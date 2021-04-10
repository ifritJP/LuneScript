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
func GoMod_BlockKind_get__allList_1116_() *LnsList{
    return GoMod_BlockKindList_
}
var GoMod_BlockKindMap_ = map[LnsInt]string {
  GoMod_BlockKind__None: "BlockKind.None",
  GoMod_BlockKind__Replace: "BlockKind.Replace",
  GoMod_BlockKind__Require: "BlockKind.Require",
}
func GoMod_BlockKind__from_1109_(arg1 LnsInt) LnsAny{
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
// for 231
func GoMod_convExp1174(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 83
func GoMod_convExp417(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 80
func GoMod_convExp379(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 237
func GoMod_convExp984(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}

// 212: decl @lune.@base.@GoMod.getReplace
func GoMod_getReplace_1126_(_map *LnsMap,tokenList *LnsList,modIndex LnsInt) {
    var prevArrow bool
    prevArrow = false
    for _, _token := range( tokenList.Items ) {
        token := _token.(string)
        if token == "=>"{
            prevArrow = true
            
        } else if prevArrow{
            _map.Set(tokenList.GetAt(modIndex).(string),token)
            break
        }
    }
}

// 224: decl @lune.@base.@GoMod.getGoMap
func GoMod_getGoMap(option *Option_Option) *GoMod_ModInfo {
    var requireMap *LnsMap
    requireMap = NewLnsMap( map[LnsAny]LnsAny{})
    var replaceMap *LnsMap
    replaceMap = NewLnsMap( map[LnsAny]LnsAny{})
    var modInfo *GoMod_ModInfo
    modInfo = NewGoMod_ModInfo(requireMap, replaceMap)
    {
        _file := GoMod_convExp1174(Lns_2DDD(Lns_io_open("go.mod", nil)))
        if _file != nil {
            file := _file.(Lns_luaStream)
            var inBlock LnsInt
            inBlock = GoMod_BlockKind__None
            for  {
                var line string
                
                {
                    _line := file.Read("*l")
                    if _line == nil{
                        break
                    } else {
                        line = _line.(string)
                    }
                }
                var trimedLine string
                trimedLine = GoMod_convExp984(Lns_2DDD(Lns_getVM().String_gsub(line,"^%s", "")))
                var tokenList *LnsList
                tokenList = Util_splitStr(trimedLine, "[^%s]+")
                if _switch1170 := inBlock; _switch1170 == GoMod_BlockKind__Require {
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^%)", nil, nil))){
                        inBlock = GoMod_BlockKind__None
                        
                    } else { 
                        requireMap.Set(tokenList.GetAt(1).(string),tokenList.GetAt(2).(string))
                    }
                } else if _switch1170 == GoMod_BlockKind__Replace {
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^%)", nil, nil))){
                        inBlock = GoMod_BlockKind__None
                        
                    } else { 
                        GoMod_getReplace_1126_(replaceMap, tokenList, 1)
                    }
                } else if _switch1170 == GoMod_BlockKind__None {
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^require%s+[^%(]", nil, nil))){
                        if tokenList.Len() == 3{
                            requireMap.Set(tokenList.GetAt(2).(string),tokenList.GetAt(3).(string))
                        }
                    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^require%s+%(", nil, nil))){
                        inBlock = GoMod_BlockKind__Require
                        
                    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^replace%s+[^%(]", nil, nil))){
                        GoMod_getReplace_1126_(replaceMap, tokenList, 2)
                    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^replace%s+%(", nil, nil))){
                        inBlock = GoMod_BlockKind__Replace
                        
                    }
                }
            }
        }
    }
    return modInfo
}

// declaration Class -- ModProjInfo
type GoMod_ModProjInfoMtd interface {
    Get_mod() string
    Get_path() string
    Get_projRoot() string
}
type GoMod_ModProjInfo struct {
    path string
    projRoot string
    mod string
    FP GoMod_ModProjInfoMtd
}
func GoMod_ModProjInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*GoMod_ModProjInfo).FP
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
func NewGoMod_ModProjInfo(arg1 string, arg2 string, arg3 string) *GoMod_ModProjInfo {
    obj := &GoMod_ModProjInfo{}
    obj.FP = obj
    obj.InitGoMod_ModProjInfo(arg1, arg2, arg3)
    return obj
}
func (self *GoMod_ModProjInfo) InitGoMod_ModProjInfo(arg1 string, arg2 string, arg3 string) {
    self.path = arg1
    self.projRoot = arg2
    self.mod = arg3
}
func (self *GoMod_ModProjInfo) Get_path() string{ return self.path }
func (self *GoMod_ModProjInfo) Get_projRoot() string{ return self.projRoot }
func (self *GoMod_ModProjInfo) Get_mod() string{ return self.mod }

// declaration Class -- ModInfo
type GoMod_ModInfoMtd interface {
    ConvLocalModulePath(arg1 string, arg2 string) LnsAny
    convPath(arg1 string, arg2 string) string
    GetLatestProjRoot() LnsAny
    getLocalModulePath(arg1 string) LnsAny
    GetLuaModulePath(arg1 string) string
    getProjRootPath(arg1 string, arg2 string)(string, string)
    Get_moduleMap() *LnsMap
    Get_replaceMap() *LnsMap
}
type GoMod_ModInfo struct {
    moduleMap *LnsMap
    replaceMap *LnsMap
    path2modProjInfo *LnsMap
    latestModProjInfo LnsAny
    FP GoMod_ModInfoMtd
}
func GoMod_ModInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*GoMod_ModInfo).FP
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
func NewGoMod_ModInfo(arg1 *LnsMap, arg2 *LnsMap) *GoMod_ModInfo {
    obj := &GoMod_ModInfo{}
    obj.FP = obj
    obj.InitGoMod_ModInfo(arg1, arg2)
    return obj
}
func (self *GoMod_ModInfo) Get_moduleMap() *LnsMap{ return self.moduleMap }
func (self *GoMod_ModInfo) Get_replaceMap() *LnsMap{ return self.replaceMap }
// 34: DeclConstr
func (self *GoMod_ModInfo) InitGoMod_ModInfo(moduleMap *LnsMap,replaceMap *LnsMap) {
    self.moduleMap = moduleMap
    
    self.replaceMap = replaceMap
    
    self.path2modProjInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.latestModProjInfo = nil
    
}

// 41: decl @lune.@base.@GoMod.ModInfo.getLatestProjRoot
func (self *GoMod_ModInfo) GetLatestProjRoot() LnsAny {
    return Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.latestModProjInfo) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*GoMod_ModProjInfo).FP.Get_projRoot()}))
}

