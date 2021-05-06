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
func Option_ModeKind_get__allList() *LnsList{
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
func Option_ModeKind__from(arg1 string) LnsAny{
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
func Option_Int2strMode_get__allList() *LnsList{
    return Option_Int2strModeList_
}
var Option_Int2strModeMap_ = map[LnsInt]string {
  Option_Int2strMode__Int2strModeDepend: "Int2strMode.Int2strModeDepend",
  Option_Int2strMode__Int2strModeNeed0: "Int2strMode.Int2strModeNeed0",
  Option_Int2strMode__Int2strModeUnneed0: "Int2strMode.Int2strModeUnneed0",
}
func Option_Int2strMode__from(arg1 LnsInt) LnsAny{
    if _, ok := Option_Int2strModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Option_Int2strMode_getTxt(arg1 LnsInt) string {
    return Option_Int2strModeMap_[arg1];
}
// for 39
func Option_convExp60(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 317
func Option_convExp825(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 312
func Option_convExp827(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 207
func Option_convExp625(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 48: decl @lune.@base.@Option.getBuildCount
func Option_getBuildCount_1009_() LnsInt {
    return 7883
}

// 75: decl @lune.@base.@Option.getRuntimeModule
func Option_getRuntimeModule() string {
    return Lns_getVM().String_format("lune.base.runtime%d", []LnsAny{Ver_luaModVersion})
}

// 200: decl @lune.@base.@Option.outputLuneMod
func Option_outputLuneMod(path LnsAny) LnsAny {
    var lune_path string
    lune_path = "runtime.lua"
    if path != nil{
        path_1601 := path.(string)
        if path_1601 != ""{
            lune_path = path_1601
            
        }
    }
    var fileObj Lns_luaStream
    
    {
        _fileObj := Option_convExp625(Lns_2DDD(Lns_io_open(lune_path, "w")))
        if _fileObj == nil{
            return Lns_getVM().String_format("failed to open -- %s", []LnsAny{lune_path})
        } else {
            fileObj = _fileObj.(Lns_luaStream)
        }
    }
    fileObj.Write("--[[\nMIT License\n\nCopyright (c) 2018,2019 ifritJP\n\nPermission is hereby granted, free of charge, to any person obtaining a copy\nof this software and associated documentation files (the \"Software\"), to deal\nin the Software without restriction, including without limitation the rights\nto use, copy, modify, merge, publish, distribute, sublicense, and/or sell\ncopies of the Software, and to permit persons to whom the Software is\nfurnished to do so, subject to the following conditions:\n\nThe above copyright notice and this permission notice shall be included in all\ncopies or substantial portions of the Software.\n\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\nIMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\nFITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\nAUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\nLIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\nOUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE\nSOFTWARE.\n]]\n")
    for _, _kind := range( LuaMod_CodeKind_get__allList().Items ) {
        kind := _kind.(LnsInt)
        fileObj.Write(LuaMod_getCode(kind))
    }
    fileObj.Close()
    return nil
}

// 247: decl @lune.@base.@Option.analyze.printUsage
func analyze__printUsage_1073_(code LnsInt) {
    Lns_print([]LnsAny{"usage:\n  <type1> [-prof] [-r] src.lns mode [mode-option]\n  <type2> -mklunemod path\n  <type3> -shebang path\n  <type4> --version\n\n* type1\n  - src.lns [common_op] ast\n  - src.lns [common_op] comp [-i] module line column\n  - src.lns [common_op] inq [-i] module line column\n  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] <lua|LUA>\n  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] [--depends dependfile] <save|SAVE> output-dir\n  - src.lns [common_op] exe\n\n  -r: use 'require( \"lune.base.runtime\" )'\n  -ol: output lua version. ver = 51 or 52 or 53.\n  -ob: output bytecompiled-code.\n      -ob0 is without debug information.\n      -ob1 is with debug information.\n  -langC: transcompile to c-lang.\n  -langGo: transcompile to golang.\n  -oc: output path of the source code transcompiled to c-lang .\n  --depends: output dependfile\n  --int2str mode: mode of int to str.\n     - depend: depends the lua version.\n     - need0: with '.0'.\n     - unneed0: without '.0'.\n\n  common_op:\n    --testing: enable test.\n    --projDir <dir>: set the project dir.\n    -u: update meta and lua on load.\n    -Werror: error by warrning.\n    --log <mode>: set log level.\n         mode: fatal, error, warn, log, info, debug, trace\n    --warning-shadowing: shadowing error convert to warning.\n    --compat-comment: backward compatibility to process the comment.\n    --disable-checking-define-abbr: disable checking for ##.\n    --uptodate <mode>: checking uptodate mode.\n            force: skip check for target lns file.\n            forceAll: skip check for all.\n            none: skip process when file is uptodate.\n            touch: touch meta file when file is uptodate.  (default)\n    --use-ipairs: use ipairs for foreach with List value.\n    --default-lazy: set lazy-loading at default.\n    --valid-luaval: enable luaval when transcompie to lua.\n    --package <name>: set the package name for the go-lang.\n    --app <name>: set the application name for the go-lang.\n\n* type2\n  path: output file path.\n"})
    Lns_getVM().OS_exit(code)
}



// 245: decl @lune.@base.@Option.analyze
func Option_analyze(argList *LnsList) *Option_Option {
    __func__ := "@lune.@base.@Option.analyze"
    var option *Option_Option
    option = NewOption_Option()
    var useStdInFlag bool
    useStdInFlag = false
    var lineNo LnsAny
    lineNo = nil
    var column LnsAny
    column = nil
    var index LnsInt
    index = 1
    {
        _file := Option_convExp827(Lns_2DDD(Lns_io_open("lune.js", "r")))
        if !Lns_IsNil( _file ) {
            file := _file.(Lns_luaStream)
            {
                _projInfo := Option_convExp825(Lns_2DDD(Option_ProjInfo1075__fromStem_1090_(Lns_car(Json_fromStr(Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( file.Read("*a")) ||
                    Lns_GetEnv().SetStackVal( "") ).(string))),nil)))
                if !Lns_IsNil( _projInfo ) {
                    projInfo := _projInfo.(*Option_ProjInfo1075)
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
        }
    }
    var getNextOp func() LnsAny
    getNextOp = func() LnsAny {
        if argList.Len() <= index{
            return nil
        }
        index = index + 1
        
        return argList.GetAt(index).(string)
    }
    Util_setDebugFlag(false)
    var uptodateOpt LnsAny
    uptodateOpt = nil
    for argList.Len() >= index {
        var arg string
        arg = argList.GetAt(index).(string)
        if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(arg,"^-", nil, nil))){
            if option.Mode != Option_ModeKind__Shebang{
                if _switch1654 := (arg); _switch1654 == "-i" {
                    useStdInFlag = true
                    
                } else if _switch1654 == "-prof" {
                    option.ValidProf = true
                    
                } else if _switch1654 == "--nodebug" {
                    Util_setDebugFlag(false)
                } else if _switch1654 == "--debug" {
                    Util_setDebugFlag(true)
                } else if _switch1654 == "-shebang" {
                    option.Mode = Option_ModeKind__Shebang
                    
                } else if _switch1654 == "--version" {
                    Lns_print([]LnsAny{Lns_getVM().String_format("LuneScript: version %s (%d:Lua%s) [%s]", []LnsAny{Ver_version, Option_getBuildCount_1009_(), Depend_getLuaVersion(), Ver_metaVersion})})
                    Lns_getVM().OS_exit(0)
                } else if _switch1654 == "--projDir" {
                    option.projDir = getNextOp()
                    
                } else if _switch1654 == "--builtin" {
                    {
                        __collection1050 := Ast_getBuiltInTypeIdMap()
                        __sorted1050 := __collection1050.CreateKeyListInt()
                        __sorted1050.Sort( LnsItemKindInt, nil )
                        for _, _typeId := range( __sorted1050.Items ) {
                            typeInfo := __collection1050.Items[ _typeId ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                            typeId := _typeId.(LnsInt)
                            Lns_print([]LnsAny{typeId, typeInfo.FP.GetTxt(nil, nil, nil)})
                        }
                    }
                    Lns_getVM().OS_exit(0)
                } else if _switch1654 == "-mklunemod" {
                    var path LnsAny
                    path = getNextOp()
                    {
                        _mess := Option_outputLuneMod(path)
                        if !Lns_IsNil( _mess ) {
                            mess := _mess.(string)
                            Util_errorLog(mess)
                            Lns_getVM().OS_exit(1)
                        }
                    }
                    Lns_getVM().OS_exit(0)
                } else if _switch1654 == "--mkbuiltin" {
                    var path string
                    
                    {
                        _path := getNextOp()
                        if _path == nil{
                            path = "."
                            
                        } else {
                            path = _path.(string)
                        }
                    }
                    option.ScriptPath = path + "/lns_builtin.lns"
                    
                    option.Mode = Option_ModeKind__Builtin
                    
                } else if _switch1654 == "-r" {
                    option.UseLuneModule = Option_getRuntimeModule()
                    
                } else if _switch1654 == "--runtime" {
                    option.UseLuneModule = getNextOp()
                    
                } else if _switch1654 == "-oc" {
                    option.BootPath = getNextOp()
                    
                } else if _switch1654 == "-u" {
                    option.UpdateOnLoad = true
                    
                } else if _switch1654 == "-Werror" {
                    option.TransCtrlInfo.StopByWarning = true
                    
                } else if _switch1654 == "--disable-checking-define-abbr" {
                    option.TransCtrlInfo.CheckingDefineAbbr = false
                    
                } else if _switch1654 == "--disable-checking-mutable" {
                    option.TransCtrlInfo.ValidCheckingMutable = false
                    
                } else if _switch1654 == "--legacy-mutable-control" {
                    option.TransCtrlInfo.LegacyMutableControl = true
                    
                } else if _switch1654 == "--compat-comment" {
                    option.TransCtrlInfo.CompatComment = true
                    
                } else if _switch1654 == "--warning-shadowing" {
                    option.TransCtrlInfo.WarningShadowing = true
                    
                } else if _switch1654 == "--valid-luaval" {
                    option.TransCtrlInfo.ValidLuaval = true
                    
                } else if _switch1654 == "--default-lazy" {
                    option.TransCtrlInfo.DefaultLazy = true
                    
                } else if _switch1654 == "--package" {
                    option.PackageName = getNextOp()
                    
                } else if _switch1654 == "--int2str" {
                    var opt LnsAny
                    opt = getNextOp()
                    if _switch1378 := opt; _switch1378 == "depend" {
                        option.FP.Get_runtimeOpt().int2strMode = Option_Int2strMode__Int2strModeDepend
                        
                    } else if _switch1378 == "need0" {
                        option.FP.Get_runtimeOpt().int2strMode = Option_Int2strMode__Int2strModeNeed0
                        
                    } else if _switch1378 == "unneed0" {
                        option.FP.Get_runtimeOpt().int2strMode = Option_Int2strMode__Int2strModeUnneed0
                        
                    } else {
                        Util_errorLog(Lns_getVM().String_format("unknown mode -- %s", []LnsAny{opt}))
                        Lns_getVM().OS_exit(1)
                    }
                } else if _switch1654 == "--app" {
                    {
                        __exp := getNextOp()
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(string)
                            option.AppName = _exp
                            
                        }
                    }
                } else if _switch1654 == "--main" {
                    {
                        __exp := getNextOp()
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(string)
                            option.MainModule = _exp
                            
                        }
                    }
                } else if _switch1654 == "--log" {
                    {
                        _txt := getNextOp()
                        if !Lns_IsNil( _txt ) {
                            txt := _txt.(string)
                            {
                                _level := Log_str2level(txt)
                                if !Lns_IsNil( _level ) {
                                    level := _level.(LnsInt)
                                    Log_setLevel(level)
                                } else {
                                    Util_errorLog(Lns_getVM().String_format("illegal level -- %s", []LnsAny{txt}))
                                }
                            }
                        }
                    }
                } else if _switch1654 == "--testing" {
                    option.Testing = true
                    
                } else if _switch1654 == "--depends" {
                    option.DependsPath = getNextOp()
                    
                } else if _switch1654 == "--use-ipairs" {
                    option.UseIpairs = true
                    
                } else if _switch1654 == "--uptodate" {
                    uptodateOpt = getNextOp()
                    
                } else if _switch1654 == "-langC" {
                    option.ConvTo = Types_Lang__C
                    
                    option.TransCtrlInfo.ValidLuaval = true
                    
                } else if _switch1654 == "-langGo" {
                    option.ConvTo = Types_Lang__Go
                    
                    option.TransCtrlInfo.ValidLuaval = true
                    
                } else if _switch1654 == "-ol" {
                    {
                        _txt := getNextOp()
                        if !Lns_IsNil( _txt ) {
                            txt := _txt.(string)
                            if _switch1598 := txt; _switch1598 == "51" {
                                option.TargetLuaVer = LuaVer_ver51
                                
                            } else if _switch1598 == "52" {
                                option.TargetLuaVer = LuaVer_ver52
                                
                            } else if _switch1598 == "53" {
                                option.TargetLuaVer = LuaVer_ver53
                                
                            }
                        }
                    }
                } else if _switch1654 == "-ob0" || _switch1654 == "-ob1" {
                    option.ByteCompile = true
                    
                    if arg == "-ob0"{
                        option.StripDebugInfo = true
                        
                    }
                } else {
                    Util_log(Lns_getVM().String_format("unknown option -- '%s'", []LnsAny{arg}))
                    Lns_getVM().OS_exit(1)
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
                    _mode := Option_ModeKind__from(arg)
                    if !Lns_IsNil( _mode ) {
                        mode := _mode.(string)
                        option.Mode = mode
                        
                    } else {
                        Util_err(Lns_getVM().String_format("unknown mode -- %s", []LnsAny{arg}))
                    }
                }
            } else { 
                if _switch1939 := (option.Mode); _switch1939 == Option_ModeKind__Complete || _switch1939 == Option_ModeKind__Inquire {
                    if Lns_op_not(option.AnalyzeModule){
                        option.AnalyzeModule = arg
                        
                    } else if Lns_op_not(lineNo){
                        lineNo = Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(arg, nil), 0)))
                        
                    } else if Lns_op_not(column){
                        column = Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(arg, nil), 0)))
                        
                        option.AnalyzePos = NewTypes_Position(Lns_unwrap( lineNo).(LnsInt), Lns_unwrap( column).(LnsInt), Util_scriptPath2Module(option.ScriptPath))
                        
                    }
                } else if _switch1939 == Option_ModeKind__Save || _switch1939 == Option_ModeKind__SaveMeta || _switch1939 == Option_ModeKind__Glue {
                    option.OutputDir = arg
                    
                } else if _switch1939 == Option_ModeKind__Shebang {
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
        uptodateOpt_1729 := uptodateOpt.(string)
        if _switch2036 := uptodateOpt_1729; _switch2036 == "force" {
            option.TransCtrlInfo.UptodateMode = &Types_CheckingUptodateMode__Force1{Util_scriptPath2Module(option.ScriptPath)}
            
        } else if _switch2036 == "forceAll" {
            option.TransCtrlInfo.UptodateMode = Types_CheckingUptodateMode__ForceAll_Obj
            
        } else if _switch2036 == "normal" {
            option.TransCtrlInfo.UptodateMode = Types_CheckingUptodateMode__Normal_Obj
            
        } else if _switch2036 == "touch" {
            option.TransCtrlInfo.UptodateMode = Types_CheckingUptodateMode__Touch_Obj
            
        } else {
            Util_errorLog("illegal mode -- " + uptodateOpt_1729)
        }
    }
    if option.Mode != Option_ModeKind__Builtin{
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( option.ScriptPath == "") ||
            Lns_GetEnv().SetStackVal( option.Mode == Option_ModeKind__Unknown) ).(bool){
            analyze__printUsage_1073_(Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( argList.Len() == 0) ||
                    Lns_GetEnv().SetStackVal( argList.GetAt(1).(string) == "") ).(bool))) &&
                Lns_GetEnv().SetStackVal( 0) ||
                Lns_GetEnv().SetStackVal( 1) ).(LnsInt))
        }
    }
    if useStdInFlag{
        if Lns_isCondTrue( option.AnalyzeModule){
            Parser_StreamParser_setStdinStream(Lns_unwrap( option.AnalyzeModule).(string))
        } else { 
            if option.ScriptPath != ""{
                Parser_StreamParser_setStdinStream(Util_scriptPath2Module(option.ScriptPath))
            }
        }
    }
    if option.ScriptPath == "@-"{
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
                option.BatchList.Insert(line)
            }
        }
    }
    Log_log(Log_Level__Log, __func__, 624, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("mode is '%s'", []LnsAny{Option_ModeKind_getTxt( option.Mode)})
    }))
    
    return option
}

