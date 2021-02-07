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

// #include <stdlib.h>
import "C"

import "unsafe"
import "fmt"
import "runtime"
import "strings"
import "sync"

var lns_globalValSym *C.char
const lns_globalValMapName = "lns_globalValMap"

var lns_luvValueCoreFreeListMutex sync.Mutex 

type Lns_luaValueCore struct {
    // lua VM の lns_globalValMap に格納しているシンボル名
    sym *C.char
    luaVM *Lns_luaVM
    typeId int
}

type Lns_luaValueCoreList struct {
    list []*Lns_luaValueCore
}

type Lns_luaValue struct {
    core *Lns_luaValueCore
}

const lns_luaValueCoreNum = 10

/**
symbol を globalVal から除外する
*/
func (luaValue *Lns_luaValue) free() {
    luaValue.core.luaVM.lns_luaValChan<- luaValue.core
}

/**
symbol を globalVal から除外する
*/
func (core *Lns_luaValueCore) free() {
    delete( core.luaVM.lns_luvValueCoreMap, core )

    top := lua_gettop( core.luaVM.vm )
    lua_pushnil( core.luaVM.vm )
    core.setValToGlobalValMap( -1 )
    lua_settop( core.luaVM.vm, top )
    


    C.free( unsafe.Pointer( core.sym ) )
    core.sym = nil
}

func (core *Lns_luaValueCore) getKey() string {
    return fmt.Sprintf( "%p", core )
}

func init() {
    lns_globalValSym = C.CString( lns_globalValMapName )
    lns_defaultPushedVal = &lns_pushedVal{}
}

func lns_luaValChanProc( luaVM *Lns_luaVM ) {
    list := make([]*Lns_luaValueCore,lns_luaValueCoreNum)
    count := 0
    for {
        core := <-luaVM.lns_luaValChan

        list[ count ] = core
        count++
        if count == lns_luaValueCoreNum {
            count = 0
            freeList := &Lns_luaValueCoreList{ list }
            
            lns_luvValueCoreFreeListMutex.Lock()
            luaVM.lns_hasLuvValueCoreFree = true
            luaVM.lns_luvValueCoreFreeList =
                append( luaVM.lns_luvValueCoreFreeList, freeList )
            lns_luvValueCoreFreeListMutex.Unlock()
            
            list = make([]*Lns_luaValueCore,lns_luaValueCoreNum)
        }
    }
}

func createVM() *Lns_luaVM {
    luaVM := &Lns_luaVM{}
    luaVM.vm = luaL_newstate()
    luaL_openlibs( luaVM.vm )
    luaVM.lns_luvValueCoreMap = map[*Lns_luaValueCore]bool{}
    luaVM.lns_luaValChan = make(chan *Lns_luaValueCore, 1000)
    luaVM.lns_luvValueCoreFreeList = []*Lns_luaValueCoreList{}
    luaVM.lns_hasLuvValueCoreFree = false
    luaVM.regexCache = newRegexpCache( 20 )

    lua_createtable( luaVM.vm )
    lua_setglobal( luaVM.vm, lns_globalValSym )
    lua_checkstack( luaVM.vm, 300 )

    Lns_initPreload( luaVM.vm )
    
    go lns_luaValChanProc( luaVM )

    return luaVM
}

func Lns_getVM() *Lns_luaVM {
    return Lns_getVMMain( cur_LnsEnv )
}

func Lns_getVMSync() *Lns_luaVM {
    return Lns_getVMMain( sync_LnsEnv )
}


func Lns_getVMMain( lnsEnv *LnsEnv ) *Lns_luaVM {
    luaVM := lnsEnv.LuaVM
    if luaVM.lns_hasLuvValueCoreFree {
        lns_luvValueCoreFreeListMutex.Lock()
        luaVM.lns_hasLuvValueCoreFree = false
        freeList := luaVM.lns_luvValueCoreFreeList
        luaVM.lns_luvValueCoreFreeList = []*Lns_luaValueCoreList{}
        lns_luvValueCoreFreeListMutex.Unlock()
        for _, coreList := range( freeList ) {
            for _, core := range( coreList.list ) {
                core.free()
            }
        }
    }
    
    return luaVM
}


func Lns_runLuaScript( script string ) {
    vm := luaL_newstate()
    if vm == nil {
        return
    }
    defer lua_close( vm )
    luaL_openlibs( vm )

    luaL_loadstring( vm, script )
    lua_pcallk( vm, 0, cLUA_MULTRET )
}

