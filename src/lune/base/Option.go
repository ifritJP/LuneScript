// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Option bool
var Option__mod__ string
// decl enum -- ModeKind 
type Option_ModeKind = string
const Option_ModeKind__Ast = "ast"
const Option_ModeKind__BootC = "bootC"
const Option_ModeKind__BuildAst = "buildAst"
const Option_ModeKind__Builtin = "builtin"
const Option_ModeKind__Complete = "comp"
const Option_ModeKind__Diag = "diag"
const Option_ModeKind__Exec = "exe"
const Option_ModeKind__Format = "format"
const Option_ModeKind__Glue = "glue"
const Option_ModeKind__GoMod = "gomod"
const Option_ModeKind__Indexer = "indexer"
const Option_ModeKind__Inquire = "inq"
const Option_ModeKind__Lua = "lua"
const Option_ModeKind__LuaMeta = "LUA"
const Option_ModeKind__MkMain = "mkmain"
const Option_ModeKind__Save = "save"
const Option_ModeKind__SaveMeta = "SAVE"
const Option_ModeKind__Shebang = "shebang"
const Option_ModeKind__Token = "token"
const Option_ModeKind__Unknown = ""
var Option_ModeKindList_ = NewLnsList( []LnsAny {
  Option_ModeKind__Unknown,
  Option_ModeKind__Token,
  Option_ModeKind__Ast,
  Option_ModeKind__Diag,
  Option_ModeKind__Complete,
  Option_ModeKind__Lua,
  Option_ModeKind__LuaMeta,
  Option_ModeKind__Save,
  Option_ModeKind__SaveMeta,
  Option_ModeKind__Exec,
  Option_ModeKind__Glue,
  Option_ModeKind__BootC,
  Option_ModeKind__Format,
  Option_ModeKind__Builtin,
  Option_ModeKind__Inquire,
  Option_ModeKind__MkMain,
  Option_ModeKind__Shebang,
  Option_ModeKind__BuildAst,
  Option_ModeKind__Indexer,
  Option_ModeKind__GoMod,
})
func Option_ModeKind_get__allList(_env *LnsEnv) *LnsList{
    return Option_ModeKindList_
}
var Option_ModeKindMap_ = map[string]string {
  Option_ModeKind__Ast: "ModeKind.Ast",
  Option_ModeKind__BootC: "ModeKind.BootC",
  Option_ModeKind__BuildAst: "ModeKind.BuildAst",
  Option_ModeKind__Builtin: "ModeKind.Builtin",
  Option_ModeKind__Complete: "ModeKind.Complete",
  Option_ModeKind__Diag: "ModeKind.Diag",
  Option_ModeKind__Exec: "ModeKind.Exec",
  Option_ModeKind__Format: "ModeKind.Format",
  Option_ModeKind__Glue: "ModeKind.Glue",
  Option_ModeKind__GoMod: "ModeKind.GoMod",
  Option_ModeKind__Indexer: "ModeKind.Indexer",
  Option_ModeKind__Inquire: "ModeKind.Inquire",
  Option_ModeKind__Lua: "ModeKind.Lua",
  Option_ModeKind__LuaMeta: "ModeKind.LuaMeta",
  Option_ModeKind__MkMain: "ModeKind.MkMain",
  Option_ModeKind__Save: "ModeKind.Save",
  Option_ModeKind__SaveMeta: "ModeKind.SaveMeta",
  Option_ModeKind__Shebang: "ModeKind.Shebang",
  Option_ModeKind__Token: "ModeKind.Token",
  Option_ModeKind__Unknown: "ModeKind.Unknown",
}
func Option_ModeKind__from(_env *LnsEnv, arg1 string) LnsAny{
    if _, ok := Option_ModeKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Option_ModeKind_getTxt(arg1 string) string {
    return Option_ModeKindMap_[arg1];
}
// decl enum -- Int2strMode 
type Option_Int2strMode = LnsInt
const Option_Int2strMode__Int2strModeDepend = 0
const Option_Int2strMode__Int2strModeNeed0 = 1
const Option_Int2strMode__Int2strModeUnneed0 = 2
var Option_Int2strModeList_ = NewLnsList( []LnsAny {
  Option_Int2strMode__Int2strModeDepend,
  Option_Int2strMode__Int2strModeNeed0,
  Option_Int2strMode__Int2strModeUnneed0,
})
func Option_Int2strMode_get__allList(_env *LnsEnv) *LnsList{
    return Option_Int2strModeList_
}
var Option_Int2strModeMap_ = map[LnsInt]string {
  Option_Int2strMode__Int2strModeDepend: "Int2strMode.Int2strModeDepend",
  Option_Int2strMode__Int2strModeNeed0: "Int2strMode.Int2strModeNeed0",
  Option_Int2strMode__Int2strModeUnneed0: "Int2strMode.Int2strModeUnneed0",
}
func Option_Int2strMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Option_Int2strModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Option_Int2strMode_getTxt(arg1 LnsInt) string {
    return Option_Int2strModeMap_[arg1];
}
// for 44
func Option_convExp0_70(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 356
func Option_convExp0_963(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 243
func Option_convExp0_749(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 361
func Option_convExp0_894(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// 85: decl @lune.@base.@Option.getRuntimeModule
func Option_getRuntimeModule(_env *LnsEnv) string {
    return _env.GetVM().String_format("lune.base.runtime%d", Lns_2DDD(Ver_luaModVersion))
}

// 236: decl @lune.@base.@Option.outputLuneMod
func Option_outputLuneMod(_env *LnsEnv, path LnsAny) LnsAny {
    var lune_path string
    lune_path = "runtime.lua"
    if path != nil{
        path_147 := path.(string)
        if path_147 != ""{
            lune_path = path_147
        }
    }
    var fileObj Lns_luaStream
    
    {
        _fileObj := Option_convExp0_749(Lns_2DDD(Lns_io_open(lune_path, "w")))
        if _fileObj == nil{
            return _env.GetVM().String_format("failed to open -- %s", Lns_2DDD(lune_path))
        } else {
            fileObj = _fileObj.(Lns_luaStream)
        }
    }
    fileObj.Write(_env, "--[[\nMIT License\n\nCopyright (c) 2018,2019 ifritJP\n\nPermission is hereby granted, free of charge, to any person obtaining a copy\nof this software and associated documentation files (the \"Software\"), to deal\nin the Software without restriction, including without limitation the rights\nto use, copy, modify, merge, publish, distribute, sublicense, and/or sell\ncopies of the Software, and to permit persons to whom the Software is\nfurnished to do so, subject to the following conditions:\n\nThe above copyright notice and this permission notice shall be included in all\ncopies or substantial portions of the Software.\n\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\nIMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\nFITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\nAUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\nLIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\nOUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE\nSOFTWARE.\n]]\n")
    for _, _kind := range( LuaMod_CodeKind_get__allList(_env).Items ) {
        kind := _kind.(LnsInt)
        fileObj.Write(_env, LuaMod_getCode(_env, kind))
    }
    fileObj.Close(_env)
    return nil
}

// 281: decl @lune.@base.@Option.analyze
func Option_analyze(_env *LnsEnv, argList *LnsList2_[string]) *Option_Option {
    __func__ := "@lune.@base.@Option.analyze"
    var option *Option_Option
    option = NewOption_Option(_env)
    var useStdInFlag bool
    useStdInFlag = false
    var lineNo LnsAny
    lineNo = nil
    var column LnsAny
    column = nil
    var index LnsInt
    index = 1
    {
        _file := Option_convExp0_963(Lns_2DDD(Lns_io_open("lune.js", "r")))
        if !Lns_IsNil( _file ) {
            file := _file.(Lns_luaStream)
            var projInfo LnsAny
            var mess LnsAny
            projInfo,mess = Option_ProjInfo1208__fromStem_3_(_env, Lns_car(Json_fromStr(_env, _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( file.Read(_env, "*a")) ||
                _env.SetStackVal( "") ).(string))),nil)
            if projInfo != nil{
                projInfo_175 := projInfo.(*Option_ProjInfo1208)
                var workArgList *LnsList2_[string]
                workArgList = NewLnsList2_[string]([]string{})
                for _, _arg := range( projInfo_175.Cmd_option.Items ) {
                    arg := _arg
                    workArgList.Insert(arg)
                }
                for _, _arg := range( argList.Items ) {
                    arg := _arg
                    workArgList.Insert(arg)
                }
                argList = workArgList
            } else {
                Lns_print(Lns_2DDD("failed to load -- lune.js", mess))
            }
            file.Close(_env)
        }
    }
    var Option_getNextOp func(_env *LnsEnv) LnsAny
    Option_getNextOp = func(_env *LnsEnv) LnsAny {
        if argList.Len() <= index{
            return nil
        }
        index = index + 1
        return argList.GetAt(index)
    }
    var Option_getNextOpNonNil func(_env *LnsEnv) string
    Option_getNextOpNonNil = func(_env *LnsEnv) string {
        {
            _nextOp := Option_getNextOp(_env)
            if !Lns_IsNil( _nextOp ) {
                nextOp := _nextOp.(string)
                return nextOp
            }
        }
        Option_analyze__printUsage_0_(_env, 1)
    // insert a dummy
        return ""
    }
    var Option_getNextOpInt func(_env *LnsEnv) LnsInt
    Option_getNextOpInt = func(_env *LnsEnv) LnsInt {
        {
            _num := Lns_tonumber(Option_getNextOpNonNil(_env), nil)
            if !Lns_IsNil( _num ) {
                num := _num.(LnsReal)
                return (LnsInt)(num)
            }
        }
        Option_analyze__printUsage_0_(_env, 1)
    // insert a dummy
        return 0
    }
    Util_setDebugFlag(_env, false)
    var uptodateOpt LnsAny
    uptodateOpt = nil
    for argList.Len() >= index {
        var arg string
        arg = argList.GetAt(index)
        if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(arg,"^-", nil, nil))){
            if option.Mode != Option_ModeKind__Shebang{
                if _switch2 := (arg); _switch2 == "-i" {
                    useStdInFlag = true
                } else if _switch2 == "--tokenizerPipeSize" {
                    AsyncTokenizer_setDefaultPipeSize(_env, Option_getNextOpInt(_env))
                } else if _switch2 == "-prof" {
                    option.ValidProf = true
                } else if _switch2 == "--noEnvArg" {
                    option.addEnvArg = false
                } else if _switch2 == "--noUseWaiter" {
                    option.TransCtrlInfo.UseWaiter = false
                    option.sortGenerateCode = false
                } else if _switch2 == "--debug-dump-ast" {
                    option.DumpDebugAst = true
                } else if _switch2 == "--unsortGenerateCode" {
                    option.sortGenerateCode = false
                } else if _switch2 == "--disableMultiPhaseAst" {
                    option.TransCtrlInfo.ValidMultiPhaseTransUnit = false
                } else if _switch2 == "--disableMultiThreadAst" {
                    option.TransCtrlInfo.ThreadPerUnitThread = 0
                } else if _switch2 == "--threadPerUnitThread" {
                    option.TransCtrlInfo.ThreadPerUnitThread = Option_getNextOpInt(_env)
                    if option.TransCtrlInfo.ThreadPerUnitThread < 0{
                        Option_analyze__printUsage_0_(_env, 1)
                    }
                } else if _switch2 == "--convGoRunnerNum" {
                    option.ConvGoRunnerNum = Option_getNextOpInt(_env)
                    if option.ConvGoRunnerNum < 0{
                        Option_analyze__printUsage_0_(_env, 1)
                    }
                } else if _switch2 == "--macroAsyncParseStmtLen" {
                    option.TransCtrlInfo.MacroAsyncParseStmtLen = Option_getNextOpInt(_env)
                    if option.TransCtrlInfo.MacroAsyncParseStmtLen < 0{
                        Option_analyze__printUsage_0_(_env, 1)
                    }
                } else if _switch2 == "--legacyNewName" {
                    option.legacyNewName = true
                } else if _switch2 == "--enableMacroAsync" {
                    option.TransCtrlInfo.ValidMacroAsync = true
                } else if _switch2 == "--disableRunner" {
                    option.enableRunner = false
                } else if _switch2 == "--disablePostBuild" {
                    option.validPostBuild = false
                } else if _switch2 == "--enableAsyncCtl" {
                    option.TransCtrlInfo.ValidAsyncCtrl = true
                } else if _switch2 == "--disableGenInherit" {
                    option.TransCtrlInfo.DefaultGenInherit = false
                } else if _switch2 == "--defaultAsync" {
                    option.TransCtrlInfo.DefaultAsync = true
                } else if _switch2 == "--limitThread" {
                    var nextArg string
                    
                    {
                        _nextArg := Option_getNextOp(_env)
                        if _nextArg == nil{
                            Option_analyze__printUsage_0_(_env, 1)
                        } else {
                            nextArg = _nextArg.(string)
                        }
                    }
                    var num LnsReal
                    
                    {
                        _num := Lns_tonumber(nextArg, nil)
                        if _num == nil{
                            Option_analyze__printUsage_0_(_env, 1)
                        } else {
                            num = _num.(LnsReal)
                        }
                    }
                    Depend_setRuntimeThreadLimit(_env, (LnsInt)(num))
                } else if _switch2 == "--nodebug" {
                    Util_setDebugFlag(_env, false)
                } else if _switch2 == "--debug" {
                    Util_setDebugFlag(_env, true)
                } else if _switch2 == "-shebang" {
                    option.Mode = Option_ModeKind__Shebang
                } else if _switch2 == "--version" {
                    Util_println(_env, Lns_2DDD(_env.GetVM().String_format("LuneScript: version %s (%d:Lua%s) [%s]", Lns_2DDD(Ver_version, Option_getBuildCount_1_(_env), Depend_getLuaVersion(_env), Ver_metaVersion))))
                    _env.GetVM().OS_exit(0)
                } else if _switch2 == "--projDir" {
                    option.projDir = Option_getNextOp(_env)
                } else if _switch2 == "--builtin" {
                    var builtin *Builtin_Builtin
                    builtin = NewBuiltin_Builtin(_env, option.TargetLuaVer, option.TransCtrlInfo)
                    _ = builtin.FP.RegistBuiltInScope(_env)
                    {
                        __forsortCollection0 := Ast_getBuiltInTypeIdMap(_env)
                        __forsortSorted0 := __forsortCollection0.CreateKeyListInt()
                        __forsortSorted0.Sort( _env, LnsItemKindInt, nil )
                        for _, _typeId := range( __forsortSorted0.Items ) {
                            builtinTypeInfo := __forsortCollection0.Items[ _typeId ]
                            typeId := _typeId
                            var parentName string
                            parentName = builtinTypeInfo.FP.Get_typeInfo(_env).FP.GetParentFullName(_env, Ast_defaultTypeNameCtrl, nil, nil)
                            var dispName string
                            dispName = builtinTypeInfo.FP.Get_typeInfo(_env).FP.Get_display_stirng(_env)
                            Util_println(_env, Lns_2DDD(typeId, _env.GetVM().String_format("%s%s", Lns_2DDD(parentName, dispName))))
                        }
                    }
                    _env.GetVM().OS_exit(0)
                } else if _switch2 == "-mklunemod" {
                    var path LnsAny
                    path = Option_getNextOp(_env)
                    {
                        _mess := Option_outputLuneMod(_env, path)
                        if !Lns_IsNil( _mess ) {
                            mess := _mess.(string)
                            Util_errorLog(_env, mess)
                            _env.GetVM().OS_exit(1)
                        }
                    }
                    _env.GetVM().OS_exit(0)
                } else if _switch2 == "--mkbuiltin" {
                    var path string
                    
                    {
                        _path := Option_getNextOp(_env)
                        if _path == nil{
                            path = "."
                        } else {
                            path = _path.(string)
                        }
                    }
                    option.ScriptPath = path + "/lns_builtin.lns"
                    option.Mode = Option_ModeKind__Builtin
                } else if _switch2 == "-r" {
                    option.UseLuneModule = Option_getRuntimeModule(_env)
                } else if _switch2 == "--runtime" {
                    option.UseLuneModule = Option_getNextOp(_env)
                } else if _switch2 == "-oc" {
                    option.BootPath = Option_getNextOp(_env)
                } else if _switch2 == "-u" {
                    option.UpdateOnLoad = true
                } else if _switch2 == "-Werror" {
                    option.TransCtrlInfo.StopByWarning = true
                } else if _switch2 == "--disable-checking-define-abbr" {
                    option.TransCtrlInfo.CheckingDefineAbbr = false
                } else if _switch2 == "--disable-checking-mutable" {
                    option.TransCtrlInfo.ValidCheckingMutable = false
                } else if _switch2 == "--legacy-mutable-control" {
                    option.TransCtrlInfo.LegacyMutableControl = true
                } else if _switch2 == "--compat-comment" {
                    option.TransCtrlInfo.CompatComment = true
                } else if _switch2 == "--warning-shadowing" {
                    option.TransCtrlInfo.WarningShadowing = true
                } else if _switch2 == "--valid-luaval" {
                    option.TransCtrlInfo.ValidLuaval = true
                } else if _switch2 == "--default-lazy" {
                    option.TransCtrlInfo.DefaultLazy = true
                } else if _switch2 == "--package" {
                    option.PackageName = Option_getNextOp(_env)
                } else if _switch2 == "--int2str" {
                    var opt LnsAny
                    opt = Option_getNextOp(_env)
                    if _switch0 := opt; _switch0 == "depend" {
                        option.FP.Get_runtimeOpt(_env).int2strMode = Option_Int2strMode__Int2strModeDepend
                    } else if _switch0 == "need0" {
                        option.FP.Get_runtimeOpt(_env).int2strMode = Option_Int2strMode__Int2strModeNeed0
                    } else if _switch0 == "unneed0" {
                        option.FP.Get_runtimeOpt(_env).int2strMode = Option_Int2strMode__Int2strModeUnneed0
                    } else {
                        Util_errorLog(_env, _env.GetVM().String_format("unknown mode -- %s", Lns_2DDD(opt)))
                        _env.GetVM().OS_exit(1)
                    }
                } else if _switch2 == "--app" {
                    {
                        __exp := Option_getNextOp(_env)
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(string)
                            option.AppName = _exp
                        }
                    }
                } else if _switch2 == "--main" {
                    {
                        __exp := Option_getNextOp(_env)
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(string)
                            option.MainModule = _exp
                        }
                    }
                } else if _switch2 == "--log" {
                    {
                        _txt := Option_getNextOp(_env)
                        if !Lns_IsNil( _txt ) {
                            txt := _txt.(string)
                            {
                                _level := Log_str2level(_env, txt)
                                if !Lns_IsNil( _level ) {
                                    level := _level.(LnsInt)
                                    Log_setLevel(_env, level)
                                } else {
                                    Util_errorLog(_env, _env.GetVM().String_format("illegal level -- %s", Lns_2DDD(txt)))
                                }
                            }
                        }
                    }
                } else if _switch2 == "--testing" {
                    option.Testing = true
                    option.TransCtrlInfo.Testing = true
                } else if _switch2 == "--enableTestBlock" {
                    option.TransCtrlInfo.Testing = true
                } else if _switch2 == "--depends" {
                    option.DependsPath = Option_getNextOp(_env)
                } else if _switch2 == "--use-ipairs" {
                    option.UseIpairs = true
                } else if _switch2 == "--uptodate" {
                    uptodateOpt = Option_getNextOp(_env)
                } else if _switch2 == "-noLua" {
                    option.NoLua = true
                } else if _switch2 == "-langC" {
                    option.ConvTo = Types_Lang__C
                    option.TransCtrlInfo.ValidLuaval = true
                } else if _switch2 == "-langGo" {
                    option.ConvTo = Types_Lang__Go
                    option.TransCtrlInfo.ValidLuaval = true
                    option.TransCtrlInfo.ValidAsyncCtrl = true
                } else if _switch2 == "-langPython" {
                    option.ConvTo = Types_Lang__Python
                    option.TransCtrlInfo.ValidLuaval = false
                    option.TransCtrlInfo.ValidAsyncCtrl = false
                } else if _switch2 == "-ol" {
                    {
                        _txt := Option_getNextOp(_env)
                        if !Lns_IsNil( _txt ) {
                            txt := _txt.(string)
                            if _switch1 := txt; _switch1 == "51" {
                                option.TargetLuaVer = LuaVer_ver51
                            } else if _switch1 == "52" {
                                option.TargetLuaVer = LuaVer_ver52
                            } else if _switch1 == "53" {
                                option.TargetLuaVer = LuaVer_ver53
                            }
                        }
                    }
                } else if _switch2 == "-ob0" || _switch2 == "-ob1" {
                    option.ByteCompile = true
                    if arg == "-ob0"{
                        option.StripDebugInfo = true
                    }
                } else {
                    Util_log(_env, _env.GetVM().String_format("unknown option -- '%s'", Lns_2DDD(arg)))
                    _env.GetVM().OS_exit(1)
                }
            } else { 
                if option.ShebangArgList.Len() == 0{
                    option.ShebangArgList.Insert(option.ScriptPath)
                }
                option.ShebangArgList.Insert(arg)
            }
        } else { 
            if option.ScriptPath == ""{
                option.ScriptPath = arg
                if option.Mode == Option_ModeKind__Shebang{
                    option.ShebangArgList.Insert(option.ScriptPath)
                }
            } else if option.Mode == ""{
                {
                    _mode := Option_ModeKind__from(_env, arg)
                    if !Lns_IsNil( _mode ) {
                        mode := _mode.(string)
                        option.Mode = mode
                    } else {
                        Util_err(_env, _env.GetVM().String_format("unknown mode -- %s", Lns_2DDD(arg)))
                    }
                }
            } else { 
                if _switch3 := (option.Mode); _switch3 == Option_ModeKind__Complete || _switch3 == Option_ModeKind__Inquire {
                    if Lns_op_not(option.AnalyzeModule){
                        option.AnalyzeModule = arg
                    } else if Lns_op_not(lineNo){
                        lineNo = Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(arg, nil), 0)))
                    } else if Lns_op_not(column){
                        column = Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(arg, nil), 0)))
                        option.AnalyzePos = NewTypes_Position(_env, Lns_unwrap( lineNo).(LnsInt), Lns_unwrap( column).(LnsInt), Util_scriptPath2Module(_env, option.ScriptPath), nil)
                    }
                } else if _switch3 == Option_ModeKind__Save || _switch3 == Option_ModeKind__SaveMeta || _switch3 == Option_ModeKind__Glue {
                    option.OutputDir = arg
                } else if _switch3 == Option_ModeKind__Shebang {
                    if option.ShebangArgList.Len() == 0{
                        option.ShebangArgList.Insert(option.ScriptPath)
                    }
                    option.ShebangArgList.Insert(arg)
                } else {
                    option.OutputDir = arg
                }
            }
        }
        index = index + 1
    }
    if uptodateOpt != nil{
        uptodateOpt_319 := uptodateOpt.(string)
        if _switch4 := uptodateOpt_319; _switch4 == "force" {
            option.TransCtrlInfo.UptodateMode = &Types_CheckingUptodateMode__Force1{Util_scriptPath2Module(_env, option.ScriptPath)}
        } else if _switch4 == "forceAll" {
            option.TransCtrlInfo.UptodateMode = Types_CheckingUptodateMode__ForceAll_Obj
        } else if _switch4 == "normal" {
            option.TransCtrlInfo.UptodateMode = Types_CheckingUptodateMode__Normal_Obj
        } else if _switch4 == "touch" {
            option.TransCtrlInfo.UptodateMode = Types_CheckingUptodateMode__Touch_Obj
        } else {
            Util_errorLog(_env, "illegal mode -- " + uptodateOpt_319)
        }
    }
    if option.Mode != Option_ModeKind__Builtin{
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( option.ScriptPath == "") ||
            _env.SetStackVal( option.Mode == Option_ModeKind__Unknown) ).(bool){
            Option_analyze__printUsage_0_(_env, _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( argList.Len() == 0) ||
                    _env.SetStackVal( argList.GetAt(1) == "") ).(bool))) &&
                _env.SetStackVal( 0) ||
                _env.SetStackVal( 1) ).(LnsInt))
        }
    }
    if useStdInFlag{
        var code string
        
        {
            _code := Lns_io_stdin.Read(_env, "*a")
            if _code == nil{
                Util_err(_env, "read error from stdin.")
            } else {
                code = _code.(string)
            }
        }
        if Lns_isCondTrue( option.AnalyzeModule){
            option.stdinFile = NewTypes_StdinFile(_env, Lns_unwrap( option.AnalyzeModule).(string), code)
        } else { 
            if option.ScriptPath != ""{
                option.stdinFile = NewTypes_StdinFile(_env, Util_scriptPath2Module(_env, option.ScriptPath), code)
            }
        }
    }
    if option.ScriptPath == "@-"{
        for  {
            var line string
            
            {
                _line := Lns_io_stdin.Read(_env, "*l")
                if _line == nil{
                    break
                } else {
                    line = _line.(string)
                }
            }
            if len(line) > 0{
                option.BatchList.Insert(line)
            }
        }
    }
    Log_log(_env, Log_Level__Log, __func__, 779, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("mode is '%s'", Lns_2DDD(Option_ModeKind_getTxt( option.Mode)))
    }))
    
    return option
}

