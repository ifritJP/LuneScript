// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_TransUnitIF bool
var TransUnitIF__mod__ string
// decl enum -- DeclClassMode 
type TransUnitIF_DeclClassMode = LnsInt
const TransUnitIF_DeclClassMode__Class = 0
const TransUnitIF_DeclClassMode__Interface = 1
const TransUnitIF_DeclClassMode__LazyModule = 3
const TransUnitIF_DeclClassMode__Module = 2
var TransUnitIF_DeclClassModeList_ = NewLnsList( []LnsAny {
  TransUnitIF_DeclClassMode__Class,
  TransUnitIF_DeclClassMode__Interface,
  TransUnitIF_DeclClassMode__Module,
  TransUnitIF_DeclClassMode__LazyModule,
})
func TransUnitIF_DeclClassMode_get__allList(_env *LnsEnv) *LnsList{
    return TransUnitIF_DeclClassModeList_
}
var TransUnitIF_DeclClassModeMap_ = map[LnsInt]string {
  TransUnitIF_DeclClassMode__Class: "DeclClassMode.Class",
  TransUnitIF_DeclClassMode__Interface: "DeclClassMode.Interface",
  TransUnitIF_DeclClassMode__LazyModule: "DeclClassMode.LazyModule",
  TransUnitIF_DeclClassMode__Module: "DeclClassMode.Module",
}
func TransUnitIF_DeclClassMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := TransUnitIF_DeclClassModeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnitIF_DeclClassMode_getTxt(arg1 LnsInt) string {
    return TransUnitIF_DeclClassModeMap_[arg1];
}
// for 325
func TransUnitIF_convExp1112(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}

// declaration Class -- Modifier
type TransUnitIF_ModifierMtd interface {
    CreateModifier(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt) *Ast_TypeInfo
    Set_validMutControl(_env *LnsEnv, arg1 bool)
}
type TransUnitIF_Modifier struct {
    validMutControl bool
    processInfo *Ast_ProcessInfo
    FP TransUnitIF_ModifierMtd
}
func TransUnitIF_Modifier2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_Modifier).FP
}
type TransUnitIF_ModifierDownCast interface {
    ToTransUnitIF_Modifier() *TransUnitIF_Modifier
}
func TransUnitIF_ModifierDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnitIF_ModifierDownCast)
    if ok { return work.ToTransUnitIF_Modifier() }
    return nil
}
func (obj *TransUnitIF_Modifier) ToTransUnitIF_Modifier() *TransUnitIF_Modifier {
    return obj
}
func NewTransUnitIF_Modifier(_env *LnsEnv, arg1 bool, arg2 *Ast_ProcessInfo) *TransUnitIF_Modifier {
    obj := &TransUnitIF_Modifier{}
    obj.FP = obj
    obj.InitTransUnitIF_Modifier(_env, arg1, arg2)
    return obj
}
func (self *TransUnitIF_Modifier) InitTransUnitIF_Modifier(_env *LnsEnv, arg1 bool, arg2 *Ast_ProcessInfo) {
    self.validMutControl = arg1
    self.processInfo = arg2
}
func (self *TransUnitIF_Modifier) Set_validMutControl(_env *LnsEnv, arg1 bool){ self.validMutControl = arg1 }
// 45: decl @lune.@base.@TransUnitIF.Modifier.createModifier
func (self *TransUnitIF_Modifier) CreateModifier(_env *LnsEnv, typeInfo *Ast_TypeInfo,mutMode LnsInt) *Ast_TypeInfo {
    if Lns_op_not(self.validMutControl){
        return typeInfo
    }
    return self.processInfo.FP.CreateModifier(_env, typeInfo, mutMode)
}


