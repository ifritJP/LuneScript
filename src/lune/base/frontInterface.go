// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_frontInterface bool
var frontInterface__mod__ string
var FrontInterface___luneScript FrontInterface_frontInterface
// for 228
func frontInterface_convExp657(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 259: decl @lune.@base.@frontInterface.setFront
func FrontInterface_setFront(newFront FrontInterface_frontInterface) {
    FrontInterface___luneScript = newFront
    
}

// 263: decl @lune.@base.@frontInterface.loadModule
func FrontInterface_loadModule(mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    return FrontInterface___luneScript.LoadModule(mod)
}

// 267: decl @lune.@base.@frontInterface.loadFromLnsTxt
func FrontInterface_loadFromLnsTxt(importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    return FrontInterface___luneScript.LoadFromLnsTxt(importModuleInfo, name, txt)
}

// 273: decl @lune.@base.@frontInterface.loadMeta
func FrontInterface_loadMeta(importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    return FrontInterface___luneScript.LoadMeta(importModuleInfo, mod)
}

// 277: decl @lune.@base.@frontInterface.searchModule
func FrontInterface_searchModule(mod string) LnsAny {
    return FrontInterface___luneScript.SearchModule(mod)
}

// 281: decl @lune.@base.@frontInterface.getLuaModulePath
func FrontInterface_getLuaModulePath(mod string) string {
    return FrontInterface___luneScript.GetLuaModulePath(mod)
}

// declaration Class -- ModuleId
var FrontInterface_ModuleId__tempId *FrontInterface_ModuleId
// 33: decl @lune.@base.@frontInterface.ModuleId.___init
func FrontInterface_ModuleId____init_1020_() {
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
// 47: DeclConstr
func (self *FrontInterface_ModuleId) InitFrontInterface_ModuleId(modTime LnsReal,buildCount LnsInt) {
    self.modTime = modTime
    
    self.buildCount = buildCount
    
    self.idStr = Lns_getVM().String_format("%f:%d", []LnsAny{modTime, buildCount})
    
}

// 53: decl @lune.@base.@frontInterface.ModuleId.getNextModuleId
func (self *FrontInterface_ModuleId) GetNextModuleId() *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(self.modTime, self.buildCount + 1)
}

// 60: decl @lune.@base.@frontInterface.ModuleId.createId
func FrontInterface_ModuleId_createId(modTime LnsReal,buildCount LnsInt) *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(modTime, buildCount)
}

// 63: decl @lune.@base.@frontInterface.ModuleId.createIdFromTxt
func FrontInterface_ModuleId_createIdFromTxt(idStr string) *FrontInterface_ModuleId {
    var modTime LnsReal
    modTime = Lns_unwrapDefault( Lns_tonumber(Lns_car(Lns_getVM().String_gsub(idStr,":.*", "")).(string), nil), 0.0).(LnsReal)
    var buildCount LnsReal
    buildCount = Lns_unwrapDefault( Lns_tonumber(Lns_car(Lns_getVM().String_gsub(idStr,".*:", "")).(string), nil), 0.0).(LnsReal)
    return NewFrontInterface_ModuleId(modTime, (LnsInt)(buildCount))
}


// declaration Class -- ModuleProvideInfo
type FrontInterface_ModuleProvideInfoMtd interface {
    Get_mutable() bool
    Get_symbolKind() LnsInt
    Get_typeInfo() *Ast_TypeInfo
}
type FrontInterface_ModuleProvideInfo struct {
    typeInfo *Ast_TypeInfo
    symbolKind LnsInt
    mutable bool
    FP FrontInterface_ModuleProvideInfoMtd
}
func FrontInterface_ModuleProvideInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_ModuleProvideInfo).FP
}
type FrontInterface_ModuleProvideInfoDownCast interface {
    ToFrontInterface_ModuleProvideInfo() *FrontInterface_ModuleProvideInfo
}
func FrontInterface_ModuleProvideInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(FrontInterface_ModuleProvideInfoDownCast)
    if ok { return work.ToFrontInterface_ModuleProvideInfo() }
    return nil
}
func (obj *FrontInterface_ModuleProvideInfo) ToFrontInterface_ModuleProvideInfo() *FrontInterface_ModuleProvideInfo {
    return obj
}
func NewFrontInterface_ModuleProvideInfo(arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool) *FrontInterface_ModuleProvideInfo {
    obj := &FrontInterface_ModuleProvideInfo{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleProvideInfo(arg1, arg2, arg3)
    return obj
}
func (self *FrontInterface_ModuleProvideInfo) InitFrontInterface_ModuleProvideInfo(arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool) {
    self.typeInfo = arg1
    self.symbolKind = arg2
    self.mutable = arg3
}
func (self *FrontInterface_ModuleProvideInfo) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }
func (self *FrontInterface_ModuleProvideInfo) Get_symbolKind() LnsInt{ return self.symbolKind }
func (self *FrontInterface_ModuleProvideInfo) Get_mutable() bool{ return self.mutable }

