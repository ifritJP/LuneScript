// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_dumpNode bool
var dumpNode__mod__ string
// for 63
func dumpNode_convExp1_59(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 510
func dumpNode_convExp3_532(arg1 []LnsAny) (string, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// 106: decl @lune.@base.@dumpNode.createFilter
func DumpNode_createFilter(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,processInfo *Ast_ProcessInfo,stream Lns_oStream) *Nodes_Filter {
    return &NewdumpNode_dumpFilter(_env, true, moduleTypeInfo, moduleTypeInfo.FP.Get_scope(_env), stream, processInfo).Nodes_Filter
}

// 113: decl @lune.@base.@dumpNode.filter
func dumpNode_filter_3_(_env *LnsEnv, node *Nodes_Node,filter *dumpNode_dumpFilter,opt *DumpNode_Opt) {
    node.FP.ProcessFilter(_env, &filter.Nodes_Filter, DumpNode_Opt2Stem(opt))
}

// 593: decl @lune.@base.@dumpNode.getTypeListTxt
func dumpNode_getTypeListTxt_4_(_env *LnsEnv, typeList *LnsList2_[*Ast_TypeInfo]) string {
    var txt string
    txt = ""
    for _index, _typeInfo := range( typeList.Items ) {
        index := _index + 1
        typeInfo := _typeInfo
        if index > 1{
            txt = txt + ", "
        }
        txt = txt + typeInfo.FP.GetTxt(_env, nil, nil, nil)
    }
    return txt
}


// 45: decl @lune.@base.@dumpNode.Opt.get
func (self *DumpNode_Opt) Get(_env *LnsEnv)(string, LnsInt) {
    return self.prefix, self.depth
}
// 48: decl @lune.@base.@dumpNode.Opt.nextOpt
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
// 62: decl @lune.@base.@dumpNode.dumpFilter.writeln
func (self *dumpNode_dumpFilter) writeln(_env *LnsEnv, opt *DumpNode_Opt,txt string) {
    var prefix string
    prefix = dumpNode_convExp1_59(Lns_2DDD(opt.FP.Get(_env)))
    self.stream.Write(_env, _env.GetVM().String_format("%s%s\n", Lns_2DDD(prefix, Lns_car(_env.GetVM().String_gsub(txt," *$", "")).(string))))
}
// 66: decl @lune.@base.@dumpNode.dumpFilter.dump
func (self *dumpNode_dumpFilter) dump(_env *LnsEnv, opt *DumpNode_Opt,node *Nodes_Node,txt string) {
    var typeStr string
    typeStr = ""
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType(_env)
    if expType.FP.Equals(_env, self.processInfo, Ast_builtinTypeNone, nil, nil){
        typeStr = _env.GetVM().String_format("(%d:%s:%s)", Lns_2DDD(expType.FP.Get_typeId(_env).Id, expType.FP.GetTxt(_env, nil, nil, nil), expType.FP.Get_kind(_env)))
    }
    var comment string
    {
        _commentList := node.FP.Get_commentList(_env)
        if !Lns_IsNil( _commentList ) {
            commentList := _commentList.(*LnsList2_[*Types_Token])
            comment = _env.GetVM().String_format("comment:%d,%d", Lns_2DDD(commentList.Len(), _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( node.FP.Get_tailComment(_env)) &&
                _env.SetStackVal( 1) ||
                _env.SetStackVal( 0) ).(LnsInt)))
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
        attrib = _env.GetVM().String_format("%s %s", Lns_2DDD(attrib, "nilacc"))
    }
    if node.FP.IsIntermediate(_env){
        attrib = _env.GetVM().String_format("%s %s", Lns_2DDD(attrib, "inter"))
    }
    if len(attrib) != 0{
        attrib = _env.GetVM().String_format("[%s]", Lns_2DDD(attrib))
    }
    self.FP.writeln(_env, opt, _env.GetVM().String_format(": %s:%s(%d:%d) %s %s %s %s", Lns_2DDD(Nodes_getNodeKindName(_env, node.FP.Get_kind(_env)), node.FP.GetIdTxt(_env), node.FP.Get_effectivePos(_env).LineNo, node.FP.Get_effectivePos(_env).Column, attrib, txt, typeStr, comment)))
}
// 117: decl @lune.@base.@dumpNode.dumpFilter.processNone
func (self *dumpNode_dumpFilter) ProcessNone(_env *LnsEnv, node *Nodes_NoneNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}
// 122: decl @lune.@base.@dumpNode.dumpFilter.processShebang
func (self *dumpNode_dumpFilter) ProcessShebang(_env *LnsEnv, node *Nodes_ShebangNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_cmd(_env))
}
// 127: decl @lune.@base.@dumpNode.dumpFilter.processLuneControl
func (self *dumpNode_dumpFilter) ProcessLuneControl(_env *LnsEnv, node *Nodes_LuneControlNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_pragma(_env).(LnsAlgeVal).GetTxt())
}
// 132: decl @lune.@base.@dumpNode.dumpFilter.processBlankLine
func (self *dumpNode_dumpFilter) ProcessBlankLine(_env *LnsEnv, node *Nodes_BlankLineNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%d", Lns_2DDD(node.FP.Get_lineNum(_env))))
}
// 137: decl @lune.@base.@dumpNode.dumpFilter.processLuneKind
func (self *dumpNode_dumpFilter) ProcessLuneKind(_env *LnsEnv, node *Nodes_LuneKindNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 143: decl @lune.@base.@dumpNode.dumpFilter.processImport
func (self *dumpNode_dumpFilter) ProcessImport(_env *LnsEnv, node *Nodes_ImportNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_info(_env).FP.Get_modulePath(_env))
}
// 148: decl @lune.@base.@dumpNode.dumpFilter.processRoot
func (self *dumpNode_dumpFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    for _, _child := range( node.FP.Get_children(_env).Items ) {
        child := _child
        dumpNode_filter_3_(_env, child, self, opt.FP.NextOpt(_env))
    }
}
// 156: decl @lune.@base.@dumpNode.dumpFilter.processSubfile
func (self *dumpNode_dumpFilter) ProcessSubfile(_env *LnsEnv, node *Nodes_SubfileNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}
// 162: decl @lune.@base.@dumpNode.dumpFilter.processRequest
func (self *dumpNode_dumpFilter) ProcessRequest(_env *LnsEnv, node *Nodes_RequestNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_processor(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 170: decl @lune.@base.@dumpNode.dumpFilter.processAsyncLock
func (self *dumpNode_dumpFilter) ProcessAsyncLock(_env *LnsEnv, node *Nodes_AsyncLockNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s, hasReturn = %s", Lns_2DDD(Nodes_LockKind_getTxt( node.FP.Get_lockKind(_env)), node.FP.Get_returnTypeList(_env) != nil)))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}
// 178: decl @lune.@base.@dumpNode.dumpFilter.processBlockSub
func (self *dumpNode_dumpFilter) ProcessBlockSub(_env *LnsEnv, node *Nodes_BlockNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s, %s, hasAsyncLockBreak:%s", Lns_2DDD(Nodes_BlockKind_getTxt( node.FP.Get_blockKind(_env)), Nodes_BreakKind_getTxt( node.FP.GetBreakKind(_env, Nodes_CheckBreakMode__Return)), node.FP.Get_hasAsyncLockBreak(_env))))
    for _, _statement := range( node.FP.Get_stmtList(_env).Items ) {
        statement := _statement
        dumpNode_filter_3_(_env, statement, self, opt.FP.NextOpt(_env))
    }
}
// 189: decl @lune.@base.@dumpNode.dumpFilter.processScope
func (self *dumpNode_dumpFilter) ProcessScope(_env *LnsEnv, node *Nodes_ScopeNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, Nodes_ScopeKind_getTxt( node.FP.Get_scopeKind(_env)))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}
// 196: decl @lune.@base.@dumpNode.dumpFilter.processStmtExp
func (self *dumpNode_dumpFilter) ProcessStmtExp(_env *LnsEnv, node *Nodes_StmtExpNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 203: decl @lune.@base.@dumpNode.dumpFilter.processDeclEnum
func (self *dumpNode_dumpFilter) ProcessDeclEnum(_env *LnsEnv, node *Nodes_DeclEnumNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_name(_env).Txt)
    var enumTypeInfo *Ast_EnumTypeInfo
    enumTypeInfo = Lns_unwrap( (Ast_EnumTypeInfoDownCastF(node.FP.Get_expType(_env).FP))).(*Ast_EnumTypeInfo)
    for _, _name := range( node.FP.Get_valueNameList(_env).Items ) {
        name := _name
        var valInfo *Ast_EnumValInfo
        valInfo = Lns_unwrap( enumTypeInfo.FP.GetEnumValInfo(_env, name.Txt)).(*Ast_EnumValInfo)
        var val LnsAny
        switch _matchExp0 := valInfo.FP.Get_val(_env).(type) {
        case *Ast_EnumLiteral__Int:
            x := _matchExp0.Val1
            val = x
        case *Ast_EnumLiteral__Real:
            x := _matchExp0.Val1
            val = x
        case *Ast_EnumLiteral__Str:
            x := _matchExp0.Val1
            val = x
        }
        self.FP.writeln(_env, opt, _env.GetVM().String_format("  %s: %s, %s", Lns_2DDD(name.Txt, valInfo.FP.Get_val(_env).(LnsAlgeVal).GetTxt(), val)))
    }
}
// 226: decl @lune.@base.@dumpNode.dumpFilter.processDeclAlge
func (self *dumpNode_dumpFilter) ProcessDeclAlge(_env *LnsEnv, node *Nodes_DeclAlgeNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var algeTypeInfo *Ast_AlgeTypeInfo
    algeTypeInfo = node.FP.Get_algeType(_env)
    self.FP.dump(_env, opt, &node.Nodes_Node, algeTypeInfo.FP.Get_rawTxt(_env))
    {
        __forsortCollection0 := algeTypeInfo.FP.Get_valInfoMap(_env)
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            valInfo := __forsortCollection0.Items[ ___forsortKey0 ]
            self.FP.writeln(_env, opt, _env.GetVM().String_format("  %s: %s", Lns_2DDD(algeTypeInfo.FP.Get_rawTxt(_env), valInfo.FP.Get_name(_env))))
        }
    }
}
// 236: decl @lune.@base.@dumpNode.dumpFilter.processNewAlgeVal
func (self *dumpNode_dumpFilter) ProcessNewAlgeVal(_env *LnsEnv, node *Nodes_NewAlgeValNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s: %s", Lns_2DDD(node.FP.Get_name(_env).Txt, node.FP.Get_expType(_env).FP.Get_display_stirng(_env))))
    for _, _exp := range( node.FP.Get_paramList(_env).Items ) {
        exp := _exp
        dumpNode_filter_3_(_env, exp, self, opt.FP.NextOpt(_env))
    }
}
// 245: decl @lune.@base.@dumpNode.dumpFilter.processProtoClass
func (self *dumpNode_dumpFilter) ProcessProtoClass(_env *LnsEnv, node *Nodes_ProtoClassNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_name(_env).Txt)
}
// 251: decl @lune.@base.@dumpNode.dumpFilter.processDeclClass
func (self *dumpNode_dumpFilter) ProcessDeclClass(_env *LnsEnv, node *Nodes_DeclClassNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s (%s)", Lns_2DDD(node.FP.Get_name(_env).Txt, self.FP.GetFull(_env, node.FP.Get_expType(_env), false))))
    {
        _baseNode := node.FP.Get_inheritInfo(_env).FP.Get_base(_env)
        if !Lns_IsNil( _baseNode ) {
            baseNode := _baseNode.(*Nodes_RefTypeNode)
            dumpNode_filter_3_(_env, &baseNode.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
    for _, _ifNode := range( node.FP.Get_inheritInfo(_env).FP.Get_impliments(_env).Items ) {
        ifNode := _ifNode
        dumpNode_filter_3_(_env, &ifNode.Nodes_Node, self, opt.FP.NextOpt(_env))
    }
    for _, _field := range( node.FP.Get_fieldList(_env).Items ) {
        field := _field
        dumpNode_filter_3_(_env, field, self, opt.FP.NextOpt(_env))
    }
}
// 267: decl @lune.@base.@dumpNode.dumpFilter.processDeclMember
func (self *dumpNode_dumpFilter) ProcessDeclMember(_env *LnsEnv, node *Nodes_DeclMemberNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_name(_env).Txt)
    dumpNode_filter_3_(_env, &node.FP.Get_refType(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}
// 273: decl @lune.@base.@dumpNode.dumpFilter.processExpMacroArgExp
func (self *dumpNode_dumpFilter) ProcessExpMacroArgExp(_env *LnsEnv, node *Nodes_ExpMacroArgExpNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s", Lns_2DDD(node.FP.Get_codeTxt(_env))))
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 279: decl @lune.@base.@dumpNode.dumpFilter.processExpMacroExp
func (self *dumpNode_dumpFilter) ProcessExpMacroExp(_env *LnsEnv, node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s", Lns_2DDD(self.FP.GetFull(_env, node.FP.Get_macroType(_env), false))))
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &expList.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
    var stmtList *LnsList2_[*Nodes_Node]
    stmtList = node.FP.Get_stmtList(_env)
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt
        dumpNode_filter_3_(_env, stmt, self, opt.FP.NextOpt(_env))
    }
}
// 291: decl @lune.@base.@dumpNode.dumpFilter.processDeclMacro
func (self *dumpNode_dumpFilter) ProcessDeclMacro(_env *LnsEnv, node *Nodes_DeclMacroNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))
}
// 296: decl @lune.@base.@dumpNode.dumpFilter.processExpMacroStat
func (self *dumpNode_dumpFilter) ProcessExpMacroStat(_env *LnsEnv, node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))
    for _, _expStr := range( node.FP.Get_expStrList(_env).Items ) {
        expStr := _expStr
        dumpNode_filter_3_(_env, expStr, self, opt.FP.NextOpt(_env))
    }
}
// 306: decl @lune.@base.@dumpNode.dumpFilter.processUnwrapSet
func (self *dumpNode_dumpFilter) ProcessUnwrapSet(_env *LnsEnv, node *Nodes_UnwrapSetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, &node.FP.Get_dstExpList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_srcExpList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    if Lns_isCondTrue( node.FP.Get_unwrapBlock(_env)){
        dumpNode_filter_3_(_env, &Lns_unwrap( node.FP.Get_unwrapBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt(_env))
    }
}
// 318: decl @lune.@base.@dumpNode.dumpFilter.processIfUnwrap
func (self *dumpNode_dumpFilter) ProcessIfUnwrap(_env *LnsEnv, node *Nodes_IfUnwrapNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    if Lns_isCondTrue( node.FP.Get_nilBlock(_env)){
        dumpNode_filter_3_(_env, &Lns_unwrap( node.FP.Get_nilBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt(_env))
    }
}
// 332: decl @lune.@base.@dumpNode.dumpFilter.processWhen
func (self *dumpNode_dumpFilter) ProcessWhen(_env *LnsEnv, node *Nodes_WhenNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var symTxt string
    symTxt = ""
    for _, _symPair := range( node.FP.Get_symPairList(_env).Items ) {
        symPair := _symPair
        symTxt = _env.GetVM().String_format("%s %s", Lns_2DDD(symTxt, symPair.FP.Get_src(_env).FP.Get_name(_env)))
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
// 347: decl @lune.@base.@dumpNode.dumpFilter.processExpExpandTuple
func (self *dumpNode_dumpFilter) ProcessExpExpandTuple(_env *LnsEnv, node *Nodes_ExpExpandTupleNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 355: decl @lune.@base.@dumpNode.dumpFilter.processDeclVar
func (self *dumpNode_dumpFilter) ProcessDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var varName string
    varName = ""
    for _index, __var := range( node.FP.Get_varList(_env).Items ) {
        index := _index + 1
        _var := __var
        if index > 1{
            varName = varName + ","
        }
        varName = _env.GetVM().String_format("%s %s", Lns_2DDD(varName, _var.FP.Get_name(_env).Txt))
    }
    if Lns_isCondTrue( node.FP.Get_unwrapBlock(_env)){
        varName = "!" + varName
    }
    varName = _env.GetVM().String_format("%s %s", Lns_2DDD(node.FP.Get_mode(_env), varName))
    self.FP.dump(_env, opt, &node.Nodes_Node, varName)
    for _, __var := range( node.FP.Get_varList(_env).Items ) {
        _var := __var
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
}
// 391: decl @lune.@base.@dumpNode.dumpFilter.processDeclArg
func (self *dumpNode_dumpFilter) ProcessDeclArg(_env *LnsEnv, node *Nodes_DeclArgNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s:%s", Lns_2DDD(node.FP.Get_name(_env).Txt, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))))
}
// 397: decl @lune.@base.@dumpNode.dumpFilter.processDeclArgDDD
func (self *dumpNode_dumpFilter) ProcessDeclArgDDD(_env *LnsEnv, node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "...")
}
// 402: decl @lune.@base.@dumpNode.dumpFilter.processExpSubDDD
func (self *dumpNode_dumpFilter) ProcessExpSubDDD(_env *LnsEnv, node *Nodes_ExpSubDDDNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("... (%d)", Lns_2DDD(node.FP.Get_remainIndex(_env))))
}
// 409: decl @lune.@base.@dumpNode.dumpFilter.processDeclForm
func (self *dumpNode_dumpFilter) ProcessDeclForm(_env *LnsEnv, node *Nodes_DeclFormNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil))
}
// 414: decl @lune.@base.@dumpNode.dumpFilter.processDeclFuncInfo
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
    self.FP.dump(_env, opt, node, _env.GetVM().String_format("%s -- stmt:%d", Lns_2DDD(name, declInfo.FP.Get_stmtNum(_env))))
    var argList *LnsList2_[*Nodes_Node]
    argList = declInfo.FP.Get_argList(_env)
    for _, _arg := range( argList.Items ) {
        arg := _arg
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
// 436: decl @lune.@base.@dumpNode.dumpFilter.processDeclFunc
func (self *dumpNode_dumpFilter) ProcessDeclFunc(_env *LnsEnv, node *Nodes_DeclFuncNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}
// 441: decl @lune.@base.@dumpNode.dumpFilter.processDeclMethod
func (self *dumpNode_dumpFilter) ProcessDeclMethod(_env *LnsEnv, node *Nodes_DeclMethodNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}
// 446: decl @lune.@base.@dumpNode.dumpFilter.processProtoMethod
func (self *dumpNode_dumpFilter) ProcessProtoMethod(_env *LnsEnv, node *Nodes_ProtoMethodNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}
// 451: decl @lune.@base.@dumpNode.dumpFilter.processDeclConstr
func (self *dumpNode_dumpFilter) ProcessDeclConstr(_env *LnsEnv, node *Nodes_DeclConstrNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}
// 457: decl @lune.@base.@dumpNode.dumpFilter.processDeclDestr
func (self *dumpNode_dumpFilter) ProcessDeclDestr(_env *LnsEnv, node *Nodes_DeclDestrNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}
// 463: decl @lune.@base.@dumpNode.dumpFilter.processExpCallSuperCtor
func (self *dumpNode_dumpFilter) ProcessExpCallSuperCtor(_env *LnsEnv, node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType(_env)
    self.FP.dump(_env, opt, &node.Nodes_Node, typeInfo.FP.GetTxt(_env, nil, nil, nil))
}
// 469: decl @lune.@base.@dumpNode.dumpFilter.processExpCallSuper
func (self *dumpNode_dumpFilter) ProcessExpCallSuper(_env *LnsEnv, node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType(_env)
    self.FP.dump(_env, opt, &node.Nodes_Node, typeInfo.FP.GetTxt(_env, nil, nil, nil))
}
// 475: decl @lune.@base.@dumpNode.dumpFilter.processRefType
func (self *dumpNode_dumpFilter) ProcessRefType(_env *LnsEnv, node *Nodes_RefTypeNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.Get_display_stirng(_env))
    dumpNode_filter_3_(_env, node.FP.Get_typeNode(_env), self, opt.FP.NextOpt(_env))
}
// 481: decl @lune.@base.@dumpNode.dumpFilter.processIf
func (self *dumpNode_dumpFilter) ProcessIf(_env *LnsEnv, node *Nodes_IfNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    var stmtList *LnsList2_[*Nodes_IfStmtInfo]
    stmtList = node.FP.Get_stmtList(_env)
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt
        if stmt.FP.Get_exp(_env).FP.Get_kind(_env) != Nodes_nodeKindEnum__None{
            dumpNode_filter_3_(_env, stmt.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
        }
        dumpNode_filter_3_(_env, &stmt.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    }
}
// 493: decl @lune.@base.@dumpNode.dumpFilter.processSwitch
func (self *dumpNode_dumpFilter) ProcessSwitch(_env *LnsEnv, node *Nodes_SwitchNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
    var caseList *LnsList2_[*Nodes_CaseInfo]
    caseList = node.FP.Get_caseList(_env)
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo
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
// 507: decl @lune.@base.@dumpNode.dumpFilter.processMatch
func (self *dumpNode_dumpFilter) ProcessMatch(_env *LnsEnv, node *Nodes_MatchNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var prefix string
    var depth LnsInt
    prefix,depth = opt.FP.Get(_env)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_val(_env), self, opt.FP.NextOpt(_env))
    var caseList *LnsList2_[*Nodes_MatchCase]
    caseList = node.FP.Get_caseList(_env)
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo
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
// 523: decl @lune.@base.@dumpNode.dumpFilter.processWhile
func (self *dumpNode_dumpFilter) ProcessWhile(_env *LnsEnv, node *Nodes_WhileNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}
// 530: decl @lune.@base.@dumpNode.dumpFilter.processRepeat
func (self *dumpNode_dumpFilter) ProcessRepeat(_env *LnsEnv, node *Nodes_RepeatNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 537: decl @lune.@base.@dumpNode.dumpFilter.processFor
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
// 548: decl @lune.@base.@dumpNode.dumpFilter.processApply
func (self *dumpNode_dumpFilter) ProcessApply(_env *LnsEnv, node *Nodes_ApplyNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var varNames string
    varNames = ""
    var varList *LnsList2_[*Ast_SymbolInfo]
    varList = node.FP.Get_varList(_env)
    for _, __var := range( varList.Items ) {
        _var := __var
        varNames = varNames + _var.FP.Get_name(_env) + " "
    }
    self.FP.dump(_env, opt, &node.Nodes_Node, varNames)
    dumpNode_filter_3_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}
// 560: decl @lune.@base.@dumpNode.dumpFilter.processForeach
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
// 571: decl @lune.@base.@dumpNode.dumpFilter.processForsort
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
// 583: decl @lune.@base.@dumpNode.dumpFilter.processExpUnwrap
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
// 604: decl @lune.@base.@dumpNode.dumpFilter.processExpCall
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
// 615: decl @lune.@base.@dumpNode.dumpFilter.processExpList
func (self *dumpNode_dumpFilter) ProcessExpList(_env *LnsEnv, node *Nodes_ExpListNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var mess string
    {
        _mRetExp := node.FP.Get_mRetExp(_env)
        if !Lns_IsNil( _mRetExp ) {
            mRetExp := _mRetExp.(*Nodes_MRetExp)
            mess = _env.GetVM().String_format("hasMRetExp (%d): ", Lns_2DDD(mRetExp.FP.Get_index(_env)))
        } else {
            mess = "noMRetExp: "
        }
    }
    for _, _expType := range( node.FP.Get_expTypeList(_env).Items ) {
        expType := _expType
        mess = _env.GetVM().String_format("%s %s", Lns_2DDD(mess, expType.FP.GetTxt(_env, nil, nil, nil)))
    }
    self.FP.dump(_env, opt, &node.Nodes_Node, mess)
    var expList *LnsList2_[*Nodes_Node]
    expList = node.FP.Get_expList(_env)
    for _, _exp := range( expList.Items ) {
        exp := _exp
        dumpNode_filter_3_(_env, exp, self, opt.FP.NextOpt(_env))
    }
}
// 635: decl @lune.@base.@dumpNode.dumpFilter.processExpMRet
func (self *dumpNode_dumpFilter) ProcessExpMRet(_env *LnsEnv, node *Nodes_ExpMRetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_mRet(_env), self, opt.FP.NextOpt(_env))
}
// 642: decl @lune.@base.@dumpNode.dumpFilter.processExpAccessMRet
func (self *dumpNode_dumpFilter) ProcessExpAccessMRet(_env *LnsEnv, node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%d", Lns_2DDD(node.FP.Get_index(_env))))
}
// 648: decl @lune.@base.@dumpNode.dumpFilter.processExpOp1
func (self *dumpNode_dumpFilter) ProcessExpOp1(_env *LnsEnv, node *Nodes_ExpOp1Node,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_op(_env).Txt)
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 655: decl @lune.@base.@dumpNode.dumpFilter.processExpToDDD
func (self *dumpNode_dumpFilter) ProcessExpToDDD(_env *LnsEnv, node *Nodes_ExpToDDDNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}
// 662: decl @lune.@base.@dumpNode.dumpFilter.processExpMultiTo1
func (self *dumpNode_dumpFilter) ProcessExpMultiTo1(_env *LnsEnv, node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 669: decl @lune.@base.@dumpNode.dumpFilter.processExpCast
func (self *dumpNode_dumpFilter) ProcessExpCast(_env *LnsEnv, node *Nodes_ExpCastNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s: %s(%d) -> %s(%d)", Lns_2DDD(Nodes_CastKind_getTxt( node.FP.Get_castKind(_env)), node.FP.Get_exp(_env).FP.Get_expType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil), node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_typeId(_env).Id, node.FP.Get_castType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil), node.FP.Get_castType(_env).FP.Get_typeId(_env).Id)))
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 681: decl @lune.@base.@dumpNode.dumpFilter.processExpParen
func (self *dumpNode_dumpFilter) ProcessExpParen(_env *LnsEnv, node *Nodes_ExpParenNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "()")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 687: decl @lune.@base.@dumpNode.dumpFilter.processExpSetVal
func (self *dumpNode_dumpFilter) ProcessExpSetVal(_env *LnsEnv, node *Nodes_ExpSetValNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("= %s", Lns_2DDD(node.FP.Get_expType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil))))
    dumpNode_filter_3_(_env, node.FP.Get_exp1(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, &node.FP.Get_exp2(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}
// 694: decl @lune.@base.@dumpNode.dumpFilter.processExpSetItem
func (self *dumpNode_dumpFilter) ProcessExpSetItem(_env *LnsEnv, node *Nodes_ExpSetItemNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    var indexSym string
    indexSym = ""
    var indexNode LnsAny
    indexNode = nil
    switch _matchExp0 := node.FP.Get_index(_env).(type) {
    case *Nodes_IndexVal__NodeIdx:
        index := _matchExp0.Val1
        indexNode = index
    case *Nodes_IndexVal__SymIdx:
        index := _matchExp0.Val1
        indexSym = index
    }
    self.FP.dump(_env, opt, &node.Nodes_Node, indexSym)
    dumpNode_filter_3_(_env, node.FP.Get_val(_env), self, opt.FP.NextOpt(_env))
    if indexNode != nil{
        indexNode_67 := indexNode.(*Nodes_Node)
        dumpNode_filter_3_(_env, indexNode_67, self, opt.FP.NextOpt(_env))
    }
    dumpNode_filter_3_(_env, node.FP.Get_exp2(_env), self, opt.FP.NextOpt(_env))
}
// 714: decl @lune.@base.@dumpNode.dumpFilter.processExpOp2
func (self *dumpNode_dumpFilter) ProcessExpOp2(_env *LnsEnv, node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s -> %s", Lns_2DDD(node.FP.Get_op(_env).Txt, node.FP.Get_expType(_env).FP.GetTxt(_env, self.FP.Get_typeNameCtrl(_env), nil, nil))))
    dumpNode_filter_3_(_env, node.FP.Get_exp1(_env), self, opt.FP.NextOpt(_env))
    dumpNode_filter_3_(_env, node.FP.Get_exp2(_env), self, opt.FP.NextOpt(_env))
}
// 722: decl @lune.@base.@dumpNode.dumpFilter.processExpNew
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
// 732: decl @lune.@base.@dumpNode.dumpFilter.processExpRef
func (self *dumpNode_dumpFilter) ProcessExpRef(_env *LnsEnv, node *Nodes_ExpRefNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s:%s", Lns_2DDD(node.FP.Get_symbolInfo(_env).FP.Get_name(_env), node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))))
}
// 738: decl @lune.@base.@dumpNode.dumpFilter.processExpRefItem
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
// 747: decl @lune.@base.@dumpNode.dumpFilter.processRefField
func (self *dumpNode_dumpFilter) ProcessRefField(_env *LnsEnv, node *Nodes_RefFieldNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s:%s:%s", Lns_2DDD(node.FP.Get_field(_env).Txt, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(node.FP.Get_symbolInfo(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_SymbolInfo).FP.Get_mutable(_env)}))) &&
        _env.SetStackVal( "mut") ||
        _env.SetStackVal( "imut") ).(string), node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))))
    dumpNode_filter_3_(_env, node.FP.Get_prefix(_env), self, opt.FP.NextOpt(_env))
}
// 756: decl @lune.@base.@dumpNode.dumpFilter.processExpOmitEnum
func (self *dumpNode_dumpFilter) ProcessExpOmitEnum(_env *LnsEnv, node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s.%s", Lns_2DDD(node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil), node.FP.Get_valToken(_env).Txt)))
}
// 763: decl @lune.@base.@dumpNode.dumpFilter.processGetField
func (self *dumpNode_dumpFilter) ProcessGetField(_env *LnsEnv, node *Nodes_GetFieldNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("get_%s:%s", Lns_2DDD(node.FP.Get_field(_env).Txt, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))))
    dumpNode_filter_3_(_env, node.FP.Get_prefix(_env), self, opt.FP.NextOpt(_env))
}
// 771: decl @lune.@base.@dumpNode.dumpFilter.processReturn
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
// 780: decl @lune.@base.@dumpNode.dumpFilter.processCondRet
func (self *dumpNode_dumpFilter) ProcessCondRet(_env *LnsEnv, node *Nodes_CondRetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 787: decl @lune.@base.@dumpNode.dumpFilter.processCondRetList
func (self *dumpNode_dumpFilter) ProcessCondRetList(_env *LnsEnv, node *Nodes_CondRetListNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env))
}
// 795: decl @lune.@base.@dumpNode.dumpFilter.processProvide
func (self *dumpNode_dumpFilter) ProcessProvide(_env *LnsEnv, node *Nodes_ProvideNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_symbol(_env).FP.Get_name(_env))
}
// 801: decl @lune.@base.@dumpNode.dumpFilter.processAlias
func (self *dumpNode_dumpFilter) ProcessAlias(_env *LnsEnv, node *Nodes_AliasNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s = %s", Lns_2DDD(node.FP.Get_newSymbol(_env).FP.Get_name(_env), node.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil))))
}
// 807: decl @lune.@base.@dumpNode.dumpFilter.processTestCase
func (self *dumpNode_dumpFilter) ProcessTestCase(_env *LnsEnv, node *Nodes_TestCaseNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_name(_env).Txt)
    dumpNode_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}
