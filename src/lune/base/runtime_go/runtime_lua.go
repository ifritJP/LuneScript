package main

// #include <stdlib.h>
// #cgo CFLAGS: -I../../../../../lctags/external/lua/lua-5.3.4/src
// #cgo LDFLAGS: -ldl -lm -llua53 -L../../../../../lctags/external/lua/lua-5.3.4/src
// #include <lauxlib.h>
// #include <lualib.h>
import "C"

import "unsafe"
//import "fmt"


type lns_luaVM struct {
    vm * C.lua_State
}

func lns_runLuaScript( script string ) {
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

func main() {
    lns_runLuaScript( "print( 'hello world' )" )
}
