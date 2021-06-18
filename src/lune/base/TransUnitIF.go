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
// for 405
func TransUnitIF_convExp1578(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 210: decl @lune.@base.@TransUnitIF.sortMess
func TransUnitIF_sortMess(_env *LnsEnv, list *LnsList) *LnsList {
    list.Sort(_env, LnsItemKindStem, LnsComp(func ( _env *LnsEnv, val1, val2 LnsAny ) bool {return sortMess___anonymous_0_( _env, val1.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess(), val2.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess())}))
    return list
}

func sortMess___anonymous_0_(_env *LnsEnv, mess1 *TransUnitIF_ErrMess,mess2 *TransUnitIF_ErrMess) bool {
    var pos1 *Types_Position
    pos1 = mess1.FP.Get_pos(_env).FP.Get_orgPos(_env)
    var pos2 *Types_Position
    pos2 = mess2.FP.Get_pos(_env).FP.Get_orgPos(_env)
    if pos1.StreamName < pos2.StreamName{
        return true
    }
    if pos1.StreamName > pos2.StreamName{
        return false
    }
    if pos1.LineNo < pos2.LineNo{
        return true
    }
    if pos1.LineNo > pos2.LineNo{
        return false
    }
    if pos1.Column < pos2.Column{
        return true
    }
    if pos1.Column > pos2.Column{
        return false
    }
    return mess1.FP.Get_mess(_env) < mess2.FP.Get_mess(_env)
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


// declaration Class -- LockedAsyncInfo
type TransUnitIF_LockedAsyncInfoMtd interface {
    Get_lockKind(_env *LnsEnv) LnsInt
    Get_loopLen(_env *LnsEnv) LnsInt
}
type TransUnitIF_LockedAsyncInfo struct {
    loopLen LnsInt
    lockKind LnsInt
    FP TransUnitIF_LockedAsyncInfoMtd
}
func TransUnitIF_LockedAsyncInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_LockedAsyncInfo).FP
}
type TransUnitIF_LockedAsyncInfoDownCast interface {
    ToTransUnitIF_LockedAsyncInfo() *TransUnitIF_LockedAsyncInfo
}
func TransUnitIF_LockedAsyncInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnitIF_LockedAsyncInfoDownCast)
    if ok { return work.ToTransUnitIF_LockedAsyncInfo() }
    return nil
}
func (obj *TransUnitIF_LockedAsyncInfo) ToTransUnitIF_LockedAsyncInfo() *TransUnitIF_LockedAsyncInfo {
    return obj
}
func NewTransUnitIF_LockedAsyncInfo(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *TransUnitIF_LockedAsyncInfo {
    obj := &TransUnitIF_LockedAsyncInfo{}
    obj.FP = obj
    obj.InitTransUnitIF_LockedAsyncInfo(_env, arg1, arg2)
    return obj
}
func (self *TransUnitIF_LockedAsyncInfo) InitTransUnitIF_LockedAsyncInfo(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) {
    self.loopLen = arg1
    self.lockKind = arg2
}
func (self *TransUnitIF_LockedAsyncInfo) Get_loopLen(_env *LnsEnv) LnsInt{ return self.loopLen }
func (self *TransUnitIF_LockedAsyncInfo) Get_lockKind(_env *LnsEnv) LnsInt{ return self.lockKind }

// declaration Class -- NSInfo
type TransUnitIF_NSInfoMtd interface {
    CanAccessLuaval(_env *LnsEnv) bool
    CanAccessNoasync(_env *LnsEnv) bool
    CanBreak(_env *LnsEnv) bool
    DecLock(_env *LnsEnv)
    Get_loopScopeQueue(_env *LnsEnv) *LnsList
    Get_nobody(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) *Types_Position
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    IncLock(_env *LnsEnv, arg1 LnsInt)
    IsLockedAsync(_env *LnsEnv) bool
    IsNoasync(_env *LnsEnv) bool
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
    validAsyncCtrl bool
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
func NewTransUnitIF_NSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Types_Position, arg3 bool) *TransUnitIF_NSInfo {
    obj := &TransUnitIF_NSInfo{}
    obj.FP = obj
    obj.InitTransUnitIF_NSInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *TransUnitIF_NSInfo) Get_nobody(_env *LnsEnv) bool{ return self.nobody }
func (self *TransUnitIF_NSInfo) Set_nobody(_env *LnsEnv, arg1 bool){ self.nobody = arg1 }
func (self *TransUnitIF_NSInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *TransUnitIF_NSInfo) Get_pos(_env *LnsEnv) *Types_Position{ return self.pos }
func (self *TransUnitIF_NSInfo) Get_loopScopeQueue(_env *LnsEnv) *LnsList{ return self.loopScopeQueue }
// advertise -- 84
func (self *TransUnitIF_NSInfo) RegisterSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) *Ast_SymbolInfo {
    return self.idSetInfo. FP.RegisterSym( _env, arg1)
}
// 106: decl @lune.@base.@TransUnitIF.NSInfo.isLockedAsync
func (self *TransUnitIF_NSInfo) IsLockedAsync(_env *LnsEnv) bool {
    return self.lockedAsyncStack.Len() > 0
}

