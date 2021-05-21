// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Meta bool
var Meta__mod__ string
// declaration Class -- _MetaInfo
type Meta__MetaInfoMtd interface {
}
type Meta__MetaInfo struct {
    G__formatVersion string
    G__enableTest bool
    G__buildId string
    G__lazyModuleList *LnsMap
    G__typeId2ClassInfoMap *LnsMap
    G__typeInfoList *LnsMap
    G__varName2InfoMap *LnsMap
    G__moduleTypeId LnsInt
    G__moduleSymbolKind LnsInt
    G__moduleMutable bool
    G__dependModuleMap *LnsMap
    G__dependIdMap *LnsMap
    G__macroName2InfoMap *LnsMap
    G__hasTest bool
    G__subModuleMap *LnsMap
    G__moduleHierarchy *LnsMap
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
func NewMeta__MetaInfo(_env *LnsEnv, arg1 string, arg2 bool, arg3 string, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsMap, arg8 LnsInt, arg9 LnsInt, arg10 bool, arg11 *LnsMap, arg12 *LnsMap, arg13 *LnsMap, arg14 bool, arg15 *LnsMap, arg16 *LnsMap) *Meta__MetaInfo {
    obj := &Meta__MetaInfo{}
    obj.FP = obj
    obj.InitMeta__MetaInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16)
    return obj
}
func (self *Meta__MetaInfo) InitMeta__MetaInfo(_env *LnsEnv, arg1 string, arg2 bool, arg3 string, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsMap, arg8 LnsInt, arg9 LnsInt, arg10 bool, arg11 *LnsMap, arg12 *LnsMap, arg13 *LnsMap, arg14 bool, arg15 *LnsMap, arg16 *LnsMap) {
    self.G__formatVersion = arg1
    self.G__enableTest = arg2
    self.G__buildId = arg3
    self.G__lazyModuleList = arg4
    self.G__typeId2ClassInfoMap = arg5
    self.G__typeInfoList = arg6
    self.G__varName2InfoMap = arg7
    self.G__moduleTypeId = arg8
    self.G__moduleSymbolKind = arg9
    self.G__moduleMutable = arg10
    self.G__dependModuleMap = arg11
    self.G__dependIdMap = arg12
    self.G__macroName2InfoMap = arg13
    self.G__hasTest = arg14
    self.G__subModuleMap = arg15
    self.G__moduleHierarchy = arg16
}

func Lns_Meta_init(_env *LnsEnv) {
    if init_Meta { return }
    init_Meta = true
    Meta__mod__ = "@lune.@base.@Meta"
    Lns_InitMod()
}
func init() {
    init_Meta = false
}
