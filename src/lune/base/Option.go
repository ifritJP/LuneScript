// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Option bool
var Option__mod__ string
// decl enum -- ModeKind 
type Option_ModeKind = string
const Option_ModeKind__Ast = "ast"
const Option_ModeKind__BootC = "bootC"
const Option_ModeKind__Builtin = "builtin"
const Option_ModeKind__Complete = "comp"
const Option_ModeKind__Diag = "diag"
const Option_ModeKind__Exec = "exe"
const Option_ModeKind__Format = "format"
const Option_ModeKind__Glue = "glue"
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
})
func Option_ModeKind_get__allList(_env *LnsEnv) *LnsList{
    return Option_ModeKindList_
}
var Option_ModeKindMap_ = map[string]string {
  Option_ModeKind__Ast: "ModeKind.Ast",
  Option_ModeKind__BootC: "ModeKind.BootC",
  Option_ModeKind__Builtin: "ModeKind.Builtin",
  Option_ModeKind__Complete: "ModeKind.Complete",
  Option_ModeKind__Diag: "ModeKind.Diag",
  Option_ModeKind__Exec: "ModeKind.Exec",
  Option_ModeKind__Format: "ModeKind.Format",
  Option_ModeKind__Glue: "ModeKind.Glue",
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
// for 41
func Option_convExp67(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 328
func Option_convExp864(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 323
func Option_convExp871(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 218
func Option_convExp664(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 51: decl @lune.@base.@Option.getBuildCount
func Option_getBuildCount_1024_(_env *LnsEnv) LnsInt {
    return 9056
}

// 78: decl @lune.@base.@Option.getRuntimeModule
func Option_getRuntimeModule(_env *LnsEnv) string {
    return _env.LuaVM.String_format("lune.base.runtime%d", []LnsAny{Ver_luaModVersion})
}

// 211: decl @lune.@base.@Option.outputLuneMod
func Option_outputLuneMod(_env *LnsEnv, path LnsAny) LnsAny {
    var lune_path string
    lune_path = "runtime.lua"
    if path != nil{
        path_125 := path.(string)
        if path_125 != ""{
            lune_path = path_125
            
        }
    }
    var fileObj Lns_luaStream
    
    {
        _fileObj := Option_convExp664(Lns_2DDD(Lns_io_open(lune_path, "w")))
        if _fileObj == nil{
            return _env.LuaVM.String_format("failed to open -- %s", []LnsAny{lune_path})
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

// 258: decl @lune.@base.@Option.analyze.printUsage
func analyze__printUsage_1176_(_env *LnsEnv, code LnsInt) {
    Lns_print([]LnsAny{"usage:\n  <type1> [-prof] [-r] src.lns mode [mode-option]\n  <type2> -mklunemod path\n  <type3> -shebang path\n  <type4> --version\n\n* type1\n  - src.lns [common_op] ast\n  - src.lns [common_op] comp [-i] module line column\n  - src.lns [common_op] inq [-i] module line column\n  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] <lua|LUA>\n  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] [--depends dependfile] <save|SAVE> output-dir\n  - src.lns [common_op] exe\n\n  -r: use 'require( \"lune.base.runtime\" )'\n  -ol: output lua version. ver = 51 or 52 or 53.\n  -ob: output bytecompiled-code.\n      -ob0 is without debug information.\n      -ob1 is with debug information.\n  -langC: transcompile to c-lang.\n  -langGo: transcompile to golang.\n  -oc: output path of the source code transcompiled to c-lang .\n  --depends: output dependfile\n  --int2str mode: mode of int to str.\n     - depend: depends the lua version.\n     - need0: with '.0'.\n     - unneed0: without '.0'.\n\n  common_op:\n    --testing: enable test.\n    --projDir <dir>: set the project dir.\n    -u: update meta and lua on load.\n    -Werror: error by warrning.\n    --log <mode>: set log level.\n         mode: fatal, error, warn, log, info, debug, trace\n    --warning-shadowing: shadowing error convert to warning.\n    --compat-comment: backward compatibility to process the comment.\n    --disable-checking-define-abbr: disable checking for ##.\n    --uptodate <mode>: checking uptodate mode.\n            force: skip check for target lns file.\n            forceAll: skip check for all.\n            none: skip process when file is uptodate.\n            touch: touch meta file when file is uptodate.  (default)\n    --use-ipairs: use ipairs for foreach with List value.\n    --default-lazy: set lazy-loading at default.\n    --valid-luaval: enable luaval when transcompie to lua.\n    --package <name>: set the package name for the go-lang.\n    --app <name>: set the application name for the go-lang.\n\n* type2\n  path: output file path.\n"})
    _env.LuaVM.OS_exit(code)
}



// 256: decl @lune.@base.@Option.analyze
func Option_analyze(_env *LnsEnv, argList *LnsList) *Option_Option {
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
        _file := Option_convExp871(Lns_2DDD(Lns_io_open("lune.js", "r")))
        if !Lns_IsNil( _file ) {
            file := _file.(Lns_luaStream)
            {
                _projInfo := Option_convExp864(Lns_2DDD(Option_ProjInfo1188__fromStem_1208_(_env, Lns_car(Json_fromStr(_env, _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( file.Read(_env, "*a")) ||
                    _env.SetStackVal( "") ).(string))),nil)))
                if !Lns_IsNil( _projInfo ) {
                    projInfo := _projInfo.(*Option_ProjInfo1188)
                    var workArgList *LnsList
                    workArgList = NewLnsList([]LnsAny{})
                    for _, _arg := range( projInfo.Cmd_option.Items ) {
                        arg := _arg.(string)
                        workArgList.Insert(arg)
                    }
                    for _, _arg := range( argList.Items ) {
                        arg := _arg.(string)
                        workArgList.Insert(arg)
                    }
                    argList = workArgList
                    
                }
            }
            file.Close(_env)
        }
    }
    var getNextOp func(_env *LnsEnv) LnsAny
    getNextOp = func(_env *LnsEnv) LnsAny {
        if argList.Len() <= index{
            return nil
        }
        index = index + 1
        
        return argList.GetAt(index).(string)
    }
    Util_setDebugFlag(_env, false)
    var uptodateOpt LnsAny
    uptodateOpt = nil
    for argList.Len() >= index {
        var arg string
        arg = argList.GetAt(index).(string)
        if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(arg,"^-", nil, nil))){
            if option.Mode != Option_ModeKind__Shebang{
                if _switch1757 := (arg); _switch1757 == "-i" {
                    useStdInFlag = true
                    
                } else if _switch1757 == "-prof" {
                    option.ValidProf = true
                    
                } else if _switch1757 == "--noEnvArg" {
                    option.addEnvArg = false
                    
                } else if _switch1757 == "--disableRunner" {
                    option.enableRunner = false
                    
                } else if _switch1757 == "--disablePostBuild" {
                    option.validPostBuild = false
                    
                } else if _switch1757 == "--enableAsyncCtl" {
                    option.TransCtrlInfo.ValidAsyncCtrl = true
                    
                } else if _switch1757 == "--nodebug" {
                    Util_setDebugFlag(_env, false)
                } else if _switch1757 == "--debug" {
                    Util_setDebugFlag(_env, true)
                } else if _switch1757 == "-shebang" {
                    option.Mode = Option_ModeKind__Shebang
                    
                } else if _switch1757 == "--version" {
                    Lns_print([]LnsAny{_env.LuaVM.String_format("LuneScript: version %s (%d:Lua%s) [%s]", []LnsAny{Ver_version, Option_getBuildCount_1024_(_env), Depend_getLuaVersion(_env), Ver_metaVersion})})
                    _env.LuaVM.OS_exit(0)
                } else if _switch1757 == "--projDir" {
                    option.projDir = getNextOp(_env)
                    
                } else if _switch1757 == "--builtin" {
                    {
                        __collection1144 := Ast_getBuiltInTypeIdMap(_env)
                        __sorted1144 := __collection1144.CreateKeyListInt()
                        __sorted1144.Sort( LnsItemKindInt, nil )
                        for _, _typeId := range( __sorted1144.Items ) {
                            builtinTypeInfo := __collection1144.Items[ _typeId ].(Ast_BuiltinTypeInfoDownCast).ToAst_BuiltinTypeInfo()
                            typeId := _typeId.(LnsInt)
                            Lns_print([]LnsAny{typeId, builtinTypeInfo.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil)})
                        }
                    }
                    _env.LuaVM.OS_exit(0)
                } else if _switch1757 == "-mklunemod" {
                    var path LnsAny
                    path = getNextOp(_env)
                    {
                        _mess := Option_outputLuneMod(_env, path)
                        if !Lns_IsNil( _mess ) {
                            mess := _mess.(string)
                            Util_errorLog(_env, mess)
                            _env.LuaVM.OS_exit(1)
                        }
                    }
                    _env.LuaVM.OS_exit(0)
                } else if _switch1757 == "--mkbuiltin" {
                    var path string
                    
                    {
                        _path := getNextOp(_env)
                        if _path == nil{
                            path = "."
                            
                        } else {
                            path = _path.(string)
                        }
                    }
                    option.ScriptPath = path + "/lns_builtin.lns"
                    
                    option.Mode = Option_ModeKind__Builtin
                    
                } else if _switch1757 == "-r" {
                    option.UseLuneModule = Option_getRuntimeModule(_env)
                    
                } else if _switch1757 == "--runtime" {
                    option.UseLuneModule = getNextOp(_env)
                    
                } else if _switch1757 == "-oc" {
                    option.BootPath = getNextOp(_env)
                    
                } else if _switch1757 == "-u" {
                    option.UpdateOnLoad = true
                    
                } else if _switch1757 == "-Werror" {
                    option.TransCtrlInfo.StopByWarning = true
                    
                } else if _switch1757 == "--disable-checking-define-abbr" {
                    option.TransCtrlInfo.CheckingDefineAbbr = false
                    
                } else if _switch1757 == "--disable-checking-mutable" {
                    option.TransCtrlInfo.ValidCheckingMutable = false
                    
                } else if _switch1757 == "--legacy-mutable-control" {
                    option.TransCtrlInfo.LegacyMutableControl = true
                    
                } else if _switch1757 == "--compat-comment" {
                    option.TransCtrlInfo.CompatComment = true
                    
                } else if _switch1757 == "--warning-shadowing" {
                    option.TransCtrlInfo.WarningShadowing = true
                    
                } else if _switch1757 == "--valid-luaval" {
                    option.TransCtrlInfo.ValidLuaval = true
                    
                } else if _switch1757 == "--default-lazy" {
                    option.TransCtrlInfo.DefaultLazy = true
                    
                } else if _switch1757 == "--package" {
                    option.PackageName = getNextOp(_env)
                    
                } else if _switch1757 == "--int2str" {
                    var opt LnsAny
                    opt = getNextOp(_env)
                    if _switch1472 := opt; _switch1472 == "depend" {
                        option.FP.Get_runtimeOpt(_env).int2strMode = Option_Int2strMode__Int2strModeDepend
                        
                    } else if _switch1472 == "need0" {
                        option.FP.Get_runtimeOpt(_env).int2strMode = Option_Int2strMode__Int2strModeNeed0
                        
                    } else if _switch1472 == "unneed0" {
                        option.FP.Get_runtimeOpt(_env).int2strMode = Option_Int2strMode__Int2strModeUnneed0
                        
                    } else {
                        Util_errorLog(_env, _env.LuaVM.String_format("unknown mode -- %s", []LnsAny{opt}))
                        _env.LuaVM.OS_exit(1)
                    }
                } else if _switch1757 == "--app" {
                    {
                        __exp := getNextOp(_env)
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(string)
                            option.AppName = _exp
                            
                        }
                    }
                } else if _switch1757 == "--main" {
                    {
                        __exp := getNextOp(_env)
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(string)
                            option.MainModule = _exp
                            
                        }
                    }
                } else if _switch1757 == "--log" {
                    {
                        _txt := getNextOp(_env)
                        if !Lns_IsNil( _txt ) {
                            txt := _txt.(string)
                            {
                                _level := Log_str2level(_env, txt)
                                if !Lns_IsNil( _level ) {
                                    level := _level.(LnsInt)
                                    Log_setLevel(_env, level)
                                } else {
                                    Util_errorLog(_env, _env.LuaVM.String_format("illegal level -- %s", []LnsAny{txt}))
                                }
                            }
                        }
                    }
                } else if _switch1757 == "--testing" {
                    option.Testing = true
                    
                } else if _switch1757 == "--depends" {
                    option.DependsPath = getNextOp(_env)
                    
                } else if _switch1757 == "--use-ipairs" {
                    option.UseIpairs = true
                    
                } else if _switch1757 == "--uptodate" {
                    uptodateOpt = getNextOp(_env)
                    
                } else if _switch1757 == "-langC" {
                    option.ConvTo = Types_Lang__C
                    
                    option.TransCtrlInfo.ValidLuaval = true
                    
                } else if _switch1757 == "-langGo" {
                    option.ConvTo = Types_Lang__Go
                    
                    option.TransCtrlInfo.ValidLuaval = true
                    
                    option.TransCtrlInfo.ValidAsyncCtrl = true
                    
                } else if _switch1757 == "-ol" {
                    {
                        _txt := getNextOp(_env)
                        if !Lns_IsNil( _txt ) {
                            txt := _txt.(string)
                            if _switch1701 := txt; _switch1701 == "51" {
                                option.TargetLuaVer = LuaVer_ver51
                                
                            } else if _switch1701 == "52" {
                                option.TargetLuaVer = LuaVer_ver52
                                
                            } else if _switch1701 == "53" {
                                option.TargetLuaVer = LuaVer_ver53
                                
                            }
                        }
                    }
                } else if _switch1757 == "-ob0" || _switch1757 == "-ob1" {
                    option.ByteCompile = true
                    
                    if arg == "-ob0"{
                        option.StripDebugInfo = true
                        
                    }
                } else {
                    Util_log(_env, _env.LuaVM.String_format("unknown option -- '%s'", []LnsAny{arg}))
                    _env.LuaVM.OS_exit(1)
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
                        Util_err(_env, _env.LuaVM.String_format("unknown mode -- %s", []LnsAny{arg}))
                    }
                }
            } else { 
                if _switch2042 := (option.Mode); _switch2042 == Option_ModeKind__Complete || _switch2042 == Option_ModeKind__Inquire {
                    if Lns_op_not(option.AnalyzeModule){
                        option.AnalyzeModule = arg
                        
                    } else if Lns_op_not(lineNo){
                        lineNo = Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(arg, nil), 0)))
                        
                    } else if Lns_op_not(column){
                        column = Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(arg, nil), 0)))
                        
                        option.AnalyzePos = NewTypes_Position(_env, Lns_unwrap( lineNo).(LnsInt), Lns_unwrap( column).(LnsInt), Util_scriptPath2Module(_env, option.ScriptPath))
                        
                    }
                } else if _switch2042 == Option_ModeKind__Save || _switch2042 == Option_ModeKind__SaveMeta || _switch2042 == Option_ModeKind__Glue {
                    option.OutputDir = arg
                    
                } else if _switch2042 == Option_ModeKind__Shebang {
                    if option.ShebangArgList.Len() == 0{
                        option.ShebangArgList.Insert(option.ScriptPath)
                    }
                    option.ShebangArgList.Insert(arg)
                } else {
                    option.OutputPath = arg
                    
                }
            }
        }
        index = index + 1
        
    }
    if uptodateOpt != nil{
        uptodateOpt_257 := uptodateOpt.(string)
        if _switch2139 := uptodateOpt_257; _switch2139 == "force" {
            option.TransCtrlInfo.UptodateMode = &Types_CheckingUptodateMode__Force1{Util_scriptPath2Module(_env, option.ScriptPath)}
            
        } else if _switch2139 == "forceAll" {
            option.TransCtrlInfo.UptodateMode = Types_CheckingUptodateMode__ForceAll_Obj
            
        } else if _switch2139 == "normal" {
            option.TransCtrlInfo.UptodateMode = Types_CheckingUptodateMode__Normal_Obj
            
        } else if _switch2139 == "touch" {
            option.TransCtrlInfo.UptodateMode = Types_CheckingUptodateMode__Touch_Obj
            
        } else {
            Util_errorLog(_env, "illegal mode -- " + uptodateOpt_257)
        }
    }
    if option.Mode != Option_ModeKind__Builtin{
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( option.ScriptPath == "") ||
            _env.SetStackVal( option.Mode == Option_ModeKind__Unknown) ).(bool){
            analyze__printUsage_1176_(_env, _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( argList.Len() == 0) ||
                    _env.SetStackVal( argList.GetAt(1).(string) == "") ).(bool))) &&
                _env.SetStackVal( 0) ||
                _env.SetStackVal( 1) ).(LnsInt))
        }
    }
    if useStdInFlag{
        if Lns_isCondTrue( option.AnalyzeModule){
            Parser_StreamParser_setStdinStream(_env, Lns_unwrap( option.AnalyzeModule).(string))
        } else { 
            if option.ScriptPath != ""{
                Parser_StreamParser_setStdinStream(_env, Util_scriptPath2Module(_env, option.ScriptPath))
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
    Log_log(_env, Log_Level__Log, __func__, 649, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.LuaVM.String_format("mode is '%s'", []LnsAny{Option_ModeKind_getTxt( option.Mode)})
    }))
    
    return option
}

