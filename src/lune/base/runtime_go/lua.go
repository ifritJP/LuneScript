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

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

// vm を再利用する際に、クリーンな状態で利用するための
// VM 初期化処理を実行するかどうかのフラグ。
// 有効にした方が安全だが、現状は無効にしていても不具合がないので
// 速度優先のために無効化しておく。
const lns_enableRevertingVM = false

var lns_globalValSym lua_rawstr

const lns_globalValMapName = "lns_globalValMap"

var lns_luvValueCoreFreeListMutex sync.Mutex

var lns_createVMMutex sync.Mutex
var lns_freeVMMap map[*Lns_luaVM]bool = map[*Lns_luaVM]bool{}

// - lua vm 内の値で、 go のデータに変換出来ない値を扱う時に、
//   Go の構造体 Lns_luaValue を定義して管理している。
// - この Lns_luaValue で管理する Lua の値が lua VM 内で開放されないように
//   参照を保持するため Lua VM 内の lns_globalValMap という object に値をセットしている。
// - この値をセットする際の object の key として C 文字列を生成している。
// - この C 文字列は、 Go の GC 対象外なので自前で開放する必要がある。
// - その開放制御を lns_luaValChanProc() で行なっている。
// - lns_luaValChanProc() は、 init() 時に go-routine で起動している。
// - lns_luaValChanProc() は、 *Lns_luaValueCore 型の channel を受信しつづけ、
//   受信した core 情報をリストに保持する。
// - そのリストが一定数(lns_luaValueCoreNum)溜ったら、
//    lns_hasLuvValueCoreFree に true をセットし、
//    lns_luvValueCoreFreeList に append する。
// - lns_luvValueCoreFreeList に格納された core は、
//   Lns_processFreeList() 実行時に開放される。
// - lns_freeVMMap へのアクセスは、 mutex で排他する。
type Lns_luaValueCore struct {
	// lns_globalValMap に格納しているシンボル名
	sym lua_rawstr
	// Lua vm
	luaVM *Lns_luaVM
	// 格納しているデータタイプ
	typeId int
}

type Lns_luaValueCoreList struct {
	list []*Lns_luaValueCore
}

type Lns_luaValue struct {
	core *Lns_luaValueCore
}

const lns_luaValueCoreNum = 10

var lns_luaValChan chan *Lns_luaValueCore
var lns_luaValProcEndChan chan bool

/**
symbol を globalVal から除外する
*/
func (luaValue *Lns_luaValue) free() {
	lns_luaValChan <- luaValue.core
}

/**
symbol を globalVal から除外する
*/
func (core *Lns_luaValueCore) free() {
	luaVM := core.luaVM
	delete(luaVM.lns_luvValueCoreMap, core)

	if !luaVM.closed {
		top := lua_gettop(luaVM.vm)
		lua_pushnil(luaVM.vm)
		core.setValToGlobalValMap(-1)
		lua_settop(luaVM.vm, top)
	}

	Lns_freeRawStr(core.sym)
	core.sym = Lns_zeroRawStr()
}

func (core *Lns_luaValueCore) getKey() string {
	return fmt.Sprintf("%p", core)
}

type Lns_luaVM struct {
	vm lua_state

	lns_luvValueCoreMap map[*Lns_luaValueCore]bool
	closed              bool
	freeCoreWorkList    []*Lns_luaValueCore
	freeCoreWorkIndex   int

	funcForReverting *Lns_luaValue

	lns_luvValueCoreFreeList []*Lns_luaValueCoreList
	lns_hasLuvValueCoreFree  bool

	regexCache *RegexpCache
}

