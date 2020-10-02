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


package main

// #include <stdlib.h>
import "C"

import "unsafe"
//import "log"
//import "runtime"


func (luaVM *Lns_luaVM) Load( txt string, opt LnsAny ) (LnsAny, LnsAny) {

    top := lua_gettop( luaVM.vm )
    defer lua_settop( luaVM.vm, top )
    
    vm := luaVM.vm
    pTxt := C.CString( txt )
    defer C.free( unsafe.Pointer( pTxt ) )
    if luaL_loadstring( vm, pTxt ) == cLUA_OK {
        return luaVM.newLuaValue( -1 ), nil
    }
    return nil, luaVM.setupFromStack( -1 )
}

func (luaVM *Lns_luaVM) RunLoadedfunc( loaded *Lns_luaValue, args[]LnsAny ) []LnsAny {

    top := lua_gettop( luaVM.vm )
    defer lua_settop( luaVM.vm, top )

    loaded.pushValFromGlobalValMap()
    for _, val := range( args ) {
        luaVM.pushAny( val )
    }
    return luaVM.lua_call( top, len( args ), cLUA_MULTRET )
}

func (luaValue *Lns_luaValue) Get1stFromMap() (LnsAny, LnsAny) {
    vm := luaValue.luaVM.vm
    top := lua_gettop( vm )
    defer lua_settop( vm, top )

    luaValue.pushValFromGlobalValMap()
    
    lua_pushnil( vm )
    if lua_next( vm, -2 ) == 0 {
        return nil, nil
    }

    ret2 := luaValue.luaVM.setupFromStack( -1 )
    ret1 := luaValue.luaVM.setupFromStack( -2 )

    return ret1, ret2
}

func (luaValue *Lns_luaValue) NextFromMap( prev LnsAny ) (LnsAny, LnsAny) {
    vm := luaValue.luaVM.vm
    top := lua_gettop( vm )
    defer lua_settop( vm, top )

    luaVM := luaValue.luaVM

    luaValue.pushValFromGlobalValMap()
    luaVM.pushAny( prev )

    if lua_next( vm, -2 ) == 0 {
        return nil, nil
    }

    ret2 := luaVM.setupFromStack( -1 )
    ret1 := luaVM.setupFromStack( -2 )

    return ret1, ret2
}