// 654: decl @lune.@base.@Option.createDefaultOption
func Option_createDefaultOption(_env *LnsEnv, pathList *LnsList,projDir LnsAny) *Option_Option {
    var option *Option_Option
    option = NewOption_Option(_env)
    if pathList.Len() == 1{
        option.ScriptPath = pathList.GetAt(1).(string)
        
    } else { 
        option.ScriptPath = "@-"
        
        for _, _path := range( pathList.Items ) {
            path := _path.(string)
            option.BatchList.Insert(path)
        }
    }
    option.UseLuneModule = Option_getRuntimeModule(_env)
    
    option.UseIpairs = true
    
    if projDir != nil{
        projDir_288 := projDir.(string)
        if projDir_288 != "/"{
            if Lns_op_not(Lns_car(_env.LuaVM.String_find(projDir_288,"/$", nil, nil))){
                option.projDir = projDir_288 + "/"
                
            } else { 
                option.projDir = projDir_288
                
            }
        }
    }
    return option
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
// 95: DeclConstr
func (self *Option_RuntimeOpt) InitOption_RuntimeOpt(_env *LnsEnv) {
    self.int2strMode = Option_Int2strMode__Int2strModeDepend
    
}


// declaration Class -- Option
type Option_OptionMtd interface {
    Get_addEnvArg(_env *LnsEnv) bool
    Get_enableRunner(_env *LnsEnv) bool
    Get_projDir(_env *LnsEnv) LnsAny
    Get_runtimeOpt(_env *LnsEnv) *Option_RuntimeOpt
    Get_validPostBuild(_env *LnsEnv) bool
    OpenDepend(_env *LnsEnv, arg1 LnsAny) LnsAny
}
type Option_Option struct {
    Mode string
    AnalyzeModule LnsAny
    AnalyzePos LnsAny
    OutputDir LnsAny
    ScriptPath string
    BatchList *LnsList
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
    TransCtrlInfo *Types_TransCtrlInfo
    ConvTo LnsAny
    Testing bool
    OutputPath LnsAny
    ShebangArgList *LnsList
    runtimeOpt *Option_RuntimeOpt
    projDir LnsAny
    addEnvArg bool
    enableRunner bool
    validPostBuild bool
    FP Option_OptionMtd
}
func Option_Option2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_Option).FP
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
// 159: DeclConstr
func (self *Option_Option) InitOption_Option(_env *LnsEnv) {
    self.validPostBuild = true
    
    self.enableRunner = true
    
    self.addEnvArg = true
    
    self.projDir = nil
    
    self.runtimeOpt = NewOption_RuntimeOpt(_env)
    
    self.ShebangArgList = NewLnsList([]LnsAny{})
    
    self.OutputPath = nil
    
    self.MainModule = ""
    
    self.AppName = nil
    
    self.PackageName = nil
    
    self.Testing = false
    
    self.ConvTo = nil
    
    self.ValidProf = false
    
    self.Mode = Option_ModeKind__Unknown
    
    self.ScriptPath = ""
    
    self.BatchList = NewLnsList([]LnsAny{})
    
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
    
}

