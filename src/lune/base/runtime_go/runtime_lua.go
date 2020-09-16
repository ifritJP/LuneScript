package main

// #include <stdlib.h>
// #cgo CFLAGS: -Ilua_link/inc
// #cgo LDFLAGS: -ldl -lm -llua5.3
// #include <lauxlib.h>
// #include <lualib.h>
import "C"

import "unsafe"
import "log"


type Lns_luaVM struct {
    vm *C.lua_State
}

var defaultVM Lns_luaVM

func Lns_getVM() *Lns_luaVM {
    return &defaultVM
}

func Lns_runLuaScript( script string ) {
    var vm * C.lua_State = C.luaL_newstate()
    if vm == nil {
        return
    }
    defer C.lua_close( vm )
    C.luaL_openlibs( vm )

    block := C.CString( script )
    defer C.free( unsafe.Pointer( block ) )
    C.luaL_loadstring( vm, block )
    C.lua_pcallk( vm, 0, C.LUA_MULTRET, 0, 0, nil )
}

/**
スタックの top の値から pStem を設定する。
 */
func (luaVM *Lns_luaVM) setupFromStack( index C.int ) LnsAny {
    vm := luaVM.vm
    switch typeId := C.lua_type( vm, index ); typeId {
    case C.LUA_TNIL:
        return nil
    case C.LUA_TNUMBER:
        num := C.lua_tonumberx( vm, index, nil )
        intVal := C.lua_tointegerx( vm, index, nil )
        if C.double(LnsInt(num)) == num {
            return intVal
        }
        return LnsReal(num)
    case C.LUA_TBOOLEAN:
        return C.lua_toboolean( vm, index ) != 0
    case C.LUA_TSTRING:
        return C.GoString( C.lua_tolstring( vm, index, nil ) )
    case C.LUA_TFUNCTION:
        // {
        //     *pStem = LNS_STEM_ANY( lns_createFormAny( _pEnv ) );
        //     C.lua_pushvalue( vm, index );
        //     lns_setAnyVal( _pEnv, pStem->val.pAny );
        // }
        log.Fatal( "not support -- func" )
      default:
        // {
        //     *pStem = LNS_STEM_ANY( lns_luaVal_new( _pEnv, lns_value_type_luaVal ) );
        //     C.lua_pushvalue( vm, index );
        //     lns_setAnyVal( _pEnv, pStem->val.pAny );
        // }
        log.Fatal( "not support -- default ", typeId )
    }
    return nil
}


/**
 * Lua の関数コール処理。
 *
 * @param stackBase 関数コール前の基準スタック位置。
 *    この関数処理後は、stackBase 位置をスタックトップとする。
 * @param argNum Lua 関数に渡す引数の数
 * @param retNum Lua 関数の戻り値の数。 不定の場合 LUA_MULTRET。
 * @param Lua 関数の戻り値。
 */
func (luaVM *Lns_luaVM) lua_call( stackBase int , argNum int, retNum int ) []LnsAny {
    top := int(C.lua_gettop( luaVM.vm )) - argNum - 1
    C.lua_callk( luaVM.vm, C.int(argNum), C.LUA_MULTRET, C.long(0), nil )

    lastRet := int(C.lua_gettop( luaVM.vm ))

    result := []LnsAny{}

    stackNum := lastRet - top
    if retNum >= 2 || retNum == C.LUA_MULTRET {
        for index := top + 1; index <= lastRet; index++ {
            result = append( result, luaVM.setupFromStack( C.int(index) ) )
        }
    } else if ( stackNum == 1 ) {
        result = append( result, luaVM.setupFromStack( C.int(-1) ) )
    }

    C.lua_settop( luaVM.vm, C.int( stackBase ) )

    return result
}



var lns_c_ptr_string *C.char

type lns_callInfoString struct {
    top int
    funcname *C.char
    string *C.char
}

func (self *lns_callInfoString) end() {
    C.free( unsafe.Pointer( self.funcname ) )
    C.free( unsafe.Pointer( self.string ) )
}

func (luaVM *Lns_luaVM) string_call_setup(
    funcName string, txt string ) *lns_callInfoString {

    callInfo := lns_callInfoString{}

    vm := luaVM.vm
    callInfo.top = int(C.lua_gettop( vm ));

    C.lua_getglobal( vm, lns_c_ptr_string );

    callInfo.funcname = C.CString( funcName )
    C.lua_getfield( vm, -1, callInfo.funcname )

    callInfo.string = C.CString( txt )
    C.lua_pushstring( vm, callInfo.string )
    
    return &callInfo
}


func (luaVM *Lns_luaVM) string_format( format string, ddd []LnsAny ) string {
    vm := luaVM.vm

    callInfo := luaVM.string_call_setup( "format", format )
    
    // ... の値を push する
    for _, val := range( ddd ) {
        switch v := val.(type) {
        case LnsInt:
            C.lua_pushinteger( vm, C.longlong( val.(LnsInt)) )
        case LnsReal:
            C.lua_pushnumber( vm, C.double(val.(LnsReal)) )
        case bool:
            if val.(bool) {
                C.lua_pushboolean( vm, C.int(1) )
            } else {
                C.lua_pushboolean( vm, C.int(0) )
            }
        case string:
            C.lua_pushstring( vm, C.CString( val.(string)) )
        default:
            log.Fatal( "not support default ", v )
        }
    }

    // 関数を実行
    return luaVM.lua_call( callInfo.top, len( ddd ) + 1, 1 )[0].(string)
}


func init() {
    defaultVM.vm = C.luaL_newstate()
    C.luaL_openlibs( defaultVM.vm )

    lns_c_ptr_string = C.CString( "string" )

}

// func main() {
//     lns_runLuaScript( "print( 'hello world' )" )
// }