type lns_pushedString struct {
    str *C.char
}

func (self *lns_pushedString) free() {
    C.free( unsafe.Pointer( self.str ) )
}

func (luaVM *Lns_luaVM) pushStr( txt string ) *lns_pushedString {
    pStr := C.CBytes([]byte(txt))
    lua_pushlstring( luaVM.vm, (*C.char)(pStr), len( txt ) )
    ret := &lns_pushedString{ (*C.char)(pStr) }
    //runtime.SetFinalizer( ret, func (obj *lns_pushedString) { obj.free() } )
    return ret
}

type lns_pushedVal struct {
    pStr *lns_pushedString
}
func (obj *lns_pushedVal) free() {
    if obj.pStr != nil {
        obj.pStr.free()
        obj.pStr = nil;
    }
}
var lns_defaultPushedVal *lns_pushedVal

type StemToLuaConv struct {
    builder *strings.Builder
    luaVM *Lns_luaVM
}

func (luaVM *Lns_luaVM) NewStemToLuaConv() *StemToLuaConv {
    return &StemToLuaConv{ &strings.Builder{}, luaVM }
}

func (self *StemToLuaConv) write( txt string ) {
    self.builder.WriteString( txt )
}


func (self *StemToLuaConv) conv( val LnsAny ) {
    if Lns_IsNil( val ) {
        self.builder.WriteString( "nil" )
    } else {
        switch val.(type) {
        case LnsInt:
            self.builder.WriteString( fmt.Sprintf( "%d", val.(LnsInt) ) )
        case LnsReal:
            self.builder.WriteString( fmt.Sprintf( "%g", val.(LnsInt) ) )
        case bool:
            if val.(bool) {
                self.builder.WriteString( "true" )
            } else {
                self.builder.WriteString( "false" )
            }
        case string:
            self.builder.WriteString(
                self.luaVM.String_format( "%q", Lns_2DDD( val ) ) )
        case *LnsList:
            val.(*LnsList).ToLuaCode( self )
        case *LnsMap:
            val.(*LnsMap).ToLuaCode( self )
        case *Lns_luaValue:
            luaVal := val.(*Lns_luaValue)
            self.builder.WriteString(
                fmt.Sprintf( "%s['%s']", lns_globalValMapName, luaVal.core.getKey() ) )
        default:
            self.builder.WriteString( fmt.Sprintf( "%v", val ) )
            //panic( fmt.Sprintf( "not supoort -- %T", val ) )
        }
    }
}




func (luaVM *Lns_luaVM) pushAny( val LnsAny ) *lns_pushedVal {
    vm := luaVM.vm
    if Lns_IsNil( val ) {
        lua_pushnil( vm )
    } else {
        switch val.(type) {
        case LnsInt:
            lua_pushinteger( vm, val.(LnsInt) )
        case LnsReal:
            lua_pushnumber( vm, val.(LnsReal) )
        case bool:
            lua_pushboolean( vm, val.(bool) )
        case string:
            return &lns_pushedVal{ luaVM.pushStr( val.(string) ) }
        case *Lns_luaValue:
            val.(*Lns_luaValue).core.pushValFromGlobalValMap()
        default:
            conv := luaVM.NewStemToLuaConv()
            conv.write( "return " )
            conv.conv( val )

            loaded, err := luaVM.Load( conv.builder.String(), nil )
            if err != nil {
                panic( fmt.Sprintf( "%s\n%T\n%s", err, val, conv.builder.String() ) )
            }
            luaval := luaVM.RunLoadedfunc( loaded.(*Lns_luaValue),  []LnsAny{} )
            
            luaval[0].(*Lns_luaValue).core.pushValFromGlobalValMap()
        }
    }
    return lns_defaultPushedVal
}

func (luaVM *Lns_luaVM) newLuaValue( index int, typeId int ) *Lns_luaValue {
    // print( fmt.Sprintf( "xxx:%d\n", lua_gettop( luaVM.vm ) ) )

    core := &Lns_luaValueCore{ luaVM: luaVM, typeId: typeId }
    core.sym = C.CString( core.getKey() )
    luaVM.lns_luvValueCoreMap[ core ] = true
    
    val := &Lns_luaValue{ core }
    runtime.SetFinalizer( val, func (obj *Lns_luaValue) { obj.free() } )
    val.core.setValToGlobalValMap( index )
    return val
}

