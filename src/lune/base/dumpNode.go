// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_dumpNode bool
var dumpNode__mod__ string
// for 59
func dumpNode_convExp137(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 103: decl @lune.@base.@dumpNode.createFilter
func DumpNode_createFilter(moduleTypeInfo *Ast_TypeInfo,processInfo *Ast_ProcessInfo,stream Lns_oStream) *Nodes_Filter {
    return &NewdumpNode_dumpFilter(true, moduleTypeInfo, moduleTypeInfo.FP.Get_scope(), stream, processInfo).Nodes_Filter
}

// 110: decl @lune.@base.@dumpNode.filter
func dumpNode_filter_1049_(node *Nodes_Node,filter *dumpNode_dumpFilter,opt *DumpNode_Opt) {
    node.FP.ProcessFilter(&filter.Nodes_Filter, DumpNode_Opt2Stem(opt))
}

// 555: decl @lune.@base.@dumpNode.getTypeListTxt
func dumpNode_getTypeListTxt_1199_(typeList *LnsList) string {
    var txt string
    txt = ""
    for _index, _typeInfo := range( typeList.Items ) {
        index := _index + 1
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            txt = txt + ", "
            
        }
        txt = txt + typeInfo.FP.GetTxt(nil, nil, nil)
        
    }
    return txt
}

// 880: decl @lune.@base.@dumpNode.Ast2Dump
func dumpNode_Ast2Dump_1316_(ast *TransUnit_ASTInfo) string {
    var stream *Util_memStream
    stream = NewUtil_memStream()
    var conv *Nodes_Filter
    conv = DumpNode_createFilter(ast.FP.Get_moduleTypeInfo(), ast.FP.Get_processInfo(), stream.FP)
    ast.FP.Get_node().FP.ProcessFilter(conv, DumpNode_Opt2Stem(NewDumpNode_Opt("", 0)))
    return stream.FP.Get_txt()
}

// declaration Class -- Opt
type DumpNode_OptMtd interface {
    Get()(string, LnsInt)
    NextOpt() *DumpNode_Opt
}
type DumpNode_Opt struct {
    prefix string
    depth LnsInt
    next LnsAny
    FP DumpNode_OptMtd
}
func DumpNode_Opt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*DumpNode_Opt).FP
}
type DumpNode_OptDownCast interface {
    ToDumpNode_Opt() *DumpNode_Opt
}
func DumpNode_OptDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(DumpNode_OptDownCast)
    if ok { return work.ToDumpNode_Opt() }
    return nil
}
func (obj *DumpNode_Opt) ToDumpNode_Opt() *DumpNode_Opt {
    return obj
}
func NewDumpNode_Opt(arg1 string, arg2 LnsInt) *DumpNode_Opt {
    obj := &DumpNode_Opt{}
    obj.FP = obj
    obj.InitDumpNode_Opt(arg1, arg2)
    return obj
}
// 36: DeclConstr
func (self *DumpNode_Opt) InitDumpNode_Opt(prefix string,depth LnsInt) {
    self.prefix = prefix
    
    self.depth = depth
    
    self.next = nil
    
}

// 41: decl @lune.@base.@dumpNode.Opt.get
func (self *DumpNode_Opt) Get()(string, LnsInt) {
    return self.prefix, self.depth
}

// 44: decl @lune.@base.@dumpNode.Opt.nextOpt
func (self *DumpNode_Opt) NextOpt() *DumpNode_Opt {
    {
        __exp := self.next
        if __exp != nil {
            _exp := __exp.(*DumpNode_Opt)
            return _exp
        }
    }
    var opt *DumpNode_Opt
    opt = NewDumpNode_Opt(self.prefix + "  ", self.depth + 1)
    self.next = opt
    
    return opt
}