// 283: decl @lune.@base.@Option.analyze.printUsage
func Option_analyze__printUsage_0_(_env *LnsEnv, code LnsInt) {
    Util_println(_env, Lns_2DDD("usage:\n  <type1> [-prof] [-r] src.lns mode [mode-option]\n  <type2> -mklunemod path\n  <type3> -shebang path\n  <type4> --version\n\n* type1\n  - src.lns [common_op] ast\n  - src.lns [common_op] comp [-i] module line column\n  - src.lns [common_op] inq [-i] module line column\n  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] <lua|LUA>\n  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] [--depends dependfile] <save|SAVE> output-dir\n  - src.lns [common_op] exe\n\n  -r: use 'require( \"lune.base.runtime\" )'\n  -ol: output lua version. ver = 51 or 52 or 53.\n  -ob: output bytecompiled-code.\n      -ob0 is without debug information.\n      -ob1 is with debug information.\n  -langC: transcompile to c-lang.\n  -langGo: transcompile to golang.\n  -langPython: transcompile to python.\n  -noLua: no transcompile to lua.\n  -oc: output path of the source code transcompiled to c-lang .\n  --depends: output dependfile\n  --int2str mode: mode of int to str.\n     - depend: depends the lua version.\n     - need0: with '.0'.\n     - unneed0: without '.0'.\n\n  common_op:\n    --testing: enable test.\n    --projDir <dir>: set the project dir.\n    -u: update meta and lua on load.\n    -Werror: error by warrning.\n    --log <mode>: set log level.\n         mode: fatal, error, warn, log, info, debug, trace\n    --warning-shadowing: shadowing error convert to warning.\n    --compat-comment: backward compatibility to process the comment.\n    --disable-checking-define-abbr: disable checking for ##.\n    --uptodate <mode>: checking uptodate mode.\n            force: skip check for target lns file.\n            forceAll: skip check for all.\n            none: skip process when file is uptodate.\n            touch: touch meta file when file is uptodate.  (default)\n    --use-ipairs: use ipairs for foreach with List value.\n    --default-lazy: set lazy-loading at default.\n    --valid-luaval: enable luaval when transcompie to lua.\n    --package <name>: set the package name for the go-lang.\n    --app <name>: set the application name for the go-lang.\n    --debug-dump-ast: dump ast for debuging.\n\n    compati_op:\n      --legacyNewName: use the legacy new method name for lua.\n\n\n\n* type2\n  path: output file path.\n"))
    _env.GetVM().OS_exit(code)
}




