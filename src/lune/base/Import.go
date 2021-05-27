// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Import bool
var Import__mod__ string
// for 1029
func Import_convExp5548(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1115
func Import_convExp6014(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 844
func Import_convExp4495(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 847
func Import_convExp4518(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 851
func Import_convExp4542(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 855
func Import_convExp4564(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 862
func Import_convExp4611(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 865
func Import_convExp4634(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 868
func Import_convExp4657(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 871
func Import_convExp4680(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 874
func Import_convExp4703(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 877
func Import_convExp4726(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 880
func Import_convExp4749(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 883
func Import_convExp4772(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 151
func Import_convExp381(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 165
func Import_convExp459(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 205
func Import_convExp701(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 231
func Import_convExp850(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 283
func Import_convExp1117(arg1 []LnsAny) *Ast_AlternateTypeInfo {
    return Lns_getFromMulti( arg1, 0 ).(*Ast_AlternateTypeInfo)
}
// for 303
func Import_convExp1236(arg1 []LnsAny) (*Ast_GenericTypeInfo, *Ast_Scope) {
    return Lns_getFromMulti( arg1, 0 ).(*Ast_GenericTypeInfo), Lns_getFromMulti( arg1, 1 ).(*Ast_Scope)
}
// for 355
func Import_convExp1495(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 374
func Import_convExp1603(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 457
func Import_convExp2076(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 471
func Import_convExp2145(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 725
func Import_convExp3841(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 921
func Import_convExp4965(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 967
func Import_convExp5207(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1033
func Import_convExp5536(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}










// declaration Class -- Import
type Import_ImportMtd interface {
    Get_importModule2ModuleInfo(_env *LnsEnv) *LnsMap
    ProcessImport(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string) *FrontInterface_ModuleInfo
    processImportFromFile(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 string, arg5 string, arg6 *LnsList, arg7 LnsInt) *FrontInterface_ModuleInfo
    processImportMain(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsInt) *FrontInterface_ModuleInfo
}
type Import_Import struct {
    transUnitIF TransUnitIF_TransUnitIF
    importModuleInfo *FrontInterface_ImportModuleInfo
    moduleType *Ast_TypeInfo
    builtinFunc *Builtin_BuiltinFuncType
    globalScope *Ast_Scope
    macroCtrl *Macro_MacroCtrl
    typeNameCtrl *Ast_TypeNameCtrl
    importModule2ModuleInfo *LnsMap
    importedAliasMap *LnsMap
    importModuleName2ModuleInfo *LnsMap
    validMutControl bool
    FP Import_ImportMtd
}
func Import_Import2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import_Import).FP
}
type Import_ImportDownCast interface {
    ToImport_Import() *Import_Import
}
func Import_ImportDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import_ImportDownCast)
    if ok { return work.ToImport_Import() }
    return nil
}
func (obj *Import_Import) ToImport_Import() *Import_Import {
    return obj
}
func NewImport_Import(_env *LnsEnv, arg1 TransUnitIF_TransUnitIF, arg2 *FrontInterface_ImportModuleInfo, arg3 *Ast_TypeInfo, arg4 *Builtin_BuiltinFuncType, arg5 *Ast_Scope, arg6 *Macro_MacroCtrl, arg7 *Ast_TypeNameCtrl, arg8 *LnsMap, arg9 bool) *Import_Import {
    obj := &Import_Import{}
    obj.FP = obj
    obj.InitImport_Import(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Import_Import) Get_importModule2ModuleInfo(_env *LnsEnv) *LnsMap{ return self.importModule2ModuleInfo }
// 67: DeclConstr
func (self *Import_Import) InitImport_Import(_env *LnsEnv, transUnitIF TransUnitIF_TransUnitIF,importModuleInfo *FrontInterface_ImportModuleInfo,moduleType *Ast_TypeInfo,builtinFunc *Builtin_BuiltinFuncType,globalScope *Ast_Scope,macroCtrl *Macro_MacroCtrl,typeNameCtrl *Ast_TypeNameCtrl,importedAliasMap *LnsMap,validMutControl bool) {
    self.validMutControl = validMutControl
    
    self.transUnitIF = transUnitIF
    
    self.importModuleInfo = importModuleInfo
    
    self.moduleType = moduleType
    
    self.builtinFunc = builtinFunc
    
    self.globalScope = globalScope
    
    self.macroCtrl = macroCtrl
    
    self.typeNameCtrl = typeNameCtrl
    
    self.importedAliasMap = importedAliasMap
    
    self.importModule2ModuleInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.importModuleName2ModuleInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
}


// 753: decl @lune.@base.@Import.Import.processImportFromFile
func (self *Import_Import) processImportFromFile(_env *LnsEnv, processInfo *Ast_ProcessInfo,lnsPath string,metaInfoStem LnsAny,orgModulePath string,modulePath string,nameList *LnsList,depth LnsInt) *FrontInterface_ModuleInfo {
    __func__ := "@lune.@base.@Import.Import.processImportFromFile"
    var metaInfo *Lns_luaValue
    metaInfo = metaInfoStem.(*Lns_luaValue)
    Log_log(_env, Log_Level__Info, __func__, 759, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.LuaVM.String_format("%s processing", []LnsAny{orgModulePath})
    }))
    
    var dependLibId2DependInfo *LnsMap
    dependLibId2DependInfo = NewLnsMap( map[LnsAny]LnsAny{})
    {
        _exp4099 := metaInfo.GetAt( "__dependModuleMap" ).(*Lns_luaValue)
        _sorted4099 := _env.LuaVM.SortMapKeyList( _exp4099 )
        _index4099, _key4099 := _sorted4099.Get1stFromMap()
        for _index4099 != nil {
            dependName := _key4099.(string)
            dependInfo := _exp4099.GetAt( _key4099 ).(*Lns_luaValue)
            var workProcessInfo *Ast_ProcessInfo
            workProcessInfo = processInfo.FP.NewUser(_env)
            workProcessInfo.FP.SwitchIdProvier(_env, Ast_IdType__Ext)
            var moduleInfo *FrontInterface_ModuleInfo
            moduleInfo = self.FP.processImportMain(_env, workProcessInfo, dependName, depth + 1)
            workProcessInfo.FP.SwitchIdProvier(_env, Ast_IdType__Base)
            var typeId LnsInt
            typeId = Lns_forceCastInt((Lns_unwrap( dependInfo.GetAt("typeId"))))
            dependLibId2DependInfo.Set(typeId,moduleInfo)
            _index4099, _key4099 = _sorted4099.NextFromMap( _index4099 )
        }
    }
    var typeId2TypeInfo *LnsMap
    typeId2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    typeId2TypeInfo.Set(Ast_rootTypeId,Ast_headTypeInfo)
    var typeId2Scope *LnsMap
    typeId2Scope = NewLnsMap( map[LnsAny]LnsAny{})
    typeId2Scope.Set(Ast_rootTypeId,self.transUnitIF.Get_scope(_env))
    typeId2TypeInfo.Set(self.builtinFunc.Lnsthread_.FP.Get_typeId(_env).Id,self.builtinFunc.Lnsthread_)
    {
        _exp4213 := metaInfo.GetAt( "__dependIdMap" ).(*Lns_luaValue)
        _key4213, _val4213 := _exp4213.Get1stFromMap()
        for _key4213 != nil {
            typeId := _key4213.(LnsInt)
            dependIdInfo := _val4213.(*Lns_luaValue)
            var dependInfo *FrontInterface_ModuleInfo
            dependInfo = Lns_unwrap( dependLibId2DependInfo.Get(Lns_unwrap( dependIdInfo.GetAt(1)).(LnsInt))).(*FrontInterface_ModuleInfo)
            var typeInfo *Ast_TypeInfo
            typeInfo = Lns_unwrap( dependInfo.FP.GetTypeInfo(_env, Lns_unwrap( dependIdInfo.GetAt(2)).(LnsInt))).(*Ast_TypeInfo)
            typeId2TypeInfo.Set(typeId,typeInfo)
            _key4213, _val4213 = _exp4213.NextFromMap( _key4213 )
        }
    }
    var moduleTypeInfo *Ast_TypeInfo
    moduleTypeInfo = Ast_headTypeInfo
    for _index, _moduleName := range( nameList.Items ) {
        index := _index + 1
        moduleName := _moduleName.(string)
        var mutable bool
        mutable = false
        if index == nameList.Len(){
            mutable = metaInfo.GetAt( "__moduleMutable" ).(bool)
            
        }
        moduleTypeInfo = self.transUnitIF.PushModule(_env, processInfo, true, moduleName, mutable)
        
        var typeId LnsInt
        typeId = Lns_unwrap( metaInfo.GetAt( "__moduleHierarchy" ).(*Lns_luaValue).GetAt(nameList.Len() - index + 1)).(LnsInt)
        typeId2TypeInfo.Set(typeId,moduleTypeInfo)
        typeId2Scope.Set(typeId,self.transUnitIF.Get_scope(_env))
    }
    for range( nameList.Items ) {
        self.transUnitIF.PopModule(_env)
    }
    for _, _symbolInfo := range( Ast_getSym2builtInTypeMap(_env).Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        typeId2TypeInfo.Set(symbolInfo.FP.Get_typeInfo(_env).FP.Get_typeId(_env).Id,symbolInfo.FP.Get_typeInfo(_env))
    }
    for _, _builtinTypeInfo := range( Ast_getBuiltInTypeIdMap(_env).Items ) {
        builtinTypeInfo := _builtinTypeInfo.(Ast_BuiltinTypeInfoDownCast).ToAst_BuiltinTypeInfo()
        var typeInfo *Ast_TypeInfo
        typeInfo = builtinTypeInfo.FP.Get_typeInfo(_env)
        typeId2TypeInfo.Set(typeInfo.FP.Get_typeId(_env).Id,typeInfo)
    }
    var newId2OldIdMap *LnsMap
    newId2OldIdMap = NewLnsMap( map[LnsAny]LnsAny{})
    var _typeInfoList *LnsList
    _typeInfoList = NewLnsList([]LnsAny{})
    var id2atomMap *LnsMap
    id2atomMap = NewLnsMap( map[LnsAny]LnsAny{})
    var _typeInfoNormalList *LnsList
    _typeInfoNormalList = NewLnsList([]LnsAny{})
    {
        _exp4854 := metaInfo.GetAt( "__typeInfoList" ).(*Lns_luaValue)
        _key4854, _val4854 := _exp4854.Get1stFromMap()
        for _key4854 != nil {
            atomInfoLua := _val4854.(*Lns_luaValue)
            var workAtomInfo LnsAny
            
            {
                _workAtomInfo := _env.LuaVM.ExpandLuavalMap(atomInfoLua)
                if _workAtomInfo == nil{
                    self.transUnitIF.Error(_env, "illegal atomInfo")
                } else {
                    workAtomInfo = _workAtomInfo
                }
            }
            var atomInfo *LnsMap
            atomInfo = workAtomInfo.(*LnsMap)
            {
                _skind := atomInfo.Get("skind")
                if !Lns_IsNil( _skind ) {
                    skind := _skind
                    var actInfo LnsAny
                    actInfo = nil
                    var mess LnsAny
                    mess = nil
                    var kind LnsInt
                    kind = Lns_unwrap( Ast_SerializeKind__from(_env, Lns_forceCastInt(skind))).(LnsInt)
                    if _switch4775 := kind; _switch4775 == Ast_SerializeKind__Enum {
                        actInfo, mess = Import_convExp4495(Lns_2DDD(Import__TypeInfoEnum__fromMap_1781_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__Alge {
                        actInfo, mess = Import_convExp4518(Lns_2DDD(Import__TypeInfoAlge__fromMap_1864_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__Module {
                        actInfo, mess = Import_convExp4542(Lns_2DDD(Import__TypeInfoModule__fromMap_1523_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__Normal {
                        var workInfo LnsAny
                        workInfo, mess = Import__TypeInfoNormal__fromMap_1718_(_env, atomInfo,nil)
                        
                        if workInfo != nil{
                            workInfo_525 := workInfo.(*Import__TypeInfoNormal)
                            _typeInfoNormalList.Insert(Import__TypeInfoNormal2Stem(workInfo_525))
                        }
                        actInfo = Import__TypeInfoDownCastF(workInfo)
                        
                    } else if _switch4775 == Ast_SerializeKind__Nilable {
                        actInfo, mess = Import_convExp4611(Lns_2DDD(Import__TypeInfoNilable__fromMap_1220_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__Alias {
                        actInfo, mess = Import_convExp4634(Lns_2DDD(Import__TypeInfoAlias__fromMap_1258_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__DDD {
                        actInfo, mess = Import_convExp4657(Lns_2DDD(Import__TypeInfoDDD__fromMap_1287_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__Alternate {
                        actInfo, mess = Import_convExp4680(Lns_2DDD(Import__TypeInfoAlternate__fromMap_1332_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__Generic {
                        actInfo, mess = Import_convExp4703(Lns_2DDD(Import__TypeInfoGeneric__fromMap_1379_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__Modifier {
                        actInfo, mess = Import_convExp4726(Lns_2DDD(Import__TypeInfoModifier__fromMap_1469_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__Box {
                        actInfo, mess = Import_convExp4749(Lns_2DDD(Import__TypeInfoBox__fromMap_1408_(_env, atomInfo,nil)))
                        
                    } else if _switch4775 == Ast_SerializeKind__Ext {
                        actInfo, mess = Import_convExp4772(Lns_2DDD(Import__TypeInfoExt__fromMap_1438_(_env, atomInfo,nil)))
                        
                    }
                    if actInfo != nil{
                        actInfo_535 := actInfo.(*Import__TypeInfo)
                        _typeInfoList.Insert(Import__TypeInfo2Stem(actInfo_535))
                        id2atomMap.Set(actInfo_535.TypeId,actInfo_535)
                    } else {
                        for _key, _val := range( atomInfo.Items ) {
                            key := _key.(string)
                            val := _val
                            Util_errorLog(_env, _env.LuaVM.String_format("table: %s:%s", []LnsAny{key, val}))
                        }
                        if mess != nil{
                            mess_541 := mess.(string)
                            Util_errorLog(_env, mess_541)
                        }
                        Util_err(_env, _env.LuaVM.String_format("_TypeInfo.%s._fromMap error", []LnsAny{Ast_SerializeKind_getTxt( kind)}))
                    }
                }
            }
            _key4854, _val4854 = _exp4854.NextFromMap( _key4854 )
        }
    }
    var orgId2MacroTypeInfo *LnsMap
    orgId2MacroTypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    var lazyModuleSet *LnsSet
    lazyModuleSet = NewLnsSet([]LnsAny{})
    {
        _exp4894 := metaInfo.GetAt( "__lazyModuleList" ).(*Lns_luaValue)
        _key4894, _val4894 := _exp4894.Get1stFromMap()
        for _key4894 != nil {
            typeId := _val4894.(LnsInt)
            lazyModuleSet.Add(typeId)
            _key4894, _val4894 = _exp4894.NextFromMap( _key4894 )
        }
    }
    var modifier *TransUnitIF_Modifier
    modifier = NewTransUnitIF_Modifier(_env, self.validMutControl, processInfo)
    var importParam *Import_ImportParam
    importParam = NewImport_ImportParam(_env, self.transUnitIF.GetLatestPos(_env), modifier, processInfo, typeId2Scope, typeId2TypeInfo, NewLnsMap( map[LnsAny]LnsAny{}), lazyModuleSet, metaInfo, self.transUnitIF.Get_scope(_env), moduleTypeInfo, Ast_ScopeAccess__Normal, id2atomMap, dependLibId2DependInfo)
    for _, _atomInfo := range( _typeInfoList.Items ) {
        atomInfo := _atomInfo.(Import__TypeInfoDownCast).ToImport__TypeInfo()
        var newTypeInfo LnsAny
        var errMess LnsAny
        newTypeInfo,errMess = atomInfo.FP.CreateTypeInfoCache(_env, importParam)
        {
            __exp := errMess
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(string)
                Util_err(_env, _env.LuaVM.String_format("Failed to createType -- %s: %s(%d): %s", []LnsAny{orgModulePath, Ast_SerializeKind_getTxt( atomInfo.Skind), atomInfo.TypeId, _exp}))
            }
        }
        if newTypeInfo != nil{
            newTypeInfo_555 := newTypeInfo.(*Ast_TypeInfo)
            if newTypeInfo_555.FP.Get_kind(_env) == Ast_TypeInfoKind__Macro{
                orgId2MacroTypeInfo.Set(atomInfo.TypeId,newTypeInfo_555)
            }
            if newTypeInfo_555.FP.Get_kind(_env) == Ast_TypeInfoKind__Set{
            }
            if newTypeInfo_555.FP.Get_accessMode(_env) == Ast_AccessMode__Global{
                if _switch5133 := newTypeInfo_555.FP.Get_kind(_env); _switch5133 == Ast_TypeInfoKind__IF || _switch5133 == Ast_TypeInfoKind__Class {
                    self.globalScope.FP.AddClass(_env, processInfo, newTypeInfo_555.FP.Get_rawTxt(_env), nil, newTypeInfo_555)
                } else if _switch5133 == Ast_TypeInfoKind__Func {
                    self.globalScope.FP.AddFunc(_env, processInfo, nil, newTypeInfo_555, Ast_AccessMode__Global, newTypeInfo_555.FP.Get_staticFlag(_env), Ast_TypeInfo_isMut(_env, newTypeInfo_555))
                } else if _switch5133 == Ast_TypeInfoKind__Enum {
                    self.globalScope.FP.AddEnum(_env, processInfo, Ast_AccessMode__Global, newTypeInfo_555.FP.Get_rawTxt(_env), nil, newTypeInfo_555)
                } else if _switch5133 == Ast_TypeInfoKind__Nilable {
                } else {
                    Util_err(_env, _env.LuaVM.String_format("%s: not support kind -- %s", []LnsAny{__func__, Ast_TypeInfoKind_getTxt( newTypeInfo_555.FP.Get_kind(_env))}))
                }
            }
        }
    }
    for _, _atomInfo := range( _typeInfoNormalList.Items ) {
        atomInfo := _atomInfo.(Import__TypeInfoNormalDownCast).ToImport__TypeInfoNormal()
        if atomInfo.Children.Len() > 0{
            importParam.FP.GetTypeInfo(_env, atomInfo.TypeId)
            var scope *Ast_Scope
            scope = Lns_unwrap( typeId2Scope.Get(atomInfo.TypeId)).(*Ast_Scope)
            for _, _childId := range( atomInfo.Children.Items ) {
                childId := _childId.(Import__IdInfoDownCast).ToImport__IdInfo()
                var typeInfo *Ast_TypeInfo
                
                {
                    _typeInfo := Import_convExp5207(Lns_2DDD(importParam.FP.GetTypeInfoFrom(_env, childId)))
                    if _typeInfo == nil{
                        Util_err(_env, _env.LuaVM.String_format("not found childId -- %s, %d, %s(%d)", []LnsAny{orgModulePath, childId.Id, atomInfo.Txt, atomInfo.TypeId}))
                    } else {
                        typeInfo = _typeInfo.(*Ast_TypeInfo)
                    }
                }
                var symbolKind LnsInt
                symbolKind = Ast_SymbolKind__Typ
                var addFlag bool
                addFlag = true
                if _switch5280 := typeInfo.FP.Get_kind(_env); _switch5280 == Ast_TypeInfoKind__Func {
                    symbolKind = Ast_SymbolKind__Fun
                    
                } else if _switch5280 == Ast_TypeInfoKind__Form || _switch5280 == Ast_TypeInfoKind__FormFunc {
                    symbolKind = Ast_SymbolKind__Typ
                    
                } else if _switch5280 == Ast_TypeInfoKind__Method {
                    symbolKind = Ast_SymbolKind__Mtd
                    
                } else if _switch5280 == Ast_TypeInfoKind__Class || _switch5280 == Ast_TypeInfoKind__Module {
                    symbolKind = Ast_SymbolKind__Typ
                    
                } else if _switch5280 == Ast_TypeInfoKind__Enum {
                    addFlag = false
                    
                }
                if addFlag{
                    scope.FP.Add(_env, processInfo, symbolKind, false, typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Func, typeInfo.FP.GetTxt(_env, nil, nil, nil), nil, typeInfo, typeInfo.FP.Get_accessMode(_env), typeInfo.FP.Get_staticFlag(_env), typeInfo.FP.Get_mutMode(_env), true, false)
                }
            }
        }
    }
    for _typeId, _typeInfo := range( typeId2TypeInfo.Items ) {
        typeId := _typeId.(LnsInt)
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        newId2OldIdMap.Set(typeInfo,typeId)
    }
    var registMember func(_env *LnsEnv, classTypeId LnsInt)
    registMember = func(_env *LnsEnv, classTypeId LnsInt) {
        __func__ := "@lune.@base.@Import.Import.processImportFromFile.registMember"
        if Lns_isCondTrue( metaInfo.GetAt( "__dependIdMap" ).(*Lns_luaValue).GetAt(classTypeId)){
            return 
        }
        var classTypeInfo *Ast_TypeInfo
        classTypeInfo = Lns_unwrap( typeId2TypeInfo.Get(classTypeId)).(*Ast_TypeInfo)
        if _switch5755 := (classTypeInfo.FP.Get_kind(_env)); _switch5755 == Ast_TypeInfoKind__Class || _switch5755 == Ast_TypeInfoKind__ExtModule {
            var scope *Ast_Scope
            scope = Lns_unwrap( typeId2Scope.Get(classTypeId)).(*Ast_Scope)
            self.transUnitIF.PushClassScope(_env, self.transUnitIF.GetLatestPos(_env), classTypeInfo, scope)
            {
                __exp := metaInfo.GetAt( "__typeId2ClassInfoMap" ).(*Lns_luaValue).GetAt(classTypeId)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Lns_luaValue)
                    var classInfo LnsAny
                    
                    {
                        _classInfo := _env.LuaVM.ExpandLuavalMap(_exp)
                        if _classInfo == nil{
                            self.transUnitIF.Error(_env, "illegal val")
                        } else {
                            classInfo = _classInfo
                        }
                    }
                    for _fieldName, _fieldInfo := range( classInfo.(*LnsMap).Items ) {
                        fieldName := _fieldName.(string)
                        fieldInfo := _fieldInfo.(*LnsMap)
                        {
                            _typeId := Import_convExp5548(Lns_2DDD(Import__IdInfo__fromStem_1169_(_env, _env.LuaVM.ExpandLuavalMap(fieldInfo.Get("typeId")),nil)))
                            if !Lns_IsNil( _typeId ) {
                                typeId := _typeId.(*Import__IdInfo)
                                var fieldTypeInfo *Ast_TypeInfo
                                fieldTypeInfo = Lns_unwrap( Lns_car(importParam.FP.GetTypeInfoFrom(_env, typeId))).(*Ast_TypeInfo)
                                _ = Import_convExp5536(Lns_2DDD(self.transUnitIF.Get_scope(_env).FP.AddMember(_env, processInfo, fieldName, nil, fieldTypeInfo, Lns_unwrap( Ast_AccessMode__from(_env, Lns_forceCastInt((Lns_unwrap( fieldInfo.Get("accessMode")))))).(LnsInt), _env.PopVal( _env.IncStack() ||
                                    _env.SetStackVal( fieldInfo.Get("staticFlag")) &&
                                    _env.SetStackVal( true) ||
                                    _env.SetStackVal( false) ).(bool), Lns_unwrap( Ast_MutMode__from(_env, Lns_forceCastInt((Lns_unwrap( fieldInfo.Get("mutMode")))))).(LnsInt))))
                            } else {
                                self.transUnitIF.Error(_env, "not found fieldInfo.typeId")
                            }
                        }
                    }
                } else {
                    self.transUnitIF.Error(_env, _env.LuaVM.String_format("not found class -- %s: %d, %s", []LnsAny{orgModulePath, classTypeId, classTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        } else if _switch5755 == Ast_TypeInfoKind__Module {
            self.transUnitIF.PushModule(_env, processInfo, true, classTypeInfo.FP.GetTxt(_env, nil, nil, nil), Ast_TypeInfo_isMut(_env, classTypeInfo))
            Log_log(_env, Log_Level__Debug, __func__, 1054, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.LuaVM.String_format("push module -- %s, %s, %d, %d, %d", []LnsAny{classTypeInfo.FP.GetTxt(_env, nil, nil, nil), _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(self.transUnitIF.Get_scope(_env).FP.Get_ownerTypeInfo(_env)) && 
                    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetFullName(_env, Ast_defaultTypeNameCtrl, self.transUnitIF.Get_scope(_env).FP, false)})/* 1:67 */)) ||
                    _env.SetStackVal( "nil") ).(string), _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(self.transUnitIF.Get_scope(_env).FP.Get_ownerTypeInfo(_env)) && 
                    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_typeId(_env)})&&
                    _env.NilAccPush(_env.NilAccPop().(*Ast_IdInfo).Id))) ||
                    _env.SetStackVal( -1) ).(LnsInt), classTypeInfo.FP.Get_typeId(_env).Id, self.transUnitIF.Get_scope(_env).FP.Get_parent(_env).FP.Get_scopeId(_env)})
            }))
            
        }
        for _, _child := range( classTypeInfo.FP.Get_children(_env).Items ) {
            child := _child.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( child.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( child.FP.Get_kind(_env) == Ast_TypeInfoKind__ExtModule) ||
                _env.SetStackVal( child.FP.Get_kind(_env) == Ast_TypeInfoKind__Module) ||
                _env.SetStackVal( child.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ).(bool){
                var oldId LnsAny
                oldId = newId2OldIdMap.Get(child)
                if Lns_isCondTrue( oldId){
                    registMember(_env, Lns_unwrap( oldId).(LnsInt))
                }
            }
        }
        if _switch5843 := classTypeInfo.FP.Get_kind(_env); _switch5843 == Ast_TypeInfoKind__Class || _switch5843 == Ast_TypeInfoKind__ExtModule {
            self.transUnitIF.PopClass(_env)
        } else if _switch5843 == Ast_TypeInfoKind__Module {
            self.transUnitIF.PopModule(_env)
        }
    }
    for _, _atomInfo := range( _typeInfoList.Items ) {
        atomInfo := _atomInfo.(Import__TypeInfoDownCast).ToImport__TypeInfo()
        {
            _workInfo := Import__TypeInfoNormalDownCastF(atomInfo.FP)
            if !Lns_IsNil( _workInfo ) {
                workInfo := _workInfo.(*Import__TypeInfoNormal)
                if workInfo.ParentId == Ast_rootTypeId{
                    registMember(_env, atomInfo.TypeId)
                }
            } else {
                {
                    _workInfo := Import__TypeInfoModuleDownCastF(atomInfo.FP)
                    if !Lns_IsNil( _workInfo ) {
                        workInfo := _workInfo.(*Import__TypeInfoModule)
                        if workInfo.ParentId == Ast_rootTypeId{
                            registMember(_env, atomInfo.TypeId)
                        }
                    }
                }
            }
        }
    }
    for _index, _moduleName := range( nameList.Items ) {
        index := _index + 1
        moduleName := _moduleName.(string)
        var mutable bool
        mutable = false
        if index == nameList.Len(){
            mutable = metaInfo.GetAt( "__moduleMutable" ).(bool)
            
        }
        self.transUnitIF.PushModule(_env, processInfo, true, moduleName, mutable)
    }
    {
        _exp6016 := metaInfo.GetAt( "__varName2InfoMap" ).(*Lns_luaValue)
        _key6016, _val6016 := _exp6016.Get1stFromMap()
        for _key6016 != nil {
            varName := _key6016.(string)
            varInfo := _val6016.(*Lns_luaValue)
            {
                _typeId := Import_convExp6014(Lns_2DDD(Import__IdInfo__fromStem_1169_(_env, _env.LuaVM.ExpandLuavalMap(varInfo.GetAt("typeId")),nil)))
                if !Lns_IsNil( _typeId ) {
                    typeId := _typeId.(*Import__IdInfo)
                    self.transUnitIF.Get_scope(_env).FP.AddStaticVar(_env, processInfo, false, true, varName, nil, Lns_unwrap( Lns_car(importParam.FP.GetTypeInfoFrom(_env, typeId))).(*Ast_TypeInfo), _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( varInfo.GetAt("mutable")) &&
                        _env.SetStackVal( Ast_MutMode__Mut) ||
                        _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt))
                } else {
                    self.transUnitIF.Error(_env, "illegal varInfo.typeId")
                }
            }
            _key6016, _val6016 = _exp6016.NextFromMap( _key6016 )
        }
    }
    var importedMacroInfoMap *LnsMap
    importedMacroInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    {
        _exp6065 := metaInfo.GetAt( "__macroName2InfoMap" ).(*Lns_luaValue)
        _key6065, _val6065 := _exp6065.Get1stFromMap()
        for _key6065 != nil {
            orgTypeId := _key6065.(LnsInt)
            macroInfoStem := _val6065
            self.macroCtrl.FP.ImportMacro(_env, processInfo, lnsPath, _env.LuaVM.ExpandLuavalMap(macroInfoStem), Lns_unwrap( orgId2MacroTypeInfo.Get(orgTypeId)).(*Ast_TypeInfo), typeId2TypeInfo, importedMacroInfoMap)
            _key6065, _val6065 = _exp6065.NextFromMap( _key6065 )
        }
    }
    var globalSymbolList *LnsList
    globalSymbolList = NewLnsList([]LnsAny{})
    for _, _symbolInfo := range( self.transUnitIF.Get_scope(_env).FP.Get_symbol2SymbolInfoMap(_env).Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if symbolInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Global{
            globalSymbolList.Insert(Ast_SymbolInfo2Stem(symbolInfo))
        }
    }
    for range( nameList.Items ) {
        self.transUnitIF.PopModule(_env)
    }
    if depth == 1{
        for _key, _val := range( importParam.ImportedAliasMap.Items ) {
            key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            val := _val.(Ast_AliasTypeInfoDownCast).ToAst_AliasTypeInfo()
            self.importedAliasMap.Set(key,val)
        }
    }
    var moduleProvideInfo *FrontInterface_ModuleProvideInfo
    moduleProvideInfo = NewFrontInterface_ModuleProvideInfo(_env, Lns_unwrap( typeId2TypeInfo.Get(metaInfo.GetAt( "__moduleTypeId" ).(LnsInt))).(*Ast_TypeInfo), Lns_unwrap( Ast_SymbolKind__from(_env, metaInfo.GetAt( "__moduleSymbolKind" ).(LnsInt))).(LnsInt), metaInfo.GetAt( "__moduleMutable" ).(bool))
    var exportInfo *Nodes_ExportInfo
    exportInfo = NewNodes_ExportInfo(_env, moduleTypeInfo, moduleProvideInfo, processInfo, globalSymbolList, importedMacroInfoMap)
    var moduleInfo *FrontInterface_ModuleInfo
    moduleInfo = NewFrontInterface_ModuleInfo(_env, orgModulePath, nameList.GetAt(nameList.Len()).(string), newId2OldIdMap, FrontInterface_ModuleId_createIdFromTxt(_env, metaInfo.GetAt( "__buildId" ).(string)), &exportInfo.FrontInterface_ExportInfo, importParam.ImportedAliasMap)
    return moduleInfo
}

// 1172: decl @lune.@base.@Import.Import.processImportMain
func (self *Import_Import) processImportMain(_env *LnsEnv, processInfo *Ast_ProcessInfo,modulePath string,depth LnsInt) *FrontInterface_ModuleInfo {
    __func__ := "@lune.@base.@Import.Import.processImportMain"
    var orgModulePath string
    orgModulePath = modulePath
    modulePath = FrontInterface_getLuaModulePath(_env, modulePath)
    
    Log_log(_env, Log_Level__Info, __func__, 1179, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.LuaVM.String_format("%s -> %s start", []LnsAny{self.moduleType.FP.GetTxt(_env, self.typeNameCtrl, nil, nil), orgModulePath})
    }))
    
    if Lns_op_not(self.importModuleInfo.FP.Add(_env, orgModulePath)){
        self.transUnitIF.Error(_env, _env.LuaVM.String_format("recursive import: %s -> %s", []LnsAny{self.importModuleInfo.FP.GetFull(_env), orgModulePath}))
    }
    {
        _moduleInfo := self.importModuleName2ModuleInfo.Get(modulePath)
        if !Lns_IsNil( _moduleInfo ) {
            moduleInfo := _moduleInfo.(*FrontInterface_ModuleInfo)
            Log_log(_env, Log_Level__Info, __func__, 1191, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.LuaVM.String_format("%s already", []LnsAny{orgModulePath})
            }))
            
            self.importModuleInfo.FP.Remove(_env)
            if depth == 1{
                self.importModule2ModuleInfo.Set(moduleInfo.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env),moduleInfo)
            }
            for _key, _val := range( moduleInfo.FP.Get_importedAliasMap(_env).Items ) {
                key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                val := _val.(Ast_AliasTypeInfoDownCast).ToAst_AliasTypeInfo()
                self.importedAliasMap.Set(key,val)
            }
            return moduleInfo
        }
    }
    var nameList *LnsList
    nameList = NewLnsList([]LnsAny{})
    {
        _form6510, _param6510, _prev6510 := _env.LuaVM.String_gmatch(modulePath, "[^%./:]+")
        for {
            _work6510 := _form6510.(*Lns_luaValue).Call( Lns_2DDD( _param6510, _prev6510 ) )
            _prev6510 = Lns_getFromMulti(_work6510,0)
            if Lns_IsNil( _prev6510 ) { break }
            txt := _prev6510.(string)
            nameList.Insert(txt)
        }
    }
    var moduleMeta *FrontInterface_ModuleMeta
    
    {
        _moduleMeta := FrontInterface_loadMeta(_env, self.importModuleInfo, orgModulePath)
        if _moduleMeta == nil{
            self.transUnitIF.Error(_env, "failed to load meta -- " + orgModulePath)
        } else {
            moduleMeta = _moduleMeta.(*FrontInterface_ModuleMeta)
        }
    }
    var moduleInfo *FrontInterface_ModuleInfo
    switch _exp6643 := moduleMeta.FP.Get_metaOrModule(_env).(type) {
    case *FrontInterface_MetaOrModule__Module:
    info := _exp6643.Val1
        moduleInfo = info
        
        {
            _exportInfo := Nodes_ExportInfoDownCastF(moduleInfo.FP.Get_exportInfo(_env).FP)
            if !Lns_IsNil( _exportInfo ) {
                exportInfo := _exportInfo.(*Nodes_ExportInfo)
                self.macroCtrl.FP.ImportMacroInfo(_env, exportInfo.FP.Get_typeId2DefMacroInfo(_env))
            }
        }
        for _, _globalSymbol := range( moduleInfo.FP.Get_exportInfo(_env).FP.Get_globalSymbolList(_env).Items ) {
            globalSymbol := _globalSymbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            self.transUnitIF.Get_scope(_env).FP.AddSymbolInfo(_env, processInfo, globalSymbol)
        }
        for _key, _val := range( moduleInfo.FP.Get_importedAliasMap(_env).Items ) {
            key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            val := _val.(Ast_AliasTypeInfoDownCast).ToAst_AliasTypeInfo()
            self.importedAliasMap.Set(key,val)
        }
    case *FrontInterface_MetaOrModule__Meta:
    metaInfo := _exp6643.Val1
        moduleInfo = self.FP.processImportFromFile(_env, processInfo, moduleMeta.FP.Get_lnsPath(_env), metaInfo, orgModulePath, modulePath, nameList, depth)
        
        moduleMeta.FP.Set_metaOrModule(_env, &FrontInterface_MetaOrModule__Module{moduleInfo})
    }
    if depth == 1{
        self.importModule2ModuleInfo.Set(moduleInfo.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env),moduleInfo)
    }
    self.importModuleName2ModuleInfo.Set(modulePath,moduleInfo)
    self.importModuleInfo.FP.Remove(_env)
    Log_log(_env, Log_Level__Info, __func__, 1245, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.LuaVM.String_format("%s complete", []LnsAny{orgModulePath})
    }))
    
    return moduleInfo
}

// 1251: decl @lune.@base.@Import.Import.processImport
func (self *Import_Import) ProcessImport(_env *LnsEnv, processInfo *Ast_ProcessInfo,modulePath string) *FrontInterface_ModuleInfo {
    return self.FP.processImportMain(_env, processInfo, modulePath, 1)
}


// declaration Class -- ImportParam
type Import_ImportParamMtd interface {
    GetTypeInfo(_env *LnsEnv, arg1 LnsInt)(LnsAny, LnsAny)
    GetTypeInfoFrom(_env *LnsEnv, arg1 *Import__IdInfo)(LnsAny, LnsAny)
}
type Import_ImportParam struct {
    Pos *Types_Position
    Modifier *TransUnitIF_Modifier
    ProcessInfo *Ast_ProcessInfo
    TypeId2Scope *LnsMap
    TypeId2TypeInfo *LnsMap
    ImportedAliasMap *LnsMap
    LazyModuleSet *LnsSet
    MetaInfo *Lns_luaValue
    Scope *Ast_Scope
    ModuleTypeInfo *Ast_TypeInfo
    ScopeAccess LnsInt
    typeId2AtomMap *LnsMap
    dependLibId2DependInfo *LnsMap
    FP Import_ImportParamMtd
}
func Import_ImportParam2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import_ImportParam).FP
}
type Import_ImportParamDownCast interface {
    ToImport_ImportParam() *Import_ImportParam
}
func Import_ImportParamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import_ImportParamDownCast)
    if ok { return work.ToImport_ImportParam() }
    return nil
}
func (obj *Import_ImportParam) ToImport_ImportParam() *Import_ImportParam {
    return obj
}
func NewImport_ImportParam(_env *LnsEnv, arg1 *Types_Position, arg2 *TransUnitIF_Modifier, arg3 *Ast_ProcessInfo, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsSet, arg8 *Lns_luaValue, arg9 *Ast_Scope, arg10 *Ast_TypeInfo, arg11 LnsInt, arg12 *LnsMap, arg13 *LnsMap) *Import_ImportParam {
    obj := &Import_ImportParam{}
    obj.FP = obj
    obj.InitImport_ImportParam(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
    return obj
}
func (self *Import_ImportParam) InitImport_ImportParam(_env *LnsEnv, arg1 *Types_Position, arg2 *TransUnitIF_Modifier, arg3 *Ast_ProcessInfo, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsSet, arg8 *Lns_luaValue, arg9 *Ast_Scope, arg10 *Ast_TypeInfo, arg11 LnsInt, arg12 *LnsMap, arg13 *LnsMap) {
    self.Pos = arg1
    self.Modifier = arg2
    self.ProcessInfo = arg3
    self.TypeId2Scope = arg4
    self.TypeId2TypeInfo = arg5
    self.ImportedAliasMap = arg6
    self.LazyModuleSet = arg7
    self.MetaInfo = arg8
    self.Scope = arg9
    self.ModuleTypeInfo = arg10
    self.ScopeAccess = arg11
    self.typeId2AtomMap = arg12
    self.dependLibId2DependInfo = arg13
}
// 160: decl @lune.@base.@Import.ImportParam.getTypeInfo
func (self *Import_ImportParam) GetTypeInfo(_env *LnsEnv, typeId LnsInt)(LnsAny, LnsAny) {
    {
        _typeInfo := self.TypeId2TypeInfo.Get(typeId)
        if !Lns_IsNil( _typeInfo ) {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo, nil
        }
    }
    {
        _atom := self.typeId2AtomMap.Get(typeId)
        if !Lns_IsNil( _atom ) {
            atom := _atom.(*Import__TypeInfo)
            var typeInfo LnsAny
            var mess LnsAny
            typeInfo,mess = atom.FP.CreateTypeInfoCache(_env, self)
            if typeInfo != nil{
                typeInfo_98 := typeInfo.(*Ast_TypeInfo)
                self.TypeId2TypeInfo.Set(typeId,typeInfo_98)
            }
            return typeInfo, mess
        }
    }
    return nil, nil
}

// 179: decl @lune.@base.@Import.ImportParam.getTypeInfoFrom
func (self *Import_ImportParam) GetTypeInfoFrom(_env *LnsEnv, typeId *Import__IdInfo)(LnsAny, LnsAny) {
    if typeId.Mod == 0{
        return self.FP.GetTypeInfo(_env, typeId.Id)
    }
    if typeId.Mod == FrontInterface_getRootDependModId(_env){
        return Ast_getRootProcessInfoRo(_env).FP.GetTypeInfo(_env, typeId.Id), nil
    }
    var moduleInfo *FrontInterface_ModuleInfo
    
    {
        _moduleInfo := self.dependLibId2DependInfo.Get(typeId.Mod)
        if _moduleInfo == nil{
            Util_err(_env, _env.LuaVM.String_format("%s, %d, %d", []LnsAny{self.ModuleTypeInfo.FP.GetTxt(_env, nil, nil, nil), typeId.Mod, typeId.Id}))
        } else {
            moduleInfo = _moduleInfo.(*FrontInterface_ModuleInfo)
        }
    }
    {
        _typeInfo := moduleInfo.FP.Get_importId2localTypeInfoMap(_env).Get(typeId.Id)
        if !Lns_IsNil( _typeInfo ) {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo, nil
        }
    }
    {
        _typeInfo := moduleInfo.FP.Get_exportInfo(_env).FP.Get_processInfo(_env).FP.GetTypeInfo(_env, typeId.Id)
        if !Lns_IsNil( _typeInfo ) {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo, nil
        }
    }
    return nil, _env.LuaVM.String_format("not found type -- %s, %d, %d", []LnsAny{self.ModuleTypeInfo.FP.GetTxt(_env, nil, nil, nil), typeId.Mod, typeId.Id})
}


// declaration Class -- _TypeInfo
type Import__TypeInfoMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfo struct {
    Skind LnsInt
    TypeId LnsInt
    FP Import__TypeInfoMtd
}
func Import__TypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfo).FP
}
type Import__TypeInfoDownCast interface {
    ToImport__TypeInfo() *Import__TypeInfo
}
func Import__TypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoDownCast)
    if ok { return work.ToImport__TypeInfo() }
    return nil
}
func (obj *Import__TypeInfo) ToImport__TypeInfo() *Import__TypeInfo {
    return obj
}
func (self *Import__TypeInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["skind"] = Lns_ToCollection( self.Skind )
    obj.Items["typeId"] = Lns_ToCollection( self.TypeId )
    return obj
}
func (self *Import__TypeInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfo_FromMapMain( newObj *Import__TypeInfo, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["skind"], false, nil); !ok {
       return false,nil,"skind:" + mess.(string)
    } else {
       newObj.Skind = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["typeId"], false, nil); !ok {
       return false,nil,"typeId:" + mess.(string)
    } else {
       newObj.TypeId = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 132: DeclConstr
func (self *Import__TypeInfo) InitImport__TypeInfo(_env *LnsEnv) {
    self.TypeId = Ast_rootTypeId
    
    self.Skind = Ast_SerializeKind__Normal
    
}


// 147: decl @lune.@base.@Import._TypeInfo.createTypeInfoCache
func (self *Import__TypeInfo) CreateTypeInfoCache(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    {
        _typeInfo := param.TypeId2TypeInfo.Get(self.TypeId)
        if !Lns_IsNil( _typeInfo ) {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo, nil
        }
    }
    var typeInfo LnsAny
    var mess LnsAny
    typeInfo,mess = self.FP.CreateTypeInfo(_env, param)
    if typeInfo != nil{
        typeInfo_83 := typeInfo.(*Ast_TypeInfo)
        param.TypeId2TypeInfo.Set(self.TypeId,typeInfo_83)
        typeInfo_83.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    }
    return typeInfo, mess
}


// declaration Class -- _IdInfo
type Import__IdInfoMtd interface {
    ToMap() *LnsMap
}
type Import__IdInfo struct {
    Id LnsInt
    Mod LnsInt
    FP Import__IdInfoMtd
}
func Import__IdInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__IdInfo).FP
}
type Import__IdInfoDownCast interface {
    ToImport__IdInfo() *Import__IdInfo
}
func Import__IdInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__IdInfoDownCast)
    if ok { return work.ToImport__IdInfo() }
    return nil
}
func (obj *Import__IdInfo) ToImport__IdInfo() *Import__IdInfo {
    return obj
}
func NewImport__IdInfo(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *Import__IdInfo {
    obj := &Import__IdInfo{}
    obj.FP = obj
    obj.InitImport__IdInfo(_env, arg1, arg2)
    return obj
}
func (self *Import__IdInfo) InitImport__IdInfo(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) {
    self.Id = arg1
    self.Mod = arg2
}
func (self *Import__IdInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["id"] = Lns_ToCollection( self.Id )
    obj.Items["mod"] = Lns_ToCollection( self.Mod )
    return obj
}
func (self *Import__IdInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__IdInfo__fromMap_1165_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__IdInfo_FromMap( arg1, paramList )
}
func Import__IdInfo__fromStem_1169_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__IdInfo_FromMap( arg1, paramList )
}
func Import__IdInfo_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__IdInfo_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__IdInfo_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__IdInfo{}
    newObj.FP = newObj
    return Import__IdInfo_FromMapMain( newObj, objMap, paramList )
}
func Import__IdInfo_FromMapMain( newObj *Import__IdInfo, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["id"], false, nil); !ok {
       return false,nil,"id:" + mess.(string)
    } else {
       newObj.Id = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["mod"], false, nil); !ok {
       return false,nil,"mod:" + mess.(string)
    } else {
       newObj.Mod = conv.(LnsInt)
    }
    return true, newObj, nil
}

// declaration Class -- _TypeInfoNilable
type Import__TypeInfoNilableMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoNilable struct {
    Import__TypeInfo
    OrgTypeId *Import__IdInfo
    FP Import__TypeInfoNilableMtd
}
func Import__TypeInfoNilable2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoNilable).FP
}
type Import__TypeInfoNilableDownCast interface {
    ToImport__TypeInfoNilable() *Import__TypeInfoNilable
}
func Import__TypeInfoNilableDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoNilableDownCast)
    if ok { return work.ToImport__TypeInfoNilable() }
    return nil
}
func (obj *Import__TypeInfoNilable) ToImport__TypeInfoNilable() *Import__TypeInfoNilable {
    return obj
}
func NewImport__TypeInfoNilable(_env *LnsEnv, arg1 *Import__IdInfo) *Import__TypeInfoNilable {
    obj := &Import__TypeInfoNilable{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoNilable(_env, arg1)
    return obj
}
func (self *Import__TypeInfoNilable) InitImport__TypeInfoNilable(_env *LnsEnv, arg1 *Import__IdInfo) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.OrgTypeId = arg1
}
func (self *Import__TypeInfoNilable) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["orgTypeId"] = Lns_ToCollection( self.OrgTypeId )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoNilable) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoNilable__fromMap_1220_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoNilable_FromMap( arg1, paramList )
}
func Import__TypeInfoNilable__fromStem_1224_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoNilable_FromMap( arg1, paramList )
}
func Import__TypeInfoNilable_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoNilable_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoNilable_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoNilable{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoNilable_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoNilable_FromMapMain( newObj *Import__TypeInfoNilable, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Import__IdInfo_FromMapSub( objMap.Items["orgTypeId"], false, nil); !ok {
       return false,nil,"orgTypeId:" + mess.(string)
    } else {
       newObj.OrgTypeId = conv.(*Import__IdInfo)
    }
    return true, newObj, nil
}
// 203: decl @lune.@base.@Import._TypeInfoNilable.createTypeInfo
func (self *Import__TypeInfoNilable) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    var orgTypeInfo *Ast_TypeInfo
    
    {
        _orgTypeInfo := Import_convExp701(Lns_2DDD(param.FP.GetTypeInfoFrom(_env, self.OrgTypeId)))
        if _orgTypeInfo == nil{
            Util_err(_env, _env.LuaVM.String_format("failed to createTypeInfo -- self.orgTypeId = (%d,%d)", []LnsAny{self.OrgTypeId.Mod, self.OrgTypeId.Id}))
        } else {
            orgTypeInfo = _orgTypeInfo.(*Ast_TypeInfo)
        }
    }
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = orgTypeInfo.FP.Get_nilableTypeInfo(_env)
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    newTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoAlias
type Import__TypeInfoAliasMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoAlias struct {
    Import__TypeInfo
    ParentId LnsInt
    rawTxt string
    srcTypeId *Import__IdInfo
    FP Import__TypeInfoAliasMtd
}
func Import__TypeInfoAlias2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoAlias).FP
}
type Import__TypeInfoAliasDownCast interface {
    ToImport__TypeInfoAlias() *Import__TypeInfoAlias
}
func Import__TypeInfoAliasDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoAliasDownCast)
    if ok { return work.ToImport__TypeInfoAlias() }
    return nil
}
func (obj *Import__TypeInfoAlias) ToImport__TypeInfoAlias() *Import__TypeInfoAlias {
    return obj
}
func NewImport__TypeInfoAlias(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 *Import__IdInfo) *Import__TypeInfoAlias {
    obj := &Import__TypeInfoAlias{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoAlias(_env, arg1, arg2, arg3)
    return obj
}
func (self *Import__TypeInfoAlias) InitImport__TypeInfoAlias(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 *Import__IdInfo) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.ParentId = arg1
    self.rawTxt = arg2
    self.srcTypeId = arg3
}
func (self *Import__TypeInfoAlias) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["rawTxt"] = Lns_ToCollection( self.rawTxt )
    obj.Items["srcTypeId"] = Lns_ToCollection( self.srcTypeId )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoAlias) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoAlias__fromMap_1258_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlias_FromMap( arg1, paramList )
}
func Import__TypeInfoAlias__fromStem_1262_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlias_FromMap( arg1, paramList )
}
func Import__TypeInfoAlias_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoAlias_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoAlias_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoAlias{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoAlias_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoAlias_FromMapMain( newObj *Import__TypeInfoAlias, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["rawTxt"], false, nil); !ok {
       return false,nil,"rawTxt:" + mess.(string)
    } else {
       newObj.rawTxt = conv.(string)
    }
    if ok,conv,mess := Import__IdInfo_FromMapSub( objMap.Items["srcTypeId"], false, nil); !ok {
       return false,nil,"srcTypeId:" + mess.(string)
    } else {
       newObj.srcTypeId = conv.(*Import__IdInfo)
    }
    return true, newObj, nil
}
// 221: decl @lune.@base.@Import._TypeInfoAlias.createTypeInfo
func (self *Import__TypeInfoAlias) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@Import._TypeInfoAlias.createTypeInfo"
    var srcTypeInfo *Ast_TypeInfo
    srcTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, self.srcTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_AliasTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateAlias(_env, param.ProcessInfo, self.rawTxt, true, Ast_AccessMode__Pub, param.ModuleTypeInfo, srcTypeInfo)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    newTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    
    {
        __ := Import_convExp850(Lns_2DDD(param.FP.GetTypeInfo(_env, self.ParentId)))
        if __ == nil{
            return nil, _env.LuaVM.String_format("%s: not found parentInfo %d %s", []LnsAny{__func__, self.ParentId, self.rawTxt})
        } else {
            _ = __.(*Ast_TypeInfo)
        }
    }
    var parentScope *Ast_Scope
    
    {
        _parentScope := param.TypeId2Scope.Get(self.ParentId)
        if _parentScope == nil{
            return nil, _env.LuaVM.String_format("%s: not found parentScope %s %s", []LnsAny{__func__, self.ParentId, self.rawTxt})
        } else {
            parentScope = _parentScope.(*Ast_Scope)
        }
    }
    parentScope.FP.AddAliasForType(_env, param.ProcessInfo, self.rawTxt, nil, &newTypeInfo.Ast_TypeInfo)
    param.ImportedAliasMap.Set(srcTypeInfo,newTypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoDDD
type Import__TypeInfoDDDMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoDDD struct {
    Import__TypeInfo
    ParentId LnsInt
    ItemTypeId *Import__IdInfo
    ExtTypeFlag bool
    FP Import__TypeInfoDDDMtd
}
func Import__TypeInfoDDD2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoDDD).FP
}
type Import__TypeInfoDDDDownCast interface {
    ToImport__TypeInfoDDD() *Import__TypeInfoDDD
}
func Import__TypeInfoDDDDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoDDDDownCast)
    if ok { return work.ToImport__TypeInfoDDD() }
    return nil
}
func (obj *Import__TypeInfoDDD) ToImport__TypeInfoDDD() *Import__TypeInfoDDD {
    return obj
}
func NewImport__TypeInfoDDD(_env *LnsEnv, arg1 LnsInt, arg2 *Import__IdInfo, arg3 bool) *Import__TypeInfoDDD {
    obj := &Import__TypeInfoDDD{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoDDD(_env, arg1, arg2, arg3)
    return obj
}
func (self *Import__TypeInfoDDD) InitImport__TypeInfoDDD(_env *LnsEnv, arg1 LnsInt, arg2 *Import__IdInfo, arg3 bool) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.ParentId = arg1
    self.ItemTypeId = arg2
    self.ExtTypeFlag = arg3
}
func (self *Import__TypeInfoDDD) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["itemTypeId"] = Lns_ToCollection( self.ItemTypeId )
    obj.Items["extTypeFlag"] = Lns_ToCollection( self.ExtTypeFlag )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoDDD) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoDDD__fromMap_1287_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoDDD_FromMap( arg1, paramList )
}
func Import__TypeInfoDDD__fromStem_1291_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoDDD_FromMap( arg1, paramList )
}
func Import__TypeInfoDDD_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoDDD_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoDDD_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoDDD{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoDDD_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoDDD_FromMapMain( newObj *Import__TypeInfoDDD, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Import__IdInfo_FromMapSub( objMap.Items["itemTypeId"], false, nil); !ok {
       return false,nil,"itemTypeId:" + mess.(string)
    } else {
       newObj.ItemTypeId = conv.(*Import__IdInfo)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["extTypeFlag"], false, nil); !ok {
       return false,nil,"extTypeFlag:" + mess.(string)
    } else {
       newObj.ExtTypeFlag = conv.(bool)
    }
    return true, newObj, nil
}
// 253: decl @lune.@base.@Import._TypeInfoDDD.createTypeInfo
func (self *Import__TypeInfoDDD) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    var itemTypeInfo *Ast_TypeInfo
    itemTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, self.ItemTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_DDDTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateDDD(_env, itemTypeInfo, true, self.ExtTypeFlag)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    newTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoAlternate
type Import__TypeInfoAlternateMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoAlternate struct {
    Import__TypeInfo
    ParentId LnsInt
    Txt string
    AccessMode LnsInt
    BaseId *Import__IdInfo
    IfList *LnsList
    BelongClassFlag bool
    AltIndex LnsInt
    FP Import__TypeInfoAlternateMtd
}
func Import__TypeInfoAlternate2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoAlternate).FP
}
type Import__TypeInfoAlternateDownCast interface {
    ToImport__TypeInfoAlternate() *Import__TypeInfoAlternate
}
func Import__TypeInfoAlternateDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoAlternateDownCast)
    if ok { return work.ToImport__TypeInfoAlternate() }
    return nil
}
func (obj *Import__TypeInfoAlternate) ToImport__TypeInfoAlternate() *Import__TypeInfoAlternate {
    return obj
}
func NewImport__TypeInfoAlternate(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *Import__IdInfo, arg5 *LnsList, arg6 bool, arg7 LnsInt) *Import__TypeInfoAlternate {
    obj := &Import__TypeInfoAlternate{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoAlternate(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Import__TypeInfoAlternate) InitImport__TypeInfoAlternate(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *Import__IdInfo, arg5 *LnsList, arg6 bool, arg7 LnsInt) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.ParentId = arg1
    self.Txt = arg2
    self.AccessMode = arg3
    self.BaseId = arg4
    self.IfList = arg5
    self.BelongClassFlag = arg6
    self.AltIndex = arg7
}
func (self *Import__TypeInfoAlternate) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["baseId"] = Lns_ToCollection( self.BaseId )
    obj.Items["ifList"] = Lns_ToCollection( self.IfList )
    obj.Items["belongClassFlag"] = Lns_ToCollection( self.BelongClassFlag )
    obj.Items["altIndex"] = Lns_ToCollection( self.AltIndex )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoAlternate) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoAlternate__fromMap_1332_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlternate_FromMap( arg1, paramList )
}
func Import__TypeInfoAlternate__fromStem_1336_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlternate_FromMap( arg1, paramList )
}
func Import__TypeInfoAlternate_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoAlternate_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoAlternate_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoAlternate{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoAlternate_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoAlternate_FromMapMain( newObj *Import__TypeInfoAlternate, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Import__IdInfo_FromMapSub( objMap.Items["baseId"], false, nil); !ok {
       return false,nil,"baseId:" + mess.(string)
    } else {
       newObj.BaseId = conv.(*Import__IdInfo)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["ifList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Import__IdInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"ifList:" + mess.(string)
    } else {
       newObj.IfList = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["belongClassFlag"], false, nil); !ok {
       return false,nil,"belongClassFlag:" + mess.(string)
    } else {
       newObj.BelongClassFlag = conv.(bool)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["altIndex"], false, nil); !ok {
       return false,nil,"altIndex:" + mess.(string)
    } else {
       newObj.AltIndex = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 275: decl @lune.@base.@Import._TypeInfoAlternate.createTypeInfo
func (self *Import__TypeInfoAlternate) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    var baseInfo *Ast_TypeInfo
    baseInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, self.BaseId))).(*Ast_TypeInfo)
    var interfaceList *LnsList
    interfaceList = NewLnsList([]LnsAny{})
    for _, _ifTypeId := range( self.IfList.Items ) {
        ifTypeId := _ifTypeId.(Import__IdInfoDownCast).ToImport__IdInfo()
        interfaceList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, ifTypeId))).(*Ast_TypeInfo)))
    }
    var newTypeInfo *Ast_AlternateTypeInfo
    newTypeInfo = Import_convExp1117(Lns_2DDD(param.ProcessInfo.FP.CreateAlternate(_env, self.BelongClassFlag, self.AltIndex, self.Txt, self.AccessMode, param.ModuleTypeInfo, baseInfo, interfaceList)))
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    newTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoGeneric
type Import__TypeInfoGenericMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoGeneric struct {
    Import__TypeInfo
    GenSrcTypeId *Import__IdInfo
    GenTypeList *LnsList
    FP Import__TypeInfoGenericMtd
}
func Import__TypeInfoGeneric2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoGeneric).FP
}
type Import__TypeInfoGenericDownCast interface {
    ToImport__TypeInfoGeneric() *Import__TypeInfoGeneric
}
func Import__TypeInfoGenericDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoGenericDownCast)
    if ok { return work.ToImport__TypeInfoGeneric() }
    return nil
}
func (obj *Import__TypeInfoGeneric) ToImport__TypeInfoGeneric() *Import__TypeInfoGeneric {
    return obj
}
func NewImport__TypeInfoGeneric(_env *LnsEnv, arg1 *Import__IdInfo, arg2 *LnsList) *Import__TypeInfoGeneric {
    obj := &Import__TypeInfoGeneric{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoGeneric(_env, arg1, arg2)
    return obj
}
func (self *Import__TypeInfoGeneric) InitImport__TypeInfoGeneric(_env *LnsEnv, arg1 *Import__IdInfo, arg2 *LnsList) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.GenSrcTypeId = arg1
    self.GenTypeList = arg2
}
func (self *Import__TypeInfoGeneric) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["genSrcTypeId"] = Lns_ToCollection( self.GenSrcTypeId )
    obj.Items["genTypeList"] = Lns_ToCollection( self.GenTypeList )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoGeneric) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoGeneric__fromMap_1379_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoGeneric_FromMap( arg1, paramList )
}
func Import__TypeInfoGeneric__fromStem_1383_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoGeneric_FromMap( arg1, paramList )
}
func Import__TypeInfoGeneric_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoGeneric_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoGeneric_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoGeneric{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoGeneric_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoGeneric_FromMapMain( newObj *Import__TypeInfoGeneric, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Import__IdInfo_FromMapSub( objMap.Items["genSrcTypeId"], false, nil); !ok {
       return false,nil,"genSrcTypeId:" + mess.(string)
    } else {
       newObj.GenSrcTypeId = conv.(*Import__IdInfo)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["genTypeList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Import__IdInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"genTypeList:" + mess.(string)
    } else {
       newObj.GenTypeList = conv.(*LnsList)
    }
    return true, newObj, nil
}
// 296: decl @lune.@base.@Import._TypeInfoGeneric.createTypeInfo
func (self *Import__TypeInfoGeneric) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    var genSrcTypeInfo *Ast_TypeInfo
    genSrcTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, self.GenSrcTypeId))).(*Ast_TypeInfo)
    var genTypeList *LnsList
    genTypeList = NewLnsList([]LnsAny{})
    for _, _typeId := range( self.GenTypeList.Items ) {
        typeId := _typeId.(Import__IdInfoDownCast).ToImport__IdInfo()
        genTypeList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, typeId))).(*Ast_TypeInfo)))
    }
    var newTypeInfo *Ast_GenericTypeInfo
    var scope *Ast_Scope
    newTypeInfo,scope = param.ProcessInfo.FP.CreateGeneric(_env, genSrcTypeInfo, genTypeList, param.ModuleTypeInfo)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    newTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    param.TypeId2Scope.Set(self.TypeId,scope)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoBox
