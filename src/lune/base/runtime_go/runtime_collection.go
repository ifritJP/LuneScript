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

package main


//import "fmt"
import "log"
import "sort"

type Lns_ToMap interface {
    ToMap() LnsMap
}

type Lns_ToCollectionIF interface {
    ToCollection() LnsAny
}

func Lns_ToCollection( val LnsAny ) LnsAny {
    if Lns_IsNil( val ) {
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

// ======== list ========

const (
    LnsItemKindStem = 0
    LnsItemKindInt = 1
    LnsItemKindReal = 2
    LnsItemKindStr = 3
)

type LnsList struct {
    Items []LnsAny
    lnsItemKind int
}

func (self *LnsList) ToCollection() LnsAny {
    list := make([]LnsAny, len(self.Items))
    for index, val := range (self.Items) {
        list[ index ] = Lns_ToCollection( val )
    }
    return NewLnsList( list )
}

func (self *LnsList) Sort() {
    sort.Sort( self )
}

func (self *LnsList) Len() int {
    return len(self.Items)
}
func (self *LnsList) Less(idx1, idx2 int) bool {
    switch self.lnsItemKind {
    case LnsItemKindInt:
        return self.Items[idx1].(LnsInt) < self.Items[idx2].(LnsInt)
    case LnsItemKindReal:
        return self.Items[idx1].(LnsReal) < self.Items[idx2].(LnsReal)
    case LnsItemKindStr:
        return self.Items[idx1].(string) < self.Items[idx2].(string)
    case LnsItemKindStem:
        return idx1 < idx2
    }
    log.Fatal( "error" )
    return false
}
func (self *LnsList) Swap(idx1, idx2 int) {
    self.Items[ idx1 ], self.Items[ idx2 ] = self.Items[ idx2 ], self.Items[ idx1 ]
}


func NewLnsList( list []LnsAny ) *LnsList {
    return &LnsList{ list, LnsItemKindStem }
}
func (lnsList *LnsList) Insert( val LnsAny ) {
    if !Lns_IsNil( val ) {
        lnsList.Items = append( lnsList.Items, val )
    }
}
func (lnsList *LnsList) Remove( index LnsAny ) LnsAny {
    if Lns_IsNil( index ) {
        ret := lnsList.Items[ len(lnsList.Items) - 1 ]
        lnsList.Items = lnsList.Items[ : len(lnsList.Items) - 1 ]
        return ret
    } else {
        work := index.(LnsInt)
        ret := lnsList.Items[ work ]
        lnsList.Items =
            append( lnsList.Items[ : work ], lnsList.Items[ work+1: ]... )
        return ret
    }
}
func (lnsList *LnsList) GetAt( index int ) LnsAny {
    return lnsList.Items[ index - 1 ]
}



// ======== set ========

type LnsSet struct {
    Items map[LnsAny]bool
}

func (self *LnsSet) ToCollection() LnsAny {
    ret := NewLnsSet([]LnsAny{})
    for key := range (self.Items) {
        ret.Add( Lns_ToCollection( key ) )
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
    return NewLnsList( list )
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



func NewLnsSet( list []LnsAny ) *LnsSet {
    set := &LnsSet{ map[LnsAny]bool{} }
    for _, val := range( list ) {
        set.Items[ val ] = true;
    }
    return set
}

func (self *LnsSet) Add( val LnsAny ) {
    self.Items[ val ] = true;
}
func (self *LnsSet) Del( val LnsAny ) {
    delete( self.Items, val )
}
func (self *LnsSet) Has( val LnsAny ) bool {
    _, has := self.Items[ val ]
    return has
}
func (self *LnsSet) And( set *LnsSet ) *LnsSet {
    delValList := NewLnsList( []LnsAny{} )
    for val := range( self.Items ) {
        if !set.Has( val ) {
            delValList.Insert( val )
        }
    }
    for _, val := range( delValList.Items ) {
        delete( self.Items, val )
    }
    return self
}
func (self *LnsSet) Or( set *LnsSet ) *LnsSet {
    for val := range( set.Items ) {
        self.Items[ val ] = true
    }
    return self
}
func (self *LnsSet) Sub( set *LnsSet ) *LnsSet {
    delValList := NewLnsList( []LnsAny{} )
    for val := range( set.Items ) {
        if set.Has( val ) {
            delValList.Insert( val )
        }
    }
    for _, val := range( delValList.Items ) {
        delete( self.Items, val )
    }
    return self
}
func (self *LnsSet) Clone() *LnsSet {
    set := NewLnsSet( []LnsAny{} )
    for val := range( set.Items ) {
        set.Items[ val ] = true
    }
    return set
}
func (self *LnsSet) Len() LnsInt {
    return len( self.Items )
}

// ======== map ========

type LnsMap map[LnsAny]LnsAny

func (self LnsMap) ToCollection() LnsAny {
    ret := LnsMap{}
    for key, val := range (self) {
        ret[ key ] = Lns_ToCollection( val )
    }
    return ret
}


func (self LnsMap) Correct() LnsMap {
    delete( self, nil )
    list := make([]LnsAny, len(self))
    index := 0
    for key, val := range self {
        if Lns_IsNil( val ) {
            list[index] = key
            index++
        }
    }
    for _, key := range list[:index] {
        delete( self, key )
    }
    return self
}

func (self LnsMap) CreateKeyListStem() *LnsList {
    list := make([]LnsAny, len(self))
    index := 0
    for key := range self {
        list[index] = key
        index++
    }    
    return NewLnsList( list )
}
func (self LnsMap) CreateKeyListInt() *LnsList {
    list := self.CreateKeyListStem()
    list.lnsItemKind = LnsItemKindInt
    return list
}

func (self LnsMap) CreateKeyListReal() *LnsList {
    list := self.CreateKeyListStem()
    list.lnsItemKind = LnsItemKindReal
    return list
}
func (self LnsMap) CreateKeyListStr() *LnsList {
    list := self.CreateKeyListStem()
    list.lnsItemKind = LnsItemKindStr
    return list
}

func (self LnsMap) Set( key, val LnsAny ) {
    if Lns_IsNil( val ) {
        delete( self, key );
    } else {
        self[ key ] = val
    }
}