// 784: decl @lune.@base.@Option.createDefaultOption
func Option_createDefaultOption(_env *LnsEnv, pathList *LnsList2_[string],projDir LnsAny) *Option_Option {
    var option *Option_Option
    option = NewOption_Option(_env)
    if pathList.Len() == 1{
        option.ScriptPath = pathList.GetAt(1)
    } else { 
        option.ScriptPath = "@-"
        for _, _path := range( pathList.Items ) {
            path := _path
            option.BatchList.Insert(path)
        }
    }
    option.UseLuneModule = Option_getRuntimeModule(_env)
    option.UseIpairs = true
    if projDir != nil{
        projDir_349 := projDir.(string)
        if projDir_349 != "/"{
            if Lns_op_not(Lns_car(_env.GetVM().String_find(projDir_349,"/$", nil, nil))){
                option.projDir = projDir_349 + "/"
            } else { 
                option.projDir = projDir_349
            }
        }
    }
    return option
}

// 55: decl @lune.@base.@Option.getBuildCount
func Option_getBuildCount_1_(_env *LnsEnv) LnsInt {
    return 13923
}


// 215: decl @lune.@base.@Option.Option.openDepend
func (self *Option_Option) OpenDepend(_env *LnsEnv, relPath LnsAny) LnsAny {
    {
        _path := self.DependsPath
        if !Lns_IsNil( _path ) {
            path := _path.(string)
            var filePath string
            if relPath != nil{
                relPath_138 := relPath.(string)
                if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(path,"/$", nil, nil))){
                    filePath = _env.GetVM().String_format("%s%s", Lns_2DDD(path, relPath_138))
                } else { 
                    filePath = _env.GetVM().String_format("%s/%s", Lns_2DDD(path, relPath_138))
                }
            } else {
                filePath = path
            }
            return Lns_car(Lns_io_open(filePath, "w"))
        }
    }
    {
        _path := self.DependsPath
        if !Lns_IsNil( _path ) {
            path := _path.(string)
            return Lns_car(Lns_io_open(path, "w"))
        }
    }
    return nil
}
// declaration Class -- RuntimeOpt
type Option_RuntimeOptMtd interface {
    Get_int2strMode(_env *LnsEnv) LnsInt
}
type Option_RuntimeOpt struct {
    int2strMode LnsInt
    FP Option_RuntimeOptMtd
}
func Option_RuntimeOpt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_RuntimeOpt).FP
}
func Option_RuntimeOpt_toSlice(slice []LnsAny) []*Option_RuntimeOpt {
    ret := make([]*Option_RuntimeOpt, len(slice))
    for index, val := range slice {
        ret[index] = val.(Option_RuntimeOptDownCast).ToOption_RuntimeOpt()
    }
    return ret
}
type Option_RuntimeOptDownCast interface {
    ToOption_RuntimeOpt() *Option_RuntimeOpt
}
func Option_RuntimeOptDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Option_RuntimeOptDownCast)
    if ok { return work.ToOption_RuntimeOpt() }
    return nil
}
func (obj *Option_RuntimeOpt) ToOption_RuntimeOpt() *Option_RuntimeOpt {
    return obj
}
func NewOption_RuntimeOpt(_env *LnsEnv) *Option_RuntimeOpt {
    obj := &Option_RuntimeOpt{}
    obj.FP = obj
    obj.InitOption_RuntimeOpt(_env)
    return obj
}
func (self *Option_RuntimeOpt) Get_int2strMode(_env *LnsEnv) LnsInt{ return self.int2strMode }
// 102: DeclConstr
func (self *Option_RuntimeOpt) InitOption_RuntimeOpt(_env *LnsEnv) {
    self.int2strMode = Option_Int2strMode__Int2strModeDepend
}