// declaration Class -- IdSetInfo
type TransUnitIF_IdSetInfoMtd interface {
    RegisterSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) *Ast_SymbolInfo
}
type TransUnitIF_IdSetInfo struct {
    anonymousFuncId *Ast_IdProvider
    anonymousVarId *Ast_IdProvider
    FP TransUnitIF_IdSetInfoMtd
}
func TransUnitIF_IdSetInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_IdSetInfo).FP
}
type TransUnitIF_IdSetInfoDownCast interface {
    ToTransUnitIF_IdSetInfo() *TransUnitIF_IdSetInfo
}
func TransUnitIF_IdSetInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnitIF_IdSetInfoDownCast)
    if ok { return work.ToTransUnitIF_IdSetInfo() }
    return nil
}
func (obj *TransUnitIF_IdSetInfo) ToTransUnitIF_IdSetInfo() *TransUnitIF_IdSetInfo {
    return obj
}
func NewTransUnitIF_IdSetInfo(_env *LnsEnv) *TransUnitIF_IdSetInfo {
    obj := &TransUnitIF_IdSetInfo{}
    obj.FP = obj
    obj.InitTransUnitIF_IdSetInfo(_env)
    return obj
}
// 61: DeclConstr
func (self *TransUnitIF_IdSetInfo) InitTransUnitIF_IdSetInfo(_env *LnsEnv) {
    self.anonymousFuncId = NewAst_IdProvider(_env, 0, 10000)
    
    self.anonymousVarId = NewAst_IdProvider(_env, 0, 10000)
    
}

// 66: decl @lune.@base.@TransUnitIF.IdSetInfo.registerSym
func (self *TransUnitIF_IdSetInfo) RegisterSym(_env *LnsEnv, symbol *Ast_SymbolInfo) *Ast_SymbolInfo {
    if symbol.FP.Get_kind(_env) == Ast_SymbolKind__Var{
        if symbol.FP.Get_name(_env) == "_"{
            var id LnsInt
            id = self.anonymousVarId.FP.GetNewId(_env)
            return &NewAst_AnonymousSymbolInfo(_env, symbol, id).Ast_SymbolInfo
        }
    }
    return symbol
}


// declaration Class -- NSInfo
type TransUnitIF_NSInfoMtd interface {
    CanAccessNoasync(_env *LnsEnv) bool
    CanBreak(_env *LnsEnv) bool
    DecLock(_env *LnsEnv)
    Get_loopScopeQueue(_env *LnsEnv) *LnsList
    Get_nobody(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) *Types_Position
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    IncLock(_env *LnsEnv)
    IsLockedAsync(_env *LnsEnv) bool
    RegisterSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) *Ast_SymbolInfo
    Set_nobody(_env *LnsEnv, arg1 bool)
}
type TransUnitIF_NSInfo struct {
    nobody bool
    typeInfo *Ast_TypeInfo
    pos *Types_Position
    loopScopeQueue *LnsList
    lockedAsyncStack *LnsList
    idSetInfo *TransUnitIF_IdSetInfo
    FP TransUnitIF_NSInfoMtd
}
func TransUnitIF_NSInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_NSInfo).FP
}
type TransUnitIF_NSInfoDownCast interface {
    ToTransUnitIF_NSInfo() *TransUnitIF_NSInfo
}
func TransUnitIF_NSInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnitIF_NSInfoDownCast)
    if ok { return work.ToTransUnitIF_NSInfo() }
    return nil
}
func (obj *TransUnitIF_NSInfo) ToTransUnitIF_NSInfo() *TransUnitIF_NSInfo {
    return obj
}
func NewTransUnitIF_NSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Types_Position) *TransUnitIF_NSInfo {
    obj := &TransUnitIF_NSInfo{}
    obj.FP = obj
    obj.InitTransUnitIF_NSInfo(_env, arg1, arg2)
    return obj
}
func (self *TransUnitIF_NSInfo) Get_nobody(_env *LnsEnv) bool{ return self.nobody }
func (self *TransUnitIF_NSInfo) Set_nobody(_env *LnsEnv, arg1 bool){ self.nobody = arg1 }
func (self *TransUnitIF_NSInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *TransUnitIF_NSInfo) Get_pos(_env *LnsEnv) *Types_Position{ return self.pos }
func (self *TransUnitIF_NSInfo) Get_loopScopeQueue(_env *LnsEnv) *LnsList{ return self.loopScopeQueue }
// advertise -- 80
func (self *TransUnitIF_NSInfo) RegisterSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) *Ast_SymbolInfo {
    return self.idSetInfo. FP.RegisterSym( _env, arg1)
}
// 99: decl @lune.@base.@TransUnitIF.NSInfo.isLockedAsync
func (self *TransUnitIF_NSInfo) IsLockedAsync(_env *LnsEnv) bool {
    return self.lockedAsyncStack.Len() > 0
}

// 103: DeclConstr
func (self *TransUnitIF_NSInfo) InitTransUnitIF_NSInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,pos *Types_Position) {
    self.idSetInfo = NewTransUnitIF_IdSetInfo(_env)
    
    self.nobody = false
    
    self.lockedAsyncStack = NewLnsList([]LnsAny{})
    
    self.loopScopeQueue = NewLnsList([]LnsAny{})
    
    self.typeInfo = typeInfo
    
    self.pos = pos
    
}