// 190: decl @lune.@base.@Option.Option.openDepend
func (self *Option_Option) OpenDepend(_env *LnsEnv, relPath LnsAny) LnsAny {
    {
        _path := self.DependsPath
        if !Lns_IsNil( _path ) {
            path := _path.(string)
            var filePath string
            if relPath != nil{
                relPath_113 := relPath.(string)
                if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(path,"/$", nil, nil))){
                    filePath = _env.LuaVM.String_format("%s%s", []LnsAny{path, relPath_113})
                    
                } else { 
                    filePath = _env.LuaVM.String_format("%s/%s", []LnsAny{path, relPath_113})
                    
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


// declaration Class -- ProjInfo
type Option_ProjInfo1188Mtd interface {
    ToMap() *LnsMap
}
type Option_ProjInfo1188 struct {
    Cmd_option *LnsList
    FP Option_ProjInfo1188Mtd
}
func Option_ProjInfo11882Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_ProjInfo1188).FP
}
type Option_ProjInfo1188DownCast interface {
    ToOption_ProjInfo1188() *Option_ProjInfo1188
}
func Option_ProjInfo1188DownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Option_ProjInfo1188DownCast)
    if ok { return work.ToOption_ProjInfo1188() }
    return nil
}
func (obj *Option_ProjInfo1188) ToOption_ProjInfo1188() *Option_ProjInfo1188 {
    return obj
}
func NewOption_ProjInfo1188(_env *LnsEnv, arg1 *LnsList) *Option_ProjInfo1188 {
    obj := &Option_ProjInfo1188{}
    obj.FP = obj
    obj.InitOption_ProjInfo1188(_env, arg1)
    return obj
}
func (self *Option_ProjInfo1188) InitOption_ProjInfo1188(_env *LnsEnv, arg1 *LnsList) {
    self.Cmd_option = arg1
}
func (self *Option_ProjInfo1188) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["cmd_option"] = Lns_ToCollection( self.Cmd_option )
    return obj
}
func (self *Option_ProjInfo1188) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Option_ProjInfo1188__fromMap_1204_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Option_ProjInfo1188_FromMap( arg1, paramList )
}
func Option_ProjInfo1188__fromStem_1208_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Option_ProjInfo1188_FromMap( arg1, paramList )
}
func Option_ProjInfo1188_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Option_ProjInfo1188_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Option_ProjInfo1188_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Option_ProjInfo1188{}
    newObj.FP = newObj
    return Option_ProjInfo1188_FromMapMain( newObj, objMap, paramList )
}
func Option_ProjInfo1188_FromMapMain( newObj *Option_ProjInfo1188, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToListSub( objMap.Items["cmd_option"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil}}); !ok {
       return false,nil,"cmd_option:" + mess.(string)
    } else {
       newObj.Cmd_option = conv.(*LnsList)
    }
    return true, newObj, nil
}

func Lns_Option_init(_env *LnsEnv) {
    if init_Option { return }
    init_Option = true
    Option__mod__ = "@lune.@base.@Option"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_Parser_init(_env)
    Lns_Json_init(_env)
    Lns_Util_init(_env)
    Lns_LuaMod_init(_env)
    Lns_Ver_init(_env)
    Lns_LuaVer_init(_env)
    Lns_LuaVer_init( _env )
    Lns_Depend_init(_env)
    Lns_Log_init(_env)
    Lns_Ast_init(_env)
    
}
func init() {
    init_Option = false
}
