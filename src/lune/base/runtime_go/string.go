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

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/golang/groupcache/lru"
)

//import "runtime"

func lns_Str_init() {
}

func str_getLineList(txt string) *LnsList {
	splited := strings.SplitAfter(txt, "\n")
	list := make([]LnsAny, len(splited))
	for index, line := range splited {
		list[index] = line
	}

	return NewLnsList(list)
}

func str_startsWith(txt, ptn string) bool {
	return strings.HasPrefix(txt, ptn)
}

func str_endsWith(txt, ptn string) bool {
	return strings.HasSuffix(txt, ptn)
}

func str_isValidStrBuilder() bool {
	return true
}

// ==== Builder
type Str_Builder struct {
	txt strings.Builder
	FP  Str_BuilderMtd
}

func str_Builder2Stem(obj LnsAny) LnsAny {
	if obj == nil {
		return nil
	}
	return obj.(*Str_Builder).FP
}

type Str_BuilderDownCast interface {
	Str_ToBuilder() *Str_Builder
}

func Str_BuilderDownCastF(multi ...LnsAny) LnsAny {
	if len(multi) == 0 {
		return nil
	}
	obj := multi[0]
	if ddd, ok := multi[0].([]LnsAny); ok {
		obj = ddd[0]
	}
	work, ok := obj.(Str_BuilderDownCast)
	if ok {
		return work.Str_ToBuilder()
	}
	return nil
}

func (obj *Str_Builder) str_ToBuilder() *Str_Builder {
	return obj
}

func newStr_Builder() *Str_Builder {
	obj := &Str_Builder{}
	obj.FP = obj
	obj.initStr_Builder()
	return obj
}
func (self *Str_Builder) initStr_Builder() {
	self.txt.Reset()
}

func (self *Str_Builder) get_txt() string {
	return self.txt.String()
}

// 106: decl @lune.@base.@Util.memStream.write
func (self *Str_Builder) add(val string) {
	self.txt.WriteString(val)
}

func (self *Str_Builder) len() LnsInt {
	return self.txt.Len()
}

func (self *Str_Builder) clear() {
	self.txt.Reset()
}

var lns_c_ptr_string lua_rawstr
var lns_c_ptr_format lua_rawstr
var lns_c_ptr_gsub lua_rawstr
var lns_c_ptr_find lua_rawstr
var lns_c_ptr_byte lua_rawstr
var lns_c_ptr_sub lua_rawstr
var lns_c_ptr_reverse lua_rawstr
var lns_c_ptr_gmatch lua_rawstr
var lns_c_ptr_dump lua_rawstr

func init() {
	lns_c_ptr_string = Lns_toRawStr("string")
	lns_c_ptr_format = Lns_toRawStr("format")
	lns_c_ptr_gsub = Lns_toRawStr("gsub")
	lns_c_ptr_find = Lns_toRawStr("find")
	lns_c_ptr_byte = Lns_toRawStr("byte")
	lns_c_ptr_sub = Lns_toRawStr("sub")
	lns_c_ptr_reverse = Lns_toRawStr("reverse")
	lns_c_ptr_gmatch = Lns_toRawStr("gmatch")
	lns_c_ptr_dump = Lns_toRawStr("dump")
}

type lns_callInfoString struct {
	top    int
	rawstr lua_rawstr
}

func (luaVM *Lns_luaVM) string_static_call_setup(funcName lua_rawstr) int {
	vm := luaVM.vm
	top := int(lua_gettop(vm))

	lua_getglobal(vm, lns_c_ptr_string)
	lua_getfield(vm, -1, funcName)

	return top
}

func (luaVM *Lns_luaVM) string_call_setup(
	funcName lua_rawstr, txt string) *lns_callInfoString {

	callInfo := &lns_callInfoString{}

	vm := luaVM.vm
	callInfo.top = luaVM.string_static_call_setup(funcName)

	callInfo.rawstr = Lns_toRawStr(txt)
	lua_pushlstring(vm, callInfo.rawstr, len(txt))
	Lns_freeRawStr(callInfo.rawstr)

	//runtime.SetFinalizer( callInfo, func (obj *lns_callInfoString) { obj.end() } )

	return callInfo
}

func (self *lns_callInfoString) call(
	luaVM *Lns_luaVM, argNum int, retNum int) []LnsAny {
	ret := luaVM.lua_call(self.top, argNum, retNum)

	self.end()

	return ret
}

func (self *lns_callInfoString) end() {
	//C.free( unsafe.Pointer( self.string ) )
}

