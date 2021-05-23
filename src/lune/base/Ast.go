// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Ast bool
var Ast__mod__ string
// decl enum -- TypeInfoKind 
type Ast_TypeInfoKind = LnsInt
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
const Ast_TypeInfoKind__ExtModule = 26
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
  Ast_TypeInfoKind__ExtModule,
})
func Ast_TypeInfoKind_get__allList(_env *LnsEnv) *LnsList{
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
  Ast_TypeInfoKind__ExtModule: "TypeInfoKind.ExtModule",
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
func Ast_TypeInfoKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_TypeInfoKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_TypeInfoKind_getTxt(arg1 LnsInt) string {
    return Ast_TypeInfoKindMap_[arg1];
}
// decl enum -- IdType 
type Ast_IdType = LnsInt
const Ast_IdType__Base = 0
const Ast_IdType__Ext = 1
var Ast_IdTypeList_ = NewLnsList( []LnsAny {
  Ast_IdType__Base,
  Ast_IdType__Ext,
})
func Ast_IdType_get__allList(_env *LnsEnv) *LnsList{
    return Ast_IdTypeList_
}
var Ast_IdTypeMap_ = map[LnsInt]string {
  Ast_IdType__Base: "IdType.Base",
  Ast_IdType__Ext: "IdType.Ext",
}
func Ast_IdType__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_IdTypeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_IdType_getTxt(arg1 LnsInt) string {
    return Ast_IdTypeMap_[arg1];
}
// decl enum -- AccessMode 
type Ast_AccessMode = LnsInt
const Ast_AccessMode__Direct = 5
const Ast_AccessMode__Global = 6
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
  Ast_AccessMode__Direct,
  Ast_AccessMode__Global,
})
func Ast_AccessMode_get__allList(_env *LnsEnv) *LnsList{
    return Ast_AccessModeList_
}
var Ast_AccessModeMap_ = map[LnsInt]string {
  Ast_AccessMode__Direct: "AccessMode.Direct",
  Ast_AccessMode__Global: "AccessMode.Global",
  Ast_AccessMode__Local: "AccessMode.Local",
  Ast_AccessMode__None: "AccessMode.None",
  Ast_AccessMode__Pri: "AccessMode.Pri",
  Ast_AccessMode__Pro: "AccessMode.Pro",
  Ast_AccessMode__Pub: "AccessMode.Pub",
}
func Ast_AccessMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_AccessModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_AccessMode_getTxt(arg1 LnsInt) string {
    return Ast_AccessModeMap_[arg1];
}
// decl enum -- SymbolKind 
type Ast_SymbolKind = LnsInt
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
func Ast_SymbolKind_get__allList(_env *LnsEnv) *LnsList{
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
func Ast_SymbolKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_SymbolKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_SymbolKind_getTxt(arg1 LnsInt) string {
    return Ast_SymbolKindMap_[arg1];
}
// decl enum -- SerializeKind 
type Ast_SerializeKind = LnsInt
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
func Ast_SerializeKind_get__allList(_env *LnsEnv) *LnsList{
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
func Ast_SerializeKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_SerializeKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_SerializeKind_getTxt(arg1 LnsInt) string {
    return Ast_SerializeKindMap_[arg1];
}
// decl enum -- MutMode 
type Ast_MutMode = LnsInt
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
func Ast_MutMode_get__allList(_env *LnsEnv) *LnsList{
    return Ast_MutModeList_
}
var Ast_MutModeMap_ = map[LnsInt]string {
  Ast_MutMode__AllMut: "MutMode.AllMut",
  Ast_MutMode__IMut: "MutMode.IMut",
  Ast_MutMode__IMutRe: "MutMode.IMutRe",
  Ast_MutMode__Mut: "MutMode.Mut",
}
func Ast_MutMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_MutModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_MutMode_getTxt(arg1 LnsInt) string {
    return Ast_MutModeMap_[arg1];
}
// decl enum -- ScopeAccess 
type Ast_ScopeAccess = LnsInt
const Ast_ScopeAccess__Full = 1
const Ast_ScopeAccess__Normal = 0
var Ast_ScopeAccessList_ = NewLnsList( []LnsAny {
  Ast_ScopeAccess__Normal,
  Ast_ScopeAccess__Full,
})
func Ast_ScopeAccess_get__allList(_env *LnsEnv) *LnsList{
    return Ast_ScopeAccessList_
}
var Ast_ScopeAccessMap_ = map[LnsInt]string {
  Ast_ScopeAccess__Full: "ScopeAccess.Full",
  Ast_ScopeAccess__Normal: "ScopeAccess.Normal",
}
func Ast_ScopeAccess__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_ScopeAccessMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_ScopeAccess_getTxt(arg1 LnsInt) string {
    return Ast_ScopeAccessMap_[arg1];
}
// decl enum -- CanEvalType 
type Ast_CanEvalType = LnsInt
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
func Ast_CanEvalType_get__allList(_env *LnsEnv) *LnsList{
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
func Ast_CanEvalType__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_CanEvalTypeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_CanEvalType_getTxt(arg1 LnsInt) string {
    return Ast_CanEvalTypeMap_[arg1];
}
// decl enum -- Async 
type Ast_Async = LnsInt
const Ast_Async__Async = 1
const Ast_Async__Noasync = 0
const Ast_Async__Transient = 2
var Ast_AsyncList_ = NewLnsList( []LnsAny {
  Ast_Async__Noasync,
  Ast_Async__Async,
  Ast_Async__Transient,
})
func Ast_Async_get__allList(_env *LnsEnv) *LnsList{
    return Ast_AsyncList_
}
var Ast_AsyncMap_ = map[LnsInt]string {
  Ast_Async__Async: "Async.Async",
  Ast_Async__Noasync: "Async.Noasync",
  Ast_Async__Transient: "Async.Transient",
}
func Ast_Async__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_AsyncMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_Async_getTxt(arg1 LnsInt) string {
    return Ast_AsyncMap_[arg1];
}
// decl enum -- MethodKind 
type Ast_MethodKind = LnsInt
const Ast_MethodKind__All = 0
const Ast_MethodKind__Object = 2
const Ast_MethodKind__Static = 1
var Ast_MethodKindList_ = NewLnsList( []LnsAny {
  Ast_MethodKind__All,
  Ast_MethodKind__Static,
  Ast_MethodKind__Object,
})
func Ast_MethodKind_get__allList(_env *LnsEnv) *LnsList{
    return Ast_MethodKindList_
}
var Ast_MethodKindMap_ = map[LnsInt]string {
  Ast_MethodKind__All: "MethodKind.All",
  Ast_MethodKind__Object: "MethodKind.Object",
  Ast_MethodKind__Static: "MethodKind.Static",
}
func Ast_MethodKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_MethodKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_MethodKind_getTxt(arg1 LnsInt) string {
    return Ast_MethodKindMap_[arg1];
}
// decl enum -- LuavalConvKind 
type Ast_LuavalConvKind = LnsInt
const Ast_LuavalConvKind__InLua = 0
const Ast_LuavalConvKind__ToLua = 1
var Ast_LuavalConvKindList_ = NewLnsList( []LnsAny {
  Ast_LuavalConvKind__InLua,
  Ast_LuavalConvKind__ToLua,
})
func Ast_LuavalConvKind_get__allList_3310_(_env *LnsEnv) *LnsList{
    return Ast_LuavalConvKindList_
}
var Ast_LuavalConvKindMap_ = map[LnsInt]string {
  Ast_LuavalConvKind__InLua: "LuavalConvKind.InLua",
  Ast_LuavalConvKind__ToLua: "LuavalConvKind.ToLua",
}
func Ast_LuavalConvKind__from_3305_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_LuavalConvKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_LuavalConvKind_getTxt(arg1 LnsInt) string {
    return Ast_LuavalConvKindMap_[arg1];
}
// decl enum -- MatchType 
type Ast_MatchType = LnsInt
const Ast_MatchType__Error = 2
const Ast_MatchType__Match = 0
const Ast_MatchType__Warn = 1
var Ast_MatchTypeList_ = NewLnsList( []LnsAny {
  Ast_MatchType__Match,
  Ast_MatchType__Warn,
  Ast_MatchType__Error,
})
func Ast_MatchType_get__allList(_env *LnsEnv) *LnsList{
    return Ast_MatchTypeList_
}
var Ast_MatchTypeMap_ = map[LnsInt]string {
  Ast_MatchType__Error: "MatchType.Error",
  Ast_MatchType__Match: "MatchType.Match",
  Ast_MatchType__Warn: "MatchType.Warn",
}
func Ast_MatchType__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Ast_MatchTypeMap_[arg1]; ok { return arg1 }
    return nil
}

func Ast_MatchType_getTxt(arg1 LnsInt) string {
    return Ast_MatchTypeMap_[arg1];
}
// decl enum -- BitOpKind 
type Ast_BitOpKind = LnsInt
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
func Ast_BitOpKind_get__allList(_env *LnsEnv) *LnsList{
    return Ast_BitOpKindList_
}
var Ast_BitOpKindMap_ = map[LnsInt]string {
  Ast_BitOpKind__And: "BitOpKind.And",
  Ast_BitOpKind__LShift: "BitOpKind.LShift",
  Ast_BitOpKind__Or: "BitOpKind.Or",
  Ast_BitOpKind__RShift: "BitOpKind.RShift",
  Ast_BitOpKind__Xor: "BitOpKind.Xor",
}
func Ast_BitOpKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
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
var Ast_rootProcessInfoRo *Ast_ProcessInfo
var Ast_rootTypeIdInfo *Ast_IdInfo
var Ast_rootTypeId LnsInt
var Ast_sym2builtInTypeMapWork *LnsMap
var Ast_sym2builtInTypeMap *LnsMap
var Ast_builtInTypeIdSetWork *LnsMap
var Ast_builtInTypeIdSet *LnsMap
var Ast_txt2AccessModeMap *LnsMap
var Ast_accessMode2txtMap *LnsMap
var Ast_rootScope *Ast_Scope
var Ast_rootScopeRo *Ast_Scope
var Ast_dummyList *LnsList
var Ast_headTypeInfo *Ast_TypeInfo
var Ast_defaultTypeNameCtrl *Ast_TypeNameCtrl
var Ast_dummyIdInfo *Ast_IdInfo
var Ast_dummySymbol *Ast_SymbolInfo
var Ast_boxRootAltType *Ast_AlternateTypeInfo
var Ast_boxRootScope *Ast_Scope
var Ast_immutableTypeSetWork *LnsSet
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
var Ast_builtinTypeRunner *Ast_TypeInfo
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
// decl alge -- LuavalResult
type Ast_LuavalResult = LnsAny
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
// decl alge -- OverrideMut
type Ast_OverrideMut = LnsAny
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
// decl alge -- EnumLiteral
type Ast_EnumLiteral = LnsAny
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
// decl alge -- OverridingType
type Ast_OverridingType = LnsAny
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
// decl alge -- CommonType
type Ast_CommonType = LnsAny
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
type Ast_filterForm func (_env *LnsEnv, arg1 *Ast_SymbolInfo) bool
// for 5788
func Ast_convExp22961(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 5884
func Ast_convExp23373(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 3278
func Ast_convExp11224(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 3439
func Ast_convExp11913(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 4646
func Ast_convExp17547(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 4646
func Ast_convExp17665(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 4646
func Ast_convExp17719(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 5613
func Ast_convExp22239(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 6340
func Ast_convExp25347(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 6378
func Ast_convExp25564(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 6401
func Ast_convExp25818(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 6424
func Ast_convExp25971(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 6450
func Ast_convExp26192(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 6458
func Ast_convExp26336(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 6779
func Ast_convExp27993(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 6787
func Ast_convExp28030(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 6795
func Ast_convExp28067(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 6803
func Ast_convExp28104(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 6178
func Ast_convExp28322(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 6178
func Ast_convExp28454(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 6178
func Ast_convExp28551(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 6902
func Ast_convExp28848(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 6908
func Ast_convExp28884(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 6999
func Ast_convExp29365(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 7036
func Ast_convExp29545(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 7041
func Ast_convExp29565(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 7279
func Ast_convExp30696(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7415
func Ast_convExp31663(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7415
func Ast_convExp31820(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7415
func Ast_convExp31985(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7415
func Ast_convExp32155(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7415
func Ast_convExp32320(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7415
func Ast_convExp32527(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7415
func Ast_convExp32703(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7415
func Ast_convExp32860(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 7454
func Ast_convExp33014(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
func ProcessInfo_switchIdProvier___anonymous_1082_(_env *LnsEnv) string {
    return "start"
}
// 209: decl @lune.@base.@Ast.getRootProcessInfo
func Ast_getRootProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return Ast_rootProcessInfo
}

// 212: decl @lune.@base.@Ast.getRootProcessInfoRo
func Ast_getRootProcessInfoRo(_env *LnsEnv) *Ast_ProcessInfo {
    return Ast_rootProcessInfoRo
}

// 327: decl @lune.@base.@Ast.getSym2builtInTypeMap
func Ast_getSym2builtInTypeMap(_env *LnsEnv) *LnsMap {
    return Ast_sym2builtInTypeMap
}

// 330: decl @lune.@base.@Ast.getBuiltInTypeIdMap
func Ast_getBuiltInTypeIdMap(_env *LnsEnv) *LnsMap {
    return Ast_builtInTypeIdSet
}

// 350: decl @lune.@base.@Ast.isBuiltin
func Ast_isBuiltin(_env *LnsEnv, typeId LnsInt) bool {
    return Ast_TypeInfo2Stem(Ast_builtInTypeIdSet.Get(typeId)) != nil
}

// 355: decl @lune.@base.@Ast.isPubToExternal
func Ast_isPubToExternal(_env *LnsEnv, mode LnsInt) bool {
    if _switch891 := mode; _switch891 == Ast_AccessMode__Pub || _switch891 == Ast_AccessMode__Pro || _switch891 == Ast_AccessMode__Global {
        return true
    }
    return false
}

// 373: decl @lune.@base.@Ast.txt2AccessMode
func Ast_txt2AccessMode(_env *LnsEnv, accessMode string) LnsAny {
    return Ast_txt2AccessModeMap.Get(accessMode)
}

// 386: decl @lune.@base.@Ast.accessMode2txt
func Ast_accessMode2txt(_env *LnsEnv, accessMode LnsInt) string {
    return Lns_unwrap( Ast_accessMode2txtMap.Get(accessMode)).(string)
}

// 400: decl @lune.@base.@Ast.isMutable
func Ast_isMutable(_env *LnsEnv, mode LnsInt) bool {
    if _switch1044 := mode; _switch1044 == Ast_MutMode__AllMut || _switch1044 == Ast_MutMode__Mut {
        return true
    }
    return false
}

// 1087: decl @lune.@base.@Ast.applyGenericDefault
func Ast_applyGenericDefault(_env *LnsEnv, typeInfo *Ast_TypeInfo,alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    {
        _genType := typeInfo.FP.ApplyGeneric(_env, alt2typeMap, moduleTypeInfo)
        if !Lns_IsNil( _genType ) {
            genType := _genType.(*Ast_TypeInfo)
            return genType
        }
    }
    return typeInfo
}


// 1115: decl @lune.@base.@Ast.getAllNameForKind
func Ast_getAllNameForKind(_env *LnsEnv, classInfo *Ast_TypeInfo,kind LnsInt,symbolKind LnsInt) *Util_OrderedSet {
    var nameSet *Util_OrderedSet
    nameSet = NewUtil_OrderedSet(_env)
    var process func(_env *LnsEnv, scope *Ast_Scope)
    process = func(_env *LnsEnv, scope *Ast_Scope) {
        {
            _inherit := scope.FP.Get_inherit(_env)
            if !Lns_IsNil( _inherit ) {
                inherit := _inherit.(*Ast_Scope)
                process(_env, inherit)
            }
        }
        {
            __collection3066 := scope.FP.Get_symbol2SymbolInfoMap(_env)
            __sorted3066 := __collection3066.CreateKeyListStr()
            __sorted3066.Sort( LnsItemKindStr, nil )
            for _, ___key3066 := range( __sorted3066.Items ) {
                symbolInfo := __collection3066.Items[ ___key3066 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if _switch3064 := symbolInfo.FP.Get_kind(_env); _switch3064 == symbolKind {
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( symbolKind == Ast_SymbolKind__Mtd) &&
                        _env.SetStackVal( symbolInfo.FP.Get_name(_env) == "__init") ).(bool)){
                    } else { 
                        var staticFlag bool
                        staticFlag = symbolInfo.FP.Get_staticFlag(_env)
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( kind == Ast_MethodKind__All) ||
                            _env.SetStackVal( kind == Ast_MethodKind__Static) &&
                            _env.SetStackVal( staticFlag) ||
                            _env.SetStackVal( kind == Ast_MethodKind__Object) &&
                            _env.SetStackVal( Lns_op_not(staticFlag)) ).(bool){
                            nameSet.FP.Add(_env, symbolInfo.FP.Get_name(_env))
                        }
                    }
                }
            }
        }
    }
    {
        _scope := classInfo.FP.Get_scope(_env)
        if !Lns_IsNil( _scope ) {
            scope := _scope.(*Ast_Scope)
            process(_env, scope)
        }
    }
    return nameSet
}

// 1149: decl @lune.@base.@Ast.getAllMethodName
func Ast_getAllMethodName(_env *LnsEnv, classInfo *Ast_TypeInfo,kind LnsInt) *Util_OrderedSet {
    return Ast_getAllNameForKind(_env, classInfo, kind, Ast_SymbolKind__Mtd)
}

// 1202: decl @lune.@base.@Ast.getScope
func Ast_getScope(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny {
    return typeInfo.FP.Get_scope(_env)
}

// 1208: decl @lune.@base.@Ast.isExtId
func Ast_isExtId(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_typeId(_env).Id >= Ast_extStartId{
        return true
    }
    return false
}

// 2335: decl @lune.@base.@Ast.dumpScope.dumpScopeSub
func dumpScope__dumpScopeSub_2200_(_env *LnsEnv, scope LnsAny,prefix string,readyIdSet *LnsSet) {
    {
        __exp := scope
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_Scope)
            if readyIdSet.Has(Ast_Scope2Stem(_exp)){
                return 
            }
            readyIdSet.Add(Ast_Scope2Stem(_exp))
            if len(prefix) > 20{
                Util_err(_env, "illegal")
            }
            {
                __collection7877 := _exp.FP.Get_symbol2SymbolInfoMap(_env)
                __sorted7877 := __collection7877.CreateKeyListStr()
                __sorted7877.Sort( LnsItemKindStr, nil )
                for _, _symbol := range( __sorted7877.Items ) {
                    symbolInfo := __collection7877.Items[ _symbol ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbol := _symbol.(string)
                    Util_log(_env, _env.LuaVM.String_format("scope: %s, %s, %s", []LnsAny{prefix, _exp, symbol}))
                    {
                        _subScope := symbolInfo.FP.Get_typeInfo(_env).FP.Get_scope(_env)
                        if !Lns_IsNil( _subScope ) {
                            subScope := _subScope.(*Ast_Scope)
                            dumpScope__dumpScopeSub_2200_(_env, subScope, prefix + "  ", readyIdSet)
                        }
                    }
                }
            }
        }
    }
}

// 2330: decl @lune.@base.@Ast.dumpScope
func Ast_dumpScope(_env *LnsEnv, workscope LnsAny,workprefix string) {
    dumpScope__dumpScopeSub_2200_(_env, workscope, workprefix, NewLnsSet([]LnsAny{}))
}


// 3348: decl @lune.@base.@Ast.isGenericType
func Ast_isGenericType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if Lns_isCondTrue( Ast_GenericTypeInfoDownCastF(typeInfo.FP)){
        return true
    }
    return false
}

// 3580: decl @lune.@base.@Ast.getEnumLiteralVal
func Ast_getEnumLiteralVal(_env *LnsEnv, obj LnsAny) LnsAny {
    switch _exp12469 := obj.(type) {
    case *Ast_EnumLiteral__Int:
    val := _exp12469.Val1
        return val
    case *Ast_EnumLiteral__Real:
    val := _exp12469.Val1
        return val
    case *Ast_EnumLiteral__Str:
    val := _exp12469.Val1
        return val
    }
// insert a dummy
    return nil
}


// 4340: decl @lune.@base.@Ast.isExtType
func Ast_isExtType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext) ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) &&
            _env.SetStackVal( typeInfo.FP.Get_extedType(_env) != typeInfo) ).(bool))) ).(bool)
}

// 4348: decl @lune.@base.@Ast.isMutableType
func Ast_isMutableType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    typeInfo = typeInfo.FP.Get_nonnilableType(_env)
    
    if Ast_immutableTypeSet.Has(Ast_TypeInfo2Stem(typeInfo)){
        return false
    }
    if _switch15846 := typeInfo.FP.Get_kind(_env); _switch15846 == Ast_TypeInfoKind__FormFunc || _switch15846 == Ast_TypeInfoKind__Enum {
        return false
    }
    return true
}

// 4445: decl @lune.@base.@Ast.addBuiltin
func Ast_addBuiltin(_env *LnsEnv, typeInfo *Ast_TypeInfo) {
    Ast_builtInTypeIdSetWork.Set(typeInfo.FP.Get_typeId(_env).Id,typeInfo)
}

// 4451: decl @lune.@base.@Ast.registBuiltin
func Ast_registBuiltin_3265_(_env *LnsEnv, idName string,typeTxt string,kind LnsInt,typeInfo *Ast_TypeInfo,nilableTypeInfo *Ast_TypeInfo,registScope bool) {
    Ast_sym2builtInTypeMapWork.Set(typeTxt,&NewAst_NormalSymbolInfo(_env, Ast_rootProcessInfo, Ast_SymbolKind__Typ, false, false, Ast_rootScope, Ast_AccessMode__Pub, false, typeTxt, nil, typeInfo, Ast_MutMode__IMut, true, false).Ast_SymbolInfo)
    if nilableTypeInfo != Ast_headTypeInfo{
        Ast_sym2builtInTypeMapWork.Set(typeTxt + "!",&NewAst_NormalSymbolInfo(_env, Ast_rootProcessInfo, Ast_SymbolKind__Typ, false, kind == Ast_TypeInfoKind__Func, Ast_rootScope, Ast_AccessMode__Pub, false, typeTxt, nil, nilableTypeInfo, Ast_MutMode__IMut, true, false).Ast_SymbolInfo)
    }
    Ast_addBuiltin(_env, typeInfo)
    Ast_addBuiltin(_env, Ast_rootProcessInfo.FP.CreateModifier(_env, typeInfo, Ast_MutMode__IMut))
    if typeInfo.FP.Get_nilableTypeInfo(_env) != Ast_headTypeInfo{
        Ast_addBuiltin(_env, typeInfo.FP.Get_nilableTypeInfo(_env))
        Ast_addBuiltin(_env, Ast_rootProcessInfo.FP.CreateModifier(_env, typeInfo.FP.Get_nilableTypeInfo(_env), Ast_MutMode__IMut))
    }
    if registScope{
        Ast_rootScope.FP.AddClass(_env, Ast_rootProcessInfo, typeTxt, nil, typeInfo)
    }
}

// 4570: decl @lune.@base.@Ast.isClass
func Ast_isClass(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
        _env.SetStackVal( typeInfo != Ast_builtinTypeString) ).(bool)
}

// 4641: decl @lune.@base.@Ast.failCreateLuavalWith
func Ast_failCreateLuavalWith_3312_(_env *LnsEnv, typeInfo *Ast_TypeInfo,convFlag LnsInt,validToCheck bool)(LnsAny, bool) {
    if Ast_isExtType(_env, typeInfo){
        return nil, true
    }
    var mess string
    mess = _env.LuaVM.String_format("not support to use the type as Luaval -- %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)})
    if _switch17762 := typeInfo.FP.Get_kind(_env); _switch17762 == Ast_TypeInfoKind__Nilable {
        return Ast_failCreateLuavalWith_3312_(_env, typeInfo.FP.Get_nonnilableType(_env), convFlag, validToCheck)
    } else if _switch17762 == Ast_TypeInfoKind__Prim {
        return nil, true
    } else if _switch17762 == Ast_TypeInfoKind__Form || _switch17762 == Ast_TypeInfoKind__IF || _switch17762 == Ast_TypeInfoKind__DDD || _switch17762 == Ast_TypeInfoKind__ExtModule {
        if Lns_op_not(validToCheck){
            return nil, false
        }
        if convFlag != Ast_LuavalConvKind__InLua{
            return mess, false
        }
        return nil, false
    } else if _switch17762 == Ast_TypeInfoKind__Stem {
        return nil, false
    } else if _switch17762 == Ast_TypeInfoKind__Class {
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
    } else if _switch17762 == Ast_TypeInfoKind__Array || _switch17762 == Ast_TypeInfoKind__List || _switch17762 == Ast_TypeInfoKind__Map {
        if Lns_op_not(validToCheck){
            return nil, false
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( convFlag != Ast_LuavalConvKind__ToLua) &&
            _env.SetStackVal( Ast_isMutable(_env, typeInfo.FP.Get_mutMode(_env))) ).(bool)){
            return "not support mutable collecion. " + mess, false
        }
        var canConv bool
        canConv = true
        for _, _itemType := range( typeInfo.FP.Get_itemTypeInfoList(_env).Items ) {
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var err LnsAny
            var work bool
            err,work = Ast_failCreateLuavalWith_3312_(_env, itemType, convFlag, validToCheck)
            if err != nil{
                err_3679 := err.(string)
                return err_3679, false
            }
            if Lns_op_not(work){
                canConv = false
                
            }
        }
        
        canConv = false
        
        return nil, canConv
    } else if _switch17762 == Ast_TypeInfoKind__FormFunc || _switch17762 == Ast_TypeInfoKind__Func {
        if Lns_op_not(validToCheck){
            return nil, false
        }
        if convFlag != Ast_LuavalConvKind__InLua{
            return mess, false
        }
        if typeInfo.FP.Get_itemTypeInfoList(_env).Len() != 0{
            return mess, false
        }
        var canConv bool
        canConv = true
        for _, _itemType := range( typeInfo.FP.Get_argTypeInfoList(_env).Items ) {
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var err LnsAny
            var work bool
            err,work = Ast_failCreateLuavalWith_3312_(_env, itemType, convFlag, validToCheck)
            if err != nil{
                err_3691 := err.(string)
                return err_3691, false
            }
            if Lns_op_not(work){
                canConv = false
                
            }
        }
        
        for _, _itemType := range( typeInfo.FP.Get_retTypeInfoList(_env).Items ) {
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var err LnsAny
            var work bool
            err,work = Ast_failCreateLuavalWith_3312_(_env, itemType, convFlag, validToCheck)
            if err != nil{
                err_3698 := err.(string)
                return err_3698, false
            }
            if Lns_op_not(work){
                canConv = false
                
            }
        }
        
        canConv = false
        
        return nil, canConv
    }
    return _env.LuaVM.String_format("not support -- %s:%s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil), Ast_TypeInfoKind_getTxt( typeInfo.FP.Get_kind(_env))}), false
}

// 4752: decl @lune.@base.@Ast.isStruct
func Ast_isStruct(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if _switch17916 := typeInfo.FP.Get_kind(_env); _switch17916 == Ast_TypeInfoKind__Class {
        if typeInfo == Ast_builtinTypeString{
            return false
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( typeInfo.FP.Get_baseTypeInfo(_env) != Ast_headTypeInfo) ||
            _env.SetStackVal( typeInfo.FP.Get_interfaceList(_env).Len() != 0) ||
            _env.SetStackVal( typeInfo.FP.Get_children(_env).Len() != 1) ).(bool){
            return false
        }
        return true
    }
    return false
}

// 4774: decl @lune.@base.@Ast.isConditionalbe
func Ast_isConditionalbe(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo) bool {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_nilable(_env)) ||
        _env.SetStackVal( typeInfo.FP.Equals(_env, processInfo, Ast_builtinTypeBool, nil, nil)) ).(bool){
        return true
    }
    return false
}











// 5588: decl @lune.@base.@Ast.applyGenericList
func Ast_applyGenericList_3567_(_env *LnsEnv, typeList *LnsList,alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo)(LnsAny, bool) {
    var typeInfoList *LnsList
    typeInfoList = NewLnsList([]LnsAny{})
    var needNew bool
    needNew = false
    for _, _srcType := range( typeList.Items ) {
        srcType := _srcType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        {
            _typeInfo := srcType.FP.ApplyGeneric(_env, alt2typeMap, moduleTypeInfo)
            if !Lns_IsNil( _typeInfo ) {
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


// 5923: decl @lune.@base.@Ast.convToExtTypeList
func Ast_convToExtTypeList(_env *LnsEnv, processInfo *Ast_ProcessInfo,list *LnsList)(LnsAny, string) {
    var extList *LnsList
    extList = NewLnsList([]LnsAny{})
    for _, _typeInfo := range( list.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        switch _exp23627 := processInfo.FP.CreateLuaval(_env, typeInfo, true).(type) {
        case *Ast_LuavalResult__OK:
        extType := _exp23627.Val1
            extList.Insert(Ast_TypeInfo2Stem(extType))
        case *Ast_LuavalResult__Err:
        err := _exp23627.Val1
            return nil, err
        }
    }
    return extList, ""
}

// 5969: decl @lune.@base.@Ast.isNumberType
func Ast_isNumberType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    return Ast_numberTypeSet.Has(Ast_TypeInfo2Stem(typeInfo.FP.Get_srcTypeInfo(_env)))
}




// 6502: decl @lune.@base.@Ast.isSettableToForm
func Ast_isSettableToForm_3956_(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo)(bool, string) {
    if typeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Direct{
        return false, "can't set the _direct function."
    }
    if typeInfo.FP.Get_argTypeInfoList(_env).Len() > 0{
        for _index, _argType := range( typeInfo.FP.Get_argTypeInfoList(_env).Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            {
                _dddType := Ast_DDDTypeInfoDownCastF(argType.FP)
                if !Lns_IsNil( _dddType ) {
                    dddType := _dddType.(*Ast_DDDTypeInfo)
                    if Lns_op_not(dddType.FP.Get_typeInfo(_env).FP.Get_nilableTypeInfo(_env).FP.Equals(_env, processInfo, Ast_builtinTypeStem_, nil, nil)){
                        return false, _env.LuaVM.String_format("mismatch arg%d", []LnsAny{index})
                    }
                } else {
                    if Lns_op_not(argType.FP.Get_srcTypeInfo(_env).FP.Equals(_env, processInfo, Ast_builtinTypeStem_, nil, nil)){
                        return false, _env.LuaVM.String_format("mismatch arg%d", []LnsAny{index})
                    }
                }
            }
        }
    }
    return true, ""
}



// 7075: decl @lune.@base.@Ast.isPrimitive
func Ast_isPrimitive(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    var srcType *Ast_TypeInfo
    srcType = typeInfo.FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env)
    if _switch29748 := srcType.FP.Get_kind(_env); _switch29748 == Ast_TypeInfoKind__Prim || _switch29748 == Ast_TypeInfoKind__Enum {
        return true
    }
    if srcType == Ast_builtinTypeString{
        return true
    }
    return false
}

// 7095: decl @lune.@base.@Ast.createProcessInfo
func Ast_createProcessInfo(_env *LnsEnv, validCheckingMutable bool,validExtType bool,validDetailError bool) *Ast_ProcessInfo {
    return Ast_ProcessInfo_createUser_1078_(_env, validCheckingMutable, validExtType, validDetailError, Ast_builtinTypeInfo2Map.FP.Clone(_env))
}


// declaration Class -- IdProvider
type Ast_IdProviderMtd interface {
    GetNewId(_env *LnsEnv) LnsInt
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
func NewAst_IdProvider(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *Ast_IdProvider {
    obj := &Ast_IdProvider{}
    obj.FP = obj
    obj.InitAst_IdProvider(_env, arg1, arg2)
    return obj
}
func (self *Ast_IdProvider) InitAst_IdProvider(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) {
    self.id = arg1
    self.maxId = arg2
}
// 37: decl @lune.@base.@Ast.IdProvider.getNewId
func (self *Ast_IdProvider) GetNewId(_env *LnsEnv) LnsInt {
    self.id = self.id + 1
    
    if self.id >= self.maxId{
        Util_err(_env, "id is over")
    }
    return self.id
}


// declaration Class -- ProcessInfo
type Ast_ProcessInfoMtd interface {
    CreateAdvertiseMethodFrom(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo) *Ast_TypeInfo
    CreateAlge(_env *LnsEnv, arg1 *Ast_Scope, arg2 *Ast_TypeInfo, arg3 bool, arg4 LnsInt, arg5 string) *Ast_AlgeTypeInfo
    CreateAlias(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 bool, arg4 LnsInt, arg5 *Ast_TypeInfo, arg6 *Ast_TypeInfo) *Ast_AliasTypeInfo
    CreateAlternate(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 string, arg4 LnsInt, arg5 *Ast_TypeInfo, arg6 LnsAny, arg7 LnsAny) *Ast_AlternateTypeInfo
    CreateArray(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 LnsInt) *Ast_TypeInfo
    CreateBox(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_TypeInfo) *Ast_TypeInfo
    CreateClass(_env *LnsEnv, arg1 bool, arg2 bool, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 *LnsList, arg7 *Ast_TypeInfo, arg8 bool, arg9 LnsInt, arg10 string) *Ast_TypeInfo
    CreateDDD(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool, arg3 bool) *Ast_DDDTypeInfo
    CreateEnum(_env *LnsEnv, arg1 *Ast_Scope, arg2 *Ast_TypeInfo, arg3 bool, arg4 LnsInt, arg5 string, arg6 *Ast_TypeInfo) *Ast_EnumTypeInfo
    CreateExtModule(_env *LnsEnv, arg1 LnsAny, arg2 *Ast_TypeInfo, arg3 bool, arg4 LnsInt, arg5 string, arg6 LnsInt, arg7 string) *Ast_TypeInfo
    CreateFunc(_env *LnsEnv, arg1 bool, arg2 bool, arg3 LnsAny, arg4 LnsInt, arg5 *Ast_TypeInfo, arg6 bool, arg7 bool, arg8 bool, arg9 LnsInt, arg10 string, arg11 LnsInt, arg12 LnsAny, arg13 LnsAny, arg14 LnsAny, arg15 LnsAny) *Ast_NormalTypeInfo
    CreateGeneric(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *LnsList, arg3 *Ast_TypeInfo) *Ast_GenericTypeInfo
    CreateList(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 LnsInt) *Ast_TypeInfo
    CreateLuaval(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) LnsAny
    CreateMap(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 *Ast_TypeInfo, arg4 *Ast_TypeInfo, arg5 LnsInt) *Ast_TypeInfo
    CreateModifier(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt) *Ast_TypeInfo
    CreateModule(_env *LnsEnv, arg1 *Ast_Scope, arg2 *Ast_TypeInfo, arg3 bool, arg4 string, arg5 bool) *Ast_TypeInfo
    CreateSet(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 LnsInt) *Ast_TypeInfo
    GetTypeInfo(_env *LnsEnv, arg1 LnsInt) LnsAny
    Get_idProv(_env *LnsEnv) *Ast_IdProvider
    Get_idProvBase(_env *LnsEnv) *Ast_IdProvider
    Get_idProvExt(_env *LnsEnv) *Ast_IdProvider
    Get_idProvScope(_env *LnsEnv) *Ast_IdProvider
    Get_idProvSym(_env *LnsEnv) *Ast_IdProvider
    get_typeInfo2Map(_env *LnsEnv) *Ast_TypeInfo2Map
    Get_validCheckingMutable(_env *LnsEnv) bool
    Get_validDetailError(_env *LnsEnv) bool
    Get_validExtType(_env *LnsEnv) bool
    newId(_env *LnsEnv, arg1 *Ast_TypeInfo) *Ast_IdInfo
    newIdForRoot(_env *LnsEnv) *Ast_IdInfo
    NewUser(_env *LnsEnv) *Ast_ProcessInfo
    setRootTypeInfo(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_TypeInfo)
    set_typeInfo2Map(_env *LnsEnv, arg1 *Ast_TypeInfo2Map)
    SwitchIdProvier(_env *LnsEnv, arg1 LnsInt)
}
type Ast_ProcessInfo struct {
    idProvSym *Ast_IdProvider
    idProvScope *Ast_IdProvider
    idProv *Ast_IdProvider
    idProvBase *Ast_IdProvider
    idProvExt *Ast_IdProvider
    typeInfo2Map LnsAny
    validExtType bool
    validCheckingMutable bool
    id2TypeInfo *LnsMap
    validDetailError bool
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
func NewAst_ProcessInfo(_env *LnsEnv, arg1 bool, arg2 *Ast_IdProvider, arg3 bool, arg4 bool, arg5 LnsAny) *Ast_ProcessInfo {
    obj := &Ast_ProcessInfo{}
    obj.FP = obj
    obj.InitAst_ProcessInfo(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Ast_ProcessInfo) Get_idProvSym(_env *LnsEnv) *Ast_IdProvider{ return self.idProvSym }
func (self *Ast_ProcessInfo) Get_idProvScope(_env *LnsEnv) *Ast_IdProvider{ return self.idProvScope }
func (self *Ast_ProcessInfo) Get_idProv(_env *LnsEnv) *Ast_IdProvider{ return self.idProv }
func (self *Ast_ProcessInfo) Get_idProvBase(_env *LnsEnv) *Ast_IdProvider{ return self.idProvBase }
func (self *Ast_ProcessInfo) Get_idProvExt(_env *LnsEnv) *Ast_IdProvider{ return self.idProvExt }
func (self *Ast_ProcessInfo) Get_validExtType(_env *LnsEnv) bool{ return self.validExtType }
func (self *Ast_ProcessInfo) Get_validCheckingMutable(_env *LnsEnv) bool{ return self.validCheckingMutable }
func (self *Ast_ProcessInfo) Get_validDetailError(_env *LnsEnv) bool{ return self.validDetailError }
// 116: decl @lune.@base.@Ast.ProcessInfo.get_typeInfo2Map
func (self *Ast_ProcessInfo) get_typeInfo2Map(_env *LnsEnv) *Ast_TypeInfo2Map {
    return Lns_unwrap( self.typeInfo2Map).(*Ast_TypeInfo2Map)
}

// 119: decl @lune.@base.@Ast.ProcessInfo.set_typeInfo2Map
func (self *Ast_ProcessInfo) set_typeInfo2Map(_env *LnsEnv, typeInfo2Map *Ast_TypeInfo2Map) {
    self.typeInfo2Map = typeInfo2Map
    
}

// 123: decl @lune.@base.@Ast.ProcessInfo.getTypeInfo
func (self *Ast_ProcessInfo) GetTypeInfo(_env *LnsEnv, id LnsInt) LnsAny {
    return self.id2TypeInfo.Get(id)
}

// 127: DeclConstr
func (self *Ast_ProcessInfo) InitAst_ProcessInfo(_env *LnsEnv, validCheckingMutable bool,idProvBase *Ast_IdProvider,validExtType bool,validDetailError bool,typeInfo2Map LnsAny) {
    self.validDetailError = validDetailError
    
    self.id2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.validCheckingMutable = validCheckingMutable
    
    self.validExtType = validExtType
    
    self.idProvBase = idProvBase
    
    self.idProvExt = NewAst_IdProvider(_env, Ast_extStartId, Ast_extMaxId)
    
    self.idProv = self.idProvBase
    
    self.idProvSym = NewAst_IdProvider(_env, 0, Ast_extMaxId)
    
    self.idProvScope = NewAst_IdProvider(_env, 0, Ast_extMaxId)
    
    self.typeInfo2Map = typeInfo2Map
    
}

// 143: decl @lune.@base.@Ast.ProcessInfo.createRoot
func Ast_ProcessInfo_createRoot_1076_(_env *LnsEnv) *Ast_ProcessInfo {
    return NewAst_ProcessInfo(_env, true, NewAst_IdProvider(_env, 0, Ast_userStartId), true, false, nil)
}

// 149: decl @lune.@base.@Ast.ProcessInfo.createUser
func Ast_ProcessInfo_createUser_1078_(_env *LnsEnv, validCheckingMutable bool,validExtType bool,validDetailError bool,typeInfo2Map LnsAny) *Ast_ProcessInfo {
    return NewAst_ProcessInfo(_env, validCheckingMutable, NewAst_IdProvider(_env, Ast_userStartId, Ast_extStartId), validExtType, validDetailError, typeInfo2Map)
}

// 158: decl @lune.@base.@Ast.ProcessInfo.switchIdProvier
func (self *Ast_ProcessInfo) SwitchIdProvier(_env *LnsEnv, idType LnsInt) {
    __func__ := "@lune.@base.@Ast.ProcessInfo.switchIdProvier"
    Log_log(_env, Log_Level__Trace, __func__, 159, Log_CreateMessage(ProcessInfo_switchIdProvier___anonymous_1082_))
    
    if idType == Ast_IdType__Base{
        self.idProv = self.idProvBase
        
    } else { 
        self.idProv = self.idProvExt
        
    }
}

// 216: decl @lune.@base.@Ast.ProcessInfo.newIdForRoot
func (self *Ast_ProcessInfo) newIdForRoot(_env *LnsEnv) *Ast_IdInfo {
    var id LnsInt
    id = self.idProv.FP.GetNewId(_env)
    return NewAst_IdInfo(_env, id, self)
}

// 220: decl @lune.@base.@Ast.ProcessInfo.setRootTypeInfo
func (self *Ast_ProcessInfo) setRootTypeInfo(_env *LnsEnv, id LnsInt,typeInfo *Ast_TypeInfo) {
    self.id2TypeInfo.Set(id,typeInfo)
}

// 224: decl @lune.@base.@Ast.ProcessInfo.newId
func (self *Ast_ProcessInfo) newId(_env *LnsEnv, typeInfo *Ast_TypeInfo) *Ast_IdInfo {
    var id LnsInt
    id = self.idProv.FP.GetNewId(_env)
    self.id2TypeInfo.Set(id,typeInfo)
    return NewAst_IdInfo(_env, id, self)
}


// 4289: decl @lune.@base.@Ast.ProcessInfo.createAlternate
func (self *Ast_ProcessInfo) CreateAlternate(_env *LnsEnv, belongClassFlag bool,altIndex LnsInt,txt string,accessMode LnsInt,parentInfo *Ast_TypeInfo,baseTypeInfo LnsAny,interfaceList LnsAny) *Ast_AlternateTypeInfo {
    return NewAst_AlternateTypeInfo(_env, self, belongClassFlag, altIndex, txt, accessMode, parentInfo, baseTypeInfo, interfaceList)
}

// 4361: decl @lune.@base.@Ast.ProcessInfo.createModifier
func (self *Ast_ProcessInfo) CreateModifier(_env *LnsEnv, srcTypeInfo *Ast_TypeInfo,mutMode LnsInt) *Ast_TypeInfo {
    srcTypeInfo = srcTypeInfo.FP.Get_srcTypeInfo(_env)
    
    if Lns_op_not(Ast_isMutableType(_env, srcTypeInfo)){
        return srcTypeInfo
    }
    if _switch15926 := mutMode; _switch15926 == Ast_MutMode__IMut || _switch15926 == Ast_MutMode__IMutRe {
        {
            __exp := self.FP.get_typeInfo2Map(_env).ImutModifierMap.Get(srcTypeInfo)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_TypeInfo)
                return _exp
            }
        }
    } else if _switch15926 == Ast_MutMode__AllMut {
        {
            __exp := self.FP.get_typeInfo2Map(_env).MutModifierMap.Get(srcTypeInfo)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_TypeInfo)
                return _exp
            }
        }
    }
    var modifier *Ast_TypeInfo
    if srcTypeInfo.FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        switch _exp15988 := self.FP.CreateLuaval(_env, self.FP.CreateModifier(_env, srcTypeInfo.FP.Get_extedType(_env), mutMode), false).(type) {
        case *Ast_LuavalResult__OK:
        workType := _exp15988.Val1
            if srcTypeInfo.FP.Get_nilable(_env){
                modifier = workType.FP.Get_nilableTypeInfo(_env)
                
            } else { 
                modifier = workType
                
            }
        case *Ast_LuavalResult__Err:
        err := _exp15988.Val1
            Util_err(_env, err)
        }
    } else { 
        modifier = &NewAst_ModifierTypeInfo(_env, self, srcTypeInfo, mutMode).Ast_TypeInfo
        
    }
    if _switch16049 := mutMode; _switch16049 == Ast_MutMode__IMut || _switch16049 == Ast_MutMode__IMutRe {
        self.FP.get_typeInfo2Map(_env).ImutModifierMap.Set(srcTypeInfo,modifier)
    } else if _switch16049 == Ast_MutMode__AllMut {
        self.FP.get_typeInfo2Map(_env).MutModifierMap.Set(srcTypeInfo,modifier)
    }
    return modifier
}

// 4783: decl @lune.@base.@Ast.ProcessInfo.createBox
func (self *Ast_ProcessInfo) CreateBox(_env *LnsEnv, accessMode LnsInt,nonnilableType *Ast_TypeInfo) *Ast_TypeInfo {
    {
        _boxType := self.FP.get_typeInfo2Map(_env).BoxMap.Get(nonnilableType)
        if !Lns_IsNil( _boxType ) {
            boxType := _boxType.(*Ast_TypeInfo)
            return boxType
        }
    }
    var boxType *Ast_BoxTypeInfo
    boxType = NewAst_BoxTypeInfo(_env, self, accessMode, nonnilableType)
    self.FP.get_typeInfo2Map(_env).BoxMap.Set(nonnilableType,&boxType.Ast_TypeInfo)
    return &boxType.Ast_TypeInfo
}

// 4827: decl @lune.@base.@Ast.ProcessInfo.createSet
func (self *Ast_ProcessInfo) CreateSet(_env *LnsEnv, accessMode LnsInt,parentInfo *Ast_TypeInfo,itemTypeInfo *LnsList,mutMode LnsInt) *Ast_TypeInfo {
    var tmpMutMode LnsInt
    if Ast_isMutable(_env, mutMode){
        tmpMutMode = mutMode
        
    } else { 
        tmpMutMode = Ast_MutMode__Mut
        
    }
    var newTypeFunc func(_env *LnsEnv, workMutMode LnsInt) *Ast_TypeInfo
    newTypeFunc = func(_env *LnsEnv, workMutMode LnsInt) *Ast_TypeInfo {
        return &NewAst_NormalTypeInfo(_env, self, false, Ast_getScope(_env, Ast_builtinTypeSet), Ast_builtinTypeSet, nil, false, false, false, Ast_AccessMode__Pub, "Set", Ast_headTypeInfo, Ast_TypeInfoKind__Set, itemTypeInfo, nil, nil, workMutMode, nil, Ast_Async__Async).Ast_TypeInfo
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = newTypeFunc(_env, tmpMutMode)
    if Ast_isMutable(_env, mutMode){
        return typeInfo
    }
    return self.FP.CreateModifier(_env, typeInfo, mutMode)
    
// insert a dummy
    return nil
}

// 4840: decl @lune.@base.@Ast.ProcessInfo.createList
func (self *Ast_ProcessInfo) CreateList(_env *LnsEnv, accessMode LnsInt,parentInfo *Ast_TypeInfo,itemTypeInfo *LnsList,mutMode LnsInt) *Ast_TypeInfo {
    var tmpMutMode LnsInt
    if Ast_isMutable(_env, mutMode){
        tmpMutMode = mutMode
        
    } else { 
        tmpMutMode = Ast_MutMode__Mut
        
    }
    var newTypeFunc func(_env *LnsEnv, workMutMode LnsInt) *Ast_TypeInfo
    newTypeFunc = func(_env *LnsEnv, workMutMode LnsInt) *Ast_TypeInfo {
        return &NewAst_NormalTypeInfo(_env, self, false, Ast_getScope(_env, Ast_builtinTypeList), Ast_builtinTypeList, nil, false, false, false, Ast_AccessMode__Pub, "List", Ast_headTypeInfo, Ast_TypeInfoKind__List, itemTypeInfo, nil, nil, workMutMode, nil, Ast_Async__Async).Ast_TypeInfo
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = newTypeFunc(_env, tmpMutMode)
    if Ast_isMutable(_env, mutMode){
        return typeInfo
    }
    return self.FP.CreateModifier(_env, typeInfo, mutMode)
    
// insert a dummy
    return nil
}

// 4853: decl @lune.@base.@Ast.ProcessInfo.createArray
func (self *Ast_ProcessInfo) CreateArray(_env *LnsEnv, accessMode LnsInt,parentInfo *Ast_TypeInfo,itemTypeInfo *LnsList,mutMode LnsInt) *Ast_TypeInfo {
    var tmpMutMode LnsInt
    if Ast_isMutable(_env, mutMode){
        tmpMutMode = mutMode
        
    } else { 
        tmpMutMode = Ast_MutMode__Mut
        
    }
    var newTypeFunc func(_env *LnsEnv, workMutMode LnsInt) *Ast_TypeInfo
    newTypeFunc = func(_env *LnsEnv, workMutMode LnsInt) *Ast_TypeInfo {
        return &NewAst_NormalTypeInfo(_env, self, false, Ast_getScope(_env, Ast_builtinTypeArray), Ast_builtinTypeArray, nil, false, false, false, Ast_AccessMode__Pub, "Array", Ast_headTypeInfo, Ast_TypeInfoKind__Array, itemTypeInfo, nil, nil, workMutMode, nil, Ast_Async__Async).Ast_TypeInfo
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = newTypeFunc(_env, tmpMutMode)
    if Ast_isMutable(_env, mutMode){
        return typeInfo
    }
    return self.FP.CreateModifier(_env, typeInfo, mutMode)
    
// insert a dummy
    return nil
}

// 4866: decl @lune.@base.@Ast.ProcessInfo.createMap
func (self *Ast_ProcessInfo) CreateMap(_env *LnsEnv, accessMode LnsInt,parentInfo *Ast_TypeInfo,keyTypeInfo *Ast_TypeInfo,valTypeInfo *Ast_TypeInfo,mutMode LnsInt) *Ast_TypeInfo {
    var tmpMutMode LnsInt
    if Ast_isMutable(_env, mutMode){
        tmpMutMode = mutMode
        
    } else { 
        tmpMutMode = Ast_MutMode__Mut
        
    }
    var newTypeFunc func(_env *LnsEnv, workMutMode LnsInt) *Ast_TypeInfo
    newTypeFunc = func(_env *LnsEnv, workMutMode LnsInt) *Ast_TypeInfo {
        return &NewAst_NormalTypeInfo(_env, self, false, Ast_getScope(_env, Ast_builtinTypeMap), Ast_builtinTypeMap, nil, false, false, false, Ast_AccessMode__Pub, "Map", Ast_headTypeInfo, Ast_TypeInfoKind__Map, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(keyTypeInfo), Ast_TypeInfo2Stem(valTypeInfo)}), nil, nil, workMutMode, nil, Ast_Async__Async).Ast_TypeInfo
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = newTypeFunc(_env, tmpMutMode)
    if Ast_isMutable(_env, mutMode){
        return typeInfo
    }
    return self.FP.CreateModifier(_env, typeInfo, mutMode)
    
// insert a dummy
    return nil
}

// 4881: decl @lune.@base.@Ast.ProcessInfo.createModule
func (self *Ast_ProcessInfo) CreateModule(_env *LnsEnv, scope *Ast_Scope,parentInfo *Ast_TypeInfo,externalFlag bool,moduleName string,mutable bool) *Ast_TypeInfo {
    {
        __exp := Ast_sym2builtInTypeMap.Get(moduleName)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_typeInfo(_env)
        }
    }
    if Parser_isLuaKeyword(_env, moduleName){
        Util_err(_env, _env.LuaVM.String_format("This symbol can not use for a class or script file. -- %s", []LnsAny{moduleName}))
    }
    var info *Ast_ModuleTypeInfo
    info = NewAst_ModuleTypeInfo(_env, self, scope, externalFlag, moduleName, parentInfo, mutable)
    return &info.Ast_TypeInfo
}

// 4901: decl @lune.@base.@Ast.ProcessInfo.createClass
func (self *Ast_ProcessInfo) CreateClass(_env *LnsEnv, classFlag bool,abstractFlag bool,scope LnsAny,baseInfo LnsAny,interfaceList LnsAny,genTypeList *LnsList,parentInfo *Ast_TypeInfo,externalFlag bool,accessMode LnsInt,className string) *Ast_TypeInfo {
    {
        __exp := Ast_sym2builtInTypeMap.Get(className)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_typeInfo(_env)
        }
    }
    if Parser_isLuaKeyword(_env, className){
        Util_err(_env, _env.LuaVM.String_format("This symbol can not use for a class or script file. -- %s", []LnsAny{className}))
    }
    var info *Ast_NormalTypeInfo
    info = NewAst_NormalTypeInfo(_env, self, abstractFlag, scope, baseInfo, interfaceList, false, externalFlag, false, accessMode, className, parentInfo, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( classFlag) &&
        _env.SetStackVal( Ast_TypeInfoKind__Class) ||
        _env.SetStackVal( Ast_TypeInfoKind__IF) ).(LnsInt), genTypeList, nil, nil, Ast_MutMode__Mut, nil, Ast_Async__Async)
    for _, _genType := range( genTypeList.Items ) {
        genType := _genType.(Ast_AlternateTypeInfoDownCast).ToAst_AlternateTypeInfo()
        genType.FP.updateParentInfo(_env, &info.Ast_TypeInfo)
    }
    return &info.Ast_TypeInfo
}

// 4935: decl @lune.@base.@Ast.ProcessInfo.createExtModule
func (self *Ast_ProcessInfo) CreateExtModule(_env *LnsEnv, scope LnsAny,parentInfo *Ast_TypeInfo,externalFlag bool,accessMode LnsInt,className string,moduleLang LnsInt,requirePath string) *Ast_TypeInfo {
    {
        __exp := Ast_sym2builtInTypeMap.Get(className)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_typeInfo(_env)
        }
    }
    if Parser_isLuaKeyword(_env, className){
        Util_err(_env, _env.LuaVM.String_format("This symbol can not use for a class or script file. -- %s", []LnsAny{className}))
    }
    var info *Ast_NormalTypeInfo
    info = NewAst_NormalTypeInfo(_env, self, false, scope, nil, nil, false, externalFlag, false, accessMode, className, parentInfo, Ast_TypeInfoKind__ExtModule, nil, nil, nil, Ast_MutMode__Mut, moduleLang, Ast_Async__Noasync)
    info.FP.set_requirePath(_env, requirePath)
    return &info.Ast_TypeInfo
}

// 4963: decl @lune.@base.@Ast.ProcessInfo.createFunc
func (self *Ast_ProcessInfo) CreateFunc(_env *LnsEnv, abstractFlag bool,builtinFlag bool,scope LnsAny,kind LnsInt,parentInfo *Ast_TypeInfo,autoFlag bool,externalFlag bool,staticFlag bool,accessMode LnsInt,funcName string,asyncMode LnsInt,altTypeList LnsAny,argTypeList LnsAny,retTypeInfoList LnsAny,mutable LnsAny) *Ast_NormalTypeInfo {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(builtinFlag)) &&
        _env.SetStackVal( Parser_isLuaKeyword(_env, funcName)) ).(bool)){
        Util_err(_env, _env.LuaVM.String_format("This symbol can not use for a function. -- %s", []LnsAny{funcName}))
    }
    var info *Ast_NormalTypeInfo
    info = NewAst_NormalTypeInfo(_env, self, abstractFlag, scope, nil, nil, autoFlag, externalFlag, staticFlag, accessMode, funcName, parentInfo, kind, Lns_unwrapDefault( altTypeList, NewLnsList([]LnsAny{})).(*LnsList), Lns_unwrapDefault( argTypeList, NewLnsList([]LnsAny{})).(*LnsList), Lns_unwrapDefault( retTypeInfoList, NewLnsList([]LnsAny{})).(*LnsList), _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mutable) &&
        _env.SetStackVal( Ast_MutMode__Mut) ||
        _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt), nil, asyncMode)
    if altTypeList != nil{
        altTypeList_3899 := altTypeList.(*LnsList)
        for _, _genType := range( altTypeList_3899.Items ) {
            genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            {
                __exp := Ast_AlternateTypeInfoDownCastF(genType.FP)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Ast_AlternateTypeInfo)
                    _exp.FP.updateParentInfo(_env, &info.Ast_TypeInfo)
                }
            }
        }
    }
    return info
}

// 4993: decl @lune.@base.@Ast.ProcessInfo.createAdvertiseMethodFrom
func (self *Ast_ProcessInfo) CreateAdvertiseMethodFrom(_env *LnsEnv, classTypeInfo *Ast_TypeInfo,typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    return &self.FP.CreateFunc(_env, false, false, Ast_getScope(_env, typeInfo), typeInfo.FP.Get_kind(_env), classTypeInfo, true, false, false, typeInfo.FP.Get_accessMode(_env), typeInfo.FP.Get_rawTxt(_env), typeInfo.FP.Get_asyncMode(_env), typeInfo.FP.Get_itemTypeInfoList(_env), typeInfo.FP.Get_argTypeInfoList(_env), typeInfo.FP.Get_retTypeInfoList(_env), Ast_TypeInfo_isMut(_env, typeInfo)).Ast_TypeInfo
}

// 5022: decl @lune.@base.@Ast.ProcessInfo.createAlias
func (self *Ast_ProcessInfo) CreateAlias(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,externalFlag bool,accessMode LnsInt,parentInfo *Ast_TypeInfo,typeInfo *Ast_TypeInfo) *Ast_AliasTypeInfo {
    return NewAst_AliasTypeInfo(_env, processInfo, name, accessMode, parentInfo, typeInfo.FP.Get_srcTypeInfo(_env), externalFlag)
}

// 5179: decl @lune.@base.@Ast.ProcessInfo.createDDD
func (self *Ast_ProcessInfo) CreateDDD(_env *LnsEnv, typeInfo *Ast_TypeInfo,externalFlag bool,extTypeFlag bool) *Ast_DDDTypeInfo {
    if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        typeInfo = typeInfo.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(Lns_car(Ast_failCreateLuavalWith_3312_(_env, typeInfo, Ast_LuavalConvKind__InLua, true)))) &&
        _env.SetStackVal( extTypeFlag) ).(bool)){
        extTypeFlag = false
        
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_nonnilableType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Ext) &&
        _env.SetStackVal( extTypeFlag) ).(bool)){
        switch _exp20414 := self.FP.CreateLuaval(_env, typeInfo, true).(type) {
        case *Ast_LuavalResult__OK:
        work := _exp20414.Val1
            typeInfo = work
            
        case *Ast_LuavalResult__Err:
        mess := _exp20414.Val1
            Util_err(_env, mess)
        }
    }
    var dddMap *LnsMap
    if extTypeFlag{
        dddMap = self.FP.get_typeInfo2Map(_env).ExtDDDMap
        
    } else { 
        dddMap = self.FP.get_typeInfo2Map(_env).DDDMap
        
    }
    {
        __exp := dddMap.Get(typeInfo)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_DDDTypeInfo)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( _exp.FP.Get_typeId(_env).Id < Ast_userStartId) &&
                _env.SetStackVal( typeInfo.FP.Get_typeId(_env).Id >= Ast_userStartId) ).(bool)){
                Util_err(_env, "on cache")
            }
            return _exp
        }
    }
    var dddType *Ast_DDDTypeInfo
    dddType = NewAst_DDDTypeInfo(_env, self, typeInfo, externalFlag, nil)
    if Lns_isCondTrue( Lns_car(Ast_failCreateLuavalWith_3312_(_env, typeInfo, Ast_LuavalConvKind__InLua, true))){
        var extDDDType *Ast_DDDTypeInfo
        extDDDType = NewAst_DDDTypeInfo(_env, self, typeInfo, externalFlag, dddType)
        if extTypeFlag{
            return extDDDType
        }
    }
    return dddType
}

// 5580: decl @lune.@base.@Ast.ProcessInfo.createGeneric
func (self *Ast_ProcessInfo) CreateGeneric(_env *LnsEnv, genSrcTypeInfo *Ast_TypeInfo,itemTypeInfoList *LnsList,moduleTypeInfo *Ast_TypeInfo) *Ast_GenericTypeInfo {
    return NewAst_GenericTypeInfo(_env, self, genSrcTypeInfo, itemTypeInfoList, moduleTypeInfo)
}

// 5847: decl @lune.@base.@Ast.ProcessInfo.createLuaval
func (self *Ast_ProcessInfo) CreateLuaval(_env *LnsEnv, luneType *Ast_TypeInfo,validToCheck bool) LnsAny {
    if Lns_op_not(self.validExtType){
        return &Ast_LuavalResult__OK{luneType, true}
    }
    if luneType.FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
        return &Ast_LuavalResult__OK{luneType, true}
    }
    if Ast_isExtType(_env, luneType){
        return &Ast_LuavalResult__OK{luneType, true}
    }
    {
        _extType := self.FP.get_typeInfo2Map(_env).ExtMap.Get(luneType)
        if !Lns_IsNil( _extType ) {
            extType := _extType.(*Ast_TypeInfo)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( extType.FP.Get_typeId(_env).Id < Ast_userStartId) &&
                _env.SetStackVal( luneType.FP.Get_typeId(_env).Id >= Ast_userStartId) ).(bool)){
                Util_err(_env, "on cache")
            }
            if extType.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                return &Ast_LuavalResult__OK{extType, false}
            }
            return &Ast_LuavalResult__OK{extType, true}
        }
    }
    var process func(_env *LnsEnv) LnsAny
    process = func(_env *LnsEnv) LnsAny {
        {
            _dddType := Ast_DDDTypeInfoDownCastF(luneType.FP)
            if !Lns_IsNil( _dddType ) {
                dddType := _dddType.(*Ast_DDDTypeInfo)
                switch _exp23347 := self.FP.CreateLuaval(_env, dddType.FP.Get_typeInfo(_env), validToCheck).(type) {
                case *Ast_LuavalResult__Err:
                mess := _exp23347.Val1
                    Util_err(_env, mess)
                case *Ast_LuavalResult__OK:
                workType := _exp23347.Val1
                    return &Ast_LuavalResult__OK{&self.FP.CreateDDD(_env, workType, dddType.FP.Get_externalFlag(_env), true).Ast_TypeInfo, false}
                }
            }
        }
        var err LnsAny
        var canConv bool
        err, canConv = Ast_failCreateLuavalWith_3312_(_env, luneType, Ast_LuavalConvKind__InLua, validToCheck)
        
        if err != nil{
            err_4600 := err.(string)
            return &Ast_LuavalResult__Err{err_4600}
        }
        if canConv{
            return &Ast_LuavalResult__OK{luneType, true}
        }
        var extType *Ast_ExtTypeInfo
        extType = NewAst_ExtTypeInfo(_env, self, luneType.FP.Get_nonnilableType(_env))
        if luneType.FP.Get_nilable(_env){
            return &Ast_LuavalResult__OK{extType.FP.Get_nilableTypeInfo(_env), false}
        }
        return &Ast_LuavalResult__OK{&extType.Ast_TypeInfo, false}
    }
    var result LnsAny
    result = process(_env)
    switch _exp23480 := result.(type) {
    case *Ast_LuavalResult__OK:
    typeInfo := _exp23480.Val1
        self.FP.get_typeInfo2Map(_env).ExtMap.Set(luneType.FP.Get_nonnilableType(_env),typeInfo.FP.Get_nonnilableType(_env))
        self.FP.get_typeInfo2Map(_env).ExtMap.Set(luneType.FP.Get_nilableTypeInfo(_env),typeInfo.FP.Get_nilableTypeInfo(_env))
    }
    return result
}

// 5975: decl @lune.@base.@Ast.ProcessInfo.createEnum
func (self *Ast_ProcessInfo) CreateEnum(_env *LnsEnv, scope *Ast_Scope,parentInfo *Ast_TypeInfo,externalFlag bool,accessMode LnsInt,enumName string,valTypeInfo *Ast_TypeInfo) *Ast_EnumTypeInfo {
    if Parser_isLuaKeyword(_env, enumName){
        Util_err(_env, _env.LuaVM.String_format("This symbol can not use for a enum. -- %s", []LnsAny{enumName}))
    }
    var info *Ast_EnumTypeInfo
    info = NewAst_EnumTypeInfo(_env, self, scope, externalFlag, accessMode, enumName, parentInfo, valTypeInfo)
    var getEnumName *Ast_NormalTypeInfo
    getEnumName = self.FP.CreateFunc(_env, false, true, nil, Ast_TypeInfoKind__Method, &info.Ast_TypeInfo, true, externalFlag, false, Ast_AccessMode__Pub, "get__txt", Ast_Async__Async, nil, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), false)
    scope.FP.AddMethod(_env, self, nil, &getEnumName.Ast_TypeInfo, Ast_AccessMode__Pub, false, false)
    var fromVal *Ast_NormalTypeInfo
    fromVal = self.FP.CreateFunc(_env, false, true, nil, Ast_TypeInfoKind__Func, &info.Ast_TypeInfo, true, externalFlag, true, Ast_AccessMode__Pub, "_from", Ast_Async__Async, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.FP.CreateModifier(_env, valTypeInfo, Ast_MutMode__IMut))}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(info.FP.Get_nilableTypeInfo(_env))}), false)
    scope.FP.AddMethod(_env, self, nil, &fromVal.Ast_TypeInfo, Ast_AccessMode__Pub, true, false)
    var allListType *Ast_TypeInfo
    allListType = self.FP.CreateList(_env, Ast_AccessMode__Pub, &info.Ast_TypeInfo, NewLnsList([]LnsAny{Ast_EnumTypeInfo2Stem(info)}), Ast_MutMode__IMut)
    var allList *Ast_NormalTypeInfo
    allList = self.FP.CreateFunc(_env, false, true, nil, Ast_TypeInfoKind__Func, &info.Ast_TypeInfo, true, externalFlag, true, Ast_AccessMode__Pub, "get__allList", Ast_Async__Async, nil, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.FP.CreateModifier(_env, allListType, Ast_MutMode__IMut))}), false)
    scope.FP.AddMethod(_env, self, nil, &allList.Ast_TypeInfo, Ast_AccessMode__Pub, true, false)
    return info
}

// 6043: decl @lune.@base.@Ast.ProcessInfo.createAlge
func (self *Ast_ProcessInfo) CreateAlge(_env *LnsEnv, scope *Ast_Scope,parentInfo *Ast_TypeInfo,externalFlag bool,accessMode LnsInt,algeName string) *Ast_AlgeTypeInfo {
    if Parser_isLuaKeyword(_env, algeName){
        Util_err(_env, _env.LuaVM.String_format("This symbol can not use for a alge. -- %s", []LnsAny{algeName}))
    }
    var info *Ast_AlgeTypeInfo
    info = NewAst_AlgeTypeInfo(_env, self, scope, externalFlag, accessMode, algeName, parentInfo)
    var getAlgeName *Ast_NormalTypeInfo
    getAlgeName = self.FP.CreateFunc(_env, false, true, nil, Ast_TypeInfoKind__Method, &info.Ast_TypeInfo, true, externalFlag, false, Ast_AccessMode__Pub, "get__txt", Ast_Async__Async, nil, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), false)
    scope.FP.AddMethod(_env, self, nil, &getAlgeName.Ast_TypeInfo, Ast_AccessMode__Pub, false, false)
    return info
}

// 7102: decl @lune.@base.@Ast.ProcessInfo.newUser
func (self *Ast_ProcessInfo) NewUser(_env *LnsEnv) *Ast_ProcessInfo {
    return Ast_ProcessInfo_createUser_1078_(_env, self.validCheckingMutable, self.validExtType, self.validDetailError, Ast_builtinTypeInfo2Map.FP.Clone(_env))
}


// declaration Class -- IdInfo
type Ast_IdInfoMtd interface {
    Equals(_env *LnsEnv, arg1 *Ast_IdInfo) bool
    Get_orgId(_env *LnsEnv) LnsInt
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    IsSwichingId(_env *LnsEnv) bool
    Set_orgId(_env *LnsEnv, arg1 LnsInt)
}
type Ast_IdInfo struct {
    Id LnsInt
    ProcessInfo *Ast_ProcessInfo
    orgId LnsAny
    FP Ast_IdInfoMtd
}
func Ast_IdInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_IdInfo).FP
}
type Ast_IdInfoDownCast interface {
    ToAst_IdInfo() *Ast_IdInfo
}
func Ast_IdInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_IdInfoDownCast)
    if ok { return work.ToAst_IdInfo() }
    return nil
}
func (obj *Ast_IdInfo) ToAst_IdInfo() *Ast_IdInfo {
    return obj
}
func NewAst_IdInfo(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_ProcessInfo) *Ast_IdInfo {
    obj := &Ast_IdInfo{}
    obj.FP = obj
    obj.InitAst_IdInfo(_env, arg1, arg2)
    return obj
}
func (self *Ast_IdInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo{ return self.ProcessInfo }
// 174: decl @lune.@base.@Ast.IdInfo.isSwichingId
func (self *Ast_IdInfo) IsSwichingId(_env *LnsEnv) bool {
    var orgId LnsInt
    
    {
        _orgId := self.orgId
        if _orgId == nil{
            return false
        } else {
            orgId = _orgId.(LnsInt)
        }
    }
    return orgId != self.Id
}

// 181: decl @lune.@base.@Ast.IdInfo.set_orgId
func (self *Ast_IdInfo) Set_orgId(_env *LnsEnv, id LnsInt) {
    self.orgId = id
    
}

// 185: decl @lune.@base.@Ast.IdInfo.get_orgId
func (self *Ast_IdInfo) Get_orgId(_env *LnsEnv) LnsInt {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.orgId) ||
        _env.SetStackVal( self.Id) ).(LnsInt)
}

// 188: DeclConstr
func (self *Ast_IdInfo) InitAst_IdInfo(_env *LnsEnv, id LnsInt,processInfo *Ast_ProcessInfo) {
    self.Id = id
    
    self.ProcessInfo = processInfo
    
    self.orgId = nil
    
}

// 194: decl @lune.@base.@Ast.IdInfo.equals
func (self *Ast_IdInfo) Equals(_env *LnsEnv, idInfo *Ast_IdInfo) bool {
    if self.FP.Get_orgId(_env) == idInfo.FP.Get_orgId(_env){
        if self.ProcessInfo == idInfo.ProcessInfo{
            return true
        }
    }
    return false
}


type Ast_ModuleInfoIF interface {
        Get_assignName(_env *LnsEnv) string
        Get_modulePath(_env *LnsEnv) string
}
func Lns_cast2Ast_ModuleInfoIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ast_ModuleInfoIF); ok { 
        return obj
    }
    return nil
}

// declaration Class -- SimpleModuleInfo
type Ast_SimpleModuleInfoMtd interface {
    Get_assignName(_env *LnsEnv) string
    Get_modulePath(_env *LnsEnv) string
}
type Ast_SimpleModuleInfo struct {
    assignName string
    modulePath string
    FP Ast_SimpleModuleInfoMtd
}
func Ast_SimpleModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_SimpleModuleInfo).FP
}
type Ast_SimpleModuleInfoDownCast interface {
    ToAst_SimpleModuleInfo() *Ast_SimpleModuleInfo
}
func Ast_SimpleModuleInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_SimpleModuleInfoDownCast)
    if ok { return work.ToAst_SimpleModuleInfo() }
    return nil
}
func (obj *Ast_SimpleModuleInfo) ToAst_SimpleModuleInfo() *Ast_SimpleModuleInfo {
    return obj
}
func NewAst_SimpleModuleInfo(_env *LnsEnv, arg1 string, arg2 string) *Ast_SimpleModuleInfo {
    obj := &Ast_SimpleModuleInfo{}
    obj.FP = obj
    obj.InitAst_SimpleModuleInfo(_env, arg1, arg2)
    return obj
}
func (self *Ast_SimpleModuleInfo) InitAst_SimpleModuleInfo(_env *LnsEnv, arg1 string, arg2 string) {
    self.assignName = arg1
    self.modulePath = arg2
}
func (self *Ast_SimpleModuleInfo) Get_assignName(_env *LnsEnv) string{ return self.assignName }
func (self *Ast_SimpleModuleInfo) Get_modulePath(_env *LnsEnv) string{ return self.modulePath }

type Ast_ModuleInfoManager interface {
        GetModuleInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny
}
func Lns_cast2Ast_ModuleInfoManager( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ast_ModuleInfoManager); ok { 
        return obj
    }
    return nil
}

// declaration Class -- DummyModuleInfoManager
var Ast_DummyModuleInfoManager__instance *Ast_DummyModuleInfoManager
// 253: decl @lune.@base.@Ast.DummyModuleInfoManager.___init
func Ast_DummyModuleInfoManager____init_1155_(_env *LnsEnv) {
    Ast_DummyModuleInfoManager__instance = NewAst_DummyModuleInfoManager(_env)
    
}

type Ast_DummyModuleInfoManagerMtd interface {
    GetModuleInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny
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
func NewAst_DummyModuleInfoManager(_env *LnsEnv) *Ast_DummyModuleInfoManager {
    obj := &Ast_DummyModuleInfoManager{}
    obj.FP = obj
    obj.InitAst_DummyModuleInfoManager(_env)
    return obj
}
func Ast_DummyModuleInfoManager_get_instance(_env *LnsEnv) *Ast_DummyModuleInfoManager{ return Ast_DummyModuleInfoManager__instance }
// 255: DeclConstr
func (self *Ast_DummyModuleInfoManager) InitAst_DummyModuleInfoManager(_env *LnsEnv) {
}

// 260: decl @lune.@base.@Ast.DummyModuleInfoManager.getModuleInfo
func (self *Ast_DummyModuleInfoManager) GetModuleInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny {
    return nil
}


type Ast_LowSymbol interface {
        Get_accessMode(_env *LnsEnv) LnsInt
        Get_convModuleParam(_env *LnsEnv) LnsAny
        Get_hasAccessFromClosure(_env *LnsEnv) bool
        Get_isLazyLoad(_env *LnsEnv) bool
        Get_kind(_env *LnsEnv) LnsInt
        Get_mutable(_env *LnsEnv) bool
        Get_name(_env *LnsEnv) string
        Get_scope(_env *LnsEnv) *Ast_Scope
        Get_staticFlag(_env *LnsEnv) bool
        Get_symbolId(_env *LnsEnv) LnsInt
        Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
}
func Lns_cast2Ast_LowSymbol( obj LnsAny ) LnsAny {
    if _, ok := obj.(Ast_LowSymbol); ok { 
        return obj
    }
    return nil
}

// declaration Class -- TypeNameCtrl
type Ast_TypeNameCtrlMtd interface {
    GetModuleName(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 string, arg3 Ast_ModuleInfoManager) string
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsAny, arg3 LnsAny) string
    Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Set_moduleTypeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo)
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
func NewAst_TypeNameCtrl(_env *LnsEnv, arg1 *Ast_TypeInfo) *Ast_TypeNameCtrl {
    obj := &Ast_TypeNameCtrl{}
    obj.FP = obj
    obj.InitAst_TypeNameCtrl(_env, arg1)
    return obj
}
func (self *Ast_TypeNameCtrl) InitAst_TypeNameCtrl(_env *LnsEnv, arg1 *Ast_TypeInfo) {
    self.moduleTypeInfo = arg1
}
func (self *Ast_TypeNameCtrl) Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.moduleTypeInfo }
func (self *Ast_TypeNameCtrl) Set_moduleTypeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo){ self.moduleTypeInfo = arg1 }

// 1154: decl @lune.@base.@Ast.TypeNameCtrl.getModuleName
func (self *Ast_TypeNameCtrl) GetModuleName(_env *LnsEnv, workTypeInfo *Ast_TypeInfo,name string,moduleInfoMan Ast_ModuleInfoManager) string {
    {
        _moduleInfo := moduleInfoMan.GetModuleInfo(_env, workTypeInfo)
        if !Lns_IsNil( _moduleInfo ) {
            moduleInfo := _moduleInfo.(Ast_ModuleInfoIF)
            var txt string
            txt = moduleInfo.Get_assignName(_env)
            return txt + "." + name
        } else {
            if self.moduleTypeInfo != workTypeInfo{
                return workTypeInfo.FP.Get_rawTxt(_env) + "." + name
            }
        }
    }
    return name
}

// 1169: decl @lune.@base.@Ast.TypeNameCtrl.getParentFullName
func (self *Ast_TypeNameCtrl) GetParentFullName(_env *LnsEnv, typeInfo *Ast_TypeInfo,importInfo LnsAny,localFlag LnsAny) string {
    var workTypeInfo *Ast_TypeInfo
    workTypeInfo = typeInfo
    var name string
    name = ""
    var moduleInfoMan Ast_ModuleInfoManager
    
    {
        _moduleInfoMan := importInfo
        if _moduleInfoMan == nil{
            moduleInfoMan = Ast_DummyModuleInfoManager_get_instance(_env).FP
            
        } else {
            moduleInfoMan = _moduleInfoMan.(Ast_ModuleInfoManager)
        }
    }
    for  {
        workTypeInfo = workTypeInfo.FP.Get_parentInfo(_env)
        
        var txt string
        txt = workTypeInfo.FP.Get_rawTxt(_env)
        if workTypeInfo == workTypeInfo.FP.Get_parentInfo(_env){
            break
        }
        if Lns_isCondTrue( localFlag){
            if workTypeInfo.FP.IsModule(_env){
                name = self.FP.GetModuleName(_env, workTypeInfo, name, moduleInfoMan)
                
                break
            }
        }
        name = txt + "." + name
        
    }
    return name
}


// declaration Class -- SymbolInfo
type Ast_SymbolInfoMtd interface {
    CanAccess(_env *LnsEnv, arg1 *Ast_Scope, arg2 LnsInt) LnsAny
    ClearValue(_env *LnsEnv)
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOrg(_env *LnsEnv) *Ast_SymbolInfo
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_canBeLeft(_env *LnsEnv) bool
    Get_canBeRight(_env *LnsEnv) bool
    Get_convModuleParam(_env *LnsEnv) LnsAny
    Get_hasAccessFromClosure(_env *LnsEnv) bool
    Get_hasValueFlag(_env *LnsEnv) bool
    Get_isLazyLoad(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_mutable(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) string
    Get_namespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_pos(_env *LnsEnv) LnsAny
    Get_posForLatestMod(_env *LnsEnv) LnsAny
    Get_posForModToRef(_env *LnsEnv) LnsAny
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_staticFlag(_env *LnsEnv) bool
    Get_symbolId(_env *LnsEnv) LnsInt
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    HasAccess(_env *LnsEnv) bool
    Set_convModuleParam(_env *LnsEnv, arg1 LnsAny)
    Set_hasAccessFromClosure(_env *LnsEnv, arg1 bool)
    Set_hasValueFlag(_env *LnsEnv, arg1 bool)
    set_namespaceTypeInfo(_env *LnsEnv, arg1 LnsAny)
    Set_posForLatestMod(_env *LnsEnv, arg1 LnsAny)
    Set_posForModToRef(_env *LnsEnv, arg1 LnsAny)
    Set_typeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo)
    UpdateValue(_env *LnsEnv, arg1 LnsAny)
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
func (self *Ast_SymbolInfo) set_namespaceTypeInfo(_env *LnsEnv, arg1 LnsAny){ self.namespaceTypeInfo = arg1 }
func (self *Ast_SymbolInfo) Get_accessMode(_env *LnsEnv) LnsInt{
// insert a dummy
    return 0
}
func (self *Ast_SymbolInfo) Get_convModuleParam(_env *LnsEnv) LnsAny{
// insert a dummy
    return nil
}
func (self *Ast_SymbolInfo) Get_hasAccessFromClosure(_env *LnsEnv) bool{
// insert a dummy
    return false
}
func (self *Ast_SymbolInfo) Get_isLazyLoad(_env *LnsEnv) bool{
// insert a dummy
    return false
}
func (self *Ast_SymbolInfo) Get_kind(_env *LnsEnv) LnsInt{
// insert a dummy
    return 0
}
func (self *Ast_SymbolInfo) Get_mutable(_env *LnsEnv) bool{
// insert a dummy
    return false
}
func (self *Ast_SymbolInfo) Get_name(_env *LnsEnv) string{
// insert a dummy
    return ""
}
func (self *Ast_SymbolInfo) Get_scope(_env *LnsEnv) *Ast_Scope{
// insert a dummy
    return nil
}
func (self *Ast_SymbolInfo) Get_staticFlag(_env *LnsEnv) bool{
// insert a dummy
    return false
}
func (self *Ast_SymbolInfo) Get_symbolId(_env *LnsEnv) LnsInt{
// insert a dummy
    return 0
}
func (self *Ast_SymbolInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{
// insert a dummy
    return nil
}
// 435: DeclConstr
func (self *Ast_SymbolInfo) InitAst_SymbolInfo(_env *LnsEnv) {
    self.namespaceTypeInfo = nil
    
}























// 483: decl @lune.@base.@Ast.SymbolInfo.hasAccess
func (self *Ast_SymbolInfo) HasAccess(_env *LnsEnv) bool {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.Get_posForModToRef(_env)) ||
        _env.SetStackVal( self.FP.Get_posForLatestMod(_env) != self.FP.Get_pos(_env)) )){
        return true
    }
    return false
}

// 490: decl @lune.@base.@Ast.SymbolInfo.updateValue
func (self *Ast_SymbolInfo) UpdateValue(_env *LnsEnv, pos LnsAny) {
    self.FP.Set_hasValueFlag(_env, true)
    self.FP.Set_posForLatestMod(_env, pos)
}

// 494: decl @lune.@base.@Ast.SymbolInfo.clearValue
func (self *Ast_SymbolInfo) ClearValue(_env *LnsEnv) {
    self.FP.Set_hasValueFlag(_env, false)
}






// 634: decl @lune.@base.@Ast.SymbolInfo.get_namespaceTypeInfo
func (self *Ast_SymbolInfo) Get_namespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    {
        __exp := self.namespaceTypeInfo
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            return _exp
        }
    }
    var work *Ast_TypeInfo
    work = self.FP.Get_scope(_env).FP.GetNamespaceTypeInfo(_env)
    self.namespaceTypeInfo = work
    
    return work
}

// 1097: decl @lune.@base.@Ast.SymbolInfo.getModule
func (self *Ast_SymbolInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.FP.Get_namespaceTypeInfo(_env).FP.GetModule(_env)
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
func NewAst_DataOwnerInfo(_env *LnsEnv, arg1 bool, arg2 *Ast_SymbolInfo) *Ast_DataOwnerInfo {
    obj := &Ast_DataOwnerInfo{}
    obj.FP = obj
    obj.InitAst_DataOwnerInfo(_env, arg1, arg2)
    return obj
}
func (self *Ast_DataOwnerInfo) InitAst_DataOwnerInfo(_env *LnsEnv, arg1 bool, arg2 *Ast_SymbolInfo) {
    self.HasData = arg1
    self.SymbolInfo = arg2
}

// declaration Class -- Scope
type Ast_ScopeMtd interface {
    AccessSymbol(_env *LnsEnv, arg1 *Ast_Scope, arg2 *Ast_SymbolInfo)
    Add(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 bool, arg4 bool, arg5 string, arg6 LnsAny, arg7 *Ast_TypeInfo, arg8 LnsInt, arg9 bool, arg10 LnsInt, arg11 bool, arg12 bool)(LnsAny, LnsAny)
    AddAlge(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 string, arg4 LnsAny, arg5 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddAlgeVal(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddAlias(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 *Types_Position, arg4 bool, arg5 LnsInt, arg6 *Ast_TypeInfo, arg7 *Ast_SymbolInfo)(LnsAny, LnsAny)
    AddAliasForType(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddAlternate(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 string, arg4 LnsAny, arg5 *Ast_TypeInfo)
    AddClass(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddClassLazy(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 bool)(LnsAny, LnsAny)
    AddEnum(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 string, arg4 LnsAny, arg5 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddEnumVal(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo)(LnsAny, LnsAny)
    AddExtModule(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 bool, arg6 LnsInt)(LnsAny, LnsAny)
    AddForm(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *Ast_TypeInfo, arg4 LnsInt)(LnsAny, LnsAny)
    AddFunc(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *Ast_TypeInfo, arg4 LnsInt, arg5 bool, arg6 bool)(LnsAny, LnsAny)
    AddIgnoredVar(_env *LnsEnv, arg1 *Ast_ProcessInfo)
    AddLocalVar(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 bool, arg4 string, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 LnsInt)(LnsAny, LnsAny)
    AddMacro(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *Ast_TypeInfo, arg4 LnsInt)(LnsAny, LnsAny)
    AddMember(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 LnsInt, arg6 bool, arg7 LnsInt)(LnsAny, LnsAny)
    AddMethod(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *Ast_TypeInfo, arg4 LnsInt, arg5 bool, arg6 bool)(LnsAny, LnsAny)
    AddModule(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Ast_ModuleInfoIF)
    AddOverrideImut(_env *LnsEnv, arg1 *Ast_SymbolInfo)
    AddStaticVar(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 bool, arg4 string, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 LnsInt)(LnsAny, LnsAny)
    AddSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo)
    AddSymbolInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_SymbolInfo)(LnsAny, LnsAny)
    AddUnwrapedVar(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 bool, arg4 string, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 LnsInt)(LnsAny, LnsAny)
    AddVar(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 string, arg4 LnsAny, arg5 *Ast_TypeInfo, arg6 LnsInt, arg7 bool)(LnsAny, LnsAny)
    FilterSymbolInfoIfField(_env *LnsEnv, arg1 *Ast_Scope, arg2 LnsInt, arg3 Ast_filterForm) bool
    FilterSymbolTypeInfo(_env *LnsEnv, arg1 *Ast_Scope, arg2 *Ast_Scope, arg3 LnsInt, arg4 Ast_filterForm)
    FilterTypeInfoField(_env *LnsEnv, arg1 LnsAny, arg2 *Ast_Scope, arg3 LnsInt, arg4 Ast_filterForm) bool
    FilterTypeInfoFieldAndIF(_env *LnsEnv, arg1 LnsAny, arg2 *Ast_Scope, arg3 LnsInt, arg4 Ast_filterForm) bool
    GetClassTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetModuleInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny
    GetNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    GetProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetSymbolInfo(_env *LnsEnv, arg1 string, arg2 *Ast_Scope, arg3 bool, arg4 LnsInt) LnsAny
    GetSymbolInfoChild(_env *LnsEnv, arg1 string) LnsAny
    GetSymbolInfoField(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 *Ast_Scope, arg4 LnsInt) LnsAny
    GetSymbolInfoIfField(_env *LnsEnv, arg1 string, arg2 *Ast_Scope, arg3 LnsInt) LnsAny
    GetSymbolTypeInfo(_env *LnsEnv, arg1 string, arg2 *Ast_Scope, arg3 *Ast_Scope, arg4 LnsInt) LnsAny
    GetTypeInfo(_env *LnsEnv, arg1 string, arg2 *Ast_Scope, arg3 bool, arg4 LnsInt) LnsAny
    GetTypeInfoChild(_env *LnsEnv, arg1 string) LnsAny
    GetTypeInfoField(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 *Ast_Scope, arg4 LnsInt) LnsAny
    Get_closureSym2NumMap(_env *LnsEnv) *LnsMap
    Get_closureSymList(_env *LnsEnv) *LnsList
    Get_closureSymMap(_env *LnsEnv) *LnsMap
    Get_inherit(_env *LnsEnv) LnsAny
    Get_ownerTypeInfo(_env *LnsEnv) LnsAny
    Get_parent(_env *LnsEnv) *Ast_Scope
    Get_scopeId(_env *LnsEnv) LnsInt
    Get_symbol2SymbolInfoMap(_env *LnsEnv) *LnsMap
    Get_validCheckingUnaccess(_env *LnsEnv) bool
    IsClosureAccess(_env *LnsEnv, arg1 *Ast_Scope, arg2 *Ast_SymbolInfo) bool
    IsInnerOf(_env *LnsEnv, arg1 *Ast_Scope) bool
    IsRoot(_env *LnsEnv) bool
    Remove(_env *LnsEnv, arg1 string)
    SetClosure(_env *LnsEnv, arg1 *Ast_SymbolInfo)
    SetData(_env *LnsEnv, arg1 *Ast_SymbolInfo)
    Set_ownerTypeInfo(_env *LnsEnv, arg1 LnsAny)
    Set_validCheckingUnaccess(_env *LnsEnv, arg1 bool)
    SwitchOwnerTypeInfo(_env *LnsEnv, arg1 LnsAny)
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
func NewAst_Scope(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 bool, arg4 LnsAny, arg5 LnsAny) *Ast_Scope {
    obj := &Ast_Scope{}
    obj.FP = obj
    obj.InitAst_Scope(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Ast_Scope) Get_scopeId(_env *LnsEnv) LnsInt{ return self.scopeId }
func (self *Ast_Scope) Get_ownerTypeInfo(_env *LnsEnv) LnsAny{ return self.ownerTypeInfo }
func (self *Ast_Scope) Get_parent(_env *LnsEnv) *Ast_Scope{ return self.parent }
func (self *Ast_Scope) Get_symbol2SymbolInfoMap(_env *LnsEnv) *LnsMap{ return self.symbol2SymbolInfoMap }
func (self *Ast_Scope) Get_inherit(_env *LnsEnv) LnsAny{ return self.inherit }
func (self *Ast_Scope) Get_closureSymMap(_env *LnsEnv) *LnsMap{ return self.closureSymMap }
func (self *Ast_Scope) Get_closureSymList(_env *LnsEnv) *LnsList{ return self.closureSymList }
func (self *Ast_Scope) Get_closureSym2NumMap(_env *LnsEnv) *LnsMap{ return self.closureSym2NumMap }
func (self *Ast_Scope) Get_validCheckingUnaccess(_env *LnsEnv) bool{ return self.validCheckingUnaccess }
func (self *Ast_Scope) Set_validCheckingUnaccess(_env *LnsEnv, arg1 bool){ self.validCheckingUnaccess = arg1 }
// 558: DeclConstr
func (self *Ast_Scope) InitAst_Scope(_env *LnsEnv, processInfo *Ast_ProcessInfo,parent LnsAny,classFlag bool,inherit LnsAny,ifScopeList LnsAny) {
    self.scopeId = processInfo.FP.Get_idProvScope(_env).FP.GetNewId(_env)
    
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


// 581: decl @lune.@base.@Ast.Scope.isRoot
func (self *Ast_Scope) IsRoot(_env *LnsEnv) bool {
    return self.parent == self
}

// 585: decl @lune.@base.@Ast.Scope.set_ownerTypeInfo
func (self *Ast_Scope) Set_ownerTypeInfo(_env *LnsEnv, owner LnsAny) {
    if Lns_op_not(self.ownerTypeInfo){
        self.ownerTypeInfo = owner
        
    }
}

// 590: decl @lune.@base.@Ast.Scope.switchOwnerTypeInfo
func (self *Ast_Scope) SwitchOwnerTypeInfo(_env *LnsEnv, owner LnsAny) {
    self.ownerTypeInfo = owner
    
}

// 595: decl @lune.@base.@Ast.Scope.getTypeInfoChild
func (self *Ast_Scope) GetTypeInfoChild(_env *LnsEnv, name string) LnsAny {
    {
        __exp := self.symbol2SymbolInfoMap.Get(name)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_typeInfo(_env)
        }
    }
    return nil
}

// 602: decl @lune.@base.@Ast.Scope.getSymbolInfoChild
func (self *Ast_Scope) GetSymbolInfoChild(_env *LnsEnv, name string) LnsAny {
    return self.symbol2SymbolInfoMap.Get(name)
}

// 606: decl @lune.@base.@Ast.Scope.setData
func (self *Ast_Scope) SetData(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) {
    self.symbolId2DataOwnerInfo.Set(symbolInfo.FP.Get_symbolId(_env),NewAst_DataOwnerInfo(_env, true, symbolInfo))
}

// 611: decl @lune.@base.@Ast.Scope.remove
func (self *Ast_Scope) Remove(_env *LnsEnv, name string) {
    self.symbol2SymbolInfoMap.Set(name,nil)
}

// 615: decl @lune.@base.@Ast.Scope.addSymbol
func (self *Ast_Scope) AddSymbol(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) {
    self.symbol2SymbolInfoMap.Set(symbolInfo.FP.Get_name(_env),symbolInfo)
}

// 619: decl @lune.@base.@Ast.Scope.addModule
func (self *Ast_Scope) AddModule(_env *LnsEnv, typeInfo *Ast_TypeInfo,moduleInfo Ast_ModuleInfoIF) {
    self.typeInfo2ModuleInfoMap.Set(typeInfo,moduleInfo)
}

// 623: decl @lune.@base.@Ast.Scope.getModuleInfo
func (self *Ast_Scope) GetModuleInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny {
    {
        _moduleInfo := self.typeInfo2ModuleInfoMap.Get(typeInfo)
        if !Lns_IsNil( _moduleInfo ) {
            moduleInfo := _moduleInfo.(Ast_ModuleInfoIF)
            return moduleInfo
        }
    }
    if self.parent != self{
        return self.parent.FP.GetModuleInfo(_env, typeInfo)
    }
    return nil
}

// 653: decl @lune.@base.@Ast.Scope.isInnerOf
func (self *Ast_Scope) IsInnerOf(_env *LnsEnv, scope *Ast_Scope) bool {
    var workScope *Ast_Scope
    workScope = self
    for workScope != Ast_rootScopeRo {
        if workScope == scope{
            return true
        }
        workScope = workScope.parent
        
    }
    return false
}

// 1263: decl @lune.@base.@Ast.Scope.getNamespaceTypeInfo
func (self *Ast_Scope) GetNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var scope *Ast_Scope
    scope = self
    for {
        {
            __exp := scope.FP.Get_ownerTypeInfo(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_TypeInfo)
                return _exp
            }
        }
        scope = scope.FP.Get_parent(_env)
        
        if scope.FP.IsRoot(_env){ break }
    }
    return typeInfo
}

// 1275: decl @lune.@base.@Ast.Scope.getModule
func (self *Ast_Scope) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.FP.GetNamespaceTypeInfo(_env).FP.GetModule(_env)
}

// 1279: decl @lune.@base.@Ast.Scope.getProcessInfo
func (self *Ast_Scope) GetProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.FP.GetModule(_env).FP.getProcessInfo(_env)
}

// 1761: decl @lune.@base.@Ast.Scope.filterTypeInfoField
func (self *Ast_Scope) FilterTypeInfoField(_env *LnsEnv, includeSelfFlag LnsAny,fromScope *Ast_Scope,access LnsInt,callback Ast_filterForm) bool {
    if self.classFlag{
        if Lns_isCondTrue( includeSelfFlag){
            {
                __collection5335 := self.symbol2SymbolInfoMap
                __sorted5335 := __collection5335.CreateKeyListStr()
                __sorted5335.Sort( LnsItemKindStr, nil )
                for _, ___key5335 := range( __sorted5335.Items ) {
                    symbolInfo := __collection5335.Items[ ___key5335 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    if Lns_isCondTrue( symbolInfo.FP.CanAccess(_env, fromScope, access)){
                        if Lns_op_not(callback(_env, symbolInfo)){
                            return false
                        }
                    }
                }
            }
        }
        {
            _scope := self.inherit
            if !Lns_IsNil( _scope ) {
                scope := _scope.(*Ast_Scope)
                if Lns_op_not(scope.FP.FilterTypeInfoField(_env, true, fromScope, access, callback)){
                    return false
                }
            }
        }
    }
    return true
}

// 1790: decl @lune.@base.@Ast.Scope.getSymbolInfoField
func (self *Ast_Scope) GetSymbolInfoField(_env *LnsEnv, name string,includeSelfFlag LnsAny,fromScope *Ast_Scope,access LnsInt) LnsAny {
    if self.classFlag{
        if Lns_isCondTrue( includeSelfFlag){
            {
                __exp := self.symbol2SymbolInfoMap.Get(name)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Ast_SymbolInfo)
                    var symbolInfo *Ast_SymbolInfo
                    
                    {
                        _symbolInfo := _exp.FP.CanAccess(_env, fromScope, access)
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
            if !Lns_IsNil( _scope ) {
                scope := _scope.(*Ast_Scope)
                var symbolInfo LnsAny
                symbolInfo = scope.FP.GetSymbolInfoField(_env, name, true, fromScope, access)
                if Lns_isCondTrue( symbolInfo){
                    return symbolInfo
                }
            }
        }
    }
    return nil
}

// 1820: decl @lune.@base.@Ast.Scope.getSymbolInfoIfField
func (self *Ast_Scope) GetSymbolInfoIfField(_env *LnsEnv, name string,fromScope *Ast_Scope,access LnsInt) LnsAny {
    if self.classFlag{
        for _, _scope := range( self.ifScopeList.Items ) {
            scope := _scope.(Ast_ScopeDownCast).ToAst_Scope()
            {
                _symbolInfo := scope.FP.GetSymbolInfoField(_env, name, true, fromScope, access)
                if !Lns_IsNil( _symbolInfo ) {
                    symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                    return symbolInfo
                }
            }
        }
    }
    {
        _scope := self.inherit
        if !Lns_IsNil( _scope ) {
            scope := _scope.(*Ast_Scope)
            {
                _symbolInfo := scope.FP.GetSymbolInfoIfField(_env, name, fromScope, access)
                if !Lns_IsNil( _symbolInfo ) {
                    symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                    return symbolInfo
                }
            }
        }
    }
    return nil
}

// 1847: decl @lune.@base.@Ast.Scope.filterSymbolInfoIfField
func (self *Ast_Scope) FilterSymbolInfoIfField(_env *LnsEnv, fromScope *Ast_Scope,access LnsInt,callback Ast_filterForm) bool {
    for _, _scope := range( self.ifScopeList.Items ) {
        scope := _scope.(Ast_ScopeDownCast).ToAst_Scope()
        if Lns_op_not(scope.FP.FilterTypeInfoField(_env, true, fromScope, access, callback)){
            return false
        }
    }
    {
        _scope := self.inherit
        if !Lns_IsNil( _scope ) {
            scope := _scope.(*Ast_Scope)
            if Lns_op_not(scope.FP.FilterSymbolInfoIfField(_env, fromScope, access, callback)){
                return false
            }
        }
    }
    return true
}

// 1869: decl @lune.@base.@Ast.Scope.getTypeInfoField
func (self *Ast_Scope) GetTypeInfoField(_env *LnsEnv, name string,includeSelfFlag LnsAny,fromScope *Ast_Scope,access LnsInt) LnsAny {
    var symbolInfo LnsAny
    symbolInfo = self.FP.GetSymbolInfoField(_env, name, includeSelfFlag, fromScope, access)
    {
        __exp := symbolInfo
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_typeInfo(_env)
        }
    }
    return nil
}

// 1893: decl @lune.@base.@Ast.Scope.filterTypeInfoFieldAndIF
func (self *Ast_Scope) FilterTypeInfoFieldAndIF(_env *LnsEnv, includeSelfFlag LnsAny,fromScope *Ast_Scope,access LnsInt,callback Ast_filterForm) bool {
    if self.classFlag{
        if Lns_isCondTrue( includeSelfFlag){
            {
                __collection5705 := self.symbol2SymbolInfoMap
                __sorted5705 := __collection5705.CreateKeyListStr()
                __sorted5705.Sort( LnsItemKindStr, nil )
                for _, ___key5705 := range( __sorted5705.Items ) {
                    symbolInfo := __collection5705.Items[ ___key5705 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    if Lns_isCondTrue( symbolInfo.FP.CanAccess(_env, fromScope, access)){
                        if Lns_op_not(callback(_env, symbolInfo)){
                            return false
                        }
                    }
                }
            }
        }
        {
            _scope := self.inherit
            if !Lns_IsNil( _scope ) {
                scope := _scope.(*Ast_Scope)
                if Lns_op_not(scope.FP.FilterTypeInfoField(_env, true, fromScope, access, callback)){
                    return false
                }
            }
        }
    }
    for _, _scope := range( self.ifScopeList.Items ) {
        scope := _scope.(Ast_ScopeDownCast).ToAst_Scope()
        if Lns_op_not(scope.FP.FilterTypeInfoField(_env, true, fromScope, access, callback)){
            return false
        }
    }
    {
        _scope := self.inherit
        if !Lns_IsNil( _scope ) {
            scope := _scope.(*Ast_Scope)
            if Lns_op_not(scope.FP.FilterSymbolInfoIfField(_env, fromScope, access, callback)){
                return false
            }
        }
    }
    return true
}

// 1941: decl @lune.@base.@Ast.Scope.getSymbolInfo
func (self *Ast_Scope) GetSymbolInfo(_env *LnsEnv, name string,fromScope *Ast_Scope,onlySameNsFlag bool,access LnsInt) LnsAny {
    {
        __exp := self.symbol2SymbolInfoMap.Get(name)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            var symbolInfo *Ast_SymbolInfo
            
            {
                _symbolInfo := _exp.FP.CanAccess(_env, fromScope, access)
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
            if !Lns_IsNil( _scope ) {
                scope := _scope.(*Ast_Scope)
                var symbolInfo LnsAny
                symbolInfo = scope.FP.GetSymbolInfoField(_env, name, true, fromScope, access)
                if Lns_isCondTrue( symbolInfo){
                    return symbolInfo
                }
            }
        }
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(onlySameNsFlag)) ||
        _env.SetStackVal( Lns_op_not(self.ownerTypeInfo)) ).(bool){
        if self.parent != self{
            return self.parent.FP.GetSymbolInfo(_env, name, fromScope, onlySameNsFlag, access)
        }
    } else { 
        var workScope *Ast_Scope
        workScope = self.parent
        for workScope.parent != workScope {
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(workScope.ownerTypeInfo) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) != Ast_TypeInfoKind__Class) &&
                _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(workScope.ownerTypeInfo) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) != Ast_TypeInfoKind__Module) ).(bool)){
                return workScope.FP.GetSymbolInfo(_env, name, fromScope, onlySameNsFlag, access)
            }
            workScope = workScope.parent
            
        }
    }
    {
        __exp := Ast_sym2builtInTypeMap.Get(name)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp
        }
    }
    return nil
}

// 2002: decl @lune.@base.@Ast.Scope.getTypeInfo
func (self *Ast_Scope) GetTypeInfo(_env *LnsEnv, name string,fromScope *Ast_Scope,onlySameNsFlag bool,access LnsInt) LnsAny {
    var symbolInfo *Ast_SymbolInfo
    
    {
        _symbolInfo := self.FP.GetSymbolInfo(_env, name, fromScope, onlySameNsFlag, access)
        if _symbolInfo == nil{
            return nil
        } else {
            symbolInfo = _symbolInfo.(*Ast_SymbolInfo)
        }
    }
    return symbolInfo.FP.Get_typeInfo(_env)
}

// 2019: decl @lune.@base.@Ast.Scope.getSymbolTypeInfo
func (self *Ast_Scope) GetSymbolTypeInfo(_env *LnsEnv, name string,fromScope *Ast_Scope,moduleScope *Ast_Scope,access LnsInt) LnsAny {
    var validThisScope bool
    validThisScope = false
    var limitSymbol bool
    limitSymbol = false
    {
        __exp := self.ownerTypeInfo
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__Func) ||
                _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) ||
                _env.SetStackVal( self == moduleScope) ||
                _env.SetStackVal( self == Ast_rootScopeRo) ).(bool){
                validThisScope = true
                
            } else if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ||
                _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__Module) ).(bool){
                limitSymbol = true
                
                validThisScope = true
                
            } else if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) ||
                _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__Alge) ).(bool){
                validThisScope = true
                
            }
        } else {
            validThisScope = true
            
        }
    }
    if validThisScope{
        {
            _symbolInfo := self.symbol2SymbolInfoMap.Get(name)
            if !Lns_IsNil( _symbolInfo ) {
                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( Lns_op_not(limitSymbol)) ||
                    _env.SetStackVal( name == "self") ||
                    _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( symbolInfo.FP.Get_typeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                        _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Typ) ).(bool))) ).(bool){
                    return symbolInfo.FP.CanAccess(_env, fromScope, access)
                }
            }
        }
    }
    if self.parent != self{
        return self.parent.FP.GetSymbolTypeInfo(_env, name, fromScope, moduleScope, access)
    }
    return Ast_sym2builtInTypeMap.Get(name)
}

// 2062: decl @lune.@base.@Ast.Scope.filterSymbolTypeInfo
func (self *Ast_Scope) FilterSymbolTypeInfo(_env *LnsEnv, fromScope *Ast_Scope,moduleScope *Ast_Scope,access LnsInt,callback Ast_filterForm) {
    if self.classFlag{
        {
            __exp := self.symbol2SymbolInfoMap.Get("self")
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_SymbolInfo)
                callback(_env, _exp)
            }
        }
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( moduleScope == fromScope) ||
        _env.SetStackVal( moduleScope == self) ||
        _env.SetStackVal( Lns_op_not(self.classFlag)) ).(bool){
        for _, _symbolInfo := range( self.symbol2SymbolInfoMap.Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if Lns_op_not(callback(_env, symbolInfo)){
                return 
            }
        }
    }
    if self.parent != self{
        self.parent.FP.FilterSymbolTypeInfo(_env, fromScope, moduleScope, access, callback)
    }
}

// 2084: decl @lune.@base.@Ast.Scope.add
func (self *Ast_Scope) Add(_env *LnsEnv, processInfo *Ast_ProcessInfo,kind LnsInt,canBeLeft bool,canBeRight bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,mutMode LnsInt,hasValueFlag bool,isLazyLoad bool)(LnsAny, LnsAny) {
    var ownerTypeInfo LnsAny
    ownerTypeInfo = nil
    if _switch6553 := kind; _switch6553 == Ast_SymbolKind__Typ || _switch6553 == Ast_SymbolKind__Fun || _switch6553 == Ast_SymbolKind__Mac {
        var existSymbol LnsAny
        if _switch6499 := typeInfo.FP.Get_kind(_env); _switch6499 == Ast_TypeInfoKind__Enum {
            if _env.NilAccFin(_env.NilAccPush(self.ownerTypeInfo) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) == Ast_TypeInfoKind__Class{
                existSymbol = self.FP.GetSymbolInfoField(_env, name, true, self, Ast_ScopeAccess__Full)
                
            } else { 
                existSymbol = self.FP.GetSymbolInfo(_env, name, self, true, Ast_ScopeAccess__Full)
                
            }
        } else if _switch6499 == Ast_TypeInfoKind__Class || _switch6499 == Ast_TypeInfoKind__Module {
            existSymbol = self.FP.GetSymbolInfoChild(_env, name)
            
        } else {
            existSymbol = self.FP.GetSymbolInfo(_env, name, self, true, Ast_ScopeAccess__Full)
            
        }
        if existSymbol != nil{
            existSymbol_1669 := existSymbol.(*Ast_SymbolInfo)
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( typeInfo.FP.Get_kind(_env) != existSymbol_1669.FP.Get_typeInfo(_env).FP.Get_kind(_env)) ||
                _env.SetStackVal( Lns_op_not(Ast_isBuiltin(_env, existSymbol_1669.FP.Get_typeInfo(_env).FP.Get_typeId(_env).Id))) ).(bool){
                return nil, existSymbol_1669
            }
        }
    } else if _switch6553 == Ast_SymbolKind__Var {
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Module{
            ownerTypeInfo = typeInfo
            
        }
    }
    var symbolInfo *Ast_NormalSymbolInfo
    symbolInfo = NewAst_NormalSymbolInfo(_env, processInfo, kind, canBeLeft, canBeRight, self, accessMode, staticFlag, name, pos, typeInfo, mutMode, hasValueFlag, isLazyLoad)
    symbolInfo.FP.set_namespaceTypeInfo(_env, ownerTypeInfo)
    self.symbol2SymbolInfoMap.Set(name,&symbolInfo.Ast_SymbolInfo)
    return &symbolInfo.Ast_SymbolInfo, nil
}

// 2133: decl @lune.@base.@Ast.Scope.addSymbolInfo
func (self *Ast_Scope) AddSymbolInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,symbolInfo *Ast_SymbolInfo)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, symbolInfo.FP.Get_kind(_env), symbolInfo.FP.Get_canBeLeft(_env), symbolInfo.FP.Get_canBeRight(_env), symbolInfo.FP.Get_name(_env), symbolInfo.FP.Get_pos(_env), symbolInfo.FP.Get_typeInfo(_env), symbolInfo.FP.Get_accessMode(_env), symbolInfo.FP.Get_staticFlag(_env), symbolInfo.FP.Get_mutMode(_env), symbolInfo.FP.Get_hasValueFlag(_env), symbolInfo.FP.Get_isLazyLoad(_env))
}

// 2143: decl @lune.@base.@Ast.Scope.addLocalVar
func (self *Ast_Scope) AddLocalVar(_env *LnsEnv, processInfo *Ast_ProcessInfo,argFlag bool,canBeLeft bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutable LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( argFlag) &&
        _env.SetStackVal( Ast_SymbolKind__Arg) ||
        _env.SetStackVal( Ast_SymbolKind__Var) ).(LnsInt), canBeLeft, name != "_", name, pos, typeInfo, Ast_AccessMode__Local, false, mutable, true, false)
}

// 2156: decl @lune.@base.@Ast.Scope.addUnwrapedVar
func (self *Ast_Scope) AddUnwrapedVar(_env *LnsEnv, processInfo *Ast_ProcessInfo,argFlag bool,canBeLeft bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutable LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( argFlag) &&
        _env.SetStackVal( Ast_SymbolKind__Arg) ||
        _env.SetStackVal( Ast_SymbolKind__Var) ).(LnsInt), canBeLeft, true, name, pos, typeInfo, Ast_AccessMode__Local, false, mutable, true, false)
}

// 2167: decl @lune.@base.@Ast.Scope.addStaticVar
func (self *Ast_Scope) AddStaticVar(_env *LnsEnv, processInfo *Ast_ProcessInfo,argFlag bool,canBeLeft bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutable LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( argFlag) &&
        _env.SetStackVal( Ast_SymbolKind__Arg) ||
        _env.SetStackVal( Ast_SymbolKind__Var) ).(LnsInt), canBeLeft, true, name, pos, typeInfo, Ast_AccessMode__Pub, true, mutable, true, false)
}

// 2177: decl @lune.@base.@Ast.Scope.addVar
func (self *Ast_Scope) AddVar(_env *LnsEnv, processInfo *Ast_ProcessInfo,accessMode LnsInt,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutable LnsInt,hasValueFlag bool)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Var, true, true, name, pos, typeInfo, accessMode, false, mutable, hasValueFlag, false)
}

// 2191: decl @lune.@base.@Ast.Scope.addEnumVal
func (self *Ast_Scope) AddEnumVal(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Mbr, false, true, name, pos, typeInfo, Ast_AccessMode__Pub, true, Ast_MutMode__Mut, true, false)
}

// 2201: decl @lune.@base.@Ast.Scope.addEnum
func (self *Ast_Scope) AddEnum(_env *LnsEnv, processInfo *Ast_ProcessInfo,accessMode LnsInt,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Typ, false, false, name, pos, typeInfo, accessMode, true, Ast_MutMode__Mut, true, false)
}

// 2211: decl @lune.@base.@Ast.Scope.addAlgeVal
func (self *Ast_Scope) AddAlgeVal(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Mbr, false, true, name, pos, typeInfo, Ast_AccessMode__Pub, true, Ast_MutMode__Mut, true, false)
}

// 2221: decl @lune.@base.@Ast.Scope.addAlge
func (self *Ast_Scope) AddAlge(_env *LnsEnv, processInfo *Ast_ProcessInfo,accessMode LnsInt,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Typ, false, false, name, pos, typeInfo, accessMode, true, Ast_MutMode__Mut, true, false)
}

// 2231: decl @lune.@base.@Ast.Scope.addAlternate
func (self *Ast_Scope) AddAlternate(_env *LnsEnv, processInfo *Ast_ProcessInfo,accessMode LnsInt,name string,pos LnsAny,typeInfo *Ast_TypeInfo) {
    self.FP.Add(_env, processInfo, Ast_SymbolKind__Typ, false, false, name, pos, typeInfo, accessMode, true, Ast_MutMode__Mut, true, false)
}

// 2240: decl @lune.@base.@Ast.Scope.addMember
func (self *Ast_Scope) AddMember(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,mutMode LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Mbr, true, true, name, pos, typeInfo, accessMode, staticFlag, mutMode, true, false)
}

// 2250: decl @lune.@base.@Ast.Scope.addMethod
func (self *Ast_Scope) AddMethod(_env *LnsEnv, processInfo *Ast_ProcessInfo,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,mutable bool)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Mtd, true, staticFlag, typeInfo.FP.Get_rawTxt(_env), pos, typeInfo, accessMode, staticFlag, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mutable) &&
        _env.SetStackVal( Ast_MutMode__Mut) ||
        _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, false)
}

// 2260: decl @lune.@base.@Ast.Scope.addFunc
func (self *Ast_Scope) AddFunc(_env *LnsEnv, processInfo *Ast_ProcessInfo,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,mutable bool)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Fun, true, true, typeInfo.FP.Get_rawTxt(_env), pos, typeInfo, accessMode, staticFlag, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mutable) &&
        _env.SetStackVal( Ast_MutMode__Mut) ||
        _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, false)
}

// 2271: decl @lune.@base.@Ast.Scope.addForm
func (self *Ast_Scope) AddForm(_env *LnsEnv, processInfo *Ast_ProcessInfo,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Typ, false, false, typeInfo.FP.Get_rawTxt(_env), pos, typeInfo, accessMode, true, Ast_MutMode__IMut, false, false)
}

// 2281: decl @lune.@base.@Ast.Scope.addMacro
func (self *Ast_Scope) AddMacro(_env *LnsEnv, processInfo *Ast_ProcessInfo,pos LnsAny,typeInfo *Ast_TypeInfo,accessMode LnsInt)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Mac, false, false, typeInfo.FP.Get_rawTxt(_env), pos, typeInfo, accessMode, true, Ast_MutMode__IMut, true, false)
}

// 2292: decl @lune.@base.@Ast.Scope.addClassLazy
func (self *Ast_Scope) AddClassLazy(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo,lazyLoad bool)(LnsAny, LnsAny) {
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Typ, false, false, name, pos, typeInfo, typeInfo.FP.Get_accessMode(_env), true, Ast_MutMode__Mut, true, lazyLoad)
}

// 2301: decl @lune.@base.@Ast.Scope.addClass
func (self *Ast_Scope) AddClass(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    return self.FP.AddClassLazy(_env, processInfo, name, pos, typeInfo, false)
}

// 2308: decl @lune.@base.@Ast.Scope.addExtModule
func (self *Ast_Scope) AddExtModule(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo,lazy bool,lang LnsInt)(LnsAny, LnsAny) {
    if lang != Types_Lang__Same{
        switch _exp7733 := processInfo.FP.CreateLuaval(_env, typeInfo, true).(type) {
        case *Ast_LuavalResult__Err:
        mess := _exp7733.Val1
            Util_err(_env, mess)
        case *Ast_LuavalResult__OK:
        luavalTypeInfo := _exp7733.Val1
            typeInfo = luavalTypeInfo
            
        }
    }
    return self.FP.Add(_env, processInfo, Ast_SymbolKind__Typ, false, false, name, pos, typeInfo, typeInfo.FP.Get_accessMode(_env), false, Ast_MutMode__Mut, true, lazy)
}

// 2362: decl @lune.@base.@Ast.Scope.setClosure
func (self *Ast_Scope) SetClosure(_env *LnsEnv, workSymbol *Ast_SymbolInfo) {
    var symbol *Ast_SymbolInfo
    symbol = workSymbol.FP.GetOrg(_env)
    var targetFuncType *Ast_TypeInfo
    targetFuncType = symbol.FP.Get_namespaceTypeInfo(_env)
    var funcType *Ast_TypeInfo
    funcType = self.FP.GetNamespaceTypeInfo(_env)
    for  {
        var funcScope *Ast_Scope
        funcScope = Lns_unwrap( funcType.FP.Get_scope(_env)).(*Ast_Scope)
        if Lns_op_not(funcScope.closureSymMap.Get(symbol.FP.Get_symbolId(_env))){
            funcScope.closureSymMap.Set(symbol.FP.Get_symbolId(_env),symbol)
            funcScope.closureSym2NumMap.Set(symbol,funcScope.closureSymList.Len())
            funcScope.closureSymList.Insert(Ast_SymbolInfo2Stem(symbol))
            funcType = funcScope.parent.FP.GetNamespaceTypeInfo(_env)
            
        } else { 
            break
        }
        if funcType == targetFuncType{
            break
        }
    }
    if Lns_op_not(symbol.FP.Get_hasAccessFromClosure(_env)){
        symbol.FP.Set_hasAccessFromClosure(_env, true)
    }
}

// 2391: decl @lune.@base.@Ast.Scope.isClosureAccess
func (self *Ast_Scope) IsClosureAccess(_env *LnsEnv, moduleScope *Ast_Scope,symbol *Ast_SymbolInfo) bool {
    var processInfo *Ast_ProcessInfo
    processInfo = moduleScope.FP.GetModule(_env).FP.getProcessInfo(_env)
    if _switch8119 := symbol.FP.Get_kind(_env); _switch8119 == Ast_SymbolKind__Var || _switch8119 == Ast_SymbolKind__Arg || _switch8119 == Ast_SymbolKind__Fun {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( symbol.FP.Get_scope(_env) == moduleScope) ||
            _env.SetStackVal( symbol.FP.Get_scope(_env) == Ast_rootScopeRo) ).(bool){
        } else if symbol.FP.Get_name(_env) == "self"{
            var funcType *Ast_TypeInfo
            funcType = self.FP.GetNamespaceTypeInfo(_env)
            if funcType.FP.Get_parentInfo(_env).FP.IsInheritFrom(_env, processInfo, symbol.FP.Get_namespaceTypeInfo(_env).FP.Get_parentInfo(_env), nil){
            } else { 
                return true
            }
        } else { 
            var funcType *Ast_TypeInfo
            funcType = self.FP.GetNamespaceTypeInfo(_env)
            if funcType != symbol.FP.Get_namespaceTypeInfo(_env){
                return true
            }
        }
    }
    return false
}

// 2431: decl @lune.@base.@Ast.Scope.accessSymbol
func (self *Ast_Scope) AccessSymbol(_env *LnsEnv, moduleScope *Ast_Scope,symbol *Ast_SymbolInfo) {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( symbol.FP.Get_kind(_env) == Ast_SymbolKind__Fun) &&
        _env.SetStackVal( self.FP.GetNamespaceTypeInfo(_env) == symbol.FP.Get_typeInfo(_env)) ).(bool)){
        return 
    }
    if self.FP.IsClosureAccess(_env, moduleScope, symbol){
        self.FP.SetClosure(_env, symbol)
    }
}

// 2558: decl @lune.@base.@Ast.Scope.getClassTypeInfo
func (self *Ast_Scope) GetClassTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    var scope *Ast_Scope
    scope = self
    for  {
        {
            _owner := scope.ownerTypeInfo
            if !Lns_IsNil( _owner ) {
                owner := _owner.(*Ast_TypeInfo)
                if _switch8527 := owner.FP.Get_kind(_env); _switch8527 == Ast_TypeInfoKind__Class || _switch8527 == Ast_TypeInfoKind__IF || _switch8527 == Ast_TypeInfoKind__Module {
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

// 4428: decl @lune.@base.@Ast.Scope.addOverrideImut
func (self *Ast_Scope) AddOverrideImut(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) {
    var typeInfo *Ast_TypeInfo
    if Ast_TypeInfo_isMut(_env, symbolInfo.FP.Get_typeInfo(_env)){
        var processInfo *Ast_ProcessInfo
        processInfo = self.FP.GetProcessInfo(_env)
        typeInfo = processInfo.FP.CreateModifier(_env, symbolInfo.FP.Get_typeInfo(_env), Ast_MutMode__IMut)
        
    } else { 
        typeInfo = symbolInfo.FP.Get_typeInfo(_env)
        
    }
    self.symbol2SymbolInfoMap.Set(symbolInfo.FP.Get_name(_env),&NewAst_AccessSymbolInfo(_env, symbolInfo, &Ast_OverrideMut__IMut{typeInfo}, false).Ast_SymbolInfo)
}

// 4574: decl @lune.@base.@Ast.Scope.addIgnoredVar
func (self *Ast_Scope) AddIgnoredVar(_env *LnsEnv, processInfo *Ast_ProcessInfo) {
    self.FP.AddLocalVar(_env, processInfo, false, true, "_", nil, Ast_builtinTypeEmpty, Ast_MutMode__Mut)
}

// 5031: decl @lune.@base.@Ast.Scope.addAlias
func (self *Ast_Scope) AddAlias(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,pos *Types_Position,externalFlag bool,accessMode LnsInt,parentInfo *Ast_TypeInfo,symbolInfo *Ast_SymbolInfo)(LnsAny, LnsAny) {
    var aliasType *Ast_AliasTypeInfo
    aliasType = self.FP.GetProcessInfo(_env).FP.CreateAlias(_env, processInfo, name, externalFlag, accessMode, parentInfo, symbolInfo.FP.Get_typeInfo(_env).FP.Get_srcTypeInfo(_env))
    return self.FP.Add(_env, processInfo, symbolInfo.FP.Get_kind(_env), false, symbolInfo.FP.Get_canBeRight(_env), name, pos, &aliasType.Ast_TypeInfo, accessMode, true, Ast_MutMode__IMut, true, false)
}

// 5044: decl @lune.@base.@Ast.Scope.addAliasForType
func (self *Ast_Scope) AddAliasForType(_env *LnsEnv, processInfo *Ast_ProcessInfo,name string,pos LnsAny,typeInfo *Ast_TypeInfo)(LnsAny, LnsAny) {
    var skind LnsInt
    skind = Ast_SymbolKind__Typ
    var canBeRight bool
    canBeRight = false
    if _switch19857 := typeInfo.FP.Get_kind(_env); _switch19857 == Ast_TypeInfoKind__Func {
        skind = Ast_SymbolKind__Fun
        
        canBeRight = true
        
    } else if _switch19857 == Ast_TypeInfoKind__Form || _switch19857 == Ast_TypeInfoKind__FormFunc {
        canBeRight = true
        
    } else if _switch19857 == Ast_TypeInfoKind__Macro {
        skind = Ast_SymbolKind__Mac
        
    }
    return self.FP.Add(_env, processInfo, skind, false, canBeRight, name, pos, typeInfo, typeInfo.FP.Get_accessMode(_env), true, Ast_MutMode__IMut, true, false)
}


// declaration Class -- TypeData
type Ast_TypeDataMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    Get_children(_env *LnsEnv) *LnsList
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
func NewAst_TypeData(_env *LnsEnv, arg1 *LnsList) *Ast_TypeData {
    obj := &Ast_TypeData{}
    obj.FP = obj
    obj.InitAst_TypeData(_env, arg1)
    return obj
}
func (self *Ast_TypeData) InitAst_TypeData(_env *LnsEnv, arg1 *LnsList) {
    self.children = arg1
}
func (self *Ast_TypeData) Get_children(_env *LnsEnv) *LnsList{ return self.children }
// 673: decl @lune.@base.@Ast.TypeData.addChildren
func (self *Ast_TypeData) AddChildren(_env *LnsEnv, child *Ast_TypeInfo) {
    self.children.Insert(Ast_TypeInfo2Stem(child))
}


// declaration Class -- SerializeInfo
type Ast_SerializeInfoMtd interface {
    Get_validChildrenMap(_env *LnsEnv) *LnsMap
    IsValidChildren(_env *LnsEnv, arg1 *Ast_IdInfo) bool
    SerializeId(_env *LnsEnv, arg1 *Ast_IdInfo) string
}
type Ast_SerializeInfo struct {
    processInfo2Id *LnsMap
    validChildrenMap *LnsMap
    FP Ast_SerializeInfoMtd
}
func Ast_SerializeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Ast_SerializeInfo).FP
}
type Ast_SerializeInfoDownCast interface {
    ToAst_SerializeInfo() *Ast_SerializeInfo
}
func Ast_SerializeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Ast_SerializeInfoDownCast)
    if ok { return work.ToAst_SerializeInfo() }
    return nil
}
func (obj *Ast_SerializeInfo) ToAst_SerializeInfo() *Ast_SerializeInfo {
    return obj
}
func NewAst_SerializeInfo(_env *LnsEnv, arg1 *LnsMap, arg2 *LnsMap) *Ast_SerializeInfo {
    obj := &Ast_SerializeInfo{}
    obj.FP = obj
    obj.InitAst_SerializeInfo(_env, arg1, arg2)
    return obj
}
func (self *Ast_SerializeInfo) InitAst_SerializeInfo(_env *LnsEnv, arg1 *LnsMap, arg2 *LnsMap) {
    self.processInfo2Id = arg1
    self.validChildrenMap = arg2
}
func (self *Ast_SerializeInfo) Get_validChildrenMap(_env *LnsEnv) *LnsMap{ return self.validChildrenMap }
// 722: decl @lune.@base.@Ast.SerializeInfo.isValidChildren
func (self *Ast_SerializeInfo) IsValidChildren(_env *LnsEnv, idInfo *Ast_IdInfo) bool {
    return Lns_op_not(self.validChildrenMap.Get(idInfo))
}

// 725: decl @lune.@base.@Ast.SerializeInfo.serializeId
func (self *Ast_SerializeInfo) SerializeId(_env *LnsEnv, idInfo *Ast_IdInfo) string {
    var id LnsInt
    id = idInfo.FP.Get_orgId(_env)
    if id >= Ast_extStartId{
        return _env.LuaVM.String_format("{ id = %d, mod = 0 }", []LnsAny{id})
    }
    var processId LnsInt
    processId = Lns_unwrap( self.processInfo2Id.Get(idInfo.ProcessInfo)).(LnsInt)
    return _env.LuaVM.String_format("{ id = %d, mod = %d }", []LnsAny{idInfo.FP.Get_orgId(_env), processId})
}


// declaration Class -- TypeInfo
type Ast_TypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
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
func (self *Ast_TypeInfo) Get_scope(_env *LnsEnv) LnsAny{ return self.scope }
func (self *Ast_TypeInfo) Get_typeData(_env *LnsEnv) *Ast_TypeData{ return self.typeData }
func (self *Ast_TypeInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo{ return self.processInfo }
// 761: decl @lune.@base.@Ast.TypeInfo.get_asyncMode
func (self *Ast_TypeInfo) Get_asyncMode(_env *LnsEnv) LnsInt {
    return Ast_Async__Async
}

// 765: decl @lune.@base.@Ast.TypeInfo.getProcessInfo
func (self *Ast_TypeInfo) getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.processInfo
}

// 769: decl @lune.@base.@Ast.TypeInfo.switchScope
func (self *Ast_TypeInfo) SwitchScope(_env *LnsEnv, scope *Ast_Scope) {
    self.scope = scope
    
    scope.FP.SwitchOwnerTypeInfo(_env, self)
}

// 779: decl @lune.@base.@Ast.TypeInfo.getOverridingType
func (self *Ast_TypeInfo) GetOverridingType(_env *LnsEnv) LnsAny {
    return nil
}

// 787: DeclConstr
func (self *Ast_TypeInfo) InitAst_TypeInfo(_env *LnsEnv, scope LnsAny,processInfo *Ast_ProcessInfo) {
    self.scope = scope
    
    {
        __exp := scope
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_Scope)
            _exp.FP.Set_ownerTypeInfo(_env, self)
        }
    }
    self.typeData = NewAst_TypeData(_env, NewLnsList([]LnsAny{}))
    
    self.processInfo = processInfo
    
}

// 798: decl @lune.@base.@Ast.TypeInfo.get_aliasSrc
func (self *Ast_TypeInfo) Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo {
    return self
}

// 802: decl @lune.@base.@Ast.TypeInfo.get_extedType
func (self *Ast_TypeInfo) Get_extedType(_env *LnsEnv) *Ast_TypeInfo {
    return self
}

// 806: decl @lune.@base.@Ast.TypeInfo.getModulePath
func Ast_TypeInfo_getModulePath(_env *LnsEnv, fullname string) string {
    return Lns_car(_env.LuaVM.String_gsub(fullname,"@", "")).(string)
}

// 811: decl @lune.@base.@Ast.TypeInfo.isModule
func (self *Ast_TypeInfo) IsModule(_env *LnsEnv) bool {
    return true
}

// 816: decl @lune.@base.@Ast.TypeInfo.getParentId
func (self *Ast_TypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return Ast_rootTypeIdInfo
}

// 821: decl @lune.@base.@Ast.TypeInfo.get_baseId
func (self *Ast_TypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return Ast_rootTypeIdInfo
}

// 826: decl @lune.@base.@Ast.TypeInfo.isInheritFrom
func (self *Ast_TypeInfo) IsInheritFrom(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    return false
}

// 833: decl @lune.@base.@Ast.TypeInfo.get_rawTxt
func (self *Ast_TypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return ""
}

// 843: decl @lune.@base.@Ast.TypeInfo.getTxtWithRaw
func (self *Ast_TypeInfo) GetTxtWithRaw(_env *LnsEnv, raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return ""
}

// 856: decl @lune.@base.@Ast.TypeInfo.getTxt
func (self *Ast_TypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}


// 868: decl @lune.@base.@Ast.TypeInfo.canEvalWith
func (self *Ast_TypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return false, nil
}

// 876: decl @lune.@base.@Ast.TypeInfo.get_abstractFlag
func (self *Ast_TypeInfo) Get_abstractFlag(_env *LnsEnv) bool {
    return false
}

// 881: decl @lune.@base.@Ast.TypeInfo.serialize
func (self *Ast_TypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    return 
}

// 887: decl @lune.@base.@Ast.TypeInfo.get_display_stirng_with
func (self *Ast_TypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    return ""
}

// 892: decl @lune.@base.@Ast.TypeInfo.get_display_stirng
func (self *Ast_TypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, "", nil)
}

// 898: decl @lune.@base.@Ast.TypeInfo.get_srcTypeInfo
func (self *Ast_TypeInfo) Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self
}

// 903: decl @lune.@base.@Ast.TypeInfo.equals
func (self *Ast_TypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if Lns_isCondTrue( checkModifer){
        return self == typeInfo
    }
    return self == typeInfo.FP.Get_srcTypeInfo(_env)
}

// 914: decl @lune.@base.@Ast.TypeInfo.get_externalFlag
func (self *Ast_TypeInfo) Get_externalFlag(_env *LnsEnv) bool {
    return false
}

// 919: decl @lune.@base.@Ast.TypeInfo.get_interfaceList
func (self *Ast_TypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList {
    return Ast_dummyList
}

// 923: decl @lune.@base.@Ast.TypeInfo.get_itemTypeInfoList
func (self *Ast_TypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList {
    return Ast_dummyList
}

// 927: decl @lune.@base.@Ast.TypeInfo.get_argTypeInfoList
func (self *Ast_TypeInfo) Get_argTypeInfoList(_env *LnsEnv) *LnsList {
    return Ast_dummyList
}

// 931: decl @lune.@base.@Ast.TypeInfo.get_retTypeInfoList
func (self *Ast_TypeInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList {
    return Ast_dummyList
}


// 941: decl @lune.@base.@Ast.TypeInfo.hasRouteNamespaceFrom
func (self *Ast_TypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, other *Ast_TypeInfo) bool {
    for  {
        if other == self{
            return true
        }
        if other.FP.Get_parentInfo(_env) == other{
            break
        }
        other = other.FP.Get_parentInfo(_env)
        
    }
    return false
}

// 954: decl @lune.@base.@Ast.TypeInfo.getModule
func (self *Ast_TypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    if self.FP.IsModule(_env){
        return self
    }
    return self.FP.Get_parentInfo(_env).FP.GetModule(_env)
}

// 961: decl @lune.@base.@Ast.TypeInfo.get_typeId
func (self *Ast_TypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo {
    return Ast_rootTypeIdInfo
}

// 965: decl @lune.@base.@Ast.TypeInfo.get_kind
func (self *Ast_TypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Root
}

// 973: decl @lune.@base.@Ast.TypeInfo.get_staticFlag
func (self *Ast_TypeInfo) Get_staticFlag(_env *LnsEnv) bool {
    return false
}

// 977: decl @lune.@base.@Ast.TypeInfo.get_accessMode
func (self *Ast_TypeInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return Ast_AccessMode__Pri
}

// 981: decl @lune.@base.@Ast.TypeInfo.get_autoFlag
func (self *Ast_TypeInfo) Get_autoFlag(_env *LnsEnv) bool {
    return false
}

// 985: decl @lune.@base.@Ast.TypeInfo.get_nonnilableType
func (self *Ast_TypeInfo) Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo {
    return self
}


// 992: decl @lune.@base.@Ast.TypeInfo.get_nilable
func (self *Ast_TypeInfo) Get_nilable(_env *LnsEnv) bool {
    return false
}

// 996: decl @lune.@base.@Ast.TypeInfo.get_nilableTypeInfo
func (self *Ast_TypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self
}

// 1004: decl @lune.@base.@Ast.TypeInfo.get_children
func (self *Ast_TypeInfo) Get_children(_env *LnsEnv) *LnsList {
    return self.typeData.FP.Get_children(_env)
}

// 1008: decl @lune.@base.@Ast.TypeInfo.addChildren
func (self *Ast_TypeInfo) AddChildren(_env *LnsEnv, child *Ast_TypeInfo) {
    self.typeData.FP.AddChildren(_env, child)
}

// 1013: decl @lune.@base.@Ast.TypeInfo.get_mutMode
func (self *Ast_TypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return Ast_MutMode__Mut
}

// 1017: decl @lune.@base.@Ast.TypeInfo.isMut
func Ast_TypeInfo_isMut(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    return Ast_isMutable(_env, typeInfo.FP.Get_mutMode(_env))
}

// 1022: decl @lune.@base.@Ast.TypeInfo.getParentFullName
func (self *Ast_TypeInfo) GetParentFullName(_env *LnsEnv, typeNameCtrl *Ast_TypeNameCtrl,importInfo LnsAny,localFlag LnsAny) string {
    return typeNameCtrl.FP.GetParentFullName(_env, self, importInfo, localFlag)
}

// 1030: decl @lune.@base.@Ast.TypeInfo.applyGeneric
func (self *Ast_TypeInfo) ApplyGeneric(_env *LnsEnv, alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    return self
}

// 1036: decl @lune.@base.@Ast.TypeInfo.get_genSrcTypeInfo
func (self *Ast_TypeInfo) Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self
}

// 1040: decl @lune.@base.@Ast.TypeInfo.serializeTypeInfoList
func (self *Ast_TypeInfo) SerializeTypeInfoList(_env *LnsEnv, serializeInfo *Ast_SerializeInfo,name string,list *LnsList,onlyPub LnsAny) string {
    var work string
    work = name
    for _, _typeInfo := range( list.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(onlyPub)) ||
            _env.SetStackVal( typeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) ).(bool){
            if len(work) != len(name){
                work = work + ", "
                
            }
            work = _env.LuaVM.String_format("%s%s", []LnsAny{work, serializeInfo.FP.SerializeId(_env, typeInfo.FP.Get_typeId(_env))})
            
        }
    }
    return work + "}, "
}

// 1056: decl @lune.@base.@Ast.TypeInfo.createScope
func Ast_TypeInfo_createScope(_env *LnsEnv, processInfo *Ast_ProcessInfo,parent LnsAny,classFlag bool,baseInfo LnsAny,interfaceList LnsAny) *Ast_Scope {
    var inheritScope LnsAny
    inheritScope = nil
    if baseInfo != nil{
        baseInfo_830 := baseInfo.(*Ast_TypeInfo)
        inheritScope = Lns_unwrap( baseInfo_830.FP.Get_scope(_env)).(*Ast_Scope)
        
    }
    var ifScopeList *LnsList
    ifScopeList = NewLnsList([]LnsAny{})
    if interfaceList != nil{
        interfaceList_833 := interfaceList.(*LnsList)
        for _, _ifType := range( interfaceList_833.Items ) {
            ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            ifScopeList.Insert(Ast_Scope2Stem(Lns_unwrap( ifType.FP.Get_scope(_env)).(*Ast_Scope)))
        }
    }
    return NewAst_Scope(_env, processInfo, parent, classFlag, inheritScope, ifScopeList)
}



// 1258: decl @lune.@base.@Ast.TypeInfo.hasBase
func (self *Ast_TypeInfo) HasBase(_env *LnsEnv) bool {
    return self.FP.Get_baseTypeInfo(_env) != Ast_headTypeInfo
}

// 1378: decl @lune.@base.@Ast.TypeInfo.isInherit
func Ast_TypeInfo_isInherit(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    var baseTypeInfo *Ast_TypeInfo
    baseTypeInfo = typeInfo.FP.Get_baseTypeInfo(_env)
    var otherTypeId *Ast_IdInfo
    otherTypeId = other.FP.Get_typeId(_env)
    if typeInfo.FP.Get_typeId(_env).FP.Equals(_env, otherTypeId){
        return true
    }
    if baseTypeInfo != Ast_headTypeInfo{
        if baseTypeInfo.FP.IsInheritFrom(_env, processInfo, other, alt2type){
            return true
        }
    }
    for _, _ifType := range( typeInfo.FP.Get_interfaceList(_env).Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if ifType.FP.IsInheritFrom(_env, processInfo, other, alt2type){
            return true
        }
    }
    return false
}

// 2451: decl @lune.@base.@Ast.TypeInfo.createAlt2typeMap
func (self *Ast_TypeInfo) CreateAlt2typeMap(_env *LnsEnv, detectFlag bool) *LnsMap {
    return Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, detectFlag)
}

// 5384: decl @lune.@base.@Ast.TypeInfo.getCommonTypeCombo
func Ast_TypeInfo_getCommonTypeCombo(_env *LnsEnv, processInfo *Ast_ProcessInfo,commonType LnsAny,otherType LnsAny,alt2type *LnsMap) LnsAny {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_builtinTypeNone
    switch _exp21268 := commonType.(type) {
    case *Ast_CommonType__Combine:
    comb := _exp21268.Val1
        return comb.FP.AndType(_env, processInfo, otherType, alt2type)
    case *Ast_CommonType__Normal:
    workTypeInfo := _exp21268.Val1
        typeInfo = workTypeInfo
        
    }
    var other *Ast_TypeInfo
    other = Ast_builtinTypeNone
    switch _exp21301 := otherType.(type) {
    case *Ast_CommonType__Combine:
    comb := _exp21301.Val1
        return comb.FP.AndType(_env, processInfo, commonType, alt2type)
    case *Ast_CommonType__Normal:
    workTypeInfo := _exp21301.Val1
        other = workTypeInfo
        
    }
    var getType func(_env *LnsEnv, workType *Ast_TypeInfo) LnsAny
    getType = func(_env *LnsEnv, workType *Ast_TypeInfo) LnsAny {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( typeInfo.FP.Get_nilable(_env)) ||
            _env.SetStackVal( other.FP.Get_nilable(_env)) ).(bool){
            workType = workType.FP.Get_nilableTypeInfo(_env)
            
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(_env, typeInfo))) ||
            _env.SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(_env, other))) ).(bool){
            workType = processInfo.FP.CreateModifier(_env, workType, Ast_MutMode__IMut)
            
        }
        return &Ast_CommonType__Normal{workType}
    }
    var type1 *Ast_TypeInfo
    type1 = typeInfo.FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env)
    var type2 *Ast_TypeInfo
    type2 = other.FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env)
    if type1 == Ast_builtinTypeNone{
        return otherType
    }
    if type2 == Ast_builtinTypeNone{
        return commonType
    }
    if type1 == Ast_builtinTypeNil{
        return &Ast_CommonType__Normal{other.FP.Get_nilableTypeInfo(_env)}
    }
    if type2 == Ast_builtinTypeNil{
        return &Ast_CommonType__Normal{typeInfo.FP.Get_nilableTypeInfo(_env)}
    }
    if Lns_car(type1.FP.CanEvalWith(_env, processInfo, type2, Ast_CanEvalType__SetOp, alt2type)).(bool){
        return getType(_env, type1)
    }
    if Lns_car(type2.FP.CanEvalWith(_env, processInfo, type1, Ast_CanEvalType__SetOp, alt2type)).(bool){
        return getType(_env, type2)
    }
    var mutMode LnsInt
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Ast_TypeInfo_isMut(_env, typeInfo)) &&
        _env.SetStackVal( Ast_TypeInfo_isMut(_env, other)) ).(bool)){
        mutMode = Ast_MutMode__Mut
        
    } else { 
        mutMode = Ast_MutMode__IMut
        
    }
    if type1.FP.Get_kind(_env) == type2.FP.Get_kind(_env){
        var getCommon func(_env *LnsEnv, workTypeInfo *Ast_TypeInfo,workOther *Ast_TypeInfo,workAlt2type *LnsMap) *Ast_TypeInfo
        getCommon = func(_env *LnsEnv, workTypeInfo *Ast_TypeInfo,workOther *Ast_TypeInfo,workAlt2type *LnsMap) *Ast_TypeInfo {
            switch _exp21582 := Ast_TypeInfo_getCommonTypeCombo(_env, processInfo, &Ast_CommonType__Normal{workTypeInfo}, &Ast_CommonType__Normal{workOther}, workAlt2type).(type) {
            case *Ast_CommonType__Normal:
            info := _exp21582.Val1
                return info
            case *Ast_CommonType__Combine:
            combine := _exp21582.Val1
                return combine.FP.Get_typeInfo(_env, processInfo)
            }
        // insert a dummy
            return nil
        }
        if _switch21772 := type1.FP.Get_kind(_env); _switch21772 == Ast_TypeInfoKind__List {
            return getType(_env, processInfo.FP.CreateList(_env, Ast_AccessMode__Local, Ast_headTypeInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(getCommon(_env, type1.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type))}), mutMode))
        } else if _switch21772 == Ast_TypeInfoKind__Array {
            return getType(_env, processInfo.FP.CreateArray(_env, Ast_AccessMode__Local, Ast_headTypeInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(getCommon(_env, type1.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type))}), mutMode))
        } else if _switch21772 == Ast_TypeInfoKind__Set {
            return getType(_env, processInfo.FP.CreateSet(_env, Ast_AccessMode__Local, Ast_headTypeInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(getCommon(_env, type1.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type))}), mutMode))
        } else if _switch21772 == Ast_TypeInfoKind__Map {
            return getType(_env, processInfo.FP.CreateMap(_env, Ast_AccessMode__Local, Ast_headTypeInfo, getCommon(_env, type1.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type), getCommon(_env, type1.FP.Get_itemTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), type2.FP.Get_itemTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type), mutMode))
        }
    }
    var work *Ast_TypeInfo
    work = type1.FP.Get_baseTypeInfo(_env)
    for work != Ast_headTypeInfo {
        if Lns_car(work.FP.CanEvalWith(_env, processInfo, type2, Ast_CanEvalType__SetOp, alt2type)).(bool){
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( typeInfo.FP.Get_nilable(_env)) ||
                _env.SetStackVal( other.FP.Get_nilable(_env)) ).(bool){
                work = work.FP.Get_nilableTypeInfo(_env)
                
            }
            if Lns_op_not(Ast_isMutable(_env, mutMode)){
                work = processInfo.FP.CreateModifier(_env, work, mutMode)
                
            }
            return &Ast_CommonType__Normal{work}
        }
        work = work.FP.Get_baseTypeInfo(_env)
        
    }
    var combine *Ast_CombineType
    combine = NewAst_CombineType(_env, typeInfo)
    return combine.FP.AndType(_env, processInfo, &Ast_CommonType__Normal{other}, alt2type)
}

// 5529: decl @lune.@base.@Ast.TypeInfo.getCommonType
func Ast_TypeInfo_getCommonType(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,other *Ast_TypeInfo,alt2type *LnsMap) *Ast_TypeInfo {
    switch _exp21946 := Ast_TypeInfo_getCommonTypeCombo(_env, processInfo, &Ast_CommonType__Normal{typeInfo}, &Ast_CommonType__Normal{other}, alt2type).(type) {
    case *Ast_CommonType__Normal:
    info := _exp21946.Val1
        return info
    case *Ast_CommonType__Combine:
    combine := _exp21946.Val1
        return combine.FP.Get_typeInfo(_env, processInfo)
    }
// insert a dummy
    return nil
}

// 6297: decl @lune.@base.@Ast.TypeInfo.checkMatchType
func Ast_TypeInfo_checkMatchType(_env *LnsEnv, processInfo *Ast_ProcessInfo,dstTypeList *LnsList,expTypeList *LnsList,allowDstShort bool,warnForFollowSrcIndex LnsAny,alt2type *LnsMap)(LnsInt, string) {
    var warnMess LnsAny
    warnMess = nil
    var checkDstTypeFrom func(_env *LnsEnv, index LnsInt,srcType *Ast_TypeInfo,srcType2nd *Ast_TypeInfo)(LnsInt, string)
    checkDstTypeFrom = func(_env *LnsEnv, index LnsInt,srcType *Ast_TypeInfo,srcType2nd *Ast_TypeInfo)(LnsInt, string) {
        var workExpType *Ast_TypeInfo
        workExpType = srcType
        {
            var _from25453 LnsInt = index
            var _to25453 LnsInt = dstTypeList.Len()
            for _work25453 := _from25453; _work25453 <= _to25453; _work25453++ {
                dstIndex := _work25453
                var workDstType *Ast_TypeInfo
                workDstType = dstTypeList.GetAt(dstIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var canEval bool
                var evalMess LnsAny
                canEval,evalMess = workDstType.FP.CanEvalWith(_env, processInfo, workExpType, Ast_CanEvalType__SetOp, alt2type)
                if Lns_op_not(canEval){
                    var message string
                    message = _env.LuaVM.String_format("exp(%d) type mismatch %s <- %s: dst %d%s", []LnsAny{dstIndex, workDstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), workExpType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), dstIndex, (_env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( evalMess) &&
                        _env.SetStackVal( _env.LuaVM.String_format(" -- %s", []LnsAny{evalMess})) ||
                        _env.SetStackVal( "") ).(string))})
                    return Ast_MatchType__Error, message
                } else if workExpType == Ast_builtinTypeAbbrNone{
                    return Ast_MatchType__Warn, Code_format(_env, Code_ID__nothing_define_abbr, _env.LuaVM.String_format("use '##', instate of '%s'.", []LnsAny{workDstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil)}))
                }
                workExpType = srcType2nd
                
            }
        }
        return Ast_MatchType__Match, ""
    }
    var checkSrcTypeFrom func(_env *LnsEnv, index LnsInt,dstType *Ast_TypeInfo)(LnsInt, string)
    checkSrcTypeFrom = func(_env *LnsEnv, index LnsInt,dstType *Ast_TypeInfo)(LnsInt, string) {
        {
            var _from25739 LnsInt = index
            var _to25739 LnsInt = expTypeList.Len()
            for _work25739 := _from25739; _work25739 <= _to25739; _work25739++ {
                srcIndex := _work25739
                var expType *Ast_TypeInfo
                expType = expTypeList.GetAt(srcIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var checkType *Ast_TypeInfo
                checkType = expType
                if expType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    if expType.FP.Get_itemTypeInfoList(_env).Len() > 0{
                        checkType = expType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        
                    } else { 
                        checkType = Ast_builtinTypeStem_
                        
                    }
                } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( srcIndex > index) &&
                    _env.SetStackVal( expType.FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr) ).(bool)){
                    return Ast_MatchType__Error, "must not use '##'"
                }
                var canEval bool
                var evalMess LnsAny
                canEval,evalMess = dstType.FP.CanEvalWith(_env, processInfo, checkType, Ast_CanEvalType__SetOp, alt2type)
                if Lns_op_not(canEval){
                    return Ast_MatchType__Error, _env.LuaVM.String_format("exp(%d) type mismatch %s <- %s: src: %d%s", []LnsAny{srcIndex, dstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), srcIndex, (_env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( evalMess) &&
                        _env.SetStackVal( _env.LuaVM.String_format(" -- %s", []LnsAny{evalMess})) ||
                        _env.SetStackVal( "") ).(string))})
                }
                if warnForFollowSrcIndex != nil{
                    warnForFollowSrcIndex_4900 := warnForFollowSrcIndex.(LnsInt)
                    if warnForFollowSrcIndex_4900 <= srcIndex{
                        var workMess string
                        workMess = _env.LuaVM.String_format("use '**' at arg(%d). %s <- %s", []LnsAny{srcIndex, dstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil)})
                        return Ast_MatchType__Warn, workMess
                    }
                }
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    _env.SetStackVal( Lns_op_not(Ast_isExtType(_env, dstType.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)))) &&
                    _env.SetStackVal( Ast_isExtType(_env, expType.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env))) ).(bool)){
                    warnMess = _env.LuaVM.String_format("exp(%d) luaval mismatch %s <- %s", []LnsAny{index, dstType.FP.GetTxt(_env, nil, nil, nil), expType.FP.GetTxt(_env, nil, nil, nil)})
                    
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
                return Ast_MatchType__Error, _env.LuaVM.String_format("over exp. expect:0, actual:%d", []LnsAny{expTypeList.Len()})
            }
            var dstType *Ast_TypeInfo
            dstType = dstTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if dstTypeList.Len() == index{
                if dstType.FP.Get_srcTypeInfo(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__DDD{
                    var isMatch bool
                    var msg LnsAny
                    isMatch,msg = dstType.FP.CanEvalWith(_env, processInfo, expType, Ast_CanEvalType__SetOp, alt2type)
                    if Lns_op_not(isMatch){
                        var workDstType *Ast_TypeInfo
                        workDstType = Ast_applyGenericDefault(_env, dstType, alt2type, dstType.FP.GetModule(_env))
                        return Ast_MatchType__Error, _env.LuaVM.String_format("exp(%d) type mismatch %s(%d) <- %s(%d): index %d%s", []LnsAny{index, workDstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), workDstType.FP.Get_typeId(_env).Id, expType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), expType.FP.Get_typeId(_env).Id, index, _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( msg) &&
                            _env.SetStackVal( _env.LuaVM.String_format(" -- %s", []LnsAny{msg})) ||
                            _env.SetStackVal( _env.LuaVM.String_format("(%s)", []LnsAny{Ast_TypeInfoKind_getTxt( dstType.FP.Get_kind(_env))})) ).(string)})
                    }
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(allowDstShort)) &&
                        _env.SetStackVal( dstTypeList.Len() < expTypeList.Len()) ).(bool)){
                        return Ast_MatchType__Error, _env.LuaVM.String_format("over exp. expect: %d: actual: %d", []LnsAny{dstTypeList.Len(), expTypeList.Len()})
                    }
                } else { 
                    var dddItemType *Ast_TypeInfo
                    dddItemType = Ast_builtinTypeStem_
                    if dstType.FP.Get_itemTypeInfoList(_env).Len() > 0{
                        dddItemType = dstType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        
                    }
                    var result LnsInt
                    var mess string
                    result,mess = checkSrcTypeFrom(_env, index, dddItemType)
                    if result != Ast_MatchType__Match{
                        return result, mess
                    }
                }
                if warnForFollowSrcIndex != nil{
                    warnForFollowSrcIndex_4924 := warnForFollowSrcIndex.(LnsInt)
                    if warnForFollowSrcIndex_4924 <= index{
                        var workMess string
                        workMess = _env.LuaVM.String_format("use '**' at arg(%d). %s <- %s", []LnsAny{index, dstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil)})
                        return Ast_MatchType__Warn, workMess
                    }
                }
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    _env.SetStackVal( Lns_op_not(Ast_isExtType(_env, dstType.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)))) &&
                    _env.SetStackVal( Ast_isExtType(_env, expType.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env))) ).(bool)){
                    warnMess = _env.LuaVM.String_format("exp(%d) luaval mismatch %s <- %s", []LnsAny{index, dstType.FP.GetTxt(_env, nil, nil, nil), expType.FP.GetTxt(_env, nil, nil, nil)})
                    
                }
                
                break
            } else if expTypeList.Len() == index{
                var srcType *Ast_TypeInfo
                srcType = expType
                var srcType2nd *Ast_TypeInfo
                srcType2nd = Ast_builtinTypeAbbrNone
                if expType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    srcType = Ast_builtinTypeStem_
                    
                    srcType2nd = Ast_builtinTypeStem_
                    
                    if expType.FP.Get_itemTypeInfoList(_env).Len() > 0{
                        srcType = expType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo(_env)
                        
                        srcType2nd = srcType
                        
                    }
                } else if expType == Ast_builtinTypeAbbr{
                    srcType2nd = Ast_builtinTypeAbbr
                    
                }
                var result LnsInt
                var mess string
                result,mess = checkDstTypeFrom(_env, index, srcType, srcType2nd)
                if result != Ast_MatchType__Match{
                    return result, mess
                }
                if warnForFollowSrcIndex != nil{
                    warnForFollowSrcIndex_4938 := warnForFollowSrcIndex.(LnsInt)
                    if warnForFollowSrcIndex_4938 <= index{
                        var workMess string
                        workMess = _env.LuaVM.String_format("use '**' at arg(%d). %s <- %s", []LnsAny{index, dstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil)})
                        return Ast_MatchType__Warn, workMess
                    }
                }
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    _env.SetStackVal( Lns_op_not(Ast_isExtType(_env, dstType.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)))) &&
                    _env.SetStackVal( Ast_isExtType(_env, expType.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env))) ).(bool)){
                    warnMess = _env.LuaVM.String_format("exp(%d) luaval mismatch %s <- %s", []LnsAny{index, dstType.FP.GetTxt(_env, nil, nil, nil), expType.FP.GetTxt(_env, nil, nil, nil)})
                    
                }
                
                break
            } else { 
                var canEval bool
                var evalMess LnsAny
                canEval,evalMess = dstType.FP.CanEvalWith(_env, processInfo, expType, Ast_CanEvalType__SetOp, alt2type)
                if Lns_op_not(canEval){
                    return Ast_MatchType__Error, _env.LuaVM.String_format("exp(%d) type mismatch %s(%s:%d) <- %s(%s:%d)%s", []LnsAny{index, dstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), Ast_TypeInfoKind_getTxt( dstType.FP.Get_kind(_env)), dstType.FP.Get_typeId(_env).Id, expType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), Ast_TypeInfoKind_getTxt( expType.FP.Get_kind(_env)), expType.FP.Get_typeId(_env).Id, (_env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( evalMess) &&
                        _env.SetStackVal( _env.LuaVM.String_format(" -- %s", []LnsAny{evalMess})) ||
                        _env.SetStackVal( "") ).(string))})
                }
                if warnForFollowSrcIndex != nil{
                    warnForFollowSrcIndex_4947 := warnForFollowSrcIndex.(LnsInt)
                    if warnForFollowSrcIndex_4947 <= index{
                        var workMess string
                        workMess = _env.LuaVM.String_format("use '**' at arg(%d). %s <- %s", []LnsAny{index, dstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil), expType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil)})
                        return Ast_MatchType__Warn, workMess
                    }
                }
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    _env.SetStackVal( Lns_op_not(Ast_isExtType(_env, dstType.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)))) &&
                    _env.SetStackVal( Ast_isExtType(_env, expType.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env))) ).(bool)){
                    warnMess = _env.LuaVM.String_format("exp(%d) luaval mismatch %s <- %s", []LnsAny{index, dstType.FP.GetTxt(_env, nil, nil, nil), expType.FP.GetTxt(_env, nil, nil, nil)})
                    
                }
                
            }
        }
    } else if Lns_op_not(allowDstShort){
        for _index, _dstType := range( dstTypeList.Items ) {
            index := _index + 1
            dstType := _dstType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(Lns_car(dstType.FP.CanEvalWith(_env, processInfo, Ast_builtinTypeNil, Ast_CanEvalType__SetOp, alt2type)).(bool)){
                return Ast_MatchType__Error, _env.LuaVM.String_format("exp(%d) type mismatch %s <- nil: short", []LnsAny{index, dstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil)})
            }
            return Ast_MatchType__Warn, Code_format(_env, Code_ID__nothing_define_abbr, _env.LuaVM.String_format("use '##', instate of '%s'.", []LnsAny{dstType.FP.GetTxt(_env, Ast_defaultTypeNameCtrl, nil, nil)}))
        }
    }
    if warnMess != nil{
        warnMess_4957 := warnMess.(string)
        return Ast_MatchType__Warn, warnMess_4957
    }
    return Ast_MatchType__Match, ""
}

// 6525: decl @lune.@base.@Ast.TypeInfo.canEvalWithBase
func Ast_TypeInfo_canEvalWithBase(_env *LnsEnv, processInfo *Ast_ProcessInfo,dest *Ast_TypeInfo,destMut bool,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if dest != dest.FP.Get_aliasSrc(_env){
        return dest.FP.Get_aliasSrc(_env).FP.CanEvalWith(_env, processInfo, other, canEvalType, alt2type)
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( dest == Ast_builtinTypeExp) ||
        _env.SetStackVal( dest == Ast_builtinTypeMultiExp) ).(bool){
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( other == Ast_builtinTypeMultiExp) &&
            _env.SetStackVal( dest != Ast_builtinTypeMultiExp) ).(bool)){
            return false, "can't eval from '__exp' to '__exps'."
        }
        if other.FP.Equals(_env, processInfo, Ast_builtinTypeAbbr, nil, nil){
            return false, ""
        }
        return true, nil
    }
    var otherMut bool
    otherMut = Ast_TypeInfo_isMut(_env, other)
    var otherSrc *Ast_TypeInfo
    otherSrc = other.FP.Get_srcTypeInfo(_env).FP.Get_aliasSrc(_env)
    if otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        if otherSrc.FP.Get_itemTypeInfoList(_env).Len() > 0{
            otherSrc = otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo(_env)
            
        } else { 
            otherSrc = Ast_builtinTypeStem_
            
        }
    }
    if otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
        if dest.FP.Get_kind(_env) != Ast_TypeInfoKind__Alternate{
            otherSrc = Ast_AlternateTypeInfo_getAssign(_env, otherSrc, alt2type).FP.Get_srcTypeInfo(_env)
            
        }
    }
    if _switch27188 := canEvalType; _switch27188 == Ast_CanEvalType__SetEq || _switch27188 == Ast_CanEvalType__SetOp || _switch27188 == Ast_CanEvalType__SetOpIMut {
        if dest == Ast_builtinTypeEmpty{
            if _switch26970 := otherSrc; _switch26970 == Ast_builtinTypeAbbr || _switch26970 == Ast_builtinTypeAbbrNone {
                return false, ""
            }
            if otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Func{
                return false, ""
            }
            return true, nil
        } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(otherMut)) &&
            _env.SetStackVal( destMut) ).(bool)){
            if processInfo.FP.Get_validCheckingMutable(_env){
                var nonNilOtherType *Ast_TypeInfo
                nonNilOtherType = otherSrc.FP.Get_nonnilableType(_env)
                if _switch27074 := nonNilOtherType.FP.Get_kind(_env); _switch27074 == Ast_TypeInfoKind__Set || _switch27074 == Ast_TypeInfoKind__Map || _switch27074 == Ast_TypeInfoKind__List || _switch27074 == Ast_TypeInfoKind__IF || _switch27074 == Ast_TypeInfoKind__Alternate {
                    return false, ""
                } else if _switch27074 == Ast_TypeInfoKind__Class {
                    if Ast_builtinTypeString != nonNilOtherType{
                        return false, ""
                    }
                } else if _switch27074 == Ast_TypeInfoKind__Prim {
                    if Ast_builtinTypeStem == nonNilOtherType{
                        return false, ""
                    }
                }
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( otherSrc != Ast_builtinTypeNil) &&
            _env.SetStackVal( otherSrc != Ast_builtinTypeString) &&
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__Prim) &&
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__Func) &&
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__Method) &&
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__Form) &&
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__FormFunc) &&
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__Enum) &&
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__Abbr) &&
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__Alternate) &&
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__Box) &&
            _env.SetStackVal( Lns_op_not(Ast_isGenericType(_env, otherSrc))) &&
            _env.SetStackVal( destMut) &&
            _env.SetStackVal( Lns_op_not(otherMut)) ).(bool)){
            if processInfo.FP.Get_validCheckingMutable(_env){
                return false, ""
            }
        }
    }
    if dest == Ast_builtinTypeStem_{
        return true, nil
    }
    if dest.FP.Get_srcTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        if dest.FP.Get_itemTypeInfoList(_env).Len() > 0{
            return dest.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.CanEvalWith(_env, processInfo, other, canEvalType, alt2type)
        }
        return true, nil
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(dest.FP.Get_nilable(_env))) &&
        _env.SetStackVal( otherSrc.FP.Get_nilable(_env)) ).(bool)){
        if canEvalType != Ast_CanEvalType__Equal{
            if dest.FP.Get_kind(_env) == Ast_TypeInfoKind__Box{
                return dest.FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
            }
            return false, ""
        } else { 
            otherSrc = otherSrc.FP.Get_nonnilableType(_env)
            
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( dest == Ast_builtinTypeStem) &&
        _env.SetStackVal( Lns_op_not(otherSrc.FP.Get_nilable(_env))) ).(bool)){
        return true, nil
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( dest == Ast_builtinTypeForm) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Func) ||
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Form) ||
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__FormFunc) ).(bool))) ).(bool)){
        if Lns_car(Ast_isSettableToForm_3956_(_env, processInfo, otherSrc)).(bool){
            return true, nil
        }
        return false, ""
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( otherSrc == Ast_builtinTypeNil) ||
        _env.SetStackVal( otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr) ).(bool){
        if dest.FP.Get_kind(_env) != Ast_TypeInfoKind__Nilable{
            return false, ""
        }
        return true, nil
    }
    if dest.FP.Get_typeId(_env).FP.Equals(_env, otherSrc.FP.Get_typeId(_env)){
        return true, nil
    }
    if dest.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        return dest.FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
    }
    {
        _extTypeInfo := Ast_ExtTypeInfoDownCastF(otherSrc.FP)
        if !Lns_IsNil( _extTypeInfo ) {
            extTypeInfo := _extTypeInfo.(*Ast_ExtTypeInfo)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( canEvalType != Ast_CanEvalType__SetEq) &&
                _env.SetStackVal( Lns_op_not(Lns_car(Ast_failCreateLuavalWith_3312_(_env, extTypeInfo.FP.Get_extedType(_env), Ast_LuavalConvKind__ToLua, false)))) ).(bool)){
                otherSrc = extTypeInfo.FP.Get_extedType(_env)
                
            }
        }
    }
    if otherSrc.FP.Get_typeId(_env).FP.Equals(_env, dest.FP.Get_typeId(_env)){
        return true, nil
    }
    if dest.FP.Get_kind(_env) != otherSrc.FP.Get_kind(_env){
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
            _env.SetStackVal( otherSrc.FP.HasBase(_env)) ).(bool)){
            return Ast_TypeInfo_canEvalWithBase(_env, processInfo, dest, destMut, otherSrc.FP.Get_baseTypeInfo(_env), canEvalType, alt2type)
        }
        if dest.FP.Get_kind(_env) == Ast_TypeInfoKind__Nilable{
            var dstNonNil *Ast_TypeInfo
            if destMut{
                dstNonNil = dest.FP.Get_nonnilableType(_env)
                
            } else { 
                dstNonNil = processInfo.FP.CreateModifier(_env, dest.FP.Get_nonnilableType(_env), Ast_MutMode__IMut)
                
            }
            if otherSrc.FP.Get_nilable(_env){
                if otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    return dstNonNil.FP.CanEvalWith(_env, processInfo, otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), canEvalType, alt2type)
                }
                return dstNonNil.FP.CanEvalWith(_env, processInfo, otherSrc.FP.Get_nonnilableType(_env), canEvalType, alt2type)
            }
            return dstNonNil.FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
        } else if Ast_isGenericType(_env, dest){
            return dest.FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
        } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( dest.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( dest.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ).(bool))) &&
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ).(bool))) ).(bool)){
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( canEvalType == Ast_CanEvalType__SetOp) ||
                _env.SetStackVal( canEvalType == Ast_CanEvalType__SetOpIMut) ).(bool){
                var result bool
                result = otherSrc.FP.IsInheritFrom(_env, processInfo, dest, alt2type)
                if result{
                    return result, nil
                }
                return false, _env.LuaVM.String_format("not inherit %s(%d) <- %s(%d)", []LnsAny{dest.FP.GetTxt(_env, nil, nil, nil), dest.FP.Get_typeId(_env).FP.Get_orgId(_env), otherSrc.FP.GetTxt(_env, nil, nil, nil), otherSrc.FP.Get_typeId(_env).FP.Get_orgId(_env)})
            }
            return false, ""
        } else if otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Enum{
            var enumTypeInfo *Ast_EnumTypeInfo
            enumTypeInfo = Lns_unwrap( (Ast_EnumTypeInfoDownCastF(otherSrc.FP))).(*Ast_EnumTypeInfo)
            return dest.FP.CanEvalWith(_env, processInfo, enumTypeInfo.FP.Get_valTypeInfo(_env), canEvalType, alt2type)
        } else if dest.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
            return dest.FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
        } else if dest.FP.Get_kind(_env) == Ast_TypeInfoKind__Box{
            return dest.FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
        } else if dest.FP.Get_kind(_env) == Ast_TypeInfoKind__Form{
            if _switch27941 := otherSrc.FP.Get_kind(_env); _switch27941 == Ast_TypeInfoKind__Form {
                return true, nil
            } else if _switch27941 == Ast_TypeInfoKind__FormFunc || _switch27941 == Ast_TypeInfoKind__Func {
                if Lns_car(Ast_isSettableToForm_3956_(_env, processInfo, otherSrc)).(bool){
                    return true, nil
                }
                return false, ""
            }
        } else if dest.FP.Get_kind(_env) == Ast_TypeInfoKind__FormFunc{
            if _switch28160 := otherSrc.FP.Get_kind(_env); _switch28160 == Ast_TypeInfoKind__FormFunc || _switch28160 == Ast_TypeInfoKind__Func {
                if otherSrc.FP.Get_accessMode(_env) == Ast_AccessMode__Direct{
                    return false, "can't set the _direct function."
                }
                {
                    var result LnsInt
                    var mess string
                    result,mess = Ast_TypeInfo_checkMatchType(_env, processInfo, dest.FP.Get_argTypeInfoList(_env), otherSrc.FP.Get_argTypeInfoList(_env), false, nil, alt2type)
                    if result == Ast_MatchType__Error{
                        return false, mess
                    }
                }
                {
                    var result LnsInt
                    var mess string
                    result,mess = Ast_TypeInfo_checkMatchType(_env, processInfo, otherSrc.FP.Get_argTypeInfoList(_env), dest.FP.Get_argTypeInfoList(_env), false, nil, alt2type)
                    if result == Ast_MatchType__Error{
                        return false, mess
                    }
                }
                {
                    var result LnsInt
                    var mess string
                    result,mess = Ast_TypeInfo_checkMatchType(_env, processInfo, dest.FP.Get_retTypeInfoList(_env), otherSrc.FP.Get_retTypeInfoList(_env), false, nil, alt2type)
                    if result == Ast_MatchType__Error{
                        return false, mess
                    }
                }
                {
                    var result LnsInt
                    var mess string
                    result,mess = Ast_TypeInfo_checkMatchType(_env, processInfo, otherSrc.FP.Get_retTypeInfoList(_env), dest.FP.Get_retTypeInfoList(_env), false, nil, alt2type)
                    if result == Ast_MatchType__Error{
                        return false, mess
                    }
                }
                if dest.FP.Get_retTypeInfoList(_env).Len() != otherSrc.FP.Get_retTypeInfoList(_env).Len(){
                    return false, _env.LuaVM.String_format("argNum %d != %d", []LnsAny{dest.FP.Get_retTypeInfoList(_env).Len(), otherSrc.FP.Get_retTypeInfoList(_env).Len()})
                }
                return true, nil
            }
        }
        return false, _env.LuaVM.String_format("illegal type -- %s, %s", []LnsAny{Ast_TypeInfoKind_getTxt( dest.FP.Get_kind(_env)), Ast_TypeInfoKind_getTxt( otherSrc.FP.Get_kind(_env))})
    }
    if _switch29211 := dest.FP.Get_kind(_env); _switch29211 == Ast_TypeInfoKind__Prim {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( dest == Ast_builtinTypeInt) &&
            _env.SetStackVal( otherSrc == Ast_builtinTypeChar) ||
            _env.SetStackVal( dest == Ast_builtinTypeChar) &&
            _env.SetStackVal( otherSrc == Ast_builtinTypeInt) ).(bool){
            return true, nil
        }
        return false, ""
    } else if _switch29211 == Ast_TypeInfoKind__List || _switch29211 == Ast_TypeInfoKind__Array || _switch29211 == Ast_TypeInfoKind__Set {
        if otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone{
            return true, nil
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( dest.FP.Get_itemTypeInfoList(_env).Len() >= 1) &&
            _env.SetStackVal( otherSrc.FP.Get_itemTypeInfoList(_env).Len() >= 1) ).(bool)){
            var ret bool
            var mess LnsAny
            ret,mess = (dest.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()).FP.CanEvalWith(_env, processInfo, otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( destMut) &&
                _env.SetStackVal( Ast_CanEvalType__SetEq) ||
                _env.SetStackVal( Ast_CanEvalType__SetOpIMut) ).(LnsInt), alt2type)
            if Lns_op_not(ret){
                return false, mess
            }
        } else { 
            return false, nil
        }
        
        return true, nil
    } else if _switch29211 == Ast_TypeInfoKind__Map {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone) &&
            _env.SetStackVal( otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone) ).(bool)){
            return true, nil
        }
        var check1 func(_env *LnsEnv) LnsAny
        check1 = func(_env *LnsEnv) LnsAny {
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( dest.FP.Get_itemTypeInfoList(_env).Len() >= 1) &&
                _env.SetStackVal( otherSrc.FP.Get_itemTypeInfoList(_env).Len() >= 1) ).(bool)){
                var ret bool
                ret,_ = (dest.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()).FP.CanEvalWith(_env, processInfo, otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( destMut) &&
                    _env.SetStackVal( Ast_CanEvalType__SetEq) ||
                    _env.SetStackVal( Ast_CanEvalType__SetOpIMut) ).(LnsInt), alt2type)
                if Lns_op_not(ret){
                    return false
                }
            } else { 
                return false
            }
            
            return true
        }
        var check2 func(_env *LnsEnv) bool
        check2 = func(_env *LnsEnv) bool {
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( dest.FP.Get_itemTypeInfoList(_env).Len() >= 2) &&
                _env.SetStackVal( otherSrc.FP.Get_itemTypeInfoList(_env).Len() >= 2) ).(bool)){
                var ret bool
                ret,_ = (dest.FP.Get_itemTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()).FP.CanEvalWith(_env, processInfo, otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( destMut) &&
                    _env.SetStackVal( Ast_CanEvalType__SetEq) ||
                    _env.SetStackVal( Ast_CanEvalType__SetOpIMut) ).(LnsInt), alt2type)
                if Lns_op_not(ret){
                    return false
                }
            } else { 
                return false
            }
            
            return true
        }
        var result1 LnsAny
        result1 = check1(_env)
        var result2 bool
        result2 = check2(_env)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( result1) &&
            _env.SetStackVal( result2) )){
            return true, nil
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(result1)) &&
            _env.SetStackVal( otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone) ||
            _env.SetStackVal( Lns_op_not(result2)) &&
            _env.SetStackVal( otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo() == Ast_builtinTypeNone) ).(bool){
            return true, nil
        }
        return false, ""
    } else if _switch29211 == Ast_TypeInfoKind__Class || _switch29211 == Ast_TypeInfoKind__IF {
        if Ast_isGenericType(_env, dest){
            return dest.FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( canEvalType == Ast_CanEvalType__SetOp) ||
            _env.SetStackVal( canEvalType == Ast_CanEvalType__SetOpIMut) ).(bool){
            var result bool
            result = otherSrc.FP.IsInheritFrom(_env, processInfo, dest, alt2type)
            if result{
                return result, nil
            }
            return false, _env.LuaVM.String_format("not inherit %s(%d) <- %s(%d)", []LnsAny{dest.FP.GetTxt(_env, nil, nil, nil), dest.FP.Get_typeId(_env).FP.Get_orgId(_env), otherSrc.FP.GetTxt(_env, nil, nil, nil), otherSrc.FP.Get_typeId(_env).FP.Get_orgId(_env)})
        }
        return false, ""
    } else if _switch29211 == Ast_TypeInfoKind__Form {
        if Lns_car(Ast_isSettableToForm_3956_(_env, processInfo, otherSrc)).(bool){
            return true, nil
        }
        return false, ""
    } else if _switch29211 == Ast_TypeInfoKind__Func || _switch29211 == Ast_TypeInfoKind__FormFunc {
        if dest.FP.Get_retTypeInfoList(_env).Len() != otherSrc.FP.Get_retTypeInfoList(_env).Len(){
            return false, _env.LuaVM.String_format("argNum %d != %d", []LnsAny{dest.FP.Get_retTypeInfoList(_env).Len(), otherSrc.FP.Get_retTypeInfoList(_env).Len()})
        }
        var argCheck LnsInt
        var argMess string
        argCheck,argMess = Ast_TypeInfo_checkMatchType(_env, processInfo, dest.FP.Get_argTypeInfoList(_env), otherSrc.FP.Get_argTypeInfoList(_env), false, nil, alt2type)
        if argCheck == Ast_MatchType__Error{
            return false, argMess
        }
        var retCheck LnsInt
        var retMess string
        retCheck,retMess = Ast_TypeInfo_checkMatchType(_env, processInfo, dest.FP.Get_retTypeInfoList(_env), otherSrc.FP.Get_retTypeInfoList(_env), false, nil, alt2type)
        if retCheck == Ast_MatchType__Error{
            return false, retMess
        }
        return true, nil
    } else if _switch29211 == Ast_TypeInfoKind__Method {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( dest.FP.Get_argTypeInfoList(_env).Len() != otherSrc.FP.Get_argTypeInfoList(_env).Len()) ||
            _env.SetStackVal( dest.FP.Get_retTypeInfoList(_env).Len() != otherSrc.FP.Get_retTypeInfoList(_env).Len()) ).(bool){
            return false, ""
        }
        for _index, _argType := range( dest.FP.Get_argTypeInfoList(_env).Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var otherArgType *Ast_TypeInfo
            otherArgType = otherSrc.FP.Get_argTypeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(argType.FP.Equals(_env, processInfo, otherArgType, alt2type, true)){
                var mess string
                mess = _env.LuaVM.String_format("unmatch arg(%d) type -- %s, %s", []LnsAny{index, argType.FP.GetTxt(_env, nil, nil, nil), otherArgType.FP.GetTxt(_env, nil, nil, nil)})
                return false, mess
            }
        }
        for _index, _retType := range( dest.FP.Get_retTypeInfoList(_env).Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var otherRetType *Ast_TypeInfo
            otherRetType = otherSrc.FP.Get_retTypeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(retType.FP.Equals(_env, processInfo, otherRetType, alt2type, true)){
                var mess string
                mess = _env.LuaVM.String_format("unmatch ret(%d) type -- %s, %s, %s", []LnsAny{index, retType.FP.GetTxt(_env, nil, nil, nil), otherRetType.FP.GetTxt(_env, nil, nil, nil), dest.FP.GetTxt(_env, nil, nil, nil)})
                return false, mess
            }
        }
        if dest.FP.Get_mutMode(_env) != otherSrc.FP.Get_mutMode(_env){
            var mess string
            mess = _env.LuaVM.String_format("unmatch mutable mode -- %s, %s", []LnsAny{Ast_MutMode_getTxt( dest.FP.Get_mutMode(_env)), Ast_MutMode_getTxt( otherSrc.FP.Get_mutMode(_env))})
            return false, mess
        }
        return true, nil
    } else if _switch29211 == Ast_TypeInfoKind__Nilable {
        var dstNonNil *Ast_TypeInfo
        if destMut{
            dstNonNil = dest.FP.Get_nonnilableType(_env)
            
        } else { 
            dstNonNil = processInfo.FP.CreateModifier(_env, dest.FP.Get_nonnilableType(_env), Ast_MutMode__IMut)
            
        }
        return dstNonNil.FP.CanEvalWith(_env, processInfo, otherSrc.FP.Get_nonnilableType(_env), canEvalType, alt2type)
    } else if _switch29211 == Ast_TypeInfoKind__Alternate {
        return dest.FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
    } else if _switch29211 == Ast_TypeInfoKind__Box {
        return dest.FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
    } else {
        return false, ""
    }
// insert a dummy
    return false,nil
}

// 7066: decl @lune.@base.@Ast.TypeInfo.getFullName
func (self *Ast_TypeInfo) GetFullName(_env *LnsEnv, typeNameCtrl *Ast_TypeNameCtrl,importInfo Ast_ModuleInfoManager,localFlag LnsAny) string {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( localFlag) &&
        _env.SetStackVal( self.FP.IsModule(_env)) )){
        return typeNameCtrl.FP.GetModuleName(_env, self, "", importInfo)
    }
    return self.FP.GetParentFullName(_env, typeNameCtrl, importInfo, localFlag) + self.FP.Get_rawTxt(_env)
}


// declaration Class -- RootTypeInfo
type Ast_RootTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_RootTypeInfo struct {
    Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_RootTypeInfo(_env *LnsEnv) *Ast_RootTypeInfo {
    obj := &Ast_RootTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_RootTypeInfo(_env)
    return obj
}
func (self *Ast_RootTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
// 1219: DeclConstr
func (self *Ast_RootTypeInfo) InitAst_RootTypeInfo(_env *LnsEnv) {
    self.InitAst_TypeInfo(_env, Ast_rootScope, Ast_rootProcessInfo)
    self.typeId = Ast_rootTypeIdInfo
    
}

// 1226: decl @lune.@base.@Ast.RootTypeInfo.get_baseTypeInfo
func (self *Ast_RootTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1230: decl @lune.@base.@Ast.RootTypeInfo.get_parentInfo
func (self *Ast_RootTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1234: decl @lune.@base.@Ast.RootTypeInfo.create
func Ast_RootTypeInfo_create_1746_(_env *LnsEnv) *Ast_RootTypeInfo {
    return NewAst_RootTypeInfo(_env)
}

// 1238: decl @lune.@base.@Ast.RootTypeInfo.get_rawTxt
func (self *Ast_RootTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return "<head>"
}

// 1242: decl @lune.@base.@Ast.RootTypeInfo.getTxt
func (self *Ast_RootTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return "<head>"
}


// declaration Class -- NormalSymbolInfo
type Ast_NormalSymbolInfoMtd interface {
    CanAccess(_env *LnsEnv, arg1 *Ast_Scope, arg2 LnsInt) LnsAny
    ClearValue(_env *LnsEnv)
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOrg(_env *LnsEnv) *Ast_SymbolInfo
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_canBeLeft(_env *LnsEnv) bool
    Get_canBeRight(_env *LnsEnv) bool
    Get_convModuleParam(_env *LnsEnv) LnsAny
    Get_hasAccessFromClosure(_env *LnsEnv) bool
    Get_hasValueFlag(_env *LnsEnv) bool
    Get_isLazyLoad(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_mutable(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) string
    Get_namespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_pos(_env *LnsEnv) LnsAny
    Get_posForLatestMod(_env *LnsEnv) LnsAny
    Get_posForModToRef(_env *LnsEnv) LnsAny
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_staticFlag(_env *LnsEnv) bool
    Get_symbolId(_env *LnsEnv) LnsInt
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    HasAccess(_env *LnsEnv) bool
    Set_convModuleParam(_env *LnsEnv, arg1 LnsAny)
    Set_hasAccessFromClosure(_env *LnsEnv, arg1 bool)
    Set_hasValueFlag(_env *LnsEnv, arg1 bool)
    set_namespaceTypeInfo(_env *LnsEnv, arg1 LnsAny)
    Set_posForLatestMod(_env *LnsEnv, arg1 LnsAny)
    Set_posForModToRef(_env *LnsEnv, arg1 LnsAny)
    Set_typeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo)
    UpdateValue(_env *LnsEnv, arg1 LnsAny)
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
func NewAst_NormalSymbolInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 bool, arg4 bool, arg5 *Ast_Scope, arg6 LnsInt, arg7 bool, arg8 string, arg9 LnsAny, arg10 *Ast_TypeInfo, arg11 LnsAny, arg12 bool, arg13 bool) *Ast_NormalSymbolInfo {
    obj := &Ast_NormalSymbolInfo{}
    obj.FP = obj
    obj.Ast_SymbolInfo.FP = obj
    obj.InitAst_NormalSymbolInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
    return obj
}
func (self *Ast_NormalSymbolInfo) Get_canBeLeft(_env *LnsEnv) bool{ return self.canBeLeft }
func (self *Ast_NormalSymbolInfo) Get_canBeRight(_env *LnsEnv) bool{ return self.canBeRight }
func (self *Ast_NormalSymbolInfo) Get_symbolId(_env *LnsEnv) LnsInt{ return self.symbolId }
func (self *Ast_NormalSymbolInfo) Get_scope(_env *LnsEnv) *Ast_Scope{ return self.scope }
func (self *Ast_NormalSymbolInfo) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Ast_NormalSymbolInfo) Get_staticFlag(_env *LnsEnv) bool{ return self.staticFlag }
func (self *Ast_NormalSymbolInfo) Get_isLazyLoad(_env *LnsEnv) bool{ return self.isLazyLoad }
func (self *Ast_NormalSymbolInfo) Get_name(_env *LnsEnv) string{ return self.name }
func (self *Ast_NormalSymbolInfo) Get_pos(_env *LnsEnv) LnsAny{ return self.pos }
func (self *Ast_NormalSymbolInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *Ast_NormalSymbolInfo) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *Ast_NormalSymbolInfo) Get_hasValueFlag(_env *LnsEnv) bool{ return self.hasValueFlag }
func (self *Ast_NormalSymbolInfo) Set_hasValueFlag(_env *LnsEnv, arg1 bool){ self.hasValueFlag = arg1 }
func (self *Ast_NormalSymbolInfo) Get_mutMode(_env *LnsEnv) LnsInt{ return self.mutMode }
func (self *Ast_NormalSymbolInfo) Get_hasAccessFromClosure(_env *LnsEnv) bool{ return self.hasAccessFromClosure }
func (self *Ast_NormalSymbolInfo) Set_hasAccessFromClosure(_env *LnsEnv, arg1 bool){ self.hasAccessFromClosure = arg1 }
func (self *Ast_NormalSymbolInfo) Get_posForLatestMod(_env *LnsEnv) LnsAny{ return self.posForLatestMod }
func (self *Ast_NormalSymbolInfo) Set_posForLatestMod(_env *LnsEnv, arg1 LnsAny){ self.posForLatestMod = arg1 }
func (self *Ast_NormalSymbolInfo) Get_posForModToRef(_env *LnsEnv) LnsAny{ return self.posForModToRef }
func (self *Ast_NormalSymbolInfo) Set_posForModToRef(_env *LnsEnv, arg1 LnsAny){ self.posForModToRef = arg1 }
func (self *Ast_NormalSymbolInfo) Get_convModuleParam(_env *LnsEnv) LnsAny{ return self.convModuleParam }
func (self *Ast_NormalSymbolInfo) Set_convModuleParam(_env *LnsEnv, arg1 LnsAny){ self.convModuleParam = arg1 }
// 1324: decl @lune.@base.@Ast.NormalSymbolInfo.get_mutable
func (self *Ast_NormalSymbolInfo) Get_mutable(_env *LnsEnv) bool {
    return Ast_isMutable(_env, self.mutMode)
}

// 1332: decl @lune.@base.@Ast.NormalSymbolInfo.getOrg
func (self *Ast_NormalSymbolInfo) GetOrg(_env *LnsEnv) *Ast_SymbolInfo {
    return &self.Ast_SymbolInfo
}

// 1336: DeclConstr
func (self *Ast_NormalSymbolInfo) InitAst_NormalSymbolInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,kind LnsInt,canBeLeft bool,canBeRight bool,scope *Ast_Scope,accessMode LnsInt,staticFlag bool,name string,pos LnsAny,typeInfo *Ast_TypeInfo,mutMode LnsAny,hasValueFlag bool,isLazyLoad bool) {
    self.InitAst_SymbolInfo(_env)
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
    
    self.symbolId = processInfo.FP.Get_idProvSym(_env).FP.GetNewId(_env)
    
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


// 2580: decl @lune.@base.@Ast.NormalSymbolInfo.canAccess
func (self *Ast_NormalSymbolInfo) CanAccess(_env *LnsEnv, fromScope *Ast_Scope,access LnsInt) LnsAny {
    if access == Ast_ScopeAccess__Full{
        return &self.Ast_SymbolInfo
    }
    if self.scope == fromScope{
        return &self.Ast_SymbolInfo
    }
    var processInfo *Ast_ProcessInfo
    processInfo = fromScope.FP.GetProcessInfo(_env)
    if _switch8742 := self.FP.Get_accessMode(_env); _switch8742 == Ast_AccessMode__Pub || _switch8742 == Ast_AccessMode__Global {
        return &self.Ast_SymbolInfo
    } else if _switch8742 == Ast_AccessMode__Pro {
        var nsClass *Ast_TypeInfo
        nsClass = self.scope.FP.GetClassTypeInfo(_env)
        var fromClass *Ast_TypeInfo
        fromClass = fromScope.FP.GetClassTypeInfo(_env)
        if fromClass.FP.IsInheritFrom(_env, processInfo, nsClass, nil){
            return &self.Ast_SymbolInfo
        }
        return nil
    } else if _switch8742 == Ast_AccessMode__Local {
        var selfMod *Ast_TypeInfo
        selfMod = self.FP.GetModule(_env)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( selfMod == Ast_headTypeInfo) ||
            _env.SetStackVal( selfMod.FP.Equals(_env, processInfo, fromScope.FP.GetModule(_env), nil, nil)) ).(bool){
            return &self.Ast_SymbolInfo
        }
        return nil
    } else if _switch8742 == Ast_AccessMode__Pri || _switch8742 == Ast_AccessMode__Direct {
        if fromScope.FP.IsInnerOf(_env, self.scope){
            return &self.Ast_SymbolInfo
        }
        return nil
    } else if _switch8742 == Ast_AccessMode__None {
        Util_err(_env, _env.LuaVM.String_format("illegl accessmode -- %s, %s", []LnsAny{self.FP.Get_accessMode(_env), self.FP.Get_name(_env)}))
    }
// insert a dummy
    return nil
}

// 4619: decl @lune.@base.@Ast.NormalSymbolInfo.set_typeInfo
func (self *Ast_NormalSymbolInfo) Set_typeInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo) {
    if self.name == "_"{
        return 
    }
    self.typeInfo = typeInfo
    
}


// declaration Class -- AutoBoxingInfo
var Ast_AutoBoxingInfo__allObj *LnsMap
// 1403: decl @lune.@base.@Ast.AutoBoxingInfo.___init
func Ast_AutoBoxingInfo____init_1834_(_env *LnsEnv) {
    Ast_AutoBoxingInfo__allObj = NewLnsMap( map[LnsAny]LnsAny{})
    
}

type Ast_AutoBoxingInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_count(_env *LnsEnv) LnsInt
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    Inc(_env *LnsEnv)
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
    Unregist(_env *LnsEnv)
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
func NewAst_AutoBoxingInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo) *Ast_AutoBoxingInfo {
    obj := &Ast_AutoBoxingInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AutoBoxingInfo(_env, arg1)
    return obj
}
func (self *Ast_AutoBoxingInfo) Get_count(_env *LnsEnv) LnsInt{ return self.count }
// 1408: DeclConstr
func (self *Ast_AutoBoxingInfo) InitAst_AutoBoxingInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo) {
    self.InitAst_TypeInfo(_env, nil, processInfo)
    self.count = 0
    
    Ast_AutoBoxingInfo__allObj.Set(&self.Ast_TypeInfo,self)
}

// 1420: decl @lune.@base.@Ast.AutoBoxingInfo.get_baseTypeInfo
func (self *Ast_AutoBoxingInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 1424: decl @lune.@base.@Ast.AutoBoxingInfo.get_parentInfo
func (self *Ast_AutoBoxingInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 1428: decl @lune.@base.@Ast.AutoBoxingInfo.get_kind
func (self *Ast_AutoBoxingInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Etc
}

// 1432: decl @lune.@base.@Ast.AutoBoxingInfo.inc
func (self *Ast_AutoBoxingInfo) Inc(_env *LnsEnv) {
    var obj *Ast_AutoBoxingInfo
    obj = Lns_unwrap( Ast_AutoBoxingInfo__allObj.Get(&self.Ast_TypeInfo)).(*Ast_AutoBoxingInfo)
    obj.count = obj.count + 1
    
}

// 1437: decl @lune.@base.@Ast.AutoBoxingInfo.unregist
func (self *Ast_AutoBoxingInfo) Unregist(_env *LnsEnv) {
    Ast_AutoBoxingInfo__allObj.Set(&self.Ast_TypeInfo,nil)
}


// declaration Class -- CanEvalCtrlTypeInfo
var Ast_CanEvalCtrlTypeInfo__detectAlt *Ast_CanEvalCtrlTypeInfo
var Ast_CanEvalCtrlTypeInfo__needAutoBoxing *Ast_CanEvalCtrlTypeInfo
var Ast_CanEvalCtrlTypeInfo__checkTypeTarget *Ast_CanEvalCtrlTypeInfo
// 1443: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.___init
func Ast_CanEvalCtrlTypeInfo____init_1856_(_env *LnsEnv) {
    Ast_CanEvalCtrlTypeInfo__detectAlt = NewAst_CanEvalCtrlTypeInfo(_env)
    
    Ast_CanEvalCtrlTypeInfo__needAutoBoxing = NewAst_CanEvalCtrlTypeInfo(_env)
    
    Ast_CanEvalCtrlTypeInfo__checkTypeTarget = NewAst_CanEvalCtrlTypeInfo(_env)
    
}

type Ast_CanEvalCtrlTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
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
func NewAst_CanEvalCtrlTypeInfo(_env *LnsEnv) *Ast_CanEvalCtrlTypeInfo {
    obj := &Ast_CanEvalCtrlTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_CanEvalCtrlTypeInfo(_env)
    return obj
}
// 1448: DeclConstr
func (self *Ast_CanEvalCtrlTypeInfo) InitAst_CanEvalCtrlTypeInfo(_env *LnsEnv) {
    self.InitAst_TypeInfo(_env, nil, Ast_rootProcessInfo)
}

// 1457: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.get_kind
func (self *Ast_CanEvalCtrlTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__CanEvalCtrl
}

// 1461: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.get_typeId
func (self *Ast_CanEvalCtrlTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo {
    return Ast_dummyIdInfo
}

// 1467: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.get_baseTypeInfo
func (self *Ast_CanEvalCtrlTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 1471: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.get_parentInfo
func (self *Ast_CanEvalCtrlTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 1475: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.createDefaultAlt2typeMap
func Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env *LnsEnv, detectFlag bool) *LnsMap {
    if detectFlag{
        var _map *LnsMap
        _map = NewLnsMap( map[LnsAny]LnsAny{})
        _map.Set(&Ast_CanEvalCtrlTypeInfo__detectAlt.Ast_TypeInfo,Ast_headTypeInfo)
        return _map
    }
    return NewLnsMap( map[LnsAny]LnsAny{})
}

// 1484: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.isValidApply
func Ast_CanEvalCtrlTypeInfo_isValidApply(_env *LnsEnv, alt2type *LnsMap) bool {
    return Ast_TypeInfo2Stem(alt2type.Get(&Ast_CanEvalCtrlTypeInfo__detectAlt.Ast_TypeInfo)) != nil
}

// 1488: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.setupNeedAutoBoxing
func Ast_CanEvalCtrlTypeInfo_setupNeedAutoBoxing(_env *LnsEnv, alt2type *LnsMap,processInfo *Ast_ProcessInfo) {
    alt2type.Set(&Ast_CanEvalCtrlTypeInfo__needAutoBoxing.Ast_TypeInfo,&NewAst_AutoBoxingInfo(_env, processInfo).Ast_TypeInfo)
}

// 1494: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.updateNeedAutoBoxing
func Ast_CanEvalCtrlTypeInfo_updateNeedAutoBoxing(_env *LnsEnv, alt2type *LnsMap) {
    {
        __exp := alt2type.Get(&Ast_CanEvalCtrlTypeInfo__needAutoBoxing.Ast_TypeInfo)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            {
                _autoBoxingInfo := Ast_AutoBoxingInfoDownCastF(_exp.FP)
                if !Lns_IsNil( _autoBoxingInfo ) {
                    autoBoxingInfo := _autoBoxingInfo.(*Ast_AutoBoxingInfo)
                    autoBoxingInfo.FP.Inc(_env)
                }
            }
        } else {
            Util_err(_env, "no exist needAutoBoxing")
        }
    }
}

// 1505: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.hasNeedAutoBoxing
func Ast_CanEvalCtrlTypeInfo_hasNeedAutoBoxing(_env *LnsEnv, alt2type *LnsMap) bool {
    {
        __exp := alt2type.Get(&Ast_CanEvalCtrlTypeInfo__needAutoBoxing.Ast_TypeInfo)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            {
                _autoBoxingInfo := Ast_AutoBoxingInfoDownCastF(_exp.FP)
                if !Lns_IsNil( _autoBoxingInfo ) {
                    autoBoxingInfo := _autoBoxingInfo.(*Ast_AutoBoxingInfo)
                    return autoBoxingInfo.FP.Get_count(_env) != 0
                }
            }
        }
    }
    return false
}

// 1515: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.finishNeedAutoBoxing
func Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(_env *LnsEnv, alt2type *LnsMap,count LnsInt) bool {
    {
        __exp := alt2type.Get(&Ast_CanEvalCtrlTypeInfo__needAutoBoxing.Ast_TypeInfo)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            {
                _autoBoxingInfo := Ast_AutoBoxingInfoDownCastF(_exp.FP)
                if !Lns_IsNil( _autoBoxingInfo ) {
                    autoBoxingInfo := _autoBoxingInfo.(*Ast_AutoBoxingInfo)
                    autoBoxingInfo.FP.Unregist(_env)
                    return autoBoxingInfo.FP.Get_count(_env) == count
                }
            }
        }
    }
    return false
}

// 1532: decl @lune.@base.@Ast.CanEvalCtrlTypeInfo.canAutoBoxing
func Ast_CanEvalCtrlTypeInfo_canAutoBoxing(_env *LnsEnv, dst *Ast_TypeInfo,src *Ast_TypeInfo) bool {
    var dstSrc *Ast_TypeInfo
    dstSrc = dst.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)
    if dstSrc.FP.Get_kind(_env) != Ast_TypeInfoKind__Box{
        return false
    }
    var srcSrc *Ast_TypeInfo
    srcSrc = src.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)
    if srcSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Box{
        return false
    }
    return true
}


// declaration Class -- NilableTypeInfo
type Ast_NilableTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_NilableTypeInfo struct {
    Ast_TypeInfo
    nonnilableType *Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_NilableTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo) *Ast_NilableTypeInfo {
    obj := &Ast_NilableTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_NilableTypeInfo(_env, arg1, arg2)
    return obj
}
func (self *Ast_NilableTypeInfo) Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo{ return self.nonnilableType }
func (self *Ast_NilableTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
// advertise -- 1546
func (self *Ast_NilableTypeInfo) AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo) {
self.nonnilableType. FP.AddChildren( _env, arg1)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap {
    return self.nonnilableType. FP.CreateAlt2typeMap( _env, arg1)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.nonnilableType. FP.GetFullName( _env, arg1,arg2,arg3)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.nonnilableType. FP.GetModule( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) GetOverridingType(_env *LnsEnv) LnsAny {
    return self.nonnilableType. FP.GetOverridingType( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.nonnilableType. FP.GetParentFullName( _env, arg1,arg2,arg3)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.nonnilableType. FP.GetParentId( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.nonnilableType. FP.getProcessInfo( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_abstractFlag(_env *LnsEnv) bool {
    return self.nonnilableType. FP.Get_abstractFlag( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return self.nonnilableType. FP.Get_accessMode( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_argTypeInfoList(_env *LnsEnv) *LnsList {
    return self.nonnilableType. FP.Get_argTypeInfoList( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_asyncMode(_env *LnsEnv) LnsInt {
    return self.nonnilableType. FP.Get_asyncMode( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_autoFlag(_env *LnsEnv) bool {
    return self.nonnilableType. FP.Get_autoFlag( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return self.nonnilableType. FP.Get_baseId( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_baseTypeInfo( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_children(_env *LnsEnv) *LnsList {
    return self.nonnilableType. FP.Get_children( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_extedType(_env *LnsEnv) *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_extedType( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_externalFlag(_env *LnsEnv) bool {
    return self.nonnilableType. FP.Get_externalFlag( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_genSrcTypeInfo( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList {
    return self.nonnilableType. FP.Get_interfaceList( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList {
    return self.nonnilableType. FP.Get_itemTypeInfoList( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return self.nonnilableType. FP.Get_mutMode( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_nilableTypeInfo( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.nonnilableType. FP.Get_parentInfo( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.nonnilableType. FP.Get_processInfo( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return self.nonnilableType. FP.Get_rawTxt( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList {
    return self.nonnilableType. FP.Get_retTypeInfoList( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_scope(_env *LnsEnv) LnsAny {
    return self.nonnilableType. FP.Get_scope( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_staticFlag(_env *LnsEnv) bool {
    return self.nonnilableType. FP.Get_staticFlag( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) Get_typeData(_env *LnsEnv) *Ast_TypeData {
    return self.nonnilableType. FP.Get_typeData( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) HasBase(_env *LnsEnv) bool {
    return self.nonnilableType. FP.HasBase( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool {
    return self.nonnilableType. FP.HasRouteNamespaceFrom( _env, arg1)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.nonnilableType. FP.IsInheritFrom( _env, arg1,arg2,arg3)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) IsModule(_env *LnsEnv) bool {
    return self.nonnilableType. FP.IsModule( _env)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo,arg2 string,arg3 *LnsList,arg4 LnsAny) string {
    return self.nonnilableType. FP.SerializeTypeInfoList( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 1546
func (self *Ast_NilableTypeInfo) SwitchScope(_env *LnsEnv, arg1 *Ast_Scope) {
self.nonnilableType. FP.SwitchScope( _env, arg1)
}
// 1550: DeclConstr
func (self *Ast_NilableTypeInfo) InitAst_NilableTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,nonnilableType *Ast_TypeInfo) {
    self.InitAst_TypeInfo(_env, nil, processInfo)
    self.nonnilableType = nonnilableType
    
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
}

// 1556: decl @lune.@base.@Ast.NilableTypeInfo.get_kind
func (self *Ast_NilableTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Nilable
}

// 1559: decl @lune.@base.@Ast.NilableTypeInfo.get_aliasSrc
func (self *Ast_NilableTypeInfo) Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1562: decl @lune.@base.@Ast.NilableTypeInfo.get_srcTypeInfo
func (self *Ast_NilableTypeInfo) Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1565: decl @lune.@base.@Ast.NilableTypeInfo.get_nilable
func (self *Ast_NilableTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return true
}

// 1570: decl @lune.@base.@Ast.NilableTypeInfo.getTxt
func (self *Ast_NilableTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 1575: decl @lune.@base.@Ast.NilableTypeInfo.getTxtWithRaw
func (self *Ast_NilableTypeInfo) GetTxtWithRaw(_env *LnsEnv, raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.nonnilableType.FP.GetTxtWithRaw(_env, raw, typeNameCtrl, importInfo, localFlag) + "!"
}

// 1584: decl @lune.@base.@Ast.NilableTypeInfo.get_display_stirng_with
func (self *Ast_NilableTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    if self.nonnilableType.FP.Get_kind(_env) == Ast_TypeInfoKind__FormFunc{
        return self.nonnilableType.FP.Get_display_stirng_with(_env, raw + "!", alt2type)
    }
    return self.nonnilableType.FP.Get_display_stirng_with(_env, raw, alt2type) + "!"
}

// 1592: decl @lune.@base.@Ast.NilableTypeInfo.get_display_stirng
func (self *Ast_NilableTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
}

// 1596: decl @lune.@base.@Ast.NilableTypeInfo.serialize
func (self *Ast_NilableTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    stream.Write(_env, _env.LuaVM.String_format("{ skind = %d, typeId = %d, orgTypeId = %s }\n", []LnsAny{Ast_SerializeKind__Nilable, self.typeId.Id, serializeInfo.FP.SerializeId(_env, self.nonnilableType.FP.Get_typeId(_env))}))
}

// 1603: decl @lune.@base.@Ast.NilableTypeInfo.equals
func (self *Ast_NilableTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if Lns_op_not(typeInfo.FP.Get_nilable(_env)){
        return false
    }
    return self.nonnilableType.FP.Equals(_env, processInfo, typeInfo.FP.Get_nonnilableType(_env), alt2type, checkModifer)
}

// 1614: decl @lune.@base.@Ast.NilableTypeInfo.applyGeneric
func (self *Ast_NilableTypeInfo) ApplyGeneric(_env *LnsEnv, alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.nonnilableType.FP.ApplyGeneric(_env, alt2typeMap, moduleTypeInfo)
    if typeInfo == self.nonnilableType{
        return &self.Ast_TypeInfo
    }
    if typeInfo != nil{
        typeInfo_1281 := typeInfo.(*Ast_TypeInfo)
        return typeInfo_1281.FP.Get_nilableTypeInfo(_env)
    }
    return nil
}

// 6124: decl @lune.@base.@Ast.NilableTypeInfo.canEvalWith
func (self *Ast_NilableTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    var otherSrc *Ast_TypeInfo
    otherSrc = other
    if &self.Ast_TypeInfo == Ast_builtinTypeStem_{
        return true, nil
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( otherSrc == Ast_builtinTypeNil) ||
        _env.SetStackVal( otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr) ).(bool){
        if self.FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Box{
            return self.FP.Get_nonnilableType(_env).FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
        }
        return true, nil
    }
    if self.typeId.FP.Equals(_env, otherSrc.FP.Get_typeId(_env)){
        return true, nil
    }
    if otherSrc.FP.Get_nilable(_env){
        if otherSrc.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            return self.FP.Get_nonnilableType(_env).FP.CanEvalWith(_env, processInfo, otherSrc.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), canEvalType, alt2type)
        }
        return self.FP.Get_nonnilableType(_env).FP.CanEvalWith(_env, processInfo, otherSrc.FP.Get_nonnilableType(_env), canEvalType, alt2type)
    }
    return self.FP.Get_nonnilableType(_env).FP.CanEvalWith(_env, processInfo, otherSrc, canEvalType, alt2type)
}


// declaration Class -- AliasTypeInfo
type Ast_AliasTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_aliasSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_AliasTypeInfo struct {
    Ast_TypeInfo
    rawTxt string
    accessMode LnsInt
    parentInfo *Ast_TypeInfo
    aliasSrcTypeInfo *Ast_TypeInfo
    externalFlag bool
    typeId *Ast_IdInfo
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
func NewAst_AliasTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 LnsInt, arg4 *Ast_TypeInfo, arg5 *Ast_TypeInfo, arg6 bool) *Ast_AliasTypeInfo {
    obj := &Ast_AliasTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AliasTypeInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Ast_AliasTypeInfo) Get_rawTxt(_env *LnsEnv) string{ return self.rawTxt }
func (self *Ast_AliasTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Ast_AliasTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_AliasTypeInfo) Get_aliasSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.aliasSrcTypeInfo }
func (self *Ast_AliasTypeInfo) Get_externalFlag(_env *LnsEnv) bool{ return self.externalFlag }
func (self *Ast_AliasTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_AliasTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.nilableTypeInfo }
// advertise -- 1631
func (self *Ast_AliasTypeInfo) AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo) {
self.aliasSrcTypeInfo. FP.AddChildren( _env, arg1)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap {
    return self.aliasSrcTypeInfo. FP.CreateAlt2typeMap( _env, arg1)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) GetOverridingType(_env *LnsEnv) LnsAny {
    return self.aliasSrcTypeInfo. FP.GetOverridingType( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.aliasSrcTypeInfo. FP.getProcessInfo( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) GetTxtWithRaw(_env *LnsEnv, arg1 string,arg2 LnsAny,arg3 LnsAny,arg4 LnsAny) string {
    return self.aliasSrcTypeInfo. FP.GetTxtWithRaw( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_abstractFlag(_env *LnsEnv) bool {
    return self.aliasSrcTypeInfo. FP.Get_abstractFlag( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_argTypeInfoList(_env *LnsEnv) *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_argTypeInfoList( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_asyncMode(_env *LnsEnv) LnsInt {
    return self.aliasSrcTypeInfo. FP.Get_asyncMode( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_autoFlag(_env *LnsEnv) bool {
    return self.aliasSrcTypeInfo. FP.Get_autoFlag( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return self.aliasSrcTypeInfo. FP.Get_baseId( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.aliasSrcTypeInfo. FP.Get_baseTypeInfo( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_children(_env *LnsEnv) *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_children( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_display_stirng_with(_env *LnsEnv, arg1 string,arg2 LnsAny) string {
    return self.aliasSrcTypeInfo. FP.Get_display_stirng_with( _env, arg1,arg2)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_extedType(_env *LnsEnv) *Ast_TypeInfo {
    return self.aliasSrcTypeInfo. FP.Get_extedType( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_interfaceList( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_itemTypeInfoList( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return self.aliasSrcTypeInfo. FP.Get_kind( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return self.aliasSrcTypeInfo. FP.Get_mutMode( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return self.aliasSrcTypeInfo. FP.Get_nilable( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.aliasSrcTypeInfo. FP.Get_processInfo( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList {
    return self.aliasSrcTypeInfo. FP.Get_retTypeInfoList( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_scope(_env *LnsEnv) LnsAny {
    return self.aliasSrcTypeInfo. FP.Get_scope( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_staticFlag(_env *LnsEnv) bool {
    return self.aliasSrcTypeInfo. FP.Get_staticFlag( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) Get_typeData(_env *LnsEnv) *Ast_TypeData {
    return self.aliasSrcTypeInfo. FP.Get_typeData( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) HasBase(_env *LnsEnv) bool {
    return self.aliasSrcTypeInfo. FP.HasBase( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool {
    return self.aliasSrcTypeInfo. FP.HasRouteNamespaceFrom( _env, arg1)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.aliasSrcTypeInfo. FP.IsInheritFrom( _env, arg1,arg2,arg3)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) IsModule(_env *LnsEnv) bool {
    return self.aliasSrcTypeInfo. FP.IsModule( _env)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo,arg2 string,arg3 *LnsList,arg4 LnsAny) string {
    return self.aliasSrcTypeInfo. FP.SerializeTypeInfoList( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 1631
func (self *Ast_AliasTypeInfo) SwitchScope(_env *LnsEnv, arg1 *Ast_Scope) {
self.aliasSrcTypeInfo. FP.SwitchScope( _env, arg1)
}
// 1640: DeclConstr
func (self *Ast_AliasTypeInfo) InitAst_AliasTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,rawTxt string,accessMode LnsInt,parentInfo *Ast_TypeInfo,aliasSrcTypeInfo *Ast_TypeInfo,externalFlag bool) {
    self.InitAst_TypeInfo(_env, nil, processInfo)
    self.rawTxt = rawTxt
    
    self.accessMode = accessMode
    
    self.parentInfo = parentInfo
    
    self.aliasSrcTypeInfo = aliasSrcTypeInfo
    
    self.externalFlag = externalFlag
    
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.nilableTypeInfo = &NewAst_NilableTypeInfo(_env, processInfo, &self.Ast_TypeInfo).Ast_TypeInfo
    
}

// 1654: decl @lune.@base.@Ast.AliasTypeInfo.getParentFullName
func (self *Ast_AliasTypeInfo) GetParentFullName(_env *LnsEnv, typeNameCtrl *Ast_TypeNameCtrl,importInfo LnsAny,localFlag LnsAny) string {
    return typeNameCtrl.FP.GetParentFullName(_env, &self.Ast_TypeInfo, importInfo, localFlag)
}

// 1661: decl @lune.@base.@Ast.AliasTypeInfo.getFullName
func (self *Ast_AliasTypeInfo) GetFullName(_env *LnsEnv, typeNameCtrl *Ast_TypeNameCtrl,importInfo Ast_ModuleInfoManager,localFlag LnsAny) string {
    return self.FP.GetParentFullName(_env, typeNameCtrl, importInfo, localFlag) + self.FP.Get_rawTxt(_env)
}

// 1669: decl @lune.@base.@Ast.AliasTypeInfo.get_aliasSrc
func (self *Ast_AliasTypeInfo) Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo {
    return self.aliasSrcTypeInfo
}

// 1673: decl @lune.@base.@Ast.AliasTypeInfo.get_nonnilableType
func (self *Ast_AliasTypeInfo) Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1677: decl @lune.@base.@Ast.AliasTypeInfo.get_srcTypeInfo
func (self *Ast_AliasTypeInfo) Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1680: decl @lune.@base.@Ast.AliasTypeInfo.get_genSrcTypeInfo
func (self *Ast_AliasTypeInfo) Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 1684: decl @lune.@base.@Ast.AliasTypeInfo.getModule
func (self *Ast_AliasTypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.FP.Get_parentInfo(_env).FP.GetModule(_env)
}

// 1689: decl @lune.@base.@Ast.AliasTypeInfo.getTxt
func (self *Ast_AliasTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.rawTxt, typeNameCtrl, importInfo, localFlag)
}

// 1695: decl @lune.@base.@Ast.AliasTypeInfo.serialize
func (self *Ast_AliasTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    var parentId *Ast_IdInfo
    parentId = self.FP.GetParentId(_env)
    stream.Write(_env, _env.LuaVM.String_format("{ skind = %d, parentId = %d, typeId = %d, rawTxt = %q, srcTypeId = %s }\n", []LnsAny{Ast_SerializeKind__Alias, parentId.Id, self.typeId.Id, self.rawTxt, serializeInfo.FP.SerializeId(_env, self.aliasSrcTypeInfo.FP.Get_typeId(_env))}))
}

// 1705: decl @lune.@base.@Ast.AliasTypeInfo.get_display_stirng
func (self *Ast_AliasTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.rawTxt, nil)
}

// 1710: decl @lune.@base.@Ast.AliasTypeInfo.getParentId
func (self *Ast_AliasTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.parentInfo.FP.Get_typeId(_env)
}

// 1714: decl @lune.@base.@Ast.AliasTypeInfo.applyGeneric
func (self *Ast_AliasTypeInfo) ApplyGeneric(_env *LnsEnv, alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.aliasSrcTypeInfo.FP.ApplyGeneric(_env, alt2typeMap, moduleTypeInfo)
    if typeInfo == self.aliasSrcTypeInfo{
        return &self.Ast_TypeInfo
    }
    return nil
}

// 1724: decl @lune.@base.@Ast.AliasTypeInfo.canEvalWith
func (self *Ast_AliasTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return self.aliasSrcTypeInfo.FP.CanEvalWith(_env, processInfo, other.FP.Get_aliasSrc(_env), canEvalType, alt2type)
}

// 1732: decl @lune.@base.@Ast.AliasTypeInfo.equals
func (self *Ast_AliasTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    return self.aliasSrcTypeInfo.FP.Equals(_env, processInfo, typeInfo.FP.Get_aliasSrc(_env), alt2type, checkModifer)
}


// declaration Class -- NilTypeInfo
type Ast_NilTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_NilTypeInfo struct {
    Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_NilTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo) *Ast_NilTypeInfo {
    obj := &Ast_NilTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_NilTypeInfo(_env, arg1)
    return obj
}
func (self *Ast_NilTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
// 2459: DeclConstr
func (self *Ast_NilTypeInfo) InitAst_NilTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo) {
    self.InitAst_TypeInfo(_env, nil, processInfo)
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
}

// 2466: decl @lune.@base.@Ast.NilTypeInfo.isModule
func (self *Ast_NilTypeInfo) IsModule(_env *LnsEnv) bool {
    return false
}

// 2471: decl @lune.@base.@Ast.NilTypeInfo.getTxt
func (self *Ast_NilTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 2477: decl @lune.@base.@Ast.NilTypeInfo.getTxtWithRaw
func (self *Ast_NilTypeInfo) GetTxtWithRaw(_env *LnsEnv, raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return "nil"
}

// 2486: decl @lune.@base.@Ast.NilTypeInfo.canEvalWith
func (self *Ast_NilTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if Lns_op_not(other.FP.Get_nilable(_env)){
        return false, _env.LuaVM.String_format("%s is not nilable.", []LnsAny{other.FP.GetTxt(_env, nil, nil, nil)})
    }
    return other.FP.Get_nilable(_env), nil
}

// 2497: decl @lune.@base.@Ast.NilTypeInfo.get_display_stirng_with
func (self *Ast_NilTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, raw, nil, nil, nil)
}

// 2501: decl @lune.@base.@Ast.NilTypeInfo.get_display_stirng
func (self *Ast_NilTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, "nil", nil)
}

// 2507: decl @lune.@base.@Ast.NilTypeInfo.equals
func (self *Ast_NilTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    return &self.Ast_TypeInfo == typeInfo
}

// 2515: decl @lune.@base.@Ast.NilTypeInfo.get_parentInfo
func (self *Ast_NilTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 2523: decl @lune.@base.@Ast.NilTypeInfo.hasRouteNamespaceFrom
func (self *Ast_NilTypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, other *Ast_TypeInfo) bool {
    return true
}

// 2528: decl @lune.@base.@Ast.NilTypeInfo.get_rawTxt
func (self *Ast_NilTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return "nil"
}

// 2532: decl @lune.@base.@Ast.NilTypeInfo.get_kind
func (self *Ast_NilTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Prim
}

// 2536: decl @lune.@base.@Ast.NilTypeInfo.get_baseTypeInfo
func (self *Ast_NilTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 2540: decl @lune.@base.@Ast.NilTypeInfo.get_nilable
func (self *Ast_NilTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return true
}

// 2544: decl @lune.@base.@Ast.NilTypeInfo.get_mutMode
func (self *Ast_NilTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return Ast_MutMode__IMut
}

// 2549: decl @lune.@base.@Ast.NilTypeInfo.getParentFullName
func (self *Ast_NilTypeInfo) GetParentFullName(_env *LnsEnv, typeNameCtrl *Ast_TypeNameCtrl,importInfo LnsAny,localFlag LnsAny) string {
    return ""
}


// declaration Class -- AccessSymbolInfo
type Ast_AccessSymbolInfoMtd interface {
    CanAccess(_env *LnsEnv, arg1 *Ast_Scope, arg2 LnsInt) LnsAny
    ClearValue(_env *LnsEnv)
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOrg(_env *LnsEnv) *Ast_SymbolInfo
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_canBeLeft(_env *LnsEnv) bool
    Get_canBeRight(_env *LnsEnv) bool
    Get_convModuleParam(_env *LnsEnv) LnsAny
    Get_hasAccessFromClosure(_env *LnsEnv) bool
    Get_hasValueFlag(_env *LnsEnv) bool
    Get_isLazyLoad(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_mutable(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) string
    Get_namespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_pos(_env *LnsEnv) LnsAny
    Get_posForLatestMod(_env *LnsEnv) LnsAny
    Get_posForModToRef(_env *LnsEnv) LnsAny
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_staticFlag(_env *LnsEnv) bool
    Get_symbolId(_env *LnsEnv) LnsInt
    Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    HasAccess(_env *LnsEnv) bool
    Set_convModuleParam(_env *LnsEnv, arg1 LnsAny)
    Set_hasAccessFromClosure(_env *LnsEnv, arg1 bool)
    Set_hasValueFlag(_env *LnsEnv, arg1 bool)
    set_namespaceTypeInfo(_env *LnsEnv, arg1 LnsAny)
    Set_posForLatestMod(_env *LnsEnv, arg1 LnsAny)
    Set_posForModToRef(_env *LnsEnv, arg1 LnsAny)
    Set_typeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo)
    UpdateValue(_env *LnsEnv, arg1 LnsAny)
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
func NewAst_AccessSymbolInfo(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 LnsAny, arg3 bool) *Ast_AccessSymbolInfo {
    obj := &Ast_AccessSymbolInfo{}
    obj.FP = obj
    obj.Ast_SymbolInfo.FP = obj
    obj.InitAst_AccessSymbolInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Ast_AccessSymbolInfo) Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbolInfo }
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) ClearValue(_env *LnsEnv) {
self.symbolInfo. FP.ClearValue( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.symbolInfo. FP.GetModule( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return self.symbolInfo. FP.Get_accessMode( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_canBeRight(_env *LnsEnv) bool {
    return self.symbolInfo. FP.Get_canBeRight( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_convModuleParam(_env *LnsEnv) LnsAny {
    return self.symbolInfo. FP.Get_convModuleParam( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_hasAccessFromClosure(_env *LnsEnv) bool {
    return self.symbolInfo. FP.Get_hasAccessFromClosure( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_hasValueFlag(_env *LnsEnv) bool {
    return self.symbolInfo. FP.Get_hasValueFlag( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_isLazyLoad(_env *LnsEnv) bool {
    return self.symbolInfo. FP.Get_isLazyLoad( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_kind(_env *LnsEnv) LnsInt {
    return self.symbolInfo. FP.Get_kind( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_name(_env *LnsEnv) string {
    return self.symbolInfo. FP.Get_name( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_namespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.symbolInfo. FP.Get_namespaceTypeInfo( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_pos(_env *LnsEnv) LnsAny {
    return self.symbolInfo. FP.Get_pos( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_posForLatestMod(_env *LnsEnv) LnsAny {
    return self.symbolInfo. FP.Get_posForLatestMod( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_posForModToRef(_env *LnsEnv) LnsAny {
    return self.symbolInfo. FP.Get_posForModToRef( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_scope(_env *LnsEnv) *Ast_Scope {
    return self.symbolInfo. FP.Get_scope( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_staticFlag(_env *LnsEnv) bool {
    return self.symbolInfo. FP.Get_staticFlag( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Get_symbolId(_env *LnsEnv) LnsInt {
    return self.symbolInfo. FP.Get_symbolId( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) HasAccess(_env *LnsEnv) bool {
    return self.symbolInfo. FP.HasAccess( _env)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Set_convModuleParam(_env *LnsEnv, arg1 LnsAny) {
self.symbolInfo. FP.Set_convModuleParam( _env, arg1)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Set_hasAccessFromClosure(_env *LnsEnv, arg1 bool) {
self.symbolInfo. FP.Set_hasAccessFromClosure( _env, arg1)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Set_hasValueFlag(_env *LnsEnv, arg1 bool) {
self.symbolInfo. FP.Set_hasValueFlag( _env, arg1)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) set_namespaceTypeInfo(_env *LnsEnv, arg1 LnsAny) {
self.symbolInfo. FP.set_namespaceTypeInfo( _env, arg1)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Set_posForLatestMod(_env *LnsEnv, arg1 LnsAny) {
self.symbolInfo. FP.Set_posForLatestMod( _env, arg1)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Set_posForModToRef(_env *LnsEnv, arg1 LnsAny) {
self.symbolInfo. FP.Set_posForModToRef( _env, arg1)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) Set_typeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) {
self.symbolInfo. FP.Set_typeInfo( _env, arg1)
}
// advertise -- 2643
func (self *Ast_AccessSymbolInfo) UpdateValue(_env *LnsEnv, arg1 LnsAny) {
self.symbolInfo. FP.UpdateValue( _env, arg1)
}
// 2649: DeclConstr
func (self *Ast_AccessSymbolInfo) InitAst_AccessSymbolInfo(_env *LnsEnv, symbolInfo *Ast_SymbolInfo,overrideMut LnsAny,overrideCanBeLeft bool) {
    self.InitAst_SymbolInfo(_env)
    self.symbolInfo = symbolInfo
    
    self.overrideMut = overrideMut
    
    self.overrideCanBeLeft = overrideCanBeLeft
    
    var symType *Ast_TypeInfo
    symType = symbolInfo.FP.Get_typeInfo(_env)
    var work *Ast_TypeInfo
    switch _exp8915 := self.overrideMut.(type) {
    case *Ast_OverrideMut__None:
        work = symType
        
    case *Ast_OverrideMut__Prefix:
    prefixTypeInfo := _exp8915.Val1
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) &&
            _env.SetStackVal( symType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
            _env.SetStackVal( prefixTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
            _env.SetStackVal( prefixTypeInfo.FP.Get_itemTypeInfoList(_env).Len() > 0) ).(bool)){
            var alt2TypeMap *LnsMap
            alt2TypeMap = prefixTypeInfo.FP.CreateAlt2typeMap(_env, false)
            var typeInfo LnsAny
            typeInfo = symType.FP.ApplyGeneric(_env, alt2TypeMap, symType.FP.GetModule(_env))
            if typeInfo != nil{
                typeInfo_2073 := typeInfo.(*Ast_TypeInfo)
                work = typeInfo_2073
                
            } else {
                work = symType
                
            }
        } else { 
            work = symType
            
        }
    case *Ast_OverrideMut__IMut:
    typeInfo := _exp8915.Val1
        work = typeInfo
        
    }
    self.overrideTypeInfo = work
    
}

// 2688: decl @lune.@base.@Ast.AccessSymbolInfo.getOrg
func (self *Ast_AccessSymbolInfo) GetOrg(_env *LnsEnv) *Ast_SymbolInfo {
    return self.symbolInfo.FP.GetOrg(_env)
}

// 2692: decl @lune.@base.@Ast.AccessSymbolInfo.canAccess
func (self *Ast_AccessSymbolInfo) CanAccess(_env *LnsEnv, fromScope *Ast_Scope,access LnsInt) LnsAny {
    if Lns_isCondTrue( self.symbolInfo.FP.CanAccess(_env, fromScope, access)){
        return &self.Ast_SymbolInfo
    }
    return nil
}

// 2701: decl @lune.@base.@Ast.AccessSymbolInfo.get_typeInfo
func (self *Ast_AccessSymbolInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.overrideTypeInfo
}

// 2705: decl @lune.@base.@Ast.AccessSymbolInfo.get_mutMode
func (self *Ast_AccessSymbolInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    switch _exp9046 := self.overrideMut.(type) {
    case *Ast_OverrideMut__None:
    case *Ast_OverrideMut__Prefix:
    prefixTypeInfo := _exp9046.Val1
        if _switch9037 := self.symbolInfo.FP.Get_mutMode(_env); _switch9037 == Ast_MutMode__AllMut || _switch9037 == Ast_MutMode__IMut || _switch9037 == Ast_MutMode__IMutRe {
            return self.symbolInfo.FP.Get_mutMode(_env)
        } else if _switch9037 == Ast_MutMode__Mut {
            if _switch9035 := prefixTypeInfo.FP.Get_mutMode(_env); _switch9035 == Ast_MutMode__AllMut {
                return Ast_MutMode__Mut
            } else if _switch9035 == Ast_MutMode__Mut || _switch9035 == Ast_MutMode__IMut || _switch9035 == Ast_MutMode__IMutRe {
                return prefixTypeInfo.FP.Get_mutMode(_env)
            }
        }
    case *Ast_OverrideMut__IMut:
        return Ast_MutMode__IMut
    }
    return self.symbolInfo.FP.Get_mutMode(_env)
}

// 2734: decl @lune.@base.@Ast.AccessSymbolInfo.get_mutable
func (self *Ast_AccessSymbolInfo) Get_mutable(_env *LnsEnv) bool {
    return Ast_isMutable(_env, self.FP.Get_mutMode(_env))
}

// 2738: decl @lune.@base.@Ast.AccessSymbolInfo.get_canBeLeft
func (self *Ast_AccessSymbolInfo) Get_canBeLeft(_env *LnsEnv) bool {
    if Lns_op_not(self.overrideCanBeLeft){
        return false
    }
    return self.symbolInfo.FP.Get_canBeLeft(_env)
}


// declaration Class -- AlternateTypeInfo
type Ast_AlternateTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    canSetFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 *LnsMap) bool
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_altIndex(_env *LnsEnv) LnsInt
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_txt(_env *LnsEnv) string
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
    updateParentInfo(_env *LnsEnv, arg1 *Ast_TypeInfo)
}
type Ast_AlternateTypeInfo struct {
    Ast_TypeInfo
    typeId *Ast_IdInfo
    txt string
    parentInfo *Ast_TypeInfo
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
func NewAst_AlternateTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 LnsInt, arg4 string, arg5 LnsInt, arg6 *Ast_TypeInfo, arg7 LnsAny, arg8 LnsAny) *Ast_AlternateTypeInfo {
    obj := &Ast_AlternateTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AlternateTypeInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Ast_AlternateTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_AlternateTypeInfo) Get_txt(_env *LnsEnv) string{ return self.txt }
func (self *Ast_AlternateTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return &self.nilableTypeInfo.Ast_TypeInfo }
func (self *Ast_AlternateTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Ast_AlternateTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.baseTypeInfo }
func (self *Ast_AlternateTypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList{ return self.interfaceList }
func (self *Ast_AlternateTypeInfo) Get_altIndex(_env *LnsEnv) LnsInt{ return self.altIndex }
// 2772: DeclConstr
func (self *Ast_AlternateTypeInfo) InitAst_AlternateTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,belongClassFlag bool,altIndex LnsInt,txt string,accessMode LnsInt,parentInfo *Ast_TypeInfo,baseTypeInfo LnsAny,interfaceList LnsAny) {
    self.InitAst_TypeInfo(_env, Ast_TypeInfo_createScope(_env, processInfo, nil, true, baseTypeInfo, interfaceList), processInfo)
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.txt = txt
    
    self.accessMode = accessMode
    
    self.parentInfo = parentInfo
    
    self.baseTypeInfo = Lns_unwrapDefault( baseTypeInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.interfaceList = Lns_unwrapDefault( interfaceList, NewLnsList([]LnsAny{})).(*LnsList)
    
    self.belongClassFlag = belongClassFlag
    
    self.altIndex = altIndex
    
    self.nilableTypeInfo = NewAst_NilableTypeInfo(_env, processInfo, &self.Ast_TypeInfo)
    
}

// 2795: decl @lune.@base.@Ast.AlternateTypeInfo.updateParentInfo
func (self *Ast_AlternateTypeInfo) updateParentInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo) {
    self.parentInfo = typeInfo
    
}

// 2800: decl @lune.@base.@Ast.AlternateTypeInfo.isModule
func (self *Ast_AlternateTypeInfo) IsModule(_env *LnsEnv) bool {
    return false
}

// 2805: decl @lune.@base.@Ast.AlternateTypeInfo.getParentId
func (self *Ast_AlternateTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.parentInfo.FP.Get_typeId(_env)
}

// 2809: decl @lune.@base.@Ast.AlternateTypeInfo.get_baseId
func (self *Ast_AlternateTypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return self.baseTypeInfo.FP.Get_typeId(_env)
}

// 2813: decl @lune.@base.@Ast.AlternateTypeInfo.get_parentInfo
func (self *Ast_AlternateTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.parentInfo
}

// 2818: decl @lune.@base.@Ast.AlternateTypeInfo.getTxt
func (self *Ast_AlternateTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 2824: decl @lune.@base.@Ast.AlternateTypeInfo.getTxtWithRaw
func (self *Ast_AlternateTypeInfo) GetTxtWithRaw(_env *LnsEnv, raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.txt
}



// 2840: decl @lune.@base.@Ast.AlternateTypeInfo.isInheritFrom
func (self *Ast_AlternateTypeInfo) IsInheritFrom(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    var workAlt2type LnsAny
    if alt2type != nil{
        alt2type_2231 := alt2type.(*LnsMap)
        var otherWork *Ast_TypeInfo
        otherWork = Ast_AlternateTypeInfo_getAssign(_env, other, alt2type_2231)
        if &self.Ast_TypeInfo == otherWork.FP.Get_srcTypeInfo(_env){
            return true
        }
        {
            _genType := alt2type_2231.Get(&self.Ast_TypeInfo)
            if !Lns_IsNil( _genType ) {
                genType := _genType.(*Ast_TypeInfo)
                return genType.FP.IsInheritFrom(_env, processInfo, otherWork, alt2type_2231)
            }
        }
        if Lns_op_not(Ast_CanEvalCtrlTypeInfo_isValidApply(_env, alt2type_2231)){
            workAlt2type = nil
            
        } else { 
            workAlt2type = alt2type_2231
            
        }
    } else {
        workAlt2type = nil
        
    }
    if &self.Ast_TypeInfo == other.FP.Get_srcTypeInfo(_env){
        return true
    }
    var check func(_env *LnsEnv) bool
    check = func(_env *LnsEnv) bool {
        if self.FP.HasBase(_env){
            if self.baseTypeInfo.FP.IsInheritFrom(_env, processInfo, other, workAlt2type){
                return true
            }
        }
        for _, _ifType := range( self.interfaceList.Items ) {
            ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if ifType.FP.IsInheritFrom(_env, processInfo, other, workAlt2type){
                return true
            }
        }
        return false
    }
    if check(_env){
        if workAlt2type != nil{
            workAlt2type_2250 := workAlt2type.(*LnsMap)
            workAlt2type_2250.Set(&self.Ast_TypeInfo,other)
        }
        return true
    }
    return false
}

// 2891: decl @lune.@base.@Ast.AlternateTypeInfo.canEvalWith
func (self *Ast_AlternateTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if &self.Ast_TypeInfo == other.FP.Get_srcTypeInfo(_env){
        return true, nil
    }
    if other.FP.Get_nilable(_env){
        return false, "is doesn't support nilable."
    }
    if Lns_op_not(self.belongClassFlag){
        {
            _altType := Ast_AlternateTypeInfoDownCastF(other.FP)
            if !Lns_IsNil( _altType ) {
                altType := _altType.(*Ast_AlternateTypeInfo)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( Lns_op_not(altType.belongClassFlag)) &&
                    _env.SetStackVal( altType.altIndex == self.altIndex) ).(bool)){
                    return true, nil
                }
            }
        }
    }
    return self.FP.canSetFrom(_env, processInfo, other, canEvalType, alt2type), nil
}

// 2917: decl @lune.@base.@Ast.AlternateTypeInfo.get_display_stirng_with
func (self *Ast_AlternateTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    if alt2type != nil{
        alt2type_2274 := alt2type.(*LnsMap)
        {
            _genType := alt2type_2274.Get(&self.Ast_TypeInfo)
            if !Lns_IsNil( _genType ) {
                genType := _genType.(*Ast_TypeInfo)
                return genType.FP.Get_display_stirng_with(_env, genType.FP.Get_rawTxt(_env), alt2type_2274)
            }
        }
    }
    return self.FP.GetTxtWithRaw(_env, raw, nil, nil, nil)
}

// 2927: decl @lune.@base.@Ast.AlternateTypeInfo.get_display_stirng
func (self *Ast_AlternateTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.txt, nil)
}

// 2933: decl @lune.@base.@Ast.AlternateTypeInfo.equals
func (self *Ast_AlternateTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if &self.Ast_TypeInfo == typeInfo{
        return true
    }
    if Lns_op_not(self.belongClassFlag){
        {
            _altType := Ast_AlternateTypeInfoDownCastF(typeInfo.FP)
            if !Lns_IsNil( _altType ) {
                altType := _altType.(*Ast_AlternateTypeInfo)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( Lns_op_not(altType.belongClassFlag)) &&
                    _env.SetStackVal( altType.altIndex == self.altIndex) ).(bool)){
                    return true
                }
            }
        }
    }
    if alt2type != nil{
        alt2type_2297 := alt2type.(*LnsMap)
        return self.FP.canSetFrom(_env, processInfo, typeInfo, nil, alt2type_2297)
    }
    return false
}

// 2961: decl @lune.@base.@Ast.AlternateTypeInfo.hasRouteNamespaceFrom
func (self *Ast_AlternateTypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, other *Ast_TypeInfo) bool {
    return true
}

// 2966: decl @lune.@base.@Ast.AlternateTypeInfo.get_rawTxt
func (self *Ast_AlternateTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return self.txt
}

// 2970: decl @lune.@base.@Ast.AlternateTypeInfo.get_kind
func (self *Ast_AlternateTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Alternate
}

// 2974: decl @lune.@base.@Ast.AlternateTypeInfo.get_nilable
func (self *Ast_AlternateTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return false
}

// 2978: decl @lune.@base.@Ast.AlternateTypeInfo.get_mutMode
func (self *Ast_AlternateTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return Ast_MutMode__Mut
}

// 2990: decl @lune.@base.@Ast.AlternateTypeInfo.serialize
func (self *Ast_AlternateTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    var parentId *Ast_IdInfo
    parentId = self.FP.GetParentId(_env)
    stream.Write(_env, _env.LuaVM.String_format("{ skind = %d, parentId = %d, typeId = %d, txt = %q, ", []LnsAny{Ast_SerializeKind__Alternate, parentId.Id, self.typeId.Id, self.txt}) + _env.LuaVM.String_format("accessMode = %d, baseId = %s, ", []LnsAny{self.accessMode, serializeInfo.FP.SerializeId(_env, self.FP.Get_baseId(_env))}) + _env.LuaVM.String_format("belongClassFlag = %s, altIndex = %d, ", []LnsAny{self.belongClassFlag, self.altIndex}))
    stream.Write(_env, self.FP.SerializeTypeInfoList(_env, serializeInfo, "ifList = {", self.interfaceList, nil))
    stream.Write(_env, "}\n")
}

// 3005: decl @lune.@base.@Ast.AlternateTypeInfo.applyGeneric
func (self *Ast_AlternateTypeInfo) ApplyGeneric(_env *LnsEnv, alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    return Ast_AlternateTypeInfo_getAssign(_env, &self.Ast_TypeInfo, alt2typeMap)
}

// 4578: decl @lune.@base.@Ast.AlternateTypeInfo.canSetFrom
func (self *Ast_AlternateTypeInfo) canSetFrom(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsAny,alt2type *LnsMap) bool {
    var otherWork *Ast_TypeInfo
    otherWork = Ast_AlternateTypeInfo_getAssign(_env, other, alt2type)
    if &self.Ast_TypeInfo == otherWork{
        return true
    }
    {
        _genType := alt2type.Get(&self.Ast_TypeInfo)
        if !Lns_IsNil( _genType ) {
            genType := _genType.(*Ast_TypeInfo)
            if canEvalType != nil{
                canEvalType_3627 := canEvalType.(LnsInt)
                return Lns_car(genType.FP.CanEvalWith(_env, processInfo, otherWork, canEvalType_3627, alt2type)).(bool)
            }
            return genType.FP.Equals(_env, processInfo, otherWork, alt2type, nil)
        }
    }
    var workAlt2type *LnsMap
    if Lns_op_not(Ast_CanEvalCtrlTypeInfo_isValidApply(_env, alt2type)){
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(Ast_isClass(_env, otherWork))) &&
            _env.SetStackVal( otherWork.FP.Get_kind(_env) != Ast_TypeInfoKind__IF) ).(bool)){
            return false
        }
        workAlt2type = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false)
        
    } else { 
        workAlt2type = alt2type
        
    }
    if self.FP.HasBase(_env){
        if Lns_op_not(other.FP.IsInheritFrom(_env, processInfo, self.baseTypeInfo, workAlt2type)){
            return false
        }
    }
    for _, _ifType := range( self.interfaceList.Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(other.FP.IsInheritFrom(_env, processInfo, ifType, workAlt2type)){
            return false
        }
    }
    workAlt2type.Set(&self.Ast_TypeInfo,otherWork)
    return true
}

// 4727: decl @lune.@base.@Ast.AlternateTypeInfo.getAssign
func Ast_AlternateTypeInfo_getAssign(_env *LnsEnv, typeInfo *Ast_TypeInfo,alt2type *LnsMap) *Ast_TypeInfo {
    if typeInfo.FP.Get_kind(_env) != Ast_TypeInfoKind__Alternate{
        return typeInfo
    }
    var otherWork *Ast_TypeInfo
    otherWork = typeInfo
    for  {
        {
            __exp := alt2type.Get(otherWork)
            if !Lns_IsNil( __exp ) {
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
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_boxingType(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_BoxTypeInfo struct {
    Ast_TypeInfo
    boxingType *Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_BoxTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsInt, arg3 *Ast_TypeInfo) *Ast_BoxTypeInfo {
    obj := &Ast_BoxTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_BoxTypeInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Ast_BoxTypeInfo) Get_boxingType(_env *LnsEnv) *Ast_TypeInfo{ return self.boxingType }
func (self *Ast_BoxTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_BoxTypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList{ return self.itemTypeInfoList }
func (self *Ast_BoxTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Ast_BoxTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return &self.nilableTypeInfo.Ast_TypeInfo }
// advertise -- 3016
func (self *Ast_BoxTypeInfo) AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo) {
self.boxingType. FP.AddChildren( _env, arg1)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.boxingType. FP.GetFullName( _env, arg1,arg2,arg3)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.boxingType. FP.GetModule( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) GetOverridingType(_env *LnsEnv) LnsAny {
    return self.boxingType. FP.GetOverridingType( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.boxingType. FP.GetParentFullName( _env, arg1,arg2,arg3)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.boxingType. FP.GetParentId( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.boxingType. FP.getProcessInfo( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_abstractFlag(_env *LnsEnv) bool {
    return self.boxingType. FP.Get_abstractFlag( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_argTypeInfoList(_env *LnsEnv) *LnsList {
    return self.boxingType. FP.Get_argTypeInfoList( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_asyncMode(_env *LnsEnv) LnsInt {
    return self.boxingType. FP.Get_asyncMode( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_autoFlag(_env *LnsEnv) bool {
    return self.boxingType. FP.Get_autoFlag( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return self.boxingType. FP.Get_baseId( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.boxingType. FP.Get_baseTypeInfo( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_children(_env *LnsEnv) *LnsList {
    return self.boxingType. FP.Get_children( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_externalFlag(_env *LnsEnv) bool {
    return self.boxingType. FP.Get_externalFlag( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.boxingType. FP.Get_genSrcTypeInfo( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList {
    return self.boxingType. FP.Get_interfaceList( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return self.boxingType. FP.Get_mutMode( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.boxingType. FP.Get_parentInfo( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.boxingType. FP.Get_processInfo( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return self.boxingType. FP.Get_rawTxt( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList {
    return self.boxingType. FP.Get_retTypeInfoList( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_staticFlag(_env *LnsEnv) bool {
    return self.boxingType. FP.Get_staticFlag( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) Get_typeData(_env *LnsEnv) *Ast_TypeData {
    return self.boxingType. FP.Get_typeData( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) HasBase(_env *LnsEnv) bool {
    return self.boxingType. FP.HasBase( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool {
    return self.boxingType. FP.HasRouteNamespaceFrom( _env, arg1)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.boxingType. FP.IsInheritFrom( _env, arg1,arg2,arg3)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) IsModule(_env *LnsEnv) bool {
    return self.boxingType. FP.IsModule( _env)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo,arg2 string,arg3 *LnsList,arg4 LnsAny) string {
    return self.boxingType. FP.SerializeTypeInfoList( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 3016
func (self *Ast_BoxTypeInfo) SwitchScope(_env *LnsEnv, arg1 *Ast_Scope) {
self.boxingType. FP.SwitchScope( _env, arg1)
}
// 3023: DeclConstr
func (self *Ast_BoxTypeInfo) InitAst_BoxTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,accessMode LnsInt,boxingType *Ast_TypeInfo) {
    self.InitAst_TypeInfo(_env, Ast_boxRootScope, processInfo)
    self.boxingType = boxingType
    
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.itemTypeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(boxingType)})
    
    self.accessMode = accessMode
    
    self.nilableTypeInfo = NewAst_NilableTypeInfo(_env, processInfo, &self.Ast_TypeInfo)
    
}

// 3036: decl @lune.@base.@Ast.BoxTypeInfo.get_scope
func (self *Ast_BoxTypeInfo) Get_scope(_env *LnsEnv) LnsAny {
    return self.Ast_TypeInfo.Get_scope(_env)
}

// 3040: decl @lune.@base.@Ast.BoxTypeInfo.get_kind
func (self *Ast_BoxTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Box
}

// 3043: decl @lune.@base.@Ast.BoxTypeInfo.get_aliasSrc
func (self *Ast_BoxTypeInfo) Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 3046: decl @lune.@base.@Ast.BoxTypeInfo.get_srcTypeInfo
func (self *Ast_BoxTypeInfo) Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 3049: decl @lune.@base.@Ast.BoxTypeInfo.get_nonnilableType
func (self *Ast_BoxTypeInfo) Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 3052: decl @lune.@base.@Ast.BoxTypeInfo.get_nilable
func (self *Ast_BoxTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return false
}

// 3055: decl @lune.@base.@Ast.BoxTypeInfo.get_extedType
func (self *Ast_BoxTypeInfo) Get_extedType(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 3061: decl @lune.@base.@Ast.BoxTypeInfo.getTxt
func (self *Ast_BoxTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 3066: decl @lune.@base.@Ast.BoxTypeInfo.getTxtWithRaw
func (self *Ast_BoxTypeInfo) GetTxtWithRaw(_env *LnsEnv, raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return "Nilable<" + self.boxingType.FP.GetTxtWithRaw(_env, raw, typeNameCtrl, importInfo, localFlag) + ">"
}

// 3074: decl @lune.@base.@Ast.BoxTypeInfo.get_display_stirng
func (self *Ast_BoxTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
}

// 3078: decl @lune.@base.@Ast.BoxTypeInfo.get_display_stirng_with
func (self *Ast_BoxTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    return _env.LuaVM.String_format("Nilable<%s>", []LnsAny{self.boxingType.FP.Get_display_stirng_with(_env, raw, alt2type)})
}

// 3083: decl @lune.@base.@Ast.BoxTypeInfo.serialize
func (self *Ast_BoxTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    stream.Write(_env, _env.LuaVM.String_format("{ skind = %d, typeId = %d, accessMode = %d, boxingType = %d }\n", []LnsAny{Ast_SerializeKind__Box, self.typeId.Id, self.accessMode, self.boxingType.FP.Get_typeId(_env).Id}))
}

// 3090: decl @lune.@base.@Ast.BoxTypeInfo.equals
func (self *Ast_BoxTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    {
        _boxType := Ast_BoxTypeInfoDownCastF(typeInfo.FP)
        if !Lns_IsNil( _boxType ) {
            boxType := _boxType.(*Ast_BoxTypeInfo)
            return self.boxingType.FP.Equals(_env, processInfo, boxType.boxingType, alt2type, checkModifer)
        }
    }
    return false
}

// 3101: decl @lune.@base.@Ast.BoxTypeInfo.createAlt2typeMap
func (self *Ast_BoxTypeInfo) CreateAlt2typeMap(_env *LnsEnv, detectFlag bool) *LnsMap {
    var _map *LnsMap
    _map = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, detectFlag)
    if self.boxingType != &Ast_boxRootAltType.Ast_TypeInfo{
        _map.Set(&Ast_boxRootAltType.Ast_TypeInfo,self.boxingType)
    }
    return _map
}

// 4795: decl @lune.@base.@Ast.BoxTypeInfo.applyGeneric
func (self *Ast_BoxTypeInfo) ApplyGeneric(_env *LnsEnv, alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.boxingType.FP.ApplyGeneric(_env, alt2typeMap, moduleTypeInfo)
    if typeInfo == self.boxingType{
        return &self.Ast_TypeInfo
    }
    return nil
}

// 6088: decl @lune.@base.@Ast.BoxTypeInfo.canEvalWith
func (self *Ast_BoxTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if &self.Ast_TypeInfo == other{
        return true, nil
    }
    if _switch24476 := canEvalType; _switch24476 == Ast_CanEvalType__SetOp || _switch24476 == Ast_CanEvalType__SetOpIMut || _switch24476 == Ast_CanEvalType__SetEq {
    } else {
        return false, nil
    }
    if other.FP.Get_nilable(_env){
        Ast_CanEvalCtrlTypeInfo_updateNeedAutoBoxing(_env, alt2type)
        return true, nil
    }
    {
        _otherBoxType := Ast_BoxTypeInfoDownCastF(other.FP)
        if !Lns_IsNil( _otherBoxType ) {
            otherBoxType := _otherBoxType.(*Ast_BoxTypeInfo)
            return self.boxingType.FP.CanEvalWith(_env, processInfo, otherBoxType.boxingType, canEvalType, alt2type)
        }
    }
    if Lns_car(self.boxingType.FP.CanEvalWith(_env, processInfo, other, canEvalType, alt2type)).(bool){
        Ast_CanEvalCtrlTypeInfo_updateNeedAutoBoxing(_env, alt2type)
        return true, nil
    }
    return false, nil
}


// declaration Class -- GenericTypeInfo
type Ast_GenericTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_GenericTypeInfo struct {
    Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_GenericTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 *Ast_TypeInfo) *Ast_GenericTypeInfo {
    obj := &Ast_GenericTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_GenericTypeInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *Ast_GenericTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_GenericTypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList{ return self.itemTypeInfoList }
func (self *Ast_GenericTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return &self.nilableTypeInfo.Ast_TypeInfo }
func (self *Ast_GenericTypeInfo) Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.genSrcTypeInfo }
// advertise -- 3119
func (self *Ast_GenericTypeInfo) AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo) {
self.genSrcTypeInfo. FP.AddChildren( _env, arg1)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.genSrcTypeInfo. FP.GetFullName( _env, arg1,arg2,arg3)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) GetOverridingType(_env *LnsEnv) LnsAny {
    return self.genSrcTypeInfo. FP.GetOverridingType( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.genSrcTypeInfo. FP.GetParentFullName( _env, arg1,arg2,arg3)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.genSrcTypeInfo. FP.GetParentId( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.genSrcTypeInfo. FP.getProcessInfo( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) GetTxt(_env *LnsEnv, arg1 LnsAny,arg2 LnsAny,arg3 LnsAny) string {
    return self.genSrcTypeInfo. FP.GetTxt( _env, arg1,arg2,arg3)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) GetTxtWithRaw(_env *LnsEnv, arg1 string,arg2 LnsAny,arg3 LnsAny,arg4 LnsAny) string {
    return self.genSrcTypeInfo. FP.GetTxtWithRaw( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_abstractFlag(_env *LnsEnv) bool {
    return self.genSrcTypeInfo. FP.Get_abstractFlag( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return self.genSrcTypeInfo. FP.Get_accessMode( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_argTypeInfoList(_env *LnsEnv) *LnsList {
    return self.genSrcTypeInfo. FP.Get_argTypeInfoList( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_asyncMode(_env *LnsEnv) LnsInt {
    return self.genSrcTypeInfo. FP.Get_asyncMode( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_autoFlag(_env *LnsEnv) bool {
    return self.genSrcTypeInfo. FP.Get_autoFlag( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return self.genSrcTypeInfo. FP.Get_baseId( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.genSrcTypeInfo. FP.Get_baseTypeInfo( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_children(_env *LnsEnv) *LnsList {
    return self.genSrcTypeInfo. FP.Get_children( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.genSrcTypeInfo. FP.Get_display_stirng( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_externalFlag(_env *LnsEnv) bool {
    return self.genSrcTypeInfo. FP.Get_externalFlag( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList {
    return self.genSrcTypeInfo. FP.Get_interfaceList( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return self.genSrcTypeInfo. FP.Get_kind( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return self.genSrcTypeInfo. FP.Get_mutMode( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return self.genSrcTypeInfo. FP.Get_nilable( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo {
    return self.genSrcTypeInfo. FP.Get_nonnilableType( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.genSrcTypeInfo. FP.Get_parentInfo( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.genSrcTypeInfo. FP.Get_processInfo( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return self.genSrcTypeInfo. FP.Get_rawTxt( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList {
    return self.genSrcTypeInfo. FP.Get_retTypeInfoList( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_scope(_env *LnsEnv) LnsAny {
    return self.genSrcTypeInfo. FP.Get_scope( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_staticFlag(_env *LnsEnv) bool {
    return self.genSrcTypeInfo. FP.Get_staticFlag( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) Get_typeData(_env *LnsEnv) *Ast_TypeData {
    return self.genSrcTypeInfo. FP.Get_typeData( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) HasBase(_env *LnsEnv) bool {
    return self.genSrcTypeInfo. FP.HasBase( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool {
    return self.genSrcTypeInfo. FP.HasRouteNamespaceFrom( _env, arg1)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) IsModule(_env *LnsEnv) bool {
    return self.genSrcTypeInfo. FP.IsModule( _env)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo,arg2 string,arg3 *LnsList,arg4 LnsAny) string {
    return self.genSrcTypeInfo. FP.SerializeTypeInfoList( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 3119
func (self *Ast_GenericTypeInfo) SwitchScope(_env *LnsEnv, arg1 *Ast_Scope) {
self.genSrcTypeInfo. FP.SwitchScope( _env, arg1)
}
// 3128: decl @lune.@base.@Ast.GenericTypeInfo.get_display_stirng_with
func (self *Ast_GenericTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    return self.genSrcTypeInfo.FP.Get_display_stirng_with(_env, raw, self.alt2typeMap)
}

// 3134: DeclConstr
func (self *Ast_GenericTypeInfo) InitAst_GenericTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,genSrcTypeInfo *Ast_TypeInfo,itemTypeInfoList *LnsList,moduleTypeInfo *Ast_TypeInfo) {
    self.InitAst_TypeInfo(_env, Ast_TypeInfo_createScope(_env, processInfo, (Lns_unwrap( genSrcTypeInfo.FP.Get_scope(_env)).(*Ast_Scope)).FP.Get_parent(_env), true, genSrcTypeInfo, nil), processInfo)
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.moduleTypeInfo = moduleTypeInfo
    
    self.itemTypeInfoList = itemTypeInfoList
    
    self.genSrcTypeInfo = genSrcTypeInfo
    
    if genSrcTypeInfo.FP.Get_itemTypeInfoList(_env).Len() != itemTypeInfoList.Len(){
        Util_err(_env, _env.LuaVM.String_format("unmatch generic type number -- %d, %d", []LnsAny{genSrcTypeInfo.FP.Get_itemTypeInfoList(_env).Len(), itemTypeInfoList.Len()}))
    }
    var alt2typeMap *LnsMap
    alt2typeMap = NewLnsMap( map[LnsAny]LnsAny{})
    var workAlt2typeMap *LnsMap
    workAlt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false)
    var hasAlter bool
    hasAlter = false
    for _index, _altTypeInfo := range( genSrcTypeInfo.FP.Get_itemTypeInfoList(_env).Items ) {
        index := _index + 1
        altTypeInfo := _altTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var itemType *Ast_TypeInfo
        itemType = itemTypeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        alt2typeMap.Set(altTypeInfo,itemType)
        if itemType.FP.ApplyGeneric(_env, workAlt2typeMap, moduleTypeInfo) != itemType{
            hasAlter = true
            
        }
    }
    self.hasAlter = hasAlter
    
    self.alt2typeMap = alt2typeMap
    
    self.nilableTypeInfo = NewAst_NilableTypeInfo(_env, processInfo, &self.Ast_TypeInfo)
    
}

// 3170: decl @lune.@base.@Ast.GenericTypeInfo.getModule
func (self *Ast_GenericTypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.moduleTypeInfo
}

// 3175: decl @lune.@base.@Ast.GenericTypeInfo.isInheritFrom
func (self *Ast_GenericTypeInfo) IsInheritFrom(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    var otherSrc *Ast_TypeInfo
    otherSrc = other.FP.Get_genSrcTypeInfo(_env)
    if Lns_op_not(self.genSrcTypeInfo.FP.IsInheritFrom(_env, processInfo, otherSrc, alt2type)){
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
            workAlt2type = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false)
            
        } else {
            workAlt2type = _workAlt2type.(*LnsMap)
        }
    }
    for _, _altType := range( otherSrc.FP.Get_itemTypeInfoList(_env).Items ) {
        altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var genType *Ast_TypeInfo
        
        {
            _genType := self.alt2typeMap.Get(altType)
            if _genType == nil{
                return false
            } else {
                genType = _genType.(*Ast_TypeInfo)
            }
        }
        var otherGenType *Ast_TypeInfo
        otherGenType = Lns_unwrap( genOther.alt2typeMap.Get(altType)).(*Ast_TypeInfo)
        if Lns_op_not(Lns_car(otherGenType.FP.CanEvalWith(_env, processInfo, genType, Ast_CanEvalType__SetEq, workAlt2type)).(bool)){
            return false
        }
    }
    return true
}

// 3215: decl @lune.@base.@Ast.GenericTypeInfo.get_aliasSrc
func (self *Ast_GenericTypeInfo) Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 3219: decl @lune.@base.@Ast.GenericTypeInfo.get_srcTypeInfo
func (self *Ast_GenericTypeInfo) Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 3222: decl @lune.@base.@Ast.GenericTypeInfo.get_extedType
func (self *Ast_GenericTypeInfo) Get_extedType(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 3227: decl @lune.@base.@Ast.GenericTypeInfo.canEvalWith
func (self *Ast_GenericTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if other.FP.Get_nilable(_env){
        return false, "GenericTypeInfo doesn't support nilable."
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Ast_TypeInfo_isMut(_env, &self.Ast_TypeInfo)) &&
        _env.SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(_env, other))) ).(bool)){
        return false, nil
    }
    var otherSrc *Ast_TypeInfo
    otherSrc = other.FP.Get_srcTypeInfo(_env)
    if &self.Ast_TypeInfo == otherSrc{
        return true, nil
    }
    var work *Ast_TypeInfo
    work = otherSrc
    for  {
        if work == Ast_headTypeInfo{
            return false, nil
        }
        for _altType, _genType := range( work.FP.CreateAlt2typeMap(_env, false).Items ) {
            altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            alt2type.Set(altType,genType)
        }
        if self.genSrcTypeInfo.FP.Equals(_env, processInfo, work.FP.Get_genSrcTypeInfo(_env), alt2type, nil){
            break
        }
        for _, _ifType := range( work.FP.Get_interfaceList(_env).Items ) {
            ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_car(self.FP.CanEvalWith(_env, processInfo, ifType, canEvalType, alt2type)).(bool){
                return true, nil
            }
        }
        work = work.FP.Get_baseTypeInfo(_env)
        
    }
    {
        _otherGen := Ast_GenericTypeInfoDownCastF(work.FP)
        if !Lns_IsNil( _otherGen ) {
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
                otherType = Ast_AlternateTypeInfo_getAssign(_env, Lns_unwrap( otherGen.alt2typeMap.Get(key)).(*Ast_TypeInfo), alt2type)
                var ret bool
                var mess LnsAny
                ret,mess = val.FP.CanEvalWith(_env, processInfo, otherType, evalType, alt2type)
                if Lns_op_not(ret){
                    return false, mess
                }
            }
        }
    }
    return true, nil
}

// 3288: decl @lune.@base.@Ast.GenericTypeInfo.equals
func (self *Ast_GenericTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if &self.Ast_TypeInfo == other{
        return true
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.Get_kind(_env) != self.FP.Get_kind(_env)) ||
        _env.SetStackVal( self.itemTypeInfoList.Len() != other.FP.Get_itemTypeInfoList(_env).Len()) ).(bool){
        return false
    }
    if Lns_op_not((Ast_GenericTypeInfoDownCastF(other.FP))){
        return false
    }
    if Lns_op_not(self.genSrcTypeInfo.FP.Equals(_env, processInfo, other.FP.Get_genSrcTypeInfo(_env), alt2type, checkModifer)){
        return false
    }
    for _index, _otherItem := range( other.FP.Get_itemTypeInfoList(_env).Items ) {
        index := _index + 1
        otherItem := _otherItem.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var typeInfo *Ast_TypeInfo
        typeInfo = self.itemTypeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(typeInfo.FP.Equals(_env, processInfo, otherItem, alt2type, checkModifer)){
            return false
        }
    }
    return true
}

// 3321: decl @lune.@base.@Ast.GenericTypeInfo.serialize
func (self *Ast_GenericTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    stream.Write(_env, _env.LuaVM.String_format("{ skind = %d, typeId = %d, genSrcTypeId = %s, genTypeList = {", []LnsAny{Ast_SerializeKind__Generic, self.typeId.Id, serializeInfo.FP.SerializeId(_env, self.genSrcTypeInfo.FP.Get_typeId(_env))}))
    var count LnsInt
    count = 0
    for _, _genType := range( self.alt2typeMap.Items ) {
        genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if count > 0{
            stream.Write(_env, ",")
        }
        stream.Write(_env, serializeInfo.FP.SerializeId(_env, genType.FP.Get_typeId(_env)))
    }
    stream.Write(_env, "} }\n")
}

// 3336: decl @lune.@base.@Ast.GenericTypeInfo.createAlt2typeMap
func (self *Ast_GenericTypeInfo) CreateAlt2typeMap(_env *LnsEnv, detectFlag bool) *LnsMap {
    var _map *LnsMap
    _map = self.genSrcTypeInfo.FP.CreateAlt2typeMap(_env, detectFlag)
    for _genType, _typeInfo := range( self.alt2typeMap.Items ) {
        genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        _map.Set(genType,typeInfo)
    }
    return _map
}

// 5609: decl @lune.@base.@Ast.GenericTypeInfo.applyGeneric
func (self *Ast_GenericTypeInfo) ApplyGeneric(_env *LnsEnv, alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    if self.genSrcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Class{
        var itemTypeInfoList LnsAny
        var newFlag bool
        itemTypeInfoList,newFlag = Ast_applyGenericList_3567_(_env, self.FP.Get_itemTypeInfoList(_env), alt2typeMap, moduleTypeInfo)
        if itemTypeInfoList != nil{
            itemTypeInfoList_4312 := itemTypeInfoList.(*LnsList)
            if newFlag{
                return &moduleTypeInfo.FP.getProcessInfo(_env).FP.CreateGeneric(_env, self.genSrcTypeInfo, itemTypeInfoList_4312, moduleTypeInfo).Ast_TypeInfo
            }
        }
    }
    var genSrcTypeInfo LnsAny
    genSrcTypeInfo = self.genSrcTypeInfo.FP.ApplyGeneric(_env, alt2typeMap, moduleTypeInfo)
    if genSrcTypeInfo == self.genSrcTypeInfo{
        return &self.Ast_TypeInfo
    }
    if Lns_op_not(self.hasAlter){
        return &self.Ast_TypeInfo
    }
    Util_errorLog(_env, _env.LuaVM.String_format("no support nest generic -- %s", []LnsAny{self.FP.GetTxt(_env, nil, nil, nil)}))
    return nil
}


// declaration Class -- ModifierTypeInfo
type Ast_ModifierTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_ModifierTypeInfo struct {
    Ast_TypeInfo
    srcTypeInfo *Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_ModifierTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt) *Ast_ModifierTypeInfo {
    obj := &Ast_ModifierTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_ModifierTypeInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Ast_ModifierTypeInfo) Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.srcTypeInfo }
func (self *Ast_ModifierTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_ModifierTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt{ return self.mutMode }
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo) {
self.srcTypeInfo. FP.AddChildren( _env, arg1)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap {
    return self.srcTypeInfo. FP.CreateAlt2typeMap( _env, arg1)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.srcTypeInfo. FP.GetFullName( _env, arg1,arg2,arg3)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.srcTypeInfo. FP.GetModule( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) GetOverridingType(_env *LnsEnv) LnsAny {
    return self.srcTypeInfo. FP.GetOverridingType( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.srcTypeInfo. FP.GetParentFullName( _env, arg1,arg2,arg3)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.srcTypeInfo. FP.GetParentId( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.srcTypeInfo. FP.getProcessInfo( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_abstractFlag(_env *LnsEnv) bool {
    return self.srcTypeInfo. FP.Get_abstractFlag( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return self.srcTypeInfo. FP.Get_accessMode( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo {
    return self.srcTypeInfo. FP.Get_aliasSrc( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_argTypeInfoList(_env *LnsEnv) *LnsList {
    return self.srcTypeInfo. FP.Get_argTypeInfoList( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_asyncMode(_env *LnsEnv) LnsInt {
    return self.srcTypeInfo. FP.Get_asyncMode( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_autoFlag(_env *LnsEnv) bool {
    return self.srcTypeInfo. FP.Get_autoFlag( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return self.srcTypeInfo. FP.Get_baseId( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.srcTypeInfo. FP.Get_baseTypeInfo( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_children(_env *LnsEnv) *LnsList {
    return self.srcTypeInfo. FP.Get_children( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_externalFlag(_env *LnsEnv) bool {
    return self.srcTypeInfo. FP.Get_externalFlag( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.srcTypeInfo. FP.Get_genSrcTypeInfo( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList {
    return self.srcTypeInfo. FP.Get_interfaceList( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList {
    return self.srcTypeInfo. FP.Get_itemTypeInfoList( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return self.srcTypeInfo. FP.Get_kind( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return self.srcTypeInfo. FP.Get_nilable( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.srcTypeInfo. FP.Get_parentInfo( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.srcTypeInfo. FP.Get_processInfo( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return self.srcTypeInfo. FP.Get_rawTxt( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList {
    return self.srcTypeInfo. FP.Get_retTypeInfoList( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_scope(_env *LnsEnv) LnsAny {
    return self.srcTypeInfo. FP.Get_scope( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_staticFlag(_env *LnsEnv) bool {
    return self.srcTypeInfo. FP.Get_staticFlag( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) Get_typeData(_env *LnsEnv) *Ast_TypeData {
    return self.srcTypeInfo. FP.Get_typeData( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) HasBase(_env *LnsEnv) bool {
    return self.srcTypeInfo. FP.HasBase( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool {
    return self.srcTypeInfo. FP.HasRouteNamespaceFrom( _env, arg1)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.srcTypeInfo. FP.IsInheritFrom( _env, arg1,arg2,arg3)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) IsModule(_env *LnsEnv) bool {
    return self.srcTypeInfo. FP.IsModule( _env)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo,arg2 string,arg3 *LnsList,arg4 LnsAny) string {
    return self.srcTypeInfo. FP.SerializeTypeInfoList( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 3355
func (self *Ast_ModifierTypeInfo) SwitchScope(_env *LnsEnv, arg1 *Ast_Scope) {
self.srcTypeInfo. FP.SwitchScope( _env, arg1)
}
// 3360: DeclConstr
func (self *Ast_ModifierTypeInfo) InitAst_ModifierTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,srcTypeInfo *Ast_TypeInfo,mutMode LnsInt) {
    self.InitAst_TypeInfo(_env, nil, processInfo)
    self.srcTypeInfo = srcTypeInfo
    
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.mutMode = mutMode
    
}

// 3367: decl @lune.@base.@Ast.ModifierTypeInfo.get_extedType
func (self *Ast_ModifierTypeInfo) Get_extedType(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 3372: decl @lune.@base.@Ast.ModifierTypeInfo.getTxt
func (self *Ast_ModifierTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 3377: decl @lune.@base.@Ast.ModifierTypeInfo.getTxtWithRaw
func (self *Ast_ModifierTypeInfo) GetTxtWithRaw(_env *LnsEnv, raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    var txt string
    txt = self.srcTypeInfo.FP.GetTxtWithRaw(_env, raw, typeNameCtrl, importInfo, localFlag)
    if _switch11703 := self.mutMode; _switch11703 == Ast_MutMode__IMut || _switch11703 == Ast_MutMode__IMutRe {
        txt = "&" + txt
        
    } else if _switch11703 == Ast_MutMode__AllMut {
        txt = "allmut " + txt
        
    }
    return txt
}

// 3394: decl @lune.@base.@Ast.ModifierTypeInfo.get_display_stirng_with
func (self *Ast_ModifierTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    var txt string
    txt = self.srcTypeInfo.FP.Get_display_stirng_with(_env, raw, alt2type)
    if Ast_isMutable(_env, self.mutMode){
        txt = "mut " + txt
        
    }
    return txt
}

// 3402: decl @lune.@base.@Ast.ModifierTypeInfo.get_display_stirng
func (self *Ast_ModifierTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
}

// 3406: decl @lune.@base.@Ast.ModifierTypeInfo.serialize
func (self *Ast_ModifierTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    stream.Write(_env, _env.LuaVM.String_format("{ skind = %d, typeId = %d, srcTypeId = %s, mutMode = %d }\n", []LnsAny{Ast_SerializeKind__Modifier, self.typeId.Id, serializeInfo.FP.SerializeId(_env, self.srcTypeInfo.FP.Get_typeId(_env)), self.mutMode}))
}

// 3414: decl @lune.@base.@Ast.ModifierTypeInfo.canEvalWith
func (self *Ast_ModifierTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    var evalType LnsInt
    if self.srcTypeInfo.FP.Get_itemTypeInfoList(_env).Len() >= 1{
        if _switch11872 := canEvalType; _switch11872 == Ast_CanEvalType__SetEq || _switch11872 == Ast_CanEvalType__SetOp {
            evalType = Ast_CanEvalType__SetOpIMut
            
        } else {
            evalType = canEvalType
            
        }
    } else { 
        evalType = canEvalType
        
    }
    var result bool
    var mess LnsAny
    result,mess = Ast_TypeInfo_canEvalWithBase(_env, processInfo, self.srcTypeInfo, Ast_TypeInfo_isMut(_env, &self.Ast_TypeInfo), other.FP.Get_srcTypeInfo(_env), evalType, alt2type)
    return result, mess
}

// 3445: decl @lune.@base.@Ast.ModifierTypeInfo.equals
func (self *Ast_ModifierTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if Lns_isCondTrue( checkModifer){
        if Ast_TypeInfo_isMut(_env, &self.Ast_TypeInfo) != Ast_TypeInfo_isMut(_env, typeInfo){
            return false
        }
    }
    return self.srcTypeInfo.FP.Equals(_env, processInfo, typeInfo.FP.Get_srcTypeInfo(_env), alt2type, checkModifer)
}

// 5004: decl @lune.@base.@Ast.ModifierTypeInfo.get_nonnilableType
func (self *Ast_ModifierTypeInfo) Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo {
    var orgType *Ast_TypeInfo
    orgType = self.srcTypeInfo.FP.Get_nonnilableType(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Ast_TypeInfo_isMut(_env, &self.Ast_TypeInfo)) ||
        _env.SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(_env, orgType))) ).(bool){
        return orgType
    }
    return self.FP.getProcessInfo(_env).FP.CreateModifier(_env, orgType, Ast_MutMode__IMut)
}

// 5012: decl @lune.@base.@Ast.ModifierTypeInfo.get_nilableTypeInfo
func (self *Ast_ModifierTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    var orgType *Ast_TypeInfo
    orgType = self.srcTypeInfo.FP.Get_nilableTypeInfo(_env)
    if Lns_op_not(Ast_TypeInfo_isMut(_env, orgType)){
        return orgType
    }
    return self.FP.getProcessInfo(_env).FP.CreateModifier(_env, orgType, Ast_MutMode__IMut)
}

// 6982: decl @lune.@base.@Ast.ModifierTypeInfo.applyGeneric
func (self *Ast_ModifierTypeInfo) ApplyGeneric(_env *LnsEnv, alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.srcTypeInfo.FP.ApplyGeneric(_env, alt2typeMap, moduleTypeInfo)
    if typeInfo == self.srcTypeInfo{
        return &self.Ast_TypeInfo
    }
    if typeInfo != nil{
        typeInfo_5156 := typeInfo.(*Ast_TypeInfo)
        return moduleTypeInfo.FP.getProcessInfo(_env).FP.CreateModifier(_env, typeInfo_5156, Ast_MutMode__IMut)
    }
    return nil
}


// declaration Class -- ModuleTypeInfo
type Ast_ModuleTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_mutable(_env *LnsEnv) bool
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_ModuleTypeInfo struct {
    Ast_TypeInfo
    externalFlag bool
    parentInfo *Ast_TypeInfo
    typeId *Ast_IdInfo
    rawTxt string
    mutable bool
    fullName string
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
func NewAst_ModuleTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_Scope, arg3 bool, arg4 string, arg5 LnsAny, arg6 bool) *Ast_ModuleTypeInfo {
    obj := &Ast_ModuleTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_ModuleTypeInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Ast_ModuleTypeInfo) Get_externalFlag(_env *LnsEnv) bool{ return self.externalFlag }
func (self *Ast_ModuleTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_ModuleTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_ModuleTypeInfo) Get_rawTxt(_env *LnsEnv) string{ return self.rawTxt }
func (self *Ast_ModuleTypeInfo) Get_mutable(_env *LnsEnv) bool{ return self.mutable }
// 3476: DeclConstr
func (self *Ast_ModuleTypeInfo) InitAst_ModuleTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,scope *Ast_Scope,externalFlag bool,txt string,parentInfo LnsAny,mutable bool) {
    self.InitAst_TypeInfo(_env, scope, processInfo)
    self.externalFlag = externalFlag
    
    self.rawTxt = txt
    
    self.parentInfo = Lns_unwrapDefault( parentInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.mutable = mutable
    
    var fullName string
    if parentInfo != nil{
        parentInfo_2829 := parentInfo.(*Ast_TypeInfo)
        parentInfo_2829.FP.AddChildren(_env, &self.Ast_TypeInfo)
        var parentFull string
        parentFull = parentInfo_2829.FP.GetParentFullName(_env, Ast_defaultTypeNameCtrl, nil, nil)
        fullName = _env.LuaVM.String_format("%s.@%s", []LnsAny{parentFull, txt})
        
    } else {
        fullName = _env.LuaVM.String_format("%s", []LnsAny{txt})
        
    }
    self.fullName = fullName
    
    scope.FP.Set_ownerTypeInfo(_env, &self.Ast_TypeInfo)
}

// 3504: decl @lune.@base.@Ast.ModuleTypeInfo.equals
func (self *Ast_ModuleTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    var other *Ast_ModuleTypeInfo
    
    {
        _other := Ast_ModuleTypeInfoDownCastF(typeInfo.FP)
        if _other == nil{
            return false
        } else {
            other = _other.(*Ast_ModuleTypeInfo)
        }
    }
    return self.fullName == other.fullName
}

// 3515: decl @lune.@base.@Ast.ModuleTypeInfo.get_baseTypeInfo
func (self *Ast_ModuleTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 3519: decl @lune.@base.@Ast.ModuleTypeInfo.isModule
func (self *Ast_ModuleTypeInfo) IsModule(_env *LnsEnv) bool {
    return true
}

// 3523: decl @lune.@base.@Ast.ModuleTypeInfo.get_accessMode
func (self *Ast_ModuleTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return Ast_AccessMode__Pub
}

// 3527: decl @lune.@base.@Ast.ModuleTypeInfo.get_kind
func (self *Ast_ModuleTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Module
}

// 3531: decl @lune.@base.@Ast.ModuleTypeInfo.getParentId
func (self *Ast_ModuleTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.parentInfo.FP.Get_typeId(_env)
}

// 3536: decl @lune.@base.@Ast.ModuleTypeInfo.getTxt
func (self *Ast_ModuleTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 3542: decl @lune.@base.@Ast.ModuleTypeInfo.getTxtWithRaw
func (self *Ast_ModuleTypeInfo) GetTxtWithRaw(_env *LnsEnv, rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return rawTxt
}

// 3549: decl @lune.@base.@Ast.ModuleTypeInfo.get_display_stirng_with
func (self *Ast_ModuleTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, raw, nil, nil, nil)
}

// 3553: decl @lune.@base.@Ast.ModuleTypeInfo.get_display_stirng
func (self *Ast_ModuleTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
}

// 3556: decl @lune.@base.@Ast.ModuleTypeInfo.canEvalWith
func (self *Ast_ModuleTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return false, nil
}

// 3563: decl @lune.@base.@Ast.ModuleTypeInfo.serialize
func (self *Ast_ModuleTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    var txt string
    txt = _env.LuaVM.String_format("{ skind = %d, parentId = %d, typeId = %d, txt = '%s', ", []LnsAny{Ast_SerializeKind__Module, self.FP.GetParentId(_env).Id, self.typeId.Id, self.rawTxt})
    stream.Write(_env, txt + "\n")
    stream.Write(_env, "}\n")
}


// declaration Class -- EnumValInfo
type Ast_EnumValInfoMtd interface {
    Get_name(_env *LnsEnv) string
    Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo
    Get_val(_env *LnsEnv) LnsAny
}
type Ast_EnumValInfo struct {
    name string
    val LnsAny
    symbolInfo *Ast_SymbolInfo
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
func NewAst_EnumValInfo(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 *Ast_SymbolInfo) *Ast_EnumValInfo {
    obj := &Ast_EnumValInfo{}
    obj.FP = obj
    obj.InitAst_EnumValInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Ast_EnumValInfo) InitAst_EnumValInfo(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 *Ast_SymbolInfo) {
    self.name = arg1
    self.val = arg2
    self.symbolInfo = arg3
}
func (self *Ast_EnumValInfo) Get_name(_env *LnsEnv) string{ return self.name }
func (self *Ast_EnumValInfo) Get_val(_env *LnsEnv) LnsAny{ return self.val }
func (self *Ast_EnumValInfo) Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbolInfo }

// declaration Class -- EnumTypeInfo
type Ast_EnumTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    AddEnumValInfo(_env *LnsEnv, arg1 *Ast_EnumValInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetEnumValInfo(_env *LnsEnv, arg1 string) LnsAny
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_name2EnumValInfo(_env *LnsEnv) *LnsMap
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    Get_val2EnumValInfo(_env *LnsEnv) *LnsMap
    Get_valTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_EnumTypeInfo struct {
    Ast_TypeInfo
    externalFlag bool
    parentInfo *Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_EnumTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_Scope, arg3 bool, arg4 LnsInt, arg5 string, arg6 LnsAny, arg7 *Ast_TypeInfo) *Ast_EnumTypeInfo {
    obj := &Ast_EnumTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_EnumTypeInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Ast_EnumTypeInfo) Get_externalFlag(_env *LnsEnv) bool{ return self.externalFlag }
func (self *Ast_EnumTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_EnumTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_EnumTypeInfo) Get_rawTxt(_env *LnsEnv) string{ return self.rawTxt }
func (self *Ast_EnumTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Ast_EnumTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.nilableTypeInfo }
func (self *Ast_EnumTypeInfo) Get_valTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.valTypeInfo }
func (self *Ast_EnumTypeInfo) Get_name2EnumValInfo(_env *LnsEnv) *LnsMap{ return self.name2EnumValInfo }
func (self *Ast_EnumTypeInfo) Get_val2EnumValInfo(_env *LnsEnv) *LnsMap{ return self.val2EnumValInfo }
// 3622: DeclConstr
func (self *Ast_EnumTypeInfo) InitAst_EnumTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,scope *Ast_Scope,externalFlag bool,accessMode LnsInt,txt string,parentInfo LnsAny,valTypeInfo *Ast_TypeInfo) {
    self.InitAst_TypeInfo(_env, scope, processInfo)
    self.externalFlag = externalFlag
    
    self.accessMode = accessMode
    
    self.rawTxt = txt
    
    self.parentInfo = Lns_unwrapDefault( parentInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.name2EnumValInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.valTypeInfo = valTypeInfo
    
    self.val2EnumValInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    {
        __exp := parentInfo
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            _exp.FP.AddChildren(_env, &self.Ast_TypeInfo)
        }
    }
    self.nilableTypeInfo = &NewAst_NilableTypeInfo(_env, processInfo, &self.Ast_TypeInfo).Ast_TypeInfo
    
    scope.FP.Set_ownerTypeInfo(_env, &self.Ast_TypeInfo)
}

// 3648: decl @lune.@base.@Ast.EnumTypeInfo.isModule
func (self *Ast_EnumTypeInfo) IsModule(_env *LnsEnv) bool {
    return false
}

// 3652: decl @lune.@base.@Ast.EnumTypeInfo.get_kind
func (self *Ast_EnumTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Enum
}

// 3657: decl @lune.@base.@Ast.EnumTypeInfo.get_baseTypeInfo
func (self *Ast_EnumTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 3660: decl @lune.@base.@Ast.EnumTypeInfo.getParentId
func (self *Ast_EnumTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.parentInfo.FP.Get_typeId(_env)
}

// 3665: decl @lune.@base.@Ast.EnumTypeInfo.getTxt
func (self *Ast_EnumTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 3671: decl @lune.@base.@Ast.EnumTypeInfo.getTxtWithRaw
func (self *Ast_EnumTypeInfo) GetTxtWithRaw(_env *LnsEnv, rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return rawTxt
}

// 3678: decl @lune.@base.@Ast.EnumTypeInfo.get_display_stirng_with
func (self *Ast_EnumTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, raw, nil, nil, nil)
}

// 3682: decl @lune.@base.@Ast.EnumTypeInfo.get_display_stirng
func (self *Ast_EnumTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
}

// 3686: decl @lune.@base.@Ast.EnumTypeInfo.canEvalWith
func (self *Ast_EnumTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    if &self.Ast_TypeInfo == other.FP.Get_srcTypeInfo(_env).FP.Get_aliasSrc(_env){
        return true, nil
    }
    return false, _env.LuaVM.String_format("%d != %d", []LnsAny{self.FP.Get_typeId(_env).Id, other.FP.Get_srcTypeInfo(_env).FP.Get_aliasSrc(_env).FP.Get_typeId(_env).Id})
}

// 3697: decl @lune.@base.@Ast.EnumTypeInfo.addEnumValInfo
func (self *Ast_EnumTypeInfo) AddEnumValInfo(_env *LnsEnv, valInfo *Ast_EnumValInfo) {
    self.name2EnumValInfo.Set(valInfo.FP.Get_name(_env),valInfo)
    self.val2EnumValInfo.Set(Ast_getEnumLiteralVal(_env, valInfo.FP.Get_val(_env)),valInfo)
}

// 3702: decl @lune.@base.@Ast.EnumTypeInfo.getEnumValInfo
func (self *Ast_EnumTypeInfo) GetEnumValInfo(_env *LnsEnv, name string) LnsAny {
    return self.name2EnumValInfo.Get(name)
}

// 3706: decl @lune.@base.@Ast.EnumTypeInfo.get_mutMode
func (self *Ast_EnumTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return Ast_MutMode__Mut
}

// 6012: decl @lune.@base.@Ast.EnumTypeInfo.serialize
func (self *Ast_EnumTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    var txt string
    txt = _env.LuaVM.String_format("{ skind = %d, parentId = %d, typeId = %d, txt = '%s',\naccessMode = %d, kind = %d, valTypeId = %d, ", []LnsAny{Ast_SerializeKind__Enum, self.FP.GetParentId(_env).Id, self.typeId.Id, self.rawTxt, self.accessMode, Ast_TypeInfoKind__Enum, self.valTypeInfo.FP.Get_typeId(_env).Id})
    stream.Write(_env, txt)
    stream.Write(_env, "enumValList = {")
    {
        __collection24172 := self.name2EnumValInfo
        __sorted24172 := __collection24172.CreateKeyListStr()
        __sorted24172.Sort( LnsItemKindStr, nil )
        for _, ___key24172 := range( __sorted24172.Items ) {
            enumValInfo := __collection24172.Items[ ___key24172 ].(Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
            stream.Write(_env, _env.LuaVM.String_format("%s = ", []LnsAny{enumValInfo.FP.Get_name(_env)}))
            switch _exp24170 := enumValInfo.FP.Get_val(_env).(type) {
            case *Ast_EnumLiteral__Int:
            val := _exp24170.Val1
                stream.Write(_env, _env.LuaVM.String_format("%d,", []LnsAny{val}))
            case *Ast_EnumLiteral__Real:
            val := _exp24170.Val1
                stream.Write(_env, _env.LuaVM.String_format("%g,", []LnsAny{val}))
            case *Ast_EnumLiteral__Str:
            val := _exp24170.Val1
                stream.Write(_env, _env.LuaVM.String_format("%q,", []LnsAny{val}))
            }
        }
    }
    stream.Write(_env, "} }\n")
}


// declaration Class -- AlgeTypeInfo
type Ast_AlgeTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    AddValInfo(_env *LnsEnv, arg1 *Ast_AlgeValInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    GetValInfo(_env *LnsEnv, arg1 string) LnsAny
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    Get_valInfoMap(_env *LnsEnv) *LnsMap
    Get_valInfoNum(_env *LnsEnv) LnsInt
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_AlgeTypeInfo struct {
    Ast_TypeInfo
    externalFlag bool
    parentInfo *Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_AlgeTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_Scope, arg3 bool, arg4 LnsInt, arg5 string, arg6 LnsAny) *Ast_AlgeTypeInfo {
    obj := &Ast_AlgeTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AlgeTypeInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Ast_AlgeTypeInfo) Get_externalFlag(_env *LnsEnv) bool{ return self.externalFlag }
func (self *Ast_AlgeTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_AlgeTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_AlgeTypeInfo) Get_rawTxt(_env *LnsEnv) string{ return self.rawTxt }
func (self *Ast_AlgeTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Ast_AlgeTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.nilableTypeInfo }
func (self *Ast_AlgeTypeInfo) Get_valInfoMap(_env *LnsEnv) *LnsMap{ return self.valInfoMap }
func (self *Ast_AlgeTypeInfo) Get_valInfoNum(_env *LnsEnv) LnsInt{ return self.valInfoNum }
// 3732: decl @lune.@base.@Ast.AlgeTypeInfo.get_baseTypeInfo
func (self *Ast_AlgeTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 3736: DeclConstr
func (self *Ast_AlgeTypeInfo) InitAst_AlgeTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,scope *Ast_Scope,externalFlag bool,accessMode LnsInt,txt string,parentInfo LnsAny) {
    self.InitAst_TypeInfo(_env, scope, processInfo)
    self.externalFlag = externalFlag
    
    self.accessMode = accessMode
    
    self.rawTxt = txt
    
    self.parentInfo = Lns_unwrapDefault( parentInfo, Ast_headTypeInfo).(*Ast_TypeInfo)
    
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.valInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.valInfoNum = 0
    
    {
        __exp := parentInfo
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            _exp.FP.AddChildren(_env, &self.Ast_TypeInfo)
        }
    }
    self.nilableTypeInfo = &NewAst_NilableTypeInfo(_env, processInfo, &self.Ast_TypeInfo).Ast_TypeInfo
    
    scope.FP.Set_ownerTypeInfo(_env, &self.Ast_TypeInfo)
}

// 3758: decl @lune.@base.@Ast.AlgeTypeInfo.getValInfo
func (self *Ast_AlgeTypeInfo) GetValInfo(_env *LnsEnv, name string) LnsAny {
    return self.valInfoMap.Get(name)
}

// 3762: decl @lune.@base.@Ast.AlgeTypeInfo.isModule
func (self *Ast_AlgeTypeInfo) IsModule(_env *LnsEnv) bool {
    return false
}

// 3766: decl @lune.@base.@Ast.AlgeTypeInfo.get_kind
func (self *Ast_AlgeTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Alge
}

// 3770: decl @lune.@base.@Ast.AlgeTypeInfo.getParentId
func (self *Ast_AlgeTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.parentInfo.FP.Get_typeId(_env)
}

// 3775: decl @lune.@base.@Ast.AlgeTypeInfo.getTxt
func (self *Ast_AlgeTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 3781: decl @lune.@base.@Ast.AlgeTypeInfo.getTxtWithRaw
func (self *Ast_AlgeTypeInfo) GetTxtWithRaw(_env *LnsEnv, rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return rawTxt
}

// 3788: decl @lune.@base.@Ast.AlgeTypeInfo.get_display_stirng_with
func (self *Ast_AlgeTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, raw, nil, nil, nil)
}

// 3792: decl @lune.@base.@Ast.AlgeTypeInfo.get_display_stirng
func (self *Ast_AlgeTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
}

// 3796: decl @lune.@base.@Ast.AlgeTypeInfo.canEvalWith
func (self *Ast_AlgeTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return &self.Ast_TypeInfo == other.FP.Get_srcTypeInfo(_env).FP.Get_aliasSrc(_env), nil
}

// 3803: decl @lune.@base.@Ast.AlgeTypeInfo.get_mutMode
func (self *Ast_AlgeTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return Ast_MutMode__Mut
}

// 3826: decl @lune.@base.@Ast.AlgeTypeInfo.addValInfo
func (self *Ast_AlgeTypeInfo) AddValInfo(_env *LnsEnv, valInfo *Ast_AlgeValInfo) {
    self.valInfoMap.Set(valInfo.FP.Get_name(_env),valInfo)
    self.valInfoNum = self.valInfoNum + 1
    
}

// 6064: decl @lune.@base.@Ast.AlgeTypeInfo.serialize
func (self *Ast_AlgeTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    var txt string
    txt = _env.LuaVM.String_format("{ skind = %d, parentId = %d, typeId = %d, txt = '%s',\naccessMode = %d, kind = %d, ", []LnsAny{Ast_SerializeKind__Alge, self.FP.GetParentId(_env).Id, self.typeId.Id, self.rawTxt, self.accessMode, Ast_TypeInfoKind__Alge})
    stream.Write(_env, txt)
    stream.Write(_env, "algeValList = {")
    var firstFlag bool
    firstFlag = true
    {
        __collection24412 := self.valInfoMap
        __sorted24412 := __collection24412.CreateKeyListStr()
        __sorted24412.Sort( LnsItemKindStr, nil )
        for _, ___key24412 := range( __sorted24412.Items ) {
            algeValInfo := __collection24412.Items[ ___key24412 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            if Lns_op_not(firstFlag){
                stream.Write(_env, ",")
            } else { 
                firstFlag = false
                
            }
            algeValInfo.FP.Serialize(_env, stream, serializeInfo)
        }
    }
    stream.Write(_env, "} }\n")
}


// declaration Class -- AlgeValInfo
type Ast_AlgeValInfoMtd interface {
    Get_algeTpye(_env *LnsEnv) *Ast_AlgeTypeInfo
    Get_name(_env *LnsEnv) string
    Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo
    Get_typeList(_env *LnsEnv) *LnsList
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
}
type Ast_AlgeValInfo struct {
    name string
    typeList *LnsList
    algeTpye *Ast_AlgeTypeInfo
    symbolInfo *Ast_SymbolInfo
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
func NewAst_AlgeValInfo(_env *LnsEnv, arg1 string, arg2 *LnsList, arg3 *Ast_AlgeTypeInfo, arg4 *Ast_SymbolInfo) *Ast_AlgeValInfo {
    obj := &Ast_AlgeValInfo{}
    obj.FP = obj
    obj.InitAst_AlgeValInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *Ast_AlgeValInfo) InitAst_AlgeValInfo(_env *LnsEnv, arg1 string, arg2 *LnsList, arg3 *Ast_AlgeTypeInfo, arg4 *Ast_SymbolInfo) {
    self.name = arg1
    self.typeList = arg2
    self.algeTpye = arg3
    self.symbolInfo = arg4
}
func (self *Ast_AlgeValInfo) Get_name(_env *LnsEnv) string{ return self.name }
func (self *Ast_AlgeValInfo) Get_typeList(_env *LnsEnv) *LnsList{ return self.typeList }
func (self *Ast_AlgeValInfo) Get_algeTpye(_env *LnsEnv) *Ast_AlgeTypeInfo{ return self.algeTpye }
func (self *Ast_AlgeValInfo) Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbolInfo }
// 3814: decl @lune.@base.@Ast.AlgeValInfo.serialize
func (self *Ast_AlgeValInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    stream.Write(_env, _env.LuaVM.String_format("{ name = '%s', typeList = {", []LnsAny{self.name}))
    for _index, _typeInfo := range( self.typeList.Items ) {
        index := _index + 1
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            stream.Write(_env, ", ")
        }
        stream.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{serializeInfo.FP.SerializeId(_env, typeInfo.FP.Get_typeId(_env))}))
    }
    stream.Write(_env, "} }")
}


// declaration Class -- NormalTypeInfo
type Ast_NormalTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    EqualsSub(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_requirePath(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    Set_mutMode(_env *LnsEnv, arg1 LnsInt)
    set_requirePath(_env *LnsEnv, arg1 string)
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
    SwitchScopeTo(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_NormalTypeInfo struct {
    Ast_TypeInfo
    externalFlag bool
    itemTypeInfoList *LnsList
    argTypeInfoList *LnsList
    retTypeInfoList *LnsList
    parentInfo *Ast_TypeInfo
    typeId *Ast_IdInfo
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
    moduleLang LnsAny
    requirePath string
    asyncMode LnsInt
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
func NewAst_NormalTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 bool, arg7 bool, arg8 bool, arg9 LnsInt, arg10 string, arg11 LnsAny, arg12 LnsInt, arg13 LnsAny, arg14 LnsAny, arg15 LnsAny, arg16 LnsAny, arg17 LnsAny, arg18 LnsInt) *Ast_NormalTypeInfo {
    obj := &Ast_NormalTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_NormalTypeInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18)
    return obj
}
func (self *Ast_NormalTypeInfo) Get_externalFlag(_env *LnsEnv) bool{ return self.externalFlag }
func (self *Ast_NormalTypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList{ return self.itemTypeInfoList }
func (self *Ast_NormalTypeInfo) Get_argTypeInfoList(_env *LnsEnv) *LnsList{ return self.argTypeInfoList }
func (self *Ast_NormalTypeInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList{ return self.retTypeInfoList }
func (self *Ast_NormalTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.parentInfo }
func (self *Ast_NormalTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_NormalTypeInfo) Get_rawTxt(_env *LnsEnv) string{ return self.rawTxt }
func (self *Ast_NormalTypeInfo) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *Ast_NormalTypeInfo) Get_staticFlag(_env *LnsEnv) bool{ return self.staticFlag }
func (self *Ast_NormalTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Ast_NormalTypeInfo) Get_autoFlag(_env *LnsEnv) bool{ return self.autoFlag }
func (self *Ast_NormalTypeInfo) Get_abstractFlag(_env *LnsEnv) bool{ return self.abstractFlag }
func (self *Ast_NormalTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.baseTypeInfo }
func (self *Ast_NormalTypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList{ return self.interfaceList }
func (self *Ast_NormalTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.nilableTypeInfo }
func (self *Ast_NormalTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt{ return self.mutMode }
func (self *Ast_NormalTypeInfo) Set_mutMode(_env *LnsEnv, arg1 LnsInt){ self.mutMode = arg1 }
func (self *Ast_NormalTypeInfo) Get_requirePath(_env *LnsEnv) string{ return self.requirePath }
func (self *Ast_NormalTypeInfo) set_requirePath(_env *LnsEnv, arg1 string){ self.requirePath = arg1 }
func (self *Ast_NormalTypeInfo) Get_asyncMode(_env *LnsEnv) LnsInt{ return self.asyncMode }
// 3885: decl @lune.@base.@Ast.NormalTypeInfo.getOverridingType
func (self *Ast_NormalTypeInfo) GetOverridingType(_env *LnsEnv) LnsAny {
    switch _exp13600 := self.overridingType.(type) {
    case *Ast_OverridingType__NotOverride:
        return nil
    case *Ast_OverridingType__Override:
    typeInfo := _exp13600.Val1
        return typeInfo
    case *Ast_OverridingType__NoReady:
        var scope *Ast_Scope
        scope = Lns_unwrap( self.parentInfo.FP.Get_scope(_env)).(*Ast_Scope)
        {
            _typeInfo := scope.FP.GetTypeInfoField(_env, self.rawTxt, false, scope, Ast_ScopeAccess__Normal)
            if !Lns_IsNil( _typeInfo ) {
                typeInfo := _typeInfo.(*Ast_TypeInfo)
                {
                    _workType := typeInfo.FP.GetOverridingType(_env)
                    if !Lns_IsNil( _workType ) {
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

// 3913: decl @lune.@base.@Ast.NormalTypeInfo.switchScopeTo
func (self *Ast_NormalTypeInfo) SwitchScopeTo(_env *LnsEnv, scope *Ast_Scope) {
    self.FP.SwitchScope(_env, scope)
}

// 3917: DeclConstr
func (self *Ast_NormalTypeInfo) InitAst_NormalTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,abstractFlag bool,scope LnsAny,baseTypeInfo LnsAny,interfaceList LnsAny,autoFlag bool,externalFlag bool,staticFlag bool,accessMode LnsInt,txt string,parentInfo LnsAny,kind LnsInt,itemTypeInfoList LnsAny,argTypeInfoList LnsAny,retTypeInfoList LnsAny,mutMode LnsAny,moduleLang LnsAny,asyncMode LnsInt) {
    self.InitAst_TypeInfo(_env, scope, processInfo)
    self.asyncMode = asyncMode
    
    if Lns_type(kind) != "number"{
        Util_printStackTrace(_env)
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( kind == Ast_TypeInfoKind__Method) &&
        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(parentInfo) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.HasBase(_env)})/* 3935:30 */)) )){
        self.overridingType = Ast_OverridingType__NoReady_Obj
        
    } else { 
        self.overridingType = Ast_OverridingType__NotOverride_Obj
        
    }
    self.requirePath = ""
    
    self.moduleLang = moduleLang
    
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
    
    var setupAlt2typeMap func(_env *LnsEnv) *LnsMap
    setupAlt2typeMap = func(_env *LnsEnv) *LnsMap {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.baseTypeInfo == Ast_headTypeInfo) &&
            _env.SetStackVal( self.interfaceList.Len() == 0) ).(bool)){
            return NewLnsMap( map[LnsAny]LnsAny{})
        }
        var alt2typeMap *LnsMap
        alt2typeMap = NewLnsMap( map[LnsAny]LnsAny{})
        if _switch14062 := kind; _switch14062 == Ast_TypeInfoKind__Set || _switch14062 == Ast_TypeInfoKind__Map || _switch14062 == Ast_TypeInfoKind__List || _switch14062 == Ast_TypeInfoKind__Array || _switch14062 == Ast_TypeInfoKind__Box {
            if self.itemTypeInfoList.Len() != self.baseTypeInfo.FP.Get_itemTypeInfoList(_env).Len(){
                Util_err(_env, _env.LuaVM.String_format("unmatch generic type number -- %d, %d", []LnsAny{self.itemTypeInfoList.Len(), self.baseTypeInfo.FP.Get_itemTypeInfoList(_env).Len()}))
            }
            for _index, _appyType := range( self.itemTypeInfoList.Items ) {
                index := _index + 1
                appyType := _appyType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var genType *Ast_TypeInfo
                genType = self.baseTypeInfo.FP.Get_itemTypeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                alt2typeMap.Set(genType,appyType)
            }
        } else if _switch14062 == Ast_TypeInfoKind__Class || _switch14062 == Ast_TypeInfoKind__IF {
            for _, _ifType := range( self.interfaceList.Items ) {
                ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                {
                    _genericType := Ast_GenericTypeInfoDownCastF(ifType.FP)
                    if !Lns_IsNil( _genericType ) {
                        genericType := _genericType.(*Ast_GenericTypeInfo)
                        for _altType, _genType := range( genericType.FP.CreateAlt2typeMap(_env, false).Items ) {
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
    self.alt2typeMap = setupAlt2typeMap(_env)
    
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    if kind == Ast_TypeInfoKind__Root{
    } else { 
        if parentInfo != nil{
            parentInfo_3281 := parentInfo.(*Ast_TypeInfo)
            parentInfo_3281.FP.AddChildren(_env, &self.Ast_TypeInfo)
        }
        var hasNilable bool
        hasNilable = false
        if _switch14171 := (kind); _switch14171 == Ast_TypeInfoKind__Prim || _switch14171 == Ast_TypeInfoKind__List || _switch14171 == Ast_TypeInfoKind__Array || _switch14171 == Ast_TypeInfoKind__Set || _switch14171 == Ast_TypeInfoKind__Map || _switch14171 == Ast_TypeInfoKind__Class || _switch14171 == Ast_TypeInfoKind__Stem || _switch14171 == Ast_TypeInfoKind__Module || _switch14171 == Ast_TypeInfoKind__IF {
            hasNilable = true
            
        } else if _switch14171 == Ast_TypeInfoKind__Func || _switch14171 == Ast_TypeInfoKind__Method || _switch14171 == Ast_TypeInfoKind__Form || _switch14171 == Ast_TypeInfoKind__FormFunc {
            hasNilable = true
            
        }
        if hasNilable{
            self.nilableTypeInfo = &NewAst_NilableTypeInfo(_env, processInfo, &self.Ast_TypeInfo).Ast_TypeInfo
            
        } else { 
            self.nilableTypeInfo = Ast_headTypeInfo
            
        }
    }
}

// 4022: decl @lune.@base.@Ast.NormalTypeInfo.createAlt2typeMap
func (self *Ast_NormalTypeInfo) CreateAlt2typeMap(_env *LnsEnv, detectFlag bool) *LnsMap {
    var _map *LnsMap
    _map = self.baseTypeInfo.FP.CreateAlt2typeMap(_env, detectFlag)
    for _genType, _typeInfo := range( self.alt2typeMap.Items ) {
        genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        _map.Set(genType,typeInfo)
    }
    return _map
}

// 4032: decl @lune.@base.@Ast.NormalTypeInfo.get_nilable
func (self *Ast_NormalTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return false
}

// 4036: decl @lune.@base.@Ast.NormalTypeInfo.isModule
func (self *Ast_NormalTypeInfo) IsModule(_env *LnsEnv) bool {
    return false
}

// 4040: decl @lune.@base.@Ast.NormalTypeInfo.getParentId
func (self *Ast_NormalTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.parentInfo.FP.Get_typeId(_env)
}

// 4044: decl @lune.@base.@Ast.NormalTypeInfo.get_baseId
func (self *Ast_NormalTypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return self.baseTypeInfo.FP.Get_typeId(_env)
}

// 4049: decl @lune.@base.@Ast.NormalTypeInfo.getTxt
func (self *Ast_NormalTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 4055: decl @lune.@base.@Ast.NormalTypeInfo.getTxtWithRaw
func (self *Ast_NormalTypeInfo) GetTxtWithRaw(_env *LnsEnv, raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    var parentTxt string
    parentTxt = ""
    if typeNameCtrl != nil{
        typeNameCtrl_3336 := typeNameCtrl.(*Ast_TypeNameCtrl)
        parentTxt = self.FP.GetParentFullName(_env, typeNameCtrl_3336, importInfo, localFlag)
        
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
            txt = txt + typeInfo.FP.GetTxt(_env, typeNameCtrl, importInfo, localFlag)
            
        }
        name = parentTxt + txt + ">"
        
    } else { 
        name = parentTxt + raw
        
    }
    return name
}

// 4083: decl @lune.@base.@Ast.NormalTypeInfo.get_display_stirng_with
func (self *Ast_NormalTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    if _switch14589 := self.kind; _switch14589 == Ast_TypeInfoKind__Func || _switch14589 == Ast_TypeInfoKind__Form || _switch14589 == Ast_TypeInfoKind__FormFunc || _switch14589 == Ast_TypeInfoKind__Method || _switch14589 == Ast_TypeInfoKind__Macro {
        var txt string
        txt = raw + "("
        for _index, _argType := range( self.argTypeInfoList.Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                txt = txt + ", "
                
            }
            txt = txt + argType.FP.Get_display_stirng(_env)
            
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
            txt = txt + retType.FP.Get_display_stirng(_env)
            
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
            txt = txt + typeInfo.FP.Get_display_stirng_with(_env, typeInfo.FP.Get_rawTxt(_env), alt2type)
            
        }
        name = parentTxt + txt + ">"
        
    } else { 
        name = parentTxt + raw
        
    }
    return name
}

// 4128: decl @lune.@base.@Ast.NormalTypeInfo.get_display_stirng
func (self *Ast_NormalTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
}

// 4133: decl @lune.@base.@Ast.NormalTypeInfo.serialize
func (self *Ast_NormalTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    if self.typeId.Id == Ast_rootTypeId{
        return 
    }
    var parentId *Ast_IdInfo
    parentId = self.FP.GetParentId(_env)
    var txt string
    txt = _env.LuaVM.String_format("{ skind=%d, parentId = %d, typeId = %d, baseId = %s, txt = '%s',\n  abstractFlag = %s, staticFlag = %s, accessMode = %d, kind = %d, mutMode = %d,\n  asyncMode = %d,     \n", []LnsAny{Ast_SerializeKind__Normal, parentId.Id, self.typeId.Id, serializeInfo.FP.SerializeId(_env, self.FP.Get_baseId(_env)), self.rawTxt, self.abstractFlag, self.staticFlag, self.accessMode, self.kind, self.mutMode, self.asyncMode})
    {
        __exp := self.moduleLang
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(LnsInt)
            txt = txt + _env.LuaVM.String_format("moduleLang = %d, ", []LnsAny{_exp})
            
        }
    }
    if self.requirePath != ""{
        txt = txt + _env.LuaVM.String_format("requirePath = \"%s\", ", []LnsAny{self.requirePath})
        
    }
    var children *LnsList
    children = NewLnsList([]LnsAny{})
    for _, _child := range( self.FP.Get_children(_env).Items ) {
        child := _child.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if serializeInfo.FP.IsValidChildren(_env, child.FP.Get_typeId(_env)){
            children.Insert(Ast_TypeInfo2Stem(child))
        }
    }
    stream.Write(_env, txt + self.FP.SerializeTypeInfoList(_env, serializeInfo, "itemTypeId = {", self.itemTypeInfoList, nil) + self.FP.SerializeTypeInfoList(_env, serializeInfo, "ifList = {", self.interfaceList, nil) + self.FP.SerializeTypeInfoList(_env, serializeInfo, "argTypeId = {", self.argTypeInfoList, nil) + self.FP.SerializeTypeInfoList(_env, serializeInfo, "retTypeId = {", self.retTypeInfoList, nil) + self.FP.SerializeTypeInfoList(_env, serializeInfo, "children = {", children, true) + "}\n")
}

// 4177: decl @lune.@base.@Ast.NormalTypeInfo.equalsSub
func (self *Ast_NormalTypeInfo) EqualsSub(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    if self.typeId.FP.Equals(_env, typeInfo.FP.Get_typeId(_env)){
        return true
    }
    if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
        return typeInfo.FP.Equals(_env, processInfo, &self.Ast_TypeInfo, alt2type, checkModifer)
    }
    {
        _aliasType := Ast_AliasTypeInfoDownCastF(typeInfo.FP)
        if !Lns_IsNil( _aliasType ) {
            aliasType := _aliasType.(*Ast_AliasTypeInfo)
            return aliasType.FP.Equals(_env, processInfo, &self.Ast_TypeInfo, alt2type, checkModifer)
        }
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.kind != typeInfo.FP.Get_kind(_env)) ||
        _env.SetStackVal( self.staticFlag != typeInfo.FP.Get_staticFlag(_env)) ||
        _env.SetStackVal( self.autoFlag != typeInfo.FP.Get_autoFlag(_env)) ||
        _env.SetStackVal( self.FP.Get_nilable(_env) != typeInfo.FP.Get_nilable(_env)) ||
        _env.SetStackVal( self.rawTxt != typeInfo.FP.Get_rawTxt(_env)) ||
        _env.SetStackVal( self.parentInfo != typeInfo.FP.Get_parentInfo(_env)) ||
        _env.SetStackVal( self.baseTypeInfo != typeInfo.FP.Get_baseTypeInfo(_env)) ).(bool){
        return false
    }
    if self.accessMode != typeInfo.FP.Get_accessMode(_env){
        if _switch15116 := self.kind; _switch15116 == Ast_TypeInfoKind__List || _switch15116 == Ast_TypeInfoKind__Map || _switch15116 == Ast_TypeInfoKind__Array || _switch15116 == Ast_TypeInfoKind__Set {
        } else {
            return false
        }
    }
    {
        if self.itemTypeInfoList.Len() != typeInfo.FP.Get_itemTypeInfoList(_env).Len(){
            return false
        }
        for _index, _item := range( self.itemTypeInfoList.Items ) {
            index := _index + 1
            item := _item.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(item.FP.Equals(_env, processInfo, typeInfo.FP.Get_itemTypeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type, checkModifer)){
                return false
            }
        }
    }
    {
        if self.retTypeInfoList.Len() != typeInfo.FP.Get_retTypeInfoList(_env).Len(){
            return false
        }
        for _index, _item := range( self.retTypeInfoList.Items ) {
            index := _index + 1
            item := _item.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(item.FP.Equals(_env, processInfo, typeInfo.FP.Get_retTypeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), alt2type, checkModifer)){
                return false
            }
        }
    }
    return true
}

// 4247: decl @lune.@base.@Ast.NormalTypeInfo.equals
func (self *Ast_NormalTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    return self.FP.EqualsSub(_env, processInfo, typeInfo, alt2type, checkModifer)
}

// 4261: decl @lune.@base.@Ast.NormalTypeInfo.create
func Ast_NormalTypeInfo_create(_env *LnsEnv, processInfo *Ast_ProcessInfo,accessMode LnsInt,abstractFlag bool,scope LnsAny,baseInfo *Ast_TypeInfo,parentInfo *Ast_TypeInfo,staticFlag bool,kind LnsInt,txt string,itemTypeInfo *LnsList,argTypeInfoList *LnsList,retTypeInfoList *LnsList,mutMode LnsAny,asyncMode LnsInt) *Ast_TypeInfo {
    if _switch15330 := kind; _switch15330 == Ast_TypeInfoKind__Class || _switch15330 == Ast_TypeInfoKind__ExtModule || _switch15330 == Ast_TypeInfoKind__IF {
        Util_err(_env, "can't use create() method. use createClass(), createExtModule().")
    }
    if kind == Ast_TypeInfoKind__Prim{
        {
            __exp := Ast_sym2builtInTypeMap.Get(txt)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_SymbolInfo)
                return _exp.FP.Get_typeInfo(_env)
            }
        }
        Util_err(_env, _env.LuaVM.String_format("not found symbol -- %s", []LnsAny{txt}))
    }
    var info *Ast_NormalTypeInfo
    info = NewAst_NormalTypeInfo(_env, processInfo, abstractFlag, scope, baseInfo, nil, false, true, staticFlag, accessMode, txt, parentInfo, kind, itemTypeInfo, argTypeInfoList, retTypeInfoList, mutMode, nil, asyncMode)
    return &info.Ast_TypeInfo
}

// 4481: decl @lune.@base.@Ast.NormalTypeInfo.createBuiltin
func Ast_NormalTypeInfo_createBuiltin_3270_(_env *LnsEnv, idName string,typeTxt string,kind LnsInt,typeDDD LnsAny,ifList LnsAny) *Ast_TypeInfo {
    var argTypeList *LnsList
    argTypeList = NewLnsList([]LnsAny{})
    var retTypeList *LnsList
    retTypeList = NewLnsList([]LnsAny{})
    if typeTxt == "form"{
        {
            __exp := typeDDD
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_TypeInfo)
                argTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(_exp)})
                
                retTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(_exp)})
                
            }
        }
    }
    var scope LnsAny
    scope = nil
    if _switch16512 := kind; _switch16512 == Ast_TypeInfoKind__Array || _switch16512 == Ast_TypeInfoKind__List || _switch16512 == Ast_TypeInfoKind__Set || _switch16512 == Ast_TypeInfoKind__Class || _switch16512 == Ast_TypeInfoKind__Module || _switch16512 == Ast_TypeInfoKind__IF || _switch16512 == Ast_TypeInfoKind__Form || _switch16512 == Ast_TypeInfoKind__FormFunc || _switch16512 == Ast_TypeInfoKind__Func || _switch16512 == Ast_TypeInfoKind__Method || _switch16512 == Ast_TypeInfoKind__Macro {
        scope = NewAst_Scope(_env, Ast_rootProcessInfo, Ast_rootScope, _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( kind == Ast_TypeInfoKind__Class) ||
            _env.SetStackVal( kind == Ast_TypeInfoKind__Module) ||
            _env.SetStackVal( kind == Ast_TypeInfoKind__IF) ||
            _env.SetStackVal( kind == Ast_TypeInfoKind__List) ||
            _env.SetStackVal( kind == Ast_TypeInfoKind__Array) ||
            _env.SetStackVal( kind == Ast_TypeInfoKind__Set) ).(bool), nil, nil)
        
    }
    var genTypeList *LnsList
    genTypeList = NewLnsList([]LnsAny{})
    if _switch16610 := kind; _switch16610 == Ast_TypeInfoKind__Array || _switch16610 == Ast_TypeInfoKind__List || _switch16610 == Ast_TypeInfoKind__Set {
        genTypeList.Insert(Ast_AlternateTypeInfo2Stem(Ast_rootProcessInfo.FP.CreateAlternate(_env, true, 1, "T", Ast_AccessMode__Pri, Ast_headTypeInfo, nil, nil)))
    } else if _switch16610 == Ast_TypeInfoKind__Map {
        genTypeList.Insert(Ast_AlternateTypeInfo2Stem(Ast_rootProcessInfo.FP.CreateAlternate(_env, true, 1, "K", Ast_AccessMode__Pri, Ast_headTypeInfo, nil, nil)))
        genTypeList.Insert(Ast_AlternateTypeInfo2Stem(Ast_rootProcessInfo.FP.CreateAlternate(_env, true, 2, "V", Ast_AccessMode__Pri, Ast_headTypeInfo, nil, nil)))
    }
    var info *Ast_NormalTypeInfo
    info = NewAst_NormalTypeInfo(_env, Ast_rootProcessInfo, false, scope, nil, ifList, false, false, false, Ast_AccessMode__Pub, typeTxt, Ast_headTypeInfo, kind, genTypeList, argTypeList, retTypeList, Ast_MutMode__Mut, nil, Ast_Async__Async)
    Ast_registBuiltin_3265_(_env, idName, typeTxt, kind, &info.Ast_TypeInfo, Ast_headTypeInfo, Ast_Scope2Stem(scope) != nil)
    return &info.Ast_TypeInfo
}

// 6190: decl @lune.@base.@Ast.NormalTypeInfo.isAvailableMapping
func Ast_NormalTypeInfo_isAvailableMapping(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,checkedTypeMap *LnsMap) bool {
    var isAvailableMappingSub func(_env *LnsEnv) bool
    isAvailableMappingSub = func(_env *LnsEnv) bool {
        if _switch25073 := typeInfo.FP.Get_kind(_env); _switch25073 == Ast_TypeInfoKind__Prim || _switch25073 == Ast_TypeInfoKind__Enum {
            return true
        } else if _switch25073 == Ast_TypeInfoKind__Alge {
            var algeTypeInfo *Ast_AlgeTypeInfo
            algeTypeInfo = Lns_unwrap( (Ast_AlgeTypeInfoDownCastF(typeInfo.FP))).(*Ast_AlgeTypeInfo)
            for _, _valInfo := range( algeTypeInfo.FP.Get_valInfoMap(_env).Items ) {
                valInfo := _valInfo.(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
                for _, _paramType := range( valInfo.FP.Get_typeList(_env).Items ) {
                    paramType := _paramType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if Lns_op_not(Ast_NormalTypeInfo_isAvailableMapping(_env, processInfo, paramType, checkedTypeMap)){
                        return false
                    }
                }
            }
            return true
        } else if _switch25073 == Ast_TypeInfoKind__Stem {
            return true
        } else if _switch25073 == Ast_TypeInfoKind__Class || _switch25073 == Ast_TypeInfoKind__IF {
            if typeInfo.FP.Equals(_env, processInfo, Ast_builtinTypeString, nil, nil){
                return true
            }
            return typeInfo.FP.IsInheritFrom(_env, processInfo, Ast_builtinTypeMapping, nil)
        } else if _switch25073 == Ast_TypeInfoKind__Alternate {
            return typeInfo.FP.IsInheritFrom(_env, processInfo, Ast_builtinTypeMapping, nil)
        } else if _switch25073 == Ast_TypeInfoKind__List || _switch25073 == Ast_TypeInfoKind__Array || _switch25073 == Ast_TypeInfoKind__Set {
            return Ast_NormalTypeInfo_isAvailableMapping(_env, processInfo, typeInfo.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), checkedTypeMap)
        } else if _switch25073 == Ast_TypeInfoKind__Map {
            if Ast_NormalTypeInfo_isAvailableMapping(_env, processInfo, typeInfo.FP.Get_itemTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), checkedTypeMap){
                var keyType *Ast_TypeInfo
                keyType = typeInfo.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( keyType.FP.Equals(_env, processInfo, Ast_builtinTypeString, nil, nil)) ||
                    _env.SetStackVal( keyType.FP.Get_kind(_env) == Ast_TypeInfoKind__Prim) ||
                    _env.SetStackVal( keyType.FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) ).(bool){
                    return true
                }
            }
            return false
        } else if _switch25073 == Ast_TypeInfoKind__Nilable {
            return Ast_NormalTypeInfo_isAvailableMapping(_env, processInfo, typeInfo.FP.Get_nonnilableType(_env), checkedTypeMap)
        } else {
            return false
        }
    // insert a dummy
        return false
    }
    typeInfo = typeInfo.FP.Get_srcTypeInfo(_env)
    
    {
        __exp := checkedTypeMap.Get(typeInfo)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(bool)
            return _exp
        }
    }
    checkedTypeMap.Set(typeInfo,true)
    var result bool
    result = isAvailableMappingSub(_env)
    checkedTypeMap.Set(typeInfo,result)
    return result
}

// 6264: decl @lune.@base.@Ast.NormalTypeInfo.isInheritFrom
func (self *Ast_NormalTypeInfo) IsInheritFrom(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    if self.FP.Get_typeId(_env).FP.Equals(_env, other.FP.Get_typeId(_env)){
        return true
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.FP.Get_kind(_env) != Ast_TypeInfoKind__Class) &&
            _env.SetStackVal( self.FP.Get_kind(_env) != Ast_TypeInfoKind__IF) ).(bool))) ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( other.FP.Get_kind(_env) != Ast_TypeInfoKind__Class) &&
            _env.SetStackVal( other.FP.Get_kind(_env) != Ast_TypeInfoKind__IF) ).(bool))) ).(bool){
        if other == Ast_builtinTypeMapping{
            return Ast_NormalTypeInfo_isAvailableMapping(_env, processInfo, &self.Ast_TypeInfo, NewLnsMap( map[LnsAny]LnsAny{}))
        }
        return false
    }
    return Ast_TypeInfo_isInherit(_env, processInfo, &self.Ast_TypeInfo, other, alt2type)
}

// 6972: decl @lune.@base.@Ast.NormalTypeInfo.canEvalWith
func (self *Ast_NormalTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return Ast_TypeInfo_canEvalWithBase(_env, processInfo, &self.Ast_TypeInfo, Ast_TypeInfo_isMut(_env, &self.Ast_TypeInfo), other, canEvalType, alt2type)
}

// 6996: decl @lune.@base.@Ast.NormalTypeInfo.applyGeneric
func (self *Ast_NormalTypeInfo) ApplyGeneric(_env *LnsEnv, alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var itemTypeInfoList *LnsList
    var needNew bool
    
    {
        _itemTypeInfoList, _needNew := Ast_applyGenericList_3567_(_env, self.itemTypeInfoList, alt2typeMap, moduleTypeInfo)
        if _itemTypeInfoList == nil{
            return nil
        } else {
            itemTypeInfoList = _itemTypeInfoList.(*LnsList)
            needNew = _needNew
        }
    }
    var processInfo *Ast_ProcessInfo
    processInfo = moduleTypeInfo.FP.getProcessInfo(_env)
    if _switch29665 := self.FP.Get_kind(_env); _switch29665 == Ast_TypeInfoKind__Set {
        if Lns_op_not(needNew){
            return &self.Ast_TypeInfo
        }
        return processInfo.FP.CreateSet(_env, self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode)
    } else if _switch29665 == Ast_TypeInfoKind__List {
        if Lns_op_not(needNew){
            return &self.Ast_TypeInfo
        }
        return processInfo.FP.CreateList(_env, self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode)
    } else if _switch29665 == Ast_TypeInfoKind__Array {
        if Lns_op_not(needNew){
            return &self.Ast_TypeInfo
        }
        return processInfo.FP.CreateArray(_env, self.accessMode, self.parentInfo, itemTypeInfoList, self.mutMode)
    } else if _switch29665 == Ast_TypeInfoKind__Map {
        if Lns_op_not(needNew){
            return &self.Ast_TypeInfo
        }
        return processInfo.FP.CreateMap(_env, self.accessMode, self.parentInfo, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), itemTypeInfoList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), self.mutMode)
    } else if _switch29665 == Ast_TypeInfoKind__Func || _switch29665 == Ast_TypeInfoKind__Form || _switch29665 == Ast_TypeInfoKind__FormFunc {
        var argTypeInfoList *LnsList
        var workArg bool
        
        {
            _argTypeInfoList, _workArg := Ast_applyGenericList_3567_(_env, self.argTypeInfoList, alt2typeMap, moduleTypeInfo)
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
            _retTypeInfoList, _workRet := Ast_applyGenericList_3567_(_env, self.retTypeInfoList, alt2typeMap, moduleTypeInfo)
            if _retTypeInfoList == nil{
                return nil
            } else {
                retTypeInfoList = _retTypeInfoList.(*LnsList)
                workRet = _workRet
            }
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( needNew) ||
            _env.SetStackVal( workArg) ||
            _env.SetStackVal( workRet) ).(bool){
            return &processInfo.FP.CreateFunc(_env, self.abstractFlag, false, Ast_getScope(_env, &self.Ast_TypeInfo), self.kind, self.parentInfo, self.autoFlag, self.externalFlag, self.staticFlag, self.accessMode, self.rawTxt, self.asyncMode, itemTypeInfoList, argTypeInfoList, retTypeInfoList, Ast_TypeInfo_isMut(_env, &self.Ast_TypeInfo)).Ast_TypeInfo
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
    Clone(_env *LnsEnv) *Ast_TypeInfo2Map
    Get_BoxMap(_env *LnsEnv) *LnsMap
    Get_DDDMap(_env *LnsEnv) *LnsMap
    Get_ExtDDDMap(_env *LnsEnv) *LnsMap
    Get_ExtMap(_env *LnsEnv) *LnsMap
    Get_ImutModifierMap(_env *LnsEnv) *LnsMap
    Get_MutModifierMap(_env *LnsEnv) *LnsMap
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
func NewAst_TypeInfo2Map(_env *LnsEnv) *Ast_TypeInfo2Map {
    obj := &Ast_TypeInfo2Map{}
    obj.FP = obj
    obj.InitAst_TypeInfo2Map(_env)
    return obj
}
func (self *Ast_TypeInfo2Map) Get_ImutModifierMap(_env *LnsEnv) *LnsMap{ return self.ImutModifierMap }
func (self *Ast_TypeInfo2Map) Get_MutModifierMap(_env *LnsEnv) *LnsMap{ return self.MutModifierMap }
func (self *Ast_TypeInfo2Map) Get_BoxMap(_env *LnsEnv) *LnsMap{ return self.BoxMap }
func (self *Ast_TypeInfo2Map) Get_DDDMap(_env *LnsEnv) *LnsMap{ return self.DDDMap }
func (self *Ast_TypeInfo2Map) Get_ExtDDDMap(_env *LnsEnv) *LnsMap{ return self.ExtDDDMap }
func (self *Ast_TypeInfo2Map) Get_ExtMap(_env *LnsEnv) *LnsMap{ return self.ExtMap }
// 4309: DeclConstr
func (self *Ast_TypeInfo2Map) InitAst_TypeInfo2Map(_env *LnsEnv) {
    self.ImutModifierMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.MutModifierMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.BoxMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.DDDMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.ExtDDDMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.ExtMap = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 4318: decl @lune.@base.@Ast.TypeInfo2Map.clone
func (self *Ast_TypeInfo2Map) Clone(_env *LnsEnv) *Ast_TypeInfo2Map {
    var obj *Ast_TypeInfo2Map
    obj = NewAst_TypeInfo2Map(_env)
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
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extTypeFlag(_env *LnsEnv) bool
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_DDDTypeInfo struct {
    Ast_TypeInfo
    typeInfo *Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_DDDTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 bool, arg4 LnsAny) *Ast_DDDTypeInfo {
    obj := &Ast_DDDTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_DDDTypeInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *Ast_DDDTypeInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *Ast_DDDTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_DDDTypeInfo) Get_externalFlag(_env *LnsEnv) bool{ return self.externalFlag }
func (self *Ast_DDDTypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList{ return self.itemTypeInfoList }
func (self *Ast_DDDTypeInfo) Get_extedType(_env *LnsEnv) *Ast_TypeInfo{ return self.extedType }
// 5077: decl @lune.@base.@Ast.DDDTypeInfo.get_extTypeFlag
func (self *Ast_DDDTypeInfo) Get_extTypeFlag(_env *LnsEnv) bool {
    return self.extedType != &self.Ast_TypeInfo
}

// 5081: decl @lune.@base.@Ast.DDDTypeInfo.get_scope
func (self *Ast_DDDTypeInfo) Get_scope(_env *LnsEnv) LnsAny {
    return nil
}

// 5086: decl @lune.@base.@Ast.DDDTypeInfo.get_baseTypeInfo
func (self *Ast_DDDTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5090: decl @lune.@base.@Ast.DDDTypeInfo.get_parentInfo
func (self *Ast_DDDTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5094: DeclConstr
func (self *Ast_DDDTypeInfo) InitAst_DDDTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,externalFlag bool,extOrgDDType LnsAny) {
    self.InitAst_TypeInfo(_env, nil, processInfo)
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.typeInfo = typeInfo
    
    self.externalFlag = externalFlag
    
    self.itemTypeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.typeInfo)})
    
    var extOrgType *Ast_DDDTypeInfo
    if extOrgDDType != nil{
        extOrgDDType_3997 := extOrgDDType.(*Ast_DDDTypeInfo)
        extOrgType = extOrgDDType_3997
        
        processInfo.FP.get_typeInfo2Map(_env).ExtDDDMap.Set(typeInfo,self)
    } else {
        extOrgType = self
        
        processInfo.FP.get_typeInfo2Map(_env).DDDMap.Set(typeInfo,self)
    }
    self.extedType = &extOrgType.Ast_TypeInfo
    
}

// 5114: decl @lune.@base.@Ast.DDDTypeInfo.isModule
func (self *Ast_DDDTypeInfo) IsModule(_env *LnsEnv) bool {
    return false
}

// 5118: decl @lune.@base.@Ast.DDDTypeInfo.canEvalWith
func (self *Ast_DDDTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return self.typeInfo.FP.CanEvalWith(_env, processInfo, other, canEvalType, alt2type)
}

// 5125: decl @lune.@base.@Ast.DDDTypeInfo.serialize
func (self *Ast_DDDTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    stream.Write(_env, _env.LuaVM.String_format("{ skind=%d, typeId = %d, itemTypeId = %s, parentId = %d, extTypeFlag = %s }\n", []LnsAny{Ast_SerializeKind__DDD, self.typeId.Id, serializeInfo.FP.SerializeId(_env, self.typeInfo.FP.Get_typeId(_env)), Ast_headTypeInfo.FP.Get_typeId(_env).Id, self.FP.Get_extTypeFlag(_env)}))
}

// 5133: decl @lune.@base.@Ast.DDDTypeInfo.get_display_stirng_with
func (self *Ast_DDDTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    var txt string
    txt = self.FP.GetTxtWithRaw(_env, raw, nil, nil, nil)
    return txt
}

// 5139: decl @lune.@base.@Ast.DDDTypeInfo.get_display_stirng
func (self *Ast_DDDTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    var txt string
    txt = self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
    return txt
}

// 5144: decl @lune.@base.@Ast.DDDTypeInfo.getModule
func (self *Ast_DDDTypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.FP.Get_typeInfo(_env).FP.GetModule(_env)
}

// 5147: decl @lune.@base.@Ast.DDDTypeInfo.get_rawTxt
func (self *Ast_DDDTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return self.FP.GetTxt(_env, nil, nil, nil)
}

// 5150: decl @lune.@base.@Ast.DDDTypeInfo.get_kind
func (self *Ast_DDDTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__DDD
}

// 5153: decl @lune.@base.@Ast.DDDTypeInfo.get_nilable
func (self *Ast_DDDTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return true
}

// 5157: decl @lune.@base.@Ast.DDDTypeInfo.get_nilableTypeInfo
func (self *Ast_DDDTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5160: decl @lune.@base.@Ast.DDDTypeInfo.get_mutMode
func (self *Ast_DDDTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return self.FP.Get_typeInfo(_env).FP.Get_mutMode(_env)
}

// 5163: decl @lune.@base.@Ast.DDDTypeInfo.get_aliasSrc
func (self *Ast_DDDTypeInfo) Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5166: decl @lune.@base.@Ast.DDDTypeInfo.get_srcTypeInfo
func (self *Ast_DDDTypeInfo) Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5169: decl @lune.@base.@Ast.DDDTypeInfo.get_accessMode
func (self *Ast_DDDTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return Ast_AccessMode__Pub
}

// 5546: decl @lune.@base.@Ast.DDDTypeInfo.getTxt
func (self *Ast_DDDTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, "...", typeNameCtrl, importInfo, localFlag)
}

// 5552: decl @lune.@base.@Ast.DDDTypeInfo.getTxtWithRaw
func (self *Ast_DDDTypeInfo) GetTxtWithRaw(_env *LnsEnv, raw string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    if self.typeInfo == Ast_builtinTypeStem_{
        if self.FP.Get_extTypeFlag(_env){
            return "Luaval<...>"
        }
        return "..."
    }
    var typeInfo *Ast_TypeInfo
    if self.FP.Get_extTypeFlag(_env){
        typeInfo = self.typeInfo.FP.Get_extedType(_env)
        
    } else { 
        typeInfo = self.typeInfo
        
    }
    var txt string
    txt = _env.LuaVM.String_format("...<%s>", []LnsAny{typeInfo.FP.GetTxt(_env, typeNameCtrl, importInfo, localFlag)})
    if self.FP.Get_extTypeFlag(_env){
        return _env.LuaVM.String_format("Luaval<%s>", []LnsAny{txt})
    }
    return txt
}


// declaration Class -- CombineType
type Ast_CombineTypeMtd interface {
    andIfSet(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *LnsSet, arg3 *LnsMap)
    AndType(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 LnsAny, arg3 *LnsMap) LnsAny
    CreateStem(_env *LnsEnv, arg1 *Ast_ProcessInfo) *Ast_TypeInfo
    Get_typeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo) *Ast_TypeInfo
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
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
func NewAst_CombineType(_env *LnsEnv, arg1 *Ast_TypeInfo) *Ast_CombineType {
    obj := &Ast_CombineType{}
    obj.FP = obj
    obj.InitAst_CombineType(_env, arg1)
    return obj
}
// 5258: DeclConstr
func (self *Ast_CombineType) InitAst_CombineType(_env *LnsEnv, typeInfo *Ast_TypeInfo) {
    self.ifSet = NewLnsSet([]LnsAny{})
    
    for _, _iftype := range( typeInfo.FP.Get_interfaceList(_env).Items ) {
        iftype := _iftype.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        self.ifSet.Add(Ast_TypeInfo2Stem(iftype))
    }
    self.nilable = typeInfo.FP.Get_nilable(_env)
    
    self.mutMode = typeInfo.FP.Get_mutMode(_env)
    
}

// 5267: decl @lune.@base.@Ast.CombineType.isInheritFrom
func (self *Ast_CombineType) IsInheritFrom(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,alt2type LnsAny) bool {
    for _ifType := range( self.ifSet.Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if ifType.FP.IsInheritFrom(_env, processInfo, other, alt2type){
            return true
        }
    }
    return false
}

// 5279: decl @lune.@base.@Ast.CombineType.andIfSet
func (self *Ast_CombineType) andIfSet(_env *LnsEnv, processInfo *Ast_ProcessInfo,ifSet *LnsSet,alt2type *LnsMap) {
    var workSet *LnsSet
    workSet = NewLnsSet([]LnsAny{})
    for _other := range( ifSet.Items ) {
        other := _other.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if self.FP.IsInheritFrom(_env, processInfo, other, alt2type){
            workSet.Add(Ast_TypeInfo2Stem(other))
        } else { 
            for _ifType := range( self.ifSet.Items ) {
                ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if other.FP.IsInheritFrom(_env, processInfo, ifType, alt2type){
                    workSet.Add(Ast_TypeInfo2Stem(ifType))
                }
            }
        }
    }
    self.ifSet = workSet
    
}

// 5298: decl @lune.@base.@Ast.CombineType.createStem
func (self *Ast_CombineType) CreateStem(_env *LnsEnv, processInfo *Ast_ProcessInfo) *Ast_TypeInfo {
    var retType *Ast_TypeInfo
    if self.nilable{
        retType = Ast_builtinTypeStem_
        
    } else { 
        retType = Ast_builtinTypeStem
        
    }
    if Ast_isMutable(_env, self.mutMode){
        return retType
    }
    return processInfo.FP.CreateModifier(_env, retType, self.mutMode)
}

// 5312: decl @lune.@base.@Ast.CombineType.get_typeInfo
func (self *Ast_CombineType) Get_typeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo) *Ast_TypeInfo {
    if self.ifSet.Len() != 1{
        return self.FP.CreateStem(_env, processInfo)
    }
    for _ifType := range( self.ifSet.Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var work *Ast_TypeInfo
        work = ifType
        if self.nilable{
            work = work.FP.Get_nilableTypeInfo(_env)
            
        }
        if Ast_isMutable(_env, self.mutMode){
            return work
        }
        return processInfo.FP.CreateModifier(_env, work, self.mutMode)
    }
    panic("illegal")
// insert a dummy
    return nil
}

// 5343: decl @lune.@base.@Ast.CombineType.andType
func (self *Ast_CombineType) AndType(_env *LnsEnv, processInfo *Ast_ProcessInfo,other LnsAny,alt2type *LnsMap) LnsAny {
    switch _exp21215 := other.(type) {
    case *Ast_CommonType__Combine:
    comboInfo := _exp21215.Val1
        self.FP.andIfSet(_env, processInfo, comboInfo.ifSet, alt2type)
        if Lns_op_not(Ast_isMutable(_env, comboInfo.mutMode)){
            self.mutMode = comboInfo.mutMode
            
        }
        return &Ast_CommonType__Combine{self}
    case *Ast_CommonType__Normal:
    typeInfo := _exp21215.Val1
        if Lns_op_not(Ast_isMutable(_env, typeInfo.FP.Get_mutMode(_env))){
            self.mutMode = typeInfo.FP.Get_mutMode(_env)
            
        }
        var ifSet *LnsSet
        ifSet = NewLnsSet([]LnsAny{})
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
            ifSet.Add(Ast_TypeInfo2Stem(typeInfo))
        } else { 
            for _, _iftype := range( typeInfo.FP.Get_interfaceList(_env).Items ) {
                iftype := _iftype.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                ifSet.Add(Ast_TypeInfo2Stem(iftype))
            }
        }
        self.FP.andIfSet(_env, processInfo, ifSet, alt2type)
        if self.ifSet.Len() != 0{
            return &Ast_CommonType__Combine{self}
        }
        return &Ast_CommonType__Normal{self.FP.CreateStem(_env, processInfo)}
    }
// insert a dummy
    return nil
}


// declaration Class -- AbbrTypeInfo
type Ast_AbbrTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_AbbrTypeInfo struct {
    Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_AbbrTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string) *Ast_AbbrTypeInfo {
    obj := &Ast_AbbrTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AbbrTypeInfo(_env, arg1, arg2)
    return obj
}
func (self *Ast_AbbrTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_AbbrTypeInfo) Get_rawTxt(_env *LnsEnv) string{ return self.rawTxt }
// 5643: decl @lune.@base.@Ast.AbbrTypeInfo.get_scope
func (self *Ast_AbbrTypeInfo) Get_scope(_env *LnsEnv) LnsAny {
    return nil
}

// 5648: decl @lune.@base.@Ast.AbbrTypeInfo.get_baseTypeInfo
func (self *Ast_AbbrTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5652: decl @lune.@base.@Ast.AbbrTypeInfo.get_parentInfo
func (self *Ast_AbbrTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5656: DeclConstr
func (self *Ast_AbbrTypeInfo) InitAst_AbbrTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,rawTxt string) {
    self.InitAst_TypeInfo(_env, nil, processInfo)
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.rawTxt = rawTxt
    
}

// 5664: decl @lune.@base.@Ast.AbbrTypeInfo.isModule
func (self *Ast_AbbrTypeInfo) IsModule(_env *LnsEnv) bool {
    return false
}

// 5669: decl @lune.@base.@Ast.AbbrTypeInfo.getTxt
func (self *Ast_AbbrTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 5675: decl @lune.@base.@Ast.AbbrTypeInfo.getTxtWithRaw
func (self *Ast_AbbrTypeInfo) GetTxtWithRaw(_env *LnsEnv, rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return rawTxt
}

// 5682: decl @lune.@base.@Ast.AbbrTypeInfo.canEvalWith
func (self *Ast_AbbrTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    return false, nil
}

// 5689: decl @lune.@base.@Ast.AbbrTypeInfo.serialize
func (self *Ast_AbbrTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    Util_err(_env, "illegal call")
}

// 5693: decl @lune.@base.@Ast.AbbrTypeInfo.get_display_stirng_with
func (self *Ast_AbbrTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, raw, nil, nil, nil)
}

// 5697: decl @lune.@base.@Ast.AbbrTypeInfo.get_display_stirng
func (self *Ast_AbbrTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
}

// 5701: decl @lune.@base.@Ast.AbbrTypeInfo.getModule
func (self *Ast_AbbrTypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5704: decl @lune.@base.@Ast.AbbrTypeInfo.get_kind
func (self *Ast_AbbrTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Abbr
}

// 5707: decl @lune.@base.@Ast.AbbrTypeInfo.get_nilable
func (self *Ast_AbbrTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return true
}

// 5710: decl @lune.@base.@Ast.AbbrTypeInfo.get_nilableTypeInfo
func (self *Ast_AbbrTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5713: decl @lune.@base.@Ast.AbbrTypeInfo.get_mutMode
func (self *Ast_AbbrTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return Ast_MutMode__IMut
}

// 5716: decl @lune.@base.@Ast.AbbrTypeInfo.get_accessMode
func (self *Ast_AbbrTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return Ast_AccessMode__Local
}


// declaration Class -- ExtTypeInfo
type Ast_ExtTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type Ast_ExtTypeInfo struct {
    Ast_TypeInfo
    typeId *Ast_IdInfo
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
func NewAst_ExtTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo) *Ast_ExtTypeInfo {
    obj := &Ast_ExtTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_ExtTypeInfo(_env, arg1, arg2)
    return obj
}
func (self *Ast_ExtTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo{ return self.typeId }
func (self *Ast_ExtTypeInfo) Get_extedType(_env *LnsEnv) *Ast_TypeInfo{ return self.extedType }
func (self *Ast_ExtTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.nilableTypeInfo }
// advertise -- 5727
func (self *Ast_ExtTypeInfo) AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo) {
self.extedType. FP.AddChildren( _env, arg1)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap {
    return self.extedType. FP.CreateAlt2typeMap( _env, arg1)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.extedType. FP.GetFullName( _env, arg1,arg2,arg3)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) GetOverridingType(_env *LnsEnv) LnsAny {
    return self.extedType. FP.GetOverridingType( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.extedType. FP.GetParentFullName( _env, arg1,arg2,arg3)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.extedType. FP.GetParentId( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.extedType. FP.getProcessInfo( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_abstractFlag(_env *LnsEnv) bool {
    return self.extedType. FP.Get_abstractFlag( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return self.extedType. FP.Get_accessMode( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_argTypeInfoList(_env *LnsEnv) *LnsList {
    return self.extedType. FP.Get_argTypeInfoList( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_asyncMode(_env *LnsEnv) LnsInt {
    return self.extedType. FP.Get_asyncMode( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_autoFlag(_env *LnsEnv) bool {
    return self.extedType. FP.Get_autoFlag( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return self.extedType. FP.Get_baseId( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.extedType. FP.Get_baseTypeInfo( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_children(_env *LnsEnv) *LnsList {
    return self.extedType. FP.Get_children( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_externalFlag(_env *LnsEnv) bool {
    return self.extedType. FP.Get_externalFlag( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.extedType. FP.Get_genSrcTypeInfo( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList {
    return self.extedType. FP.Get_interfaceList( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList {
    return self.extedType. FP.Get_itemTypeInfoList( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return self.extedType. FP.Get_mutMode( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.extedType. FP.Get_parentInfo( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.extedType. FP.Get_processInfo( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return self.extedType. FP.Get_rawTxt( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList {
    return self.extedType. FP.Get_retTypeInfoList( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_scope(_env *LnsEnv) LnsAny {
    return self.extedType. FP.Get_scope( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_staticFlag(_env *LnsEnv) bool {
    return self.extedType. FP.Get_staticFlag( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) Get_typeData(_env *LnsEnv) *Ast_TypeData {
    return self.extedType. FP.Get_typeData( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) HasBase(_env *LnsEnv) bool {
    return self.extedType. FP.HasBase( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool {
    return self.extedType. FP.HasRouteNamespaceFrom( _env, arg1)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.extedType. FP.IsInheritFrom( _env, arg1,arg2,arg3)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) IsModule(_env *LnsEnv) bool {
    return self.extedType. FP.IsModule( _env)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo,arg2 string,arg3 *LnsList,arg4 LnsAny) string {
    return self.extedType. FP.SerializeTypeInfoList( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 5727
func (self *Ast_ExtTypeInfo) SwitchScope(_env *LnsEnv, arg1 *Ast_Scope) {
self.extedType. FP.SwitchScope( _env, arg1)
}
// 5733: DeclConstr
func (self *Ast_ExtTypeInfo) InitAst_ExtTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,extedType *Ast_TypeInfo) {
    self.InitAst_TypeInfo(_env, extedType.FP.Get_scope(_env), processInfo)
    self.typeId = processInfo.FP.newId(_env, &self.Ast_TypeInfo)
    
    self.extedType = extedType
    
    self.nilableTypeInfo = &NewAst_NilableTypeInfo(_env, processInfo, &self.Ast_TypeInfo).Ast_TypeInfo
    
}

// 5744: decl @lune.@base.@Ast.ExtTypeInfo.getTxt
func (self *Ast_ExtTypeInfo) GetTxt(_env *LnsEnv, typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, self.FP.Get_rawTxt(_env), typeNameCtrl, importInfo, localFlag)
}

// 5750: decl @lune.@base.@Ast.ExtTypeInfo.getTxtWithRaw
func (self *Ast_ExtTypeInfo) GetTxtWithRaw(_env *LnsEnv, rawTxt string,typeNameCtrl LnsAny,importInfo LnsAny,localFlag LnsAny) string {
    return _env.LuaVM.String_format("Luaval<%s>", []LnsAny{self.extedType.FP.GetTxtWithRaw(_env, rawTxt, typeNameCtrl, importInfo, localFlag)})
}

// 5758: decl @lune.@base.@Ast.ExtTypeInfo.equals
func (self *Ast_ExtTypeInfo) Equals(_env *LnsEnv, processInfo *Ast_ProcessInfo,typeInfo *Ast_TypeInfo,alt2type LnsAny,checkModifer LnsAny) bool {
    {
        _extTypeInfo := Ast_ExtTypeInfoDownCastF(typeInfo.FP)
        if !Lns_IsNil( _extTypeInfo ) {
            extTypeInfo := _extTypeInfo.(*Ast_ExtTypeInfo)
            return self.extedType.FP.Equals(_env, processInfo, extTypeInfo.extedType, alt2type, checkModifer)
        }
    }
    if Lns_isCondTrue( Lns_car(Ast_failCreateLuavalWith_3312_(_env, self.extedType, Ast_LuavalConvKind__InLua, false))){
        return false
    }
    return self.extedType.FP.Equals(_env, processInfo, typeInfo, alt2type, checkModifer)
}

// 5774: decl @lune.@base.@Ast.ExtTypeInfo.canEvalWith
func (self *Ast_ExtTypeInfo) CanEvalWith(_env *LnsEnv, processInfo *Ast_ProcessInfo,other *Ast_TypeInfo,canEvalType LnsInt,alt2type *LnsMap)(bool, LnsAny) {
    {
        _extTypeInfo := Ast_ExtTypeInfoDownCastF(other.FP.Get_nonnilableType(_env).FP)
        if !Lns_IsNil( _extTypeInfo ) {
            extTypeInfo := _extTypeInfo.(*Ast_ExtTypeInfo)
            var otherExtedType *Ast_TypeInfo
            if other.FP.Get_nilable(_env){
                otherExtedType = extTypeInfo.extedType.FP.Get_nilableTypeInfo(_env)
                
            } else { 
                otherExtedType = extTypeInfo.extedType
                
            }
            return self.extedType.FP.CanEvalWith(_env, processInfo, otherExtedType, canEvalType, alt2type)
        }
    }
    {
        __exp := Ast_convExp22961(Lns_2DDD(Ast_failCreateLuavalWith_3312_(_env, other, Ast_LuavalConvKind__ToLua, true)))
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(string)
            return false, _exp
        }
    }
    return true, nil
}

// 5798: decl @lune.@base.@Ast.ExtTypeInfo.serialize
func (self *Ast_ExtTypeInfo) Serialize(_env *LnsEnv, stream Lns_oStream,serializeInfo *Ast_SerializeInfo) {
    stream.Write(_env, _env.LuaVM.String_format("{ skind = %d, typeId = %d, extedTypeId = %s }\n", []LnsAny{Ast_SerializeKind__Ext, self.typeId.Id, serializeInfo.FP.SerializeId(_env, self.extedType.FP.Get_typeId(_env))}))
}

// 5805: decl @lune.@base.@Ast.ExtTypeInfo.get_display_stirng_with
func (self *Ast_ExtTypeInfo) Get_display_stirng_with(_env *LnsEnv, raw string,alt2type LnsAny) string {
    return self.FP.GetTxtWithRaw(_env, raw, nil, nil, nil)
}

// 5809: decl @lune.@base.@Ast.ExtTypeInfo.get_display_stirng
func (self *Ast_ExtTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.FP.Get_display_stirng_with(_env, self.FP.Get_rawTxt(_env), nil)
}

// 5813: decl @lune.@base.@Ast.ExtTypeInfo.getModule
func (self *Ast_ExtTypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return Ast_headTypeInfo
}

// 5816: decl @lune.@base.@Ast.ExtTypeInfo.get_kind
func (self *Ast_ExtTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return Ast_TypeInfoKind__Ext
}

// 5819: decl @lune.@base.@Ast.ExtTypeInfo.get_aliasSrc
func (self *Ast_ExtTypeInfo) Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5822: decl @lune.@base.@Ast.ExtTypeInfo.get_srcTypeInfo
func (self *Ast_ExtTypeInfo) Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5825: decl @lune.@base.@Ast.ExtTypeInfo.get_nonnilableType
func (self *Ast_ExtTypeInfo) Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo {
    return &self.Ast_TypeInfo
}

// 5828: decl @lune.@base.@Ast.ExtTypeInfo.get_nilable
func (self *Ast_ExtTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return false
}

// 5831: decl @lune.@base.@Ast.ExtTypeInfo.applyGeneric
func (self *Ast_ExtTypeInfo) ApplyGeneric(_env *LnsEnv, alt2typeMap *LnsMap,moduleTypeInfo *Ast_TypeInfo) LnsAny {
    var typeInfo LnsAny
    typeInfo = self.extedType.FP.ApplyGeneric(_env, alt2typeMap, moduleTypeInfo)
    if typeInfo != self.extedType{
        Util_err(_env, _env.LuaVM.String_format("not support generics -- %s", []LnsAny{self.extedType.FP.GetTxt(_env, nil, nil, nil)}))
    }
    return &self.Ast_TypeInfo
}


// declaration Class -- AndExpTypeInfo
type Ast_AndExpTypeInfoMtd interface {
    AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo)
    ApplyGeneric(_env *LnsEnv, arg1 *LnsMap, arg2 *Ast_TypeInfo) LnsAny
    CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *LnsMap)(bool, LnsAny)
    CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap
    Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 LnsAny) bool
    GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager, arg3 LnsAny) string
    GetModule(_env *LnsEnv) *Ast_TypeInfo
    GetOverridingType(_env *LnsEnv) LnsAny
    GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl, arg2 LnsAny, arg3 LnsAny) string
    GetParentId(_env *LnsEnv) *Ast_IdInfo
    getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo
    GetTxt(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 LnsAny) string
    GetTxtWithRaw(_env *LnsEnv, arg1 string, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny) string
    Get_abstractFlag(_env *LnsEnv) bool
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo
    Get_argTypeInfoList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsInt
    Get_autoFlag(_env *LnsEnv) bool
    Get_baseId(_env *LnsEnv) *Ast_IdInfo
    Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_children(_env *LnsEnv) *LnsList
    Get_display_stirng(_env *LnsEnv) string
    Get_display_stirng_with(_env *LnsEnv, arg1 string, arg2 LnsAny) string
    Get_exp1(_env *LnsEnv) *Ast_TypeInfo
    Get_exp2(_env *LnsEnv) *Ast_TypeInfo
    Get_extedType(_env *LnsEnv) *Ast_TypeInfo
    Get_externalFlag(_env *LnsEnv) bool
    Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_interfaceList(_env *LnsEnv) *LnsList
    Get_itemTypeInfoList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_mutMode(_env *LnsEnv) LnsInt
    Get_nilable(_env *LnsEnv) bool
    Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo
    Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_rawTxt(_env *LnsEnv) string
    Get_result(_env *LnsEnv) *Ast_TypeInfo
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_scope(_env *LnsEnv) LnsAny
    Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_staticFlag(_env *LnsEnv) bool
    Get_typeData(_env *LnsEnv) *Ast_TypeData
    Get_typeId(_env *LnsEnv) *Ast_IdInfo
    HasBase(_env *LnsEnv) bool
    HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 LnsAny) bool
    IsModule(_env *LnsEnv) bool
    Serialize(_env *LnsEnv, arg1 Lns_oStream, arg2 *Ast_SerializeInfo)
    SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    SwitchScope(_env *LnsEnv, arg1 *Ast_Scope)
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
func NewAst_AndExpTypeInfo(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 *Ast_TypeInfo, arg4 *Ast_TypeInfo) *Ast_AndExpTypeInfo {
    obj := &Ast_AndExpTypeInfo{}
    obj.FP = obj
    obj.Ast_TypeInfo.FP = obj
    obj.InitAst_AndExpTypeInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *Ast_AndExpTypeInfo) Get_exp1(_env *LnsEnv) *Ast_TypeInfo{ return self.exp1 }
func (self *Ast_AndExpTypeInfo) Get_exp2(_env *LnsEnv) *Ast_TypeInfo{ return self.exp2 }
func (self *Ast_AndExpTypeInfo) Get_result(_env *LnsEnv) *Ast_TypeInfo{ return self.result }
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) AddChildren(_env *LnsEnv, arg1 *Ast_TypeInfo) {
self.result. FP.AddChildren( _env, arg1)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) ApplyGeneric(_env *LnsEnv, arg1 *LnsMap,arg2 *Ast_TypeInfo) LnsAny {
    return self.result. FP.ApplyGeneric( _env, arg1,arg2)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) CanEvalWith(_env *LnsEnv, arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsInt,arg4 *LnsMap)(bool, LnsAny) {
    return self.result. FP.CanEvalWith( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) CreateAlt2typeMap(_env *LnsEnv, arg1 bool) *LnsMap {
    return self.result. FP.CreateAlt2typeMap( _env, arg1)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Equals(_env *LnsEnv, arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny,arg4 LnsAny) bool {
    return self.result. FP.Equals( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) GetFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 Ast_ModuleInfoManager,arg3 LnsAny) string {
    return self.result. FP.GetFullName( _env, arg1,arg2,arg3)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) GetModule(_env *LnsEnv) *Ast_TypeInfo {
    return self.result. FP.GetModule( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) GetOverridingType(_env *LnsEnv) LnsAny {
    return self.result. FP.GetOverridingType( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) GetParentFullName(_env *LnsEnv, arg1 *Ast_TypeNameCtrl,arg2 LnsAny,arg3 LnsAny) string {
    return self.result. FP.GetParentFullName( _env, arg1,arg2,arg3)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) GetParentId(_env *LnsEnv) *Ast_IdInfo {
    return self.result. FP.GetParentId( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) getProcessInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.result. FP.getProcessInfo( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) GetTxt(_env *LnsEnv, arg1 LnsAny,arg2 LnsAny,arg3 LnsAny) string {
    return self.result. FP.GetTxt( _env, arg1,arg2,arg3)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) GetTxtWithRaw(_env *LnsEnv, arg1 string,arg2 LnsAny,arg3 LnsAny,arg4 LnsAny) string {
    return self.result. FP.GetTxtWithRaw( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_abstractFlag(_env *LnsEnv) bool {
    return self.result. FP.Get_abstractFlag( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_accessMode(_env *LnsEnv) LnsInt {
    return self.result. FP.Get_accessMode( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_aliasSrc(_env *LnsEnv) *Ast_TypeInfo {
    return self.result. FP.Get_aliasSrc( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_argTypeInfoList(_env *LnsEnv) *LnsList {
    return self.result. FP.Get_argTypeInfoList( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_asyncMode(_env *LnsEnv) LnsInt {
    return self.result. FP.Get_asyncMode( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_autoFlag(_env *LnsEnv) bool {
    return self.result. FP.Get_autoFlag( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_baseId(_env *LnsEnv) *Ast_IdInfo {
    return self.result. FP.Get_baseId( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_baseTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.result. FP.Get_baseTypeInfo( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_children(_env *LnsEnv) *LnsList {
    return self.result. FP.Get_children( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_display_stirng(_env *LnsEnv) string {
    return self.result. FP.Get_display_stirng( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_display_stirng_with(_env *LnsEnv, arg1 string,arg2 LnsAny) string {
    return self.result. FP.Get_display_stirng_with( _env, arg1,arg2)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_extedType(_env *LnsEnv) *Ast_TypeInfo {
    return self.result. FP.Get_extedType( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_externalFlag(_env *LnsEnv) bool {
    return self.result. FP.Get_externalFlag( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_genSrcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.result. FP.Get_genSrcTypeInfo( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_interfaceList(_env *LnsEnv) *LnsList {
    return self.result. FP.Get_interfaceList( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_itemTypeInfoList(_env *LnsEnv) *LnsList {
    return self.result. FP.Get_itemTypeInfoList( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_kind(_env *LnsEnv) LnsInt {
    return self.result. FP.Get_kind( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_mutMode(_env *LnsEnv) LnsInt {
    return self.result. FP.Get_mutMode( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_nilable(_env *LnsEnv) bool {
    return self.result. FP.Get_nilable( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_nilableTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.result. FP.Get_nilableTypeInfo( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_nonnilableType(_env *LnsEnv) *Ast_TypeInfo {
    return self.result. FP.Get_nonnilableType( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_parentInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.result. FP.Get_parentInfo( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo {
    return self.result. FP.Get_processInfo( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_rawTxt(_env *LnsEnv) string {
    return self.result. FP.Get_rawTxt( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList {
    return self.result. FP.Get_retTypeInfoList( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_scope(_env *LnsEnv) LnsAny {
    return self.result. FP.Get_scope( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_srcTypeInfo(_env *LnsEnv) *Ast_TypeInfo {
    return self.result. FP.Get_srcTypeInfo( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_staticFlag(_env *LnsEnv) bool {
    return self.result. FP.Get_staticFlag( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_typeData(_env *LnsEnv) *Ast_TypeData {
    return self.result. FP.Get_typeData( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Get_typeId(_env *LnsEnv) *Ast_IdInfo {
    return self.result. FP.Get_typeId( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) HasBase(_env *LnsEnv) bool {
    return self.result. FP.HasBase( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) HasRouteNamespaceFrom(_env *LnsEnv, arg1 *Ast_TypeInfo) bool {
    return self.result. FP.HasRouteNamespaceFrom( _env, arg1)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) IsInheritFrom(_env *LnsEnv, arg1 *Ast_ProcessInfo,arg2 *Ast_TypeInfo,arg3 LnsAny) bool {
    return self.result. FP.IsInheritFrom( _env, arg1,arg2,arg3)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) IsModule(_env *LnsEnv) bool {
    return self.result. FP.IsModule( _env)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) Serialize(_env *LnsEnv, arg1 Lns_oStream,arg2 *Ast_SerializeInfo) {
self.result. FP.Serialize( _env, arg1,arg2)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) SerializeTypeInfoList(_env *LnsEnv, arg1 *Ast_SerializeInfo,arg2 string,arg3 *LnsList,arg4 LnsAny) string {
    return self.result. FP.SerializeTypeInfoList( _env, arg1,arg2,arg3,arg4)
}
// advertise -- 5944
func (self *Ast_AndExpTypeInfo) SwitchScope(_env *LnsEnv, arg1 *Ast_Scope) {
self.result. FP.SwitchScope( _env, arg1)
}
// 5949: DeclConstr
func (self *Ast_AndExpTypeInfo) InitAst_AndExpTypeInfo(_env *LnsEnv, processInfo *Ast_ProcessInfo,exp1 *Ast_TypeInfo,exp2 *Ast_TypeInfo,result *Ast_TypeInfo) {
    self.InitAst_TypeInfo(_env, result.FP.Get_scope(_env), processInfo)
    self.exp1 = exp1
    
    self.exp2 = exp2
    
    self.result = result
    
}


// declaration Class -- RefTypeInfo
type Ast_RefTypeInfoMtd interface {
    Get_itemRefTypeList(_env *LnsEnv) *LnsList
    Get_pos(_env *LnsEnv) *Types_Position
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
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
func NewAst_RefTypeInfo(_env *LnsEnv, arg1 *Types_Position, arg2 *LnsList, arg3 *Ast_TypeInfo) *Ast_RefTypeInfo {
    obj := &Ast_RefTypeInfo{}
    obj.FP = obj
    obj.InitAst_RefTypeInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Ast_RefTypeInfo) InitAst_RefTypeInfo(_env *LnsEnv, arg1 *Types_Position, arg2 *LnsList, arg3 *Ast_TypeInfo) {
    self.pos = arg1
    self.itemRefTypeList = arg2
    self.typeInfo = arg3
}
func (self *Ast_RefTypeInfo) Get_pos(_env *LnsEnv) *Types_Position{ return self.pos }
func (self *Ast_RefTypeInfo) Get_itemRefTypeList(_env *LnsEnv) *LnsList{ return self.itemRefTypeList }
func (self *Ast_RefTypeInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }

// declaration Class -- TypeAnalyzer
type Ast_TypeAnalyzerMtd interface {
    AnalyzeType(_env *LnsEnv, arg1 *Ast_Scope, arg2 Parser_PushbackParser, arg3 LnsInt, arg4 bool, arg5 bool)(LnsAny, LnsAny, LnsAny)
    AnalyzeTypeFromTxt(_env *LnsEnv, arg1 string, arg2 *Ast_Scope, arg3 LnsInt, arg4 bool)(LnsAny, LnsAny, LnsAny)
    AnalyzeTypeItemList(_env *LnsEnv, arg1 bool, arg2 bool, arg3 bool, arg4 *Ast_TypeInfo, arg5 *Types_Position)(LnsAny, LnsAny, LnsAny)
    analyzeTypeSub(_env *LnsEnv, arg1 bool)(LnsAny, LnsAny, LnsAny)
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
func NewAst_TypeAnalyzer(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Ast_TypeInfo, arg3 *Ast_TypeInfo, arg4 *Ast_Scope, arg5 LnsInt, arg6 bool) *Ast_TypeAnalyzer {
    obj := &Ast_TypeAnalyzer{}
    obj.FP = obj
    obj.InitAst_TypeAnalyzer(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
// 7145: DeclConstr
func (self *Ast_TypeAnalyzer) InitAst_TypeAnalyzer(_env *LnsEnv, processInfo *Ast_ProcessInfo,parentInfo *Ast_TypeInfo,moduleType *Ast_TypeInfo,moduleScope *Ast_Scope,scopeAccess LnsInt,validMutControl bool) {
    self.processInfo = processInfo
    
    self.parentInfo = parentInfo
    
    self.moduleType = moduleType
    
    self.moduleScope = moduleScope
    
    self.scopeAccess = scopeAccess
    
    self.validMutControl = validMutControl
    
    self.scope = Ast_rootScope
    
    self.accessMode = Ast_AccessMode__Local
    
    self.parentPub = false
    
    self.parser = NewParser_DefaultPushbackParser(_env, &NewParser_DummyParser(_env).Parser_Parser).FP
    
}



// 7170: decl @lune.@base.@Ast.TypeAnalyzer.analyzeType
func (self *Ast_TypeAnalyzer) AnalyzeType(_env *LnsEnv, scope *Ast_Scope,parser Parser_PushbackParser,accessMode LnsInt,allowDDD bool,parentPub bool)(LnsAny, LnsAny, LnsAny) {
    self.scope = scope
    
    self.parser = parser
    
    self.accessMode = accessMode
    
    self.parentPub = parentPub
    
    return self.FP.analyzeTypeSub(_env, allowDDD)
}

// 7182: decl @lune.@base.@Ast.TypeAnalyzer.analyzeTypeFromTxt
func (self *Ast_TypeAnalyzer) AnalyzeTypeFromTxt(_env *LnsEnv, txt string,scope *Ast_Scope,accessMode LnsInt,parentPub bool)(LnsAny, LnsAny, LnsAny) {
    var stream *Parser_TxtStream
    stream = NewParser_TxtStream(_env, txt)
    var parser *Parser_DefaultPushbackParser
    parser = NewParser_DefaultPushbackParser(_env, &NewParser_StreamParser(_env, stream.FP, "test", nil, nil).Parser_Parser)
    return self.FP.AnalyzeType(_env, scope, parser.FP, accessMode, true, parentPub)
}

// 7193: decl @lune.@base.@Ast.TypeAnalyzer.analyzeTypeSub
func (self *Ast_TypeAnalyzer) analyzeTypeSub(_env *LnsEnv, allowDDD bool)(LnsAny, LnsAny, LnsAny) {
    var firstToken *Types_Token
    firstToken = self.parser.GetTokenNoErr(_env)
    var token *Types_Token
    token = firstToken
    var refFlag bool
    refFlag = false
    if token.Txt == "&"{
        refFlag = true
        
        token = self.parser.GetTokenNoErr(_env)
        
    }
    var mutFlag bool
    mutFlag = false
    if token.Txt == "mut"{
        mutFlag = true
        
        token = self.parser.GetTokenNoErr(_env)
        
    }
    var typeInfo *Ast_TypeInfo
    if token.Txt == "..."{
        typeInfo = Ast_builtinTypeDDD
        
    } else { 
        var symbol *Ast_SymbolInfo
        
        {
            _symbol := self.scope.FP.GetSymbolTypeInfo(_env, token.Txt, self.scope, self.moduleScope, self.scopeAccess)
            if _symbol == nil{
                return nil, token.Pos, "not found type -- " + token.Txt
            } else {
                symbol = _symbol.(*Ast_SymbolInfo)
            }
        }
        if symbol.FP.Get_kind(_env) != Ast_SymbolKind__Typ{
            return nil, token.Pos, _env.LuaVM.String_format("illegal type -- %s", []LnsAny{symbol.FP.Get_name(_env)})
        }
        typeInfo = symbol.FP.Get_typeInfo(_env)
        
    }
    return self.FP.AnalyzeTypeItemList(_env, allowDDD, refFlag, mutFlag, typeInfo, token.Pos)
}

// 7233: decl @lune.@base.@Ast.TypeAnalyzer.analyzeTypeItemList
func (self *Ast_TypeAnalyzer) AnalyzeTypeItemList(_env *LnsEnv, allowDDD bool,refFlag bool,mutFlag bool,typeInfo *Ast_TypeInfo,pos *Types_Position)(LnsAny, LnsAny, LnsAny) {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.parentPub) &&
        _env.SetStackVal( Ast_isPubToExternal(_env, self.accessMode)) &&
        _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, typeInfo.FP.Get_accessMode(_env)))) ).(bool)){
        return nil, pos, _env.LuaVM.String_format("This type must be public. -- %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)})
    }
    var token *Types_Token
    token = self.parser.GetTokenNoErr(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Consecutive) &&
        _env.SetStackVal( token.Txt == "!") ).(bool)){
        typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
        
        token = self.parser.GetTokenNoErr(_env)
        
    }
    var genericRefList *LnsList
    genericRefList = NewLnsList([]LnsAny{})
    for  {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Txt == "[") ||
            _env.SetStackVal( token.Txt == "[@") ).(bool){
            if token.Txt == "["{
                typeInfo = self.processInfo.FP.CreateList(_env, self.accessMode, self.parentInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), Ast_MutMode__Mut)
                
            } else { 
                typeInfo = self.processInfo.FP.CreateArray(_env, self.accessMode, self.parentInfo, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), Ast_MutMode__Mut)
                
            }
            token = self.parser.GetTokenNoErr(_env)
            
            if token.Txt != "]"{
                return nil, token.Pos, "not found -- ']'"
            }
            
        } else if token.Txt == "<"{
            var genericList *LnsList
            genericList = NewLnsList([]LnsAny{})
            var nextToken *Types_Token
            nextToken = Parser_getEofToken(_env)
            for {
                var refType LnsAny
                refType = Ast_convExp30696(Lns_2DDD(self.FP.analyzeTypeSub(_env, false)))
                if refType != nil{
                    refType_5336 := refType.(*Ast_RefTypeInfo)
                    genericRefList.Insert(Ast_RefTypeInfo2Stem(refType_5336))
                    genericList.Insert(Ast_TypeInfo2Stem(refType_5336.FP.Get_typeInfo(_env)))
                }
                nextToken = self.parser.GetTokenNoErr(_env)
                
                if nextToken.Txt != ","{ break }
            }
            if nextToken.Txt != ">"{
                return nil, nextToken.Pos, "not found -- ']'"
            }
            
            if _switch31317 := typeInfo.FP.Get_kind(_env); _switch31317 == Ast_TypeInfoKind__Map {
                if genericList.Len() != 2{
                    return nil, pos, "Key or value type is unknown"
                } else { 
                    typeInfo = self.processInfo.FP.CreateMap(_env, self.accessMode, self.parentInfo, genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), genericList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__Mut)
                    
                }
            } else if _switch31317 == Ast_TypeInfoKind__List {
                if genericList.Len() != 1{
                    return nil, pos, _env.LuaVM.String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = self.processInfo.FP.CreateList(_env, self.accessMode, self.parentInfo, genericList, Ast_MutMode__Mut)
                
            } else if _switch31317 == Ast_TypeInfoKind__Array {
                if genericList.Len() != 1{
                    return nil, pos, _env.LuaVM.String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = self.processInfo.FP.CreateArray(_env, self.accessMode, self.parentInfo, genericList, Ast_MutMode__Mut)
                
            } else if _switch31317 == Ast_TypeInfoKind__Set {
                if genericList.Len() != 1{
                    return nil, pos, _env.LuaVM.String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = self.processInfo.FP.CreateSet(_env, self.accessMode, self.parentInfo, genericList, Ast_MutMode__Mut)
                
            } else if _switch31317 == Ast_TypeInfoKind__DDD {
                if genericList.Len() != 1{
                    return nil, pos, _env.LuaVM.String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = &self.processInfo.FP.CreateDDD(_env, genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), false, false).Ast_TypeInfo
                
            } else if _switch31317 == Ast_TypeInfoKind__Class || _switch31317 == Ast_TypeInfoKind__IF {
                if genericList.Len() != typeInfo.FP.Get_itemTypeInfoList(_env).Len(){
                    return nil, pos, _env.LuaVM.String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                for _, _itemType := range( genericList.Items ) {
                    itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if itemType.FP.Get_nilable(_env){
                        var mess string
                        mess = _env.LuaVM.String_format("can't use nilable type -- %s", []LnsAny{itemType.FP.GetTxt(_env, nil, nil, nil)})
                        return nil, pos, mess
                    }
                }
                typeInfo = &self.processInfo.FP.CreateGeneric(_env, typeInfo, genericList, self.moduleType).Ast_TypeInfo
                
            } else if _switch31317 == Ast_TypeInfoKind__Box {
                if genericList.Len() != 1{
                    return nil, pos, _env.LuaVM.String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                typeInfo = self.processInfo.FP.CreateBox(_env, self.accessMode, genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                
            } else if _switch31317 == Ast_TypeInfoKind__Ext {
                if genericList.Len() != 1{
                    return nil, pos, _env.LuaVM.String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()})
                }
                
                switch _exp31295 := self.processInfo.FP.CreateLuaval(_env, genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), true).(type) {
                case *Ast_LuavalResult__OK:
                extTypeInfo := _exp31295.Val1
                    typeInfo = extTypeInfo
                    
                case *Ast_LuavalResult__Err:
                err := _exp31295.Val1
                    return nil, pos, err
                }
            } else {
                return nil, pos, _env.LuaVM.String_format("not support generic: %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)})
            }
        } else { 
            self.parser.Pushback(_env)
            break
        }
        token = self.parser.GetTokenNoErr(_env)
        
    }
    if token.Txt == "!"{
        typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
        
        self.parser.GetTokenNoErr(_env)
    }
    if Lns_op_not(allowDDD){
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            return nil, pos, _env.LuaVM.String_format("invalid type. -- '%s'", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)})
        }
    }
    if refFlag{
        if self.validMutControl{
            typeInfo = self.processInfo.FP.CreateModifier(_env, typeInfo, Ast_MutMode__IMut)
            
        }
    }
    return NewAst_RefTypeInfo(_env, pos, genericRefList, typeInfo), nil, nil
}


func Lns_Ast_init(_env *LnsEnv) {
    if init_Ast { return }
    init_Ast = true
    Ast__mod__ = "@lune.@base.@Ast"
    Lns_InitMod()
    Lns_Parser_init(_env)
    Lns_Util_init(_env)
    Lns_Code_init(_env)
    Lns_Log_init(_env)
    Lns_Types_init(_env)
    Ast_extStartId = 100000
    Ast_extMaxId = 10000000
    Ast_userStartId = 1000
    Ast_rootProcessInfo = Ast_ProcessInfo_createRoot_1076_(_env)
    Ast_rootProcessInfoRo = Ast_rootProcessInfo
    Ast_rootTypeIdInfo = Ast_rootProcessInfo.FP.newIdForRoot(_env)
    Ast_rootTypeId = Ast_rootTypeIdInfo.Id
    Ast_DummyModuleInfoManager____init_1155_(_env)
    Ast_sym2builtInTypeMapWork = NewLnsMap( map[LnsAny]LnsAny{})
    Ast_sym2builtInTypeMap = Ast_sym2builtInTypeMapWork
    Ast_builtInTypeIdSetWork = NewLnsMap( map[LnsAny]LnsAny{})
    Ast_builtInTypeIdSet = Ast_builtInTypeIdSetWork
    Ast_txt2AccessModeMap = NewLnsMap( map[LnsAny]LnsAny{"none":Ast_AccessMode__None,"pub":Ast_AccessMode__Pub,"pro":Ast_AccessMode__Pro,"pri":Ast_AccessMode__Pri,"local":Ast_AccessMode__Local,"_direct":Ast_AccessMode__Direct,"global":Ast_AccessMode__Global,})
    Ast_accessMode2txtMap = NewLnsMap( map[LnsAny]LnsAny{Ast_AccessMode__None:"none",Ast_AccessMode__Pub:"pub",Ast_AccessMode__Pro:"pro",Ast_AccessMode__Pri:"pri",Ast_AccessMode__Local:"local",Ast_AccessMode__Direct:"_direct",Ast_AccessMode__Global:"global",})
    Ast_rootScope = NewAst_Scope(_env, Ast_rootProcessInfo, nil, false, nil, nil)
    Ast_rootScopeRo = Ast_rootScope
    Ast_dummyList = NewLnsList([]LnsAny{})
    Ast_headTypeInfo = &Ast_RootTypeInfo_create_1746_(_env).Ast_TypeInfo
    Ast_rootProcessInfo.FP.setRootTypeInfo(_env, Ast_rootTypeId, Ast_headTypeInfo)
    Ast_defaultTypeNameCtrl = NewAst_TypeNameCtrl(_env, Ast_headTypeInfo)
    Ast_AutoBoxingInfo____init_1834_(_env)
    Ast_dummyIdInfo = NewAst_IdInfo(_env, 1, Ast_rootProcessInfo)
    Ast_CanEvalCtrlTypeInfo____init_1856_(_env)
    Ast_dummySymbol = Lns_unwrap( Lns_car(Ast_rootScope.FP.AddLocalVar(_env, Ast_rootProcessInfo, false, false, "$$", nil, Ast_headTypeInfo, Ast_MutMode__IMut))).(*Ast_SymbolInfo)
    Ast_boxRootAltType = NewAst_AlternateTypeInfo(_env, Ast_rootProcessInfo, true, 1, "_T", Ast_AccessMode__Pub, Ast_headTypeInfo, nil, nil)
    Ast_boxRootScope = NewAst_Scope(_env, Ast_rootProcessInfo, Ast_rootScope, true, nil, nil)
    Ast_rootProcessInfo.FP.set_typeInfo2Map(_env, NewAst_TypeInfo2Map(_env))
    Ast_immutableTypeSetWork = NewLnsSet([]LnsAny{})
    Ast_immutableTypeSet = Ast_immutableTypeSetWork
    Ast_addBuiltin(_env, Ast_headTypeInfo)
    Ast_builtinTypeNone = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "__None", "", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeEmpty = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "__Empty", "::", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeNeverRet = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "__NRet", "__", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeStem = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Stem", "stem", Ast_TypeInfoKind__Stem, nil, nil)
    Ast_builtinTypeStem_ = Ast_builtinTypeStem.FP.Get_nilableTypeInfo(_env)
    Ast_builtinTypeBool = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Bool", "bool", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeInt = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Int", "int", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeReal = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Real", "real", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeChar = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "char", "__char", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeMapping = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Mapping", "Mapping", Ast_TypeInfoKind__IF, nil, nil)
    Ast_builtinTypeRunner = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "__Runner", "__Runner", Ast_TypeInfoKind__IF, nil, nil)
    Ast_builtinTypeAsyncItem = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "__AsyncItem", "__AsyncItem", Ast_TypeInfoKind__IF, nil, nil)
    Ast_builtinTypeString = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "String", "str", Ast_TypeInfoKind__Class, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeMapping)}))
    Ast_builtinTypeMap = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Map", "Map", Ast_TypeInfoKind__Map, nil, nil)
    Ast_builtinTypeSet = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Set", "Set", Ast_TypeInfoKind__Set, nil, nil)
    Ast_builtinTypeList = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "List", "List", Ast_TypeInfoKind__List, nil, nil)
    Ast_builtinTypeArray = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Array", "Array", Ast_TypeInfoKind__Array, nil, nil)
    Ast_immutableTypeSetWork.Add(Ast_TypeInfo2Stem(Ast_builtinTypeBool))
    Ast_immutableTypeSetWork.Add(Ast_TypeInfo2Stem(Ast_builtinTypeInt))
    Ast_immutableTypeSetWork.Add(Ast_TypeInfo2Stem(Ast_builtinTypeReal))
    Ast_immutableTypeSetWork.Add(Ast_TypeInfo2Stem(Ast_builtinTypeChar))
    Ast_immutableTypeSetWork.Add(Ast_TypeInfo2Stem(Ast_builtinTypeString))
    Ast_builtinTypeBox = NewAst_BoxTypeInfo(_env, Ast_rootProcessInfo, Ast_AccessMode__Pub, &Ast_boxRootAltType.Ast_TypeInfo)
    Ast_registBuiltin_3265_(_env, "Nilable", "Nilable", Ast_TypeInfoKind__Box, &Ast_builtinTypeBox.Ast_TypeInfo, Ast_headTypeInfo, true)
    Ast_builtinTypeNil = &NewAst_NilTypeInfo(_env, Ast_rootProcessInfo).Ast_TypeInfo
    Ast_registBuiltin_3265_(_env, "Nil", "nil", Ast_TypeInfoKind__Prim, Ast_builtinTypeNil, Ast_headTypeInfo, false)
    Ast_builtinTypeDDD = &Ast_rootProcessInfo.FP.CreateDDD(_env, Ast_builtinTypeStem_, true, false).Ast_TypeInfo
    Ast_registBuiltin_3265_(_env, "DDD", "...", Ast_TypeInfoKind__DDD, Ast_builtinTypeDDD, Ast_headTypeInfo, false)
    Ast_builtinTypeForm = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Form", "form", Ast_TypeInfoKind__Form, Ast_builtinTypeDDD, nil)
    Ast_immutableTypeSetWork.Add(Ast_TypeInfo2Stem(Ast_builtinTypeForm))
    Ast_builtinTypeSymbol = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Symbol", "sym", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeStat = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Stat", "stat", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeExp = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Exp", "__exp", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeMultiExp = Ast_NormalTypeInfo_createBuiltin_3270_(_env, "Exps", "__exps", Ast_TypeInfoKind__Prim, nil, nil)
    Ast_builtinTypeAbbr = &NewAst_AbbrTypeInfo(_env, Ast_rootProcessInfo, "##").Ast_TypeInfo
    Ast_builtinTypeAbbrNone = &NewAst_AbbrTypeInfo(_env, Ast_rootProcessInfo, "[##]").Ast_TypeInfo
    switch _exp23519 := Ast_rootProcessInfo.FP.CreateLuaval(_env, Ast_builtinTypeStem, true).(type) {
    case *Ast_LuavalResult__OK:
    typeInfo := _exp23519.Val1
        Ast_builtinTypeLua = typeInfo
        
    default:
        Util_err(_env, "illegal")
    }
    Ast_registBuiltin_3265_(_env, "Luaval", "Luaval", Ast_TypeInfoKind__Ext, Ast_builtinTypeLua, Ast_headTypeInfo, false)
    Ast_builtinTypeDDDLua = &Ast_rootProcessInfo.FP.CreateDDD(_env, Ast_builtinTypeStem_, true, true).Ast_TypeInfo
    Ast_registBuiltin_3265_(_env, "__LuaDDD", "__LuaDDD", Ast_TypeInfoKind__Ext, Ast_builtinTypeDDDLua, Ast_headTypeInfo, false)
    Ast_numberTypeSet = NewLnsSet([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt), Ast_TypeInfo2Stem(Ast_builtinTypeChar), Ast_TypeInfo2Stem(Ast_builtinTypeReal)})
    Ast_builtinTypeInfo2Map = Ast_rootProcessInfo.FP.get_typeInfo2Map(_env).FP.Clone(_env)
    Ast_bitBinOpMap = NewLnsMap( map[LnsAny]LnsAny{"&":Ast_BitOpKind__And,"|":Ast_BitOpKind__Or,"~":Ast_BitOpKind__Xor,"|>>":Ast_BitOpKind__RShift,"|<<":Ast_BitOpKind__LShift,})
    Ast_compOpSet = NewLnsSet([]LnsAny{"==", "~="})
    Ast_mathCompOpSet = NewLnsSet([]LnsAny{"<", "<=", ">", ">="})
}
func init() {
    init_Ast = false
}
