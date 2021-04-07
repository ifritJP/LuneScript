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
func GoMod_BlockKind_get__allList_1050_() *LnsList{
    return GoMod_BlockKindList_
}
var GoMod_BlockKindMap_ = map[LnsInt]string {
  GoMod_BlockKind__None: "BlockKind.None",
  GoMod_BlockKind__Replace: "BlockKind.Replace",
  GoMod_BlockKind__Require: "BlockKind.Require",
}
func GoMod_BlockKind__from_1043_(arg1 LnsInt) LnsAny{
    if _, ok := GoMod_BlockKindMap_[arg1]; ok { return arg1 }
    return nil
}

func GoMod_BlockKind_getTxt(arg1 LnsInt) string {
    return GoMod_BlockKindMap_[arg1];
}
// for 67
func GoMod_convExp556(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 73
func GoMod_convExp366(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 48: decl @lune.@base.@GoMod.getReplace
func GoMod_getReplace_1060_(_map *LnsMap,tokenList *LnsList,modIndex LnsInt) {
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

// 60: decl @lune.@base.@GoMod.getGoMap
func GoMod_getGoMap(option *Option_Option) *GoMod_ModInfo {
    var requireMap *LnsMap
    requireMap = NewLnsMap( map[LnsAny]LnsAny{})
    var replaceMap *LnsMap
    replaceMap = NewLnsMap( map[LnsAny]LnsAny{})
    var modInfo *GoMod_ModInfo
    modInfo = NewGoMod_ModInfo(requireMap, replaceMap)
    {
        _file := GoMod_convExp556(Lns_2DDD(Lns_io_open("go.mod", nil)))
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
                trimedLine = GoMod_convExp366(Lns_2DDD(Lns_getVM().String_gsub(line,"^%s", "")))
                var tokenList *LnsList
                tokenList = Util_splitStr(trimedLine, "[^%s]+")
                if _switch552 := inBlock; _switch552 == GoMod_BlockKind__Require {
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^%)", nil, nil))){
                        inBlock = GoMod_BlockKind__None
                        
                    } else { 
                        requireMap.Set(tokenList.GetAt(1).(string),tokenList.GetAt(2).(string))
                    }
                } else if _switch552 == GoMod_BlockKind__Replace {
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^%)", nil, nil))){
                        inBlock = GoMod_BlockKind__None
                        
                    } else { 
                        GoMod_getReplace_1060_(replaceMap, tokenList, 1)
                    }
                } else if _switch552 == GoMod_BlockKind__None {
                    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^require%s+[^%(]", nil, nil))){
                        if tokenList.Len() == 3{
                            requireMap.Set(tokenList.GetAt(2).(string),tokenList.GetAt(3).(string))
                        }
                    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^require%s+%(", nil, nil))){
                        inBlock = GoMod_BlockKind__Require
                        
                    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^replace%s+[^%(]", nil, nil))){
                        GoMod_getReplace_1060_(replaceMap, tokenList, 2)
                    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(line,"^replace%s+%(", nil, nil))){
                        inBlock = GoMod_BlockKind__Require
                        
                    }
                }
            }
        }
    }
    return modInfo
}

// declaration Class -- ModInfo
type GoMod_ModInfoMtd interface {
    GetModulePath(arg1 string) LnsAny
    Get_moduleMap() *LnsMap
    Get_replaceMap() *LnsMap
}
type GoMod_ModInfo struct {
    moduleMap *LnsMap
    replaceMap *LnsMap
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
func (self *GoMod_ModInfo) InitGoMod_ModInfo(arg1 *LnsMap, arg2 *LnsMap) {
    self.moduleMap = arg1
    self.replaceMap = arg2
}
func (self *GoMod_ModInfo) Get_moduleMap() *LnsMap{ return self.moduleMap }
func (self *GoMod_ModInfo) Get_replaceMap() *LnsMap{ return self.replaceMap }
// 12: decl @lune.@base.@GoMod.ModInfo.getModulePath
func (self *GoMod_ModInfo) GetModulePath(path string) LnsAny {
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
                    aChar_533 := aChar.(LnsInt)
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( aChar_533 >= 65) &&
                        Lns_GetEnv().SetStackVal( aChar_533 <= 90) ).(bool)){
                        gomod = Lns_getVM().String_format("%s!%c", []LnsAny{gomod, aChar_533 - 65 + 97})
                        
                    } else { 
                        gomod = Lns_getVM().String_format("%s%c", []LnsAny{gomod, aChar_533})
                        
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
