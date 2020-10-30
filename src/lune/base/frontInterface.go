// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_frontInterface bool
var frontInterface__mod__ string
var FrontInterface___luneScript FrontInterface_frontInterface
// for 135
func frontInterface_convExp350(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 162: decl @lune.@base.@frontInterface.setFront
func FrontInterface_setFront(newFront FrontInterface_frontInterface) {
    FrontInterface___luneScript = newFront
    
}

// 166: decl @lune.@base.@frontInterface.loadModule
func FrontInterface_loadModule(mod string)(LnsAny, LnsAny) {
    return FrontInterface___luneScript.LoadModule(mod)
}

// 170: decl @lune.@base.@frontInterface.loadFromLnsTxt
func FrontInterface_loadFromLnsTxt(importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    return FrontInterface___luneScript.LoadFromLnsTxt(importModuleInfo, name, txt)
}

// 176: decl @lune.@base.@frontInterface.loadMeta
func FrontInterface_loadMeta(importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    return FrontInterface___luneScript.LoadMeta(importModuleInfo, mod)
}

// 180: decl @lune.@base.@frontInterface.searchModule
func FrontInterface_searchModule(mod string) LnsAny {
    return FrontInterface___luneScript.SearchModule(mod)
}

// declaration Class -- ModuleId
var FrontInterface_ModuleId__tempId *FrontInterface_ModuleId
// 32: decl @lune.@base.@frontInterface.ModuleId.___init
func FrontInterface_ModuleId____init_1021_() {
    FrontInterface_ModuleId__tempId = NewFrontInterface_ModuleId(0.0, 0)
    
}

type FrontInterface_ModuleIdMtd interface {
    GetNextModuleId() *FrontInterface_ModuleId
    Get_buildCount() LnsInt
    Get_idStr() string
    Get_modTime() LnsReal
}
type FrontInterface_ModuleId struct {
    modTime LnsReal
    buildCount LnsInt
    idStr string
    TempId *FrontInterface_ModuleId
    FP FrontInterface_ModuleIdMtd
}
func FrontInterface_ModuleId2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_ModuleId).FP
}
type FrontInterface_ModuleIdDownCast interface {
    ToFrontInterface_ModuleId() *FrontInterface_ModuleId
}
func FrontInterface_ModuleIdDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(FrontInterface_ModuleIdDownCast)
    if ok { return work.ToFrontInterface_ModuleId() }
    return nil
}
func (obj *FrontInterface_ModuleId) ToFrontInterface_ModuleId() *FrontInterface_ModuleId {
    return obj
}
func NewFrontInterface_ModuleId(arg1 LnsReal, arg2 LnsInt) *FrontInterface_ModuleId {
    obj := &FrontInterface_ModuleId{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleId(arg1, arg2)
    return obj
}
func (self *FrontInterface_ModuleId) Get_modTime() LnsReal{ return self.modTime }
func (self *FrontInterface_ModuleId) Get_buildCount() LnsInt{ return self.buildCount }
func (self *FrontInterface_ModuleId) Get_idStr() string{ return self.idStr }
// 46: DeclConstr
func (self *FrontInterface_ModuleId) InitFrontInterface_ModuleId(modTime LnsReal,buildCount LnsInt) {
    self.modTime = modTime
    
    self.buildCount = buildCount
    
    self.idStr = Lns_getVM().String_format("%f:%d", []LnsAny{modTime, buildCount})
    
}

// 52: decl @lune.@base.@frontInterface.ModuleId.getNextModuleId
func (self *FrontInterface_ModuleId) GetNextModuleId() *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(self.modTime, self.buildCount + 1)
}

// 59: decl @lune.@base.@frontInterface.ModuleId.createId
func FrontInterface_ModuleId_createId(modTime LnsReal,buildCount LnsInt) *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(modTime, buildCount)
}

// 62: decl @lune.@base.@frontInterface.ModuleId.createIdFromTxt
func FrontInterface_ModuleId_createIdFromTxt(idStr string) *FrontInterface_ModuleId {
    var modTime LnsReal
    modTime = Lns_unwrapDefault( Lns_tonumber(Lns_car(Lns_getVM().String_gsub(idStr,":.*", "")).(string), nil), 0.0).(LnsReal)
    var buildCount LnsReal
    buildCount = Lns_unwrapDefault( Lns_tonumber(Lns_car(Lns_getVM().String_gsub(idStr,".*:", "")).(string), nil), 0.0).(LnsReal)
    return NewFrontInterface_ModuleId(modTime, (LnsInt)(buildCount))
}


