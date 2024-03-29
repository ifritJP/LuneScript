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
// decl enum -- StmtKind 
type TransUnitIF_StmtKind = LnsInt
const TransUnitIF_StmtKind__Apply = 6
const TransUnitIF_StmtKind__For = 3
const TransUnitIF_StmtKind__Foreach = 4
const TransUnitIF_StmtKind__Forsort = 5
const TransUnitIF_StmtKind__Match = 2
const TransUnitIF_StmtKind__Switch = 1
var TransUnitIF_StmtKindList_ = NewLnsList( []LnsAny {
  TransUnitIF_StmtKind__Switch,
  TransUnitIF_StmtKind__Match,
  TransUnitIF_StmtKind__For,
  TransUnitIF_StmtKind__Foreach,
  TransUnitIF_StmtKind__Forsort,
  TransUnitIF_StmtKind__Apply,
})
func TransUnitIF_StmtKind_get__allList(_env *LnsEnv) *LnsList{
    return TransUnitIF_StmtKindList_
}
var TransUnitIF_StmtKindMap_ = map[LnsInt]string {
  TransUnitIF_StmtKind__Apply: "StmtKind.Apply",
  TransUnitIF_StmtKind__For: "StmtKind.For",
  TransUnitIF_StmtKind__Foreach: "StmtKind.Foreach",
  TransUnitIF_StmtKind__Forsort: "StmtKind.Forsort",
  TransUnitIF_StmtKind__Match: "StmtKind.Match",
  TransUnitIF_StmtKind__Switch: "StmtKind.Switch",
}
func TransUnitIF_StmtKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := TransUnitIF_StmtKindMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnitIF_StmtKind_getTxt(arg1 LnsInt) string {
    return TransUnitIF_StmtKindMap_[arg1];
}
// decl enum -- SetNSInfo 
type TransUnitIF_SetNSInfo = LnsInt
const TransUnitIF_SetNSInfo__FromScope = 1
const TransUnitIF_SetNSInfo__None = 0
var TransUnitIF_SetNSInfoList_ = NewLnsList( []LnsAny {
  TransUnitIF_SetNSInfo__None,
  TransUnitIF_SetNSInfo__FromScope,
})
func TransUnitIF_SetNSInfo_get__allList(_env *LnsEnv) *LnsList{
    return TransUnitIF_SetNSInfoList_
}
var TransUnitIF_SetNSInfoMap_ = map[LnsInt]string {
  TransUnitIF_SetNSInfo__FromScope: "SetNSInfo.FromScope",
  TransUnitIF_SetNSInfo__None: "SetNSInfo.None",
}
func TransUnitIF_SetNSInfo__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := TransUnitIF_SetNSInfoMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnitIF_SetNSInfo_getTxt(arg1 LnsInt) string {
    return TransUnitIF_SetNSInfoMap_[arg1];
}
// for 583
func TransUnitIF_convExp3_369(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 352: decl @lune.@base.@TransUnitIF.sortMess
func TransUnitIF_sortMess(_env *LnsEnv, list *LnsList2_[*TransUnitIF_ErrMess]) *LnsList2_[*TransUnitIF_ErrMess] {
    list.Sort(_env, LnsItemKindStem, LnsComp(func ( _env *LnsEnv, val1, val2 LnsAny ) bool {return TransUnitIF_sortMess___anonymous_0_( _env, val1.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess(), val2.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess())}))
    return list
}

func TransUnitIF_sortMess___anonymous_0_(_env *LnsEnv, mess1 *TransUnitIF_ErrMess,mess2 *TransUnitIF_ErrMess) bool {
    var pos1 Types_Position
    pos1 = mess1.FP.Get_pos(_env).Get_orgPos(_env)
    var pos2 Types_Position
    pos2 = mess2.FP.Get_pos(_env).Get_orgPos(_env)
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
// 46: decl @lune.@base.@TransUnitIF.Modifier.createModifier
func (self *TransUnitIF_Modifier) CreateModifier(_env *LnsEnv, typeInfo *Ast_TypeInfo,mutMode LnsInt) *Ast_TypeInfo {
    if Lns_op_not(self.validMutControl){
        return typeInfo
    }
    return self.processInfo.FP.CreateModifier(_env, typeInfo, mutMode)
}
// 67: decl @lune.@base.@TransUnitIF.IdSetInfo.registerSym
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
// 118: decl @lune.@base.@TransUnitIF.NSInfo.pushLoopScope
func (self *TransUnitIF_NSInfo) PushLoopScope(_env *LnsEnv, scope *Ast_Scope) {
    self.loopScopeQueue.Insert(NewTransUnitIF_LoopInfo(_env, scope))
}
// 121: decl @lune.@base.@TransUnitIF.NSInfo.popLoopScope
func (self *TransUnitIF_NSInfo) PopLoopScope(_env *LnsEnv) {
    self.loopScopeQueue.Remove(nil)
}
// 124: decl @lune.@base.@TransUnitIF.NSInfo.hasAsyncLockBreak
func (self *TransUnitIF_NSInfo) HasAsyncLockBreak(_env *LnsEnv) bool {
    if self.loopScopeQueue.Len() > 0{
        var loopInfo *TransUnitIF_LoopInfo
        loopInfo = self.loopScopeQueue.GetAt(self.loopScopeQueue.Len())
        return loopInfo.FP.Get_hasAsyncLockBreak(_env)
    }
    return false
}
// 155: decl @lune.@base.@TransUnitIF.NSInfo.addStmtNum
func (self *TransUnitIF_NSInfo) AddStmtNum(_env *LnsEnv, num LnsInt) {
    self.stmtNum = self.stmtNum + num
}
// 162: decl @lune.@base.@TransUnitIF.NSInfo.isLockedAsync
func (self *TransUnitIF_NSInfo) IsLockedAsync(_env *LnsEnv) bool {
    return self.lockedAsyncStack.Len() > 0
}
// 166: decl @lune.@base.@TransUnitIF.NSInfo.isNoasync
func (self *TransUnitIF_NSInfo) IsNoasync(_env *LnsEnv) bool {
    if self.typeInfo.FP.Get_asyncMode(_env) == Ast_Async__Noasync{
        return true
    }
    for _, _info := range( self.lockedAsyncStack.Items ) {
        info := _info
        if _switch0 := info.FP.Get_lockKind(_env); _switch0 == Nodes_LockKind__AsyncLock || _switch0 == Nodes_LockKind__LuaLock {
            return true
        }
    }
    return false
}
// 202: decl @lune.@base.@TransUnitIF.NSInfo.duplicate
func (self *TransUnitIF_NSInfo) Duplicate(_env *LnsEnv) *TransUnitIF_NSInfo {
    var typeData *Ast_TypeData
    typeData = NewAst_TypeData(_env)
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = NewTransUnitIF_NSInfo(_env, self.typeInfo, NewAst_SimpleTypeDataAccessor(_env, typeData).FP, self.pos, self.validAsyncCtrl)
    typeData.FP.AddFrom(_env, self.typeDataAccessor.Get_typeData(_env))
    return nsInfo
}
// 213: decl @lune.@base.@TransUnitIF.NSInfo.getNextStmtId
func (self *TransUnitIF_NSInfo) GetNextStmtId(_env *LnsEnv, stmtKind LnsInt) LnsInt {
    var id LnsInt
    id = self.stmtIdList.GetAt(stmtKind)
    self.stmtIdList.Set(stmtKind,id + 1)
    return id
}
// 220: decl @lune.@base.@TransUnitIF.NSInfo.incLock
func (self *TransUnitIF_NSInfo) IncLock(_env *LnsEnv, lockKind LnsInt) {
    self.lockedAsyncStack.Insert(NewTransUnitIF_LockedAsyncInfo(_env, self.loopScopeQueue.Len(), lockKind))
}
// 224: decl @lune.@base.@TransUnitIF.NSInfo.decLock
func (self *TransUnitIF_NSInfo) DecLock(_env *LnsEnv) {
    self.lockedAsyncStack.Remove(nil)
}
// 237: decl @lune.@base.@TransUnitIF.NSInfo.canBreak
func (self *TransUnitIF_NSInfo) CanBreak(_env *LnsEnv) bool {
    return self.loopScopeQueue.Len() > 0
}
// 260: decl @lune.@base.@TransUnitIF.NSInfo.setAsyncLockBreak
func (self *TransUnitIF_NSInfo) SetAsyncLockBreak(_env *LnsEnv) bool {
    var len LnsInt
    len = self.lockedAsyncStack.Len()
    if len == 0{
        return false
    }
    var loopQueueLen LnsInt
    loopQueueLen = self.loopScopeQueue.Len()
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( loopQueueLen > 0) &&
        _env.SetStackVal( self.lockedAsyncStack.GetAt(len).FP.Get_loopLen(_env) >= loopQueueLen) ).(bool)){
        var loopInfo *TransUnitIF_LoopInfo
        loopInfo = self.loopScopeQueue.GetAt(loopQueueLen)
        loopInfo.FP.Set_hasAsyncLockBreak(_env, true)
        return true
    }
    return false
}
// 278: decl @lune.@base.@TransUnitIF.NSInfo.canAccessNoasync
func (self *TransUnitIF_NSInfo) CanAccessNoasync(_env *LnsEnv) bool {
    var len LnsInt
    len = self.lockedAsyncStack.Len()
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.typeInfo.FP.Get_asyncMode(_env) == Ast_Async__Noasync) ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( len > 0) &&
            _env.SetStackVal( self.lockedAsyncStack.GetAt(len).FP.Get_lockKind(_env) != Nodes_LockKind__LuaDepend) ).(bool))) ).(bool){
        return true
    }
    return false
}
// 291: decl @lune.@base.@TransUnitIF.NSInfo.canAccessLuaval
func (self *TransUnitIF_NSInfo) CanAccessLuaval(_env *LnsEnv) bool {
    if Lns_op_not(self.validAsyncCtrl){
        return true
    }
    for _, _stack := range( self.lockedAsyncStack.Items ) {
        stack := _stack
        if stack.FP.Get_lockKind(_env) != Nodes_LockKind__AsyncLock{
            return true
        }
    }
    return false
}
// 303: decl @lune.@base.@TransUnitIF.NSInfo.addCondRet
func (self *TransUnitIF_NSInfo) AddCondRet(_env *LnsEnv, nodeManager *Nodes_NodeManager,pos Types_Position,inTestBlock bool,isInAnalyzeArgMode bool,expOkType *Ast_TypeInfo,exp *Nodes_Node,kind LnsInt) *Nodes_CondRetNode {
    self.condRetCount = self.condRetCount + 1
    if kind != Nodes_CondRetKind__Two{
        if exp.FP.Get_expTypeList(_env).Len() > 1{
            exp = &Nodes_ExpMultiTo1Node_create(_env, nodeManager, pos, inTestBlock, isInAnalyzeArgMode, NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](exp.FP.Get_expType(_env))), exp).Nodes_Node
        }
    }
    var condRetNode *Nodes_CondRetNode
    condRetNode = Nodes_CondRetNode_create(_env, nodeManager, pos, inTestBlock, isInAnalyzeArgMode, NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](expOkType)), exp, kind, self.condRetCount)
    self.condRetNodeList.Insert(condRetNode)
    return condRetNode
}
// 325: decl @lune.@base.@TransUnitIF.NSInfo.clearCondRetNodeList
func (self *TransUnitIF_NSInfo) ClearCondRetNodeList(_env *LnsEnv) {
    if self.condRetNodeList.Len() > 0{
        self.condRetNodeList = NewLnsList2_[*Nodes_CondRetNode]([]*Nodes_CondRetNode{})
    }
}
// 424: decl @lune.@base.@TransUnitIF.TransUnitBase.getCurrentNamespaceTypeInfo
func (self *TransUnitIF_TransUnitBase) GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.scope.FP.GetNamespaceTypeInfo(_env)
}
// 431: decl @lune.@base.@TransUnitIF.TransUnitBase.error
func (self *TransUnitIF_TransUnitBase) Error(_env *LnsEnv, mess string) {
    self.FP.ErrorAt(_env, self.FP.GetLatestPos(_env), mess)
}
// 435: decl @lune.@base.@TransUnitIF.TransUnitBase.get_curNsInfo
func (self *TransUnitIF_TransUnitBase) Get_curNsInfo(_env *LnsEnv) *TransUnitIF_NSInfo {
    {
        __exp := self.curNsInfo
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*TransUnitIF_NSInfo)
            return _exp
        }
    }
    self.FP.Error(_env, _env.GetVM().String_format("not found NSInfo for %s", Lns_2DDD(self.FP.GetCurrentNamespaceTypeInfo(_env))))