// declaration Class -- dumpFilter
type dumpNode_dumpFilterMtd interface {
    DefaultProcess(arg1 *Nodes_Node, arg2 LnsAny)
    dump(arg1 *DumpNode_Opt, arg2 *Nodes_Node, arg3 string)
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
    processDeclFuncInfo(arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo, arg3 *DumpNode_Opt)
    ProcessDeclMacro(arg1 *Nodes_DeclMacroNode, arg2 LnsAny)
    ProcessDeclMember(arg1 *Nodes_DeclMemberNode, arg2 LnsAny)
    ProcessDeclMethod(arg1 *Nodes_DeclMethodNode, arg2 LnsAny)
    ProcessDeclVar(arg1 *Nodes_DeclVarNode, arg2 LnsAny)
    ProcessEnv(arg1 *Nodes_EnvNode, arg2 LnsAny)
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
    writeln(arg1 *DumpNode_Opt, arg2 string)
}
type dumpNode_dumpFilter struct {
    Nodes_Filter
    stream Lns_oStream
    processInfo *Ast_ProcessInfo
    FP dumpNode_dumpFilterMtd
}
func dumpNode_dumpFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*dumpNode_dumpFilter).FP
}
type dumpNode_dumpFilterDownCast interface {
    TodumpNode_dumpFilter() *dumpNode_dumpFilter
}
func dumpNode_dumpFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(dumpNode_dumpFilterDownCast)
    if ok { return work.TodumpNode_dumpFilter() }
    return nil
}
func (obj *dumpNode_dumpFilter) TodumpNode_dumpFilter() *dumpNode_dumpFilter {
    return obj
}
func NewdumpNode_dumpFilter(arg1 bool, arg2 LnsAny, arg3 LnsAny, arg4 Lns_oStream, arg5 *Ast_ProcessInfo) *dumpNode_dumpFilter {
    obj := &dumpNode_dumpFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitdumpNode_dumpFilter(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *dumpNode_dumpFilter) InitdumpNode_dumpFilter(arg1 bool, arg2 LnsAny, arg3 LnsAny, arg4 Lns_oStream, arg5 *Ast_ProcessInfo) {
    self.Nodes_Filter.InitNodes_Filter( arg1,arg2,arg3)
    self.stream = arg4
    self.processInfo = arg5
}
// 58: decl @lune.@base.@dumpNode.dumpFilter.writeln
func (self *dumpNode_dumpFilter) writeln(opt *DumpNode_Opt,txt string) {
    var prefix string
    prefix = dumpNode_convExp137(Lns_2DDD(opt.FP.Get()))
    self.stream.Write(Lns_getVM().String_format("%s%s\n", []LnsAny{prefix, Lns_car(Lns_getVM().String_gsub(txt," *$", "")).(string)}))
}

// 62: decl @lune.@base.@dumpNode.dumpFilter.dump
func (self *dumpNode_dumpFilter) dump(opt *DumpNode_Opt,node *Nodes_Node,txt string) {
    var typeStr string
    typeStr = ""
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType()
    if expType.FP.Equals(self.processInfo, Ast_builtinTypeNone, nil, nil){
        typeStr = Lns_getVM().String_format("(%d:%s:%s)", []LnsAny{expType.FP.Get_typeId(), expType.FP.GetTxt(nil, nil, nil), expType.FP.Get_kind()})
        
    }
    var comment string
    {
        _commentList := node.FP.Get_commentList()
        if _commentList != nil {
            commentList := _commentList.(*LnsList)
            comment = Lns_getVM().String_format("comment:%d,%d", []LnsAny{commentList.Len(), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( node.FP.Get_tailComment()) &&
                Lns_GetEnv().SetStackVal( 1) ||
                Lns_GetEnv().SetStackVal( 0) ).(LnsInt)})
            
        } else {
            if Lns_isCondTrue( node.FP.Get_tailComment()){
                comment = "comment:0,1"
                
            } else { 
                comment = ""
                
            }
        }
    }
    var attrib string
    attrib = ""
    if node.FP.HasNilAccess(){
        attrib = Lns_getVM().String_format("%s %s", []LnsAny{attrib, "nilacc"})
        
    }
    if node.FP.IsThreading(){
        attrib = Lns_getVM().String_format("%s %s", []LnsAny{attrib, "thread"})
        
    }
    if len(attrib) != 0{
        attrib = Lns_getVM().String_format("[%s]", []LnsAny{attrib})
        
    }
    self.FP.writeln(opt, Lns_getVM().String_format(": %s:%d(%d:%d) %s %s %s %s", []LnsAny{Nodes_getNodeKindName(node.FP.Get_kind()), node.FP.Get_id(), node.FP.Get_effectivePos().LineNo, node.FP.Get_effectivePos().Column, attrib, txt, typeStr, comment}))
}

// 114: decl @lune.@base.@dumpNode.dumpFilter.processNone
func (self *dumpNode_dumpFilter) ProcessNone(node *Nodes_NoneNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
}

// 119: decl @lune.@base.@dumpNode.dumpFilter.processLuneControl
func (self *dumpNode_dumpFilter) ProcessLuneControl(node *Nodes_LuneControlNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_pragma().(LnsAlgeVal).GetTxt())
}

// 124: decl @lune.@base.@dumpNode.dumpFilter.processBlankLine
func (self *dumpNode_dumpFilter) ProcessBlankLine(node *Nodes_BlankLineNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%d", []LnsAny{node.FP.Get_lineNum()}))
}

// 129: decl @lune.@base.@dumpNode.dumpFilter.processLuneKind
func (self *dumpNode_dumpFilter) ProcessLuneKind(node *Nodes_LuneKindNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
}

// 135: decl @lune.@base.@dumpNode.dumpFilter.processImport
func (self *dumpNode_dumpFilter) ProcessImport(node *Nodes_ImportNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_modulePath())
}

// 140: decl @lune.@base.@dumpNode.dumpFilter.processRoot
func (self *dumpNode_dumpFilter) ProcessRoot(node *Nodes_RootNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    for _, _child := range( node.FP.Get_children().Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(child, self, opt.FP.NextOpt())
    }
}

// 148: decl @lune.@base.@dumpNode.dumpFilter.processSubfile
func (self *dumpNode_dumpFilter) ProcessSubfile(node *Nodes_SubfileNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
}

// 154: decl @lune.@base.@dumpNode.dumpFilter.processEnv
func (self *dumpNode_dumpFilter) ProcessEnv(node *Nodes_EnvNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
}

// 161: decl @lune.@base.@dumpNode.dumpFilter.processBlockSub
func (self *dumpNode_dumpFilter) ProcessBlockSub(node *Nodes_BlockNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    for _, _statement := range( node.FP.Get_stmtList().Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(statement, self, opt.FP.NextOpt())
    }
}