type Import__TypeInfoBoxMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoBox struct {
    Import__TypeInfo
    AccessMode LnsInt
    BoxingType LnsInt
    FP Import__TypeInfoBoxMtd
}
func Import__TypeInfoBox2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoBox).FP
}
type Import__TypeInfoBoxDownCast interface {
    ToImport__TypeInfoBox() *Import__TypeInfoBox
}
func Import__TypeInfoBoxDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoBoxDownCast)
    if ok { return work.ToImport__TypeInfoBox() }
    return nil
}
func (obj *Import__TypeInfoBox) ToImport__TypeInfoBox() *Import__TypeInfoBox {
    return obj
}
func NewImport__TypeInfoBox(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *Import__TypeInfoBox {
    obj := &Import__TypeInfoBox{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoBox(_env, arg1, arg2)
    return obj
}
func (self *Import__TypeInfoBox) InitImport__TypeInfoBox(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.AccessMode = arg1
    self.BoxingType = arg2
}
func (self *Import__TypeInfoBox) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["boxingType"] = Lns_ToCollection( self.BoxingType )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoBox) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoBox__fromMap_1408_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoBox_FromMap( arg1, paramList )
}
func Import__TypeInfoBox__fromStem_1412_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoBox_FromMap( arg1, paramList )
}
func Import__TypeInfoBox_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoBox_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoBox_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoBox{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoBox_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoBox_FromMapMain( newObj *Import__TypeInfoBox, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["boxingType"], false, nil); !ok {
       return false,nil,"boxingType:" + mess.(string)
    } else {
       newObj.BoxingType = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 316: decl @lune.@base.@Import._TypeInfoBox.createTypeInfo
func (self *Import__TypeInfoBox) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    var boxingType *Ast_TypeInfo
    boxingType = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(_env, self.BoxingType))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateBox(_env, self.AccessMode, boxingType)
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    newTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoExt
type Import__TypeInfoExtMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoExt struct {
    Import__TypeInfo
    ExtedTypeId *Import__IdInfo
    FP Import__TypeInfoExtMtd
}
func Import__TypeInfoExt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoExt).FP
}
type Import__TypeInfoExtDownCast interface {
    ToImport__TypeInfoExt() *Import__TypeInfoExt
}
func Import__TypeInfoExtDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoExtDownCast)
    if ok { return work.ToImport__TypeInfoExt() }
    return nil
}
func (obj *Import__TypeInfoExt) ToImport__TypeInfoExt() *Import__TypeInfoExt {
    return obj
}
func NewImport__TypeInfoExt(_env *LnsEnv, arg1 *Import__IdInfo) *Import__TypeInfoExt {
    obj := &Import__TypeInfoExt{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoExt(_env, arg1)
    return obj
}
func (self *Import__TypeInfoExt) InitImport__TypeInfoExt(_env *LnsEnv, arg1 *Import__IdInfo) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.ExtedTypeId = arg1
}
func (self *Import__TypeInfoExt) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["extedTypeId"] = Lns_ToCollection( self.ExtedTypeId )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoExt) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoExt__fromMap_1438_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoExt_FromMap( arg1, paramList )
}
func Import__TypeInfoExt__fromStem_1442_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoExt_FromMap( arg1, paramList )
}
func Import__TypeInfoExt_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoExt_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoExt_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoExt{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoExt_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoExt_FromMapMain( newObj *Import__TypeInfoExt, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Import__IdInfo_FromMapSub( objMap.Items["extedTypeId"], false, nil); !ok {
       return false,nil,"extedTypeId:" + mess.(string)
    } else {
       newObj.ExtedTypeId = conv.(*Import__IdInfo)
    }
    return true, newObj, nil
}
// 329: decl @lune.@base.@Import._TypeInfoExt.createTypeInfo
func (self *Import__TypeInfoExt) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    var extedType *Ast_TypeInfo
    extedType = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, self.ExtedTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_TypeInfo
    switch _exp1418 := param.ProcessInfo.FP.CreateLuaval(_env, extedType, true).(type) {
    case *Ast_LuavalResult__OK:
    extType := _exp1418.Val1
        newTypeInfo = extType
        
    case *Ast_LuavalResult__Err:
    mess := _exp1418.Val1
        Util_err(_env, mess)
    }
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    newTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoModifier
type Import__TypeInfoModifierMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoModifier struct {
    Import__TypeInfo
    SrcTypeId *Import__IdInfo
    MutMode LnsInt
    FP Import__TypeInfoModifierMtd
}
func Import__TypeInfoModifier2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoModifier).FP
}
type Import__TypeInfoModifierDownCast interface {
    ToImport__TypeInfoModifier() *Import__TypeInfoModifier
}
func Import__TypeInfoModifierDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoModifierDownCast)
    if ok { return work.ToImport__TypeInfoModifier() }
    return nil
}
func (obj *Import__TypeInfoModifier) ToImport__TypeInfoModifier() *Import__TypeInfoModifier {
    return obj
}
func NewImport__TypeInfoModifier(_env *LnsEnv, arg1 *Import__IdInfo, arg2 LnsInt) *Import__TypeInfoModifier {
    obj := &Import__TypeInfoModifier{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoModifier(_env, arg1, arg2)
    return obj
}
func (self *Import__TypeInfoModifier) InitImport__TypeInfoModifier(_env *LnsEnv, arg1 *Import__IdInfo, arg2 LnsInt) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.SrcTypeId = arg1
    self.MutMode = arg2
}
func (self *Import__TypeInfoModifier) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["srcTypeId"] = Lns_ToCollection( self.SrcTypeId )
    obj.Items["mutMode"] = Lns_ToCollection( self.MutMode )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoModifier) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoModifier__fromMap_1469_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoModifier_FromMap( arg1, paramList )
}
func Import__TypeInfoModifier__fromStem_1473_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoModifier_FromMap( arg1, paramList )
}
func Import__TypeInfoModifier_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoModifier_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoModifier_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoModifier{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoModifier_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoModifier_FromMapMain( newObj *Import__TypeInfoModifier, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Import__IdInfo_FromMapSub( objMap.Items["srcTypeId"], false, nil); !ok {
       return false,nil,"srcTypeId:" + mess.(string)
    } else {
       newObj.SrcTypeId = conv.(*Import__IdInfo)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["mutMode"], false, nil); !ok {
       return false,nil,"mutMode:" + mess.(string)
    } else {
       newObj.MutMode = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 352: decl @lune.@base.@Import._TypeInfoModifier.createTypeInfo
func (self *Import__TypeInfoModifier) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    var srcTypeInfo *Ast_TypeInfo
    
    {
        _srcTypeInfo := Import_convExp1495(Lns_2DDD(param.FP.GetTypeInfoFrom(_env, self.SrcTypeId)))
        if _srcTypeInfo == nil{
            return nil, _env.LuaVM.String_format("not found srcType -- %d", []LnsAny{self.SrcTypeId.Id})
        } else {
            srcTypeInfo = _srcTypeInfo.(*Ast_TypeInfo)
        }
    }
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = param.Modifier.FP.CreateModifier(_env, srcTypeInfo, self.MutMode)
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    newTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoModule
type Import__TypeInfoModuleMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoModule struct {
    Import__TypeInfo
    ParentId LnsInt
    Txt string
    FP Import__TypeInfoModuleMtd
}
func Import__TypeInfoModule2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoModule).FP
}
type Import__TypeInfoModuleDownCast interface {
    ToImport__TypeInfoModule() *Import__TypeInfoModule
}
func Import__TypeInfoModuleDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoModuleDownCast)
    if ok { return work.ToImport__TypeInfoModule() }
    return nil
}
func (obj *Import__TypeInfoModule) ToImport__TypeInfoModule() *Import__TypeInfoModule {
    return obj
}
func NewImport__TypeInfoModule(_env *LnsEnv, arg1 LnsInt, arg2 string) *Import__TypeInfoModule {
    obj := &Import__TypeInfoModule{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoModule(_env, arg1, arg2)
    return obj
}
func (self *Import__TypeInfoModule) InitImport__TypeInfoModule(_env *LnsEnv, arg1 LnsInt, arg2 string) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.ParentId = arg1
    self.Txt = arg2
}
func (self *Import__TypeInfoModule) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoModule) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoModule__fromMap_1523_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoModule_FromMap( arg1, paramList )
}
func Import__TypeInfoModule__fromStem_1527_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoModule_FromMap( arg1, paramList )
}
func Import__TypeInfoModule_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoModule_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoModule_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoModule{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoModule_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoModule_FromMapMain( newObj *Import__TypeInfoModule, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    return true, newObj, nil
}
// 369: decl @lune.@base.@Import._TypeInfoModule.createTypeInfo
func (self *Import__TypeInfoModule) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@Import._TypeInfoModule.createTypeInfo"
    var parentInfo *Ast_TypeInfo
    parentInfo = Ast_headTypeInfo
    if self.ParentId != Ast_rootTypeId{
        var workTypeInfo *Ast_TypeInfo
        
        {
            _workTypeInfo := Import_convExp1603(Lns_2DDD(param.FP.GetTypeInfo(_env, self.ParentId)))
            if _workTypeInfo == nil{
                Util_err(_env, _env.LuaVM.String_format("not found parentInfo %d %s", []LnsAny{self.ParentId, self.Txt}))
            } else {
                workTypeInfo = _workTypeInfo.(*Ast_TypeInfo)
            }
        }
        parentInfo = workTypeInfo
        
    }
    var parentScope *Ast_Scope
    
    {
        _parentScope := param.TypeId2Scope.Get(self.ParentId)
        if _parentScope == nil{
            return nil, _env.LuaVM.String_format("%s: not found parentScope %s %s", []LnsAny{__func__, self.ParentId, self.Txt})
        } else {
            parentScope = _parentScope.(*Ast_Scope)
        }
    }
    var newTypeInfo LnsAny
    newTypeInfo = parentScope.FP.GetTypeInfoChild(_env, self.Txt)
    {
        __ := newTypeInfo
        if !Lns_IsNil( __ ) {
            panic("internal error")
        } else {
            var scope *Ast_Scope
            scope = NewAst_Scope(_env, param.ProcessInfo, parentScope, true, nil, nil)
            var mutable bool
            mutable = false
            if self.TypeId == param.MetaInfo.GetAt( "__moduleTypeId" ).(LnsInt){
                mutable = param.MetaInfo.GetAt( "__moduleMutable" ).(bool)
                
            }
            var workTypeInfo *Ast_TypeInfo
            workTypeInfo = param.ProcessInfo.FP.CreateModule(_env, scope, parentInfo, true, self.Txt, mutable)
            newTypeInfo = workTypeInfo
            
            param.TypeId2Scope.Set(self.TypeId,scope)
            param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
            workTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
            parentScope.FP.AddClass(_env, param.ProcessInfo, self.Txt, nil, workTypeInfo)
            Log_log(_env, Log_Level__Info, __func__, 415, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.LuaVM.String_format("new module -- %s, %s, %d, %d, %d", []LnsAny{self.Txt, workTypeInfo.FP.GetFullName(_env, Ast_defaultTypeNameCtrl, parentScope.FP, false), self.TypeId, workTypeInfo.FP.Get_typeId(_env).Id, parentScope.FP.Get_scopeId(_env)})
            }))
            
        }
    }
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoNormal
type Import__TypeInfoNormalMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoNormal struct {
    Import__TypeInfo
    ParentId LnsInt
    AbstractFlag bool
    BaseId *Import__IdInfo
    Txt string
    StaticFlag bool
    AccessMode LnsInt
    Kind LnsInt
    MutMode LnsInt
    AsyncMode LnsInt
    IfList *LnsList
    ItemTypeId *LnsList
    ArgTypeId *LnsList
    RetTypeId *LnsList
    Children *LnsList
    ModuleLang LnsAny
    RequirePath LnsAny
    FP Import__TypeInfoNormalMtd
}
func Import__TypeInfoNormal2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoNormal).FP
}
type Import__TypeInfoNormalDownCast interface {
    ToImport__TypeInfoNormal() *Import__TypeInfoNormal
}
func Import__TypeInfoNormalDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoNormalDownCast)
    if ok { return work.ToImport__TypeInfoNormal() }
    return nil
}
func (obj *Import__TypeInfoNormal) ToImport__TypeInfoNormal() *Import__TypeInfoNormal {
    return obj
}
func NewImport__TypeInfoNormal(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Import__IdInfo, arg4 string, arg5 bool, arg6 LnsInt, arg7 LnsInt, arg8 LnsInt, arg9 LnsInt, arg10 *LnsList, arg11 *LnsList, arg12 *LnsList, arg13 *LnsList, arg14 *LnsList, arg15 LnsAny, arg16 LnsAny) *Import__TypeInfoNormal {
    obj := &Import__TypeInfoNormal{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoNormal(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16)
    return obj
}
func (self *Import__TypeInfoNormal) InitImport__TypeInfoNormal(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Import__IdInfo, arg4 string, arg5 bool, arg6 LnsInt, arg7 LnsInt, arg8 LnsInt, arg9 LnsInt, arg10 *LnsList, arg11 *LnsList, arg12 *LnsList, arg13 *LnsList, arg14 *LnsList, arg15 LnsAny, arg16 LnsAny) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.ParentId = arg1
    self.AbstractFlag = arg2
    self.BaseId = arg3
    self.Txt = arg4
    self.StaticFlag = arg5
    self.AccessMode = arg6
    self.Kind = arg7
    self.MutMode = arg8
    self.AsyncMode = arg9
    self.IfList = arg10
    self.ItemTypeId = arg11
    self.ArgTypeId = arg12
    self.RetTypeId = arg13
    self.Children = arg14
    self.ModuleLang = arg15
    self.RequirePath = arg16
}
func (self *Import__TypeInfoNormal) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["abstractFlag"] = Lns_ToCollection( self.AbstractFlag )
    obj.Items["baseId"] = Lns_ToCollection( self.BaseId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    obj.Items["staticFlag"] = Lns_ToCollection( self.StaticFlag )
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["kind"] = Lns_ToCollection( self.Kind )
    obj.Items["mutMode"] = Lns_ToCollection( self.MutMode )
    obj.Items["asyncMode"] = Lns_ToCollection( self.AsyncMode )
    obj.Items["ifList"] = Lns_ToCollection( self.IfList )
    obj.Items["itemTypeId"] = Lns_ToCollection( self.ItemTypeId )
    obj.Items["argTypeId"] = Lns_ToCollection( self.ArgTypeId )
    obj.Items["retTypeId"] = Lns_ToCollection( self.RetTypeId )
    obj.Items["children"] = Lns_ToCollection( self.Children )
    obj.Items["moduleLang"] = Lns_ToCollection( self.ModuleLang )
    obj.Items["requirePath"] = Lns_ToCollection( self.RequirePath )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoNormal) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoNormal__fromMap_1718_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoNormal_FromMap( arg1, paramList )
}
func Import__TypeInfoNormal__fromStem_1722_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoNormal_FromMap( arg1, paramList )
}
func Import__TypeInfoNormal_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoNormal_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoNormal_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoNormal{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoNormal_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoNormal_FromMapMain( newObj *Import__TypeInfoNormal, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["abstractFlag"], false, nil); !ok {
       return false,nil,"abstractFlag:" + mess.(string)
    } else {
       newObj.AbstractFlag = conv.(bool)
    }
    if ok,conv,mess := Import__IdInfo_FromMapSub( objMap.Items["baseId"], false, nil); !ok {
       return false,nil,"baseId:" + mess.(string)
    } else {
       newObj.BaseId = conv.(*Import__IdInfo)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["staticFlag"], false, nil); !ok {
       return false,nil,"staticFlag:" + mess.(string)
    } else {
       newObj.StaticFlag = conv.(bool)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["kind"], false, nil); !ok {
       return false,nil,"kind:" + mess.(string)
    } else {
       newObj.Kind = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["mutMode"], false, nil); !ok {
       return false,nil,"mutMode:" + mess.(string)
    } else {
       newObj.MutMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["asyncMode"], false, nil); !ok {
       return false,nil,"asyncMode:" + mess.(string)
    } else {
       newObj.AsyncMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["ifList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Import__IdInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"ifList:" + mess.(string)
    } else {
       newObj.IfList = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["itemTypeId"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Import__IdInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"itemTypeId:" + mess.(string)
    } else {
       newObj.ItemTypeId = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["argTypeId"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Import__IdInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"argTypeId:" + mess.(string)
    } else {
       newObj.ArgTypeId = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["retTypeId"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Import__IdInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"retTypeId:" + mess.(string)
    } else {
       newObj.RetTypeId = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["children"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Import__IdInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"children:" + mess.(string)
    } else {
       newObj.Children = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["moduleLang"], true, nil); !ok {
       return false,nil,"moduleLang:" + mess.(string)
    } else {
       newObj.ModuleLang = conv
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["requirePath"], true, nil); !ok {
       return false,nil,"requirePath:" + mess.(string)
    } else {
       newObj.RequirePath = conv
    }
    return true, newObj, nil
}
// 444: decl @lune.@base.@Import._TypeInfoNormal.createTypeInfo
func (self *Import__TypeInfoNormal) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@Import._TypeInfoNormal.createTypeInfo"
    var newTypeInfo LnsAny
    newTypeInfo = nil
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.ParentId != Ast_rootTypeId) ||
        _env.SetStackVal( Lns_op_not(Ast_getBuiltInTypeIdMap(_env).Get(self.TypeId))) ||
        _env.SetStackVal( self.Kind == Ast_TypeInfoKind__List) ||
        _env.SetStackVal( self.Kind == Ast_TypeInfoKind__Array) ||
        _env.SetStackVal( self.Kind == Ast_TypeInfoKind__Map) ||
        _env.SetStackVal( self.Kind == Ast_TypeInfoKind__Set) ).(bool){
        var parentInfo *Ast_TypeInfo
        parentInfo = Ast_headTypeInfo
        if self.ParentId != Ast_rootTypeId{
            var workTypeInfo *Ast_TypeInfo
            
            {
                _workTypeInfo := Import_convExp2076(Lns_2DDD(param.FP.GetTypeInfo(_env, self.ParentId)))
                if _workTypeInfo == nil{
                    return nil, _env.LuaVM.String_format("not found parentInfo %d %s", []LnsAny{self.ParentId, self.Txt})
                } else {
                    workTypeInfo = _workTypeInfo.(*Ast_TypeInfo)
                }
            }
            parentInfo = workTypeInfo
            
        }
        var itemTypeInfo *LnsList
        itemTypeInfo = NewLnsList([]LnsAny{})
        for _, _typeId := range( self.ItemTypeId.Items ) {
            typeId := _typeId.(Import__IdInfoDownCast).ToImport__IdInfo()
            itemTypeInfo.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, typeId))).(*Ast_TypeInfo)))
        }
        var argTypeInfo *LnsList
        argTypeInfo = NewLnsList([]LnsAny{})
        for _index, _typeId := range( self.ArgTypeId.Items ) {
            index := _index + 1
            typeId := _typeId.(Import__IdInfoDownCast).ToImport__IdInfo()
            var argType LnsAny
            var mess LnsAny
            argType,mess = param.FP.GetTypeInfoFrom(_env, typeId)
            if argType != nil{
                argType_328 := argType.(*Ast_TypeInfo)
                argTypeInfo.Insert(Ast_TypeInfo2Stem(argType_328))
            } else {
                var errmess string
                errmess = _env.LuaVM.String_format("not found arg (index:%d) -- %s.%s, %d, %d. %s", []LnsAny{index, parentInfo.FP.GetTxt(_env, nil, nil, nil), self.Txt, typeId.Id, self.ArgTypeId.Len(), mess})
                return nil, errmess
            }
        }
        var retTypeInfo *LnsList
        retTypeInfo = NewLnsList([]LnsAny{})
        for _, _typeId := range( self.RetTypeId.Items ) {
            typeId := _typeId.(Import__IdInfoDownCast).ToImport__IdInfo()
            retTypeInfo.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, typeId))).(*Ast_TypeInfo)))
        }
        var baseInfo *Ast_TypeInfo
        baseInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, self.BaseId))).(*Ast_TypeInfo)
        var interfaceList *LnsList
        interfaceList = NewLnsList([]LnsAny{})
        for _, _ifTypeId := range( self.IfList.Items ) {
            ifTypeId := _ifTypeId.(Import__IdInfoDownCast).ToImport__IdInfo()
            interfaceList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, ifTypeId))).(*Ast_TypeInfo)))
        }
        var parentScope *Ast_Scope
        
        {
            _parentScope := param.TypeId2Scope.Get(self.ParentId)
            if _parentScope == nil{
                return nil, _env.LuaVM.String_format("%s: not found parentScope %s %s", []LnsAny{__func__, self.ParentId, self.Txt})
            } else {
                parentScope = _parentScope.(*Ast_Scope)
            }
        }
        if self.Txt != ""{
            newTypeInfo = parentScope.FP.GetTypeInfoChild(_env, self.Txt)
            
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( newTypeInfo) &&
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.Kind == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( self.Kind == Ast_TypeInfoKind__ExtModule) ||
                _env.SetStackVal( self.Kind == Ast_TypeInfoKind__IF) ).(bool))) )){
            panic("internal error")
        } else { 
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.Kind == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( self.Kind == Ast_TypeInfoKind__IF) ).(bool){
                Log_log(_env, Log_Level__Debug, __func__, 515, Log_CreateMessage(func(_env *LnsEnv) string {
                    return _env.LuaVM.String_format("new type -- %d, %s -- %s, %d", []LnsAny{self.ParentId, self.Txt, _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(parentScope.FP.Get_ownerTypeInfo(_env)) && 
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetFullName(_env, Ast_defaultTypeNameCtrl, parentScope.FP, false)})/* 1:65 */)) ||
                        _env.SetStackVal( "nil") ).(string), _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(parentScope.FP.Get_ownerTypeInfo(_env)) && 
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_typeId(_env)})&&
                        _env.NilAccPush(_env.NilAccPop().(*Ast_IdInfo).Id))) ||
                        _env.SetStackVal( -1) ).(LnsInt)})
                }))
                
                var baseScope *Ast_Scope
                baseScope = Lns_unwrap( baseInfo.FP.Get_scope(_env)).(*Ast_Scope)
                var ifScopeList *LnsList
                ifScopeList = NewLnsList([]LnsAny{})
                for _, _ifType := range( interfaceList.Items ) {
                    ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    ifScopeList.Insert(Ast_Scope2Stem(Lns_unwrap( ifType.FP.Get_scope(_env)).(*Ast_Scope)))
                }
                var scope *Ast_Scope
                scope = NewAst_Scope(_env, param.ProcessInfo, parentScope, true, baseScope, ifScopeList)
                var altTypeList *LnsList
                altTypeList = NewLnsList([]LnsAny{})
                for _, _itemType := range( itemTypeInfo.Items ) {
                    itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    altTypeList.Insert(Ast_AlternateTypeInfo2Stem(Lns_unwrap( (Ast_AlternateTypeInfoDownCastF(itemType.FP))).(*Ast_AlternateTypeInfo)))
                }
                var workTypeInfo *Ast_TypeInfo
                workTypeInfo = param.ProcessInfo.FP.CreateClass(_env, self.Kind == Ast_TypeInfoKind__Class, self.AbstractFlag, scope, baseInfo, interfaceList, altTypeList, parentInfo, true, Ast_AccessMode__Pub, self.Txt)
                parentScope.FP.AddClassLazy(_env, param.ProcessInfo, self.Txt, nil, workTypeInfo, param.LazyModuleSet.Has(self.TypeId))
                newTypeInfo = workTypeInfo
                
                param.TypeId2Scope.Set(self.TypeId,scope)
                param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
                workTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
            } else if self.Kind == Ast_TypeInfoKind__ExtModule{
                Log_log(_env, Log_Level__Debug, __func__, 553, Log_CreateMessage(func(_env *LnsEnv) string {
                    return _env.LuaVM.String_format("new type -- %d, %s -- %s, %d", []LnsAny{self.ParentId, self.Txt, _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(parentScope.FP.Get_ownerTypeInfo(_env)) && 
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetFullName(_env, Ast_defaultTypeNameCtrl, parentScope.FP, false)})/* 1:65 */)) ||
                        _env.SetStackVal( "nil") ).(string), _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(parentScope.FP.Get_ownerTypeInfo(_env)) && 
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_typeId(_env)})&&
                        _env.NilAccPush(_env.NilAccPop().(*Ast_IdInfo).Id))) ||
                        _env.SetStackVal( -1) ).(LnsInt)})
                }))
                
                var scope *Ast_Scope
                scope = NewAst_Scope(_env, param.ProcessInfo, parentScope, true, nil, NewLnsList([]LnsAny{}))
                var workTypeInfo *Ast_TypeInfo
                workTypeInfo = param.ProcessInfo.FP.CreateExtModule(_env, scope, parentInfo, true, Ast_AccessMode__Pub, self.Txt, Lns_unwrap( self.ModuleLang).(LnsInt), Lns_unwrap( self.RequirePath).(string))
                parentScope.FP.AddExtModule(_env, param.ProcessInfo, self.Txt, nil, workTypeInfo, param.LazyModuleSet.Has(self.TypeId), Lns_unwrap( self.ModuleLang).(LnsInt))
                newTypeInfo = workTypeInfo
                
                param.TypeId2Scope.Set(self.TypeId,scope)
                param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
                workTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
            } else { 
                var scope LnsAny
                scope = nil
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( self.Kind == Ast_TypeInfoKind__Func) ||
                    _env.SetStackVal( self.Kind == Ast_TypeInfoKind__Method) ).(bool){
                    scope = NewAst_Scope(_env, param.ProcessInfo, parentScope, false, nil, nil)
                    
                }
                var typeInfoKind LnsInt
                typeInfoKind = self.Kind
                var accessMode LnsInt
                accessMode = self.AccessMode
                var workTypeInfo *Ast_TypeInfo
                workTypeInfo = Ast_NormalTypeInfo_create(_env, param.ProcessInfo, accessMode, self.AbstractFlag, scope, baseInfo, parentInfo, self.StaticFlag, typeInfoKind, self.Txt, itemTypeInfo, argTypeInfo, retTypeInfo, self.MutMode, self.AsyncMode)
                newTypeInfo = workTypeInfo
                
                param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
                workTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
                if _switch3204 := self.Kind; _switch3204 == Ast_TypeInfoKind__Func || _switch3204 == Ast_TypeInfoKind__Method || _switch3204 == Ast_TypeInfoKind__Macro || _switch3204 == Ast_TypeInfoKind__Form || _switch3204 == Ast_TypeInfoKind__FormFunc {
                    var symbolKind LnsInt
                    symbolKind = Ast_SymbolKind__Fun
                    if _switch3142 := self.Kind; _switch3142 == Ast_TypeInfoKind__Method {
                        symbolKind = Ast_SymbolKind__Mtd
                        
                    } else if _switch3142 == Ast_TypeInfoKind__Macro {
                        symbolKind = Ast_SymbolKind__Mac
                        
                    } else if _switch3142 == Ast_TypeInfoKind__Form || _switch3142 == Ast_TypeInfoKind__FormFunc {
                        symbolKind = Ast_SymbolKind__Typ
                        
                    }
                    var workParentScope *Ast_Scope
                    workParentScope = Lns_unwrap( param.TypeId2Scope.Get(self.ParentId)).(*Ast_Scope)
                    workParentScope.FP.Add(_env, param.ProcessInfo, symbolKind, false, self.Kind == Ast_TypeInfoKind__Func, self.Txt, nil, workTypeInfo, accessMode, self.StaticFlag, Ast_MutMode__IMut, true, false)
                    param.TypeId2Scope.Set(self.TypeId,scope)
                }
            }
        }
    } else { 
        newTypeInfo = param.Scope.FP.GetTypeInfo(_env, self.Txt, param.Scope, false, param.ScopeAccess)
        
        if newTypeInfo != nil{
            newTypeInfo_375 := newTypeInfo.(*Ast_TypeInfo)
            param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo_375)
            newTypeInfo_375.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
        } else {
            for _key, _val := range( self.FP.ToMap().Items ) {
                key := _key.(string)
                val := _val
                Util_errorLog(_env, _env.LuaVM.String_format("error: illegal self %s:%s", []LnsAny{key, val}))
            }
        }
    }
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoEnum
type Import__TypeInfoEnumMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoEnum struct {
    Import__TypeInfo
    ParentId LnsInt
    Txt string
    AccessMode LnsInt
    ValTypeId LnsInt
    EnumValList *LnsMap
    FP Import__TypeInfoEnumMtd
}
func Import__TypeInfoEnum2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoEnum).FP
}
type Import__TypeInfoEnumDownCast interface {
    ToImport__TypeInfoEnum() *Import__TypeInfoEnum
}
func Import__TypeInfoEnumDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoEnumDownCast)
    if ok { return work.ToImport__TypeInfoEnum() }
    return nil
}
func (obj *Import__TypeInfoEnum) ToImport__TypeInfoEnum() *Import__TypeInfoEnum {
    return obj
}
func NewImport__TypeInfoEnum(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 LnsInt, arg5 *LnsMap) *Import__TypeInfoEnum {
    obj := &Import__TypeInfoEnum{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoEnum(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Import__TypeInfoEnum) InitImport__TypeInfoEnum(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 LnsInt, arg5 *LnsMap) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.ParentId = arg1
    self.Txt = arg2
    self.AccessMode = arg3
    self.ValTypeId = arg4
    self.EnumValList = arg5
}
func (self *Import__TypeInfoEnum) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["valTypeId"] = Lns_ToCollection( self.ValTypeId )
    obj.Items["enumValList"] = Lns_ToCollection( self.EnumValList )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoEnum) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoEnum__fromMap_1781_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoEnum_FromMap( arg1, paramList )
}
func Import__TypeInfoEnum__fromStem_1785_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoEnum_FromMap( arg1, paramList )
}
func Import__TypeInfoEnum_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoEnum_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoEnum_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoEnum{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoEnum_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoEnum_FromMapMain( newObj *Import__TypeInfoEnum, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["valTypeId"], false, nil); !ok {
       return false,nil,"valTypeId:" + mess.(string)
    } else {
       newObj.ValTypeId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToLnsMapSub( objMap.Items["enumValList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil},Lns_ToObjParam{
            Lns_ToStemSub, false,nil}}); !ok {
       return false,nil,"enumValList:" + mess.(string)
    } else {
       newObj.EnumValList = conv.(*LnsMap)
    }
    return true, newObj, nil
}
// 649: decl @lune.@base.@Import._TypeInfoEnum.createTypeInfo
func (self *Import__TypeInfoEnum) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    var accessMode LnsInt
    accessMode = Lns_unwrap( Ast_AccessMode__from(_env, LnsInt(self.AccessMode))).(LnsInt)
    var parentInfo *Ast_TypeInfo
    parentInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(_env, self.ParentId))).(*Ast_TypeInfo)
    var parentScope *Ast_Scope
    parentScope = Lns_unwrap( param.TypeId2Scope.Get(self.ParentId)).(*Ast_Scope)
    var scope *Ast_Scope
    scope = NewAst_Scope(_env, param.ProcessInfo, parentScope, true, nil, nil)
    param.TypeId2Scope.Set(self.TypeId,scope)
    var valTypeInfo *Ast_TypeInfo
    valTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(_env, self.ValTypeId))).(*Ast_TypeInfo)
    var enumTypeInfo *Ast_EnumTypeInfo
    enumTypeInfo = param.ProcessInfo.FP.CreateEnum(_env, scope, parentInfo, true, accessMode, self.Txt, valTypeInfo)
    var newTypeInfo *Ast_EnumTypeInfo
    newTypeInfo = enumTypeInfo
    param.TypeId2TypeInfo.Set(self.TypeId,&enumTypeInfo.Ast_TypeInfo)
    enumTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    var getEnumLiteral func(_env *LnsEnv, val LnsAny) LnsAny
    getEnumLiteral = func(_env *LnsEnv, val LnsAny) LnsAny {
        if _switch3519 := valTypeInfo; _switch3519 == Ast_builtinTypeInt {
            return &Ast_EnumLiteral__Int{Lns_forceCastInt(val)}
        } else if _switch3519 == Ast_builtinTypeReal {
            return &Ast_EnumLiteral__Real{Lns_forceCastReal(val)}
        } else if _switch3519 == Ast_builtinTypeString {
            return &Ast_EnumLiteral__Str{val.(string)}
        }
        return nil
    }
    for _valName, _valData := range( self.EnumValList.Items ) {
        valName := _valName.(string)
        valData := _valData
        var val LnsAny
        
        {
            _val := getEnumLiteral(_env, valData)
            if _val == nil{
                return nil, _env.LuaVM.String_format("unknown enum val type -- %s", []LnsAny{valTypeInfo.FP.GetTxt(_env, nil, nil, nil)})
            } else {
                val = _val
            }
        }
        var evalValSym *Ast_SymbolInfo
        evalValSym = Lns_unwrap( Lns_car(scope.FP.AddEnumVal(_env, param.ProcessInfo, valName, nil, &enumTypeInfo.Ast_TypeInfo))).(*Ast_SymbolInfo)
        enumTypeInfo.FP.AddEnumValInfo(_env, NewAst_EnumValInfo(_env, valName, val, evalValSym))
    }
    parentScope.FP.AddEnum(_env, param.ProcessInfo, accessMode, self.Txt, nil, &enumTypeInfo.Ast_TypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoAlgeVal
type Import__TypeInfoAlgeValMtd interface {
    ToMap() *LnsMap
}
type Import__TypeInfoAlgeVal struct {
    Name string
    TypeList *LnsList
    FP Import__TypeInfoAlgeValMtd
}
func Import__TypeInfoAlgeVal2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoAlgeVal).FP
}
type Import__TypeInfoAlgeValDownCast interface {
    ToImport__TypeInfoAlgeVal() *Import__TypeInfoAlgeVal
}
func Import__TypeInfoAlgeValDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoAlgeValDownCast)
    if ok { return work.ToImport__TypeInfoAlgeVal() }
    return nil
}
func (obj *Import__TypeInfoAlgeVal) ToImport__TypeInfoAlgeVal() *Import__TypeInfoAlgeVal {
    return obj
}
func NewImport__TypeInfoAlgeVal(_env *LnsEnv, arg1 string, arg2 *LnsList) *Import__TypeInfoAlgeVal {
    obj := &Import__TypeInfoAlgeVal{}
    obj.FP = obj
    obj.InitImport__TypeInfoAlgeVal(_env, arg1, arg2)
    return obj
}
func (self *Import__TypeInfoAlgeVal) InitImport__TypeInfoAlgeVal(_env *LnsEnv, arg1 string, arg2 *LnsList) {
    self.Name = arg1
    self.TypeList = arg2
}
func (self *Import__TypeInfoAlgeVal) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["name"] = Lns_ToCollection( self.Name )
    obj.Items["typeList"] = Lns_ToCollection( self.TypeList )
    return obj
}
func (self *Import__TypeInfoAlgeVal) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoAlgeVal__fromMap_1805_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlgeVal_FromMap( arg1, paramList )
}
func Import__TypeInfoAlgeVal__fromStem_1809_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlgeVal_FromMap( arg1, paramList )
}
func Import__TypeInfoAlgeVal_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoAlgeVal_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoAlgeVal_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoAlgeVal{}
    newObj.FP = newObj
    return Import__TypeInfoAlgeVal_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoAlgeVal_FromMapMain( newObj *Import__TypeInfoAlgeVal, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["name"], false, nil); !ok {
       return false,nil,"name:" + mess.(string)
    } else {
       newObj.Name = conv.(string)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["typeList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Import__IdInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"typeList:" + mess.(string)
    } else {
       newObj.TypeList = conv.(*LnsList)
    }
    return true, newObj, nil
}

