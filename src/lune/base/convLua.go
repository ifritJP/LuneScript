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
type convLua_outputMacroStmtBlock_11_ func (_env *LnsEnv)
// for 3783
func convLua_convExp18876(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 3549
func convLua_convExp17910(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 3808
func convLua_convExp1956(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 298
func convLua_convExp2409(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1346
func convLua_convExp7332(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1347
func convLua_convExp7347(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1357
func convLua_convExp7409(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1365
func convLua_convExp7460(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1433
func convLua_convExp7829(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1495
func convLua_convExp8102(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1826
func convLua_convExp9694(arg1 []LnsAny) (string, bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(bool), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 2933
func convLua_convExp15065(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2948
func convLua_convExp15141(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 3711
func convLua_convExp18441(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 3781
func convLua_convExp18624(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 78: decl @lune.@base.@convLua.getSymbolTxt
func convLua_getSymbolTxt_5_(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) string {
    if symbolInfo.FP.Get_name(_env) == "_"{
        {
            _annonymous := Ast_AnonymousSymbolInfoDownCastF(symbolInfo.FP)
            if !Lns_IsNil( _annonymous ) {
                annonymous := _annonymous.(*Ast_AnonymousSymbolInfo)
                return _env.LuaVM.String_format("_%d", []LnsAny{annonymous.FP.Get_anonymousId(_env)})
            } else {
                Util_err(_env, _env.LuaVM.String_format("can't cast to AnonymousSymbolInfo. -- %s:%s", []LnsAny{_env.NilAccFin(_env.NilAccPush(symbolInfo.FP.Get_pos(_env)) && 
                _env.NilAccPush(_env.NilAccPop().(*Types_Position).LineNo)), _env.NilAccFin(_env.NilAccPush(symbolInfo.FP.Get_pos(_env)) && 
                _env.NilAccPush(_env.NilAccPop().(*Types_Position).Column))}))
            }
        }
    }
    return symbolInfo.FP.Get_name(_env)
}

// 231: decl @lune.@base.@convLua.filter
func convLua_filter_7_(_env *LnsEnv, node *Nodes_Node,filter *convLua_ConvFilter,parent *Nodes_Node) {
    node.FP.ProcessFilter(_env, &filter.Nodes_Filter, ConvLua_Opt2Stem(NewConvLua_Opt(_env, parent)))
}

// 1274: decl @lune.@base.@convLua.isGenericType
func convLua_isGenericType_10_(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if Ast_isGenericType(_env, typeInfo){
        return true
    }
    if _switch1 := typeInfo.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Class || _switch1 == Ast_TypeInfoKind__IF {
        if typeInfo.FP.Get_itemTypeInfoList(_env).Len() > 0{
            return true
        }
    }
    return false
}

// 3627: decl @lune.@base.@convLua.createFilter
func ConvLua_createFilter(_env *LnsEnv, streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,moduleTypeInfo *Ast_TypeInfo,processInfo *Ast_ProcessInfo,moduleSymbolKind LnsInt,builtinFunc *Builtin_BuiltinFuncType,useLuneRuntime LnsAny,targetLuaVer *LuaVer_LuaVerInfo,enableTest bool,useIpairs bool) *ConvLua_FilterInfo {
    var convFilter *convLua_ConvFilter
    convFilter = NewconvLua_ConvFilter(_env, streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, processInfo, moduleSymbolKind, builtinFunc, useLuneRuntime, targetLuaVer, enableTest, useIpairs)
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
    outputAlter2MapFunc(_env *LnsEnv, arg1 Lns_oStream, arg2 *LnsMap)
    OutputDeclMacro(_env *LnsEnv, arg1 string, arg2 *LnsList, arg3 convLua_outputMacroStmtBlock_11_)
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
    ProcessDeclVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode, arg2 LnsAny)
    ProcessExpAccessMRet(_env *LnsEnv, arg1 *Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(_env *LnsEnv, arg1 *Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(_env *LnsEnv, arg1 *Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(_env *LnsEnv, arg1 *Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpList(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 LnsAny)
    processExpListSub(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsList, arg3 LnsAny)
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
    ProcessReturn(_env *LnsEnv, arg1 *Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(_env *LnsEnv, arg1 *Nodes_RootNode, arg2 LnsAny)
    ProcessScope(_env *LnsEnv, arg1 *Nodes_ScopeNode, arg2 LnsAny)
    ProcessShebang(_env *LnsEnv, arg1 *Nodes_ShebangNode, arg2 LnsAny)
    ProcessStmtExp(_env *LnsEnv, arg1 *Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(_env *LnsEnv, arg1 *Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(_env *LnsEnv, arg1 *Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(_env *LnsEnv, arg1 *Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(_env *LnsEnv, arg1 *Nodes_TestCaseNode, arg2 LnsAny)
    ProcessUnboxing(_env *LnsEnv, arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(_env *LnsEnv, arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(_env *LnsEnv, arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(_env *LnsEnv, arg1 *Nodes_WhileNode, arg2 LnsAny)
    process__func__symbol(_env *LnsEnv, arg1 bool, arg2 *Ast_TypeInfo, arg3 string)
    pushIndent(_env *LnsEnv, arg1 LnsAny)
    Write(_env *LnsEnv, arg1 string)(LnsAny, LnsAny)
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
    indentQueue *LnsList
    curLineNo LnsInt
    classId2TypeInfo *LnsMap
    classId2MemberList *LnsMap
    pubEnumId2EnumTypeInfo *LnsMap
    pubAlgeId2AlgeTypeInfo *LnsMap
    pubVarName2InfoMap *LnsMap
    pubFuncName2InfoMap *LnsMap
    needIndent bool
    macroDepth LnsInt
    macroVarSymSet *LnsSet
    moduleTypeInfo *Ast_TypeInfo
    moduleSymbolKind LnsInt
    needModuleObj bool
    useLuneRuntime LnsAny
    targetLuaVer *LuaVer_LuaVerInfo
    enableTest bool
    useIpairs bool
    processInfo *Ast_ProcessInfo
    moduleType2SymbolMap *LnsMap
    builtinFunc *Builtin_BuiltinFuncType
    builtinSym2code *LnsMap
    FP convLua_ConvFilterMtd
}
func convLua_ConvFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convLua_ConvFilter).FP
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
func NewconvLua_ConvFilter(_env *LnsEnv, arg1 string, arg2 Lns_oStream, arg3 Lns_oStream, arg4 LnsInt, arg5 bool, arg6 *Ast_TypeInfo, arg7 *Ast_ProcessInfo, arg8 LnsInt, arg9 *Builtin_BuiltinFuncType, arg10 LnsAny, arg11 *LuaVer_LuaVerInfo, arg12 bool, arg13 bool) *convLua_ConvFilter {
    obj := &convLua_ConvFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvLua_ConvFilter(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
    return obj
}
// 129: DeclConstr
func (self *convLua_ConvFilter) InitconvLua_ConvFilter(_env *LnsEnv, streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,moduleTypeInfo *Ast_TypeInfo,processInfo *Ast_ProcessInfo,moduleSymbolKind LnsInt,builtinFunc *Builtin_BuiltinFuncType,useLuneRuntime LnsAny,targetLuaVer *LuaVer_LuaVerInfo,enableTest bool,useIpairs bool) {
    self.InitNodes_Filter(_env, true, moduleTypeInfo, moduleTypeInfo.FP.Get_scope(_env))
    if stream == metaStream{
        Util_err(_env, "streamName == stream")
    }
    self.builtinFunc = builtinFunc
    
    self.moduleType2SymbolMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.processInfo = processInfo
    
    self.enableTest = enableTest
    
    self.macroVarSymSet = NewLnsSet([]LnsAny{})
    
    self.needModuleObj = true
    
    self.indentQueue = NewLnsList([]LnsAny{0})
    
    self.moduleSymbolKind = moduleSymbolKind
    
    self.macroDepth = 0
    
    self.streamName = streamName
    
    self.stream = stream
    
    self.metaStream = metaStream
    
    self.outMetaFlag = false
    
    self.convMode = convMode
    
    self.inMacro = inMacro
    
    self.curLineNo = 1
    
    self.classId2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.classId2MemberList = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.pubVarName2InfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.pubFuncName2InfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.pubEnumId2EnumTypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.pubAlgeId2AlgeTypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.needIndent = false
    
    self.moduleTypeInfo = moduleTypeInfo
    
    self.useLuneRuntime = useLuneRuntime
    
    self.targetLuaVer = targetLuaVer
    
    self.useIpairs = useIpairs
    
    self.builtinSym2code = NewLnsMap( map[LnsAny]LnsAny{builtinFunc.G__lns_runmode_Sync_sym:_env.LuaVM.String_format("%d", []LnsAny{0}),builtinFunc.G__lns_runmode_Queue_sym:_env.LuaVM.String_format("%d", []LnsAny{1}),builtinFunc.G__lns_runmode_Skip_sym:_env.LuaVM.String_format("%d", []LnsAny{2}),})
    
}

// 178: decl @lune.@base.@convLua.ConvFilter.get_indent
func (self *convLua_ConvFilter) get_indent(_env *LnsEnv) LnsInt {
    if self.indentQueue.Len() > 0{
        return self.indentQueue.GetAt(self.indentQueue.Len()).(LnsInt)
    }
    return 0
}

// 185: decl @lune.@base.@convLua.ConvFilter.getCanonicalName
func (self *convLua_ConvFilter) getCanonicalName(_env *LnsEnv, typeInfo *Ast_TypeInfo,localFlag bool) string {
    var enumName string
    enumName = typeInfo.FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), localFlag)
    var moduleType *Ast_TypeInfo
    moduleType = typeInfo.FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.GetModule(_env)
    var canonical string
    canonical = Lns_car(_env.LuaVM.String_gsub(enumName,"&", "")).(string)
    {
        _assignSym := self.moduleType2SymbolMap.Get(moduleType)
        if !Lns_IsNil( _assignSym ) {
            assignSym := _assignSym.(*Ast_SymbolInfo)
            if assignSym.FP.Get_isLazyLoad(_env){
                var index LnsInt
                index = Lns_unwrap( Lns_car(_env.LuaVM.String_find(canonical,".", 1, true))).(LnsInt)
                return _env.LuaVM.String_format("%s().%s", []LnsAny{_env.LuaVM.String_sub(canonical,1, index - 1), _env.LuaVM.String_sub(canonical,index + 1, nil)})
            }
        }
    }
    return canonical
}

// 200: decl @lune.@base.@convLua.ConvFilter.getFullName
func (self *convLua_ConvFilter) getFullName(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    return self.FP.getCanonicalName(_env, typeInfo, true)
}

// 204: decl @lune.@base.@convLua.ConvFilter.close
func (self *convLua_ConvFilter) Close(_env *LnsEnv) {
}

// 206: decl @lune.@base.@convLua.ConvFilter.flush
func (self *convLua_ConvFilter) Flush(_env *LnsEnv) {
}

// 209: decl @lune.@base.@convLua.ConvFilter.write
func (self *convLua_ConvFilter) Write(_env *LnsEnv, txt string)(LnsAny, LnsAny) {
    var stream Lns_oStream
    stream = self.stream
    if self.outMetaFlag{
        stream = self.metaStream
        
    }
    for _, _line := range( Str_getLineList(_env, txt).Items ) {
        line := _line.(string)
        if self.needIndent{
            stream.Write(_env, _env.LuaVM.String_rep(" ", self.FP.get_indent(_env)))
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


// 246: decl @lune.@base.@convLua.ConvFilter.processBlankLine
func (self *convLua_ConvFilter) ProcessBlankLine(_env *LnsEnv, node *Nodes_BlankLineNode,_opt LnsAny) {
}

// 246: decl @lune.@base.@convLua.ConvFilter.processDeclForm
func (self *convLua_ConvFilter) ProcessDeclForm(_env *LnsEnv, node *Nodes_DeclFormNode,_opt LnsAny) {
}

// 246: decl @lune.@base.@convLua.ConvFilter.processProtoMethod
func (self *convLua_ConvFilter) ProcessProtoMethod(_env *LnsEnv, node *Nodes_ProtoMethodNode,_opt LnsAny) {
}

// 264: decl @lune.@base.@convLua.ConvFilter.pushIndent
func (self *convLua_ConvFilter) pushIndent(_env *LnsEnv, newIndent LnsAny) {
    var indent LnsInt
    indent = Lns_unwrapDefault( newIndent, self.FP.get_indent(_env) + convLua_stepIndent).(LnsInt)
    self.indentQueue.Insert(indent)
}

// 269: decl @lune.@base.@convLua.ConvFilter.popIndent
func (self *convLua_ConvFilter) popIndent(_env *LnsEnv) {
    if self.indentQueue.Len() == 0{
        Util_err(_env, "self.indentQueue == 0")
    }
    self.indentQueue.Remove(nil)
}

// 277: decl @lune.@base.@convLua.ConvFilter.writeln
func (self *convLua_ConvFilter) writeln(_env *LnsEnv, txt string) {
    self.FP.Write(_env, txt)
    self.FP.Write(_env, "\n")
    self.needIndent = true
    
}

// 283: decl @lune.@base.@convLua.ConvFilter.processNone
func (self *convLua_ConvFilter) ProcessNone(_env *LnsEnv, node *Nodes_NoneNode,_opt LnsAny) {
}

// 288: decl @lune.@base.@convLua.ConvFilter.processShebang
func (self *convLua_ConvFilter) ProcessShebang(_env *LnsEnv, node *Nodes_ShebangNode,_opt LnsAny) {
}

// 293: decl @lune.@base.@convLua.ConvFilter.processImport
func (self *convLua_ConvFilter) ProcessImport(_env *LnsEnv, node *Nodes_ImportNode,_opt LnsAny) {
    var info *Nodes_ImportInfo
    info = node.FP.Get_info(_env)
    var modulePath string
    modulePath = info.FP.Get_modulePath(_env)
    var modSym string
    modSym = convLua_convExp2409(Lns_2DDD(_env.LuaVM.String_gsub(modulePath,".*%.", "")))
    modSym = info.FP.Get_assignName(_env)
    
    self.FP.Write(_env, _env.LuaVM.String_format("local %s = _lune.", []LnsAny{modSym}))
    if _switch1 := info.FP.Get_lazy(_env); _switch1 == Nodes_LazyLoad__Off {
        self.FP.Write(_env, "loadModule")
    } else if _switch1 == Nodes_LazyLoad__On || _switch1 == Nodes_LazyLoad__Auto {
        self.FP.Write(_env, "_lazyImport")
    }
    self.FP.Write(_env, _env.LuaVM.String_format("( '%s' )", []LnsAny{modulePath}))
}

// 339: decl @lune.@base.@convLua.ConvFilter.outputMeta
func (self *convLua_ConvFilter) outputMeta(_env *LnsEnv, node *Nodes_RootNode) {
    if self.convMode == ConvLua_ConvMode__Convert{
        return 
    }
    self.outMetaFlag = true
    
    self.FP.writeln(_env, "local _moduleObj = {}")
    self.FP.writeln(_env, "----- meta -----")
    self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.__version = '%s'", []LnsAny{Ver_version}))
    self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.__formatVersion = '%s'", []LnsAny{Ver_metaVersion}))
    self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.__buildId = %q", []LnsAny{node.FP.Get_moduleId(_env).FP.GetNextModuleId(_env).FP.Get_idStr(_env)}))
    self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.__enableTest = %s", []LnsAny{self.enableTest}))
    self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.__hasTest = %s", []LnsAny{node.FP.Get_nodeManager(_env).FP.GetTestCaseNodeList(_env).Len() != 0}))
    self.FP.Write(_env, "_moduleObj.__lazyModuleList = {")
    {
        var firstFlag bool
        firstFlag = true
        for _, _declClass := range( node.FP.Get_nodeManager(_env).FP.GetDeclClassNodeList(_env).Items ) {
            declClass := _declClass.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( declClass.FP.Get_lazyLoad(_env) != Nodes_LazyLoad__Off) &&
                _env.SetStackVal( Ast_isPubToExternal(_env, declClass.FP.Get_accessMode(_env))) ).(bool)){
                if Lns_op_not(firstFlag){
                    self.FP.Write(_env, ",")
                } else { 
                    firstFlag = false
                    
                }
                self.FP.Write(_env, _env.LuaVM.String_format("%d", []LnsAny{declClass.FP.Get_expType(_env).FP.Get_typeId(_env).Id}))
            }
        }
    }
    self.FP.writeln(_env, "}")
    var importModuleType2Index *LnsMap
    importModuleType2Index = NewLnsMap( map[LnsAny]LnsAny{})
    var importProcessInfo2Index *LnsMap
    importProcessInfo2Index = NewLnsMap( map[LnsAny]LnsAny{})
    importProcessInfo2Index.Set(Ast_getRootProcessInfoRo(_env),FrontInterface_getRootDependModId(_env))
    importProcessInfo2Index.Set(self.processInfo,0)
    importProcessInfo2Index.Set(self.processInfo.FP.Get_orgInfo(_env),0)
    var importNameMap *LnsMap
    importNameMap = NewLnsMap( map[LnsAny]LnsAny{})
    {
        for _, _moduleInfo := range( node.FP.Get_importModule2moduleInfo(_env).Items ) {
            moduleInfo := _moduleInfo.(FrontInterface_ModuleInfoDownCast).ToFrontInterface_ModuleInfo()
            importNameMap.Set(moduleInfo.FP.Get_fullName(_env),moduleInfo)
        }
        var index LnsInt
        index = 0
        {
            __forsortCollection1 := importNameMap
            __forsortSorted1 := __forsortCollection1.CreateKeyListStr()
            __forsortSorted1.Sort( _env, LnsItemKindStr, nil )
            for _, ___forsortKey1 := range( __forsortSorted1.Items ) {
                moduleInfo := __forsortCollection1.Items[ ___forsortKey1 ].(FrontInterface_ModuleInfoDownCast).ToFrontInterface_ModuleInfo()
                index = index + 1
                
                importModuleType2Index.Set(moduleInfo.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env),index)
                importProcessInfo2Index.Set(moduleInfo.FP.Get_exportInfo(_env).FP.Get_processInfo(_env),index)
                importProcessInfo2Index.Set(moduleInfo.FP.Get_exportInfo(_env).FP.Get_processInfo(_env).FP.Get_orgInfo(_env),index)
            }
        }
    }
    var typeId2TypeInfo *LnsMap
    typeId2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    {
        var work *Ast_TypeInfo
        work = node.FP.Get_moduleTypeInfo(_env)
        for Ast_TypeInfo_hasParent(_env, work) {
            typeId2TypeInfo.Set(work.FP.Get_typeId(_env),work)
            work = work.FP.Get_parentInfo(_env)
            
        }
    }
    var pickupClassMap *LnsMap
    pickupClassMap = NewLnsMap( map[LnsAny]LnsAny{})
    var checkExportTypeInfo func(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny
    checkExportTypeInfo = func(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny {
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
    var isDependOnExt func(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool
    isDependOnExt = func(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
        if Ast_isExtId(_env, typeInfo){
            return true
        }
        return self.moduleTypeInfo.FP.Get_processInfo(_env) != typeInfo.FP.Get_processInfo(_env)
    }
    var pickupTypeId func(_env *LnsEnv, typeInfo *Ast_TypeInfo,forceFlag LnsAny,pickupChildFlag LnsAny)
    pickupTypeId = func(_env *LnsEnv, typeInfo *Ast_TypeInfo,forceFlag LnsAny,pickupChildFlag LnsAny) {
        if typeInfo.FP.Get_typeId(_env) == Ast_rootTypeIdInfo{
            return 
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(forceFlag)) &&
            _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, typeInfo.FP.Get_accessMode(_env)))) ).(bool)){
            return 
        }
        if Lns_isCondTrue( typeId2TypeInfo.Get(typeInfo.FP.Get_typeId(_env))){
            if isDependOnExt(_env, typeInfo){
                return 
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( pickupChildFlag) &&
                _env.SetStackVal( Lns_op_not(typeInfo.FP.Get_nilable(_env))) )){
                for _, _itemTypeInfo := range( typeInfo.FP.Get_children(_env).Items ) {
                    itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if Ast_isPubToExternal(_env, itemTypeInfo.FP.Get_accessMode(_env)){
                        if _switch1 := itemTypeInfo.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Class || _switch1 == Ast_TypeInfoKind__IF || _switch1 == Ast_TypeInfoKind__Form || _switch1 == Ast_TypeInfoKind__FormFunc || _switch1 == Ast_TypeInfoKind__Func || _switch1 == Ast_TypeInfoKind__Method {
                            pickupTypeId(_env, itemTypeInfo, true, true)
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
            pickupTypeId(_env, typeInfo.FP.Get_extedType(_env), true, true)
        }
        if typeInfo != typeInfo.FP.Get_srcTypeInfo(_env){
            pickupTypeId(_env, typeInfo.FP.Get_srcTypeInfo(_env), true, false)
        } else if typeInfo.FP.Get_nilable(_env){
            pickupTypeId(_env, typeInfo.FP.Get_nonnilableType(_env), true, false)
        } else { 
            if isDependOnExt(_env, typeInfo){
                return 
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ).(bool){
                pickupClassMap.Set(typeInfo.FP.Get_typeId(_env),typeInfo)
            }
            if Lns_op_not(typeInfo.FP.Get_externalFlag(_env)){
                if _switch2 := typeInfo.FP.Get_kind(_env); _switch2 == Ast_TypeInfoKind__IF || _switch2 == Ast_TypeInfoKind__Class || _switch2 == Ast_TypeInfoKind__Form || _switch2 == Ast_TypeInfoKind__FormFunc || _switch2 == Ast_TypeInfoKind__Alge || _switch2 == Ast_TypeInfoKind__Enum || _switch2 == Ast_TypeInfoKind__Map || _switch2 == Ast_TypeInfoKind__Set || _switch2 == Ast_TypeInfoKind__List || _switch2 == Ast_TypeInfoKind__Array || _switch2 == Ast_TypeInfoKind__Alternate || _switch2 == Ast_TypeInfoKind__Box {
                    pickupTypeId(_env, typeInfo.FP.Get_nilableTypeInfo(_env), true, false)
                }
            }
            var parentInfo *Ast_TypeInfo
            parentInfo = typeInfo.FP.Get_parentInfo(_env)
            pickupTypeId(_env, parentInfo, true, false)
            pickupTypeId(_env, typeInfo.FP.Get_genSrcTypeInfo(_env), true, false)
            var baseInfo *Ast_TypeInfo
            baseInfo = typeInfo.FP.Get_baseTypeInfo(_env)
            if baseInfo.FP.Get_typeId(_env) != Ast_rootTypeIdInfo{
                pickupTypeId(_env, baseInfo, true, true)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_interfaceList(_env).Items ) {
                itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                pickupTypeId(_env, itemTypeInfo, true, true)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_itemTypeInfoList(_env).Items ) {
                itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                pickupTypeId(_env, itemTypeInfo, true, false)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_argTypeInfoList(_env).Items ) {
                itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                pickupTypeId(_env, itemTypeInfo, true, false)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_retTypeInfoList(_env).Items ) {
                itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                pickupTypeId(_env, itemTypeInfo, true, true)
            }
            if Lns_isCondTrue( pickupChildFlag){
                for _, _itemTypeInfo := range( typeInfo.FP.Get_children(_env).Items ) {
                    itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if itemTypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
                        if _switch3 := itemTypeInfo.FP.Get_kind(_env); _switch3 == Ast_TypeInfoKind__Class || _switch3 == Ast_TypeInfoKind__IF || _switch3 == Ast_TypeInfoKind__Form || _switch3 == Ast_TypeInfoKind__FormFunc || _switch3 == Ast_TypeInfoKind__Func || _switch3 == Ast_TypeInfoKind__Method {
                            pickupTypeId(_env, itemTypeInfo, true, true)
                        }
                    }
                }
            }
        }
    }
    var validChildrenSet *LnsMap
    validChildrenSet = NewLnsMap( map[LnsAny]LnsAny{})
    {
        var typeInfo *Ast_TypeInfo
        typeInfo = self.moduleTypeInfo
        for Ast_TypeInfo_hasParent(_env, typeInfo) {
            validChildrenSet.Set(typeInfo.FP.Get_parentInfo(_env),NewLnsMap( map[LnsAny]LnsAny{typeInfo.FP.Get_typeId(_env):typeInfo,}))
            typeInfo = typeInfo.FP.Get_parentInfo(_env)
            
        }
        pickupTypeId(_env, self.moduleTypeInfo, true, nil)
    }
    var typeId2ClassMap *LnsMap
    typeId2ClassMap = node.FP.Get_typeId2ClassMap(_env)
    for _, _namespaceInfo := range( typeId2ClassMap.Items ) {
        namespaceInfo := _namespaceInfo.(Nodes_NamespaceInfoDownCast).ToNodes_NamespaceInfo()
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( namespaceInfo.TypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) &&
            _env.SetStackVal( Lns_op_not(namespaceInfo.TypeInfo.FP.Get_externalFlag(_env))) ).(bool)){
            pickupClassMap.Set(namespaceInfo.TypeInfo.FP.Get_typeId(_env),namespaceInfo.TypeInfo)
        }
    }
    var classId2TypeInfo *LnsMap
    classId2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    for _idInfo, _classTypeInfo := range( self.classId2TypeInfo.Items ) {
        idInfo := _idInfo.(Ast_IdInfoDownCast).ToAst_IdInfo()
        classTypeInfo := _classTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        classId2TypeInfo.Set(idInfo.Id,classTypeInfo)
    }
    var serializeInfo *Ast_SerializeInfo
    serializeInfo = NewAst_SerializeInfo(_env, importProcessInfo2Index, NewLnsMap( map[LnsAny]LnsAny{}))
    self.FP.writeln(_env, "local __typeId2ClassInfoMap = {}")
    self.FP.writeln(_env, "_moduleObj.__typeId2ClassInfoMap = __typeId2ClassInfoMap")
    {
        __forsortCollection2 := classId2TypeInfo
        __forsortSorted2 := __forsortCollection2.CreateKeyListInt()
        __forsortSorted2.Sort( _env, LnsItemKindInt, nil )
        for _, _classTypeId := range( __forsortSorted2.Items ) {
            classTypeInfo := __forsortCollection2.Items[ _classTypeId ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            classTypeId := _classTypeId.(LnsInt)
            if classTypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
                pickupTypeId(_env, classTypeInfo, true, _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( validChildrenSet.Get(classTypeInfo) == nil) &&
                    _env.SetStackVal( Lns_op_not(classTypeInfo.FP.Get_externalFlag(_env))) ).(bool))
                pickupClassMap.Set(classTypeInfo.FP.Get_typeId(_env),nil)
                self.FP.writeln(_env, "do")
                self.FP.pushIndent(_env, nil)
                self.FP.writeln(_env, _env.LuaVM.String_format("local __classInfo%d = {}", []LnsAny{classTypeId}))
                self.FP.writeln(_env, _env.LuaVM.String_format("__typeId2ClassInfoMap[ %d ] = __classInfo%d", []LnsAny{classTypeId, classTypeId}))
                for _, _memberNode := range( Lns_unwrap( self.classId2MemberList.Get(classTypeInfo.FP.Get_typeId(_env))).(*LnsList).Items ) {
                    memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
                    if Ast_isPubToExternal(_env, memberNode.FP.Get_accessMode(_env)){
                        var memberName string
                        memberName = memberNode.FP.Get_name(_env).Txt
                        var memberTypeInfo *Ast_TypeInfo
                        memberTypeInfo = memberNode.FP.Get_expType(_env)
                        self.FP.writeln(_env, _env.LuaVM.String_format("__classInfo%d.%s = {", []LnsAny{classTypeId, memberName}))
                        self.FP.writeln(_env, _env.LuaVM.String_format("  name='%s', staticFlag = %s, mutMode = %d,", []LnsAny{memberName, memberNode.FP.Get_staticFlag(_env), memberNode.FP.Get_symbolInfo(_env).FP.Get_mutMode(_env)}) + _env.LuaVM.String_format("accessMode = %d, typeId = %s }", []LnsAny{memberNode.FP.Get_accessMode(_env), serializeInfo.FP.SerializeId(_env, memberTypeInfo.FP.Get_typeId(_env))}))
                        pickupTypeId(_env, memberTypeInfo, true, nil)
                    }
                }
                self.FP.popIndent(_env)
                self.FP.writeln(_env, "end")
            }
        }
    }
    var pickupedClassMap *LnsMap
    pickupedClassMap = NewLnsMap( map[LnsAny]LnsAny{})
    for  {
        var workClassMap *LnsMap
        workClassMap = NewLnsMap( map[LnsAny]LnsAny{})
        var hasWorkClassFlag bool
        hasWorkClassFlag = false
        for _classTypeId, _classTypeInfo := range( pickupClassMap.Items ) {
            classTypeId := _classTypeId.(Ast_IdInfoDownCast).ToAst_IdInfo()
            classTypeInfo := _classTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(pickupedClassMap.Get(classTypeId)){
                pickupedClassMap.Set(classTypeId,classTypeInfo)
                if Lns_isCondTrue( checkExportTypeInfo(_env, classTypeInfo)){
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
                classTypeInfo := __forsortCollection3.Items[ _classTypeId ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                classTypeId := _classTypeId.(LnsInt)
                if Lns_op_not(Ast_isBuiltin(_env, classTypeId)){
                    var scope *Ast_Scope
                    
                    {
                        _scope := classTypeInfo.FP.Get_scope(_env)
                        if _scope == nil{
                            Util_err(_env, _env.LuaVM.String_format("%s.scope is nil", []LnsAny{classTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
                        } else {
                            scope = _scope.(*Ast_Scope)
                        }
                    }
                    pickupTypeId(_env, classTypeInfo, true, _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( validChildrenSet.Get(classTypeInfo) == nil) &&
                        _env.SetStackVal( Lns_op_not(classTypeInfo.FP.Get_externalFlag(_env))) ).(bool))
                    if Lns_isCondTrue( checkExportTypeInfo(_env, classTypeInfo)){
                        self.FP.writeln(_env, "do")
                        self.FP.pushIndent(_env, nil)
                        self.FP.writeln(_env, _env.LuaVM.String_format("local __classInfo%d = {}", []LnsAny{classTypeId}))
                        self.FP.writeln(_env, _env.LuaVM.String_format("__typeId2ClassInfoMap[ %d ] = __classInfo%d", []LnsAny{classTypeId, classTypeId}))
                        {
                            __forsortCollection4 := scope.FP.Get_symbol2SymbolInfoMap(_env)
                            __forsortSorted4 := __forsortCollection4.CreateKeyListStr()
                            __forsortSorted4.Sort( _env, LnsItemKindStr, nil )
                            for _, _fieldName := range( __forsortSorted4.Items ) {
                                symbolInfo := __forsortCollection4.Items[ _fieldName ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                                fieldName := _fieldName.(string)
                                var typeInfo *Ast_TypeInfo
                                typeInfo = symbolInfo.FP.Get_typeInfo(_env)
                                if _env.PopVal( _env.IncStack() ||
                                    _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) ||
                                    _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Var) ).(bool){
                                    if symbolInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
                                        self.FP.writeln(_env, _env.LuaVM.String_format("__classInfo%d.%s = {", []LnsAny{classTypeId, fieldName}))
                                        self.FP.writeln(_env, _env.LuaVM.String_format("  name='%s', staticFlag = %s, ", []LnsAny{fieldName, symbolInfo.FP.Get_staticFlag(_env)}) + _env.LuaVM.String_format("accessMode = %d, typeId = %d }", []LnsAny{symbolInfo.FP.Get_accessMode(_env), typeInfo.FP.Get_typeId(_env).Id}))
                                        pickupTypeId(_env, typeInfo, nil, nil)
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
            pickupTypeId(_env, macroTypeInfo, true, nil)
            self.FP.writeln(_env, "do")
            self.FP.pushIndent(_env, nil)
            self.FP.writeln(_env, "local info = {}")
            self.FP.writeln(_env, _env.LuaVM.String_format("__macroName2InfoMap[ %d ] = info", []LnsAny{macroTypeInfo.FP.Get_typeId(_env).Id}))
            var pos *Types_Position
            pos = macroDeclNode.FP.Get_pos(_env).FP.Get_orgPos(_env)
            self.FP.writeln(_env, _env.LuaVM.String_format("info.pos = {%d,%d}", []LnsAny{pos.LineNo, pos.Column}))
            self.FP.writeln(_env, _env.LuaVM.String_format("info.name = %q", []LnsAny{declInfo.FP.Get_name(_env).Txt}))
            self.FP.Write(_env, "info.argList = {")
            for _index, _argNode := range( declInfo.FP.Get_argList(_env).Items ) {
                index := _index + 1
                argNode := _argNode.(Nodes_DeclArgNodeDownCast).ToNodes_DeclArgNode()
                if index != 1{
                    self.FP.Write(_env, ",")
                }
                self.FP.Write(_env, _env.LuaVM.String_format("{name=%q,typeId=%d}", []LnsAny{argNode.FP.Get_name(_env).Txt, argNode.FP.Get_expType(_env).FP.Get_typeId(_env).Id}))
            }
            self.FP.writeln(_env, "}")
            self.FP.Write(_env, "info.symList = {")
            var firstFlag bool
            firstFlag = true
            {
                __forsortCollection5 := macroInfo.FP.Get_symbol2MacroValInfoMap(_env)
                __forsortSorted5 := __forsortCollection5.CreateKeyListStr()
                __forsortSorted5.Sort( _env, LnsItemKindStr, nil )
                for _, _name := range( __forsortSorted5.Items ) {
                    symInfo := __forsortCollection5.Items[ _name ].(Nodes_MacroValInfoDownCast).ToNodes_MacroValInfo()
                    name := _name.(string)
                    if firstFlag{
                        firstFlag = false
                        
                    } else { 
                        self.FP.Write(_env, ",")
                    }
                    self.FP.Write(_env, _env.LuaVM.String_format("{name=%q,typeId=%d}", []LnsAny{name, symInfo.TypeInfo.FP.Get_typeId(_env).Id}))
                    pickupTypeId(_env, symInfo.TypeInfo, true, nil)
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
                    workFilter = NewconvLua_ConvFilter(_env, declInfo.FP.Get_name(_env).Txt, memStream.FP, NewUtil_NullOStream(_env).FP, ConvLua_ConvMode__Convert, false, Ast_headTypeInfo, self.processInfo, Ast_SymbolKind__Typ, self.builtinFunc, self.useLuneRuntime, self.targetLuaVer, self.enableTest, self.useIpairs)
                    workFilter.macroDepth = workFilter.macroDepth + 1
                    
                    workFilter.FP.ProcessBlock(_env, stmtBlock, ConvLua_Opt2Stem(NewConvLua_Opt(_env, &node.Nodes_Node)))
                    workFilter.macroDepth = workFilter.macroDepth - 1
                    
                    memStream.FP.Close(_env)
                    self.FP.writeln(_env, _env.LuaVM.String_format("info.stmtBlock = %s", []LnsAny{Parser_quoteStr(_env, memStream.FP.Get_txt(_env))}))
                }
            }
            self.FP.writeln(_env, "info.tokenList = {")
            var prevLineNo LnsInt
            prevLineNo = -1
            for _index, _token := range( declInfo.FP.Get_tokenList(_env).Items ) {
                index := _index + 1
                token := _token.(Types_TokenDownCast).ToTypes_Token()
                if index > 1{
                    self.FP.Write(_env, ",")
                }
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( prevLineNo != -1) &&
                    _env.SetStackVal( prevLineNo != token.Pos.LineNo) ).(bool)){
                    self.FP.Write(_env, _env.LuaVM.String_format("{%d,%s},", []LnsAny{Types_TokenKind__Dlmt, Parser_quoteStr(_env, "\n")}))
                }
                prevLineNo = token.Pos.LineNo
                
                self.FP.Write(_env, _env.LuaVM.String_format("{%d,%s}", []LnsAny{token.Kind, Parser_quoteStr(_env, token.Txt)}))
            }
            self.FP.writeln(_env, "}")
            self.FP.popIndent(_env)
            self.FP.writeln(_env, "end")
        }
    }
    self.FP.writeln(_env, "local __varName2InfoMap = {}")
    self.FP.writeln(_env, "_moduleObj.__varName2InfoMap = __varName2InfoMap")
    {
        __forsortCollection6 := self.pubVarName2InfoMap
        __forsortSorted6 := __forsortCollection6.CreateKeyListStr()
        __forsortSorted6.Sort( _env, LnsItemKindStr, nil )
        for _, _varName := range( __forsortSorted6.Items ) {
            varInfo := __forsortCollection6.Items[ _varName ].(convLua_PubVerInfoDownCast).ToconvLua_PubVerInfo()
            varName := _varName.(string)
            self.FP.writeln(_env, _env.LuaVM.String_format("__varName2InfoMap.%s = {", []LnsAny{varName}))
            self.FP.writeln(_env, _env.LuaVM.String_format("  name='%s', accessMode = %d, typeId = %s, mutable = %s }", []LnsAny{varName, varInfo.AccessMode, serializeInfo.FP.SerializeId(_env, varInfo.TypeInfo.FP.Get_typeId(_env)), true}))
            pickupTypeId(_env, varInfo.TypeInfo, true, nil)
        }
    }
    {
        __forsortCollection7 := self.pubFuncName2InfoMap
        __forsortSorted7 := __forsortCollection7.CreateKeyListStr()
        __forsortSorted7.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey7 := range( __forsortSorted7.Items ) {
            funcInfo := __forsortCollection7.Items[ ___forsortKey7 ].(convLua_PubFuncInfoDownCast).ToconvLua_PubFuncInfo()
            pickupTypeId(_env, funcInfo.TypeInfo, true, nil)
        }
    }
    for _, _aliasNode := range( node.FP.Get_nodeManager(_env).FP.GetAliasNodeList(_env).Items ) {
        aliasNode := _aliasNode.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
        pickupTypeId(_env, aliasNode.FP.Get_expType(_env), false, nil)
    }
    self.FP.writeln(_env, "local __typeInfoList = {}")
    self.FP.writeln(_env, "_moduleObj.__typeInfoList = __typeInfoList")
    var listIndex LnsInt
    listIndex = 1
    var outputDepend func(_env *LnsEnv, typeInfo *Ast_TypeInfo,moduleTypeInfo *Ast_TypeInfo) LnsInt
    outputDepend = func(_env *LnsEnv, typeInfo *Ast_TypeInfo,moduleTypeInfo *Ast_TypeInfo) LnsInt {
        var typeId *Ast_IdInfo
        typeId = typeInfo.FP.Get_typeId(_env)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.processInfo.FP.Get_orgInfo(_env) == typeId.FP.Get_processInfo(_env)) ||
            _env.SetStackVal( Lns_op_not(Ast_TypeInfo_hasParent(_env, moduleTypeInfo))) ).(bool){
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
    var wroteTypeIdSet *LnsSet
    wroteTypeIdSet = NewLnsSet([]LnsAny{})
    var outputTypeInfo func(_env *LnsEnv, typeInfo *Ast_TypeInfo)
    outputTypeInfo = func(_env *LnsEnv, typeInfo *Ast_TypeInfo) {
        if Ast_isBuiltin(_env, typeInfo.FP.Get_typeId(_env).Id){
            return 
        }
        if Lns_op_not(Ast_TypeInfo_hasParent(_env, typeInfo)){
            return 
        }
        if _switch1 := typeInfo.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Class || _switch1 == Ast_TypeInfoKind__IF {
            if _switch2 := typeInfo.FP.Get_accessMode(_env); _switch2 == Ast_AccessMode__Pub || _switch2 == Ast_AccessMode__Pro || _switch2 == Ast_AccessMode__Global {
            } else {
                Util_errorLog(_env, _env.LuaVM.String_format("skip: %s %s", []LnsAny{typeInfo.FP.Get_accessMode(_env), self.FP.getFullName(_env, typeInfo)}))
                return 
            }
        }
        var typeId *Ast_IdInfo
        typeId = typeInfo.FP.Get_typeId(_env)
        if wroteTypeIdSet.Has(Ast_IdInfo2Stem(typeId)){
            return 
        }
        wroteTypeIdSet.Add(Ast_IdInfo2Stem(typeId))
        self.FP.Write(_env, _env.LuaVM.String_format("__typeInfoList[%d] = ", []LnsAny{listIndex}))
        listIndex = listIndex + 1
        
        var validChildren *LnsMap
        
        {
            _validChildren := validChildrenSet.Get(typeInfo)
            if _validChildren == nil{
                validChildren = typeId2TypeInfo
                
            } else {
                validChildren = _validChildren.(*LnsMap)
            }
        }
        typeInfo.FP.Serialize(_env, self.FP, NewAst_SerializeInfo(_env, importProcessInfo2Index, validChildren))
    }
    for _typeId, _typeInfo := range( self.pubEnumId2EnumTypeInfo.Items ) {
        typeId := _typeId.(Ast_IdInfoDownCast).ToAst_IdInfo()
        typeInfo := _typeInfo.(Ast_EnumTypeInfoDownCast).ToAst_EnumTypeInfo()
        typeId2TypeInfo.Set(typeId,&typeInfo.Ast_TypeInfo)
    }
    for _typeId, _typeInfo := range( self.pubAlgeId2AlgeTypeInfo.Items ) {
        typeId := _typeId.(Ast_IdInfoDownCast).ToAst_IdInfo()
        typeInfo := _typeInfo.(Ast_AlgeTypeInfoDownCast).ToAst_AlgeTypeInfo()
        typeId2TypeInfo.Set(typeId,&typeInfo.Ast_TypeInfo)
        {
            __forsortCollection8 := typeInfo.FP.Get_valInfoMap(_env)
            __forsortSorted8 := __forsortCollection8.CreateKeyListStr()
            __forsortSorted8.Sort( _env, LnsItemKindStr, nil )
            for _, ___forsortKey8 := range( __forsortSorted8.Items ) {
                valInfo := __forsortCollection8.Items[ ___forsortKey8 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
                for _, _valType := range( valInfo.FP.Get_typeList(_env).Items ) {
                    valType := _valType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    pickupTypeId(_env, valType, true, nil)
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
            typeId2TypeInfo.Set(aliasType.FP.Get_typeId(_env),&aliasType.Ast_TypeInfo)
            typeId2TypeInfo.Set(aliasSrcType.FP.Get_typeId(_env),aliasSrcType)
        }
    }
    self.FP.writeln(_env, "local __dependIdMap = {}")
    self.FP.writeln(_env, "_moduleObj.__dependIdMap = __dependIdMap")
    var exportNeedModuleTypeInfo *LnsSet
    exportNeedModuleTypeInfo = NewLnsSet([]LnsAny{})
    {
        var module2TypeList *LnsMap
        module2TypeList = NewLnsMap( map[LnsAny]LnsAny{})
        for _, _typeInfo := range( typeId2TypeInfo.Items ) {
            typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var modIndex LnsInt
            modIndex = Lns_unwrap( importProcessInfo2Index.Get(typeInfo.FP.Get_typeId(_env).FP.Get_processInfo(_env))).(LnsInt)
            var _map *LnsMap
            
            {
                __map := module2TypeList.Get(modIndex)
                if __map == nil{
                    _map = NewLnsMap( map[LnsAny]LnsAny{})
                    
                    module2TypeList.Set(modIndex,_map)
                } else {
                    _map = __map.(*LnsMap)
                }
            }
            _map.Set(typeInfo.FP.Get_typeId(_env).Id,typeInfo)
        }
        {
            __forsortCollection9 := module2TypeList
            __forsortSorted9 := __forsortCollection9.CreateKeyListInt()
            __forsortSorted9.Sort( _env, LnsItemKindInt, nil )
            for _, ___forsortKey9 := range( __forsortSorted9.Items ) {
                _map := __forsortCollection9.Items[ ___forsortKey9 ].(*LnsMap)
                {
                    __forsortCollection10 := _map
                    __forsortSorted10 := __forsortCollection10.CreateKeyListInt()
                    __forsortSorted10.Sort( _env, LnsItemKindInt, nil )
                    for _, ___forsortKey10 := range( __forsortSorted10.Items ) {
                        typeInfo := __forsortCollection10.Items[ ___forsortKey10 ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        var moduleTypeInfo *Ast_TypeInfo
                        moduleTypeInfo = typeInfo.FP.GetModule(_env)
                        exportNeedModuleTypeInfo.Add(Ast_TypeInfo2Stem(moduleTypeInfo))
                        var ret LnsInt
                        ret = outputDepend(_env, typeInfo, moduleTypeInfo)
                        if ret == convLua_ExportIdKind__Normal{
                            outputTypeInfo(_env, typeInfo)
                        }
                    }
                }
            }
        }
    }
    for _moduleTypeInfo := range( node.FP.Get_useModuleMacroSet(_env).Items ) {
        moduleTypeInfo := _moduleTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        exportNeedModuleTypeInfo.Add(Ast_TypeInfo2Stem(moduleTypeInfo))
    }
    self.FP.writeln(_env, "local __dependModuleMap = {}")
    self.FP.writeln(_env, "_moduleObj.__dependModuleMap = __dependModuleMap")
    {
        __forsortCollection11 := importNameMap
        __forsortSorted11 := __forsortCollection11.CreateKeyListStr()
        __forsortSorted11.Sort( _env, LnsItemKindStr, nil )
        for _, _name := range( __forsortSorted11.Items ) {
            moduleInfo := __forsortCollection11.Items[ _name ].(FrontInterface_ModuleInfoDownCast).ToFrontInterface_ModuleInfo()
            name := _name.(string)
            var moduleTypeInfo *Ast_TypeInfo
            moduleTypeInfo = moduleInfo.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env)
            self.FP.writeln(_env, _env.LuaVM.String_format("__dependModuleMap[ '%s' ] = { typeId = %d, use = %s, buildId = %q }", []LnsAny{name, Lns_unwrap( importModuleType2Index.Get(moduleTypeInfo)).(LnsInt), exportNeedModuleTypeInfo.Has(Ast_TypeInfo2Stem(moduleTypeInfo)), (Lns_unwrap( node.FP.Get_importModule2moduleInfo(_env).Get(moduleTypeInfo)).(*FrontInterface_ModuleInfo)).FP.Get_moduleId(_env).FP.Get_idStr(_env)}))
        }
    }
    self.FP.Write(_env, "_moduleObj.__subModuleMap = {")
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
                        self.FP.Write(_env, ",")
                    }
                    self.FP.Write(_env, _env.LuaVM.String_format("%q", []LnsAny{usePath}))
                }
            }
        }
    }
    self.FP.writeln(_env, "}")
    self.FP.Write(_env, "_moduleObj.__moduleHierarchy = {")
    var workType *Ast_TypeInfo
    workType = node.FP.Get_moduleTypeInfo(_env)
    for Ast_TypeInfo_hasParent(_env, workType) {
        self.FP.Write(_env, _env.LuaVM.String_format("%d,", []LnsAny{workType.FP.Get_typeId(_env).Id}))
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
    self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.__moduleTypeId = %d", []LnsAny{moduleTypeInfo.FP.Get_typeId(_env).Id}))
    self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.__moduleSymbolKind = %d", []LnsAny{moduleSymbolKind}))
    self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.__moduleMutable = %s", []LnsAny{Ast_TypeInfo_isMut(_env, moduleTypeInfo)}))
    self.FP.writeln(_env, "----- meta -----")
    self.FP.writeln(_env, "return _moduleObj")
    self.outMetaFlag = false
    
}

// 990: decl @lune.@base.@convLua.ConvFilter.processRoot
func (self *convLua_ConvFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    self.FP.writeln(_env, _env.LuaVM.String_format("--%s", []LnsAny{self.streamName}))
    self.needModuleObj = Nodes_ProvideNode2Stem(node.FP.Get_provideNode(_env)) == nil
    
    if self.needModuleObj{
        self.FP.writeln(_env, "local _moduleObj = {}")
    }
    if self.enableTest{
        self.FP.writeln(_env, "_moduleObj.__enableTest = true")
    }
    self.FP.writeln(_env, _env.LuaVM.String_format("local __mod__ = '%s'", []LnsAny{node.FP.Get_moduleTypeInfo(_env).FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), nil)}))
    var luneSymbol string
    luneSymbol = _env.LuaVM.String_format("_lune%d", []LnsAny{Ver_luaModVersion})
    {
        _runtime := self.useLuneRuntime
        if !Lns_IsNil( _runtime ) {
            runtime := _runtime.(string)
            self.FP.writeln(_env, _env.LuaVM.String_format("local _lune = require( \"%s\" )", []LnsAny{runtime}))
        } else {
            self.FP.writeln(_env, "local _lune = {}")
            self.FP.writeln(_env, _env.LuaVM.String_format("if %s then\n   _lune = %s\nend", []LnsAny{luneSymbol, luneSymbol}))
            if node.FP.Get_luneHelperInfo(_env).UseAlge{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Alge))
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__AlgeMapping))
            }
            if node.FP.Get_luneHelperInfo(_env).UseSet{
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__SetOp))
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__SetMapping))
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( node.FP.Get_luneHelperInfo(_env).UseUnpack) &&
                _env.SetStackVal( Lns_op_not(self.targetLuaVer.FP.Get_hasTableUnpack(_env))) ).(bool)){
                self.FP.writeln(_env, LuaMod_getCode(_env, LuaMod_CodeKind__Unpack))
            }
            if node.FP.Get_luneHelperInfo(_env).UseLoad{
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
        }
    }
    self.FP.writeln(_env, _env.LuaVM.String_format("if not %s then\n   %s = _lune\nend", []LnsAny{luneSymbol, luneSymbol}))
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
    var children *LnsList
    children = node.FP.Get_children(_env)
    for _, _child := range( children.Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        convLua_filter_7_(_env, child, self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
    }
    {
        __exp := node.FP.Get_provideNode(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ProvideNode)
            self.FP.Write(_env, "return ")
            self.FP.Write(_env, _exp.FP.Get_symbol(_env).FP.Get_name(_env))
            self.FP.writeln(_env, "")
        } else {
            self.FP.writeln(_env, "return _moduleObj")
        }
    }
    self.FP.outputMeta(_env, node)
}

// 1100: decl @lune.@base.@convLua.ConvFilter.processSubfile
func (self *convLua_ConvFilter) ProcessSubfile(_env *LnsEnv, node *Nodes_SubfileNode,_opt LnsAny) {
}

// 1105: decl @lune.@base.@convLua.ConvFilter.processAsyncLock
func (self *convLua_ConvFilter) ProcessAsyncLock(_env *LnsEnv, node *Nodes_AsyncLockNode,_opt LnsAny) {
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
}

// 1112: decl @lune.@base.@convLua.ConvFilter.processBlockSub
func (self *convLua_ConvFilter) ProcessBlockSub(_env *LnsEnv, node *Nodes_BlockNode,_opt LnsAny) {
    var word string
    word = ""
    if _switch1 := node.FP.Get_blockKind(_env); _switch1 == Nodes_BlockKind__If || _switch1 == Nodes_BlockKind__Elseif {
        word = "then"
        
    } else if _switch1 == Nodes_BlockKind__Else {
        word = ""
        
    } else if _switch1 == Nodes_BlockKind__While {
        word = "do"
        
    } else if _switch1 == Nodes_BlockKind__Repeat {
        word = ""
        
    } else if _switch1 == Nodes_BlockKind__For {
        word = "do"
        
    } else if _switch1 == Nodes_BlockKind__Apply {
        word = "do"
        
    } else if _switch1 == Nodes_BlockKind__Foreach {
        word = "do"
        
    } else if _switch1 == Nodes_BlockKind__Macro {
        word = ""
        
    } else if _switch1 == Nodes_BlockKind__Func {
        word = ""
        
    } else if _switch1 == Nodes_BlockKind__Default {
        word = ""
        
    } else if _switch1 == Nodes_BlockKind__Block || _switch1 == Nodes_BlockKind__AsyncLock {
        word = "do"
        
    } else if _switch1 == Nodes_BlockKind__LetUnwrap {
        word = ""
        
    } else if _switch1 == Nodes_BlockKind__LetUnwrapThenDo {
        word = ""
        
    } else if _switch1 == Nodes_BlockKind__IfUnwrap {
        word = ""
        
    }
    self.FP.writeln(_env, word)
    self.FP.pushIndent(_env, nil)
    if Lns_isCondTrue( _env.NilAccFin(_env.NilAccPush(node.FP.Get_scope(_env).FP.GetSymbolInfoChild(_env, "_")) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_SymbolInfo).FP.Get_posForLatestMod(_env)}))){
        self.FP.writeln(_env, "local _")
    }
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList(_env)
    for _, _statement := range( stmtList.Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        convLua_filter_7_(_env, statement, self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
    }
    self.FP.popIndent(_env)
    if _switch2 := node.FP.Get_blockKind(_env); _switch2 == Nodes_BlockKind__Block || _switch2 == Nodes_BlockKind__AsyncLock {
        self.FP.writeln(_env, "end")
    }
}

// 1182: decl @lune.@base.@convLua.ConvFilter.processLoadRuntime
func (self *convLua_ConvFilter) processLoadRuntime(_env *LnsEnv) {
    {
        __exp := self.useLuneRuntime
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(string)
            self.FP.writeln(_env, _env.LuaVM.String_format("local _lune = require( \"%s\" )", []LnsAny{_exp}))
        } else {
            self.FP.writeln(_env, _env.LuaVM.String_format("local _lune = require( \"%s\" )", []LnsAny{Option_getRuntimeModule(_env)}))
        }
    }
}

// 1192: decl @lune.@base.@convLua.ConvFilter.processScope
func (self *convLua_ConvFilter) ProcessScope(_env *LnsEnv, node *Nodes_ScopeNode,_opt LnsAny) {
    if node.FP.Get_scopeKind(_env) == Nodes_ScopeKind__Root{
        self.FP.processLoadRuntime(_env)
    }
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
}

// 1202: decl @lune.@base.@convLua.ConvFilter.processStmtExp
func (self *convLua_ConvFilter) ProcessStmtExp(_env *LnsEnv, node *Nodes_StmtExpNode,_opt LnsAny) {
    convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
}

// 1208: decl @lune.@base.@convLua.ConvFilter.processDeclEnum
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
        enumFullName = _env.LuaVM.String_format("%s.%s", []LnsAny{self.FP.getFullName(_env, parentInfo), enumFullName})
        
        access = ""
        
        isTopNS = false
        
    }
    self.FP.writeln(_env, _env.LuaVM.String_format("%s%s = {}", []LnsAny{access, enumFullName}))
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( isTopNS) &&
        _env.SetStackVal( node.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) ).(bool)){
        if self.needModuleObj{
            self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.%s = %s", []LnsAny{enumFullName, enumFullName}))
        }
    }
    if typeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
        self.pubEnumId2EnumTypeInfo.Set(typeInfo.FP.Get_typeId(_env),typeInfo)
    }
    self.FP.writeln(_env, _env.LuaVM.String_format("%s._val2NameMap = {}", []LnsAny{enumFullName}))
    self.FP.writeln(_env, _env.LuaVM.String_format("function %s:_getTxt( val )\n   local name = self._val2NameMap[ val ]\n   if name then\n      return string.format( \"%s.%%s\", name )\n   end\n   return string.format( \"illegal val -- %%s\", val )\nend\nfunction %s._from( val )\n   if %s._val2NameMap[ val ] then\n      return val\n   end\n   return nil\nend\n    ", []LnsAny{enumFullName, enumFullName, enumFullName, enumFullName}))
    self.FP.writeln(_env, _env.LuaVM.String_format("%s.__allList = {}\nfunction %s.get__allList()\n   return %s.__allList\nend\n", []LnsAny{enumFullName, enumFullName, enumFullName}))
    for _index, _valName := range( node.FP.Get_valueNameList(_env).Items ) {
        index := _index + 1
        valName := _valName.(Types_TokenDownCast).ToTypes_Token()
        var valInfo *Ast_EnumValInfo
        valInfo = Lns_unwrap( typeInfo.FP.GetEnumValInfo(_env, valName.Txt)).(*Ast_EnumValInfo)
        var valTxt string
        valTxt = _env.LuaVM.String_format("%s", []LnsAny{Ast_getEnumLiteralVal(_env, valInfo.FP.Get_val(_env))})
        if typeInfo.FP.Get_valTypeInfo(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeString, nil, nil){
            valTxt = _env.LuaVM.String_format("'%s'", []LnsAny{Ast_getEnumLiteralVal(_env, valInfo.FP.Get_val(_env))})
            
        }
        self.FP.writeln(_env, _env.LuaVM.String_format("%s.%s = %s", []LnsAny{enumFullName, valName.Txt, valTxt}))
        self.FP.writeln(_env, _env.LuaVM.String_format("%s._val2NameMap[%s] = '%s'", []LnsAny{enumFullName, valTxt, valName.Txt}))
        self.FP.writeln(_env, _env.LuaVM.String_format("%s.__allList[%d] = %s.%s", []LnsAny{enumFullName, index, enumFullName, valName.Txt}))
    }
}

// 1290: decl @lune.@base.@convLua.ConvFilter.getMapInfo
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
            if Ast_NormalTypeInfo_isAvailableMapping(_env, self.processInfo, nonnilableType, NewLnsMap( map[LnsAny]LnsAny{})){
                funcTxt = _env.LuaVM.String_format("%s._fromMap", []LnsAny{self.FP.getFullName(_env, nonnilableType)})
                
                if convLua_isGenericType_10_(_env, nonnilableType){
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
        funcTxt = _env.LuaVM.String_format("%s._from", []LnsAny{self.FP.getFullName(_env, nonnilableType)})
        
    } else if _switch1 == Ast_TypeInfoKind__Prim {
        if _switch2 := nonnilableType; _switch2 == Ast_builtinTypeInt {
            funcTxt = "_lune._toInt"
            
        } else if _switch2 == Ast_builtinTypeReal {
            funcTxt = "_lune._toReal"
            
        } else if _switch2 == Ast_builtinTypeBool {
            funcTxt = "_lune._toBool"
            
        } else {
            Util_err(_env, _env.LuaVM.String_format("unknown type -- %s", []LnsAny{nonnilableType.FP.GetTxt(_env, nil, nil, nil)}))
        }
    } else if _switch1 == Ast_TypeInfoKind__Map {
        funcTxt = "_lune._toMap"
        
        var itemList *LnsList
        itemList = nonnilableType.FP.Get_itemTypeInfoList(_env)
        var keyFuncTxt string
        var keyNilable bool
        var keyChild string
        keyFuncTxt,keyNilable,keyChild = self.FP.getMapInfo(_env, itemList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        var valFuncTxt string
        var valNilable bool
        var valChild string
        valFuncTxt,valNilable,valChild = self.FP.getMapInfo(_env, itemList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        child = _env.LuaVM.String_format("{ { func = %s, nilable = %s, child = %s }, \n", []LnsAny{keyFuncTxt, keyNilable, keyChild}) + _env.LuaVM.String_format("{ func = %s, nilable = %s, child = %s } }", []LnsAny{valFuncTxt, valNilable, valChild})
        
    } else if _switch1 == Ast_TypeInfoKind__Set {
        funcTxt = "_lune._toSet"
        
        var itemList *LnsList
        itemList = nonnilableType.FP.Get_itemTypeInfoList(_env)
        var valFuncTxt string
        var valNilable bool
        var valChild string
        valFuncTxt,valNilable,valChild = self.FP.getMapInfo(_env, itemList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        child = _env.LuaVM.String_format("{ func = %s, nilable = %s, child = %s }", []LnsAny{valFuncTxt, valNilable, valChild})
        
    } else if _switch1 == Ast_TypeInfoKind__List || _switch1 == Ast_TypeInfoKind__Array {
        funcTxt = "_lune._toList"
        
        var itemList *LnsList
        itemList = nonnilableType.FP.Get_itemTypeInfoList(_env)
        var valFuncTxt string
        var valNilable bool
        var valChild string
        valFuncTxt,valNilable,valChild = self.FP.getMapInfo(_env, itemList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        child = _env.LuaVM.String_format("{ { func = %s, nilable = %s, child = %s } }", []LnsAny{valFuncTxt, valNilable, valChild})
        
    } else if _switch1 == Ast_TypeInfoKind__Alternate {
        var prefix string
        prefix = _env.LuaVM.String_format("obj.__alt2mapFunc.%s", []LnsAny{nonnilableType.FP.Get_rawTxt(_env)})
        funcTxt = _env.LuaVM.String_format("%s.func", []LnsAny{prefix})
        
        child = _env.LuaVM.String_format("%s.child", []LnsAny{prefix})
        
    }
    return funcTxt, typeInfo.FP.Get_nilable(_env), child
}

// 1380: decl @lune.@base.@convLua.ConvFilter.processDeclAlge
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
        algeFullName = _env.LuaVM.String_format("%s.%s", []LnsAny{self.FP.getFullName(_env, parentInfo), algeFullName})
        
        access = ""
        
        isTopNS = false
        
    }
    self.FP.writeln(_env, _env.LuaVM.String_format("%s%s = {}", []LnsAny{access, algeFullName}))
    self.FP.writeln(_env, _env.LuaVM.String_format("%s._name2Val = {}", []LnsAny{algeFullName}))
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( isTopNS) &&
        _env.SetStackVal( node.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) ).(bool)){
        if self.needModuleObj{
            self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.%s = %s", []LnsAny{algeFullName, algeFullName}))
        }
    }
    if typeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
        self.pubAlgeId2AlgeTypeInfo.Set(typeInfo.FP.Get_typeId(_env),typeInfo)
    }
    self.FP.writeln(_env, _env.LuaVM.String_format("function %s:_getTxt( val )\n   local name = val[ 1 ]\n   if name then\n      return string.format( \"%s.%%s\", name )\n   end\n   return string.format( \"illegal val -- %%s\", val )\nend\n", []LnsAny{algeFullName, algeFullName}))
    self.FP.writeln(_env, _env.LuaVM.String_format("function %s._from( val )\n   return _lune._AlgeFrom( %s, val )\nend\n", []LnsAny{algeFullName, algeFullName}))
    {
        __forsortCollection1 := node.FP.Get_algeType(_env).FP.Get_valInfoMap(_env)
        __forsortSorted1 := __forsortCollection1.CreateKeyListStr()
        __forsortSorted1.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey1 := range( __forsortSorted1.Items ) {
            valInfo := __forsortCollection1.Items[ ___forsortKey1 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            self.FP.Write(_env, _env.LuaVM.String_format("%s.%s = { \"%s\"", []LnsAny{algeFullName, valInfo.FP.Get_name(_env), valInfo.FP.Get_name(_env)}))
            if valInfo.FP.Get_typeList(_env).Len() > 0{
                self.FP.Write(_env, ", {")
                for _index, _paramType := range( valInfo.FP.Get_typeList(_env).Items ) {
                    index := _index + 1
                    paramType := _paramType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if index > 1{
                        self.FP.Write(_env, ",")
                    }
                    if Ast_NormalTypeInfo_isAvailableMapping(_env, self.processInfo, &node.FP.Get_algeType(_env).Ast_TypeInfo, NewLnsMap( map[LnsAny]LnsAny{})){
                        var funcTxt string
                        var nilable bool
                        var child string
                        funcTxt,nilable,child = self.FP.getMapInfo(_env, paramType)
                        self.FP.Write(_env, _env.LuaVM.String_format("{ func=%s, nilable=%s, child=%s }", []LnsAny{funcTxt, nilable, child}))
                    } else { 
                        self.FP.Write(_env, "{}")
                    }
                }
                self.FP.Write(_env, "}")
            }
            self.FP.writeln(_env, "}")
            self.FP.writeln(_env, _env.LuaVM.String_format("%s._name2Val[\"%s\"] = %s.%s", []LnsAny{algeFullName, valInfo.FP.Get_name(_env), algeFullName, valInfo.FP.Get_name(_env)}))
        }
    }
}

// 1448: decl @lune.@base.@convLua.ConvFilter.processNewAlgeVal
func (self *convLua_ConvFilter) ProcessNewAlgeVal(_env *LnsEnv, node *Nodes_NewAlgeValNode,_opt LnsAny) {
    var valInfo *Ast_AlgeValInfo
    valInfo = node.FP.Get_valInfo(_env)
    self.FP.Write(_env, _env.LuaVM.String_format("_lune.newAlge( %s.%s", []LnsAny{self.FP.getFullName(_env, &node.FP.Get_algeTypeInfo(_env).Ast_TypeInfo), valInfo.FP.Get_name(_env)}))
    if valInfo.FP.Get_typeList(_env).Len() > 0{
        self.FP.Write(_env, ", {")
        for _index, _exp := range( node.FP.Get_paramList(_env).Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            if index > 1{
                self.FP.Write(_env, ",")
            }
            convLua_filter_7_(_env, exp, self, &node.Nodes_Node)
        }
        self.FP.Write(_env, "}")
    }
    self.FP.Write(_env, ")")
}

// 1469: decl @lune.@base.@convLua.ConvFilter.getDestrClass
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

// 1482: decl @lune.@base.@convLua.ConvFilter.outputAlter2MapFunc
func (self *convLua_ConvFilter) outputAlter2MapFunc(_env *LnsEnv, stream Lns_oStream,alt2Map *LnsMap) {
    stream.Write(_env, "{")
    for _altType, _assinType := range( alt2Map.Items ) {
        altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        assinType := _assinType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if altType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
            if assinType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                stream.Write(_env, _env.LuaVM.String_format("%s = obj.__alt2mapFunc.%s,", []LnsAny{altType.FP.Get_rawTxt(_env), assinType.FP.Get_rawTxt(_env)}))
            } else { 
                var funcTxt string
                var nilable bool
                var child string
                funcTxt,nilable,child = self.FP.getMapInfo(_env, assinType)
                stream.Write(_env, _env.LuaVM.String_format("%s = { func=%s, nilable=%s, child=%s },", []LnsAny{altType.FP.Get_rawTxt(_env), funcTxt, nilable, child}))
            }
        }
    }
    stream.Write(_env, "}")
}

// 1507: decl @lune.@base.@convLua.ConvFilter.processProtoClass
func (self *convLua_ConvFilter) ProcessProtoClass(_env *LnsEnv, node *Nodes_ProtoClassNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("local %s = {}", []LnsAny{node.FP.Get_name(_env).Txt}))
}

// 1513: decl @lune.@base.@convLua.ConvFilter.processDeclClass
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
            self.FP.Write(_env, _env.LuaVM.String_format("local %s = ", []LnsAny{className}))
            if node.FP.Get_lazyLoad(_env) == Nodes_LazyLoad__Off{
                self.FP.Write(_env, "require")
            } else { 
                self.FP.Write(_env, "_lune._lazyRequire")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("( %s )", []LnsAny{_exp.Txt}))
            if _switch1 := node.FP.Get_accessMode(_env); _switch1 == Ast_AccessMode__Pub || _switch1 == Ast_AccessMode__Pro {
                if self.needModuleObj{
                    self.FP.writeln(_env, "")
                    self.FP.Write(_env, _env.LuaVM.String_format("_moduleObj.%s = %s", []LnsAny{className, className}))
                }
            }
            return 
        }
    }
    if Lns_op_not(node.FP.Get_hasPrototype(_env)){
        self.FP.writeln(_env, _env.LuaVM.String_format("local %s = {}", []LnsAny{className}))
    }
    var ifTxt string
    ifTxt = ""
    if classTypeInfo.FP.Get_interfaceList(_env).Len() > 0{
        for _, _ifType := range( classTypeInfo.FP.Get_interfaceList(_env).Items ) {
            ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            ifTxt = ifTxt + self.FP.getFullName(_env, ifType) + ","
            
        }
        ifTxt = _env.LuaVM.String_format("ifList = {%s}", []LnsAny{ifTxt})
        
    }
    var baseInfo *Ast_TypeInfo
    baseInfo = classTypeInfo.FP.Get_baseTypeInfo(_env)
    var baseTxt string
    baseTxt = ""
    if baseInfo.FP.Get_typeId(_env) != Ast_rootTypeIdInfo{
        baseTxt = _env.LuaVM.String_format("__index = %s", []LnsAny{self.FP.getFullName(_env, baseInfo)})
        
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( len(ifTxt) > 0) ||
        _env.SetStackVal( len(baseTxt) > 0) ).(bool){
        var metaTxt string
        metaTxt = baseTxt
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( len(baseTxt) > 0) &&
            _env.SetStackVal( len(ifTxt) > 0) ).(bool)){
            metaTxt = _env.LuaVM.String_format("%s,%s", []LnsAny{baseTxt, ifTxt})
            
        } else if len(ifTxt) > 0{
            metaTxt = ifTxt
            
        }
        self.FP.writeln(_env, _env.LuaVM.String_format("setmetatable( %s, { %s } )", []LnsAny{className, metaTxt}))
    }
    if nodeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
        if self.needModuleObj{
            self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.%s = %s", []LnsAny{className, className}))
        }
    }
    for _, _declNode := range( node.FP.Get_declStmtList(_env).Items ) {
        declNode := _declNode.(Nodes_NodeDownCast).ToNodes_Node()
        convLua_filter_7_(_env, declNode, self, &node.Nodes_Node)
    }
    var hasConstrFlag bool
    hasConstrFlag = false
    var memberList *LnsList
    memberList = NewLnsList([]LnsAny{})
    var fieldList *LnsList
    fieldList = nodeInfo.FP.Get_fieldList(_env)
    var outerMethodSet *LnsSet
    outerMethodSet = nodeInfo.FP.Get_outerMethodSet(_env)
    var methodNameSet *LnsSet
    methodNameSet = NewLnsSet([]LnsAny{})
    if classTypeInfo.FP.Get_kind(_env) != Ast_TypeInfoKind__IF{
        for _, _field := range( fieldList.Items ) {
            field := _field.(Nodes_NodeDownCast).ToNodes_Node()
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
                        memberList.Insert(Nodes_DeclMemberNode2Stem(declMemberNode))
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
                convLua_filter_7_(_env, field, self, &node.Nodes_Node)
            }
        }
    }
    var destTxt string
    destTxt = ""
    {
        __exp := self.FP.getDestrClass(_env, node.FP.Get_expType(_env))
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            destTxt = _env.LuaVM.String_format(", __gc = %s.__free", []LnsAny{_exp.FP.GetTxt(_env, nil, nil, nil)})
            
        }
    }
    self.FP.writeln(_env, _env.LuaVM.String_format("function %s.setmeta( obj )\n  setmetatable( obj, { __index = %s %s } )\nend", []LnsAny{className, className, destTxt}))
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
                        superArgTxt = _env.LuaVM.String_format("%s__superarg%d", []LnsAny{superArgTxt, index})
                        
                    }
                }
            }
        }
        for _, _member := range( memberList.Items ) {
            member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
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
        self.FP.writeln(_env, _env.LuaVM.String_format("function %s.new( %s )\n   local obj = {}\n   %s.setmeta( obj )\n   if obj.__init then\n      obj:__init( %s )\n   end\n   return obj\nend\nfunction %s:__init( %s )\n", []LnsAny{className, argTxt, className, argTxt, className, argTxt}))
        self.FP.pushIndent(_env, nil)
        if baseInfo != Ast_headTypeInfo{
            if Lns_isCondTrue( (Lns_unwrap( baseInfo.FP.Get_scope(_env)).(*Ast_Scope)).FP.GetSymbolInfoChild(_env, "__init")){
                self.FP.Write(_env, _env.LuaVM.String_format("%s.__init( self", []LnsAny{self.FP.getFullName(_env, baseInfo)}))
                if len(superArgTxt) > 0{
                    self.FP.writeln(_env, _env.LuaVM.String_format(", %s )", []LnsAny{superArgTxt}))
                } else { 
                    self.FP.writeln(_env, ")")
                }
            }
        }
        for _, _member := range( memberList.Items ) {
            member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            var memberName string
            memberName = member.FP.Get_name(_env).Txt
            self.FP.writeln(_env, _env.LuaVM.String_format("self.%s = %s", []LnsAny{memberName, memberName}))
        }
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "end")
    }
    for _, _memberNode := range( nodeInfo.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
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
            self.FP.writeln(_env, _env.LuaVM.String_format("function %s%s%s()\n   return %s.%s\nend", []LnsAny{className, delimit, getterName, prefix, memberName}))
            methodNameSet.Add(getterName)
        }
        var setterName string
        setterName = "set_" + memberName
        autoFlag = Lns_op_not(methodNameSet.Has(setterName))
        
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( memberNode.FP.Get_setterMode(_env) != Ast_AccessMode__None) &&
            _env.SetStackVal( autoFlag) ).(bool)){
            self.FP.writeln(_env, _env.LuaVM.String_format("function %s%s%s( %s )\n   %s.%s = %s\nend", []LnsAny{className, delimit, setterName, memberName, prefix, memberName, memberName}))
            methodNameSet.Add(setterName)
        }
    }
    for _, _advertiseInfo := range( node.FP.Get_advertiseList(_env).Items ) {
        advertiseInfo := _advertiseInfo.(Nodes_AdvertiseInfoDownCast).ToNodes_AdvertiseInfo()
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
                    self.FP.writeln(_env, _env.LuaVM.String_format("function %s:%s( ... )\n   return self.%s:%s( ... )\nend\n", []LnsAny{className, childName, memberName, childName}))
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
                    initStmt := _initStmt.(Nodes_NodeDownCast).ToNodes_Node()
                    convLua_filter_7_(_env, initStmt, self, &node.Nodes_Node)
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
        if convLua_isGenericType_10_(_env, classTypeInfo){
            declArgTxt = "val, __alt2mapFunc"
            
            argTxt = "{ __alt2mapFunc = __alt2mapFunc }, val"
            
        }
        self.FP.writeln(_env, _env.LuaVM.String_format("function %s:_toMap()\n  return self\nend\nfunction %s._fromMap( %s )\n  local obj, mes = %s._fromMapSub( %s )\n  if obj then\n     %s.setmeta( obj )\n  end\n  return obj, mes\nend\nfunction %s._fromStem( %s )\n  return %s._fromMap( %s )\nend\n", []LnsAny{className, className, declArgTxt, className, argTxt, className, className, declArgTxt, className, declArgTxt}))
        self.FP.writeln(_env, _env.LuaVM.String_format("function %s._fromMapSub( obj, val )", []LnsAny{className}))
        if classTypeInfo.FP.Get_baseTypeInfo(_env) != Ast_headTypeInfo{
            self.FP.writeln(_env, _env.LuaVM.String_format("   local result, mes = %s._fromMapSub( obj, val )\n   if not result then\n      return nil, mes\n   end\n", []LnsAny{self.FP.getFullName(_env, classTypeInfo.FP.Get_baseTypeInfo(_env))}))
        }
        self.FP.writeln(_env, "   local memInfo = {}")
        for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            var funcTxt string
            var nilable bool
            var child string
            funcTxt,nilable,child = self.FP.getMapInfo(_env, memberNode.FP.Get_expType(_env))
            self.FP.writeln(_env, _env.LuaVM.String_format("   table.insert( memInfo, { name = \"%s\", func = %s, nilable = %s, child = %s } )", []LnsAny{memberNode.FP.Get_name(_env).Txt, funcTxt, nilable, child}))
        }
        self.FP.writeln(_env, "   local result, mess = _lune._fromMap( obj, val, memInfo )\n   if not result then\n      return nil, mess\n   end\n   return obj\nend")
    }
}

