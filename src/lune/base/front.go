// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_front bool
var front__mod__ string
var front_forceUpdateMeta bool
// decl alge -- UptodateInfo
type front_UptodateInfo = LnsAny
type front_UptodateInfo__Update struct{
Val1 *FrontInterface_ModuleInfo
Val2 *TransUnit_ASTInfo
}
func (self *front_UptodateInfo__Update) GetTxt() string {
return "UptodateInfo.Update"
}
type front_UptodateInfo__Uptodate struct{
Val1 *FrontInterface_ModuleMeta
}
func (self *front_UptodateInfo__Uptodate) GetTxt() string {
return "UptodateInfo.Uptodate"
}
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
// decl alge -- BuildMode
type front_BuildMode = LnsAny
type front_BuildMode__CreateAst struct{
}
var front_BuildMode__CreateAst_Obj = &front_BuildMode__CreateAst{}
func (self *front_BuildMode__CreateAst) GetTxt() string {
return "BuildMode.CreateAst"
}
type front_BuildMode__Output struct{
Val1 Lns_oStream
Val2 Lns_oStream
}
func (self *front_BuildMode__Output) GetTxt() string {
return "BuildMode.Output"
}
type front_BuildMode__Save struct{
}
var front_BuildMode__Save_Obj = &front_BuildMode__Save{}
func (self *front_BuildMode__Save) GetTxt() string {
return "BuildMode.Save"
}
type Front_AstCallback func (_env *LnsEnv, arg1 *TransUnit_ASTInfo)
type front_OpenOStreamForConvert_1308_ func (_env *LnsEnv, arg1 LnsAny)(LnsAny, LnsAny, LnsAny)
type front_CloseOStreamForConvert_1310_ func (_env *LnsEnv, arg1 LnsAny,arg2 LnsAny,arg3 LnsAny)
// for 1018: ExpCast
func conv2Form5369( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1737: ExpCast
func conv2Form8799( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1839: ExpCast
func conv2Form9159( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 315
func front_convExp1291(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 393
func front_convExp1684(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 412
func front_convExp1778(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 688
func front_convExp3182(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 697
func front_convExp3211(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1563
func front_convExp7911(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1602
func front_convExp8083(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1616
func front_convExp8164(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1812
func front_convExp9101(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1838
func front_convExp9177(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 411
func front_convExp1757(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 438
func front_convExp1911(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 439
func front_convExp1930(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 479
func front_convExp2118(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 681
func front_convExp3114(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1268
func front_convExp6505(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1578
func front_convExp7953(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1581
func front_convExp7971(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 325
func front_convExp1328(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 241
func front_convExp875(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 337
func front_convExp1353(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 372
func front_convExp1555(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 397
func front_convExp1674(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 559
func front_convExp2408(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 586
func front_convExp2561(arg1 []LnsAny) (LnsAny, string, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 631
func front_convExp2817(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 669
func front_convExp2992(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 737
func front_convExp3474(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 802
func front_convExp3872(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 845
func front_convExp4200(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 854
func front_convExp4258(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 876
func front_convExp4372(arg1 []LnsAny) (*FrontInterface_ModuleMeta, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 )
}
// for 924
func front_convExp4676(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 949
func front_convExp4774(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 976
func front_convExp5070(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 985
func front_convExp5161(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1139
func front_convExp5903(arg1 []LnsAny) (LnsAny, LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 )
}
// for 1156
func front_convExp5981(arg1 []LnsAny) (LnsAny, LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 )
}
// for 1222
func front_convExp6291(arg1 []LnsAny) (LnsAny, LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 )
}
// for 1259
func front_convExp6448(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1331
func front_convExp6783(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1336
func front_convExp6822(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1363
func front_convExp6949(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1365
func front_convExp6964(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1369
func front_convExp6978(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1370
func front_convExp6994(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1414
func front_convExp7155(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1473
func front_convExp7404(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1474
func front_convExp7417(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1515
func front_convExp7655(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1516
func front_convExp7668(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1518
func front_convExp7685(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1551
func front_convExp7835(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1607
func front_convExp8107(arg1 []LnsAny) (bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1684
func front_convExp8377(arg1 []LnsAny) (*FrontInterface_ModuleId, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 1 )
}
// for 1716
func front_convExp8532(arg1 []LnsAny) (bool, LnsAny, *Util_memStream) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 ).(*Util_memStream)
}
// for 1826
func front_convExp9123(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1836
func front_convExp9142(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
func front__anonymous_1004_(_env *LnsEnv, ver LnsInt) {
    LuaVer_setCurVer(_env, ver)
}
// 249: decl @lune.@base.@front.createPaser
func front_createPaser_1135_(_env *LnsEnv, path string,mod string) *Parser_Parser {
    var parser LnsAny
    parser = Parser_StreamParser_create(_env, path, false, mod)
    {
        __exp := parser
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Parser_StreamParser)
            return &_exp.Parser_Parser
        }
    }
    panic("failed to open " + path)
// insert a dummy
    return nil
}

// 310: decl @lune.@base.@front.loadFromChunk
func front_loadFromChunk_1147_(_env *LnsEnv, chunk LnsAny,err LnsAny) LnsAny {
    if err != nil{
        err_237 := err.(string)
        Util_errorLog(_env, err_237)
    }
    if chunk != nil{
        chunk_239 := chunk.(*Lns_luaValue)
        {
            _work := front_convExp1291(Lns_2DDD(_env.LuaVM.RunLoadedfunc(chunk_239,Lns_2DDD([]LnsAny{}))[0]))
            if !Lns_IsNil( _work ) {
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

// 323: decl @lune.@base.@front.loadFromLuaTxt
func Front_loadFromLuaTxt(_env *LnsEnv, txt string) LnsAny {
    return front_loadFromChunk_1147_(front_convExp1328(_env, Lns_2DDD(_env.LuaVM.Load(txt, nil))))
}

// 336: decl @lune.@base.@front.byteCompileFromLuaTxt
func front_byteCompileFromLuaTxt_1151_(_env *LnsEnv, txt string,stripDebugInfo bool) string {
    var chunk LnsAny
    var err LnsAny
    chunk,err = _env.LuaVM.Load(txt, nil)
    if chunk != nil{
        chunk_254 := chunk.(*Lns_luaValue)
        return _env.LuaVM.String_dump(chunk_254, stripDebugInfo)
    }
    panic(Lns_unwrapDefault( err, "load error").(string))
// insert a dummy
    return ""
}

// 406: decl @lune.@base.@front.getMetaInfo
func front_getMetaInfo_1198_(_env *LnsEnv, lnsPath string,mod string,outdir LnsAny)(LnsAny, string, string) {
    var moduleMetaPath string
    moduleMetaPath = lnsPath
    if outdir != nil{
        outdir_315 := outdir.(string)
        moduleMetaPath = _env.LuaVM.String_format("%s/%s", []LnsAny{outdir_315, Lns_car(_env.LuaVM.String_gsub(mod,"%.", "/")).(string)})
        
    }
    moduleMetaPath = front_convExp1757(Lns_2DDD(_env.LuaVM.String_gsub(moduleMetaPath,"%.lns$", ".meta")))
    
    {
        _meta, _metaCode := front_MetaForBuildId_LoadFromMeta_1196_(_env, moduleMetaPath)
        if !Lns_IsNil( _meta ) && !Lns_IsNil( _metaCode ) {
            meta := _meta.(*front_MetaForBuildId)
            metaCode := _metaCode.(string)
            return meta, moduleMetaPath, metaCode
        }
    }
    return nil, moduleMetaPath, ""
}


// 471: decl @lune.@base.@front.getModuleId
func front_getModuleId_1204_(_env *LnsEnv, lnsPath string,mod string,outdir LnsAny,metaInfo LnsAny) *FrontInterface_ModuleId {
    var buildCount LnsInt
    buildCount = 0
    var fileTime LnsReal
    
    {
        _fileTime := Depend_getFileLastModifiedTime(_env, lnsPath)
        if _fileTime == nil{
            return FrontInterface_ModuleId__tempId
        } else {
            fileTime = _fileTime.(LnsReal)
        }
    }
    if Lns_op_not(metaInfo){
        metaInfo = front_convExp2118(Lns_2DDD(front_getMetaInfo_1198_(_env, lnsPath, mod, outdir)))
        
    }
    if metaInfo != nil{
        metaInfo_360 := metaInfo.(*front_MetaForBuildId)
        var buildId *FrontInterface_ModuleId
        buildId = metaInfo_360.FP.CreateModuleId(_env)
        buildCount = buildId.FP.Get_buildCount(_env)
        
    }
    return FrontInterface_ModuleId_createId(_env, fileTime, buildCount)
}

func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1215_(_env *LnsEnv) string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1217_(_env *LnsEnv) string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1219_(_env *LnsEnv) string {
    return "NeedUpdate"
}


func Front_getModuleIdAndCheckUptodate___anonymous_1223_(_env *LnsEnv) string {
    return "not found meta"
}
// 637: decl @lune.@base.@front.createModuleInfo
func front_createModuleInfo_1227_(_env *LnsEnv, ast *TransUnit_ASTInfo,mod string,moduleId *FrontInterface_ModuleId) *FrontInterface_ModuleInfo {
    var importedAliasMap *LnsMap
    importedAliasMap = NewLnsMap( map[LnsAny]LnsAny{})
    var rootNode *Nodes_RootNode
    rootNode = Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode)
    for _, _node := range( rootNode.FP.Get_nodeManager(_env).FP.GetAliasNodeList(_env).Items ) {
        node := _node.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
        importedAliasMap.Set(node.FP.Get_typeInfo(_env),Ast_AliasTypeInfoDownCastF(node.FP.Get_expType(_env).FP))
    }
    var exportInfo *FrontInterface_ExportInfo
    exportInfo = ast.FP.Get_exportInfo(_env)
    return NewFrontInterface_ModuleInfo(_env, mod, exportInfo.FP.Get_moduleTypeInfo(_env).FP.Get_rawTxt(_env), NewLnsMap( map[LnsAny]LnsAny{}), moduleId, exportInfo, importedAliasMap)
}





















func Front_convertToLua___anonymous_1328_(_env *LnsEnv, luaCodeStream LnsAny,metaStream LnsAny,dependStream LnsAny) {
    if dependStream != nil{
        dependStream_848 := dependStream.(Lns_oStream)
        dependStream_848.Close(_env)
    }
}
// 1413: decl @lune.@base.@front.Front.saveToLua.txt2ModuleId
func Front_saveToLua__txt2ModuleId_1357_(_env *LnsEnv, txt string) *FrontInterface_ModuleId {
    var buildIdTxt string
    buildIdTxt = front_convExp7155(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(txt,"^_moduleObj.__buildId = ", "")).(string),"\"", "")))
    return FrontInterface_ModuleId_createIdFromTxt(_env, buildIdTxt)
}







// 1662: decl @lune.@base.@front.convertLnsCode2LuaCode
func Front_convertLnsCode2LuaCode(_env *LnsEnv, lnsCode string,path string) string {
    var option *Option_Option
    option = NewOption_Option(_env)
    option.ScriptPath = path
    
    option.UseLuneModule = Option_getRuntimeModule(_env)
    
    option.UseIpairs = true
    
    var front *front_Front
    front = Newfront_Front(_env, option, nil)
    return front.FP.ConvertLns2LuaCode(_env, NewFrontInterface_ImportModuleInfo(_env), NewParser_TxtStream(_env, lnsCode).FP, path)
}





// 1766: decl @lune.@base.@front.build
func Front_build(_env *LnsEnv, option *Option_Option,astCallback Front_AstCallback) {
    var front *front_Front
    front = Newfront_Front(_env, option, nil)
    front.FP.Build(_env, front_BuildMode__CreateAst_Obj, astCallback)
}


// 1867: decl @lune.@base.@front.exec
func Front_exec(_env *LnsEnv, args *LnsList) {
    var version LnsReal
    version = Lns_unwrapDefault( Lns_tonumber(Lns_car(_env.LuaVM.String_gsub(Depend_getLuaVersion(_env),"^[^%d]+", "")).(string), nil), 0.0).(LnsReal)
    if version < 5.1{
        Lns_io_stderr.Write(_env, _env.LuaVM.String_format("LuneScript doesn't support this lua version(%s). %s\n", []LnsAny{version, "please use the version >= 5.1."}))
        _env.LuaVM.OS_exit(1)
    }
    var option *Option_Option
    option = Option_analyze(_env, args)
    var front *front_Front
    front = Newfront_Front(_env, option, nil)
    front.FP.Exec(_env)
}

// 1886: decl @lune.@base.@front.setFront
func Front_setFront(_env *LnsEnv, bindModuleList *LnsList) {
    var option *Option_Option
    option = Option_createDefaultOption(_env, NewLnsList([]LnsAny{"dummy.lns"}), nil)
    Newfront_Front(_env, option, bindModuleList)
}

// 1891: decl @lune.@base.@front.__main
func Front___main(_env *LnsEnv, argList *LnsList) LnsInt {
    Lns_front_init( _env )
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    for _index, _arg := range( argList.Items ) {
        index := _index + 1
        arg := _arg.(string)
        if index > 1{
            list.Insert(arg)
        }
    }
    Front_exec(_env, list)
    return 0
}

// declaration Class -- LoadInfo
type front_LoadInfoMtd interface {
}
type front_LoadInfo struct {
    Mod LnsAny
    Meta *FrontInterface_ModuleMeta
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
func Newfront_LoadInfo(_env *LnsEnv, arg1 LnsAny, arg2 *FrontInterface_ModuleMeta) *front_LoadInfo {
    obj := &front_LoadInfo{}
    obj.FP = obj
    obj.Initfront_LoadInfo(_env, arg1, arg2)
    return obj
}
func (self *front_LoadInfo) Initfront_LoadInfo(_env *LnsEnv, arg1 LnsAny, arg2 *FrontInterface_ModuleMeta) {
    self.Mod = arg1
    self.Meta = arg2
}

// declaration Class -- ModuleMgr
type front_ModuleMgrMtd interface {
    Add(_env *LnsEnv, arg1 string, arg2 *TransUnit_ASTInfo, arg3 *FrontInterface_ModuleId) *FrontInterface_ModuleInfo
    AddMeta(_env *LnsEnv, arg1 string, arg2 *FrontInterface_ModuleMeta)
    createModuleInfo(_env *LnsEnv, arg1 *TransUnit_ASTInfo, arg2 string, arg3 *FrontInterface_ModuleId) *FrontInterface_ModuleInfo
    Get(_env *LnsEnv, arg1 string) LnsAny
    GetAst(_env *LnsEnv, arg1 string) LnsAny
    GetMeta(_env *LnsEnv, arg1 string) LnsAny
    GetModList(_env *LnsEnv) *LnsList
}
type front_ModuleMgr struct {
    mod2info *Util_OrderdMap
    loadedMetaMap *LnsMap
    FP front_ModuleMgrMtd
}
func front_ModuleMgr2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_ModuleMgr).FP
}
type front_ModuleMgrDownCast interface {
    Tofront_ModuleMgr() *front_ModuleMgr
}
func front_ModuleMgrDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(front_ModuleMgrDownCast)
    if ok { return work.Tofront_ModuleMgr() }
    return nil
}
func (obj *front_ModuleMgr) Tofront_ModuleMgr() *front_ModuleMgr {
    return obj
}
func Newfront_ModuleMgr(_env *LnsEnv) *front_ModuleMgr {
    obj := &front_ModuleMgr{}
    obj.FP = obj
    obj.Initfront_ModuleMgr(_env)
    return obj
}
// 80: DeclConstr
func (self *front_ModuleMgr) Initfront_ModuleMgr(_env *LnsEnv) {
    self.mod2info = NewUtil_OrderdMap(_env)
    
    self.loadedMetaMap = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 85: decl @lune.@base.@front.ModuleMgr.get
func (self *front_ModuleMgr) Get(_env *LnsEnv, mod string) LnsAny {
    return self.mod2info.FP.Get_map(_env).Get(mod)
}

// 89: decl @lune.@base.@front.ModuleMgr.getAst
func (self *front_ModuleMgr) GetAst(_env *LnsEnv, mod string) LnsAny {
    var info LnsAny
    
    {
        _info := self.FP.Get(_env, mod)
        if _info == nil{
            return nil
        } else {
            info = _info
        }
    }
    switch _exp180 := info.(type) {
    case *front_UptodateInfo__Update:
    ast := _exp180.Val2
        return ast
    case *front_UptodateInfo__Uptodate:
        return nil
    }
// insert a dummy
    return nil
}

// 103: decl @lune.@base.@front.ModuleMgr.getModList
func (self *front_ModuleMgr) GetModList(_env *LnsEnv) *LnsList {
    return self.mod2info.FP.Get_keyList(_env)
}

// 108: decl @lune.@base.@front.ModuleMgr.createModuleInfo
func (self *front_ModuleMgr) createModuleInfo(_env *LnsEnv, ast *TransUnit_ASTInfo,mod string,moduleId *FrontInterface_ModuleId) *FrontInterface_ModuleInfo {
    var importedAliasMap *LnsMap
    importedAliasMap = NewLnsMap( map[LnsAny]LnsAny{})
    var rootNode *Nodes_RootNode
    rootNode = Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode)
    for _, _node := range( rootNode.FP.Get_nodeManager(_env).FP.GetAliasNodeList(_env).Items ) {
        node := _node.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
        importedAliasMap.Set(node.FP.Get_typeInfo(_env),Ast_AliasTypeInfoDownCastF(node.FP.Get_expType(_env).FP))
    }
    var exportInfo *FrontInterface_ExportInfo
    exportInfo = ast.FP.Get_exportInfo(_env)
    return NewFrontInterface_ModuleInfo(_env, mod, exportInfo.FP.Get_moduleTypeInfo(_env).FP.Get_rawTxt(_env), NewLnsMap( map[LnsAny]LnsAny{}), moduleId, exportInfo, importedAliasMap)
}

// 123: decl @lune.@base.@front.ModuleMgr.add
func (self *front_ModuleMgr) Add(_env *LnsEnv, mod string,ast *TransUnit_ASTInfo,moduleId *FrontInterface_ModuleId) *FrontInterface_ModuleInfo {
    if Lns_isCondTrue( self.loadedMetaMap.Get(mod)){
        if mod == "lune.base.Testing"{
        } else { 
            Util_err(_env, _env.LuaVM.String_format("the meta is already loaded -- %s", []LnsAny{mod}))
        }
    }
    var moduleInfo *FrontInterface_ModuleInfo
    moduleInfo = self.FP.createModuleInfo(_env, ast, mod, moduleId)
    self.mod2info.FP.Add(_env, mod, &front_UptodateInfo__Update{moduleInfo, ast})
    return moduleInfo
}

// 141: decl @lune.@base.@front.ModuleMgr.addMeta
func (self *front_ModuleMgr) AddMeta(_env *LnsEnv, mod string,meta *FrontInterface_ModuleMeta) {
    if Lns_op_not(self.FP.Get(_env, mod)){
        self.mod2info.FP.Add(_env, mod, &front_UptodateInfo__Uptodate{meta})
    }
    self.loadedMetaMap.Set(mod,meta)
}

// 148: decl @lune.@base.@front.ModuleMgr.getMeta
func (self *front_ModuleMgr) GetMeta(_env *LnsEnv, mod string) LnsAny {
    return self.loadedMetaMap.Get(mod)
}


// declaration Class -- Front
type front_FrontMtd interface {
    Build(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny)
    CheckDiag(_env *LnsEnv, arg1 string)
    checkUptodateMeta(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny) LnsAny
    Complete(_env *LnsEnv, arg1 string)
    convert(_env *LnsEnv, arg1 *TransUnit_ASTInfo, arg2 string, arg3 Lns_oStream, arg4 Lns_oStream, arg5 LnsInt, arg6 bool)
    convertFromAst(_env *LnsEnv, arg1 *TransUnit_ASTInfo, arg2 string, arg3 LnsInt)(string, string)
    ConvertLns2LuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 Lns_iStream, arg3 string) string
    ConvertLuaToStreamFromScript(_env *LnsEnv, arg1 LnsAny, arg2 *FrontInterface_ModuleId, arg3 LnsAny, arg4 LnsInt, arg5 string, arg6 string, arg7 bool, arg8 bool, arg9 front_OpenOStreamForConvert_1308_, arg10 LnsAny) LnsAny
    convertToLua(_env *LnsEnv, arg1 string, arg2 LnsInt, arg3 Lns_oStream, arg4 Lns_oStream) LnsAny
    createAst(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 *Parser_Parser, arg3 string, arg4 *FrontInterface_ModuleId, arg5 LnsAny, arg6 LnsInt, arg7 LnsAny) *TransUnit_ASTInfo
    CreateGlue(_env *LnsEnv, arg1 string)
    createGoOption(_env *LnsEnv, arg1 string) *ConvGo_Option
    createPaser(_env *LnsEnv, arg1 string) *Parser_Parser
    DumpAst(_env *LnsEnv, arg1 string)
    DumpTokenize(_env *LnsEnv, arg1 string)
    Error(_env *LnsEnv, arg1 string)
    Exec(_env *LnsEnv)
    Format(_env *LnsEnv, arg1 string)
    getGoAppName(_env *LnsEnv) string
    getLoadInfo(_env *LnsEnv, arg1 string) LnsAny
    GetLuaModulePath(_env *LnsEnv, arg1 string) string
    getModuleIdAndCheckUptodate(_env *LnsEnv, arg1 string, arg2 string)(*FrontInterface_ModuleId, LnsAny)
    Inquire(_env *LnsEnv, arg1 string)
    loadFile(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string)(*FrontInterface_ModuleMeta, LnsAny)
    loadFileToLuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string)(*FrontInterface_ModuleMeta, string)
    LoadFromLnsTxt(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string) LnsAny
    loadLua(_env *LnsEnv, arg1 string) LnsAny
    LoadMeta(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string) LnsAny
    LoadModule(_env *LnsEnv, arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
    loadParserToLuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 *Parser_Parser, arg3 string)(*FrontInterface_ModuleMeta, string)
    outputBootC(_env *LnsEnv, arg1 string)
    OutputBuiltin(_env *LnsEnv, arg1 string)
    SaveToC(_env *LnsEnv, arg1 string, arg2 *TransUnit_ASTInfo)
    SaveToGo(_env *LnsEnv, arg1 string, arg2 *TransUnit_ASTInfo)
    SaveToLua(_env *LnsEnv, arg1 *front_UpdateInfo)(bool, LnsAny, *Util_memStream)
    scriptPath2Module(_env *LnsEnv, arg1 string) string
    searchLuaFile(_env *LnsEnv, arg1 string, arg2 LnsAny) LnsAny
    SearchModule(_env *LnsEnv, arg1 string) LnsAny
    searchModuleFile(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny) LnsAny
    setLoadInfo(_env *LnsEnv, arg1 string, arg2 *front_LoadInfo)
}
type front_Front struct {
    option *Option_Option
    loadedMap *LnsMap
    loadedMapTest *LnsMap
    convertedMap *LnsMap
    gomodMap *GoMod_ModInfo
    bindModuleSet *LnsSet
    moduleMgr *front_ModuleMgr
    targetSet *LnsSet
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
func Newfront_Front(_env *LnsEnv, arg1 *Option_Option, arg2 LnsAny) *front_Front {
    obj := &front_Front{}
    obj.FP = obj
    obj.Initfront_Front(_env, arg1, arg2)
    return obj
}
// 176: DeclConstr
func (self *front_Front) Initfront_Front(_env *LnsEnv, option *Option_Option,bindModuleList LnsAny) {
    self.targetSet = NewLnsSet([]LnsAny{})
    
    self.bindModuleSet = NewLnsSet([]LnsAny{})
    
    if bindModuleList != nil{
        bindModuleList_123 := bindModuleList.(*LnsList)
        for _, _mod := range( bindModuleList_123.Items ) {
            mod := _mod.(string)
            self.bindModuleSet.Add(mod)
        }
    }
    self.moduleMgr = Newfront_ModuleMgr(_env)
    
    self.gomodMap = GoMod_getGoMap(_env)
    
    self.option = option
    
    self.loadedMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.loadedMapTest = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.convertedMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    FrontInterface_setFront(_env, self.FP)
    var loadedMap *LnsMap
    loadedMap = NewLnsMap( map[LnsAny]LnsAny{})
    {
        _exp688 := Depend_getLoadedMod(_env)
        _key688, _val688 := _exp688.Get1stFromMap()
        for _key688 != nil {
            mod := _key688.(string)
            modval := _val688
            if mod == "lune.base.Testing"{
                loadedMap.Set(mod,modval)
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( option.Testing) &&
                _env.SetStackVal( modval.(*Lns_luaValue).GetAt("__enableTest")) ||
                _env.SetStackVal( Lns_op_not(option.Testing)) &&
                _env.SetStackVal( modval.(*Lns_luaValue).GetAt("__enableTest")) )){
                loadedMap.Set(mod,modval)
            }
            _key688, _val688 = _exp688.NextFromMap( _key688 )
        }
    }
    self.preloadedModMap = loadedMap
    
}

// 209: decl @lune.@base.@front.Front.getLoadInfo
func (self *front_Front) getLoadInfo(_env *LnsEnv, mod string) LnsAny {
    if self.option.Testing{
        return self.loadedMapTest.Get(mod)
    }
    return self.loadedMap.Get(mod)
}

// 216: decl @lune.@base.@front.Front.setLoadInfo
func (self *front_Front) setLoadInfo(_env *LnsEnv, mod string,info *front_LoadInfo) {
    if self.option.Testing{
        self.loadedMapTest.Set(mod,info)
    }
    self.loadedMap.Set(mod,info)
}







// 234: decl @lune.@base.@front.Front.error
func (self *front_Front) Error(_env *LnsEnv, message string) {
    Util_errorLog(_env, message)
    Util_printStackTrace(_env)
    _env.LuaVM.OS_exit(1)
}

// 240: decl @lune.@base.@front.Front.loadLua
func (self *front_Front) loadLua(_env *LnsEnv, path string) LnsAny {
    var chunk LnsAny
    var err LnsAny
    chunk,err = _env.LuaVM.Loadfile(path)
    if chunk != nil{
        chunk_179 := chunk.(*Lns_luaValue)
        return Lns_unwrap( Lns_car(_env.LuaVM.RunLoadedfunc(chunk_179,Lns_2DDD([]LnsAny{}))[0]))
    }
    Util_errorLog(_env, Lns_unwrapDefault( err, _env.LuaVM.String_format("load error -- %s.", []LnsAny{path})).(string))
    return nil
}

// 258: decl @lune.@base.@front.Front.scriptPath2Module
func (self *front_Front) scriptPath2Module(_env *LnsEnv, path string) string {
    return Util_scriptPath2ModuleFromProjDir(_env, path, self.option.FP.Get_projDir(_env))
}

// 262: decl @lune.@base.@front.Front.createPaser
func (self *front_Front) createPaser(_env *LnsEnv, scriptPath string) *Parser_Parser {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    return front_createPaser_1135_(_env, scriptPath, mod)
}

// 269: decl @lune.@base.@front.Front.createAst
func (self *front_Front) createAst(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parser *Parser_Parser,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *TransUnit_ASTInfo {
    {
        __exp := self.moduleMgr.FP.Get(_env, mod)
        if !Lns_IsNil( __exp ) {
            _exp := __exp
            switch _exp1079 := _exp.(type) {
            case *front_UptodateInfo__Update:
            ast := _exp1079.Val2
                return ast
            case *front_UptodateInfo__Uptodate:
                Util_err(_env, _env.LuaVM.String_format("can't load the multiple module -- %s", []LnsAny{mod}))
            }
        }
    }
    var transUnit *TransUnit_TransUnit
    transUnit = NewTransUnit_TransUnit(_env, moduleId, importModuleInfo, &NewConvLua_MacroEvalImp(_env).Nodes_MacroEval, analyzeModule, analyzeMode, pos, self.option.TargetLuaVer, self.option.TransCtrlInfo)
    var ast *TransUnit_ASTInfo
    ast = transUnit.FP.CreateAST(_env, parser, false, mod)
    self.moduleMgr.FP.Add(_env, mod, ast, moduleId)
    return ast
}

// 296: decl @lune.@base.@front.Front.convert
func (self *front_Front) convert(_env *LnsEnv, ast *TransUnit_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool) {
    var exportInfo *FrontInterface_ExportInfo
    exportInfo = ast.FP.Get_exportInfo(_env)
    var conv *Nodes_Filter
    conv = ConvLua_createFilter(_env, streamName, stream, metaStream, convMode, inMacro, exportInfo.FP.Get_moduleTypeInfo(_env), exportInfo.FP.Get_processInfo(_env), exportInfo.FP.Get_provideInfo(_env).FP.Get_symbolKind(_env), self.option.UseLuneModule, self.option.TargetLuaVer, self.option.Testing, self.option.UseIpairs)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvLua_Opt2Stem(NewConvLua_Opt(_env, ast.FP.Get_node(_env))))
}

// 345: decl @lune.@base.@front.Front.convertFromAst
func (self *front_Front) convertFromAst(_env *LnsEnv, ast *TransUnit_ASTInfo,streamName string,mode LnsInt)(string, string) {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var metaStream *Util_memStream
    metaStream = NewUtil_memStream(_env)
    self.FP.convert(_env, ast, streamName, stream.FP, metaStream.FP, mode, false)
    return metaStream.FP.Get_txt(_env), stream.FP.Get_txt(_env)
}

// 359: decl @lune.@base.@front.Front.loadFromLnsTxt
func (self *front_Front) LoadFromLnsTxt(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    var transUnit *TransUnit_TransUnit
    transUnit = NewTransUnit_TransUnit(_env, FrontInterface_ModuleId__tempId, importModuleInfo, &NewConvLua_MacroEvalImp(_env).Nodes_MacroEval, nil, nil, nil, self.option.TargetLuaVer, self.option.TransCtrlInfo)
    var stream *Parser_TxtStream
    stream = NewParser_TxtStream(_env, txt)
    var parser *Parser_StreamParser
    parser = NewParser_StreamParser(_env, stream.FP, name, false, nil)
    var ast *TransUnit_ASTInfo
    ast = transUnit.FP.CreateAST(_env, &parser.Parser_Parser, false, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(_env, ast, name, ConvLua_ConvMode__ConvMeta)
    return Lns_unwrap( Front_loadFromLuaTxt(_env, luaTxt))
}

// 418: decl @lune.@base.@front.Front.searchModuleFile
func (self *front_Front) searchModuleFile(_env *LnsEnv, mod string,suffix string,addPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.searchModuleFile"
    switch _exp1827 := self.gomodMap.FP.ConvLocalModulePath(_env, mod, suffix).(type) {
    case *GoMod_GoModResult__NotGo:
    case *GoMod_GoModResult__NotFound:
        return nil
    case *GoMod_GoModResult__Found:
    info := _exp1827.Val1
        return info.FP.Get_path(_env)
    }
    var lnsSearchPath string
    lnsSearchPath = Lns_package_path
    if addPath != nil{
        addPath_332 := addPath.(string)
        lnsSearchPath = _env.LuaVM.String_format("%s/?%s;%s", []LnsAny{addPath_332, suffix, Lns_package_path})
        
    }
    {
        _projDir := self.option.FP.Get_projDir(_env)
        if !Lns_IsNil( _projDir ) {
            projDir := _projDir.(string)
            lnsSearchPath = _env.LuaVM.String_format("%s;%s", []LnsAny{Util_pathJoin(_env, projDir, _env.LuaVM.String_format("?%s", []LnsAny{suffix})), Lns_package_path})
            
        }
    }
    lnsSearchPath = front_convExp1911(Lns_2DDD(_env.LuaVM.String_gsub(lnsSearchPath,"%.lua$", suffix)))
    
    lnsSearchPath = front_convExp1930(Lns_2DDD(_env.LuaVM.String_gsub(lnsSearchPath,"%.lua;", suffix + ";")))
    
    var foundPath string
    
    {
        _foundPath := Depend_searchpath(_env, mod, lnsSearchPath)
        if _foundPath == nil{
            var latestProjRoot string
            
            {
                _latestProjRoot := self.gomodMap.FP.GetLatestProjRoot(_env)
                if _latestProjRoot == nil{
                    return nil
                } else {
                    latestProjRoot = _latestProjRoot.(string)
                }
            }
            var latestProjSearchPath string
            latestProjSearchPath = Util_pathJoin(_env, latestProjRoot, "?" + suffix)
            {
                __exp := Depend_searchpath(_env, mod, latestProjSearchPath)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(string)
                    foundPath = _exp
                    
                } else {
                    Log_log(_env, Log_Level__Err, __func__, 451, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.LuaVM.String_format("not found at %s", []LnsAny{latestProjSearchPath})
                    }))
                    
                    return nil
                }
            }
        } else {
            foundPath = _foundPath.(string)
        }
    }
    return Lns_car(_env.LuaVM.String_gsub(foundPath,"^./", "")).(string)
}

// 511: decl @lune.@base.@front.Front.getModuleIdAndCheckUptodate
func (self *front_Front) getModuleIdAndCheckUptodate(_env *LnsEnv, lnsPath string,mod string)(*FrontInterface_ModuleId, LnsAny) {
    __func__ := "@lune.@base.@front.Front.getModuleIdAndCheckUptodate"
    var uptodate LnsAny
    uptodate = front_ModuleUptodate__NeedUpdate_Obj
    switch _exp2217 := self.option.TransCtrlInfo.UptodateMode.(type) {
    case *Types_CheckingUptodateMode__ForceAll:
        return FrontInterface_ModuleId__tempId, uptodate
    case *Types_CheckingUptodateMode__Force1:
    forceMod := _exp2217.Val1
        if mod == forceMod{
            return FrontInterface_ModuleId__tempId, uptodate
        }
    case *Types_CheckingUptodateMode__Normal:
    case *Types_CheckingUptodateMode__Touch:
    }
    var checkDependUptodate func(_env *LnsEnv, metaTime LnsReal,metaInfo *front_MetaForBuildId,metaCode string) LnsAny
    checkDependUptodate = func(_env *LnsEnv, metaTime LnsReal,metaInfo *front_MetaForBuildId,metaCode string) LnsAny {
        __func__ := "@lune.@base.@front.Front.getModuleIdAndCheckUptodate.checkDependUptodate"
        for _depMod, _dependItem := range( metaInfo.G__dependModuleMap.Items ) {
            depMod := _depMod.(string)
            dependItem := _dependItem.(front_DependMetaInfoDownCast).Tofront_DependMetaInfo()
            var modMetaPath string
            
            {
                _modMetaPath := self.FP.searchModuleFile(_env, depMod, ".meta", self.option.OutputDir)
                if _modMetaPath == nil{
                    Log_log(_env, Log_Level__Debug, __func__, 547, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1215_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    modMetaPath = _modMetaPath.(string)
                }
            }
            var time LnsReal
            
            {
                _time := Depend_getFileLastModifiedTime(_env, modMetaPath)
                if _time == nil{
                    Log_log(_env, Log_Level__Debug, __func__, 552, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1217_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    time = _time.(LnsReal)
                }
            }
            if time > metaTime{
                var dependMeta *front_MetaForBuildId
                
                {
                    _dependMeta := front_convExp2408(Lns_2DDD(front_MetaForBuildId_LoadFromMeta_1196_(_env, modMetaPath)))
                    if _dependMeta == nil{
                        Log_log(_env, Log_Level__Debug, __func__, 560, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1219_))
                        
                        return front_ModuleUptodate__NeedUpdate_Obj
                    } else {
                        dependMeta = _dependMeta.(*front_MetaForBuildId)
                    }
                }
                var orgMetaModuleId *FrontInterface_ModuleId
                orgMetaModuleId = FrontInterface_ModuleId_createIdFromTxt(_env, dependItem.BuildId)
                var metaModuleId *FrontInterface_ModuleId
                metaModuleId = dependMeta.FP.CreateModuleId(_env)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( metaModuleId.FP.Get_buildCount(_env) != 0) &&
                    _env.SetStackVal( metaModuleId.FP.Get_buildCount(_env) != orgMetaModuleId.FP.Get_buildCount(_env)) ).(bool)){
                    Log_log(_env, Log_Level__Debug, __func__, 570, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.LuaVM.String_format("NeedUpdate: %s, %d, %d", []LnsAny{modMetaPath, metaModuleId.FP.Get_buildCount(_env), orgMetaModuleId.FP.Get_buildCount(_env)})
                    }))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                }
            }
        }
        if self.option.TransCtrlInfo.UptodateMode == Types_CheckingUptodateMode__Touch_Obj{
            return &front_ModuleUptodate__NeedTouch{metaCode, metaInfo}
        }
        return &front_ModuleUptodate__Uptodate{metaInfo}
    }
    var metaInfo LnsAny
    var metaPath string
    var metaCode string
    metaInfo,metaPath,metaCode = front_getMetaInfo_1198_(_env, lnsPath, mod, self.option.OutputDir)
    if metaInfo != nil{
        metaInfo_415 := metaInfo.(*front_MetaForBuildId)
        if metaInfo_415.G__enableTest == self.option.Testing{
            var buildId *FrontInterface_ModuleId
            buildId = FrontInterface_ModuleId_createIdFromTxt(_env, metaInfo_415.G__buildId)
            if buildId != FrontInterface_ModuleId__tempId{
                var lnsTime LnsAny
                lnsTime = Depend_getFileLastModifiedTime(_env, lnsPath)
                var metaTime LnsAny
                metaTime = Depend_getFileLastModifiedTime(_env, metaPath)
                if lnsTime != nil && metaTime != nil{
                    lnsTime_422 := lnsTime.(LnsReal)
                    metaTime_423 := metaTime.(LnsReal)
                    if lnsTime_422 == buildId.FP.Get_modTime(_env){
                        uptodate = checkDependUptodate(_env, metaTime_423, metaInfo_415, metaCode)
                        
                    }
                }
            }
        } else { 
        }
    } else {
        Log_log(_env, Log_Level__Debug, __func__, 608, Log_CreateMessage(Front_getModuleIdAndCheckUptodate___anonymous_1223_))
        
    }
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_1204_(_env, lnsPath, mod, self.option.OutputDir, metaInfo)
    if moduleId == FrontInterface_ModuleId__tempId{
        Util_err(_env, _env.LuaVM.String_format("not found -- %s", []LnsAny{lnsPath}))
    }
    return moduleId, uptodate
}

// 621: decl @lune.@base.@front.Front.convertLns2LuaCode
func (self *front_Front) ConvertLns2LuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,stream Lns_iStream,streamName string) string {
    var mod string
    mod = self.FP.scriptPath2Module(_env, streamName)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, &NewParser_StreamParser(_env, stream, streamName, false, nil).Parser_Parser, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(_env, ast, streamName, ConvLua_ConvMode__ConvMeta)
    return luaTxt
}

// 658: decl @lune.@base.@front.Front.loadParserToLuaCode
func (self *front_Front) loadParserToLuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parser *Parser_Parser,mod string)(*FrontInterface_ModuleMeta, string) {
    __func__ := "@lune.@base.@front.Front.loadParserToLuaCode"
    var path string
    path = parser.FP.GetStreamName(_env)
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_1204_(_env, path, mod, nil, nil)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, parser, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(_env, ast, path, ConvLua_ConvMode__ConvMeta)
    Log_log(_env, Log_Level__Trace, __func__, 670, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.LuaVM.String_format("Meta = %s", []LnsAny{metaTxt})
    }))
    
    if self.option.UpdateOnLoad{
        var saveFile func(_env *LnsEnv, suffix string,txt string,byteCompile bool,stripDebugInfo bool,checkUpdate bool)
        saveFile = func(_env *LnsEnv, suffix string,txt string,byteCompile bool,stripDebugInfo bool,checkUpdate bool) {
            var newpath string
            newpath = ""
            {
                _dir := self.option.OutputDir
                if !Lns_IsNil( _dir ) {
                    dir := _dir.(string)
                    newpath = _env.LuaVM.String_format("%s/%s%s", []LnsAny{dir, Lns_car(_env.LuaVM.String_gsub(mod,"%.", "/")).(string), suffix})
                    
                } else {
                    newpath = front_convExp3114(Lns_2DDD(_env.LuaVM.String_gsub(path,".lns$", suffix)))
                    
                }
            }
            var saveTxt string
            saveTxt = txt
            if byteCompile{
                saveTxt = front_byteCompileFromLuaTxt_1151_(_env, saveTxt, stripDebugInfo)
                
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(front_forceUpdateMeta)) &&
                _env.SetStackVal( checkUpdate) ).(bool)){
                {
                    _fileObj := front_convExp3182(Lns_2DDD(Lns_io_open(newpath, nil)))
                    if !Lns_IsNil( _fileObj ) {
                        fileObj := _fileObj.(Lns_luaStream)
                        var oldTxt LnsAny
                        oldTxt = fileObj.Read(_env, "*a")
                        fileObj.Close(_env)
                        if saveTxt == oldTxt{
                            return 
                        }
                    }
                }
            }
            {
                _fileObj := front_convExp3211(Lns_2DDD(Lns_io_open(newpath, "w")))
                if !Lns_IsNil( _fileObj ) {
                    fileObj := _fileObj.(Lns_luaStream)
                    fileObj.Write(_env, saveTxt)
                    fileObj.Close(_env)
                }
            }
        }
        saveFile(_env, ".lua", luaTxt, self.option.ByteCompile, self.option.StripDebugInfo, false)
        saveFile(_env, ".meta", metaTxt, self.option.ByteCompile, true, true)
    }
    var meta *FrontInterface_ModuleMeta
    meta = NewFrontInterface_ModuleMeta(_env, Lns_unwrap( Front_loadFromLuaTxt(_env, metaTxt)), path, front_createModuleInfo_1227_(_env, ast, mod, moduleId))
    return meta, luaTxt
}

// 716: decl @lune.@base.@front.Front.loadFileToLuaCode
func (self *front_Front) loadFileToLuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,path string,mod string)(*FrontInterface_ModuleMeta, string) {
    __func__ := "@lune.@base.@front.Front.loadFileToLuaCode"
    Log_log(_env, Log_Level__Log, __func__, 720, Log_CreateMessage(func(_env *LnsEnv) string {
        __func__ := "@lune.@base.@front.Front.loadFileToLuaCode.<anonymous>"
        return _env.LuaVM.String_format("%s: %s", []LnsAny{__func__, mod})
    }))
    
    return self.FP.loadParserToLuaCode(_env, importModuleInfo, front_createPaser_1135_(_env, path, mod), mod)
}

// 731: decl @lune.@base.@front.Front.loadFile
func (self *front_Front) loadFile(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,path string,mod string)(*FrontInterface_ModuleMeta, LnsAny) {
    __func__ := "@lune.@base.@front.Front.loadFile"
    Log_log(_env, Log_Level__Info, __func__, 735, Log_CreateMessage(func(_env *LnsEnv) string {
        __func__ := "@lune.@base.@front.Front.loadFile.<anonymous>"
        return _env.LuaVM.String_format("start %s:%s", []LnsAny{__func__, mod})
    }))
    
    var meta *FrontInterface_ModuleMeta
    var luaTxt string
    meta,luaTxt = self.FP.loadFileToLuaCode(_env, importModuleInfo, path, mod)
    {
        _preLoadInfo := self.preloadedModMap.Get(mod)
        if !Lns_IsNil( _preLoadInfo ) {
            preLoadInfo := _preLoadInfo
            return meta, preLoadInfo
        }
    }
    return meta, Lns_unwrap( Front_loadFromLuaTxt(_env, luaTxt))
}

// 748: decl @lune.@base.@front.Front.searchModule
func (self *front_Front) SearchModule(_env *LnsEnv, mod string) LnsAny {
    return self.FP.searchModuleFile(_env, mod, ".lns", nil)
}

// 752: decl @lune.@base.@front.Front.searchLuaFile
func (self *front_Front) searchLuaFile(_env *LnsEnv, moduleFullName string,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(_env, moduleFullName, ".lua", addSearchPath)
}

// 772: decl @lune.@base.@front.Front.checkUptodateMeta
func (self *front_Front) checkUptodateMeta(_env *LnsEnv, lnsPath string,metaPath string,addSearchPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.checkUptodateMeta"
    var metaObj LnsAny
    
    {
        _metaObj := self.FP.loadLua(_env, metaPath)
        if _metaObj == nil{
            Log_log(_env, Log_Level__Warn, __func__, 776, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.LuaVM.String_format("load error -- %s", []LnsAny{metaPath})
            }))
            
            return nil
        } else {
            metaObj = _metaObj
        }
    }
    var meta *Lns_luaValue
    meta = metaObj.(*Lns_luaValue)
    if meta.GetAt( "__formatVersion" ).(string) != Ver_metaVersion{
        Log_log(_env, Log_Level__Warn, __func__, 781, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.LuaVM.String_format("unmatch meta version -- %s", []LnsAny{metaPath})
        }))
        
        return nil
    }
    if meta.GetAt( "__hasTest" ).(bool){
        if meta.GetAt( "__enableTest" ).(bool) != self.option.Testing{
            Log_log(_env, Log_Level__Warn, __func__, 787, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.LuaVM.String_format("unmatch test setting -- %s", []LnsAny{metaPath})
            }))
            
            return nil
        }
    }
    {
        _exp4061 := meta.GetAt( "__dependModuleMap" ).(*Lns_luaValue)
        _key4061, _ := _exp4061.Get1stFromMap()
        for _key4061 != nil {
            moduleFullName := _key4061.(string)
            {
                _moduleLnsPath := self.FP.SearchModule(_env, moduleFullName)
                if !Lns_IsNil( _moduleLnsPath ) {
                    moduleLnsPath := _moduleLnsPath.(string)
                    {
                        _moduleLuaPath := self.FP.searchLuaFile(_env, moduleFullName, addSearchPath)
                        if !Lns_IsNil( _moduleLuaPath ) {
                            moduleLuaPath := _moduleLuaPath.(string)
                            if Lns_op_not(Util_getReadyCode(_env, moduleLnsPath, metaPath)){
                                Log_log(_env, Log_Level__Warn, __func__, 798, Log_CreateMessage(func(_env *LnsEnv) string {
                                    return _env.LuaVM.String_format("not ready -- %s, %s", []LnsAny{moduleLnsPath, metaPath})
                                }))
                                
                                return nil
                            }
                            var moduleMetaPath string
                            moduleMetaPath = front_convExp3872(Lns_2DDD(_env.LuaVM.String_gsub(moduleLuaPath,"%.lua$", ".meta")))
                            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( Depend_existFile(_env, moduleMetaPath)) &&
                                _env.SetStackVal( Lns_op_not(Util_getReadyCode(_env, moduleMetaPath, metaPath))) ).(bool)){
                                Log_log(_env, Log_Level__Warn, __func__, 806, Log_CreateMessage(func(_env *LnsEnv) string {
                                    return _env.LuaVM.String_format("not ready -- %s, %s", []LnsAny{moduleMetaPath, metaPath})
                                }))
                                
                                return nil
                            }
                        } else {
                            Log_log(_env, Log_Level__Warn, __func__, 811, Log_CreateMessage(func(_env *LnsEnv) string {
                                return _env.LuaVM.String_format("not found .lua file for -- %s", []LnsAny{moduleFullName})
                            }))
                            
                            return nil
                        }
                    }
                } else {
                    Log_log(_env, Log_Level__Warn, __func__, 816, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.LuaVM.String_format("not found .lns file -- %s", []LnsAny{moduleFullName})
                    }))
                    
                    return nil
                }
            }
            _key4061, _ = _exp4061.NextFromMap( _key4061 )
        }
    }
    return NewFrontInterface_ModuleMeta(_env, meta, lnsPath, nil)
}

// 829: decl @lune.@base.@front.Front.loadModule
func (self *front_Front) LoadModule(_env *LnsEnv, mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    __func__ := "@lune.@base.@front.Front.loadModule"
    var orgMod string
    orgMod = mod
    mod = self.gomodMap.FP.GetLuaModulePath(_env, mod)
    
    if Lns_op_not(self.FP.getLoadInfo(_env, mod)){
        {
            _luaTxt := self.convertedMap.Get(mod)
            if !Lns_IsNil( _luaTxt ) {
                luaTxt := _luaTxt.(string)
                {
                    _meta := self.moduleMgr.FP.GetMeta(_env, mod)
                    if !Lns_IsNil( _meta ) {
                        meta := _meta.(*FrontInterface_ModuleMeta)
                        self.FP.setLoadInfo(_env, mod, Newfront_LoadInfo(_env, Lns_unwrap( Front_loadFromLuaTxt(_env, luaTxt)), meta))
                    } else {
                        panic(_env.LuaVM.String_format("nothing meta -- %s", []LnsAny{mod}))
                    }
                }
            } else {
                {
                    _lnsPath := self.FP.SearchModule(_env, orgMod)
                    if !Lns_IsNil( _lnsPath ) {
                        lnsPath := _lnsPath.(string)
                        var luaPath LnsAny
                        luaPath = front_convExp4200(Lns_2DDD(_env.LuaVM.String_gsub(lnsPath, "%.lns$", ".lua")))
                        {
                            _dir := self.option.OutputDir
                            if !Lns_IsNil( _dir ) {
                                dir := _dir.(string)
                                luaPath = self.FP.searchLuaFile(_env, orgMod, dir)
                                
                            }
                        }
                        var loadVal LnsAny
                        loadVal = nil
                        if luaPath != nil{
                            luaPath_583 := luaPath.(string)
                            if Util_getReadyCode(_env, lnsPath, luaPath_583){
                                var metaPath string
                                metaPath = front_convExp4258(Lns_2DDD(_env.LuaVM.String_gsub(luaPath_583, "%.lua$", ".meta")))
                                if Util_getReadyCode(_env, lnsPath, metaPath){
                                    {
                                        _preLoadInfo := self.preloadedModMap.Get(mod)
                                        if !Lns_IsNil( _preLoadInfo ) {
                                            preLoadInfo := _preLoadInfo
                                            loadVal = preLoadInfo
                                            
                                        } else {
                                            loadVal = self.FP.loadLua(_env, luaPath_583)
                                            
                                        }
                                    }
                                    {
                                        __exp := loadVal
                                        if !Lns_IsNil( __exp ) {
                                            _exp := __exp
                                            {
                                                _meta := self.FP.checkUptodateMeta(_env, lnsPath, metaPath, self.option.OutputDir)
                                                if !Lns_IsNil( _meta ) {
                                                    meta := _meta.(*FrontInterface_ModuleMeta)
                                                    self.FP.setLoadInfo(_env, mod, Newfront_LoadInfo(_env, _exp, meta))
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
                            var meta *FrontInterface_ModuleMeta
                            var workVal LnsAny
                            meta,workVal = self.FP.loadFile(_env, NewFrontInterface_ImportModuleInfo(_env), lnsPath, mod)
                            self.FP.setLoadInfo(_env, mod, Newfront_LoadInfo(_env, workVal, meta))
                        }
                    } else {
                        if self.bindModuleSet.Has(mod){
                            Log_log(_env, Log_Level__Warn, __func__, 883, Log_CreateMessage(func(_env *LnsEnv) string {
                                return _env.LuaVM.String_format("load from the binding -- %s", []LnsAny{mod})
                            }))
                            
                            var workMod LnsAny
                            workMod = Lns_require(mod)
                            var meta *FrontInterface_ModuleMeta
                            meta = NewFrontInterface_ModuleMeta(_env, Lns_unwrap( Front_loadFromLuaTxt(_env, "return {}")), Lns_car(_env.LuaVM.String_gsub(mod,"%.", "/")).(string) + ".lns", nil)
                            self.FP.setLoadInfo(_env, mod, Newfront_LoadInfo(_env, workMod, meta))
                        }
                    }
                }
            }
        }
    }
    {
        __exp := self.FP.getLoadInfo(_env, mod)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*front_LoadInfo)
            return _exp.Mod, _exp.Meta
        }
    }
    panic(_env.LuaVM.String_format("load error, %s", []LnsAny{mod}))
// insert a dummy
    return nil,nil
}

// 904: decl @lune.@base.@front.Front.getLuaModulePath
func (self *front_Front) GetLuaModulePath(_env *LnsEnv, mod string) string {
    return self.gomodMap.FP.GetLuaModulePath(_env, mod)
}

// 908: decl @lune.@base.@front.Front.loadMeta
func (self *front_Front) LoadMeta(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    __func__ := "@lune.@base.@front.Front.loadMeta"
    var orgMod string
    orgMod = mod
    mod = self.gomodMap.FP.GetLuaModulePath(_env, mod)
    
    if Lns_op_not(self.moduleMgr.FP.GetMeta(_env, mod)){
        {
            __exp := self.FP.getLoadInfo(_env, mod)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*front_LoadInfo)
                self.moduleMgr.FP.AddMeta(_env, mod, _exp.Meta)
            } else {
                {
                    _lnsPath := self.FP.SearchModule(_env, orgMod)
                    if !Lns_IsNil( _lnsPath ) {
                        lnsPath := _lnsPath.(string)
                        var meta LnsAny
                        meta = nil
                        if Lns_op_not(self.targetSet.Has(mod)){
                            var luaPath LnsAny
                            luaPath = front_convExp4676(Lns_2DDD(_env.LuaVM.String_gsub(lnsPath, "%.lns$", ".lua")))
                            {
                                _dir := self.option.OutputDir
                                if !Lns_IsNil( _dir ) {
                                    dir := _dir.(string)
                                    luaPath = self.FP.searchLuaFile(_env, orgMod, dir)
                                    
                                }
                            }
                            if luaPath != nil{
                                luaPath_630 := luaPath.(string)
                                var forceFlag bool
                                switch _exp4746 := self.option.TransCtrlInfo.UptodateMode.(type) {
                                case *Types_CheckingUptodateMode__ForceAll:
                                    forceFlag = true
                                    
                                case *Types_CheckingUptodateMode__Force1:
                                forceMod := _exp4746.Val1
                                    forceFlag = mod == forceMod
                                    
                                case *Types_CheckingUptodateMode__Normal:
                                    forceFlag = false
                                    
                                case *Types_CheckingUptodateMode__Touch:
                                    forceFlag = false
                                    
                                }
                                if Lns_op_not(forceFlag){
                                    if Util_getReadyCode(_env, lnsPath, luaPath_630){
                                        var metaPath string
                                        metaPath = front_convExp4774(Lns_2DDD(_env.LuaVM.String_gsub(luaPath_630, "%.lua$", ".meta")))
                                        if Util_getReadyCode(_env, lnsPath, metaPath){
                                            meta = self.FP.checkUptodateMeta(_env, lnsPath, metaPath, self.option.OutputDir)
                                            
                                        } else { 
                                            Log_log(_env, Log_Level__Warn, __func__, 954, Log_CreateMessage(func(_env *LnsEnv) string {
                                                return _env.LuaVM.String_format("%s not ready meta %s, %s", []LnsAny{orgMod, lnsPath, metaPath})
                                            }))
                                            
                                        }
                                    } else { 
                                        Log_log(_env, Log_Level__Warn, __func__, 958, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.LuaVM.String_format("%s not ready lua %s, %s", []LnsAny{orgMod, lnsPath, luaPath_630})
                                        }))
                                        
                                    }
                                } else { 
                                    Log_log(_env, Log_Level__Warn, __func__, 962, Log_CreateMessage(func(_env *LnsEnv) string {
                                        return _env.LuaVM.String_format("force analyze -- %s", []LnsAny{orgMod})
                                    }))
                                    
                                }
                            } else {
                                Log_log(_env, Log_Level__Warn, __func__, 966, Log_CreateMessage(func(_env *LnsEnv) string {
                                    return _env.LuaVM.String_format("%s not found lua in %s", []LnsAny{orgMod, self.option.OutputDir})
                                }))
                                
                            }
                        }
                        if meta != nil{
                            meta_654 := meta.(*FrontInterface_ModuleMeta)
                            self.moduleMgr.FP.AddMeta(_env, mod, meta_654)
                        } else {
                            var metawork *FrontInterface_ModuleMeta
                            var luaTxt string
                            metawork,luaTxt = self.FP.loadFileToLuaCode(_env, importModuleInfo, lnsPath, orgMod)
                            self.moduleMgr.FP.AddMeta(_env, mod, metawork)
                            self.convertedMap.Set(mod,luaTxt)
                        }
                    } else {
                        {
                            _lnsCode := Depend_getBindLns(_env, mod)
                            if !Lns_IsNil( _lnsCode ) {
                                lnsCode := _lnsCode.(string)
                                var path string
                                path = Lns_car(_env.LuaVM.String_gsub(mod,"%.", "/")).(string) + ".lns"
                                var parser *Parser_StreamParser
                                parser = NewParser_StreamParser(_env, NewParser_TxtStream(_env, lnsCode).FP, path, false, nil)
                                var meta *FrontInterface_ModuleMeta
                                var luaTxt string
                                meta,luaTxt = self.FP.loadParserToLuaCode(_env, importModuleInfo, &parser.Parser_Parser, mod)
                                self.moduleMgr.FP.AddMeta(_env, mod, meta)
                                self.convertedMap.Set(mod,luaTxt)
                            }
                        }
                    }
                }
            }
        }
    }
    {
        _meta := self.moduleMgr.FP.GetMeta(_env, mod)
        if !Lns_IsNil( _meta ) {
            meta := _meta.(*FrontInterface_ModuleMeta)
            return meta
        }
    }
    return nil
}

// 1000: decl @lune.@base.@front.Front.dumpTokenize
func (self *front_Front) DumpTokenize(_env *LnsEnv, scriptPath string) {
    var parser *Parser_Parser
    parser = self.FP.createPaser(_env, scriptPath)
    for  {
        var token *Types_Token
        
        {
            _token := parser.FP.GetToken(_env)
            if _token == nil{
                break
            } else {
                token = _token.(*Types_Token)
            }
        }
        Lns_print([]LnsAny{Types_TokenKind_getTxt( token.Kind), token.Pos.LineNo, token.Pos.Column, token.Txt})
    }
}

// 1012: decl @lune.@base.@front.Front.dumpAst
func (self *front_Front) DumpAst(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    Depend_profile(_env, self.option.ValidProf, conv2Form5369(func(_env *LnsEnv) {
        var ast *TransUnit_ASTInfo
        ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), self.FP.createPaser(_env, scriptPath), mod, front_getModuleId_1204_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, DumpNode_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt(_env, "", 0)))
    }), scriptPath + ".profi")
}

// 1031: decl @lune.@base.@front.Front.format
func (self *front_Front) Format(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), self.FP.createPaser(_env, scriptPath), mod, front_getModuleId_1204_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, Formatter_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), Lns_io_stdout), Formatter_Opt2Stem(NewFormatter_Opt(_env, ast.FP.Get_node(_env))))
}

// 1046: decl @lune.@base.@front.Front.checkDiag
func (self *front_Front) CheckDiag(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    Util_setErrorCode(_env, 0)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), self.FP.createPaser(_env, scriptPath), mod, front_getModuleId_1204_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Diag, nil)
}

// 1056: decl @lune.@base.@front.Front.complete
func (self *front_Front) Complete(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), self.FP.createPaser(_env, scriptPath), mod, front_getModuleId_1204_(_env, scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Complete, self.option.AnalyzePos)
}

// 1064: decl @lune.@base.@front.Front.inquire
func (self *front_Front) Inquire(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), self.FP.createPaser(_env, scriptPath), mod, front_getModuleId_1204_(_env, scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Inquire, self.option.AnalyzePos)
}

// 1073: decl @lune.@base.@front.Front.createGlue
func (self *front_Front) CreateGlue(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), self.FP.createPaser(_env, scriptPath), mod, front_getModuleId_1204_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    var filter *Nodes_Filter
    filter = GlueFilter_createFilter(_env, self.option.OutputDir)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, filter, 0)
}

// 1107: decl @lune.@base.@front.Front.convertLuaToStreamFromScript
func (self *front_Front) ConvertLuaToStreamFromScript(_env *LnsEnv, parser LnsAny,moduleId *FrontInterface_ModuleId,uptodate LnsAny,convMode LnsInt,path string,mod string,byteCompile bool,stripDebugInfo bool,openOStream front_OpenOStreamForConvert_1308_,closeOStream LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.convertLuaToStreamFromScript"
    Log_log(_env, Log_Level__Log, __func__, 1113, Log_CreateMessage(func(_env *LnsEnv) string {
        return path
    }))
    
    var outputDependInfo func(_env *LnsEnv, stream LnsAny,metaInfo LnsAny)
    outputDependInfo = func(_env *LnsEnv, stream LnsAny,metaInfo LnsAny) {
        if stream != nil{
            stream_747 := stream.(Lns_oStream)
            if metaInfo != nil{
                metaInfo_749 := metaInfo.(*front_MetaForBuildId)
                var dependInfo *OutputDepend_DependInfo
                dependInfo = NewOutputDepend_DependInfo(_env, mod)
                for _dependMod, _ := range( metaInfo_749.G__dependModuleMap.Items ) {
                    dependMod := _dependMod.(string)
                    dependInfo.FP.AddImpotModule(_env, dependMod)
                }
                for _, _subMod := range( metaInfo_749.G__subModuleMap.Items ) {
                    subMod := _subMod.(string)
                    dependInfo.FP.AddSubMod(_env, subMod)
                }
                dependInfo.FP.Output(_env, stream_747)
            } else {
                Util_err(_env, "metaInfo is nil")
            }
        }
    }
    var retAst LnsAny
    retAst = nil
    switch _exp6351 := uptodate.(type) {
    case *front_ModuleUptodate__Uptodate:
    metaInfo := _exp6351.Val1
        Util_errorLog(_env, "uptodate -- " + path)
        var stream LnsAny
        var metaStream LnsAny
        var dependsStream LnsAny
        stream,metaStream,dependsStream = openOStream(_env, uptodate)
        outputDependInfo(_env, dependsStream, metaInfo)
        if closeOStream != nil{
            closeOStream_764 := closeOStream.(front_CloseOStreamForConvert_1310_)
            closeOStream_764(_env, stream, metaStream, dependsStream)
        }
    case *front_ModuleUptodate__NeedUpdate:
        var ast *TransUnit_ASTInfo
        ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), front_createPaser_1135_(_env, path, mod), mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
        retAst = ast
        
        var streamDst LnsAny
        var metaStreamDst LnsAny
        var dependsStreamDst LnsAny
        streamDst,metaStreamDst,dependsStreamDst = openOStream(_env, uptodate)
        if streamDst != nil && metaStreamDst != nil{
            streamDst_771 := streamDst.(Lns_oStream)
            metaStreamDst_772 := metaStreamDst.(Lns_oStream)
            var streamMem *Util_memStream
            streamMem = NewUtil_memStream(_env)
            var metaStreamMem *Util_memStream
            metaStreamMem = NewUtil_memStream(_env)
            var dependsStreamMem *Util_memStream
            dependsStreamMem = NewUtil_memStream(_env)
            var stream Lns_oStream
            var metaStream Lns_oStream
            var dependsStream LnsAny
            if Str_isValidStrBuilder(_env){
                stream = streamMem.FP
                
                metaStream = metaStreamMem.FP
                
                dependsStream = dependsStreamMem.FP
                
            } else { 
                stream = streamDst_771
                
                metaStream = metaStreamDst_772
                
                dependsStream = dependsStreamDst
                
            }
            var outStream Lns_oStream
            var oMetaStream Lns_oStream
            outStream,oMetaStream = stream, metaStream
            if dependsStream != nil{
                dependsStream_784 := dependsStream.(Lns_oStream)
                ast.FP.Get_node(_env).FP.ProcessFilter(_env, OutputDepend_createFilter(_env, dependsStream_784), 1)
            }
            var byteStream *Util_memStream
            byteStream = NewUtil_memStream(_env)
            var byteMetaStream *Util_memStream
            byteMetaStream = NewUtil_memStream(_env)
            if byteCompile{
                outStream = byteStream.FP
                
                oMetaStream = byteMetaStream.FP
                
            }
            self.FP.convert(_env, ast, path, outStream, oMetaStream, convMode, false)
            if byteCompile{
                stream.Write(_env, front_byteCompileFromLuaTxt_1151_(_env, byteStream.FP.Get_txt(_env), stripDebugInfo))
                if metaStream != stream{
                    metaStream.Write(_env, front_byteCompileFromLuaTxt_1151_(_env, byteMetaStream.FP.Get_txt(_env), true))
                }
            }
            if Str_isValidStrBuilder(_env){
                streamDst_771.Write(_env, streamMem.FP.Get_txt(_env))
                metaStreamDst_772.Write(_env, metaStreamMem.FP.Get_txt(_env))
                Lns_NilAccFinCall2( _env.NilAccFin(_env.NilAccPush(dependsStreamDst) && 
                Lns_NilAccCall2( _env, func () (LnsAny,LnsAny) { return _env.NilAccPop().(Lns_oStream).Write(_env, dependsStreamMem.FP.Get_txt(_env))})/* 1208:16 */))
            }
        } else {
            Util_err(_env, "failed to open lua stream or meta stream")
        }
        if closeOStream != nil{
            closeOStream_793 := closeOStream.(front_CloseOStreamForConvert_1310_)
            closeOStream_793(_env, streamDst, metaStreamDst, dependsStreamDst)
        }
    case *front_ModuleUptodate__NeedTouch:
    metaCode := _exp6351.Val1
    metaInfo := _exp6351.Val2
        Util_errorLog(_env, "touch -- " + path)
        var stream LnsAny
        var metaStream LnsAny
        var dependsStream LnsAny
        stream,metaStream,dependsStream = openOStream(_env, uptodate)
        if self.option.Mode == Option_ModeKind__SaveMeta{
            if metaStream != nil{
                metaStream_802 := metaStream.(Lns_oStream)
                metaStream_802.Write(_env, metaCode)
            } else {
                Util_err(_env, "failed to open meta stream")
            }
        }
        outputDependInfo(_env, dependsStream, metaInfo)
        if closeOStream != nil{
            closeOStream_805 := closeOStream.(front_CloseOStreamForConvert_1310_)
            closeOStream_805(_env, stream, metaStream, dependsStream)
        }
    }
    return retAst
}

// 1242: decl @lune.@base.@front.Front.getGoAppName
func (self *front_Front) getGoAppName(_env *LnsEnv) string {
    var appName string
    
    {
        _appName := self.option.AppName
        if _appName == nil{
            appName = self.gomodMap.FP.Get_name(_env)
            
        } else {
            appName = _appName.(string)
        }
    }
    return appName
}

// 1249: decl @lune.@base.@front.Front.createGoOption
func (self *front_Front) createGoOption(_env *LnsEnv, scriptPath string) *ConvGo_Option {
    var packageName string
    {
        __exp := self.option.PackageName
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(string)
            packageName = _exp
            
        } else {
            if Lns_op_not(Lns_car(_env.LuaVM.String_find(scriptPath,"/", nil, nil))){
                packageName = "main"
                
            } else { 
                var parentPath string
                parentPath = front_convExp6448(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(scriptPath,"/[^/]+$", "")).(string),".*/", "")))
                if len(parentPath) == 0{
                    packageName = "main"
                    
                } else if parentPath == "."{
                    packageName = "main"
                    
                } else if parentPath == ".."{
                    packageName = "main"
                    
                } else { 
                    packageName = front_convExp6505(Lns_2DDD(_env.LuaVM.String_gsub(parentPath,"[^%w]", "")))
                    
                }
            }
        }
    }
    return NewConvGo_Option(_env, packageName, self.FP.getGoAppName(_env), self.option.MainModule, self.option.FP.Get_addEnvArg(_env))
}

// 1277: decl @lune.@base.@front.Front.convertToLua
func (self *front_Front) convertToLua(_env *LnsEnv, scriptPath string,convMode LnsInt,streamLua Lns_oStream,streamMeta Lns_oStream) LnsAny {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    var parser *Parser_Parser
    parser = front_createPaser_1135_(_env, scriptPath, mod)
    var ast LnsAny
    ast = self.FP.ConvertLuaToStreamFromScript(_env, parser, FrontInterface_ModuleId__tempId, front_ModuleUptodate__NeedUpdate_Obj, convMode, scriptPath, mod, self.option.ByteCompile, self.option.StripDebugInfo, front_OpenOStreamForConvert_1308_(func(_env *LnsEnv, mode LnsAny)(LnsAny, LnsAny, LnsAny) {
        return streamLua, streamMeta, self.option.FP.OpenDepend(_env, nil)
    }), front_CloseOStreamForConvert_1310_(Front_convertToLua___anonymous_1328_))
    if ast != nil{
        ast_851 := ast.(*TransUnit_ASTInfo)
        if _switch6713 := self.option.ConvTo; _switch6713 == Types_Lang__Go {
            var conv *Nodes_Filter
            conv = ConvGo_createFilter(_env, self.option.Testing, "stdout", streamLua, ast_851, self.FP.createGoOption(_env, scriptPath))
            ast_851.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast_851.FP.Get_node(_env))))
        }
    }
    return ast
}

