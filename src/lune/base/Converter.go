// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Converter bool
var Converter__mod__ string
// decl alge -- CreateAstResult
type Converter_CreateAstResult = LnsAny
type Converter_CreateAstResult__Ast struct{
Val1 *AstInfo_ASTInfo
}
func (self *Converter_CreateAstResult__Ast) GetTxt() string {
return "CreateAstResult.Ast"
}
type Converter_CreateAstResult__Creater struct{
Val1 *Converter_AstCreater
}
func (self *Converter_CreateAstResult__Creater) GetTxt() string {
return "CreateAstResult.Creater"
}
type Converter_ConverterFunc func (_env *LnsEnv)
// for 330
func Converter_convExp0_2290(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 339
func Converter_convExp0_2361(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 76
func Converter_convExp0_1419(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 220
func Converter_convExp0_1749(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 279
func Converter_convExp0_2000(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 280
func Converter_convExp0_2013(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 335
func Converter_convExp0_2314(arg1 []LnsAny) (bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 458
func Converter_convExp0_2561(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 535
func Converter_convExp0_2691(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 604
func Converter_convExp0_2764(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 44: decl @lune.@base.@Converter.createModuleInfo
func Converter_createModuleInfo_0_(_env *LnsEnv, ast *AstInfo_ASTInfo,mod string,moduleId *FrontInterface_ModuleId) *FrontInterface_ModuleInfo {
    var exportInfo *FrontInterface_ExportInfo
    exportInfo = ast.FP.Get_exportInfo(_env)
    return NewFrontInterface_ModuleInfo(_env, exportInfo)
}

// 57: decl @lune.@base.@Converter.ast2LuaMain
func Converter_ast2LuaMain(_env *LnsEnv, ast *AstInfo_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,option *Option_Option) *ConvLua_FilterInfo {
    var exportInfo *FrontInterface_ExportInfo
    exportInfo = ast.FP.Get_exportInfo(_env)
    var conv *ConvLua_FilterInfo
    conv = ConvLua_createFilter(_env, streamName, stream, metaStream, convMode, inMacro, exportInfo.FP.Get_moduleTypeInfo(_env), exportInfo.FP.Get_processInfo(_env), exportInfo.FP.Get_provideInfo(_env).FP.Get_symbolKind(_env), ast.FP.Get_builtinFunc(_env), option.UseLuneModule, option.TargetLuaVer, option.Testing, option.UseIpairs, NewConvLua_Option(_env, option.MainModule, option.FP.Get_legacyNewName(_env)))
    return conv
}

// 73: decl @lune.@base.@Converter.byteCompileFromLuaTxt
func Converter_byteCompileFromLuaTxt(_env *LnsEnv, txt string,stripDebugInfo bool) string {
    var ret string
        var chunk LnsAny
        var err LnsAny
        chunk,err = _env.GetVM().Load(txt, nil)
        if chunk != nil{
            chunk_248 := chunk.(*Lns_luaValue)
            ret = _env.GetVM().String_dump(chunk_248, stripDebugInfo)
        } else {
            Util_err(_env, Lns_unwrapDefault( err, "load error").(string))
        }
    return ret
}



// 205: decl @lune.@base.@Converter.getAstFromResult
func Converter_getAstFromResult_6_(_env *LnsEnv, result LnsAny) *AstInfo_ASTInfo {
    switch _matchExp0 := result.(type) {
    case *Converter_CreateAstResult__Ast:
        ast := _matchExp0.Val1
        return ast
    case *Converter_CreateAstResult__Creater:
        creater := _matchExp0.Val1
        return Lns_car(creater.FP.GetAst(_env)).(*AstInfo_ASTInfo)
    }
// insert a dummy
    return nil
}

// 216: decl @lune.@base.@Converter.closeStreams
func Converter_closeStreams(_env *LnsEnv, stream LnsAny,metaStream LnsAny,dependStream LnsAny,metaPath string,saveMetaFlag bool) {
    var Converter_checkDiff func(_env *LnsEnv, oldStream *Util_TxtStream,newStream *Util_TxtStream)(bool, string)
    Converter_checkDiff = func(_env *LnsEnv, oldStream *Util_TxtStream,newStream *Util_TxtStream)(bool, string) {
        __func__ := "@lune.@base.@Converter.closeStreams.checkDiff"
        var headEndPos LnsInt
        headEndPos = 0
        var tailBeginPos LnsInt
        tailBeginPos = 0
        var oldBuildIdLine string
        oldBuildIdLine = ""
        var newBuildIdLine string
        newBuildIdLine = ""
        for  {
            var newLine LnsAny
            newLine = newStream.FP.Read(_env, "*l")
            var oldLine LnsAny
            oldLine = oldStream.FP.Read(_env, "*l")
            if oldLine != nil{
                oldLine_297 := oldLine.(string)
                if len(oldBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(oldLine_297,"^_moduleObj.__buildId", nil, nil))){
                        oldBuildIdLine = oldLine_297
                    }
                }
            }
            
            if newLine != nil{
                newLine_301 := newLine.(string)
                if len(newBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(newLine_301,"^_moduleObj.__buildId", nil, nil))){
                        newBuildIdLine = newLine_301
                    }
                }
            }
            
            if newLine != oldLine{
                var cont bool
                cont = false
                if newLine != nil && oldLine != nil{
                    newLine_307 := newLine.(string)
                    oldLine_308 := oldLine.(string)
                    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(oldLine_308,"^_moduleObj.__buildId", nil, nil))){
                        if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(newLine_307,"^_moduleObj.__buildId", nil, nil))){
                            tailBeginPos = newStream.FP.Get_lineNo(_env)
                            cont = true
                        }
                    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_car(_env.GetVM().String_find(oldLine_308,"^__dependModuleMap.*buildId =", nil, nil))) &&
                        _env.SetStackVal( Lns_car(_env.GetVM().String_find(newLine_307,"^__dependModuleMap.*buildId =", nil, nil))) )){
                        var oldSub string
                        oldSub = Converter_convExp0_2000(Lns_2DDD(_env.GetVM().String_gsub(oldLine_308,"buildId =.*", "")))
                        var newSub string
                        newSub = Converter_convExp0_2013(Lns_2DDD(_env.GetVM().String_gsub(newLine_307,"buildId =.*", "")))
                        if oldSub == newSub{
                            cont = true
                        }
                    }
                }
                if Lns_op_not(cont){
                    Log_log(_env, Log_Level__Debug, __func__, 287, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.GetVM().String_format("<%s>, <%s>", []LnsAny{oldLine, newLine})
                    }))
                    
                    return false, ""
                }
            } else { 
                if tailBeginPos == 0{
                    headEndPos = newStream.FP.Get_lineNo(_env)
                }
                if Lns_op_not(oldLine){
                    if tailBeginPos == 0{
                        return true, oldStream.FP.Get_txt(_env)
                    }
                    var oldBuildId *FrontInterface_ModuleId
                    oldBuildId = Converter_closeStreams__txt2ModuleId_0_(_env, oldBuildIdLine)
                    var newBuildId *FrontInterface_ModuleId
                    newBuildId = Converter_closeStreams__txt2ModuleId_0_(_env, newBuildIdLine)
                    var worlBuildId *FrontInterface_ModuleId
                    worlBuildId = FrontInterface_ModuleId_createId(_env, newBuildId.FP.Get_modTime(_env), oldBuildId.FP.Get_buildCount(_env))
                    var buildIdLine string
                    buildIdLine = _env.GetVM().String_format("_moduleObj.__buildId = %q", []LnsAny{worlBuildId.FP.Get_idStr(_env)})
                    var txt string
                    txt = _env.GetVM().String_format("%s%s\n%s", []LnsAny{newStream.FP.GetSubstring(_env, 1, headEndPos), buildIdLine, newStream.FP.GetSubstring(_env, tailBeginPos, nil)})
                    return true, txt
                }
            }
        }
    // insert a dummy
        return false,""
    }
    if stream != nil{
        stream_328 := stream.(Lns_oStream)
        stream_328.Close(_env)
    }
    if dependStream != nil{
        dependStream_330 := dependStream.(Lns_oStream)
        dependStream_330.Close(_env)
    }
    if metaStream != nil{
        metaStream_332 := metaStream.(*Util_memStream)
        if saveMetaFlag{
            var newMetaTxt string
            newMetaTxt = metaStream_332.FP.Get_txt(_env)
            var oldMetaTxt string
            oldMetaTxt = ""
            {
                _oldFileObj := Converter_convExp0_2290(Lns_2DDD(Lns_io_open(metaPath, nil)))
                if !Lns_IsNil( _oldFileObj ) {
                    oldFileObj := _oldFileObj.(Lns_luaStream)
                    oldMetaTxt = Lns_unwrapDefault( oldFileObj.Read(_env, "*a"), "").(string)
                    oldFileObj.Close(_env)
                }
            }
            var sameFlag bool
            var txt string
            sameFlag,txt = Converter_checkDiff(_env, NewUtil_TxtStream(_env, oldMetaTxt), NewUtil_TxtStream(_env, newMetaTxt))
            var Converter_saveMeta func(_env *LnsEnv, meta string)
            Converter_saveMeta = func(_env *LnsEnv, meta string) {
                {
                    _fileObj := Converter_convExp0_2361(Lns_2DDD(Lns_io_open(metaPath, "w")))
                    if !Lns_IsNil( _fileObj ) {
                        fileObj := _fileObj.(Lns_luaStream)
                        fileObj.Write(_env, meta)
                        fileObj.Close(_env)
                    } else {
                        Util_err(_env, _env.GetVM().String_format("failed to open -- %s", []LnsAny{metaPath}))
                    }
                }
            }
            if Lns_op_not(sameFlag){
                Converter_saveMeta(_env, newMetaTxt)
            } else { 
                if txt != ""{
                    Converter_saveMeta(_env, txt)
                }
            }
        }
    }
}

