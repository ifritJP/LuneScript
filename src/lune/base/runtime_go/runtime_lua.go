/*
MIT License

Copyright (c) 2018,2020 ifritJP

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package runtimelns

// #include <stdlib.h>
import "C"

import "unsafe"
import "fmt"
import "runtime"

var lns_globalValSym *C.char

var defaultVM Lns_luaVM

type Lns_luaValue struct {
    // lua VM の lns_globalValMap に格納しているシンボル名
    sym *C.char
    luaVM *Lns_luaVM
}

func init() {
    defaultVM.vm = luaL_newstate()
    luaL_openlibs( defaultVM.vm )

    lns_globalValSym = C.CString( "lns_globalValMap" )
    lua_createtable( defaultVM.vm )
    lua_setglobal( defaultVM.vm, lns_globalValSym )
    
    lns_defaultPushedVal = &lns_pushedVal{}
}


func Lns_getVM() *Lns_luaVM {
    return &defaultVM
}

func Lns_runLuaScript( script string ) {
    vm := luaL_newstate()
    if vm == nil {
        return
    }
    defer lua_close( vm )
    luaL_openlibs( vm )

    luaL_loadstring( vm, script )
    lua_pcallk( vm, 0, cLUA_MULTRET )
}

type lns_pushedString struct {
    str *C.char
}

func (self *lns_pushedString) free() {
    C.free( unsafe.Pointer( self.str ) )
}

func (luaVM *Lns_luaVM) pushStr( txt string ) *lns_pushedString {
    pStr := C.CBytes([]byte(txt))
    lua_pushlstring( luaVM.vm, (*C.char)(pStr), len( txt ) )
    ret := &lns_pushedString{ (*C.char)(pStr) }
    //runtime.SetFinalizer( ret, func (obj *lns_pushedString) { obj.free() } )
    return ret
}

type lns_pushedVal struct {
    pStr *lns_pushedString
}
func (obj *lns_pushedVal) free() {
    if obj.pStr != nil {
        obj.pStr.free()
        obj.pStr = nil;
    }
}
var lns_defaultPushedVal *lns_pushedVal

func (luaVM *Lns_luaVM) pushAny( val LnsAny ) *lns_pushedVal {
    vm := luaVM.vm
    if Lns_IsNil( val ) {
        lua_pushnil( vm )
    } else {
        switch val.(type) {
        case LnsInt:
            lua_pushinteger( vm, val.(LnsInt) )
        case LnsReal:
            lua_pushnumber( vm, val.(LnsReal) )
        case bool:
            lua_pushboolean( vm, val.(bool) )
        case string:
            return &lns_pushedVal{ luaVM.pushStr( val.(string) ) }
        case *Lns_luaValue:
            val.(*Lns_luaValue).pushValFromGlobalValMap()
        default:
            panic( fmt.Sprintf( "not supoort -- %v", val ) )
        }
    }
    return lns_defaultPushedVal
}

func (luaVM *Lns_luaVM) newLuaValue( index int ) *Lns_luaValue {
    val := &Lns_luaValue{ luaVM: luaVM }
    val.sym = C.CString( fmt.Sprintf( "%p", &val ) )
    runtime.SetFinalizer( val, func (obj *Lns_luaValue) { obj.free() } )
    val.setValToGlobalValMap( index )
    return val
}

/**
スタックの指定 index の値を取得する
 */
func (luaVM *Lns_luaVM) setupFromStack( index int ) LnsAny {
    vm := luaVM.vm
    switch typeId := lua_type( vm, index ); typeId {
    case cLUA_TNIL:
        return nil
    case cLUA_TNUMBER:
        num := lua_tonumberx( vm, index )
        intVal := LnsInt(lua_tointegerx( vm, index ))
        if LnsReal(LnsInt(num)) == num {
            return intVal
        }
        return LnsReal(num)
    case cLUA_TBOOLEAN:
        return lua_toboolean( vm, index )
    case cLUA_TSTRING:
        return lua_tolstring( vm, index )
    default:
        return luaVM.newLuaValue( -1 )
    }
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
    argTop := int(lua_gettop( luaVM.vm )) - argNum
    lua_callk( luaVM.vm, argNum, cLUA_MULTRET )

    lastRet := int(lua_gettop( luaVM.vm ))

    result := []LnsAny{}

    stackNum := lastRet - argTop + 1
    if retNum >= 2 || retNum == cLUA_MULTRET {
        for index := argTop; index <= lastRet; index++ {
            result = append( result, luaVM.setupFromStack( index ) )
        }
    } else if ( stackNum == 1 ) {
        result = append( result, luaVM.setupFromStack( -1 ) )
    }

    lua_settop( luaVM.vm, stackBase )

    return result
}


