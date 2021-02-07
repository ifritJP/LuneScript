// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_glueFilter bool
var glueFilter__mod__ string
// for 399
func glueFilter_convExp1514(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 331
func glueFilter_convExp1186(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 407
func glueFilter_convExp1553(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 67: decl @lune.@base.@glueFilter.getDeclFuncInfo
func glueFilter_getDeclFuncInfo_1028_(node *Nodes_Node) *Nodes_DeclFuncInfo {
    {
        _work := Nodes_DeclConstrNodeDownCastF(node.FP)
        if _work != nil {
            work := _work.(*Nodes_DeclConstrNode)
            return work.FP.Get_declInfo()
        }
    }
    {
        _work := Nodes_DeclDestrNodeDownCastF(node.FP)
        if _work != nil {
            work := _work.(*Nodes_DeclDestrNode)
            return work.FP.Get_declInfo()
        }
    }
    {
        _work := Nodes_DeclMethodNodeDownCastF(node.FP)
        if _work != nil {
            work := _work.(*Nodes_DeclMethodNode)
            return work.FP.Get_declInfo()
        }
    }
    {
        _work := Nodes_DeclFuncNodeDownCastF(node.FP)
        if _work != nil {
            work := _work.(*Nodes_DeclFuncNode)
            return work.FP.Get_declInfo()
        }
    }
    Util_err("failed to get DeclFuncInfo")
// insert a dummy
    return nil
}

// 85: decl @lune.@base.@glueFilter.getFuncName
func glueFilter_getFuncName_1034_(name string) string {
    if name == "__free"{
        return "__gc"
    }
    return name
}

// 391: decl @lune.@base.@glueFilter.createFilter
func GlueFilter_createFilter(outputDir LnsAny) *Nodes_Filter {
    return &NewglueFilter_glueFilter(outputDir).Nodes_Filter
}


// declaration Class -- glueGenerator
type glueFilter_glueGeneratorMtd interface {
    getArgInfo(arg1 *Nodes_Node)(string, string, *Ast_TypeInfo, string)
    OutputClass(arg1 string, arg2 *Nodes_DeclClassNode, arg3 string)
    outputCommonFunc(arg1 string)
    outputFuncReg(arg1 string, arg2 *LnsList)
    outputMethod(arg1 *Nodes_Node, arg2 string)
    outputPrototype(arg1 *Nodes_Node)
    outputPrototypeList(arg1 *LnsList)
    outputUserPrototype(arg1 *Nodes_Node, arg2 string)
    outputUserPrototypeList(arg1 *LnsList, arg2 string)
    write(arg1 string)
    writeHeader(arg1 string)
}
type glueFilter_glueGenerator struct {
    srcStream Lns_oStream
    headerStream Lns_oStream
    FP glueFilter_glueGeneratorMtd
}
func glueFilter_glueGenerator2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*glueFilter_glueGenerator).FP
}
type glueFilter_glueGeneratorDownCast interface {
    ToglueFilter_glueGenerator() *glueFilter_glueGenerator
}
func glueFilter_glueGeneratorDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(glueFilter_glueGeneratorDownCast)
    if ok { return work.ToglueFilter_glueGenerator() }
    return nil
}
func (obj *glueFilter_glueGenerator) ToglueFilter_glueGenerator() *glueFilter_glueGenerator {
    return obj
}
func NewglueFilter_glueGenerator(arg1 Lns_oStream, arg2 Lns_oStream) *glueFilter_glueGenerator {
    obj := &glueFilter_glueGenerator{}
    obj.FP = obj
    obj.InitglueFilter_glueGenerator(arg1, arg2)
    return obj
}
func (self *glueFilter_glueGenerator) InitglueFilter_glueGenerator(arg1 Lns_oStream, arg2 Lns_oStream) {
    self.srcStream = arg1
    self.headerStream = arg2
}
// 35: decl @lune.@base.@glueFilter.glueGenerator.write
func (self *glueFilter_glueGenerator) write(txt string) {
    self.srcStream.Write(txt)
}

// 39: decl @lune.@base.@glueFilter.glueGenerator.writeHeader
func (self *glueFilter_glueGenerator) writeHeader(txt string) {
    self.headerStream.Write(txt)
}

