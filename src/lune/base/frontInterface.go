// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_frontInterface bool
var frontInterface__mod__ string
var FrontInterface_dummyFront FrontInterface_frontInterface
// decl alge -- MetaOrModule
type FrontInterface_MetaOrModule = LnsAny
type FrontInterface_MetaOrModule__Export struct{
Val1 *FrontInterface_ExportInfo
}
func (self *FrontInterface_MetaOrModule__Export) GetTxt() string {
return "MetaOrModule.Export"
}
type FrontInterface_MetaOrModule__MetaRaw struct{
Val1 LnsAny
}
func (self *FrontInterface_MetaOrModule__MetaRaw) GetTxt() string {
return "MetaOrModule.MetaRaw"
}
type FrontInterface_MetaOrModule__Module struct{
Val1 *FrontInterface_ModuleInfo
Val2 *FrontInterface_ExportInfo
}
func (self *FrontInterface_MetaOrModule__Module) GetTxt() string {
return "MetaOrModule.Module"
}
// for 375
func frontInterface_convExp0_1554(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 141: decl @lune.@base.@frontInterface.getRootDependModId
func FrontInterface_getRootDependModId(_env *LnsEnv) LnsInt {
    return -1
}

// 370: decl @lune.@base.@frontInterface.dummyLoadModule
func FrontInterface_dummyLoadModule(_env *LnsEnv, mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    var modVal LnsAny
    var moduleMeta *FrontInterface_ModuleMeta
    Lns_LockEnvSync( _env, 373, func () {
        var emptyTable LnsAny
        var loaded LnsAny
        loaded = frontInterface_convExp0_1554(Lns_2DDD(_env.GetVM().Load("return {}", nil)))
        if loaded != nil{
            loaded_333 := loaded.(*Lns_luaValue)
            emptyTable = Lns_unwrap( Lns_car(_env.GetVM().RunLoadedfunc(loaded_333,Lns_2DDD([]LnsAny{}))))
        } else {
            Util_err(_env, "load error")
        }
        moduleMeta = NewFrontInterface_ModuleMeta(_env, Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string) + ".lns", &FrontInterface_MetaOrModule__MetaRaw{emptyTable})
        modVal = Lns_require(mod)
    })
    return modVal, moduleMeta
}

