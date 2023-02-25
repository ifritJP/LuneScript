// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_GoMod bool
var GoMod__mod__ string
// decl enum -- BlockKind 
type GoMod_BlockKind = LnsInt
const GoMod_BlockKind__None = 0
const GoMod_BlockKind__Replace = 2
const GoMod_BlockKind__Require = 1
var GoMod_BlockKindList_ = NewLnsList( []LnsAny {
  GoMod_BlockKind__None,
  GoMod_BlockKind__Require,
  GoMod_BlockKind__Replace,
})
func GoMod_BlockKind_get__allList_2_(_env *LnsEnv) *LnsList{
    return GoMod_BlockKindList_
}
var GoMod_BlockKindMap_ = map[LnsInt]string {
  GoMod_BlockKind__None: "BlockKind.None",
  GoMod_BlockKind__Replace: "BlockKind.Replace",
  GoMod_BlockKind__Require: "BlockKind.Require",
}
func GoMod_BlockKind__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := GoMod_BlockKindMap_[arg1]; ok { return arg1 }
    return nil
}

func GoMod_BlockKind_getTxt(arg1 LnsInt) string {
    return GoMod_BlockKindMap_[arg1];
}
// decl alge -- GoModResult
type GoMod_GoModResult = LnsAny
type GoMod_GoModResult__Found struct{
Val1 *GoMod_ModProjInfo
}
func (self *GoMod_GoModResult__Found) GetTxt() string {
return "GoModResult.Found"
}
type GoMod_GoModResult__NotFound struct{
}
var GoMod_GoModResult__NotFound_Obj = &GoMod_GoModResult__NotFound{}
func (self *GoMod_GoModResult__NotFound) GetTxt() string {
return "GoModResult.NotFound"
}
type GoMod_GoModResult__NotGo struct{
}
var GoMod_GoModResult__NotGo_Obj = &GoMod_GoModResult__NotGo{}
func (self *GoMod_GoModResult__NotGo) GetTxt() string {
return "GoModResult.NotGo"
}
// for 314
func GoMod_convExp0_1732(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 176
func GoMod_convExp0_838(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 173
func GoMod_convExp0_800(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 248
func GoMod_convExp0_1248(arg1 []LnsAny) (string, string) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 320
func GoMod_convExp0_1517(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 74
func GoMod_convExp0_335(arg1 []LnsAny) []LnsAny {
    return Lns_2DDD( arg1[0:])
}
// for 147
func GoMod_convExp0_636(arg1 []LnsAny) []LnsAny {
    return Lns_2DDD( arg1[0:])
}
// 298: decl @lune.@base.@GoMod.getReplace
func GoMod_getReplace_4_(_env *LnsEnv, _map *LnsMap2_[string,string],tokenList *LnsList2_[string],modIndex LnsInt) {
    var prevArrow bool
    prevArrow = false
    for _, _token := range( tokenList.Items ) {
        token := _token
        if token == "=>"{
            prevArrow = true
        } else if prevArrow{
            _map.Set(tokenList.GetAt(modIndex),token)
            break
        }
    }
}

// 310: decl @lune.@base.@GoMod.getGoMap
func GoMod_getGoMap(_env *LnsEnv) *GoMod_ModInfo {
    var requireMap *LnsMap2_[string,string]
    requireMap = NewLnsMap2_[string,string]( map[string]string{})
    var replaceMap *LnsMap2_[string,string]
    replaceMap = NewLnsMap2_[string,string]( map[string]string{})
    var name string
    name = "lnsc"
    {
        _file := GoMod_convExp0_1732(Lns_2DDD(Lns_io_open("go.mod", nil)))
        if !Lns_IsNil( _file ) {
            file := _file.(Lns_luaStream)
            var inBlock LnsInt
            inBlock = GoMod_BlockKind__None
            for  {
                var line string
                
                {
                    _line := file.Read(_env, "*l")
                    if _line == nil{
                        break
                    } else {
                        line = _line.(string)
                    }
                }
                var trimedLine string
                trimedLine = GoMod_convExp0_1517(Lns_2DDD(_env.GetVM().String_gsub(line,"^%s", "")))
                var tokenList *LnsList2_[string]
                tokenList = Util_splitStr(_env, trimedLine, "[^%s]+")
                if _switch0 := inBlock; _switch0 == GoMod_BlockKind__Require {
                    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(line,"^%)", nil, nil))){
                        inBlock = GoMod_BlockKind__None
                    } else { 
                        requireMap.Set(tokenList.GetAt(1),tokenList.GetAt(2))
                    }
                } else if _switch0 == GoMod_BlockKind__Replace {
                    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(line,"^%)", nil, nil))){
                        inBlock = GoMod_BlockKind__None
                    } else { 
                        GoMod_getReplace_4_(_env, replaceMap, tokenList, 1)
                    }
                } else if _switch0 == GoMod_BlockKind__None {
                    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(line,"^module%s+", nil, nil))){
                        name = tokenList.GetAt(2)
                    } else if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(line,"^require%s+[^%(]", nil, nil))){
                        if tokenList.Len() == 3{
                            requireMap.Set(tokenList.GetAt(2),tokenList.GetAt(3))
                        }
                    } else if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(line,"^require%s+%(", nil, nil))){
                        inBlock = GoMod_BlockKind__Require
                    } else if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(line,"^replace%s+[^%(]", nil, nil))){
                        GoMod_getReplace_4_(_env, replaceMap, tokenList, 2)
                    } else if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(line,"^replace%s+%(", nil, nil))){
                        inBlock = GoMod_BlockKind__Replace
                    }
                }
            }
            file.Close(_env)
        }
    }
    var modInfo *GoMod_ModInfo
    modInfo = NewGoMod_ModInfo(_env, name, requireMap, replaceMap)
    return modInfo
}