// 170: decl @lune.@base.@dumpNode.dumpFilter.processScope
func (self *dumpNode_dumpFilter) ProcessScope(node *Nodes_ScopeNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Nodes_ScopeKind_getTxt( node.FP.Get_scopeKind()))
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
}

// 177: decl @lune.@base.@dumpNode.dumpFilter.processStmtExp
func (self *dumpNode_dumpFilter) ProcessStmtExp(node *Nodes_StmtExpNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
}

// 184: decl @lune.@base.@dumpNode.dumpFilter.processDeclEnum
func (self *dumpNode_dumpFilter) ProcessDeclEnum(node *Nodes_DeclEnumNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_name().Txt)
    var enumTypeInfo *Ast_EnumTypeInfo
    enumTypeInfo = Lns_unwrap( (Ast_EnumTypeInfoDownCastF(node.FP.Get_expType().FP))).(*Ast_EnumTypeInfo)
    for _, _name := range( node.FP.Get_valueNameList().Items ) {
        name := _name.(Types_TokenDownCast).ToTypes_Token()
        var valInfo *Ast_EnumValInfo
        valInfo = Lns_unwrap( enumTypeInfo.FP.GetEnumValInfo(name.Txt)).(*Ast_EnumValInfo)
        var val LnsAny
        switch _exp943 := valInfo.FP.Get_val().(type) {
        case *Ast_EnumLiteral__Int:
        x := _exp943.Val1
            val = x
            
        case *Ast_EnumLiteral__Real:
        x := _exp943.Val1
            val = x
            
        case *Ast_EnumLiteral__Str:
        x := _exp943.Val1
            val = x
            
        }
        self.FP.writeln(opt, Lns_getVM().String_format("  %s: %s, %s", []LnsAny{name.Txt, valInfo.FP.Get_val().(LnsAlgeVal).GetTxt(), val}))
    }
}

// 207: decl @lune.@base.@dumpNode.dumpFilter.processDeclAlge
func (self *dumpNode_dumpFilter) ProcessDeclAlge(node *Nodes_DeclAlgeNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var algeTypeInfo *Ast_AlgeTypeInfo
    algeTypeInfo = node.FP.Get_algeType()
    self.FP.dump(opt, &node.Nodes_Node, algeTypeInfo.FP.Get_rawTxt())
    {
        __collection1024 := algeTypeInfo.FP.Get_valInfoMap()
        __sorted1024 := __collection1024.CreateKeyListStr()
        __sorted1024.Sort( LnsItemKindStr, nil )
        for _, ___key1024 := range( __sorted1024.Items ) {
            valInfo := __collection1024.Items[ ___key1024 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            self.FP.writeln(opt, Lns_getVM().String_format("  %s: %s", []LnsAny{algeTypeInfo.FP.Get_rawTxt(), valInfo.FP.Get_name()}))
        }
    }
}

// 217: decl @lune.@base.@dumpNode.dumpFilter.processNewAlgeVal
func (self *dumpNode_dumpFilter) ProcessNewAlgeVal(node *Nodes_NewAlgeValNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_name().Txt)
    for _, _exp := range( node.FP.Get_paramList().Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(exp, self, opt.FP.NextOpt())
    }
}

// 226: decl @lune.@base.@dumpNode.dumpFilter.processProtoClass
func (self *dumpNode_dumpFilter) ProcessProtoClass(node *Nodes_ProtoClassNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_name().Txt)
}

// 232: decl @lune.@base.@dumpNode.dumpFilter.processDeclClass
func (self *dumpNode_dumpFilter) ProcessDeclClass(node *Nodes_DeclClassNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_name().Txt)
    for _, _field := range( node.FP.Get_fieldList().Items ) {
        field := _field.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(field, self, opt.FP.NextOpt())
    }
}

// 240: decl @lune.@base.@dumpNode.dumpFilter.processDeclMember
func (self *dumpNode_dumpFilter) ProcessDeclMember(node *Nodes_DeclMemberNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_name().Txt)
    dumpNode_filter_1049_(&node.FP.Get_refType().Nodes_Node, self, opt.FP.NextOpt())
}

// 246: decl @lune.@base.@dumpNode.dumpFilter.processExpMacroExp
func (self *dumpNode_dumpFilter) ProcessExpMacroExp(node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList()
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(stmt, self, opt.FP.NextOpt())
    }
}

// 255: decl @lune.@base.@dumpNode.dumpFilter.processDeclMacro
func (self *dumpNode_dumpFilter) ProcessDeclMacro(node *Nodes_DeclMacroNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_expType().FP.GetTxt(nil, nil, nil))
}

// 260: decl @lune.@base.@dumpNode.dumpFilter.processExpMacroStat
func (self *dumpNode_dumpFilter) ProcessExpMacroStat(node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_expType().FP.GetTxt(nil, nil, nil))
    for _, _expStr := range( node.FP.Get_expStrList().Items ) {
        expStr := _expStr.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(expStr, self, opt.FP.NextOpt())
    }
}

