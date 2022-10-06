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
Val2 *AstInfo_ASTInfo
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
type Front_BuildMode = LnsAny
type Front_BuildMode__CreateAst struct{
}
var Front_BuildMode__CreateAst_Obj = &Front_BuildMode__CreateAst{}
func (self *Front_BuildMode__CreateAst) GetTxt() string {
return "BuildMode.CreateAst"
}
type Front_BuildMode__Output struct{
Val1 Lns_oStream
Val2 Lns_oStream
}
func (self *Front_BuildMode__Output) GetTxt() string {
return "BuildMode.Output"
}
type Front_BuildMode__Save struct{
}
var Front_BuildMode__Save_Obj = &Front_BuildMode__Save{}
func (self *Front_BuildMode__Save) GetTxt() string {
return "BuildMode.Save"
}
type Front_AstCallback func (_env *LnsEnv, arg1 *AstInfo_ASTInfo)
// for 1626: ExpCast
func conv2Form0_2180( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1731: ExpCast
func conv2Form0_2364( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1751: ExpCast
func conv2Form0_2431( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1241: ExpCast
func conv2Form4_128( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 394
func front_convExp1_848(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 466
func front_convExp1_1146(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 488
func front_convExp1_1222(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 830
func front_convExp2_1367(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 839
func front_convExp2_1396(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1750
func front_convExp0_2442(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1821
func front_convExp0_2881(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 411
func front_convExp1_884(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 472
func front_convExp1_1133(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 487
func front_convExp1_1201(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 517
func front_convExp1_1370(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 518
func front_convExp1_1389(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 559
func front_convExp2_41(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 741
func front_convExp2_837(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 770
func front_convExp2_989(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 823
func front_convExp2_1297(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1767
func front_convExp0_2536(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, string, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1770
func front_convExp0_2562(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1774
func front_convExp0_2588(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1778
func front_convExp0_2614(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1782
func front_convExp0_2640(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1786
func front_convExp0_2666(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1790
func front_convExp0_2692(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1836
func front_convExp0_2928(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, string, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 271
func front_convExp1_285(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 295
func front_convExp1_396(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 324
func front_convExp1_469(arg1 []LnsAny) *AstInfo_ASTInfo {
    return Lns_getFromMulti( arg1, 0 ).(*AstInfo_ASTInfo)
}
// for 341
func front_convExp1_551(arg1 []LnsAny) (*AstInfo_ASTInfo, *FrontInterface_ModuleInfo, *FrontInterface_ModuleMeta) {
    return Lns_getFromMulti( arg1, 0 ).(*AstInfo_ASTInfo), Lns_getFromMulti( arg1, 1 ).(*FrontInterface_ModuleInfo), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleMeta)
}
// for 445
func front_convExp1_1057(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 639
func front_convExp2_310(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 666
func front_convExp2_461(arg1 []LnsAny) (LnsAny, string, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 705
func front_convExp2_650(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 710
func front_convExp2_695(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 732
func front_convExp2_779(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 761
func front_convExp2_931(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 811
func front_convExp2_1175(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 893
func front_convExp3_83(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 959
func front_convExp3_479(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1002
func front_convExp3_782(arg1 []LnsAny) (string, LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1019
func front_convExp3_880(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1028
func front_convExp3_940(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1050
func front_convExp3_1058(arg1 []LnsAny) (*FrontInterface_ModuleMeta, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 )
}
// for 1139
func front_convExp3_1474(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1166
func front_convExp3_1581(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1192
func front_convExp3_1892(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1203
func front_convExp3_1959(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1320
func front_convExp4_397(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1327
func front_convExp4_459(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1363
func front_convExp4_557(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1365
func front_convExp4_572(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1369
func front_convExp4_586(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1370
func front_convExp4_602(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1379
func front_convExp4_627(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1409
func front_convExp4_716(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1410
func front_convExp4_730(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1411
func front_convExp4_743(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1413
func front_convExp4_760(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1576
func front_convExp0_1581(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1591
func front_convExp0_1648(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1592
func front_convExp0_1661(arg1 []LnsAny) (*FrontInterface_ModuleId, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 1 )
}
// for 1599
func front_convExp0_1706(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1655
func front_convExp0_2057(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1699
func front_convExp0_2272(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1725
func front_convExp0_2341(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1728
func front_convExp0_2358(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1748
func front_convExp0_2414(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1801
func front_convExp0_2744(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1819
func front_convExp0_2837(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1845
func front_convExp0_2978(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}









func front_Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_0_(_env *LnsEnv) string {
    return "NeedUpdate"
}
func front_Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1_(_env *LnsEnv) string {
    return "NeedUpdate"
}
func front_Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_2_(_env *LnsEnv) string {
    return "NeedUpdate"
}

func front_Front_getModuleIdAndCheckUptodate___anonymous_1_(_env *LnsEnv) string {
    return "not found meta"
}









func front__anonymous_0_(_env *LnsEnv, ver LnsInt) {
    LuaVer_setCurVer(_env, ver)
}
// 282: decl @lune.@base.@front.createPaser
func front_createPaser_6_(_env *LnsEnv, path string,mod string,stdinFile LnsAny) *Parser_Parser {
    return &Parser_StreamParser_create(_env, &Types_ParserSrc__LnsPath{path, mod, nil}, false, stdinFile, nil).Parser_Parser
}

// 378: decl @lune.@base.@front.ast2Lua
func front_ast2Lua_7_(_env *LnsEnv, ast *AstInfo_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,option *Option_Option) {
    var conv *ConvLua_FilterInfo
    conv = Converter_ast2LuaMain(_env, ast, streamName, stream, metaStream, convMode, inMacro, option)
    conv.FP.OutputLuaAndMeta(_env, Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode))
}

// 387: decl @lune.@base.@front.loadFromChunk
func front_loadFromChunk_8_(_env *LnsEnv, chunk LnsAny,err LnsAny) LnsAny {
    if err != nil{
        err_66 := err.(string)
        Util_errorLog(_env, err_66)
    }
    var ret LnsAny
    Lns_LockEnvSync( _env, 392, func () {
        if chunk != nil{
            chunk_70 := chunk.(*Lns_luaValue)
            {
                _work := front_convExp1_848(Lns_2DDD(_env.GetVM().RunLoadedfunc(chunk_70,Lns_2DDD([]LnsAny{}))[0]))
                if !Lns_IsNil( _work ) {
                    work := _work
                    ret = work
                } else {
                    ret = nil
                }
            }
        } else {
            panic("failed to error")
        }
    })
    return ret
}

// 407: decl @lune.@base.@front.loadFromLuaTxt
func Front_loadFromLuaTxt(_env *LnsEnv, txt string) LnsAny {
    var chunk LnsAny
    var err LnsAny
    Lns_LockEnvSync( _env, 410, func () {
        chunk, err = _env.GetVM().Load(txt, nil)
    })
    return front_loadFromChunk_8_(_env, chunk, err)
}

// 482: decl @lune.@base.@front.getMetaInfo
func front_getMetaInfo_12_(_env *LnsEnv, lnsPath string,mod string,outdir LnsAny)(LnsAny, string, string) {
    var moduleMetaPath string
    moduleMetaPath = lnsPath
    if outdir != nil{
        outdir_99 := outdir.(string)
        moduleMetaPath = _env.GetVM().String_format("%s/%s", []LnsAny{outdir_99, Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string)})
    }
    moduleMetaPath = front_convExp1_1201(Lns_2DDD(_env.GetVM().String_gsub(moduleMetaPath,"%.lns$", ".meta")))
    {
        _meta, _metaCode := front_MetaForBuildId_LoadFromMeta_5_(_env, moduleMetaPath)
        if !Lns_IsNil( _meta ) && !Lns_IsNil( _metaCode ) {
            meta := _meta.(*front_MetaForBuildId)
            metaCode := _metaCode.(string)
            return meta, moduleMetaPath, metaCode
        }
    }
    return nil, moduleMetaPath, ""
}

// 551: decl @lune.@base.@front.getModuleId
func front_getModuleId_13_(_env *LnsEnv, lnsPath string,mod string,outdir LnsAny,metaInfo LnsAny) *FrontInterface_ModuleId {
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
        metaInfo = front_convExp2_41(Lns_2DDD(front_getMetaInfo_12_(_env, lnsPath, mod, outdir)))
    }
    if metaInfo != nil{
        metaInfo_7 := metaInfo.(*front_MetaForBuildId)
        var buildId *FrontInterface_ModuleId
        buildId = metaInfo_7.FP.CreateModuleId(_env)
        buildCount = buildId.FP.Get_buildCount(_env)
    }
    return FrontInterface_ModuleId_createId(_env, fileTime, buildCount)
}




// 1303: decl @lune.@base.@front.outputDependInfo
func front_outputDependInfo_15_(_env *LnsEnv, stream LnsAny,metaInfo *front_MetaForBuildId,mod string) {
    if stream != nil{
        stream_20 := stream.(Lns_oStream)
        var dependInfo *OutputDepend_DependInfo
        dependInfo = NewOutputDepend_DependInfo(_env, mod)
        for _dependMod, _ := range( metaInfo.G__dependModuleMap.Items ) {
            dependMod := _dependMod.(string)
            dependInfo.FP.AddImpotModule(_env, dependMod)
        }
        for _, _subMod := range( metaInfo.G__subModuleMap.Items ) {
            subMod := _subMod.(string)
            dependInfo.FP.AddSubMod(_env, subMod)
        }
        dependInfo.FP.Output(_env, stream_20)
    }
}


// 1521: decl @lune.@base.@front.convertLnsCode2LuaCodeWithOpt
func Front_convertLnsCode2LuaCodeWithOpt(_env *LnsEnv, option *Option_Option,lnsCode string,path string,baseDir LnsAny) string {
    var front *front_Front
    front = Newfront_Front(_env, option, nil)
    return front.FP.convertLns2LuaCode(_env, NewFrontInterface_ImportModuleInfo(_env), TransUnit_AnalyzeMode__Compile, &Types_ParserSrc__LnsCode{lnsCode, path, nil}, baseDir, NewUtil_TxtStream(_env, lnsCode).FP, path)
}

// 1532: decl @lune.@base.@front.convertLnsCode2LuaCode
func Front_convertLnsCode2LuaCode(_env *LnsEnv, lnsCode string,path string,baseDir LnsAny) string {
    var option *Option_Option
    option = NewOption_Option(_env)
    option.ScriptPath = path
    option.UseLuneModule = Option_getRuntimeModule(_env)
    option.UseIpairs = true
    return Front_convertLnsCode2LuaCodeWithOpt(_env, option, lnsCode, path, baseDir)
}

// 1542: decl @lune.@base.@front.getBaseDir
func front_getBaseDir_19_(_env *LnsEnv, path string,userProjDir LnsAny)(string, LnsAny) {
    var projDir string
    if userProjDir != nil{
        userProjDir_88 := userProjDir.(string)
        projDir = userProjDir_88
    } else {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_car(_env.GetVM().String_find(path,"^/", nil, nil))) ||
            _env.SetStackVal( Lns_car(_env.GetVM().String_find(path,"^%.%./", nil, nil))) )){
            var parentPath string
            parentPath = Util_parentPath(_env, path)
            {
                _workDir := Util_searchProjDir(_env, parentPath)
                if !Lns_IsNil( _workDir ) {
                    workDir := _workDir.(string)
                    projDir = workDir
                } else {
                    projDir = parentPath
                }
            }
        } else { 
            return path, nil
        }
    }
    var scriptPath string
    if len(projDir) == 1{
        scriptPath = _env.GetVM().String_sub(path,2, nil)
    } else { 
        scriptPath = _env.GetVM().String_sub(path,len(projDir) + 2, nil)
    }
    return scriptPath, projDir
}

// 1573: decl @lune.@base.@front.getParseSrcAndBaseDir
func front_getParseSrcAndBaseDir_20_(_env *LnsEnv, path string,userProjDir LnsAny)(LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    var workPath string
    var projDir LnsAny
    workPath,projDir = front_getBaseDir_19_(_env, path, userProjDir)
    var mod string
    mod = Util_scriptPath2ModuleFromProjDir(_env, workPath, projDir)
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_13_(_env, path, mod, nil, nil)
    return &Types_ParserSrc__LnsPath{path, mod, nil}, mod, moduleId, projDir
}





// 1675: decl @lune.@base.@front.buildWithBuildMode
func Front_buildWithBuildMode(_env *LnsEnv, option *Option_Option,buildMode LnsAny,astCallback LnsAny) {
    var front *front_Front
    Lns_LockEnvSync( _env, 1679, func () {
        front = Newfront_Front(_env, option, nil)
    })
    front.FP.Build(_env, buildMode, astCallback)
}

// 1685: decl @lune.@base.@front.build
func Front_build(_env *LnsEnv, option *Option_Option,astCallback LnsAny) {
    Front_buildWithBuildMode(_env, option, Front_BuildMode__CreateAst_Obj, astCallback)
}

func front_Front_exec___anonymous_1_(_env *LnsEnv, ast *AstInfo_ASTInfo) {
    Util_println(_env, []LnsAny{ast.FP.Get_streamName(_env)})
}
func front_Front_exec___anonymous_2_(_env *LnsEnv, ast *AstInfo_ASTInfo) {
    var indexer *NodeIndexer_Indexer
    indexer = NewNodeIndexer_Indexer(_env, ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env))
    indexer.FP.Start(_env, ast.FP.Get_node(_env), NewLnsSet([]LnsAny{Nodes_NodeKind_get_Switch(_env), Nodes_NodeKind_get_Match(_env), Nodes_NodeKind_get_For(_env), Nodes_NodeKind_get_Apply(_env)}))
    indexer.FP.Dump(_env)
}
// 1875: decl @lune.@base.@front.exec
func Front_exec(_env *LnsEnv, args *LnsList) {
    var version LnsReal
    version = Lns_unwrapDefault( Lns_tonumber(Lns_car(_env.GetVM().String_gsub(Depend_getLuaVersion(_env),"^[^%d]+", "")).(string), nil), 0.0).(LnsReal)
    if version < 5.1{
        Util_errorLog(_env, _env.GetVM().String_format("LuneScript doesn't support this lua version(%s). %s\n", []LnsAny{version, "please use the version >= 5.1."}))
        _env.GetVM().OS_exit(1)
    }
    var option *Option_Option
    option = Option_analyze(_env, args)
    var front *front_Front
    front = Newfront_Front(_env, option, nil)
    front.FP.Exec(_env)
}

// 1894: decl @lune.@base.@front.setFront
func Front_setFront(_env *LnsEnv, bindModuleList *LnsList) {
    var option *Option_Option
    option = Option_createDefaultOption(_env, NewLnsList([]LnsAny{"dummy.lns"}), nil)
    Newfront_Front(_env, option, bindModuleList)
}

// 1899: decl @lune.@base.@front.__main
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




// 94: decl @lune.@base.@front.ModuleMgr.get
func (self *front_ModuleMgr) Get(_env *LnsEnv, mod string) LnsAny {
    return self.mod2info.FP.Get_map(_env).Get(mod)
}
// 98: decl @lune.@base.@front.ModuleMgr.getAst
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
    switch _matchExp0 := info.(type) {
    case *front_UptodateInfo__Update:
        ast := _matchExp0.Val2
        return ast
    case *front_UptodateInfo__Uptodate:
        return nil
    }
// insert a dummy
    return nil
}
// 112: decl @lune.@base.@front.ModuleMgr.getModList
func (self *front_ModuleMgr) GetModList(_env *LnsEnv) *LnsList {
    return self.mod2info.FP.Get_keyList(_env)
}
// 117: decl @lune.@base.@front.ModuleMgr.add
func (self *front_ModuleMgr) Add(_env *LnsEnv, ast *AstInfo_ASTInfo,moduleInfo *FrontInterface_ModuleInfo) {
    var mod string
    mod = moduleInfo.FP.Get_fullName(_env)
    self.mod2info.FP.Add(_env, mod, &front_UptodateInfo__Update{moduleInfo, ast}, true)
}
// 136: decl @lune.@base.@front.ModuleMgr.addMeta
func (self *front_ModuleMgr) AddMeta(_env *LnsEnv, mod string,meta *FrontInterface_ModuleMeta) {
    __func__ := "@lune.@base.@front.ModuleMgr.addMeta"
    if Lns_op_not(self.FP.Get(_env, mod)){
        Log_log(_env, Log_Level__Log, __func__, 138, Log_CreateMessage(func(_env *LnsEnv) string {
            return mod
        }))
        
        self.mod2info.FP.Add(_env, mod, &front_UptodateInfo__Uptodate{meta}, false)
    }
    self.loadedMetaMap.Set(mod,meta)
}
// 144: decl @lune.@base.@front.ModuleMgr.getMeta
func (self *front_ModuleMgr) GetMeta(_env *LnsEnv, mod string) LnsAny {
    return self.loadedMetaMap.Get(mod)
}
// 184: decl @lune.@base.@front.Front.regConvertedMap
func (self *front_Front) regConvertedMap(_env *LnsEnv, mod string,luaTxt string,meta *FrontInterface_ModuleMeta) {
    self.moduleMgr.FP.AddMeta(_env, mod, meta)
    self.convertedMap.Set(mod,luaTxt)
}
// 235: decl @lune.@base.@front.Front.getLoadInfo
func (self *front_Front) getLoadInfo(_env *LnsEnv, mod string) LnsAny {
    if self.option.Testing{
        return self.loadedMapTest.Get(mod)
    }
    return self.loadedMap.Get(mod)
}
// 242: decl @lune.@base.@front.Front.setLoadInfo
func (self *front_Front) setLoadInfo(_env *LnsEnv, mod string,info *front_LoadInfo) {
    if self.option.Testing{
        self.loadedMapTest.Set(mod,info)
    }
    self.loadedMap.Set(mod,info)
}
// 262: decl @lune.@base.@front.Front.error
func (self *front_Front) Error(_env *LnsEnv, message string) {
    Util_errorLog(_env, message)
    Util_printStackTrace(_env)
    _env.GetVM().OS_exit(1)
}
// 268: decl @lune.@base.@front.Front.loadLua
func (self *front_Front) loadLua(_env *LnsEnv, path string) LnsAny {
    var ret LnsAny
    Lns_LockEnvSync( _env, 270, func () {
        var chunk LnsAny
        var err LnsAny
        chunk,err = _env.GetVM().Loadfile(path)
        if chunk != nil{
            chunk_30 := chunk.(*Lns_luaValue)
            ret = Lns_unwrap( Lns_car(_env.GetVM().RunLoadedfunc(chunk_30,Lns_2DDD([]LnsAny{}))[0]))
        } else {
            Util_errorLog(_env, Lns_unwrapDefault( err, _env.GetVM().String_format("load error -- %s.", []LnsAny{path})).(string))
            ret = nil
        }
    })
    return ret
}
// 288: decl @lune.@base.@front.Front.scriptPath2Module
func (self *front_Front) scriptPath2Module(_env *LnsEnv, path string,baseDir LnsAny)(string, LnsAny) {
    var mod string
    mod = Util_scriptPath2ModuleFromProjDir(_env, path, baseDir)
    return mod, baseDir
}
// 293: decl @lune.@base.@front.Front.createPaser
func (self *front_Front) createPaser(_env *LnsEnv, scriptPath string,baseDir LnsAny) *Parser_Parser {
    var mod string
    mod = front_convExp1_396(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath, baseDir)))
    return front_createPaser_6_(_env, scriptPath, mod, self.option.FP.Get_stdinFile(_env))
}
// 299: decl @lune.@base.@front.Front.createAstSub
func (self *front_Front) createAstSub(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,baseDir LnsAny,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) LnsAny {
    {
        __exp := self.moduleMgr.FP.Get(_env, mod)
        if !Lns_IsNil( __exp ) {
            _exp := __exp
            switch _matchExp0 := _exp.(type) {
            case *front_UptodateInfo__Update:
                ast := _matchExp0.Val2
                return &Converter_CreateAstResult__Ast{ast}
            case *front_UptodateInfo__Uptodate:
                Util_err(_env, _env.GetVM().String_format("can't load the multiple module -- %s", []LnsAny{mod}))
            }
        }
    }
    {
        _creater := self.mod2astCreate.Get(mod)
        if !Lns_IsNil( _creater ) {
            creater := _creater.(*Converter_AstCreater)
            var ast *AstInfo_ASTInfo
            ast = front_convExp1_469(Lns_2DDD(creater.FP.GetAst(_env)))
            return &Converter_CreateAstResult__Ast{ast}
        }
    }
    var astCreater *Converter_AstCreater
    astCreater = NewConverter_AstCreater(_env, importModuleInfo, parserSrc, mod, baseDir, moduleId, analyzeModule, analyzeMode, pos, self.builtinFunc, self.option)
    self.mod2astCreate.Set(mod,astCreater)
    return &Converter_CreateAstResult__Creater{astCreater}
}
// 335: decl @lune.@base.@front.Front.applyAstResult
func (self *front_Front) applyAstResult(_env *LnsEnv, result LnsAny) *AstInfo_ASTInfo {
    __func__ := "@lune.@base.@front.Front.applyAstResult"
    switch _matchExp0 := result.(type) {
    case *Converter_CreateAstResult__Ast:
        ast := _matchExp0.Val1
        return ast
    case *Converter_CreateAstResult__Creater:
        astCreater := _matchExp0.Val1
        var ast *AstInfo_ASTInfo
        var moduleInfo *FrontInterface_ModuleInfo
        var meta *FrontInterface_ModuleMeta
        ast,moduleInfo,meta = astCreater.FP.GetAst(_env)
        Log_log(_env, Log_Level__Log, __func__, 342, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.GetVM().String_format("applyAstResult -- %s", []LnsAny{ast.FP.Get_exportInfo(_env).FP.Get_fullName(_env)})
        }))
        
        self.moduleMgr.FP.Add(_env, ast, moduleInfo)
        self.moduleMgr.FP.AddMeta(_env, ast.FP.Get_exportInfo(_env).FP.Get_fullName(_env), meta)
        self.mod2loader.Set(ast.FP.Get_exportInfo(_env).FP.Get_fullName(_env),nil)
        self.mod2astCreate.Set(ast.FP.Get_exportInfo(_env).FP.Get_fullName(_env),nil)
        return ast
    }
// insert a dummy
    return nil
}
// 356: decl @lune.@base.@front.Front.createAst
func (self *front_Front) createAst(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,baseDir LnsAny,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *AstInfo_ASTInfo {
    var ast *AstInfo_ASTInfo
    ast = self.FP.applyAstResult(_env, self.FP.createAstSub(_env, importModuleInfo, parserSrc, baseDir, mod, moduleId, analyzeModule, analyzeMode, pos))
    if self.option.DumpDebugAst{
        Util_println(_env, []LnsAny{_env.GetVM().String_format("=== ast:%s ==========", []LnsAny{mod})})
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, DumpNode_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt(_env, "", 0)))
    }
    return ast
}
// 416: decl @lune.@base.@front.Front.convertFromAst
func (self *front_Front) convertFromAst(_env *LnsEnv, ast *AstInfo_ASTInfo,streamName string,mode LnsInt)(string, string) {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var metaStream *Util_memStream
    metaStream = NewUtil_memStream(_env)
    front_ast2Lua_7_(_env, ast, streamName, stream.FP, metaStream.FP, mode, false, self.option)
    return metaStream.FP.Get_txt(_env), stream.FP.Get_txt(_env)
}
// 430: decl @lune.@base.@front.Front.loadFromLnsTxt
func (self *front_Front) LoadFromLnsTxt(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,baseDir LnsAny,name string,txt string) LnsAny {
    var transUnit *TransUnit_TransUnitCtrl
    transUnit = NewTransUnit_TransUnitCtrl(_env, FrontInterface_ModuleId__tempId, importModuleInfo, &NewConvLua_MacroEvalImp(_env, self.builtinFunc).Nodes_MacroEval, false, nil, nil, nil, self.option.TargetLuaVer, self.option.TransCtrlInfo, self.builtinFunc)
    var ast *AstInfo_ASTInfo
    ast = transUnit.FP.CreateAST(_env, &Types_ParserSrc__LnsCode{txt, name, nil}, false, baseDir, self.option.FP.Get_stdinFile(_env), false, _env.GetVM().String_format("$load%d", []LnsAny{self.loadCount}), nil)
    self.loadCount = self.loadCount + 1
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(_env, ast, name, ConvLua_ConvMode__ConvMeta)
    return Lns_unwrap( Front_loadFromLuaTxt(_env, luaTxt))
}
// 494: decl @lune.@base.@front.Front.searchModuleFile
func (self *front_Front) searchModuleFile(_env *LnsEnv, mod string,suffix string,baseDir LnsAny,outputDir LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.searchModuleFile"
    switch _matchExp0 := self.gomodMap.FP.ConvLocalModulePath(_env, mod, suffix, baseDir).(type) {
    case *GoMod_GoModResult__NotGo:
    case *GoMod_GoModResult__NotFound:
        return nil
    case *GoMod_GoModResult__Found:
        info := _matchExp0.Val1
        return info.FP.Get_path(_env)
    }
    var lnsSearchPath string
    lnsSearchPath = Lns_package_path
    if outputDir != nil{
        outputDir_110 := outputDir.(string)
        lnsSearchPath = _env.GetVM().String_format("%s/?%s;%s", []LnsAny{outputDir_110, suffix, Lns_package_path})
    }
    if baseDir != nil{
        baseDir_112 := baseDir.(string)
        lnsSearchPath = _env.GetVM().String_format("%s/?%s;%s", []LnsAny{baseDir_112, suffix, Lns_package_path})
    }
    {
        _projDir := self.option.FP.Get_projDir(_env)
        if !Lns_IsNil( _projDir ) {
            projDir := _projDir.(string)
            lnsSearchPath = _env.GetVM().String_format("%s;%s", []LnsAny{Util_pathJoin(_env, projDir, _env.GetVM().String_format("?%s", []LnsAny{suffix})), Lns_package_path})
        }
    }
    lnsSearchPath = front_convExp1_1370(Lns_2DDD(_env.GetVM().String_gsub(lnsSearchPath,"%.lua$", suffix)))
    lnsSearchPath = front_convExp1_1389(Lns_2DDD(_env.GetVM().String_gsub(lnsSearchPath,"%.lua;", suffix + ";")))
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
                    Log_log(_env, Log_Level__Err, __func__, 530, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.GetVM().String_format("not found '%s' at %s", []LnsAny{mod, latestProjSearchPath})
                    }))
                    
                    return nil
                }
            }
        } else {
            foundPath = _foundPath.(string)
        }
    }
    return Lns_car(_env.GetVM().String_gsub(foundPath,"^./", "")).(string)
}
// 591: decl @lune.@base.@front.Front.getModuleIdAndCheckUptodate
func (self *front_Front) getModuleIdAndCheckUptodate(_env *LnsEnv, lnsPath string,mod string)(*FrontInterface_ModuleId, LnsAny) {
    __func__ := "@lune.@base.@front.Front.getModuleIdAndCheckUptodate"
    var uptodate LnsAny
    uptodate = front_ModuleUptodate__NeedUpdate_Obj
    switch _matchExp0 := self.option.TransCtrlInfo.UptodateMode.(type) {
    case *Types_CheckingUptodateMode__ForceAll:
        return FrontInterface_ModuleId__tempId, uptodate
    case *Types_CheckingUptodateMode__Force1:
        forceMod := _matchExp0.Val1
        if mod == forceMod{
            return FrontInterface_ModuleId__tempId, uptodate
        }
    case *Types_CheckingUptodateMode__Normal:
    case *Types_CheckingUptodateMode__Touch:
    }
    var front_checkDependUptodate func(_env *LnsEnv, metaTime LnsReal,metaInfo *front_MetaForBuildId,metaCode string) LnsAny
    front_checkDependUptodate = func(_env *LnsEnv, metaTime LnsReal,metaInfo *front_MetaForBuildId,metaCode string) LnsAny {
        __func__ := "@lune.@base.@front.Front.getModuleIdAndCheckUptodate.checkDependUptodate"
        for _depMod, _dependItem := range( metaInfo.G__dependModuleMap.Items ) {
            depMod := _depMod.(string)
            dependItem := _dependItem.(front_DependMetaInfoDownCast).Tofront_DependMetaInfo()
            var modMetaPath string
            
            {
                _modMetaPath := self.FP.searchModuleFile(_env, depMod, ".meta", nil, self.option.OutputDir)
                if _modMetaPath == nil{
                    Log_log(_env, Log_Level__Debug, __func__, 627, Log_CreateMessage(front_Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_0_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    modMetaPath = _modMetaPath.(string)
                }
            }
            var time LnsReal
            
            {
                _time := Depend_getFileLastModifiedTime(_env, modMetaPath)
                if _time == nil{
                    Log_log(_env, Log_Level__Debug, __func__, 632, Log_CreateMessage(front_Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    time = _time.(LnsReal)
                }
            }
            if time > metaTime{
                var dependMeta *front_MetaForBuildId
                
                {
                    _dependMeta := front_convExp2_310(Lns_2DDD(front_MetaForBuildId_LoadFromMeta_5_(_env, modMetaPath)))
                    if _dependMeta == nil{
                        Log_log(_env, Log_Level__Debug, __func__, 640, Log_CreateMessage(front_Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_2_))
                        
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
                    Log_log(_env, Log_Level__Debug, __func__, 650, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.GetVM().String_format("NeedUpdate: %s, %d, %d", []LnsAny{modMetaPath, metaModuleId.FP.Get_buildCount(_env), orgMetaModuleId.FP.Get_buildCount(_env)})
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
    metaInfo,metaPath,metaCode = front_getMetaInfo_12_(_env, lnsPath, mod, self.option.OutputDir)
    if metaInfo != nil{
        metaInfo_52 := metaInfo.(*front_MetaForBuildId)
        if metaInfo_52.G__enableTest == self.option.Testing{
            var buildId *FrontInterface_ModuleId
            buildId = FrontInterface_ModuleId_createIdFromTxt(_env, metaInfo_52.G__buildId)
            if buildId != FrontInterface_ModuleId__tempId{
                var lnsTime LnsAny
                lnsTime = Depend_getFileLastModifiedTime(_env, lnsPath)
                var metaTime LnsAny
                metaTime = Depend_getFileLastModifiedTime(_env, metaPath)
                if lnsTime != nil && metaTime != nil{
                    lnsTime_59 := lnsTime.(LnsReal)
                    metaTime_60 := metaTime.(LnsReal)
                    if lnsTime_59 == buildId.FP.Get_modTime(_env){
                        uptodate = front_checkDependUptodate(_env, metaTime_60, metaInfo_52, metaCode)
                    }
                }
            }
        } else { 
        }
    } else {
        Log_log(_env, Log_Level__Debug, __func__, 688, Log_CreateMessage(front_Front_getModuleIdAndCheckUptodate___anonymous_1_))
        
    }
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_13_(_env, lnsPath, mod, self.option.OutputDir, metaInfo)
    if moduleId == FrontInterface_ModuleId__tempId{
        Util_err(_env, _env.GetVM().String_format("not found -- %s", []LnsAny{lnsPath}))
    }
    return moduleId, uptodate
}
// 701: decl @lune.@base.@front.Front.convertLns2LuaCode
func (self *front_Front) convertLns2LuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,analyzeMode LnsInt,parserSrc LnsAny,baseDir LnsAny,stream Lns_iStream,streamName string) string {
    var mod string
    mod = front_convExp2_650(Lns_2DDD(self.FP.scriptPath2Module(_env, streamName, baseDir)))
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, parserSrc, baseDir, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, analyzeMode, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(_env, ast, streamName, ConvLua_ConvMode__ConvMeta)
    return luaTxt
}
// 715: decl @lune.@base.@front.Front.getGoAppName
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
// 722: decl @lune.@base.@front.Front.createGoOption
func (self *front_Front) createGoOption(_env *LnsEnv, scriptPath string) *ConvGo_Option {
    var packageName string
    {
        __exp := self.option.PackageName
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(string)
            packageName = _exp
        } else {
            if Lns_op_not(Lns_car(_env.GetVM().String_find(scriptPath,"/", nil, nil))){
                packageName = "main"
            } else { 
                var parentPath string
                parentPath = front_convExp2_779(Lns_2DDD(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(scriptPath,"/[^/]+$", "")).(string),".*/", "")))
                if len(parentPath) == 0{
                    packageName = "main"
                } else if parentPath == "."{
                    packageName = "main"
                } else if parentPath == ".."{
                    packageName = "main"
                } else { 
                    packageName = front_convExp2_837(Lns_2DDD(_env.GetVM().String_gsub(parentPath,"[^%w]", "")))
                }
            }
        }
    }
    return NewConvGo_Option(_env, packageName, self.FP.getGoAppName(_env), self.option.MainModule, self.option.FP.Get_addEnvArg(_env), self.option.ConvGoRunnerNum, self.option.FP.Get_sortGenerateCode(_env))
}
// 751: decl @lune.@base.@front.Front.createPythonOption
func (self *front_Front) createPythonOption(_env *LnsEnv, scriptPath string) *ConvPython_Option {
    var packageName string
    {
        __exp := self.option.PackageName
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(string)
            packageName = _exp
        } else {
            if Lns_op_not(Lns_car(_env.GetVM().String_find(scriptPath,"/", nil, nil))){
                packageName = "main"
            } else { 
                var parentPath string
                parentPath = front_convExp2_931(Lns_2DDD(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(scriptPath,"/[^/]+$", "")).(string),".*/", "")))
                if len(parentPath) == 0{
                    packageName = "main"
                } else if parentPath == "."{
                    packageName = "main"
                } else if parentPath == ".."{
                    packageName = "main"
                } else { 
                    packageName = front_convExp2_989(Lns_2DDD(_env.GetVM().String_gsub(parentPath,"[^%w]", "")))
                }
            }
        }
    }
    return NewConvPython_Option(_env, packageName, self.FP.getGoAppName(_env), self.option.MainModule, false, self.option.ConvGoRunnerNum)
}
// 780: decl @lune.@base.@front.Front.convertToLanguage
func (self *front_Front) convertToLanguage(_env *LnsEnv, ast *AstInfo_ASTInfo,stream Lns_oStream,path string) {
    if _switch0 := self.option.ConvTo; _switch0 == Types_Lang__Go {
        var conv *Nodes_Filter
        conv = ConvGo_createFilter(_env, self.option.Testing, path, stream, ast, self.FP.createGoOption(_env, path))
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast.FP.Get_node(_env))))
    } else if _switch0 == Types_Lang__Python {
        var conv *Nodes_Filter
        conv = ConvPython_createFilter(_env, self.option.Testing, path, stream, ast, self.FP.createPythonOption(_env, path))
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvPython_Opt2Stem(NewConvPython_Opt(_env, ast.FP.Get_node(_env))))
    }
}
// 802: decl @lune.@base.@front.Front.loadParserToLuaCode
func (self *front_Front) loadParserToLuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,path string,mod string,baseDir LnsAny)(*FrontInterface_ModuleMeta, string) {
    __func__ := "@lune.@base.@front.Front.loadParserToLuaCode"
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_13_(_env, path, mod, nil, nil)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, parserSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(_env, ast, path, ConvLua_ConvMode__ConvMeta)
    Log_log(_env, Log_Level__Trace, __func__, 812, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("Meta = %s", []LnsAny{metaTxt})
    }))
    
    if self.option.UpdateOnLoad{
        var front_saveFile func(_env *LnsEnv, suffix string,txt string,byteCompile bool,stripDebugInfo bool,checkUpdate bool)
        front_saveFile = func(_env *LnsEnv, suffix string,txt string,byteCompile bool,stripDebugInfo bool,checkUpdate bool) {
            var newpath string
            newpath = ""
            {
                _dir := self.option.OutputDir
                if !Lns_IsNil( _dir ) {
                    dir := _dir.(string)
                    newpath = _env.GetVM().String_format("%s/%s%s", []LnsAny{dir, Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string), suffix})
                } else {
                    newpath = front_convExp2_1297(Lns_2DDD(_env.GetVM().String_gsub(path,".lns$", suffix)))
                }
            }
            var saveTxt string
            saveTxt = txt
            if byteCompile{
                saveTxt = Converter_byteCompileFromLuaTxt(_env, saveTxt, stripDebugInfo)
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(front_forceUpdateMeta)) &&
                _env.SetStackVal( checkUpdate) ).(bool)){
                {
                    _fileObj := front_convExp2_1367(Lns_2DDD(Lns_io_open(newpath, nil)))
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
                _fileObj := front_convExp2_1396(Lns_2DDD(Lns_io_open(newpath, "w")))
                if !Lns_IsNil( _fileObj ) {
                    fileObj := _fileObj.(Lns_luaStream)
                    fileObj.Write(_env, saveTxt)
                    fileObj.Close(_env)
                }
            }
        }
        front_saveFile(_env, ".lua", luaTxt, self.option.ByteCompile, self.option.StripDebugInfo, false)
        front_saveFile(_env, ".meta", metaTxt, self.option.ByteCompile, true, true)
        {
            var memStream *Util_memStream
            memStream = NewUtil_memStream(_env)
            self.FP.convertToLanguage(_env, ast, memStream.FP, path)
            {
                _convTo := self.option.ConvTo
                if !Lns_IsNil( _convTo ) {
                    convTo := _convTo.(LnsInt)
                    var suffix LnsAny
                    if _switch0 := convTo; _switch0 == Types_Lang__Same || _switch0 == Types_Lang__Lua {
                        suffix = nil
                    } else if _switch0 == Types_Lang__Go {
                        suffix = ".go"
                    } else if _switch0 == Types_Lang__Python {
                        suffix = ".py"
                    } else if _switch0 == Types_Lang__C {
                        suffix = ".c"
                    }
                    if suffix != nil{
                        suffix_145 := suffix.(string)
                        front_saveFile(_env, suffix_145, memStream.FP.Get_txt(_env), false, false, false)
                    }
                }
            }
        }
    }
    return Lns_unwrap( self.moduleMgr.FP.GetMeta(_env, mod)).(*FrontInterface_ModuleMeta), luaTxt
}
// 887: decl @lune.@base.@front.Front.loadFile
func (self *front_Front) loadFile(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,baseDir LnsAny,path string,mod string)(*FrontInterface_ModuleMeta, LnsAny) {
    __func__ := "@lune.@base.@front.Front.loadFile"
    Log_log(_env, Log_Level__Info, __func__, 891, Log_CreateMessage(func(_env *LnsEnv) string {
        __func__ := "@lune.@base.@front.Front.loadFile.<anonymous>"
        return _env.GetVM().String_format("start %s:%s", []LnsAny{__func__, mod})
    }))
    
    var meta *FrontInterface_ModuleMeta
    var luaTxt string
    meta,luaTxt = self.FP.loadParserToLuaCode(_env, importModuleInfo, &Types_ParserSrc__LnsPath{path, mod, nil}, path, mod, baseDir)
    {
        _preLoadInfo := self.preloadedModMap.Get(mod)
        if !Lns_IsNil( _preLoadInfo ) {
            preLoadInfo := _preLoadInfo
            return meta, preLoadInfo
        }
    }
    return meta, Lns_unwrap( Front_loadFromLuaTxt(_env, luaTxt))
}
// 904: decl @lune.@base.@front.Front.searchModule
func (self *front_Front) SearchModule(_env *LnsEnv, mod string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(_env, mod, ".lns", baseDir, addSearchPath)
}
// 908: decl @lune.@base.@front.Front.searchLuaFile
func (self *front_Front) searchLuaFile(_env *LnsEnv, moduleFullName string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(_env, moduleFullName, ".lua", baseDir, addSearchPath)
}
// 920: decl @lune.@base.@front.Front.checkUptodateMeta
func (self *front_Front) checkUptodateMeta(_env *LnsEnv, lnsPath string,metaPath string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.checkUptodateMeta"
    var metaObj LnsAny
    
    {
        _metaObj := self.FP.loadLua(_env, metaPath)
        if _metaObj == nil{
            Log_log(_env, Log_Level__Warn, __func__, 925, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.GetVM().String_format("load error -- %s", []LnsAny{metaPath})
            }))
            
            return nil
        } else {
            metaObj = _metaObj
        }
    }
    var meta *Lns_luaValue
    var valid bool
    valid = true
    Lns_LockEnvSync( _env, 931, func () {
        meta = metaObj.(*Lns_luaValue)
        if meta.GetAt( "__formatVersion" ).(string) != Ver_metaVersion{
            Log_log(_env, Log_Level__Warn, __func__, 934, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.GetVM().String_format("unmatch meta version -- %s", []LnsAny{metaPath})
            }))
            
            valid = false
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( valid) &&
            _env.SetStackVal( meta.GetAt( "__hasTest" ).(bool)) ).(bool)){
            if meta.GetAt( "__enableTest" ).(bool) != self.option.Testing{
                Log_log(_env, Log_Level__Warn, __func__, 940, Log_CreateMessage(func(_env *LnsEnv) string {
                    return _env.GetVM().String_format("unmatch test setting -- %s", []LnsAny{metaPath})
                }))
                
                valid = false
            }
        }
        if valid{
            {
                _foreachExp0 := meta.GetAt( "__dependModuleMap" ).(*Lns_luaValue)
                _foreachKey0, _ := _foreachExp0.Get1stFromMap()
                for _foreachKey0 != nil {
                    moduleFullName := _foreachKey0.(string)
                    {
                        _moduleLnsPath := self.FP.SearchModule(_env, moduleFullName, baseDir, addSearchPath)
                        if !Lns_IsNil( _moduleLnsPath ) {
                            moduleLnsPath := _moduleLnsPath.(string)
                            {
                                _moduleLuaPath := self.FP.searchLuaFile(_env, moduleFullName, baseDir, addSearchPath)
                                if !Lns_IsNil( _moduleLuaPath ) {
                                    moduleLuaPath := _moduleLuaPath.(string)
                                    if Lns_op_not(Util_getReadyCode(_env, moduleLnsPath, metaPath)){
                                        Log_log(_env, Log_Level__Warn, __func__, 954, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.GetVM().String_format("not ready -- %s, %s", []LnsAny{moduleLnsPath, metaPath})
                                        }))
                                        
                                        valid = false
                                        break
                                    }
                                    var moduleMetaPath string
                                    moduleMetaPath = front_convExp3_479(Lns_2DDD(_env.GetVM().String_gsub(moduleLuaPath,"%.lua$", ".meta")))
                                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                        _env.SetStackVal( Depend_existFile(_env, moduleMetaPath)) &&
                                        _env.SetStackVal( Lns_op_not(Util_getReadyCode(_env, moduleMetaPath, metaPath))) ).(bool)){
                                        Log_log(_env, Log_Level__Warn, __func__, 963, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.GetVM().String_format("not ready -- %s, %s", []LnsAny{moduleMetaPath, metaPath})
                                        }))
                                        
                                        valid = false
                                        break
                                    }
                                } else {
                                    Log_log(_env, Log_Level__Warn, __func__, 969, Log_CreateMessage(func(_env *LnsEnv) string {
                                        return _env.GetVM().String_format("not found .lua file for -- %s", []LnsAny{moduleFullName})
                                    }))
                                    
                                    valid = false
                                    break
                                }
                            }
                        } else {
                            Log_log(_env, Log_Level__Warn, __func__, 975, Log_CreateMessage(func(_env *LnsEnv) string {
                                return _env.GetVM().String_format("not found .lns file -- %s", []LnsAny{moduleFullName})
                            }))
                            
                            valid = false
                            break
                        }
                    }
                    _foreachKey0, _ = _foreachExp0.NextFromMap( _foreachKey0 )
                }
            }
        }
    })
    if Lns_op_not(valid){
        return nil
    }
    return NewFrontInterface_ModuleMeta(_env, lnsPath, &FrontInterface_MetaOrModule__MetaRaw{meta})
}
// 997: decl @lune.@base.@front.Front.loadModuleWithBaseDir
func (self *front_Front) loadModuleWithBaseDir(_env *LnsEnv, orgMod string,baseDir LnsAny)(LnsAny, *FrontInterface_ModuleMeta) {
    __func__ := "@lune.@base.@front.Front.loadModuleWithBaseDir"
    Log_log(_env, Log_Level__Info, __func__, 1000, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("loadModule  -- %s:%s", []LnsAny{orgMod, baseDir})
    }))
    
    var mod string
    _,_,mod = self.gomodMap.FP.GetLuaModulePath(_env, orgMod, baseDir)
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
                        panic(_env.GetVM().String_format("nothing meta -- %s", []LnsAny{mod}))
                    }
                }
            } else {
                {
                    _lnsPath := self.FP.SearchModule(_env, orgMod, baseDir, nil)
                    if !Lns_IsNil( _lnsPath ) {
                        lnsPath := _lnsPath.(string)
                        var luaPath LnsAny
                        luaPath = front_convExp3_880(Lns_2DDD(_env.GetVM().String_gsub(lnsPath, "%.lns$", ".lua")))
                        {
                            _dir := self.option.OutputDir
                            if !Lns_IsNil( _dir ) {
                                dir := _dir.(string)
                                luaPath = self.FP.searchLuaFile(_env, orgMod, nil, dir)
                            }
                        }
                        var loadVal LnsAny
                        loadVal = nil
                        if luaPath != nil{
                            luaPath_67 := luaPath.(string)
                            if Util_getReadyCode(_env, lnsPath, luaPath_67){
                                var metaPath string
                                metaPath = front_convExp3_940(Lns_2DDD(_env.GetVM().String_gsub(luaPath_67, "%.lua$", ".meta")))
                                if Util_getReadyCode(_env, lnsPath, metaPath){
                                    {
                                        _preLoadInfo := self.preloadedModMap.Get(mod)
                                        if !Lns_IsNil( _preLoadInfo ) {
                                            preLoadInfo := _preLoadInfo
                                            loadVal = preLoadInfo
                                        } else {
                                            loadVal = self.FP.loadLua(_env, luaPath_67)
                                        }
                                    }
                                    {
                                        __exp := loadVal
                                        if !Lns_IsNil( __exp ) {
                                            _exp := __exp
                                            {
                                                _meta := self.FP.checkUptodateMeta(_env, lnsPath, metaPath, baseDir, self.option.OutputDir)
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
                            meta,workVal = self.FP.loadFile(_env, NewFrontInterface_ImportModuleInfo(_env), baseDir, lnsPath, mod)
                            self.FP.setLoadInfo(_env, mod, Newfront_LoadInfo(_env, workVal, meta))
                        }
                    } else {
                        if self.bindModuleSet.Has(mod){
                            Log_log(_env, Log_Level__Warn, __func__, 1057, Log_CreateMessage(func(_env *LnsEnv) string {
                                return _env.GetVM().String_format("load from the binding -- %s", []LnsAny{mod})
                            }))
                            
                            var workMod LnsAny
                            Lns_LockEnvSync( _env, 1061, func () {
                                workMod = Lns_require(mod)
                            })
                            var meta *FrontInterface_ModuleMeta
                            meta = NewFrontInterface_ModuleMeta(_env, Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string) + ".lns", &FrontInterface_MetaOrModule__MetaRaw{Lns_unwrap( Front_loadFromLuaTxt(_env, "return {}"))})
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
    panic(_env.GetVM().String_format("load error, %s", []LnsAny{mod}))
// insert a dummy
    return nil,nil
}
// 1088: decl @lune.@base.@front.Front.loadModule
func (self *front_Front) LoadModule(_env *LnsEnv, mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    return self.FP.loadModuleWithBaseDir(_env, mod, nil)
}
// 1095: decl @lune.@base.@front.Front.getLuaModulePath
func (self *front_Front) GetLuaModulePath(_env *LnsEnv, mod string,baseDir LnsAny)(string, LnsAny, string) {
    return self.gomodMap.FP.GetLuaModulePath(_env, mod, baseDir)
}
// 1099: decl @lune.@base.@front.Front.loadMeta
func (self *front_Front) LoadMeta(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,mod string,orgMod string,baseDir LnsAny,loader FrontInterface_ModuleLoader) LnsAny {
    __func__ := "@lune.@base.@front.Front.loadMeta"
    {
        _creater := self.mod2astCreate.Get(orgMod)
        if !Lns_IsNil( _creater ) {
            creater := _creater.(*Converter_AstCreater)
            var exportInfo *FrontInterface_ExportInfo
            
            {
                _exportInfo := creater.FP.GetExportInfo(_env)
                if _exportInfo == nil{
                    return nil
                } else {
                    exportInfo = _exportInfo.(*FrontInterface_ExportInfo)
                }
            }
            var meta *FrontInterface_ModuleMeta
            meta = NewFrontInterface_ModuleMeta(_env, exportInfo.FP.Get_streamName(_env), &FrontInterface_MetaOrModule__Export{exportInfo})
            self.moduleMgr.FP.AddMeta(_env, orgMod, meta)
        }
    }
    {
        _workLoader := self.mod2loader.Get(orgMod)
        if !Lns_IsNil( _workLoader ) {
            workLoader := _workLoader.(FrontInterface_ModuleLoader)
            var exportInfo *FrontInterface_ExportInfo
            
            {
                _exportInfo := workLoader.GetExportInfo(_env)
                if _exportInfo == nil{
                    return nil
                } else {
                    exportInfo = _exportInfo.(*FrontInterface_ExportInfo)
                }
            }
            var meta *FrontInterface_ModuleMeta
            meta = NewFrontInterface_ModuleMeta(_env, exportInfo.FP.Get_streamName(_env), &FrontInterface_MetaOrModule__Export{exportInfo})
            self.moduleMgr.FP.AddMeta(_env, orgMod, meta)
            return meta
        }
    }
    if Lns_op_not(self.moduleMgr.FP.GetMeta(_env, orgMod)){
        self.mod2loader.Set(orgMod,loader)
        {
            __exp := self.FP.getLoadInfo(_env, orgMod)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*front_LoadInfo)
                self.moduleMgr.FP.AddMeta(_env, orgMod, _exp.Meta)
            } else {
                {
                    _lnsPath := self.FP.SearchModule(_env, mod, baseDir, nil)
                    if !Lns_IsNil( _lnsPath ) {
                        lnsPath := _lnsPath.(string)
                        var meta LnsAny
                        meta = nil
                        if Lns_op_not(self.targetSet.Has(orgMod)){
                            var luaPath LnsAny
                            luaPath = front_convExp3_1474(Lns_2DDD(_env.GetVM().String_gsub(lnsPath, "%.lns$", ".lua")))
                            if Lns_op_not(baseDir){
                                {
                                    _dir := self.option.OutputDir
                                    if !Lns_IsNil( _dir ) {
                                        dir := _dir.(string)
                                        luaPath = self.FP.searchLuaFile(_env, orgMod, nil, dir)
                                    }
                                }
                            }
                            if luaPath != nil{
                                luaPath_119 := luaPath.(string)
                                var forceFlag bool
                                switch _matchExp0 := self.option.TransCtrlInfo.UptodateMode.(type) {
                                case *Types_CheckingUptodateMode__ForceAll:
                                    forceFlag = true
                                case *Types_CheckingUptodateMode__Force1:
                                    forceMod := _matchExp0.Val1
                                    forceFlag = orgMod == forceMod
                                case *Types_CheckingUptodateMode__Normal:
                                    forceFlag = false
                                case *Types_CheckingUptodateMode__Touch:
                                    forceFlag = false
                                }
                                if Lns_op_not(forceFlag){
                                    if Util_getReadyCode(_env, lnsPath, luaPath_119){
                                        var metaPath string
                                        metaPath = front_convExp3_1581(Lns_2DDD(_env.GetVM().String_gsub(luaPath_119, "%.lua$", ".meta")))
                                        if Util_getReadyCode(_env, lnsPath, metaPath){
                                            meta = self.FP.checkUptodateMeta(_env, lnsPath, metaPath, baseDir, self.option.OutputDir)
                                        } else { 
                                            Log_log(_env, Log_Level__Warn, __func__, 1171, Log_CreateMessage(func(_env *LnsEnv) string {
                                                return _env.GetVM().String_format("%s not ready meta %s, %s", []LnsAny{orgMod, lnsPath, metaPath})
                                            }))
                                            
                                        }
                                    } else { 
                                        Log_log(_env, Log_Level__Warn, __func__, 1175, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.GetVM().String_format("%s not ready lua %s, %s", []LnsAny{orgMod, lnsPath, luaPath_119})
                                        }))
                                        
                                    }
                                } else { 
                                    Log_log(_env, Log_Level__Warn, __func__, 1179, Log_CreateMessage(func(_env *LnsEnv) string {
                                        return _env.GetVM().String_format("force analyze -- %s", []LnsAny{orgMod})
                                    }))
                                    
                                }
                            } else {
                                Log_log(_env, Log_Level__Warn, __func__, 1183, Log_CreateMessage(func(_env *LnsEnv) string {
                                    return _env.GetVM().String_format("%s not found lua in %s", []LnsAny{orgMod, self.option.OutputDir})
                                }))
                                
                            }
                        }
                        if meta != nil{
                            meta_143 := meta.(*FrontInterface_ModuleMeta)
                            self.moduleMgr.FP.AddMeta(_env, orgMod, meta_143)
                        } else {
                            var metawork *FrontInterface_ModuleMeta
                            var luaTxt string
                            metawork,luaTxt = self.FP.loadParserToLuaCode(_env, importModuleInfo, &Types_ParserSrc__LnsPath{lnsPath, orgMod, nil}, lnsPath, orgMod, baseDir)
                            self.FP.regConvertedMap(_env, orgMod, luaTxt, metawork)
                        }
                    } else {
                        {
                            _lnsCode := Depend_getBindLns(_env, orgMod)
                            if !Lns_IsNil( _lnsCode ) {
                                lnsCode := _lnsCode.(string)
                                var path string
                                path = Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string) + ".lns"
                                var meta *FrontInterface_ModuleMeta
                                var luaTxt string
                                meta,luaTxt = self.FP.loadParserToLuaCode(_env, importModuleInfo, &Types_ParserSrc__LnsCode{lnsCode, orgMod, nil}, path, orgMod, baseDir)
                                self.FP.regConvertedMap(_env, orgMod, luaTxt, meta)
                            }
                        }
                    }
                }
            }
        }
    }
    {
        _meta := self.moduleMgr.FP.GetMeta(_env, orgMod)
        if !Lns_IsNil( _meta ) {
            meta := _meta.(*FrontInterface_ModuleMeta)
            return meta
        }
    }
    return nil
}
// 1221: decl @lune.@base.@front.Front.dumpTokenize
func (self *front_Front) DumpTokenize(_env *LnsEnv, scriptPath string,baseDir LnsAny) {
    var parser *Parser_Parser
    parser = self.FP.createPaser(_env, scriptPath, baseDir)
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
        Util_println(_env, []LnsAny{Types_TokenKind_getTxt( token.Kind), token.Pos.LineNo, token.Pos.Column, token.Txt})
    }
}
// 1234: decl @lune.@base.@front.Front.dumpAst
func (self *front_Front) DumpAst(_env *LnsEnv, parserSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    Depend_profile(_env, self.option.ValidProf, conv2Form4_128(func(_env *LnsEnv) {
        var ast *AstInfo_ASTInfo
        ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), parserSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, DumpNode_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt(_env, "", 0)))
    }), mod + ".profi")
}
// 1255: decl @lune.@base.@front.Front.format
func (self *front_Front) Format(_env *LnsEnv, parserSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), parserSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, Formatter_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), Lns_io_stdout), Formatter_Opt2Stem(NewFormatter_Opt(_env, ast.FP.Get_node(_env))))
}
// 1268: decl @lune.@base.@front.Front.checkDiag
func (self *front_Front) CheckDiag(_env *LnsEnv, parserSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    Util_setErrorCode(_env, 0)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), parserSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Diag, nil)
}
// 1276: decl @lune.@base.@front.Front.complete
func (self *front_Front) Complete(_env *LnsEnv, parserSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), parserSrc, baseDir, mod, moduleId, self.option.AnalyzeModule, TransUnit_AnalyzeMode__Complete, self.option.AnalyzePos)
}
// 1284: decl @lune.@base.@front.Front.inquire
func (self *front_Front) Inquire(_env *LnsEnv, parserSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), parserSrc, baseDir, mod, moduleId, self.option.AnalyzeModule, TransUnit_AnalyzeMode__Inquire, self.option.AnalyzePos)
}
// 1293: decl @lune.@base.@front.Front.createGlue
func (self *front_Front) CreateGlue(_env *LnsEnv, parserSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), parserSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var filter *Nodes_Filter
    filter = GlueFilter_createFilter(_env, self.option.OutputDir)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, filter, 0)
}
// 1316: decl @lune.@base.@front.Front.convertToLua
func (self *front_Front) convertToLua(_env *LnsEnv, scriptPath string,baseDir LnsAny,convMode LnsInt,streamLua Lns_oStream,streamMeta Lns_oStream) LnsAny {
    var mod string
    mod = front_convExp4_397(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath, baseDir)))
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_13_(_env, scriptPath, mod, nil, nil)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{scriptPath, mod, nil}, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(_env, ast, scriptPath, convMode)
    streamLua.Write(_env, luaTxt)
    streamMeta.Write(_env, metaTxt)
    self.FP.convertToLanguage(_env, ast, streamLua, scriptPath)
    return ast
}
// 1339: decl @lune.@base.@front.Front.saveToGo
func (self *front_Front) saveToGo(_env *LnsEnv, scriptPath string,astResult LnsAny,mod string) LnsAny {
    return NewConverter_GoConverter(_env, scriptPath, astResult, mod, self.option, self.FP.createGoOption(_env, scriptPath))
}
// 1350: decl @lune.@base.@front.Front.saveToPython
func (self *front_Front) saveToPython(_env *LnsEnv, scriptPath string,astResult LnsAny,mod string) LnsAny {
    return NewConverter_PythonConverter(_env, scriptPath, astResult, mod, self.option, self.FP.createPythonOption(_env, scriptPath))
}
// 1362: decl @lune.@base.@front.Front.saveToC
func (self *front_Front) SaveToC(_env *LnsEnv, scriptPath string,ast *AstInfo_ASTInfo) {
    var cPath string
    cPath = front_convExp4_557(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".c")))
    var file Lns_luaStream
    
    {
        _file := front_convExp4_572(Lns_2DDD(Lns_io_open(cPath, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var hPath string
    hPath = front_convExp4_586(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".h")))
    var hFile Lns_luaStream
    
    {
        _hFile := front_convExp4_602(Lns_2DDD(Lns_io_open(hPath, "w")))
        if _hFile == nil{
            return 
        } else {
            hFile = _hFile.(Lns_luaStream)
        }
    }
    file.Close(_env)
    hFile.Close(_env)
}
// 1378: decl @lune.@base.@front.Front.outputBuiltin
func (self *front_Front) OutputBuiltin(_env *LnsEnv, scriptPath string) {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, "lns_builtin", nil)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsCode{"", mod, nil}, baseDir, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    self.FP.SaveToC(_env, scriptPath, ast)
}
// 1400: decl @lune.@base.@front.Front.saveToLua
func (self *front_Front) saveToLua(_env *LnsEnv, updateInfo *front_UpdateInfo,baseDir LnsAny) LnsAny {
    var scriptPath string
    scriptPath = updateInfo.FP.Get_scriptPath(_env)
    var dependsPath LnsAny
    dependsPath = updateInfo.FP.Get_dependsPath(_env)
    var moduleId *FrontInterface_ModuleId
    moduleId = updateInfo.FP.Get_moduleId(_env)
    var uptodate LnsAny
    uptodate = updateInfo.FP.Get_uptodate(_env)
    var mod string
    mod = front_convExp4_716(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath, baseDir)))
    var luaPath string
    luaPath = front_convExp4_730(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".lua")))
    var metaPath string
    metaPath = front_convExp4_743(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".meta")))
    if Lns_isCondTrue( self.option.OutputDir){
        var filename string
        filename = front_convExp4_760(Lns_2DDD(_env.GetVM().String_gsub(mod,"%.", "/")))
        luaPath = _env.GetVM().String_format("%s/%s.lua", []LnsAny{self.option.OutputDir, filename})
        metaPath = _env.GetVM().String_format("%s/%s.meta", []LnsAny{self.option.OutputDir, filename})
    }
    var convMode LnsInt
    convMode = ConvLua_ConvMode__ConvMeta
    if luaPath == scriptPath{
        Util_errorLog(_env, _env.GetVM().String_format("%s is illegal filename.", []LnsAny{luaPath}))
        return nil
    }
    switch _matchExp0 := uptodate.(type) {
    case *front_ModuleUptodate__NeedUpdate:
        var result LnsAny
        result = self.FP.createAstSub(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{scriptPath, mod, nil}, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
        var luaConv LnsAny
        if self.option.NoLua{
            luaConv = nil
        } else { 
            luaConv = NewConverter_LuaConverter(_env, luaPath, metaPath, dependsPath, result, convMode, scriptPath, self.option.ByteCompile, self.option.StripDebugInfo, self.option)
        }
        var goConv LnsAny
        goConv = nil
        var pyConv LnsAny
        pyConv = nil
        if _switch0 := self.option.ConvTo; _switch0 == Types_Lang__Python {
            pyConv = self.FP.saveToPython(_env, scriptPath, result, mod)
        } else if _switch0 == Types_Lang__Go {
            goConv = self.FP.saveToGo(_env, scriptPath, result, mod)
        }
        return Converter_ConverterFunc(func(_env *LnsEnv) {
            _env.NilAccFin(_env.NilAccPush(luaConv) && 
            Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*Converter_LuaConverter).FP.SaveLua(_env)})/* 1471:13 */)
            _env.NilAccFin(_env.NilAccPush(goConv) && 
            Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*Converter_GoConverter).FP.SaveGo(_env)})/* 1472:13 */)
            _env.NilAccFin(_env.NilAccPush(pyConv) && 
            Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*Converter_PythonConverter).FP.SavePython(_env)})/* 1473:13 */)
        })
    case *front_ModuleUptodate__Uptodate:
        metaInfo := _matchExp0.Val1
        Util_errorLog(_env, "uptodate -- " + scriptPath)
        var dependsStream LnsAny
        dependsStream = self.option.FP.OpenDepend(_env, dependsPath)
        front_outputDependInfo_15_(_env, dependsStream, metaInfo, mod)
        Converter_closeStreams(_env, nil, nil, dependsStream, metaPath, false)
        return nil
    case *front_ModuleUptodate__NeedTouch:
        metaCode := _matchExp0.Val1
        metaInfo := _matchExp0.Val2
        Util_errorLog(_env, "touch -- " + scriptPath)
        var dependsStream LnsAny
        dependsStream = self.option.FP.OpenDepend(_env, dependsPath)
        var metaMemStream *Util_memStream
        metaMemStream = NewUtil_memStream(_env)
        if self.option.Mode == Option_ModeKind__SaveMeta{
            metaMemStream.FP.Write(_env, metaCode)
        }
        front_outputDependInfo_15_(_env, dependsStream, metaInfo, mod)
        Converter_closeStreams(_env, nil, metaMemStream, dependsStream, metaPath, self.option.Mode == Option_ModeKind__SaveMeta)
        return nil
    }