// 113: decl @lune.@base.@TransUnitIF.NSInfo.incLock
func (self *TransUnitIF_NSInfo) IncLock(_env *LnsEnv) {
    self.lockedAsyncStack.Insert(self.loopScopeQueue.Len())
}

// 116: decl @lune.@base.@TransUnitIF.NSInfo.decLock
func (self *TransUnitIF_NSInfo) DecLock(_env *LnsEnv) {
    self.lockedAsyncStack.Remove(nil)
}

// 129: decl @lune.@base.@TransUnitIF.NSInfo.canBreak
func (self *TransUnitIF_NSInfo) CanBreak(_env *LnsEnv) bool {
    var len LnsInt
    len = self.lockedAsyncStack.Len()
    var loopQueueLen LnsInt
    loopQueueLen = self.loopScopeQueue.Len()
    if len == 0{
        return loopQueueLen > 0
    }
    return self.lockedAsyncStack.GetAt(len).(LnsInt) < loopQueueLen
}

// 141: decl @lune.@base.@TransUnitIF.NSInfo.canAccessNoasync
func (self *TransUnitIF_NSInfo) CanAccessNoasync(_env *LnsEnv) bool {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.typeInfo.FP.Get_asyncMode(_env) == Ast_Async__Noasync) ||
        _env.SetStackVal( self.lockedAsyncStack.Len() > 0) ).(bool){
        return true
    }
    return false
}


type TransUnitIF_TransUnitIF interface {
        Error(_env *LnsEnv, arg1 string)
        Get_scope(_env *LnsEnv) *Ast_Scope
        PopClass(_env *LnsEnv)
        PopModule(_env *LnsEnv)
        PushClassScope(_env *LnsEnv, arg1 *Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
        PushModule(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *TransUnitIF_NSInfo
        PushModuleLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *Ast_TypeInfo
}
func Lns_cast2TransUnitIF_TransUnitIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(TransUnitIF_TransUnitIF); ok { 
        return obj
    }
    return nil
}

// declaration Class -- TransUnitBase
type TransUnitIF_TransUnitBaseMtd interface {
    AddErrMess(_env *LnsEnv, arg1 *Types_Position, arg2 string)
    Error(_env *LnsEnv, arg1 string)
    ErrorAt(_env *LnsEnv, arg1 *Types_Position, arg2 string)
    GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    GetCurrentNamespaceTypeInfoMut(_env *LnsEnv) *Ast_TypeInfo
    GetLatestPos(_env *LnsEnv) *Types_Position
    Get_errMessList(_env *LnsEnv) *LnsList
    Get_globalScope(_env *LnsEnv) *Ast_Scope
    Get_scope(_env *LnsEnv) *Ast_Scope
    NewNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Types_Position) *TransUnitIF_NSInfo
    PopClass(_env *LnsEnv)
    PopModule(_env *LnsEnv)
    PopScope(_env *LnsEnv)
    PushClass(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *TransUnitIF_NSInfo
    PushClassLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *Ast_TypeInfo
    PushClassScope(_env *LnsEnv, arg1 *Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
    PushModule(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *TransUnitIF_NSInfo
    PushModuleLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *Ast_TypeInfo
    PushScope(_env *LnsEnv, arg1 bool, arg2 LnsAny, arg3 LnsAny) *Ast_Scope
}
type TransUnitIF_TransUnitBase struct {
    ProcessInfo *Ast_ProcessInfo
    Scope *Ast_Scope
    GlobalScope *Ast_Scope
    Namespace2Scope *LnsMap
    NsInfoMap *LnsMap
    ErrMessList *LnsList
    TypeNameCtrl *Ast_TypeNameCtrl
    TypeId2ClassMap *LnsMap
    FP TransUnitIF_TransUnitBaseMtd
}
func TransUnitIF_TransUnitBase2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_TransUnitBase).FP
}
type TransUnitIF_TransUnitBaseDownCast interface {
    ToTransUnitIF_TransUnitBase() *TransUnitIF_TransUnitBase
}
func TransUnitIF_TransUnitBaseDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnitIF_TransUnitBaseDownCast)
    if ok { return work.ToTransUnitIF_TransUnitBase() }
    return nil
}
func (obj *TransUnitIF_TransUnitBase) ToTransUnitIF_TransUnitBase() *TransUnitIF_TransUnitBase {
    return obj
}
func (self *TransUnitIF_TransUnitBase) Get_globalScope(_env *LnsEnv) *Ast_Scope{ return self.GlobalScope }
func (self *TransUnitIF_TransUnitBase) Get_errMessList(_env *LnsEnv) *LnsList{ return self.ErrMessList }
// 190: decl @lune.@base.@TransUnitIF.TransUnitBase.get_scope
func (self *TransUnitIF_TransUnitBase) Get_scope(_env *LnsEnv) *Ast_Scope {
    return self.Scope
}