// 219: decl @lune.@base.@Converter.closeStreams.txt2ModuleId
func Converter_closeStreams__txt2ModuleId_0_(_env *LnsEnv, txt string) *FrontInterface_ModuleId {
    var buildIdTxt string
    buildIdTxt = Converter_convExp0_1749(Lns_2DDD(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(txt,"^_moduleObj.__buildId = ", "")).(string),"\"", "")))
    return FrontInterface_ModuleId_createIdFromTxt(_env, buildIdTxt)
}









// 103: decl @lune.@base.@Converter.AstCreater.createAst
func (self *Converter_AstCreater) createAst(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,baseDir LnsAny,stdinFile LnsAny,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *AstInfo_ASTInfo {
    var transUnit *TransUnit_TransUnitCtrl
    transUnit = NewTransUnit_TransUnitCtrl(_env, self.moduleId, importModuleInfo, &NewConvLua_MacroEvalImp(_env, self.builtinFunc).Nodes_MacroEval, true, analyzeModule, analyzeMode, pos, self.option.TargetLuaVer, self.option.TransCtrlInfo, self.builtinFunc)
    return transUnit.FP.CreateAST(_env, parserSrc, true, baseDir, stdinFile, false, self.mod, TransUnit_ReadyExportInfo(func(_env *LnsEnv, exportInfo *FrontInterface_ExportInfo) {
        self.exportInfo = exportInfo
        {
            __exp := self.exportInfoReadyFlag
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(G__lns_Flag)
                _exp.Set(_env)
            }
        }
    }))
}
// 169: decl @lune.@base.@Converter.AstCreater.runMain
func (self *Converter_AstCreater) RunMain(_env *LnsEnv) {
    self.converter(_env)
}
// 174: decl @lune.@base.@Converter.AstCreater.getAst
func (self *Converter_AstCreater) GetAst(_env *LnsEnv)(*AstInfo_ASTInfo, *FrontInterface_ModuleInfo, *FrontInterface_ModuleMeta) {
    LnsJoin(_env, self.FP)
    var exportInfo *FrontInterface_ExportInfo
    
    {
        _exportInfo := self.exportInfo
        if _exportInfo == nil{
            Util_err(_env, _env.GetVM().String_format("exportInfo is nil -- %s", []LnsAny{self.mod}))
        } else {
            exportInfo = _exportInfo.(*FrontInterface_ExportInfo)
        }
    }
    var moduleMeta *FrontInterface_ModuleMeta
    moduleMeta = NewFrontInterface_ModuleMeta(_env, exportInfo.FP.Get_streamName(_env), &FrontInterface_MetaOrModule__Export{exportInfo})
    return Lns_unwrap( self.ast).(*AstInfo_ASTInfo), Lns_unwrap( self.moduleInfo).(*FrontInterface_ModuleInfo), moduleMeta
}
// 186: decl @lune.@base.@Converter.AstCreater.getExportInfo
func (self *Converter_AstCreater) GetExportInfo(_env *LnsEnv) LnsAny {
    __func__ := "@lune.@base.@Converter.AstCreater.getExportInfo"
    {
        __exp := self.exportInfoReadyFlag
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(G__lns_Flag)
            _exp.Wait(_env)
        }
    }
    if Lns_op_not(self.exportInfo){
        Log_log(_env, Log_Level__Err, __func__, 191, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.GetVM().String_format("exportInfo is nil -- %s", []LnsAny{self.mod})
        }))
        
    }
    return self.exportInfo
}
// 427: decl @lune.@base.@Converter.LuaConverter.runMain
func (self *Converter_LuaConverter) RunMain(_env *LnsEnv) {
    self.converterFunc(_env)
}
// 431: decl @lune.@base.@Converter.LuaConverter.saveLua
func (self *Converter_LuaConverter) SaveLua(_env *LnsEnv) {
    LnsJoin(_env, self.FP)
    var ast *AstInfo_ASTInfo
    ast = Converter_getAstFromResult_6_(_env, self.astResult)
    _env.NilAccFin(_env.NilAccPush(self.filterInfo) && 
    Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*ConvLua_FilterInfo).FP.OutputMeta(_env, Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode))})/* 435:7 */)
    if self.byteCompile{
        self.streamMem.FP.Write(_env, Converter_byteCompileFromLuaTxt(_env, self.byteStream.FP.Get_txt(_env), self.stripDebugInfo))
        self.metaStreamMem.FP.Write(_env, Converter_byteCompileFromLuaTxt(_env, self.byteMetaStream.FP.Get_txt(_env), true))
    }
    var luaCode string
    luaCode = self.streamMem.FP.Get_txt(_env)
    var metaTxt string
    metaTxt = self.metaStreamMem.FP.Get_txt(_env)
    var dependTxt LnsAny
    var needDepends bool
    needDepends = self.option.DependsPath != nil
    if needDepends{
        dependTxt = self.dependsStreamMem.FP.Get_txt(_env)
    } else { 
        dependTxt = nil
    }
    var streamDst Lns_luaStream
    
    {
        _streamDst := Converter_convExp0_2561(Lns_2DDD(Lns_io_open(self.luaPath, "w")))
        if _streamDst == nil{
            Util_err(_env, _env.GetVM().String_format("write open error -- %s", []LnsAny{self.luaPath}))
        } else {
            streamDst = _streamDst.(Lns_luaStream)
        }
    }
    var dependsStreamDst LnsAny
    dependsStreamDst = self.option.FP.OpenDepend(_env, self.dependsPath)
    streamDst.Write(_env, luaCode)
    var metaMemStream *Util_memStream
    metaMemStream = NewUtil_memStream(_env)
    metaMemStream.FP.Write(_env, metaTxt)
    if dependsStreamDst != nil{
        dependsStreamDst_366 := dependsStreamDst.(Lns_oStream)
        dependsStreamDst_366.Write(_env, Lns_unwrap( dependTxt).(string))
    }
    Converter_closeStreams(_env, streamDst, metaMemStream, dependsStreamDst, self.metaPath, self.option.Mode == Option_ModeKind__SaveMeta)
}
// 523: decl @lune.@base.@Converter.GoConverter.runMain
func (self *Converter_GoConverter) RunMain(_env *LnsEnv) {
    self.converter(_env)
}
// 527: decl @lune.@base.@Converter.GoConverter.saveGo
func (self *Converter_GoConverter) SaveGo(_env *LnsEnv) {
    LnsJoin(_env, self.FP)
    if Lns_op_not(self.validFlag){
        return 
    }
    var file Lns_luaStream
    
    {
        _file := Converter_convExp0_2691(Lns_2DDD(Lns_io_open(self.path, "w")))
        if _file == nil{
            Util_err(_env, _env.GetVM().String_format("can't open file -- %s", []LnsAny{self.path}))
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    file.Write(_env, self.memStream.FP.Get_txt(_env))
    file.Close(_env)
}
// 592: decl @lune.@base.@Converter.PythonConverter.runMain
func (self *Converter_PythonConverter) RunMain(_env *LnsEnv) {
    self.converter(_env)
}
// 596: decl @lune.@base.@Converter.PythonConverter.savePython
func (self *Converter_PythonConverter) SavePython(_env *LnsEnv) {
    LnsJoin(_env, self.FP)
    if Lns_op_not(self.validFlag){
        return 
    }
    var file Lns_luaStream
    
    {
        _file := Converter_convExp0_2764(Lns_2DDD(Lns_io_open(self.path, "w")))
        if _file == nil{
            Util_err(_env, _env.GetVM().String_format("can't open file -- %s", []LnsAny{self.path}))
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    file.Write(_env, self.memStream.FP.Get_txt(_env))
    file.Close(_env)
}
// declaration Class -- AstCreater
type Converter_AstCreaterMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    createAst(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 LnsInt, arg7 LnsAny) *AstInfo_ASTInfo
    GetAst(_env *LnsEnv)(*AstInfo_ASTInfo, *FrontInterface_ModuleInfo, *FrontInterface_ModuleMeta)
    GetExportInfo(_env *LnsEnv) LnsAny
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    Start(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny) bool
}
type Converter_AstCreater struct {
    Runner_Runner
    option *Option_Option
    builtinFunc *Builtin_BuiltinFuncType
    mod string
    moduleId *FrontInterface_ModuleId
    ast LnsAny
    moduleInfo LnsAny
    converter Converter_ConverterFunc
    exportInfo LnsAny
    exportInfoReadyFlag LnsAny
    FP Converter_AstCreaterMtd
}
func Converter_AstCreater2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Converter_AstCreater).FP
}
      
func Converter_AstCreater_toSlice_Runner_Runner(slice []LnsAny) []*Runner_Runner {
   ret := make([]*Runner_Runner, len(slice))
   for index, val := range slice {
      ret[index] = &val.(Converter_AstCreaterDownCast).ToConverter_AstCreater().Runner_Runner
   }
   return ret
}
      
func Converter_AstCreater_toSlice__IF[T any](slice []LnsAny) []T {
   ret := make([]T, len(slice))
   for index, val := range slice {
      ret[index] = val.(Converter_AstCreaterDownCast).ToConverter_AstCreater().FP.(T)
   }
   return ret
}
type Converter_AstCreaterDownCast interface {
    ToConverter_AstCreater() *Converter_AstCreater
}
func Converter_AstCreaterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Converter_AstCreaterDownCast)
    if ok { return work.ToConverter_AstCreater() }
    return nil
}
func (obj *Converter_AstCreater) ToConverter_AstCreater() *Converter_AstCreater {
    return obj
}
func NewConverter_AstCreater(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 LnsAny, arg5 *FrontInterface_ModuleId, arg6 LnsAny, arg7 LnsInt, arg8 LnsAny, arg9 *Builtin_BuiltinFuncType, arg10 *Option_Option) *Converter_AstCreater {
    obj := &Converter_AstCreater{}
    obj.FP = obj
    obj.Runner_Runner.FP = obj
    obj.InitConverter_AstCreater(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
// 126: DeclConstr
func (self *Converter_AstCreater) InitConverter_AstCreater(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,mod string,baseDir LnsAny,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny,builtinFunc *Builtin_BuiltinFuncType,option *Option_Option) {
    self.InitRunner_Runner(_env)
    self.option = option
    self.builtinFunc = builtinFunc
    self.mod = mod
    self.moduleId = moduleId
    self.moduleInfo = nil
    self.exportInfo = nil
    self.ast = nil
    self.exportInfoReadyFlag = LnsCreateSyncFlag(_env)
    self.converter = Converter_ConverterFunc(func(_env *LnsEnv) {
        __func__ := "@lune.@base.@Converter.AstCreater.__init.<anonymous>"
        var ast *AstInfo_ASTInfo
        ast = self.FP.createAst(_env, importModuleInfo, parserSrc, baseDir, option.FP.Get_stdinFile(_env), analyzeModule, analyzeMode, pos)
        self.ast = ast
        self.moduleInfo = Converter_createModuleInfo_0_(_env, ast, self.mod, self.moduleId)
        Log_log(_env, Log_Level__Log, __func__, 149, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.GetVM().String_format("generated AST -- %s", []LnsAny{mod})
        }))
        
    })
    var lnsPath string
    switch _matchExp0 := parserSrc.(type) {
    case *Types_ParserSrc__LnsCode:
        path := _matchExp0.Val2
        lnsPath = path
    case *Types_ParserSrc__LnsPath:
        path := _matchExp0.Val2
        lnsPath = path
    case *Types_ParserSrc__Parser:
        path := _matchExp0.Val1
        lnsPath = path
    }
    self.FP.Start(_env, 0, _env.GetVM().String_format("createAst - %s", []LnsAny{lnsPath}))
}


