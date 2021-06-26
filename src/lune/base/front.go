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
// decl alge -- CreateAstResult
type front_CreateAstResult = LnsAny
type front_CreateAstResult__Ast struct{
Val1 *AstInfo_ASTInfo
}
func (self *front_CreateAstResult__Ast) GetTxt() string {
return "CreateAstResult.Ast"
}
type front_CreateAstResult__Creater struct{
Val1 *front_AstCreater
}
func (self *front_CreateAstResult__Creater) GetTxt() string {
return "CreateAstResult.Creater"
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
type Front_AstCallback func (_env *LnsEnv, arg1 *AstInfo_ASTInfo)
type front_ConverterFunc_6_ func (_env *LnsEnv)
// for 1309: ExpCast
func conv2Form3_1409( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1963: ExpCast
func conv2Form5_536( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 2090: ExpCast
func conv2Form5_941( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 531
func front_convExp1_1043(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 616
func front_convExp1_1390(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 638
func front_convExp1_1466(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 936
func front_convExp2_1110(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 945
func front_convExp2_1139(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1492
func front_convExp4_568(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1501
func front_convExp4_639(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2065
func front_convExp5_866(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2089
func front_convExp5_961(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 547
func front_convExp1_1078(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 622
func front_convExp1_1377(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 637
func front_convExp1_1445(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 667
func front_convExp1_1614(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 668
func front_convExp1_1633(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 709
func front_convExp2_41(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 891
func front_convExp2_835(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 929
func front_convExp2_1041(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1090
func front_convExp3_32(arg1 []LnsAny) (string, LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 1690
func front_convExp0_2089(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 385
func front_convExp1_477(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 409
func front_convExp1_590(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 455
func front_convExp1_681(arg1 []LnsAny) *AstInfo_ASTInfo {
    return Lns_getFromMulti( arg1, 0 ).(*AstInfo_ASTInfo)
}
// for 472
func front_convExp1_758(arg1 []LnsAny) (*AstInfo_ASTInfo, *FrontInterface_ModuleInfo, *FrontInterface_ModuleMeta) {
    return Lns_getFromMulti( arg1, 0 ).(*AstInfo_ASTInfo), Lns_getFromMulti( arg1, 1 ).(*FrontInterface_ModuleInfo), Lns_getFromMulti( arg1, 2 ).(*FrontInterface_ModuleMeta)
}
// for 555
func front_convExp1_1104(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 595
func front_convExp1_1301(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 789
func front_convExp2_310(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 816
func front_convExp2_461(arg1 []LnsAny) (LnsAny, string, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 855
func front_convExp2_648(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 860
func front_convExp2_693(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 882
func front_convExp2_777(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 917
func front_convExp2_919(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 985
func front_convExp2_1359(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1051
func front_convExp2_1755(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1104
func front_convExp3_128(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1113
func front_convExp3_188(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1135
func front_convExp3_306(arg1 []LnsAny) (*FrontInterface_ModuleMeta, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 )
}
// for 1211
func front_convExp3_706(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1238
func front_convExp3_813(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1264
func front_convExp3_1124(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1274
func front_convExp3_1202(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1306
func front_convExp3_1313(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1326
func front_convExp3_1424(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1342
func front_convExp3_1510(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1350
func front_convExp3_1571(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1358
func front_convExp3_1628(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1369
func front_convExp3_1685(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1382
func front_convExp4_28(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1441
func front_convExp4_279(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1442
func front_convExp4_292(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1497
func front_convExp4_592(arg1 []LnsAny) (bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1620
func front_convExp4_838(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1655
func front_convExp4_975(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1662
func front_convExp4_1037(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1736
func front_convExp4_1154(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1760
func front_convExp4_1208(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1762
func front_convExp4_1223(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1766
func front_convExp4_1237(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1767
func front_convExp4_1253(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1776
func front_convExp4_1276(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1806
func front_convExp4_1363(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 1807
func front_convExp4_1377(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1808
func front_convExp4_1390(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1810
func front_convExp4_1407(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1930
func front_convExp5_103(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1931
func front_convExp5_116(arg1 []LnsAny) (*FrontInterface_ModuleId, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 1 )
}
// for 1936
func front_convExp5_154(arg1 []LnsAny) (string, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 )
}
// for 2075
func front_convExp5_905(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2087
func front_convExp5_924(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2105
func front_convExp5_1024(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
func front__anonymous_0_(_env *LnsEnv, ver LnsInt) {
    LuaVer_setCurVer(_env, ver)
}
// 82: decl @lune.@base.@front.createModuleInfo
func front_createModuleInfo_4_(_env *LnsEnv, ast *AstInfo_ASTInfo,mod string,moduleId *FrontInterface_ModuleId) *FrontInterface_ModuleInfo {
    var exportInfo *FrontInterface_ExportInfo
    exportInfo = ast.FP.Get_exportInfo(_env)
    return NewFrontInterface_ModuleInfo(_env, exportInfo)
}



// 396: decl @lune.@base.@front.createPaser
func front_createPaser_9_(_env *LnsEnv, path string,mod string,stdinFile LnsAny) *Parser_Parser {
    return &Parser_StreamParser_create(_env, &Types_ParserSrc__LnsPath{path, mod, nil}, false, stdinFile, nil).Parser_Parser
}

// 419: decl @lune.@base.@front.getAstFromResult
func front_getAstFromResult_11_(_env *LnsEnv, result LnsAny) *AstInfo_ASTInfo {
    switch _matchExp0 := result.(type) {
    case *front_CreateAstResult__Ast:
    ast := _matchExp0.Val1
        return ast
    case *front_CreateAstResult__Creater:
    creater := _matchExp0.Val1
        return Lns_car(creater.FP.getAst(_env)).(*AstInfo_ASTInfo)
    }
// insert a dummy
    return nil
}

// 500: decl @lune.@base.@front.ast2LuaMain
func front_ast2LuaMain_12_(_env *LnsEnv, ast *AstInfo_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,option *Option_Option) *ConvLua_FilterInfo {
    var exportInfo *FrontInterface_ExportInfo
    exportInfo = ast.FP.Get_exportInfo(_env)
    var conv *ConvLua_FilterInfo
    conv = ConvLua_createFilter(_env, streamName, stream, metaStream, convMode, inMacro, exportInfo.FP.Get_moduleTypeInfo(_env), exportInfo.FP.Get_processInfo(_env), exportInfo.FP.Get_provideInfo(_env).FP.Get_symbolKind(_env), ast.FP.Get_builtinFunc(_env), option.UseLuneModule, option.TargetLuaVer, option.Testing, option.UseIpairs)
    return conv
}

// 515: decl @lune.@base.@front.ast2Lua
func front_ast2Lua_13_(_env *LnsEnv, ast *AstInfo_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,option *Option_Option) {
    var conv *ConvLua_FilterInfo
    conv = front_ast2LuaMain_12_(_env, ast, streamName, stream, metaStream, convMode, inMacro, option)
    conv.FP.OutputLuaAndMeta(_env, Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode))
}

// 524: decl @lune.@base.@front.loadFromChunk
func front_loadFromChunk_14_(_env *LnsEnv, chunk LnsAny,err LnsAny) LnsAny {
    if err != nil{
        err_90 := err.(string)
        Util_errorLog(_env, err_90)
    }
    var ret LnsAny
    Lns_LockEnvSync( _env, func () {
        if chunk != nil{
            chunk_94 := chunk.(*Lns_luaValue)
            {
                _work := front_convExp1_1043(Lns_2DDD(_env.GetVM().RunLoadedfunc(chunk_94,Lns_2DDD([]LnsAny{}))[0]))
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

// 543: decl @lune.@base.@front.loadFromLuaTxt
func Front_loadFromLuaTxt(_env *LnsEnv, txt string) LnsAny {
    var chunk LnsAny
    var err LnsAny
    Lns_LockEnvSync( _env, func () {
        chunk, err = _env.GetVM().Load(txt, nil)
    })
    return front_loadFromChunk_14_(_env, chunk, err)
}

// 552: decl @lune.@base.@front.byteCompileFromLuaTxt
func front_byteCompileFromLuaTxt_16_(_env *LnsEnv, txt string,stripDebugInfo bool) string {
    var ret string
        var chunk LnsAny
        var err LnsAny
        chunk,err = _env.GetVM().Load(txt, nil)
        if chunk != nil{
            chunk_109 := chunk.(*Lns_luaValue)
            ret = _env.GetVM().String_dump(chunk_109, stripDebugInfo)
        } else {
            panic(Lns_unwrapDefault( err, "load error").(string))
        }
    return ret
}

// 632: decl @lune.@base.@front.getMetaInfo
func front_getMetaInfo_19_(_env *LnsEnv, lnsPath string,mod string,outdir LnsAny)(LnsAny, string, string) {
    var moduleMetaPath string
    moduleMetaPath = lnsPath
    if outdir != nil{
        outdir_131 := outdir.(string)
        moduleMetaPath = _env.GetVM().String_format("%s/%s", []LnsAny{outdir_131, Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string)})
    }
    moduleMetaPath = front_convExp1_1445(Lns_2DDD(_env.GetVM().String_gsub(moduleMetaPath,"%.lns$", ".meta")))
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

// 701: decl @lune.@base.@front.getModuleId
func front_getModuleId_20_(_env *LnsEnv, lnsPath string,mod string,outdir LnsAny,metaInfo LnsAny) *FrontInterface_ModuleId {
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
        metaInfo = front_convExp2_41(Lns_2DDD(front_getMetaInfo_19_(_env, lnsPath, mod, outdir)))
    }
    if metaInfo != nil{
        metaInfo_7 := metaInfo.(*front_MetaForBuildId)
        var buildId *FrontInterface_ModuleId
        buildId = metaInfo_7.FP.CreateModuleId(_env)
        buildCount = buildId.FP.Get_buildCount(_env)
    }
    return FrontInterface_ModuleId_createId(_env, fileTime, buildCount)
}

// 1378: decl @lune.@base.@front.closeStreams
func front_closeStreams_22_(_env *LnsEnv, stream LnsAny,metaStream LnsAny,dependStream LnsAny,metaPath string,saveMetaFlag bool) {
    var checkDiff func(_env *LnsEnv, oldStream *Util_TxtStream,newStream *Util_TxtStream)(bool, string)
    checkDiff = func(_env *LnsEnv, oldStream *Util_TxtStream,newStream *Util_TxtStream)(bool, string) {
        __func__ := "@lune.@base.@front.closeStreams.checkDiff"
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
                oldLine_23 := oldLine.(string)
                if len(oldBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(oldLine_23,"^_moduleObj.__buildId", nil, nil))){
                        oldBuildIdLine = oldLine_23
                    }
                }
            }
            
            if newLine != nil{
                newLine_27 := newLine.(string)
                if len(newBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(newLine_27,"^_moduleObj.__buildId", nil, nil))){
                        newBuildIdLine = newLine_27
                    }
                }
            }
            
            if newLine != oldLine{
                var cont bool
                cont = false
                if newLine != nil && oldLine != nil{
                    newLine_33 := newLine.(string)
                    oldLine_34 := oldLine.(string)
                    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(oldLine_34,"^_moduleObj.__buildId", nil, nil))){
                        if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(newLine_33,"^_moduleObj.__buildId", nil, nil))){
                            tailBeginPos = newStream.FP.Get_lineNo(_env)
                            cont = true
                        }
                    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_car(_env.GetVM().String_find(oldLine_34,"^__dependModuleMap.*buildId =", nil, nil))) &&
                        _env.SetStackVal( Lns_car(_env.GetVM().String_find(newLine_33,"^__dependModuleMap.*buildId =", nil, nil))) )){
                        var oldSub string
                        oldSub = front_convExp4_279(Lns_2DDD(_env.GetVM().String_gsub(oldLine_34,"buildId =.*", "")))
                        var newSub string
                        newSub = front_convExp4_292(Lns_2DDD(_env.GetVM().String_gsub(newLine_33,"buildId =.*", "")))
                        if oldSub == newSub{
                            cont = true
                        }
                    }
                }
                if Lns_op_not(cont){
                    Log_log(_env, Log_Level__Debug, __func__, 1449, Log_CreateMessage(func(_env *LnsEnv) string {
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
                    oldBuildId = closeStreams__txt2ModuleId_0_(_env, oldBuildIdLine)
                    var newBuildId *FrontInterface_ModuleId
                    newBuildId = closeStreams__txt2ModuleId_0_(_env, newBuildIdLine)
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
        stream_54 := stream.(Lns_oStream)
        stream_54.Close(_env)
    }
    if dependStream != nil{
        dependStream_56 := dependStream.(Lns_oStream)
        dependStream_56.Close(_env)
    }
    if metaStream != nil{
        metaStream_58 := metaStream.(*Util_memStream)
        if saveMetaFlag{
            var newMetaTxt string
            newMetaTxt = metaStream_58.FP.Get_txt(_env)
            var oldMetaTxt string
            oldMetaTxt = ""
            {
                _oldFileObj := front_convExp4_568(Lns_2DDD(Lns_io_open(metaPath, nil)))
                if !Lns_IsNil( _oldFileObj ) {
                    oldFileObj := _oldFileObj.(Lns_luaStream)
                    oldMetaTxt = Lns_unwrapDefault( oldFileObj.Read(_env, "*a"), "").(string)
                    oldFileObj.Close(_env)
                }
            }
            var sameFlag bool
            var txt string
            sameFlag,txt = checkDiff(_env, NewUtil_TxtStream(_env, oldMetaTxt), NewUtil_TxtStream(_env, newMetaTxt))
            var saveMeta func(_env *LnsEnv, meta string)
            saveMeta = func(_env *LnsEnv, meta string) {
                {
                    _fileObj := front_convExp4_639(Lns_2DDD(Lns_io_open(metaPath, "w")))
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
                saveMeta(_env, newMetaTxt)
            } else { 
                if txt != ""{
                    saveMeta(_env, txt)
                }
            }
        }
    }
}


// 1638: decl @lune.@base.@front.outputDependInfo
func front_outputDependInfo_24_(_env *LnsEnv, stream LnsAny,metaInfo *front_MetaForBuildId,mod string) {
    if stream != nil{
        stream_95 := stream.(Lns_oStream)
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
        dependInfo.FP.Output(_env, stream_95)
    }
}


// 1908: decl @lune.@base.@front.convertLnsCode2LuaCode
func Front_convertLnsCode2LuaCode(_env *LnsEnv, lnsCode string,path string,baseDir LnsAny) string {
    var option *Option_Option
    option = NewOption_Option(_env)
    option.ScriptPath = path
    option.UseLuneModule = Option_getRuntimeModule(_env)
    option.UseIpairs = true
    var front *front_Front
    front = Newfront_Front(_env, option, nil)
    return front.FP.ConvertLns2LuaCode(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsCode{lnsCode, path, nil}, baseDir, NewUtil_TxtStream(_env, lnsCode).FP, path)
}

// 2008: decl @lune.@base.@front.build
func Front_build(_env *LnsEnv, option *Option_Option,astCallback Front_AstCallback) {
    var front *front_Front
    Lns_LockEnvSync( _env, func () {
        front = Newfront_Front(_env, option, nil)
    })
    front.FP.Build(_env, front_BuildMode__CreateAst_Obj, astCallback)
}

// 2135: decl @lune.@base.@front.exec
func Front_exec(_env *LnsEnv, args *LnsList) {
    var version LnsReal
    version = Lns_unwrapDefault( Lns_tonumber(Lns_car(_env.GetVM().String_gsub(Depend_getLuaVersion(_env),"^[^%d]+", "")).(string), nil), 0.0).(LnsReal)
    if version < 5.1{
        Lns_io_stderr.Write(_env, _env.GetVM().String_format("LuneScript doesn't support this lua version(%s). %s\n", []LnsAny{version, "please use the version >= 5.1."}))
        _env.GetVM().OS_exit(1)
    }
    var option *Option_Option
    option = Option_analyze(_env, args)
    var front *front_Front
    front = Newfront_Front(_env, option, nil)
    front.FP.Exec(_env)
}

// 2154: decl @lune.@base.@front.setFront
func Front_setFront(_env *LnsEnv, bindModuleList *LnsList) {
    var option *Option_Option
    option = Option_createDefaultOption(_env, NewLnsList([]LnsAny{"dummy.lns"}), nil)
    Newfront_Front(_env, option, bindModuleList)
}

// 2159: decl @lune.@base.@front.__main
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





func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_0_(_env *LnsEnv) string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1_(_env *LnsEnv) string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_2_(_env *LnsEnv) string {
    return "NeedUpdate"
}


func Front_getModuleIdAndCheckUptodate___anonymous_1_(_env *LnsEnv) string {
    return "not found meta"
}
















// 1381: decl @lune.@base.@front.closeStreams.txt2ModuleId
func closeStreams__txt2ModuleId_0_(_env *LnsEnv, txt string) *FrontInterface_ModuleId {
    var buildIdTxt string
    buildIdTxt = front_convExp4_28(Lns_2DDD(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(txt,"^_moduleObj.__buildId = ", "")).(string),"\"", "")))
    return FrontInterface_ModuleId_createIdFromTxt(_env, buildIdTxt)
}










func Front_exec___anonymous_1_(_env *LnsEnv, ast *AstInfo_ASTInfo) {
    Lns_print([]LnsAny{ast.FP.Get_streamName(_env)})
}
func Front_exec___anonymous_2_(_env *LnsEnv, ast *AstInfo_ASTInfo) {
    var indexer *NodeIndexer_Indexer
    indexer = NewNodeIndexer_Indexer(_env, ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env))
    indexer.FP.Start(_env, ast.FP.Get_node(_env), NewLnsSet([]LnsAny{Nodes_NodeKind_get_Switch(_env), Nodes_NodeKind_get_Match(_env), Nodes_NodeKind_get_For(_env), Nodes_NodeKind_get_Apply(_env)}))
    indexer.FP.Dump(_env)
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

// declaration Class -- AstCreater
type front_AstCreaterMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    createAst(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 LnsInt, arg7 LnsAny) *AstInfo_ASTInfo
    getAst(_env *LnsEnv)(*AstInfo_ASTInfo, *FrontInterface_ModuleInfo, *FrontInterface_ModuleMeta)
    getExportInfo(_env *LnsEnv) LnsAny
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    Start(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny)
}
type front_AstCreater struct {
    Runner_Runner
    option *Option_Option
    builtinFunc *Builtin_BuiltinFuncType
    mod string
    moduleId *FrontInterface_ModuleId
    ast LnsAny
    moduleInfo LnsAny
    converter front_ConverterFunc_6_
    exportInfo LnsAny
    exportInfoReadyFlag LnsAny
    FP front_AstCreaterMtd
}
func front_AstCreater2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_AstCreater).FP
}
type front_AstCreaterDownCast interface {
    Tofront_AstCreater() *front_AstCreater
}
func front_AstCreaterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(front_AstCreaterDownCast)
    if ok { return work.Tofront_AstCreater() }
    return nil
}
func (obj *front_AstCreater) Tofront_AstCreater() *front_AstCreater {
    return obj
}
func Newfront_AstCreater(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 LnsAny, arg5 *FrontInterface_ModuleId, arg6 LnsAny, arg7 LnsInt, arg8 LnsAny, arg9 *Builtin_BuiltinFuncType, arg10 *Option_Option) *front_AstCreater {
    obj := &front_AstCreater{}
    obj.FP = obj
    obj.Runner_Runner.FP = obj
    obj.Initfront_AstCreater(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}

// declaration Class -- Front
type front_FrontMtd interface {
    applyAstResult(_env *LnsEnv, arg1 LnsAny) *AstInfo_ASTInfo
    Build(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny)
    CheckDiag(_env *LnsEnv, arg1 string)
    checkUptodateMeta(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny, arg4 LnsAny) LnsAny
    Complete(_env *LnsEnv, arg1 string)
    convertFromAst(_env *LnsEnv, arg1 *AstInfo_ASTInfo, arg2 string, arg3 LnsInt)(string, string)
    ConvertLns2LuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 Lns_iStream, arg5 string) string
    convertToLua(_env *LnsEnv, arg1 string, arg2 LnsInt, arg3 Lns_oStream, arg4 Lns_oStream) LnsAny
    createAst(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 string, arg5 *FrontInterface_ModuleId, arg6 LnsAny, arg7 LnsInt, arg8 LnsAny) *AstInfo_ASTInfo
    createAstSub(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 string, arg5 *FrontInterface_ModuleId, arg6 LnsAny, arg7 LnsInt, arg8 LnsAny) LnsAny
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
    GetLuaModulePath(_env *LnsEnv, arg1 string, arg2 LnsAny)(string, LnsAny, string)
    getModuleIdAndCheckUptodate(_env *LnsEnv, arg1 string, arg2 string)(*FrontInterface_ModuleId, LnsAny)
    Inquire(_env *LnsEnv, arg1 string)
    loadFile(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string)(*FrontInterface_ModuleMeta, LnsAny)
    LoadFromLnsTxt(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string) LnsAny
    loadLua(_env *LnsEnv, arg1 string) LnsAny
    LoadMeta(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string, arg4 LnsAny, arg5 FrontInterface_ModuleLoader) LnsAny
    LoadModule(_env *LnsEnv, arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
    loadParserToLuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string, arg5 LnsAny)(*FrontInterface_ModuleMeta, string)
    outputBootC(_env *LnsEnv, arg1 string)
    OutputBuiltin(_env *LnsEnv, arg1 string)
    SaveToC(_env *LnsEnv, arg1 string, arg2 *AstInfo_ASTInfo)
    saveToGo(_env *LnsEnv, arg1 string, arg2 LnsAny) LnsAny
    saveToLua(_env *LnsEnv, arg1 *front_UpdateInfo) LnsAny
    scriptPath2Module(_env *LnsEnv, arg1 string)(string, LnsAny)
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

// declaration Class -- LuaConverter
type front_LuaConverterMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    SaveLua(_env *LnsEnv)
    Start(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny)
}
type front_LuaConverter struct {
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
    converterFunc front_ConverterFunc_6_
    filterInfo LnsAny
    FP front_LuaConverterMtd
}
func front_LuaConverter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_LuaConverter).FP
}
type front_LuaConverterDownCast interface {
    Tofront_LuaConverter() *front_LuaConverter
}
func front_LuaConverterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(front_LuaConverterDownCast)
    if ok { return work.Tofront_LuaConverter() }
    return nil
}
func (obj *front_LuaConverter) Tofront_LuaConverter() *front_LuaConverter {
    return obj
}
func Newfront_LuaConverter(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny, arg4 LnsAny, arg5 LnsInt, arg6 string, arg7 bool, arg8 bool, arg9 *Option_Option) *front_LuaConverter {
    obj := &front_LuaConverter{}
    obj.FP = obj
    obj.Runner_Runner.FP = obj
    obj.Initfront_LuaConverter(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}

// declaration Class -- GoConverter
type front_GoConverterMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    saveGo(_env *LnsEnv)
    Start(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny)
}
type front_GoConverter struct {
    Runner_Runner
    memStream *Util_memStream
    path string
    converter front_ConverterFunc_6_
    validFlag bool
    FP front_GoConverterMtd
}
func front_GoConverter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*front_GoConverter).FP
}
type front_GoConverterDownCast interface {
    Tofront_GoConverter() *front_GoConverter
}
func front_GoConverterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(front_GoConverterDownCast)
    if ok { return work.Tofront_GoConverter() }
    return nil
}
func (obj *front_GoConverter) Tofront_GoConverter() *front_GoConverter {
    return obj
}
func Newfront_GoConverter(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 *Option_Option, arg4 *ConvGo_Option) *front_GoConverter {
    obj := &front_GoConverter{}
    obj.FP = obj
    obj.Runner_Runner.FP = obj
    obj.Initfront_GoConverter(_env, arg1, arg2, arg3, arg4)
    return obj
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
    Lns_LuneControl_init(_env)
    Lns_Nodes_init(_env)
    Lns_Ast_init(_env)
    Lns_Runner_init(_env)
    Lns_Builtin_init(_env)
    Lns_NodeIndexer_init(_env)
    Depend_setup(_env, Depend_UpdateVer(front__anonymous_0_))
    front_forceUpdateMeta = true
}
func init() {
    init_front = false
}
// 102: DeclConstr
func (self *front_ModuleMgr) Initfront_ModuleMgr(_env *LnsEnv) {
    self.mod2info = NewUtil_OrderdMap(_env)
    self.loadedMetaMap = NewLnsMap( map[LnsAny]LnsAny{})
}
// 107: decl @lune.@base.@front.ModuleMgr.get
func (self *front_ModuleMgr) Get(_env *LnsEnv, mod string) LnsAny {
    return self.mod2info.FP.Get_map(_env).Get(mod)
}
// 111: decl @lune.@base.@front.ModuleMgr.getAst
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
// 125: decl @lune.@base.@front.ModuleMgr.getModList
func (self *front_ModuleMgr) GetModList(_env *LnsEnv) *LnsList {
    return self.mod2info.FP.Get_keyList(_env)
}
// 130: decl @lune.@base.@front.ModuleMgr.add
func (self *front_ModuleMgr) Add(_env *LnsEnv, ast *AstInfo_ASTInfo,moduleInfo *FrontInterface_ModuleInfo) {
    var mod string
    mod = moduleInfo.FP.Get_fullName(_env)
    self.mod2info.FP.Add(_env, mod, &front_UptodateInfo__Update{moduleInfo, ast})
}
// 149: decl @lune.@base.@front.ModuleMgr.addMeta
func (self *front_ModuleMgr) AddMeta(_env *LnsEnv, mod string,meta *FrontInterface_ModuleMeta) {
    if Lns_op_not(self.FP.Get(_env, mod)){
        self.mod2info.FP.Add(_env, mod, &front_UptodateInfo__Uptodate{meta})
    }
    self.loadedMetaMap.Set(mod,meta)
}
// 156: decl @lune.@base.@front.ModuleMgr.getMeta
func (self *front_ModuleMgr) GetMeta(_env *LnsEnv, mod string) LnsAny {
    return self.loadedMetaMap.Get(mod)
}
// 178: decl @lune.@base.@front.AstCreater.createAst
func (self *front_AstCreater) createAst(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,baseDir LnsAny,stdinFile LnsAny,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *AstInfo_ASTInfo {
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
// 201: DeclConstr
func (self *front_AstCreater) Initfront_AstCreater(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,mod string,baseDir LnsAny,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny,builtinFunc *Builtin_BuiltinFuncType,option *Option_Option) {
    self.InitRunner_Runner(_env)
    self.option = option
    self.builtinFunc = builtinFunc
    self.mod = mod
    self.moduleId = moduleId
    self.moduleInfo = nil
    self.exportInfo = nil
    self.ast = nil
    self.exportInfoReadyFlag = LnsCreateSyncFlag(_env)
    self.converter = front_ConverterFunc_6_(func(_env *LnsEnv) {
        __func__ := "@lune.@base.@front.AstCreater.__init.<anonymous>"
        var ast *AstInfo_ASTInfo
        ast = self.FP.createAst(_env, importModuleInfo, parserSrc, baseDir, option.FP.Get_stdinFile(_env), analyzeModule, analyzeMode, pos)
        self.ast = ast
        self.moduleInfo = front_createModuleInfo_4_(_env, ast, self.mod, self.moduleId)
        Log_log(_env, Log_Level__Log, __func__, 224, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.GetVM().String_format("generated AST -- %s", []LnsAny{mod})
        }))
        
    })
    var lnsPath string
    switch _matchExp0 := parserSrc.(type) {
    case *Types_ParserSrc__LnsCode:
    path := _matchExp0.Val2
        lnsPath = path
    case *Types_ParserSrc__LnsPath:
    path := _matchExp0.Val1
        lnsPath = path
    case *Types_ParserSrc__Parser:
    path := _matchExp0.Val1
        lnsPath = path
    }
    self.FP.Start(_env, 0, _env.GetVM().String_format("createAst - %s", []LnsAny{lnsPath}))
}
// 244: decl @lune.@base.@front.AstCreater.runMain
func (self *front_AstCreater) RunMain(_env *LnsEnv) {
    self.converter(_env)
}
// 249: decl @lune.@base.@front.AstCreater.getAst
func (self *front_AstCreater) getAst(_env *LnsEnv)(*AstInfo_ASTInfo, *FrontInterface_ModuleInfo, *FrontInterface_ModuleMeta) {
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
// 261: decl @lune.@base.@front.AstCreater.getExportInfo
func (self *front_AstCreater) getExportInfo(_env *LnsEnv) LnsAny {
    __func__ := "@lune.@base.@front.AstCreater.getExportInfo"
    {
        __exp := self.exportInfoReadyFlag
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(G__lns_Flag)
            _exp.Wait(_env)
        }
    }
    if Lns_op_not(self.exportInfo){
        Log_log(_env, Log_Level__Err, __func__, 266, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.GetVM().String_format("exportInfo is nil -- %s", []LnsAny{self.mod})
        }))
        
    }
    return self.exportInfo
}
// 305: DeclConstr
func (self *front_Front) Initfront_Front(_env *LnsEnv, option *Option_Option,bindModuleList LnsAny) {
    self.mod2loader = NewLnsMap( map[LnsAny]LnsAny{})
    self.mod2astCreate = NewLnsMap( map[LnsAny]LnsAny{})
    self.loadCount = 0
    self.targetSet = NewLnsSet([]LnsAny{})
    self.bindModuleSet = NewLnsSet([]LnsAny{})
    if bindModuleList != nil{
        bindModuleList_171 := bindModuleList.(*LnsList)
        for _, _mod := range( bindModuleList_171.Items ) {
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
// 349: decl @lune.@base.@front.Front.getLoadInfo
func (self *front_Front) getLoadInfo(_env *LnsEnv, mod string) LnsAny {
    if self.option.Testing{
        return self.loadedMapTest.Get(mod)
    }
    return self.loadedMap.Get(mod)
}
// 356: decl @lune.@base.@front.Front.setLoadInfo
func (self *front_Front) setLoadInfo(_env *LnsEnv, mod string,info *front_LoadInfo) {
    if self.option.Testing{
        self.loadedMapTest.Set(mod,info)
    }
    self.loadedMap.Set(mod,info)
}
// 376: decl @lune.@base.@front.Front.error
func (self *front_Front) Error(_env *LnsEnv, message string) {
    Util_errorLog(_env, message)
    Util_printStackTrace(_env)
    _env.GetVM().OS_exit(1)
}
// 382: decl @lune.@base.@front.Front.loadLua
func (self *front_Front) loadLua(_env *LnsEnv, path string) LnsAny {
    var ret LnsAny
    Lns_LockEnvSync( _env, func () {
        var chunk LnsAny
        var err LnsAny
        chunk,err = _env.GetVM().Loadfile(path)
        if chunk != nil{
            chunk_48 := chunk.(*Lns_luaValue)
            ret = Lns_unwrap( Lns_car(_env.GetVM().RunLoadedfunc(chunk_48,Lns_2DDD([]LnsAny{}))[0]))
        } else {
            Util_errorLog(_env, Lns_unwrapDefault( err, _env.GetVM().String_format("load error -- %s.", []LnsAny{path})).(string))
            ret = nil
        }
    })
    return ret
}
// 402: decl @lune.@base.@front.Front.scriptPath2Module
func (self *front_Front) scriptPath2Module(_env *LnsEnv, path string)(string, LnsAny) {
    var mod string
    mod = Util_scriptPath2ModuleFromProjDir(_env, path, self.option.FP.Get_projDir(_env))
    return mod, self.option.FP.Get_projDir(_env)
}
// 407: decl @lune.@base.@front.Front.createPaser
func (self *front_Front) createPaser(_env *LnsEnv, scriptPath string) *Parser_Parser {
    var mod string
    mod = front_convExp1_590(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath)))
    return front_createPaser_9_(_env, scriptPath, mod, self.option.FP.Get_stdinFile(_env))
}
// 430: decl @lune.@base.@front.Front.createAstSub
func (self *front_Front) createAstSub(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,baseDir LnsAny,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) LnsAny {
    {
        __exp := self.moduleMgr.FP.Get(_env, mod)
        if !Lns_IsNil( __exp ) {
            _exp := __exp
            switch _matchExp0 := _exp.(type) {
            case *front_UptodateInfo__Update:
            ast := _matchExp0.Val2
                return &front_CreateAstResult__Ast{ast}
            case *front_UptodateInfo__Uptodate:
                Util_err(_env, _env.GetVM().String_format("can't load the multiple module -- %s", []LnsAny{mod}))
            }
        }
    }
    {
        _creater := self.mod2astCreate.Get(mod)
        if !Lns_IsNil( _creater ) {
            creater := _creater.(*front_AstCreater)
            var ast *AstInfo_ASTInfo
            ast = front_convExp1_681(Lns_2DDD(creater.FP.getAst(_env)))
            return &front_CreateAstResult__Ast{ast}
        }
    }
    var astCreater *front_AstCreater
    astCreater = Newfront_AstCreater(_env, importModuleInfo, parserSrc, mod, baseDir, moduleId, analyzeModule, analyzeMode, pos, self.builtinFunc, self.option)
    self.mod2astCreate.Set(mod,astCreater)
    return &front_CreateAstResult__Creater{astCreater}
}
// 466: decl @lune.@base.@front.Front.applyAstResult
func (self *front_Front) applyAstResult(_env *LnsEnv, result LnsAny) *AstInfo_ASTInfo {
    __func__ := "@lune.@base.@front.Front.applyAstResult"
    switch _matchExp0 := result.(type) {
    case *front_CreateAstResult__Ast:
    ast := _matchExp0.Val1
        return ast
    case *front_CreateAstResult__Creater:
    astCreater := _matchExp0.Val1
        var ast *AstInfo_ASTInfo
        var moduleInfo *FrontInterface_ModuleInfo
        var meta *FrontInterface_ModuleMeta
        ast,moduleInfo,meta = astCreater.FP.getAst(_env)
        Log_log(_env, Log_Level__Log, __func__, 473, Log_CreateMessage(func(_env *LnsEnv) string {
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
// 487: decl @lune.@base.@front.Front.createAst
func (self *front_Front) createAst(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,baseDir LnsAny,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *AstInfo_ASTInfo {
    return self.FP.applyAstResult(_env, self.FP.createAstSub(_env, importModuleInfo, parserSrc, baseDir, mod, moduleId, analyzeModule, analyzeMode, pos))
}
// 566: decl @lune.@base.@front.Front.convertFromAst
func (self *front_Front) convertFromAst(_env *LnsEnv, ast *AstInfo_ASTInfo,streamName string,mode LnsInt)(string, string) {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var metaStream *Util_memStream
    metaStream = NewUtil_memStream(_env)
    front_ast2Lua_13_(_env, ast, streamName, stream.FP, metaStream.FP, mode, false, self.option)
    return metaStream.FP.Get_txt(_env), stream.FP.Get_txt(_env)
}
// 580: decl @lune.@base.@front.Front.loadFromLnsTxt
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
// 644: decl @lune.@base.@front.Front.searchModuleFile
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
        outputDir_142 := outputDir.(string)
        lnsSearchPath = _env.GetVM().String_format("%s/?%s;%s", []LnsAny{outputDir_142, suffix, Lns_package_path})
    }
    if baseDir != nil{
        baseDir_144 := baseDir.(string)
        lnsSearchPath = _env.GetVM().String_format("%s/?%s;%s", []LnsAny{baseDir_144, suffix, Lns_package_path})
    }
    {
        _projDir := self.option.FP.Get_projDir(_env)
        if !Lns_IsNil( _projDir ) {
            projDir := _projDir.(string)
            lnsSearchPath = _env.GetVM().String_format("%s;%s", []LnsAny{Util_pathJoin(_env, projDir, _env.GetVM().String_format("?%s", []LnsAny{suffix})), Lns_package_path})
        }
    }
    lnsSearchPath = front_convExp1_1614(Lns_2DDD(_env.GetVM().String_gsub(lnsSearchPath,"%.lua$", suffix)))
    lnsSearchPath = front_convExp1_1633(Lns_2DDD(_env.GetVM().String_gsub(lnsSearchPath,"%.lua;", suffix + ";")))
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
                    Log_log(_env, Log_Level__Err, __func__, 680, Log_CreateMessage(func(_env *LnsEnv) string {
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
// 741: decl @lune.@base.@front.Front.getModuleIdAndCheckUptodate
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
    var checkDependUptodate func(_env *LnsEnv, metaTime LnsReal,metaInfo *front_MetaForBuildId,metaCode string) LnsAny
    checkDependUptodate = func(_env *LnsEnv, metaTime LnsReal,metaInfo *front_MetaForBuildId,metaCode string) LnsAny {
        __func__ := "@lune.@base.@front.Front.getModuleIdAndCheckUptodate.checkDependUptodate"
        for _depMod, _dependItem := range( metaInfo.G__dependModuleMap.Items ) {
            depMod := _depMod.(string)
            dependItem := _dependItem.(front_DependMetaInfoDownCast).Tofront_DependMetaInfo()
            var modMetaPath string
            
            {
                _modMetaPath := self.FP.searchModuleFile(_env, depMod, ".meta", nil, self.option.OutputDir)
                if _modMetaPath == nil{
                    Log_log(_env, Log_Level__Debug, __func__, 777, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_0_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    modMetaPath = _modMetaPath.(string)
                }
            }
            var time LnsReal
            
            {
                _time := Depend_getFileLastModifiedTime(_env, modMetaPath)
                if _time == nil{
                    Log_log(_env, Log_Level__Debug, __func__, 782, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1_))
                    
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
                        Log_log(_env, Log_Level__Debug, __func__, 790, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_2_))
                        
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
                    Log_log(_env, Log_Level__Debug, __func__, 800, Log_CreateMessage(func(_env *LnsEnv) string {
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
    metaInfo,metaPath,metaCode = front_getMetaInfo_19_(_env, lnsPath, mod, self.option.OutputDir)
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
                        uptodate = checkDependUptodate(_env, metaTime_60, metaInfo_52, metaCode)
                    }
                }
            }
        } else { 
        }
    } else {
        Log_log(_env, Log_Level__Debug, __func__, 838, Log_CreateMessage(Front_getModuleIdAndCheckUptodate___anonymous_1_))
        
    }
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_20_(_env, lnsPath, mod, self.option.OutputDir, metaInfo)
    if moduleId == FrontInterface_ModuleId__tempId{
        Util_err(_env, _env.GetVM().String_format("not found -- %s", []LnsAny{lnsPath}))
    }
    return moduleId, uptodate
}
// 851: decl @lune.@base.@front.Front.convertLns2LuaCode
func (self *front_Front) ConvertLns2LuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,baseDir LnsAny,stream Lns_iStream,streamName string) string {
    var mod string
    mod = front_convExp2_648(Lns_2DDD(self.FP.scriptPath2Module(_env, streamName)))
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, parserSrc, baseDir, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(_env, ast, streamName, ConvLua_ConvMode__ConvMeta)
    return luaTxt
}
// 865: decl @lune.@base.@front.Front.getGoAppName
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
// 872: decl @lune.@base.@front.Front.createGoOption
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
                parentPath = front_convExp2_777(Lns_2DDD(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(scriptPath,"/[^/]+$", "")).(string),".*/", "")))
                if len(parentPath) == 0{
                    packageName = "main"
                } else if parentPath == "."{
                    packageName = "main"
                } else if parentPath == ".."{
                    packageName = "main"
                } else { 
                    packageName = front_convExp2_835(Lns_2DDD(_env.GetVM().String_gsub(parentPath,"[^%w]", "")))
                }
            }
        }
    }
    return NewConvGo_Option(_env, packageName, self.FP.getGoAppName(_env), self.option.MainModule, self.option.FP.Get_addEnvArg(_env), self.option.FP.Get_enableRunner(_env))
}
// 907: decl @lune.@base.@front.Front.loadParserToLuaCode
func (self *front_Front) loadParserToLuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,path string,mod string,baseDir LnsAny)(*FrontInterface_ModuleMeta, string) {
    __func__ := "@lune.@base.@front.Front.loadParserToLuaCode"
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_20_(_env, path, mod, nil, nil)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, parserSrc, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(_env, ast, path, ConvLua_ConvMode__ConvMeta)
    Log_log(_env, Log_Level__Trace, __func__, 918, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("Meta = %s", []LnsAny{metaTxt})
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
                    newpath = _env.GetVM().String_format("%s/%s%s", []LnsAny{dir, Lns_car(_env.GetVM().String_gsub(mod,"%.", "/")).(string), suffix})
                } else {
                    newpath = front_convExp2_1041(Lns_2DDD(_env.GetVM().String_gsub(path,".lns$", suffix)))
                }
            }
            var saveTxt string
            saveTxt = txt
            if byteCompile{
                saveTxt = front_byteCompileFromLuaTxt_16_(_env, saveTxt, stripDebugInfo)
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(front_forceUpdateMeta)) &&
                _env.SetStackVal( checkUpdate) ).(bool)){
                {
                    _fileObj := front_convExp2_1110(Lns_2DDD(Lns_io_open(newpath, nil)))
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
                _fileObj := front_convExp2_1139(Lns_2DDD(Lns_io_open(newpath, "w")))
                if !Lns_IsNil( _fileObj ) {
                    fileObj := _fileObj.(Lns_luaStream)
                    fileObj.Write(_env, saveTxt)
                    fileObj.Close(_env)
                }
            }
        }
        saveFile(_env, ".lua", luaTxt, self.option.ByteCompile, self.option.StripDebugInfo, false)
        saveFile(_env, ".meta", metaTxt, self.option.ByteCompile, true, true)
        if _switch0 := self.option.ConvTo; _switch0 == Types_Lang__Go {
            var memStream *Util_memStream
            memStream = NewUtil_memStream(_env)
            var conv *Nodes_Filter
            conv = ConvGo_createFilter(_env, self.option.Testing, path, memStream.FP, ast, self.FP.createGoOption(_env, path))
            ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast.FP.Get_node(_env))))
            saveFile(_env, ".go", memStream.FP.Get_txt(_env), false, false, false)
        }
    }
    return Lns_unwrap( self.moduleMgr.FP.GetMeta(_env, mod)).(*FrontInterface_ModuleMeta), luaTxt
}
// 979: decl @lune.@base.@front.Front.loadFile
func (self *front_Front) loadFile(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,baseDir LnsAny,path string,mod string)(*FrontInterface_ModuleMeta, LnsAny) {
    __func__ := "@lune.@base.@front.Front.loadFile"
    Log_log(_env, Log_Level__Info, __func__, 983, Log_CreateMessage(func(_env *LnsEnv) string {
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
// 996: decl @lune.@base.@front.Front.searchModule
func (self *front_Front) SearchModule(_env *LnsEnv, mod string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(_env, mod, ".lns", baseDir, addSearchPath)
}
// 1000: decl @lune.@base.@front.Front.searchLuaFile
func (self *front_Front) searchLuaFile(_env *LnsEnv, moduleFullName string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(_env, moduleFullName, ".lua", baseDir, addSearchPath)
}
// 1012: decl @lune.@base.@front.Front.checkUptodateMeta
func (self *front_Front) checkUptodateMeta(_env *LnsEnv, lnsPath string,metaPath string,baseDir LnsAny,addSearchPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.checkUptodateMeta"
    var metaObj LnsAny
    
    {
        _metaObj := self.FP.loadLua(_env, metaPath)
        if _metaObj == nil{
            Log_log(_env, Log_Level__Warn, __func__, 1017, Log_CreateMessage(func(_env *LnsEnv) string {
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
    Lns_LockEnvSync( _env, func () {
        meta = metaObj.(*Lns_luaValue)
        if meta.GetAt( "__formatVersion" ).(string) != Ver_metaVersion{
            Log_log(_env, Log_Level__Warn, __func__, 1026, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.GetVM().String_format("unmatch meta version -- %s", []LnsAny{metaPath})
            }))
            
            valid = false
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( valid) &&
            _env.SetStackVal( meta.GetAt( "__hasTest" ).(bool)) ).(bool)){
            if meta.GetAt( "__enableTest" ).(bool) != self.option.Testing{
                Log_log(_env, Log_Level__Warn, __func__, 1032, Log_CreateMessage(func(_env *LnsEnv) string {
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
                                        Log_log(_env, Log_Level__Warn, __func__, 1046, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.GetVM().String_format("not ready -- %s, %s", []LnsAny{moduleLnsPath, metaPath})
                                        }))
                                        
                                        valid = false
                                        break
                                    }
                                    var moduleMetaPath string
                                    moduleMetaPath = front_convExp2_1755(Lns_2DDD(_env.GetVM().String_gsub(moduleLuaPath,"%.lua$", ".meta")))
                                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                        _env.SetStackVal( Depend_existFile(_env, moduleMetaPath)) &&
                                        _env.SetStackVal( Lns_op_not(Util_getReadyCode(_env, moduleMetaPath, metaPath))) ).(bool)){
                                        Log_log(_env, Log_Level__Warn, __func__, 1055, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.GetVM().String_format("not ready -- %s, %s", []LnsAny{moduleMetaPath, metaPath})
                                        }))
                                        
                                        valid = false
                                        break
                                    }
                                } else {
                                    Log_log(_env, Log_Level__Warn, __func__, 1061, Log_CreateMessage(func(_env *LnsEnv) string {
                                        return _env.GetVM().String_format("not found .lua file for -- %s", []LnsAny{moduleFullName})
                                    }))
                                    
                                    valid = false
                                    break
                                }
                            }
                        } else {
                            Log_log(_env, Log_Level__Warn, __func__, 1067, Log_CreateMessage(func(_env *LnsEnv) string {
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
// 1087: decl @lune.@base.@front.Front.loadModule
func (self *front_Front) LoadModule(_env *LnsEnv, mod string)(LnsAny, *FrontInterface_ModuleMeta) {
    __func__ := "@lune.@base.@front.Front.loadModule"
    var orgMod string
    orgMod = mod
    var baseDir LnsAny
    baseDir = nil
    _, _, mod = self.gomodMap.FP.GetLuaModulePath(_env, mod, baseDir)
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
                        luaPath = front_convExp3_128(Lns_2DDD(_env.GetVM().String_gsub(lnsPath, "%.lns$", ".lua")))
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
                            luaPath_17 := luaPath.(string)
                            if Util_getReadyCode(_env, lnsPath, luaPath_17){
                                var metaPath string
                                metaPath = front_convExp3_188(Lns_2DDD(_env.GetVM().String_gsub(luaPath_17, "%.lua$", ".meta")))
                                if Util_getReadyCode(_env, lnsPath, metaPath){
                                    {
                                        _preLoadInfo := self.preloadedModMap.Get(mod)
                                        if !Lns_IsNil( _preLoadInfo ) {
                                            preLoadInfo := _preLoadInfo
                                            loadVal = preLoadInfo
                                        } else {
                                            loadVal = self.FP.loadLua(_env, luaPath_17)
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
                            Log_log(_env, Log_Level__Warn, __func__, 1142, Log_CreateMessage(func(_env *LnsEnv) string {
                                return _env.GetVM().String_format("load from the binding -- %s", []LnsAny{mod})
                            }))
                            
                            var workMod LnsAny
                            Lns_LockEnvSync( _env, func () {
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
// 1167: decl @lune.@base.@front.Front.getLuaModulePath
func (self *front_Front) GetLuaModulePath(_env *LnsEnv, mod string,baseDir LnsAny)(string, LnsAny, string) {
    return self.gomodMap.FP.GetLuaModulePath(_env, mod, baseDir)
}
// 1171: decl @lune.@base.@front.Front.loadMeta
func (self *front_Front) LoadMeta(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,mod string,orgMod string,baseDir LnsAny,loader FrontInterface_ModuleLoader) LnsAny {
    __func__ := "@lune.@base.@front.Front.loadMeta"
    {
        _creater := self.mod2astCreate.Get(orgMod)
        if !Lns_IsNil( _creater ) {
            creater := _creater.(*front_AstCreater)
            var exportInfo *FrontInterface_ExportInfo
            
            {
                _exportInfo := creater.FP.getExportInfo(_env)
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
                            luaPath = front_convExp3_706(Lns_2DDD(_env.GetVM().String_gsub(lnsPath, "%.lns$", ".lua")))
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
                                luaPath_68 := luaPath.(string)
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
                                    if Util_getReadyCode(_env, lnsPath, luaPath_68){
                                        var metaPath string
                                        metaPath = front_convExp3_813(Lns_2DDD(_env.GetVM().String_gsub(luaPath_68, "%.lua$", ".meta")))
                                        if Util_getReadyCode(_env, lnsPath, metaPath){
                                            meta = self.FP.checkUptodateMeta(_env, lnsPath, metaPath, baseDir, self.option.OutputDir)
                                        } else { 
                                            Log_log(_env, Log_Level__Warn, __func__, 1243, Log_CreateMessage(func(_env *LnsEnv) string {
                                                return _env.GetVM().String_format("%s not ready meta %s, %s", []LnsAny{orgMod, lnsPath, metaPath})
                                            }))
                                            
                                        }
                                    } else { 
                                        Log_log(_env, Log_Level__Warn, __func__, 1247, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.GetVM().String_format("%s not ready lua %s, %s", []LnsAny{orgMod, lnsPath, luaPath_68})
                                        }))
                                        
                                    }
                                } else { 
                                    Log_log(_env, Log_Level__Warn, __func__, 1251, Log_CreateMessage(func(_env *LnsEnv) string {
                                        return _env.GetVM().String_format("force analyze -- %s", []LnsAny{orgMod})
                                    }))
                                    
                                }
                            } else {
                                Log_log(_env, Log_Level__Warn, __func__, 1255, Log_CreateMessage(func(_env *LnsEnv) string {
                                    return _env.GetVM().String_format("%s not found lua in %s", []LnsAny{orgMod, self.option.OutputDir})
                                }))
                                
                            }
                        }
                        if meta != nil{
                            meta_92 := meta.(*FrontInterface_ModuleMeta)
                            self.moduleMgr.FP.AddMeta(_env, orgMod, meta_92)
                        } else {
                            var metawork *FrontInterface_ModuleMeta
                            var luaTxt string
                            metawork,luaTxt = self.FP.loadParserToLuaCode(_env, importModuleInfo, &Types_ParserSrc__LnsPath{lnsPath, orgMod, nil}, lnsPath, orgMod, baseDir)
                            self.moduleMgr.FP.AddMeta(_env, orgMod, metawork)
                            self.convertedMap.Set(orgMod,luaTxt)
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
                                self.moduleMgr.FP.AddMeta(_env, orgMod, meta)
                                self.convertedMap.Set(orgMod,luaTxt)
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
// 1291: decl @lune.@base.@front.Front.dumpTokenize
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
// 1303: decl @lune.@base.@front.Front.dumpAst
func (self *front_Front) DumpAst(_env *LnsEnv, scriptPath string) {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, scriptPath)
    Depend_profile(_env, self.option.ValidProf, conv2Form3_1409(func(_env *LnsEnv) {
        var ast *AstInfo_ASTInfo
        ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{scriptPath, mod, nil}, baseDir, mod, front_getModuleId_20_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, DumpNode_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt(_env, "", 0)))
    }), scriptPath + ".profi")
}
// 1323: decl @lune.@base.@front.Front.format
func (self *front_Front) Format(_env *LnsEnv, scriptPath string) {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, scriptPath)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{scriptPath, mod, nil}, baseDir, mod, front_getModuleId_20_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, Formatter_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), Lns_io_stdout), Formatter_Opt2Stem(NewFormatter_Opt(_env, ast.FP.Get_node(_env))))
}
// 1339: decl @lune.@base.@front.Front.checkDiag
func (self *front_Front) CheckDiag(_env *LnsEnv, scriptPath string) {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, scriptPath)
    Util_setErrorCode(_env, 0)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{scriptPath, mod, nil}, baseDir, mod, front_getModuleId_20_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Diag, nil)
}
// 1349: decl @lune.@base.@front.Front.complete
func (self *front_Front) Complete(_env *LnsEnv, scriptPath string) {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, scriptPath)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{scriptPath, mod, nil}, baseDir, mod, front_getModuleId_20_(_env, scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Complete, self.option.AnalyzePos)
}
// 1357: decl @lune.@base.@front.Front.inquire
func (self *front_Front) Inquire(_env *LnsEnv, scriptPath string) {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, scriptPath)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{scriptPath, mod, nil}, baseDir, mod, front_getModuleId_20_(_env, scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Inquire, self.option.AnalyzePos)
}
// 1366: decl @lune.@base.@front.Front.createGlue
func (self *front_Front) CreateGlue(_env *LnsEnv, scriptPath string) {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, scriptPath)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{scriptPath, mod, nil}, baseDir, mod, front_getModuleId_20_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    var filter *Nodes_Filter
    filter = GlueFilter_createFilter(_env, self.option.OutputDir)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, filter, 0)
}
// 1651: decl @lune.@base.@front.Front.convertToLua
func (self *front_Front) convertToLua(_env *LnsEnv, scriptPath string,convMode LnsInt,streamLua Lns_oStream,streamMeta Lns_oStream) LnsAny {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, scriptPath)
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_20_(_env, scriptPath, mod, nil, nil)
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{scriptPath, mod, nil}, baseDir, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(_env, ast, scriptPath, convMode)
    streamLua.Write(_env, luaTxt)
    streamMeta.Write(_env, metaTxt)
    if _switch0 := self.option.ConvTo; _switch0 == Types_Lang__Go {
        var conv *Nodes_Filter
        conv = ConvGo_createFilter(_env, self.option.Testing, "stdout", streamLua, ast, self.FP.createGoOption(_env, scriptPath))
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast.FP.Get_node(_env))))
    }
    return ast
}
// 1750: decl @lune.@base.@front.Front.saveToGo
func (self *front_Front) saveToGo(_env *LnsEnv, scriptPath string,astResult LnsAny) LnsAny {
    return Newfront_GoConverter(_env, scriptPath, astResult, self.option, self.FP.createGoOption(_env, scriptPath))
}
// 1759: decl @lune.@base.@front.Front.saveToC
func (self *front_Front) SaveToC(_env *LnsEnv, scriptPath string,ast *AstInfo_ASTInfo) {
    var cPath string
    cPath = front_convExp4_1208(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".c")))
    var file Lns_luaStream
    
    {
        _file := front_convExp4_1223(Lns_2DDD(Lns_io_open(cPath, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var hPath string
    hPath = front_convExp4_1237(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".h")))
    var hFile Lns_luaStream
    
    {
        _hFile := front_convExp4_1253(Lns_2DDD(Lns_io_open(hPath, "w")))
        if _hFile == nil{
            return 
        } else {
            hFile = _hFile.(Lns_luaStream)
        }
    }
    file.Close(_env)
    hFile.Close(_env)
}
// 1775: decl @lune.@base.@front.Front.outputBuiltin
func (self *front_Front) OutputBuiltin(_env *LnsEnv, scriptPath string) {
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, "lns_builtin")
    var ast *AstInfo_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsCode{"", mod, nil}, baseDir, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    self.FP.SaveToC(_env, scriptPath, ast)
}
// 1797: decl @lune.@base.@front.Front.saveToLua
func (self *front_Front) saveToLua(_env *LnsEnv, updateInfo *front_UpdateInfo) LnsAny {
    var scriptPath string
    scriptPath = updateInfo.FP.Get_scriptPath(_env)
    var dependsPath LnsAny
    dependsPath = updateInfo.FP.Get_dependsPath(_env)
    var moduleId *FrontInterface_ModuleId
    moduleId = updateInfo.FP.Get_moduleId(_env)
    var uptodate LnsAny
    uptodate = updateInfo.FP.Get_uptodate(_env)
    var mod string
    var baseDir LnsAny
    mod,baseDir = self.FP.scriptPath2Module(_env, scriptPath)
    var luaPath string
    luaPath = front_convExp4_1377(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".lua")))
    var metaPath string
    metaPath = front_convExp4_1390(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".meta")))
    if Lns_isCondTrue( self.option.OutputDir){
        var filename string
        filename = front_convExp4_1407(Lns_2DDD(_env.GetVM().String_gsub(mod,"%.", "/")))
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
        var luaConv *front_LuaConverter
        luaConv = Newfront_LuaConverter(_env, luaPath, metaPath, dependsPath, result, convMode, scriptPath, self.option.ByteCompile, self.option.StripDebugInfo, self.option)
        var goConv LnsAny
        goConv = nil
        if _switch0 := self.option.ConvTo; _switch0 == Types_Lang__Go {
            goConv = self.FP.saveToGo(_env, scriptPath, result)
        }
        return front_ConverterFunc_6_(func(_env *LnsEnv) {
            luaConv.FP.SaveLua(_env)
            _env.NilAccFin(_env.NilAccPush(goConv) && 
            Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*front_GoConverter).FP.saveGo(_env)})/* 1860:13 */)
        })
    case *front_ModuleUptodate__Uptodate:
    metaInfo := _matchExp0.Val1
        Util_errorLog(_env, "uptodate -- " + scriptPath)
        var dependsStream LnsAny
        dependsStream = self.option.FP.OpenDepend(_env, dependsPath)
        front_outputDependInfo_24_(_env, dependsStream, metaInfo, mod)
        front_closeStreams_22_(_env, nil, nil, dependsStream, metaPath, false)
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
        front_outputDependInfo_24_(_env, dependsStream, metaInfo, mod)
        front_closeStreams_22_(_env, nil, metaMemStream, dependsStream, metaPath, self.option.Mode == Option_ModeKind__SaveMeta)
        return nil
    }
// insert a dummy
    return nil
}
// 1890: decl @lune.@base.@front.Front.outputBootC
func (self *front_Front) outputBootC(_env *LnsEnv, scriptPath string) {
}
// 1928: decl @lune.@base.@front.Front.build
func (self *front_Front) Build(_env *LnsEnv, buildMode LnsAny,astCallback LnsAny) {
    var createUpdateInfo func(_env *LnsEnv, scriptPath string,dependsPath LnsAny) *front_UpdateInfo
    createUpdateInfo = func(_env *LnsEnv, scriptPath string,dependsPath LnsAny) *front_UpdateInfo {
        var mod string
        mod = front_convExp5_103(Lns_2DDD(self.FP.scriptPath2Module(_env, scriptPath)))
        var moduleId *FrontInterface_ModuleId
        var uptodate LnsAny
        moduleId,uptodate = self.FP.getModuleIdAndCheckUptodate(_env, scriptPath, mod)
        return Newfront_UpdateInfo(_env, scriptPath, dependsPath, moduleId, uptodate)
    }
    var process func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo) LnsAny
    process = func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo) LnsAny {
        var mod string
        var baseDir LnsAny
        mod,baseDir = self.FP.scriptPath2Module(_env, updateInfo.FP.Get_scriptPath(_env))
        switch _matchExp0 := buildMode.(type) {
        case *front_BuildMode__Save:
            return self.FP.saveToLua(_env, updateInfo)
        case *front_BuildMode__Output:
        streamLua := _matchExp0.Val1
        streamMeta := _matchExp0.Val2
            self.FP.convertToLua(_env, updateInfo.FP.Get_scriptPath(_env), ConvLua_ConvMode__ConvMeta, streamLua, streamMeta)
        case *front_BuildMode__CreateAst:
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(self.mod2astCreate.Get(mod))) &&
                _env.SetStackVal( Lns_op_not(self.moduleMgr.FP.GetAst(_env, mod))) ).(bool)){
                var result LnsAny
                result = self.FP.createAstSub(_env, NewFrontInterface_ImportModuleInfo(_env), &Types_ParserSrc__LnsPath{updateInfo.FP.Get_scriptPath(_env), mod, nil}, baseDir, mod, updateInfo.FP.Get_moduleId(_env), nil, TransUnit_AnalyzeMode__Compile, nil)
                return front_ConverterFunc_6_(func(_env *LnsEnv) {
                    self.FP.applyAstResult(_env, result)
                })
            }
        }
        return nil
    }
    Depend_profile(_env, self.option.ValidProf, conv2Form5_536(func(_env *LnsEnv) {
        if self.option.ScriptPath == "@-"{
            for _, _path := range( self.option.BatchList.Items ) {
                path := _path.(string)
                self.targetSet.Add(Lns_car(self.FP.scriptPath2Module(_env, path)).(string))
            }
            var postProcessMap *LnsMap
            postProcessMap = NewLnsMap( map[LnsAny]LnsAny{})
            for _index, _path := range( self.option.BatchList.Items ) {
                index := _index + 1
                path := _path.(string)
                var updateInfo *front_UpdateInfo
                updateInfo = createUpdateInfo(_env, path, Lns_car(_env.GetVM().String_gsub(path,".lns$", ".d")).(string))
                Lns_print([]LnsAny{_env.GetVM().String_format("%s: start...", []LnsAny{updateInfo.FP.Get_scriptPath(_env)})})
                {
                    __exp := process(_env, false, updateInfo)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(front_ConverterFunc_6_)
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
                    postProcess := __forsortCollection0.Items[ _index ].(front_ConverterFunc_6_)
                    index := _index.(LnsInt)
                    var prev LnsReal
                    prev = _env.GetVM().OS_clock()
                    var path string
                    path = self.option.BatchList.GetAt(index).(string)
                    Lns_print([]LnsAny{_env.GetVM().String_format("%s: waiting...", []LnsAny{path})})
                    postProcess(_env)
                    Lns_print([]LnsAny{_env.GetVM().String_format("%s: done %g msec", []LnsAny{path, (_env.GetVM().OS_clock() - prev) * LnsReal(1000)})})
                }
            }
        } else { 
            {
                _postProcess := process(_env, true, createUpdateInfo(_env, self.option.ScriptPath, nil))
                if !Lns_IsNil( _postProcess ) {
                    postProcess := _postProcess.(front_ConverterFunc_6_)
                    postProcess(_env)
                }
            }
        }
        if astCallback != nil{
            astCallback_52 := astCallback.(Front_AstCallback)
            for _, _mod := range( self.moduleMgr.FP.GetModList(_env).Items ) {
                mod := _mod.(string)
                {
                    __exp := self.moduleMgr.FP.GetAst(_env, mod)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*AstInfo_ASTInfo)
                        astCallback_52(_env, _exp)
                    }
                }
            }
        }
    }), self.option.ScriptPath + ".profi")
}
// 2016: decl @lune.@base.@front.Front.exec
func (self *front_Front) Exec(_env *LnsEnv) {
    __func__ := "@lune.@base.@front.Front.exec"
    Log_log(_env, Log_Level__Trace, __func__, 2018, Log_CreateMessage(func(_env *LnsEnv) string {
        return Option_ModeKind_getTxt( self.option.Mode)
    }))
    
    if _switch0 := self.option.Mode; _switch0 == Option_ModeKind__Token {
        self.FP.DumpTokenize(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Ast {
        self.FP.DumpAst(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Format {
        self.FP.Format(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Diag {
        self.FP.CheckDiag(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Complete {
        self.FP.Complete(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Inquire {
        self.FP.Inquire(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Glue {
        self.FP.CreateGlue(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Lua || _switch0 == Option_ModeKind__LuaMeta {
        var convMode LnsInt
        if self.option.Mode == Option_ModeKind__Lua{
            convMode = ConvLua_ConvMode__Convert
        } else { 
            convMode = ConvLua_ConvMode__ConvMeta
        }
        self.FP.convertToLua(_env, self.option.ScriptPath, convMode, Lns_io_stdout, Lns_io_stdout)
    } else if _switch0 == Option_ModeKind__Save || _switch0 == Option_ModeKind__SaveMeta {
        self.FP.Build(_env, front_BuildMode__Save_Obj, nil)
    } else if _switch0 == Option_ModeKind__BuildAst {
        self.FP.Build(_env, front_BuildMode__CreateAst_Obj, Front_AstCallback(Front_exec___anonymous_1_))
    } else if _switch0 == Option_ModeKind__Shebang {
        Lns_LockEnvSync( _env, func () {
            {
                _modObj := front_convExp5_866(Lns_2DDD(self.FP.LoadModule(_env, Lns_car(self.FP.scriptPath2Module(_env, self.option.ScriptPath)).(string))))
                if !Lns_IsNil( _modObj ) {
                    modObj := _modObj
                    var code LnsInt
                    code = Depend_runMain(_env, modObj.(*Lns_luaValue).GetAt("__main"), self.option.ShebangArgList)
                    _env.GetVM().OS_exit(code)
                }
            }
        })
    } else if _switch0 == Option_ModeKind__Exec {
        DependLuaOnLns_runLuaOnLns(_env, "", nil, false)
        _ = front_convExp5_905(Lns_2DDD(self.FP.LoadModule(_env, Lns_car(self.FP.scriptPath2Module(_env, self.option.ScriptPath)).(string))))
        if self.option.Testing{
            var code string
            code = "local Testing = require( \"lune.base.Testing\" )\nreturn function( path )\n  Testing.run( path );\n  Testing.outputAllResult( io.stdout );\nend\n"
            Lns_LockEnvSync( _env, func () {
                var loaded LnsAny
                var mess LnsAny
                loaded,mess = _env.GetVM().Load(code, nil)
                if loaded != nil{
                    loaded_92 := loaded.(*Lns_luaValue)
                    {
                        _mod := front_convExp5_961(Lns_2DDD(_env.GetVM().RunLoadedfunc(loaded_92,Lns_2DDD([]LnsAny{}))[0]))
                        if !Lns_IsNil( _mod ) {
                            mod := _mod
                            _env.GetVM().RunLoadedfunc((mod.(*Lns_luaValue)),Lns_2DDD([]LnsAny{Lns_car(self.FP.scriptPath2Module(_env, self.option.ScriptPath)).(string)}))
                        }
                    }
                } else {
                    Lns_print([]LnsAny{mess})
                }
            })
        }
    } else if _switch0 == Option_ModeKind__BootC {
        self.FP.outputBootC(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__Builtin {
        self.FP.OutputBuiltin(_env, self.option.ScriptPath)
    } else if _switch0 == Option_ModeKind__MkMain {
        var mod string
        mod = front_convExp5_1024(Lns_2DDD(self.FP.scriptPath2Module(_env, self.option.ScriptPath)))
        {
            _mess := ConvGo_outputGoMain(_env, self.FP.getGoAppName(_env), mod, self.option.Testing, self.option.OutputPath, self.option.FP.Get_runtimeOpt(_env))
            if !Lns_IsNil( _mess ) {
                mess := _mess.(string)
                Util_errorLog(_env, mess)
            }
        }
    } else if _switch0 == Option_ModeKind__Indexer {
        self.FP.Build(_env, front_BuildMode__CreateAst_Obj, Front_AstCallback(Front_exec___anonymous_2_))
    } else {
        Lns_print([]LnsAny{"illegal mode"})
    }
}
// 610: decl @lune.@base.@front.MetaForBuildId.createModuleId
func (self *front_MetaForBuildId) CreateModuleId(_env *LnsEnv) *FrontInterface_ModuleId {
    return FrontInterface_ModuleId_createIdFromTxt(_env, self.G__buildId)
}
// 615: decl @lune.@base.@front.MetaForBuildId.LoadFromMeta
func front_MetaForBuildId_LoadFromMeta_5_(_env *LnsEnv, metaPath string)(LnsAny, LnsAny) {
    {
        _fileObj := front_convExp1_1390(Lns_2DDD(Lns_io_open(metaPath, nil)))
        if !Lns_IsNil( _fileObj ) {
            fileObj := _fileObj.(Lns_luaStream)
            var luaCode LnsAny
            luaCode = fileObj.Read(_env, "*a")
            fileObj.Close(_env)
            if luaCode != nil{
                luaCode_125 := luaCode.(string)
                var meta LnsAny
                    meta = front_convExp1_1377(Lns_2DDD(front_MetaForBuildId__fromStem_4_(_env, _env.GetVM().ExpandLuavalMap(Front_loadFromLuaTxt(_env, luaCode_125)),nil)))
                return meta, luaCode_125
            }
        }
    }
    return nil, nil
}
// 1538: DeclConstr
func (self *front_LuaConverter) Initfront_LuaConverter(_env *LnsEnv, luaPath string,metaPath string,dependsPath LnsAny,astResult LnsAny,convMode LnsInt,path string,byteCompile bool,stripDebugInfo bool,option *Option_Option) {
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
    self.converterFunc = front_ConverterFunc_6_(func(_env *LnsEnv) {
        var stream *Util_memStream
        stream = self.streamMem
        var metaStream *Util_memStream
        metaStream = self.metaStreamMem
        var outStream *Util_memStream
        var oMetaStream *Util_memStream
        outStream,oMetaStream = stream, metaStream
        var ast *AstInfo_ASTInfo
        ast = front_getAstFromResult_11_(_env, self.astResult)
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
        filterInfo = front_ast2LuaMain_12_(_env, ast, path, outStream.FP, oMetaStream.FP, convMode, false, option)
        self.filterInfo = filterInfo
        filterInfo.FP.OutputLua(_env, Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode))
    })
    self.FP.Start(_env, 1, _env.GetVM().String_format("convlua -- %s", []LnsAny{path}))
}
// 1589: decl @lune.@base.@front.LuaConverter.runMain
func (self *front_LuaConverter) RunMain(_env *LnsEnv) {
    self.converterFunc(_env)
}
// 1593: decl @lune.@base.@front.LuaConverter.saveLua
func (self *front_LuaConverter) SaveLua(_env *LnsEnv) {
    LnsJoin(_env, self.FP)
    var ast *AstInfo_ASTInfo
    ast = front_getAstFromResult_11_(_env, self.astResult)
    _env.NilAccFin(_env.NilAccPush(self.filterInfo) && 
    Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*ConvLua_FilterInfo).FP.OutputMeta(_env, Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode))})/* 1597:7 */)
    if self.byteCompile{
        self.streamMem.FP.Write(_env, front_byteCompileFromLuaTxt_16_(_env, self.byteStream.FP.Get_txt(_env), self.stripDebugInfo))
        self.metaStreamMem.FP.Write(_env, front_byteCompileFromLuaTxt_16_(_env, self.byteMetaStream.FP.Get_txt(_env), true))
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
        _streamDst := front_convExp4_838(Lns_2DDD(Lns_io_open(self.luaPath, "w")))
        if _streamDst == nil{
            panic(_env.GetVM().String_format("write open error -- %s", []LnsAny{self.luaPath}))
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
        dependsStreamDst_92 := dependsStreamDst.(Lns_oStream)
        dependsStreamDst_92.Write(_env, Lns_unwrap( dependTxt).(string))
    }
    front_closeStreams_22_(_env, streamDst, metaMemStream, dependsStreamDst, self.metaPath, self.option.Mode == Option_ModeKind__SaveMeta)
}
// 1684: DeclConstr
func (self *front_GoConverter) Initfront_GoConverter(_env *LnsEnv, scriptPath string,astResult LnsAny,option *Option_Option,goOpt *ConvGo_Option) {
    self.InitRunner_Runner(_env)
    self.validFlag = true
    var path string
    path = front_convExp0_2089(Lns_2DDD(_env.GetVM().String_gsub(scriptPath,"%.lns$", ".go")))
    {
        _dir := option.OutputDir
        if !Lns_IsNil( _dir ) {
            dir := _dir.(string)
            path = _env.GetVM().String_format("%s/%s", []LnsAny{dir, path})
        }
    }
    self.path = path
    self.memStream = NewUtil_memStream(_env)
    self.converter = front_ConverterFunc_6_(func(_env *LnsEnv) {
        var ast *AstInfo_ASTInfo
        ast = front_getAstFromResult_11_(_env, astResult)
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
        conv = ConvGo_createFilter(_env, option.Testing, path, self.memStream.FP, ast, goOpt)
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast.FP.Get_node(_env))))
    })
    self.FP.Start(_env, 1, _env.GetVM().String_format("convgo -- %s", []LnsAny{scriptPath}))
}
// 1724: decl @lune.@base.@front.GoConverter.runMain
func (self *front_GoConverter) RunMain(_env *LnsEnv) {
    self.converter(_env)
}
// 1728: decl @lune.@base.@front.GoConverter.saveGo
func (self *front_GoConverter) saveGo(_env *LnsEnv) {
    LnsJoin(_env, self.FP)
    if Lns_op_not(self.validFlag){
        return 
    }
    var file Lns_luaStream
    
    {
        _file := front_convExp4_1154(Lns_2DDD(Lns_io_open(self.path, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    file.Write(_env, self.memStream.FP.Get_txt(_env))
    file.Close(_env)
}
