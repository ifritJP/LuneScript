// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_Ast bool
var Ast__mod__ string
// decl enum -- IdType 
const Ast_IdType__Base = 0
const Ast_IdType__Ext = 1
var Ast_IdTypeList_ = NewLnsList( []LnsAny {
  Ast_IdType__Base,
  Ast_IdType__Ext,
})
func Ast_IdType_get__allList() *LnsList{
    return Ast_IdTypeList_
}
var Ast_IdTypeMap_ = map[LnsInt]string {
  Ast_IdType__Base: "IdType.Base",
  Ast_IdType__Ext: "IdType.Ext",
}
func Ast_IdType__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_IdTypeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_IdType_getTxt(arg1 LnsInt) string {
    return Ast_IdTypeMap_[arg1];
}
// decl enum -- AccessMode 
const Ast_AccessMode__Global = 5
const Ast_AccessMode__Local = 4
const Ast_AccessMode__None = 0
const Ast_AccessMode__Pri = 3
const Ast_AccessMode__Pro = 2
const Ast_AccessMode__Pub = 1
var Ast_AccessModeList_ = NewLnsList( []LnsAny {
  Ast_AccessMode__None,
  Ast_AccessMode__Pub,
  Ast_AccessMode__Pro,
  Ast_AccessMode__Pri,
  Ast_AccessMode__Local,
  Ast_AccessMode__Global,
})
func Ast_AccessMode_get__allList() *LnsList{
    return Ast_AccessModeList_
}
var Ast_AccessModeMap_ = map[LnsInt]string {
  Ast_AccessMode__Global: "AccessMode.Global",
  Ast_AccessMode__Local: "AccessMode.Local",
  Ast_AccessMode__None: "AccessMode.None",
  Ast_AccessMode__Pri: "AccessMode.Pri",
  Ast_AccessMode__Pro: "AccessMode.Pro",
  Ast_AccessMode__Pub: "AccessMode.Pub",
}
func Ast_AccessMode__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_AccessModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_AccessMode_getTxt(arg1 LnsInt) string {
    return Ast_AccessModeMap_[arg1];
}
// decl enum -- SymbolKind 
const Ast_SymbolKind__Arg = 5
const Ast_SymbolKind__Fun = 3
const Ast_SymbolKind__Mac = 6
const Ast_SymbolKind__Mbr = 1
const Ast_SymbolKind__Mtd = 2
const Ast_SymbolKind__Typ = 0
const Ast_SymbolKind__Var = 4
var Ast_SymbolKindList_ = NewLnsList( []LnsAny {
  Ast_SymbolKind__Typ,
  Ast_SymbolKind__Mbr,
  Ast_SymbolKind__Mtd,
  Ast_SymbolKind__Fun,
  Ast_SymbolKind__Var,
  Ast_SymbolKind__Arg,
  Ast_SymbolKind__Mac,
})
func Ast_SymbolKind_get__allList() *LnsList{
    return Ast_SymbolKindList_
}
var Ast_SymbolKindMap_ = map[LnsInt]string {
  Ast_SymbolKind__Arg: "SymbolKind.Arg",
  Ast_SymbolKind__Fun: "SymbolKind.Fun",
  Ast_SymbolKind__Mac: "SymbolKind.Mac",
  Ast_SymbolKind__Mbr: "SymbolKind.Mbr",
  Ast_SymbolKind__Mtd: "SymbolKind.Mtd",
  Ast_SymbolKind__Typ: "SymbolKind.Typ",
  Ast_SymbolKind__Var: "SymbolKind.Var",
}
func Ast_SymbolKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_SymbolKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_SymbolKind_getTxt(arg1 LnsInt) string {
    return Ast_SymbolKindMap_[arg1];
}
// decl enum -- SerializeKind 
const Ast_SerializeKind__Alge = 5
const Ast_SerializeKind__Alias = 7
const Ast_SerializeKind__Alternate = 8
const Ast_SerializeKind__Box = 10
const Ast_SerializeKind__DDD = 6
const Ast_SerializeKind__Enum = 4
const Ast_SerializeKind__Ext = 11
const Ast_SerializeKind__Generic = 9
const Ast_SerializeKind__Modifier = 1
const Ast_SerializeKind__Module = 2
const Ast_SerializeKind__Nilable = 0
const Ast_SerializeKind__Normal = 3
var Ast_SerializeKindList_ = NewLnsList( []LnsAny {
  Ast_SerializeKind__Nilable,
  Ast_SerializeKind__Modifier,
  Ast_SerializeKind__Module,
  Ast_SerializeKind__Normal,
  Ast_SerializeKind__Enum,
  Ast_SerializeKind__Alge,
  Ast_SerializeKind__DDD,
  Ast_SerializeKind__Alias,
  Ast_SerializeKind__Alternate,
  Ast_SerializeKind__Generic,
  Ast_SerializeKind__Box,
  Ast_SerializeKind__Ext,
})
func Ast_SerializeKind_get__allList() *LnsList{
    return Ast_SerializeKindList_
}
var Ast_SerializeKindMap_ = map[LnsInt]string {
  Ast_SerializeKind__Alge: "SerializeKind.Alge",
  Ast_SerializeKind__Alias: "SerializeKind.Alias",
  Ast_SerializeKind__Alternate: "SerializeKind.Alternate",
  Ast_SerializeKind__Box: "SerializeKind.Box",
  Ast_SerializeKind__DDD: "SerializeKind.DDD",
  Ast_SerializeKind__Enum: "SerializeKind.Enum",
  Ast_SerializeKind__Ext: "SerializeKind.Ext",
  Ast_SerializeKind__Generic: "SerializeKind.Generic",
  Ast_SerializeKind__Modifier: "SerializeKind.Modifier",
  Ast_SerializeKind__Module: "SerializeKind.Module",
  Ast_SerializeKind__Nilable: "SerializeKind.Nilable",
  Ast_SerializeKind__Normal: "SerializeKind.Normal",
}
func Ast_SerializeKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_SerializeKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_SerializeKind_getTxt(arg1 LnsInt) string {
    return Ast_SerializeKindMap_[arg1];
}
// decl enum -- TypeInfoKind 
const Ast_TypeInfoKind__Abbr = 16
const Ast_TypeInfoKind__Alge = 14
const Ast_TypeInfoKind__Alternate = 18
const Ast_TypeInfoKind__Array = 4
const Ast_TypeInfoKind__Box = 19
const Ast_TypeInfoKind__CanEvalCtrl = 20
const Ast_TypeInfoKind__Class = 6
const Ast_TypeInfoKind__CombineIF = 25
const Ast_TypeInfoKind__DDD = 15
const Ast_TypeInfoKind__Enum = 11
const Ast_TypeInfoKind__Etc = 21
const Ast_TypeInfoKind__Ext = 24
const Ast_TypeInfoKind__Form = 22
const Ast_TypeInfoKind__FormFunc = 23
const Ast_TypeInfoKind__Func = 8
const Ast_TypeInfoKind__IF = 7
const Ast_TypeInfoKind__List = 3
const Ast_TypeInfoKind__Macro = 1
const Ast_TypeInfoKind__Map = 5
const Ast_TypeInfoKind__Method = 9
const Ast_TypeInfoKind__Module = 12
const Ast_TypeInfoKind__Nilable = 10
const Ast_TypeInfoKind__Prim = 2
const Ast_TypeInfoKind__Root = 0
const Ast_TypeInfoKind__Set = 17
const Ast_TypeInfoKind__Stem = 13
var Ast_TypeInfoKindList_ = NewLnsList( []LnsAny {
  Ast_TypeInfoKind__Root,
  Ast_TypeInfoKind__Macro,
  Ast_TypeInfoKind__Prim,
  Ast_TypeInfoKind__List,
  Ast_TypeInfoKind__Array,
  Ast_TypeInfoKind__Map,
  Ast_TypeInfoKind__Class,
  Ast_TypeInfoKind__IF,
  Ast_TypeInfoKind__Func,
  Ast_TypeInfoKind__Method,
  Ast_TypeInfoKind__Nilable,
  Ast_TypeInfoKind__Enum,
  Ast_TypeInfoKind__Module,
  Ast_TypeInfoKind__Stem,
  Ast_TypeInfoKind__Alge,
  Ast_TypeInfoKind__DDD,
  Ast_TypeInfoKind__Abbr,
  Ast_TypeInfoKind__Set,
  Ast_TypeInfoKind__Alternate,
  Ast_TypeInfoKind__Box,
  Ast_TypeInfoKind__CanEvalCtrl,
  Ast_TypeInfoKind__Etc,
  Ast_TypeInfoKind__Form,
  Ast_TypeInfoKind__FormFunc,
  Ast_TypeInfoKind__Ext,
  Ast_TypeInfoKind__CombineIF,
})
func Ast_TypeInfoKind_get__allList() *LnsList{
    return Ast_TypeInfoKindList_
}
var Ast_TypeInfoKindMap_ = map[LnsInt]string {
  Ast_TypeInfoKind__Abbr: "TypeInfoKind.Abbr",
  Ast_TypeInfoKind__Alge: "TypeInfoKind.Alge",
  Ast_TypeInfoKind__Alternate: "TypeInfoKind.Alternate",
  Ast_TypeInfoKind__Array: "TypeInfoKind.Array",
  Ast_TypeInfoKind__Box: "TypeInfoKind.Box",
  Ast_TypeInfoKind__CanEvalCtrl: "TypeInfoKind.CanEvalCtrl",
  Ast_TypeInfoKind__Class: "TypeInfoKind.Class",
  Ast_TypeInfoKind__CombineIF: "TypeInfoKind.CombineIF",
  Ast_TypeInfoKind__DDD: "TypeInfoKind.DDD",
  Ast_TypeInfoKind__Enum: "TypeInfoKind.Enum",
  Ast_TypeInfoKind__Etc: "TypeInfoKind.Etc",
  Ast_TypeInfoKind__Ext: "TypeInfoKind.Ext",
  Ast_TypeInfoKind__Form: "TypeInfoKind.Form",
  Ast_TypeInfoKind__FormFunc: "TypeInfoKind.FormFunc",
  Ast_TypeInfoKind__Func: "TypeInfoKind.Func",
  Ast_TypeInfoKind__IF: "TypeInfoKind.IF",
  Ast_TypeInfoKind__List: "TypeInfoKind.List",
  Ast_TypeInfoKind__Macro: "TypeInfoKind.Macro",
  Ast_TypeInfoKind__Map: "TypeInfoKind.Map",
  Ast_TypeInfoKind__Method: "TypeInfoKind.Method",
  Ast_TypeInfoKind__Module: "TypeInfoKind.Module",
  Ast_TypeInfoKind__Nilable: "TypeInfoKind.Nilable",
  Ast_TypeInfoKind__Prim: "TypeInfoKind.Prim",
  Ast_TypeInfoKind__Root: "TypeInfoKind.Root",
  Ast_TypeInfoKind__Set: "TypeInfoKind.Set",
  Ast_TypeInfoKind__Stem: "TypeInfoKind.Stem",
}
func Ast_TypeInfoKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_TypeInfoKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_TypeInfoKind_getTxt(arg1 LnsInt) string {
    return Ast_TypeInfoKindMap_[arg1];
}
// decl enum -- MutMode 
const Ast_MutMode__AllMut = 3
const Ast_MutMode__IMut = 0
const Ast_MutMode__IMutRe = 1
const Ast_MutMode__Mut = 2
var Ast_MutModeList_ = NewLnsList( []LnsAny {
  Ast_MutMode__IMut,
  Ast_MutMode__IMutRe,
  Ast_MutMode__Mut,
  Ast_MutMode__AllMut,
})
func Ast_MutMode_get__allList() *LnsList{
    return Ast_MutModeList_
}
var Ast_MutModeMap_ = map[LnsInt]string {
  Ast_MutMode__AllMut: "MutMode.AllMut",
  Ast_MutMode__IMut: "MutMode.IMut",
  Ast_MutMode__IMutRe: "MutMode.IMutRe",
  Ast_MutMode__Mut: "MutMode.Mut",
}
func Ast_MutMode__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_MutModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_MutMode_getTxt(arg1 LnsInt) string {
    return Ast_MutModeMap_[arg1];
}
// decl enum -- ScopeAccess 
const Ast_ScopeAccess__Full = 1
const Ast_ScopeAccess__Normal = 0
var Ast_ScopeAccessList_ = NewLnsList( []LnsAny {
  Ast_ScopeAccess__Normal,
  Ast_ScopeAccess__Full,
})
func Ast_ScopeAccess_get__allList() *LnsList{
    return Ast_ScopeAccessList_
}
var Ast_ScopeAccessMap_ = map[LnsInt]string {
  Ast_ScopeAccess__Full: "ScopeAccess.Full",
  Ast_ScopeAccess__Normal: "ScopeAccess.Normal",
}
func Ast_ScopeAccess__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_ScopeAccessMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_ScopeAccess_getTxt(arg1 LnsInt) string {
    return Ast_ScopeAccessMap_[arg1];
}
// decl enum -- CanEvalType 
const Ast_CanEvalType__Comp = 5
const Ast_CanEvalType__Equal = 3
const Ast_CanEvalType__Logical = 6
const Ast_CanEvalType__Math = 4
const Ast_CanEvalType__SetEq = 2
const Ast_CanEvalType__SetOp = 1
const Ast_CanEvalType__SetOpIMut = 0
var Ast_CanEvalTypeList_ = NewLnsList( []LnsAny {
  Ast_CanEvalType__SetOpIMut,
  Ast_CanEvalType__SetOp,
  Ast_CanEvalType__SetEq,
  Ast_CanEvalType__Equal,
  Ast_CanEvalType__Math,
  Ast_CanEvalType__Comp,
  Ast_CanEvalType__Logical,
})
func Ast_CanEvalType_get__allList() *LnsList{
    return Ast_CanEvalTypeList_
}
var Ast_CanEvalTypeMap_ = map[LnsInt]string {
  Ast_CanEvalType__Comp: "CanEvalType.Comp",
  Ast_CanEvalType__Equal: "CanEvalType.Equal",
  Ast_CanEvalType__Logical: "CanEvalType.Logical",
  Ast_CanEvalType__Math: "CanEvalType.Math",
  Ast_CanEvalType__SetEq: "CanEvalType.SetEq",
  Ast_CanEvalType__SetOp: "CanEvalType.SetOp",
  Ast_CanEvalType__SetOpIMut: "CanEvalType.SetOpIMut",
}
func Ast_CanEvalType__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_CanEvalTypeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_CanEvalType_getTxt(arg1 LnsInt) string {
    return Ast_CanEvalTypeMap_[arg1];
}
// decl enum -- MethodKind 
const Ast_MethodKind__All = 0
const Ast_MethodKind__Object = 2
const Ast_MethodKind__Static = 1
var Ast_MethodKindList_ = NewLnsList( []LnsAny {
  Ast_MethodKind__All,
  Ast_MethodKind__Static,
  Ast_MethodKind__Object,
})
func Ast_MethodKind_get__allList() *LnsList{
    return Ast_MethodKindList_
}
var Ast_MethodKindMap_ = map[LnsInt]string {
  Ast_MethodKind__All: "MethodKind.All",
  Ast_MethodKind__Object: "MethodKind.Object",
  Ast_MethodKind__Static: "MethodKind.Static",
}
func Ast_MethodKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_MethodKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_MethodKind_getTxt(arg1 LnsInt) string {
    return Ast_MethodKindMap_[arg1];
}
// decl enum -- LuavalConvKind 
const Ast_LuavalConvKind__InLua = 0
const Ast_LuavalConvKind__ToLua = 1
var Ast_LuavalConvKindList_ = NewLnsList( []LnsAny {
  Ast_LuavalConvKind__InLua,
  Ast_LuavalConvKind__ToLua,
})
func Ast_LuavalConvKind_get__allList_4231_() *LnsList{
    return Ast_LuavalConvKindList_
}
var Ast_LuavalConvKindMap_ = map[LnsInt]string {
  Ast_LuavalConvKind__InLua: "LuavalConvKind.InLua",
  Ast_LuavalConvKind__ToLua: "LuavalConvKind.ToLua",
}
func Ast_LuavalConvKind__from_4224_(arg1 LnsInt) LnsAny{
    if _, ok := Ast_LuavalConvKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_LuavalConvKind_getTxt(arg1 LnsInt) string {
    return Ast_LuavalConvKindMap_[arg1];
}
// decl enum -- MatchType 
const Ast_MatchType__Error = 2
const Ast_MatchType__Match = 0
const Ast_MatchType__Warn = 1
var Ast_MatchTypeList_ = NewLnsList( []LnsAny {
  Ast_MatchType__Match,
  Ast_MatchType__Warn,
  Ast_MatchType__Error,
})
func Ast_MatchType_get__allList() *LnsList{
    return Ast_MatchTypeList_
}
var Ast_MatchTypeMap_ = map[LnsInt]string {
  Ast_MatchType__Error: "MatchType.Error",
  Ast_MatchType__Match: "MatchType.Match",
  Ast_MatchType__Warn: "MatchType.Warn",
}
func Ast_MatchType__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_MatchTypeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_MatchType_getTxt(arg1 LnsInt) string {
    return Ast_MatchTypeMap_[arg1];
}
// decl enum -- BitOpKind 
const Ast_BitOpKind__And = 0
const Ast_BitOpKind__LShift = 3
const Ast_BitOpKind__Or = 1
const Ast_BitOpKind__RShift = 4
const Ast_BitOpKind__Xor = 2
var Ast_BitOpKindList_ = NewLnsList( []LnsAny {
  Ast_BitOpKind__And,
  Ast_BitOpKind__Or,
  Ast_BitOpKind__Xor,
  Ast_BitOpKind__LShift,
  Ast_BitOpKind__RShift,
})
func Ast_BitOpKind_get__allList() *LnsList{
    return Ast_BitOpKindList_
}
var Ast_BitOpKindMap_ = map[LnsInt]string {
  Ast_BitOpKind__And: "BitOpKind.And",
  Ast_BitOpKind__LShift: "BitOpKind.LShift",
  Ast_BitOpKind__Or: "BitOpKind.Or",
  Ast_BitOpKind__RShift: "BitOpKind.RShift",
  Ast_BitOpKind__Xor: "BitOpKind.Xor",
}
func Ast_BitOpKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Ast_BitOpKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_BitOpKind_getTxt(arg1 LnsInt) string {
    return Ast_BitOpKindMap_[arg1];
}
var Ast_extStartId LnsInt
var Ast_extMaxId LnsInt
var Ast_userStartId LnsInt
var Ast_rootProcessInfo *Ast_ProcessInfo
var Ast_rootTypeId LnsInt
var Ast_sym2builtInTypeMap *LnsMap
var Ast_builtInTypeIdSet *LnsMap
var Ast_txt2AccessModeMap *LnsMap
var Ast_accessMode2txtMap *LnsMap
var Ast_rootScope *Ast_Scope
var Ast_dummyList *LnsList
var Ast_rootChildren *LnsList
var Ast_headTypeInfo *Ast_TypeInfo
var Ast_defaultTypeNameCtrl *Ast_TypeNameCtrl
var Ast_dummySymbol *Ast_SymbolInfo
var Ast_boxRootAltType *Ast_AlternateTypeInfo
var Ast_boxRootScope *Ast_Scope
var Ast_immutableTypeSet *LnsSet
var Ast_builtinTypeNone *Ast_TypeInfo
var Ast_builtinTypeEmpty *Ast_TypeInfo
var Ast_builtinTypeNeverRet *Ast_TypeInfo
var Ast_builtinTypeStem *Ast_TypeInfo
var Ast_builtinTypeStem_ *Ast_TypeInfo
var Ast_builtinTypeBool *Ast_TypeInfo
var Ast_builtinTypeInt *Ast_TypeInfo
var Ast_builtinTypeReal *Ast_TypeInfo
var Ast_builtinTypeChar *Ast_TypeInfo
var Ast_builtinTypeMapping *Ast_TypeInfo
var Ast_builtinTypeAsyncItem *Ast_TypeInfo
var Ast_builtinTypeString *Ast_TypeInfo
var Ast_builtinTypeMap *Ast_TypeInfo
var Ast_builtinTypeSet *Ast_TypeInfo
var Ast_builtinTypeList *Ast_TypeInfo
var Ast_builtinTypeArray *Ast_TypeInfo
var Ast_builtinTypeBox *Ast_BoxTypeInfo
var Ast_builtinTypeNil *Ast_TypeInfo
var Ast_builtinTypeDDD *Ast_TypeInfo
var Ast_builtinTypeForm *Ast_TypeInfo
var Ast_builtinTypeSymbol *Ast_TypeInfo
var Ast_builtinTypeStat *Ast_TypeInfo
var Ast_builtinTypeExp *Ast_TypeInfo
var Ast_builtinTypeMultiExp *Ast_TypeInfo
var Ast_builtinTypeAbbr *Ast_TypeInfo
var Ast_builtinTypeAbbrNone *Ast_TypeInfo
var Ast_builtinTypeLua *Ast_TypeInfo
var Ast_builtinTypeDDDLua *Ast_TypeInfo
var Ast_numberTypeSet *LnsSet
var Ast_builtinTypeInfo2Map *Ast_TypeInfo2Map
var Ast_bitBinOpMap *LnsMap
var Ast_compOpSet *LnsSet
var Ast_mathCompOpSet *LnsSet
type Ast_LuavalResult__Err struct{
Val1 string
}
func (self *Ast_LuavalResult__Err) GetTxt() string {
return "LuavalResult.Err"
}
type Ast_LuavalResult__OK struct{
Val1 *Ast_TypeInfo
Val2 bool
}
func (self *Ast_LuavalResult__OK) GetTxt() string {
return "LuavalResult.OK"
}
type Ast_OverrideMut__IMut struct{
Val1 *Ast_TypeInfo
}
func (self *Ast_OverrideMut__IMut) GetTxt() string {
return "OverrideMut.IMut"
}
type Ast_OverrideMut__None struct{
}
var Ast_OverrideMut__None_Obj = &Ast_OverrideMut__None{}
func (self *Ast_OverrideMut__None) GetTxt() string {
return "OverrideMut.None"
}
type Ast_OverrideMut__Prefix struct{
Val1 *Ast_TypeInfo
}
func (self *Ast_OverrideMut__Prefix) GetTxt() string {
return "OverrideMut.Prefix"
}
type Ast_EnumLiteral__Int struct{
Val1 LnsInt
}
func (self *Ast_EnumLiteral__Int) GetTxt() string {
return "EnumLiteral.Int"
}
type Ast_EnumLiteral__Real struct{
Val1 LnsReal
}
func (self *Ast_EnumLiteral__Real) GetTxt() string {
return "EnumLiteral.Real"
}
type Ast_EnumLiteral__Str struct{
Val1 string
}
func (self *Ast_EnumLiteral__Str) GetTxt() string {
return "EnumLiteral.Str"
}
type Ast_OverridingType__NoReady struct{
}
var Ast_OverridingType__NoReady_Obj = &Ast_OverridingType__NoReady{}
func (self *Ast_OverridingType__NoReady) GetTxt() string {
return "OverridingType.NoReady"
}
type Ast_OverridingType__NotOverride struct{
}
var Ast_OverridingType__NotOverride_Obj = &Ast_OverridingType__NotOverride{}
func (self *Ast_OverridingType__NotOverride) GetTxt() string {
return "OverridingType.NotOverride"
}
type Ast_OverridingType__Override struct{
Val1 *Ast_TypeInfo
}
func (self *Ast_OverridingType__Override) GetTxt() string {
return "OverridingType.Override"
}
type Ast_CommonType__Combine struct{
Val1 *Ast_CombineType
}
func (self *Ast_CommonType__Combine) GetTxt() string {
return "CommonType.Combine"
}
type Ast_CommonType__Normal struct{
Val1 *Ast_TypeInfo
}
func (self *Ast_CommonType__Normal) GetTxt() string {
return "CommonType.Normal"
}
type Ast_filterForm func (arg1 *Ast_SymbolInfo) bool
// for 5460
func Ast_convExp21900(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 6856
func Ast_convExp29173(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 6991
func Ast_convExp30142(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 6991
func Ast_convExp30300(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 6991
func Ast_convExp30466(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 6991
func Ast_convExp30637(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 6991
func Ast_convExp30803(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 6991
func Ast_convExp31009(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 6991
func Ast_convExp31186(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 6991
func Ast_convExp31344(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7030
func Ast_convExp31499(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
func ProcessInfo_switchIdProvier___anonymous_1070_() string {
    return "start"
}
// 115: decl @lune.@base.@Ast.getRootProcessInfo
func Ast_getRootProcessInfo() *Ast_ProcessInfo {
    return Ast_rootProcessInfo
}

// 208: decl @lune.@base.@Ast.getSym2builtInTypeMap
func Ast_getSym2builtInTypeMap() *LnsMap {
    return Ast_sym2builtInTypeMap
}

// 212: decl @lune.@base.@Ast.getBuiltInTypeIdMap
func Ast_getBuiltInTypeIdMap() *LnsMap {
    return Ast_builtInTypeIdSet
}

// 270: decl @lune.@base.@Ast.isBuiltin
func Ast_isBuiltin(typeId LnsInt) bool {
    return Ast_TypeInfo2Stem(Ast_builtInTypeIdSet.Items[typeId]) != nil
}

// 275: decl @lune.@base.@Ast.isPubToExternal
func Ast_isPubToExternal(mode LnsInt) bool {
    if _switch596 := mode; _switch596 == Ast_AccessMode__Pub || _switch596 == Ast_AccessMode__Pro || _switch596 == Ast_AccessMode__Global {
        return true
    }
    return false
}

// 291: decl @lune.@base.@Ast.txt2AccessMode
func Ast_txt2AccessMode(accessMode string) LnsAny {
    return Ast_txt2AccessModeMap.Items[accessMode]
}

// 302: decl @lune.@base.@Ast.accessMode2txt
func Ast_accessMode2txt(accessMode LnsInt) string {
    return Lns_unwrap( Ast_accessMode2txtMap.Items[accessMode]).(string)
}

// 316: decl @lune.@base.@Ast.isMutable
func Ast_isMutable(mode LnsInt) bool {
    if _switch809 := mode; _switch809 == Ast_MutMode__AllMut || _switch809 == Ast_MutMode__Mut {
        return true
    }
    return false
}

// 962: decl @lune.@base.@Ast.applyGenericDefault
func Ast_applyGenericDefault(typeInfo *Ast_TypeInfo,alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    {
        _genType := typeInfo.FP.ApplyGeneric(alt2typeMap, moduleTypeInfo)
        if _genType != nil {
            genType := _genType.(*Ast_TypeInfo)
            return genType
        }
    }
    return typeInfo
}


// 990: decl @lune.@base.@Ast.getAllNameForKind
func Ast_getAllNameForKind(classInfo *Ast_TypeInfo,kind LnsInt,symbolKind LnsInt) *Util_OrderedSet {
    var nameSet *Util_OrderedSet
    nameSet = NewUtil_OrderedSet()
    var process func(scope *Ast_Scope)
    process = func(scope *Ast_Scope) {
        {
            _inherit := scope.FP.Get_inherit()
            if _inherit != nil {
                inherit := _inherit.(*Ast_Scope)
                process(inherit)
            }
        }
        {
            __collection2730 := scope.FP.Get_symbol2SymbolInfoMap()
            __sorted2730 := __collection2730.CreateKeyListStr()
            __sorted2730.Sort( LnsItemKindStr, nil )
            for _, ___key2730 := range( __sorted2730.Items ) {
                symbolInfo := __collection2730.Items[ ___key2730 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if _switch2728 := symbolInfo.FP.Get_kind(); _switch2728 == symbolKind {
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( symbolKind == Ast_SymbolKind__Mtd) &&
                        Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_name() == "__init") ).(bool)){
                    } else { 
                        var staticFlag bool
                        staticFlag = symbolInfo.FP.Get_staticFlag()
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( kind == Ast_MethodKind__All) ||
                            Lns_GetEnv().SetStackVal( kind == Ast_MethodKind__Static) &&
                            Lns_GetEnv().SetStackVal( staticFlag) ||
                            Lns_GetEnv().SetStackVal( kind == Ast_MethodKind__Object) &&
                            Lns_GetEnv().SetStackVal( Lns_op_not(staticFlag)) ).(bool){
                            nameSet.FP.Add(symbolInfo.FP.Get_name())
                        }
                    }
                }
            }
        }
    }
    {
        _scope := classInfo.FP.Get_scope()
        if _scope != nil {
            scope := _scope.(*Ast_Scope)
            process(scope)
        }
    }
    return nameSet
}

// 1024: decl @lune.@base.@Ast.getAllMethodName
func Ast_getAllMethodName(classInfo *Ast_TypeInfo,kind LnsInt) *Util_OrderedSet {
    return Ast_getAllNameForKind(classInfo, kind, Ast_SymbolKind__Mtd)
}

// 1077: decl @lune.@base.@Ast.getScope
func Ast_getScope(typeInfo *Ast_TypeInfo) LnsAny {
    return typeInfo.FP.Get_scope()
}

// 1083: decl @lune.@base.@Ast.isExtId
func Ast_isExtId(typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_typeId() >= Ast_extStartId{
        return true
    }
    return false
}

// 2100: decl @lune.@base.@Ast.dumpScope.dumpScopeSub
func dumpScope__dumpScopeSub_2620_(scope LnsAny,prefix string,readyIdSet *LnsSet) {
    {
        __exp := scope
        if __exp != nil {
            _exp := __exp.(*Ast_Scope)
            if readyIdSet.Has(Ast_Scope2Stem(_exp)){
                return 
            }
            readyIdSet.Add(Ast_Scope2Stem(_exp))
            if len(prefix) > 20{
                Util_err("illegal")
            }
            {
                __collection7114 := _exp.FP.Get_symbol2SymbolInfoMap()
                __sorted7114 := __collection7114.CreateKeyListStr()
                __sorted7114.Sort( LnsItemKindStr, nil )
                for _, _symbol := range( __sorted7114.Items ) {
                    symbolInfo := __collection7114.Items[ _symbol ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbol := _symbol.(string)
                    Util_log(Lns_getVM().String_format("scope: %s, %s, %s", []LnsAny{prefix, _exp, symbol}))
                    {
                        _subScope := symbolInfo.FP.Get_typeInfo().FP.Get_scope()
                        if _subScope != nil {
                            subScope := _subScope.(*Ast_Scope)
                            dumpScope__dumpScopeSub_2620_(subScope, prefix + "  ", readyIdSet)
                        }
                    }
                }
            }
        }
    }
}

// 2095: decl @lune.@base.@Ast.dumpScope
func Ast_dumpScope(workscope LnsAny,workprefix string) {
    dumpScope__dumpScopeSub_2620_(workscope, workprefix, NewLnsSet([]LnsAny{}))
}


// 3095: decl @lune.@base.@Ast.isGenericType
func Ast_isGenericType(typeInfo *Ast_TypeInfo) bool {
    if Lns_isCondTrue( Ast_GenericTypeInfoDownCastF(typeInfo.FP)){
        return true
    }
    return false
}

// 3308: decl @lune.@base.@Ast.getEnumLiteralVal
func Ast_getEnumLiteralVal(obj LnsAny) LnsAny {
    switch _exp11591 := obj.(type) {
    case *Ast_EnumLiteral__Int:
    val := _exp11591.Val1
        return val
    case *Ast_EnumLiteral__Real:
    val := _exp11591.Val1
        return val
    case *Ast_EnumLiteral__Str:
    val := _exp11591.Val1
        return val
    }
// insert a dummy
    return nil
}


// 4037: decl @lune.@base.@Ast.isExtType
func Ast_isExtType(typeInfo *Ast_TypeInfo) bool {
    return Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Ext) ||
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__DDD) &&
            Lns_GetEnv().SetStackVal( typeInfo.FP.Get_extedType() != typeInfo) ).(bool))) ).(bool)
}

// 4044: decl @lune.@base.@Ast.isMutableType
func Ast_isMutableType(typeInfo *Ast_TypeInfo) bool {
    typeInfo = typeInfo.FP.Get_nonnilableType()
    
    if Ast_immutableTypeSet.Has(Ast_TypeInfo2Stem(typeInfo)){
        return false
    }
    if _switch14856 := typeInfo.FP.Get_kind(); _switch14856 == Ast_TypeInfoKind__FormFunc || _switch14856 == Ast_TypeInfoKind__Enum {
        return false
    }
    return true
}

// 4147: decl @lune.@base.@Ast.addBuiltin
func Ast_addBuiltin(typeInfo *Ast_TypeInfo) {
    Ast_builtInTypeIdSet.Set(typeInfo.FP.Get_typeId(),typeInfo)
}

// 4151: decl @lune.@base.@Ast.registBuiltin
func Ast_registBuiltin_4166_(idName string,typeTxt string,kind LnsInt,typeInfo *Ast_TypeInfo,nilableTypeInfo *Ast_TypeInfo,registScope bool) {
    Ast_sym2builtInTypeMap.Set(typeTxt,&NewAst_NormalSymbolInfo(Ast_rootProcessInfo, Ast_SymbolKind__Typ, false, false, Ast_rootScope, Ast_AccessMode__Pub, false, typeTxt, nil, typeInfo, Ast_MutMode__IMut, true, false).Ast_SymbolInfo)
    if nilableTypeInfo != Ast_headTypeInfo{
        Ast_sym2builtInTypeMap.Set(typeTxt + "!",&NewAst_NormalSymbolInfo(Ast_rootProcessInfo, Ast_SymbolKind__Typ, false, kind == Ast_TypeInfoKind__Func, Ast_rootScope, Ast_AccessMode__Pub, false, typeTxt, nil, nilableTypeInfo, Ast_MutMode__IMut, true, false).Ast_SymbolInfo)
    }
    Ast_addBuiltin(typeInfo)
    Ast_addBuiltin(Ast_rootProcessInfo.FP.CreateModifier(typeInfo, Ast_MutMode__IMut))
    if typeInfo.FP.Get_nilableTypeInfo() != Ast_headTypeInfo{
        Ast_addBuiltin(typeInfo.FP.Get_nilableTypeInfo())
        Ast_addBuiltin(Ast_rootProcessInfo.FP.CreateModifier(typeInfo.FP.Get_nilableTypeInfo(), Ast_MutMode__IMut))
    }
    if registScope{
        Ast_rootScope.FP.AddClass(Ast_rootProcessInfo, typeTxt, nil, typeInfo)
    }
}

// 4271: decl @lune.@base.@Ast.isClass
func Ast_isClass(typeInfo *Ast_TypeInfo) bool {
    return Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Class) &&
        Lns_GetEnv().SetStackVal( typeInfo != Ast_builtinTypeString) ).(bool)
}

// 4340: decl @lune.@base.@Ast.failCreateLuavalWith
func Ast_failCreateLuavalWith_4234_(typeInfo *Ast_TypeInfo,convFlag LnsInt,validToCheck bool)(LnsAny, bool) {
    if Ast_isExtType(typeInfo){
        return nil, true
    }
    var mess string
    mess = Lns_getVM().String_format("not support to use the type as Luaval -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)})
    if _switch16783 := typeInfo.FP.Get_kind(); _switch16783 == Ast_TypeInfoKind__Nilable {
        return Ast_failCreateLuavalWith_4234_(typeInfo.FP.Get_nonnilableType(), convFlag, validToCheck)
    } else if _switch16783 == Ast_TypeInfoKind__Prim {
        return nil, true
    } else if _switch16783 == Ast_TypeInfoKind__Form || _switch16783 == Ast_TypeInfoKind__IF || _switch16783 == Ast_TypeInfoKind__DDD {
        if Lns_op_not(validToCheck){
            return nil, false
        }
        if convFlag != Ast_LuavalConvKind__InLua{
            return mess, false
        }
        return nil, false
    } else if _switch16783 == Ast_TypeInfoKind__Stem {
        return nil, false
    } else if _switch16783 == Ast_TypeInfoKind__Class {
        if typeInfo != Ast_builtinTypeString{
            if Lns_op_not(validToCheck){
                return nil, false
            }
            if convFlag != Ast_LuavalConvKind__InLua{
                return mess, false
            }
            return nil, false
        }
        return nil, true
    } else if _switch16783 == Ast_TypeInfoKind__Array || _switch16783 == Ast_TypeInfoKind__List || _switch16783 == Ast_TypeInfoKind__Map {
        if Lns_op_not(validToCheck){
            return nil, false
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( convFlag != Ast_LuavalConvKind__ToLua) &&
            Lns_GetEnv().SetStackVal( Ast_isMutable(typeInfo.FP.Get_mutMode())) ).(bool)){
            return "not support mutable collecion. " + mess, false
        }
        var canConv bool
        canConv = true
        for _, _itemType := range( typeInfo.FP.Get_itemTypeInfoList().Items ) {
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var err LnsAny
            var work bool
            err,work = Ast_failCreateLuavalWith_4234_(itemType, convFlag, validToCheck)
            if err != nil{
                err_3819 := err.(string)
                return err_3819, false
            }
            if Lns_op_not(work){
                canConv = false
                
            }
        }
        
        canConv = false
        
        return nil, canConv
    } else if _switch16783 == Ast_TypeInfoKind__FormFunc {
        if Lns_op_not(validToCheck){
            return nil, false
        }
        if convFlag != Ast_LuavalConvKind__InLua{
            return mess, false
        }
        if typeInfo.FP.Get_itemTypeInfoList().Len() != 0{
            return mess, false
        }
        var canConv bool
        canConv = true
        for _, _itemType := range( typeInfo.FP.Get_argTypeInfoList().Items ) {
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var err LnsAny
            var work bool
            err,work = Ast_failCreateLuavalWith_4234_(itemType, convFlag, validToCheck)
            if err != nil{
                err_3831 := err.(string)
                return err_3831, false
            }
            if Lns_op_not(work){
                canConv = false
                
            }
        }
        
        for _, _itemType := range( typeInfo.FP.Get_retTypeInfoList().Items ) {
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var err LnsAny
            var work bool
            err,work = Ast_failCreateLuavalWith_4234_(itemType, convFlag, validToCheck)
            if err != nil{
                err_3838 := err.(string)
                return err_3838, false
            }
            if Lns_op_not(work){
                canConv = false
                
            }
        }
        
        canConv = false
        
        return nil, canConv
    }
    return Lns_getVM().String_format("not support -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}), false
}

// 4451: decl @lune.@base.@Ast.isStruct
func Ast_isStruct(typeInfo *Ast_TypeInfo) bool {
    if _switch16933 := typeInfo.FP.Get_kind(); _switch16933 == Ast_TypeInfoKind__Class {
        if typeInfo == Ast_builtinTypeString{
            return false
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( typeInfo.FP.Get_baseTypeInfo() != Ast_headTypeInfo) ||
            Lns_GetEnv().SetStackVal( typeInfo.FP.Get_interfaceList().Len() != 0) ||
            Lns_GetEnv().SetStackVal( typeInfo.FP.Get_children().Len() != 1) ).(bool){
            return false
        }
        return true
    }
    return false
}

// 4474: decl @lune.@base.@Ast.isConditionalbe
func Ast_isConditionalbe(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo) bool {
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_nilable()) ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Equals(processInfo, Ast_builtinTypeBool, nil, nil)) ).(bool){
        return true
    }
    return false
}











// 5252: decl @lune.@base.@Ast.applyGenericList
func Ast_applyGenericList_4604_(typeList *LnsList,alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo)(LnsAny, bool) {
    var typeInfoList *LnsList
    typeInfoList = NewLnsList([]LnsAny{})
    var needNew bool
    needNew = false
    for _, _srcType := range( typeList.Items ) {
        srcType := _srcType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        {
            _typeInfo := srcType.FP.ApplyGeneric(alt2typeMap, moduleTypeInfo)
            if _typeInfo != nil {
                typeInfo := _typeInfo.(*Ast_TypeInfo)
                typeInfoList.Insert(Ast_TypeInfo2Stem(typeInfo))
                if srcType != typeInfo{
                    needNew = true
                    
                }
            } else {
                return nil, false
            }
        }
    }
    return typeInfoList, needNew
}


// 5595: decl @lune.@base.@Ast.convToExtTypeList
func Ast_convToExtTypeList(processInfo *Ast_ProcessInfo,list *LnsList)(LnsAny, string) {
    var extList *LnsList
    extList = NewLnsList([]LnsAny{})
    for _, _typeInfo := range( list.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        switch _exp22562 := processInfo.FP.CreateLuaval(typeInfo, true).(type) {
        case *Ast_LuavalResult__OK:
        extType := _exp22562.Val1
            extList.Insert(Ast_TypeInfo2Stem(extType))
        case *Ast_LuavalResult__Err:
        err := _exp22562.Val1
            return nil, err
        }
    }
    return extList, ""
}

// 5639: decl @lune.@base.@Ast.isNumberType
func Ast_isNumberType(typeInfo *Ast_TypeInfo) bool {
    return Ast_numberTypeSet.Has(Ast_TypeInfo2Stem(typeInfo.FP.Get_srcTypeInfo()))
}




// 6153: decl @lune.@base.@Ast.isSettableToForm
func Ast_isSettableToForm_5182_(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_argTypeInfoList().Len() > 0{
        for _, _argType := range( typeInfo.FP.Get_argTypeInfoList().Items ) {
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            {
                _dddType := Ast_DDDTypeInfoDownCastF(argType.FP)
                if _dddType != nil {
                    dddType := _dddType.(*Ast_DDDTypeInfo)
                    if Lns_op_not(dddType.FP.Get_typeInfo().FP.Get_nilableTypeInfo().FP.Equals(processInfo, Ast_builtinTypeStem_, nil, nil)){
                        return false
                    }
                } else {
                    if Lns_op_not(argType.FP.Get_srcTypeInfo().FP.Equals(processInfo, Ast_builtinTypeStem_, nil, nil)){
                        return false
                    }
                }
            }
        }
    }
    return true
}



// 6682: decl @lune.@base.@Ast.createProcessInfo
func Ast_createProcessInfo(validExtType bool) *Ast_ProcessInfo {
    return Ast_ProcessInfo_createUser_1064_(validExtType, Ast_builtinTypeInfo2Map.FP.Clone())
}


// declaration Class -- IdProvider
type Ast_IdProviderMtd interface {
    GetNewId() LnsInt
    Get_id() LnsInt
    Increment()
}
type Ast_IdProvider struct {
    id LnsInt
    maxId LnsInt
    FP Ast_IdProviderMtd
}
func Ast_IdProvider2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_IdProvider).FP
}
type Ast_IdProviderDownCast interface {
    ToAst_IdProvider() *Ast_IdProvider
}
func Ast_IdProviderDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_IdProviderDownCast)
    if ok { return work.ToAst_IdProvider() }
    return nil
}
func (obj *Ast_IdProvider) ToAst_IdProvider() *Ast_IdProvider {
    return obj
}
func NewAst_IdProvider(arg1 LnsInt, arg2 LnsInt) *Ast_IdProvider {
    obj := &Ast_IdProvider{}
    obj.FP = obj
    obj.InitAst_IdProvider(arg1, arg2)
    return obj
}
func (self *Ast_IdProvider) InitAst_IdProvider(arg1 LnsInt, arg2 LnsInt) {
    self.id = arg1
    self.maxId = arg2
}
func (self *Ast_IdProvider) Get_id() LnsInt{ return self.id }
// 34: decl @lune.@base.@Ast.IdProvider.increment
func (self *Ast_IdProvider) Increment() {
    self.id = self.id + 1
    
    if self.id >= self.maxId{
        Util_err("id is over")
    }
}

// 41: decl @lune.@base.@Ast.IdProvider.getNewId
func (self *Ast_IdProvider) GetNewId() LnsInt {
    self.id = self.id + 1
    
    if self.id >= self.maxId{
        Util_err("id is over")
    }
    return self.id
}


// declaration Class -- ProcessInfo
type Ast_ProcessInfoMtd interface {
    CreateAdvertiseMethodFrom(arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo) *Ast_TypeInfo
    CreateAlge(arg1 *Ast_Scope, arg2 *Ast_TypeInfo, arg3 bool, arg4 LnsInt, arg5 string) *Ast_AlgeTypeInfo
    CreateAlias(arg1 *Ast_ProcessInfo, arg2 string, arg3 bool, arg4 LnsInt, arg5 *Ast_TypeInfo, arg6 *Ast_TypeInfo) *Ast_AliasTypeInfo
    CreateAlternate(arg1 bool, arg2 LnsInt, arg3 string, arg4 LnsInt, arg5 *Ast_TypeInfo, arg6 LnsAny, arg7 LnsAny) *Ast_AlternateTypeInfo
    CreateArray(arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 LnsInt) *Ast_TypeInfo
    CreateBox(arg1 LnsInt, arg2 *Ast_TypeInfo) *Ast_TypeInfo
    CreateClass(arg1 bool, arg2 bool, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 *LnsList, arg7 *Ast_TypeInfo, arg8 bool, arg9 LnsInt, arg10 string) *Ast_TypeInfo
    CreateDDD(arg1 *Ast_TypeInfo, arg2 bool, arg3 bool) *Ast_DDDTypeInfo
    CreateEnum(arg1 *Ast_Scope, arg2 *Ast_TypeInfo, arg3 bool, arg4 LnsInt, arg5 string, arg6 *Ast_TypeInfo) *Ast_EnumTypeInfo
    CreateFunc(arg1 bool, arg2 bool, arg3 LnsAny, arg4 LnsInt, arg5 *Ast_TypeInfo, arg6 bool, arg7 bool, arg8 bool, arg9 LnsInt, arg10 string, arg11 LnsAny, arg12 LnsAny, arg13 LnsAny, arg14 LnsAny) *Ast_TypeInfo
    CreateGeneric(arg1 *Ast_TypeInfo, arg2 *LnsList, arg3 *Ast_TypeInfo) *Ast_GenericTypeInfo
    CreateList(arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 LnsInt) *Ast_TypeInfo
    CreateLuaval(arg1 *Ast_TypeInfo, arg2 bool) LnsAny
    CreateMap(arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 *Ast_TypeInfo, arg4 *Ast_TypeInfo, arg5 LnsInt) *Ast_TypeInfo
    CreateModifier(arg1 *Ast_TypeInfo, arg2 LnsInt) *Ast_TypeInfo
    CreateModule(arg1 *Ast_Scope, arg2 *Ast_TypeInfo, arg3 bool, arg4 string, arg5 bool) *Ast_TypeInfo
    CreateSet(arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 LnsInt) *Ast_TypeInfo
    Get_idProv() *Ast_IdProvider
    Get_idProvBase() *Ast_IdProvider
    Get_idProvExt() *Ast_IdProvider
    Get_idProvScope() *Ast_IdProvider
    Get_idProvSym() *Ast_IdProvider
    get_typeInfo2Map() *Ast_TypeInfo2Map
    Get_validExtType() bool
    set_typeInfo2Map(arg1 *Ast_TypeInfo2Map)
    SwitchIdProvier(arg1 LnsInt)
}
type Ast_ProcessInfo struct {
    idProvSym *Ast_IdProvider
    idProvScope *Ast_IdProvider
    idProv *Ast_IdProvider
    idProvBase *Ast_IdProvider
    idProvExt *Ast_IdProvider
    typeInfo2Map LnsAny
    validExtType bool
    FP Ast_ProcessInfoMtd
}
func Ast_ProcessInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_ProcessInfo).FP
}
type Ast_ProcessInfoDownCast interface {
    ToAst_ProcessInfo() *Ast_ProcessInfo
}
func Ast_ProcessInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_ProcessInfoDownCast)
    if ok { return work.ToAst_ProcessInfo() }
    return nil
}
func (obj *Ast_ProcessInfo) ToAst_ProcessInfo() *Ast_ProcessInfo {
    return obj
}
func NewAst_ProcessInfo(arg1 *Ast_IdProvider, arg2 bool, arg3 LnsAny) *Ast_ProcessInfo {
    obj := &Ast_ProcessInfo{}
    obj.FP = obj
    obj.InitAst_ProcessInfo(arg1, arg2, arg3)
    return obj
}
func (self *Ast_ProcessInfo) Get_idProvSym() *Ast_IdProvider{ return self.idProvSym }
func (self *Ast_ProcessInfo) Get_idProvScope() *Ast_IdProvider{ return self.idProvScope }
func (self *Ast_ProcessInfo) Get_idProv() *Ast_IdProvider{ return self.idProv }
func (self *Ast_ProcessInfo) Get_idProvBase() *Ast_IdProvider{ return self.idProvBase }
func (self *Ast_ProcessInfo) Get_idProvExt() *Ast_IdProvider{ return self.idProvExt }
func (self *Ast_ProcessInfo) Get_validExtType() bool{ return self.validExtType }
// 74: decl @lune.@base.@Ast.ProcessInfo.get_typeInfo2Map
func (self *Ast_ProcessInfo) get_typeInfo2Map() *Ast_TypeInfo2Map {
    return Lns_unwrap( self.typeInfo2Map).(*Ast_TypeInfo2Map)
}

// 77: decl @lune.@base.@Ast.ProcessInfo.set_typeInfo2Map
func (self *Ast_ProcessInfo) set_typeInfo2Map(typeInfo2Map *Ast_TypeInfo2Map) {
    self.typeInfo2Map = typeInfo2Map
    
}

// 81: DeclConstr
func (self *Ast_ProcessInfo) InitAst_ProcessInfo(idProvBase *Ast_IdProvider,validExtType bool,typeInfo2Map LnsAny) {
    self.validExtType = validExtType
    
    self.idProvBase = idProvBase
    
    self.idProvExt = NewAst_IdProvider(Ast_extStartId, Ast_extMaxId)
    
    self.idProv = self.idProvBase
    
    self.idProvSym = NewAst_IdProvider(0, Ast_extMaxId)
    
    self.idProvScope = NewAst_IdProvider(0, Ast_extMaxId)
    
    self.typeInfo2Map = typeInfo2Map
    
}

// 92: decl @lune.@base.@Ast.ProcessInfo.createRoot
func Ast_ProcessInfo_createRoot_1061_() *Ast_ProcessInfo {
    return NewAst_ProcessInfo(NewAst_IdProvider(0, Ast_userStartId), true, nil)
}

// 97: decl @lune.@base.@Ast.ProcessInfo.createUser
func Ast_ProcessInfo_createUser_1064_(validExtType bool,typeInfo2Map LnsAny) *Ast_ProcessInfo {
    return NewAst_ProcessInfo(NewAst_IdProvider(Ast_userStartId, Ast_extStartId), validExtType, typeInfo2Map)
}

// 104: decl @lune.@base.@Ast.ProcessInfo.switchIdProvier
func (self *Ast_ProcessInfo) SwitchIdProvier(idType LnsInt) {
    __func__ := "@lune.@base.@Ast.ProcessInfo.switchIdProvier"
    Log_log(Log_Level__Trace, __func__, 105, Log_CreateMessage(ProcessInfo_switchIdProvier___anonymous_1070_))
    
    if idType == Ast_IdType__Base{
        self.idProv = self.idProvBase
        
    } else { 
        self.idProv = self.idProvExt
        
    }
}


// 3986: decl @lune.@base.@Ast.ProcessInfo.createAlternate
func (self *Ast_ProcessInfo) CreateAlternate(belongClassFlag bool,altIndex LnsInt,txt string,accessMode LnsInt,moduleTypeInfo *Ast_TypeInfo,baseTypeInfo LnsAny,interfaceList LnsAny) *Ast_AlternateTypeInfo {
    return NewAst_AlternateTypeInfo(self, belongClassFlag, altIndex, txt, accessMode, moduleTypeInfo, baseTypeInfo, interfaceList)
}

// 4057: decl @lune.@base.@Ast.ProcessInfo.createModifier
func (self *Ast_ProcessInfo) CreateModifier(srcTypeInfo *Ast_TypeInfo,mutMode LnsInt) *Ast_TypeInfo {
    srcTypeInfo = srcTypeInfo.FP.Get_srcTypeInfo()
    
    if Lns_op_not(Ast_isMutableType(srcTypeInfo)){
        return srcTypeInfo
    }
    if _switch14936 := mutMode; _switch14936 == Ast_MutMode__IMut || _switch14936 == Ast_MutMode__IMutRe {
        {
            __exp := self.FP.get_typeInfo2Map().ImutModifierMap.Items[srcTypeInfo]
            if __exp != nil {
                _exp := __exp.(*Ast_TypeInfo)
                return _exp
            }
        }
    } else if _switch14936 == Ast_MutMode__AllMut {
        {
            __exp := self.FP.get_typeInfo2Map().MutModifierMap.Items[srcTypeInfo]
            if __exp != nil {
                _exp := __exp.(*Ast_TypeInfo)
                return _exp
            }
        }
    }
    self.FP.Get_idProv().FP.Increment()
    var modifier *Ast_TypeInfo
    if srcTypeInfo.FP.Get_nonnilableType().FP.Get_kind() == Ast_TypeInfoKind__Ext{
        switch _exp15002 := self.FP.CreateLuaval(self.FP.CreateModifier(srcTypeInfo.FP.Get_extedType(), mutMode), false).(type) {
        case *Ast_LuavalResult__OK:
        workType := _exp15002.Val1
            if srcTypeInfo.FP.Get_nilable(){
                modifier = workType.FP.Get_nilableTypeInfo()
                
            } else { 
                modifier = workType
                
            }
        case *Ast_LuavalResult__Err:
        err := _exp15002.Val1
            Util_err(err)
        }
    } else { 
        modifier = &NewAst_ModifierTypeInfo(nil, self, srcTypeInfo, self.FP.Get_idProv().FP.Get_id(), mutMode).Ast_TypeInfo
        
    }
    if _switch15069 := mutMode; _switch15069 == Ast_MutMode__IMut || _switch15069 == Ast_MutMode__IMutRe {
        self.FP.get_typeInfo2Map().ImutModifierMap.Set(srcTypeInfo,modifier)
    } else if _switch15069 == Ast_MutMode__AllMut {
        self.FP.get_typeInfo2Map().MutModifierMap.Set(srcTypeInfo,modifier)
    }
    return modifier
}

// 4483: decl @lune.@base.@Ast.ProcessInfo.createBox
func (self *Ast_ProcessInfo) CreateBox(accessMode LnsInt,nonnilableType *Ast_TypeInfo) *Ast_TypeInfo {
    {
        _boxType := self.FP.get_typeInfo2Map().BoxMap.Items[nonnilableType]
        if _boxType != nil {
            boxType := _boxType.(*Ast_TypeInfo)
            return boxType
        }
    }
    var boxType *Ast_BoxTypeInfo
    boxType = NewAst_BoxTypeInfo(self, self.FP.Get_idProv().FP.GetNewId(), accessMode, nonnilableType)
    self.FP.get_typeInfo2Map().BoxMap.Set(nonnilableType,&boxType.Ast_TypeInfo)
    return &boxType.Ast_TypeInfo
}

// 4529: decl @lune.@base.@Ast.ProcessInfo.createSet
func (self *Ast_ProcessInfo) CreateSet(accessMode LnsInt,parentInfo *Ast_TypeInfo,itemTypeInfo *LnsList,mutMode LnsInt) *Ast_TypeInfo {
    var tmpMutMode LnsInt
    if Ast_isMutable(mutMode){
        tmpMutMode = mutMode
        
    } else { 
        tmpMutMode = Ast_MutMode__Mut
        
    }
    self.FP.Get_idProv().FP.Increment()
    var newTypeFunc func(workMutMode LnsInt) *Ast_TypeInfo
    newTypeFunc = func(workMutMode LnsInt) *Ast_TypeInfo {
        return &NewAst_NormalTypeInfo(self, false, Ast_getScope(Ast_builtinTypeSet), Ast_builtinTypeSet, nil, false, false, false, Ast_AccessMode__Pub, "Set", Ast_headTypeInfo, self.FP.Get_idProv().FP.Get_id(), Ast_TypeInfoKind__Set, itemTypeInfo, nil, nil, workMutMode).Ast_TypeInfo
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = newTypeFunc(tmpMutMode)
    if Ast_isMutable(mutMode){
        return typeInfo
    }
    return self.FP.CreateModifier(typeInfo, mutMode)
    
// insert a dummy
    return nil
}

// 4542: decl @lune.@base.@Ast.ProcessInfo.createList
func (self *Ast_ProcessInfo) CreateList(accessMode LnsInt,parentInfo *Ast_TypeInfo,itemTypeInfo *LnsList,mutMode LnsInt) *Ast_TypeInfo {
    var tmpMutMode LnsInt
    if Ast_isMutable(mutMode){
        tmpMutMode = mutMode
        
    } else { 
        tmpMutMode = Ast_MutMode__Mut
        
    }
    self.FP.Get_idProv().FP.Increment()
    var newTypeFunc func(workMutMode LnsInt) *Ast_TypeInfo
    newTypeFunc = func(workMutMode LnsInt) *Ast_TypeInfo {
        return &NewAst_NormalTypeInfo(self, false, Ast_getScope(Ast_builtinTypeList), Ast_builtinTypeList, nil, false, false, false, Ast_AccessMode__Pub, "List", Ast_headTypeInfo, self.FP.Get_idProv().FP.Get_id(), Ast_TypeInfoKind__List, itemTypeInfo, nil, nil, workMutMode).Ast_TypeInfo
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = newTypeFunc(tmpMutMode)
    if Ast_isMutable(mutMode){
        return typeInfo
    }
    return self.FP.CreateModifier(typeInfo, mutMode)
    
// insert a dummy
    return nil
}

// 4555: decl @lune.@base.@Ast.ProcessInfo.createArray
func (self *Ast_ProcessInfo) CreateArray(accessMode LnsInt,parentInfo *Ast_TypeInfo,itemTypeInfo *LnsList,mutMode LnsInt) *Ast_TypeInfo {
    var tmpMutMode LnsInt
    if Ast_isMutable(mutMode){
        tmpMutMode = mutMode
        
    } else { 
        tmpMutMode = Ast_MutMode__Mut
        
    }
    self.FP.Get_idProv().FP.Increment()
    var newTypeFunc func(workMutMode LnsInt) *Ast_TypeInfo
    newTypeFunc = func(workMutMode LnsInt) *Ast_TypeInfo {
        return &NewAst_NormalTypeInfo(self, false, Ast_getScope(Ast_builtinTypeArray), Ast_builtinTypeArray, nil, false, false, false, Ast_AccessMode__Pub, "Array", Ast_headTypeInfo, self.FP.Get_idProv().FP.Get_id(), Ast_TypeInfoKind__Array, itemTypeInfo, nil, nil, workMutMode).Ast_TypeInfo
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = newTypeFunc(tmpMutMode)
    if Ast_isMutable(mutMode){
        return typeInfo
    }
    return self.FP.CreateModifier(typeInfo, mutMode)
    
// insert a dummy
    return nil
}

// 4568: decl @lune.@base.@Ast.ProcessInfo.createMap
func (self *Ast_ProcessInfo) CreateMap(accessMode LnsInt,parentInfo *Ast_TypeInfo,keyTypeInfo *Ast_TypeInfo,valTypeInfo *Ast_TypeInfo,mutMode LnsInt) *Ast_TypeInfo {
    var tmpMutMode LnsInt
    if Ast_isMutable(mutMode){
        tmpMutMode = mutMode
        
    } else { 
        tmpMutMode = Ast_MutMode__Mut
        
    }
    self.FP.Get_idProv().FP.Increment()
    var newTypeFunc func(workMutMode LnsInt) *Ast_TypeInfo
    newTypeFunc = func(workMutMode LnsInt) *Ast_TypeInfo {
        return &NewAst_NormalTypeInfo(self, false, Ast_getScope(Ast_builtinTypeMap), Ast_builtinTypeMap, nil, false, false, false, Ast_AccessMode__Pub, "Map", Ast_headTypeInfo, self.FP.Get_idProv().FP.Get_id(), Ast_TypeInfoKind__Map, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(keyTypeInfo), Ast_TypeInfo2Stem(valTypeInfo)}), nil, nil, workMutMode).Ast_TypeInfo
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = newTypeFunc(tmpMutMode)
    if Ast_isMutable(mutMode){
        return typeInfo
    }
    return self.FP.CreateModifier(typeInfo, mutMode)
    
// insert a dummy
    return nil
}

// 4583: decl @lune.@base.@Ast.ProcessInfo.createModule
func (self *Ast_ProcessInfo) CreateModule(scope *Ast_Scope,parentInfo *Ast_TypeInfo,externalFlag bool,moduleName string,mutable bool) *Ast_TypeInfo {
    {
        __exp := Ast_sym2builtInTypeMap.Items[moduleName]
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_typeInfo()
        }
    }
    if Parser_isLuaKeyword(moduleName){
        Util_err(Lns_getVM().String_format("This symbol can not use for a class or script file. -- %s", []LnsAny{moduleName}))
    }
    var info *Ast_ModuleTypeInfo
    info = NewAst_ModuleTypeInfo(self, scope, externalFlag, moduleName, parentInfo, self.FP.Get_idProv().FP.GetNewId(), mutable)
    return &info.Ast_TypeInfo
}

// 4604: decl @lune.@base.@Ast.ProcessInfo.createClass
func (self *Ast_ProcessInfo) CreateClass(classFlag bool,abstractFlag bool,scope LnsAny,baseInfo LnsAny,interfaceList LnsAny,genTypeList *LnsList,parentInfo *Ast_TypeInfo,externalFlag bool,accessMode LnsInt,className string) *Ast_TypeInfo {
    {
        __exp := Ast_sym2builtInTypeMap.Items[className]
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_typeInfo()
        }
    }
    if Parser_isLuaKeyword(className){
        Util_err(Lns_getVM().String_format("This symbol can not use for a class or script file. -- %s", []LnsAny{className}))
    }
    var info *Ast_NormalTypeInfo
    info = NewAst_NormalTypeInfo(self, abstractFlag, scope, baseInfo, interfaceList, false, externalFlag, false, accessMode, className, parentInfo, self.FP.Get_idProv().FP.GetNewId(), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( classFlag) &&
        Lns_GetEnv().SetStackVal( Ast_TypeInfoKind__Class) ||
        Lns_GetEnv().SetStackVal( Ast_TypeInfoKind__IF) ).(LnsInt), genTypeList, nil, nil, Ast_MutMode__Mut)
    return &info.Ast_TypeInfo
}

// 4633: decl @lune.@base.@Ast.ProcessInfo.createFunc
func (self *Ast_ProcessInfo) CreateFunc(abstractFlag bool,builtinFlag bool,scope LnsAny,kind LnsInt,parentInfo *Ast_TypeInfo,autoFlag bool,externalFlag bool,staticFlag bool,accessMode LnsInt,funcName string,altTypeList LnsAny,argTypeList LnsAny,retTypeInfoList LnsAny,mutable LnsAny) *Ast_TypeInfo {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(builtinFlag)) &&
        Lns_GetEnv().SetStackVal( Parser_isLuaKeyword(funcName)) ).(bool)){
        Util_err(Lns_getVM().String_format("This symbol can not use for a function. -- %s", []LnsAny{funcName}))
    }
    var info *Ast_NormalTypeInfo
    info = NewAst_NormalTypeInfo(self, abstractFlag, scope, nil, nil, autoFlag, externalFlag, staticFlag, accessMode, funcName, parentInfo, self.FP.Get_idProv().FP.GetNewId(), kind, Lns_unwrapDefault( altTypeList, NewLnsList([]LnsAny{})).(*LnsList), Lns_unwrapDefault( argTypeList, NewLnsList([]LnsAny{})).(*LnsList), Lns_unwrapDefault( retTypeInfoList, NewLnsList([]LnsAny{})).(*LnsList), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mutable) &&
        Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
        Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt))
    return &info.Ast_TypeInfo
}

// 4654: decl @lune.@base.@Ast.ProcessInfo.createAdvertiseMethodFrom
func (self *Ast_ProcessInfo) CreateAdvertiseMethodFrom(classTypeInfo *Ast_TypeInfo,typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    return self.FP.CreateFunc(false, false, Ast_getScope(typeInfo), typeInfo.FP.Get_kind(), classTypeInfo, true, false, false, typeInfo.FP.Get_accessMode(), typeInfo.FP.Get_rawTxt(), typeInfo.FP.Get_itemTypeInfoList(), typeInfo.FP.Get_argTypeInfoList(), typeInfo.FP.Get_retTypeInfoList(), Ast_TypeInfo_isMut(typeInfo))
}

// 4683: decl @lune.@base.@Ast.ProcessInfo.createAlias
func (self *Ast_ProcessInfo) CreateAlias(processInfo *Ast_ProcessInfo,name string,externalFlag bool,accessMode LnsInt,parentInfo *Ast_TypeInfo,typeInfo *Ast_TypeInfo) *Ast_AliasTypeInfo {
    return NewAst_AliasTypeInfo(processInfo, name, accessMode, parentInfo, typeInfo.FP.Get_srcTypeInfo(), externalFlag)
}

// 4839: decl @lune.@base.@Ast.ProcessInfo.createDDD
func (self *Ast_ProcessInfo) CreateDDD(typeInfo *Ast_TypeInfo,externalFlag bool,extTypeFlag bool) *Ast_DDDTypeInfo {
    if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__DDD{
        typeInfo = typeInfo.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Ast_failCreateLuavalWith_4234_(typeInfo, Ast_LuavalConvKind__InLua, true)))) &&
        Lns_GetEnv().SetStackVal( extTypeFlag) ).(bool)){
        extTypeFlag = false
        
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_nonnilableType().FP.Get_kind() != Ast_TypeInfoKind__Ext) &&
        Lns_GetEnv().SetStackVal( extTypeFlag) ).(bool)){
        switch _exp19315 := self.FP.CreateLuaval(typeInfo, true).(type) {
        case *Ast_LuavalResult__OK:
        work := _exp19315.Val1
            typeInfo = work
            
        case *Ast_LuavalResult__Err:
        mess := _exp19315.Val1
            Util_err(mess)
        }
    }
    var dddMap *LnsMap
    if extTypeFlag{
        dddMap = self.FP.get_typeInfo2Map().ExtDDDMap
        
    } else { 
        dddMap = self.FP.get_typeInfo2Map().DDDMap
        
    }
    {
        __exp := dddMap.Items[typeInfo]
        if __exp != nil {
            _exp := __exp.(*Ast_DDDTypeInfo)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( _exp.FP.Get_typeId() < Ast_userStartId) &&
                Lns_GetEnv().SetStackVal( typeInfo.FP.Get_typeId() >= Ast_userStartId) ).(bool)){
                Util_err("on cache")
            }
            return _exp
        }
    }
    var dddType *Ast_DDDTypeInfo
    dddType = NewAst_DDDTypeInfo(self, self.FP.Get_idProv().FP.GetNewId(), typeInfo, externalFlag, nil)
    if Lns_isCondTrue( Lns_car(Ast_failCreateLuavalWith_4234_(typeInfo, Ast_LuavalConvKind__InLua, true))){
        var extDDDType *Ast_DDDTypeInfo
        extDDDType = NewAst_DDDTypeInfo(self, self.FP.Get_idProv().FP.GetNewId(), typeInfo, externalFlag, dddType)
        if extTypeFlag{
            return extDDDType
        }
    }
    return dddType
}

// 5243: decl @lune.@base.@Ast.ProcessInfo.createGeneric
func (self *Ast_ProcessInfo) CreateGeneric(genSrcTypeInfo *Ast_TypeInfo,itemTypeInfoList *LnsList,moduleTypeInfo *Ast_TypeInfo) *Ast_GenericTypeInfo {
    self.FP.Get_idProv().FP.Increment()
    return NewAst_GenericTypeInfo(self, genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo)
}

// 5518: decl @lune.@base.@Ast.ProcessInfo.createLuaval
func (self *Ast_ProcessInfo) CreateLuaval(luneType *Ast_TypeInfo,validToCheck bool) LnsAny {
    if Lns_op_not(self.validExtType){
        return &Ast_LuavalResult__OK{luneType, true}
    }
    if luneType.FP.Get_kind() == Ast_TypeInfoKind__Method{
        return &Ast_LuavalResult__OK{luneType, true}
    }
    if Ast_isExtType(luneType){
        return &Ast_LuavalResult__OK{luneType, true}
    }
    {
        _extType := self.FP.get_typeInfo2Map().ExtMap.Items[luneType]
        if _extType != nil {
            extType := _extType.(*Ast_TypeInfo)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( extType.FP.Get_typeId() < Ast_userStartId) &&
                Lns_GetEnv().SetStackVal( luneType.FP.Get_typeId() >= Ast_userStartId) ).(bool)){
                Util_err("on cache")
            }
            if extType.FP.Get_kind() == Ast_TypeInfoKind__Ext{
                return &Ast_LuavalResult__OK{extType, false}
            }
            return &Ast_LuavalResult__OK{extType, true}
        }
    }
    var process func() LnsAny
    process = func() LnsAny {
        {
            _dddType := Ast_DDDTypeInfoDownCastF(luneType.FP)
            if _dddType != nil {
                dddType := _dddType.(*Ast_DDDTypeInfo)
                switch _exp22280 := self.FP.CreateLuaval(dddType.FP.Get_typeInfo(), validToCheck).(type) {
                case *Ast_LuavalResult__Err:
                mess := _exp22280.Val1
                    Util_err(mess)
                case *Ast_LuavalResult__OK:
                workType := _exp22280.Val1
                    return &Ast_LuavalResult__OK{&self.FP.CreateDDD(workType, dddType.FP.Get_externalFlag(), true).Ast_TypeInfo, false}
                }
            }
        }
        var err LnsAny
        var canConv bool
        err, canConv = Ast_failCreateLuavalWith_4234_(luneType, Ast_LuavalConvKind__InLua, validToCheck)
        
        if err != nil{
            err_4720 := err.(string)
            return &Ast_LuavalResult__Err{err_4720}
        }
        if canConv{
            return &Ast_LuavalResult__OK{luneType, true}
        }
        self.FP.Get_idProv().FP.Increment()
        var extType *Ast_ExtTypeInfo
        extType = NewAst_ExtTypeInfo(self, luneType.FP.Get_nonnilableType())
        if luneType.FP.Get_nilable(){
            return &Ast_LuavalResult__OK{extType.FP.Get_nilableTypeInfo(), false}
        }
        return &Ast_LuavalResult__OK{&extType.Ast_TypeInfo, false}
    }
    var result LnsAny
    result = process()
    switch _exp22418 := result.(type) {
    case *Ast_LuavalResult__OK:
    typeInfo := _exp22418.Val1
        self.FP.get_typeInfo2Map().ExtMap.Set(luneType.FP.Get_nonnilableType(),typeInfo.FP.Get_nonnilableType())
        self.FP.get_typeInfo2Map().ExtMap.Set(luneType.FP.Get_nilableTypeInfo(),typeInfo.FP.Get_nilableTypeInfo())
    }
    return result
}

// 5645: decl @lune.@base.@Ast.ProcessInfo.createEnum
func (self *Ast_ProcessInfo) CreateEnum(scope *Ast_Scope,parentInfo *Ast_TypeInfo,externalFlag bool,accessMode LnsInt,enumName string,valTypeInfo *Ast_TypeInfo) *Ast_EnumTypeInfo {
    if Parser_isLuaKeyword(enumName){
        Util_err(Lns_getVM().String_format("This symbol can not use for a enum. -- %s", []LnsAny{enumName}))
    }
    var info *Ast_EnumTypeInfo
    info = NewAst_EnumTypeInfo(self, scope, externalFlag, accessMode, enumName, parentInfo, self.FP.Get_idProv().FP.GetNewId(), valTypeInfo)
    var getEnumName *Ast_TypeInfo
    getEnumName = self.FP.CreateFunc(false, true, nil, Ast_TypeInfoKind__Method, &info.Ast_TypeInfo, true, externalFlag, false, Ast_AccessMode__Pub, "get__txt", nil, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), false)
    scope.FP.AddMethod(self, nil, getEnumName, Ast_AccessMode__Pub, false, false)
    var fromVal *Ast_TypeInfo
    fromVal = self.FP.CreateFunc(false, true, nil, Ast_TypeInfoKind__Func, &info.Ast_TypeInfo, true, externalFlag, true, Ast_AccessMode__Pub, "_from", nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.FP.CreateModifier(valTypeInfo, Ast_MutMode__IMut))}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(info.FP.Get_nilableTypeInfo())}), false)
    scope.FP.AddMethod(self, nil, fromVal, Ast_AccessMode__Pub, true, false)
    var allListType *Ast_TypeInfo
    allListType = self.FP.CreateList(Ast_AccessMode__Pub, &info.Ast_TypeInfo, NewLnsList([]LnsAny{Ast_EnumTypeInfo2Stem(info)}), Ast_MutMode__IMut)
    var allList *Ast_TypeInfo
    allList = self.FP.CreateFunc(false, true, nil, Ast_TypeInfoKind__Func, &info.Ast_TypeInfo, true, externalFlag, true, Ast_AccessMode__Pub, "get__allList", nil, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.FP.CreateModifier(allListType, Ast_MutMode__IMut))}), false)
    scope.FP.AddMethod(self, nil, allList, Ast_AccessMode__Pub, true, false)
    return info
}

// 5712: decl @lune.@base.@Ast.ProcessInfo.createAlge
func (self *Ast_ProcessInfo) CreateAlge(scope *Ast_Scope,parentInfo *Ast_TypeInfo,externalFlag bool,accessMode LnsInt,algeName string) *Ast_AlgeTypeInfo {
    if Parser_isLuaKeyword(algeName){
        Util_err(Lns_getVM().String_format("This symbol can not use for a alge. -- %s", []LnsAny{algeName}))
    }
    var info *Ast_AlgeTypeInfo
    info = NewAst_AlgeTypeInfo(self, scope, externalFlag, accessMode, algeName, parentInfo, self.FP.Get_idProv().FP.GetNewId())
    var getAlgeName *Ast_TypeInfo
    getAlgeName = self.FP.CreateFunc(false, true, nil, Ast_TypeInfoKind__Method, &info.Ast_TypeInfo, true, externalFlag, false, Ast_AccessMode__Pub, "get__txt", nil, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), false)
    scope.FP.AddMethod(self, nil, getAlgeName, Ast_AccessMode__Pub, false, false)
    return info
}


type Ast_ModuleInfoIF interface {
        Get_assignName() string
        Get_modulePath() string
}
func Lns_cast2Ast_ModuleInfoIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ast_ModuleInfoIF); ok { 
        return obj
    }
    return nil
}

type Ast_ModuleInfoManager interface {
        GetModuleInfo(arg1 *Ast_TypeInfo) LnsAny
}
func Lns_cast2Ast_ModuleInfoManager( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ast_ModuleInfoManager); ok { 
        return obj
    }
    return nil
}

// declaration Class -- DummyModuleInfoManager
var Ast_DummyModuleInfoManager__instance *Ast_DummyModuleInfoManager
// 138: decl @lune.@base.@Ast.DummyModuleInfoManager.___init
func Ast_DummyModuleInfoManager____init_1141_() {
    Ast_DummyModuleInfoManager__instance = NewAst_DummyModuleInfoManager()
    
}

type Ast_DummyModuleInfoManagerMtd interface {
    GetModuleInfo(arg1 *Ast_TypeInfo) LnsAny
}
type Ast_DummyModuleInfoManager struct {
    instance *Ast_DummyModuleInfoManager
    FP Ast_DummyModuleInfoManagerMtd
}
func Ast_DummyModuleInfoManager2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_DummyModuleInfoManager).FP
}
type Ast_DummyModuleInfoManagerDownCast interface {
    ToAst_DummyModuleInfoManager() *Ast_DummyModuleInfoManager
}
func Ast_DummyModuleInfoManagerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_DummyModuleInfoManagerDownCast)
    if ok { return work.ToAst_DummyModuleInfoManager() }
    return nil
}
func (obj *Ast_DummyModuleInfoManager) ToAst_DummyModuleInfoManager() *Ast_DummyModuleInfoManager {
    return obj
}
func NewAst_DummyModuleInfoManager() *Ast_DummyModuleInfoManager {
    obj := &Ast_DummyModuleInfoManager{}
    obj.FP = obj
    obj.InitAst_DummyModuleInfoManager()
    return obj
}
func Ast_DummyModuleInfoManager_get_instance() *Ast_DummyModuleInfoManager{ return Ast_DummyModuleInfoManager__instance }
// 140: DeclConstr
func (self *Ast_DummyModuleInfoManager) InitAst_DummyModuleInfoManager() {
}

// 145: decl @lune.@base.@Ast.DummyModuleInfoManager.getModuleInfo
func (self *Ast_DummyModuleInfoManager) GetModuleInfo(typeInfo *Ast_TypeInfo) LnsAny {
    return nil
}


type Ast_LowSymbol interface {
        Get_accessMode() LnsInt
        Get_convModuleParam() LnsAny
        Get_hasAccessFromClosure() bool
        Get_isLazyLoad() bool
        Get_kind() LnsInt
        Get_mutable() bool
        Get_name() string
        Get_scope() *Ast_Scope
        Get_staticFlag() bool
        Get_symbolId() LnsInt
        Get_typeInfo() *Ast_TypeInfo
}
func Lns_cast2Ast_LowSymbol( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ast_LowSymbol); ok { 
        return obj
    }
    return nil
}

// declaration Class -- TypeNameCtrl
type Ast_TypeNameCtrlMtd interface {
    GetModuleName(arg1 *Ast_TypeInfo, arg2 string, arg3 Ast_ModuleInfoManager) string
    GetParentFullName(arg1 *Ast_TypeInfo, arg2 LnsAny, arg3 LnsAny) string
    Get_moduleTypeInfo() *Ast_TypeInfo
    Set_moduleTypeInfo(arg1 *Ast_TypeInfo)
}
type Ast_TypeNameCtrl struct {
    moduleTypeInfo *Ast_TypeInfo
    FP Ast_TypeNameCtrlMtd
}
func Ast_TypeNameCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_TypeNameCtrl).FP
}
type Ast_TypeNameCtrlDownCast interface {
    ToAst_TypeNameCtrl() *Ast_TypeNameCtrl
}
func Ast_TypeNameCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_TypeNameCtrlDownCast)
    if ok { return work.ToAst_TypeNameCtrl() }
    return nil
}
func (obj *Ast_TypeNameCtrl) ToAst_TypeNameCtrl() *Ast_TypeNameCtrl {
    return obj
}
func NewAst_TypeNameCtrl(arg1 *Ast_TypeInfo) *Ast_TypeNameCtrl {
    obj := &Ast_TypeNameCtrl{}
    obj.FP = obj
    obj.InitAst_TypeNameCtrl(arg1)
    return obj
}
func (self *Ast_TypeNameCtrl) InitAst_TypeNameCtrl(arg1 *Ast_TypeInfo) {
    self.moduleTypeInfo = arg1
}
func (self *Ast_TypeNameCtrl) Get_moduleTypeInfo() *Ast_TypeInfo{ return self.moduleTypeInfo }
func (self *Ast_TypeNameCtrl) Set_moduleTypeInfo(arg1 *Ast_TypeInfo){ self.moduleTypeInfo = arg1 }

// 1029: decl @lune.@base.@Ast.TypeNameCtrl.getModuleName
func (self *Ast_TypeNameCtrl) GetModuleName(workTypeInfo *Ast_TypeInfo,name string,moduleInfoMan Ast_ModuleInfoManager) string {
    {
        _moduleInfo := moduleInfoMan.GetModuleInfo(workTypeInfo)
        if _moduleInfo != nil {
            moduleInfo := _moduleInfo.(Ast_ModuleInfoIF)
            var txt string
            txt = moduleInfo.Get_assignName()
            return txt + "." + name
        } else {
            if self.moduleTypeInfo != workTypeInfo{
                return workTypeInfo.FP.Get_rawTxt() + "." + name
            }
        }
    }
    return name
}

// 1044: decl @lune.@base.@Ast.TypeNameCtrl.getParentFullName
func (self *Ast_TypeNameCtrl) GetParentFullName(typeInfo *Ast_TypeInfo,importInfo LnsAny,localFlag LnsAny) string {
    var workTypeInfo *Ast_TypeInfo
    workTypeInfo = typeInfo
    var name string
    name = ""
    var moduleInfoMan Ast_ModuleInfoManager
    
    {
        _moduleInfoMan := importInfo
        if _moduleInfoMan == nil{
            moduleInfoMan = Ast_DummyModuleInfoManager_get_instance().FP
            
        } else {
            moduleInfoMan = _moduleInfoMan.(Ast_ModuleInfoManager)
        }
    }
    for  {
        workTypeInfo = workTypeInfo.FP.Get_parentInfo()
        
        var txt string
        txt = workTypeInfo.FP.Get_rawTxt()
        if workTypeInfo == workTypeInfo.FP.Get_parentInfo(){
            break
        }
        if Lns_isCondTrue( localFlag){
            if workTypeInfo.FP.IsModule(){
                name = self.FP.GetModuleName(workTypeInfo, name, moduleInfoMan)
                
                break
            }
        }
        name = txt + "." + name
        
    }
    return name
}


// declaration Class -- SymbolInfo
type Ast_SymbolInfoMtd interface {
    CanAccess(arg1 *Ast_Scope, arg2 LnsInt) LnsAny
    ClearValue()
    GetModule() *Ast_TypeInfo
    GetOrg() *Ast_SymbolInfo
    Get_accessMode() LnsInt
    Get_canBeLeft() bool
    Get_canBeRight() bool
    Get_convModuleParam() LnsAny
    Get_hasAccessFromClosure() bool
    Get_hasValueFlag() bool
    Get_isLazyLoad() bool
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_mutable() bool
    Get_name() string
    Get_namespaceTypeInfo() *Ast_TypeInfo
    Get_pos() LnsAny
    Get_posForLatestMod() LnsAny
    Get_posForModToRef() LnsAny
    Get_scope() *Ast_Scope
    Get_staticFlag() bool
    Get_symbolId() LnsInt
    Get_typeInfo() *Ast_TypeInfo
    HasAccess() bool
    Set_convModuleParam(arg1 LnsAny)
    Set_hasAccessFromClosure(arg1 bool)
    Set_hasValueFlag(arg1 bool)
    Set_posForLatestMod(arg1 LnsAny)
    Set_posForModToRef(arg1 LnsAny)
    Set_typeInfo(arg1 *Ast_TypeInfo)
    UpdateValue(arg1 LnsAny)
}
type Ast_SymbolInfo struct {
    namespaceTypeInfo LnsAny
    FP Ast_SymbolInfoMtd
}
func Ast_SymbolInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_SymbolInfo).FP
}
type Ast_SymbolInfoDownCast interface {
    ToAst_SymbolInfo() *Ast_SymbolInfo
}
func Ast_SymbolInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_SymbolInfoDownCast)
    if ok { return work.ToAst_SymbolInfo() }
    return nil
}
func (obj *Ast_SymbolInfo) ToAst_SymbolInfo() *Ast_SymbolInfo {
    return obj
}
func (self *Ast_SymbolInfo) Get_accessMode() LnsInt{
// insert a dummy
    return 0
}
func (self *Ast_SymbolInfo) Get_convModuleParam() LnsAny{
// insert a dummy
    return nil
}
func (self *Ast_SymbolInfo) Get_hasAccessFromClosure() bool{
// insert a dummy
    return false
}
func (self *Ast_SymbolInfo) Get_isLazyLoad() bool{
// insert a dummy
    return false
}
func (self *Ast_SymbolInfo) Get_kind() LnsInt{
// insert a dummy
    return 0
}
func (self *Ast_SymbolInfo) Get_mutable() bool{
// insert a dummy
    return false
}
func (self *Ast_SymbolInfo) Get_name() string{
// insert a dummy
    return ""
}
func (self *Ast_SymbolInfo) Get_scope() *Ast_Scope{
// insert a dummy
    return nil
}
func (self *Ast_SymbolInfo) Get_staticFlag() bool{
// insert a dummy
    return false
}
func (self *Ast_SymbolInfo) Get_symbolId() LnsInt{
// insert a dummy
    return 0
}
func (self *Ast_SymbolInfo) Get_typeInfo() *Ast_TypeInfo{
// insert a dummy
    return nil
}
// 351: DeclConstr
func (self *Ast_SymbolInfo) InitAst_SymbolInfo() {
    self.namespaceTypeInfo = nil
    
}























// 399: decl @lune.@base.@Ast.SymbolInfo.hasAccess
func (self *Ast_SymbolInfo) HasAccess() bool {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.FP.Get_posForModToRef()) ||
        Lns_GetEnv().SetStackVal( self.FP.Get_posForLatestMod() != self.FP.Get_pos()) )){
        return true
    }
    return false
}

// 406: decl @lune.@base.@Ast.SymbolInfo.updateValue
func (self *Ast_SymbolInfo) UpdateValue(pos LnsAny) {
    self.FP.Set_hasValueFlag(true)
    self.FP.Set_posForLatestMod(pos)
}

// 410: decl @lune.@base.@Ast.SymbolInfo.clearValue
func (self *Ast_SymbolInfo) ClearValue() {
    self.FP.Set_hasValueFlag(false)
}






// 550: decl @lune.@base.@Ast.SymbolInfo.get_namespaceTypeInfo
func (self *Ast_SymbolInfo) Get_namespaceTypeInfo() *Ast_TypeInfo {
    {
        __exp := self.namespaceTypeInfo
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            return _exp
        }
    }
    var work *Ast_TypeInfo
    work = self.FP.Get_scope().FP.GetNamespaceTypeInfo()
    self.namespaceTypeInfo = work
    
    return work
}

// 972: decl @lune.@base.@Ast.SymbolInfo.getModule
func (self *Ast_SymbolInfo) GetModule() *Ast_TypeInfo {
    return self.FP.Get_namespaceTypeInfo().FP.GetModule()
}


// declaration Class -- DataOwnerInfo
type Ast_DataOwnerInfoMtd interface {
}
type Ast_DataOwnerInfo struct {
    HasData bool
    SymbolInfo *Ast_SymbolInfo
    FP Ast_DataOwnerInfoMtd
}
func Ast_DataOwnerInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_DataOwnerInfo).FP
}
type Ast_DataOwnerInfoDownCast interface {
    ToAst_DataOwnerInfo() *Ast_DataOwnerInfo
}
func Ast_DataOwnerInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_DataOwnerInfoDownCast)
    if ok { return work.ToAst_DataOwnerInfo() }
    return nil
}
func (obj *Ast_DataOwnerInfo) ToAst_DataOwnerInfo() *Ast_DataOwnerInfo {
    return obj
}
func NewAst_DataOwnerInfo(arg1 bool, arg2 *Ast_SymbolInfo) *Ast_DataOwnerInfo {
    obj := &Ast_DataOwnerInfo{}
    obj.FP = obj
    obj.InitAst_DataOwnerInfo(arg1, arg2)
    return obj
}
func (self *Ast_DataOwnerInfo) InitAst_DataOwnerInfo(arg1 bool, arg2 *Ast_SymbolInfo) {
    self.HasData = arg1
    self.SymbolInfo = arg2
}

// declaration Class -- Scope
type Ast_ScopeMtd interface {
    AccessSymbol(arg1 *Ast_Scope, arg2 *Ast_SymbolInfo)
    Add(arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 bool, arg4 bool, arg5 string, arg6 LnsAny, arg7 *Ast_TypeInfo, arg8 LnsInt, arg9 bool, arg10 LnsInt, arg11 bool, arg12 bool)(LnsAny, LnsAny)
    AddAlge(arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 string, arg4 LnsAny, arg5 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddAlgeVal(arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddAlias(arg1 *Ast_ProcessInfo, arg2 string, arg3 *Types_Position, arg4 bool, arg5 LnsInt, arg6 *Ast_TypeInfo, arg7 *Ast_SymbolInfo)(LnsAny, LnsAny)
    AddAliasForType(arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddAlternate(arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 string, arg4 LnsAny, arg5 *Ast_TypeInfo)
    AddClass(arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddClassLazy(arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 bool)(LnsAny, LnsAny)
    AddEnum(arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 string, arg4 LnsAny, arg5 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddEnumVal(arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddForm(arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *Ast_TypeInfo, arg4 LnsInt)
    AddFunc(arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *Ast_TypeInfo, arg4 LnsInt, arg5 bool, arg6 bool)(LnsAny, LnsAny)
    AddIgnoredVar(arg1 *Ast_ProcessInfo)
    AddLocalVar(arg1 *Ast_ProcessInfo, arg2 bool, arg3 bool, arg4 string, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 LnsInt)(LnsAny, LnsAny)
    AddMacro(arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *Ast_TypeInfo, arg4 LnsInt)(LnsAny, LnsAny)
    AddMember(arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 LnsInt, arg6 bool, arg7 LnsInt)(LnsAny, LnsAny)
    AddMethod(arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *Ast_TypeInfo, arg4 LnsInt, arg5 bool, arg6 bool)(LnsAny, LnsAny)
    AddModule(arg1 *Ast_TypeInfo, arg2 Ast_ModuleInfoIF)
    AddOverrideImut(arg1 *Ast_SymbolInfo)
    AddStaticVar(arg1 *Ast_ProcessInfo, arg2 bool, arg3 bool, arg4 string, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 LnsInt)(LnsAny, LnsAny)
    AddSymbol(arg1 *Ast_SymbolInfo)
    AddUnwrapedVar(arg1 *Ast_ProcessInfo, arg2 bool, arg3 bool, arg4 string, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 LnsInt)(LnsAny, LnsAny)
    AddVar(arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 string, arg4 LnsAny, arg5 *Ast_TypeInfo, arg6 LnsInt, arg7 bool)(LnsAny, LnsAny)
    FilterSymbolInfoIfField(arg1 *Ast_Scope, arg2 LnsInt, arg3 Ast_filterForm) bool
    FilterSymbolTypeInfo(arg1 *Ast_Scope, arg2 *Ast_Scope, arg3 LnsInt, arg4 Ast_filterForm)
    FilterTypeInfoField(arg1 LnsAny, arg2 *Ast_Scope, arg3 LnsInt, arg4 Ast_filterForm) bool
    GetClassTypeInfo() *Ast_TypeInfo
    GetModule() *Ast_TypeInfo
    GetModuleInfo(arg1 *Ast_TypeInfo) LnsAny
    GetNamespaceTypeInfo() *Ast_TypeInfo
    GetProcessInfo() *Ast_ProcessInfo
    GetSymbolInfo(arg1 string, arg2 *Ast_Scope, arg3 bool, arg4 LnsInt) LnsAny
    GetSymbolInfoChild(arg1 string) LnsAny
    GetSymbolInfoField(arg1 string, arg2 LnsAny, arg3 *Ast_Scope, arg4 LnsInt) LnsAny
    GetSymbolInfoIfField(arg1 string, arg2 *Ast_Scope, arg3 LnsInt) LnsAny
    GetSymbolTypeInfo(arg1 string, arg2 *Ast_Scope, arg3 *Ast_Scope, arg4 LnsInt) LnsAny
    GetTypeInfo(arg1 string, arg2 *Ast_Scope, arg3 bool, arg4 LnsInt) LnsAny
    GetTypeInfoChild(arg1 string) LnsAny
    GetTypeInfoField(arg1 string, arg2 LnsAny, arg3 *Ast_Scope, arg4 LnsInt) LnsAny
    Get_closureSym2NumMap() *LnsMap
    Get_closureSymList() *LnsList
    Get_closureSymMap() *LnsMap
    Get_inherit() LnsAny
    Get_ownerTypeInfo() LnsAny
    Get_parent() *Ast_Scope
    Get_scopeId() LnsInt
    Get_symbol2SymbolInfoMap() *LnsMap
    Get_validCheckingUnaccess() bool
    IsClosureAccess(arg1 *Ast_Scope, arg2 *Ast_SymbolInfo) bool
    IsInnerOf(arg1 *Ast_Scope) bool
    IsRoot() bool
    Remove(arg1 string)
    SetClosure(arg1 *Ast_SymbolInfo)
    SetData(arg1 *Ast_SymbolInfo)
    Set_ownerTypeInfo(arg1 LnsAny)
    Set_validCheckingUnaccess(arg1 bool)
    SwitchOwnerTypeInfo(arg1 LnsAny)
}
type Ast_Scope struct {
    scopeId LnsInt
    ownerTypeInfo LnsAny
    parent *Ast_Scope
    symbol2SymbolInfoMap *LnsMap
    classFlag bool
    inherit LnsAny
    ifScopeList *LnsList
    symbolId2DataOwnerInfo *LnsMap
    closureSymMap *LnsMap
    closureSymList *LnsList
    closureSym2NumMap *LnsMap
    typeInfo2ModuleInfoMap *LnsMap
    validCheckingUnaccess bool
    FP Ast_ScopeMtd
}
func Ast_Scope2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_Scope).FP
}
type Ast_ScopeDownCast interface {
    ToAst_Scope() *Ast_Scope
}
func Ast_ScopeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_ScopeDownCast)
    if ok { return work.ToAst_Scope() }
    return nil
}
func (obj *Ast_Scope) ToAst_Scope() *Ast_Scope {
    return obj
}
func NewAst_Scope(arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 bool, arg4 LnsAny, arg5 LnsAny) *Ast_Scope {
    obj := &Ast_Scope{}
    obj.FP = obj
    obj.InitAst_Scope(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Ast_Scope) Get_scopeId() LnsInt{ return self.scopeId }
func (self *Ast_Scope) Get_ownerTypeInfo() LnsAny{ return self.ownerTypeInfo }
func (self *Ast_Scope) Get_parent() *Ast_Scope{ return self.parent }
func (self *Ast_Scope) Get_symbol2SymbolInfoMap() *LnsMap{ return self.symbol2SymbolInfoMap }
func (self *Ast_Scope) Get_inherit() LnsAny{ return self.inherit }
func (self *Ast_Scope) Get_closureSymMap() *LnsMap{ return self.closureSymMap }
func (self *Ast_Scope) Get_closureSymList() *LnsList{ return self.closureSymList }
func (self *Ast_Scope) Get_closureSym2NumMap() *LnsMap{ return self.closureSym2NumMap }
func (self *Ast_Scope) Get_validCheckingUnaccess() bool{ return self.validCheckingUnaccess }
func (self *Ast_Scope) Set_validCheckingUnaccess(arg1 bool){ self.validCheckingUnaccess = arg1 }
// 474: DeclConstr
func (self *Ast_Scope) InitAst_Scope(processInfo *Ast_ProcessInfo,parent LnsAny,classFlag bool,inherit LnsAny,ifScopeList LnsAny) {
    self.scopeId = processInfo.FP.Get_idProvScope().FP.GetNewId()
    
    self.typeInfo2ModuleInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.closureSymMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.closureSym2NumMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.closureSymList = NewLnsList([]LnsAny{})
    
    self.parent = Lns_unwrapDefault( parent, self).(*Ast_Scope)
    
    self.symbol2SymbolInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.inherit = inherit
    
    self.classFlag = classFlag
    
    self.symbolId2DataOwnerInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.ifScopeList = Lns_unwrapDefault( ifScopeList, NewLnsList([]LnsAny{})).(*LnsList)
    
    self.ownerTypeInfo = nil
    
    self.validCheckingUnaccess = true
    
}


// 497: decl @lune.@base.@Ast.Scope.isRoot
func (self *Ast_Scope) IsRoot() bool {
    return self.parent == self
}

// 501: decl @lune.@base.@Ast.Scope.set_ownerTypeInfo
func (self *Ast_Scope) Set_ownerTypeInfo(owner LnsAny) {
    if Lns_op_not(self.ownerTypeInfo){
        self.ownerTypeInfo = owner
        
    }
}

// 506: decl @lune.@base.@Ast.Scope.switchOwnerTypeInfo
func (self *Ast_Scope) SwitchOwnerTypeInfo(owner LnsAny) {
    self.ownerTypeInfo = owner
    
}

// 511: decl @lune.@base.@Ast.Scope.getTypeInfoChild
func (self *Ast_Scope) GetTypeInfoChild(name string) LnsAny {
    {
        __exp := self.symbol2SymbolInfoMap.Items[name]
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_typeInfo()
        }
    }
    return nil
}

// 518: decl @lune.@base.@Ast.Scope.getSymbolInfoChild
func (self *Ast_Scope) GetSymbolInfoChild(name string) LnsAny {
    return self.symbol2SymbolInfoMap.Items[name]
}

// 522: decl @lune.@base.@Ast.Scope.setData
func (self *Ast_Scope) SetData(symbolInfo *Ast_SymbolInfo) {
    self.symbolId2DataOwnerInfo.Set(symbolInfo.FP.Get_symbolId(),NewAst_DataOwnerInfo(true, symbolInfo))
}

// 527: decl @lune.@base.@Ast.Scope.remove
func (self *Ast_Scope) Remove(name string) {
    self.symbol2SymbolInfoMap.Set(name,nil)
}

// 531: decl @lune.@base.@Ast.Scope.addSymbol
func (self *Ast_Scope) AddSymbol(symbolInfo *Ast_SymbolInfo) {
    self.symbol2SymbolInfoMap.Set(symbolInfo.FP.Get_name(),symbolInfo)
}

// 535: decl @lune.@base.@Ast.Scope.addModule
func (self *Ast_Scope) AddModule(typeInfo *Ast_TypeInfo,moduleInfo Ast_ModuleInfoIF) {
    self.typeInfo2ModuleInfoMap.Set(typeInfo,moduleInfo)
}

// 539: decl @lune.@base.@Ast.Scope.getModuleInfo
func (self *Ast_Scope) GetModuleInfo(typeInfo *Ast_TypeInfo) LnsAny {
    {
        _moduleInfo := self.typeInfo2ModuleInfoMap.Items[typeInfo]
        if _moduleInfo != nil {
            moduleInfo := _moduleInfo.(Ast_ModuleInfoIF)
            return moduleInfo
        }
    }
    if self.parent != self{
        return self.parent.FP.GetModuleInfo(typeInfo)
    }
    return nil
}

// 567: decl @lune.@base.@Ast.Scope.isInnerOf
func (self *Ast_Scope) IsInnerOf(scope *Ast_Scope) bool {
    var workScope *Ast_Scope
    workScope = self
    for workScope != Ast_rootScope {
        if workScope == scope{
            return true
        }
        workScope = workScope.parent
        
    }
    return false
}

// 1126: decl @lune.@base.@Ast.Scope.getNamespaceTypeInfo
func (self *Ast_Scope) GetNamespaceTypeInfo() *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var scope *Ast_Scope
    scope = self
    for {
        {
            __exp := scope.FP.Get_ownerTypeInfo()
            if __exp != nil {
                _exp := __exp.(*Ast_TypeInfo)
                return _exp
            }
        }
        scope = scope.FP.Get_parent()
        
        if scope.FP.IsRoot(){ break }
    }
    return typeInfo
}

// 1138: decl @lune.@base.@Ast.Scope.getModule
func (self *Ast_Scope) GetModule() *Ast_TypeInfo {
    return self.FP.GetNamespaceTypeInfo().FP.GetModule()
}

// 1142: decl @lune.@base.@Ast.Scope.getProcessInfo
func (self *Ast_Scope) GetProcessInfo() *Ast_ProcessInfo {
    return self.FP.GetModule().FP.getProcessInfo()
}

// 1614: decl @lune.@base.@Ast.Scope.filterTypeInfoField
func (self *Ast_Scope) FilterTypeInfoField(includeSelfFlag LnsAny,fromScope *Ast_Scope,access LnsInt,callback Ast_filterForm) bool {
    if self.classFlag{
        if Lns_isCondTrue( includeSelfFlag){
            {
                __collection4911 := self.symbol2SymbolInfoMap
                __sorted4911 := __collection4911.CreateKeyListStr()
                __sorted4911.Sort( LnsItemKindStr, nil )
                for _, ___key4911 := range( __sorted4911.Items ) {
                    symbolInfo := __collection4911.Items[ ___key4911 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    if Lns_isCondTrue( symbolInfo.FP.CanAccess(fromScope, access)){
                        if Lns_op_not(callback(symbolInfo)){
                            return false
                        }
                    }
                }
            }
        }
        {
            _scope := self.inherit
            if _scope != nil {
                scope := _scope.(*Ast_Scope)
                if Lns_op_not(scope.FP.FilterTypeInfoField(true, fromScope, access, callback)){
                    return false
                }
            }
        }
    }
    return true
}

// 1643: decl @lune.@base.@Ast.Scope.getSymbolInfoField
func (self *Ast_Scope) GetSymbolInfoField(name string,includeSelfFlag LnsAny,fromScope *Ast_Scope,access LnsInt) LnsAny {
    if self.classFlag{
        if Lns_isCondTrue( includeSelfFlag){
            {
                __exp := self.symbol2SymbolInfoMap.Items[name]
                if __exp != nil {
                    _exp := __exp.(*Ast_SymbolInfo)
                    var symbolInfo *Ast_SymbolInfo
                    
                    {
                        _symbolInfo := _exp.FP.CanAccess(fromScope, access)
                        if _symbolInfo == nil{
                            return nil
                        } else {
                            symbolInfo = _symbolInfo.(*Ast_SymbolInfo)
                        }
                    }
                    return symbolInfo
                }
            }
        }
        {
            _scope := self.inherit
            if _scope != nil {
                scope := _scope.(*Ast_Scope)
                var symbolInfo LnsAny
                symbolInfo = scope.FP.GetSymbolInfoField(name, true, fromScope, access)
                if Lns_isCondTrue( symbolInfo){
                    return symbolInfo
                }
            }
        }
    }
    return nil
}

// 1673: decl @lune.@base.@Ast.Scope.getSymbolInfoIfField
func (self *Ast_Scope) GetSymbolInfoIfField(name string,fromScope *Ast_Scope,access LnsInt) LnsAny {
    if self.classFlag{
        for _, _scope := range( self.ifScopeList.Items ) {
            scope := _scope.(Ast_ScopeDownCast).ToAst_Scope()
            {
                _symbolInfo := scope.FP.GetSymbolInfoField(name, true, fromScope, access)
                if _symbolInfo != nil {
                    symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                    return symbolInfo
                }
            }
        }
    }
    {
        _scope := self.inherit
        if _scope != nil {
            scope := _scope.(*Ast_Scope)
            {
                _symbolInfo := scope.FP.GetSymbolInfoIfField(name, fromScope, access)
                if _symbolInfo != nil {
                    symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                    return symbolInfo
                }
            }
        }
    }
    return nil
}

// 1700: decl @lune.@base.@Ast.Scope.filterSymbolInfoIfField
func (self *Ast_Scope) FilterSymbolInfoIfField(fromScope *Ast_Scope,access LnsInt,callback Ast_filterForm) bool {
    for _, _scope := range( self.ifScopeList.Items ) {
        scope := _scope.(Ast_ScopeDownCast).ToAst_Scope()
        if Lns_op_not(scope.FP.FilterTypeInfoField(true, fromScope, access, callback)){
            return false
        }
    }
    {
        _scope := self.inherit
        if _scope != nil {
            scope := _scope.(*Ast_Scope)
            if Lns_op_not(scope.FP.FilterSymbolInfoIfField(fromScope, access, callback)){
                return false
            }
        }
    }
    return true
}

// 1722: decl @lune.@base.@Ast.Scope.getTypeInfoField
func (self *Ast_Scope) GetTypeInfoField(name string,includeSelfFlag LnsAny,fromScope *Ast_Scope,access LnsInt) LnsAny {
    var symbolInfo LnsAny
    symbolInfo = self.FP.GetSymbolInfoField(name, includeSelfFlag, fromScope, access)
    {
        __exp := symbolInfo
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_typeInfo()
        }
    }
    return nil
}

// 1743: decl @lune.@base.@Ast.Scope.getSymbolInfo
func (self *Ast_Scope) GetSymbolInfo(name string,fromScope *Ast_Scope,onlySameNsFlag bool,access LnsInt) LnsAny {
    {
        __exp := self.symbol2SymbolInfoMap.Items[name]
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            var symbolInfo *Ast_SymbolInfo
            
            {
                _symbolInfo := _exp.FP.CanAccess(fromScope, access)
                if _symbolInfo == nil{
                    return nil
                } else {
                    symbolInfo = _symbolInfo.(*Ast_SymbolInfo)
                }
            }
            return symbolInfo
        }
    }
    if Lns_op_not(onlySameNsFlag){
        {
            _scope := self.inherit
            if _scope != nil {
                scope := _scope.(*Ast_Scope)
                var symbolInfo LnsAny
                symbolInfo = scope.FP.GetSymbolInfoField(name, true, fromScope, access)
                if Lns_isCondTrue( symbolInfo){
                    return symbolInfo
                }
            }
        }
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(onlySameNsFlag)) ||
        Lns_GetEnv().SetStackVal( Lns_op_not(self.ownerTypeInfo)) ).(bool){
        if self.parent != self{
            return self.parent.FP.GetSymbolInfo(name, fromScope, onlySameNsFlag, access)
        }
    } else { 
        var workScope *Ast_Scope
        workScope = self.parent
        for workScope.parent != workScope {
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workScope.ownerTypeInfo) && 
                Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_kind()})) != Ast_TypeInfoKind__Class) &&
                Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(workScope.ownerTypeInfo) && 
                Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_kind()})) != Ast_TypeInfoKind__Module) ).(bool)){
                return workScope.FP.GetSymbolInfo(name, fromScope, onlySameNsFlag, access)
            }
            workScope = workScope.parent
            
        }
    }
    {
        __exp := Ast_sym2builtInTypeMap.Items[name]
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp
        }
    }
    return nil
}

// 1804: decl @lune.@base.@Ast.Scope.getTypeInfo
func (self *Ast_Scope) GetTypeInfo(name string,fromScope *Ast_Scope,onlySameNsFlag bool,access LnsInt) LnsAny {
    var symbolInfo *Ast_SymbolInfo
    
    {
        _symbolInfo := self.FP.GetSymbolInfo(name, fromScope, onlySameNsFlag, access)
        if _symbolInfo == nil{
            return nil
        } else {
            symbolInfo = _symbolInfo.(*Ast_SymbolInfo)
        }
    }
    return symbolInfo.FP.Get_typeInfo()
}

// 1821: decl @lune.@base.@Ast.Scope.getSymbolTypeInfo
func (self *Ast_Scope) GetSymbolTypeInfo(name string,fromScope *Ast_Scope,moduleScope *Ast_Scope,access LnsInt) LnsAny {
    var validThisScope bool
    validThisScope = false
    var limitSymbol bool
    limitSymbol = false
    {
        __exp := self.ownerTypeInfo
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__Func) ||
                Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__Method) ||
                Lns_GetEnv().SetStackVal( self == moduleScope) ||
                Lns_GetEnv().SetStackVal( self == Ast_rootScope) ).(bool){
                validThisScope = true
                
            } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__IF) ||
                Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__Module) ).(bool){
                limitSymbol = true
                
                validThisScope = true
                
            } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__Enum) ||
                Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__Alge) ).(bool){
                validThisScope = true
                
            }
        } else {
            validThisScope = true
            
        }
    }
    if validThisScope{
        {
            _symbolInfo := self.symbol2SymbolInfoMap.Items[name]
            if _symbolInfo != nil {
                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_op_not(limitSymbol)) ||
                    Lns_GetEnv().SetStackVal( name == "self") ||
                    Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_typeInfo().FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
                        Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Typ) ).(bool))) ).(bool){
                    return symbolInfo.FP.CanAccess(fromScope, access)
                }
            }
        }
    }
    if self.parent != self{
        return self.parent.FP.GetSymbolTypeInfo(name, fromScope, moduleScope, access)
    }
    return Ast_sym2builtInTypeMap.Items[name]
}

// 1864: decl @lune.@base.@Ast.Scope.filterSymbolTypeInfo
func (self *Ast_Scope) FilterSymbolTypeInfo(fromScope *Ast_Scope,moduleScope *Ast_Scope,access LnsInt,callback Ast_filterForm) {
    if self.classFlag{
        {
            __exp := self.symbol2SymbolInfoMap.Items["self"]
            if __exp != nil {
                _exp := __exp.(*Ast_SymbolInfo)
                callback(_exp)
            }
        }
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( moduleScope == fromScope) ||
        Lns_GetEnv().SetStackVal( moduleScope == self) ||
        Lns_GetEnv().SetStackVal( Lns_op_not(self.classFlag)) ).(bool){
        for _, _symbolInfo := range( self.symbol2SymbolInfoMap.Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if Lns_op_not(callback(symbolInfo)){
                return 
            }
        }
    }
    if self.parent != self{
        self.parent.FP.FilterSymbolTypeInfo(fromScope, moduleScope, access, callback)
    }
}

// 1886: decl @lune.@base.@Ast.Scope.add
func (self *Ast_Scope) Add(processInfo *Ast_ProcessInfo,kind LnsInt,canBeLeft bool,canBeRight bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,mutMode LnsInt,hasValueFlag bool,isLazyLoad bool)(LnsAny, LnsAny) {
    if _switch5961 := kind; _switch5961 == Ast_SymbolKind__Typ || _switch5961 == Ast_SymbolKind__Fun || _switch5961 == Ast_SymbolKind__Mac {
        var existSymbol LnsAny
        if _switch5927 := typeInfo.FP.Get_kind(); _switch5927 == Ast_TypeInfoKind__Enum {
            if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.ownerTypeInfo) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_kind()})) == Ast_TypeInfoKind__Class{
                existSymbol = self.FP.GetSymbolInfoField(name, true, self, Ast_ScopeAccess__Full)
                
            } else { 
                existSymbol = self.FP.GetSymbolInfo(name, self, true, Ast_ScopeAccess__Full)
                
            }
        } else if _switch5927 == Ast_TypeInfoKind__Class || _switch5927 == Ast_TypeInfoKind__Module {
            existSymbol = self.FP.GetSymbolInfoChild(name)
            
        } else {
            existSymbol = self.FP.GetSymbolInfo(name, self, true, Ast_ScopeAccess__Full)
            
        }
        if existSymbol != nil{
            existSymbol_1867 := existSymbol.(*Ast_SymbolInfo)
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( typeInfo.FP.Get_kind() != existSymbol_1867.FP.Get_typeInfo().FP.Get_kind()) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isBuiltin(existSymbol_1867.FP.Get_typeInfo().FP.Get_typeId()))) ).(bool){
                return nil, existSymbol_1867
            }
        }
    }
    var symbolInfo *Ast_NormalSymbolInfo
    symbolInfo = NewAst_NormalSymbolInfo(processInfo, kind, canBeLeft, canBeRight, self, accessMode, staticFlag, name, pos, typeInfo, mutMode, hasValueFlag, isLazyLoad)
    self.symbol2SymbolInfoMap.Set(name,&symbolInfo.Ast_SymbolInfo)
    return &symbolInfo.Ast_SymbolInfo, nil
}

// 1928: decl @lune.@base.@Ast.Scope.addLocalVar
func (self *Ast_Scope) AddLocalVar(processInfo *Ast_ProcessInfo,argFlag bool,canBeLeft bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutable LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( argFlag) &&
        Lns_GetEnv().SetStackVal( Ast_SymbolKind__Arg) ||
        Lns_GetEnv().SetStackVal( Ast_SymbolKind__Var) ).(LnsInt), canBeLeft, name != "_", name, pos, typeInfo, Ast_AccessMode__Local, false, mutable, true, false)
}

// 1941: decl @lune.@base.@Ast.Scope.addUnwrapedVar
func (self *Ast_Scope) AddUnwrapedVar(processInfo *Ast_ProcessInfo,argFlag bool,canBeLeft bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutable LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( argFlag) &&
        Lns_GetEnv().SetStackVal( Ast_SymbolKind__Arg) ||
        Lns_GetEnv().SetStackVal( Ast_SymbolKind__Var) ).(LnsInt), canBeLeft, true, name, pos, typeInfo, Ast_AccessMode__Local, false, mutable, true, false)
}

// 1952: decl @lune.@base.@Ast.Scope.addStaticVar
func (self *Ast_Scope) AddStaticVar(processInfo *Ast_ProcessInfo,argFlag bool,canBeLeft bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutable LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( argFlag) &&
        Lns_GetEnv().SetStackVal( Ast_SymbolKind__Arg) ||
        Lns_GetEnv().SetStackVal( Ast_SymbolKind__Var) ).(LnsInt), canBeLeft, true, name, pos, typeInfo, Ast_AccessMode__Local, true, mutable, true, false)
}

// 1962: decl @lune.@base.@Ast.Scope.addVar
func (self *Ast_Scope) AddVar(processInfo *Ast_ProcessInfo,accessMode LnsInt,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutable LnsInt,hasValueFlag bool)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Var, true, true, name, pos, typeInfo, accessMode, false, mutable, hasValueFlag, false)
}

// 1976: decl @lune.@base.@Ast.Scope.addEnumVal
func (self *Ast_Scope) AddEnumVal(processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Mbr, false, true, name, pos, typeInfo, Ast_AccessMode__Pub, true, Ast_MutMode__Mut, true, false)
}

// 1986: decl @lune.@base.@Ast.Scope.addEnum
func (self *Ast_Scope) AddEnum(processInfo *Ast_ProcessInfo,accessMode LnsInt,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Typ, false, false, name, pos, typeInfo, accessMode, true, Ast_MutMode__Mut, true, false)
}

// 1996: decl @lune.@base.@Ast.Scope.addAlgeVal
func (self *Ast_Scope) AddAlgeVal(processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Mbr, false, true, name, pos, typeInfo, Ast_AccessMode__Pub, true, Ast_MutMode__Mut, true, false)
}

// 2006: decl @lune.@base.@Ast.Scope.addAlge
func (self *Ast_Scope) AddAlge(processInfo *Ast_ProcessInfo,accessMode LnsInt,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Typ, false, false, name, pos, typeInfo, accessMode, true, Ast_MutMode__Mut, true, false)
}

// 2016: decl @lune.@base.@Ast.Scope.addAlternate
func (self *Ast_Scope) AddAlternate(processInfo *Ast_ProcessInfo,accessMode LnsInt,name string,pos LnsAny,typeInfo *Ast_TypeInfo) {
    self.FP.Add(processInfo, Ast_SymbolKind__Typ, false, false, name, pos, typeInfo, accessMode, true, Ast_MutMode__Mut, true, false)
}

// 2025: decl @lune.@base.@Ast.Scope.addMember
func (self *Ast_Scope) AddMember(processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,mutMode LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Mbr, true, true, name, pos, typeInfo, accessMode, staticFlag, mutMode, true, false)
}

// 2035: decl @lune.@base.@Ast.Scope.addMethod
func (self *Ast_Scope) AddMethod(processInfo *Ast_ProcessInfo,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,mutable bool)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Mtd, true, staticFlag, typeInfo.FP.Get_rawTxt(), pos, typeInfo, accessMode, staticFlag, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mutable) &&
        Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
        Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, false)
}

// 2045: decl @lune.@base.@Ast.Scope.addFunc
func (self *Ast_Scope) AddFunc(processInfo *Ast_ProcessInfo,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,mutable bool)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Fun, true, true, typeInfo.FP.Get_rawTxt(), pos, typeInfo, accessMode, staticFlag, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mutable) &&
        Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
        Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, false)
}

// 2056: decl @lune.@base.@Ast.Scope.addForm
func (self *Ast_Scope) AddForm(processInfo *Ast_ProcessInfo,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt) {
    self.FP.Add(processInfo, Ast_SymbolKind__Typ, false, false, typeInfo.FP.Get_rawTxt(), pos, typeInfo, accessMode, true, Ast_MutMode__IMut, false, false)
}

// 2065: decl @lune.@base.@Ast.Scope.addMacro
func (self *Ast_Scope) AddMacro(processInfo *Ast_ProcessInfo,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Mac, false, false, typeInfo.FP.Get_rawTxt(), pos, typeInfo, accessMode, true, Ast_MutMode__IMut, true, false)
}

// 2076: decl @lune.@base.@Ast.Scope.addClassLazy
func (self *Ast_Scope) AddClassLazy(processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo,lazyLoad bool)(LnsAny, LnsAny) {
    return self.FP.Add(processInfo, Ast_SymbolKind__Typ, false, false, name, pos, typeInfo, typeInfo.FP.Get_accessMode(), true, Ast_MutMode__Mut, true, lazyLoad)
}

// 2085: decl @lune.@base.@Ast.Scope.addClass
func (self *Ast_Scope) AddClass(processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.AddClassLazy(processInfo, name, pos, typeInfo, false)
}

// 2127: decl @lune.@base.@Ast.Scope.setClosure
func (self *Ast_Scope) SetClosure(workSymbol *Ast_SymbolInfo) {
    var symbol *Ast_SymbolInfo
    symbol = workSymbol.FP.GetOrg()
    var targetFuncType *Ast_TypeInfo
    targetFuncType = symbol.FP.Get_namespaceTypeInfo()
    var funcType *Ast_TypeInfo
    funcType = self.FP.GetNamespaceTypeInfo()
    for  {
        var funcScope *Ast_Scope
        funcScope = Lns_unwrap( funcType.FP.Get_scope()).(*Ast_Scope)
        if Lns_op_not(funcScope.closureSymMap.Items[symbol.FP.Get_symbolId()]){
            funcScope.closureSymMap.Set(symbol.FP.Get_symbolId(),symbol)
            funcScope.closureSym2NumMap.Set(symbol,funcScope.closureSymList.Len())
            funcScope.closureSymList.Insert(Ast_SymbolInfo2Stem(symbol))
            funcType = funcScope.parent.FP.GetNamespaceTypeInfo()
            
        } else { 
            break
        }
        if funcType == targetFuncType{
            break
        }
    }
    if Lns_op_not(symbol.FP.Get_hasAccessFromClosure()){
        symbol.FP.Set_hasAccessFromClosure(true)
    }
}

// 2156: decl @lune.@base.@Ast.Scope.isClosureAccess
func (self *Ast_Scope) IsClosureAccess(moduleScope *Ast_Scope,symbol *Ast_SymbolInfo) bool {
    var processInfo *Ast_ProcessInfo
    processInfo = moduleScope.FP.GetModule().FP.getProcessInfo()
    if _switch7356 := symbol.FP.Get_kind(); _switch7356 == Ast_SymbolKind__Var || _switch7356 == Ast_SymbolKind__Arg || _switch7356 == Ast_SymbolKind__Fun {
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symbol.FP.Get_scope() == moduleScope) ||
            Lns_GetEnv().SetStackVal( symbol.FP.Get_scope() == Ast_rootScope) ).(bool){
        } else if symbol.FP.Get_name() == "self"{
            var funcType *Ast_TypeInfo
            funcType = self.FP.GetNamespaceTypeInfo()
            if funcType.FP.Get_parentInfo().FP.IsInheritFrom(processInfo, symbol.FP.Get_namespaceTypeInfo().FP.Get_parentInfo(), nil){
            } else { 
                return true
            }
        } else { 
            var funcType *Ast_TypeInfo
            funcType = self.FP.GetNamespaceTypeInfo()
            if funcType != symbol.FP.Get_namespaceTypeInfo(){
                return true
            }
        }
    }
    return false
}

// 2196: decl @lune.@base.@Ast.Scope.accessSymbol
func (self *Ast_Scope) AccessSymbol(moduleScope *Ast_Scope,symbol *Ast_SymbolInfo) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( symbol.FP.Get_kind() == Ast_SymbolKind__Fun) &&
        Lns_GetEnv().SetStackVal( self.FP.GetNamespaceTypeInfo() == symbol.FP.Get_typeInfo()) ).(bool)){
        return 
    }
    if self.FP.IsClosureAccess(moduleScope, symbol){
        self.FP.SetClosure(symbol)
    }
}

// 2321: decl @lune.@base.@Ast.Scope.getClassTypeInfo
func (self *Ast_Scope) GetClassTypeInfo() *Ast_TypeInfo {
    var scope *Ast_Scope
    scope = self
    for  {
        {
            _owner := scope.ownerTypeInfo
            if _owner != nil {
                owner := _owner.(*Ast_TypeInfo)
                if _switch7741 := owner.FP.Get_kind(); _switch7741 == Ast_TypeInfoKind__Class || _switch7741 == Ast_TypeInfoKind__IF || _switch7741 == Ast_TypeInfoKind__Module {
                    return owner
                }
            }
        }
        if scope.parent == scope{
            break
        }
        scope = scope.parent
        
    }
    return Ast_headTypeInfo
}

// 4126: decl @lune.@base.@Ast.Scope.addOverrideImut
func (self *Ast_Scope) AddOverrideImut(symbolInfo *Ast_SymbolInfo) {
    var typeInfo *Ast_TypeInfo
    if Ast_TypeInfo_isMut(symbolInfo.FP.Get_typeInfo()){
        var processInfo *Ast_ProcessInfo
        processInfo = self.FP.GetProcessInfo()
        typeInfo = processInfo.FP.CreateModifier(symbolInfo.FP.Get_typeInfo(), Ast_MutMode__IMut)
        
    } else { 
        typeInfo = symbolInfo.FP.Get_typeInfo()
        
    }
    self.symbol2SymbolInfoMap.Set(symbolInfo.FP.Get_name(),&NewAst_AccessSymbolInfo(symbolInfo, &Ast_OverrideMut__IMut{typeInfo}, false).Ast_SymbolInfo)
}

// 4275: decl @lune.@base.@Ast.Scope.addIgnoredVar
func (self *Ast_Scope) AddIgnoredVar(processInfo *Ast_ProcessInfo) {
    self.FP.AddLocalVar(processInfo, false, true, "_", nil, Ast_builtinTypeEmpty, Ast_MutMode__Mut)
}

// 4692: decl @lune.@base.@Ast.Scope.addAlias
func (self *Ast_Scope) AddAlias(processInfo *Ast_ProcessInfo,name string,pos *Types_Position,externalFlag bool,accessMode LnsInt,parentInfo *Ast_TypeInfo,symbolInfo *Ast_SymbolInfo)(LnsAny, LnsAny) {
    var aliasType *Ast_AliasTypeInfo
    aliasType = self.FP.GetProcessInfo().FP.CreateAlias(processInfo, name, externalFlag, accessMode, parentInfo, symbolInfo.FP.Get_typeInfo().FP.Get_srcTypeInfo())
    return self.FP.Add(processInfo, symbolInfo.FP.Get_kind(), false, symbolInfo.FP.Get_canBeRight(), name, pos, &aliasType.Ast_TypeInfo, accessMode, true, Ast_MutMode__IMut, true, false)
}

// 4705: decl @lune.@base.@Ast.Scope.addAliasForType
func (self *Ast_Scope) AddAliasForType(processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    var skind LnsInt
    skind = Ast_SymbolKind__Typ
    var canBeRight bool
    canBeRight = false
    if _switch18767 := typeInfo.FP.Get_kind(); _switch18767 == Ast_TypeInfoKind__Func {
        skind = Ast_SymbolKind__Fun
        
        canBeRight = true
        
    } else if _switch18767 == Ast_TypeInfoKind__Form || _switch18767 == Ast_TypeInfoKind__FormFunc {
        canBeRight = true
        
    } else if _switch18767 == Ast_TypeInfoKind__Macro {
        skind = Ast_SymbolKind__Mac
        
    }
    return self.FP.Add(processInfo, skind, false, canBeRight, name, pos, typeInfo, typeInfo.FP.Get_accessMode(), true, Ast_MutMode__IMut, true, false)
}


// declaration Class -- TypeData
type Ast_TypeDataMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    Get_children() *LnsList
}
type Ast_TypeData struct {
    children *LnsList
    FP Ast_TypeDataMtd
}
func Ast_TypeData2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_TypeData).FP
}
type Ast_TypeDataDownCast interface {
    ToAst_TypeData() *Ast_TypeData
}
func Ast_TypeDataDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_TypeDataDownCast)
    if ok { return work.ToAst_TypeData() }
    return nil
}
func (obj *Ast_TypeData) ToAst_TypeData() *Ast_TypeData {
    return obj
}
func NewAst_TypeData(arg1 *LnsList) *Ast_TypeData {
    obj := &Ast_TypeData{}
    obj.FP = obj
    obj.InitAst_TypeData(arg1)
    return obj
}
func (self *Ast_TypeData) InitAst_TypeData(arg1 *LnsList) {
    self.children = arg1
}
func (self *Ast_TypeData) Get_children() *LnsList{ return self.children }
// 587: decl @lune.@base.@Ast.TypeData.addChildren
func (self *Ast_TypeData) AddChildren(child *Ast_TypeInfo) {
    self.children.Insert(Ast_TypeInfo2Stem(child))
}


// declaration Class -- TypeInfo
type Ast_TypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_TypeInfo struct {
    scope LnsAny
    typeData *Ast_TypeData
    processInfo *Ast_ProcessInfo
    FP Ast_TypeInfoMtd
}
func Ast_TypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_TypeInfo).FP
}
type Ast_TypeInfoDownCast interface {
    ToAst_TypeInfo() *Ast_TypeInfo
}
func Ast_TypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_TypeInfoDownCast)
    if ok { return work.ToAst_TypeInfo() }
    return nil
}
func (obj *Ast_TypeInfo) ToAst_TypeInfo() *Ast_TypeInfo {
    return obj
}
func (self *Ast_TypeInfo) Get_scope() LnsAny{ return self.scope }
func (self *Ast_TypeInfo) Get_typeData() *Ast_TypeData{ return self.typeData }
func (self *Ast_TypeInfo) Get_processInfo() *Ast_ProcessInfo{ return self.processInfo }
// 641: decl @lune.@base.@Ast.TypeInfo.getProcessInfo
func (self *Ast_TypeInfo) getProcessInfo() *Ast_ProcessInfo {
    return self.processInfo
}

// 645: decl @lune.@base.@Ast.TypeInfo.switchScope
func (self *Ast_TypeInfo) SwitchScope(scope *Ast_Scope) {
    self.scope = scope
    
    scope.FP.SwitchOwnerTypeInfo(self)
}

// 655: decl @lune.@base.@Ast.TypeInfo.getOverridingType
func (self *Ast_TypeInfo) GetOverridingType() LnsAny {
    return nil
}

// 663: DeclConstr
func (self *Ast_TypeInfo) InitAst_TypeInfo(scope LnsAny,processInfo *Ast_ProcessInfo) {
    self.scope = scope
    
    {
        __exp := scope
        if __exp != nil {
            _exp := __exp.(*Ast_Scope)
            _exp.FP.Set_ownerTypeInfo(self)
        }
    }
    self.typeData = NewAst_TypeData(NewLnsList([]LnsAny{}))
    
    self.processInfo = processInfo
    
}

// 674: decl @lune.@base.@Ast.TypeInfo.get_aliasSrc
func (self *Ast_TypeInfo) Get_aliasSrc() *Ast_TypeInfo {
    return self
}

// 678: decl @lune.@base.@Ast.TypeInfo.get_extedType
func (self *Ast_TypeInfo) Get_extedType() *Ast_TypeInfo {
    return self
}

// 682: decl @lune.@base.@Ast.TypeInfo.getModulePath
func Ast_TypeInfo_getModulePath(fullname string) string {
    return Lns_car(Lns_getVM().String_gsub(fullname,"@", "")).(string)
}

// 687: decl @lune.@base.@Ast.TypeInfo.isModule
func (self *Ast_TypeInfo) IsModule() bool {
    return true
}

// 692: decl @lune.@base.@Ast.TypeInfo.getParentId
func (self *Ast_TypeInfo) GetParentId() LnsInt {
    return Ast_rootTypeId
}

// 697: decl @lune.@base.@Ast.TypeInfo.get_baseId
func (self *Ast_TypeInfo) Get_baseId() LnsInt {
    return Ast_rootTypeId
}

// 702: decl @lune.@base.@Ast.TypeInfo.isInheritFrom
func (self *Ast_TypeInfo) IsInheritFrom(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    return false
}

// 709: decl @lune.@base.@Ast.TypeInfo.get_rawTxt
func (self *Ast_TypeInfo) Get_rawTxt() string {
    return ""
}

// 719: decl @lune.@base.@Ast.TypeInfo.getTxtWithRaw
func (self *Ast_TypeInfo) GetTxtWithRaw(raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return ""
}

// 732: decl @lune.@base.@Ast.TypeInfo.getTxt
func (self *Ast_TypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}


// 744: decl @lune.@base.@Ast.TypeInfo.canEvalWith
func (self *Ast_TypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return false, nil
}

// 752: decl @lune.@base.@Ast.TypeInfo.get_abstractFlag
func (self *Ast_TypeInfo) Get_abstractFlag() bool {
    return false
}

// 757: decl @lune.@base.@Ast.TypeInfo.serialize
func (self *Ast_TypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    return 
}

// 763: decl @lune.@base.@Ast.TypeInfo.get_display_stirng_with
func (self *Ast_TypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    return ""
}

// 768: decl @lune.@base.@Ast.TypeInfo.get_display_stirng
func (self *Ast_TypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with("", nil)
}

// 774: decl @lune.@base.@Ast.TypeInfo.get_srcTypeInfo
func (self *Ast_TypeInfo) Get_srcTypeInfo() *Ast_TypeInfo {
    return self
}

// 779: decl @lune.@base.@Ast.TypeInfo.equals
func (self *Ast_TypeInfo) Equals(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if Lns_isCondTrue( checkModifer){
        return self == typeInfo
    }
    return self == typeInfo.FP.Get_srcTypeInfo()
}

// 790: decl @lune.@base.@Ast.TypeInfo.get_externalFlag
func (self *Ast_TypeInfo) Get_externalFlag() bool {
    return false
}

// 795: decl @lune.@base.@Ast.TypeInfo.get_interfaceList
func (self *Ast_TypeInfo) Get_interfaceList() *LnsList {
    return Ast_dummyList
}

// 799: decl @lune.@base.@Ast.TypeInfo.get_itemTypeInfoList
func (self *Ast_TypeInfo) Get_itemTypeInfoList() *LnsList {
    return Ast_dummyList
}

// 803: decl @lune.@base.@Ast.TypeInfo.get_argTypeInfoList
func (self *Ast_TypeInfo) Get_argTypeInfoList() *LnsList {
    return Ast_dummyList
}

// 807: decl @lune.@base.@Ast.TypeInfo.get_retTypeInfoList
func (self *Ast_TypeInfo) Get_retTypeInfoList() *LnsList {
    return Ast_dummyList
}


// 817: decl @lune.@base.@Ast.TypeInfo.hasRouteNamespaceFrom
func (self *Ast_TypeInfo) HasRouteNamespaceFrom(other *Ast_TypeInfo) bool {
    for  {
        if other == self{
            return true
        }
        if other.FP.Get_parentInfo() == other{
            break
        }
        other = other.FP.Get_parentInfo()
        
    }
    return false
}

// 830: decl @lune.@base.@Ast.TypeInfo.getModule
func (self *Ast_TypeInfo) GetModule() *Ast_TypeInfo {
    if self.FP.IsModule(){
        return self
    }
    return self.FP.Get_parentInfo().FP.GetModule()
}

// 837: decl @lune.@base.@Ast.TypeInfo.get_typeId
func (self *Ast_TypeInfo) Get_typeId() LnsInt {
    return Ast_rootTypeId
}

// 841: decl @lune.@base.@Ast.TypeInfo.get_kind
func (self *Ast_TypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Root
}

// 849: decl @lune.@base.@Ast.TypeInfo.get_staticFlag
func (self *Ast_TypeInfo) Get_staticFlag() bool {
    return false
}

// 853: decl @lune.@base.@Ast.TypeInfo.get_accessMode
func (self *Ast_TypeInfo) Get_accessMode() LnsInt {
    return Ast_AccessMode__Pri
}

// 857: decl @lune.@base.@Ast.TypeInfo.get_autoFlag
func (self *Ast_TypeInfo) Get_autoFlag() bool {
    return false
}

// 861: decl @lune.@base.@Ast.TypeInfo.get_nonnilableType
func (self *Ast_TypeInfo) Get_nonnilableType() *Ast_TypeInfo {
    return self
}


// 868: decl @lune.@base.@Ast.TypeInfo.get_nilable
func (self *Ast_TypeInfo) Get_nilable() bool {
    return false
}

// 872: decl @lune.@base.@Ast.TypeInfo.get_nilableTypeInfo
func (self *Ast_TypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo {
    return self
}

// 880: decl @lune.@base.@Ast.TypeInfo.get_children
func (self *Ast_TypeInfo) Get_children() *LnsList {
    return self.typeData.FP.Get_children()
}

// 884: decl @lune.@base.@Ast.TypeInfo.addChildren
func (self *Ast_TypeInfo) AddChildren(child *Ast_TypeInfo) {
    self.typeData.FP.AddChildren(child)
}

// 889: decl @lune.@base.@Ast.TypeInfo.get_mutMode
func (self *Ast_TypeInfo) Get_mutMode() LnsInt {
    return Ast_MutMode__Mut
}

// 893: decl @lune.@base.@Ast.TypeInfo.isMut
func Ast_TypeInfo_isMut(typeInfo *Ast_TypeInfo) bool {
    return Ast_isMutable(typeInfo.FP.Get_mutMode())
}

// 898: decl @lune.@base.@Ast.TypeInfo.getParentFullName
func (self *Ast_TypeInfo) GetParentFullName(typeNameCtrl *Ast_TypeNameCtrl,importInfo LnsAny,localFlag LnsAny) string {
    return typeNameCtrl.FP.GetParentFullName(self, importInfo, localFlag)
}

// 906: decl @lune.@base.@Ast.TypeInfo.applyGeneric
func (self *Ast_TypeInfo) ApplyGeneric(alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    return self
}

// 912: decl @lune.@base.@Ast.TypeInfo.get_genSrcTypeInfo
func (self *Ast_TypeInfo) Get_genSrcTypeInfo() *Ast_TypeInfo {
    return self
}

// 916: decl @lune.@base.@Ast.TypeInfo.serializeTypeInfoList
func (self *Ast_TypeInfo) SerializeTypeInfoList(name string,list *LnsList,onlyPub LnsAny) string {
    var work string
    work = name
    for _, _typeInfo := range( list.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(onlyPub)) ||
            Lns_GetEnv().SetStackVal( typeInfo.FP.Get_accessMode() == Ast_AccessMode__Pub) ).(bool){
            if len(work) != len(name){
                work = work + ", "
                
            }
            work = Lns_getVM().String_format("%s%d", []LnsAny{work, typeInfo.FP.Get_typeId()})
            
        }
    }
    return work + "}, "
}

// 931: decl @lune.@base.@Ast.TypeInfo.createScope
func Ast_TypeInfo_createScope(processInfo *Ast_ProcessInfo,parent LnsAny,classFlag bool,baseInfo LnsAny,interfaceList LnsAny) *Ast_Scope {
    var inheritScope LnsAny
    inheritScope = nil
    if baseInfo != nil{
        baseInfo_1071 := baseInfo.(*Ast_TypeInfo)
        inheritScope = Lns_unwrap( baseInfo_1071.scope).(*Ast_Scope)
        
    }
    var ifScopeList *LnsList
    ifScopeList = NewLnsList([]LnsAny{})
    if interfaceList != nil{
        interfaceList_1074 := interfaceList.(*LnsList)
        for _, _ifType := range( interfaceList_1074.Items ) {
            ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            ifScopeList.Insert(Ast_Scope2Stem(Lns_unwrap( ifType.scope).(*Ast_Scope)))
        }
    }
    return NewAst_Scope(processInfo, parent, classFlag, inheritScope, ifScopeList)
}



// 1121: decl @lune.@base.@Ast.TypeInfo.hasBase
func (self *Ast_TypeInfo) HasBase() bool {
    return self.FP.Get_baseTypeInfo() != Ast_headTypeInfo
}

// 1241: decl @lune.@base.@Ast.TypeInfo.isInherit
func Ast_TypeInfo_isInherit(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    var baseTypeInfo *Ast_TypeInfo
    baseTypeInfo = typeInfo.FP.Get_baseTypeInfo()
    var otherTypeId LnsInt
    otherTypeId = other.FP.Get_typeId()
    if typeInfo.FP.Get_typeId() == otherTypeId{
        return true
    }
    if baseTypeInfo != Ast_headTypeInfo{
        if baseTypeInfo.FP.IsInheritFrom(processInfo, other, alt2type){
            return true
        }
    }
    for _, _ifType := range( typeInfo.FP.Get_interfaceList().Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if ifType.FP.IsInheritFrom(processInfo, other, alt2type){
            return true
        }
    }
    return false
}

// 2216: decl @lune.@base.@Ast.TypeInfo.createAlt2typeMap
func (self *Ast_TypeInfo) CreateAlt2typeMap(detectFlag bool) *LnsMap {
    return Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(detectFlag)
}

// 5047: decl @lune.@base.@Ast.TypeInfo.getCommonTypeCombo
func Ast_TypeInfo_getCommonTypeCombo(processInfo *Ast_ProcessInfo,commonType LnsAny,otherType LnsAny,alt2type *LnsMap) LnsAny {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_builtinTypeNone
    switch _exp20173 := commonType.(type) {
    case *Ast_CommonType__Combine:
    comb := _exp20173.Val1
        return comb.FP.AndType(processInfo, otherType, alt2type)
    case *Ast_CommonType__Normal:
    workTypeInfo := _exp20173.Val1
        typeInfo = workTypeInfo
        
    }
    var other *Ast_TypeInfo
    other = Ast_builtinTypeNone
    switch _exp20204 := otherType.(type) {
    case *Ast_CommonType__Combine:
    comb := _exp20204.Val1
        return comb.FP.AndType(processInfo, commonType, alt2type)
    case *Ast_CommonType__Normal:
    workTypeInfo := _exp20204.Val1
        other = workTypeInfo
        
    }
    var getType func(workType *Ast_TypeInfo) LnsAny
    getType = func(workType *Ast_TypeInfo) LnsAny {
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( typeInfo.FP.Get_nilable()) ||
            Lns_GetEnv().SetStackVal( other.FP.Get_nilable()) ).(bool){
            workType = workType.FP.Get_nilableTypeInfo()
            
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(typeInfo))) ||
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(other))) ).(bool){
            workType = processInfo.FP.CreateModifier(workType, Ast_MutMode__IMut)
            
        }
        return &Ast_CommonType__Normal{workType}
    }
    var type1 *Ast_TypeInfo
    type1 = typeInfo.FP.Get_nonnilableType().FP.Get_srcTypeInfo()
    var type2 *Ast_TypeInfo
    type2 = other.FP.Get_nonnilableType().FP.Get_srcTypeInfo()
    if type1 == Ast_builtinTypeNone{
        return otherType
    }
    if type2 == Ast_builtinTypeNone{
        return commonType
    }
    if type1 == Ast_builtinTypeNil{
        return &Ast_CommonType__Normal{other.FP.Get_nilableTypeInfo()}
    }
    if type2 == Ast_builtinTypeNil{
        return &Ast_CommonType__Normal{typeInfo.FP.Get_nilableTypeInfo()}
    }
    if Lns_car(type1.FP.CanEvalWith(processInfo, type2, Ast_CanEvalType__SetOp, alt2type)).(bool){
        return getType(type1)
    }
    if Lns_car(type2.FP.CanEvalWith(processInfo, type1, Ast_CanEvalType__SetOp, alt2type)).(bool){
        return getType(type2)
    }
    var mutMode LnsInt
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(typeInfo)) &&
        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(other)) ).(bool)){
        mutMode = Ast_MutMode__Mut
        
    } else { 
        mutMode = Ast_MutMode__IMut
        
    }
    if type1.FP.Get_kind() == type2.FP.Get_kind(){
        var getCommon func(workTypeInfo *Ast_TypeInfo,workOther *Ast_TypeInfo,workAlt2type *LnsMap) *Ast_TypeInfo
        getCommon = func(workTypeInfo *Ast_TypeInfo,workOther *Ast_TypeInfo,workAlt2type *LnsMap) *Ast_TypeInfo {
            switch _exp20483 := Ast_TypeInfo_getCommonTypeCombo(processInfo, &Ast_CommonType__Normal{workTypeInfo}, &Ast_CommonType__Normal{workOther}, workAlt2type).(type) {
            case *Ast_CommonType__Normal:
            info := _exp20483.Val1
                return info
            case *Ast_CommonType__Combine:
            combine := _exp20483.Val1
                return combine.FP.Get_typeInfo(processInfo)
            }
        // insert a dummy
            return nil
        }
        if _switch20673 := type1.FP.Get_kind(); _switch20673 == Ast_TypeInfoKind__List {
            return getType(processInfo.FP.CreateList(Ast_AccessMode__Local, Ast_headTypeInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(getCommon(type1.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type))}), mutMode))
        } else if _switch20673 == Ast_TypeInfoKind__Array {
            return getType(processInfo.FP.CreateArray(Ast_AccessMode__Local, Ast_headTypeInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(getCommon(type1.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type))}), mutMode))
        } else if _switch20673 == Ast_TypeInfoKind__Set {
            return getType(processInfo.FP.CreateSet(Ast_AccessMode__Local, Ast_headTypeInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(getCommon(type1.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type))}), mutMode))
        } else if _switch20673 == Ast_TypeInfoKind__Map {
            return getType(processInfo.FP.CreateMap(Ast_AccessMode__Local, Ast_headTypeInfo, getCommon(type1.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type), getCommon(type1.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type), mutMode))
        }
    }
    var work *Ast_TypeInfo
    work = type1.FP.Get_baseTypeInfo()
    for work != Ast_headTypeInfo {
        if Lns_car(work.FP.CanEvalWith(processInfo, type2, Ast_CanEvalType__SetOp, alt2type)).(bool){
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( typeInfo.FP.Get_nilable()) ||
                Lns_GetEnv().SetStackVal( other.FP.Get_nilable()) ).(bool){
                work = work.FP.Get_nilableTypeInfo()
                
            }
            if Lns_op_not(Ast_isMutable(mutMode)){
                work = processInfo.FP.CreateModifier(work, mutMode)
                
            }
            return &Ast_CommonType__Normal{work}
        }
        work = work.FP.Get_baseTypeInfo()
        
    }
    var combine *Ast_CombineType
    combine = NewAst_CombineType(typeInfo)
    return combine.FP.AndType(processInfo, &Ast_CommonType__Normal{other}, alt2type)
}

// 5192: decl @lune.@base.@Ast.TypeInfo.getCommonType
func Ast_TypeInfo_getCommonType(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,other *Ast_TypeInfo,alt2type *LnsMap) *Ast_TypeInfo {
    switch _exp20845 := Ast_TypeInfo_getCommonTypeCombo(processInfo, &Ast_CommonType__Normal{typeInfo}, &Ast_CommonType__Normal{other}, alt2type).(type) {
    case *Ast_CommonType__Normal:
    info := _exp20845.Val1
        return info
    case *Ast_CommonType__Combine:
    combine := _exp20845.Val1
        return combine.FP.Get_typeInfo(processInfo)
    }
// insert a dummy
    return nil
}

// 5968: decl @lune.@base.@Ast.TypeInfo.checkMatchType
func Ast_TypeInfo_checkMatchType(processInfo *Ast_ProcessInfo,dstTypeList *LnsList,expTypeList *LnsList,allowDstShort bool,warnForFollowSrcIndex LnsAny,alt2type *LnsMap)(LnsInt, string) {
    var warnMess LnsAny
    warnMess = nil
    var checkDstTypeFrom func(index LnsInt,srcType *Ast_TypeInfo,srcType2nd *Ast_TypeInfo)(LnsInt, string)
    checkDstTypeFrom = func(index LnsInt,srcType *Ast_TypeInfo,srcType2nd *Ast_TypeInfo)(LnsInt, string) {
        var workExpType *Ast_TypeInfo
        workExpType = srcType
        {
            var _from24364 LnsInt = index
            var _to24364 LnsInt = dstTypeList.Len()
            for _work24364 := _from24364; _work24364 <= _to24364; _work24364++ {
                dstIndex := _work24364
                var workDstType *Ast_TypeInfo
                workDstType = dstTypeList.GetAt(dstIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if Lns_op_not(Lns_car(workDstType.FP.CanEvalWith(processInfo, workExpType, Ast_CanEvalType__SetOp, alt2type)).(bool)){
                    var message string
                    message = Lns_getVM().String_format("exp(%d) type mismatch %s <- %s: dst %d", []LnsAny{dstIndex, workDstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), workExpType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), dstIndex})
                    return Ast_MatchType__Error, message
                } else if workExpType == Ast_builtinTypeAbbrNone{
                    return Ast_MatchType__Warn, Code_format(Code_ID__nothing_define_abbr, Lns_getVM().String_format("use '##', instate of '%s'.", []LnsAny{workDstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil)}))
                }
                workExpType = srcType2nd
                
            }
        }
        return Ast_MatchType__Match, ""
    }
    var checkSrcTypeFrom func(index LnsInt,dstType *Ast_TypeInfo)(LnsInt, string)
    checkSrcTypeFrom = func(index LnsInt,dstType *Ast_TypeInfo)(LnsInt, string) {
        {
            var _from24620 LnsInt = index
            var _to24620 LnsInt = expTypeList.Len()
            for _work24620 := _from24620; _work24620 <= _to24620; _work24620++ {
                srcIndex := _work24620
                var expType *Ast_TypeInfo
                expType = expTypeList.GetAt(srcIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var checkType *Ast_TypeInfo
                checkType = expType
                if expType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                    if expType.FP.Get_itemTypeInfoList().Len() > 0{
                        checkType = expType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        
                    } else { 
                        checkType = Ast_builtinTypeStem_
                        
                    }
                } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( srcIndex > index) &&
                    Lns_GetEnv().SetStackVal( expType.FP.Get_kind() == Ast_TypeInfoKind__Abbr) ).(bool)){
                    return Ast_MatchType__Error, "must not use '##'"
                }
                if Lns_op_not(Lns_car(dstType.FP.CanEvalWith(processInfo, checkType, Ast_CanEvalType__SetOp, alt2type)).(bool)){
                    return Ast_MatchType__Error, Lns_getVM().String_format("exp(%d) type mismatch %s <- %s: src: %d", []LnsAny{srcIndex, dstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), srcIndex})
                }
                if warnForFollowSrcIndex != nil{
                    warnForFollowSrcIndex_5012 := warnForFollowSrcIndex.(LnsInt)
                    if warnForFollowSrcIndex_5012 <= srcIndex{
                        var workMess string
                        workMess = Lns_getVM().String_format("use '**' at arg(%d). %s <- %s", []LnsAny{srcIndex, dstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil)})
                        return Ast_MatchType__Warn, workMess
                    }
                }
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isExtType(dstType.FP.Get_srcTypeInfo().FP.Get_nonnilableType()))) &&
                    Lns_GetEnv().SetStackVal( Ast_isExtType(expType.FP.Get_srcTypeInfo().FP.Get_nonnilableType())) ).(bool)){
                    warnMess = Lns_getVM().String_format("exp(%d) luaval mismatch %s <- %s", []LnsAny{index, dstType.FP.GetTxt(nil, nil, nil), expType.FP.GetTxt(nil, nil, nil)})
                    
                }
                
            }
        }
        return Ast_MatchType__Match, ""
    }
    if expTypeList.Len() > 0{
        for _index, _expType := range( expTypeList.Items ) {
            index := _index + 1
            expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if dstTypeList.Len() == 0{
                return Ast_MatchType__Error, Lns_getVM().String_format("over exp. expect:0, actual:%d", []LnsAny{expTypeList.Len()})
            }
            var dstType *Ast_TypeInfo
            dstType = dstTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if dstTypeList.Len() == index{
                if dstType.FP.Get_srcTypeInfo().FP.Get_kind() != Ast_TypeInfoKind__DDD{
                    var isMatch bool
                    var msg LnsAny
                    isMatch,msg = dstType.FP.CanEvalWith(processInfo, expType, Ast_CanEvalType__SetOp, alt2type)
                    if Lns_op_not(isMatch){
                        return Ast_MatchType__Error, Lns_getVM().String_format("exp(%d) type mismatch %s <- %s: index %d%s", []LnsAny{index, Ast_applyGenericDefault(dstType, alt2type, dstType.FP.GetModule()).FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), index, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( msg) &&
                            Lns_GetEnv().SetStackVal( Lns_getVM().String_format(" -- %s", []LnsAny{msg})) ||
                            Lns_GetEnv().SetStackVal( "") ).(string)})
                    }
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(allowDstShort)) &&
                        Lns_GetEnv().SetStackVal( dstTypeList.Len() < expTypeList.Len()) ).(bool)){
                        return Ast_MatchType__Error, Lns_getVM().String_format("over exp. expect: %d: actual: %d", []LnsAny{dstTypeList.Len(), expTypeList.Len()})
                    }
                } else { 
                    var dddItemType *Ast_TypeInfo
                    dddItemType = Ast_builtinTypeStem_
                    if dstType.FP.Get_itemTypeInfoList().Len() > 0{
                        dddItemType = dstType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        
                    }
                    var result LnsInt
                    var mess string
                    result,mess = checkSrcTypeFrom(index, dddItemType)
                    if result != Ast_MatchType__Match{
                        return result, mess
                    }
                }
                if warnForFollowSrcIndex != nil{
                    warnForFollowSrcIndex_5035 := warnForFollowSrcIndex.(LnsInt)
                    if warnForFollowSrcIndex_5035 <= index{
                        var workMess string
                        workMess = Lns_getVM().String_format("use '**' at arg(%d). %s <- %s", []LnsAny{index, dstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil)})
                        return Ast_MatchType__Warn, workMess
                    }
                }
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isExtType(dstType.FP.Get_srcTypeInfo().FP.Get_nonnilableType()))) &&
                    Lns_GetEnv().SetStackVal( Ast_isExtType(expType.FP.Get_srcTypeInfo().FP.Get_nonnilableType())) ).(bool)){
                    warnMess = Lns_getVM().String_format("exp(%d) luaval mismatch %s <- %s", []LnsAny{index, dstType.FP.GetTxt(nil, nil, nil), expType.FP.GetTxt(nil, nil, nil)})
                    
                }
                
                break
            } else if expTypeList.Len() == index{
                var srcType *Ast_TypeInfo
                srcType = expType
                var srcType2nd *Ast_TypeInfo
                srcType2nd = Ast_builtinTypeAbbrNone
                if expType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                    srcType = Ast_builtinTypeStem_
                    
                    srcType2nd = Ast_builtinTypeStem_
                    
                    if expType.FP.Get_itemTypeInfoList().Len() > 0{
                        srcType = expType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        
                        srcType2nd = srcType
                        
                    }
                } else if expType == Ast_builtinTypeAbbr{
                    srcType2nd = Ast_builtinTypeAbbr
                    
                }
                var result LnsInt
                var mess string
                result,mess = checkDstTypeFrom(index, srcType, srcType2nd)
                if result != Ast_MatchType__Match{
                    return result, mess
                }
                if warnForFollowSrcIndex != nil{
                    warnForFollowSrcIndex_5049 := warnForFollowSrcIndex.(LnsInt)
                    if warnForFollowSrcIndex_5049 <= index{
                        var workMess string
                        workMess = Lns_getVM().String_format("use '**' at arg(%d). %s <- %s", []LnsAny{index, dstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil)})
                        return Ast_MatchType__Warn, workMess
                    }
                }
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isExtType(dstType.FP.Get_srcTypeInfo().FP.Get_nonnilableType()))) &&
                    Lns_GetEnv().SetStackVal( Ast_isExtType(expType.FP.Get_srcTypeInfo().FP.Get_nonnilableType())) ).(bool)){
                    warnMess = Lns_getVM().String_format("exp(%d) luaval mismatch %s <- %s", []LnsAny{index, dstType.FP.GetTxt(nil, nil, nil), expType.FP.GetTxt(nil, nil, nil)})
                    
                }
                
                break
            } else { 
                if Lns_op_not(Lns_car(dstType.FP.CanEvalWith(processInfo, expType, Ast_CanEvalType__SetOp, alt2type)).(bool)){
                    return Ast_MatchType__Error, Lns_getVM().String_format("exp(%d) type mismatch %s(%d) <- %s(%d)", []LnsAny{index, dstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), dstType.FP.Get_typeId(), expType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), expType.FP.Get_typeId()})
                }
                if warnForFollowSrcIndex != nil{
                    warnForFollowSrcIndex_5056 := warnForFollowSrcIndex.(LnsInt)
                    if warnForFollowSrcIndex_5056 <= index{
                        var workMess string
                        workMess = Lns_getVM().String_format("use '**' at arg(%d). %s <- %s", []LnsAny{index, dstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil)})
                        return Ast_MatchType__Warn, workMess
                    }
                }
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isExtType(dstType.FP.Get_srcTypeInfo().FP.Get_nonnilableType()))) &&
                    Lns_GetEnv().SetStackVal( Ast_isExtType(expType.FP.Get_srcTypeInfo().FP.Get_nonnilableType())) ).(bool)){
                    warnMess = Lns_getVM().String_format("exp(%d) luaval mismatch %s <- %s", []LnsAny{index, dstType.FP.GetTxt(nil, nil, nil), expType.FP.GetTxt(nil, nil, nil)})
                    
                }
                
            }
        }
    } else if Lns_op_not(allowDstShort){
        for _index, _dstType := range( dstTypeList.Items ) {
            index := _index + 1
            dstType := _dstType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(Lns_car(dstType.FP.CanEvalWith(processInfo, Ast_builtinTypeNil, Ast_CanEvalType__SetOp, alt2type)).(bool)){
                return Ast_MatchType__Error, Lns_getVM().String_format("exp(%d) type mismatch %s <- nil: short", []LnsAny{index, dstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil)})
            }
            return Ast_MatchType__Warn, Code_format(Code_ID__nothing_define_abbr, Lns_getVM().String_format("use '##', instate of '%s'.", []LnsAny{dstType.FP.GetTxt(Ast_defaultTypeNameCtrl, nil, nil)}))
        }
    }
    if warnMess != nil{
        warnMess_5066 := warnMess.(string)
        return Ast_MatchType__Warn, warnMess_5066
    }
    return Ast_MatchType__Match, ""
}

// 6173: decl @lune.@base.@Ast.TypeInfo.canEvalWithBase
func Ast_TypeInfo_canEvalWithBase(processInfo *Ast_ProcessInfo,dest *Ast_TypeInfo,destMut bool,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if dest != dest.FP.Get_aliasSrc(){
        return dest.FP.Get_aliasSrc().FP.CanEvalWith(processInfo, other, canEvalType, alt2type)
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( dest == Ast_builtinTypeExp) ||
        Lns_GetEnv().SetStackVal( dest == Ast_builtinTypeMultiExp) ).(bool){
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( other == Ast_builtinTypeMultiExp) &&
            Lns_GetEnv().SetStackVal( dest != Ast_builtinTypeMultiExp) ).(bool)){
            return false, "can't eval from '__exp' to '__exps'."
        }
        if other.FP.Equals(processInfo, Ast_builtinTypeAbbr, nil, nil){
            return false, nil
        }
        return true, nil
    }
    var otherMut bool
    otherMut = Ast_TypeInfo_isMut(other)
    var otherSrc *Ast_TypeInfo
    otherSrc = other.FP.Get_srcTypeInfo().FP.Get_aliasSrc()
    if otherSrc.FP.Get_kind() == Ast_TypeInfoKind__DDD{
        if otherSrc.FP.Get_itemTypeInfoList().Len() > 0{
            otherSrc = otherSrc.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo()
            
        } else { 
            otherSrc = Ast_builtinTypeStem_
            
        }
    }
    if otherSrc.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
        if dest.FP.Get_kind() != Ast_TypeInfoKind__Alternate{
            otherSrc = Ast_AlternateTypeInfo_getAssign(otherSrc, alt2type).FP.Get_srcTypeInfo()
            
        }
    }
    if _switch25941 := canEvalType; _switch25941 == Ast_CanEvalType__SetEq || _switch25941 == Ast_CanEvalType__SetOp || _switch25941 == Ast_CanEvalType__SetOpIMut {
        if dest == Ast_builtinTypeEmpty{
            if _switch25748 := otherSrc; _switch25748 == Ast_builtinTypeAbbr || _switch25748 == Ast_builtinTypeAbbrNone {
                return false, nil
            }
            if otherSrc.FP.Get_kind() == Ast_TypeInfoKind__Func{
                return false, nil
            }
            return true, nil
        } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(otherMut)) &&
            Lns_GetEnv().SetStackVal( destMut) ).(bool)){
            var nonNilOtherType *Ast_TypeInfo
            nonNilOtherType = otherSrc.FP.Get_nonnilableType()
            if _switch25837 := nonNilOtherType.FP.Get_kind(); _switch25837 == Ast_TypeInfoKind__Set || _switch25837 == Ast_TypeInfoKind__Map || _switch25837 == Ast_TypeInfoKind__List || _switch25837 == Ast_TypeInfoKind__IF || _switch25837 == Ast_TypeInfoKind__Alternate {
                return false, nil
            } else if _switch25837 == Ast_TypeInfoKind__Class {
                if Ast_builtinTypeString != nonNilOtherType{
                    return false, nil
                }
            } else if _switch25837 == Ast_TypeInfoKind__Prim {
                if Ast_builtinTypeStem == nonNilOtherType{
                    return false, nil
                }
            }
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( otherSrc != Ast_builtinTypeNil) &&
            Lns_GetEnv().SetStackVal( otherSrc != Ast_builtinTypeString) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() != Ast_TypeInfoKind__Prim) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() != Ast_TypeInfoKind__Func) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() != Ast_TypeInfoKind__Method) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() != Ast_TypeInfoKind__Form) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() != Ast_TypeInfoKind__FormFunc) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() != Ast_TypeInfoKind__Enum) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() != Ast_TypeInfoKind__Abbr) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() != Ast_TypeInfoKind__Alternate) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() != Ast_TypeInfoKind__Box) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isGenericType(otherSrc))) &&
            Lns_GetEnv().SetStackVal( destMut) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(otherMut)) ).(bool)){
            return false, nil
        }
    }
    if dest == Ast_builtinTypeStem_{
        return true, nil
    }
    if dest.FP.Get_srcTypeInfo().FP.Get_kind() == Ast_TypeInfoKind__DDD{
        if dest.FP.Get_itemTypeInfoList().Len() > 0{
            return dest.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.CanEvalWith(processInfo, other, canEvalType, alt2type)
        }
        return true, nil
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(dest.FP.Get_nilable())) &&
        Lns_GetEnv().SetStackVal( otherSrc.FP.Get_nilable()) ).(bool)){
        if canEvalType != Ast_CanEvalType__Equal{
            if dest.FP.Get_kind() == Ast_TypeInfoKind__Box{
                return dest.FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
            }
            return false, nil
        } else { 
            otherSrc = otherSrc.FP.Get_nonnilableType()
            
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( dest == Ast_builtinTypeStem) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(otherSrc.FP.Get_nilable())) ).(bool)){
        return true, nil
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( dest == Ast_builtinTypeForm) &&
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() == Ast_TypeInfoKind__Func) ||
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() == Ast_TypeInfoKind__Form) ||
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() == Ast_TypeInfoKind__FormFunc) ).(bool))) ).(bool)){
        return Ast_isSettableToForm_5182_(processInfo, otherSrc), nil
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( otherSrc == Ast_builtinTypeNil) ||
        Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() == Ast_TypeInfoKind__Abbr) ).(bool){
        if dest.FP.Get_kind() != Ast_TypeInfoKind__Nilable{
            return false, nil
        }
        return true, nil
    }
    if dest.FP.Get_typeId() == otherSrc.FP.Get_typeId(){
        return true, nil
    }
    if dest.FP.Get_kind() == Ast_TypeInfoKind__Ext{
        return dest.FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
    }
    {
        _extTypeInfo := Ast_ExtTypeInfoDownCastF(otherSrc.FP)
        if _extTypeInfo != nil {
            extTypeInfo := _extTypeInfo.(*Ast_ExtTypeInfo)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( canEvalType != Ast_CanEvalType__SetEq) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Ast_failCreateLuavalWith_4234_(extTypeInfo.FP.Get_extedType(), Ast_LuavalConvKind__ToLua, false)))) ).(bool)){
                otherSrc = extTypeInfo.FP.Get_extedType()
                
            }
        }
    }
    if dest.FP.Get_kind() != otherSrc.FP.Get_kind(){
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.HasBase()) ).(bool)){
            return Ast_TypeInfo_canEvalWithBase(processInfo, dest, destMut, otherSrc.FP.Get_baseTypeInfo(), canEvalType, alt2type)
        }
        if dest.FP.Get_kind() == Ast_TypeInfoKind__Nilable{
            var dstNonNil *Ast_TypeInfo
            if destMut{
                dstNonNil = dest.FP.Get_nonnilableType()
                
            } else { 
                dstNonNil = processInfo.FP.CreateModifier(dest.FP.Get_nonnilableType(), Ast_MutMode__IMut)
                
            }
            if otherSrc.FP.Get_nilable(){
                if otherSrc.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                    return dstNonNil.FP.CanEvalWith(processInfo, otherSrc.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), canEvalType, alt2type)
                }
                return dstNonNil.FP.CanEvalWith(processInfo, otherSrc.FP.Get_nonnilableType(), canEvalType, alt2type)
            }
            return dstNonNil.FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
        } else if Ast_isGenericType(dest){
            return dest.FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
        } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( dest.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( dest.FP.Get_kind() == Ast_TypeInfoKind__IF) ).(bool))) &&
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() == Ast_TypeInfoKind__IF) ).(bool))) ).(bool)){
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( canEvalType == Ast_CanEvalType__SetOp) ||
                Lns_GetEnv().SetStackVal( canEvalType == Ast_CanEvalType__SetOpIMut) ).(bool){
                return otherSrc.FP.IsInheritFrom(processInfo, dest, alt2type), nil
            }
            return false, nil
        } else if otherSrc.FP.Get_kind() == Ast_TypeInfoKind__Enum{
            var enumTypeInfo *Ast_EnumTypeInfo
            enumTypeInfo = Lns_unwrap( (Ast_EnumTypeInfoDownCastF(otherSrc.FP))).(*Ast_EnumTypeInfo)
            return dest.FP.CanEvalWith(processInfo, enumTypeInfo.FP.Get_valTypeInfo(), canEvalType, alt2type)
        } else if dest.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
            return dest.FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
        } else if dest.FP.Get_kind() == Ast_TypeInfoKind__Box{
            return dest.FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
        } else if dest.FP.Get_kind() == Ast_TypeInfoKind__Form{
            if _switch26597 := otherSrc.FP.Get_kind(); _switch26597 == Ast_TypeInfoKind__Form {
                return true, nil
            } else if _switch26597 == Ast_TypeInfoKind__FormFunc || _switch26597 == Ast_TypeInfoKind__Func {
                return Ast_isSettableToForm_5182_(processInfo, otherSrc), nil
            }
        } else if dest.FP.Get_kind() == Ast_TypeInfoKind__FormFunc{
            if _switch26802 := otherSrc.FP.Get_kind(); _switch26802 == Ast_TypeInfoKind__FormFunc || _switch26802 == Ast_TypeInfoKind__Func {
                {
                    var result LnsInt
                    var mess string
                    result,mess = Ast_TypeInfo_checkMatchType(processInfo, dest.FP.Get_argTypeInfoList(), otherSrc.FP.Get_argTypeInfoList(), false, nil, alt2type)
                    if result == Ast_MatchType__Error{
                        return false, mess
                    }
                }
                {
                    var result LnsInt
                    var mess string
                    result,mess = Ast_TypeInfo_checkMatchType(processInfo, otherSrc.FP.Get_argTypeInfoList(), dest.FP.Get_argTypeInfoList(), false, nil, alt2type)
                    if result == Ast_MatchType__Error{
                        return false, mess
                    }
                }
                {
                    var result LnsInt
                    var mess string
                    result,mess = Ast_TypeInfo_checkMatchType(processInfo, dest.FP.Get_retTypeInfoList(), otherSrc.FP.Get_retTypeInfoList(), false, nil, alt2type)
                    if result == Ast_MatchType__Error{
                        return false, mess
                    }
                }
                {
                    var result LnsInt
                    var mess string
                    result,mess = Ast_TypeInfo_checkMatchType(processInfo, otherSrc.FP.Get_retTypeInfoList(), dest.FP.Get_retTypeInfoList(), false, nil, alt2type)
                    if result == Ast_MatchType__Error{
                        return false, mess
                    }
                }
                if dest.FP.Get_retTypeInfoList().Len() != otherSrc.FP.Get_retTypeInfoList().Len(){
                    return false, Lns_getVM().String_format("argNum %d != %d", []LnsAny{dest.FP.Get_retTypeInfoList().Len(), otherSrc.FP.Get_retTypeInfoList().Len()})
                }
                return true, nil
            }
        }
        return false, nil
    }
    if _switch27773 := dest.FP.Get_kind(); _switch27773 == Ast_TypeInfoKind__Prim {
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( dest == Ast_builtinTypeInt) &&
            Lns_GetEnv().SetStackVal( otherSrc == Ast_builtinTypeChar) ||
            Lns_GetEnv().SetStackVal( dest == Ast_builtinTypeChar) &&
            Lns_GetEnv().SetStackVal( otherSrc == Ast_builtinTypeInt) ).(bool){
            return true, nil
        }
        return false, nil
    } else if _switch27773 == Ast_TypeInfoKind__List || _switch27773 == Ast_TypeInfoKind__Array || _switch27773 == Ast_TypeInfoKind__Set {
        if otherSrc.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone{
            return true, nil
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( dest.FP.Get_itemTypeInfoList().Len() >= 1) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_itemTypeInfoList().Len() >= 1) ).(bool)){
            var ret bool
            var mess LnsAny
            ret,mess = (dest.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()).FP.CanEvalWith(processInfo, otherSrc.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( destMut) &&
                Lns_GetEnv().SetStackVal( Ast_CanEvalType__SetEq) ||
                Lns_GetEnv().SetStackVal( Ast_CanEvalType__SetOpIMut) ).(LnsInt), alt2type)
            if Lns_op_not(ret){
                return false, mess
            }
        } else { 
            return false, nil
        }
        
        return true, nil
    } else if _switch27773 == Ast_TypeInfoKind__Map {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone) ).(bool)){
            return true, nil
        }
        var check1 func() LnsAny
        check1 = func() LnsAny {
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( dest.FP.Get_itemTypeInfoList().Len() >= 1) &&
                Lns_GetEnv().SetStackVal( otherSrc.FP.Get_itemTypeInfoList().Len() >= 1) ).(bool)){
                var ret bool
                ret,_ = (dest.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()).FP.CanEvalWith(processInfo, otherSrc.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( destMut) &&
                    Lns_GetEnv().SetStackVal( Ast_CanEvalType__SetEq) ||
                    Lns_GetEnv().SetStackVal( Ast_CanEvalType__SetOpIMut) ).(LnsInt), alt2type)
                if Lns_op_not(ret){
                    return false
                }
            } else { 
                return false
            }
            
            return true
        }
        var check2 func() bool
        check2 = func() bool {
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( dest.FP.Get_itemTypeInfoList().Len() >= 2) &&
                Lns_GetEnv().SetStackVal( otherSrc.FP.Get_itemTypeInfoList().Len() >= 2) ).(bool)){
                var ret bool
                ret,_ = (dest.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()).FP.CanEvalWith(processInfo, otherSrc.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( destMut) &&
                    Lns_GetEnv().SetStackVal( Ast_CanEvalType__SetEq) ||
                    Lns_GetEnv().SetStackVal( Ast_CanEvalType__SetOpIMut) ).(LnsInt), alt2type)
                if Lns_op_not(ret){
                    return false
                }
            } else { 
                return false
            }
            
            return true
        }
        var result1 LnsAny
        result1 = check1()
        var result2 bool
        result2 = check2()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( result1) &&
            Lns_GetEnv().SetStackVal( result2) )){
            return true, nil
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(result1)) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone) ||
            Lns_GetEnv().SetStackVal( Lns_op_not(result2)) &&
            Lns_GetEnv().SetStackVal( otherSrc.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone) ).(bool){
            return true, nil
        }
        return false, nil
    } else if _switch27773 == Ast_TypeInfoKind__Class || _switch27773 == Ast_TypeInfoKind__IF {
        if Ast_isGenericType(dest){
            return dest.FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( canEvalType == Ast_CanEvalType__SetOp) ||
            Lns_GetEnv().SetStackVal( canEvalType == Ast_CanEvalType__SetOpIMut) ).(bool){
            return otherSrc.FP.IsInheritFrom(processInfo, dest, alt2type), nil
        }
        return false, nil
    } else if _switch27773 == Ast_TypeInfoKind__Form {
        return Ast_isSettableToForm_5182_(processInfo, otherSrc), nil
    } else if _switch27773 == Ast_TypeInfoKind__Func || _switch27773 == Ast_TypeInfoKind__FormFunc {
        if dest.FP.Get_retTypeInfoList().Len() != otherSrc.FP.Get_retTypeInfoList().Len(){
            return false, Lns_getVM().String_format("argNum %d != %d", []LnsAny{dest.FP.Get_retTypeInfoList().Len(), otherSrc.FP.Get_retTypeInfoList().Len()})
        }
        var argCheck LnsInt
        var argMess string
        argCheck,argMess = Ast_TypeInfo_checkMatchType(processInfo, dest.FP.Get_argTypeInfoList(), otherSrc.FP.Get_argTypeInfoList(), false, nil, alt2type)
        if argCheck == Ast_MatchType__Error{
            return false, argMess
        }
        var retCheck LnsInt
        var retMess string
        retCheck,retMess = Ast_TypeInfo_checkMatchType(processInfo, dest.FP.Get_retTypeInfoList(), otherSrc.FP.Get_retTypeInfoList(), false, nil, alt2type)
        if retCheck == Ast_MatchType__Error{
            return false, retMess
        }
        return true, nil
    } else if _switch27773 == Ast_TypeInfoKind__Method {
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( dest.FP.Get_argTypeInfoList().Len() != otherSrc.FP.Get_argTypeInfoList().Len()) ||
            Lns_GetEnv().SetStackVal( dest.FP.Get_retTypeInfoList().Len() != otherSrc.FP.Get_retTypeInfoList().Len()) ).(bool){
            return false, nil
        }
        for _index, _argType := range( dest.FP.Get_argTypeInfoList().Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var otherArgType *Ast_TypeInfo
            otherArgType = otherSrc.FP.Get_argTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(argType.FP.Equals(processInfo, otherArgType, alt2type, nil)){
                var mess string
                mess = Lns_getVM().String_format("unmatch arg(%d) type -- %s, %s", []LnsAny{index, argType.FP.GetTxt(nil, nil, nil), otherArgType.FP.GetTxt(nil, nil, nil)})
                return false, mess
            }
        }
        for _index, _retType := range( dest.FP.Get_retTypeInfoList().Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var otherRetType *Ast_TypeInfo
            otherRetType = otherSrc.FP.Get_retTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(retType.FP.Equals(processInfo, otherRetType, alt2type, nil)){
                var mess string
                mess = Lns_getVM().String_format("unmatch ret(%d) type -- %s, %s, %s", []LnsAny{index, retType.FP.GetTxt(nil, nil, nil), otherRetType.FP.GetTxt(nil, nil, nil), dest.FP.GetTxt(nil, nil, nil)})
                return false, mess
            }
        }
        if dest.FP.Get_mutMode() != otherSrc.FP.Get_mutMode(){
            var mess string
            mess = Lns_getVM().String_format("unmatch mutable mode -- %s, %s", []LnsAny{Ast_MutMode_getTxt( dest.FP.Get_mutMode()), Ast_MutMode_getTxt( otherSrc.FP.Get_mutMode())})
            return false, mess
        }
        return true, nil
    } else if _switch27773 == Ast_TypeInfoKind__Nilable {
        var dstNonNil *Ast_TypeInfo
        if destMut{
            dstNonNil = dest.FP.Get_nonnilableType()
            
        } else { 
            dstNonNil = processInfo.FP.CreateModifier(dest.FP.Get_nonnilableType(), Ast_MutMode__IMut)
            
        }
        return dstNonNil.FP.CanEvalWith(processInfo, otherSrc.FP.Get_nonnilableType(), canEvalType, alt2type)
    } else if _switch27773 == Ast_TypeInfoKind__Alternate {
        return dest.FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
    } else if _switch27773 == Ast_TypeInfoKind__Box {
        return dest.FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
    } else {
        return false, nil
    }
// insert a dummy
    return false,nil
}

// 6667: decl @lune.@base.@Ast.TypeInfo.getFullName
func (self *Ast_TypeInfo) GetFullName(typeNameCtrl *Ast_TypeNameCtrl,importInfo Ast_ModuleInfoManager,localFlag LnsAny) string {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( localFlag) &&
        Lns_GetEnv().SetStackVal( self.FP.IsModule()) )){
        return typeNameCtrl.FP.GetModuleName(self, "", importInfo)
    }
    return self.FP.GetParentFullName(typeNameCtrl, importInfo, localFlag) + self.FP.Get_rawTxt()
}


// declaration Class -- RootTypeInfo
type Ast_RootTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_RootTypeInfo struct {
    Ast_TypeInfo
    typeId LnsInt
    FP Ast_RootTypeInfoMtd
}
func Ast_RootTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_RootTypeInfo).FP
}
type Ast_RootTypeInfoDownCast interface {
    ToAst_RootTypeInfo() *Ast_RootTypeInfo
}
func Ast_RootTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_RootTypeInfoDownCast)
    if ok { return work.ToAst_RootTypeInfo() }
    return nil
}
func (obj *Ast_RootTypeInfo) ToAst_RootTypeInfo() *Ast_RootTypeInfo {
    return obj
}
func NewAst_RootTypeInfo() *Ast_RootTypeInfo {
    obj := &Ast_RootTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_RootTypeInfo()
    return obj
}
func (self *Ast_RootTypeInfo) Get_typeId() LnsInt{ return self.typeId }
// 1094: DeclConstr
func (self *Ast_RootTypeInfo) InitAst_RootTypeInfo() {
    self.InitAst_TypeInfo(Ast_rootScope, Ast_rootProcessInfo)
    self.typeId = Ast_rootProcessInfo.FP.Get_idProv().FP.Get_id()
    
}

// 1100: decl @lune.@base.@Ast.RootTypeInfo.get_baseTypeInfo
func (self *Ast_RootTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1104: decl @lune.@base.@Ast.RootTypeInfo.get_parentInfo
func (self *Ast_RootTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1108: decl @lune.@base.@Ast.RootTypeInfo.create
func Ast_RootTypeInfo_create_1961_() *Ast_RootTypeInfo {
    return NewAst_RootTypeInfo()
}


// declaration Class -- NormalSymbolInfo
type Ast_NormalSymbolInfoMtd interface {
    CanAccess(arg1 *Ast_Scope, arg2 LnsInt) LnsAny
    ClearValue()
    GetModule() *Ast_TypeInfo
    GetOrg() *Ast_SymbolInfo
    Get_accessMode() LnsInt
    Get_canBeLeft() bool
    Get_canBeRight() bool
    Get_convModuleParam() LnsAny
    Get_hasAccessFromClosure() bool
    Get_hasValueFlag() bool
    Get_isLazyLoad() bool
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_mutable() bool
    Get_name() string
    Get_namespaceTypeInfo() *Ast_TypeInfo
    Get_pos() LnsAny
    Get_posForLatestMod() LnsAny
    Get_posForModToRef() LnsAny
    Get_scope() *Ast_Scope
    Get_staticFlag() bool
    Get_symbolId() LnsInt
    Get_typeInfo() *Ast_TypeInfo
    HasAccess() bool
    Set_convModuleParam(arg1 LnsAny)
    Set_hasAccessFromClosure(arg1 bool)
    Set_hasValueFlag(arg1 bool)
    Set_posForLatestMod(arg1 LnsAny)
    Set_posForModToRef(arg1 LnsAny)
    Set_typeInfo(arg1 *Ast_TypeInfo)
    UpdateValue(arg1 LnsAny)
}
type Ast_NormalSymbolInfo struct {
    Ast_SymbolInfo
    canBeLeft bool
    canBeRight bool
    symbolId LnsInt
    scope *Ast_Scope
    accessMode LnsInt
    staticFlag bool
    isLazyLoad bool
    name string
    pos LnsAny
    typeInfo *Ast_TypeInfo
    kind LnsInt
    hasValueFlag bool
    mutMode LnsInt
    hasAccessFromClosure bool
    posForLatestMod LnsAny
    posForModToRef LnsAny
    convModuleParam LnsAny
    FP Ast_NormalSymbolInfoMtd
}
func Ast_NormalSymbolInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_NormalSymbolInfo).FP
}
type Ast_NormalSymbolInfoDownCast interface {
    ToAst_NormalSymbolInfo() *Ast_NormalSymbolInfo
}
func Ast_NormalSymbolInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_NormalSymbolInfoDownCast)
    if ok { return work.ToAst_NormalSymbolInfo() }
    return nil
}
func (obj *Ast_NormalSymbolInfo) ToAst_NormalSymbolInfo() *Ast_NormalSymbolInfo {
    return obj
}
func NewAst_NormalSymbolInfo(arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 bool, arg4 bool, arg5 *Ast_Scope, arg6 LnsInt, arg7 bool, arg8 string, arg9 LnsAny, arg10 *Ast_TypeInfo, arg11 LnsAny, arg12 bool, arg13 bool) *Ast_NormalSymbolInfo {
    obj := &Ast_NormalSymbolInfo{}
    obj.FP = obj
    obj.Ast_SymbolInfo.FP = obj
    obj.InitAst_NormalSymbolInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
    return obj
}
func (self *Ast_NormalSymbolInfo) Get_canBeLeft() bool{ return self.canBeLeft }
func (self *Ast_NormalSymbolInfo) Get_canBeRight() bool{ return self.canBeRight }
func (self *Ast_NormalSymbolInfo) Get_symbolId() LnsInt{ return self.symbolId }
func (self *Ast_NormalSymbolInfo) Get_scope() *Ast_Scope{ return self.scope }
func (self *Ast_NormalSymbolInfo) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Ast_NormalSymbolInfo) Get_staticFlag() bool{ return self.staticFlag }
func (self *Ast_NormalSymbolInfo) Get_isLazyLoad() bool{ return self.isLazyLoad }
func (self *Ast_NormalSymbolInfo) Get_name() string{ return self.name }
func (self *Ast_NormalSymbolInfo) Get_pos() LnsAny{ return self.pos }
func (self *Ast_NormalSymbolInfo) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }
func (self *Ast_NormalSymbolInfo) Get_kind() LnsInt{ return self.kind }
func (self *Ast_NormalSymbolInfo) Get_hasValueFlag() bool{ return self.hasValueFlag }
func (self *Ast_NormalSymbolInfo) Set_hasValueFlag(arg1 bool){ self.hasValueFlag = arg1 }
func (self *Ast_NormalSymbolInfo) Get_mutMode() LnsInt{ return self.mutMode }
func (self *Ast_NormalSymbolInfo) Get_hasAccessFromClosure() bool{ return self.hasAccessFromClosure }
func (self *Ast_NormalSymbolInfo) Set_hasAccessFromClosure(arg1 bool){ self.hasAccessFromClosure = arg1 }
func (self *Ast_NormalSymbolInfo) Get_posForLatestMod() LnsAny{ return self.posForLatestMod }
func (self *Ast_NormalSymbolInfo) Set_posForLatestMod(arg1 LnsAny){ self.posForLatestMod = arg1 }
func (self *Ast_NormalSymbolInfo) Get_posForModToRef() LnsAny{ return self.posForModToRef }
func (self *Ast_NormalSymbolInfo) Set_posForModToRef(arg1 LnsAny){ self.posForModToRef = arg1 }
func (self *Ast_NormalSymbolInfo) Get_convModuleParam() LnsAny{ return self.convModuleParam }
func (self *Ast_NormalSymbolInfo) Set_convModuleParam(arg1 LnsAny){ self.convModuleParam = arg1 }
// 1187: decl @lune.@base.@Ast.NormalSymbolInfo.get_mutable
func (self *Ast_NormalSymbolInfo) Get_mutable() bool {
    return Ast_isMutable(self.mutMode)
}

// 1195: decl @lune.@base.@Ast.NormalSymbolInfo.getOrg
func (self *Ast_NormalSymbolInfo) GetOrg() *Ast_SymbolInfo {
    return &self.Ast_SymbolInfo
}

// 1199: DeclConstr
func (self *Ast_NormalSymbolInfo) InitAst_NormalSymbolInfo(processInfo *Ast_ProcessInfo,kind LnsInt,canBeLeft bool,canBeRight bool,scope *Ast_Scope,accessMode LnsInt,staticFlag bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutMode LnsAny,hasValueFlag bool,isLazyLoad bool) {
    self.InitAst_SymbolInfo()
    self.convModuleParam = nil
    
    self.hasAccessFromClosure = false
    
    if hasValueFlag{
        self.posForLatestMod = pos
        
    } else { 
        self.posForLatestMod = nil
        
    }
    self.posForModToRef = nil
    
    self.kind = kind
    
    self.canBeLeft = canBeLeft
    
    self.canBeRight = canBeRight
    
    self.symbolId = processInfo.FP.Get_idProvSym().FP.GetNewId()
    
    self.scope = scope
    
    self.accessMode = accessMode
    
    self.staticFlag = staticFlag
    
    self.name = name
    
    self.pos = pos
    
    self.typeInfo = typeInfo
    
    self.mutMode = Lns_unwrapDefault( mutMode, Ast_MutMode__IMut).(LnsInt)
    
    self.hasValueFlag = hasValueFlag
    
    self.isLazyLoad = isLazyLoad
    
}


// 2343: decl @lune.@base.@Ast.NormalSymbolInfo.canAccess
func (self *Ast_NormalSymbolInfo) CanAccess(fromScope *Ast_Scope,access LnsInt) LnsAny {
    if access == Ast_ScopeAccess__Full{
        return &self.Ast_SymbolInfo
    }
    if self.scope == fromScope{
        return &self.Ast_SymbolInfo
    }
    var processInfo *Ast_ProcessInfo
    processInfo = fromScope.FP.GetProcessInfo()
    if _switch7902 := self.FP.Get_accessMode(); _switch7902 == Ast_AccessMode__Pub || _switch7902 == Ast_AccessMode__Global {
        return &self.Ast_SymbolInfo
    } else if _switch7902 == Ast_AccessMode__Pro {
        var nsClass *Ast_TypeInfo
        nsClass = self.scope.FP.GetClassTypeInfo()
        var fromClass *Ast_TypeInfo
        fromClass = fromScope.FP.GetClassTypeInfo()
        if fromClass.FP.IsInheritFrom(processInfo, nsClass, nil){
            return &self.Ast_SymbolInfo
        }
        return nil
    } else if _switch7902 == Ast_AccessMode__Local {
        return &self.Ast_SymbolInfo
    } else if _switch7902 == Ast_AccessMode__Pri {
        if fromScope.FP.IsInnerOf(self.scope){
            return &self.Ast_SymbolInfo
        }
        return nil
    }
    Util_err(Lns_getVM().String_format("illegl accessmode -- %s, %s", []LnsAny{self.FP.Get_accessMode(), self.FP.Get_name()}))
// insert a dummy
    return nil
}

// 4320: decl @lune.@base.@Ast.NormalSymbolInfo.set_typeInfo
func (self *Ast_NormalSymbolInfo) Set_typeInfo(typeInfo *Ast_TypeInfo) {
    if self.name == "_"{
        return 
    }
    self.typeInfo = typeInfo
    
}


// declaration Class -- AutoBoxingInfo
var Ast_AutoBoxingInfo__allObj *LnsMap
// 1266: decl @lune.@base.@Ast.AutoBoxingInfo.___init
func Ast_AutoBoxingInfo____init_2085_() {
    Ast_AutoBoxingInfo__allObj = NewLnsMap( map[LnsAny]LnsAny{})
    
}

type Ast_AutoBoxingInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_count() LnsInt
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    Inc()
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
    Unregist()
}
type Ast_AutoBoxingInfo struct {
    Ast_TypeInfo
    allObj *LnsMap
    count LnsInt
    FP Ast_AutoBoxingInfoMtd
}
func Ast_AutoBoxingInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_AutoBoxingInfo).FP
}
type Ast_AutoBoxingInfoDownCast interface {
    ToAst_AutoBoxingInfo() *Ast_AutoBoxingInfo
}
func Ast_AutoBoxingInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_AutoBoxingInfoDownCast)
    if ok { return work.ToAst_AutoBoxingInfo() }
    return nil
}
func (obj *Ast_AutoBoxingInfo) ToAst_AutoBoxingInfo() *Ast_AutoBoxingInfo {
    return obj
}
func NewAst_AutoBoxingInfo(arg1 *Ast_ProcessInfo) *Ast_AutoBoxingInfo {
    obj := &Ast_AutoBoxingInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AutoBoxingInfo(arg1)
    return obj
}
func (self *Ast_AutoBoxingInfo) Get_count() LnsInt{ return self.count }
// 1271: DeclConstr
func (self *Ast_AutoBoxingInfo) InitAst_AutoBoxingInfo(processInfo *Ast_ProcessInfo) {
    self.InitAst_TypeInfo(nil, processInfo)
    self.count = 0
    
    Ast_AutoBoxingInfo__allObj.Set(&self.Ast_TypeInfo,self)
}

// 1283: decl @lune.@base.@Ast.AutoBoxingInfo.get_baseTypeInfo
func (self *Ast_AutoBoxingInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 1287: decl @lune.@base.@Ast.AutoBoxingInfo.get_parentInfo
func (self *Ast_AutoBoxingInfo) Get_parentInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 1291: decl @lune.@base.@Ast.AutoBoxingInfo.get_kind
func (self *Ast_AutoBoxingInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Etc
}

// 1295: decl @lune.@base.@Ast.AutoBoxingInfo.inc
func (self *Ast_AutoBoxingInfo) Inc() {
    var obj *Ast_AutoBoxingInfo
    obj = Lns_unwrap( Ast_AutoBoxingInfo__allObj.Items[&self.Ast_TypeInfo]).(*Ast_AutoBoxingInfo)
    obj.count = obj.count + 1
    
}

// 1300: decl @lune.@base.@Ast.AutoBoxingInfo.unregist
func (self *Ast_AutoBoxingInfo) Unregist() {
    Ast_AutoBoxingInfo__allObj.Set(&self.Ast_TypeInfo,nil)
}


// declaration Class -- CanEvalCtrlTypeInfo
var Ast_CanEvalCtrlTypeInfo__detectAlt *Ast_CanEvalCtrlTypeInfo
var Ast_CanEvalCtrlTypeInfo__needAutoBoxing *Ast_CanEvalCtrlTypeInfo
var Ast_CanEvalCtrlTypeInfo__checkTypeTarget *Ast_CanEvalCtrlTypeInfo
// 1305: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.___init
func Ast_CanEvalCtrlTypeInfo____init_2121_() {
    Ast_CanEvalCtrlTypeInfo__detectAlt = NewAst_CanEvalCtrlTypeInfo()
    
    Ast_CanEvalCtrlTypeInfo__needAutoBoxing = NewAst_CanEvalCtrlTypeInfo()
    
    Ast_CanEvalCtrlTypeInfo__checkTypeTarget = NewAst_CanEvalCtrlTypeInfo()
    
}

type Ast_CanEvalCtrlTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_CanEvalCtrlTypeInfo struct {
    Ast_TypeInfo
    DetectAlt *Ast_CanEvalCtrlTypeInfo
    NeedAutoBoxing *Ast_CanEvalCtrlTypeInfo
    CheckTypeTarget *Ast_CanEvalCtrlTypeInfo
    FP Ast_CanEvalCtrlTypeInfoMtd
}
func Ast_CanEvalCtrlTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_CanEvalCtrlTypeInfo).FP
}
type Ast_CanEvalCtrlTypeInfoDownCast interface {
    ToAst_CanEvalCtrlTypeInfo() *Ast_CanEvalCtrlTypeInfo
}
func Ast_CanEvalCtrlTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_CanEvalCtrlTypeInfoDownCast)
    if ok { return work.ToAst_CanEvalCtrlTypeInfo() }
    return nil
}
func (obj *Ast_CanEvalCtrlTypeInfo) ToAst_CanEvalCtrlTypeInfo() *Ast_CanEvalCtrlTypeInfo {
    return obj
}
func NewAst_CanEvalCtrlTypeInfo() *Ast_CanEvalCtrlTypeInfo {
    obj := &Ast_CanEvalCtrlTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_CanEvalCtrlTypeInfo()
    return obj
}
// 1310: DeclConstr
func (self *Ast_CanEvalCtrlTypeInfo) InitAst_CanEvalCtrlTypeInfo() {
    self.InitAst_TypeInfo(nil, Ast_rootProcessInfo)
}

// 1319: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.get_kind
func (self *Ast_CanEvalCtrlTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__CanEvalCtrl
}

// 1323: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.get_typeId
func (self *Ast_CanEvalCtrlTypeInfo) Get_typeId() LnsInt {
    return -1
}

// 1328: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.get_baseTypeInfo
func (self *Ast_CanEvalCtrlTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 1332: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.get_parentInfo
func (self *Ast_CanEvalCtrlTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 1336: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap
func Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(detectFlag bool) *LnsMap {
    if detectFlag{
        var _map *LnsMap
        _map = NewLnsMap( map[LnsAny]LnsAny{})
        _map.Set(&Ast_CanEvalCtrlTypeInfo__detectAlt.Ast_TypeInfo,Ast_headTypeInfo)
        return _map
    }
    return NewLnsMap( map[LnsAny]LnsAny{})
}

// 1345: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.isValidApply
func Ast_CanEvalCtrlTypeInfo_isValidApply(alt2type *LnsMap) bool {
    return Ast_TypeInfo2Stem(alt2type.Items[&Ast_CanEvalCtrlTypeInfo__detectAlt.Ast_TypeInfo]) != nil
}

// 1349: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.setupNeedAutoBoxing
func Ast_CanEvalCtrlTypeInfo_setupNeedAutoBoxing(alt2type *LnsMap,processInfo *Ast_ProcessInfo) {
    alt2type.Set(&Ast_CanEvalCtrlTypeInfo__needAutoBoxing.Ast_TypeInfo,&NewAst_AutoBoxingInfo(processInfo).Ast_TypeInfo)
}

// 1355: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.updateNeedAutoBoxing
func Ast_CanEvalCtrlTypeInfo_updateNeedAutoBoxing(alt2type *LnsMap) {
    {
        __exp := alt2type.Items[&Ast_CanEvalCtrlTypeInfo__needAutoBoxing.Ast_TypeInfo]
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            {
                _autoBoxingInfo := Ast_AutoBoxingInfoDownCastF(_exp.FP)
                if _autoBoxingInfo != nil {
                    autoBoxingInfo := _autoBoxingInfo.(*Ast_AutoBoxingInfo)
                    autoBoxingInfo.FP.Inc()
                }
            }
        } else {
            Util_err("no exist needAutoBoxing")
        }
    }
}

// 1366: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.hasNeedAutoBoxing
func Ast_CanEvalCtrlTypeInfo_hasNeedAutoBoxing(alt2type *LnsMap) bool {
    {
        __exp := alt2type.Items[&Ast_CanEvalCtrlTypeInfo__needAutoBoxing.Ast_TypeInfo]
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            {
                _autoBoxingInfo := Ast_AutoBoxingInfoDownCastF(_exp.FP)
                if _autoBoxingInfo != nil {
                    autoBoxingInfo := _autoBoxingInfo.(*Ast_AutoBoxingInfo)
                    return autoBoxingInfo.FP.Get_count() != 0
                }
            }
        }
    }
    return false
}

// 1376: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.finishNeedAutoBoxing
func Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(alt2type *LnsMap,count LnsInt) bool {
    {
        __exp := alt2type.Items[&Ast_CanEvalCtrlTypeInfo__needAutoBoxing.Ast_TypeInfo]
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            {
                _autoBoxingInfo := Ast_AutoBoxingInfoDownCastF(_exp.FP)
                if _autoBoxingInfo != nil {
                    autoBoxingInfo := _autoBoxingInfo.(*Ast_AutoBoxingInfo)
                    autoBoxingInfo.FP.Unregist()
                    return autoBoxingInfo.FP.Get_count() == count
                }
            }
        }
    }
    return false
}

// 1393: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.canAutoBoxing
func Ast_CanEvalCtrlTypeInfo_canAutoBoxing(dst *Ast_TypeInfo,src *Ast_TypeInfo) bool {
    var dstSrc *Ast_TypeInfo
    dstSrc = dst.FP.Get_srcTypeInfo().FP.Get_nonnilableType()
    if dstSrc.FP.Get_kind() != Ast_TypeInfoKind__Box{
        return false
    }
    var srcSrc *Ast_TypeInfo
    srcSrc = src.FP.Get_srcTypeInfo().FP.Get_nonnilableType()
    if srcSrc.FP.Get_kind() == Ast_TypeInfoKind__Box{
        return false
    }
    return true
}


// declaration Class -- NilableTypeInfo
type Ast_NilableTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_NilableTypeInfo struct {
    Ast_TypeInfo
    nonnilableType *Ast_TypeInfo
    typeId LnsInt
    FP Ast_NilableTypeInfoMtd
}
func Ast_NilableTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_NilableTypeInfo).FP
}
type Ast_NilableTypeInfoDownCast interface {
    ToAst_NilableTypeInfo() *Ast_NilableTypeInfo
}
func Ast_NilableTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_NilableTypeInfoDownCast)
    if ok { return work.ToAst_NilableTypeInfo() }
    return nil
}
func (obj *Ast_NilableTypeInfo) ToAst_NilableTypeInfo() *Ast_NilableTypeInfo {
    return obj
}
func NewAst_NilableTypeInfo(arg1 LnsAny, arg2 *Ast_ProcessInfo, arg3 *Ast_TypeInfo, arg4 LnsInt) *Ast_NilableTypeInfo {
    obj := &Ast_NilableTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_NilableTypeInfo(arg1, arg2, arg3, arg4)
    return obj
}
func (self *Ast_NilableTypeInfo) InitAst_NilableTypeInfo(arg1 LnsAny, arg2 *Ast_ProcessInfo, arg3 *Ast_TypeInfo, arg4 LnsInt) {
    self.Ast_TypeInfo.InitAst_TypeInfo( arg1,arg2)
    self.nonnilableType = arg3
    self.typeId = arg4
}
func (self *Ast_NilableTypeInfo) Get_nonnilableType() *Ast_TypeInfo{ return self.nonnilableType }
func (self *Ast_NilableTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_NilableTypeInfo) AddChildren(arg1 *Ast_TypeInfo) {
self.nonnilableType. FP.AddChildren( arg1)
}
func (self *Ast_NilableTypeInfo) CreateAlt2typeMap(arg1 bool) *LnsMap {
    return self.nonnilableType. FP.CreateAlt2typeMap( arg1)
}
func (self *Ast_NilableTypeInfo) GetFullName(arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.nonnilableType. FP.GetFullName( arg1,arg2,arg3)
}
func (self *Ast_NilableTypeInfo) GetModule() *Ast_TypeInfo {
    return self.nonnilableType. FP.GetModule( )
}
func (self *Ast_NilableTypeInfo) GetOverridingType() LnsAny {
    return self.nonnilableType. FP.GetOverridingType( )
}
func (self *Ast_NilableTypeInfo) GetParentFullName(arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.nonnilableType. FP.GetParentFullName( arg1,arg2,arg3)
}
func (self *Ast_NilableTypeInfo) GetParentId() LnsInt {
    return self.nonnilableType. FP.GetParentId( )
}
func (self *Ast_NilableTypeInfo) getProcessInfo() *Ast_ProcessInfo {
    return self.nonnilableType. FP.getProcessInfo( )
}
func (self *Ast_NilableTypeInfo) Get_abstractFlag() bool {
    return self.nonnilableType. FP.Get_abstractFlag( )
}
func (self *Ast_NilableTypeInfo) Get_accessMode() LnsInt {
    return self.nonnilableType. FP.Get_accessMode( )
}
func (self *Ast_NilableTypeInfo) Get_argTypeInfoList() *LnsList {
    return self.nonnilableType. FP.Get_argTypeInfoList( )
}
func (self *Ast_NilableTypeInfo) Get_autoFlag() bool {
    return self.nonnilableType. FP.Get_autoFlag( )
}
func (self *Ast_NilableTypeInfo) Get_baseId() LnsInt {
    return self.nonnilableType. FP.Get_baseId( )
}
func (self *Ast_NilableTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_baseTypeInfo( )
}
func (self *Ast_NilableTypeInfo) Get_children() *LnsList {
    return self.nonnilableType. FP.Get_children( )
}
func (self *Ast_NilableTypeInfo) Get_extedType() *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_extedType( )
}
func (self *Ast_NilableTypeInfo) Get_externalFlag() bool {
    return self.nonnilableType. FP.Get_externalFlag( )
}
func (self *Ast_NilableTypeInfo) Get_genSrcTypeInfo() *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_genSrcTypeInfo( )
}
func (self *Ast_NilableTypeInfo) Get_interfaceList() *LnsList {
    return self.nonnilableType. FP.Get_interfaceList( )
}
func (self *Ast_NilableTypeInfo) Get_itemTypeInfoList() *LnsList {
    return self.nonnilableType. FP.Get_itemTypeInfoList( )
}
func (self *Ast_NilableTypeInfo) Get_mutMode() LnsInt {
    return self.nonnilableType. FP.Get_mutMode( )
}
func (self *Ast_NilableTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_nilableTypeInfo( )
}
func (self *Ast_NilableTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_parentInfo( )
}
func (self *Ast_NilableTypeInfo) Get_processInfo() *Ast_ProcessInfo {
    return self.nonnilableType. FP.Get_processInfo( )
}
func (self *Ast_NilableTypeInfo) Get_rawTxt() string {
    return self.nonnilableType. FP.Get_rawTxt( )
}
func (self *Ast_NilableTypeInfo) Get_retTypeInfoList() *LnsList {
    return self.nonnilableType. FP.Get_retTypeInfoList( )
}
func (self *Ast_NilableTypeInfo) Get_scope() LnsAny {
    return self.nonnilableType. FP.Get_scope( )
}
func (self *Ast_NilableTypeInfo) Get_staticFlag() bool {
    return self.nonnilableType. FP.Get_staticFlag( )
}
func (self *Ast_NilableTypeInfo) Get_typeData() *Ast_TypeData {
    return self.nonnilableType. FP.Get_typeData( )
}
func (self *Ast_NilableTypeInfo) HasBase() bool {
    return self.nonnilableType. FP.HasBase( )
}
func (self *Ast_NilableTypeInfo) HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool {
    return self.nonnilableType. FP.HasRouteNamespaceFrom( arg1)
}
func (self *Ast_NilableTypeInfo) IsInheritFrom(arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.nonnilableType. FP.IsInheritFrom( arg1,arg2,arg3)
}
func (self *Ast_NilableTypeInfo) IsModule() bool {
    return self.nonnilableType. FP.IsModule( )
}
func (self *Ast_NilableTypeInfo) SerializeTypeInfoList(arg1 string,arg2 *LnsList,arg3 LnsAny) string {
    return self.nonnilableType. FP.SerializeTypeInfoList( arg1,arg2,arg3)
}
func (self *Ast_NilableTypeInfo) SwitchScope(arg1 *Ast_Scope) {
self.nonnilableType. FP.SwitchScope( arg1)
}
// 1411: decl @lune.@base.@Ast.NilableTypeInfo.get_kind
func (self *Ast_NilableTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Nilable
}

// 1414: decl @lune.@base.@Ast.NilableTypeInfo.get_aliasSrc
func (self *Ast_NilableTypeInfo) Get_aliasSrc() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1417: decl @lune.@base.@Ast.NilableTypeInfo.get_srcTypeInfo
func (self *Ast_NilableTypeInfo) Get_srcTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1420: decl @lune.@base.@Ast.NilableTypeInfo.get_nilable
func (self *Ast_NilableTypeInfo) Get_nilable() bool {
    return true
}

// 1425: decl @lune.@base.@Ast.NilableTypeInfo.getTxt
func (self *Ast_NilableTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 1430: decl @lune.@base.@Ast.NilableTypeInfo.getTxtWithRaw
func (self *Ast_NilableTypeInfo) GetTxtWithRaw(raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.nonnilableType.FP.GetTxtWithRaw(raw, typeNameCtrl, importInfo, localFlag) + "!"
}

// 1439: decl @lune.@base.@Ast.NilableTypeInfo.get_display_stirng_with
func (self *Ast_NilableTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    if self.nonnilableType.FP.Get_kind() == Ast_TypeInfoKind__FormFunc{
        return self.nonnilableType.FP.Get_display_stirng_with(raw + "!", alt2type)
    }
    return self.nonnilableType.FP.Get_display_stirng_with(raw, alt2type) + "!"
}

// 1447: decl @lune.@base.@Ast.NilableTypeInfo.get_display_stirng
func (self *Ast_NilableTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
}

// 1451: decl @lune.@base.@Ast.NilableTypeInfo.serialize
func (self *Ast_NilableTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    stream.Write(Lns_getVM().String_format("{ skind = %d, typeId = %d, orgTypeId = %d }\n", []LnsAny{Ast_SerializeKind__Nilable, self.typeId, self.nonnilableType.FP.Get_typeId()}))
}

// 1457: decl @lune.@base.@Ast.NilableTypeInfo.equals
func (self *Ast_NilableTypeInfo) Equals(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if Lns_op_not(typeInfo.FP.Get_nilable()){
        return false
    }
    return self.nonnilableType.FP.Equals(processInfo, typeInfo.FP.Get_nonnilableType(), alt2type, checkModifer)
}

// 1468: decl @lune.@base.@Ast.NilableTypeInfo.applyGeneric
func (self *Ast_NilableTypeInfo) ApplyGeneric(alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.nonnilableType.FP.ApplyGeneric(alt2typeMap, moduleTypeInfo)
    if typeInfo == self.nonnilableType{
        return &self.Ast_TypeInfo
    }
    if typeInfo != nil{
        typeInfo_1502 := typeInfo.(*Ast_TypeInfo)
        return typeInfo_1502.FP.Get_nilableTypeInfo()
    }
    return nil
}

// 5795: decl @lune.@base.@Ast.NilableTypeInfo.canEvalWith
func (self *Ast_NilableTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    var otherSrc *Ast_TypeInfo
    otherSrc = other
    if &self.Ast_TypeInfo == Ast_builtinTypeStem_{
        return true, nil
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( otherSrc == Ast_builtinTypeNil) ||
        Lns_GetEnv().SetStackVal( otherSrc.FP.Get_kind() == Ast_TypeInfoKind__Abbr) ).(bool){
        if self.FP.Get_nonnilableType().FP.Get_kind() == Ast_TypeInfoKind__Box{
            return self.FP.Get_nonnilableType().FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
        }
        return true, nil
    }
    if self.typeId == otherSrc.FP.Get_typeId(){
        return true, nil
    }
    if otherSrc.FP.Get_nilable(){
        if otherSrc.FP.Get_kind() == Ast_TypeInfoKind__DDD{
            return self.FP.Get_nonnilableType().FP.CanEvalWith(processInfo, otherSrc.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), canEvalType, alt2type)
        }
        return self.FP.Get_nonnilableType().FP.CanEvalWith(processInfo, otherSrc.FP.Get_nonnilableType(), canEvalType, alt2type)
    }
    return self.FP.Get_nonnilableType().FP.CanEvalWith(processInfo, otherSrc, canEvalType, alt2type)
}


// declaration Class -- AliasTypeInfo
type Ast_AliasTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_aliasSrcTypeInfo() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_AliasTypeInfo struct {
    Ast_TypeInfo
    rawTxt string
    accessMode LnsInt
    parentInfo *Ast_TypeInfo
    aliasSrcTypeInfo *Ast_TypeInfo
    externalFlag bool
    typeId LnsInt
    nilableTypeInfo *Ast_TypeInfo
    FP Ast_AliasTypeInfoMtd
}
func Ast_AliasTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_AliasTypeInfo).FP
}
type Ast_AliasTypeInfoDownCast interface {
    ToAst_AliasTypeInfo() *Ast_AliasTypeInfo
}
func Ast_AliasTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_AliasTypeInfoDownCast)
    if ok { return work.ToAst_AliasTypeInfo() }
    return nil
}
func (obj *Ast_AliasTypeInfo) ToAst_AliasTypeInfo() *Ast_AliasTypeInfo {
    return obj
}
func NewAst_AliasTypeInfo(arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsInt, arg4 *Ast_TypeInfo, arg5 *Ast_TypeInfo, arg6 bool) *Ast_AliasTypeInfo {
    obj := &Ast_AliasTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AliasTypeInfo(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Ast_AliasTypeInfo) Get_rawTxt() string{ return self.rawTxt }
func (self *Ast_AliasTypeInfo) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Ast_AliasTypeInfo) Get_parentInfo() *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_AliasTypeInfo) Get_aliasSrcTypeInfo() *Ast_TypeInfo{ return self.aliasSrcTypeInfo }
func (self *Ast_AliasTypeInfo) Get_externalFlag() bool{ return self.externalFlag }
func (self *Ast_AliasTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_AliasTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo{ return self.nilableTypeInfo }
func (self *Ast_AliasTypeInfo) AddChildren(arg1 *Ast_TypeInfo) {
self.aliasSrcTypeInfo. FP.AddChildren( arg1)
}
func (self *Ast_AliasTypeInfo) CreateAlt2typeMap(arg1 bool) *LnsMap {
    return self.aliasSrcTypeInfo. FP.CreateAlt2typeMap( arg1)
}
func (self *Ast_AliasTypeInfo) GetOverridingType() LnsAny {
    return self.aliasSrcTypeInfo. FP.GetOverridingType( )
}
func (self *Ast_AliasTypeInfo) getProcessInfo() *Ast_ProcessInfo {
    return self.aliasSrcTypeInfo. FP.getProcessInfo( )
}
func (self *Ast_AliasTypeInfo) GetTxtWithRaw(arg1 string,arg2 LnsAny,arg3 LnsAny,arg4 LnsAny) string {
    return self.aliasSrcTypeInfo. FP.GetTxtWithRaw( arg1,arg2,arg3,arg4)
}
func (self *Ast_AliasTypeInfo) Get_abstractFlag() bool {
    return self.aliasSrcTypeInfo. FP.Get_abstractFlag( )
}
func (self *Ast_AliasTypeInfo) Get_argTypeInfoList() *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_argTypeInfoList( )
}
func (self *Ast_AliasTypeInfo) Get_autoFlag() bool {
    return self.aliasSrcTypeInfo. FP.Get_autoFlag( )
}
func (self *Ast_AliasTypeInfo) Get_baseId() LnsInt {
    return self.aliasSrcTypeInfo. FP.Get_baseId( )
}
func (self *Ast_AliasTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return self.aliasSrcTypeInfo. FP.Get_baseTypeInfo( )
}
func (self *Ast_AliasTypeInfo) Get_children() *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_children( )
}
func (self *Ast_AliasTypeInfo) Get_display_stirng_with(arg1 string,arg2 LnsAny) string {
    return self.aliasSrcTypeInfo. FP.Get_display_stirng_with( arg1,arg2)
}
func (self *Ast_AliasTypeInfo) Get_extedType() *Ast_TypeInfo {
    return self.aliasSrcTypeInfo. FP.Get_extedType( )
}
func (self *Ast_AliasTypeInfo) Get_interfaceList() *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_interfaceList( )
}
func (self *Ast_AliasTypeInfo) Get_itemTypeInfoList() *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_itemTypeInfoList( )
}
func (self *Ast_AliasTypeInfo) Get_kind() LnsInt {
    return self.aliasSrcTypeInfo. FP.Get_kind( )
}
func (self *Ast_AliasTypeInfo) Get_mutMode() LnsInt {
    return self.aliasSrcTypeInfo. FP.Get_mutMode( )
}
func (self *Ast_AliasTypeInfo) Get_nilable() bool {
    return self.aliasSrcTypeInfo. FP.Get_nilable( )
}
func (self *Ast_AliasTypeInfo) Get_processInfo() *Ast_ProcessInfo {
    return self.aliasSrcTypeInfo. FP.Get_processInfo( )
}
func (self *Ast_AliasTypeInfo) Get_retTypeInfoList() *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_retTypeInfoList( )
}
func (self *Ast_AliasTypeInfo) Get_scope() LnsAny {
    return self.aliasSrcTypeInfo. FP.Get_scope( )
}
func (self *Ast_AliasTypeInfo) Get_staticFlag() bool {
    return self.aliasSrcTypeInfo. FP.Get_staticFlag( )
}
func (self *Ast_AliasTypeInfo) Get_typeData() *Ast_TypeData {
    return self.aliasSrcTypeInfo. FP.Get_typeData( )
}
func (self *Ast_AliasTypeInfo) HasBase() bool {
    return self.aliasSrcTypeInfo. FP.HasBase( )
}
func (self *Ast_AliasTypeInfo) HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool {
    return self.aliasSrcTypeInfo. FP.HasRouteNamespaceFrom( arg1)
}
func (self *Ast_AliasTypeInfo) IsInheritFrom(arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.aliasSrcTypeInfo. FP.IsInheritFrom( arg1,arg2,arg3)
}
func (self *Ast_AliasTypeInfo) IsModule() bool {
    return self.aliasSrcTypeInfo. FP.IsModule( )
}
func (self *Ast_AliasTypeInfo) SerializeTypeInfoList(arg1 string,arg2 *LnsList,arg3 LnsAny) string {
    return self.aliasSrcTypeInfo. FP.SerializeTypeInfoList( arg1,arg2,arg3)
}
func (self *Ast_AliasTypeInfo) SwitchScope(arg1 *Ast_Scope) {
self.aliasSrcTypeInfo. FP.SwitchScope( arg1)
}
// 1494: DeclConstr
func (self *Ast_AliasTypeInfo) InitAst_AliasTypeInfo(processInfo *Ast_ProcessInfo,rawTxt string,accessMode LnsInt,parentInfo *Ast_TypeInfo,aliasSrcTypeInfo *Ast_TypeInfo,externalFlag bool) {
    self.InitAst_TypeInfo(nil, processInfo)
    self.rawTxt = rawTxt
    
    self.accessMode = accessMode
    
    self.parentInfo = parentInfo
    
    self.aliasSrcTypeInfo = aliasSrcTypeInfo
    
    self.externalFlag = externalFlag
    
    self.typeId = processInfo.FP.Get_idProv().FP.GetNewId()
    
    self.nilableTypeInfo = &NewAst_NilableTypeInfo(nil, processInfo, &self.Ast_TypeInfo, processInfo.FP.Get_idProv().FP.GetNewId()).Ast_TypeInfo
    
}

// 1509: decl @lune.@base.@Ast.AliasTypeInfo.getParentFullName
func (self *Ast_AliasTypeInfo) GetParentFullName(typeNameCtrl *Ast_TypeNameCtrl,importInfo LnsAny,localFlag LnsAny) string {
    return typeNameCtrl.FP.GetParentFullName(&self.Ast_TypeInfo, importInfo, localFlag)
}

// 1516: decl @lune.@base.@Ast.AliasTypeInfo.getFullName
func (self *Ast_AliasTypeInfo) GetFullName(typeNameCtrl *Ast_TypeNameCtrl,importInfo Ast_ModuleInfoManager,localFlag LnsAny) string {
    return self.FP.GetParentFullName(typeNameCtrl, importInfo, localFlag) + self.FP.Get_rawTxt()
}

// 1524: decl @lune.@base.@Ast.AliasTypeInfo.get_aliasSrc
func (self *Ast_AliasTypeInfo) Get_aliasSrc() *Ast_TypeInfo {
    return self.aliasSrcTypeInfo
}

// 1528: decl @lune.@base.@Ast.AliasTypeInfo.get_nonnilableType
func (self *Ast_AliasTypeInfo) Get_nonnilableType() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1532: decl @lune.@base.@Ast.AliasTypeInfo.get_srcTypeInfo
func (self *Ast_AliasTypeInfo) Get_srcTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1535: decl @lune.@base.@Ast.AliasTypeInfo.get_genSrcTypeInfo
func (self *Ast_AliasTypeInfo) Get_genSrcTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1539: decl @lune.@base.@Ast.AliasTypeInfo.getModule
func (self *Ast_AliasTypeInfo) GetModule() *Ast_TypeInfo {
    return self.FP.Get_parentInfo().FP.GetModule()
}

// 1544: decl @lune.@base.@Ast.AliasTypeInfo.getTxt
func (self *Ast_AliasTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.rawTxt, typeNameCtrl, importInfo, localFlag)
}

// 1550: decl @lune.@base.@Ast.AliasTypeInfo.serialize
func (self *Ast_AliasTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    var parentId LnsInt
    parentId = self.FP.GetParentId()
    stream.Write(Lns_getVM().String_format("{ skind = %d, parentId = %d, typeId = %d, rawTxt = %q, srcTypeId = %d }\n", []LnsAny{Ast_SerializeKind__Alias, parentId, self.typeId, self.rawTxt, self.aliasSrcTypeInfo.FP.Get_typeId()}))
}

// 1558: decl @lune.@base.@Ast.AliasTypeInfo.get_display_stirng
func (self *Ast_AliasTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.rawTxt, nil)
}

// 1563: decl @lune.@base.@Ast.AliasTypeInfo.getParentId
func (self *Ast_AliasTypeInfo) GetParentId() LnsInt {
    return self.parentInfo.FP.Get_typeId()
}

// 1567: decl @lune.@base.@Ast.AliasTypeInfo.applyGeneric
func (self *Ast_AliasTypeInfo) ApplyGeneric(alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.aliasSrcTypeInfo.FP.ApplyGeneric(alt2typeMap, moduleTypeInfo)
    if typeInfo == self.aliasSrcTypeInfo{
        return &self.Ast_TypeInfo
    }
    return nil
}

// 1577: decl @lune.@base.@Ast.AliasTypeInfo.canEvalWith
func (self *Ast_AliasTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return self.aliasSrcTypeInfo.FP.CanEvalWith(processInfo, other.FP.Get_aliasSrc(), canEvalType, alt2type)
}

// 1585: decl @lune.@base.@Ast.AliasTypeInfo.equals
func (self *Ast_AliasTypeInfo) Equals(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    return self.aliasSrcTypeInfo.FP.Equals(processInfo, typeInfo.FP.Get_aliasSrc(), alt2type, checkModifer)
}


// declaration Class -- NilTypeInfo
type Ast_NilTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_NilTypeInfo struct {
    Ast_TypeInfo
    typeId LnsInt
    FP Ast_NilTypeInfoMtd
}
func Ast_NilTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_NilTypeInfo).FP
}
type Ast_NilTypeInfoDownCast interface {
    ToAst_NilTypeInfo() *Ast_NilTypeInfo
}
func Ast_NilTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_NilTypeInfoDownCast)
    if ok { return work.ToAst_NilTypeInfo() }
    return nil
}
func (obj *Ast_NilTypeInfo) ToAst_NilTypeInfo() *Ast_NilTypeInfo {
    return obj
}
func NewAst_NilTypeInfo(arg1 *Ast_ProcessInfo) *Ast_NilTypeInfo {
    obj := &Ast_NilTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_NilTypeInfo(arg1)
    return obj
}
func (self *Ast_NilTypeInfo) Get_typeId() LnsInt{ return self.typeId }
// 2224: DeclConstr
func (self *Ast_NilTypeInfo) InitAst_NilTypeInfo(processInfo *Ast_ProcessInfo) {
    self.InitAst_TypeInfo(nil, processInfo)
    processInfo.FP.Get_idProv().FP.Increment()
    self.typeId = processInfo.FP.Get_idProv().FP.Get_id()
    
}

// 2232: decl @lune.@base.@Ast.NilTypeInfo.isModule
func (self *Ast_NilTypeInfo) IsModule() bool {
    return false
}

// 2237: decl @lune.@base.@Ast.NilTypeInfo.getTxt
func (self *Ast_NilTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 2243: decl @lune.@base.@Ast.NilTypeInfo.getTxtWithRaw
func (self *Ast_NilTypeInfo) GetTxtWithRaw(raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return "nil"
}

// 2252: decl @lune.@base.@Ast.NilTypeInfo.canEvalWith
func (self *Ast_NilTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return other.FP.Get_nilable(), nil
}

// 2260: decl @lune.@base.@Ast.NilTypeInfo.get_display_stirng_with
func (self *Ast_NilTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(raw, nil, nil, nil)
}

// 2264: decl @lune.@base.@Ast.NilTypeInfo.get_display_stirng
func (self *Ast_NilTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with("nil", nil)
}

// 2270: decl @lune.@base.@Ast.NilTypeInfo.equals
func (self *Ast_NilTypeInfo) Equals(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    return &self.Ast_TypeInfo == typeInfo
}

// 2278: decl @lune.@base.@Ast.NilTypeInfo.get_parentInfo
func (self *Ast_NilTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 2286: decl @lune.@base.@Ast.NilTypeInfo.hasRouteNamespaceFrom
func (self *Ast_NilTypeInfo) HasRouteNamespaceFrom(other *Ast_TypeInfo) bool {
    return true
}

// 2291: decl @lune.@base.@Ast.NilTypeInfo.get_rawTxt
func (self *Ast_NilTypeInfo) Get_rawTxt() string {
    return "nil"
}

// 2295: decl @lune.@base.@Ast.NilTypeInfo.get_kind
func (self *Ast_NilTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Prim
}

// 2299: decl @lune.@base.@Ast.NilTypeInfo.get_baseTypeInfo
func (self *Ast_NilTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 2303: decl @lune.@base.@Ast.NilTypeInfo.get_nilable
func (self *Ast_NilTypeInfo) Get_nilable() bool {
    return true
}

// 2307: decl @lune.@base.@Ast.NilTypeInfo.get_mutMode
func (self *Ast_NilTypeInfo) Get_mutMode() LnsInt {
    return Ast_MutMode__IMut
}

// 2312: decl @lune.@base.@Ast.NilTypeInfo.getParentFullName
func (self *Ast_NilTypeInfo) GetParentFullName(typeNameCtrl *Ast_TypeNameCtrl,importInfo LnsAny,localFlag LnsAny) string {
    return ""
}


// declaration Class -- AccessSymbolInfo
type Ast_AccessSymbolInfoMtd interface {
    CanAccess(arg1 *Ast_Scope, arg2 LnsInt) LnsAny
    ClearValue()
    GetModule() *Ast_TypeInfo
    GetOrg() *Ast_SymbolInfo
    Get_accessMode() LnsInt
    Get_canBeLeft() bool
    Get_canBeRight() bool
    Get_convModuleParam() LnsAny
    Get_hasAccessFromClosure() bool
    Get_hasValueFlag() bool
    Get_isLazyLoad() bool
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_mutable() bool
    Get_name() string
    Get_namespaceTypeInfo() *Ast_TypeInfo
    Get_pos() LnsAny
    Get_posForLatestMod() LnsAny
    Get_posForModToRef() LnsAny
    Get_scope() *Ast_Scope
    Get_staticFlag() bool
    Get_symbolId() LnsInt
    Get_symbolInfo() *Ast_SymbolInfo
    Get_typeInfo() *Ast_TypeInfo
    HasAccess() bool
    Set_convModuleParam(arg1 LnsAny)
    Set_hasAccessFromClosure(arg1 bool)
    Set_hasValueFlag(arg1 bool)
    Set_posForLatestMod(arg1 LnsAny)
    Set_posForModToRef(arg1 LnsAny)
    Set_typeInfo(arg1 *Ast_TypeInfo)
    UpdateValue(arg1 LnsAny)
}
type Ast_AccessSymbolInfo struct {
    Ast_SymbolInfo
    symbolInfo *Ast_SymbolInfo
    overrideMut LnsAny
    overrideCanBeLeft bool
    overrideTypeInfo *Ast_TypeInfo
    FP Ast_AccessSymbolInfoMtd
}
func Ast_AccessSymbolInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_AccessSymbolInfo).FP
}
type Ast_AccessSymbolInfoDownCast interface {
    ToAst_AccessSymbolInfo() *Ast_AccessSymbolInfo
}
func Ast_AccessSymbolInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_AccessSymbolInfoDownCast)
    if ok { return work.ToAst_AccessSymbolInfo() }
    return nil
}
func (obj *Ast_AccessSymbolInfo) ToAst_AccessSymbolInfo() *Ast_AccessSymbolInfo {
    return obj
}
func NewAst_AccessSymbolInfo(arg1 *Ast_SymbolInfo, arg2 LnsAny, arg3 bool) *Ast_AccessSymbolInfo {
    obj := &Ast_AccessSymbolInfo{}
    obj.FP = obj
    obj.Ast_SymbolInfo.FP = obj
    obj.InitAst_AccessSymbolInfo(arg1, arg2, arg3)
    return obj
}
func (self *Ast_AccessSymbolInfo) Get_symbolInfo() *Ast_SymbolInfo{ return self.symbolInfo }
func (self *Ast_AccessSymbolInfo) ClearValue() {
self.symbolInfo. FP.ClearValue( )
}
func (self *Ast_AccessSymbolInfo) GetModule() *Ast_TypeInfo {
    return self.symbolInfo. FP.GetModule( )
}
func (self *Ast_AccessSymbolInfo) Get_accessMode() LnsInt {
    return self.symbolInfo. FP.Get_accessMode( )
}
func (self *Ast_AccessSymbolInfo) Get_canBeRight() bool {
    return self.symbolInfo. FP.Get_canBeRight( )
}
func (self *Ast_AccessSymbolInfo) Get_convModuleParam() LnsAny {
    return self.symbolInfo. FP.Get_convModuleParam( )
}
func (self *Ast_AccessSymbolInfo) Get_hasAccessFromClosure() bool {
    return self.symbolInfo. FP.Get_hasAccessFromClosure( )
}
func (self *Ast_AccessSymbolInfo) Get_hasValueFlag() bool {
    return self.symbolInfo. FP.Get_hasValueFlag( )
}
func (self *Ast_AccessSymbolInfo) Get_isLazyLoad() bool {
    return self.symbolInfo. FP.Get_isLazyLoad( )
}
func (self *Ast_AccessSymbolInfo) Get_kind() LnsInt {
    return self.symbolInfo. FP.Get_kind( )
}
func (self *Ast_AccessSymbolInfo) Get_name() string {
    return self.symbolInfo. FP.Get_name( )
}
func (self *Ast_AccessSymbolInfo) Get_namespaceTypeInfo() *Ast_TypeInfo {
    return self.symbolInfo. FP.Get_namespaceTypeInfo( )
}
func (self *Ast_AccessSymbolInfo) Get_pos() LnsAny {
    return self.symbolInfo. FP.Get_pos( )
}
func (self *Ast_AccessSymbolInfo) Get_posForLatestMod() LnsAny {
    return self.symbolInfo. FP.Get_posForLatestMod( )
}
func (self *Ast_AccessSymbolInfo) Get_posForModToRef() LnsAny {
    return self.symbolInfo. FP.Get_posForModToRef( )
}
func (self *Ast_AccessSymbolInfo) Get_scope() *Ast_Scope {
    return self.symbolInfo. FP.Get_scope( )
}
func (self *Ast_AccessSymbolInfo) Get_staticFlag() bool {
    return self.symbolInfo. FP.Get_staticFlag( )
}
func (self *Ast_AccessSymbolInfo) Get_symbolId() LnsInt {
    return self.symbolInfo. FP.Get_symbolId( )
}
func (self *Ast_AccessSymbolInfo) HasAccess() bool {
    return self.symbolInfo. FP.HasAccess( )
}
func (self *Ast_AccessSymbolInfo) Set_convModuleParam(arg1 LnsAny) {
self.symbolInfo. FP.Set_convModuleParam( arg1)
}
func (self *Ast_AccessSymbolInfo) Set_hasAccessFromClosure(arg1 bool) {
self.symbolInfo. FP.Set_hasAccessFromClosure( arg1)
}
func (self *Ast_AccessSymbolInfo) Set_hasValueFlag(arg1 bool) {
self.symbolInfo. FP.Set_hasValueFlag( arg1)
}
func (self *Ast_AccessSymbolInfo) Set_posForLatestMod(arg1 LnsAny) {
self.symbolInfo. FP.Set_posForLatestMod( arg1)
}
func (self *Ast_AccessSymbolInfo) Set_posForModToRef(arg1 LnsAny) {
self.symbolInfo. FP.Set_posForModToRef( arg1)
}
func (self *Ast_AccessSymbolInfo) Set_typeInfo(arg1 *Ast_TypeInfo) {
self.symbolInfo. FP.Set_typeInfo( arg1)
}
func (self *Ast_AccessSymbolInfo) UpdateValue(arg1 LnsAny) {
self.symbolInfo. FP.UpdateValue( arg1)
}
// 2402: DeclConstr
func (self *Ast_AccessSymbolInfo) InitAst_AccessSymbolInfo(symbolInfo *Ast_SymbolInfo,overrideMut LnsAny,overrideCanBeLeft bool) {
    self.InitAst_SymbolInfo()
    self.symbolInfo = symbolInfo
    
    self.overrideMut = overrideMut
    
    self.overrideCanBeLeft = overrideCanBeLeft
    
    var symType *Ast_TypeInfo
    symType = symbolInfo.FP.Get_typeInfo()
    var work *Ast_TypeInfo
    switch _exp8090 := self.overrideMut.(type) {
    case *Ast_OverrideMut__None:
        work = symType
        
    case *Ast_OverrideMut__Prefix:
    prefixTypeInfo := _exp8090.Val1
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( self.symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mbr) &&
            Lns_GetEnv().SetStackVal( symType.FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
            Lns_GetEnv().SetStackVal( prefixTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Class) &&
            Lns_GetEnv().SetStackVal( prefixTypeInfo.FP.Get_itemTypeInfoList().Len() > 0) ).(bool)){
            var alt2TypeMap *LnsMap
            alt2TypeMap = prefixTypeInfo.FP.CreateAlt2typeMap(false)
            var typeInfo LnsAny
            typeInfo = symType.FP.ApplyGeneric(alt2TypeMap, symType.FP.GetModule())
            if typeInfo != nil{
                typeInfo_2242 := typeInfo.(*Ast_TypeInfo)
                work = typeInfo_2242
                
            } else {
                work = symType
                
            }
        } else { 
            work = symType
            
        }
    case *Ast_OverrideMut__IMut:
    typeInfo := _exp8090.Val1
        work = typeInfo
        
    }
    self.overrideTypeInfo = work
    
}

// 2441: decl @lune.@base.@Ast.AccessSymbolInfo.getOrg
func (self *Ast_AccessSymbolInfo) GetOrg() *Ast_SymbolInfo {
    return self.symbolInfo.FP.GetOrg()
}

// 2445: decl @lune.@base.@Ast.AccessSymbolInfo.canAccess
func (self *Ast_AccessSymbolInfo) CanAccess(fromScope *Ast_Scope,access LnsInt) LnsAny {
    if Lns_isCondTrue( self.symbolInfo.FP.CanAccess(fromScope, access)){
        return &self.Ast_SymbolInfo
    }
    return nil
}

// 2454: decl @lune.@base.@Ast.AccessSymbolInfo.get_typeInfo
func (self *Ast_AccessSymbolInfo) Get_typeInfo() *Ast_TypeInfo {
    return self.overrideTypeInfo
}

// 2458: decl @lune.@base.@Ast.AccessSymbolInfo.get_mutMode
func (self *Ast_AccessSymbolInfo) Get_mutMode() LnsInt {
    switch _exp8218 := self.overrideMut.(type) {
    case *Ast_OverrideMut__None:
    case *Ast_OverrideMut__Prefix:
    prefixTypeInfo := _exp8218.Val1
        if _switch8211 := self.symbolInfo.FP.Get_mutMode(); _switch8211 == Ast_MutMode__AllMut || _switch8211 == Ast_MutMode__IMut || _switch8211 == Ast_MutMode__IMutRe {
            return self.symbolInfo.FP.Get_mutMode()
        } else if _switch8211 == Ast_MutMode__Mut {
            if _switch8209 := prefixTypeInfo.FP.Get_mutMode(); _switch8209 == Ast_MutMode__AllMut {
                return Ast_MutMode__Mut
            } else if _switch8209 == Ast_MutMode__Mut || _switch8209 == Ast_MutMode__IMut || _switch8209 == Ast_MutMode__IMutRe {
                return prefixTypeInfo.FP.Get_mutMode()
            }
        }
    case *Ast_OverrideMut__IMut:
        return Ast_MutMode__IMut
    }
    return self.symbolInfo.FP.Get_mutMode()
}

// 2487: decl @lune.@base.@Ast.AccessSymbolInfo.get_mutable
func (self *Ast_AccessSymbolInfo) Get_mutable() bool {
    return Ast_isMutable(self.FP.Get_mutMode())
}

// 2491: decl @lune.@base.@Ast.AccessSymbolInfo.get_canBeLeft
func (self *Ast_AccessSymbolInfo) Get_canBeLeft() bool {
    if Lns_op_not(self.overrideCanBeLeft){
        return false
    }
    return self.symbolInfo.FP.Get_canBeLeft()
}


// declaration Class -- AlternateTypeInfo
type Ast_AlternateTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    canSetFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 *LnsMap) bool
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_altIndex() LnsInt
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_txt() string
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_AlternateTypeInfo struct {
    Ast_TypeInfo
    typeId LnsInt
    txt string
    moduleTypeInfo *Ast_TypeInfo
    nilableTypeInfo *Ast_NilableTypeInfo
    accessMode LnsInt
    baseTypeInfo *Ast_TypeInfo
    interfaceList *LnsList
    belongClassFlag bool
    altIndex LnsInt
    FP Ast_AlternateTypeInfoMtd
}
func Ast_AlternateTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_AlternateTypeInfo).FP
}
type Ast_AlternateTypeInfoDownCast interface {
    ToAst_AlternateTypeInfo() *Ast_AlternateTypeInfo
}
func Ast_AlternateTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_AlternateTypeInfoDownCast)
    if ok { return work.ToAst_AlternateTypeInfo() }
    return nil
}
func (obj *Ast_AlternateTypeInfo) ToAst_AlternateTypeInfo() *Ast_AlternateTypeInfo {
    return obj
}
func NewAst_AlternateTypeInfo(arg1 *Ast_ProcessInfo, arg2 bool, arg3 LnsInt, arg4 string, arg5 LnsInt, arg6 *Ast_TypeInfo, arg7 LnsAny, arg8 LnsAny) *Ast_AlternateTypeInfo {
    obj := &Ast_AlternateTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AlternateTypeInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Ast_AlternateTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_AlternateTypeInfo) Get_txt() string{ return self.txt }
func (self *Ast_AlternateTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo{ return &self.nilableTypeInfo.Ast_TypeInfo }
func (self *Ast_AlternateTypeInfo) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Ast_AlternateTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo{ return self.baseTypeInfo }
func (self *Ast_AlternateTypeInfo) Get_interfaceList() *LnsList{ return self.interfaceList }
func (self *Ast_AlternateTypeInfo) Get_altIndex() LnsInt{ return self.altIndex }
// 2525: DeclConstr
func (self *Ast_AlternateTypeInfo) InitAst_AlternateTypeInfo(processInfo *Ast_ProcessInfo,belongClassFlag bool,altIndex LnsInt,txt string,accessMode LnsInt,moduleTypeInfo *Ast_TypeInfo,baseTypeInfo LnsAny,interfaceList LnsAny) {
    self.InitAst_TypeInfo(Ast_TypeInfo_createScope(processInfo, nil, true, baseTypeInfo, interfaceList), processInfo)
    processInfo.FP.Get_idProv().FP.Increment()
    self.typeId = processInfo.FP.Get_idProv().FP.Get_id()
    
    self.txt = txt
    
    self.accessMode = accessMode
    
    self.moduleTypeInfo = moduleTypeInfo
    
    self.baseTypeInfo = Lns_unwrapDefault( baseTypeInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.interfaceList = Lns_unwrapDefault( interfaceList, NewLnsList([]LnsAny{})).(*LnsList)
    
    self.belongClassFlag = belongClassFlag
    
    self.altIndex = altIndex
    
    processInfo.FP.Get_idProv().FP.Increment()
    self.nilableTypeInfo = NewAst_NilableTypeInfo(nil, processInfo, &self.Ast_TypeInfo, processInfo.FP.Get_idProv().FP.Get_id())
    
}

// 2549: decl @lune.@base.@Ast.AlternateTypeInfo.isModule
func (self *Ast_AlternateTypeInfo) IsModule() bool {
    return false
}

// 2554: decl @lune.@base.@Ast.AlternateTypeInfo.getParentId
func (self *Ast_AlternateTypeInfo) GetParentId() LnsInt {
    return self.moduleTypeInfo.FP.Get_typeId()
}

// 2558: decl @lune.@base.@Ast.AlternateTypeInfo.get_baseId
func (self *Ast_AlternateTypeInfo) Get_baseId() LnsInt {
    return self.baseTypeInfo.FP.Get_typeId()
}

// 2562: decl @lune.@base.@Ast.AlternateTypeInfo.get_parentInfo
func (self *Ast_AlternateTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return self.moduleTypeInfo
}

// 2567: decl @lune.@base.@Ast.AlternateTypeInfo.getTxt
func (self *Ast_AlternateTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 2573: decl @lune.@base.@Ast.AlternateTypeInfo.getTxtWithRaw
func (self *Ast_AlternateTypeInfo) GetTxtWithRaw(raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.txt
}



// 2589: decl @lune.@base.@Ast.AlternateTypeInfo.isInheritFrom
func (self *Ast_AlternateTypeInfo) IsInheritFrom(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    var workAlt2type LnsAny
    if alt2type != nil{
        alt2type_2388 := alt2type.(*LnsMap)
        var otherWork *Ast_TypeInfo
        otherWork = Ast_AlternateTypeInfo_getAssign(other, alt2type_2388)
        if &self.Ast_TypeInfo == otherWork.FP.Get_srcTypeInfo(){
            return true
        }
        {
            _genType := alt2type_2388.Items[&self.Ast_TypeInfo]
            if _genType != nil {
                genType := _genType.(*Ast_TypeInfo)
                return genType.FP.IsInheritFrom(processInfo, otherWork, alt2type_2388)
            }
        }
        if Lns_op_not(Ast_CanEvalCtrlTypeInfo_isValidApply(alt2type_2388)){
            workAlt2type = nil
            
        } else { 
            workAlt2type = alt2type_2388
            
        }
    } else {
        workAlt2type = nil
        
    }
    if &self.Ast_TypeInfo == other.FP.Get_srcTypeInfo(){
        return true
    }
    var check func() bool
    check = func() bool {
        if self.FP.HasBase(){
            if self.baseTypeInfo.FP.IsInheritFrom(processInfo, other, workAlt2type){
                return true
            }
        }
        for _, _ifType := range( self.interfaceList.Items ) {
            ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if ifType.FP.IsInheritFrom(processInfo, other, workAlt2type){
                return true
            }
        }
        return false
    }
    if check(){
        if workAlt2type != nil{
            workAlt2type_2407 := workAlt2type.(*LnsMap)
            workAlt2type_2407.Set(&self.Ast_TypeInfo,other)
        }
        return true
    }
    return false
}

// 2640: decl @lune.@base.@Ast.AlternateTypeInfo.canEvalWith
func (self *Ast_AlternateTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if &self.Ast_TypeInfo == other.FP.Get_srcTypeInfo(){
        return true, nil
    }
    if other.FP.Get_nilable(){
        return false, nil
    }
    if Lns_op_not(self.belongClassFlag){
        {
            _altType := Ast_AlternateTypeInfoDownCastF(other.FP)
            if _altType != nil {
                altType := _altType.(*Ast_AlternateTypeInfo)
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_op_not(altType.belongClassFlag)) &&
                    Lns_GetEnv().SetStackVal( altType.altIndex == self.altIndex) ).(bool)){
                    return true, nil
                }
            }
        }
    }
    return self.FP.canSetFrom(processInfo, other, canEvalType, alt2type), nil
}

// 2666: decl @lune.@base.@Ast.AlternateTypeInfo.get_display_stirng_with
func (self *Ast_AlternateTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    if alt2type != nil{
        alt2type_2431 := alt2type.(*LnsMap)
        {
            _genType := alt2type_2431.Items[&self.Ast_TypeInfo]
            if _genType != nil {
                genType := _genType.(*Ast_TypeInfo)
                return genType.FP.Get_display_stirng_with(genType.FP.Get_rawTxt(), alt2type_2431)
            }
        }
    }
    return self.FP.GetTxtWithRaw(raw, nil, nil, nil)
}

// 2676: decl @lune.@base.@Ast.AlternateTypeInfo.get_display_stirng
func (self *Ast_AlternateTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.txt, nil)
}

// 2682: decl @lune.@base.@Ast.AlternateTypeInfo.equals
func (self *Ast_AlternateTypeInfo) Equals(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if &self.Ast_TypeInfo == typeInfo{
        return true
    }
    if Lns_op_not(self.belongClassFlag){
        {
            _altType := Ast_AlternateTypeInfoDownCastF(typeInfo.FP)
            if _altType != nil {
                altType := _altType.(*Ast_AlternateTypeInfo)
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_op_not(altType.belongClassFlag)) &&
                    Lns_GetEnv().SetStackVal( altType.altIndex == self.altIndex) ).(bool)){
                    return true
                }
            }
        }
    }
    if alt2type != nil{
        alt2type_2454 := alt2type.(*LnsMap)
        return self.FP.canSetFrom(processInfo, typeInfo, nil, alt2type_2454)
    }
    return false
}

// 2710: decl @lune.@base.@Ast.AlternateTypeInfo.hasRouteNamespaceFrom
func (self *Ast_AlternateTypeInfo) HasRouteNamespaceFrom(other *Ast_TypeInfo) bool {
    return true
}

// 2715: decl @lune.@base.@Ast.AlternateTypeInfo.get_rawTxt
func (self *Ast_AlternateTypeInfo) Get_rawTxt() string {
    return self.txt
}

// 2719: decl @lune.@base.@Ast.AlternateTypeInfo.get_kind
func (self *Ast_AlternateTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Alternate
}

// 2723: decl @lune.@base.@Ast.AlternateTypeInfo.get_nilable
func (self *Ast_AlternateTypeInfo) Get_nilable() bool {
    return false
}

// 2727: decl @lune.@base.@Ast.AlternateTypeInfo.get_mutMode
func (self *Ast_AlternateTypeInfo) Get_mutMode() LnsInt {
    return Ast_MutMode__Mut
}

// 2732: decl @lune.@base.@Ast.AlternateTypeInfo.getParentFullName
func (self *Ast_AlternateTypeInfo) GetParentFullName(typeNameCtrl *Ast_TypeNameCtrl,importInfo LnsAny,localFlag LnsAny) string {
    return ""
}

// 2739: decl @lune.@base.@Ast.AlternateTypeInfo.serialize
func (self *Ast_AlternateTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    var parentId LnsInt
    parentId = self.FP.GetParentId()
    stream.Write(Lns_getVM().String_format("{ skind = %d, parentId = %d, typeId = %d, txt = %q, ", []LnsAny{Ast_SerializeKind__Alternate, parentId, self.typeId, self.txt}) + Lns_getVM().String_format("accessMode = %d, baseId = %d, ", []LnsAny{self.accessMode, self.FP.Get_baseId()}) + Lns_getVM().String_format("belongClassFlag = %s, altIndex = %d, ", []LnsAny{self.belongClassFlag, self.altIndex}))
    stream.Write(self.FP.SerializeTypeInfoList("ifList = {", self.interfaceList, nil))
    stream.Write("}\n")
}

// 2751: decl @lune.@base.@Ast.AlternateTypeInfo.applyGeneric
func (self *Ast_AlternateTypeInfo) ApplyGeneric(alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    return Ast_AlternateTypeInfo_getAssign(&self.Ast_TypeInfo, alt2typeMap)
}

// 4279: decl @lune.@base.@Ast.AlternateTypeInfo.canSetFrom
func (self *Ast_AlternateTypeInfo) canSetFrom(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsAny,alt2type *LnsMap) bool {
    var otherWork *Ast_TypeInfo
    otherWork = Ast_AlternateTypeInfo_getAssign(other, alt2type)
    if &self.Ast_TypeInfo == otherWork{
        return true
    }
    {
        _genType := alt2type.Items[&self.Ast_TypeInfo]
        if _genType != nil {
            genType := _genType.(*Ast_TypeInfo)
            if canEvalType != nil{
                canEvalType_3767 := canEvalType.(LnsInt)
                return Lns_car(genType.FP.CanEvalWith(processInfo, otherWork, canEvalType_3767, alt2type)).(bool)
            }
            return genType.FP.Equals(processInfo, otherWork, alt2type, nil)
        }
    }
    var workAlt2type *LnsMap
    if Lns_op_not(Ast_CanEvalCtrlTypeInfo_isValidApply(alt2type)){
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isClass(otherWork))) &&
            Lns_GetEnv().SetStackVal( otherWork.FP.Get_kind() != Ast_TypeInfoKind__IF) ).(bool)){
            return false
        }
        workAlt2type = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false)
        
    } else { 
        workAlt2type = alt2type
        
    }
    if self.FP.HasBase(){
        if Lns_op_not(other.FP.IsInheritFrom(processInfo, self.baseTypeInfo, workAlt2type)){
            return false
        }
    }
    for _, _ifType := range( self.interfaceList.Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(other.FP.IsInheritFrom(processInfo, ifType, workAlt2type)){
            return false
        }
    }
    workAlt2type.Set(&self.Ast_TypeInfo,otherWork)
    return true
}

// 4426: decl @lune.@base.@Ast.AlternateTypeInfo.getAssign
func Ast_AlternateTypeInfo_getAssign(typeInfo *Ast_TypeInfo,alt2type *LnsMap) *Ast_TypeInfo {
    if typeInfo.FP.Get_kind() != Ast_TypeInfoKind__Alternate{
        return typeInfo
    }
    var otherWork *Ast_TypeInfo
    otherWork = typeInfo
    for  {
        {
            __exp := alt2type.Items[otherWork]
            if __exp != nil {
                _exp := __exp.(*Ast_TypeInfo)
                if _exp != otherWork{
                    otherWork = _exp
                    
                } else { 
                    return otherWork
                }
            } else {
                return otherWork
            }
        }
    }
// insert a dummy
    return nil
}


// declaration Class -- BoxTypeInfo
type Ast_BoxTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_boxingType() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_BoxTypeInfo struct {
    Ast_TypeInfo
    boxingType *Ast_TypeInfo
    typeId LnsInt
    itemTypeInfoList *LnsList
    accessMode LnsInt
    nilableTypeInfo *Ast_NilableTypeInfo
    FP Ast_BoxTypeInfoMtd
}
func Ast_BoxTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_BoxTypeInfo).FP
}
type Ast_BoxTypeInfoDownCast interface {
    ToAst_BoxTypeInfo() *Ast_BoxTypeInfo
}
func Ast_BoxTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_BoxTypeInfoDownCast)
    if ok { return work.ToAst_BoxTypeInfo() }
    return nil
}
func (obj *Ast_BoxTypeInfo) ToAst_BoxTypeInfo() *Ast_BoxTypeInfo {
    return obj
}
func NewAst_BoxTypeInfo(arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 LnsInt, arg4 *Ast_TypeInfo) *Ast_BoxTypeInfo {
    obj := &Ast_BoxTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_BoxTypeInfo(arg1, arg2, arg3, arg4)
    return obj
}
func (self *Ast_BoxTypeInfo) Get_boxingType() *Ast_TypeInfo{ return self.boxingType }
func (self *Ast_BoxTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_BoxTypeInfo) Get_itemTypeInfoList() *LnsList{ return self.itemTypeInfoList }
func (self *Ast_BoxTypeInfo) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Ast_BoxTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo{ return &self.nilableTypeInfo.Ast_TypeInfo }
func (self *Ast_BoxTypeInfo) AddChildren(arg1 *Ast_TypeInfo) {
self.boxingType. FP.AddChildren( arg1)
}
func (self *Ast_BoxTypeInfo) GetFullName(arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.boxingType. FP.GetFullName( arg1,arg2,arg3)
}
func (self *Ast_BoxTypeInfo) GetModule() *Ast_TypeInfo {
    return self.boxingType. FP.GetModule( )
}
func (self *Ast_BoxTypeInfo) GetOverridingType() LnsAny {
    return self.boxingType. FP.GetOverridingType( )
}
func (self *Ast_BoxTypeInfo) GetParentFullName(arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.boxingType. FP.GetParentFullName( arg1,arg2,arg3)
}
func (self *Ast_BoxTypeInfo) GetParentId() LnsInt {
    return self.boxingType. FP.GetParentId( )
}
func (self *Ast_BoxTypeInfo) getProcessInfo() *Ast_ProcessInfo {
    return self.boxingType. FP.getProcessInfo( )
}
func (self *Ast_BoxTypeInfo) Get_abstractFlag() bool {
    return self.boxingType. FP.Get_abstractFlag( )
}
func (self *Ast_BoxTypeInfo) Get_argTypeInfoList() *LnsList {
    return self.boxingType. FP.Get_argTypeInfoList( )
}
func (self *Ast_BoxTypeInfo) Get_autoFlag() bool {
    return self.boxingType. FP.Get_autoFlag( )
}
func (self *Ast_BoxTypeInfo) Get_baseId() LnsInt {
    return self.boxingType. FP.Get_baseId( )
}
func (self *Ast_BoxTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return self.boxingType. FP.Get_baseTypeInfo( )
}
func (self *Ast_BoxTypeInfo) Get_children() *LnsList {
    return self.boxingType. FP.Get_children( )
}
func (self *Ast_BoxTypeInfo) Get_externalFlag() bool {
    return self.boxingType. FP.Get_externalFlag( )
}
func (self *Ast_BoxTypeInfo) Get_genSrcTypeInfo() *Ast_TypeInfo {
    return self.boxingType. FP.Get_genSrcTypeInfo( )
}
func (self *Ast_BoxTypeInfo) Get_interfaceList() *LnsList {
    return self.boxingType. FP.Get_interfaceList( )
}
func (self *Ast_BoxTypeInfo) Get_mutMode() LnsInt {
    return self.boxingType. FP.Get_mutMode( )
}
func (self *Ast_BoxTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return self.boxingType. FP.Get_parentInfo( )
}
func (self *Ast_BoxTypeInfo) Get_processInfo() *Ast_ProcessInfo {
    return self.boxingType. FP.Get_processInfo( )
}
func (self *Ast_BoxTypeInfo) Get_rawTxt() string {
    return self.boxingType. FP.Get_rawTxt( )
}
func (self *Ast_BoxTypeInfo) Get_retTypeInfoList() *LnsList {
    return self.boxingType. FP.Get_retTypeInfoList( )
}
func (self *Ast_BoxTypeInfo) Get_staticFlag() bool {
    return self.boxingType. FP.Get_staticFlag( )
}
func (self *Ast_BoxTypeInfo) Get_typeData() *Ast_TypeData {
    return self.boxingType. FP.Get_typeData( )
}
func (self *Ast_BoxTypeInfo) HasBase() bool {
    return self.boxingType. FP.HasBase( )
}
func (self *Ast_BoxTypeInfo) HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool {
    return self.boxingType. FP.HasRouteNamespaceFrom( arg1)
}
func (self *Ast_BoxTypeInfo) IsInheritFrom(arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.boxingType. FP.IsInheritFrom( arg1,arg2,arg3)
}
func (self *Ast_BoxTypeInfo) IsModule() bool {
    return self.boxingType. FP.IsModule( )
}
func (self *Ast_BoxTypeInfo) SerializeTypeInfoList(arg1 string,arg2 *LnsList,arg3 LnsAny) string {
    return self.boxingType. FP.SerializeTypeInfoList( arg1,arg2,arg3)
}
func (self *Ast_BoxTypeInfo) SwitchScope(arg1 *Ast_Scope) {
self.boxingType. FP.SwitchScope( arg1)
}
// 2769: DeclConstr
func (self *Ast_BoxTypeInfo) InitAst_BoxTypeInfo(processInfo *Ast_ProcessInfo,typeId LnsInt,accessMode LnsInt,boxingType *Ast_TypeInfo) {
    self.InitAst_TypeInfo(Ast_boxRootScope, processInfo)
    self.boxingType = boxingType
    
    self.typeId = typeId
    
    self.itemTypeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(boxingType)})
    
    self.accessMode = accessMode
    
    processInfo.FP.Get_idProv().FP.Increment()
    self.nilableTypeInfo = NewAst_NilableTypeInfo(nil, processInfo, &self.Ast_TypeInfo, processInfo.FP.Get_idProv().FP.Get_id())
    
}

// 2783: decl @lune.@base.@Ast.BoxTypeInfo.get_scope
func (self *Ast_BoxTypeInfo) Get_scope() LnsAny {
    return self.Ast_TypeInfo.Get_scope()
}

// 2787: decl @lune.@base.@Ast.BoxTypeInfo.get_kind
func (self *Ast_BoxTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Box
}

// 2790: decl @lune.@base.@Ast.BoxTypeInfo.get_aliasSrc
func (self *Ast_BoxTypeInfo) Get_aliasSrc() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 2793: decl @lune.@base.@Ast.BoxTypeInfo.get_srcTypeInfo
func (self *Ast_BoxTypeInfo) Get_srcTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 2796: decl @lune.@base.@Ast.BoxTypeInfo.get_nonnilableType
func (self *Ast_BoxTypeInfo) Get_nonnilableType() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 2799: decl @lune.@base.@Ast.BoxTypeInfo.get_nilable
func (self *Ast_BoxTypeInfo) Get_nilable() bool {
    return false
}

// 2802: decl @lune.@base.@Ast.BoxTypeInfo.get_extedType
func (self *Ast_BoxTypeInfo) Get_extedType() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 2808: decl @lune.@base.@Ast.BoxTypeInfo.getTxt
func (self *Ast_BoxTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 2813: decl @lune.@base.@Ast.BoxTypeInfo.getTxtWithRaw
func (self *Ast_BoxTypeInfo) GetTxtWithRaw(raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return "Nilable<" + self.boxingType.FP.GetTxtWithRaw(raw, typeNameCtrl, importInfo, localFlag) + ">"
}

// 2821: decl @lune.@base.@Ast.BoxTypeInfo.get_display_stirng
func (self *Ast_BoxTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
}

// 2825: decl @lune.@base.@Ast.BoxTypeInfo.get_display_stirng_with
func (self *Ast_BoxTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    return Lns_getVM().String_format("Nilable<%s>", []LnsAny{self.boxingType.FP.Get_display_stirng_with(raw, alt2type)})
}

// 2830: decl @lune.@base.@Ast.BoxTypeInfo.serialize
func (self *Ast_BoxTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    stream.Write(Lns_getVM().String_format("{ skind = %d, typeId = %d, accessMode = %d, boxingType = %d }\n", []LnsAny{Ast_SerializeKind__Box, self.typeId, self.accessMode, self.boxingType.FP.Get_typeId()}))
}

// 2837: decl @lune.@base.@Ast.BoxTypeInfo.equals
func (self *Ast_BoxTypeInfo) Equals(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    {
        _boxType := Ast_BoxTypeInfoDownCastF(typeInfo.FP)
        if _boxType != nil {
            boxType := _boxType.(*Ast_BoxTypeInfo)
            return self.boxingType.FP.Equals(processInfo, boxType.boxingType, alt2type, checkModifer)
        }
    }
    return false
}

// 2848: decl @lune.@base.@Ast.BoxTypeInfo.createAlt2typeMap
func (self *Ast_BoxTypeInfo) CreateAlt2typeMap(detectFlag bool) *LnsMap {
    var _map *LnsMap
    _map = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(detectFlag)
    if self.boxingType != &Ast_boxRootAltType.Ast_TypeInfo{
        _map.Set(&Ast_boxRootAltType.Ast_TypeInfo,self.boxingType)
    }
    return _map
}

// 4496: decl @lune.@base.@Ast.BoxTypeInfo.applyGeneric
func (self *Ast_BoxTypeInfo) ApplyGeneric(alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.boxingType.FP.ApplyGeneric(alt2typeMap, moduleTypeInfo)
    if typeInfo == self.boxingType{
        return &self.Ast_TypeInfo
    }
    if typeInfo != nil{
        typeInfo_3886 := typeInfo.(*Ast_TypeInfo)
        return moduleTypeInfo.FP.getProcessInfo().FP.CreateBox(self.accessMode, typeInfo_3886)
    }
    return nil
}

// 5759: decl @lune.@base.@Ast.BoxTypeInfo.canEvalWith
func (self *Ast_BoxTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if &self.Ast_TypeInfo == other{
        return true, nil
    }
    if _switch23428 := canEvalType; _switch23428 == Ast_CanEvalType__SetOp || _switch23428 == Ast_CanEvalType__SetOpIMut || _switch23428 == Ast_CanEvalType__SetEq {
    } else {
        return false, nil
    }
    if other.FP.Get_nilable(){
        Ast_CanEvalCtrlTypeInfo_updateNeedAutoBoxing(alt2type)
        return true, nil
    }
    {
        _otherBoxType := Ast_BoxTypeInfoDownCastF(other.FP)
        if _otherBoxType != nil {
            otherBoxType := _otherBoxType.(*Ast_BoxTypeInfo)
            return self.boxingType.FP.CanEvalWith(processInfo, otherBoxType.boxingType, canEvalType, alt2type)
        }
    }
    if Lns_car(self.boxingType.FP.CanEvalWith(processInfo, other, canEvalType, alt2type)).(bool){
        Ast_CanEvalCtrlTypeInfo_updateNeedAutoBoxing(alt2type)
        return true, nil
    }
    return false, nil
}


// declaration Class -- GenericTypeInfo
type Ast_GenericTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_GenericTypeInfo struct {
    Ast_TypeInfo
    typeId LnsInt
    itemTypeInfoList *LnsList
    nilableTypeInfo *Ast_NilableTypeInfo
    genSrcTypeInfo *Ast_TypeInfo
    moduleTypeInfo *Ast_TypeInfo
    alt2typeMap *LnsMap
    hasAlter bool
    FP Ast_GenericTypeInfoMtd
}
func Ast_GenericTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_GenericTypeInfo).FP
}
type Ast_GenericTypeInfoDownCast interface {
    ToAst_GenericTypeInfo() *Ast_GenericTypeInfo
}
func Ast_GenericTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_GenericTypeInfoDownCast)
    if ok { return work.ToAst_GenericTypeInfo() }
    return nil
}
func (obj *Ast_GenericTypeInfo) ToAst_GenericTypeInfo() *Ast_GenericTypeInfo {
    return obj
}
func NewAst_GenericTypeInfo(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 *Ast_TypeInfo) *Ast_GenericTypeInfo {
    obj := &Ast_GenericTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_GenericTypeInfo(arg1, arg2, arg3, arg4)
    return obj
}
func (self *Ast_GenericTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_GenericTypeInfo) Get_itemTypeInfoList() *LnsList{ return self.itemTypeInfoList }
func (self *Ast_GenericTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo{ return &self.nilableTypeInfo.Ast_TypeInfo }
func (self *Ast_GenericTypeInfo) Get_genSrcTypeInfo() *Ast_TypeInfo{ return self.genSrcTypeInfo }
func (self *Ast_GenericTypeInfo) AddChildren(arg1 *Ast_TypeInfo) {
self.genSrcTypeInfo. FP.AddChildren( arg1)
}
func (self *Ast_GenericTypeInfo) GetFullName(arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.genSrcTypeInfo. FP.GetFullName( arg1,arg2,arg3)
}
func (self *Ast_GenericTypeInfo) GetOverridingType() LnsAny {
    return self.genSrcTypeInfo. FP.GetOverridingType( )
}
func (self *Ast_GenericTypeInfo) GetParentFullName(arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.genSrcTypeInfo. FP.GetParentFullName( arg1,arg2,arg3)
}
func (self *Ast_GenericTypeInfo) GetParentId() LnsInt {
    return self.genSrcTypeInfo. FP.GetParentId( )
}
func (self *Ast_GenericTypeInfo) getProcessInfo() *Ast_ProcessInfo {
    return self.genSrcTypeInfo. FP.getProcessInfo( )
}
func (self *Ast_GenericTypeInfo) GetTxt(arg1 LnsAny,arg2 LnsAny,arg3 LnsAny) string {
    return self.genSrcTypeInfo. FP.GetTxt( arg1,arg2,arg3)
}
func (self *Ast_GenericTypeInfo) GetTxtWithRaw(arg1 string,arg2 LnsAny,arg3 LnsAny,arg4 LnsAny) string {
    return self.genSrcTypeInfo. FP.GetTxtWithRaw( arg1,arg2,arg3,arg4)
}
func (self *Ast_GenericTypeInfo) Get_abstractFlag() bool {
    return self.genSrcTypeInfo. FP.Get_abstractFlag( )
}
func (self *Ast_GenericTypeInfo) Get_accessMode() LnsInt {
    return self.genSrcTypeInfo. FP.Get_accessMode( )
}
func (self *Ast_GenericTypeInfo) Get_argTypeInfoList() *LnsList {
    return self.genSrcTypeInfo. FP.Get_argTypeInfoList( )
}
func (self *Ast_GenericTypeInfo) Get_autoFlag() bool {
    return self.genSrcTypeInfo. FP.Get_autoFlag( )
}
func (self *Ast_GenericTypeInfo) Get_baseId() LnsInt {
    return self.genSrcTypeInfo. FP.Get_baseId( )
}
func (self *Ast_GenericTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return self.genSrcTypeInfo. FP.Get_baseTypeInfo( )
}
func (self *Ast_GenericTypeInfo) Get_children() *LnsList {
    return self.genSrcTypeInfo. FP.Get_children( )
}
func (self *Ast_GenericTypeInfo) Get_display_stirng() string {
    return self.genSrcTypeInfo. FP.Get_display_stirng( )
}
func (self *Ast_GenericTypeInfo) Get_externalFlag() bool {
    return self.genSrcTypeInfo. FP.Get_externalFlag( )
}
func (self *Ast_GenericTypeInfo) Get_interfaceList() *LnsList {
    return self.genSrcTypeInfo. FP.Get_interfaceList( )
}
func (self *Ast_GenericTypeInfo) Get_kind() LnsInt {
    return self.genSrcTypeInfo. FP.Get_kind( )
}
func (self *Ast_GenericTypeInfo) Get_mutMode() LnsInt {
    return self.genSrcTypeInfo. FP.Get_mutMode( )
}
func (self *Ast_GenericTypeInfo) Get_nilable() bool {
    return self.genSrcTypeInfo. FP.Get_nilable( )
}
func (self *Ast_GenericTypeInfo) Get_nonnilableType() *Ast_TypeInfo {
    return self.genSrcTypeInfo. FP.Get_nonnilableType( )
}
func (self *Ast_GenericTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return self.genSrcTypeInfo. FP.Get_parentInfo( )
}
func (self *Ast_GenericTypeInfo) Get_processInfo() *Ast_ProcessInfo {
    return self.genSrcTypeInfo. FP.Get_processInfo( )
}
func (self *Ast_GenericTypeInfo) Get_rawTxt() string {
    return self.genSrcTypeInfo. FP.Get_rawTxt( )
}
func (self *Ast_GenericTypeInfo) Get_retTypeInfoList() *LnsList {
    return self.genSrcTypeInfo. FP.Get_retTypeInfoList( )
}
func (self *Ast_GenericTypeInfo) Get_scope() LnsAny {
    return self.genSrcTypeInfo. FP.Get_scope( )
}
func (self *Ast_GenericTypeInfo) Get_staticFlag() bool {
    return self.genSrcTypeInfo. FP.Get_staticFlag( )
}
func (self *Ast_GenericTypeInfo) Get_typeData() *Ast_TypeData {
    return self.genSrcTypeInfo. FP.Get_typeData( )
}
func (self *Ast_GenericTypeInfo) HasBase() bool {
    return self.genSrcTypeInfo. FP.HasBase( )
}
func (self *Ast_GenericTypeInfo) HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool {
    return self.genSrcTypeInfo. FP.HasRouteNamespaceFrom( arg1)
}
func (self *Ast_GenericTypeInfo) IsModule() bool {
    return self.genSrcTypeInfo. FP.IsModule( )
}
func (self *Ast_GenericTypeInfo) SerializeTypeInfoList(arg1 string,arg2 *LnsList,arg3 LnsAny) string {
    return self.genSrcTypeInfo. FP.SerializeTypeInfoList( arg1,arg2,arg3)
}
func (self *Ast_GenericTypeInfo) SwitchScope(arg1 *Ast_Scope) {
self.genSrcTypeInfo. FP.SwitchScope( arg1)
}
// 2875: decl @lune.@base.@Ast.GenericTypeInfo.get_display_stirng_with
func (self *Ast_GenericTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    return self.genSrcTypeInfo.FP.Get_display_stirng_with(raw, self.alt2typeMap)
}

// 2881: DeclConstr
func (self *Ast_GenericTypeInfo) InitAst_GenericTypeInfo(processInfo *Ast_ProcessInfo,genSrcTypeInfo *Ast_TypeInfo,itemTypeInfoList *LnsList,moduleTypeInfo *Ast_TypeInfo) {
    self.InitAst_TypeInfo(Ast_TypeInfo_createScope(processInfo, (Lns_unwrap( genSrcTypeInfo.FP.Get_scope()).(*Ast_Scope)).FP.Get_parent(), true, genSrcTypeInfo, nil), processInfo)
    processInfo.FP.Get_idProv().FP.Increment()
    self.typeId = processInfo.FP.Get_idProv().FP.Get_id()
    
    self.moduleTypeInfo = moduleTypeInfo
    
    self.itemTypeInfoList = itemTypeInfoList
    
    self.genSrcTypeInfo = genSrcTypeInfo
    
    if genSrcTypeInfo.FP.Get_itemTypeInfoList().Len() != itemTypeInfoList.Len(){
        Util_err(Lns_getVM().String_format("unmatch generic type number -- %d, %d", []LnsAny{genSrcTypeInfo.FP.Get_itemTypeInfoList().Len(), itemTypeInfoList.Len()}))
    }
    var alt2typeMap *LnsMap
    alt2typeMap = NewLnsMap( map[LnsAny]LnsAny{})
    var workAlt2typeMap *LnsMap
    workAlt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false)
    var hasAlter bool
    hasAlter = false
    for _index, _altTypeInfo := range( genSrcTypeInfo.FP.Get_itemTypeInfoList().Items ) {
        index := _index + 1
        altTypeInfo := _altTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var itemType *Ast_TypeInfo
        itemType = itemTypeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        alt2typeMap.Set(altTypeInfo,itemType)
        if itemType.FP.ApplyGeneric(workAlt2typeMap, moduleTypeInfo) != itemType{
            hasAlter = true
            
        }
    }
    self.hasAlter = hasAlter
    
    self.alt2typeMap = alt2typeMap
    
    processInfo.FP.Get_idProv().FP.Increment()
    self.nilableTypeInfo = NewAst_NilableTypeInfo(nil, processInfo, &self.Ast_TypeInfo, processInfo.FP.Get_idProv().FP.Get_id())
    
}

// 2918: decl @lune.@base.@Ast.GenericTypeInfo.getModule
func (self *Ast_GenericTypeInfo) GetModule() *Ast_TypeInfo {
    return self.moduleTypeInfo
}

// 2923: decl @lune.@base.@Ast.GenericTypeInfo.isInheritFrom
func (self *Ast_GenericTypeInfo) IsInheritFrom(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    var otherSrc *Ast_TypeInfo
    otherSrc = other.FP.Get_genSrcTypeInfo()
    if Lns_op_not(self.genSrcTypeInfo.FP.IsInheritFrom(processInfo, otherSrc, alt2type)){
        return false
    }
    var genOther *Ast_GenericTypeInfo
    
    {
        _genOther := Ast_GenericTypeInfoDownCastF(other.FP)
        if _genOther == nil{
            return true
        } else {
            genOther = _genOther.(*Ast_GenericTypeInfo)
        }
    }
    var workAlt2type *LnsMap
    
    {
        _workAlt2type := alt2type
        if _workAlt2type == nil{
            workAlt2type = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false)
            
        } else {
            workAlt2type = _workAlt2type.(*LnsMap)
        }
    }
    for _, _altType := range( otherSrc.FP.Get_itemTypeInfoList().Items ) {
        altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var genType *Ast_TypeInfo
        
        {
            _genType := self.alt2typeMap.Items[altType]
            if _genType == nil{
                return false
            } else {
                genType = _genType.(*Ast_TypeInfo)
            }
        }
        var otherGenType *Ast_TypeInfo
        otherGenType = Lns_unwrap( genOther.alt2typeMap.Items[altType]).(*Ast_TypeInfo)
        if Lns_op_not(Lns_car(otherGenType.FP.CanEvalWith(processInfo, genType, Ast_CanEvalType__SetEq, workAlt2type)).(bool)){
            return false
        }
    }
    return true
}

// 2963: decl @lune.@base.@Ast.GenericTypeInfo.get_aliasSrc
func (self *Ast_GenericTypeInfo) Get_aliasSrc() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 2967: decl @lune.@base.@Ast.GenericTypeInfo.get_srcTypeInfo
func (self *Ast_GenericTypeInfo) Get_srcTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 2970: decl @lune.@base.@Ast.GenericTypeInfo.get_extedType
func (self *Ast_GenericTypeInfo) Get_extedType() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 2975: decl @lune.@base.@Ast.GenericTypeInfo.canEvalWith
func (self *Ast_GenericTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if other.FP.Get_nilable(){
        return false, nil
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(&self.Ast_TypeInfo)) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(other))) ).(bool)){
        return false, nil
    }
    var otherSrc *Ast_TypeInfo
    otherSrc = other.FP.Get_srcTypeInfo()
    if &self.Ast_TypeInfo == otherSrc{
        return true, nil
    }
    var work *Ast_TypeInfo
    work = otherSrc
    for  {
        if work == Ast_headTypeInfo{
            return false, nil
        }
        for _altType, _genType := range( work.FP.CreateAlt2typeMap(false).Items ) {
            altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            alt2type.Set(altType,genType)
        }
        if self.genSrcTypeInfo.FP.Equals(processInfo, work.FP.Get_genSrcTypeInfo(), alt2type, nil){
            break
        }
        for _, _ifType := range( work.FP.Get_interfaceList().Items ) {
            ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_car(self.FP.CanEvalWith(processInfo, ifType, canEvalType, alt2type)).(bool){
                return true, nil
            }
        }
        work = work.FP.Get_baseTypeInfo()
        
    }
    {
        _otherGen := Ast_GenericTypeInfoDownCastF(work.FP)
        if _otherGen != nil {
            otherGen := _otherGen.(*Ast_GenericTypeInfo)
            var evalType LnsInt
            if canEvalType == Ast_CanEvalType__SetOp{
                evalType = Ast_CanEvalType__SetEq
                
            } else { 
                evalType = canEvalType
                
            }
            for _key, _val := range( self.alt2typeMap.Items ) {
                key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                val := _val.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var otherType *Ast_TypeInfo
                otherType = Ast_AlternateTypeInfo_getAssign(Lns_unwrap( otherGen.alt2typeMap.Items[key]).(*Ast_TypeInfo), alt2type)
                var ret bool
                var mess LnsAny
                ret,mess = val.FP.CanEvalWith(processInfo, otherType, evalType, alt2type)
                if Lns_op_not(ret){
                    return false, mess
                }
            }
        }
    }
    return true, nil
}

// 3036: decl @lune.@base.@Ast.GenericTypeInfo.equals
func (self *Ast_GenericTypeInfo) Equals(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if &self.Ast_TypeInfo == other{
        return true
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.FP.Get_kind() != self.FP.Get_kind()) ||
        Lns_GetEnv().SetStackVal( self.itemTypeInfoList.Len() != other.FP.Get_itemTypeInfoList().Len()) ).(bool){
        return false
    }
    if Lns_op_not((Ast_GenericTypeInfoDownCastF(other.FP))){
        return false
    }
    if Lns_op_not(self.genSrcTypeInfo.FP.Equals(processInfo, other.FP.Get_genSrcTypeInfo(), alt2type, checkModifer)){
        return false
    }
    for _index, _otherItem := range( other.FP.Get_itemTypeInfoList().Items ) {
        index := _index + 1
        otherItem := _otherItem.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var typeInfo *Ast_TypeInfo
        typeInfo = self.itemTypeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(typeInfo.FP.Equals(processInfo, otherItem, alt2type, checkModifer)){
            return false
        }
    }
    return true
}

// 3069: decl @lune.@base.@Ast.GenericTypeInfo.serialize
func (self *Ast_GenericTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    stream.Write(Lns_getVM().String_format("{ skind = %d, typeId = %d, genSrcTypeId = %d, genTypeList = {", []LnsAny{Ast_SerializeKind__Generic, self.typeId, self.genSrcTypeInfo.FP.Get_typeId()}))
    var count LnsInt
    count = 0
    for _, _genType := range( self.alt2typeMap.Items ) {
        genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if count > 0{
            stream.Write(",")
        }
        stream.Write(Lns_getVM().String_format("%d", []LnsAny{genType.FP.Get_typeId()}))
    }
    stream.Write("} }\n")
}

// 3083: decl @lune.@base.@Ast.GenericTypeInfo.createAlt2typeMap
func (self *Ast_GenericTypeInfo) CreateAlt2typeMap(detectFlag bool) *LnsMap {
    var _map *LnsMap
    _map = self.genSrcTypeInfo.FP.CreateAlt2typeMap(detectFlag)
    for _genType, _typeInfo := range( self.alt2typeMap.Items ) {
        genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        _map.Set(genType,typeInfo)
    }
    return _map
}

// 5273: decl @lune.@base.@Ast.GenericTypeInfo.applyGeneric
func (self *Ast_GenericTypeInfo) ApplyGeneric(alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    if self.genSrcTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Class{
        var itemTypeInfoList LnsAny
        var newFlag bool
        itemTypeInfoList,newFlag = Ast_applyGenericList_4604_(self.FP.Get_itemTypeInfoList(), alt2typeMap, moduleTypeInfo)
        if itemTypeInfoList != nil{
            itemTypeInfoList_4431 := itemTypeInfoList.(*LnsList)
            if newFlag{
                return &moduleTypeInfo.FP.getProcessInfo().FP.CreateGeneric(self.genSrcTypeInfo, itemTypeInfoList_4431, moduleTypeInfo).Ast_TypeInfo
            }
        }
    }
    var genSrcTypeInfo LnsAny
    genSrcTypeInfo = self.genSrcTypeInfo.FP.ApplyGeneric(alt2typeMap, moduleTypeInfo)
    if genSrcTypeInfo == self.genSrcTypeInfo{
        return &self.Ast_TypeInfo
    }
    if Lns_op_not(self.hasAlter){
        return &self.Ast_TypeInfo
    }
    Util_errorLog(Lns_getVM().String_format("no support nest generic -- %s", []LnsAny{self.FP.GetTxt(nil, nil, nil)}))
    return nil
}


// declaration Class -- ModifierTypeInfo
type Ast_ModifierTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_ModifierTypeInfo struct {
    Ast_TypeInfo
    srcTypeInfo *Ast_TypeInfo
    typeId LnsInt
    mutMode LnsInt
    FP Ast_ModifierTypeInfoMtd
}
func Ast_ModifierTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_ModifierTypeInfo).FP
}
type Ast_ModifierTypeInfoDownCast interface {
    ToAst_ModifierTypeInfo() *Ast_ModifierTypeInfo
}
func Ast_ModifierTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_ModifierTypeInfoDownCast)
    if ok { return work.ToAst_ModifierTypeInfo() }
    return nil
}
func (obj *Ast_ModifierTypeInfo) ToAst_ModifierTypeInfo() *Ast_ModifierTypeInfo {
    return obj
}
func NewAst_ModifierTypeInfo(arg1 LnsAny, arg2 *Ast_ProcessInfo, arg3 *Ast_TypeInfo, arg4 LnsInt, arg5 LnsInt) *Ast_ModifierTypeInfo {
    obj := &Ast_ModifierTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_ModifierTypeInfo(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Ast_ModifierTypeInfo) InitAst_ModifierTypeInfo(arg1 LnsAny, arg2 *Ast_ProcessInfo, arg3 *Ast_TypeInfo, arg4 LnsInt, arg5 LnsInt) {
    self.Ast_TypeInfo.InitAst_TypeInfo( arg1,arg2)
    self.srcTypeInfo = arg3
    self.typeId = arg4
    self.mutMode = arg5
}
func (self *Ast_ModifierTypeInfo) Get_srcTypeInfo() *Ast_TypeInfo{ return self.srcTypeInfo }
func (self *Ast_ModifierTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_ModifierTypeInfo) Get_mutMode() LnsInt{ return self.mutMode }
func (self *Ast_ModifierTypeInfo) AddChildren(arg1 *Ast_TypeInfo) {
self.srcTypeInfo. FP.AddChildren( arg1)
}
func (self *Ast_ModifierTypeInfo) CreateAlt2typeMap(arg1 bool) *LnsMap {
    return self.srcTypeInfo. FP.CreateAlt2typeMap( arg1)
}
func (self *Ast_ModifierTypeInfo) GetFullName(arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.srcTypeInfo. FP.GetFullName( arg1,arg2,arg3)
}
func (self *Ast_ModifierTypeInfo) GetModule() *Ast_TypeInfo {
    return self.srcTypeInfo. FP.GetModule( )
}
func (self *Ast_ModifierTypeInfo) GetOverridingType() LnsAny {
    return self.srcTypeInfo. FP.GetOverridingType( )
}
func (self *Ast_ModifierTypeInfo) GetParentFullName(arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.srcTypeInfo. FP.GetParentFullName( arg1,arg2,arg3)
}
func (self *Ast_ModifierTypeInfo) GetParentId() LnsInt {
    return self.srcTypeInfo. FP.GetParentId( )
}
func (self *Ast_ModifierTypeInfo) getProcessInfo() *Ast_ProcessInfo {
    return self.srcTypeInfo. FP.getProcessInfo( )
}
func (self *Ast_ModifierTypeInfo) Get_abstractFlag() bool {
    return self.srcTypeInfo. FP.Get_abstractFlag( )
}
func (self *Ast_ModifierTypeInfo) Get_accessMode() LnsInt {
    return self.srcTypeInfo. FP.Get_accessMode( )
}
func (self *Ast_ModifierTypeInfo) Get_aliasSrc() *Ast_TypeInfo {
    return self.srcTypeInfo. FP.Get_aliasSrc( )
}
func (self *Ast_ModifierTypeInfo) Get_argTypeInfoList() *LnsList {
    return self.srcTypeInfo. FP.Get_argTypeInfoList( )
}
func (self *Ast_ModifierTypeInfo) Get_autoFlag() bool {
    return self.srcTypeInfo. FP.Get_autoFlag( )
}
func (self *Ast_ModifierTypeInfo) Get_baseId() LnsInt {
    return self.srcTypeInfo. FP.Get_baseId( )
}
func (self *Ast_ModifierTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return self.srcTypeInfo. FP.Get_baseTypeInfo( )
}
func (self *Ast_ModifierTypeInfo) Get_children() *LnsList {
    return self.srcTypeInfo. FP.Get_children( )
}
func (self *Ast_ModifierTypeInfo) Get_externalFlag() bool {
    return self.srcTypeInfo. FP.Get_externalFlag( )
}
func (self *Ast_ModifierTypeInfo) Get_genSrcTypeInfo() *Ast_TypeInfo {
    return self.srcTypeInfo. FP.Get_genSrcTypeInfo( )
}
func (self *Ast_ModifierTypeInfo) Get_interfaceList() *LnsList {
    return self.srcTypeInfo. FP.Get_interfaceList( )
}
func (self *Ast_ModifierTypeInfo) Get_itemTypeInfoList() *LnsList {
    return self.srcTypeInfo. FP.Get_itemTypeInfoList( )
}
func (self *Ast_ModifierTypeInfo) Get_kind() LnsInt {
    return self.srcTypeInfo. FP.Get_kind( )
}
func (self *Ast_ModifierTypeInfo) Get_nilable() bool {
    return self.srcTypeInfo. FP.Get_nilable( )
}
func (self *Ast_ModifierTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return self.srcTypeInfo. FP.Get_parentInfo( )
}
func (self *Ast_ModifierTypeInfo) Get_processInfo() *Ast_ProcessInfo {
    return self.srcTypeInfo. FP.Get_processInfo( )
}
func (self *Ast_ModifierTypeInfo) Get_rawTxt() string {
    return self.srcTypeInfo. FP.Get_rawTxt( )
}
func (self *Ast_ModifierTypeInfo) Get_retTypeInfoList() *LnsList {
    return self.srcTypeInfo. FP.Get_retTypeInfoList( )
}
func (self *Ast_ModifierTypeInfo) Get_scope() LnsAny {
    return self.srcTypeInfo. FP.Get_scope( )
}
func (self *Ast_ModifierTypeInfo) Get_staticFlag() bool {
    return self.srcTypeInfo. FP.Get_staticFlag( )
}
func (self *Ast_ModifierTypeInfo) Get_typeData() *Ast_TypeData {
    return self.srcTypeInfo. FP.Get_typeData( )
}
func (self *Ast_ModifierTypeInfo) HasBase() bool {
    return self.srcTypeInfo. FP.HasBase( )
}
func (self *Ast_ModifierTypeInfo) HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool {
    return self.srcTypeInfo. FP.HasRouteNamespaceFrom( arg1)
}
func (self *Ast_ModifierTypeInfo) IsInheritFrom(arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.srcTypeInfo. FP.IsInheritFrom( arg1,arg2,arg3)
}
func (self *Ast_ModifierTypeInfo) IsModule() bool {
    return self.srcTypeInfo. FP.IsModule( )
}
func (self *Ast_ModifierTypeInfo) SerializeTypeInfoList(arg1 string,arg2 *LnsList,arg3 LnsAny) string {
    return self.srcTypeInfo. FP.SerializeTypeInfoList( arg1,arg2,arg3)
}
func (self *Ast_ModifierTypeInfo) SwitchScope(arg1 *Ast_Scope) {
self.srcTypeInfo. FP.SwitchScope( arg1)
}
// 3107: decl @lune.@base.@Ast.ModifierTypeInfo.get_extedType
func (self *Ast_ModifierTypeInfo) Get_extedType() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 3112: decl @lune.@base.@Ast.ModifierTypeInfo.getTxt
func (self *Ast_ModifierTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 3117: decl @lune.@base.@Ast.ModifierTypeInfo.getTxtWithRaw
func (self *Ast_ModifierTypeInfo) GetTxtWithRaw(raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    var txt string
    txt = self.srcTypeInfo.FP.GetTxtWithRaw(raw, typeNameCtrl, importInfo, localFlag)
    if Lns_op_not(Ast_isMutable(self.mutMode)){
        txt = "&" + txt
        
    }
    return txt
}

// 3129: decl @lune.@base.@Ast.ModifierTypeInfo.get_display_stirng_with
func (self *Ast_ModifierTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    var txt string
    txt = self.srcTypeInfo.FP.Get_display_stirng_with(raw, alt2type)
    if Ast_isMutable(self.mutMode){
        txt = "mut " + txt
        
    }
    return txt
}

// 3137: decl @lune.@base.@Ast.ModifierTypeInfo.get_display_stirng
func (self *Ast_ModifierTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
}

// 3141: decl @lune.@base.@Ast.ModifierTypeInfo.serialize
func (self *Ast_ModifierTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    stream.Write(Lns_getVM().String_format("{ skind = %d, typeId = %d, srcTypeId = %d, mutMode = %d }\n", []LnsAny{Ast_SerializeKind__Modifier, self.typeId, self.srcTypeInfo.FP.Get_typeId(), self.mutMode}))
}

// 3148: decl @lune.@base.@Ast.ModifierTypeInfo.canEvalWith
func (self *Ast_ModifierTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    var evalType LnsInt
    if self.srcTypeInfo.FP.Get_itemTypeInfoList().Len() >= 1{
        if _switch11019 := canEvalType; _switch11019 == Ast_CanEvalType__SetEq || _switch11019 == Ast_CanEvalType__SetOp {
            evalType = Ast_CanEvalType__SetOpIMut
            
        } else {
            evalType = canEvalType
            
        }
    } else { 
        evalType = canEvalType
        
    }
    return Ast_TypeInfo_canEvalWithBase(processInfo, self.srcTypeInfo, Ast_TypeInfo_isMut(&self.Ast_TypeInfo), other.FP.Get_srcTypeInfo(), evalType, alt2type)
}

// 3178: decl @lune.@base.@Ast.ModifierTypeInfo.equals
func (self *Ast_ModifierTypeInfo) Equals(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if Lns_isCondTrue( checkModifer){
        if Ast_TypeInfo_isMut(&self.Ast_TypeInfo) != Ast_TypeInfo_isMut(typeInfo){
            return false
        }
    }
    return self.srcTypeInfo.FP.Equals(processInfo, typeInfo.FP.Get_srcTypeInfo(), alt2type, checkModifer)
}

// 4665: decl @lune.@base.@Ast.ModifierTypeInfo.get_nonnilableType
func (self *Ast_ModifierTypeInfo) Get_nonnilableType() *Ast_TypeInfo {
    var orgType *Ast_TypeInfo
    orgType = self.srcTypeInfo.FP.Get_nonnilableType()
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(&self.Ast_TypeInfo)) ||
        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(orgType))) ).(bool){
        return orgType
    }
    return self.FP.getProcessInfo().FP.CreateModifier(orgType, Ast_MutMode__IMut)
}

// 4673: decl @lune.@base.@Ast.ModifierTypeInfo.get_nilableTypeInfo
func (self *Ast_ModifierTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo {
    var orgType *Ast_TypeInfo
    orgType = self.srcTypeInfo.FP.Get_nilableTypeInfo()
    if Lns_op_not(Ast_TypeInfo_isMut(orgType)){
        return orgType
    }
    return self.FP.getProcessInfo().FP.CreateModifier(orgType, Ast_MutMode__IMut)
}

// 6583: decl @lune.@base.@Ast.ModifierTypeInfo.applyGeneric
func (self *Ast_ModifierTypeInfo) ApplyGeneric(alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.srcTypeInfo.FP.ApplyGeneric(alt2typeMap, moduleTypeInfo)
    if typeInfo == self.srcTypeInfo{
        return &self.Ast_TypeInfo
    }
    if typeInfo != nil{
        typeInfo_5250 := typeInfo.(*Ast_TypeInfo)
        return moduleTypeInfo.FP.getProcessInfo().FP.CreateModifier(typeInfo_5250, Ast_MutMode__IMut)
    }
    return nil
}


// declaration Class -- ModuleTypeInfo
type Ast_ModuleTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_mutable() bool
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_ModuleTypeInfo struct {
    Ast_TypeInfo
    externalFlag bool
    parentInfo *Ast_TypeInfo
    typeId LnsInt
    rawTxt string
    mutable bool
    FP Ast_ModuleTypeInfoMtd
}
func Ast_ModuleTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_ModuleTypeInfo).FP
}
type Ast_ModuleTypeInfoDownCast interface {
    ToAst_ModuleTypeInfo() *Ast_ModuleTypeInfo
}
func Ast_ModuleTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_ModuleTypeInfoDownCast)
    if ok { return work.ToAst_ModuleTypeInfo() }
    return nil
}
func (obj *Ast_ModuleTypeInfo) ToAst_ModuleTypeInfo() *Ast_ModuleTypeInfo {
    return obj
}
func NewAst_ModuleTypeInfo(arg1 *Ast_ProcessInfo, arg2 *Ast_Scope, arg3 bool, arg4 string, arg5 LnsAny, arg6 LnsInt, arg7 bool) *Ast_ModuleTypeInfo {
    obj := &Ast_ModuleTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_ModuleTypeInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Ast_ModuleTypeInfo) Get_externalFlag() bool{ return self.externalFlag }
func (self *Ast_ModuleTypeInfo) Get_parentInfo() *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_ModuleTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_ModuleTypeInfo) Get_rawTxt() string{ return self.rawTxt }
func (self *Ast_ModuleTypeInfo) Get_mutable() bool{ return self.mutable }
// 3208: DeclConstr
func (self *Ast_ModuleTypeInfo) InitAst_ModuleTypeInfo(processInfo *Ast_ProcessInfo,scope *Ast_Scope,externalFlag bool,txt string,parentInfo LnsAny,typeId LnsInt,mutable bool) {
    self.InitAst_TypeInfo(scope, processInfo)
    self.externalFlag = externalFlag
    
    self.rawTxt = txt
    
    self.parentInfo = Lns_unwrapDefault( parentInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.typeId = typeId
    
    self.mutable = mutable
    
    {
        __exp := parentInfo
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            _exp.FP.AddChildren(&self.Ast_TypeInfo)
        }
    }
    processInfo.FP.Get_idProv().FP.Increment()
    scope.FP.Set_ownerTypeInfo(&self.Ast_TypeInfo)
}

// 3230: decl @lune.@base.@Ast.ModuleTypeInfo.get_baseTypeInfo
func (self *Ast_ModuleTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 3234: decl @lune.@base.@Ast.ModuleTypeInfo.isModule
func (self *Ast_ModuleTypeInfo) IsModule() bool {
    return true
}

// 3238: decl @lune.@base.@Ast.ModuleTypeInfo.get_accessMode
func (self *Ast_ModuleTypeInfo) Get_accessMode() LnsInt {
    return Ast_AccessMode__Pub
}

// 3242: decl @lune.@base.@Ast.ModuleTypeInfo.get_kind
func (self *Ast_ModuleTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Module
}

// 3246: decl @lune.@base.@Ast.ModuleTypeInfo.getParentId
func (self *Ast_ModuleTypeInfo) GetParentId() LnsInt {
    return self.parentInfo.FP.Get_typeId()
}

// 3251: decl @lune.@base.@Ast.ModuleTypeInfo.getTxt
func (self *Ast_ModuleTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 3257: decl @lune.@base.@Ast.ModuleTypeInfo.getTxtWithRaw
func (self *Ast_ModuleTypeInfo) GetTxtWithRaw(rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return rawTxt
}

// 3264: decl @lune.@base.@Ast.ModuleTypeInfo.get_display_stirng_with
func (self *Ast_ModuleTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(raw, nil, nil, nil)
}

// 3268: decl @lune.@base.@Ast.ModuleTypeInfo.get_display_stirng
func (self *Ast_ModuleTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
}

// 3271: decl @lune.@base.@Ast.ModuleTypeInfo.canEvalWith
func (self *Ast_ModuleTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return false, nil
}

// 3278: decl @lune.@base.@Ast.ModuleTypeInfo.serialize
func (self *Ast_ModuleTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    var txt string
    txt = Lns_getVM().String_format("{ skind = %d, parentId = %d, typeId = %d, txt = '%s', ", []LnsAny{Ast_SerializeKind__Module, self.FP.GetParentId(), self.typeId, self.rawTxt})
    stream.Write(txt + "\n")
    stream.Write("children = {")
    var set *LnsMap
    
    {
        _set := validChildrenSet
        if _set == nil{
            set = NewLnsMap( map[LnsAny]LnsAny{})
            
        } else {
            set = _set.(*LnsMap)
        }
    }
    if Lns_isCondTrue( validChildrenSet){
        for _, _child := range( self.FP.Get_children().Items ) {
            child := _child.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( set.Items[child.FP.Get_typeId()]) &&
                Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( child.FP.Get_accessMode() == Ast_AccessMode__Pub) ||
                    Lns_GetEnv().SetStackVal( child.FP.Get_accessMode() == Ast_AccessMode__Global) ).(bool))) )){
                stream.Write(Lns_getVM().String_format("%d, ", []LnsAny{child.FP.Get_typeId()}))
            }
        }
    }
    stream.Write("} }\n")
}


// declaration Class -- EnumValInfo
type Ast_EnumValInfoMtd interface {
    Get_name() string
    Get_val() LnsAny
}
type Ast_EnumValInfo struct {
    name string
    val LnsAny
    FP Ast_EnumValInfoMtd
}
func Ast_EnumValInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_EnumValInfo).FP
}
type Ast_EnumValInfoDownCast interface {
    ToAst_EnumValInfo() *Ast_EnumValInfo
}
func Ast_EnumValInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_EnumValInfoDownCast)
    if ok { return work.ToAst_EnumValInfo() }
    return nil
}
func (obj *Ast_EnumValInfo) ToAst_EnumValInfo() *Ast_EnumValInfo {
    return obj
}
func NewAst_EnumValInfo(arg1 string, arg2 LnsAny) *Ast_EnumValInfo {
    obj := &Ast_EnumValInfo{}
    obj.FP = obj
    obj.InitAst_EnumValInfo(arg1, arg2)
    return obj
}
func (self *Ast_EnumValInfo) InitAst_EnumValInfo(arg1 string, arg2 LnsAny) {
    self.name = arg1
    self.val = arg2
}
func (self *Ast_EnumValInfo) Get_name() string{ return self.name }
func (self *Ast_EnumValInfo) Get_val() LnsAny{ return self.val }

// declaration Class -- EnumTypeInfo
type Ast_EnumTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    AddEnumValInfo(arg1 *Ast_EnumValInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetEnumValInfo(arg1 string) LnsAny
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_name2EnumValInfo() *LnsMap
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    Get_val2EnumValInfo() *LnsMap
    Get_valTypeInfo() *Ast_TypeInfo
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_EnumTypeInfo struct {
    Ast_TypeInfo
    externalFlag bool
    parentInfo *Ast_TypeInfo
    typeId LnsInt
    rawTxt string
    accessMode LnsInt
    nilableTypeInfo *Ast_TypeInfo
    valTypeInfo *Ast_TypeInfo
    name2EnumValInfo *LnsMap
    val2EnumValInfo *LnsMap
    FP Ast_EnumTypeInfoMtd
}
func Ast_EnumTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_EnumTypeInfo).FP
}
type Ast_EnumTypeInfoDownCast interface {
    ToAst_EnumTypeInfo() *Ast_EnumTypeInfo
}
func Ast_EnumTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_EnumTypeInfoDownCast)
    if ok { return work.ToAst_EnumTypeInfo() }
    return nil
}
func (obj *Ast_EnumTypeInfo) ToAst_EnumTypeInfo() *Ast_EnumTypeInfo {
    return obj
}
func NewAst_EnumTypeInfo(arg1 *Ast_ProcessInfo, arg2 *Ast_Scope, arg3 bool, arg4 LnsInt, arg5 string, arg6 LnsAny, arg7 LnsInt, arg8 *Ast_TypeInfo) *Ast_EnumTypeInfo {
    obj := &Ast_EnumTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_EnumTypeInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Ast_EnumTypeInfo) Get_externalFlag() bool{ return self.externalFlag }
func (self *Ast_EnumTypeInfo) Get_parentInfo() *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_EnumTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_EnumTypeInfo) Get_rawTxt() string{ return self.rawTxt }
func (self *Ast_EnumTypeInfo) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Ast_EnumTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo{ return self.nilableTypeInfo }
func (self *Ast_EnumTypeInfo) Get_valTypeInfo() *Ast_TypeInfo{ return self.valTypeInfo }
func (self *Ast_EnumTypeInfo) Get_name2EnumValInfo() *LnsMap{ return self.name2EnumValInfo }
func (self *Ast_EnumTypeInfo) Get_val2EnumValInfo() *LnsMap{ return self.val2EnumValInfo }
// 3349: DeclConstr
func (self *Ast_EnumTypeInfo) InitAst_EnumTypeInfo(processInfo *Ast_ProcessInfo,scope *Ast_Scope,externalFlag bool,accessMode LnsInt,txt string,parentInfo LnsAny,typeId LnsInt,valTypeInfo *Ast_TypeInfo) {
    self.InitAst_TypeInfo(scope, processInfo)
    self.externalFlag = externalFlag
    
    self.accessMode = accessMode
    
    self.rawTxt = txt
    
    self.parentInfo = Lns_unwrapDefault( parentInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.typeId = typeId
    
    self.name2EnumValInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.valTypeInfo = valTypeInfo
    
    self.val2EnumValInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    {
        __exp := parentInfo
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            _exp.FP.AddChildren(&self.Ast_TypeInfo)
        }
    }
    self.nilableTypeInfo = &NewAst_NilableTypeInfo(nil, processInfo, &self.Ast_TypeInfo, typeId + 1).Ast_TypeInfo
    
    processInfo.FP.Get_idProv().FP.Increment()
    scope.FP.Set_ownerTypeInfo(&self.Ast_TypeInfo)
}

// 3376: decl @lune.@base.@Ast.EnumTypeInfo.isModule
func (self *Ast_EnumTypeInfo) IsModule() bool {
    return false
}

// 3380: decl @lune.@base.@Ast.EnumTypeInfo.get_kind
func (self *Ast_EnumTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Enum
}

// 3385: decl @lune.@base.@Ast.EnumTypeInfo.get_baseTypeInfo
func (self *Ast_EnumTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 3388: decl @lune.@base.@Ast.EnumTypeInfo.getParentId
func (self *Ast_EnumTypeInfo) GetParentId() LnsInt {
    return self.parentInfo.FP.Get_typeId()
}

// 3393: decl @lune.@base.@Ast.EnumTypeInfo.getTxt
func (self *Ast_EnumTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 3399: decl @lune.@base.@Ast.EnumTypeInfo.getTxtWithRaw
func (self *Ast_EnumTypeInfo) GetTxtWithRaw(rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return rawTxt
}

// 3406: decl @lune.@base.@Ast.EnumTypeInfo.get_display_stirng_with
func (self *Ast_EnumTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(raw, nil, nil, nil)
}

// 3410: decl @lune.@base.@Ast.EnumTypeInfo.get_display_stirng
func (self *Ast_EnumTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
}

// 3414: decl @lune.@base.@Ast.EnumTypeInfo.canEvalWith
func (self *Ast_EnumTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return &self.Ast_TypeInfo == other.FP.Get_srcTypeInfo().FP.Get_aliasSrc(), nil
}

// 3421: decl @lune.@base.@Ast.EnumTypeInfo.addEnumValInfo
func (self *Ast_EnumTypeInfo) AddEnumValInfo(valInfo *Ast_EnumValInfo) {
    self.name2EnumValInfo.Set(valInfo.FP.Get_name(),valInfo)
    self.val2EnumValInfo.Set(Ast_getEnumLiteralVal(valInfo.FP.Get_val()),valInfo)
}

// 3426: decl @lune.@base.@Ast.EnumTypeInfo.getEnumValInfo
func (self *Ast_EnumTypeInfo) GetEnumValInfo(name string) LnsAny {
    return self.name2EnumValInfo.Items[name]
}

// 3430: decl @lune.@base.@Ast.EnumTypeInfo.get_mutMode
func (self *Ast_EnumTypeInfo) Get_mutMode() LnsInt {
    return Ast_MutMode__Mut
}

// 5682: decl @lune.@base.@Ast.EnumTypeInfo.serialize
func (self *Ast_EnumTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    var txt string
    txt = Lns_getVM().String_format("{ skind = %d, parentId = %d, typeId = %d, txt = '%s',\naccessMode = %d, kind = %d, valTypeId = %d, ", []LnsAny{Ast_SerializeKind__Enum, self.FP.GetParentId(), self.typeId, self.rawTxt, self.accessMode, Ast_TypeInfoKind__Enum, self.valTypeInfo.FP.Get_typeId()})
    stream.Write(txt)
    stream.Write("enumValList = {")
    {
        __collection23123 := self.name2EnumValInfo
        __sorted23123 := __collection23123.CreateKeyListStr()
        __sorted23123.Sort( LnsItemKindStr, nil )
        for _, ___key23123 := range( __sorted23123.Items ) {
            enumValInfo := __collection23123.Items[ ___key23123 ].(Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
            stream.Write(Lns_getVM().String_format("%s = ", []LnsAny{enumValInfo.FP.Get_name()}))
            switch _exp23121 := enumValInfo.FP.Get_val().(type) {
            case *Ast_EnumLiteral__Int:
            val := _exp23121.Val1
                stream.Write(Lns_getVM().String_format("%d,", []LnsAny{val}))
            case *Ast_EnumLiteral__Real:
            val := _exp23121.Val1
                stream.Write(Lns_getVM().String_format("%g,", []LnsAny{val}))
            case *Ast_EnumLiteral__Str:
            val := _exp23121.Val1
                stream.Write(Lns_getVM().String_format("%q,", []LnsAny{val}))
            }
        }
    }
    stream.Write("} }\n")
}


// declaration Class -- AlgeTypeInfo
type Ast_AlgeTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    AddValInfo(arg1 *Ast_AlgeValInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    GetValInfo(arg1 string) LnsAny
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    Get_valInfoMap() *LnsMap
    Get_valInfoNum() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_AlgeTypeInfo struct {
    Ast_TypeInfo
    externalFlag bool
    parentInfo *Ast_TypeInfo
    typeId LnsInt
    rawTxt string
    accessMode LnsInt
    nilableTypeInfo *Ast_TypeInfo
    valInfoMap *LnsMap
    valInfoNum LnsInt
    FP Ast_AlgeTypeInfoMtd
}
func Ast_AlgeTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_AlgeTypeInfo).FP
}
type Ast_AlgeTypeInfoDownCast interface {
    ToAst_AlgeTypeInfo() *Ast_AlgeTypeInfo
}
func Ast_AlgeTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_AlgeTypeInfoDownCast)
    if ok { return work.ToAst_AlgeTypeInfo() }
    return nil
}
func (obj *Ast_AlgeTypeInfo) ToAst_AlgeTypeInfo() *Ast_AlgeTypeInfo {
    return obj
}
func NewAst_AlgeTypeInfo(arg1 *Ast_ProcessInfo, arg2 *Ast_Scope, arg3 bool, arg4 LnsInt, arg5 string, arg6 LnsAny, arg7 LnsInt) *Ast_AlgeTypeInfo {
    obj := &Ast_AlgeTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AlgeTypeInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Ast_AlgeTypeInfo) Get_externalFlag() bool{ return self.externalFlag }
func (self *Ast_AlgeTypeInfo) Get_parentInfo() *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_AlgeTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_AlgeTypeInfo) Get_rawTxt() string{ return self.rawTxt }
func (self *Ast_AlgeTypeInfo) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Ast_AlgeTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo{ return self.nilableTypeInfo }
func (self *Ast_AlgeTypeInfo) Get_valInfoMap() *LnsMap{ return self.valInfoMap }
func (self *Ast_AlgeTypeInfo) Get_valInfoNum() LnsInt{ return self.valInfoNum }
// 3456: decl @lune.@base.@Ast.AlgeTypeInfo.get_baseTypeInfo
func (self *Ast_AlgeTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 3460: DeclConstr
func (self *Ast_AlgeTypeInfo) InitAst_AlgeTypeInfo(processInfo *Ast_ProcessInfo,scope *Ast_Scope,externalFlag bool,accessMode LnsInt,txt string,parentInfo LnsAny,typeId LnsInt) {
    self.InitAst_TypeInfo(scope, processInfo)
    self.externalFlag = externalFlag
    
    self.accessMode = accessMode
    
    self.rawTxt = txt
    
    self.parentInfo = Lns_unwrapDefault( parentInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.typeId = typeId
    
    self.valInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.valInfoNum = 0
    
    {
        __exp := parentInfo
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            _exp.FP.AddChildren(&self.Ast_TypeInfo)
        }
    }
    self.nilableTypeInfo = &NewAst_NilableTypeInfo(nil, processInfo, &self.Ast_TypeInfo, typeId + 1).Ast_TypeInfo
    
    processInfo.FP.Get_idProv().FP.Increment()
    scope.FP.Set_ownerTypeInfo(&self.Ast_TypeInfo)
}

// 3484: decl @lune.@base.@Ast.AlgeTypeInfo.getValInfo
func (self *Ast_AlgeTypeInfo) GetValInfo(name string) LnsAny {
    return self.valInfoMap.Items[name]
}

// 3488: decl @lune.@base.@Ast.AlgeTypeInfo.isModule
func (self *Ast_AlgeTypeInfo) IsModule() bool {
    return false
}

// 3492: decl @lune.@base.@Ast.AlgeTypeInfo.get_kind
func (self *Ast_AlgeTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Alge
}

// 3496: decl @lune.@base.@Ast.AlgeTypeInfo.getParentId
func (self *Ast_AlgeTypeInfo) GetParentId() LnsInt {
    return self.parentInfo.FP.Get_typeId()
}

// 3501: decl @lune.@base.@Ast.AlgeTypeInfo.getTxt
func (self *Ast_AlgeTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 3507: decl @lune.@base.@Ast.AlgeTypeInfo.getTxtWithRaw
func (self *Ast_AlgeTypeInfo) GetTxtWithRaw(rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return rawTxt
}

// 3514: decl @lune.@base.@Ast.AlgeTypeInfo.get_display_stirng_with
func (self *Ast_AlgeTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(raw, nil, nil, nil)
}

// 3518: decl @lune.@base.@Ast.AlgeTypeInfo.get_display_stirng
func (self *Ast_AlgeTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
}

// 3522: decl @lune.@base.@Ast.AlgeTypeInfo.canEvalWith
func (self *Ast_AlgeTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return &self.Ast_TypeInfo == other.FP.Get_srcTypeInfo().FP.Get_aliasSrc(), nil
}

// 3529: decl @lune.@base.@Ast.AlgeTypeInfo.get_mutMode
func (self *Ast_AlgeTypeInfo) Get_mutMode() LnsInt {
    return Ast_MutMode__Mut
}

// 3551: decl @lune.@base.@Ast.AlgeTypeInfo.addValInfo
func (self *Ast_AlgeTypeInfo) AddValInfo(valInfo *Ast_AlgeValInfo) {
    self.valInfoMap.Set(valInfo.FP.Get_name(),valInfo)
    self.valInfoNum = self.valInfoNum + 1
    
}

// 5734: decl @lune.@base.@Ast.AlgeTypeInfo.serialize
func (self *Ast_AlgeTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    var txt string
    txt = Lns_getVM().String_format("{ skind = %d, parentId = %d, typeId = %d, txt = '%s',\naccessMode = %d, kind = %d, ", []LnsAny{Ast_SerializeKind__Alge, self.FP.GetParentId(), self.typeId, self.rawTxt, self.accessMode, Ast_TypeInfoKind__Alge})
    stream.Write(txt)
    stream.Write("algeValList = {")
    var firstFlag bool
    firstFlag = true
    {
        __collection23364 := self.valInfoMap
        __sorted23364 := __collection23364.CreateKeyListStr()
        __sorted23364.Sort( LnsItemKindStr, nil )
        for _, ___key23364 := range( __sorted23364.Items ) {
            algeValInfo := __collection23364.Items[ ___key23364 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            if Lns_op_not(firstFlag){
                stream.Write(",")
            } else { 
                firstFlag = false
                
            }
            algeValInfo.FP.Serialize(stream)
        }
    }
    stream.Write("} }\n")
}


// declaration Class -- AlgeValInfo
type Ast_AlgeValInfoMtd interface {
    Get_algeTpye() *Ast_AlgeTypeInfo
    Get_name() string
    Get_typeList() *LnsList
    Serialize(arg1 Lns_oStream)
}
type Ast_AlgeValInfo struct {
    name string
    typeList *LnsList
    algeTpye *Ast_AlgeTypeInfo
    FP Ast_AlgeValInfoMtd
}
func Ast_AlgeValInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_AlgeValInfo).FP
}
type Ast_AlgeValInfoDownCast interface {
    ToAst_AlgeValInfo() *Ast_AlgeValInfo
}
func Ast_AlgeValInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_AlgeValInfoDownCast)
    if ok { return work.ToAst_AlgeValInfo() }
    return nil
}
func (obj *Ast_AlgeValInfo) ToAst_AlgeValInfo() *Ast_AlgeValInfo {
    return obj
}
func NewAst_AlgeValInfo(arg1 string, arg2 *LnsList, arg3 *Ast_AlgeTypeInfo) *Ast_AlgeValInfo {
    obj := &Ast_AlgeValInfo{}
    obj.FP = obj
    obj.InitAst_AlgeValInfo(arg1, arg2, arg3)
    return obj
}
func (self *Ast_AlgeValInfo) InitAst_AlgeValInfo(arg1 string, arg2 *LnsList, arg3 *Ast_AlgeTypeInfo) {
    self.name = arg1
    self.typeList = arg2
    self.algeTpye = arg3
}
func (self *Ast_AlgeValInfo) Get_name() string{ return self.name }
func (self *Ast_AlgeValInfo) Get_typeList() *LnsList{ return self.typeList }
func (self *Ast_AlgeValInfo) Get_algeTpye() *Ast_AlgeTypeInfo{ return self.algeTpye }
// 3539: decl @lune.@base.@Ast.AlgeValInfo.serialize
func (self *Ast_AlgeValInfo) Serialize(stream Lns_oStream) {
    stream.Write(Lns_getVM().String_format("{ name = '%s', typeList = {", []LnsAny{self.name}))
    for _index, _typeInfo := range( self.typeList.Items ) {
        index := _index + 1
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            stream.Write(", ")
        }
        stream.Write(Lns_getVM().String_format("%d", []LnsAny{typeInfo.FP.Get_typeId()}))
    }
    stream.Write("} }")
}


// declaration Class -- NormalTypeInfo
type Ast_NormalTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    EqualsSub(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    Set_mutMode(arg1 LnsInt)
    SwitchScope(arg1 *Ast_Scope)
    SwitchScopeTo(arg1 *Ast_Scope)
}
type Ast_NormalTypeInfo struct {
    Ast_TypeInfo
    externalFlag bool
    itemTypeInfoList *LnsList
    argTypeInfoList *LnsList
    retTypeInfoList *LnsList
    parentInfo *Ast_TypeInfo
    typeId LnsInt
    rawTxt string
    kind LnsInt
    staticFlag bool
    accessMode LnsInt
    autoFlag bool
    abstractFlag bool
    baseTypeInfo *Ast_TypeInfo
    interfaceList *LnsList
    nilableTypeInfo *Ast_TypeInfo
    mutMode LnsInt
    alt2typeMap *LnsMap
    overridingType LnsAny
    FP Ast_NormalTypeInfoMtd
}
func Ast_NormalTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_NormalTypeInfo).FP
}
type Ast_NormalTypeInfoDownCast interface {
    ToAst_NormalTypeInfo() *Ast_NormalTypeInfo
}
func Ast_NormalTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_NormalTypeInfoDownCast)
    if ok { return work.ToAst_NormalTypeInfo() }
    return nil
}
func (obj *Ast_NormalTypeInfo) ToAst_NormalTypeInfo() *Ast_NormalTypeInfo {
    return obj
}
func NewAst_NormalTypeInfo(arg1 *Ast_ProcessInfo, arg2 bool, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 bool, arg7 bool, arg8 bool, arg9 LnsInt, arg10 string, arg11 LnsAny, arg12 LnsInt, arg13 LnsInt, arg14 LnsAny, arg15 LnsAny, arg16 LnsAny, arg17 LnsAny) *Ast_NormalTypeInfo {
    obj := &Ast_NormalTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_NormalTypeInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17)
    return obj
}
func (self *Ast_NormalTypeInfo) Get_externalFlag() bool{ return self.externalFlag }
func (self *Ast_NormalTypeInfo) Get_itemTypeInfoList() *LnsList{ return self.itemTypeInfoList }
func (self *Ast_NormalTypeInfo) Get_argTypeInfoList() *LnsList{ return self.argTypeInfoList }
func (self *Ast_NormalTypeInfo) Get_retTypeInfoList() *LnsList{ return self.retTypeInfoList }
func (self *Ast_NormalTypeInfo) Get_parentInfo() *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_NormalTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_NormalTypeInfo) Get_rawTxt() string{ return self.rawTxt }
func (self *Ast_NormalTypeInfo) Get_kind() LnsInt{ return self.kind }
func (self *Ast_NormalTypeInfo) Get_staticFlag() bool{ return self.staticFlag }
func (self *Ast_NormalTypeInfo) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Ast_NormalTypeInfo) Get_autoFlag() bool{ return self.autoFlag }
func (self *Ast_NormalTypeInfo) Get_abstractFlag() bool{ return self.abstractFlag }
func (self *Ast_NormalTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo{ return self.baseTypeInfo }
func (self *Ast_NormalTypeInfo) Get_interfaceList() *LnsList{ return self.interfaceList }
func (self *Ast_NormalTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo{ return self.nilableTypeInfo }
func (self *Ast_NormalTypeInfo) Get_mutMode() LnsInt{ return self.mutMode }
func (self *Ast_NormalTypeInfo) Set_mutMode(arg1 LnsInt){ self.mutMode = arg1 }
// 3606: decl @lune.@base.@Ast.NormalTypeInfo.getOverridingType
func (self *Ast_NormalTypeInfo) GetOverridingType() LnsAny {
    switch _exp12689 := self.overridingType.(type) {
    case *Ast_OverridingType__NotOverride:
        return nil
    case *Ast_OverridingType__Override:
    typeInfo := _exp12689.Val1
        return typeInfo
    case *Ast_OverridingType__NoReady:
        var scope *Ast_Scope
        scope = Lns_unwrap( self.parentInfo.FP.Get_scope()).(*Ast_Scope)
        {
            _typeInfo := scope.FP.GetTypeInfoField(self.rawTxt, false, scope, Ast_ScopeAccess__Normal)
            if _typeInfo != nil {
                typeInfo := _typeInfo.(*Ast_TypeInfo)
                {
                    _workType := typeInfo.FP.GetOverridingType()
                    if _workType != nil {
                        workType := _workType.(*Ast_TypeInfo)
                        self.overridingType = &Ast_OverridingType__Override{workType}
                        
                        return workType
                    } else {
                        self.overridingType = &Ast_OverridingType__Override{typeInfo}
                        
                        return typeInfo
                    }
                }
            } else {
                self.overridingType = Ast_OverridingType__NotOverride_Obj
                
                return nil
            }
        }
    }
// insert a dummy
    return nil
}

// 3634: decl @lune.@base.@Ast.NormalTypeInfo.switchScopeTo
func (self *Ast_NormalTypeInfo) SwitchScopeTo(scope *Ast_Scope) {
    self.FP.SwitchScope(scope)
}

// 3638: DeclConstr
func (self *Ast_NormalTypeInfo) InitAst_NormalTypeInfo(processInfo *Ast_ProcessInfo,abstractFlag bool,scope LnsAny,baseTypeInfo LnsAny,interfaceList LnsAny,autoFlag bool,externalFlag bool,staticFlag bool,accessMode LnsInt,txt string,parentInfo LnsAny,typeId LnsInt,kind LnsInt,itemTypeInfoList LnsAny,argTypeInfoList LnsAny,retTypeInfoList LnsAny,mutMode LnsAny) {
    self.InitAst_TypeInfo(scope, processInfo)
    if Lns_type(kind) != "number"{
        Util_printStackTrace()
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( kind == Ast_TypeInfoKind__Method) &&
        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentInfo) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.HasBase()})/* 3653:30 */)) )){
        self.overridingType = Ast_OverridingType__NoReady_Obj
        
    } else { 
        self.overridingType = Ast_OverridingType__NotOverride_Obj
        
    }
    self.abstractFlag = abstractFlag
    
    self.baseTypeInfo = Lns_unwrapDefault( baseTypeInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.interfaceList = Lns_unwrapDefault( interfaceList, NewLnsList([]LnsAny{})).(*LnsList)
    
    self.autoFlag = autoFlag
    
    self.externalFlag = externalFlag
    
    self.staticFlag = staticFlag
    
    self.accessMode = accessMode
    
    self.rawTxt = txt
    
    self.kind = kind
    
    self.itemTypeInfoList = Lns_unwrapDefault( itemTypeInfoList, NewLnsList([]LnsAny{})).(*LnsList)
    
    self.argTypeInfoList = Lns_unwrapDefault( argTypeInfoList, NewLnsList([]LnsAny{})).(*LnsList)
    
    self.retTypeInfoList = Lns_unwrapDefault( retTypeInfoList, NewLnsList([]LnsAny{})).(*LnsList)
    
    self.parentInfo = Lns_unwrapDefault( parentInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.mutMode = Lns_unwrapDefault( mutMode, Ast_MutMode__IMut).(LnsInt)
    
    var setupAlt2typeMap func() *LnsMap
    setupAlt2typeMap = func() *LnsMap {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( self.baseTypeInfo == Ast_headTypeInfo) &&
            Lns_GetEnv().SetStackVal( self.interfaceList.Len() == 0) ).(bool)){
            return NewLnsMap( map[LnsAny]LnsAny{})
        }
        var alt2typeMap *LnsMap
        alt2typeMap = NewLnsMap( map[LnsAny]LnsAny{})
        if _switch13122 := kind; _switch13122 == Ast_TypeInfoKind__Set || _switch13122 == Ast_TypeInfoKind__Map || _switch13122 == Ast_TypeInfoKind__List || _switch13122 == Ast_TypeInfoKind__Array || _switch13122 == Ast_TypeInfoKind__Box {
            if self.itemTypeInfoList.Len() != self.baseTypeInfo.FP.Get_itemTypeInfoList().Len(){
                Util_err(Lns_getVM().String_format("unmatch generic type number -- %d, %d", []LnsAny{self.itemTypeInfoList.Len(), self.baseTypeInfo.FP.Get_itemTypeInfoList().Len()}))
            }
            for _index, _appyType := range( self.itemTypeInfoList.Items ) {
                index := _index + 1
                appyType := _appyType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var genType *Ast_TypeInfo
                genType = self.baseTypeInfo.FP.Get_itemTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                alt2typeMap.Set(genType,appyType)
            }
        } else if _switch13122 == Ast_TypeInfoKind__Class || _switch13122 == Ast_TypeInfoKind__IF {
            for _, _ifType := range( self.interfaceList.Items ) {
                ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                {
                    _genericType := Ast_GenericTypeInfoDownCastF(ifType.FP)
                    if _genericType != nil {
                        genericType := _genericType.(*Ast_GenericTypeInfo)
                        for _altType, _genType := range( genericType.FP.CreateAlt2typeMap(false).Items ) {
                            altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                            genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                            alt2typeMap.Set(altType,genType)
                        }
                    }
                }
            }
        }
        return alt2typeMap
    }
    self.alt2typeMap = setupAlt2typeMap()
    
    self.typeId = typeId
    
    if kind == Ast_TypeInfoKind__Root{
    } else { 
        if parentInfo != nil{
            parentInfo_3405 := parentInfo.(*Ast_TypeInfo)
            parentInfo_3405.FP.AddChildren(&self.Ast_TypeInfo)
        }
        var hasNilable bool
        hasNilable = false
        if _switch13224 := (kind); _switch13224 == Ast_TypeInfoKind__Prim || _switch13224 == Ast_TypeInfoKind__List || _switch13224 == Ast_TypeInfoKind__Array || _switch13224 == Ast_TypeInfoKind__Set || _switch13224 == Ast_TypeInfoKind__Map || _switch13224 == Ast_TypeInfoKind__Class || _switch13224 == Ast_TypeInfoKind__Stem || _switch13224 == Ast_TypeInfoKind__Module || _switch13224 == Ast_TypeInfoKind__IF {
            hasNilable = true
            
        } else if _switch13224 == Ast_TypeInfoKind__Func || _switch13224 == Ast_TypeInfoKind__Method || _switch13224 == Ast_TypeInfoKind__Form || _switch13224 == Ast_TypeInfoKind__FormFunc {
            hasNilable = true
            
        }
        if hasNilable{
            self.nilableTypeInfo = &NewAst_NilableTypeInfo(nil, processInfo, &self.Ast_TypeInfo, typeId + 1).Ast_TypeInfo
            
            processInfo.FP.Get_idProv().FP.Increment()
        } else { 
            self.nilableTypeInfo = Ast_headTypeInfo
            
        }
        processInfo.FP.Get_idProv().FP.Increment()
    }
}

// 3742: decl @lune.@base.@Ast.NormalTypeInfo.createAlt2typeMap
func (self *Ast_NormalTypeInfo) CreateAlt2typeMap(detectFlag bool) *LnsMap {
    var _map *LnsMap
    _map = self.baseTypeInfo.FP.CreateAlt2typeMap(detectFlag)
    for _genType, _typeInfo := range( self.alt2typeMap.Items ) {
        genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        _map.Set(genType,typeInfo)
    }
    return _map
}

// 3752: decl @lune.@base.@Ast.NormalTypeInfo.get_nilable
func (self *Ast_NormalTypeInfo) Get_nilable() bool {
    return false
}

// 3756: decl @lune.@base.@Ast.NormalTypeInfo.isModule
func (self *Ast_NormalTypeInfo) IsModule() bool {
    return false
}

// 3760: decl @lune.@base.@Ast.NormalTypeInfo.getParentId
func (self *Ast_NormalTypeInfo) GetParentId() LnsInt {
    return self.parentInfo.FP.Get_typeId()
}

// 3764: decl @lune.@base.@Ast.NormalTypeInfo.get_baseId
func (self *Ast_NormalTypeInfo) Get_baseId() LnsInt {
    return self.baseTypeInfo.FP.Get_typeId()
}

// 3769: decl @lune.@base.@Ast.NormalTypeInfo.getTxt
func (self *Ast_NormalTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 3775: decl @lune.@base.@Ast.NormalTypeInfo.getTxtWithRaw
func (self *Ast_NormalTypeInfo) GetTxtWithRaw(raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    var parentTxt string
    parentTxt = ""
    if typeNameCtrl != nil{
        typeNameCtrl_3460 := typeNameCtrl.(*Ast_TypeNameCtrl)
        parentTxt = self.FP.GetParentFullName(typeNameCtrl_3460, importInfo, localFlag)
        
    }
    var name string
    if self.itemTypeInfoList.Len() > 0{
        var txt string
        txt = raw + "<"
        for _index, _typeInfo := range( self.itemTypeInfoList.Items ) {
            index := _index + 1
            typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                txt = txt + ","
                
            }
            txt = txt + typeInfo.FP.GetTxt(typeNameCtrl, importInfo, localFlag)
            
        }
        name = parentTxt + txt + ">"
        
    } else { 
        name = parentTxt + raw
        
    }
    return name
}

// 3803: decl @lune.@base.@Ast.NormalTypeInfo.get_display_stirng_with
func (self *Ast_NormalTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    if _switch13661 := self.kind; _switch13661 == Ast_TypeInfoKind__Func || _switch13661 == Ast_TypeInfoKind__Form || _switch13661 == Ast_TypeInfoKind__FormFunc || _switch13661 == Ast_TypeInfoKind__Method || _switch13661 == Ast_TypeInfoKind__Macro {
        var txt string
        txt = raw + "("
        for _index, _argType := range( self.argTypeInfoList.Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                txt = txt + ", "
                
            }
            txt = txt + argType.FP.Get_display_stirng()
            
        }
        txt = txt + ")"
        
        for _index, _retType := range( self.retTypeInfoList.Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index == 1{
                txt = txt + ": "
                
            } else { 
                txt = txt + ", "
                
            }
            txt = txt + retType.FP.Get_display_stirng()
            
        }
        return txt
    }
    var parentTxt string
    parentTxt = ""
    var name string
    if self.itemTypeInfoList.Len() > 0{
        var txt string
        txt = raw + "<"
        for _index, _typeInfo := range( self.itemTypeInfoList.Items ) {
            index := _index + 1
            typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                txt = txt + ","
                
            }
            txt = txt + typeInfo.FP.Get_display_stirng_with(typeInfo.FP.Get_rawTxt(), alt2type)
            
        }
        name = parentTxt + txt + ">"
        
    } else { 
        name = parentTxt + raw
        
    }
    return name
}

// 3848: decl @lune.@base.@Ast.NormalTypeInfo.get_display_stirng
func (self *Ast_NormalTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
}

// 3853: decl @lune.@base.@Ast.NormalTypeInfo.serialize
func (self *Ast_NormalTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    if self.typeId == Ast_rootTypeId{
        return 
    }
    var parentId LnsInt
    parentId = self.FP.GetParentId()
    var txt string
    txt = Lns_getVM().String_format("{ skind=%d, parentId = %d, typeId = %d, baseId = %d, txt = '%s',\n        abstractFlag = %s, staticFlag = %s, accessMode = %d, kind = %d, mutMode = %d, ", []LnsAny{Ast_SerializeKind__Normal, parentId, self.typeId, self.FP.Get_baseId(), self.rawTxt, self.abstractFlag, self.staticFlag, self.accessMode, self.kind, self.mutMode})
    var children *LnsList
    children = NewLnsList([]LnsAny{})
    var set *LnsMap
    
    {
        _set := validChildrenSet
        if _set == nil{
            set = NewLnsMap( map[LnsAny]LnsAny{})
            
        } else {
            set = _set.(*LnsMap)
        }
    }
    for _, _child := range( self.FP.Get_children().Items ) {
        child := _child.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_isCondTrue( set.Items[child.FP.Get_typeId()]){
            children.Insert(Ast_TypeInfo2Stem(child))
        }
    }
    stream.Write(txt + self.FP.SerializeTypeInfoList("itemTypeId = {", self.itemTypeInfoList, nil) + self.FP.SerializeTypeInfoList("ifList = {", self.interfaceList, nil) + self.FP.SerializeTypeInfoList("argTypeId = {", self.argTypeInfoList, nil) + self.FP.SerializeTypeInfoList("retTypeId = {", self.retTypeInfoList, nil) + self.FP.SerializeTypeInfoList("children = {", children, true) + "}\n")
}

// 3884: decl @lune.@base.@Ast.NormalTypeInfo.equalsSub
func (self *Ast_NormalTypeInfo) EqualsSub(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if self.typeId == typeInfo.FP.Get_typeId(){
        return true
    }
    if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
        return typeInfo.FP.Equals(processInfo, &self.Ast_TypeInfo, alt2type, checkModifer)
    }
    {
        _aliasType := Ast_AliasTypeInfoDownCastF(typeInfo.FP)
        if _aliasType != nil {
            aliasType := _aliasType.(*Ast_AliasTypeInfo)
            return aliasType.FP.Equals(processInfo, &self.Ast_TypeInfo, alt2type, checkModifer)
        }
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.kind != typeInfo.FP.Get_kind()) ||
        Lns_GetEnv().SetStackVal( self.staticFlag != typeInfo.FP.Get_staticFlag()) ||
        Lns_GetEnv().SetStackVal( self.autoFlag != typeInfo.FP.Get_autoFlag()) ||
        Lns_GetEnv().SetStackVal( self.FP.Get_nilable() != typeInfo.FP.Get_nilable()) ||
        Lns_GetEnv().SetStackVal( self.rawTxt != typeInfo.FP.Get_rawTxt()) ||
        Lns_GetEnv().SetStackVal( self.parentInfo != typeInfo.FP.Get_parentInfo()) ||
        Lns_GetEnv().SetStackVal( self.baseTypeInfo != typeInfo.FP.Get_baseTypeInfo()) ).(bool){
        return false
    }
    if self.accessMode != typeInfo.FP.Get_accessMode(){
        if _switch14135 := self.kind; _switch14135 == Ast_TypeInfoKind__List || _switch14135 == Ast_TypeInfoKind__Map || _switch14135 == Ast_TypeInfoKind__Array || _switch14135 == Ast_TypeInfoKind__Set {
        } else {
            return false
        }
    }
    {
        if self.itemTypeInfoList.Len() != typeInfo.FP.Get_itemTypeInfoList().Len(){
            return false
        }
        for _index, _item := range( self.itemTypeInfoList.Items ) {
            index := _index + 1
            item := _item.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(item.FP.Equals(processInfo, typeInfo.FP.Get_itemTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type, checkModifer)){
                return false
            }
        }
    }
    {
        if self.retTypeInfoList.Len() != typeInfo.FP.Get_retTypeInfoList().Len(){
            return false
        }
        for _index, _item := range( self.retTypeInfoList.Items ) {
            index := _index + 1
            item := _item.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(item.FP.Equals(processInfo, typeInfo.FP.Get_retTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type, checkModifer)){
                return false
            }
        }
    }
    return true
}

// 3954: decl @lune.@base.@Ast.NormalTypeInfo.equals
func (self *Ast_NormalTypeInfo) Equals(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    return self.FP.EqualsSub(processInfo, typeInfo, alt2type, checkModifer)
}

// 3962: decl @lune.@base.@Ast.NormalTypeInfo.create
func Ast_NormalTypeInfo_create(processInfo *Ast_ProcessInfo,accessMode LnsInt,abstractFlag bool,scope LnsAny,baseInfo *Ast_TypeInfo,interfaceList *LnsList,parentInfo *Ast_TypeInfo,staticFlag bool,kind LnsInt,txt string,itemTypeInfo *LnsList,argTypeInfoList *LnsList,retTypeInfoList *LnsList,mutMode LnsAny) *Ast_TypeInfo {
    if kind == Ast_TypeInfoKind__Prim{
        {
            __exp := Ast_sym2builtInTypeMap.Items[txt]
            if __exp != nil {
                _exp := __exp.(*Ast_SymbolInfo)
                return _exp.FP.Get_typeInfo()
            }
        }
        Util_err(Lns_getVM().String_format("not found symbol -- %s", []LnsAny{txt}))
    }
    processInfo.FP.Get_idProv().FP.Increment()
    var info *Ast_NormalTypeInfo
    info = NewAst_NormalTypeInfo(processInfo, abstractFlag, scope, baseInfo, interfaceList, false, true, staticFlag, accessMode, txt, parentInfo, processInfo.FP.Get_idProv().FP.Get_id(), kind, itemTypeInfo, argTypeInfoList, retTypeInfoList, mutMode)
    return &info.Ast_TypeInfo
}

// 4180: decl @lune.@base.@Ast.NormalTypeInfo.createBuiltin
func Ast_NormalTypeInfo_createBuiltin(idName string,typeTxt string,kind LnsInt,typeDDD LnsAny,ifList LnsAny) *Ast_TypeInfo {
    var typeId LnsInt
    typeId = Ast_rootProcessInfo.FP.Get_idProv().FP.GetNewId()
    if kind == Ast_TypeInfoKind__Root{
        typeId = Ast_rootTypeId
        
    }
    var argTypeList *LnsList
    argTypeList = NewLnsList([]LnsAny{})
    var retTypeList *LnsList
    retTypeList = NewLnsList([]LnsAny{})
    if typeTxt == "form"{
        {
            __exp := typeDDD
            if __exp != nil {
                _exp := __exp.(*Ast_TypeInfo)
                argTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(_exp)})
                
                retTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(_exp)})
                
            }
        }
    }
    var scope LnsAny
    scope = nil
    if _switch15552 := kind; _switch15552 == Ast_TypeInfoKind__Array || _switch15552 == Ast_TypeInfoKind__List || _switch15552 == Ast_TypeInfoKind__Set || _switch15552 == Ast_TypeInfoKind__Class || _switch15552 == Ast_TypeInfoKind__Module || _switch15552 == Ast_TypeInfoKind__IF || _switch15552 == Ast_TypeInfoKind__Form || _switch15552 == Ast_TypeInfoKind__FormFunc || _switch15552 == Ast_TypeInfoKind__Func || _switch15552 == Ast_TypeInfoKind__Method || _switch15552 == Ast_TypeInfoKind__Macro {
        scope = NewAst_Scope(Ast_rootProcessInfo, Ast_rootScope, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( kind == Ast_TypeInfoKind__Class) ||
            Lns_GetEnv().SetStackVal( kind == Ast_TypeInfoKind__Module) ||
            Lns_GetEnv().SetStackVal( kind == Ast_TypeInfoKind__IF) ||
            Lns_GetEnv().SetStackVal( kind == Ast_TypeInfoKind__List) ||
            Lns_GetEnv().SetStackVal( kind == Ast_TypeInfoKind__Array) ||
            Lns_GetEnv().SetStackVal( kind == Ast_TypeInfoKind__Set) ).(bool), nil, nil)
        
    }
    var genTypeList *LnsList
    genTypeList = NewLnsList([]LnsAny{})
    if _switch15650 := kind; _switch15650 == Ast_TypeInfoKind__Array || _switch15650 == Ast_TypeInfoKind__List || _switch15650 == Ast_TypeInfoKind__Set {
        genTypeList.Insert(Ast_AlternateTypeInfo2Stem(Ast_rootProcessInfo.FP.CreateAlternate(true, 1, "T", Ast_AccessMode__Pri, Ast_headTypeInfo, nil, nil)))
    } else if _switch15650 == Ast_TypeInfoKind__Map {
        genTypeList.Insert(Ast_AlternateTypeInfo2Stem(Ast_rootProcessInfo.FP.CreateAlternate(true, 1, "K", Ast_AccessMode__Pri, Ast_headTypeInfo, nil, nil)))
        genTypeList.Insert(Ast_AlternateTypeInfo2Stem(Ast_rootProcessInfo.FP.CreateAlternate(true, 2, "V", Ast_AccessMode__Pri, Ast_headTypeInfo, nil, nil)))
    }
    var info *Ast_NormalTypeInfo
    info = NewAst_NormalTypeInfo(Ast_rootProcessInfo, false, scope, nil, ifList, false, false, false, Ast_AccessMode__Pub, typeTxt, Ast_headTypeInfo, typeId, kind, genTypeList, argTypeList, retTypeList, Ast_MutMode__Mut)
    Ast_registBuiltin_4166_(idName, typeTxt, kind, &info.Ast_TypeInfo, Ast_headTypeInfo, Ast_Scope2Stem(scope) != nil)
    return &info.Ast_TypeInfo
}

// 5861: decl @lune.@base.@Ast.NormalTypeInfo.isAvailableMapping
func Ast_NormalTypeInfo_isAvailableMapping(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,checkedTypeMap *LnsMap) bool {
    var isAvailableMappingSub func() bool
    isAvailableMappingSub = func() bool {
        if _switch24023 := typeInfo.FP.Get_kind(); _switch24023 == Ast_TypeInfoKind__Prim || _switch24023 == Ast_TypeInfoKind__Enum {
            return true
        } else if _switch24023 == Ast_TypeInfoKind__Alge {
            var algeTypeInfo *Ast_AlgeTypeInfo
            algeTypeInfo = Lns_unwrap( (Ast_AlgeTypeInfoDownCastF(typeInfo.FP))).(*Ast_AlgeTypeInfo)
            for _, _valInfo := range( algeTypeInfo.FP.Get_valInfoMap().Items ) {
                valInfo := _valInfo.(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
                for _, _paramType := range( valInfo.FP.Get_typeList().Items ) {
                    paramType := _paramType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if Lns_op_not(Ast_NormalTypeInfo_isAvailableMapping(processInfo, paramType, checkedTypeMap)){
                        return false
                    }
                }
            }
            return true
        } else if _switch24023 == Ast_TypeInfoKind__Stem {
            return true
        } else if _switch24023 == Ast_TypeInfoKind__Class || _switch24023 == Ast_TypeInfoKind__IF {
            if typeInfo.FP.Equals(processInfo, Ast_builtinTypeString, nil, nil){
                return true
            }
            return typeInfo.FP.IsInheritFrom(processInfo, Ast_builtinTypeMapping, nil)
        } else if _switch24023 == Ast_TypeInfoKind__Alternate {
            return typeInfo.FP.IsInheritFrom(processInfo, Ast_builtinTypeMapping, nil)
        } else if _switch24023 == Ast_TypeInfoKind__List || _switch24023 == Ast_TypeInfoKind__Array || _switch24023 == Ast_TypeInfoKind__Set {
            return Ast_NormalTypeInfo_isAvailableMapping(processInfo, typeInfo.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), checkedTypeMap)
        } else if _switch24023 == Ast_TypeInfoKind__Map {
            if Ast_NormalTypeInfo_isAvailableMapping(processInfo, typeInfo.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), checkedTypeMap){
                var keyType *Ast_TypeInfo
                keyType = typeInfo.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( keyType.FP.Equals(processInfo, Ast_builtinTypeString, nil, nil)) ||
                    Lns_GetEnv().SetStackVal( keyType.FP.Get_kind() == Ast_TypeInfoKind__Prim) ||
                    Lns_GetEnv().SetStackVal( keyType.FP.Get_kind() == Ast_TypeInfoKind__Enum) ).(bool){
                    return true
                }
            }
            return false
        } else if _switch24023 == Ast_TypeInfoKind__Nilable {
            return Ast_NormalTypeInfo_isAvailableMapping(processInfo, typeInfo.FP.Get_nonnilableType(), checkedTypeMap)
        } else {
            return false
        }
    // insert a dummy
        return false
    }
    typeInfo = typeInfo.FP.Get_srcTypeInfo()
    
    {
        __exp := checkedTypeMap.Items[typeInfo]
        if __exp != nil {
            _exp := __exp.(bool)
            return _exp
        }
    }
    checkedTypeMap.Set(typeInfo,true)
    var result bool
    result = isAvailableMappingSub()
    checkedTypeMap.Set(typeInfo,result)
    return result
}

// 5935: decl @lune.@base.@Ast.NormalTypeInfo.isInheritFrom
func (self *Ast_NormalTypeInfo) IsInheritFrom(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    if self.FP.Get_typeId() == other.FP.Get_typeId(){
        return true
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( self.FP.Get_kind() != Ast_TypeInfoKind__Class) &&
            Lns_GetEnv().SetStackVal( self.FP.Get_kind() != Ast_TypeInfoKind__IF) ).(bool))) ||
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( other.FP.Get_kind() != Ast_TypeInfoKind__Class) &&
            Lns_GetEnv().SetStackVal( other.FP.Get_kind() != Ast_TypeInfoKind__IF) ).(bool))) ).(bool){
        if other == Ast_builtinTypeMapping{
            return Ast_NormalTypeInfo_isAvailableMapping(processInfo, &self.Ast_TypeInfo, NewLnsMap( map[LnsAny]LnsAny{}))
        }
        return false
    }
    return Ast_TypeInfo_isInherit(processInfo, &self.Ast_TypeInfo, other, alt2type)
}

// 6573: decl @lune.@base.@Ast.NormalTypeInfo.canEvalWith
func (self *Ast_NormalTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return Ast_TypeInfo_canEvalWithBase(processInfo, &self.Ast_TypeInfo, Ast_TypeInfo_isMut(&self.Ast_TypeInfo), other, canEvalType, alt2type)
}

// 6597: decl @lune.@base.@Ast.NormalTypeInfo.applyGeneric
func (self *Ast_NormalTypeInfo) ApplyGeneric(alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var itemTypeInfoList *LnsList
    var needNew bool
    
    {
        _itemTypeInfoList, _needNew := Ast_applyGenericList_4604_(self.itemTypeInfoList, alt2typeMap, moduleTypeInfo)
        if _itemTypeInfoList == nil{
            return nil
        } else {
            itemTypeInfoList = _itemTypeInfoList.(*LnsList)
            needNew = _needNew
        }
    }
    var processInfo *Ast_ProcessInfo
    processInfo = moduleTypeInfo.FP.getProcessInfo()
    if _switch28222 := self.FP.Get_kind(); _switch28222 == Ast_TypeInfoKind__Set {
        if Lns_op_not(needNew){
            return &self.Ast_TypeInfo
        }
        return processInfo.FP.CreateSet(self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode)
    } else if _switch28222 == Ast_TypeInfoKind__List {
        if Lns_op_not(needNew){
            return &self.Ast_TypeInfo
        }
        return processInfo.FP.CreateList(self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode)
    } else if _switch28222 == Ast_TypeInfoKind__Array {
        if Lns_op_not(needNew){
            return &self.Ast_TypeInfo
        }
        return processInfo.FP.CreateArray(self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode)
    } else if _switch28222 == Ast_TypeInfoKind__Map {
        if Lns_op_not(needNew){
            return &self.Ast_TypeInfo
        }
        return processInfo.FP.CreateMap(self.accessMode, self.parentInfo, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), itemTypeInfoList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), self.mutMode)
    } else if _switch28222 == Ast_TypeInfoKind__Func || _switch28222 == Ast_TypeInfoKind__Form || _switch28222 == Ast_TypeInfoKind__FormFunc {
        var argTypeInfoList *LnsList
        var workArg bool
        
        {
            _argTypeInfoList, _workArg := Ast_applyGenericList_4604_(self.argTypeInfoList, alt2typeMap, moduleTypeInfo)
            if _argTypeInfoList == nil{
                return nil
            } else {
                argTypeInfoList = _argTypeInfoList.(*LnsList)
                workArg = _workArg
            }
        }
        var retTypeInfoList *LnsList
        var workRet bool
        
        {
            _retTypeInfoList, _workRet := Ast_applyGenericList_4604_(self.retTypeInfoList, alt2typeMap, moduleTypeInfo)
            if _retTypeInfoList == nil{
                return nil
            } else {
                retTypeInfoList = _retTypeInfoList.(*LnsList)
                workRet = _workRet
            }
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( needNew) ||
            Lns_GetEnv().SetStackVal( workArg) ||
            Lns_GetEnv().SetStackVal( workRet) ).(bool){
            return processInfo.FP.CreateFunc(self.abstractFlag, false, Ast_getScope(&self.Ast_TypeInfo), self.kind, self.parentInfo, self.autoFlag, self.externalFlag, self.staticFlag, self.accessMode, self.rawTxt, itemTypeInfoList, argTypeInfoList, retTypeInfoList, Ast_TypeInfo_isMut(&self.Ast_TypeInfo))
        }
        return &self.Ast_TypeInfo
    } else {
        if self.itemTypeInfoList.Len() == 0{
            return &self.Ast_TypeInfo
        }
        return nil
    }
// insert a dummy
    return nil
}


// declaration Class -- TypeInfo2Map
type Ast_TypeInfo2MapMtd interface {
    Clone() *Ast_TypeInfo2Map
    Get_BoxMap() *LnsMap
    Get_DDDMap() *LnsMap
    Get_ExtDDDMap() *LnsMap
    Get_ExtMap() *LnsMap
    Get_ImutModifierMap() *LnsMap
    Get_MutModifierMap() *LnsMap
}
type Ast_TypeInfo2Map struct {
    ImutModifierMap *LnsMap
    MutModifierMap *LnsMap
    BoxMap *LnsMap
    DDDMap *LnsMap
    ExtDDDMap *LnsMap
    ExtMap *LnsMap
    FP Ast_TypeInfo2MapMtd
}
func Ast_TypeInfo2Map2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_TypeInfo2Map).FP
}
type Ast_TypeInfo2MapDownCast interface {
    ToAst_TypeInfo2Map() *Ast_TypeInfo2Map
}
func Ast_TypeInfo2MapDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_TypeInfo2MapDownCast)
    if ok { return work.ToAst_TypeInfo2Map() }
    return nil
}
func (obj *Ast_TypeInfo2Map) ToAst_TypeInfo2Map() *Ast_TypeInfo2Map {
    return obj
}
func NewAst_TypeInfo2Map() *Ast_TypeInfo2Map {
    obj := &Ast_TypeInfo2Map{}
    obj.FP = obj
    obj.InitAst_TypeInfo2Map()
    return obj
}
func (self *Ast_TypeInfo2Map) Get_ImutModifierMap() *LnsMap{ return self.ImutModifierMap }
func (self *Ast_TypeInfo2Map) Get_MutModifierMap() *LnsMap{ return self.MutModifierMap }
func (self *Ast_TypeInfo2Map) Get_BoxMap() *LnsMap{ return self.BoxMap }
func (self *Ast_TypeInfo2Map) Get_DDDMap() *LnsMap{ return self.DDDMap }
func (self *Ast_TypeInfo2Map) Get_ExtDDDMap() *LnsMap{ return self.ExtDDDMap }
func (self *Ast_TypeInfo2Map) Get_ExtMap() *LnsMap{ return self.ExtMap }
// 4006: DeclConstr
func (self *Ast_TypeInfo2Map) InitAst_TypeInfo2Map() {
    self.ImutModifierMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.MutModifierMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.BoxMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.DDDMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.ExtDDDMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.ExtMap = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 4015: decl @lune.@base.@Ast.TypeInfo2Map.clone
func (self *Ast_TypeInfo2Map) Clone() *Ast_TypeInfo2Map {
    var obj *Ast_TypeInfo2Map
    obj = NewAst_TypeInfo2Map()
    for _key, _val := range( self.ImutModifierMap.Items ) {
        key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        val := _val.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        obj.ImutModifierMap.Set(key,val)
    }
    
    for _key, _val := range( self.MutModifierMap.Items ) {
        key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        val := _val.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        obj.MutModifierMap.Set(key,val)
    }
    
    for _key, _val := range( self.BoxMap.Items ) {
        key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        val := _val.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        obj.BoxMap.Set(key,val)
    }
    
    for _key, _val := range( self.DDDMap.Items ) {
        key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        val := _val.(Ast_DDDTypeInfoDownCast).ToAst_DDDTypeInfo()
        obj.DDDMap.Set(key,val)
    }
    
    for _key, _val := range( self.ExtDDDMap.Items ) {
        key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        val := _val.(Ast_DDDTypeInfoDownCast).ToAst_DDDTypeInfo()
        obj.ExtDDDMap.Set(key,val)
    }
    
    for _key, _val := range( self.ExtMap.Items ) {
        key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        val := _val.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        obj.ExtMap.Set(key,val)
    }
    
    return obj
}


// declaration Class -- DDDTypeInfo
type Ast_DDDTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extTypeFlag() bool
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    Get_typeInfo() *Ast_TypeInfo
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_DDDTypeInfo struct {
    Ast_TypeInfo
    typeInfo *Ast_TypeInfo
    typeId LnsInt
    externalFlag bool
    itemTypeInfoList *LnsList
    extedType *Ast_TypeInfo
    FP Ast_DDDTypeInfoMtd
}
func Ast_DDDTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_DDDTypeInfo).FP
}
type Ast_DDDTypeInfoDownCast interface {
    ToAst_DDDTypeInfo() *Ast_DDDTypeInfo
}
func Ast_DDDTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_DDDTypeInfoDownCast)
    if ok { return work.ToAst_DDDTypeInfo() }
    return nil
}
func (obj *Ast_DDDTypeInfo) ToAst_DDDTypeInfo() *Ast_DDDTypeInfo {
    return obj
}
func NewAst_DDDTypeInfo(arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 *Ast_TypeInfo, arg4 bool, arg5 LnsAny) *Ast_DDDTypeInfo {
    obj := &Ast_DDDTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_DDDTypeInfo(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Ast_DDDTypeInfo) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }
func (self *Ast_DDDTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_DDDTypeInfo) Get_externalFlag() bool{ return self.externalFlag }
func (self *Ast_DDDTypeInfo) Get_itemTypeInfoList() *LnsList{ return self.itemTypeInfoList }
func (self *Ast_DDDTypeInfo) Get_extedType() *Ast_TypeInfo{ return self.extedType }
// 4738: decl @lune.@base.@Ast.DDDTypeInfo.get_extTypeFlag
func (self *Ast_DDDTypeInfo) Get_extTypeFlag() bool {
    return self.extedType != &self.Ast_TypeInfo
}

// 4742: decl @lune.@base.@Ast.DDDTypeInfo.get_scope
func (self *Ast_DDDTypeInfo) Get_scope() LnsAny {
    return nil
}

// 4747: decl @lune.@base.@Ast.DDDTypeInfo.get_baseTypeInfo
func (self *Ast_DDDTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 4751: decl @lune.@base.@Ast.DDDTypeInfo.get_parentInfo
func (self *Ast_DDDTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 4755: DeclConstr
func (self *Ast_DDDTypeInfo) InitAst_DDDTypeInfo(processInfo *Ast_ProcessInfo,typeId LnsInt,typeInfo *Ast_TypeInfo,externalFlag bool,extOrgDDType LnsAny) {
    self.InitAst_TypeInfo(nil, processInfo)
    self.typeId = typeId
    
    self.typeInfo = typeInfo
    
    self.externalFlag = externalFlag
    
    self.itemTypeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.typeInfo)})
    
    var extOrgType *Ast_DDDTypeInfo
    if extOrgDDType != nil{
        extOrgDDType_4111 := extOrgDDType.(*Ast_DDDTypeInfo)
        extOrgType = extOrgDDType_4111
        
        processInfo.FP.get_typeInfo2Map().ExtDDDMap.Set(typeInfo,self)
    } else {
        extOrgType = self
        
        processInfo.FP.get_typeInfo2Map().DDDMap.Set(typeInfo,self)
    }
    self.extedType = &extOrgType.Ast_TypeInfo
    
}

// 4775: decl @lune.@base.@Ast.DDDTypeInfo.isModule
func (self *Ast_DDDTypeInfo) IsModule() bool {
    return false
}

// 4779: decl @lune.@base.@Ast.DDDTypeInfo.canEvalWith
func (self *Ast_DDDTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return self.typeInfo.FP.CanEvalWith(processInfo, other, canEvalType, alt2type)
}

// 4786: decl @lune.@base.@Ast.DDDTypeInfo.serialize
func (self *Ast_DDDTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    stream.Write(Lns_getVM().String_format("{ skind=%d, typeId = %d, itemTypeId = %d, parentId = %d, extTypeFlag = %s }\n", []LnsAny{Ast_SerializeKind__DDD, self.typeId, self.typeInfo.FP.Get_typeId(), Ast_headTypeInfo.FP.Get_typeId(), self.FP.Get_extTypeFlag()}))
}

// 4793: decl @lune.@base.@Ast.DDDTypeInfo.get_display_stirng_with
func (self *Ast_DDDTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    var txt string
    txt = self.FP.GetTxtWithRaw(raw, nil, nil, nil)
    return txt
}

// 4799: decl @lune.@base.@Ast.DDDTypeInfo.get_display_stirng
func (self *Ast_DDDTypeInfo) Get_display_stirng() string {
    var txt string
    txt = self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
    return txt
}

// 4804: decl @lune.@base.@Ast.DDDTypeInfo.getModule
func (self *Ast_DDDTypeInfo) GetModule() *Ast_TypeInfo {
    return self.typeInfo.FP.GetModule()
}

// 4807: decl @lune.@base.@Ast.DDDTypeInfo.get_rawTxt
func (self *Ast_DDDTypeInfo) Get_rawTxt() string {
    return self.FP.GetTxt(nil, nil, nil)
}

// 4810: decl @lune.@base.@Ast.DDDTypeInfo.get_kind
func (self *Ast_DDDTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__DDD
}

// 4813: decl @lune.@base.@Ast.DDDTypeInfo.get_nilable
func (self *Ast_DDDTypeInfo) Get_nilable() bool {
    return true
}

// 4817: decl @lune.@base.@Ast.DDDTypeInfo.get_nilableTypeInfo
func (self *Ast_DDDTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 4820: decl @lune.@base.@Ast.DDDTypeInfo.get_mutMode
func (self *Ast_DDDTypeInfo) Get_mutMode() LnsInt {
    return self.typeInfo.FP.Get_mutMode()
}

// 4823: decl @lune.@base.@Ast.DDDTypeInfo.get_aliasSrc
func (self *Ast_DDDTypeInfo) Get_aliasSrc() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 4826: decl @lune.@base.@Ast.DDDTypeInfo.get_srcTypeInfo
func (self *Ast_DDDTypeInfo) Get_srcTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 4829: decl @lune.@base.@Ast.DDDTypeInfo.get_accessMode
func (self *Ast_DDDTypeInfo) Get_accessMode() LnsInt {
    return Ast_AccessMode__Pub
}

// 5209: decl @lune.@base.@Ast.DDDTypeInfo.getTxt
func (self *Ast_DDDTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw("...", typeNameCtrl, importInfo, localFlag)
}

// 5215: decl @lune.@base.@Ast.DDDTypeInfo.getTxtWithRaw
func (self *Ast_DDDTypeInfo) GetTxtWithRaw(raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    if self.typeInfo == Ast_builtinTypeStem_{
        if self.FP.Get_extTypeFlag(){
            return "Luaval<...>"
        }
        return "..."
    }
    var typeInfo *Ast_TypeInfo
    if self.FP.Get_extTypeFlag(){
        typeInfo = self.typeInfo.FP.Get_extedType()
        
    } else { 
        typeInfo = self.typeInfo
        
    }
    var txt string
    txt = Lns_getVM().String_format("...<%s>", []LnsAny{typeInfo.FP.GetTxt(typeNameCtrl, importInfo, localFlag)})
    if self.FP.Get_extTypeFlag(){
        return Lns_getVM().String_format("Luaval<%s>", []LnsAny{txt})
    }
    return txt
}


// declaration Class -- CombineType
type Ast_CombineTypeMtd interface {
    andIfSet(arg1 *Ast_ProcessInfo, arg2 *LnsSet, arg3 *LnsMap)
    AndType(arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *LnsMap) LnsAny
    CreateStem(arg1 *Ast_ProcessInfo) *Ast_TypeInfo
    Get_typeInfo(arg1 *Ast_ProcessInfo) *Ast_TypeInfo
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
}
type Ast_CombineType struct {
    ifSet *LnsSet
    nilable bool
    mutMode LnsInt
    FP Ast_CombineTypeMtd
}
func Ast_CombineType2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_CombineType).FP
}
type Ast_CombineTypeDownCast interface {
    ToAst_CombineType() *Ast_CombineType
}
func Ast_CombineTypeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_CombineTypeDownCast)
    if ok { return work.ToAst_CombineType() }
    return nil
}
func (obj *Ast_CombineType) ToAst_CombineType() *Ast_CombineType {
    return obj
}
func NewAst_CombineType(arg1 *Ast_TypeInfo) *Ast_CombineType {
    obj := &Ast_CombineType{}
    obj.FP = obj
    obj.InitAst_CombineType(arg1)
    return obj
}
// 4921: DeclConstr
func (self *Ast_CombineType) InitAst_CombineType(typeInfo *Ast_TypeInfo) {
    self.ifSet = NewLnsSet([]LnsAny{})
    
    for _, _iftype := range( typeInfo.FP.Get_interfaceList().Items ) {
        iftype := _iftype.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        self.ifSet.Add(Ast_TypeInfo2Stem(iftype))
    }
    self.nilable = typeInfo.FP.Get_nilable()
    
    self.mutMode = typeInfo.FP.Get_mutMode()
    
}

// 4930: decl @lune.@base.@Ast.CombineType.isInheritFrom
func (self *Ast_CombineType) IsInheritFrom(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    for _ifType := range( self.ifSet.Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if ifType.FP.IsInheritFrom(processInfo, other, alt2type){
            return true
        }
    }
    return false
}

// 4942: decl @lune.@base.@Ast.CombineType.andIfSet
func (self *Ast_CombineType) andIfSet(processInfo *Ast_ProcessInfo,ifSet *LnsSet,alt2type *LnsMap) {
    var workSet *LnsSet
    workSet = NewLnsSet([]LnsAny{})
    for _other := range( ifSet.Items ) {
        other := _other.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if self.FP.IsInheritFrom(processInfo, other, alt2type){
            workSet.Add(Ast_TypeInfo2Stem(other))
        } else { 
            for _ifType := range( self.ifSet.Items ) {
                ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if other.FP.IsInheritFrom(processInfo, ifType, alt2type){
                    workSet.Add(Ast_TypeInfo2Stem(ifType))
                }
            }
        }
    }
    self.ifSet = workSet
    
}

// 4961: decl @lune.@base.@Ast.CombineType.createStem
func (self *Ast_CombineType) CreateStem(processInfo *Ast_ProcessInfo) *Ast_TypeInfo {
    var retType *Ast_TypeInfo
    if self.nilable{
        retType = Ast_builtinTypeStem_
        
    } else { 
        retType = Ast_builtinTypeStem
        
    }
    if Ast_isMutable(self.mutMode){
        return retType
    }
    return processInfo.FP.CreateModifier(retType, self.mutMode)
}

// 4975: decl @lune.@base.@Ast.CombineType.get_typeInfo
func (self *Ast_CombineType) Get_typeInfo(processInfo *Ast_ProcessInfo) *Ast_TypeInfo {
    if self.ifSet.Len() != 1{
        return self.FP.CreateStem(processInfo)
    }
    for _ifType := range( self.ifSet.Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var work *Ast_TypeInfo
        work = ifType
        if self.nilable{
            work = work.FP.Get_nilableTypeInfo()
            
        }
        if Ast_isMutable(self.mutMode){
            return work
        }
        return processInfo.FP.CreateModifier(work, self.mutMode)
    }
    panic("illegal")
// insert a dummy
    return nil
}

// 5006: decl @lune.@base.@Ast.CombineType.andType
func (self *Ast_CombineType) AndType(processInfo *Ast_ProcessInfo,other LnsAny,alt2type *LnsMap) LnsAny {
    switch _exp20122 := other.(type) {
    case *Ast_CommonType__Combine:
    comboInfo := _exp20122.Val1
        self.FP.andIfSet(processInfo, comboInfo.ifSet, alt2type)
        if Lns_op_not(Ast_isMutable(comboInfo.mutMode)){
            self.mutMode = comboInfo.mutMode
            
        }
        return &Ast_CommonType__Combine{self}
    case *Ast_CommonType__Normal:
    typeInfo := _exp20122.Val1
        if Lns_op_not(Ast_isMutable(typeInfo.FP.Get_mutMode())){
            self.mutMode = typeInfo.FP.Get_mutMode()
            
        }
        var ifSet *LnsSet
        ifSet = NewLnsSet([]LnsAny{})
        if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__IF{
            ifSet.Add(Ast_TypeInfo2Stem(typeInfo))
        } else { 
            for _, _iftype := range( typeInfo.FP.Get_interfaceList().Items ) {
                iftype := _iftype.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                ifSet.Add(Ast_TypeInfo2Stem(iftype))
            }
        }
        self.FP.andIfSet(processInfo, ifSet, alt2type)
        if self.ifSet.Len() != 0{
            return &Ast_CommonType__Combine{self}
        }
        return &Ast_CommonType__Normal{self.FP.CreateStem(processInfo)}
    }
// insert a dummy
    return nil
}


// declaration Class -- AbbrTypeInfo
type Ast_AbbrTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_AbbrTypeInfo struct {
    Ast_TypeInfo
    typeId LnsInt
    rawTxt string
    FP Ast_AbbrTypeInfoMtd
}
func Ast_AbbrTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_AbbrTypeInfo).FP
}
type Ast_AbbrTypeInfoDownCast interface {
    ToAst_AbbrTypeInfo() *Ast_AbbrTypeInfo
}
func Ast_AbbrTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_AbbrTypeInfoDownCast)
    if ok { return work.ToAst_AbbrTypeInfo() }
    return nil
}
func (obj *Ast_AbbrTypeInfo) ToAst_AbbrTypeInfo() *Ast_AbbrTypeInfo {
    return obj
}
func NewAst_AbbrTypeInfo(arg1 *Ast_ProcessInfo, arg2 string) *Ast_AbbrTypeInfo {
    obj := &Ast_AbbrTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AbbrTypeInfo(arg1, arg2)
    return obj
}
func (self *Ast_AbbrTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_AbbrTypeInfo) Get_rawTxt() string{ return self.rawTxt }
// 5307: decl @lune.@base.@Ast.AbbrTypeInfo.get_scope
func (self *Ast_AbbrTypeInfo) Get_scope() LnsAny {
    return nil
}

// 5312: decl @lune.@base.@Ast.AbbrTypeInfo.get_baseTypeInfo
func (self *Ast_AbbrTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5316: decl @lune.@base.@Ast.AbbrTypeInfo.get_parentInfo
func (self *Ast_AbbrTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5320: DeclConstr
func (self *Ast_AbbrTypeInfo) InitAst_AbbrTypeInfo(processInfo *Ast_ProcessInfo,rawTxt string) {
    self.InitAst_TypeInfo(nil, processInfo)
    var typeId LnsInt
    typeId = processInfo.FP.Get_idProv().FP.Get_id() + 1
    processInfo.FP.Get_idProv().FP.Increment()
    self.typeId = typeId
    
    self.rawTxt = rawTxt
    
}

// 5331: decl @lune.@base.@Ast.AbbrTypeInfo.isModule
func (self *Ast_AbbrTypeInfo) IsModule() bool {
    return false
}

// 5336: decl @lune.@base.@Ast.AbbrTypeInfo.getTxt
func (self *Ast_AbbrTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 5342: decl @lune.@base.@Ast.AbbrTypeInfo.getTxtWithRaw
func (self *Ast_AbbrTypeInfo) GetTxtWithRaw(rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return rawTxt
}

// 5349: decl @lune.@base.@Ast.AbbrTypeInfo.canEvalWith
func (self *Ast_AbbrTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return false, nil
}

// 5356: decl @lune.@base.@Ast.AbbrTypeInfo.serialize
func (self *Ast_AbbrTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    Util_err("illegal call")
}

// 5360: decl @lune.@base.@Ast.AbbrTypeInfo.get_display_stirng_with
func (self *Ast_AbbrTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(raw, nil, nil, nil)
}

// 5364: decl @lune.@base.@Ast.AbbrTypeInfo.get_display_stirng
func (self *Ast_AbbrTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
}

// 5368: decl @lune.@base.@Ast.AbbrTypeInfo.getModule
func (self *Ast_AbbrTypeInfo) GetModule() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5371: decl @lune.@base.@Ast.AbbrTypeInfo.get_kind
func (self *Ast_AbbrTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Abbr
}

// 5374: decl @lune.@base.@Ast.AbbrTypeInfo.get_nilable
func (self *Ast_AbbrTypeInfo) Get_nilable() bool {
    return true
}

// 5377: decl @lune.@base.@Ast.AbbrTypeInfo.get_nilableTypeInfo
func (self *Ast_AbbrTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5380: decl @lune.@base.@Ast.AbbrTypeInfo.get_mutMode
func (self *Ast_AbbrTypeInfo) Get_mutMode() LnsInt {
    return Ast_MutMode__IMut
}

// 5383: decl @lune.@base.@Ast.AbbrTypeInfo.get_accessMode
func (self *Ast_AbbrTypeInfo) Get_accessMode() LnsInt {
    return Ast_AccessMode__Local
}


// declaration Class -- ExtTypeInfo
type Ast_ExtTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_ExtTypeInfo struct {
    Ast_TypeInfo
    typeId LnsInt
    extedType *Ast_TypeInfo
    nilableTypeInfo *Ast_TypeInfo
    FP Ast_ExtTypeInfoMtd
}
func Ast_ExtTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_ExtTypeInfo).FP
}
type Ast_ExtTypeInfoDownCast interface {
    ToAst_ExtTypeInfo() *Ast_ExtTypeInfo
}
func Ast_ExtTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_ExtTypeInfoDownCast)
    if ok { return work.ToAst_ExtTypeInfo() }
    return nil
}
func (obj *Ast_ExtTypeInfo) ToAst_ExtTypeInfo() *Ast_ExtTypeInfo {
    return obj
}
func NewAst_ExtTypeInfo(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo) *Ast_ExtTypeInfo {
    obj := &Ast_ExtTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_ExtTypeInfo(arg1, arg2)
    return obj
}
func (self *Ast_ExtTypeInfo) Get_typeId() LnsInt{ return self.typeId }
func (self *Ast_ExtTypeInfo) Get_extedType() *Ast_TypeInfo{ return self.extedType }
func (self *Ast_ExtTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo{ return self.nilableTypeInfo }
func (self *Ast_ExtTypeInfo) AddChildren(arg1 *Ast_TypeInfo) {
self.extedType. FP.AddChildren( arg1)
}
func (self *Ast_ExtTypeInfo) CreateAlt2typeMap(arg1 bool) *LnsMap {
    return self.extedType. FP.CreateAlt2typeMap( arg1)
}
func (self *Ast_ExtTypeInfo) GetFullName(arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.extedType. FP.GetFullName( arg1,arg2,arg3)
}
func (self *Ast_ExtTypeInfo) GetOverridingType() LnsAny {
    return self.extedType. FP.GetOverridingType( )
}
func (self *Ast_ExtTypeInfo) GetParentFullName(arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.extedType. FP.GetParentFullName( arg1,arg2,arg3)
}
func (self *Ast_ExtTypeInfo) GetParentId() LnsInt {
    return self.extedType. FP.GetParentId( )
}
func (self *Ast_ExtTypeInfo) getProcessInfo() *Ast_ProcessInfo {
    return self.extedType. FP.getProcessInfo( )
}
func (self *Ast_ExtTypeInfo) Get_abstractFlag() bool {
    return self.extedType. FP.Get_abstractFlag( )
}
func (self *Ast_ExtTypeInfo) Get_accessMode() LnsInt {
    return self.extedType. FP.Get_accessMode( )
}
func (self *Ast_ExtTypeInfo) Get_argTypeInfoList() *LnsList {
    return self.extedType. FP.Get_argTypeInfoList( )
}
func (self *Ast_ExtTypeInfo) Get_autoFlag() bool {
    return self.extedType. FP.Get_autoFlag( )
}
func (self *Ast_ExtTypeInfo) Get_baseId() LnsInt {
    return self.extedType. FP.Get_baseId( )
}
func (self *Ast_ExtTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return self.extedType. FP.Get_baseTypeInfo( )
}
func (self *Ast_ExtTypeInfo) Get_children() *LnsList {
    return self.extedType. FP.Get_children( )
}
func (self *Ast_ExtTypeInfo) Get_externalFlag() bool {
    return self.extedType. FP.Get_externalFlag( )
}
func (self *Ast_ExtTypeInfo) Get_genSrcTypeInfo() *Ast_TypeInfo {
    return self.extedType. FP.Get_genSrcTypeInfo( )
}
func (self *Ast_ExtTypeInfo) Get_interfaceList() *LnsList {
    return self.extedType. FP.Get_interfaceList( )
}
func (self *Ast_ExtTypeInfo) Get_itemTypeInfoList() *LnsList {
    return self.extedType. FP.Get_itemTypeInfoList( )
}
func (self *Ast_ExtTypeInfo) Get_mutMode() LnsInt {
    return self.extedType. FP.Get_mutMode( )
}
func (self *Ast_ExtTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return self.extedType. FP.Get_parentInfo( )
}
func (self *Ast_ExtTypeInfo) Get_processInfo() *Ast_ProcessInfo {
    return self.extedType. FP.Get_processInfo( )
}
func (self *Ast_ExtTypeInfo) Get_rawTxt() string {
    return self.extedType. FP.Get_rawTxt( )
}
func (self *Ast_ExtTypeInfo) Get_retTypeInfoList() *LnsList {
    return self.extedType. FP.Get_retTypeInfoList( )
}
func (self *Ast_ExtTypeInfo) Get_scope() LnsAny {
    return self.extedType. FP.Get_scope( )
}
func (self *Ast_ExtTypeInfo) Get_staticFlag() bool {
    return self.extedType. FP.Get_staticFlag( )
}
func (self *Ast_ExtTypeInfo) Get_typeData() *Ast_TypeData {
    return self.extedType. FP.Get_typeData( )
}
func (self *Ast_ExtTypeInfo) HasBase() bool {
    return self.extedType. FP.HasBase( )
}
func (self *Ast_ExtTypeInfo) HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool {
    return self.extedType. FP.HasRouteNamespaceFrom( arg1)
}
func (self *Ast_ExtTypeInfo) IsInheritFrom(arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.extedType. FP.IsInheritFrom( arg1,arg2,arg3)
}
func (self *Ast_ExtTypeInfo) IsModule() bool {
    return self.extedType. FP.IsModule( )
}
func (self *Ast_ExtTypeInfo) SerializeTypeInfoList(arg1 string,arg2 *LnsList,arg3 LnsAny) string {
    return self.extedType. FP.SerializeTypeInfoList( arg1,arg2,arg3)
}
func (self *Ast_ExtTypeInfo) SwitchScope(arg1 *Ast_Scope) {
self.extedType. FP.SwitchScope( arg1)
}
// 5400: DeclConstr
func (self *Ast_ExtTypeInfo) InitAst_ExtTypeInfo(processInfo *Ast_ProcessInfo,extedType *Ast_TypeInfo) {
    self.InitAst_TypeInfo(extedType.FP.Get_scope(), processInfo)
    var typeId LnsInt
    typeId = processInfo.FP.Get_idProv().FP.Get_id() + 1
    processInfo.FP.Get_idProv().FP.Increment()
    self.typeId = typeId
    
    self.extedType = extedType
    
    self.nilableTypeInfo = &NewAst_NilableTypeInfo(nil, processInfo, &self.Ast_TypeInfo, typeId + 1).Ast_TypeInfo
    
    processInfo.FP.Get_idProv().FP.Increment()
}

// 5416: decl @lune.@base.@Ast.ExtTypeInfo.getTxt
func (self *Ast_ExtTypeInfo) GetTxt(typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(self.FP.Get_rawTxt(), typeNameCtrl, importInfo, localFlag)
}

// 5422: decl @lune.@base.@Ast.ExtTypeInfo.getTxtWithRaw
func (self *Ast_ExtTypeInfo) GetTxtWithRaw(rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return Lns_getVM().String_format("Luaval<%s>", []LnsAny{self.extedType.FP.GetTxtWithRaw(rawTxt, typeNameCtrl, importInfo, localFlag)})
}

// 5430: decl @lune.@base.@Ast.ExtTypeInfo.equals
func (self *Ast_ExtTypeInfo) Equals(processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    {
        _extTypeInfo := Ast_ExtTypeInfoDownCastF(typeInfo.FP)
        if _extTypeInfo != nil {
            extTypeInfo := _extTypeInfo.(*Ast_ExtTypeInfo)
            return self.extedType.FP.Equals(processInfo, extTypeInfo.extedType, alt2type, checkModifer)
        }
    }
    if Lns_isCondTrue( Lns_car(Ast_failCreateLuavalWith_4234_(self.extedType, Ast_LuavalConvKind__InLua, false))){
        return false
    }
    return self.extedType.FP.Equals(processInfo, typeInfo, alt2type, checkModifer)
}

// 5446: decl @lune.@base.@Ast.ExtTypeInfo.canEvalWith
func (self *Ast_ExtTypeInfo) CanEvalWith(processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    {
        _extTypeInfo := Ast_ExtTypeInfoDownCastF(other.FP.Get_nonnilableType().FP)
        if _extTypeInfo != nil {
            extTypeInfo := _extTypeInfo.(*Ast_ExtTypeInfo)
            var otherExtedType *Ast_TypeInfo
            if other.FP.Get_nilable(){
                otherExtedType = extTypeInfo.extedType.FP.Get_nilableTypeInfo()
                
            } else { 
                otherExtedType = extTypeInfo.extedType
                
            }
            return self.extedType.FP.CanEvalWith(processInfo, otherExtedType, canEvalType, alt2type)
        }
    }
    {
        __exp := Ast_convExp21900(Lns_2DDD(Ast_failCreateLuavalWith_4234_(other, Ast_LuavalConvKind__ToLua, true)))
        if __exp != nil {
            _exp := __exp.(string)
            return false, _exp
        }
    }
    return true, nil
}

// 5470: decl @lune.@base.@Ast.ExtTypeInfo.serialize
func (self *Ast_ExtTypeInfo) Serialize(stream Lns_oStream,validChildrenSet LnsAny) {
    stream.Write(Lns_getVM().String_format("{ skind = %d, typeId = %d, extedTypeId = %d }\n", []LnsAny{Ast_SerializeKind__Ext, self.typeId, self.extedType.FP.Get_typeId()}))
}

// 5476: decl @lune.@base.@Ast.ExtTypeInfo.get_display_stirng_with
func (self *Ast_ExtTypeInfo) Get_display_stirng_with(raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(raw, nil, nil, nil)
}

// 5480: decl @lune.@base.@Ast.ExtTypeInfo.get_display_stirng
func (self *Ast_ExtTypeInfo) Get_display_stirng() string {
    return self.FP.Get_display_stirng_with(self.FP.Get_rawTxt(), nil)
}

// 5484: decl @lune.@base.@Ast.ExtTypeInfo.getModule
func (self *Ast_ExtTypeInfo) GetModule() *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5487: decl @lune.@base.@Ast.ExtTypeInfo.get_kind
func (self *Ast_ExtTypeInfo) Get_kind() LnsInt {
    return Ast_TypeInfoKind__Ext
}

// 5490: decl @lune.@base.@Ast.ExtTypeInfo.get_aliasSrc
func (self *Ast_ExtTypeInfo) Get_aliasSrc() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5493: decl @lune.@base.@Ast.ExtTypeInfo.get_srcTypeInfo
func (self *Ast_ExtTypeInfo) Get_srcTypeInfo() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5496: decl @lune.@base.@Ast.ExtTypeInfo.get_nonnilableType
func (self *Ast_ExtTypeInfo) Get_nonnilableType() *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5499: decl @lune.@base.@Ast.ExtTypeInfo.get_nilable
func (self *Ast_ExtTypeInfo) Get_nilable() bool {
    return false
}

// 5502: decl @lune.@base.@Ast.ExtTypeInfo.applyGeneric
func (self *Ast_ExtTypeInfo) ApplyGeneric(alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.extedType.FP.ApplyGeneric(alt2typeMap, moduleTypeInfo)
    if typeInfo != self.extedType{
        Util_err(Lns_getVM().String_format("not support generics -- %s", []LnsAny{self.extedType.FP.GetTxt(nil, nil, nil)}))
    }
    return &self.Ast_TypeInfo
}


// declaration Class -- AndExpTypeInfo
type Ast_AndExpTypeInfoMtd interface {
    AddChildren(arg1 *Ast_TypeInfo)
    ApplyGeneric(arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(arg1 bool) *LnsMap
    Equals(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule() *Ast_TypeInfo
    GetOverridingType() LnsAny
    GetParentFullName(arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId() LnsInt
    getProcessInfo() *Ast_ProcessInfo
    GetTxt(arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag() bool
    Get_accessMode() LnsInt
    Get_aliasSrc() *Ast_TypeInfo
    Get_argTypeInfoList() *LnsList
    Get_autoFlag() bool
    Get_baseId() LnsInt
    Get_baseTypeInfo() *Ast_TypeInfo
    Get_children() *LnsList
    Get_display_stirng() string
    Get_display_stirng_with(arg1 string, arg2 LnsAny) string
    Get_exp1() *Ast_TypeInfo
    Get_exp2() *Ast_TypeInfo
    Get_extedType() *Ast_TypeInfo
    Get_externalFlag() bool
    Get_genSrcTypeInfo() *Ast_TypeInfo
    Get_interfaceList() *LnsList
    Get_itemTypeInfoList() *LnsList
    Get_kind() LnsInt
    Get_mutMode() LnsInt
    Get_nilable() bool
    Get_nilableTypeInfo() *Ast_TypeInfo
    Get_nonnilableType() *Ast_TypeInfo
    Get_parentInfo() *Ast_TypeInfo
    Get_processInfo() *Ast_ProcessInfo
    Get_rawTxt() string
    Get_result() *Ast_TypeInfo
    Get_retTypeInfoList() *LnsList
    Get_scope() LnsAny
    Get_srcTypeInfo() *Ast_TypeInfo
    Get_staticFlag() bool
    Get_typeData() *Ast_TypeData
    Get_typeId() LnsInt
    HasBase() bool
    HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool
    IsInheritFrom(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule() bool
    Serialize(arg1 Lns_oStream, arg2 LnsAny)
    SerializeTypeInfoList(arg1 string, arg2 *LnsList, arg3 LnsAny) string
    SwitchScope(arg1 *Ast_Scope)
}
type Ast_AndExpTypeInfo struct {
    Ast_TypeInfo
    exp1 *Ast_TypeInfo
    exp2 *Ast_TypeInfo
    result *Ast_TypeInfo
    FP Ast_AndExpTypeInfoMtd
}
func Ast_AndExpTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_AndExpTypeInfo).FP
}
type Ast_AndExpTypeInfoDownCast interface {
    ToAst_AndExpTypeInfo() *Ast_AndExpTypeInfo
}
func Ast_AndExpTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_AndExpTypeInfoDownCast)
    if ok { return work.ToAst_AndExpTypeInfo() }
    return nil
}
func (obj *Ast_AndExpTypeInfo) ToAst_AndExpTypeInfo() *Ast_AndExpTypeInfo {
    return obj
}
func NewAst_AndExpTypeInfo(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 *Ast_TypeInfo, arg4 *Ast_TypeInfo) *Ast_AndExpTypeInfo {
    obj := &Ast_AndExpTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AndExpTypeInfo(arg1, arg2, arg3, arg4)
    return obj
}
func (self *Ast_AndExpTypeInfo) Get_exp1() *Ast_TypeInfo{ return self.exp1 }
func (self *Ast_AndExpTypeInfo) Get_exp2() *Ast_TypeInfo{ return self.exp2 }
func (self *Ast_AndExpTypeInfo) Get_result() *Ast_TypeInfo{ return self.result }
func (self *Ast_AndExpTypeInfo) AddChildren(arg1 *Ast_TypeInfo) {
self.result. FP.AddChildren( arg1)
}
func (self *Ast_AndExpTypeInfo) ApplyGeneric(arg1 *LnsMap,arg2 *Ast_TypeInfo) LnsAny {
    return self.result. FP.ApplyGeneric( arg1,arg2)
}
func (self *Ast_AndExpTypeInfo) CanEvalWith(arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsInt,arg4 *LnsMap)(bool, LnsAny) {
    return self.result. FP.CanEvalWith( arg1,arg2,arg3,arg4)
}
func (self *Ast_AndExpTypeInfo) CreateAlt2typeMap(arg1 bool) *LnsMap {
    return self.result. FP.CreateAlt2typeMap( arg1)
}
func (self *Ast_AndExpTypeInfo) Equals(arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny,arg4 LnsAny) bool {
    return self.result. FP.Equals( arg1,arg2,arg3,arg4)
}
func (self *Ast_AndExpTypeInfo) GetFullName(arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.result. FP.GetFullName( arg1,arg2,arg3)
}
func (self *Ast_AndExpTypeInfo) GetModule() *Ast_TypeInfo {
    return self.result. FP.GetModule( )
}
func (self *Ast_AndExpTypeInfo) GetOverridingType() LnsAny {
    return self.result. FP.GetOverridingType( )
}
func (self *Ast_AndExpTypeInfo) GetParentFullName(arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.result. FP.GetParentFullName( arg1,arg2,arg3)
}
func (self *Ast_AndExpTypeInfo) GetParentId() LnsInt {
    return self.result. FP.GetParentId( )
}
func (self *Ast_AndExpTypeInfo) getProcessInfo() *Ast_ProcessInfo {
    return self.result. FP.getProcessInfo( )
}
func (self *Ast_AndExpTypeInfo) GetTxt(arg1 LnsAny,arg2 LnsAny,arg3 LnsAny) string {
    return self.result. FP.GetTxt( arg1,arg2,arg3)
}
func (self *Ast_AndExpTypeInfo) GetTxtWithRaw(arg1 string,arg2 LnsAny,arg3 LnsAny,arg4 LnsAny) string {
    return self.result. FP.GetTxtWithRaw( arg1,arg2,arg3,arg4)
}
func (self *Ast_AndExpTypeInfo) Get_abstractFlag() bool {
    return self.result. FP.Get_abstractFlag( )
}
func (self *Ast_AndExpTypeInfo) Get_accessMode() LnsInt {
    return self.result. FP.Get_accessMode( )
}
func (self *Ast_AndExpTypeInfo) Get_aliasSrc() *Ast_TypeInfo {
    return self.result. FP.Get_aliasSrc( )
}
func (self *Ast_AndExpTypeInfo) Get_argTypeInfoList() *LnsList {
    return self.result. FP.Get_argTypeInfoList( )
}
func (self *Ast_AndExpTypeInfo) Get_autoFlag() bool {
    return self.result. FP.Get_autoFlag( )
}
func (self *Ast_AndExpTypeInfo) Get_baseId() LnsInt {
    return self.result. FP.Get_baseId( )
}
func (self *Ast_AndExpTypeInfo) Get_baseTypeInfo() *Ast_TypeInfo {
    return self.result. FP.Get_baseTypeInfo( )
}
func (self *Ast_AndExpTypeInfo) Get_children() *LnsList {
    return self.result. FP.Get_children( )
}
func (self *Ast_AndExpTypeInfo) Get_display_stirng() string {
    return self.result. FP.Get_display_stirng( )
}
func (self *Ast_AndExpTypeInfo) Get_display_stirng_with(arg1 string,arg2 LnsAny) string {
    return self.result. FP.Get_display_stirng_with( arg1,arg2)
}
func (self *Ast_AndExpTypeInfo) Get_extedType() *Ast_TypeInfo {
    return self.result. FP.Get_extedType( )
}
func (self *Ast_AndExpTypeInfo) Get_externalFlag() bool {
    return self.result. FP.Get_externalFlag( )
}
func (self *Ast_AndExpTypeInfo) Get_genSrcTypeInfo() *Ast_TypeInfo {
    return self.result. FP.Get_genSrcTypeInfo( )
}
func (self *Ast_AndExpTypeInfo) Get_interfaceList() *LnsList {
    return self.result. FP.Get_interfaceList( )
}
func (self *Ast_AndExpTypeInfo) Get_itemTypeInfoList() *LnsList {
    return self.result. FP.Get_itemTypeInfoList( )
}
func (self *Ast_AndExpTypeInfo) Get_kind() LnsInt {
    return self.result. FP.Get_kind( )
}
func (self *Ast_AndExpTypeInfo) Get_mutMode() LnsInt {
    return self.result. FP.Get_mutMode( )
}
func (self *Ast_AndExpTypeInfo) Get_nilable() bool {
    return self.result. FP.Get_nilable( )
}
func (self *Ast_AndExpTypeInfo) Get_nilableTypeInfo() *Ast_TypeInfo {
    return self.result. FP.Get_nilableTypeInfo( )
}
func (self *Ast_AndExpTypeInfo) Get_nonnilableType() *Ast_TypeInfo {
    return self.result. FP.Get_nonnilableType( )
}
func (self *Ast_AndExpTypeInfo) Get_parentInfo() *Ast_TypeInfo {
    return self.result. FP.Get_parentInfo( )
}
func (self *Ast_AndExpTypeInfo) Get_processInfo() *Ast_ProcessInfo {
    return self.result. FP.Get_processInfo( )
}
func (self *Ast_AndExpTypeInfo) Get_rawTxt() string {
    return self.result. FP.Get_rawTxt( )
}
func (self *Ast_AndExpTypeInfo) Get_retTypeInfoList() *LnsList {
    return self.result. FP.Get_retTypeInfoList( )
}
func (self *Ast_AndExpTypeInfo) Get_scope() LnsAny {
    return self.result. FP.Get_scope( )
}
func (self *Ast_AndExpTypeInfo) Get_srcTypeInfo() *Ast_TypeInfo {
    return self.result. FP.Get_srcTypeInfo( )
}
func (self *Ast_AndExpTypeInfo) Get_staticFlag() bool {
    return self.result. FP.Get_staticFlag( )
}
func (self *Ast_AndExpTypeInfo) Get_typeData() *Ast_TypeData {
    return self.result. FP.Get_typeData( )
}
func (self *Ast_AndExpTypeInfo) Get_typeId() LnsInt {
    return self.result. FP.Get_typeId( )
}
func (self *Ast_AndExpTypeInfo) HasBase() bool {
    return self.result. FP.HasBase( )
}
func (self *Ast_AndExpTypeInfo) HasRouteNamespaceFrom(arg1 *Ast_TypeInfo) bool {
    return self.result. FP.HasRouteNamespaceFrom( arg1)
}
func (self *Ast_AndExpTypeInfo) IsInheritFrom(arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.result. FP.IsInheritFrom( arg1,arg2,arg3)
}
func (self *Ast_AndExpTypeInfo) IsModule() bool {
    return self.result. FP.IsModule( )
}
func (self *Ast_AndExpTypeInfo) Serialize(arg1 Lns_oStream,arg2 LnsAny) {
self.result. FP.Serialize( arg1,arg2)
}
func (self *Ast_AndExpTypeInfo) SerializeTypeInfoList(arg1 string,arg2 *LnsList,arg3 LnsAny) string {
    return self.result. FP.SerializeTypeInfoList( arg1,arg2,arg3)
}
func (self *Ast_AndExpTypeInfo) SwitchScope(arg1 *Ast_Scope) {
self.result. FP.SwitchScope( arg1)
}
// 5621: DeclConstr
func (self *Ast_AndExpTypeInfo) InitAst_AndExpTypeInfo(processInfo *Ast_ProcessInfo,exp1 *Ast_TypeInfo,exp2 *Ast_TypeInfo,result *Ast_TypeInfo) {
    self.InitAst_TypeInfo(result.FP.Get_scope(), processInfo)
    self.exp1 = exp1
    
    self.exp2 = exp2
    
    self.result = result
    
}


// declaration Class -- RefTypeInfo
type Ast_RefTypeInfoMtd interface {
    Get_itemRefTypeList() *LnsList
    Get_pos() *Types_Position
    Get_typeInfo() *Ast_TypeInfo
}
type Ast_RefTypeInfo struct {
    pos *Types_Position
    itemRefTypeList *LnsList
    typeInfo *Ast_TypeInfo
    FP Ast_RefTypeInfoMtd
}
func Ast_RefTypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_RefTypeInfo).FP
}
type Ast_RefTypeInfoDownCast interface {
    ToAst_RefTypeInfo() *Ast_RefTypeInfo
}
func Ast_RefTypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_RefTypeInfoDownCast)
    if ok { return work.ToAst_RefTypeInfo() }
    return nil
}
func (obj *Ast_RefTypeInfo) ToAst_RefTypeInfo() *Ast_RefTypeInfo {
    return obj
}
func NewAst_RefTypeInfo(arg1 *Types_Position, arg2 *LnsList, arg3 *Ast_TypeInfo) *Ast_RefTypeInfo {
    obj := &Ast_RefTypeInfo{}
    obj.FP = obj
    obj.InitAst_RefTypeInfo(arg1, arg2, arg3)
    return obj
}
func (self *Ast_RefTypeInfo) InitAst_RefTypeInfo(arg1 *Types_Position, arg2 *LnsList, arg3 *Ast_TypeInfo) {
    self.pos = arg1
    self.itemRefTypeList = arg2
    self.typeInfo = arg3
}
func (self *Ast_RefTypeInfo) Get_pos() *Types_Position{ return self.pos }
func (self *Ast_RefTypeInfo) Get_itemRefTypeList() *LnsList{ return self.itemRefTypeList }
func (self *Ast_RefTypeInfo) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }

// declaration Class -- TypeAnalyzer
type Ast_TypeAnalyzerMtd interface {
    AnalyzeType(arg1 *Ast_Scope, arg2 Parser_PushbackParser, arg3 LnsInt, arg4 bool, arg5 bool)(LnsAny, LnsAny, LnsAny)
    AnalyzeTypeFromTxt(arg1 string, arg2 *Ast_Scope, arg3 LnsInt, arg4 bool)(LnsAny, LnsAny, LnsAny)
    AnalyzeTypeItemList(arg1 bool, arg2 bool, arg3 bool, arg4 *Ast_TypeInfo, arg5 *Types_Position)(LnsAny, LnsAny, LnsAny)
    analyzeTypeSub(arg1 bool)(LnsAny, LnsAny, LnsAny)
}
type Ast_TypeAnalyzer struct {
    parser Parser_PushbackParser
    parentInfo *Ast_TypeInfo
    moduleType *Ast_TypeInfo
    moduleScope *Ast_Scope
    scope *Ast_Scope
    scopeAccess LnsInt
    accessMode LnsInt
    parentPub bool
    validMutControl bool
    processInfo *Ast_ProcessInfo
    FP Ast_TypeAnalyzerMtd
}
func Ast_TypeAnalyzer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_TypeAnalyzer).FP
}
type Ast_TypeAnalyzerDownCast interface {
    ToAst_TypeAnalyzer() *Ast_TypeAnalyzer
}
func Ast_TypeAnalyzerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_TypeAnalyzerDownCast)
    if ok { return work.ToAst_TypeAnalyzer() }
    return nil
}
func (obj *Ast_TypeAnalyzer) ToAst_TypeAnalyzer() *Ast_TypeAnalyzer {
    return obj
}
func NewAst_TypeAnalyzer(arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 *Ast_TypeInfo, arg4 *Ast_Scope, arg5 LnsInt, arg6 bool) *Ast_TypeAnalyzer {
    obj := &Ast_TypeAnalyzer{}
    obj.FP = obj
    obj.InitAst_TypeAnalyzer(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
// 6723: DeclConstr
func (self *Ast_TypeAnalyzer) InitAst_TypeAnalyzer(processInfo *Ast_ProcessInfo,parentInfo *Ast_TypeInfo,moduleType *Ast_TypeInfo,moduleScope *Ast_Scope,scopeAccess LnsInt,validMutControl bool) {
    self.processInfo = processInfo
    
    self.parentInfo = parentInfo
    
    self.moduleType = moduleType
    
    self.moduleScope = moduleScope
    
    self.scopeAccess = scopeAccess
    
    self.validMutControl = validMutControl
    
    self.scope = Ast_rootScope
    
    self.accessMode = Ast_AccessMode__Local
    
    self.parentPub = false
    
    self.parser = NewParser_DefaultPushbackParser(&NewParser_DummyParser().Parser_Parser).FP
    
}



// 6748: decl @lune.@base.@Ast.TypeAnalyzer.analyzeType
func (self *Ast_TypeAnalyzer) AnalyzeType(scope *Ast_Scope,parser Parser_PushbackParser,accessMode LnsInt,allowDDD bool,parentPub bool)(LnsAny, LnsAny, LnsAny) {
    self.scope = scope
    
    self.parser = parser
    
    self.accessMode = accessMode
    
    self.parentPub = parentPub
    
    return self.FP.analyzeTypeSub(allowDDD)
}

// 6759: decl @lune.@base.@Ast.TypeAnalyzer.analyzeTypeFromTxt
func (self *Ast_TypeAnalyzer) AnalyzeTypeFromTxt(txt string,scope *Ast_Scope,accessMode LnsInt,parentPub bool)(LnsAny, LnsAny, LnsAny) {
    var stream *Parser_TxtStream
    stream = NewParser_TxtStream(txt)
    var parser *Parser_DefaultPushbackParser
    parser = NewParser_DefaultPushbackParser(&NewParser_StreamParser(stream.FP, "test", nil).Parser_Parser)
    return self.FP.AnalyzeType(scope, parser.FP, accessMode, true, parentPub)
}

// 6770: decl @lune.@base.@Ast.TypeAnalyzer.analyzeTypeSub
func (self *Ast_TypeAnalyzer) analyzeTypeSub(allowDDD bool)(LnsAny, LnsAny, LnsAny) {
    var firstToken *Types_Token
    firstToken = self.parser.GetTokenNoErr()
    var token *Types_Token
    token = firstToken
    var refFlag bool
    refFlag = false
    if token.Txt == "&"{
        refFlag = true
        
        token = self.parser.GetTokenNoErr()
        
    }
    var mutFlag bool
    mutFlag = false
    if token.Txt == "mut"{
        mutFlag = true
        
        token = self.parser.GetTokenNoErr()
        
    }
    var typeInfo *Ast_TypeInfo
    if token.Txt == "..."{
        typeInfo = Ast_builtinTypeDDD
        
    } else { 
        var symbol *Ast_SymbolInfo
        
        {
            _symbol := self.scope.FP.GetSymbolTypeInfo(token.Txt, self.scope, self.moduleScope, self.scopeAccess)
            if _symbol == nil{
                return nil, token.Pos, "not found type -- " + token.Txt
            } else {
                symbol = _symbol.(*Ast_SymbolInfo)
            }
        }
        if symbol.FP.Get_kind() != Ast_SymbolKind__Typ{
            return nil, token.Pos, Lns_getVM().String_format("illegal type -- %s", []LnsAny{symbol.FP.Get_name()})
        }
        typeInfo = symbol.FP.Get_typeInfo()
        
    }
    return self.FP.AnalyzeTypeItemList(allowDDD, refFlag, mutFlag, typeInfo, token.Pos)
}

// 6810: decl @lune.@base.@Ast.TypeAnalyzer.analyzeTypeItemList
func (self *Ast_TypeAnalyzer) AnalyzeTypeItemList(allowDDD bool,refFlag bool,mutFlag bool,typeInfo *Ast_TypeInfo,pos *Types_Position)(LnsAny, LnsAny, LnsAny) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.parentPub) &&
        Lns_GetEnv().SetStackVal( Ast_isPubToExternal(self.accessMode)) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(typeInfo.FP.Get_accessMode()))) ).(bool)){
        return nil, pos, Lns_getVM().String_format("This type must be public. -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)})
    }
    var token *Types_Token
    token = self.parser.GetTokenNoErr()
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( token.Consecutive) &&
        Lns_GetEnv().SetStackVal( token.Txt == "!") ).(bool)){
        typeInfo = typeInfo.FP.Get_nilableTypeInfo()
        
        token = self.parser.GetTokenNoErr()
        
    }
    var genericRefList *LnsList
    genericRefList = NewLnsList([]LnsAny{})
    for  {
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == "[") ||
            Lns_GetEnv().SetStackVal( token.Txt == "[@") ).(bool){
            if token.Txt == "["{
                typeInfo = self.processInfo.FP.CreateList(self.accessMode, self.parentInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), Ast_MutMode__Mut)
                
            } else { 
                typeInfo = self.processInfo.FP.CreateArray(self.accessMode, self.parentInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), Ast_MutMode__Mut)
                
            }
            token = self.parser.GetTokenNoErr()
            
            if token.Txt != "]"{
                return nil, token.Pos, "not found -- ']'"
            }
            
        } else if token.Txt == "<"{
            var genericList *LnsList
            genericList = NewLnsList([]LnsAny{})
            var nextToken *Types_Token
            nextToken = Parser_getEofToken()
            for {
                var refType LnsAny
                refType = Ast_convExp29173(Lns_2DDD(self.FP.analyzeTypeSub(false)))
                if refType != nil{
                    refType_5417 := refType.(*Ast_RefTypeInfo)
                    genericRefList.Insert(Ast_RefTypeInfo2Stem(refType_5417))
                    genericList.Insert(Ast_TypeInfo2Stem(refType_5417.FP.Get_typeInfo()))
                }
                nextToken = self.parser.GetTokenNoErr()
                
                if nextToken.Txt != ","{ break }
            }
            if nextToken.Txt != ">"{
                return nil, nextToken.Pos, "not found -- ']'"
            }
            
            if _switch29800 := typeInfo.FP.Get_kind(); _switch29800 == Ast_TypeInfoKind__Map {
                if genericList.Len() != 2{
                    return nil, pos, "Key or value type is unknown"
                } else { 
                    typeInfo = self.processInfo.FP.CreateMap(self.accessMode, self.parentInfo, genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), genericList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__Mut)
                    
                }
            } else if _switch29800 == Ast_TypeInfoKind__List {
                if genericList.Len() != 1{
                    return nil, pos, Lns_getVM().String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = self.processInfo.FP.CreateList(self.accessMode, self.parentInfo, genericList, Ast_MutMode__Mut)
                
            } else if _switch29800 == Ast_TypeInfoKind__Array {
                if genericList.Len() != 1{
                    return nil, pos, Lns_getVM().String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = self.processInfo.FP.CreateArray(self.accessMode, self.parentInfo, genericList, Ast_MutMode__Mut)
                
            } else if _switch29800 == Ast_TypeInfoKind__Set {
                if genericList.Len() != 1{
                    return nil, pos, Lns_getVM().String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = self.processInfo.FP.CreateSet(self.accessMode, self.parentInfo, genericList, Ast_MutMode__Mut)
                
            } else if _switch29800 == Ast_TypeInfoKind__DDD {
                if genericList.Len() != 1{
                    return nil, pos, Lns_getVM().String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = &self.processInfo.FP.CreateDDD(genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), false, false).Ast_TypeInfo
                
            } else if _switch29800 == Ast_TypeInfoKind__Class || _switch29800 == Ast_TypeInfoKind__IF {
                if genericList.Len() != typeInfo.FP.Get_itemTypeInfoList().Len(){
                    return nil, pos, Lns_getVM().String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                for _, _itemType := range( genericList.Items ) {
                    itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if itemType.FP.Get_nilable(){
                        var mess string
                        mess = Lns_getVM().String_format("can't use nilable type -- %s", []LnsAny{itemType.FP.GetTxt(nil, nil, nil)})
                        return nil, pos, mess
                    }
                }
                typeInfo = &self.processInfo.FP.CreateGeneric(typeInfo, genericList, self.moduleType).Ast_TypeInfo
                
            } else if _switch29800 == Ast_TypeInfoKind__Box {
                if genericList.Len() != 1{
                    return nil, pos, Lns_getVM().String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = self.processInfo.FP.CreateBox(self.accessMode, genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                
            } else if _switch29800 == Ast_TypeInfoKind__Ext {
                if genericList.Len() != 1{
                    return nil, pos, Lns_getVM().String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                switch _exp29778 := self.processInfo.FP.CreateLuaval(genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), true).(type) {
                case *Ast_LuavalResult__OK:
                extTypeInfo := _exp29778.Val1
                    typeInfo = extTypeInfo
                    
                case *Ast_LuavalResult__Err:
                err := _exp29778.Val1
                    return nil, pos, err
                }
            } else {
                return nil, pos, Lns_getVM().String_format("not support generic: %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)})
            }
        } else { 
            self.parser.Pushback()
            break
        }
        token = self.parser.GetTokenNoErr()
        
    }
    if token.Txt == "!"{
        typeInfo = typeInfo.FP.Get_nilableTypeInfo()
        
        self.parser.GetTokenNoErr()
    }
    if Lns_op_not(allowDDD){
        if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__DDD{
            return nil, pos, Lns_getVM().String_format("invalid type. -- '%s'", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)})
        }
    }
    if refFlag{
        if self.validMutControl{
            typeInfo = self.processInfo.FP.CreateModifier(typeInfo, Ast_MutMode__IMut)
            
        }
    }
    return NewAst_RefTypeInfo(pos, genericRefList, typeInfo), nil, nil
}


func Lns_Ast_init() {
    if init_Ast { return }
    init_Ast = true
    Ast__mod__ = "@lune.@base.@Ast"
    Lns_InitMod()
    Lns_Parser_init()
    Lns_Util_init()
    Lns_Code_init()
    Lns_Log_init()
    Ast_extStartId = 100000
    Ast_extMaxId = 10000000
    Ast_userStartId = 1000
    Ast_rootProcessInfo = Ast_ProcessInfo_createRoot_1061_()
    Ast_rootTypeId = Ast_rootProcessInfo.FP.Get_idProv().FP.GetNewId()
    Ast_DummyModuleInfoManager____init_1141_()
    Ast_sym2builtInTypeMap = NewLnsMap( map[LnsAny]LnsAny{})
    Ast_builtInTypeIdSet = NewLnsMap( map[LnsAny]LnsAny{})
    Ast_txt2AccessModeMap = NewLnsMap( map[LnsAny]LnsAny{})
    Ast_txt2AccessModeMap.Set("none",Ast_AccessMode__None)
    Ast_txt2AccessModeMap.Set("pub",Ast_AccessMode__Pub)
    Ast_txt2AccessModeMap.Set("pro",Ast_AccessMode__Pro)
    Ast_txt2AccessModeMap.Set("pri",Ast_AccessMode__Pri)
    Ast_txt2AccessModeMap.Set("local",Ast_AccessMode__Local)
    Ast_txt2AccessModeMap.Set("global",Ast_AccessMode__Global)
    Ast_accessMode2txtMap = NewLnsMap( map[LnsAny]LnsAny{})
    Ast_accessMode2txtMap.Set(Ast_AccessMode__None,"none")
    Ast_accessMode2txtMap.Set(Ast_AccessMode__Pub,"pub")
    Ast_accessMode2txtMap.Set(Ast_AccessMode__Pro,"pro")
    Ast_accessMode2txtMap.Set(Ast_AccessMode__Pri,"pri")
    Ast_accessMode2txtMap.Set(Ast_AccessMode__Local,"local")
    Ast_accessMode2txtMap.Set(Ast_AccessMode__Global,"global")
    Ast_rootScope = NewAst_Scope(Ast_rootProcessInfo, nil, false, nil, nil)
    Ast_dummyList = NewLnsList([]LnsAny{})
    _ = NewLnsList([]LnsAny{})
    Ast_headTypeInfo = &Ast_RootTypeInfo_create_1961_().Ast_TypeInfo
    Ast_defaultTypeNameCtrl = NewAst_TypeNameCtrl(Ast_headTypeInfo)
    Ast_AutoBoxingInfo____init_2085_()
    Ast_CanEvalCtrlTypeInfo____init_2121_()
    Ast_dummySymbol = Lns_unwrap( Lns_car(Ast_rootScope.FP.AddLocalVar(Ast_rootProcessInfo, false, false, "$$", nil, Ast_headTypeInfo, Ast_MutMode__IMut))).(*Ast_SymbolInfo)
    Ast_boxRootAltType = NewAst_AlternateTypeInfo(Ast_rootProcessInfo, true, 1, "_T", Ast_AccessMode__Pub, Ast_headTypeInfo, nil, nil)
    Ast_boxRootScope = NewAst_Scope(Ast_rootProcessInfo, Ast_rootScope, true, nil, nil)
    Ast_rootProcessInfo.FP.set_typeInfo2Map(NewAst_TypeInfo2Map())
    Ast_immutableTypeSet = NewLnsSet([]LnsAny{})
    Ast_rootProcessInfo.FP.Get_idProv().FP.Increment()
    Ast_builtinTypeNone = Ast_NormalTypeInfo_createBuiltin("__None", "", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeEmpty = Ast_NormalTypeInfo_createBuiltin("__Empty", "::", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeNeverRet = Ast_NormalTypeInfo_createBuiltin("__NRet", "__", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeStem = Ast_NormalTypeInfo_createBuiltin("Stem", "stem", Ast_TypeInfoKind__Stem, nil, nil)
    Ast_builtinTypeStem_ = Ast_builtinTypeStem.FP.Get_nilableTypeInfo()
    Ast_builtinTypeBool = Ast_NormalTypeInfo_createBuiltin("Bool", "bool", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeInt = Ast_NormalTypeInfo_createBuiltin("Int", "int", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeReal = Ast_NormalTypeInfo_createBuiltin("Real", "real", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeChar = Ast_NormalTypeInfo_createBuiltin("char", "__char", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeMapping = Ast_NormalTypeInfo_createBuiltin("Mapping", "Mapping", Ast_TypeInfoKind__IF, nil, nil)
    Ast_builtinTypeAsyncItem = Ast_NormalTypeInfo_createBuiltin("__AsyncItem", "__AsyncItem", Ast_TypeInfoKind__IF, nil, nil)
    Ast_builtinTypeString = Ast_NormalTypeInfo_createBuiltin("String", "str", Ast_TypeInfoKind__Class, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeMapping)}))
    Ast_builtinTypeMap = Ast_NormalTypeInfo_createBuiltin("Map", "Map", Ast_TypeInfoKind__Map, nil, nil)
    Ast_builtinTypeSet = Ast_NormalTypeInfo_createBuiltin("Set", "Set", Ast_TypeInfoKind__Set, nil, nil)
    Ast_builtinTypeList = Ast_NormalTypeInfo_createBuiltin("List", "List", Ast_TypeInfoKind__List, nil, nil)
    Ast_builtinTypeArray = Ast_NormalTypeInfo_createBuiltin("Array", "Array", Ast_TypeInfoKind__Array, nil, nil)
    Ast_immutableTypeSet.Add(Ast_TypeInfo2Stem(Ast_builtinTypeBool))
    Ast_immutableTypeSet.Add(Ast_TypeInfo2Stem(Ast_builtinTypeInt))
    Ast_immutableTypeSet.Add(Ast_TypeInfo2Stem(Ast_builtinTypeReal))
    Ast_immutableTypeSet.Add(Ast_TypeInfo2Stem(Ast_builtinTypeChar))
    Ast_immutableTypeSet.Add(Ast_TypeInfo2Stem(Ast_builtinTypeString))
    Ast_builtinTypeBox = NewAst_BoxTypeInfo(Ast_rootProcessInfo, Ast_rootProcessInfo.FP.Get_idProv().FP.GetNewId(), Ast_AccessMode__Pub, &Ast_boxRootAltType.Ast_TypeInfo)
    Ast_registBuiltin_4166_("Nilable", "Nilable", Ast_TypeInfoKind__Box, &Ast_builtinTypeBox.Ast_TypeInfo, Ast_headTypeInfo, true)
    Ast_builtinTypeNil = &NewAst_NilTypeInfo(Ast_rootProcessInfo).Ast_TypeInfo
    Ast_registBuiltin_4166_("Nil", "nil", Ast_TypeInfoKind__Prim, Ast_builtinTypeNil, Ast_headTypeInfo, false)
    Ast_builtinTypeDDD = &Ast_rootProcessInfo.FP.CreateDDD(Ast_builtinTypeStem_, true, false).Ast_TypeInfo
    Ast_registBuiltin_4166_("DDD", "...", Ast_TypeInfoKind__DDD, Ast_builtinTypeDDD, Ast_headTypeInfo, false)
    Ast_builtinTypeForm = Ast_NormalTypeInfo_createBuiltin("Form", "form", Ast_TypeInfoKind__Form, Ast_builtinTypeDDD, nil)
    Ast_immutableTypeSet.Add(Ast_TypeInfo2Stem(Ast_builtinTypeForm))
    Ast_builtinTypeSymbol = Ast_NormalTypeInfo_createBuiltin("Symbol", "sym", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeStat = Ast_NormalTypeInfo_createBuiltin("Stat", "stat", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeExp = Ast_NormalTypeInfo_createBuiltin("Exp", "__exp", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeMultiExp = Ast_NormalTypeInfo_createBuiltin("Exps", "__exps", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeAbbr = &NewAst_AbbrTypeInfo(Ast_rootProcessInfo, "##").Ast_TypeInfo
    Ast_builtinTypeAbbrNone = &NewAst_AbbrTypeInfo(Ast_rootProcessInfo, "[##]").Ast_TypeInfo
    switch _exp22456 := Ast_rootProcessInfo.FP.CreateLuaval(Ast_builtinTypeStem, true).(type) {
    case *Ast_LuavalResult__OK:
    typeInfo := _exp22456.Val1
        Ast_builtinTypeLua = typeInfo
        
    default:
        Util_err("illegal")
    }
    Ast_registBuiltin_4166_("Luaval", "Luaval", Ast_TypeInfoKind__Ext, Ast_builtinTypeLua, Ast_headTypeInfo, false)
    Ast_builtinTypeDDDLua = &Ast_rootProcessInfo.FP.CreateDDD(Ast_builtinTypeStem_, true, true).Ast_TypeInfo
    Ast_registBuiltin_4166_("__LuaDDD", "__LuaDDD", Ast_TypeInfoKind__Ext, Ast_builtinTypeDDDLua, Ast_headTypeInfo, false)
    Ast_numberTypeSet = NewLnsSet([]LnsAny{})
    Ast_numberTypeSet.Add(Ast_TypeInfo2Stem(Ast_builtinTypeInt))
    Ast_numberTypeSet.Add(Ast_TypeInfo2Stem(Ast_builtinTypeChar))
    Ast_numberTypeSet.Add(Ast_TypeInfo2Stem(Ast_builtinTypeReal))
    Ast_builtinTypeInfo2Map = Ast_rootProcessInfo.FP.get_typeInfo2Map().FP.Clone()
    Ast_bitBinOpMap = NewLnsMap( map[LnsAny]LnsAny{"&":Ast_BitOpKind__And,"|":Ast_BitOpKind__Or,"~":Ast_BitOpKind__Xor,"|>>":Ast_BitOpKind__RShift,"|<<":Ast_BitOpKind__LShift,})
    Ast_compOpSet = NewLnsSet([]LnsAny{"==", "~="})
    Ast_mathCompOpSet = NewLnsSet([]LnsAny{"<", "<=", ">", ">="})
}
func init() {
    init_Ast = false
}