// 194: DeclConstr
func (self *TransUnitIF_TransUnitBase) InitTransUnitIF_TransUnitBase(_env *LnsEnv, ctrl_info *Types_TransCtrlInfo,processInfo *Ast_ProcessInfo) {
    self.TypeId2ClassMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.TypeNameCtrl = Ast_defaultTypeNameCtrl
    
    self.ErrMessList = NewLnsList([]LnsAny{})
    
    self.Namespace2Scope = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.ProcessInfo = processInfo
    
    self.GlobalScope = NewAst_Scope(_env, processInfo, processInfo.FP.Get_topScope(_env), true, nil, nil)
    
    self.Scope = NewAst_Scope(_env, processInfo, self.GlobalScope, true, nil, nil)
    
    self.NsInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    var subRootTypeInfo *Ast_TypeInfo
    subRootTypeInfo = self.ProcessInfo.FP.Get_dummyParentType(_env)
    self.NsInfoMap.Set(subRootTypeInfo,NewTransUnitIF_NSInfo(_env, subRootTypeInfo, NewTypes_Position(_env, 0, 0, "@builtin@")))
}

// 211: decl @lune.@base.@TransUnitIF.TransUnitBase.addErrMess
func (self *TransUnitIF_TransUnitBase) AddErrMess(_env *LnsEnv, pos *Types_Position,mess string) {
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(mess,"type mismatch.*<- &", nil, nil))){
        mess = mess + ". if your code is the old style, use the opiton '--legacy-mutable-control'."
        
    }
    self.ErrMessList.Insert(_env.LuaVM.String_format("%s:%d:%d: error: %s", []LnsAny{pos.StreamName, pos.LineNo, pos.Column, mess}))
}



// 241: decl @lune.@base.@TransUnitIF.TransUnitBase.error
func (self *TransUnitIF_TransUnitBase) Error(_env *LnsEnv, mess string) {
    self.FP.ErrorAt(_env, self.FP.GetLatestPos(_env), mess)
}

// 245: decl @lune.@base.@TransUnitIF.TransUnitBase.pushScope
func (self *TransUnitIF_TransUnitBase) PushScope(_env *LnsEnv, classFlag bool,baseInfo LnsAny,interfaceList LnsAny) *Ast_Scope {
    self.Scope = Ast_TypeInfo_createScope(_env, self.ProcessInfo, self.Scope, classFlag, baseInfo, interfaceList)
    
    return self.Scope
}

// 253: decl @lune.@base.@TransUnitIF.TransUnitBase.popScope
func (self *TransUnitIF_TransUnitBase) PopScope(_env *LnsEnv) {
    self.Scope = self.Scope.FP.Get_parent(_env)
    
}

// 261: decl @lune.@base.@TransUnitIF.TransUnitBase.newNSInfo
func (self *TransUnitIF_TransUnitBase) NewNSInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,pos *Types_Position) *TransUnitIF_NSInfo {
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = NewTransUnitIF_NSInfo(_env, typeInfo, pos)
    self.NsInfoMap.Set(typeInfo,nsInfo)
    return nsInfo
}

// 278: decl @lune.@base.@TransUnitIF.TransUnitBase.getCurrentNamespaceTypeInfoMut
func (self *TransUnitIF_TransUnitBase) GetCurrentNamespaceTypeInfoMut(_env *LnsEnv) *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = self.Scope.FP.GetNamespaceTypeInfo(_env)
    var nsInfo *TransUnitIF_NSInfo
    
    {
        _nsInfo := self.NsInfoMap.Get(typeInfo)
        if _nsInfo == nil{
            self.FP.Error(_env, _env.LuaVM.String_format("not found nsInfo -- %s (%d)", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil), typeInfo.FP.Get_typeId(_env).Id}))
        } else {
            nsInfo = _nsInfo.(*TransUnitIF_NSInfo)
        }
    }
    return nsInfo.FP.Get_typeInfo(_env)
}