// 270: decl @lune.@base.@dumpNode.dumpFilter.processUnwrapSet
func (self *dumpNode_dumpFilter) ProcessUnwrapSet(node *Nodes_UnwrapSetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(&node.FP.Get_dstExpList().Nodes_Node, self, opt.FP.NextOpt())
    dumpNode_filter_1049_(&node.FP.Get_srcExpList().Nodes_Node, self, opt.FP.NextOpt())
    if Lns_isCondTrue( node.FP.Get_unwrapBlock()){
        dumpNode_filter_1049_(&Lns_unwrap( node.FP.Get_unwrapBlock()).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt())
    }
}

// 282: decl @lune.@base.@dumpNode.dumpFilter.processIfUnwrap
func (self *dumpNode_dumpFilter) ProcessIfUnwrap(node *Nodes_IfUnwrapNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    for _, _expNode := range( node.FP.Get_expList().FP.Get_expList().Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(expNode, self, opt.FP.NextOpt())
    }
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
    if Lns_isCondTrue( node.FP.Get_nilBlock()){
        dumpNode_filter_1049_(&Lns_unwrap( node.FP.Get_nilBlock()).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt())
    }
}

// 295: decl @lune.@base.@dumpNode.dumpFilter.processWhen
func (self *dumpNode_dumpFilter) ProcessWhen(node *Nodes_WhenNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var symTxt string
    symTxt = ""
    for _, _symPair := range( node.FP.Get_symPairList().Items ) {
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        symTxt = Lns_getVM().String_format("%s %s", []LnsAny{symTxt, symPair.FP.Get_src().FP.Get_name()})
        
    }
    self.FP.dump(opt, &node.Nodes_Node, symTxt)
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
    {
        __exp := node.FP.Get_elseBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 311: decl @lune.@base.@dumpNode.dumpFilter.processDeclVar
func (self *dumpNode_dumpFilter) ProcessDeclVar(node *Nodes_DeclVarNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var varName string
    varName = ""
    for _index, __var := range( node.FP.Get_varList().Items ) {
        index := _index + 1
        _var := __var.(Nodes_VarInfoDownCast).ToNodes_VarInfo()
        if index > 1{
            varName = varName + ","
            
        }
        varName = Lns_getVM().String_format("%s %s", []LnsAny{varName, _var.FP.Get_name().Txt})
        
    }
    if Lns_isCondTrue( node.FP.Get_unwrapBlock()){
        varName = "!" + varName
        
    }
    varName = Lns_getVM().String_format("%s %s", []LnsAny{node.FP.Get_mode(), varName})
    
    self.FP.dump(opt, &node.Nodes_Node, varName)
    for _, __var := range( node.FP.Get_varList().Items ) {
        _var := __var.(Nodes_VarInfoDownCast).ToNodes_VarInfo()
        {
            __exp := _var.FP.Get_refType()
            if __exp != nil {
                _exp := __exp.(*Nodes_RefTypeNode)
                dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
            }
        }
    }
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
    {
        __exp := node.FP.Get_unwrapBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
    {
        __exp := node.FP.Get_thenBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
    {
        __exp := node.FP.Get_syncBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 350: decl @lune.@base.@dumpNode.dumpFilter.processDeclArg
func (self *dumpNode_dumpFilter) ProcessDeclArg(node *Nodes_DeclArgNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s:%s", []LnsAny{node.FP.Get_name().Txt, node.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
}

// 356: decl @lune.@base.@dumpNode.dumpFilter.processDeclArgDDD
func (self *dumpNode_dumpFilter) ProcessDeclArgDDD(node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "...")
}

// 361: decl @lune.@base.@dumpNode.dumpFilter.processExpSubDDD
func (self *dumpNode_dumpFilter) ProcessExpSubDDD(node *Nodes_ExpSubDDDNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("... (%d)", []LnsAny{node.FP.Get_remainIndex()}))
}

// 368: decl @lune.@base.@dumpNode.dumpFilter.processDeclForm
func (self *dumpNode_dumpFilter) ProcessDeclForm(node *Nodes_DeclFormNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_expType().FP.GetTxt(self.FP.Get_typeNameCtrl(), nil, nil))
}

// 373: decl @lune.@base.@dumpNode.dumpFilter.processDeclFuncInfo
func (self *dumpNode_dumpFilter) processDeclFuncInfo(node *Nodes_Node,declInfo *Nodes_DeclFuncInfo,opt *DumpNode_Opt) {
    var name string
    
    {
        _name := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(declInfo.FP.Get_name()) && 
        Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Types_Token).Txt))
        if _name == nil{
            name = "<anonymous>"
            
        } else {
            name = _name.(string)
        }
    }
    name = node.FP.Get_expType().FP.Get_display_stirng_with(name, nil)
    
    if Ast_TypeInfo_isMut(node.FP.Get_expType()){
        name = name + " mut"
        
    }
    {
        __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(node.FP.Get_expType().FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_closureSymList()}))
        if __exp != nil {
            _exp := __exp.(*LnsList)
            if _exp.Len() > 0{
                name = name + " closure"
                
            }
        }
    }
    self.FP.dump(opt, node, name)
    var argList *LnsList
    argList = declInfo.FP.Get_argList()
    for _, _arg := range( argList.Items ) {
        arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(arg, self, opt.FP.NextOpt())
    }
    {
        __exp := declInfo.FP.Get_body()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 397: decl @lune.@base.@dumpNode.dumpFilter.processDeclFunc
func (self *dumpNode_dumpFilter) ProcessDeclFunc(node *Nodes_DeclFuncNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo(), opt)
}

