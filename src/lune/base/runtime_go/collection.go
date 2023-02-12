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
	"reflect"
)

type Lns_ToMap interface {
	ToMap() *LnsMap
}

type Lns_ToCollectionIF interface {
	ToCollection() LnsAny
}

func Lns_ToCollection(val LnsAny) LnsAny {
	if Lns_IsNil(val) {
		return nil
	}
	switch val.(type) {
	case LnsInt:
		return val
	case LnsReal:
		return val
	case bool:
		return val
	case string:
		return val
	case Lns_ToCollectionIF:
		return val.(Lns_ToCollectionIF).ToCollection()
	default:
		return val.(Lns_ToMap).ToMap()
	}
}

func Lns_FromStemGetAt(obj LnsAny, index LnsAny, nilAccess bool) LnsAny {
	if nilAccess {
		if Lns_IsNil(obj) {
			return nil
		}
	}
	if mapObj, ok := obj.(*LnsMap); ok {
		return mapObj.Items[index]
	}
	return obj.(*LnsList).Items[index.(LnsInt)-1]
}

func Lns_valFromGo(goObj reflect.Value) LnsAny {
	kind := goObj.Kind()
	if kind == reflect.Bool {
		return goObj.Bool()
	}
	if kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64 ||
		kind == reflect.Uintptr {
		return LnsInt(goObj.Int())
	}
	if kind == reflect.Float32 ||
		kind == reflect.Float64 {
		ret := goObj.Float()
		return LnsReal(ret)
	}
	if kind == reflect.String {
		return goObj.String()
	}

	if kind == reflect.Array || kind == reflect.Slice {
		list := make([]LnsAny, goObj.Len())
		for index := 0; index < goObj.Len(); index++ {
			val := Lns_valFromGo(goObj.Index(index))
			list[index] = val
		}
		return NewLnsList(list)
	}
	if kind == reflect.Map {
		mapObj := NewLnsMap(map[LnsAny]LnsAny{})
		iter := goObj.MapRange()
		for iter.Next() {
			key := Lns_valFromGo(iter.Key())
			val := Lns_valFromGo(iter.Value())
			mapObj.Items[key] = val
		}
		return mapObj
	}
	if kind == reflect.Interface {
		return Lns_valFromGo(reflect.ValueOf(goObj.Interface()))
	}
	// kind == reflect.Struct
	return goObj.Interface()
}

func Lns_valToGo(val LnsAny, jsonMode bool) LnsAny {
	if Lns_IsNil(val) {
		return nil
	} else {
		switch val.(type) {
		case LnsInt:
			return val
		case LnsReal:
			return val
		case bool:
			return val
		case string:
			return val
		case *LnsList:
			lnsList := val.(*LnsList)
			list := make([]LnsAny, len(lnsList.Items))
			for index := 0; index < len(lnsList.Items); index++ {
				list[index] = Lns_valToGo(lnsList.Items[index], jsonMode)
			}
			return list
		case *LnsMap:
			lnsMap := val.(*LnsMap)
			if jsonMode {
				mapObj := map[string]LnsAny{}
				for key, val := range lnsMap.Items {
					mapObj[key.(string)] = Lns_valToGo(val, jsonMode)
				}
				return mapObj
			} else {
				mapObj := map[LnsAny]LnsAny{}
				for key, val := range lnsMap.Items {
					mapObj[Lns_valToGo(key, jsonMode)] = Lns_valToGo(val, jsonMode)
				}
				return mapObj
			}
		default:
			return val
		}
	}
}

// ======== list ========

const (
	LnsItemKindUnknown = 0
	LnsItemKindStem    = 1
	LnsItemKindInt     = 2
	LnsItemKindReal    = 3
	LnsItemKindStr     = 4
)

type LnsList struct {
	Items       []LnsAny
	lnsItemKind int
}

func Lns_listFromGo(goObj LnsAny) *LnsList {
	if Lns_IsNil(goObj) {
		return nil
	}
	return Lns_valFromGo(reflect.ValueOf(goObj)).(*LnsList)
}