// 43: decl @lune.@base.@glueFilter.glueGenerator.getArgInfo
func (self *glueFilter_glueGenerator) getArgInfo(argNode *Nodes_Node)(string, string, *Ast_TypeInfo, string) {
    var argType *Ast_TypeInfo
    argType = argNode.FP.Get_expType()
    var orgType *Ast_TypeInfo
    orgType = (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( argType.FP.Get_nilable()) &&
        Lns_GetEnv().SetStackVal( argType.FP.Get_nonnilableType()) ||
        Lns_GetEnv().SetStackVal( argType) ).(*Ast_TypeInfo)).FP.Get_srcTypeInfo()
    var typeTxt string
    typeTxt = ""
    var nilableTypeTxt string
    nilableTypeTxt = ""
    if orgType == Ast_builtinTypeInt{
        typeTxt = "int"
        
        nilableTypeTxt = typeTxt + " *"
        
    } else if orgType == Ast_builtinTypeReal{
        typeTxt = "double"
        
        nilableTypeTxt = typeTxt + " *"
        
    } else if orgType == Ast_builtinTypeString{
        typeTxt = "const char *"
        
        nilableTypeTxt = typeTxt
        
    }
    var argName string
    argName = ""
    {
        __exp := Nodes_DeclArgNodeDownCastF(argNode.FP)
        if __exp != nil {
            _exp := __exp.(*Nodes_DeclArgNode)
            argName = _exp.FP.Get_name().Txt
            
        }
    }
    return typeTxt, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( argType.FP.Get_nilable()) &&
        Lns_GetEnv().SetStackVal( nilableTypeTxt) ||
        Lns_GetEnv().SetStackVal( typeTxt) ).(string), orgType, argName
}

// 92: decl @lune.@base.@glueFilter.glueGenerator.outputPrototype
func (self *glueFilter_glueGenerator) outputPrototype(node *Nodes_Node) {
    var name string
    name = glueFilter_getFuncName_1034_(node.FP.Get_expType().FP.Get_rawTxt())
    self.FP.write(Lns_getVM().String_format("static int lns_glue_%s( lua_State * pLua )", []LnsAny{name}))
}

// 99: decl @lune.@base.@glueFilter.glueGenerator.outputUserPrototype
func (self *glueFilter_glueGenerator) outputUserPrototype(node *Nodes_Node,gluePrefix string) {
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType()
    self.FP.writeHeader(Lns_getVM().String_format("extern int %s%s( lua_State * pLua", []LnsAny{gluePrefix, glueFilter_getFuncName_1034_(expType.FP.Get_rawTxt())}))
    var declInfo *Nodes_DeclFuncInfo
    declInfo = glueFilter_getDeclFuncInfo_1028_(node)
    for _, _argNode := range( declInfo.FP.Get_argList().Items ) {
        argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
        var typeTxt string
        var argTypeTxt string
        var argType *Ast_TypeInfo
        var argName string
        typeTxt,argTypeTxt,argType,argName = self.FP.getArgInfo(argNode)
        if typeTxt != ""{
            self.FP.writeHeader(Lns_getVM().String_format(", %s %s", []LnsAny{argTypeTxt, argName}))
            if argType == Ast_builtinTypeString{
                self.FP.writeHeader(Lns_getVM().String_format(", int size_%s", []LnsAny{argName}))
            }
        }
    }
    self.FP.writeHeader(" )")
}

// 122: decl @lune.@base.@glueFilter.glueGenerator.outputPrototypeList
func (self *glueFilter_glueGenerator) outputPrototypeList(methodNodeList *LnsList) {
    for _, _node := range( methodNodeList.Items ) {
        node := _node.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.outputPrototype(node)
        self.FP.write(";\n")
    }
}

// 130: decl @lune.@base.@glueFilter.glueGenerator.outputUserPrototypeList
func (self *glueFilter_glueGenerator) outputUserPrototypeList(methodNodeList *LnsList,gluePrefix string) {
    for _, _node := range( methodNodeList.Items ) {
        node := _node.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.outputUserPrototype(node, gluePrefix)
        self.FP.writeHeader(";\n")
    }
}

