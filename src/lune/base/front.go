// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_front bool
var front__mod__ string
var front_forceUpdateMeta bool
// decl alge -- ModuleUptodate
type front_ModuleUptodate = LnsAny
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
type front_OpenOStreamForConvert_1321_ func (arg1 LnsAny)(LnsAny, LnsAny, LnsAny)
type front_CloseOStreamForConvert_1324_ func (arg1 LnsAny,arg2 LnsAny,arg3 LnsAny)
// for 800: ExpCast
func conv2Form4120( src func ()) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        src()
        return []LnsAny{}
    }
}
// for 1485: ExpCast
func conv2Form7581( src func ()) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        src()
        return []LnsAny{}
    }
}
// for 1530: ExpCast
func conv2Form7686( src func ()) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        src()
        return []LnsAny{}
    }
}
// for 229
func front_convExp791(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 298
func front_convExp1183(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 336
func front_convExp1399(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 550
func front_convExp2452(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 558
func front_convExp2481(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1312
func front_convExp6627(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1358
func front_convExp6822(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1372
func front_convExp6903(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1510
func front_convExp7629(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1529
func front_convExp7703(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 316
func front_convExp1256(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 348
func front_convExp1456(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 349
func front_convExp1475(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 378
func front_convExp1566(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 543
func front_convExp2389(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1032
func front_convExp5190(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 302
func front_convExp1173(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 445
func front_convExp1835(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 646
func front_convExp3026(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 687
func front_convExp3327(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 696
func front_convExp3385(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 742
func front_convExp3625(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 751
func front_convExp3683(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1023
func front_convExp5133(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1079
func front_convExp5432(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1084
func front_convExp5470(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1111
func front_convExp5597(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1113
func front_convExp5612(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1117
func front_convExp5626(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1118
func front_convExp5642(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1161
func front_convExp5848(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1220
func front_convExp6099(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1221
func front_convExp6112(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1262
func front_convExp6349(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1263
func front_convExp6362(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1265
func front_convExp6379(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1297
func front_convExp6537(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1408
func front_convExp7020(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1409
func front_convExp7051(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1517
func front_convExp7650(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
func front__anonymous_1007_(ver LnsInt) {
    LuaVer_setCurVer(ver)
}
// 180: decl @lune.@base.@front.createPaser
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

// 189: decl @lune.@base.@front.scriptPath2Module
func Front_scriptPath2Module(path string) string {
    return Util_scriptPath2Module(path)
}

// 224: decl @lune.@base.@front.loadFromChunk
func front_loadFromChunk_1114_(chunk LnsAny,err LnsAny) LnsAny {
    if err != nil{
        err_5871 := err.(string)
        Util_errorLog(err_5871)
    }
    if chunk != nil{
        chunk_5873 := chunk.(*Lns_luaValue)
        {
            _work := front_convExp791(Lns_2DDD(Lns_getVM().RunLoadedfunc(chunk_5873,Lns_2DDD([]LnsAny{}))[0]))
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

// 237: decl @lune.@base.@front.loadFromLuaTxt
func front_loadFromLuaTxt_1117_(txt string) LnsAny {
    return front_loadFromChunk_1114_(Lns_getVM().Load(txt, nil))
}

// 242: decl @lune.@base.@front.byteCompileFromLuaTxt
func front_byteCompileFromLuaTxt_1120_(txt string,stripDebugInfo bool) string {
    var chunk LnsAny
    var err LnsAny
    chunk,err = Lns_getVM().Load(txt, nil)
    if chunk != nil{
        chunk_5888 := chunk.(*Lns_luaValue)
        return Lns_getVM().String_dump(chunk_5888, stripDebugInfo)
    }
    panic(Lns_unwrapDefault( err, "load error").(string))
// insert a dummy
    return ""
}

// 311: decl @lune.@base.@front.getMetaInfo
func front_getMetaInfo_1186_(lnsPath string,mod string,outdir LnsAny)(LnsAny, string, string) {
    var moduleMetaPath string
    moduleMetaPath = lnsPath
    if outdir != nil{
        outdir_5949 := outdir.(string)
        moduleMetaPath = Lns_getVM().String_format("%s/%s", []LnsAny{outdir_5949, Lns_car(Lns_getVM().String_gsub(mod,"%.", "/")).(string)})
        
    }
    moduleMetaPath = front_convExp1256(Lns_2DDD(Lns_getVM().String_gsub(moduleMetaPath,"%.lns$", ".meta")))
    
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

// 370: decl @lune.@base.@front.getModuleId
func front_getModuleId_1196_(lnsPath string,mod string,outdir LnsAny,metaInfo LnsAny) *FrontInterface_ModuleId {
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
        metaInfo = front_convExp1566(Lns_2DDD(front_getMetaInfo_1186_(lnsPath, mod, outdir)))
        
    }
    if metaInfo != nil{
        metaInfo_5988 := metaInfo.(*front_MetaForBuildId)
        var buildId *FrontInterface_ModuleId
        buildId = metaInfo_5988.FP.CreateModuleId()
        buildCount = buildId.FP.Get_buildCount()
        
    }
    return FrontInterface_ModuleId_createId(fileTime, buildCount)
}

func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1210_() string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1213_() string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1216_() string {
    return "NeedUpdate"
}


func Front_getModuleIdAndCheckUptodate___anonymous_1222_() string {
    return "not found meta"
}















func Front_convertToLua___anonymous_1350_(stream LnsAny,metaStream LnsAny,dependStream LnsAny) {
    if dependStream != nil{
        dependStream_6406 := dependStream.(Lns_oStream)
        dependStream_6406.Close()
    }
}
// 1160: decl @lune.@base.@front.Front.saveToLua.txt2ModuleId
func Front_saveToLua__txt2ModuleId_1391_(txt string) *FrontInterface_ModuleId {
    var buildIdTxt string
    buildIdTxt = front_convExp5848(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(txt,"^_moduleObj.__buildId = ", "")).(string),"\"", "")))
    return FrontInterface_ModuleId_createIdFromTxt(buildIdTxt)
}







// 1422: decl @lune.@base.@front.convertLnsCode2LuaCode
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




// 1558: decl @lune.@base.@front.exec
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

// 1575: decl @lune.@base.@front.setFront
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
    ConvertLuaToStreamFromScript(arg1 LnsAny, arg2 *FrontInterface_ModuleId, arg3 LnsAny, arg4 LnsInt, arg5 string, arg6 string, arg7 bool, arg8 bool, arg9 front_OpenOStreamForConvert_1321_, arg10 LnsAny) LnsAny
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
    gomodMap *GoMod_ModInfo
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
// 115: DeclConstr
func (self *front_Front) Initfront_Front(option *Option_Option) {
    self.gomodMap = GoMod_getGoMap(option)
    
    self.option = option
    
    self.loadedMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.loadedMapTest = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.loadedMetaMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.convertedMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    FrontInterface_setFront(self.FP)
    var loadedMap *LnsMap
    loadedMap = NewLnsMap( map[LnsAny]LnsAny{})
    {
        _exp260 := Depend_getLoadedMod()
        _key260, _val260 := _exp260.Get1stFromMap()
        for _key260 != nil {
            mod := _key260.(string)
            modval := _val260
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
            _key260, _val260 = _exp260.NextFromMap( _key260 )
        }
    }
    self.preloadedModMap = loadedMap
    
}

// 141: decl @lune.@base.@front.Front.getLoadInfo
func (self *front_Front) getLoadInfo(mod string) LnsAny {
    if self.option.Testing{
        return self.loadedMapTest.Items[mod]
    }
    return self.loadedMap.Items[mod]
}

// 148: decl @lune.@base.@front.Front.setLoadInfo
func (self *front_Front) setLoadInfo(mod string,info *front_LoadInfo) {
    if self.option.Testing{
        self.loadedMapTest.Set(mod,info)
    }
    self.loadedMap.Set(mod,info)
}






// 165: decl @lune.@base.@front.Front.error
func (self *front_Front) Error(message string) {
    Util_errorLog(message)
    Util_printStackTrace()
    Lns_getVM().OS_exit(1)
}

// 171: decl @lune.@base.@front.Front.loadLua
func (self *front_Front) loadLua(path string) LnsAny {
    var chunk LnsAny
    var err LnsAny
    chunk,err = Lns_getVM().Loadfile(path)
    if chunk != nil{
        chunk_5823 := chunk.(*Lns_luaValue)
        return Lns_unwrap( Lns_car(Lns_getVM().RunLoadedfunc(chunk_5823,Lns_2DDD([]LnsAny{}))[0]))
    }
    Util_errorLog(Lns_unwrapDefault( err, Lns_getVM().String_format("load error -- %s.", []LnsAny{path})).(string))
    return nil
}

// 193: decl @lune.@base.@front.Front.createPaser
func (self *front_Front) createPaser(scriptPath string) *Parser_Parser {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    return front_createPaser_1099_(scriptPath, mod)
}

// 200: decl @lune.@base.@front.Front.createAst
func (self *front_Front) createAst(importModuleInfo *FrontInterface_ImportModuleInfo,parser *Parser_Parser,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *TransUnit_ASTInfo {
    var transUnit *TransUnit_TransUnit
    transUnit = NewTransUnit_TransUnit(moduleId, importModuleInfo, &NewConvLua_MacroEvalImp().Nodes_MacroEval, analyzeModule, analyzeMode, pos, self.option.TargetLuaVer, self.option.TransCtrlInfo)
    return transUnit.FP.CreateAST(parser, false, mod)
}

// 212: decl @lune.@base.@front.Front.convert
func (self *front_Front) convert(ast *TransUnit_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool) {
    var conv *Nodes_Filter
    conv = ConvLua_createFilter(streamName, stream, metaStream, convMode, inMacro, ast.FP.Get_moduleTypeInfo(), ast.FP.Get_processInfo(), ast.FP.Get_moduleSymbolKind(), self.option.UseLuneModule, self.option.TargetLuaVer, self.option.Testing, self.option.UseIpairs)
    ast.FP.Get_node().FP.ProcessFilter(conv, ConvLua_Opt2Stem(NewConvLua_Opt(ast.FP.Get_node())))
}

// 251: decl @lune.@base.@front.Front.convertFromAst
func (self *front_Front) convertFromAst(ast *TransUnit_ASTInfo,streamName string,mode LnsInt)(string, string) {
    var stream *Util_memStream
    stream = NewUtil_memStream()
    var metaStream *Util_memStream
    metaStream = NewUtil_memStream()
    self.FP.convert(ast, streamName, stream.FP, metaStream.FP, mode, false)
    return metaStream.FP.Get_txt(), stream.FP.Get_txt()
}

// 265: decl @lune.@base.@front.Front.loadFromLnsTxt
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

// 323: decl @lune.@base.@front.Front.searchModuleFile
func (self *front_Front) searchModuleFile(mod string,suffix string,addPath LnsAny) LnsAny {
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(mod,"^go/", nil, nil))){
        var workMod string
        workMod = Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(mod,"^go/", "")).(string),"%.", "/")).(string),":", ".")).(string) + suffix
        var pathList *LnsList
        pathList = NewLnsList([]LnsAny{Util_pathJoin("vendor", workMod)})
        {
            _gopath := self.gomodMap.FP.GetModulePath(workMod)
            if _gopath != nil {
                gopath := _gopath.(string)
                pathList.Insert(gopath)
            }
        }
        for _, _path := range( pathList.Items ) {
            path := _path.(string)
            {
                _fileObj := front_convExp1399(Lns_2DDD(Lns_io_open(path, nil)))
                if _fileObj != nil {
                    fileObj := _fileObj.(Lns_luaStream)
                    fileObj.Close()
                    return path
                }
            }
        }
        return nil
    }
    var lnsSearchPath string
    lnsSearchPath = Lns_package_path
    if addPath != nil{
        addPath_5971 := addPath.(string)
        lnsSearchPath = Lns_getVM().String_format("%s/?%s;%s", []LnsAny{addPath_5971, suffix, Lns_package_path})
        
    }
    lnsSearchPath = front_convExp1456(Lns_2DDD(Lns_getVM().String_gsub(lnsSearchPath,"%.lua$", suffix)))
    
    lnsSearchPath = front_convExp1475(Lns_2DDD(Lns_getVM().String_gsub(lnsSearchPath,"%.lua;", suffix + ";")))
    
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

// 410: decl @lune.@base.@front.Front.getModuleIdAndCheckUptodate
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
        for _depMod, _dependItem := range( metaInfo.G__dependModuleMap.Items ) {
            depMod := _depMod.(string)
            dependItem := _dependItem.(front_DependMetaInfoDownCast).Tofront_DependMetaInfo()
            var modMetaPath string
            
            {
                _modMetaPath := self.FP.searchModuleFile(depMod, ".meta", self.option.OutputDir)
                if _modMetaPath == nil{
                    Log_log(Log_Level__Debug, __func__, 433, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1210_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    modMetaPath = _modMetaPath.(string)
                }
            }
            var time LnsReal
            
            {
                _time := Depend_getFileLastModifiedTime(modMetaPath)
                if _time == nil{
                    Log_log(Log_Level__Debug, __func__, 438, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1213_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    time = _time.(LnsReal)
                }
            }
            if time > metaTime{
                var dependMeta *front_MetaForBuildId
                
                {
                    _dependMeta := front_convExp1835(Lns_2DDD(front_MetaForBuildId_LoadFromMeta_1183_(modMetaPath)))
                    if _dependMeta == nil{
                        Log_log(Log_Level__Debug, __func__, 446, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1216_))
                        
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
                    Log_log(Log_Level__Debug, __func__, 456, Log_CreateMessage(func() string {
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
        metaInfo_6038 := metaInfo.(*front_MetaForBuildId)
        if metaInfo_6038.G__enableTest == self.option.Testing{
            var buildId *FrontInterface_ModuleId
            buildId = FrontInterface_ModuleId_createIdFromTxt(metaInfo_6038.G__buildId)
            if buildId != FrontInterface_ModuleId__tempId{
                var lnsTime LnsAny
                lnsTime = Depend_getFileLastModifiedTime(lnsPath)
                var metaTime LnsAny
                metaTime = Depend_getFileLastModifiedTime(metaPath)
                if lnsTime != nil && metaTime != nil{
                    lnsTime_6045 := lnsTime.(LnsReal)
                    metaTime_6046 := metaTime.(LnsReal)
                    if lnsTime_6045 == buildId.FP.Get_modTime(){
                        uptodate = checkDependUptodate(metaTime_6046, metaInfo_6038, metaCode)
                        
                    }
                }
            }
        } else { 
        }
    } else {
        Log_log(Log_Level__Debug, __func__, 494, Log_CreateMessage(Front_getModuleIdAndCheckUptodate___anonymous_1222_))
        
    }
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_1196_(lnsPath, mod, self.option.OutputDir, metaInfo)
    if moduleId == FrontInterface_ModuleId__tempId{
        Util_err(Lns_getVM().String_format("not found -- %s", []LnsAny{lnsPath}))
    }
    return moduleId, uptodate
}

// 507: decl @lune.@base.@front.Front.convertLns2LuaCode
func (self *front_Front) ConvertLns2LuaCode(importModuleInfo *FrontInterface_ImportModuleInfo,stream Lns_iStream,streamName string) string {
    var mod string
    mod = Front_scriptPath2Module(streamName)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(importModuleInfo, &NewParser_StreamParser(stream, streamName, false).Parser_Parser, mod, FrontInterface_ModuleId_createId(0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(ast, streamName, ConvLua_ConvMode__Exec)
    return luaTxt
}

// 526: decl @lune.@base.@front.Front.loadFileToLuaCode
func (self *front_Front) loadFileToLuaCode(importModuleInfo *FrontInterface_ImportModuleInfo,path string,mod string)(LnsAny, string) {
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(importModuleInfo, front_createPaser_1099_(path, mod), mod, front_getModuleId_1196_(path, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
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
                    newpath = front_convExp2389(Lns_2DDD(Lns_getVM().String_gsub(path,".lns$", suffix)))
                    
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
                    _fileObj := front_convExp2452(Lns_2DDD(Lns_io_open(newpath, nil)))
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
                _fileObj := front_convExp2481(Lns_2DDD(Lns_io_open(newpath, "w")))
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

// 578: decl @lune.@base.@front.Front.loadFile
func (self *front_Front) loadFile(importModuleInfo *FrontInterface_ImportModuleInfo,path string,mod string)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@front.Front.loadFile"
    Log_log(Log_Level__Info, __func__, 582, Log_CreateMessage(func() string {
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

// 594: decl @lune.@base.@front.Front.searchModule
func (self *front_Front) SearchModule(mod string) LnsAny {
    return self.FP.searchModuleFile(mod, ".lns", nil)
}

// 598: decl @lune.@base.@front.Front.searchLuaFile
func (self *front_Front) searchLuaFile(moduleFullName string,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(moduleFullName, ".lua", addSearchPath)
}

// 618: decl @lune.@base.@front.Front.checkUptodateMeta
func (self *front_Front) checkUptodateMeta(metaPath string,addSearchPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.checkUptodateMeta"
    var metaObj LnsAny
    
    {
        _metaObj := self.FP.loadLua(metaPath)
        if _metaObj == nil{
            Log_log(Log_Level__Warn, __func__, 620, Log_CreateMessage(func() string {
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
        Log_log(Log_Level__Warn, __func__, 625, Log_CreateMessage(func() string {
            return Lns_getVM().String_format("unmatch meta version -- %s", []LnsAny{metaPath})
        }))
        
        return nil
    }
    if meta.GetAt( "__hasTest" ).(bool){
        if meta.GetAt( "__enableTest" ).(bool) != self.option.Testing{
            Log_log(Log_Level__Warn, __func__, 631, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("unmatch test setting -- %s", []LnsAny{metaPath})
            }))
            
            return nil
        }
    }
    {
        _exp3215 := meta.GetAt( "__dependModuleMap" ).(*Lns_luaValue)
        _key3215, _ := _exp3215.Get1stFromMap()
        for _key3215 != nil {
            moduleFullName := _key3215.(string)
            {
                _lnsPath := self.FP.SearchModule(moduleFullName)
                if _lnsPath != nil {
                    lnsPath := _lnsPath.(string)
                    {
                        _moduleLuaPath := self.FP.searchLuaFile(moduleFullName, addSearchPath)
                        if _moduleLuaPath != nil {
                            moduleLuaPath := _moduleLuaPath.(string)
                            if Lns_op_not(Util_getReadyCode(lnsPath, metaPath)){
                                Log_log(Log_Level__Warn, __func__, 642, Log_CreateMessage(func() string {
                                    return Lns_getVM().String_format("not ready -- %s, %s", []LnsAny{lnsPath, metaPath})
                                }))
                                
                                return nil
                            }
                            var moduleMetaPath string
                            moduleMetaPath = front_convExp3026(Lns_2DDD(Lns_getVM().String_gsub(moduleLuaPath,"%.lua$", ".meta")))
                            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                Lns_GetEnv().SetStackVal( Depend_existFile(moduleMetaPath)) &&
                                Lns_GetEnv().SetStackVal( Lns_op_not(Util_getReadyCode(moduleMetaPath, metaPath))) ).(bool)){
                                Log_log(Log_Level__Warn, __func__, 650, Log_CreateMessage(func() string {
                                    return Lns_getVM().String_format("not ready -- %s, %s", []LnsAny{moduleMetaPath, metaPath})
                                }))
                                
                                return nil
                            }
                        } else {
                            Log_log(Log_Level__Warn, __func__, 655, Log_CreateMessage(func() string {
                                return Lns_getVM().String_format("not found .lua file for -- %s", []LnsAny{moduleFullName})
                            }))
                            
                            return nil
                        }
                    }
                } else {
                    Log_log(Log_Level__Warn, __func__, 660, Log_CreateMessage(func() string {
                        return Lns_getVM().String_format("not found .lns file -- %s", []LnsAny{moduleFullName})
                    }))
                    
                    return nil
                }
            }
            _key3215, _ = _exp3215.NextFromMap( _key3215 )
        }
    }
    return meta
}

// 674: decl @lune.@base.@front.Front.loadModule
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
                        luaPath = front_convExp3327(Lns_2DDD(Lns_getVM().String_gsub(lnsPath, "%.lns$", ".lua")))
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
                            luaPath_6180 := luaPath.(string)
                            if Util_getReadyCode(lnsPath, luaPath_6180){
                                var metaPath string
                                metaPath = front_convExp3385(Lns_2DDD(Lns_getVM().String_gsub(luaPath_6180, "%.lua$", ".meta")))
                                if Util_getReadyCode(lnsPath, metaPath){
                                    {
                                        _preLoadInfo := self.preloadedModMap.Items[mod]
                                        if _preLoadInfo != nil {
                                            preLoadInfo := _preLoadInfo
                                            loadVal = preLoadInfo
                                            
                                        } else {
                                            loadVal = self.FP.loadLua(luaPath_6180)
                                            
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

// 733: decl @lune.@base.@front.Front.loadMeta
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
                        luaPath = front_convExp3625(Lns_2DDD(Lns_getVM().String_gsub(lnsPath, "%.lns$", ".lua")))
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
                            luaPath_6214 := luaPath.(string)
                            if Util_getReadyCode(lnsPath, luaPath_6214){
                                var metaPath string
                                metaPath = front_convExp3683(Lns_2DDD(Lns_getVM().String_gsub(luaPath_6214, "%.lua$", ".meta")))
                                if Util_getReadyCode(lnsPath, metaPath){
                                    meta = self.FP.checkUptodateMeta(metaPath, self.option.OutputDir)
                                    
                                    if meta != nil{
                                        meta_6219 := meta
                                        self.loadedMetaMap.Set(mod,meta_6219)
                                    }
                                } else { 
                                    Log_log(Log_Level__Warn, __func__, 759, Log_CreateMessage(func() string {
                                        return Lns_getVM().String_format("%s not ready meta %s, %s", []LnsAny{mod, lnsPath, metaPath})
                                    }))
                                    
                                }
                            } else { 
                                Log_log(Log_Level__Warn, __func__, 763, Log_CreateMessage(func() string {
                                    return Lns_getVM().String_format("%s not ready lua %s, %s", []LnsAny{mod, lnsPath, luaPath_6214})
                                }))
                                
                            }
                        } else {
                            Log_log(Log_Level__Warn, __func__, 767, Log_CreateMessage(func() string {
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

// 782: decl @lune.@base.@front.Front.dumpTokenize
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

// 794: decl @lune.@base.@front.Front.dumpAst
func (self *front_Front) DumpAst(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    Depend_profile(self.option.ValidProf, conv2Form4120(func() {
        var ast *TransUnit_ASTInfo
        ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1196_(scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
        ast.FP.Get_node().FP.ProcessFilter(DumpNode_createFilter(ast.FP.Get_moduleTypeInfo(), ast.FP.Get_processInfo(), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt("", 0)))
    }), scriptPath + ".profi")
}

// 811: decl @lune.@base.@front.Front.format
func (self *front_Front) Format(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1196_(scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    ast.FP.Get_node().FP.ProcessFilter(Formatter_createFilter(ast.FP.Get_moduleTypeInfo(), Lns_io_stdout), Formatter_Opt2Stem(NewFormatter_Opt(ast.FP.Get_node())))
}

// 824: decl @lune.@base.@front.Front.checkDiag
func (self *front_Front) CheckDiag(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    Util_setErrorCode(0)
    self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1196_(scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Diag, nil)
}

// 834: decl @lune.@base.@front.Front.complete
func (self *front_Front) Complete(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1196_(scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Complete, self.option.AnalyzePos)
}

// 842: decl @lune.@base.@front.Front.inquire
func (self *front_Front) Inquire(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1196_(scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Inquire, self.option.AnalyzePos)
}

// 851: decl @lune.@base.@front.Front.createGlue
func (self *front_Front) CreateGlue(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module(scriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(scriptPath), mod, front_getModuleId_1196_(scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    var filter *Nodes_Filter
    filter = GlueFilter_createFilter(self.option.OutputDir)
    ast.FP.Get_node().FP.ProcessFilter(filter, 0)
}

// 884: decl @lune.@base.@front.Front.convertLuaToStreamFromScript
func (self *front_Front) ConvertLuaToStreamFromScript(parser LnsAny,moduleId *FrontInterface_ModuleId,uptodate LnsAny,convMode LnsInt,path string,mod string,byteCompile bool,stripDebugInfo bool,openOStream front_OpenOStreamForConvert_1321_,closeOStream LnsAny) LnsAny {
    var outputDependInfo func(stream LnsAny,metaInfo LnsAny)
    outputDependInfo = func(stream LnsAny,metaInfo LnsAny) {
        if stream != nil{
            stream_6310 := stream.(Lns_oStream)
            if metaInfo != nil{
                metaInfo_6312 := metaInfo.(*front_MetaForBuildId)
                var dependInfo *OutputDepend_DependInfo
                dependInfo = NewOutputDepend_DependInfo(mod)
                for _dependMod, _ := range( metaInfo_6312.G__dependModuleMap.Items ) {
                    dependMod := _dependMod.(string)
                    dependInfo.FP.AddImpotModule(dependMod)
                }
                for _, _subMod := range( metaInfo_6312.G__subModuleMap.Items ) {
                    subMod := _subMod.(string)
                    dependInfo.FP.AddSubMod(subMod)
                }
                dependInfo.FP.Output(stream_6310)
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
    switch _exp4947 := uptodate.(type) {
    case *front_ModuleUptodate__Uptodate:
    metaInfo := _exp4947.Val1
        Util_errorLog("uptodate -- " + path)
        outputDependInfo(dependsStream, metaInfo)
    case *front_ModuleUptodate__NeedUpdate:
        if stream != nil && metaStream != nil{
            stream_6336 := stream.(Lns_oStream)
            metaStream_6337 := metaStream.(Lns_oStream)
            var ast *TransUnit_ASTInfo
            ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), front_createPaser_1099_(path, mod), mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
            retAst = ast
            
            if dependsStream != nil{
                dependsStream_6340 := dependsStream.(Lns_oStream)
                ast.FP.Get_node().FP.ProcessFilter(OutputDepend_createFilter(dependsStream_6340), 1)
            }
            var outStream Lns_oStream
            outStream = stream_6336
            var oMetaStream Lns_oStream
            oMetaStream = metaStream_6337
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
                stream_6336.Write(front_byteCompileFromLuaTxt_1120_(byteStream.FP.Get_txt(), stripDebugInfo))
                if metaStream_6337 != stream_6336{
                    metaStream_6337.Write(front_byteCompileFromLuaTxt_1120_(byteMetaStream.FP.Get_txt(), true))
                }
            }
        } else {
            Util_err("failed to open lua stream or meta stream")
        }
    case *front_ModuleUptodate__NeedTouch:
    metaCode := _exp4947.Val1
    metaInfo := _exp4947.Val2
        Util_errorLog("touch -- " + path)
        if self.option.Mode == Option_ModeKind__SaveMeta{
            if metaStream != nil{
                metaStream_6354 := metaStream.(Lns_oStream)
                metaStream_6354.Write(metaCode)
            } else {
                Util_err("failed to open meta stream")
            }
        }
        outputDependInfo(dependsStream, metaInfo)
    }
    if Str_isValidStrBuilder(){
        if streamDst != nil{
            streamDst_6365 := streamDst.(Lns_oStream)
            streamDst_6365.Write(streamMem.FP.Get_txt())
        }
        
        if metaStreamDst != nil{
            metaStreamDst_6367 := metaStreamDst.(Lns_oStream)
            metaStreamDst_6367.Write(metaStreamMem.FP.Get_txt())
        }
        
        if dependsStreamDst != nil{
            dependsStreamDst_6369 := dependsStreamDst.(Lns_oStream)
            dependsStreamDst_6369.Write(dependsStreamMem.FP.Get_txt())
        }
        
    }
    if closeOStream != nil{
        closeOStream_6371 := closeOStream.(front_CloseOStreamForConvert_1324_)
        closeOStream_6371(stream, metaStream, dependsStream)
    }
    return retAst
}

// 1013: decl @lune.@base.@front.Front.createGoOption
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
                parentPath = front_convExp5133(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(scriptPath,"/[^/]+$", "")).(string),".*/", "")))
                if len(parentPath) == 0{
                    packageName = "main"
                    
                } else if parentPath == "."{
                    packageName = "main"
                    
                } else if parentPath == ".."{
                    packageName = "main"
                    
                } else { 
                    packageName = front_convExp5190(Lns_2DDD(Lns_getVM().String_gsub(parentPath,"[^%w]", "")))
                    
                }
            }
        }
    }
    return NewConvGo_Option(packageName, self.option.AppName, self.option.MainModule)
}

// 1039: decl @lune.@base.@front.Front.convertToLua
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
    ast = self.FP.ConvertLuaToStreamFromScript(parser, FrontInterface_ModuleId__tempId, front_ModuleUptodate__NeedUpdate_Obj, convMode, scriptPath, mod, self.option.ByteCompile, self.option.StripDebugInfo, front_OpenOStreamForConvert_1321_(func(mode LnsAny)(LnsAny, LnsAny, LnsAny) {
        return Lns_io_stdout, Lns_io_stdout, self.option.FP.OpenDepend(nil)
    }), front_CloseOStreamForConvert_1324_(Front_convertToLua___anonymous_1350_))
    if ast != nil{
        ast_6409 := ast.(*TransUnit_ASTInfo)
        if _switch5405 := self.option.ConvTo; _switch5405 == Types_Lang__Go {
            var conv *Nodes_Filter
            conv = ConvGo_createFilter(self.option.Testing, "stdout", Lns_io_stdout, ast_6409, self.FP.createGoOption(scriptPath))
            ast_6409.FP.Get_node().FP.ProcessFilter(conv, ConvGo_Opt2Stem(NewConvGo_Opt(ast_6409.FP.Get_node())))
        }
    }
}

// 1078: decl @lune.@base.@front.Front.saveToGo
func (self *front_Front) SaveToGo(scriptPath string,ast *TransUnit_ASTInfo) {
    var path string
    path = front_convExp5432(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".go")))
    {
        _dir := self.option.OutputDir
        if _dir != nil {
            dir := _dir.(string)
            path = Lns_getVM().String_format("%s/%s", []LnsAny{dir, path})
            
        }
    }
    var file Lns_luaStream
    
    {
        _file := front_convExp5470(Lns_2DDD(Lns_io_open(path, "w")))
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

// 1110: decl @lune.@base.@front.Front.saveToC
func (self *front_Front) SaveToC(scriptPath string,ast *TransUnit_ASTInfo) {
    var cPath string
    cPath = front_convExp5597(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".c")))
    var file Lns_luaStream
    
    {
        _file := front_convExp5612(Lns_2DDD(Lns_io_open(cPath, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var hPath string
    hPath = front_convExp5626(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".h")))
    var hFile Lns_luaStream
    
    {
        _hFile := front_convExp5642(Lns_2DDD(Lns_io_open(hPath, "w")))
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

// 1130: decl @lune.@base.@front.Front.outputBuiltin
func (self *front_Front) OutputBuiltin(scriptPath string) {
    var mod string
    mod = Front_scriptPath2Module("lns_builtin")
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), &NewParser_DummyParser().Parser_Parser, mod, FrontInterface_ModuleId_createId(0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    self.FP.SaveToC(scriptPath, ast)
}

// 1152: decl @lune.@base.@front.Front.saveToLua
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
                oldLine_6496 := oldLine.(string)
                if len(oldBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(oldLine_6496,"^_moduleObj.__buildId", nil, nil))){
                        oldBuildIdLine = oldLine_6496
                        
                    }
                }
            }
            
            if newLine != nil{
                newLine_6500 := newLine.(string)
                if len(newBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(newLine_6500,"^_moduleObj.__buildId", nil, nil))){
                        newBuildIdLine = newLine_6500
                        
                    }
                }
            }
            
            if newLine != oldLine{
                var cont bool
                cont = false
                if newLine != nil && oldLine != nil{
                    newLine_6506 := newLine.(string)
                    oldLine_6507 := oldLine.(string)
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(oldLine_6507,"^_moduleObj.__buildId", nil, nil))){
                        if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(newLine_6506,"^_moduleObj.__buildId", nil, nil))){
                            tailBeginPos = newStream.FP.Get_lineNo()
                            
                            cont = true
                            
                        }
                    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(oldLine_6507,"^__dependModuleMap.*buildId =", nil, nil))) &&
                        Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(newLine_6506,"^__dependModuleMap.*buildId =", nil, nil))) )){
                        var oldSub string
                        oldSub = front_convExp6099(Lns_2DDD(Lns_getVM().String_gsub(oldLine_6507,"buildId =.*", "")))
                        var newSub string
                        newSub = front_convExp6112(Lns_2DDD(Lns_getVM().String_gsub(newLine_6506,"buildId =.*", "")))
                        if oldSub == newSub{
                            cont = true
                            
                        }
                    }
                }
                if Lns_op_not(cont){
                    Log_log(Log_Level__Debug, __func__, 1228, Log_CreateMessage(func() string {
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
                    oldBuildId = Front_saveToLua__txt2ModuleId_1391_(oldBuildIdLine)
                    var newBuildId *FrontInterface_ModuleId
                    newBuildId = Front_saveToLua__txt2ModuleId_1391_(newBuildIdLine)
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
    luaPath = front_convExp6349(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".lua")))
    var metaPath string
    metaPath = front_convExp6362(Lns_2DDD(Lns_getVM().String_gsub(scriptPath,"%.lns$", ".meta")))
    if Lns_isCondTrue( self.option.OutputDir){
        var filename string
        filename = front_convExp6379(Lns_2DDD(Lns_getVM().String_gsub(mod,"%.", "/")))
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
        ast = self.FP.ConvertLuaToStreamFromScript(parser, moduleId, uptodate, convMode, scriptPath, mod, self.option.ByteCompile, self.option.StripDebugInfo, front_OpenOStreamForConvert_1321_(func(mode LnsAny)(LnsAny, LnsAny, LnsAny) {
            var openLuaStream func() LnsAny
            openLuaStream = func() LnsAny {
                var fileObj Lns_luaStream
                
                {
                    _fileObj := front_convExp6537(Lns_2DDD(Lns_io_open(luaPath, "w")))
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
                        __exp := front_convExp6627(Lns_2DDD(Lns_io_open(tempMetaPath, "w+")))
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
        }), front_CloseOStreamForConvert_1324_(func(stream LnsAny,metaStream LnsAny,dependStream LnsAny) {
            if stream != nil{
                stream_6573 := stream.(Lns_oStream)
                stream_6573.Close()
            }
            if dependStream != nil{
                dependStream_6575 := dependStream.(Lns_oStream)
                dependStream_6575.Close()
            }
            if metaFileObj != nil{
                metaFileObj_6577 := metaFileObj.(Lns_luaStream)
                metaFileObj_6577.Flush()
                metaFileObj_6577.Seek("set", 0)
                var newMetaTxt string
                
                {
                    _newMetaTxt := metaFileObj_6577.Read("*a")
                    if _newMetaTxt == nil{
                        Util_err(Lns_getVM().String_format("faled to read meta. -- %s.", []LnsAny{tempMetaPath}))
                    } else {
                        newMetaTxt = _newMetaTxt.(string)
                    }
                }
                metaFileObj_6577.Close()
                var oldMetaTxt string
                oldMetaTxt = ""
                {
                    _oldFileObj := front_convExp6822(Lns_2DDD(Lns_io_open(metaPath, nil)))
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
                            _fileObj := front_convExp6903(Lns_2DDD(Lns_io_open(metaPath, "w")))
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
        ast_6594 := ast.(*TransUnit_ASTInfo)
        if _switch6984 := self.option.ConvTo; _switch6984 == Types_Lang__C {
            self.FP.SaveToC(scriptPath, ast_6594)
        } else if _switch6984 == Types_Lang__Go {
            self.FP.SaveToGo(scriptPath, ast_6594)
        }
    }
    return updateFlag
}

// 1404: decl @lune.@base.@front.Front.outputBootC
func (self *front_Front) outputBootC(scriptPath string) {
    var stream Lns_oStream
    {
        _path := self.option.BootPath
        if _path != nil {
            path := _path.(string)
            var cPath string
            cPath = front_convExp7020(Lns_2DDD(Lns_getVM().String_gsub(path,"%.lns$", ".c")))
            var file Lns_luaStream
            
            {
                _file := front_convExp7051(Lns_2DDD(Lns_io_open(cPath, "w")))
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

// 1435: decl @lune.@base.@front.Front.exec
func (self *front_Front) Exec() {
    __func__ := "@lune.@base.@front.Front.exec"
    Log_log(Log_Level__Trace, __func__, 1437, Log_CreateMessage(func() string {
        return Option_ModeKind_getTxt( self.option.Mode)
    }))
    
    if _switch7808 := self.option.Mode; _switch7808 == Option_ModeKind__Token {
        self.FP.DumpTokenize(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__Ast {
        self.FP.DumpAst(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__Format {
        self.FP.Format(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__Diag {
        self.FP.CheckDiag(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__Complete {
        self.FP.Complete(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__Inquire {
        self.FP.Inquire(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__Glue {
        self.FP.CreateGlue(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__Lua || _switch7808 == Option_ModeKind__LuaMeta {
        self.FP.convertToLua(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__Save || _switch7808 == Option_ModeKind__SaveMeta {
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
        Depend_profile(self.option.ValidProf, conv2Form7581(func() {
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
    } else if _switch7808 == Option_ModeKind__Shebang {
        {
            _modObj := front_convExp7629(Lns_2DDD(self.FP.LoadModule(Front_scriptPath2Module(self.option.ScriptPath))))
            if _modObj != nil {
                modObj := _modObj
                var code LnsInt
                code = Depend_runMain(modObj.(*Lns_luaValue).GetAt("__main"), self.option.ShebangArgList)
                Lns_getVM().OS_exit(code)
            }
        }
    } else if _switch7808 == Option_ModeKind__Exec {
        _ = front_convExp7650(Lns_2DDD(self.FP.LoadModule(Front_scriptPath2Module(self.option.ScriptPath))))
        if self.option.Testing{
            var code string
            code = "local Testing = require( \"lune.base.Testing\" )\nreturn function( path )\n  Testing.run( path );\n  Testing.outputAllResult( io.stdout );\nend\n"
            var loaded LnsAny
            var mess LnsAny
            loaded,mess = Lns_getVM().Load(code, nil)
            if loaded != nil{
                loaded_6672 := loaded.(*Lns_luaValue)
                {
                    _mod := front_convExp7703(Lns_2DDD(Lns_getVM().RunLoadedfunc(loaded_6672,Lns_2DDD([]LnsAny{}))[0]))
                    if _mod != nil {
                        mod := _mod
                        Lns_getVM().RunLoadedfunc((mod.(*Lns_luaValue)),Lns_2DDD([]LnsAny{Front_scriptPath2Module(self.option.ScriptPath)}))
                    }
                }
            } else {
                Lns_print([]LnsAny{mess})
            }
        }
    } else if _switch7808 == Option_ModeKind__BootC {
        self.FP.outputBootC(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__Builtin {
        self.FP.OutputBuiltin(self.option.ScriptPath)
    } else if _switch7808 == Option_ModeKind__MkMain {
        var mod string
        mod = Front_scriptPath2Module(self.option.ScriptPath)
        {
            _mess := ConvGo_outputGoMain(self.option.AppName, mod, self.option.Testing, self.option.OutputPath, self.option.FP.Get_runtimeOpt())
            if _mess != nil {
                mess := _mess.(string)
                Util_errorLog(mess)
            }
        }
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
    G__buildId string
    G__dependModuleMap *LnsMap
    G__subModuleMap *LnsList
    G__enableTest bool
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
    self.G__buildId = arg1
    self.G__dependModuleMap = arg2
    self.G__subModuleMap = arg3
    self.G__enableTest = arg4
}
func (self *front_MetaForBuildId) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["__buildId"] = Lns_ToCollection( self.G__buildId )
    obj.Items["__dependModuleMap"] = Lns_ToCollection( self.G__dependModuleMap )
    obj.Items["__subModuleMap"] = Lns_ToCollection( self.G__subModuleMap )
    obj.Items["__enableTest"] = Lns_ToCollection( self.G__enableTest )
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
       newObj.G__buildId = conv.(string)
    }
    if ok,conv,mess := Lns_ToLnsMapSub( objMap.Items["__dependModuleMap"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil},Lns_ToObjParam{
            front_DependMetaInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"__dependModuleMap:" + mess.(string)
    } else {
       newObj.G__dependModuleMap = conv.(*LnsMap)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["__subModuleMap"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil}}); !ok {
       return false,nil,"__subModuleMap:" + mess.(string)
    } else {
       newObj.G__subModuleMap = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["__enableTest"], false, nil); !ok {
       return false,nil,"__enableTest:" + mess.(string)
    } else {
       newObj.G__enableTest = conv.(bool)
    }
    return true, newObj, nil
}
// 292: decl @lune.@base.@front.MetaForBuildId.createModuleId
func (self *front_MetaForBuildId) CreateModuleId() *FrontInterface_ModuleId {
    return FrontInterface_ModuleId_createIdFromTxt(self.G__buildId)
}

// 297: decl @lune.@base.@front.MetaForBuildId.LoadFromMeta
func front_MetaForBuildId_LoadFromMeta_1183_(metaPath string)(LnsAny, LnsAny) {
    {
        _fileObj := front_convExp1183(Lns_2DDD(Lns_io_open(metaPath, nil)))
        if _fileObj != nil {
            fileObj := _fileObj.(Lns_luaStream)
            var luaCode LnsAny
            luaCode = fileObj.Read("*a")
            fileObj.Close()
            if luaCode != nil{
                luaCode_5939 := luaCode.(string)
                var meta LnsAny
                meta = front_convExp1173(Lns_2DDD(front_MetaForBuildId__fromStem_1179_(Lns_getVM().ExpandLuavalMap(front_loadFromLuaTxt_1117_(luaCode_5939)),nil)))
                return meta, luaCode_5939
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
    Lns_GoMod_init()
    Lns_Meta_init()
    Depend_setup(Depend_UpdateVer(front__anonymous_1007_))
    front_forceUpdateMeta = true
}
func init() {
    init_front = false
}
