// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_TestingMacro bool
var TestingMacro__mod__ string
func Lns_TestingMacro_init(_env *LnsEnv) {
    if init_TestingMacro { return }
    init_TestingMacro = true
    TestingMacro__mod__ = "@lune.@base.@TestingMacro"
    Lns_InitMod()
}
func init() {
    init_TestingMacro = false
}
