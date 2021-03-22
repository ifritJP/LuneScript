// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Util bool
var Util__mod__ string
func Lns_Util_init() {
    if init_Util { return }
    init_Util = true
    Util__mod__ = "@lune.@Util"
    Lns_InitMod()
}
func init() {
    init_Util = false
}