// 1842: decl @lune.@base.@convLua.ConvFilter.processDeclMember
func (self *convLua_ConvFilter) ProcessDeclMember(_env *LnsEnv, node *Nodes_DeclMemberNode,_opt LnsAny) {
}

// 1849: decl @lune.@base.@convLua.ConvFilter.processExpMacroExp
func (self *convLua_ConvFilter) ProcessExpMacroExp(_env *LnsEnv, node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    for _, _stmt := range( node.FP.Get_stmtList(_env).Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        convLua_filter_7_(_env, stmt, self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
    }
}

// 1861: decl @lune.@base.@convLua.ConvFilter.outputDeclMacro
func (self *convLua_ConvFilter) OutputDeclMacro(_env *LnsEnv, name string,argNameList *LnsList,callback convLua_outputMacroStmtBlock_11_) {
    self.FP.Write(_env, _env.LuaVM.String_format("local function %s(", []LnsAny{name}))
    self.FP.writeln(_env, "__macroArgs )")
    self.FP.pushIndent(_env, nil)
    self.FP.writeln(_env, _env.LuaVM.String_format("local _lune = require( \"%s\" )", []LnsAny{Option_getRuntimeModule(_env)}))
    self.FP.writeln(_env, "local __var = __macroArgs.__var")
    for _, _argName := range( argNameList.Items ) {
        argName := _argName.(string)
        self.FP.writeln(_env, _env.LuaVM.String_format("local %s = __macroArgs.%s", []LnsAny{argName, argName}))
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
    self.FP.writeln(_env, _env.LuaVM.String_format("return %s", []LnsAny{name}))
}

// 1905: decl @lune.@base.@convLua.ConvFilter.processExpMacroStatList
func (self *convLua_ConvFilter) ProcessExpMacroStatList(_env *LnsEnv, node *Nodes_ExpMacroStatListNode,_opt LnsAny) {
    self.FP.Write(_env, "__expStatList(")
    convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, ")")
}

// 1913: decl @lune.@base.@convLua.ConvFilter.processDeclMacro
func (self *convLua_ConvFilter) ProcessDeclMacro(_env *LnsEnv, node *Nodes_DeclMacroNode,_opt LnsAny) {
    if self.inMacro{
        var macroInfo *Nodes_DeclMacroInfo
        macroInfo = node.FP.Get_declInfo(_env)
        var argNameList *LnsList
        argNameList = NewLnsList([]LnsAny{})
        for _, _arg := range( macroInfo.FP.Get_argList(_env).Items ) {
            arg := _arg.(Nodes_DeclArgNodeDownCast).ToNodes_DeclArgNode()
            argNameList.Insert(arg.FP.Get_name(_env).Txt)
        }
        self.FP.OutputDeclMacro(_env, macroInfo.FP.Get_name(_env).Txt, argNameList, convLua_outputMacroStmtBlock_11_(func(_env *LnsEnv) {
            {
                _stmtBlock := macroInfo.FP.Get_stmtBlock(_env)
                if !Lns_IsNil( _stmtBlock ) {
                    stmtBlock := _stmtBlock.(*Nodes_BlockNode)
                    convLua_filter_7_(_env, &stmtBlock.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }))
    }
}

// 1932: decl @lune.@base.@convLua.ConvFilter.processExpMacroStat
func (self *convLua_ConvFilter) ProcessExpMacroStat(_env *LnsEnv, node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    if node.FP.Get_expStrList(_env).Len() == 0{
        self.FP.Write(_env, "''")
    } else { 
        for _index, _token := range( node.FP.Get_expStrList(_env).Items ) {
            index := _index + 1
            token := _token.(Nodes_NodeDownCast).ToNodes_Node()
            if index != 1{
                self.FP.Write(_env, "..")
            }
            convLua_filter_7_(_env, token, self, &node.Nodes_Node)
        }
    }
}

// 1950: decl @lune.@base.@convLua.ConvFilter.processExpNew
func (self *convLua_ConvFilter) ProcessExpNew(_env *LnsEnv, node *Nodes_ExpNewNode,_opt LnsAny) {
    convLua_filter_7_(_env, node.FP.Get_symbol(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, ".new(")
    {
        __exp := node.FP.Get_argList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write(_env, ")")
}

// 1961: decl @lune.@base.@convLua.ConvFilter.process__func__symbol
func (self *convLua_ConvFilter) process__func__symbol(_env *LnsEnv, has__func__Symbol bool,parentType *Ast_TypeInfo,funcName string) {
    if has__func__Symbol{
        var nameSpace string
        nameSpace = self.FP.getCanonicalName(_env, parentType, false)
        if funcName == ""{
            funcName = "<anonymous>"
            
        }
        self.FP.pushIndent(_env, nil)
        self.FP.writeln(_env, _env.LuaVM.String_format("local __func__ = '%s.%s'", []LnsAny{nameSpace, funcName}))
        self.FP.popIndent(_env)
    }
}

// 1975: decl @lune.@base.@convLua.ConvFilter.processDeclConstr
func (self *convLua_ConvFilter) ProcessDeclConstr(_env *LnsEnv, node *Nodes_DeclConstrNode,_opt LnsAny) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo(_env)
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = Lns_unwrap( declInfo.FP.Get_classTypeInfo(_env)).(*Ast_TypeInfo)
    var className string
    className = self.FP.getFullName(_env, classTypeInfo)
    self.FP.Write(_env, _env.LuaVM.String_format("function %s.new( ", []LnsAny{className}))
    var argTxt string
    argTxt = ""
    self.FP.Write(_env, argTxt)
    var argList *LnsList
    argList = declInfo.FP.Get_argList(_env)
    for _, _arg := range( argList.Items ) {
        arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
        if len(argTxt) > 0{
            self.FP.Write(_env, ", ")
            argTxt = argTxt + ", "
            
        }
        convLua_filter_7_(_env, arg, self, &node.Nodes_Node)
        {
            __exp := Nodes_DeclArgNodeDownCastF(arg.FP)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_DeclArgNode)
                argTxt = argTxt + _exp.FP.Get_name(_env).Txt
                
            } else {
                var name *Types_Token
                name = Lns_unwrap( node.FP.Get_declInfo(_env).FP.Get_name(_env)).(*Types_Token)
                Util_err(_env, _env.LuaVM.String_format("not support ... in macro -- %s", []LnsAny{name.Txt}))
            }
        }
    }
    self.FP.writeln(_env, " )")
    self.FP.pushIndent(_env, nil)
    self.FP.writeln(_env, "local obj = {}")
    self.FP.writeln(_env, _env.LuaVM.String_format("%s.setmeta( obj )", []LnsAny{className}))
    self.FP.writeln(_env, _env.LuaVM.String_format("if obj.__init then obj:__init( %s ); end", []LnsAny{argTxt}))
    self.FP.writeln(_env, "return obj")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
    self.FP.Write(_env, _env.LuaVM.String_format("function %s:__init(%s) ", []LnsAny{className, argTxt}))
    {
        __exp := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.process__func__symbol(_env, declInfo.FP.Get_has__func__Symbol(_env), node.FP.Get_expType(_env).FP.Get_parentInfo(_env), "__init")
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, "end")
}

// 2030: decl @lune.@base.@convLua.ConvFilter.processDeclDestr
func (self *convLua_ConvFilter) ProcessDeclDestr(_env *LnsEnv, node *Nodes_DeclDestrNode,_opt LnsAny) {
    self.FP.writeln(_env, _env.LuaVM.String_format("function %s.__free( self )", []LnsAny{_env.NilAccFin(_env.NilAccPush(node.FP.Get_declInfo(_env).FP.Get_classTypeInfo(_env)) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetTxt(_env, nil, nil, nil)})/* 2034:9 */)}))
    self.FP.process__func__symbol(_env, node.FP.Get_declInfo(_env).FP.Get_has__func__Symbol(_env), node.FP.Get_expType(_env).FP.Get_parentInfo(_env), "__free")
    convLua_filter_7_(_env, &Lns_unwrap( node.FP.Get_declInfo(_env).FP.Get_body(_env)).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = node.FP.Get_expType(_env).FP.Get_parentInfo(_env)
    {
        __exp := self.FP.getDestrClass(_env, classTypeInfo.FP.Get_baseTypeInfo(_env))
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            self.FP.writeln(_env, _env.LuaVM.String_format("%s.__free( self )", []LnsAny{_exp.FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    self.FP.writeln(_env, "end")
}

// 2050: decl @lune.@base.@convLua.ConvFilter.processExpCallSuperCtor
func (self *convLua_ConvFilter) ProcessExpCallSuperCtor(_env *LnsEnv, node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType(_env)
    self.FP.Write(_env, _env.LuaVM.String_format("%s.%s( self", []LnsAny{self.FP.getFullName(_env, typeInfo), node.FP.Get_methodType(_env).FP.Get_rawTxt(_env)}))
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(_env, ",")
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, ")")
}

// 2065: decl @lune.@base.@convLua.ConvFilter.processExpCallSuper
func (self *convLua_ConvFilter) ProcessExpCallSuper(_env *LnsEnv, node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType(_env)
    self.FP.Write(_env, _env.LuaVM.String_format("%s.%s( self", []LnsAny{self.FP.getFullName(_env, typeInfo), node.FP.Get_methodType(_env).FP.Get_rawTxt(_env)}))
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(_env, ",")
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write(_env, ")")
}

// 2081: decl @lune.@base.@convLua.ConvFilter.processDeclMethod
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
    self.FP.Write(_env, _env.LuaVM.String_format("function %s%s%s( ", []LnsAny{classTypeInfo.FP.Get_rawTxt(_env), delimit, methodName}))
    var argList *LnsList
    argList = declInfo.FP.Get_argList(_env)
    for _index, _arg := range( argList.Items ) {
        index := _index + 1
        arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        convLua_filter_7_(_env, arg, self, &node.Nodes_Node)
    }
    self.FP.writeln(_env, " )")
    {
        __exp := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.process__func__symbol(_env, declInfo.FP.Get_has__func__Symbol(_env), node.FP.Get_expType(_env).FP.Get_parentInfo(_env), methodName)
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, "end")
}

// 2114: decl @lune.@base.@convLua.ConvFilter.processUnwrapSet
func (self *convLua_ConvFilter) ProcessUnwrapSet(_env *LnsEnv, node *Nodes_UnwrapSetNode,_opt LnsAny) {
    var dstExpList *Nodes_ExpListNode
    dstExpList = node.FP.Get_dstExpList(_env)
    convLua_filter_7_(_env, &dstExpList.Nodes_Node, self, &node.Nodes_Node)
    self.FP.Write(_env, " = ")
    convLua_filter_7_(_env, &node.FP.Get_srcExpList(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
    self.FP.Write(_env, "if ")
    for _index, _expNode := range( dstExpList.FP.Get_expList(_env).Items ) {
        index := _index + 1
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        if index > 1{
            self.FP.Write(_env, " or ")
        }
        self.FP.Write(_env, "nil == ")
        convLua_filter_7_(_env, expNode, self, &node.Nodes_Node)
    }
    self.FP.writeln(_env, " then")
    self.FP.pushIndent(_env, nil)
    for _index, _expNode := range( dstExpList.FP.Get_expList(_env).Items ) {
        index := _index + 1
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.Write(_env, _env.LuaVM.String_format("local _exp%d = ", []LnsAny{index}))
        convLua_filter_7_(_env, expNode, self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
    }
    if Lns_isCondTrue( node.FP.Get_unwrapBlock(_env)){
        convLua_filter_7_(_env, &Lns_unwrap( node.FP.Get_unwrapBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
    }
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}

// 2147: decl @lune.@base.@convLua.ConvFilter.processExpListSub
func (self *convLua_ConvFilter) processExpListSub(_env *LnsEnv, parent *Nodes_Node,expList *LnsList,mRetExp LnsAny) {
    var mRetIndex LnsAny
    mRetIndex = _env.NilAccFin(_env.NilAccPush(mRetExp) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
    for _index, _exp := range( expList.Items ) {
        index := _index + 1
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
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
            self.FP.Write(_env, ", ")
        }
        convLua_filter_7_(_env, exp, self, parent)
        if index == mRetIndex{
            break
        }
    }
}

// 2174: decl @lune.@base.@convLua.ConvFilter.processExpMRet
func (self *convLua_ConvFilter) ProcessExpMRet(_env *LnsEnv, node *Nodes_ExpMRetNode,_opt LnsAny) {
    convLua_filter_7_(_env, node.FP.Get_mRet(_env), self, &node.Nodes_Node)
}

// 2180: decl @lune.@base.@convLua.ConvFilter.processIfUnwrap
func (self *convLua_ConvFilter) ProcessIfUnwrap(_env *LnsEnv, node *Nodes_IfUnwrapNode,_opt LnsAny) {
    self.FP.writeln(_env, "do")
    self.FP.pushIndent(_env, nil)
    self.FP.Write(_env, "local ")
    for _index, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        index := _index + 1
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        self.FP.Write(_env, convLua_getSymbolTxt_5_(_env, varSym))
        if index != node.FP.Get_varSymList(_env).Len(){
            self.FP.Write(_env, ", ")
        }
    }
    self.FP.Write(_env, " = ")
    self.FP.processExpListSub(_env, &node.Nodes_Node, node.FP.Get_expList(_env).FP.Get_expList(_env), node.FP.Get_expList(_env).FP.Get_mRetExp(_env))
    self.FP.writeln(_env, "")
    self.FP.Write(_env, "if ")
    var hasSym bool
    hasSym = false
    for _, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if varSym.FP.Get_name(_env) != "_"{
            if hasSym{
                self.FP.Write(_env, " and  ")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("%s ~= nil", []LnsAny{convLua_getSymbolTxt_5_(_env, varSym)}))
            hasSym = true
            
        }
    }
    self.FP.Write(_env, " then")
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    {
        __exp := node.FP.Get_nilBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write(_env, "else")
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, "end")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}

// 2223: decl @lune.@base.@convLua.ConvFilter.processWhen
func (self *convLua_ConvFilter) ProcessWhen(_env *LnsEnv, node *Nodes_WhenNode,_opt LnsAny) {
    self.FP.Write(_env, "if ")
    for _index, _symPair := range( node.FP.Get_symPairList(_env).Items ) {
        index := _index + 1
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        self.FP.Write(_env, _env.LuaVM.String_format("%s ~= nil", []LnsAny{symPair.FP.Get_src(_env).FP.Get_name(_env)}))
        if index != node.FP.Get_symPairList(_env).Len(){
            self.FP.Write(_env, " and ")
        }
    }
    self.FP.Write(_env, " then")
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    {
        __exp := node.FP.Get_elseBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write(_env, "else")
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(_env, "end")
}

// 2245: decl @lune.@base.@convLua.ConvFilter.processDeclVar
func (self *convLua_ConvFilter) ProcessDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,_opt LnsAny) {
    if Lns_isCondTrue( node.FP.Get_syncBlock(_env)){
        self.FP.writeln(_env, "do")
        self.FP.pushIndent(_env, nil)
        for _, _varInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
            varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            self.FP.writeln(_env, _env.LuaVM.String_format("local _sync_%s", []LnsAny{convLua_getSymbolTxt_5_(_env, varInfo)}))
        }
        self.FP.writeln(_env, "do")
        self.FP.pushIndent(_env, nil)
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_mode(_env) != Nodes_DeclVarMode__Unwrap) &&
        _env.SetStackVal( node.FP.Get_accessMode(_env) != Ast_AccessMode__Global) ).(bool)){
        self.FP.Write(_env, "local ")
    }
    var varList *LnsList
    varList = node.FP.Get_symbolInfoList(_env)
    var varNameList *LnsList
    varNameList = NewLnsList([]LnsAny{})
    for _index, __var := range( varList.Items ) {
        index := _index + 1
        _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        var name string
        name = convLua_getSymbolTxt_5_(_env, _var)
        self.FP.Write(_env, name)
        varNameList.Insert(name)
    }
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(_env, " = ")
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        } else {
            self.FP.writeln(_env, "")
        }
    }
    {
        __exp := node.FP.Get_unwrapBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.writeln(_env, "")
            self.FP.Write(_env, "if ")
            for _index, _varName := range( varNameList.Items ) {
                index := _index + 1
                varName := _varName.(string)
                if index > 1{
                    self.FP.Write(_env, " or ")
                }
                self.FP.Write(_env, " nil == " + varName)
            }
            self.FP.writeln(_env, " then")
            self.FP.pushIndent(_env, nil)
            for _, _varName := range( varNameList.Items ) {
                varName := _varName.(string)
                self.FP.writeln(_env, _env.LuaVM.String_format("local _%s = %s", []LnsAny{varName, varName}))
            }
            self.FP.popIndent(_env)
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
            {
                _thenBlock := node.FP.Get_thenBlock(_env)
                if !Lns_IsNil( _thenBlock ) {
                    thenBlock := _thenBlock.(*Nodes_BlockNode)
                    self.FP.writeln(_env, "else")
                    self.FP.pushIndent(_env, nil)
                    convLua_filter_7_(_env, &thenBlock.Nodes_Node, self, &node.Nodes_Node)
                    self.FP.popIndent(_env)
                }
            }
            self.FP.writeln(_env, "end")
        }
    }
    {
        __exp := node.FP.Get_syncBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
            var syncVarNameList *LnsList
            syncVarNameList = NewLnsList([]LnsAny{})
            for _, _varInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
                varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                var name string
                name = convLua_getSymbolTxt_5_(_env, varInfo)
                syncVarNameList.Insert(name)
                self.FP.writeln(_env, _env.LuaVM.String_format("_sync_%s = %s", []LnsAny{name, name}))
            }
            self.FP.popIndent(_env)
            self.FP.writeln(_env, "end")
            for _, _name := range( syncVarNameList.Items ) {
                name := _name.(string)
                self.FP.writeln(_env, _env.LuaVM.String_format("%s = _sync_%s", []LnsAny{name, name}))
            }
            self.FP.popIndent(_env)
            self.FP.writeln(_env, "end")
        }
    }
    if _switch1 := node.FP.Get_accessMode(_env); _switch1 == Ast_AccessMode__Pub || _switch1 == Ast_AccessMode__Global {
        self.FP.writeln(_env, "")
        for _index, _varName := range( varNameList.Items ) {
            index := _index + 1
            varName := _varName.(string)
            var name string
            name = varName
            if self.needModuleObj{
                self.FP.writeln(_env, _env.LuaVM.String_format("_moduleObj.%s = %s", []LnsAny{name, name}))
            }
            self.pubVarName2InfoMap.Set(name,NewconvLua_PubVerInfo(_env, node.FP.Get_staticFlag(_env), node.FP.Get_accessMode(_env), node.FP.Get_symbolInfoList(_env).GetAt(index).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_mutable(_env), node.FP.Get_typeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))
        }
    }
    if self.macroDepth > 0{
        self.FP.writeln(_env, "")
        for _, _symbolInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            var varName string
            varName = convLua_getSymbolTxt_5_(_env, symbolInfo)
            self.FP.writeln(_env, _env.LuaVM.String_format("table.insert( macroVar.__names, '%s' )", []LnsAny{varName}))
            self.FP.writeln(_env, _env.LuaVM.String_format("macroVar.%s = %s", []LnsAny{varName, varName}))
            self.macroVarSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
        }
    }
}

