// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_convLua bool
var convLua__mod__ string
// decl enum -- ConvMode 
const ConvLua_ConvMode__ConvMeta = 2
const ConvLua_ConvMode__Convert = 1
const ConvLua_ConvMode__Exec = 0
var ConvLua_ConvModeList_ = NewLnsList( []LnsAny {
  ConvLua_ConvMode__Exec,
  ConvLua_ConvMode__Convert,
  ConvLua_ConvMode__ConvMeta,
})
func ConvLua_ConvMode_get__allList() *LnsList{
    return ConvLua_ConvModeList_
}
var ConvLua_ConvModeMap_ = map[LnsInt]string {
  ConvLua_ConvMode__ConvMeta: "ConvMode.ConvMeta",
  ConvLua_ConvMode__Convert: "ConvMode.Convert",
  ConvLua_ConvMode__Exec: "ConvMode.Exec",
}
func ConvLua_ConvMode__from(arg1 LnsInt) LnsAny{
    if _, ok := ConvLua_ConvModeMap_[arg1]; ok { return arg1 }
    return nil
}

func ConvLua_ConvMode_getTxt(arg1 LnsInt) string {
    return ConvLua_ConvModeMap_[arg1];
}
var convLua_stepIndent LnsInt
type convLua_outputMacroStmtBlock_1378_ func ()
// for 3501
func convLua_convExp17690(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 3302
func convLua_convExp16544(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// for 254
func convLua_convExp970(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2704
func convLua_convExp13476(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2719
func convLua_convExp13551(arg1 []LnsAny) LnsInt {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt)
}
// 73: decl @lune.@base.@convLua.getSymTxt
func convLua_getSymTxt_1068_(name string,id string) string {
    if name == "_"{
        return Lns_getVM().String_format("_%s", []LnsAny{id})
    }
    return name
}

// 192: decl @lune.@base.@convLua.filter
func convLua_filter_1163_(node *Nodes_Node,filter *convLua_convFilter,parent *Nodes_Node) {
    node.FP.ProcessFilter(&filter.Nodes_Filter, ConvLua_Opt2Stem(NewConvLua_Opt(parent)))
}






// 1095: decl @lune.@base.@convLua.isGenericType
func convLua_isGenericType_1316_(typeInfo *Ast_TypeInfo) bool {
    if Ast_isGenericType(typeInfo){
        return true
    }
    if _switch5084 := typeInfo.FP.Get_kind(); _switch5084 == Ast_TypeInfoKind__Class || _switch5084 == Ast_TypeInfoKind__IF {
        if typeInfo.FP.Get_itemTypeInfoList().Len() > 0{
            return true
        }
    }
    return false
}





// 3358: decl @lune.@base.@convLua.createFilter
func ConvLua_createFilter(streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,moduleTypeInfo *Ast_TypeInfo,processInfo *Ast_ProcessInfo,moduleSymbolKind LnsInt,useLuneRuntime LnsAny,targetLuaVer *LuaVer_LuaVerInfo,enableTest bool,useIpairs bool) *Nodes_Filter {
    return &NewconvLua_convFilter(streamName, stream, metaStream, convMode, inMacro, moduleTypeInfo, processInfo, moduleSymbolKind, useLuneRuntime, targetLuaVer, enableTest, useIpairs).Nodes_Filter
}

// 3369: decl @lune.@base.@convLua.runLuaOnLns
func convLua_runLuaOnLns_1650_(code string)(LnsAny, string) {
    var loadFunc LnsAny
    var err string
    loadFunc,err = DependLuaOnLns_runLuaOnLns(code)
    if loadFunc != nil{
        loadFunc_7075 := loadFunc.(*Lns_luaValue)
        var mod LnsAny
        mod = Lns_getVM().RunLoadedfunc(loadFunc_7075,[]LnsAny{})[0]
        if mod != nil{
            mod_7078 := mod
            return mod_7078, ""
        }
        return nil, "load error"
    }
    return nil, err
}



// 3434: decl @lune.@base.@convLua.createAst
func ConvLua_createAst(scriptPath string,lnsCode string,option *Option_Option) *TransUnit_ASTInfo {
    var inStream *Parser_TxtStream
    inStream = NewParser_TxtStream(lnsCode)
    var moduleId *FrontInterface_ModuleId
    moduleId = FrontInterface_ModuleId_createId(0.0, 0)
    var moduleInfo *FrontInterface_ImportModuleInfo
    moduleInfo = NewFrontInterface_ImportModuleInfo()
    var mod string
    mod = Util_scriptPath2Module(scriptPath)
    var transUnit *TransUnit_TransUnit
    transUnit = NewTransUnit_TransUnit(moduleId, moduleInfo, &NewConvLua_MacroEvalImp().Nodes_MacroEval, nil, TransUnit_AnalyzeMode__Compile, nil, option.TargetLuaVer, option.TransCtrlInfo)
    var parser *Parser_StreamParser
    parser = NewParser_StreamParser(inStream.FP, scriptPath, false)
    return transUnit.FP.CreateAST(&parser.Parser_Parser, false, mod)
}

// 3450: decl @lune.@base.@convLua.getTestLnsCode
func ConvLua_getTestLnsCode() string {
    var lnsCode string
    lnsCode = "pub let mut outputList:List<str> = [];\nfn Print( ... ) {\n   let args = [ ... ];\n   let mut format = \"\";\n   foreach _, index in args {\n      if index > 1 {\n         format = \"%s\\t\" (format);\n      }\n      format = format .. \"%s\";\n   }\n   outputList.insert( string.format( format, ... ) );\n}\n\nPrint( \"hello world\" );\nmacro _hoge( list:List<int>) {\n  {\n     let mut statList:List<stat> = [];\n     foreach val in list {\n        statList.insert( `{ Print( ,,val + 10 ); } );\n     }\n  }\n   ,,statList;\n}\n_hoge( [1,2,3] );\n_ = 1;\n"
    return lnsCode
}

// 3482: decl @lune.@base.@convLua.Ast2Code
func convLua_Ast2Code_1689_(option *Option_Option,ast *TransUnit_ASTInfo,streamName string)(string, string) {
    var stream *Util_memStream
    stream = NewUtil_memStream()
    var metaStream *Util_memStream
    metaStream = NewUtil_memStream()
    var conv *Nodes_Filter
    conv = ConvLua_createFilter(streamName, stream.FP, metaStream.FP, ConvLua_ConvMode__ConvMeta, false, ast.FP.Get_moduleTypeInfo(), ast.FP.Get_processInfo(), ast.FP.Get_moduleSymbolKind(), option.UseLuneModule, option.TargetLuaVer, option.Testing, option.UseIpairs)
    ast.FP.Get_node().FP.ProcessFilter(conv, ConvLua_Opt2Stem(NewConvLua_Opt(ast.FP.Get_node())))
    return metaStream.FP.Get_txt(), stream.FP.Get_txt()
}

// 3498: decl @lune.@base.@convLua.runTestCode
func ConvLua_runTestCode(ctrl *Testing_Ctrl,luaCode string) {
    var loaded LnsAny
    var mess LnsAny
    loaded,mess = Lns_getVM().Load(luaCode, nil)
    if loaded != nil{
        loaded_7149 := loaded.(*Lns_luaValue)
        {
            _mod := convLua_convExp17690(Lns_2DDD(Lns_getVM().RunLoadedfunc(loaded_7149,Lns_2DDD([]LnsAny{}))[0]))
            if _mod != nil {
                mod := _mod
                {
                    _listobj := mod.(*Lns_luaValue).GetAt("outputList")
                    if _listobj != nil {
                        listobj := _listobj
                        var strList *Lns_luaValue
                        strList = listobj.(*Lns_luaValue)
                        if ctrl.FP.CheckEq(strList.Len(), 4, "#strList", "4", nil, convLua__mod__, 3504){
                            ctrl.FP.CheckEq(strList.GetAt(1).(string), "hello world", "strList[ 1 ]", "\"hello world\"", nil, convLua__mod__, 3505)
                            ctrl.FP.CheckEq(strList.GetAt(2).(string), "11", "strList[ 2 ]", "\"11\"", nil, convLua__mod__, 3506)
                            ctrl.FP.CheckEq(strList.GetAt(3).(string), "12", "strList[ 3 ]", "\"12\"", nil, convLua__mod__, 3507)
                            ctrl.FP.CheckEq(strList.GetAt(4).(string), "13", "strList[ 4 ]", "\"13\"", nil, convLua__mod__, 3508)
                        }
                    }
                }
            }
        }
    } else {
        Lns_print([]LnsAny{mess})
    }
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
func NewconvLua_PubVerInfo(arg1 bool, arg2 LnsInt, arg3 bool, arg4 *Ast_TypeInfo) *convLua_PubVerInfo {
    obj := &convLua_PubVerInfo{}
    obj.FP = obj
    obj.InitconvLua_PubVerInfo(arg1, arg2, arg3, arg4)
    return obj
}
func (self *convLua_PubVerInfo) InitconvLua_PubVerInfo(arg1 bool, arg2 LnsInt, arg3 bool, arg4 *Ast_TypeInfo) {
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
func NewconvLua_PubFuncInfo(arg1 LnsInt, arg2 *Ast_TypeInfo) *convLua_PubFuncInfo {
    obj := &convLua_PubFuncInfo{}
    obj.FP = obj
    obj.InitconvLua_PubFuncInfo(arg1, arg2)
    return obj
}
func (self *convLua_PubFuncInfo) InitconvLua_PubFuncInfo(arg1 LnsInt, arg2 *Ast_TypeInfo) {
    self.AccessMode = arg1
    self.TypeInfo = arg2
}

// declaration Class -- ModuleInfo
type convLua_ModuleInfoMtd interface {
    Get_assignName() string
    Get_modulePath() string
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
func NewconvLua_ModuleInfo(arg1 string, arg2 string) *convLua_ModuleInfo {
    obj := &convLua_ModuleInfo{}
    obj.FP = obj
    obj.InitconvLua_ModuleInfo(arg1, arg2)
    return obj
}
func (self *convLua_ModuleInfo) InitconvLua_ModuleInfo(arg1 string, arg2 string) {
    self.assignName = arg1
    self.modulePath = arg2
}
func (self *convLua_ModuleInfo) Get_assignName() string{ return self.assignName }
func (self *convLua_ModuleInfo) Get_modulePath() string{ return self.modulePath }

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
func NewConvLua_Opt(arg1 *Nodes_Node) *ConvLua_Opt {
    obj := &ConvLua_Opt{}
    obj.FP = obj
    obj.InitConvLua_Opt(arg1)
    return obj
}
func (self *ConvLua_Opt) InitConvLua_Opt(arg1 *Nodes_Node) {
    self.Node = arg1
}

// declaration Class -- convFilter
type convLua_convFilterMtd interface {
    Close()
    DefaultProcess(arg1 *Nodes_Node, arg2 LnsAny)
    Flush()
    getCanonicalName(arg1 *Ast_TypeInfo, arg2 bool) string
    getDestrClass(arg1 *Ast_TypeInfo) LnsAny
    getFullName(arg1 *Ast_TypeInfo) string
    getMapInfo(arg1 *Ast_TypeInfo)(string, bool, string)
    get_indent() LnsInt
    Get_moduleInfoManager() Ast_ModuleInfoManager
    Get_optStack() *LnsList
    Get_typeNameCtrl() *Ast_TypeNameCtrl
    outputAlter2MapFunc(arg1 Lns_oStream, arg2 *LnsMap)
    OutputDeclMacro(arg1 string, arg2 *LnsList, arg3 convLua_outputMacroStmtBlock_1378_)
    outputMeta(arg1 *Nodes_RootNode)
    popIndent()
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
    ProcessExpDDD(arg1 *Nodes_ExpDDDNode, arg2 LnsAny)
    ProcessExpList(arg1 *Nodes_ExpListNode, arg2 LnsAny)
    processExpListSub(arg1 *Nodes_Node, arg2 *LnsList, arg3 LnsAny)
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
    processLoadRuntime()
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
    process__func__symbol(arg1 bool, arg2 *Ast_TypeInfo, arg3 string)
    pushIndent(arg1 LnsAny)
    Write(arg1 string)(LnsAny, LnsAny)
    writeln(arg1 string)
}
type convLua_convFilter struct {
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
    FP convLua_convFilterMtd
}
func convLua_convFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convLua_convFilter).FP
}
type convLua_convFilterDownCast interface {
    ToconvLua_convFilter() *convLua_convFilter
}
func convLua_convFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convLua_convFilterDownCast)
    if ok { return work.ToconvLua_convFilter() }
    return nil
}
func (obj *convLua_convFilter) ToconvLua_convFilter() *convLua_convFilter {
    return obj
}
func NewconvLua_convFilter(arg1 string, arg2 Lns_oStream, arg3 Lns_oStream, arg4 LnsInt, arg5 bool, arg6 *Ast_TypeInfo, arg7 *Ast_ProcessInfo, arg8 LnsInt, arg9 LnsAny, arg10 *LuaVer_LuaVerInfo, arg11 bool, arg12 bool) *convLua_convFilter {
    obj := &convLua_convFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvLua_convFilter(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
    return obj
}
// 114: DeclConstr
func (self *convLua_convFilter) InitconvLua_convFilter(streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,moduleTypeInfo *Ast_TypeInfo,processInfo *Ast_ProcessInfo,moduleSymbolKind LnsInt,useLuneRuntime LnsAny,targetLuaVer *LuaVer_LuaVerInfo,enableTest bool,useIpairs bool) {
    self.InitNodes_Filter(true, moduleTypeInfo, moduleTypeInfo.FP.Get_scope())
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
    
}

// 148: decl @lune.@base.@convLua.convFilter.get_indent
func (self *convLua_convFilter) get_indent() LnsInt {
    if self.indentQueue.Len() > 0{
        return self.indentQueue.GetAt(self.indentQueue.Len()).(LnsInt)
    }
    return 0
}

// 155: decl @lune.@base.@convLua.convFilter.getCanonicalName
func (self *convLua_convFilter) getCanonicalName(typeInfo *Ast_TypeInfo,localFlag bool) string {
    var enumName string
    enumName = typeInfo.FP.GetFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), localFlag)
    return Lns_getVM().String_format("%s", []LnsAny{Lns_car(Lns_getVM().String_gsub(enumName,"&", "")).(string)})
}

// 161: decl @lune.@base.@convLua.convFilter.getFullName
func (self *convLua_convFilter) getFullName(typeInfo *Ast_TypeInfo) string {
    return self.FP.getCanonicalName(typeInfo, true)
}

// 165: decl @lune.@base.@convLua.convFilter.close
func (self *convLua_convFilter) Close() {
}

// 167: decl @lune.@base.@convLua.convFilter.flush
func (self *convLua_convFilter) Flush() {
}

// 170: decl @lune.@base.@convLua.convFilter.write
func (self *convLua_convFilter) Write(txt string)(LnsAny, LnsAny) {
    var stream Lns_oStream
    stream = self.stream
    if self.outMetaFlag{
        stream = self.metaStream
        
    }
    for _, _line := range( Str_getLineList(txt).Items ) {
        line := _line.(string)
        if self.needIndent{
            stream.Write(Lns_getVM().String_rep(" ", self.FP.get_indent()))
            self.needIndent = false
            
        }
        if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( len(line) > 0) &&
            Lns_setStackVal( LnsInt(line[len(line)-1]) == 10) ).(bool)){
            self.curLineNo = self.curLineNo + 1
            
        }
        stream.Write(line)
    }
    return self.FP, nil
}


// 207: decl @lune.@base.@convLua.convFilter.processBlankLine
func (self *convLua_convFilter) ProcessBlankLine(node *Nodes_BlankLineNode,_opt LnsAny) {
}

// 207: decl @lune.@base.@convLua.convFilter.processDeclForm
func (self *convLua_convFilter) ProcessDeclForm(node *Nodes_DeclFormNode,_opt LnsAny) {
}

// 207: decl @lune.@base.@convLua.convFilter.processProtoMethod
func (self *convLua_convFilter) ProcessProtoMethod(node *Nodes_ProtoMethodNode,_opt LnsAny) {
}

// 225: decl @lune.@base.@convLua.convFilter.pushIndent
func (self *convLua_convFilter) pushIndent(newIndent LnsAny) {
    var indent LnsInt
    indent = Lns_unwrapDefault( newIndent, self.FP.get_indent() + convLua_stepIndent).(LnsInt)
    self.indentQueue.Insert(indent)
}

// 230: decl @lune.@base.@convLua.convFilter.popIndent
func (self *convLua_convFilter) popIndent() {
    if self.indentQueue.Len() == 0{
        Util_err("self.indentQueue == 0")
    }
    self.indentQueue.Remove(nil)
}

// 238: decl @lune.@base.@convLua.convFilter.writeln
func (self *convLua_convFilter) writeln(txt string) {
    self.FP.Write(txt)
    self.FP.Write("\n")
    self.needIndent = true
    
}

// 244: decl @lune.@base.@convLua.convFilter.processNone
func (self *convLua_convFilter) ProcessNone(node *Nodes_NoneNode,_opt LnsAny) {
}

// 250: decl @lune.@base.@convLua.convFilter.processImport
func (self *convLua_convFilter) ProcessImport(node *Nodes_ImportNode,_opt LnsAny) {
    var module string
    module = node.FP.Get_modulePath()
    var moduleName string
    moduleName = convLua_convExp970(Lns_2DDD(Lns_getVM().String_gsub(module,".*%.", "")))
    moduleName = node.FP.Get_assignName()
    
    self.FP.Write(Lns_getVM().String_format("local %s = _lune.loadModule( '%s' )", []LnsAny{moduleName, module}))
}