// 110: decl @lune.@base.@TransUnitIF.NSInfo.isNoasync
func (self *TransUnitIF_NSInfo) IsNoasync(_env *LnsEnv) bool {
    if self.typeInfo.FP.Get_asyncMode(_env) == Ast_Async__Noasync{
        return true
    }
    for _, _info := range( self.lockedAsyncStack.Items ) {
        info := _info.(TransUnitIF_LockedAsyncInfoDownCast).ToTransUnitIF_LockedAsyncInfo()
        if _switch1 := info.FP.Get_lockKind(_env); _switch1 == Nodes_LockKind__AsyncLock || _switch1 == Nodes_LockKind__LuaLock {
            return true
        }
    }
    return false
}

// 124: DeclConstr
func (self *TransUnitIF_NSInfo) InitTransUnitIF_NSInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,pos *Types_Position,validAsyncCtrl bool) {
    self.idSetInfo = NewTransUnitIF_IdSetInfo(_env)
    self.nobody = false
    self.lockedAsyncStack = NewLnsList([]LnsAny{})
    self.loopScopeQueue = NewLnsList([]LnsAny{})
    self.typeInfo = typeInfo
    self.pos = pos
    self.validAsyncCtrl = validAsyncCtrl
}

// 137: decl @lune.@base.@TransUnitIF.NSInfo.incLock
func (self *TransUnitIF_NSInfo) IncLock(_env *LnsEnv, lockKind LnsInt) {
    self.lockedAsyncStack.Insert(TransUnitIF_LockedAsyncInfo2Stem(NewTransUnitIF_LockedAsyncInfo(_env, self.loopScopeQueue.Len(), lockKind)))
}

// 141: decl @lune.@base.@TransUnitIF.NSInfo.decLock
func (self *TransUnitIF_NSInfo) DecLock(_env *LnsEnv) {
    self.lockedAsyncStack.Remove(nil)
}

// 154: decl @lune.@base.@TransUnitIF.NSInfo.canBreak
func (self *TransUnitIF_NSInfo) CanBreak(_env *LnsEnv) bool {
    var len LnsInt
    len = self.lockedAsyncStack.Len()
    var loopQueueLen LnsInt
    loopQueueLen = self.loopScopeQueue.Len()
    if len == 0{
        return loopQueueLen > 0
    }
    return self.lockedAsyncStack.GetAt(len).(TransUnitIF_LockedAsyncInfoDownCast).ToTransUnitIF_LockedAsyncInfo().FP.Get_loopLen(_env) < loopQueueLen
}

// 166: decl @lune.@base.@TransUnitIF.NSInfo.canAccessNoasync
func (self *TransUnitIF_NSInfo) CanAccessNoasync(_env *LnsEnv) bool {
    var len LnsInt
    len = self.lockedAsyncStack.Len()
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.typeInfo.FP.Get_asyncMode(_env) == Ast_Async__Noasync) ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( len > 0) &&
            _env.SetStackVal( self.lockedAsyncStack.GetAt(len).(TransUnitIF_LockedAsyncInfoDownCast).ToTransUnitIF_LockedAsyncInfo().FP.Get_lockKind(_env) != Nodes_LockKind__Unsafe) ).(bool))) ).(bool){
        return true
    }
    return false
}

