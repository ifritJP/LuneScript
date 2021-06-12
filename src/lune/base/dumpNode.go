// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_dumpNode bool
var dumpNode__mod__ string
// for 62
func dumpNode_convExp139(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 475
func dumpNode_convExp2537(arg1 []LnsAny) (string, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// 102: decl @lune.@base.@dumpNode.createFilter
func DumpNode_createFilter(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,processInfo *Ast_ProcessInfo,stream Lns_oStream) *Nodes_Filter {
    return &NewdumpNode_dumpFilter(_env, true, moduleTypeInfo, moduleTypeInfo.FP.Get_scope(_env), stream, processInfo).Nodes_Filter
}

// 109: decl @lune.@base.@dumpNode.filter
func dumpNode_filter_3_(_env *LnsEnv, node *Nodes_Node,filter *dumpNode_dumpFilter,opt *DumpNode_Opt) {
    node.FP.ProcessFilter(_env, &filter.Nodes_Filter, DumpNode_Opt2Stem(opt))
}

// 558: decl @lune.@base.@dumpNode.getTypeListTxt
func dumpNode_getTypeListTxt_4_(_env *LnsEnv, typeList *LnsList) string {
    var txt string
    txt = ""
    for _index, _typeInfo := range( typeList.Items ) {
        index := _index + 1
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            txt = txt + ", "
            
        }
        txt = txt + typeInfo.FP.GetTxt(_env, nil, nil, nil)
        
    }
    return txt
}


// declaration Class -- Opt
type DumpNode_OptMtd interface {
    Get(_env *LnsEnv)(string, LnsInt)
    NextOpt(_env *LnsEnv) *DumpNode_Opt
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
func NewDumpNode_Opt(_env *LnsEnv, arg1 string, arg2 LnsInt) *DumpNode_Opt {
    obj := &DumpNode_Opt{}
    obj.FP = obj
    obj.InitDumpNode_Opt(_env, arg1, arg2)
    return obj
}
// 39: DeclConstr
func (self *DumpNode_Opt) InitDumpNode_Opt(_env *LnsEnv, prefix string,depth LnsInt) {
    self.prefix = prefix
    
    self.depth = depth
    
    self.next = nil
    
}

// 44: decl @lune.@base.@dumpNode.Opt.get
func (self *DumpNode_Opt) Get(_env *LnsEnv)(string, LnsInt) {
    return self.prefix, self.depth
}

// 47: decl @lune.@base.@dumpNode.Opt.nextOpt
func (self *DumpNode_Opt) NextOpt(_env *LnsEnv) *DumpNode_Opt {
    {
        __exp := self.next
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*DumpNode_Opt)
            return _exp
        }
    }
    var opt *DumpNode_Opt
    opt = NewDumpNode_Opt(_env, self.prefix + "  ", self.depth + 1)
    self.next = opt
    
    return opt
}


// declaration Class -- dumpFilter
type dumpNode_dumpFilterMtd interface {
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    dump(_env *LnsEnv, arg1 *DumpNode_Opt, arg2 *Nodes_Node, arg3 string)
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
    processDeclFuncInfo(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo, arg3 *DumpNode_Opt)
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
    writeln(_env *LnsEnv, arg1 *DumpNode_Opt, arg2 string)
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
func NewdumpNode_dumpFilter(_env *LnsEnv, arg1 bool, arg2 LnsAny, arg3 LnsAny, arg4 Lns_oStream, arg5 *Ast_ProcessInfo) *dumpNode_dumpFilter {
    obj := &dumpNode_dumpFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitdumpNode_dumpFilter(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *dumpNode_dumpFilter) InitdumpNode_dumpFilter(_env *LnsEnv, arg1 bool, arg2 LnsAny, arg3 LnsAny, arg4 Lns_oStream, arg5 *Ast_ProcessInfo) {
    self.Nodes_Filter.InitNodes_Filter( _env, arg1,arg2,arg3)
    self.stream = arg4
    self.processInfo = arg5
}
// 61: decl @lune.@base.@dumpNode.dumpFilter.writeln
func (self *dumpNode_dumpFilter) writeln(_env *LnsEnv, opt *DumpNode_Opt,txt string) {
    var prefix string
    prefix = dumpNode_convExp139(Lns_2DDD(opt.FP.Get(_env)))
    self.stream.Write(_env, _env.LuaVM.String_format("%s%s\n", []LnsAny{prefix, Lns_car(_env.LuaVM.String_gsub(txt," *$", "")).(string)}))
}

// 65: decl @lune.@base.@dumpNode.dumpFilter.dump
func (self *dumpNode_dumpFilter) dump(_env *LnsEnv, opt *DumpNode_Opt,node *Nodes_Node,txt string) {
    var typeStr string
    typeStr = ""
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType(_env)
    if expType.FP.Equals(_env, self.processInfo, Ast_builtinTypeNone, nil, nil){
        typeStr = _env.LuaVM.String_format("(%d:%s:%s)", []LnsAny{expType.FP.Get_typeId(_env).Id, expType.FP.GetTxt(_env, nil, nil, nil), expType.FP.Get_kind(_env)})
        
    }
    var comment string
    {
        _commentList := node.FP.Get_commentList(_env)
        if !Lns_IsNil( _commentList ) {
            commentList := _commentList.(*LnsList)
            comment = _env.LuaVM.String_format("comment:%d,%d", []LnsAny{commentList.Len(), _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( node.FP.Get_tailComment(_env)) &&
                _env.SetStackVal( 1) ||
                _env.SetStackVal( 0) ).(LnsInt)})
            
        } else {
            if Lns_isCondTrue( node.FP.Get_tailComment(_env)){
                comment = "comment:0,1"
                
            } else { 
                comment = ""
                
            }
        }
    }
    var attrib string
    attrib = ""
    if node.FP.HasNilAccess(_env){
        attrib = _env.LuaVM.String_format("%s %s", []LnsAny{attrib, "nilacc"})
        
    }
    if len(attrib) != 0{
        attrib = _env.LuaVM.String_format("[%s]", []LnsAny{attrib})
        
    }
    self.FP.writeln(_env, opt, _env.LuaVM.String_format(": %s:%d(%d:%d) %s %s %s %s", []LnsAny{Nodes_getNodeKindName(_env, node.FP.Get_kind(_env)), node.FP.Get_id(_env), node.FP.Get_effectivePos(_env).LineNo, node.FP.Get_effectivePos(_env).Column, attrib, txt, typeStr, comment}))
}

// 113: decl @lune.@base.@dumpNode.dumpFilter.processNone
func (self *dumpNode_dumpFilter) ProcessNone(_env *LnsEnv, node *Nodes_NoneNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}

// 118: decl @lune.@base.@dumpNode.dumpFilter.processShebang
func (self *dumpNode_dumpFilter) ProcessShebang(_env *LnsEnv, node *Nodes_ShebangNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_cmd(_env))
}

