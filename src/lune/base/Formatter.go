// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Formatter bool
var Formatter__mod__ string
// 58: decl @lune.@base.@Formatter.createFilter
func Formatter_createFilter(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,stream Lns_oStream) *Nodes_Filter {
    return &NewFormatter_FormatterFilter(_env, moduleTypeInfo, moduleTypeInfo.FP.Get_scope(_env), stream).Nodes_Filter
}

// 68: decl @lune.@base.@Formatter.filter
func Formatter_filter_3_(_env *LnsEnv, node *Nodes_Node,filter *Formatter_FormatterFilter,opt *Formatter_Opt) {
    filter.FP.OutputHeadComment(_env, node)
    node.FP.ProcessFilter(_env, &filter.Nodes_Filter, Formatter_Opt2Stem(opt))
}



// 40: decl @lune.@base.@Formatter.Opt.nextOpt
func (self *Formatter_Opt) NextOpt(_env *LnsEnv, parent *Nodes_Node) *Formatter_Opt {
    return NewFormatter_Opt(_env, parent)
}
// 62: decl @lune.@base.@Formatter.FormatterFilter.outputHeadComment
func (self *Formatter_FormatterFilter) OutputHeadComment(_env *LnsEnv, node *Nodes_Node) {
    for _, _commentNode := range( Lns_unwrapDefault( node.FP.Get_commentList(_env), NewLnsList([]LnsAny{})).(*LnsList).Items ) {
        commentNode := _commentNode.(Types_TokenDownCast).ToTypes_Token()
        self.FP.Writeln(_env, commentNode.Txt)
    }
}
// 77: decl @lune.@base.@Formatter.FormatterFilter.processNone
func (self *Formatter_FormatterFilter) ProcessNone(_env *LnsEnv, node *Nodes_NoneNode,_opt LnsAny) {
}
// 83: decl @lune.@base.@Formatter.FormatterFilter.processShebang
func (self *Formatter_FormatterFilter) ProcessShebang(_env *LnsEnv, node *Nodes_ShebangNode,_opt LnsAny) {
    self.FP.Writeln(_env, node.FP.Get_cmd(_env))
}
// 89: decl @lune.@base.@Formatter.FormatterFilter.processBlankLine
func (self *Formatter_FormatterFilter) ProcessBlankLine(_env *LnsEnv, node *Nodes_BlankLineNode,_opt LnsAny) {
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = node.FP.Get_lineNum(_env)
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            self.FP.Writeln(_env, "")
        }
    }
}
// 97: decl @lune.@base.@Formatter.FormatterFilter.processImport
func (self *Formatter_FormatterFilter) ProcessImport(_env *LnsEnv, node *Nodes_ImportNode,_opt LnsAny) {
    var info *Nodes_ImportInfo
    info = node.FP.Get_info(_env)
    self.FP.Write(_env, _env.GetVM().String_format("import %s", []LnsAny{info.FP.Get_modulePath(_env)}))
    if Lns_op_not(Lns_car(_env.GetVM().String_find(info.FP.Get_modulePath(_env),"%." + info.FP.Get_assignName(_env) + "$", nil, nil))){
        self.FP.Write(_env, _env.GetVM().String_format(" as %s", []LnsAny{info.FP.Get_assignName(_env)}))
    }
    self.FP.Writeln(_env, ";")
}
// 107: decl @lune.@base.@Formatter.FormatterFilter.processRoot
func (self *Formatter_FormatterFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    for _, _child := range( node.FP.Get_children(_env).Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_3_(_env, child, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
}
// 114: decl @lune.@base.@Formatter.FormatterFilter.processSubfile
func (self *Formatter_FormatterFilter) ProcessSubfile(_env *LnsEnv, node *Nodes_SubfileNode,_opt LnsAny) {
}
// 119: decl @lune.@base.@Formatter.FormatterFilter.processLuneControl
func (self *Formatter_FormatterFilter) ProcessLuneControl(_env *LnsEnv, node *Nodes_LuneControlNode,_opt LnsAny) {
    self.FP.Writeln(_env, "_lune_control ")
    switch _matchExp0 := node.FP.Get_pragma(_env).(type) {
    case *LuneControl_Pragma__default__init:
        self.FP.Writeln(_env, "default__init")
    case *LuneControl_Pragma__default__init_old:
        self.FP.Writeln(_env, "default__init_old")
    case *LuneControl_Pragma__disable_mut_control:
        self.FP.Writeln(_env, "disable_mut_control")
    case *LuneControl_Pragma__ignore_symbol_:
        self.FP.Writeln(_env, "ignore_symbol_")
    case *LuneControl_Pragma__load__lune_module:
        self.FP.Writeln(_env, "load__lune_module")
    case *LuneControl_Pragma__limit_conv_code:
    codeSet := _matchExp0.Val1
        self.FP.Writeln(_env, "limit_conv_code")
        for _code := range( codeSet.Items ) {
            code := _code.(string)
            self.FP.Write(_env, _env.GetVM().String_format(" %s", []LnsAny{code}))
        }
    case *LuneControl_Pragma__use_async:
        self.FP.Writeln(_env, "use_async")
    case *LuneControl_Pragma__run_async_pipe:
        self.FP.Writeln(_env, "run_async_pipe")
    case *LuneControl_Pragma__default_async_func:
        self.FP.Writeln(_env, "default_async_func")
    case *LuneControl_Pragma__default_async_all:
        self.FP.Writeln(_env, "default_async_all")
    case *LuneControl_Pragma__default_async_this_class:
        self.FP.Writeln(_env, "default_async_this_class")
    case *LuneControl_Pragma__default_noasync_this_class:
        self.FP.Writeln(_env, "default_noasync_this_class")
    case *LuneControl_Pragma__use_macro_special_var:
        self.FP.Writeln(_env, "use_macro_special_var")
    case *LuneControl_Pragma__single_phase_ast:
        self.FP.Writeln(_env, "single_phase_ast")
    }
    self.FP.Writeln(_env, ";")
}
// 173: decl @lune.@base.@Formatter.FormatterFilter.processRequest
func (self *Formatter_FormatterFilter) ProcessRequest(_env *LnsEnv, node *Nodes_RequestNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "__request ")
    Formatter_filter_3_(_env, node.FP.Get_processor(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, " ")
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 182: decl @lune.@base.@Formatter.FormatterFilter.processAsyncLock
func (self *Formatter_FormatterFilter) ProcessAsyncLock(_env *LnsEnv, node *Nodes_AsyncLockNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    if _switch0 := node.FP.Get_lockKind(_env); _switch0 == Nodes_LockKind__AsyncLock {
        self.FP.Writeln(_env, "__asyncLock {")
    } else if _switch0 == Nodes_LockKind__LuaGo {
        self.FP.Writeln(_env, "__luago {")
    } else if _switch0 == Nodes_LockKind__LuaLock {
        self.FP.Writeln(_env, "__luaLock {")
    } else if _switch0 == Nodes_LockKind__LuaDepend {
        self.FP.Writeln(_env, "__luaDepend {")
    }
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Writeln(_env, "}")
}
// 202: decl @lune.@base.@Formatter.FormatterFilter.processBlockSub
func (self *Formatter_FormatterFilter) ProcessBlockSub(_env *LnsEnv, node *Nodes_BlockNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    for _, _statement := range( node.FP.Get_stmtList(_env).Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_3_(_env, statement, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
    self.FP.PopIndent(_env)
    if _switch0 := node.FP.Get_blockKind(_env); _switch0 == Nodes_BlockKind__LetUnwrap || _switch0 == Nodes_BlockKind__Repeat {
        self.FP.Write(_env, "}")
    } else {
        self.FP.Writeln(_env, "}")
    }
}
// 223: decl @lune.@base.@Formatter.FormatterFilter.processStmtExp
func (self *Formatter_FormatterFilter) ProcessStmtExp(_env *LnsEnv, node *Nodes_StmtExpNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    {
        _tailComment := node.FP.Get_tailComment(_env)
        if !Lns_IsNil( _tailComment ) {
            tailComment := _tailComment.(*Types_Token)
            self.FP.Write(_env, "; ")
            self.FP.Writeln(_env, tailComment.Txt)
        } else {
            self.FP.Writeln(_env, ";")
        }
    }
}
// 237: decl @lune.@base.@Formatter.FormatterFilter.processDeclEnum
func (self *Formatter_FormatterFilter) ProcessDeclEnum(_env *LnsEnv, node *Nodes_DeclEnumNode,_opt LnsAny) {
    self.FP.Writeln(_env, _env.GetVM().String_format("enum %s {", []LnsAny{node.FP.Get_name(_env).Txt}))
    self.FP.PushIndent(_env, nil)
    for _, _name := range( node.FP.Get_valueNameList(_env).Items ) {
        name := _name.(Types_TokenDownCast).ToTypes_Token()
        self.FP.Write(_env, _env.GetVM().String_format("%s", []LnsAny{name.Txt}))
        self.FP.Writeln(_env, ",")
    }
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 252: decl @lune.@base.@Formatter.FormatterFilter.processDeclAlge
func (self *Formatter_FormatterFilter) ProcessDeclAlge(_env *LnsEnv, node *Nodes_DeclAlgeNode,_opt LnsAny) {
}
// 260: decl @lune.@base.@Formatter.FormatterFilter.processNewAlgeVal
func (self *Formatter_FormatterFilter) ProcessNewAlgeVal(_env *LnsEnv, node *Nodes_NewAlgeValNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    for _, _exp := range( node.FP.Get_paramList(_env).Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_3_(_env, exp, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
}
// 269: decl @lune.@base.@Formatter.FormatterFilter.outputDeclClass
func (self *Formatter_FormatterFilter) outputDeclClass(_env *LnsEnv, protoFlag bool,classType *Ast_TypeInfo,gluePrefix LnsAny,moduleName LnsAny) {
    if classType.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
        self.FP.Write(_env, "pub ")
    }
    if _switch0 := classType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Class {
        if Lns_isCondTrue( moduleName){
            self.FP.Write(_env, "module ")
        } else { 
            self.FP.Write(_env, "class ")
        }
    } else if _switch0 == Ast_TypeInfoKind__IF {
        self.FP.Write(_env, "interface ")
    }
    self.FP.Write(_env, classType.FP.Get_rawTxt(_env))
    if classType.FP.Get_itemTypeInfoList(_env).Len() > 0{
        self.FP.Write(_env, "<")
        for _index, _genType := range( classType.FP.Get_itemTypeInfoList(_env).Items ) {
            index := _index + 1
            genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index > 1{
                self.FP.Write(_env, ",")
            }
            self.FP.Write(_env, genType.FP.GetTxt(_env, nil, nil, nil))
        }
        self.FP.Write(_env, ">")
    }
    if moduleName != nil{
        moduleName_85 := moduleName.(*Types_Token)
        self.FP.Write(_env, " require ")
        self.FP.Write(_env, _env.GetVM().String_format("%s ", []LnsAny{moduleName_85.Txt}))
        if gluePrefix != nil{
            gluePrefix_87 := gluePrefix.(string)
            self.FP.Write(_env, "glue ")
            self.FP.Write(_env, gluePrefix_87)
        }
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( classType.FP.HasBase(_env)) ||
        _env.SetStackVal( classType.FP.Get_interfaceList(_env).Len() > 0) ).(bool){
        self.FP.Write(_env, " extend ")
        if classType.FP.HasBase(_env){
            self.FP.Write(_env, classType.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil))
        }
        if classType.FP.Get_interfaceList(_env).Len() > 0{
            self.FP.Write(_env, "(")
            for _index, _ifType := range( classType.FP.Get_interfaceList(_env).Items ) {
                index := _index + 1
                ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if index > 1{
                    self.FP.Write(_env, ",")
                }
                self.FP.Write(_env, ifType.FP.GetTxt(_env, nil, nil, nil))
            }
            self.FP.Write(_env, ")")
        }
    }
}
// 329: decl @lune.@base.@Formatter.FormatterFilter.processProtoClass
func (self *Formatter_FormatterFilter) ProcessProtoClass(_env *LnsEnv, node *Nodes_ProtoClassNode,_opt LnsAny) {
    self.FP.outputDeclClass(_env, true, node.FP.Get_expType(_env), nil, nil)
}
// 335: decl @lune.@base.@Formatter.FormatterFilter.processDeclClass
func (self *Formatter_FormatterFilter) ProcessDeclClass(_env *LnsEnv, node *Nodes_DeclClassNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.outputDeclClass(_env, false, node.FP.Get_expType(_env), node.FP.Get_gluePrefix(_env), node.FP.Get_moduleName(_env))
    self.FP.Writeln(_env, "")
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    for _, _stmt := range( node.FP.Get_allStmtList(_env).Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_3_(_env, stmt, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 409: decl @lune.@base.@Formatter.FormatterFilter.processDeclAdvertise
func (self *Formatter_FormatterFilter) ProcessDeclAdvertise(_env *LnsEnv, node *Nodes_DeclAdvertiseNode,_opt LnsAny) {
    self.FP.Writeln(_env, _env.GetVM().String_format("advertise %s;", []LnsAny{node.FP.Get_advInfo(_env).FP.Get_member(_env).FP.Get_name(_env).Txt}))
}
// 415: decl @lune.@base.@Formatter.FormatterFilter.processDeclMember
func (self *Formatter_FormatterFilter) ProcessDeclMember(_env *LnsEnv, node *Nodes_DeclMemberNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    if _switch0 := node.FP.Get_accessMode(_env); _switch0 == Ast_AccessMode__Pub || _switch0 == Ast_AccessMode__Pro || _switch0 == Ast_AccessMode__None || _switch0 == Ast_AccessMode__Local {
        self.FP.Write(_env, Ast_accessMode2txt(_env, node.FP.Get_accessMode(_env)))
        self.FP.Write(_env, " ")
    }
    if node.FP.Get_staticFlag(_env){
        self.FP.Write(_env, "static ")
    }
    self.FP.Write(_env, "let ")
    var symbol *Ast_SymbolInfo
    symbol = node.FP.Get_symbolInfo(_env)
    if symbol.FP.Get_mutable(_env){
        self.FP.Write(_env, "mut ")
    }
    self.FP.Write(_env, symbol.FP.Get_name(_env))
    self.FP.Write(_env, ":")
    Formatter_filter_3_(_env, &node.FP.Get_refType(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    if node.FP.Get_getterMode(_env) == Ast_AccessMode__None{
        if node.FP.Get_setterMode(_env) != Ast_AccessMode__None{
            self.FP.Write(_env, "{none,")
            self.FP.Write(_env, Ast_accessMode2txt(_env, node.FP.Get_setterMode(_env)))
            self.FP.Write(_env, "}")
        }
    } else { 
        self.FP.Write(_env, "{")
        self.FP.Write(_env, Ast_accessMode2txt(_env, node.FP.Get_getterMode(_env)))
        if Lns_op_not(node.FP.Get_getterMutable(_env)){
            self.FP.Write(_env, "&")
        }
        if node.FP.Get_getterRetType(_env) != symbol.FP.Get_typeInfo(_env){
            self.FP.Write(_env, ":")
            self.FP.Write(_env, node.FP.Get_getterRetType(_env).FP.GetTxt(_env, nil, nil, nil))
        }
        if node.FP.Get_setterMode(_env) != Ast_AccessMode__None{
            self.FP.Write(_env, ",")
            self.FP.Write(_env, Ast_accessMode2txt(_env, node.FP.Get_setterMode(_env)))
        }
        self.FP.Write(_env, "}")
    }
    self.FP.Writeln(_env, ";")
}
// 465: decl @lune.@base.@Formatter.FormatterFilter.processExpMacroExp
func (self *Formatter_FormatterFilter) ProcessExpMacroExp(_env *LnsEnv, node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList(_env)
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_3_(_env, stmt, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
}
// 474: decl @lune.@base.@Formatter.FormatterFilter.processDeclMacro
func (self *Formatter_FormatterFilter) ProcessDeclMacro(_env *LnsEnv, node *Nodes_DeclMacroNode,_opt LnsAny) {
}
// 479: decl @lune.@base.@Formatter.FormatterFilter.processExpMacroStat
func (self *Formatter_FormatterFilter) ProcessExpMacroStat(_env *LnsEnv, node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    for _, _strNode := range( node.FP.Get_expStrList(_env).Items ) {
        strNode := _strNode.(Nodes_NodeDownCast).ToNodes_Node()
        Formatter_filter_3_(_env, strNode, self, opt.FP.NextOpt(_env, strNode))
    }
}
// 489: decl @lune.@base.@Formatter.FormatterFilter.processUnwrapSet
func (self *Formatter_FormatterFilter) ProcessUnwrapSet(_env *LnsEnv, node *Nodes_UnwrapSetNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, &node.FP.Get_dstExpList(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    Formatter_filter_3_(_env, &node.FP.Get_srcExpList(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    if Lns_isCondTrue( node.FP.Get_unwrapBlock(_env)){
        Formatter_filter_3_(_env, &Lns_unwrap( node.FP.Get_unwrapBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
}
// 501: decl @lune.@base.@Formatter.FormatterFilter.processIfUnwrap
func (self *Formatter_FormatterFilter) ProcessIfUnwrap(_env *LnsEnv, node *Nodes_IfUnwrapNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "if! ")
    if node.FP.Get_varSymList(_env).Len() != 0{
        self.FP.Write(_env, "let ")
        for _index, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
            index := _index + 1
            varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if index > 1{
                self.FP.Write(_env, ",")
            }
            self.FP.Write(_env, varSym.FP.Get_name(_env))
        }
        self.FP.Write(_env, " = ")
    }
    Formatter_filter_3_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    if Lns_isCondTrue( node.FP.Get_nilBlock(_env)){
        self.FP.Write(_env, "else ")
        Formatter_filter_3_(_env, &Lns_unwrap( node.FP.Get_nilBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
}
// 523: decl @lune.@base.@Formatter.FormatterFilter.processWhen
func (self *Formatter_FormatterFilter) ProcessWhen(_env *LnsEnv, node *Nodes_WhenNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "when!")
    var symTxt string
    symTxt = " "
    for _, _symPair := range( node.FP.Get_symPairList(_env).Items ) {
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        symTxt = _env.GetVM().String_format("%s%s ", []LnsAny{symTxt, symPair.FP.Get_src(_env).FP.Get_name(_env)})
    }
    self.FP.Write(_env, symTxt)
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    {
        __exp := node.FP.Get_elseBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write(_env, "else ")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
}
// 542: decl @lune.@base.@Formatter.FormatterFilter.processDeclVar
func (self *Formatter_FormatterFilter) ProcessDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    if _switch0 := node.FP.Get_accessMode(_env); _switch0 == Ast_AccessMode__Pub {
        self.FP.Write(_env, "pub ")
    }
    if _switch1 := node.FP.Get_mode(_env); _switch1 == Nodes_DeclVarMode__Let {
        if node.FP.Get_unwrapFlag(_env){
            self.FP.Write(_env, "let! ")
        } else { 
            self.FP.Write(_env, "let ")
        }
    } else if _switch1 == Nodes_DeclVarMode__Unwrap {
        self.FP.Write(_env, "unwrap! ")
    }
    for _index, _symInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
        index := _index + 1
        symInfo := _symInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        if symInfo.FP.Get_mutable(_env){
            self.FP.Write(_env, "mut ")
        }
        self.FP.Write(_env, symInfo.FP.Get_name(_env))
        if node.FP.Get_varList(_env).Len() >= index{
            var varInfo *Nodes_VarInfo
            varInfo = node.FP.Get_varList(_env).GetAt(index).(Nodes_VarInfoDownCast).ToNodes_VarInfo()
            {
                _varType := varInfo.FP.Get_refType(_env)
                if !Lns_IsNil( _varType ) {
                    varType := _varType.(*Nodes_RefTypeNode)
                    self.FP.Write(_env, ":")
                    Formatter_filter_3_(_env, &varType.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
                }
            }
        }
    }
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(_env, " = ")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
    {
        __exp := node.FP.Get_unwrapBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write(_env, " ")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
    {
        __exp := node.FP.Get_thenBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write(_env, "then")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
    {
        __exp := node.FP.Get_syncBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write(_env, "do")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
    self.FP.Writeln(_env, ";")
}
// 600: decl @lune.@base.@Formatter.FormatterFilter.processDeclArg
func (self *Formatter_FormatterFilter) ProcessDeclArg(_env *LnsEnv, node *Nodes_DeclArgNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, node.FP.Get_symbolInfo(_env).FP.Get_name(_env))
    self.FP.Write(_env, ":")
    {
        _refType := node.FP.Get_argType(_env)
        if !Lns_IsNil( _refType ) {
            refType := _refType.(*Nodes_RefTypeNode)
            Formatter_filter_3_(_env, &refType.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        } else {
            self.FP.Write(_env, node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))
        }
    }
}
// 613: decl @lune.@base.@Formatter.FormatterFilter.processDeclArgDDD
func (self *Formatter_FormatterFilter) ProcessDeclArgDDD(_env *LnsEnv, node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    self.FP.Write(_env, "...")
}
// 618: decl @lune.@base.@Formatter.FormatterFilter.processExpSubDDD
func (self *Formatter_FormatterFilter) ProcessExpSubDDD(_env *LnsEnv, node *Nodes_ExpSubDDDNode,_opt LnsAny) {
}
// 625: decl @lune.@base.@Formatter.FormatterFilter.processDeclForm
func (self *Formatter_FormatterFilter) ProcessDeclForm(_env *LnsEnv, node *Nodes_DeclFormNode,_opt LnsAny) {
}
// 630: decl @lune.@base.@Formatter.FormatterFilter.processDeclFuncInfo
func (self *Formatter_FormatterFilter) processDeclFuncInfo(_env *LnsEnv, node *Nodes_Node,declInfo *Nodes_DeclFuncInfo,opt *Formatter_Opt) {
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_expType(_env)
    if funcType.FP.Get_accessMode(_env) == Ast_AccessMode__Pub{
        self.FP.Write(_env, "pub ")
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( declInfo.FP.Get_staticFlag(_env)) &&
        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(declInfo.FP.Get_name(_env)) && 
        _env.NilAccPush(_env.NilAccPop().(*Types_Token).Txt)) == "__init") ).(bool)){
        self.FP.Write(_env, "__init")
    } else { 
        self.FP.Write(_env, "fn ")
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( opt.FP.Get_parent(_env).FP.Get_kind(_env) != Nodes_NodeKind_get_DeclClass(_env)) &&
            _env.SetStackVal( funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) ).(bool)){
            var classType *Ast_TypeInfo
            classType = funcType.FP.Get_parentInfo(_env)
            self.FP.Write(_env, classType.FP.Get_rawTxt(_env))
            self.FP.Write(_env, ".")
        }
        {
            _nameToken := declInfo.FP.Get_name(_env)
            if !Lns_IsNil( _nameToken ) {
                nameToken := _nameToken.(*Types_Token)
                self.FP.Write(_env, nameToken.Txt)
            }
        }
        self.FP.Write(_env, "(")
        var argList *LnsList
        argList = declInfo.FP.Get_argList(_env)
        if argList.Len() != 0{
            self.FP.Write(_env, " ")
        }
        for _index, _arg := range( argList.Items ) {
            index := _index + 1
            arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
            if index > 1{
                self.FP.Write(_env, ", ")
            }
            Formatter_filter_3_(_env, arg, self, opt.FP.NextOpt(_env, node))
        }
        if argList.Len() != 0{
            self.FP.Write(_env, " ")
        }
        self.FP.Write(_env, ")")
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Ast_TypeInfo_isMut(_env, funcType)) &&
        _env.SetStackVal( declInfo.FP.Get_kind(_env) == Nodes_FuncKind__Mtd) ).(bool)){
        self.FP.Write(_env, " mut")
    }
    {
        _asyncMode := declInfo.FP.Get_asyncMode(_env)
        if !Lns_IsNil( _asyncMode ) {
            asyncMode := _asyncMode.(LnsInt)
            if _switch0 := asyncMode; _switch0 == Ast_Async__Async {
                self.FP.Write(_env, " __async")
            } else if _switch0 == Ast_Async__Noasync {
                self.FP.Write(_env, " __noasync")
            } else if _switch0 == Ast_Async__Transient {
                self.FP.Write(_env, " __trans")
            }
        }
    }
    if funcType.FP.Get_retTypeInfoList(_env).Len() != 0{
        self.FP.Write(_env, " : ")
        for _index, _retType := range( funcType.FP.Get_retTypeInfoList(_env).Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index > 1{
                self.FP.Write(_env, ", ")
            }
            self.FP.Write(_env, retType.FP.GetTxt(_env, nil, nil, nil))
        }
    }
    {
        __exp := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write(_env, " ")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, node))
        } else {
            self.FP.Writeln(_env, ";")
        }
    }
}
// 707: decl @lune.@base.@Formatter.FormatterFilter.processDeclFunc
func (self *Formatter_FormatterFilter) ProcessDeclFunc(_env *LnsEnv, node *Nodes_DeclFuncNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}
// 712: decl @lune.@base.@Formatter.FormatterFilter.processDeclMethod
func (self *Formatter_FormatterFilter) ProcessDeclMethod(_env *LnsEnv, node *Nodes_DeclMethodNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}
// 717: decl @lune.@base.@Formatter.FormatterFilter.processDeclConstr
func (self *Formatter_FormatterFilter) ProcessDeclConstr(_env *LnsEnv, node *Nodes_DeclConstrNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}
// 723: decl @lune.@base.@Formatter.FormatterFilter.processDeclDestr
func (self *Formatter_FormatterFilter) ProcessDeclDestr(_env *LnsEnv, node *Nodes_DeclDestrNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.processDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), opt)
}
// 729: decl @lune.@base.@Formatter.FormatterFilter.processExpCallSuperCtor
func (self *Formatter_FormatterFilter) ProcessExpCallSuperCtor(_env *LnsEnv, node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "super(")
    {
        _expListNode := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expListNode ) {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            Formatter_filter_3_(_env, &expListNode.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
    self.FP.Writeln(_env, ");")
}
// 739: decl @lune.@base.@Formatter.FormatterFilter.processExpCallSuper
func (self *Formatter_FormatterFilter) ProcessExpCallSuper(_env *LnsEnv, node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "super(")
    {
        _expListNode := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expListNode ) {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            Formatter_filter_3_(_env, &expListNode.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
    self.FP.Writeln(_env, ")")
}
// 749: decl @lune.@base.@Formatter.FormatterFilter.processRefType
func (self *Formatter_FormatterFilter) ProcessRefType(_env *LnsEnv, node *Nodes_RefTypeNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    {
        _mutMode := node.FP.Get_mutMode(_env)
        if !Lns_IsNil( _mutMode ) {
            mutMode := _mutMode.(LnsInt)
            if _switch0 := mutMode; _switch0 == Ast_MutMode__IMut {
                self.FP.Write(_env, "&")
            } else if _switch0 == Ast_MutMode__AllMut {
                self.FP.Write(_env, "allmut")
            }
        }
    }
    Formatter_filter_3_(_env, node.FP.Get_name(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType(_env).FP.Get_nonnilableType(_env)
    if _switch1 := expType.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__List || _switch1 == Ast_TypeInfoKind__Set {
        self.FP.Write(_env, "<")
        Formatter_filter_3_(_env, node.FP.Get_itemNodeList(_env).GetAt(1).(Nodes_NodeDownCast).ToNodes_Node(), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        self.FP.Write(_env, ">")
    } else if _switch1 == Ast_TypeInfoKind__Map {
        self.FP.Write(_env, "<")
        Formatter_filter_3_(_env, node.FP.Get_itemNodeList(_env).GetAt(1).(Nodes_NodeDownCast).ToNodes_Node(), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        self.FP.Write(_env, ",")
        Formatter_filter_3_(_env, node.FP.Get_itemNodeList(_env).GetAt(2).(Nodes_NodeDownCast).ToNodes_Node(), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        self.FP.Write(_env, ">")
    }
    if node.FP.Get_expType(_env).FP.Get_nilable(_env){
        self.FP.Write(_env, "!")
    }
}
// 786: decl @lune.@base.@Formatter.FormatterFilter.processIf
func (self *Formatter_FormatterFilter) ProcessIf(_env *LnsEnv, node *Nodes_IfNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList(_env)
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if _switch0 := stmt.FP.Get_kind(_env); _switch0 == Nodes_IfKind__If {
            self.FP.Write(_env, "if ")
        } else if _switch0 == Nodes_IfKind__ElseIf {
            self.FP.Write(_env, "elseif ")
        } else if _switch0 == Nodes_IfKind__Else {
            self.FP.Write(_env, "else ")
        }
        Formatter_filter_3_(_env, stmt.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        self.FP.Write(_env, " ")
        Formatter_filter_3_(_env, &stmt.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
}
// 807: decl @lune.@base.@Formatter.FormatterFilter.processSwitch
func (self *Formatter_FormatterFilter) ProcessSwitch(_env *LnsEnv, node *Nodes_SwitchNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    var caseList *LnsList
    caseList = node.FP.Get_caseList(_env)
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        Formatter_filter_3_(_env, &caseInfo.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        Formatter_filter_3_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
    {
        __exp := node.FP.Get_default(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_BlockNode)
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
}
// 821: decl @lune.@base.@Formatter.FormatterFilter.processMatch
func (self *Formatter_FormatterFilter) ProcessMatch(_env *LnsEnv, node *Nodes_MatchNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_val(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    var caseList *LnsList
    caseList = node.FP.Get_caseList(_env)
    for _, _caseInfo := range( caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        Formatter_filter_3_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
    {
        __exp := node.FP.Get_defaultBlock(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            Formatter_filter_3_(_env, _exp, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
}
// 834: decl @lune.@base.@Formatter.FormatterFilter.processWhile
func (self *Formatter_FormatterFilter) ProcessWhile(_env *LnsEnv, node *Nodes_WhileNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "while ")
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, " ")
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 842: decl @lune.@base.@Formatter.FormatterFilter.processRepeat
func (self *Formatter_FormatterFilter) ProcessRepeat(_env *LnsEnv, node *Nodes_RepeatNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "repeat ")
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, " ")
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Writeln(_env, ";")
}
// 851: decl @lune.@base.@Formatter.FormatterFilter.processFor
func (self *Formatter_FormatterFilter) ProcessFor(_env *LnsEnv, node *Nodes_ForNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, _env.GetVM().String_format("for %s = ", []LnsAny{node.FP.Get_val(_env).FP.Get_name(_env)}))
    Formatter_filter_3_(_env, node.FP.Get_init(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, ", ")
    Formatter_filter_3_(_env, node.FP.Get_to(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    {
        __exp := node.FP.Get_delta(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write(_env, ", ")
            Formatter_filter_3_(_env, _exp, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
    self.FP.Write(_env, " ")
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 865: decl @lune.@base.@Formatter.FormatterFilter.processApply
func (self *Formatter_FormatterFilter) ProcessApply(_env *LnsEnv, node *Nodes_ApplyNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "apply ")
    for _index, __var := range( node.FP.Get_varList(_env).Items ) {
        index := _index + 1
        _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        self.FP.Write(_env, _var.FP.Get_name(_env))
    }
    self.FP.Write(_env, " of ")
    Formatter_filter_3_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, " ")
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 881: decl @lune.@base.@Formatter.FormatterFilter.processForeach
func (self *Formatter_FormatterFilter) ProcessForeach(_env *LnsEnv, node *Nodes_ForeachNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "foreach ")
    self.FP.Write(_env, node.FP.Get_val(_env).FP.Get_name(_env))
    {
        _key := node.FP.Get_key(_env)
        if !Lns_IsNil( _key ) {
            key := _key.(*Ast_SymbolInfo)
            self.FP.Write(_env, ", ")
            self.FP.Write(_env, key.FP.Get_name(_env))
        }
    }
    self.FP.Write(_env, " in ")
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Writeln(_env, "")
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 895: decl @lune.@base.@Formatter.FormatterFilter.processForsort
func (self *Formatter_FormatterFilter) ProcessForsort(_env *LnsEnv, node *Nodes_ForsortNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "forsort ")
    self.FP.Write(_env, node.FP.Get_val(_env).FP.Get_name(_env))
    {
        _key := node.FP.Get_key(_env)
        if !Lns_IsNil( _key ) {
            key := _key.(*Ast_SymbolInfo)
            self.FP.Write(_env, ", ")
            self.FP.Write(_env, key.FP.Get_name(_env))
        }
    }
    self.FP.Write(_env, " in ")
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Writeln(_env, "")
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 910: decl @lune.@base.@Formatter.FormatterFilter.processExpUnwrap
func (self *Formatter_FormatterFilter) ProcessExpUnwrap(_env *LnsEnv, node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "unwrap ")
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    {
        __exp := node.FP.Get_default(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write(_env, " default ")
            Formatter_filter_3_(_env, _exp, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
}
// 932: decl @lune.@base.@Formatter.FormatterFilter.processExpCall
func (self *Formatter_FormatterFilter) ProcessExpCall(_env *LnsEnv, node *Nodes_ExpCallNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_func(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    if node.FP.Get_nilAccess(_env){
        self.FP.Write(_env, "$(")
    } else { 
        self.FP.Write(_env, "(")
    }
    {
        __exp := node.FP.Get_argList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(_env, " ")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
            self.FP.Write(_env, " ")
        }
    }
    self.FP.Write(_env, ")")
}
// 949: decl @lune.@base.@Formatter.FormatterFilter.processExpList
func (self *Formatter_FormatterFilter) ProcessExpList(_env *LnsEnv, node *Nodes_ExpListNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    var expList *LnsList
    expList = node.FP.Get_expList(_env)
    for _index, _exp := range( expList.Items ) {
        index := _index + 1
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        if index > 1{
            if exp.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpAccessMRet(_env){
                break
            }
            if exp.FP.Get_expType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Abbr{
                self.FP.Write(_env, ", ")
            } else { 
                self.FP.Write(_env, " ")
            }
        }
        Formatter_filter_3_(_env, exp, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
}
// 970: decl @lune.@base.@Formatter.FormatterFilter.processExpMRet
func (self *Formatter_FormatterFilter) ProcessExpMRet(_env *LnsEnv, node *Nodes_ExpMRetNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_mRet(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 976: decl @lune.@base.@Formatter.FormatterFilter.processExpAccessMRet
func (self *Formatter_FormatterFilter) ProcessExpAccessMRet(_env *LnsEnv, node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
}
// 982: decl @lune.@base.@Formatter.FormatterFilter.processExpOp1
func (self *Formatter_FormatterFilter) ProcessExpOp1(_env *LnsEnv, node *Nodes_ExpOp1Node,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, node.FP.Get_op(_env).Txt)
    if _switch0 := node.FP.Get_op(_env).Txt; _switch0 == "not" {
        self.FP.Write(_env, " ")
    }
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 994: decl @lune.@base.@Formatter.FormatterFilter.processExpToDDD
func (self *Formatter_FormatterFilter) ProcessExpToDDD(_env *LnsEnv, node *Nodes_ExpToDDDNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 1001: decl @lune.@base.@Formatter.FormatterFilter.processExpMultiTo1
func (self *Formatter_FormatterFilter) ProcessExpMultiTo1(_env *LnsEnv, node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 1008: decl @lune.@base.@Formatter.FormatterFilter.processExpCast
func (self *Formatter_FormatterFilter) ProcessExpCast(_env *LnsEnv, node *Nodes_ExpCastNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    {
        _refType := node.FP.Get_castTypeNode(_env)
        if !Lns_IsNil( _refType ) {
            refType := _refType.(*Nodes_RefTypeNode)
            self.FP.Write(_env, node.FP.Get_castOpe(_env))
            Formatter_filter_3_(_env, &refType.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
}
// 1017: decl @lune.@base.@Formatter.FormatterFilter.processExpParen
func (self *Formatter_FormatterFilter) ProcessExpParen(_env *LnsEnv, node *Nodes_ExpParenNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "(")
    Formatter_filter_3_(_env, node.FP.Get_exp(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, ")")
}
// 1024: decl @lune.@base.@Formatter.FormatterFilter.processExpSetVal
func (self *Formatter_FormatterFilter) ProcessExpSetVal(_env *LnsEnv, node *Nodes_ExpSetValNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_exp1(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, " = ")
    Formatter_filter_3_(_env, &node.FP.Get_exp2(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 1031: decl @lune.@base.@Formatter.FormatterFilter.processExpSetItem
func (self *Formatter_FormatterFilter) ProcessExpSetItem(_env *LnsEnv, node *Nodes_ExpSetItemNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_val(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    switch _matchExp0 := node.FP.Get_index(_env).(type) {
    case *Nodes_IndexVal__NodeIdx:
    index := _matchExp0.Val1
        self.FP.Write(_env, "[")
        Formatter_filter_3_(_env, index, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        self.FP.Write(_env, "]")
    case *Nodes_IndexVal__SymIdx:
    index := _matchExp0.Val1
        self.FP.Write(_env, _env.GetVM().String_format(".%s", []LnsAny{index}))
    }
    self.FP.Write(_env, " = ")
    Formatter_filter_3_(_env, node.FP.Get_exp2(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 1048: decl @lune.@base.@Formatter.FormatterFilter.processExpOp2
func (self *Formatter_FormatterFilter) ProcessExpOp2(_env *LnsEnv, node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_exp1(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, _env.GetVM().String_format(" %s ", []LnsAny{node.FP.Get_op(_env).Txt}))
    Formatter_filter_3_(_env, node.FP.Get_exp2(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 1055: decl @lune.@base.@Formatter.FormatterFilter.processExpNew
func (self *Formatter_FormatterFilter) ProcessExpNew(_env *LnsEnv, node *Nodes_ExpNewNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "new ")
    Formatter_filter_3_(_env, node.FP.Get_symbol(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, "(")
    {
        __exp := node.FP.Get_argList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
    self.FP.Write(_env, ")")
}
// 1067: decl @lune.@base.@Formatter.FormatterFilter.processExpRef
func (self *Formatter_FormatterFilter) ProcessExpRef(_env *LnsEnv, node *Nodes_ExpRefNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_symbolInfo(_env).FP.Get_name(_env))
}
// 1072: decl @lune.@base.@Formatter.FormatterFilter.processExpRefItem
func (self *Formatter_FormatterFilter) ProcessExpRefItem(_env *LnsEnv, node *Nodes_ExpRefItemNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_val(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    if node.FP.Get_nilAccess(_env){
        self.FP.Write(_env, "$")
    }
    {
        __exp := node.FP.Get_index(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write(_env, "[ ")
            Formatter_filter_3_(_env, _exp, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
            self.FP.Write(_env, " ]")
        } else {
            {
                __exp := node.FP.Get_symbol(_env)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(string)
                    self.FP.Write(_env, ".")
                    self.FP.Write(_env, _exp)
                }
            }
        }
    }
}
// 1091: decl @lune.@base.@Formatter.FormatterFilter.processRefField
func (self *Formatter_FormatterFilter) ProcessRefField(_env *LnsEnv, node *Nodes_RefFieldNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_prefix(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    if node.FP.Get_nilAccess(_env){
        self.FP.Write(_env, "$.")
    } else { 
        self.FP.Write(_env, ".")
    }
    self.FP.Write(_env, node.FP.Get_field(_env).Txt)
}
// 1103: decl @lune.@base.@Formatter.FormatterFilter.processExpOmitEnum
func (self *Formatter_FormatterFilter) ProcessExpOmitEnum(_env *LnsEnv, node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    self.FP.Write(_env, ".")
    self.FP.Write(_env, node.FP.Get_valToken(_env).Txt)
}
// 1110: decl @lune.@base.@Formatter.FormatterFilter.processGetField
func (self *Formatter_FormatterFilter) ProcessGetField(_env *LnsEnv, node *Nodes_GetFieldNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_prefix(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    if node.FP.Get_nilAccess(_env){
        self.FP.Write(_env, "$")
    }
    self.FP.Write(_env, ".$")
    self.FP.Write(_env, node.FP.Get_field(_env).Txt)
}
// 1121: decl @lune.@base.@Formatter.FormatterFilter.processReturn
func (self *Formatter_FormatterFilter) ProcessReturn(_env *LnsEnv, node *Nodes_ReturnNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "return")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(_env, " ")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
    self.FP.Writeln(_env, ";")
}
// 1133: decl @lune.@base.@Formatter.FormatterFilter.processProvide
func (self *Formatter_FormatterFilter) ProcessProvide(_env *LnsEnv, node *Nodes_ProvideNode,_opt LnsAny) {
}
// 1139: decl @lune.@base.@Formatter.FormatterFilter.processAlias
func (self *Formatter_FormatterFilter) ProcessAlias(_env *LnsEnv, node *Nodes_AliasNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, _env.GetVM().String_format("alias %s = ", []LnsAny{node.FP.Get_newSymbol(_env).FP.Get_name(_env)}))
    Formatter_filter_3_(_env, node.FP.Get_srcNode(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Writeln(_env, ";")
}
// 1147: decl @lune.@base.@Formatter.FormatterFilter.processTestCase
func (self *Formatter_FormatterFilter) ProcessTestCase(_env *LnsEnv, node *Nodes_TestCaseNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, _env.GetVM().String_format("__test %s(%s) {", []LnsAny{node.FP.Get_name(_env).Txt, node.FP.Get_ctrlName(_env)}))
    Formatter_filter_3_(_env, &node.FP.Get_block(_env).Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    self.FP.Write(_env, "}")
}
// 1155: decl @lune.@base.@Formatter.FormatterFilter.processBoxing
func (self *Formatter_FormatterFilter) ProcessBoxing(_env *LnsEnv, node *Nodes_BoxingNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_src(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 1162: decl @lune.@base.@Formatter.FormatterFilter.processUnboxing
func (self *Formatter_FormatterFilter) ProcessUnboxing(_env *LnsEnv, node *Nodes_UnboxingNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    Formatter_filter_3_(_env, node.FP.Get_src(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
}
// 1169: decl @lune.@base.@Formatter.FormatterFilter.processLiteralList
func (self *Formatter_FormatterFilter) ProcessLiteralList(_env *LnsEnv, node *Nodes_LiteralListNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "[")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(_env, " ")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
            self.FP.Write(_env, " ")
        }
    }
    self.FP.Write(_env, "]")
}
// 1182: decl @lune.@base.@Formatter.FormatterFilter.processLiteralSet
func (self *Formatter_FormatterFilter) ProcessLiteralSet(_env *LnsEnv, node *Nodes_LiteralSetNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "(@")
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(_env, " ")
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
            self.FP.Write(_env, " ")
        }
    }
    self.FP.Write(_env, ")")
}
// 1195: decl @lune.@base.@Formatter.FormatterFilter.processLiteralMap
func (self *Formatter_FormatterFilter) ProcessLiteralMap(_env *LnsEnv, node *Nodes_LiteralMapNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, "{")
    var pairList *LnsList
    pairList = node.FP.Get_pairList(_env)
    for _index, _pair := range( pairList.Items ) {
        index := _index + 1
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        Formatter_filter_3_(_env, pair.FP.Get_key(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        self.FP.Write(_env, ":")
        Formatter_filter_3_(_env, pair.FP.Get_val(_env), self, opt.FP.NextOpt(_env, &node.Nodes_Node))
    }
    self.FP.Write(_env, "}")
}
// 1213: decl @lune.@base.@Formatter.FormatterFilter.processLiteralArray
func (self *Formatter_FormatterFilter) ProcessLiteralArray(_env *LnsEnv, node *Nodes_LiteralArrayNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    {
        __exp := node.FP.Get_expList(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            Formatter_filter_3_(_env, &_exp.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
        }
    }
}
// 1221: decl @lune.@base.@Formatter.FormatterFilter.processLiteralChar
func (self *Formatter_FormatterFilter) ProcessLiteralChar(_env *LnsEnv, node *Nodes_LiteralCharNode,_opt LnsAny) {
    self.FP.Write(_env, _env.GetVM().String_format("?%s", []LnsAny{node.FP.Get_token(_env).Txt}))
}
// 1226: decl @lune.@base.@Formatter.FormatterFilter.processLiteralInt
func (self *Formatter_FormatterFilter) ProcessLiteralInt(_env *LnsEnv, node *Nodes_LiteralIntNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
}
// 1231: decl @lune.@base.@Formatter.FormatterFilter.processLiteralReal
func (self *Formatter_FormatterFilter) ProcessLiteralReal(_env *LnsEnv, node *Nodes_LiteralRealNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
}
// 1236: decl @lune.@base.@Formatter.FormatterFilter.processLiteralString
func (self *Formatter_FormatterFilter) ProcessLiteralString(_env *LnsEnv, node *Nodes_LiteralStringNode,_opt LnsAny) {
    opt := _opt.(*Formatter_Opt)
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
    {
        _expList := node.FP.Get_orgParam(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.Write(_env, " ( ")
            Formatter_filter_3_(_env, &expList.Nodes_Node, self, opt.FP.NextOpt(_env, &node.Nodes_Node))
            self.FP.Write(_env, " )")
        }
    }
}
// 1248: decl @lune.@base.@Formatter.FormatterFilter.processLiteralBool
func (self *Formatter_FormatterFilter) ProcessLiteralBool(_env *LnsEnv, node *Nodes_LiteralBoolNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
}
// 1253: decl @lune.@base.@Formatter.FormatterFilter.processLiteralNil
func (self *Formatter_FormatterFilter) ProcessLiteralNil(_env *LnsEnv, node *Nodes_LiteralNilNode,_opt LnsAny) {
    self.FP.Write(_env, "nil")
}
// 1258: decl @lune.@base.@Formatter.FormatterFilter.processBreak
func (self *Formatter_FormatterFilter) ProcessBreak(_env *LnsEnv, node *Nodes_BreakNode,_opt LnsAny) {
    self.FP.Writeln(_env, "break;")
}
// 1264: decl @lune.@base.@Formatter.FormatterFilter.processLiteralSymbol
func (self *Formatter_FormatterFilter) ProcessLiteralSymbol(_env *LnsEnv, node *Nodes_LiteralSymbolNode,_opt LnsAny) {
}
// 1270: decl @lune.@base.@Formatter.FormatterFilter.processAbbr
func (self *Formatter_FormatterFilter) ProcessAbbr(_env *LnsEnv, node *Nodes_AbbrNode,_opt LnsAny) {
    self.FP.Write(_env, "##")
}
// declaration Class -- Opt
type Formatter_OptMtd interface {
    Get_parent(_env *LnsEnv) *Nodes_Node
    NextOpt(_env *LnsEnv, arg1 *Nodes_Node) *Formatter_Opt
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
func NewFormatter_Opt(_env *LnsEnv, arg1 *Nodes_Node) *Formatter_Opt {
    obj := &Formatter_Opt{}
    obj.FP = obj
    obj.InitFormatter_Opt(_env, arg1)
    return obj
}
func (self *Formatter_Opt) Get_parent(_env *LnsEnv) *Nodes_Node{ return self.parent }
// 37: DeclConstr
func (self *Formatter_Opt) InitFormatter_Opt(_env *LnsEnv, parent *Nodes_Node) {
    self.parent = parent
}


// declaration Class -- FormatterFilter
type Formatter_FormatterFilterMtd interface {
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    GetFull(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl
    outputDeclClass(_env *LnsEnv, arg1 bool, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny)
    OutputHeadComment(_env *LnsEnv, arg1 *Nodes_Node)
    PopIndent(_env *LnsEnv)
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
    processDeclFuncInfo(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo, arg3 *Formatter_Opt)
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
    ProcessUnboxing(_env *LnsEnv, arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(_env *LnsEnv, arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(_env *LnsEnv, arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(_env *LnsEnv, arg1 *Nodes_WhileNode, arg2 LnsAny)
    PushIndent(_env *LnsEnv, arg1 LnsAny)
    ReturnToSource(_env *LnsEnv)
    SwitchToHeader(_env *LnsEnv)
    Write(_env *LnsEnv, arg1 string)
    WriteRaw(_env *LnsEnv, arg1 string)
    Writeln(_env *LnsEnv, arg1 string)
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
func NewFormatter_FormatterFilter(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsAny, arg3 Lns_oStream) *Formatter_FormatterFilter {
    obj := &Formatter_FormatterFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitFormatter_FormatterFilter(_env, arg1, arg2, arg3)
    return obj
}
// advertise -- 45
func (self *Formatter_FormatterFilter) PopIndent(_env *LnsEnv) {
self.stream. FP.PopIndent( _env)
}
// advertise -- 45
func (self *Formatter_FormatterFilter) PushIndent(_env *LnsEnv, arg1 LnsAny) {
self.stream. FP.PushIndent( _env, arg1)
}
// advertise -- 45
func (self *Formatter_FormatterFilter) ReturnToSource(_env *LnsEnv) {
self.stream. FP.ReturnToSource( _env)
}
// advertise -- 45
func (self *Formatter_FormatterFilter) SwitchToHeader(_env *LnsEnv) {
self.stream. FP.SwitchToHeader( _env)
}
// advertise -- 45
func (self *Formatter_FormatterFilter) Write(_env *LnsEnv, arg1 string) {
self.stream. FP.Write( _env, arg1)
}
// advertise -- 45
func (self *Formatter_FormatterFilter) WriteRaw(_env *LnsEnv, arg1 string) {
self.stream. FP.WriteRaw( _env, arg1)
}
// advertise -- 45
func (self *Formatter_FormatterFilter) Writeln(_env *LnsEnv, arg1 string) {
self.stream. FP.Writeln( _env, arg1)
}
// 48: DeclConstr
func (self *Formatter_FormatterFilter) InitFormatter_FormatterFilter(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,moduleInfoManager LnsAny,stream Lns_oStream) {
    self.InitNodes_Filter(_env, true, moduleTypeInfo, moduleInfoManager)
    self.stream = NewUtil_SimpleSourceOStream(_env, stream, nil, 3)
}


func Lns_Formatter_init(_env *LnsEnv) {
    if init_Formatter { return }
    init_Formatter = true
    Formatter__mod__ = "@lune.@base.@Formatter"
    Lns_InitMod()
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Parser_init(_env)
    Lns_Types_init(_env)
    Lns_Util_init(_env)
    Lns_LuneControl_init(_env)
}
func init() {
    init_Formatter = false
}