func (luaVM *Lns_luaVM) freeLuaValueCore() {
	luaVM.freeCoreWorkIndex = 0
	freeList := &Lns_luaValueCoreList{luaVM.freeCoreWorkList}

	lns_luvValueCoreFreeListMutex.Lock()
	luaVM.lns_hasLuvValueCoreFree = true
	luaVM.lns_luvValueCoreFreeList =
		append(luaVM.lns_luvValueCoreFreeList, freeList)
	if luaVM.closed {
		// closed の場合は、ここで開放する
		Lns_processFreeListMain(luaVM.lns_luvValueCoreFreeList)
		luaVM.lns_hasLuvValueCoreFree = false
		luaVM.lns_luvValueCoreFreeList = []*Lns_luaValueCoreList{}
	}
	lns_luvValueCoreFreeListMutex.Unlock()

	luaVM.freeCoreWorkList = make([]*Lns_luaValueCore, lns_luaValueCoreNum)
}

func lns_luaValChanProc() {
	for {
		core := <-lns_luaValChan

		if core == nil {
			break
		}

		luaVM := core.luaVM
		list := luaVM.freeCoreWorkList

		list[luaVM.freeCoreWorkIndex] = core
		luaVM.freeCoreWorkIndex++

		if luaVM.freeCoreWorkIndex == lns_luaValueCoreNum {
			luaVM.freeLuaValueCore()
		}
	}
	lns_luaValProcEndChan <- true
}

var lnsLnsMap map[string]string

func AddlnsLnsInfo(key string, code string) {
	lnsLnsMap[key] = code
}
func depend_getBindLns(mod string) LnsAny {
	if code, has := lnsLnsMap[mod]; has {
		return code
	}
	return nil
}

func init() {
	lnsLnsMap = map[string]string{}
	lns_globalValSym = Lns_toRawStr(lns_globalValMapName)
	lns_defaultPushedVal = &lns_pushedVal{}
	lns_luaValChan = make(chan *Lns_luaValueCore, 1000)
	lns_luaValProcEndChan = make(chan bool, 0)

	go lns_luaValChanProc()

}

func getVMFromCache() *Lns_luaVM {
	var luaVM *Lns_luaVM = nil

	lns_createVMMutex.Lock()
	defer lns_createVMMutex.Unlock()

	for vm := range lns_freeVMMap {
		luaVM = vm
		break
	}
	if luaVM != nil {
		// 空き情報から削除
		delete(lns_freeVMMap, luaVM)
	}
	return luaVM
}

func createVM() *Lns_luaVM {
	luaVM := getVMFromCache()

	if luaVM == nil {
		luaVM = &Lns_luaVM{}
		luaVM.vm = LuaL_newstate(300)
		LuaL_openlibs(luaVM.vm)
		luaVM.lns_luvValueCoreMap = map[*Lns_luaValueCore]bool{}
		luaVM.freeCoreWorkList = make([]*Lns_luaValueCore, lns_luaValueCoreNum)
		luaVM.freeCoreWorkIndex = 0

		luaVM.closed = false
		luaVM.lns_luvValueCoreFreeList = []*Lns_luaValueCoreList{}
		luaVM.lns_hasLuvValueCoreFree = false
		luaVM.regexCache = newRegexpCache(20)

		lua_createtable(luaVM.vm)
		lua_setglobal(luaVM.vm, lns_globalValSym)

		Lns_initPreload(luaVM.vm)

		if lns_enableRevertingVM {
			// デフォルトの global 値を lns_builtinGlobals に保持しておく。
			// デフォルトの package.loaded 値を lns_builtinLoaded に保持しておく
			// これらは closeVM() 時に、 VM の状態をデフォルトに戻すために保持しておく
			luaVM.run_script(`
lns_builtinGlobals = {}
lns_builtinLoaded = {}
lns_builtinGlobals[ "lns_builtinLoaded" ] = true
lns_builtinGlobals[ "lns_builtinGlobals" ] = true
for key, val in pairs( _G ) do
   lns_builtinGlobals[ key ] = val
end
for key, val in pairs( package.loaded ) do
   lns_builtinLoaded[ key ] = val
end
`, []LnsAny{})

			// createVM() で保持しておいたデフォルト値以外を開放する処理を生成しておく
			loaded, err := luaVM.Load(`
local function release( target, default )
   local work = {}
   for key, val in pairs( target ) do
       if not default[ key ] then
          table.insert( work, key )
       end
   end
   for key, val in ipairs( work ) do
       target[ val ] = nil
   end
end
release( _G, lns_builtinGlobals )
release( package.loaded, lns_builtinLoaded )
collectgarbage()
`, nil)

			if err != nil {
				panic(err)
			}
			luaVM.funcForReverting = loaded.(*Lns_luaValue)
		}
	}

	return luaVM
}

