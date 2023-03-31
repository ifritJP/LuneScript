// This code is transcompiled by LuneScript.
package lns
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Util bool
var Util__mod__ string
var Util_validLog bool
// for 11
func Util_convExp0_34(arg1 []LnsAny) []LnsAny {
    return Lns_2DDD( arg1[0:])
}
// 5: decl @lns.@Util.setLog
func Util_setLog(_env *LnsEnv, valid bool) {
    Util_validLog = valid
}

// 9: decl @lns.@Util.log
func Util_log(_env *LnsEnv, ddd []LnsAny) {
    if Util_validLog{
        Lns_print(ddd)
    }
}

func Lns_Util_init(_env *LnsEnv) {
    if init_Util { return }
    init_Util = true
    Util__mod__ = "@lns.@Util"
    Lns_InitMod()
    Util_validLog = false
}
func init() {
    init_Util = false
}