// 284: decl @lune.@base.@convLua.convFilter.outputMeta
func (self *convLua_convFilter) outputMeta(node *Nodes_RootNode) {
    if self.convMode == ConvLua_ConvMode__Convert{
        return 
    }
    self.outMetaFlag = true
    
    if self.stream != self.metaStream{
        self.FP.writeln("local _moduleObj = {}")
    }
    self.FP.writeln("----- meta -----")
    self.FP.writeln(Lns_getVM().String_format("_moduleObj.__version = '%s'", []LnsAny{Ver_version}))
    self.FP.writeln(Lns_getVM().String_format("_moduleObj.__formatVersion = '%s'", []LnsAny{Ver_metaVersion}))
    self.FP.writeln(Lns_getVM().String_format("_moduleObj.__buildId = %q", []LnsAny{node.FP.Get_moduleId().FP.GetNextModuleId().FP.Get_idStr()}))
    self.FP.writeln(Lns_getVM().String_format("_moduleObj.__enableTest = %s", []LnsAny{self.enableTest}))
    self.FP.writeln(Lns_getVM().String_format("_moduleObj.__hasTest = %s", []LnsAny{node.FP.Get_nodeManager().FP.GetTestCaseNodeList().Len() != 0}))
    var importModuleType2Index *LnsMap
    importModuleType2Index = NewLnsMap( map[LnsAny]LnsAny{})
    var importNameMap *LnsMap
    importNameMap = NewLnsMap( map[LnsAny]LnsAny{})
    {
        for _typeInfo, _moduleInfo := range( node.FP.Get_importModule2moduleInfo().Items ) {
            typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            moduleInfo := _moduleInfo.(Nodes_ModuleInfoDownCast).ToNodes_ModuleInfo()
            importNameMap.Set(moduleInfo.FP.Get_fullName(),typeInfo)
        }
        var index LnsInt
        index = 0
        {
            __collection1209 := importNameMap
            __sorted1209 := __collection1209.CreateKeyListStr()
            __sorted1209.Sort( LnsItemKindStr, nil )
            for _, ___key1209 := range( __sorted1209.Items ) {
                typeInfo := __collection1209.Items[ ___key1209 ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                index = index + 1
                
                importModuleType2Index.Set(typeInfo,index)
            }
        }
    }
    var typeId2TypeInfo *LnsMap
    typeId2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    var pickupClassMap *LnsMap
    pickupClassMap = NewLnsMap( map[LnsAny]LnsAny{})
    var checkExportTypeInfo func(typeInfo *Ast_TypeInfo) LnsAny
    checkExportTypeInfo = func(typeInfo *Ast_TypeInfo) LnsAny {
        var moduleTypeInfo *Ast_TypeInfo
        moduleTypeInfo = typeInfo.FP.GetModule()
        var typeId LnsInt
        typeId = typeInfo.FP.Get_typeId()
        return Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( typeId2TypeInfo.Items[typeId]) &&
            Lns_setStackVal( Lns_op_not(Ast_isBuiltin(typeId))) &&
            Lns_setStackVal( (Lns_popVal( Lns_incStack() ||
                Lns_setStackVal( moduleTypeInfo.FP.HasRouteNamespaceFrom(node.FP.Get_moduleTypeInfo())) ||
                Lns_setStackVal( typeInfo.FP.Get_srcTypeInfo() != typeInfo) ||
                Lns_setStackVal( moduleTypeInfo.FP.Equals(self.processInfo, Ast_headTypeInfo, nil, nil)) ).(bool))) )
    }
    var isDependOnExt func(typeInfo *Ast_TypeInfo) bool
    isDependOnExt = func(typeInfo *Ast_TypeInfo) bool {
        if Ast_isExtId(typeInfo){
            return true
        }
        return self.moduleTypeInfo.FP.Get_processInfo() != typeInfo.FP.Get_processInfo()
    }
    var pickupTypeId func(typeInfo *Ast_TypeInfo,forceFlag LnsAny,pickupChildFlag LnsAny)
    pickupTypeId = func(typeInfo *Ast_TypeInfo,forceFlag LnsAny,pickupChildFlag LnsAny) {
        if typeInfo.FP.Get_typeId() == Ast_rootTypeId{
            return 
        }
        if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( Lns_op_not(forceFlag)) &&
            Lns_setStackVal( Lns_op_not(Ast_isPubToExternal(typeInfo.FP.Get_accessMode()))) ).(bool)){
            return 
        }
        if Lns_isCondTrue( typeId2TypeInfo.Items[typeInfo.FP.Get_typeId()]){
            if isDependOnExt(typeInfo){
                return 
            }
            if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                Lns_setStackVal( pickupChildFlag) &&
                Lns_setStackVal( Lns_op_not(typeInfo.FP.Get_nilable())) )){
                for _, _itemTypeInfo := range( typeInfo.FP.Get_children().Items ) {
                    itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if Ast_isPubToExternal(itemTypeInfo.FP.Get_accessMode()){
                        if _switch1446 := itemTypeInfo.FP.Get_kind(); _switch1446 == Ast_TypeInfoKind__Class || _switch1446 == Ast_TypeInfoKind__IF || _switch1446 == Ast_TypeInfoKind__Form || _switch1446 == Ast_TypeInfoKind__FormFunc || _switch1446 == Ast_TypeInfoKind__Func || _switch1446 == Ast_TypeInfoKind__Method {
                            pickupTypeId(itemTypeInfo, true, true)
                        }
                    }
                }
            }
            return 
        }
        typeId2TypeInfo.Set(typeInfo.FP.Get_typeId(),typeInfo)
        if typeInfo.FP.IsModule(){
            return 
        }
        if Ast_isBuiltin(typeInfo.FP.Get_srcTypeInfo().FP.Get_typeId()){
            return 
        }
        if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Ext{
            pickupTypeId(typeInfo.FP.Get_extedType(), true, true)
        }
        if typeInfo != typeInfo.FP.Get_srcTypeInfo(){
            pickupTypeId(typeInfo.FP.Get_srcTypeInfo(), true, false)
        } else if typeInfo.FP.Get_nilable(){
            pickupTypeId(typeInfo.FP.Get_nonnilableType(), true, false)
        } else { 
            if isDependOnExt(typeInfo){
                return 
            }
            if Lns_popVal( Lns_incStack() ||
                Lns_setStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
                Lns_setStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__IF) ).(bool){
                pickupClassMap.Set(typeInfo.FP.Get_typeId(),typeInfo)
            }
            if Lns_op_not(typeInfo.FP.Get_externalFlag()){
                if _switch1652 := typeInfo.FP.Get_kind(); _switch1652 == Ast_TypeInfoKind__IF || _switch1652 == Ast_TypeInfoKind__Class || _switch1652 == Ast_TypeInfoKind__Form || _switch1652 == Ast_TypeInfoKind__FormFunc || _switch1652 == Ast_TypeInfoKind__Alge || _switch1652 == Ast_TypeInfoKind__Enum || _switch1652 == Ast_TypeInfoKind__Map || _switch1652 == Ast_TypeInfoKind__Set || _switch1652 == Ast_TypeInfoKind__List || _switch1652 == Ast_TypeInfoKind__Array || _switch1652 == Ast_TypeInfoKind__Alternate || _switch1652 == Ast_TypeInfoKind__Box {
                    pickupTypeId(typeInfo.FP.Get_nilableTypeInfo(), true, false)
                    var imutType *Ast_TypeInfo
                    imutType = self.processInfo.FP.CreateModifier(typeInfo, Ast_MutMode__IMut)
                    pickupTypeId(imutType, true, false)
                }
            }
            var parentInfo *Ast_TypeInfo
            parentInfo = typeInfo.FP.Get_parentInfo()
            pickupTypeId(parentInfo, true, false)
            pickupTypeId(typeInfo.FP.Get_genSrcTypeInfo(), true, false)
            var baseInfo *Ast_TypeInfo
            baseInfo = typeInfo.FP.Get_baseTypeInfo()
            if baseInfo.FP.Get_typeId() != Ast_rootTypeId{
                pickupTypeId(baseInfo, true, true)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_interfaceList().Items ) {
                itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                pickupTypeId(itemTypeInfo, true, true)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_itemTypeInfoList().Items ) {
                itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                pickupTypeId(itemTypeInfo, true, false)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_argTypeInfoList().Items ) {
                itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                pickupTypeId(itemTypeInfo, true, false)
            }
            for _, _itemTypeInfo := range( typeInfo.FP.Get_retTypeInfoList().Items ) {
                itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                pickupTypeId(itemTypeInfo, true, true)
            }
            if Lns_isCondTrue( pickupChildFlag){
                for _, _itemTypeInfo := range( typeInfo.FP.Get_children().Items ) {
                    itemTypeInfo := _itemTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if itemTypeInfo.FP.Get_accessMode() == Ast_AccessMode__Pub{
                        if _switch1816 := itemTypeInfo.FP.Get_kind(); _switch1816 == Ast_TypeInfoKind__Class || _switch1816 == Ast_TypeInfoKind__IF || _switch1816 == Ast_TypeInfoKind__Form || _switch1816 == Ast_TypeInfoKind__FormFunc || _switch1816 == Ast_TypeInfoKind__Func || _switch1816 == Ast_TypeInfoKind__Method {
                            pickupTypeId(itemTypeInfo, true, true)
                        }
                    }
                }
            }
        }
    }
    var classId2TypeInfo *LnsMap
    classId2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    var validChildrenSet *LnsMap
    validChildrenSet = NewLnsMap( map[LnsAny]LnsAny{})
    {
        var typeInfo *Ast_TypeInfo
        typeInfo = self.moduleTypeInfo
        for typeInfo != Ast_headTypeInfo {
            validChildrenSet.Set(typeInfo.FP.Get_parentInfo(),NewLnsMap( map[LnsAny]LnsAny{typeInfo.FP.Get_typeId():typeInfo,}))
            typeInfo = typeInfo.FP.Get_parentInfo()
            
        }
        pickupTypeId(self.moduleTypeInfo, true, nil)
    }
    var typeId2ClassMap *LnsMap
    typeId2ClassMap = node.FP.Get_typeId2ClassMap()
    for _, _namespaceInfo := range( typeId2ClassMap.Items ) {
        namespaceInfo := _namespaceInfo.(Nodes_NamespaceInfoDownCast).ToNodes_NamespaceInfo()
        if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( namespaceInfo.TypeInfo.FP.Get_accessMode() == Ast_AccessMode__Pub) &&
            Lns_setStackVal( Lns_op_not(namespaceInfo.TypeInfo.FP.Get_externalFlag())) ).(bool)){
            pickupClassMap.Set(namespaceInfo.TypeInfo.FP.Get_typeId(),namespaceInfo.TypeInfo)
        }
    }
    classId2TypeInfo = self.classId2TypeInfo
    
    self.FP.writeln("local __typeId2ClassInfoMap = {}")
    self.FP.writeln("_moduleObj.__typeId2ClassInfoMap = __typeId2ClassInfoMap")
    {
        __collection2179 := classId2TypeInfo
        __sorted2179 := __collection2179.CreateKeyListInt()
        __sorted2179.Sort( LnsItemKindInt, nil )
        for _, _classTypeId := range( __sorted2179.Items ) {
            classTypeInfo := __collection2179.Items[ _classTypeId ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            classTypeId := _classTypeId.(LnsInt)
            if classTypeInfo.FP.Get_accessMode() == Ast_AccessMode__Pub{
                pickupTypeId(classTypeInfo, true, Lns_popVal( Lns_incStack() ||
                    Lns_setStackVal( validChildrenSet.Items[classTypeInfo] == nil) &&
                    Lns_setStackVal( Lns_op_not(classTypeInfo.FP.Get_externalFlag())) ).(bool))
                pickupClassMap.Set(classTypeId,nil)
                self.FP.writeln("do")
                self.FP.pushIndent(nil)
                self.FP.writeln(Lns_getVM().String_format("local __classInfo%d = {}", []LnsAny{classTypeId}))
                self.FP.writeln(Lns_getVM().String_format("__typeId2ClassInfoMap[ %d ] = __classInfo%d", []LnsAny{classTypeId, classTypeId}))
                for _, _memberNode := range( Lns_unwrap( self.classId2MemberList.Items[classTypeId]).(*LnsList).Items ) {
                    memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
                    if memberNode.FP.Get_accessMode() != Ast_AccessMode__Pri{
                        var memberName string
                        memberName = memberNode.FP.Get_name().Txt
                        var memberTypeInfo *Ast_TypeInfo
                        memberTypeInfo = memberNode.FP.Get_expType()
                        self.FP.writeln(Lns_getVM().String_format("__classInfo%d.%s = {", []LnsAny{classTypeId, memberName}))
                        self.FP.writeln(Lns_getVM().String_format("  name='%s', staticFlag = %s, mutMode = %d,", []LnsAny{memberName, memberNode.FP.Get_staticFlag(), memberNode.FP.Get_symbolInfo().FP.Get_mutMode()}) + Lns_getVM().String_format("accessMode = %d, typeId = %d }", []LnsAny{memberNode.FP.Get_accessMode(), memberTypeInfo.FP.Get_typeId()}))
                        pickupTypeId(memberTypeInfo, true, nil)
                    }
                }
                self.FP.popIndent()
                self.FP.writeln("end")
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
            classTypeId := _classTypeId.(LnsInt)
            classTypeInfo := _classTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(pickupedClassMap.Items[classTypeId]){
                pickupedClassMap.Set(classTypeId,classTypeInfo)
                workClassMap.Set(classTypeId,classTypeInfo)
                hasWorkClassFlag = true
                
            }
        }
        if Lns_op_not(hasWorkClassFlag){
            break
        }
        {
            __collection2488 := workClassMap
            __sorted2488 := __collection2488.CreateKeyListInt()
            __sorted2488.Sort( LnsItemKindInt, nil )
            for _, _classTypeId := range( __sorted2488.Items ) {
                classTypeInfo := __collection2488.Items[ _classTypeId ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                classTypeId := _classTypeId.(LnsInt)
                if Lns_op_not(Ast_isBuiltin(classTypeId)){
                    var scope *Ast_Scope
                    
                    {
                        _scope := classTypeInfo.FP.Get_scope()
                        if _scope == nil{
                            Util_err(Lns_getVM().String_format("%s.scope is nil", []LnsAny{classTypeInfo.FP.GetTxt(nil, nil, nil)}))
                        } else {
                            scope = _scope.(*Ast_Scope)
                        }
                    }
                    pickupTypeId(classTypeInfo, true, Lns_popVal( Lns_incStack() ||
                        Lns_setStackVal( validChildrenSet.Items[classTypeInfo] == nil) &&
                        Lns_setStackVal( Lns_op_not(classTypeInfo.FP.Get_externalFlag())) ).(bool))
                    if Lns_isCondTrue( checkExportTypeInfo(classTypeInfo)){
                        self.FP.writeln("do")
                        self.FP.pushIndent(nil)
                        self.FP.writeln(Lns_getVM().String_format("local __classInfo%s = {}", []LnsAny{classTypeId}))
                        self.FP.writeln(Lns_getVM().String_format("__typeId2ClassInfoMap[ %d ] = __classInfo%d", []LnsAny{classTypeId, classTypeId}))
                        {
                            __collection2466 := scope.FP.Get_symbol2SymbolInfoMap()
                            __sorted2466 := __collection2466.CreateKeyListStr()
                            __sorted2466.Sort( LnsItemKindStr, nil )
                            for _, _fieldName := range( __sorted2466.Items ) {
                                symbolInfo := __collection2466.Items[ _fieldName ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                                fieldName := _fieldName.(string)
                                var typeInfo *Ast_TypeInfo
                                typeInfo = symbolInfo.FP.Get_typeInfo()
                                if Lns_popVal( Lns_incStack() ||
                                    Lns_setStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mbr) ||
                                    Lns_setStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Var) ).(bool){
                                    if symbolInfo.FP.Get_accessMode() == Ast_AccessMode__Pub{
                                        self.FP.writeln(Lns_getVM().String_format("__classInfo%d.%s = {", []LnsAny{classTypeId, fieldName}))
                                        self.FP.writeln(Lns_getVM().String_format("  name='%s', staticFlag = %s, ", []LnsAny{fieldName, symbolInfo.FP.Get_staticFlag()}) + Lns_getVM().String_format("accessMode = %d, typeId = %d }", []LnsAny{symbolInfo.FP.Get_accessMode(), typeInfo.FP.Get_typeId()}))
                                        pickupTypeId(typeInfo, nil, nil)
                                    }
                                }
                            }
                        }
                        self.FP.popIndent()
                        self.FP.writeln("end")
                    }
                }
            }
        }
    }
    self.FP.writeln("local __macroName2InfoMap = {}")
    self.FP.writeln("_moduleObj.__macroName2InfoMap = __macroName2InfoMap")
    for _, _macroDeclNode := range( node.FP.Get_nodeManager().FP.GetDeclMacroNodeList().Items ) {
        macroDeclNode := _macroDeclNode.(Nodes_DeclMacroNodeDownCast).ToNodes_DeclMacroNode()
        var declInfo *Nodes_DeclMacroInfo
        declInfo = macroDeclNode.FP.Get_declInfo()
        if declInfo.FP.Get_pubFlag(){
            var macroInfo *Nodes_MacroInfo
            macroInfo = Lns_unwrap( node.FP.Get_typeId2MacroInfo().Items[macroDeclNode.FP.Get_expType().FP.Get_typeId()]).(*Nodes_MacroInfo)
            var macroTypeInfo *Ast_TypeInfo
            macroTypeInfo = macroDeclNode.FP.Get_expType()
            pickupTypeId(macroTypeInfo, true, nil)
            self.FP.writeln("do")
            self.FP.pushIndent(nil)
            self.FP.writeln("local info = {}")
            self.FP.writeln(Lns_getVM().String_format("__macroName2InfoMap[ %d ] = info", []LnsAny{macroTypeInfo.FP.Get_typeId()}))
            self.FP.writeln(Lns_getVM().String_format("info.name = %q", []LnsAny{declInfo.FP.Get_name().Txt}))
            self.FP.Write("info.argList = {")
            for _index, _argNode := range( declInfo.FP.Get_argList().Items ) {
                index := _index + 1
                argNode := _argNode.(Nodes_DeclArgNodeDownCast).ToNodes_DeclArgNode()
                if index != 1{
                    self.FP.Write(",")
                }
                self.FP.Write(Lns_getVM().String_format("{name=%q,typeId=%d}", []LnsAny{argNode.FP.Get_name().Txt, argNode.FP.Get_expType().FP.Get_typeId()}))
            }
            self.FP.writeln("}")
            self.FP.Write("info.symList = {")
            var firstFlag bool
            firstFlag = true
            {
                __collection2727 := macroInfo.Symbol2MacroValInfoMap
                __sorted2727 := __collection2727.CreateKeyListStr()
                __sorted2727.Sort( LnsItemKindStr, nil )
                for _, _name := range( __sorted2727.Items ) {
                    symInfo := __collection2727.Items[ _name ].(Nodes_MacroValInfoDownCast).ToNodes_MacroValInfo()
                    name := _name.(string)
                    if firstFlag{
                        firstFlag = false
                        
                    } else { 
                        self.FP.Write(",")
                    }
                    self.FP.Write(Lns_getVM().String_format("{name=%q,typeId=%d}", []LnsAny{name, symInfo.TypeInfo.FP.Get_typeId()}))
                    pickupTypeId(symInfo.TypeInfo, true, nil)
                }
            }
            self.FP.writeln("}")
            {
                _stmtBlock := declInfo.FP.Get_stmtBlock()
                if _stmtBlock != nil {
                    stmtBlock := _stmtBlock.(*Nodes_BlockNode)
                    var memStream *Util_memStream
                    memStream = NewUtil_memStream()
                    var workFilter *convLua_convFilter
                    workFilter = NewconvLua_convFilter(declInfo.FP.Get_name().Txt, memStream.FP, memStream.FP, ConvLua_ConvMode__Convert, false, Ast_headTypeInfo, self.processInfo, Ast_SymbolKind__Typ, self.useLuneRuntime, self.targetLuaVer, self.enableTest, self.useIpairs)
                    workFilter.macroDepth = workFilter.macroDepth + 1
                    
                    workFilter.FP.ProcessBlock(stmtBlock, ConvLua_Opt2Stem(NewConvLua_Opt(&node.Nodes_Node)))
                    workFilter.macroDepth = workFilter.macroDepth - 1
                    
                    memStream.FP.Close()
                    self.FP.writeln(Lns_getVM().String_format("info.stmtBlock = %q", []LnsAny{memStream.FP.Get_txt()}))
                }
            }
            self.FP.writeln("info.tokenList = {")
            var prevLineNo LnsInt
            prevLineNo = -1
            for _index, _token := range( declInfo.FP.Get_tokenList().Items ) {
                index := _index + 1
                token := _token.(Parser_TokenDownCast).ToParser_Token()
                if index > 1{
                    self.FP.Write(",")
                }
                if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                    Lns_setStackVal( prevLineNo != -1) &&
                    Lns_setStackVal( prevLineNo != token.Pos.LineNo) ).(bool)){
                    self.FP.Write(Lns_getVM().String_format("{%d,%q},", []LnsAny{Parser_TokenKind__Dlmt, "\n"}))
                }
                prevLineNo = token.Pos.LineNo
                
                self.FP.Write(Lns_getVM().String_format("{%d,%q}", []LnsAny{token.Kind, token.Txt}))
            }
            self.FP.writeln("}")
            self.FP.popIndent()
            self.FP.writeln("end")
        }
    }
    self.FP.writeln("local __varName2InfoMap = {}")
    self.FP.writeln("_moduleObj.__varName2InfoMap = __varName2InfoMap")
    {
        __collection3050 := self.pubVarName2InfoMap
        __sorted3050 := __collection3050.CreateKeyListStr()
        __sorted3050.Sort( LnsItemKindStr, nil )
        for _, _varName := range( __sorted3050.Items ) {
            varInfo := __collection3050.Items[ _varName ].(convLua_PubVerInfoDownCast).ToconvLua_PubVerInfo()
            varName := _varName.(string)
            self.FP.writeln(Lns_getVM().String_format("__varName2InfoMap.%s = {", []LnsAny{varName}))
            self.FP.writeln(Lns_getVM().String_format("  name='%s', accessMode = %d, typeId = %d, mutable = %s }", []LnsAny{varName, varInfo.AccessMode, varInfo.TypeInfo.FP.Get_typeId(), true}))
            pickupTypeId(varInfo.TypeInfo, true, nil)
        }
    }
    {
        __collection3067 := self.pubFuncName2InfoMap
        __sorted3067 := __collection3067.CreateKeyListStr()
        __sorted3067.Sort( LnsItemKindStr, nil )
        for _, ___key3067 := range( __sorted3067.Items ) {
            funcInfo := __collection3067.Items[ ___key3067 ].(convLua_PubFuncInfoDownCast).ToconvLua_PubFuncInfo()
            pickupTypeId(funcInfo.TypeInfo, true, nil)
        }
    }
    for _, _aliasNode := range( node.FP.Get_nodeManager().FP.GetAliasNodeList().Items ) {
        aliasNode := _aliasNode.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
        pickupTypeId(aliasNode.FP.Get_expType(), false, nil)
    }
    self.FP.writeln("local __typeInfoList = {}")
    self.FP.writeln("_moduleObj.__typeInfoList = __typeInfoList")
    var listIndex LnsInt
    listIndex = 1
    var outputDepend func(typeInfo *Ast_TypeInfo,moduleTypeInfo *Ast_TypeInfo) bool
    outputDepend = func(typeInfo *Ast_TypeInfo,moduleTypeInfo *Ast_TypeInfo) bool {
        {
            _moduleIndex := importModuleType2Index.Items[moduleTypeInfo]
            if _moduleIndex != nil {
                moduleIndex := _moduleIndex.(LnsInt)
                var moduleInfo *Nodes_ModuleInfo
                moduleInfo = Lns_unwrap( node.FP.Get_importModule2moduleInfo().Items[moduleTypeInfo]).(*Nodes_ModuleInfo)
                {
                    _extId := moduleInfo.FP.Get_localTypeInfo2importIdMap().Items[typeInfo]
                    if _extId != nil {
                        extId := _extId.(LnsInt)
                        self.FP.writeln(Lns_getVM().String_format("__dependIdMap[ %d ] = { %d, %d } -- %s", []LnsAny{typeInfo.FP.Get_typeId(), moduleIndex, extId, typeInfo.FP.GetTxt(nil, nil, nil)}))
                        return true
                    }
                }
            }
        }
        return false
    }
    var wroteTypeIdSet *LnsSet
    wroteTypeIdSet = NewLnsSet([]LnsAny{})
    var outputTypeInfo func(typeInfo *Ast_TypeInfo)
    outputTypeInfo = func(typeInfo *Ast_TypeInfo) {
        var force bool
        force = false
        if Ast_isExtId(typeInfo){
            var moduleTypeInfo *Ast_TypeInfo
            moduleTypeInfo = typeInfo.FP.GetModule()
            {
                _moduleInfo := node.FP.Get_importModule2moduleInfo().Items[moduleTypeInfo]
                if _moduleInfo != nil {
                    moduleInfo := _moduleInfo.(*Nodes_ModuleInfo)
                    if Lns_isCondTrue( moduleInfo.FP.Get_localTypeInfo2importIdMap().Items[typeInfo]){
                        return 
                    }
                }
            }
            if moduleTypeInfo == self.moduleTypeInfo{
                force = true
                
            }
        }
        if _switch3302 := typeInfo.FP.Get_kind(); _switch3302 == Ast_TypeInfoKind__Class || _switch3302 == Ast_TypeInfoKind__IF {
            if _switch3300 := typeInfo.FP.Get_accessMode(); _switch3300 == Ast_AccessMode__Pub || _switch3300 == Ast_AccessMode__Pro || _switch3300 == Ast_AccessMode__Global {
            } else {
                Util_errorLog(Lns_getVM().String_format("skip: %s %s", []LnsAny{typeInfo.FP.Get_accessMode(), self.FP.getFullName(typeInfo)}))
                return 
            }
        }
        var typeId LnsInt
        typeId = typeInfo.FP.Get_typeId()
        if wroteTypeIdSet.Has(typeId){
            return 
        }
        wroteTypeIdSet.Add(typeId)
        if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( force) ||
            Lns_setStackVal( checkExportTypeInfo(typeInfo)) )){
            self.FP.Write(Lns_getVM().String_format("__typeInfoList[%d] = ", []LnsAny{listIndex}))
            listIndex = listIndex + 1
            
            var validChildren LnsAny
            validChildren = validChildrenSet.Items[typeInfo]
            if Lns_op_not(validChildren){
                validChildren = typeId2TypeInfo
                
            }
            typeInfo.FP.Serialize(self.FP, validChildren)
        }
    }
    for _typeId, _typeInfo := range( self.pubEnumId2EnumTypeInfo.Items ) {
        typeId := _typeId.(LnsInt)
        typeInfo := _typeInfo.(Ast_EnumTypeInfoDownCast).ToAst_EnumTypeInfo()
        typeId2TypeInfo.Set(typeId,&typeInfo.Ast_TypeInfo)
    }
    {
        __collection3458 := self.pubAlgeId2AlgeTypeInfo
        __sorted3458 := __collection3458.CreateKeyListInt()
        __sorted3458.Sort( LnsItemKindInt, nil )
        for _, _typeId := range( __sorted3458.Items ) {
            typeInfo := __collection3458.Items[ _typeId ].(Ast_AlgeTypeInfoDownCast).ToAst_AlgeTypeInfo()
            typeId := _typeId.(LnsInt)
            typeId2TypeInfo.Set(typeId,&typeInfo.Ast_TypeInfo)
            {
                __collection3456 := typeInfo.FP.Get_valInfoMap()
                __sorted3456 := __collection3456.CreateKeyListStr()
                __sorted3456.Sort( LnsItemKindStr, nil )
                for _, ___key3456 := range( __sorted3456.Items ) {
                    valInfo := __collection3456.Items[ ___key3456 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
                    for _, _valType := range( valInfo.FP.Get_typeList().Items ) {
                        valType := _valType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        pickupTypeId(valType, true, nil)
                    }
                }
            }
        }
    }
    self.FP.writeln("local __dependIdMap = {}")
    self.FP.writeln("_moduleObj.__dependIdMap = __dependIdMap")
    var exportNeedModuleTypeInfo *LnsSet
    exportNeedModuleTypeInfo = NewLnsSet([]LnsAny{})
    {
        __collection3539 := typeId2TypeInfo
        __sorted3539 := __collection3539.CreateKeyListInt()
        __sorted3539.Sort( LnsItemKindInt, nil )
        for _, ___key3539 := range( __sorted3539.Items ) {
            typeInfo := __collection3539.Items[ ___key3539 ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var valid bool
            valid = false
            var moduleTypeInfo *Ast_TypeInfo
            moduleTypeInfo = typeInfo.FP.GetModule()
            exportNeedModuleTypeInfo.Add(Ast_TypeInfo2Stem(moduleTypeInfo))
            if outputDepend(typeInfo, moduleTypeInfo){
                valid = true
                
            }
            if Lns_op_not(valid){
                outputTypeInfo(typeInfo)
            }
        }
    }
    for _moduleTypeInfo := range( node.FP.Get_useModuleMacroSet().Items ) {
        moduleTypeInfo := _moduleTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        exportNeedModuleTypeInfo.Add(Ast_TypeInfo2Stem(moduleTypeInfo))
    }
    self.FP.writeln("local __dependModuleMap = {}")
    self.FP.writeln("_moduleObj.__dependModuleMap = __dependModuleMap")
    {
        __collection3619 := importNameMap
        __sorted3619 := __collection3619.CreateKeyListStr()
        __sorted3619.Sort( LnsItemKindStr, nil )
        for _, _name := range( __sorted3619.Items ) {
            moduleTypeInfo := __collection3619.Items[ _name ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            name := _name.(string)
            self.FP.writeln(Lns_getVM().String_format("__dependModuleMap[ '%s' ] = { typeId = %d, use = %s, buildId = %q }", []LnsAny{name, Lns_unwrap( importModuleType2Index.Items[moduleTypeInfo]).(LnsInt), exportNeedModuleTypeInfo.Has(Ast_TypeInfo2Stem(moduleTypeInfo)), (Lns_unwrap( node.FP.Get_importModule2moduleInfo().Items[moduleTypeInfo]).(*Nodes_ModuleInfo)).FP.Get_moduleId().FP.Get_idStr()}))
        }
    }
    self.FP.Write("_moduleObj.__subModuleMap = {")
    {
        var firstFlag bool
        firstFlag = true
        for _, _subfileNode := range( node.FP.Get_nodeManager().FP.GetSubfileNodeList().Items ) {
            subfileNode := _subfileNode.(Nodes_SubfileNodeDownCast).ToNodes_SubfileNode()
            {
                _usePath := subfileNode.FP.Get_usePath()
                if _usePath != nil {
                    usePath := _usePath.(string)
                    if firstFlag{
                        firstFlag = false
                        
                    } else { 
                        self.FP.Write(",")
                    }
                    self.FP.Write(Lns_getVM().String_format("%q", []LnsAny{usePath}))
                }
            }
        }
    }
    self.FP.writeln("}")
    var moduleTypeInfo *Ast_TypeInfo
    moduleTypeInfo = self.moduleTypeInfo
    var moduleSymbolKind LnsInt
    moduleSymbolKind = Ast_SymbolKind__Typ
    {
        __exp := node.FP.Get_provideNode()
        if __exp != nil {
            _exp := __exp.(*Nodes_ProvideNode)
            moduleTypeInfo = _exp.FP.Get_symbol().FP.Get_typeInfo()
            
            moduleSymbolKind = _exp.FP.Get_symbol().FP.Get_kind()
            
        }
    }
    self.FP.writeln(Lns_getVM().String_format("_moduleObj.__moduleTypeId = %d", []LnsAny{moduleTypeInfo.FP.Get_typeId()}))
    self.FP.writeln(Lns_getVM().String_format("_moduleObj.__moduleSymbolKind = %d", []LnsAny{moduleSymbolKind}))
    self.FP.writeln(Lns_getVM().String_format("_moduleObj.__moduleMutable = %s", []LnsAny{Ast_TypeInfo_isMut(moduleTypeInfo)}))
    self.FP.writeln("----- meta -----")
    if self.stream != self.metaStream{
        self.FP.writeln("return _moduleObj")
    }
    self.outMetaFlag = false
    
}

// 837: decl @lune.@base.@convLua.convFilter.processRoot
func (self *convLua_convFilter) ProcessRoot(node *Nodes_RootNode,_opt LnsAny) {
    self.FP.writeln(Lns_getVM().String_format("--%s", []LnsAny{self.streamName}))
    self.needModuleObj = Nodes_ProvideNode2Stem(node.FP.Get_provideNode()) == nil
    
    if self.needModuleObj{
        self.FP.writeln("local _moduleObj = {}")
    }
    self.FP.writeln(Lns_getVM().String_format("local __mod__ = '%s'", []LnsAny{node.FP.Get_moduleTypeInfo().FP.GetFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), nil)}))
    var luneSymbol string
    luneSymbol = Lns_getVM().String_format("_lune%d", []LnsAny{Ver_luaModVersion})
    {
        _runtime := self.useLuneRuntime
        if _runtime != nil {
            runtime := _runtime.(string)
            self.FP.writeln(Lns_getVM().String_format("local _lune = require( \"%s\" )", []LnsAny{runtime}))
        } else {
            self.FP.writeln("local _lune = {}")
            self.FP.writeln(Lns_getVM().String_format("if %s then\n   _lune = %s\nend", []LnsAny{luneSymbol, luneSymbol}))
            if node.FP.Get_luneHelperInfo().UseAlge{
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__Alge))
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__AlgeMapping))
            }
            if node.FP.Get_luneHelperInfo().UseSet{
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__SetOp))
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__SetMapping))
            }
            if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                Lns_setStackVal( node.FP.Get_luneHelperInfo().UseUnpack) &&
                Lns_setStackVal( Lns_op_not(self.targetLuaVer.FP.Get_hasTableUnpack())) ).(bool)){
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__Unpack))
            }
            if node.FP.Get_luneHelperInfo().UseLoad{
                self.FP.writeln(self.targetLuaVer.FP.GetLoadCode())
            }
            if node.FP.Get_luneHelperInfo().UseNilAccess{
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__NilAcc))
            }
            if node.FP.Get_luneHelperInfo().UseUnwrapExp{
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__Unwrap))
            }
            if node.FP.Get_luneHelperInfo().HasMappingClassDef{
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__Mapping))
            }
            if node.FP.Get_nodeManager().FP.GetImportNodeList().Len() != 0{
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__LoadModule))
            }
            if node.FP.Get_nodeManager().FP.GetExpCastNodeList().Len() != 0{
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__InstanceOf))
                self.FP.writeln(LuaMod_getCode(LuaMod_CodeKind__Cast))
            }
        }
    }
    self.FP.writeln(Lns_getVM().String_format("if not %s then\n   %s = _lune\nend", []LnsAny{luneSymbol, luneSymbol}))
    if Lns_isCondTrue( Lns_NilAccFin(Lns_NilAccPush(node.FP.Get_moduleScope().FP.GetSymbolInfoChild("_")) && 
    Lns_NilAccCall1( func () LnsAny { return Lns_NilAccPop().(*Ast_SymbolInfo).FP.Get_posForLatestMod()}))){
        self.FP.writeln("local _")
    }
    var children *LnsList
    children = node.FP.Get_children()
    for _, _child := range( children.Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        convLua_filter_1163_(child, self, &node.Nodes_Node)
        self.FP.writeln("")
    }
    self.FP.outputMeta(node)
    {
        __exp := node.FP.Get_provideNode()
        if __exp != nil {
            _exp := __exp.(*Nodes_ProvideNode)
            self.FP.Write("return ")
            self.FP.Write(_exp.FP.Get_symbol().FP.Get_name())
            self.FP.writeln("")
        } else {
            self.FP.writeln("return _moduleObj")
        }
    }
}

