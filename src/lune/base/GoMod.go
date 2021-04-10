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
func GoMod_BlockKind_get__allList_1083_() *LnsList{
    return GoMod_BlockKindList_
}
var GoMod_BlockKindMap_ = map[LnsInt]string {
  GoMod_BlockKind__None: "BlockKind.None",
  GoMod_BlockKind__Replace: "BlockKind.Replace",
  GoMod_BlockKind__Require: "BlockKind.Require",
}
func GoMod_BlockKind__from_1076_(arg1 LnsInt) LnsAny{
    if _, ok := GoMod_BlockKindMap_[arg1]; ok { return arg1 }
    return nil
}

func GoMod_BlockKind_getTxt(arg1 LnsInt) string {
    return GoMod_BlockKindMap_[arg1];
}
// decl alge -- GoModResult
type GoMod_GoModResult = LnsAny
type GoMod_GoModResult__Found struct{
Val1 string
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
// for 170
func GoMod_convExp1010(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 117
func GoMod_convExp538(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 114
func GoMod_convExp500(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 176
func GoMod_convExp820(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 151: decl @lune.@base.@GoMod.getReplace
func GoMod_getReplace_1093_(_map *LnsMap,tokenList *LnsList,modIndex LnsInt) {
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

// 163: decl @lune.@base.@GoMod.getGoMap
func GoMod_getGoMap(option *Option_Option) *GoMod_ModInfo {
    var requireMap *LnsMap
    requireMap = NewLnsMap( map[LnsAny]LnsAny{})
    var replaceMap *LnsMap
    replaceMap = NewLnsMap( map[LnsAny]LnsAny{})
    var modInfo *GoMod_ModInfo
    modInfo = NewGoMod_ModInfo(requireMap, replaceMap)
    {
        _file := GoMod_convExp1010(Lns_2DDD(Lns_io_open("go.mod", nil)))
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
                trimedLine = GoMod_convExp820(Lns_2DDD(Lns_getVM().String_gsub(line,"^%s", "")))
                var tokenList *LnsList
                tokenList = Util_splitStr(trimedLine, "[^%s]+")
                if _switch1006 := inBlock; _switch1006 == GoMod_BlockKind__Require {
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^%)", nil, nil))){
                        inBlock = GoMod_BlockKind__None
                        
                    } else { 
                        requireMap.Set(tokenList.GetAt(1).(string),tokenList.GetAt(2).(string))
                    }
                } else if _switch1006 == GoMod_BlockKind__Replace {
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^%)", nil, nil))){
                        inBlock = GoMod_BlockKind__None
                        
                    } else { 
                        GoMod_getReplace_1093_(replaceMap, tokenList, 1)
                    }
                } else if _switch1006 == GoMod_BlockKind__None {
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^require%s+[^%(]", nil, nil))){
                        if tokenList.Len() == 3{
                            requireMap.Set(tokenList.GetAt(2).(string),tokenList.GetAt(3).(string))
                        }
                    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^require%s+%(", nil, nil))){
                        inBlock = GoMod_BlockKind__Require
                        
                    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^replace%s+[^%(]", nil, nil))){
                        GoMod_getReplace_1093_(replaceMap, tokenList, 2)
                    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^replace%s+%(", nil, nil))){
                        inBlock = GoMod_BlockKind__Replace
                        
                    }
                }
            }
        }
    }
    return modInfo
}

// declaration Class -- ModInfo
type GoMod_ModInfoMtd interface {
    ConvLocalModulePath(arg1 string, arg2 string) LnsAny
    convPath(arg1 string, arg2 string) string
    getLocalModulePath(arg1 string) LnsAny
    GetLuaModulePath(arg1 string) string
    Get_moduleMap() *LnsMap
    Get_replaceMap() *LnsMap
}
type GoMod_ModInfo struct {
    moduleMap *LnsMap
    replaceMap *LnsMap
    workPath2convPath *LnsMap
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
// 21: DeclConstr
func (self *GoMod_ModInfo) InitGoMod_ModInfo(moduleMap *LnsMap,replaceMap *LnsMap) {
    self.moduleMap = moduleMap
    
    self.replaceMap = replaceMap
    
    self.workPath2convPath = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 27: decl @lune.@base.@GoMod.ModInfo.getLocalModulePath
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
                    aChar_546 := aChar.(LnsInt)
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( aChar_546 >= 65) &&
                        Lns_GetEnv().SetStackVal( aChar_546 <= 90) ).(bool)){
                        gomod = Lns_getVM().String_format("%s!%c", []LnsAny{gomod, aChar_546 - 65 + 97})
                        
                    } else { 
                        gomod = Lns_getVM().String_format("%s%c", []LnsAny{gomod, aChar_546})
                        
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

// 56: decl @lune.@base.@GoMod.ModInfo.convPath
func (self *GoMod_ModInfo) convPath(mod string,suffix string) string {
    return Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(mod,"^go/", "")).(string),"%.", "/")).(string),":", ".")).(string) + suffix
}