func (self *Lns_luaVM) closeVM() {

	if lns_enableRevertingVM {
		// vm の状態を、初期化状態に戻す処理を実行
		self.RunLoadedfunc(self.funcForReverting, nil)
	}

	// freeList に登録された core 情報の開放
	Lns_processFreeList(self)

	{
		lns_createVMMutex.Lock()
		// 空き vm を登録
		lns_freeVMMap[self] = true
		lns_createVMMutex.Unlock()
	}

	// ここでは lua_close() しない。
	// 最終的に Lns_shutdownAllVM() をコールした時に lua_close() する。
	// lua_close(self.vm)
}

// 全ての LUA VM を停止させる
//
// これを実行した後は、 createVM() を実行してはならない。
// もし createVM() を実行したい場合は、
//
func Lns_shutdownAllVM() {
	// vm.funcForReverting を開放するために nil をセットする
	for vm := range lns_freeVMMap {
		vm.funcForReverting = nil
	}

	// VM を止める前に GC を呼ぶ。
	// 通常は、この GC で lns_luaValueCore が lns_luaValChanProc で処理されているはず。
	runtime.GC()

	// lns_luaValChanProc への終了要求
	lns_luaValChan <- nil

	// 終了待ち
	<-lns_luaValProcEndChan

	// Mutex ロック
	lns_luvValueCoreFreeListMutex.Lock()
	defer lns_luvValueCoreFreeListMutex.Unlock()

	lns_createVMMutex.Lock()
	defer lns_createVMMutex.Unlock()

	// vm を close する
	for vm := range lns_freeVMMap {
		vm.freeLuaValueCore()
		lua_close(vm.vm)
		vm.closed = true
	}
	lns_freeVMMap = map[*Lns_luaVM]bool{}

	// 念のため、 lns_luaValueCore のインスタンスが残っていた時のために、
	// 起動しておく。
	// go routine が一つ残ってしまうが、下手にメモリリークするよりはマシ。
	go lns_luaValChanProc()
}

func Lns_getVM() *Lns_luaVM {
	return Lns_processFreeList(cur_LnsEnv.LuaVM)
}

func Lns_getVMSync() *Lns_luaVM {
	return Lns_processFreeList(sync_LnsEnv.LuaVM)
}

func Lns_processFreeListMain(freeList []*Lns_luaValueCoreList) {
	for _, coreList := range freeList {
		for _, core := range coreList.list {
			if core == nil {
				break
			}
			core.free()
		}
	}
}

func Lns_processFreeList(luaVM *Lns_luaVM) *Lns_luaVM {
	if luaVM.lns_hasLuvValueCoreFree {
		lns_luvValueCoreFreeListMutex.Lock()
		luaVM.lns_hasLuvValueCoreFree = false
		freeList := luaVM.lns_luvValueCoreFreeList
		luaVM.lns_luvValueCoreFreeList = []*Lns_luaValueCoreList{}
		lns_luvValueCoreFreeListMutex.Unlock()
		Lns_processFreeListMain(freeList)
	}

	return luaVM
}

type lns_pushedString struct {
	str lua_rawstr
}

func (self *lns_pushedString) free() {
	Lns_freeRawStr(self.str)
}

func (luaVM *Lns_luaVM) pushStr(txt string) *lns_pushedString {
	pStr := Lns_toRawStr(txt)
	lua_pushlstring(luaVM.vm, pStr, len(txt))
	ret := &lns_pushedString{pStr}
	return ret
}

