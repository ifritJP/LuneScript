// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/lune/base/runtime_go"
var init_Meta bool
var Meta__mod__ string
// declaration Class -- _MetaInfo
type Meta__MetaInfoMtd interface {
}
type Meta__MetaInfo struct {
    __formatVersion string
    __enableTest bool
    __buildId string
    __lazyModuleList *LnsMap
    __typeId2ClassInfoMap *LnsMap
    __typeInfoList *LnsMap
    __varName2InfoMap *LnsMap
    __moduleTypeId LnsInt
    __moduleSymbolKind LnsInt
    __moduleMutable bool
    __dependModuleMap *LnsMap
    __dependIdMap *LnsMap
    __macroName2InfoMap *LnsMap
    __hasTest bool
    __subModuleMap *LnsMap
    FP Meta__MetaInfoMtd
}
func Meta__MetaInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Meta__MetaInfo).FP
}
type Meta__MetaInfoDownCast interface {
    ToMeta__MetaInfo() *Meta__MetaInfo
}
func Meta__MetaInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Meta__MetaInfoDownCast)
    if ok { return work.ToMeta__MetaInfo() }
    return nil
}
func (obj *Meta__MetaInfo) ToMeta__MetaInfo() *Meta__MetaInfo {
    return obj
}
func NewMeta__MetaInfo(arg1 string, arg2 bool, arg3 string, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsMap, arg8 LnsInt, arg9 LnsInt, arg10 bool, arg11 *LnsMap, arg12 *LnsMap, arg13 *LnsMap, arg14 bool, arg15 *LnsMap) *Meta__MetaInfo {
    obj := &Meta__MetaInfo{}
    obj.FP = obj
    obj.InitMeta__MetaInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15)
    return obj
}
func (self *Meta__MetaInfo) InitMeta__MetaInfo(arg1 string, arg2 bool, arg3 string, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsMap, arg8 LnsInt, arg9 LnsInt, arg10 bool, arg11 *LnsMap, arg12 *LnsMap, arg13 *LnsMap, arg14 bool, arg15 *LnsMap) {
    self.__formatVersion = arg1
    self.__enableTest = arg2
    self.__buildId = arg3
    self.__lazyModuleList = arg4
    self.__typeId2ClassInfoMap = arg5
    self.__typeInfoList = arg6
    self.__varName2InfoMap = arg7
    self.__moduleTypeId = arg8
    self.__moduleSymbolKind = arg9
    self.__moduleMutable = arg10
    self.__dependModuleMap = arg11
    self.__dependIdMap = arg12
    self.__macroName2InfoMap = arg13
    self.__hasTest = arg14
    self.__subModuleMap = arg15
}

func Lns_Meta_init() {
    if init_Meta { return }
    init_Meta = true
    Meta__mod__ = "@lune.@base.@Meta"
    Lns_InitMod()
}
func init() {
    init_Meta = false
}
