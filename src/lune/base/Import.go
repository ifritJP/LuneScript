// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Import bool
var Import__mod__ string
// for 804
func Import_convExp4303(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 807
func Import_convExp4326(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 811
func Import_convExp4358(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 822
func Import_convExp4427(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 825
func Import_convExp4450(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 828
func Import_convExp4473(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 831
func Import_convExp4496(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 834
func Import_convExp4519(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 837
func Import_convExp4542(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 840
func Import_convExp4565(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 843
func Import_convExp4588(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Import__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 186
func Import_convExp567(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 209
func Import_convExp696(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 328
func Import_convExp1294(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 346
func Import_convExp1392(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 424
func Import_convExp1907(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 694
func Import_convExp3679(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 989
func Import_convExp5324(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}










// declaration Class -- Import
type Import_ImportMtd interface {
    Get_importModule2ModuleInfo() *LnsMap
    ProcessImport(arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsInt) *FrontInterface_ModuleInfo
    processImportSub(arg1 *Ast_ProcessInfo, arg2 *FrontInterface_ModuleMeta, arg3 string, arg4 string, arg5 *LnsList, arg6 LnsInt) *FrontInterface_ModuleInfo
}
type Import_Import struct {
    transUnitIF TransUnitIF_TransUnitIF
    importModuleInfo *FrontInterface_ImportModuleInfo
    moduleType *Ast_TypeInfo
    builtinFunc *Builtin_BuiltinFuncType
    globalScope *Ast_Scope
    macroCtrl *Macro_MacroCtrl
    typeNameCtrl *Ast_TypeNameCtrl
    helperInfo *Nodes_LuneHelperInfo
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
func NewImport_Import(arg1 TransUnitIF_TransUnitIF, arg2 *FrontInterface_ImportModuleInfo, arg3 *Ast_TypeInfo, arg4 *Builtin_BuiltinFuncType, arg5 *Ast_Scope, arg6 *Macro_MacroCtrl, arg7 *Ast_TypeNameCtrl, arg8 *Nodes_LuneHelperInfo, arg9 *LnsMap, arg10 bool) *Import_Import {
    obj := &Import_Import{}
    obj.FP = obj
    obj.InitImport_Import(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Import_Import) Get_importModule2ModuleInfo() *LnsMap{ return self.importModule2ModuleInfo }
// 67: DeclConstr
func (self *Import_Import) InitImport_Import(transUnitIF TransUnitIF_TransUnitIF,importModuleInfo *FrontInterface_ImportModuleInfo,moduleType *Ast_TypeInfo,builtinFunc *Builtin_BuiltinFuncType,globalScope *Ast_Scope,macroCtrl *Macro_MacroCtrl,typeNameCtrl *Ast_TypeNameCtrl,helperInfo *Nodes_LuneHelperInfo,importedAliasMap *LnsMap,validMutControl bool) {
    self.validMutControl = validMutControl
    
    self.transUnitIF = transUnitIF
    
    self.importModuleInfo = importModuleInfo
    
    self.moduleType = moduleType
    
    self.builtinFunc = builtinFunc
    
    self.globalScope = globalScope
    
    self.macroCtrl = macroCtrl
    
    self.typeNameCtrl = typeNameCtrl
    
    self.helperInfo = helperInfo
    
    self.importedAliasMap = importedAliasMap
    
    self.importModule2ModuleInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.importModuleName2ModuleInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
}


// 722: decl @lune.@base.@Import.Import.processImportSub
func (self *Import_Import) processImportSub(processInfo *Ast_ProcessInfo,moduleMeta *FrontInterface_ModuleMeta,orgModulePath string,modulePath string,nameList *LnsList,depth LnsInt) *FrontInterface_ModuleInfo {
    __func__ := "@lune.@base.@Import.Import.processImportSub"
    var metaInfo *Lns_luaValue
    metaInfo = moduleMeta.FP.Get_metaInfo().(*Lns_luaValue)
    Log_log(Log_Level__Info, __func__, 728, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s processing", []LnsAny{orgModulePath})
    }))
    
    var dependLibId2DependInfo *LnsMap
    dependLibId2DependInfo = NewLnsMap( map[LnsAny]LnsAny{})
    {
        _exp3916 := metaInfo.GetAt( "__dependModuleMap" ).(*Lns_luaValue)
        _sorted3916 := Lns_getVM().SortMapKeyList( _exp3916 )
        _index3916, _key3916 := _sorted3916.Get1stFromMap()
        for _index3916 != nil {
            dependName := _key3916.(string)
            dependInfo := _exp3916.GetAt( _key3916 ).(*Lns_luaValue)
            if Lns_isCondTrue( dependInfo.GetAt("use")){
                var moduleInfo *FrontInterface_ModuleInfo
                moduleInfo = self.FP.ProcessImport(processInfo, dependName, depth + 1)
                var typeId LnsInt
                typeId = Lns_forceCastInt((Lns_unwrap( dependInfo.GetAt("typeId"))))
                dependLibId2DependInfo.Set(typeId,moduleInfo)
            }
            _index3916, _key3916 = _sorted3916.NextFromMap( _index3916 )
        }
    }
    var typeId2TypeInfo *LnsMap
    typeId2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    typeId2TypeInfo.Set(Ast_rootTypeId,Ast_headTypeInfo)
    var typeId2Scope *LnsMap
    typeId2Scope = NewLnsMap( map[LnsAny]LnsAny{})
    typeId2Scope.Set(Ast_rootTypeId,self.transUnitIF.Get_scope())
    typeId2TypeInfo.Set(self.builtinFunc.Lnsthread_.FP.Get_typeId().Id,self.builtinFunc.Lnsthread_)
    typeId2Scope.Set(self.builtinFunc.Lnsthread_.FP.Get_typeId().Id,self.builtinFunc.Lnsthread_.FP.Get_scope())
    {
        _exp4066 := metaInfo.GetAt( "__dependIdMap" ).(*Lns_luaValue)
        _key4066, _val4066 := _exp4066.Get1stFromMap()
        for _key4066 != nil {
            typeId := _key4066.(LnsInt)
            dependIdInfo := _val4066.(*Lns_luaValue)
            var dependInfo *FrontInterface_ModuleInfo
            dependInfo = Lns_unwrap( dependLibId2DependInfo.Get(Lns_unwrap( dependIdInfo.GetAt(1)).(LnsInt))).(*FrontInterface_ModuleInfo)
            var typeInfo *Ast_TypeInfo
            typeInfo = Lns_unwrap( dependInfo.FP.GetTypeInfo(Lns_unwrap( dependIdInfo.GetAt(2)).(LnsInt))).(*Ast_TypeInfo)
            typeId2TypeInfo.Set(typeId,typeInfo)
            {
                __exp := Ast_getScope(typeInfo)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Ast_Scope)
                    typeId2Scope.Set(typeId,_exp)
                }
            }
            _key4066, _val4066 = _exp4066.NextFromMap( _key4066 )
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
        moduleTypeInfo = self.transUnitIF.PushModule(processInfo, true, moduleName, mutable)
        
    }
    for range( nameList.Items ) {
        self.transUnitIF.PopModule()
    }
    for _, _symbolInfo := range( Ast_getSym2builtInTypeMap().Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        typeId2TypeInfo.Set(symbolInfo.FP.Get_typeInfo().FP.Get_typeId().Id,symbolInfo.FP.Get_typeInfo())
    }
    for _, _builtinTypeInfo := range( Ast_getBuiltInTypeIdMap().Items ) {
        builtinTypeInfo := _builtinTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        typeId2TypeInfo.Set(builtinTypeInfo.FP.Get_typeId().Id,builtinTypeInfo)
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
        _exp4670 := metaInfo.GetAt( "__typeInfoList" ).(*Lns_luaValue)
        _key4670, _val4670 := _exp4670.Get1stFromMap()
        for _key4670 != nil {
            atomInfoLua := _val4670.(*Lns_luaValue)
            var workAtomInfo LnsAny
            
            {
                _workAtomInfo := Lns_getVM().ExpandLuavalMap(atomInfoLua)
                if _workAtomInfo == nil{
                    self.transUnitIF.Error("illegal atomInfo")
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
                    kind = Lns_unwrap( Ast_SerializeKind__from(Lns_forceCastInt(skind))).(LnsInt)
                    if _switch4591 := kind; _switch4591 == Ast_SerializeKind__Enum {
                        actInfo, mess = Import_convExp4303(Lns_2DDD(Import__TypeInfoEnum__fromMap_1341_(atomInfo,nil)))
                        
                    } else if _switch4591 == Ast_SerializeKind__Alge {
                        actInfo, mess = Import_convExp4326(Lns_2DDD(Import__TypeInfoAlge__fromMap_1383_(atomInfo,nil)))
                        
                        self.helperInfo.UseAlge = true
                        
                    } else if _switch4591 == Ast_SerializeKind__Module {
                        actInfo, mess = Import_convExp4358(Lns_2DDD(Import__TypeInfoModule__fromMap_1253_(atomInfo,nil)))
                        
                    } else if _switch4591 == Ast_SerializeKind__Normal {
                        var workInfo LnsAny
                        workInfo, mess = Import__TypeInfoNormal__fromMap_1319_(atomInfo,nil)
                        
                        if workInfo != nil{
                            workInfo_4093 := workInfo.(*Import__TypeInfoNormal)
                            _typeInfoNormalList.Insert(Import__TypeInfoNormal2Stem(workInfo_4093))
                        }
                        actInfo = Import__TypeInfoDownCastF(workInfo)
                        
                    } else if _switch4591 == Ast_SerializeKind__Nilable {
                        actInfo, mess = Import_convExp4427(Lns_2DDD(Import__TypeInfoNilable__fromMap_1101_(atomInfo,nil)))
                        
                    } else if _switch4591 == Ast_SerializeKind__Alias {
                        actInfo, mess = Import_convExp4450(Lns_2DDD(Import__TypeInfoAlias__fromMap_1118_(atomInfo,nil)))
                        
                    } else if _switch4591 == Ast_SerializeKind__DDD {
                        actInfo, mess = Import_convExp4473(Lns_2DDD(Import__TypeInfoDDD__fromMap_1135_(atomInfo,nil)))
                        
                    } else if _switch4591 == Ast_SerializeKind__Alternate {
                        actInfo, mess = Import_convExp4496(Lns_2DDD(Import__TypeInfoAlternate__fromMap_1159_(atomInfo,nil)))
                        
                    } else if _switch4591 == Ast_SerializeKind__Generic {
                        actInfo, mess = Import_convExp4519(Lns_2DDD(Import__TypeInfoGeneric__fromMap_1183_(atomInfo,nil)))
                        
                    } else if _switch4591 == Ast_SerializeKind__Modifier {
                        actInfo, mess = Import_convExp4542(Lns_2DDD(Import__TypeInfoModifier__fromMap_1234_(atomInfo,nil)))
                        
                    } else if _switch4591 == Ast_SerializeKind__Box {
                        actInfo, mess = Import_convExp4565(Lns_2DDD(Import__TypeInfoBox__fromMap_1200_(atomInfo,nil)))
                        
                    } else if _switch4591 == Ast_SerializeKind__Ext {
                        actInfo, mess = Import_convExp4588(Lns_2DDD(Import__TypeInfoExt__fromMap_1217_(atomInfo,nil)))
                        
                    }
                    if actInfo != nil{
                        actInfo_4103 := actInfo.(*Import__TypeInfo)
                        _typeInfoList.Insert(Import__TypeInfo2Stem(actInfo_4103))
                        id2atomMap.Set(actInfo_4103.TypeId,actInfo_4103)
                    } else {
                        for _key, _val := range( atomInfo.Items ) {
                            key := _key.(string)
                            val := _val
                            Util_errorLog(Lns_getVM().String_format("table: %s:%s", []LnsAny{key, val}))
                        }
                        if mess != nil{
                            mess_4109 := mess.(string)
                            Util_errorLog(mess_4109)
                        }
                        Util_err(Lns_getVM().String_format("_TypeInfo.%s._fromMap error", []LnsAny{Ast_SerializeKind_getTxt( kind)}))
                    }
                }
            }
            _key4670, _val4670 = _exp4670.NextFromMap( _key4670 )
        }
    }
    var orgId2MacroTypeInfo *LnsMap
    orgId2MacroTypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    var lazyModuleSet *LnsSet
    lazyModuleSet = NewLnsSet([]LnsAny{})
    {
        _exp4710 := metaInfo.GetAt( "__lazyModuleList" ).(*Lns_luaValue)
        _key4710, _val4710 := _exp4710.Get1stFromMap()
        for _key4710 != nil {
            typeId := _val4710.(LnsInt)
            lazyModuleSet.Add(typeId)
            _key4710, _val4710 = _exp4710.NextFromMap( _key4710 )
        }
    }
    var modifier *TransUnitIF_Modifier
    modifier = NewTransUnitIF_Modifier(self.validMutControl, processInfo)
    var importParam *Import_ImportParam
    importParam = NewImport_ImportParam(self.transUnitIF.GetLatestPos(), modifier, processInfo, typeId2Scope, typeId2TypeInfo, NewLnsMap( map[LnsAny]LnsAny{}), lazyModuleSet, metaInfo, self.transUnitIF.Get_scope(), moduleTypeInfo, Ast_ScopeAccess__Normal, id2atomMap)
    for _, _atomInfo := range( _typeInfoList.Items ) {
        atomInfo := _atomInfo.(Import__TypeInfoDownCast).ToImport__TypeInfo()
        var newTypeInfo LnsAny
        var errMess LnsAny
        newTypeInfo,errMess = atomInfo.FP.CreateTypeInfoCache(importParam)
        {
            __exp := errMess
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(string)
                Util_err(Lns_getVM().String_format("Failed to createType -- %s: %s(%d): %s", []LnsAny{orgModulePath, Ast_SerializeKind_getTxt( atomInfo.Skind), atomInfo.TypeId, _exp}))
            }
        }
        if newTypeInfo != nil{
            newTypeInfo_4123 := newTypeInfo.(*Ast_TypeInfo)
            if newTypeInfo_4123.FP.Get_kind() == Ast_TypeInfoKind__Macro{
                orgId2MacroTypeInfo.Set(atomInfo.TypeId,newTypeInfo_4123)
            }
            if newTypeInfo_4123.FP.Get_kind() == Ast_TypeInfoKind__Set{
                self.helperInfo.UseSet = true
                
            }
            if newTypeInfo_4123.FP.Get_accessMode() == Ast_AccessMode__Global{
                if _switch4955 := newTypeInfo_4123.FP.Get_kind(); _switch4955 == Ast_TypeInfoKind__IF || _switch4955 == Ast_TypeInfoKind__Class {
                    self.globalScope.FP.AddClass(processInfo, newTypeInfo_4123.FP.Get_rawTxt(), nil, newTypeInfo_4123)
                } else if _switch4955 == Ast_TypeInfoKind__Func {
                    self.globalScope.FP.AddFunc(processInfo, nil, newTypeInfo_4123, Ast_AccessMode__Global, newTypeInfo_4123.FP.Get_staticFlag(), Ast_TypeInfo_isMut(newTypeInfo_4123))
                } else if _switch4955 == Ast_TypeInfoKind__Enum {
                    self.globalScope.FP.AddEnum(processInfo, Ast_AccessMode__Global, newTypeInfo_4123.FP.Get_rawTxt(), nil, newTypeInfo_4123)
                } else if _switch4955 == Ast_TypeInfoKind__Nilable {
                } else {
                    Util_err(Lns_getVM().String_format("%s: not support kind -- %s", []LnsAny{__func__, Ast_TypeInfoKind_getTxt( newTypeInfo_4123.FP.Get_kind())}))
                }
            }
        }
    }
    for _, _atomInfo := range( _typeInfoNormalList.Items ) {
        atomInfo := _atomInfo.(Import__TypeInfoNormalDownCast).ToImport__TypeInfoNormal()
        if atomInfo.Children.Len() > 0{
            var scope *Ast_Scope
            scope = Lns_unwrap( typeId2Scope.Get(atomInfo.TypeId)).(*Ast_Scope)
            for _, _childId := range( atomInfo.Children.Items ) {
                childId := _childId.(Import__IdInfoDownCast).ToImport__IdInfo()
                var typeInfo *Ast_TypeInfo
                
                {
                    _typeInfo := typeId2TypeInfo.Get(childId.Id)
                    if _typeInfo == nil{
                        Util_err(Lns_getVM().String_format("not found childId -- %s, %d, %s(%d)", []LnsAny{orgModulePath, childId.Id, atomInfo.Txt, atomInfo.TypeId}))
                    } else {
                        typeInfo = _typeInfo.(*Ast_TypeInfo)
                    }
                }
                var symbolKind LnsInt
                symbolKind = Ast_SymbolKind__Typ
                var addFlag bool
                addFlag = true
                if _switch5090 := typeInfo.FP.Get_kind(); _switch5090 == Ast_TypeInfoKind__Func {
                    symbolKind = Ast_SymbolKind__Fun
                    
                } else if _switch5090 == Ast_TypeInfoKind__Form || _switch5090 == Ast_TypeInfoKind__FormFunc {
                    symbolKind = Ast_SymbolKind__Typ
                    
                } else if _switch5090 == Ast_TypeInfoKind__Method {
                    symbolKind = Ast_SymbolKind__Mtd
                    
                } else if _switch5090 == Ast_TypeInfoKind__Class || _switch5090 == Ast_TypeInfoKind__Module {
                    symbolKind = Ast_SymbolKind__Typ
                    
                } else if _switch5090 == Ast_TypeInfoKind__Enum {
                    addFlag = false
                    
                }
                if addFlag{
                    scope.FP.Add(processInfo, symbolKind, false, typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Func, typeInfo.FP.GetTxt(nil, nil, nil), nil, typeInfo, typeInfo.FP.Get_accessMode(), typeInfo.FP.Get_staticFlag(), typeInfo.FP.Get_mutMode(), true, false)
                }
            }
        }
    }
    for _typeId, _typeInfo := range( typeId2TypeInfo.Items ) {
        typeId := _typeId.(LnsInt)
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        newId2OldIdMap.Set(typeInfo,typeId)
    }
    var registMember func(classTypeId LnsInt)
    registMember = func(classTypeId LnsInt) {
        __func__ := "@lune.@base.@Import.Import.processImportSub.registMember"
        if Lns_isCondTrue( metaInfo.GetAt( "__dependIdMap" ).(*Lns_luaValue).GetAt(classTypeId)){
            return 
        }
        var classTypeInfo *Ast_TypeInfo
        classTypeInfo = Lns_unwrap( typeId2TypeInfo.Get(classTypeId)).(*Ast_TypeInfo)
        if _switch5553 := (classTypeInfo.FP.Get_kind()); _switch5553 == Ast_TypeInfoKind__Class || _switch5553 == Ast_TypeInfoKind__ExtModule {
            self.transUnitIF.PushClassScope(self.transUnitIF.GetLatestPos(), classTypeInfo)
            {
                __exp := metaInfo.GetAt( "__typeId2ClassInfoMap" ).(*Lns_luaValue).GetAt(classTypeId)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Lns_luaValue)
                    var classInfo LnsAny
                    
                    {
                        _classInfo := Lns_getVM().ExpandLuavalMap(_exp)
                        if _classInfo == nil{
                            self.transUnitIF.Error("illegal val")
                        } else {
                            classInfo = _classInfo
                        }
                    }
                    for _fieldName, _fieldInfo := range( classInfo.(*LnsMap).Items ) {
                        fieldName := _fieldName.(string)
                        fieldInfo := _fieldInfo.(*LnsMap)
                        {
                            _typeId := fieldInfo.Get("typeId")
                            if !Lns_IsNil( _typeId ) {
                                typeId := _typeId
                                var fieldTypeInfo *Ast_TypeInfo
                                fieldTypeInfo = Lns_unwrap( typeId2TypeInfo.Get(Lns_forceCastInt(typeId))).(*Ast_TypeInfo)
                                _ = Import_convExp5324(Lns_2DDD(self.transUnitIF.Get_scope().FP.AddMember(processInfo, fieldName, nil, fieldTypeInfo, Lns_unwrap( Ast_AccessMode__from(Lns_forceCastInt((Lns_unwrap( fieldInfo.Get("accessMode")))))).(LnsInt), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                    Lns_GetEnv().SetStackVal( fieldInfo.Get("staticFlag")) &&
                                    Lns_GetEnv().SetStackVal( true) ||
                                    Lns_GetEnv().SetStackVal( false) ).(bool), Lns_unwrap( Ast_MutMode__from(Lns_forceCastInt((Lns_unwrap( fieldInfo.Get("mutMode")))))).(LnsInt))))
                            } else {
                                self.transUnitIF.Error("not found fieldInfo.typeId")
                            }
                        }
                    }
                } else {
                    self.transUnitIF.Error(Lns_getVM().String_format("not found class -- %s: %d, %s", []LnsAny{orgModulePath, classTypeId, classTypeInfo.FP.GetTxt(nil, nil, nil)}))
                }
            }
        } else if _switch5553 == Ast_TypeInfoKind__Module {
            self.transUnitIF.PushModule(processInfo, true, classTypeInfo.FP.GetTxt(nil, nil, nil), Ast_TypeInfo_isMut(classTypeInfo))
            Log_log(Log_Level__Debug, __func__, 1010, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("push module -- %s, %s, %d, %d, %d", []LnsAny{classTypeInfo.FP.GetTxt(nil, nil, nil), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.transUnitIF.Get_scope().FP.Get_ownerTypeInfo()) && 
                    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.GetFullName(Ast_defaultTypeNameCtrl, self.transUnitIF.Get_scope().FP, false)})/* 1:72 */)) ||
                    Lns_GetEnv().SetStackVal( "nil") ).(string), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.transUnitIF.Get_scope().FP.Get_ownerTypeInfo()) && 
                    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_typeId()})&&
                    Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Ast_IdInfo).Id))) ||
                    Lns_GetEnv().SetStackVal( -1) ).(LnsInt), classTypeInfo.FP.Get_typeId().Id, self.transUnitIF.Get_scope().FP.Get_parent().FP.Get_scopeId()})
            }))
            
        }
        for _, _child := range( classTypeInfo.FP.Get_children().Items ) {
            child := _child.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( child.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( child.FP.Get_kind() == Ast_TypeInfoKind__ExtModule) ||
                Lns_GetEnv().SetStackVal( child.FP.Get_kind() == Ast_TypeInfoKind__Module) ||
                Lns_GetEnv().SetStackVal( child.FP.Get_kind() == Ast_TypeInfoKind__IF) ).(bool){
                var oldId LnsAny
                oldId = newId2OldIdMap.Get(child)
                if Lns_isCondTrue( oldId){
                    registMember(Lns_unwrap( oldId).(LnsInt))
                }
            }
        }
        if _switch5641 := classTypeInfo.FP.Get_kind(); _switch5641 == Ast_TypeInfoKind__Class || _switch5641 == Ast_TypeInfoKind__ExtModule {
            self.transUnitIF.PopClass()
        } else if _switch5641 == Ast_TypeInfoKind__Module {
            self.transUnitIF.PopModule()
        }
    }
    for _, _atomInfo := range( _typeInfoList.Items ) {
        atomInfo := _atomInfo.(Import__TypeInfoDownCast).ToImport__TypeInfo()
        {
            _workInfo := Import__TypeInfoNormalDownCastF(atomInfo.FP)
            if !Lns_IsNil( _workInfo ) {
                workInfo := _workInfo.(*Import__TypeInfoNormal)
                if workInfo.ParentId == Ast_rootTypeId{
                    registMember(atomInfo.TypeId)
                }
            } else {
                {
                    _workInfo := Import__TypeInfoModuleDownCastF(atomInfo.FP)
                    if !Lns_IsNil( _workInfo ) {
                        workInfo := _workInfo.(*Import__TypeInfoModule)
                        if workInfo.ParentId == Ast_rootTypeId{
                            registMember(atomInfo.TypeId)
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
        self.transUnitIF.PushModule(processInfo, true, moduleName, mutable)
    }
    {
        _exp5803 := metaInfo.GetAt( "__varName2InfoMap" ).(*Lns_luaValue)
        _key5803, _val5803 := _exp5803.Get1stFromMap()
        for _key5803 != nil {
            varName := _key5803.(string)
            varInfo := _val5803.(*Lns_luaValue)
            {
                _typeId := varInfo.GetAt("typeId")
                if !Lns_IsNil( _typeId ) {
                    typeId := _typeId
                    self.transUnitIF.Get_scope().FP.AddStaticVar(processInfo, false, true, varName, nil, Lns_unwrap( typeId2TypeInfo.Get(Lns_forceCastInt(typeId))).(*Ast_TypeInfo), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( varInfo.GetAt("mutable")) &&
                        Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
                        Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt))
                } else {
                    self.transUnitIF.Error("illegal varInfo.typeId")
                }
            }
            _key5803, _val5803 = _exp5803.NextFromMap( _key5803 )
        }
    }
    {
        _exp5837 := metaInfo.GetAt( "__macroName2InfoMap" ).(*Lns_luaValue)
        _key5837, _val5837 := _exp5837.Get1stFromMap()
        for _key5837 != nil {
            orgTypeId := _key5837.(LnsInt)
            macroInfoStem := _val5837
            self.macroCtrl.FP.ImportMacro(processInfo, moduleMeta.FP.Get_lnsPath(), Lns_getVM().ExpandLuavalMap(macroInfoStem), Lns_unwrap( orgId2MacroTypeInfo.Get(orgTypeId)).(*Ast_TypeInfo), typeId2TypeInfo)
            _key5837, _val5837 = _exp5837.NextFromMap( _key5837 )
        }
    }
    for range( nameList.Items ) {
        self.transUnitIF.PopModule()
    }
    if depth == 1{
        for _key, _val := range( importParam.ImportedAliasMap.Items ) {
            key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            val := _val.(Ast_AliasTypeInfoDownCast).ToAst_AliasTypeInfo()
            self.importedAliasMap.Set(key,val)
        }
    }
    var moduleProvideInfo *FrontInterface_ModuleProvideInfo
    moduleProvideInfo = NewFrontInterface_ModuleProvideInfo(Lns_unwrap( typeId2TypeInfo.Get(metaInfo.GetAt( "__moduleTypeId" ).(LnsInt))).(*Ast_TypeInfo), Lns_unwrap( Ast_SymbolKind__from(metaInfo.GetAt( "__moduleSymbolKind" ).(LnsInt))).(LnsInt), metaInfo.GetAt( "__moduleMutable" ).(bool))
    var moduleInfo *FrontInterface_ModuleInfo
    moduleInfo = NewFrontInterface_ModuleInfo(orgModulePath, nameList.GetAt(nameList.Len()).(string), newId2OldIdMap, FrontInterface_ModuleId_createIdFromTxt(metaInfo.GetAt( "__buildId" ).(string)), processInfo, moduleProvideInfo, moduleTypeInfo, importParam.ImportedAliasMap)
    return moduleInfo
}

// 1117: decl @lune.@base.@Import.Import.processImport
func (self *Import_Import) ProcessImport(processInfo *Ast_ProcessInfo,modulePath string,depth LnsInt) *FrontInterface_ModuleInfo {
    __func__ := "@lune.@base.@Import.Import.processImport"
    var orgModulePath string
    orgModulePath = modulePath
    modulePath = FrontInterface_getLuaModulePath(modulePath)
    
    Log_log(Log_Level__Info, __func__, 1124, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s -> %s start", []LnsAny{self.moduleType.FP.GetTxt(self.typeNameCtrl, nil, nil), orgModulePath})
    }))
    
    if Lns_op_not(self.importModuleInfo.FP.Add(orgModulePath)){
        self.transUnitIF.Error(Lns_getVM().String_format("recursive import: %s -> %s", []LnsAny{self.importModuleInfo.FP.GetFull(), orgModulePath}))
    }
    {
        _moduleInfo := self.importModuleName2ModuleInfo.Get(modulePath)
        if !Lns_IsNil( _moduleInfo ) {
            moduleInfo := _moduleInfo.(*FrontInterface_ModuleInfo)
            Log_log(Log_Level__Info, __func__, 1136, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("%s already", []LnsAny{orgModulePath})
            }))
            
            self.importModuleInfo.FP.Remove()
            for _key, _val := range( moduleInfo.FP.Get_importedAliasMap().Items ) {
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
        _form6206, _param6206, _prev6206 := Lns_getVM().String_gmatch(modulePath, "[^%./:]+")
        for {
            _work6206 := _form6206.(*Lns_luaValue).Call( Lns_2DDD( _param6206, _prev6206 ) )
            _prev6206 = Lns_getFromMulti(_work6206,0)
            if Lns_IsNil( _prev6206 ) { break }
            txt := _prev6206.(string)
            nameList.Insert(txt)
        }
    }
    var moduleMeta *FrontInterface_ModuleMeta
    
    {
        _moduleMeta := FrontInterface_loadMeta(self.importModuleInfo, orgModulePath)
        if _moduleMeta == nil{
            self.transUnitIF.Error("failed to load meta -- " + orgModulePath)
        } else {
            moduleMeta = _moduleMeta.(*FrontInterface_ModuleMeta)
        }
    }
    var moduleInfo *FrontInterface_ModuleInfo
    moduleInfo = self.FP.processImportSub(processInfo, moduleMeta, orgModulePath, modulePath, nameList, depth)
    self.importModule2ModuleInfo.Set(moduleInfo.FP.Get_moduleTypeInfo(),moduleInfo)
    self.importModuleName2ModuleInfo.Set(modulePath,moduleInfo)
    self.importModuleInfo.FP.Remove()
    Log_log(Log_Level__Info, __func__, 1169, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s complete", []LnsAny{orgModulePath})
    }))
    
    return moduleInfo
}


