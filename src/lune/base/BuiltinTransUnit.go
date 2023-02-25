// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_BuiltinTransUnit bool
var BuiltinTransUnit__mod__ string
// for 103
func BuiltinTransUnit_convExp0_444(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 40: decl @lune.@base.@BuiltinTransUnit.TransUnit.get_scope
func (self *BuiltinTransUnit_TransUnit) Get_scope(_env *LnsEnv) *Ast_Scope {
    return self.scope
}
// 51: decl @lune.@base.@BuiltinTransUnit.TransUnit.error
func (self *BuiltinTransUnit_TransUnit) Error(_env *LnsEnv, mess string) {
    Util_err(_env, mess)
}
// 55: decl @lune.@base.@BuiltinTransUnit.TransUnit.pushScope
func (self *BuiltinTransUnit_TransUnit) PushScope(_env *LnsEnv, scopeKind LnsInt,baseInfo LnsAny,interfaceList LnsAny) *Ast_Scope {
    self.scope = Ast_TypeInfo_createScope(_env, self.processInfo, self.scope, scopeKind, baseInfo, interfaceList)
    return self.scope
}
// 63: decl @lune.@base.@BuiltinTransUnit.TransUnit.popScope
func (self *BuiltinTransUnit_TransUnit) PopScope(_env *LnsEnv) {
    self.scope = self.scope.FP.Get_outerScope(_env)
}
// 67: decl @lune.@base.@BuiltinTransUnit.TransUnit.getCurrentNamespaceTypeInfo
func (self *BuiltinTransUnit_TransUnit) getCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_getBuiltinMut(_env, self.scope.FP.GetNamespaceTypeInfo(_env))
}
// 71: decl @lune.@base.@BuiltinTransUnit.TransUnit.pushModuleLow
func (self *BuiltinTransUnit_TransUnit) PushModuleLow(_env *LnsEnv, processInfo *Ast_ProcessInfo,externalFlag bool,name string,mutable bool) *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var modName string
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(name,"^@", nil, nil))){
        modName = name
    } else { 
        modName = _env.GetVM().String_format("@%s", Lns_2DDD(name))
    }
    {
        __exp := self.scope.FP.GetTypeInfoChild(_env, modName)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            typeInfo = _exp
            {
                _scope := self.namespace2Scope.Get(typeInfo)
                if !Lns_IsNil( _scope ) {
                    scope := _scope.(*Ast_Scope)
                    self.scope = scope
                } else {
                    self.FP.Error(_env, _env.GetVM().String_format("not found scope -- %s", Lns_2DDD(name)))
                }
            }
        } else {
            var parentInfo *Ast_TypeInfo
            parentInfo = self.FP.getCurrentNamespaceTypeInfo(_env)
            var parentScope *Ast_Scope
            parentScope = self.scope
            var scope *Ast_Scope
            scope = self.FP.PushScope(_env, Ast_ScopeKind__Module, nil, nil)
            var newType *Ast_TypeInfo
            newType = processInfo.FP.CreateModule(_env, scope, parentInfo, parentInfo.FP, externalFlag, modName, mutable)
            typeInfo = newType
            self.namespace2Scope.Set(typeInfo,scope)
            Ast_addBuiltinMut(_env, newType, scope)
            var existSym LnsAny
            _,existSym = parentScope.FP.AddClass(_env, processInfo, modName, nil, typeInfo)
            if existSym != nil{
                existSym_98 := existSym.(*Ast_SymbolInfo)
                self.FP.Error(_env, _env.GetVM().String_format("module symbols exist -- %s.%s -- %s.%s", Lns_2DDD(existSym_98.FP.Get_namespaceTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil), existSym_98.FP.Get_name(_env), parentInfo.FP.GetTxt(_env, nil, nil, nil), modName)))
            }
        }
    }
    return typeInfo
}
// 114: decl @lune.@base.@BuiltinTransUnit.TransUnit.popModule
func (self *BuiltinTransUnit_TransUnit) PopModule(_env *LnsEnv) {
    self.FP.PopScope(_env)
}
// 121: decl @lune.@base.@BuiltinTransUnit.TransUnit.pushClassScope
func (self *BuiltinTransUnit_TransUnit) PushClassScope(_env *LnsEnv, errPos Types_Position,classTypeInfo *Ast_TypeInfo,scope *Ast_Scope) {
    if self.scope != _env.NilAccFin(_env.NilAccPush(classTypeInfo.FP.Get_scope(_env)) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_parent(_env)})){
        var classParentName string
        var classParentTypeId *Ast_IdInfo
        {
            _parentType := _env.NilAccFin(_env.NilAccPush(classTypeInfo.FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_parent(_env)})&&
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_ownerTypeInfo(_env)}))
            if !Lns_IsNil( _parentType ) {
                parentType := _parentType.(*Ast_TypeInfo)
                classParentName = parentType.FP.GetTxt(_env, nil, nil, nil)
                classParentTypeId = parentType.FP.Get_typeId(_env)
            } else {
                classParentName = "nil"
                classParentTypeId = Ast_dummyIdInfo
            }
        }
        self.FP.Error(_env, _env.GetVM().String_format("This class does not exist in this scope. -- %s -- %s(%d), %s(%d)", Lns_2DDD(classTypeInfo.FP.GetTxt(_env, nil, nil, nil), _env.NilAccFin(_env.NilAccPush(self.scope.FP.Get_ownerTypeInfo(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetTxt(_env, nil, nil, nil)})/* 137:15 */), _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(self.scope.FP.Get_ownerTypeInfo(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_typeId(_env)})&&
            _env.NilAccPush(_env.NilAccPop().(*Ast_IdInfo).Id))) ||
            _env.SetStackVal( -1) ).(LnsInt), classParentName, classParentTypeId.Id)))
    }
    self.scope = scope
}
// 149: decl @lune.@base.@BuiltinTransUnit.TransUnit.pushClassLow
func (self *BuiltinTransUnit_TransUnit) PushClassLow(_env *LnsEnv, processInfo *Ast_ProcessInfo,errPos Types_Position,mode LnsInt,finalFlag bool,abstractFlag bool,baseInfo LnsAny,interfaceList LnsAny,genTypeList LnsAny,externalFlag bool,name string,allowMultiple bool,accessMode LnsInt,defNamespace LnsAny) *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    {
        __exp := self.scope.FP.GetTypeInfo(_env, name, self.scope, true, Ast_ScopeAccess__Normal)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            typeInfo = _exp
            if typeInfo.FP.Get_abstractFlag(_env) != abstractFlag{
                self.FP.Error(_env, _env.GetVM().String_format("mismatch class(%s) abstract for prototype", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil))))
            }
            if typeInfo.FP.Get_finalFlag(_env) != finalFlag{
                self.FP.Error(_env, _env.GetVM().String_format("mismatch class(%s) final for prototype. %s, %s", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil), typeInfo.FP.Get_finalFlag(_env), finalFlag)))
            }
            if typeInfo.FP.Get_accessMode(_env) != accessMode{
                self.FP.Error(_env, _env.GetVM().String_format("mismatch class(%s) accessmode(%s) for prototype accessmode(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil), Ast_AccessMode_getTxt( accessMode), Ast_AccessMode_getTxt( typeInfo.FP.Get_accessMode(_env)))))
            }
            if baseInfo != nil{
                baseInfo_115 := baseInfo.(*Ast_TypeInfo)
                if typeInfo.FP.Get_baseTypeInfo(_env) != baseInfo_115{
                    self.FP.Error(_env, _env.GetVM().String_format("mismatch class(%s) base class(%s) for prototype base class(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil), baseInfo_115.FP.GetTxt(_env, nil, nil, nil), typeInfo.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil))))
                }
            } else {
                if typeInfo.FP.HasBase(_env){
                    self.FP.Error(_env, _env.GetVM().String_format("mismatch class(%s) base class(None) for prototype base class(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil), typeInfo.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil))))
                }
            }
            {
                var typeList *LnsList2_[*Ast_TypeInfo]
                typeList = Lns_unwrapDefault( interfaceList, NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})).(*LnsList2_[*Ast_TypeInfo])
                if typeInfo.FP.Get_interfaceList(_env).Len() == typeList.Len(){
                    for _index, _protoType := range( typeInfo.FP.Get_interfaceList(_env).Items ) {
                        index := _index + 1
                        protoType := _protoType
                        if protoType != typeList.GetAt(index){
                            self.FP.Error(_env, _env.GetVM().String_format("mismatch class(%s) %s(%s) for prototype %s(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil), "interface", typeList.GetAt(index).FP.GetTxt(_env, nil, nil, nil), "interface", protoType.FP.GetTxt(_env, nil, nil, nil))))
                        }
                    }
                } else { 
                    self.FP.Error(_env, _env.GetVM().String_format("mismatch class(%s) %s(%d) for prototype %s(%d)", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil), "interface", typeList.Len(), "interface", typeInfo.FP.Get_interfaceList(_env).Len())))
                }
            }
            
            {
                var typeList *LnsList2_[*Ast_AlternateTypeInfo]
                typeList = Lns_unwrapDefault( genTypeList, NewLnsList2_[*Ast_AlternateTypeInfo]([]*Ast_AlternateTypeInfo{})).(*LnsList2_[*Ast_AlternateTypeInfo])
                if typeInfo.FP.Get_itemTypeInfoList(_env).Len() == typeList.Len(){
                    for _index, _protoType := range( typeInfo.FP.Get_itemTypeInfoList(_env).Items ) {
                        index := _index + 1
                        protoType := _protoType
                        if protoType != &typeList.GetAt(index).Ast_TypeInfo{
                            self.FP.Error(_env, _env.GetVM().String_format("mismatch class(%s) %s(%s) for prototype %s(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil), "generics", typeList.GetAt(index).FP.GetTxt(_env, nil, nil, nil), "generics", protoType.FP.GetTxt(_env, nil, nil, nil))))
                        }
                    }
                } else { 
                    self.FP.Error(_env, _env.GetVM().String_format("mismatch class(%s) %s(%d) for prototype %s(%d)", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil), "generics", typeList.Len(), "generics", typeInfo.FP.Get_itemTypeInfoList(_env).Len())))
                }
            }
            
            {
                _scope := self.namespace2Scope.Get(typeInfo)
                if !Lns_IsNil( _scope ) {
                    scope := _scope.(*Ast_Scope)
                    self.scope = scope
                } else {
                    {
                        _scope := Ast_TypeInfo_getBuiltinInfo(_env, typeInfo).FP.Get_scope(_env)
                        if !Lns_IsNil( _scope ) {
                            scope := _scope.(*Ast_Scope)
                            self.scope = scope
                        } else {
                            self.FP.Error(_env, _env.GetVM().String_format("not find scope -- %s", Lns_2DDD(name)))
                        }
                    }
                }
            }
            if _switch0 := (typeInfo.FP.Get_kind(_env)); _switch0 == Ast_TypeInfoKind__Class {
                if mode == TransUnitIF_DeclClassMode__Interface{
                    self.FP.Error(_env, _env.GetVM().String_format("define interface already -- %s", Lns_2DDD(name)))
                }
            } else if _switch0 == Ast_TypeInfoKind__IF {
                if mode != TransUnitIF_DeclClassMode__Interface{
                    self.FP.Error(_env, _env.GetVM().String_format("define class already -- %s", Lns_2DDD(name)))
                }
            }
        } else {
            var parentInfo *Ast_TypeInfo
            parentInfo = self.FP.getCurrentNamespaceTypeInfo(_env)
            var parentScope *Ast_Scope
            parentScope = self.scope
            var scope *Ast_Scope
            scope = self.FP.PushScope(_env, Ast_ScopeKind__Class, baseInfo, interfaceList)
            var workGenTypeList *LnsList2_[*Ast_AlternateTypeInfo]
            if genTypeList != nil{
                genTypeList_158 := genTypeList.(*LnsList2_[*Ast_AlternateTypeInfo])
                workGenTypeList = genTypeList_158
            } else {
                workGenTypeList = NewLnsList2_[*Ast_AlternateTypeInfo]([]*Ast_AlternateTypeInfo{})
            }
            var newType *Ast_TypeInfo
            newType = processInfo.FP.CreateClassAsync(_env, mode != TransUnitIF_DeclClassMode__Interface, finalFlag, abstractFlag, scope, baseInfo, interfaceList, workGenTypeList, parentInfo, parentInfo.FP, externalFlag, accessMode, name)
            typeInfo = newType
            self.namespace2Scope.Set(typeInfo,scope)
            Ast_addBuiltinMut(_env, newType, scope)
            parentScope.FP.AddClassLazy(_env, processInfo, name, errPos, typeInfo, mode == TransUnitIF_DeclClassMode__LazyModule)
        }
    }
    if genTypeList != nil{
        genTypeList_162 := genTypeList.(*LnsList2_[*Ast_AlternateTypeInfo])
        for _, _genType := range( genTypeList_162.Items ) {
            genType := _genType
            self.scope.FP.AddAlternate(_env, processInfo, accessMode, genType.FP.Get_txt(_env), errPos, &genType.Ast_TypeInfo)
        }
    }
    return typeInfo
}
// 282: decl @lune.@base.@BuiltinTransUnit.TransUnit.popClass
func (self *BuiltinTransUnit_TransUnit) PopClass(_env *LnsEnv) {
    self.FP.PopScope(_env)
}
// declaration Class -- TransUnit
type BuiltinTransUnit_TransUnitMtd interface {
    Error(_env *LnsEnv, arg1 string)
    getCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_scope(_env *LnsEnv) *Ast_Scope
    PopClass(_env *LnsEnv)
    PopModule(_env *LnsEnv)
    PopScope(_env *LnsEnv)
    PushClassLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsAny, arg9 bool, arg10 string, arg11 bool, arg12 LnsInt, arg13 LnsAny) *Ast_TypeInfo
    PushClassScope(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
    PushModuleLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *Ast_TypeInfo
    PushScope(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny) *Ast_Scope
}
type BuiltinTransUnit_TransUnit struct {
    processInfo *Ast_ProcessInfo
    scope *Ast_Scope
    namespace2Scope *LnsMap2_[*Ast_TypeInfo,*Ast_Scope]
    FP BuiltinTransUnit_TransUnitMtd
}
func BuiltinTransUnit_TransUnit2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*BuiltinTransUnit_TransUnit).FP
}
func BuiltinTransUnit_TransUnit_toSlice(slice []LnsAny) []*BuiltinTransUnit_TransUnit {
    ret := make([]*BuiltinTransUnit_TransUnit, len(slice))
    for index, val := range slice {
        ret[index] = val.(BuiltinTransUnit_TransUnitDownCast).ToBuiltinTransUnit_TransUnit()
    }
    return ret
}
type BuiltinTransUnit_TransUnitDownCast interface {
    ToBuiltinTransUnit_TransUnit() *BuiltinTransUnit_TransUnit
}
func BuiltinTransUnit_TransUnitDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(BuiltinTransUnit_TransUnitDownCast)
    if ok { return work.ToBuiltinTransUnit_TransUnit() }
    return nil
}
func (obj *BuiltinTransUnit_TransUnit) ToBuiltinTransUnit_TransUnit() *BuiltinTransUnit_TransUnit {
    return obj
}
func NewBuiltinTransUnit_TransUnit(_env *LnsEnv, arg1 *Types_TransCtrlInfo, arg2 *Ast_ProcessInfo) *BuiltinTransUnit_TransUnit {
    obj := &BuiltinTransUnit_TransUnit{}
    obj.FP = obj
    obj.InitBuiltinTransUnit_TransUnit(_env, arg1, arg2)
    return obj
}
// 44: DeclConstr
func (self *BuiltinTransUnit_TransUnit) InitBuiltinTransUnit_TransUnit(_env *LnsEnv, ctrl_info *Types_TransCtrlInfo,processInfo *Ast_ProcessInfo) {
    self.namespace2Scope = NewLnsMap2_[*Ast_TypeInfo,*Ast_Scope]( map[*Ast_TypeInfo]*Ast_Scope{})
    self.processInfo = processInfo
    self.scope = processInfo.FP.Get_topScope(_env)
}


func Lns_BuiltinTransUnit_init(_env *LnsEnv) {
    if init_BuiltinTransUnit { return }
    init_BuiltinTransUnit = true
    BuiltinTransUnit__mod__ = "@lune.@base.@BuiltinTransUnit"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_Parser_init(_env)
    Lns_Util_init(_env)
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_TransUnitIF_init(_env)
}
func init() {
    init_BuiltinTransUnit = false
}