// 287: decl @lune.@base.@TransUnitIF.TransUnitBase.getCurrentNamespaceTypeInfo
func (self *TransUnitIF_TransUnitBase) GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.Scope.FP.GetNamespaceTypeInfo(_env)
}

// 291: decl @lune.@base.@TransUnitIF.TransUnitBase.pushModule
func (self *TransUnitIF_TransUnitBase) PushModule(_env *LnsEnv, processInfo *Ast_ProcessInfo,externalFlag bool,name string,mutable bool) *TransUnitIF_NSInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var modName string
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(name,"^@", nil, nil))){
        modName = name
        
    } else { 
        modName = _env.LuaVM.String_format("@%s", []LnsAny{name})
        
    }
    var nsInfo *TransUnitIF_NSInfo
    {
        __exp := self.Scope.FP.GetTypeInfoChild(_env, modName)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            typeInfo = _exp
            
            {
                _scope := self.Namespace2Scope.Get(typeInfo)
                if !Lns_IsNil( _scope ) {
                    scope := _scope.(*Ast_Scope)
                    self.Scope = scope
                    
                } else {
                    self.FP.Error(_env, _env.LuaVM.String_format("not found scope -- %s", []LnsAny{name}))
                }
            }
            nsInfo = Lns_unwrap( self.NsInfoMap.Get(typeInfo)).(*TransUnitIF_NSInfo)
            
        } else {
            var parentInfo *Ast_TypeInfo
            parentInfo = self.FP.GetCurrentNamespaceTypeInfoMut(_env)
            var parentScope *Ast_Scope
            parentScope = self.Scope
            var scope *Ast_Scope
            scope = self.FP.PushScope(_env, true, nil, nil)
            var newType *Ast_TypeInfo
            newType = processInfo.FP.CreateModule(_env, scope, parentInfo, externalFlag, modName, mutable)
            typeInfo = newType
            
            self.Namespace2Scope.Set(typeInfo,scope)
            nsInfo = self.FP.NewNSInfo(_env, newType, self.FP.GetLatestPos(_env))
            
            var existSym LnsAny
            _,existSym = parentScope.FP.AddClass(_env, processInfo, modName, nil, typeInfo)
            if existSym != nil{
                existSym_212 := existSym.(*Ast_SymbolInfo)
                self.FP.AddErrMess(_env, self.FP.GetLatestPos(_env), _env.LuaVM.String_format("module symbols exist -- %s.%s -- %s.%s", []LnsAny{existSym_212.FP.Get_namespaceTypeInfo(_env).FP.GetFullName(_env, self.TypeNameCtrl, parentScope.FP, false), existSym_212.FP.Get_name(_env), parentInfo.FP.GetFullName(_env, self.TypeNameCtrl, parentScope.FP, false), modName}))
            }
        }
    }
    if Lns_op_not(self.TypeId2ClassMap.Get(typeInfo.FP.Get_typeId(_env))){
        var namespace *Nodes_NamespaceInfo
        namespace = NewNodes_NamespaceInfo(_env, modName, self.Scope, typeInfo)
        self.TypeId2ClassMap.Set(typeInfo.FP.Get_typeId(_env),namespace)
    }
    return nsInfo
}

// 344: decl @lune.@base.@TransUnitIF.TransUnitBase.pushModuleLow
func (self *TransUnitIF_TransUnitBase) PushModuleLow(_env *LnsEnv, processInfo *Ast_ProcessInfo,externalFlag bool,name string,mutable bool) *Ast_TypeInfo {
    return self.FP.PushModule(_env, processInfo, externalFlag, name, mutable).FP.Get_typeInfo(_env)
}

// 351: decl @lune.@base.@TransUnitIF.TransUnitBase.popModule
func (self *TransUnitIF_TransUnitBase) PopModule(_env *LnsEnv) {
    self.FP.PopScope(_env)
}