// 141: decl @lune.@base.@glueFilter.glueGenerator.outputFuncReg
func (self *glueFilter_glueGenerator) outputFuncReg(symbolName string,methodNodeList *LnsList) {
    self.FP.write(Lns_getVM().String_format("static const luaL_Reg %s[] = {\n", []LnsAny{symbolName}))
    for _, _node := range( methodNodeList.Items ) {
        node := _node.(Nodes_NodeDownCast).ToNodes_Node()
        {
            _nameToken := glueFilter_getDeclFuncInfo_1028_(node).FP.Get_name()
            if _nameToken != nil {
                nameToken := _nameToken.(*Types_Token)
                var name string
                name = glueFilter_getFuncName_1034_(nameToken.Txt)
                self.FP.write(Lns_getVM().String_format("  { \"%s\", lns_glue_%s },\n", []LnsAny{name, name}))
            }
        }
    }
    self.FP.write("  { NULL, NULL }\n};\n")
}

// 154: decl @lune.@base.@glueFilter.glueGenerator.outputCommonFunc
func (self *glueFilter_glueGenerator) outputCommonFunc(moduleSymbolFull string) {
    self.FP.writeHeader(Lns_getVM().String_format("extern int luaopen_%s( lua_State * pLua );\nextern void * lns_glue_get_%s( lua_State * pLua, int index );\nextern void * lns_glue_new_%s( lua_State * pLua, size_t size );\n", []LnsAny{moduleSymbolFull, moduleSymbolFull, moduleSymbolFull}))
    self.FP.write(Lns_getVM().String_format("void * lns_glue_get_%s( lua_State * pLua, int index )\n{\n    return luaL_checkudata( pLua, index, s_full_class_name);\n}\n\nstatic void lns_glue_setupObjMethod(\n    lua_State * pLua, const char * pName, const luaL_Reg * pReg )\n{\n    luaL_newmetatable(pLua, pName );\n    lua_pushvalue(pLua, -1);\n    lua_setfield(pLua, -2, \"__index\");\n\n#if LUA_VERSION_NUM >= 502\n    luaL_setfuncs(pLua, pReg, 0);\n\n    lua_pop(pLua, 1);\n#else\n    luaL_register(pLua, NULL, pReg );\n\n    lua_pop(pLua, 1);\n#endif\n}\n\nvoid * lns_glue_new_%s( lua_State * pLua, size_t size )\n{\n    void * pBuf = lua_newuserdata( pLua, size );\n    if ( pBuf == NULL ) {\n        return NULL;\n    }\n    \n#if LUA_VERSION_NUM >= 502\n    luaL_setmetatable( pLua, s_full_class_name );\n#else\n    luaL_getmetatable( pLua, s_full_class_name );\n    lua_setmetatable( pLua, -2 );\n#endif\n\n    return pBuf;\n}\n\nint luaopen_%s( lua_State * pLua )\n{\n    lns_glue_setupObjMethod( pLua, s_full_class_name, s_lua_method_info );\n\n#if LUA_VERSION_NUM >= 502\n    luaL_newlib( pLua, s_lua_func_info );\n#else\n    luaL_register( pLua, s_full_class_name, s_lua_func_info );\n#endif\n    return 1;\n}\n", []LnsAny{moduleSymbolFull, moduleSymbolFull, moduleSymbolFull}))
}