// declaration Class -- ImportModuleInfo
type FrontInterface_ImportModuleInfoMtd interface {
    Add(arg1 string) bool
    GetFull() string
    Remove()
}
type FrontInterface_ImportModuleInfo struct {
    orderedSet *Util_OrderedSet
    FP FrontInterface_ImportModuleInfoMtd
}
func FrontInterface_ImportModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_ImportModuleInfo).FP
}
type FrontInterface_ImportModuleInfoDownCast interface {
    ToFrontInterface_ImportModuleInfo() *FrontInterface_ImportModuleInfo
}
func FrontInterface_ImportModuleInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(FrontInterface_ImportModuleInfoDownCast)
    if ok { return work.ToFrontInterface_ImportModuleInfo() }
    return nil
}
func (obj *FrontInterface_ImportModuleInfo) ToFrontInterface_ImportModuleInfo() *FrontInterface_ImportModuleInfo {
    return obj
}
func NewFrontInterface_ImportModuleInfo() *FrontInterface_ImportModuleInfo {
    obj := &FrontInterface_ImportModuleInfo{}
    obj.FP = obj
    obj.InitFrontInterface_ImportModuleInfo()
    return obj
}
// 76: DeclConstr
func (self *FrontInterface_ImportModuleInfo) InitFrontInterface_ImportModuleInfo() {
    self.orderedSet = NewUtil_OrderedSet()
    
}

// 80: decl @lune.@base.@frontInterface.ImportModuleInfo.add
func (self *FrontInterface_ImportModuleInfo) Add(modulePath string) bool {
    return self.orderedSet.FP.Add(modulePath)
}

// 84: decl @lune.@base.@frontInterface.ImportModuleInfo.remove
func (self *FrontInterface_ImportModuleInfo) Remove() {
    self.orderedSet.FP.RemoveLast()
}

// 88: decl @lune.@base.@frontInterface.ImportModuleInfo.getFull
func (self *FrontInterface_ImportModuleInfo) GetFull() string {
    var txt string
    txt = ""
    for _, _modulePath := range( self.orderedSet.FP.Get_list().Items ) {
        modulePath := _modulePath.(string)
        txt = Lns_getVM().String_format("%s -> %s", []LnsAny{txt, modulePath})
        
    }
    return txt
}


type FrontInterface_frontInterface interface {
        Error(arg1 string)
        LoadFromLnsTxt(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string) LnsAny
        LoadMeta(arg1 *FrontInterface_ImportModuleInfo, arg2 string) LnsAny
        LoadModule(arg1 string)(LnsAny, LnsAny)
        SearchModule(arg1 string) LnsAny
}
func Lns_cast2FrontInterface_frontInterface( obj LnsAny ) LnsAny {
    if _, ok := obj.(FrontInterface_frontInterface); ok { 
        return obj
    }
    return nil
}

// declaration Class -- dummyFront
type frontInterface_dummyFrontMtd interface {
    Error(arg1 string)
    LoadFromLnsTxt(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string) LnsAny
    LoadMeta(arg1 *FrontInterface_ImportModuleInfo, arg2 string) LnsAny
    LoadModule(arg1 string)(LnsAny, LnsAny)
    SearchModule(arg1 string) LnsAny
}
type frontInterface_dummyFront struct {
    FP frontInterface_dummyFrontMtd
}
func frontInterface_dummyFront2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*frontInterface_dummyFront).FP
}
type frontInterface_dummyFrontDownCast interface {
    TofrontInterface_dummyFront() *frontInterface_dummyFront
}
func frontInterface_dummyFrontDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(frontInterface_dummyFrontDownCast)
    if ok { return work.TofrontInterface_dummyFront() }
    return nil
}
func (obj *frontInterface_dummyFront) TofrontInterface_dummyFront() *frontInterface_dummyFront {
    return obj
}
func NewfrontInterface_dummyFront() *frontInterface_dummyFront {
    obj := &frontInterface_dummyFront{}
    obj.FP = obj
    obj.InitfrontInterface_dummyFront()
    return obj
}
func (self *frontInterface_dummyFront) InitfrontInterface_dummyFront() {
}
// 133: decl @lune.@base.@frontInterface.dummyFront.loadModule
func (self *frontInterface_dummyFront) LoadModule(mod string)(LnsAny, LnsAny) {
    var loaded LnsAny
    loaded = frontInterface_convExp350(Lns_2DDD(Lns_getVM().Load("return {}", nil)))
    var emptyTable LnsAny
    if loaded != nil{
        loaded_224 := loaded.(*Lns_luaValue)
        emptyTable = Lns_unwrap( Lns_car(Lns_getVM().RunLoadedfunc(loaded_224,Lns_2DDD([]LnsAny{}))[0]))
        
    } else {
        panic("load error")
    }
    return Lns_require(mod), emptyTable
}

// 144: decl @lune.@base.@frontInterface.dummyFront.loadMeta
func (self *frontInterface_dummyFront) LoadMeta(importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 147: decl @lune.@base.@frontInterface.dummyFront.loadFromLnsTxt
func (self *frontInterface_dummyFront) LoadFromLnsTxt(importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 152: decl @lune.@base.@frontInterface.dummyFront.searchModule
func (self *frontInterface_dummyFront) SearchModule(mod string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 155: decl @lune.@base.@frontInterface.dummyFront.error
func (self *frontInterface_dummyFront) Error(message string) {
    panic("not implements")
}


func Lns_frontInterface_init() {
    if init_frontInterface { return }
    init_frontInterface = true
    frontInterface__mod__ = "@lune.@base.@frontInterface"
    Lns_InitMod()
    Lns_Util_init()
    FrontInterface_ModuleId____init_1021_()
    FrontInterface___luneScript = NewfrontInterface_dummyFront().FP
}
func init() {
    init_frontInterface = false
}