// 358: decl @lune.@base.@TransUnitIF.TransUnitBase.pushClassScope
func (self *TransUnitIF_TransUnitBase) PushClassScope(_env *LnsEnv, errPos *Types_Position,classTypeInfo *Ast_TypeInfo,scope *Ast_Scope) {
    if self.Scope != _env.NilAccFin(_env.NilAccPush(classTypeInfo.FP.Get_scope(_env)) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_parent(_env)})){
        var classParentName string
        var classParentTypeId *Ast_IdInfo
        {
            _parentType := _env.NilAccFin(_env.NilAccPush(classTypeInfo.FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_parent(_env)})&&
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_ownerTypeInfo(_env)}))
            if !Lns_IsNil( _parentType ) {
                parentType := _parentType.(*Ast_TypeInfo)
                classParentName = parentType.FP.GetFullName(_env, self.TypeNameCtrl, Lns_unwrap( parentType.FP.Get_scope(_env)).(*Ast_Scope).FP, false)
                
                classParentTypeId = parentType.FP.Get_typeId(_env)
                
            } else {
                classParentName = "nil"
                
                classParentTypeId = Ast_dummyIdInfo
                
            }
        }
        self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("This class does not exist in this scope. -- %s -- %s(%d), %s(%d)", []LnsAny{classTypeInfo.FP.GetTxt(_env, nil, nil, nil), _env.NilAccFin(_env.NilAccPush(self.Scope.FP.Get_ownerTypeInfo(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetFullName(_env, self.TypeNameCtrl, self.Scope.FP, false)})/* 376:15 */), _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(self.Scope.FP.Get_ownerTypeInfo(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_typeId(_env)})&&
            _env.NilAccPush(_env.NilAccPop().(*Ast_IdInfo).Id))) ||
            _env.SetStackVal( -1) ).(LnsInt), classParentName, classParentTypeId.Id}))
    }
    self.Scope = scope
    
}

