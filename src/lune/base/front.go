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
type front_OpenOStreamForConvert_1316_ func (arg1 LnsAny)(LnsAny, LnsAny, LnsAny)
type front_CloseOStreamForConvert_1319_ func (arg1 LnsAny,arg2 LnsAny,arg3 LnsAny)
// for 733: ExpCast
func conv2Form3957( src func ()) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        src()
        return []LnsAny{}
    }
}
// for 1116: ExpCast
func conv2Form6347( src func ()) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        src()
        return []LnsAny{}
    }
}
// for 185
func front_convExp743(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 254
func front_convExp1135(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 489
func front_convExp2295(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 497
func front_convExp2324(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1168
func front_convExp6033(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1214
func front_convExp6225(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1228
func front_convExp6306(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 272
func front_convExp1208(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 286
func front_convExp1299(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 287
func front_convExp1319(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 317
func front_convExp1409(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 482
func front_convExp2232(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 258
func front_convExp1125(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 384
func front_convExp1678(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 585
func front_convExp2869(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 620
func front_convExp3170(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 629
func front_convExp3228(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 675
func front_convExp3468(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 684
func front_convExp3526(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 954
func front_convExp4932(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 959
func front_convExp4970(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 975
func front_convExp5039(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 977
func front_convExp5054(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 981
func front_convExp5070(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 982
func front_convExp5086(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1013
func front_convExp5240(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1072
func front_convExp5490(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1073
func front_convExp5503(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1117
func front_convExp5753(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1118
func front_convExp5768(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1120
func front_convExp5785(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1153
func front_convExp5943(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1265
func front_convExp6432(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1266
func front_convExp6463(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1327
func front_convExp6743(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
func front__anonymous_1007_(ver LnsInt) {
    LuaVer_setCurVer(ver)
}
// 136: decl @lune.@base.@front.createPaser
func front_createPaser_1098_(path string,mod string) *Parser_Parser {
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

// 145: decl @lune.@base.@front.scriptPath2Module
func Front_scriptPath2Module(path string) string {
    return Util_scriptPath2Module(path)
}

// 180: decl @lune.@base.@front.loadFromChunk
func front_loadFromChunk_1113_(chunk LnsAny,err LnsAny) LnsAny {
    if err != nil{
        err_5602 := err.(string)
        Util_errorLog(err_5602)
    }
    if chunk != nil{
        chunk_5604 := chunk.(*Lns_luaValue)
        {
            _work := front_convExp743(Lns_2DDD(Lns_getVM().RunLoadedfunc(chunk_5604,Lns_2DDD([]LnsAny{}))[0]))
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

// 193: decl @lune.@base.@front.loadFromLuaTxt
func front_loadFromLuaTxt_1116_(txt string) LnsAny {
    return front_loadFromChunk_1113_(Lns_getVM().Load(txt, nil))
}

// 198: decl @lune.@base.@front.byteCompileFromLuaTxt
func front_byteCompileFromLuaTxt_1119_(txt string,stripDebugInfo bool) string {
    var chunk LnsAny
    var err LnsAny
    chunk,err = Lns_getVM().Load(txt, nil)
    if chunk != nil{
        chunk_5619 := chunk.(*Lns_luaValue)
        return Lns_getVM().String_dump(chunk_5619, stripDebugInfo)
    }
    panic(Lns_unwrapDefault( err, "load error").(string))
// insert a dummy
    return ""
}

// 267: decl @lune.@base.@front.getMetaInfo
func front_getMetaInfo_1185_(lnsPath string,mod string,outdir LnsAny)(LnsAny, string, string) {
    var moduleMetaPath string
    moduleMetaPath = lnsPath
    if outdir != nil{
        outdir_5680 := outdir.(string)
        moduleMetaPath = Lns_getVM().String_format("%s/%s", []LnsAny{outdir_5680, Lns_car(Lns_getVM().String_gsub(mod,"%.", "/")).(string)})
        
    }
    moduleMetaPath = front_convExp1208(Lns_2DDD(Lns_getVM().String_gsub(moduleMetaPath,"%.lns$", ".meta")))
    
    {
        _meta, _metaCode := front_MetaForBuildId_LoadFromMeta_1182_(moduleMetaPath)
        if _meta != nil && _metaCode != nil {
            meta := _meta.(*front_MetaForBuildId)
            metaCode := _metaCode.(string)
            return meta, moduleMetaPath, metaCode
        }
    }
    return nil, moduleMetaPath, ""
}

// 309: decl @lune.@base.@front.getModuleId
func front_getModuleId_1191_(lnsPath string,mod string,outdir LnsAny,metaInfo LnsAny) *FrontInterface_ModuleId {
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
        metaInfo = front_convExp1409(Lns_2DDD(front_getMetaInfo_1185_(lnsPath, mod, outdir)))
        
    }
    if metaInfo != nil{
        metaInfo_5710 := metaInfo.(*front_MetaForBuildId)
        var buildId *FrontInterface_ModuleId
        buildId = metaInfo_5710.FP.CreateModuleId()
        buildCount = buildId.FP.Get_buildCount()
        
    }
    return FrontInterface_ModuleId_createId(fileTime, buildCount)
}

func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1205_() string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1208_() string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1211_() string {
    return "NeedUpdate"
}


func Front_getModuleIdAndCheckUptodate___anonymous_1217_() string {
    return "not found meta"
}















func Front_convertToLua___anonymous_1334_(stream LnsAny,metaStream LnsAny,dependStream LnsAny) {
    if dependStream != nil{
        dependStream_6083 := dependStream.(Lns_oStream)
        dependStream_6083.Close()
    }
}
// 1012: decl @lune.@base.@front.Front.saveToLua.txt2ModuleId
func Front_saveToLua__txt2ModuleId_1349_(txt string) *FrontInterface_ModuleId {
    var buildIdTxt string
    buildIdTxt = front_convExp5240(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(txt,"^_moduleObj.__buildId = ", "")).(string),"\"", "")))
    return FrontInterface_ModuleId_createIdFromTxt(buildIdTxt)
}








// 1279: decl @lune.@base.@front.convertLnsCode2LuaCode
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


// 1348: decl @lune.@base.@front.exec
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

// 1365: decl @lune.@base.@front.setFront
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
    CheckDiag()
    checkUptodateMeta(arg1 string, arg2 LnsAny) LnsAny
    Complete()
    convert(arg1 *TransUnit_ASTInfo, arg2 string, arg3 Lns_oStream, arg4 Lns_oStream, arg5 LnsInt, arg6 bool)
    convertFromAst(arg1 *TransUnit_ASTInfo, arg2 string, arg3 LnsInt)(string, string)
    ConvertLns2LuaCode(arg1 *FrontInterface_ImportModuleInfo, arg2 Lns_iStream, arg3 string) string
    ConvertLuaToStreamFromScript(arg1 bool, arg2 LnsInt, arg3 string, arg4 string, arg5 bool, arg6 bool, arg7 front_OpenOStreamForConvert_1316_, arg8 LnsAny) LnsAny
    convertToLua()
    createAst(arg1 *FrontInterface_ImportModuleInfo, arg2 *Parser_Parser, arg3 string, arg4 *FrontInterface_ModuleId, arg5 LnsAny, arg6 LnsInt, arg7 LnsAny) *TransUnit_ASTInfo
    CreateGlue()
    createPaser() *Parser_Parser
    DumpAst()
    DumpTokenize()
    Error(arg1 string)
    Exec()
    Format()
    getLoadInfo(arg1 string) LnsAny
    getModuleIdAndCheckUptodate(arg1 string, arg2 string)(*FrontInterface_ModuleId, LnsAny)
    Inquire()
    loadFile(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string)(LnsAny, LnsAny)
    loadFileToLuaCode(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string)(LnsAny, string)
    LoadFromLnsTxt(arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string) LnsAny
    loadLua(arg1 string) LnsAny
    LoadMeta(arg1 *FrontInterface_ImportModuleInfo, arg2 string) LnsAny
    LoadModule(arg1 string)(LnsAny, LnsAny)
    outputBootC()
    OutputBuiltin()
    SaveToC(arg1 *TransUnit_ASTInfo)
    SaveToGo(arg1 *TransUnit_ASTInfo)
    SaveToLua() bool
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
// 77: DeclConstr
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
        _exp211 := Depend_getLoadedMod()
        _key211, _val211 := _exp211.Get1stFromMap()
        for _key211 != nil {
            mod := _key211.(string)
            modval := _val211
            if mod == "lune.base.Testing"{
                loadedMap.Set(mod,modval)
            }
            _key211, _val211 = _exp211.NextFromMap( _key211 )
        }
    }
    self.preloadedModMap = loadedMap
    
}

// 97: decl @lune.@base.@front.Front.getLoadInfo
func (self *front_Front) getLoadInfo(mod string) LnsAny {
    if self.option.Testing{
        return self.loadedMapTest.Items[mod]
    }
    return self.loadedMap.Items[mod]
}

// 104: decl @lune.@base.@front.Front.setLoadInfo
func (self *front_Front) setLoadInfo(mod string,info *front_LoadInfo) {
    if self.option.Testing{
        self.loadedMapTest.Set(mod,info)
    }
    self.loadedMap.Set(mod,info)
}






// 121: decl @lune.@base.@front.Front.error
func (self *front_Front) Error(message string) {
    Util_errorLog(message)
    Util_printStackTrace()
    Lns_getVM().OS_exit(1)
}

// 127: decl @lune.@base.@front.Front.loadLua
func (self *front_Front) loadLua(path string) LnsAny {
    var chunk LnsAny
    var err LnsAny
    chunk,err = Lns_getVM().Loadfile(path)
    if chunk != nil{
        chunk_5555 := chunk.(*Lns_luaValue)
        return Lns_unwrap( Lns_car(Lns_getVM().RunLoadedfunc(chunk_5555,Lns_2DDD([]LnsAny{}))[0]))
    }
    Util_errorLog(Lns_unwrapDefault( err, Lns_getVM().String_format("load error -- %s.", []LnsAny{path})).(string))
    return nil
}

// 149: decl @lune.@base.@front.Front.createPaser
func (self *front_Front) createPaser() *Parser_Parser {
    var mod string
    mod = Front_scriptPath2Module(self.option.ScriptPath)
    return front_createPaser_1098_(self.option.ScriptPath, mod)
}

// 156: decl @lune.@base.@front.Front.createAst
func (self *front_Front) createAst(importModuleInfo *FrontInterface_ImportModuleInfo,parser *Parser_Parser,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *TransUnit_ASTInfo {
    var transUnit *TransUnit_TransUnit
    transUnit = NewTransUnit_TransUnit(moduleId, importModuleInfo, &NewConvLua_MacroEvalImp().Nodes_MacroEval, analyzeModule, analyzeMode, pos, self.option.TargetLuaVer, self.option.TransCtrlInfo)
    return transUnit.FP.CreateAST(parser, false, mod)
}

// 168: decl @lune.@base.@front.Front.convert
func (self *front_Front) convert(ast *TransUnit_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool) {
    var conv *Nodes_Filter
    conv = ConvLua_createFilter(streamName, stream, metaStream, convMode, inMacro, ast.FP.Get_moduleTypeInfo(), ast.FP.Get_processInfo(), ast.FP.Get_moduleSymbolKind(), self.option.UseLuneModule, self.option.TargetLuaVer, self.option.Testing, self.option.UseIpairs)
    ast.FP.Get_node().FP.ProcessFilter(conv, ConvLua_Opt2Stem(NewConvLua_Opt(ast.FP.Get_node())))
}

// 207: decl @lune.@base.@front.Front.convertFromAst
func (self *front_Front) convertFromAst(ast *TransUnit_ASTInfo,streamName string,mode LnsInt)(string, string) {
    var stream *Util_memStream
    stream = NewUtil_memStream()
    var metaStream *Util_memStream
    metaStream = NewUtil_memStream()
    self.FP.convert(ast, streamName, stream.FP, metaStream.FP, mode, false)
    return metaStream.FP.Get_txt(), stream.FP.Get_txt()
}

// 221: decl @lune.@base.@front.Front.loadFromLnsTxt
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
    return Lns_unwrap( front_loadFromLuaTxt_1116_(luaTxt))
}

// 279: decl @lune.@base.@front.Front.searchModuleFile
func (self *front_Front) searchModuleFile(mod string,suffix string,addPath LnsAny) LnsAny {
    var lnsSearchPath string
    lnsSearchPath = Lns_package_path
    if addPath != nil{
        addPath_5693 := addPath.(string)
        lnsSearchPath = Lns_getVM().String_format("%s/?%s;%s", []LnsAny{addPath_5693, suffix, Lns_package_path})
        
    }
    lnsSearchPath = front_convExp1299(Lns_2DDD(Lns_getVM().String_gsub(lnsSearchPath,"%.lua$", suffix)))
    
    lnsSearchPath = front_convExp1319(Lns_2DDD(Lns_getVM().String_gsub(lnsSearchPath,"%.lua;", suffix + ";")))
    
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

// 349: decl @lune.@base.@front.Front.getModuleIdAndCheckUptodate
func (self *front_Front) getModuleIdAndCheckUptodate(lnsPath string,mod string)(*FrontInterface_ModuleId, LnsAny) {
    __func__ := "@lune.@base.@front.Front.getModuleIdAndCheckUptodate"
    var uptodate LnsAny
    uptodate = front_ModuleUptodate__NeedUpdate_Obj
    if self.option.TransCtrlInfo.UptodateMode == Option_CheckingUptodateMode__Force{
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
                    Log_log(Log_Level__Debug, __func__, 372, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1205_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    modMetaPath = _modMetaPath.(string)
                }
            }
            var time LnsReal
            
            {
                _time := Depend_getFileLastModifiedTime(modMetaPath)
                if _time == nil{
                    Log_log(Log_Level__Debug, __func__, 377, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1208_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    time = _time.(LnsReal)
                }
            }
            if time > metaTime{
                var dependMeta *front_MetaForBuildId
                
                {
                    _dependMeta := front_convExp1678(Lns_2DDD(front_MetaForBuildId_LoadFromMeta_1182_(modMetaPath)))
                    if _dependMeta == nil{
                        Log_log(Log_Level__Debug, __func__, 385, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1211_))
                        
                        return front_ModuleUptodate__NeedUpdate_Obj
                    } else {
                        dependMeta = _dependMeta.(*front_MetaForBuildId)
                    }
                }
                var orgMetaModuleId *FrontInterface_ModuleId
                orgMetaModuleId = FrontInterface_ModuleId_createIdFromTxt(dependItem.BuildId)
                var metaModuleId *FrontInterface_ModuleId
                metaModuleId = dependMeta.FP.CreateModuleId()
                if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                    Lns_setStackVal( metaModuleId.FP.Get_buildCount() != 0) &&
                    Lns_setStackVal( metaModuleId.FP.Get_buildCount() != orgMetaModuleId.FP.Get_buildCount()) ).(bool)){
                    Log_log(Log_Level__Debug, __func__, 395, Log_CreateMessage(func() string {
                        return Lns_getVM().String_format("NeedUpdate: %s, %d, %d", []LnsAny{modMetaPath, metaModuleId.FP.Get_buildCount(), orgMetaModuleId.FP.Get_buildCount()})
                    }))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                }
            }
        }
        if self.option.TransCtrlInfo.UptodateMode == Option_CheckingUptodateMode__Touch{
            return &front_ModuleUptodate__NeedTouch{metaCode, metaInfo}
        }
        return &front_ModuleUptodate__Uptodate{metaInfo}
    }
    var metaInfo LnsAny
    var metaPath string
    var metaCode string
    metaInfo,metaPath,metaCode = front_getMetaInfo_1185_(lnsPath, mod, self.option.OutputDir)
    if metaInfo != nil{
        metaInfo_5760 := metaInfo.(*front_MetaForBuildId)
        if metaInfo_5760.__enableTest == self.option.Testing{
            var buildId *FrontInterface_ModuleId
            buildId = FrontInterface_ModuleId_createIdFromTxt(metaInfo_5760.__buildId)
            if buildId != FrontInterface_ModuleId__tempId{
                var lnsTime LnsAny
                lnsTime = Depend_getFileLastModifiedTime(lnsPath)
                var metaTime LnsAny
                metaTime = Depend_getFileLastModifiedTime(metaPath)
                if lnsTime != nil && metaTime != nil{
                    lnsTime_5767 := lnsTime.(LnsReal)
                    metaTime_5768 := metaTime.(LnsReal)
                    if lnsTime_5767 == buildId.FP.Get_modTime(){
                        uptodate = checkDependUptodate(metaTime_5768, metaInfo_5760, metaCode)
                        
                    }
                }
            }
        } else { 
        }
    } else {
        Log_log(Log_Level__Debug, __func__, 433, Log_CreateMessage(Front_getModuleIdAndCheckUptodate___anonymous_1217_))
        
    }
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_1191_(lnsPath, mod, self.option.OutputDir, metaInfo)
    if moduleId == FrontInterface_ModuleId__tempId{
        Util_err(Lns_getVM().String_format("not found -- %s", []LnsAny{lnsPath}))
    }
    return moduleId, uptodate
}

// 446: decl @lune.@base.@front.Front.convertLns2LuaCode
func (self *front_Front) ConvertLns2LuaCode(importModuleInfo *FrontInterface_ImportModuleInfo,stream Lns_iStream,streamName string) string {
    var mod string
    mod = Front_scriptPath2Module(streamName)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(importModuleInfo, &NewParser_StreamParser(stream, streamName, false).Parser_Parser, mod, FrontInterface_ModuleId_createId(0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(ast, streamName, ConvLua_ConvMode__Exec)
    return luaTxt
}

// 465: decl @lune.@base.@front.Front.loadFileToLuaCode
func (self *front_Front) loadFileToLuaCode(importModuleInfo *FrontInterface_ImportModuleInfo,path string,mod string)(LnsAny, string) {
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(importModuleInfo, front_createPaser_1098_(path, mod), mod, front_getModuleId_1191_(path, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
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
                    newpath = front_convExp2232(Lns_2DDD(Lns_getVM().String_gsub(path,".lns$", suffix)))
                    
                }
            }
            var saveTxt string
            saveTxt = txt
            if byteCompile{
                saveTxt = front_byteCompileFromLuaTxt_1119_(saveTxt, stripDebugInfo)
                
            }
            if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                Lns_setStackVal( Lns_op_not(front_forceUpdateMeta)) &&
                Lns_setStackVal( checkUpdate) ).(bool)){
                {
                    _fileObj := front_convExp2295(Lns_2DDD(Lns_io_open(newpath, nil)))
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
                _fileObj := front_convExp2324(Lns_2DDD(Lns_io_open(newpath, "w")))
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
    meta = Lns_unwrap( front_loadFromLuaTxt_1116_(metaTxt))
    return meta, luaTxt
}

// 517: decl @lune.@base.@front.Front.loadFile
func (self *front_Front) loadFile(importModuleInfo *FrontInterface_ImportModuleInfo,path string,mod string)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@front.Front.loadFile"
    Log_log(Log_Level__Info, __func__, 521, Log_CreateMessage(func() string {
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
    return meta, Lns_unwrap( front_loadFromLuaTxt_1116_(luaTxt))
}

// 533: decl @lune.@base.@front.Front.searchModule
func (self *front_Front) SearchModule(mod string) LnsAny {
    return self.FP.searchModuleFile(mod, ".lns", nil)
}

// 537: decl @lune.@base.@front.Front.searchLuaFile
func (self *front_Front) searchLuaFile(moduleFullName string,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(moduleFullName, ".lua", addSearchPath)
}

// 557: decl @lune.@base.@front.Front.checkUptodateMeta
func (self *front_Front) checkUptodateMeta(metaPath string,addSearchPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.checkUptodateMeta"
    var metaObj LnsAny
    
    {
        _metaObj := self.FP.loadLua(metaPath)
        if _metaObj == nil{
            Log_log(Log_Level__Warn, __func__, 559, Log_CreateMessage(func() string {
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
        Log_log(Log_Level__Warn, __func__, 564, Log_CreateMessage(func() string {
            return Lns_getVM().String_format("unmatch meta version -- %s", []LnsAny{metaPath})
        }))
        
        return nil
    }
    if meta.GetAt( "__hasTest" ).(bool){
        if meta.GetAt( "__enableTest" ).(bool) != self.option.Testing{
            Log_log(Log_Level__Warn, __func__, 570, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("unmatch test setting -- %s", []LnsAny{metaPath})
            }))
            
            return nil
        }
    }
    {
        _exp3058 := meta.GetAt( "__dependModuleMap" ).(*Lns_luaValue)
        _key3058, _ := _exp3058.Get1stFromMap()
        for _key3058 != nil {
            moduleFullName := _key3058.(string)
            {
                _lnsPath := self.FP.SearchModule(moduleFullName)
                if _lnsPath != nil {
                    lnsPath := _lnsPath.(string)
                    {
                        _moduleLuaPath := self.FP.searchLuaFile(moduleFullName, addSearchPath)
                        if _moduleLuaPath != nil {
                            moduleLuaPath := _moduleLuaPath.(string)
                            if Lns_op_not(Util_getReadyCode(lnsPath, metaPath)){
                                Log_log(Log_Level__Warn, __func__, 581, Log_CreateMessage(func() string {
                                    return Lns_getVM().String_format("not ready -- %s, %s", []LnsAny{lnsPath, metaPath})
                                }))
                                
                                return nil
                            }
                            var moduleMetaPath string
                            moduleMetaPath = front_convExp2869(Lns_2DDD(Lns_getVM().String_gsub(moduleLuaPath,"%.lua$", ".meta")))
                            if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                                Lns_setStackVal( Depend_existFile(moduleMetaPath)) &&
                                Lns_setStackVal( Lns_op_not(Util_getReadyCode(moduleMetaPath, metaPath))) ).(bool)){
                                Log_log(Log_Level__Warn, __func__, 589, Log_CreateMessage(func() string {
                                    return Lns_getVM().String_format("not ready -- %s, %s", []LnsAny{moduleMetaPath, metaPath})
                                }))
                                
                                return nil
                            }
                        } else {
                            Log_log(Log_Level__Warn, __func__, 594, Log_CreateMessage(func() string {
                                return Lns_getVM().String_format("not found .lua file for -- %s", []LnsAny{moduleFullName})
                            }))
                            
                            return nil
                        }
                    }
                } else {
                    Log_log(Log_Level__Warn, __func__, 599, Log_CreateMessage(func() string {
                        return Lns_getVM().String_format("not found .lns file -- %s", []LnsAny{moduleFullName})
                    }))
                    
                    return nil
                }
            }
            _key3058, _ = _exp3058.NextFromMap( _key3058 )
        }
    }
    return meta
}

// 607: decl @lune.@base.@front.Front.loadModule
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
                        self.FP.setLoadInfo(mod, Newfront_LoadInfo(Lns_unwrap( front_loadFromLuaTxt_1116_(luaTxt)), meta))
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
                        luaPath = front_convExp3170(Lns_2DDD(Lns_getVM().String_gsub(lnsPath, "%.lns$", ".lua")))
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
                            luaPath_5902 := luaPath.(string)
                            if Util_getReadyCode(lnsPath, luaPath_5902){
                                var metaPath string
                                metaPath = front_convExp3228(Lns_2DDD(Lns_getVM().String_gsub(luaPath_5902, "%.lua$", ".meta")))
                                if Util_getReadyCode(lnsPath, metaPath){
                                    {
                                        _preLoadInfo := self.preloadedModMap.Items[mod]
                                        if _preLoadInfo != nil {
                                            preLoadInfo := _preLoadInfo
                                            loadVal = preLoadInfo
                                            
                                        } else {
                                            loadVal = self.FP.loadLua(luaPath_5902)
                                            
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

// 666: decl @lune.@base.@front.Front.loadMeta
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
                        luaPath = front_convExp3468(Lns_2DDD(Lns_getVM().String_gsub(lnsPath, "%.lns$", ".lua")))
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
                            luaPath_5936 := luaPath.(string)
                            if Util_getReadyCode(lnsPath, luaPath_5936){
                                var metaPath string
                                metaPath = front_convExp3526(Lns_2DDD(Lns_getVM().String_gsub(luaPath_5936, "%.lua$", ".meta")))
                                if Util_getReadyCode(lnsPath, metaPath){
                                    meta = self.FP.checkUptodateMeta(metaPath, self.option.OutputDir)
                                    
                                    if meta != nil{
                                        meta_5941 := meta
                                        self.loadedMetaMap.Set(mod,meta_5941)
                                    }
                                } else { 
                                    Log_log(Log_Level__Warn, __func__, 692, Log_CreateMessage(func() string {
                                        return Lns_getVM().String_format("%s not ready meta %s, %s", []LnsAny{mod, lnsPath, metaPath})
                                    }))
                                    
                                }
                            } else { 
                                Log_log(Log_Level__Warn, __func__, 696, Log_CreateMessage(func() string {
                                    return Lns_getVM().String_format("%s not ready lua %s, %s", []LnsAny{mod, lnsPath, luaPath_5936})
                                }))
                                
                            }
                        } else {
                            Log_log(Log_Level__Warn, __func__, 700, Log_CreateMessage(func() string {
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

// 715: decl @lune.@base.@front.Front.dumpTokenize
func (self *front_Front) DumpTokenize() {
    var parser *Parser_Parser
    parser = self.FP.createPaser()
    for  {
        var token *Parser_Token
        
        {
            _token := parser.FP.GetToken()
            if _token == nil{
                break
            } else {
                token = _token.(*Parser_Token)
            }
        }
        Lns_print([]LnsAny{Parser_TokenKind_getTxt( token.Kind), token.Pos.LineNo, token.Pos.Column, token.Txt})
    }
}

// 727: decl @lune.@base.@front.Front.dumpAst
func (self *front_Front) DumpAst() {
    var mod string
    mod = Front_scriptPath2Module(self.option.ScriptPath)
    Depend_profile(self.option.ValidProf, conv2Form3957(func() {
        var ast *TransUnit_ASTInfo
        ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(), mod, front_getModuleId_1191_(self.option.ScriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
        ast.FP.Get_node().FP.ProcessFilter(DumpNode_createFilter(ast.FP.Get_moduleTypeInfo(), ast.FP.Get_processInfo(), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt("", 0)))
    }), self.option.ScriptPath + ".profi")
}

// 744: decl @lune.@base.@front.Front.format
func (self *front_Front) Format() {
    var mod string
    mod = Front_scriptPath2Module(self.option.ScriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(), mod, front_getModuleId_1191_(self.option.ScriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    ast.FP.Get_node().FP.ProcessFilter(Formatter_createFilter(ast.FP.Get_moduleTypeInfo(), Lns_io_stdout), Formatter_Opt2Stem(NewFormatter_Opt(ast.FP.Get_node())))
}

// 757: decl @lune.@base.@front.Front.checkDiag
func (self *front_Front) CheckDiag() {
    var mod string
    mod = Front_scriptPath2Module(self.option.ScriptPath)
    Util_setErrorCode(0)
    self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(), mod, front_getModuleId_1191_(self.option.ScriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Diag, nil)
}

// 766: decl @lune.@base.@front.Front.complete
func (self *front_Front) Complete() {
    var mod string
    mod = Front_scriptPath2Module(self.option.ScriptPath)
    self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(), mod, front_getModuleId_1191_(self.option.ScriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Complete, self.option.AnalyzePos)
}

// 773: decl @lune.@base.@front.Front.inquire
func (self *front_Front) Inquire() {
    var mod string
    mod = Front_scriptPath2Module(self.option.ScriptPath)
    self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(), mod, front_getModuleId_1191_(self.option.ScriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Inquire, self.option.AnalyzePos)
}

// 781: decl @lune.@base.@front.Front.createGlue
func (self *front_Front) CreateGlue() {
    var mod string
    mod = Front_scriptPath2Module(self.option.ScriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), self.FP.createPaser(), mod, front_getModuleId_1191_(self.option.ScriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    var filter *Nodes_Filter
    filter = GlueFilter_createFilter(self.option.OutputDir)
    ast.FP.Get_node().FP.ProcessFilter(filter, 0)
}

// 815: decl @lune.@base.@front.Front.convertLuaToStreamFromScript
func (self *front_Front) ConvertLuaToStreamFromScript(checkUptodate bool,convMode LnsInt,path string,mod string,byteCompile bool,stripDebugInfo bool,openOStream front_OpenOStreamForConvert_1316_,closeOStream LnsAny) LnsAny {
    var outputDependInfo func(stream LnsAny,metaInfo LnsAny)
    outputDependInfo = func(stream LnsAny,metaInfo LnsAny) {
        if stream != nil{
            stream_6023 := stream.(Lns_oStream)
            if metaInfo != nil{
                metaInfo_6025 := metaInfo.(*front_MetaForBuildId)
                var dependInfo *OutputDepend_DependInfo
                dependInfo = NewOutputDepend_DependInfo(mod)
                for _dependMod, _ := range( metaInfo_6025.__dependModuleMap.Items ) {
                    dependMod := _dependMod.(string)
                    dependInfo.FP.AddImpotModule(dependMod)
                }
                for _, _subMod := range( metaInfo_6025.__subModuleMap.Items ) {
                    subMod := _subMod.(string)
                    dependInfo.FP.AddSubMod(subMod)
                }
                dependInfo.FP.Output(stream_6023)
            } else {
                Util_err("metaInfo is nil")
            }
        }
    }
    var retAst LnsAny
    retAst = nil
    var moduleId *FrontInterface_ModuleId
    var uptodate LnsAny
    if checkUptodate{
        moduleId, uptodate = self.FP.getModuleIdAndCheckUptodate(path, mod)
        
    } else { 
        moduleId, uptodate = FrontInterface_ModuleId__tempId, front_ModuleUptodate__NeedUpdate_Obj
        
    }
    var stream LnsAny
    var metaStream LnsAny
    var dependsStream LnsAny
    stream,metaStream,dependsStream = openOStream(uptodate)
    switch _exp4719 := uptodate.(type) {
    case *front_ModuleUptodate__Uptodate:
    metaInfo := _exp4719.Val1
        Util_errorLog("uptodate -- " + path)
        outputDependInfo(dependsStream, metaInfo)
    case *front_ModuleUptodate__NeedUpdate:
        if stream != nil && metaStream != nil{
            stream_6045 := stream.(Lns_oStream)
            metaStream_6046 := metaStream.(Lns_oStream)
            var ast *TransUnit_ASTInfo
            ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), front_createPaser_1098_(path, mod), mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
            retAst = ast
            
            if dependsStream != nil{
                dependsStream_6049 := dependsStream.(Lns_oStream)
                ast.FP.Get_node().FP.ProcessFilter(OutputDepend_createFilter(dependsStream_6049), 1)
            }
            var outStream Lns_oStream
            outStream = stream_6045
            var oMetaStream Lns_oStream
            oMetaStream = metaStream_6046
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
                stream_6045.Write(front_byteCompileFromLuaTxt_1119_(byteStream.FP.Get_txt(), stripDebugInfo))
                if metaStream_6046 != stream_6045{
                    metaStream_6046.Write(front_byteCompileFromLuaTxt_1119_(byteMetaStream.FP.Get_txt(), true))
                }
            }
        } else {
            Util_err("failed to open lua stream or meta stream")
        }
    case *front_ModuleUptodate__NeedTouch:
    metaCode := _exp4719.Val1
    metaInfo := _exp4719.Val2
        Util_errorLog("touch -- " + path)
        if self.option.Mode == Option_ModeKind__SaveMeta{
            if metaStream != nil{
                metaStream_6063 := metaStream.(Lns_oStream)
                metaStream_6063.Write(metaCode)
            } else {
                Util_err("failed to open meta stream")
            }
        }
        outputDependInfo(dependsStream, metaInfo)
    }
    if closeOStream != nil{
        closeOStream_6066 := closeOStream.(front_CloseOStreamForConvert_1319_)
        closeOStream_6066(stream, metaStream, dependsStream)
    }
    return retAst
}

// 917: decl @lune.@base.@front.Front.convertToLua
func (self *front_Front) convertToLua() {
    var mod string
    mod = Front_scriptPath2Module(self.option.ScriptPath)
    var convMode LnsInt
    convMode = ConvLua_ConvMode__Convert
    if self.option.Mode == Option_ModeKind__LuaMeta{
        convMode = ConvLua_ConvMode__ConvMeta
        
    }
    var ast LnsAny
    ast = self.FP.ConvertLuaToStreamFromScript(false, convMode, self.option.ScriptPath, mod, self.option.ByteCompile, self.option.StripDebugInfo, front_OpenOStreamForConvert_1316_(func(mode LnsAny)(LnsAny, LnsAny, LnsAny) {
        return Lns_io_stdout, Lns_io_stdout, self.option.FP.OpenDepend()
    }), front_CloseOStreamForConvert_1319_(Front_convertToLua___anonymous_1334_))
    if ast != nil{
        ast_6086 := ast.(*TransUnit_ASTInfo)
        if _switch4906 := self.option.ConvTo; _switch4906 == Option_Conv__Go {
            var conv *Nodes_Filter
            conv = ConvGo_createFilter(self.option.Testing, "stdout", Lns_io_stdout, ast_6086)
            ast_6086.FP.Get_node().FP.ProcessFilter(conv, ConvGo_Opt2Stem(NewConvGo_Opt(ast_6086.FP.Get_node())))
        }
    }
}

// 953: decl @lune.@base.@front.Front.saveToGo
func (self *front_Front) SaveToGo(ast *TransUnit_ASTInfo) {
    var path string
    path = front_convExp4932(Lns_2DDD(Lns_getVM().String_gsub(self.option.ScriptPath,"%.lns$", ".go")))
    {
        _dir := self.option.OutputDir
        if _dir != nil {
            dir := _dir.(string)
            path = Lns_getVM().String_format("%s/%s", []LnsAny{dir, path})
            
        }
    }
    var file Lns_luaStream
    
    {
        _file := front_convExp4970(Lns_2DDD(Lns_io_open(path, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var conv *Nodes_Filter
    conv = ConvGo_createFilter(self.option.Testing, path, file, ast)
    ast.FP.Get_node().FP.ProcessFilter(conv, ConvGo_Opt2Stem(NewConvGo_Opt(ast.FP.Get_node())))
    file.Close()
}

// 974: decl @lune.@base.@front.Front.saveToC
func (self *front_Front) SaveToC(ast *TransUnit_ASTInfo) {
    var cPath string
    cPath = front_convExp5039(Lns_2DDD(Lns_getVM().String_gsub(self.option.ScriptPath,"%.lns$", ".c")))
    var file Lns_luaStream
    
    {
        _file := front_convExp5054(Lns_2DDD(Lns_io_open(cPath, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var hPath string
    hPath = front_convExp5070(Lns_2DDD(Lns_getVM().String_gsub(self.option.ScriptPath,"%.lns$", ".h")))
    var hFile Lns_luaStream
    
    {
        _hFile := front_convExp5086(Lns_2DDD(Lns_io_open(hPath, "w")))
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

// 994: decl @lune.@base.@front.Front.outputBuiltin
func (self *front_Front) OutputBuiltin() {
    var mod string
    mod = Front_scriptPath2Module("lns_builtin")
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(NewFrontInterface_ImportModuleInfo(), &NewParser_DummyParser().Parser_Parser, mod, FrontInterface_ModuleId_createId(0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    self.FP.SaveToC(ast)
}

// 1009: decl @lune.@base.@front.Front.saveToLua
func (self *front_Front) SaveToLua() bool {
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
                oldLine_6147 := oldLine.(string)
                if len(oldBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(oldLine_6147,"^_moduleObj.__buildId", nil, nil))){
                        oldBuildIdLine = oldLine_6147
                        
                    }
                }
            }
            
            if newLine != nil{
                newLine_6151 := newLine.(string)
                if len(newBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(newLine_6151,"^_moduleObj.__buildId", nil, nil))){
                        newBuildIdLine = newLine_6151
                        
                    }
                }
            }
            
            if newLine != oldLine{
                var cont bool
                cont = false
                if newLine != nil && oldLine != nil{
                    newLine_6157 := newLine.(string)
                    oldLine_6158 := oldLine.(string)
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(oldLine_6158,"^_moduleObj.__buildId", nil, nil))){
                        if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(newLine_6157,"^_moduleObj.__buildId", nil, nil))){
                            tailBeginPos = newStream.FP.Get_lineNo()
                            
                            cont = true
                            
                        }
                    } else if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
                        Lns_setStackVal( Lns_car(Lns_getVM().String_find(oldLine_6158,"^__dependModuleMap.*buildId =", nil, nil))) &&
                        Lns_setStackVal( Lns_car(Lns_getVM().String_find(newLine_6157,"^__dependModuleMap.*buildId =", nil, nil))) )){
                        var oldSub string
                        oldSub = front_convExp5490(Lns_2DDD(Lns_getVM().String_gsub(oldLine_6158,"buildId =.*", "")))
                        var newSub string
                        newSub = front_convExp5503(Lns_2DDD(Lns_getVM().String_gsub(newLine_6157,"buildId =.*", "")))
                        if oldSub == newSub{
                            cont = true
                            
                        }
                    }
                }
                if Lns_op_not(cont){
                    Log_log(Log_Level__Debug, __func__, 1080, Log_CreateMessage(func() string {
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
                    oldBuildId = Front_saveToLua__txt2ModuleId_1349_(oldBuildIdLine)
                    var newBuildId *FrontInterface_ModuleId
                    newBuildId = Front_saveToLua__txt2ModuleId_1349_(newBuildIdLine)
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
    mod = Front_scriptPath2Module(self.option.ScriptPath)
    Depend_profile(self.option.ValidProf, conv2Form6347(func() {
        var luaPath string
        luaPath = front_convExp5753(Lns_2DDD(Lns_getVM().String_gsub(self.option.ScriptPath,"%.lns$", ".lua")))
        var metaPath string
        metaPath = front_convExp5768(Lns_2DDD(Lns_getVM().String_gsub(self.option.ScriptPath,"%.lns$", ".meta")))
        if Lns_isCondTrue( self.option.OutputDir){
            var filename string
            filename = front_convExp5785(Lns_2DDD(Lns_getVM().String_gsub(mod,"%.", "/")))
            luaPath = Lns_getVM().String_format("%s/%s.lua", []LnsAny{self.option.OutputDir, filename})
            
            metaPath = Lns_getVM().String_format("%s/%s.meta", []LnsAny{self.option.OutputDir, filename})
            
        }
        if luaPath == self.option.ScriptPath{
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
            ast = self.FP.ConvertLuaToStreamFromScript(true, convMode, self.option.ScriptPath, mod, self.option.ByteCompile, self.option.StripDebugInfo, front_OpenOStreamForConvert_1316_(func(mode LnsAny)(LnsAny, LnsAny, LnsAny) {
                var openLuaStream func() LnsAny
                openLuaStream = func() LnsAny {
                    var fileObj Lns_luaStream
                    
                    {
                        _fileObj := front_convExp5943(Lns_2DDD(Lns_io_open(luaPath, "w")))
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
                            __exp := front_convExp6033(Lns_2DDD(Lns_io_open(tempMetaPath, "w+")))
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
                return stream, metaStream, self.option.FP.OpenDepend()
            }), front_CloseOStreamForConvert_1319_(func(stream LnsAny,metaStream LnsAny,dependStream LnsAny) {
                if stream != nil{
                    stream_6226 := stream.(Lns_oStream)
                    stream_6226.Close()
                }
                if dependStream != nil{
                    dependStream_6228 := dependStream.(Lns_oStream)
                    dependStream_6228.Close()
                }
                if metaFileObj != nil{
                    metaFileObj_6230 := metaFileObj.(Lns_luaStream)
                    metaFileObj_6230.Flush()
                    metaFileObj_6230.Seek("set", 0)
                    var newMetaTxt string
                    
                    {
                        _newMetaTxt := metaFileObj_6230.Read("*a")
                        if _newMetaTxt == nil{
                            Util_err(Lns_getVM().String_format("faled to read meta. -- %s.", []LnsAny{tempMetaPath}))
                        } else {
                            newMetaTxt = _newMetaTxt.(string)
                        }
                    }
                    metaFileObj_6230.Close()
                    var oldMetaTxt string
                    oldMetaTxt = ""
                    {
                        _oldFileObj := front_convExp6225(Lns_2DDD(Lns_io_open(metaPath, nil)))
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
                                _fileObj := front_convExp6306(Lns_2DDD(Lns_io_open(metaPath, "w")))
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
    }), self.option.ScriptPath + ".profi")
    if updateFlag{
        Lns_getVM().String_gsub(self.option.ScriptPath,"%.lns$", ".lua")
    }
    if ast != nil{
        ast_6247 := ast.(*TransUnit_ASTInfo)
        if _switch6399 := self.option.ConvTo; _switch6399 == Option_Conv__C {
            self.FP.SaveToC(ast_6247)
        } else if _switch6399 == Option_Conv__Go {
            self.FP.SaveToGo(ast_6247)
        }
    }
    return updateFlag
}

// 1261: decl @lune.@base.@front.Front.outputBootC
func (self *front_Front) outputBootC() {
    var stream Lns_oStream
    {
        _path := self.option.BootPath
        if _path != nil {
            path := _path.(string)
            var cPath string
            cPath = front_convExp6432(Lns_2DDD(Lns_getVM().String_gsub(path,"%.lns$", ".c")))
            var file Lns_luaStream
            
            {
                _file := front_convExp6463(Lns_2DDD(Lns_io_open(cPath, "w")))
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
    initModule = Front_scriptPath2Module(self.option.ScriptPath)
    ConvCC_outputBootcode(stream, initModule)
}

// 1292: decl @lune.@base.@front.Front.exec
func (self *front_Front) Exec() {
    __func__ := "@lune.@base.@front.Front.exec"
    Log_log(Log_Level__Trace, __func__, 1294, Log_CreateMessage(func() string {
        return Option_ModeKind_getTxt( self.option.Mode)
    }))
    
    if _switch6805 := self.option.Mode; _switch6805 == Option_ModeKind__Token {
        self.FP.DumpTokenize()
    } else if _switch6805 == Option_ModeKind__Ast {
        self.FP.DumpAst()
    } else if _switch6805 == Option_ModeKind__Format {
        self.FP.Format()
    } else if _switch6805 == Option_ModeKind__Diag {
        self.FP.CheckDiag()
    } else if _switch6805 == Option_ModeKind__Complete {
        self.FP.Complete()
    } else if _switch6805 == Option_ModeKind__Inquire {
        self.FP.Inquire()
    } else if _switch6805 == Option_ModeKind__Glue {
        self.FP.CreateGlue()
    } else if _switch6805 == Option_ModeKind__Lua || _switch6805 == Option_ModeKind__LuaMeta {
        self.FP.convertToLua()
    } else if _switch6805 == Option_ModeKind__Save || _switch6805 == Option_ModeKind__SaveMeta {
        self.FP.SaveToLua()
    } else if _switch6805 == Option_ModeKind__Exec {
        _ = front_convExp6743(Lns_2DDD(self.FP.LoadModule(Front_scriptPath2Module(self.option.ScriptPath))))
        if self.option.Testing{
            Testing_run(Front_scriptPath2Module(self.option.ScriptPath))
            Testing_outputAllResult(Lns_io_stdout)
        }
    } else if _switch6805 == Option_ModeKind__BootC {
        self.FP.outputBootC()
    } else if _switch6805 == Option_ModeKind__Builtin {
        self.FP.OutputBuiltin()
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
func front_DependMetaInfo__fromMap_1142_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_DependMetaInfo_FromMap( arg1, paramList )
}
func front_DependMetaInfo__fromStem_1145_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
func front_MetaForBuildId__fromMap_1175_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_MetaForBuildId_FromMap( arg1, paramList )
}
func front_MetaForBuildId__fromStem_1178_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 248: decl @lune.@base.@front.MetaForBuildId.createModuleId
func (self *front_MetaForBuildId) CreateModuleId() *FrontInterface_ModuleId {
    return FrontInterface_ModuleId_createIdFromTxt(self.__buildId)
}

// 253: decl @lune.@base.@front.MetaForBuildId.LoadFromMeta
func front_MetaForBuildId_LoadFromMeta_1182_(metaPath string)(LnsAny, LnsAny) {
    {
        _fileObj := front_convExp1135(Lns_2DDD(Lns_io_open(metaPath, nil)))
        if _fileObj != nil {
            fileObj := _fileObj.(Lns_luaStream)
            var luaCode LnsAny
            luaCode = fileObj.Read("*a")
            fileObj.Close()
            if luaCode != nil{
                luaCode_5670 := luaCode.(string)
                var meta LnsAny
                meta = front_convExp1125(Lns_2DDD(front_MetaForBuildId__fromStem_1178_(Lns_getVM().ExpandLuavalMap(front_loadFromLuaTxt_1116_(luaCode_5670)),nil)))
                return meta, luaCode_5670
            }
        }
    }
    return nil, nil
}


func Lns_front_init() {
    if init_front { return }
    init_front = true
    front__mod__ = "@lune.@base.@front"
    Lns_InitMod()
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