// 1317: decl @lune.@base.@front.Front.saveToGo
func (self *front_Front) SaveToGo(_env *LnsEnv, scriptPath string,ast *TransUnit_ASTInfo) {
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
        switch _exp6768 := pragma.(type) {
        case *LuneControl_Pragma__limit_conv_code:
        codeSet := _exp6768.Val1
            if Lns_op_not(codeSet.Has(LuneControl_Code__Go)){
                return 
            }
        }
    }
    var path string
    path = front_convExp6783(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".go")))
    {
        _dir := self.option.OutputDir
        if !Lns_IsNil( _dir ) {
            dir := _dir.(string)
            path = _env.LuaVM.String_format("%s/%s", []LnsAny{dir, path})
            
        }
    }
    var file Lns_luaStream
    
    {
        _file := front_convExp6822(Lns_2DDD(Lns_io_open(path, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var memStream *Util_memStream
    memStream = NewUtil_memStream(_env)
    var dstStream Lns_oStream
    if Str_isValidStrBuilder(_env){
        dstStream = memStream.FP
        
    } else { 
        dstStream = file
        
    }
    var conv *Nodes_Filter
    conv = ConvGo_createFilter(_env, self.option.Testing, path, dstStream, ast, self.FP.createGoOption(_env, scriptPath))
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast.FP.Get_node(_env))))
    if Str_isValidStrBuilder(_env){
        file.Write(_env, memStream.FP.Get_txt(_env))
    }
    file.Close(_env)
}

