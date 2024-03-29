// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_convLua bool
var convLua__mod__ string
// decl enum -- ConvMode 
type ConvLua_ConvMode = LnsInt
const ConvLua_ConvMode__ConvMeta = 1
const ConvLua_ConvMode__Convert = 0
var ConvLua_ConvModeList_ = NewLnsList( []LnsAny {
  ConvLua_ConvMode__Convert,
  ConvLua_ConvMode__ConvMeta,
})
func ConvLua_ConvMode_get__allList(_env *LnsEnv) *LnsList{
    return ConvLua_ConvModeList_
}
var ConvLua_ConvModeMap_ = map[LnsInt]string {
  ConvLua_ConvMode__ConvMeta: "ConvMode.ConvMeta",
  ConvLua_ConvMode__Convert: "ConvMode.Convert",
}
func ConvLua_ConvMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := ConvLua_ConvModeMap_[arg1]; ok { return arg1 }
    return nil
}

func ConvLua_ConvMode_getTxt(arg1 LnsInt) string {
    return ConvLua_ConvModeMap_[arg1];
}
// decl enum -- ExportIdKind 
type convLua_ExportIdKind = LnsInt
const convLua_ExportIdKind__Depend = 1
const convLua_ExportIdKind__Discarded = 0
const convLua_ExportIdKind__Normal = 2
var convLua_ExportIdKindList_ = NewLnsList( []LnsAny {
  convLua_ExportIdKind__Discarded,
  convLua_ExportIdKind__Depend,
  convLua_ExportIdKind__Normal,
})
func convLua_ExportIdKind_get__allList_2_(_env *LnsEnv) *LnsList{
    return convLua_ExportIdKindList_
}
var convLua_ExportIdKindMap_ = map[LnsInt]string {
  convLua_ExportIdKind__Depend: "ExportIdKind.Depend",
  convLua_ExportIdKind__Discarded: "ExportIdKind.Discarded",
  convLua_ExportIdKind__Normal: "ExportIdKind.Normal",
}
func convLua_ExportIdKind__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := convLua_ExportIdKindMap_[arg1]; ok { return arg1 }
    return nil
}

func convLua_ExportIdKind_getTxt(arg1 LnsInt) string {
    return convLua_ExportIdKindMap_[arg1];
}
var convLua_stepIndent LnsInt
var convLua_asyncLockBreakSym string
type convLua_outputMacroStmtBlock_13_ func (_env *LnsEnv)
// for 4117
func convLua_convExp0_4806(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 3869
func convLua_convExp0_3743(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 1570
func convLua_convExp2_2343(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1571
func convLua_convExp2_2358(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1581
func convLua_convExp2_2420(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1589
func convLua_convExp2_2471(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1657
func convLua_convExp2_2840(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1719
func convLua_convExp2_3113(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 2051
func convLua_convExp2_4717(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 3234
func convLua_convExp4_2135(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 3249
func convLua_convExp4_2211(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 4040
func convLua_convExp0_4334(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 4115
func convLua_convExp0_4554(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}

// 79: decl @lune.@base.@convLua.getSymbolTxt
func convLua_getSymbolTxt_5_(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) string {
    if symbolInfo.FP.Get_name(_env) == "_"{
        {
            _annonymous := Ast_AnonymousSymbolInfoDownCastF(symbolInfo.FP)
            if !Lns_IsNil( _annonymous ) {
                annonymous := _annonymous.(*Ast_AnonymousSymbolInfo)
                return _env.GetVM().String_format("_%d", Lns_2DDD(annonymous.FP.Get_anonymousId(_env)))
            } else {
                Util_err(_env, _env.GetVM().String_format("can't cast to AnonymousSymbolInfo. -- %s:%s", Lns_2DDD(_env.NilAccFin(_env.NilAccPush(symbolInfo.FP.Get_pos(_env)) && 
                _env.NilAccPush(_env.NilAccPop().(Types_Position).LineNo)), _env.NilAccFin(_env.NilAccPush(symbolInfo.FP.Get_pos(_env)) && 
                _env.NilAccPush(_env.NilAccPop().(Types_Position).Column)))))
            }
        }
    }
    return symbolInfo.FP.Get_name(_env)
}

// 279: decl @lune.@base.@convLua.filter
func convLua_filter_9_(_env *LnsEnv, node *Nodes_Node,filter *convLua_ConvFilter,parent *Nodes_Node) {
    node.FP.ProcessFilter(_env, &filter.Nodes_Filter, ConvLua_Opt2Stem(NewConvLua_Opt(_env, parent)))
}






// 1498: decl @lune.@base.@convLua.isGenericType
func convLua_isGenericType_12_(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if Ast_isGenericType(_env, typeInfo){
        return true
    }
    if _switch0 := typeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Class || _switch0 == Ast_TypeInfoKind__IF {
        if typeInfo.FP.Get_itemTypeInfoList(_env).Len() > 0{
            return true
        }
    }
    return false
}






// 3952: decl @lune.@base.@convLua.createFilter
func ConvLua_createFilter(_env *LnsEnv, streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,moduleTypeInfo *Ast_TypeInfo,processInfo *Ast_ProcessInfo,moduleSymbolKind LnsInt,builtinFunc *Builtin_BuiltinFuncType,useLuneRuntime LnsAny,targetLuaVer *LuaVer_LuaVerInfo,enableTest bool,useIpairs bool,option *ConvLua_Option) *ConvLua_FilterInfo {
    var convFilter *convLua_ConvFilter
    convFilter = NewconvLua_ConvFilter(_env, streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, processInfo, moduleSymbolKind, builtinFunc, useLuneRuntime, targetLuaVer, enableTest, useIpairs, option)
    return NewConvLua_FilterInfo(_env, &convFilter.Nodes_Filter)
}








// declaration Class -- PubVerInfo
type convLua_PubVerInfoMtd interface {
}
type convLua_PubVerInfo struct {
    StaticFlag bool
    AccessMode LnsInt
    Mutable bool
    TypeInfo *Ast_TypeInfo
    FP convLua_PubVerInfoMtd
}
func convLua_PubVerInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convLua_PubVerInfo).FP
}
func convLua_PubVerInfo_toSlice(slice []LnsAny) []*convLua_PubVerInfo {
    ret := make([]*convLua_PubVerInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(convLua_PubVerInfoDownCast).ToconvLua_PubVerInfo()
    }
    return ret
}
type convLua_PubVerInfoDownCast interface {
    ToconvLua_PubVerInfo() *convLua_PubVerInfo
}
func convLua_PubVerInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convLua_PubVerInfoDownCast)
    if ok { return work.ToconvLua_PubVerInfo() }
    return nil
}
func (obj *convLua_PubVerInfo) ToconvLua_PubVerInfo() *convLua_PubVerInfo {
    return obj
}
func NewconvLua_PubVerInfo(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 bool, arg4 *Ast_TypeInfo) *convLua_PubVerInfo {
    obj := &convLua_PubVerInfo{}
    obj.FP = obj
    obj.InitconvLua_PubVerInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *convLua_PubVerInfo) InitconvLua_PubVerInfo(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 bool, arg4 *Ast_TypeInfo) {
    self.StaticFlag = arg1
    self.AccessMode = arg2
    self.Mutable = arg3
    self.TypeInfo = arg4
}

// declaration Class -- PubFuncInfo
type convLua_PubFuncInfoMtd interface {
}
type convLua_PubFuncInfo struct {
    AccessMode LnsInt
    TypeInfo *Ast_TypeInfo
    FP convLua_PubFuncInfoMtd
}
func convLua_PubFuncInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convLua_PubFuncInfo).FP
}
func convLua_PubFuncInfo_toSlice(slice []LnsAny) []*convLua_PubFuncInfo {
    ret := make([]*convLua_PubFuncInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(convLua_PubFuncInfoDownCast).ToconvLua_PubFuncInfo()
    }
    return ret
}
type convLua_PubFuncInfoDownCast interface {
    ToconvLua_PubFuncInfo() *convLua_PubFuncInfo
}
func convLua_PubFuncInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convLua_PubFuncInfoDownCast)
    if ok { return work.ToconvLua_PubFuncInfo() }
    return nil
}
func (obj *convLua_PubFuncInfo) ToconvLua_PubFuncInfo() *convLua_PubFuncInfo {
    return obj
}
func NewconvLua_PubFuncInfo(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_TypeInfo) *convLua_PubFuncInfo {
    obj := &convLua_PubFuncInfo{}
    obj.FP = obj
    obj.InitconvLua_PubFuncInfo(_env, arg1, arg2)
    return obj
}
func (self *convLua_PubFuncInfo) InitconvLua_PubFuncInfo(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_TypeInfo) {
    self.AccessMode = arg1
    self.TypeInfo = arg2
}

// declaration Class -- ModuleInfo
type convLua_ModuleInfoMtd interface {
    Get_assignName(_env *LnsEnv) string
    Get_modulePath(_env *LnsEnv) string
}
type convLua_ModuleInfo struct {
    assignName string
    modulePath string
    FP convLua_ModuleInfoMtd
}
func convLua_ModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convLua_ModuleInfo).FP
}
func convLua_ModuleInfo_toSlice(slice []LnsAny) []*convLua_ModuleInfo {
    ret := make([]*convLua_ModuleInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(convLua_ModuleInfoDownCast).ToconvLua_ModuleInfo()
    }
    return ret
}
type convLua_ModuleInfoDownCast interface {
    ToconvLua_ModuleInfo() *convLua_ModuleInfo
}
func convLua_ModuleInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convLua_ModuleInfoDownCast)
    if ok { return work.ToconvLua_ModuleInfo() }
    return nil
}
func (obj *convLua_ModuleInfo) ToconvLua_ModuleInfo() *convLua_ModuleInfo {
    return obj
}
func NewconvLua_ModuleInfo(_env *LnsEnv, arg1 string, arg2 string) *convLua_ModuleInfo {
    obj := &convLua_ModuleInfo{}
    obj.FP = obj
    obj.InitconvLua_ModuleInfo(_env, arg1, arg2)
    return obj
}
func (self *convLua_ModuleInfo) InitconvLua_ModuleInfo(_env *LnsEnv, arg1 string, arg2 string) {
    self.assignName = arg1
    self.modulePath = arg2
}
func (self *convLua_ModuleInfo) Get_assignName(_env *LnsEnv) string{ return self.assignName }
func (self *convLua_ModuleInfo) Get_modulePath(_env *LnsEnv) string{ return self.modulePath }

// declaration Class -- Opt
type ConvLua_OptMtd interface {
}
type ConvLua_Opt struct {
    Node *Nodes_Node
    FP ConvLua_OptMtd
}
func ConvLua_Opt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvLua_Opt).FP
}
func ConvLua_Opt_toSlice(slice []LnsAny) []*ConvLua_Opt {
    ret := make([]*ConvLua_Opt, len(slice))
    for index, val := range slice {
        ret[index] = val.(ConvLua_OptDownCast).ToConvLua_Opt()
    }
    return ret
}
type ConvLua_OptDownCast interface {
    ToConvLua_Opt() *ConvLua_Opt
}
func ConvLua_OptDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvLua_OptDownCast)
    if ok { return work.ToConvLua_Opt() }
    return nil
}
func (obj *ConvLua_Opt) ToConvLua_Opt() *ConvLua_Opt {
    return obj
}
func NewConvLua_Opt(_env *LnsEnv, arg1 *Nodes_Node) *ConvLua_Opt {
    obj := &ConvLua_Opt{}
    obj.FP = obj
    obj.InitConvLua_Opt(_env, arg1)
    return obj
}
func (self *ConvLua_Opt) InitConvLua_Opt(_env *LnsEnv, arg1 *Nodes_Node) {
    self.Node = arg1
}

// declaration Class -- Option
type ConvLua_OptionMtd interface {
    Get_legacyNewMode(_env *LnsEnv) bool
    Get_mainModule(_env *LnsEnv) string
}
type ConvLua_Option struct {
    mainModule string
    legacyNewMode bool
    FP ConvLua_OptionMtd
}
func ConvLua_Option2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvLua_Option).FP
}
func ConvLua_Option_toSlice(slice []LnsAny) []*ConvLua_Option {
    ret := make([]*ConvLua_Option, len(slice))
    for index, val := range slice {
        ret[index] = val.(ConvLua_OptionDownCast).ToConvLua_Option()
    }
    return ret
}
type ConvLua_OptionDownCast interface {
    ToConvLua_Option() *ConvLua_Option
}
func ConvLua_OptionDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvLua_OptionDownCast)
    if ok { return work.ToConvLua_Option() }
    return nil
}
func (obj *ConvLua_Option) ToConvLua_Option() *ConvLua_Option {
    return obj
}
func NewConvLua_Option(_env *LnsEnv, arg1 string, arg2 bool) *ConvLua_Option {
    obj := &ConvLua_Option{}
    obj.FP = obj
    obj.InitConvLua_Option(_env, arg1, arg2)
    return obj
}
func (self *ConvLua_Option) InitConvLua_Option(_env *LnsEnv, arg1 string, arg2 bool) {
    self.mainModule = arg1
    self.legacyNewMode = arg2
}
func (self *ConvLua_Option) Get_mainModule(_env *LnsEnv) string{ return self.mainModule }
func (self *ConvLua_Option) Get_legacyNewMode(_env *LnsEnv) bool{ return self.legacyNewMode }

// declaration Class -- NewMethod
type convLua_NewMethodMtd interface {
}
type convLua_NewMethod struct {
    NewName string
    SetmetaName string
    FP convLua_NewMethodMtd
}
func convLua_NewMethod2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convLua_NewMethod).FP
}
func convLua_NewMethod_toSlice(slice []LnsAny) []*convLua_NewMethod {
    ret := make([]*convLua_NewMethod, len(slice))
    for index, val := range slice {
        ret[index] = val.(convLua_NewMethodDownCast).ToconvLua_NewMethod()
    }
    return ret
}
type convLua_NewMethodDownCast interface {
    ToconvLua_NewMethod() *convLua_NewMethod
}
func convLua_NewMethodDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convLua_NewMethodDownCast)
    if ok { return work.ToconvLua_NewMethod() }
    return nil
}
func (obj *convLua_NewMethod) ToconvLua_NewMethod() *convLua_NewMethod {
    return obj
}
func NewconvLua_NewMethod(_env *LnsEnv, arg1 string, arg2 string) *convLua_NewMethod {
    obj := &convLua_NewMethod{}
    obj.FP = obj
    obj.InitconvLua_NewMethod(_env, arg1, arg2)
    return obj
}
// 99: DeclConstr
func (self *convLua_NewMethod) InitconvLua_NewMethod(_env *LnsEnv, newName string,setmetaName string) {
    self.NewName = newName
    self.SetmetaName = setmetaName
}