// 2358: decl @lune.@base.@convLua.ConvFilter.processDeclArg
func (self *convLua_ConvFilter) ProcessDeclArg(_env *LnsEnv, node *Nodes_DeclArgNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{node.FP.Get_name(_env).Txt}))
}

// 2366: decl @lune.@base.@convLua.ConvFilter.processDeclArgDDD
func (self *convLua_ConvFilter) ProcessDeclArgDDD(_env *LnsEnv, node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    self.FP.Write(_env, "...")
}

// 2372: decl @lune.@base.@convLua.ConvFilter.processDeclFunc
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
    self.FP.Write(_env, _env.LuaVM.String_format("%sfunction %s( ", []LnsAny{letTxt, name}))
    var argList *LnsList
    argList = declInfo.FP.Get_argList(_env)
    self.FP.processExpListSub(_env, &node.Nodes_Node, argList, nil)
    self.FP.writeln(_env, " )")
    {
        __exp := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.process__func__symbol(_env, declInfo.FP.Get_has__func__Symbol(_env), node.FP.Get_expType(_env).FP.Get_parentInfo(_env), name)
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write(_env, "end")
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType(_env)
    if _switch1 := expType.FP.Get_accessMode(_env); _switch1 == Ast_AccessMode__Pub || _switch1 == Ast_AccessMode__Global {
        if self.needModuleObj{
            self.FP.writeln(_env, "")
            self.FP.Write(_env, _env.LuaVM.String_format("_moduleObj.%s = %s", []LnsAny{name, name}))
        }
        self.pubFuncName2InfoMap.Set(name,NewconvLua_PubFuncInfo(_env, declInfo.FP.Get_accessMode(_env), node.FP.Get_expType(_env)))
    }
}