// 1362: decl @lune.@base.@front.Front.saveToC
func (self *front_Front) SaveToC(_env *LnsEnv, scriptPath string,ast *TransUnit_ASTInfo) {
    var cPath string
    cPath = front_convExp6949(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".c")))
    var file Lns_luaStream
    
    {
        _file := front_convExp6964(Lns_2DDD(Lns_io_open(cPath, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var hPath string
    hPath = front_convExp6978(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".h")))
    var hFile Lns_luaStream
    
    {
        _hFile := front_convExp6994(Lns_2DDD(Lns_io_open(hPath, "w")))
        if _hFile == nil{
            return 
        } else {
            hFile = _hFile.(Lns_luaStream)
        }
    }
    file.Close(_env)
    hFile.Close(_env)
}

// 1382: decl @lune.@base.@front.Front.outputBuiltin
func (self *front_Front) OutputBuiltin(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, "lns_builtin")
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &NewParser_DummyParser(_env).Parser_Parser, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    self.FP.SaveToC(_env, scriptPath, ast)
}

// 1404: decl @lune.@base.@front.Front.saveToLua
func (self *front_Front) SaveToLua(_env *LnsEnv, updateInfo *front_UpdateInfo)(bool, LnsAny, *Util_memStream) {
    var scriptPath string
    scriptPath = updateInfo.FP.Get_scriptPath(_env)
    var dependsPath LnsAny
    dependsPath = updateInfo.FP.Get_dependsPath(_env)
    var parser LnsAny
    parser = updateInfo.FP.Get_parser(_env)
    var moduleId *FrontInterface_ModuleId
    moduleId = updateInfo.FP.Get_moduleId(_env)
    var uptodate LnsAny
    uptodate = updateInfo.FP.Get_uptodate(_env)
    var checkDiff func(_env *LnsEnv, oldStream *Parser_TxtStream,newStream *Parser_TxtStream)(bool, string)
    checkDiff = func(_env *LnsEnv, oldStream *Parser_TxtStream,newStream *Parser_TxtStream)(bool, string) {
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
            newLine = newStream.FP.Read(_env, "*l")
            var oldLine LnsAny
            oldLine = oldStream.FP.Read(_env, "*l")
            if oldLine != nil{
                oldLine_945 := oldLine.(string)
                if len(oldBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(oldLine_945,"^_moduleObj.__buildId", nil, nil))){
                        oldBuildIdLine = oldLine_945
                        
                    }
                }
            }
            
            if newLine != nil{
                newLine_949 := newLine.(string)
                if len(newBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(newLine_949,"^_moduleObj.__buildId", nil, nil))){
                        newBuildIdLine = newLine_949
                        
                    }
                }
            }
            
            if newLine != oldLine{
                var cont bool
                cont = false
                if newLine != nil && oldLine != nil{
                    newLine_955 := newLine.(string)
                    oldLine_956 := oldLine.(string)
                    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(oldLine_956,"^_moduleObj.__buildId", nil, nil))){
                        if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(newLine_955,"^_moduleObj.__buildId", nil, nil))){
                            tailBeginPos = newStream.FP.Get_lineNo(_env)
                            
                            cont = true
                            
                        }
                    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_car(_env.LuaVM.String_find(oldLine_956,"^__dependModuleMap.*buildId =", nil, nil))) &&
                        _env.SetStackVal( Lns_car(_env.LuaVM.String_find(newLine_955,"^__dependModuleMap.*buildId =", nil, nil))) )){
                        var oldSub string
                        oldSub = front_convExp7404(Lns_2DDD(_env.LuaVM.String_gsub(oldLine_956,"buildId =.*", "")))
                        var newSub string
                        newSub = front_convExp7417(Lns_2DDD(_env.LuaVM.String_gsub(newLine_955,"buildId =.*", "")))
                        if oldSub == newSub{
                            cont = true
                            
                        }
                    }
                }
                if Lns_op_not(cont){
                    Log_log(_env, Log_Level__Debug, __func__, 1481, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.LuaVM.String_format("<%s>, <%s>", []LnsAny{oldLine, newLine})
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
                    oldBuildId = Front_saveToLua__txt2ModuleId_1357_(_env, oldBuildIdLine)
                    var newBuildId *FrontInterface_ModuleId
                    newBuildId = Front_saveToLua__txt2ModuleId_1357_(_env, newBuildIdLine)
                    var worlBuildId *FrontInterface_ModuleId
                    worlBuildId = FrontInterface_ModuleId_createId(_env, newBuildId.FP.Get_modTime(_env), oldBuildId.FP.Get_buildCount(_env))
                    var buildIdLine string
                    buildIdLine = _env.LuaVM.String_format("_moduleObj.__buildId = %q", []LnsAny{worlBuildId.FP.Get_idStr(_env)})
                    var txt string
                    txt = _env.LuaVM.String_format("%s%s\n%s", []LnsAny{newStream.FP.GetSubstring(_env, 1, headEndPos), buildIdLine, newStream.FP.GetSubstring(_env, tailBeginPos, nil)})
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
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    var luaPath string
    luaPath = front_convExp7655(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".lua")))
    var metaPath string
    metaPath = front_convExp7668(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".meta")))
    if Lns_isCondTrue( self.option.OutputDir){
        var filename string
        filename = front_convExp7685(Lns_2DDD(_env.LuaVM.String_gsub(mod,"%.", "/")))
        luaPath = _env.LuaVM.String_format("%s/%s.lua", []LnsAny{self.option.OutputDir, filename})
        
        metaPath = _env.LuaVM.String_format("%s/%s.meta", []LnsAny{self.option.OutputDir, filename})
        
    }
    var metaMemStream *Util_memStream
    metaMemStream = NewUtil_memStream(_env)
    if luaPath == scriptPath{
        Util_errorLog(_env, _env.LuaVM.String_format("%s is illegal filename.", []LnsAny{luaPath}))
    } else { 
        var convMode LnsInt
        convMode = ConvLua_ConvMode__ConvMeta
        var metaFileObj LnsAny
        metaFileObj = nil
        var tempMetaPath string
        tempMetaPath = metaPath + ".tmp"
        ast = self.FP.ConvertLuaToStreamFromScript(_env, parser, moduleId, uptodate, convMode, scriptPath, mod, self.option.ByteCompile, self.option.StripDebugInfo, front_OpenOStreamForConvert_1308_(func(_env *LnsEnv, mode LnsAny)(LnsAny, LnsAny, LnsAny) {
            var openLuaStream func(_env *LnsEnv) LnsAny
            openLuaStream = func(_env *LnsEnv) LnsAny {
                var fileObj Lns_luaStream
                
                {
                    _fileObj := front_convExp7835(Lns_2DDD(Lns_io_open(luaPath, "w")))
                    if _fileObj == nil{
                        panic(_env.LuaVM.String_format("write open error -- %s", []LnsAny{luaPath}))
                    } else {
                        fileObj = _fileObj.(Lns_luaStream)
                    }
                }
                return fileObj
            }
            var openStreams func(_env *LnsEnv, luaFlag bool)(LnsAny, LnsAny)
            openStreams = func(_env *LnsEnv, luaFlag bool)(LnsAny, LnsAny) {
                var stream LnsAny
                stream = nil
                if luaFlag{
                    stream = openLuaStream(_env)
                    
                }
                if self.option.Mode == Option_ModeKind__SaveMeta{
                    {
                        __exp := front_convExp7911(Lns_2DDD(Lns_io_open(tempMetaPath, "w+")))
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(Lns_luaStream)
                            metaFileObj = _exp
                            
                        } else {
                            panic(_env.LuaVM.String_format("write open error -- %s", []LnsAny{metaPath}))
                        }
                    }
                }
                return stream, metaMemStream.FP
            }
            var stream LnsAny
            stream = nil
            var metaStream LnsAny
            metaStream = stream
            switch mode.(type) {
            case *front_ModuleUptodate__Uptodate:
            case *front_ModuleUptodate__NeedTouch:
                stream, metaStream = openStreams(_env, false)
                
            default:
                stream, metaStream = openStreams(_env, true)
                
            }
            return stream, metaStream, self.option.FP.OpenDepend(_env, dependsPath)
        }), front_CloseOStreamForConvert_1310_(func(_env *LnsEnv, stream LnsAny,metaStream LnsAny,dependStream LnsAny) {
            if stream != nil{
                stream_1021 := stream.(Lns_oStream)
                stream_1021.Close(_env)
            }
            if dependStream != nil{
                dependStream_1023 := dependStream.(Lns_oStream)
                dependStream_1023.Close(_env)
            }
            if metaFileObj != nil{
                metaFileObj_1025 := metaFileObj.(Lns_luaStream)
                var newMetaTxt string
                newMetaTxt = metaMemStream.FP.Get_txt(_env)
                metaFileObj_1025.Write(_env, newMetaTxt)
                metaFileObj_1025.Close(_env)
                var oldMetaTxt string
                oldMetaTxt = ""
                {
                    _oldFileObj := front_convExp8083(Lns_2DDD(Lns_io_open(metaPath, nil)))
                    if !Lns_IsNil( _oldFileObj ) {
                        oldFileObj := _oldFileObj.(Lns_luaStream)
                        oldMetaTxt = Lns_unwrapDefault( oldFileObj.Read(_env, "*a"), "").(string)
                        
                        oldFileObj.Close(_env)
                    }
                }
                var sameFlag bool
                var txt string
                sameFlag,txt = checkDiff(_env, NewParser_TxtStream(_env, oldMetaTxt), NewParser_TxtStream(_env, newMetaTxt))
                if Lns_op_not(sameFlag){
                    _env.LuaVM.OS_rename(tempMetaPath, metaPath)
                } else { 
                    _env.LuaVM.OS_remove(tempMetaPath)
                    if txt != ""{
                        {
                            _fileObj := front_convExp8164(Lns_2DDD(Lns_io_open(metaPath, "w")))
                            if !Lns_IsNil( _fileObj ) {
                                fileObj := _fileObj.(Lns_luaStream)
                                fileObj.Write(_env, txt)
                                fileObj.Close(_env)
                            }
                        }
                    } else { 
                        updateFlag = false
                        
                    }
                }
            }
        }))
        
    }
    if ast != nil{
        ast_1039 := ast.(*TransUnit_ASTInfo)
        if _switch8230 := self.option.ConvTo; _switch8230 == Types_Lang__C {
            self.FP.SaveToC(_env, scriptPath, ast_1039)
        } else if _switch8230 == Types_Lang__Go {
            self.FP.SaveToGo(_env, scriptPath, ast_1039)
        }
    }
    return updateFlag, ast, metaMemStream
}

