// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_glueFilter bool
var glueFilter__mod__ string
// for 402
func glueFilter_convExp0_1518(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 112
func glueFilter_convExp0_532(arg1 []LnsAny) (string, string, *Ast_TypeInfo, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*Ast_TypeInfo), Lns_getFromMulti( arg1, 3 ).(string)
}
// for 247
func glueFilter_convExp0_809(arg1 []LnsAny) (string, string, *Ast_TypeInfo, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*Ast_TypeInfo), Lns_getFromMulti( arg1, 3 ).(string)
}
// for 334
func glueFilter_convExp0_1245(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 410
func glueFilter_convExp0_1558(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 70: decl @lune.@base.@glueFilter.getDeclFuncInfo
func glueFilter_getDeclFuncInfo_1_(_env *LnsEnv, node *Nodes_Node) *Nodes_DeclFuncInfo {
    {
        _work := Nodes_DeclConstrNodeDownCastF(node.FP)
        if !Lns_IsNil( _work ) {
            work := _work.(*Nodes_DeclConstrNode)
            return work.FP.Get_declInfo(_env)
        }
    }
    {
        _work := Nodes_DeclDestrNodeDownCastF(node.FP)
        if !Lns_IsNil( _work ) {
            work := _work.(*Nodes_DeclDestrNode)
            return work.FP.Get_declInfo(_env)
        }
    }
    {
        _work := Nodes_DeclMethodNodeDownCastF(node.FP)
        if !Lns_IsNil( _work ) {
            work := _work.(*Nodes_DeclMethodNode)
            return work.FP.Get_declInfo(_env)
        }
    }
    {
        _work := Nodes_DeclFuncNodeDownCastF(node.FP)
        if !Lns_IsNil( _work ) {
            work := _work.(*Nodes_DeclFuncNode)
            return work.FP.Get_declInfo(_env)
        }
    }
    Util_err(_env, "failed to get DeclFuncInfo")
// insert a dummy
    return nil
}

// 88: decl @lune.@base.@glueFilter.getFuncName
func glueFilter_getFuncName_2_(_env *LnsEnv, name string) string {
    if name == "__free"{
        return "__gc"
    }
    return name
}

// 394: decl @lune.@base.@glueFilter.createFilter
func GlueFilter_createFilter(_env *LnsEnv, outputDir LnsAny) *Nodes_Filter {
    return &NewglueFilter_glueFilter(_env, outputDir).Nodes_Filter
}


// 38: decl @lune.@base.@glueFilter.glueGenerator.write
func (self *glueFilter_glueGenerator) write(_env *LnsEnv, txt string) {
    self.srcStream.Write(_env, txt)
}
// 42: decl @lune.@base.@glueFilter.glueGenerator.writeHeader
func (self *glueFilter_glueGenerator) writeHeader(_env *LnsEnv, txt string) {
    self.headerStream.Write(_env, txt)
}
// 46: decl @lune.@base.@glueFilter.glueGenerator.getArgInfo
func (self *glueFilter_glueGenerator) getArgInfo(_env *LnsEnv, argNode *Nodes_Node)(string, string, *Ast_TypeInfo, string) {
    var argType *Ast_TypeInfo
    argType = argNode.FP.Get_expType(_env)
    var orgType *Ast_TypeInfo
    orgType = (_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( argType.FP.Get_nilable(_env)) &&
        _env.SetStackVal( argType.FP.Get_nonnilableType(_env)) ||
        _env.SetStackVal( argType) ).(*Ast_TypeInfo)).FP.Get_srcTypeInfo(_env)
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
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_DeclArgNode)
            argName = _exp.FP.Get_name(_env).Txt
        }
    }
    return typeTxt, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( argType.FP.Get_nilable(_env)) &&
        _env.SetStackVal( nilableTypeTxt) ||
        _env.SetStackVal( typeTxt) ).(string), orgType, argName
}
// 95: decl @lune.@base.@glueFilter.glueGenerator.outputPrototype
func (self *glueFilter_glueGenerator) outputPrototype(_env *LnsEnv, node *Nodes_Node) {
    var name string
    name = glueFilter_getFuncName_2_(_env, node.FP.Get_expType(_env).FP.Get_rawTxt(_env))
    self.FP.write(_env, _env.GetVM().String_format("static int lns_glue_%s( lua_State * pLua )", Lns_2DDD(name)))
}
// 102: decl @lune.@base.@glueFilter.glueGenerator.outputUserPrototype
func (self *glueFilter_glueGenerator) outputUserPrototype(_env *LnsEnv, node *Nodes_Node,gluePrefix string) {
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType(_env)
    self.FP.writeHeader(_env, _env.GetVM().String_format("extern int %s%s( lua_State * pLua", Lns_2DDD(gluePrefix, glueFilter_getFuncName_2_(_env, expType.FP.Get_rawTxt(_env)))))
    var declInfo *Nodes_DeclFuncInfo
    declInfo = glueFilter_getDeclFuncInfo_1_(_env, node)
    for _, _argNode := range( declInfo.FP.Get_argList(_env).Items ) {
        argNode := _argNode
        var typeTxt string
        var argTypeTxt string
        var argType *Ast_TypeInfo
        var argName string
        typeTxt,argTypeTxt,argType,argName = self.FP.getArgInfo(_env, argNode)
        if typeTxt != ""{
            self.FP.writeHeader(_env, _env.GetVM().String_format(", %s %s", Lns_2DDD(argTypeTxt, argName)))
            if argType == Ast_builtinTypeString{
                self.FP.writeHeader(_env, _env.GetVM().String_format(", int size_%s", Lns_2DDD(argName)))
            }
        }
    }
    self.FP.writeHeader(_env, " )")
}
// 125: decl @lune.@base.@glueFilter.glueGenerator.outputPrototypeList
func (self *glueFilter_glueGenerator) outputPrototypeList(_env *LnsEnv, methodNodeList *LnsList2_[*Nodes_Node]) {
    for _, _node := range( methodNodeList.Items ) {
        node := _node
        self.FP.outputPrototype(_env, node)
        self.FP.write(_env, ";\n")
    }
}
// 133: decl @lune.@base.@glueFilter.glueGenerator.outputUserPrototypeList
func (self *glueFilter_glueGenerator) outputUserPrototypeList(_env *LnsEnv, methodNodeList *LnsList2_[*Nodes_Node],gluePrefix string) {
    for _, _node := range( methodNodeList.Items ) {
        node := _node
        self.FP.outputUserPrototype(_env, node, gluePrefix)
        self.FP.writeHeader(_env, ";\n")
    }
}
// 144: decl @lune.@base.@glueFilter.glueGenerator.outputFuncReg
func (self *glueFilter_glueGenerator) outputFuncReg(_env *LnsEnv, symbolName string,methodNodeList *LnsList2_[*Nodes_Node]) {
    self.FP.write(_env, _env.GetVM().String_format("static const luaL_Reg %s[] = {\n", Lns_2DDD(symbolName)))
    for _, _node := range( methodNodeList.Items ) {
        node := _node
        {
            _nameToken := glueFilter_getDeclFuncInfo_1_(_env, node).FP.Get_name(_env)
            if !Lns_IsNil( _nameToken ) {
                nameToken := _nameToken.(*Types_Token)
                var name string
                name = glueFilter_getFuncName_2_(_env, nameToken.Txt)
                self.FP.write(_env, _env.GetVM().String_format("  { \"%s\", lns_glue_%s },\n", Lns_2DDD(name, name)))
            }
        }
    }
    self.FP.write(_env, "  { NULL, NULL }\n};\n")
}
// 157: decl @lune.@base.@glueFilter.glueGenerator.outputCommonFunc
func (self *glueFilter_glueGenerator) outputCommonFunc(_env *LnsEnv, moduleSymbolFull string) {
    self.FP.writeHeader(_env, _env.GetVM().String_format("extern int luaopen_%s( lua_State * pLua );\nextern void * lns_glue_get_%s( lua_State * pLua, int index );\nextern void * lns_glue_new_%s( lua_State * pLua, size_t size );\n", Lns_2DDD(moduleSymbolFull, moduleSymbolFull, moduleSymbolFull)))
    self.FP.write(_env, _env.GetVM().String_format("void * lns_glue_get_%s( lua_State * pLua, int index )\n{\n    return luaL_checkudata( pLua, index, s_full_class_name);\n}\n\nstatic void lns_glue_setupObjMethod(\n    lua_State * pLua, const char * pName, const luaL_Reg * pReg )\n{\n    luaL_newmetatable(pLua, pName );\n    lua_pushvalue(pLua, -1);\n    lua_setfield(pLua, -2, \"__index\");\n\n#if LUA_VERSION_NUM >= 502\n    luaL_setfuncs(pLua, pReg, 0);\n\n    lua_pop(pLua, 1);\n#else\n    luaL_register(pLua, NULL, pReg );\n\n    lua_pop(pLua, 1);\n#endif\n}\n\nvoid * lns_glue_new_%s( lua_State * pLua, size_t size )\n{\n    void * pBuf = lua_newuserdata( pLua, size );\n    if ( pBuf == NULL ) {\n        return NULL;\n    }\n    \n#if LUA_VERSION_NUM >= 502\n    luaL_setmetatable( pLua, s_full_class_name );\n#else\n    luaL_getmetatable( pLua, s_full_class_name );\n    lua_setmetatable( pLua, -2 );\n#endif\n\n    return pBuf;\n}\n\nint luaopen_%s( lua_State * pLua )\n{\n    lns_glue_setupObjMethod( pLua, s_full_class_name, s_lua_method_info );\n\n#if LUA_VERSION_NUM >= 502\n    luaL_newlib( pLua, s_lua_func_info );\n#else\n    luaL_register( pLua, s_full_class_name, s_lua_func_info );\n#endif\n    return 1;\n}\n", Lns_2DDD(moduleSymbolFull, moduleSymbolFull, moduleSymbolFull)))
}
// 230: decl @lune.@base.@glueFilter.glueGenerator.outputMethod
func (self *glueFilter_glueGenerator) outputMethod(_env *LnsEnv, node *Nodes_Node,gluePrefix string) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = glueFilter_getDeclFuncInfo_1_(_env, node)
    var name string
    {
        __exp := declInfo.FP.Get_name(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Types_Token)
            name = gluePrefix + glueFilter_getFuncName_2_(_env, _exp.Txt)
        } else {
            return 
        }
    }
    self.FP.outputPrototype(_env, node)
    self.FP.write(_env, "{\n")
    var glueArgInfoList *LnsList2_[*glueFilter_GlueArgInfo]
    glueArgInfoList = NewLnsList2_[*glueFilter_GlueArgInfo]([]*glueFilter_GlueArgInfo{})
    for _index, _argNode := range( declInfo.FP.Get_argList(_env).Items ) {
        index := _index + 1
        argNode := _argNode
        var typeTxt string
        var argTypeTxt string
        var argType *Ast_TypeInfo
        var argName string
        typeTxt,argTypeTxt,argType,argName = self.FP.getArgInfo(_env, argNode)
        if typeTxt != ""{
            var addVal LnsInt
            addVal = _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( declInfo.FP.Get_staticFlag(_env)) &&
                _env.SetStackVal( 0) ||
                _env.SetStackVal( 1) ).(LnsInt)
            var callArgName string
            callArgName = argName
            self.FP.write(_env, _env.GetVM().String_format("  %s %s = ", Lns_2DDD(typeTxt, argName)))
            if _switch0 := argType; _switch0 == Ast_builtinTypeInt {
                self.FP.write(_env, "0;\n")
            } else if _switch0 == Ast_builtinTypeReal {
                self.FP.write(_env, "0.0;\n")
            } else if _switch0 == Ast_builtinTypeString {
                self.FP.write(_env, "NULL;\n")
            }
            if argNode.FP.Get_expType(_env).FP.Get_nilable(_env){
                if argType != Ast_builtinTypeString{
                    callArgName = _env.GetVM().String_format("_p%s", Lns_2DDD(argName))
                    self.FP.write(_env, _env.GetVM().String_format("  %s %s = NULL;\n", Lns_2DDD(argTypeTxt, callArgName)))
                }
            }
            var setTxt string
            setTxt = ""
            var callTxt string
            callTxt = ""
            if _switch1 := argType; _switch1 == Ast_builtinTypeInt {
                setTxt = _env.GetVM().String_format("  %s = luaL_checkinteger( pLua, %d );\n", Lns_2DDD(argName, index + addVal))
                callTxt = callArgName
            } else if _switch1 == Ast_builtinTypeReal {
                setTxt = _env.GetVM().String_format("  %s = luaL_checknumber( pLua, %d );\n", Lns_2DDD(argName, index + addVal))
                callTxt = callArgName
            } else if _switch1 == Ast_builtinTypeString {
                self.FP.write(_env, _env.GetVM().String_format("  size_t size_%s = 0;\n", Lns_2DDD(argName)))
                setTxt = _env.GetVM().String_format("  %s = luaL_checklstring( pLua, %d, &size_%s );\n", Lns_2DDD(argName, index + addVal, argName))
                callTxt = _env.GetVM().String_format("%s, size_%s", Lns_2DDD(argName, argName))
            }
            glueArgInfoList.Insert(NewglueFilter_GlueArgInfo(_env, index + addVal, argName, callArgName, callTxt, setTxt, argNode.FP.Get_expType(_env)))
        }
    }
    for _, _glueArgInfo := range( glueArgInfoList.Items ) {
        glueArgInfo := _glueArgInfo
        if glueArgInfo.FP.Get_typeInfo(_env).FP.Get_nilable(_env){
            self.FP.write(_env, _env.GetVM().String_format("  if ( !lua_isnoneornil( pLua, %d ) ) {\n", Lns_2DDD(glueArgInfo.FP.Get_index(_env))))
            self.FP.write(_env, "  ")
        }
        self.FP.write(_env, glueArgInfo.FP.Get_setTxt(_env))
        if glueArgInfo.FP.Get_typeInfo(_env).FP.Get_nilable(_env){
            if glueArgInfo.FP.Get_callArgName(_env) != glueArgInfo.FP.Get_argName(_env){
                self.FP.write(_env, _env.GetVM().String_format("    %s = &%s;\n", Lns_2DDD(glueArgInfo.FP.Get_callArgName(_env), glueArgInfo.FP.Get_argName(_env))))
            }
            self.FP.write(_env, "  }\n")
        }
    }
    self.FP.write(_env, _env.GetVM().String_format("  return %s( pLua", Lns_2DDD(name)))
    for _, _glueArgInfo := range( glueArgInfoList.Items ) {
        glueArgInfo := _glueArgInfo
        self.FP.write(_env, ", ")
        self.FP.write(_env, glueArgInfo.FP.Get_callTxt(_env))
    }
    self.FP.write(_env, ");\n")
    self.FP.write(_env, "}\n")
}
// 331: decl @lune.@base.@glueFilter.glueGenerator.outputClass
func (self *glueFilter_glueGenerator) OutputClass(_env *LnsEnv, moduleFullName string,node *Nodes_DeclClassNode,gluePrefix string) {
    var moduleSymbolFull string
    moduleSymbolFull = glueFilter_convExp0_1245(Lns_2DDD(_env.GetVM().String_gsub(moduleFullName,"%.", "_")))
    var staticMethodNodeList *LnsList2_[*Nodes_Node]
    staticMethodNodeList = NewLnsList2_[*Nodes_Node]([]*Nodes_Node{})
    var methodNodeList *LnsList2_[*Nodes_Node]
    methodNodeList = NewLnsList2_[*Nodes_Node]([]*Nodes_Node{})
    for _, _fieldNode := range( node.FP.Get_fieldList(_env).Items ) {
        fieldNode := _fieldNode
        {
            _methodNode := Nodes_DeclMethodNodeDownCastF(fieldNode.FP)
            if !Lns_IsNil( _methodNode ) {
                methodNode := _methodNode.(*Nodes_DeclMethodNode)
                if methodNode.FP.Get_declInfo(_env).FP.Get_staticFlag(_env){
                    staticMethodNodeList.Insert(&methodNode.Nodes_Node)
                } else { 
                    methodNodeList.Insert(&methodNode.Nodes_Node)
                }
            } else {
                {
                    _methodNode := Nodes_DeclDestrNodeDownCastF(fieldNode.FP)
                    if !Lns_IsNil( _methodNode ) {
                        methodNode := _methodNode.(*Nodes_DeclDestrNode)
                        methodNodeList.Insert(&methodNode.Nodes_Node)
                    }
                }
            }
        }
    }
    self.FP.writeHeader(_env, "#include <lauxlib.h>\n")
    self.FP.outputUserPrototypeList(_env, staticMethodNodeList, gluePrefix)
    self.FP.outputUserPrototypeList(_env, methodNodeList, gluePrefix)
    self.FP.write(_env, _env.GetVM().String_format("#include \"%s_glue.h\"\n", Lns_2DDD(moduleSymbolFull)))
    self.FP.outputPrototypeList(_env, staticMethodNodeList)
    self.FP.outputPrototypeList(_env, methodNodeList)
    self.FP.write(_env, _env.GetVM().String_format("static const char * s_full_class_name = \"%s\";\n", Lns_2DDD(moduleFullName)))
    self.FP.outputFuncReg(_env, "s_lua_func_info", staticMethodNodeList)
    self.FP.outputFuncReg(_env, "s_lua_method_info", methodNodeList)
    self.FP.outputCommonFunc(_env, moduleSymbolFull)
    for _, _methodNode := range( methodNodeList.Items ) {
        methodNode := _methodNode
        self.FP.outputMethod(_env, methodNode, gluePrefix)
    }
    for _, _methodNode := range( staticMethodNodeList.Items ) {
        methodNode := _methodNode
        self.FP.outputMethod(_env, methodNode, gluePrefix)
    }
}
// 398: decl @lune.@base.@glueFilter.glueFilter.processRoot
func (self *glueFilter_glueFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_dummy LnsAny) {
    var glueFilter_createFile func(_env *LnsEnv, filename string) Lns_oStream
    glueFilter_createFile = func(_env *LnsEnv, filename string) Lns_oStream {
        var filePath string
        filePath = _env.GetVM().String_format("%s/%s", Lns_2DDD(Lns_unwrapDefault( self.outputDir, ".").(string), filename))
        {
            __exp := glueFilter_convExp0_1518(Lns_2DDD(Lns_io_open(filePath, "w")))
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(Lns_luaStream)
                return _exp
            }
        }
        Util_err(_env, _env.GetVM().String_format("open error -- %s ", Lns_2DDD(filePath)))
    // insert a dummy
        return nil
    }
    for _, _declClassNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclClassNodeList(_env).Items ) {
        declClassNode := _declClassNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        {
            _moduleName := declClassNode.FP.Get_moduleName(_env)
            if !Lns_IsNil( _moduleName ) {
                moduleName := _moduleName.(*Types_Token)
                var moduleSymbolName string
                moduleSymbolName = glueFilter_convExp0_1558(Lns_2DDD(_env.GetVM().String_gsub(moduleName.FP.GetExcludedDelimitTxt(_env),"%.", "_")))
                {
                    __exp := declClassNode.FP.Get_gluePrefix(_env)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(string)
                        var cFile Lns_oStream
                        cFile = glueFilter_createFile(_env, moduleSymbolName + "_glue.c")
                        var hFile Lns_oStream
                        hFile = glueFilter_createFile(_env, moduleSymbolName + "_glue.h")
                        var glue *glueFilter_glueGenerator
                        glue = NewglueFilter_glueGenerator(_env, cFile, hFile)
                        glue.FP.OutputClass(_env, moduleSymbolName, declClassNode, _exp)
                        cFile.Close(_env)
                        hFile.Close(_env)
                    }
                }
            }
        }
    }
}
// declaration Class -- glueGenerator
type glueFilter_glueGeneratorMtd interface {
    getArgInfo(_env *LnsEnv, arg1 *Nodes_Node)(string, string, *Ast_TypeInfo, string)
    OutputClass(_env *LnsEnv, arg1 string, arg2 *Nodes_DeclClassNode, arg3 string)
    outputCommonFunc(_env *LnsEnv, arg1 string)
    outputFuncReg(_env *LnsEnv, arg1 string, arg2 *LnsList2_[*Nodes_Node])
    outputMethod(_env *LnsEnv, arg1 *Nodes_Node, arg2 string)
    outputPrototype(_env *LnsEnv, arg1 *Nodes_Node)
    outputPrototypeList(_env *LnsEnv, arg1 *LnsList2_[*Nodes_Node])
    outputUserPrototype(_env *LnsEnv, arg1 *Nodes_Node, arg2 string)
    outputUserPrototypeList(_env *LnsEnv, arg1 *LnsList2_[*Nodes_Node], arg2 string)
    write(_env *LnsEnv, arg1 string)
    writeHeader(_env *LnsEnv, arg1 string)
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
func glueFilter_glueGenerator_toSlice(slice []LnsAny) []*glueFilter_glueGenerator {
    ret := make([]*glueFilter_glueGenerator, len(slice))
    for index, val := range slice {
        ret[index] = val.(glueFilter_glueGeneratorDownCast).ToglueFilter_glueGenerator()
    }
    return ret
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
func NewglueFilter_glueGenerator(_env *LnsEnv, arg1 Lns_oStream, arg2 Lns_oStream) *glueFilter_glueGenerator {
    obj := &glueFilter_glueGenerator{}
    obj.FP = obj
    obj.InitglueFilter_glueGenerator(_env, arg1, arg2)
    return obj
}
func (self *glueFilter_glueGenerator) InitglueFilter_glueGenerator(_env *LnsEnv, arg1 Lns_oStream, arg2 Lns_oStream) {
    self.srcStream = arg1
    self.headerStream = arg2
}

// declaration Class -- GlueArgInfo
type glueFilter_GlueArgInfoMtd interface {
    Get_argName(_env *LnsEnv) string
    Get_callArgName(_env *LnsEnv) string
    Get_callTxt(_env *LnsEnv) string
    Get_index(_env *LnsEnv) LnsInt
    Get_setTxt(_env *LnsEnv) string
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
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
func glueFilter_GlueArgInfo_toSlice(slice []LnsAny) []*glueFilter_GlueArgInfo {
    ret := make([]*glueFilter_GlueArgInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(glueFilter_GlueArgInfoDownCast).ToglueFilter_GlueArgInfo()
    }
    return ret
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
func NewglueFilter_GlueArgInfo(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 string, arg4 string, arg5 string, arg6 *Ast_TypeInfo) *glueFilter_GlueArgInfo {
    obj := &glueFilter_GlueArgInfo{}
    obj.FP = obj
    obj.InitglueFilter_GlueArgInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *glueFilter_GlueArgInfo) InitglueFilter_GlueArgInfo(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 string, arg4 string, arg5 string, arg6 *Ast_TypeInfo) {
    self.index = arg1
    self.argName = arg2
    self.callArgName = arg3
    self.callTxt = arg4
    self.setTxt = arg5
    self.typeInfo = arg6
}
func (self *glueFilter_GlueArgInfo) Get_index(_env *LnsEnv) LnsInt{ return self.index }
func (self *glueFilter_GlueArgInfo) Get_argName(_env *LnsEnv) string{ return self.argName }
func (self *glueFilter_GlueArgInfo) Get_callArgName(_env *LnsEnv) string{ return self.callArgName }
func (self *glueFilter_GlueArgInfo) Get_callTxt(_env *LnsEnv) string{ return self.callTxt }
func (self *glueFilter_GlueArgInfo) Get_setTxt(_env *LnsEnv) string{ return self.setTxt }
func (self *glueFilter_GlueArgInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }

// declaration Class -- glueFilter
type glueFilter_glueFilterMtd interface {
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
func glueFilter_glueFilter_toSlice(slice []LnsAny) []*glueFilter_glueFilter {
    ret := make([]*glueFilter_glueFilter, len(slice))
    for index, val := range slice {
        ret[index] = val.(glueFilter_glueFilterDownCast).ToglueFilter_glueFilter()
    }
    return ret
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
func NewglueFilter_glueFilter(_env *LnsEnv, arg1 LnsAny) *glueFilter_glueFilter {
    obj := &glueFilter_glueFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitglueFilter_glueFilter(_env, arg1)
    return obj
}
// 388: DeclConstr
func (self *glueFilter_glueFilter) InitglueFilter_glueFilter(_env *LnsEnv, outputDir LnsAny) {
    self.InitNodes_Filter(_env, false, nil, nil)
    self.outputDir = outputDir
}


func Lns_glueFilter_init(_env *LnsEnv) {
    if init_glueFilter { return }
    init_glueFilter = true
    glueFilter__mod__ = "@lune.@base.@glueFilter"
    Lns_InitMod()
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Util_init(_env)
}
func init() {
    init_glueFilter = false
}