type lns_pushedVal struct {
	pStr *lns_pushedString
}

func (obj *lns_pushedVal) free() {
	if obj.pStr != nil {
		obj.pStr.free()
		obj.pStr = nil
	}
}

var lns_defaultPushedVal *lns_pushedVal

type StemToLuaConv struct {
	builder *strings.Builder
	luaVM   *Lns_luaVM
}

func (luaVM *Lns_luaVM) NewStemToLuaConv() *StemToLuaConv {
	return &StemToLuaConv{&strings.Builder{}, luaVM}
}

func (self *StemToLuaConv) write(txt string) {
	self.builder.WriteString(txt)
}

func (self *StemToLuaConv) conv(val LnsAny) {
	if Lns_IsNil(val) {
		self.builder.WriteString("nil")
	} else {
		switch val.(type) {
		case LnsInt:
			self.builder.WriteString(fmt.Sprintf("%d", val.(LnsInt)))
		case LnsReal:
			self.builder.WriteString(fmt.Sprintf("%g", val.(LnsInt)))
		case bool:
			if val.(bool) {
				self.builder.WriteString("true")
			} else {
				self.builder.WriteString("false")
			}
		case string:
			self.builder.WriteString(
				self.luaVM.String_format("%q", Lns_2DDD(val)))
		case *LnsList:
			val.(*LnsList).ToLuaCode(self)
		case *LnsMap:
			val.(*LnsMap).ToLuaCode(self)
		case *Lns_luaValue:
			luaVal := val.(*Lns_luaValue)
			self.builder.WriteString(
				fmt.Sprintf("%s['%s']", lns_globalValMapName, luaVal.core.getKey()))
		default:
			self.builder.WriteString(fmt.Sprintf("%v", val))
			//panic( fmt.Sprintf( "not supoort -- %T", val ) )
		}
	}
}

func (luaVM *Lns_luaVM) pushAny(val LnsAny) *lns_pushedVal {
	vm := luaVM.vm
	if Lns_IsNil(val) {
		lua_pushnil(vm)
	} else {
		switch val.(type) {
		case LnsInt:
			lua_pushinteger(vm, val.(LnsInt))
		case LnsReal:
			lua_pushnumber(vm, val.(LnsReal))
		case bool:
			lua_pushboolean(vm, val.(bool))
		case string:
			return &lns_pushedVal{luaVM.pushStr(val.(string))}
		case *Lns_luaValue:
			val.(*Lns_luaValue).core.pushValFromGlobalValMap()
		default:
			conv := luaVM.NewStemToLuaConv()
			conv.write("return ")
			conv.conv(val)

			loaded, err := luaVM.Load(conv.builder.String(), nil)
			if err != nil {
				panic(fmt.Sprintf("%s\n%T\n%s", err, val, conv.builder.String()))
			}
			luaval := luaVM.RunLoadedfunc(loaded.(*Lns_luaValue), []LnsAny{})

			luaval[0].(*Lns_luaValue).core.pushValFromGlobalValMap()
		}
	}
	return lns_defaultPushedVal
}

func (luaVM *Lns_luaVM) newLuaValue(index int, typeId int) *Lns_luaValue {
	// print( fmt.Sprintf( "xxx:%d\n", lua_gettop( luaVM.vm ) ) )

	core := &Lns_luaValueCore{luaVM: luaVM, typeId: typeId}
	core.sym = Lns_toRawStr(core.getKey())
	luaVM.lns_luvValueCoreMap[core] = true

	val := &Lns_luaValue{core}
	runtime.SetFinalizer(val, func(obj *Lns_luaValue) { obj.free() })
	val.core.setValToGlobalValMap(index)
	return val
}