func string_index(index LnsInt, len LnsInt) LnsInt {
	if index >= 0 {
		return index
	} else if -index > len {
		return 0
	}
	return len + index + 1
}

func String_formatGo(format string, ddd []LnsAny) (string, bool) {
	if len(ddd) == 0 {
		return format, true
	}

	orgFormat := format

	type formatInfo struct {
		// % までの文字列
		part string
		// % の後の文字
		formType string
		// % に対する引数
		arg LnsAny
	}

	// %書式を確認し、%書式と引数の組み合わせを infoList に格納する
	infoList := []formatInfo{}
	argIndex := 0
	for {
		offset := strings.Index(format, "%")
		if offset < 0 || offset >= len(format) {
			infoList = append(infoList, formatInfo{format, "", nil})
			break
		}
		var formType string
		switch format[offset+1] {
		case 's':
			formType = "v"
		case 'd':
			formType = "d"
		case 'c':
			formType = "c"
		case 'g':
			formType = "g"
		case 'f':
			formType = "f"
		case '%':
			formType = "%"
		default:
			return "", false
		}
		var arg LnsAny
		if formType == "%" {
			arg = nil
		} else {
			if len(ddd) < argIndex {
				panic(fmt.Sprintf("error -- %d, %v", argIndex, orgFormat))
			}
			arg = ddd[argIndex]
			argIndex++
		}

		infoList = append(infoList, formatInfo{format[:offset+1], formType, arg})
		format = format[offset+2:]
	}

	// % 書式を展開する。
	goformat := strings.Builder{}
	for _, info := range infoList {
		switch info.formType {
		case "":
			goformat.WriteString(info.part)
		case "%":
			goformat.WriteString(info.part)
		case "v":
			txt := fmt.Sprintf(info.part+info.formType, Lns_ToString(info.arg))
			goformat.WriteString(txt)
		default:
			txt := fmt.Sprintf(info.part+info.formType, info.arg)
			goformat.WriteString(txt)
		}
	}
	return goformat.String(), true
}

func (luaVM *Lns_luaVM) String_format(format string, ddd []LnsAny) string {

	result, ok := String_formatGo(format, ddd)
	if ok {
		return result
	}

	callInfo := luaVM.string_call_setup(lns_c_ptr_format, format)

	// ... の値を push する
	for _, val := range ddd {
		switch val.(type) {
		case *LnsList:
			pVal := luaVM.pushAny(fmt.Sprintf("table:%v", val))
			defer pVal.free()
		case *LnsSet:
			pVal := luaVM.pushAny(fmt.Sprintf("table:%v", val))
			defer pVal.free()
		case *LnsMap:
			pVal := luaVM.pushAny(fmt.Sprintf("table:%v", val))
			defer pVal.free()
		case LnsInt:
			luaVM.pushAny(val)
		case LnsReal:
			luaVM.pushAny(val)
		case bool:
			luaVM.pushAny(val)
		case string:
			luaVM.pushAny(val)
		case *Lns_luaValue:
			luaVM.pushAny(val)
		default:
			if Lns_IsNil(val) {
				luaVM.pushAny(val)
			} else {
				pVal := luaVM.pushAny(fmt.Sprintf("table:%v", val))
				defer pVal.free()
			}
		}
	}

	return callInfo.call(luaVM, 1+len(ddd), 1)[0].(string)
}

func (luaVM *Lns_luaVM) String_replace(txt, src, dst string) string {
	return strings.ReplaceAll(txt, src, dst)
}

func (luaVM *Lns_luaVM) String_gsub(
	txt string, src string, dst string) (string, LnsInt) {

	callInfo := luaVM.string_call_setup(lns_c_ptr_gsub, txt)

	pSrc := luaVM.pushStr(src)
	defer pSrc.free()
	pDst := luaVM.pushStr(dst)
	defer pDst.free()

	ret := callInfo.call(luaVM, 3, 2)
	return ret[0].(string), LnsInt(ret[1].(int))
}

type RegexpCache struct {
	cache *lru.Cache
}

func newRegexpCache(num LnsInt) *RegexpCache {
	return &RegexpCache{lru.New(num)}
}