// 45: decl @lune.@base.@GoMod.ModInfo.getLocalModulePath
func (self *GoMod_ModInfo) getLocalModulePath(path string) LnsAny {
    for _mod, _ver := range( self.moduleMap.Items ) {
        mod := _mod.(string)
        ver := _ver.(string)
        if Lns_car(Lns_getVM().String_find(path,mod, 1, true)) == 1{
            var relativeName string
            relativeName = Lns_getVM().String_sub(path,len(mod) + 2, nil)
            {
                _replacePath := self.replaceMap.Items[mod]
                if _replacePath != nil {
                    replacePath := _replacePath.(string)
                    return Util_pathJoin(replacePath, relativeName)
                }
            }
            var gomod string
            gomod = ""
            for _, _aChar := range( NewLnsList(Lns_getVM().String_byte(mod,1, len(mod))).Items ) {
                aChar := _aChar
                if aChar != nil{
                    aChar_560 := aChar.(LnsInt)
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( aChar_560 >= 65) &&
                        Lns_GetEnv().SetStackVal( aChar_560 <= 90) ).(bool)){
                        gomod = Lns_getVM().String_format("%s!%c", []LnsAny{gomod, aChar_560 - 65 + 97})
                        
                    } else { 
                        gomod = Lns_getVM().String_format("%s%c", []LnsAny{gomod, aChar_560})
                        
                    }
                }
            }
            gomod = Lns_getVM().String_format("%s@%s", []LnsAny{gomod, ver})
            
            {
                _gopath := Depend_getGOPATH()
                if _gopath != nil {
                    gopath := _gopath.(string)
                    var modpath string
                    modpath = Util_pathJoin(gopath, Lns_getVM().String_format("pkg/mod/%s", []LnsAny{gomod}))
                    return Util_pathJoin(modpath, relativeName)
                }
            }
        }
    }
    return nil
}

// 74: decl @lune.@base.@GoMod.ModInfo.convPath
func (self *GoMod_ModInfo) convPath(mod string,suffix string) string {
    return Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(mod,"^go/", "")).(string),"%.", "/")).(string),":", ".")).(string) + suffix
}