// 123: decl @lune.@base.@dumpNode.dumpFilter.processLuneControl
func (self *dumpNode_dumpFilter) ProcessLuneControl(_env *LnsEnv, node *Nodes_LuneControlNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_pragma(_env).(LnsAlgeVal).GetTxt())
}

// 128: decl @lune.@base.@dumpNode.dumpFilter.processBlankLine
func (self *dumpNode_dumpFilter) ProcessBlankLine(_env *LnsEnv, node *Nodes_BlankLineNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%d", []LnsAny{node.FP.Get_lineNum(_env)}))
}

// 133: decl @lune.@base.@dumpNode.dumpFilter.processLuneKind
func (self *dumpNode_dumpFilter) ProcessLuneKind(_env *LnsEnv, node *Nodes_LuneKindNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}

// 139: decl @lune.@base.@dumpNode.dumpFilter.processImport
func (self *dumpNode_dumpFilter) ProcessImport(_env *LnsEnv, node *Nodes_ImportNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_info(_env).FP.Get_modulePath(_env))
}

// 144: decl @lune.@base.@dumpNode.dumpFilter.processRoot
func (self *dumpNode_dumpFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    for _, _child := range( node.FP.Get_children(_env).Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, child, self, opt.FP.NextOpt(_env))
    }
}

// 152: decl @lune.@base.@dumpNode.dumpFilter.processSubfile
func (self *dumpNode_dumpFilter) ProcessSubfile(_env *LnsEnv, node *Nodes_SubfileNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}

// 158: decl @lune.@base.@dumpNode.dumpFilter.processAsyncLock
func (self *dumpNode_dumpFilter) ProcessAsyncLock(_env *LnsEnv, node *Nodes_AsyncLockNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 165: decl @lune.@base.@dumpNode.dumpFilter.processBlockSub
func (self *dumpNode_dumpFilter) ProcessBlockSub(_env *LnsEnv, node *Nodes_BlockNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    for _, _statement := range( node.FP.Get_stmtList(_env).Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, statement, self, opt.FP.NextOpt(_env))
    }
}

// 174: decl @lune.@base.@dumpNode.dumpFilter.processScope
func (self *dumpNode_dumpFilter) ProcessScope(_env *LnsEnv, node *Nodes_ScopeNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, Nodes_ScopeKind_getTxt( node.FP.Get_scopeKind(_env)))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 181: decl @lune.@base.@dumpNode.dumpFilter.processStmtExp
func (self *dumpNode_dumpFilter) ProcessStmtExp(_env *LnsEnv, node *Nodes_StmtExpNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}

// 188: decl @lune.@base.@dumpNode.dumpFilter.processDeclEnum
func (self *dumpNode_dumpFilter) ProcessDeclEnum(_env *LnsEnv, node *Nodes_DeclEnumNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_name(_env).Txt)
    var enumTypeInfo *Ast_EnumTypeInfo
    enumTypeInfo = Lns_unwrap( (Ast_EnumTypeInfoDownCastF(node.FP.Get_expType(_env).FP))).(*Ast_EnumTypeInfo)
    for _, _name := range( node.FP.Get_valueNameList(_env).Items ) {
        name := _name.(Types_TokenDownCast).ToTypes_Token()
        var valInfo *Ast_EnumValInfo
        valInfo = Lns_unwrap( enumTypeInfo.FP.GetEnumValInfo(_env, name.Txt)).(*Ast_EnumValInfo)
        var val LnsAny
        switch _matchExp1 := valInfo.FP.Get_val(_env).(type) {
        case *Ast_EnumLiteral__Int:
        x := _matchExp1.Val1
            val = x
            
        case *Ast_EnumLiteral__Real:
        x := _matchExp1.Val1
            val = x
            
        case *Ast_EnumLiteral__Str:
        x := _matchExp1.Val1
            val = x
            
        }
        self.FP.writeln(_env, opt, _env.LuaVM.String_format("  %s: %s, %s", []LnsAny{name.Txt, valInfo.FP.Get_val(_env).(LnsAlgeVal).GetTxt(), val}))
    }
}

