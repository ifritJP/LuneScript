// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_Code bool
var Code__mod__ string
// decl enum -- ID 
const Code_ID__nothing_define_abbr = 0
var Code_IDList_ = NewLnsList( []LnsAny {
  Code_ID__nothing_define_abbr,
})
func Code_ID_get__allList() *LnsList{
    return Code_IDList_
}
var Code_IDMap_ = map[LnsInt]string {
  Code_ID__nothing_define_abbr: "ID.nothing_define_abbr",
}
func Code_ID__from(arg1 LnsInt) LnsAny{
    if _, ok := Code_IDMap_[arg1]; ok { return arg1 }
    return nil
}

func Code_ID_getTxt(arg1 LnsInt) string {
    return Code_IDMap_[arg1];
}
// 29: decl @lune.@base.@Code.format
func Code_format(id LnsInt,mess string) string {
    return Lns_getVM().String_format("%05d:%s", []LnsAny{id, mess})
}

// 33: decl @lune.@base.@Code.isMessageOf
func Code_isMessageOf(id LnsInt,mess string) bool {
    var pat string
    pat = Lns_getVM().String_format("^%05d:", []LnsAny{id})
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(mess,pat, nil, nil))){
        return true
    }
    return false
}

func Lns_Code_init() {
    if init_Code { return }
    init_Code = true
    Code__mod__ = "@lune.@base.@Code"
    Lns_InitMod()
}
func init() {
    init_Code = false
}
