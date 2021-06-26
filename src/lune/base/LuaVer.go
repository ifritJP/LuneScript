// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_LuaVer bool
var LuaVer__mod__ string
// decl enum -- BitOp 
type LuaVer_BitOp = LnsInt
const LuaVer_BitOp__Cant = 2
const LuaVer_BitOp__HasMod = 1
const LuaVer_BitOp__HasOp = 0
var LuaVer_BitOpList_ = NewLnsList( []LnsAny {
  LuaVer_BitOp__HasOp,
  LuaVer_BitOp__HasMod,
  LuaVer_BitOp__Cant,
})
func LuaVer_BitOp_get__allList(_env *LnsEnv) *LnsList{
    return LuaVer_BitOpList_
}
var LuaVer_BitOpMap_ = map[LnsInt]string {
  LuaVer_BitOp__Cant: "BitOp.Cant",
  LuaVer_BitOp__HasMod: "BitOp.HasMod",
  LuaVer_BitOp__HasOp: "BitOp.HasOp",
}
func LuaVer_BitOp__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := LuaVer_BitOpMap_[arg1]; ok { return arg1 }
    return nil
}

func LuaVer_BitOp_getTxt(arg1 LnsInt) string {
    return LuaVer_BitOpMap_[arg1];
}
// decl enum -- VerKind 
type LuaVer_VerKind = LnsInt
const LuaVer_VerKind__v51 = 51
const LuaVer_VerKind__v52 = 52
const LuaVer_VerKind__v53 = 53
var LuaVer_VerKindList_ = NewLnsList( []LnsAny {
  LuaVer_VerKind__v51,
  LuaVer_VerKind__v52,
  LuaVer_VerKind__v53,
})
func LuaVer_VerKind_get__allList(_env *LnsEnv) *LnsList{
    return LuaVer_VerKindList_
}
var LuaVer_VerKindMap_ = map[LnsInt]string {
  LuaVer_VerKind__v51: "VerKind.v51",
  LuaVer_VerKind__v52: "VerKind.v52",
  LuaVer_VerKind__v53: "VerKind.v53",
}
func LuaVer_VerKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := LuaVer_VerKindMap_[arg1]; ok { return arg1 }
    return nil
}

func LuaVer_VerKind_getTxt(arg1 LnsInt) string {
    return LuaVer_VerKindMap_[arg1];
}
var LuaVer_ver51 *LuaVer_LuaVerInfo
var LuaVer_ver52 *LuaVer_LuaVerInfo
var LuaVer_ver53 *LuaVer_LuaVerInfo
var LuaVer_kind2verMap *LnsMap
var LuaVer_curVer LnsAny
// 79: decl @lune.@base.@LuaVer.setCurVer
func LuaVer_setCurVer(_env *LnsEnv, ver LnsInt) {
    var verKind LnsInt
    
    {
        _verKind := LuaVer_VerKind__from(_env, ver)
        if _verKind == nil{
            return 
        } else {
            verKind = _verKind.(LnsInt)
        }
    }
    LuaVer_curVer = Lns_unwrap( LuaVer_kind2verMap.Get(verKind)).(*LuaVer_LuaVerInfo)
}

// 85: decl @lune.@base.@LuaVer.getCurVer
func LuaVer_getCurVer(_env *LnsEnv) *LuaVer_LuaVerInfo {
    return Lns_unwrap( LuaVer_curVer).(*LuaVer_LuaVerInfo)
}

