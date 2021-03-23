// +build azuregolua

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

import "github.com/Azure/golua/lua"
import "github.com/Azure/golua/std"

//import "sync"
import "fmt"
import "math"


type lua_rawstr = string
type lua_state = *lua.State


func Lns_toRawStr( str string ) lua_rawstr {
    return str
}

func Lns_freeRawStr( pStr lua_rawstr ) {
}

const zeroRawStr = ""
func Lns_zeroRawStr() lua_rawstr {
    return zeroRawStr
}

type lnsSrcInfo struct {
  codeC lua_rawstr
  len int
}
var lnsSrcMap map[string] *lnsSrcInfo

func AddlnsSrcInfo( key string, code []byte ) {
    lnsSrcMap[ key ] = &lnsSrcInfo{ string( code ), len( code ) }
}

// lua の package.preload に登録される関数。
// lua の第一引数にロードするモジュール名が渡される。
// 戻り値として、ロード後の値を push する。
func LnsPreload( vm lua_state ) int {
    lua_pushnil( vm )
    return 1
    
    // mod := lua_tolstring( vm, 1 );

    // if srcInfo, ok := lnsSrcMap[ mod ]; !ok {
    //     lua_pushnil( vm )
    //     return 1
    // }

    // if code := lua.luaL_loadbufferx(
    //     vm, srcInfo.codeC, srcInfo.len, mod, "bt" ); code != lua.LUA_OK {
    //     panic( lua_tolstring( vm, -1 ))
    // }
    // if code := C.lua_pcallk( vm, 0, 1, 0, 0, nil ); code != C.LUA_OK {
    //     panic( lua_tolstring( vm, -1 ));
    // }
    // return 1;
}
func Lns_initPreload( vm lua_state ) {
    // for key := range( lnsSrcMap ) {
    //     lua.registPreload( vm, key )
    // }
}

var cLUA_MULTRET int
var cLUA_TBOOLEAN int
var cLUA_TFUNCTION int
var cLUA_TTABLE int
var cLUA_TNIL int
var cLUA_TNUMBER int
var cLUA_TSTRING int

func init() {
    lnsSrcMap = map[string] *lnsSrcInfo{}
    cLUA_MULTRET = int(lua.MultRets)
    cLUA_TBOOLEAN = int(lua.BoolType)
    cLUA_TFUNCTION = int(lua.FuncType)
    cLUA_TTABLE = int(lua.TableType)
    cLUA_TNIL = int(lua.NilType)
    cLUA_TNUMBER = int(lua.NumberType)
    cLUA_TSTRING = int(lua.StringType)
}

func Lns_getLoadFuncName() string {
    return "load"
}

func Depend_getLuaVersion() string {
    return "5.3"
}
type Depend_UpdateVer func ( ver LnsInt )
func Depend_setup( callback Depend_UpdateVer) {
    callback( 53 );
}

func lns_ToStringFromRead( val LnsReal ) string {
    if digit, frac := math.Modf( val ); frac == 0 {
        return fmt.Sprintf( "%g.0", digit )
    }
    return fmt.Sprintf( "%g", val )
}

func lua_popGet( vm lua_state ) lua.Value {
    return vm.Pop()
}

// luaL api ======================
func LuaL_loadstring( vm lua_state, txt string ) error {
    return vm.LoadText( txt )
}
func LuaL_newstate( stackSize int ) lua_state {
    vm := lua.NewState()
    vm.CheckStack( stackSize )
    return vm
}
func LuaL_openlibs( vm lua_state ) {
    std.Open( vm )
}

// lua api ======================

func lua_callk( vm lua_state, argNum int, retNum int ) {
    vm.Call( argNum, retNum )
}
func lua_close(vm lua_state) {
    vm.Close()
}
func lua_replace(vm lua_state, index int ) {
    vm.Replace( index )
}
func lua_createtable(vm lua_state) {
    vm.NewTable()
}
func lua_gettable(vm lua_state, index int ) {
    vm.GetTable( index )
}
func lua_getfield(vm lua_state, index int, pSym lua_rawstr ) {
    vm.GetField( index, pSym )
}
func lua_getglobal(vm lua_state, pSym lua_rawstr) {
    vm.GetGlobal( pSym )
}
func lua_gettop(vm lua_state) int {
    return vm.Top()
}
func lua_pcallk(vm lua_state, argNum int, retNum int ) error {
    return vm.PCall( argNum, retNum, 0 )
}
func lua_pushboolean(vm lua_state, flag bool ) {
    vm.Push( flag )
}
func lua_pushinteger(vm lua_state, val LnsInt ) {
    vm.Push( val )
}
func lua_pushnil(vm lua_state) {
    vm.Push( nil )
}
func lua_pushnumber(vm lua_state, val LnsReal ) {
    vm.Push( val )
}
func lua_pushlstring(vm lua_state, pSym lua_rawstr, len int ) {
    vm.Push( pSym )
}
func lua_pushvalue(vm lua_state, index int ) {
    vm.PushIndex( index )
}
func lua_setfield(vm lua_state, index int, pSym lua_rawstr ) {
    vm.SetField( index, pSym )
}
func lua_setglobal(vm lua_state, pSym lua_rawstr ) {
    vm.SetGlobal( pSym )
}
func lua_settop(vm lua_state, index int) {
    vm.SetTop( index )
}
func lua_pop(vm lua_state, num int) {
    vm.PopN( num )
}
func lua_toboolean(vm lua_state, index int) bool {
    return vm.ToBool( index )
}
func lua_tointegerx(vm lua_state, index int) LnsInt {
    return LnsInt(vm.ToInt( index ))
}
func lua_tolstring(vm lua_state, index int ) string {
    return vm.ToString( index )
}
func lua_tonumberx(vm lua_state, index int) LnsReal {
    return LnsReal(vm.ToNumber( index ))
}
func lua_type(vm lua_state, index int) int {
    val := vm.CheckAny(index)
    return int(val.Type())
}
func lua_next(vm lua_state, index int) int {
    if vm.Next( index ) {
        return 1
    }
    return 0
}
func lua_len( vm lua_state, index int) LnsInt {
    val := vm.CheckAny(index)
    switch val.Type() {
    case lua.StringType:
        return len( val.String() )
    case lua.TableType:
        return vm.ToTable( index ).Length()
    }
    return 0
}