// 78: decl @lune.@base.@GoMod.ModInfo.getProjRootPath
func (self *GoMod_ModInfo) getProjRootPath(mod string,path string)(string, string) {
    var convPath string
    convPath = GoMod_convExp379(Lns_2DDD(Lns_getVM().String_gsub(self.FP.convPath(mod, ".lns"),"github%.com/[^/]+/[^/]+/", "")))
    var projRoot string
    projRoot = Lns_getVM().String_sub(path,1, len(path) - len(convPath))
    if projRoot != "/"{
        projRoot = GoMod_convExp417(Lns_2DDD(Lns_getVM().String_gsub(projRoot,"/$", "")))
        
    }
    path = Util_parentPath(path)
    
    var modList *LnsList
    modList = Util_splitStr(mod, "[^%.]+")
    var startIndex LnsInt
    startIndex = modList.Len()
    {
        var _from509 LnsInt = 1
        var _to509 LnsInt = modList.Len()
        for _work509 := _from509; _work509 <= _to509; _work509++ {
            modIndex := _work509
            if Depend_existFile(Util_pathJoin(path, "lune.js")){
                startIndex = modIndex
                
                break
            }
            if path == projRoot{
                startIndex = modIndex
                
                break
            }
            path = Util_parentPath(path)
            
        }
    }
    var convMod string
    convMod = ""
    {
        var _from562 LnsInt = modList.Len() - startIndex + 1
        var _to562 LnsInt = modList.Len()
        for _work562 := _from562; _work562 <= _to562; _work562++ {
            index := _work562
            if convMod != ""{
                convMod = Lns_getVM().String_format("%s.", []LnsAny{convMod})
                
            }
            convMod = convMod + modList.GetAt(index).(string)
            
        }
    }
    return path, convMod
}

// 118: decl @lune.@base.@GoMod.ModInfo.convLocalModulePath
func (self *GoMod_ModInfo) ConvLocalModulePath(mod string,suffix string) LnsAny {
    __func__ := "@lune.@base.@GoMod.ModInfo.convLocalModulePath"
    if Lns_op_not(Lns_car(Lns_getVM().String_find(mod,"^go/", nil, nil))){
        return GoMod_GoModResult__NotGo_Obj
    }
    var workMod string
    workMod = self.FP.convPath(mod, suffix)
    {
        __exp := self.path2modProjInfo.Items[workMod]
        if __exp != nil {
            _exp := __exp.(*GoMod_ModProjInfo)
            return &GoMod_GoModResult__Found{_exp}
        }
    }
    var pathList *LnsList
    pathList = NewLnsList([]LnsAny{})
    {
        __exp := self.FP.getLocalModulePath(workMod)
        if __exp != nil {
            _exp := __exp.(string)
            pathList.Insert(_exp)
        }
    }
    pathList.Insert(Util_pathJoin("vendor", workMod))
    for _, _path := range( pathList.Items ) {
        path := _path.(string)
        if Depend_existFile(path){
            var projRoot string
            var convMod string
            projRoot,convMod = self.FP.getProjRootPath(mod, path)
            var projInfo *GoMod_ModProjInfo
            projInfo = NewGoMod_ModProjInfo(path, projRoot, convMod)
            self.path2modProjInfo.Set(workMod,projInfo)
            return &GoMod_GoModResult__Found{projInfo}
        } else { 
            Log_log(Log_Level__Log, __func__, 142, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("not found %s", []LnsAny{path})
            }))
            
        }
    }
    return GoMod_GoModResult__NotFound_Obj
}

// 159: decl @lune.@base.@GoMod.ModInfo.getLuaModulePath
func (self *GoMod_ModInfo) GetLuaModulePath(mod string) string {
    var info *GoMod_ModProjInfo
    switch _exp827 := self.FP.ConvLocalModulePath(mod, ".lns").(type) {
    case *GoMod_GoModResult__NotGo:
        return mod
    case *GoMod_GoModResult__NotFound:
        return mod
    case *GoMod_GoModResult__Found:
    workInfo := _exp827.Val1
        info = workInfo
        
    }
    return info.FP.Get_mod()
}


func Lns_GoMod_init() {
    if init_GoMod { return }
    init_GoMod = true
    GoMod__mod__ = "@lune.@base.@GoMod"
    Lns_InitMod()
    Lns_Option_init()
    Lns_Types_init()
    Lns_Util_init()
    Lns_LuaVer_init()
    Lns_Depend_init()
    Lns_Log_init()
}
func init() {
    init_GoMod = false
}