// 211: decl @lune.@base.@dumpNode.dumpFilter.processDeclAlge
func (self *dumpNode_dumpFilter) ProcessDeclAlge(_env *LnsEnv, node *Nodes_DeclAlgeNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var algeTypeInfo *Ast_AlgeTypeInfo
    algeTypeInfo = node.FP.Get_algeType(_env)
    self.FP.dump(_env, opt, &node.Nodes_Node, algeTypeInfo.FP.Get_rawTxt(_env))
    {
        __forsortCollection1 := algeTypeInfo.FP.Get_valInfoMap(_env)
        __forsortSorted1 := __forsortCollection1.CreateKeyListStr()
        __forsortSorted1.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey1 := range( __forsortSorted1.Items ) {
            valInfo := __forsortCollection1.Items[ ___forsortKey1 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            self.FP.writeln(_env, opt, _env.LuaVM.String_format("  %s: %s", []LnsAny{algeTypeInfo.FP.Get_rawTxt(_env), valInfo.FP.Get_name(_env)}))
        }
    }
}

// 221: decl @lune.@base.@dumpNode.dumpFilter.processNewAlgeVal
func (self *dumpNode_dumpFilter) ProcessNewAlgeVal(_env *LnsEnv, node *Nodes_NewAlgeValNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_name(_env).Txt)
    for _, _exp := range( node.FP.Get_paramList(_env).Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, exp, self, opt.FP.NextOpt(_env))
    }
}

// 230: decl @lune.@base.@dumpNode.dumpFilter.processProtoClass
func (self *dumpNode_dumpFilter) ProcessProtoClass(_env *LnsEnv, node *Nodes_ProtoClassNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_name(_env).Txt)
}

// 236: decl @lune.@base.@dumpNode.dumpFilter.processDeclClass
func (self *dumpNode_dumpFilter) ProcessDeclClass(_env *LnsEnv, node *Nodes_DeclClassNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s (%s)", []LnsAny{node.FP.Get_name(_env).Txt, self.FP.GetFull(_env, node.FP.Get_expType(_env), false)}))
    for _, _field := range( node.FP.Get_fieldList(_env).Items ) {
        field := _field.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, field, self, opt.FP.NextOpt(_env))
    }
}

// 246: decl @lune.@base.@dumpNode.dumpFilter.processDeclMember
func (self *dumpNode_dumpFilter) ProcessDeclMember(_env *LnsEnv, node *Nodes_DeclMemberNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_name(_env).Txt)
    dumpNode_filter_3_(_env, &node.FP.Get_refType(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 252: decl @lune.@base.@dumpNode.dumpFilter.processExpMacroExp
func (self *dumpNode_dumpFilter) ProcessExpMacroExp(_env *LnsEnv, node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList(_env)
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, stmt, self, opt.FP.NextOpt(_env))
    }
}

// 261: decl @lune.@base.@dumpNode.dumpFilter.processDeclMacro
func (self *dumpNode_dumpFilter) ProcessDeclMacro(_env *LnsEnv, node *Nodes_DeclMacroNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))
}

// 266: decl @lune.@base.@dumpNode.dumpFilter.processExpMacroStat
func (self *dumpNode_dumpFilter) ProcessExpMacroStat(_env *LnsEnv, node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))
    for _, _expStr := range( node.FP.Get_expStrList(_env).Items ) {
        expStr := _expStr.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, expStr, self, opt.FP.NextOpt(_env))
    }
}