// 2414: decl @lune.@base.@convLua.ConvFilter.processRefType
func (self *convLua_ConvFilter) ProcessRefType(_env *LnsEnv, node *Nodes_RefTypeNode,_opt LnsAny) {
    if _switch1 := node.FP.Get_mutMode(_env); _switch1 == Ast_MutMode__IMut {
        self.FP.Write(_env, "&")
    } else if _switch1 == Ast_MutMode__AllMut {
        self.FP.Write(_env, "+")
    }
    convLua_filter_7_(_env, node.FP.Get_name(_env), self, &node.Nodes_Node)
    if node.FP.Get_array(_env) == "array"{
        self.FP.Write(_env, "[@]")
    } else if node.FP.Get_array(_env) == "list"{
        self.FP.Write(_env, "[]")
    }
}

// 2434: decl @lune.@base.@convLua.ConvFilter.processIf
func (self *convLua_ConvFilter) ProcessIf(_env *LnsEnv, node *Nodes_IfNode,_opt LnsAny) {
    var valList *LnsList
    valList = node.FP.Get_stmtList(_env)
    for _index, _val := range( valList.Items ) {
        index := _index + 1
        val := _val.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if index == 1{
            self.FP.Write(_env, "if ")
            convLua_filter_7_(_env, val.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else if val.FP.Get_kind(_env) == Nodes_IfKind__ElseIf{
            self.FP.Write(_env, "elseif ")
            convLua_filter_7_(_env, val.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else { 
            self.FP.writeln(_env, "else")
        }
        self.FP.Write(_env, " ")
        convLua_filter_7_(_env, &val.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    }
    self.FP.writeln(_env, "end")
}

// 2456: decl @lune.@base.@convLua.ConvFilter.processSwitch
func (self *convLua_ConvFilter) ProcessSwitch(_env *LnsEnv, node *Nodes_SwitchNode,_opt LnsAny) {
    self.FP.writeln(_env, "do")
    self.FP.pushIndent(_env, nil)
    self.FP.Write(_env, "local _switchExp = ")
    convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
    if node.FP.Get_caseList(_env).Len() > 0{
        for _index, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
            index := _index + 1
            caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
            if index == 1{
                self.FP.Write(_env, "if ")
            } else { 
                self.FP.Write(_env, "elseif ")
            }
            var expList *Nodes_ExpListNode
            expList = caseInfo.FP.Get_expList(_env)
            for _listIndex, _expNode := range( expList.FP.Get_expList(_env).Items ) {
                listIndex := _listIndex + 1
                expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                if listIndex != 1{
                    self.FP.Write(_env, " or ")
                }
                self.FP.Write(_env, "_switchExp == ")
                convLua_filter_7_(_env, expNode, self, &node.Nodes_Node)
            }
            self.FP.Write(_env, " then")
            convLua_filter_7_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        }
        {
            __exp := node.FP.Get_default(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_BlockNode)
                self.FP.writeln(_env, "else ")
                self.FP.pushIndent(_env, nil)
                convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
                self.FP.popIndent(_env)
            }
        }
        self.FP.writeln(_env, "end")
    }
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}

// 2499: decl @lune.@base.@convLua.ConvFilter.processMatch
func (self *convLua_ConvFilter) ProcessMatch(_env *LnsEnv, node *Nodes_MatchNode,_opt LnsAny) {
    self.FP.writeln(_env, "do")
    self.FP.pushIndent(_env, nil)
    self.FP.Write(_env, "local _matchExp = ")
    convLua_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
    if node.FP.Get_caseList(_env).Len() > 0{
        var fullName string
        fullName = self.FP.getFullName(_env, &node.FP.Get_algeTypeInfo(_env).Ast_TypeInfo)
        for _index, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
            index := _index + 1
            caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
            if index == 1{
                self.FP.Write(_env, "if ")
            } else { 
                self.FP.Write(_env, "elseif ")
            }
            self.FP.writeln(_env, _env.LuaVM.String_format("_matchExp[1] == %s.%s[1] then", []LnsAny{fullName, caseInfo.FP.Get_valInfo(_env).FP.Get_name(_env)}))
            for _paramNum, _paramSym := range( caseInfo.FP.Get_valParamNameList(_env).Items ) {
                paramNum := _paramNum + 1
                paramSym := _paramSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.writeln(_env, _env.LuaVM.String_format("   local %s = _matchExp[2][%d]", []LnsAny{paramSym.FP.Get_name(_env), paramNum}))
            }
            convLua_filter_7_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        }
        {
            __exp := node.FP.Get_defaultBlock(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_Node)
                self.FP.writeln(_env, "else ")
                self.FP.pushIndent(_env, nil)
                convLua_filter_7_(_env, _exp, self, &node.Nodes_Node)
                self.FP.popIndent(_env)
            }
        }
        self.FP.writeln(_env, "end")
    }
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}

// 2538: decl @lune.@base.@convLua.ConvFilter.processWhile
func (self *convLua_ConvFilter) ProcessWhile(_env *LnsEnv, node *Nodes_WhileNode,_opt LnsAny) {
    self.FP.Write(_env, "while ")
    convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, " ")
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
}