// 929: decl @lune.@base.@convLua.convFilter.processSubfile
func (self *convLua_convFilter) ProcessSubfile(node *Nodes_SubfileNode,_opt LnsAny) {
}

// 935: decl @lune.@base.@convLua.convFilter.processBlockSub
func (self *convLua_convFilter) ProcessBlockSub(node *Nodes_BlockNode,_opt LnsAny) {
    var word string
    word = ""
    if _switch4487 := node.FP.Get_blockKind(); _switch4487 == Nodes_BlockKind__If || _switch4487 == Nodes_BlockKind__Elseif {
        word = "then"
        
    } else if _switch4487 == Nodes_BlockKind__Else {
        word = ""
        
    } else if _switch4487 == Nodes_BlockKind__While {
        word = "do"
        
    } else if _switch4487 == Nodes_BlockKind__Repeat {
        word = ""
        
    } else if _switch4487 == Nodes_BlockKind__For {
        word = "do"
        
    } else if _switch4487 == Nodes_BlockKind__Apply {
        word = "do"
        
    } else if _switch4487 == Nodes_BlockKind__Foreach {
        word = "do"
        
    } else if _switch4487 == Nodes_BlockKind__Macro {
        word = ""
        
    } else if _switch4487 == Nodes_BlockKind__Func {
        word = ""
        
    } else if _switch4487 == Nodes_BlockKind__Default {
        word = ""
        
    } else if _switch4487 == Nodes_BlockKind__Block {
        word = "do"
        
    } else if _switch4487 == Nodes_BlockKind__LetUnwrap {
        word = ""
        
    } else if _switch4487 == Nodes_BlockKind__LetUnwrapThenDo {
        word = ""
        
    } else if _switch4487 == Nodes_BlockKind__IfUnwrap {
        word = ""
        
    }
    self.FP.writeln(word)
    self.FP.pushIndent(nil)
    if Lns_isCondTrue( Lns_NilAccFin(Lns_NilAccPush(node.FP.Get_scope().FP.GetSymbolInfoChild("_")) && 
    Lns_NilAccCall1( func () LnsAny { return Lns_NilAccPop().(*Ast_SymbolInfo).FP.Get_posForLatestMod()}))){
        self.FP.writeln("local _")
    }
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList()
    for _, _statement := range( stmtList.Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        convLua_filter_1163_(statement, self, &node.Nodes_Node)
        self.FP.writeln("")
    }
    self.FP.popIndent()
    if node.FP.Get_blockKind() == Nodes_BlockKind__Block{
        self.FP.writeln("end")
    }
}

// 1003: decl @lune.@base.@convLua.convFilter.processLoadRuntime
func (self *convLua_convFilter) processLoadRuntime() {
    {
        __exp := self.useLuneRuntime
        if __exp != nil {
            _exp := __exp.(string)
            self.FP.writeln(Lns_getVM().String_format("local _lune = require( \"%s\" )", []LnsAny{_exp}))
        } else {
            self.FP.writeln(Lns_getVM().String_format("local _lune = require( \"%s\" )", []LnsAny{Option_getRuntimeModule()}))
        }
    }
}

// 1013: decl @lune.@base.@convLua.convFilter.processScope
func (self *convLua_convFilter) ProcessScope(node *Nodes_ScopeNode,_opt LnsAny) {
    if node.FP.Get_scopeKind() == Nodes_ScopeKind__Root{
        self.FP.processLoadRuntime()
    }
    convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
}

// 1023: decl @lune.@base.@convLua.convFilter.processStmtExp
func (self *convLua_convFilter) ProcessStmtExp(node *Nodes_StmtExpNode,_opt LnsAny) {
    convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
}

// 1029: decl @lune.@base.@convLua.convFilter.processDeclEnum
func (self *convLua_convFilter) ProcessDeclEnum(node *Nodes_DeclEnumNode,_opt LnsAny) {
    var access string
    access = Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( node.FP.Get_accessMode() == Ast_AccessMode__Global) &&
        Lns_setStackVal( "") ||
        Lns_setStackVal( "local ") ).(string)
    var enumFullName string
    enumFullName = node.FP.Get_name().Txt
    var typeInfo *Ast_EnumTypeInfo
    typeInfo = Lns_unwrap( Ast_EnumTypeInfoDownCastF(node.FP.Get_expType().FP)).(*Ast_EnumTypeInfo)
    var parentInfo *Ast_TypeInfo
    parentInfo = typeInfo.FP.Get_parentInfo()
    var isTopNS bool
    isTopNS = true
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( parentInfo != Ast_headTypeInfo) &&
        Lns_setStackVal( parentInfo.FP.Get_kind() == Ast_TypeInfoKind__Class) ).(bool)){
        enumFullName = Lns_getVM().String_format("%s.%s", []LnsAny{self.FP.getFullName(parentInfo), enumFullName})
        
        access = ""
        
        isTopNS = false
        
    }
    self.FP.writeln(Lns_getVM().String_format("%s%s = {}", []LnsAny{access, enumFullName}))
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( isTopNS) &&
        Lns_setStackVal( node.FP.Get_accessMode() == Ast_AccessMode__Pub) ).(bool)){
        if self.needModuleObj{
            self.FP.writeln(Lns_getVM().String_format("_moduleObj.%s = %s", []LnsAny{enumFullName, enumFullName}))
        }
    }
    if typeInfo.FP.Get_accessMode() == Ast_AccessMode__Pub{
        self.pubEnumId2EnumTypeInfo.Set(typeInfo.FP.Get_typeId(),typeInfo)
    }
    self.FP.writeln(Lns_getVM().String_format("%s._val2NameMap = {}", []LnsAny{enumFullName}))
    self.FP.writeln(Lns_getVM().String_format("function %s:_getTxt( val )\n   local name = self._val2NameMap[ val ]\n   if name then\n      return string.format( \"%s.%%s\", name )\n   end\n   return string.format( \"illegal val -- %%s\", val )\nend\nfunction %s._from( val )\n   if %s._val2NameMap[ val ] then\n      return val\n   end\n   return nil\nend\n    ", []LnsAny{enumFullName, enumFullName, enumFullName, enumFullName}))
    self.FP.writeln(Lns_getVM().String_format("%s.__allList = {}\nfunction %s.get__allList()\n   return %s.__allList\nend\n", []LnsAny{enumFullName, enumFullName, enumFullName}))
    for _index, _valName := range( node.FP.Get_valueNameList().Items ) {
        index := _index + 1
        valName := _valName.(Parser_TokenDownCast).ToParser_Token()
        var valInfo *Ast_EnumValInfo
        valInfo = Lns_unwrap( typeInfo.FP.GetEnumValInfo(valName.Txt)).(*Ast_EnumValInfo)
        var valTxt string
        valTxt = Lns_getVM().String_format("%s", []LnsAny{Ast_getEnumLiteralVal(valInfo.FP.Get_val())})
        if typeInfo.FP.Get_valTypeInfo().FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil){
            valTxt = Lns_getVM().String_format("'%s'", []LnsAny{Ast_getEnumLiteralVal(valInfo.FP.Get_val())})
            
        }
        self.FP.writeln(Lns_getVM().String_format("%s.%s = %s", []LnsAny{enumFullName, valName.Txt, valTxt}))
        self.FP.writeln(Lns_getVM().String_format("%s._val2NameMap[%s] = '%s'", []LnsAny{enumFullName, valTxt, valName.Txt}))
        self.FP.writeln(Lns_getVM().String_format("%s.__allList[%d] = %s.%s", []LnsAny{enumFullName, index, enumFullName, valName.Txt}))
    }
}

