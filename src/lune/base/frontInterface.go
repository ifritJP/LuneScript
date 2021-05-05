// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_frontInterface bool
var frontInterface__mod__ string
var FrontInterface___luneScript FrontInterface_frontInterface
// for 150
func frontInterface_convExp362(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 181: decl @lune.@base.@frontInterface.setFront
func FrontInterface_setFront(newFront FrontInterface_frontInterface) {
    FrontInterface___luneScript = newFront
    
}

// 185: decl @lune.@base.@frontInterface.loadModule
func FrontInterface_loadModule(mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    return FrontInterface___luneScript.LoadModule(mod)
}

// 189: decl @lune.@base.@frontInterface.loadFromLnsTxt
func FrontInterface_loadFromLnsTxt(importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    return FrontInterface___luneScript.LoadFromLnsTxt(importModuleInfo, name, txt)
}

// 195: decl @lune.@base.@frontInterface.loadMeta
func FrontInterface_loadMeta(importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    return FrontInterface___luneScript.LoadMeta(importModuleInfo, mod)
}

// 199: decl @lune.@base.@frontInterface.searchModule
func FrontInterface_searchModule(mod string) LnsAny {
    return FrontInterface___luneScript.SearchModule(mod)
}

// 203: decl @lune.@base.@frontInterface.getLuaModulePath
func FrontInterface_getLuaModulePath(mod string) string {
    return FrontInterface___luneScript.GetLuaModulePath(mod)
}

// declaration Class -- ModuleMeta
type FrontInterface_ModuleMetaMtd interface {
    Get_lnsPath() string
    Get_metaInfo() LnsAny
}
type FrontInterface_ModuleMeta struct {
    metaInfo LnsAny
    lnsPath string
    FP FrontInterface_ModuleMetaMtd
}
func FrontInterface_ModuleMeta2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_ModuleMeta).FP
}
type FrontInterface_ModuleMetaDownCast interface {
    ToFrontInterface_ModuleMeta() *FrontInterface_ModuleMeta
}
func FrontInterface_ModuleMetaDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(FrontInterface_ModuleMetaDownCast)
    if ok { return work.ToFrontInterface_ModuleMeta() }
    return nil
}
func (obj *FrontInterface_ModuleMeta) ToFrontInterface_ModuleMeta() *FrontInterface_ModuleMeta {
    return obj
}
func NewFrontInterface_ModuleMeta(arg1 LnsAny, arg2 string) *FrontInterface_ModuleMeta {
    obj := &FrontInterface_ModuleMeta{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleMeta(arg1, arg2)
    return obj
}
func (self *FrontInterface_ModuleMeta) InitFrontInterface_ModuleMeta(arg1 LnsAny, arg2 string) {
    self.metaInfo = arg1
    self.lnsPath = arg2
}
func (self *FrontInterface_ModuleMeta) Get_metaInfo() LnsAny{ return self.metaInfo }
func (self *FrontInterface_ModuleMeta) Get_lnsPath() string{ return self.lnsPath }

// declaration Class -- ModuleId
var FrontInterface_ModuleId__tempId *FrontInterface_ModuleId
// 38: decl @lune.@base.@frontInterface.ModuleId.___init
func FrontInterface_ModuleId____init_1046_() {
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
// 52: DeclConstr
func (self *FrontInterface_ModuleId) InitFrontInterface_ModuleId(modTime LnsReal,buildCount LnsInt) {
    self.modTime = modTime
    
    self.buildCount = buildCount
    
    self.idStr = Lns_getVM().String_format("%f:%d", []LnsAny{modTime, buildCount})
    
}

// 58: decl @lune.@base.@frontInterface.ModuleId.getNextModuleId
func (self *FrontInterface_ModuleId) GetNextModuleId() *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(self.modTime, self.buildCount + 1)
}

// 65: decl @lune.@base.@frontInterface.ModuleId.createId
func FrontInterface_ModuleId_createId(modTime LnsReal,buildCount LnsInt) *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(modTime, buildCount)
}

// 68: decl @lune.@base.@frontInterface.ModuleId.createIdFromTxt
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
// 82: DeclConstr
func (self *FrontInterface_ImportModuleInfo) InitFrontInterface_ImportModuleInfo() {
    self.orderedSet = NewUtil_OrderedSet()
    
}

// 86: decl @lune.@base.@frontInterface.ImportModuleInfo.add
func (self *FrontInterface_ImportModuleInfo) Add(modulePath string) bool {
    return self.orderedSet.FP.Add(modulePath)
}

// 90: decl @lune.@base.@frontInterface.ImportModuleInfo.remove
func (self *FrontInterface_ImportModuleInfo) Remove() {
    self.orderedSet.FP.RemoveLast()
}

// 94: decl @lune.@base.@frontInterface.ImportModuleInfo.getFull
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
        GetLuaModulePath(arg1 string) string
        LoadFromLnsTxt(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string) LnsAny
        LoadMeta(arg1 *FrontInterface_ImportModuleInfo, arg2 string) LnsAny
        LoadModule(arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
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
    GetLuaModulePath(arg1 string) string
    LoadFromLnsTxt(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string) LnsAny
    LoadMeta(arg1 *FrontInterface_ImportModuleInfo, arg2 string) LnsAny
    LoadModule(arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
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
// 148: decl @lune.@base.@frontInterface.dummyFront.loadModule
func (self *frontInterface_dummyFront) LoadModule(mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    var loaded LnsAny
    loaded = frontInterface_convExp362(Lns_2DDD(Lns_getVM().Load("return {}", nil)))
    var emptyTable LnsAny
    if loaded != nil{
        loaded_1499 := loaded.(*Lns_luaValue)
        emptyTable = Lns_unwrap( Lns_car(Lns_getVM().RunLoadedfunc(loaded_1499,Lns_2DDD([]LnsAny{}))[0]))
        
    } else {
        panic("load error")
    }
    var meta *FrontInterface_ModuleMeta
    meta = NewFrontInterface_ModuleMeta(emptyTable, Lns_car(Lns_getVM().String_gsub(mod,"%.", "/")).(string) + ".lns")
    return Lns_require(mod), meta
}

// 160: decl @lune.@base.@frontInterface.dummyFront.loadMeta
func (self *frontInterface_dummyFront) LoadMeta(importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 163: decl @lune.@base.@frontInterface.dummyFront.loadFromLnsTxt
func (self *frontInterface_dummyFront) LoadFromLnsTxt(importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 168: decl @lune.@base.@frontInterface.dummyFront.getLuaModulePath
func (self *frontInterface_dummyFront) GetLuaModulePath(mod string) string {
    panic("not implements")
// insert a dummy
    return ""
}

// 171: decl @lune.@base.@frontInterface.dummyFront.searchModule
func (self *frontInterface_dummyFront) SearchModule(mod string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 174: decl @lune.@base.@frontInterface.dummyFront.error
func (self *frontInterface_dummyFront) Error(message string) {
    panic("not implements")
}


func Lns_frontInterface_init() {
    if init_frontInterface { return }
    init_frontInterface = true
    frontInterface__mod__ = "@lune.@base.@frontInterface"
    Lns_InitMod()
    Lns_Util_init()
    Lns_Ast_init()
    FrontInterface_ModuleId____init_1046_()
    FrontInterface___luneScript = NewfrontInterface_dummyFront().FP
}
func init() {
    init_frontInterface = false
}