// 2549: decl @lune.@base.@convLua.ConvFilter.processRepeat
func (self *convLua_ConvFilter) ProcessRepeat(_env *LnsEnv, node *Nodes_RepeatNode,_opt LnsAny) {
    self.FP.Write(_env, "repeat ")
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Write(_env, "until ")
    convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
}

// 2558: decl @lune.@base.@convLua.ConvFilter.processFor
func (self *convLua_ConvFilter) ProcessFor(_env *LnsEnv, node *Nodes_ForNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("for %s = ", []LnsAny{convLua_getSymbolTxt_5_(_env, node.FP.Get_val(_env))}))
    convLua_filter_7_(_env, node.FP.Get_init(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, ", ")
    convLua_filter_7_(_env, node.FP.Get_to(_env), self, &node.Nodes_Node)
    {
        __exp := node.FP.Get_delta(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write(_env, ", ")
            convLua_filter_7_(_env, _exp, self, &node.Nodes_Node)
        }
    }
    self.FP.Write(_env, " ")
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
}

// 2574: decl @lune.@base.@convLua.ConvFilter.processApply
func (self *convLua_ConvFilter) ProcessApply(_env *LnsEnv, node *Nodes_ApplyNode,_opt LnsAny) {
    self.FP.Write(_env, "for ")
    var varList *LnsList
    varList = node.FP.Get_varList(_env)
    for _index, __var := range( varList.Items ) {
        index := _index + 1
        _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        self.FP.Write(_env, convLua_getSymbolTxt_5_(_env, _var))
    }
    self.FP.Write(_env, " in ")
    convLua_filter_7_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Write(_env, " ")
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
}

