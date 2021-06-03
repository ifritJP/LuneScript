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
// decl alge -- CreateAstResult
type front_CreateAstResult = LnsAny
type front_CreateAstResult__Ast struct{
Val1 *TransUnit_ASTInfo
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
type Front_AstCallback func (_env *LnsEnv, arg1 *TransUnit_ASTInfo)
type front_ConverterFuncSub_1145_ func (_env *LnsEnv)
type front_ConverterFunc_1147_ func (_env *LnsEnv) LnsAny
// for 1110: ExpCast
func conv2Form5722( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1796: ExpCast
func conv2Form9167( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 1913: ExpCast
func conv2Form9527( src func (_env *LnsEnv)) LnsForm {
    return func (_env *LnsEnv,  argList []LnsAny) []LnsAny {
        src(_env)
        return []LnsAny{}
    }
}
// for 447
func front_convExp1811(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 518
func front_convExp2210(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 537
func front_convExp2304(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 796
func front_convExp3599(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 805
func front_convExp3628(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1293
func front_convExp6645(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1302
func front_convExp6716(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1893
func front_convExp9469(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1912
func front_convExp9545(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 536
func front_convExp2283(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 563
func front_convExp2437(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 564
func front_convExp2456(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 604
func front_convExp2644(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 789
func front_convExp3531(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1476
func front_convExp7526(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 457
func front_convExp1848(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, LnsAny, LnsAny) {
    return _env, Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 322
func front_convExp1223(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 380
func front_convExp1454(arg1 []LnsAny) *TransUnit_ASTInfo {
    return Lns_getFromMulti( arg1, 0 ).(*TransUnit_ASTInfo)
}
// for 392
func front_convExp1529(arg1 []LnsAny) (*TransUnit_ASTInfo, *FrontInterface_ModuleInfo) {
    return Lns_getFromMulti( arg1, 0 ).(*TransUnit_ASTInfo), Lns_getFromMulti( arg1, 1 ).(*FrontInterface_ModuleInfo)
}
// for 461
func front_convExp1873(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 497
func front_convExp2081(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 522
func front_convExp2200(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 684
func front_convExp2934(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 711
func front_convExp3087(arg1 []LnsAny) (LnsAny, string, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string), Lns_getFromMulti( arg1, 2 ).(string)
}
// for 755
func front_convExp3333(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 777
func front_convExp3409(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 834
func front_convExp3798(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 889
func front_convExp4194(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 932
func front_convExp4523(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 941
func front_convExp4581(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 963
func front_convExp4695(arg1 []LnsAny) (*FrontInterface_ModuleMeta, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 )
}
// for 1003
func front_convExp4941(arg1 []LnsAny) (*TransUnit_ASTInfo, *FrontInterface_ModuleInfo) {
    return Lns_getFromMulti( arg1, 0 ).(*TransUnit_ASTInfo), Lns_getFromMulti( arg1, 1 ).(*FrontInterface_ModuleInfo)
}
// for 1016
func front_convExp5035(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1041
func front_convExp5133(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1067
func front_convExp5438(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1076
func front_convExp5512(arg1 []LnsAny) (*FrontInterface_ModuleMeta, string) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleMeta), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1183
func front_convExp6110(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1242
func front_convExp6358(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1243
func front_convExp6371(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1298
func front_convExp6669(arg1 []LnsAny) (bool, string) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1419
func front_convExp7241(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1467
func front_convExp7469(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1495
func front_convExp7646(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1521
func front_convExp7769(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1546
func front_convExp7910(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1583
func front_convExp8031(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1585
func front_convExp8046(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1589
func front_convExp8060(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1590
func front_convExp8076(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1630
func front_convExp8220(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1631
func front_convExp8233(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1633
func front_convExp8250(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1769
func front_convExp8787(arg1 []LnsAny) (*FrontInterface_ModuleId, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*FrontInterface_ModuleId), Lns_getFromMulti( arg1, 1 )
}
// for 1900
func front_convExp9491(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1910
func front_convExp9510(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
func front__anonymous_1008_(_env *LnsEnv, ver LnsInt) {
    LuaVer_setCurVer(_env, ver)
}
// 79: decl @lune.@base.@front.createModuleInfo
func front_createModuleInfo_1034_(_env *LnsEnv, ast *TransUnit_ASTInfo,mod string,moduleId *FrontInterface_ModuleId) *FrontInterface_ModuleInfo {
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


// 330: decl @lune.@base.@front.createPaser
func front_createPaser_1345_(_env *LnsEnv, path string,mod string) *Parser_Parser {
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

// 418: decl @lune.@base.@front.ast2LuaMain
func front_ast2LuaMain_1405_(_env *LnsEnv, ast *TransUnit_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,option *Option_Option) *ConvLua_FilterInfo {
    var exportInfo *FrontInterface_ExportInfo
    exportInfo = ast.FP.Get_exportInfo(_env)
    var conv *ConvLua_FilterInfo
    conv = ConvLua_createFilter(_env, streamName, stream, metaStream, convMode, inMacro, exportInfo.FP.Get_moduleTypeInfo(_env), exportInfo.FP.Get_processInfo(_env), exportInfo.FP.Get_provideInfo(_env).FP.Get_symbolKind(_env), ast.FP.Get_builtinFunc(_env), option.UseLuneModule, option.TargetLuaVer, option.Testing, option.UseIpairs)
    return conv
}

// 433: decl @lune.@base.@front.ast2Lua
func front_ast2Lua_1411_(_env *LnsEnv, ast *TransUnit_ASTInfo,streamName string,stream Lns_oStream,metaStream Lns_oStream,convMode LnsInt,inMacro bool,option *Option_Option) {
    var conv *ConvLua_FilterInfo
    conv = front_ast2LuaMain_1405_(_env, ast, streamName, stream, metaStream, convMode, inMacro, option)
    conv.FP.OutputLuaAndMeta(_env, Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode))
}

// 442: decl @lune.@base.@front.loadFromChunk
func front_loadFromChunk_1416_(_env *LnsEnv, chunk LnsAny,err LnsAny) LnsAny {
    if err != nil{
        err_320 := err.(string)
        Util_errorLog(_env, err_320)
    }
    if chunk != nil{
        chunk_322 := chunk.(*Lns_luaValue)
        {
            _work := front_convExp1811(Lns_2DDD(_env.CommonLuaVM.RunLoadedfunc(chunk_322,Lns_2DDD([]LnsAny{}))[0]))
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

// 455: decl @lune.@base.@front.loadFromLuaTxt
func Front_loadFromLuaTxt(_env *LnsEnv, txt string) LnsAny {
    return front_loadFromChunk_1416_(front_convExp1848(_env, Lns_2DDD(_env.CommonLuaVM.Load(txt, nil))))
}

// 460: decl @lune.@base.@front.byteCompileFromLuaTxt
func front_byteCompileFromLuaTxt_1428_(_env *LnsEnv, txt string,stripDebugInfo bool) string {
    var chunk LnsAny
    var err LnsAny
    chunk,err = _env.CommonLuaVM.Load(txt, nil)
    if chunk != nil{
        chunk_337 := chunk.(*Lns_luaValue)
        return _env.CommonLuaVM.String_dump(chunk_337, stripDebugInfo)
    }
    panic(Lns_unwrapDefault( err, "load error").(string))
// insert a dummy
    return ""
}

// 531: decl @lune.@base.@front.getMetaInfo
func front_getMetaInfo_1527_(_env *LnsEnv, lnsPath string,mod string,outdir LnsAny)(LnsAny, string, string) {
    var moduleMetaPath string
    moduleMetaPath = lnsPath
    if outdir != nil{
        outdir_396 := outdir.(string)
        moduleMetaPath = _env.LuaVM.String_format("%s/%s", []LnsAny{outdir_396, Lns_car(_env.LuaVM.String_gsub(mod,"%.", "/")).(string)})
        
    }
    moduleMetaPath = front_convExp2283(Lns_2DDD(_env.LuaVM.String_gsub(moduleMetaPath,"%.lns$", ".meta")))
    
    {
        _meta, _metaCode := front_MetaForBuildId_LoadFromMeta_1513_(_env, moduleMetaPath)
        if !Lns_IsNil( _meta ) && !Lns_IsNil( _metaCode ) {
            meta := _meta.(*front_MetaForBuildId)
            metaCode := _metaCode.(string)
            return meta, moduleMetaPath, metaCode
        }
    }
    return nil, moduleMetaPath, ""
}


// 596: decl @lune.@base.@front.getModuleId
func front_getModuleId_1579_(_env *LnsEnv, lnsPath string,mod string,outdir LnsAny,metaInfo LnsAny) *FrontInterface_ModuleId {
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
        metaInfo = front_convExp2644(Lns_2DDD(front_getMetaInfo_1527_(_env, lnsPath, mod, outdir)))
        
    }
    if metaInfo != nil{
        metaInfo_441 := metaInfo.(*front_MetaForBuildId)
        var buildId *FrontInterface_ModuleId
        buildId = metaInfo_441.FP.CreateModuleId(_env)
        buildCount = buildId.FP.Get_buildCount(_env)
        
    }
    return FrontInterface_ModuleId_createId(_env, fileTime, buildCount)
}

func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1612_(_env *LnsEnv) string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1621_(_env *LnsEnv) string {
    return "NeedUpdate"
}
func Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1630_(_env *LnsEnv) string {
    return "NeedUpdate"
}


func Front_getModuleIdAndCheckUptodate___anonymous_1667_(_env *LnsEnv) string {
    return "not found meta"
}
















// 1182: decl @lune.@base.@front.closeStreams.txt2ModuleId
func closeStreams__txt2ModuleId_2086_(_env *LnsEnv, txt string) *FrontInterface_ModuleId {
    var buildIdTxt string
    buildIdTxt = front_convExp6110(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(txt,"^_moduleObj.__buildId = ", "")).(string),"\"", "")))
    return FrontInterface_ModuleId_createIdFromTxt(_env, buildIdTxt)
}




// 1179: decl @lune.@base.@front.closeStreams
func front_closeStreams_2084_(_env *LnsEnv, stream LnsAny,metaStream LnsAny,dependStream LnsAny,metaPath string,saveMetaFlag bool) {
    var checkDiff func(_env *LnsEnv, oldStream *Parser_TxtStream,newStream *Parser_TxtStream)(bool, string)
    checkDiff = func(_env *LnsEnv, oldStream *Parser_TxtStream,newStream *Parser_TxtStream)(bool, string) {
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
                oldLine_814 := oldLine.(string)
                if len(oldBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(oldLine_814,"^_moduleObj.__buildId", nil, nil))){
                        oldBuildIdLine = oldLine_814
                        
                    }
                }
            }
            
            if newLine != nil{
                newLine_818 := newLine.(string)
                if len(newBuildIdLine) == 0{
                    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(newLine_818,"^_moduleObj.__buildId", nil, nil))){
                        newBuildIdLine = newLine_818
                        
                    }
                }
            }
            
            if newLine != oldLine{
                var cont bool
                cont = false
                if newLine != nil && oldLine != nil{
                    newLine_824 := newLine.(string)
                    oldLine_825 := oldLine.(string)
                    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(oldLine_825,"^_moduleObj.__buildId", nil, nil))){
                        if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(newLine_824,"^_moduleObj.__buildId", nil, nil))){
                            tailBeginPos = newStream.FP.Get_lineNo(_env)
                            
                            cont = true
                            
                        }
                    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_car(_env.LuaVM.String_find(oldLine_825,"^__dependModuleMap.*buildId =", nil, nil))) &&
                        _env.SetStackVal( Lns_car(_env.LuaVM.String_find(newLine_824,"^__dependModuleMap.*buildId =", nil, nil))) )){
                        var oldSub string
                        oldSub = front_convExp6358(Lns_2DDD(_env.LuaVM.String_gsub(oldLine_825,"buildId =.*", "")))
                        var newSub string
                        newSub = front_convExp6371(Lns_2DDD(_env.LuaVM.String_gsub(newLine_824,"buildId =.*", "")))
                        if oldSub == newSub{
                            cont = true
                            
                        }
                    }
                }
                if Lns_op_not(cont){
                    Log_log(_env, Log_Level__Debug, __func__, 1250, Log_CreateMessage(func(_env *LnsEnv) string {
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
                    oldBuildId = closeStreams__txt2ModuleId_2086_(_env, oldBuildIdLine)
                    var newBuildId *FrontInterface_ModuleId
                    newBuildId = closeStreams__txt2ModuleId_2086_(_env, newBuildIdLine)
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
    if stream != nil{
        stream_845 := stream.(Lns_oStream)
        stream_845.Close(_env)
    }
    if dependStream != nil{
        dependStream_847 := dependStream.(Lns_oStream)
        dependStream_847.Close(_env)
    }
    if metaStream != nil{
        metaStream_849 := metaStream.(*Util_memStream)
        if saveMetaFlag{
            var newMetaTxt string
            newMetaTxt = metaStream_849.FP.Get_txt(_env)
            var oldMetaTxt string
            oldMetaTxt = ""
            {
                _oldFileObj := front_convExp6645(Lns_2DDD(Lns_io_open(metaPath, nil)))
                if !Lns_IsNil( _oldFileObj ) {
                    oldFileObj := _oldFileObj.(Lns_luaStream)
                    oldMetaTxt = Lns_unwrapDefault( oldFileObj.Read(_env, "*a"), "").(string)
                    
                    oldFileObj.Close(_env)
                }
            }
            var sameFlag bool
            var txt string
            sameFlag,txt = checkDiff(_env, NewParser_TxtStream(_env, oldMetaTxt), NewParser_TxtStream(_env, newMetaTxt))
            var saveMeta func(_env *LnsEnv, meta string)
            saveMeta = func(_env *LnsEnv, meta string) {
                {
                    _fileObj := front_convExp6716(Lns_2DDD(Lns_io_open(metaPath, "w")))
                    if !Lns_IsNil( _fileObj ) {
                        fileObj := _fileObj.(Lns_luaStream)
                        fileObj.Write(_env, meta)
                        fileObj.Close(_env)
                    } else {
                        Util_err(_env, _env.LuaVM.String_format("failed to open -- %s", []LnsAny{metaPath}))
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


// 1437: decl @lune.@base.@front.outputDependInfo
func front_outputDependInfo_2249_(_env *LnsEnv, stream LnsAny,metaInfo *front_MetaForBuildId,mod string) {
    if stream != nil{
        stream_934 := stream.(Lns_oStream)
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
        dependInfo.FP.Output(_env, stream_934)
    }
}






// 1746: decl @lune.@base.@front.convertLnsCode2LuaCode
func Front_convertLnsCode2LuaCode(_env *LnsEnv, lnsCode string,path string) string {
    var option *Option_Option
    option = NewOption_Option(_env)
    option.ScriptPath = path
    
    option.UseLuneModule = Option_getRuntimeModule(_env)
    
    option.UseIpairs = true
    
    var front *front_Front
    front = Newfront_Front(_env, option, nil)
    return front.FP.ConvertLns2LuaCode(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsCode{lnsCode, path}, NewParser_TxtStream(_env, lnsCode).FP, path)
}




// 1847: decl @lune.@base.@front.build
func Front_build(_env *LnsEnv, option *Option_Option,astCallback Front_AstCallback) {
    var front *front_Front
    front = Newfront_Front(_env, option, nil)
    front.FP.Build(_env, front_BuildMode__CreateAst_Obj, astCallback)
}


// 1941: decl @lune.@base.@front.exec
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

// 1960: decl @lune.@base.@front.setFront
func Front_setFront(_env *LnsEnv, bindModuleList *LnsList) {
    var option *Option_Option
    option = Option_createDefaultOption(_env, NewLnsList([]LnsAny{"dummy.lns"}), nil)
    Newfront_Front(_env, option, bindModuleList)
}

// 1965: decl @lune.@base.@front.__main
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
    Add(_env *LnsEnv, arg1 *TransUnit_ASTInfo, arg2 *FrontInterface_ModuleInfo)
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
// 101: DeclConstr
func (self *front_ModuleMgr) Initfront_ModuleMgr(_env *LnsEnv) {
    self.mod2info = NewUtil_OrderdMap(_env)
    
    self.loadedMetaMap = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 106: decl @lune.@base.@front.ModuleMgr.get
func (self *front_ModuleMgr) Get(_env *LnsEnv, mod string) LnsAny {
    return self.mod2info.FP.Get_map(_env).Get(mod)
}

// 110: decl @lune.@base.@front.ModuleMgr.getAst
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
    switch _exp278 := info.(type) {
    case *front_UptodateInfo__Update:
    ast := _exp278.Val2
        return ast
    case *front_UptodateInfo__Uptodate:
        return nil
    }
// insert a dummy
    return nil
}

// 124: decl @lune.@base.@front.ModuleMgr.getModList
func (self *front_ModuleMgr) GetModList(_env *LnsEnv) *LnsList {
    return self.mod2info.FP.Get_keyList(_env)
}

// 129: decl @lune.@base.@front.ModuleMgr.add
func (self *front_ModuleMgr) Add(_env *LnsEnv, ast *TransUnit_ASTInfo,moduleInfo *FrontInterface_ModuleInfo) {
    var mod string
    mod = moduleInfo.FP.Get_fullName(_env)
    if Lns_isCondTrue( self.loadedMetaMap.Get(mod)){
        if mod == "lune.base.Testing"{
        } else { 
            Util_err(_env, _env.LuaVM.String_format("the meta is already loaded -- %s", []LnsAny{mod}))
        }
    }
    self.mod2info.FP.Add(_env, mod, &front_UptodateInfo__Update{moduleInfo, ast})
    self.loadedMetaMap.Set(mod,NewFrontInterface_ModuleMeta(_env, ast.FP.Get_streamName(_env), &FrontInterface_MetaOrModule__Module{moduleInfo}))
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


// declaration Class -- AstCreater
type front_AstCreaterMtd interface {
    createAst(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 LnsAny, arg4 LnsInt, arg5 LnsAny) *TransUnit_ASTInfo
    getAst(_env *LnsEnv)(*TransUnit_ASTInfo, *FrontInterface_ModuleInfo)
}
type front_AstCreater struct {
    option *Option_Option
    builtinFunc *Builtin_BuiltinFuncType
    mod string
    moduleId *FrontInterface_ModuleId
    ast LnsAny
    moduleInfo LnsAny
    converter front_ConverterFunc_1147_
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
func Newfront_AstCreater(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 *FrontInterface_ModuleId, arg5 LnsAny, arg6 LnsInt, arg7 LnsAny, arg8 *Builtin_BuiltinFuncType, arg9 *Option_Option) *front_AstCreater {
    obj := &front_AstCreater{}
    obj.FP = obj
    obj.Initfront_AstCreater(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
// 175: decl @lune.@base.@front.AstCreater.createAst
func (self *front_AstCreater) createAst(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *TransUnit_ASTInfo {
    var transUnit *TransUnit_TransUnit
    transUnit = NewTransUnit_TransUnit(_env, self.moduleId, importModuleInfo, &NewConvLua_MacroEvalImp(_env, self.builtinFunc).Nodes_MacroEval, analyzeModule, analyzeMode, pos, self.option.TargetLuaVer, self.option.TransCtrlInfo, self.builtinFunc)
    return transUnit.FP.CreateAST(_env, parserSrc, false, self.mod)
}

// 189: DeclConstr
func (self *front_AstCreater) Initfront_AstCreater(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny,builtinFunc *Builtin_BuiltinFuncType,option *Option_Option) {
    self.option = option
    
    self.builtinFunc = builtinFunc
    
    self.mod = mod
    
    self.moduleId = moduleId
    
    self.moduleInfo = nil
    
    self.converter = front_ConverterFunc_1147_(func(_env *LnsEnv) LnsAny {
        var ast *TransUnit_ASTInfo
        ast = self.FP.createAst(_env, importModuleInfo, parserSrc, analyzeModule, analyzeMode, pos)
        self.ast = ast
        
        self.moduleInfo = front_createModuleInfo_1034_(_env, ast, self.mod, self.moduleId)
        
        return nil
    })
    
    self.converter(_env)
}

// 213: decl @lune.@base.@front.AstCreater.getAst
func (self *front_AstCreater) getAst(_env *LnsEnv)(*TransUnit_ASTInfo, *FrontInterface_ModuleInfo) {
    return Lns_unwrap( self.ast).(*TransUnit_ASTInfo), Lns_unwrap( self.moduleInfo).(*FrontInterface_ModuleInfo)
}


// declaration Class -- Front
type front_FrontMtd interface {
    Build(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny)
    CheckDiag(_env *LnsEnv, arg1 string)
    checkUptodateMeta(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny) LnsAny
    Complete(_env *LnsEnv, arg1 string)
    convertFromAst(_env *LnsEnv, arg1 *TransUnit_ASTInfo, arg2 string, arg3 LnsInt)(string, string)
    ConvertLns2LuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 Lns_iStream, arg4 string) string
    convertToLua(_env *LnsEnv, arg1 string, arg2 LnsInt, arg3 Lns_oStream, arg4 Lns_oStream) LnsAny
    createAst(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 *FrontInterface_ModuleId, arg5 LnsAny, arg6 LnsInt, arg7 LnsAny) *TransUnit_ASTInfo
    createAstSub(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 *FrontInterface_ModuleId, arg5 LnsAny, arg6 LnsInt, arg7 LnsAny) LnsAny
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
    LoadFromLnsTxt(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string, arg3 string) LnsAny
    loadLua(_env *LnsEnv, arg1 string) LnsAny
    LoadMeta(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 string) LnsAny
    LoadModule(_env *LnsEnv, arg1 string)(LnsAny, *FrontInterface_ModuleMeta)
    loadParserToLuaCode(_env *LnsEnv, arg1 *FrontInterface_ImportModuleInfo, arg2 LnsAny, arg3 string, arg4 string)(*FrontInterface_ModuleMeta, string)
    outputBootC(_env *LnsEnv, arg1 string)
    OutputBuiltin(_env *LnsEnv, arg1 string)
    processAstCreater(_env *LnsEnv, arg1 *front_AstCreater) *TransUnit_ASTInfo
    SaveToC(_env *LnsEnv, arg1 string, arg2 *TransUnit_ASTInfo)
    saveToGo(_env *LnsEnv, arg1 string, arg2 *TransUnit_ASTInfo) LnsAny
    saveToLua(_env *LnsEnv, arg1 *front_UpdateInfo) LnsAny
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
    mod2astCreate *LnsMap
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
// 249: DeclConstr
func (self *front_Front) Initfront_Front(_env *LnsEnv, option *Option_Option,bindModuleList LnsAny) {
    self.mod2astCreate = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.loadCount = 0
    
    self.targetSet = NewLnsSet([]LnsAny{})
    
    self.bindModuleSet = NewLnsSet([]LnsAny{})
    
    if bindModuleList != nil{
        bindModuleList_165 := bindModuleList.(*LnsList)
        for _, _mod := range( bindModuleList_165.Items ) {
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
        _exp1036 := Depend_getLoadedMod(_env)
        _key1036, _val1036 := _exp1036.Get1stFromMap()
        for _key1036 != nil {
            mod := _key1036.(string)
            modval := _val1036
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
            _key1036, _val1036 = _exp1036.NextFromMap( _key1036 )
        }
    }
    self.preloadedModMap = loadedMap
    
}

// 290: decl @lune.@base.@front.Front.getLoadInfo
func (self *front_Front) getLoadInfo(_env *LnsEnv, mod string) LnsAny {
    if self.option.Testing{
        return self.loadedMapTest.Get(mod)
    }
    return self.loadedMap.Get(mod)
}

// 297: decl @lune.@base.@front.Front.setLoadInfo
func (self *front_Front) setLoadInfo(_env *LnsEnv, mod string,info *front_LoadInfo) {
    if self.option.Testing{
        self.loadedMapTest.Set(mod,info)
    }
    self.loadedMap.Set(mod,info)
}







// 315: decl @lune.@base.@front.Front.error
func (self *front_Front) Error(_env *LnsEnv, message string) {
    Util_errorLog(_env, message)
    Util_printStackTrace(_env)
    _env.LuaVM.OS_exit(1)
}

// 321: decl @lune.@base.@front.Front.loadLua
func (self *front_Front) loadLua(_env *LnsEnv, path string) LnsAny {
    var chunk LnsAny
    var err LnsAny
    chunk,err = _env.CommonLuaVM.Loadfile(path)
    if chunk != nil{
        chunk_223 := chunk.(*Lns_luaValue)
        return Lns_unwrap( Lns_car(_env.CommonLuaVM.RunLoadedfunc(chunk_223,Lns_2DDD([]LnsAny{}))[0]))
    }
    Util_errorLog(_env, Lns_unwrapDefault( err, _env.LuaVM.String_format("load error -- %s.", []LnsAny{path})).(string))
    return nil
}

// 339: decl @lune.@base.@front.Front.scriptPath2Module
func (self *front_Front) scriptPath2Module(_env *LnsEnv, path string) string {
    return Util_scriptPath2ModuleFromProjDir(_env, path, self.option.FP.Get_projDir(_env))
}

// 343: decl @lune.@base.@front.Front.createPaser
func (self *front_Front) createPaser(_env *LnsEnv, scriptPath string) *Parser_Parser {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    return front_createPaser_1345_(_env, scriptPath, mod)
}

// 355: decl @lune.@base.@front.Front.createAstSub
func (self *front_Front) createAstSub(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) LnsAny {
    {
        __exp := self.moduleMgr.FP.Get(_env, mod)
        if !Lns_IsNil( __exp ) {
            _exp := __exp
            switch _exp1436 := _exp.(type) {
            case *front_UptodateInfo__Update:
            ast := _exp1436.Val2
                return &front_CreateAstResult__Ast{ast}
            case *front_UptodateInfo__Uptodate:
                Util_err(_env, _env.LuaVM.String_format("can't load the multiple module -- %s", []LnsAny{mod}))
            }
        }
    }
    {
        _creater := self.mod2astCreate.Get(mod)
        if !Lns_IsNil( _creater ) {
            creater := _creater.(*front_AstCreater)
            var ast *TransUnit_ASTInfo
            ast = front_convExp1454(Lns_2DDD(creater.FP.getAst(_env)))
            return &front_CreateAstResult__Ast{ast}
        }
    }
    var astCreater *front_AstCreater
    astCreater = Newfront_AstCreater(_env, importModuleInfo, parserSrc, mod, moduleId, analyzeModule, analyzeMode, pos, self.builtinFunc, self.option)
    self.mod2astCreate.Set(mod,astCreater)
    return &front_CreateAstResult__Creater{astCreater}
}

// 391: decl @lune.@base.@front.Front.processAstCreater
func (self *front_Front) processAstCreater(_env *LnsEnv, astCreater *front_AstCreater) *TransUnit_ASTInfo {
    var ast *TransUnit_ASTInfo
    var moduleInfo *FrontInterface_ModuleInfo
    ast,moduleInfo = astCreater.FP.getAst(_env)
    self.moduleMgr.FP.Add(_env, ast, moduleInfo)
    return ast
}

// 398: decl @lune.@base.@front.Front.createAst
func (self *front_Front) createAst(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,mod string,moduleId *FrontInterface_ModuleId,analyzeModule LnsAny,analyzeMode LnsInt,pos LnsAny) *TransUnit_ASTInfo {
    switch _exp1615 := self.FP.createAstSub(_env, importModuleInfo, parserSrc, mod, moduleId, analyzeModule, analyzeMode, pos).(type) {
    case *front_CreateAstResult__Ast:
    ast := _exp1615.Val1
        return ast
    case *front_CreateAstResult__Creater:
    astCreater := _exp1615.Val1
        return self.FP.processAstCreater(_env, astCreater)
    }
// insert a dummy
    return nil
}

// 469: decl @lune.@base.@front.Front.convertFromAst
func (self *front_Front) convertFromAst(_env *LnsEnv, ast *TransUnit_ASTInfo,streamName string,mode LnsInt)(string, string) {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var metaStream *Util_memStream
    metaStream = NewUtil_memStream(_env)
    front_ast2Lua_1411_(_env, ast, streamName, stream.FP, metaStream.FP, mode, false, self.option)
    return metaStream.FP.Get_txt(_env), stream.FP.Get_txt(_env)
}

// 483: decl @lune.@base.@front.Front.loadFromLnsTxt
func (self *front_Front) LoadFromLnsTxt(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,name string,txt string) LnsAny {
    var transUnit *TransUnit_TransUnit
    transUnit = NewTransUnit_TransUnit(_env, FrontInterface_ModuleId__tempId, importModuleInfo, &NewConvLua_MacroEvalImp(_env, self.builtinFunc).Nodes_MacroEval, nil, nil, nil, self.option.TargetLuaVer, self.option.TransCtrlInfo, self.builtinFunc)
    var ast *TransUnit_ASTInfo
    ast = transUnit.FP.CreateAST(_env, &Parser_ParserSrc__LnsCode{txt, name}, false, _env.LuaVM.String_format("$load%d", []LnsAny{self.loadCount}))
    self.loadCount = self.loadCount + 1
    
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(_env, ast, name, ConvLua_ConvMode__ConvMeta)
    return Lns_unwrap( Front_loadFromLuaTxt(_env, luaTxt))
}

// 543: decl @lune.@base.@front.Front.searchModuleFile
func (self *front_Front) searchModuleFile(_env *LnsEnv, mod string,suffix string,addPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.searchModuleFile"
    switch _exp2353 := self.gomodMap.FP.ConvLocalModulePath(_env, mod, suffix).(type) {
    case *GoMod_GoModResult__NotGo:
    case *GoMod_GoModResult__NotFound:
        return nil
    case *GoMod_GoModResult__Found:
    info := _exp2353.Val1
        return info.FP.Get_path(_env)
    }
    var lnsSearchPath string
    lnsSearchPath = Lns_package_path
    if addPath != nil{
        addPath_413 := addPath.(string)
        lnsSearchPath = _env.LuaVM.String_format("%s/?%s;%s", []LnsAny{addPath_413, suffix, Lns_package_path})
        
    }
    {
        _projDir := self.option.FP.Get_projDir(_env)
        if !Lns_IsNil( _projDir ) {
            projDir := _projDir.(string)
            lnsSearchPath = _env.LuaVM.String_format("%s;%s", []LnsAny{Util_pathJoin(_env, projDir, _env.LuaVM.String_format("?%s", []LnsAny{suffix})), Lns_package_path})
            
        }
    }
    lnsSearchPath = front_convExp2437(Lns_2DDD(_env.LuaVM.String_gsub(lnsSearchPath,"%.lua$", suffix)))
    
    lnsSearchPath = front_convExp2456(Lns_2DDD(_env.LuaVM.String_gsub(lnsSearchPath,"%.lua;", suffix + ";")))
    
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
                    Log_log(_env, Log_Level__Err, __func__, 576, Log_CreateMessage(func(_env *LnsEnv) string {
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

// 636: decl @lune.@base.@front.Front.getModuleIdAndCheckUptodate
func (self *front_Front) getModuleIdAndCheckUptodate(_env *LnsEnv, lnsPath string,mod string)(*FrontInterface_ModuleId, LnsAny) {
    __func__ := "@lune.@base.@front.Front.getModuleIdAndCheckUptodate"
    var uptodate LnsAny
    uptodate = front_ModuleUptodate__NeedUpdate_Obj
    switch _exp2743 := self.option.TransCtrlInfo.UptodateMode.(type) {
    case *Types_CheckingUptodateMode__ForceAll:
        return FrontInterface_ModuleId__tempId, uptodate
    case *Types_CheckingUptodateMode__Force1:
    forceMod := _exp2743.Val1
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
                    Log_log(_env, Log_Level__Debug, __func__, 672, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1612_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    modMetaPath = _modMetaPath.(string)
                }
            }
            var time LnsReal
            
            {
                _time := Depend_getFileLastModifiedTime(_env, modMetaPath)
                if _time == nil{
                    Log_log(_env, Log_Level__Debug, __func__, 677, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1621_))
                    
                    return front_ModuleUptodate__NeedUpdate_Obj
                } else {
                    time = _time.(LnsReal)
                }
            }
            if time > metaTime{
                var dependMeta *front_MetaForBuildId
                
                {
                    _dependMeta := front_convExp2934(Lns_2DDD(front_MetaForBuildId_LoadFromMeta_1513_(_env, modMetaPath)))
                    if _dependMeta == nil{
                        Log_log(_env, Log_Level__Debug, __func__, 685, Log_CreateMessage(Front_getModuleIdAndCheckUptodate_checkDependUptodate___anonymous_1630_))
                        
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
                    Log_log(_env, Log_Level__Debug, __func__, 695, Log_CreateMessage(func(_env *LnsEnv) string {
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
    metaInfo,metaPath,metaCode = front_getMetaInfo_1527_(_env, lnsPath, mod, self.option.OutputDir)
    if metaInfo != nil{
        metaInfo_496 := metaInfo.(*front_MetaForBuildId)
        if metaInfo_496.G__enableTest == self.option.Testing{
            var buildId *FrontInterface_ModuleId
            buildId = FrontInterface_ModuleId_createIdFromTxt(_env, metaInfo_496.G__buildId)
            if buildId != FrontInterface_ModuleId__tempId{
                var lnsTime LnsAny
                lnsTime = Depend_getFileLastModifiedTime(_env, lnsPath)
                var metaTime LnsAny
                metaTime = Depend_getFileLastModifiedTime(_env, metaPath)
                if lnsTime != nil && metaTime != nil{
                    lnsTime_503 := lnsTime.(LnsReal)
                    metaTime_504 := metaTime.(LnsReal)
                    if lnsTime_503 == buildId.FP.Get_modTime(_env){
                        uptodate = checkDependUptodate(_env, metaTime_504, metaInfo_496, metaCode)
                        
                    }
                }
            }
        } else { 
        }
    } else {
        Log_log(_env, Log_Level__Debug, __func__, 733, Log_CreateMessage(Front_getModuleIdAndCheckUptodate___anonymous_1667_))
        
    }
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_1579_(_env, lnsPath, mod, self.option.OutputDir, metaInfo)
    if moduleId == FrontInterface_ModuleId__tempId{
        Util_err(_env, _env.LuaVM.String_format("not found -- %s", []LnsAny{lnsPath}))
    }
    return moduleId, uptodate
}

// 746: decl @lune.@base.@front.Front.convertLns2LuaCode
func (self *front_Front) ConvertLns2LuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,stream Lns_iStream,streamName string) string {
    var mod string
    mod = self.FP.scriptPath2Module(_env, streamName)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, parserSrc, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    var luaTxt string
    _,luaTxt = self.FP.convertFromAst(_env, ast, streamName, ConvLua_ConvMode__ConvMeta)
    return luaTxt
}

// 767: decl @lune.@base.@front.Front.loadParserToLuaCode
func (self *front_Front) loadParserToLuaCode(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,parserSrc LnsAny,path string,mod string)(*FrontInterface_ModuleMeta, string) {
    __func__ := "@lune.@base.@front.Front.loadParserToLuaCode"
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_1579_(_env, path, mod, nil, nil)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, importModuleInfo, parserSrc, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(_env, ast, path, ConvLua_ConvMode__ConvMeta)
    Log_log(_env, Log_Level__Trace, __func__, 778, Log_CreateMessage(func(_env *LnsEnv) string {
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
                    newpath = front_convExp3531(Lns_2DDD(_env.LuaVM.String_gsub(path,".lns$", suffix)))
                    
                }
            }
            var saveTxt string
            saveTxt = txt
            if byteCompile{
                saveTxt = front_byteCompileFromLuaTxt_1428_(_env, saveTxt, stripDebugInfo)
                
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(front_forceUpdateMeta)) &&
                _env.SetStackVal( checkUpdate) ).(bool)){
                {
                    _fileObj := front_convExp3599(Lns_2DDD(Lns_io_open(newpath, nil)))
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
                _fileObj := front_convExp3628(Lns_2DDD(Lns_io_open(newpath, "w")))
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
    meta = NewFrontInterface_ModuleMeta(_env, path, &FrontInterface_MetaOrModule__Module{front_createModuleInfo_1034_(_env, ast, mod, moduleId)})
    return meta, luaTxt
}

// 828: decl @lune.@base.@front.Front.loadFile
func (self *front_Front) loadFile(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,path string,mod string)(*FrontInterface_ModuleMeta, LnsAny) {
    __func__ := "@lune.@base.@front.Front.loadFile"
    Log_log(_env, Log_Level__Info, __func__, 832, Log_CreateMessage(func(_env *LnsEnv) string {
        __func__ := "@lune.@base.@front.Front.loadFile.<anonymous>"
        return _env.LuaVM.String_format("start %s:%s", []LnsAny{__func__, mod})
    }))
    
    var meta *FrontInterface_ModuleMeta
    var luaTxt string
    meta,luaTxt = self.FP.loadParserToLuaCode(_env, importModuleInfo, &Parser_ParserSrc__LnsPath{path, mod}, path, mod)
    {
        _preLoadInfo := self.preloadedModMap.Get(mod)
        if !Lns_IsNil( _preLoadInfo ) {
            preLoadInfo := _preLoadInfo
            return meta, preLoadInfo
        }
    }
    return meta, Lns_unwrap( Front_loadFromLuaTxt(_env, luaTxt))
}

// 845: decl @lune.@base.@front.Front.searchModule
func (self *front_Front) SearchModule(_env *LnsEnv, mod string) LnsAny {
    return self.FP.searchModuleFile(_env, mod, ".lns", nil)
}

// 849: decl @lune.@base.@front.Front.searchLuaFile
func (self *front_Front) searchLuaFile(_env *LnsEnv, moduleFullName string,addSearchPath LnsAny) LnsAny {
    return self.FP.searchModuleFile(_env, moduleFullName, ".lua", addSearchPath)
}

// 859: decl @lune.@base.@front.Front.checkUptodateMeta
func (self *front_Front) checkUptodateMeta(_env *LnsEnv, lnsPath string,metaPath string,addSearchPath LnsAny) LnsAny {
    __func__ := "@lune.@base.@front.Front.checkUptodateMeta"
    var metaObj LnsAny
    
    {
        _metaObj := self.FP.loadLua(_env, metaPath)
        if _metaObj == nil{
            Log_log(_env, Log_Level__Warn, __func__, 863, Log_CreateMessage(func(_env *LnsEnv) string {
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
        Log_log(_env, Log_Level__Warn, __func__, 868, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.LuaVM.String_format("unmatch meta version -- %s", []LnsAny{metaPath})
        }))
        
        return nil
    }
    if meta.GetAt( "__hasTest" ).(bool){
        if meta.GetAt( "__enableTest" ).(bool) != self.option.Testing{
            Log_log(_env, Log_Level__Warn, __func__, 874, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.LuaVM.String_format("unmatch test setting -- %s", []LnsAny{metaPath})
            }))
            
            return nil
        }
    }
    {
        _exp4383 := meta.GetAt( "__dependModuleMap" ).(*Lns_luaValue)
        _key4383, _ := _exp4383.Get1stFromMap()
        for _key4383 != nil {
            moduleFullName := _key4383.(string)
            {
                _moduleLnsPath := self.FP.SearchModule(_env, moduleFullName)
                if !Lns_IsNil( _moduleLnsPath ) {
                    moduleLnsPath := _moduleLnsPath.(string)
                    {
                        _moduleLuaPath := self.FP.searchLuaFile(_env, moduleFullName, addSearchPath)
                        if !Lns_IsNil( _moduleLuaPath ) {
                            moduleLuaPath := _moduleLuaPath.(string)
                            if Lns_op_not(Util_getReadyCode(_env, moduleLnsPath, metaPath)){
                                Log_log(_env, Log_Level__Warn, __func__, 885, Log_CreateMessage(func(_env *LnsEnv) string {
                                    return _env.LuaVM.String_format("not ready -- %s, %s", []LnsAny{moduleLnsPath, metaPath})
                                }))
                                
                                return nil
                            }
                            var moduleMetaPath string
                            moduleMetaPath = front_convExp4194(Lns_2DDD(_env.LuaVM.String_gsub(moduleLuaPath,"%.lua$", ".meta")))
                            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( Depend_existFile(_env, moduleMetaPath)) &&
                                _env.SetStackVal( Lns_op_not(Util_getReadyCode(_env, moduleMetaPath, metaPath))) ).(bool)){
                                Log_log(_env, Log_Level__Warn, __func__, 893, Log_CreateMessage(func(_env *LnsEnv) string {
                                    return _env.LuaVM.String_format("not ready -- %s, %s", []LnsAny{moduleMetaPath, metaPath})
                                }))
                                
                                return nil
                            }
                        } else {
                            Log_log(_env, Log_Level__Warn, __func__, 898, Log_CreateMessage(func(_env *LnsEnv) string {
                                return _env.LuaVM.String_format("not found .lua file for -- %s", []LnsAny{moduleFullName})
                            }))
                            
                            return nil
                        }
                    }
                } else {
                    Log_log(_env, Log_Level__Warn, __func__, 903, Log_CreateMessage(func(_env *LnsEnv) string {
                        return _env.LuaVM.String_format("not found .lns file -- %s", []LnsAny{moduleFullName})
                    }))
                    
                    return nil
                }
            }
            _key4383, _ = _exp4383.NextFromMap( _key4383 )
        }
    }
    return NewFrontInterface_ModuleMeta(_env, lnsPath, &FrontInterface_MetaOrModule__Meta{meta})
}

// 916: decl @lune.@base.@front.Front.loadModule
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
                        luaPath = front_convExp4523(Lns_2DDD(_env.LuaVM.String_gsub(lnsPath, "%.lns$", ".lua")))
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
                            luaPath_645 := luaPath.(string)
                            if Util_getReadyCode(_env, lnsPath, luaPath_645){
                                var metaPath string
                                metaPath = front_convExp4581(Lns_2DDD(_env.LuaVM.String_gsub(luaPath_645, "%.lua$", ".meta")))
                                if Util_getReadyCode(_env, lnsPath, metaPath){
                                    {
                                        _preLoadInfo := self.preloadedModMap.Get(mod)
                                        if !Lns_IsNil( _preLoadInfo ) {
                                            preLoadInfo := _preLoadInfo
                                            loadVal = preLoadInfo
                                            
                                        } else {
                                            loadVal = self.FP.loadLua(_env, luaPath_645)
                                            
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
                            Log_log(_env, Log_Level__Warn, __func__, 970, Log_CreateMessage(func(_env *LnsEnv) string {
                                return _env.LuaVM.String_format("load from the binding -- %s", []LnsAny{mod})
                            }))
                            
                            var workMod LnsAny
                            workMod = Lns_require(mod)
                            var meta *FrontInterface_ModuleMeta
                            meta = NewFrontInterface_ModuleMeta(_env, Lns_car(_env.LuaVM.String_gsub(mod,"%.", "/")).(string) + ".lns", &FrontInterface_MetaOrModule__Meta{Lns_unwrap( Front_loadFromLuaTxt(_env, "return {}"))})
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

// 991: decl @lune.@base.@front.Front.getLuaModulePath
func (self *front_Front) GetLuaModulePath(_env *LnsEnv, mod string) string {
    return self.gomodMap.FP.GetLuaModulePath(_env, mod)
}

// 995: decl @lune.@base.@front.Front.loadMeta
func (self *front_Front) LoadMeta(_env *LnsEnv, importModuleInfo *FrontInterface_ImportModuleInfo,mod string) LnsAny {
    __func__ := "@lune.@base.@front.Front.loadMeta"
    var orgMod string
    orgMod = mod
    mod = self.gomodMap.FP.GetLuaModulePath(_env, mod)
    
    {
        _creater := self.mod2astCreate.Get(mod)
        if !Lns_IsNil( _creater ) {
            creater := _creater.(*front_AstCreater)
            var ast *TransUnit_ASTInfo
            var moduleInfo *FrontInterface_ModuleInfo
            ast,moduleInfo = creater.FP.getAst(_env)
            return NewFrontInterface_ModuleMeta(_env, ast.FP.Get_streamName(_env), &FrontInterface_MetaOrModule__Module{moduleInfo})
        }
    }
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
                            luaPath = front_convExp5035(Lns_2DDD(_env.LuaVM.String_gsub(lnsPath, "%.lns$", ".lua")))
                            {
                                _dir := self.option.OutputDir
                                if !Lns_IsNil( _dir ) {
                                    dir := _dir.(string)
                                    luaPath = self.FP.searchLuaFile(_env, orgMod, dir)
                                    
                                }
                            }
                            if luaPath != nil{
                                luaPath_696 := luaPath.(string)
                                var forceFlag bool
                                switch _exp5105 := self.option.TransCtrlInfo.UptodateMode.(type) {
                                case *Types_CheckingUptodateMode__ForceAll:
                                    forceFlag = true
                                    
                                case *Types_CheckingUptodateMode__Force1:
                                forceMod := _exp5105.Val1
                                    forceFlag = mod == forceMod
                                    
                                case *Types_CheckingUptodateMode__Normal:
                                    forceFlag = false
                                    
                                case *Types_CheckingUptodateMode__Touch:
                                    forceFlag = false
                                    
                                }
                                if Lns_op_not(forceFlag){
                                    if Util_getReadyCode(_env, lnsPath, luaPath_696){
                                        var metaPath string
                                        metaPath = front_convExp5133(Lns_2DDD(_env.LuaVM.String_gsub(luaPath_696, "%.lua$", ".meta")))
                                        if Util_getReadyCode(_env, lnsPath, metaPath){
                                            meta = self.FP.checkUptodateMeta(_env, lnsPath, metaPath, self.option.OutputDir)
                                            
                                        } else { 
                                            Log_log(_env, Log_Level__Warn, __func__, 1046, Log_CreateMessage(func(_env *LnsEnv) string {
                                                return _env.LuaVM.String_format("%s not ready meta %s, %s", []LnsAny{orgMod, lnsPath, metaPath})
                                            }))
                                            
                                        }
                                    } else { 
                                        Log_log(_env, Log_Level__Warn, __func__, 1050, Log_CreateMessage(func(_env *LnsEnv) string {
                                            return _env.LuaVM.String_format("%s not ready lua %s, %s", []LnsAny{orgMod, lnsPath, luaPath_696})
                                        }))
                                        
                                    }
                                } else { 
                                    Log_log(_env, Log_Level__Warn, __func__, 1054, Log_CreateMessage(func(_env *LnsEnv) string {
                                        return _env.LuaVM.String_format("force analyze -- %s", []LnsAny{orgMod})
                                    }))
                                    
                                }
                            } else {
                                Log_log(_env, Log_Level__Warn, __func__, 1058, Log_CreateMessage(func(_env *LnsEnv) string {
                                    return _env.LuaVM.String_format("%s not found lua in %s", []LnsAny{orgMod, self.option.OutputDir})
                                }))
                                
                            }
                        }
                        if meta != nil{
                            meta_720 := meta.(*FrontInterface_ModuleMeta)
                            self.moduleMgr.FP.AddMeta(_env, mod, meta_720)
                        } else {
                            var metawork *FrontInterface_ModuleMeta
                            var luaTxt string
                            metawork,luaTxt = self.FP.loadParserToLuaCode(_env, importModuleInfo, &Parser_ParserSrc__LnsPath{lnsPath, orgMod}, lnsPath, orgMod)
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
                                var meta *FrontInterface_ModuleMeta
                                var luaTxt string
                                meta,luaTxt = self.FP.loadParserToLuaCode(_env, importModuleInfo, &Parser_ParserSrc__LnsCode{lnsCode, mod}, path, mod)
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

// 1092: decl @lune.@base.@front.Front.dumpTokenize
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

// 1104: decl @lune.@base.@front.Front.dumpAst
func (self *front_Front) DumpAst(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    Depend_profile(_env, self.option.ValidProf, conv2Form5722(func(_env *LnsEnv) {
        var ast *TransUnit_ASTInfo
        ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsPath{scriptPath, mod}, mod, front_getModuleId_1579_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, DumpNode_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env), Lns_io_stdout), DumpNode_Opt2Stem(NewDumpNode_Opt(_env, "", 0)))
    }), scriptPath + ".profi")
}

// 1124: decl @lune.@base.@front.Front.format
func (self *front_Front) Format(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsPath{scriptPath, mod}, mod, front_getModuleId_1579_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, Formatter_createFilter(_env, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), Lns_io_stdout), Formatter_Opt2Stem(NewFormatter_Opt(_env, ast.FP.Get_node(_env))))
}

// 1140: decl @lune.@base.@front.Front.checkDiag
func (self *front_Front) CheckDiag(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    Util_setErrorCode(_env, 0)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsPath{scriptPath, mod}, mod, front_getModuleId_1579_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Diag, nil)
}

// 1150: decl @lune.@base.@front.Front.complete
func (self *front_Front) Complete(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsPath{scriptPath, mod}, mod, front_getModuleId_1579_(_env, scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Complete, self.option.AnalyzePos)
}

// 1158: decl @lune.@base.@front.Front.inquire
func (self *front_Front) Inquire(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsPath{scriptPath, mod}, mod, front_getModuleId_1579_(_env, scriptPath, mod, nil, nil), self.option.AnalyzeModule, TransUnit_AnalyzeMode__Inquire, self.option.AnalyzePos)
}

// 1167: decl @lune.@base.@front.Front.createGlue
func (self *front_Front) CreateGlue(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsPath{scriptPath, mod}, mod, front_getModuleId_1579_(_env, scriptPath, mod, nil, nil), nil, TransUnit_AnalyzeMode__Compile, nil)
    var filter *Nodes_Filter
    filter = GlueFilter_createFilter(_env, self.option.OutputDir)
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, filter, 0)
}

// 1450: decl @lune.@base.@front.Front.getGoAppName
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

// 1457: decl @lune.@base.@front.Front.createGoOption
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
                parentPath = front_convExp7469(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(scriptPath,"/[^/]+$", "")).(string),".*/", "")))
                if len(parentPath) == 0{
                    packageName = "main"
                    
                } else if parentPath == "."{
                    packageName = "main"
                    
                } else if parentPath == ".."{
                    packageName = "main"
                    
                } else { 
                    packageName = front_convExp7526(Lns_2DDD(_env.LuaVM.String_gsub(parentPath,"[^%w]", "")))
                    
                }
            }
        }
    }
    return NewConvGo_Option(_env, packageName, self.FP.getGoAppName(_env), self.option.MainModule, self.option.FP.Get_addEnvArg(_env), self.option.FP.Get_enableRunner(_env))
}

// 1485: decl @lune.@base.@front.Front.convertToLua
func (self *front_Front) convertToLua(_env *LnsEnv, scriptPath string,convMode LnsInt,streamLua Lns_oStream,streamMeta Lns_oStream) LnsAny {
    var mod string
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    var moduleId *FrontInterface_ModuleId
    moduleId = front_getModuleId_1579_(_env, scriptPath, mod, nil, nil)
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsPath{scriptPath, mod}, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
    var metaTxt string
    var luaTxt string
    metaTxt,luaTxt = self.FP.convertFromAst(_env, ast, scriptPath, convMode)
    streamLua.Write(_env, luaTxt)
    streamMeta.Write(_env, metaTxt)
    if _switch7716 := self.option.ConvTo; _switch7716 == Types_Lang__Go {
        var conv *Nodes_Filter
        conv = ConvGo_createFilter(_env, self.option.Testing, "stdout", streamLua, ast, self.FP.createGoOption(_env, scriptPath))
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast.FP.Get_node(_env))))
    }
    return ast
}

// 1560: decl @lune.@base.@front.Front.saveToGo
func (self *front_Front) saveToGo(_env *LnsEnv, scriptPath string,ast *TransUnit_ASTInfo) LnsAny {
    var rootNode *Nodes_RootNode
    
    {
        _rootNode := Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)
        if _rootNode == nil{
            return nil
        } else {
            rootNode = _rootNode.(*Nodes_RootNode)
        }
    }
    for _pragma := range( rootNode.FP.Get_luneHelperInfo(_env).PragmaSet.Items ) {
        pragma := _pragma
        switch _exp7983 := pragma.(type) {
        case *LuneControl_Pragma__limit_conv_code:
        codeSet := _exp7983.Val1
            if Lns_op_not(codeSet.Has(LuneControl_Code__Go)){
                return nil
            }
        }
    }
    return Newfront_GoConverter(_env, scriptPath, ast, self.option, self.FP.createGoOption(_env, scriptPath))
}

// 1582: decl @lune.@base.@front.Front.saveToC
func (self *front_Front) SaveToC(_env *LnsEnv, scriptPath string,ast *TransUnit_ASTInfo) {
    var cPath string
    cPath = front_convExp8031(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".c")))
    var file Lns_luaStream
    
    {
        _file := front_convExp8046(Lns_2DDD(Lns_io_open(cPath, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    var hPath string
    hPath = front_convExp8060(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".h")))
    var hFile Lns_luaStream
    
    {
        _hFile := front_convExp8076(Lns_2DDD(Lns_io_open(hPath, "w")))
        if _hFile == nil{
            return 
        } else {
            hFile = _hFile.(Lns_luaStream)
        }
    }
    file.Close(_env)
    hFile.Close(_env)
}

// 1598: decl @lune.@base.@front.Front.outputBuiltin
func (self *front_Front) OutputBuiltin(_env *LnsEnv, scriptPath string) {
    var mod string
    mod = self.FP.scriptPath2Module(_env, "lns_builtin")
    var ast *TransUnit_ASTInfo
    ast = self.FP.createAst(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsCode{"", mod}, mod, FrontInterface_ModuleId_createId(_env, 0.0, 0), nil, TransUnit_AnalyzeMode__Compile, nil)
    self.FP.SaveToC(_env, scriptPath, ast)
}

// 1620: decl @lune.@base.@front.Front.saveToLua
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
    mod = self.FP.scriptPath2Module(_env, scriptPath)
    var luaPath string
    luaPath = front_convExp8220(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".lua")))
    var metaPath string
    metaPath = front_convExp8233(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".meta")))
    if Lns_isCondTrue( self.option.OutputDir){
        var filename string
        filename = front_convExp8250(Lns_2DDD(_env.LuaVM.String_gsub(mod,"%.", "/")))
        luaPath = _env.LuaVM.String_format("%s/%s.lua", []LnsAny{self.option.OutputDir, filename})
        
        metaPath = _env.LuaVM.String_format("%s/%s.meta", []LnsAny{self.option.OutputDir, filename})
        
    }
    var convMode LnsInt
    convMode = ConvLua_ConvMode__ConvMeta
    if luaPath == scriptPath{
        Util_errorLog(_env, _env.LuaVM.String_format("%s is illegal filename.", []LnsAny{luaPath}))
        return nil
    }
    switch _exp8643 := uptodate.(type) {
    case *front_ModuleUptodate__NeedUpdate:
        var process func(_env *LnsEnv, ast *TransUnit_ASTInfo) LnsAny
        process = func(_env *LnsEnv, ast *TransUnit_ASTInfo) LnsAny {
            var luaConv *front_LuaConverter
            luaConv = Newfront_LuaConverter(_env, luaPath, metaPath, dependsPath, ast, convMode, scriptPath, self.option.ByteCompile, self.option.StripDebugInfo, self.option)
            var goConv LnsAny
            goConv = nil
            if _switch8401 := self.option.ConvTo; _switch8401 == Types_Lang__C {
                self.FP.SaveToC(_env, scriptPath, ast)
            } else if _switch8401 == Types_Lang__Go {
                goConv = self.FP.saveToGo(_env, scriptPath, ast)
                
            }
            return front_ConverterFuncSub_1145_(func(_env *LnsEnv) {
                luaConv.FP.SaveLua(_env)
                _env.NilAccFin(_env.NilAccPush(goConv) && 
                Lns_NilAccCall0( _env, func () {_env.NilAccPop().(*front_GoConverter).FP.saveGo(_env)})/* 1679:16 */)
            })
        }
        var result LnsAny
        result = self.FP.createAstSub(_env, NewFrontInterface_ImportModuleInfo(_env), &Parser_ParserSrc__LnsPath{scriptPath, mod}, mod, moduleId, nil, TransUnit_AnalyzeMode__Compile, nil)
        switch _exp8501 := result.(type) {
        case *front_CreateAstResult__Ast:
        ast := _exp8501.Val1
            return front_ConverterFunc_1147_(func(_env *LnsEnv) LnsAny {
                return process(_env, ast)
            })
        case *front_CreateAstResult__Creater:
        astCreater := _exp8501.Val1
            return front_ConverterFunc_1147_(func(_env *LnsEnv) LnsAny {
                return process(_env, self.FP.processAstCreater(_env, astCreater))
            })
        }
    case *front_ModuleUptodate__Uptodate:
    metaInfo := _exp8643.Val1
        Util_errorLog(_env, "uptodate -- " + scriptPath)
        var dependsStream LnsAny
        dependsStream = self.option.FP.OpenDepend(_env, dependsPath)
        front_outputDependInfo_2249_(_env, dependsStream, metaInfo, mod)
        front_closeStreams_2084_(_env, nil, nil, dependsStream, metaPath, false)
        return nil
    case *front_ModuleUptodate__NeedTouch:
    metaCode := _exp8643.Val1
    metaInfo := _exp8643.Val2
        Util_errorLog(_env, "touch -- " + scriptPath)
        var dependsStream LnsAny
        dependsStream = self.option.FP.OpenDepend(_env, dependsPath)
        var metaMemStream *Util_memStream
        metaMemStream = NewUtil_memStream(_env)
        if self.option.Mode == Option_ModeKind__SaveMeta{
            metaMemStream.FP.Write(_env, metaCode)
        }
        front_outputDependInfo_2249_(_env, dependsStream, metaInfo, mod)
        front_closeStreams_2084_(_env, nil, metaMemStream, dependsStream, metaPath, self.option.Mode == Option_ModeKind__SaveMeta)
        return nil
    }
// insert a dummy
    return nil
}

// 1728: decl @lune.@base.@front.Front.outputBootC
func (self *front_Front) outputBootC(_env *LnsEnv, scriptPath string) {
}

// 1766: decl @lune.@base.@front.Front.build
func (self *front_Front) Build(_env *LnsEnv, buildMode LnsAny,astCallback LnsAny) {
    var createUpdateInfo func(_env *LnsEnv, scriptPath string,dependsPath LnsAny) *front_UpdateInfo
    createUpdateInfo = func(_env *LnsEnv, scriptPath string,dependsPath LnsAny) *front_UpdateInfo {
        var mod string
        mod = self.FP.scriptPath2Module(_env, scriptPath)
        var moduleId *FrontInterface_ModuleId
        var uptodate LnsAny
        moduleId,uptodate = self.FP.getModuleIdAndCheckUptodate(_env, scriptPath, mod)
        return Newfront_UpdateInfo(_env, scriptPath, dependsPath, moduleId, uptodate)
    }
    var process func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo) LnsAny
    process = func(_env *LnsEnv, oneShot bool,updateInfo *front_UpdateInfo) LnsAny {
        var mod string
        mod = self.FP.scriptPath2Module(_env, updateInfo.FP.Get_scriptPath(_env))
        switch _exp8892 := buildMode.(type) {
        case *front_BuildMode__Save:
            return self.FP.saveToLua(_env, updateInfo)
        case *front_BuildMode__Output:
        streamLua := _exp8892.Val1
        streamMeta := _exp8892.Val2
            self.FP.convertToLua(_env, updateInfo.FP.Get_scriptPath(_env), ConvLua_ConvMode__ConvMeta, streamLua, streamMeta)
        case *front_BuildMode__CreateAst:
            if Lns_op_not(self.moduleMgr.FP.GetAst(_env, mod)){
                self.FP.convertToLua(_env, updateInfo.FP.Get_scriptPath(_env), ConvLua_ConvMode__ConvMeta, NewUtil_NullOStream(_env).FP, NewUtil_NullOStream(_env).FP)
            }
        }
        return nil
    }
    Depend_profile(_env, self.option.ValidProf, conv2Form9167(func(_env *LnsEnv) {
        if self.option.ScriptPath == "@-"{
            for _, _path := range( self.option.BatchList.Items ) {
                path := _path.(string)
                self.targetSet.Add(self.FP.scriptPath2Module(_env, path))
            }
            var postProcessList *LnsList
            postProcessList = NewLnsList([]LnsAny{})
            for _, _path := range( self.option.BatchList.Items ) {
                path := _path.(string)
                var updateInfo *front_UpdateInfo
                updateInfo = createUpdateInfo(_env, path, Lns_car(_env.LuaVM.String_gsub(path,".lns$", ".d")).(string))
                var prev LnsReal
                prev = _env.LuaVM.OS_clock()
                Lns_print([]LnsAny{_env.LuaVM.String_format("%s: processing...", []LnsAny{updateInfo.FP.Get_scriptPath(_env)})})
                {
                    __exp := process(_env, false, updateInfo)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(front_ConverterFunc_1147_)
                        if self.option.FP.Get_validPostBuild(_env){
                            postProcessList.Insert(_exp)
                        } else { 
                            _exp(_env)
                        }
                    }
                }
                Lns_print([]LnsAny{_env.LuaVM.String_format("%s: done %g", []LnsAny{updateInfo.FP.Get_scriptPath(_env), _env.LuaVM.OS_clock() - prev})})
            }
            var subProcessList *LnsList
            subProcessList = NewLnsList([]LnsAny{})
            for _, _postProcess := range( postProcessList.Items ) {
                postProcess := _postProcess.(front_ConverterFunc_1147_)
                {
                    _subProcess := postProcess(_env)
                    if !Lns_IsNil( _subProcess ) {
                        subProcess := _subProcess.(front_ConverterFuncSub_1145_)
                        subProcessList.Insert(subProcess)
                    }
                }
            }
            for _, _subProcess := range( subProcessList.Items ) {
                subProcess := _subProcess.(front_ConverterFuncSub_1145_)
                subProcess(_env)
            }
        } else { 
            {
                _postProcess := process(_env, true, createUpdateInfo(_env, self.option.ScriptPath, nil))
                if !Lns_IsNil( _postProcess ) {
                    postProcess := _postProcess.(front_ConverterFunc_1147_)
                    {
                        _subProcess := postProcess(_env)
                        if !Lns_IsNil( _subProcess ) {
                            subProcess := _subProcess.(front_ConverterFuncSub_1145_)
                            subProcess(_env)
                        }
                    }
                }
            }
        }
        if astCallback != nil{
            astCallback_1169 := astCallback.(Front_AstCallback)
            for _, _key := range( self.moduleMgr.FP.GetModList(_env).Items ) {
                key := _key.(string)
                {
                    __exp := self.moduleMgr.FP.GetAst(_env, key)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*TransUnit_ASTInfo)
                        astCallback_1169(_env, _exp)
                    }
                }
            }
        }
    }), self.option.ScriptPath + ".profi")
}

// 1852: decl @lune.@base.@front.Front.exec
func (self *front_Front) Exec(_env *LnsEnv) {
    __func__ := "@lune.@base.@front.Front.exec"
    Log_log(_env, Log_Level__Trace, __func__, 1854, Log_CreateMessage(func(_env *LnsEnv) string {
        return Option_ModeKind_getTxt( self.option.Mode)
    }))
    
    if _switch9651 := self.option.Mode; _switch9651 == Option_ModeKind__Token {
        self.FP.DumpTokenize(_env, self.option.ScriptPath)
    } else if _switch9651 == Option_ModeKind__Ast {
        self.FP.DumpAst(_env, self.option.ScriptPath)
    } else if _switch9651 == Option_ModeKind__Format {
        self.FP.Format(_env, self.option.ScriptPath)
    } else if _switch9651 == Option_ModeKind__Diag {
        self.FP.CheckDiag(_env, self.option.ScriptPath)
    } else if _switch9651 == Option_ModeKind__Complete {
        self.FP.Complete(_env, self.option.ScriptPath)
    } else if _switch9651 == Option_ModeKind__Inquire {
        self.FP.Inquire(_env, self.option.ScriptPath)
    } else if _switch9651 == Option_ModeKind__Glue {
        self.FP.CreateGlue(_env, self.option.ScriptPath)
    } else if _switch9651 == Option_ModeKind__Lua || _switch9651 == Option_ModeKind__LuaMeta {
        var convMode LnsInt
        if self.option.Mode == Option_ModeKind__Lua{
            convMode = ConvLua_ConvMode__Convert
            
        } else { 
            convMode = ConvLua_ConvMode__ConvMeta
            
        }
        self.FP.convertToLua(_env, self.option.ScriptPath, convMode, Lns_io_stdout, Lns_io_stdout)
    } else if _switch9651 == Option_ModeKind__Save || _switch9651 == Option_ModeKind__SaveMeta {
        self.FP.Build(_env, front_BuildMode__Save_Obj, nil)
    } else if _switch9651 == Option_ModeKind__Shebang {
        {
            _modObj := front_convExp9469(Lns_2DDD(self.FP.LoadModule(_env, self.FP.scriptPath2Module(_env, self.option.ScriptPath))))
            if !Lns_IsNil( _modObj ) {
                modObj := _modObj
                var code LnsInt
                code = Depend_runMain(_env, modObj.(*Lns_luaValue).GetAt("__main"), self.option.ShebangArgList)
                _env.LuaVM.OS_exit(code)
            }
        }
    } else if _switch9651 == Option_ModeKind__Exec {
        _ = front_convExp9491(Lns_2DDD(self.FP.LoadModule(_env, self.FP.scriptPath2Module(_env, self.option.ScriptPath))))
        if self.option.Testing{
            var code string
            code = "local Testing = require( \"lune.base.Testing\" )\nreturn function( path )\n  Testing.run( path );\n  Testing.outputAllResult( io.stdout );\nend\n"
            var loaded LnsAny
            var mess LnsAny
            loaded,mess = _env.CommonLuaVM.Load(code, nil)
            if loaded != nil{
                loaded_1209 := loaded.(*Lns_luaValue)
                {
                    _mod := front_convExp9545(Lns_2DDD(_env.CommonLuaVM.RunLoadedfunc(loaded_1209,Lns_2DDD([]LnsAny{}))[0]))
                    if !Lns_IsNil( _mod ) {
                        mod := _mod
                        _env.CommonLuaVM.RunLoadedfunc((mod.(*Lns_luaValue)),Lns_2DDD([]LnsAny{self.FP.scriptPath2Module(_env, self.option.ScriptPath)}))
                    }
                }
            } else {
                Lns_print([]LnsAny{mess})
            }
        }
    } else if _switch9651 == Option_ModeKind__BootC {
        self.FP.outputBootC(_env, self.option.ScriptPath)
    } else if _switch9651 == Option_ModeKind__Builtin {
        self.FP.OutputBuiltin(_env, self.option.ScriptPath)
    } else if _switch9651 == Option_ModeKind__MkMain {
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
func front_DependMetaInfo__fromMap_1473_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_DependMetaInfo_FromMap( arg1, paramList )
}
func front_DependMetaInfo__fromStem_1477_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
func front_MetaForBuildId__fromMap_1505_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return front_MetaForBuildId_FromMap( arg1, paramList )
}
func front_MetaForBuildId__fromStem_1509_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
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
// 512: decl @lune.@base.@front.MetaForBuildId.createModuleId
func (self *front_MetaForBuildId) CreateModuleId(_env *LnsEnv) *FrontInterface_ModuleId {
    return FrontInterface_ModuleId_createIdFromTxt(_env, self.G__buildId)
}

// 517: decl @lune.@base.@front.MetaForBuildId.LoadFromMeta
func front_MetaForBuildId_LoadFromMeta_1513_(_env *LnsEnv, metaPath string)(LnsAny, LnsAny) {
    {
        _fileObj := front_convExp2210(Lns_2DDD(Lns_io_open(metaPath, nil)))
        if !Lns_IsNil( _fileObj ) {
            fileObj := _fileObj.(Lns_luaStream)
            var luaCode LnsAny
            luaCode = fileObj.Read(_env, "*a")
            fileObj.Close(_env)
            if luaCode != nil{
                luaCode_386 := luaCode.(string)
                var meta LnsAny
                meta = front_convExp2200(Lns_2DDD(front_MetaForBuildId__fromStem_1509_(_env, _env.CommonLuaVM.ExpandLuavalMap(Front_loadFromLuaTxt(_env, luaCode_386)),nil)))
                return meta, luaCode_386
            }
        }
    }
    return nil, nil
}


// declaration Class -- LuaConverter
type front_LuaConverterMtd interface {
    Join(_env *LnsEnv)
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    SaveLua(_env *LnsEnv)
    Start(_env *LnsEnv)
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
    converterFunc front_ConverterFunc_1147_
    filterInfo *ConvLua_FilterInfo
    rootNode *Nodes_RootNode
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
func Newfront_LuaConverter(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny, arg4 *TransUnit_ASTInfo, arg5 LnsInt, arg6 string, arg7 bool, arg8 bool, arg9 *Option_Option) *front_LuaConverter {
    obj := &front_LuaConverter{}
    obj.FP = obj
    obj.Runner_Runner.FP = obj
    obj.Initfront_LuaConverter(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
// 1339: DeclConstr
func (self *front_LuaConverter) Initfront_LuaConverter(_env *LnsEnv, luaPath string,metaPath string,dependsPath LnsAny,ast *TransUnit_ASTInfo,convMode LnsInt,path string,byteCompile bool,stripDebugInfo bool,option *Option_Option) {
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
    
    self.rootNode = Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode)
    
    self.converterFunc = front_ConverterFunc_1147_(func(_env *LnsEnv) LnsAny {
        var stream *Util_memStream
        stream = self.streamMem
        var metaStream *Util_memStream
        metaStream = self.metaStreamMem
        var outStream *Util_memStream
        var oMetaStream *Util_memStream
        outStream,oMetaStream = stream, metaStream
        var needDepends bool
        needDepends = option.DependsPath != nil
        if needDepends{
            ast.FP.Get_node(_env).FP.ProcessFilter(_env, OutputDepend_createFilter(_env, self.dependsStreamMem.FP), 1)
        }
        if byteCompile{
            outStream = self.byteStream
            
            oMetaStream = self.byteMetaStream
            
        }
        self.filterInfo = front_ast2LuaMain_1405_(_env, ast, path, outStream.FP, oMetaStream.FP, convMode, false, option)
        
        self.filterInfo.FP.OutputLua(_env, self.rootNode)
        return nil
    })
    
    self.FP.Start(_env)
}

// 1388: decl @lune.@base.@front.LuaConverter.runMain
func (self *front_LuaConverter) RunMain(_env *LnsEnv) {
    self.converterFunc(_env)
}

// 1392: decl @lune.@base.@front.LuaConverter.saveLua
func (self *front_LuaConverter) SaveLua(_env *LnsEnv) {
    self.FP.Join(_env)
    self.filterInfo.FP.OutputMeta(_env, self.rootNode)
    if self.byteCompile{
        self.streamMem.FP.Write(_env, front_byteCompileFromLuaTxt_1428_(_env, self.byteStream.FP.Get_txt(_env), self.stripDebugInfo))
        self.metaStreamMem.FP.Write(_env, front_byteCompileFromLuaTxt_1428_(_env, self.byteMetaStream.FP.Get_txt(_env), true))
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
        _streamDst := front_convExp7241(Lns_2DDD(Lns_io_open(self.luaPath, "w")))
        if _streamDst == nil{
            panic(_env.LuaVM.String_format("write open error -- %s", []LnsAny{self.luaPath}))
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
        dependsStreamDst_926 := dependsStreamDst.(Lns_oStream)
        dependsStreamDst_926.Write(_env, Lns_unwrap( dependTxt).(string))
    }
    front_closeStreams_2084_(_env, streamDst, metaMemStream, dependsStreamDst, self.metaPath, self.option.Mode == Option_ModeKind__SaveMeta)
}


// declaration Class -- GoConverter
type front_GoConverterMtd interface {
    Join(_env *LnsEnv)
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    saveGo(_env *LnsEnv)
    Start(_env *LnsEnv)
}
type front_GoConverter struct {
    Runner_Runner
    memStream *Util_memStream
    path string
    converter front_ConverterFunc_1147_
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
func Newfront_GoConverter(_env *LnsEnv, arg1 string, arg2 *TransUnit_ASTInfo, arg3 *Option_Option, arg4 *ConvGo_Option) *front_GoConverter {
    obj := &front_GoConverter{}
    obj.FP = obj
    obj.Runner_Runner.FP = obj
    obj.Initfront_GoConverter(_env, arg1, arg2, arg3, arg4)
    return obj
}
// 1516: DeclConstr
func (self *front_GoConverter) Initfront_GoConverter(_env *LnsEnv, scriptPath string,ast *TransUnit_ASTInfo,option *Option_Option,goOpt *ConvGo_Option) {
    self.InitRunner_Runner(_env)
    var path string
    path = front_convExp7769(Lns_2DDD(_env.LuaVM.String_gsub(scriptPath,"%.lns$", ".go")))
    {
        _dir := option.OutputDir
        if !Lns_IsNil( _dir ) {
            dir := _dir.(string)
            path = _env.LuaVM.String_format("%s/%s", []LnsAny{dir, path})
            
        }
    }
    self.path = path
    
    self.memStream = NewUtil_memStream(_env)
    
    self.converter = front_ConverterFunc_1147_(func(_env *LnsEnv) LnsAny {
        var conv *Nodes_Filter
        conv = ConvGo_createFilter(_env, option.Testing, path, self.memStream.FP, ast, goOpt)
        ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast.FP.Get_node(_env))))
        return nil
    })
    
    self.FP.Start(_env)
}

// 1539: decl @lune.@base.@front.GoConverter.runMain
func (self *front_GoConverter) RunMain(_env *LnsEnv) {
    self.converter(_env)
}

// 1543: decl @lune.@base.@front.GoConverter.saveGo
func (self *front_GoConverter) saveGo(_env *LnsEnv) {
    self.FP.Join(_env)
    var file Lns_luaStream
    
    {
        _file := front_convExp7910(Lns_2DDD(Lns_io_open(self.path, "w")))
        if _file == nil{
            return 
        } else {
            file = _file.(Lns_luaStream)
        }
    }
    file.Write(_env, self.memStream.FP.Get_txt(_env))
    file.Close(_env)
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
    Lns_Runner_init(_env)
    Lns_Builtin_init(_env)
    Depend_setup(_env, Depend_UpdateVer(front__anonymous_1008_))
    front_forceUpdateMeta = true
}
func init() {
    init_front = false
}
