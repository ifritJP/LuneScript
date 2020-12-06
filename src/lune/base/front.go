// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_front bool
var front__mod__ string
var front_forceUpdateMeta bool
type front_ModuleUptodate__NeedTouch struct{
Val1 string
Val2 *front_MetaForBuildId
}
func (self *front_ModuleUptodate__NeedTouch) GetTxt() string {
return "ModuleUptodate.NeedTouch"
}
type front_ModuleUptodate__NeedUpdate struct{
}
var front_ModuleUptodate__NeedUpdate_Obj = &front_ModuleUptodate__NeedUpdate{}
func (self *front_ModuleUptodate__NeedUpdate) GetTxt() string {
return "ModuleUptodate.NeedUpdate"
}
type front_ModuleUptodate__Uptodate struct{
Val1 *front_MetaForBuildId
}
func (self *front_ModuleUptodate__Uptodate) GetTxt() string {
return "ModuleUptodate.Uptodate"
}
type front_OpenOStreamForConvert_1317_ func (arg1 LnsAny)(LnsAny, LnsAny, LnsAny)
type front_CloseOStreamForConvert_1320_ func (arg1 LnsAny,arg2 LnsAny,arg3 LnsAny)
// for 741: ExpCast
func conv2Form3993( src func ()) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        src()
        return []LnsAny{}
    }
}
// for 1426: ExpCast
func conv2Form7446( src func ()) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        src()
        return []LnsAny{}
    }
}
// for 1464: ExpCast
func conv2Form7507( src func ()) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        src()
        return []LnsAny{}
    }
}
// for 193
func front_convExp773(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 262
func front_convExp1165(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 497
func front_convExp2325(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 505
func front_convExp2354(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1253
func front_convExp6492(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1299
func front_convExp6687(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1313
func front_convExp6768(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1463
func front_convExp7524(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 280
func front_convExp1238(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 294
func front_convExp1329(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 295
func front_convExp1349(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 325
func front_convExp1439(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 490
func front_convExp2262(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 973
func front_convExp5063(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 266
func front_convExp1155(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 392
func front_convExp1708(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 593
func front_convExp2899(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 628
func front_convExp3200(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 637
func front_convExp3258(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 683
func front_convExp3498(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 692
func front_convExp3556(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 964
func front_convExp5006(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1020
func front_convExp5297(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1025
func front_convExp5335(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1052
func front_convExp5462(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1054
func front_convExp5477(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1058
func front_convExp5491(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1059
func front_convExp5507(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1102
func front_convExp5713(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1161
func front_convExp5964(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1162
func front_convExp5977(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1203
func front_convExp6214(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1204
func front_convExp6227(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1206
func front_convExp6244(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1238
func front_convExp6402(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1349
func front_convExp6885(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1350
func front_convExp6916(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1451
func front_convExp7471(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
func front__anonymous_1007_(ver LnsInt) {
    LuaVer_setCurVer(ver)
}
// 144: decl @lune.@base.@front.createPaser
func front_createPaser_1099_(path string,mod string) *Parser_Parser {
    var parser LnsAny
    parser = Parser_StreamParser_create(path, false, mod)
    {
        __exp := parser
        if __exp != nil {
            _exp := __exp.(*Parser_StreamParser)
            return &_exp.Parser_Parser
        }
    }
    panic("failed to open " + path)
// insert a dummy
    return nil
}

// 153: decl @lune.@base.@front.scriptPath2Module
func Front_scriptPath2Module(path string) string {
    return Util_scriptPath2Module(path)
}

// 188: decl @lune.@base.@front.loadFromChunk
func front_loadFromChunk_1114_(chunk LnsAny,err LnsAny) LnsAny {
    if err != nil{
        err_5755 := err.(string)
        Util_errorLog(err_5755)
    }
    if chunk != nil{
        chunk_5757 := chunk.(*Lns_luaValue)
        {
            _work := front_convExp773(Lns_2DDD(Lns_getVM().RunLoadedfunc(chunk_5757,Lns_2DDD([]LnsAny{}))[0]))
            if _work != nil {
                work := _work
                return work
            }
        }
        return nil
    }
    panic("failed to error")
// insert a dummy
    return nil
}

// 201: decl @lune.@base.@front.loadFromLuaTxt
func front_loadFromLuaTxt_1117_(txt string) LnsAny {
    return front_loadFromChunk_1114_(Lns_getVM().Load(txt, nil))
}

// 206: decl @lune.@base.@front.byteCompileFromLuaTxt
func front_byteCompileFromLuaTxt_1120_(txt string,stripDebugInfo bool) string {
    var chunk LnsAny
    var err LnsAny
    chunk,err = Lns_getVM().Load(txt, nil)
    if chunk != nil{
        chunk_5772 := chunk.(*Lns_luaValue)
        return Lns_getVM().String_dump(chunk_5772, stripDebugInfo)
    }
    panic(Lns_unwrapDefault( err, "load error").(string))
// insert a dummy
    return ""
}

// 275: decl @lune.@base.@front.getMetaInfo
func front_getMetaInfo_1186_(lnsPath string,mod string,outdir LnsAny)(LnsAny, string, string) {
    var moduleMetaPath string
    moduleMetaPath = lnsPath
    if outdir != nil{
        outdir_5833 := outdir.(string)
        moduleMetaPath = Lns_getVM().String_format("%s/%s", []LnsAny{outdir_5833, Lns_car(Lns_getVM().String_gsub(mod,"%.", "/")).(string)})
        
    }
    moduleMetaPath = front_convExp1238(Lns_2DDD(Lns_getVM().String_gsub(moduleMetaPath,"%.lns$", ".meta")))
    
    {
        _meta, _metaCode := front_MetaForBuildId_LoadFromMeta_1183_(moduleMetaPath)
        if _meta != nil && _metaCode != nil {
            meta := _meta.(*front_MetaForBuildId)
            metaCode := _metaCode.(string)
            return meta, moduleMetaPath, metaCode
        }
    }
    return nil, moduleMetaPath, ""
}

// 317: decl @lune.@base.@front.getModuleId
func front_getModuleId_1192_(lnsPath string,mod string,outdir LnsAny,metaInfo LnsAny) *FrontInterface_ModuleId {
    var buildCount LnsInt
    buildCount = 0
    var fileTime LnsReal
    
    {
        _fileTime := Depend_getFileLastModifiedTime(lnsPath)
        if _fileTime == nil{
            return FrontInterface_ModuleId__tempId
        } else {
            fileTime = _fileTime.(LnsReal)
        }
    }
    if Lns_op_not(metaInfo){
        metaInfo = front_convExp1439(Lns_2DDD(front_getMetaInfo_1186_(lnsPath, mod, outdir)))
        
    }
    if metaInfo != nil{
        metaInfo_5863 := metaInfo.(*front_MetaForBuildId)
        var buildId *FrontInterface_ModuleId
        buildId = metaInfo_5863.FP.CreateModuleId()
        buildCount = buildId.FP.Get_buildCount()
        
    }
    return FrontInterface_ModuleId_createId(fileTime, buildCount)
}

func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1206_() string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1209_() string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1212_() string {
    return "NeedUpdate"
}


func Front_getModuleIdAndCheckUptodate___anonymous_1218_() string {
    return "not found meta"
}















func Front_convertToLua___anonymous_1346_(stream LnsAny,metaStream LnsAny,dependStream LnsAny) {
    if dependStream != nil{
        dependStream_6281 := dependStream.(Lns_oStream)
        dependStream_6281.Close()
    }
}
// 1101: decl @lune.@base.@front.Front.saveToLua.txt2ModuleId
func Front_saveToLua__txt2ModuleId_1387_(txt string) *FrontInterface_ModuleId {
    var buildIdTxt string
    buildIdTxt = front_convExp5713(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(txt,"^_moduleObj.__buildId = ", "")).(string),"\"", "")))
    return FrontInterface_ModuleId_createIdFromTxt(buildIdTxt)
}







// 1363: decl @lune.@base.@front.convertLnsCode2LuaCode
func Front_convertLnsCode2LuaCode(lnsCode string,path string) string {
    var option *Option_Option
    option = NewOption_Option()
    option.ScriptPath = path
    
    option.UseLuneModule = Option_getRuntimeModule()
    
    option.UseIpairs = true
    
    var front *front_Front
    front = Newfront_Front(option)
    return front.FP.ConvertLns2LuaCode(NewFrontInterface_ImportModuleInfo(), NewParser_TxtStream(lnsCode).FP, path)
}




// 1483: decl @lune.@base.@front.exec
func Front_exec(args *LnsList) {
    var version LnsReal
    version = Lns_unwrapDefault( Lns_tonumber(Lns_car(Lns_getVM().String_gsub(Depend_getLuaVersion(),"^[^%d]+", "")).(string), nil), 0.0).(LnsReal)
    if version < 5.1{
        Lns_io_stderr.Write(Lns_getVM().String_format("LuneScript doesn't support this lua version(%s). %s\n", []LnsAny{version, "please use the version >= 5.1."}))
        Lns_getVM().OS_exit(1)
    }
    var option *Option_Option
    option = Option_analyze(args)
    var front *front_Front
    front = Newfront_Front(option)
    front.FP.Exec()
}

// 1500: decl @lune.@base.@front.setFront
func Front_setFront() {
    var option *Option_Option
    option = Option_createDefaultOption("dummy.lns")
    Newfront_Front(option)
}

// declaration Class -- LoadInfo
type front_LoadInfoMtd interface {
}
type front_LoadInfo struct {
    Mod LnsAny
    Meta LnsAny
    FP front_LoadInfoMtd
}
func front_LoadInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_LoadInfo).FP
}
type front_LoadInfoDownCast interface {
    Tofront_LoadInfo() *front_LoadInfo
}
func front_LoadInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(front_LoadInfoDownCast)
    if ok { return work.Tofront_LoadInfo() }
    return nil
}
func (obj *front_LoadInfo) Tofront_LoadInfo() *front_LoadInfo {
    return obj
}
func Newfront_LoadInfo(arg1 LnsAny, arg2 LnsAny) *front_LoadInfo {
    obj := &front_LoadInfo{}
    obj.FP = obj
    obj.Initfront_LoadInfo(arg1, arg2)
    return obj
}
func (self *front_LoadInfo) Initfront_LoadInfo(arg1 LnsAny, arg2 LnsAny) {
    self.Mod = arg1
    self.Meta = arg2
}

// declaration Class -- Front
type front_FrontMtd interface {
    CheckDiag(arg1 string)
    checkUptodateMeta(arg1 string, arg2 LnsAny) LnsAny
    Complete(arg1 string)
    convert(arg1 *TransUnit_ASTInfo, arg2 string, arg3 Lns_oStream, arg4 Lns_oStream, arg5 LnsInt, arg6 bool)
    convertFromAst(arg1 *TransUnit_ASTInfo, arg2 string, arg3 LnsInt)(string, string)
    ConvertLns2LuaCode(arg1 *FrontInterface_ImportModuleInfo, arg2 Lns_iStream, arg3 string) string
    ConvertLuaToStreamFromScript(arg1 LnsAny, arg2 *FrontInterface_ModuleId, arg3 LnsAny, arg4 LnsInt, arg5 string, arg6 string, arg7 bool, arg8 bool, arg9 front_OpenOStreamForConvert_1317_, arg10 LnsAny) LnsAny
    convertToLua(arg1 string)
    createAst(arg1 *FrontInterface_ImportModuleInfo, arg2 *Parser_Parser, arg3 string, arg4 *FrontInterface_ModuleId, arg5 LnsAny, arg6 LnsInt, arg7 LnsAny) *TransUnit_ASTInfo
    CreateGlue(arg1 string)
    createGoOption(arg1 string) *ConvGo_Option
    createPaser(arg1 string) *Parser_Parser
    DumpAst(arg1 string)
    DumpTokenize(arg1 string)
    Error(arg1 string)
    Exec()
    Format(arg1 string)
    getLoadInfo(arg1 string) LnsAny
    getModuleIdAndCheckUptodate(arg1 string, arg2 string)(*FrontInterface_ModuleId, LnsAny)
    Inquire(arg1 string)
    loadFile(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string)(LnsAny, LnsAny)
    loadFileToLuaCode(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string)(LnsAny, string)
    LoadFromLnsTxt(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string) LnsAny
    loadLua(arg1 string) LnsAny
    LoadMeta(arg1 *FrontInterface_ImportModuleInfo, arg2 string) LnsAny
    LoadModule(arg1 string)(LnsAny, LnsAny)
    outputBootC(arg1 string)
    OutputBuiltin(arg1 string)
    SaveToC(arg1 string, arg2 *TransUnit_ASTInfo)
    SaveToGo(arg1 string, arg2 *TransUnit_ASTInfo)
    SaveToLua(arg1 *front_UpdateInfo) bool
    searchLuaFile(arg1 string, arg2 LnsAny) LnsAny
    SearchModule(arg1 string) LnsAny
    searchModuleFile(arg1 string, arg2 string, arg3 LnsAny) LnsAny
    setLoadInfo(arg1 string, arg2 *front_LoadInfo)
}
type front_Front struct {
    option *Option_Option
    loadedMap *LnsMap
    loadedMapTest *LnsMap
    loadedMetaMap *LnsMap
    convertedMap *LnsMap
    preloadedModMap *LnsMap
    FP front_FrontMtd
}
func front_Front2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_Front).FP
}
type front_FrontDownCast interface {
    Tofront_Front() *front_Front
}
func front_FrontDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(front_FrontDownCast)
    if ok { return work.Tofront_Front() }
    return nil
}
func (obj *front_Front) Tofront_Front() *front_Front {
    return obj
}
func Newfront_Front(arg1 *Option_Option) *front_Front {
    obj := &front_Front{}
    obj.FP = obj
    obj.Initfront_Front(arg1)
    return obj
}
// 80: DeclConstr
func (self *front_Front) Initfront_Front(option *Option_Option) {
    self.option = option
    
    self.loadedMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.loadedMapTest = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.loadedMetaMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.convertedMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    FrontInterface_setFront(self.FP)
    var loadedMap *LnsMap
    loadedMap = NewLnsMap( map[LnsAny]LnsAny{})
    {
        _exp242 := Depend_getLoadedMod()
        _key242, _val242 := _exp242.Get1stFromMap()
        for _key242 != nil {
            mod := _key242.(string)
            modval := _val242
            if mod == "lune.base.Testing"{
                loadedMap.Set(mod,modval)
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( option.Testing) &&
                Lns_GetEnv().SetStackVal( modval.(*Lns_luaValue).GetAt("__enableTest")) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(option.Testing)) &&
                Lns_GetEnv().SetStackVal( modval.(*Lns_luaValue).GetAt("__enableTest")) )){
                loadedMap.Set(mod,modval)
            }
            _key242, _val242 = _exp242.NextFromMap( _key242 )
        }
    }
    self.preloadedModMap = loadedMap
    
}

// 105: decl @lune.@base.@front.Front.getLoadInfo
func (self *front_Front) getLoadInfo(mod string) LnsAny {
    if self.option.Testing{
        return self.loadedMapTest.Items[mod]
    }
    return self.loadedMap.Items[mod]
}

// 112: decl @lune.@base.@front.Front.setLoadInfo
func (self *front_Front) setLoadInfo(mod string,info *front_LoadInfo) {
    if self.option.Testing{
        self.loadedMapTest.Set(mod,info)
    }
    self.loadedMap.Set(mod,info)
}






// 129: decl @lune.@base.@front.Front.error
func (self *front_Front) Error(message string) {
    Util_errorLog(message)
    Util_printStackTrace()
    Lns_getVM().OS_exit(1)
}

// 135: decl @lune.@base.@front.Front.loadLua
func (self *front_Front) loadLua(path string) LnsAny {
    var chunk LnsAny
    var err LnsAny
    chunk,err = Lns_getVM().Loadfile(path)
    if chunk != nil{
        chunk_5707 := chunk.(*Lns_luaValue)
        return Lns_unwrap( Lns_car(Lns_getVM().RunLoadedfunc(chunk_5707,Lns_2DDD([]LnsAny{}))[0]))
    }
    Util_errorLog(Lns_unwrapDefault( err, Lns_getVM().String_format("load error -- %s.", []LnsAny{path})).(string))
    return nil
}

// 157: decl @lune.@base.@front.Front.createPaser
func (self *front_Front) createPaser(scriptPath string) *Parser_Parser {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    return front_createPaser_1099_(scriptPath, mod)
}

// 164: decl @lune.@base.@front.Front.createAst
func (self *front_Front) createAst(importModuleInfo *FrontInterface_ImportModuleInfo,parser *Parser_Parser,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *TransUnit_ASTInfo {
    var transUnit *TransUnit_TransUnit
    transUnit = NewTransUnit_TransUnit(moduleId, importModuleInfo, &NewConvLua_MacroEvalImp().Nodes_MacroEval, analyzeModule, analyzeMode, pos, self.option.TargetLuaVer, self.option.TransCtrlInfo)
    return transUnit.FP.CreateAST(parser, false, mod)
}

// 176: decl @lune.@base.@front.Front.convert
func (self *front_Front) convert(ast *TransUnit_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool) {
    var conv *Nodes_Filter
    conv = ConvLua_createFilter(streamName, stream, metaStream, convMode, inMacro, ast.FP.Get_moduleTypeInfo(), ast.FP.Get_processInfo(), ast.FP.Get_moduleSymbolKind(), self.option.UseLuneModule, self.option.TargetLuaVer, self.option.Testing, self.option.UseIpairs)
    ast.FP.Get_node().FP.ProcessFilter(conv, ConvLua_Opt2Stem(NewConvLua_Opt(ast.FP.Get_node())))
}

// 215: decl @lune.@base.@front.Front.convertFromAst
func (self *front_Front) convertFromAst(ast *TransUnit_ASTInfo,streamName string,mode LnsInt)(string, string) {
    var stream *Util_memStream
    stream = NewUtil_memStream()
    var metaStream *Util_memStream
    metaStream = NewUtil_memStream()
    self.FP.convert(ast, streamName, stream.FP, metaStream.FP, mode, false)
    return metaStream.FP.Get_txt(), stream.FP.Get_txt()
}

// 229: decl @lune.@base.@front.Front.loadFromLnsTxt
func (self *front_Front) LoadFromLnsTxt(importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    var transUnit *TransUnit_TransUnit
    transUnit = NewTransUnit_TransUnit(FrontInterface_ModuleId__tempId, importModuleInfo, &NewConvLua_MacroEvalImp().Nodes_MacroEval, nil, nil, nil, self.option.TargetLuaVer, self.option.TransCtrlInfo)
    var stream *Parser_TxtStream
    stream = NewParser_TxtStream(txt)
    var parser *Parser_StreamParser
    parser = NewParser_StreamParser(stream.FP, name, false)
    var ast *TransUnit_ASTInfo
    ast = transUnit.FP.CreateAST(&parser.Parser_Parser, false, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(ast, name, ConvLua_ConvMode__Exec)
    return Lns_unwrap( front_loadFromLuaTxt_1117_(luaTxt))
}

// 287: decl @lune.@base.@front.Front.searchModuleFile
func (self *front_Front) searchModuleFile(mod string,suffix string,addPath LnsAny) LnsAny {
    var lnsSearchPath string
    lnsSearchPath = Lns_package_path
    if addPath != nil{
        addPath_5846 := addPath.(string)
        lnsSearchPath = Lns_getVM().String_format("%s/?%s;%s", []LnsAny{addPath_5846, suffix, Lns_package_path})
        
    }
    lnsSearchPath = front_convExp1329(Lns_2DDD(Lns_getVM().String_gsub(lnsSearchPath,"%.lua$", suffix)))
    
    lnsSearchPath = front_convExp1349(Lns_2DDD(Lns_getVM().String_gsub(lnsSearchPath,"%.lua;", suffix + ";")))
    
    var foundPath string
    
    {
        _foundPath := Depend_searchpath(mod, lnsSearchPath)
        if _foundPath == nil{
            return nil
        } else {
            foundPath = _foundPath.(string)
        }
    }
    return Lns_car(Lns_getVM().String_gsub(foundPath,"^./", "")).(string)
}

// 357: decl @lune.@base.@front.Front.getModuleIdAndCheckUptodate
func (self *front_Front) getModuleIdAndCheckUptodate(lnsPath string,mod string)(*FrontInterface_ModuleId, LnsAny) {
    __func__ := "@lune.@base.@front.Front.getModuleIdAndCheckUptodate"
    var uptodate LnsAny
    uptodate = front_ModuleUptodate__NeedUpdate_Obj
    if self.option.TransCtrlInfo.UptodateMode == Types_CheckingUptodateMode__Force{
        return FrontInterface_ModuleId__tempId, uptodate
    }
    var checkDependUptodate func(metaTime LnsReal,metaInfo *front_MetaForBuildId,metaCode string) LnsAny
    checkDependUptodate = func(metaTime LnsReal,metaInfo *front_MetaForBuildId,metaCode string) LnsAny {
        __func__ := "@lune.@base.@front.Front.getModuleIdAndCheckUptodate.checkDependUptodate"
        for _depMod, _dependItem := range( metaInfo.__dependModuleMap.Items ) {
            depMod := _depMod.(string)
            dependItem := _dependItem.(front_DependMetaInfoDownCast).Tofront_DependMetaInfo()
            var modMetaPath string
            
            {
                _modMetaPath := self.FP.searchModuleFile(depMod, ".meta", self.option.OutputDir)
                if _modMetaPath == nil{
                    Log_log(Log_Level__Debug, __func__, 380, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1206_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    modMetaPath = _modMetaPath.(string)
                }
            }
            var time LnsReal
            
            {
                _time := Depend_getFileLastModifiedTime(modMetaPath)
                if _time == nil{
                    Log_log(Log_Level__Debug, __func__, 385, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1209_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    time = _time.(LnsReal)
                }
            }
            if time > metaTime{
                var dependMeta *front_MetaForBuildId
                
                {
                    _dependMeta := front_convExp1708(Lns_2DDD(front_MetaForBuildId_LoadFromMeta_1183_(modMetaPath)))
                    if _dependMeta == nil{
                        Log_log(Log_Level__Debug, __func__, 393, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1212_))
                        
                        return front_ModuleUptodate__NeedUpdate_Obj
                    } else {
                        dependMeta = _dependMeta.(*front_MetaForBuildId)
                    }
                }
                var orgMetaModuleId *FrontInterface_ModuleId
                orgMetaModuleId = FrontInterface_ModuleId_createIdFromTxt(dependItem.BuildId)
                var metaModuleId *FrontInterface_ModuleId
                metaModuleId = dependMeta.FP.CreateModuleId()
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( metaModuleId.FP.Get_buildCount() != 0) &&
                    Lns_GetEnv().SetStackVal( metaModuleId.FP.Get_buildCount() != orgMetaModuleId.FP.Get_buildCount()) ).(bool)){
                    Log_log(Log_Level__Debug, __func__, 403, Log_CreateMessage(func() string {
                        return Lns_getVM().String_format("NeedUpdate: %s, %d, %d", []LnsAny{modMetaPath, metaModuleId.FP.Get_buildCount(), orgMetaModuleId.FP.Get_buildCount()})
                    }))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                }
            }
        }
        if self.option.TransCtrlInfo.UptodateMode == Types_CheckingUptodateMode__Touch{
            return &front_ModuleUptodate__NeedTouch{metaCode, metaInfo}
        }
        return &front_ModuleUptodate__Uptodate{metaInfo}
    }
    var metaInfo LnsAny
    var metaPath string
    var metaCode string
    metaInfo,metaPath,metaCode = front_getMetaInfo_1186_(lnsPath, mod, self.option.OutputDir)
    if metaInfo != nil{
        metaInfo_5913 := metaInfo.(*front_MetaForBuildId)
        if metaInfo_5913.__enableTest == self.option.Testing{
            var buildId *FrontInterface_ModuleId
            buildId = FrontInterface_ModuleId_createIdFromTxt(metaInfo_5913.__buildId)
            if buildId != FrontInterface_ModuleId__tempId{
                var lnsTime LnsAny
                lnsTime = Depend_getFileLastModifiedTime(lnsPath)
                var metaTime LnsAny
                metaTime = Depend_getFileLastModifiedTime(metaPath)
                if lnsTime != nil && metaTime != nil{
                    lnsTime_5920 := lnsTime.(LnsReal)
                    metaTime_5921 := metaTime.(LnsReal)
                    if lnsTime_5920 == buildId.FP.Get_modTime(){
                        uptodate = checkDependUptodate(metaTime_5921, metaInfo_5913, metaCode)
                        
                    }
                }
            }
        } else { 
        }
    } else {
        Log_log(Log_Level__Debug, __func__, 441, Log_CreateMessage(Front_getModuleIdAndCheckUptodate___anonymous_1218_))
        
    }
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_1192_(lnsPath, mod, self.option.OutputDir, metaInfo)
    if moduleId == FrontInterface_ModuleId__tempId{
        Util_err(Lns_getVM().String_format("not found -- %s", []LnsAny{lnsPath}))
    }
    return moduleId, uptodate
}

// 454: decl @lune.@base.@front.Front.convertLns2LuaCode
func (self *front_Front) ConvertLns2LuaCode(importModuleInfo *FrontInterface_ImportModuleInfo,stream Lns_iStream,streamName string) string {
    var mod string
    mod = Front_scriptPath2Module(streamName)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(importModuleInfo, &NewParser_StreamParser(stream, streamName, false).Parser_Parser, mod, FrontInterface_ModuleId_createId(0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(ast, streamName, ConvLua_ConvMode__Exec)
    return luaTxt
}

// 473: decl @lune.@base.@front.Front.loadFileToLuaCode
func (self *front_Front) loadFileToLuaCode(importModuleInfo *FrontInterface_ImportModuleInfo,path string,mod string)(LnsAny, string) {
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(importModuleInfo, front_createPaser_1099_(path, mod), mod, front_getModuleId_1192_(path, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(ast, path, ConvLua_ConvMode__Exec)
    if self.option.UpdateOnLoad{
        var saveFile func(suffix string,txt string,byteCompile bool,stripDebugInfo bool,checkUpdate bool)
        saveFile = func(suffix string,txt string,byteCompile bool,stripDebugInfo bool,checkUpdate bool) {
            var newpath string
            newpath = ""
            {
                _dir := self.option.OutputDir
                if _dir != nil {
                    dir := _dir.(string)
                    newpath = Lns_getVM().String_format("%s/%s%s", []LnsAny{dir, Lns_car(Lns_getVM().String_gsub(mod,"%.", "/")).(string), suffix})
                    
                } else {
                    newpath = front_convExp2262(Lns_2DDD(Lns_getVM().String_gsub(path,".lns$", suffix)))
                    
                }
            }
            var saveTxt string
            saveTxt = txt
            if byteCompile{
                saveTxt = front_byteCompileFromLuaTxt_1120_(saveTxt, stripDebugInfo)
                
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(front_forceUpdateMeta)) &&
                Lns_GetEnv().SetStackVal( checkUpdate) ).(bool)){
                {
                    _fileObj := front_convExp2325(Lns_2DDD(Lns_io_open(newpath, nil)))
                    if _fileObj != nil {
                        fileObj := _fileObj.(Lns_luaStream)
                        var oldTxt LnsAny
                        oldTxt = fileObj.Read("*a")
                        if saveTxt == oldTxt{
                            return 
                        }
                    }
                }
            }
            {
                _fileObj := front_convExp2354(Lns_2DDD(Lns_io_open(newpath, "w")))
                if _fileObj != nil {
                    fileObj := _fileObj.(Lns_luaStream)
                    fileObj.Write(saveTxt)
                    fileObj.Close()
                }
            }
        }
        saveFile(".lua", luaTxt, self.option.ByteCompile, self.option.StripDebugInfo, false)
        saveFile(".meta", metaTxt, self.option.ByteCompile, true, true)
    }
    var meta LnsAny
    meta = Lns_unwrap( front_loadFromLuaTxt_1117_(metaTxt))
    return meta, luaTxt
}

// 525: decl @lune.@base.@front.Front.loadFile
func (self *front_Front) loadFile(importModuleInfo *FrontInterface_ImportModuleInfo,path string,mod string)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@front.Front.loadFile"
    Log_log(Log_Level__Info, __func__, 529, Log_CreateMessage(func() string {
        __func__ := "@lune.@base.@front.Front.loadFile.<anonymous>"
        return Lns_getVM().String_format("start %s:%s", []LnsAny{__func__, mod})
    }))
    
    var meta LnsAny
    var luaTxt string
    meta,luaTxt = self.FP.loadFileToLuaCode(importModuleInfo, path, mod)
    {
        _preLoadInfo := self.preloadedModMap.Items[mod]
        if _preLoadInfo != nil {
            preLoadInfo := _preLoadInfo
            return meta, preLoadInfo
        }
    }
    return meta, Lns_unwrap( front_loadFromLuaTxt_1117_(luaTxt))
}

// 541: decl @lune.@base.@front.Front.searchModule
func (self *front_Front) SearchModule(mod string) LnsAny {
    return self.FP.searchModuleFile(mod, ".lns", nil)
}

// 545: decl @lune.@base.@front.Front.searchLuaFile
func (self *front_Front) searchLuaFile(moduleFullName string,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(moduleFullName, ".lua", addSearchPath)
}

// 565: decl @lune.@base.@front.Front.checkUptodateMeta
func (self *front_Front) checkUptodateMeta(metaPath string,addSearchPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.checkUptodateMeta"
    var metaObj LnsAny
    
    {
        _metaObj := self.FP.loadLua(metaPath)
        if _metaObj == nil{
            Log_log(Log_Level__Warn, __func__, 567, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("load error -- %s", []LnsAny{metaPath})
            }))
            
            return nil
        } else {
            metaObj = _metaObj
        }
    }
    var meta *Lns_luaValue
    meta = metaObj.(*Lns_luaValue)
    if meta.GetAt( "__formatVersion" ).(string) != Ver_metaVersion{
        Log_log(Log_Level__Warn, __func__, 572, Log_CreateMessage(func() string {
            return Lns_getVM().String_format("unmatch meta version -- %s", []LnsAny{metaPath})
        }))
        
        return nil
    }
    if meta.GetAt( "__hasTest" ).(bool){
        if meta.GetAt( "__enableTest" ).(bool) != self.option.Testing{
            Log_log(Log_Level__Warn, __func__, 578, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("unmatch test setting -- %s", []LnsAny{metaPath})
            }))
            
            return nil
        }
    }
    {
        _exp3088 := meta.GetAt( "__dependModuleMap" ).(*Lns_luaValue)
        _key3088, _ := _exp3088.Get1stFromMap()
        for _key3088 != nil {
            moduleFullName := _key3088.(string)
            {
                _lnsPath := self.FP.SearchModule(moduleFullName)
                if _lnsPath != nil {
                    lnsPath := _lnsPath.(string)
                    {
                        _moduleLuaPath := self.FP.searchLuaFile(moduleFullName, addSearchPath)
                        if _moduleLuaPath != nil {
                            moduleLuaPath := _moduleLuaPath.(string)
                            if Lns_op_not(Util_getReadyCode(lnsPath, metaPath)){
                                Log_log(Log_Level__Warn, __func__, 589, Log_CreateMessage(func() string {
                                    return Lns_getVM().String_format("not ready -- %s, %s", []LnsAny{lnsPath, metaPath})
                                }))
                                
                                return nil
                            }
                            var moduleMetaPath string
                            moduleMetaPath = front_convExp2899(Lns_2DDD(Lns_getVM().String_gsub(moduleLuaPath,"%.lua$", ".meta")))
                            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                Lns_GetEnv().SetStackVal( Depend_existFile(moduleMetaPath)) &&
                                Lns_GetEnv().SetStackVal( Lns_op_not(Util_getReadyCode(moduleMetaPath, metaPath))) ).(bool)){
                                Log_log(Log_Level__Warn, __func__, 597, Log_CreateMessage(func() string {
                                    return Lns_getVM().String_format("not ready -- %s, %s", []LnsAny{moduleMetaPath, metaPath})
                                }))
                                
                                return nil
                            }
                        } else {
                            Log_log(Log_Level__Warn, __func__, 602, Log_CreateMessage(func() string {
                                return Lns_getVM().String_format("not found .lua file for -- %s", []LnsAny{moduleFullName})
                            }))
                            
                            return nil
                        }
                    }
                } else {
                    Log_log(Log_Level__Warn, __func__, 607, Log_CreateMessage(func() string {
                        return Lns_getVM().String_format("not found .lns file -- %s", []LnsAny{moduleFullName})
                    }))
                    
                    return nil
                }
            }
            _key3088, _ = _exp3088.NextFromMap( _key3088 )
        }
    }
    return meta
}

// 615: decl @lune.@base.@front.Front.loadModule
func (self *front_Front) LoadModule(mod string)(LnsAny, LnsAny) {
    if Lns_op_not(self.FP.getLoadInfo(mod)){
        {
            _luaTxt := self.convertedMap.Items[mod]
            if _luaTxt != nil {
                luaTxt := _luaTxt.(string)
                {
                    _meta := self.loadedMetaMap.Items[mod]
                    if _meta != nil {
                        meta := _meta
                        self.FP.setLoadInfo(mod, Newfront_LoadInfo(Lns_unwrap( front_loadFromLuaTxt_1117_(luaTxt)), meta))
                    } else {
                        panic(Lns_getVM().String_format("nothing meta -- %s", []LnsAny{mod}))
                    }
                }
            } else {
                {
                    _lnsPath := self.FP.SearchModule(mod)
                    if _lnsPath != nil {
                        lnsPath := _lnsPath.(string)
                        var luaPath LnsAny
                        luaPath = front_convExp3200(Lns_2DDD(Lns_getVM().String_gsub(lnsPath, "%.lns$", ".lua")))
                        {
                            _dir := self.option.OutputDir
                            if _dir != nil {
                                dir := _dir.(string)
                                luaPath = self.FP.searchLuaFile(mod, dir)
                                
                            }
                        }
                        var loadVal LnsAny
                        loadVal = nil
                        if luaPath != nil{
                            luaPath_6055 := luaPath.(string)
                            if Util_getReadyCode(lnsPath, luaPath_6055){
                                var metaPath string
                                metaPath = front_convExp3258(Lns_2DDD(Lns_getVM().String_gsub(luaPath_6055, "%.lua$", ".meta")))
                                if Util_getReadyCode(lnsPath, metaPath){
                                    {
                                        _preLoadInfo := self.preloadedModMap.Items[mod]
                                        if _preLoadInfo != nil {
                                            preLoadInfo := _preLoadInfo
                                            loadVal = preLoadInfo
                                            
                                        } else {
                                            loadVal = self.FP.loadLua(luaPath_6055)
                                            
                                        }
                                    }
                                    {
                                        __exp := loadVal
                                        if __exp != nil {
                                            _exp := __exp
                                            {
                                                _meta := self.FP.checkUptodateMeta(metaPath, self.option.OutputDir)
                                                if _meta != nil {
                                                    meta := _meta
                                                    self.FP.setLoadInfo(mod, Newfront_LoadInfo(_exp, meta))
                                                } else {
                                                    loadVal = nil
                                                    
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        if loadVal == nil{
                            var meta LnsAny
                            var workVal LnsAny
                            meta,workVal = self.FP.loadFile(NewFrontInterface_ImportModuleInfo(), lnsPath, mod)
                            self.FP.setLoadInfo(mod, Newfront_LoadInfo(workVal, meta))
                        }
                    }
                }
            }
        }
    }
    {
        __exp := self.FP.getLoadInfo(mod)
        if __exp != nil {
            _exp := __exp.(*front_LoadInfo)
            return _exp.Mod, _exp.Meta
        }
    }
    panic(Lns_getVM().String_format("load error, %s", []LnsAny{mod}))
// insert a dummy
    return nil,nil
}

// 674: decl @lune.@base.@front.Front.loadMeta
func (self *front_Front) LoadMeta(importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    __func__ := "@lune.@base.@front.Front.loadMeta"
    if self.loadedMetaMap.Items[mod] == nil{
        {
            __exp := self.FP.getLoadInfo(mod)
            if __exp != nil {
                _exp := __exp.(*front_LoadInfo)
                self.loadedMetaMap.Set(mod,_exp.Meta)
            } else {
                {
                    _lnsPath := self.FP.SearchModule(mod)
                    if _lnsPath != nil {
                        lnsPath := _lnsPath.(string)
                        var luaPath LnsAny
                        luaPath = front_convExp3498(Lns_2DDD(Lns_getVM().String_gsub(lnsPath, "%.lns$", ".lua")))
                        {
                            _dir := self.option.OutputDir
                            if _dir != nil {
                                dir := _dir.(string)
                                luaPath = self.FP.searchLuaFile(mod, dir)
                                
                            }
                        }
                        var meta LnsAny
                        meta = nil
                        if luaPath != nil{
                            luaPath_6089 := luaPath.(string)
                            if Util_getReadyCode(lnsPath, luaPath_6089){
                                var metaPath string
                                metaPath = front_convExp3556(Lns_2DDD(Lns_getVM().String_gsub(luaPath_6089, "%.lua$", ".meta")))
                                if Util_getReadyCode(lnsPath, metaPath){
                                    meta = self.FP.checkUptodateMeta(metaPath, self.option.OutputDir)
                                    
                                    if meta != nil{
                                        meta_6094 := meta
                                        self.loadedMetaMap.Set(mod,meta_6094)
                                    }
                                } else { 
                                    Log_log(Log_Level__Warn, __func__, 700, Log_CreateMessage(func() string {
                                        return Lns_getVM().String_format("%s not ready meta %s, %s", []LnsAny{mod, lnsPath, metaPath})
                                    }))
                                    
                                }
                            } else { 
                                Log_log(Log_Level__Warn, __func__, 704, Log_CreateMessage(func() string {
                                    return Lns_getVM().String_format("%s not ready lua %s, %s", []LnsAny{mod, lnsPath, luaPath_6089})
                                }))
                                
                            }
                        } else {
                            Log_log(Log_Level__Warn, __func__, 708, Log_CreateMessage(func() string {
                                return Lns_getVM().String_format("%s not found lua in %s", []LnsAny{mod, self.option.OutputDir})
                            }))
                            
                        }
                        if meta == nil{
                            var metawork LnsAny
                            var luaTxt string
                            metawork,luaTxt = self.FP.loadFileToLuaCode(importModuleInfo, lnsPath, mod)
                            self.loadedMetaMap.Set(mod,metawork)
                            self.convertedMap.Set(mod,luaTxt)
                        }
                    }
                }
            }
        }
    }
    return self.loadedMetaMap.Items[mod]
}

// 723: decl @lune.@base.@front.Front.dumpTokenize
func (self *front_Front) DumpTokenize(scriptPath string) {
    var parser *Parser_Parser
    parser = self.FP.createPaser(scriptPath)
    for  {
        var token *Types_Token
        
        {
            _token := parser.FP.GetToken()
            if _token == nil{
                break
            } else {
                token = _token.(*Types_Token)
            }
        }
        Lns_print([]LnsAny{Types_TokenKind_getTxt( token.Kind), token.Pos.LineNo, token.Pos.Column, token.Txt})
    }
}

// 735: decl @lune.@base.@front.Front.dumpAst
func (self *front_Front) DumpAst(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    Depend_profile(self.option.ValidProf, conv2Form3993(func() {
        var ast *TransUnit_ASTInfo
        ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1192_(scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
        ast.FP.Get_node().FP.ProcessFilter(DumpNode_createFilter(ast.FP.Get_moduleTypeInfo(), ast.FP.Get_processInfo(), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt("", 0)))
    }), scriptPath + ".profi")
}

// 752: decl @lune.@base.@front.Front.format
func (self *front_Front) Format(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1192_(scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    ast.FP.Get_node().FP.ProcessFilter(Formatter_createFilter(ast.FP.Get_moduleTypeInfo(), Lns_io_stdout), Formatter_Opt2Stem(NewFormatter_Opt(ast.FP.Get_node())))
}

// 765: decl @lune.@base.@front.Front.checkDiag
func (self *front_Front) CheckDiag(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    Util_setErrorCode(0)
    self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1192_(scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Diag, nil)
}

// 775: decl @lune.@base.@front.Front.complete
func (self *front_Front) Complete(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1192_(scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Complete, self.option.AnalyzePos)
}

// 783: decl @lune.@base.@front.Front.inquire
func (self *front_Front) Inquire(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1192_(scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Inquire, self.option.AnalyzePos)
}

// 792: decl @lune.@base.@front.Front.createGlue
func (self *front_Front) CreateGlue(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1192_(scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    var filter *Nodes_Filter
    filter = GlueFilter_createFilter(self.option.OutputDir)
    ast.FP.Get_node().FP.ProcessFilter(filter, 0)
}

// 825: decl @lune.@base.@front.Front.convertLuaToStreamFromScript
func (self *front_Front) ConvertLuaToStreamFromScript(parser LnsAny,moduleId *FrontInterface_ModuleId,uptodate LnsAny,convMode LnsInt,path string,mod string,byteCompile bool,stripDebugInfo bool,openOStream front_OpenOStreamForConvert_1317_,closeOStream LnsAny) LnsAny {
    var outputDependInfo func(stream LnsAny,metaInfo LnsAny)
    outputDependInfo = func(stream LnsAny,metaInfo LnsAny) {
        if stream != nil{
            stream_6185 := stream.(Lns_oStream)
            if metaInfo != nil{
                metaInfo_6187 := metaInfo.(*front_MetaForBuildId)
                var dependInfo *OutputDepend_DependInfo
                dependInfo = NewOutputDepend_DependInfo(mod)
                for _dependMod, _ := range( metaInfo_6187.__dependModuleMap.Items ) {
                    dependMod := _dependMod.(string)
                    dependInfo.FP.AddImpotModule(dependMod)
                }
                for _, _subMod := range( metaInfo_6187.__subModuleMap.Items ) {
                    subMod := _subMod.(string)
                    dependInfo.FP.AddSubMod(subMod)
                }
                dependInfo.FP.Output(stream_6185)
            } else {
                Util_err("metaInfo is nil")
            }
        }
    }
    var streamDst LnsAny
    var metaStreamDst LnsAny
    var dependsStreamDst LnsAny
    streamDst,metaStreamDst,dependsStreamDst = openOStream(uptodate)
    var streamMem *Util_memStream
    streamMem = NewUtil_memStream()
    var metaStreamMem *Util_memStream
    metaStreamMem = NewUtil_memStream()
    var dependsStreamMem *Util_memStream
    dependsStreamMem = NewUtil_memStream()
    var stream LnsAny
    var metaStream LnsAny
    var dependsStream LnsAny
    if Str_isValidStrBuilder(){
        stream = streamMem.FP
        
        metaStream = metaStreamMem.FP
        
        dependsStream = dependsStreamMem.FP
        
    } else { 
        stream = streamDst
        
        metaStream = metaStreamDst
        
        dependsStream = dependsStreamDst
        
    }
    var retAst LnsAny
    retAst = nil
    switch _exp4820 := uptodate.(type) {
    case *front_ModuleUptodate__Uptodate:
    metaInfo := _exp4820.Val1
        Util_errorLog("uptodate -- " + path)
        outputDependInfo(dependsStream, metaInfo)
    case *front_ModuleUptodate__NeedUpdate:
        if stream != nil && metaStream != nil{
            stream_6211 := stream.(Lns_oStream)
            metaStream_6212 := metaStream.(Lns_oStream)
            var ast *TransUnit_ASTInfo
            ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), front_createPaser_1099_(path, mod), mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
            retAst = ast
            
            if dependsStream != nil{
                dependsStream_6215 := dependsStream.(Lns_oStream)
                ast.FP.Get_node().FP.ProcessFilter(OutputDepend_createFilter(dependsStream_6215), 1)
            }
            var outStream Lns_oStream
            outStream = stream_6211
            var oMetaStream Lns_oStream
            oMetaStream = metaStream_6212
            var byteStream *Util_memStream
            byteStream = NewUtil_memStream()
            var byteMetaStream *Util_memStream
            byteMetaStream = NewUtil_memStream()
            if byteCompile{
                outStream = byteStream.FP
                
                oMetaStream = byteMetaStream.FP
                
            }
            self.FP.convert(ast, path, outStream, oMetaStream, convMode, false)
            if byteCompile{
                stream_6211.Write(front_byteCompileFromLuaTxt_1120_(byteStream.FP.Get_txt(), stripDebugInfo))
                if metaStream_6212 != stream_6211{
                    metaStream_6212.Write(front_byteCompileFromLuaTxt_1120_(byteMetaStream.FP.Get_txt(), true))
                }
            }
        } else {
            Util_err("failed to open lua stream or meta stream")
        }
    case *front_ModuleUptodate__NeedTouch:
    metaCode := _exp4820.Val1
    metaInfo := _exp4820.Val2
        Util_errorLog("touch -- " + path)
        if self.option.Mode == Option_ModeKind__SaveMeta{
            if metaStream != nil{
                metaStream_6229 := metaStream.(Lns_oStream)
                metaStream_6229.Write(metaCode)
            } else {
                Util_err("failed to open meta stream")
            }
        }
        outputDependInfo(dependsStream, metaInfo)
    }
    if Str_isValidStrBuilder(){
        if streamDst != nil{
            streamDst_6240 := streamDst.(Lns_oStream)
            streamDst_6240.Write(streamMem.FP.Get_txt())
        }
        
        if metaStreamDst != nil{
            metaStreamDst_6242 := metaStreamDst.(Lns_oStream)
            metaStreamDst_6242.Write(metaStreamMem.FP.Get_txt())
        }
        
        if dependsStreamDst != nil{
            dependsStreamDst_6244 := dependsStreamDst.(Lns_oStream)
            dependsStreamDst_6244.Write(dependsStreamMem.FP.Get_txt())
        }
        
    }
    if closeOStream != nil{
        closeOStream_6246 := closeOStream.(front_CloseOStreamForConvert_1320_)
        closeOStream_6246(stream, metaStream, dependsStream)
    }
    return retAst
}

// 954: decl @lune.@base.@front.Front.createGoOption
func (self *front_Front) createGoOption(scriptPath string) *ConvGo_Option {
    var packageName string
    {
        __exp := self.option.PackageName
        if __exp != nil {
            _exp := __exp.(string)
            packageName = _exp
            
        } else {
            if Lns_op_not(Lns_car(Lns_getVM().String_find(scriptPath,"/", nil, nil))){
                packageName = "main"
                
            } else { 
                var parentPath string
                parentPath = front_convExp5006(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(scriptPath,"/[^/]+$", "")).(string),".*/", "")))
                if len(parentPath) == 0{
                    packageName = "main"
                    
                } else if parentPath == "."{
                    packageName = "main"
                    
                } else if parentPath == ".."{
                    packageName = "main"
                    
                } else { 
                    packageName = front_convExp5063(Lns_2DDD(Lns_getVM().String_gsub(parentPath,"[^%w]", "")))
                    
                }
            }
        }
    }
    return NewConvGo_Option(packageName)
}

// 980: decl @lune.@base.@front.Front.convertToLua
func (self *front_Front) convertToLua(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    var convMode LnsInt
    convMode = ConvLua_ConvMode__Convert
    if self.option.Mode == Option_ModeKind__LuaMeta{
        convMode = ConvLua_ConvMode__ConvMeta
        
    }
    var parser *Parser_Parser
    parser = front_createPaser_1099_(scriptPath, mod)
    var ast LnsAny
    ast = self.FP.ConvertLuaToStreamFromScript(parser, FrontInterface_ModuleId__tempId, front_ModuleUptodate__NeedUpdate_Obj, convMode, scriptPath, mod, self.option.ByteCompile, self.option.StripDebugInfo, front_OpenOStreamForConvert_1317_(func(mode LnsAny)(LnsAny, LnsAny, LnsAny) {
        return Lns_io_stdout, Lns_io_stdout, self.option.FP.OpenDepend(nil)
    }), front_CloseOStreamForConvert_1320_(Front_convertToLua___anonymous_1346_))
    if ast != nil{
        ast_6284 := ast.(*TransUnit_ASTInfo)
        if _switch5270 := self.option.ConvTo; _switch5270 == Types_Conv__Go {
            var conv *Nodes_Filter
            conv = ConvGo_createFilter(self.option.Testing, "stdout", Lns_io_stdout, ast_6284, self.FP.createGoOption(scriptPath))
            ast_6284.FP.Get_node().FP.ProcessFilter(conv, ConvGo_Opt2Stem(NewConvGo_Opt(ast_6284.FP.Get_node())))
        }
    }
}

// 1019: decl @lune.@base.@front.Front.saveToGo
func (self *front_Front) SaveToGo(scriptPath string,ast *TransUnit_ASTInfo) {
    var path string
    path = front_convExp5297(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".go")))
    {
        _dir := self.option.OutputDir
        if _dir != nil {
            dir := _dir.(string)
            path = Lns_getVM().String_format("%s/%s", []LnsAny{dir, path})
            
        }
    }
    var file Lns_luaStream
    
    {
        _file := front_convExp5335(Lns_2DDD(Lns_io_open(path, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var memStream *Util_memStream
    memStream = NewUtil_memStream()
    var dstStream Lns_oStream
    if Str_isValidStrBuilder(){
        dstStream = memStream.FP
        
    } else { 
        dstStream = file
        
    }
    var conv *Nodes_Filter
    conv = ConvGo_createFilter(self.option.Testing, path, dstStream, ast, self.FP.createGoOption(scriptPath))
    ast.FP.Get_node().FP.ProcessFilter(conv, ConvGo_Opt2Stem(NewConvGo_Opt(ast.FP.Get_node())))
    if Str_isValidStrBuilder(){
        file.Write(memStream.FP.Get_txt())
    }
    file.Close()
}

// 1051: decl @lune.@base.@front.Front.saveToC
func (self *front_Front) SaveToC(scriptPath string,ast *TransUnit_ASTInfo) {
    var cPath string
    cPath = front_convExp5462(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".c")))
    var file Lns_luaStream
    
    {
        _file := front_convExp5477(Lns_2DDD(Lns_io_open(cPath, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var hPath string
    hPath = front_convExp5491(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".h")))
    var hFile Lns_luaStream
    
    {
        _hFile := front_convExp5507(Lns_2DDD(Lns_io_open(hPath, "w")))
        if _hFile == nil{
            return 
        } else {
            hFile = _hFile.(Lns_luaStream)
        }
    }
    var conv *Nodes_Filter
    conv = ConvCC_createFilter(self.option.Testing, self.option.Mode == Option_ModeKind__Builtin, cPath, file, hFile, ast)
    ast.FP.Get_node().FP.ProcessFilter(conv, ConvCC_Opt2Stem(NewConvCC_Opt(ast.FP.Get_node())))
    file.Close()
    hFile.Close()
}

// 1071: decl @lune.@base.@front.Front.outputBuiltin
func (self *front_Front) OutputBuiltin(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module("lns_builtin")
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), &NewParser_DummyParser().Parser_Parser, mod, FrontInterface_ModuleId_createId(0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    self.FP.SaveToC(scriptPath, ast)
}

// 1093: decl @lune.@base.@front.Front.saveToLua
func (self *front_Front) SaveToLua(updateInfo *front_UpdateInfo) bool {
    var scriptPath string
    scriptPath = updateInfo.FP.Get_scriptPath()
    var dependsPath LnsAny
    dependsPath = updateInfo.FP.Get_dependsPath()
    var parser LnsAny
    parser = updateInfo.FP.Get_parser()
    var moduleId *FrontInterface_ModuleId
    moduleId = updateInfo.FP.Get_moduleId()
    var uptodate LnsAny
    uptodate = updateInfo.FP.Get_uptodate()
    var checkDiff func(oldStream *Parser_TxtStream,newStream *Parser_TxtStream)(bool, string)
    checkDiff = func(oldStream *Parser_TxtStream,newStream *Parser_TxtStream)(bool, string) {
        __func__ := "@lune.@base.@front.Front.saveToLua.checkDiff"
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
            newLine = newStream.FP.Read("*l")
            var oldLine LnsAny
            oldLine = oldStream.FP.Read("*l")
            if oldLine != nil{
                oldLine_6371 := oldLine.(string)
                if len(oldBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(oldLine_6371,"^_moduleObj.__buildId", nil, nil))){
                        oldBuildIdLine = oldLine_6371
                        
                    }
                }
            }
            
            if newLine != nil{
                newLine_6375 := newLine.(string)
                if len(newBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(newLine_6375,"^_moduleObj.__buildId", nil, nil))){
                        newBuildIdLine = newLine_6375
                        
                    }
                }
            }
            
            if newLine != oldLine{
                var cont bool
                cont = false
                if newLine != nil && oldLine != nil{
                    newLine_6381 := newLine.(string)
                    oldLine_6382 := oldLine.(string)
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(oldLine_6382,"^_moduleObj.__buildId", nil, nil))){
                        if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(newLine_6381,"^_moduleObj.__buildId", nil, nil))){
                            tailBeginPos = newStream.FP.Get_lineNo()
                            
                            cont = true
                            
                        }
                    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(oldLine_6382,"^__dependModuleMap.*buildId =", nil, nil))) &&
                        Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(newLine_6381,"^__dependModuleMap.*buildId =", nil, nil))) )){
                        var oldSub string
                        oldSub = front_convExp5964(Lns_2DDD(Lns_getVM().String_gsub(oldLine_6382,"buildId =.*", "")))
                        var newSub string
                        newSub = front_convExp5977(Lns_2DDD(Lns_getVM().String_gsub(newLine_6381,"buildId =.*", "")))
                        if oldSub == newSub{
                            cont = true
                            
                        }
                    }
                }
                if Lns_op_not(cont){
                    Log_log(Log_Level__Debug, __func__, 1169, Log_CreateMessage(func() string {
                        return Lns_getVM().String_format("<%s>, <%s>", []LnsAny{oldLine, newLine})
                    }))
                    
                    return false, ""
                }
            } else { 
                if tailBeginPos == 0{
                    headEndPos = newStream.FP.Get_lineNo()
                    
                }
                if Lns_op_not(oldLine){
                    if tailBeginPos == 0{
                        return true, oldStream.FP.Get_txt()
                    }
                    var oldBuildId *FrontInterface_ModuleId
                    oldBuildId = Front_saveToLua__txt2ModuleId_1387_(oldBuildIdLine)
                    var newBuildId *FrontInterface_ModuleId
                    newBuildId = Front_saveToLua__txt2ModuleId_1387_(newBuildIdLine)
                    var worlBuildId *FrontInterface_ModuleId
                    worlBuildId = FrontInterface_ModuleId_createId(newBuildId.FP.Get_modTime(), oldBuildId.FP.Get_buildCount())
                    var buildIdLine string
                    buildIdLine = Lns_getVM().String_format("_moduleObj.__buildId = %q", []LnsAny{worlBuildId.FP.Get_idStr()})
                    var txt string
                    txt = Lns_getVM().String_format("%s%s\n%s", []LnsAny{newStream.FP.GetSubstring(1, headEndPos), buildIdLine, newStream.FP.GetSubstring(tailBeginPos, nil)})
                    return true, txt
                }
            }
        }
    // insert a dummy
        return false,""
    }
    var updateFlag bool
    updateFlag = true
    var ast LnsAny
    ast = nil
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    var luaPath string
    luaPath = front_convExp6214(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".lua")))
    var metaPath string
    metaPath = front_convExp6227(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".meta")))
    if Lns_isCondTrue( self.option.OutputDir){
        var filename string
        filename = front_convExp6244(Lns_2DDD(Lns_getVM().String_gsub(mod,"%.", "/")))
        luaPath = Lns_getVM().String_format("%s/%s.lua", []LnsAny{self.option.OutputDir, filename})
        
        metaPath = Lns_getVM().String_format("%s/%s.meta", []LnsAny{self.option.OutputDir, filename})
        
    }
    if luaPath == scriptPath{
        Util_errorLog(Lns_getVM().String_format("%s is illegal filename.", []LnsAny{luaPath}))
    } else { 
        var convMode LnsInt
        convMode = ConvLua_ConvMode__Convert
        if self.option.Mode == Option_ModeKind__SaveMeta{
            convMode = ConvLua_ConvMode__ConvMeta
            
        }
        var metaFileObj LnsAny
        metaFileObj = nil
        var tempMetaPath string
        tempMetaPath = metaPath + ".tmp"
        ast = self.FP.ConvertLuaToStreamFromScript(parser, moduleId, uptodate, convMode, scriptPath, mod, self.option.ByteCompile, self.option.StripDebugInfo, front_OpenOStreamForConvert_1317_(func(mode LnsAny)(LnsAny, LnsAny, LnsAny) {
            var openLuaStream func() LnsAny
            openLuaStream = func() LnsAny {
                var fileObj Lns_luaStream
                
                {
                    _fileObj := front_convExp6402(Lns_2DDD(Lns_io_open(luaPath, "w")))
                    if _fileObj == nil{
                        panic(Lns_getVM().String_format("write open error -- %s", []LnsAny{luaPath}))
                    } else {
                        fileObj = _fileObj.(Lns_luaStream)
                    }
                }
                return fileObj
            }
            var openStreams func(luaFlag bool)(LnsAny, LnsAny)
            openStreams = func(luaFlag bool)(LnsAny, LnsAny) {
                var stream LnsAny
                stream = nil
                if luaFlag{
                    stream = openLuaStream()
                    
                }
                var metaStream LnsAny
                metaStream = stream
                if self.option.Mode == Option_ModeKind__SaveMeta{
                    {
                        __exp := front_convExp6492(Lns_2DDD(Lns_io_open(tempMetaPath, "w+")))
                        if __exp != nil {
                            _exp := __exp.(Lns_luaStream)
                            metaFileObj = _exp
                            
                            metaStream = _exp
                            
                        } else {
                            panic(Lns_getVM().String_format("write open error -- %s", []LnsAny{metaPath}))
                        }
                    }
                }
                return stream, metaStream
            }
            var stream LnsAny
            stream = nil
            var metaStream LnsAny
            metaStream = stream
            switch mode.(type) {
            case *front_ModuleUptodate__Uptodate:
            case *front_ModuleUptodate__NeedTouch:
                stream, metaStream = openStreams(false)
                
            default:
                stream, metaStream = openStreams(true)
                
            }
            return stream, metaStream, self.option.FP.OpenDepend(dependsPath)
        }), front_CloseOStreamForConvert_1320_(func(stream LnsAny,metaStream LnsAny,dependStream LnsAny) {
            if stream != nil{
                stream_6448 := stream.(Lns_oStream)
                stream_6448.Close()
            }
            if dependStream != nil{
                dependStream_6450 := dependStream.(Lns_oStream)
                dependStream_6450.Close()
            }
            if metaFileObj != nil{
                metaFileObj_6452 := metaFileObj.(Lns_luaStream)
                metaFileObj_6452.Flush()
                metaFileObj_6452.Seek("set", 0)
                var newMetaTxt string
                
                {
                    _newMetaTxt := metaFileObj_6452.Read("*a")
                    if _newMetaTxt == nil{
                        Util_err(Lns_getVM().String_format("faled to read meta. -- %s.", []LnsAny{tempMetaPath}))
                    } else {
                        newMetaTxt = _newMetaTxt.(string)
                    }
                }
                metaFileObj_6452.Close()
                var oldMetaTxt string
                oldMetaTxt = ""
                {
                    _oldFileObj := front_convExp6687(Lns_2DDD(Lns_io_open(metaPath, nil)))
                    if _oldFileObj != nil {
                        oldFileObj := _oldFileObj.(Lns_luaStream)
                        oldMetaTxt = Lns_unwrapDefault( oldFileObj.Read("*a"), "").(string)
                        
                        oldFileObj.Close()
                    }
                }
                var sameFlag bool
                var txt string
                sameFlag,txt = checkDiff(NewParser_TxtStream(oldMetaTxt), NewParser_TxtStream(newMetaTxt))
                if Lns_op_not(sameFlag){
                    Lns_getVM().OS_rename(tempMetaPath, metaPath)
                } else { 
                    Lns_getVM().OS_remove(tempMetaPath)
                    if txt != ""{
                        {
                            _fileObj := front_convExp6768(Lns_2DDD(Lns_io_open(metaPath, "w")))
                            if _fileObj != nil {
                                fileObj := _fileObj.(Lns_luaStream)
                                fileObj.Write(txt)
                                fileObj.Close()
                            }
                        }
                    } else { 
                        updateFlag = false
                        
                    }
                }
            }
        }))
        
    }
    if updateFlag{
        Lns_getVM().String_gsub(scriptPath,"%.lns$", ".lua")
    }
    if ast != nil{
        ast_6469 := ast.(*TransUnit_ASTInfo)
        if _switch6849 := self.option.ConvTo; _switch6849 == Types_Conv__C {
            self.FP.SaveToC(scriptPath, ast_6469)
        } else if _switch6849 == Types_Conv__Go {
            self.FP.SaveToGo(scriptPath, ast_6469)
        }
    }
    return updateFlag
}

// 1345: decl @lune.@base.@front.Front.outputBootC
func (self *front_Front) outputBootC(scriptPath string) {
    var stream Lns_oStream
    {
        _path := self.option.BootPath
        if _path != nil {
            path := _path.(string)
            var cPath string
            cPath = front_convExp6885(Lns_2DDD(Lns_getVM().String_gsub(path,"%.lns$", ".c")))
            var file Lns_luaStream
            
            {
                _file := front_convExp6916(Lns_2DDD(Lns_io_open(cPath, "w")))
                if _file == nil{
                    Lns_print([]LnsAny{Lns_getVM().String_format("failed to open file -- %s", []LnsAny{cPath})})
                    return 
                } else {
                    file = _file.(Lns_luaStream)
                }
            }
            stream = file
            
        } else {
            stream = Lns_io_stdout
            
        }
    }
    var initModule string
    initModule = Front_scriptPath2Module(scriptPath)
    ConvCC_outputBootcode(stream, initModule)
}

// 1376: decl @lune.@base.@front.Front.exec
func (self *front_Front) Exec() {
    __func__ := "@lune.@base.@front.Front.exec"
    Log_log(Log_Level__Trace, __func__, 1378, Log_CreateMessage(func() string {
        return Option_ModeKind_getTxt( self.option.Mode)
    }))
    
    if _switch7581 := self.option.Mode; _switch7581 == Option_ModeKind__Token {
        self.FP.DumpTokenize(self.option.ScriptPath)
    } else if _switch7581 == Option_ModeKind__Ast {
        self.FP.DumpAst(self.option.ScriptPath)
    } else if _switch7581 == Option_ModeKind__Format {
        self.FP.Format(self.option.ScriptPath)
    } else if _switch7581 == Option_ModeKind__Diag {
        self.FP.CheckDiag(self.option.ScriptPath)
    } else if _switch7581 == Option_ModeKind__Complete {
        self.FP.Complete(self.option.ScriptPath)
    } else if _switch7581 == Option_ModeKind__Inquire {
        self.FP.Inquire(self.option.ScriptPath)
    } else if _switch7581 == Option_ModeKind__Glue {
        self.FP.CreateGlue(self.option.ScriptPath)
    } else if _switch7581 == Option_ModeKind__Lua || _switch7581 == Option_ModeKind__LuaMeta {
        self.FP.convertToLua(self.option.ScriptPath)
    } else if _switch7581 == Option_ModeKind__Save || _switch7581 == Option_ModeKind__SaveMeta {
        var createUpdateInfo func(scriptPath string,dependsPath LnsAny) *front_UpdateInfo
        createUpdateInfo = func(scriptPath string,dependsPath LnsAny) *front_UpdateInfo {
            var mod string
            mod = Front_scriptPath2Module(scriptPath)
            var moduleId *FrontInterface_ModuleId
            var uptodate LnsAny
            moduleId,uptodate = self.FP.getModuleIdAndCheckUptodate(scriptPath, mod)
            var parser LnsAny
            switch uptodate.(type) {
            case *front_ModuleUptodate__NeedUpdate:
                parser = front_createPaser_1099_(scriptPath, mod)
                
            case *front_ModuleUptodate__NeedTouch:
                parser = nil
                
            case *front_ModuleUptodate__Uptodate:
                parser = nil
                
            }
            return Newfront_UpdateInfo(scriptPath, dependsPath, parser, moduleId, uptodate)
        }
        Depend_profile(self.option.ValidProf, conv2Form7446(func() {
            if self.option.ScriptPath == "@-"{
                var updateList *LnsList
                updateList = NewLnsList([]LnsAny{})
                for  {
                    var line string
                    
                    {
                        _line := Lns_io_stdin.Read("*l")
                        if _line == nil{
                            break
                        } else {
                            line = _line.(string)
                        }
                    }
                    if len(line) > 0{
                        updateList.Insert(front_UpdateInfo2Stem(createUpdateInfo(line, Lns_car(Lns_getVM().String_gsub(line,".lns$", ".d")).(string))))
                    }
                }
                for _, _updateInfo := range( updateList.Items ) {
                    updateInfo := _updateInfo.(front_UpdateInfoDownCast).Tofront_UpdateInfo()
                    var prev LnsReal
                    prev = Lns_getVM().OS_clock()
                    self.FP.SaveToLua(updateInfo)
                    Lns_print([]LnsAny{Lns_getVM().String_format("%s:%g", []LnsAny{updateInfo.FP.Get_scriptPath(), Lns_getVM().OS_clock() - prev})})
                }
            } else { 
                self.FP.SaveToLua(createUpdateInfo(self.option.ScriptPath, nil))
            }
        }), self.option.ScriptPath + ".profi")
    } else if _switch7581 == Option_ModeKind__Exec {
        _ = front_convExp7471(Lns_2DDD(self.FP.LoadModule(Front_scriptPath2Module(self.option.ScriptPath))))
        if self.option.Testing{
            var code string
            code = "local Testing = require( \"lune.base.Testing\" )\nreturn function( path )\n  Testing.run( path );\n  Testing.outputAllResult( io.stdout );\nend\n"
            var loaded LnsAny
            var mess LnsAny
            loaded,mess = Lns_getVM().Load(code, nil)
            if loaded != nil{
                loaded_6543 := loaded.(*Lns_luaValue)
                {
                    _mod := front_convExp7524(Lns_2DDD(Lns_getVM().RunLoadedfunc(loaded_6543,Lns_2DDD([]LnsAny{}))[0]))
                    if _mod != nil {
                        mod := _mod
                        Lns_getVM().RunLoadedfunc((mod.(*Lns_luaValue)),Lns_2DDD([]LnsAny{Front_scriptPath2Module(self.option.ScriptPath)}))
                    }
                }
            } else {
                Lns_print([]LnsAny{mess})
            }
        }
    } else if _switch7581 == Option_ModeKind__BootC {
        self.FP.outputBootC(self.option.ScriptPath)
    } else if _switch7581 == Option_ModeKind__Builtin {
        self.FP.OutputBuiltin(self.option.ScriptPath)
    } else {
        Lns_print([]LnsAny{"illegal mode"})
    }
}


// declaration Class -- DependMetaInfo
type front_DependMetaInfoMtd interface {
    ToMap() *LnsMap
}
type front_DependMetaInfo struct {
    Use bool
    BuildId string
    FP front_DependMetaInfoMtd
}
func front_DependMetaInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_DependMetaInfo).FP
}
type front_DependMetaInfoDownCast interface {
    Tofront_DependMetaInfo() *front_DependMetaInfo
}
func front_DependMetaInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(front_DependMetaInfoDownCast)
    if ok { return work.Tofront_DependMetaInfo() }
    return nil
}
func (obj *front_DependMetaInfo) Tofront_DependMetaInfo() *front_DependMetaInfo {
    return obj
}
func Newfront_DependMetaInfo(arg1 bool, arg2 string) *front_DependMetaInfo {
    obj := &front_DependMetaInfo{}
    obj.FP = obj
    obj.Initfront_DependMetaInfo(arg1, arg2)
    return obj
}
func (self *front_DependMetaInfo) Initfront_DependMetaInfo(arg1 bool, arg2 string) {
    self.Use = arg1
    self.BuildId = arg2
}
func (self *front_DependMetaInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["use"] = Lns_ToCollection( self.Use )
    obj.Items["buildId"] = Lns_ToCollection( self.BuildId )
    return obj
}
func (self *front_DependMetaInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func front_DependMetaInfo__fromMap_1143_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_DependMetaInfo_FromMap( arg1, paramList )
}
func front_DependMetaInfo__fromStem_1146_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_DependMetaInfo_FromMap( arg1, paramList )
}
func front_DependMetaInfo_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := front_DependMetaInfo_FromMapSub(obj,false, paramList);
    return conv,mess
}
func front_DependMetaInfo_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &front_DependMetaInfo{}
    newObj.FP = newObj
    return front_DependMetaInfo_FromMapMain( newObj, objMap, paramList )
}
func front_DependMetaInfo_FromMapMain( newObj *front_DependMetaInfo, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["use"], false, nil); !ok {
       return false,nil,"use:" + mess.(string)
    } else {
       newObj.Use = conv.(bool)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["buildId"], false, nil); !ok {
       return false,nil,"buildId:" + mess.(string)
    } else {
       newObj.BuildId = conv.(string)
    }
    return true, newObj, nil
}

// declaration Class -- MetaForBuildId
type front_MetaForBuildIdMtd interface {
    ToMap() *LnsMap
    CreateModuleId() *FrontInterface_ModuleId
}
type front_MetaForBuildId struct {
    __buildId string
    __dependModuleMap *LnsMap
    __subModuleMap *LnsList
    __enableTest bool
    FP front_MetaForBuildIdMtd
}
func front_MetaForBuildId2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_MetaForBuildId).FP
}
type front_MetaForBuildIdDownCast interface {
    Tofront_MetaForBuildId() *front_MetaForBuildId
}
func front_MetaForBuildIdDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(front_MetaForBuildIdDownCast)
    if ok { return work.Tofront_MetaForBuildId() }
    return nil
}
func (obj *front_MetaForBuildId) Tofront_MetaForBuildId() *front_MetaForBuildId {
    return obj
}
func Newfront_MetaForBuildId(arg1 string, arg2 *LnsMap, arg3 *LnsList, arg4 bool) *front_MetaForBuildId {
    obj := &front_MetaForBuildId{}
    obj.FP = obj
    obj.Initfront_MetaForBuildId(arg1, arg2, arg3, arg4)
    return obj
}
func (self *front_MetaForBuildId) Initfront_MetaForBuildId(arg1 string, arg2 *LnsMap, arg3 *LnsList, arg4 bool) {
    self.__buildId = arg1
    self.__dependModuleMap = arg2
    self.__subModuleMap = arg3
    self.__enableTest = arg4
}
func (self *front_MetaForBuildId) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["__buildId"] = Lns_ToCollection( self.__buildId )
    obj.Items["__dependModuleMap"] = Lns_ToCollection( self.__dependModuleMap )
    obj.Items["__subModuleMap"] = Lns_ToCollection( self.__subModuleMap )
    obj.Items["__enableTest"] = Lns_ToCollection( self.__enableTest )
    return obj
}
func (self *front_MetaForBuildId) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func front_MetaForBuildId__fromMap_1176_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_MetaForBuildId_FromMap( arg1, paramList )
}
func front_MetaForBuildId__fromStem_1179_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_MetaForBuildId_FromMap( arg1, paramList )
}
func front_MetaForBuildId_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := front_MetaForBuildId_FromMapSub(obj,false, paramList);
    return conv,mess
}
func front_MetaForBuildId_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &front_MetaForBuildId{}
    newObj.FP = newObj
    return front_MetaForBuildId_FromMapMain( newObj, objMap, paramList )
}
func front_MetaForBuildId_FromMapMain( newObj *front_MetaForBuildId, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["__buildId"], false, nil); !ok {
       return false,nil,"__buildId:" + mess.(string)
    } else {
       newObj.__buildId = conv.(string)
    }
    if ok,conv,mess := Lns_ToLnsMapSub( objMap.Items["__dependModuleMap"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil},Lns_ToObjParam{
            front_DependMetaInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"__dependModuleMap:" + mess.(string)
    } else {
       newObj.__dependModuleMap = conv.(*LnsMap)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["__subModuleMap"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil}}); !ok {
       return false,nil,"__subModuleMap:" + mess.(string)
    } else {
       newObj.__subModuleMap = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["__enableTest"], false, nil); !ok {
       return false,nil,"__enableTest:" + mess.(string)
    } else {
       newObj.__enableTest = conv.(bool)
    }
    return true, newObj, nil
}
// 256: decl @lune.@base.@front.MetaForBuildId.createModuleId
func (self *front_MetaForBuildId) CreateModuleId() *FrontInterface_ModuleId {
    return FrontInterface_ModuleId_createIdFromTxt(self.__buildId)
}