// 2592: decl @lune.@base.@convLua.ConvFilter.processForeach
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
    self.FP.Write(_env, "for ")
    if keySym != nil{
        keySym_1521 := keySym.(*Ast_SymbolInfo)
        self.FP.Write(_env, convLua_getSymbolTxt_5_(_env, keySym_1521))
    } else {
        self.FP.Write(_env, "__index")
    }
    self.FP.Write(_env, ", ")
    if valSym != nil{
        valSym_1524 := valSym.(*Ast_SymbolInfo)
        self.FP.Write(_env, convLua_getSymbolTxt_5_(_env, valSym_1524))
    } else {
        self.FP.Write(_env, "__val")
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.useIpairs) &&
        _env.SetStackVal( node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__List) ).(bool)){
        if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilable(_env){
            self.FP.Write(_env, " in pairs( ")
        } else { 
            self.FP.Write(_env, " in ipairs( ")
        }
    } else { 
        self.FP.Write(_env, " in pairs( ")
    }
    convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, " ) ")
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
}

// 2642: decl @lune.@base.@convLua.ConvFilter.processForsort
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
    self.FP.Write(_env, "local __map = ")
    convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.writeln(_env, "")
    self.FP.writeln(_env, "for __key in pairs( __map ) do")
    self.FP.pushIndent(_env, nil)
    self.FP.writeln(_env, "table.insert( __sorted, __key )")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
    self.FP.writeln(_env, "table.sort( __sorted )")
    self.FP.Write(_env, "for __index, ")
    var key string
    key = "__key"
    if keySym != nil{
        keySym_1538 := keySym.(*Ast_SymbolInfo)
        key = convLua_getSymbolTxt_5_(_env, keySym_1538)
        
    }
    self.FP.Write(_env, key)
    self.FP.writeln(_env, " in ipairs( __sorted ) do")
    self.FP.pushIndent(_env, nil)
    if valSym != nil{
        valSym_1540 := valSym.(*Ast_SymbolInfo)
        self.FP.writeln(_env, _env.LuaVM.String_format("local %s = __map[ %s ]", []LnsAny{convLua_getSymbolTxt_5_(_env, valSym_1540), key}))
    }
    convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln(_env, "end")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
    self.FP.popIndent(_env)
    self.FP.writeln(_env, "end")
}

// 2694: decl @lune.@base.@convLua.ConvFilter.processExpUnwrap
func (self *convLua_ConvFilter) ProcessExpUnwrap(_env *LnsEnv, node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    {
        __exp := node.FP.Get_default(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write(_env, "_lune.unwrapDefault( ")
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ", ")
            convLua_filter_7_(_env, _exp, self, &node.Nodes_Node)
            self.FP.Write(_env, ")")
        } else {
            self.FP.Write(_env, "_lune.unwrap( ")
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ")")
        }
    }
}