// 402: decl @lune.@base.@dumpNode.dumpFilter.processDeclMethod
func (self *dumpNode_dumpFilter) ProcessDeclMethod(node *Nodes_DeclMethodNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo(), opt)
}

// 407: decl @lune.@base.@dumpNode.dumpFilter.processProtoMethod
func (self *dumpNode_dumpFilter) ProcessProtoMethod(node *Nodes_ProtoMethodNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo(), opt)
}

// 412: decl @lune.@base.@dumpNode.dumpFilter.processDeclConstr
func (self *dumpNode_dumpFilter) ProcessDeclConstr(node *Nodes_DeclConstrNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo(), opt)
}

// 418: decl @lune.@base.@dumpNode.dumpFilter.processDeclDestr
func (self *dumpNode_dumpFilter) ProcessDeclDestr(node *Nodes_DeclDestrNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
}

// 424: decl @lune.@base.@dumpNode.dumpFilter.processExpCallSuperCtor
func (self *dumpNode_dumpFilter) ProcessExpCallSuperCtor(node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType()
    self.FP.dump(opt, &node.Nodes_Node, typeInfo.FP.GetTxt(nil, nil, nil))
}

// 430: decl @lune.@base.@dumpNode.dumpFilter.processExpCallSuper
func (self *dumpNode_dumpFilter) ProcessExpCallSuper(node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType()
    self.FP.dump(opt, &node.Nodes_Node, typeInfo.FP.GetTxt(nil, nil, nil))
}

// 436: decl @lune.@base.@dumpNode.dumpFilter.processRefType
func (self *dumpNode_dumpFilter) ProcessRefType(node *Nodes_RefTypeNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( node.FP.Get_refFlag()) &&
        Lns_GetEnv().SetStackVal( "&") ||
        Lns_GetEnv().SetStackVal( "") ).(string)) + (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( node.FP.Get_mutFlag()) &&
        Lns_GetEnv().SetStackVal( "mut ") ||
        Lns_GetEnv().SetStackVal( "") ).(string)))
    dumpNode_filter_1049_(node.FP.Get_name(), self, opt.FP.NextOpt())
}

// 444: decl @lune.@base.@dumpNode.dumpFilter.processIf
func (self *dumpNode_dumpFilter) ProcessIf(node *Nodes_IfNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList()
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if stmt.FP.Get_exp().FP.Get_kind() != Nodes_nodeKindEnum__None{
            dumpNode_filter_1049_(stmt.FP.Get_exp(), self, opt.FP.NextOpt())
        }
        dumpNode_filter_1049_(&stmt.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
    }
}

// 456: decl @lune.@base.@dumpNode.dumpFilter.processSwitch
func (self *dumpNode_dumpFilter) ProcessSwitch(node *Nodes_SwitchNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
    var caseList *LnsList
    caseList = node.FP.Get_caseList()
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        dumpNode_filter_1049_(&caseInfo.FP.Get_expList().Nodes_Node, self, opt.FP.NextOpt())
        dumpNode_filter_1049_(&caseInfo.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
    }
    {
        __exp := node.FP.Get_default()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 470: decl @lune.@base.@dumpNode.dumpFilter.processMatch
func (self *dumpNode_dumpFilter) ProcessMatch(node *Nodes_MatchNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var prefix string
    var depth LnsInt
    prefix,depth = opt.FP.Get()
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_val(), self, opt.FP.NextOpt())
    var caseList *LnsList
    caseList = node.FP.Get_caseList()
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        dumpNode_filter_1049_(&caseInfo.FP.Get_block().Nodes_Node, self, NewDumpNode_Opt(prefix + "  " + caseInfo.FP.Get_valInfo().FP.Get_name(), depth + 1))
    }
    {
        __exp := node.FP.Get_defaultBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            dumpNode_filter_1049_(_exp, self, opt.FP.NextOpt())
        }
    }
}

// 485: decl @lune.@base.@dumpNode.dumpFilter.processWhile
func (self *dumpNode_dumpFilter) ProcessWhile(node *Nodes_WhileNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
}

// 492: decl @lune.@base.@dumpNode.dumpFilter.processRepeat
func (self *dumpNode_dumpFilter) ProcessRepeat(node *Nodes_RepeatNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
}

// 499: decl @lune.@base.@dumpNode.dumpFilter.processFor
func (self *dumpNode_dumpFilter) ProcessFor(node *Nodes_ForNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_val().FP.Get_name())
    dumpNode_filter_1049_(node.FP.Get_init(), self, opt.FP.NextOpt())
    dumpNode_filter_1049_(node.FP.Get_to(), self, opt.FP.NextOpt())
    {
        __exp := node.FP.Get_delta()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            dumpNode_filter_1049_(_exp, self, opt.FP.NextOpt())
        }
    }
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
}

// 510: decl @lune.@base.@dumpNode.dumpFilter.processApply
func (self *dumpNode_dumpFilter) ProcessApply(node *Nodes_ApplyNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var varNames string
    varNames = ""
    var varList *LnsList
    varList = node.FP.Get_varList()
    for _, __var := range( varList.Items ) {
        _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        varNames = varNames + _var.FP.Get_name() + " "
        
    }
    self.FP.dump(opt, &node.Nodes_Node, varNames)
    dumpNode_filter_1049_(&node.FP.Get_expList().Nodes_Node, self, opt.FP.NextOpt())
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
}