// 1111: decl @lune.@base.@convLua.convFilter.getMapInfo
func (self *convLua_convFilter) getMapInfo(typeInfo *Ast_TypeInfo)(string, bool, string) {
    var nonnilableType *Ast_TypeInfo
    nonnilableType = typeInfo.FP.Get_srcTypeInfo()
    if typeInfo.FP.Get_nilable(){
        nonnilableType = typeInfo.FP.Get_nonnilableType()
        
    }
    var child string
    child = "{}"
    var funcTxt string
    funcTxt = ""
    if _switch5538 := nonnilableType.FP.Get_kind(); _switch5538 == Ast_TypeInfoKind__Stem {
        funcTxt = "_lune._toStem"
        
    } else if _switch5538 == Ast_TypeInfoKind__Class || _switch5538 == Ast_TypeInfoKind__IF {
        if Lns_op_not(nonnilableType.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil)){
            funcTxt = Lns_getVM().String_format("%s._fromMap", []LnsAny{self.FP.getFullName(nonnilableType)})
            
            if convLua_isGenericType_1316_(nonnilableType){
                var memStream *Util_memStream
                memStream = NewUtil_memStream()
                self.FP.outputAlter2MapFunc(memStream.FP, nonnilableType.FP.CreateAlt2typeMap(false))
                child = memStream.FP.Get_txt()
                
            }
        } else { 
            funcTxt = "_lune._toStr"
            
        }
    } else if _switch5538 == Ast_TypeInfoKind__Enum || _switch5538 == Ast_TypeInfoKind__Alge {
        funcTxt = Lns_getVM().String_format("%s._from", []LnsAny{self.FP.getFullName(nonnilableType)})
        
    } else if _switch5538 == Ast_TypeInfoKind__Prim {
        if _switch5318 := nonnilableType; _switch5318 == Ast_builtinTypeInt {
            funcTxt = "_lune._toInt"
            
        } else if _switch5318 == Ast_builtinTypeReal {
            funcTxt = "_lune._toReal"
            
        } else if _switch5318 == Ast_builtinTypeBool {
            funcTxt = "_lune._toBool"
            
        } else {
            Util_err(Lns_getVM().String_format("unknown type -- %s", []LnsAny{nonnilableType.FP.GetTxt(nil, nil, nil)}))
        }
    } else if _switch5538 == Ast_TypeInfoKind__Map {
        funcTxt = "_lune._toMap"
        
        var itemList *LnsList
        itemList = nonnilableType.FP.Get_itemTypeInfoList()
        var keyFuncTxt string
        var keyNilable bool
        var keyChild string
        keyFuncTxt,keyNilable,keyChild = self.FP.getMapInfo(itemList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        var valFuncTxt string
        var valNilable bool
        var valChild string
        valFuncTxt,valNilable,valChild = self.FP.getMapInfo(itemList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        child = Lns_getVM().String_format("{ { func = %s, nilable = %s, child = %s }, \n", []LnsAny{keyFuncTxt, keyNilable, keyChild}) + Lns_getVM().String_format("{ func = %s, nilable = %s, child = %s } }", []LnsAny{valFuncTxt, valNilable, valChild})
        
    } else if _switch5538 == Ast_TypeInfoKind__Set {
        funcTxt = "_lune._toSet"
        
        var itemList *LnsList
        itemList = nonnilableType.FP.Get_itemTypeInfoList()
        var valFuncTxt string
        var valNilable bool
        var valChild string
        valFuncTxt,valNilable,valChild = self.FP.getMapInfo(itemList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        child = Lns_getVM().String_format("{ func = %s, nilable = %s, child = %s }", []LnsAny{valFuncTxt, valNilable, valChild})
        
    } else if _switch5538 == Ast_TypeInfoKind__List || _switch5538 == Ast_TypeInfoKind__Array {
        funcTxt = "_lune._toList"
        
        var itemList *LnsList
        itemList = nonnilableType.FP.Get_itemTypeInfoList()
        var valFuncTxt string
        var valNilable bool
        var valChild string
        valFuncTxt,valNilable,valChild = self.FP.getMapInfo(itemList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        child = Lns_getVM().String_format("{ { func = %s, nilable = %s, child = %s } }", []LnsAny{valFuncTxt, valNilable, valChild})
        
    } else if _switch5538 == Ast_TypeInfoKind__Alternate {
        var prefix string
        prefix = Lns_getVM().String_format("obj.__alt2mapFunc.%s", []LnsAny{nonnilableType.FP.Get_rawTxt()})
        funcTxt = Lns_getVM().String_format("%s.func", []LnsAny{prefix})
        
        child = Lns_getVM().String_format("%s.child", []LnsAny{prefix})
        
    }
    return funcTxt, typeInfo.FP.Get_nilable(), child
}

// 1195: decl @lune.@base.@convLua.convFilter.processDeclAlge
func (self *convLua_convFilter) ProcessDeclAlge(node *Nodes_DeclAlgeNode,_opt LnsAny) {
    var access string
    access = Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( node.FP.Get_accessMode() == Ast_AccessMode__Global) &&
        Lns_setStackVal( "") ||
        Lns_setStackVal( "local ") ).(string)
    var algeFullName string
    algeFullName = node.FP.Get_algeType().FP.Get_rawTxt()
    var typeInfo *Ast_AlgeTypeInfo
    typeInfo = Lns_unwrap( Ast_AlgeTypeInfoDownCastF(node.FP.Get_expType().FP)).(*Ast_AlgeTypeInfo)
    var parentInfo *Ast_TypeInfo
    parentInfo = typeInfo.FP.Get_parentInfo()
    var isTopNS bool
    isTopNS = true
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( parentInfo != Ast_headTypeInfo) &&
        Lns_setStackVal( parentInfo.FP.Get_kind() == Ast_TypeInfoKind__Class) ).(bool)){
        algeFullName = Lns_getVM().String_format("%s.%s", []LnsAny{self.FP.getFullName(parentInfo), algeFullName})
        
        access = ""
        
        isTopNS = false
        
    }
    self.FP.writeln(Lns_getVM().String_format("%s%s = {}", []LnsAny{access, algeFullName}))
    self.FP.writeln(Lns_getVM().String_format("%s._name2Val = {}", []LnsAny{algeFullName}))
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( isTopNS) &&
        Lns_setStackVal( node.FP.Get_accessMode() == Ast_AccessMode__Pub) ).(bool)){
        if self.needModuleObj{
            self.FP.writeln(Lns_getVM().String_format("_moduleObj.%s = %s", []LnsAny{algeFullName, algeFullName}))
        }
    }
    if typeInfo.FP.Get_accessMode() == Ast_AccessMode__Pub{
        self.pubAlgeId2AlgeTypeInfo.Set(typeInfo.FP.Get_typeId(),typeInfo)
    }
    self.FP.writeln(Lns_getVM().String_format("function %s:_getTxt( val )\n   local name = val[ 1 ]\n   if name then\n      return string.format( \"%s.%%s\", name )\n   end\n   return string.format( \"illegal val -- %%s\", val )\nend\n", []LnsAny{algeFullName, algeFullName}))
    self.FP.writeln(Lns_getVM().String_format("function %s._from( val )\n   return _lune._AlgeFrom( %s, val )\nend\n", []LnsAny{algeFullName, algeFullName}))
    {
        __collection5896 := node.FP.Get_algeType().FP.Get_valInfoMap()
        __sorted5896 := __collection5896.CreateKeyListStr()
        __sorted5896.Sort( LnsItemKindStr, nil )
        for _, ___key5896 := range( __sorted5896.Items ) {
            valInfo := __collection5896.Items[ ___key5896 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            self.FP.Write(Lns_getVM().String_format("%s.%s = { \"%s\"", []LnsAny{algeFullName, valInfo.FP.Get_name(), valInfo.FP.Get_name()}))
            if valInfo.FP.Get_typeList().Len() > 0{
                self.FP.Write(", {")
                for _index, _paramType := range( valInfo.FP.Get_typeList().Items ) {
                    index := _index + 1
                    paramType := _paramType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if index > 1{
                        self.FP.Write(",")
                    }
                    var funcTxt string
                    var nilable bool
                    var child string
                    funcTxt,nilable,child = self.FP.getMapInfo(paramType)
                    self.FP.Write(Lns_getVM().String_format("{ func=%s, nilable=%s, child=%s }", []LnsAny{funcTxt, nilable, child}))
                }
                self.FP.Write("}")
            }
            self.FP.writeln("}")
            self.FP.writeln(Lns_getVM().String_format("%s._name2Val[\"%s\"] = %s.%s", []LnsAny{algeFullName, valInfo.FP.Get_name(), algeFullName, valInfo.FP.Get_name()}))
        }
    }
}

// 1256: decl @lune.@base.@convLua.convFilter.processNewAlgeVal
func (self *convLua_convFilter) ProcessNewAlgeVal(node *Nodes_NewAlgeValNode,_opt LnsAny) {
    var valInfo *Ast_AlgeValInfo
    valInfo = node.FP.Get_valInfo()
    self.FP.Write(Lns_getVM().String_format("_lune.newAlge( %s.%s", []LnsAny{self.FP.getFullName(&node.FP.Get_algeTypeInfo().Ast_TypeInfo), valInfo.FP.Get_name()}))
    if valInfo.FP.Get_typeList().Len() > 0{
        self.FP.Write(", {")
        for _index, _exp := range( node.FP.Get_paramList().Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            if index > 1{
                self.FP.Write(",")
            }
            convLua_filter_1163_(exp, self, &node.Nodes_Node)
        }
        self.FP.Write("}")
    }
    self.FP.Write(")")
}

// 1277: decl @lune.@base.@convLua.convFilter.getDestrClass
func (self *convLua_convFilter) getDestrClass(classTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo *Ast_TypeInfo
    typeInfo = classTypeInfo
    for Lns_op_not(typeInfo.FP.Equals(self.processInfo, Ast_headTypeInfo, nil, nil)) {
        var scope *Ast_Scope
        scope = Lns_unwrap( typeInfo.FP.Get_scope()).(*Ast_Scope)
        if Lns_isCondTrue( scope.FP.GetTypeInfoChild("__free")){
            return typeInfo
        }
        typeInfo = typeInfo.FP.Get_baseTypeInfo()
        
    }
    return nil
}

// 1290: decl @lune.@base.@convLua.convFilter.outputAlter2MapFunc
func (self *convLua_convFilter) outputAlter2MapFunc(stream Lns_oStream,alt2Map *LnsMap) {
    stream.Write("{")
    for _altType, _assinType := range( alt2Map.Items ) {
        altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        assinType := _assinType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if altType.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
            if assinType.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
                stream.Write(Lns_getVM().String_format("%s = obj.__alt2mapFunc.%s,", []LnsAny{altType.FP.Get_rawTxt(), assinType.FP.Get_rawTxt()}))
            } else { 
                var funcTxt string
                var nilable bool
                var child string
                funcTxt,nilable,child = self.FP.getMapInfo(assinType)
                stream.Write(Lns_getVM().String_format("%s = { func=%s, nilable=%s, child=%s },", []LnsAny{altType.FP.Get_rawTxt(), funcTxt, nilable, child}))
            }
        }
    }
    stream.Write("}")
}

// 1315: decl @lune.@base.@convLua.convFilter.processProtoClass
func (self *convLua_convFilter) ProcessProtoClass(node *Nodes_ProtoClassNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("local %s = {}", []LnsAny{node.FP.Get_name().Txt}))
}