// declaration Class -- _TypeInfoAlge
type Import__TypeInfoAlgeMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(_env *LnsEnv, arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoAlge struct {
    Import__TypeInfo
    ParentId LnsInt
    Txt string
    AccessMode LnsInt
    AlgeValList *LnsList
    FP Import__TypeInfoAlgeMtd
}
func Import__TypeInfoAlge2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import__TypeInfoAlge).FP
}
type Import__TypeInfoAlgeDownCast interface {
    ToImport__TypeInfoAlge() *Import__TypeInfoAlge
}
func Import__TypeInfoAlgeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import__TypeInfoAlgeDownCast)
    if ok { return work.ToImport__TypeInfoAlge() }
    return nil
}
func (obj *Import__TypeInfoAlge) ToImport__TypeInfoAlge() *Import__TypeInfoAlge {
    return obj
}
func NewImport__TypeInfoAlge(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *LnsList) *Import__TypeInfoAlge {
    obj := &Import__TypeInfoAlge{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoAlge(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *Import__TypeInfoAlge) InitImport__TypeInfoAlge(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *LnsList) {
    self.Import__TypeInfo.InitImport__TypeInfo( _env)
    self.ParentId = arg1
    self.Txt = arg2
    self.AccessMode = arg3
    self.AlgeValList = arg4
}
func (self *Import__TypeInfoAlge) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["algeValList"] = Lns_ToCollection( self.AlgeValList )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoAlge) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoAlge__fromMap_1864_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlge_FromMap( arg1, paramList )
}
func Import__TypeInfoAlge__fromStem_1868_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlge_FromMap( arg1, paramList )
}
func Import__TypeInfoAlge_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Import__TypeInfoAlge_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Import__TypeInfoAlge_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Import__TypeInfoAlge{}
    newObj.FP = newObj
    newObj.Import__TypeInfo.FP = newObj
    return Import__TypeInfoAlge_FromMapMain( newObj, objMap, paramList )
}
func Import__TypeInfoAlge_FromMapMain( newObj *Import__TypeInfoAlge, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := Import__TypeInfo_FromMapMain( &newObj.Import__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["algeValList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Import__TypeInfoAlgeVal_FromMapSub, false,nil}}); !ok {
       return false,nil,"algeValList:" + mess.(string)
    } else {
       newObj.AlgeValList = conv.(*LnsList)
    }
    return true, newObj, nil
}
// 703: decl @lune.@base.@Import._TypeInfoAlge.createTypeInfo
func (self *Import__TypeInfoAlge) CreateTypeInfo(_env *LnsEnv, param *Import_ImportParam)(LnsAny, LnsAny) {
    var accessMode LnsInt
    accessMode = Lns_unwrap( Ast_AccessMode__from(_env, LnsInt(self.AccessMode))).(LnsInt)
    var parentInfo *Ast_TypeInfo
    parentInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(_env, self.ParentId))).(*Ast_TypeInfo)
    var parentScope *Ast_Scope
    parentScope = Lns_unwrap( param.TypeId2Scope.Get(self.ParentId)).(*Ast_Scope)
    var scope *Ast_Scope
    scope = NewAst_Scope(_env, param.ProcessInfo, parentScope, true, nil, nil)
    param.TypeId2Scope.Set(self.TypeId,scope)
    var algeTypeInfo *Ast_AlgeTypeInfo
    algeTypeInfo = param.ProcessInfo.FP.CreateAlge(_env, scope, parentInfo, true, accessMode, self.Txt)
    var newTypeInfo *Ast_AlgeTypeInfo
    newTypeInfo = algeTypeInfo
    param.TypeId2TypeInfo.Set(self.TypeId,&algeTypeInfo.Ast_TypeInfo)
    algeTypeInfo.FP.Get_typeId(_env).FP.Set_orgId(_env, self.TypeId)
    for _, _valInfo := range( self.AlgeValList.Items ) {
        valInfo := _valInfo.(Import__TypeInfoAlgeValDownCast).ToImport__TypeInfoAlgeVal()
        var typeInfoList *LnsList
        typeInfoList = NewLnsList([]LnsAny{})
        for _, _orgTypeId := range( valInfo.TypeList.Items ) {
            orgTypeId := _orgTypeId.(Import__IdInfoDownCast).ToImport__IdInfo()
            typeInfoList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(_env, orgTypeId))).(*Ast_TypeInfo)))
        }
        var algeValSym LnsAny
        algeValSym = Import_convExp3841(Lns_2DDD(scope.FP.AddAlgeVal(_env, param.ProcessInfo, valInfo.Name, nil, &algeTypeInfo.Ast_TypeInfo)))
        var algeVal *Ast_AlgeValInfo
        algeVal = NewAst_AlgeValInfo(_env, valInfo.Name, typeInfoList, algeTypeInfo, Lns_unwrap( algeValSym).(*Ast_SymbolInfo))
        algeTypeInfo.FP.AddValInfo(_env, algeVal)
    }
    parentScope.FP.AddAlge(_env, param.ProcessInfo, accessMode, self.Txt, nil, &algeTypeInfo.Ast_TypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- DependModuleInfo
type Import_DependModuleInfoMtd interface {
    GetTypeInfo(_env *LnsEnv, arg1 LnsInt) *Ast_TypeInfo
}
type Import_DependModuleInfo struct {
    id LnsInt
    metaTypeId2TypeInfoMap *LnsMap
    FP Import_DependModuleInfoMtd
}
func Import_DependModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Import_DependModuleInfo).FP
}
type Import_DependModuleInfoDownCast interface {
    ToImport_DependModuleInfo() *Import_DependModuleInfo
}
func Import_DependModuleInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Import_DependModuleInfoDownCast)
    if ok { return work.ToImport_DependModuleInfo() }
    return nil
}
func (obj *Import_DependModuleInfo) ToImport_DependModuleInfo() *Import_DependModuleInfo {
    return obj
}
func NewImport_DependModuleInfo(_env *LnsEnv, arg1 LnsInt, arg2 *LnsMap) *Import_DependModuleInfo {
    obj := &Import_DependModuleInfo{}
    obj.FP = obj
    obj.InitImport_DependModuleInfo(_env, arg1, arg2)
    return obj
}
func (self *Import_DependModuleInfo) InitImport_DependModuleInfo(_env *LnsEnv, arg1 LnsInt, arg2 *LnsMap) {
    self.id = arg1
    self.metaTypeId2TypeInfoMap = arg2
}
// 748: decl @lune.@base.@Import.DependModuleInfo.getTypeInfo
func (self *Import_DependModuleInfo) GetTypeInfo(_env *LnsEnv, metaTypeId LnsInt) *Ast_TypeInfo {
    return Lns_unwrap( self.metaTypeId2TypeInfoMap.Get(metaTypeId)).(*Ast_TypeInfo)
}


func Lns_Import_init(_env *LnsEnv) {
    if init_Import { return }
    init_Import = true
    Import__mod__ = "@lune.@base.@Import"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_Meta_init(_env)
    Lns_Parser_init(_env)
    Lns_Util_init(_env)
    Lns_Ast_init(_env)
    Lns_Macro_init(_env)
    Lns_Nodes_init(_env)
    Lns_frontInterface_init(_env)
    Lns_Log_init(_env)
    Lns_TransUnitIF_init(_env)
    Lns_Builtin_init(_env)
}
func init() {
    init_Import = false
}
