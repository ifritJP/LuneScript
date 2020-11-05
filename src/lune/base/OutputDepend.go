// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_OutputDepend bool
var OutputDepend__mod__ string
// for 125
func OutputDepend_convExp487(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 92: decl @lune.@base.@OutputDepend.createFilter
func OutputDepend_createFilter(stream Lns_oStream) *Nodes_Filter {
    return &NewOutputDepend_convFilter(stream).Nodes_Filter
}

// 134: decl @lune.@base.@OutputDepend.Ast2Depend
func OutputDepend_Ast2Depend_1094_(ast *TransUnit_ASTInfo) string {
    var stream *Util_memStream
    stream = NewUtil_memStream()
    var conv *Nodes_Filter
    conv = OutputDepend_createFilter(stream.FP)
    ast.FP.Get_node().FP.ProcessFilter(conv, 0)
    return stream.FP.Get_txt()
}

// declaration Class -- DependInfo
type OutputDepend_DependInfoMtd interface {
    AddImpotModule(arg1 string)
    AddSubMod(arg1 string)
    Output(arg1 Lns_oStream)
}
type OutputDepend_DependInfo struct {
    targetModule string
    importModuleList *LnsList
    subModList *LnsList
    FP OutputDepend_DependInfoMtd
}
func OutputDepend_DependInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*OutputDepend_DependInfo).FP
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
func NewOutputDepend_DependInfo(arg1 string) *OutputDepend_DependInfo {
    obj := &OutputDepend_DependInfo{}
    obj.FP = obj
    obj.InitOutputDepend_DependInfo(arg1)
    return obj
}
// 35: DeclConstr
func (self *OutputDepend_DependInfo) InitOutputDepend_DependInfo(targetModule string) {
    self.targetModule = Ast_TypeInfo_getModulePath(targetModule)
    
    self.importModuleList = NewLnsList([]LnsAny{})
    
    self.subModList = NewLnsList([]LnsAny{})
    
}

// 41: decl @lune.@base.@OutputDepend.DependInfo.addImpotModule
func (self *OutputDepend_DependInfo) AddImpotModule(mod string) {
    self.importModuleList.Insert(mod)
}

// 44: decl @lune.@base.@OutputDepend.DependInfo.addSubMod
func (self *OutputDepend_DependInfo) AddSubMod(path string) {
    self.subModList.Insert(path)
}

// 49: decl @lune.@base.@OutputDepend.DependInfo.output
func (self *OutputDepend_DependInfo) Output(stream Lns_oStream) {
    stream.Write(Lns_getVM().String_format("%s.meta: \\\n", []LnsAny{Lns_car(Lns_getVM().String_gsub(self.targetModule,"%.", "/")).(string)}))
    stream.Write(Lns_getVM().String_format("  %s.lns \\\n", []LnsAny{Lns_car(Lns_getVM().String_gsub(self.targetModule,"%.", "/")).(string)}))
    for _, _mod := range( self.importModuleList.Items ) {
        mod := _mod.(string)
        stream.Write(Lns_getVM().String_format("  %s.meta \\\n", []LnsAny{Lns_car(Lns_getVM().String_gsub(mod,"%.", "/")).(string)}))
    }
    for _, _path := range( self.subModList.Items ) {
        path := _path.(string)
        stream.Write(Lns_getVM().String_format("  %s.lns \\\n", []LnsAny{Lns_car(Lns_getVM().String_gsub(path,"%.", "/")).(string)}))
    }
}