// insert a dummy
    return nil
}
// 441: decl @lune.@base.@TransUnitIF.TransUnitBase.set_curNsInfo
func (self *TransUnitIF_TransUnitBase) Set_curNsInfo(_env *LnsEnv, nsInfo *TransUnitIF_NSInfo) {
    self.curNsInfo = nsInfo
}
// 445: decl @lune.@base.@TransUnitIF.TransUnitBase.get_scope
func (self *TransUnitIF_TransUnitBase) Get_scope(_env *LnsEnv) *Ast_Scope {
    return self.scope
}
// 448: decl @lune.@base.@TransUnitIF.TransUnitBase.get_scopeRO
func (self *TransUnitIF_TransUnitBase) Get_scopeRO(_env *LnsEnv) *Ast_Scope {
    return self.scope
}
// 452: decl @lune.@base.@TransUnitIF.TransUnitBase.getNSInfo
func (self *TransUnitIF_TransUnitBase) GetNSInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo) *TransUnitIF_NSInfo {
    var nsInfo *TransUnitIF_NSInfo
    
    {
        _nsInfo := self.NsInfoMap.Get(typeInfo)
        if _nsInfo == nil{
            self.FP.Error(_env, _env.GetVM().String_format("not found TypeInfo -- %s", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            nsInfo = _nsInfo.(*TransUnitIF_NSInfo)
        }
    }
    return nsInfo
}
// 459: decl @lune.@base.@TransUnitIF.TransUnitBase.setScope
func (self *TransUnitIF_TransUnitBase) SetScope(_env *LnsEnv, scope *Ast_Scope,setNSInfo LnsInt) {
    self.scope = scope
    if _switch0 := setNSInfo; _switch0 == TransUnitIF_SetNSInfo__None {
    } else if _switch0 == TransUnitIF_SetNSInfo__FromScope {
        self.curNsInfo = self.FP.GetNSInfo(_env, self.FP.GetCurrentNamespaceTypeInfo(_env))
    }
}
// 491: decl @lune.@base.@TransUnitIF.TransUnitBase.addErrMess
func (self *TransUnitIF_TransUnitBase) AddErrMess(_env *LnsEnv, pos Types_Position,mess string) {
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(mess,"type mismatch.*<- &", nil, nil))){
        mess = mess + ". if your code is the old style, use the opiton '--legacy-mutable-control'."
    }
    self.ErrMessList.Insert(NewTransUnitIF_ErrMess(_env, _env.GetVM().String_format("%s: error: %s", Lns_2DDD(pos.GetDisplayTxt(_env), mess)), pos))
}
// 500: decl @lune.@base.@TransUnitIF.TransUnitBase.pushScope
func (self *TransUnitIF_TransUnitBase) PushScope(_env *LnsEnv, scopeKind LnsInt,baseInfo LnsAny,interfaceList LnsAny) *Ast_Scope {
    self.scope = Ast_TypeInfo_createScope(_env, self.ProcessInfo, self.scope, scopeKind, baseInfo, interfaceList)
    return self.scope
}
// 508: decl @lune.@base.@TransUnitIF.TransUnitBase.popScope
func (self *TransUnitIF_TransUnitBase) PopScope(_env *LnsEnv) {
    self.scope = self.scope.FP.Get_outerScope(_env)
    var nsInfo LnsAny
    nsInfo = self.NsInfoMap.Get(self.FP.GetCurrentNamespaceTypeInfo(_env))
    if nsInfo != self.curNsInfo{
        {
            _curNsInfo := self.curNsInfo
            if !Lns_IsNil( _curNsInfo ) {
                curNsInfo := _curNsInfo.(*TransUnitIF_NSInfo)
                if curNsInfo.FP.Get_condRetNodeList(_env).Len() != 0{
                    self.FP.Error(_env, "internal error. condRetNodeList is not nil.")
                }
            }
        }
        self.curNsInfo = nsInfo
    }
}
// 522: decl @lune.@base.@TransUnitIF.TransUnitBase.newNSInfoWithTypeData
func (self *TransUnitIF_TransUnitBase) NewNSInfoWithTypeData(_env *LnsEnv, typeInfo *Ast_TypeInfo,typeDataAccessor Ast_TypeDataAccessor,pos Types_Position) *TransUnitIF_NSInfo {
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = NewTransUnitIF_NSInfo(_env, typeInfo, typeDataAccessor, pos, self.Ctrl_info.ValidAsyncCtrl)
    self.NsInfoMap.Set(typeInfo,nsInfo)
    return nsInfo
}
// 536: decl @lune.@base.@TransUnitIF.TransUnitBase.newNSInfo
func (self *TransUnitIF_TransUnitBase) NewNSInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,pos Types_Position) *TransUnitIF_NSInfo {
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = NewTransUnitIF_NSInfo(_env, typeInfo, typeInfo.FP, pos, self.Ctrl_info.ValidAsyncCtrl)
    self.NsInfoMap.Set(typeInfo,nsInfo)
    return nsInfo
}
// 546: decl @lune.@base.@TransUnitIF.TransUnitBase.pushModule
func (self *TransUnitIF_TransUnitBase) PushModule(_env *LnsEnv, processInfo *Ast_ProcessInfo,externalFlag bool,name string,mutable bool) *TransUnitIF_NSInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var modName string
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(name,"^@", nil, nil))){
        modName = name
    } else { 
        modName = _env.GetVM().String_format("@%s", Lns_2DDD(name))
    }
    var nsInfo *TransUnitIF_NSInfo
    {
        __exp := self.scope.FP.GetTypeInfoChild(_env, modName)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            typeInfo = _exp
            {
                _scope := self.Namespace2Scope.Get(typeInfo)
                if !Lns_IsNil( _scope ) {
                    scope := _scope.(*Ast_Scope)
                    self.scope = scope
                } else {
                    self.FP.Error(_env, _env.GetVM().String_format("not found scope -- %s", Lns_2DDD(name)))
                }
            }
            nsInfo = Lns_unwrap( self.NsInfoMap.Get(typeInfo)).(*TransUnitIF_NSInfo)
        } else {
            var parentNsInfo *TransUnitIF_NSInfo
            parentNsInfo = self.FP.Get_curNsInfo(_env)
            var parentInfo *Ast_TypeInfo
            parentInfo = parentNsInfo.FP.Get_typeInfo(_env)
            var parentScope *Ast_Scope
            parentScope = self.scope
            var scope *Ast_Scope
            scope = self.FP.PushScope(_env, Ast_ScopeKind__Module, nil, nil)
            var newType *Ast_TypeInfo
            newType = processInfo.FP.CreateModule(_env, scope, parentInfo, parentNsInfo.FP.Get_typeDataAccessor(_env), externalFlag, modName, mutable)
            typeInfo = newType
            self.Namespace2Scope.Set(typeInfo,scope)
            nsInfo = self.FP.NewNSInfo(_env, newType, self.FP.GetLatestPos(_env))
            var existSym LnsAny
            _,existSym = parentScope.FP.AddClass(_env, processInfo, modName, nil, typeInfo)
            if existSym != nil{
                existSym_31 := existSym.(*Ast_SymbolInfo)
                self.FP.AddErrMess(_env, self.FP.GetLatestPos(_env), _env.GetVM().String_format("module symbols exist -- %s.%s -- %s.%s", Lns_2DDD(existSym_31.FP.Get_namespaceTypeInfo(_env).FP.GetFullName(_env, self.TypeNameCtrl, parentScope.FP, false), existSym_31.FP.Get_name(_env), parentInfo.FP.GetFullName(_env, self.TypeNameCtrl, parentScope.FP, false), modName)))
            }
        }
    }
    if Lns_op_not(self.TypeId2ClassMap.Get(typeInfo.FP.Get_typeId(_env))){
        var namespace *Nodes_NamespaceInfo
        namespace = NewNodes_NamespaceInfo(_env, modName, self.scope, typeInfo)
        self.TypeId2ClassMap.Set(typeInfo.FP.Get_typeId(_env),namespace)
    }
    self.curNsInfo = nsInfo
    return nsInfo
}
// 604: decl @lune.@base.@TransUnitIF.TransUnitBase.pushModuleLow
func (self *TransUnitIF_TransUnitBase) PushModuleLow(_env *LnsEnv, processInfo *Ast_ProcessInfo,externalFlag bool,name string,mutable bool) *Ast_TypeInfo {
    return self.FP.PushModule(_env, processInfo, externalFlag, name, mutable).FP.Get_typeInfo(_env)
}
// 611: decl @lune.@base.@TransUnitIF.TransUnitBase.popModule
func (self *TransUnitIF_TransUnitBase) PopModule(_env *LnsEnv) {
    self.FP.PopScope(_env)
}
// 618: decl @lune.@base.@TransUnitIF.TransUnitBase.pushClassScope
func (self *TransUnitIF_TransUnitBase) PushClassScope(_env *LnsEnv, errPos Types_Position,classTypeInfo *Ast_TypeInfo,scope *Ast_Scope) {
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
                classParentName = parentType.FP.GetFullName(_env, self.TypeNameCtrl, Lns_unwrap( parentType.FP.Get_scope(_env)).(*Ast_Scope).FP, false)
                classParentTypeId = parentType.FP.Get_typeId(_env)
            } else {
                classParentName = "nil"
                classParentTypeId = Ast_dummyIdInfo
            }
        }
        self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("This class does not exist in this scope. -- %s -- %s(%d), %s(%d)", Lns_2DDD(classTypeInfo.FP.GetTxt(_env, nil, nil, nil), _env.NilAccFin(_env.NilAccPush(self.scope.FP.Get_ownerTypeInfo(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetFullName(_env, self.TypeNameCtrl, self.scope.FP, false)})/* 636:15 */), _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(self.scope.FP.Get_ownerTypeInfo(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_typeId(_env)})&&
            _env.NilAccPush(_env.NilAccPop().(*Ast_IdInfo).Id))) ||
            _env.SetStackVal( -1) ).(LnsInt), classParentName, classParentTypeId.Id)))
    }
    self.scope = scope
    self.curNsInfo = self.NsInfoMap.Get(classTypeInfo)
}
// 651: decl @lune.@base.@TransUnitIF.TransUnitBase.pushClass
func (self *TransUnitIF_TransUnitBase) PushClass(_env *LnsEnv, processInfo *Ast_ProcessInfo,errPos Types_Position,mode LnsInt,finalFlag bool,abstractFlag bool,baseInfo LnsAny,interfaceList LnsAny,genTypeList LnsAny,externalFlag bool,name string,allowMultiple bool,accessMode LnsInt,defNamespace LnsAny) *TransUnitIF_NSInfo {
    var nsInfo *TransUnitIF_NSInfo
    var typeInfo *Ast_TypeInfo
    {
        __exp := self.scope.FP.GetTypeInfo(_env, name, self.scope, true, Ast_ScopeAccess__Normal)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            typeInfo = _exp
            nsInfo = Lns_unwrap( self.NsInfoMap.Get(typeInfo)).(*TransUnitIF_NSInfo)
            if _env.NilAccFin(_env.NilAccPush(typeInfo.FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_parent(_env)})) != self.scope{
                self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("multiple class(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil))))
                self.FP.Error(_env, "stop by error")
            }
            if typeInfo.FP.Get_abstractFlag(_env) != abstractFlag{
                self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) abstract for prototype", Lns_2DDD(typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil))))
            }
            if typeInfo.FP.Get_accessMode(_env) != accessMode{
                self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) accessmode(%s) for prototype accessmode(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), Ast_AccessMode_getTxt( accessMode), Ast_AccessMode_getTxt( typeInfo.FP.Get_accessMode(_env)))))
            }
            if baseInfo != nil{
                baseInfo_18 := baseInfo.(*Ast_TypeInfo)
                if typeInfo.FP.Get_baseTypeInfo(_env) != baseInfo_18{
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) base class(%s) for prototype base class(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), baseInfo_18.FP.GetTxt(_env, nil, nil, nil), typeInfo.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil))))
                }
            } else {
                if typeInfo.FP.HasBase(_env){
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) base class(None) for prototype base class(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), typeInfo.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil))))
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
                            self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) %s(%s) for prototype %s(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), "interface", typeList.GetAt(index).FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), "interface", protoType.FP.GetTxt(_env, nil, nil, nil))))
                        }
                    }
                } else { 
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) %s(%d) for prototype %s(%d)", Lns_2DDD(typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), "interface", typeList.Len(), "interface", typeInfo.FP.Get_interfaceList(_env).Len())))
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
                            self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) %s(%s) for prototype %s(%s)", Lns_2DDD(typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), "generics", typeList.GetAt(index).FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), "generics", protoType.FP.GetTxt(_env, nil, nil, nil))))
                        }
                    }
                } else { 
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("mismatch class(%s) %s(%d) for prototype %s(%d)", Lns_2DDD(typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), "generics", typeList.Len(), "generics", typeInfo.FP.Get_itemTypeInfoList(_env).Len())))
                }
            }
            
            {
                _scope := self.Namespace2Scope.Get(typeInfo)
                if !Lns_IsNil( _scope ) {
                    scope := _scope.(*Ast_Scope)
                    self.scope = scope
                } else {
                    self.FP.ErrorAt(_env, errPos, _env.GetVM().String_format("not find scope -- %s", Lns_2DDD(name)))
                }
            }
            if _switch0 := (typeInfo.FP.Get_kind(_env)); _switch0 == Ast_TypeInfoKind__Class {
                if mode == TransUnitIF_DeclClassMode__Interface{
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("define interface already -- %s", Lns_2DDD(name)))
                    Util_printStackTrace(_env)
                }
            } else if _switch0 == Ast_TypeInfoKind__IF {
                if mode != TransUnitIF_DeclClassMode__Interface{
                    self.FP.AddErrMess(_env, errPos, _env.GetVM().String_format("define class already -- %s", Lns_2DDD(name)))
                    Util_printStackTrace(_env)
                }
            }
        } else {
            var parentNsInfo *TransUnitIF_NSInfo
            parentNsInfo = self.FP.Get_curNsInfo(_env)
            var parentScope *Ast_Scope
            parentScope = self.scope
            var scope *Ast_Scope
            scope = self.FP.PushScope(_env, Ast_ScopeKind__Class, baseInfo, interfaceList)
            var workGenTypeList *LnsList2_[*Ast_AlternateTypeInfo]
            if genTypeList != nil{
                genTypeList_58 := genTypeList.(*LnsList2_[*Ast_AlternateTypeInfo])
                workGenTypeList = genTypeList_58
            } else {
                workGenTypeList = NewLnsList2_[*Ast_AlternateTypeInfo]([]*Ast_AlternateTypeInfo{})
            }
            var newType *Ast_TypeInfo
            newType = processInfo.FP.CreateClassAsync(_env, mode != TransUnitIF_DeclClassMode__Interface, finalFlag, abstractFlag, scope, baseInfo, interfaceList, workGenTypeList, parentNsInfo.FP.Get_typeInfo(_env), parentNsInfo.FP.Get_typeDataAccessor(_env), externalFlag, accessMode, name)
            typeInfo = newType
            self.Namespace2Scope.Set(typeInfo,scope)
            nsInfo = self.FP.NewNSInfo(_env, newType, errPos)
            parentScope.FP.AddClassLazy(_env, processInfo, name, errPos, typeInfo, mode == TransUnitIF_DeclClassMode__LazyModule)
        }
    }
    if genTypeList != nil{
        genTypeList_62 := genTypeList.(*LnsList2_[*Ast_AlternateTypeInfo])
        for _, _genType := range( genTypeList_62.Items ) {
            genType := _genType
            self.scope.FP.AddAlternate(_env, processInfo, accessMode, genType.FP.Get_txt(_env), errPos, &genType.Ast_TypeInfo)
        }
    }
    var namespace *Nodes_NamespaceInfo
    
    {
        _namespace := defNamespace
        if _namespace == nil{
            namespace = NewNodes_NamespaceInfo(_env, name, self.scope, typeInfo)
        } else {
            namespace = _namespace.(*Nodes_NamespaceInfo)
        }
    }
    self.TypeId2ClassMap.Set(typeInfo.FP.Get_typeId(_env),namespace)
    self.curNsInfo = nsInfo
    return nsInfo
}
// 811: decl @lune.@base.@TransUnitIF.TransUnitBase.pushClassLow
func (self *TransUnitIF_TransUnitBase) PushClassLow(_env *LnsEnv, processInfo *Ast_ProcessInfo,errPos Types_Position,mode LnsInt,finalFlag bool,abstractFlag bool,baseInfo LnsAny,interfaceList LnsAny,genTypeList LnsAny,externalFlag bool,name string,allowMultiple bool,accessMode LnsInt,defNamespace LnsAny) *Ast_TypeInfo {
    return self.FP.PushClass(_env, processInfo, errPos, mode, finalFlag, abstractFlag, baseInfo, interfaceList, genTypeList, externalFlag, name, allowMultiple, accessMode, defNamespace).FP.Get_typeInfo(_env)
}
// 825: decl @lune.@base.@TransUnitIF.TransUnitBase.popClass
func (self *TransUnitIF_TransUnitBase) PopClass(_env *LnsEnv) {
    self.FP.PopScope(_env)
}
// 834: decl @lune.@base.@TransUnitIF.SimpeTransUnit.errorAt
func (self *TransUnitIF_SimpeTransUnit) ErrorAt(_env *LnsEnv, pos Types_Position,mess string) {
    self.FP.AddErrMess(_env, pos, mess)
    for _, _errmess := range( self.ErrMessList.Items ) {
        errmess := _errmess
        Util_errorLog(_env, errmess.FP.Get_mess(_env))
    }
    {
        _nearCode := self.nearCode
        if !Lns_IsNil( _nearCode ) {
            nearCode := _nearCode.(string)
            Util_println(_env, Lns_2DDD("------ near code -----", self.macroMode))
            Util_println(_env, Lns_2DDD(nearCode))
            Util_println(_env, Lns_2DDD("------"))
        }
    }
    Util_err(_env, "has error")
}
// 848: decl @lune.@base.@TransUnitIF.SimpeTransUnit.getLatestPos
func (self *TransUnitIF_SimpeTransUnit) GetLatestPos(_env *LnsEnv) Types_Position {
    return self.latestPos
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
func TransUnitIF_Modifier_toSlice(slice []LnsAny) []*TransUnitIF_Modifier {
    ret := make([]*TransUnitIF_Modifier, len(slice))
    for index, val := range slice {
        ret[index] = val.(TransUnitIF_ModifierDownCast).ToTransUnitIF_Modifier()
    }
    return ret
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
func TransUnitIF_IdSetInfo_toSlice(slice []LnsAny) []*TransUnitIF_IdSetInfo {
    ret := make([]*TransUnitIF_IdSetInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(TransUnitIF_IdSetInfoDownCast).ToTransUnitIF_IdSetInfo()
    }
    return ret
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
// 62: DeclConstr
func (self *TransUnitIF_IdSetInfo) InitTransUnitIF_IdSetInfo(_env *LnsEnv) {
    self.anonymousFuncId = NewAst_IdProvider(_env, 0, 10000)
    self.anonymousVarId = NewAst_IdProvider(_env, 0, 10000)
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
func TransUnitIF_LockedAsyncInfo_toSlice(slice []LnsAny) []*TransUnitIF_LockedAsyncInfo {
    ret := make([]*TransUnitIF_LockedAsyncInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(TransUnitIF_LockedAsyncInfoDownCast).ToTransUnitIF_LockedAsyncInfo()
    }
    return ret
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

// declaration Class -- LoopInfo
type TransUnitIF_LoopInfoMtd interface {
    Get_hasAsyncLockBreak(_env *LnsEnv) bool
    Get_scope(_env *LnsEnv) *Ast_Scope
    Set_hasAsyncLockBreak(_env *LnsEnv, arg1 bool)
}
type TransUnitIF_LoopInfo struct {
    scope *Ast_Scope
    hasAsyncLockBreak bool
    FP TransUnitIF_LoopInfoMtd
}
func TransUnitIF_LoopInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_LoopInfo).FP
}
func TransUnitIF_LoopInfo_toSlice(slice []LnsAny) []*TransUnitIF_LoopInfo {
    ret := make([]*TransUnitIF_LoopInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(TransUnitIF_LoopInfoDownCast).ToTransUnitIF_LoopInfo()
    }
    return ret
}
type TransUnitIF_LoopInfoDownCast interface {
    ToTransUnitIF_LoopInfo() *TransUnitIF_LoopInfo
}
func TransUnitIF_LoopInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnitIF_LoopInfoDownCast)
    if ok { return work.ToTransUnitIF_LoopInfo() }
    return nil
}
func (obj *TransUnitIF_LoopInfo) ToTransUnitIF_LoopInfo() *TransUnitIF_LoopInfo {
    return obj
}
func NewTransUnitIF_LoopInfo(_env *LnsEnv, arg1 *Ast_Scope) *TransUnitIF_LoopInfo {
    obj := &TransUnitIF_LoopInfo{}
    obj.FP = obj
    obj.InitTransUnitIF_LoopInfo(_env, arg1)
    return obj
}
func (self *TransUnitIF_LoopInfo) Get_scope(_env *LnsEnv) *Ast_Scope{ return self.scope }
func (self *TransUnitIF_LoopInfo) Get_hasAsyncLockBreak(_env *LnsEnv) bool{ return self.hasAsyncLockBreak }
func (self *TransUnitIF_LoopInfo) Set_hasAsyncLockBreak(_env *LnsEnv, arg1 bool){ self.hasAsyncLockBreak = arg1 }
// 101: DeclConstr
func (self *TransUnitIF_LoopInfo) InitTransUnitIF_LoopInfo(_env *LnsEnv, scope *Ast_Scope) {
    self.scope = scope
    self.hasAsyncLockBreak = false
}


// declaration Class -- NSInfo
type TransUnitIF_NSInfoMtd interface {
    AddCondRet(_env *LnsEnv, arg1 *Nodes_NodeManager, arg2 Types_Position, arg3 bool, arg4 bool, arg5 *Ast_TypeInfo, arg6 *Nodes_Node, arg7 LnsInt) *Nodes_CondRetNode
    AddStmtNum(_env *LnsEnv, arg1 LnsInt)
    CanAccessLuaval(_env *LnsEnv) bool
    CanAccessNoasync(_env *LnsEnv) bool
    CanBreak(_env *LnsEnv) bool
    ClearCondRetNodeList(_env *LnsEnv)
    DecLock(_env *LnsEnv)
    Duplicate(_env *LnsEnv) *TransUnitIF_NSInfo
    GetNextStmtId(_env *LnsEnv, arg1 LnsInt) LnsInt
    Get_condRetNodeList(_env *LnsEnv) *LnsList2_[*Nodes_CondRetNode]
    Get_nobody(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_stmtNum(_env *LnsEnv) LnsInt
    Get_typeDataAccessor(_env *LnsEnv) Ast_TypeDataAccessor
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    HasAsyncLockBreak(_env *LnsEnv) bool
    IncLock(_env *LnsEnv, arg1 LnsInt)
    IsLockedAsync(_env *LnsEnv) bool
    IsNoasync(_env *LnsEnv) bool
    PopLoopScope(_env *LnsEnv)
    PushLoopScope(_env *LnsEnv, arg1 *Ast_Scope)
    RegisterSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) *Ast_SymbolInfo
    SetAsyncLockBreak(_env *LnsEnv) bool
    Set_nobody(_env *LnsEnv, arg1 bool)
}
type TransUnitIF_NSInfo struct {
    nobody bool
    typeInfo *Ast_TypeInfo
    typeDataAccessor Ast_TypeDataAccessor
    pos Types_Position
    loopScopeQueue *LnsList2_[*TransUnitIF_LoopInfo]
    lockedAsyncStack *LnsList2_[*TransUnitIF_LockedAsyncInfo]
    idSetInfo *TransUnitIF_IdSetInfo
    condRetNodeList *LnsList2_[*Nodes_CondRetNode]
    condRetCount LnsInt
    validAsyncCtrl bool
    stmtIdList *LnsList2_[LnsInt]
    stmtNum LnsInt
    FP TransUnitIF_NSInfoMtd
}
func TransUnitIF_NSInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_NSInfo).FP
}
func TransUnitIF_NSInfo_toSlice(slice []LnsAny) []*TransUnitIF_NSInfo {
    ret := make([]*TransUnitIF_NSInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(TransUnitIF_NSInfoDownCast).ToTransUnitIF_NSInfo()
    }
    return ret
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
func NewTransUnitIF_NSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Ast_TypeDataAccessor, arg3 Types_Position, arg4 bool) *TransUnitIF_NSInfo {
    obj := &TransUnitIF_NSInfo{}
    obj.FP = obj
    obj.InitTransUnitIF_NSInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *TransUnitIF_NSInfo) Get_nobody(_env *LnsEnv) bool{ return self.nobody }