/**
symbol を globalVal から除外する
*/
func (luaValue *Lns_luaValue) free() {
    lua_pushnil( luaValue.luaVM.vm )
    luaValue.setValToGlobalValMap( -1 )
    lua_pop( luaValue.luaVM.vm, 1 )

    C.free( unsafe.Pointer( luaValue.sym ) )
    luaValue.sym = nil
}

func (luaValue *Lns_luaValue) Call( argList []LnsAny ) []LnsAny {

    luaVM := luaValue.luaVM

    top := lua_gettop( luaVM.vm )
    
    luaValue.pushValFromGlobalValMap()
    for _, arg := range( argList ) {
        pVal := luaVM.pushAny( arg )
        defer pVal.free()
    }

    return luaVM.lua_call( top, len( argList ), cLUA_MULTRET )
}


/**
index 指定のスタックの値を globalVal[ symbol ] にセットする
*/
func (luaValue *Lns_luaValue) setValToGlobalValMap( index int ) {
    vm := luaValue.luaVM.vm
    // globalVal をスタックトップに push
    lua_getglobal( vm, lns_globalValSym )
    // 指定の値をスタックトップに push
    if index < 0 {
        lua_pushvalue( vm, index - 1 )
    } else {
        lua_pushvalue( vm, index )
    }
    // スタックトップの値を globalVal[ symbol ] にセットし、スタックトップは削除。
    lua_setfield( vm, -2, luaValue.sym )
    // globalVal をスタックから除外
    lua_pop( vm, 1 )
}

/**
globalVal[ symbol ] をスタックトックに push する
*/
func (luaValue *Lns_luaValue) pushValFromGlobalValMap() {
    vm := luaValue.luaVM.vm
    // globalVal をスタックトップに push
    lua_getglobal( vm, lns_globalValSym )
    // globalVal[ symbol ] をスタックトップに push
    lua_getfield( vm, -1, luaValue.sym )
    // globalVal をスタックから除外
    lua_copy( vm, -1, -2 )
    lua_pop( vm, 1 )
}


func (luaVM *Lns_luaVM) CallStatic(
    packName string, funcname string, args[] LnsAny ) []LnsAny {
    
    vm := luaVM.vm
    top := lua_gettop( vm )
    defer lua_settop( vm, top )

    var argPos int
    if packName == "" {
        argPos = 1
        pFuncname := C.CString( funcname )
        defer C.free( unsafe.Pointer( pFuncname ) )
        lua_getglobal( vm, pFuncname )
    } else {
        argPos = 2
        pPackName := C.CString( packName )
        defer C.free( unsafe.Pointer( pPackName ) )
        pFuncname := C.CString( funcname )
        defer C.free( unsafe.Pointer( pFuncname ) )
        lua_getglobal( vm, pPackName )
        lua_getfield(vm, -1, pFuncname )
    }

    
    for _, val := range( args ) {
        pVal := luaVM.pushAny( val )
        defer pVal.free()
    }
    if lua_pcallk( vm, len( args ), cLUA_MULTRET ) != cLUA_OK {
        panic( lua_tolstring( vm, -1 ) )
    }
    ret := []LnsAny{}
    nowTop := lua_gettop( vm )
    for index := top + argPos; index <= nowTop; index++ {
        ret = append( ret, luaVM.setupFromStack( index ) )
    }
    
    return ret
}

func (obj *Lns_luaValue) CallMethod( funcname string, args[] LnsAny ) []LnsAny {
    luaVM := obj.luaVM
    
    vm := luaVM.vm
    top := lua_gettop( vm )
    defer lua_settop( vm, top )

    pFuncname := C.CString( funcname )
    defer C.free( unsafe.Pointer( pFuncname ) )

    obj.pushValFromGlobalValMap() // obj を push
    lua_getfield(vm, -1, pFuncname ) // obj の pFuncname のメソッドを取得
    obj.pushValFromGlobalValMap() // obj を引数に push
    
    for _, val := range( args ) {
        pVal := luaVM.pushAny( val ) // arg を push
        defer pVal.free()
    }
    if lua_pcallk( vm, len( args ) + 1, cLUA_MULTRET ) != cLUA_OK {
        panic( lua_tolstring( vm, -1 ) )
    }
    ret := []LnsAny{}
    nowTop := lua_gettop( vm )
    for index := top + 2; index <= nowTop; index++ {
        ret = append( ret, luaVM.setupFromStack( index ) )
    }
    
    return ret
}