// 2712: decl @lune.@base.@convLua.ConvFilter.processExpCall
func (self *convLua_ConvFilter) ProcessExpCall(_env *LnsEnv, node *Nodes_ExpCallNode,_opt LnsAny) {
    var wroteFuncFlag bool
    wroteFuncFlag = false
    var setArgFlag bool
    setArgFlag = false
    var fieldCall func(_env *LnsEnv) bool
    fieldCall = func(_env *LnsEnv) bool {
        var fieldNode *Nodes_RefFieldNode
        
        {
            _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func(_env).FP)
            if _fieldNode == nil{
                return true
            } else {
                fieldNode = _fieldNode.(*Nodes_RefFieldNode)
            }
        }
        if node.FP.Get_func(_env).FP.Get_expType(_env) == self.builtinFunc.G__lns_runtime_log{
            return false
        }
        if node.FP.Get_func(_env).FP.Get_expType(_env) == self.builtinFunc.List___new{
            self.FP.Write(_env, "{}")
            return false
        }
        var prefixNode *Nodes_Node
        prefixNode = fieldNode.FP.Get_prefix(_env)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( prefixNode.FP.Get_expType(_env).FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeAsyncItem, nil)) &&
            _env.SetStackVal( node.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_rawTxt(_env) == "_createPipe") ).(bool)){
            self.FP.Write(_env, "nil")
            return false
        }
        var processSet func(_env *LnsEnv) bool
        processSet = func(_env *LnsEnv) bool {
            setArgFlag = true
            
            wroteFuncFlag = true
            
            if _switch1 := fieldNode.FP.Get_field(_env).Txt; _switch1 == "add" || _switch1 == "del" {
                convLua_filter_7_(_env, prefixNode, self, &fieldNode.Nodes_Node)
                self.FP.Write(_env, "[")
                {
                    _argList := node.FP.Get_argList(_env)
                    if !Lns_IsNil( _argList ) {
                        argList := _argList.(*Nodes_ExpListNode)
                        convLua_filter_7_(_env, &argList.Nodes_Node, self, &fieldNode.Nodes_Node)
                    }
                }
                self.FP.Write(_env, "]")
                if _switch2 := fieldNode.FP.Get_field(_env).Txt; _switch2 == "add" {
                    self.FP.Write(_env, "= true")
                } else if _switch2 == "del" {
                    self.FP.Write(_env, "= nil")
                }
                return false
            }
            self.FP.Write(_env, _env.LuaVM.String_format("_lune._Set_%s(", []LnsAny{fieldNode.FP.Get_field(_env).Txt}))
            convLua_filter_7_(_env, prefixNode, self, &fieldNode.Nodes_Node)
            return true
        }
        var prefixType *Ast_TypeInfo
        prefixType = prefixNode.FP.Get_expType(_env)
        var processEnumAlge func(_env *LnsEnv)
        processEnumAlge = func(_env *LnsEnv) {
            wroteFuncFlag = true
            
            var fieldExpType *Ast_TypeInfo
            fieldExpType = fieldNode.FP.Get_expType(_env)
            var canonicalName string
            canonicalName = self.FP.getFullName(_env, prefixType)
            var methodName string
            methodName = fieldNode.FP.Get_field(_env).Txt
            var delimit string
            delimit = ":"
            if methodName == "get__txt"{
                methodName = "_getTxt"
                
            }
            if fieldExpType.FP.Get_kind(_env) == Ast_TypeInfoKind__Func{
                delimit = "."
                
            }
            self.FP.Write(_env, _env.LuaVM.String_format("%s%s%s( ", []LnsAny{canonicalName, delimit, methodName}))
            if fieldExpType.FP.Get_staticFlag(_env){
                setArgFlag = false
                
            } else { 
                convLua_filter_7_(_env, prefixNode, self, &fieldNode.Nodes_Node)
                setArgFlag = true
                
            }
        }
        if node.FP.Get_nilAccess(_env){
            wroteFuncFlag = true
            
            setArgFlag = true
            
            if _switch1 := prefixType.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__List || _switch1 == Ast_TypeInfoKind__Array {
                self.FP.Write(_env, _env.LuaVM.String_format("_lune.nilacc( table.%s, nil, 'list', ", []LnsAny{fieldNode.FP.Get_field(_env).Txt}))
                convLua_filter_7_(_env, prefixNode, self, &fieldNode.Nodes_Node)
            } else {
                self.FP.Write(_env, "_lune.nilacc( ")
                convLua_filter_7_(_env, prefixNode, self, &fieldNode.Nodes_Node)
                self.FP.Write(_env, _env.LuaVM.String_format(", '%s', 'callmtd' ", []LnsAny{fieldNode.FP.Get_field(_env).Txt}))
            }
        } else { 
            if _switch2 := prefixType.FP.Get_kind(_env); _switch2 == Ast_TypeInfoKind__List || _switch2 == Ast_TypeInfoKind__Array {
                setArgFlag = true
                
                wroteFuncFlag = true
                
                self.FP.Write(_env, _env.LuaVM.String_format("table.%s( ", []LnsAny{fieldNode.FP.Get_field(_env).Txt}))
                convLua_filter_7_(_env, prefixNode, self, &fieldNode.Nodes_Node)
            } else if _switch2 == Ast_TypeInfoKind__Set {
                if Lns_op_not(processSet(_env)){
                    return false
                }
            } else if _switch2 == Ast_TypeInfoKind__Enum || _switch2 == Ast_TypeInfoKind__Alge {
                processEnumAlge(_env)
            } else if _switch2 == Ast_TypeInfoKind__Box {
                convLua_filter_7_(_env, prefixNode, self, &fieldNode.Nodes_Node)
                self.FP.Write(_env, "[1]")
                return false
            } else if _switch2 == Ast_TypeInfoKind__Class {
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( prefixType.FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeMapping, nil)) &&
                    _env.SetStackVal( convLua_isGenericType_10_(_env, prefixType)) &&
                    _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( fieldNode.FP.Get_field(_env).Txt == "_fromMap") ||
                        _env.SetStackVal( fieldNode.FP.Get_field(_env).Txt == "_fromStem") ).(bool))) ).(bool)){
                    wroteFuncFlag = true
                    
                    setArgFlag = true
                    
                    convLua_filter_7_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
                    self.FP.Write(_env, "( ")
                    {
                        _argList := node.FP.Get_argList(_env)
                        if !Lns_IsNil( _argList ) {
                            argList := _argList.(*Nodes_ExpListNode)
                            convLua_filter_7_(_env, &argList.Nodes_Node, self, &node.Nodes_Node)
                            self.FP.Write(_env, ", ")
                        }
                    }
                    self.FP.outputAlter2MapFunc(_env, self.FP, prefixType.FP.CreateAlt2typeMap(_env, false))
                    self.FP.Write(_env, ")")
                    return false
                }
            }
        }
        return true
    }
    if Lns_op_not(fieldCall(_env)){
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
                self.FP.Write(_env, _env.LuaVM.String_format("%s.%s( self ", []LnsAny{self.FP.getFullName(_env, funcType.FP.Get_parentInfo(_env)), funcType.FP.Get_rawTxt(_env)}))
            } else { 
                if _switch1 := refNode.FP.Get_expType(_env); _switch1 == self.builtinFunc.Lns_expandLuavalMap {
                    wroteFuncFlag = true
                    
                    self.FP.Write(_env, "(")
                } else if _switch1 == self.builtinFunc.Lns___run {
                    self.FP.Write(_env, "_lune._run(")
                    wroteFuncFlag = true
                    
                } else if _switch1 == self.builtinFunc.Lns___join {
                    return 
                }
            }
        }
    }
    if Lns_op_not(wroteFuncFlag){
        if node.FP.Get_nilAccess(_env){
            self.FP.Write(_env, "_lune.nilacc( ")
            convLua_filter_7_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ", nil, 'call'")
            wroteFuncFlag = true
            
            setArgFlag = true
            
        } else { 
            convLua_filter_7_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, "( ")
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
            var expList *LnsList
            expList = NewLnsList([]LnsAny{})
            for _, _expNode := range( argList.FP.Get_expList(_env).Items ) {
                expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                if expNode.FP.Get_expType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Abbr{
                    {
                        _toDDD := Nodes_ExpToDDDNodeDownCastF(expNode.FP)
                        if !Lns_IsNil( _toDDD ) {
                            toDDD := _toDDD.(*Nodes_ExpToDDDNode)
                            for _, _appNode := range( toDDD.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
                                appNode := _appNode.(Nodes_NodeDownCast).ToNodes_Node()
                                expList.Insert(Nodes_Node2Stem(appNode))
                            }
                        } else {
                            expList.Insert(Nodes_Node2Stem(expNode))
                        }
                    }
                }
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( wroteFuncFlag) &&
                _env.SetStackVal( setArgFlag) ).(bool)){
                if expList.Len() > 0{
                    self.FP.Write(_env, ", ")
                }
            }
            if convStrFlag{
                var opList *LnsList
                opList = NewLnsList([]LnsAny{})
                if expList.Len() > 0{
                    var literal LnsAny
                    literal = convLua_convExp15065(Lns_2DDD(expList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node().FP.GetLiteral(_env)))
                    if literal != nil{
                        literal_1625 := literal
                        switch _matchExp1 := literal_1625.(type) {
                        case *Nodes_Literal__Str:
                        txt := _matchExp1.Val1
                            opList = TransUnit_findForm(_env, txt)
                            
                        }
                    }
                }
                for _index, _argNode := range( expList.Items ) {
                    index := _index + 1
                    argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                    var filtered bool
                    filtered = false
                    if index > 1{
                        self.FP.Write(_env, ", ")
                        if index - 1 <= opList.Len(){
                            var formType LnsInt
                            formType = convLua_convExp15141(Lns_2DDD(TransUnit_isMatchStringFormatType(_env, opList.GetAt(index - 1).(string), argNode.FP.Get_expType(_env), self.targetLuaVer)))
                            if formType == TransUnit_FormType__NeedConv{
                                self.FP.Write(_env, "tostring( ")
                                convLua_filter_7_(_env, argNode, self, &node.Nodes_Node)
                                self.FP.Write(_env, ")")
                                filtered = true
                                
                            }
                        }
                    }
                    if Lns_op_not(filtered){
                        convLua_filter_7_(_env, argNode, self, &node.Nodes_Node)
                    }
                }
            } else { 
                convLua_filter_7_(_env, &argList.Nodes_Node, self, &node.Nodes_Node)
            }
        }
    }
    self.FP.Write(_env, " )")
}

// 2970: decl @lune.@base.@convLua.ConvFilter.processExpList
func (self *convLua_ConvFilter) ProcessExpList(_env *LnsEnv, node *Nodes_ExpListNode,_opt LnsAny) {
    var expList *LnsList
    expList = node.FP.Get_expList(_env)
    self.FP.processExpListSub(_env, &node.Nodes_Node, expList, node.FP.Get_mRetExp(_env))
}

// 2979: decl @lune.@base.@convLua.ConvFilter.processExpOp1
func (self *convLua_ConvFilter) ProcessExpOp1(_env *LnsEnv, node *Nodes_ExpOp1Node,_opt LnsAny) {
    var op string
    op = node.FP.Get_op(_env).Txt
    if op == ",,,"{
        convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    } else if op == ",,,,"{
        if node.FP.Get_macroMode(_env) != Nodes_MacroMode__None{
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else { 
            self.FP.Write(_env, "__luneSym2Str( ")
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, " )")
        }
    } else if op == ",,"{
        if _switch1 := node.FP.Get_exp(_env).FP.Get_expType(_env); _switch1 == Ast_builtinTypeInt || _switch1 == Ast_builtinTypeReal || _switch1 == Ast_builtinTypeBool {
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else {
            self.FP.Write(_env, "__luneGetLocal( ")
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, " )")
        }
    } else if op == "~"{
        if self.targetLuaVer.FP.Get_hasBitOp(_env) == LuaVer_BitOp__HasOp{
            self.FP.Write(_env, op)
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else { 
            self.FP.Write(_env, "bit32.bnot( ")
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, " )")
        }
    } else { 
        if op == "not"{
            op = op + " "
            
        }
        self.FP.Write(_env, op)
        convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    }
}

// 3028: decl @lune.@base.@convLua.ConvFilter.processExpToDDD
func (self *convLua_ConvFilter) ProcessExpToDDD(_env *LnsEnv, node *Nodes_ExpToDDDNode,_opt LnsAny) {
    self.FP.processExpListSub(_env, &node.Nodes_Node, node.FP.Get_expList(_env).FP.Get_expList(_env), node.FP.Get_expList(_env).FP.Get_mRetExp(_env))
}

// 3034: decl @lune.@base.@convLua.ConvFilter.processExpMultiTo1
func (self *convLua_ConvFilter) ProcessExpMultiTo1(_env *LnsEnv, node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
}

// 3040: decl @lune.@base.@convLua.ConvFilter.processExpCast
func (self *convLua_ConvFilter) ProcessExpCast(_env *LnsEnv, node *Nodes_ExpCastNode,_opt LnsAny) {
    if _switch1 := node.FP.Get_castKind(_env); _switch1 == Nodes_CastKind__Force {
        if node.FP.Get_expType(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeInt, nil, nil){
            self.FP.Write(_env, "math.floor(")
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ")")
        } else if node.FP.Get_expType(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeReal, nil, nil){
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, " * 1.0")
        } else { 
            convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        }
    } else if _switch1 == Nodes_CastKind__Normal {
        self.FP.Write(_env, "_lune.__Cast( ")
        convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        var castKind LnsInt
        var classObj string
        classObj = "nil"
        if _switch2 := node.FP.Get_expType(_env).FP.Get_nonnilableType(_env); _switch2 == Ast_builtinTypeInt {
            castKind = LuaMod_CastKind__Int
            
        } else if _switch2 == Ast_builtinTypeReal {
            castKind = LuaMod_CastKind__Real
            
        } else if _switch2 == Ast_builtinTypeString {
            castKind = LuaMod_CastKind__Str
            
        } else {
            castKind = LuaMod_CastKind__Class
            
            classObj = self.FP.getFullName(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env))
            
        }
        self.FP.Write(_env, _env.LuaVM.String_format(", %d, %s )", []LnsAny{castKind, classObj}))
    } else if _switch1 == Nodes_CastKind__Implicit {
        convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    }
}

// 3087: decl @lune.@base.@convLua.ConvFilter.processExpParen
func (self *convLua_ConvFilter) ProcessExpParen(_env *LnsEnv, node *Nodes_ExpParenNode,_opt LnsAny) {
    self.FP.Write(_env, "(")
    convLua_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, " )")
}

// 3096: decl @lune.@base.@convLua.ConvFilter.processExpSetVal
func (self *convLua_ConvFilter) ProcessExpSetVal(_env *LnsEnv, node *Nodes_ExpSetValNode,_opt LnsAny) {
    convLua_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, " = ")
    convLua_filter_7_(_env, &node.FP.Get_exp2(_env).Nodes_Node, self, &node.Nodes_Node)
}

// 3103: decl @lune.@base.@convLua.ConvFilter.processExpSetItem
func (self *convLua_ConvFilter) ProcessExpSetItem(_env *LnsEnv, node *Nodes_ExpSetItemNode,_opt LnsAny) {
    convLua_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, "[")
    switch _matchExp1 := node.FP.Get_index(_env).(type) {
    case *Nodes_IndexVal__NodeIdx:
    index := _matchExp1.Val1
        convLua_filter_7_(_env, index, self, &node.Nodes_Node)
    case *Nodes_IndexVal__SymIdx:
    index := _matchExp1.Val1
        self.FP.Write(_env, _env.LuaVM.String_format("'%s'", []LnsAny{index}))
    }
    self.FP.Write(_env, "]")
    self.FP.Write(_env, " = ")
    convLua_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
}

// 3121: decl @lune.@base.@convLua.ConvFilter.processExpOp2
func (self *convLua_ConvFilter) ProcessExpOp2(_env *LnsEnv, node *Nodes_ExpOp2Node,_opt LnsAny) {
    var intCast bool
    intCast = false
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_expType(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeInt, nil, nil)) &&
        _env.SetStackVal( node.FP.Get_op(_env).Txt == "/") ).(bool)){
        intCast = true
        
        self.FP.Write(_env, "math.floor(")
    }
    var opTxt string
    opTxt = node.FP.Get_op(_env).Txt
    {
        __exp := Ast_bitBinOpMap.Get(opTxt)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(LnsInt)
            if self.targetLuaVer.FP.Get_hasBitOp(_env) == LuaVer_BitOp__HasOp{
                if _switch1 := _exp; _switch1 == Ast_BitOpKind__LShift {
                    opTxt = "<<"
                    
                } else if _switch1 == Ast_BitOpKind__RShift {
                    opTxt = ">>"
                    
                }
                convLua_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
                self.FP.Write(_env, " " + opTxt + " ")
                convLua_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
            } else { 
                var binfunc string
                binfunc = ""
                var exp2Mod string
                exp2Mod = ""
                if _switch2 := _exp; _switch2 == Ast_BitOpKind__And {
                    binfunc = "band"
                    
                } else if _switch2 == Ast_BitOpKind__Or {
                    binfunc = "bor"
                    
                } else if _switch2 == Ast_BitOpKind__Xor {
                    binfunc = "bxor"
                    
                } else if _switch2 == Ast_BitOpKind__LShift {
                    binfunc = "lshift"
                    
                } else if _switch2 == Ast_BitOpKind__RShift {
                    binfunc = "lshift"
                    
                    exp2Mod = "-"
                    
                }
                self.FP.Write(_env, _env.LuaVM.String_format("bit32.%s(", []LnsAny{binfunc}))
                convLua_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
                self.FP.Write(_env, ", ")
                self.FP.Write(_env, exp2Mod)
                convLua_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
                self.FP.Write(_env, " )")
            }
        } else {
            convLua_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, " " + opTxt + " ")
            convLua_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
        }
    }
    if intCast{
        self.FP.Write(_env, ")")
    }
}

// 3191: decl @lune.@base.@convLua.ConvFilter.processExpRef
func (self *convLua_ConvFilter) ProcessExpRef(_env *LnsEnv, node *Nodes_ExpRefNode,_opt LnsAny) {
    if _switch1 := node.FP.Get_symbolInfo(_env).FP.Get_name(_env); _switch1 == "super" {
        var funcType *Ast_TypeInfo
        funcType = node.FP.Get_expType(_env)
        self.FP.Write(_env, _env.LuaVM.String_format("%s.%s", []LnsAny{self.FP.getFullName(_env, funcType.FP.Get_parentInfo(_env)), funcType.FP.Get_rawTxt(_env)}))
    } else {
        var builtinFunc *Builtin_BuiltinFuncType
        builtinFunc = self.builtinFunc
        if node.FP.Get_expType(_env).FP.Equals(_env, self.processInfo, builtinFunc.Lns__load, nil, nil){
            self.FP.Write(_env, "_lune." + self.targetLuaVer.FP.Get_loadStrFuncName(_env))
        } else { 
            if self.macroVarSymSet.Has(Ast_SymbolInfo2Stem(node.FP.Get_symbolInfo(_env).FP.GetOrg(_env))){
                self.FP.Write(_env, "macroVar.")
            } else { 
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( node.FP.Get_symbolInfo(_env).FP.Get_accessMode(_env) == Ast_AccessMode__Pub) &&
                    _env.SetStackVal( node.FP.Get_symbolInfo(_env).FP.Get_kind(_env) == Ast_SymbolKind__Var) ).(bool)){
                    if self.needModuleObj{
                        self.FP.Write(_env, "_moduleObj.")
                    }
                }
            }
            self.FP.Write(_env, node.FP.Get_symbolInfo(_env).FP.Get_name(_env))
            if node.FP.Get_symbolInfo(_env).FP.Get_isLazyLoad(_env){
                self.FP.Write(_env, "()")
            }
        }
    }
}