// 179: decl @lune.@base.@TransUnitIF.NSInfo.canAccessLuaval
func (self *TransUnitIF_NSInfo) CanAccessLuaval(_env *LnsEnv) bool {
    if Lns_op_not(self.validAsyncCtrl){
        return true
    }
    if self.lockedAsyncStack.Len() > 0{
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

// declaration Class -- ErrMess
type TransUnitIF_ErrMessMtd interface {
    Get_mess(_env *LnsEnv) string
    Get_pos(_env *LnsEnv) *Types_Position
}
type TransUnitIF_ErrMess struct {
    Mess string
    Pos *Types_Position
    FP TransUnitIF_ErrMessMtd
}
func TransUnitIF_ErrMess2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_ErrMess).FP
}
type TransUnitIF_ErrMessDownCast interface {
    ToTransUnitIF_ErrMess() *TransUnitIF_ErrMess
}
func TransUnitIF_ErrMessDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnitIF_ErrMessDownCast)
    if ok { return work.ToTransUnitIF_ErrMess() }
    return nil
}
func (obj *TransUnitIF_ErrMess) ToTransUnitIF_ErrMess() *TransUnitIF_ErrMess {
    return obj
}
func NewTransUnitIF_ErrMess(_env *LnsEnv, arg1 string, arg2 *Types_Position) *TransUnitIF_ErrMess {
    obj := &TransUnitIF_ErrMess{}
    obj.FP = obj
    obj.InitTransUnitIF_ErrMess(_env, arg1, arg2)
    return obj
}
func (self *TransUnitIF_ErrMess) InitTransUnitIF_ErrMess(_env *LnsEnv, arg1 string, arg2 *Types_Position) {
    self.Mess = arg1
    self.Pos = arg2
}
func (self *TransUnitIF_ErrMess) Get_mess(_env *LnsEnv) string{ return self.Mess }
func (self *TransUnitIF_ErrMess) Get_pos(_env *LnsEnv) *Types_Position{ return self.Pos }

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
    Ctrl_info *Types_TransCtrlInfo
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
// 269: decl @lune.@base.@TransUnitIF.TransUnitBase.get_scope
func (self *TransUnitIF_TransUnitBase) Get_scope(_env *LnsEnv) *Ast_Scope {
    return self.Scope
}

// 273: DeclConstr
func (self *TransUnitIF_TransUnitBase) InitTransUnitIF_TransUnitBase(_env *LnsEnv, ctrl_info *Types_TransCtrlInfo,processInfo *Ast_ProcessInfo) {
    self.Ctrl_info = ctrl_info
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
    self.NsInfoMap.Set(subRootTypeInfo,NewTransUnitIF_NSInfo(_env, subRootTypeInfo, NewTypes_Position(_env, 0, 0, "@builtin@"), ctrl_info.ValidAsyncCtrl))
}

// 292: decl @lune.@base.@TransUnitIF.TransUnitBase.addErrMess
func (self *TransUnitIF_TransUnitBase) AddErrMess(_env *LnsEnv, pos *Types_Position,mess string) {
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(mess,"type mismatch.*<- &", nil, nil))){
        mess = mess + ". if your code is the old style, use the opiton '--legacy-mutable-control'."
    }
    self.ErrMessList.Insert(TransUnitIF_ErrMess2Stem(NewTransUnitIF_ErrMess(_env, _env.GetVM().String_format("%s: error: %s", []LnsAny{pos.FP.GetDisplayTxt(_env), mess}), pos)))
}



// 321: decl @lune.@base.@TransUnitIF.TransUnitBase.error
func (self *TransUnitIF_TransUnitBase) Error(_env *LnsEnv, mess string) {
    self.FP.ErrorAt(_env, self.FP.GetLatestPos(_env), mess)
}

// 325: decl @lune.@base.@TransUnitIF.TransUnitBase.pushScope
func (self *TransUnitIF_TransUnitBase) PushScope(_env *LnsEnv, classFlag bool,baseInfo LnsAny,interfaceList LnsAny) *Ast_Scope {
    self.Scope = Ast_TypeInfo_createScope(_env, self.ProcessInfo, self.Scope, classFlag, baseInfo, interfaceList)
    return self.Scope
}

