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
	"math"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type LnsInt = int
type LnsReal = float64
type LnsAny = interface{}

var LnsNone interface{} = nil
var Lns_package_path string

type LnsEnv struct {
	// and or 演算で利用するスタック
	valStack []LnsAny
	// nil アクセス演算子で利用するスタック
	nilAccStack []LnsAny
	// gsub などの runtime 移植していない Lua API を動かすための LuaVM
	LuaVM *Lns_luaVM
	// load などの、全体を通して共通で動作させる必要のある VM
	CommonLuaVM *Lns_luaVM
	// async 用の Env かどうか。現在の設定。 __asyncLock している間は変る。
	async bool
	// async 用の Env かどうか。起動時の設定。
	orgAsync bool
	// runner の名前。 runner 実行時にセットする。
	runnerName string
	// runner の id
	runnerId int
}

// デフォルトのシングルタスクで使用する LnsEnv
var cur_LnsEnv *LnsEnv

// / __nosync を排他するための mutex
var sync_LnsEnvMutex sync.Mutex

func Lns_GetEnv() *LnsEnv {
	return cur_LnsEnv
}

func (self *LnsEnv) GetVM() *Lns_luaVM {
	if self.async {
		return self.LuaVM
	}
	return self.CommonLuaVM
}

/*
*
各モジュールを初期化する際に実行する関数。

import エラーを回避するため、
敢てランタイムの関数のどれか一つを呼んでいる。

実際の初期化関数は、 Lns_InitModOnce() で行なう。
*/
func Lns_InitMod() {
}

/*
*

	go 側から直接 Async 用の LnsEnv を生成する場合に利用する。

	通常は main や run の内部処理から生成するので、これを直接は利用しない。

	使用が終ったら Lns_ReleaseEnv() を実行して開放する。
*/
func Lns_createAsyncEnv(runnerName string) *LnsEnv {
	return lns_threadMgrInfo.createAsyncEnv(runnerName)
}

/*
go 側から直接 LnsEnv を生成する場合に利用する。

	Lns_createEnv() で生成した env を開放する。
*/
func Lns_releaseEnv(env *LnsEnv) {
	env.LuaVM.closeVM()
}

func createEnv(async bool, runnerName string, runnerId int) *LnsEnv {
	env := &LnsEnv{}
	env.valStack = make([]LnsAny, 2)
	env.nilAccStack = make([]LnsAny, 2)
	env.runnerName = runnerName
	env.runnerId = runnerId
	env.LuaVM = requestVM()
	if async {
		env.CommonLuaVM = cur_LnsEnv.LuaVM
	} else {
		env.CommonLuaVM = env.LuaVM
	}
	env.async = async
	env.orgAsync = async

	return env
}

func exitRuntime(code LnsInt) {
	if validRuntimeLog {
		fmt.Printf("----------\n")
		fmt.Printf("vmFreeMap = %d\n", len(lns_freeVMMap))
		fmt.Printf("totalReqVM = %d\n", lns_countOfRequestToCreateVM)
		lns_threadMgrInfo.dump(os.Stdout)
		fmt.Printf("----------\n")
	}
	os.Exit(code)
}

// 整数を文字列に変換する場合、 .0 を付加するかどうかのモード
type Int2strMode_t int

const (
	// Lua のバージョンに依存。
	// lua5.1 は .0 不要。
	// lua5.3 は .0 必要。
	Int2strModeDepend Int2strMode_t = iota
	// .0 を付加する
	Int2strModeNeed0
	// .0 不要
	Int2strModeUnneed0
)

type LnsRuntimeOpt struct {
	Int2strMode Int2strMode_t
}

var lnsRuntimeOpt LnsRuntimeOpt

var launchTime time.Time

/*
*
各モジュールを初期化する際に実行する関数。
*/
func Lns_InitModOnce(opts ...LnsRuntimeOpt) {

	if len(opts) == 0 {
		lnsRuntimeOpt = LnsRuntimeOpt{Int2strMode: Int2strModeDepend}
	} else {
		lnsRuntimeOpt = opts[0]
	}

	// go producerLuaVM()

	cur_LnsEnv = createEnv(false, "main", 0)

	Lns_package_path = cur_LnsEnv.LuaVM.GetPackagePath()

	// __asyncLock 用にロックしておく
	sync_LnsEnvMutex.Lock()

	launchTime = time.Now()
}