func (self *RegexpCache) createRegexp(src string) *regexp.Regexp {

	if exp, ok := self.cache.Get(src); ok {
		return exp.(*regexp.Regexp)
	}

	offset := 0
	remain := len(src)

	processQuote := func(inClass bool) string {
		remain--
		offset++
		if remain <= 0 {
			return ""
		}
		cls := ""
		switch src[offset] {
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
			return fmt.Sprintf("\\x%02X", src[offset])
		}
		if inClass {
			return cls
		}
		return fmt.Sprintf("[%s]", cls)
	}

	reg := ""
	if remain >= 2 && src[offset] == '^' {
		reg += "^"

		offset++
		remain--
	}

	inClass := false
	firstFlag := true
	for remain > 0 {
		switch src[offset] {
		case '%':
			if quote := processQuote(inClass); quote == "" {
				return nil
			} else {
				reg += quote
			}
		case '[':
			if inClass {
				return nil
			}
			inClass = true
			reg += src[offset : offset+1]
		case ']':
			if !inClass {
				return nil
			}
			inClass = false
			reg += src[offset : offset+1]
		case '^':
			reg += src[offset : offset+1]
		case '$':
			reg += src[offset : offset+1]
		case '*':
			reg += src[offset : offset+1]
		case '+':
			reg += src[offset : offset+1]
		case '-':
			if inClass || firstFlag {
				reg += src[offset : offset+1]
			} else {
				return nil
			}
		case '|':
			reg += fmt.Sprintf("\\x%02X", src[offset])
		case '{':
			reg += fmt.Sprintf("\\x%02X", src[offset])
		case '}':
			reg += fmt.Sprintf("\\x%02X", src[offset])
		case '(':
			return nil
		case '\\':
			reg += fmt.Sprintf("\\x%02X", src[offset])
		case ':':
			reg += fmt.Sprintf("\\x%02X", src[offset])
		default:
			reg += src[offset : offset+1]
		}

		if firstFlag {
			firstFlag = false
		}
		remain--
		offset++
	}

	re := regexp.MustCompile(reg)
	self.cache.Add(src, re)

	return re
}

var goFIndDummyRet = []LnsAny{}

func (self *RegexpCache) goFind(txt string, src string, index LnsInt) (bool, []LnsAny) {

	dummyRet := goFIndDummyRet

	re := self.createRegexp(src)
	if re == nil {
		return false, dummyRet
	}

	loc := re.FindStringIndex(txt)
	if loc == nil {
		return true, dummyRet
	}

	if index <= 0 {
		index = 1
	}
	return true, []LnsAny{index + loc[0], index + loc[1] - 1}
}

// goFind の結果と lua ラインタイムの find の結果を比べるかどうか。
// デバッグ用途。
var checkFindResult = false

func (luaVM *Lns_luaVM) String_find(
	txt string, src string, index LnsAny, plain LnsAny) []LnsAny {

	target := txt
	offset := 1
	if work, ok := index.(LnsInt); ok {
		offset = string_index(work, len(txt))
		if offset > 0 {
			if offset > len(txt) {
				return goFIndDummyRet
			}
			target = txt[offset-1:]
		}
	}

	if plain == true {
		findIndex := strings.Index(target, src)
		if findIndex < 0 {
			return goFIndDummyRet
		}
		return []LnsAny{offset + findIndex, offset + findIndex + len(src) - 1}
	}
	var goResult []LnsAny
	if ok, result := luaVM.regexCache.goFind(target, src, offset); ok {
		if !checkFindResult {
			return result
		}
		goResult = result
	}
	if validRuntimeDebugLog {
		print(fmt.Sprintf("log: fallback -- string.find()  %s, %s, %v\n", txt, src, index))
	}

	// goFind で処理できなかった場合は、 lua のランタイムで処理する

	callInfo := luaVM.string_call_setup(lns_c_ptr_find, txt)

	pSrc := luaVM.pushStr(src)
	defer pSrc.free()
	luaVM.pushAny(index) // int なので defer 不要
	luaVM.pushAny(plain) // bool なので defer 不要

	ret := callInfo.call(luaVM, 4, 2)

	if checkFindResult {
		if len(ret) != len(goResult) ||
			ret[0] != goResult[0] ||
			(len(ret) > 1 && ret[1] != goResult[1]) {
			mess := fmt.Sprintf(
				"hoge: %s %s %d %d %d %s %s\n", txt, src, index,
				len(ret), len(goResult), ret[0], goResult[0])
			if len(ret) > 1 && len(goResult) > 1 {
				mess = fmt.Sprintf("%s %s %s\n", mess, ret[1], goResult[1])
			}
			panic(mess)
		}
	}

	return ret
}