// 629: decl @lune.@base.@Option.createDefaultOption
func Option_createDefaultOption(pathList *LnsList,projDir LnsAny) *Option_Option {
    var option *Option_Option
    option = NewOption_Option()
    if pathList.Len() == 1{
        option.ScriptPath = pathList.GetAt(1).(string)
        
    } else { 
        option.ScriptPath = "@-"
        
        for _, _path := range( pathList.Items ) {
            path := _path.(string)
            option.BatchList.Insert(path)
        }
    }
    option.UseLuneModule = Option_getRuntimeModule()
    
    option.UseIpairs = true
    
    if projDir != nil{
        projDir_1760 := projDir.(string)
        if projDir_1760 != "/"{
            if Lns_op_not(Lns_car(Lns_getVM().String_find(projDir_1760,"/$", nil, nil))){
                option.projDir = projDir_1760 + "/"
                
            } else { 
                option.projDir = projDir_1760
                
            }
        }
    }
    return option
}

// declaration Class -- RuntimeOpt
type Option_RuntimeOptMtd interface {
    Get_int2strMode() LnsInt
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
func NewOption_RuntimeOpt() *Option_RuntimeOpt {
    obj := &Option_RuntimeOpt{}
    obj.FP = obj
    obj.InitOption_RuntimeOpt()
    return obj
}
func (self *Option_RuntimeOpt) Get_int2strMode() LnsInt{ return self.int2strMode }
// 92: DeclConstr
func (self *Option_RuntimeOpt) InitOption_RuntimeOpt() {
    self.int2strMode = Option_Int2strMode__Int2strModeDepend
    
}


// declaration Class -- Option
type Option_OptionMtd interface {
    Get_projDir() LnsAny
    Get_runtimeOpt() *Option_RuntimeOpt
    OpenDepend(arg1 LnsAny) LnsAny
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
func NewOption_Option() *Option_Option {
    obj := &Option_Option{}
    obj.FP = obj
    obj.InitOption_Option()
    return obj
}
func (self *Option_Option) Get_runtimeOpt() *Option_RuntimeOpt{ return self.runtimeOpt }
func (self *Option_Option) Get_projDir() LnsAny{ return self.projDir }
// 151: DeclConstr
func (self *Option_Option) InitOption_Option() {
    self.projDir = nil
    
    self.runtimeOpt = NewOption_RuntimeOpt()
    
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
    
    self.TargetLuaVer = LuaVer_getCurVer()
    
    self.TransCtrlInfo = Types_TransCtrlInfo_create_normal()
    
    self.BootPath = nil
    
    self.UseIpairs = false
    
    self.AnalyzeModule = nil
    
    self.OutputDir = nil
    
    self.AnalyzePos = nil
    
    self.DependsPath = nil
    
}

// 179: decl @lune.@base.@Option.Option.openDepend
func (self *Option_Option) OpenDepend(relPath LnsAny) LnsAny {
    {
        _path := self.DependsPath
        if !Lns_IsNil( _path ) {
            path := _path.(string)
            var filePath string
            if relPath != nil{
                relPath_1589 := relPath.(string)
                if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(path,"/$", nil, nil))){
                    filePath = Lns_getVM().String_format("%s%s", []LnsAny{path, relPath_1589})
                    
                } else { 
                    filePath = Lns_getVM().String_format("%s/%s", []LnsAny{path, relPath_1589})
                    
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
type Option_ProjInfo1075Mtd interface {
    ToMap() *LnsMap
}
type Option_ProjInfo1075 struct {
    Cmd_option *LnsList
    FP Option_ProjInfo1075Mtd
}
func Option_ProjInfo10752Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_ProjInfo1075).FP
}
type Option_ProjInfo1075DownCast interface {
    ToOption_ProjInfo1075() *Option_ProjInfo1075
}
func Option_ProjInfo1075DownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Option_ProjInfo1075DownCast)
    if ok { return work.ToOption_ProjInfo1075() }
    return nil
}
func (obj *Option_ProjInfo1075) ToOption_ProjInfo1075() *Option_ProjInfo1075 {
    return obj
}
func NewOption_ProjInfo1075(arg1 *LnsList) *Option_ProjInfo1075 {
    obj := &Option_ProjInfo1075{}
    obj.FP = obj
    obj.InitOption_ProjInfo1075(arg1)
    return obj
}
func (self *Option_ProjInfo1075) InitOption_ProjInfo1075(arg1 *LnsList) {
    self.Cmd_option = arg1
}
func (self *Option_ProjInfo1075) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["cmd_option"] = Lns_ToCollection( self.Cmd_option )
    return obj
}
func (self *Option_ProjInfo1075) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Option_ProjInfo1075__fromMap_1088_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Option_ProjInfo1075_FromMap( arg1, paramList )
}
func Option_ProjInfo1075__fromStem_1090_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Option_ProjInfo1075_FromMap( arg1, paramList )
}
func Option_ProjInfo1075_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Option_ProjInfo1075_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Option_ProjInfo1075_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Option_ProjInfo1075{}
    newObj.FP = newObj
    return Option_ProjInfo1075_FromMapMain( newObj, objMap, paramList )
}
func Option_ProjInfo1075_FromMapMain( newObj *Option_ProjInfo1075, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToListSub( objMap.Items["cmd_option"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil}}); !ok {
       return false,nil,"cmd_option:" + mess.(string)
    } else {
       newObj.Cmd_option = conv.(*LnsList)
    }
    return true, newObj, nil
}

func Lns_Option_init() {
    if init_Option { return }
    init_Option = true
    Option__mod__ = "@lune.@base.@Option"
    Lns_InitMod()
    Lns_Types_init()
    Lns_Parser_init()
    Lns_Json_init()
    Lns_Util_init()
    Lns_LuaMod_init()
    Lns_Ver_init()
    Lns_LuaVer_init()
    Lns_LuaVer_init()
    Lns_Depend_init()
    Lns_Log_init()
    Lns_Ast_init()
    
}
func init() {
    init_Option = false
}