// declaration Class -- Option
type Option_OptionMtd interface {
    Get_addEnvArg(_env *LnsEnv) bool
    Get_enableRunner(_env *LnsEnv) bool
    Get_legacyNewName(_env *LnsEnv) bool
    Get_projDir(_env *LnsEnv) LnsAny
    Get_runtimeOpt(_env *LnsEnv) *Option_RuntimeOpt
    Get_sortGenerateCode(_env *LnsEnv) bool
    Get_stdinFile(_env *LnsEnv) LnsAny
    Get_validPostBuild(_env *LnsEnv) bool
    OpenDepend(_env *LnsEnv, arg1 LnsAny) LnsAny
    Set_stdinFile(_env *LnsEnv, arg1 LnsAny)
}
type Option_Option struct {
    Mode string
    AnalyzeModule LnsAny
    AnalyzePos LnsAny
    OutputDir LnsAny
    ScriptPath string
    BatchList *LnsList2_[string]
    ValidProf bool
    UseLuneModule LnsAny
    DependsPath LnsAny
    UpdateOnLoad bool
    ByteCompile bool
    StripDebugInfo bool
    TargetLuaVer *LuaVer_LuaVerInfo
    BootPath LnsAny
    UseIpairs bool
    PackageName LnsAny
    AppName LnsAny
    MainModule string
    DumpDebugAst bool
    TransCtrlInfo *Types_TransCtrlInfo
    ConvTo LnsAny
    NoLua bool
    Testing bool
    ShebangArgList *LnsList
    runtimeOpt *Option_RuntimeOpt
    projDir LnsAny
    addEnvArg bool
    enableRunner bool
    validPostBuild bool
    stdinFile LnsAny
    legacyNewName bool
    ConvGoRunnerNum LnsInt
    sortGenerateCode bool
    FP Option_OptionMtd
}
func Option_Option2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_Option).FP
}
func Option_Option_toSlice(slice []LnsAny) []*Option_Option {
    ret := make([]*Option_Option, len(slice))
    for index, val := range slice {
        ret[index] = val.(Option_OptionDownCast).ToOption_Option()
    }
    return ret
}
type Option_OptionDownCast interface {
    ToOption_Option() *Option_Option
}
func Option_OptionDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Option_OptionDownCast)
    if ok { return work.ToOption_Option() }
    return nil
}
func (obj *Option_Option) ToOption_Option() *Option_Option {
    return obj
}
func NewOption_Option(_env *LnsEnv) *Option_Option {
    obj := &Option_Option{}
    obj.FP = obj
    obj.InitOption_Option(_env)
    return obj
}
func (self *Option_Option) Get_runtimeOpt(_env *LnsEnv) *Option_RuntimeOpt{ return self.runtimeOpt }
func (self *Option_Option) Get_projDir(_env *LnsEnv) LnsAny{ return self.projDir }
func (self *Option_Option) Get_addEnvArg(_env *LnsEnv) bool{ return self.addEnvArg }
func (self *Option_Option) Get_enableRunner(_env *LnsEnv) bool{ return self.enableRunner }
func (self *Option_Option) Get_validPostBuild(_env *LnsEnv) bool{ return self.validPostBuild }
func (self *Option_Option) Get_stdinFile(_env *LnsEnv) LnsAny{ return self.stdinFile }
func (self *Option_Option) Set_stdinFile(_env *LnsEnv, arg1 LnsAny){ self.stdinFile = arg1 }
func (self *Option_Option) Get_legacyNewName(_env *LnsEnv) bool{ return self.legacyNewName }
func (self *Option_Option) Get_sortGenerateCode(_env *LnsEnv) bool{ return self.sortGenerateCode }
// 179: DeclConstr
func (self *Option_Option) InitOption_Option(_env *LnsEnv) {
    self.sortGenerateCode = true
    self.DumpDebugAst = false
    self.legacyNewName = false
    self.stdinFile = nil
    self.validPostBuild = true
    self.enableRunner = true
    self.addEnvArg = true
    self.projDir = nil
    self.runtimeOpt = NewOption_RuntimeOpt(_env)
    self.ShebangArgList = NewLnsList([]LnsAny{})
    self.MainModule = ""
    self.AppName = nil
    self.PackageName = nil
    self.Testing = false
    self.ConvTo = nil
    self.NoLua = false
    self.ValidProf = false
    self.Mode = Option_ModeKind__Unknown
    self.ScriptPath = ""
    self.BatchList = NewLnsList2_[string]([]string{})
    self.UseLuneModule = nil
    self.UpdateOnLoad = false
    self.ByteCompile = false
    self.StripDebugInfo = false
    self.TargetLuaVer = LuaVer_getCurVer(_env)
    self.TransCtrlInfo = Types_TransCtrlInfo_create_normal(_env)
    self.BootPath = nil
    self.UseIpairs = false
    self.AnalyzeModule = nil
    self.OutputDir = nil
    self.AnalyzePos = nil
    self.DependsPath = nil
    self.ConvGoRunnerNum = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( true) &&
        _env.SetStackVal( 4) ||
        _env.SetStackVal( 0) ).(LnsInt)
}