// declaration Class -- convFilter
type OutputDepend_convFilterMtd interface {
    DefaultProcess(arg1 *Nodes_Node, arg2 LnsAny)
    Get_moduleInfoManager() Ast_ModuleInfoManager
    Get_optStack() *LnsList
    Get_typeNameCtrl() *Ast_TypeNameCtrl
    ProcessAbbr(arg1 *Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(arg1 *Nodes_AliasNode, arg2 LnsAny)
    ProcessApply(arg1 *Nodes_ApplyNode, arg2 LnsAny)
    ProcessBlankLine(arg1 *Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(arg1 *Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(arg1 *Nodes_BreakNode, arg2 LnsAny)
    ProcessConvStat(arg1 *Nodes_ConvStatNode, arg2 LnsAny)
    ProcessDeclAdvertise(arg1 *Nodes_DeclAdvertiseNode, arg2 LnsAny)
    ProcessDeclAlge(arg1 *Nodes_DeclAlgeNode, arg2 LnsAny)
    ProcessDeclArg(arg1 *Nodes_DeclArgNode, arg2 LnsAny)
    ProcessDeclArgDDD(arg1 *Nodes_DeclArgDDDNode, arg2 LnsAny)
    ProcessDeclClass(arg1 *Nodes_DeclClassNode, arg2 LnsAny)
    ProcessDeclConstr(arg1 *Nodes_DeclConstrNode, arg2 LnsAny)
    ProcessDeclDestr(arg1 *Nodes_DeclDestrNode, arg2 LnsAny)
    ProcessDeclEnum(arg1 *Nodes_DeclEnumNode, arg2 LnsAny)
    ProcessDeclForm(arg1 *Nodes_DeclFormNode, arg2 LnsAny)
    ProcessDeclFunc(arg1 *Nodes_DeclFuncNode, arg2 LnsAny)
    ProcessDeclMacro(arg1 *Nodes_DeclMacroNode, arg2 LnsAny)
    ProcessDeclMember(arg1 *Nodes_DeclMemberNode, arg2 LnsAny)
    ProcessDeclMethod(arg1 *Nodes_DeclMethodNode, arg2 LnsAny)
    ProcessDeclVar(arg1 *Nodes_DeclVarNode, arg2 LnsAny)
    ProcessExpAccessMRet(arg1 *Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(arg1 *Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(arg1 *Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(arg1 *Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(arg1 *Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpList(arg1 *Nodes_ExpListNode, arg2 LnsAny)
    ProcessExpMRet(arg1 *Nodes_ExpMRetNode, arg2 LnsAny)
    ProcessExpMacroArgExp(arg1 *Nodes_ExpMacroArgExpNode, arg2 LnsAny)
    ProcessExpMacroExp(arg1 *Nodes_ExpMacroExpNode, arg2 LnsAny)
    ProcessExpMacroStat(arg1 *Nodes_ExpMacroStatNode, arg2 LnsAny)
    ProcessExpMacroStatList(arg1 *Nodes_ExpMacroStatListNode, arg2 LnsAny)
    ProcessExpMultiTo1(arg1 *Nodes_ExpMultiTo1Node, arg2 LnsAny)
    ProcessExpNew(arg1 *Nodes_ExpNewNode, arg2 LnsAny)
    ProcessExpOmitEnum(arg1 *Nodes_ExpOmitEnumNode, arg2 LnsAny)
    ProcessExpOp1(arg1 *Nodes_ExpOp1Node, arg2 LnsAny)
    ProcessExpOp2(arg1 *Nodes_ExpOp2Node, arg2 LnsAny)
    ProcessExpParen(arg1 *Nodes_ExpParenNode, arg2 LnsAny)
    ProcessExpRef(arg1 *Nodes_ExpRefNode, arg2 LnsAny)
    ProcessExpRefItem(arg1 *Nodes_ExpRefItemNode, arg2 LnsAny)
    ProcessExpSetItem(arg1 *Nodes_ExpSetItemNode, arg2 LnsAny)
    ProcessExpSetVal(arg1 *Nodes_ExpSetValNode, arg2 LnsAny)
    ProcessExpSubDDD(arg1 *Nodes_ExpSubDDDNode, arg2 LnsAny)
    ProcessExpToDDD(arg1 *Nodes_ExpToDDDNode, arg2 LnsAny)
    ProcessExpUnwrap(arg1 *Nodes_ExpUnwrapNode, arg2 LnsAny)
    ProcessFor(arg1 *Nodes_ForNode, arg2 LnsAny)
    ProcessForeach(arg1 *Nodes_ForeachNode, arg2 LnsAny)
    ProcessForsort(arg1 *Nodes_ForsortNode, arg2 LnsAny)
    ProcessGetField(arg1 *Nodes_GetFieldNode, arg2 LnsAny)
    ProcessIf(arg1 *Nodes_IfNode, arg2 LnsAny)
    ProcessIfUnwrap(arg1 *Nodes_IfUnwrapNode, arg2 LnsAny)
    ProcessImport(arg1 *Nodes_ImportNode, arg2 LnsAny)
    ProcessLiteralArray(arg1 *Nodes_LiteralArrayNode, arg2 LnsAny)
    ProcessLiteralBool(arg1 *Nodes_LiteralBoolNode, arg2 LnsAny)
    ProcessLiteralChar(arg1 *Nodes_LiteralCharNode, arg2 LnsAny)
    ProcessLiteralInt(arg1 *Nodes_LiteralIntNode, arg2 LnsAny)
    ProcessLiteralList(arg1 *Nodes_LiteralListNode, arg2 LnsAny)
    ProcessLiteralMap(arg1 *Nodes_LiteralMapNode, arg2 LnsAny)
    ProcessLiteralNil(arg1 *Nodes_LiteralNilNode, arg2 LnsAny)
    ProcessLiteralReal(arg1 *Nodes_LiteralRealNode, arg2 LnsAny)
    ProcessLiteralSet(arg1 *Nodes_LiteralSetNode, arg2 LnsAny)
    ProcessLiteralString(arg1 *Nodes_LiteralStringNode, arg2 LnsAny)
    ProcessLiteralSymbol(arg1 *Nodes_LiteralSymbolNode, arg2 LnsAny)
    ProcessLuneControl(arg1 *Nodes_LuneControlNode, arg2 LnsAny)
    ProcessLuneKind(arg1 *Nodes_LuneKindNode, arg2 LnsAny)
    ProcessMatch(arg1 *Nodes_MatchNode, arg2 LnsAny)
    ProcessNewAlgeVal(arg1 *Nodes_NewAlgeValNode, arg2 LnsAny)
    ProcessNone(arg1 *Nodes_NoneNode, arg2 LnsAny)
    ProcessProtoClass(arg1 *Nodes_ProtoClassNode, arg2 LnsAny)
    ProcessProtoMethod(arg1 *Nodes_ProtoMethodNode, arg2 LnsAny)
    ProcessProvide(arg1 *Nodes_ProvideNode, arg2 LnsAny)
    ProcessRefField(arg1 *Nodes_RefFieldNode, arg2 LnsAny)
    ProcessRefType(arg1 *Nodes_RefTypeNode, arg2 LnsAny)
    ProcessRepeat(arg1 *Nodes_RepeatNode, arg2 LnsAny)
    ProcessReturn(arg1 *Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(arg1 *Nodes_RootNode, arg2 LnsAny)
    ProcessScope(arg1 *Nodes_ScopeNode, arg2 LnsAny)
    ProcessStmtExp(arg1 *Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(arg1 *Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(arg1 *Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(arg1 *Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(arg1 *Nodes_TestCaseNode, arg2 LnsAny)
    ProcessUnboxing(arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(arg1 *Nodes_WhileNode, arg2 LnsAny)
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
func NewOutputDepend_convFilter(arg1 Lns_oStream) *OutputDepend_convFilter {
    obj := &OutputDepend_convFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitOutputDepend_convFilter(arg1)
    return obj
}
// 66: DeclConstr
func (self *OutputDepend_convFilter) InitOutputDepend_convFilter(stream Lns_oStream) {
    self.InitNodes_Filter(false, nil, nil)
    self.stream = stream
    
}

// 73: decl @lune.@base.@OutputDepend.convFilter.processRoot
func (self *OutputDepend_convFilter) ProcessRoot(node *Nodes_RootNode,_dummy LnsAny) {
    var moduleFull string
    moduleFull = node.FP.Get_moduleTypeInfo().FP.GetFullName(self.FP.Get_typeNameCtrl(), Ast_DummyModuleInfoManager_get_instance().FP, nil)
    var dependInfo *OutputDepend_DependInfo
    dependInfo = NewOutputDepend_DependInfo(moduleFull)
    for _, _impNode := range( node.FP.Get_nodeManager().FP.GetImportNodeList().Items ) {
        impNode := _impNode.(Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        dependInfo.FP.AddImpotModule(impNode.FP.Get_modulePath())
    }
    for _, _subfileNode := range( node.FP.Get_nodeManager().FP.GetSubfileNodeList().Items ) {
        subfileNode := _subfileNode.(Nodes_SubfileNodeDownCast).ToNodes_SubfileNode()
        {
            _usePath := subfileNode.FP.Get_usePath()
            if _usePath != nil {
                usePath := _usePath.(string)
                dependInfo.FP.AddSubMod(usePath)
            }
        }
    }
    dependInfo.FP.Output(self.stream)
}


// declaration Class -- Front
type OutputDepend_FrontMtd interface {
    Error(arg1 string)
    LoadFromLnsTxt(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string) LnsAny
    LoadMeta(arg1 *FrontInterface_ImportModuleInfo, arg2 string) LnsAny
    LoadModule(arg1 string)(LnsAny, LnsAny)
    SearchModule(arg1 string) LnsAny
}
type OutputDepend_Front struct {
    FP OutputDepend_FrontMtd
}
func OutputDepend_Front2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*OutputDepend_Front).FP
}
type OutputDepend_FrontDownCast interface {
    ToOutputDepend_Front() *OutputDepend_Front
}
func OutputDepend_FrontDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(OutputDepend_FrontDownCast)
    if ok { return work.ToOutputDepend_Front() }
    return nil
}
func (obj *OutputDepend_Front) ToOutputDepend_Front() *OutputDepend_Front {
    return obj
}
func NewOutputDepend_Front() *OutputDepend_Front {
    obj := &OutputDepend_Front{}
    obj.FP = obj
    obj.InitOutputDepend_Front()
    return obj
}
func (self *OutputDepend_Front) InitOutputDepend_Front() {
}
// 102: decl @lune.@base.@OutputDepend.Front.loadModule
func (self *OutputDepend_Front) LoadModule(mod string)(LnsAny, LnsAny) {
    Util_err("not implements")
// insert a dummy
    return nil,nil
}

// 105: decl @lune.@base.@OutputDepend.Front.loadFromLnsTxt
func (self *OutputDepend_Front) LoadFromLnsTxt(importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    Util_err("not implements")
// insert a dummy
    return nil
}

// 111: decl @lune.@base.@OutputDepend.Front.searchModule
func (self *OutputDepend_Front) SearchModule(mod string) LnsAny {
    Util_err("not implements")
// insert a dummy
    return nil
}

// 114: decl @lune.@base.@OutputDepend.Front.error
func (self *OutputDepend_Front) Error(message string) {
    Util_err(message)
}

// 118: decl @lune.@base.@OutputDepend.Front.loadMeta
func (self *OutputDepend_Front) LoadMeta(importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    var metaPath string
    metaPath = Lns_car(Lns_getVM().String_gsub(mod,"%.", "/")).(string) + ".meta"
    Lns_print([]LnsAny{metaPath})
    var loaded LnsAny
    var mess LnsAny
    loaded,mess = Lns_getVM().Loadfile(metaPath)
    if loaded != nil{
        loaded_5503 := loaded.(*Lns_luaValue)
        var meta LnsAny
        meta = OutputDepend_convExp487(Lns_2DDD(Lns_getVM().RunLoadedfunc(loaded_5503,Lns_2DDD([]LnsAny{}))[0]))
        return meta
    }
    Lns_print([]LnsAny{mess})
    return nil
}


func Lns_OutputDepend_init() {
    if init_OutputDepend { return }
    init_OutputDepend = true
    OutputDepend__mod__ = "@lune.@base.@OutputDepend"
    Lns_InitMod()
    Lns_Nodes_init()
    Lns_Util_init()
    Lns_Ast_init()
    Lns_TransUnit_init()
    Lns_convLua_init()
    Lns_frontInterface_init()
}
func init() {
    init_OutputDepend = false
}
