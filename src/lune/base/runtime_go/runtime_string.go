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
import "fmt"
import "strings"
//import "runtime"

var lns_c_ptr_string *C.char
var lns_c_ptr_format *C.char
var lns_c_ptr_gsub *C.char
var lns_c_ptr_find *C.char
var lns_c_ptr_byte *C.char
var lns_c_ptr_sub *C.char
var lns_c_ptr_reverse *C.char
var lns_c_ptr_gmatch *C.char
var lns_c_ptr_dump *C.char

func init() {
    lns_c_ptr_string = C.CString( "string" )
    lns_c_ptr_format = C.CString( "format" )
    lns_c_ptr_gsub = C.CString( "gsub" )
    lns_c_ptr_find = C.CString( "find" )
    lns_c_ptr_byte = C.CString( "byte" )
    lns_c_ptr_sub = C.CString( "sub" )
    lns_c_ptr_reverse = C.CString( "reverse" )
    lns_c_ptr_gmatch = C.CString( "gmatch" )
    lns_c_ptr_dump = C.CString( "dump" )
}

type lns_callInfoString struct {
    top int
    string *C.char
}

func (luaVM *Lns_luaVM) string_static_call_setup( funcName *C.char ) int {
    vm := luaVM.vm
    top := int(lua_gettop( vm ))

    lua_getglobal( vm, lns_c_ptr_string )
    lua_getfield( vm, -1, funcName )

    return top
}


func (luaVM *Lns_luaVM) string_call_setup(
    funcName *C.char, txt string ) *lns_callInfoString {

    callInfo := &lns_callInfoString{}

    vm := luaVM.vm
    callInfo.top = luaVM.string_static_call_setup( funcName )

    callInfo.string = C.CString( txt )
    lua_pushlstring( vm, callInfo.string, len( txt ) )

    //runtime.SetFinalizer( callInfo, func (obj *lns_callInfoString) { obj.end() } )
    
    return callInfo
}


func (self *lns_callInfoString) call( luaVM *Lns_luaVM, retNum int ) []LnsAny {
    now := int(lua_gettop( luaVM.vm ))

    ret := luaVM.lua_call( self.top, now - self.top - 2, retNum )

    self.end()
    
    return ret
}


func (self *lns_callInfoString) end() {
    C.free( unsafe.Pointer( self.string ) )
}


func (luaVM *Lns_luaVM) String_format( format string, ddd []LnsAny ) string {

    callInfo := luaVM.string_call_setup( lns_c_ptr_format, format )
    
    // ... の値を push する
    for _, val := range( ddd ) {
        switch val.(type) {
        case *LnsList:
            pVal := luaVM.pushAny( fmt.Sprintf("table:%v", val ) )
            defer pVal.free()
        case *LnsSet:
            pVal := luaVM.pushAny( fmt.Sprintf("table:%v", val ) )
            defer pVal.free()
        case *LnsMap:
            pVal := luaVM.pushAny( fmt.Sprintf("table:%v", val ) )
            defer pVal.free()
        default:
            pVal := luaVM.pushAny( val )
            defer pVal.free()
        }
    }

    return callInfo.call( luaVM, 1 )[0].(string)
}

func (luaVM *Lns_luaVM) String_gsub(
    txt string, src string, dst string ) (string,LnsInt) {

    callInfo := luaVM.string_call_setup( lns_c_ptr_gsub, txt )

    pSrc := luaVM.pushStr( src )
    defer pSrc.free()
    pDst := luaVM.pushStr( dst )
    defer pDst.free()

    ret := callInfo.call( luaVM, 2 );
    return ret[0].(string), LnsInt(ret[1].(int))
}

func (luaVM *Lns_luaVM) String_find(
    txt string, src string, index LnsAny, plain LnsAny ) (LnsAny,LnsAny) {

    callInfo := luaVM.string_call_setup( lns_c_ptr_find, txt )

    pSrc := luaVM.pushStr( src )
    defer pSrc.free()
    luaVM.pushAny( index )
    luaVM.pushAny( plain )
    
    ret := callInfo.call( luaVM, 2 )
    return Lns_getFromMulti( ret, 0 ), Lns_getFromMulti( ret, 1 )
}

func (luaVM *Lns_luaVM) String_byte(
    txt string, from LnsAny, to LnsAny ) []LnsAny {

    callInfo := luaVM.string_call_setup( lns_c_ptr_byte, txt )

    luaVM.pushAny( from )
    luaVM.pushAny( to )
    
    ret := callInfo.call( luaVM, cLUA_MULTRET )
    return ret
}

func (luaVM *Lns_luaVM) String_rep( txt string, num LnsInt ) string {

    ret := ""
    for count := 1; count <= num; count++ {
        ret += txt;
    }
    return ret
}

func (luaVM *Lns_luaVM) String_sub( txt string, from LnsInt, to LnsAny ) string {
    callInfo := luaVM.string_call_setup( lns_c_ptr_sub, txt )

    luaVM.pushAny( from )
    luaVM.pushAny( to )
    return callInfo.call( luaVM, 1 )[0].(string)
}

func (luaVM *Lns_luaVM) String_lower( txt string ) string {
    return strings.ToLower( txt )
}

func (luaVM *Lns_luaVM) String_upper( txt string ) string {
    return strings.ToUpper( txt )
}

func (luaVM *Lns_luaVM) String_reverse( txt string ) string {
    callInfo := luaVM.string_call_setup( lns_c_ptr_reverse, txt )
    return callInfo.call( luaVM, 1 )[0].(string)
}


func (luaVM *Lns_luaVM) String_gmatch( txt string, pat string ) (LnsAny, LnsAny, LnsAny) {

    callInfo := luaVM.string_call_setup( lns_c_ptr_gmatch, txt )

    pSrc := luaVM.pushStr( pat )
    defer pSrc.free()
    
    ret := callInfo.call( luaVM, 3 )
    return Lns_getFromMulti( ret, 0 ),Lns_getFromMulti( ret, 1 ),Lns_getFromMulti( ret, 2 )
}

func (luaVM *Lns_luaVM) String_dump( form *Lns_luaValue, flag LnsAny ) string {

    top := luaVM.string_static_call_setup( lns_c_ptr_dump )

    form.pushValFromGlobalValMap()
    luaVM.pushAny( flag )
    
    ret := luaVM.lua_call( top, 2, 1 )
    return ret[0].(string)
}