// 1321: decl @lune.@base.@convLua.convFilter.processDeclClass
func (self *convLua_convFilter) ProcessDeclClass(node *Nodes_DeclClassNode,_opt LnsAny) {
    var nodeInfo *Nodes_DeclClassNode
    nodeInfo = node
    var classNameToken *Parser_Token
    classNameToken = nodeInfo.FP.Get_name()
    var className string
    className = classNameToken.Txt
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = node.FP.Get_expType()
    var classTypeId LnsInt
    classTypeId = classTypeInfo.FP.Get_typeId()
    if nodeInfo.FP.Get_accessMode() == Ast_AccessMode__Pub{
        self.classId2TypeInfo.Set(classTypeId,classTypeInfo)
    }
    self.classId2MemberList.Set(classTypeId,nodeInfo.FP.Get_memberList())
    {
        __exp := node.FP.Get_moduleName()
        if __exp != nil {
            _exp := __exp.(*Parser_Token)
            self.FP.Write(Lns_getVM().String_format("local %s = require( %s )", []LnsAny{className, _exp.Txt}))
            if _switch6335 := node.FP.Get_accessMode(); _switch6335 == Ast_AccessMode__Pub || _switch6335 == Ast_AccessMode__Pro {
                if self.needModuleObj{
                    self.FP.writeln("")
                    self.FP.Write(Lns_getVM().String_format("_moduleObj.%s = %s", []LnsAny{className, className}))
                }
            }
            return 
        }
    }
    if Lns_op_not(node.FP.Get_hasPrototype()){
        self.FP.writeln(Lns_getVM().String_format("local %s = {}", []LnsAny{className}))
    }
    var ifTxt string
    ifTxt = ""
    if classTypeInfo.FP.Get_interfaceList().Len() > 0{
        for _, _ifType := range( classTypeInfo.FP.Get_interfaceList().Items ) {
            ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            ifTxt = ifTxt + self.FP.getFullName(ifType) + ","
            
        }
        ifTxt = Lns_getVM().String_format("ifList = {%s}", []LnsAny{ifTxt})
        
    }
    var baseInfo *Ast_TypeInfo
    baseInfo = classTypeInfo.FP.Get_baseTypeInfo()
    var baseTxt string
    baseTxt = ""
    if baseInfo.FP.Get_typeId() != Ast_rootTypeId{
        baseTxt = Lns_getVM().String_format("__index = %s", []LnsAny{self.FP.getFullName(baseInfo)})
        
    }
    if Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( len(ifTxt) > 0) ||
        Lns_setStackVal( len(baseTxt) > 0) ).(bool){
        var metaTxt string
        metaTxt = baseTxt
        if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( len(baseTxt) > 0) &&
            Lns_setStackVal( len(ifTxt) > 0) ).(bool)){
            metaTxt = Lns_getVM().String_format("%s,%s", []LnsAny{baseTxt, ifTxt})
            
        } else if len(ifTxt) > 0{
            metaTxt = ifTxt
            
        }
        self.FP.writeln(Lns_getVM().String_format("setmetatable( %s, { %s } )", []LnsAny{className, metaTxt}))
    }
    if nodeInfo.FP.Get_accessMode() == Ast_AccessMode__Pub{
        if self.needModuleObj{
            self.FP.writeln(Lns_getVM().String_format("_moduleObj.%s = %s", []LnsAny{className, className}))
        }
    }
    for _, _declNode := range( node.FP.Get_declStmtList().Items ) {
        declNode := _declNode.(Nodes_NodeDownCast).ToNodes_Node()
        convLua_filter_1163_(declNode, self, &node.Nodes_Node)
    }
    var hasConstrFlag bool
    hasConstrFlag = false
    var memberList *LnsList
    memberList = NewLnsList([]LnsAny{})
    var fieldList *LnsList
    fieldList = nodeInfo.FP.Get_fieldList()
    var outerMethodSet *LnsSet
    outerMethodSet = nodeInfo.FP.Get_outerMethodSet()
    var methodNameSet *LnsSet
    methodNameSet = NewLnsSet([]LnsAny{})
    if classTypeInfo.FP.Get_kind() != Ast_TypeInfoKind__IF{
        for _, _field := range( fieldList.Items ) {
            field := _field.(Nodes_NodeDownCast).ToNodes_Node()
            var ignoreFlag bool
            ignoreFlag = false
            if field.FP.Get_kind() == Nodes_NodeKind_get_DeclConstr(){
                hasConstrFlag = true
                
                methodNameSet.Add("__init")
            }
            if field.FP.Get_kind() == Nodes_NodeKind_get_DeclDestr(){
                methodNameSet.Add("__free")
            }
            {
                _declMemberNode := Nodes_DeclMemberNodeDownCastF(field.FP)
                if _declMemberNode != nil {
                    declMemberNode := _declMemberNode.(*Nodes_DeclMemberNode)
                    if Lns_op_not(declMemberNode.FP.Get_staticFlag()){
                        memberList.Insert(Nodes_DeclMemberNode2Stem(declMemberNode))
                    }
                }
            }
            {
                _methodNode := Nodes_DeclMethodNodeDownCastF(field.FP)
                if _methodNode != nil {
                    methodNode := _methodNode.(*Nodes_DeclMethodNode)
                    var declInfo *Nodes_DeclFuncInfo
                    declInfo = methodNode.FP.Get_declInfo()
                    var methodNameToken *Parser_Token
                    methodNameToken = Lns_unwrap( declInfo.FP.Get_name()).(*Parser_Token)
                    if outerMethodSet.Has(methodNameToken.Txt){
                        ignoreFlag = true
                        
                    }
                    methodNameSet.Add(methodNameToken.Txt)
                }
            }
            if (Lns_op_not(ignoreFlag)){
                convLua_filter_1163_(field, self, &node.Nodes_Node)
            }
        }
    }
    var destTxt string
    destTxt = ""
    {
        __exp := self.FP.getDestrClass(node.FP.Get_expType())
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            destTxt = Lns_getVM().String_format(", __gc = %s.__free", []LnsAny{_exp.FP.GetTxt(nil, nil, nil)})
            
        }
    }
    self.FP.writeln(Lns_getVM().String_format("function %s.setmeta( obj )\n  setmetatable( obj, { __index = %s %s } )\nend", []LnsAny{className, className, destTxt}))
    if Lns_op_not(hasConstrFlag){
        methodNameSet.Add("__init")
        var oldFlag bool
        oldFlag = node.FP.Get_hasOldCtor()
        var superArgTxt string
        superArgTxt = ""
        var thisArgTxt string
        thisArgTxt = ""
        if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( Lns_op_not(oldFlag)) &&
            Lns_setStackVal( baseInfo != Ast_headTypeInfo) ).(bool)){
            {
                _superInit := (Lns_unwrap( baseInfo.FP.Get_scope()).(*Ast_Scope)).FP.GetSymbolInfoChild("__init")
                if _superInit != nil {
                    superInit := _superInit.(*Ast_SymbolInfo)
                    for _index, _ := range( superInit.FP.Get_typeInfo().FP.Get_argTypeInfoList().Items ) {
                        index := _index + 1
                        if len(superArgTxt) > 0{
                            superArgTxt = superArgTxt + ", "
                            
                        }
                        superArgTxt = Lns_getVM().String_format("%s__superarg%d", []LnsAny{superArgTxt, index})
                        
                    }
                }
            }
        }
        for _, _member := range( memberList.Items ) {
            member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if len(thisArgTxt) > 0{
                thisArgTxt = thisArgTxt + ", "
                
            }
            thisArgTxt = thisArgTxt + member.FP.Get_name().Txt
            
        }
        var argTxt string
        argTxt = superArgTxt
        if thisArgTxt != ""{
            if len(argTxt) > 0{
                argTxt = argTxt + ","
                
            }
            argTxt = argTxt + thisArgTxt
            
        }
        self.FP.writeln(Lns_getVM().String_format("function %s.new( %s )\n   local obj = {}\n   %s.setmeta( obj )\n   if obj.__init then\n      obj:__init( %s )\n   end\n   return obj\nend\nfunction %s:__init( %s )\n", []LnsAny{className, argTxt, className, argTxt, className, argTxt}))
        self.FP.pushIndent(nil)
        if baseInfo != Ast_headTypeInfo{
            if Lns_isCondTrue( (Lns_unwrap( baseInfo.FP.Get_scope()).(*Ast_Scope)).FP.GetSymbolInfoChild("__init")){
                self.FP.Write(Lns_getVM().String_format("%s.__init( self", []LnsAny{self.FP.getFullName(baseInfo)}))
                if len(superArgTxt) > 0{
                    self.FP.writeln(Lns_getVM().String_format(", %s )", []LnsAny{superArgTxt}))
                } else { 
                    self.FP.writeln(")")
                }
            }
        }
        for _, _member := range( memberList.Items ) {
            member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            var memberName string
            memberName = member.FP.Get_name().Txt
            self.FP.writeln(Lns_getVM().String_format("self.%s = %s", []LnsAny{memberName, memberName}))
        }
        self.FP.popIndent()
        self.FP.writeln("end")
    }
    for _, _memberNode := range( nodeInfo.FP.Get_memberList().Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var memberNameToken *Parser_Token
        memberNameToken = memberNode.FP.Get_name()
        var memberName string
        memberName = memberNameToken.Txt
        var getterName string
        getterName = "get_" + memberName
        var autoFlag bool
        autoFlag = Lns_op_not(methodNameSet.Has(getterName))
        var prefix string
        var delimit string
        if memberNode.FP.Get_staticFlag(){
            prefix = className
            
            delimit = "."
            
        } else { 
            prefix = "self"
            
            delimit = ":"
            
        }
        if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( memberNode.FP.Get_getterMode() != Ast_AccessMode__None) &&
            Lns_setStackVal( autoFlag) ).(bool)){
            self.FP.writeln(Lns_getVM().String_format("function %s%s%s()\n   return %s.%s\nend", []LnsAny{className, delimit, getterName, prefix, memberName}))
            methodNameSet.Add(getterName)
        }
        var setterName string
        setterName = "set_" + memberName
        autoFlag = Lns_op_not(methodNameSet.Has(setterName))
        
        if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( memberNode.FP.Get_setterMode() != Ast_AccessMode__None) &&
            Lns_setStackVal( autoFlag) ).(bool)){
            self.FP.writeln(Lns_getVM().String_format("function %s%s%s( %s )\n   %s.%s = %s\nend", []LnsAny{className, delimit, setterName, memberName, prefix, memberName, memberName}))
            methodNameSet.Add(setterName)
        }
    }
    for _, _advertiseInfo := range( node.FP.Get_advertiseList().Items ) {
        advertiseInfo := _advertiseInfo.(Nodes_AdvertiseInfoDownCast).ToNodes_AdvertiseInfo()
        var memberName string
        memberName = advertiseInfo.FP.Get_member().FP.Get_name().Txt
        var memberType *Ast_TypeInfo
        memberType = advertiseInfo.FP.Get_member().FP.Get_expType()
        for _, _mtdName := range( Ast_getAllMethodName(memberType, Ast_MethodKind__Object).FP.Get_list().Items ) {
            mtdName := _mtdName.(string)
            var mbrScope *Ast_Scope
            mbrScope = Lns_unwrap( memberType.FP.Get_scope()).(*Ast_Scope)
            var child *Ast_TypeInfo
            child = Lns_unwrap( mbrScope.FP.GetTypeInfoField(mtdName, true, mbrScope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
            if child.FP.Get_accessMode() != Ast_AccessMode__Pri{
                var childName string
                childName = advertiseInfo.FP.Get_prefix() + child.FP.GetTxt(nil, nil, nil)
                if Lns_op_not(methodNameSet.Has(childName)){
                    self.FP.writeln(Lns_getVM().String_format("function %s:%s( ... )\n   return self.%s:%s( ... )\nend\n", []LnsAny{className, childName, memberName, childName}))
                }
            }
        }
    }
    {
        _initBlock := Lns_NilAccFin(Lns_NilAccPush(nodeInfo.FP.Get_initBlock().FP.Get_func()) && 
        Lns_NilAccCall1( func () LnsAny { return Lns_NilAccPop().(*Nodes_DeclMethodNode).FP.Get_declInfo()})&&
        Lns_NilAccCall1( func () LnsAny { return Lns_NilAccPop().(*Nodes_DeclFuncInfo).FP.Get_body()}))
        if _initBlock != nil {
            initBlock := _initBlock.(*Nodes_BlockNode)
            if initBlock.FP.Get_stmtList().Len() > 0{
                self.FP.writeln("do")
                self.FP.pushIndent(nil)
                for _, _initStmt := range( initBlock.FP.Get_stmtList().Items ) {
                    initStmt := _initStmt.(Nodes_NodeDownCast).ToNodes_Node()
                    convLua_filter_1163_(initStmt, self, &node.Nodes_Node)
                    self.FP.writeln("")
                }
                self.FP.popIndent()
                self.FP.writeln("end")
            }
        }
    }
    if classTypeInfo.FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, nil){
        var declArgTxt string
        declArgTxt = "val"
        var argTxt string
        argTxt = "{}, val"
        if convLua_isGenericType_1316_(classTypeInfo){
            declArgTxt = "val, __alt2mapFunc"
            
            argTxt = "{ __alt2mapFunc = __alt2mapFunc }, val"
            
        }
        self.FP.writeln(Lns_getVM().String_format("function %s:_toMap()\n  return self\nend\nfunction %s._fromMap( %s )\n  local obj, mes = %s._fromMapSub( %s )\n  if obj then\n     %s.setmeta( obj )\n  end\n  return obj, mes\nend\nfunction %s._fromStem( %s )\n  return %s._fromMap( %s )\nend\n", []LnsAny{className, className, declArgTxt, className, argTxt, className, className, declArgTxt, className, declArgTxt}))
        self.FP.writeln(Lns_getVM().String_format("function %s._fromMapSub( obj, val )", []LnsAny{className}))
        if classTypeInfo.FP.Get_baseTypeInfo() != Ast_headTypeInfo{
            self.FP.writeln(Lns_getVM().String_format("   local result, mes = %s._fromMapSub( obj, val )\n   if not result then\n      return nil, mes\n   end\n", []LnsAny{self.FP.getFullName(classTypeInfo.FP.Get_baseTypeInfo())}))
        }
        self.FP.writeln("   local memInfo = {}")
        for _, _memberNode := range( node.FP.Get_memberList().Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            var funcTxt string
            var nilable bool
            var child string
            funcTxt,nilable,child = self.FP.getMapInfo(memberNode.FP.Get_expType())
            self.FP.writeln(Lns_getVM().String_format("   table.insert( memInfo, { name = \"%s\", func = %s, nilable = %s, child = %s } )", []LnsAny{memberNode.FP.Get_name().Txt, funcTxt, nilable, child}))
        }
        self.FP.writeln("   local result, mess = _lune._fromMap( obj, val, memInfo )\n   if not result then\n      return nil, mess\n   end\n   return obj\nend")
    }
}

// 1645: decl @lune.@base.@convLua.convFilter.processDeclMember
func (self *convLua_convFilter) ProcessDeclMember(node *Nodes_DeclMemberNode,_opt LnsAny) {
}

// 1652: decl @lune.@base.@convLua.convFilter.processExpMacroExp
func (self *convLua_convFilter) ProcessExpMacroExp(node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    for _, _stmt := range( node.FP.Get_stmtList().Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        convLua_filter_1163_(stmt, self, &node.Nodes_Node)
        self.FP.writeln("")
    }
}

// 1664: decl @lune.@base.@convLua.convFilter.outputDeclMacro
func (self *convLua_convFilter) OutputDeclMacro(name string,argNameList *LnsList,callback convLua_outputMacroStmtBlock_1378_) {
    self.FP.Write(Lns_getVM().String_format("local function %s(", []LnsAny{name}))
    self.FP.writeln("__macroArgs )")
    self.FP.pushIndent(nil)
    self.FP.writeln(Lns_getVM().String_format("local _lune = require( \"%s\" )", []LnsAny{Option_getRuntimeModule()}))
    self.FP.writeln("local __var = __macroArgs.__var")
    for _, _argName := range( argNameList.Items ) {
        argName := _argName.(string)
        self.FP.writeln(Lns_getVM().String_format("local %s = __macroArgs.%s", []LnsAny{argName, argName}))
    }
    self.FP.writeln("local macroVar = {}")
    self.FP.writeln("macroVar.__names = {}")
    self.FP.writeln("local function __expStatList( list )\n  local ret = \"\"\n  for index, txt in ipairs( list ) do\n    ret = string.format( \"%s %s \", ret, txt )\n  end\n  return ret\nend\n")
    self.macroDepth = self.macroDepth + 1
    
    callback()
    self.macroDepth = self.macroDepth - 1
    
    self.FP.writeln("")
    self.FP.writeln("macroVar.__var = __var")
    self.FP.writeln("return macroVar")
    self.FP.popIndent()
    self.FP.writeln("end")
    self.FP.writeln(Lns_getVM().String_format("return %s", []LnsAny{name}))
}

// 1708: decl @lune.@base.@convLua.convFilter.processExpMacroStatList
func (self *convLua_convFilter) ProcessExpMacroStatList(node *Nodes_ExpMacroStatListNode,_opt LnsAny) {
    self.FP.Write("__expStatList(")
    convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Write(")")
}

// 1716: decl @lune.@base.@convLua.convFilter.processDeclMacro
func (self *convLua_convFilter) ProcessDeclMacro(node *Nodes_DeclMacroNode,_opt LnsAny) {
    if self.inMacro{
        var macroInfo *Nodes_DeclMacroInfo
        macroInfo = node.FP.Get_declInfo()
        var argNameList *LnsList
        argNameList = NewLnsList([]LnsAny{})
        for _, _arg := range( macroInfo.FP.Get_argList().Items ) {
            arg := _arg.(Nodes_DeclArgNodeDownCast).ToNodes_DeclArgNode()
            argNameList.Insert(arg.FP.Get_name().Txt)
        }
        self.FP.OutputDeclMacro(macroInfo.FP.Get_name().Txt, argNameList, convLua_outputMacroStmtBlock_1378_(func() {
            {
                _stmtBlock := macroInfo.FP.Get_stmtBlock()
                if _stmtBlock != nil {
                    stmtBlock := _stmtBlock.(*Nodes_BlockNode)
                    convLua_filter_1163_(&stmtBlock.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }))
    }
}

// 1735: decl @lune.@base.@convLua.convFilter.processExpMacroStat
func (self *convLua_convFilter) ProcessExpMacroStat(node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    if node.FP.Get_expStrList().Len() == 0{
        self.FP.Write("''")
    } else { 
        for _index, _token := range( node.FP.Get_expStrList().Items ) {
            index := _index + 1
            token := _token.(Nodes_NodeDownCast).ToNodes_Node()
            if index != 1{
                self.FP.Write("..")
            }
            convLua_filter_1163_(token, self, &node.Nodes_Node)
        }
    }
}

// 1753: decl @lune.@base.@convLua.convFilter.processExpNew
func (self *convLua_convFilter) ProcessExpNew(node *Nodes_ExpNewNode,_opt LnsAny) {
    convLua_filter_1163_(node.FP.Get_symbol(), self, &node.Nodes_Node)
    self.FP.Write(".new(")
    {
        __exp := node.FP.Get_argList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write(")")
}

// 1764: decl @lune.@base.@convLua.convFilter.process__func__symbol
func (self *convLua_convFilter) process__func__symbol(has__func__Symbol bool,parentType *Ast_TypeInfo,funcName string) {
    if has__func__Symbol{
        var nameSpace string
        nameSpace = self.FP.getCanonicalName(parentType, false)
        if funcName == ""{
            funcName = "<anonymous>"
            
        }
        self.FP.pushIndent(nil)
        self.FP.writeln(Lns_getVM().String_format("local __func__ = '%s.%s'", []LnsAny{nameSpace, funcName}))
        self.FP.popIndent()
    }
}

// 1778: decl @lune.@base.@convLua.convFilter.processDeclConstr
func (self *convLua_convFilter) ProcessDeclConstr(node *Nodes_DeclConstrNode,_opt LnsAny) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo()
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = Lns_unwrap( declInfo.FP.Get_classTypeInfo()).(*Ast_TypeInfo)
    var className string
    className = self.FP.getFullName(classTypeInfo)
    self.FP.Write(Lns_getVM().String_format("function %s.new( ", []LnsAny{className}))
    var argTxt string
    argTxt = ""
    self.FP.Write(argTxt)
    var argList *LnsList
    argList = declInfo.FP.Get_argList()
    for _, _arg := range( argList.Items ) {
        arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
        if len(argTxt) > 0{
            self.FP.Write(", ")
            argTxt = argTxt + ", "
            
        }
        convLua_filter_1163_(arg, self, &node.Nodes_Node)
        {
            __exp := Nodes_DeclArgNodeDownCastF(arg.FP)
            if __exp != nil {
                _exp := __exp.(*Nodes_DeclArgNode)
                argTxt = argTxt + _exp.FP.Get_name().Txt
                
            } else {
                var name *Parser_Token
                name = Lns_unwrap( node.FP.Get_declInfo().FP.Get_name()).(*Parser_Token)
                Util_err(Lns_getVM().String_format("not support ... in macro -- %s", []LnsAny{name.Txt}))
            }
        }
    }
    self.FP.writeln(" )")
    self.FP.pushIndent(nil)
    self.FP.writeln("local obj = {}")
    self.FP.writeln(Lns_getVM().String_format("%s.setmeta( obj )", []LnsAny{className}))
    self.FP.writeln(Lns_getVM().String_format("if obj.__init then obj:__init( %s ); end", []LnsAny{argTxt}))
    self.FP.writeln("return obj")
    self.FP.popIndent()
    self.FP.writeln("end")
    self.FP.Write(Lns_getVM().String_format("function %s:__init(%s) ", []LnsAny{className, argTxt}))
    {
        __exp := declInfo.FP.Get_body()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.process__func__symbol(declInfo.FP.Get_has__func__Symbol(), node.FP.Get_expType().FP.Get_parentInfo(), "__init")
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln("end")
}

// 1833: decl @lune.@base.@convLua.convFilter.processDeclDestr
func (self *convLua_convFilter) ProcessDeclDestr(node *Nodes_DeclDestrNode,_opt LnsAny) {
    self.FP.writeln(Lns_getVM().String_format("function %s.__free( self )", []LnsAny{Lns_NilAccFin(Lns_NilAccPush(node.FP.Get_declInfo().FP.Get_classTypeInfo()) && 
    Lns_NilAccCall1( func () LnsAny { return Lns_NilAccPop().(*Ast_TypeInfo).FP.GetTxt(nil, nil, nil)})/* 1837:9 */)}))
    self.FP.process__func__symbol(node.FP.Get_declInfo().FP.Get_has__func__Symbol(), node.FP.Get_expType().FP.Get_parentInfo(), "__free")
    convLua_filter_1163_(&Lns_unwrap( node.FP.Get_declInfo().FP.Get_body()).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = node.FP.Get_expType().FP.Get_parentInfo()
    {
        __exp := self.FP.getDestrClass(classTypeInfo.FP.Get_baseTypeInfo())
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            self.FP.writeln(Lns_getVM().String_format("%s.__free( self )", []LnsAny{_exp.FP.GetTxt(nil, nil, nil)}))
        }
    }
    self.FP.writeln("end")
}

// 1853: decl @lune.@base.@convLua.convFilter.processExpCallSuperCtor
func (self *convLua_convFilter) ProcessExpCallSuperCtor(node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType()
    self.FP.Write(Lns_getVM().String_format("%s.%s( self", []LnsAny{self.FP.getFullName(typeInfo), node.FP.Get_methodType().FP.Get_rawTxt()}))
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(",")
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln(")")
}

// 1867: decl @lune.@base.@convLua.convFilter.processExpCallSuper
func (self *convLua_convFilter) ProcessExpCallSuper(node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = node.FP.Get_superType()
    self.FP.Write(Lns_getVM().String_format("%s.%s( self", []LnsAny{self.FP.getFullName(typeInfo), node.FP.Get_methodType().FP.Get_rawTxt()}))
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(",")
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write(")")
}

// 1882: decl @lune.@base.@convLua.convFilter.processDeclMethod
func (self *convLua_convFilter) ProcessDeclMethod(node *Nodes_DeclMethodNode,_opt LnsAny) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo()
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = Lns_unwrap( declInfo.FP.Get_classTypeInfo()).(*Ast_TypeInfo)
    var delimit string
    delimit = ":"
    if declInfo.FP.Get_staticFlag(){
        delimit = "."
        
    }
    var methodNodeToken *Parser_Token
    methodNodeToken = Lns_unwrap( declInfo.FP.Get_name()).(*Parser_Token)
    var methodName string
    methodName = methodNodeToken.Txt
    self.FP.Write(Lns_getVM().String_format("function %s%s%s( ", []LnsAny{classTypeInfo.FP.Get_rawTxt(), delimit, methodName}))
    var argList *LnsList
    argList = declInfo.FP.Get_argList()
    for _index, _arg := range( argList.Items ) {
        index := _index + 1
        arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
        if index > 1{
            self.FP.Write(", ")
        }
        convLua_filter_1163_(arg, self, &node.Nodes_Node)
    }
    self.FP.writeln(" )")
    {
        __exp := declInfo.FP.Get_body()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.process__func__symbol(declInfo.FP.Get_has__func__Symbol(), node.FP.Get_expType().FP.Get_parentInfo(), methodName)
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln("end")
}

// 1915: decl @lune.@base.@convLua.convFilter.processUnwrapSet
func (self *convLua_convFilter) ProcessUnwrapSet(node *Nodes_UnwrapSetNode,_opt LnsAny) {
    var dstExpList *Nodes_ExpListNode
    dstExpList = node.FP.Get_dstExpList()
    convLua_filter_1163_(&dstExpList.Nodes_Node, self, &node.Nodes_Node)
    self.FP.Write(" = ")
    convLua_filter_1163_(&node.FP.Get_srcExpList().Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln("")
    self.FP.Write("if ")
    for _index, _expNode := range( dstExpList.FP.Get_expList().Items ) {
        index := _index + 1
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        if index > 1{
            self.FP.Write(" or ")
        }
        self.FP.Write("nil == ")
        convLua_filter_1163_(expNode, self, &node.Nodes_Node)
    }
    self.FP.writeln(" then")
    self.FP.pushIndent(nil)
    for _index, _expNode := range( dstExpList.FP.Get_expList().Items ) {
        index := _index + 1
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.Write(Lns_getVM().String_format("local _exp%d = ", []LnsAny{index}))
        convLua_filter_1163_(expNode, self, &node.Nodes_Node)
        self.FP.writeln("")
    }
    if Lns_isCondTrue( node.FP.Get_unwrapBlock()){
        convLua_filter_1163_(&Lns_unwrap( node.FP.Get_unwrapBlock()).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
    }
    self.FP.popIndent()
    self.FP.writeln("end")
}

// 1948: decl @lune.@base.@convLua.convFilter.processExpListSub
func (self *convLua_convFilter) processExpListSub(parent *Nodes_Node,expList *LnsList,mRetExp LnsAny) {
    var mRetIndex LnsAny
    mRetIndex = Lns_NilAccFin(Lns_NilAccPush(mRetExp) && 
    Lns_NilAccCall1( func () LnsAny { return Lns_NilAccPop().(*Nodes_MRetExp).FP.Get_index()}))
    for _index, _exp := range( expList.Items ) {
        index := _index + 1
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        if exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr{
            break
        }
        {
            _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
            if _castNode != nil {
                castNode := _castNode.(*Nodes_ExpCastNode)
                if castNode.FP.Get_castKind() == Nodes_CastKind__Implicit{
                    if castNode.FP.Get_exp().FP.Get_kind() == Nodes_NodeKind_get_ExpAccessMRet(){
                        break
                    }
                }
            }
        }
        if index > 1{
            self.FP.Write(", ")
        }
        convLua_filter_1163_(exp, self, parent)
        if index == mRetIndex{
            break
        }
    }
}

// 1975: decl @lune.@base.@convLua.convFilter.processExpMRet
func (self *convLua_convFilter) ProcessExpMRet(node *Nodes_ExpMRetNode,_opt LnsAny) {
    convLua_filter_1163_(node.FP.Get_mRet(), self, &node.Nodes_Node)
}

// 1981: decl @lune.@base.@convLua.convFilter.processIfUnwrap
func (self *convLua_convFilter) ProcessIfUnwrap(node *Nodes_IfUnwrapNode,_opt LnsAny) {
    self.FP.writeln("do")
    self.FP.pushIndent(nil)
    self.FP.Write("local ")
    for _index, _varSym := range( node.FP.Get_varSymList().Items ) {
        index := _index + 1
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        self.FP.Write(convLua_getSymTxt_1068_(varSym.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{varSym.FP.Get_symbolId()})))
        if index != node.FP.Get_varSymList().Len(){
            self.FP.Write(", ")
        }
    }
    self.FP.Write(" = ")
    self.FP.processExpListSub(&node.Nodes_Node, node.FP.Get_expList().FP.Get_expList(), node.FP.Get_expList().FP.Get_mRetExp())
    self.FP.writeln("")
    self.FP.Write("if ")
    var hasSym bool
    hasSym = false
    for _, _varSym := range( node.FP.Get_varSymList().Items ) {
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if varSym.FP.Get_name() != "_"{
            if hasSym{
                self.FP.Write(" and  ")
            }
            self.FP.Write(Lns_getVM().String_format("%s ~= nil", []LnsAny{convLua_getSymTxt_1068_(varSym.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{varSym.FP.Get_symbolId()}))}))
            hasSym = true
            
        }
    }
    self.FP.Write(" then")
    convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    {
        __exp := node.FP.Get_nilBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write("else")
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln("end")
    self.FP.popIndent()
    self.FP.writeln("end")
}

// 2024: decl @lune.@base.@convLua.convFilter.processWhen
func (self *convLua_convFilter) ProcessWhen(node *Nodes_WhenNode,_opt LnsAny) {
    self.FP.Write("if ")
    for _index, _symPair := range( node.FP.Get_symPairList().Items ) {
        index := _index + 1
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        self.FP.Write(Lns_getVM().String_format("%s ~= nil", []LnsAny{symPair.FP.Get_src().FP.Get_name()}))
        if index != node.FP.Get_symPairList().Len(){
            self.FP.Write(" and ")
        }
    }
    self.FP.Write(" then")
    convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    {
        __exp := node.FP.Get_elseBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write("else")
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.writeln("end")
}

// 2046: decl @lune.@base.@convLua.convFilter.processDeclVar
func (self *convLua_convFilter) ProcessDeclVar(node *Nodes_DeclVarNode,_opt LnsAny) {
    if Lns_isCondTrue( node.FP.Get_syncBlock()){
        self.FP.writeln("do")
        self.FP.pushIndent(nil)
        for _, _varInfo := range( node.FP.Get_syncVarList().Items ) {
            varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            self.FP.writeln(Lns_getVM().String_format("local _sync_%s", []LnsAny{convLua_getSymTxt_1068_(varInfo.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{varInfo.FP.Get_symbolId()}))}))
        }
        self.FP.writeln("do")
        self.FP.pushIndent(nil)
    }
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( node.FP.Get_mode() != Nodes_DeclVarMode__Unwrap) &&
        Lns_setStackVal( node.FP.Get_accessMode() != Ast_AccessMode__Global) ).(bool)){
        self.FP.Write("local ")
    }
    var varList *LnsList
    varList = node.FP.Get_symbolInfoList()
    var varNameList *LnsList
    varNameList = NewLnsList([]LnsAny{})
    for _index, __var := range( varList.Items ) {
        index := _index + 1
        _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(", ")
        }
        var name string
        name = convLua_getSymTxt_1068_(_var.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{_var.FP.Get_symbolId()}))
        self.FP.Write(name)
        varNameList.Insert(name)
    }
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.Write(" = ")
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        } else {
            self.FP.writeln("")
        }
    }
    {
        __exp := node.FP.Get_unwrapBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.writeln("")
            self.FP.Write("if ")
            for _index, _varName := range( varNameList.Items ) {
                index := _index + 1
                varName := _varName.(string)
                if index > 1{
                    self.FP.Write(" or ")
                }
                self.FP.Write(" nil == " + varName)
            }
            self.FP.writeln(" then")
            self.FP.pushIndent(nil)
            for _, _varName := range( varNameList.Items ) {
                varName := _varName.(string)
                self.FP.writeln(Lns_getVM().String_format("local _%s = %s", []LnsAny{varName, varName}))
            }
            self.FP.popIndent()
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
            {
                _thenBlock := node.FP.Get_thenBlock()
                if _thenBlock != nil {
                    thenBlock := _thenBlock.(*Nodes_BlockNode)
                    self.FP.writeln("else")
                    self.FP.pushIndent(nil)
                    convLua_filter_1163_(&thenBlock.Nodes_Node, self, &node.Nodes_Node)
                    self.FP.popIndent()
                }
            }
            self.FP.writeln("end")
        }
    }
    {
        __exp := node.FP.Get_syncBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
            var syncVarNameList *LnsList
            syncVarNameList = NewLnsList([]LnsAny{})
            for _, _varInfo := range( node.FP.Get_syncVarList().Items ) {
                varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                var name string
                name = convLua_getSymTxt_1068_(varInfo.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{varInfo.FP.Get_symbolId()}))
                syncVarNameList.Insert(name)
                self.FP.writeln(Lns_getVM().String_format("_sync_%s = %s", []LnsAny{name, name}))
            }
            self.FP.popIndent()
            self.FP.writeln("end")
            for _, _name := range( syncVarNameList.Items ) {
                name := _name.(string)
                self.FP.writeln(Lns_getVM().String_format("%s = _sync_%s", []LnsAny{name, name}))
            }
            self.FP.popIndent()
            self.FP.writeln("end")
        }
    }
    if node.FP.Get_accessMode() == Ast_AccessMode__Pub{
        self.FP.writeln("")
        for _index, _varName := range( varNameList.Items ) {
            index := _index + 1
            varName := _varName.(string)
            var name string
            name = varName
            if self.needModuleObj{
                self.FP.writeln(Lns_getVM().String_format("_moduleObj.%s = %s", []LnsAny{name, name}))
            }
            self.pubVarName2InfoMap.Set(name,NewconvLua_PubVerInfo(node.FP.Get_staticFlag(), node.FP.Get_accessMode(), node.FP.Get_symbolInfoList().GetAt(index).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_mutable(), node.FP.Get_typeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))
        }
    }
    if self.macroDepth > 0{
        self.FP.writeln("")
        for _, _symbolInfo := range( node.FP.Get_symbolInfoList().Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            var varName string
            varName = convLua_getSymTxt_1068_(symbolInfo.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{symbolInfo.FP.Get_symbolId()}))
            self.FP.writeln(Lns_getVM().String_format("table.insert( macroVar.__names, '%s' )", []LnsAny{varName}))
            self.FP.writeln(Lns_getVM().String_format("macroVar.%s = %s", []LnsAny{varName, varName}))
            self.macroVarSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
        }
    }
}