// 67: decl @lune.@base.@GoMod.ModInfo.getGoModPath
func (self *GoMod_ModInfo) getGoModPath(_env *LnsEnv, ver string,mod string) *LnsList2_[string] {
    var pathList *LnsList2_[string]
    pathList = NewLnsList2_[string]([]string{})
    {
        _gopath := Depend_getGOPATH(_env)
        if !Lns_IsNil( _gopath ) {
            gopath := _gopath.(string)
            pathList.Insert(Util_pathJoin(_env, gopath, _env.GetVM().String_format("src/%s", Lns_2DDD(mod))))
            var gomod string
            gomod = ""
            for _, _aChar := range( NewLnsList2_[LnsAny](Lns_2Slice[LnsAny]( Lns_2DDD(_env.GetVM().String_byte(mod,1, len(mod))))).Items ) {
                aChar := _aChar
                if aChar != nil{
                    aChar_107 := aChar.(LnsInt)
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( aChar_107 >= 65) &&
                        _env.SetStackVal( aChar_107 <= 90) ).(bool)){
                        gomod = _env.GetVM().String_format("%s!%c", Lns_2DDD(gomod, aChar_107 - 65 + 97))
                    } else { 
                        gomod = _env.GetVM().String_format("%s%c", Lns_2DDD(gomod, aChar_107))
                    }
                }
            }
            gomod = _env.GetVM().String_format("%s@%s", Lns_2DDD(gomod, ver))
            pathList.Insert(Util_pathJoin(_env, gopath, _env.GetVM().String_format("pkg/mod/%s", Lns_2DDD(gomod))))
        }
    }
    return pathList
}
// 90: decl @lune.@base.@GoMod.ModInfo.getModPathList
func (self *GoMod_ModInfo) GetModPathList(_env *LnsEnv) *LnsList {
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    for _mod, _ver := range( self.moduleMap.Items ) {
        mod := _mod
        ver := _ver
        {
            __exp := self.replaceMap.Get(mod)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(string)
                list.Insert(_exp)
            } else {
                for _, _path := range( self.FP.getGoModPath(_env, ver, mod).Items ) {
                    path := _path
                    if Depend_existFile(_env, path){
                        list.Insert(path)
                        break
                    }
                }
            }
        }
    }
    return list
}
// 125: decl @lune.@base.@GoMod.ModInfo.getLatestProjRoot
func (self *GoMod_ModInfo) GetLatestProjRoot(_env *LnsEnv) LnsAny {
    return _env.NilAccFin(_env.NilAccPush(self.latestModProjInfo) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*GoMod_ModProjInfo).FP.Get_projRoot(_env)}))
}
// 129: decl @lune.@base.@GoMod.ModInfo.getLocalModulePathList
func (self *GoMod_ModInfo) getLocalModulePathList(_env *LnsEnv, path string) *LnsList2_[string] {
    var pathList *LnsList2_[string]
    pathList = NewLnsList2_[string]([]string{})
    for _mod, _ver := range( self.moduleMap.Items ) {
        mod := _mod
        ver := _ver
        if Lns_car(_env.GetVM().String_find(path,mod, 1, true)) == 1{
            var relativeName string
            relativeName = _env.GetVM().String_sub(path,len(mod) + 2, nil)
            {
                _replacePath := self.replaceMap.Get(mod)
                if !Lns_IsNil( _replacePath ) {
                    replacePath := _replacePath.(string)
                    pathList.Insert(Util_pathJoin(_env, replacePath, relativeName))
                    break
                }
            }
            {
                _gopath := Depend_getGOPATH(_env)
                if !Lns_IsNil( _gopath ) {
                    gopath := _gopath.(string)
                    pathList.Insert(Util_pathJoin(_env, gopath, _env.GetVM().String_format("src/%s", Lns_2DDD(path))))
                    var gomod string
                    gomod = ""
                    for _, _aChar := range( NewLnsList2_[LnsAny](Lns_2Slice[LnsAny]( Lns_2DDD(_env.GetVM().String_byte(mod,1, len(mod))))).Items ) {
                        aChar := _aChar
                        if aChar != nil{
                            aChar_137 := aChar.(LnsInt)
                            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( aChar_137 >= 65) &&
                                _env.SetStackVal( aChar_137 <= 90) ).(bool)){
                                gomod = _env.GetVM().String_format("%s!%c", Lns_2DDD(gomod, aChar_137 - 65 + 97))
                            } else { 
                                gomod = _env.GetVM().String_format("%s%c", Lns_2DDD(gomod, aChar_137))
                            }
                        }
                    }
                    gomod = _env.GetVM().String_format("%s@%s", Lns_2DDD(gomod, ver))
                    pathList.Insert(Util_pathJoin(_env, gopath, _env.GetVM().String_format("pkg/mod/%s/%s", Lns_2DDD(gomod, relativeName))))
                }
            }
            break
        }
    }
    return pathList
}
// 167: decl @lune.@base.@GoMod.ModInfo.convPath
func (self *GoMod_ModInfo) convPath(_env *LnsEnv, mod string,suffix string) string {
    return Lns_car(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(mod,"^go/", "")).(string),"%.", "/")).(string),":", ".")).(string) + suffix
}
// 171: decl @lune.@base.@GoMod.ModInfo.getProjRootPath
func (self *GoMod_ModInfo) getProjRootPath(_env *LnsEnv, mod string,path string)(string, string) {
    var convPath string
    convPath = GoMod_convExp0_800(Lns_2DDD(_env.GetVM().String_gsub(self.FP.convPath(_env, mod, ".lns"),"github%.com/[^/]+/[^/]+/", "")))
    var projRoot string
    projRoot = _env.GetVM().String_sub(path,1, len(path) - len(convPath))
    if projRoot != "/"{
        projRoot = GoMod_convExp0_838(Lns_2DDD(_env.GetVM().String_gsub(projRoot,"/$", "")))
    }
    path = Util_parentPath(_env, path)
    var modList *LnsList2_[string]
    modList = Util_splitStr(_env, mod, "[^%.]+")
    var startIndex LnsInt
    startIndex = modList.Len()
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = modList.Len()
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            modIndex := _forWork0
            if Depend_existFile(_env, Util_pathJoin(_env, path, "lune.js")){
                startIndex = modIndex
                break
            }
            if path == projRoot{
                startIndex = modIndex
                break
            }
            path = Util_parentPath(_env, path)
        }
    }
    var convMod string
    convMod = ""
    {
        var _forFrom1 LnsInt = modList.Len() - startIndex + 1
        var _forTo1 LnsInt = modList.Len()
        for _forWork1 := _forFrom1; _forWork1 <= _forTo1; _forWork1++ {
            index := _forWork1
            if convMod != ""{
                convMod = _env.GetVM().String_format("%s.", Lns_2DDD(convMod))
            }
            convMod = convMod + modList.GetAt(index)
        }
    }
    return path, convMod
}
// 211: decl @lune.@base.@GoMod.ModInfo.convLocalModulePath
func (self *GoMod_ModInfo) ConvLocalModulePath(_env *LnsEnv, mod string,suffix string,baseDir LnsAny) LnsAny {
    __func__ := "@lune.@base.@GoMod.ModInfo.convLocalModulePath"
    if Lns_op_not(Lns_car(_env.GetVM().String_find(mod,"^go/", nil, nil))){
        if baseDir != nil{
            baseDir_158 := baseDir.(string)
            var goModDir LnsAny
            goModDir = nil
            goModDir = self.goModDir2Path.Get(baseDir_158)
            if goModDir != nil{
                goModDir_161 := goModDir.(string)
                mod = _env.GetVM().String_format("go/%s.%s", Lns_2DDD(Lns_car(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(goModDir_161,"%.", ":")).(string),"/", ".")).(string), mod))
            }
            if Lns_op_not(goModDir){
                Log_log(_env, Log_Level__Log, __func__, 220, Log_CreateMessage(func(_env *LnsEnv) string {
                    return _env.GetVM().String_format("not found baseDir -- %s", Lns_2DDD(baseDir_158))
                }))
                
            }
        } else {
            return GoMod_GoModResult__NotGo_Obj
        }
    }
    var workMod string
    workMod = self.FP.convPath(_env, mod, suffix)
    {
        __exp := self.path2modProjInfo.Get(workMod)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*GoMod_ModProjInfo)
            self.latestModProjInfo = _exp
            return &GoMod_GoModResult__Found{_exp}
        }
    }
    var pathList *LnsList2_[string]
    pathList = self.FP.getLocalModulePathList(_env, workMod)
    if Lns_op_not(Lns_car(_env.GetVM().String_find(mod,"^go/", nil, nil))){
        if baseDir != nil{
            baseDir_172 := baseDir.(string)
            pathList.Insert(Util_pathJoin(_env, baseDir_172, workMod))
        }
    }
    pathList.Insert(Util_pathJoin(_env, "vendor", workMod))
    for _, _path := range( pathList.Items ) {
        path := _path
        if Depend_existFile(_env, path){
            var projRoot string
            var convMod string
            projRoot,convMod = self.FP.getProjRootPath(_env, mod, path)
            var projInfo *GoMod_ModProjInfo
            projInfo = NewGoMod_ModProjInfo(_env, path, projRoot, convMod, mod)
            self.path2modProjInfo.Set(workMod,projInfo)
            self.latestModProjInfo = projInfo
            return &GoMod_GoModResult__Found{projInfo}
        } else { 
            Log_log(_env, Log_Level__Log, __func__, 254, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.GetVM().String_format("not found %s", Lns_2DDD(path))
            }))
            
        }
    }
    return GoMod_GoModResult__NotFound_Obj
}
// 275: decl @lune.@base.@GoMod.ModInfo.getLuaModulePath
func (self *GoMod_ModInfo) GetLuaModulePath(_env *LnsEnv, mod string,baseDir LnsAny)(string, LnsAny, string) {
    var info *GoMod_ModProjInfo
    switch _matchExp0 := self.FP.ConvLocalModulePath(_env, mod, ".lns", baseDir).(type) {
    case *GoMod_GoModResult__NotGo:
        return mod, nil, mod
    case *GoMod_GoModResult__NotFound:
        return mod, nil, mod
    case *GoMod_GoModResult__Found:
        workInfo := _matchExp0.Val1
        info = workInfo
    }
    return info.FP.Get_mod(_env), info.FP.Get_projRoot(_env), info.FP.Get_fullMod(_env)
}
// declaration Class -- ModProjInfo
type GoMod_ModProjInfoMtd interface {
    Get_fullMod(_env *LnsEnv) string
    Get_mod(_env *LnsEnv) string
    Get_path(_env *LnsEnv) string
    Get_projRoot(_env *LnsEnv) string
}
type GoMod_ModProjInfo struct {
    path string
    projRoot string
    mod string
    fullMod string
    FP GoMod_ModProjInfoMtd
}
func GoMod_ModProjInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*GoMod_ModProjInfo).FP
}
func GoMod_ModProjInfo_toSlice(slice []LnsAny) []*GoMod_ModProjInfo {
    ret := make([]*GoMod_ModProjInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(GoMod_ModProjInfoDownCast).ToGoMod_ModProjInfo()
    }
    return ret
}
type GoMod_ModProjInfoDownCast interface {
    ToGoMod_ModProjInfo() *GoMod_ModProjInfo
}
func GoMod_ModProjInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(GoMod_ModProjInfoDownCast)
    if ok { return work.ToGoMod_ModProjInfo() }
    return nil
}
func (obj *GoMod_ModProjInfo) ToGoMod_ModProjInfo() *GoMod_ModProjInfo {
    return obj
}
func NewGoMod_ModProjInfo(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 string) *GoMod_ModProjInfo {
    obj := &GoMod_ModProjInfo{}
    obj.FP = obj
    obj.InitGoMod_ModProjInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *GoMod_ModProjInfo) InitGoMod_ModProjInfo(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 string) {
    self.path = arg1
    self.projRoot = arg2
    self.mod = arg3
    self.fullMod = arg4
}
func (self *GoMod_ModProjInfo) Get_path(_env *LnsEnv) string{ return self.path }
func (self *GoMod_ModProjInfo) Get_projRoot(_env *LnsEnv) string{ return self.projRoot }
func (self *GoMod_ModProjInfo) Get_mod(_env *LnsEnv) string{ return self.mod }
func (self *GoMod_ModProjInfo) Get_fullMod(_env *LnsEnv) string{ return self.fullMod }