// 55: decl @lune.@base.@LuaVer.LuaVerInfo.isSupport
func (self *LuaVer_LuaVerInfo) IsSupport(_env *LnsEnv, symbol string) bool {
    return Lns_op_not(self.noSupportSymMap.Has(symbol))
}
// 59: decl @lune.@base.@LuaVer.LuaVerInfo.getLoadCode
func (self *LuaVer_LuaVerInfo) GetLoadCode(_env *LnsEnv) string {
    return LuaMod_getCode(_env, self.loadKind)
}
// declaration Class -- LuaVerInfo
type LuaVer_LuaVerInfoMtd interface {
    GetLoadCode(_env *LnsEnv) string
    Get_canFormStem2Str(_env *LnsEnv) bool
    Get_canUseMetaGc(_env *LnsEnv) bool
    Get_hasBitOp(_env *LnsEnv) LnsInt
    Get_hasSearchPath(_env *LnsEnv) bool
    Get_hasTableUnpack(_env *LnsEnv) bool
    Get_loadStrFuncName(_env *LnsEnv) string
    Get_verKind(_env *LnsEnv) LnsInt
    IsSupport(_env *LnsEnv, arg1 string) bool
}
type LuaVer_LuaVerInfo struct {
    verKind LnsInt
    hasBitOp LnsInt
    hasTableUnpack bool
    canFormStem2Str bool
    hasSearchPath bool
    loadStrFuncName string
    canUseMetaGc bool
    loadKind LnsInt
    noSupportSymMap *LnsSet
    FP LuaVer_LuaVerInfoMtd
}
func LuaVer_LuaVerInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*LuaVer_LuaVerInfo).FP
}
type LuaVer_LuaVerInfoDownCast interface {
    ToLuaVer_LuaVerInfo() *LuaVer_LuaVerInfo
}
func LuaVer_LuaVerInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(LuaVer_LuaVerInfoDownCast)
    if ok { return work.ToLuaVer_LuaVerInfo() }
    return nil
}
func (obj *LuaVer_LuaVerInfo) ToLuaVer_LuaVerInfo() *LuaVer_LuaVerInfo {
    return obj
}
func NewLuaVer_LuaVerInfo(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 bool, arg4 bool, arg5 bool, arg6 string, arg7 bool, arg8 LnsInt, arg9 *LnsSet) *LuaVer_LuaVerInfo {
    obj := &LuaVer_LuaVerInfo{}
    obj.FP = obj
    obj.InitLuaVer_LuaVerInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *LuaVer_LuaVerInfo) InitLuaVer_LuaVerInfo(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 bool, arg4 bool, arg5 bool, arg6 string, arg7 bool, arg8 LnsInt, arg9 *LnsSet) {
    self.verKind = arg1
    self.hasBitOp = arg2
    self.hasTableUnpack = arg3
    self.canFormStem2Str = arg4
    self.hasSearchPath = arg5
    self.loadStrFuncName = arg6
    self.canUseMetaGc = arg7
    self.loadKind = arg8
    self.noSupportSymMap = arg9
}
func (self *LuaVer_LuaVerInfo) Get_verKind(_env *LnsEnv) LnsInt{ return self.verKind }
func (self *LuaVer_LuaVerInfo) Get_hasBitOp(_env *LnsEnv) LnsInt{ return self.hasBitOp }
func (self *LuaVer_LuaVerInfo) Get_hasTableUnpack(_env *LnsEnv) bool{ return self.hasTableUnpack }
func (self *LuaVer_LuaVerInfo) Get_canFormStem2Str(_env *LnsEnv) bool{ return self.canFormStem2Str }
func (self *LuaVer_LuaVerInfo) Get_hasSearchPath(_env *LnsEnv) bool{ return self.hasSearchPath }
func (self *LuaVer_LuaVerInfo) Get_loadStrFuncName(_env *LnsEnv) string{ return self.loadStrFuncName }
func (self *LuaVer_LuaVerInfo) Get_canUseMetaGc(_env *LnsEnv) bool{ return self.canUseMetaGc }

func Lns_LuaVer_init(_env *LnsEnv) {
    if init_LuaVer { return }
    init_LuaVer = true
    LuaVer__mod__ = "@lune.@base.@LuaVer"
    Lns_InitMod()
    Lns_LuaMod_init(_env)
    LuaVer_ver51 = NewLuaVer_LuaVerInfo(_env, LuaVer_VerKind__v51, LuaVer_BitOp__Cant, false, false, false, "loadstring51", false, LuaMod_CodeKind__LoadStr51, NewLnsSet([]LnsAny{"package.searchpath"}))
    LuaVer_ver52 = NewLuaVer_LuaVerInfo(_env, LuaVer_VerKind__v52, LuaVer_BitOp__HasMod, true, true, true, "loadstring52", true, LuaMod_CodeKind__LoadStr52, NewLnsSet([]LnsAny{}))
    LuaVer_ver53 = NewLuaVer_LuaVerInfo(_env, LuaVer_VerKind__v53, LuaVer_BitOp__HasOp, true, true, true, "loadstring52", true, LuaMod_CodeKind__LoadStr52, NewLnsSet([]LnsAny{}))
    LuaVer_kind2verMap = NewLnsMap( map[LnsAny]LnsAny{LuaVer_VerKind__v51:LuaVer_ver51,LuaVer_VerKind__v52:LuaVer_ver52,LuaVer_VerKind__v53:LuaVer_ver53,})
    LuaVer_curVer = nil
}
func init() {
    init_LuaVer = false
}