// 3227: decl @lune.@base.@convLua.ConvFilter.processExpRefItem
func (self *convLua_ConvFilter) ProcessExpRefItem(_env *LnsEnv, node *Nodes_ExpRefItemNode,_opt LnsAny) {
    if node.FP.Get_nilAccess(_env){
        self.FP.Write(_env, "_lune.nilacc( ")
        convLua_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, ", nil, 'item', ")
        {
            __exp := node.FP.Get_index(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_Node)
                convLua_filter_7_(_env, _exp, self, &node.Nodes_Node)
            } else {
                self.FP.Write(_env, _env.LuaVM.String_format("'%s'", []LnsAny{Lns_unwrap( node.FP.Get_symbol(_env)).(string)}))
            }
        }
        self.FP.Write(_env, ")")
    } else { 
        if node.FP.Get_val(_env).FP.Get_expType(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeString, nil, nil){
            self.FP.Write(_env, "string.byte( ")
            convLua_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ", ")
            {
                __exp := node.FP.Get_index(_env)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Nodes_Node)
                    convLua_filter_7_(_env, _exp, self, &node.Nodes_Node)
                } else {
                    panic("index is nil")
                }
            }
            self.FP.Write(_env, " )")
        } else { 
            convLua_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, "[")
            {
                __exp := node.FP.Get_index(_env)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Nodes_Node)
                    convLua_filter_7_(_env, _exp, self, &node.Nodes_Node)
                } else {
                    self.FP.Write(_env, _env.LuaVM.String_format("'%s'", []LnsAny{Lns_unwrap( node.FP.Get_symbol(_env)).(string)}))
                }
            }
            self.FP.Write(_env, "]")
        }
    }
}

// 3269: decl @lune.@base.@convLua.ConvFilter.processRefField
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
                    self.FP.Write(_env, code)
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
        self.FP.Write(_env, "_lune.nilacc( ")
        convLua_filter_7_(_env, prefix, self, &node.Nodes_Node)
        self.FP.Write(_env, _env.LuaVM.String_format(", \"%s\" )", []LnsAny{node.FP.Get_field(_env).Txt}))
    } else { 
        convLua_filter_7_(_env, prefix, self, &node.Nodes_Node)
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
        self.FP.Write(_env, delimit + fieldToken.Txt)
        {
            _symbolInfo := node.FP.Get_symbolInfo(_env)
            if !Lns_IsNil( _symbolInfo ) {
                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                if symbolInfo.FP.Get_isLazyLoad(_env){
                    self.FP.Write(_env, "()")
                }
            }
        }
    }
}

// 3308: decl @lune.@base.@convLua.ConvFilter.processExpOmitEnum
func (self *convLua_ConvFilter) ProcessExpOmitEnum(_env *LnsEnv, node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    {
        _aliasType := node.FP.Get_aliasType(_env)
        if !Lns_IsNil( _aliasType ) {
            aliasType := _aliasType.(*Ast_AliasTypeInfo)
            self.FP.Write(_env, self.FP.getFullName(_env, &aliasType.Ast_TypeInfo))
        } else {
            self.FP.Write(_env, self.FP.getFullName(_env, node.FP.Get_expType(_env)))
        }
    }
    self.FP.Write(_env, _env.LuaVM.String_format(".%s", []LnsAny{node.FP.Get_valToken(_env).Txt}))
}

// 3319: decl @lune.@base.@convLua.ConvFilter.processGetField
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
        self.FP.Write(_env, _env.LuaVM.String_format("%s:_getTxt( ", []LnsAny{self.FP.getFullName(_env, prefixType)}))
        convLua_filter_7_(_env, prefixNode, self, &node.Nodes_Node)
        self.FP.writeln(_env, ")")
    } else { 
        if node.FP.Get_nilAccess(_env){
            fieldTxt = _env.LuaVM.String_format("get_%s", []LnsAny{fieldTxt})
            
            self.FP.Write(_env, "_lune.nilacc( ")
            convLua_filter_7_(_env, prefixNode, self, &node.Nodes_Node)
            self.FP.Write(_env, _env.LuaVM.String_format(", '%s', 'callmtd' )", []LnsAny{fieldTxt}))
        } else { 
            fieldTxt = _env.LuaVM.String_format("get_%s()", []LnsAny{fieldTxt})
            
            convLua_filter_7_(_env, prefixNode, self, &node.Nodes_Node)
            var delimit string
            delimit = "."
            if node.FP.Get_getterTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
                delimit = ":"
                
            } else { 
                delimit = "."
                
            }
            self.FP.Write(_env, delimit + fieldTxt)
        }
    }
}

// 3358: decl @lune.@base.@convLua.ConvFilter.processReturn
func (self *convLua_ConvFilter) ProcessReturn(_env *LnsEnv, node *Nodes_ReturnNode,_opt LnsAny) {
    self.FP.Write(_env, "return ")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
}

// 3368: decl @lune.@base.@convLua.ConvFilter.processLuneKind
func (self *convLua_ConvFilter) ProcessLuneKind(_env *LnsEnv, node *Nodes_LuneKindNode,_opt LnsAny) {
    {
        _workNode := Nodes_ExpCastNodeDownCastF(node.FP.Get_exp(_env).FP)
        if !Lns_IsNil( _workNode ) {
            workNode := _workNode.(*Nodes_ExpCastNode)
            if workNode.FP.Get_castKind(_env) == Nodes_CastKind__Implicit{
                self.FP.Write(_env, _env.LuaVM.String_format("%d", []LnsAny{workNode.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env)}))
            }
        } else {
            self.FP.Write(_env, _env.LuaVM.String_format("%d", []LnsAny{node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env)}))
        }
    }
}

// 3381: decl @lune.@base.@convLua.ConvFilter.processTestCase
func (self *convLua_ConvFilter) ProcessTestCase(_env *LnsEnv, node *Nodes_TestCaseNode,_opt LnsAny) {
    if self.enableTest{
        self.FP.writeln(_env, "do")
        self.FP.pushIndent(_env, nil)
        convLua_filter_7_(_env, node.FP.Get_impNode(_env), self, &node.Nodes_Node)
        self.FP.writeln(_env, "")
        self.FP.writeln(_env, _env.LuaVM.String_format("local function testcase( %s ) ", []LnsAny{node.FP.Get_ctrlName(_env)}))
        convLua_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        self.FP.writeln(_env, "end")
        self.FP.writeln(_env, _env.LuaVM.String_format("__t.registerTestcase( \"%s\", \"%s\", testcase )", []LnsAny{Lns_car(_env.LuaVM.String_gsub(self.moduleTypeInfo.FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), nil),"@", "")).(string), node.FP.Get_name(_env).Txt}))
        self.FP.popIndent(_env)
        self.FP.writeln(_env, "end")
    }
}

// 3403: decl @lune.@base.@convLua.ConvFilter.processTestBlock
func (self *convLua_ConvFilter) ProcessTestBlock(_env *LnsEnv, node *Nodes_TestBlockNode,_opt LnsAny) {
    if self.enableTest{
        for _, _statement := range( node.FP.Get_stmtList(_env).Items ) {
            statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
            convLua_filter_7_(_env, statement, self, &node.Nodes_Node)
            self.FP.writeln(_env, "")
        }
    }
}

// 3414: decl @lune.@base.@convLua.ConvFilter.processProvide
func (self *convLua_ConvFilter) ProcessProvide(_env *LnsEnv, node *Nodes_ProvideNode,_opt LnsAny) {
}

// 3419: decl @lune.@base.@convLua.ConvFilter.processAlias
func (self *convLua_ConvFilter) ProcessAlias(_env *LnsEnv, node *Nodes_AliasNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("local %s = ", []LnsAny{node.FP.Get_newSymbol(_env).FP.Get_name(_env)}))
    convLua_filter_7_(_env, node.FP.Get_srcNode(_env), self, &node.Nodes_Node)
    if Ast_isPubToExternal(_env, node.FP.Get_expType(_env).FP.Get_accessMode(_env)){
        self.FP.Write(_env, _env.LuaVM.String_format("\n_moduleObj.%s = %s", []LnsAny{node.FP.Get_newSymbol(_env).FP.Get_name(_env), node.FP.Get_newSymbol(_env).FP.Get_name(_env)}))
    }
}

// 3429: decl @lune.@base.@convLua.ConvFilter.processBoxing
func (self *convLua_ConvFilter) ProcessBoxing(_env *LnsEnv, node *Nodes_BoxingNode,_opt LnsAny) {
    self.FP.Write(_env, "{")
    convLua_filter_7_(_env, node.FP.Get_src(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, "}")
}

// 3439: decl @lune.@base.@convLua.ConvFilter.processUnboxing
func (self *convLua_ConvFilter) ProcessUnboxing(_env *LnsEnv, node *Nodes_UnboxingNode,_opt LnsAny) {
    convLua_filter_7_(_env, node.FP.Get_src(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, "[1]")
}

// 3447: decl @lune.@base.@convLua.ConvFilter.processLiteralList
func (self *convLua_ConvFilter) ProcessLiteralList(_env *LnsEnv, node *Nodes_LiteralListNode,_opt LnsAny) {
    self.FP.Write(_env, "{")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write(_env, "}")
}

// 3459: decl @lune.@base.@convLua.ConvFilter.processLiteralSet
func (self *convLua_ConvFilter) ProcessLiteralSet(_env *LnsEnv, node *Nodes_LiteralSetNode,_opt LnsAny) {
    self.FP.Write(_env, "{")
    {
        _expListNode := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expListNode ) {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            for _index, _expNode := range( expListNode.FP.Get_expList(_env).Items ) {
                index := _index + 1
                expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.Write(_env, ", ")
                }
                self.FP.Write(_env, "[")
                convLua_filter_7_(_env, expNode, self, &node.Nodes_Node)
                self.FP.Write(_env, "] = true")
            }
        }
    }
    self.FP.Write(_env, "}")
}

// 3477: decl @lune.@base.@convLua.ConvFilter.processLiteralMap
func (self *convLua_ConvFilter) ProcessLiteralMap(_env *LnsEnv, node *Nodes_LiteralMapNode,_opt LnsAny) {
    self.FP.Write(_env, "{")
    var pairList *LnsList
    pairList = node.FP.Get_pairList(_env)
    for _index, _pair := range( pairList.Items ) {
        index := _index + 1
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        self.FP.Write(_env, "[")
        convLua_filter_7_(_env, pair.FP.Get_key(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, "] = ")
        convLua_filter_7_(_env, pair.FP.Get_val(_env), self, &node.Nodes_Node)
    }
    self.FP.Write(_env, "}")
}

// 3495: decl @lune.@base.@convLua.ConvFilter.processLiteralArray
func (self *convLua_ConvFilter) ProcessLiteralArray(_env *LnsEnv, node *Nodes_LiteralArrayNode,_opt LnsAny) {
    self.FP.Write(_env, "{")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_7_(_env, &_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write(_env, "}")
}

// 3508: decl @lune.@base.@convLua.ConvFilter.processLiteralChar
func (self *convLua_ConvFilter) ProcessLiteralChar(_env *LnsEnv, node *Nodes_LiteralCharNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("%d", []LnsAny{node.FP.Get_num(_env)}))
}

// 3514: decl @lune.@base.@convLua.ConvFilter.processLiteralInt
func (self *convLua_ConvFilter) ProcessLiteralInt(_env *LnsEnv, node *Nodes_LiteralIntNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
}

// 3522: decl @lune.@base.@convLua.ConvFilter.processLiteralReal
func (self *convLua_ConvFilter) ProcessLiteralReal(_env *LnsEnv, node *Nodes_LiteralRealNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
}

// 3529: decl @lune.@base.@convLua.ConvFilter.processLiteralString
func (self *convLua_ConvFilter) ProcessLiteralString(_env *LnsEnv, node *Nodes_LiteralStringNode,_opt LnsAny) {
    var txt string
    txt = node.FP.Get_token(_env).Txt
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(txt, "^```", nil, nil))){
        txt = "[==[" + _env.LuaVM.String_sub(txt,4, -4) + "]==]"
        
    }
    var opList *LnsList
    opList = TransUnit_findForm(_env, txt)
    {
        _expList := node.FP.Get_orgParam(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            var mRetIndex LnsAny
            mRetIndex = _env.NilAccFin(_env.NilAccPush(expList.FP.Get_mRetExp(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
            self.FP.Write(_env, _env.LuaVM.String_format("string.format( %s, ", []LnsAny{txt}))
            for _index, _val := range( expList.FP.Get_expList(_env).Items ) {
                index := _index + 1
                val := _val.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.Write(_env, ", ")
                }
                var matchFlag LnsInt
                matchFlag = TransUnit_FormType__Match
                if index <= opList.Len(){
                    matchFlag = convLua_convExp17910(Lns_2DDD(TransUnit_isMatchStringFormatType(_env, opList.GetAt(index).(string), val.FP.Get_expType(_env), self.targetLuaVer)))
                    
                }
                if matchFlag == TransUnit_FormType__NeedConv{
                    self.FP.Write(_env, "tostring( ")
                    convLua_filter_7_(_env, val, self, &node.Nodes_Node)
                    self.FP.Write(_env, ")")
                } else { 
                    convLua_filter_7_(_env, val, self, &node.Nodes_Node)
                }
                if index == mRetIndex{
                    break
                }
            }
            self.FP.Write(_env, ")")
        } else {
            self.FP.Write(_env, txt)
        }
    }
}

// 3571: decl @lune.@base.@convLua.ConvFilter.processLiteralBool
func (self *convLua_ConvFilter) ProcessLiteralBool(_env *LnsEnv, node *Nodes_LiteralBoolNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
}

// 3577: decl @lune.@base.@convLua.ConvFilter.processLiteralNil
func (self *convLua_ConvFilter) ProcessLiteralNil(_env *LnsEnv, node *Nodes_LiteralNilNode,_opt LnsAny) {
    self.FP.Write(_env, "nil")
}

// 3583: decl @lune.@base.@convLua.ConvFilter.processBreak
func (self *convLua_ConvFilter) ProcessBreak(_env *LnsEnv, node *Nodes_BreakNode,_opt LnsAny) {
    self.FP.Write(_env, "break")
}

// 3589: decl @lune.@base.@convLua.ConvFilter.processLiteralSymbol
func (self *convLua_ConvFilter) ProcessLiteralSymbol(_env *LnsEnv, node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{node.FP.Get_token(_env).Txt}))
}

// 3595: decl @lune.@base.@convLua.ConvFilter.processLuneControl
func (self *convLua_ConvFilter) ProcessLuneControl(_env *LnsEnv, node *Nodes_LuneControlNode,_opt LnsAny) {
    switch node.FP.Get_pragma(_env).(type) {
    case *LuneControl_Pragma__load__lune_module:
        self.FP.processLoadRuntime(_env)
    }
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
// 3608: decl @lune.@base.@convLua.FilterInfo.outputLuaAndMeta
func (self *ConvLua_FilterInfo) OutputLuaAndMeta(_env *LnsEnv, node *Nodes_RootNode) {
    node.FP.ProcessFilter(_env, self.filter, ConvLua_Opt2Stem(NewConvLua_Opt(_env, &node.Nodes_Node)))
}

// 3615: decl @lune.@base.@convLua.FilterInfo.outputLua
func (self *ConvLua_FilterInfo) OutputLua(_env *LnsEnv, node *Nodes_RootNode) {
    node.FP.ProcessFilter(_env, self.filter, ConvLua_Opt2Stem(NewConvLua_Opt(_env, &node.Nodes_Node)))
}

// 3620: decl @lune.@base.@convLua.FilterInfo.outputMeta
func (self *ConvLua_FilterInfo) OutputMeta(_env *LnsEnv, node *Nodes_RootNode) {
}


// declaration Class -- MacroEvalImp
type ConvLua_MacroEvalImpMtd interface {
    EvalFromCodeToLuaCode(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
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
// 3650: decl @lune.@base.@convLua.MacroEvalImp.evalFromCodeToLuaCode
func (self *ConvLua_MacroEvalImp) EvalFromCodeToLuaCode(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,argNameList *LnsList,code LnsAny) string {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var conv *convLua_ConvFilter
    conv = NewconvLua_ConvFilter(_env, "macro", stream.FP, NewUtil_NullOStream(_env).FP, ConvLua_ConvMode__ConvMeta, true, Ast_headTypeInfo, processInfo, Ast_SymbolKind__Typ, self.builtinFunc, nil, LuaVer_getCurVer(_env), false, true)
    conv.FP.OutputDeclMacro(_env, name, argNameList, convLua_outputMacroStmtBlock_11_(func(_env *LnsEnv) {
        if code != nil{
            code_1863 := code.(string)
            conv.FP.Write(_env, code_1863)
        }
    }))
    return stream.FP.Get_txt(_env)
}

// 3669: decl @lune.@base.@convLua.MacroEvalImp.evalToLuaCode
func (self *ConvLua_MacroEvalImp) EvalToLuaCode(_env *LnsEnv, processInfo *Ast_ProcessInfo,node *Nodes_DeclMacroNode) string {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var conv *convLua_ConvFilter
    conv = NewconvLua_ConvFilter(_env, "macro", stream.FP, NewUtil_NullOStream(_env).FP, ConvLua_ConvMode__ConvMeta, true, Ast_headTypeInfo, processInfo, Ast_SymbolKind__Typ, self.builtinFunc, nil, LuaVer_getCurVer(_env), false, true)
    conv.FP.ProcessDeclMacro(_env, node, NewConvLua_Opt(_env, &node.Nodes_Node))
    return stream.FP.Get_txt(_env)
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
    Lns_Parser_init(_env)
    Lns_Types_init(_env)
    Lns_Log_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Option_init(_env)
    Lns_frontInterface_init(_env)
    Lns_Builtin_init(_env)
    convLua_stepIndent = 3
    
    
    
    Lns_Testing_init(_env)
}
func init() {
    init_convLua = false
}