// 58: decl @lune.@base.@frontInterface.ModuleId.getNextModuleId
func (self *FrontInterface_ModuleId) GetNextModuleId(_env *LnsEnv) *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(_env, self.modTime, self.buildCount + 1)
}
// 65: decl @lune.@base.@frontInterface.ModuleId.createId
func FrontInterface_ModuleId_createId(_env *LnsEnv, modTime LnsReal,buildCount LnsInt) *FrontInterface_ModuleId {
    return NewFrontInterface_ModuleId(_env, modTime, buildCount)
}
// 68: decl @lune.@base.@frontInterface.ModuleId.createIdFromTxt
func FrontInterface_ModuleId_createIdFromTxt(_env *LnsEnv, idStr string) *FrontInterface_ModuleId {
    var modTime LnsReal
    modTime = Lns_unwrapDefault( Lns_tonumber(Lns_car(_env.GetVM().String_gsub(idStr,":.*", "")).(string), nil), 0.0).(LnsReal)
    var buildCount LnsReal
    buildCount = Lns_unwrapDefault( Lns_tonumber(Lns_car(_env.GetVM().String_gsub(idStr,".*:", "")).(string), nil), 0.0).(LnsReal)
    return NewFrontInterface_ModuleId(_env, modTime, (LnsInt)(buildCount))
}
// 115: decl @lune.@base.@frontInterface.LuneHelperInfo.mergeFrom
func (self *FrontInterface_LuneHelperInfo) MergeFrom(_env *LnsEnv, src *FrontInterface_LuneHelperInfo) {
    self.UseNilAccess = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseNilAccess) ||
        _env.SetStackVal( src.UseNilAccess) ).(bool)
    
    self.UseUnwrapExp = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseUnwrapExp) ||
        _env.SetStackVal( src.UseUnwrapExp) ).(bool)
    
    self.HasMappingClassDef = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.HasMappingClassDef) ||
        _env.SetStackVal( src.HasMappingClassDef) ).(bool)
    
    self.UseLoad = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseLoad) ||
        _env.SetStackVal( src.UseLoad) ).(bool)
    
    self.UseUnpack = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseUnpack) ||
        _env.SetStackVal( src.UseUnpack) ).(bool)
    
    self.UseAlge = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseAlge) ||
        _env.SetStackVal( src.UseAlge) ).(bool)
    
    self.UseSet = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseSet) ||
        _env.SetStackVal( src.UseSet) ).(bool)
    
    self.CallAnonymous = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.CallAnonymous) ||
        _env.SetStackVal( src.CallAnonymous) ).(bool)
    
    self.UseLazyLoad = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseLazyLoad) ||
        _env.SetStackVal( src.UseLazyLoad) ).(bool)
    
    self.UseLazyRequire = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseLazyRequire) ||
        _env.SetStackVal( src.UseLazyRequire) ).(bool)
    
    self.UseRun = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseRun) ||
        _env.SetStackVal( src.UseRun) ).(bool)
    
    self.UseStrReplace = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseStrReplace) ||
        _env.SetStackVal( src.UseStrReplace) ).(bool)
    
    self.UseResult = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseResult) ||
        _env.SetStackVal( src.UseResult) ).(bool)
    
    self.UseError = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.UseError) ||
        _env.SetStackVal( src.UseError) ).(bool)
    
    for _val := range( src.PragmaSet.Items ) {
        val := _val
        self.PragmaSet.Add(val)
    }
}
// 199: decl @lune.@base.@frontInterface.ExportInfo.get_modulePath
func (self *FrontInterface_ExportInfo) Get_modulePath(_env *LnsEnv) string {
    return self.fullName
}
// 203: decl @lune.@base.@frontInterface.ExportInfo.getTypeInfo
func (self *FrontInterface_ExportInfo) GetTypeInfo(_env *LnsEnv, localTypeId LnsInt) LnsAny {
    {
        _typeInfo := self.importId2localTypeInfoMap.Get(localTypeId)
        if !Lns_IsNil( _typeInfo ) {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo
        }
    }
    return nil
}
// 211: decl @lune.@base.@frontInterface.ExportInfo.assign
func (self *FrontInterface_ExportInfo) Assign(_env *LnsEnv, assignName string) *FrontInterface_ExportInfo {
    var info *FrontInterface_ExportInfo
    info = NewFrontInterface_ExportInfo(_env, self.moduleTypeInfo, self.provideInfo, self.processInfo, self.globalSymbolList, self.importedAliasMap, self.moduleId, self.fullName, assignName, self.streamName, NewLnsMap2_[*Ast_TypeInfo,LnsInt]( map[*Ast_TypeInfo]LnsInt{}))
    info.importId2localTypeInfoMap = self.importId2localTypeInfoMap
    return info
}
// 256: decl @lune.@base.@frontInterface.ImportModuleInfo.add
func (self *FrontInterface_ImportModuleInfo) Add(_env *LnsEnv, modulePath string) bool {
    return self.orderedSet.FP.Add(_env, modulePath)
}
// 260: decl @lune.@base.@frontInterface.ImportModuleInfo.remove
func (self *FrontInterface_ImportModuleInfo) Remove(_env *LnsEnv) {
    self.orderedSet.FP.RemoveLast(_env)
}
// 264: decl @lune.@base.@frontInterface.ImportModuleInfo.getFull
func (self *FrontInterface_ImportModuleInfo) GetFull(_env *LnsEnv) string {
    var txt string
    txt = ""
    for _, _modulePath := range( self.orderedSet.FP.Get_list(_env).Items ) {
        modulePath := _modulePath.(string)
        txt = _env.GetVM().String_format("%s -> %s", Lns_2DDD(txt, modulePath))
    }
    return txt
}
// 272: decl @lune.@base.@frontInterface.ImportModuleInfo.clone
func (self *FrontInterface_ImportModuleInfo) Clone(_env *LnsEnv) *FrontInterface_ImportModuleInfo {
    var info *FrontInterface_ImportModuleInfo
    info = NewFrontInterface_ImportModuleInfo(_env)
    for _, _mod := range( self.orderedSet.FP.Get_list(_env).Items ) {
        mod := _mod.(string)
        info.FP.Add(_env, mod)
    }
    return info
}
// 280: decl @lune.@base.@frontInterface.ImportModuleInfo.len
func (self *FrontInterface_ImportModuleInfo) Len(_env *LnsEnv) LnsInt {
    return self.orderedSet.FP.Get_list(_env).Len()
}
// 284: decl @lune.@base.@frontInterface.ImportModuleInfo.list
func (self *FrontInterface_ImportModuleInfo) List(_env *LnsEnv) *LnsList {
    return self.orderedSet.FP.Get_list(_env)
}
// 342: decl @lune.@base.@frontInterface.FrontAccessor.loadModule
func (self *FrontInterface_FrontAccessor) LoadModule(_env *LnsEnv, mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    return self.frontIF.LoadModule(_env, mod)
}
// 346: decl @lune.@base.@frontInterface.FrontAccessor.loadMeta
func (self *FrontInterface_FrontAccessor) LoadMeta(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,mod string,orgMod string,baseDir LnsAny,loader FrontInterface_ModuleLoader) LnsAny {
    return self.frontIF.LoadMeta(_env, importModuleInfo, mod, orgMod, baseDir, loader)
}
// 352: decl @lune.@base.@frontInterface.FrontAccessor.loadFromLnsTxt
func (self *FrontInterface_FrontAccessor) LoadFromLnsTxt(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,baseDir LnsAny,name string,txt string) LnsAny {
    return self.frontIF.LoadFromLnsTxt(_env, importModuleInfo, baseDir, name, txt)
}
// 359: decl @lune.@base.@frontInterface.FrontAccessor.getLuaModulePath
func (self *FrontInterface_FrontAccessor) GetLuaModulePath(_env *LnsEnv, mod string,baseDir LnsAny)(string, LnsAny, string) {
    return self.frontIF.GetLuaModulePath(_env, mod, baseDir)
}
// 362: decl @lune.@base.@frontInterface.FrontAccessor.searchModule
func (self *FrontInterface_FrontAccessor) SearchModule(_env *LnsEnv, mod string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    return self.frontIF.SearchModule(_env, mod, baseDir, addSearchPath)
}
// 365: decl @lune.@base.@frontInterface.FrontAccessor.error
func (self *FrontInterface_FrontAccessor) Error(_env *LnsEnv, message string) {
    self.frontIF.Error(_env, message)
}
// 389: decl @lune.@base.@frontInterface.DummyFront.loadModule
func (self *frontInterface_DummyFront) LoadModule(_env *LnsEnv, mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    return FrontInterface_dummyLoadModule(_env, mod)
}
// 392: decl @lune.@base.@frontInterface.DummyFront.loadMeta
func (self *frontInterface_DummyFront) LoadMeta(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,mod string,orgMod string,baseDir LnsAny,loader FrontInterface_ModuleLoader) LnsAny {
    __func__ := "@lune.@base.@frontInterface.DummyFront.loadMeta"
    Util_err(_env, _env.GetVM().String_format("not implements: %s", Lns_2DDD(__func__)))
// insert a dummy
    return nil
}
// 397: decl @lune.@base.@frontInterface.DummyFront.loadFromLnsTxt
func (self *frontInterface_DummyFront) LoadFromLnsTxt(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,baseDir LnsAny,name string,txt string) LnsAny {
    __func__ := "@lune.@base.@frontInterface.DummyFront.loadFromLnsTxt"
    Util_err(_env, _env.GetVM().String_format("not implements: %s", Lns_2DDD(__func__)))
// insert a dummy
    return nil
}
// 403: decl @lune.@base.@frontInterface.DummyFront.getLuaModulePath
func (self *frontInterface_DummyFront) GetLuaModulePath(_env *LnsEnv, mod string,baseDir LnsAny)(string, LnsAny, string) {
    return mod, nil, mod
}
// 407: decl @lune.@base.@frontInterface.DummyFront.searchModule
func (self *frontInterface_DummyFront) SearchModule(_env *LnsEnv, mod string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@frontInterface.DummyFront.searchModule"
    Util_err(_env, _env.GetVM().String_format("not implements: %s", Lns_2DDD(__func__)))
// insert a dummy
    return nil
}
// 410: decl @lune.@base.@frontInterface.DummyFront.error
func (self *frontInterface_DummyFront) Error(_env *LnsEnv, message string) {
    Util_err(_env, message)
}
// declaration Class -- ModuleId
var FrontInterface_ModuleId__tempId *FrontInterface_ModuleId
// 38: decl @lune.@base.@frontInterface.ModuleId.___init
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
func FrontInterface_ModuleId_toSlice(slice []LnsAny) []*FrontInterface_ModuleId {
    ret := make([]*FrontInterface_ModuleId, len(slice))
    for index, val := range slice {
        ret[index] = val.(FrontInterface_ModuleIdDownCast).ToFrontInterface_ModuleId()
    }
    return ret
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
// 52: DeclConstr
func (self *FrontInterface_ModuleId) InitFrontInterface_ModuleId(_env *LnsEnv, modTime LnsReal,buildCount LnsInt) {
    self.modTime = modTime
    self.buildCount = buildCount
    self.idStr = _env.GetVM().String_format("%f:%d", Lns_2DDD(modTime, buildCount))
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
func FrontInterface_ModuleProvideInfo_toSlice(slice []LnsAny) []*FrontInterface_ModuleProvideInfo {
    ret := make([]*FrontInterface_ModuleProvideInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(FrontInterface_ModuleProvideInfoDownCast).ToFrontInterface_ModuleProvideInfo()
    }
    return ret
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
    MergeFrom(_env *LnsEnv, arg1 *FrontInterface_LuneHelperInfo)
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
    PragmaSet *LnsSet2_[LnsAny]
    UseLazyLoad bool
    UseLazyRequire bool
    UseRun bool
    UseStrReplace bool
    UseResult bool
    UseError bool
    FP FrontInterface_LuneHelperInfoMtd
}
func FrontInterface_LuneHelperInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_LuneHelperInfo).FP
}
func FrontInterface_LuneHelperInfo_toSlice(slice []LnsAny) []*FrontInterface_LuneHelperInfo {
    ret := make([]*FrontInterface_LuneHelperInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(FrontInterface_LuneHelperInfoDownCast).ToFrontInterface_LuneHelperInfo()
    }
    return ret
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
// 97: DeclConstr
func (self *FrontInterface_LuneHelperInfo) InitFrontInterface_LuneHelperInfo(_env *LnsEnv) {
    self.UseStrReplace = false
    self.UseNilAccess = false
    self.UseUnwrapExp = false
    self.HasMappingClassDef = false
    self.UseLoad = false
    self.UseUnpack = false
    self.UseAlge = false
    self.UseSet = false
    self.CallAnonymous = false
    self.PragmaSet = NewLnsSet2_[LnsAny]([]LnsAny{})
    self.UseLazyLoad = false
    self.UseLazyRequire = false
    self.UseRun = false
    self.UseResult = false
    self.UseError = false
}


// declaration Class -- ExportInfo
type FrontInterface_ExportInfoMtd interface {
    Assign(_env *LnsEnv, arg1 string) *FrontInterface_ExportInfo
    GetTypeInfo(_env *LnsEnv, arg1 LnsInt) LnsAny
    Get_assignName(_env *LnsEnv) string
    Get_fullName(_env *LnsEnv) string
    Get_globalSymbolList(_env *LnsEnv) *LnsList2_[*Ast_SymbolInfo]
    Get_importId2localTypeInfoMap(_env *LnsEnv) *LnsMap2_[LnsInt,*Ast_TypeInfo]
    Get_importedAliasMap(_env *LnsEnv) *LnsMap2_[*Ast_TypeInfo,*Ast_AliasTypeInfo]
    Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId
    Get_modulePath(_env *LnsEnv) string
    Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_provideInfo(_env *LnsEnv) *FrontInterface_ModuleProvideInfo
    Get_streamName(_env *LnsEnv) string
    Set_importId2localTypeInfoMap(_env *LnsEnv, arg1 *LnsMap2_[LnsInt,*Ast_TypeInfo])
}
type FrontInterface_ExportInfo struct {
    moduleTypeInfo *Ast_TypeInfo
    provideInfo *FrontInterface_ModuleProvideInfo
    processInfo *Ast_ProcessInfo
    globalSymbolList *LnsList2_[*Ast_SymbolInfo]
    importedAliasMap *LnsMap2_[*Ast_TypeInfo,*Ast_AliasTypeInfo]
    moduleId *FrontInterface_ModuleId
    fullName string
    assignName string
    streamName string
    importId2localTypeInfoMap *LnsMap2_[LnsInt,*Ast_TypeInfo]
    FP FrontInterface_ExportInfoMtd
}
func FrontInterface_ExportInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_ExportInfo).FP
}
func FrontInterface_ExportInfo_toSlice(slice []LnsAny) []*FrontInterface_ExportInfo {
    ret := make([]*FrontInterface_ExportInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(FrontInterface_ExportInfoDownCast).ToFrontInterface_ExportInfo()
    }
    return ret
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
func NewFrontInterface_ExportInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *FrontInterface_ModuleProvideInfo, arg3 *Ast_ProcessInfo, arg4 *LnsList2_[*Ast_SymbolInfo], arg5 *LnsMap2_[*Ast_TypeInfo,*Ast_AliasTypeInfo], arg6 *FrontInterface_ModuleId, arg7 string, arg8 string, arg9 string, arg10 *LnsMap2_[*Ast_TypeInfo,LnsInt]) *FrontInterface_ExportInfo {
    obj := &FrontInterface_ExportInfo{}
    obj.FP = obj
    obj.InitFrontInterface_ExportInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *FrontInterface_ExportInfo) Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.moduleTypeInfo }
