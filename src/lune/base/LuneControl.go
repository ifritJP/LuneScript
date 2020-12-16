// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/lune/base/runtime_go"
var init_LuneControl bool
var LuneControl__mod__ string
// decl enum -- Code 
const LuneControl_Code__C = "c"
const LuneControl_Code__Go = "go"
const LuneControl_Code__Lua = "lua"
var LuneControl_CodeList_ = NewLnsList( []LnsAny {
  LuneControl_Code__Lua,
  LuneControl_Code__C,
  LuneControl_Code__Go,
})
func LuneControl_Code_get__allList() *LnsList{
    return LuneControl_CodeList_
}
var LuneControl_CodeMap_ = map[string]string {
  LuneControl_Code__C: "Code.C",
  LuneControl_Code__Go: "Code.Go",
  LuneControl_Code__Lua: "Code.Lua",
}
func LuneControl_Code__from(arg1 string) LnsAny{
    if _, ok := LuneControl_CodeMap_[arg1]; ok { return arg1 }
    return nil
}

func LuneControl_Code_getTxt(arg1 string) string {
    return LuneControl_CodeMap_[arg1];
}
type LuneControl_Pragma__default__init struct{
}
var LuneControl_Pragma__default__init_Obj = &LuneControl_Pragma__default__init{}
func (self *LuneControl_Pragma__default__init) GetTxt() string {
return "Pragma.default__init"
}
type LuneControl_Pragma__default__init_old struct{
}
var LuneControl_Pragma__default__init_old_Obj = &LuneControl_Pragma__default__init_old{}
func (self *LuneControl_Pragma__default__init_old) GetTxt() string {
return "Pragma.default__init_old"
}
type LuneControl_Pragma__disable_mut_control struct{
}
var LuneControl_Pragma__disable_mut_control_Obj = &LuneControl_Pragma__disable_mut_control{}
func (self *LuneControl_Pragma__disable_mut_control) GetTxt() string {
return "Pragma.disable_mut_control"
}
type LuneControl_Pragma__ignore_symbol_ struct{
}
var LuneControl_Pragma__ignore_symbol__Obj = &LuneControl_Pragma__ignore_symbol_{}
func (self *LuneControl_Pragma__ignore_symbol_) GetTxt() string {
return "Pragma.ignore_symbol_"
}
type LuneControl_Pragma__limit_conv_code struct{
Val1 *LnsSet
}
func (self *LuneControl_Pragma__limit_conv_code) GetTxt() string {
return "Pragma.limit_conv_code"
}
type LuneControl_Pragma__load__lune_module struct{
}
var LuneControl_Pragma__load__lune_module_Obj = &LuneControl_Pragma__load__lune_module{}
func (self *LuneControl_Pragma__load__lune_module) GetTxt() string {
return "Pragma.load__lune_module"
}
type LuneControl_Pragma__run_async_pipe struct{
}
var LuneControl_Pragma__run_async_pipe_Obj = &LuneControl_Pragma__run_async_pipe{}
func (self *LuneControl_Pragma__run_async_pipe) GetTxt() string {
return "Pragma.run_async_pipe"
}
type LuneControl_Pragma__use_async struct{
}
var LuneControl_Pragma__use_async_Obj = &LuneControl_Pragma__use_async{}
func (self *LuneControl_Pragma__use_async) GetTxt() string {
return "Pragma.use_async"
}
func Lns_LuneControl_init() {
    if init_LuneControl { return }
    init_LuneControl = true
    LuneControl__mod__ = "@lune.@base.@LuneControl"
    Lns_InitMod()
}
func init() {
    init_LuneControl = false
}