// 390: decl @lune.@base.@TransUnitIF.TransUnitBase.pushClass
func (self *TransUnitIF_TransUnitBase) PushClass(_env *LnsEnv, processInfo *Ast_ProcessInfo,errPos *Types_Position,mode LnsInt,abstractFlag bool,baseInfo LnsAny,interfaceList LnsAny,genTypeList LnsAny,externalFlag bool,name string,allowMultiple bool,accessMode LnsInt,defNamespace LnsAny) *TransUnitIF_NSInfo {
    var typeInfo *Ast_TypeInfo
    var nsInfo *TransUnitIF_NSInfo
    {
        __exp := self.Scope.FP.GetTypeInfo(_env, name, self.Scope, true, Ast_ScopeAccess__Normal)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            typeInfo = _exp
            
            nsInfo = Lns_unwrap( self.NsInfoMap.Get(typeInfo)).(*TransUnitIF_NSInfo)
            
            if _env.NilAccFin(_env.NilAccPush(typeInfo.FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_parent(_env)})) != self.Scope{
                self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("multiple class(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
                self.FP.Error(_env, "stop by error")
            }
            if typeInfo.FP.Get_abstractFlag(_env) != abstractFlag{
                self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("mismatch class(%s) abstract for prototpye", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
            }
            if typeInfo.FP.Get_accessMode(_env) != accessMode{
                self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("mismatch class(%s) accessmode(%s) for prototpye accessmode(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), Ast_AccessMode_getTxt( accessMode), Ast_AccessMode_getTxt( typeInfo.FP.Get_accessMode(_env))}))
            }
            if baseInfo != nil{
                baseInfo_264 := baseInfo.(*Ast_TypeInfo)
                if typeInfo.FP.Get_baseTypeInfo(_env) != baseInfo_264{
                    self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("mismatch class(%s) base class(%s) for prototpye base class(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), baseInfo_264.FP.GetTxt(_env, nil, nil, nil), typeInfo.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            } else {
                if typeInfo.FP.HasBase(_env){
                    self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("mismatch class(%s) base class(None) for prototpye base class(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), typeInfo.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
            var compareList func(_env *LnsEnv, protoList *LnsList,typeList *LnsList,message string)
            compareList = func(_env *LnsEnv, protoList *LnsList,typeList *LnsList,message string) {
                if protoList.Len() == typeList.Len(){
                    for _index, _protoType := range( protoList.Items ) {
                        index := _index + 1
                        protoType := _protoType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if protoType != typeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(){
                            self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("mismatch class(%s) %s(%s) for prototpye %s(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), message, typeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), message, protoType.FP.GetTxt(_env, nil, nil, nil)}))
                        }
                    }
                } else { 
                    self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("mismatch class(%s) %s(%d) for prototpye %s(%d)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), message, typeList.Len(), message, protoList.Len()}))
                }
            }
            compareList(_env, typeInfo.FP.Get_interfaceList(_env), Lns_unwrapDefault( interfaceList, NewLnsList([]LnsAny{})).(*LnsList), "interface")
            compareList(_env, typeInfo.FP.Get_itemTypeInfoList(_env), Lns_unwrapDefault( genTypeList, NewLnsList([]LnsAny{})).(*LnsList), "generics")
            {
                _scope := self.Namespace2Scope.Get(typeInfo)
                if !Lns_IsNil( _scope ) {
                    scope := _scope.(*Ast_Scope)
                    self.Scope = scope
                    
                } else {
                    self.FP.ErrorAt(_env, errPos, _env.LuaVM.String_format("not find scope -- %s", []LnsAny{name}))
                }
            }
            if _switch1 := (typeInfo.FP.Get_kind(_env)); _switch1 == Ast_TypeInfoKind__Class {
                if mode == TransUnitIF_DeclClassMode__Interface{
                    self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("define interface already -- %s", []LnsAny{name}))
                    Util_printStackTrace(_env)
                }
            } else if _switch1 == Ast_TypeInfoKind__IF {
                if mode != TransUnitIF_DeclClassMode__Interface{
                    self.FP.AddErrMess(_env, errPos, _env.LuaVM.String_format("define class already -- %s", []LnsAny{name}))
                    Util_printStackTrace(_env)
                }
            }
        } else {
            var parentInfo *Ast_TypeInfo
            parentInfo = self.FP.GetCurrentNamespaceTypeInfoMut(_env)
            var parentScope *Ast_Scope
            parentScope = self.Scope
            var scope *Ast_Scope
            scope = self.FP.PushScope(_env, true, baseInfo, interfaceList)
            var workGenTypeList *LnsList
            if genTypeList != nil{
                genTypeList_293 := genTypeList.(*LnsList)
                workGenTypeList = genTypeList_293
                
            } else {
                workGenTypeList = NewLnsList([]LnsAny{})
                
            }
            var newType *Ast_TypeInfo
            newType = processInfo.FP.CreateClassAsync(_env, mode != TransUnitIF_DeclClassMode__Interface, abstractFlag, scope, baseInfo, interfaceList, workGenTypeList, parentInfo, externalFlag, accessMode, name)
            typeInfo = newType
            
            self.Namespace2Scope.Set(typeInfo,scope)
            nsInfo = self.FP.NewNSInfo(_env, newType, errPos)
            
            parentScope.FP.AddClassLazy(_env, processInfo, name, errPos, typeInfo, mode == TransUnitIF_DeclClassMode__LazyModule)
        }
    }
    if genTypeList != nil{
        genTypeList_297 := genTypeList.(*LnsList)
        for _, _genType := range( genTypeList_297.Items ) {
            genType := _genType.(Ast_AlternateTypeInfoDownCast).ToAst_AlternateTypeInfo()
            self.Scope.FP.AddAlternate(_env, processInfo, accessMode, genType.FP.Get_txt(_env), errPos, &genType.Ast_TypeInfo)
        }
    }
    var namespace *Nodes_NamespaceInfo
    
    {
        _namespace := defNamespace
        if _namespace == nil{
            namespace = NewNodes_NamespaceInfo(_env, name, self.Scope, typeInfo)
            
        } else {
            namespace = _namespace.(*Nodes_NamespaceInfo)
        }
    }
    self.TypeId2ClassMap.Set(typeInfo.FP.Get_typeId(_env),namespace)
    return nsInfo
}

// 542: decl @lune.@base.@TransUnitIF.TransUnitBase.pushClassLow
func (self *TransUnitIF_TransUnitBase) PushClassLow(_env *LnsEnv, processInfo *Ast_ProcessInfo,errPos *Types_Position,mode LnsInt,abstractFlag bool,baseInfo LnsAny,interfaceList LnsAny,genTypeList LnsAny,externalFlag bool,name string,allowMultiple bool,accessMode LnsInt,defNamespace LnsAny) *Ast_TypeInfo {
    return self.FP.PushClass(_env, processInfo, errPos, mode, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace).FP.Get_typeInfo(_env)
}

// 556: decl @lune.@base.@TransUnitIF.TransUnitBase.popClass
func (self *TransUnitIF_TransUnitBase) PopClass(_env *LnsEnv) {
    self.FP.PopScope(_env)
}


// declaration Class -- SimpeTransUnit
type TransUnitIF_SimpeTransUnitMtd interface {
    AddErrMess(_env *LnsEnv, arg1 *Types_Position, arg2 string)
    Error(_env *LnsEnv, arg1 string)
    ErrorAt(_env *LnsEnv, arg1 *Types_Position, arg2 string)
    GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    GetCurrentNamespaceTypeInfoMut(_env *LnsEnv) *Ast_TypeInfo
    GetLatestPos(_env *LnsEnv) *Types_Position
    Get_errMessList(_env *LnsEnv) *LnsList
    Get_globalScope(_env *LnsEnv) *Ast_Scope
    Get_scope(_env *LnsEnv) *Ast_Scope
    NewNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Types_Position) *TransUnitIF_NSInfo
    PopClass(_env *LnsEnv)
    PopModule(_env *LnsEnv)
    PopScope(_env *LnsEnv)
    PushClass(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *TransUnitIF_NSInfo
    PushClassLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *Ast_TypeInfo
    PushClassScope(_env *LnsEnv, arg1 *Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
    PushModule(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *TransUnitIF_NSInfo
    PushModuleLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *Ast_TypeInfo
    PushScope(_env *LnsEnv, arg1 bool, arg2 LnsAny, arg3 LnsAny) *Ast_Scope
}
type TransUnitIF_SimpeTransUnit struct {
    TransUnitIF_TransUnitBase
    latestPos *Types_Position
    macroMode string
    nearCode LnsAny
    FP TransUnitIF_SimpeTransUnitMtd
}
func TransUnitIF_SimpeTransUnit2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_SimpeTransUnit).FP
}
type TransUnitIF_SimpeTransUnitDownCast interface {
    ToTransUnitIF_SimpeTransUnit() *TransUnitIF_SimpeTransUnit
}
func TransUnitIF_SimpeTransUnitDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnitIF_SimpeTransUnitDownCast)
    if ok { return work.ToTransUnitIF_SimpeTransUnit() }
    return nil
}
func (obj *TransUnitIF_SimpeTransUnit) ToTransUnitIF_SimpeTransUnit() *TransUnitIF_SimpeTransUnit {
    return obj
}
func NewTransUnitIF_SimpeTransUnit(_env *LnsEnv, arg1 *Types_TransCtrlInfo, arg2 *Ast_ProcessInfo, arg3 *Types_Position, arg4 string, arg5 LnsAny) *TransUnitIF_SimpeTransUnit {
    obj := &TransUnitIF_SimpeTransUnit{}
    obj.FP = obj
    obj.TransUnitIF_TransUnitBase.FP = obj
    obj.InitTransUnitIF_SimpeTransUnit(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *TransUnitIF_SimpeTransUnit) InitTransUnitIF_SimpeTransUnit(_env *LnsEnv, arg1 *Types_TransCtrlInfo, arg2 *Ast_ProcessInfo, arg3 *Types_Position, arg4 string, arg5 LnsAny) {
    self.TransUnitIF_TransUnitBase.InitTransUnitIF_TransUnitBase( _env, arg1,arg2)
    self.latestPos = arg3
    self.macroMode = arg4
    self.nearCode = arg5
}
// 565: decl @lune.@base.@TransUnitIF.SimpeTransUnit.errorAt
func (self *TransUnitIF_SimpeTransUnit) ErrorAt(_env *LnsEnv, pos *Types_Position,mess string) {
    self.FP.AddErrMess(_env, pos, mess)
    for _, _errmess := range( self.ErrMessList.Items ) {
        errmess := _errmess.(string)
        Util_errorLog(_env, errmess)
    }
    {
        _nearCode := self.nearCode
        if !Lns_IsNil( _nearCode ) {
            nearCode := _nearCode.(string)
            Lns_print([]LnsAny{"------ near code -----", self.macroMode})
            Lns_print([]LnsAny{nearCode})
            Lns_print([]LnsAny{"------"})
        }
    }
    Util_err(_env, "has error")
}

// 579: decl @lune.@base.@TransUnitIF.SimpeTransUnit.getLatestPos
func (self *TransUnitIF_SimpeTransUnit) GetLatestPos(_env *LnsEnv) *Types_Position {
    return self.latestPos
}


func Lns_TransUnitIF_init(_env *LnsEnv) {
    if init_TransUnitIF { return }
    init_TransUnitIF = true
    TransUnitIF__mod__ = "@lune.@base.@TransUnitIF"
    Lns_InitMod()
    Lns_Parser_init(_env)
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Util_init(_env)
    Lns_Types_init(_env)
}
func init() {
    init_TransUnitIF = false
}