func Lns_ToListSub(
	obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny) {
	if Lns_IsNil(obj) {
		if nilable {
			return true, nil, nil
		}
		return false, nil, "nil"
	}
	itemParam := paramList[0]
	if val, ok := obj.(*LnsList); ok {
		list := make([]LnsAny, len(val.Items))
		for index, val := range val.Items {
			success, conved, mess :=
				itemParam.Func(val, itemParam.Nilable, itemParam.Child)
			if !success {
				return false, nil, fmt.Sprintf("%d:%s", index+1, mess)
			}
			list[index] = conved
		}
		return true, NewLnsList(list), nil
	} else if val, ok := obj.(*LnsMap); ok {
		// lua は list と map に明確な差がないので、
		// 本来 list のデータも map になる可能性があるため、
		// map からも処理できるようにする。
		list := make([]LnsAny, len(val.Items))
		for index := 1; index <= len(val.Items); index++ {
			if val, ok := val.Items[index]; ok {
				success, conved, mess :=
					itemParam.Func(val, itemParam.Nilable, itemParam.Child)
				if !success {
					return false, nil, fmt.Sprintf("%d:%s", index, mess)
				}
				list[index-1] = conved
			} else {
				return false, nil, fmt.Sprintf("%d:%s", index, "no index")
			}
		}
		return true, NewLnsList(list), nil
	}
	return false, nil, "no list"
}

func (self *LnsList) ToCollection() LnsAny {
	list := make([]LnsAny, len(self.Items))
	for index, val := range self.Items {
		list[index] = Lns_ToCollection(val)
	}
	return NewLnsList(list)
}

func (self *LnsList) Len() int {
	return len(self.Items)
}
func (self *LnsList) Less(idx1, idx2 int) bool {
	val1 := self.Items[idx1]
	val2 := self.Items[idx2]
	switch self.lnsItemKind {
	case LnsItemKindInt:
		return val1.(LnsInt) < val2.(LnsInt)
	case LnsItemKindReal:
		return val1.(LnsReal) < val2.(LnsReal)
	case LnsItemKindStr:
		return val1.(string) < val2.(string)
	case LnsItemKindStem:
		switch val1.(type) {
		case LnsInt:
			cval1 := val1.(LnsInt)
			switch val2.(type) {
			case LnsInt:
				return cval1 < val2.(LnsInt)
			case LnsReal:
				return LnsReal(cval1) < val2.(LnsReal)
			default:
				return true
			}
		case LnsReal:
			cval1 := val1.(LnsReal)
			switch val2.(type) {
			case LnsInt:
				return cval1 < LnsReal(val2.(LnsInt))
			case LnsReal:
				return cval1 < val2.(LnsReal)
			default:
				return true
			}
		case string:
			cval1 := val1.(string)
			switch val2.(type) {
			case LnsInt:
				return false
			case LnsReal:
				return false
			case string:
				cval2 := val2.(string)
				return cval1 < cval2
			default:
				return true
			}
		default:
			switch val2.(type) {
			case LnsInt:
				return false
			case LnsReal:
				return false
			case string:
				return false
			default:
				return idx1 < idx2
			}
		}
	}
	panic("error")
	return false
}
func (self *LnsList) Swap(idx1, idx2 int) {
	self.Items[idx1], self.Items[idx2] = self.Items[idx2], self.Items[idx1]
}

func NewLnsList(list []LnsAny) *LnsList {
	return &LnsList{list, LnsItemKindUnknown}
}
func (lnsList *LnsList) Insert(val LnsAny) {
	if !Lns_IsNil(val) {
		lnsList.Items = append(lnsList.Items, val)
	}
}
func (lnsList *LnsList) Remove(index LnsAny) LnsAny {
	if Lns_IsNil(index) {
		ret := lnsList.Items[len(lnsList.Items)-1]
		lnsList.Items = lnsList.Items[:len(lnsList.Items)-1]
		return ret
	} else {
		work := index.(LnsInt) - 1
		ret := lnsList.Items[work]
		lnsList.Items =
			append(lnsList.Items[:work], lnsList.Items[work+1:]...)
		return ret
	}
}
func (lnsList *LnsList) GetAt(index int) LnsAny {
	return lnsList.Items[index-1]
}
func (lnsList *LnsList) Set(index int, val LnsAny) {
	index--
	if len(lnsList.Items) > index {
		lnsList.Items[index] = val
	} else {
		if len(lnsList.Items) == index {
			lnsList.Items = append(lnsList.Items, val)
		} else {
			panic(fmt.Sprintf("illegal index -- %d", index))
		}
	}
}
func (lnsList *LnsList) Unpack() []LnsAny {
	return lnsList.Items
}