// declaration Class -- ImportParam
type Import_ImportParamMtd interface {
    GetTypeInfo(arg1 LnsInt)(LnsAny, LnsAny)
    GetTypeInfoFrom(arg1 *Import__IdInfo)(LnsAny, LnsAny)
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
func NewImport_ImportParam(arg1 *Types_Position, arg2 *TransUnitIF_Modifier, arg3 *Ast_ProcessInfo, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsSet, arg8 *Lns_luaValue, arg9 *Ast_Scope, arg10 *Ast_TypeInfo, arg11 LnsInt, arg12 *LnsMap) *Import_ImportParam {
    obj := &Import_ImportParam{}
    obj.FP = obj
    obj.InitImport_ImportParam(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
    return obj
}
func (self *Import_ImportParam) InitImport_ImportParam(arg1 *Types_Position, arg2 *TransUnitIF_Modifier, arg3 *Ast_ProcessInfo, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsSet, arg8 *Lns_luaValue, arg9 *Ast_Scope, arg10 *Ast_TypeInfo, arg11 LnsInt, arg12 *LnsMap) {
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
}
// 157: decl @lune.@base.@Import.ImportParam.getTypeInfo
func (self *Import_ImportParam) GetTypeInfo(typeId LnsInt)(LnsAny, LnsAny) {
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
            typeInfo,mess = atom.FP.CreateTypeInfoCache(self)
            if typeInfo != nil{
                typeInfo_3676 := typeInfo.(*Ast_TypeInfo)
                self.TypeId2TypeInfo.Set(typeId,typeInfo_3676)
            }
            return typeInfo, mess
        }
    }
    return nil, nil
}