// declaration Class -- ProjInfo
type Option_ProjInfo1208Mtd interface {
    ToMap() *LnsMap
}
type Option_ProjInfo1208 struct {
    Cmd_option *LnsList2_[string]
    FP Option_ProjInfo1208Mtd
}
func Option_ProjInfo12082Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_ProjInfo1208).FP
}
func Option_ProjInfo1208_toSlice(slice []LnsAny) []*Option_ProjInfo1208 {
    ret := make([]*Option_ProjInfo1208, len(slice))
    for index, val := range slice {
        ret[index] = val.(Option_ProjInfo1208DownCast).ToOption_ProjInfo1208()
    }
    return ret
}
type Option_ProjInfo1208DownCast interface {
    ToOption_ProjInfo1208() *Option_ProjInfo1208
}
func Option_ProjInfo1208DownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Option_ProjInfo1208DownCast)
    if ok { return work.ToOption_ProjInfo1208() }
    return nil
}
func (obj *Option_ProjInfo1208) ToOption_ProjInfo1208() *Option_ProjInfo1208 {
    return obj
}
func NewOption_ProjInfo1208(_env *LnsEnv, arg1 *LnsList2_[string]) *Option_ProjInfo1208 {
    obj := &Option_ProjInfo1208{}
    obj.FP = obj
    obj.InitOption_ProjInfo1208(_env, arg1)
    return obj
}
func (self *Option_ProjInfo1208) InitOption_ProjInfo1208(_env *LnsEnv, arg1 *LnsList2_[string]) {
    self.Cmd_option = arg1
}
func (self *Option_ProjInfo1208) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["cmd_option"] = Lns_ToCollection( self.Cmd_option )
    return obj
}
func (self *Option_ProjInfo1208) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Option_ProjInfo1208__fromMap_2_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Option_ProjInfo1208_FromMap( arg1, paramList )
}
func Option_ProjInfo1208__fromStem_3_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Option_ProjInfo1208_FromMap( arg1, paramList )
}
func Option_ProjInfo1208_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Option_ProjInfo1208_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Option_ProjInfo1208_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Option_ProjInfo1208{}
    newObj.FP = newObj
    return Option_ProjInfo1208_FromMapMain( newObj, objMap, paramList )
}
func Option_ProjInfo1208_FromMapMain( newObj *Option_ProjInfo1208, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToList2Sub[string]( objMap.Items["cmd_option"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil}}); !ok {
       return false,nil,"cmd_option:" + mess.(string)
    } else {
       newObj.Cmd_option = conv.(*LnsList2_[string])
    }
    return true, newObj, nil
}

func Lns_Option_init(_env *LnsEnv) {
    if init_Option { return }
    init_Option = true
    Option__mod__ = "@lune.@base.@Option"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_Tokenizer_init(_env)
    Lns_AsyncTokenizer_init(_env)
    Lns_Json_init(_env)
    Lns_Util_init(_env)
    Lns_LuaMod_init(_env)
    Lns_Ver_init(_env)
    Lns_LuaVer_init(_env)
    Lns_LuaVer_init( _env )
    Lns_Depend_init(_env)
    Lns_Log_init(_env)
    Lns_Ast_init(_env)
    Lns_Builtin_init(_env)
    
}
func init() {
    init_Option = false
}