// declaration Class -- ConvFilter
type convLua_ConvFilterMtd interface {
    Close(_env *LnsEnv)
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    Flush(_env *LnsEnv)
    getCanonicalName(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    getDestrClass(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny
    GetFull(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    getFullName(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    getMapInfo(_env *LnsEnv, arg1 *Ast_TypeInfo)(string, bool, string)
    get_indent(_env *LnsEnv) LnsInt
    Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl
    outputAlter2MapFunc(_env *LnsEnv, arg1 Lns_oStream, arg2 *LnsMap2_[*Ast_TypeInfo,*Ast_TypeInfo])
    outputCondRet(_env *LnsEnv, arg1 *Nodes_CondRetNode)
    outputCondRetInfo(_env *LnsEnv, arg1 *Nodes_CondRetInfo)
    OutputDeclMacro(_env *LnsEnv, arg1 string, arg2 *LnsList2_[string], arg3 convLua_outputMacroStmtBlock_13_)
    outputMainDirect(_env *LnsEnv)
    outputMeta(_env *LnsEnv, arg1 *Nodes_RootNode)
    popIndent(_env *LnsEnv)
    ProcessAbbr(_env *LnsEnv, arg1 *Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(_env *LnsEnv, arg1 *Nodes_AliasNode, arg2 LnsAny)
    ProcessApply(_env *LnsEnv, arg1 *Nodes_ApplyNode, arg2 LnsAny)
    ProcessAsyncLock(_env *LnsEnv, arg1 *Nodes_AsyncLockNode, arg2 LnsAny)
    ProcessBlankLine(_env *LnsEnv, arg1 *Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(_env *LnsEnv, arg1 *Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(_env *LnsEnv, arg1 *Nodes_BreakNode, arg2 LnsAny)
    ProcessCondRet(_env *LnsEnv, arg1 *Nodes_CondRetNode, arg2 LnsAny)
    ProcessCondRetList(_env *LnsEnv, arg1 *Nodes_CondRetListNode, arg2 LnsAny)
    ProcessConvStat(_env *LnsEnv, arg1 *Nodes_ConvStatNode, arg2 LnsAny)
    ProcessDeclAdvertise(_env *LnsEnv, arg1 *Nodes_DeclAdvertiseNode, arg2 LnsAny)
    ProcessDeclAlge(_env *LnsEnv, arg1 *Nodes_DeclAlgeNode, arg2 LnsAny)
    ProcessDeclArg(_env *LnsEnv, arg1 *Nodes_DeclArgNode, arg2 LnsAny)
    ProcessDeclArgDDD(_env *LnsEnv, arg1 *Nodes_DeclArgDDDNode, arg2 LnsAny)
    ProcessDeclClass(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 LnsAny)
    ProcessDeclConstr(_env *LnsEnv, arg1 *Nodes_DeclConstrNode, arg2 LnsAny)
    ProcessDeclDestr(_env *LnsEnv, arg1 *Nodes_DeclDestrNode, arg2 LnsAny)
    ProcessDeclEnum(_env *LnsEnv, arg1 *Nodes_DeclEnumNode, arg2 LnsAny)
    ProcessDeclForm(_env *LnsEnv, arg1 *Nodes_DeclFormNode, arg2 LnsAny)
    ProcessDeclFunc(_env *LnsEnv, arg1 *Nodes_DeclFuncNode, arg2 LnsAny)
    ProcessDeclMacro(_env *LnsEnv, arg1 *Nodes_DeclMacroNode, arg2 LnsAny)
    ProcessDeclMember(_env *LnsEnv, arg1 *Nodes_DeclMemberNode, arg2 LnsAny)
    ProcessDeclMethod(_env *LnsEnv, arg1 *Nodes_DeclMethodNode, arg2 LnsAny)
    ProcessDeclTuple(_env *LnsEnv, arg1 *Nodes_DeclTupleNode, arg2 LnsAny)
    ProcessDeclVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode, arg2 LnsAny)
    ProcessExpAccessMRet(_env *LnsEnv, arg1 *Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(_env *LnsEnv, arg1 *Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(_env *LnsEnv, arg1 *Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(_env *LnsEnv, arg1 *Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpExpandTuple(_env *LnsEnv, arg1 *Nodes_ExpExpandTupleNode, arg2 LnsAny)
    ProcessExpList(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 LnsAny)
    processExpListSub(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsList2_[*Nodes_Node], arg3 LnsAny)
    ProcessExpMRet(_env *LnsEnv, arg1 *Nodes_ExpMRetNode, arg2 LnsAny)
    ProcessExpMacroArgExp(_env *LnsEnv, arg1 *Nodes_ExpMacroArgExpNode, arg2 LnsAny)
    ProcessExpMacroExp(_env *LnsEnv, arg1 *Nodes_ExpMacroExpNode, arg2 LnsAny)
    ProcessExpMacroStat(_env *LnsEnv, arg1 *Nodes_ExpMacroStatNode, arg2 LnsAny)
    ProcessExpMacroStatList(_env *LnsEnv, arg1 *Nodes_ExpMacroStatListNode, arg2 LnsAny)
    ProcessExpMultiTo1(_env *LnsEnv, arg1 *Nodes_ExpMultiTo1Node, arg2 LnsAny)
    ProcessExpNew(_env *LnsEnv, arg1 *Nodes_ExpNewNode, arg2 LnsAny)
    ProcessExpOmitEnum(_env *LnsEnv, arg1 *Nodes_ExpOmitEnumNode, arg2 LnsAny)
    ProcessExpOp1(_env *LnsEnv, arg1 *Nodes_ExpOp1Node, arg2 LnsAny)
    ProcessExpOp2(_env *LnsEnv, arg1 *Nodes_ExpOp2Node, arg2 LnsAny)
    ProcessExpParen(_env *LnsEnv, arg1 *Nodes_ExpParenNode, arg2 LnsAny)
    ProcessExpRef(_env *LnsEnv, arg1 *Nodes_ExpRefNode, arg2 LnsAny)
    ProcessExpRefItem(_env *LnsEnv, arg1 *Nodes_ExpRefItemNode, arg2 LnsAny)
    ProcessExpSetItem(_env *LnsEnv, arg1 *Nodes_ExpSetItemNode, arg2 LnsAny)
    ProcessExpSetVal(_env *LnsEnv, arg1 *Nodes_ExpSetValNode, arg2 LnsAny)
    ProcessExpSubDDD(_env *LnsEnv, arg1 *Nodes_ExpSubDDDNode, arg2 LnsAny)
    ProcessExpToDDD(_env *LnsEnv, arg1 *Nodes_ExpToDDDNode, arg2 LnsAny)
    ProcessExpUnwrap(_env *LnsEnv, arg1 *Nodes_ExpUnwrapNode, arg2 LnsAny)
    ProcessFor(_env *LnsEnv, arg1 *Nodes_ForNode, arg2 LnsAny)
    ProcessForeach(_env *LnsEnv, arg1 *Nodes_ForeachNode, arg2 LnsAny)
    ProcessForsort(_env *LnsEnv, arg1 *Nodes_ForsortNode, arg2 LnsAny)
    ProcessGetField(_env *LnsEnv, arg1 *Nodes_GetFieldNode, arg2 LnsAny)
    ProcessIf(_env *LnsEnv, arg1 *Nodes_IfNode, arg2 LnsAny)
    ProcessIfUnwrap(_env *LnsEnv, arg1 *Nodes_IfUnwrapNode, arg2 LnsAny)
    ProcessImport(_env *LnsEnv, arg1 *Nodes_ImportNode, arg2 LnsAny)
    ProcessLiteralArray(_env *LnsEnv, arg1 *Nodes_LiteralArrayNode, arg2 LnsAny)
    ProcessLiteralBool(_env *LnsEnv, arg1 *Nodes_LiteralBoolNode, arg2 LnsAny)
    ProcessLiteralChar(_env *LnsEnv, arg1 *Nodes_LiteralCharNode, arg2 LnsAny)
    ProcessLiteralInt(_env *LnsEnv, arg1 *Nodes_LiteralIntNode, arg2 LnsAny)
    ProcessLiteralList(_env *LnsEnv, arg1 *Nodes_LiteralListNode, arg2 LnsAny)
    ProcessLiteralMap(_env *LnsEnv, arg1 *Nodes_LiteralMapNode, arg2 LnsAny)
    ProcessLiteralNil(_env *LnsEnv, arg1 *Nodes_LiteralNilNode, arg2 LnsAny)
    ProcessLiteralReal(_env *LnsEnv, arg1 *Nodes_LiteralRealNode, arg2 LnsAny)
    ProcessLiteralSet(_env *LnsEnv, arg1 *Nodes_LiteralSetNode, arg2 LnsAny)
    ProcessLiteralString(_env *LnsEnv, arg1 *Nodes_LiteralStringNode, arg2 LnsAny)
    ProcessLiteralSymbol(_env *LnsEnv, arg1 *Nodes_LiteralSymbolNode, arg2 LnsAny)
    processLoadRuntime(_env *LnsEnv)
    ProcessLuneControl(_env *LnsEnv, arg1 *Nodes_LuneControlNode, arg2 LnsAny)
    ProcessLuneKind(_env *LnsEnv, arg1 *Nodes_LuneKindNode, arg2 LnsAny)
    ProcessMatch(_env *LnsEnv, arg1 *Nodes_MatchNode, arg2 LnsAny)
    ProcessNewAlgeVal(_env *LnsEnv, arg1 *Nodes_NewAlgeValNode, arg2 LnsAny)
    ProcessNone(_env *LnsEnv, arg1 *Nodes_NoneNode, arg2 LnsAny)
    ProcessProtoClass(_env *LnsEnv, arg1 *Nodes_ProtoClassNode, arg2 LnsAny)
    ProcessProtoMethod(_env *LnsEnv, arg1 *Nodes_ProtoMethodNode, arg2 LnsAny)
    ProcessProvide(_env *LnsEnv, arg1 *Nodes_ProvideNode, arg2 LnsAny)
    ProcessRefField(_env *LnsEnv, arg1 *Nodes_RefFieldNode, arg2 LnsAny)
    ProcessRefType(_env *LnsEnv, arg1 *Nodes_RefTypeNode, arg2 LnsAny)
    ProcessRepeat(_env *LnsEnv, arg1 *Nodes_RepeatNode, arg2 LnsAny)
    ProcessRequest(_env *LnsEnv, arg1 *Nodes_RequestNode, arg2 LnsAny)
    ProcessReturn(_env *LnsEnv, arg1 *Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(_env *LnsEnv, arg1 *Nodes_RootNode, arg2 LnsAny)
    ProcessScope(_env *LnsEnv, arg1 *Nodes_ScopeNode, arg2 LnsAny)
    ProcessShebang(_env *LnsEnv, arg1 *Nodes_ShebangNode, arg2 LnsAny)
    ProcessStmtExp(_env *LnsEnv, arg1 *Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(_env *LnsEnv, arg1 *Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(_env *LnsEnv, arg1 *Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(_env *LnsEnv, arg1 *Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(_env *LnsEnv, arg1 *Nodes_TestCaseNode, arg2 LnsAny)
    ProcessTupleConst(_env *LnsEnv, arg1 *Nodes_TupleConstNode, arg2 LnsAny)
    ProcessUnboxing(_env *LnsEnv, arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(_env *LnsEnv, arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(_env *LnsEnv, arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(_env *LnsEnv, arg1 *Nodes_WhileNode, arg2 LnsAny)
    process__func__symbol(_env *LnsEnv, arg1 bool, arg2 *Ast_TypeInfo, arg3 string)
    pushIndent(_env *LnsEnv, arg1 LnsAny)
    Write(_env *LnsEnv, arg1 string)(LnsAny, LnsAny)
    WriteRaw(_env *LnsEnv, arg1 string)
    writeRawln(_env *LnsEnv, arg1 string)
    writeln(_env *LnsEnv, arg1 string)
}
type convLua_ConvFilter struct {
    Nodes_Filter
    streamName string
    stream Lns_oStream
    metaStream Lns_oStream
    outMetaFlag bool
    convMode LnsInt
    inMacro bool
    indentQueue *LnsList2_[LnsInt]
    curLineNo LnsInt
    classId2TypeInfo *LnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]
    classId2MemberList *LnsMap2_[*Ast_IdInfo,*LnsList2_[*Nodes_DeclMemberNode]]
    pubEnumId2EnumTypeInfo *LnsMap2_[*Ast_IdInfo,*Ast_EnumTypeInfo]
    pubAlgeId2AlgeTypeInfo *LnsMap2_[*Ast_IdInfo,*Ast_AlgeTypeInfo]
    pubVarName2InfoMap *LnsMap2_[string,*convLua_PubVerInfo]
    pubFuncName2InfoMap *LnsMap2_[string,*convLua_PubFuncInfo]
    needIndent bool
    macroDepth LnsInt
    macroVarSymSet *LnsSet2_[*Ast_SymbolInfo]
    moduleTypeInfo *Ast_TypeInfo
    moduleSymbolKind LnsInt
    needModuleObj bool
    useLuneRuntime LnsAny
    targetLuaVer *LuaVer_LuaVerInfo
    enableTest bool
    useIpairs bool
    processInfo *Ast_ProcessInfo
    moduleType2SymbolMap *LnsMap2_[*Ast_TypeInfo,*Ast_SymbolInfo]
    builtinFunc *Builtin_BuiltinFuncType
    builtinSym2code *LnsMap2_[*Ast_SymbolInfo,string]
    option *ConvLua_Option
    newMethod *convLua_NewMethod
    FP convLua_ConvFilterMtd
}
func convLua_ConvFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convLua_ConvFilter).FP
}
func convLua_ConvFilter_toSlice(slice []LnsAny) []*convLua_ConvFilter {
    ret := make([]*convLua_ConvFilter, len(slice))
    for index, val := range slice {
        ret[index] = val.(convLua_ConvFilterDownCast).ToconvLua_ConvFilter()
    }
    return ret
}
type convLua_ConvFilterDownCast interface {
    ToconvLua_ConvFilter() *convLua_ConvFilter
}
func convLua_ConvFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convLua_ConvFilterDownCast)
    if ok { return work.ToconvLua_ConvFilter() }
    return nil
}
func (obj *convLua_ConvFilter) ToconvLua_ConvFilter() *convLua_ConvFilter {
    return obj
}
func NewconvLua_ConvFilter(_env *LnsEnv, arg1 string, arg2 Lns_oStream, arg3 Lns_oStream, arg4 LnsInt, arg5 bool, arg6 *Ast_TypeInfo, arg7 *Ast_ProcessInfo, arg8 LnsInt, arg9 *Builtin_BuiltinFuncType, arg10 LnsAny, arg11 *LuaVer_LuaVerInfo, arg12 bool, arg13 bool, arg14 *ConvLua_Option) *convLua_ConvFilter {
    obj := &convLua_ConvFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvLua_ConvFilter(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14)
    return obj
}
// 153: DeclConstr
func (self *convLua_ConvFilter) InitconvLua_ConvFilter(_env *LnsEnv, streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,moduleTypeInfo *Ast_TypeInfo,processInfo *Ast_ProcessInfo,moduleSymbolKind LnsInt,builtinFunc *Builtin_BuiltinFuncType,useLuneRuntime LnsAny,targetLuaVer *LuaVer_LuaVerInfo,enableTest bool,useIpairs bool,option *ConvLua_Option) {
    self.InitNodes_Filter(_env, true, moduleTypeInfo, moduleTypeInfo.FP.Get_scope(_env))
    if stream == metaStream{
        Util_err(_env, "streamName == stream")
    }
    self.newMethod = convLua_NewMethod_create_1_(_env, option.FP.Get_legacyNewMode(_env))
    self.option = option
    self.builtinFunc = builtinFunc
    self.moduleType2SymbolMap = NewLnsMap2_[*Ast_TypeInfo,*Ast_SymbolInfo]( map[*Ast_TypeInfo]*Ast_SymbolInfo{})
    self.processInfo = processInfo
    self.enableTest = enableTest
    self.macroVarSymSet = NewLnsSet2_[*Ast_SymbolInfo]([]*Ast_SymbolInfo{})
    self.needModuleObj = true
    self.indentQueue = NewLnsList2_[LnsInt](Lns_2DDDGen[LnsInt](0))
    self.moduleSymbolKind = moduleSymbolKind
    self.macroDepth = 0
    self.streamName = streamName
    self.stream = stream
    self.metaStream = metaStream
    self.outMetaFlag = false
    self.convMode = convMode
    self.inMacro = inMacro
    self.curLineNo = 1
    self.classId2TypeInfo = NewLnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]( map[*Ast_IdInfo]*Ast_TypeInfo{})
    self.classId2MemberList = NewLnsMap2_[*Ast_IdInfo,*LnsList2_[*Nodes_DeclMemberNode]]( map[*Ast_IdInfo]*LnsList2_[*Nodes_DeclMemberNode]{})
    self.pubVarName2InfoMap = NewLnsMap2_[string,*convLua_PubVerInfo]( map[string]*convLua_PubVerInfo{})
    self.pubFuncName2InfoMap = NewLnsMap2_[string,*convLua_PubFuncInfo]( map[string]*convLua_PubFuncInfo{})
    self.pubEnumId2EnumTypeInfo = NewLnsMap2_[*Ast_IdInfo,*Ast_EnumTypeInfo]( map[*Ast_IdInfo]*Ast_EnumTypeInfo{})
    self.pubAlgeId2AlgeTypeInfo = NewLnsMap2_[*Ast_IdInfo,*Ast_AlgeTypeInfo]( map[*Ast_IdInfo]*Ast_AlgeTypeInfo{})
    self.needIndent = false
    self.moduleTypeInfo = moduleTypeInfo
    self.useLuneRuntime = useLuneRuntime
    self.targetLuaVer = targetLuaVer
    self.useIpairs = useIpairs
    self.builtinSym2code = NewLnsMap2_[*Ast_SymbolInfo,string]( map[*Ast_SymbolInfo]string{builtinFunc.G__lns_runmode_Sync_sym:_env.GetVM().String_format("%d", Lns_2DDD(0)),builtinFunc.G__lns_runmode_Queue_sym:_env.GetVM().String_format("%d", Lns_2DDD(1)),builtinFunc.G__lns_runmode_Skip_sym:_env.GetVM().String_format("%d", Lns_2DDD(2)),builtinFunc.G__lns_capability_async_sym:"false",})
}



// declaration Class -- FilterInfo
type ConvLua_FilterInfoMtd interface {
    Get_filter(_env *LnsEnv) *Nodes_Filter
    OutputLua(_env *LnsEnv, arg1 *Nodes_RootNode)
    OutputLuaAndMeta(_env *LnsEnv, arg1 *Nodes_RootNode)
    OutputMeta(_env *LnsEnv, arg1 *Nodes_RootNode)
}
type ConvLua_FilterInfo struct {
    filter *Nodes_Filter
    FP ConvLua_FilterInfoMtd
}
func ConvLua_FilterInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvLua_FilterInfo).FP
}
func ConvLua_FilterInfo_toSlice(slice []LnsAny) []*ConvLua_FilterInfo {
    ret := make([]*ConvLua_FilterInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(ConvLua_FilterInfoDownCast).ToConvLua_FilterInfo()
    }
    return ret
}
type ConvLua_FilterInfoDownCast interface {
    ToConvLua_FilterInfo() *ConvLua_FilterInfo
}
func ConvLua_FilterInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvLua_FilterInfoDownCast)
    if ok { return work.ToConvLua_FilterInfo() }
    return nil
}
func (obj *ConvLua_FilterInfo) ToConvLua_FilterInfo() *ConvLua_FilterInfo {
    return obj
}
func NewConvLua_FilterInfo(_env *LnsEnv, arg1 *Nodes_Filter) *ConvLua_FilterInfo {
    obj := &ConvLua_FilterInfo{}
    obj.FP = obj
    obj.InitConvLua_FilterInfo(_env, arg1)
    return obj
}
func (self *ConvLua_FilterInfo) InitConvLua_FilterInfo(_env *LnsEnv, arg1 *Nodes_Filter) {
    self.filter = arg1
}
func (self *ConvLua_FilterInfo) Get_filter(_env *LnsEnv) *Nodes_Filter{ return self.filter }

// declaration Class -- MacroEvalImp
type ConvLua_MacroEvalImpMtd interface {
    EvalFromCodeToLuaCode(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 *LnsList2_[string], arg4 LnsAny) string
    EvalToLuaCode(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Nodes_DeclMacroNode) string
}
type ConvLua_MacroEvalImp struct {
    Nodes_MacroEval
    builtinFunc *Builtin_BuiltinFuncType
    FP ConvLua_MacroEvalImpMtd
}
func ConvLua_MacroEvalImp2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvLua_MacroEvalImp).FP
}
func ConvLua_MacroEvalImp_toSlice(slice []LnsAny) []*ConvLua_MacroEvalImp {
    ret := make([]*ConvLua_MacroEvalImp, len(slice))
    for index, val := range slice {
        ret[index] = val.(ConvLua_MacroEvalImpDownCast).ToConvLua_MacroEvalImp()
    }
    return ret
}
type ConvLua_MacroEvalImpDownCast interface {
    ToConvLua_MacroEvalImp() *ConvLua_MacroEvalImp
}
func ConvLua_MacroEvalImpDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvLua_MacroEvalImpDownCast)
    if ok { return work.ToConvLua_MacroEvalImp() }
    return nil
}
func (obj *ConvLua_MacroEvalImp) ToConvLua_MacroEvalImp() *ConvLua_MacroEvalImp {
    return obj
}
func NewConvLua_MacroEvalImp(_env *LnsEnv, arg1 *Builtin_BuiltinFuncType) *ConvLua_MacroEvalImp {
    obj := &ConvLua_MacroEvalImp{}
    obj.FP = obj
    obj.Nodes_MacroEval.FP = obj
    obj.InitConvLua_MacroEvalImp(_env, arg1)
    return obj
}
func (self *ConvLua_MacroEvalImp) InitConvLua_MacroEvalImp(_env *LnsEnv, arg1 *Builtin_BuiltinFuncType) {
    self.Nodes_MacroEval.InitNodes_MacroEval( _env)
    self.builtinFunc = arg1
}

func Lns_convLua_init(_env *LnsEnv) {
    if init_convLua { return }
    init_convLua = true
    convLua__mod__ = "@lune.@base.@convLua"
    Lns_InitMod()
    Lns_Ver_init(_env)
    Lns_Str_init(_env)
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Util_init(_env)
    Lns_AstInfo_init(_env)
    Lns_TransUnit_init(_env)
    Lns_LuaMod_init(_env)
    Lns_LuaVer_init(_env)
    Lns_Tokenizer_init(_env)
    Lns_Types_init(_env)
    Lns_Log_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Option_init(_env)
    Lns_frontInterface_init(_env)
    Lns_Builtin_init(_env)
    convLua_stepIndent = 3
    
    
    
    convLua_asyncLockBreakSym = "_break"
}
func init() {
    init_convLua = false
}
// 103: decl @lune.@base.@convLua.NewMethod.create
func convLua_NewMethod_create_1_(_env *LnsEnv, legacy bool) *convLua_NewMethod {
    if legacy{
        return NewconvLua_NewMethod(_env, "new", "setmeta")
    } else { 
        return NewconvLua_NewMethod(_env, "_new", "_setmeta")
    }
// insert a dummy
    return nil
}
// 204: decl @lune.@base.@convLua.ConvFilter.get_indent
func (self *convLua_ConvFilter) get_indent(_env *LnsEnv) LnsInt {
    if self.indentQueue.Len() > 0{
        return self.indentQueue.GetAt(self.indentQueue.Len())
    }
    return 0
}
// 211: decl @lune.@base.@convLua.ConvFilter.getCanonicalName
func (self *convLua_ConvFilter) getCanonicalName(_env *LnsEnv, typeInfo *Ast_TypeInfo,localFlag bool) string {
    if _switch0 := typeInfo.FP.Get_genSrcTypeInfo(_env); _switch0 == self.builtinFunc.G__ret_ {
        return "_lune.Result"
    } else if _switch0 == self.builtinFunc.G__lns_it_ {
        return "_lune.__lns_it"
    }
    var enumName string
    enumName = typeInfo.FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), localFlag)
    var moduleType *Ast_TypeInfo
    moduleType = typeInfo.FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.GetModule(_env)
    var canonical string
    canonical = Lns_car(_env.GetVM().String_gsub(enumName,"&", "")).(string)
    {
        _assignSym := self.moduleType2SymbolMap.Get(moduleType)
        if !Lns_IsNil( _assignSym ) {
            assignSym := _assignSym.(*Ast_SymbolInfo)
            if assignSym.FP.Get_isLazyLoad(_env){
                var index LnsInt
                index = Lns_unwrap( Lns_car(_env.GetVM().String_find(canonical,".", 1, true))).(LnsInt)
                return _env.GetVM().String_format("%s().%s", Lns_2DDD(_env.GetVM().String_sub(canonical,1, index - 1), _env.GetVM().String_sub(canonical,index + 1, nil)))
            }
        }
    }
    return canonical
}
// 236: decl @lune.@base.@convLua.ConvFilter.getFullName
func (self *convLua_ConvFilter) getFullName(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    return self.FP.getCanonicalName(_env, typeInfo, true)
}
// 240: decl @lune.@base.@convLua.ConvFilter.close
func (self *convLua_ConvFilter) Close(_env *LnsEnv) {
}
// 242: decl @lune.@base.@convLua.ConvFilter.flush
func (self *convLua_ConvFilter) Flush(_env *LnsEnv) {
}
// 245: decl @lune.@base.@convLua.ConvFilter.writeRaw
func (self *convLua_ConvFilter) WriteRaw(_env *LnsEnv, txt string) {
    var stream Lns_oStream
    stream = self.stream
    if self.outMetaFlag{
        stream = self.metaStream
    }
    if self.needIndent{
        stream.Write(_env, _env.GetVM().String_rep(" ", self.FP.get_indent(_env)))
        self.needIndent = false
    }
    stream.Write(_env, txt)
}
// 257: decl @lune.@base.@convLua.ConvFilter.write
func (self *convLua_ConvFilter) Write(_env *LnsEnv, txt string)(LnsAny, LnsAny) {
    var stream Lns_oStream
    stream = self.stream
    if self.outMetaFlag{
        stream = self.metaStream
    }
    for _, _line := range( Str_getLineList(_env, txt).Items ) {
        line := _line.(string)
        if self.needIndent{
            stream.Write(_env, _env.GetVM().String_rep(" ", self.FP.get_indent(_env)))
            self.needIndent = false
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( len(line) > 0) &&
            _env.SetStackVal( LnsInt(line[len(line)-1]) == 10) ).(bool)){
            self.curLineNo = self.curLineNo + 1
        }
        stream.Write(_env, line)
    }
    return self.FP, nil
}
// 294: decl @lune.@base.@convLua.ConvFilter.processBlankLine
func (self *convLua_ConvFilter) ProcessBlankLine(_env *LnsEnv, node *Nodes_BlankLineNode,_opt LnsAny) {
}
// 294: decl @lune.@base.@convLua.ConvFilter.processDeclForm
func (self *convLua_ConvFilter) ProcessDeclForm(_env *LnsEnv, node *Nodes_DeclFormNode,_opt LnsAny) {
}
// 294: decl @lune.@base.@convLua.ConvFilter.processProtoMethod
func (self *convLua_ConvFilter) ProcessProtoMethod(_env *LnsEnv, node *Nodes_ProtoMethodNode,_opt LnsAny) {
}
// 312: decl @lune.@base.@convLua.ConvFilter.pushIndent
func (self *convLua_ConvFilter) pushIndent(_env *LnsEnv, newIndent LnsAny) {
    var indent LnsInt
    indent = Lns_unwrapDefault( newIndent, self.FP.get_indent(_env) + convLua_stepIndent).(LnsInt)
    self.indentQueue.Insert(LnsInt(indent))
}
// 317: decl @lune.@base.@convLua.ConvFilter.popIndent
func (self *convLua_ConvFilter) popIndent(_env *LnsEnv) {
    if self.indentQueue.Len() == 0{
        Util_err(_env, "self.indentQueue == 0")
    }
    self.indentQueue.Remove(nil)
}
// 325: decl @lune.@base.@convLua.ConvFilter.writeln
func (self *convLua_ConvFilter) writeln(_env *LnsEnv, txt string) {
    self.FP.Write(_env, txt)
    self.FP.WriteRaw(_env, "\n")
    self.needIndent = true
}
// 331: decl @lune.@base.@convLua.ConvFilter.writeRawln
func (self *convLua_ConvFilter) writeRawln(_env *LnsEnv, txt string) {
    self.FP.WriteRaw(_env, txt)
    self.FP.WriteRaw(_env, "\n")
    self.needIndent = true
}
// 337: decl @lune.@base.@convLua.ConvFilter.processNone
func (self *convLua_ConvFilter) ProcessNone(_env *LnsEnv, node *Nodes_NoneNode,_opt LnsAny) {
}
// 342: decl @lune.@base.@convLua.ConvFilter.processShebang
func (self *convLua_ConvFilter) ProcessShebang(_env *LnsEnv, node *Nodes_ShebangNode,_opt LnsAny) {
}
// 347: decl @lune.@base.@convLua.ConvFilter.processImport
func (self *convLua_ConvFilter) ProcessImport(_env *LnsEnv, node *Nodes_ImportNode,_opt LnsAny) {
    var info *Nodes_ImportInfo
    info = node.FP.Get_info(_env)
    var modulePath string
    modulePath = info.FP.Get_modulePath(_env)
    var modSym string
    modSym = info.FP.Get_assignName(_env)
    self.FP.WriteRaw(_env, _env.GetVM().String_format("local %s = _lune.", Lns_2DDD(modSym)))
    if _switch0 := info.FP.Get_lazy(_env); _switch0 == Nodes_LazyLoad__Off {
        self.FP.WriteRaw(_env, "loadModule")
    } else if _switch0 == Nodes_LazyLoad__On || _switch0 == Nodes_LazyLoad__Auto {
        self.FP.WriteRaw(_env, "_lazyImport")
    }
    self.FP.writeRawln(_env, _env.GetVM().String_format("( '%s' )", Lns_2DDD(modulePath)))
}
// 392: decl @lune.@base.@convLua.ConvFilter.outputMeta
func (self *convLua_ConvFilter) outputMeta(_env *LnsEnv, node *Nodes_RootNode) {
    if self.convMode == ConvLua_ConvMode__Convert{
        return 
    }
    self.outMetaFlag = true
    self.FP.writeln(_env, "local _moduleObj = {}")
    self.FP.writeln(_env, "----- meta -----")
    self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.__version = '%s'", Lns_2DDD(Ver_version)))
    self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.__formatVersion = '%s'", Lns_2DDD(Ver_metaVersion)))
    self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.__buildId = %q", Lns_2DDD(node.FP.Get_moduleId(_env).FP.GetNextModuleId(_env).FP.Get_idStr(_env))))
    self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.__enableTest = %s", Lns_2DDD(self.enableTest)))
    self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.__hasTest = %s", Lns_2DDD(node.FP.Get_nodeManager(_env).FP.GetTestCaseNodeList(_env).Len() != 0)))
    self.FP.WriteRaw(_env, "_moduleObj.__lazyModuleList = {")
    {
        var firstFlag bool
        firstFlag = true
        for _, _declClass := range( node.FP.Get_nodeManager(_env).FP.GetDeclClassNodeList(_env).Items ) {
            declClass := _declClass.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( declClass.FP.Get_lazyLoad(_env) != Nodes_LazyLoad__Off) &&
                _env.SetStackVal( Ast_isPubToExternal(_env, declClass.FP.Get_accessMode(_env))) ).(bool)){
                if Lns_op_not(firstFlag){
                    self.FP.WriteRaw(_env, ",")
                } else { 
                    firstFlag = false
                }
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%d", Lns_2DDD(declClass.FP.Get_expType(_env).FP.Get_typeId(_env).Id)))
            }
        }
    }
    self.FP.writeln(_env, "}")
    var importModuleType2Index *LnsMap2_[*Ast_TypeInfo,LnsInt]
    importModuleType2Index = NewLnsMap2_[*Ast_TypeInfo,LnsInt]( map[*Ast_TypeInfo]LnsInt{})
    var importProcessInfo2Index *LnsMap2_[*Ast_ProcessInfo,LnsInt]
    importProcessInfo2Index = NewLnsMap2_[*Ast_ProcessInfo,LnsInt]( map[*Ast_ProcessInfo]LnsInt{})
    importProcessInfo2Index.Set(Ast_getRootProcessInfoRo(_env),FrontInterface_getRootDependModId(_env))
    importProcessInfo2Index.Set(self.processInfo,0)
    importProcessInfo2Index.Set(self.processInfo.FP.Get_orgInfo(_env),0)
    var importNameMap *LnsMap2_[string,*FrontInterface_ExportInfo]
    importNameMap = NewLnsMap2_[string,*FrontInterface_ExportInfo]( map[string]*FrontInterface_ExportInfo{})
    {
        for _, _exportInfo := range( node.FP.Get_importModule2moduleInfo(_env).Items ) {
            exportInfo := _exportInfo
            importNameMap.Set(exportInfo.FP.Get_fullName(_env),exportInfo)
        }
        var index LnsInt
        index = 0
        {
            __forsortCollection0 := importNameMap
            __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
            __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
            for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
                exportInfo := __forsortCollection0.Items[ ___forsortKey0 ]
                index = index + 1
                importModuleType2Index.Set(exportInfo.FP.Get_moduleTypeInfo(_env),index)
                importProcessInfo2Index.Set(exportInfo.FP.Get_processInfo(_env),index)
                importProcessInfo2Index.Set(exportInfo.FP.Get_processInfo(_env).FP.Get_orgInfo(_env),index)
            }
        }
    }
    var typeId2TypeInfo *LnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]
    typeId2TypeInfo = NewLnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]( map[*Ast_IdInfo]*Ast_TypeInfo{})
    {
        var work *Ast_TypeInfo
        work = node.FP.Get_moduleTypeInfo(_env)
        for Ast_TypeInfo_hasParent(_env, work) {
            typeId2TypeInfo.Set(work.FP.Get_typeId(_env),work)
            work = work.FP.Get_parentInfo(_env)
        }
    }
    var pickupClassMap *LnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]
    pickupClassMap = NewLnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]( map[*Ast_IdInfo]*Ast_TypeInfo{})
    var convLua_checkExportTypeInfo func(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny
    convLua_checkExportTypeInfo = func(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny {
        var moduleTypeInfo *Ast_TypeInfo
        moduleTypeInfo = typeInfo.FP.GetModule(_env)
        var typeId *Ast_IdInfo
        typeId = typeInfo.FP.Get_typeId(_env)
        return _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( typeId2TypeInfo.Get(typeId)) &&
            _env.SetStackVal( Lns_op_not(Ast_isBuiltin(_env, typeId.Id))) &&
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( moduleTypeInfo.FP.HasRouteNamespaceFrom(_env, node.FP.Get_moduleTypeInfo(_env))) ||
                _env.SetStackVal( typeInfo.FP.Get_srcTypeInfo(_env) != typeInfo) ||
                _env.SetStackVal( moduleTypeInfo.FP.Equals(_env, self.processInfo, Ast_headTypeInfo, nil, nil)) ).(bool))) )
    }
    var convLua_isDependOnExt func(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool
    convLua_isDependOnExt = func(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
        if Ast_isExtId(_env, typeInfo){
            return true
        }
        return self.moduleTypeInfo.FP.Get_processInfo(_env) != typeInfo.FP.Get_typeId(_env).FP.Get_processInfo(_env)
    }
    var convLua_pickupTypeId func(_env *LnsEnv, typeInfo *Ast_TypeInfo,forceFlag LnsAny,pickupChildFlag LnsAny)
    convLua_pickupTypeId = func(_env *LnsEnv, typeInfo *Ast_TypeInfo,forceFlag LnsAny,pickupChildFlag LnsAny) {
        if typeInfo.FP.Get_typeId(_env) == Ast_rootTypeIdInfo{
            return 
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(forceFlag)) &&
            _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, typeInfo.FP.Get_accessMode(_env)))) ).(bool)){
            return 
        }
        if Lns_isCondTrue( typeId2TypeInfo.Get(typeInfo.FP.Get_typeId(_env))){
            if convLua_isDependOnExt(_env, typeInfo){
                return 
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( pickupChildFlag) &&
                _env.SetStackVal( Lns_op_not(typeInfo.FP.Get_nilable(_env))) )){
                for _, _itemTypeInfo := range( typeInfo.FP.Get_children(_env).Items ) {
                    itemTypeInfo := _itemTypeInfo
                    if Ast_isPubToExternal(_env, itemTypeInfo.FP.Get_accessMode(_env)){
                        if _switch0 := itemTypeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Class || _switch0 == Ast_TypeInfoKind__IF || _switch0 == Ast_TypeInfoKind__Form || _switch0 == Ast_TypeInfoKind__FormFunc || _switch0 == Ast_TypeInfoKind__Func || _switch0 == Ast_TypeInfoKind__Method {
                            convLua_pickupTypeId(_env, itemTypeInfo, true, true)
                        }
                    }
                }
            }
            return 
        }
        typeId2TypeInfo.Set(typeInfo.FP.Get_typeId(_env),typeInfo)
        if typeInfo.FP.IsModule(_env){
            return 
        }
        if Ast_isBuiltin(_env, typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_typeId(_env).Id){
            return 
        }
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
            convLua_pickupTypeId(_env, typeInfo.FP.Get_extedType(_env), true, true)
        }
        if typeInfo != typeInfo.FP.Get_srcTypeInfo(_env){
            convLua_pickupTypeId(_env, typeInfo.FP.Get_srcTypeInfo(_env), true, false)
        } else if typeInfo.FP.Get_nilable(_env){
            convLua_pickupTypeId(_env, typeInfo.FP.Get_nonnilableType(_env), true, false)
        } else { 
            if convLua_isDependOnExt(_env, typeInfo){
                return 
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ).(bool){
                pickupClassMap.Set(typeInfo.FP.Get_typeId(_env),typeInfo)
            }
            if Lns_op_not(typeInfo.FP.Get_externalFlag(_env)){
                if _switch1 := typeInfo.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__IF || _switch1 == Ast_TypeInfoKind__Class || _switch1 == Ast_TypeInfoKind__Form || _switch1 == Ast_TypeInfoKind__FormFunc || _switch1 == Ast_TypeInfoKind__Alge || _switch1 == Ast_TypeInfoKind__Enum || _switch1 == Ast_TypeInfoKind__Map || _switch1 == Ast_TypeInfoKind__Set || _switch1 == Ast_TypeInfoKind__List || _switch1 == Ast_TypeInfoKind__Array || _switch1 == Ast_TypeInfoKind__Tuple || _switch1 == Ast_TypeInfoKind__Alternate || _switch1 == Ast_TypeInfoKind__Box {
                    convLua_pickupTypeId(_env, typeInfo.FP.Get_nilableTypeInfo(_env), true, false)
                }
            }
            var parentInfo *Ast_TypeInfo
            parentInfo = typeInfo.FP.Get_parentInfo(_env)
            convLua_pickupTypeId(_env, parentInfo, true, false)
            convLua_pickupTypeId(_env, typeInfo.FP.Get_genSrcTypeInfo(_env), true, false)
            var baseInfo *Ast_TypeInfo
            baseInfo = typeInfo.FP.Get_baseTypeInfo(_env)
            if baseInfo.FP.Get_typeId(_env) != Ast_rootTypeIdInfo{
                convLua_pickupTypeId(_env, baseInfo, true, true)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_interfaceList(_env).Items ) {
                itemTypeInfo := _itemTypeInfo
                convLua_pickupTypeId(_env, itemTypeInfo, true, true)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_itemTypeInfoList(_env).Items ) {
                itemTypeInfo := _itemTypeInfo
                convLua_pickupTypeId(_env, itemTypeInfo, true, false)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_argTypeInfoList(_env).Items ) {
                itemTypeInfo := _itemTypeInfo
                convLua_pickupTypeId(_env, itemTypeInfo, true, false)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_retTypeInfoList(_env).Items ) {
                itemTypeInfo := _itemTypeInfo
                convLua_pickupTypeId(_env, itemTypeInfo, true, true)
            }
            if Lns_isCondTrue( pickupChildFlag){
                for _, _itemTypeInfo := range( typeInfo.FP.Get_children(_env).Items ) {
                    itemTypeInfo := _itemTypeInfo
                    if itemTypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
                        if _switch2 := itemTypeInfo.FP.Get_kind(_env); _switch2 == Ast_TypeInfoKind__Class || _switch2 == Ast_TypeInfoKind__IF || _switch2 == Ast_TypeInfoKind__Form || _switch2 == Ast_TypeInfoKind__FormFunc || _switch2 == Ast_TypeInfoKind__Func || _switch2 == Ast_TypeInfoKind__Method {
                            convLua_pickupTypeId(_env, itemTypeInfo, true, true)
                        }
                    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( itemTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) &&
                        _env.SetStackVal( itemTypeInfo.FP.Get_rawTxt(_env) == "__free") ).(bool)){
                        typeId2TypeInfo.Set(itemTypeInfo.FP.Get_typeId(_env),itemTypeInfo)
                    }
                }
            }
        }
    }
    var validChildrenSet *LnsMap2_[*Ast_TypeInfo,*LnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]]
    validChildrenSet = NewLnsMap2_[*Ast_TypeInfo,*LnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]]( map[*Ast_TypeInfo]*LnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]{})
    {
        var typeInfo *Ast_TypeInfo
        typeInfo = self.moduleTypeInfo
        for Ast_TypeInfo_hasParent(_env, typeInfo) {
            validChildrenSet.Set(typeInfo.FP.Get_parentInfo(_env),NewLnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]( map[*Ast_IdInfo]*Ast_TypeInfo{typeInfo.FP.Get_typeId(_env):typeInfo,}))
            typeInfo = typeInfo.FP.Get_parentInfo(_env)
        }
        convLua_pickupTypeId(_env, self.moduleTypeInfo, true, nil)
    }
    var typeId2ClassMap *LnsMap2_[*Ast_IdInfo,*Nodes_NamespaceInfo]
    typeId2ClassMap = node.FP.Get_typeId2ClassMap(_env)
    for _, _namespaceInfo := range( typeId2ClassMap.Items ) {
        namespaceInfo := _namespaceInfo
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( namespaceInfo.TypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) &&
            _env.SetStackVal( Lns_op_not(namespaceInfo.TypeInfo.FP.Get_externalFlag(_env))) ).(bool)){
            pickupClassMap.Set(namespaceInfo.TypeInfo.FP.Get_typeId(_env),namespaceInfo.TypeInfo)
        }
    }
    var classId2TypeInfo *LnsMap2_[LnsInt,*Ast_TypeInfo]
    classId2TypeInfo = NewLnsMap2_[LnsInt,*Ast_TypeInfo]( map[LnsInt]*Ast_TypeInfo{})
    for _idInfo, _classTypeInfo := range( self.classId2TypeInfo.Items ) {
        idInfo := _idInfo
        classTypeInfo := _classTypeInfo
        classId2TypeInfo.Set(idInfo.Id,classTypeInfo)
    }
    var serializeInfo *Ast_SerializeInfo
    serializeInfo = NewAst_SerializeInfo(_env, importProcessInfo2Index, NewLnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]( map[*Ast_IdInfo]*Ast_TypeInfo{}))
    self.FP.writeln(_env, "local __typeId2ClassInfoMap = {}")
    self.FP.writeln(_env, "_moduleObj.__typeId2ClassInfoMap = __typeId2ClassInfoMap")
    {
        __forsortCollection1 := classId2TypeInfo
        __forsortSorted1 := __forsortCollection1.CreateKeyListInt()
        __forsortSorted1.Sort( _env, LnsItemKindInt, nil )
        for _, _classTypeId := range( __forsortSorted1.Items ) {
            classTypeInfo := __forsortCollection1.Items[ _classTypeId ]
            classTypeId := _classTypeId
            if classTypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
                convLua_pickupTypeId(_env, classTypeInfo, true, _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( validChildrenSet.Get(classTypeInfo) == nil) &&
                    _env.SetStackVal( Lns_op_not(classTypeInfo.FP.Get_externalFlag(_env))) ).(bool))
                pickupClassMap.Set(classTypeInfo.FP.Get_typeId(_env),nil)
                self.FP.writeln(_env, "do")
                self.FP.pushIndent(_env, nil)
                self.FP.writeln(_env, _env.GetVM().String_format("local __classInfo%d = {}", Lns_2DDD(classTypeId)))
                self.FP.writeln(_env, _env.GetVM().String_format("__typeId2ClassInfoMap[ %d ] = __classInfo%d", Lns_2DDD(classTypeId, classTypeId)))
                for _, _memberNode := range( Lns_unwrap( self.classId2MemberList.Get(classTypeInfo.FP.Get_typeId(_env))).(*LnsList2_[*Nodes_DeclMemberNode]).Items ) {
                    memberNode := _memberNode
                    if Ast_isPubToExternal(_env, memberNode.FP.Get_accessMode(_env)){
                        var memberName string
                        memberName = memberNode.FP.Get_name(_env).Txt
                        var memberTypeInfo *Ast_TypeInfo
                        memberTypeInfo = memberNode.FP.Get_expType(_env)
                        self.FP.writeln(_env, _env.GetVM().String_format("__classInfo%d.%s = {", Lns_2DDD(classTypeId, memberName)))
                        self.FP.writeln(_env, _env.GetVM().String_format("  name='%s', staticFlag = %s, mutMode = %d,", Lns_2DDD(memberName, memberNode.FP.Get_staticFlag(_env), memberNode.FP.Get_symbolInfo(_env).FP.Get_mutMode(_env))) + _env.GetVM().String_format("accessMode = %d, typeId = %s }", Lns_2DDD(memberNode.FP.Get_accessMode(_env), serializeInfo.FP.SerializeId(_env, memberTypeInfo.FP.Get_typeId(_env)))))
                        convLua_pickupTypeId(_env, memberTypeInfo, true, nil)
                    }
                }
                self.FP.popIndent(_env)
                self.FP.writeln(_env, "end")
            }
        }
    }
    var pickupedClassMap *LnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]
    pickupedClassMap = NewLnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]( map[*Ast_IdInfo]*Ast_TypeInfo{})
    for  {
        var workClassMap *LnsMap2_[LnsInt,*Ast_TypeInfo]
        workClassMap = NewLnsMap2_[LnsInt,*Ast_TypeInfo]( map[LnsInt]*Ast_TypeInfo{})
        var hasWorkClassFlag bool
        hasWorkClassFlag = false
        for _classTypeId, _classTypeInfo := range( pickupClassMap.Items ) {
            classTypeId := _classTypeId
            classTypeInfo := _classTypeInfo
            if Lns_op_not(pickupedClassMap.Get(classTypeId)){
                pickupedClassMap.Set(classTypeId,classTypeInfo)
                if Lns_isCondTrue( convLua_checkExportTypeInfo(_env, classTypeInfo)){
                    workClassMap.Set(classTypeId.Id,classTypeInfo)
                    hasWorkClassFlag = true
                }
            }
        }
        if Lns_op_not(hasWorkClassFlag){
            break
        }
        {
            __forsortCollection3 := workClassMap
            __forsortSorted3 := __forsortCollection3.CreateKeyListInt()
            __forsortSorted3.Sort( _env, LnsItemKindInt, nil )
            for _, _classTypeId := range( __forsortSorted3.Items ) {
                classTypeInfo := __forsortCollection3.Items[ _classTypeId ]
                classTypeId := _classTypeId
                if Lns_op_not(Ast_isBuiltin(_env, classTypeId)){
                    var scope *Ast_Scope
                    
                    {
                        _scope := classTypeInfo.FP.Get_scope(_env)
                        if _scope == nil{
                            Util_err(_env, _env.GetVM().String_format("%s.scope is nil", Lns_2DDD(classTypeInfo.FP.GetTxt(_env, nil, nil, nil))))
                        } else {
                            scope = _scope.(*Ast_Scope)
                        }
                    }
                    convLua_pickupTypeId(_env, classTypeInfo, true, _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( validChildrenSet.Get(classTypeInfo) == nil) &&
                        _env.SetStackVal( Lns_op_not(classTypeInfo.FP.Get_externalFlag(_env))) ).(bool))
                    if Lns_isCondTrue( convLua_checkExportTypeInfo(_env, classTypeInfo)){
                        self.FP.writeln(_env, "do")
                        self.FP.pushIndent(_env, nil)
                        self.FP.writeln(_env, _env.GetVM().String_format("local __classInfo%d = {}", Lns_2DDD(classTypeId)))
                        self.FP.writeln(_env, _env.GetVM().String_format("__typeId2ClassInfoMap[ %d ] = __classInfo%d", Lns_2DDD(classTypeId, classTypeId)))
                        {
                            __forsortCollection2 := scope.FP.Get_symbol2SymbolInfoMap(_env)
                            __forsortSorted2 := __forsortCollection2.CreateKeyListStr()
                            __forsortSorted2.Sort( _env, LnsItemKindStr, nil )
                            for _, _fieldName := range( __forsortSorted2.Items ) {
                                symbolInfo := __forsortCollection2.Items[ _fieldName ]
                                fieldName := _fieldName
                                var typeInfo *Ast_TypeInfo
                                typeInfo = symbolInfo.FP.Get_typeInfo(_env)
                                if _env.PopVal( _env.IncStack() ||
                                    _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) ||
                                    _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Var) ).(bool){
                                    if symbolInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
                                        self.FP.writeln(_env, _env.GetVM().String_format("__classInfo%d.%s = {", Lns_2DDD(classTypeId, fieldName)))
                                        self.FP.writeln(_env, _env.GetVM().String_format("  name='%s', staticFlag = %s, ", Lns_2DDD(fieldName, symbolInfo.FP.Get_staticFlag(_env))) + _env.GetVM().String_format("accessMode = %d, typeId = %d }", Lns_2DDD(symbolInfo.FP.Get_accessMode(_env), typeInfo.FP.Get_typeId(_env).Id)))
                                        convLua_pickupTypeId(_env, typeInfo, nil, nil)
                                    }
                                }
                            }
                        }
                        self.FP.popIndent(_env)
                        self.FP.writeln(_env, "end")
                    }
                }
            }
        }
    }
    self.FP.writeln(_env, "local __macroName2InfoMap = {}")
    self.FP.writeln(_env, "_moduleObj.__macroName2InfoMap = __macroName2InfoMap")
    for _, _macroDeclNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclMacroNodeList(_env).Items ) {
        macroDeclNode := _macroDeclNode.(Nodes_DeclMacroNodeDownCast).ToNodes_DeclMacroNode()
        var declInfo *Nodes_DeclMacroInfo
        declInfo = macroDeclNode.FP.Get_declInfo(_env)
        if declInfo.FP.Get_pubFlag(_env){
            var macroInfo *Nodes_MacroInfo
            macroInfo = Lns_unwrap( node.FP.Get_typeId2MacroInfo(_env).Get(macroDeclNode.FP.Get_expType(_env).FP.Get_typeId(_env))).(*Nodes_MacroInfo)
            var macroTypeInfo *Ast_TypeInfo
            macroTypeInfo = macroDeclNode.FP.Get_expType(_env)
            convLua_pickupTypeId(_env, macroTypeInfo, true, nil)
            self.FP.writeln(_env, "do")
            self.FP.pushIndent(_env, nil)
            self.FP.writeln(_env, "local info = {}")
            self.FP.writeln(_env, _env.GetVM().String_format("__macroName2InfoMap[ %d ] = info", Lns_2DDD(macroTypeInfo.FP.Get_typeId(_env).Id)))
            var pos Types_Position
            pos = macroDeclNode.FP.Get_pos(_env).Get_orgPos(_env)
            self.FP.writeln(_env, _env.GetVM().String_format("info.pos = {%d,%d}", Lns_2DDD(pos.LineNo, pos.Column)))
            self.FP.writeln(_env, _env.GetVM().String_format("info.name = %q", Lns_2DDD(declInfo.FP.Get_name(_env).Txt)))
            self.FP.WriteRaw(_env, "info.argList = {")
            for _index, _argNode := range( declInfo.FP.Get_argList(_env).Items ) {
                index := _index + 1
                argNode := _argNode
                if index != 1{
                    self.FP.WriteRaw(_env, ",")
                }
                self.FP.WriteRaw(_env, _env.GetVM().String_format("{name=%q,typeId=%d}", Lns_2DDD(argNode.FP.Get_name(_env).Txt, argNode.FP.Get_expType(_env).FP.Get_typeId(_env).Id)))
            }
            self.FP.writeln(_env, "}")
            self.FP.WriteRaw(_env, "info.symList = {")
            var firstFlag bool
            firstFlag = true
            {
                __forsortCollection4 := macroInfo.FP.Get_symbol2MacroValInfoMap(_env)
                __forsortSorted4 := __forsortCollection4.CreateKeyListStr()
                __forsortSorted4.Sort( _env, LnsItemKindStr, nil )
                for _, _name := range( __forsortSorted4.Items ) {
                    symInfo := __forsortCollection4.Items[ _name ]
                    name := _name
                    if firstFlag{
                        firstFlag = false
                    } else { 
                        self.FP.WriteRaw(_env, ",")
                    }
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("{name=%q,typeId=%d}", Lns_2DDD(name, symInfo.TypeInfo.FP.Get_typeId(_env).Id)))
                    convLua_pickupTypeId(_env, symInfo.TypeInfo, true, nil)
                }
            }
            self.FP.writeln(_env, "}")
            {
                _stmtBlock := declInfo.FP.Get_stmtBlock(_env)
                if !Lns_IsNil( _stmtBlock ) {
                    stmtBlock := _stmtBlock.(*Nodes_BlockNode)
                    var memStream *Util_memStream
                    memStream = NewUtil_memStream(_env)
                    var workFilter *convLua_ConvFilter
                    workFilter = NewconvLua_ConvFilter(_env, declInfo.FP.Get_name(_env).Txt, memStream.FP, NewUtil_NullOStream(_env).FP, ConvLua_ConvMode__Convert, false, Ast_headTypeInfo, self.processInfo, Ast_SymbolKind__Typ, self.builtinFunc, self.useLuneRuntime, self.targetLuaVer, self.enableTest, self.useIpairs, self.option)
                    workFilter.macroDepth = workFilter.macroDepth + 1
                    workFilter.FP.ProcessBlock(_env, stmtBlock, ConvLua_Opt2Stem(NewConvLua_Opt(_env, &node.Nodes_Node)))
                    workFilter.macroDepth = workFilter.macroDepth - 1
                    memStream.FP.Close(_env)
                    self.FP.writeln(_env, _env.GetVM().String_format("info.stmtBlock = %s", Lns_2DDD(Tokenizer_quoteStr(_env, memStream.FP.Get_txt(_env)))))
                }
            }
            self.FP.writeln(_env, "info.tokenList = {")
            var prevLineNo LnsInt
            prevLineNo = -1
            for _index, _token := range( declInfo.FP.Get_tokenList(_env).Items ) {
                index := _index + 1
                token := _token
                if index > 1{
                    self.FP.WriteRaw(_env, ",")
                }
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( prevLineNo != -1) &&
                    _env.SetStackVal( prevLineNo != token.Pos.LineNo) ).(bool)){
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("{%d,%s},", Lns_2DDD(Types_TokenKind__Dlmt, Tokenizer_quoteStr(_env, "\n"))))
                }
                prevLineNo = token.Pos.LineNo
                self.FP.WriteRaw(_env, _env.GetVM().String_format("{%d,%s}", Lns_2DDD(token.Kind, Tokenizer_quoteStr(_env, token.Txt))))
            }
            self.FP.writeln(_env, "}")
            self.FP.popIndent(_env)
            self.FP.writeln(_env, "end")
        }
    }
    self.FP.writeln(_env, "local __varName2InfoMap = {}")
    self.FP.writeln(_env, "_moduleObj.__varName2InfoMap = __varName2InfoMap")
    {
        __forsortCollection5 := self.pubVarName2InfoMap
        __forsortSorted5 := __forsortCollection5.CreateKeyListStr()
        __forsortSorted5.Sort( _env, LnsItemKindStr, nil )
        for _, _varName := range( __forsortSorted5.Items ) {
            varInfo := __forsortCollection5.Items[ _varName ]
            varName := _varName
            self.FP.writeln(_env, _env.GetVM().String_format("__varName2InfoMap.%s = {", Lns_2DDD(varName)))
            self.FP.writeln(_env, _env.GetVM().String_format("  name='%s', accessMode = %d, typeId = %s, mutable = %s }", Lns_2DDD(varName, varInfo.AccessMode, serializeInfo.FP.SerializeId(_env, varInfo.TypeInfo.FP.Get_typeId(_env)), true)))
            convLua_pickupTypeId(_env, varInfo.TypeInfo, true, nil)
        }
    }
    {
        __forsortCollection6 := self.pubFuncName2InfoMap
        __forsortSorted6 := __forsortCollection6.CreateKeyListStr()
        __forsortSorted6.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey6 := range( __forsortSorted6.Items ) {
            funcInfo := __forsortCollection6.Items[ ___forsortKey6 ]
            convLua_pickupTypeId(_env, funcInfo.TypeInfo, true, nil)
        }
    }
    for _, _aliasNode := range( node.FP.Get_nodeManager(_env).FP.GetAliasNodeList(_env).Items ) {
        aliasNode := _aliasNode.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
        convLua_pickupTypeId(_env, aliasNode.FP.Get_expType(_env), false, nil)
    }
    self.FP.writeln(_env, "local __typeInfoList = {}")
    self.FP.writeln(_env, "_moduleObj.__typeInfoList = __typeInfoList")
    var listIndex LnsInt
    listIndex = 1
    var convLua_outputDepend func(_env *LnsEnv, typeInfo *Ast_TypeInfo,moduleTypeInfo *Ast_TypeInfo) LnsInt
    convLua_outputDepend = func(_env *LnsEnv, typeInfo *Ast_TypeInfo,moduleTypeInfo *Ast_TypeInfo) LnsInt {
        var typeId *Ast_IdInfo
        typeId = typeInfo.FP.Get_typeId(_env)
        if self.processInfo.FP.Get_orgInfo(_env) == typeId.FP.Get_processInfo(_env){
            return convLua_ExportIdKind__Normal
        }
        if typeId.FP.IsSwichingId(_env){
            return convLua_ExportIdKind__Discarded
        }
        if Ast_isExtId(_env, typeInfo){
            return convLua_ExportIdKind__Normal
        }
        return convLua_ExportIdKind__Discarded
    }
    var wroteTypeIdSet *LnsSet2_[*Ast_IdInfo]
    wroteTypeIdSet = NewLnsSet2_[*Ast_IdInfo]([]*Ast_IdInfo{})
    var convLua_outputTypeInfo func(_env *LnsEnv, typeInfo *Ast_TypeInfo)
    convLua_outputTypeInfo = func(_env *LnsEnv, typeInfo *Ast_TypeInfo) {
        if Ast_isBuiltin(_env, typeInfo.FP.Get_typeId(_env).Id){
            return 
        }
        if Lns_op_not(Ast_TypeInfo_hasParent(_env, typeInfo)){
            return 
        }
        if _switch1 := typeInfo.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Class || _switch1 == Ast_TypeInfoKind__IF {
            if _switch0 := typeInfo.FP.Get_accessMode(_env); _switch0 == Ast_AccessMode__Pub || _switch0 == Ast_AccessMode__Pro || _switch0 == Ast_AccessMode__Global {
            } else {
                Util_errorLog(_env, _env.GetVM().String_format("skip: %s %s", Lns_2DDD(typeInfo.FP.Get_accessMode(_env), self.FP.getFullName(_env, typeInfo))))
                return 
            }
        } else if _switch1 == Ast_TypeInfoKind__Module {
            if Lns_op_not(typeInfo.FP.HasRouteNamespaceFrom(_env, self.moduleTypeInfo)){
                return 
            }
        }
        var typeId *Ast_IdInfo
        typeId = typeInfo.FP.Get_typeId(_env)
        if wroteTypeIdSet.Has(typeId){
            return 
        }
        wroteTypeIdSet.Add(typeId)
        self.FP.WriteRaw(_env, _env.GetVM().String_format("__typeInfoList[%d] = ", Lns_2DDD(listIndex)))
        listIndex = listIndex + 1
        var validChildren *LnsMap2_[*Ast_IdInfo,*Ast_TypeInfo]
        
        {
            _validChildren := validChildrenSet.Get(typeInfo)
            if _validChildren == nil{
                validChildren = typeId2TypeInfo
            } else {
                validChildren = _validChildren.(*LnsMap2_[*Ast_IdInfo,*Ast_TypeInfo])
            }
        }
        typeInfo.FP.Serialize(_env, self.FP, NewAst_SerializeInfo(_env, importProcessInfo2Index, validChildren))
    }
    for _, _typeInfo := range( self.pubEnumId2EnumTypeInfo.Items ) {
        typeInfo := _typeInfo
        convLua_pickupTypeId(_env, &typeInfo.Ast_TypeInfo, true, nil)
    }
    for _, _typeInfo := range( self.pubAlgeId2AlgeTypeInfo.Items ) {
        typeInfo := _typeInfo
        convLua_pickupTypeId(_env, &typeInfo.Ast_TypeInfo, true, nil)
        {
            __forsortCollection7 := typeInfo.FP.Get_valInfoMap(_env)
            __forsortSorted7 := __forsortCollection7.CreateKeyListStr()
            __forsortSorted7.Sort( _env, LnsItemKindStr, nil )
            for _, ___forsortKey7 := range( __forsortSorted7.Items ) {
                valInfo := __forsortCollection7.Items[ ___forsortKey7 ]
                for _, _valType := range( valInfo.FP.Get_typeList(_env).Items ) {
                    valType := _valType
                    convLua_pickupTypeId(_env, valType, true, nil)
                }
            }
        }
    }
    for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetAliasNodeList(_env).Items ) {
        workNode := _workNode.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
        if Ast_isPubToExternal(_env, workNode.FP.Get_expType(_env).FP.Get_accessMode(_env)){
            var aliasType *Ast_AliasTypeInfo
            aliasType = Lns_unwrap( Ast_AliasTypeInfoDownCastF(workNode.FP.Get_expType(_env).FP)).(*Ast_AliasTypeInfo)
            var aliasSrcType *Ast_TypeInfo
            aliasSrcType = aliasType.FP.Get_aliasSrc(_env)
            convLua_pickupTypeId(_env, &aliasType.Ast_TypeInfo, true, nil)
            convLua_pickupTypeId(_env, aliasSrcType, true, nil)
        }
    }
    self.FP.writeln(_env, "local __dependIdMap = {}")
    self.FP.writeln(_env, "_moduleObj.__dependIdMap = __dependIdMap")
    var exportNeedModuleTypeInfo *LnsSet2_[*Ast_TypeInfo]
    exportNeedModuleTypeInfo = NewLnsSet2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    {
        var module2TypeList *LnsMap2_[LnsInt,*LnsMap2_[LnsInt,*Ast_TypeInfo]]
        module2TypeList = NewLnsMap2_[LnsInt,*LnsMap2_[LnsInt,*Ast_TypeInfo]]( map[LnsInt]*LnsMap2_[LnsInt,*Ast_TypeInfo]{})
        for _, _typeInfo := range( typeId2TypeInfo.Items ) {
            typeInfo := _typeInfo
            var modIndex LnsInt
            modIndex = Lns_unwrap( importProcessInfo2Index.Get(typeInfo.FP.Get_typeId(_env).FP.Get_processInfo(_env))).(LnsInt)
            var _map *LnsMap2_[LnsInt,*Ast_TypeInfo]
            
            {
                __map := module2TypeList.Get(modIndex)
                if __map == nil{
                    _map = NewLnsMap2_[LnsInt,*Ast_TypeInfo]( map[LnsInt]*Ast_TypeInfo{})
                    module2TypeList.Set(modIndex,_map)
                } else {
                    _map = __map.(*LnsMap2_[LnsInt,*Ast_TypeInfo])
                }
            }
            _map.Set(typeInfo.FP.Get_typeId(_env).Id,typeInfo)
        }
        {
            __forsortCollection9 := module2TypeList
            __forsortSorted9 := __forsortCollection9.CreateKeyListInt()
            __forsortSorted9.Sort( _env, LnsItemKindInt, nil )
            for _, ___forsortKey9 := range( __forsortSorted9.Items ) {
                _map := __forsortCollection9.Items[ ___forsortKey9 ]
                {
                    __forsortCollection8 := _map
                    __forsortSorted8 := __forsortCollection8.CreateKeyListInt()
                    __forsortSorted8.Sort( _env, LnsItemKindInt, nil )
                    for _, ___forsortKey8 := range( __forsortSorted8.Items ) {
                        typeInfo := __forsortCollection8.Items[ ___forsortKey8 ]
                        var moduleTypeInfo *Ast_TypeInfo
                        moduleTypeInfo = typeInfo.FP.GetModule(_env)
                        exportNeedModuleTypeInfo.Add(moduleTypeInfo)
                        var ret LnsInt
                        ret = convLua_outputDepend(_env, typeInfo, moduleTypeInfo)
                        if ret == convLua_ExportIdKind__Normal{
                            convLua_outputTypeInfo(_env, typeInfo)
                        }
                    }
                }
            }
        }
    }
    for _moduleTypeInfo := range( node.FP.Get_useModuleMacroSet(_env).Items ) {
        moduleTypeInfo := _moduleTypeInfo
        exportNeedModuleTypeInfo.Add(moduleTypeInfo)
    }
    self.FP.writeln(_env, "local __dependModuleMap = {}")
    self.FP.writeln(_env, "_moduleObj.__dependModuleMap = __dependModuleMap")
    {
        __forsortCollection10 := importNameMap
        __forsortSorted10 := __forsortCollection10.CreateKeyListStr()
        __forsortSorted10.Sort( _env, LnsItemKindStr, nil )
        for _, _name := range( __forsortSorted10.Items ) {
            exportInfo := __forsortCollection10.Items[ _name ]
            name := _name
            var moduleTypeInfo *Ast_TypeInfo
            moduleTypeInfo = exportInfo.FP.Get_moduleTypeInfo(_env)
            self.FP.writeln(_env, _env.GetVM().String_format("__dependModuleMap[ '%s' ] = { typeId = %d, use = %s, buildId = %q }", Lns_2DDD(name, Lns_unwrap( importModuleType2Index.Get(moduleTypeInfo)).(LnsInt), exportNeedModuleTypeInfo.Has(moduleTypeInfo), (Lns_unwrap( node.FP.Get_importModule2moduleInfo(_env).Get(moduleTypeInfo)).(*FrontInterface_ExportInfo)).FP.Get_moduleId(_env).FP.Get_idStr(_env))))
        }
    }
    self.FP.WriteRaw(_env, "_moduleObj.__subModuleMap = {")
    {
        var firstFlag bool
        firstFlag = true
        for _, _subfileNode := range( node.FP.Get_nodeManager(_env).FP.GetSubfileNodeList(_env).Items ) {
            subfileNode := _subfileNode.(Nodes_SubfileNodeDownCast).ToNodes_SubfileNode()
            {
                _usePath := subfileNode.FP.Get_usePath(_env)
                if !Lns_IsNil( _usePath ) {
                    usePath := _usePath.(string)
                    if firstFlag{
                        firstFlag = false
                    } else { 
                        self.FP.WriteRaw(_env, ",")
                    }
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("%q", Lns_2DDD(usePath)))
                }
            }
        }
    }
    self.FP.writeln(_env, "}")
    self.FP.WriteRaw(_env, "_moduleObj.__moduleHierarchy = {")
    var workType *Ast_TypeInfo
    workType = node.FP.Get_moduleTypeInfo(_env)
    for Ast_TypeInfo_hasParent(_env, workType) {
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%d,", Lns_2DDD(workType.FP.Get_typeId(_env).Id)))
        workType = workType.FP.Get_parentInfo(_env)
    }
    self.FP.writeln(_env, "}")
    var moduleTypeInfo *Ast_TypeInfo
    moduleTypeInfo = self.moduleTypeInfo
    var moduleSymbolKind LnsInt
    moduleSymbolKind = Ast_SymbolKind__Typ
    {
        __exp := node.FP.Get_provideNode(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ProvideNode)
            moduleTypeInfo = _exp.FP.Get_symbol(_env).FP.Get_typeInfo(_env)
            moduleSymbolKind = _exp.FP.Get_symbol(_env).FP.Get_kind(_env)
        }
    }
    self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.__moduleTypeId = %d", Lns_2DDD(moduleTypeInfo.FP.Get_typeId(_env).Id)))
    self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.__moduleSymbolKind = %d", Lns_2DDD(moduleSymbolKind)))
    self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.__moduleMutable = %s", Lns_2DDD(Ast_TypeInfo_isMut(_env, moduleTypeInfo))))
    self.FP.writeln(_env, "----- meta -----")
    self.FP.writeln(_env, "return _moduleObj")
    self.outMetaFlag = false
}
// 1055: decl @lune.@base.@convLua.ConvFilter.outputMainDirect
func (self *convLua_ConvFilter) outputMainDirect(_env *LnsEnv) {
    var code string
    code = _env.GetVM().String_format("do\n   local loaded, mess = _lune.%s( [=[\nif _lune and _lune._shebang then\n  return nil\nelse\n  return arg\nend\n]=] )\n   if loaded ~= nil then\n      local args = loaded(  )\n      do\n         local obj = (args )\n         if obj ~= nil then\n            local work = obj\n            local argList = {\"\"}\n            do\n               local _exp = work[0]\n               if _exp ~= nil then\n                  argList[1] = _exp\n               end\n            end\n            for key, val in pairs( work ) do\n               if key > 0 then\n                  table.insert( argList, val )\n               end\n            end\n            __main( argList )\n         else\n            -- print( \"via lnsc\" )\n         end\n      end\n   else\n      error( mess )\n   end\nend\n", Lns_2DDD(self.targetLuaVer.FP.Get_loadStrFuncName(_env)))
    self.FP.writeln(_env, code)
}
// 1096: decl @lune.@base.@convLua.ConvFilter.processRoot
func (self *convLua_ConvFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convLua.ConvFilter.processRoot"
    Log_log(_env, Log_Level__Log, __func__, 1099, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("streamName: %s, enableTest: %s", Lns_2DDD(self.streamName, self.enableTest))
    }))
    
    self.FP.writeln(_env, _env.GetVM().String_format("--%s", Lns_2DDD(self.streamName)))
    self.needModuleObj = Nodes_ProvideNode2Stem(node.FP.Get_provideNode(_env)) == nil
    if self.needModuleObj{
        self.FP.writeln(_env, "local _moduleObj = {}")
    }
    if self.enableTest{
        self.FP.writeln(_env, "_moduleObj.__enableTest = true")
    }
    self.FP.writeln(_env, _env.GetVM().String_format("local __mod__ = '%s'", Lns_2DDD(node.FP.Get_moduleTypeInfo(_env).FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), nil))))
    var luneSymbol string
    luneSymbol = _env.GetVM().String_format("_lune%d", Lns_2DDD(Ver_luaModVersion))
    {
        _runtime := self.useLuneRuntime
        if !Lns_IsNil( _runtime ) {
            runtime := _runtime.(string)
            self.FP.writeln(_env, _env.GetVM().String_format("local _lune = require( \"%s\" )", Lns_2DDD(runtime)))
        } else {
            self.FP.writeln(_env, "local _lune = {}")
            self.FP.writeln(_env, _env.GetVM().String_format("if %s then\n   _lune = %s\nend", Lns_2DDD(luneSymbol, luneSymbol)))
            if node.FP.Get_luneHelperInfo(_env).UseResult{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Result))
            }
            if node.FP.Get_luneHelperInfo(_env).UseError{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Error))
            }
            if node.FP.Get_luneHelperInfo(_env).UseAlge{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Alge))
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__AlgeMapping))
            }
            if node.FP.Get_luneHelperInfo(_env).UseSet{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__SetOp))
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__SetMapping))
            }
            if Lns_isCondTrue( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( node.FP.Get_luneHelperInfo(_env).UseUnpack) &&
                _env.SetStackVal( Lns_op_not(self.targetLuaVer.FP.Get_hasTableUnpack(_env))) ).(bool))){
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Unpack))
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( node.FP.Get_luneHelperInfo(_env).UseLoad) ||
                _env.SetStackVal( self.option.FP.Get_mainModule(_env) == Lns_car(_env.GetVM().String_gsub(self.FP.getCanonicalName(_env, self.moduleTypeInfo, false),"@", "")).(string)) ).(bool){
                self.FP.writeln(_env, self.targetLuaVer.FP.GetLoadCode(_env))
            }
            if node.FP.Get_luneHelperInfo(_env).UseNilAccess{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__NilAcc))
            }
            if node.FP.Get_luneHelperInfo(_env).UseUnwrapExp{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Unwrap))
            }
            if node.FP.Get_luneHelperInfo(_env).HasMappingClassDef{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Mapping))
            }
            if node.FP.Get_nodeManager(_env).FP.GetImportNodeList(_env).Len() != 0{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__LoadModule))
            }
            if node.FP.Get_nodeManager(_env).FP.GetExpCastNodeList(_env).Len() != 0{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__InstanceOf))
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Cast))
            }
            if node.FP.Get_luneHelperInfo(_env).UseLazyLoad{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__LazyLoad))
            }
            if node.FP.Get_luneHelperInfo(_env).UseLazyRequire{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__LazyRequire))
            }
            if node.FP.Get_luneHelperInfo(_env).UseRun{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Run))
            }
            if node.FP.Get_luneHelperInfo(_env).UseStrReplace{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__StrReplace))
            }
        }
    }
    self.FP.writeln(_env, _env.GetVM().String_format("if not %s then\n   %s = _lune\nend", Lns_2DDD(luneSymbol, luneSymbol)))
    for _, _importNode := range( node.FP.Get_nodeManager(_env).FP.GetImportNodeList(_env).Items ) {
        importNode := _importNode.(Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        var info *Nodes_ImportInfo
        info = importNode.FP.Get_info(_env)
        if info.FP.Get_lazy(_env) != Nodes_LazyLoad__Off{
            self.moduleType2SymbolMap.Set(info.FP.Get_moduleTypeInfo(_env),info.FP.Get_symbolInfo(_env))
        }
    }
    if Lns_isCondTrue( _env.NilAccFin(_env.NilAccPush(node.FP.Get_moduleScope(_env).FP.GetSymbolInfoChild(_env, "_")) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_SymbolInfo).FP.Get_posForLatestMod(_env)}))){
        self.FP.writeln(_env, "local _")
    }
    var children *LnsList2_[*Nodes_Node]
    children = node.FP.Get_children(_env)
    for _, _child := range( children.Items ) {
        child := _child
        convLua_filter_9_(_env, child, self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
    }
    if self.option.FP.Get_mainModule(_env) == Lns_car(_env.GetVM().String_gsub(self.FP.getCanonicalName(_env, self.moduleTypeInfo, false),"@", "")).(string){
        self.FP.outputMainDirect(_env)
    }
    {
        __exp := node.FP.Get_provideNode(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ProvideNode)
            self.FP.WriteRaw(_env, "return ")
            self.FP.WriteRaw(_env, _exp.FP.Get_symbol(_env).FP.Get_name(_env))
            self.FP.writeln(_env, "")
        } else {
            self.FP.writeln(_env, "return _moduleObj")
        }
    }
    self.FP.outputMeta(_env, node)
}
// 1227: decl @lune.@base.@convLua.ConvFilter.processSubfile
func (self *convLua_ConvFilter) ProcessSubfile(_env *LnsEnv, node *Nodes_SubfileNode,_opt LnsAny) {
}
// 1232: decl @lune.@base.@convLua.ConvFilter.processRequest
func (self *convLua_ConvFilter) ProcessRequest(_env *LnsEnv, node *Nodes_RequestNode,_opt LnsAny) {
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
}
// 1240: decl @lune.@base.@convLua.ConvFilter.processAsyncLock
func (self *convLua_ConvFilter) ProcessAsyncLock(_env *LnsEnv, node *Nodes_AsyncLockNode,_opt LnsAny) {
    if Lns_isCondTrue( node.FP.Get_returnTypeList(_env)){
        self.FP.writeln(_env, "do")
        self.FP.pushIndent(_env, nil)
        self.FP.writeln(_env, "function _lock()")
        self.FP.pushIndent(_env, nil)
        convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "end")
        self.FP.writeln(_env, "local _val = { _lock() }")
        self.FP.writeln(_env, "if #_val ~= 0 then return table.unpack( _val ) end")
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "end")
    } else { 
        convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    }
    if node.FP.Get_hasAsyncLockBreak(_env){
        self.FP.writeln(_env, _env.GetVM().String_format("if %s then break end", Lns_2DDD(convLua_asyncLockBreakSym)))
    }
}
// 1264: decl @lune.@base.@convLua.ConvFilter.processBlockSub
func (self *convLua_ConvFilter) ProcessBlockSub(_env *LnsEnv, node *Nodes_BlockNode,_opt LnsAny) {
    var word string
    word = ""
    if _switch0 := node.FP.Get_blockKind(_env); _switch0 == Nodes_BlockKind__If || _switch0 == Nodes_BlockKind__Elseif {
        word = "then"
    } else if _switch0 == Nodes_BlockKind__Else {
        word = ""
    } else if _switch0 == Nodes_BlockKind__While {
        word = "do"
    } else if _switch0 == Nodes_BlockKind__Repeat {
        word = ""
    } else if _switch0 == Nodes_BlockKind__For {
        word = "do"
    } else if _switch0 == Nodes_BlockKind__Apply {
        word = "do"
    } else if _switch0 == Nodes_BlockKind__Foreach {
        word = "do"
    } else if _switch0 == Nodes_BlockKind__Macro {
        word = ""
    } else if _switch0 == Nodes_BlockKind__Func {
        word = ""
    } else if _switch0 == Nodes_BlockKind__Default {
        word = ""
    } else if _switch0 == Nodes_BlockKind__Block || _switch0 == Nodes_BlockKind__AsyncLock {
        word = "do"
    } else if _switch0 == Nodes_BlockKind__LetUnwrap {
        word = ""
    } else if _switch0 == Nodes_BlockKind__LetUnwrapThenDo {
        word = ""
    } else if _switch0 == Nodes_BlockKind__IfUnwrap {
        word = ""
    }
    self.FP.writeln(_env, word)
    self.FP.pushIndent(_env, nil)
    if node.FP.Get_hasAsyncLockBreak(_env){
        self.FP.writeln(_env, _env.GetVM().String_format("local %s = false", Lns_2DDD(convLua_asyncLockBreakSym)))
    }
    if Lns_isCondTrue( _env.NilAccFin(_env.NilAccPush(node.FP.Get_scope(_env).FP.GetSymbolInfoChild(_env, "_")) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_SymbolInfo).FP.Get_posForLatestMod(_env)}))){
        self.FP.writeln(_env, "local _")
    }
    var stmtList *LnsList2_[*Nodes_Node]
    stmtList = node.FP.Get_stmtList(_env)
    for _, _statement := range( stmtList.Items ) {
        statement := _statement
        convLua_filter_9_(_env, statement, self, &node.Nodes_Node)
    }
    self.FP.popIndent(_env)
    if _switch1 := node.FP.Get_blockKind(_env); _switch1 == Nodes_BlockKind__Block || _switch1 == Nodes_BlockKind__AsyncLock {
        self.FP.writeln(_env, "end")
    }
}
// 1339: decl @lune.@base.@convLua.ConvFilter.processLoadRuntime
func (self *convLua_ConvFilter) processLoadRuntime(_env *LnsEnv) {
    {
        __exp := self.useLuneRuntime
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(string)
            self.FP.writeln(_env, _env.GetVM().String_format("local _lune = require( \"%s\" )", Lns_2DDD(_exp)))
        } else {
            self.FP.writeln(_env, _env.GetVM().String_format("local _lune = require( \"%s\" )", Lns_2DDD(Option_getRuntimeModule(_env))))
        }
    }
}
// 1349: decl @lune.@base.@convLua.ConvFilter.processScope
func (self *convLua_ConvFilter) ProcessScope(_env *LnsEnv, node *Nodes_ScopeNode,_opt LnsAny) {
    if node.FP.Get_scopeKind(_env) == Nodes_ScopeKind__Root{
        self.FP.processLoadRuntime(_env)
    }
    convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
}
// 1358: decl @lune.@base.@convLua.ConvFilter.processCondRet
func (self *convLua_ConvFilter) ProcessCondRet(_env *LnsEnv, node *Nodes_CondRetNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("_cond%s", Lns_2DDD(node.FP.Get_order(_env))))
}
// 1364: decl @lune.@base.@convLua.ConvFilter.outputCondRet
func (self *convLua_ConvFilter) outputCondRet(_env *LnsEnv, node *Nodes_CondRetNode) {
    var symName string
    symName = _env.GetVM().String_format("_cond%d", Lns_2DDD(node.FP.Get_order(_env)))
    if _switch0 := node.FP.Get_condKind(_env); _switch0 == Nodes_CondRetKind__Nilable {
        self.FP.Write(_env, _env.GetVM().String_format("local %s = ", Lns_2DDD(symName)))
        convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
        self.FP.writeln(_env, _env.GetVM().String_format("if %s == nil then return nil end", Lns_2DDD(symName)))
    } else if _switch0 == Nodes_CondRetKind__Ret {
        self.FP.writeln(_env, _env.GetVM().String_format("local %s", Lns_2DDD(symName)))
        self.FP.writeln(_env, "do")
        self.FP.pushIndent(_env, nil)
        self.FP.Write(_env, "local _matchExp = ")
        convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
        self.FP.writeln(_env, "if _matchExp[1] == _lune.Result.Err[1] then")
        self.FP.pushIndent(_env, nil)
        self.FP.writeln(_env, "return _matchExp")
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "elseif _matchExp[1] == _lune.Result.Ok[1] then")
        self.FP.pushIndent(_env, nil)
        self.FP.writeln(_env, _env.GetVM().String_format("%s = _matchExp[2][1]", Lns_2DDD(symName)))
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "end")
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "end")
    } else if _switch0 == Nodes_CondRetKind__Two {
        self.FP.writeln(_env, _env.GetVM().String_format("local %s", Lns_2DDD(symName)))
        self.FP.writeln(_env, "do")
        self.FP.pushIndent(_env, nil)
        self.FP.writeln(_env, "local __condWork")
        self.FP.Write(_env, _env.GetVM().String_format("%s, __condWork = ", Lns_2DDD(symName)))
        convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
        self.FP.writeln(_env, _env.GetVM().String_format("if %s == nil then return nil, __condWork end", Lns_2DDD(symName)))
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "end")
    }
}
// 1409: decl @lune.@base.@convLua.ConvFilter.processCondRetList
func (self *convLua_ConvFilter) ProcessCondRetList(_env *LnsEnv, node *Nodes_CondRetListNode,_opt LnsAny) {
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
}
// 1415: decl @lune.@base.@convLua.ConvFilter.outputCondRetInfo
func (self *convLua_ConvFilter) outputCondRetInfo(_env *LnsEnv, info *Nodes_CondRetInfo) {
    for _, _condRetNode := range( info.FP.Get_list(_env).Items ) {
        condRetNode := _condRetNode
        self.FP.outputCondRet(_env, condRetNode)
    }
}
// 1421: decl @lune.@base.@convLua.ConvFilter.processStmtExp
func (self *convLua_ConvFilter) ProcessStmtExp(_env *LnsEnv, node *Nodes_StmtExpNode,_opt LnsAny) {
    {
        _condRetListNode := Nodes_CondRetListNodeDownCastF(node.FP.Get_exp(_env).FP)
        if !Lns_IsNil( _condRetListNode ) {
            condRetListNode := _condRetListNode.(*Nodes_CondRetListNode)
            self.FP.outputCondRetInfo(_env, condRetListNode.FP.Get_info(_env))
        }
    }
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
}
// 1432: decl @lune.@base.@convLua.ConvFilter.processDeclEnum
func (self *convLua_ConvFilter) ProcessDeclEnum(_env *LnsEnv, node *Nodes_DeclEnumNode,_opt LnsAny) {
    var access string
    access = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_accessMode(_env) == Ast_AccessMode__Global) &&
        _env.SetStackVal( "") ||
        _env.SetStackVal( "local ") ).(string)
    var enumFullName string
    enumFullName = node.FP.Get_name(_env).Txt
    var typeInfo *Ast_EnumTypeInfo
    typeInfo = Lns_unwrap( Ast_EnumTypeInfoDownCastF(node.FP.Get_expType(_env).FP.Get_aliasSrc(_env).FP)).(*Ast_EnumTypeInfo)
    var parentInfo *Ast_TypeInfo
    parentInfo = typeInfo.FP.Get_parentInfo(_env)
    var isTopNS bool
    isTopNS = true
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Ast_TypeInfo_hasParent(_env, &typeInfo.Ast_TypeInfo)) &&
        _env.SetStackVal( parentInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ).(bool)){
        enumFullName = _env.GetVM().String_format("%s.%s", Lns_2DDD(self.FP.getFullName(_env, parentInfo), enumFullName))
        access = ""
        isTopNS = false
    }
    self.FP.writeln(_env, _env.GetVM().String_format("%s%s = {}", Lns_2DDD(access, enumFullName)))
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( isTopNS) &&
        _env.SetStackVal( node.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) ).(bool)){
        if self.needModuleObj{
            self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.%s = %s", Lns_2DDD(enumFullName, enumFullName)))
        }
    }
    if typeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
        self.pubEnumId2EnumTypeInfo.Set(typeInfo.FP.Get_typeId(_env),typeInfo)
    }
    self.FP.writeln(_env, _env.GetVM().String_format("%s._val2NameMap = {}", Lns_2DDD(enumFullName)))
    self.FP.writeln(_env, _env.GetVM().String_format("function %s:_getTxt( val )\n   local name = self._val2NameMap[ val ]\n   if name then\n      return string.format( \"%s.%%s\", name )\n   end\n   return string.format( \"illegal val -- %%s\", val )\nend\nfunction %s._from( val )\n   if %s._val2NameMap[ val ] then\n      return val\n   end\n   return nil\nend\n    ", Lns_2DDD(enumFullName, enumFullName, enumFullName, enumFullName)))
    self.FP.writeln(_env, _env.GetVM().String_format("%s.__allList = {}\nfunction %s.get__allList()\n   return %s.__allList\nend\n", Lns_2DDD(enumFullName, enumFullName, enumFullName)))
    for _index, _valName := range( node.FP.Get_valueNameList(_env).Items ) {
        index := _index + 1
        valName := _valName
        var valInfo *Ast_EnumValInfo
        valInfo = Lns_unwrap( typeInfo.FP.GetEnumValInfo(_env, valName.Txt)).(*Ast_EnumValInfo)
        var valTxt string
        valTxt = _env.GetVM().String_format("%s", Lns_2DDD(Ast_getEnumLiteralVal(_env, valInfo.FP.Get_val(_env))))
        if typeInfo.FP.Get_valTypeInfo(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeString, nil, nil){
            valTxt = _env.GetVM().String_format("'%s'", Lns_2DDD(Ast_getEnumLiteralVal(_env, valInfo.FP.Get_val(_env))))
        }
        self.FP.writeln(_env, _env.GetVM().String_format("%s.%s = %s", Lns_2DDD(enumFullName, valName.Txt, valTxt)))
        self.FP.writeln(_env, _env.GetVM().String_format("%s._val2NameMap[%s] = '%s'", Lns_2DDD(enumFullName, valTxt, valName.Txt)))
        self.FP.writeln(_env, _env.GetVM().String_format("%s.__allList[%d] = %s.%s", Lns_2DDD(enumFullName, index, enumFullName, valName.Txt)))
    }
}
// 1514: decl @lune.@base.@convLua.ConvFilter.getMapInfo
func (self *convLua_ConvFilter) getMapInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo)(string, bool, string) {
    var nonnilableType *Ast_TypeInfo
    nonnilableType = typeInfo.FP.Get_srcTypeInfo(_env)
    if typeInfo.FP.Get_nilable(_env){
        nonnilableType = typeInfo.FP.Get_nonnilableType(_env)
    }
    var child string
    child = "{}"
    var funcTxt string
    funcTxt = ""
    if _switch1 := nonnilableType.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Stem {
        funcTxt = "_lune._toStem"
    } else if _switch1 == Ast_TypeInfoKind__Class || _switch1 == Ast_TypeInfoKind__IF {
        if Lns_op_not(nonnilableType.FP.Equals(_env, self.processInfo, Ast_builtinTypeString, nil, nil)){
            if Lns_car(Ast_NormalTypeInfo_isAvailableMapping(_env, self.processInfo, nonnilableType, NewLnsMap2_[*Ast_TypeInfo,bool]( map[*Ast_TypeInfo]bool{}))).(bool){
                funcTxt = _env.GetVM().String_format("%s._fromMap", Lns_2DDD(self.FP.getFullName(_env, nonnilableType)))
                if convLua_isGenericType_12_(_env, nonnilableType){
                    var memStream *Util_memStream
                    memStream = NewUtil_memStream(_env)
                    self.FP.outputAlter2MapFunc(_env, memStream.FP, nonnilableType.FP.CreateAlt2typeMap(_env, false))
                    child = memStream.FP.Get_txt(_env)
                }
            } else { 
                funcTxt = "nil"
            }
        } else { 
            funcTxt = "_lune._toStr"
        }
    } else if _switch1 == Ast_TypeInfoKind__Enum || _switch1 == Ast_TypeInfoKind__Alge {
        funcTxt = _env.GetVM().String_format("%s._from", Lns_2DDD(self.FP.getFullName(_env, nonnilableType)))
    } else if _switch1 == Ast_TypeInfoKind__Prim {
        if _switch0 := nonnilableType; _switch0 == Ast_builtinTypeInt {
            funcTxt = "_lune._toInt"
        } else if _switch0 == Ast_builtinTypeReal {
            funcTxt = "_lune._toReal"
        } else if _switch0 == Ast_builtinTypeBool {
            funcTxt = "_lune._toBool"
        } else {
            Util_err(_env, _env.GetVM().String_format("unknown type -- %s", Lns_2DDD(nonnilableType.FP.GetTxt(_env, nil, nil, nil))))
        }
    } else if _switch1 == Ast_TypeInfoKind__Map {
        funcTxt = "_lune._toMap"
        var itemList *LnsList2_[*Ast_TypeInfo]
        itemList = nonnilableType.FP.Get_itemTypeInfoList(_env)
        var keyFuncTxt string
        var keyNilable bool
        var keyChild string
        keyFuncTxt,keyNilable,keyChild = self.FP.getMapInfo(_env, itemList.GetAt(1))
        var valFuncTxt string
        var valNilable bool
        var valChild string
        valFuncTxt,valNilable,valChild = self.FP.getMapInfo(_env, itemList.GetAt(2))
        child = _env.GetVM().String_format("{ { func = %s, nilable = %s, child = %s }, \n", Lns_2DDD(keyFuncTxt, keyNilable, keyChild)) + _env.GetVM().String_format("{ func = %s, nilable = %s, child = %s } }", Lns_2DDD(valFuncTxt, valNilable, valChild))
    } else if _switch1 == Ast_TypeInfoKind__Set {
        funcTxt = "_lune._toSet"
        var itemList *LnsList2_[*Ast_TypeInfo]
        itemList = nonnilableType.FP.Get_itemTypeInfoList(_env)
        var valFuncTxt string
        var valNilable bool
        var valChild string
        valFuncTxt,valNilable,valChild = self.FP.getMapInfo(_env, itemList.GetAt(1))
        child = _env.GetVM().String_format("{ func = %s, nilable = %s, child = %s }", Lns_2DDD(valFuncTxt, valNilable, valChild))
    } else if _switch1 == Ast_TypeInfoKind__List || _switch1 == Ast_TypeInfoKind__Array {
        funcTxt = "_lune._toList"
        var itemList *LnsList2_[*Ast_TypeInfo]
        itemList = nonnilableType.FP.Get_itemTypeInfoList(_env)
        var valFuncTxt string
        var valNilable bool
        var valChild string
        valFuncTxt,valNilable,valChild = self.FP.getMapInfo(_env, itemList.GetAt(1))
        child = _env.GetVM().String_format("{ { func = %s, nilable = %s, child = %s } }", Lns_2DDD(valFuncTxt, valNilable, valChild))
    } else if _switch1 == Ast_TypeInfoKind__Alternate {
        var prefix string
        prefix = _env.GetVM().String_format("obj.__alt2mapFunc.%s", Lns_2DDD(nonnilableType.FP.Get_rawTxt(_env)))
        funcTxt = _env.GetVM().String_format("%s.func", Lns_2DDD(prefix))
        child = _env.GetVM().String_format("%s.child", Lns_2DDD(prefix))
    }
    return funcTxt, typeInfo.FP.Get_nilable(_env), child
}
// 1604: decl @lune.@base.@convLua.ConvFilter.processDeclAlge
func (self *convLua_ConvFilter) ProcessDeclAlge(_env *LnsEnv, node *Nodes_DeclAlgeNode,_opt LnsAny) {
    var access string
    access = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_accessMode(_env) == Ast_AccessMode__Global) &&
        _env.SetStackVal( "") ||
        _env.SetStackVal( "local ") ).(string)
    var algeFullName string
    algeFullName = node.FP.Get_algeType(_env).FP.Get_rawTxt(_env)
    var typeInfo *Ast_AlgeTypeInfo
    typeInfo = Lns_unwrap( Ast_AlgeTypeInfoDownCastF(node.FP.Get_expType(_env).FP)).(*Ast_AlgeTypeInfo)
    var parentInfo *Ast_TypeInfo
    parentInfo = typeInfo.FP.Get_parentInfo(_env)
    var isTopNS bool
    isTopNS = true
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Ast_TypeInfo_hasParent(_env, &typeInfo.Ast_TypeInfo)) &&
        _env.SetStackVal( parentInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ).(bool)){
        algeFullName = _env.GetVM().String_format("%s.%s", Lns_2DDD(self.FP.getFullName(_env, parentInfo), algeFullName))
        access = ""
        isTopNS = false
    }
    self.FP.writeln(_env, _env.GetVM().String_format("%s%s = {}", Lns_2DDD(access, algeFullName)))
    self.FP.writeln(_env, _env.GetVM().String_format("%s._name2Val = {}", Lns_2DDD(algeFullName)))
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( isTopNS) &&
        _env.SetStackVal( node.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) ).(bool)){
        if self.needModuleObj{
            self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.%s = %s", Lns_2DDD(algeFullName, algeFullName)))
        }
    }
    if typeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
        self.pubAlgeId2AlgeTypeInfo.Set(typeInfo.FP.Get_typeId(_env),typeInfo)
    }
    self.FP.writeln(_env, _env.GetVM().String_format("function %s:_getTxt( val )\n   local name = val[ 1 ]\n   if name then\n      return string.format( \"%s.%%s\", name )\n   end\n   return string.format( \"illegal val -- %%s\", val )\nend\n", Lns_2DDD(algeFullName, algeFullName)))
    self.FP.writeln(_env, _env.GetVM().String_format("function %s._from( val )\n   return _lune._AlgeFrom( %s, val )\nend\n", Lns_2DDD(algeFullName, algeFullName)))
    {
        __forsortCollection0 := node.FP.Get_algeType(_env).FP.Get_valInfoMap(_env)
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            valInfo := __forsortCollection0.Items[ ___forsortKey0 ]
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.%s = { \"%s\"", Lns_2DDD(algeFullName, valInfo.FP.Get_name(_env), valInfo.FP.Get_name(_env))))
            if valInfo.FP.Get_typeList(_env).Len() > 0{
                self.FP.WriteRaw(_env, ", {")
                for _index, _paramType := range( valInfo.FP.Get_typeList(_env).Items ) {
                    index := _index + 1
                    paramType := _paramType
                    if index > 1{
                        self.FP.WriteRaw(_env, ",")
                    }
                    if Lns_car(Ast_NormalTypeInfo_isAvailableMapping(_env, self.processInfo, &node.FP.Get_algeType(_env).Ast_TypeInfo, NewLnsMap2_[*Ast_TypeInfo,bool]( map[*Ast_TypeInfo]bool{}))).(bool){
                        var funcTxt string
                        var nilable bool
                        var child string
                        funcTxt,nilable,child = self.FP.getMapInfo(_env, paramType)
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("{ func=%s, nilable=%s, child=%s }", Lns_2DDD(funcTxt, nilable, child)))
                    } else { 
                        self.FP.WriteRaw(_env, "{}")
                    }
                }
                self.FP.WriteRaw(_env, "}")
            }
            self.FP.writeln(_env, "}")
            self.FP.writeln(_env, _env.GetVM().String_format("%s._name2Val[\"%s\"] = %s.%s", Lns_2DDD(algeFullName, valInfo.FP.Get_name(_env), algeFullName, valInfo.FP.Get_name(_env))))
        }
    }
}
// 1672: decl @lune.@base.@convLua.ConvFilter.processNewAlgeVal
func (self *convLua_ConvFilter) ProcessNewAlgeVal(_env *LnsEnv, node *Nodes_NewAlgeValNode,_opt LnsAny) {
    var valInfo *Ast_AlgeValInfo
    valInfo = node.FP.Get_valInfo(_env)
    self.FP.WriteRaw(_env, _env.GetVM().String_format("_lune.newAlge( %s.%s", Lns_2DDD(self.FP.getFullName(_env, &node.FP.Get_algeTypeInfo(_env).Ast_TypeInfo), valInfo.FP.Get_name(_env))))
    if valInfo.FP.Get_typeList(_env).Len() > 0{
        self.FP.WriteRaw(_env, ", {")
        for _index, _exp := range( node.FP.Get_paramList(_env).Items ) {
            index := _index + 1
            exp := _exp
            if index > 1{
                self.FP.WriteRaw(_env, ",")
            }
            convLua_filter_9_(_env, exp, self, &node.Nodes_Node)
        }
        self.FP.WriteRaw(_env, "}")
    }
    self.FP.WriteRaw(_env, ")")
}
// 1693: decl @lune.@base.@convLua.ConvFilter.getDestrClass
func (self *convLua_ConvFilter) getDestrClass(_env *LnsEnv, classTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo *Ast_TypeInfo
    typeInfo = classTypeInfo
    for typeInfo != Ast_headTypeInfo {
        var scope *Ast_Scope
        scope = Lns_unwrap( typeInfo.FP.Get_scope(_env)).(*Ast_Scope)
        if Lns_isCondTrue( scope.FP.GetTypeInfoChild(_env, "__free")){
            return typeInfo
        }
        typeInfo = typeInfo.FP.Get_baseTypeInfo(_env)
    }
    return nil
}
// 1706: decl @lune.@base.@convLua.ConvFilter.outputAlter2MapFunc
func (self *convLua_ConvFilter) outputAlter2MapFunc(_env *LnsEnv, stream Lns_oStream,alt2Map *LnsMap2_[*Ast_TypeInfo,*Ast_TypeInfo]) {
    stream.Write(_env, "{")
    for _altType, _assinType := range( alt2Map.Items ) {
        altType := _altType
        assinType := _assinType
        if altType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
            if assinType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                stream.Write(_env, _env.GetVM().String_format("%s = obj.__alt2mapFunc.%s,", Lns_2DDD(altType.FP.Get_rawTxt(_env), assinType.FP.Get_rawTxt(_env))))
            } else { 
                var funcTxt string
                var nilable bool
                var child string
                funcTxt,nilable,child = self.FP.getMapInfo(_env, assinType)
                stream.Write(_env, _env.GetVM().String_format("%s = { func=%s, nilable=%s, child=%s },", Lns_2DDD(altType.FP.Get_rawTxt(_env), funcTxt, nilable, child)))
            }
        }
    }
    stream.Write(_env, "}")
}
// 1731: decl @lune.@base.@convLua.ConvFilter.processProtoClass
func (self *convLua_ConvFilter) ProcessProtoClass(_env *LnsEnv, node *Nodes_ProtoClassNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("local %s = {}", Lns_2DDD(node.FP.Get_name(_env).Txt)))
}
// 1737: decl @lune.@base.@convLua.ConvFilter.processDeclClass
func (self *convLua_ConvFilter) ProcessDeclClass(_env *LnsEnv, node *Nodes_DeclClassNode,_opt LnsAny) {
    var nodeInfo *Nodes_DeclClassNode
    nodeInfo = node
    var classNameToken *Types_Token
    classNameToken = nodeInfo.FP.Get_name(_env)
    var className string
    className = classNameToken.Txt
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = node.FP.Get_expType(_env)
    var classTypeId *Ast_IdInfo
    classTypeId = classTypeInfo.FP.Get_typeId(_env)
    if nodeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
        self.classId2TypeInfo.Set(classTypeId,classTypeInfo)
    }
    self.classId2MemberList.Set(classTypeId,nodeInfo.FP.Get_memberList(_env))
    {
        __exp := node.FP.Get_moduleName(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Types_Token)
            self.FP.WriteRaw(_env, _env.GetVM().String_format("local %s = ", Lns_2DDD(className)))
            if node.FP.Get_lazyLoad(_env) == Nodes_LazyLoad__Off{
                self.FP.WriteRaw(_env, "require")
            } else { 
                self.FP.WriteRaw(_env, "_lune._lazyRequire")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("( %s )", Lns_2DDD(_exp.Txt)))
            if _switch0 := node.FP.Get_accessMode(_env); _switch0 == Ast_AccessMode__Pub || _switch0 == Ast_AccessMode__Pro {
                if self.needModuleObj{
                    self.FP.writeln(_env, "")
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("_moduleObj.%s = %s", Lns_2DDD(className, className)))
                }
            }
            return 
        }
    }
    if Lns_op_not(node.FP.Get_hasPrototype(_env)){
        self.FP.writeln(_env, _env.GetVM().String_format("local %s = {}", Lns_2DDD(className)))
    }
    var ifTxt string
    ifTxt = ""
    if classTypeInfo.FP.Get_interfaceList(_env).Len() > 0{
        for _, _ifType := range( classTypeInfo.FP.Get_interfaceList(_env).Items ) {
            ifType := _ifType
            ifTxt = ifTxt + self.FP.getFullName(_env, ifType) + ","
        }
        ifTxt = _env.GetVM().String_format("ifList = {%s}", Lns_2DDD(ifTxt))
    }
    var baseInfo *Ast_TypeInfo
    baseInfo = classTypeInfo.FP.Get_baseTypeInfo(_env)
    var baseTxt string
    baseTxt = ""
    if baseInfo.FP.Get_typeId(_env) != Ast_rootTypeIdInfo{
        baseTxt = _env.GetVM().String_format("__index = %s", Lns_2DDD(self.FP.getFullName(_env, baseInfo)))
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( len(ifTxt) > 0) ||
        _env.SetStackVal( len(baseTxt) > 0) ).(bool){
        var metaTxt string
        metaTxt = baseTxt
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( len(baseTxt) > 0) &&
            _env.SetStackVal( len(ifTxt) > 0) ).(bool)){
            metaTxt = _env.GetVM().String_format("%s,%s", Lns_2DDD(baseTxt, ifTxt))
        } else if len(ifTxt) > 0{
            metaTxt = ifTxt
        }
        self.FP.writeln(_env, _env.GetVM().String_format("setmetatable( %s, { %s } )", Lns_2DDD(className, metaTxt)))
    }
    if nodeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
        if self.needModuleObj{
            self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.%s = %s", Lns_2DDD(className, className)))
        }
    }
    for _, _declNode := range( node.FP.Get_declStmtList(_env).Items ) {
        declNode := _declNode
        convLua_filter_9_(_env, declNode, self, &node.Nodes_Node)
    }
    var hasConstrFlag bool
    hasConstrFlag = false
    var memberList *LnsList2_[*Nodes_DeclMemberNode]
    memberList = NewLnsList2_[*Nodes_DeclMemberNode]([]*Nodes_DeclMemberNode{})
    var fieldList *LnsList2_[*Nodes_Node]
    fieldList = nodeInfo.FP.Get_fieldList(_env)
    var outerMethodSet *LnsSet2_[string]
    outerMethodSet = nodeInfo.FP.Get_outerMethodSet(_env)
    var methodNameSet *LnsSet2_[string]
    methodNameSet = NewLnsSet2_[string]([]string{})
    if classTypeInfo.FP.Get_kind(_env) != Ast_TypeInfoKind__IF{
        for _, _field := range( fieldList.Items ) {
            field := _field
            var ignoreFlag bool
            ignoreFlag = false
            if field.FP.Get_kind(_env) == Nodes_NodeKind_get_DeclConstr(_env){
                hasConstrFlag = true
                methodNameSet.Add("__init")
            }
            if field.FP.Get_kind(_env) == Nodes_NodeKind_get_DeclDestr(_env){
                methodNameSet.Add("__free")
            }
            {
                _declMemberNode := Nodes_DeclMemberNodeDownCastF(field.FP)
                if !Lns_IsNil( _declMemberNode ) {
                    declMemberNode := _declMemberNode.(*Nodes_DeclMemberNode)
                    if Lns_op_not(declMemberNode.FP.Get_staticFlag(_env)){
                        memberList.Insert(declMemberNode)
                    }
                }
            }
            {
                _methodNode := Nodes_DeclMethodNodeDownCastF(field.FP)
                if !Lns_IsNil( _methodNode ) {
                    methodNode := _methodNode.(*Nodes_DeclMethodNode)
                    var declInfo *Nodes_DeclFuncInfo
                    declInfo = methodNode.FP.Get_declInfo(_env)
                    var methodNameToken *Types_Token
                    methodNameToken = Lns_unwrap( declInfo.FP.Get_name(_env)).(*Types_Token)
                    if outerMethodSet.Has(methodNameToken.Txt){
                        ignoreFlag = true
                    }
                    methodNameSet.Add(methodNameToken.Txt)
                }
            }
            if (Lns_op_not(ignoreFlag)){
                convLua_filter_9_(_env, field, self, &node.Nodes_Node)
            }
        }
    }
    var destTxt string
    destTxt = ""
    {
        __exp := self.FP.getDestrClass(_env, node.FP.Get_expType(_env))
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            destTxt = _env.GetVM().String_format(", __gc = %s.__free", Lns_2DDD(self.FP.getFullName(_env, _exp)))
        }
    }
    self.FP.writeln(_env, _env.GetVM().String_format("function %s.%s( obj )\n  setmetatable( obj, { __index = %s %s } )\nend", Lns_2DDD(className, self.newMethod.SetmetaName, className, destTxt)))
    if Lns_op_not(hasConstrFlag){
        methodNameSet.Add("__init")
        var oldFlag bool
        oldFlag = node.FP.Get_hasOldCtor(_env)
        var superArgTxt string
        superArgTxt = ""
        var thisArgTxt string
        thisArgTxt = ""
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(oldFlag)) &&
            _env.SetStackVal( baseInfo != Ast_headTypeInfo) ).(bool)){
            {
                _superInit := (Lns_unwrap( baseInfo.FP.Get_scope(_env)).(*Ast_Scope)).FP.GetSymbolInfoChild(_env, "__init")
                if !Lns_IsNil( _superInit ) {
                    superInit := _superInit.(*Ast_SymbolInfo)
                    for _index, _ := range( superInit.FP.Get_typeInfo(_env).FP.Get_argTypeInfoList(_env).Items ) {
                        index := _index + 1
                        if len(superArgTxt) > 0{
                            superArgTxt = superArgTxt + ", "
                        }
                        superArgTxt = _env.GetVM().String_format("%s__superarg%d", Lns_2DDD(superArgTxt, index))
                    }
                }
            }
        }
        for _, _member := range( memberList.Items ) {
            member := _member
            if len(thisArgTxt) > 0{
                thisArgTxt = thisArgTxt + ", "
            }
            thisArgTxt = thisArgTxt + member.FP.Get_name(_env).Txt
        }
        var argTxt string
        argTxt = superArgTxt
        if thisArgTxt != ""{
            if len(argTxt) > 0{
                argTxt = argTxt + ","
            }
            argTxt = argTxt + thisArgTxt
        }
        self.FP.writeln(_env, _env.GetVM().String_format("function %s.%s( %s )\n   local obj = {}\n   %s.%s( obj )\n   if obj.__init then\n      obj:__init( %s )\n   end\n   return obj\nend\nfunction %s:__init( %s )\n", Lns_2DDD(className, self.newMethod.NewName, argTxt, className, self.newMethod.SetmetaName, argTxt, className, argTxt)))
        self.FP.pushIndent(_env, nil)
        if baseInfo != Ast_headTypeInfo{
            if Lns_isCondTrue( (Lns_unwrap( baseInfo.FP.Get_scope(_env)).(*Ast_Scope)).FP.GetSymbolInfoChild(_env, "__init")){
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.__init( self", Lns_2DDD(self.FP.getFullName(_env, baseInfo))))
                if len(superArgTxt) > 0{
                    self.FP.writeln(_env, _env.GetVM().String_format(", %s )", Lns_2DDD(superArgTxt)))
                } else { 
                    self.FP.writeln(_env, ")")
                }
            }
        }
        for _, _member := range( memberList.Items ) {
            member := _member
            var memberName string
            memberName = member.FP.Get_name(_env).Txt
            self.FP.writeln(_env, _env.GetVM().String_format("self.%s = %s", Lns_2DDD(memberName, memberName)))
        }
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "end")
    }
    for _, _memberNode := range( nodeInfo.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode
        var memberNameToken *Types_Token
        memberNameToken = memberNode.FP.Get_name(_env)
        var memberName string
        memberName = memberNameToken.Txt
        var getterName string
        getterName = "get_" + memberName
        var autoFlag bool
        autoFlag = Lns_op_not(methodNameSet.Has(getterName))
        var prefix string
        var delimit string
        if memberNode.FP.Get_staticFlag(_env){
            prefix = className
            delimit = "."
        } else { 
            prefix = "self"
            delimit = ":"
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( memberNode.FP.Get_getterMode(_env) != Ast_AccessMode__None) &&
            _env.SetStackVal( autoFlag) ).(bool)){
            self.FP.writeln(_env, _env.GetVM().String_format("function %s%s%s()\n   return %s.%s\nend", Lns_2DDD(className, delimit, getterName, prefix, memberName)))
            methodNameSet.Add(getterName)
        }
        var setterName string
        setterName = "set_" + memberName
        autoFlag = Lns_op_not(methodNameSet.Has(setterName))
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( memberNode.FP.Get_setterMode(_env) != Ast_AccessMode__None) &&
            _env.SetStackVal( autoFlag) ).(bool)){
            self.FP.writeln(_env, _env.GetVM().String_format("function %s%s%s( %s )\n   %s.%s = %s\nend", Lns_2DDD(className, delimit, setterName, memberName, prefix, memberName, memberName)))
            methodNameSet.Add(setterName)
        }
    }
    for _, _advertiseInfo := range( node.FP.Get_advertiseList(_env).Items ) {
        advertiseInfo := _advertiseInfo
        var memberName string
        memberName = advertiseInfo.FP.Get_member(_env).FP.Get_name(_env).Txt
        var memberType *Ast_TypeInfo
        memberType = advertiseInfo.FP.Get_member(_env).FP.Get_expType(_env)
        for _, _mtdName := range( Ast_getAllMethodName(_env, memberType, Ast_MethodKind__Object).FP.Get_list(_env).Items ) {
            mtdName := _mtdName.(string)
            var mbrScope *Ast_Scope
            mbrScope = Lns_unwrap( memberType.FP.Get_scope(_env)).(*Ast_Scope)
            var child *Ast_TypeInfo
            child = Lns_unwrap( mbrScope.FP.GetTypeInfoField(_env, mtdName, true, mbrScope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
            if child.FP.Get_accessMode(_env) != Ast_AccessMode__Pri{
                var childName string
                childName = advertiseInfo.FP.Get_prefix(_env) + child.FP.GetTxt(_env, nil, nil, nil)
                if Lns_op_not(methodNameSet.Has(childName)){
                    self.FP.writeln(_env, _env.GetVM().String_format("function %s:%s( ... )\n   return self.%s:%s( ... )\nend\n", Lns_2DDD(className, childName, memberName, childName)))
                }
            }
        }
    }
    {
        _initBlock := _env.NilAccFin(_env.NilAccPush(nodeInfo.FP.Get_initBlock(_env).FP.Get_func(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_DeclMethodNode).FP.Get_declInfo(_env)})&&
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_DeclFuncInfo).FP.Get_body(_env)}))
        if !Lns_IsNil( _initBlock ) {
            initBlock := _initBlock.(*Nodes_BlockNode)
            if initBlock.FP.Get_stmtList(_env).Len() > 0{
                self.FP.writeln(_env, "do")
                self.FP.pushIndent(_env, nil)
                for _, _initStmt := range( initBlock.FP.Get_stmtList(_env).Items ) {
                    initStmt := _initStmt
                    convLua_filter_9_(_env, initStmt, self, &node.Nodes_Node)
                    self.FP.writeln(_env, "")
                }
                self.FP.popIndent(_env)
                self.FP.writeln(_env, "end")
            }
        }
    }
    if classTypeInfo.FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeMapping, nil){
        var declArgTxt string
        declArgTxt = "val"
        var argTxt string
        argTxt = "{}, val"
        if convLua_isGenericType_12_(_env, classTypeInfo){
            declArgTxt = "val, __alt2mapFunc"
            argTxt = "{ __alt2mapFunc = __alt2mapFunc }, val"
        }
        self.FP.writeln(_env, _env.GetVM().String_format("function %s:_toMap()\n  return self\nend\nfunction %s._fromMap( %s )\n  local obj, mes = %s._fromMapSub( %s )\n  if obj then\n     %s.%s( obj )\n  end\n  return obj, mes\nend\nfunction %s._fromStem( %s )\n  return %s._fromMap( %s )\nend\n", Lns_2DDD(className, className, declArgTxt, className, argTxt, className, self.newMethod.SetmetaName, className, declArgTxt, className, declArgTxt)))
        self.FP.writeln(_env, _env.GetVM().String_format("function %s._fromMapSub( obj, val )", Lns_2DDD(className)))
        if classTypeInfo.FP.Get_baseTypeInfo(_env) != Ast_headTypeInfo{
            self.FP.writeln(_env, _env.GetVM().String_format("   local result, mes = %s._fromMapSub( obj, val )\n   if not result then\n      return nil, mes\n   end\n", Lns_2DDD(self.FP.getFullName(_env, classTypeInfo.FP.Get_baseTypeInfo(_env)))))
        }
        self.FP.writeln(_env, "   local memInfo = {}")
        for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
            memberNode := _memberNode
            var funcTxt string
            var nilable bool
            var child string
            funcTxt,nilable,child = self.FP.getMapInfo(_env, memberNode.FP.Get_expType(_env))
            self.FP.writeln(_env, _env.GetVM().String_format("   table.insert( memInfo, { name = \"%s\", func = %s, nilable = %s, child = %s } )", Lns_2DDD(memberNode.FP.Get_name(_env).Txt, funcTxt, nilable, child)))
        }
        self.FP.writeln(_env, "   local result, mess = _lune._fromMap( obj, val, memInfo )\n   if not result then\n      return nil, mess\n   end\n   return obj\nend")
    }
}
// 2067: decl @lune.@base.@convLua.ConvFilter.processDeclMember
func (self *convLua_ConvFilter) ProcessDeclMember(_env *LnsEnv, node *Nodes_DeclMemberNode,_opt LnsAny) {
}
// 2074: decl @lune.@base.@convLua.ConvFilter.processExpMacroArgExp
func (self *convLua_ConvFilter) ProcessExpMacroArgExp(_env *LnsEnv, node *Nodes_ExpMacroArgExpNode,_opt LnsAny) {
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
}
// 2081: decl @lune.@base.@convLua.ConvFilter.processExpMacroExp
func (self *convLua_ConvFilter) ProcessExpMacroExp(_env *LnsEnv, node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    for _, _stmt := range( node.FP.Get_stmtList(_env).Items ) {
        stmt := _stmt
        convLua_filter_9_(_env, stmt, self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
    }
}
// 2093: decl @lune.@base.@convLua.ConvFilter.outputDeclMacro
func (self *convLua_ConvFilter) OutputDeclMacro(_env *LnsEnv, name string,argNameList *LnsList2_[string],callback convLua_outputMacroStmtBlock_13_) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("local function %s(", Lns_2DDD(name)))
    self.FP.writeln(_env, "__macroArgs )")
    self.FP.pushIndent(_env, nil)
    self.FP.writeln(_env, _env.GetVM().String_format("local _lune = require( \"%s\" )", Lns_2DDD(Option_getRuntimeModule(_env))))
    self.FP.writeln(_env, "local __var = __macroArgs.__var")
    for _, _argName := range( argNameList.Items ) {
        argName := _argName
        self.FP.writeln(_env, _env.GetVM().String_format("local %s = __macroArgs.%s", Lns_2DDD(argName, argName)))
    }
    self.FP.writeln(_env, "local macroVar = {}")
    self.FP.writeln(_env, "macroVar.__names = {}")
    self.FP.writeln(_env, "local function __expStatList( list )\n  local ret = \"\"\n  for index, txt in ipairs( list ) do\n    ret = string.format( \"%s %s \", ret, txt )\n  end\n  return ret\nend\n")
    self.macroDepth = self.macroDepth + 1
    callback(_env)
    self.macroDepth = self.macroDepth - 1
    self.FP.writeln(_env, "")
    self.FP.writeln(_env, "macroVar.__var = __var")
    self.FP.writeln(_env, "return macroVar")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
    self.FP.writeln(_env, _env.GetVM().String_format("return %s", Lns_2DDD(name)))
}
// 2137: decl @lune.@base.@convLua.ConvFilter.processExpMacroStatList
func (self *convLua_ConvFilter) ProcessExpMacroStatList(_env *LnsEnv, node *Nodes_ExpMacroStatListNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "__expStatList(")
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, ")")
}
// 2145: decl @lune.@base.@convLua.ConvFilter.processDeclMacro
func (self *convLua_ConvFilter) ProcessDeclMacro(_env *LnsEnv, node *Nodes_DeclMacroNode,_opt LnsAny) {
    if self.inMacro{
        var macroInfo *Nodes_DeclMacroInfo
        macroInfo = node.FP.Get_declInfo(_env)
        var argNameList *LnsList2_[string]
        argNameList = NewLnsList2_[string]([]string{})
        for _, _arg := range( macroInfo.FP.Get_argList(_env).Items ) {
            arg := _arg
            argNameList.Insert(arg.FP.Get_name(_env).Txt)
        }
        self.FP.OutputDeclMacro(_env, macroInfo.FP.Get_name(_env).Txt, argNameList, convLua_outputMacroStmtBlock_13_(func(_env *LnsEnv) {
            {
                _stmtBlock := macroInfo.FP.Get_stmtBlock(_env)
                if !Lns_IsNil( _stmtBlock ) {
                    stmtBlock := _stmtBlock.(*Nodes_BlockNode)
                    convLua_filter_9_(_env, &stmtBlock.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }))
    }
}
// 2164: decl @lune.@base.@convLua.ConvFilter.processExpMacroStat
func (self *convLua_ConvFilter) ProcessExpMacroStat(_env *LnsEnv, node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    if node.FP.Get_expStrList(_env).Len() == 0{
        self.FP.WriteRaw(_env, "''")
    } else { 
        for _index, _token := range( node.FP.Get_expStrList(_env).Items ) {
            index := _index + 1
            token := _token
            if index != 1{
                self.FP.WriteRaw(_env, "..")
            }
            convLua_filter_9_(_env, token, self, &node.Nodes_Node)
        }
    }
}
// 2182: decl @lune.@base.@convLua.ConvFilter.processExpNew
func (self *convLua_ConvFilter) ProcessExpNew(_env *LnsEnv, node *Nodes_ExpNewNode,_opt LnsAny) {
    convLua_filter_9_(_env, node.FP.Get_symbol(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, _env.GetVM().String_format(".%s(", Lns_2DDD(self.newMethod.NewName)))
    {
        __exp := node.FP.Get_argList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.WriteRaw(_env, ")")
}
// 2193: decl @lune.@base.@convLua.ConvFilter.process__func__symbol
func (self *convLua_ConvFilter) process__func__symbol(_env *LnsEnv, has__func__Symbol bool,parentType *Ast_TypeInfo,funcName string) {
    if has__func__Symbol{
        var nameSpace string
        nameSpace = self.FP.getCanonicalName(_env, parentType, false)
        if funcName == ""{
            funcName = "<anonymous>"
        }
        self.FP.pushIndent(_env, nil)
        self.FP.writeln(_env, _env.GetVM().String_format("local __func__ = '%s.%s'", Lns_2DDD(nameSpace, funcName)))
        self.FP.popIndent(_env)
    }
}
// 2207: decl @lune.@base.@convLua.ConvFilter.processDeclConstr
func (self *convLua_ConvFilter) ProcessDeclConstr(_env *LnsEnv, node *Nodes_DeclConstrNode,_opt LnsAny) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo(_env)
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = Lns_unwrap( declInfo.FP.Get_classTypeInfo(_env)).(*Ast_TypeInfo)
    var className string
    className = self.FP.getFullName(_env, classTypeInfo)
    self.FP.WriteRaw(_env, _env.GetVM().String_format("function %s.%s( ", Lns_2DDD(className, self.newMethod.NewName)))
    var argTxt string
    argTxt = ""
    self.FP.WriteRaw(_env, argTxt)
    var argList *LnsList2_[*Nodes_Node]
    argList = declInfo.FP.Get_argList(_env)
    for _, _arg := range( argList.Items ) {
        arg := _arg
        if len(argTxt) > 0{
            self.FP.WriteRaw(_env, ", ")
            argTxt = argTxt + ", "
        }
        convLua_filter_9_(_env, arg, self, &node.Nodes_Node)
        {
            __exp := Nodes_DeclArgNodeDownCastF(arg.FP)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_DeclArgNode)
                argTxt = argTxt + _exp.FP.Get_name(_env).Txt
            } else {
                var name *Types_Token
                name = Lns_unwrap( node.FP.Get_declInfo(_env).FP.Get_name(_env)).(*Types_Token)
                Util_err(_env, _env.GetVM().String_format("not support ... in macro -- %s", Lns_2DDD(name.Txt)))
            }
        }
    }
    self.FP.writeln(_env, " )")
    self.FP.pushIndent(_env, nil)
    self.FP.writeln(_env, "local obj = {}")
    self.FP.writeln(_env, _env.GetVM().String_format("%s.%s( obj )", Lns_2DDD(className, self.newMethod.SetmetaName)))
    self.FP.writeln(_env, _env.GetVM().String_format("if obj.__init then obj:__init( %s ); end", Lns_2DDD(argTxt)))
    self.FP.writeln(_env, "return obj")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
    self.FP.WriteRaw(_env, _env.GetVM().String_format("function %s:__init(%s) ", Lns_2DDD(className, argTxt)))
    {
        __exp := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.process__func__symbol(_env, declInfo.FP.Get_has__func__Symbol(_env), node.FP.Get_expType(_env).FP.Get_parentInfo(_env), "__init")
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, "end")
}
// 2262: decl @lune.@base.@convLua.ConvFilter.processDeclDestr
func (self *convLua_ConvFilter) ProcessDeclDestr(_env *LnsEnv, node *Nodes_DeclDestrNode,_opt LnsAny) {
    self.FP.writeln(_env, _env.GetVM().String_format("function %s.__free( self )", Lns_2DDD(_env.NilAccFin(_env.NilAccPush(node.FP.Get_declInfo(_env).FP.Get_classTypeInfo(_env)) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetTxt(_env, nil, nil, nil)})/* 2266:9 */))))
    self.FP.process__func__symbol(_env, node.FP.Get_declInfo(_env).FP.Get_has__func__Symbol(_env), node.FP.Get_expType(_env).FP.Get_parentInfo(_env), "__free")
    convLua_filter_9_(_env, &Lns_unwrap( node.FP.Get_declInfo(_env).FP.Get_body(_env)).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = node.FP.Get_expType(_env).FP.Get_parentInfo(_env)
    {
        __exp := self.FP.getDestrClass(_env, classTypeInfo.FP.Get_baseTypeInfo(_env))
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            self.FP.writeln(_env, _env.GetVM().String_format("%s.__free( self )", Lns_2DDD(self.FP.getFullName(_env, _exp))))
        }
    }
    self.FP.writeln(_env, "end")
}
// 2282: decl @lune.@base.@convLua.ConvFilter.processExpCallSuperCtor
func (self *convLua_ConvFilter) ProcessExpCallSuperCtor(_env *LnsEnv, node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType(_env)
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.%s( self", Lns_2DDD(self.FP.getFullName(_env, typeInfo), node.FP.Get_methodType(_env).FP.Get_rawTxt(_env))))
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.WriteRaw(_env, ",")
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, ")")
}
// 2297: decl @lune.@base.@convLua.ConvFilter.processExpCallSuper
func (self *convLua_ConvFilter) ProcessExpCallSuper(_env *LnsEnv, node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType(_env)
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.%s( self", Lns_2DDD(self.FP.getFullName(_env, typeInfo), node.FP.Get_methodType(_env).FP.Get_rawTxt(_env))))
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.WriteRaw(_env, ",")
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.WriteRaw(_env, ")")
}
// 2313: decl @lune.@base.@convLua.ConvFilter.processDeclMethod
func (self *convLua_ConvFilter) ProcessDeclMethod(_env *LnsEnv, node *Nodes_DeclMethodNode,_opt LnsAny) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo(_env)
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = Lns_unwrap( declInfo.FP.Get_classTypeInfo(_env)).(*Ast_TypeInfo)
    var delimit string
    delimit = ":"
    if declInfo.FP.Get_staticFlag(_env){
        delimit = "."
    }
    var methodNodeToken *Types_Token
    methodNodeToken = Lns_unwrap( declInfo.FP.Get_name(_env)).(*Types_Token)
    var methodName string
    methodName = methodNodeToken.Txt
    self.FP.WriteRaw(_env, _env.GetVM().String_format("function %s%s%s( ", Lns_2DDD(classTypeInfo.FP.Get_rawTxt(_env), delimit, methodName)))
    var argList *LnsList2_[*Nodes_Node]
    argList = declInfo.FP.Get_argList(_env)
    for _index, _arg := range( argList.Items ) {
        index := _index + 1
        arg := _arg
        if index > 1{
            self.FP.WriteRaw(_env, ", ")
        }
        convLua_filter_9_(_env, arg, self, &node.Nodes_Node)
    }
    self.FP.writeln(_env, " )")
    {
        __exp := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.process__func__symbol(_env, declInfo.FP.Get_has__func__Symbol(_env), node.FP.Get_expType(_env).FP.Get_parentInfo(_env), methodName)
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, "end")
}
// 2346: decl @lune.@base.@convLua.ConvFilter.processUnwrapSet
func (self *convLua_ConvFilter) ProcessUnwrapSet(_env *LnsEnv, node *Nodes_UnwrapSetNode,_opt LnsAny) {
    var dstExpList *Nodes_ExpListNode
    dstExpList = node.FP.Get_dstExpList(_env)
    convLua_filter_9_(_env, &dstExpList.Nodes_Node, self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, " = ")
    convLua_filter_9_(_env, &node.FP.Get_srcExpList(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
    self.FP.WriteRaw(_env, "if ")
    for _index, _expNode := range( dstExpList.FP.Get_expList(_env).Items ) {
        index := _index + 1
        expNode := _expNode
        if index > 1{
            self.FP.WriteRaw(_env, " or ")
        }
        self.FP.WriteRaw(_env, "nil == ")
        convLua_filter_9_(_env, expNode, self, &node.Nodes_Node)
    }
    self.FP.writeln(_env, " then")
    self.FP.pushIndent(_env, nil)
    for _index, _expNode := range( dstExpList.FP.Get_expList(_env).Items ) {
        index := _index + 1
        expNode := _expNode
        self.FP.WriteRaw(_env, _env.GetVM().String_format("local _exp%d = ", Lns_2DDD(index)))
        convLua_filter_9_(_env, expNode, self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
    }
    if Lns_isCondTrue( node.FP.Get_unwrapBlock(_env)){
        convLua_filter_9_(_env, &Lns_unwrap( node.FP.Get_unwrapBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
    }
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}
// 2379: decl @lune.@base.@convLua.ConvFilter.processExpListSub
func (self *convLua_ConvFilter) processExpListSub(_env *LnsEnv, parent *Nodes_Node,expList *LnsList2_[*Nodes_Node],mRetExp LnsAny) {
    var mRetIndex LnsAny
    mRetIndex = _env.NilAccFin(_env.NilAccPush(mRetExp) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
    for _index, _exp := range( expList.Items ) {
        index := _index + 1
        exp := _exp
        if exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
            break
        }
        {
            _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
            if !Lns_IsNil( _castNode ) {
                castNode := _castNode.(*Nodes_ExpCastNode)
                if castNode.FP.Get_castKind(_env) == Nodes_CastKind__Implicit{
                    if castNode.FP.Get_exp(_env).FP.Get_kind(_env) == Nodes_NodeKind_get_ExpAccessMRet(_env){
                        break
                    }
                }
            }
        }
        if index > 1{
            self.FP.WriteRaw(_env, ", ")
        }
        convLua_filter_9_(_env, exp, self, parent)
        if index == mRetIndex{
            break
        }
    }
}
// 2406: decl @lune.@base.@convLua.ConvFilter.processExpMRet
func (self *convLua_ConvFilter) ProcessExpMRet(_env *LnsEnv, node *Nodes_ExpMRetNode,_opt LnsAny) {
    convLua_filter_9_(_env, node.FP.Get_mRet(_env), self, &node.Nodes_Node)
}
// 2412: decl @lune.@base.@convLua.ConvFilter.processIfUnwrap
func (self *convLua_ConvFilter) ProcessIfUnwrap(_env *LnsEnv, node *Nodes_IfUnwrapNode,_opt LnsAny) {
    {
        __exp := node.FP.Get_condRetInfo(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_CondRetInfo)
            self.FP.outputCondRetInfo(_env, _exp)
        }
    }
    self.FP.writeln(_env, "do")
    self.FP.pushIndent(_env, nil)
    self.FP.WriteRaw(_env, "local ")
    for _index, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        index := _index + 1
        varSym := _varSym
        self.FP.WriteRaw(_env, convLua_getSymbolTxt_5_(_env, varSym))
        if index != node.FP.Get_varSymList(_env).Len(){
            self.FP.WriteRaw(_env, ", ")
        }
    }
    self.FP.WriteRaw(_env, " = ")
    self.FP.processExpListSub(_env, &node.Nodes_Node, node.FP.Get_expList(_env).FP.Get_expList(_env), node.FP.Get_expList(_env).FP.Get_mRetExp(_env))
    self.FP.writeln(_env, "")
    self.FP.WriteRaw(_env, "if ")
    var hasSym bool
    hasSym = false
    for _, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        varSym := _varSym
        if varSym.FP.Get_name(_env) != "_"{
            if hasSym{
                self.FP.WriteRaw(_env, " and  ")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s ~= nil", Lns_2DDD(convLua_getSymbolTxt_5_(_env, varSym))))
            hasSym = true
        }
    }
    self.FP.WriteRaw(_env, " then")
    convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    {
        __exp := node.FP.Get_nilBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.WriteRaw(_env, "else")
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, "end")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}
// 2459: decl @lune.@base.@convLua.ConvFilter.processWhen
func (self *convLua_ConvFilter) ProcessWhen(_env *LnsEnv, node *Nodes_WhenNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "if ")
    for _index, _symPair := range( node.FP.Get_symPairList(_env).Items ) {
        index := _index + 1
        symPair := _symPair
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s ~= nil", Lns_2DDD(symPair.FP.Get_src(_env).FP.Get_name(_env))))
        if index != node.FP.Get_symPairList(_env).Len(){
            self.FP.WriteRaw(_env, " and ")
        }
    }
    self.FP.WriteRaw(_env, " then")
    convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    {
        __exp := node.FP.Get_elseBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.WriteRaw(_env, "else")
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, "end")
}
// 2481: decl @lune.@base.@convLua.ConvFilter.processExpExpandTuple
func (self *convLua_ConvFilter) ProcessExpExpandTuple(_env *LnsEnv, node *Nodes_ExpExpandTupleNode,_opt LnsAny) {
    self.FP.Write(_env, "table.unpack(")
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, ")")
}
// 2490: decl @lune.@base.@convLua.ConvFilter.processDeclVar
func (self *convLua_ConvFilter) ProcessDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,_opt LnsAny) {
    {
        _condRetInfo := node.FP.Get_condRetInfo(_env)
        if !Lns_IsNil( _condRetInfo ) {
            condRetInfo := _condRetInfo.(*Nodes_CondRetInfo)
            self.FP.outputCondRetInfo(_env, condRetInfo)
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_mode(_env) != Nodes_DeclVarMode__Unwrap) &&
        _env.SetStackVal( node.FP.Get_accessMode(_env) != Ast_AccessMode__Global) ).(bool)){
        self.FP.WriteRaw(_env, "local ")
    }
    var varList *LnsList2_[*Ast_SymbolInfo]
    varList = node.FP.Get_symbolInfoList(_env)
    var varNameList *LnsList2_[string]
    varNameList = NewLnsList2_[string]([]string{})
    var convLua_outputPub func(_env *LnsEnv)
    convLua_outputPub = func(_env *LnsEnv) {
        if _switch0 := node.FP.Get_accessMode(_env); _switch0 == Ast_AccessMode__Pub || _switch0 == Ast_AccessMode__Global {
            self.FP.writeln(_env, "")
            for _, _varName := range( varNameList.Items ) {
                varName := _varName
                var name string
                name = varName
                if self.needModuleObj{
                    self.FP.writeln(_env, _env.GetVM().String_format("_moduleObj.%s = %s", Lns_2DDD(name, name)))
                }
            }
        }
    }
    for _index, __var := range( varList.Items ) {
        index := _index + 1
        _var := __var
        if index > 1{
            self.FP.WriteRaw(_env, ", ")
        }
        var name string
        name = convLua_getSymbolTxt_5_(_env, _var)
        self.FP.WriteRaw(_env, name)
        varNameList.Insert(name)
    }
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.WriteRaw(_env, " = ")
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        } else {
            self.FP.writeln(_env, "")
        }
    }
    {
        __exp := node.FP.Get_unwrapBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.writeln(_env, "")
            self.FP.WriteRaw(_env, "if ")
            var hasVar bool
            hasVar = false
            for _, _varName := range( varNameList.Items ) {
                varName := _varName
                if varName != "_"{
                    if hasVar{
                        self.FP.WriteRaw(_env, " or ")
                    }
                    hasVar = true
                    self.FP.WriteRaw(_env, " nil == " + varName)
                }
            }
            self.FP.writeln(_env, " then")
            self.FP.pushIndent(_env, nil)
            for _, _varName := range( varNameList.Items ) {
                varName := _varName
                self.FP.writeln(_env, _env.GetVM().String_format("local _%s = %s", Lns_2DDD(varName, varName)))
            }
            self.FP.popIndent(_env)
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
            {
                _thenBlock := node.FP.Get_thenBlock(_env)
                if !Lns_IsNil( _thenBlock ) {
                    thenBlock := _thenBlock.(*Nodes_BlockNode)
                    self.FP.writeln(_env, "else")
                    self.FP.pushIndent(_env, nil)
                    convLua_outputPub(_env)
                    convLua_filter_9_(_env, &thenBlock.Nodes_Node, self, &node.Nodes_Node)
                    self.FP.popIndent(_env)
                } else {
                    if _switch0 := node.FP.Get_accessMode(_env); _switch0 == Ast_AccessMode__Pub || _switch0 == Ast_AccessMode__Global {
                        self.FP.writeln(_env, "else")
                        self.FP.pushIndent(_env, nil)
                        convLua_outputPub(_env)
                        self.FP.popIndent(_env)
                    }
                }
            }
            self.FP.writeln(_env, "end")
        } else {
            convLua_outputPub(_env)
        }
    }
    if _switch1 := node.FP.Get_accessMode(_env); _switch1 == Ast_AccessMode__Pub || _switch1 == Ast_AccessMode__Global {
        for _index, _varName := range( varNameList.Items ) {
            index := _index + 1
            varName := _varName
            var name string
            name = varName
            self.pubVarName2InfoMap.Set(name,NewconvLua_PubVerInfo(_env, node.FP.Get_staticFlag(_env), node.FP.Get_accessMode(_env), node.FP.Get_symbolInfoList(_env).GetAt(index).FP.Get_mutable(_env), node.FP.Get_typeInfoList(_env).GetAt(index)))
        }
    }
    if self.macroDepth > 0{
        self.FP.writeln(_env, "")
        for _, _symbolInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
            symbolInfo := _symbolInfo
            var varName string
            varName = convLua_getSymbolTxt_5_(_env, symbolInfo)
            self.FP.writeln(_env, _env.GetVM().String_format("table.insert( macroVar.__names, '%s' )", Lns_2DDD(varName)))
            self.FP.writeln(_env, _env.GetVM().String_format("macroVar.%s = %s", Lns_2DDD(varName, varName)))
            self.macroVarSymSet.Add(symbolInfo)
        }
    }
    self.FP.writeln(_env, "")
}
// 2607: decl @lune.@base.@convLua.ConvFilter.processDeclArg
func (self *convLua_ConvFilter) ProcessDeclArg(_env *LnsEnv, node *Nodes_DeclArgNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s", Lns_2DDD(node.FP.Get_name(_env).Txt)))
}
// 2615: decl @lune.@base.@convLua.ConvFilter.processDeclArgDDD
func (self *convLua_ConvFilter) ProcessDeclArgDDD(_env *LnsEnv, node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "...")
}
// 2621: decl @lune.@base.@convLua.ConvFilter.processDeclFunc
func (self *convLua_ConvFilter) ProcessDeclFunc(_env *LnsEnv, node *Nodes_DeclFuncNode,_opt LnsAny) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo(_env)
    var nameToken LnsAny
    nameToken = declInfo.FP.Get_name(_env)
    var name string
    name = ""
    {
        __exp := nameToken
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Types_Token)
            name = _exp.Txt
        }
    }
    var letTxt string
    letTxt = ""
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( declInfo.FP.Get_accessMode(_env) != Ast_AccessMode__Global) &&
        _env.SetStackVal( len(name) != 0) ).(bool)){
        letTxt = "local "
    }
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%sfunction %s( ", Lns_2DDD(letTxt, name)))
    var argList *LnsList2_[*Nodes_Node]
    argList = declInfo.FP.Get_argList(_env)
    self.FP.processExpListSub(_env, &node.Nodes_Node, argList, nil)
    self.FP.writeln(_env, " )")
    {
        __exp := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.process__func__symbol(_env, declInfo.FP.Get_has__func__Symbol(_env), node.FP.Get_expType(_env).FP.Get_parentInfo(_env), name)
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.WriteRaw(_env, "end")
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType(_env)
    if _switch0 := expType.FP.Get_accessMode(_env); _switch0 == Ast_AccessMode__Pub || _switch0 == Ast_AccessMode__Global {
        if self.needModuleObj{
            self.FP.writeln(_env, "")
            self.FP.WriteRaw(_env, _env.GetVM().String_format("_moduleObj.%s = %s", Lns_2DDD(name, name)))
        }
        self.pubFuncName2InfoMap.Set(name,NewconvLua_PubFuncInfo(_env, declInfo.FP.Get_accessMode(_env), node.FP.Get_expType(_env)))
    }
    self.FP.writeln(_env, "")
}
// 2664: decl @lune.@base.@convLua.ConvFilter.processRefType
func (self *convLua_ConvFilter) ProcessRefType(_env *LnsEnv, node *Nodes_RefTypeNode,_opt LnsAny) {
    if _switch0 := node.FP.Get_mutMode(_env); _switch0 == Ast_MutMode__IMut {
        self.FP.WriteRaw(_env, "&")
    } else if _switch0 == Ast_MutMode__AllMut {
        self.FP.WriteRaw(_env, "+")
    }
    convLua_filter_9_(_env, node.FP.Get_typeNode(_env), self, &node.Nodes_Node)
    if node.FP.Get_array(_env) == "array"{
        self.FP.WriteRaw(_env, "[@]")
    } else if node.FP.Get_array(_env) == "list"{
        self.FP.WriteRaw(_env, "[]")
    }
}
// 2684: decl @lune.@base.@convLua.ConvFilter.processIf
func (self *convLua_ConvFilter) ProcessIf(_env *LnsEnv, node *Nodes_IfNode,_opt LnsAny) {
    var valList *LnsList2_[*Nodes_IfStmtInfo]
    valList = node.FP.Get_stmtList(_env)
    for _index, _val := range( valList.Items ) {
        index := _index + 1
        val := _val
        {
            __exp := val.FP.Get_condRetInfo(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_CondRetInfo)
                self.FP.outputCondRetInfo(_env, _exp)
            }
        }
        if index == 1{
            self.FP.WriteRaw(_env, "if ")
            convLua_filter_9_(_env, val.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else if val.FP.Get_kind(_env) == Nodes_IfKind__ElseIf{
            self.FP.WriteRaw(_env, "elseif ")
            convLua_filter_9_(_env, val.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else { 
            self.FP.writeln(_env, "else")
        }
        self.FP.WriteRaw(_env, " ")
        convLua_filter_9_(_env, &val.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    }
    self.FP.writeln(_env, "end")
}
// 2709: decl @lune.@base.@convLua.ConvFilter.processSwitch
func (self *convLua_ConvFilter) ProcessSwitch(_env *LnsEnv, node *Nodes_SwitchNode,_opt LnsAny) {
    self.FP.writeln(_env, "do")
    self.FP.pushIndent(_env, nil)
    self.FP.WriteRaw(_env, "local _switchExp = ")
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
    if node.FP.Get_caseList(_env).Len() > 0{
        for _index, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
            index := _index + 1
            caseInfo := _caseInfo
            if index == 1{
                self.FP.WriteRaw(_env, "if ")
            } else { 
                self.FP.WriteRaw(_env, "elseif ")
            }
            var expList *Nodes_ExpListNode
            expList = caseInfo.FP.Get_expList(_env)
            for _listIndex, _expNode := range( expList.FP.Get_expList(_env).Items ) {
                listIndex := _listIndex + 1
                expNode := _expNode
                if listIndex != 1{
                    self.FP.WriteRaw(_env, " or ")
                }
                self.FP.WriteRaw(_env, "_switchExp == ")
                convLua_filter_9_(_env, expNode, self, &node.Nodes_Node)
            }
            self.FP.WriteRaw(_env, " then")
            convLua_filter_9_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        }
        {
            __exp := node.FP.Get_default(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_BlockNode)
                self.FP.writeln(_env, "else ")
                self.FP.pushIndent(_env, nil)
                convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
                self.FP.popIndent(_env)
            }
        }
        self.FP.writeln(_env, "end")
    }
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}
// 2752: decl @lune.@base.@convLua.ConvFilter.processMatch
func (self *convLua_ConvFilter) ProcessMatch(_env *LnsEnv, node *Nodes_MatchNode,_opt LnsAny) {
    self.FP.writeln(_env, "do")
    self.FP.pushIndent(_env, nil)
    self.FP.WriteRaw(_env, "local _matchExp = ")
    convLua_filter_9_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
    if node.FP.Get_caseList(_env).Len() > 0{
        var fullName string
        fullName = self.FP.getFullName(_env, &node.FP.Get_algeTypeInfo(_env).Ast_TypeInfo)
        for _index, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
            index := _index + 1
            caseInfo := _caseInfo
            if index == 1{
                self.FP.WriteRaw(_env, "if ")
            } else { 
                self.FP.WriteRaw(_env, "elseif ")
            }
            self.FP.writeln(_env, _env.GetVM().String_format("_matchExp[1] == %s.%s[1] then", Lns_2DDD(fullName, caseInfo.FP.Get_valInfo(_env).FP.Get_name(_env))))
            for _paramNum, _paramSym := range( caseInfo.FP.Get_valParamNameList(_env).Items ) {
                paramNum := _paramNum + 1
                paramSym := _paramSym
                self.FP.writeln(_env, _env.GetVM().String_format("   local %s = _matchExp[2][%d]", Lns_2DDD(paramSym.FP.Get_name(_env), paramNum)))
            }
            convLua_filter_9_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        }
        {
            __exp := node.FP.Get_defaultBlock(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_Node)
                self.FP.writeln(_env, "else ")
                self.FP.pushIndent(_env, nil)
                convLua_filter_9_(_env, _exp, self, &node.Nodes_Node)
                self.FP.popIndent(_env)
            }
        }
        self.FP.writeln(_env, "end")
    }
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}
// 2791: decl @lune.@base.@convLua.ConvFilter.processWhile
func (self *convLua_ConvFilter) ProcessWhile(_env *LnsEnv, node *Nodes_WhileNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "while ")
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, " ")
    convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
}
// 2802: decl @lune.@base.@convLua.ConvFilter.processRepeat
func (self *convLua_ConvFilter) ProcessRepeat(_env *LnsEnv, node *Nodes_RepeatNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "repeat ")
    convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, "until ")
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
}
// 2812: decl @lune.@base.@convLua.ConvFilter.processFor
func (self *convLua_ConvFilter) ProcessFor(_env *LnsEnv, node *Nodes_ForNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("for %s = ", Lns_2DDD(convLua_getSymbolTxt_5_(_env, node.FP.Get_val(_env)))))
    convLua_filter_9_(_env, node.FP.Get_init(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, ", ")
    convLua_filter_9_(_env, node.FP.Get_to(_env), self, &node.Nodes_Node)
    {
        __exp := node.FP.Get_delta(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            self.FP.WriteRaw(_env, ", ")
            convLua_filter_9_(_env, _exp, self, &node.Nodes_Node)
        }
    }
    self.FP.WriteRaw(_env, " ")
    convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
}
// 2828: decl @lune.@base.@convLua.ConvFilter.processApply
func (self *convLua_ConvFilter) ProcessApply(_env *LnsEnv, node *Nodes_ApplyNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "for ")
    var varList *LnsList2_[*Ast_SymbolInfo]
    varList = node.FP.Get_varList(_env)
    for _index, __var := range( varList.Items ) {
        index := _index + 1
        _var := __var
        if index > 1{
            self.FP.WriteRaw(_env, ", ")
        }
        self.FP.WriteRaw(_env, convLua_getSymbolTxt_5_(_env, _var))
    }
    self.FP.WriteRaw(_env, " in ")
    convLua_filter_9_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, " ")
    convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
}
// 2846: decl @lune.@base.@convLua.ConvFilter.processForeach
func (self *convLua_ConvFilter) ProcessForeach(_env *LnsEnv, node *Nodes_ForeachNode,_opt LnsAny) {
    var keySym LnsAny
    var valSym LnsAny
    if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Set{
        keySym = node.FP.Get_val(_env)
        valSym = node.FP.Get_key(_env)
    } else { 
        keySym = node.FP.Get_key(_env)
        valSym = node.FP.Get_val(_env)
    }
    self.FP.WriteRaw(_env, "for ")
    if keySym != nil{
        keySym_22 := keySym.(*Ast_SymbolInfo)
        self.FP.WriteRaw(_env, convLua_getSymbolTxt_5_(_env, keySym_22))
    } else {
        self.FP.WriteRaw(_env, "__index")
    }
    self.FP.WriteRaw(_env, ", ")
    if valSym != nil{
        valSym_25 := valSym.(*Ast_SymbolInfo)
        self.FP.WriteRaw(_env, convLua_getSymbolTxt_5_(_env, valSym_25))
    } else {
        self.FP.WriteRaw(_env, "__val")
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.useIpairs) &&
        _env.SetStackVal( node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__List) ).(bool)){
        if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env).GetAt(1).FP.Get_nilable(_env){
            self.FP.WriteRaw(_env, " in pairs( ")
        } else { 
            self.FP.WriteRaw(_env, " in ipairs( ")
        }
    } else { 
        self.FP.WriteRaw(_env, " in pairs( ")
    }
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, " ) ")
    convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
}
// 2896: decl @lune.@base.@convLua.ConvFilter.processForsort
func (self *convLua_ConvFilter) ProcessForsort(_env *LnsEnv, node *Nodes_ForsortNode,_opt LnsAny) {
    var keySym LnsAny
    var valSym LnsAny
    if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Set{
        keySym = node.FP.Get_val(_env)
        valSym = node.FP.Get_key(_env)
    } else { 
        keySym = node.FP.Get_key(_env)
        valSym = node.FP.Get_val(_env)
    }
    self.FP.writeln(_env, "do")
    self.FP.pushIndent(_env, nil)
    self.FP.writeln(_env, "local __sorted = {}")
    self.FP.WriteRaw(_env, "local __map = ")
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
    self.FP.writeln(_env, "for __key in pairs( __map ) do")
    self.FP.pushIndent(_env, nil)
    self.FP.writeln(_env, "table.insert( __sorted, __key )")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
    self.FP.writeln(_env, "table.sort( __sorted )")
    self.FP.WriteRaw(_env, "for __index, ")
    var key string
    key = "__key"
    if keySym != nil{
        keySym_39 := keySym.(*Ast_SymbolInfo)
        key = convLua_getSymbolTxt_5_(_env, keySym_39)
    }
    self.FP.WriteRaw(_env, key)
    self.FP.writeln(_env, " in ipairs( __sorted ) do")
    self.FP.pushIndent(_env, nil)
    if valSym != nil{
        valSym_41 := valSym.(*Ast_SymbolInfo)
        self.FP.writeln(_env, _env.GetVM().String_format("local %s = __map[ %s ]", Lns_2DDD(convLua_getSymbolTxt_5_(_env, valSym_41), key)))
    }
    convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}
