// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_frontInterface bool
var frontInterface__mod__ string
var FrontInterface___luneScript FrontInterface_frontInterface
// decl alge -- MetaOrModule
type FrontInterface_MetaOrModule = LnsAny
type FrontInterface_MetaOrModule__Meta struct{
Val1 LnsAny
}
func (self *FrontInterface_MetaOrModule__Meta) GetTxt() string {
return "MetaOrModule.Meta"
}
type FrontInterface_MetaOrModule__Module struct{
Val1 *FrontInterface_ModuleInfo
}
func (self *FrontInterface_MetaOrModule__Module) GetTxt() string {
return "MetaOrModule.Module"
}
// for 290
func frontInterface_convExp881(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 109: decl @lune.@base.@frontInterface.getRootDependModId
func FrontInterface_getRootDependModId(_env *LnsEnv) LnsInt {
    return -1
}

// 325: decl @lune.@base.@frontInterface.setFront
func FrontInterface_setFront(_env *LnsEnv, newFront FrontInterface_frontInterface) {
    FrontInterface___luneScript = newFront
    
}

// 329: decl @lune.@base.@frontInterface.loadModule
func FrontInterface_loadModule(_env *LnsEnv, mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    return FrontInterface___luneScript.LoadModule(_env, mod)
}

// 333: decl @lune.@base.@frontInterface.loadFromLnsTxt
func FrontInterface_loadFromLnsTxt(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,baseDir LnsAny,name string,txt string) LnsAny {
    return FrontInterface___luneScript.LoadFromLnsTxt(_env, importModuleInfo, baseDir, name, txt)
}

// 340: decl @lune.@base.@frontInterface.loadMeta
func FrontInterface_loadMeta(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,mod string,orgMod string,baseDir LnsAny,loader FrontInterface_ModuleLoader) LnsAny {
    return FrontInterface___luneScript.LoadMeta(_env, importModuleInfo, mod, orgMod, baseDir, loader)
}

// 346: decl @lune.@base.@frontInterface.searchModule
func FrontInterface_searchModule(_env *LnsEnv, mod string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    return FrontInterface___luneScript.SearchModule(_env, mod, baseDir, addSearchPath)
}

// 350: decl @lune.@base.@frontInterface.getLuaModulePath
func FrontInterface_getLuaModulePath(_env *LnsEnv, mod string,baseDir LnsAny)(string, LnsAny, string) {
    return FrontInterface___luneScript.GetLuaModulePath(_env, mod, baseDir)
}

// declaration Class -- ModuleId
var FrontInterface_ModuleId__tempId *FrontInterface_ModuleId
// 37: decl @lune.@base.@frontInterface.ModuleId.___init
func FrontInterface_ModuleId____init_5_(_env *LnsEnv) {
    FrontInterface_ModuleId__tempId = NewFrontInterface_ModuleId(_env, 0.0, 0)
    
}

type FrontInterface_ModuleIdMtd interface {
    GetNextModuleId(_env *LnsEnv) *FrontInterface_ModuleId
    Get_buildCount(_env *LnsEnv) LnsInt
    Get_idStr(_env *LnsEnv) string
    Get_modTime(_env *LnsEnv) LnsReal
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
func NewFrontInterface_ModuleId(_env *LnsEnv, arg1 LnsReal, arg2 LnsInt) *FrontInterface_ModuleId {
    obj := &FrontInterface_ModuleId{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleId(_env, arg1, arg2)
    return obj
}
func (self *FrontInterface_ModuleId) Get_modTime(_env *LnsEnv) LnsReal{ return self.modTime }
func (self *FrontInterface_ModuleId) Get_buildCount(_env *LnsEnv) LnsInt{ return self.buildCount }
func (self *FrontInterface_ModuleId) Get_idStr(_env *LnsEnv) string{ return self.idStr }
// 51: DeclConstr
func (self *FrontInterface_ModuleId) InitFrontInterface_ModuleId(_env *LnsEnv, modTime LnsReal,buildCount LnsInt) {
    self.modTime = modTime
    
    self.buildCount = buildCount
    
    self.idStr = _env.LuaVM.String_format("%f:%d", []LnsAny{modTime, buildCount})
    
}

// 57: decl @lune.@base.@frontInterface.ModuleId.getNextModuleId
func (self *FrontInterface_ModuleId) GetNextModuleId(_env *LnsEnv) *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(_env, self.modTime, self.buildCount + 1)
}

// 64: decl @lune.@base.@frontInterface.ModuleId.createId
func FrontInterface_ModuleId_createId(_env *LnsEnv, modTime LnsReal,buildCount LnsInt) *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(_env, modTime, buildCount)
}

// 67: decl @lune.@base.@frontInterface.ModuleId.createIdFromTxt
func FrontInterface_ModuleId_createIdFromTxt(_env *LnsEnv, idStr string) *FrontInterface_ModuleId {
    var modTime LnsReal
    modTime = Lns_unwrapDefault( Lns_tonumber(Lns_car(_env.LuaVM.String_gsub(idStr,":.*", "")).(string), nil), 0.0).(LnsReal)
    var buildCount LnsReal
    buildCount = Lns_unwrapDefault( Lns_tonumber(Lns_car(_env.LuaVM.String_gsub(idStr,".*:", "")).(string), nil), 0.0).(LnsReal)
    return NewFrontInterface_ModuleId(_env, modTime, (LnsInt)(buildCount))
}


// declaration Class -- ModuleProvideInfo
type FrontInterface_ModuleProvideInfoMtd interface {
    Get_mutable(_env *LnsEnv) bool
    Get_symbolKind(_env *LnsEnv) LnsInt
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
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
func NewFrontInterface_ModuleProvideInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool) *FrontInterface_ModuleProvideInfo {
    obj := &FrontInterface_ModuleProvideInfo{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleProvideInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *FrontInterface_ModuleProvideInfo) InitFrontInterface_ModuleProvideInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool) {
    self.typeInfo = arg1
    self.symbolKind = arg2
    self.mutable = arg3
}
func (self *FrontInterface_ModuleProvideInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *FrontInterface_ModuleProvideInfo) Get_symbolKind(_env *LnsEnv) LnsInt{ return self.symbolKind }
func (self *FrontInterface_ModuleProvideInfo) Get_mutable(_env *LnsEnv) bool{ return self.mutable }

// declaration Class -- LuneHelperInfo
type FrontInterface_LuneHelperInfoMtd interface {
}
type FrontInterface_LuneHelperInfo struct {
    UseNilAccess bool
    UseUnwrapExp bool
    HasMappingClassDef bool
    UseLoad bool
    UseUnpack bool
    UseAlge bool
    UseSet bool
    CallAnonymous bool
    PragmaSet *LnsSet
    UseLazyLoad bool
    UseLazyRequire bool
    UseRun bool
    FP FrontInterface_LuneHelperInfoMtd
}
func FrontInterface_LuneHelperInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_LuneHelperInfo).FP
}
type FrontInterface_LuneHelperInfoDownCast interface {
    ToFrontInterface_LuneHelperInfo() *FrontInterface_LuneHelperInfo
}
func FrontInterface_LuneHelperInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(FrontInterface_LuneHelperInfoDownCast)
    if ok { return work.ToFrontInterface_LuneHelperInfo() }
    return nil
}
func (obj *FrontInterface_LuneHelperInfo) ToFrontInterface_LuneHelperInfo() *FrontInterface_LuneHelperInfo {
    return obj
}
func NewFrontInterface_LuneHelperInfo(_env *LnsEnv) *FrontInterface_LuneHelperInfo {
    obj := &FrontInterface_LuneHelperInfo{}
    obj.FP = obj
    obj.InitFrontInterface_LuneHelperInfo(_env)
    return obj
}
// 93: DeclConstr
func (self *FrontInterface_LuneHelperInfo) InitFrontInterface_LuneHelperInfo(_env *LnsEnv) {
    self.UseNilAccess = false
    
    self.UseUnwrapExp = false
    
    self.HasMappingClassDef = false
    
    self.UseLoad = false
    
    self.UseUnpack = false
    
    self.UseAlge = false
    
    self.UseSet = false
    
    self.CallAnonymous = false
    
    self.PragmaSet = NewLnsSet([]LnsAny{})
    
    self.UseLazyLoad = false
    
    self.UseLazyRequire = false
    
    self.UseRun = false
    
}