/**
スタックの指定 index の値を取得する

@param index スタック位置
@param passTable 指定位置の値がテーブルの場合の処理方法の指定。
   true の場合、 Go の値に変換せずにそのまま Lns_luaVal として返す。
   false の場合、 LnsMap に変換して返す。
*/
func (luaVM *Lns_luaVM) setupFromStack(index int, passTable bool) LnsAny {
	vm := luaVM.vm
	switch typeId := lua_type(vm, index); typeId {
	case cLUA_TNIL:
		return nil
	case cLUA_TNUMBER:
		num := lua_tonumberx(vm, index)
		intVal := LnsInt(lua_tointegerx(vm, index))
		if LnsReal(LnsInt(num)) == num {
			return intVal
		}
		return LnsReal(num)
	case cLUA_TBOOLEAN:
		return lua_toboolean(vm, index)
	case cLUA_TSTRING:
		return lua_tolstring(vm, index)
	case cLUA_TTABLE:
		if passTable {
			return luaVM.newLuaValue(index, typeId)
		}
		retVal := NewLnsMap(map[LnsAny]LnsAny{})
		top := lua_gettop(vm)
		defer lua_settop(vm, top)

		// tbl を push
		lua_pushvalue(vm, index)
		lua_pushnil(vm)

		for true {
			if lua_next(vm, -2) == 0 {
				return retVal
			}
			// kye = (index:-2) , val = (index:-1)
			val := luaVM.setupFromStack(-1, false)
			if val == nil {
				return retVal
			}
			key := luaVM.setupFromStack(-2, false)
			if key == nil {
				return retVal
			}
			retVal.Items[key] = val

			// val を削除して、 key は lua_next() 用に残す
			lua_pop(vm, 1)
		}
		return retVal
	default:
		return luaVM.newLuaValue(index, typeId)
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
func (luaVM *Lns_luaVM) lua_call(stackBase int, argNum int, retNum int) []LnsAny {

	argTop := int(lua_gettop(luaVM.vm)) - argNum
	lua_callk(luaVM.vm, argNum, cLUA_MULTRET)

	lastRet := int(lua_gettop(luaVM.vm))

	result := []LnsAny{}

	stackNum := lastRet - argTop + 1
	if retNum >= 2 || retNum == cLUA_MULTRET {
		for index := argTop; index <= lastRet; index++ {
			result = append(result, luaVM.setupFromStack(index, true))
		}
	} else if stackNum == 1 {
		result = append(result, luaVM.setupFromStack(-1, true))
	}

	lua_settop(luaVM.vm, stackBase)

	return result
}

func (luaValue *Lns_luaValue) Call(argList []LnsAny) []LnsAny {

	luaVM := luaValue.core.luaVM

	top := lua_gettop(luaVM.vm)

	luaValue.core.pushValFromGlobalValMap()
	for _, arg := range argList {
		pVal := luaVM.pushAny(arg)
		defer pVal.free()
	}

	return luaVM.lua_call(top, len(argList), cLUA_MULTRET)
}

/**
index 指定のスタックの値を globalVal[ symbol ] にセットする
*/
func (core *Lns_luaValueCore) setValToGlobalValMap(index int) {
	vm := core.luaVM.vm
	top := lua_gettop(vm)
	defer lua_settop(vm, top)

	// globalVal をスタックトップに push
	lua_getglobal(vm, lns_globalValSym)
	// 指定の値をスタックトップに push
	if index < 0 {
		lua_pushvalue(vm, index-1)
	} else {
		lua_pushvalue(vm, index)
	}
	// スタックトップの値を globalVal[ symbol ] にセットし、スタックトップは削除。
	lua_setfield(vm, -2, core.sym)
}

/**
globalVal[ symbol ] をスタックトックに push する
*/
func (core *Lns_luaValueCore) pushValFromGlobalValMap() {
	vm := core.luaVM.vm
	// globalVal をスタックトップに push
	lua_getglobal(vm, lns_globalValSym)
	// globalVal[ symbol ] をスタックトップに push
	lua_getfield(vm, -1, core.sym)
	// globalVal をスタックから除外
	// lua_copy( vm, -1, -2 )
	// lua_pop( vm, 1 )
	lua_replace(vm, lua_gettop(vm)-1)

	// 本来なら不要だが念の為のチェック。
	// どこかでバグがあると、
	// 記録していたタイプと異なるタイプが保持されている可能性がある。
	if lua_type(vm, -1) != core.typeId {
		panic(fmt.Sprintf("hoge: %d %d\n", core.typeId, lua_type(vm, -1)))
	}
}

/**
Lua の関数を実行する。

@param packName Lua のパッケージ名。 string.format() を実行する場合 "string"。
@param funcname 実行する関数名。string.format() を実行する場合、 "format"。
@return []LnsAny 実行結果。
*/
func (luaVM *Lns_luaVM) CallStatic(
	packName string, funcname string, args []LnsAny) []LnsAny {

	vm := luaVM.vm
	top := lua_gettop(vm)
	defer lua_settop(vm, top)

	var argPos int
	if packName == "" {
		argPos = 1
		pFuncname := Lns_toRawStr(funcname)
		defer Lns_freeRawStr(pFuncname)
		lua_getglobal(vm, pFuncname)
	} else {
		argPos = 2
		pPackName := Lns_toRawStr(packName)
		defer Lns_freeRawStr(pPackName)
		pFuncname := Lns_toRawStr(funcname)
		defer Lns_freeRawStr(pFuncname)
		lua_getglobal(vm, pPackName)
		lua_getfield(vm, -1, pFuncname)
	}

	for _, val := range args {
		pVal := luaVM.pushAny(val)
		defer pVal.free()
	}
	if err := lua_pcallk(vm, len(args), cLUA_MULTRET); err != nil {
		panic(err.Error())
	}
	ret := []LnsAny{}
	nowTop := lua_gettop(vm)
	for index := top + argPos; index <= nowTop; index++ {
		ret = append(ret, luaVM.setupFromStack(index, true))
	}

	return ret
}

func (obj *Lns_luaValue) CallMethod(funcname string, args []LnsAny) []LnsAny {
	luaVM := obj.core.luaVM

	vm := luaVM.vm
	top := lua_gettop(vm)
	defer lua_settop(vm, top)

	pFuncname := Lns_toRawStr(funcname)
	defer Lns_freeRawStr(pFuncname)

	obj.core.pushValFromGlobalValMap() // obj を push
	lua_getfield(vm, -1, pFuncname)    // obj の pFuncname のメソッドを取得
	obj.core.pushValFromGlobalValMap() // obj を引数に push

	for _, val := range args {
		pVal := luaVM.pushAny(val) // arg を push
		defer pVal.free()
	}
	if err := lua_pcallk(vm, len(args)+1, cLUA_MULTRET); err != nil {
		panic(err.Error())
	}
	ret := []LnsAny{}
	nowTop := lua_gettop(vm)
	for index := top + 2; index <= nowTop; index++ {
		ret = append(ret, luaVM.setupFromStack(index, true))
	}

	return ret
}

func (obj *Lns_luaValue) GetAt(index LnsAny) LnsAny {
	luaVM := obj.core.luaVM

	vm := luaVM.vm
	top := lua_gettop(vm)
	defer lua_settop(vm, top)

	obj.core.pushValFromGlobalValMap() // obj を push
	pVal := luaVM.pushAny(index)       // arg を push
	defer pVal.free()
	lua_gettable(vm, -2) // obj[arg] を push
	return luaVM.setupFromStack(-1, true)
}

func (obj *Lns_luaValue) Len() LnsInt {
	luaVM := obj.core.luaVM
	vm := luaVM.vm
	top := lua_gettop(vm)
	defer lua_settop(vm, top)
	obj.core.pushValFromGlobalValMap() // obj を push

	len := lua_len(vm, -1)

	return len
}