// 2158: decl @lune.@base.@convLua.convFilter.processDeclArg
func (self *convLua_convFilter) ProcessDeclArg(node *Nodes_DeclArgNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{node.FP.Get_name().Txt}))
}

// 2166: decl @lune.@base.@convLua.convFilter.processDeclArgDDD
func (self *convLua_convFilter) ProcessDeclArgDDD(node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    self.FP.Write("...")
}

// 2172: decl @lune.@base.@convLua.convFilter.processExpDDD
func (self *convLua_convFilter) ProcessExpDDD(node *Nodes_ExpDDDNode,_opt LnsAny) {
    self.FP.Write("...")
}

// 2178: decl @lune.@base.@convLua.convFilter.processDeclFunc
func (self *convLua_convFilter) ProcessDeclFunc(node *Nodes_DeclFuncNode,_opt LnsAny) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo()
    var nameToken LnsAny
    nameToken = declInfo.FP.Get_name()
    var name string
    name = ""
    {
        __exp := nameToken
        if __exp != nil {
            _exp := __exp.(*Parser_Token)
            name = _exp.Txt
            
        }
    }
    var letTxt string
    letTxt = ""
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( declInfo.FP.Get_accessMode() != Ast_AccessMode__Global) &&
        Lns_setStackVal( len(name) != 0) ).(bool)){
        letTxt = "local "
        
    }
    self.FP.Write(Lns_getVM().String_format("%sfunction %s( ", []LnsAny{letTxt, name}))
    var argList *LnsList
    argList = declInfo.FP.Get_argList()
    self.FP.processExpListSub(&node.Nodes_Node, argList, nil)
    self.FP.writeln(" )")
    {
        __exp := declInfo.FP.Get_body()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.process__func__symbol(declInfo.FP.Get_has__func__Symbol(), node.FP.Get_expType().FP.Get_parentInfo(), name)
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write("end")
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType()
    if expType.FP.Get_accessMode() == Ast_AccessMode__Pub{
        if self.needModuleObj{
            self.FP.writeln("")
            self.FP.Write(Lns_getVM().String_format("_moduleObj.%s = %s", []LnsAny{name, name}))
        }
        self.pubFuncName2InfoMap.Set(name,NewconvLua_PubFuncInfo(declInfo.FP.Get_accessMode(), node.FP.Get_expType()))
    }
}

// 2219: decl @lune.@base.@convLua.convFilter.processRefType
func (self *convLua_convFilter) ProcessRefType(node *Nodes_RefTypeNode,_opt LnsAny) {
    self.FP.Write((Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( node.FP.Get_refFlag()) &&
        Lns_setStackVal( "&") ||
        Lns_setStackVal( "") ).(string)) + (Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( node.FP.Get_mutFlag()) &&
        Lns_setStackVal( "mut ") ||
        Lns_setStackVal( "") ).(string)))
    convLua_filter_1163_(node.FP.Get_name(), self, &node.Nodes_Node)
    if node.FP.Get_array() == "array"{
        self.FP.Write("[@]")
    } else if node.FP.Get_array() == "list"{
        self.FP.Write("[]")
    }
}

// 2233: decl @lune.@base.@convLua.convFilter.processIf
func (self *convLua_convFilter) ProcessIf(node *Nodes_IfNode,_opt LnsAny) {
    var valList *LnsList
    valList = node.FP.Get_stmtList()
    for _index, _val := range( valList.Items ) {
        index := _index + 1
        val := _val.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if index == 1{
            self.FP.Write("if ")
            convLua_filter_1163_(val.FP.Get_exp(), self, &node.Nodes_Node)
        } else if val.FP.Get_kind() == Nodes_IfKind__ElseIf{
            self.FP.Write("elseif ")
            convLua_filter_1163_(val.FP.Get_exp(), self, &node.Nodes_Node)
        } else { 
            self.FP.writeln("else")
        }
        self.FP.Write(" ")
        convLua_filter_1163_(&val.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    }
    self.FP.writeln("end")
}

// 2255: decl @lune.@base.@convLua.convFilter.processSwitch
func (self *convLua_convFilter) ProcessSwitch(node *Nodes_SwitchNode,_opt LnsAny) {
    self.FP.writeln("do")
    self.FP.pushIndent(nil)
    self.FP.Write("local _switchExp = ")
    convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.writeln("")
    if node.FP.Get_caseList().Len() > 0{
        for _index, _caseInfo := range( node.FP.Get_caseList().Items ) {
            index := _index + 1
            caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
            if index == 1{
                self.FP.Write("if ")
            } else { 
                self.FP.Write("elseif ")
            }
            var expList *Nodes_ExpListNode
            expList = caseInfo.FP.Get_expList()
            for _listIndex, _expNode := range( expList.FP.Get_expList().Items ) {
                listIndex := _listIndex + 1
                expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                if listIndex != 1{
                    self.FP.Write(" or ")
                }
                self.FP.Write("_switchExp == ")
                convLua_filter_1163_(expNode, self, &node.Nodes_Node)
            }
            self.FP.Write(" then")
            convLua_filter_1163_(&caseInfo.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
        }
        {
            __exp := node.FP.Get_default()
            if __exp != nil {
                _exp := __exp.(*Nodes_BlockNode)
                self.FP.writeln("else ")
                self.FP.pushIndent(nil)
                convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
                self.FP.popIndent()
            }
        }
        self.FP.writeln("end")
    }
    self.FP.popIndent()
    self.FP.writeln("end")
}

// 2298: decl @lune.@base.@convLua.convFilter.processMatch
func (self *convLua_convFilter) ProcessMatch(node *Nodes_MatchNode,_opt LnsAny) {
    self.FP.writeln("do")
    self.FP.pushIndent(nil)
    self.FP.Write("local _matchExp = ")
    convLua_filter_1163_(node.FP.Get_val(), self, &node.Nodes_Node)
    self.FP.writeln("")
    if node.FP.Get_caseList().Len() > 0{
        var fullName string
        fullName = self.FP.getFullName(&node.FP.Get_algeTypeInfo().Ast_TypeInfo)
        for _index, _caseInfo := range( node.FP.Get_caseList().Items ) {
            index := _index + 1
            caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
            if index == 1{
                self.FP.Write("if ")
            } else { 
                self.FP.Write("elseif ")
            }
            self.FP.writeln(Lns_getVM().String_format("_matchExp[1] == %s.%s[1] then", []LnsAny{fullName, caseInfo.FP.Get_valInfo().FP.Get_name()}))
            for _paramNum, _paramSym := range( caseInfo.FP.Get_valParamNameList().Items ) {
                paramNum := _paramNum + 1
                paramSym := _paramSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.writeln(Lns_getVM().String_format("   local %s = _matchExp[2][%d]", []LnsAny{paramSym.FP.Get_name(), paramNum}))
            }
            convLua_filter_1163_(&caseInfo.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
        }
        {
            __exp := node.FP.Get_defaultBlock()
            if __exp != nil {
                _exp := __exp.(*Nodes_Node)
                self.FP.writeln("else ")
                self.FP.pushIndent(nil)
                convLua_filter_1163_(_exp, self, &node.Nodes_Node)
                self.FP.popIndent()
            }
        }
        self.FP.writeln("end")
    }
    self.FP.popIndent()
    self.FP.writeln("end")
}

// 2337: decl @lune.@base.@convLua.convFilter.processWhile
func (self *convLua_convFilter) ProcessWhile(node *Nodes_WhileNode,_opt LnsAny) {
    self.FP.Write("while ")
    convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Write(" ")
    convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln("end")
}

// 2348: decl @lune.@base.@convLua.convFilter.processRepeat
func (self *convLua_convFilter) ProcessRepeat(node *Nodes_RepeatNode,_opt LnsAny) {
    self.FP.Write("repeat ")
    convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.Write("until ")
    convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
}

// 2357: decl @lune.@base.@convLua.convFilter.processFor
func (self *convLua_convFilter) ProcessFor(node *Nodes_ForNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("for %s = ", []LnsAny{convLua_getSymTxt_1068_(node.FP.Get_val().FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{node.FP.Get_val().FP.Get_symbolId()}))}))
    convLua_filter_1163_(node.FP.Get_init(), self, &node.Nodes_Node)
    self.FP.Write(", ")
    convLua_filter_1163_(node.FP.Get_to(), self, &node.Nodes_Node)
    {
        __exp := node.FP.Get_delta()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write(", ")
            convLua_filter_1163_(_exp, self, &node.Nodes_Node)
        }
    }
    self.FP.Write(" ")
    convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln("end")
}

// 2373: decl @lune.@base.@convLua.convFilter.processApply
func (self *convLua_convFilter) ProcessApply(node *Nodes_ApplyNode,_opt LnsAny) {
    self.FP.Write("for ")
    var varList *LnsList
    varList = node.FP.Get_varList()
    for _index, __var := range( varList.Items ) {
        index := _index + 1
        _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(", ")
        }
        self.FP.Write(convLua_getSymTxt_1068_(_var.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{_var.FP.Get_symbolId()})))
    }
    self.FP.Write(" in ")
    convLua_filter_1163_(&node.FP.Get_expList().Nodes_Node, self, &node.Nodes_Node)
    self.FP.Write(" ")
    convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln("end")
}

// 2391: decl @lune.@base.@convLua.convFilter.processForeach
func (self *convLua_convFilter) ProcessForeach(node *Nodes_ForeachNode,_opt LnsAny) {
    var keySym LnsAny
    var valSym LnsAny
    if node.FP.Get_exp().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Set{
        keySym = node.FP.Get_val()
        
        valSym = node.FP.Get_key()
        
    } else { 
        keySym = node.FP.Get_key()
        
        valSym = node.FP.Get_val()
        
    }
    self.FP.Write("for ")
    if keySym != nil{
        keySym_6549 := keySym.(*Ast_SymbolInfo)
        self.FP.Write(convLua_getSymTxt_1068_(keySym_6549.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{keySym_6549.FP.Get_symbolId()})))
    } else {
        self.FP.Write("__index")
    }
    self.FP.Write(", ")
    if valSym != nil{
        valSym_6552 := valSym.(*Ast_SymbolInfo)
        self.FP.Write(convLua_getSymTxt_1068_(valSym_6552.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{valSym_6552.FP.Get_symbolId()})))
    } else {
        self.FP.Write("__val")
    }
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( self.useIpairs) &&
        Lns_setStackVal( node.FP.Get_exp().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__List) ).(bool)){
        if node.FP.Get_exp().FP.Get_expType().FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilable(){
            self.FP.Write(" in pairs( ")
        } else { 
            self.FP.Write(" in ipairs( ")
        }
    } else { 
        self.FP.Write(" in pairs( ")
    }
    convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Write(" ) ")
    convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln("end")
}

// 2441: decl @lune.@base.@convLua.convFilter.processForsort
func (self *convLua_convFilter) ProcessForsort(node *Nodes_ForsortNode,_opt LnsAny) {
    var keySym LnsAny
    var valSym LnsAny
    if node.FP.Get_exp().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Set{
        keySym = node.FP.Get_val()
        
        valSym = node.FP.Get_key()
        
    } else { 
        keySym = node.FP.Get_key()
        
        valSym = node.FP.Get_val()
        
    }
    self.FP.writeln("do")
    self.FP.pushIndent(nil)
    self.FP.writeln("local __sorted = {}")
    self.FP.Write("local __map = ")
    convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.writeln("")
    self.FP.writeln("for __key in pairs( __map ) do")
    self.FP.pushIndent(nil)
    self.FP.writeln("table.insert( __sorted, __key )")
    self.FP.popIndent()
    self.FP.writeln("end")
    self.FP.writeln("table.sort( __sorted )")
    self.FP.Write("for __index, ")
    var key string
    key = "__key"
    if keySym != nil{
        keySym_6571 := keySym.(*Ast_SymbolInfo)
        key = convLua_getSymTxt_1068_(keySym_6571.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{keySym_6571.FP.Get_symbolId()}))
        
    }
    self.FP.Write(key)
    self.FP.writeln(" in ipairs( __sorted ) do")
    self.FP.pushIndent(nil)
    if valSym != nil{
        valSym_6573 := valSym.(*Ast_SymbolInfo)
        self.FP.writeln(Lns_getVM().String_format("local %s = __map[ %s ]", []LnsAny{convLua_getSymTxt_1068_(valSym_6573.FP.Get_name(), Lns_getVM().String_format("%d", []LnsAny{valSym_6573.FP.Get_symbolId()})), key}))
    }
    convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.writeln("end")
    self.FP.popIndent()
    self.FP.writeln("end")
    self.FP.popIndent()
    self.FP.writeln("end")
}