// 64: decl @lune.@base.@GoMod.ModInfo.convLocalModulePath
func (self *GoMod_ModInfo) ConvLocalModulePath(mod string,suffix string) LnsAny {
    if Lns_op_not(Lns_car(Lns_getVM().String_find(mod,"^go/", nil, nil))){
        return GoMod_GoModResult__NotGo_Obj
    }
    var workMod string
    workMod = self.FP.convPath(mod, suffix)
    {
        __exp := self.workPath2convPath.Items[workMod]
        if __exp != nil {
            _exp := __exp.(string)
            return &GoMod_GoModResult__Found{_exp}
        }
    }
    var pathList *LnsList
    pathList = NewLnsList([]LnsAny{Util_pathJoin("vendor", workMod)})
    {
        __exp := self.FP.getLocalModulePath(workMod)
        if __exp != nil {
            _exp := __exp.(string)
            pathList.Insert(_exp)
        }
    }
    for _, _path := range( pathList.Items ) {
        path := _path.(string)
        if Depend_existFile(path){
            self.workPath2convPath.Set(workMod,path)
            return &GoMod_GoModResult__Found{path}
        }
    }
    return GoMod_GoModResult__NotFound_Obj
}

// 100: decl @lune.@base.@GoMod.ModInfo.getLuaModulePath
func (self *GoMod_ModInfo) GetLuaModulePath(mod string) string {
    var path string
    switch _exp479 := self.FP.ConvLocalModulePath(mod, ".lns").(type) {
    case *GoMod_GoModResult__NotGo:
        return mod
    case *GoMod_GoModResult__NotFound:
        return mod
    case *GoMod_GoModResult__Found:
    workPath := _exp479.Val1
        path = workPath
        
    }
    var convPath string
    convPath = GoMod_convExp500(Lns_2DDD(Lns_getVM().String_gsub(self.FP.convPath(mod, ".lns"),"github%.com/[^/]+/[^/]+/", "")))
    var projRoot string
    projRoot = Lns_getVM().String_sub(path,1, len(path) - len(convPath))
    if projRoot != "/"{
        projRoot = GoMod_convExp538(Lns_2DDD(Lns_getVM().String_gsub(projRoot,"/$", "")))
        
    }
    path = Util_parentPath(path)
    
    var modList *LnsList
    modList = Util_splitStr(mod, "[^%.]+")
    {
        var _from666 LnsInt = 1
        var _to666 LnsInt = modList.Len()
        for _work666 := _from666; _work666 <= _to666; _work666++ {
            modIndex := _work666
            if Depend_existFile(Util_pathJoin(path, "lune.js")){
                var convMod string
                convMod = ""
                {
                    var _from637 LnsInt = modList.Len() - modIndex + 1
                    var _to637 LnsInt = modList.Len()
                    for _work637 := _from637; _work637 <= _to637; _work637++ {
                        index := _work637
                        if convMod != ""{
                            convMod = Lns_getVM().String_format("%s.", []LnsAny{convMod})
                            
                        }
                        convMod = convMod + modList.GetAt(index).(string)
                        
                    }
                }
                return convMod
            }
            if path == projRoot{
                break
            }
            path = Util_parentPath(path)
            
        }
    }
    return mod
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
}
func init() {
    init_GoMod = false
}