func (self *TransUnitIF_NSInfo) Set_nobody(_env *LnsEnv, arg1 bool){ self.nobody = arg1 }
func (self *TransUnitIF_NSInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *TransUnitIF_NSInfo) Get_typeDataAccessor(_env *LnsEnv) Ast_TypeDataAccessor{ return self.typeDataAccessor }
func (self *TransUnitIF_NSInfo) Get_pos(_env *LnsEnv) Types_Position{ return self.pos }
func (self *TransUnitIF_NSInfo) Get_condRetNodeList(_env *LnsEnv) *LnsList2_[*Nodes_CondRetNode]{ return self.condRetNodeList }
func (self *TransUnitIF_NSInfo) Get_stmtNum(_env *LnsEnv) LnsInt{ return self.stmtNum }
// advertise -- 107
func (self *TransUnitIF_NSInfo) RegisterSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) *Ast_SymbolInfo {
    return self.idSetInfo. FP.RegisterSym( _env, arg1)
}
// 180: DeclConstr
func (self *TransUnitIF_NSInfo) InitTransUnitIF_NSInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,typeDataAccessor Ast_TypeDataAccessor,pos Types_Position,validAsyncCtrl bool) {
    self.stmtIdList = NewLnsList2_[LnsInt]([]LnsInt{})
    for range( TransUnitIF_StmtKind_get__allList(_env).Items ) {
        self.stmtIdList.Insert(LnsInt(0))
    }
    self.idSetInfo = NewTransUnitIF_IdSetInfo(_env)
    self.nobody = false
    self.lockedAsyncStack = NewLnsList2_[*TransUnitIF_LockedAsyncInfo]([]*TransUnitIF_LockedAsyncInfo{})
    self.loopScopeQueue = NewLnsList2_[*TransUnitIF_LoopInfo]([]*TransUnitIF_LoopInfo{})
    self.typeDataAccessor = typeDataAccessor
    self.typeInfo = typeInfo
    self.pos = pos
    self.validAsyncCtrl = validAsyncCtrl
    self.stmtNum = 0
    self.condRetNodeList = NewLnsList2_[*Nodes_CondRetNode]([]*Nodes_CondRetNode{})
    self.condRetCount = 0
}