/**
スタックの指定 index の値を取得する

@param index スタック位置
@param passTable 指定位置の値がテーブルの場合の処理方法の指定。
   true の場合、 Go の値に変換せずにそのまま Lns_luaVal として返す。
   false の場合、 LnsMap に変換して返す。
 */
func (luaVM *Lns_luaVM) setupFromStack( index int, passTable bool ) LnsAny {
    vm := luaVM.vm
    switch typeId := lua_type( vm, index ); typeId {
    case cLUA_TNIL:
        return nil
    case cLUA_TNUMBER:
        num := lua_tonumberx( vm, index )
        intVal := LnsInt(lua_tointegerx( vm, index ))
        if LnsReal(LnsInt(num)) == num {
            return intVal
        }
        return LnsReal(num)
    case cLUA_TBOOLEAN:
        return lua_toboolean( vm, index )
    case cLUA_TSTRING:
        return lua_tolstring( vm, index )
    case cLUA_TTABLE:
        if passTable {
            return luaVM.newLuaValue( index, typeId )
        }
        retVal := NewLnsMap( map[LnsAny]LnsAny{} )
        top := lua_gettop( vm )
        defer lua_settop( vm, top )

        // tbl を push
        lua_pushvalue( vm, index );
        lua_pushnil( vm )

        for true {
            if lua_next( vm, -2 ) == 0 {
                return retVal
            }
            // kye = (index:-2) , val = (index:-1)
            val := luaVM.setupFromStack( -1, false );
            if val == nil {
                return retVal
            }
            key := luaVM.setupFromStack( -2, false );
            if key == nil {
                return retVal
            }
            retVal.Items[ key ] = val

            // val を削除して、 key は lua_next() 用に残す
            lua_pop( vm, 1 )
        }
        return retVal
    default:
        return luaVM.newLuaValue( index, typeId )
    }
}


/**
 * Lua の関数コール処理。
 *
 * @param stackBase 関数コール前の基準スタック位置。
 *    この関数処理後は、stackBase 位置をスタックトップとする。
 * @param argNum Lua 関数に渡す引数の数
 * @param retNum Lua 関数の戻り値の数。 不定の場合 LUA_MULTRET。
 * @param Lua 関数の戻り値。
 */
func (luaVM *Lns_luaVM) lua_call( stackBase int , argNum int, retNum int ) []LnsAny {

    argTop := int(lua_gettop( luaVM.vm )) - argNum
    lua_callk( luaVM.vm, argNum, cLUA_MULTRET )

    lastRet := int(lua_gettop( luaVM.vm ))

    result := []LnsAny{}

    stackNum := lastRet - argTop + 1
    if retNum >= 2 || retNum == cLUA_MULTRET {
        for index := argTop; index <= lastRet; index++ {
            result = append( result, luaVM.setupFromStack( index, true ) )
        }
    } else if ( stackNum == 1 ) {
        result = append( result, luaVM.setupFromStack( -1, true ) )
    }

    lua_settop( luaVM.vm, stackBase )

    return result
}


func (luaValue *Lns_luaValue) Call( argList []LnsAny ) []LnsAny {

    luaVM := luaValue.core.luaVM

    top := lua_gettop( luaVM.vm )
    
    luaValue.core.pushValFromGlobalValMap()
    for _, arg := range( argList ) {
        pVal := luaVM.pushAny( arg )
        defer pVal.free()
    }

    return luaVM.lua_call( top, len( argList ), cLUA_MULTRET )
}


/**
index 指定のスタックの値を globalVal[ symbol ] にセットする
*/
func (core *Lns_luaValueCore) setValToGlobalValMap( index int ) {
    vm := core.luaVM.vm
    top := lua_gettop( vm )
    defer lua_settop( vm, top )
    
    // globalVal をスタックトップに push
    lua_getglobal( vm, lns_globalValSym )
    // 指定の値をスタックトップに push
    if index < 0 {
        lua_pushvalue( vm, index - 1 )
    } else {
        lua_pushvalue( vm, index )
    }
    // スタックトップの値を globalVal[ symbol ] にセットし、スタックトップは削除。
    lua_setfield( vm, -2, core.sym )
}

