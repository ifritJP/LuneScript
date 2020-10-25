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
import "strings"
//import "runtime"

// import "github.com/yuin/gopher-lua/pm"

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
    C.free( unsafe.Pointer( callInfo.string ) )

    //runtime.SetFinalizer( callInfo, func (obj *lns_callInfoString) { obj.end() } )
    
    return callInfo
}


func (self *lns_callInfoString) call(
    luaVM *Lns_luaVM, argNum int, retNum int ) []LnsAny {
    ret := luaVM.lua_call( self.top, argNum, retNum )

    self.end()
    
    return ret
}


func (self *lns_callInfoString) end() {
    //C.free( unsafe.Pointer( self.string ) )
}

func string_index( index LnsInt, len LnsInt ) LnsInt {
    if index >= 0 {
        return index;
    } else if -index > len {
        return 0;
    }
    return len + index + 1;
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

    return callInfo.call( luaVM, 1 + len( ddd ), 1 )[0].(string)
}

func (luaVM *Lns_luaVM) String_gsub(
    txt string, src string, dst string ) (string,LnsInt) {

    callInfo := luaVM.string_call_setup( lns_c_ptr_gsub, txt )

    pSrc := luaVM.pushStr( src )
    defer pSrc.free()
    pDst := luaVM.pushStr( dst )
    defer pDst.free()

    ret := callInfo.call( luaVM, 3, 2 );
    return ret[0].(string), LnsInt(ret[1].(int))
}

func (luaVM *Lns_luaVM) String_find(
    txt string, src string, index LnsAny, plain LnsAny ) []LnsAny {

    callInfo := luaVM.string_call_setup( lns_c_ptr_find, txt )

    pSrc := luaVM.pushStr( src )
    defer pSrc.free()
    luaVM.pushAny( index ) // int なので defer 不要
    luaVM.pushAny( plain ) // bool なので defer 不要
    
    ret := callInfo.call( luaVM, 4, 2 )
    return ret

    // offset := 1
    // if work, ok := index.(LnsInt); ok {
    //     offset = work
    // }
    // if offset < 0 {
    //     return []LnsAny{}
    // } else if offset == 0 {
    //     offset = 1;
    // }
    
    // matches, err := pm.Find( src, []byte(txt), offset - 1, 1 )
	// if err != nil {
	// 	panic( err )
	// }
	// if len( matches ) == 0 {
    //     return []LnsAny{}
	// }
	// md := matches[0]
    // return []LnsAny{ md.Capture(0) + 1, md.Capture(1) }
    
	// // L.Push(LNumber(md.Capture(0) + 1))
	// // L.Push(LNumber(md.Capture(1)))
	// // for i := 2; i < md.CaptureLength(); i += 2 {
	// // 	if md.IsPosCapture(i) {
	// // 		L.Push(LNumber(md.Capture(i)))
	// // 	} else {
	// // 		L.Push(LString(str[md.Capture(i):md.Capture(i+1)]))
	// // 	}
	// // }
	// // return md.CaptureLength()/2 + 1
}

func (luaVM *Lns_luaVM) String_byte(
    txt string, from LnsAny, to LnsAny ) []LnsAny {

    callInfo := luaVM.string_call_setup( lns_c_ptr_byte, txt )

    luaVM.pushAny( from ) // int なので defer 不要
    luaVM.pushAny( to ) // int なので defer 不要
    
    ret := callInfo.call( luaVM, 3, cLUA_MULTRET )
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
    // callInfo := luaVM.string_call_setup( lns_c_ptr_sub, txt )
    // luaVM.pushAny( from ) // int なので defer 不要
    // luaVM.pushAny( to ) // int なので defer 不要
    // return callInfo.call( luaVM, 3, 1 )[0].(string)

    from = string_index( from, len( txt ) )
    if from < 1 {
        from = 1
    }
    end := len( txt )
    if work, ok := to.(LnsInt); ok {
        end = string_index( work, len( txt ) )
    }
    if from <= end {
        if end > len(txt) {
            end = len(txt)
        }
        return txt[from - 1:end]
    }
    return ""
}

func (luaVM *Lns_luaVM) String_lower( txt string ) string {
    return strings.ToLower( txt )
}

func (luaVM *Lns_luaVM) String_upper( txt string ) string {
    return strings.ToUpper( txt )
}

func (luaVM *Lns_luaVM) String_reverse( txt string ) string {
    callInfo := luaVM.string_call_setup( lns_c_ptr_reverse, txt )
    return callInfo.call( luaVM, 1, 1 )[0].(string)
}


func (luaVM *Lns_luaVM) String_gmatch( txt string, pat string ) (LnsAny, LnsAny, LnsAny) {

    callInfo := luaVM.string_call_setup( lns_c_ptr_gmatch, txt )

    pSrc := luaVM.pushStr( pat )
    defer pSrc.free()
    
    ret := callInfo.call( luaVM, 2, 3 )
    return Lns_getFromMulti( ret, 0 ),Lns_getFromMulti( ret, 1 ),Lns_getFromMulti( ret, 2 )
}

func (luaVM *Lns_luaVM) String_dump( form *Lns_luaValue, flag LnsAny ) string {

    top := luaVM.string_static_call_setup( lns_c_ptr_dump )

    form.core.pushValFromGlobalValMap()
    luaVM.pushAny( flag ) // bool なので defer 不要
    
    ret := luaVM.lua_call( top, 2, 1 )
    return ret[0].(string)
}