// 522: decl @lune.@base.@dumpNode.dumpFilter.processForeach
func (self *dumpNode_dumpFilter) ProcessForeach(node *Nodes_ForeachNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var index string
    index = ""
    {
        __exp := node.FP.Get_key()
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            index = _exp.FP.Get_name()
            
        }
    }
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_val().FP.Get_name() + " " + index)
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
}

// 533: decl @lune.@base.@dumpNode.dumpFilter.processForsort
func (self *dumpNode_dumpFilter) ProcessForsort(node *Nodes_ForsortNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var index string
    index = ""
    {
        __exp := node.FP.Get_key()
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            index = _exp.FP.Get_name()
            
        }
    }
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_val().FP.Get_name() + " " + index)
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
}

// 545: decl @lune.@base.@dumpNode.dumpFilter.processExpUnwrap
func (self *dumpNode_dumpFilter) ProcessExpUnwrap(node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
    {
        __exp := node.FP.Get_default()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            dumpNode_filter_1049_(_exp, self, opt.FP.NextOpt())
        }
    }
}

// 566: decl @lune.@base.@dumpNode.dumpFilter.processExpCall
func (self *dumpNode_dumpFilter) ProcessExpCall(node *Nodes_ExpCallNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var mess string
    mess = dumpNode_getTypeListTxt_1199_(node.FP.Get_expTypeList())
    self.FP.dump(opt, &node.Nodes_Node, mess)
    dumpNode_filter_1049_(node.FP.Get_func(), self, opt.FP.NextOpt())
    {
        __exp := node.FP.Get_argList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 577: decl @lune.@base.@dumpNode.dumpFilter.processExpList
func (self *dumpNode_dumpFilter) ProcessExpList(node *Nodes_ExpListNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var mess string
    {
        _mRetExp := node.FP.Get_mRetExp()
        if _mRetExp != nil {
            mRetExp := _mRetExp.(*Nodes_MRetExp)
            mess = Lns_getVM().String_format("hasMRetExp (%d): ", []LnsAny{mRetExp.FP.Get_index()})
            
        } else {
            mess = "noMRetExp: "
            
        }
    }
    for _, _expType := range( node.FP.Get_expTypeList().Items ) {
        expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        mess = Lns_getVM().String_format("%s %s", []LnsAny{mess, expType.FP.GetTxt(nil, nil, nil)})
        
    }
    self.FP.dump(opt, &node.Nodes_Node, mess)
    var expList *LnsList
    expList = node.FP.Get_expList()
    for _, _exp := range( expList.Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(exp, self, opt.FP.NextOpt())
    }
}

// 597: decl @lune.@base.@dumpNode.dumpFilter.processExpMRet
func (self *dumpNode_dumpFilter) ProcessExpMRet(node *Nodes_ExpMRetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_mRet(), self, opt.FP.NextOpt())
}

// 604: decl @lune.@base.@dumpNode.dumpFilter.processExpAccessMRet
func (self *dumpNode_dumpFilter) ProcessExpAccessMRet(node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%d", []LnsAny{node.FP.Get_index()}))
}

// 610: decl @lune.@base.@dumpNode.dumpFilter.processExpOp1
func (self *dumpNode_dumpFilter) ProcessExpOp1(node *Nodes_ExpOp1Node,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_op().Txt)
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
}

// 617: decl @lune.@base.@dumpNode.dumpFilter.processExpToDDD
func (self *dumpNode_dumpFilter) ProcessExpToDDD(node *Nodes_ExpToDDDNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(&node.FP.Get_expList().Nodes_Node, self, opt.FP.NextOpt())
}

// 624: decl @lune.@base.@dumpNode.dumpFilter.processExpMultiTo1
func (self *dumpNode_dumpFilter) ProcessExpMultiTo1(node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
}

// 631: decl @lune.@base.@dumpNode.dumpFilter.processExpCast
func (self *dumpNode_dumpFilter) ProcessExpCast(node *Nodes_ExpCastNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s(%d) -> %s(%d)", []LnsAny{node.FP.Get_exp().FP.Get_expType().FP.GetTxt(self.FP.Get_typeNameCtrl(), nil, nil), node.FP.Get_exp().FP.Get_expType().FP.Get_typeId(), node.FP.Get_castType().FP.GetTxt(self.FP.Get_typeNameCtrl(), nil, nil), node.FP.Get_castType().FP.Get_typeId()}))
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
}

// 642: decl @lune.@base.@dumpNode.dumpFilter.processExpParen
func (self *dumpNode_dumpFilter) ProcessExpParen(node *Nodes_ExpParenNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "()")
    dumpNode_filter_1049_(node.FP.Get_exp(), self, opt.FP.NextOpt())
}

// 648: decl @lune.@base.@dumpNode.dumpFilter.processExpSetVal
func (self *dumpNode_dumpFilter) ProcessExpSetVal(node *Nodes_ExpSetValNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("= %s", []LnsAny{node.FP.Get_expType().FP.GetTxt(self.FP.Get_typeNameCtrl(), nil, nil)}))
    dumpNode_filter_1049_(node.FP.Get_exp1(), self, opt.FP.NextOpt())
    dumpNode_filter_1049_(&node.FP.Get_exp2().Nodes_Node, self, opt.FP.NextOpt())
}