type TransUnitIF_TransUnitIF interface {
        Error(_env *LnsEnv, arg1 string)
        Get_scope(_env *LnsEnv) *Ast_Scope
        PopClass(_env *LnsEnv)
        PopModule(_env *LnsEnv)
        PushClassScope(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
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
    Get_pos(_env *LnsEnv) Types_Position
}
type TransUnitIF_ErrMess struct {
    Mess string
    Pos Types_Position
    FP TransUnitIF_ErrMessMtd
}
func TransUnitIF_ErrMess2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_ErrMess).FP
}
func TransUnitIF_ErrMess_toSlice(slice []LnsAny) []*TransUnitIF_ErrMess {
    ret := make([]*TransUnitIF_ErrMess, len(slice))
    for index, val := range slice {
        ret[index] = val.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
    }
    return ret
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
func NewTransUnitIF_ErrMess(_env *LnsEnv, arg1 string, arg2 Types_Position) *TransUnitIF_ErrMess {
    obj := &TransUnitIF_ErrMess{}
    obj.FP = obj
    obj.InitTransUnitIF_ErrMess(_env, arg1, arg2)
    return obj
}
func (self *TransUnitIF_ErrMess) InitTransUnitIF_ErrMess(_env *LnsEnv, arg1 string, arg2 Types_Position) {
    self.Mess = arg1
    self.Pos = arg2
}
func (self *TransUnitIF_ErrMess) Get_mess(_env *LnsEnv) string{ return self.Mess }
func (self *TransUnitIF_ErrMess) Get_pos(_env *LnsEnv) Types_Position{ return self.Pos }

// declaration Class -- TransUnitBase
type TransUnitIF_TransUnitBaseMtd interface {
    AddErrMess(_env *LnsEnv, arg1 Types_Position, arg2 string)
    Error(_env *LnsEnv, arg1 string)
    ErrorAt(_env *LnsEnv, arg1 Types_Position, arg2 string)
    GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    GetLatestPos(_env *LnsEnv) Types_Position
    GetNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) *TransUnitIF_NSInfo
    Get_curNsInfo(_env *LnsEnv) *TransUnitIF_NSInfo
    Get_errMessList(_env *LnsEnv) *LnsList2_[*TransUnitIF_ErrMess]
    Get_globalScope(_env *LnsEnv) *Ast_Scope
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_scopeRO(_env *LnsEnv) *Ast_Scope
    NewNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Types_Position) *TransUnitIF_NSInfo
    NewNSInfoWithTypeData(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Ast_TypeDataAccessor, arg3 Types_Position) *TransUnitIF_NSInfo
    PopClass(_env *LnsEnv)
    PopModule(_env *LnsEnv)
    PopScope(_env *LnsEnv)
    PushClass(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsAny, arg9 bool, arg10 string, arg11 bool, arg12 LnsInt, arg13 LnsAny) *TransUnitIF_NSInfo
    PushClassLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsAny, arg9 bool, arg10 string, arg11 bool, arg12 LnsInt, arg13 LnsAny) *Ast_TypeInfo
    PushClassScope(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
    PushModule(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *TransUnitIF_NSInfo
    PushModuleLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *Ast_TypeInfo
    PushScope(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny) *Ast_Scope
    SetScope(_env *LnsEnv, arg1 *Ast_Scope, arg2 LnsInt)
    Set_curNsInfo(_env *LnsEnv, arg1 *TransUnitIF_NSInfo)
}
type TransUnitIF_TransUnitBase struct {
    ProcessInfo *Ast_ProcessInfo
    scope *Ast_Scope
    GlobalScope *Ast_Scope
    Namespace2Scope *LnsMap2_[*Ast_TypeInfo,*Ast_Scope]
    NsInfoMap *LnsMap2_[*Ast_TypeInfo,*TransUnitIF_NSInfo]
    curNsInfo LnsAny
    ErrMessList *LnsList2_[*TransUnitIF_ErrMess]
    TypeNameCtrl *Ast_TypeNameCtrl
    TypeId2ClassMap *LnsMap2_[*Ast_IdInfo,*Nodes_NamespaceInfo]
    Ctrl_info *Types_TransCtrlInfo
    FP TransUnitIF_TransUnitBaseMtd
}
func TransUnitIF_TransUnitBase2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnitIF_TransUnitBase).FP
}
func TransUnitIF_TransUnitBase_toSlice(slice []LnsAny) []*TransUnitIF_TransUnitBase {
    ret := make([]*TransUnitIF_TransUnitBase, len(slice))
    for index, val := range slice {
        ret[index] = val.(TransUnitIF_TransUnitBaseDownCast).ToTransUnitIF_TransUnitBase()
    }
    return ret
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
func (self *TransUnitIF_TransUnitBase) Get_errMessList(_env *LnsEnv) *LnsList2_[*TransUnitIF_ErrMess]{ return self.ErrMessList }


// 470: DeclConstr
func (self *TransUnitIF_TransUnitBase) InitTransUnitIF_TransUnitBase(_env *LnsEnv, ctrl_info *Types_TransCtrlInfo,processInfo *Ast_ProcessInfo) {
    self.Ctrl_info = ctrl_info
    self.TypeId2ClassMap = NewLnsMap2_[*Ast_IdInfo,*Nodes_NamespaceInfo]( map[*Ast_IdInfo]*Nodes_NamespaceInfo{})
    self.TypeNameCtrl = Ast_defaultTypeNameCtrl
    self.ErrMessList = NewLnsList2_[*TransUnitIF_ErrMess]([]*TransUnitIF_ErrMess{})
    self.Namespace2Scope = NewLnsMap2_[*Ast_TypeInfo,*Ast_Scope]( map[*Ast_TypeInfo]*Ast_Scope{})
    self.ProcessInfo = processInfo
    self.GlobalScope = NewAst_Scope(_env, processInfo, processInfo.FP.Get_topScope(_env), Ast_ScopeKind__Module, nil, nil)
    self.scope = NewAst_Scope(_env, processInfo, self.GlobalScope, Ast_ScopeKind__Module, nil, nil)
    self.NsInfoMap = NewLnsMap2_[*Ast_TypeInfo,*TransUnitIF_NSInfo]( map[*Ast_TypeInfo]*TransUnitIF_NSInfo{})
    var subRootTypeInfo *Ast_TypeInfo
    subRootTypeInfo = self.ProcessInfo.FP.Get_dummyParentType(_env)
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = NewTransUnitIF_NSInfo(_env, subRootTypeInfo, subRootTypeInfo.FP, NewTypes_Position(_env, 0, 0, "@builtin@", nil), ctrl_info.ValidAsyncCtrl)
    self.curNsInfo = nsInfo
    self.NsInfoMap.Set(subRootTypeInfo,nsInfo)
}


// declaration Class -- SimpeTransUnit
type TransUnitIF_SimpeTransUnitMtd interface {
    AddErrMess(_env *LnsEnv, arg1 Types_Position, arg2 string)
    Error(_env *LnsEnv, arg1 string)
    ErrorAt(_env *LnsEnv, arg1 Types_Position, arg2 string)
    GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    GetLatestPos(_env *LnsEnv) Types_Position
    GetNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) *TransUnitIF_NSInfo
    Get_curNsInfo(_env *LnsEnv) *TransUnitIF_NSInfo
    Get_errMessList(_env *LnsEnv) *LnsList2_[*TransUnitIF_ErrMess]
    Get_globalScope(_env *LnsEnv) *Ast_Scope
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_scopeRO(_env *LnsEnv) *Ast_Scope
    NewNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Types_Position) *TransUnitIF_NSInfo
    NewNSInfoWithTypeData(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Ast_TypeDataAccessor, arg3 Types_Position) *TransUnitIF_NSInfo
    PopClass(_env *LnsEnv)
    PopModule(_env *LnsEnv)
    PopScope(_env *LnsEnv)
    PushClass(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsAny, arg9 bool, arg10 string, arg11 bool, arg12 LnsInt, arg13 LnsAny) *TransUnitIF_NSInfo
    PushClassLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsAny, arg9 bool, arg10 string, arg11 bool, arg12 LnsInt, arg13 LnsAny) *Ast_TypeInfo
    PushClassScope(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
    PushModule(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *TransUnitIF_NSInfo
    PushModuleLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *Ast_TypeInfo
    PushScope(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny) *Ast_Scope
    SetScope(_env *LnsEnv, arg1 *Ast_Scope, arg2 LnsInt)
    Set_curNsInfo(_env *LnsEnv, arg1 *TransUnitIF_NSInfo)
}
type TransUnitIF_SimpeTransUnit struct {
    TransUnitIF_TransUnitBase
    latestPos Types_Position
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
func TransUnitIF_SimpeTransUnit_toSlice(slice []LnsAny) []*TransUnitIF_SimpeTransUnit {
    ret := make([]*TransUnitIF_SimpeTransUnit, len(slice))
    for index, val := range slice {
        ret[index] = val.(TransUnitIF_SimpeTransUnitDownCast).ToTransUnitIF_SimpeTransUnit()
    }
    return ret
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
func NewTransUnitIF_SimpeTransUnit(_env *LnsEnv, arg1 *Types_TransCtrlInfo, arg2 *Ast_ProcessInfo, arg3 Types_Position, arg4 string, arg5 LnsAny) *TransUnitIF_SimpeTransUnit {
    obj := &TransUnitIF_SimpeTransUnit{}
    obj.FP = obj
    obj.TransUnitIF_TransUnitBase.FP = obj
    obj.InitTransUnitIF_SimpeTransUnit(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *TransUnitIF_SimpeTransUnit) InitTransUnitIF_SimpeTransUnit(_env *LnsEnv, arg1 *Types_TransCtrlInfo, arg2 *Ast_ProcessInfo, arg3 Types_Position, arg4 string, arg5 LnsAny) {
    self.TransUnitIF_TransUnitBase.InitTransUnitIF_TransUnitBase( _env, arg1,arg2)
    self.latestPos = arg3
    self.macroMode = arg4
    self.nearCode = arg5
}

func Lns_TransUnitIF_init(_env *LnsEnv) {
    if init_TransUnitIF { return }
    init_TransUnitIF = true
    TransUnitIF__mod__ = "@lune.@base.@TransUnitIF"
    Lns_InitMod()
    Lns_Tokenizer_init(_env)
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Util_init(_env)
    Lns_Types_init(_env)
}
func init() {
    init_TransUnitIF = false
}
