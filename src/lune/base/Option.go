// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_Option bool
var Option__mod__ string
// decl enum -- ModeKind 
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
const Option_ModeKind__Save = "save"
const Option_ModeKind__SaveMeta = "SAVE"
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
  Option_ModeKind__Save: "ModeKind.Save",
  Option_ModeKind__SaveMeta: "ModeKind.SaveMeta",
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
// for 39
func Option_convExp60(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 267
func Option_convExp715(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 262
func Option_convExp717(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 166
func Option_convExp517(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// 48: decl @lune.@base.@Option.getBuildCount
func Option_getBuildCount_1015_() LnsInt {
    return 6331
}

// 73: decl @lune.@base.@Option.getRuntimeModule
func Option_getRuntimeModule() string {
    return Lns_getVM().String_format("lune.base.runtime%d", []LnsAny{Ver_luaModVersion})
}

// 159: decl @lune.@base.@Option.outputLuneMod
func Option_outputLuneMod(path LnsAny) LnsAny {
    var lune_path string
    lune_path = "runtime.lua"
    if path != nil{
        path_2458 := path.(string)
        if path_2458 != ""{
            lune_path = path_2458
            
        }
    }
    var fileObj Lns_luaStream
    
    {
        _fileObj := Option_convExp517(Lns_2DDD(Lns_io_open(lune_path, "w")))
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

// 206: decl @lune.@base.@Option.analyze.printUsage
func analyze__printUsage_1060_(code LnsInt) {
    Lns_print([]LnsAny{"usage:\n  <type1> [-prof] [-r] src.lns mode [mode-option]\n  <type2> -mklunemod path\n  <type3> --version\n\n* type1\n  - src.lns [common_op] ast\n  - src.lns [common_op] comp [-i] module line column\n  - src.lns [common_op] inq [-i] module line column\n  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] <lua|LUA>\n  - src.lns [common_op] [-ol ver] [-ob<0|1>] [-dmr] [--depends dependfile] <save|SAVE> output-dir\n  - src.lns [common_op] exe\n\n  -r: use 'require( \"lune.base.runtime\" )'\n  -ol: output lua version. ver = 51 or 52 or 53.\n  -ob: output bytecompiled-code.\n      -ob0 is without debug information.\n      -ob1 is with debug information.\n  -langC: transcompile to c-lang.\n  -langGo: transcompile to golang.\n  -oc: output path of the source code transcompiled to c-lang .\n  --depends: output dependfile\n\n  common_op:\n    -u: update meta and lua on load.\n    -Werror: error by warrning.\n    --log <mode>: set log level.\n         mode: fatal, error, warn, log, info, debug, trace\n    --warning-shadowing: shadowing error convert to warning.\n    --compat-comment: backward compatibility to process the comment.\n    --disable-checking-define-abbr: disable checking for ##.\n    --uptodate <mode>: checking uptodate mode.\n            mode: skip check.\n            none: skip process when file is uptodate.\n            touch: touch meta file when file is uptodate.  (default)\n    --use-ipairs: use ipairs for foreach with List value.\n    --default-lazy: set lazy-loading at default.\n    --valid-luaval: enable luaval when transcompie to lua.\n    --package <name>: set the package name for the go-lang.\n\n* type2\n  dir: output directory.\n"})
    Lns_getVM().OS_exit(code)
}



// 204: decl @lune.@base.@Option.analyze
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
        _file := Option_convExp717(Lns_2DDD(Lns_io_open("lune.js", "r")))
        if _file != nil {
            file := _file.(Lns_luaStream)
            {
                _projInfo := Option_convExp715(Lns_2DDD(Option_ProjInfo1063__fromStem_1084_(Lns_car(Json_fromStr(Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( file.Read("*a")) ||
                    Lns_GetEnv().SetStackVal( "") ).(string))),nil)))
                if _projInfo != nil {
                    projInfo := _projInfo.(*Option_ProjInfo1063)
                    var workArgList *LnsList
                    workArgList = NewLnsList([]LnsAny{})
                    for _, _arg := range( argList.Items ) {
                        arg := _arg.(string)
                        workArgList.Insert(arg)
                    }
                    for _, _arg := range( projInfo.Cmd_option.Items ) {
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
    for argList.Len() >= index {
        var arg string
        arg = argList.GetAt(index).(string)
        if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(arg,"^-", nil, nil))){
            if _switch1400 := (arg); _switch1400 == "-i" {
                useStdInFlag = true
                
            } else if _switch1400 == "-prof" {
                option.ValidProf = true
                
            } else if _switch1400 == "--nodebug" {
                Util_setDebugFlag(false)
            } else if _switch1400 == "--debug" {
                Util_setDebugFlag(true)
            } else if _switch1400 == "--version" {
                Lns_print([]LnsAny{Lns_getVM().String_format("LuneScript: version %s (%d:%s) [%s]", []LnsAny{Ver_version, Option_getBuildCount_1015_(), Depend_getLuaVersion(), Ver_metaVersion})})
                Lns_getVM().OS_exit(0)
            } else if _switch1400 == "--builtin" {
                {
                    __collection902 := Ast_getBuiltInTypeIdMap()
                    __sorted902 := __collection902.CreateKeyListInt()
                    __sorted902.Sort( LnsItemKindInt, nil )
                    for _, _typeId := range( __sorted902.Items ) {
                        typeInfo := __collection902.Items[ _typeId ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        typeId := _typeId.(LnsInt)
                        Lns_print([]LnsAny{typeId, typeInfo.FP.GetTxt(nil, nil, nil)})
                    }
                }
                Lns_getVM().OS_exit(0)
            } else if _switch1400 == "-mklunemod" {
                var path LnsAny
                path = getNextOp()
                {
                    _mess := Option_outputLuneMod(path)
                    if _mess != nil {
                        mess := _mess.(string)
                        Util_errorLog(mess)
                        Lns_getVM().OS_exit(1)
                    }
                }
                Lns_getVM().OS_exit(0)
            } else if _switch1400 == "--mkbuiltin" {
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
                
            } else if _switch1400 == "-r" {
                option.UseLuneModule = Option_getRuntimeModule()
                
            } else if _switch1400 == "--runtime" {
                option.UseLuneModule = getNextOp()
                
            } else if _switch1400 == "-oc" {
                option.BootPath = getNextOp()
                
            } else if _switch1400 == "-u" {
                option.UpdateOnLoad = true
                
            } else if _switch1400 == "-Werror" {
                option.TransCtrlInfo.StopByWarning = true
                
            } else if _switch1400 == "--disable-checking-define-abbr" {
                option.TransCtrlInfo.CheckingDefineAbbr = false
                
            } else if _switch1400 == "--compat-comment" {
                option.TransCtrlInfo.CompatComment = true
                
            } else if _switch1400 == "--warning-shadowing" {
                option.TransCtrlInfo.WarningShadowing = true
                
            } else if _switch1400 == "--valid-luaval" {
                option.TransCtrlInfo.ValidLuaval = true
                
            } else if _switch1400 == "--default-lazy" {
                option.TransCtrlInfo.DefaultLazy = true
                
            } else if _switch1400 == "--package" {
                option.PackageName = getNextOp()
                
            } else if _switch1400 == "--log" {
                {
                    _txt := getNextOp()
                    if _txt != nil {
                        txt := _txt.(string)
                        {
                            _level := Log_str2level(txt)
                            if _level != nil {
                                level := _level.(LnsInt)
                                Log_setLevel(level)
                            } else {
                                Util_errorLog(Lns_getVM().String_format("illegal level -- %s", []LnsAny{txt}))
                            }
                        }
                    }
                }
            } else if _switch1400 == "--testing" {
                option.Testing = true
                
            } else if _switch1400 == "--depends" {
                option.DependsPath = getNextOp()
                
            } else if _switch1400 == "--use-ipairs" {
                option.UseIpairs = true
                
            } else if _switch1400 == "--uptodate" {
                {
                    _txt := getNextOp()
                    if _txt != nil {
                        txt := _txt.(string)
                        {
                            _mode := Types_CheckingUptodateMode__from(txt)
                            if _mode != nil {
                                mode := _mode.(string)
                                option.TransCtrlInfo.UptodateMode = mode
                                
                            } else {
                                Util_errorLog("illegal mode -- " + txt)
                            }
                        }
                    }
                }
            } else if _switch1400 == "-langC" {
                option.ConvTo = Types_Conv__C
                
                option.TransCtrlInfo.ValidLuaval = true
                
            } else if _switch1400 == "-langGo" {
                option.ConvTo = Types_Conv__Go
                
                option.TransCtrlInfo.ValidLuaval = true
                
            } else if _switch1400 == "-ol" {
                {
                    _txt := getNextOp()
                    if _txt != nil {
                        txt := _txt.(string)
                        if _switch1344 := txt; _switch1344 == "51" {
                            option.TargetLuaVer = LuaVer_ver51
                            
                        } else if _switch1344 == "52" {
                            option.TargetLuaVer = LuaVer_ver52
                            
                        } else if _switch1344 == "53" {
                            option.TargetLuaVer = LuaVer_ver53
                            
                        }
                    }
                }
            } else if _switch1400 == "-ob0" || _switch1400 == "-ob1" {
                option.ByteCompile = true
                
                if arg == "-ob0"{
                    option.StripDebugInfo = true
                    
                }
            } else {
                Util_log(Lns_getVM().String_format("unknown option -- %s", []LnsAny{arg}))
                Lns_getVM().OS_exit(1)
            }
        } else { 
            if option.ScriptPath == ""{
                option.ScriptPath = arg
                
            } else if option.Mode == ""{
                {
                    _mode := Option_ModeKind__from(arg)
                    if _mode != nil {
                        mode := _mode.(string)
                        option.Mode = mode
                        
                    } else {
                        Util_err(Lns_getVM().String_format("unknown mode -- %s", []LnsAny{arg}))
                    }
                }
            } else { 
                if _switch1583 := (option.Mode); _switch1583 == Option_ModeKind__Complete || _switch1583 == Option_ModeKind__Inquire {
                    if Lns_op_not(option.AnalyzeModule){
                        option.AnalyzeModule = arg
                        
                    } else if Lns_op_not(lineNo){
                        lineNo = Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(arg, nil), 0)))
                        
                    } else if Lns_op_not(column){
                        column = Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(arg, nil), 0)))
                        
                        option.AnalyzePos = NewTypes_Position(Lns_unwrap( lineNo).(LnsInt), Lns_unwrap( column).(LnsInt), Util_scriptPath2Module(option.ScriptPath))
                        
                    }
                } else if _switch1583 == Option_ModeKind__Save || _switch1583 == Option_ModeKind__SaveMeta || _switch1583 == Option_ModeKind__Glue {
                    option.OutputDir = arg
                    
                } else {
                }
            }
        }
        index = index + 1
        
    }
    if option.Mode != Option_ModeKind__Builtin{
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( option.ScriptPath == "") ||
            Lns_GetEnv().SetStackVal( option.Mode == Option_ModeKind__Unknown) ).(bool){
            analyze__printUsage_1060_(Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
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
    Log_log(Log_Level__Log, __func__, 490, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("mode is '%s'", []LnsAny{Option_ModeKind_getTxt( option.Mode)})
    }))
    
    return option
}

// 495: decl @lune.@base.@Option.createDefaultOption
func Option_createDefaultOption(path string) *Option_Option {
    var option *Option_Option
    option = NewOption_Option()
    option.ScriptPath = path
    
    option.UseLuneModule = Option_getRuntimeModule()
    
    option.UseIpairs = true
    
    return option
}

// declaration Class -- Option
type Option_OptionMtd interface {
    OpenDepend(arg1 LnsAny) LnsAny
}
type Option_Option struct {
    Mode string
    AnalyzeModule LnsAny
    AnalyzePos LnsAny
    OutputDir LnsAny
    ScriptPath string
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
    TransCtrlInfo *Types_TransCtrlInfo
    ConvTo LnsAny
    Testing bool
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
// 117: DeclConstr
func (self *Option_Option) InitOption_Option() {
    self.PackageName = nil
    
    self.Testing = false
    
    self.ConvTo = nil
    
    self.ValidProf = false
    
    self.Mode = Option_ModeKind__Unknown
    
    self.ScriptPath = ""
    
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

// 138: decl @lune.@base.@Option.Option.openDepend
func (self *Option_Option) OpenDepend(relPath LnsAny) LnsAny {
    {
        _path := self.DependsPath
        if _path != nil {
            path := _path.(string)
            var filePath string
            if relPath != nil{
                relPath_2446 := relPath.(string)
                if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(path,"/$", nil, nil))){
                    filePath = Lns_getVM().String_format("%s%s", []LnsAny{path, relPath_2446})
                    
                } else { 
                    filePath = Lns_getVM().String_format("%s/%s", []LnsAny{path, relPath_2446})
                    
                }
            } else {
                filePath = path
                
            }
            return Lns_car(Lns_io_open(filePath, "w"))
        }
    }
    {
        _path := self.DependsPath
        if _path != nil {
            path := _path.(string)
            return Lns_car(Lns_io_open(path, "w"))
        }
    }
    return nil
}


// declaration Class -- ProjInfo
type Option_ProjInfo1063Mtd interface {
    ToMap() *LnsMap
}
type Option_ProjInfo1063 struct {
    Cmd_option *LnsList
    FP Option_ProjInfo1063Mtd
}
func Option_ProjInfo10632Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Option_ProjInfo1063).FP
}
type Option_ProjInfo1063DownCast interface {
    ToOption_ProjInfo1063() *Option_ProjInfo1063
}
func Option_ProjInfo1063DownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Option_ProjInfo1063DownCast)
    if ok { return work.ToOption_ProjInfo1063() }
    return nil
}
func (obj *Option_ProjInfo1063) ToOption_ProjInfo1063() *Option_ProjInfo1063 {
    return obj
}
func NewOption_ProjInfo1063(arg1 *LnsList) *Option_ProjInfo1063 {
    obj := &Option_ProjInfo1063{}
    obj.FP = obj
    obj.InitOption_ProjInfo1063(arg1)
    return obj
}
func (self *Option_ProjInfo1063) InitOption_ProjInfo1063(arg1 *LnsList) {
    self.Cmd_option = arg1
}
func (self *Option_ProjInfo1063) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["cmd_option"] = Lns_ToCollection( self.Cmd_option )
    return obj
}
func (self *Option_ProjInfo1063) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Option_ProjInfo1063__fromMap_1081_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Option_ProjInfo1063_FromMap( arg1, paramList )
}
func Option_ProjInfo1063__fromStem_1084_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Option_ProjInfo1063_FromMap( arg1, paramList )
}
func Option_ProjInfo1063_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Option_ProjInfo1063_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Option_ProjInfo1063_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Option_ProjInfo1063{}
    newObj.FP = newObj
    return Option_ProjInfo1063_FromMapMain( newObj, objMap, paramList )
}
func Option_ProjInfo1063_FromMapMain( newObj *Option_ProjInfo1063, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
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
