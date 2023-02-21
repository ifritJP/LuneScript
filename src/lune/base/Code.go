// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Code bool
var Code__mod__ string
// decl enum -- ID 
type Code_ID = LnsInt
const Code_ID__nothing_define_abbr = 0
var Code_IDList_ = NewLnsList( []LnsAny {
  Code_ID__nothing_define_abbr,
})
func Code_ID_get__allList(_env *LnsEnv) *LnsList{
    return Code_IDList_
}
var Code_IDMap_ = map[LnsInt]string {
  Code_ID__nothing_define_abbr: "ID.nothing_define_abbr",
}
func Code_ID__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Code_IDMap_[arg1]; ok { return arg1 }
    return nil
}

func Code_ID_getTxt(arg1 LnsInt) string {
    return Code_IDMap_[arg1];
}
// 31: decl @lune.@base.@Code.format
func Code_format(_env *LnsEnv, id LnsInt,mess string) string {
    return _env.GetVM().String_format("%05d:%s", Lns_2DDD(id, mess))
}

// 35: decl @lune.@base.@Code.isMessageOf
func Code_isMessageOf(_env *LnsEnv, id LnsInt,mess string) bool {
    var pat string
    pat = _env.GetVM().String_format("^%05d:", Lns_2DDD(id))
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(mess,pat, nil, nil))){
        return true
    }
    return false
}

func Lns_Code_init(_env *LnsEnv) {
    if init_Code { return }
    init_Code = true
    Code__mod__ = "@lune.@base.@Code"
    Lns_InitMod()
}
func init() {
    init_Code = false
}