func (LnsList *LnsList) ToLuaCode(conv *StemToLuaConv) {
	conv.write("{")
	for _, val := range LnsList.Items {
		conv.conv(val)
		conv.write(",")
	}
	conv.write("}")
}

// ======== GenList =======

type LnsList2_[T any] struct {
	Items       []T
	lnsItemKind int
}

func Lns_ToList2Sub[T any](
	obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny) {
	if Lns_IsNil(obj) {
		if nilable {
			return true, nil, nil
		}
		return false, nil, "nil"
	}
	itemParam := paramList[0]
	if val, ok := obj.(*LnsList); ok {
		list := make([]T, len(val.Items))
		for index, val := range val.Items {
			success, conved, mess :=
				itemParam.Func(val, itemParam.Nilable, itemParam.Child)
			if !success {
				return false, nil, fmt.Sprintf("%d:%s", index+1, mess)
			}
			list[index] = conved.(T)
		}
		return true, NewLnsList2_[T](list), nil
	} else if val, ok := obj.(*LnsMap); ok {
		// lua は list と map に明確な差がないので、
		// 本来 list のデータも map になる可能性があるため、
		// map からも処理できるようにする。
		list := make([]T, len(val.Items))
		for index := 1; index <= len(val.Items); index++ {
			if val, ok := val.Items[index]; ok {
				success, conved, mess :=
					itemParam.Func(val, itemParam.Nilable, itemParam.Child)
				if !success {
					return false, nil, fmt.Sprintf("%d:%s", index, mess)
				}
				list[index-1] = conved.(T)
			} else {
				return false, nil, fmt.Sprintf("%d:%s", index, "no index")
			}
		}
		return true, NewLnsList2_[T](list), nil
	}
	return false, nil, "no list"
}

func (self *LnsList2_[T]) ToCollection() LnsAny {
	list := make([]LnsAny, len(self.Items))
	for index, val := range self.Items {
		list[index] = Lns_ToCollection(val)
	}
	return NewLnsList(list)
}

func (self *LnsList2_[T]) Len() int {
	return len(self.Items)
}
func (self *LnsList2_[T]) Less(idx1, idx2 int) bool {
	var val1 any = self.Items[idx1]
	var val2 any = self.Items[idx2]
	switch self.lnsItemKind {
	case LnsItemKindInt:
		return val1.(LnsInt) < val2.(LnsInt)
	case LnsItemKindReal:
		return val1.(LnsReal) < val2.(LnsReal)
	case LnsItemKindStr:
		return val1.(string) < val2.(string)
	case LnsItemKindStem:
		switch val1.(type) {
		case LnsInt:
			cval1 := val1.(LnsInt)
			switch val2.(type) {
			case LnsInt:
				return cval1 < val2.(LnsInt)
			case LnsReal:
				return LnsReal(cval1) < val2.(LnsReal)
			default:
				return true
			}
		case LnsReal:
			cval1 := val1.(LnsReal)
			switch val2.(type) {
			case LnsInt:
				return cval1 < LnsReal(val2.(LnsInt))
			case LnsReal:
				return cval1 < val2.(LnsReal)
			default:
				return true
			}
		case string:
			cval1 := val1.(string)
			switch val2.(type) {
			case LnsInt:
				return false
			case LnsReal:
				return false
			case string:
				cval2 := val2.(string)
				return cval1 < cval2
			default:
				return true
			}
		default:
			switch val2.(type) {
			case LnsInt:
				return false
			case LnsReal:
				return false
			case string:
				return false
			default:
				return idx1 < idx2
			}
		}
	}
	panic("error")
	return false
}
func (self *LnsList2_[T]) Swap(idx1, idx2 int) {
	self.Items[idx1], self.Items[idx2] = self.Items[idx2], self.Items[idx1]
}