// declaration Class -- LuaConverter
type Converter_LuaConverterMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    SaveLua(_env *LnsEnv)
    Start(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny) bool
}
type Converter_LuaConverter struct {
    Runner_Runner
    luaPath string
    dependsPath LnsAny
    metaPath string
    option *Option_Option
    byteCompile bool
    stripDebugInfo bool
    byteStream *Util_memStream
    byteMetaStream *Util_memStream
    streamMem *Util_memStream
    metaStreamMem *Util_memStream
    dependsStreamMem *Util_memStream
    astResult LnsAny
    converterFunc Converter_ConverterFunc
    filterInfo LnsAny
    FP Converter_LuaConverterMtd
}
func Converter_LuaConverter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Converter_LuaConverter).FP
}
      
func Converter_LuaConverter_toSlice_Runner_Runner(slice []LnsAny) []*Runner_Runner {
   ret := make([]*Runner_Runner, len(slice))
   for index, val := range slice {
      ret[index] = &val.(Converter_LuaConverterDownCast).ToConverter_LuaConverter().Runner_Runner
   }
   return ret
}
      
func Converter_LuaConverter_toSlice__IF[T any](slice []LnsAny) []T {
   ret := make([]T, len(slice))
   for index, val := range slice {
      ret[index] = val.(Converter_LuaConverterDownCast).ToConverter_LuaConverter().FP.(T)
   }
   return ret
}
type Converter_LuaConverterDownCast interface {
    ToConverter_LuaConverter() *Converter_LuaConverter
}
func Converter_LuaConverterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Converter_LuaConverterDownCast)
    if ok { return work.ToConverter_LuaConverter() }
    return nil
}
func (obj *Converter_LuaConverter) ToConverter_LuaConverter() *Converter_LuaConverter {
    return obj
}
func NewConverter_LuaConverter(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny, arg4 LnsAny, arg5 LnsInt, arg6 string, arg7 bool, arg8 bool, arg9 *Option_Option) *Converter_LuaConverter {
    obj := &Converter_LuaConverter{}
    obj.FP = obj
    obj.Runner_Runner.FP = obj
    obj.InitConverter_LuaConverter(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
// 376: DeclConstr
func (self *Converter_LuaConverter) InitConverter_LuaConverter(_env *LnsEnv, luaPath string,metaPath string,dependsPath LnsAny,astResult LnsAny,convMode LnsInt,path string,byteCompile bool,stripDebugInfo bool,option *Option_Option) {
    self.InitRunner_Runner(_env)
    self.luaPath = luaPath
    self.metaPath = metaPath
    self.dependsPath = dependsPath
    self.option = option
    self.stripDebugInfo = stripDebugInfo
    self.byteCompile = byteCompile
    self.byteStream = NewUtil_memStream(_env)
    self.byteMetaStream = NewUtil_memStream(_env)
    self.streamMem = NewUtil_memStream(_env)
    self.metaStreamMem = NewUtil_memStream(_env)
    self.dependsStreamMem = NewUtil_memStream(_env)
    self.astResult = astResult
    self.filterInfo = nil
    self.converterFunc = Converter_ConverterFunc(func(_env *LnsEnv) {
        var stream *Util_memStream
        stream = self.streamMem
        var metaStream *Util_memStream
        metaStream = self.metaStreamMem
        var outStream *Util_memStream
        var oMetaStream *Util_memStream
        outStream,oMetaStream = stream, metaStream
        var ast *AstInfo_ASTInfo
        ast = Converter_getAstFromResult_6_(_env, self.astResult)
        var needDepends bool
        needDepends = option.DependsPath != nil
        if needDepends{
            ast.FP.Get_node(_env).FP.ProcessFilter(_env, OutputDepend_createFilter(_env, self.dependsStreamMem.FP), 1)
        }
        if byteCompile{
            outStream = self.byteStream
            oMetaStream = self.byteMetaStream
        }
        var filterInfo *ConvLua_FilterInfo
        filterInfo = Converter_ast2LuaMain(_env, ast, path, outStream.FP, oMetaStream.FP, convMode, false, option)
        self.filterInfo = filterInfo
        filterInfo.FP.OutputLua(_env, Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode))
    })
    self.FP.Start(_env, 1, _env.GetVM().String_format("convlua -- %s", []LnsAny{path}))
}


// declaration Class -- GoConverter
type Converter_GoConverterMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    SaveGo(_env *LnsEnv)
    Start(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny) bool
}
type Converter_GoConverter struct {
    Runner_Runner
    memStream *Util_memStream
    path string
    converter Converter_ConverterFunc
    validFlag bool
    FP Converter_GoConverterMtd
}
func Converter_GoConverter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Converter_GoConverter).FP
}
      
