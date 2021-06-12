// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Ver bool
var Ver__mod__ string
var Ver_version string
var Ver_metaVersion string
var Ver_luaModVersion LnsInt
func Lns_Ver_init(_env *LnsEnv) {
    if init_Ver { return }
    init_Ver = true
    Ver__mod__ = "@lune.@base.@Ver"
    Lns_InitMod()
    Ver_version = "1.3.0"
    Ver_metaVersion = "1.0.138"
    Ver_luaModVersion = 4
}
func init() {
    init_Ver = false
}