// 814: decl @lune.@base.@dumpNode.dumpFilter.processTestBlock
func (self *dumpNode_dumpFilter) ProcessTestBlock(_env *LnsEnv, node *Nodes_TestBlockNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    for _, _statement := range( node.FP.Get_stmtList(_env).Items ) {
        statement := _statement
        dumpNode_filter_3_(_env, statement, self, opt.FP.NextOpt(_env))
    }
}
// 823: decl @lune.@base.@dumpNode.dumpFilter.processBoxing
func (self *dumpNode_dumpFilter) ProcessBoxing(_env *LnsEnv, node *Nodes_BoxingNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_src(_env), self, opt.FP.NextOpt(_env))
}
// 830: decl @lune.@base.@dumpNode.dumpFilter.processUnboxing
func (self *dumpNode_dumpFilter) ProcessUnboxing(_env *LnsEnv, node *Nodes_UnboxingNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
    dumpNode_filter_3_(_env, node.FP.Get_src(_env), self, opt.FP.NextOpt(_env))
}
// 837: decl @lune.@base.@dumpNode.dumpFilter.processTupleConst
func (self *dumpNode_dumpFilter) ProcessTupleConst(_env *LnsEnv, node *Nodes_TupleConstNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "(=)")
    dumpNode_filter_3_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env))
}
// 844: decl @lune.@base.@dumpNode.dumpFilter.processLiteralList
func (self *dumpNode_dumpFilter) ProcessLiteralList(_env *LnsEnv, node *Nodes_LiteralListNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}
// 852: decl @lune.@base.@dumpNode.dumpFilter.processLiteralSet
func (self *dumpNode_dumpFilter) ProcessLiteralSet(_env *LnsEnv, node *Nodes_LiteralSetNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            dumpNode_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env))
        }
    }
}
// 860: decl @lune.@base.@dumpNode.dumpFilter.processLiteralMap
func (self *dumpNode_dumpFilter) ProcessLiteralMap(_env *LnsEnv, node *Nodes_LiteralMapNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))
    var pairList *LnsList2_[*Nodes_PairItem]
    pairList = node.FP.Get_pairList(_env)
    for _, _pair := range( pairList.Items ) {
        pair := _pair
        dumpNode_filter_3_(_env, pair.FP.Get_key(_env), self, opt.FP.NextOpt(_env))
        dumpNode_filter_3_(_env, pair.FP.Get_val(_env), self, opt.FP.NextOpt(_env))
    }
}
// 871: decl @lune.@base.@dumpNode.dumpFilter.processLiteralArray
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
// 879: decl @lune.@base.@dumpNode.dumpFilter.processLiteralChar
func (self *dumpNode_dumpFilter) ProcessLiteralChar(_env *LnsEnv, node *Nodes_LiteralCharNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s(%s)", Lns_2DDD(node.FP.Get_num(_env), node.FP.Get_token(_env).Txt)))
}
// 885: decl @lune.@base.@dumpNode.dumpFilter.processLiteralInt
func (self *dumpNode_dumpFilter) ProcessLiteralInt(_env *LnsEnv, node *Nodes_LiteralIntNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s(%s)", Lns_2DDD(node.FP.Get_num(_env), node.FP.Get_token(_env).Txt)))
}
// 891: decl @lune.@base.@dumpNode.dumpFilter.processLiteralReal
func (self *dumpNode_dumpFilter) ProcessLiteralReal(_env *LnsEnv, node *Nodes_LiteralRealNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s(%s)", Lns_2DDD(node.FP.Get_num(_env), node.FP.Get_token(_env).Txt)))
}
// 897: decl @lune.@base.@dumpNode.dumpFilter.processLiteralString
func (self *dumpNode_dumpFilter) ProcessLiteralString(_env *LnsEnv, node *Nodes_LiteralStringNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, node.FP.Get_token(_env).Txt)
    {
        _expList := node.FP.Get_orgParam(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            for _, _param := range( expList.FP.Get_expList(_env).Items ) {
                param := _param
                dumpNode_filter_3_(_env, param, self, opt.FP.NextOpt(_env))
            }
        }
    }
}
// 907: decl @lune.@base.@dumpNode.dumpFilter.processLiteralBool
func (self *dumpNode_dumpFilter) ProcessLiteralBool(_env *LnsEnv, node *Nodes_LiteralBoolNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_token(_env).Txt == "true") &&
        _env.SetStackVal( "true") ||
        _env.SetStackVal( "false") ).(string))
}
// 913: decl @lune.@base.@dumpNode.dumpFilter.processLiteralNil
func (self *dumpNode_dumpFilter) ProcessLiteralNil(_env *LnsEnv, node *Nodes_LiteralNilNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}
// 918: decl @lune.@base.@dumpNode.dumpFilter.processBreak
func (self *dumpNode_dumpFilter) ProcessBreak(_env *LnsEnv, node *Nodes_BreakNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "")
}
// 923: decl @lune.@base.@dumpNode.dumpFilter.processLiteralSymbol
func (self *dumpNode_dumpFilter) ProcessLiteralSymbol(_env *LnsEnv, node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, _env.GetVM().String_format("%s: %s", Lns_2DDD(node.FP.Get_token(_env).Txt, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))))
}
// 929: decl @lune.@base.@dumpNode.dumpFilter.processAbbr
func (self *dumpNode_dumpFilter) ProcessAbbr(_env *LnsEnv, node *Nodes_AbbrNode,_opt LnsAny) {
    opt := _opt.(*DumpNode_Opt)
    self.FP.dump(_env, opt, &node.Nodes_Node, "##")
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
func DumpNode_Opt_toSlice(slice []LnsAny) []*DumpNode_Opt {
    ret := make([]*DumpNode_Opt, len(slice))
    for index, val := range slice {
        ret[index] = val.(DumpNode_OptDownCast).ToDumpNode_Opt()
    }
    return ret
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
// 40: DeclConstr
func (self *DumpNode_Opt) InitDumpNode_Opt(_env *LnsEnv, prefix string,depth LnsInt) {
    self.prefix = prefix
    self.depth = depth
    self.next = nil
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
    processDeclFuncInfo(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo, arg3 *DumpNode_Opt)
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
func dumpNode_dumpFilter_toSlice(slice []LnsAny) []*dumpNode_dumpFilter {
    ret := make([]*dumpNode_dumpFilter, len(slice))
    for index, val := range slice {
        ret[index] = val.(dumpNode_dumpFilterDownCast).TodumpNode_dumpFilter()
    }
    return ret
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

func Lns_dumpNode_init(_env *LnsEnv) {
    if init_dumpNode { return }
    init_dumpNode = true
    dumpNode__mod__ = "@lune.@base.@dumpNode"
    Lns_InitMod()
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Util_init(_env)
}
func init() {
    init_dumpNode = false
}