func Converter_GoConverter_toSlice_Runner_Runner(slice []LnsAny) []*Runner_Runner {
   ret := make([]*Runner_Runner, len(slice))
   for index, val := range slice {
      ret[index] = &val.(Converter_GoConverterDownCast).ToConverter_GoConverter().Runner_Runner
   }
   return ret
}
      
func Converter_GoConverter_toSlice__IF[T any](slice []LnsAny) []T {
   ret := make([]T, len(slice))
   for index, val := range slice {
      ret[index] = val.(Converter_GoConverterDownCast).ToConverter_GoConverter().FP.(T)
   }
   return ret
}
type Converter_GoConverterDownCast interface {
    ToConverter_GoConverter() *Converter_GoConverter
}
func Converter_GoConverterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Converter_GoConverterDownCast)
    if ok { return work.ToConverter_GoConverter() }
    return nil
}
func (obj *Converter_GoConverter) ToConverter_GoConverter() *Converter_GoConverter {
    return obj
}
func NewConverter_GoConverter(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 string, arg4 *Option_Option, arg5 *ConvGo_Option) *Converter_GoConverter {
    obj := &Converter_GoConverter{}
    obj.FP = obj
    obj.Runner_Runner.FP = obj
    obj.InitConverter_GoConverter(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
// 483: DeclConstr
func (self *Converter_GoConverter) InitConverter_GoConverter(_env *LnsEnv, scriptPath string,astResult LnsAny,mod string,option *Option_Option,goOpt *ConvGo_Option) {
    self.InitRunner_Runner(_env)
    self.validFlag = true
    var path string
    path = Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string) + ".go"
    {
        _dir := option.OutputDir
        if !Lns_IsNil( _dir ) {
            dir := _dir.(string)
            path = _env.GetVM().String_format("%s/%s", []LnsAny{dir, path})
        }
    }
    self.path = path
    self.memStream = NewUtil_memStream(_env)
    self.converter = Converter_ConverterFunc(func(_env *LnsEnv) {
        var ast *AstInfo_ASTInfo
        ast = Converter_getAstFromResult_6_(_env, astResult)
        var rootNode *Nodes_RootNode
        
        {
            _rootNode := Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)
            if _rootNode == nil{
                return 
            } else {
                rootNode = _rootNode.(*Nodes_RootNode)
            }
        }
        for _pragma := range( rootNode.FP.Get_luneHelperInfo(_env).PragmaSet.Items ) {
            pragma := _pragma
            switch _matchExp0 := pragma.(type) {
            case *LuneControl_Pragma__limit_conv_code:
                codeSet := _matchExp0.Val1
                if Lns_op_not(codeSet.Has(LuneControl_Code__Go)){
                    self.validFlag = false
                    return 
                }
            }
        }
        var conv *Nodes_Filter
        conv = ConvGo_createFilter(_env, option.Testing, scriptPath, self.memStream.FP, ast, goOpt)
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast.FP.Get_node(_env))))
    })
    self.FP.Start(_env, 1, _env.GetVM().String_format("convgo -- %s", []LnsAny{scriptPath}))
}