// 276: decl @lune.@base.@dumpNode.dumpFilter.processUnwrapSet
func (self *dumpNode_dumpFilter) ProcessUnwrapSet(_env *LnsEnv, node *Nodes_UnwrapSetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, &node.FP.Get_dstExpList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_srcExpList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    if Lns_isCondTrue( node.FP.Get_unwrapBlock(_env)){
        dumpNode_filter_3_(_env, &Lns_unwrap( node.FP.Get_unwrapBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt(_env))
    }
}

// 288: decl @lune.@base.@dumpNode.dumpFilter.processIfUnwrap
func (self *dumpNode_dumpFilter) ProcessIfUnwrap(_env *LnsEnv, node *Nodes_IfUnwrapNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    for _, _expNode := range( node.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, expNode, self, opt.FP.NextOpt(_env))
    }
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    if Lns_isCondTrue( node.FP.Get_nilBlock(_env)){
        dumpNode_filter_3_(_env, &Lns_unwrap( node.FP.Get_nilBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt(_env))
    }
}

// 301: decl @lune.@base.@dumpNode.dumpFilter.processWhen
func (self *dumpNode_dumpFilter) ProcessWhen(_env *LnsEnv, node *Nodes_WhenNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var symTxt string
    symTxt = ""
    for _, _symPair := range( node.FP.Get_symPairList(_env).Items ) {
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        symTxt = _env.LuaVM.String_format("%s %s", []LnsAny{symTxt, symPair.FP.Get_src(_env).FP.Get_name(_env)})
        
    }
    self.FP.dump(_env, opt, &node.Nodes_Node, symTxt)
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    {
        __exp := node.FP.Get_elseBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 317: decl @lune.@base.@dumpNode.dumpFilter.processDeclVar
func (self *dumpNode_dumpFilter) ProcessDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var varName string
    varName = ""
    for _index, __var := range( node.FP.Get_varList(_env).Items ) {
        index := _index + 1
        _var := __var.(Nodes_VarInfoDownCast).ToNodes_VarInfo()
        if index > 1{
            varName = varName + ","
            
        }
        varName = _env.LuaVM.String_format("%s %s", []LnsAny{varName, _var.FP.Get_name(_env).Txt})
        
    }
    if Lns_isCondTrue( node.FP.Get_unwrapBlock(_env)){
        varName = "!" + varName
        
    }
    varName = _env.LuaVM.String_format("%s %s", []LnsAny{node.FP.Get_mode(_env), varName})
    
    self.FP.dump(_env, opt, &node.Nodes_Node, varName)
    for _, __var := range( node.FP.Get_varList(_env).Items ) {
        _var := __var.(Nodes_VarInfoDownCast).ToNodes_VarInfo()
        {
            __exp := _var.FP.Get_refType(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_RefTypeNode)
                dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
            }
        }
    }
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
    {
        __exp := node.FP.Get_unwrapBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
    {
        __exp := node.FP.Get_thenBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
    {
        __exp := node.FP.Get_syncBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 356: decl @lune.@base.@dumpNode.dumpFilter.processDeclArg
func (self *dumpNode_dumpFilter) ProcessDeclArg(_env *LnsEnv, node *Nodes_DeclArgNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s:%s", []LnsAny{node.FP.Get_name(_env).Txt, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
}

// 362: decl @lune.@base.@dumpNode.dumpFilter.processDeclArgDDD
func (self *dumpNode_dumpFilter) ProcessDeclArgDDD(_env *LnsEnv, node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "...")
}

// 367: decl @lune.@base.@dumpNode.dumpFilter.processExpSubDDD
func (self *dumpNode_dumpFilter) ProcessExpSubDDD(_env *LnsEnv, node *Nodes_ExpSubDDDNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("... (%d)", []LnsAny{node.FP.Get_remainIndex(_env)}))
}

// 374: decl @lune.@base.@dumpNode.dumpFilter.processDeclForm
func (self *dumpNode_dumpFilter) ProcessDeclForm(_env *LnsEnv, node *Nodes_DeclFormNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil))
}

// 379: decl @lune.@base.@dumpNode.dumpFilter.processDeclFuncInfo
func (self *dumpNode_dumpFilter) processDeclFuncInfo(_env *LnsEnv, node *Nodes_Node,declInfo *Nodes_DeclFuncInfo,opt *DumpNode_Opt) {
    var name string
    
    {
        _name := _env.NilAccFin(_env.NilAccPush(declInfo.FP.Get_name(_env)) && 
        _env.NilAccPush(_env.NilAccPop().(*Types_Token).Txt))
        if _name == nil{
            name = "<anonymous>"
            
        } else {
            name = _name.(string)
        }
    }
    name = node.FP.Get_expType(_env).FP.Get_display_stirng_with(_env, name, nil)
    
    if Ast_TypeInfo_isMut(_env, node.FP.Get_expType(_env)){
        name = name + " mut"
        
    }
    if Lns_isCondTrue( _env.NilAccFin(_env.NilAccPush(node.FP.Get_expType(_env).FP.Get_scope(_env)) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_hasClosureAccess(_env)}))){
        name = name + " closure"
        
    }
    self.FP.dump(_env, opt, node, name)
    var argList *LnsList
    argList = declInfo.FP.Get_argList(_env)
    for _, _arg := range( argList.Items ) {
        arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, arg, self, opt.FP.NextOpt(_env))
    }
    {
        __exp := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 401: decl @lune.@base.@dumpNode.dumpFilter.processDeclFunc
func (self *dumpNode_dumpFilter) ProcessDeclFunc(_env *LnsEnv, node *Nodes_DeclFuncNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}

// 406: decl @lune.@base.@dumpNode.dumpFilter.processDeclMethod
func (self *dumpNode_dumpFilter) ProcessDeclMethod(_env *LnsEnv, node *Nodes_DeclMethodNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}

// 411: decl @lune.@base.@dumpNode.dumpFilter.processProtoMethod
func (self *dumpNode_dumpFilter) ProcessProtoMethod(_env *LnsEnv, node *Nodes_ProtoMethodNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}

// 416: decl @lune.@base.@dumpNode.dumpFilter.processDeclConstr
func (self *dumpNode_dumpFilter) ProcessDeclConstr(_env *LnsEnv, node *Nodes_DeclConstrNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}

// 422: decl @lune.@base.@dumpNode.dumpFilter.processDeclDestr
func (self *dumpNode_dumpFilter) ProcessDeclDestr(_env *LnsEnv, node *Nodes_DeclDestrNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}

// 428: decl @lune.@base.@dumpNode.dumpFilter.processExpCallSuperCtor
func (self *dumpNode_dumpFilter) ProcessExpCallSuperCtor(_env *LnsEnv, node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType(_env)
    self.FP.dump(_env, opt, &node.Nodes_Node, typeInfo.FP.GetTxt(_env, nil, nil, nil))
}

// 434: decl @lune.@base.@dumpNode.dumpFilter.processExpCallSuper
func (self *dumpNode_dumpFilter) ProcessExpCallSuper(_env *LnsEnv, node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType(_env)
    self.FP.dump(_env, opt, &node.Nodes_Node, typeInfo.FP.GetTxt(_env, nil, nil, nil))
}

// 440: decl @lune.@base.@dumpNode.dumpFilter.processRefType
func (self *dumpNode_dumpFilter) ProcessRefType(_env *LnsEnv, node *Nodes_RefTypeNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_name(_env), self, opt.FP.NextOpt(_env))
}