// declaration Class -- ModInfo
type GoMod_ModInfoMtd interface {
    ConvLocalModulePath(_env *LnsEnv, arg1 string, arg2 string, arg3 LnsAny) LnsAny
    convPath(_env *LnsEnv, arg1 string, arg2 string) string
    getGoModPath(_env *LnsEnv, arg1 string, arg2 string) *LnsList2_[string]
    GetLatestProjRoot(_env *LnsEnv) LnsAny
    getLocalModulePathList(_env *LnsEnv, arg1 string) *LnsList2_[string]
    GetLuaModulePath(_env *LnsEnv, arg1 string, arg2 LnsAny)(string, LnsAny, string)
    GetModPathList(_env *LnsEnv) *LnsList
    getProjRootPath(_env *LnsEnv, arg1 string, arg2 string)(string, string)
    Get_moduleMap(_env *LnsEnv) *LnsMap2_[string,string]
    Get_name(_env *LnsEnv) string
    Get_replaceMap(_env *LnsEnv) *LnsMap2_[string,string]
}
type GoMod_ModInfo struct {
    name string
    moduleMap *LnsMap2_[string,string]
    replaceMap *LnsMap2_[string,string]
    path2modProjInfo *LnsMap2_[string,*GoMod_ModProjInfo]
    goModDir2Path *LnsMap2_[string,string]
    latestModProjInfo LnsAny
    FP GoMod_ModInfoMtd
}
func GoMod_ModInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*GoMod_ModInfo).FP
}
func GoMod_ModInfo_toSlice(slice []LnsAny) []*GoMod_ModInfo {
    ret := make([]*GoMod_ModInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(GoMod_ModInfoDownCast).ToGoMod_ModInfo()
    }
    return ret
}
type GoMod_ModInfoDownCast interface {
    ToGoMod_ModInfo() *GoMod_ModInfo
}
func GoMod_ModInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(GoMod_ModInfoDownCast)
    if ok { return work.ToGoMod_ModInfo() }
    return nil
}
func (obj *GoMod_ModInfo) ToGoMod_ModInfo() *GoMod_ModInfo {
    return obj
}
func NewGoMod_ModInfo(_env *LnsEnv, arg1 string, arg2 *LnsMap2_[string,string], arg3 *LnsMap2_[string,string]) *GoMod_ModInfo {
    obj := &GoMod_ModInfo{}
    obj.FP = obj
    obj.InitGoMod_ModInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *GoMod_ModInfo) Get_name(_env *LnsEnv) string{ return self.name }