func Lns_IsNil(val LnsAny) bool {
	if val == nil {
		return true
	}
	// switch val.(type) {
	// case LnsInt:
	// 	return false
	// case LnsReal:
	// 	return false
	// case bool:
	// 	return false
	// case string:
	// 	return false
	// default:
	// 	value := reflect.ValueOf(val)
	// 	if value.Kind() == reflect.Struct {
	// 		return false
	// 	}
	// 	return value.IsNil()
	// }
	value := reflect.ValueOf(val)
	switch value.Kind() {
	case reflect.Int:
		return false
	case reflect.Float64:
		return false
	case reflect.Bool:
		return false
	case reflect.String:
		return false
	case reflect.Struct:
		return false
	default:
		return value.IsNil()
	}
}

func Lns_type(val LnsAny) string {
	if val == nil {
		return "nil"
	}
	switch val.(type) {
	case LnsInt:
		return "number"
	case LnsReal:
		return "number"
	case bool:
		return "boolean"
	case string:
		return "string"
	case *Lns_luaValue:
		switch val.(*Lns_luaValue).core.typeId {
		case cLUA_TFUNCTION:
			return "function"
		case cLUA_TTABLE:
			return "table"
		default:
			return fmt.Sprintf("<notsupport>:%d, %d, %s",
				val.(*Lns_luaValue).core.typeId, cLUA_TTABLE,
				val.(*Lns_luaValue).core.typeId == cLUA_TTABLE)
		}
	default:
		value := reflect.ValueOf(val)
		if value.Kind() == reflect.Func {
			return "function"
		}
		return "table"
	}
}

func Lns_tonumber(val string, base LnsAny) LnsAny {
	if Lns_IsNil(base) {
		if str_startsWith(val, "0x") || str_startsWith(val, "0X") {
			if dig, err := strconv.ParseInt(val[2:], 16, 64); err != nil {
				return nil
			} else {
				return LnsReal(dig)
			}
		}
		f, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil
		}
		return f
	}
	if bs, ok := base.(LnsInt); ok {
		if dig, err := strconv.ParseInt(val, bs, 64); err != nil {
			return nil
		} else {
			return LnsReal(dig)
		}
	} else {
		return nil
	}
}

func Lns_print(multi []LnsAny) {
	for index, val := range multi {
		if index != 0 {
			fmt.Print("\t")
		}
		fmt.Print(Lns_ToString(val))
	}
	fmt.Print("\n")
}

func Lns_require(val string) LnsAny {
	panic("not support")
	return nil
}

func Lns_unwrap(val LnsAny) LnsAny {
	if Lns_IsNil(val) {
		panic("unwrap nil")
	}
	return val
}
func Lns_unwrapDefault(val, def LnsAny) LnsAny {
	if Lns_IsNil(val) {
		return def
	}
	return val
}

func Lns_cast2LnsInt(val LnsAny) LnsAny {
	if _, ok := val.(LnsInt); ok {
		return val
	}
	return nil
}
func Lns_cast2LnsReal(val LnsAny) LnsAny {
	if _, ok := val.(LnsReal); ok {
		return val
	}
	return nil
}
func Lns_cast2bool(val LnsAny) LnsAny {
	if _, ok := val.(bool); ok {
		return val
	}
	return nil
}
func Lns_cast2string(val LnsAny) LnsAny {
	if _, ok := val.(string); ok {
		return val
	}
	return nil
}

func Lns_forceCastInt(val LnsAny) LnsInt {
	switch val.(type) {
	case LnsInt:
		return val.(LnsInt)
	case LnsReal:
		return LnsInt(val.(LnsReal))
	}
	panic(fmt.Sprintf("illegal type -- %T", val))
}

func Lns_forceCastReal(val LnsAny) LnsReal {
	switch val.(type) {
	case LnsInt:
		return LnsReal(val.(LnsInt))
	case LnsReal:
		return val.(LnsReal)
	}
	panic(fmt.Sprintf("illegal type -- %T", val))
}

