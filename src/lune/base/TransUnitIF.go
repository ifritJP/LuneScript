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
    Set_nobody(_env *LnsEnv, arg1 bool)
}
type TransUnitIF_NSInfo struct {
    nobody bool
    typeInfo *Ast_TypeInfo
    pos *Types_Position
    loopScopeQueue *LnsList
    lockedAsyncStack *LnsList
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
// 67: decl @lune.@base.@TransUnitIF.NSInfo.isLockedAsync
func (self *TransUnitIF_NSInfo) IsLockedAsync(_env *LnsEnv) bool {
    return self.lockedAsyncStack.Len() > 0
}

// 71: DeclConstr
func (self *TransUnitIF_NSInfo) InitTransUnitIF_NSInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,pos *Types_Position) {
    self.nobody = false
    
    self.lockedAsyncStack = NewLnsList([]LnsAny{})
    
    self.loopScopeQueue = NewLnsList([]LnsAny{})
    
    self.typeInfo = typeInfo
    
    self.pos = pos
    
}

// 80: decl @lune.@base.@TransUnitIF.NSInfo.incLock
func (self *TransUnitIF_NSInfo) IncLock(_env *LnsEnv) {
    self.lockedAsyncStack.Insert(self.loopScopeQueue.Len())
}

// 83: decl @lune.@base.@TransUnitIF.NSInfo.decLock
func (self *TransUnitIF_NSInfo) DecLock(_env *LnsEnv) {
    self.lockedAsyncStack.Remove(nil)
}

// 96: decl @lune.@base.@TransUnitIF.NSInfo.canBreak
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

// 108: decl @lune.@base.@TransUnitIF.NSInfo.canAccessNoasync
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
        GetLatestPos(_env *LnsEnv) *Types_Position
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

func Lns_TransUnitIF_init(_env *LnsEnv) {
    if init_TransUnitIF { return }
    init_TransUnitIF = true
    TransUnitIF__mod__ = "@lune.@base.@TransUnitIF"
    Lns_InitMod()
    Lns_Parser_init(_env)
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Types_init(_env)
}
func init() {
    init_TransUnitIF = false
}