func (self *GoMod_ModInfo) Get_moduleMap(_env *LnsEnv) *LnsMap2_[string,string]{ return self.moduleMap }
func (self *GoMod_ModInfo) Get_replaceMap(_env *LnsEnv) *LnsMap2_[string,string]{ return self.replaceMap }
// 107: DeclConstr
func (self *GoMod_ModInfo) InitGoMod_ModInfo(_env *LnsEnv, name string,moduleMap *LnsMap2_[string,string],replaceMap *LnsMap2_[string,string]) {
    self.name = name
    self.moduleMap = moduleMap
    self.replaceMap = replaceMap
    self.goModDir2Path = NewLnsMap2_[string,string]( map[string]string{})
    self.path2modProjInfo = NewLnsMap2_[string,*GoMod_ModProjInfo]( map[string]*GoMod_ModProjInfo{})
    self.latestModProjInfo = nil
    for _mod, _dst := range( replaceMap.Items ) {
        mod := _mod
        dst := _dst
        self.goModDir2Path.Set(dst,mod)
    }
    for _mod, _ver := range( moduleMap.Items ) {
        mod := _mod
        ver := _ver
        for _, _path := range( self.FP.getGoModPath(_env, ver, mod).Items ) {
            path := _path
            self.goModDir2Path.Set(path,mod)
        }
    }
}


func Lns_GoMod_init(_env *LnsEnv) {
    if init_GoMod { return }
    init_GoMod = true
    GoMod__mod__ = "@lune.@base.@GoMod"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_Util_init(_env)
    Lns_LuaVer_init( _env )
    Lns_Depend_init(_env)
    Lns_Log_init(_env)
}
func init() {
    init_GoMod = false
}