// 2494: decl @lune.@base.@convLua.convFilter.processExpUnwrap
func (self *convLua_convFilter) ProcessExpUnwrap(node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    {
        __exp := node.FP.Get_default()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            self.FP.Write("_lune.unwrapDefault( ")
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(", ")
            convLua_filter_1163_(_exp, self, &node.Nodes_Node)
            self.FP.Write(")")
        } else {
            self.FP.Write("_lune.unwrap( ")
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(")")
        }
    }
}

// 2512: decl @lune.@base.@convLua.convFilter.processExpCall
func (self *convLua_convFilter) ProcessExpCall(node *Nodes_ExpCallNode,_opt LnsAny) {
    var wroteFuncFlag bool
    wroteFuncFlag = false
    var setArgFlag bool
    setArgFlag = false
    var fieldCall func() bool
    fieldCall = func() bool {
        var fieldNode *Nodes_RefFieldNode
        
        {
            _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func().FP)
            if _fieldNode == nil{
                return true
            } else {
                fieldNode = _fieldNode.(*Nodes_RefFieldNode)
            }
        }
        var prefixNode *Nodes_Node
        prefixNode = fieldNode.FP.Get_prefix()
        var processSet func() bool
        processSet = func() bool {
            setArgFlag = true
            
            wroteFuncFlag = true
            
            if _switch12639 := fieldNode.FP.Get_field().Txt; _switch12639 == "add" || _switch12639 == "del" {
                convLua_filter_1163_(prefixNode, self, &fieldNode.Nodes_Node)
                self.FP.Write("[")
                {
                    _argList := node.FP.Get_argList()
                    if _argList != nil {
                        argList := _argList.(*Nodes_ExpListNode)
                        convLua_filter_1163_(&argList.Nodes_Node, self, &fieldNode.Nodes_Node)
                    }
                }
                self.FP.Write("]")
                if _switch12633 := fieldNode.FP.Get_field().Txt; _switch12633 == "add" {
                    self.FP.Write("= true")
                } else if _switch12633 == "del" {
                    self.FP.Write("= nil")
                }
                return false
            }
            self.FP.Write(Lns_getVM().String_format("_lune._Set_%s(", []LnsAny{fieldNode.FP.Get_field().Txt}))
            convLua_filter_1163_(prefixNode, self, &fieldNode.Nodes_Node)
            return true
        }
        var prefixType *Ast_TypeInfo
        prefixType = prefixNode.FP.Get_expType()
        var processEnumAlge func()
        processEnumAlge = func() {
            wroteFuncFlag = true
            
            var fieldExpType *Ast_TypeInfo
            fieldExpType = fieldNode.FP.Get_expType()
            var canonicalName string
            canonicalName = self.FP.getFullName(prefixType)
            var methodName string
            methodName = fieldNode.FP.Get_field().Txt
            var delimit string
            delimit = ":"
            if methodName == "get__txt"{
                methodName = "_getTxt"
                
            }
            if fieldExpType.FP.Get_kind() == Ast_TypeInfoKind__Func{
                delimit = "."
                
            }
            self.FP.Write(Lns_getVM().String_format("%s%s%s( ", []LnsAny{canonicalName, delimit, methodName}))
            if fieldExpType.FP.Get_staticFlag(){
                setArgFlag = false
                
            } else { 
                convLua_filter_1163_(prefixNode, self, &fieldNode.Nodes_Node)
                setArgFlag = true
                
            }
        }
        if node.FP.Get_nilAccess(){
            wroteFuncFlag = true
            
            setArgFlag = true
            
            if _switch12893 := prefixType.FP.Get_kind(); _switch12893 == Ast_TypeInfoKind__List || _switch12893 == Ast_TypeInfoKind__Array {
                self.FP.Write(Lns_getVM().String_format("_lune.nilacc( table.%s, nil, 'list', ", []LnsAny{fieldNode.FP.Get_field().Txt}))
                convLua_filter_1163_(prefixNode, self, &fieldNode.Nodes_Node)
            } else {
                self.FP.Write("_lune.nilacc( ")
                convLua_filter_1163_(prefixNode, self, &fieldNode.Nodes_Node)
                self.FP.Write(Lns_getVM().String_format(", '%s', 'callmtd' ", []LnsAny{fieldNode.FP.Get_field().Txt}))
            }
        } else { 
            if _switch13139 := prefixType.FP.Get_kind(); _switch13139 == Ast_TypeInfoKind__List || _switch13139 == Ast_TypeInfoKind__Array {
                setArgFlag = true
                
                wroteFuncFlag = true
                
                self.FP.Write(Lns_getVM().String_format("table.%s( ", []LnsAny{fieldNode.FP.Get_field().Txt}))
                convLua_filter_1163_(prefixNode, self, &fieldNode.Nodes_Node)
            } else if _switch13139 == Ast_TypeInfoKind__Set {
                if Lns_op_not(processSet()){
                    return false
                }
            } else if _switch13139 == Ast_TypeInfoKind__Enum || _switch13139 == Ast_TypeInfoKind__Alge {
                processEnumAlge()
            } else if _switch13139 == Ast_TypeInfoKind__Box {
                convLua_filter_1163_(prefixNode, self, &fieldNode.Nodes_Node)
                self.FP.Write("[1]")
                return false
            } else if _switch13139 == Ast_TypeInfoKind__Class {
                if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                    Lns_setStackVal( prefixType.FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, nil)) &&
                    Lns_setStackVal( convLua_isGenericType_1316_(prefixType)) &&
                    Lns_setStackVal( (Lns_popVal( Lns_incStack() ||
                        Lns_setStackVal( fieldNode.FP.Get_field().Txt == "_fromMap") ||
                        Lns_setStackVal( fieldNode.FP.Get_field().Txt == "_fromStem") ).(bool))) ).(bool)){
                    wroteFuncFlag = true
                    
                    setArgFlag = true
                    
                    convLua_filter_1163_(node.FP.Get_func(), self, &node.Nodes_Node)
                    self.FP.Write("( ")
                    {
                        _argList := node.FP.Get_argList()
                        if _argList != nil {
                            argList := _argList.(*Nodes_ExpListNode)
                            convLua_filter_1163_(&argList.Nodes_Node, self, &node.Nodes_Node)
                            self.FP.Write(", ")
                        }
                    }
                    self.FP.outputAlter2MapFunc(self.FP, prefixType.FP.CreateAlt2typeMap(false))
                    self.FP.Write(")")
                    return false
                }
            }
        }
        return true
    }
    if Lns_op_not(fieldCall()){
        return 
    }
    {
        _refNode := Nodes_ExpRefNodeDownCastF(node.FP.Get_func().FP)
        if _refNode != nil {
            refNode := _refNode.(*Nodes_ExpRefNode)
            if refNode.FP.Get_symbolInfo().FP.Get_name() == "super"{
                wroteFuncFlag = true
                
                setArgFlag = true
                
                var funcType *Ast_TypeInfo
                funcType = refNode.FP.Get_expType()
                self.FP.Write(Lns_getVM().String_format("%s.%s( self ", []LnsAny{self.FP.getFullName(funcType.FP.Get_parentInfo()), funcType.FP.Get_rawTxt()}))
            } else if refNode.FP.Get_expType() == TransUnit_getBuiltinFunc().Lns_expandLuavalMap{
                wroteFuncFlag = true
                
                self.FP.Write("(")
            }
        }
    }
    if Lns_op_not(wroteFuncFlag){
        if node.FP.Get_nilAccess(){
            self.FP.Write("_lune.nilacc( ")
            convLua_filter_1163_(node.FP.Get_func(), self, &node.Nodes_Node)
            self.FP.Write(", nil, 'call'")
            wroteFuncFlag = true
            
            setArgFlag = true
            
        } else { 
            convLua_filter_1163_(node.FP.Get_func(), self, &node.Nodes_Node)
            self.FP.Write("( ")
        }
    }
    var convStrFlag bool
    convStrFlag = false
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( Lns_op_not(self.targetLuaVer.FP.Get_canFormStem2Str())) &&
        Lns_setStackVal( TransUnit_isStrFormFunc(node.FP.Get_func().FP.Get_expType())) ).(bool)){
        convStrFlag = true
        
    }
    {
        _argList := node.FP.Get_argList()
        if _argList != nil {
            argList := _argList.(*Nodes_ExpListNode)
            var expList *LnsList
            expList = NewLnsList([]LnsAny{})
            for _, _expNode := range( argList.FP.Get_expList().Items ) {
                expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                if expNode.FP.Get_expType().FP.Get_kind() != Ast_TypeInfoKind__Abbr{
                    {
                        _toDDD := Nodes_ExpToDDDNodeDownCastF(expNode.FP)
                        if _toDDD != nil {
                            toDDD := _toDDD.(*Nodes_ExpToDDDNode)
                            for _, _appNode := range( toDDD.FP.Get_expList().FP.Get_expList().Items ) {
                                appNode := _appNode.(Nodes_NodeDownCast).ToNodes_Node()
                                expList.Insert(Nodes_Node2Stem(appNode))
                            }
                        } else {
                            expList.Insert(Nodes_Node2Stem(expNode))
                        }
                    }
                }
            }
            if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                Lns_setStackVal( wroteFuncFlag) &&
                Lns_setStackVal( setArgFlag) ).(bool)){
                if expList.Len() > 0{
                    self.FP.Write(", ")
                }
            }
            if convStrFlag{
                var opList *LnsList
                opList = NewLnsList([]LnsAny{})
                if expList.Len() > 0{
                    var literal LnsAny
                    literal = convLua_convExp13476(Lns_2DDD(expList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node().FP.GetLiteral()))
                    if literal != nil{
                        literal_6662 := literal
                        switch _exp13495 := literal_6662.(type) {
                        case *Nodes_Literal__Str:
                        txt := _exp13495.Val1
                            opList = TransUnit_findForm(txt)
                            
                        }
                    }
                }
                for _index, _argNode := range( expList.Items ) {
                    index := _index + 1
                    argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                    var filtered bool
                    filtered = false
                    if index > 1{
                        self.FP.Write(", ")
                        if index - 1 <= opList.Len(){
                            var formType LnsInt
                            formType = convLua_convExp13551(Lns_2DDD(TransUnit_isMatchStringFormatType(opList.GetAt(index - 1).(string), argNode.FP.Get_expType(), self.targetLuaVer)))
                            if formType == TransUnit_FormType__NeedConv{
                                self.FP.Write("tostring( ")
                                convLua_filter_1163_(argNode, self, &node.Nodes_Node)
                                self.FP.Write(")")
                                filtered = true
                                
                            }
                        }
                    }
                    if Lns_op_not(filtered){
                        convLua_filter_1163_(argNode, self, &node.Nodes_Node)
                    }
                }
            } else { 
                convLua_filter_1163_(&argList.Nodes_Node, self, &node.Nodes_Node)
            }
        }
    }
    self.FP.Write(" )")
}

// 2741: decl @lune.@base.@convLua.convFilter.processExpList
func (self *convLua_convFilter) ProcessExpList(node *Nodes_ExpListNode,_opt LnsAny) {
    var expList *LnsList
    expList = node.FP.Get_expList()
    self.FP.processExpListSub(&node.Nodes_Node, expList, node.FP.Get_mRetExp())
}

// 2750: decl @lune.@base.@convLua.convFilter.processExpOp1
func (self *convLua_convFilter) ProcessExpOp1(node *Nodes_ExpOp1Node,_opt LnsAny) {
    var op string
    op = node.FP.Get_op().Txt
    if op == ",,,"{
        convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
    } else if op == ",,,,"{
        if node.FP.Get_macroMode() != Nodes_MacroMode__None{
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
        } else { 
            self.FP.Write("__luneSym2Str( ")
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(" )")
        }
    } else if op == ",,"{
        if _switch13846 := node.FP.Get_exp().FP.Get_expType(); _switch13846 == Ast_builtinTypeInt || _switch13846 == Ast_builtinTypeReal || _switch13846 == Ast_builtinTypeBool {
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
        } else {
            self.FP.Write("__luneGetLocal( ")
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(" )")
        }
    } else if op == "~"{
        if self.targetLuaVer.FP.Get_hasBitOp() == LuaVer_BitOp__HasOp{
            self.FP.Write(op)
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
        } else { 
            self.FP.Write("bit32.bnot( ")
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(" )")
        }
    } else { 
        if op == "not"{
            op = op + " "
            
        }
        self.FP.Write(op)
        convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
    }
}

// 2799: decl @lune.@base.@convLua.convFilter.processExpToDDD
func (self *convLua_convFilter) ProcessExpToDDD(node *Nodes_ExpToDDDNode,_opt LnsAny) {
    self.FP.processExpListSub(&node.Nodes_Node, node.FP.Get_expList().FP.Get_expList(), node.FP.Get_expList().FP.Get_mRetExp())
}

// 2805: decl @lune.@base.@convLua.convFilter.processExpMultiTo1
func (self *convLua_convFilter) ProcessExpMultiTo1(node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
}

// 2811: decl @lune.@base.@convLua.convFilter.processExpCast
func (self *convLua_convFilter) ProcessExpCast(node *Nodes_ExpCastNode,_opt LnsAny) {
    if _switch14265 := node.FP.Get_castKind(); _switch14265 == Nodes_CastKind__Force {
        if node.FP.Get_expType().FP.Equals(self.processInfo, Ast_builtinTypeInt, nil, nil){
            self.FP.Write("math.floor(")
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(")")
        } else if node.FP.Get_expType().FP.Equals(self.processInfo, Ast_builtinTypeReal, nil, nil){
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(" * 1.0")
        } else { 
            convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
        }
    } else if _switch14265 == Nodes_CastKind__Normal {
        self.FP.Write("_lune.__Cast( ")
        convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
        var castKind LnsInt
        var classObj string
        classObj = "nil"
        if _switch14229 := node.FP.Get_expType().FP.Get_nonnilableType(); _switch14229 == Ast_builtinTypeInt {
            castKind = LuaMod_CastKind__Int
            
        } else if _switch14229 == Ast_builtinTypeReal {
            castKind = LuaMod_CastKind__Real
            
        } else if _switch14229 == Ast_builtinTypeString {
            castKind = LuaMod_CastKind__Str
            
        } else {
            castKind = LuaMod_CastKind__Class
            
            classObj = self.FP.getFullName(node.FP.Get_expType().FP.Get_nonnilableType())
            
        }
        self.FP.Write(Lns_getVM().String_format(", %d, %s )", []LnsAny{castKind, classObj}))
    } else if _switch14265 == Nodes_CastKind__Implicit {
        convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
    }
}

// 2858: decl @lune.@base.@convLua.convFilter.processExpParen
func (self *convLua_convFilter) ProcessExpParen(node *Nodes_ExpParenNode,_opt LnsAny) {
    self.FP.Write("(")
    convLua_filter_1163_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Write(" )")
}

// 2867: decl @lune.@base.@convLua.convFilter.processExpSetVal
func (self *convLua_convFilter) ProcessExpSetVal(node *Nodes_ExpSetValNode,_opt LnsAny) {
    convLua_filter_1163_(node.FP.Get_exp1(), self, &node.Nodes_Node)
    self.FP.Write(" = ")
    convLua_filter_1163_(&node.FP.Get_exp2().Nodes_Node, self, &node.Nodes_Node)
}

// 2874: decl @lune.@base.@convLua.convFilter.processExpSetItem
func (self *convLua_convFilter) ProcessExpSetItem(node *Nodes_ExpSetItemNode,_opt LnsAny) {
    convLua_filter_1163_(node.FP.Get_val(), self, &node.Nodes_Node)
    self.FP.Write("[")
    switch _exp14423 := node.FP.Get_index().(type) {
    case *Nodes_IndexVal__NodeIdx:
    index := _exp14423.Val1
        convLua_filter_1163_(index, self, &node.Nodes_Node)
    case *Nodes_IndexVal__SymIdx:
    index := _exp14423.Val1
        self.FP.Write(Lns_getVM().String_format("'%s'", []LnsAny{index}))
    }
    self.FP.Write("]")
    self.FP.Write(" = ")
    convLua_filter_1163_(node.FP.Get_exp2(), self, &node.Nodes_Node)
}

// 2892: decl @lune.@base.@convLua.convFilter.processExpOp2
func (self *convLua_convFilter) ProcessExpOp2(node *Nodes_ExpOp2Node,_opt LnsAny) {
    var intCast bool
    intCast = false
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( node.FP.Get_expType().FP.Equals(self.processInfo, Ast_builtinTypeInt, nil, nil)) &&
        Lns_setStackVal( node.FP.Get_op().Txt == "/") ).(bool)){
        intCast = true
        
        self.FP.Write("math.floor(")
    }
    var opTxt string
    opTxt = node.FP.Get_op().Txt
    {
        __exp := Ast_bitBinOpMap.Items[opTxt]
        if __exp != nil {
            _exp := __exp.(LnsInt)
            if self.targetLuaVer.FP.Get_hasBitOp() == LuaVer_BitOp__HasOp{
                if _switch14553 := _exp; _switch14553 == Ast_BitOpKind__LShift {
                    opTxt = "<<"
                    
                } else if _switch14553 == Ast_BitOpKind__RShift {
                    opTxt = ">>"
                    
                }
                convLua_filter_1163_(node.FP.Get_exp1(), self, &node.Nodes_Node)
                self.FP.Write(" " + opTxt + " ")
                convLua_filter_1163_(node.FP.Get_exp2(), self, &node.Nodes_Node)
            } else { 
                var binfunc string
                binfunc = ""
                var exp2Mod string
                exp2Mod = ""
                if _switch14670 := _exp; _switch14670 == Ast_BitOpKind__And {
                    binfunc = "band"
                    
                } else if _switch14670 == Ast_BitOpKind__Or {
                    binfunc = "bor"
                    
                } else if _switch14670 == Ast_BitOpKind__Xor {
                    binfunc = "bxor"
                    
                } else if _switch14670 == Ast_BitOpKind__LShift {
                    binfunc = "lshift"
                    
                } else if _switch14670 == Ast_BitOpKind__RShift {
                    binfunc = "lshift"
                    
                    exp2Mod = "-"
                    
                }
                self.FP.Write(Lns_getVM().String_format("bit32.%s(", []LnsAny{binfunc}))
                convLua_filter_1163_(node.FP.Get_exp1(), self, &node.Nodes_Node)
                self.FP.Write(", ")
                self.FP.Write(exp2Mod)
                convLua_filter_1163_(node.FP.Get_exp2(), self, &node.Nodes_Node)
                self.FP.Write(" )")
            }
        } else {
            convLua_filter_1163_(node.FP.Get_exp1(), self, &node.Nodes_Node)
            self.FP.Write(" " + opTxt + " ")
            convLua_filter_1163_(node.FP.Get_exp2(), self, &node.Nodes_Node)
        }
    }
    if intCast{
        self.FP.Write(")")
    }
}