// 176: decl @lune.@base.@Import.ImportParam.getTypeInfoFrom
func (self *Import_ImportParam) GetTypeInfoFrom(typeId *Import__IdInfo)(LnsAny, LnsAny) {
    return self.FP.GetTypeInfo(typeId.Id)
}


// declaration Class -- _TypeInfo
type Import__TypeInfoMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
// 130: DeclConstr
func (self *Import__TypeInfo) InitImport__TypeInfo() {
    self.TypeId = Ast_rootTypeId
    
    self.Skind = Ast_SerializeKind__Normal
    
}


// 145: decl @lune.@base.@Import._TypeInfo.createTypeInfoCache
func (self *Import__TypeInfo) CreateTypeInfoCache(param *Import_ImportParam)(LnsAny, LnsAny) {
    {
        _typeInfo := param.TypeId2TypeInfo.Get(self.TypeId)
        if !Lns_IsNil( _typeInfo ) {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo, nil
        }
    }
    var typeInfo LnsAny
    var mess LnsAny
    typeInfo,mess = self.FP.CreateTypeInfo(param)
    if typeInfo != nil{
        typeInfo_3661 := typeInfo.(*Ast_TypeInfo)
        param.TypeId2TypeInfo.Set(self.TypeId,typeInfo_3661)
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
func NewImport__IdInfo(arg1 LnsInt, arg2 LnsInt) *Import__IdInfo {
    obj := &Import__IdInfo{}
    obj.FP = obj
    obj.InitImport__IdInfo(arg1, arg2)
    return obj
}
func (self *Import__IdInfo) InitImport__IdInfo(arg1 LnsInt, arg2 LnsInt) {
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
func Import__IdInfo__fromMap_1081_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__IdInfo_FromMap( arg1, paramList )
}
func Import__IdInfo__fromStem_1083_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
}
type Import__TypeInfoNilable struct {
    Import__TypeInfo
    OrgTypeId LnsInt
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
func NewImport__TypeInfoNilable(arg1 LnsInt) *Import__TypeInfoNilable {
    obj := &Import__TypeInfoNilable{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoNilable(arg1)
    return obj
}
func (self *Import__TypeInfoNilable) InitImport__TypeInfoNilable(arg1 LnsInt) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
    self.OrgTypeId = arg1
}
func (self *Import__TypeInfoNilable) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["orgTypeId"] = Lns_ToCollection( self.OrgTypeId )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoNilable) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoNilable__fromMap_1101_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoNilable_FromMap( arg1, paramList )
}
func Import__TypeInfoNilable__fromStem_1103_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["orgTypeId"], false, nil); !ok {
       return false,nil,"orgTypeId:" + mess.(string)
    } else {
       newObj.OrgTypeId = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 184: decl @lune.@base.@Import._TypeInfoNilable.createTypeInfo
func (self *Import__TypeInfoNilable) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    var orgTypeInfo *Ast_TypeInfo
    
    {
        _orgTypeInfo := Import_convExp567(Lns_2DDD(param.FP.GetTypeInfo(self.OrgTypeId)))
        if _orgTypeInfo == nil{
            Util_err(Lns_getVM().String_format("failed to createTypeInfo -- self.orgTypeId = %d", []LnsAny{self.OrgTypeId}))
        } else {
            orgTypeInfo = _orgTypeInfo.(*Ast_TypeInfo)
        }
    }
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = orgTypeInfo.FP.Get_nilableTypeInfo()
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoAlias
type Import__TypeInfoAliasMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoAlias(arg1 LnsInt, arg2 string, arg3 *Import__IdInfo) *Import__TypeInfoAlias {
    obj := &Import__TypeInfoAlias{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoAlias(arg1, arg2, arg3)
    return obj
}
func (self *Import__TypeInfoAlias) InitImport__TypeInfoAlias(arg1 LnsInt, arg2 string, arg3 *Import__IdInfo) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
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
func Import__TypeInfoAlias__fromMap_1118_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlias_FromMap( arg1, paramList )
}
func Import__TypeInfoAlias__fromStem_1120_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 200: decl @lune.@base.@Import._TypeInfoAlias.createTypeInfo
func (self *Import__TypeInfoAlias) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@Import._TypeInfoAlias.createTypeInfo"
    var srcTypeInfo *Ast_TypeInfo
    srcTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(self.srcTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_AliasTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateAlias(param.ProcessInfo, self.rawTxt, true, Ast_AccessMode__Pub, param.ModuleTypeInfo, srcTypeInfo)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    
    {
        __ := Import_convExp696(Lns_2DDD(param.FP.GetTypeInfo(self.ParentId)))
        if __ == nil{
            return nil, Lns_getVM().String_format("%s: not found parentInfo %d %s", []LnsAny{__func__, self.ParentId, self.rawTxt})
        } else {
            _ = __.(*Ast_TypeInfo)
        }
    }
    var parentScope *Ast_Scope
    
    {
        _parentScope := param.TypeId2Scope.Get(self.ParentId)
        if _parentScope == nil{
            return nil, Lns_getVM().String_format("%s: not found parentScope %s %s", []LnsAny{__func__, self.ParentId, self.rawTxt})
        } else {
            parentScope = _parentScope.(*Ast_Scope)
        }
    }
    parentScope.FP.AddAliasForType(param.ProcessInfo, self.rawTxt, nil, &newTypeInfo.Ast_TypeInfo)
    param.ImportedAliasMap.Set(srcTypeInfo,newTypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoDDD
type Import__TypeInfoDDDMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoDDD(arg1 LnsInt, arg2 *Import__IdInfo, arg3 bool) *Import__TypeInfoDDD {
    obj := &Import__TypeInfoDDD{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoDDD(arg1, arg2, arg3)
    return obj
}
func (self *Import__TypeInfoDDD) InitImport__TypeInfoDDD(arg1 LnsInt, arg2 *Import__IdInfo, arg3 bool) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
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
func Import__TypeInfoDDD__fromMap_1135_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoDDD_FromMap( arg1, paramList )
}
func Import__TypeInfoDDD__fromStem_1137_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 231: decl @lune.@base.@Import._TypeInfoDDD.createTypeInfo
func (self *Import__TypeInfoDDD) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    var itemTypeInfo *Ast_TypeInfo
    itemTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(self.ItemTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_DDDTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateDDD(itemTypeInfo, true, self.ExtTypeFlag)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoAlternate
type Import__TypeInfoAlternateMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoAlternate(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *Import__IdInfo, arg5 *LnsList, arg6 bool, arg7 LnsInt) *Import__TypeInfoAlternate {
    obj := &Import__TypeInfoAlternate{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoAlternate(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Import__TypeInfoAlternate) InitImport__TypeInfoAlternate(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *Import__IdInfo, arg5 *LnsList, arg6 bool, arg7 LnsInt) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
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
func Import__TypeInfoAlternate__fromMap_1159_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlternate_FromMap( arg1, paramList )
}
func Import__TypeInfoAlternate__fromStem_1161_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 252: decl @lune.@base.@Import._TypeInfoAlternate.createTypeInfo
func (self *Import__TypeInfoAlternate) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    var baseInfo *Ast_TypeInfo
    baseInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(self.BaseId))).(*Ast_TypeInfo)
    var interfaceList *LnsList
    interfaceList = NewLnsList([]LnsAny{})
    for _, _ifTypeId := range( self.IfList.Items ) {
        ifTypeId := _ifTypeId.(Import__IdInfoDownCast).ToImport__IdInfo()
        interfaceList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(ifTypeId))).(*Ast_TypeInfo)))
    }
    var newTypeInfo *Ast_AlternateTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateAlternate(self.BelongClassFlag, self.AltIndex, self.Txt, self.AccessMode, param.ModuleTypeInfo, baseInfo, interfaceList)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoGeneric
type Import__TypeInfoGenericMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoGeneric(arg1 *Import__IdInfo, arg2 *LnsList) *Import__TypeInfoGeneric {
    obj := &Import__TypeInfoGeneric{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoGeneric(arg1, arg2)
    return obj
}
func (self *Import__TypeInfoGeneric) InitImport__TypeInfoGeneric(arg1 *Import__IdInfo, arg2 *LnsList) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
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
func Import__TypeInfoGeneric__fromMap_1183_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoGeneric_FromMap( arg1, paramList )
}
func Import__TypeInfoGeneric__fromStem_1185_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 272: decl @lune.@base.@Import._TypeInfoGeneric.createTypeInfo
func (self *Import__TypeInfoGeneric) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    var genSrcTypeInfo *Ast_TypeInfo
    genSrcTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(self.GenSrcTypeId))).(*Ast_TypeInfo)
    var genTypeList *LnsList
    genTypeList = NewLnsList([]LnsAny{})
    for _, _typeId := range( self.GenTypeList.Items ) {
        typeId := _typeId.(Import__IdInfoDownCast).ToImport__IdInfo()
        genTypeList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(typeId))).(*Ast_TypeInfo)))
    }
    var newTypeInfo *Ast_GenericTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateGeneric(genSrcTypeInfo, genTypeList, param.ModuleTypeInfo)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    param.TypeId2Scope.Set(self.TypeId,Ast_getScope(&newTypeInfo.Ast_TypeInfo))
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoBox
type Import__TypeInfoBoxMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoBox(arg1 LnsInt, arg2 LnsInt) *Import__TypeInfoBox {
    obj := &Import__TypeInfoBox{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoBox(arg1, arg2)
    return obj
}
func (self *Import__TypeInfoBox) InitImport__TypeInfoBox(arg1 LnsInt, arg2 LnsInt) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
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
func Import__TypeInfoBox__fromMap_1200_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoBox_FromMap( arg1, paramList )
}
func Import__TypeInfoBox__fromStem_1202_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 291: decl @lune.@base.@Import._TypeInfoBox.createTypeInfo
func (self *Import__TypeInfoBox) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    var boxingType *Ast_TypeInfo
    boxingType = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.BoxingType))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateBox(self.AccessMode, boxingType)
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoExt
type Import__TypeInfoExtMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoExt(arg1 *Import__IdInfo) *Import__TypeInfoExt {
    obj := &Import__TypeInfoExt{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoExt(arg1)
    return obj
}
func (self *Import__TypeInfoExt) InitImport__TypeInfoExt(arg1 *Import__IdInfo) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
    self.ExtedTypeId = arg1
}
func (self *Import__TypeInfoExt) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["extedTypeId"] = Lns_ToCollection( self.ExtedTypeId )
    return self.Import__TypeInfo.ToMapSetup( obj )
}
func (self *Import__TypeInfoExt) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Import__TypeInfoExt__fromMap_1217_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoExt_FromMap( arg1, paramList )
}
func Import__TypeInfoExt__fromStem_1219_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 303: decl @lune.@base.@Import._TypeInfoExt.createTypeInfo
func (self *Import__TypeInfoExt) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    var extedType *Ast_TypeInfo
    extedType = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(self.ExtedTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_TypeInfo
    switch _exp1227 := param.ProcessInfo.FP.CreateLuaval(extedType, true).(type) {
    case *Ast_LuavalResult__OK:
    extType := _exp1227.Val1
        newTypeInfo = extType
        
    case *Ast_LuavalResult__Err:
    mess := _exp1227.Val1
        Util_err(mess)
    }
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoModifier
type Import__TypeInfoModifierMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoModifier(arg1 *Import__IdInfo, arg2 LnsInt) *Import__TypeInfoModifier {
    obj := &Import__TypeInfoModifier{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoModifier(arg1, arg2)
    return obj
}
func (self *Import__TypeInfoModifier) InitImport__TypeInfoModifier(arg1 *Import__IdInfo, arg2 LnsInt) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
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
func Import__TypeInfoModifier__fromMap_1234_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoModifier_FromMap( arg1, paramList )
}
func Import__TypeInfoModifier__fromStem_1236_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 325: decl @lune.@base.@Import._TypeInfoModifier.createTypeInfo
func (self *Import__TypeInfoModifier) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    var srcTypeInfo *Ast_TypeInfo
    
    {
        _srcTypeInfo := Import_convExp1294(Lns_2DDD(param.FP.GetTypeInfoFrom(self.SrcTypeId)))
        if _srcTypeInfo == nil{
            return nil, Lns_getVM().String_format("not found srcType -- %d", []LnsAny{self.SrcTypeId.Id})
        } else {
            srcTypeInfo = _srcTypeInfo.(*Ast_TypeInfo)
        }
    }
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = param.Modifier.FP.CreateModifier(srcTypeInfo, self.MutMode)
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoModule
type Import__TypeInfoModuleMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoModule(arg1 LnsInt, arg2 string) *Import__TypeInfoModule {
    obj := &Import__TypeInfoModule{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoModule(arg1, arg2)
    return obj
}
func (self *Import__TypeInfoModule) InitImport__TypeInfoModule(arg1 LnsInt, arg2 string) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
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
func Import__TypeInfoModule__fromMap_1253_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoModule_FromMap( arg1, paramList )
}
func Import__TypeInfoModule__fromStem_1255_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 341: decl @lune.@base.@Import._TypeInfoModule.createTypeInfo
func (self *Import__TypeInfoModule) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@Import._TypeInfoModule.createTypeInfo"
    var parentInfo *Ast_TypeInfo
    parentInfo = Ast_headTypeInfo
    if self.ParentId != Ast_rootTypeId{
        var workTypeInfo *Ast_TypeInfo
        
        {
            _workTypeInfo := Import_convExp1392(Lns_2DDD(param.FP.GetTypeInfo(self.ParentId)))
            if _workTypeInfo == nil{
                Util_err(Lns_getVM().String_format("not found parentInfo %d %s", []LnsAny{self.ParentId, self.Txt}))
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
            return nil, Lns_getVM().String_format("%s: not found parentScope %s %s", []LnsAny{__func__, self.ParentId, self.Txt})
        } else {
            parentScope = _parentScope.(*Ast_Scope)
        }
    }
    var newTypeInfo LnsAny
    newTypeInfo = parentScope.FP.GetTypeInfoChild(self.Txt)
    {
        __exp := newTypeInfo
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            param.TypeId2Scope.Set(self.TypeId,Ast_getScope(_exp))
            if Lns_op_not(_exp.FP.Get_scope()){
                return nil, Lns_getVM().String_format("not found scope %s %d %s %s %s", []LnsAny{parentScope, self.ParentId, self.TypeId, self.Txt, _exp.FP.GetTxt(nil, nil, nil)})
            }
            param.TypeId2TypeInfo.Set(self.TypeId,_exp)
        } else {
            var scope *Ast_Scope
            scope = NewAst_Scope(param.ProcessInfo, parentScope, true, nil, nil)
            var mutable bool
            mutable = false
            if self.TypeId == param.MetaInfo.GetAt( "__moduleTypeId" ).(LnsInt){
                mutable = param.MetaInfo.GetAt( "__moduleMutable" ).(bool)
                
            }
            var workTypeInfo *Ast_TypeInfo
            workTypeInfo = param.ProcessInfo.FP.CreateModule(scope, parentInfo, true, self.Txt, mutable)
            newTypeInfo = workTypeInfo
            
            param.TypeId2Scope.Set(self.TypeId,scope)
            param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
            parentScope.FP.AddClass(param.ProcessInfo, self.Txt, nil, workTypeInfo)
            Log_log(Log_Level__Info, __func__, 383, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("new module -- %s, %s, %d, %d, %d", []LnsAny{self.Txt, workTypeInfo.FP.GetFullName(Ast_defaultTypeNameCtrl, parentScope.FP, false), self.TypeId, workTypeInfo.FP.Get_typeId().Id, parentScope.FP.Get_scopeId()})
            }))
            
        }
    }
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoNormal
type Import__TypeInfoNormalMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoNormal(arg1 LnsInt, arg2 bool, arg3 *Import__IdInfo, arg4 string, arg5 bool, arg6 LnsInt, arg7 LnsInt, arg8 LnsInt, arg9 *LnsList, arg10 *LnsList, arg11 *LnsList, arg12 *LnsList, arg13 *LnsList, arg14 LnsAny, arg15 LnsAny) *Import__TypeInfoNormal {
    obj := &Import__TypeInfoNormal{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoNormal(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15)
    return obj
}
func (self *Import__TypeInfoNormal) InitImport__TypeInfoNormal(arg1 LnsInt, arg2 bool, arg3 *Import__IdInfo, arg4 string, arg5 bool, arg6 LnsInt, arg7 LnsInt, arg8 LnsInt, arg9 *LnsList, arg10 *LnsList, arg11 *LnsList, arg12 *LnsList, arg13 *LnsList, arg14 LnsAny, arg15 LnsAny) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
    self.ParentId = arg1
    self.AbstractFlag = arg2
    self.BaseId = arg3
    self.Txt = arg4
    self.StaticFlag = arg5
    self.AccessMode = arg6
    self.Kind = arg7
    self.MutMode = arg8
    self.IfList = arg9
    self.ItemTypeId = arg10
    self.ArgTypeId = arg11
    self.RetTypeId = arg12
    self.Children = arg13
    self.ModuleLang = arg14
    self.RequirePath = arg15
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
func Import__TypeInfoNormal__fromMap_1319_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoNormal_FromMap( arg1, paramList )
}
func Import__TypeInfoNormal__fromStem_1321_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 411: decl @lune.@base.@Import._TypeInfoNormal.createTypeInfo
func (self *Import__TypeInfoNormal) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@Import._TypeInfoNormal.createTypeInfo"
    var newTypeInfo LnsAny
    newTypeInfo = nil
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.ParentId != Ast_rootTypeId) ||
        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_getBuiltInTypeIdMap().Get(self.TypeId))) ||
        Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__List) ||
        Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Array) ||
        Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Map) ||
        Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Set) ).(bool){
        var parentInfo *Ast_TypeInfo
        parentInfo = Ast_headTypeInfo
        if self.ParentId != Ast_rootTypeId{
            var workTypeInfo *Ast_TypeInfo
            
            {
                _workTypeInfo := Import_convExp1907(Lns_2DDD(param.FP.GetTypeInfo(self.ParentId)))
                if _workTypeInfo == nil{
                    return nil, Lns_getVM().String_format("not found parentInfo %d %s", []LnsAny{self.ParentId, self.Txt})
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
            itemTypeInfo.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(typeId))).(*Ast_TypeInfo)))
        }
        var argTypeInfo *LnsList
        argTypeInfo = NewLnsList([]LnsAny{})
        for _index, _typeId := range( self.ArgTypeId.Items ) {
            index := _index + 1
            typeId := _typeId.(Import__IdInfoDownCast).ToImport__IdInfo()
            var argType LnsAny
            var mess LnsAny
            argType,mess = param.FP.GetTypeInfoFrom(typeId)
            if argType != nil{
                argType_3896 := argType.(*Ast_TypeInfo)
                argTypeInfo.Insert(Ast_TypeInfo2Stem(argType_3896))
            } else {
                var errmess string
                errmess = Lns_getVM().String_format("not found arg (index:%d) -- %s.%s, %d, %d. %s", []LnsAny{index, parentInfo.FP.GetTxt(nil, nil, nil), self.Txt, typeId.Id, self.ArgTypeId.Len(), mess})
                return nil, errmess
            }
        }
        var retTypeInfo *LnsList
        retTypeInfo = NewLnsList([]LnsAny{})
        for _, _typeId := range( self.RetTypeId.Items ) {
            typeId := _typeId.(Import__IdInfoDownCast).ToImport__IdInfo()
            retTypeInfo.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(typeId))).(*Ast_TypeInfo)))
        }
        var baseInfo *Ast_TypeInfo
        baseInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(self.BaseId))).(*Ast_TypeInfo)
        var interfaceList *LnsList
        interfaceList = NewLnsList([]LnsAny{})
        for _, _ifTypeId := range( self.IfList.Items ) {
            ifTypeId := _ifTypeId.(Import__IdInfoDownCast).ToImport__IdInfo()
            interfaceList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(ifTypeId))).(*Ast_TypeInfo)))
        }
        var parentScope *Ast_Scope
        
        {
            _parentScope := param.TypeId2Scope.Get(self.ParentId)
            if _parentScope == nil{
                return nil, Lns_getVM().String_format("%s: not found parentScope %s %s", []LnsAny{__func__, self.ParentId, self.Txt})
            } else {
                parentScope = _parentScope.(*Ast_Scope)
            }
        }
        if self.Txt != ""{
            newTypeInfo = parentScope.FP.GetTypeInfoChild(self.Txt)
            
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( newTypeInfo) &&
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__ExtModule) ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__IF) ).(bool))) )){
            {
                __exp := newTypeInfo
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Ast_TypeInfo)
                    param.TypeId2Scope.Set(self.TypeId,Ast_getScope(_exp))
                    if Lns_op_not(_exp.FP.Get_scope()){
                        Util_err(Lns_getVM().String_format("not found scope %s %s %s %s %s", []LnsAny{parentScope, self.ParentId, self.TypeId, self.Txt, _exp.FP.GetTxt(nil, nil, nil)}))
                    }
                    param.TypeId2TypeInfo.Set(self.TypeId,_exp)
                }
            }
        } else { 
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__IF) ).(bool){
                Log_log(Log_Level__Debug, __func__, 490, Log_CreateMessage(func() string {
                    return Lns_getVM().String_format("new type -- %d, %s -- %s, %d", []LnsAny{self.ParentId, self.Txt, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentScope.FP.Get_ownerTypeInfo()) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.GetFullName(Ast_defaultTypeNameCtrl, parentScope.FP, false)})/* 1:65 */)) ||
                        Lns_GetEnv().SetStackVal( "nil") ).(string), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentScope.FP.Get_ownerTypeInfo()) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_typeId()})&&
                        Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Ast_IdInfo).Id))) ||
                        Lns_GetEnv().SetStackVal( -1) ).(LnsInt)})
                }))
                
                var baseScope *Ast_Scope
                baseScope = Lns_unwrap( Ast_getScope(baseInfo)).(*Ast_Scope)
                var ifScopeList *LnsList
                ifScopeList = NewLnsList([]LnsAny{})
                for _, _ifType := range( interfaceList.Items ) {
                    ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    ifScopeList.Insert(Ast_Scope2Stem(Lns_unwrap( ifType.FP.Get_scope()).(*Ast_Scope)))
                }
                var scope *Ast_Scope
                scope = NewAst_Scope(param.ProcessInfo, parentScope, true, baseScope, ifScopeList)
                var altTypeList *LnsList
                altTypeList = NewLnsList([]LnsAny{})
                for _, _itemType := range( itemTypeInfo.Items ) {
                    itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    altTypeList.Insert(Ast_AlternateTypeInfo2Stem(Lns_unwrap( (Ast_AlternateTypeInfoDownCastF(itemType.FP))).(*Ast_AlternateTypeInfo)))
                }
                var workTypeInfo *Ast_TypeInfo
                workTypeInfo = param.ProcessInfo.FP.CreateClass(self.Kind == Ast_TypeInfoKind__Class, self.AbstractFlag, scope, baseInfo, interfaceList, altTypeList, parentInfo, true, Ast_AccessMode__Pub, self.Txt)
                parentScope.FP.AddClassLazy(param.ProcessInfo, self.Txt, nil, workTypeInfo, param.LazyModuleSet.Has(self.TypeId))
                newTypeInfo = workTypeInfo
                
                param.TypeId2Scope.Set(self.TypeId,scope)
                param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
            } else if self.Kind == Ast_TypeInfoKind__ExtModule{
                Log_log(Log_Level__Debug, __func__, 527, Log_CreateMessage(func() string {
                    return Lns_getVM().String_format("new type -- %d, %s -- %s, %d", []LnsAny{self.ParentId, self.Txt, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentScope.FP.Get_ownerTypeInfo()) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.GetFullName(Ast_defaultTypeNameCtrl, parentScope.FP, false)})/* 1:65 */)) ||
                        Lns_GetEnv().SetStackVal( "nil") ).(string), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentScope.FP.Get_ownerTypeInfo()) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_typeId()})&&
                        Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Ast_IdInfo).Id))) ||
                        Lns_GetEnv().SetStackVal( -1) ).(LnsInt)})
                }))
                
                var scope *Ast_Scope
                scope = NewAst_Scope(param.ProcessInfo, parentScope, true, nil, NewLnsList([]LnsAny{}))
                var workTypeInfo *Ast_TypeInfo
                workTypeInfo = param.ProcessInfo.FP.CreateExtModule(scope, parentInfo, true, Ast_AccessMode__Pub, self.Txt, Lns_unwrap( self.ModuleLang).(LnsInt), Lns_unwrap( self.RequirePath).(string))
                parentScope.FP.AddExtModule(param.ProcessInfo, self.Txt, nil, workTypeInfo, param.LazyModuleSet.Has(self.TypeId), Lns_unwrap( self.ModuleLang).(LnsInt))
                newTypeInfo = workTypeInfo
                
                param.TypeId2Scope.Set(self.TypeId,scope)
                param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
            } else { 
                var scope LnsAny
                scope = nil
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Func) ||
                    Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Method) ).(bool){
                    scope = NewAst_Scope(param.ProcessInfo, parentScope, false, nil, nil)
                    
                }
                var typeInfoKind LnsInt
                typeInfoKind = self.Kind
                var accessMode LnsInt
                accessMode = self.AccessMode
                var workTypeInfo *Ast_TypeInfo
                workTypeInfo = Ast_NormalTypeInfo_create(param.ProcessInfo, accessMode, self.AbstractFlag, scope, baseInfo, parentInfo, self.StaticFlag, typeInfoKind, self.Txt, itemTypeInfo, argTypeInfo, retTypeInfo, self.MutMode)
                newTypeInfo = workTypeInfo
                
                param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
                if _switch3070 := self.Kind; _switch3070 == Ast_TypeInfoKind__Func || _switch3070 == Ast_TypeInfoKind__Method || _switch3070 == Ast_TypeInfoKind__Macro || _switch3070 == Ast_TypeInfoKind__Form || _switch3070 == Ast_TypeInfoKind__FormFunc {
                    var symbolKind LnsInt
                    symbolKind = Ast_SymbolKind__Fun
                    if _switch3008 := self.Kind; _switch3008 == Ast_TypeInfoKind__Method {
                        symbolKind = Ast_SymbolKind__Mtd
                        
                    } else if _switch3008 == Ast_TypeInfoKind__Macro {
                        symbolKind = Ast_SymbolKind__Mac
                        
                    } else if _switch3008 == Ast_TypeInfoKind__Form || _switch3008 == Ast_TypeInfoKind__FormFunc {
                        symbolKind = Ast_SymbolKind__Typ
                        
                    }
                    var workParentScope *Ast_Scope
                    workParentScope = Lns_unwrap( param.TypeId2Scope.Get(self.ParentId)).(*Ast_Scope)
                    workParentScope.FP.Add(param.ProcessInfo, symbolKind, false, self.Kind == Ast_TypeInfoKind__Func, self.Txt, nil, workTypeInfo, accessMode, self.StaticFlag, Ast_MutMode__IMut, true, false)
                    param.TypeId2Scope.Set(self.TypeId,scope)
                }
            }
        }
    } else { 
        newTypeInfo = param.Scope.FP.GetTypeInfo(self.Txt, param.Scope, false, param.ScopeAccess)
        
        if Lns_op_not(newTypeInfo){
            for _key, _val := range( self.FP.ToMap().Items ) {
                key := _key.(string)
                val := _val
                Util_errorLog(Lns_getVM().String_format("error: illegal self %s:%s", []LnsAny{key, val}))
            }
        }
        param.TypeId2TypeInfo.Set(self.TypeId,Lns_unwrap( newTypeInfo).(*Ast_TypeInfo))
    }
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoEnum
type Import__TypeInfoEnumMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoEnum(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 LnsInt, arg5 *LnsMap) *Import__TypeInfoEnum {
    obj := &Import__TypeInfoEnum{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoEnum(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Import__TypeInfoEnum) InitImport__TypeInfoEnum(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 LnsInt, arg5 *LnsMap) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
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
func Import__TypeInfoEnum__fromMap_1341_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoEnum_FromMap( arg1, paramList )
}
func Import__TypeInfoEnum__fromStem_1343_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 619: decl @lune.@base.@Import._TypeInfoEnum.createTypeInfo
func (self *Import__TypeInfoEnum) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    var accessMode LnsInt
    accessMode = Lns_unwrap( Ast_AccessMode__from(LnsInt(self.AccessMode))).(LnsInt)
    var parentInfo *Ast_TypeInfo
    parentInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.ParentId))).(*Ast_TypeInfo)
    var parentScope *Ast_Scope
    parentScope = Lns_unwrap( Ast_getScope(parentInfo)).(*Ast_Scope)
    var scope *Ast_Scope
    scope = NewAst_Scope(param.ProcessInfo, parentScope, true, nil, nil)
    param.TypeId2Scope.Set(self.TypeId,scope)
    var valTypeInfo *Ast_TypeInfo
    valTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.ValTypeId))).(*Ast_TypeInfo)
    var enumTypeInfo *Ast_EnumTypeInfo
    enumTypeInfo = param.ProcessInfo.FP.CreateEnum(scope, parentInfo, true, accessMode, self.Txt, valTypeInfo)
    var newTypeInfo *Ast_EnumTypeInfo
    newTypeInfo = enumTypeInfo
    param.TypeId2TypeInfo.Set(self.TypeId,&enumTypeInfo.Ast_TypeInfo)
    var getEnumLiteral func(val LnsAny) LnsAny
    getEnumLiteral = func(val LnsAny) LnsAny {
        if _switch3367 := valTypeInfo; _switch3367 == Ast_builtinTypeInt {
            return &Ast_EnumLiteral__Int{Lns_forceCastInt(val)}
        } else if _switch3367 == Ast_builtinTypeReal {
            return &Ast_EnumLiteral__Real{Lns_forceCastReal(val)}
        } else if _switch3367 == Ast_builtinTypeString {
            return &Ast_EnumLiteral__Str{val.(string)}
        }
        return nil
    }
    for _valName, _valData := range( self.EnumValList.Items ) {
        valName := _valName.(string)
        valData := _valData
        var val LnsAny
        
        {
            _val := getEnumLiteral(valData)
            if _val == nil{
                return nil, Lns_getVM().String_format("unknown enum val type -- %s", []LnsAny{valTypeInfo.FP.GetTxt(nil, nil, nil)})
            } else {
                val = _val
            }
        }
        var evalValSym *Ast_SymbolInfo
        evalValSym = Lns_unwrap( Lns_car(scope.FP.AddEnumVal(param.ProcessInfo, valName, nil, &enumTypeInfo.Ast_TypeInfo))).(*Ast_SymbolInfo)
        enumTypeInfo.FP.AddEnumValInfo(NewAst_EnumValInfo(valName, val, evalValSym))
    }
    parentScope.FP.AddEnum(param.ProcessInfo, accessMode, self.Txt, nil, &enumTypeInfo.Ast_TypeInfo)
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
func NewImport__TypeInfoAlgeVal(arg1 string, arg2 *LnsList) *Import__TypeInfoAlgeVal {
    obj := &Import__TypeInfoAlgeVal{}
    obj.FP = obj
    obj.InitImport__TypeInfoAlgeVal(arg1, arg2)
    return obj
}
func (self *Import__TypeInfoAlgeVal) InitImport__TypeInfoAlgeVal(arg1 string, arg2 *LnsList) {
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
func Import__TypeInfoAlgeVal__fromMap_1358_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlgeVal_FromMap( arg1, paramList )
}
func Import__TypeInfoAlgeVal__fromStem_1360_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    CreateTypeInfo(arg1 *Import_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *Import_ImportParam)(LnsAny, LnsAny)
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
func NewImport__TypeInfoAlge(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *LnsList) *Import__TypeInfoAlge {
    obj := &Import__TypeInfoAlge{}
    obj.FP = obj
    obj.Import__TypeInfo.FP = obj
    obj.InitImport__TypeInfoAlge(arg1, arg2, arg3, arg4)
    return obj
}
func (self *Import__TypeInfoAlge) InitImport__TypeInfoAlge(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *LnsList) {
    self.Import__TypeInfo.InitImport__TypeInfo( )
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
func Import__TypeInfoAlge__fromMap_1383_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Import__TypeInfoAlge_FromMap( arg1, paramList )
}
func Import__TypeInfoAlge__fromStem_1385_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 673: decl @lune.@base.@Import._TypeInfoAlge.createTypeInfo
func (self *Import__TypeInfoAlge) CreateTypeInfo(param *Import_ImportParam)(LnsAny, LnsAny) {
    var accessMode LnsInt
    accessMode = Lns_unwrap( Ast_AccessMode__from(LnsInt(self.AccessMode))).(LnsInt)
    var parentInfo *Ast_TypeInfo
    parentInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.ParentId))).(*Ast_TypeInfo)
    var parentScope *Ast_Scope
    parentScope = Lns_unwrap( Ast_getScope(parentInfo)).(*Ast_Scope)
    var scope *Ast_Scope
    scope = NewAst_Scope(param.ProcessInfo, parentScope, true, nil, nil)
    param.TypeId2Scope.Set(self.TypeId,scope)
    var algeTypeInfo *Ast_AlgeTypeInfo
    algeTypeInfo = param.ProcessInfo.FP.CreateAlge(scope, parentInfo, true, accessMode, self.Txt)
    var newTypeInfo *Ast_AlgeTypeInfo
    newTypeInfo = algeTypeInfo
    param.TypeId2TypeInfo.Set(self.TypeId,&algeTypeInfo.Ast_TypeInfo)
    for _, _valInfo := range( self.AlgeValList.Items ) {
        valInfo := _valInfo.(Import__TypeInfoAlgeValDownCast).ToImport__TypeInfoAlgeVal()
        var typeInfoList *LnsList
        typeInfoList = NewLnsList([]LnsAny{})
        for _, _orgTypeId := range( valInfo.TypeList.Items ) {
            orgTypeId := _orgTypeId.(Import__IdInfoDownCast).ToImport__IdInfo()
            typeInfoList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfoFrom(orgTypeId))).(*Ast_TypeInfo)))
        }
        var algeValSym LnsAny
        algeValSym = Import_convExp3679(Lns_2DDD(scope.FP.AddAlgeVal(param.ProcessInfo, valInfo.Name, nil, &algeTypeInfo.Ast_TypeInfo)))
        var algeVal *Ast_AlgeValInfo
        algeVal = NewAst_AlgeValInfo(valInfo.Name, typeInfoList, algeTypeInfo, Lns_unwrap( algeValSym).(*Ast_SymbolInfo))
        algeTypeInfo.FP.AddValInfo(algeVal)
    }
    parentScope.FP.AddAlge(param.ProcessInfo, accessMode, self.Txt, nil, &algeTypeInfo.Ast_TypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- DependModuleInfo
type Import_DependModuleInfoMtd interface {
    GetTypeInfo(arg1 LnsInt) *Ast_TypeInfo
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
func NewImport_DependModuleInfo(arg1 LnsInt, arg2 *LnsMap) *Import_DependModuleInfo {
    obj := &Import_DependModuleInfo{}
    obj.FP = obj
    obj.InitImport_DependModuleInfo(arg1, arg2)
    return obj
}
func (self *Import_DependModuleInfo) InitImport_DependModuleInfo(arg1 LnsInt, arg2 *LnsMap) {
    self.id = arg1
    self.metaTypeId2TypeInfoMap = arg2
}
// 717: decl @lune.@base.@Import.DependModuleInfo.getTypeInfo
func (self *Import_DependModuleInfo) GetTypeInfo(metaTypeId LnsInt) *Ast_TypeInfo {
    return Lns_unwrap( self.metaTypeId2TypeInfoMap.Get(metaTypeId)).(*Ast_TypeInfo)
}


func Lns_Import_init() {
    if init_Import { return }
    init_Import = true
    Import__mod__ = "@lune.@base.@Import"
    Lns_InitMod()
    Lns_Types_init()
    Lns_Meta_init()
    Lns_Parser_init()
    Lns_Util_init()
    Lns_Ast_init()
    Lns_Macro_init()
    Lns_Nodes_init()
    Lns_frontInterface_init()
    Lns_Log_init()
    Lns_TransUnitIF_init()
    Lns_Builtin_init()
}
func init() {
    init_Import = false
}