func (luaVM *Lns_luaVM) String_byte(
	txt string, from LnsAny, to LnsAny) []LnsAny {

	callInfo := luaVM.string_call_setup(lns_c_ptr_byte, txt)

	luaVM.pushAny(from) // int なので defer 不要
	luaVM.pushAny(to)   // int なので defer 不要

	ret := callInfo.call(luaVM, 3, cLUA_MULTRET)
	return ret
}

func (luaVM *Lns_luaVM) String_rep(txt string, num LnsInt) string {
	var builder strings.Builder
	for count := 1; count <= num; count++ {
		builder.WriteString(txt)
	}
	return builder.String()
}

func (luaVM *Lns_luaVM) String_sub(txt string, from LnsInt, to LnsAny) string {
	// callInfo := luaVM.string_call_setup( lns_c_ptr_sub, txt )
	// luaVM.pushAny( from ) // int なので defer 不要
	// luaVM.pushAny( to ) // int なので defer 不要
	// return callInfo.call( luaVM, 3, 1 )[0].(string)

	from = string_index(from, len(txt))
	if from < 1 {
		from = 1
	}
	end := len(txt)
	if work, ok := to.(LnsInt); ok {
		end = string_index(work, len(txt))
	}
	if from <= end {
		if end > len(txt) {
			end = len(txt)
		}
		return txt[from-1 : end]
	}
	return ""
}

func (luaVM *Lns_luaVM) String_lower(txt string) string {
	return strings.ToLower(txt)
}

func (luaVM *Lns_luaVM) String_upper(txt string) string {
	return strings.ToUpper(txt)
}

func (luaVM *Lns_luaVM) String_reverse(txt string) string {
	callInfo := luaVM.string_call_setup(lns_c_ptr_reverse, txt)
	return callInfo.call(luaVM, 1, 1)[0].(string)
}

func (luaVM *Lns_luaVM) String_gmatch(txt string, pat string) (LnsAny, LnsAny, LnsAny) {

	callInfo := luaVM.string_call_setup(lns_c_ptr_gmatch, txt)

	pSrc := luaVM.pushStr(pat)
	defer pSrc.free()

	ret := callInfo.call(luaVM, 2, 3)
	return Lns_getFromMulti(ret, 0), Lns_getFromMulti(ret, 1), Lns_getFromMulti(ret, 2)
}

func (luaVM *Lns_luaVM) String_dump(form *Lns_luaValue, flag LnsAny) string {

	top := luaVM.string_static_call_setup(lns_c_ptr_dump)

	form.core.pushValFromGlobalValMap()
	luaVM.pushAny(flag) // bool なので defer 不要

	ret := luaVM.lua_call(top, 2, 1)
	return ret[0].(string)
}

func Lns_stringJoin( listObj, sepNil, startNil, endNil LnsAny ) string {
    if list, ok := listObj.(*LnsList); ok {
        return Lns_string_Join( list, sepNil, startNil, endNil )
    }
    if list, ok := listObj.(*LnsList); ok {
        return Lns_string_Join( list, sepNil, startNil, endNil )
    }
    panic( fmt.Sprintf( "illegal type -- %v", listObj ) )
}

func Lns_string_Join( list *LnsList, sepNil, startNil, endNil LnsAny ) string {
    var sep string
    if Lns_IsNil( sepNil ) {
        sep = ""
    } else {
        sep = sepNil.(string)
    }
    start := 0
    if !Lns_IsNil( startNil ) {
        start = startNil.(LnsInt) - 1
    }
    end := len( list.Items ) - 1
    if !Lns_IsNil( endNil ) {
        end = endNil.(LnsInt) - 1
    }
    if end < start {
        return ""
    }
    if end == start {
        return list.Items[ start ].(string)
    }
    

	var builder strings.Builder
	for index, txtObj := range (list.Items) {
        txt := txtObj.(string)
        if start <= index {
            builder.WriteString(txt)
            if index >= end {
                break
            }
            if len( sep ) != 0 {
                builder.WriteString(sep)
            }
        }
	}
	return builder.String()
}

func Lns_string__Join( list *LnsList2_[string], sepNil, startNil, endNil LnsAny ) string {
    var sep string
    if Lns_IsNil( sepNil ) {
        sep = ""
    } else {
        sep = sepNil.(string)
    }
    start := 0
    if !Lns_IsNil( startNil ) {
        start = startNil.(LnsInt) - 1
    }
    end := len( list.Items ) - 1
    if !Lns_IsNil( endNil ) {
        end = endNil.(LnsInt) - 1
    }
    if end < start {
        return ""
    }
    if end == start {
        return list.Items[ start ]
    }

    slice := list.Items[start:end+1]
    return strings.Join( slice, sep )
}