//go:build cgo && !gopherlua && !azuregolua
// +build cgo,!gopherlua,!azuregolua

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

// #include <string.h>
// #include <stdlib.h>
// #cgo pkg-config: lua-5.3
// #include <lauxlib.h>
// #include <lualib.h>
//
// extern int LnsPreload( lua_State * pVM );
// // pName のパッケージをロードする関数に LnsPreload をセットする
// static void registPreload( lua_State * pVM, char * pName ) {
//    int top = lua_gettop( pVM );
//    // package.preload を push
//    lua_getglobal( pVM, "package");
//    lua_getfield( pVM, -1, "preload" );
//    // LnsPreload を push
//    lua_pushcfunction( pVM, LnsPreload );
//    // package.preload[ pName ] に LnsPreload をセット
//    lua_setfield( pVM, -2, pName );
//    lua_settop( pVM, top );
// }
import "C"
import (
	"fmt"
	"math"
	"unsafe"
)

//import "sync"

type lua_rawstr = *C.char
type lua_state = *C.lua_State

type lua_int = C.longlong
type lua_num = C.double
type lua_bool = C.int

func Lns_toRawStr(str string) lua_rawstr {
	return C.CString(str)
}

func Lns_freeRawStr(pStr lua_rawstr) {
	C.free(unsafe.Pointer(pStr))
}

func Lns_zeroRawStr() lua_rawstr {
	return nil
}

type lnsSrcInfo struct {
	codeC lua_rawstr
	len   int
}

var lnsSrcMap map[string]*lnsSrcInfo

func AddlnsSrcInfo(key string, code []byte) {
	lnsSrcMap[key] = &lnsSrcInfo{C.CString(string(code)), len(code)}
}

// lua の package.preload に登録される関数。
// lua の第一引数にロードするモジュール名が渡される。
// 戻り値として、ロード後の値を push する。
//
//export LnsPreload
func LnsPreload(vm *C.lua_State) C.int {
	mod := lua_tolstring(vm, 1)

	srcInfo := lnsSrcMap[mod]

	pModule := C.CString(mod)
	defer C.free(unsafe.Pointer(pModule))
	pMode := C.CString("bt")
	defer C.free(unsafe.Pointer(pMode))

	if code := C.luaL_loadbufferx(
		vm, srcInfo.codeC, C.ulong(srcInfo.len), pModule, pMode); code != C.LUA_OK {
		panic(lua_tolstring(vm, -1))
	}
	if code := C.lua_pcallk(vm, 0, 1, 0, 0, nil); code != C.LUA_OK {
		panic(lua_tolstring(vm, -1))
	}
	return 1
}
func Lns_initPreload(vm *C.lua_State) {
	for key := range lnsSrcMap {
		pKey := C.CString(key)
		defer C.free(unsafe.Pointer(pKey))
		C.registPreload(vm, pKey)
	}
}

//func Lns_initPreload( vm *C.lua_State ) {}

var cLUA_MULTRET int
var cLUA_TBOOLEAN int
var cLUA_TFUNCTION int
var cLUA_TTABLE int
var cLUA_TNIL int
var cLUA_TNUMBER int
var cLUA_TSTRING int

func init() {
	lnsSrcMap = map[string]*lnsSrcInfo{}
	cLUA_MULTRET = int(C.LUA_MULTRET)
	cLUA_TBOOLEAN = int(C.LUA_TBOOLEAN)
	cLUA_TFUNCTION = int(C.LUA_TFUNCTION)
	cLUA_TTABLE = int(C.LUA_TTABLE)
	cLUA_TNIL = int(C.LUA_TNIL)
	cLUA_TNUMBER = int(C.LUA_TNUMBER)
	cLUA_TSTRING = int(C.LUA_TSTRING)
}

func Lns_getLoadFuncName() string {
	return "load"
}

func depend_getLuaVersion() string {
	return "5.3"
}
func depend_setup(callback func(ver LnsInt)) {
	callback(53)
}

func lns_ToStringFromReal(val LnsReal) string {
	if digit, frac := math.Modf(val); frac == 0 {
		return fmt.Sprintf("%g.0", digit)
	}
	return fmt.Sprintf("%g", val)
}