// 655: decl @lune.@base.@dumpNode.dumpFilter.processExpSetItem
func (self *dumpNode_dumpFilter) ProcessExpSetItem(node *Nodes_ExpSetItemNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var indexSym string
    indexSym = ""
    var indexNode LnsAny
    indexNode = nil
    switch _exp3807 := node.FP.Get_index().(type) {
    case *Nodes_IndexVal__NodeIdx:
    index := _exp3807.Val1
        indexNode = index
        
    case *Nodes_IndexVal__SymIdx:
    index := _exp3807.Val1
        indexSym = index
        
    }
    self.FP.dump(opt, &node.Nodes_Node, indexSym)
    dumpNode_filter_1049_(node.FP.Get_val(), self, opt.FP.NextOpt())
    if indexNode != nil{
        indexNode_5707 := indexNode.(*Nodes_Node)
        dumpNode_filter_1049_(indexNode_5707, self, opt.FP.NextOpt())
    }
    dumpNode_filter_1049_(node.FP.Get_exp2(), self, opt.FP.NextOpt())
}

// 675: decl @lune.@base.@dumpNode.dumpFilter.processExpOp2
func (self *dumpNode_dumpFilter) ProcessExpOp2(node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s -> %s", []LnsAny{node.FP.Get_op().Txt, node.FP.Get_expType().FP.GetTxt(self.FP.Get_typeNameCtrl(), nil, nil)}))
    dumpNode_filter_1049_(node.FP.Get_exp1(), self, opt.FP.NextOpt())
    dumpNode_filter_1049_(node.FP.Get_exp2(), self, opt.FP.NextOpt())
}

// 683: decl @lune.@base.@dumpNode.dumpFilter.processExpNew
func (self *dumpNode_dumpFilter) ProcessExpNew(node *Nodes_ExpNewNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_symbol(), self, opt.FP.NextOpt())
    {
        __exp := node.FP.Get_argList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 693: decl @lune.@base.@dumpNode.dumpFilter.processExpRef
func (self *dumpNode_dumpFilter) ProcessExpRef(node *Nodes_ExpRefNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s:%s", []LnsAny{node.FP.Get_symbolInfo().FP.Get_name(), node.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
}

// 699: decl @lune.@base.@dumpNode.dumpFilter.processExpRefItem
func (self *dumpNode_dumpFilter) ProcessExpRefItem(node *Nodes_ExpRefItemNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "seq[exp] " + node.FP.Get_expType().FP.GetTxt(nil, nil, nil))
    dumpNode_filter_1049_(node.FP.Get_val(), self, opt.FP.NextOpt())
    {
        __exp := node.FP.Get_index()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            dumpNode_filter_1049_(_exp, self, opt.FP.NextOpt())
        }
    }
}

// 708: decl @lune.@base.@dumpNode.dumpFilter.processRefField
func (self *dumpNode_dumpFilter) ProcessRefField(node *Nodes_RefFieldNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s:%s:%s", []LnsAny{node.FP.Get_field().Txt, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(node.FP.Get_symbolInfo()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_SymbolInfo).FP.Get_mutable()}))) &&
        Lns_GetEnv().SetStackVal( "mut") ||
        Lns_GetEnv().SetStackVal( "imut") ).(string), node.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
    dumpNode_filter_1049_(node.FP.Get_prefix(), self, opt.FP.NextOpt())
}

// 717: decl @lune.@base.@dumpNode.dumpFilter.processExpOmitEnum
func (self *dumpNode_dumpFilter) ProcessExpOmitEnum(node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s.%s", []LnsAny{node.FP.Get_expType().FP.GetTxt(nil, nil, nil), node.FP.Get_valToken().Txt}))
}

// 724: decl @lune.@base.@dumpNode.dumpFilter.processGetField
func (self *dumpNode_dumpFilter) ProcessGetField(node *Nodes_GetFieldNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("get_%s:%s", []LnsAny{node.FP.Get_field().Txt, node.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
    dumpNode_filter_1049_(node.FP.Get_prefix(), self, opt.FP.NextOpt())
}

// 732: decl @lune.@base.@dumpNode.dumpFilter.processReturn
func (self *dumpNode_dumpFilter) ProcessReturn(node *Nodes_ReturnNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 742: decl @lune.@base.@dumpNode.dumpFilter.processProvide
func (self *dumpNode_dumpFilter) ProcessProvide(node *Nodes_ProvideNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_symbol().FP.Get_name())
}

// 748: decl @lune.@base.@dumpNode.dumpFilter.processAlias
func (self *dumpNode_dumpFilter) ProcessAlias(node *Nodes_AliasNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s = %s", []LnsAny{node.FP.Get_newSymbol().FP.Get_name(), node.FP.Get_typeInfo().FP.GetTxt(nil, nil, nil)}))
}

// 754: decl @lune.@base.@dumpNode.dumpFilter.processTestCase
func (self *dumpNode_dumpFilter) ProcessTestCase(node *Nodes_TestCaseNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_name().Txt)
    dumpNode_filter_1049_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt())
}