// 227: decl @lune.@base.@glueFilter.glueGenerator.outputMethod
func (self *glueFilter_glueGenerator) outputMethod(node *Nodes_Node,gluePrefix string) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = glueFilter_getDeclFuncInfo_1028_(node)
    var name string
    {
        __exp := declInfo.FP.Get_name()
        if __exp != nil {
            _exp := __exp.(*Types_Token)
            name = gluePrefix + glueFilter_getFuncName_1034_(_exp.Txt)
            
        } else {
            return 
        }
    }
    self.FP.outputPrototype(node)
    self.FP.write("{\n")
    var glueArgInfoList *LnsList
    glueArgInfoList = NewLnsList([]LnsAny{})
    for _index, _argNode := range( declInfo.FP.Get_argList().Items ) {
        index := _index + 1
        argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
        var typeTxt string
        var argTypeTxt string
        var argType *Ast_TypeInfo
        var argName string
        typeTxt,argTypeTxt,argType,argName = self.FP.getArgInfo(argNode)
        if typeTxt != ""{
            var addVal LnsInt
            addVal = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( declInfo.FP.Get_staticFlag()) &&
                Lns_GetEnv().SetStackVal( 0) ||
                Lns_GetEnv().SetStackVal( 1) ).(LnsInt)
            var callArgName string
            callArgName = argName
            self.FP.write(Lns_getVM().String_format("  %s %s = ", []LnsAny{typeTxt, argName}))
            if _switch816 := argType; _switch816 == Ast_builtinTypeInt {
                self.FP.write("0;\n")
            } else if _switch816 == Ast_builtinTypeReal {
                self.FP.write("0.0;\n")
            } else if _switch816 == Ast_builtinTypeString {
                self.FP.write("NULL;\n")
            }
            if argNode.FP.Get_expType().FP.Get_nilable(){
                if argType != Ast_builtinTypeString{
                    callArgName = Lns_getVM().String_format("_p%s", []LnsAny{argName})
                    
                    self.FP.write(Lns_getVM().String_format("  %s %s = NULL;\n", []LnsAny{argTypeTxt, callArgName}))
                }
            }
            var setTxt string
            setTxt = ""
            var callTxt string
            callTxt = ""
            if _switch987 := argType; _switch987 == Ast_builtinTypeInt {
                setTxt = Lns_getVM().String_format("  %s = luaL_checkinteger( pLua, %d );\n", []LnsAny{argName, index + addVal})
                
                callTxt = callArgName
                
            } else if _switch987 == Ast_builtinTypeReal {
                setTxt = Lns_getVM().String_format("  %s = luaL_checknumber( pLua, %d );\n", []LnsAny{argName, index + addVal})
                
                callTxt = callArgName
                
            } else if _switch987 == Ast_builtinTypeString {
                self.FP.write(Lns_getVM().String_format("  size_t size_%s = 0;\n", []LnsAny{argName}))
                setTxt = Lns_getVM().String_format("  %s = luaL_checklstring( pLua, %d, &size_%s );\n", []LnsAny{argName, index + addVal, argName})
                
                callTxt = Lns_getVM().String_format("%s, size_%s", []LnsAny{argName, argName})
                
            }
            glueArgInfoList.Insert(glueFilter_GlueArgInfo2Stem(NewglueFilter_GlueArgInfo(index + addVal, argName, callArgName, callTxt, setTxt, argNode.FP.Get_expType())))
        }
    }
    for _, _glueArgInfo := range( glueArgInfoList.Items ) {
        glueArgInfo := _glueArgInfo.(glueFilter_GlueArgInfoDownCast).ToglueFilter_GlueArgInfo()
        if glueArgInfo.FP.Get_typeInfo().FP.Get_nilable(){
            self.FP.write(Lns_getVM().String_format("  if ( !lua_isnoneornil( pLua, %d ) ) {\n", []LnsAny{glueArgInfo.FP.Get_index()}))
            self.FP.write("  ")
        }
        self.FP.write(glueArgInfo.FP.Get_setTxt())
        if glueArgInfo.FP.Get_typeInfo().FP.Get_nilable(){
            if glueArgInfo.FP.Get_callArgName() != glueArgInfo.FP.Get_argName(){
                self.FP.write(Lns_getVM().String_format("    %s = &%s;\n", []LnsAny{glueArgInfo.FP.Get_callArgName(), glueArgInfo.FP.Get_argName()}))
            }
            self.FP.write("  }\n")
        }
    }
    self.FP.write(Lns_getVM().String_format("  return %s( pLua", []LnsAny{name}))
    for _, _glueArgInfo := range( glueArgInfoList.Items ) {
        glueArgInfo := _glueArgInfo.(glueFilter_GlueArgInfoDownCast).ToglueFilter_GlueArgInfo()
        self.FP.write(", ")
        self.FP.write(glueArgInfo.FP.Get_callTxt())
    }
    self.FP.write(");\n")
    self.FP.write("}\n")
}

