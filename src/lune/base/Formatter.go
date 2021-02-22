// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Formatter bool
var Formatter__mod__ string
// 54: decl @lune.@base.@Formatter.createFilter
func Formatter_createFilter(moduleTypeInfo *Ast_TypeInfo,stream Lns_oStream) *Nodes_Filter {
    return &NewFormatter_FormatterFilter(moduleTypeInfo, moduleTypeInfo.FP.Get_scope(), stream).Nodes_Filter
}

// 64: decl @lune.@base.@Formatter.filter
func Formatter_filter_1066_(node *Nodes_Node,filter *Formatter_FormatterFilter,opt *Formatter_Opt) {
    filter.FP.OutputHeadComment(node)
    node.FP.ProcessFilter(&filter.Nodes_Filter, Formatter_Opt2Stem(opt))
}



// declaration Class -- Opt
type Formatter_OptMtd interface {
    Get_parent() *Nodes_Node
    NextOpt(arg1 *Nodes_Node) *Formatter_Opt
}
type Formatter_Opt struct {
    parent *Nodes_Node
    FP Formatter_OptMtd
}
func Formatter_Opt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_Opt).FP
}
type Formatter_OptDownCast interface {
    ToFormatter_Opt() *Formatter_Opt
}
func Formatter_OptDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_OptDownCast)
    if ok { return work.ToFormatter_Opt() }
    return nil
}
func (obj *Formatter_Opt) ToFormatter_Opt() *Formatter_Opt {
    return obj
}
func NewFormatter_Opt(arg1 *Nodes_Node) *Formatter_Opt {
    obj := &Formatter_Opt{}
    obj.FP = obj
    obj.InitFormatter_Opt(arg1)
    return obj
}
func (self *Formatter_Opt) Get_parent() *Nodes_Node{ return self.parent }
// 33: DeclConstr
func (self *Formatter_Opt) InitFormatter_Opt(parent *Nodes_Node) {
    self.parent = parent
    
}

// 36: decl @lune.@base.@Formatter.Opt.nextOpt
func (self *Formatter_Opt) NextOpt(parent *Nodes_Node) *Formatter_Opt {
    return NewFormatter_Opt(parent)
}


// declaration Class -- FormatterFilter
type Formatter_FormatterFilterMtd interface {
    DefaultProcess(arg1 *Nodes_Node, arg2 LnsAny)
    Get_moduleInfoManager() Ast_ModuleInfoManager
    Get_optStack() *LnsList
    Get_typeNameCtrl() *Ast_TypeNameCtrl
    outputDeclClass(arg1 bool, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny)
    OutputHeadComment(arg1 *Nodes_Node)
    PopIndent()
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
    processDeclFuncInfo(arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo, arg3 *Formatter_Opt)
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
    PushIndent(arg1 LnsAny)
    ReturnToSource()
    SwitchToHeader()
    Write(arg1 string)
    Writeln(arg1 string)
}
type Formatter_FormatterFilter struct {
    Nodes_Filter
    stream *Util_SimpleSourceOStream
    FP Formatter_FormatterFilterMtd
}
func Formatter_FormatterFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Formatter_FormatterFilter).FP
}
type Formatter_FormatterFilterDownCast interface {
    ToFormatter_FormatterFilter() *Formatter_FormatterFilter
}
func Formatter_FormatterFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Formatter_FormatterFilterDownCast)
    if ok { return work.ToFormatter_FormatterFilter() }
    return nil
}
func (obj *Formatter_FormatterFilter) ToFormatter_FormatterFilter() *Formatter_FormatterFilter {
    return obj
}
func NewFormatter_FormatterFilter(arg1 *Ast_TypeInfo, arg2 LnsAny, arg3 Lns_oStream) *Formatter_FormatterFilter {
    obj := &Formatter_FormatterFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitFormatter_FormatterFilter(arg1, arg2, arg3)
    return obj
}
func (self *Formatter_FormatterFilter) PopIndent() {
self.stream. FP.PopIndent( )
}
func (self *Formatter_FormatterFilter) PushIndent(arg1 LnsAny) {
self.stream. FP.PushIndent( arg1)
}
func (self *Formatter_FormatterFilter) ReturnToSource() {
self.stream. FP.ReturnToSource( )
}
func (self *Formatter_FormatterFilter) SwitchToHeader() {
self.stream. FP.SwitchToHeader( )
}
func (self *Formatter_FormatterFilter) Write(arg1 string) {
self.stream. FP.Write( arg1)
}
func (self *Formatter_FormatterFilter) Writeln(arg1 string) {
self.stream. FP.Writeln( arg1)
}
// 44: DeclConstr
func (self *Formatter_FormatterFilter) InitFormatter_FormatterFilter(moduleTypeInfo *Ast_TypeInfo,moduleInfoManager LnsAny,stream Lns_oStream) {
    self.InitNodes_Filter(true, moduleTypeInfo, moduleInfoManager)
    self.stream = NewUtil_SimpleSourceOStream(stream, nil, 3)
    
}

// 58: decl @lune.@base.@Formatter.FormatterFilter.outputHeadComment
func (self *Formatter_FormatterFilter) OutputHeadComment(node *Nodes_Node) {
    for _, _commentNode := range( Lns_unwrapDefault( node.FP.Get_commentList(), NewLnsList([]LnsAny{})).(*LnsList).Items ) {
        commentNode := _commentNode.(Types_TokenDownCast).ToTypes_Token()
        self.FP.Writeln(commentNode.Txt)
    }
}