// 761: decl @lune.@base.@dumpNode.dumpFilter.processTestBlock
func (self *dumpNode_dumpFilter) ProcessTestBlock(node *Nodes_TestBlockNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    for _, _statement := range( node.FP.Get_stmtList().Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_1049_(statement, self, opt.FP.NextOpt())
    }
}

// 770: decl @lune.@base.@dumpNode.dumpFilter.processBoxing
func (self *dumpNode_dumpFilter) ProcessBoxing(node *Nodes_BoxingNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_src(), self, opt.FP.NextOpt())
}

// 777: decl @lune.@base.@dumpNode.dumpFilter.processUnboxing
func (self *dumpNode_dumpFilter) ProcessUnboxing(node *Nodes_UnboxingNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
    dumpNode_filter_1049_(node.FP.Get_src(), self, opt.FP.NextOpt())
}

// 784: decl @lune.@base.@dumpNode.dumpFilter.processLiteralList
func (self *dumpNode_dumpFilter) ProcessLiteralList(node *Nodes_LiteralListNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "[]")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 792: decl @lune.@base.@dumpNode.dumpFilter.processLiteralSet
func (self *dumpNode_dumpFilter) ProcessLiteralSet(node *Nodes_LiteralSetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "(@)")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 800: decl @lune.@base.@dumpNode.dumpFilter.processLiteralMap
func (self *dumpNode_dumpFilter) ProcessLiteralMap(node *Nodes_LiteralMapNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "{}")
    var pairList *LnsList
    pairList = node.FP.Get_pairList()
    for _, _pair := range( pairList.Items ) {
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        dumpNode_filter_1049_(pair.FP.Get_key(), self, opt.FP.NextOpt())
        dumpNode_filter_1049_(pair.FP.Get_val(), self, opt.FP.NextOpt())
    }
}

// 811: decl @lune.@base.@dumpNode.dumpFilter.processLiteralArray
func (self *dumpNode_dumpFilter) ProcessLiteralArray(node *Nodes_LiteralArrayNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "[@]")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_1049_(&_exp.Nodes_Node, self, opt.FP.NextOpt())
        }
    }
}

// 819: decl @lune.@base.@dumpNode.dumpFilter.processLiteralChar
func (self *dumpNode_dumpFilter) ProcessLiteralChar(node *Nodes_LiteralCharNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s(%s)", []LnsAny{node.FP.Get_num(), node.FP.Get_token().Txt}))
}

// 825: decl @lune.@base.@dumpNode.dumpFilter.processLiteralInt
func (self *dumpNode_dumpFilter) ProcessLiteralInt(node *Nodes_LiteralIntNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s(%s)", []LnsAny{node.FP.Get_num(), node.FP.Get_token().Txt}))
}

// 831: decl @lune.@base.@dumpNode.dumpFilter.processLiteralReal
func (self *dumpNode_dumpFilter) ProcessLiteralReal(node *Nodes_LiteralRealNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_getVM().String_format("%s(%s)", []LnsAny{node.FP.Get_num(), node.FP.Get_token().Txt}))
}

// 837: decl @lune.@base.@dumpNode.dumpFilter.processLiteralString
func (self *dumpNode_dumpFilter) ProcessLiteralString(node *Nodes_LiteralStringNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_token().Txt)
    {
        _expList := node.FP.Get_orgParam()
        if _expList != nil {
            expList := _expList.(*Nodes_ExpListNode)
            for _, _param := range( expList.FP.Get_expList().Items ) {
                param := _param.(Nodes_NodeDownCast).ToNodes_Node()
                dumpNode_filter_1049_(param, self, opt.FP.NextOpt())
            }
        }
    }
}

// 847: decl @lune.@base.@dumpNode.dumpFilter.processLiteralBool
func (self *dumpNode_dumpFilter) ProcessLiteralBool(node *Nodes_LiteralBoolNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( node.FP.Get_token().Txt == "true") &&
        Lns_GetEnv().SetStackVal( "true") ||
        Lns_GetEnv().SetStackVal( "false") ).(string))
}

// 853: decl @lune.@base.@dumpNode.dumpFilter.processLiteralNil
func (self *dumpNode_dumpFilter) ProcessLiteralNil(node *Nodes_LiteralNilNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
}

// 858: decl @lune.@base.@dumpNode.dumpFilter.processBreak
func (self *dumpNode_dumpFilter) ProcessBreak(node *Nodes_BreakNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "")
}

// 863: decl @lune.@base.@dumpNode.dumpFilter.processLiteralSymbol
func (self *dumpNode_dumpFilter) ProcessLiteralSymbol(node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, node.FP.Get_token().Txt)
}

// 869: decl @lune.@base.@dumpNode.dumpFilter.processAbbr
func (self *dumpNode_dumpFilter) ProcessAbbr(node *Nodes_AbbrNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(opt, &node.Nodes_Node, "##")
}


func Lns_dumpNode_init() {
    if init_dumpNode { return }
    init_dumpNode = true
    dumpNode__mod__ = "@lune.@base.@dumpNode"
    Lns_InitMod()
    Lns_Ast_init()
    Lns_Nodes_init()
    Lns_LuneControl_init()
    Lns_Util_init()
    Lns_convLua_init()
    Lns_TransUnit_init()
}
func init() {
    init_dumpNode = false
}
