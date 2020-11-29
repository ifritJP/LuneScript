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

// +build cgo

package runtimelns

// #include <string.h>
// #include <stdlib.h>
// #cgo pkg-config: lua-5.3
// #include <lauxlib.h>
// #include <lualib.h>
import "C"
import "unsafe"
//import "sync"

type lua_int = C.longlong
type lua_num = C.double
type lua_bool = C.int

var cLUA_MULTRET int
var cLUA_TBOOLEAN int
var cLUA_TFUNCTION int
var cLUA_TTABLE int
var cLUA_TNIL int
var cLUA_TNUMBER int
var cLUA_TSTRING int
var cLUA_OK int

func init() {
    cLUA_MULTRET = int(C.LUA_MULTRET)
    cLUA_TBOOLEAN = int(C.LUA_TBOOLEAN)
    cLUA_TFUNCTION = int(C.LUA_TFUNCTION)
    cLUA_TTABLE = int(C.LUA_TTABLE)
    cLUA_TNIL = int(C.LUA_TNIL)
    cLUA_TNUMBER = int(C.LUA_TNUMBER)
    cLUA_TSTRING = int(C.LUA_TSTRING)
    cLUA_OK = int(C.LUA_OK)
}

type Lns_luaVM struct {
    vm *C.lua_State

    lns_luvValueCoreMap map[*Lns_luaValueCore]bool
    lns_luaValChan chan *Lns_luaValueCore

    lns_luvValueCoreFreeList []*Lns_luaValueCoreList
    lns_hasLuvValueCoreFree bool

    regexCache *RegexpCache
}

// luaL api ======================
func luaL_loadstring( vm *C.lua_State, txt string ) int {
    pTxt := C.CString( txt )
    defer C.free( unsafe.Pointer( pTxt ) )
    return int(C.luaL_loadstring( vm, pTxt ))
}
func luaL_newstate() *C.lua_State {
    return C.luaL_newstate()
}
func luaL_openlibs( vm *C.lua_State ) {
    C.luaL_openlibs( vm )
}

// lua api ======================

func lua_checkstack( vm *C.lua_State, size int ) int {
    return int( C.lua_checkstack( vm, C.int( size ) ) )
}

func lua_callk( vm *C.lua_State, argNum int, retNum int ) {
    C.lua_callk( vm, C.int(argNum), C.int(retNum), C.long(0), nil )
}
func lua_close(vm *C.lua_State) {
    C.lua_close( vm )
}
func lua_copy(vm *C.lua_State, from int, to int ) {
    C.lua_copy(vm, C.int(from), C.int(to))
}
func lua_createtable(vm *C.lua_State) {
    C.lua_createtable(vm,C.int(0), C.int(0))
}
func lua_geti(vm *C.lua_State, index int, fieldPos int ) int {
    return int(C.lua_geti(vm, C.int(index), C.longlong(fieldPos) ))
}
func lua_gettable(vm *C.lua_State, index int ) int {
    return int(C.lua_gettable(vm, C.int(index) ))
}
func lua_getfield(vm *C.lua_State, index int, pSym *C.char ) int {
    return int(C.lua_getfield(vm, C.int(index), pSym ))
}
func lua_getglobal(vm *C.lua_State, pSym *C.char) int {
    return int(C.lua_getglobal(vm, pSym))
}
func lua_gettop(vm *C.lua_State) int {
    return int( C.lua_gettop(vm) )
}
func lua_pcallk(vm *C.lua_State, argNum int, retNum int ) int {
    return int(C.lua_pcallk(vm, C.int( argNum ), C.int( retNum ), 0, 0, nil ))
}
func lua_pushboolean(vm *C.lua_State, flag bool ) {
    if flag {
        C.lua_pushboolean( vm, lua_bool( 1 ) )
    } else {
        C.lua_pushboolean( vm, lua_bool( 0 ) )
    }
}
func lua_pushinteger(vm *C.lua_State, val LnsInt ) {
    C.lua_pushinteger(vm, lua_int(val))
}
func lua_pushnil(vm *C.lua_State) {
    C.lua_pushnil(vm)
}
func lua_pushnumber(vm *C.lua_State, val LnsReal ) {
    C.lua_pushnumber(vm, lua_num( val ) )
}
func lua_pushlstring(vm *C.lua_State, pSym *C.char, len int ) {
    C.lua_pushlstring(vm, pSym, C.ulong( len ) )
}
func lua_pushvalue(vm *C.lua_State, index int ) {
    C.lua_pushvalue(vm, C.int( index ))
}
func lua_setfield(vm *C.lua_State, index int, pSym *C.char ) {
    C.lua_setfield(vm, C.int( index ), pSym)
}
func lua_setglobal(vm *C.lua_State, pSym *C.char ) {
    C.lua_setglobal(vm, pSym)
}
func lua_settop(vm *C.lua_State, index int) {
    C.lua_settop(vm, C.int( index ))
}
func lua_pop(vm *C.lua_State, num int) {
    C.lua_settop(vm, C.int( -(num)-1 ))
}
func lua_toboolean(vm *C.lua_State, index int) bool {
    return C.lua_toboolean(vm, C.int( index ) ) != 0
}
func lua_tointegerx(vm *C.lua_State, index int) LnsInt {
    return LnsInt( C.lua_tointegerx(vm, C.int( index ), nil ) )
}
func lua_tolstring(vm *C.lua_State, index int ) string {
    var len C.ulong
    pstr := C.lua_tolstring(vm, C.int( index ), &len )
    str := C.GoStringN( pstr, C.int( len ) )
    return str
}
func lua_tonumberx(vm *C.lua_State, index int) LnsReal {
    return LnsReal( C.lua_tonumberx(vm, C.int( index ), nil ) )
}
func lua_type(vm *C.lua_State, index int) int {
    return int(C.lua_type(vm, C.int( index )))
}
func lua_next(vm *C.lua_State, index int) int {
    return int(C.lua_next(vm, C.int( index )))
}
func lua_len( vm *C.lua_State, index int) {
    C.lua_len(vm, C.int( index ))
}