// 446: decl @lune.@base.@dumpNode.dumpFilter.processIf
func (self *dumpNode_dumpFilter) ProcessIf(_env *LnsEnv, node *Nodes_IfNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList(_env)
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if stmt.FP.Get_exp(_env).FP.Get_kind(_env) != Nodes_nodeKindEnum__None{
            dumpNode_filter_3_(_env, stmt.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
        }
        dumpNode_filter_3_(_env, &stmt.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    }
}

// 458: decl @lune.@base.@dumpNode.dumpFilter.processSwitch
func (self *dumpNode_dumpFilter) ProcessSwitch(_env *LnsEnv, node *Nodes_SwitchNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
    var caseList *LnsList
    caseList = node.FP.Get_caseList(_env)
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        dumpNode_filter_3_(_env, &caseInfo.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
        dumpNode_filter_3_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    }
    {
        __exp := node.FP.Get_default(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 472: decl @lune.@base.@dumpNode.dumpFilter.processMatch
func (self *dumpNode_dumpFilter) ProcessMatch(_env *LnsEnv, node *Nodes_MatchNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var prefix string
    var depth LnsInt
    prefix,depth = opt.FP.Get(_env)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_val(_env), self, opt.FP.NextOpt(_env))
    var caseList *LnsList
    caseList = node.FP.Get_caseList(_env)
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        dumpNode_filter_3_(_env, &caseInfo.FP.Get_valExpRef(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
        dumpNode_filter_3_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, NewDumpNode_Opt(_env, prefix + "  " + caseInfo.FP.Get_valInfo(_env).FP.Get_name(_env), depth + 1))
    }
    {
        __exp := node.FP.Get_defaultBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            dumpNode_filter_3_(_env, _exp, self, opt.FP.NextOpt(_env))
        }
    }
}

// 488: decl @lune.@base.@dumpNode.dumpFilter.processWhile
func (self *dumpNode_dumpFilter) ProcessWhile(_env *LnsEnv, node *Nodes_WhileNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 495: decl @lune.@base.@dumpNode.dumpFilter.processRepeat
func (self *dumpNode_dumpFilter) ProcessRepeat(_env *LnsEnv, node *Nodes_RepeatNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}

// 502: decl @lune.@base.@dumpNode.dumpFilter.processFor
func (self *dumpNode_dumpFilter) ProcessFor(_env *LnsEnv, node *Nodes_ForNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_val(_env).FP.Get_name(_env))
    dumpNode_filter_3_(_env, node.FP.Get_init(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, node.FP.Get_to(_env), self, opt.FP.NextOpt(_env))
    {
        __exp := node.FP.Get_delta(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            dumpNode_filter_3_(_env, _exp, self, opt.FP.NextOpt(_env))
        }
    }
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 513: decl @lune.@base.@dumpNode.dumpFilter.processApply
func (self *dumpNode_dumpFilter) ProcessApply(_env *LnsEnv, node *Nodes_ApplyNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var varNames string
    varNames = ""
    var varList *LnsList
    varList = node.FP.Get_varList(_env)
    for _, __var := range( varList.Items ) {
        _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        varNames = varNames + _var.FP.Get_name(_env) + " "
        
    }
    self.FP.dump(_env, opt, &node.Nodes_Node, varNames)
    dumpNode_filter_3_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 525: decl @lune.@base.@dumpNode.dumpFilter.processForeach
func (self *dumpNode_dumpFilter) ProcessForeach(_env *LnsEnv, node *Nodes_ForeachNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var index string
    index = ""
    {
        __exp := node.FP.Get_key(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            index = _exp.FP.Get_name(_env)
            
        }
    }
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_val(_env).FP.Get_name(_env) + " " + index)
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 536: decl @lune.@base.@dumpNode.dumpFilter.processForsort
func (self *dumpNode_dumpFilter) ProcessForsort(_env *LnsEnv, node *Nodes_ForsortNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var index string
    index = ""
    {
        __exp := node.FP.Get_key(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            index = _exp.FP.Get_name(_env)
            
        }
    }
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_val(_env).FP.Get_name(_env) + " " + index)
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 548: decl @lune.@base.@dumpNode.dumpFilter.processExpUnwrap
func (self *dumpNode_dumpFilter) ProcessExpUnwrap(_env *LnsEnv, node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
    {
        __exp := node.FP.Get_default(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            dumpNode_filter_3_(_env, _exp, self, opt.FP.NextOpt(_env))
        }
    }
}

// 569: decl @lune.@base.@dumpNode.dumpFilter.processExpCall
func (self *dumpNode_dumpFilter) ProcessExpCall(_env *LnsEnv, node *Nodes_ExpCallNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var mess string
    mess = dumpNode_getTypeListTxt_4_(_env, node.FP.Get_expTypeList(_env))
    self.FP.dump(_env, opt, &node.Nodes_Node, mess)
    dumpNode_filter_3_(_env, node.FP.Get_func(_env), self, opt.FP.NextOpt(_env))
    {
        __exp := node.FP.Get_argList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 580: decl @lune.@base.@dumpNode.dumpFilter.processExpList
func (self *dumpNode_dumpFilter) ProcessExpList(_env *LnsEnv, node *Nodes_ExpListNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var mess string
    {
        _mRetExp := node.FP.Get_mRetExp(_env)
        if !Lns_IsNil( _mRetExp ) {
            mRetExp := _mRetExp.(*Nodes_MRetExp)
            mess = _env.LuaVM.String_format("hasMRetExp (%d): ", []LnsAny{mRetExp.FP.Get_index(_env)})
            
        } else {
            mess = "noMRetExp: "
            
        }
    }
    for _, _expType := range( node.FP.Get_expTypeList(_env).Items ) {
        expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        mess = _env.LuaVM.String_format("%s %s", []LnsAny{mess, expType.FP.GetTxt(_env, nil, nil, nil)})
        
    }
    self.FP.dump(_env, opt, &node.Nodes_Node, mess)
    var expList *LnsList
    expList = node.FP.Get_expList(_env)
    for _, _exp := range( expList.Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, exp, self, opt.FP.NextOpt(_env))
    }
}

// 600: decl @lune.@base.@dumpNode.dumpFilter.processExpMRet
func (self *dumpNode_dumpFilter) ProcessExpMRet(_env *LnsEnv, node *Nodes_ExpMRetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_mRet(_env), self, opt.FP.NextOpt(_env))
}

// 607: decl @lune.@base.@dumpNode.dumpFilter.processExpAccessMRet
func (self *dumpNode_dumpFilter) ProcessExpAccessMRet(_env *LnsEnv, node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%d", []LnsAny{node.FP.Get_index(_env)}))
}

// 613: decl @lune.@base.@dumpNode.dumpFilter.processExpOp1
func (self *dumpNode_dumpFilter) ProcessExpOp1(_env *LnsEnv, node *Nodes_ExpOp1Node,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_op(_env).Txt)
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}

// 620: decl @lune.@base.@dumpNode.dumpFilter.processExpToDDD
func (self *dumpNode_dumpFilter) ProcessExpToDDD(_env *LnsEnv, node *Nodes_ExpToDDDNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 627: decl @lune.@base.@dumpNode.dumpFilter.processExpMultiTo1
func (self *dumpNode_dumpFilter) ProcessExpMultiTo1(_env *LnsEnv, node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}

// 634: decl @lune.@base.@dumpNode.dumpFilter.processExpCast
func (self *dumpNode_dumpFilter) ProcessExpCast(_env *LnsEnv, node *Nodes_ExpCastNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s(%d) -> %s(%d)", []LnsAny{node.FP.Get_exp(_env).FP.Get_expType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil), node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_typeId(_env).Id, node.FP.Get_castType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil), node.FP.Get_castType(_env).FP.Get_typeId(_env).Id}))
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}

// 645: decl @lune.@base.@dumpNode.dumpFilter.processExpParen
func (self *dumpNode_dumpFilter) ProcessExpParen(_env *LnsEnv, node *Nodes_ExpParenNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "()")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}

// 651: decl @lune.@base.@dumpNode.dumpFilter.processExpSetVal
func (self *dumpNode_dumpFilter) ProcessExpSetVal(_env *LnsEnv, node *Nodes_ExpSetValNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("= %s", []LnsAny{node.FP.Get_expType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil)}))
    dumpNode_filter_3_(_env, node.FP.Get_exp1(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_exp2(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 658: decl @lune.@base.@dumpNode.dumpFilter.processExpSetItem
func (self *dumpNode_dumpFilter) ProcessExpSetItem(_env *LnsEnv, node *Nodes_ExpSetItemNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var indexSym string
    indexSym = ""
    var indexNode LnsAny
    indexNode = nil
    switch _matchExp1 := node.FP.Get_index(_env).(type) {
    case *Nodes_IndexVal__NodeIdx:
    index := _matchExp1.Val1
        indexNode = index
        
    case *Nodes_IndexVal__SymIdx:
    index := _matchExp1.Val1
        indexSym = index
        
    }
    self.FP.dump(_env, opt, &node.Nodes_Node, indexSym)
    dumpNode_filter_3_(_env, node.FP.Get_val(_env), self, opt.FP.NextOpt(_env))
    if indexNode != nil{
        indexNode_609 := indexNode.(*Nodes_Node)
        dumpNode_filter_3_(_env, indexNode_609, self, opt.FP.NextOpt(_env))
    }
    dumpNode_filter_3_(_env, node.FP.Get_exp2(_env), self, opt.FP.NextOpt(_env))
}

// 678: decl @lune.@base.@dumpNode.dumpFilter.processExpOp2
func (self *dumpNode_dumpFilter) ProcessExpOp2(_env *LnsEnv, node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s -> %s", []LnsAny{node.FP.Get_op(_env).Txt, node.FP.Get_expType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil)}))
    dumpNode_filter_3_(_env, node.FP.Get_exp1(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, node.FP.Get_exp2(_env), self, opt.FP.NextOpt(_env))
}

// 686: decl @lune.@base.@dumpNode.dumpFilter.processExpNew
func (self *dumpNode_dumpFilter) ProcessExpNew(_env *LnsEnv, node *Nodes_ExpNewNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_symbol(_env), self, opt.FP.NextOpt(_env))
    {
        __exp := node.FP.Get_argList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 696: decl @lune.@base.@dumpNode.dumpFilter.processExpRef
func (self *dumpNode_dumpFilter) ProcessExpRef(_env *LnsEnv, node *Nodes_ExpRefNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s:%s", []LnsAny{node.FP.Get_symbolInfo(_env).FP.Get_name(_env), node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
}

// 702: decl @lune.@base.@dumpNode.dumpFilter.processExpRefItem
func (self *dumpNode_dumpFilter) ProcessExpRefItem(_env *LnsEnv, node *Nodes_ExpRefItemNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "seq[exp] " + node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))
    dumpNode_filter_3_(_env, node.FP.Get_val(_env), self, opt.FP.NextOpt(_env))
    {
        __exp := node.FP.Get_index(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            dumpNode_filter_3_(_env, _exp, self, opt.FP.NextOpt(_env))
        }
    }
}

// 711: decl @lune.@base.@dumpNode.dumpFilter.processRefField
func (self *dumpNode_dumpFilter) ProcessRefField(_env *LnsEnv, node *Nodes_RefFieldNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s:%s:%s", []LnsAny{node.FP.Get_field(_env).Txt, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(node.FP.Get_symbolInfo(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_SymbolInfo).FP.Get_mutable(_env)}))) &&
        _env.SetStackVal( "mut") ||
        _env.SetStackVal( "imut") ).(string), node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
    dumpNode_filter_3_(_env, node.FP.Get_prefix(_env), self, opt.FP.NextOpt(_env))
}

// 720: decl @lune.@base.@dumpNode.dumpFilter.processExpOmitEnum
func (self *dumpNode_dumpFilter) ProcessExpOmitEnum(_env *LnsEnv, node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s.%s", []LnsAny{node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil), node.FP.Get_valToken(_env).Txt}))
}

// 727: decl @lune.@base.@dumpNode.dumpFilter.processGetField
func (self *dumpNode_dumpFilter) ProcessGetField(_env *LnsEnv, node *Nodes_GetFieldNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("get_%s:%s", []LnsAny{node.FP.Get_field(_env).Txt, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
    dumpNode_filter_3_(_env, node.FP.Get_prefix(_env), self, opt.FP.NextOpt(_env))
}

// 735: decl @lune.@base.@dumpNode.dumpFilter.processReturn
func (self *dumpNode_dumpFilter) ProcessReturn(_env *LnsEnv, node *Nodes_ReturnNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 744: decl @lune.@base.@dumpNode.dumpFilter.processProvide
func (self *dumpNode_dumpFilter) ProcessProvide(_env *LnsEnv, node *Nodes_ProvideNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_symbol(_env).FP.Get_name(_env))
}

// 750: decl @lune.@base.@dumpNode.dumpFilter.processAlias
func (self *dumpNode_dumpFilter) ProcessAlias(_env *LnsEnv, node *Nodes_AliasNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s = %s", []LnsAny{node.FP.Get_newSymbol(_env).FP.Get_name(_env), node.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
}

// 756: decl @lune.@base.@dumpNode.dumpFilter.processTestCase
func (self *dumpNode_dumpFilter) ProcessTestCase(_env *LnsEnv, node *Nodes_TestCaseNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_name(_env).Txt)
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}

// 763: decl @lune.@base.@dumpNode.dumpFilter.processTestBlock
func (self *dumpNode_dumpFilter) ProcessTestBlock(_env *LnsEnv, node *Nodes_TestBlockNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    for _, _statement := range( node.FP.Get_stmtList(_env).Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        dumpNode_filter_3_(_env, statement, self, opt.FP.NextOpt(_env))
    }
}

// 772: decl @lune.@base.@dumpNode.dumpFilter.processBoxing
func (self *dumpNode_dumpFilter) ProcessBoxing(_env *LnsEnv, node *Nodes_BoxingNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_src(_env), self, opt.FP.NextOpt(_env))
}

// 779: decl @lune.@base.@dumpNode.dumpFilter.processUnboxing
func (self *dumpNode_dumpFilter) ProcessUnboxing(_env *LnsEnv, node *Nodes_UnboxingNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_src(_env), self, opt.FP.NextOpt(_env))
}

// 786: decl @lune.@base.@dumpNode.dumpFilter.processLiteralList
func (self *dumpNode_dumpFilter) ProcessLiteralList(_env *LnsEnv, node *Nodes_LiteralListNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "[]")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 794: decl @lune.@base.@dumpNode.dumpFilter.processLiteralSet
func (self *dumpNode_dumpFilter) ProcessLiteralSet(_env *LnsEnv, node *Nodes_LiteralSetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "(@)")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 802: decl @lune.@base.@dumpNode.dumpFilter.processLiteralMap
func (self *dumpNode_dumpFilter) ProcessLiteralMap(_env *LnsEnv, node *Nodes_LiteralMapNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "{}")
    var pairList *LnsList
    pairList = node.FP.Get_pairList(_env)
    for _, _pair := range( pairList.Items ) {
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        dumpNode_filter_3_(_env, pair.FP.Get_key(_env), self, opt.FP.NextOpt(_env))
        dumpNode_filter_3_(_env, pair.FP.Get_val(_env), self, opt.FP.NextOpt(_env))
    }
}

// 813: decl @lune.@base.@dumpNode.dumpFilter.processLiteralArray
func (self *dumpNode_dumpFilter) ProcessLiteralArray(_env *LnsEnv, node *Nodes_LiteralArrayNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "[@]")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}

// 821: decl @lune.@base.@dumpNode.dumpFilter.processLiteralChar
func (self *dumpNode_dumpFilter) ProcessLiteralChar(_env *LnsEnv, node *Nodes_LiteralCharNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s(%s)", []LnsAny{node.FP.Get_num(_env), node.FP.Get_token(_env).Txt}))
}

// 827: decl @lune.@base.@dumpNode.dumpFilter.processLiteralInt
func (self *dumpNode_dumpFilter) ProcessLiteralInt(_env *LnsEnv, node *Nodes_LiteralIntNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s(%s)", []LnsAny{node.FP.Get_num(_env), node.FP.Get_token(_env).Txt}))
}

// 833: decl @lune.@base.@dumpNode.dumpFilter.processLiteralReal
func (self *dumpNode_dumpFilter) ProcessLiteralReal(_env *LnsEnv, node *Nodes_LiteralRealNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.LuaVM.String_format("%s(%s)", []LnsAny{node.FP.Get_num(_env), node.FP.Get_token(_env).Txt}))
}

// 839: decl @lune.@base.@dumpNode.dumpFilter.processLiteralString
func (self *dumpNode_dumpFilter) ProcessLiteralString(_env *LnsEnv, node *Nodes_LiteralStringNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_token(_env).Txt)
    {
        _expList := node.FP.Get_orgParam(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            for _, _param := range( expList.FP.Get_expList(_env).Items ) {
                param := _param.(Nodes_NodeDownCast).ToNodes_Node()
                dumpNode_filter_3_(_env, param, self, opt.FP.NextOpt(_env))
            }
        }
    }
}