// luaL api ======================
func LuaL_loadstring(vm *C.lua_State, txt string) int {
	pTxt := C.CString(txt)
	defer C.free(unsafe.Pointer(pTxt))
	return int(C.luaL_loadstring(vm, pTxt))
}
func LuaL_newstate(stackSize int) *C.lua_State {
	vm := C.luaL_newstate()
	C.lua_checkstack(vm, C.int(stackSize))
	return vm
}
func LuaL_openlibs(vm *C.lua_State) {
	C.luaL_openlibs(vm)
}

// lua api ======================

func lua_callk(vm *C.lua_State, argNum int, retNum int) {
	C.lua_callk(vm, C.int(argNum), C.int(retNum), C.long(0), nil)
}
func lua_close(vm *C.lua_State) {
	C.lua_close(vm)
}
func lua_replace(vm *C.lua_State, index int) {
	C.lua_copy(vm, C.int(-1), C.int(index))
	lua_pop(vm, 1)
}
func lua_createtable(vm *C.lua_State) {
	C.lua_createtable(vm, C.int(0), C.int(0))
}

//	func lua_geti(vm *C.lua_State, index int, fieldPos int ) int {
//	    return int(C.lua_geti(vm, C.int(index), C.longlong(fieldPos) ))
//	}
func lua_gettable(vm *C.lua_State, index int) {
	C.lua_gettable(vm, C.int(index))
}
func lua_getfield(vm *C.lua_State, index int, pSym lua_rawstr) {
	C.lua_getfield(vm, C.int(index), pSym)
}
func lua_getglobal(vm *C.lua_State, pSym lua_rawstr) {
	C.lua_getglobal(vm, pSym)
}
func lua_gettop(vm *C.lua_State) int {
	return int(C.lua_gettop(vm))
}
func lua_pcallk(vm *C.lua_State, argNum int, retNum int) error {
	ret := int(C.lua_pcallk(vm, C.int(argNum), C.int(retNum), 0, 0, nil))
	if ret != C.LUA_OK {
		return fmt.Errorf("%s", lua_tolstring(vm, -1))
	}
	return nil
}
func lua_pushboolean(vm *C.lua_State, flag bool) {
	if flag {
		C.lua_pushboolean(vm, lua_bool(1))
	} else {
		C.lua_pushboolean(vm, lua_bool(0))
	}
}
func lua_pushinteger(vm *C.lua_State, val LnsInt) {
	C.lua_pushinteger(vm, lua_int(val))
}
func lua_pushnil(vm *C.lua_State) {
	C.lua_pushnil(vm)
}
func lua_pushnumber(vm *C.lua_State, val LnsReal) {
	C.lua_pushnumber(vm, lua_num(val))
}
func lua_pushlstring(vm *C.lua_State, pSym lua_rawstr, len int) {
	C.lua_pushlstring(vm, pSym, C.ulong(len))
}
func lua_pushvalue(vm *C.lua_State, index int) {
	C.lua_pushvalue(vm, C.int(index))
}
func lua_setfield(vm *C.lua_State, index int, pSym lua_rawstr) {
	C.lua_setfield(vm, C.int(index), pSym)
}
func lua_setglobal(vm *C.lua_State, pSym lua_rawstr) {
	C.lua_setglobal(vm, pSym)
}
func lua_settop(vm *C.lua_State, index int) {
	C.lua_settop(vm, C.int(index))
}
func lua_pop(vm *C.lua_State, num int) {
	C.lua_settop(vm, C.int(-(num)-1))
}
func lua_toboolean(vm *C.lua_State, index int) bool {
	return C.lua_toboolean(vm, C.int(index)) != 0
}
func lua_tointegerx(vm *C.lua_State, index int) LnsInt {
	return LnsInt(C.lua_tointegerx(vm, C.int(index), nil))
}
func lua_tolstring(vm *C.lua_State, index int) string {
	var len C.ulong
	pstr := C.lua_tolstring(vm, C.int(index), &len)
	str := C.GoStringN(pstr, C.int(len))
	return str
}
func lua_tonumberx(vm *C.lua_State, index int) LnsReal {
	return LnsReal(C.lua_tonumberx(vm, C.int(index), nil))
}
func lua_type(vm *C.lua_State, index int) int {
	return int(C.lua_type(vm, C.int(index)))
}
func lua_next(vm *C.lua_State, index int) int {
	return int(C.lua_next(vm, C.int(index)))
}
func lua_len(vm *C.lua_State, index int) LnsInt {
	C.lua_len(vm, C.int(index))
	len := lua_tointegerx(vm, -1)
	lua_pop(vm, 1)
	return len
}