// 333: decl @lune.@base.@TransUnitIF.TransUnitBase.popScope
func (self *TransUnitIF_TransUnitBase) PopScope(_env *LnsEnv) {
    self.Scope = self.Scope.FP.Get_outerScope(_env)
}

// 341: decl @lune.@base.@TransUnitIF.TransUnitBase.newNSInfo
func (self *TransUnitIF_TransUnitBase) NewNSInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,pos *Types_Position) *TransUnitIF_NSInfo {
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = NewTransUnitIF_NSInfo(_env, typeInfo, pos, self.Ctrl_info.ValidAsyncCtrl)
    self.NsInfoMap.Set(typeInfo,nsInfo)
    return nsInfo
}

// 358: decl @lune.@base.@TransUnitIF.TransUnitBase.getCurrentNamespaceTypeInfoMut
func (self *TransUnitIF_TransUnitBase) GetCurrentNamespaceTypeInfoMut(_env *LnsEnv) *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = self.Scope.FP.GetNamespaceTypeInfo(_env)
    var nsInfo *TransUnitIF_NSInfo
    
    {
        _nsInfo := self.NsInfoMap.Get(typeInfo)
        if _nsInfo == nil{
            self.FP.Error(_env, _env.GetVM().String_format("not found nsInfo -- %s (%d)", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil), typeInfo.FP.Get_typeId(_env).Id}))
        } else {
            nsInfo = _nsInfo.(*TransUnitIF_NSInfo)
        }
    }
    return nsInfo.FP.Get_typeInfo(_env)
}

// 367: decl @lune.@base.@TransUnitIF.TransUnitBase.getCurrentNamespaceTypeInfo
func (self *TransUnitIF_TransUnitBase) GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.Scope.FP.GetNamespaceTypeInfo(_env)
}

// 371: decl @lune.@base.@TransUnitIF.TransUnitBase.pushModule
func (self *TransUnitIF_TransUnitBase) PushModule(_env *LnsEnv, processInfo *Ast_ProcessInfo,externalFlag bool,name string,mutable bool) *TransUnitIF_NSInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var modName string
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(name,"^@", nil, nil))){
        modName = name
    } else { 
        modName = _env.GetVM().String_format("@%s", []LnsAny{name})
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
                    self.FP.Error(_env, _env.GetVM().String_format("not found scope -- %s", []LnsAny{name}))
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
                existSym_321 := existSym.(*Ast_SymbolInfo)
                self.FP.AddErrMess(_env, self.FP.GetLatestPos(_env), _env.GetVM().String_format("module symbols exist -- %s.%s -- %s.%s", []LnsAny{existSym_321.FP.Get_namespaceTypeInfo(_env).FP.GetFullName(_env, self.TypeNameCtrl, parentScope.FP, false), existSym_321.FP.Get_name(_env), parentInfo.FP.GetFullName(_env, self.TypeNameCtrl, parentScope.FP, false), modName}))
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

// 424: decl @lune.@base.@TransUnitIF.TransUnitBase.pushModuleLow
func (self *TransUnitIF_TransUnitBase) PushModuleLow(_env *LnsEnv, processInfo *Ast_ProcessInfo,externalFlag bool,name string,mutable bool) *Ast_TypeInfo {
    return self.FP.PushModule(_env, processInfo, externalFlag, name, mutable).FP.Get_typeInfo(_env)
}

// 431: decl @lune.@base.@TransUnitIF.TransUnitBase.popModule
func (self *TransUnitIF_TransUnitBase) PopModule(_env *LnsEnv) {
    self.FP.PopScope(_env)
}

// 438: decl @lune.@base.@TransUnitIF.TransUnitBase.pushClassScope
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
        self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("This class does not exist in this scope. -- %s -- %s(%d), %s(%d)", []LnsAny{classTypeInfo.FP.GetTxt(_env, nil, nil, nil), _env.NilAccFin(_env.NilAccPush(self.Scope.FP.Get_ownerTypeInfo(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetFullName(_env, self.TypeNameCtrl, self.Scope.FP, false)})/* 456:15 */), _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(self.Scope.FP.Get_ownerTypeInfo(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_typeId(_env)})&&
            _env.NilAccPush(_env.NilAccPop().(*Ast_IdInfo).Id))) ||
            _env.SetStackVal( -1) ).(LnsInt), classParentName, classParentTypeId.Id}))
    }
    self.Scope = scope
}