// 2962: decl @lune.@base.@convLua.convFilter.processExpRef
func (self *convLua_convFilter) ProcessExpRef(node *Nodes_ExpRefNode,_opt LnsAny) {
    if _switch14950 := node.FP.Get_symbolInfo().FP.Get_name(); _switch14950 == "super" {
        var funcType *Ast_TypeInfo
        funcType = node.FP.Get_expType()
        self.FP.Write(Lns_getVM().String_format("%s.%s", []LnsAny{self.FP.getFullName(funcType.FP.Get_parentInfo()), funcType.FP.Get_rawTxt()}))
    } else {
        var builtinFunc *TransUnit_BuiltinFuncType
        builtinFunc = TransUnit_getBuiltinFunc()
        if node.FP.Get_expType().FP.Equals(self.processInfo, builtinFunc.Lns__load, nil, nil){
            self.FP.Write("_lune." + self.targetLuaVer.FP.Get_loadStrFuncName())
        } else { 
            if self.macroVarSymSet.Has(Ast_SymbolInfo2Stem(node.FP.Get_symbolInfo().FP.GetOrg())){
                self.FP.Write("macroVar.")
            } else { 
                if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                    Lns_setStackVal( node.FP.Get_symbolInfo().FP.Get_accessMode() == Ast_AccessMode__Pub) &&
                    Lns_setStackVal( node.FP.Get_symbolInfo().FP.Get_kind() == Ast_SymbolKind__Var) ).(bool)){
                    if self.needModuleObj{
                        self.FP.Write("_moduleObj.")
                    }
                }
            }
            self.FP.Write(node.FP.Get_symbolInfo().FP.Get_name())
        }
    }
}

// 2995: decl @lune.@base.@convLua.convFilter.processExpRefItem
func (self *convLua_convFilter) ProcessExpRefItem(node *Nodes_ExpRefItemNode,_opt LnsAny) {
    if node.FP.Get_nilAccess(){
        self.FP.Write("_lune.nilacc( ")
        convLua_filter_1163_(node.FP.Get_val(), self, &node.Nodes_Node)
        self.FP.Write(", nil, 'item', ")
        {
            __exp := node.FP.Get_index()
            if __exp != nil {
                _exp := __exp.(*Nodes_Node)
                convLua_filter_1163_(_exp, self, &node.Nodes_Node)
            } else {
                self.FP.Write(Lns_getVM().String_format("'%s'", []LnsAny{Lns_unwrap( node.FP.Get_symbol()).(string)}))
            }
        }
        self.FP.Write(")")
    } else { 
        if node.FP.Get_val().FP.Get_expType().FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil){
            self.FP.Write("string.byte( ")
            convLua_filter_1163_(node.FP.Get_val(), self, &node.Nodes_Node)
            self.FP.Write(", ")
            {
                __exp := node.FP.Get_index()
                if __exp != nil {
                    _exp := __exp.(*Nodes_Node)
                    convLua_filter_1163_(_exp, self, &node.Nodes_Node)
                } else {
                    panic("index is nil")
                }
            }
            self.FP.Write(" )")
        } else { 
            convLua_filter_1163_(node.FP.Get_val(), self, &node.Nodes_Node)
            self.FP.Write("[")
            {
                __exp := node.FP.Get_index()
                if __exp != nil {
                    _exp := __exp.(*Nodes_Node)
                    convLua_filter_1163_(_exp, self, &node.Nodes_Node)
                } else {
                    self.FP.Write(Lns_getVM().String_format("'%s'", []LnsAny{Lns_unwrap( node.FP.Get_symbol()).(string)}))
                }
            }
            self.FP.Write("]")
        }
    }
}

// 3037: decl @lune.@base.@convLua.convFilter.processRefField
func (self *convLua_convFilter) ProcessRefField(node *Nodes_RefFieldNode,_opt LnsAny) {
    opt := _opt.(*ConvLua_Opt)
    var parent *Nodes_Node
    parent = opt.Node
    var prefix *Nodes_Node
    prefix = node.FP.Get_prefix()
    if node.FP.Get_nilAccess(){
        self.FP.Write("_lune.nilacc( ")
        convLua_filter_1163_(prefix, self, &node.Nodes_Node)
        self.FP.Write(Lns_getVM().String_format(", \"%s\" )", []LnsAny{node.FP.Get_field().Txt}))
    } else { 
        convLua_filter_1163_(prefix, self, &node.Nodes_Node)
        var delimit string
        delimit = "."
        if parent.FP.Get_kind() == Nodes_NodeKind_get_ExpCall(){
            if node.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Method{
                delimit = ":"
                
            } else { 
                delimit = "."
                
            }
        }
        var fieldToken *Parser_Token
        fieldToken = node.FP.Get_field()
        self.FP.Write(delimit + fieldToken.Txt)
    }
}

// 3065: decl @lune.@base.@convLua.convFilter.processExpOmitEnum
func (self *convLua_convFilter) ProcessExpOmitEnum(node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("%s.%s", []LnsAny{self.FP.getFullName(node.FP.Get_expType()), node.FP.Get_valToken().Txt}))
}

// 3071: decl @lune.@base.@convLua.convFilter.processGetField
func (self *convLua_convFilter) ProcessGetField(node *Nodes_GetFieldNode,_opt LnsAny) {
    var prefixNode *Nodes_Node
    prefixNode = node.FP.Get_prefix()
    var prefixType *Ast_TypeInfo
    prefixType = prefixNode.FP.Get_expType()
    var fieldTxt string
    fieldTxt = node.FP.Get_field().Txt
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( fieldTxt == "_txt") &&
        Lns_setStackVal( (Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( prefixType.FP.Get_kind() == Ast_TypeInfoKind__Enum) ||
            Lns_setStackVal( prefixType.FP.Get_kind() == Ast_TypeInfoKind__Alge) ).(bool))) ).(bool)){
        self.FP.Write(Lns_getVM().String_format("%s:_getTxt( ", []LnsAny{self.FP.getFullName(prefixType)}))
        convLua_filter_1163_(prefixNode, self, &node.Nodes_Node)
        self.FP.writeln(")")
    } else { 
        if node.FP.Get_nilAccess(){
            fieldTxt = Lns_getVM().String_format("get_%s", []LnsAny{fieldTxt})
            
            self.FP.Write("_lune.nilacc( ")
            convLua_filter_1163_(prefixNode, self, &node.Nodes_Node)
            self.FP.Write(Lns_getVM().String_format(", '%s', 'callmtd' )", []LnsAny{fieldTxt}))
        } else { 
            fieldTxt = Lns_getVM().String_format("get_%s()", []LnsAny{fieldTxt})
            
            convLua_filter_1163_(prefixNode, self, &node.Nodes_Node)
            var delimit string
            delimit = "."
            if node.FP.Get_getterTypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Method{
                delimit = ":"
                
            } else { 
                delimit = "."
                
            }
            self.FP.Write(delimit + fieldTxt)
        }
    }
}

// 3111: decl @lune.@base.@convLua.convFilter.processReturn
func (self *convLua_convFilter) ProcessReturn(node *Nodes_ReturnNode,_opt LnsAny) {
    self.FP.Write("return ")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
}

// 3121: decl @lune.@base.@convLua.convFilter.processLuneKind
func (self *convLua_convFilter) ProcessLuneKind(node *Nodes_LuneKindNode,_opt LnsAny) {
    {
        _workNode := Nodes_ExpCastNodeDownCastF(node.FP.Get_exp().FP)
        if _workNode != nil {
            workNode := _workNode.(*Nodes_ExpCastNode)
            if workNode.FP.Get_castKind() == Nodes_CastKind__Implicit{
                self.FP.Write(Lns_getVM().String_format("%d", []LnsAny{workNode.FP.Get_exp().FP.Get_expType().FP.Get_kind()}))
            }
        } else {
            self.FP.Write(Lns_getVM().String_format("%d", []LnsAny{node.FP.Get_exp().FP.Get_expType().FP.Get_kind()}))
        }
    }
}

// 3134: decl @lune.@base.@convLua.convFilter.processTestCase
func (self *convLua_convFilter) ProcessTestCase(node *Nodes_TestCaseNode,_opt LnsAny) {
    if self.enableTest{
        self.FP.writeln("do")
        self.FP.pushIndent(nil)
        convLua_filter_1163_(node.FP.Get_impNode(), self, &node.Nodes_Node)
        self.FP.writeln("")
        self.FP.writeln(Lns_getVM().String_format("local function testcase( %s ) ", []LnsAny{node.FP.Get_ctrlName()}))
        convLua_filter_1163_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
        self.FP.writeln("end")
        self.FP.writeln(Lns_getVM().String_format("__t.registerTestcase( \"%s\", \"%s\", testcase )", []LnsAny{Lns_car(Lns_getVM().String_gsub(self.moduleTypeInfo.FP.GetFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), nil),"@", "")).(string), node.FP.Get_name().Txt}))
        self.FP.popIndent()
        self.FP.writeln("end")
    }
}

// 3156: decl @lune.@base.@convLua.convFilter.processTestBlock
func (self *convLua_convFilter) ProcessTestBlock(node *Nodes_TestBlockNode,_opt LnsAny) {
    if self.enableTest{
        for _, _statement := range( node.FP.Get_stmtList().Items ) {
            statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
            convLua_filter_1163_(statement, self, &node.Nodes_Node)
            self.FP.writeln("")
        }
    }
}

// 3167: decl @lune.@base.@convLua.convFilter.processProvide
func (self *convLua_convFilter) ProcessProvide(node *Nodes_ProvideNode,_opt LnsAny) {
}

// 3172: decl @lune.@base.@convLua.convFilter.processAlias
func (self *convLua_convFilter) ProcessAlias(node *Nodes_AliasNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("local %s = ", []LnsAny{node.FP.Get_newName()}))
    convLua_filter_1163_(node.FP.Get_srcNode(), self, &node.Nodes_Node)
    if Ast_isPubToExternal(node.FP.Get_expType().FP.Get_accessMode()){
        self.FP.Write(Lns_getVM().String_format("\n_moduleObj.%s = %s", []LnsAny{node.FP.Get_newName(), node.FP.Get_newName()}))
    }
}

// 3182: decl @lune.@base.@convLua.convFilter.processBoxing
func (self *convLua_convFilter) ProcessBoxing(node *Nodes_BoxingNode,_opt LnsAny) {
    self.FP.Write("{")
    convLua_filter_1163_(node.FP.Get_src(), self, &node.Nodes_Node)
    self.FP.Write("}")
}

// 3192: decl @lune.@base.@convLua.convFilter.processUnboxing
func (self *convLua_convFilter) ProcessUnboxing(node *Nodes_UnboxingNode,_opt LnsAny) {
    convLua_filter_1163_(node.FP.Get_src(), self, &node.Nodes_Node)
    self.FP.Write("[1]")
}

// 3200: decl @lune.@base.@convLua.convFilter.processLiteralList
func (self *convLua_convFilter) ProcessLiteralList(node *Nodes_LiteralListNode,_opt LnsAny) {
    self.FP.Write("{")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write("}")
}

// 3212: decl @lune.@base.@convLua.convFilter.processLiteralSet
func (self *convLua_convFilter) ProcessLiteralSet(node *Nodes_LiteralSetNode,_opt LnsAny) {
    self.FP.Write("{")
    {
        _expListNode := node.FP.Get_expList()
        if _expListNode != nil {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            for _index, _expNode := range( expListNode.FP.Get_expList().Items ) {
                index := _index + 1
                expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.Write(", ")
                }
                self.FP.Write("[")
                convLua_filter_1163_(expNode, self, &node.Nodes_Node)
                self.FP.Write("] = true")
            }
        }
    }
    self.FP.Write("}")
}

// 3230: decl @lune.@base.@convLua.convFilter.processLiteralMap
func (self *convLua_convFilter) ProcessLiteralMap(node *Nodes_LiteralMapNode,_opt LnsAny) {
    self.FP.Write("{")
    var pairList *LnsList
    pairList = node.FP.Get_pairList()
    for _index, _pair := range( pairList.Items ) {
        index := _index + 1
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        if index > 1{
            self.FP.Write(", ")
        }
        self.FP.Write("[")
        convLua_filter_1163_(pair.FP.Get_key(), self, &node.Nodes_Node)
        self.FP.Write("] = ")
        convLua_filter_1163_(pair.FP.Get_val(), self, &node.Nodes_Node)
    }
    self.FP.Write("}")
}

// 3248: decl @lune.@base.@convLua.convFilter.processLiteralArray
func (self *convLua_convFilter) ProcessLiteralArray(node *Nodes_LiteralArrayNode,_opt LnsAny) {
    self.FP.Write("{")
    {
        __exp := node.FP.Get_expList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            convLua_filter_1163_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Write("}")
}

// 3261: decl @lune.@base.@convLua.convFilter.processLiteralChar
func (self *convLua_convFilter) ProcessLiteralChar(node *Nodes_LiteralCharNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("%d", []LnsAny{node.FP.Get_num()}))
}

// 3267: decl @lune.@base.@convLua.convFilter.processLiteralInt
func (self *convLua_convFilter) ProcessLiteralInt(node *Nodes_LiteralIntNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 3275: decl @lune.@base.@convLua.convFilter.processLiteralReal
func (self *convLua_convFilter) ProcessLiteralReal(node *Nodes_LiteralRealNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 3282: decl @lune.@base.@convLua.convFilter.processLiteralString
func (self *convLua_convFilter) ProcessLiteralString(node *Nodes_LiteralStringNode,_opt LnsAny) {
    var txt string
    txt = node.FP.Get_token().Txt
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(txt, "^```", nil, nil))){
        txt = "[==[" + Lns_getVM().String_sub(txt,4, -4) + "]==]"
        
    }
    var opList *LnsList
    opList = TransUnit_findForm(txt)
    {
        _expList := node.FP.Get_orgParam()
        if _expList != nil {
            expList := _expList.(*Nodes_ExpListNode)
            var mRetIndex LnsAny
            mRetIndex = Lns_NilAccFin(Lns_NilAccPush(expList.FP.Get_mRetExp()) && 
            Lns_NilAccCall1( func () LnsAny { return Lns_NilAccPop().(*Nodes_MRetExp).FP.Get_index()}))
            self.FP.Write(Lns_getVM().String_format("string.format( %s, ", []LnsAny{txt}))
            for _index, _val := range( expList.FP.Get_expList().Items ) {
                index := _index + 1
                val := _val.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.Write(", ")
                }
                var matchFlag LnsInt
                matchFlag = TransUnit_FormType__Match
                if index <= opList.Len(){
                    matchFlag = convLua_convExp16544(Lns_2DDD(TransUnit_isMatchStringFormatType(opList.GetAt(index).(string), val.FP.Get_expType(), self.targetLuaVer)))
                    
                }
                if matchFlag == TransUnit_FormType__NeedConv{
                    self.FP.Write("tostring( ")
                    convLua_filter_1163_(val, self, &node.Nodes_Node)
                    self.FP.Write(")")
                } else { 
                    convLua_filter_1163_(val, self, &node.Nodes_Node)
                }
                if index == mRetIndex{
                    break
                }
            }
            self.FP.Write(")")
        } else {
            self.FP.Write(txt)
        }
    }
}

// 3324: decl @lune.@base.@convLua.convFilter.processLiteralBool
func (self *convLua_convFilter) ProcessLiteralBool(node *Nodes_LiteralBoolNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 3330: decl @lune.@base.@convLua.convFilter.processLiteralNil
func (self *convLua_convFilter) ProcessLiteralNil(node *Nodes_LiteralNilNode,_opt LnsAny) {
    self.FP.Write("nil")
}

// 3336: decl @lune.@base.@convLua.convFilter.processBreak
func (self *convLua_convFilter) ProcessBreak(node *Nodes_BreakNode,_opt LnsAny) {
    self.FP.Write("break")
}

// 3342: decl @lune.@base.@convLua.convFilter.processLiteralSymbol
func (self *convLua_convFilter) ProcessLiteralSymbol(node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{node.FP.Get_token().Txt}))
}

// 3348: decl @lune.@base.@convLua.convFilter.processLuneControl
func (self *convLua_convFilter) ProcessLuneControl(node *Nodes_LuneControlNode,_opt LnsAny) {
    switch node.FP.Get_pragma().(type) {
    case *LuneControl_Pragma__load__lune_module:
        self.FP.processLoadRuntime()
    }
}


// declaration Class -- MacroEvalImp
type ConvLua_MacroEvalImpMtd interface {
    Eval(arg1 *Ast_ProcessInfo, arg2 *Nodes_DeclMacroNode) *Lns_luaValue
    EvalFromCode(arg1 *Ast_ProcessInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) *Lns_luaValue
    evalFromMacroCode(arg1 string) *Lns_luaValue
}
type ConvLua_MacroEvalImp struct {
    Nodes_MacroEval
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
func NewConvLua_MacroEvalImp() *ConvLua_MacroEvalImp {
    obj := &ConvLua_MacroEvalImp{}
    obj.FP = obj
    obj.Nodes_MacroEval.FP = obj
    obj.InitConvLua_MacroEvalImp()
    return obj
}
func (self *ConvLua_MacroEvalImp) InitConvLua_MacroEvalImp() {
    self.Nodes_MacroEval.InitNodes_MacroEval( )
}
// 3383: decl @lune.@base.@convLua.MacroEvalImp.evalFromMacroCode
func (self *ConvLua_MacroEvalImp) evalFromMacroCode(code string) *Lns_luaValue {
    __func__ := "@lune.@base.@convLua.MacroEvalImp.evalFromMacroCode"
    Log_log(Log_Level__Trace, __func__, 3385, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("macro: %s", []LnsAny{code})
    }))
    
    var _func LnsAny
    var err string
    _func,err = convLua_runLuaOnLns_1650_(code)
    if _func != nil{
        func_7090 := _func
        return func_7090.(*Lns_luaValue)
    }
    Util_err(err)
// insert a dummy
    return nil
}

// 3394: decl @lune.@base.@convLua.MacroEvalImp.evalFromCode
func (self *ConvLua_MacroEvalImp) EvalFromCode(processInfo *Ast_ProcessInfo,name string,argNameList *LnsList,code LnsAny) *Lns_luaValue {
    var stream *Util_memStream
    stream = NewUtil_memStream()
    var conv *convLua_convFilter
    conv = NewconvLua_convFilter("macro", stream.FP, stream.FP, ConvLua_ConvMode__Exec, true, Ast_headTypeInfo, processInfo, Ast_SymbolKind__Typ, nil, LuaVer_getCurVer(), false, true)
    conv.FP.OutputDeclMacro(name, argNameList, convLua_outputMacroStmtBlock_1378_(func() {
        if code != nil{
            code_7104 := code.(string)
            conv.FP.Write(code_7104)
        }
    }))
    return self.FP.evalFromMacroCode(stream.FP.Get_txt())
}

// 3415: decl @lune.@base.@convLua.MacroEvalImp.eval
func (self *ConvLua_MacroEvalImp) Eval(processInfo *Ast_ProcessInfo,node *Nodes_DeclMacroNode) *Lns_luaValue {
    var stream *Util_memStream
    stream = NewUtil_memStream()
    var conv *convLua_convFilter
    conv = NewconvLua_convFilter("macro", stream.FP, stream.FP, ConvLua_ConvMode__Exec, true, Ast_headTypeInfo, processInfo, Ast_SymbolKind__Typ, nil, LuaVer_getCurVer(), false, true)
    conv.FP.ProcessDeclMacro(node, NewConvLua_Opt(&node.Nodes_Node))
    return self.FP.evalFromMacroCode(stream.FP.Get_txt())
}


func Lns_convLua_init() {
    if init_convLua { return }
    init_convLua = true
    convLua__mod__ = "@lune.@base.@convLua"
    Lns_InitMod()
    Lns_Ver_init()
    Lns_Str_init()
    Lns_Ast_init()
    Lns_Nodes_init()
    Lns_Util_init()
    Lns_TransUnit_init()
    Lns_LuaMod_init()
    Lns_LuaVer_init()
    Lns_Parser_init()
    Lns_Log_init()
    Lns_LuneControl_init()
    Lns_Option_init()
    Lns_DependLuaOnLns_init()
    convLua_stepIndent = 3
    
    
    
    Lns_frontInterface_init()
    Lns_Testing_init()
}
func init() {
    init_convLua = false
}