// 2948: decl @lune.@base.@convLua.ConvFilter.processExpUnwrap
func (self *convLua_ConvFilter) ProcessExpUnwrap(_env *LnsEnv, node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    {
        __exp := node.FP.Get_default(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            self.FP.WriteRaw(_env, "_lune.unwrapDefault( ")
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ", ")
            convLua_filter_9_(_env, _exp, self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        } else {
            self.FP.WriteRaw(_env, "_lune.unwrap( ")
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        }
    }
}
// 2966: decl @lune.@base.@convLua.ConvFilter.processExpCall
func (self *convLua_ConvFilter) ProcessExpCall(_env *LnsEnv, node *Nodes_ExpCallNode,_opt LnsAny) {
    var wroteFuncFlag bool
    wroteFuncFlag = false
    var setArgFlag bool
    setArgFlag = false
    var convLua_fieldCall func(_env *LnsEnv) bool
    convLua_fieldCall = func(_env *LnsEnv) bool {
        var fieldNode *Nodes_RefFieldNode
        
        {
            _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func(_env).FP)
            if _fieldNode == nil{
                return true
            } else {
                fieldNode = _fieldNode.(*Nodes_RefFieldNode)
            }
        }
        if _switch0 := node.FP.Get_func(_env).FP.Get_expType(_env); _switch0 == self.builtinFunc.G__lns_runtime_log || _switch0 == self.builtinFunc.G__lns_runtime_enableLog || _switch0 == self.builtinFunc.G__lns_runtime_enableDebugLog || _switch0 == self.builtinFunc.G__lns_runtime_dumpLog || _switch0 == self.builtinFunc.G__processor_end {
            return false
        } else if _switch0 == self.builtinFunc.G__lns_runtime_time {
            self.FP.WriteRaw(_env, "math.floor( os.time() * 1000 )")
            return false
        } else if _switch0 == self.builtinFunc.G__lns_sync_createProcesser {
            self.FP.WriteRaw(_env, "{}")
            return false
        } else if _switch0 == self.builtinFunc.G__lns_sync_createFlag {
            self.FP.WriteRaw(_env, "nil")
            return false
        } else if _switch0 == self.builtinFunc.G__lns_sync__createPipe {
            self.FP.WriteRaw(_env, "nil")
            return false
        }
        var prefixNode *Nodes_Node
        prefixNode = fieldNode.FP.Get_prefix(_env)
        var fieldTxt string
        if _switch1 := node.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_nonnilableType(_env); _switch1 == self.builtinFunc.Str_replace {
            wroteFuncFlag = true
            setArgFlag = true
            self.FP.WriteRaw(_env, "_lune.replace( ")
            convLua_filter_9_(_env, prefixNode, self, &fieldNode.Nodes_Node)
            return true
        } else if _switch1 == self.builtinFunc.String__join || _switch1 == self.builtinFunc.String___join {
            wroteFuncFlag = true
            self.FP.WriteRaw(_env, "table.concat( ")
            return true
        } else {
            fieldTxt = fieldNode.FP.Get_field(_env).Txt
        }
        var convLua_processSet func(_env *LnsEnv) bool
        convLua_processSet = func(_env *LnsEnv) bool {
            setArgFlag = true
            wroteFuncFlag = true
            if _switch1 := fieldTxt; _switch1 == "add" || _switch1 == "del" {
                convLua_filter_9_(_env, prefixNode, self, &fieldNode.Nodes_Node)
                self.FP.WriteRaw(_env, "[")
                {
                    _argList := node.FP.Get_argList(_env)
                    if !Lns_IsNil( _argList ) {
                        argList := _argList.(*Nodes_ExpListNode)
                        convLua_filter_9_(_env, &argList.Nodes_Node, self, &fieldNode.Nodes_Node)
                    }
                }
                self.FP.WriteRaw(_env, "]")
                if _switch0 := fieldTxt; _switch0 == "add" {
                    self.FP.WriteRaw(_env, "= true")
                } else if _switch0 == "del" {
                    self.FP.WriteRaw(_env, "= nil")
                }
                return false
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("_lune._Set_%s(", Lns_2DDD(fieldTxt)))
            convLua_filter_9_(_env, prefixNode, self, &fieldNode.Nodes_Node)
            return true
        }
        var prefixType *Ast_TypeInfo
        prefixType = prefixNode.FP.Get_expType(_env)
        var convLua_processEnumAlge func(_env *LnsEnv)
        convLua_processEnumAlge = func(_env *LnsEnv) {
            wroteFuncFlag = true
            var fieldExpType *Ast_TypeInfo
            fieldExpType = fieldNode.FP.Get_expType(_env)
            var canonicalName string
            canonicalName = self.FP.getFullName(_env, prefixType)
            var methodName string
            methodName = fieldTxt
            var delimit string
            delimit = ":"
            if methodName == "get__txt"{
                methodName = "_getTxt"
            }
            if fieldExpType.FP.Get_kind(_env) == Ast_TypeInfoKind__Func{
                delimit = "."
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s%s%s( ", Lns_2DDD(canonicalName, delimit, methodName)))
            if fieldExpType.FP.Get_staticFlag(_env){
                setArgFlag = false
            } else { 
                convLua_filter_9_(_env, prefixNode, self, &fieldNode.Nodes_Node)
                setArgFlag = true
            }
        }
        if node.FP.Get_nilAccess(_env){
            wroteFuncFlag = true
            setArgFlag = true
            if _switch2 := prefixType.FP.Get_kind(_env); _switch2 == Ast_TypeInfoKind__List || _switch2 == Ast_TypeInfoKind__Array {
                self.FP.WriteRaw(_env, _env.GetVM().String_format("_lune.nilacc( table.%s, nil, 'list', ", Lns_2DDD(fieldTxt)))
                convLua_filter_9_(_env, prefixNode, self, &fieldNode.Nodes_Node)
            } else {
                self.FP.WriteRaw(_env, "_lune.nilacc( ")
                convLua_filter_9_(_env, prefixNode, self, &fieldNode.Nodes_Node)
                self.FP.WriteRaw(_env, _env.GetVM().String_format(", '%s', 'callmtd' ", Lns_2DDD(fieldTxt)))
            }
        } else { 
            if _switch3 := prefixType.FP.Get_kind(_env); _switch3 == Ast_TypeInfoKind__List || _switch3 == Ast_TypeInfoKind__Array || _switch3 == Ast_TypeInfoKind__Tuple {
                setArgFlag = true
                wroteFuncFlag = true
                self.FP.WriteRaw(_env, _env.GetVM().String_format("table.%s( ", Lns_2DDD(fieldTxt)))
                convLua_filter_9_(_env, prefixNode, self, &fieldNode.Nodes_Node)
            } else if _switch3 == Ast_TypeInfoKind__Set {
                if Lns_op_not(convLua_processSet(_env)){
                    return false
                }
            } else if _switch3 == Ast_TypeInfoKind__Enum || _switch3 == Ast_TypeInfoKind__Alge {
                convLua_processEnumAlge(_env)
            } else if _switch3 == Ast_TypeInfoKind__Box {
                convLua_filter_9_(_env, prefixNode, self, &fieldNode.Nodes_Node)
                self.FP.WriteRaw(_env, "[1]")
                return false
            } else if _switch3 == Ast_TypeInfoKind__Class {
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( prefixType.FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeMapping, nil)) &&
                    _env.SetStackVal( convLua_isGenericType_12_(_env, prefixType)) &&
                    _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( fieldTxt == "_fromMap") ||
                        _env.SetStackVal( fieldTxt == "_fromStem") ).(bool))) ).(bool)){
                    wroteFuncFlag = true
                    setArgFlag = true
                    convLua_filter_9_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
                    self.FP.WriteRaw(_env, "( ")
                    {
                        _argList := node.FP.Get_argList(_env)
                        if !Lns_IsNil( _argList ) {
                            argList := _argList.(*Nodes_ExpListNode)
                            convLua_filter_9_(_env, &argList.Nodes_Node, self, &node.Nodes_Node)
                            self.FP.WriteRaw(_env, ", ")
                        }
                    }
                    self.FP.outputAlter2MapFunc(_env, self.FP, prefixType.FP.CreateAlt2typeMap(_env, false))
                    self.FP.WriteRaw(_env, ")")
                    return false
                }
            }
        }
        return true
    }
    if Lns_op_not(convLua_fieldCall(_env)){
        return 
    }
    {
        _refNode := Nodes_ExpRefNodeDownCastF(node.FP.Get_func(_env).FP)
        if !Lns_IsNil( _refNode ) {
            refNode := _refNode.(*Nodes_ExpRefNode)
            if refNode.FP.Get_symbolInfo(_env).FP.Get_name(_env) == "super"{
                wroteFuncFlag = true
                setArgFlag = true
                var funcType *Ast_TypeInfo
                funcType = refNode.FP.Get_expType(_env)
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.%s( self ", Lns_2DDD(self.FP.getFullName(_env, funcType.FP.Get_parentInfo(_env)), funcType.FP.Get_rawTxt(_env))))
            } else { 
                if _switch0 := refNode.FP.Get_expType(_env); _switch0 == self.builtinFunc.Lns_expandLuavalMap {
                    wroteFuncFlag = true
                    self.FP.WriteRaw(_env, "(")
                } else if _switch0 == self.builtinFunc.Lns___run {
                    self.FP.WriteRaw(_env, "_lune._run(")
                    wroteFuncFlag = true
                } else if _switch0 == self.builtinFunc.Lns___join {
                    return 
                } else if _switch0 == self.builtinFunc.Lns___serr {
                    self.FP.WriteRaw(_env, "_lune.LnsErr.create(")
                    wroteFuncFlag = true
                }
            }
        }
    }
    if Lns_op_not(wroteFuncFlag){
        if node.FP.Get_nilAccess(_env){
            self.FP.WriteRaw(_env, "_lune.nilacc( ")
            convLua_filter_9_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ", nil, 'call'")
            wroteFuncFlag = true
            setArgFlag = true
        } else { 
            convLua_filter_9_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, "( ")
        }
    }
    var convStrFlag bool
    convStrFlag = false
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(self.targetLuaVer.FP.Get_canFormStem2Str(_env))) &&
        _env.SetStackVal( self.builtinFunc.FP.IsStrFormFunc(_env, node.FP.Get_func(_env).FP.Get_expType(_env))) ).(bool)){
        convStrFlag = true
    }
    {
        _argList := node.FP.Get_argList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            var expList *LnsList2_[*Nodes_Node]
            expList = NewLnsList2_[*Nodes_Node]([]*Nodes_Node{})
            for _, _expNode := range( argList.FP.Get_expList(_env).Items ) {
                expNode := _expNode
                if expNode.FP.Get_expType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Abbr{
                    {
                        _toDDD := Nodes_ExpToDDDNodeDownCastF(expNode.FP)
                        if !Lns_IsNil( _toDDD ) {
                            toDDD := _toDDD.(*Nodes_ExpToDDDNode)
                            for _, _appNode := range( toDDD.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
                                appNode := _appNode
                                expList.Insert(appNode)
                            }
                        } else {
                            expList.Insert(expNode)
                        }
                    }
                }
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( wroteFuncFlag) &&
                _env.SetStackVal( setArgFlag) ).(bool)){
                if expList.Len() > 0{
                    self.FP.WriteRaw(_env, ", ")
                }
            }
            if convStrFlag{
                var opList *LnsList2_[string]
                opList = NewLnsList2_[string]([]string{})
                if expList.Len() > 0{
                    var literal LnsAny
                    literal = convLua_convExp4_2135(Lns_2DDD(expList.GetAt(1).FP.GetLiteral(_env)))
                    if literal != nil{
                        literal_133 := literal
                        switch _matchExp0 := literal_133.(type) {
                        case *Nodes_Literal__Str:
                            txt := _matchExp0.Val1
                            opList = TransUnit_findForm(_env, txt)
                        }
                    }
                }
                for _index, _argNode := range( expList.Items ) {
                    index := _index + 1
                    argNode := _argNode
                    var filtered bool
                    filtered = false
                    if index > 1{
                        self.FP.WriteRaw(_env, ", ")
                        if index - 1 <= opList.Len(){
                            var formType LnsInt
                            formType = convLua_convExp4_2211(Lns_2DDD(TransUnit_isMatchStringFormatType(_env, opList.GetAt(index - 1), argNode.FP.Get_expType(_env), self.targetLuaVer)))
                            if formType == TransUnit_FormType__NeedConv{
                                self.FP.WriteRaw(_env, "tostring( ")
                                convLua_filter_9_(_env, argNode, self, &node.Nodes_Node)
                                self.FP.WriteRaw(_env, ")")
                                filtered = true
                            }
                        }
                    }
                    if Lns_op_not(filtered){
                        convLua_filter_9_(_env, argNode, self, &node.Nodes_Node)
                    }
                }
            } else { 
                convLua_filter_9_(_env, &argList.Nodes_Node, self, &node.Nodes_Node)
            }
        }
    }
    self.FP.WriteRaw(_env, " )")
}
// 3271: decl @lune.@base.@convLua.ConvFilter.processExpList
func (self *convLua_ConvFilter) ProcessExpList(_env *LnsEnv, node *Nodes_ExpListNode,_opt LnsAny) {
    var expList *LnsList2_[*Nodes_Node]
    expList = node.FP.Get_expList(_env)
    self.FP.processExpListSub(_env, &node.Nodes_Node, expList, node.FP.Get_mRetExp(_env))
}
// 3280: decl @lune.@base.@convLua.ConvFilter.processExpOp1
func (self *convLua_ConvFilter) ProcessExpOp1(_env *LnsEnv, node *Nodes_ExpOp1Node,_opt LnsAny) {
    var op string
    op = node.FP.Get_op(_env).Txt
    if op == ",,,"{
        convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    } else if op == ",,,,"{
        if node.FP.Get_macroMode(_env) != Nodes_MacroMode__None{
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else { 
            self.FP.WriteRaw(_env, "__luneSym2Str( ")
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, " )")
        }
    } else if op == ",,"{
        if _switch0 := node.FP.Get_exp(_env).FP.Get_expType(_env); _switch0 == Ast_builtinTypeInt || _switch0 == Ast_builtinTypeReal || _switch0 == Ast_builtinTypeBool {
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else {
            self.FP.WriteRaw(_env, "__luneGetLocal( ")
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, " )")
        }
    } else if op == "~"{
        if self.targetLuaVer.FP.Get_hasBitOp(_env) == LuaVer_BitOp__HasOp{
            self.FP.WriteRaw(_env, op)
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else { 
            self.FP.WriteRaw(_env, "bit32.bnot( ")
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, " )")
        }
    } else { 
        if op == "not"{
            op = op + " "
        }
        self.FP.WriteRaw(_env, op)
        convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    }
}
// 3329: decl @lune.@base.@convLua.ConvFilter.processExpToDDD
func (self *convLua_ConvFilter) ProcessExpToDDD(_env *LnsEnv, node *Nodes_ExpToDDDNode,_opt LnsAny) {
    self.FP.processExpListSub(_env, &node.Nodes_Node, node.FP.Get_expList(_env).FP.Get_expList(_env), node.FP.Get_expList(_env).FP.Get_mRetExp(_env))
}
// 3335: decl @lune.@base.@convLua.ConvFilter.processExpMultiTo1
func (self *convLua_ConvFilter) ProcessExpMultiTo1(_env *LnsEnv, node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
}
// 3341: decl @lune.@base.@convLua.ConvFilter.processExpCast
func (self *convLua_ConvFilter) ProcessExpCast(_env *LnsEnv, node *Nodes_ExpCastNode,_opt LnsAny) {
    if _switch1 := node.FP.Get_castKind(_env); _switch1 == Nodes_CastKind__Force {
        if node.FP.Get_expType(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeInt, nil, nil){
            self.FP.WriteRaw(_env, "math.floor(")
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        } else if node.FP.Get_expType(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeReal, nil, nil){
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, " * 1.0")
        } else { 
            convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        }
    } else if _switch1 == Nodes_CastKind__Normal {
        self.FP.WriteRaw(_env, "_lune.__Cast( ")
        convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        var castKind LnsInt
        var classObj string
        classObj = "nil"
        if _switch0 := node.FP.Get_expType(_env).FP.Get_nonnilableType(_env); _switch0 == Ast_builtinTypeInt {
            castKind = LuaMod_CastKind__Int
        } else if _switch0 == Ast_builtinTypeReal {
            castKind = LuaMod_CastKind__Real
        } else if _switch0 == Ast_builtinTypeString {
            castKind = LuaMod_CastKind__Str
        } else {
            castKind = LuaMod_CastKind__Class
            classObj = self.FP.getFullName(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env))
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format(", %d, %s )", Lns_2DDD(castKind, classObj)))
    } else if _switch1 == Nodes_CastKind__Implicit {
        convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    }
}
// 3388: decl @lune.@base.@convLua.ConvFilter.processExpParen
func (self *convLua_ConvFilter) ProcessExpParen(_env *LnsEnv, node *Nodes_ExpParenNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "(")
    convLua_filter_9_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, " )")
}
// 3397: decl @lune.@base.@convLua.ConvFilter.processExpSetVal
func (self *convLua_ConvFilter) ProcessExpSetVal(_env *LnsEnv, node *Nodes_ExpSetValNode,_opt LnsAny) {
    {
        __exp := node.FP.Get_condRetInfo(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_CondRetInfo)
            self.FP.outputCondRetInfo(_env, _exp)
        }
    }
    convLua_filter_9_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, " = ")
    convLua_filter_9_(_env, &node.FP.Get_exp2(_env).Nodes_Node, self, &node.Nodes_Node)
}
// 3408: decl @lune.@base.@convLua.ConvFilter.processExpSetItem
func (self *convLua_ConvFilter) ProcessExpSetItem(_env *LnsEnv, node *Nodes_ExpSetItemNode,_opt LnsAny) {
    {
        __exp := node.FP.Get_condRetInfo(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_CondRetInfo)
            self.FP.outputCondRetInfo(_env, _exp)
        }
    }
    convLua_filter_9_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, "[")
    switch _matchExp0 := node.FP.Get_index(_env).(type) {
    case *Nodes_IndexVal__NodeIdx:
        index := _matchExp0.Val1
        convLua_filter_9_(_env, index, self, &node.Nodes_Node)
    case *Nodes_IndexVal__SymIdx:
        index := _matchExp0.Val1
        self.FP.WriteRaw(_env, _env.GetVM().String_format("'%s'", Lns_2DDD(index)))
    }
    self.FP.WriteRaw(_env, "]")
    self.FP.WriteRaw(_env, " = ")
    convLua_filter_9_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
}
// 3430: decl @lune.@base.@convLua.ConvFilter.processExpOp2
func (self *convLua_ConvFilter) ProcessExpOp2(_env *LnsEnv, node *Nodes_ExpOp2Node,_opt LnsAny) {
    var intCast bool
    intCast = false
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_expType(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeInt, nil, nil)) &&
        _env.SetStackVal( node.FP.Get_op(_env).Txt == "/") ).(bool)){
        intCast = true
        self.FP.WriteRaw(_env, "math.floor(")
    }
    var opTxt string
    opTxt = node.FP.Get_op(_env).Txt
    {
        __exp := Ast_bitBinOpMap.Get(opTxt)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(LnsInt)
            if self.targetLuaVer.FP.Get_hasBitOp(_env) == LuaVer_BitOp__HasOp{
                if _switch0 := _exp; _switch0 == Ast_BitOpKind__LShift {
                    opTxt = "<<"
                } else if _switch0 == Ast_BitOpKind__RShift {
                    opTxt = ">>"
                }
                convLua_filter_9_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
                self.FP.WriteRaw(_env, " " + opTxt + " ")
                convLua_filter_9_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
            } else { 
                var binfunc string
                binfunc = ""
                var exp2Mod string
                exp2Mod = ""
                if _switch1 := _exp; _switch1 == Ast_BitOpKind__And {
                    binfunc = "band"
                } else if _switch1 == Ast_BitOpKind__Or {
                    binfunc = "bor"
                } else if _switch1 == Ast_BitOpKind__Xor {
                    binfunc = "bxor"
                } else if _switch1 == Ast_BitOpKind__LShift {
                    binfunc = "lshift"
                } else if _switch1 == Ast_BitOpKind__RShift {
                    binfunc = "lshift"
                    exp2Mod = "-"
                }
                self.FP.WriteRaw(_env, _env.GetVM().String_format("bit32.%s(", Lns_2DDD(binfunc)))
                convLua_filter_9_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
                self.FP.WriteRaw(_env, ", ")
                self.FP.WriteRaw(_env, exp2Mod)
                convLua_filter_9_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
                self.FP.WriteRaw(_env, " )")
            }
        } else {
            convLua_filter_9_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, " " + opTxt + " ")
            convLua_filter_9_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
        }
    }
    if intCast{
        self.FP.WriteRaw(_env, ")")
    }
}
// 3500: decl @lune.@base.@convLua.ConvFilter.processExpRef
func (self *convLua_ConvFilter) ProcessExpRef(_env *LnsEnv, node *Nodes_ExpRefNode,_opt LnsAny) {
    if _switch0 := node.FP.Get_symbolInfo(_env).FP.Get_name(_env); _switch0 == "super" {
        var funcType *Ast_TypeInfo
        funcType = node.FP.Get_expType(_env)
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.%s", Lns_2DDD(self.FP.getFullName(_env, funcType.FP.Get_parentInfo(_env)), funcType.FP.Get_rawTxt(_env))))
    } else {
        var builtinFunc *Builtin_BuiltinFuncType
        builtinFunc = self.builtinFunc
        if node.FP.Get_expType(_env).FP.Equals(_env, self.processInfo, builtinFunc.Lns__load, nil, nil){
            self.FP.WriteRaw(_env, "_lune." + self.targetLuaVer.FP.Get_loadStrFuncName(_env))
        } else { 
            if self.macroVarSymSet.Has(node.FP.Get_symbolInfo(_env).FP.GetOrg(_env)){
                self.FP.WriteRaw(_env, "macroVar.")
            } else { 
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( node.FP.Get_symbolInfo(_env).FP.Get_accessMode(_env) == Ast_AccessMode__Pub) &&
                    _env.SetStackVal( node.FP.Get_symbolInfo(_env).FP.Get_kind(_env) == Ast_SymbolKind__Var) ).(bool)){
                    if self.needModuleObj{
                        self.FP.WriteRaw(_env, "_moduleObj.")
                    }
                }
            }
            self.FP.WriteRaw(_env, node.FP.Get_symbolInfo(_env).FP.Get_name(_env))
            if node.FP.Get_symbolInfo(_env).FP.Get_isLazyLoad(_env){
                self.FP.WriteRaw(_env, "()")
            }
        }
    }
}
// 3536: decl @lune.@base.@convLua.ConvFilter.processExpRefItem
func (self *convLua_ConvFilter) ProcessExpRefItem(_env *LnsEnv, node *Nodes_ExpRefItemNode,_opt LnsAny) {
    if node.FP.Get_nilAccess(_env){
        self.FP.WriteRaw(_env, "_lune.nilacc( ")
        convLua_filter_9_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, ", nil, 'item', ")
        {
            __exp := node.FP.Get_index(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_Node)
                convLua_filter_9_(_env, _exp, self, &node.Nodes_Node)
            } else {
                self.FP.WriteRaw(_env, _env.GetVM().String_format("'%s'", Lns_2DDD(Lns_unwrap( node.FP.Get_symbol(_env)).(string))))
            }
        }
        self.FP.WriteRaw(_env, ")")
    } else { 
        if node.FP.Get_val(_env).FP.Get_expType(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeString, nil, nil){
            self.FP.WriteRaw(_env, "string.byte( ")
            convLua_filter_9_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ", ")
            {
                __exp := node.FP.Get_index(_env)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Nodes_Node)
                    convLua_filter_9_(_env, _exp, self, &node.Nodes_Node)
                } else {
                    Util_err(_env, "index is nil")
                }
            }
            self.FP.WriteRaw(_env, " )")
        } else { 
            convLua_filter_9_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, "[")
            {
                __exp := node.FP.Get_index(_env)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Nodes_Node)
                    convLua_filter_9_(_env, _exp, self, &node.Nodes_Node)
                } else {
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("'%s'", Lns_2DDD(Lns_unwrap( node.FP.Get_symbol(_env)).(string))))
                }
            }
            self.FP.WriteRaw(_env, "]")
        }
    }
}
// 3579: decl @lune.@base.@convLua.ConvFilter.processRefField
func (self *convLua_ConvFilter) ProcessRefField(_env *LnsEnv, node *Nodes_RefFieldNode,_opt LnsAny) {
    opt := _opt.(*ConvLua_Opt)
    {
        _symbol := node.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( _symbol ) {
            symbol := _symbol.(*Ast_SymbolInfo)
            {
                _code := self.builtinSym2code.Get(symbol.FP.GetOrg(_env))
                if !Lns_IsNil( _code ) {
                    code := _code.(string)
                    self.FP.WriteRaw(_env, code)
                    return 
                }
            }
        }
    }
    var parent *Nodes_Node
    parent = opt.Node
    var prefix *Nodes_Node
    prefix = node.FP.Get_prefix(_env)
    if node.FP.Get_nilAccess(_env){
        self.FP.WriteRaw(_env, "_lune.nilacc( ")
        convLua_filter_9_(_env, prefix, self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, _env.GetVM().String_format(", \"%s\" )", Lns_2DDD(node.FP.Get_field(_env).Txt)))
    } else { 
        convLua_filter_9_(_env, prefix, self, &node.Nodes_Node)
        var delimit string
        delimit = "."
        if parent.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpCall(_env){
            if node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
                delimit = ":"
            } else { 
                delimit = "."
            }
        }
        var fieldToken *Types_Token
        fieldToken = node.FP.Get_field(_env)
        self.FP.WriteRaw(_env, delimit + fieldToken.Txt)
        {
            _symbolInfo := node.FP.Get_symbolInfo(_env)
            if !Lns_IsNil( _symbolInfo ) {
                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                if symbolInfo.FP.Get_isLazyLoad(_env){
                    self.FP.WriteRaw(_env, "()")
                }
            }
        }
    }
}
// 3618: decl @lune.@base.@convLua.ConvFilter.processExpOmitEnum
func (self *convLua_ConvFilter) ProcessExpOmitEnum(_env *LnsEnv, node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    {
        _aliasType := node.FP.Get_aliasType(_env)
        if !Lns_IsNil( _aliasType ) {
            aliasType := _aliasType.(*Ast_AliasTypeInfo)
            self.FP.WriteRaw(_env, self.FP.getFullName(_env, &aliasType.Ast_TypeInfo))
        } else {
            self.FP.WriteRaw(_env, self.FP.getFullName(_env, node.FP.Get_expType(_env)))
        }
    }
    self.FP.WriteRaw(_env, _env.GetVM().String_format(".%s", Lns_2DDD(node.FP.Get_valToken(_env).Txt)))
}
// 3629: decl @lune.@base.@convLua.ConvFilter.processGetField
func (self *convLua_ConvFilter) ProcessGetField(_env *LnsEnv, node *Nodes_GetFieldNode,_opt LnsAny) {
    var prefixNode *Nodes_Node
    prefixNode = node.FP.Get_prefix(_env)
    var prefixType *Ast_TypeInfo
    prefixType = prefixNode.FP.Get_expType(_env)
    var fieldTxt string
    fieldTxt = node.FP.Get_field(_env).Txt
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( fieldTxt == "_txt") &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) ||
            _env.SetStackVal( prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alge) ).(bool))) ).(bool)){
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s:_getTxt( ", Lns_2DDD(self.FP.getFullName(_env, prefixType))))
        convLua_filter_9_(_env, prefixNode, self, &node.Nodes_Node)
        self.FP.writeln(_env, ")")
    } else { 
        if node.FP.Get_nilAccess(_env){
            fieldTxt = _env.GetVM().String_format("get_%s", Lns_2DDD(fieldTxt))
            self.FP.WriteRaw(_env, "_lune.nilacc( ")
            convLua_filter_9_(_env, prefixNode, self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, _env.GetVM().String_format(", '%s', 'callmtd' )", Lns_2DDD(fieldTxt)))
        } else { 
            fieldTxt = _env.GetVM().String_format("get_%s()", Lns_2DDD(fieldTxt))
            convLua_filter_9_(_env, prefixNode, self, &node.Nodes_Node)
            var delimit string
            delimit = "."
            if node.FP.Get_getterTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
                delimit = ":"
            } else { 
                delimit = "."
            }
            self.FP.WriteRaw(_env, delimit + fieldTxt)
        }
    }
}
// 3668: decl @lune.@base.@convLua.ConvFilter.processReturn
func (self *convLua_ConvFilter) ProcessReturn(_env *LnsEnv, node *Nodes_ReturnNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "return ")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, "")
}
// 3679: decl @lune.@base.@convLua.ConvFilter.processLuneKind
func (self *convLua_ConvFilter) ProcessLuneKind(_env *LnsEnv, node *Nodes_LuneKindNode,_opt LnsAny) {
    {
        _workNode := Nodes_ExpCastNodeDownCastF(node.FP.Get_exp(_env).FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes_ExpCastNode)
            if workNode.FP.Get_castKind(_env) == Nodes_CastKind__Implicit{
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%d", Lns_2DDD(workNode.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env))))
            }
        } else {
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%d", Lns_2DDD(node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env))))
        }
    }
}
// 3692: decl @lune.@base.@convLua.ConvFilter.processTestCase
func (self *convLua_ConvFilter) ProcessTestCase(_env *LnsEnv, node *Nodes_TestCaseNode,_opt LnsAny) {
    if self.enableTest{
        self.FP.writeln(_env, "do")
        self.FP.pushIndent(_env, nil)
        convLua_filter_9_(_env, node.FP.Get_impNode(_env), self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
        self.FP.writeln(_env, _env.GetVM().String_format("local function testcase( %s ) ", Lns_2DDD(node.FP.Get_ctrlName(_env))))
        convLua_filter_9_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        self.FP.writeln(_env, "end")
        self.FP.writeln(_env, _env.GetVM().String_format("__t.registerTestcase( \"%s\", \"%s\", testcase )", Lns_2DDD(Lns_car(_env.GetVM().String_gsub(self.moduleTypeInfo.FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), nil),"@", "")).(string), node.FP.Get_name(_env).Txt)))
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "end")
    }
}
// 3714: decl @lune.@base.@convLua.ConvFilter.processTestBlock
func (self *convLua_ConvFilter) ProcessTestBlock(_env *LnsEnv, node *Nodes_TestBlockNode,_opt LnsAny) {
    if self.enableTest{
        for _, _statement := range( node.FP.Get_stmtList(_env).Items ) {
            statement := _statement
            convLua_filter_9_(_env, statement, self, &node.Nodes_Node)
            self.FP.writeln(_env, "")
        }
    }
}
// 3725: decl @lune.@base.@convLua.ConvFilter.processProvide
func (self *convLua_ConvFilter) ProcessProvide(_env *LnsEnv, node *Nodes_ProvideNode,_opt LnsAny) {
}
// 3730: decl @lune.@base.@convLua.ConvFilter.processAlias
func (self *convLua_ConvFilter) ProcessAlias(_env *LnsEnv, node *Nodes_AliasNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("local %s = ", Lns_2DDD(node.FP.Get_newSymbol(_env).FP.Get_name(_env))))
    convLua_filter_9_(_env, node.FP.Get_srcNode(_env), self, &node.Nodes_Node)
    if Ast_isPubToExternal(_env, node.FP.Get_expType(_env).FP.Get_accessMode(_env)){
        self.FP.WriteRaw(_env, _env.GetVM().String_format("\n_moduleObj.%s = %s", Lns_2DDD(node.FP.Get_newSymbol(_env).FP.Get_name(_env), node.FP.Get_newSymbol(_env).FP.Get_name(_env))))
    }
}
// 3740: decl @lune.@base.@convLua.ConvFilter.processBoxing
func (self *convLua_ConvFilter) ProcessBoxing(_env *LnsEnv, node *Nodes_BoxingNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "{")
    convLua_filter_9_(_env, node.FP.Get_src(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, "}")
}
// 3750: decl @lune.@base.@convLua.ConvFilter.processUnboxing
func (self *convLua_ConvFilter) ProcessUnboxing(_env *LnsEnv, node *Nodes_UnboxingNode,_opt LnsAny) {
    convLua_filter_9_(_env, node.FP.Get_src(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, "[1]")
}
// 3757: decl @lune.@base.@convLua.ConvFilter.processTupleConst
func (self *convLua_ConvFilter) ProcessTupleConst(_env *LnsEnv, node *Nodes_TupleConstNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "{")
    convLua_filter_9_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, "}")
}
// 3767: decl @lune.@base.@convLua.ConvFilter.processLiteralList
func (self *convLua_ConvFilter) ProcessLiteralList(_env *LnsEnv, node *Nodes_LiteralListNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "{")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.WriteRaw(_env, "}")
}
// 3779: decl @lune.@base.@convLua.ConvFilter.processLiteralSet
func (self *convLua_ConvFilter) ProcessLiteralSet(_env *LnsEnv, node *Nodes_LiteralSetNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "{")
    {
        _expListNode := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expListNode ) {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            for _index, _expNode := range( expListNode.FP.Get_expList(_env).Items ) {
                index := _index + 1
                expNode := _expNode
                if index > 1{
                    self.FP.WriteRaw(_env, ", ")
                }
                self.FP.WriteRaw(_env, "[")
                convLua_filter_9_(_env, expNode, self, &node.Nodes_Node)
                self.FP.WriteRaw(_env, "] = true")
            }
        }
    }
    self.FP.WriteRaw(_env, "}")
}
// 3797: decl @lune.@base.@convLua.ConvFilter.processLiteralMap
func (self *convLua_ConvFilter) ProcessLiteralMap(_env *LnsEnv, node *Nodes_LiteralMapNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "{")
    var pairList *LnsList2_[*Nodes_PairItem]
    pairList = node.FP.Get_pairList(_env)
    for _index, _pair := range( pairList.Items ) {
        index := _index + 1
        pair := _pair
        if index > 1{
            self.FP.WriteRaw(_env, ", ")
        }
        self.FP.WriteRaw(_env, "[")
        convLua_filter_9_(_env, pair.FP.Get_key(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, "] = ")
        convLua_filter_9_(_env, pair.FP.Get_val(_env), self, &node.Nodes_Node)
    }
    self.FP.WriteRaw(_env, "}")
}
// 3815: decl @lune.@base.@convLua.ConvFilter.processLiteralArray
func (self *convLua_ConvFilter) ProcessLiteralArray(_env *LnsEnv, node *Nodes_LiteralArrayNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "{")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_9_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.WriteRaw(_env, "}")
}
// 3828: decl @lune.@base.@convLua.ConvFilter.processLiteralChar
func (self *convLua_ConvFilter) ProcessLiteralChar(_env *LnsEnv, node *Nodes_LiteralCharNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%d", Lns_2DDD(node.FP.Get_num(_env))))
}
// 3834: decl @lune.@base.@convLua.ConvFilter.processLiteralInt
func (self *convLua_ConvFilter) ProcessLiteralInt(_env *LnsEnv, node *Nodes_LiteralIntNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, node.FP.Get_token(_env).Txt)
}
// 3842: decl @lune.@base.@convLua.ConvFilter.processLiteralReal
func (self *convLua_ConvFilter) ProcessLiteralReal(_env *LnsEnv, node *Nodes_LiteralRealNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, node.FP.Get_token(_env).Txt)
}
// 3849: decl @lune.@base.@convLua.ConvFilter.processLiteralString
func (self *convLua_ConvFilter) ProcessLiteralString(_env *LnsEnv, node *Nodes_LiteralStringNode,_opt LnsAny) {
    var txt string
    txt = node.FP.Get_token(_env).Txt
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(txt, "^```", nil, nil))){
        txt = "[==[" + _env.GetVM().String_sub(txt,4, -4) + "]==]"
    }
    var opList *LnsList2_[string]
    opList = TransUnit_findForm(_env, txt)
    {
        _expList := node.FP.Get_orgParam(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            var mRetIndex LnsAny
            mRetIndex = _env.NilAccFin(_env.NilAccPush(expList.FP.Get_mRetExp(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
            self.FP.WriteRaw(_env, _env.GetVM().String_format("string.format( %s, ", Lns_2DDD(txt)))
            for _index, _val := range( expList.FP.Get_expList(_env).Items ) {
                index := _index + 1
                val := _val
                if index > 1{
                    self.FP.WriteRaw(_env, ", ")
                }
                var matchFlag LnsInt
                matchFlag = TransUnit_FormType__Match
                if index <= opList.Len(){
                    matchFlag = convLua_convExp0_3743(Lns_2DDD(TransUnit_isMatchStringFormatType(_env, opList.GetAt(index), val.FP.Get_expType(_env), self.targetLuaVer)))
                }
                if matchFlag == TransUnit_FormType__NeedConv{
                    self.FP.WriteRaw(_env, "tostring( ")
                    convLua_filter_9_(_env, val, self, &node.Nodes_Node)
                    self.FP.WriteRaw(_env, ")")
                } else { 
                    convLua_filter_9_(_env, val, self, &node.Nodes_Node)
                }
                if index == mRetIndex{
                    break
                }
            }
            self.FP.WriteRaw(_env, ")")
        } else {
            self.FP.WriteRaw(_env, txt)
        }
    }
}
// 3891: decl @lune.@base.@convLua.ConvFilter.processLiteralBool
func (self *convLua_ConvFilter) ProcessLiteralBool(_env *LnsEnv, node *Nodes_LiteralBoolNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, node.FP.Get_token(_env).Txt)
}
// 3897: decl @lune.@base.@convLua.ConvFilter.processLiteralNil
func (self *convLua_ConvFilter) ProcessLiteralNil(_env *LnsEnv, node *Nodes_LiteralNilNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "nil")
}
// 3903: decl @lune.@base.@convLua.ConvFilter.processBreak
func (self *convLua_ConvFilter) ProcessBreak(_env *LnsEnv, node *Nodes_BreakNode,_opt LnsAny) {
    if node.FP.Get_asyncLockBreak(_env){
        self.FP.writeRawln(_env, _env.GetVM().String_format("%s = true", Lns_2DDD(convLua_asyncLockBreakSym)))
        self.FP.writeRawln(_env, "return")
    } else { 
        self.FP.writeRawln(_env, "break")
    }
}
// 3914: decl @lune.@base.@convLua.ConvFilter.processLiteralSymbol
func (self *convLua_ConvFilter) ProcessLiteralSymbol(_env *LnsEnv, node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s", Lns_2DDD(node.FP.Get_token(_env).Txt)))
}
// 3920: decl @lune.@base.@convLua.ConvFilter.processLuneControl
func (self *convLua_ConvFilter) ProcessLuneControl(_env *LnsEnv, node *Nodes_LuneControlNode,_opt LnsAny) {
    switch node.FP.Get_pragma(_env).(type) {
    case *LuneControl_Pragma__load__lune_module:
        self.FP.processLoadRuntime(_env)
    }
}
// 3933: decl @lune.@base.@convLua.FilterInfo.outputLuaAndMeta
func (self *ConvLua_FilterInfo) OutputLuaAndMeta(_env *LnsEnv, node *Nodes_RootNode) {
    node.FP.ProcessFilter(_env, self.filter, ConvLua_Opt2Stem(NewConvLua_Opt(_env, &node.Nodes_Node)))
}
// 3940: decl @lune.@base.@convLua.FilterInfo.outputLua
func (self *ConvLua_FilterInfo) OutputLua(_env *LnsEnv, node *Nodes_RootNode) {
    node.FP.ProcessFilter(_env, self.filter, ConvLua_Opt2Stem(NewConvLua_Opt(_env, &node.Nodes_Node)))
}
// 3945: decl @lune.@base.@convLua.FilterInfo.outputMeta
func (self *ConvLua_FilterInfo) OutputMeta(_env *LnsEnv, node *Nodes_RootNode) {
}
// 3976: decl @lune.@base.@convLua.MacroEvalImp.evalFromCodeToLuaCode
func (self *ConvLua_MacroEvalImp) EvalFromCodeToLuaCode(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,argNameList *LnsList2_[string],code LnsAny) string {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var conv *convLua_ConvFilter
    conv = NewconvLua_ConvFilter(_env, "macro", stream.FP, NewUtil_NullOStream(_env).FP, ConvLua_ConvMode__ConvMeta, true, Ast_headTypeInfo, processInfo, Ast_SymbolKind__Typ, self.builtinFunc, nil, LuaVer_getCurVer(_env), false, true, NewConvLua_Option(_env, "", false))
    conv.FP.OutputDeclMacro(_env, name, argNameList, convLua_outputMacroStmtBlock_13_(func(_env *LnsEnv) {
        if code != nil{
            code_922 := code.(string)
            conv.FP.Write(_env, code_922)
        }
    }))
    return stream.FP.Get_txt(_env)
}
// 3996: decl @lune.@base.@convLua.MacroEvalImp.evalToLuaCode
func (self *ConvLua_MacroEvalImp) EvalToLuaCode(_env *LnsEnv, processInfo *Ast_ProcessInfo,node *Nodes_DeclMacroNode) string {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var conv *convLua_ConvFilter
    conv = NewconvLua_ConvFilter(_env, "macro", stream.FP, NewUtil_NullOStream(_env).FP, ConvLua_ConvMode__ConvMeta, true, Ast_headTypeInfo, processInfo, Ast_SymbolKind__Typ, self.builtinFunc, nil, LuaVer_getCurVer(_env), false, true, NewConvLua_Option(_env, "", false))
    conv.FP.ProcessDeclMacro(_env, node, NewConvLua_Opt(_env, &node.Nodes_Node))
    return stream.FP.Get_txt(_env)
}
