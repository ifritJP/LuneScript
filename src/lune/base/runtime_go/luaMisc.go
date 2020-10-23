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

//import "unsafe"
//import "log"
//import "runtime"

func (luaVM *Lns_luaVM) ExpandLuavalMap( stem LnsAny ) LnsAny {
    if Lns_IsNil( stem ) {
        return nil
    }

    pVal := luaVM.pushAny( stem )
    defer pVal.free()
    return luaVM.setupFromStack( -1, false );
}

func (luaVM *Lns_luaVM) Loadfile( path string ) (LnsAny, LnsAny) {

    ret := luaVM.CallStatic( "", "loadfile", []LnsAny{ path } )
    if len( ret ) == 1 {
        return ret[ 0 ], nil
    }
    return ret[ 0 ], ret[ 1 ]
}

func (luaVM *Lns_luaVM) Load( txt string, opt LnsAny ) (LnsAny, LnsAny) {

    ret := luaVM.CallStatic( "", "load", []LnsAny{ txt, opt } )
    if len( ret ) == 1 {
        return ret[ 0 ], nil
    }
    return ret[ 0 ], ret[ 1 ]
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

    ret2 := luaValue.luaVM.setupFromStack( -1, true )
    ret1 := luaValue.luaVM.setupFromStack( -2, true )

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

    ret2 := luaVM.setupFromStack( -1, true )
    ret1 := luaVM.setupFromStack( -2, true )

    return ret1, ret2
}



//===============
func (luaVM *Lns_luaVM) IO_open( path string, opt LnsAny ) LnsAny {
    ret := luaVM.CallStatic( "io", "open", []LnsAny{ path, opt } )
    return Lns_getFromMulti( ret, 0 )
}
func (luaVM *Lns_luaVM) IO_popen( path string ) LnsAny {
    ret := luaVM.CallStatic( "io", "popen", []LnsAny{ path } )
    return Lns_getFromMulti( ret, 0 )
}
func (luaVM *Lns_luaVM) Package_searchpath( name string, path string ) LnsAny {
    ret := luaVM.CallStatic( "package", "searchpath", []LnsAny{ name, path } )
    return Lns_getFromMulti( ret, 0 )
}
func (luaVM *Lns_luaVM) OS_clock() LnsReal {
    ret := luaVM.CallStatic( "os", "clock", []LnsAny{} )
    clock := Lns_getFromMulti( ret, 0 )
    if val, ok := clock.(LnsReal); ok {
        return val;
    }
    return LnsReal(clock.(LnsInt))
}
func (luaVM *Lns_luaVM) OS_exit( code LnsAny ) {
    luaVM.CallStatic( "os", "exit", []LnsAny{ code } )
}
func (luaVM *Lns_luaVM) OS_remove( path string ) (LnsAny, LnsAny) {
    ret := luaVM.CallStatic( "os", "remove", []LnsAny{ path } )
    return Lns_getFromMulti( ret, 0 ), Lns_getFromMulti( ret, 1 )
}
func (luaVM *Lns_luaVM) OS_date( format LnsAny, time LnsAny ) LnsAny {
    ret := luaVM.CallStatic( "os", "date", []LnsAny{ format, time } )
    return Lns_getFromMulti( ret, 0 )
}
func (luaVM *Lns_luaVM) OS_time( tbl LnsAny ) LnsAny {
    ret := luaVM.CallStatic( "os", "time", []LnsAny{ tbl } )
    return Lns_getFromMulti( ret, 0 )
}
func (luaVM *Lns_luaVM) OS_difftime( tm1 LnsAny, tm2 LnsAny ) LnsInt {
    ret := luaVM.CallStatic( "os", "difftime", []LnsAny{ tm1, tm2 } )
    return Lns_getFromMulti( ret, 0 ).(LnsInt)
}
func (luaVM *Lns_luaVM) OS_rename( src string, dst string ) (LnsAny, LnsAny) {
    ret := luaVM.CallStatic( "os", "rename", []LnsAny{ src, dst } )
    return Lns_getFromMulti( ret, 0 ), Lns_getFromMulti( ret, 0 )
}
func (luaVM *Lns_luaVM) Math_random( mVal LnsAny, nVal LnsAny ) LnsReal {
    var args []LnsAny
    if nVal == nil && mVal == nil {
        args = []LnsAny{}
    } else if nVal == nil {
        args = []LnsAny{ mVal }
    } else {
        args = []LnsAny{ mVal, nVal }
    }
    ret := luaVM.CallStatic( "math", "random", args )
    return Lns_getFromMulti( ret, 0 ).(LnsReal)
}
func (luaVM *Lns_luaVM) Math_randomseed( seed LnsInt ) {
    luaVM.CallStatic( "math", "randomseed", []LnsAny{ seed } )
}
