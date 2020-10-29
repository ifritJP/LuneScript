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
import "regexp"
//import "runtime"

import "github.com/golang/groupcache/lru"

func Lns_Str_init() {
}

func Str_getLineList( txt string ) *LnsList {
    list := NewLnsList( []LnsAny{} )
    for _, line := range( strings.SplitAfter( txt, "\n" ) ) {
        list.Insert( line )
    }
    return list
}

func Str_startsWith( txt, ptn string ) bool {
    return strings.HasPrefix( txt, ptn )
}

func Str_endsWith( txt, ptn string ) bool {
    return strings.HasSuffix( txt, ptn )
}



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

    regexpCache = lru.New( 20 )
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

var regexpCache *lru.Cache

func createRegexp( src string ) *regexp.Regexp {

    if exp, ok := regexpCache.Get( src ); ok {
        return exp.(*regexp.Regexp);
    }
    
    offset := 0
    remain := len( src )
    
    processQuote := func ( inClass bool ) string {
        remain--
        offset++
        if remain <= 0 {
            return ""
        }
        cls := ""
        switch src[ offset ] {
        case 'a':
            cls = "[:alpha:]"
        case 'c':
            cls = "[:cntrl:]"
        case 'd':
            cls = "[:digit:]"
        case 'g':
            cls = "[:graph:]"
        case 'l':
            cls = "[:lower:]"
        case 'p':
            cls = "[:punct:]"
        case 's':
            cls = "[:space:]"
        case 'u':
            cls = "[:upper:]"
        case 'w':
            cls = "[:alnum:]"
        case 'x':
            cls = "[:xdigit:]"
        default:
            return fmt.Sprintf( "\\x%02X", src[ offset ] )
        }
        if inClass {
            return cls
        }
        return fmt.Sprintf( "[%s]", cls )
    }

    reg := ""
    if remain >= 2 && src[ offset ] == '^' {
        reg += "^"

        offset++
        remain--
    }

    
    inClass := false
    firstFlag := true
    for remain > 0 {
        switch src[ offset ] {
        case '%':
            if quote := processQuote( inClass ); quote == "" {
                return nil
            } else {
                reg += quote
            }
        case '[':
            if inClass {
                return nil
            }
            inClass = true
            reg += src[ offset:offset + 1 ]
        case ']':
            if !inClass {
                return nil
            }
            inClass = false
            reg += src[ offset:offset + 1 ]
        case '^':
            reg += src[ offset:offset + 1 ]
        case '$':
            reg += src[ offset:offset + 1 ]
        case '*':
            reg += src[ offset:offset + 1 ]
        case '+':
            reg += src[ offset:offset + 1 ]
        case '-':
            if inClass || firstFlag {
                reg += src[ offset:offset + 1 ]
            } else {
                return nil
            }
        case '|':
            reg += fmt.Sprintf( "\\x%02X", src[ offset ] )
        case '{':
            reg += fmt.Sprintf( "\\x%02X", src[ offset ] )
        case '}':
            reg += fmt.Sprintf( "\\x%02X", src[ offset ] )
        case '(':
            return nil
        case '\\':
            reg += fmt.Sprintf( "\\x%02X", src[ offset ] )
        case ':':
            reg += fmt.Sprintf( "\\x%02X", src[ offset ] )
        default:
            reg += src[ offset:offset + 1 ]
        }

        if firstFlag {
            firstFlag = false
        }
        remain--
        offset++        
    }

    re := regexp.MustCompile( reg )
    regexpCache.Add( src, re )
    
    return re
}


func goFind( txt string, src string, index LnsInt ) (bool, []LnsAny) {

    dummyRet := []LnsAny{ nil }

    re := createRegexp( src )
    if re == nil {
        return false, dummyRet
    }

    loc := re.FindStringIndex( txt )
    if loc == nil {
        return true, dummyRet
    }

    if index <= 0 {
        index = 1
    }
    return true, []LnsAny{ index + loc[0], index  + loc[1] - 1 }
}

// goFind の結果と lua ラインタイムの find の結果を比べるかどうか。
// デバッグ用途。
var checkFindResult = false
func (luaVM *Lns_luaVM) String_find(
    txt string, src string, index LnsAny, plain LnsAny ) []LnsAny {

    target := txt
    offset := 1
    if work, ok := index.(LnsInt); ok {
        offset = string_index( work, len( txt ) )
        if offset > 0 {
            target = txt[ offset - 1: ]
        }
    }
    
    if plain == true {
        findIndex := strings.Index( target, src )
        if findIndex < 0 {
            return []LnsAny{}
        }
        return []LnsAny{ offset + findIndex, offset + findIndex + len( src ) - 1 }
    }
    var goResult []LnsAny
    if ok, result := goFind( target, src, offset ); ok {
        if !checkFindResult {
            return result;
        }
        goResult = result
    }
    print( fmt.Sprintf( "log: fallback -- string.find()  %s, %s, %v\n", txt, src, index ) )


    // goFind で処理できなかった場合は、 lua のランタイムで処理する
    

    callInfo := luaVM.string_call_setup( lns_c_ptr_find, txt )

    pSrc := luaVM.pushStr( src )
    defer pSrc.free()
    luaVM.pushAny( index ) // int なので defer 不要
    luaVM.pushAny( plain ) // bool なので defer 不要
    
    ret := callInfo.call( luaVM, 4, 2 )

    if checkFindResult {
        if len( ret ) != len( goResult ) ||
            ret[ 0 ] != goResult[ 0 ] ||
            ( len( ret ) > 1 && ret[ 1 ] != goResult[ 1 ] ) {
            mess := fmt.Sprintf(
                "hoge: %s %s %d %d %d %s %s\n", txt, src, index,
                len( ret ), len( goResult ), ret[0], goResult[0] )
            if len( ret ) > 1 && len( goResult ) > 1 {
                mess = fmt.Sprintf( "%s %s %s\n", mess, ret[1], goResult[1] )
            }
            panic( mess )
        }
    }
    
    return ret
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