type Lns_ToObjParam struct {
	Func    func(obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny)
	Nilable bool
	Child   []Lns_ToObjParam
}
type Lns_ToObj func(obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny)

func Lns_ToStemSub(
	obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny) {
	if Lns_IsNil(obj) {
		if nilable {
			return true, nil, nil
		}
		return false, nil, "nil"
	}
	return true, obj, nil
}
func Lns_ToIntSub(
	obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny) {
	if Lns_IsNil(obj) {
		if nilable {
			return true, nil, nil
		}
		return false, nil, "nil"
	}
	if _, ok := obj.(LnsInt); ok {
		return true, obj, nil
	}
	if val, ok := obj.(LnsReal); ok {
		if math.Ceil(val) == val {
			return true, LnsInt(val), nil
		}
		return true, val, nil
	}
	return false, nil, fmt.Sprintf("no int: %s", reflect.ValueOf(obj).Kind())
}
func Lns_ToRealSub(obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny) {
	if Lns_IsNil(obj) {
		if nilable {
			return true, nil, nil
		}
		return false, nil, "nil"
	}
	if _, ok := obj.(LnsReal); ok {
		return true, obj, nil
	}
	return false, nil, "no real"
}
func Lns_ToBoolSub(obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny) {
	if Lns_IsNil(obj) {
		if nilable {
			return true, nil, nil
		}
		return false, nil, "nil"
	}
	if _, ok := obj.(bool); ok {
		return true, obj, nil
	}
	return false, nil, "no bool"
}
func Lns_ToStrSub(obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny) {
	if Lns_IsNil(obj) {
		if nilable {
			return true, nil, nil
		}
		return false, nil, "nil"
	}
	if _, ok := obj.(string); ok {
		return true, obj, nil
	}
	return false, nil, "no str"
}

type LnsAlgeVal interface {
	GetTxt() string
}

func Lns_getFromMulti(multi []LnsAny, index LnsInt) LnsAny {
	if len(multi) > index {
		return multi[index]
	}
	return nil
}

func (self *LnsEnv) NilAccPush(obj interface{}) bool {
	if Lns_IsNil(obj) {
		return false
	}
	self.nilAccStack = append(self.nilAccStack, obj)
	return true
}

// func Lns_NilAccLast( obj interface{} ) bool {
//     self.nilAccStack = append( self.nilAccStack, obj )
//     return true
// }

func (self *LnsEnv) NilAccPop() LnsAny {
	obj := self.nilAccStack[len(self.nilAccStack)-1]
	self.nilAccStack = self.nilAccStack[:len(self.nilAccStack)-1]
	return obj
}

func (self *LnsEnv) NilAccFin(ret bool) LnsAny {
	if ret {
		return self.NilAccPop()
	}
	return nil
}

func Lns_NilAccCall0(self *LnsEnv, call func()) bool {
	call()
	return self.NilAccPush(true)
}
func Lns_NilAccCall1(self *LnsEnv, call func() LnsAny) bool {
	return self.NilAccPush(call())
}
func Lns_NilAccCall2(self *LnsEnv, call func() (LnsAny, LnsAny)) bool {
	return self.NilAccPush(Lns_2DDD(call()))
}
func Lns_NilAccFinCall2(ret LnsAny) (LnsAny, LnsAny) {
	if Lns_IsNil(ret) {
		return nil, nil
	}
	list := ret.([]LnsAny)
	return list[0], list[1]
}
func Lns_NilAccCall3(self *LnsEnv, call func() (LnsAny, LnsAny, LnsAny)) bool {
	return self.NilAccPush(Lns_2DDD(call()))
}
func Lns_NilAccFinCall3(ret LnsAny) (LnsAny, LnsAny, LnsAny) {
	if Lns_IsNil(ret) {
		return nil, nil, nil
	}
	list := ret.([]LnsAny)
	return list[0], list[1], list[2]
}

func Lns_isCondTrue(stem LnsAny) bool {
	if Lns_IsNil(stem) {
		return false
	}
	switch stem.(type) {
	case bool:
		return stem.(bool)
	default:
		return true
	}
}