// insert a dummy
    return nil
}
// 1503: decl @lune.@base.@front.Front.outputBootC
func (self *front_Front) outputBootC(_env *LnsEnv, scriptPath string) {
}
// 1588: decl @lune.@base.@front.Front.build
func (self *front_Front) Build(_env *LnsEnv, buildMode LnsAny,astCallback LnsAny) {
    var front_createUpdateInfo func(_env *LnsEnv, scriptPath string,dependsPath LnsAny,baseDir LnsAny) *front_UpdateInfo
    front_createUpdateInfo = func(_env *LnsEnv, scriptPath string,dependsPath LnsAny,baseDir LnsAny) *front_UpdateInfo {
        var mod string
        mod = front_convExp0_1648(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath, baseDir)))
        var moduleId *FrontInterface_ModuleId
        var uptodate LnsAny
        moduleId,uptodate = self.FP.getModuleIdAndCheckUptodate(_env, scriptPath, mod)
        return Newfront_UpdateInfo(_env, scriptPath, dependsPath, moduleId, uptodate)
    }
    var front_process func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo,baseDir LnsAny) LnsAny
    front_process = func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo,baseDir LnsAny) LnsAny {
        var mod string
        mod = front_convExp0_1706(Lns_2DDD(self.FP.scriptPath2Module(_env, updateInfo.FP.Get_scriptPath(_env), baseDir)))
        switch _matchExp0 := buildMode.(type) {
        case *Front_BuildMode__Save:
            return self.FP.saveToLua(_env, updateInfo, baseDir)
        case *Front_BuildMode__Output:
            streamLua := _matchExp0.Val1
            streamMeta := _matchExp0.Val2
            self.FP.convertToLua(_env, updateInfo.FP.Get_scriptPath(_env), baseDir, ConvLua_ConvMode__ConvMeta, streamLua, streamMeta)
        case *Front_BuildMode__CreateAst:
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(self.mod2astCreate.Get(mod))) &&
                _env.SetStackVal( Lns_op_not(self.moduleMgr.FP.GetAst(_env, mod))) ).(bool)){
                var result LnsAny
                result = self.FP.createAstSub(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{updateInfo.FP.Get_scriptPath(_env), mod, nil}, baseDir, mod, updateInfo.FP.Get_moduleId(_env), nil, TransUnit_AnalyzeMode__Compile, nil)
                return Converter_ConverterFunc(func(_env *LnsEnv) {
                    self.FP.applyAstResult(_env, result)
                })
            }
        }
        return nil
    }
    Depend_profile(_env, self.option.ValidProf, conv2Form0_2180(func(_env *LnsEnv) {
        __func__ := "@lune.@base.@front.Front.build.<anonymous>"
        if self.option.ScriptPath == "@-"{
            var baseDir LnsAny
            baseDir = self.option.FP.Get_projDir(_env)
            for _, _path := range( self.option.BatchList.Items ) {
                path := _path.(string)
                self.targetSet.Add(Lns_car(self.FP.scriptPath2Module(_env, path, baseDir)).(string))
            }
            var postProcessMap *LnsMap
            postProcessMap = NewLnsMap( map[LnsAny]LnsAny{})
            for _index, _path := range( self.option.BatchList.Items ) {
                index := _index + 1
                path := _path.(string)
                var updateInfo *front_UpdateInfo
                updateInfo = front_createUpdateInfo(_env, path, Lns_car(_env.GetVM().String_gsub(path,".lns$", ".d")).(string), baseDir)
                Util_println(_env, []LnsAny{_env.GetVM().String_format("%s: start...", []LnsAny{updateInfo.FP.Get_scriptPath(_env)})})
                {
                    __exp := front_process(_env, false, updateInfo, nil)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(Converter_ConverterFunc)
                        if self.option.FP.Get_validPostBuild(_env){
                            postProcessMap.Set(index,_exp)
                        } else { 
                            _exp(_env)
                        }
                    }
                }
            }
            {
                __forsortCollection0 := postProcessMap
                __forsortSorted0 := __forsortCollection0.CreateKeyListInt()
                __forsortSorted0.Sort( _env, LnsItemKindInt, nil )
                for _, _index := range( __forsortSorted0.Items ) {
                    postProcess := __forsortCollection0.Items[ _index ].(Converter_ConverterFunc)
                    index := _index.(LnsInt)
                    var prev LnsReal
                    prev = _env.GetVM().OS_clock()
                    var path string
                    path = self.option.BatchList.GetAt(index).(string)
                    Util_println(_env, []LnsAny{_env.GetVM().String_format("%s: waiting...", []LnsAny{path})})
                    postProcess(_env)
                    Util_println(_env, []LnsAny{_env.GetVM().String_format("%s: done %g msec", []LnsAny{path, (_env.GetVM().OS_clock() - prev) * LnsReal(1000)})})
                }
            }
        } else { 
            var baseDir LnsAny
            _,baseDir = front_getBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env))
            {
                _postProcess := front_process(_env, true, front_createUpdateInfo(_env, self.option.ScriptPath, nil, baseDir), baseDir)
                if !Lns_IsNil( _postProcess ) {
                    postProcess := _postProcess.(Converter_ConverterFunc)
                    postProcess(_env)
                }
            }
        }
        if astCallback != nil{
            astCallback_580 := astCallback.(Front_AstCallback)
            for _, _mod := range( self.moduleMgr.FP.GetModList(_env).Items ) {
                mod := _mod.(string)
                {
                    __exp := self.moduleMgr.FP.GetAst(_env, mod)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*AstInfo_ASTInfo)
                        astCallback_580(_env, _exp)
                    } else {
                        Log_log(_env, Log_Level__Err, __func__, 1668, Log_CreateMessage(func(_env *LnsEnv) string {
                            return _env.GetVM().String_format("not found AST -- %s", []LnsAny{mod})
                        }))
                        
                    }
                }
            }
        }
    }), self.option.ScriptPath + ".profi")
}
// 1690: decl @lune.@base.@front.Front.executeLns
func (self *front_Front) executeLns(_env *LnsEnv, path string,baseDir LnsAny) {
    __func__ := "@lune.@base.@front.Front.executeLns"
    var mod string
    mod = Util_scriptPath2ModuleFromProjDir(_env, path, baseDir)
    {
        var parserSrc LnsAny
        parserSrc = &Types_ParserSrc__LnsPath{path, mod, nil}
        var luaCode string
        _,luaCode = self.FP.loadParserToLuaCode(_env, NewFrontInterface_ImportModuleInfo(_env), parserSrc, path, mod, baseDir)
        Log_log(_env, Log_Level__Debug, __func__, 1701, Log_CreateMessage(func(_env *LnsEnv) string {
            return "luacode: " + luaCode
        }))
        
        Lns_LockEnvSync( _env, 1703, func () {
            var subModPreLoad string
            subModPreLoad = "return function( submod2Code, dumpCode )\n   local preloadFunc = function( mod )\n      code = submod2Code[ mod ]\n      local loadFunc = loadstring or load -- lua5.1 and lua5.2\n      local loaded, mess = loadFunc( code )\n      if not loaded then\n         error( mess )\n      end\n      return loaded()\n   end\n   for mod, code in pairs( submod2Code ) do\n      if dumpCode then\n         print( string.format( \"mod: %s %s\", mod, code ) )\n      end\n      package.preload[ mod ] = preloadFunc\n   end\nend\n"
            var loaded *Lns_luaValue
            
            {
                _loaded := front_convExp0_2341(Lns_2DDD(_env.GetVM().Load(subModPreLoad, nil)))
                if _loaded == nil{
                    panic("failed to subModPreLoad")
                } else {
                    loaded = _loaded.(*Lns_luaValue)
                }
            }
            var preloadFunc LnsAny
            
            {
                _preloadFunc := front_convExp0_2358(Lns_2DDD(_env.GetVM().RunLoadedfunc(loaded,Lns_2DDD([]LnsAny{}))[0]))
                if _preloadFunc == nil{
                    panic("failed to preloadFunc")
                } else {
                    preloadFunc = _preloadFunc
                }
            }
            _env.GetVM().RunLoadedfunc((preloadFunc.(*Lns_luaValue)),Lns_2DDD([]LnsAny{self.convertedMap, Log_getLevel(_env) >= Log_Level__Debug}))
        })
        Front_loadFromLuaTxt(_env, luaCode)
    }
    if self.option.Testing{
        var code string
        code = "local Testing = require( \"lune.base.Testing\" )\nreturn function( path )\n  Testing.run( path );\n  Testing.outputAllResult( io.stdout );\nend\n"
        Lns_LockEnvSync( _env, 1747, func () {
            var loaded LnsAny
            var mess LnsAny
            loaded,mess = _env.GetVM().Load(code, nil)
            if loaded != nil{
                loaded_614 := loaded.(*Lns_luaValue)
                {
                    _testFunc := front_convExp0_2442(Lns_2DDD(_env.GetVM().RunLoadedfunc(loaded_614,Lns_2DDD([]LnsAny{}))[0]))
                    if !Lns_IsNil( _testFunc ) {
                        testFunc := _testFunc
                        _env.GetVM().RunLoadedfunc((testFunc.(*Lns_luaValue)),Lns_2DDD([]LnsAny{mod}))
                    }
                }
            } else {
                Util_println(_env, []LnsAny{mess})
            }
        })
    }
}
// 1760: decl @lune.@base.@front.Front.exec
func (self *front_Front) Exec(_env *LnsEnv) {
    __func__ := "@lune.@base.@front.Front.exec"
    Log_log(_env, Log_Level__Trace, __func__, 1762, Log_CreateMessage(func(_env *LnsEnv) string {
        return Option_ModeKind_getTxt( self.option.Mode)
    }))
    
    if _switch0 := self.option.Mode; _switch0 == Option_ModeKind__Token {
        self.FP.DumpTokenize(front_convExp0_2536(_env, Lns_2DDD(front_getBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Ast {
        self.FP.DumpAst(front_convExp0_2562(_env, Lns_2DDD(front_getParseSrcAndBaseDir_20_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Format {
        self.FP.Format(front_convExp0_2588(_env, Lns_2DDD(front_getParseSrcAndBaseDir_20_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Diag {
        self.FP.CheckDiag(front_convExp0_2614(_env, Lns_2DDD(front_getParseSrcAndBaseDir_20_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Complete {
        self.FP.Complete(front_convExp0_2640(_env, Lns_2DDD(front_getParseSrcAndBaseDir_20_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Inquire {
        self.FP.Inquire(front_convExp0_2666(_env, Lns_2DDD(front_getParseSrcAndBaseDir_20_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Glue {
        self.FP.CreateGlue(front_convExp0_2692(_env, Lns_2DDD(front_getParseSrcAndBaseDir_20_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Lua || _switch0 == Option_ModeKind__LuaMeta {
        var convMode LnsInt
        if self.option.Mode == Option_ModeKind__Lua{
            convMode = ConvLua_ConvMode__Convert
        } else { 
            convMode = ConvLua_ConvMode__ConvMeta
        }
        var scriptPath string
        var baseDir LnsAny
        scriptPath,baseDir = front_getBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env))
        self.FP.convertToLua(_env, scriptPath, baseDir, convMode, Lns_io_stdout, Lns_io_stdout)
    } else if _switch0 == Option_ModeKind__Save || _switch0 == Option_ModeKind__SaveMeta {
        self.FP.Build(_env, Front_BuildMode__Save_Obj, nil)
    } else if _switch0 == Option_ModeKind__BuildAst {
        self.FP.Build(_env, Front_BuildMode__CreateAst_Obj, Front_AstCallback(front_Front_exec___anonymous_1_))
    } else if _switch0 == Option_ModeKind__Shebang {
        Depend_setupShebang(_env)
        Lns_LockEnvSync( _env, 1817, func () {
            var scriptPath string
            var baseDir LnsAny
            scriptPath,baseDir = front_getBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env))
            {
                _modObj := front_convExp0_2881(Lns_2DDD(self.FP.loadModuleWithBaseDir(_env, Lns_car(self.FP.scriptPath2Module(_env, scriptPath, baseDir)).(string), baseDir)))
                if !Lns_IsNil( _modObj ) {
                    modObj := _modObj
                    var code LnsInt
                    code = Depend_runMain(_env, modObj.(*Lns_luaValue).GetAt("__main"), self.option.ShebangArgList)
                    _env.GetVM().OS_exit(code)
                }
            }
        })
    } else if _switch0 == Option_ModeKind__GoMod {
        for _, _path := range( self.gomodMap.FP.GetModPathList(_env).Items ) {
            path := _path.(string)
            Util_println(_env, []LnsAny{path})
        }
    } else if _switch0 == Option_ModeKind__Exec {
        self.FP.executeLns(front_convExp0_2928(_env, Lns_2DDD(front_getBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__BootC {
        self.FP.outputBootC(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Builtin {
        self.FP.OutputBuiltin(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__MkMain {
        var mod string
        mod = front_convExp0_2978(Lns_2DDD(self.FP.scriptPath2Module(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env))))
        {
            _mess := ConvGo_outputGoMain(_env, self.FP.getGoAppName(_env), mod, self.option.Testing, self.option.OutputDir, self.option.FP.Get_runtimeOpt(_env))
            if !Lns_IsNil( _mess ) {
                mess := _mess.(string)
                Util_errorLog(_env, mess)
            }
        }
    } else if _switch0 == Option_ModeKind__Indexer {
        self.FP.Build(_env, Front_BuildMode__CreateAst_Obj, Front_AstCallback(front_Front_exec___anonymous_2_))
    } else {
        Util_println(_env, []LnsAny{"illegal mode"})
    }
}
// 460: decl @lune.@base.@front.MetaForBuildId.createModuleId
func (self *front_MetaForBuildId) CreateModuleId(_env *LnsEnv) *FrontInterface_ModuleId {
    return FrontInterface_ModuleId_createIdFromTxt(_env, self.G__buildId)
}
// 465: decl @lune.@base.@front.MetaForBuildId.LoadFromMeta
func front_MetaForBuildId_LoadFromMeta_5_(_env *LnsEnv, metaPath string)(LnsAny, LnsAny) {
    {
        _fileObj := front_convExp1_1146(Lns_2DDD(Lns_io_open(metaPath, nil)))
        if !Lns_IsNil( _fileObj ) {
            fileObj := _fileObj.(Lns_luaStream)
            var luaCode LnsAny
            luaCode = fileObj.Read(_env, "*a")
            fileObj.Close(_env)
            if luaCode != nil{
                luaCode_93 := luaCode.(string)
                var meta LnsAny
                    meta = front_convExp1_1133(Lns_2DDD(front_MetaForBuildId__fromStem_4_(_env, _env.GetVM().ExpandLuavalMap(Front_loadFromLuaTxt(_env, luaCode_93)),nil)))
                return meta, luaCode_93
            }
        }
    }
    return nil, nil
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
    Add(_env *LnsEnv, arg1 *AstInfo_ASTInfo, arg2 *FrontInterface_ModuleInfo)
    AddMeta(_env *LnsEnv, arg1 string, arg2 *FrontInterface_ModuleMeta)
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
// 89: DeclConstr
func (self *front_ModuleMgr) Initfront_ModuleMgr(_env *LnsEnv) {
    self.mod2info = NewUtil_OrderdMap(_env)
    self.loadedMetaMap = NewLnsMap( map[LnsAny]LnsAny{})
}


// declaration Class -- Front
type front_FrontMtd interface {
    applyAstResult(_env *LnsEnv, arg1 LnsAny) *AstInfo_ASTInfo
    Build(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny)
    CheckDiag(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    checkUptodateMeta(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny, arg4 LnsAny) LnsAny
    Complete(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    convertFromAst(_env *LnsEnv, arg1 *AstInfo_ASTInfo, arg2 string, arg3 LnsInt)(string, string)
    convertLns2LuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsInt, arg3 LnsAny, arg4 LnsAny, arg5 Lns_iStream, arg6 string) string
    convertToLanguage(_env *LnsEnv, arg1 *AstInfo_ASTInfo, arg2 Lns_oStream, arg3 string)
    convertToLua(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsInt, arg4 Lns_oStream, arg5 Lns_oStream) LnsAny
    createAst(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 string, arg5 *FrontInterface_ModuleId, arg6 LnsAny, arg7 LnsInt, arg8 LnsAny) *AstInfo_ASTInfo
    createAstSub(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 string, arg5 *FrontInterface_ModuleId, arg6 LnsAny, arg7 LnsInt, arg8 LnsAny) LnsAny
    CreateGlue(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    createGoOption(_env *LnsEnv, arg1 string) *ConvGo_Option
    createPaser(_env *LnsEnv, arg1 string, arg2 LnsAny) *Parser_Parser
    createPythonOption(_env *LnsEnv, arg1 string) *ConvPython_Option
    DumpAst(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    DumpTokenize(_env *LnsEnv, arg1 string, arg2 LnsAny)
    Error(_env *LnsEnv, arg1 string)
    Exec(_env *LnsEnv)
    executeLns(_env *LnsEnv, arg1 string, arg2 LnsAny)
    Format(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    getGoAppName(_env *LnsEnv) string
    getLoadInfo(_env *LnsEnv, arg1 string) LnsAny
    GetLuaModulePath(_env *LnsEnv, arg1 string, arg2 LnsAny)(string, LnsAny, string)
    getModuleIdAndCheckUptodate(_env *LnsEnv, arg1 string, arg2 string)(*FrontInterface_ModuleId, LnsAny)
    Inquire(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    loadFile(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string)(*FrontInterface_ModuleMeta, LnsAny)
    LoadFromLnsTxt(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string) LnsAny
    loadLua(_env *LnsEnv, arg1 string) LnsAny
    LoadMeta(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string, arg4 LnsAny, arg5 FrontInterface_ModuleLoader) LnsAny
    LoadModule(_env *LnsEnv, arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
    loadModuleWithBaseDir(_env *LnsEnv, arg1 string, arg2 LnsAny)(LnsAny, *FrontInterface_ModuleMeta)
    loadParserToLuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny)(*FrontInterface_ModuleMeta, string)
    outputBootC(_env *LnsEnv, arg1 string)
    OutputBuiltin(_env *LnsEnv, arg1 string)
    regConvertedMap(_env *LnsEnv, arg1 string, arg2 string, arg3 *FrontInterface_ModuleMeta)
    SaveToC(_env *LnsEnv, arg1 string, arg2 *AstInfo_ASTInfo)
    saveToGo(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 string) LnsAny
    saveToLua(_env *LnsEnv, arg1 *front_UpdateInfo, arg2 LnsAny) LnsAny
    saveToPython(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 string) LnsAny
    scriptPath2Module(_env *LnsEnv, arg1 string, arg2 LnsAny)(string, LnsAny)
    searchLuaFile(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny) LnsAny
    SearchModule(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny) LnsAny
    searchModuleFile(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny, arg4 LnsAny) LnsAny
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
    mod2astCreate *LnsMap
    mod2loader *LnsMap
    preloadedModMap *LnsMap
    loadCount LnsInt
    builtinFunc *Builtin_BuiltinFuncType
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
// 189: DeclConstr
func (self *front_Front) Initfront_Front(_env *LnsEnv, option *Option_Option,bindModuleList LnsAny) {
    self.mod2loader = NewLnsMap( map[LnsAny]LnsAny{})
    self.mod2astCreate = NewLnsMap( map[LnsAny]LnsAny{})
    self.loadCount = 0
    self.targetSet = NewLnsSet([]LnsAny{})
    self.bindModuleSet = NewLnsSet([]LnsAny{})
    if bindModuleList != nil{
        bindModuleList_107 := bindModuleList.(*LnsList)
        for _, _mod := range( bindModuleList_107.Items ) {
            mod := _mod.(string)
            self.bindModuleSet.Add(mod)
        }
    }
    self.moduleMgr = Newfront_ModuleMgr(_env)
    self.gomodMap = GoMod_getGoMap(_env)
    DependLuaOnLns_addGoModPath(_env, self.gomodMap.FP.GetModPathList(_env))
    self.option = option
    self.loadedMap = NewLnsMap( map[LnsAny]LnsAny{})
    self.loadedMapTest = NewLnsMap( map[LnsAny]LnsAny{})
    self.convertedMap = NewLnsMap( map[LnsAny]LnsAny{})
    {
        var builtin *Builtin_Builtin
        builtin = NewBuiltin_Builtin(_env, self.option.TargetLuaVer, option.TransCtrlInfo)
        self.builtinFunc = builtin.FP.RegistBuiltInScope(_env)
    }
    FrontInterface_setFront(_env, self.FP)
    var loadedMap *LnsMap
    loadedMap = NewLnsMap( map[LnsAny]LnsAny{})
        {
            _foreachExp1 := Depend_getLoadedMod(_env)
            _foreachKey1, _foreachVal1 := _foreachExp1.Get1stFromMap()
            for _foreachKey1 != nil {
                mod := _foreachKey1.(string)
                modval := _foreachVal1
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
                _foreachKey1, _foreachVal1 = _foreachExp1.NextFromMap( _foreachKey1 )
            }
        }
    self.preloadedModMap = loadedMap
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
func front_DependMetaInfo__fromMap_2_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_DependMetaInfo_FromMap( arg1, paramList )
}
func front_DependMetaInfo__fromStem_3_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
func front_MetaForBuildId__fromMap_3_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_MetaForBuildId_FromMap( arg1, paramList )
}
func front_MetaForBuildId__fromStem_4_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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

// declaration Class -- UpdateInfo
type front_UpdateInfoMtd interface {
    Get_dependsPath(_env *LnsEnv) LnsAny
    Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId
    Get_scriptPath(_env *LnsEnv) string
    Get_uptodate(_env *LnsEnv) LnsAny
}
type front_UpdateInfo struct {
    scriptPath string
    dependsPath LnsAny
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
func Newfront_UpdateInfo(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 *FrontInterface_ModuleId, arg4 LnsAny) *front_UpdateInfo {
    obj := &front_UpdateInfo{}
    obj.FP = obj
    obj.Initfront_UpdateInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *front_UpdateInfo) Initfront_UpdateInfo(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 *FrontInterface_ModuleId, arg4 LnsAny) {
    self.scriptPath = arg1
    self.dependsPath = arg2
    self.moduleId = arg3
    self.uptodate = arg4
}
func (self *front_UpdateInfo) Get_scriptPath(_env *LnsEnv) string{ return self.scriptPath }
func (self *front_UpdateInfo) Get_dependsPath(_env *LnsEnv) LnsAny{ return self.dependsPath }
func (self *front_UpdateInfo) Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId{ return self.moduleId }
func (self *front_UpdateInfo) Get_uptodate(_env *LnsEnv) LnsAny{ return self.uptodate }

func Lns_front_init(_env *LnsEnv) {
    if init_front { return }
    init_front = true
    front__mod__ = "@lune.@base.@front"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_frontInterface_init(_env)
    Lns_Parser_init(_env)
    Lns_convLua_init(_env)
    Lns_convGo_init(_env)
    Lns_convPython_init(_env)
    Lns_AstInfo_init(_env)
    Lns_TransUnit_init(_env)
    Lns_Util_init(_env)
    Lns_Option_init(_env)
    Lns_dumpNode_init(_env)
    Lns_glueFilter_init(_env)
    Lns_LuaVer_init( _env )
    Lns_Depend_init(_env)
    Lns_DependLuaOnLns_init(_env)
    Lns_OutputDepend_init(_env)
    Lns_Ver_init(_env)
    Lns_LuaVer_init(_env)
    Lns_Log_init(_env)
    Lns_Formatter_init(_env)
    Lns_Testing_init(_env)
    Lns_GoMod_init(_env)
    Lns_Meta_init(_env)
    Lns_Nodes_init(_env)
    Lns_Builtin_init(_env)
    Lns_NodeIndexer_init(_env)
    Lns_Converter_init(_env)
    Depend_setup(_env, Depend_UpdateVer(front__anonymous_0_))
    front_forceUpdateMeta = true
}
func init() {
    init_front = false
}
