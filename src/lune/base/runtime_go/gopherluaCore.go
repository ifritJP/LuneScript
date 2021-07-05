// +build gopherlua

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

	"github.com/yuin/gopher-lua"
)

//import "sync"

type lua_rawstr = string
type lua_state = *lua.LState

func Lns_toRawStr(str string) lua_rawstr {
	return str
}

func Lns_freeRawStr(pStr lua_rawstr) {
}

const zeroRawStr = ""

func Lns_zeroRawStr() lua_rawstr {
	return zeroRawStr
}

type lnsSrcInfo struct {
	codeC lua_rawstr
	len   int
}

var lnsSrcMap map[string]*lnsSrcInfo

func AddlnsSrcInfo(key string, code []byte) {
	lnsSrcMap[key] = &lnsSrcInfo{string(code), len(code)}
}

var readyBind bool

func Lns_initPreload(vm lua_state) {
	if !readyBind {
		init_bind()
		readyBind = true
	}
	for key := range lnsSrcMap {
		code := lnsSrcMap[key]
		vm.PreloadModule(key, func(vms lua_state) int {
			if err := LuaL_loadstring(vms, code.codeC); err != nil {
				panic(err)
			}
			lua_callk(vms, 0, cLUA_MULTRET)
			return 1
		})
	}
}

var cLUA_MULTRET int
var cLUA_TBOOLEAN int
var cLUA_TFUNCTION int
var cLUA_TTABLE int
var cLUA_TNIL int
var cLUA_TNUMBER int
var cLUA_TSTRING int

func init() {
	readyBind = false
	lnsSrcMap = map[string]*lnsSrcInfo{}
	cLUA_MULTRET = int(lua.MultRet)
	cLUA_TBOOLEAN = int(lua.LTBool)
	cLUA_TFUNCTION = int(lua.LTFunction)
	cLUA_TTABLE = int(lua.LTTable)
	cLUA_TNIL = int(lua.LTNil)
	cLUA_TNUMBER = int(lua.LTNumber)
	cLUA_TSTRING = int(lua.LTString)
}

func Lns_getLoadFuncName() string {
	return "loadstring"
}

func depend_getLuaVersion() string {
	return "5.1"
}

func depend_setup(callback func(ver LnsInt)) {
	callback(51)
}

func lns_ToStringFromReal(val LnsReal) string {
	return fmt.Sprintf("%g", val)
}

func lua_popGet(vm lua_state) lua.LValue {
	val := vm.Get(-1)
	vm.Pop(1)
	return val
}

// luaL api ======================
func LuaL_loadstring(vm lua_state, txt string) error {
	if fn, err := vm.LoadString(txt); err != nil {
		return err
	} else {
		vm.Push(fn)
	}
	return nil
}
func LuaL_newstate(stackSize int) lua_state {
	return lua.NewState(lua.Options{
		CallStackSize: stackSize, RegistrySize: stackSize, RegistryMaxSize: stackSize})
}
func LuaL_openlibs(vm lua_state) {
	vm.OpenLibs()
}

// lua api ======================

func lua_callk(vm lua_state, argNum int, retNum int) {
	vm.Call(argNum, retNum)
}
func lua_close(vm lua_state) {
	vm.Close()
}
func lua_replace(vm lua_state, index int) {
	vm.Replace(index, lua_popGet(vm))
}
func lua_createtable(vm lua_state) {
	vm.Push(vm.NewTable())
}
func lua_gettable(vm lua_state, index int) {
	obj := vm.Get(index)
	key := lua_popGet(vm)
	vm.Push(vm.GetTable(obj, key))
}
func lua_getfield(vm lua_state, index int, pSym lua_rawstr) {
	vm.Push(vm.GetField(vm.Get(index), pSym))
}
func lua_getglobal(vm lua_state, pSym lua_rawstr) {
	vm.Push(vm.GetGlobal(pSym))
}
func lua_gettop(vm lua_state) int {
	return vm.GetTop()
}
func lua_pcallk(vm lua_state, argNum int, retNum int) error {
	return vm.PCall(argNum, retNum, nil)
}
func lua_pushboolean(vm lua_state, flag bool) {
	if flag {
		vm.Push(lua.LTrue)
	} else {
		vm.Push(lua.LFalse)
	}
}
func lua_pushinteger(vm lua_state, val LnsInt) {
	vm.Push(lua.LNumber(val))
}
func lua_pushnil(vm lua_state) {
	vm.Push(lua.LNil)
}
func lua_pushnumber(vm lua_state, val LnsReal) {
	vm.Push(lua.LNumber(val))
}
func lua_pushlstring(vm lua_state, pSym lua_rawstr, len int) {
	vm.Push(lua.LString(pSym))
}
func lua_pushvalue(vm lua_state, index int) {
	vm.Push(vm.Get(index))
}
func lua_setfield(vm lua_state, index int, pSym lua_rawstr) {
	obj := vm.Get(index)
	val := lua_popGet(vm)
	vm.SetField(obj, pSym, val)
}
func lua_setglobal(vm lua_state, pSym lua_rawstr) {
	val := lua_popGet(vm)
	vm.SetGlobal(pSym, val)
}
func lua_settop(vm lua_state, index int) {
	vm.SetTop(index)
}
func lua_pop(vm lua_state, num int) {
	vm.Pop(num)
}
func lua_toboolean(vm lua_state, index int) bool {
	return vm.ToBool(index)
}
func lua_tointegerx(vm lua_state, index int) LnsInt {
	return vm.ToInt(index)
}
func lua_tolstring(vm lua_state, index int) string {
	return vm.ToString(index)
}
func lua_tonumberx(vm lua_state, index int) LnsReal {
	return LnsReal(vm.ToNumber(index))
}
func lua_type(vm lua_state, index int) int {
	return int(vm.Get(index).Type())
}
func lua_next(vm lua_state, index int) int {
	tbl := vm.ToTable(index)
	key := lua_popGet(vm)
	nextKey, nextVal := tbl.Next(key)
	if nextKey == nil {
		return 0
	}
	vm.Push(nextKey)
	vm.Push(nextVal)
	return 1
}
func lua_len(vm lua_state, index int) LnsInt {
	return vm.ObjLen(vm.Get(index))
}