// 470: decl @lune.@base.@TransUnitIF.TransUnitBase.pushClass
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
                self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("multiple class(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
                self.FP.Error(_env, "stop by error")
            }
            if typeInfo.FP.Get_abstractFlag(_env) != abstractFlag{
                self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) abstract for prototpye", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
            }
            if typeInfo.FP.Get_accessMode(_env) != accessMode{
                self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) accessmode(%s) for prototpye accessmode(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), Ast_AccessMode_getTxt( accessMode), Ast_AccessMode_getTxt( typeInfo.FP.Get_accessMode(_env))}))
            }
            if baseInfo != nil{
                baseInfo_342 := baseInfo.(*Ast_TypeInfo)
                if typeInfo.FP.Get_baseTypeInfo(_env) != baseInfo_342{
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) base class(%s) for prototpye base class(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), baseInfo_342.FP.GetTxt(_env, nil, nil, nil), typeInfo.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            } else {
                if typeInfo.FP.HasBase(_env){
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) base class(None) for prototpye base class(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), typeInfo.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
            var compareList func(_env *LnsEnv, protoList *LnsList,typeList *LnsList,message string)
            compareList = func(_env *LnsEnv, protoList *LnsList,typeList *LnsList,message string) {
                if protoList.Len() == typeList.Len(){
                    for _index, _protoType := range( protoList.Items ) {
                        index := _index + 1
                        protoType := _protoType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if protoType != typeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(){
                            self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) %s(%s) for prototpye %s(%s)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), message, typeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), message, protoType.FP.GetTxt(_env, nil, nil, nil)}))
                        }
                    }
                } else { 
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) %s(%d) for prototpye %s(%d)", []LnsAny{typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), message, typeList.Len(), message, protoList.Len()}))
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
                    self.FP.ErrorAt(_env, errPos, _env.GetVM().String_format("not find scope -- %s", []LnsAny{name}))
                }
            }
            if _switch1 := (typeInfo.FP.Get_kind(_env)); _switch1 == Ast_TypeInfoKind__Class {
                if mode == TransUnitIF_DeclClassMode__Interface{
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("define interface already -- %s", []LnsAny{name}))
                    Util_printStackTrace(_env)
                }
            } else if _switch1 == Ast_TypeInfoKind__IF {
                if mode != TransUnitIF_DeclClassMode__Interface{
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("define class already -- %s", []LnsAny{name}))
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
                genTypeList_371 := genTypeList.(*LnsList)
                workGenTypeList = genTypeList_371
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
        genTypeList_375 := genTypeList.(*LnsList)
        for _, _genType := range( genTypeList_375.Items ) {
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

// 622: decl @lune.@base.@TransUnitIF.TransUnitBase.pushClassLow
func (self *TransUnitIF_TransUnitBase) PushClassLow(_env *LnsEnv, processInfo *Ast_ProcessInfo,errPos *Types_Position,mode LnsInt,abstractFlag bool,baseInfo LnsAny,interfaceList LnsAny,genTypeList LnsAny,externalFlag bool,name string,allowMultiple bool,accessMode LnsInt,defNamespace LnsAny) *Ast_TypeInfo {
    return self.FP.PushClass(_env, processInfo, errPos, mode, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace).FP.Get_typeInfo(_env)
}

// 636: decl @lune.@base.@TransUnitIF.TransUnitBase.popClass
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
// 645: decl @lune.@base.@TransUnitIF.SimpeTransUnit.errorAt
func (self *TransUnitIF_SimpeTransUnit) ErrorAt(_env *LnsEnv, pos *Types_Position,mess string) {
    self.FP.AddErrMess(_env, pos, mess)
    for _, _errmess := range( self.ErrMessList.Items ) {
        errmess := _errmess.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
        Util_errorLog(_env, errmess.FP.Get_mess(_env))
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

// 659: decl @lune.@base.@TransUnitIF.SimpeTransUnit.getLatestPos
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