func (self *FrontInterface_ExportInfo) Get_provideInfo(_env *LnsEnv) *FrontInterface_ModuleProvideInfo{ return self.provideInfo }
func (self *FrontInterface_ExportInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo{ return self.processInfo }
func (self *FrontInterface_ExportInfo) Get_globalSymbolList(_env *LnsEnv) *LnsList2_[*Ast_SymbolInfo]{ return self.globalSymbolList }
func (self *FrontInterface_ExportInfo) Get_importedAliasMap(_env *LnsEnv) *LnsMap2_[*Ast_TypeInfo,*Ast_AliasTypeInfo]{ return self.importedAliasMap }
func (self *FrontInterface_ExportInfo) Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId{ return self.moduleId }
func (self *FrontInterface_ExportInfo) Get_fullName(_env *LnsEnv) string{ return self.fullName }
func (self *FrontInterface_ExportInfo) Get_assignName(_env *LnsEnv) string{ return self.assignName }
func (self *FrontInterface_ExportInfo) Get_streamName(_env *LnsEnv) string{ return self.streamName }
func (self *FrontInterface_ExportInfo) Get_importId2localTypeInfoMap(_env *LnsEnv) *LnsMap2_[LnsInt,*Ast_TypeInfo]{ return self.importId2localTypeInfoMap }
func (self *FrontInterface_ExportInfo) Set_importId2localTypeInfoMap(_env *LnsEnv, arg1 *LnsMap2_[LnsInt,*Ast_TypeInfo]){ self.importId2localTypeInfoMap = arg1 }
// 171: DeclConstr
func (self *FrontInterface_ExportInfo) InitFrontInterface_ExportInfo(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,provideInfo *FrontInterface_ModuleProvideInfo,processInfo *Ast_ProcessInfo,globalSymbolList *LnsList2_[*Ast_SymbolInfo],importedAliasMap *LnsMap2_[*Ast_TypeInfo,*Ast_AliasTypeInfo],moduleId *FrontInterface_ModuleId,fullName string,assignName string,streamName string,idMap *LnsMap2_[*Ast_TypeInfo,LnsInt]) {
    self.moduleTypeInfo = moduleTypeInfo
    self.provideInfo = provideInfo
    self.processInfo = processInfo
    self.globalSymbolList = globalSymbolList
    self.importedAliasMap = importedAliasMap
    self.moduleId = moduleId
    self.fullName = fullName
    self.assignName = assignName
    self.streamName = streamName
    var importId2localTypeInfoMap *LnsMap2_[LnsInt,*Ast_TypeInfo]
    importId2localTypeInfoMap = NewLnsMap2_[LnsInt,*Ast_TypeInfo]( map[LnsInt]*Ast_TypeInfo{})
    for _typeInfo, _importId := range( idMap.Items ) {
        typeInfo := _typeInfo
        importId := _importId
        importId2localTypeInfoMap.Set(importId,typeInfo)
    }
    self.importId2localTypeInfoMap = importId2localTypeInfoMap
}


// declaration Class -- ModuleInfo
type FrontInterface_ModuleInfoMtd interface {
    Assign(_env *LnsEnv, arg1 string) *FrontInterface_ExportInfo
    GetTypeInfo(_env *LnsEnv, arg1 LnsInt) LnsAny
    Get_assignName(_env *LnsEnv) string
    Get_exportInfo(_env *LnsEnv) *FrontInterface_ExportInfo
    Get_fullName(_env *LnsEnv) string
    Get_globalSymbolList(_env *LnsEnv) *LnsList2_[*Ast_SymbolInfo]
    Get_importId2localTypeInfoMap(_env *LnsEnv) *LnsMap2_[LnsInt,*Ast_TypeInfo]
    Get_importedAliasMap(_env *LnsEnv) *LnsMap2_[*Ast_TypeInfo,*Ast_AliasTypeInfo]
    Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId
    Get_modulePath(_env *LnsEnv) string
    Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_provideInfo(_env *LnsEnv) *FrontInterface_ModuleProvideInfo
    Get_streamName(_env *LnsEnv) string
    Set_importId2localTypeInfoMap(_env *LnsEnv, arg1 *LnsMap2_[LnsInt,*Ast_TypeInfo])
}
type FrontInterface_ModuleInfo struct {
    exportInfo *FrontInterface_ExportInfo
    FP FrontInterface_ModuleInfoMtd
}
func FrontInterface_ModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_ModuleInfo).FP
}
func FrontInterface_ModuleInfo_toSlice(slice []LnsAny) []*FrontInterface_ModuleInfo {
    ret := make([]*FrontInterface_ModuleInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(FrontInterface_ModuleInfoDownCast).ToFrontInterface_ModuleInfo()
    }
    return ret
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
func NewFrontInterface_ModuleInfo(_env *LnsEnv, arg1 *FrontInterface_ExportInfo) *FrontInterface_ModuleInfo {
    obj := &FrontInterface_ModuleInfo{}
    obj.FP = obj
    obj.InitFrontInterface_ModuleInfo(_env, arg1)
    return obj
}
func (self *FrontInterface_ModuleInfo) Get_exportInfo(_env *LnsEnv) *FrontInterface_ExportInfo{ return self.exportInfo }
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Assign(_env *LnsEnv, arg1 string) *FrontInterface_ExportInfo {
    return self.exportInfo. FP.Assign( _env, arg1)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) GetTypeInfo(_env *LnsEnv, arg1 LnsInt) LnsAny {
    return self.exportInfo. FP.GetTypeInfo( _env, arg1)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_assignName(_env *LnsEnv) string {
    return self.exportInfo. FP.Get_assignName( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_fullName(_env *LnsEnv) string {
    return self.exportInfo. FP.Get_fullName( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_globalSymbolList(_env *LnsEnv) *LnsList2_[*Ast_SymbolInfo] {
    return self.exportInfo. FP.Get_globalSymbolList( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_importId2localTypeInfoMap(_env *LnsEnv) *LnsMap2_[LnsInt,*Ast_TypeInfo] {
    return self.exportInfo. FP.Get_importId2localTypeInfoMap( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_importedAliasMap(_env *LnsEnv) *LnsMap2_[*Ast_TypeInfo,*Ast_AliasTypeInfo] {
    return self.exportInfo. FP.Get_importedAliasMap( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId {
    return self.exportInfo. FP.Get_moduleId( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_modulePath(_env *LnsEnv) string {
    return self.exportInfo. FP.Get_modulePath( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.exportInfo. FP.Get_moduleTypeInfo( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.exportInfo. FP.Get_processInfo( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_provideInfo(_env *LnsEnv) *FrontInterface_ModuleProvideInfo {
    return self.exportInfo. FP.Get_provideInfo( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Get_streamName(_env *LnsEnv) string {
    return self.exportInfo. FP.Get_streamName( _env)
}
// advertise -- 223
func (self *FrontInterface_ModuleInfo) Set_importId2localTypeInfoMap(_env *LnsEnv, arg1 *LnsMap2_[LnsInt,*Ast_TypeInfo]) {
self.exportInfo. FP.Set_importId2localTypeInfoMap( _env, arg1)
}
// 229: DeclConstr
func (self *FrontInterface_ModuleInfo) InitFrontInterface_ModuleInfo(_env *LnsEnv, exportInfo *FrontInterface_ExportInfo) {
    self.exportInfo = exportInfo
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
func FrontInterface_ModuleMeta_toSlice(slice []LnsAny) []*FrontInterface_ModuleMeta {
    ret := make([]*FrontInterface_ModuleMeta, len(slice))
    for index, val := range slice {
        ret[index] = val.(FrontInterface_ModuleMetaDownCast).ToFrontInterface_ModuleMeta()
    }
    return ret
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
func FrontInterface_ImportModuleInfo_toSlice(slice []LnsAny) []*FrontInterface_ImportModuleInfo {
    ret := make([]*FrontInterface_ImportModuleInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(FrontInterface_ImportModuleInfoDownCast).ToFrontInterface_ImportModuleInfo()
    }
    return ret
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
// 252: DeclConstr
func (self *FrontInterface_ImportModuleInfo) InitFrontInterface_ImportModuleInfo(_env *LnsEnv) {
    self.orderedSet = NewUtil_OrderedSet(_env)
}


type FrontInterface_ModuleLoader interface {
        GetExportInfo(_env *LnsEnv) LnsAny
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

// declaration Class -- FrontAccessor
type FrontInterface_FrontAccessorMtd interface {
    Error(_env *LnsEnv, arg1 string)
    GetLuaModulePath(_env *LnsEnv, arg1 string, arg2 LnsAny)(string, LnsAny, string)
    LoadFromLnsTxt(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string) LnsAny
    LoadMeta(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string, arg4 LnsAny, arg5 FrontInterface_ModuleLoader) LnsAny
    LoadModule(_env *LnsEnv, arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
    SearchModule(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny) LnsAny
}
type FrontInterface_FrontAccessor struct {
    frontIF FrontInterface_frontInterface
    FP FrontInterface_FrontAccessorMtd
}
func FrontInterface_FrontAccessor2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*FrontInterface_FrontAccessor).FP
}
func FrontInterface_FrontAccessor_toSlice(slice []LnsAny) []*FrontInterface_FrontAccessor {
    ret := make([]*FrontInterface_FrontAccessor, len(slice))
    for index, val := range slice {
        ret[index] = val.(FrontInterface_FrontAccessorDownCast).ToFrontInterface_FrontAccessor()
    }
    return ret
}
type FrontInterface_FrontAccessorDownCast interface {
    ToFrontInterface_FrontAccessor() *FrontInterface_FrontAccessor
}
func FrontInterface_FrontAccessorDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(FrontInterface_FrontAccessorDownCast)
    if ok { return work.ToFrontInterface_FrontAccessor() }
    return nil
}
func (obj *FrontInterface_FrontAccessor) ToFrontInterface_FrontAccessor() *FrontInterface_FrontAccessor {
    return obj
}
func NewFrontInterface_FrontAccessor(_env *LnsEnv, arg1 FrontInterface_frontInterface) *FrontInterface_FrontAccessor {
    obj := &FrontInterface_FrontAccessor{}
    obj.FP = obj
    obj.InitFrontInterface_FrontAccessor(_env, arg1)
    return obj
}
func (self *FrontInterface_FrontAccessor) InitFrontInterface_FrontAccessor(_env *LnsEnv, arg1 FrontInterface_frontInterface) {
    self.frontIF = arg1
}

// declaration Class -- DummyFront
type frontInterface_DummyFrontMtd interface {
    Error(_env *LnsEnv, arg1 string)
    GetLuaModulePath(_env *LnsEnv, arg1 string, arg2 LnsAny)(string, LnsAny, string)
    LoadFromLnsTxt(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string) LnsAny
    LoadMeta(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string, arg4 LnsAny, arg5 FrontInterface_ModuleLoader) LnsAny
    LoadModule(_env *LnsEnv, arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
    SearchModule(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny) LnsAny
}
type frontInterface_DummyFront struct {
    FP frontInterface_DummyFrontMtd
}
func frontInterface_DummyFront2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*frontInterface_DummyFront).FP
}
func frontInterface_DummyFront_toSlice(slice []LnsAny) []*frontInterface_DummyFront {
    ret := make([]*frontInterface_DummyFront, len(slice))
    for index, val := range slice {
        ret[index] = val.(frontInterface_DummyFrontDownCast).TofrontInterface_DummyFront()
    }
    return ret
}
type frontInterface_DummyFrontDownCast interface {
    TofrontInterface_DummyFront() *frontInterface_DummyFront
}
func frontInterface_DummyFrontDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(frontInterface_DummyFrontDownCast)
    if ok { return work.TofrontInterface_DummyFront() }
    return nil
}
func (obj *frontInterface_DummyFront) TofrontInterface_DummyFront() *frontInterface_DummyFront {
    return obj
}
func NewfrontInterface_DummyFront(_env *LnsEnv) *frontInterface_DummyFront {
    obj := &frontInterface_DummyFront{}
    obj.FP = obj
    obj.InitfrontInterface_DummyFront(_env)
    return obj
}
func (self *frontInterface_DummyFront) InitfrontInterface_DummyFront(_env *LnsEnv) {
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
    FrontInterface_dummyFront = NewfrontInterface_DummyFront(_env).FP
}
func init() {
    init_frontInterface = false
}