// 328: decl @lune.@base.@glueFilter.glueGenerator.outputClass
func (self *glueFilter_glueGenerator) OutputClass(moduleFullName string,node *Nodes_DeclClassNode,gluePrefix string) {
    var moduleSymbolFull string
    moduleSymbolFull = glueFilter_convExp1186(Lns_2DDD(Lns_getVM().String_gsub(moduleFullName,"%.", "_")))
    var staticMethodNodeList *LnsList
    staticMethodNodeList = NewLnsList([]LnsAny{})
    var methodNodeList *LnsList
    methodNodeList = NewLnsList([]LnsAny{})
    for _, _fieldNode := range( node.FP.Get_fieldList().Items ) {
        fieldNode := _fieldNode.(Nodes_NodeDownCast).ToNodes_Node()
        {
            _methodNode := Nodes_DeclMethodNodeDownCastF(fieldNode.FP)
            if _methodNode != nil {
                methodNode := _methodNode.(*Nodes_DeclMethodNode)
                if methodNode.FP.Get_declInfo().FP.Get_staticFlag(){
                    staticMethodNodeList.Insert(Nodes_DeclMethodNode2Stem(methodNode))
                } else { 
                    methodNodeList.Insert(Nodes_DeclMethodNode2Stem(methodNode))
                }
            } else {
                {
                    _methodNode := Nodes_DeclDestrNodeDownCastF(fieldNode.FP)
                    if _methodNode != nil {
                        methodNode := _methodNode.(*Nodes_DeclDestrNode)
                        methodNodeList.Insert(Nodes_DeclDestrNode2Stem(methodNode))
                    }
                }
            }
        }
    }
    self.FP.writeHeader("#include <lauxlib.h>\n")
    self.FP.outputUserPrototypeList(staticMethodNodeList, gluePrefix)
    self.FP.outputUserPrototypeList(methodNodeList, gluePrefix)
    self.FP.write(Lns_getVM().String_format("#include \"%s_glue.h\"\n", []LnsAny{moduleSymbolFull}))
    self.FP.outputPrototypeList(staticMethodNodeList)
    self.FP.outputPrototypeList(methodNodeList)
    self.FP.write(Lns_getVM().String_format("static const char * s_full_class_name = \"%s\";\n", []LnsAny{moduleFullName}))
    self.FP.outputFuncReg("s_lua_func_info", staticMethodNodeList)
    self.FP.outputFuncReg("s_lua_method_info", methodNodeList)
    self.FP.outputCommonFunc(moduleSymbolFull)
    for _, _methodNode := range( methodNodeList.Items ) {
        methodNode := _methodNode.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.outputMethod(methodNode, gluePrefix)
    }
    for _, _methodNode := range( staticMethodNodeList.Items ) {
        methodNode := _methodNode.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.outputMethod(methodNode, gluePrefix)
    }
}


// declaration Class -- GlueArgInfo
type glueFilter_GlueArgInfoMtd interface {
    Get_argName() string
    Get_callArgName() string
    Get_callTxt() string
    Get_index() LnsInt
    Get_setTxt() string
    Get_typeInfo() *Ast_TypeInfo
}
type glueFilter_GlueArgInfo struct {
    index LnsInt
    argName string
    callArgName string
    callTxt string
    setTxt string
    typeInfo *Ast_TypeInfo
    FP glueFilter_GlueArgInfoMtd
}
func glueFilter_GlueArgInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*glueFilter_GlueArgInfo).FP
}
type glueFilter_GlueArgInfoDownCast interface {
    ToglueFilter_GlueArgInfo() *glueFilter_GlueArgInfo
}
func glueFilter_GlueArgInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(glueFilter_GlueArgInfoDownCast)
    if ok { return work.ToglueFilter_GlueArgInfo() }
    return nil
}
func (obj *glueFilter_GlueArgInfo) ToglueFilter_GlueArgInfo() *glueFilter_GlueArgInfo {
    return obj
}
func NewglueFilter_GlueArgInfo(arg1 LnsInt, arg2 string, arg3 string, arg4 string, arg5 string, arg6 *Ast_TypeInfo) *glueFilter_GlueArgInfo {
    obj := &glueFilter_GlueArgInfo{}
    obj.FP = obj
    obj.InitglueFilter_GlueArgInfo(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *glueFilter_GlueArgInfo) InitglueFilter_GlueArgInfo(arg1 LnsInt, arg2 string, arg3 string, arg4 string, arg5 string, arg6 *Ast_TypeInfo) {
    self.index = arg1
    self.argName = arg2
    self.callArgName = arg3
    self.callTxt = arg4
    self.setTxt = arg5
    self.typeInfo = arg6
}
func (self *glueFilter_GlueArgInfo) Get_index() LnsInt{ return self.index }
func (self *glueFilter_GlueArgInfo) Get_argName() string{ return self.argName }
func (self *glueFilter_GlueArgInfo) Get_callArgName() string{ return self.callArgName }
func (self *glueFilter_GlueArgInfo) Get_callTxt() string{ return self.callTxt }
func (self *glueFilter_GlueArgInfo) Get_setTxt() string{ return self.setTxt }
func (self *glueFilter_GlueArgInfo) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }

// declaration Class -- glueFilter
type glueFilter_glueFilterMtd interface {
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
}
type glueFilter_glueFilter struct {
    Nodes_Filter
    outputDir LnsAny
    FP glueFilter_glueFilterMtd
}
func glueFilter_glueFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*glueFilter_glueFilter).FP
}
type glueFilter_glueFilterDownCast interface {
    ToglueFilter_glueFilter() *glueFilter_glueFilter
}
func glueFilter_glueFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(glueFilter_glueFilterDownCast)
    if ok { return work.ToglueFilter_glueFilter() }
    return nil
}
func (obj *glueFilter_glueFilter) ToglueFilter_glueFilter() *glueFilter_glueFilter {
    return obj
}
func NewglueFilter_glueFilter(arg1 LnsAny) *glueFilter_glueFilter {
    obj := &glueFilter_glueFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitglueFilter_glueFilter(arg1)
    return obj
}
// 385: DeclConstr
func (self *glueFilter_glueFilter) InitglueFilter_glueFilter(outputDir LnsAny) {
    self.InitNodes_Filter(false, nil, nil)
    self.outputDir = outputDir
    
}

// 395: decl @lune.@base.@glueFilter.glueFilter.processRoot
func (self *glueFilter_glueFilter) ProcessRoot(node *Nodes_RootNode,_dummy LnsAny) {
    var createFile func(filename string) Lns_oStream
    createFile = func(filename string) Lns_oStream {
        var filePath string
        filePath = Lns_getVM().String_format("%s/%s", []LnsAny{Lns_unwrapDefault( self.outputDir, ".").(string), filename})
        {
            __exp := glueFilter_convExp1514(Lns_2DDD(Lns_io_open(filePath, "w")))
            if __exp != nil {
                _exp := __exp.(Lns_luaStream)
                return _exp
            }
        }
        panic(Lns_getVM().String_format("open error -- %s ", []LnsAny{filePath}))
    // insert a dummy
        return nil
    }
    for _, _declClassNode := range( node.FP.Get_nodeManager().FP.GetDeclClassNodeList().Items ) {
        declClassNode := _declClassNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        {
            _moduleName := declClassNode.FP.Get_moduleName()
            if _moduleName != nil {
                moduleName := _moduleName.(*Types_Token)
                var moduleSymbolName string
                moduleSymbolName = glueFilter_convExp1553(Lns_2DDD(Lns_getVM().String_gsub(moduleName.FP.GetExcludedDelimitTxt(),"%.", "_")))
                {
                    __exp := declClassNode.FP.Get_gluePrefix()
                    if __exp != nil {
                        _exp := __exp.(string)
                        var glue *glueFilter_glueGenerator
                        glue = NewglueFilter_glueGenerator(createFile(moduleSymbolName + "_glue.c"), createFile(moduleSymbolName + "_glue.h"))
                        glue.FP.OutputClass(moduleSymbolName, declClassNode, _exp)
                    }
                }
            }
        }
    }
}


func Lns_glueFilter_init() {
    if init_glueFilter { return }
    init_glueFilter = true
    glueFilter__mod__ = "@lune.@base.@glueFilter"
    Lns_InitMod()
    Lns_Ast_init()
    Lns_Nodes_init()
    Lns_Util_init()
}
func init() {
    init_glueFilter = false
}