func NewLnsList2_[T any](list []T) *LnsList2_[T] {
	return &LnsList2_[T]{list, LnsItemKindUnknown}
}
func (lnsList *LnsList2_[T]) Insert(val T) {
	if !Lns_IsNil(val) {
		lnsList.Items = append(lnsList.Items, val)
	}
}
func (lnsList *LnsList2_[T]) Remove(index LnsAny) T {
	if Lns_IsNil(index) {
		ret := lnsList.Items[len(lnsList.Items)-1]
		lnsList.Items = lnsList.Items[:len(lnsList.Items)-1]
		return ret
	} else {
		work := index.(LnsInt) - 1
		ret := lnsList.Items[work]
		lnsList.Items =
			append(lnsList.Items[:work], lnsList.Items[work+1:]...)
		return ret
	}
}
func (lnsList *LnsList2_[T]) GetAt(index int) T {
	return lnsList.Items[index-1]
}
func (lnsList *LnsList2_[T]) Set(index int, val T) {
	index--
	if len(lnsList.Items) > index {
		lnsList.Items[index] = val
	} else {
		if len(lnsList.Items) == index {
			lnsList.Items = append(lnsList.Items, val)
		} else {
			panic(fmt.Sprintf("illegal index -- %d", index))
		}
	}
}
func (lnsList *LnsList2_[T]) Unpack() []T {
	ret := make([]T, len(lnsList.Items))
	for index := 0; index < len(ret); index++ {
		ret[index] = lnsList.Items[index]
	}
	return ret
}

func (LnsList2_ *LnsList2_[T]) ToLuaCode(conv *StemToLuaConv) {
	conv.write("{")
	for _, val := range LnsList2_.Items {
		conv.conv(val)
		conv.write(",")
	}
	conv.write("}")
}

// ======== set ========

type LnsSet struct {
	Items map[LnsAny]bool
}

func Lns_ToSetSub(
	obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny) {
	if Lns_IsNil(obj) {
		if nilable {
			return true, nil, nil
		}
		return false, nil, "nil"
	}
	itemParam := paramList[0]
	if val, ok := obj.(*LnsSet); ok {
		list := make([]LnsAny, len(val.Items))
		index := 0
		for key := range val.Items {
			success, conved, mess :=
				itemParam.Func(key, itemParam.Nilable, itemParam.Child)
			if !success {
				return false, nil, fmt.Sprintf("%s:%s", key, mess)
			}
			list[index] = conved
			index++
		}
		return true, NewLnsSet(list), nil
	}
	return false, nil, "no set"
}

func (self *LnsSet) ToCollection() LnsAny {
	ret := NewLnsSet([]LnsAny{})
	for key := range self.Items {
		ret.Add(Lns_ToCollection(key))
	}
	return ret
}

func (self *LnsSet) CreateKeyListStem() *LnsList {
	list := make([]LnsAny, len(self.Items))
	index := 0
	for key := range self.Items {
		list[index] = key
		index++
	}
	return NewLnsList(list)
}
func (self *LnsSet) CreateKeyListInt() *LnsList {
	list := self.CreateKeyListStem()
	list.lnsItemKind = LnsItemKindInt
	return list
}

func (self *LnsSet) CreateKeyListReal() *LnsList {
	list := self.CreateKeyListStem()
	list.lnsItemKind = LnsItemKindReal
	return list
}
func (self *LnsSet) CreateKeyListStr() *LnsList {
	list := self.CreateKeyListStem()
	list.lnsItemKind = LnsItemKindStr
	return list
}

func NewLnsSet(list []LnsAny) *LnsSet {
	set := &LnsSet{map[LnsAny]bool{}}
	for _, val := range list {
		set.Items[val] = true
	}
	return set
}

func (self *LnsSet) Add(val LnsAny) {
	self.Items[val] = true
}
func (self *LnsSet) Del(val LnsAny) {
	delete(self.Items, val)
}
func (self *LnsSet) Has(val LnsAny) bool {
	_, has := self.Items[val]
	return has
}
func (self *LnsSet) And(set *LnsSet) *LnsSet {
	delValList := NewLnsList([]LnsAny{})
	for val := range self.Items {
		if !set.Has(val) {
			delValList.Insert(val)
		}
	}
	for _, val := range delValList.Items {
		delete(self.Items, val)
	}
	return self
}
func (self *LnsSet) Or(set *LnsSet) *LnsSet {
	for val := range set.Items {
		self.Items[val] = true
	}
	return self
}
func (self *LnsSet) Sub(set *LnsSet) *LnsSet {
	delValList := NewLnsList([]LnsAny{})
	for val := range set.Items {
		if set.Has(val) {
			delValList.Insert(val)
		}
	}
	for _, val := range delValList.Items {
		delete(self.Items, val)
	}
	return self
}
func (self *LnsSet) Clone() *LnsSet {
	set := NewLnsSet([]LnsAny{})
	for val := range self.Items {
		set.Items[val] = true
	}
	return set
}
func (self *LnsSet) Len() LnsInt {
	return len(self.Items)
}

