// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_OutputDepend bool
var OutputDepend__mod__ string
// for 146
func OutputDepend_convExp0_610(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 144
func OutputDepend_convExp0_569(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 96: decl @lune.@base.@OutputDepend.createFilter
func OutputDepend_createFilter(_env *LnsEnv, stream Lns_oStream) *Nodes_Filter {
    return &NewOutputDepend_convFilter(_env, stream).Nodes_Filter
}


// 45: decl @lune.@base.@OutputDepend.DependInfo.addImpotModule
func (self *OutputDepend_DependInfo) AddImpotModule(_env *LnsEnv, mod string) {
    self.importModuleList.Insert(mod)
}
// 48: decl @lune.@base.@OutputDepend.DependInfo.addSubMod
func (self *OutputDepend_DependInfo) AddSubMod(_env *LnsEnv, path string) {
    self.subModList.Insert(path)
}
// 53: decl @lune.@base.@OutputDepend.DependInfo.output
func (self *OutputDepend_DependInfo) Output(_env *LnsEnv, stream Lns_oStream) {
    stream.Write(_env, _env.GetVM().String_format("%s.meta: \\\n", Lns_2DDD(Lns_car(_env.GetVM().String_gsub(self.targetModule,"%.", "/")).(string))))
    stream.Write(_env, _env.GetVM().String_format("  %s.lns \\\n", Lns_2DDD(Lns_car(_env.GetVM().String_gsub(self.targetModule,"%.", "/")).(string))))
    for _, _mod := range( self.importModuleList.Items ) {
        mod := _mod
        stream.Write(_env, _env.GetVM().String_format("  %s.meta \\\n", Lns_2DDD(Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string))))
    }
    for _, _path := range( self.subModList.Items ) {
        path := _path
        stream.Write(_env, _env.GetVM().String_format("  %s.lns \\\n", Lns_2DDD(Lns_car(_env.GetVM().String_gsub(path,"%.", "/")).(string))))
    }
}
// 77: decl @lune.@base.@OutputDepend.convFilter.processRoot
func (self *OutputDepend_convFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_dummy LnsAny) {
    var moduleFull string
    moduleFull = node.FP.Get_moduleTypeInfo(_env).FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), Ast_DummyModuleInfoManager_get_instance(_env).FP, nil)
    var dependInfo *OutputDepend_DependInfo
    dependInfo = NewOutputDepend_DependInfo(_env, moduleFull)
    for _, _impNode := range( node.FP.Get_nodeManager(_env).FP.GetImportNodeList(_env).Items ) {
        impNode := _impNode.(Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        dependInfo.FP.AddImpotModule(_env, impNode.FP.Get_info(_env).FP.Get_modulePath(_env))
    }
    for _, _subfileNode := range( node.FP.Get_nodeManager(_env).FP.GetSubfileNodeList(_env).Items ) {
        subfileNode := _subfileNode.(Nodes_SubfileNodeDownCast).ToNodes_SubfileNode()
        {
            _usePath := subfileNode.FP.Get_usePath(_env)
            if !Lns_IsNil( _usePath ) {
                usePath := _usePath.(string)
                dependInfo.FP.AddSubMod(_env, usePath)
            }
        }
    }
    dependInfo.FP.Output(_env, self.stream)
}
// declaration Class -- DependInfo
type OutputDepend_DependInfoMtd interface {
    AddImpotModule(_env *LnsEnv, arg1 string)
    AddSubMod(_env *LnsEnv, arg1 string)
    Output(_env *LnsEnv, arg1 Lns_oStream)
}
type OutputDepend_DependInfo struct {
    targetModule string
    importModuleList *LnsList2_[string]
    subModList *LnsList2_[string]
    FP OutputDepend_DependInfoMtd
}
func OutputDepend_DependInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*OutputDepend_DependInfo).FP
}
func OutputDepend_DependInfo_toSlice(slice []LnsAny) []*OutputDepend_DependInfo {
    ret := make([]*OutputDepend_DependInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(OutputDepend_DependInfoDownCast).ToOutputDepend_DependInfo()
    }
    return ret
}
type OutputDepend_DependInfoDownCast interface {
    ToOutputDepend_DependInfo() *OutputDepend_DependInfo
}
func OutputDepend_DependInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(OutputDepend_DependInfoDownCast)
    if ok { return work.ToOutputDepend_DependInfo() }
    return nil
}
func (obj *OutputDepend_DependInfo) ToOutputDepend_DependInfo() *OutputDepend_DependInfo {
    return obj
}
func NewOutputDepend_DependInfo(_env *LnsEnv, arg1 string) *OutputDepend_DependInfo {
    obj := &OutputDepend_DependInfo{}
    obj.FP = obj
    obj.InitOutputDepend_DependInfo(_env, arg1)
    return obj
}
// 39: DeclConstr
func (self *OutputDepend_DependInfo) InitOutputDepend_DependInfo(_env *LnsEnv, targetModule string) {
    self.targetModule = Ast_TypeInfo_getModulePath(_env, targetModule)
    self.importModuleList = NewLnsList2_[string]([]string{})
    self.subModList = NewLnsList2_[string]([]string{})
}


// declaration Class -- convFilter
type OutputDepend_convFilterMtd interface {
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    GetFull(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl
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
}
type OutputDepend_convFilter struct {
    Nodes_Filter
    stream Lns_oStream
    FP OutputDepend_convFilterMtd
}
func OutputDepend_convFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*OutputDepend_convFilter).FP
}
func OutputDepend_convFilter_toSlice(slice []LnsAny) []*OutputDepend_convFilter {
    ret := make([]*OutputDepend_convFilter, len(slice))
    for index, val := range slice {
        ret[index] = val.(OutputDepend_convFilterDownCast).ToOutputDepend_convFilter()
    }
    return ret
}
type OutputDepend_convFilterDownCast interface {
    ToOutputDepend_convFilter() *OutputDepend_convFilter
}
func OutputDepend_convFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(OutputDepend_convFilterDownCast)
    if ok { return work.ToOutputDepend_convFilter() }
    return nil
}
func (obj *OutputDepend_convFilter) ToOutputDepend_convFilter() *OutputDepend_convFilter {
    return obj
}
func NewOutputDepend_convFilter(_env *LnsEnv, arg1 Lns_oStream) *OutputDepend_convFilter {
    obj := &OutputDepend_convFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitOutputDepend_convFilter(_env, arg1)
    return obj
}
// 70: DeclConstr
func (self *OutputDepend_convFilter) InitOutputDepend_convFilter(_env *LnsEnv, stream Lns_oStream) {
    self.InitNodes_Filter(_env, false, nil, nil)
    self.stream = stream
}



func Lns_OutputDepend_init(_env *LnsEnv) {
    if init_OutputDepend { return }
    init_OutputDepend = true
    OutputDepend__mod__ = "@lune.@base.@OutputDepend"
    Lns_InitMod()
    Lns_Nodes_init(_env)
    Lns_Util_init(_env)
    Lns_Ast_init(_env)
    Lns_AstInfo_init(_env)
}
func init() {
    init_OutputDepend = false
}
