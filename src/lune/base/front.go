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
// for 1657: ExpCast
func conv2Form0_2211( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1763: ExpCast
func conv2Form0_2310( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 29: ExpCast
func conv2Form0_2385( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 29: ExpCast
func conv2Form0_2460( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1802: ExpCast
func conv2Form0_2655( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1267: ExpCast
func conv2Form4_131( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 402
func front_convExp1_863(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 497
func front_convExp1_1245(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 849
func front_convExp2_1393(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1801
func front_convExp0_2666(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 419
func front_convExp1_900(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 481
func front_convExp1_1156(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 496
func front_convExp1_1224(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 526
func front_convExp1_1393(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 527
func front_convExp1_1412(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 568
func front_convExp2_41(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 750
func front_convExp2_837(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 779
func front_convExp2_989(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 833
func front_convExp2_1297(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1153
func front_convExp3_1495(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1818
func front_convExp0_2760(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, string, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1821
func front_convExp0_2786(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1825
func front_convExp0_2812(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1829
func front_convExp0_2838(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1833
func front_convExp0_2864(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1837
func front_convExp0_2890(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1841
func front_convExp0_2916(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 3 )
}
// for 1870
func front_convExp0_3064(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, string, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1897
func front_convExp0_3146(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, string, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 278
func front_convExp1_285(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 302
func front_convExp1_399(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 331
func front_convExp1_474(arg1 []LnsAny) *AstInfo_ASTInfo {
    return Lns_getFromMulti( arg1, 0 ).(*AstInfo_ASTInfo)
}
// for 349
func front_convExp1_566(arg1 []LnsAny) (*AstInfo_ASTInfo, *FrontInterface_ModuleInfo, *FrontInterface_ModuleMeta) {
    return Lns_getFromMulti( arg1, 0 ).(*AstInfo_ASTInfo), Lns_getFromMulti( arg1, 1 ).(*FrontInterface_ModuleInfo), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleMeta)
}
// for 454
func front_convExp1_1083(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 648
func front_convExp2_310(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 675
func front_convExp2_461(arg1 []LnsAny) (LnsAny, string, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 714
func front_convExp2_650(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 719
func front_convExp2_695(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 741
func front_convExp2_779(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 770
func front_convExp2_931(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 820
func front_convExp2_1175(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 905
func front_convExp3_85(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 972
func front_convExp3_481(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1015
func front_convExp3_784(arg1 []LnsAny) (string, LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1032
func front_convExp3_883(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1041
func front_convExp3_943(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1063
func front_convExp3_1061(arg1 []LnsAny) (*FrontInterface_ModuleMeta, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 )
}
// for 1150
func front_convExp3_1466(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1180
func front_convExp3_1603(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1217
func front_convExp3_1955(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1228
func front_convExp3_2022(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1352
func front_convExp4_432(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1359
func front_convExp4_496(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1395
func front_convExp4_594(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1397
func front_convExp4_609(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1401
func front_convExp4_623(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1402
func front_convExp4_639(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1411
func front_convExp4_664(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1441
func front_convExp4_753(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1442
func front_convExp4_767(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1443
func front_convExp4_780(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1445
func front_convExp4_797(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1607
func front_convExp0_1610(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1622
func front_convExp0_1679(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1623
func front_convExp0_1692(arg1 []LnsAny) (*FrontInterface_ModuleId, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 1 )
}
// for 1630
func front_convExp0_1737(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1686
func front_convExp0_2088(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1757
func front_convExp0_2286(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1760
func front_convExp0_2304(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1776
func front_convExp0_2540(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1788
func front_convExp0_2619(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1799
func front_convExp0_2638(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1852
func front_convExp0_2968(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1869
func front_convExp0_3077(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1906
func front_convExp0_3196(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 23
func front_convExp0_2361(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 23
func front_convExp0_2436(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 26
func front_convExp0_2379(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 26
func front_convExp0_2454(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
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
// 289: decl @lune.@base.@front.createPaser
func front_createPaser_6_(_env *LnsEnv, baseDir LnsAny,path string,mod string,stdinFile LnsAny) *Tokenizer_Tokenizer {
    return &Tokenizer_StreamTokenizer_create(_env, &Types_TokenizerSrc__LnsPath{baseDir, path, mod, nil}, false, stdinFile, nil).Tokenizer_Tokenizer
}

// 386: decl @lune.@base.@front.ast2Lua
func front_ast2Lua_7_(_env *LnsEnv, ast *AstInfo_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,option *Option_Option) {
    var conv *ConvLua_FilterInfo
    conv = Converter_ast2LuaMain(_env, ast, streamName, stream, metaStream, convMode, inMacro, option)
    conv.FP.OutputLuaAndMeta(_env, Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode))
}

// 395: decl @lune.@base.@front.loadFromChunk
func front_loadFromChunk_8_(_env *LnsEnv, chunk LnsAny,err LnsAny) LnsAny {
    if err != nil{
        err_66 := err.(string)
        Util_errorLog(_env, err_66)
    }
    var ret LnsAny
    Lns_LockEnvSync( _env, 400, func () {
        if chunk != nil{
            chunk_70 := chunk.(*Lns_luaValue)
            {
                _work := front_convExp1_863(Lns_2DDD(_env.GetVM().RunLoadedfunc(chunk_70,Lns_2DDD([]LnsAny{}))))
                if !Lns_IsNil( _work ) {
                    work := _work
                    ret = work
                } else {
                    ret = nil
                }
            }
        } else {
            Util_err(_env, "failed to error")
        }
    })
    return ret
}

// 415: decl @lune.@base.@front.loadFromLuaTxt
func Front_loadFromLuaTxt(_env *LnsEnv, txt string) LnsAny {
    var chunk LnsAny
    var err LnsAny
    Lns_LockEnvSync( _env, 418, func () {
        chunk, err = _env.GetVM().Load(txt, nil)
    })
    return front_loadFromChunk_8_(_env, chunk, err)
}

// 491: decl @lune.@base.@front.getMetaInfo
func front_getMetaInfo_12_(_env *LnsEnv, lnsPath string,mod string,outdir LnsAny)(LnsAny, string, string) {
    var moduleMetaPath string
    moduleMetaPath = lnsPath
    if outdir != nil{
        outdir_99 := outdir.(string)
        moduleMetaPath = _env.GetVM().String_format("%s/%s", Lns_2DDD(outdir_99, Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string)))
    }
    moduleMetaPath = front_convExp1_1224(Lns_2DDD(_env.GetVM().String_gsub(moduleMetaPath,"%.lns$", ".meta")))
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

// 560: decl @lune.@base.@front.getModuleId
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




// 1335: decl @lune.@base.@front.outputDependInfo
func front_outputDependInfo_15_(_env *LnsEnv, stream LnsAny,metaInfo *front_MetaForBuildId,mod string) {
    if stream != nil{
        stream_21 := stream.(Lns_oStream)
        var dependInfo *OutputDepend_DependInfo
        dependInfo = NewOutputDepend_DependInfo(_env, mod)
        for _dependMod, _ := range( metaInfo.G__dependModuleMap.Items ) {
            dependMod := _dependMod
            dependInfo.FP.AddImpotModule(_env, dependMod)
        }
        for _, _subMod := range( metaInfo.G__subModuleMap.Items ) {
            subMod := _subMod
            dependInfo.FP.AddSubMod(_env, subMod)
        }
        dependInfo.FP.Output(_env, stream_21)
    }
}


// 1562: decl @lune.@base.@front.convertLnsCode2LuaCode
func Front_convertLnsCode2LuaCode(_env *LnsEnv, lnsCode string,path string,baseDir LnsAny) string {
    var option *Option_Option
    option = NewOption_Option(_env)
    option.ScriptPath = path
    option.UseLuneModule = Option_getRuntimeModule(_env)
    option.UseIpairs = true
    var front *Front_Front
    front = NewFront_Front(_env, option, nil)
    return front.FP.ConvertLnsCode2LuaCodeWithOpt(_env, lnsCode, path, baseDir)
}

// 1573: decl @lune.@base.@front.getBaseDir
func front_getBaseDir_18_(_env *LnsEnv, path string,userProjDir LnsAny)(string, LnsAny) {
    var projDir string
    if userProjDir != nil{
        userProjDir_89 := userProjDir.(string)
        projDir = userProjDir_89
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

// 1604: decl @lune.@base.@front.getParseSrcAndBaseDir
func front_getParseSrcAndBaseDir_19_(_env *LnsEnv, path string,userProjDir LnsAny)(LnsAny, string, *FrontInterface_ModuleId, LnsAny) {
    var workPath string
    var projDir LnsAny
    workPath,projDir = front_getBaseDir_18_(_env, path, userProjDir)
    var mod string
    mod = Util_scriptPath2ModuleFromProjDir(_env, workPath, projDir)
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_13_(_env, path, mod, nil, nil)
    return &Types_TokenizerSrc__LnsPath{projDir, workPath, mod, nil}, mod, moduleId, projDir
}





// 1706: decl @lune.@base.@front.buildWithBuildMode
func Front_buildWithBuildMode(_env *LnsEnv, option *Option_Option,buildMode LnsAny,astCallback LnsAny) {
    var front *Front_Front
    Lns_LockEnvSync( _env, 1710, func () {
        front = NewFront_Front(_env, option, nil)
    })
    front.FP.Build(_env, buildMode, astCallback)
}

// 1716: decl @lune.@base.@front.build
func Front_build(_env *LnsEnv, option *Option_Option,astCallback LnsAny) {
    Front_buildWithBuildMode(_env, option, Front_BuildMode__CreateAst_Obj, astCallback)
}

func front_Front_exec___anonymous_1_(_env *LnsEnv, ast *AstInfo_ASTInfo) {
    Util_println(_env, Lns_2DDD(ast.FP.Get_streamName(_env)))
}
func front_Front_exec___anonymous_2_(_env *LnsEnv, ast *AstInfo_ASTInfo) {
    var indexer *NodeIndexer_Indexer
    indexer = NewNodeIndexer_Indexer(_env, ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env))
    indexer.FP.Start(_env, ast.FP.Get_node(_env), NewLnsSet2_[LnsInt](Lns_2DDDGen[LnsInt](Nodes_NodeKind_get_Switch(_env), Nodes_NodeKind_get_Match(_env), Nodes_NodeKind_get_For(_env), Nodes_NodeKind_get_Apply(_env))))
    indexer.FP.Dump(_env)
}
// 1936: decl @lune.@base.@front.exec
func Front_exec(_env *LnsEnv, args *LnsList2_[string]) {
    var version LnsReal
    version = Lns_unwrapDefault( Lns_tonumber(Lns_car(_env.GetVM().String_gsub(Depend_getLuaVersion(_env),"^[^%d]+", "")).(string), nil), 0.0).(LnsReal)
    if version < 5.1{
        Util_errorLog(_env, _env.GetVM().String_format("LuneScript doesn't support this lua version(%s). %s\n", Lns_2DDD(version, "please use the version >= 5.1.")))
        _env.GetVM().OS_exit(1)
    }
    var option *Option_Option
    option = Option_analyze(_env, args)
    var front *Front_Front
    front = NewFront_Front(_env, option, nil)
    front.FP.Exec(_env)
}

// 1953: decl @lune.@base.@front.__main
func Front___main(_env *LnsEnv, argList *LnsList) LnsInt {
    Lns_front_init( _env )
    var list *LnsList2_[string]
    list = NewLnsList2_[string]([]string{})
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




// 96: decl @lune.@base.@front.ModuleMgr.get
func (self *front_ModuleMgr) Get(_env *LnsEnv, mod string) LnsAny {
    return self.mod2info.FP.Get_map(_env).Get(mod)
}
// 100: decl @lune.@base.@front.ModuleMgr.getAst
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
// 114: decl @lune.@base.@front.ModuleMgr.getModList
func (self *front_ModuleMgr) GetModList(_env *LnsEnv) *LnsList {
    return self.mod2info.FP.Get_keyList(_env)
}
// 119: decl @lune.@base.@front.ModuleMgr.add
func (self *front_ModuleMgr) Add(_env *LnsEnv, ast *AstInfo_ASTInfo,moduleInfo *FrontInterface_ModuleInfo) {
    var mod string
    mod = moduleInfo.FP.Get_fullName(_env)
    self.mod2info.FP.Add(_env, mod, &front_UptodateInfo__Update{moduleInfo, ast}, true)
}
// 138: decl @lune.@base.@front.ModuleMgr.addMeta
func (self *front_ModuleMgr) AddMeta(_env *LnsEnv, mod string,meta *FrontInterface_ModuleMeta) {
    __func__ := "@lune.@base.@front.ModuleMgr.addMeta"
    if Lns_op_not(self.FP.Get(_env, mod)){
        Log_log(_env, Log_Level__Log, __func__, 140, Log_CreateMessage(func(_env *LnsEnv) string {
            return mod
        }))
        
        self.mod2info.FP.Add(_env, mod, &front_UptodateInfo__Uptodate{meta}, false)
    }
    self.loadedMetaMap.Set(mod,meta)
}
// 146: decl @lune.@base.@front.ModuleMgr.getMeta
func (self *front_ModuleMgr) GetMeta(_env *LnsEnv, mod string) LnsAny {
    return self.loadedMetaMap.Get(mod)
}
// 186: decl @lune.@base.@front.Front.regConvertedMap
func (self *Front_Front) regConvertedMap(_env *LnsEnv, mod string,luaTxt string,meta *FrontInterface_ModuleMeta) {
    self.moduleMgr.FP.AddMeta(_env, mod, meta)
    self.convertedMap.Set(mod,luaTxt)
}
// 242: decl @lune.@base.@front.Front.getLoadInfo
func (self *Front_Front) getLoadInfo(_env *LnsEnv, mod string) LnsAny {
    if self.option.Testing{
        return self.loadedMapTest.Get(mod)
    }
    return self.loadedMap.Get(mod)
}
// 249: decl @lune.@base.@front.Front.setLoadInfo
func (self *Front_Front) setLoadInfo(_env *LnsEnv, mod string,info *front_LoadInfo) {
    if self.option.Testing{
        self.loadedMapTest.Set(mod,info)
    }
    self.loadedMap.Set(mod,info)
}
// 269: decl @lune.@base.@front.Front.error
func (self *Front_Front) Error(_env *LnsEnv, message string) {
    Util_errorLog(_env, message)
    Util_printStackTrace(_env)
    _env.GetVM().OS_exit(1)
}
// 275: decl @lune.@base.@front.Front.loadLua
func (self *Front_Front) loadLua(_env *LnsEnv, path string) LnsAny {
    var ret LnsAny
    Lns_LockEnvSync( _env, 277, func () {
        var chunk LnsAny
        var err LnsAny
        chunk,err = _env.GetVM().Loadfile(path)
        if chunk != nil{
            chunk_30 := chunk.(*Lns_luaValue)
            ret = Lns_unwrap( Lns_car(_env.GetVM().RunLoadedfunc(chunk_30,Lns_2DDD([]LnsAny{}))))
        } else {
            Util_errorLog(_env, Lns_unwrapDefault( err, _env.GetVM().String_format("load error -- %s.", Lns_2DDD(path))).(string))
            ret = nil
        }
    })
    return ret
}
// 295: decl @lune.@base.@front.Front.scriptPath2Module
func (self *Front_Front) scriptPath2Module(_env *LnsEnv, path string,baseDir LnsAny)(string, LnsAny) {
    var mod string
    mod = Util_scriptPath2ModuleFromProjDir(_env, path, baseDir)
    return mod, baseDir
}
// 300: decl @lune.@base.@front.Front.createPaser
func (self *Front_Front) createPaser(_env *LnsEnv, scriptPath string,baseDir LnsAny) *Tokenizer_Tokenizer {
    var mod string
    mod = front_convExp1_399(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath, baseDir)))
    return front_createPaser_6_(_env, baseDir, scriptPath, mod, self.option.FP.Get_stdinFile(_env))
}
// 306: decl @lune.@base.@front.Front.createAstSub
func (self *Front_Front) createAstSub(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,tokenizerSrc LnsAny,baseDir LnsAny,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) LnsAny {
    {
        __exp := self.moduleMgr.FP.Get(_env, mod)
        if !Lns_IsNil( __exp ) {
            _exp := __exp
            switch _matchExp0 := _exp.(type) {
            case *front_UptodateInfo__Update:
                ast := _matchExp0.Val2
                return &Converter_CreateAstResult__Ast{ast}
            case *front_UptodateInfo__Uptodate:
                Util_err(_env, _env.GetVM().String_format("can't load the multiple module -- %s", Lns_2DDD(mod)))
            }
        }
    }
    {
        _creater := self.mod2astCreate.Get(mod)
        if !Lns_IsNil( _creater ) {
            creater := _creater.(*Converter_AstCreater)
            var ast *AstInfo_ASTInfo
            ast = front_convExp1_474(Lns_2DDD(creater.FP.GetAst(_env)))
            return &Converter_CreateAstResult__Ast{ast}
        }
    }
    var astCreater *Converter_AstCreater
    astCreater = NewConverter_AstCreater(_env, NewFrontInterface_FrontAccessor(_env, self.FP), importModuleInfo, tokenizerSrc, mod, baseDir, moduleId, analyzeModule, analyzeMode, pos, self.builtinFunc, self.option)
    self.mod2astCreate.Set(mod,astCreater)
    return &Converter_CreateAstResult__Creater{astCreater}
}
// 343: decl @lune.@base.@front.Front.applyAstResult
func (self *Front_Front) applyAstResult(_env *LnsEnv, result LnsAny) *AstInfo_ASTInfo {
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
        Log_log(_env, Log_Level__Log, __func__, 350, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.GetVM().String_format("applyAstResult -- %s", Lns_2DDD(ast.FP.Get_exportInfo(_env).FP.Get_fullName(_env)))
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
// 364: decl @lune.@base.@front.Front.createAst
func (self *Front_Front) createAst(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,tokenizerSrc LnsAny,baseDir LnsAny,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *AstInfo_ASTInfo {
    var ast *AstInfo_ASTInfo
    ast = self.FP.applyAstResult(_env, self.FP.createAstSub(_env, importModuleInfo, tokenizerSrc, baseDir, mod, moduleId, analyzeModule, analyzeMode, pos))
    if self.option.DumpDebugAst{
        Util_println(_env, Lns_2DDD(_env.GetVM().String_format("=== ast:%s ==========", Lns_2DDD(mod))))
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, DumpNode_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt(_env, "", 0)))
    }
    return ast
}
// 424: decl @lune.@base.@front.Front.convertFromAst
func (self *Front_Front) convertFromAst(_env *LnsEnv, ast *AstInfo_ASTInfo,streamName string,mode LnsInt)(string, string) {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var metaStream *Util_memStream
    metaStream = NewUtil_memStream(_env)
    front_ast2Lua_7_(_env, ast, streamName, stream.FP, metaStream.FP, mode, false, self.option)
    return metaStream.FP.Get_txt(_env), stream.FP.Get_txt(_env)
}
// 438: decl @lune.@base.@front.Front.loadFromLnsTxt
func (self *Front_Front) LoadFromLnsTxt(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,baseDir LnsAny,name string,txt string) LnsAny {
    var transUnit *TransUnit_TransUnitCtrl
    transUnit = NewTransUnit_TransUnitCtrl(_env, NewFrontInterface_FrontAccessor(_env, self.FP), FrontInterface_ModuleId__tempId, importModuleInfo, &NewConvLua_MacroEvalImp(_env, self.builtinFunc).Nodes_MacroEval, false, nil, nil, nil, self.option.TargetLuaVer, self.option.TransCtrlInfo, self.builtinFunc)
    var ast *AstInfo_ASTInfo
    ast = transUnit.FP.CreateAST(_env, &Types_TokenizerSrc__LnsCode{txt, name, nil}, false, baseDir, self.option.FP.Get_stdinFile(_env), false, _env.GetVM().String_format("$load%d", Lns_2DDD(self.loadCount)), nil)
    self.loadCount = self.loadCount + 1
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(_env, ast, name, ConvLua_ConvMode__ConvMeta)
    return Lns_unwrap( Front_loadFromLuaTxt(_env, luaTxt))
}
// 503: decl @lune.@base.@front.Front.searchModuleFile
func (self *Front_Front) searchModuleFile(_env *LnsEnv, mod string,suffix string,baseDir LnsAny,outputDir LnsAny) LnsAny {
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
        lnsSearchPath = _env.GetVM().String_format("%s/?%s;%s", Lns_2DDD(outputDir_110, suffix, Lns_package_path))
    }
    if baseDir != nil{
        baseDir_112 := baseDir.(string)
        lnsSearchPath = _env.GetVM().String_format("%s/?%s;%s", Lns_2DDD(baseDir_112, suffix, Lns_package_path))
    }
    {
        _projDir := self.option.FP.Get_projDir(_env)
        if !Lns_IsNil( _projDir ) {
            projDir := _projDir.(string)
            lnsSearchPath = _env.GetVM().String_format("%s;%s", Lns_2DDD(Util_pathJoin(_env, projDir, _env.GetVM().String_format("?%s", Lns_2DDD(suffix))), Lns_package_path))
        }
    }
    lnsSearchPath = front_convExp1_1393(Lns_2DDD(_env.GetVM().String_gsub(lnsSearchPath,"%.lua$", suffix)))
    lnsSearchPath = front_convExp1_1412(Lns_2DDD(_env.GetVM().String_gsub(lnsSearchPath,"%.lua;", suffix + ";")))
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
                    Log_log(_env, Log_Level__Err, __func__, 539, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.GetVM().String_format("not found '%s' at %s", Lns_2DDD(mod, latestProjSearchPath))
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
// 600: decl @lune.@base.@front.Front.getModuleIdAndCheckUptodate
func (self *Front_Front) getModuleIdAndCheckUptodate(_env *LnsEnv, lnsPath string,mod string)(*FrontInterface_ModuleId, LnsAny) {
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
            depMod := _depMod
            dependItem := _dependItem
            var modMetaPath string
            
            {
                _modMetaPath := self.FP.searchModuleFile(_env, depMod, ".meta", nil, self.option.OutputDir)
                if _modMetaPath == nil{
                    Log_log(_env, Log_Level__Debug, __func__, 636, Log_CreateMessage(front_Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_0_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    modMetaPath = _modMetaPath.(string)
                }
            }
            var time LnsReal
            
            {
                _time := Depend_getFileLastModifiedTime(_env, modMetaPath)
                if _time == nil{
                    Log_log(_env, Log_Level__Debug, __func__, 641, Log_CreateMessage(front_Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1_))
                    
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
                        Log_log(_env, Log_Level__Debug, __func__, 649, Log_CreateMessage(front_Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_2_))
                        
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
                    Log_log(_env, Log_Level__Debug, __func__, 659, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.GetVM().String_format("NeedUpdate: %s, %d, %d", Lns_2DDD(modMetaPath, metaModuleId.FP.Get_buildCount(_env), orgMetaModuleId.FP.Get_buildCount(_env)))
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
        Log_log(_env, Log_Level__Debug, __func__, 697, Log_CreateMessage(front_Front_getModuleIdAndCheckUptodate___anonymous_1_))
        
    }
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_13_(_env, lnsPath, mod, self.option.OutputDir, metaInfo)
    if moduleId == FrontInterface_ModuleId__tempId{
        Util_err(_env, _env.GetVM().String_format("not found -- %s", Lns_2DDD(lnsPath)))
    }
    return moduleId, uptodate
}
// 710: decl @lune.@base.@front.Front.convertLns2LuaCode
func (self *Front_Front) ConvertLns2LuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,analyzeMode LnsInt,tokenizerSrc LnsAny,baseDir LnsAny,stream Lns_iStream,streamName string) string {
    var mod string
    mod = front_convExp2_650(Lns_2DDD(self.FP.scriptPath2Module(_env, streamName, baseDir)))
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, tokenizerSrc, baseDir, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, analyzeMode, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(_env, ast, streamName, ConvLua_ConvMode__ConvMeta)
    return luaTxt
}
// 724: decl @lune.@base.@front.Front.getGoAppName
func (self *Front_Front) getGoAppName(_env *LnsEnv) string {
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
// 731: decl @lune.@base.@front.Front.createGoOption
func (self *Front_Front) createGoOption(_env *LnsEnv, scriptPath string) *ConvGo_Option {
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
// 760: decl @lune.@base.@front.Front.createPythonOption
func (self *Front_Front) createPythonOption(_env *LnsEnv, scriptPath string) *ConvPython_Option {
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
// 789: decl @lune.@base.@front.Front.convertToLanguage
func (self *Front_Front) convertToLanguage(_env *LnsEnv, ast *AstInfo_ASTInfo,stream Lns_oStream,path string) {
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
// 811: decl @lune.@base.@front.Front.loadTokenizerToLuaCode
func (self *Front_Front) loadTokenizerToLuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,tokenizerSrc LnsAny,path string,mod string,baseDir LnsAny)(*FrontInterface_ModuleMeta, string) {
    __func__ := "@lune.@base.@front.Front.loadTokenizerToLuaCode"
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_13_(_env, path, mod, nil, nil)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(_env, ast, path, ConvLua_ConvMode__ConvMeta)
    Log_log(_env, Log_Level__Trace, __func__, 821, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("Meta = %s", Lns_2DDD(metaTxt))
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
                    newpath = _env.GetVM().String_format("%s/%s%s", Lns_2DDD(dir, Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string), suffix))
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
                    _fileObj := Util_openRd(_env, newpath)
                    if !Lns_IsNil( _fileObj ) {
                        fileObj := _fileObj.(Lns_iStream)
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
                _fileObj := front_convExp2_1393(Lns_2DDD(Lns_io_open(newpath, "w")))
                if !Lns_IsNil( _fileObj ) {
                    fileObj := _fileObj.(Lns_luaStream)
                    fileObj.Write(_env, saveTxt)
                    fileObj.Close(_env)
                }
            }
        }
        front_saveFile(_env, ".lua", luaTxt, self.option.ByteCompile, self.option.StripDebugInfo, false)
        if self.option.Mode == Option_ModeKind__SaveMeta{
            front_saveFile(_env, ".meta", metaTxt, self.option.ByteCompile, true, true)
        }
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
                        suffix_146 := suffix.(string)
                        front_saveFile(_env, suffix_146, memStream.FP.Get_txt(_env), false, false, false)
                    }
                }
            }
        }
    }
    return Lns_unwrap( self.moduleMgr.FP.GetMeta(_env, mod)).(*FrontInterface_ModuleMeta), luaTxt
}
// 899: decl @lune.@base.@front.Front.loadFile
func (self *Front_Front) loadFile(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,baseDir LnsAny,path string,mod string)(*FrontInterface_ModuleMeta, LnsAny) {
    __func__ := "@lune.@base.@front.Front.loadFile"
    Log_log(_env, Log_Level__Info, __func__, 903, Log_CreateMessage(func(_env *LnsEnv) string {
        __func__ := "@lune.@base.@front.Front.loadFile.<anonymous>"
        return _env.GetVM().String_format("start %s:%s", Lns_2DDD(__func__, mod))
    }))
    
    var meta *FrontInterface_ModuleMeta
    var luaTxt string
    meta,luaTxt = self.FP.loadTokenizerToLuaCode(_env, importModuleInfo, &Types_TokenizerSrc__LnsPath{baseDir, path, mod, nil}, path, mod, baseDir)
    {
        _preLoadInfo := self.preloadedModMap.Get(mod)
        if !Lns_IsNil( _preLoadInfo ) {
            preLoadInfo := _preLoadInfo
            return meta, preLoadInfo
        }
    }
    return meta, Lns_unwrap( Front_loadFromLuaTxt(_env, luaTxt))
}
// 917: decl @lune.@base.@front.Front.searchModule
func (self *Front_Front) SearchModule(_env *LnsEnv, mod string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(_env, mod, ".lns", baseDir, addSearchPath)
}
// 921: decl @lune.@base.@front.Front.searchLuaFile
func (self *Front_Front) searchLuaFile(_env *LnsEnv, moduleFullName string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(_env, moduleFullName, ".lua", baseDir, addSearchPath)
}
// 933: decl @lune.@base.@front.Front.checkUptodateMeta
func (self *Front_Front) checkUptodateMeta(_env *LnsEnv, lnsPath string,metaPath string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.checkUptodateMeta"
    var metaObj LnsAny
    
    {
        _metaObj := self.FP.loadLua(_env, metaPath)
        if _metaObj == nil{
            Log_log(_env, Log_Level__Warn, __func__, 938, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.GetVM().String_format("load error -- %s", Lns_2DDD(metaPath))
            }))
            
            return nil
        } else {
            metaObj = _metaObj
        }
    }
    var meta *Lns_luaValue
    var valid bool
    valid = true
    Lns_LockEnvSync( _env, 944, func () {
        meta = metaObj.(*Lns_luaValue)
        if meta.GetAt( "__formatVersion" ).(string) != Ver_metaVersion{
            Log_log(_env, Log_Level__Warn, __func__, 947, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.GetVM().String_format("unmatch meta version -- %s", Lns_2DDD(metaPath))
            }))
            
            valid = false
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( valid) &&
            _env.SetStackVal( meta.GetAt( "__hasTest" ).(bool)) ).(bool)){
            if meta.GetAt( "__enableTest" ).(bool) != self.option.Testing{
                Log_log(_env, Log_Level__Warn, __func__, 953, Log_CreateMessage(func(_env *LnsEnv) string {
                    return _env.GetVM().String_format("unmatch test setting -- %s", Lns_2DDD(metaPath))
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
                                        Log_log(_env, Log_Level__Warn, __func__, 967, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.GetVM().String_format("not ready -- %s, %s", Lns_2DDD(moduleLnsPath, metaPath))
                                        }))
                                        
                                        valid = false
                                        break
                                    }
                                    var moduleMetaPath string
                                    moduleMetaPath = front_convExp3_481(Lns_2DDD(_env.GetVM().String_gsub(moduleLuaPath,"%.lua$", ".meta")))
                                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                        _env.SetStackVal( Depend_existFile(_env, moduleMetaPath)) &&
                                        _env.SetStackVal( Lns_op_not(Util_getReadyCode(_env, moduleMetaPath, metaPath))) ).(bool)){
                                        Log_log(_env, Log_Level__Warn, __func__, 976, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.GetVM().String_format("not ready -- %s, %s", Lns_2DDD(moduleMetaPath, metaPath))
                                        }))
                                        
                                        valid = false
                                        break
                                    }
                                } else {
                                    Log_log(_env, Log_Level__Warn, __func__, 982, Log_CreateMessage(func(_env *LnsEnv) string {
                                        return _env.GetVM().String_format("not found .lua file for -- %s", Lns_2DDD(moduleFullName))
                                    }))
                                    
                                    valid = false
                                    break
                                }
                            }
                        } else {
                            Log_log(_env, Log_Level__Warn, __func__, 988, Log_CreateMessage(func(_env *LnsEnv) string {
                                return _env.GetVM().String_format("not found .lns file -- %s", Lns_2DDD(moduleFullName))
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
// 1010: decl @lune.@base.@front.Front.loadModuleWithBaseDir
func (self *Front_Front) loadModuleWithBaseDir(_env *LnsEnv, orgMod string,baseDir LnsAny)(LnsAny, *FrontInterface_ModuleMeta) {
    __func__ := "@lune.@base.@front.Front.loadModuleWithBaseDir"
    Log_log(_env, Log_Level__Info, __func__, 1013, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("loadModule  -- %s:%s", Lns_2DDD(orgMod, baseDir))
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
                        Util_err(_env, _env.GetVM().String_format("nothing meta -- %s", Lns_2DDD(mod)))
                    }
                }
            } else {
                {
                    _lnsPath := self.FP.SearchModule(_env, orgMod, baseDir, nil)
                    if !Lns_IsNil( _lnsPath ) {
                        lnsPath := _lnsPath.(string)
                        var luaPath LnsAny
                        luaPath = front_convExp3_883(Lns_2DDD(_env.GetVM().String_gsub(lnsPath, "%.lns$", ".lua")))
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
                                metaPath = front_convExp3_943(Lns_2DDD(_env.GetVM().String_gsub(luaPath_67, "%.lua$", ".meta")))
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
                            Log_log(_env, Log_Level__Warn, __func__, 1070, Log_CreateMessage(func(_env *LnsEnv) string {
                                return _env.GetVM().String_format("load from the binding -- %s", Lns_2DDD(mod))
                            }))
                            
                            var workMod LnsAny
                            Lns_LockEnvSync( _env, 1074, func () {
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
    Util_err(_env, _env.GetVM().String_format("load error, %s", Lns_2DDD(mod)))
// insert a dummy
    return nil,nil
}
// 1101: decl @lune.@base.@front.Front.loadModule
func (self *Front_Front) LoadModule(_env *LnsEnv, mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    return self.FP.loadModuleWithBaseDir(_env, mod, nil)
}
// 1108: decl @lune.@base.@front.Front.getLuaModulePath
func (self *Front_Front) GetLuaModulePath(_env *LnsEnv, mod string,baseDir LnsAny)(string, LnsAny, string) {
    return self.gomodMap.FP.GetLuaModulePath(_env, mod, baseDir)
}
// 1112: decl @lune.@base.@front.Front.loadMeta
func (self *Front_Front) LoadMeta(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,mod string,orgMod string,baseDir LnsAny,loader FrontInterface_ModuleLoader) LnsAny {
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
                        var luaPath LnsAny
                        luaPath = front_convExp3_1466(Lns_2DDD(_env.GetVM().String_gsub(lnsPath, "%.lns$", ".lua")))
                        if Lns_op_not(self.targetSet.Has(orgMod)){
                            luaPath = front_convExp3_1495(Lns_2DDD(_env.GetVM().String_gsub(lnsPath, "%.lns$", ".lua")))
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
                                        metaPath = front_convExp3_1603(Lns_2DDD(_env.GetVM().String_gsub(luaPath_119, "%.lua$", ".meta")))
                                        if Util_getReadyCode(_env, lnsPath, metaPath){
                                            meta = self.FP.checkUptodateMeta(_env, lnsPath, metaPath, baseDir, self.option.OutputDir)
                                        } else { 
                                            Log_log(_env, Log_Level__Warn, __func__, 1185, Log_CreateMessage(func(_env *LnsEnv) string {
                                                return _env.GetVM().String_format("%s not ready meta %s, %s", Lns_2DDD(orgMod, lnsPath, metaPath))
                                            }))
                                            
                                        }
                                    } else { 
                                        Log_log(_env, Log_Level__Warn, __func__, 1189, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.GetVM().String_format("%s not ready lua %s, %s", Lns_2DDD(orgMod, lnsPath, luaPath_119))
                                        }))
                                        
                                    }
                                } else { 
                                    Log_log(_env, Log_Level__Warn, __func__, 1193, Log_CreateMessage(func(_env *LnsEnv) string {
                                        return _env.GetVM().String_format("force analyze -- %s", Lns_2DDD(orgMod))
                                    }))
                                    
                                }
                            } else {
                                Log_log(_env, Log_Level__Warn, __func__, 1197, Log_CreateMessage(func(_env *LnsEnv) string {
                                    return _env.GetVM().String_format("%s not found lua in %s", Lns_2DDD(orgMod, self.option.OutputDir))
                                }))
                                
                            }
                        }
                        if meta != nil{
                            meta_143 := meta.(*FrontInterface_ModuleMeta)
                            self.moduleMgr.FP.AddMeta(_env, orgMod, meta_143)
                            if _switch0 := self.option.Mode; _switch0 == Option_ModeKind__Exec || _switch0 == Option_ModeKind__Shebang {
                                if luaPath != nil{
                                    luaPath_146 := luaPath.(string)
                                    {
                                        _luaCode := Util_readFile(_env, luaPath_146)
                                        if !Lns_IsNil( _luaCode ) {
                                            luaCode := _luaCode.(string)
                                            self.FP.regConvertedMap(_env, orgMod, luaCode, meta_143)
                                        }
                                    }
                                }
                            }
                        } else {
                            var metawork *FrontInterface_ModuleMeta
                            var luaTxt string
                            metawork,luaTxt = self.FP.loadTokenizerToLuaCode(_env, importModuleInfo, &Types_TokenizerSrc__LnsPath{nil, lnsPath, orgMod, nil}, lnsPath, orgMod, baseDir)
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
                                meta,luaTxt = self.FP.loadTokenizerToLuaCode(_env, importModuleInfo, &Types_TokenizerSrc__LnsCode{lnsCode, orgMod, nil}, path, orgMod, baseDir)
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
// 1246: decl @lune.@base.@front.Front.dumpTokenize
func (self *Front_Front) DumpTokenize(_env *LnsEnv, scriptPath string,baseDir LnsAny) {
    var tokenizer *Tokenizer_Tokenizer
    tokenizer = self.FP.createPaser(_env, scriptPath, baseDir)
    for  {
        var token *Types_Token
        
        {
            _token := tokenizer.FP.GetToken(_env)
            if _token == nil{
                break
            } else {
                token = _token.(*Types_Token)
            }
        }
        Util_println(_env, Lns_2DDD(Types_TokenKind_getTxt( token.Kind), token.Consecutive, token.Pos.LineNo, token.Pos.Column, token.Txt))
    }
}
// 1260: decl @lune.@base.@front.Front.dumpAst
func (self *Front_Front) DumpAst(_env *LnsEnv, tokenizerSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    Depend_profile(_env, self.option.ValidProf, conv2Form4_131(func(_env *LnsEnv) {
        var ast *AstInfo_ASTInfo
        ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, DumpNode_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt(_env, "", 0)))
    }), mod + ".profi")
}
// 1281: decl @lune.@base.@front.Front.format
func (self *Front_Front) Format(_env *LnsEnv, tokenizerSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, Formatter_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), Lns_io_stdout, false), Formatter_Opt2Stem(NewFormatter_Opt(_env, ast.FP.Get_node(_env))))
}
// 1294: decl @lune.@base.@front.Front.checkDiag
func (self *Front_Front) CheckDiag(_env *LnsEnv, tokenizerSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    Util_setErrorCode(_env, 0)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Diag, nil)
}
// 1302: decl @lune.@base.@front.Front.complete
func (self *Front_Front) complete(_env *LnsEnv, tokenizerSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), tokenizerSrc, baseDir, mod, moduleId, self.option.AnalyzeModule, TransUnit_AnalyzeMode__Complete, self.option.AnalyzePos)
}
// 1310: decl @lune.@base.@front.Front.completeFromCode
func (self *Front_Front) CompleteFromCode(_env *LnsEnv, lnsCode string,mod string,baseDir LnsAny) {
    self.FP.complete(_env, &Types_TokenizerSrc__LnsCode{lnsCode, mod, nil}, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), baseDir)
}
// 1316: decl @lune.@base.@front.Front.inquire
func (self *Front_Front) Inquire(_env *LnsEnv, tokenizerSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), tokenizerSrc, baseDir, mod, moduleId, self.option.AnalyzeModule, TransUnit_AnalyzeMode__Inquire, self.option.AnalyzePos)
}
// 1325: decl @lune.@base.@front.Front.createGlue
func (self *Front_Front) CreateGlue(_env *LnsEnv, tokenizerSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,baseDir LnsAny) {
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), tokenizerSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var filter *Nodes_Filter
    filter = GlueFilter_createFilter(_env, self.option.OutputDir)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, filter, 0)
}
// 1348: decl @lune.@base.@front.Front.convertToLua
func (self *Front_Front) convertToLua(_env *LnsEnv, scriptPath string,baseDir LnsAny,convMode LnsInt,streamLua Lns_oStream,streamMeta Lns_oStream) LnsAny {
    var mod string
    mod = front_convExp4_432(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath, baseDir)))
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_13_(_env, scriptPath, mod, nil, nil)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_TokenizerSrc__LnsPath{baseDir, scriptPath, mod, nil}, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(_env, ast, scriptPath, convMode)
    streamLua.Write(_env, luaTxt)
    streamMeta.Write(_env, metaTxt)
    self.FP.convertToLanguage(_env, ast, streamLua, scriptPath)
    return ast
}
// 1371: decl @lune.@base.@front.Front.saveToGo
func (self *Front_Front) saveToGo(_env *LnsEnv, scriptPath string,astResult LnsAny,mod string) LnsAny {
    return NewConverter_GoConverter(_env, scriptPath, astResult, mod, self.option, self.FP.createGoOption(_env, scriptPath))
}
// 1382: decl @lune.@base.@front.Front.saveToPython
func (self *Front_Front) saveToPython(_env *LnsEnv, scriptPath string,astResult LnsAny,mod string) LnsAny {
    return NewConverter_PythonConverter(_env, scriptPath, astResult, mod, self.option, self.FP.createPythonOption(_env, scriptPath))
}
// 1394: decl @lune.@base.@front.Front.saveToC
func (self *Front_Front) SaveToC(_env *LnsEnv, scriptPath string,ast *AstInfo_ASTInfo) {
    var cPath string
    cPath = front_convExp4_594(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".c")))
    var file Lns_luaStream
    
    {
        _file := front_convExp4_609(Lns_2DDD(Lns_io_open(cPath, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var hPath string
    hPath = front_convExp4_623(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".h")))
    var hFile Lns_luaStream
    
    {
        _hFile := front_convExp4_639(Lns_2DDD(Lns_io_open(hPath, "w")))
        if _hFile == nil{
            return 
        } else {
            hFile = _hFile.(Lns_luaStream)
        }
    }
    file.Close(_env)
    hFile.Close(_env)
}
// 1410: decl @lune.@base.@front.Front.outputBuiltin
func (self *Front_Front) OutputBuiltin(_env *LnsEnv, scriptPath string) {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, "lns_builtin", nil)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_TokenizerSrc__LnsCode{"", mod, nil}, baseDir, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    self.FP.SaveToC(_env, scriptPath, ast)
}
// 1432: decl @lune.@base.@front.Front.saveToLua
func (self *Front_Front) saveToLua(_env *LnsEnv, updateInfo *front_UpdateInfo,baseDir LnsAny) LnsAny {
    var scriptPath string
    scriptPath = updateInfo.FP.Get_scriptPath(_env)
    var dependsPath LnsAny
    dependsPath = updateInfo.FP.Get_dependsPath(_env)
    var moduleId *FrontInterface_ModuleId
    moduleId = updateInfo.FP.Get_moduleId(_env)
    var uptodate LnsAny
    uptodate = updateInfo.FP.Get_uptodate(_env)
    var mod string
    mod = front_convExp4_753(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath, baseDir)))
    var luaPath string
    luaPath = front_convExp4_767(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".lua")))
    var metaPath string
    metaPath = front_convExp4_780(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".meta")))
    if Lns_isCondTrue( self.option.OutputDir){
        var filename string
        filename = front_convExp4_797(Lns_2DDD(_env.GetVM().String_gsub(mod,"%.", "/")))
        luaPath = _env.GetVM().String_format("%s/%s.lua", Lns_2DDD(self.option.OutputDir, filename))
        metaPath = _env.GetVM().String_format("%s/%s.meta", Lns_2DDD(self.option.OutputDir, filename))
    }
    var convMode LnsInt
    convMode = ConvLua_ConvMode__ConvMeta
    if luaPath == scriptPath{
        Util_errorLog(_env, _env.GetVM().String_format("%s is illegal filename.", Lns_2DDD(luaPath)))
        return nil
    }
    switch _matchExp0 := uptodate.(type) {
    case *front_ModuleUptodate__NeedUpdate:
        var result LnsAny
        result = self.FP.createAstSub(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_TokenizerSrc__LnsPath{baseDir, scriptPath, mod, nil}, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
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
            Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*Converter_LuaConverter).FP.SaveLua(_env)})/* 1503:13 */)
            _env.NilAccFin(_env.NilAccPush(goConv) && 
            Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*Converter_GoConverter).FP.SaveGo(_env)})/* 1504:13 */)
            _env.NilAccFin(_env.NilAccPush(pyConv) && 
            Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*Converter_PythonConverter).FP.SavePython(_env)})/* 1505:13 */)
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
// 1535: decl @lune.@base.@front.Front.outputBootC
func (self *Front_Front) outputBootC(_env *LnsEnv, scriptPath string) {
}
// 1553: decl @lune.@base.@front.Front.convertLnsCode2LuaCodeWithOpt
func (self *Front_Front) ConvertLnsCode2LuaCodeWithOpt(_env *LnsEnv, lnsCode string,path string,baseDir LnsAny) string {
    return self.FP.ConvertLns2LuaCode(_env, NewFrontInterface_ImportModuleInfo(_env), TransUnit_AnalyzeMode__Compile, &Types_TokenizerSrc__LnsCode{lnsCode, path, nil}, baseDir, NewUtil_TxtStream(_env, lnsCode).FP, path)
}
// 1619: decl @lune.@base.@front.Front.build
func (self *Front_Front) Build(_env *LnsEnv, buildMode LnsAny,astCallback LnsAny) {
    var front_createUpdateInfo func(_env *LnsEnv, scriptPath string,dependsPath LnsAny,baseDir LnsAny) *front_UpdateInfo
    front_createUpdateInfo = func(_env *LnsEnv, scriptPath string,dependsPath LnsAny,baseDir LnsAny) *front_UpdateInfo {
        var mod string
        mod = front_convExp0_1679(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath, baseDir)))
        var moduleId *FrontInterface_ModuleId
        var uptodate LnsAny
        moduleId,uptodate = self.FP.getModuleIdAndCheckUptodate(_env, scriptPath, mod)
        return Newfront_UpdateInfo(_env, scriptPath, dependsPath, moduleId, uptodate)
    }
    var front_process func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo,baseDir LnsAny) LnsAny
    front_process = func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo,baseDir LnsAny) LnsAny {
        var mod string
        mod = front_convExp0_1737(Lns_2DDD(self.FP.scriptPath2Module(_env, updateInfo.FP.Get_scriptPath(_env), baseDir)))
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
                result = self.FP.createAstSub(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_TokenizerSrc__LnsPath{baseDir, updateInfo.FP.Get_scriptPath(_env), mod, nil}, baseDir, mod, updateInfo.FP.Get_moduleId(_env), nil, TransUnit_AnalyzeMode__Compile, nil)
                return Converter_ConverterFunc(func(_env *LnsEnv) {
                    self.FP.applyAstResult(_env, result)
                })
            }
        }
        return nil
    }
    Depend_profile(_env, self.option.ValidProf, conv2Form0_2211(func(_env *LnsEnv) {
        __func__ := "@lune.@base.@front.Front.build.<anonymous>"
        if self.option.ScriptPath == "@-"{
            var baseDir LnsAny
            baseDir = self.option.FP.Get_projDir(_env)
            for _, _path := range( self.option.BatchList.Items ) {
                path := _path
                self.targetSet.Add(Lns_car(self.FP.scriptPath2Module(_env, path, baseDir)).(string))
            }
            var postProcessMap *LnsMap2_[LnsInt,Converter_ConverterFunc]
            postProcessMap = NewLnsMap2_[LnsInt,Converter_ConverterFunc]( map[LnsInt]Converter_ConverterFunc{})
            for _index, _path := range( self.option.BatchList.Items ) {
                index := _index + 1
                path := _path
                var updateInfo *front_UpdateInfo
                updateInfo = front_createUpdateInfo(_env, path, Lns_car(_env.GetVM().String_gsub(path,".lns$", ".d")).(string), baseDir)
                Util_println(_env, Lns_2DDD(_env.GetVM().String_format("%s: start...", Lns_2DDD(updateInfo.FP.Get_scriptPath(_env)))))
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
                    postProcess := __forsortCollection0.Items[ _index ]
                    index := _index
                    var prev LnsReal
                    prev = _env.GetVM().OS_clock()
                    var path string
                    path = self.option.BatchList.GetAt(index)
                    Util_println(_env, Lns_2DDD(_env.GetVM().String_format("%s: waiting...", Lns_2DDD(path))))
                    postProcess(_env)
                    Util_println(_env, Lns_2DDD(_env.GetVM().String_format("%s: done %g msec", Lns_2DDD(path, (_env.GetVM().OS_clock() - prev) * LnsReal(1000)))))
                }
            }
        } else { 
            var baseDir LnsAny
            _,baseDir = front_getBaseDir_18_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env))
            {
                _postProcess := front_process(_env, true, front_createUpdateInfo(_env, self.option.ScriptPath, nil, baseDir), baseDir)
                if !Lns_IsNil( _postProcess ) {
                    postProcess := _postProcess.(Converter_ConverterFunc)
                    postProcess(_env)
                }
            }
        }
        if astCallback != nil{
            astCallback_595 := astCallback.(Front_AstCallback)
            for _, _mod := range( self.moduleMgr.FP.GetModList(_env).Items ) {
                mod := _mod.(string)
                {
                    __exp := self.moduleMgr.FP.GetAst(_env, mod)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*AstInfo_ASTInfo)
                        astCallback_595(_env, _exp)
                    } else {
                        Log_log(_env, Log_Level__Err, __func__, 1699, Log_CreateMessage(func(_env *LnsEnv) string {
                            return _env.GetVM().String_format("not found AST -- %s", Lns_2DDD(mod))
                        }))
                        
                    }
                }
            }
        }
    }), self.option.ScriptPath + ".profi")
}
// 1721: decl @lune.@base.@front.Front.setupPreloadWithImportedModules
func (self *Front_Front) SetupPreloadWithImportedModules(_env *LnsEnv, asyncFlag bool) {
    if asyncFlag{
            {
                var subModPreLoad string
                subModPreLoad = "return function( submod2Code, dumpCode )\n   local preloadFunc = function( mod )\n      code = submod2Code[ mod ]\n      local loadFunc = loadstring or load -- lua5.1 and lua5.2\n      local loaded, mess = loadFunc( code )\n      if not loaded then\n         error( mess )\n      end\n      return loaded()\n   end\n   for mod, code in pairs( submod2Code ) do\n      if dumpCode then\n         print( string.format( \"mod: %s %s\", mod, code ) )\n      end\n      package.preload[ mod ] = preloadFunc\n   end\nend\n"
                var loaded *Lns_luaValue
                
                {
                    _loaded := front_convExp0_2361(Lns_2DDD(_env.GetVM().Load(subModPreLoad, nil)))
                    if _loaded == nil{
                        Util_err(_env, "failed to subModPreLoad")
                    } else {
                        loaded = _loaded.(*Lns_luaValue)
                    }
                }
                var preloadFunc LnsAny
                
                {
                    _preloadFunc := front_convExp0_2379(Lns_2DDD(_env.GetVM().RunLoadedfunc(loaded,Lns_2DDD([]LnsAny{}))))
                    if _preloadFunc == nil{
                        Util_err(_env, "failed to preloadFunc")
                    } else {
                        preloadFunc = _preloadFunc
                    }
                }
                _env.GetVM().RunLoadedfunc((preloadFunc.(*Lns_luaValue)),Lns_2DDD(Lns_2DDD(self.convertedMap, Log_getLevel(_env) >= Log_Level__Debug)))
            }
    } else { 
        Lns_LockEnvSync( _env, 1729, func () {
            {
                var subModPreLoad string
                subModPreLoad = "return function( submod2Code, dumpCode )\n   local preloadFunc = function( mod )\n      code = submod2Code[ mod ]\n      local loadFunc = loadstring or load -- lua5.1 and lua5.2\n      local loaded, mess = loadFunc( code )\n      if not loaded then\n         error( mess )\n      end\n      return loaded()\n   end\n   for mod, code in pairs( submod2Code ) do\n      if dumpCode then\n         print( string.format( \"mod: %s %s\", mod, code ) )\n      end\n      package.preload[ mod ] = preloadFunc\n   end\nend\n"
                var loaded *Lns_luaValue
                
                {
                    _loaded := front_convExp0_2436(Lns_2DDD(_env.GetVM().Load(subModPreLoad, nil)))
                    if _loaded == nil{
                        Util_err(_env, "failed to subModPreLoad")
                    } else {
                        loaded = _loaded.(*Lns_luaValue)
                    }
                }
                var preloadFunc LnsAny
                
                {
                    _preloadFunc := front_convExp0_2454(Lns_2DDD(_env.GetVM().RunLoadedfunc(loaded,Lns_2DDD([]LnsAny{}))))
                    if _preloadFunc == nil{
                        Util_err(_env, "failed to preloadFunc")
                    } else {
                        preloadFunc = _preloadFunc
                    }
                }
                _env.GetVM().RunLoadedfunc((preloadFunc.(*Lns_luaValue)),Lns_2DDD(Lns_2DDD(self.convertedMap, Log_getLevel(_env) >= Log_Level__Debug)))
            }
        })
    }
    
}
// 1767: decl @lune.@base.@front.Front.executeLns
func (self *Front_Front) executeLns(_env *LnsEnv, path string,baseDir LnsAny)(LnsAny, string) {
    __func__ := "@lune.@base.@front.Front.executeLns"
    var mod string
    mod = Util_scriptPath2ModuleFromProjDir(_env, path, baseDir)
    var tokenizerSrc LnsAny
    tokenizerSrc = &Types_TokenizerSrc__LnsPath{baseDir, path, mod, nil}
    var luaCode string
    _,luaCode = self.FP.loadTokenizerToLuaCode(_env, NewFrontInterface_ImportModuleInfo(_env), tokenizerSrc, path, mod, baseDir)
    Log_log(_env, Log_Level__Debug, __func__, 1778, Log_CreateMessage(func(_env *LnsEnv) string {
        return "luacode: " + luaCode
    }))
    
    self.FP.SetupPreloadWithImportedModules(_env, false)
    return Front_loadFromLuaTxt(_env, luaCode), mod
}
// 1787: decl @lune.@base.@front.Front.executeLnsAndTest
func (self *Front_Front) executeLnsAndTest(_env *LnsEnv, path string,baseDir LnsAny) {
    var mod string
    _,mod = self.FP.executeLns(_env, path, baseDir)
    if self.option.Testing{
        var code string
        code = "local Testing = require( \"lune.base.Testing\" )\nreturn function( path )\n  Testing.run( path );\n  Testing.outputAllResult( io.stdout );\nend\n"
        Lns_LockEnvSync( _env, 1798, func () {
            var loaded LnsAny
            var mess LnsAny
            loaded,mess = _env.GetVM().Load(code, nil)
            if loaded != nil{
                loaded_657 := loaded.(*Lns_luaValue)
                {
                    _testFunc := front_convExp0_2666(Lns_2DDD(_env.GetVM().RunLoadedfunc(loaded_657,Lns_2DDD([]LnsAny{}))))
                    if !Lns_IsNil( _testFunc ) {
                        testFunc := _testFunc
                        _env.GetVM().RunLoadedfunc((testFunc.(*Lns_luaValue)),Lns_2DDD(Lns_2DDD(mod)))
                    }
                }
            } else {
                Util_println(_env, Lns_2DDD(mess))
            }
        })
    }
}
// 1811: decl @lune.@base.@front.Front.exec
func (self *Front_Front) Exec(_env *LnsEnv) {
    __func__ := "@lune.@base.@front.Front.exec"
    Log_log(_env, Log_Level__Trace, __func__, 1813, Log_CreateMessage(func(_env *LnsEnv) string {
        return Option_ModeKind_getTxt( self.option.Mode)
    }))
    
    if _switch0 := self.option.Mode; _switch0 == Option_ModeKind__Token {
        self.FP.DumpTokenize(front_convExp0_2760(_env, Lns_2DDD(front_getBaseDir_18_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Ast {
        self.FP.DumpAst(front_convExp0_2786(_env, Lns_2DDD(front_getParseSrcAndBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Format {
        self.FP.Format(front_convExp0_2812(_env, Lns_2DDD(front_getParseSrcAndBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Diag {
        self.FP.CheckDiag(front_convExp0_2838(_env, Lns_2DDD(front_getParseSrcAndBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Complete {
        self.FP.complete(front_convExp0_2864(_env, Lns_2DDD(front_getParseSrcAndBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Inquire {
        self.FP.Inquire(front_convExp0_2890(_env, Lns_2DDD(front_getParseSrcAndBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Glue {
        self.FP.CreateGlue(front_convExp0_2916(_env, Lns_2DDD(front_getParseSrcAndBaseDir_19_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__Lua || _switch0 == Option_ModeKind__LuaMeta {
        var convMode LnsInt
        if self.option.Mode == Option_ModeKind__Lua{
            convMode = ConvLua_ConvMode__Convert
        } else { 
            convMode = ConvLua_ConvMode__ConvMeta
        }
        var scriptPath string
        var baseDir LnsAny
        scriptPath,baseDir = front_getBaseDir_18_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env))
        self.FP.convertToLua(_env, scriptPath, baseDir, convMode, Lns_io_stdout, Lns_io_stdout)
    } else if _switch0 == Option_ModeKind__Save || _switch0 == Option_ModeKind__SaveMeta {
        self.FP.Build(_env, Front_BuildMode__Save_Obj, nil)
    } else if _switch0 == Option_ModeKind__BuildAst {
        self.FP.Build(_env, Front_BuildMode__CreateAst_Obj, Front_AstCallback(front_Front_exec___anonymous_1_))
    } else if _switch0 == Option_ModeKind__Shebang {
        Depend_setupShebang(_env)
        Lns_LockEnvSync( _env, 1868, func () {
            var modObj LnsAny
            
            {
                _modObj := front_convExp0_3077(Lns_2DDD(self.FP.executeLns(front_convExp0_3064(_env, Lns_2DDD(front_getBaseDir_18_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))))
                if _modObj == nil{
                    _env.GetVM().OS_exit(1)
                } else {
                    modObj = _modObj
                }
            }
            var code LnsInt
            code = Depend_runMain(_env, modObj.(*Lns_luaValue).GetAt("__main"), self.option.ShebangArgList)
            _env.GetVM().OS_exit(code)
        })
    } else if _switch0 == Option_ModeKind__GoMod {
        for _, _path := range( self.gomodMap.FP.GetModPathList(_env).Items ) {
            path := _path.(string)
            Util_println(_env, Lns_2DDD(path))
        }
    } else if _switch0 == Option_ModeKind__Exec {
        self.FP.executeLnsAndTest(front_convExp0_3146(_env, Lns_2DDD(front_getBaseDir_18_(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env)))))
    } else if _switch0 == Option_ModeKind__BootC {
        self.FP.outputBootC(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Builtin {
        self.FP.OutputBuiltin(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__MkMain {
        var mod string
        mod = front_convExp0_3196(Lns_2DDD(self.FP.scriptPath2Module(_env, self.option.ScriptPath, self.option.FP.Get_projDir(_env))))
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
        Util_println(_env, Lns_2DDD("illegal mode"))
    }
}
// 469: decl @lune.@base.@front.MetaForBuildId.createModuleId
func (self *front_MetaForBuildId) CreateModuleId(_env *LnsEnv) *FrontInterface_ModuleId {
    return FrontInterface_ModuleId_createIdFromTxt(_env, self.G__buildId)
}
// 474: decl @lune.@base.@front.MetaForBuildId.LoadFromMeta
func front_MetaForBuildId_LoadFromMeta_5_(_env *LnsEnv, metaPath string)(LnsAny, LnsAny) {
    {
        _fileObj := Util_openRd(_env, metaPath)
        if !Lns_IsNil( _fileObj ) {
            fileObj := _fileObj.(Lns_iStream)
            var luaCode LnsAny
            luaCode = fileObj.Read(_env, "*a")
            fileObj.Close(_env)
            if luaCode != nil{
                luaCode_93 := luaCode.(string)
                var meta LnsAny
                    meta = front_convExp1_1156(Lns_2DDD(front_MetaForBuildId__fromStem_4_(_env, _env.GetVM().ExpandLuavalMap(Front_loadFromLuaTxt(_env, luaCode_93)),nil)))
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
func front_LoadInfo_toSlice(slice []LnsAny) []*front_LoadInfo {
    ret := make([]*front_LoadInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(front_LoadInfoDownCast).Tofront_LoadInfo()
    }
    return ret
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
    loadedMetaMap *LnsMap2_[string,*FrontInterface_ModuleMeta]
    FP front_ModuleMgrMtd
}
func front_ModuleMgr2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_ModuleMgr).FP
}
func front_ModuleMgr_toSlice(slice []LnsAny) []*front_ModuleMgr {
    ret := make([]*front_ModuleMgr, len(slice))
    for index, val := range slice {
        ret[index] = val.(front_ModuleMgrDownCast).Tofront_ModuleMgr()
    }
    return ret
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
// 91: DeclConstr
func (self *front_ModuleMgr) Initfront_ModuleMgr(_env *LnsEnv) {
    self.mod2info = NewUtil_OrderdMap(_env)
    self.loadedMetaMap = NewLnsMap2_[string,*FrontInterface_ModuleMeta]( map[string]*FrontInterface_ModuleMeta{})
}


// declaration Class -- Front
type Front_FrontMtd interface {
    applyAstResult(_env *LnsEnv, arg1 LnsAny) *AstInfo_ASTInfo
    Build(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny)
    CheckDiag(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    checkUptodateMeta(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny, arg4 LnsAny) LnsAny
    complete(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    CompleteFromCode(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny)
    convertFromAst(_env *LnsEnv, arg1 *AstInfo_ASTInfo, arg2 string, arg3 LnsInt)(string, string)
    ConvertLns2LuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsInt, arg3 LnsAny, arg4 LnsAny, arg5 Lns_iStream, arg6 string) string
    ConvertLnsCode2LuaCodeWithOpt(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny) string
    convertToLanguage(_env *LnsEnv, arg1 *AstInfo_ASTInfo, arg2 Lns_oStream, arg3 string)
    convertToLua(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsInt, arg4 Lns_oStream, arg5 Lns_oStream) LnsAny
    createAst(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 string, arg5 *FrontInterface_ModuleId, arg6 LnsAny, arg7 LnsInt, arg8 LnsAny) *AstInfo_ASTInfo
    createAstSub(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 string, arg5 *FrontInterface_ModuleId, arg6 LnsAny, arg7 LnsInt, arg8 LnsAny) LnsAny
    CreateGlue(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    createGoOption(_env *LnsEnv, arg1 string) *ConvGo_Option
    createPaser(_env *LnsEnv, arg1 string, arg2 LnsAny) *Tokenizer_Tokenizer
    createPythonOption(_env *LnsEnv, arg1 string) *ConvPython_Option
    DumpAst(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 *FrontInterface_ModuleId, arg4 LnsAny)
    DumpTokenize(_env *LnsEnv, arg1 string, arg2 LnsAny)
    Error(_env *LnsEnv, arg1 string)
    Exec(_env *LnsEnv)
    executeLns(_env *LnsEnv, arg1 string, arg2 LnsAny)(LnsAny, string)
    executeLnsAndTest(_env *LnsEnv, arg1 string, arg2 LnsAny)
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
    loadTokenizerToLuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny)(*FrontInterface_ModuleMeta, string)
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
    SetupPreloadWithImportedModules(_env *LnsEnv, arg1 bool)
}
type Front_Front struct {
    option *Option_Option
    loadedMap *LnsMap2_[string,*front_LoadInfo]
    loadedMapTest *LnsMap2_[string,*front_LoadInfo]
    convertedMap *LnsMap2_[string,string]
    gomodMap *GoMod_ModInfo
    bindModuleSet *LnsSet2_[string]
    moduleMgr *front_ModuleMgr
    targetSet *LnsSet2_[string]
    mod2astCreate *LnsMap2_[string,*Converter_AstCreater]
    mod2loader *LnsMap2_[string,FrontInterface_ModuleLoader]
    preloadedModMap *LnsMap2_[string,LnsAny]
    loadCount LnsInt
    builtinFunc *Builtin_BuiltinFuncType
    FP Front_FrontMtd
}
func Front_Front2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Front_Front).FP
}
func Front_Front_toSlice(slice []LnsAny) []*Front_Front {
    ret := make([]*Front_Front, len(slice))
    for index, val := range slice {
        ret[index] = val.(Front_FrontDownCast).ToFront_Front()
    }
    return ret
}
type Front_FrontDownCast interface {
    ToFront_Front() *Front_Front
}
func Front_FrontDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Front_FrontDownCast)
    if ok { return work.ToFront_Front() }
    return nil
}
func (obj *Front_Front) ToFront_Front() *Front_Front {
    return obj
}
func NewFront_Front(_env *LnsEnv, arg1 *Option_Option, arg2 LnsAny) *Front_Front {
    obj := &Front_Front{}
    obj.FP = obj
    obj.InitFront_Front(_env, arg1, arg2)
    return obj
}
// 191: DeclConstr
func (self *Front_Front) InitFront_Front(_env *LnsEnv, option *Option_Option,bindModuleList LnsAny) {
    self.mod2loader = NewLnsMap2_[string,FrontInterface_ModuleLoader]( map[string]FrontInterface_ModuleLoader{})
    self.mod2astCreate = NewLnsMap2_[string,*Converter_AstCreater]( map[string]*Converter_AstCreater{})
    self.loadCount = 0
    self.targetSet = NewLnsSet2_[string]([]string{})
    self.bindModuleSet = NewLnsSet2_[string]([]string{})
    if bindModuleList != nil{
        bindModuleList_106 := bindModuleList.(*LnsList2_[string])
        for _, _mod := range( bindModuleList_106.Items ) {
            mod := _mod
            self.bindModuleSet.Add(mod)
        }
    }
    self.moduleMgr = Newfront_ModuleMgr(_env)
    var gomodDir string
    if option.ScriptPath != ""{
        gomodDir = Util_parentPath(_env, option.ScriptPath)
    } else { 
        gomodDir = ""
    }
    self.gomodMap = GoMod_getGoMap(_env, gomodDir)
    DependLuaOnLns_addGoModPath(_env, self.gomodMap.FP.GetModPathList(_env))
    self.option = option
    self.loadedMap = NewLnsMap2_[string,*front_LoadInfo]( map[string]*front_LoadInfo{})
    self.loadedMapTest = NewLnsMap2_[string,*front_LoadInfo]( map[string]*front_LoadInfo{})
    self.convertedMap = NewLnsMap2_[string,string]( map[string]string{})
    {
        var builtin *Builtin_Builtin
        builtin = NewBuiltin_Builtin(_env, self.option.TargetLuaVer, option.TransCtrlInfo)
        self.builtinFunc = builtin.FP.RegistBuiltInScope(_env)
    }
    var loadedMap *LnsMap2_[string,LnsAny]
    loadedMap = NewLnsMap2_[string,LnsAny]( map[string]LnsAny{})
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
func front_DependMetaInfo_toSlice(slice []LnsAny) []*front_DependMetaInfo {
    ret := make([]*front_DependMetaInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(front_DependMetaInfoDownCast).Tofront_DependMetaInfo()
    }
    return ret
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
    G__dependModuleMap *LnsMap2_[string,*front_DependMetaInfo]
    G__subModuleMap *LnsList2_[string]
    G__enableTest bool
    FP front_MetaForBuildIdMtd
}
func front_MetaForBuildId2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_MetaForBuildId).FP
}
func front_MetaForBuildId_toSlice(slice []LnsAny) []*front_MetaForBuildId {
    ret := make([]*front_MetaForBuildId, len(slice))
    for index, val := range slice {
        ret[index] = val.(front_MetaForBuildIdDownCast).Tofront_MetaForBuildId()
    }
    return ret
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
func Newfront_MetaForBuildId(_env *LnsEnv, arg1 string, arg2 *LnsMap2_[string,*front_DependMetaInfo], arg3 *LnsList2_[string], arg4 bool) *front_MetaForBuildId {
    obj := &front_MetaForBuildId{}
    obj.FP = obj
    obj.Initfront_MetaForBuildId(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *front_MetaForBuildId) Initfront_MetaForBuildId(_env *LnsEnv, arg1 string, arg2 *LnsMap2_[string,*front_DependMetaInfo], arg3 *LnsList2_[string], arg4 bool) {
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
    if ok,conv,mess := Lns_ToLnsMap2Sub[string,*front_DependMetaInfo]( objMap.Items["__dependModuleMap"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil},Lns_ToObjParam{
            front_DependMetaInfo_FromMapSub, false,nil}}); !ok {
       return false,nil,"__dependModuleMap:" + mess.(string)
    } else {
       newObj.G__dependModuleMap = conv.(*LnsMap2_[string,*front_DependMetaInfo])
    }
    if ok,conv,mess := Lns_ToList2Sub[string]( objMap.Items["__subModuleMap"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil}}); !ok {
       return false,nil,"__subModuleMap:" + mess.(string)
    } else {
       newObj.G__subModuleMap = conv.(*LnsList2_[string])
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
func front_UpdateInfo_toSlice(slice []LnsAny) []*front_UpdateInfo {
    ret := make([]*front_UpdateInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(front_UpdateInfoDownCast).Tofront_UpdateInfo()
    }
    return ret
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
    Lns_Tokenizer_init(_env)
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