// 1644: decl @lune.@base.@front.Front.outputBootC
func (self *front_Front) outputBootC(_env *LnsEnv, scriptPath string) {
}

// 1681: decl @lune.@base.@front.Front.build
func (self *front_Front) Build(_env *LnsEnv, buildMode LnsAny,astCallback LnsAny) {
    var createUpdateInfo func(_env *LnsEnv, scriptPath string,dependsPath LnsAny) *front_UpdateInfo
    createUpdateInfo = func(_env *LnsEnv, scriptPath string,dependsPath LnsAny) *front_UpdateInfo {
        var mod string
        mod = self.FP.scriptPath2Module(_env, scriptPath)
        var moduleId *FrontInterface_ModuleId
        var uptodate LnsAny
        moduleId,uptodate = self.FP.getModuleIdAndCheckUptodate(_env, scriptPath, mod)
        var parser LnsAny
        switch uptodate.(type) {
        case *front_ModuleUptodate__NeedUpdate:
            parser = front_createPaser_1135_(_env, scriptPath, mod)
            
        case *front_ModuleUptodate__NeedTouch:
            parser = nil
            
        case *front_ModuleUptodate__Uptodate:
            parser = nil
            
        }
        return Newfront_UpdateInfo(_env, scriptPath, dependsPath, parser, moduleId, uptodate)
    }
    var process func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo)
    process = func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo) {
        var mod string
        mod = self.FP.scriptPath2Module(_env, updateInfo.FP.Get_scriptPath(_env))
        var createMeta func(_env *LnsEnv, ast LnsAny,metaStream *Util_memStream)
        createMeta = func(_env *LnsEnv, ast LnsAny,metaStream *Util_memStream) {
            if Lns_op_not(oneShot){
                if ast != nil{
                    ast_1093 := ast.(*TransUnit_ASTInfo)
                    var metaInfo *FrontInterface_ModuleMeta
                    metaInfo = NewFrontInterface_ModuleMeta(_env, Lns_unwrap( Front_loadFromLuaTxt(_env, metaStream.FP.Get_txt(_env))), updateInfo.FP.Get_scriptPath(_env), front_createModuleInfo_1227_(_env, ast_1093, mod, updateInfo.FP.Get_moduleId(_env)))
                    self.moduleMgr.FP.AddMeta(_env, mod, metaInfo)
                }
            }
        }
        switch _exp8612 := buildMode.(type) {
        case *front_BuildMode__Save:
            var ast LnsAny
            var streamMeta *Util_memStream
            _,ast,streamMeta = self.FP.SaveToLua(_env, updateInfo)
            createMeta(_env, ast, streamMeta)
        case *front_BuildMode__Output:
        streamLua := _exp8612.Val1
        streamMeta := _exp8612.Val2
            self.FP.convertToLua(_env, updateInfo.FP.Get_scriptPath(_env), ConvLua_ConvMode__ConvMeta, streamLua, streamMeta)
        case *front_BuildMode__CreateAst:
            if Lns_isCondTrue( self.moduleMgr.FP.GetAst(_env, mod)){
                return 
            }
            var streamMeta *Util_memStream
            streamMeta = NewUtil_memStream(_env)
            var ast LnsAny
            ast = self.FP.convertToLua(_env, updateInfo.FP.Get_scriptPath(_env), ConvLua_ConvMode__ConvMeta, NewUtil_NullOStream(_env).FP, streamMeta.FP)
            createMeta(_env, ast, streamMeta)
        }
    }
    Depend_profile(_env, self.option.ValidProf, conv2Form8799(func(_env *LnsEnv) {
        if self.option.ScriptPath == "@-"{
            for _, _path := range( self.option.BatchList.Items ) {
                path := _path.(string)
                self.targetSet.Add(self.FP.scriptPath2Module(_env, path))
            }
            for _, _path := range( self.option.BatchList.Items ) {
                path := _path.(string)
                var updateInfo *front_UpdateInfo
                updateInfo = createUpdateInfo(_env, path, Lns_car(_env.LuaVM.String_gsub(path,".lns$", ".d")).(string))
                var prev LnsReal
                prev = _env.LuaVM.OS_clock()
                Lns_print([]LnsAny{_env.LuaVM.String_format("%s: processing...", []LnsAny{updateInfo.FP.Get_scriptPath(_env)})})
                process(_env, false, updateInfo)
                Lns_print([]LnsAny{_env.LuaVM.String_format("%s: done %g", []LnsAny{updateInfo.FP.Get_scriptPath(_env), _env.LuaVM.OS_clock() - prev})})
            }
        } else { 
            process(_env, true, createUpdateInfo(_env, self.option.ScriptPath, nil))
        }
        if astCallback != nil{
            astCallback_1117 := astCallback.(Front_AstCallback)
            for _, _key := range( self.moduleMgr.FP.GetModList(_env).Items ) {
                key := _key.(string)
                {
                    __exp := self.moduleMgr.FP.GetAst(_env, key)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*TransUnit_ASTInfo)
                        astCallback_1117(_env, _exp)
                    }
                }
            }
        }
    }), self.option.ScriptPath + ".profi")
}

