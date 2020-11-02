// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_Types bool
var Types__mod__ string
// decl enum -- Conv 
const Types_Conv__C = 0
const Types_Conv__Go = 1
var Types_ConvList_ = NewLnsList( []LnsAny {
  Types_Conv__C,
  Types_Conv__Go,
})
func Types_Conv_get__allList() *LnsList{
    return Types_ConvList_
}
var Types_ConvMap_ = map[LnsInt]string {
  Types_Conv__C: "Conv.C",
  Types_Conv__Go: "Conv.Go",
}
func Types_Conv__from(arg1 LnsInt) LnsAny{
    if _, ok := Types_ConvMap_[arg1]; ok { return arg1 }
    return nil
}

func Types_Conv_getTxt(arg1 LnsInt) string {
    return Types_ConvMap_[arg1];
}
// decl enum -- CheckingUptodateMode 
const Types_CheckingUptodateMode__Force = "force"
const Types_CheckingUptodateMode__Normal = "none"
const Types_CheckingUptodateMode__Touch = "touch"
var Types_CheckingUptodateModeList_ = NewLnsList( []LnsAny {
  Types_CheckingUptodateMode__Force,
  Types_CheckingUptodateMode__Normal,
  Types_CheckingUptodateMode__Touch,
})
func Types_CheckingUptodateMode_get__allList() *LnsList{
    return Types_CheckingUptodateModeList_
}
var Types_CheckingUptodateModeMap_ = map[string]string {
  Types_CheckingUptodateMode__Force: "CheckingUptodateMode.Force",
  Types_CheckingUptodateMode__Normal: "CheckingUptodateMode.Normal",
  Types_CheckingUptodateMode__Touch: "CheckingUptodateMode.Touch",
}
func Types_CheckingUptodateMode__from(arg1 string) LnsAny{
    if _, ok := Types_CheckingUptodateModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Types_CheckingUptodateMode_getTxt(arg1 string) string {
    return Types_CheckingUptodateModeMap_[arg1];
}
// declaration Class -- TransCtrlInfo
type Types_TransCtrlInfoMtd interface {
}
type Types_TransCtrlInfo struct {
    WarningShadowing bool
    CompatComment bool
    CheckingDefineAbbr bool
    StopByWarning bool
    UptodateMode string
    ValidLuaval bool
    FP Types_TransCtrlInfoMtd
}
func Types_TransCtrlInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Types_TransCtrlInfo).FP
}
type Types_TransCtrlInfoDownCast interface {
    ToTypes_TransCtrlInfo() *Types_TransCtrlInfo
}
func Types_TransCtrlInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Types_TransCtrlInfoDownCast)
    if ok { return work.ToTypes_TransCtrlInfo() }
    return nil
}
func (obj *Types_TransCtrlInfo) ToTypes_TransCtrlInfo() *Types_TransCtrlInfo {
    return obj
}
func NewTypes_TransCtrlInfo(arg1 bool, arg2 bool, arg3 bool, arg4 bool, arg5 string, arg6 bool) *Types_TransCtrlInfo {
    obj := &Types_TransCtrlInfo{}
    obj.FP = obj
    obj.InitTypes_TransCtrlInfo(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Types_TransCtrlInfo) InitTypes_TransCtrlInfo(arg1 bool, arg2 bool, arg3 bool, arg4 bool, arg5 string, arg6 bool) {
    self.WarningShadowing = arg1
    self.CompatComment = arg2
    self.CheckingDefineAbbr = arg3
    self.StopByWarning = arg4
    self.UptodateMode = arg5
    self.ValidLuaval = arg6
}
// 53: decl @lune.@base.@Types.TransCtrlInfo.create_normal
func Types_TransCtrlInfo_create_normal() *Types_TransCtrlInfo {
    return NewTypes_TransCtrlInfo(false, false, true, false, Types_CheckingUptodateMode__Touch, false)
}


func Lns_Types_init() {
    if init_Types { return }
    init_Types = true
    Types__mod__ = "@lune.@base.@Types"
    Lns_InitMod()
}
func init() {
    init_Types = false
}
