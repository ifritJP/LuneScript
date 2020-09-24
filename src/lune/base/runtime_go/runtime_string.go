package main

// #include <stdlib.h>
// #cgo CFLAGS: -Ilua_link/inc
// #include <lauxlib.h>
// #include <lualib.h>
import "C"

import "unsafe"
import "log"
import "runtime"


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

    callInfo := &lns_callInfoString{}

    vm := luaVM.vm
    callInfo.top = int(C.lua_gettop( vm ));

    C.lua_getglobal( vm, lns_c_ptr_string );

    callInfo.funcname = C.CString( funcName )
    C.lua_getfield( vm, -1, callInfo.funcname )

    callInfo.string = C.CString( txt )
    C.lua_pushstring( vm, callInfo.string )

    runtime.SetFinalizer( callInfo, func (obj *lns_callInfoString) { obj.end() } )
    
    return callInfo
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
            pVal := luaVM.pushStr( val.(string) )
            defer pVal.free()
        default:
            log.Fatal( "not support default ", v )
        }
    }

    // 関数を実行
    return luaVM.lua_call( callInfo.top, len( ddd ) + 1, 1 )[0].(string)
}

func (luaVM *Lns_luaVM) string_gsub(
    format string, src string, dst string ) (string,LnsInt) {

    callInfo := luaVM.string_call_setup( "gsub", format )

    pSrc := luaVM.pushStr( src )
    defer pSrc.free()
    pDst := luaVM.pushStr( dst )
    defer pDst.free()

    // 関数を実行
    ret := luaVM.lua_call( callInfo.top, 3, 2 );
    return ret[0].(string), LnsInt(ret[1].(C.longlong))
}


func init() {
    lns_c_ptr_string = C.CString( "string" )
}