// ======== map ========

type LnsMap struct {
	Items map[LnsAny]LnsAny
}

func Lns_mapFromGo(goObj LnsAny) *LnsMap {
	if Lns_IsNil(goObj) {
		return nil
	}
	return Lns_valFromGo(reflect.ValueOf(goObj)).(*LnsMap)
}

func Lns_ToLnsMapSub(
	obj LnsAny, nilable bool, paramList []Lns_ToObjParam) (bool, LnsAny, LnsAny) {
	if Lns_IsNil(obj) {
		if nilable {
			return true, nil, nil
		}
		return false, nil, "nil"
	}
	keyParam := paramList[0]
	itemParam := paramList[1]
	if lnsMap, ok := obj.(*LnsMap); ok {
		newMap := NewLnsMap(map[LnsAny]LnsAny{})
		for key, val := range lnsMap.Items {
			successKey, convedKey, messKey :=
				keyParam.Func(key, keyParam.Nilable, keyParam.Child)
			if !successKey {
				return false, nil, fmt.Sprintf(".%s:%s", key, messKey)
			}
			successVal, convedVal, messVal :=
				itemParam.Func(val, itemParam.Nilable, itemParam.Child)
			if !successVal {
				return false, nil, fmt.Sprintf(".%s:%s", val, messVal)
			}
			newMap.Items[convedKey] = convedVal
		}
		return true, newMap, nil
	}
	return false, nil, "no map"
}

func (self *LnsMap) ToCollection() LnsAny {
	ret := NewLnsMap(map[LnsAny]LnsAny{})
	for key, val := range self.Items {
		ret.Items[key] = Lns_ToCollection(val)
	}
	return ret
}

func (self *LnsMap) Correct() *LnsMap {
	delete(self.Items, nil)
	list := make([]LnsAny, len(self.Items))
	index := 0
	for key, val := range self.Items {
		if Lns_IsNil(val) {
			list[index] = key
			index++
		}
	}
	for _, key := range list[:index] {
		delete(self.Items, key)
	}
	return self
}

func (self *LnsMap) CreateKeyListStem() *LnsList {
	list := make([]LnsAny, len(self.Items))
	index := 0
	for key := range self.Items {
		list[index] = key
		index++
	}
	return NewLnsList(list)
}
func (self *LnsMap) CreateKeyListInt() *LnsList {
	list := self.CreateKeyListStem()
	list.lnsItemKind = LnsItemKindInt
	return list
}

func (self *LnsMap) CreateKeyListReal() *LnsList {
	list := self.CreateKeyListStem()
	list.lnsItemKind = LnsItemKindReal
	return list
}
func (self *LnsMap) CreateKeyListStr() *LnsList {
	list := self.CreateKeyListStem()
	list.lnsItemKind = LnsItemKindStr
	return list
}

func NewLnsMap(arg map[LnsAny]LnsAny) *LnsMap {
	return &LnsMap{arg}
}

func (self *LnsMap) Set(key, val LnsAny) {
	if Lns_IsNil(val) {
		delete(self.Items, key)
	} else {
		self.Items[key] = val
	}
}

func (self *LnsMap) Del(key LnsAny) {
	delete(self.Items, key)
}

func (self *LnsMap) Get(key LnsAny) LnsAny {
	val, has := self.Items[key]
	if !has {
		return nil
	}
	return val
}

func (LnsMap *LnsMap) ToLuaCode(conv *StemToLuaConv) {
	conv.write("{")
	for key, val := range LnsMap.Items {
		conv.write("[ ")
		conv.conv(key)
		conv.write(" ] =")
		conv.conv(val)
		conv.write(",")
	}
	conv.write("}")
}