// declaration Class -- PythonConverter
type Converter_PythonConverterMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    SavePython(_env *LnsEnv)
    Start(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny) bool
}
type Converter_PythonConverter struct {
    Runner_Runner
    memStream *Util_memStream
    path string
    converter Converter_ConverterFunc
    validFlag bool
    FP Converter_PythonConverterMtd
}
func Converter_PythonConverter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Converter_PythonConverter).FP
}
      
func Converter_PythonConverter_toSlice_Runner_Runner(slice []LnsAny) []*Runner_Runner {
   ret := make([]*Runner_Runner, len(slice))
   for index, val := range slice {
      ret[index] = &val.(Converter_PythonConverterDownCast).ToConverter_PythonConverter().Runner_Runner
   }
   return ret
}
      
func Converter_PythonConverter_toSlice__IF[T any](slice []LnsAny) []T {
   ret := make([]T, len(slice))
   for index, val := range slice {
      ret[index] = val.(Converter_PythonConverterDownCast).ToConverter_PythonConverter().FP.(T)
   }
   return ret
}
type Converter_PythonConverterDownCast interface {
    ToConverter_PythonConverter() *Converter_PythonConverter
}
func Converter_PythonConverterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Converter_PythonConverterDownCast)
    if ok { return work.ToConverter_PythonConverter() }
    return nil
}
func (obj *Converter_PythonConverter) ToConverter_PythonConverter() *Converter_PythonConverter {
    return obj
}
func NewConverter_PythonConverter(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 string, arg4 *Option_Option, arg5 *ConvPython_Option) *Converter_PythonConverter {
    obj := &Converter_PythonConverter{}
    obj.FP = obj
    obj.Runner_Runner.FP = obj
    obj.InitConverter_PythonConverter(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
// 552: DeclConstr
func (self *Converter_PythonConverter) InitConverter_PythonConverter(_env *LnsEnv, scriptPath string,astResult LnsAny,mod string,option *Option_Option,pythonOpt *ConvPython_Option) {
    self.InitRunner_Runner(_env)
    self.validFlag = true
    var path string
    path = Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string) + ".py"
    {
        _dir := option.OutputDir
        if !Lns_IsNil( _dir ) {
            dir := _dir.(string)
            path = _env.GetVM().String_format("%s/%s", []LnsAny{dir, path})
        }
    }
    self.path = path
    self.memStream = NewUtil_memStream(_env)
    self.converter = Converter_ConverterFunc(func(_env *LnsEnv) {
        var ast *AstInfo_ASTInfo
        ast = Converter_getAstFromResult_6_(_env, astResult)
        var rootNode *Nodes_RootNode
        
        {
            _rootNode := Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)
            if _rootNode == nil{
                return 
            } else {
                rootNode = _rootNode.(*Nodes_RootNode)
            }
        }
        for _pragma := range( rootNode.FP.Get_luneHelperInfo(_env).PragmaSet.Items ) {
            pragma := _pragma
            switch _matchExp0 := pragma.(type) {
            case *LuneControl_Pragma__limit_conv_code:
                codeSet := _matchExp0.Val1
                if Lns_op_not(codeSet.Has(LuneControl_Code__Python)){
                    self.validFlag = false
                    return 
                }
            }
        }
        var conv *Nodes_Filter
        conv = ConvPython_createFilter(_env, option.Testing, scriptPath, self.memStream.FP, ast, pythonOpt)
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvPython_Opt2Stem(NewConvPython_Opt(_env, ast.FP.Get_node(_env))))
    })
    self.FP.Start(_env, 1, _env.GetVM().String_format("convpython -- %s", []LnsAny{scriptPath}))
}


func Lns_Converter_init(_env *LnsEnv) {
    if init_Converter { return }
    init_Converter = true
    Converter__mod__ = "@lune.@base.@Converter"
    Lns_InitMod()
    Lns_Runner_init(_env)
    Lns_Util_init(_env)
    Lns_AstInfo_init(_env)
    Lns_Option_init(_env)
    Lns_Builtin_init(_env)
    Lns_frontInterface_init(_env)
    Lns_Types_init(_env)
    Lns_TransUnit_init(_env)
    Lns_Parser_init(_env)
    Lns_convLua_init(_env)
    Lns_convGo_init(_env)
    Lns_convPython_init(_env)
    Lns_Log_init(_env)
    Lns_Nodes_init(_env)
    Lns_LuneControl_init(_env)
    Lns_OutputDepend_init(_env)
}
func init() {
    init_Converter = false
}