// declaration Class -- ModuleInfo
type FrontInterface_ModuleInfoMtd interface {
    Assign(arg1 string) *FrontInterface_ModuleInfo
    GetImportTypeId(arg1 *Ast_TypeInfo) LnsAny
    GetTypeInfo(arg1 LnsInt) LnsAny
    Get_assignName() string
    Get_fullName() string
    Get_importId2localTypeInfoMap() *LnsMap
    Get_importedAliasMap() *LnsMap
    Get_localTypeInfo2importIdMap() *LnsMap
    Get_moduleId() *FrontInterface_ModuleId
    Get_modulePath() string
    Get_moduleProvide() *FrontInterface_ModuleProvideInfo
    Get_moduleTypeInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
}
type FrontInterface_ModuleInfo struct {
    fullName string
    localTypeInfo2importIdMap *LnsMap
    importId2localTypeInfoMap *LnsMap
    assignName string
    moduleId *FrontInterface_ModuleId
    importedAliasMap *LnsMap
    processInfo *Ast_ProcessInfo
    moduleProvide *FrontInterface_ModuleProvideInfo
    moduleTypeInfo *Ast_TypeInfo
    FP FrontInterface_ModuleInfoMtd
}
func FrontInterface_ModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_ModuleInfo).FP
}
type FrontInterface_ModuleInfoDownCast interface {
    ToFrontInterface_ModuleInfo() *FrontInterface_ModuleInfo
}
func FrontInterface_ModuleInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(FrontInterface_ModuleInfoDownCast)
    if ok { return work.ToFrontInterface_ModuleInfo() }
    return nil
}
func (obj *FrontInterface_ModuleInfo) ToFrontInterface_ModuleInfo() *FrontInterface_ModuleInfo {
    return obj
}
func NewFrontInterface_ModuleInfo(arg1 string, arg2 string, arg3 *LnsMap, arg4 *FrontInterface_ModuleId, arg5 *Ast_ProcessInfo, arg6 *FrontInterface_ModuleProvideInfo, arg7 *Ast_TypeInfo, arg8 *LnsMap) *FrontInterface_ModuleInfo {
    obj := &FrontInterface_ModuleInfo{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *FrontInterface_ModuleInfo) Get_fullName() string{ return self.fullName }
func (self *FrontInterface_ModuleInfo) Get_localTypeInfo2importIdMap() *LnsMap{ return self.localTypeInfo2importIdMap }
func (self *FrontInterface_ModuleInfo) Get_importId2localTypeInfoMap() *LnsMap{ return self.importId2localTypeInfoMap }
func (self *FrontInterface_ModuleInfo) Get_assignName() string{ return self.assignName }
func (self *FrontInterface_ModuleInfo) Get_moduleId() *FrontInterface_ModuleId{ return self.moduleId }
func (self *FrontInterface_ModuleInfo) Get_importedAliasMap() *LnsMap{ return self.importedAliasMap }
func (self *FrontInterface_ModuleInfo) Get_processInfo() *Ast_ProcessInfo{ return self.processInfo }
func (self *FrontInterface_ModuleInfo) Get_moduleProvide() *FrontInterface_ModuleProvideInfo{ return self.moduleProvide }
func (self *FrontInterface_ModuleInfo) Get_moduleTypeInfo() *Ast_TypeInfo{ return self.moduleTypeInfo }
// 97: DeclConstr
func (self *FrontInterface_ModuleInfo) InitFrontInterface_ModuleInfo(fullName string,assignName string,idMap *LnsMap,moduleId *FrontInterface_ModuleId,processInfo *Ast_ProcessInfo,moduleProvide *FrontInterface_ModuleProvideInfo,moduleTypeInfo *Ast_TypeInfo,importedAliasMap *LnsMap) {
    self.moduleTypeInfo = moduleTypeInfo
    
    self.moduleProvide = moduleProvide
    
    self.moduleId = moduleId
    
    self.fullName = fullName
    
    self.assignName = assignName
    
    self.localTypeInfo2importIdMap = idMap
    
    self.importId2localTypeInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    for _typeInfo, _importId := range( idMap.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        importId := _importId.(LnsInt)
        self.importId2localTypeInfoMap.Set(importId,typeInfo)
    }
    self.processInfo = processInfo
    
    self.importedAliasMap = importedAliasMap
    
}

// 117: decl @lune.@base.@frontInterface.ModuleInfo.getImportTypeId
func (self *FrontInterface_ModuleInfo) GetImportTypeId(typeInfo *Ast_TypeInfo) LnsAny {
    {
        _typeId := self.localTypeInfo2importIdMap.Get(typeInfo)
        if !Lns_IsNil( _typeId ) {
            typeId := _typeId.(LnsInt)
            return typeId
        }
    }
    return nil
}

// 127: decl @lune.@base.@frontInterface.ModuleInfo.getTypeInfo
func (self *FrontInterface_ModuleInfo) GetTypeInfo(localTypeId LnsInt) LnsAny {
    {
        _typeInfo := self.importId2localTypeInfoMap.Get(localTypeId)
        if !Lns_IsNil( _typeInfo ) {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo
        }
    }
    return nil
}

// 135: decl @lune.@base.@frontInterface.ModuleInfo.get_modulePath
func (self *FrontInterface_ModuleInfo) Get_modulePath() string {
    return self.fullName
}

// 139: decl @lune.@base.@frontInterface.ModuleInfo.assign
func (self *FrontInterface_ModuleInfo) Assign(assignName string) *FrontInterface_ModuleInfo {
    return NewFrontInterface_ModuleInfo(self.fullName, assignName, self.localTypeInfo2importIdMap, self.moduleId, self.processInfo, self.moduleProvide, self.moduleTypeInfo, self.importedAliasMap)
}


// declaration Class -- ModuleMeta
type FrontInterface_ModuleMetaMtd interface {
    Get_lnsPath() string
    Get_metaInfo() LnsAny
    Get_moduleInfo() LnsAny
}
type FrontInterface_ModuleMeta struct {
    metaInfo LnsAny
    lnsPath string
    moduleInfo LnsAny
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
func NewFrontInterface_ModuleMeta(arg1 LnsAny, arg2 string, arg3 LnsAny) *FrontInterface_ModuleMeta {
    obj := &FrontInterface_ModuleMeta{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleMeta(arg1, arg2, arg3)
    return obj
}
func (self *FrontInterface_ModuleMeta) InitFrontInterface_ModuleMeta(arg1 LnsAny, arg2 string, arg3 LnsAny) {
    self.metaInfo = arg1
    self.lnsPath = arg2
    self.moduleInfo = arg3
}
func (self *FrontInterface_ModuleMeta) Get_metaInfo() LnsAny{ return self.metaInfo }
func (self *FrontInterface_ModuleMeta) Get_lnsPath() string{ return self.lnsPath }
func (self *FrontInterface_ModuleMeta) Get_moduleInfo() LnsAny{ return self.moduleInfo }

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
// 160: DeclConstr
func (self *FrontInterface_ImportModuleInfo) InitFrontInterface_ImportModuleInfo() {
    self.orderedSet = NewUtil_OrderedSet()
    
}

// 164: decl @lune.@base.@frontInterface.ImportModuleInfo.add
func (self *FrontInterface_ImportModuleInfo) Add(modulePath string) bool {
    return self.orderedSet.FP.Add(modulePath)
}

// 168: decl @lune.@base.@frontInterface.ImportModuleInfo.remove
func (self *FrontInterface_ImportModuleInfo) Remove() {
    self.orderedSet.FP.RemoveLast()
}

// 172: decl @lune.@base.@frontInterface.ImportModuleInfo.getFull
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
// 226: decl @lune.@base.@frontInterface.dummyFront.loadModule
func (self *frontInterface_dummyFront) LoadModule(mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    var loaded LnsAny
    loaded = frontInterface_convExp657(Lns_2DDD(Lns_getVM().Load("return {}", nil)))
    var emptyTable LnsAny
    if loaded != nil{
        loaded_1562 := loaded.(*Lns_luaValue)
        emptyTable = Lns_unwrap( Lns_car(Lns_getVM().RunLoadedfunc(loaded_1562,Lns_2DDD([]LnsAny{}))[0]))
        
    } else {
        panic("load error")
    }
    var meta *FrontInterface_ModuleMeta
    meta = NewFrontInterface_ModuleMeta(emptyTable, Lns_car(Lns_getVM().String_gsub(mod,"%.", "/")).(string) + ".lns", nil)
    return Lns_require(mod), meta
}

// 238: decl @lune.@base.@frontInterface.dummyFront.loadMeta
func (self *frontInterface_dummyFront) LoadMeta(importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 241: decl @lune.@base.@frontInterface.dummyFront.loadFromLnsTxt
func (self *frontInterface_dummyFront) LoadFromLnsTxt(importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 246: decl @lune.@base.@frontInterface.dummyFront.getLuaModulePath
func (self *frontInterface_dummyFront) GetLuaModulePath(mod string) string {
    panic("not implements")
// insert a dummy
    return ""
}

// 249: decl @lune.@base.@frontInterface.dummyFront.searchModule
func (self *frontInterface_dummyFront) SearchModule(mod string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 252: decl @lune.@base.@frontInterface.dummyFront.error
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
    FrontInterface_ModuleId____init_1020_()
    FrontInterface___luneScript = NewfrontInterface_dummyFront().FP
}
func init() {
    init_frontInterface = false
}