// 73: decl @lune.@base.@Formatter.FormatterFilter.processNone
func (self *Formatter_FormatterFilter) ProcessNone(node *Nodes_NoneNode,_opt LnsAny) {
}

// 79: decl @lune.@base.@Formatter.FormatterFilter.processBlankLine
func (self *Formatter_FormatterFilter) ProcessBlankLine(node *Nodes_BlankLineNode,_opt LnsAny) {
    {
        var _from240 LnsInt = 1
        var _to240 LnsInt = node.FP.Get_lineNum()
        for _work240 := _from240; _work240 <= _to240; _work240++ {
            self.FP.Writeln("")
        }
    }
}

// 87: decl @lune.@base.@Formatter.FormatterFilter.processImport
func (self *Formatter_FormatterFilter) ProcessImport(node *Nodes_ImportNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("import %s", []LnsAny{node.FP.Get_modulePath()}))
    if Lns_op_not(Lns_car(Lns_getVM().String_find(node.FP.Get_modulePath(),"%." + node.FP.Get_assignName() + "$", nil, nil))){
        self.FP.Write(Lns_getVM().String_format(" as %s", []LnsAny{node.FP.Get_assignName()}))
    }
    self.FP.Writeln(";")
}

// 96: decl @lune.@base.@Formatter.FormatterFilter.processRoot
func (self *Formatter_FormatterFilter) ProcessRoot(node *Nodes_RootNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    for _, _child := range( node.FP.Get_children().Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_1066_(child, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
}

// 103: decl @lune.@base.@Formatter.FormatterFilter.processSubfile
func (self *Formatter_FormatterFilter) ProcessSubfile(node *Nodes_SubfileNode,_opt LnsAny) {
}

// 108: decl @lune.@base.@Formatter.FormatterFilter.processEnv
func (self *Formatter_FormatterFilter) ProcessEnv(node *Nodes_EnvNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("__envLock")
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 114: decl @lune.@base.@Formatter.FormatterFilter.processBlockSub
func (self *Formatter_FormatterFilter) ProcessBlockSub(node *Nodes_BlockNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    for _, _statement := range( node.FP.Get_stmtList().Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_1066_(statement, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
    self.FP.PopIndent()
    if _switch479 := node.FP.Get_blockKind(); _switch479 == Nodes_BlockKind__LetUnwrap || _switch479 == Nodes_BlockKind__Repeat {
        self.FP.Write("}")
    } else {
        self.FP.Writeln("}")
    }
}

// 135: decl @lune.@base.@Formatter.FormatterFilter.processStmtExp
func (self *Formatter_FormatterFilter) ProcessStmtExp(node *Nodes_StmtExpNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
    {
        _tailComment := node.FP.Get_tailComment()
        if _tailComment != nil {
            tailComment := _tailComment.(*Types_Token)
            self.FP.Write("; ")
            self.FP.Writeln(tailComment.Txt)
        } else {
            self.FP.Writeln(";")
        }
    }
}

// 149: decl @lune.@base.@Formatter.FormatterFilter.processDeclEnum
func (self *Formatter_FormatterFilter) ProcessDeclEnum(node *Nodes_DeclEnumNode,_opt LnsAny) {
    self.FP.Writeln(Lns_getVM().String_format("enum %s {", []LnsAny{node.FP.Get_name().Txt}))
    self.FP.PushIndent(nil)
    for _, _name := range( node.FP.Get_valueNameList().Items ) {
        name := _name.(Types_TokenDownCast).ToTypes_Token()
        self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{name.Txt}))
        self.FP.Writeln(",")
    }
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 164: decl @lune.@base.@Formatter.FormatterFilter.processDeclAlge
func (self *Formatter_FormatterFilter) ProcessDeclAlge(node *Nodes_DeclAlgeNode,_opt LnsAny) {
}

// 172: decl @lune.@base.@Formatter.FormatterFilter.processNewAlgeVal
func (self *Formatter_FormatterFilter) ProcessNewAlgeVal(node *Nodes_NewAlgeValNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    for _, _exp := range( node.FP.Get_paramList().Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_1066_(exp, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
}

// 181: decl @lune.@base.@Formatter.FormatterFilter.outputDeclClass
func (self *Formatter_FormatterFilter) outputDeclClass(protoFlag bool,classType *Ast_TypeInfo,gluePrefix LnsAny,moduleName LnsAny) {
    if classType.FP.Get_accessMode() == Ast_AccessMode__Pub{
        self.FP.Write("pub ")
    }
    if _switch742 := classType.FP.Get_kind(); _switch742 == Ast_TypeInfoKind__Class {
        if Lns_isCondTrue( moduleName){
            self.FP.Write("module ")
        } else { 
            self.FP.Write("class ")
        }
    } else if _switch742 == Ast_TypeInfoKind__IF {
        self.FP.Write("interface ")
    }
    self.FP.Write(classType.FP.Get_rawTxt())
    if classType.FP.Get_itemTypeInfoList().Len() > 0{
        self.FP.Write("<")
        for _index, _genType := range( classType.FP.Get_itemTypeInfoList().Items ) {
            index := _index + 1
            genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index > 1{
                self.FP.Write(",")
            }
            self.FP.Write(genType.FP.GetTxt(nil, nil, nil))
        }
        self.FP.Write(">")
    }
    if moduleName != nil{
        moduleName_5278 := moduleName.(*Types_Token)
        self.FP.Write(" require ")
        self.FP.Write(Lns_getVM().String_format("%s ", []LnsAny{moduleName_5278.Txt}))
        if gluePrefix != nil{
            gluePrefix_5280 := gluePrefix.(string)
            self.FP.Write("glue ")
            self.FP.Write(gluePrefix_5280)
        }
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( classType.FP.HasBase()) ||
        Lns_GetEnv().SetStackVal( classType.FP.Get_interfaceList().Len() > 0) ).(bool){
        self.FP.Write(" extend ")
        if classType.FP.HasBase(){
            self.FP.Write(classType.FP.Get_baseTypeInfo().FP.GetTxt(nil, nil, nil))
        }
        if classType.FP.Get_interfaceList().Len() > 0{
            self.FP.Write("(")
            for _index, _ifType := range( classType.FP.Get_interfaceList().Items ) {
                index := _index + 1
                ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if index > 1{
                    self.FP.Write(",")
                }
                self.FP.Write(ifType.FP.GetTxt(nil, nil, nil))
            }
            self.FP.Write(")")
        }
    }
}

// 241: decl @lune.@base.@Formatter.FormatterFilter.processProtoClass
func (self *Formatter_FormatterFilter) ProcessProtoClass(node *Nodes_ProtoClassNode,_opt LnsAny) {
    self.FP.outputDeclClass(true, node.FP.Get_expType(), nil, nil)
}

// 247: decl @lune.@base.@Formatter.FormatterFilter.processDeclClass
func (self *Formatter_FormatterFilter) ProcessDeclClass(node *Nodes_DeclClassNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.outputDeclClass(false, node.FP.Get_expType(), node.FP.Get_gluePrefix(), node.FP.Get_moduleName())
    self.FP.Writeln("")
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    for _, _stmt := range( node.FP.Get_allStmtList().Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_1066_(stmt, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 321: decl @lune.@base.@Formatter.FormatterFilter.processDeclAdvertise
func (self *Formatter_FormatterFilter) ProcessDeclAdvertise(node *Nodes_DeclAdvertiseNode,_opt LnsAny) {
    self.FP.Writeln(Lns_getVM().String_format("advertise %s;", []LnsAny{node.FP.Get_advInfo().FP.Get_member().FP.Get_name().Txt}))
}

// 327: decl @lune.@base.@Formatter.FormatterFilter.processDeclMember
func (self *Formatter_FormatterFilter) ProcessDeclMember(node *Nodes_DeclMemberNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    if _switch1155 := node.FP.Get_accessMode(); _switch1155 == Ast_AccessMode__Pub || _switch1155 == Ast_AccessMode__Pro || _switch1155 == Ast_AccessMode__None || _switch1155 == Ast_AccessMode__Local {
        self.FP.Write(Ast_accessMode2txt(node.FP.Get_accessMode()))
        self.FP.Write(" ")
    }
    if node.FP.Get_staticFlag(){
        self.FP.Write("static ")
    }
    self.FP.Write("let ")
    var symbol *Ast_SymbolInfo
    symbol = node.FP.Get_symbolInfo()
    if symbol.FP.Get_mutable(){
        self.FP.Write("mut ")
    }
    self.FP.Write(symbol.FP.Get_name())
    self.FP.Write(":")
    Formatter_filter_1066_(&node.FP.Get_refType().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    if node.FP.Get_getterMode() == Ast_AccessMode__None{
        if node.FP.Get_setterMode() != Ast_AccessMode__None{
            self.FP.Write("{none,")
            self.FP.Write(Ast_accessMode2txt(node.FP.Get_setterMode()))
            self.FP.Write("}")
        }
    } else { 
        self.FP.Write("{")
        self.FP.Write(Ast_accessMode2txt(node.FP.Get_getterMode()))
        if Lns_op_not(node.FP.Get_getterMutable()){
            self.FP.Write("&")
        }
        if node.FP.Get_getterRetType() != symbol.FP.Get_typeInfo(){
            self.FP.Write(":")
            self.FP.Write(node.FP.Get_getterRetType().FP.GetTxt(nil, nil, nil))
        }
        if node.FP.Get_setterMode() != Ast_AccessMode__None{
            self.FP.Write(",")
            self.FP.Write(Ast_accessMode2txt(node.FP.Get_setterMode()))
        }
        self.FP.Write("}")
    }
    self.FP.Writeln(";")
}

// 377: decl @lune.@base.@Formatter.FormatterFilter.processExpMacroExp
func (self *Formatter_FormatterFilter) ProcessExpMacroExp(node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList()
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_1066_(stmt, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
}

// 386: decl @lune.@base.@Formatter.FormatterFilter.processDeclMacro
func (self *Formatter_FormatterFilter) ProcessDeclMacro(node *Nodes_DeclMacroNode,_opt LnsAny) {
}

// 391: decl @lune.@base.@Formatter.FormatterFilter.processExpMacroStat
func (self *Formatter_FormatterFilter) ProcessExpMacroStat(node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    for _, _strNode := range( node.FP.Get_expStrList().Items ) {
        strNode := _strNode.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_1066_(strNode, self, opt.FP.NextOpt(strNode))
    }
}

// 401: decl @lune.@base.@Formatter.FormatterFilter.processUnwrapSet
func (self *Formatter_FormatterFilter) ProcessUnwrapSet(node *Nodes_UnwrapSetNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(&node.FP.Get_dstExpList().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    Formatter_filter_1066_(&node.FP.Get_srcExpList().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    if Lns_isCondTrue( node.FP.Get_unwrapBlock()){
        Formatter_filter_1066_(&Lns_unwrap( node.FP.Get_unwrapBlock()).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
}

// 413: decl @lune.@base.@Formatter.FormatterFilter.processIfUnwrap
func (self *Formatter_FormatterFilter) ProcessIfUnwrap(node *Nodes_IfUnwrapNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("if! ")
    if node.FP.Get_varSymList().Len() != 0{
        self.FP.Write("let ")
        for _index, _varSym := range( node.FP.Get_varSymList().Items ) {
            index := _index + 1
            varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if index > 1{
                self.FP.Write(",")
            }
            self.FP.Write(varSym.FP.Get_name())
        }
        self.FP.Write(" = ")
    }
    Formatter_filter_1066_(&node.FP.Get_expList().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    if Lns_isCondTrue( node.FP.Get_nilBlock()){
        self.FP.Write("else ")
        Formatter_filter_1066_(&Lns_unwrap( node.FP.Get_nilBlock()).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
}

// 435: decl @lune.@base.@Formatter.FormatterFilter.processWhen
func (self *Formatter_FormatterFilter) ProcessWhen(node *Nodes_WhenNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("when!")
    var symTxt string
    symTxt = " "
    for _, _symPair := range( node.FP.Get_symPairList().Items ) {
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        symTxt = Lns_getVM().String_format("%s%s ", []LnsAny{symTxt, symPair.FP.Get_src().FP.Get_name()})
        
    }
    self.FP.Write(symTxt)
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    {
        __exp := node.FP.Get_elseBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write("else ")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
}

// 454: decl @lune.@base.@Formatter.FormatterFilter.processDeclVar
func (self *Formatter_FormatterFilter) ProcessDeclVar(node *Nodes_DeclVarNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    if _switch1853 := node.FP.Get_accessMode(); _switch1853 == Ast_AccessMode__Pub {
        self.FP.Write("pub ")
    }
    if _switch1896 := node.FP.Get_mode(); _switch1896 == Nodes_DeclVarMode__Let {
        if node.FP.Get_unwrapFlag(){
            self.FP.Write("let! ")
        } else { 
            self.FP.Write("let ")
        }
    } else if _switch1896 == Nodes_DeclVarMode__Unwrap {
        self.FP.Write("unwrap! ")
    }
    for _index, _symInfo := range( node.FP.Get_symbolInfoList().Items ) {
        index := _index + 1
        symInfo := _symInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(", ")
        }
        if symInfo.FP.Get_mutable(){
            self.FP.Write("mut ")
        }
        self.FP.Write(symInfo.FP.Get_name())
        if node.FP.Get_varList().Len() >= index{
            var varInfo *Nodes_VarInfo
            varInfo = node.FP.Get_varList().GetAt(index).(Nodes_VarInfoDownCast).ToNodes_VarInfo()
            {
                _varType := varInfo.FP.Get_refType()
                if _varType != nil {
                    varType := _varType.(*Nodes_RefTypeNode)
                    self.FP.Write(":")
                    Formatter_filter_1066_(&varType.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
                }
            }
        }
    }
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(" = ")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
    {
        __exp := node.FP.Get_unwrapBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write(" ")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
    {
        __exp := node.FP.Get_thenBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write("then")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
    {
        __exp := node.FP.Get_syncBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write("do")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
    self.FP.Writeln(";")
}

// 512: decl @lune.@base.@Formatter.FormatterFilter.processDeclArg
func (self *Formatter_FormatterFilter) ProcessDeclArg(node *Nodes_DeclArgNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(node.FP.Get_symbolInfo().FP.Get_name())
    self.FP.Write(":")
    {
        _refType := node.FP.Get_argType()
        if _refType != nil {
            refType := _refType.(*Nodes_RefTypeNode)
            Formatter_filter_1066_(&refType.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        } else {
            self.FP.Write(node.FP.Get_expType().FP.GetTxt(nil, nil, nil))
        }
    }
}

// 525: decl @lune.@base.@Formatter.FormatterFilter.processDeclArgDDD
func (self *Formatter_FormatterFilter) ProcessDeclArgDDD(node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    self.FP.Write("...")
}

// 530: decl @lune.@base.@Formatter.FormatterFilter.processExpSubDDD
func (self *Formatter_FormatterFilter) ProcessExpSubDDD(node *Nodes_ExpSubDDDNode,_opt LnsAny) {
}

// 537: decl @lune.@base.@Formatter.FormatterFilter.processDeclForm
func (self *Formatter_FormatterFilter) ProcessDeclForm(node *Nodes_DeclFormNode,_opt LnsAny) {
}

// 542: decl @lune.@base.@Formatter.FormatterFilter.processDeclFuncInfo
func (self *Formatter_FormatterFilter) processDeclFuncInfo(node *Nodes_Node,declInfo *Nodes_DeclFuncInfo,opt *Formatter_Opt) {
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_expType()
    if funcType.FP.Get_accessMode() == Ast_AccessMode__Pub{
        self.FP.Write("pub ")
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( declInfo.FP.Get_staticFlag()) &&
        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(declInfo.FP.Get_name()) && 
        Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Types_Token).Txt)) == "__init") ).(bool)){
        self.FP.Write("__init")
    } else { 
        self.FP.Write("fn ")
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( opt.FP.Get_parent().FP.Get_kind() != Nodes_NodeKind_get_DeclClass()) &&
            Lns_GetEnv().SetStackVal( funcType.FP.Get_kind() == Ast_TypeInfoKind__Method) ).(bool)){
            var classType *Ast_TypeInfo
            classType = funcType.FP.Get_parentInfo()
            self.FP.Write(classType.FP.Get_rawTxt())
            self.FP.Write(".")
        }
        {
            _nameToken := declInfo.FP.Get_name()
            if _nameToken != nil {
                nameToken := _nameToken.(*Types_Token)
                self.FP.Write(nameToken.Txt)
            }
        }
        self.FP.Write("(")
        var argList *LnsList
        argList = declInfo.FP.Get_argList()
        if argList.Len() != 0{
            self.FP.Write(" ")
        }
        for _index, _arg := range( argList.Items ) {
            index := _index + 1
            arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
            if index > 1{
                self.FP.Write(", ")
            }
            Formatter_filter_1066_(arg, self, opt.FP.NextOpt(node))
        }
        if argList.Len() != 0{
            self.FP.Write(" ")
        }
        self.FP.Write(")")
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(funcType)) &&
        Lns_GetEnv().SetStackVal( declInfo.FP.Get_kind() == Nodes_FuncKind__Mtd) ).(bool)){
        self.FP.Write(" mut")
    }
    if funcType.FP.Get_retTypeInfoList().Len() != 0{
        self.FP.Write(" : ")
        for _index, _retType := range( funcType.FP.Get_retTypeInfoList().Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index > 1{
                self.FP.Write(", ")
            }
            self.FP.Write(retType.FP.GetTxt(nil, nil, nil))
        }
    }
    {
        __exp := declInfo.FP.Get_body()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write(" ")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(node))
        } else {
            self.FP.Writeln(";")
        }
    }
}

// 606: decl @lune.@base.@Formatter.FormatterFilter.processDeclFunc
func (self *Formatter_FormatterFilter) ProcessDeclFunc(node *Nodes_DeclFuncNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.processDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo(), opt)
}

// 611: decl @lune.@base.@Formatter.FormatterFilter.processDeclMethod
func (self *Formatter_FormatterFilter) ProcessDeclMethod(node *Nodes_DeclMethodNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.processDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo(), opt)
}

// 616: decl @lune.@base.@Formatter.FormatterFilter.processDeclConstr
func (self *Formatter_FormatterFilter) ProcessDeclConstr(node *Nodes_DeclConstrNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.processDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo(), opt)
}

// 622: decl @lune.@base.@Formatter.FormatterFilter.processDeclDestr
func (self *Formatter_FormatterFilter) ProcessDeclDestr(node *Nodes_DeclDestrNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.processDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo(), opt)
}

// 628: decl @lune.@base.@Formatter.FormatterFilter.processExpCallSuperCtor
func (self *Formatter_FormatterFilter) ProcessExpCallSuperCtor(node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("super(")
    {
        _expListNode := node.FP.Get_expList()
        if _expListNode != nil {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            Formatter_filter_1066_(&expListNode.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
    self.FP.Writeln(");")
}

// 638: decl @lune.@base.@Formatter.FormatterFilter.processExpCallSuper
func (self *Formatter_FormatterFilter) ProcessExpCallSuper(node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("super(")
    {
        _expListNode := node.FP.Get_expList()
        if _expListNode != nil {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            Formatter_filter_1066_(&expListNode.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
    self.FP.Writeln(")")
}

// 648: decl @lune.@base.@Formatter.FormatterFilter.processRefType
func (self *Formatter_FormatterFilter) ProcessRefType(node *Nodes_RefTypeNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    if Lns_op_not(Ast_TypeInfo_isMut(node.FP.Get_expType())){
        self.FP.Write("&")
    }
    Formatter_filter_1066_(node.FP.Get_name(), self, opt.FP.NextOpt(&node.Nodes_Node))
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType().FP.Get_nonnilableType()
    if _switch2975 := expType.FP.Get_kind(); _switch2975 == Ast_TypeInfoKind__List || _switch2975 == Ast_TypeInfoKind__Set {
        self.FP.Write("<")
        Formatter_filter_1066_(node.FP.Get_itemNodeList().GetAt(1).(Nodes_NodeDownCast).ToNodes_Node(), self, opt.FP.NextOpt(&node.Nodes_Node))
        self.FP.Write(">")
    } else if _switch2975 == Ast_TypeInfoKind__Map {
        self.FP.Write("<")
        Formatter_filter_1066_(node.FP.Get_itemNodeList().GetAt(1).(Nodes_NodeDownCast).ToNodes_Node(), self, opt.FP.NextOpt(&node.Nodes_Node))
        self.FP.Write(",")
        Formatter_filter_1066_(node.FP.Get_itemNodeList().GetAt(2).(Nodes_NodeDownCast).ToNodes_Node(), self, opt.FP.NextOpt(&node.Nodes_Node))
        self.FP.Write(">")
    }
    if node.FP.Get_expType().FP.Get_nilable(){
        self.FP.Write("!")
    }
}

// 678: decl @lune.@base.@Formatter.FormatterFilter.processIf
func (self *Formatter_FormatterFilter) ProcessIf(node *Nodes_IfNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList()
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if _switch3050 := stmt.FP.Get_kind(); _switch3050 == Nodes_IfKind__If {
            self.FP.Write("if ")
        } else if _switch3050 == Nodes_IfKind__ElseIf {
            self.FP.Write("elseif ")
        } else if _switch3050 == Nodes_IfKind__Else {
            self.FP.Write("else ")
        }
        Formatter_filter_1066_(stmt.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
        self.FP.Write(" ")
        Formatter_filter_1066_(&stmt.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
}

// 699: decl @lune.@base.@Formatter.FormatterFilter.processSwitch
func (self *Formatter_FormatterFilter) ProcessSwitch(node *Nodes_SwitchNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
    var caseList *LnsList
    caseList = node.FP.Get_caseList()
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        Formatter_filter_1066_(&caseInfo.FP.Get_expList().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        Formatter_filter_1066_(&caseInfo.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
    {
        __exp := node.FP.Get_default()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
}

// 713: decl @lune.@base.@Formatter.FormatterFilter.processMatch
func (self *Formatter_FormatterFilter) ProcessMatch(node *Nodes_MatchNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_val(), self, opt.FP.NextOpt(&node.Nodes_Node))
    var caseList *LnsList
    caseList = node.FP.Get_caseList()
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        Formatter_filter_1066_(&caseInfo.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
    {
        __exp := node.FP.Get_defaultBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            Formatter_filter_1066_(_exp, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
}

// 726: decl @lune.@base.@Formatter.FormatterFilter.processWhile
func (self *Formatter_FormatterFilter) ProcessWhile(node *Nodes_WhileNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("while ")
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Write(" ")
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 734: decl @lune.@base.@Formatter.FormatterFilter.processRepeat
func (self *Formatter_FormatterFilter) ProcessRepeat(node *Nodes_RepeatNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("repeat ")
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Write(" ")
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Writeln(";")
}

// 743: decl @lune.@base.@Formatter.FormatterFilter.processFor
func (self *Formatter_FormatterFilter) ProcessFor(node *Nodes_ForNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(Lns_getVM().String_format("for %s = ", []LnsAny{node.FP.Get_val().FP.Get_name()}))
    Formatter_filter_1066_(node.FP.Get_init(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Write(", ")
    Formatter_filter_1066_(node.FP.Get_to(), self, opt.FP.NextOpt(&node.Nodes_Node))
    {
        __exp := node.FP.Get_delta()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write(", ")
            Formatter_filter_1066_(_exp, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
    self.FP.Write(" ")
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 757: decl @lune.@base.@Formatter.FormatterFilter.processApply
func (self *Formatter_FormatterFilter) ProcessApply(node *Nodes_ApplyNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("apply ")
    for _index, __var := range( node.FP.Get_varList().Items ) {
        index := _index + 1
        _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(", ")
        }
        self.FP.Write(_var.FP.Get_name())
    }
    self.FP.Write(" of ")
    Formatter_filter_1066_(&node.FP.Get_expList().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Write(" ")
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 773: decl @lune.@base.@Formatter.FormatterFilter.processForeach
func (self *Formatter_FormatterFilter) ProcessForeach(node *Nodes_ForeachNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("foreach ")
    self.FP.Write(node.FP.Get_val().FP.Get_name())
    {
        _key := node.FP.Get_key()
        if _key != nil {
            key := _key.(*Ast_SymbolInfo)
            self.FP.Write(", ")
            self.FP.Write(key.FP.Get_name())
        }
    }
    self.FP.Write(" in ")
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Writeln("")
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 787: decl @lune.@base.@Formatter.FormatterFilter.processForsort
func (self *Formatter_FormatterFilter) ProcessForsort(node *Nodes_ForsortNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("forsort ")
    self.FP.Write(node.FP.Get_val().FP.Get_name())
    {
        _key := node.FP.Get_key()
        if _key != nil {
            key := _key.(*Ast_SymbolInfo)
            self.FP.Write(", ")
            self.FP.Write(key.FP.Get_name())
        }
    }
    self.FP.Write(" in ")
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Writeln("")
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 802: decl @lune.@base.@Formatter.FormatterFilter.processExpUnwrap
func (self *Formatter_FormatterFilter) ProcessExpUnwrap(node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("unwrap ")
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
    {
        __exp := node.FP.Get_default()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write(" default ")
            Formatter_filter_1066_(_exp, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
}

// 824: decl @lune.@base.@Formatter.FormatterFilter.processExpCall
func (self *Formatter_FormatterFilter) ProcessExpCall(node *Nodes_ExpCallNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_func(), self, opt.FP.NextOpt(&node.Nodes_Node))
    if node.FP.Get_nilAccess(){
        self.FP.Write("$(")
    } else { 
        self.FP.Write("(")
    }
    {
        __exp := node.FP.Get_argList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(" ")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
            self.FP.Write(" ")
        }
    }
    self.FP.Write(")")
}

// 841: decl @lune.@base.@Formatter.FormatterFilter.processExpList
func (self *Formatter_FormatterFilter) ProcessExpList(node *Nodes_ExpListNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    var expList *LnsList
    expList = node.FP.Get_expList()
    for _index, _exp := range( expList.Items ) {
        index := _index + 1
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        if index > 1{
            if exp.FP.Get_kind() == Nodes_NodeKind_get_ExpAccessMRet(){
                break
            }
            if exp.FP.Get_expType().FP.Get_kind() != Ast_TypeInfoKind__Abbr{
                self.FP.Write(", ")
            } else { 
                self.FP.Write(" ")
            }
        }
        Formatter_filter_1066_(exp, self, opt.FP.NextOpt(&node.Nodes_Node))
    }
}

// 862: decl @lune.@base.@Formatter.FormatterFilter.processExpMRet
func (self *Formatter_FormatterFilter) ProcessExpMRet(node *Nodes_ExpMRetNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_mRet(), self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 868: decl @lune.@base.@Formatter.FormatterFilter.processExpAccessMRet
func (self *Formatter_FormatterFilter) ProcessExpAccessMRet(node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
}

// 874: decl @lune.@base.@Formatter.FormatterFilter.processExpOp1
func (self *Formatter_FormatterFilter) ProcessExpOp1(node *Nodes_ExpOp1Node,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(node.FP.Get_op().Txt)
    if _switch4286 := node.FP.Get_op().Txt; _switch4286 == "not" {
        self.FP.Write(" ")
    }
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 886: decl @lune.@base.@Formatter.FormatterFilter.processExpToDDD
func (self *Formatter_FormatterFilter) ProcessExpToDDD(node *Nodes_ExpToDDDNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(&node.FP.Get_expList().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 893: decl @lune.@base.@Formatter.FormatterFilter.processExpMultiTo1
func (self *Formatter_FormatterFilter) ProcessExpMultiTo1(node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 900: decl @lune.@base.@Formatter.FormatterFilter.processExpCast
func (self *Formatter_FormatterFilter) ProcessExpCast(node *Nodes_ExpCastNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 906: decl @lune.@base.@Formatter.FormatterFilter.processExpParen
func (self *Formatter_FormatterFilter) ProcessExpParen(node *Nodes_ExpParenNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("(")
    Formatter_filter_1066_(node.FP.Get_exp(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Write(")")
}

// 913: decl @lune.@base.@Formatter.FormatterFilter.processExpSetVal
func (self *Formatter_FormatterFilter) ProcessExpSetVal(node *Nodes_ExpSetValNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_exp1(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Write(" = ")
    Formatter_filter_1066_(&node.FP.Get_exp2().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 920: decl @lune.@base.@Formatter.FormatterFilter.processExpSetItem
func (self *Formatter_FormatterFilter) ProcessExpSetItem(node *Nodes_ExpSetItemNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_val(), self, opt.FP.NextOpt(&node.Nodes_Node))
    switch _exp4590 := node.FP.Get_index().(type) {
    case *Nodes_IndexVal__NodeIdx:
    index := _exp4590.Val1
        self.FP.Write("[")
        Formatter_filter_1066_(index, self, opt.FP.NextOpt(&node.Nodes_Node))
        self.FP.Write("]")
    case *Nodes_IndexVal__SymIdx:
    index := _exp4590.Val1
        self.FP.Write(Lns_getVM().String_format(".%s", []LnsAny{index}))
    }
    self.FP.Write(" = ")
    Formatter_filter_1066_(node.FP.Get_exp2(), self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 937: decl @lune.@base.@Formatter.FormatterFilter.processExpOp2
func (self *Formatter_FormatterFilter) ProcessExpOp2(node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_exp1(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Write(Lns_getVM().String_format(" %s ", []LnsAny{node.FP.Get_op().Txt}))
    Formatter_filter_1066_(node.FP.Get_exp2(), self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 944: decl @lune.@base.@Formatter.FormatterFilter.processExpNew
func (self *Formatter_FormatterFilter) ProcessExpNew(node *Nodes_ExpNewNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("new ")
    Formatter_filter_1066_(node.FP.Get_symbol(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Write("(")
    {
        __exp := node.FP.Get_argList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
    self.FP.Write(")")
}

// 956: decl @lune.@base.@Formatter.FormatterFilter.processExpRef
func (self *Formatter_FormatterFilter) ProcessExpRef(node *Nodes_ExpRefNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_symbolInfo().FP.Get_name())
}

// 961: decl @lune.@base.@Formatter.FormatterFilter.processExpRefItem
func (self *Formatter_FormatterFilter) ProcessExpRefItem(node *Nodes_ExpRefItemNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_val(), self, opt.FP.NextOpt(&node.Nodes_Node))
    if node.FP.Get_nilAccess(){
        self.FP.Write("$")
    }
    {
        __exp := node.FP.Get_index()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write("[ ")
            Formatter_filter_1066_(_exp, self, opt.FP.NextOpt(&node.Nodes_Node))
            self.FP.Write(" ]")
        } else {
            {
                __exp := node.FP.Get_symbol()
                if __exp != nil {
                    _exp := __exp.(string)
                    self.FP.Write(".")
                    self.FP.Write(_exp)
                }
            }
        }
    }
}

// 980: decl @lune.@base.@Formatter.FormatterFilter.processRefField
func (self *Formatter_FormatterFilter) ProcessRefField(node *Nodes_RefFieldNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_prefix(), self, opt.FP.NextOpt(&node.Nodes_Node))
    if node.FP.Get_nilAccess(){
        self.FP.Write("$.")
    } else { 
        self.FP.Write(".")
    }
    self.FP.Write(node.FP.Get_field().Txt)
}

// 992: decl @lune.@base.@Formatter.FormatterFilter.processExpOmitEnum
func (self *Formatter_FormatterFilter) ProcessExpOmitEnum(node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    self.FP.Write(".")
    self.FP.Write(node.FP.Get_valToken().Txt)
}

// 999: decl @lune.@base.@Formatter.FormatterFilter.processGetField
func (self *Formatter_FormatterFilter) ProcessGetField(node *Nodes_GetFieldNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_prefix(), self, opt.FP.NextOpt(&node.Nodes_Node))
    if node.FP.Get_nilAccess(){
        self.FP.Write("$")
    }
    self.FP.Write(".$")
    self.FP.Write(node.FP.Get_field().Txt)
}

// 1010: decl @lune.@base.@Formatter.FormatterFilter.processReturn
func (self *Formatter_FormatterFilter) ProcessReturn(node *Nodes_ReturnNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("return")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(" ")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
    self.FP.Writeln(";")
}

// 1022: decl @lune.@base.@Formatter.FormatterFilter.processProvide
func (self *Formatter_FormatterFilter) ProcessProvide(node *Nodes_ProvideNode,_opt LnsAny) {
}

// 1028: decl @lune.@base.@Formatter.FormatterFilter.processAlias
func (self *Formatter_FormatterFilter) ProcessAlias(node *Nodes_AliasNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(Lns_getVM().String_format("alias %s = ", []LnsAny{node.FP.Get_newSymbol().FP.Get_name()}))
    Formatter_filter_1066_(node.FP.Get_srcNode(), self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Writeln(";")
}

// 1036: decl @lune.@base.@Formatter.FormatterFilter.processTestCase
func (self *Formatter_FormatterFilter) ProcessTestCase(node *Nodes_TestCaseNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(Lns_getVM().String_format("__test %s(%s) {", []LnsAny{node.FP.Get_name().Txt, node.FP.Get_ctrlName()}))
    Formatter_filter_1066_(&node.FP.Get_block().Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
    self.FP.Write("}")
}

// 1044: decl @lune.@base.@Formatter.FormatterFilter.processBoxing
func (self *Formatter_FormatterFilter) ProcessBoxing(node *Nodes_BoxingNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_src(), self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 1051: decl @lune.@base.@Formatter.FormatterFilter.processUnboxing
func (self *Formatter_FormatterFilter) ProcessUnboxing(node *Nodes_UnboxingNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_1066_(node.FP.Get_src(), self, opt.FP.NextOpt(&node.Nodes_Node))
}

// 1058: decl @lune.@base.@Formatter.FormatterFilter.processLiteralList
func (self *Formatter_FormatterFilter) ProcessLiteralList(node *Nodes_LiteralListNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("[")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(" ")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
            self.FP.Write(" ")
        }
    }
    self.FP.Write("]")
}

// 1071: decl @lune.@base.@Formatter.FormatterFilter.processLiteralSet
func (self *Formatter_FormatterFilter) ProcessLiteralSet(node *Nodes_LiteralSetNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("(@")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(" ")
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
            self.FP.Write(" ")
        }
    }
    self.FP.Write(")")
}

// 1084: decl @lune.@base.@Formatter.FormatterFilter.processLiteralMap
func (self *Formatter_FormatterFilter) ProcessLiteralMap(node *Nodes_LiteralMapNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write("{")
    var pairList *LnsList
    pairList = node.FP.Get_pairList()
    for _index, _pair := range( pairList.Items ) {
        index := _index + 1
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        if index > 1{
            self.FP.Write(", ")
        }
        Formatter_filter_1066_(pair.FP.Get_key(), self, opt.FP.NextOpt(&node.Nodes_Node))
        self.FP.Write(":")
        Formatter_filter_1066_(pair.FP.Get_val(), self, opt.FP.NextOpt(&node.Nodes_Node))
    }
    self.FP.Write("}")
}

// 1102: decl @lune.@base.@Formatter.FormatterFilter.processLiteralArray
func (self *Formatter_FormatterFilter) ProcessLiteralArray(node *Nodes_LiteralArrayNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            Formatter_filter_1066_(&_exp.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
        }
    }
}

// 1110: decl @lune.@base.@Formatter.FormatterFilter.processLiteralChar
func (self *Formatter_FormatterFilter) ProcessLiteralChar(node *Nodes_LiteralCharNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("?%s", []LnsAny{node.FP.Get_token().Txt}))
}

// 1115: decl @lune.@base.@Formatter.FormatterFilter.processLiteralInt
func (self *Formatter_FormatterFilter) ProcessLiteralInt(node *Nodes_LiteralIntNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 1120: decl @lune.@base.@Formatter.FormatterFilter.processLiteralReal
func (self *Formatter_FormatterFilter) ProcessLiteralReal(node *Nodes_LiteralRealNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 1125: decl @lune.@base.@Formatter.FormatterFilter.processLiteralString
func (self *Formatter_FormatterFilter) ProcessLiteralString(node *Nodes_LiteralStringNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(node.FP.Get_token().Txt)
    {
        _expList := node.FP.Get_orgParam()
        if _expList != nil {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.Write(" ( ")
            Formatter_filter_1066_(&expList.Nodes_Node, self, opt.FP.NextOpt(&node.Nodes_Node))
            self.FP.Write(" )")
        }
    }
}

// 1137: decl @lune.@base.@Formatter.FormatterFilter.processLiteralBool
func (self *Formatter_FormatterFilter) ProcessLiteralBool(node *Nodes_LiteralBoolNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 1142: decl @lune.@base.@Formatter.FormatterFilter.processLiteralNil
func (self *Formatter_FormatterFilter) ProcessLiteralNil(node *Nodes_LiteralNilNode,_opt LnsAny) {
    self.FP.Write("nil")
}

// 1147: decl @lune.@base.@Formatter.FormatterFilter.processBreak
func (self *Formatter_FormatterFilter) ProcessBreak(node *Nodes_BreakNode,_opt LnsAny) {
    self.FP.Writeln("break;")
}

// 1153: decl @lune.@base.@Formatter.FormatterFilter.processLiteralSymbol
func (self *Formatter_FormatterFilter) ProcessLiteralSymbol(node *Nodes_LiteralSymbolNode,_opt LnsAny) {
}

// 1159: decl @lune.@base.@Formatter.FormatterFilter.processAbbr
func (self *Formatter_FormatterFilter) ProcessAbbr(node *Nodes_AbbrNode,_opt LnsAny) {
    self.FP.Write("##")
}


func Lns_Formatter_init() {
    if init_Formatter { return }
    init_Formatter = true
    Formatter__mod__ = "@lune.@base.@Formatter"
    Lns_InitMod()
    Lns_Ast_init()
    Lns_Nodes_init()
    Lns_Parser_init()
    Lns_Util_init()
}
func init() {
    init_Formatter = false
}