func Lns_op_not(stem LnsAny) bool {
	return !Lns_isCondTrue(stem)
}

/** 多値返却の先頭を返す
 */
func Lns_car(multi ...LnsAny) LnsAny {
	if len(multi) == 0 {
		return nil
	}
	if Lns_IsNil(multi[0]) {
		return nil
	}
	if ddd, ok := multi[0].([]LnsAny); ok {
		return Lns_car(ddd...)
	}
	return multi[0]
}

func Lns_2DDD(multi ...LnsAny) []LnsAny {
	if len(multi) == 0 {
		return multi
	}
	switch multi[len(multi)-1].(type) {
	case []LnsAny:
		ddd := multi[len(multi)-1].([]LnsAny)
		newMulti := multi[:len(multi)-1]
		for _, val := range ddd {
			newMulti = append(newMulti, val)
		}
		return newMulti
	}
	return multi
}

func Lns_2DDDGen[T any](multi ...LnsAny) []T {
	if len(multi) == 0 {
		return []T{}
	}
	newMulti := make([]T, len(multi)-1)
	for index, val := range multi[:len(multi)-1] {
		newMulti[index] = val.(T)
	}
	switch multi[len(multi)-1].(type) {
	case []T:
		ddd := multi[len(multi)-1].([]T)
		for _, val := range ddd {
			newMulti = append(newMulti, val)
		}
		return newMulti
	}
	return append(newMulti, multi[len(multi)-1].(T))
}

func Lns_2Slice[T any](slice []LnsAny) []T {
	ret := make([]T, len(slice))
	for index, val := range slice {
		ret[index] = val.(T)
	}
	return ret
}

func Lns_ToString(val LnsAny) string {
	if Lns_IsNil(val) {
		return "nil"
	}
	switch val.(type) {
	case LnsInt:
		return fmt.Sprintf("%d", val)
	case LnsReal:
		real := val.(LnsReal)
		switch lnsRuntimeOpt.Int2strMode {
		case Int2strModeUnneed0:
			return fmt.Sprintf("%g", real)
		case Int2strModeNeed0:
			if digit, frac := math.Modf(real); frac == 0 {
				return fmt.Sprintf("%g.0", digit)
			}
			return fmt.Sprintf("%g", real)
		}
		return lns_ToStringFromReal(val.(LnsReal))
	case bool:
		return fmt.Sprintf("%v", val)
	case string:
		return val.(string)
	default:
		value := reflect.ValueOf(val)
		if value.Kind() == reflect.Func {
			return fmt.Sprintf("function:%T", val)
		}
		if value.CanAddr() {
			return fmt.Sprintf("table:%T(%p)", val, value.Addr())
		}
		return fmt.Sprintf("table:%T(%p)", val, val)
	}
}

/**
 * スタックを一段上げる
 */
func (self *LnsEnv) IncStack() bool {
	self.valStack = append(self.valStack, nil)
	return false
}

/**
 * 値 pVal をスタックの top にセットし、値 pVal の条件判定結果を返す
 *
 * スタックに lns_ddd_t は詰めない。
 * 呼び出し側で lns_ddd_t の先頭要素を指定すること。
 *
 * @param pVal スタックに詰む値
 * @return pVal の条件判定結果。 lns_isCondTrue()。
 */
func (self *LnsEnv) SetStackVal(val LnsAny) bool {
	self.valStack[len(self.valStack)-1] = val
	return Lns_isCondTrue(val)
}

/**
 * スタックから値を pop する。
 *
 * @return pop した値。
 */
func (self *LnsEnv) PopVal(dummy bool) LnsAny {
	val := self.valStack[len(self.valStack)-1]
	self.valStack = self.valStack[:len(self.valStack)-1]
	return val
}

var validRuntimeDebugLog = false

func LnsEnableDebugLog(_env *LnsEnv, valid bool) {
	validRuntimeDebugLog = valid
}

func LnsGetTime(_env *LnsEnv) LnsInt {
	return LnsInt(time.Now().Sub(launchTime).Milliseconds())
}