// 849: decl @lune.@base.@dumpNode.dumpFilter.processLiteralBool
func (self *dumpNode_dumpFilter) ProcessLiteralBool(_env *LnsEnv, node *Nodes_LiteralBoolNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_token(_env).Txt == "true") &&
        _env.SetStackVal( "true") ||
        _env.SetStackVal( "false") ).(string))
}

// 855: decl @lune.@base.@dumpNode.dumpFilter.processLiteralNil
func (self *dumpNode_dumpFilter) ProcessLiteralNil(_env *LnsEnv, node *Nodes_LiteralNilNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}

// 860: decl @lune.@base.@dumpNode.dumpFilter.processBreak
func (self *dumpNode_dumpFilter) ProcessBreak(_env *LnsEnv, node *Nodes_BreakNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}

// 865: decl @lune.@base.@dumpNode.dumpFilter.processLiteralSymbol
func (self *dumpNode_dumpFilter) ProcessLiteralSymbol(_env *LnsEnv, node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_token(_env).Txt)
}

// 871: decl @lune.@base.@dumpNode.dumpFilter.processAbbr
func (self *dumpNode_dumpFilter) ProcessAbbr(_env *LnsEnv, node *Nodes_AbbrNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "##")
}


func Lns_dumpNode_init(_env *LnsEnv) {
    if init_dumpNode { return }
    init_dumpNode = true
    dumpNode__mod__ = "@lune.@base.@dumpNode"
    Lns_InitMod()
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Util_init(_env)
    Lns_convLua_init(_env)
    Lns_AstInfo_init(_env)
}
func init() {
    init_dumpNode = false
}