// declaration Class -- ExportInfo
type FrontInterface_ExportInfoMtd interface {
    Get_globalSymbolList(_env *LnsEnv) *LnsList
    Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_provideInfo(_env *LnsEnv) *FrontInterface_ModuleProvideInfo
}
type FrontInterface_ExportInfo struct {
    moduleTypeInfo *Ast_TypeInfo
    provideInfo *FrontInterface_ModuleProvideInfo
    processInfo *Ast_ProcessInfo
    globalSymbolList *LnsList
    FP FrontInterface_ExportInfoMtd
}
func FrontInterface_ExportInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_ExportInfo).FP
}
type FrontInterface_ExportInfoDownCast interface {
    ToFrontInterface_ExportInfo() *FrontInterface_ExportInfo
}
func FrontInterface_ExportInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(FrontInterface_ExportInfoDownCast)
    if ok { return work.ToFrontInterface_ExportInfo() }
    return nil
}
func (obj *FrontInterface_ExportInfo) ToFrontInterface_ExportInfo() *FrontInterface_ExportInfo {
    return obj
}
func NewFrontInterface_ExportInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *FrontInterface_ModuleProvideInfo, arg3 *Ast_ProcessInfo, arg4 *LnsList) *FrontInterface_ExportInfo {
    obj := &FrontInterface_ExportInfo{}
    obj.FP = obj
    obj.InitFrontInterface_ExportInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *FrontInterface_ExportInfo) InitFrontInterface_ExportInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *FrontInterface_ModuleProvideInfo, arg3 *Ast_ProcessInfo, arg4 *LnsList) {
    self.moduleTypeInfo = arg1
    self.provideInfo = arg2
    self.processInfo = arg3
    self.globalSymbolList = arg4
}
func (self *FrontInterface_ExportInfo) Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.moduleTypeInfo }
func (self *FrontInterface_ExportInfo) Get_provideInfo(_env *LnsEnv) *FrontInterface_ModuleProvideInfo{ return self.provideInfo }
func (self *FrontInterface_ExportInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo{ return self.processInfo }
func (self *FrontInterface_ExportInfo) Get_globalSymbolList(_env *LnsEnv) *LnsList{ return self.globalSymbolList }

// declaration Class -- ModuleInfo
type FrontInterface_ModuleInfoMtd interface {
    Assign(_env *LnsEnv, arg1 string) *FrontInterface_ModuleInfo
    GetTypeInfo(_env *LnsEnv, arg1 LnsInt) LnsAny
    Get_assignName(_env *LnsEnv) string
    Get_exportInfo(_env *LnsEnv) *FrontInterface_ExportInfo
    Get_fullName(_env *LnsEnv) string
    Get_importId2localTypeInfoMap(_env *LnsEnv) *LnsMap
    Get_importedAliasMap(_env *LnsEnv) *LnsMap
    Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId
    Get_modulePath(_env *LnsEnv) string
    Get_streamName(_env *LnsEnv) string
}
type FrontInterface_ModuleInfo struct {
    streamName string
    fullName string
    importId2localTypeInfoMap *LnsMap
    assignName string
    moduleId *FrontInterface_ModuleId
    importedAliasMap *LnsMap
    exportInfo *FrontInterface_ExportInfo
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
func NewFrontInterface_ModuleInfo(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 *LnsMap, arg5 *FrontInterface_ModuleId, arg6 *FrontInterface_ExportInfo, arg7 *LnsMap) *FrontInterface_ModuleInfo {
    obj := &FrontInterface_ModuleInfo{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *FrontInterface_ModuleInfo) Get_streamName(_env *LnsEnv) string{ return self.streamName }
func (self *FrontInterface_ModuleInfo) Get_fullName(_env *LnsEnv) string{ return self.fullName }
func (self *FrontInterface_ModuleInfo) Get_importId2localTypeInfoMap(_env *LnsEnv) *LnsMap{ return self.importId2localTypeInfoMap }
func (self *FrontInterface_ModuleInfo) Get_assignName(_env *LnsEnv) string{ return self.assignName }
func (self *FrontInterface_ModuleInfo) Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId{ return self.moduleId }
func (self *FrontInterface_ModuleInfo) Get_importedAliasMap(_env *LnsEnv) *LnsMap{ return self.importedAliasMap }
func (self *FrontInterface_ModuleInfo) Get_exportInfo(_env *LnsEnv) *FrontInterface_ExportInfo{ return self.exportInfo }
// 143: DeclConstr
func (self *FrontInterface_ModuleInfo) InitFrontInterface_ModuleInfo(_env *LnsEnv, streamName string,fullName string,assignName string,idMap *LnsMap,moduleId *FrontInterface_ModuleId,exportInfo *FrontInterface_ExportInfo,importedAliasMap *LnsMap) {
    self.streamName = streamName
    
    self.exportInfo = exportInfo
    
    self.moduleId = moduleId
    
    self.fullName = fullName
    
    self.assignName = assignName
    
    var importId2localTypeInfoMap *LnsMap
    importId2localTypeInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    for _typeInfo, _importId := range( idMap.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        importId := _importId.(LnsInt)
        importId2localTypeInfoMap.Set(importId,typeInfo)
    }
    self.importId2localTypeInfoMap = importId2localTypeInfoMap
    
    self.importedAliasMap = importedAliasMap
    
}

// 161: decl @lune.@base.@frontInterface.ModuleInfo.getTypeInfo
func (self *FrontInterface_ModuleInfo) GetTypeInfo(_env *LnsEnv, localTypeId LnsInt) LnsAny {
    {
        _typeInfo := self.importId2localTypeInfoMap.Get(localTypeId)
        if !Lns_IsNil( _typeInfo ) {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo
        }
    }
    return nil
}

// 168: decl @lune.@base.@frontInterface.ModuleInfo.get_modulePath
func (self *FrontInterface_ModuleInfo) Get_modulePath(_env *LnsEnv) string {
    return self.fullName
}

// 172: decl @lune.@base.@frontInterface.ModuleInfo.assign
func (self *FrontInterface_ModuleInfo) Assign(_env *LnsEnv, assignName string) *FrontInterface_ModuleInfo {
    var info *FrontInterface_ModuleInfo
    info = NewFrontInterface_ModuleInfo(_env, self.streamName, self.fullName, assignName, NewLnsMap( map[LnsAny]LnsAny{}), self.moduleId, self.exportInfo, self.importedAliasMap)
    info.importId2localTypeInfoMap = self.importId2localTypeInfoMap
    
    return info
}


// declaration Class -- ModuleMeta
type FrontInterface_ModuleMetaMtd interface {
    Get_lnsPath(_env *LnsEnv) string
    Get_metaOrModule(_env *LnsEnv) LnsAny
    Set_metaOrModule(_env *LnsEnv, arg1 LnsAny)
}
type FrontInterface_ModuleMeta struct {
    lnsPath string
    metaOrModule LnsAny
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
func NewFrontInterface_ModuleMeta(_env *LnsEnv, arg1 string, arg2 LnsAny) *FrontInterface_ModuleMeta {
    obj := &FrontInterface_ModuleMeta{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleMeta(_env, arg1, arg2)
    return obj
}
func (self *FrontInterface_ModuleMeta) InitFrontInterface_ModuleMeta(_env *LnsEnv, arg1 string, arg2 LnsAny) {
    self.lnsPath = arg1
    self.metaOrModule = arg2
}
func (self *FrontInterface_ModuleMeta) Get_lnsPath(_env *LnsEnv) string{ return self.lnsPath }
func (self *FrontInterface_ModuleMeta) Get_metaOrModule(_env *LnsEnv) LnsAny{ return self.metaOrModule }
func (self *FrontInterface_ModuleMeta) Set_metaOrModule(_env *LnsEnv, arg1 LnsAny){ self.metaOrModule = arg1 }

// declaration Class -- ImportModuleInfo
type FrontInterface_ImportModuleInfoMtd interface {
    Add(_env *LnsEnv, arg1 string) bool
    Clone(_env *LnsEnv) *FrontInterface_ImportModuleInfo
    GetFull(_env *LnsEnv) string
    Len(_env *LnsEnv) LnsInt
    List(_env *LnsEnv) *LnsList
    Remove(_env *LnsEnv)
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
func NewFrontInterface_ImportModuleInfo(_env *LnsEnv) *FrontInterface_ImportModuleInfo {
    obj := &FrontInterface_ImportModuleInfo{}
    obj.FP = obj
    obj.InitFrontInterface_ImportModuleInfo(_env)
    return obj
}
// 200: DeclConstr
func (self *FrontInterface_ImportModuleInfo) InitFrontInterface_ImportModuleInfo(_env *LnsEnv) {
    self.orderedSet = NewUtil_OrderedSet(_env)
    
}

// 204: decl @lune.@base.@frontInterface.ImportModuleInfo.add
func (self *FrontInterface_ImportModuleInfo) Add(_env *LnsEnv, modulePath string) bool {
    return self.orderedSet.FP.Add(_env, modulePath)
}

// 208: decl @lune.@base.@frontInterface.ImportModuleInfo.remove
func (self *FrontInterface_ImportModuleInfo) Remove(_env *LnsEnv) {
    self.orderedSet.FP.RemoveLast(_env)
}

// 212: decl @lune.@base.@frontInterface.ImportModuleInfo.getFull
func (self *FrontInterface_ImportModuleInfo) GetFull(_env *LnsEnv) string {
    var txt string
    txt = ""
    for _, _modulePath := range( self.orderedSet.FP.Get_list(_env).Items ) {
        modulePath := _modulePath.(string)
        txt = _env.LuaVM.String_format("%s -> %s", []LnsAny{txt, modulePath})
        
    }
    return txt
}

// 220: decl @lune.@base.@frontInterface.ImportModuleInfo.clone
func (self *FrontInterface_ImportModuleInfo) Clone(_env *LnsEnv) *FrontInterface_ImportModuleInfo {
    var info *FrontInterface_ImportModuleInfo
    info = NewFrontInterface_ImportModuleInfo(_env)
    for _, _mod := range( self.orderedSet.FP.Get_list(_env).Items ) {
        mod := _mod.(string)
        info.FP.Add(_env, mod)
    }
    return info
}

// 228: decl @lune.@base.@frontInterface.ImportModuleInfo.len
func (self *FrontInterface_ImportModuleInfo) Len(_env *LnsEnv) LnsInt {
    return self.orderedSet.FP.Get_list(_env).Len()
}

// 232: decl @lune.@base.@frontInterface.ImportModuleInfo.list
func (self *FrontInterface_ImportModuleInfo) List(_env *LnsEnv) *LnsList {
    return self.orderedSet.FP.Get_list(_env)
}


type FrontInterface_ModuleLoader interface {
        GetModuleInfo(_env *LnsEnv) LnsAny
}
func Lns_cast2FrontInterface_ModuleLoader( obj LnsAny ) LnsAny {
    if _, ok := obj.(FrontInterface_ModuleLoader); ok { 
        return obj
    }
    return nil
}

type FrontInterface_frontInterface interface {
        Error(_env *LnsEnv, arg1 string)
        GetLuaModulePath(_env *LnsEnv, arg1 string, arg2 LnsAny)(string, LnsAny, string)
        LoadFromLnsTxt(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string) LnsAny
        LoadMeta(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string, arg4 LnsAny, arg5 FrontInterface_ModuleLoader) LnsAny
        LoadModule(_env *LnsEnv, arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
        SearchModule(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny) LnsAny
}
func Lns_cast2FrontInterface_frontInterface( obj LnsAny ) LnsAny {
    if _, ok := obj.(FrontInterface_frontInterface); ok { 
        return obj
    }
    return nil
}

// declaration Class -- dummyFront
type frontInterface_dummyFrontMtd interface {
    Error(_env *LnsEnv, arg1 string)
    GetLuaModulePath(_env *LnsEnv, arg1 string, arg2 LnsAny)(string, LnsAny, string)
    LoadFromLnsTxt(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string) LnsAny
    LoadMeta(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string, arg4 LnsAny, arg5 FrontInterface_ModuleLoader) LnsAny
    LoadModule(_env *LnsEnv, arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
    SearchModule(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny) LnsAny
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
func NewfrontInterface_dummyFront(_env *LnsEnv) *frontInterface_dummyFront {
    obj := &frontInterface_dummyFront{}
    obj.FP = obj
    obj.InitfrontInterface_dummyFront(_env)
    return obj
}
func (self *frontInterface_dummyFront) InitfrontInterface_dummyFront(_env *LnsEnv) {
}
// 288: decl @lune.@base.@frontInterface.dummyFront.loadModule
func (self *frontInterface_dummyFront) LoadModule(_env *LnsEnv, mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    var loaded LnsAny
    loaded = frontInterface_convExp881(Lns_2DDD(_env.CommonLuaVM.Load("return {}", nil)))
    var emptyTable LnsAny
    if loaded != nil{
        loaded_216 := loaded.(*Lns_luaValue)
        emptyTable = Lns_unwrap( Lns_car(_env.CommonLuaVM.RunLoadedfunc(loaded_216,Lns_2DDD([]LnsAny{}))[0]))
        
    } else {
        panic("load error")
    }
    var meta *FrontInterface_ModuleMeta
    meta = NewFrontInterface_ModuleMeta(_env, Lns_car(_env.LuaVM.String_gsub(mod,"%.", "/")).(string) + ".lns", &FrontInterface_MetaOrModule__Meta{emptyTable})
    return Lns_require(mod), meta
}

// 301: decl @lune.@base.@frontInterface.dummyFront.loadMeta
func (self *frontInterface_dummyFront) LoadMeta(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,mod string,orgMod string,baseDir LnsAny,loader FrontInterface_ModuleLoader) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 306: decl @lune.@base.@frontInterface.dummyFront.loadFromLnsTxt
func (self *frontInterface_dummyFront) LoadFromLnsTxt(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,baseDir LnsAny,name string,txt string) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 312: decl @lune.@base.@frontInterface.dummyFront.getLuaModulePath
func (self *frontInterface_dummyFront) GetLuaModulePath(_env *LnsEnv, mod string,baseDir LnsAny)(string, LnsAny, string) {
    panic("not implements")
// insert a dummy
    return "",nil,""
}

// 315: decl @lune.@base.@frontInterface.dummyFront.searchModule
func (self *frontInterface_dummyFront) SearchModule(_env *LnsEnv, mod string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    panic("not implements")
// insert a dummy
    return nil
}

// 318: decl @lune.@base.@frontInterface.dummyFront.error
func (self *frontInterface_dummyFront) Error(_env *LnsEnv, message string) {
    panic("not implements")
}


func Lns_frontInterface_init(_env *LnsEnv) {
    if init_frontInterface { return }
    init_frontInterface = true
    frontInterface__mod__ = "@lune.@base.@frontInterface"
    Lns_InitMod()
    Lns_Util_init(_env)
    Lns_Ast_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Runner_init(_env)
    FrontInterface_ModuleId____init_5_(_env)
    FrontInterface___luneScript = NewfrontInterface_dummyFront(_env).FP
}
func init() {
    init_frontInterface = false
}