// 1771: decl @lune.@base.@front.Front.exec
func (self *front_Front) Exec(_env *LnsEnv) {
    __func__ := "@lune.@base.@front.Front.exec"
    Log_log(_env, Log_Level__Trace, __func__, 1773, Log_CreateMessage(func(_env *LnsEnv) string {
        return Option_ModeKind_getTxt( self.option.Mode)
    }))
    
    if _switch9283 := self.option.Mode; _switch9283 == Option_ModeKind__Token {
        self.FP.DumpTokenize(_env, self.option.ScriptPath)
    } else if _switch9283 == Option_ModeKind__Ast {
        self.FP.DumpAst(_env, self.option.ScriptPath)
    } else if _switch9283 == Option_ModeKind__Format {
        self.FP.Format(_env, self.option.ScriptPath)
    } else if _switch9283 == Option_ModeKind__Diag {
        self.FP.CheckDiag(_env, self.option.ScriptPath)
    } else if _switch9283 == Option_ModeKind__Complete {
        self.FP.Complete(_env, self.option.ScriptPath)
    } else if _switch9283 == Option_ModeKind__Inquire {
        self.FP.Inquire(_env, self.option.ScriptPath)
    } else if _switch9283 == Option_ModeKind__Glue {
        self.FP.CreateGlue(_env, self.option.ScriptPath)
    } else if _switch9283 == Option_ModeKind__Lua || _switch9283 == Option_ModeKind__LuaMeta {
        var convMode LnsInt
        if self.option.Mode == Option_ModeKind__Lua{
            convMode = ConvLua_ConvMode__Convert
            
        } else { 
            convMode = ConvLua_ConvMode__ConvMeta
            
        }
        self.FP.convertToLua(_env, self.option.ScriptPath, convMode, Lns_io_stdout, Lns_io_stdout)
    } else if _switch9283 == Option_ModeKind__Save || _switch9283 == Option_ModeKind__SaveMeta {
        self.FP.Build(_env, front_BuildMode__Save_Obj, nil)
    } else if _switch9283 == Option_ModeKind__Shebang {
        {
            _modObj := front_convExp9101(Lns_2DDD(self.FP.LoadModule(_env, self.FP.scriptPath2Module(_env, self.option.ScriptPath))))
            if !Lns_IsNil( _modObj ) {
                modObj := _modObj
                var code LnsInt
                code = Depend_runMain(_env, modObj.(*Lns_luaValue).GetAt("__main"), self.option.ShebangArgList)
                _env.LuaVM.OS_exit(code)
            }
        }
    } else if _switch9283 == Option_ModeKind__Exec {
        _ = front_convExp9123(Lns_2DDD(self.FP.LoadModule(_env, self.FP.scriptPath2Module(_env, self.option.ScriptPath))))
        if self.option.Testing{
            var code string
            code = "local Testing = require( \"lune.base.Testing\" )\nreturn function( path )\n  Testing.run( path );\n  Testing.outputAllResult( io.stdout );\nend\n"
            var loaded LnsAny
            var mess LnsAny
            loaded,mess = _env.LuaVM.Load(code, nil)
            if loaded != nil{
                loaded_1157 := loaded.(*Lns_luaValue)
                {
                    _mod := front_convExp9177(Lns_2DDD(_env.LuaVM.RunLoadedfunc(loaded_1157,Lns_2DDD([]LnsAny{}))[0]))
                    if !Lns_IsNil( _mod ) {
                        mod := _mod
                        _env.LuaVM.RunLoadedfunc((mod.(*Lns_luaValue)),Lns_2DDD([]LnsAny{self.FP.scriptPath2Module(_env, self.option.ScriptPath)}))
                    }
                }
            } else {
                Lns_print([]LnsAny{mess})
            }
        }
    } else if _switch9283 == Option_ModeKind__BootC {
        self.FP.outputBootC(_env, self.option.ScriptPath)
    } else if _switch9283 == Option_ModeKind__Builtin {
        self.FP.OutputBuiltin(_env, self.option.ScriptPath)
    } else if _switch9283 == Option_ModeKind__MkMain {
        var mod string
        mod = self.FP.scriptPath2Module(_env, self.option.ScriptPath)
        {
            _mess := ConvGo_outputGoMain(_env, self.FP.getGoAppName(_env), mod, self.option.Testing, self.option.OutputPath, self.option.FP.Get_runtimeOpt(_env))
            if !Lns_IsNil( _mess ) {
                mess := _mess.(string)
                Util_errorLog(_env, mess)
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
func Newfront_DependMetaInfo(_env *LnsEnv, arg1 bool, arg2 string) *front_DependMetaInfo {
    obj := &front_DependMetaInfo{}
    obj.FP = obj
    obj.Initfront_DependMetaInfo(_env, arg1, arg2)
    return obj
}
func (self *front_DependMetaInfo) Initfront_DependMetaInfo(_env *LnsEnv, arg1 bool, arg2 string) {
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
func front_DependMetaInfo__fromMap_1167_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_DependMetaInfo_FromMap( arg1, paramList )
}
func front_DependMetaInfo__fromStem_1169_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
    CreateModuleId(_env *LnsEnv) *FrontInterface_ModuleId
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
func Newfront_MetaForBuildId(_env *LnsEnv, arg1 string, arg2 *LnsMap, arg3 *LnsList, arg4 bool) *front_MetaForBuildId {
    obj := &front_MetaForBuildId{}
    obj.FP = obj
    obj.Initfront_MetaForBuildId(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *front_MetaForBuildId) Initfront_MetaForBuildId(_env *LnsEnv, arg1 string, arg2 *LnsMap, arg3 *LnsList, arg4 bool) {
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
func front_MetaForBuildId__fromMap_1191_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_MetaForBuildId_FromMap( arg1, paramList )
}
func front_MetaForBuildId__fromStem_1193_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 387: decl @lune.@base.@front.MetaForBuildId.createModuleId
func (self *front_MetaForBuildId) CreateModuleId(_env *LnsEnv) *FrontInterface_ModuleId {
    return FrontInterface_ModuleId_createIdFromTxt(_env, self.G__buildId)
}

// 392: decl @lune.@base.@front.MetaForBuildId.LoadFromMeta
func front_MetaForBuildId_LoadFromMeta_1196_(_env *LnsEnv, metaPath string)(LnsAny, LnsAny) {
    {
        _fileObj := front_convExp1684(Lns_2DDD(Lns_io_open(metaPath, nil)))
        if !Lns_IsNil( _fileObj ) {
            fileObj := _fileObj.(Lns_luaStream)
            var luaCode LnsAny
            luaCode = fileObj.Read(_env, "*a")
            fileObj.Close(_env)
            if luaCode != nil{
                luaCode_305 := luaCode.(string)
                var meta LnsAny
                meta = front_convExp1674(Lns_2DDD(front_MetaForBuildId__fromStem_1193_(_env, _env.LuaVM.ExpandLuavalMap(Front_loadFromLuaTxt(_env, luaCode_305)),nil)))
                return meta, luaCode_305
            }
        }
    }
    return nil, nil
}


// declaration Class -- UpdateInfo
type front_UpdateInfoMtd interface {
    Get_dependsPath(_env *LnsEnv) LnsAny
    Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId
    Get_parser(_env *LnsEnv) LnsAny
    Get_scriptPath(_env *LnsEnv) string
    Get_uptodate(_env *LnsEnv) LnsAny
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
func Newfront_UpdateInfo(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 *FrontInterface_ModuleId, arg5 LnsAny) *front_UpdateInfo {
    obj := &front_UpdateInfo{}
    obj.FP = obj
    obj.Initfront_UpdateInfo(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *front_UpdateInfo) Initfront_UpdateInfo(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 *FrontInterface_ModuleId, arg5 LnsAny) {
    self.scriptPath = arg1
    self.dependsPath = arg2
    self.parser = arg3
    self.moduleId = arg4
    self.uptodate = arg5
}
func (self *front_UpdateInfo) Get_scriptPath(_env *LnsEnv) string{ return self.scriptPath }
func (self *front_UpdateInfo) Get_dependsPath(_env *LnsEnv) LnsAny{ return self.dependsPath }
func (self *front_UpdateInfo) Get_parser(_env *LnsEnv) LnsAny{ return self.parser }
func (self *front_UpdateInfo) Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId{ return self.moduleId }
func (self *front_UpdateInfo) Get_uptodate(_env *LnsEnv) LnsAny{ return self.uptodate }

func Lns_front_init(_env *LnsEnv) {
    if init_front { return }
    init_front = true
    front__mod__ = "@lune.@base.@front"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_Str_init(_env)
    Lns_frontInterface_init(_env)
    Lns_Parser_init(_env)
    Lns_convLua_init(_env)
    Lns_convGo_init(_env)
    Lns_TransUnit_init(_env)
    Lns_Util_init(_env)
    Lns_Option_init(_env)
    Lns_dumpNode_init(_env)
    Lns_glueFilter_init(_env)
    Lns_LuaVer_init( _env )
    Lns_Depend_init(_env)
    Lns_OutputDepend_init(_env)
    Lns_Ver_init(_env)
    Lns_LuaVer_init(_env)
    Lns_Log_init(_env)
    Lns_Formatter_init(_env)
    Lns_Testing_init(_env)
    Lns_GoMod_init(_env)
    Lns_Meta_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Nodes_init(_env)
    Lns_Ast_init(_env)
    Depend_setup(_env, Depend_UpdateVer(front__anonymous_1004_))
    front_forceUpdateMeta = true
}
func init() {
    init_front = false
}