// 261: decl @lune.@base.@front.MetaForBuildId.LoadFromMeta
func front_MetaForBuildId_LoadFromMeta_1183_(metaPath string)(LnsAny, LnsAny) {
    {
        _fileObj := front_convExp1165(Lns_2DDD(Lns_io_open(metaPath, nil)))
        if _fileObj != nil {
            fileObj := _fileObj.(Lns_luaStream)
            var luaCode LnsAny
            luaCode = fileObj.Read("*a")
            fileObj.Close()
            if luaCode != nil{
                luaCode_5823 := luaCode.(string)
                var meta LnsAny
                meta = front_convExp1155(Lns_2DDD(front_MetaForBuildId__fromStem_1179_(Lns_getVM().ExpandLuavalMap(front_loadFromLuaTxt_1117_(luaCode_5823)),nil)))
                return meta, luaCode_5823
            }
        }
    }
    return nil, nil
}


// declaration Class -- UpdateInfo
type front_UpdateInfoMtd interface {
    Get_dependsPath() LnsAny
    Get_moduleId() *FrontInterface_ModuleId
    Get_parser() LnsAny
    Get_scriptPath() string
    Get_uptodate() LnsAny
}
type front_UpdateInfo struct {
    scriptPath string
    dependsPath LnsAny
    parser LnsAny
    moduleId *FrontInterface_ModuleId
    uptodate LnsAny
    FP front_UpdateInfoMtd
}
func front_UpdateInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_UpdateInfo).FP
}
type front_UpdateInfoDownCast interface {
    Tofront_UpdateInfo() *front_UpdateInfo
}
func front_UpdateInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(front_UpdateInfoDownCast)
    if ok { return work.Tofront_UpdateInfo() }
    return nil
}
func (obj *front_UpdateInfo) Tofront_UpdateInfo() *front_UpdateInfo {
    return obj
}
func Newfront_UpdateInfo(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 *FrontInterface_ModuleId, arg5 LnsAny) *front_UpdateInfo {
    obj := &front_UpdateInfo{}
    obj.FP = obj
    obj.Initfront_UpdateInfo(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *front_UpdateInfo) Initfront_UpdateInfo(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 *FrontInterface_ModuleId, arg5 LnsAny) {
    self.scriptPath = arg1
    self.dependsPath = arg2
    self.parser = arg3
    self.moduleId = arg4
    self.uptodate = arg5
}
func (self *front_UpdateInfo) Get_scriptPath() string{ return self.scriptPath }
func (self *front_UpdateInfo) Get_dependsPath() LnsAny{ return self.dependsPath }
func (self *front_UpdateInfo) Get_parser() LnsAny{ return self.parser }
func (self *front_UpdateInfo) Get_moduleId() *FrontInterface_ModuleId{ return self.moduleId }
func (self *front_UpdateInfo) Get_uptodate() LnsAny{ return self.uptodate }

func Lns_front_init() {
    if init_front { return }
    init_front = true
    front__mod__ = "@lune.@base.@front"
    Lns_InitMod()
    Lns_Types_init()
    Lns_Str_init()
    Lns_frontInterface_init()
    Lns_Parser_init()
    Lns_convLua_init()
    Lns_convCC_init()
    Lns_convGo_init()
    Lns_TransUnit_init()
    Lns_Util_init()
    Lns_Option_init()
    Lns_dumpNode_init()
    Lns_glueFilter_init()
    Lns_LuaVer_init()
    Lns_Depend_init()
    Lns_OutputDepend_init()
    Lns_Ver_init()
    Lns_LuaVer_init()
    Lns_Log_init()
    Lns_Formatter_init()
    Lns_Testing_init()
    Lns_Meta_init()
    Depend_setup(Depend_UpdateVer(front__anonymous_1007_))
    front_forceUpdateMeta = true
}
func init() {
    init_front = false
}