/**
globalVal[ symbol ] をスタックトックに push する
*/
func (core *Lns_luaValueCore) pushValFromGlobalValMap() {
    vm := core.luaVM.vm
    // globalVal をスタックトップに push
    lua_getglobal( vm, lns_globalValSym )
    // globalVal[ symbol ] をスタックトップに push
    lua_getfield( vm, -1, core.sym )
    // globalVal をスタックから除外
    lua_copy( vm, -1, -2 )
    lua_pop( vm, 1 )

    // 本来なら不要だが念の為のチェック。
    // どこかでバグがあると、
    // 記録していたタイプと異なるタイプが保持されている可能性がある。
    if lua_type( vm, -1 ) != core.typeId {
        panic( fmt.Sprintf( "hoge: %d %d\n", core.typeId, lua_type( vm, -1 ) ) )
    }
}


/**
Lua の関数を実行する。

@param packName Lua のパッケージ名。 string.format() を実行する場合 "string"。
@param funcname 実行する関数名。string.format() を実行する場合、 "format"。
@return []LnsAny 実行結果。
*/
func (luaVM *Lns_luaVM) CallStatic(
    packName string, funcname string, args[] LnsAny ) []LnsAny {
    
    vm := luaVM.vm
    top := lua_gettop( vm )
    defer lua_settop( vm, top )

    var argPos int
    if packName == "" {
        argPos = 1
        pFuncname := C.CString( funcname )
        defer C.free( unsafe.Pointer( pFuncname ) )
        lua_getglobal( vm, pFuncname )
    } else {
        argPos = 2
        pPackName := C.CString( packName )
        defer C.free( unsafe.Pointer( pPackName ) )
        pFuncname := C.CString( funcname )
        defer C.free( unsafe.Pointer( pFuncname ) )
        lua_getglobal( vm, pPackName )
        lua_getfield(vm, -1, pFuncname )
    }

    
    for _, val := range( args ) {
        pVal := luaVM.pushAny( val )
        defer pVal.free()
    }
    if lua_pcallk( vm, len( args ), cLUA_MULTRET ) != cLUA_OK {
        panic( lua_tolstring( vm, -1 ) )
    }
    ret := []LnsAny{}
    nowTop := lua_gettop( vm )
    for index := top + argPos; index <= nowTop; index++ {
        ret = append( ret, luaVM.setupFromStack( index, true ) )
    }
    
    return ret
}

func (obj *Lns_luaValue) CallMethod( funcname string, args[] LnsAny ) []LnsAny {
    luaVM := obj.core.luaVM
    
    vm := luaVM.vm
    top := lua_gettop( vm )
    defer lua_settop( vm, top )

    pFuncname := C.CString( funcname )
    defer C.free( unsafe.Pointer( pFuncname ) )

    obj.core.pushValFromGlobalValMap() // obj を push
    lua_getfield(vm, -1, pFuncname ) // obj の pFuncname のメソッドを取得
    obj.core.pushValFromGlobalValMap() // obj を引数に push
    
    for _, val := range( args ) {
        pVal := luaVM.pushAny( val ) // arg を push
        defer pVal.free()
    }
    if lua_pcallk( vm, len( args ) + 1, cLUA_MULTRET ) != cLUA_OK {
        panic( lua_tolstring( vm, -1 ) )
    }
    ret := []LnsAny{}
    nowTop := lua_gettop( vm )
    for index := top + 2; index <= nowTop; index++ {
        ret = append( ret, luaVM.setupFromStack( index, true ) )
    }
    
    return ret
}

func (obj *Lns_luaValue) GetAt( index LnsAny ) LnsAny {
    luaVM := obj.core.luaVM
    
    vm := luaVM.vm
    top := lua_gettop( vm )
    defer lua_settop( vm, top )

    obj.core.pushValFromGlobalValMap() // obj を push
    pVal := luaVM.pushAny( index ) // arg を push
    defer pVal.free()
    lua_gettable( vm, -2 ); // obj[arg] を push
    return luaVM.setupFromStack( -1, true )
}

func (obj *Lns_luaValue) Len() LnsInt {
    luaVM := obj.core.luaVM
    vm := luaVM.vm
    top := lua_gettop( vm )
    defer lua_settop( vm, top )
    obj.core.pushValFromGlobalValMap() // obj を push

    lua_len( vm, -1 )
    return lua_tointegerx(vm, -1 )
}
