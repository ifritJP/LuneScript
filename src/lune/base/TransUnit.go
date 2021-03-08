// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_TransUnit bool
var TransUnit__mod__ string
// decl enum -- DeclClassMode 
type TransUnit_DeclClassMode = LnsInt
const TransUnit_DeclClassMode__Class = 0
const TransUnit_DeclClassMode__Interface = 1
const TransUnit_DeclClassMode__LazyModule = 3
const TransUnit_DeclClassMode__Module = 2
var TransUnit_DeclClassModeList_ = NewLnsList( []LnsAny {
  TransUnit_DeclClassMode__Class,
  TransUnit_DeclClassMode__Interface,
  TransUnit_DeclClassMode__Module,
  TransUnit_DeclClassMode__LazyModule,
})
func TransUnit_DeclClassMode_get__allList_1019_() *LnsList{
    return TransUnit_DeclClassModeList_
}
var TransUnit_DeclClassModeMap_ = map[LnsInt]string {
  TransUnit_DeclClassMode__Class: "DeclClassMode.Class",
  TransUnit_DeclClassMode__Interface: "DeclClassMode.Interface",
  TransUnit_DeclClassMode__LazyModule: "DeclClassMode.LazyModule",
  TransUnit_DeclClassMode__Module: "DeclClassMode.Module",
}
func TransUnit_DeclClassMode__from_1012_(arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_DeclClassModeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_DeclClassMode_getTxt(arg1 LnsInt) string {
    return TransUnit_DeclClassModeMap_[arg1];
}
// decl enum -- DeclFuncMode 
type TransUnit_DeclFuncMode = LnsInt
const TransUnit_DeclFuncMode__Class = 1
const TransUnit_DeclFuncMode__Func = 0
const TransUnit_DeclFuncMode__Glue = 3
const TransUnit_DeclFuncMode__Module = 2
var TransUnit_DeclFuncModeList_ = NewLnsList( []LnsAny {
  TransUnit_DeclFuncMode__Func,
  TransUnit_DeclFuncMode__Class,
  TransUnit_DeclFuncMode__Module,
  TransUnit_DeclFuncMode__Glue,
})
func TransUnit_DeclFuncMode_get__allList_1034_() *LnsList{
    return TransUnit_DeclFuncModeList_
}
var TransUnit_DeclFuncModeMap_ = map[LnsInt]string {
  TransUnit_DeclFuncMode__Class: "DeclFuncMode.Class",
  TransUnit_DeclFuncMode__Func: "DeclFuncMode.Func",
  TransUnit_DeclFuncMode__Glue: "DeclFuncMode.Glue",
  TransUnit_DeclFuncMode__Module: "DeclFuncMode.Module",
}
func TransUnit_DeclFuncMode__from_1027_(arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_DeclFuncModeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_DeclFuncMode_getTxt(arg1 LnsInt) string {
    return TransUnit_DeclFuncModeMap_[arg1];
}
// decl enum -- ExpSymbolMode 
type TransUnit_ExpSymbolMode = LnsInt
const TransUnit_ExpSymbolMode__Field = 2
const TransUnit_ExpSymbolMode__FieldNil = 3
const TransUnit_ExpSymbolMode__Fn = 1
const TransUnit_ExpSymbolMode__Get = 4
const TransUnit_ExpSymbolMode__GetNil = 5
const TransUnit_ExpSymbolMode__Symbol = 0
var TransUnit_ExpSymbolModeList_ = NewLnsList( []LnsAny {
  TransUnit_ExpSymbolMode__Symbol,
  TransUnit_ExpSymbolMode__Fn,
  TransUnit_ExpSymbolMode__Field,
  TransUnit_ExpSymbolMode__FieldNil,
  TransUnit_ExpSymbolMode__Get,
  TransUnit_ExpSymbolMode__GetNil,
})
func TransUnit_ExpSymbolMode_get__allList_1049_() *LnsList{
    return TransUnit_ExpSymbolModeList_
}
var TransUnit_ExpSymbolModeMap_ = map[LnsInt]string {
  TransUnit_ExpSymbolMode__Field: "ExpSymbolMode.Field",
  TransUnit_ExpSymbolMode__FieldNil: "ExpSymbolMode.FieldNil",
  TransUnit_ExpSymbolMode__Fn: "ExpSymbolMode.Fn",
  TransUnit_ExpSymbolMode__Get: "ExpSymbolMode.Get",
  TransUnit_ExpSymbolMode__GetNil: "ExpSymbolMode.GetNil",
  TransUnit_ExpSymbolMode__Symbol: "ExpSymbolMode.Symbol",
}
func TransUnit_ExpSymbolMode__from_1042_(arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_ExpSymbolModeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_ExpSymbolMode_getTxt(arg1 LnsInt) string {
    return TransUnit_ExpSymbolModeMap_[arg1];
}
// decl enum -- AnalyzeMode 
type TransUnit_AnalyzeMode = LnsInt
const TransUnit_AnalyzeMode__Compile = 0
const TransUnit_AnalyzeMode__Complete = 2
const TransUnit_AnalyzeMode__Diag = 1
const TransUnit_AnalyzeMode__Inquire = 3
var TransUnit_AnalyzeModeList_ = NewLnsList( []LnsAny {
  TransUnit_AnalyzeMode__Compile,
  TransUnit_AnalyzeMode__Diag,
  TransUnit_AnalyzeMode__Complete,
  TransUnit_AnalyzeMode__Inquire,
})
func TransUnit_AnalyzeMode_get__allList() *LnsList{
    return TransUnit_AnalyzeModeList_
}
var TransUnit_AnalyzeModeMap_ = map[LnsInt]string {
  TransUnit_AnalyzeMode__Compile: "AnalyzeMode.Compile",
  TransUnit_AnalyzeMode__Complete: "AnalyzeMode.Complete",
  TransUnit_AnalyzeMode__Diag: "AnalyzeMode.Diag",
  TransUnit_AnalyzeMode__Inquire: "AnalyzeMode.Inquire",
}
func TransUnit_AnalyzeMode__from(arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_AnalyzeModeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_AnalyzeMode_getTxt(arg1 LnsInt) string {
    return TransUnit_AnalyzeModeMap_[arg1];
}
// decl enum -- AnalyzingState 
type TransUnit_AnalyzingState = LnsInt
const TransUnit_AnalyzingState__ClassMethod = 3
const TransUnit_AnalyzingState__Constructor = 1
const TransUnit_AnalyzingState__Func = 4
const TransUnit_AnalyzingState__InitBlock = 2
const TransUnit_AnalyzingState__Other = 0
var TransUnit_AnalyzingStateList_ = NewLnsList( []LnsAny {
  TransUnit_AnalyzingState__Other,
  TransUnit_AnalyzingState__Constructor,
  TransUnit_AnalyzingState__InitBlock,
  TransUnit_AnalyzingState__ClassMethod,
  TransUnit_AnalyzingState__Func,
})
func TransUnit_AnalyzingState_get__allList_1260_() *LnsList{
    return TransUnit_AnalyzingStateList_
}
var TransUnit_AnalyzingStateMap_ = map[LnsInt]string {
  TransUnit_AnalyzingState__ClassMethod: "AnalyzingState.ClassMethod",
  TransUnit_AnalyzingState__Constructor: "AnalyzingState.Constructor",
  TransUnit_AnalyzingState__Func: "AnalyzingState.Func",
  TransUnit_AnalyzingState__InitBlock: "AnalyzingState.InitBlock",
  TransUnit_AnalyzingState__Other: "AnalyzingState.Other",
}
func TransUnit_AnalyzingState__from_1253_(arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_AnalyzingStateMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_AnalyzingState_getTxt(arg1 LnsInt) string {
    return TransUnit_AnalyzingStateMap_[arg1];
}
// decl enum -- SymbolMode 
type TransUnit_SymbolMode = LnsInt
const TransUnit_SymbolMode__MustNot_ = 1
const TransUnit_SymbolMode__MustNot_Or_ = 2
const TransUnit_SymbolMode__Must_ = 0
var TransUnit_SymbolModeList_ = NewLnsList( []LnsAny {
  TransUnit_SymbolMode__Must_,
  TransUnit_SymbolMode__MustNot_,
  TransUnit_SymbolMode__MustNot_Or_,
})
func TransUnit_SymbolMode_get__allList_4708_() *LnsList{
    return TransUnit_SymbolModeList_
}
var TransUnit_SymbolModeMap_ = map[LnsInt]string {
  TransUnit_SymbolMode__MustNot_: "SymbolMode.MustNot_",
  TransUnit_SymbolMode__MustNot_Or_: "SymbolMode.MustNot_Or_",
  TransUnit_SymbolMode__Must_: "SymbolMode.Must_",
}
func TransUnit_SymbolMode__from_4701_(arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_SymbolModeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_SymbolMode_getTxt(arg1 LnsInt) string {
    return TransUnit_SymbolModeMap_[arg1];
}
// decl enum -- TentativeMode 
type TransUnit_TentativeMode = LnsInt
const TransUnit_TentativeMode__Finish = 5
const TransUnit_TentativeMode__Ignore = 0
const TransUnit_TentativeMode__Loop = 1
const TransUnit_TentativeMode__Merge = 4
const TransUnit_TentativeMode__Simple = 2
const TransUnit_TentativeMode__Start = 3
var TransUnit_TentativeModeList_ = NewLnsList( []LnsAny {
  TransUnit_TentativeMode__Ignore,
  TransUnit_TentativeMode__Loop,
  TransUnit_TentativeMode__Simple,
  TransUnit_TentativeMode__Start,
  TransUnit_TentativeMode__Merge,
  TransUnit_TentativeMode__Finish,
})
func TransUnit_TentativeMode_get__allList_4791_() *LnsList{
    return TransUnit_TentativeModeList_
}
var TransUnit_TentativeModeMap_ = map[LnsInt]string {
  TransUnit_TentativeMode__Finish: "TentativeMode.Finish",
  TransUnit_TentativeMode__Ignore: "TentativeMode.Ignore",
  TransUnit_TentativeMode__Loop: "TentativeMode.Loop",
  TransUnit_TentativeMode__Merge: "TentativeMode.Merge",
  TransUnit_TentativeMode__Simple: "TentativeMode.Simple",
  TransUnit_TentativeMode__Start: "TentativeMode.Start",
}
func TransUnit_TentativeMode__from_4784_(arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_TentativeModeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_TentativeMode_getTxt(arg1 LnsInt) string {
    return TransUnit_TentativeModeMap_[arg1];
}
// decl enum -- FormType 
type TransUnit_FormType = LnsInt
const TransUnit_FormType__Match = 0
const TransUnit_FormType__NeedConv = 1
const TransUnit_FormType__Unmatch = 2
var TransUnit_FormTypeList_ = NewLnsList( []LnsAny {
  TransUnit_FormType__Match,
  TransUnit_FormType__NeedConv,
  TransUnit_FormType__Unmatch,
})
func TransUnit_FormType_get__allList() *LnsList{
    return TransUnit_FormTypeList_
}
var TransUnit_FormTypeMap_ = map[LnsInt]string {
  TransUnit_FormType__Match: "FormType.Match",
  TransUnit_FormType__NeedConv: "FormType.NeedConv",
  TransUnit_FormType__Unmatch: "FormType.Unmatch",
}
func TransUnit_FormType__from(arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_FormTypeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_FormType_getTxt(arg1 LnsInt) string {
    return TransUnit_FormTypeMap_[arg1];
}
var TransUnit_op2levelMap *LnsMap
var TransUnit_op1levelMap *LnsMap
var TransUnit_opTopLevel LnsInt
var TransUnit_quotedChar2Code *LnsMap
var TransUnit_builtinFunc *TransUnit_BuiltinFuncType
var TransUnit_readyBuiltin bool
var TransUnit_specialSymbolSet *LnsSet
var TransUnit_builtinKeywordSet *LnsSet
var TransUnit_CantOverrideMethods *LnsSet
type TransUnit_checkImplicitCastCallback_1263_ func (arg1 *Ast_TypeInfo,arg2 *Nodes_Node) LnsAny
type TransUnit_checkImplicitCastCallback_6090_ func (arg1 *Ast_TypeInfo,arg2 *Nodes_Node) LnsAny
type TransUnit_checkCompForm_6433_ func (arg1 *Writer_JSON,arg2 string)
// for 2298
func TransUnit_convExp20589(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 3906
func TransUnit_convExp28928(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 4085
func TransUnit_convExp29876(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 348
func TransUnit_convExp33994(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 396
func TransUnit_convExp54460(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1193
func TransUnit_convExp58158(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1189
func TransUnit_convExp58160(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1351
func TransUnit_convExp58829(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2585
func TransUnit_convExp65190(arg1 []LnsAny) (LnsAny, LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 )
}
// for 3195
func TransUnit_convExp68291(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 53
func TransUnit_convExp73619(arg1 []LnsAny) (LnsAny, LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 )
}
// for 2290
func TransUnit_convExp20471(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 3224
func TransUnit_convExp25147(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3227
func TransUnit_convExp25170(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3231
func TransUnit_convExp25202(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3242
func TransUnit_convExp25271(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3245
func TransUnit_convExp25294(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3248
func TransUnit_convExp25317(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3251
func TransUnit_convExp25340(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3254
func TransUnit_convExp25363(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3257
func TransUnit_convExp25386(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3260
func TransUnit_convExp25409(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 3263
func TransUnit_convExp25432(arg1 []LnsAny) (LnsAny, LnsAny) {
    return TransUnit__TypeInfoDownCastF(Lns_getFromMulti( arg1, 0 )), Lns_getFromMulti( arg1, 1 )
}
// for 995
func TransUnit_convExp37295(arg1 []LnsAny) (*Types_Token, LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 )
}
// for 1071
func TransUnit_convExp37723(arg1 []LnsAny) (*Types_Token, LnsAny, *LnsList) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 ).(*LnsList)
}
// for 2121
func TransUnit_convExp62768(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1514
func TransUnit_convExp6455(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1532
func TransUnit_convExp6553(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1610
func TransUnit_convExp7058(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1899
func TransUnit_convExp8956(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2293
func TransUnit_convExp20496(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2314
func TransUnit_convExp20644(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 3158
func TransUnit_convExp24743(arg1 []LnsAny) (*Lns_luaValue, *LnsMap) {
    return Lns_getFromMulti( arg1, 0 ).(*Lns_luaValue), Lns_getFromMulti( arg1, 1 ).(*LnsMap)
}
// for 3404
func TransUnit_convExp26152(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 3827
func TransUnit_convExp28383(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2413
func TransUnit_convExp44812(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 834
func TransUnit_convExp56376(arg1 []LnsAny) (LnsInt, *LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*LnsMap), Lns_getFromMulti( arg1, 2 )
}
// for 903
func TransUnit_convExp56721(arg1 []LnsAny) (LnsInt, *LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*LnsMap), Lns_getFromMulti( arg1, 2 )
}
// for 2120
func TransUnit_convExp62753(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2122
func TransUnit_convExp62783(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2126
func TransUnit_convExp62823(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 3611
func TransUnit_convExp70613(arg1 []LnsAny) (LnsInt, *LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*LnsMap), Lns_getFromMulti( arg1, 2 )
}

// 90: decl @lune.@base.@TransUnit.clearThePosForModToRef
func TransUnit_clearThePosForModToRef_1103_(scope *Ast_Scope,moduleScope *Ast_Scope) *LnsList {
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    scope.FP.FilterSymbolTypeInfo(scope, moduleScope, Ast_ScopeAccess__Normal, Ast_filterForm(func(symInfo *Ast_SymbolInfo) bool {
        if symInfo.FP.Get_kind() == Ast_SymbolKind__Var{
            list.Insert(TransUnit_RefAccessSymPos2Stem(NewTransUnit_RefAccessSymPos(symInfo, symInfo.FP.Get_posForModToRef())))
            symInfo.FP.Set_posForModToRef(nil)
        }
        return true
    }))
    return list
}







// 1990: decl @lune.@base.@TransUnit.getBuiltinFunc
func TransUnit_getBuiltinFunc() *TransUnit_BuiltinFuncType {
    return TransUnit_builtinFunc
}



















// 1994: decl @lune.@base.@TransUnit.setupBuiltinTypeInfo
func TransUnit_setupBuiltinTypeInfo_3435_(name string,fieldName string,symInfo *Ast_SymbolInfo) {
    var typeInfo *Ast_TypeInfo
    typeInfo = symInfo.FP.Get_typeInfo()
    var process_ func()
    process_ = func() {
        if _switch15073 := fieldName; _switch15073 == "_fcall" {
            TransUnit_builtinFunc.Lns__fcall = typeInfo
            
            TransUnit_builtinFunc.Lns__fcall_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "_kind" {
            TransUnit_builtinFunc.Lns__kind = typeInfo
            
            TransUnit_builtinFunc.Lns__kind_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "_load" {
            TransUnit_builtinFunc.Lns__load = typeInfo
            
            TransUnit_builtinFunc.Lns__load_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "collectgarbage" {
            TransUnit_builtinFunc.Lns_collectgarbage = typeInfo
            
            TransUnit_builtinFunc.Lns_collectgarbage_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "error" {
            TransUnit_builtinFunc.Lns_error = typeInfo
            
            TransUnit_builtinFunc.Lns_error_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "expandLuavalMap" {
            TransUnit_builtinFunc.Lns_expandLuavalMap = typeInfo
            
            TransUnit_builtinFunc.Lns_expandLuavalMap_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "loadfile" {
            TransUnit_builtinFunc.Lns_loadfile = typeInfo
            
            TransUnit_builtinFunc.Lns_loadfile_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "print" {
            TransUnit_builtinFunc.Lns_print = typeInfo
            
            TransUnit_builtinFunc.Lns_print_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "require" {
            TransUnit_builtinFunc.Lns_require = typeInfo
            
            TransUnit_builtinFunc.Lns_require_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "tonumber" {
            TransUnit_builtinFunc.Lns_tonumber = typeInfo
            
            TransUnit_builtinFunc.Lns_tonumber_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "tostring" {
            TransUnit_builtinFunc.Lns_tostring = typeInfo
            
            TransUnit_builtinFunc.Lns_tostring_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15073 == "type" {
            TransUnit_builtinFunc.Lns_type = typeInfo
            
            TransUnit_builtinFunc.Lns_type_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_iStream func()
    process_iStream = func() {
        if _switch15163 := fieldName; _switch15163 == "__attrib" {
            TransUnit_builtinFunc.Istream___attrib = typeInfo
            
            TransUnit_builtinFunc.Istream___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15163 == "close" {
            TransUnit_builtinFunc.Istream_close = typeInfo
            
            TransUnit_builtinFunc.Istream_close_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15163 == "read" {
            TransUnit_builtinFunc.Istream_read = typeInfo
            
            TransUnit_builtinFunc.Istream_read_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_oStream func()
    process_oStream = func() {
        if _switch15281 := fieldName; _switch15281 == "__attrib" {
            TransUnit_builtinFunc.Ostream___attrib = typeInfo
            
            TransUnit_builtinFunc.Ostream___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15281 == "close" {
            TransUnit_builtinFunc.Ostream_close = typeInfo
            
            TransUnit_builtinFunc.Ostream_close_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15281 == "flush" {
            TransUnit_builtinFunc.Ostream_flush = typeInfo
            
            TransUnit_builtinFunc.Ostream_flush_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15281 == "write" {
            TransUnit_builtinFunc.Ostream_write = typeInfo
            
            TransUnit_builtinFunc.Ostream_write_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_luaStream func()
    process_luaStream = func() {
        if _switch15455 := fieldName; _switch15455 == "__attrib" {
            TransUnit_builtinFunc.Luastream___attrib = typeInfo
            
            TransUnit_builtinFunc.Luastream___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15455 == "close" {
            TransUnit_builtinFunc.Luastream_close = typeInfo
            
            TransUnit_builtinFunc.Luastream_close_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15455 == "flush" {
            TransUnit_builtinFunc.Luastream_flush = typeInfo
            
            TransUnit_builtinFunc.Luastream_flush_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15455 == "read" {
            TransUnit_builtinFunc.Luastream_read = typeInfo
            
            TransUnit_builtinFunc.Luastream_read_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15455 == "seek" {
            TransUnit_builtinFunc.Luastream_seek = typeInfo
            
            TransUnit_builtinFunc.Luastream_seek_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15455 == "write" {
            TransUnit_builtinFunc.Luastream_write = typeInfo
            
            TransUnit_builtinFunc.Luastream_write_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_Mapping func()
    process_Mapping = func() {
        if _switch15517 := fieldName; _switch15517 == "__attrib" {
            TransUnit_builtinFunc.Mapping___attrib = typeInfo
            
            TransUnit_builtinFunc.Mapping___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15517 == "_toMap" {
            TransUnit_builtinFunc.Mapping__toMap = typeInfo
            
            TransUnit_builtinFunc.Mapping__toMap_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process___pipe func()
    process___pipe = func() {
        if _switch15579 := fieldName; _switch15579 == "get" {
            TransUnit_builtinFunc.G__pipe_get = typeInfo
            
            TransUnit_builtinFunc.G__pipe_get_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15579 == "put" {
            TransUnit_builtinFunc.G__pipe_put = typeInfo
            
            TransUnit_builtinFunc.G__pipe_put_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_LnsThread func()
    process_LnsThread = func() {
        if _switch15641 := fieldName; _switch15641 == "__init" {
            TransUnit_builtinFunc.Lnsthread___init = typeInfo
            
            TransUnit_builtinFunc.Lnsthread___init_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15641 == "loop" {
            TransUnit_builtinFunc.Lnsthread_loop = typeInfo
            
            TransUnit_builtinFunc.Lnsthread_loop_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_io func()
    process_io = func() {
        if _switch15815 := fieldName; _switch15815 == "__attrib" {
            TransUnit_builtinFunc.Io___attrib = typeInfo
            
            TransUnit_builtinFunc.Io___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15815 == "open" {
            TransUnit_builtinFunc.Io_open = typeInfo
            
            TransUnit_builtinFunc.Io_open_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15815 == "popen" {
            TransUnit_builtinFunc.Io_popen = typeInfo
            
            TransUnit_builtinFunc.Io_popen_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15815 == "stderr" {
            TransUnit_builtinFunc.Io_stderr = typeInfo
            
            TransUnit_builtinFunc.Io_stderr_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15815 == "stdin" {
            TransUnit_builtinFunc.Io_stdin = typeInfo
            
            TransUnit_builtinFunc.Io_stdin_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15815 == "stdout" {
            TransUnit_builtinFunc.Io_stdout = typeInfo
            
            TransUnit_builtinFunc.Io_stdout_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_package func()
    process_package = func() {
        if _switch15905 := fieldName; _switch15905 == "__attrib" {
            TransUnit_builtinFunc.Package___attrib = typeInfo
            
            TransUnit_builtinFunc.Package___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15905 == "path" {
            TransUnit_builtinFunc.Package_path = typeInfo
            
            TransUnit_builtinFunc.Package_path_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch15905 == "searchpath" {
            TransUnit_builtinFunc.Package_searchpath = typeInfo
            
            TransUnit_builtinFunc.Package_searchpath_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_os func()
    process_os = func() {
        if _switch16135 := fieldName; _switch16135 == "__attrib" {
            TransUnit_builtinFunc.Os___attrib = typeInfo
            
            TransUnit_builtinFunc.Os___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16135 == "clock" {
            TransUnit_builtinFunc.Os_clock = typeInfo
            
            TransUnit_builtinFunc.Os_clock_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16135 == "date" {
            TransUnit_builtinFunc.Os_date = typeInfo
            
            TransUnit_builtinFunc.Os_date_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16135 == "difftime" {
            TransUnit_builtinFunc.Os_difftime = typeInfo
            
            TransUnit_builtinFunc.Os_difftime_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16135 == "exit" {
            TransUnit_builtinFunc.Os_exit = typeInfo
            
            TransUnit_builtinFunc.Os_exit_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16135 == "remove" {
            TransUnit_builtinFunc.Os_remove = typeInfo
            
            TransUnit_builtinFunc.Os_remove_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16135 == "rename" {
            TransUnit_builtinFunc.Os_rename = typeInfo
            
            TransUnit_builtinFunc.Os_rename_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16135 == "time" {
            TransUnit_builtinFunc.Os_time = typeInfo
            
            TransUnit_builtinFunc.Os_time_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_string func()
    process_string = func() {
        if _switch16477 := fieldName; _switch16477 == "__attrib" {
            TransUnit_builtinFunc.String___attrib = typeInfo
            
            TransUnit_builtinFunc.String___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "byte" {
            TransUnit_builtinFunc.String_byte = typeInfo
            
            TransUnit_builtinFunc.String_byte_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "dump" {
            TransUnit_builtinFunc.String_dump = typeInfo
            
            TransUnit_builtinFunc.String_dump_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "find" {
            TransUnit_builtinFunc.String_find = typeInfo
            
            TransUnit_builtinFunc.String_find_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "format" {
            TransUnit_builtinFunc.String_format = typeInfo
            
            TransUnit_builtinFunc.String_format_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "gmatch" {
            TransUnit_builtinFunc.String_gmatch = typeInfo
            
            TransUnit_builtinFunc.String_gmatch_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "gsub" {
            TransUnit_builtinFunc.String_gsub = typeInfo
            
            TransUnit_builtinFunc.String_gsub_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "lower" {
            TransUnit_builtinFunc.String_lower = typeInfo
            
            TransUnit_builtinFunc.String_lower_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "rep" {
            TransUnit_builtinFunc.String_rep = typeInfo
            
            TransUnit_builtinFunc.String_rep_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "reverse" {
            TransUnit_builtinFunc.String_reverse = typeInfo
            
            TransUnit_builtinFunc.String_reverse_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "sub" {
            TransUnit_builtinFunc.String_sub = typeInfo
            
            TransUnit_builtinFunc.String_sub_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16477 == "upper" {
            TransUnit_builtinFunc.String_upper = typeInfo
            
            TransUnit_builtinFunc.String_upper_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_str func()
    process_str = func() {
        if _switch16791 := fieldName; _switch16791 == "__attrib" {
            TransUnit_builtinFunc.Str___attrib = typeInfo
            
            TransUnit_builtinFunc.Str___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "byte" {
            TransUnit_builtinFunc.Str_byte = typeInfo
            
            TransUnit_builtinFunc.Str_byte_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "find" {
            TransUnit_builtinFunc.Str_find = typeInfo
            
            TransUnit_builtinFunc.Str_find_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "format" {
            TransUnit_builtinFunc.Str_format = typeInfo
            
            TransUnit_builtinFunc.Str_format_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "gmatch" {
            TransUnit_builtinFunc.Str_gmatch = typeInfo
            
            TransUnit_builtinFunc.Str_gmatch_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "gsub" {
            TransUnit_builtinFunc.Str_gsub = typeInfo
            
            TransUnit_builtinFunc.Str_gsub_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "lower" {
            TransUnit_builtinFunc.Str_lower = typeInfo
            
            TransUnit_builtinFunc.Str_lower_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "rep" {
            TransUnit_builtinFunc.Str_rep = typeInfo
            
            TransUnit_builtinFunc.Str_rep_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "reverse" {
            TransUnit_builtinFunc.Str_reverse = typeInfo
            
            TransUnit_builtinFunc.Str_reverse_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "sub" {
            TransUnit_builtinFunc.Str_sub = typeInfo
            
            TransUnit_builtinFunc.Str_sub_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16791 == "upper" {
            TransUnit_builtinFunc.Str_upper = typeInfo
            
            TransUnit_builtinFunc.Str_upper_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_List func()
    process_List = func() {
        if _switch16937 := fieldName; _switch16937 == "__less" {
            TransUnit_builtinFunc.List___less = typeInfo
            
            TransUnit_builtinFunc.List___less_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16937 == "insert" {
            TransUnit_builtinFunc.List_insert = typeInfo
            
            TransUnit_builtinFunc.List_insert_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16937 == "remove" {
            TransUnit_builtinFunc.List_remove = typeInfo
            
            TransUnit_builtinFunc.List_remove_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16937 == "sort" {
            TransUnit_builtinFunc.List_sort = typeInfo
            
            TransUnit_builtinFunc.List_sort_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch16937 == "unpack" {
            TransUnit_builtinFunc.List_unpack = typeInfo
            
            TransUnit_builtinFunc.List_unpack_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_Array func()
    process_Array = func() {
        if _switch17027 := fieldName; _switch17027 == "__less" {
            TransUnit_builtinFunc.Array___less = typeInfo
            
            TransUnit_builtinFunc.Array___less_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17027 == "sort" {
            TransUnit_builtinFunc.Array_sort = typeInfo
            
            TransUnit_builtinFunc.Array_sort_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17027 == "unpack" {
            TransUnit_builtinFunc.Array_unpack = typeInfo
            
            TransUnit_builtinFunc.Array_unpack_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_Set func()
    process_Set = func() {
        if _switch17257 := fieldName; _switch17257 == "add" {
            TransUnit_builtinFunc.Set_add = typeInfo
            
            TransUnit_builtinFunc.Set_add_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17257 == "and" {
            TransUnit_builtinFunc.Set_and = typeInfo
            
            TransUnit_builtinFunc.Set_and_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17257 == "clone" {
            TransUnit_builtinFunc.Set_clone = typeInfo
            
            TransUnit_builtinFunc.Set_clone_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17257 == "del" {
            TransUnit_builtinFunc.Set_del = typeInfo
            
            TransUnit_builtinFunc.Set_del_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17257 == "has" {
            TransUnit_builtinFunc.Set_has = typeInfo
            
            TransUnit_builtinFunc.Set_has_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17257 == "len" {
            TransUnit_builtinFunc.Set_len = typeInfo
            
            TransUnit_builtinFunc.Set_len_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17257 == "or" {
            TransUnit_builtinFunc.Set_or = typeInfo
            
            TransUnit_builtinFunc.Set_or_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17257 == "sub" {
            TransUnit_builtinFunc.Set_sub = typeInfo
            
            TransUnit_builtinFunc.Set_sub_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_math func()
    process_math = func() {
        if _switch17347 := fieldName; _switch17347 == "__attrib" {
            TransUnit_builtinFunc.Math___attrib = typeInfo
            
            TransUnit_builtinFunc.Math___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17347 == "random" {
            TransUnit_builtinFunc.Math_random = typeInfo
            
            TransUnit_builtinFunc.Math_random_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17347 == "randomseed" {
            TransUnit_builtinFunc.Math_randomseed = typeInfo
            
            TransUnit_builtinFunc.Math_randomseed_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_debug func()
    process_debug = func() {
        if _switch17437 := fieldName; _switch17437 == "__attrib" {
            TransUnit_builtinFunc.Debug___attrib = typeInfo
            
            TransUnit_builtinFunc.Debug___attrib_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17437 == "getinfo" {
            TransUnit_builtinFunc.Debug_getinfo = typeInfo
            
            TransUnit_builtinFunc.Debug_getinfo_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        } else if _switch17437 == "getlocal" {
            TransUnit_builtinFunc.Debug_getlocal = typeInfo
            
            TransUnit_builtinFunc.Debug_getlocal_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    var process_Nilable func()
    process_Nilable = func() {
        if _switch17471 := fieldName; _switch17471 == "val" {
            TransUnit_builtinFunc.Nilable_val = typeInfo
            
            TransUnit_builtinFunc.Nilable_val_sym = symInfo
            
            TransUnit_builtinFunc.FP.Register(symInfo)
        }
    }
    if _switch17622 := name; _switch17622 == "" {
        process_()
    } else if _switch17622 == "iStream" {
        process_iStream()
    } else if _switch17622 == "oStream" {
        process_oStream()
    } else if _switch17622 == "luaStream" {
        process_luaStream()
    } else if _switch17622 == "Mapping" {
        process_Mapping()
    } else if _switch17622 == "__pipe" {
        process___pipe()
    } else if _switch17622 == "LnsThread" {
        process_LnsThread()
    } else if _switch17622 == "io" {
        process_io()
    } else if _switch17622 == "package" {
        process_package()
    } else if _switch17622 == "os" {
        process_os()
    } else if _switch17622 == "string" {
        process_string()
    } else if _switch17622 == "str" {
        process_str()
    } else if _switch17622 == "List" {
        process_List()
    } else if _switch17622 == "Array" {
        process_Array()
    } else if _switch17622 == "Set" {
        process_Set()
    } else if _switch17622 == "math" {
        process_math()
    } else if _switch17622 == "debug" {
        process_debug()
    } else if _switch17622 == "Nilable" {
        process_Nilable()
    }
}

// 2004: decl @lune.@base.@TransUnit.getBuiltInInfo
func TransUnit_getBuiltInInfo_3512_() *LnsList {
    return NewLnsList([]LnsAny{NewLnsMap( map[LnsAny]LnsAny{"":NewLnsMap( map[LnsAny]LnsAny{"_fcall":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"form", "&..."}),"ret":NewLnsList([]LnsAny{""}),}),"_kind":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"int"}),}),"_load":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "stem!"}),"ret":NewLnsList([]LnsAny{"Luaval<form>!", "str!"}),}),"collectgarbage":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),}),"error":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"__"}),}),"expandLuavalMap":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"Luaval<&stem>!"}),"ret":NewLnsList([]LnsAny{"&stem!"}),}),"loadfile":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<form>!", "str!"}),}),"print":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&..."}),"ret":NewLnsList([]LnsAny{}),}),"require":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<stem>!"}),}),"tonumber":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int!"}),"ret":NewLnsList([]LnsAny{"real!"}),}),"tostring":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&stem"}),"ret":NewLnsList([]LnsAny{"str"}),}),"type":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&stem!"}),"ret":NewLnsList([]LnsAny{"str"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"iStream":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"close":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"read":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"oStream":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"close":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"flush":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"write":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"stem!", "str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"luaStream":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"inplements":NewLnsList([]LnsAny{"iStream", "oStream"}),"type":NewLnsList([]LnsAny{"interface"}),}),"close":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"flush":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"read":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"seek":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int"}),"ret":NewLnsList([]LnsAny{"int!", "str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"write":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"stem!", "str!"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Mapping":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"interface"}),}),"_toMap":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"__pipe<T>":NewLnsMap( map[LnsAny]LnsAny{"get":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"T!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"put":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T!"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"LnsThread":NewLnsMap( map[LnsAny]LnsAny{"__init":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"method"}),}),"loop":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"abstract"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"io":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"open":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str!"}),"ret":NewLnsList([]LnsAny{"luaStream!", "str!"}),}),"popen":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<luaStream>!"}),}),"stderr":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"oStream"}),}),"stdin":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"iStream"}),}),"stdout":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"oStream"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"package":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"path":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"var"}),"typeInfo":NewLnsList([]LnsAny{"str"}),}),"searchpath":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"str!"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"os":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"clock":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"real"}),}),"date":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str!", "stem!"}),"ret":NewLnsList([]LnsAny{"Luaval<stem>!"}),}),"difftime":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&stem", "&stem"}),"ret":NewLnsList([]LnsAny{"int"}),}),"exit":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!"}),"ret":NewLnsList([]LnsAny{"__"}),}),"remove":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"bool!", "str!"}),}),"rename":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"stem!", "str!"}),}),"time":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"stem!"}),"ret":NewLnsList([]LnsAny{"stem!"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"string":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"byte":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int!", "int!"}),"ret":NewLnsList([]LnsAny{"...<int>"}),}),"dump":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Luaval<form>", "bool!"}),"ret":NewLnsList([]LnsAny{"str"}),}),"find":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str", "int!", "bool!"}),"ret":NewLnsList([]LnsAny{"...<int>"}),}),"format":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "..."}),"ret":NewLnsList([]LnsAny{"str"}),}),"gmatch":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"Luaval<form>", "stem!", "stem!"}),}),"gsub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str", "str"}),"ret":NewLnsList([]LnsAny{"str", "int"}),}),"lower":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"str"}),}),"rep":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int"}),"ret":NewLnsList([]LnsAny{"str"}),}),"reverse":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"str"}),}),"sub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int", "int!"}),"ret":NewLnsList([]LnsAny{"str"}),}),"upper":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"str"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"str":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"inplements":NewLnsList([]LnsAny{"Mapping"}),}),"byte":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!", "int!"}),"ret":NewLnsList([]LnsAny{"...<int!>"}),"type":NewLnsList([]LnsAny{"method"}),}),"find":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "int!", "bool!"}),"ret":NewLnsList([]LnsAny{"...<int>"}),"type":NewLnsList([]LnsAny{"method"}),}),"format":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&..."}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"gmatch":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str"}),"ret":NewLnsList([]LnsAny{"Luaval<form>", "stem!", "stem!"}),"type":NewLnsList([]LnsAny{"method"}),}),"gsub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"str", "str"}),"ret":NewLnsList([]LnsAny{"str", "int"}),"type":NewLnsList([]LnsAny{"method"}),}),"lower":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"rep":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int"}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"reverse":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"sub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int", "int!"}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),"upper":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"str"}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"List<T>":NewLnsMap( map[LnsAny]LnsAny{"__less":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T", "T"}),"ret":NewLnsList([]LnsAny{"bool"}),"type":NewLnsList([]LnsAny{"formfunc"}),}),"insert":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{""}),"type":NewLnsList([]LnsAny{"mut"}),}),"remove":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!"}),"ret":NewLnsList([]LnsAny{"T!"}),"type":NewLnsList([]LnsAny{"mut"}),}),"sort":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"__less!"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"unpack":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"..."}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Array<T>":NewLnsMap( map[LnsAny]LnsAny{"__less":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T", "T"}),"ret":NewLnsList([]LnsAny{"bool"}),"type":NewLnsList([]LnsAny{"formfunc"}),}),"sort":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"__less!"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"unpack":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"..."}),"type":NewLnsList([]LnsAny{"method"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Set<T>":NewLnsMap( map[LnsAny]LnsAny{"add":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"and":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Set<T>"}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"mut"}),}),"clone":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"method"}),}),"del":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{}),"type":NewLnsList([]LnsAny{"mut"}),}),"has":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"T"}),"ret":NewLnsList([]LnsAny{"bool"}),"type":NewLnsList([]LnsAny{"method"}),}),"len":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"int"}),"type":NewLnsList([]LnsAny{"method"}),}),"or":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Set<T>"}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"mut"}),}),"sub":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"&Set<T>"}),"ret":NewLnsList([]LnsAny{"Set<T>"}),"type":NewLnsList([]LnsAny{"mut"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"math":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"random":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int!", "int!"}),"ret":NewLnsList([]LnsAny{"real"}),}),"randomseed":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int"}),"ret":NewLnsList([]LnsAny{}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"debug":NewLnsMap( map[LnsAny]LnsAny{"__attrib":NewLnsMap( map[LnsAny]LnsAny{"type":NewLnsList([]LnsAny{"module"}),}),"getinfo":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int"}),"ret":NewLnsList([]LnsAny{"Map<str,stem>!"}),}),"getlocal":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{"int", "int"}),"ret":NewLnsList([]LnsAny{"str!", "stem!"}),}),}),}), NewLnsMap( map[LnsAny]LnsAny{"Nilable<_T>":NewLnsMap( map[LnsAny]LnsAny{"val":NewLnsMap( map[LnsAny]LnsAny{"arg":NewLnsList([]LnsAny{}),"ret":NewLnsList([]LnsAny{"_T!"}),"type":NewLnsList([]LnsAny{"method"}),}),}),})})
}

// 2182: decl @lune.@base.@TransUnit.isStrFormFunc
func TransUnit_isStrFormFunc(typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_srcTypeInfo() == TransUnit_builtinFunc.String_format{
        return true
    }
    return false
}




// 2453: decl @lune.@base.@TransUnit.TransUnit.registBuiltInScope.processCopyAlterList
func TransUnit_registBuiltInScope__processCopyAlterList_4630_(alterList *LnsList,typeList *LnsList) {
    for _, _typeInfo := range( typeList.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        alterList.Insert(Ast_AlternateTypeInfo2Stem(Lns_unwrap( Ast_AlternateTypeInfoDownCastF(typeInfo.FP)).(*Ast_AlternateTypeInfo)))
    }
}















// 632: decl @lune.@base.@TransUnit.TransUnit.createAST.createId2proto
func TransUnit_createAST__createId2proto_5304_(_map *LnsMap) *LnsMap {
    var id2proto *LnsMap
    id2proto = NewLnsMap( map[LnsAny]LnsAny{})
    for _protoType, _ := range( _map.Items ) {
        protoType := _protoType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        id2proto.Set(protoType.FP.Get_typeId(),protoType)
    }
    return id2proto
}















func TransUnit_analyzeInitExp___anonymous_5892_(dstType *Ast_TypeInfo,expNode *Nodes_Node) LnsAny {
    return nil
}




// 1183: decl @lune.@base.@TransUnit.findForm
func TransUnit_findForm(format string) *LnsList {
    var remain string
    remain = format
    var opList *LnsList
    opList = NewLnsList([]LnsAny{})
    for  {
        var pos LnsAny
        var endPos LnsAny
        pos,endPos = nil, nil
        {
            _index, _endIndex := TransUnit_convExp58160(Lns_2DDD(Lns_getVM().String_find(remain,"^%%[%d]*%a", nil, nil)))
            if _index != nil && _endIndex != nil {
                index := _index.(LnsInt)
                endIndex := _endIndex.(LnsInt)
                pos, endPos = index, endIndex
                
            } else {
                {
                    _index, _endIndex := TransUnit_convExp58158(Lns_2DDD(Lns_getVM().String_find(remain,"[^%%]%%[%d]*%a", nil, nil)))
                    if _index != nil && _endIndex != nil {
                        index := _index.(LnsInt)
                        endIndex := _endIndex.(LnsInt)
                        pos, endPos = index + 1, endIndex
                        
                    }
                }
            }
        }
        if pos != nil && endPos != nil{
            pos_10171 := pos.(LnsInt)
            endPos_10172 := endPos.(LnsInt)
            var op string
            op = Lns_getVM().String_sub(remain,pos_10171, endPos_10172)
            opList.Insert(op)
            remain = Lns_getVM().String_sub(remain,endPos_10172 + 1, nil)
            
        } else {
            break
        }
    }
    return opList
}

// 1215: decl @lune.@base.@TransUnit.isMatchStringFormatType
func TransUnit_isMatchStringFormatType(opKind string,argType *Ast_TypeInfo,luaVer *LuaVer_LuaVerInfo)(LnsInt, *Ast_TypeInfo) {
    {
        _enumType := Ast_EnumTypeInfoDownCastF(argType.FP.Get_srcTypeInfo().FP.Get_aliasSrc().FP)
        if _enumType != nil {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            argType = enumType.FP.Get_valTypeInfo()
            
        }
    }
    if _switch58366 := LnsInt(opKind[len(opKind)-1]); _switch58366 == 115 {
        if argType.FP.Get_srcTypeInfo() != Ast_builtinTypeString{
            if Lns_op_not(luaVer.FP.Get_canFormStem2Str()){
                return TransUnit_FormType__NeedConv, Ast_builtinTypeString
            }
        }
    } else if _switch58366 == 113 {
        if argType.FP.Get_srcTypeInfo() != Ast_builtinTypeString{
            return TransUnit_FormType__Unmatch, Ast_builtinTypeString
        }
    } else if _switch58366 == 65 || _switch58366 == 97 || _switch58366 == 69 || _switch58366 == 101 || _switch58366 == 102 || _switch58366 == 71 || _switch58366 == 103 {
        if argType.FP.Get_srcTypeInfo() != Ast_builtinTypeReal{
            return TransUnit_FormType__Unmatch, Ast_builtinTypeReal
        }
    } else {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( argType.FP.Get_srcTypeInfo() != Ast_builtinTypeInt) &&
            Lns_GetEnv().SetStackVal( argType.FP.Get_srcTypeInfo() != Ast_builtinTypeChar) ).(bool)){
            return TransUnit_FormType__Unmatch, Ast_builtinTypeInt
        }
    }
    return TransUnit_FormType__Match, Ast_builtinTypeNone
}








func TransUnit_analyzeExpSymbol___anonymous_6521_(workSymbolInfo *Ast_SymbolInfo) bool {
    Lns_print([]LnsAny{"sym", workSymbolInfo.FP.Get_name()})
    return true
}
// 2790: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOpSet.process
func TransUnit_analyzeExpOpSet__process_6535_(lValNode *Nodes_Node) LnsAny {
    var refItemNode *Nodes_ExpRefItemNode
    
    {
        _refItemNode := Nodes_ExpRefItemNodeDownCastF(lValNode.FP)
        if _refItemNode == nil{
            return nil
        } else {
            refItemNode = _refItemNode.(*Nodes_ExpRefItemNode)
        }
    }
    if _switch66217 := refItemNode.FP.Get_val().FP.Get_expType().FP.Get_kind(); _switch66217 == Ast_TypeInfoKind__List || _switch66217 == Ast_TypeInfoKind__Map {
        return refItemNode
    }
    return nil
}

// 2957: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOpEquals.getType
func TransUnit_analyzeExpOpEquals__getType_6565_(typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    var workType *Ast_TypeInfo
    workType = typeInfo.FP.Get_nonnilableType().FP.Get_srcTypeInfo()
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( workType.FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
        Lns_GetEnv().SetStackVal( workType.FP.HasBase()) ).(bool)){
        return workType.FP.Get_baseTypeInfo()
    }
    return workType
}




// declaration Class -- AccessSymPos
type TransUnit_AccessSymPosMtd interface {
    Get_pos() *Types_Position
    Get_symbol() *Ast_SymbolInfo
}
type TransUnit_AccessSymPos struct {
    symbol *Ast_SymbolInfo
    pos *Types_Position
    FP TransUnit_AccessSymPosMtd
}
func TransUnit_AccessSymPos2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_AccessSymPos).FP
}
type TransUnit_AccessSymPosDownCast interface {
    ToTransUnit_AccessSymPos() *TransUnit_AccessSymPos
}
func TransUnit_AccessSymPosDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_AccessSymPosDownCast)
    if ok { return work.ToTransUnit_AccessSymPos() }
    return nil
}
func (obj *TransUnit_AccessSymPos) ToTransUnit_AccessSymPos() *TransUnit_AccessSymPos {
    return obj
}
func NewTransUnit_AccessSymPos(arg1 *Ast_SymbolInfo, arg2 *Types_Position) *TransUnit_AccessSymPos {
    obj := &TransUnit_AccessSymPos{}
    obj.FP = obj
    obj.InitTransUnit_AccessSymPos(arg1, arg2)
    return obj
}
func (self *TransUnit_AccessSymPos) InitTransUnit_AccessSymPos(arg1 *Ast_SymbolInfo, arg2 *Types_Position) {
    self.symbol = arg1
    self.pos = arg2
}
func (self *TransUnit_AccessSymPos) Get_symbol() *Ast_SymbolInfo{ return self.symbol }
func (self *TransUnit_AccessSymPos) Get_pos() *Types_Position{ return self.pos }

// declaration Class -- RefAccessSymPos
type TransUnit_RefAccessSymPosMtd interface {
    Get_pos() LnsAny
    Get_symbol() *Ast_SymbolInfo
}
type TransUnit_RefAccessSymPos struct {
    symbol *Ast_SymbolInfo
    pos LnsAny
    FP TransUnit_RefAccessSymPosMtd
}
func TransUnit_RefAccessSymPos2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_RefAccessSymPos).FP
}
type TransUnit_RefAccessSymPosDownCast interface {
    ToTransUnit_RefAccessSymPos() *TransUnit_RefAccessSymPos
}
func TransUnit_RefAccessSymPosDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_RefAccessSymPosDownCast)
    if ok { return work.ToTransUnit_RefAccessSymPos() }
    return nil
}
func (obj *TransUnit_RefAccessSymPos) ToTransUnit_RefAccessSymPos() *TransUnit_RefAccessSymPos {
    return obj
}
func NewTransUnit_RefAccessSymPos(arg1 *Ast_SymbolInfo, arg2 LnsAny) *TransUnit_RefAccessSymPos {
    obj := &TransUnit_RefAccessSymPos{}
    obj.FP = obj
    obj.InitTransUnit_RefAccessSymPos(arg1, arg2)
    return obj
}
func (self *TransUnit_RefAccessSymPos) InitTransUnit_RefAccessSymPos(arg1 *Ast_SymbolInfo, arg2 LnsAny) {
    self.symbol = arg1
    self.pos = arg2
}
func (self *TransUnit_RefAccessSymPos) Get_symbol() *Ast_SymbolInfo{ return self.symbol }
func (self *TransUnit_RefAccessSymPos) Get_pos() LnsAny{ return self.pos }

// declaration Class -- TentativeSymbol
type TransUnit_TentativeSymbolMtd interface {
    AddAccessSym(arg1 *Ast_SymbolInfo)
    AddAccessSymPos(arg1 *TransUnit_AccessSymPos)
    CheckAndExclude(arg1 *Ast_SymbolInfo) bool
    ClearAccessSym()
    Finish(arg1 bool) LnsAny
    Get_accessSymList() *LnsList
    Get_initSymSet() *LnsSet
    Get_scope() *Ast_Scope
    merge(arg1 bool) bool
    ModSym(arg1 *Ast_SymbolInfo)
    NewSet(arg1 *Ast_Scope)
    Regist(arg1 *Ast_SymbolInfo, arg2 *Types_Position) bool
    Skip()
    SyncPos()
}
type TransUnit_TentativeSymbol struct {
    symbolSet *LnsSet
    initSymSet *LnsSet
    oldSymbolSet LnsAny
    parent LnsAny
    scope *Ast_Scope
    skipFlag bool
    loopFlag bool
    accessSymList *LnsList
    modSymbolSet *LnsSet
    accessSymSet *LnsSet
    sym2posForModToRef *LnsList
    FP TransUnit_TentativeSymbolMtd
}
func TransUnit_TentativeSymbol2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_TentativeSymbol).FP
}
type TransUnit_TentativeSymbolDownCast interface {
    ToTransUnit_TentativeSymbol() *TransUnit_TentativeSymbol
}
func TransUnit_TentativeSymbolDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_TentativeSymbolDownCast)
    if ok { return work.ToTransUnit_TentativeSymbol() }
    return nil
}
func (obj *TransUnit_TentativeSymbol) ToTransUnit_TentativeSymbol() *TransUnit_TentativeSymbol {
    return obj
}
func NewTransUnit_TentativeSymbol(arg1 LnsAny, arg2 *Ast_Scope, arg3 *Ast_Scope, arg4 bool, arg5 LnsAny) *TransUnit_TentativeSymbol {
    obj := &TransUnit_TentativeSymbol{}
    obj.FP = obj
    obj.InitTransUnit_TentativeSymbol(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *TransUnit_TentativeSymbol) Get_initSymSet() *LnsSet{ return self.initSymSet }
func (self *TransUnit_TentativeSymbol) Get_scope() *Ast_Scope{ return self.scope }
func (self *TransUnit_TentativeSymbol) Get_accessSymList() *LnsList{ return self.accessSymList }
// 150: DeclConstr
func (self *TransUnit_TentativeSymbol) InitTransUnit_TentativeSymbol(parent LnsAny,scope *Ast_Scope,moduleScope *Ast_Scope,loopFlag bool,refAccessSymPosList LnsAny) {
    self.symbolSet = NewLnsSet([]LnsAny{})
    
    self.oldSymbolSet = nil
    
    self.parent = parent
    
    self.scope = scope
    
    self.skipFlag = false
    
    self.loopFlag = loopFlag
    
    self.accessSymList = NewLnsList([]LnsAny{})
    
    self.initSymSet = NewLnsSet([]LnsAny{})
    
    self.accessSymSet = NewLnsSet([]LnsAny{})
    
    self.modSymbolSet = NewLnsSet([]LnsAny{})
    
    var list *LnsList
    if refAccessSymPosList != nil{
        refAccessSymPosList_5523 := refAccessSymPosList.(*LnsList)
        list = refAccessSymPosList_5523
        
    } else {
        if loopFlag{
            list = TransUnit_clearThePosForModToRef_1103_(scope, moduleScope)
            
        } else { 
            list = NewLnsList([]LnsAny{})
            
        }
    }
    self.sym2posForModToRef = list
    
}

// 179: decl @lune.@base.@TransUnit.TentativeSymbol.modSym
func (self *TransUnit_TentativeSymbol) ModSym(accessSym *Ast_SymbolInfo) {
    self.modSymbolSet.Add(Ast_SymbolInfo2Stem(accessSym.FP.GetOrg()))
}

// 182: decl @lune.@base.@TransUnit.TentativeSymbol.addAccessSym
func (self *TransUnit_TentativeSymbol) AddAccessSym(accessSym *Ast_SymbolInfo) {
    self.accessSymSet.Add(Ast_SymbolInfo2Stem(accessSym.FP.GetOrg()))
}

// 191: decl @lune.@base.@TransUnit.TentativeSymbol.addAccessSymPos
func (self *TransUnit_TentativeSymbol) AddAccessSymPos(accessSym *TransUnit_AccessSymPos) {
    self.accessSymList.Insert(TransUnit_AccessSymPos2Stem(accessSym))
}

// 197: decl @lune.@base.@TransUnit.TentativeSymbol.clearAccessSym
func (self *TransUnit_TentativeSymbol) ClearAccessSym() {
    if self.accessSymList.Len() != 0{
        self.accessSymList = NewLnsList([]LnsAny{})
        
    }
}

// 209: decl @lune.@base.@TransUnit.TentativeSymbol.checkAndExclude
func (self *TransUnit_TentativeSymbol) CheckAndExclude(symbolInfo *Ast_SymbolInfo) bool {
    symbolInfo = symbolInfo.FP.GetOrg()
    
    if self.symbolSet.Has(Ast_SymbolInfo2Stem(symbolInfo)){
        self.symbolSet.Del(Ast_SymbolInfo2Stem(symbolInfo))
        return true
    }
    return false
}

// 221: decl @lune.@base.@TransUnit.TentativeSymbol.regist
func (self *TransUnit_TentativeSymbol) Regist(symbolInfo *Ast_SymbolInfo,pos *Types_Position) bool {
    self.symbolSet.Add(Ast_SymbolInfo2Stem(symbolInfo.FP.GetOrg()))
    self.initSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo.FP.GetOrg()))
    symbolInfo.FP.Set_hasValueFlag(true)
    if self.FP.Get_scope().FP.IsInnerOf(symbolInfo.FP.Get_scope()){
        if Lns_op_not(symbolInfo.FP.Get_mutable()){
            var work *TransUnit_TentativeSymbol
            work = self
            for  {
                {
                    __exp := work.parent
                    if __exp != nil {
                        _exp := __exp.(*TransUnit_TentativeSymbol)
                        if work.scope == symbolInfo.FP.Get_scope(){
                            break
                        }
                        if work.loopFlag{
                            return false
                        }
                        work = _exp
                        
                    } else {
                        break
                    }
                }
            }
        }
    }
    return true
}

// 255: decl @lune.@base.@TransUnit.TentativeSymbol.skip
func (self *TransUnit_TentativeSymbol) Skip() {
    for _symbolInfo := range( self.symbolSet.Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        symbolInfo.FP.ClearValue()
    }
    self.skipFlag = true
    
}

// 268: decl @lune.@base.@TransUnit.TentativeSymbol.merge
func (self *TransUnit_TentativeSymbol) merge(finishFlag bool) bool {
    if self.skipFlag{
        self.skipFlag = false
        
        {
            _other := self.oldSymbolSet
            if _other != nil {
                other := _other.(*LnsSet)
                self.symbolSet = other.Clone()
                
            }
        }
        if finishFlag{
            for _symbolInfo := range( self.symbolSet.Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                symbolInfo.FP.UpdateValue(symbolInfo.FP.Get_posForLatestMod())
            }
        }
        return self.oldSymbolSet != nil
    }
    {
        _other := self.oldSymbolSet
        if _other != nil {
            other := _other.(*LnsSet)
            var mergedSet *LnsSet
            mergedSet = self.symbolSet.Clone().And(other)
            if finishFlag{
                for _symbolInfo := range( self.symbolSet.Clone().Or(other).Sub(mergedSet).Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.ClearValue()
                }
            } else { 
                for _symbolInfo := range( self.symbolSet.Clone().Or(other).Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.ClearValue()
                }
            }
            self.symbolSet = mergedSet
            
        } else {
            if Lns_op_not(finishFlag){
                for _symbolInfo := range( self.symbolSet.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.ClearValue()
                }
            }
        }
    }
    return true
}

// 305: decl @lune.@base.@TransUnit.TentativeSymbol.syncPos
func (self *TransUnit_TentativeSymbol) SyncPos() {
    if self.loopFlag{
        for _, _info := range( self.sym2posForModToRef.Items ) {
            info := _info.(TransUnit_RefAccessSymPosDownCast).ToTransUnit_RefAccessSymPos()
            var symbol *Ast_SymbolInfo
            symbol = info.FP.Get_symbol()
            if Lns_isCondTrue( symbol.FP.Get_posForModToRef()){
                symbol.FP.Set_posForModToRef(symbol.FP.Get_posForLatestMod())
            } else { 
                symbol.FP.Set_posForModToRef(info.FP.Get_pos())
            }
        }
    }
}

// 330: decl @lune.@base.@TransUnit.TentativeSymbol.finish
func (self *TransUnit_TentativeSymbol) Finish(complete bool) LnsAny {
    self.FP.SyncPos()
    self.FP.merge(true)
    {
        _parent := self.parent
        if _parent != nil {
            parent := _parent.(*TransUnit_TentativeSymbol)
            if complete{
                for _symbolInfo := range( self.symbolSet.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    if symbolInfo.FP.Get_hasValueFlag(){
                        if parent.scope.FP.IsInnerOf(symbolInfo.FP.Get_scope()){
                            parent.symbolSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
                        }
                    }
                }
            } else { 
                for _symbolInfo := range( self.symbolSet.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.ClearValue()
                }
            }
            for _symbolInfo := range( self.initSymSet.Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if symbolInfo.FP.Get_scope() != self.scope{
                    parent.initSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
                }
            }
            
            for _symbolInfo := range( self.accessSymSet.Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if symbolInfo.FP.Get_scope() != self.scope{
                    parent.accessSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
                }
            }
            
            for _symbolInfo := range( self.modSymbolSet.Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if symbolInfo.FP.Get_scope() != self.scope{
                    parent.modSymbolSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
                }
            }
            
            return parent
        }
    }
    return nil
}

// 366: decl @lune.@base.@TransUnit.TentativeSymbol.newSet
func (self *TransUnit_TentativeSymbol) NewSet(scope *Ast_Scope) {
    if self.FP.merge(false){
        self.oldSymbolSet = self.symbolSet
        
    }
    self.symbolSet = NewLnsSet([]LnsAny{})
    
    self.scope = scope
    
}


// declaration Class -- ClosureFun
type TransUnit_ClosureFunMtd interface {
    Check() bool
}
type TransUnit_ClosureFun struct {
    symbol *Ast_SymbolInfo
    fromScope *Ast_Scope
    FP TransUnit_ClosureFunMtd
}
func TransUnit_ClosureFun2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_ClosureFun).FP
}
type TransUnit_ClosureFunDownCast interface {
    ToTransUnit_ClosureFun() *TransUnit_ClosureFun
}
func TransUnit_ClosureFunDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_ClosureFunDownCast)
    if ok { return work.ToTransUnit_ClosureFun() }
    return nil
}
func (obj *TransUnit_ClosureFun) ToTransUnit_ClosureFun() *TransUnit_ClosureFun {
    return obj
}
func NewTransUnit_ClosureFun(arg1 *Ast_SymbolInfo, arg2 *Ast_Scope) *TransUnit_ClosureFun {
    obj := &TransUnit_ClosureFun{}
    obj.FP = obj
    obj.InitTransUnit_ClosureFun(arg1, arg2)
    return obj
}
func (self *TransUnit_ClosureFun) InitTransUnit_ClosureFun(arg1 *Ast_SymbolInfo, arg2 *Ast_Scope) {
    self.symbol = arg1
    self.fromScope = arg2
}
// 401: decl @lune.@base.@TransUnit.ClosureFun.check
func (self *TransUnit_ClosureFun) Check() bool {
    {
        __exp := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.symbol.FP.Get_typeInfo().FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_closureSymList()}))
        if __exp != nil {
            _exp := __exp.(*LnsList)
            if _exp.Len() > 0{
                self.fromScope.FP.SetClosure(self.symbol)
                return true
            }
        }
    }
    return false
}

// 414: decl @lune.@base.@TransUnit.ClosureFun.checkList
func TransUnit_ClosureFun_checkList_1280_(list *LnsList) {
    var workList *LnsList
    workList = list
    var remainList *LnsList
    remainList = NewLnsList([]LnsAny{})
    for  {
        var update bool
        update = false
        for _, _closureFun := range( workList.Items ) {
            closureFun := _closureFun.(TransUnit_ClosureFunDownCast).ToTransUnit_ClosureFun()
            if closureFun.FP.Check(){
                update = true
                
            } else { 
                remainList.Insert(TransUnit_ClosureFun2Stem(closureFun))
            }
        }
        if Lns_op_not(update){
            break
        }
        workList = remainList
        
        remainList = NewLnsList([]LnsAny{})
        
    }
}


// declaration Class -- AccessSymbolSet
type TransUnit_AccessSymbolSetMtd interface {
    Add(arg1 *Ast_SymbolInfo)
    ApplyPos(arg1 LnsAny)
    Get_accessSym2Pos() *LnsMap
}
type TransUnit_AccessSymbolSet struct {
    accessSym2Pos *LnsMap
    FP TransUnit_AccessSymbolSetMtd
}
func TransUnit_AccessSymbolSet2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_AccessSymbolSet).FP
}
type TransUnit_AccessSymbolSetDownCast interface {
    ToTransUnit_AccessSymbolSet() *TransUnit_AccessSymbolSet
}
func TransUnit_AccessSymbolSetDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_AccessSymbolSetDownCast)
    if ok { return work.ToTransUnit_AccessSymbolSet() }
    return nil
}
func (obj *TransUnit_AccessSymbolSet) ToTransUnit_AccessSymbolSet() *TransUnit_AccessSymbolSet {
    return obj
}
func NewTransUnit_AccessSymbolSet() *TransUnit_AccessSymbolSet {
    obj := &TransUnit_AccessSymbolSet{}
    obj.FP = obj
    obj.InitTransUnit_AccessSymbolSet()
    return obj
}
func (self *TransUnit_AccessSymbolSet) Get_accessSym2Pos() *LnsMap{ return self.accessSym2Pos }
// 452: DeclConstr
func (self *TransUnit_AccessSymbolSet) InitTransUnit_AccessSymbolSet() {
    self.accessSym2Pos = NewLnsMap( map[LnsAny]LnsAny{})
    
}

// 455: decl @lune.@base.@TransUnit.AccessSymbolSet.add
func (self *TransUnit_AccessSymbolSet) Add(symbol *Ast_SymbolInfo) {
    self.accessSym2Pos.Set(symbol,symbol.FP.Get_posForLatestMod())
}

// 458: decl @lune.@base.@TransUnit.AccessSymbolSet.applyPos
func (self *TransUnit_AccessSymbolSet) ApplyPos(excludeSymList LnsAny) {
    var set *LnsSet
    
    {
        _set := excludeSymList
        if _set == nil{
            set = NewLnsSet([]LnsAny{})
            
        } else {
            set = _set.(*LnsSet)
        }
    }
    for _symbol, _pos := range( self.accessSym2Pos.Items ) {
        symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        pos := _pos.(Types_PositionDownCast).ToTypes_Position()
        if Lns_op_not(set.Has(Ast_SymbolInfo2Stem(symbol.FP.GetOrg()))){
            symbol.FP.Set_posForModToRef(pos)
        }
    }
}


// declaration Class -- AccessSymbolSetQueue
type TransUnit_AccessSymbolSetQueueMtd interface {
    Add(arg1 *Ast_SymbolInfo)
    GetMap() *LnsMap
    Pop(arg1 LnsAny)
    Push()
}
type TransUnit_AccessSymbolSetQueue struct {
    queue *LnsList
    FP TransUnit_AccessSymbolSetQueueMtd
}
func TransUnit_AccessSymbolSetQueue2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_AccessSymbolSetQueue).FP
}
type TransUnit_AccessSymbolSetQueueDownCast interface {
    ToTransUnit_AccessSymbolSetQueue() *TransUnit_AccessSymbolSetQueue
}
func TransUnit_AccessSymbolSetQueueDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_AccessSymbolSetQueueDownCast)
    if ok { return work.ToTransUnit_AccessSymbolSetQueue() }
    return nil
}
func (obj *TransUnit_AccessSymbolSetQueue) ToTransUnit_AccessSymbolSetQueue() *TransUnit_AccessSymbolSetQueue {
    return obj
}
func NewTransUnit_AccessSymbolSetQueue() *TransUnit_AccessSymbolSetQueue {
    obj := &TransUnit_AccessSymbolSetQueue{}
    obj.FP = obj
    obj.InitTransUnit_AccessSymbolSetQueue()
    return obj
}
// 472: DeclConstr
func (self *TransUnit_AccessSymbolSetQueue) InitTransUnit_AccessSymbolSetQueue() {
    self.queue = NewLnsList([]LnsAny{})
    
}

// 475: decl @lune.@base.@TransUnit.AccessSymbolSetQueue.push
func (self *TransUnit_AccessSymbolSetQueue) Push() {
    self.queue.Insert(TransUnit_AccessSymbolSet2Stem(NewTransUnit_AccessSymbolSet()))
}

// 478: decl @lune.@base.@TransUnit.AccessSymbolSetQueue.pop
func (self *TransUnit_AccessSymbolSetQueue) Pop(symbolList LnsAny) {
    self.queue.GetAt(self.queue.Len()).(TransUnit_AccessSymbolSetDownCast).ToTransUnit_AccessSymbolSet().FP.ApplyPos(symbolList)
    self.queue.Remove(nil)
}

// 482: decl @lune.@base.@TransUnit.AccessSymbolSetQueue.add
func (self *TransUnit_AccessSymbolSetQueue) Add(symbol *Ast_SymbolInfo) {
    self.queue.GetAt(self.queue.Len()).(TransUnit_AccessSymbolSetDownCast).ToTransUnit_AccessSymbolSet().FP.Add(symbol)
}

// 485: decl @lune.@base.@TransUnit.AccessSymbolSetQueue.getMap
func (self *TransUnit_AccessSymbolSetQueue) GetMap() *LnsMap {
    return self.queue.GetAt(self.queue.Len()).(TransUnit_AccessSymbolSetDownCast).ToTransUnit_AccessSymbolSet().FP.Get_accessSym2Pos()
}


type TransUnit_TransUnitIF interface {
        CreateModifier(arg1 *Ast_TypeInfo, arg2 LnsInt) *Ast_TypeInfo
}
func Lns_cast2TransUnit_TransUnitIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(TransUnit_TransUnitIF); ok { 
        return obj
    }
    return nil
}

// declaration Class -- TransUnit
type TransUnit_TransUnitMtd interface {
    MultiTo1(arg1 *Nodes_Node) *Nodes_Node
    accessSymbol(arg1 *Ast_SymbolInfo, arg2 bool)
    addAccessor(arg1 *Nodes_DeclMemberNode, arg2 *LnsSet, arg3 *Ast_Scope, arg4 *Ast_TypeInfo)
    addDefaultConstructor(arg1 *Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope, arg4 *LnsList, arg5 *LnsSet, arg6 bool)
    addErrMess(arg1 *Types_Position, arg2 string)
    addLocalVar(arg1 *Types_Position, arg2 bool, arg3 bool, arg4 string, arg5 *Ast_TypeInfo, arg6 LnsInt, arg7 LnsAny) *Ast_SymbolInfo
    addMethod(arg1 *Ast_TypeInfo, arg2 *Nodes_Node, arg3 string)
    addWarnErrMess(arg1 *Types_Position, arg2 bool, arg3 string)
    addWarnMess(arg1 *Types_Position, arg2 string)
    analyzeAccessClassField(arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 *Types_Token)(*Ast_TypeInfo, LnsAny, bool)
    analyzeAlias(arg1 LnsInt, arg2 *Types_Token) *Nodes_AliasNode
    analyzeApply(arg1 *Types_Token) *Nodes_ApplyNode
    analyzeBlock(arg1 LnsInt, arg2 LnsInt, arg3 LnsAny, arg4 LnsAny) *Nodes_BlockNode
    analyzeClassBody(arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 LnsInt, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 *Types_Token, arg8 LnsAny, arg9 LnsAny, arg10 LnsInt, arg11 *Types_Token)(*Nodes_DeclClassNode, *Types_Token, *LnsSet)
    analyzeDecl(arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token) LnsAny
    analyzeDeclAlge(arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclAlgeNode
    analyzeDeclAlternateType(arg1 bool, arg2 *Types_Token, arg3 LnsInt)(*Types_Token, *LnsList)
    analyzeDeclArgList(arg1 LnsInt, arg2 *Ast_Scope, arg3 *LnsList, arg4 bool) *Types_Token
    analyzeDeclClass(arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 LnsInt) *Nodes_DeclClassNode
    analyzeDeclEnum(arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclEnumNode
    analyzeDeclForm(arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclFormNode
    analyzeDeclFunc(arg1 LnsInt, arg2 bool, arg3 bool, arg4 LnsInt, arg5 bool, arg6 LnsAny, arg7 *Types_Token, arg8 LnsAny) *Nodes_Node
    analyzeDeclMacro(arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclMacroNode
    analyzeDeclMacroSub(arg1 LnsInt, arg2 *Types_Token, arg3 *Types_Token, arg4 *Ast_Scope, arg5 *Ast_TypeInfo, arg6 *LnsList) *Nodes_DeclMacroNode
    analyzeDeclMember(arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool, arg4 *Types_Token) *Nodes_DeclMemberNode
    analyzeDeclMethod(arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool, arg4 bool, arg5 LnsInt, arg6 bool, arg7 *Types_Token, arg8 *Types_Token) *Nodes_Node
    analyzeDeclProto(arg1 LnsInt, arg2 *Types_Token) *Nodes_Node
    analyzeDeclVar(arg1 LnsInt, arg2 LnsInt, arg3 *Types_Token) *Nodes_Node
    analyzeEnvLock(arg1 *Types_Token) *Nodes_Node
    analyzeExp(arg1 bool, arg2 bool, arg3 bool, arg4 LnsAny, arg5 LnsAny) *Nodes_Node
    analyzeExpCall(arg1 *Types_Token, arg2 *Nodes_Node, arg3 *Types_Token)(*Nodes_Node, *Types_Token)
    analyzeExpCast(arg1 *Types_Token, arg2 string, arg3 *Nodes_Node) *Nodes_Node
    analyzeExpCont(arg1 *Types_Token, arg2 *Nodes_Node, arg3 bool, arg4 bool) *Nodes_Node
    analyzeExpField(arg1 *Types_Token, arg2 *Types_Token, arg3 LnsInt, arg4 *Nodes_Node) *Nodes_Node
    analyzeExpList(arg1 bool, arg2 bool, arg3 bool, arg4 LnsAny, arg5 LnsAny, arg6 LnsAny) *Nodes_ExpListNode
    analyzeExpMacroStat(arg1 *Types_Token) *Nodes_ExpMacroStatNode
    analyzeExpOneRVal(arg1 bool, arg2 bool, arg3 LnsAny, arg4 LnsAny) *Nodes_Node
    analyzeExpOp2(arg1 *Types_Token, arg2 *Nodes_Node, arg3 LnsAny) *Nodes_Node
    analyzeExpOpEquals(arg1 *Types_Position, arg2 *Types_Token, arg3 *Nodes_Node, arg4 *Nodes_Node)(*Nodes_Node, *Nodes_Node)
    analyzeExpOpSet(arg1 *Nodes_Node, arg2 *Types_Token, arg3 *LnsList) *Nodes_Node
    analyzeExpRefItem(arg1 *Types_Token, arg2 *Nodes_Node, arg3 bool) *Nodes_Node
    analyzeExpSymbol(arg1 *Types_Token, arg2 *Types_Token, arg3 LnsInt, arg4 LnsAny, arg5 bool, arg6 bool) *Nodes_Node
    analyzeExpUnwrap(arg1 *Types_Token) *Nodes_Node
    analyzeExtend(arg1 LnsInt, arg2 *Types_Position)(*Types_Token, LnsAny, *LnsList, *LnsMap)
    analyzeFor(arg1 *Types_Token) *Nodes_ForNode
    analyzeForeach(arg1 *Types_Token, arg2 bool) *Nodes_Node
    analyzeFuncBlock(arg1 LnsInt, arg2 *Types_Token, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 string, arg6 *Ast_Scope, arg7 *LnsList) *Nodes_BlockNode
    analyzeIf(arg1 *Types_Token) *Nodes_Node
    analyzeIfUnwrap(arg1 *Types_Token) *Nodes_IfUnwrapNode
    analyzeImport(arg1 *Types_Token) *Nodes_Node
    analyzeImportFor(arg1 *Types_Position, arg2 string, arg3 string, arg4 bool, arg5 LnsInt) *Nodes_Node
    analyzeInitExp(arg1 *Types_Position, arg2 LnsInt, arg3 bool, arg4 *LnsList, arg5 *LnsList)(*LnsList, *LnsList, *LnsList, LnsAny)
    analyzeLetAndInitExp(arg1 *Types_Position, arg2 bool, arg3 LnsInt, arg4 LnsInt, arg5 bool)(*LnsList, *LnsList, *LnsList, LnsAny)
    analyzeListConst(arg1 *Types_Token, arg2 LnsAny) *Nodes_Node
    analyzeListItems(arg1 *Types_Position, arg2 *Types_Token, arg3 string, arg4 LnsAny)(LnsAny, *Ast_TypeInfo)
    analyzeLuneControl(arg1 *Types_Token) LnsAny
    analyzeMapConst(arg1 *Types_Token, arg2 LnsAny) *Nodes_LiteralMapNode
    analyzeMatch(arg1 *Types_Token) *Nodes_MatchNode
    analyzeNewAlge(arg1 *Types_Token, arg2 *Ast_AlgeTypeInfo, arg3 LnsAny) *Nodes_NewAlgeValNode
    analyzeProvide(arg1 *Types_Token) *Nodes_ProvideNode
    analyzePushClass(arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsInt, arg9 *LnsList)(*Types_Token, *Ast_TypeInfo)
    analyzeRefType(arg1 LnsInt, arg2 bool, arg3 bool) *Nodes_RefTypeNode
    analyzeRefTypeWithSymbol(arg1 LnsInt, arg2 bool, arg3 bool, arg4 bool, arg5 *Nodes_Node, arg6 bool) *Nodes_RefTypeNode
    analyzeRepeat(arg1 *Types_Token) *Nodes_RepeatNode
    analyzeRetTypeList(arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 bool)(*LnsList, *Types_Token)
    analyzeReturn(arg1 *Types_Token) *Nodes_ReturnNode
    analyzeScope(arg1 *Types_Token) *Nodes_ScopeNode
    analyzeSetConst(arg1 *Types_Token, arg2 LnsAny) *Nodes_Node
    analyzeStatement(arg1 LnsAny) LnsAny
    analyzeStatementList(arg1 *LnsList, arg2 LnsAny)(LnsAny, LnsInt)
    analyzeStatementListSubfile(arg1 *LnsList) LnsAny
    analyzeStrConst(arg1 *Types_Token, arg2 *Types_Token) *Nodes_Node
    analyzeSubfile(arg1 *Types_Token) *Nodes_SubfileNode
    analyzeSuper(arg1 *Types_Token) *Nodes_Node
    analyzeSwitch(arg1 *Types_Token) *Nodes_SwitchNode
    analyzeTest(arg1 *Types_Token) *Nodes_Node
    analyzeTestCase(arg1 *Types_Token) *Nodes_TestCaseNode
    analyzeUnwrap(arg1 *Types_Token) *Nodes_Node
    analyzeWhen(arg1 *Types_Token) *Nodes_Node
    analyzeWhile(arg1 *Types_Token) *Nodes_WhileNode
    checkAccesSym()
    checkAlgeComp(arg1 *Types_Token, arg2 *Ast_AlgeTypeInfo)
    checkArgForSort(arg1 *Types_Token, arg2 *LnsList, arg3 *Nodes_ExpListNode)
    checkArgForStringForm(arg1 *Types_Token, arg2 *Nodes_ExpListNode)
    checkComp(arg1 *Types_Token, arg2 TransUnit_checkCompForm_6433_)
    checkEnumComp(arg1 *Types_Token, arg2 *Ast_EnumTypeInfo)
    checkFieldComp(arg1 bool, arg2 *Types_Token, arg3 *Nodes_Node)
    checkImplicitCast(arg1 *LnsMap, arg2 bool, arg3 *LnsList, arg4 *Nodes_ExpListNode, arg5 TransUnit_checkImplicitCastCallback_1263_) LnsAny
    checkLiteralEmptyCollection(arg1 *Types_Position, arg2 string, arg3 *Ast_TypeInfo)
    checkMatchType(arg1 string, arg2 *Types_Position, arg3 *LnsList, arg4 LnsAny, arg5 bool, arg6 bool, arg7 LnsAny)(LnsInt, *LnsMap, LnsAny, *LnsList)
    checkMatchValType(arg1 *Types_Position, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 *LnsList, arg5 LnsAny)(LnsInt, *LnsMap, LnsAny)
    checkNextToken(arg1 string) *Types_Token
    checkOverrideMethod(arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo) *LnsList
    checkOverriededMethodOfAllClass()
    checkPublic(arg1 *Types_Position, arg2 *Ast_TypeInfo)
    checkShadowing(arg1 *Types_Position, arg2 string, arg3 *Ast_Scope)
    checkStringFormat(arg1 *Types_Position, arg2 string, arg3 *LnsList)
    checkSymbol(arg1 *Types_Token, arg2 LnsInt) *Types_Token
    checkSymbolComp(arg1 *Types_Token)
    checkSymbolHavingValue(arg1 *Types_Position, arg2 *LnsList)
    checkThreading(arg1 *Types_Position) bool
    checkToken(arg1 *Types_Token, arg2 string) *Types_Token
    convToExtTypeList(arg1 *Types_Position, arg2 *Ast_TypeInfo, arg3 *LnsList) *LnsList
    CreateAST(arg1 *Parser_Parser, arg2 bool, arg3 LnsAny) *TransUnit_ASTInfo
    createExpList(arg1 *Types_Position, arg2 *LnsList, arg3 *LnsList, arg4 bool, arg5 LnsAny) *Nodes_ExpListNode
    createExpListNode(arg1 *Nodes_ExpListNode, arg2 *LnsList) *Nodes_ExpListNode
    createExtType(arg1 *Types_Position, arg2 *Ast_TypeInfo) *Ast_TypeInfo
    CreateModifier(arg1 *Ast_TypeInfo, arg2 LnsInt) *Ast_TypeInfo
    createNoneNode(arg1 *Types_Position) *Nodes_Node
    CreatePosition(arg1 LnsInt, arg2 LnsInt) *Types_Position
    dumpComp(arg1 Writer_Writer, arg2 string, arg3 *Ast_SymbolInfo, arg4 bool) bool
    dumpFieldComp(arg1 Writer_Writer, arg2 bool, arg3 *Ast_TypeInfo, arg4 string, arg5 LnsAny)
    dumpSymbolComp(arg1 Writer_Writer, arg2 *Ast_Scope, arg3 string)
    dumpSymbolType(arg1 string, arg2 *Ast_TypeInfo)
    Error(arg1 string)
    ErrorAt(arg1 *Types_Position, arg2 string)
    errorShadowing(arg1 *Types_Position, arg2 LnsAny)
    evalMacro(arg1 *Types_Token, arg2 *Ast_TypeInfo, arg3 LnsAny) *Nodes_ExpMacroExpNode
    evalMacroOp(arg1 *Types_Token, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 Macro_EvalMacroCallback)
    finishTentativeSymbol(arg1 bool)
    getContinueToken()(*Types_Token, bool)
    getCurrentClass() *Ast_TypeInfo
    getCurrentNamespaceScope() *Ast_Scope
    getCurrentNamespaceTypeInfo() *Ast_TypeInfo
    GetStreamName() string
    getSymbolToken(arg1 LnsInt) *Types_Token
    getToken(arg1 LnsAny) *Types_Token
    GetTokenNoErr() *Types_Token
    Get_errMessList() *LnsList
    Get_warnMessList() *LnsList
    inAnalyzingState(arg1 LnsInt) bool
    isTargetToken(arg1 *Types_Token) bool
    isTargetTokenPos(arg1 string, arg2 *Types_Position) bool
    mergeTentativeSymbol(arg1 *Ast_Scope)
    NewPushback(arg1 *LnsList)
    popAnalyzingState()
    popClass()
    popModule()
    popScope()
    prepareExpCall(arg1 *Types_Position, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 *Ast_TypeInfo)(*LnsMap, LnsAny)
    prepareTentativeSymbol(arg1 *Ast_Scope, arg2 bool, arg3 LnsAny)
    processAddFunc(arg1 bool, arg2 *Ast_Scope, arg3 *Types_Token, arg4 *Ast_TypeInfo, arg5 *LnsMap) *Ast_SymbolInfo
    processCaseDefault(arg1 *Types_Token, arg2 LnsInt, arg3 *Types_Token, arg4 bool)(LnsAny, bool)
    processFunc(arg1 *Types_Token, arg2 *Types_Token, arg3 LnsAny, arg4 *Nodes_Node, arg5 *Ast_TypeInfo, arg6 *LnsMap, arg7 *LnsList, arg8 *Ast_TypeInfo, arg9 LnsAny) *Nodes_Node
    processImport(arg1 string, arg2 LnsInt)(*Lns_luaValue, *LnsMap, *Nodes_ModuleInfo)
    pushAnalyzingState(arg1 LnsInt)
    pushClass(arg1 *Types_Position, arg2 LnsInt, arg3 bool, arg4 LnsAny, arg5 LnsAny, arg6 LnsAny, arg7 bool, arg8 string, arg9 bool, arg10 LnsInt, arg11 LnsAny) *Ast_TypeInfo
    pushClassScope(arg1 *Types_Position, arg2 *Ast_TypeInfo)
    pushExtModule(arg1 bool, arg2 string, arg3 LnsInt, arg4 *Types_Position, arg5 bool, arg6 LnsInt, arg7 string) *Ast_TypeInfo
    pushModule(arg1 bool, arg2 string, arg3 bool) *Ast_TypeInfo
    pushScope(arg1 bool, arg2 LnsAny, arg3 LnsAny) *Ast_Scope
    Pushback()
    PushbackStr(arg1 string, arg2 string)
    PushbackToken(arg1 *Types_Token)
    registBuiltInScope()
    supportLang(arg1 string) bool
}
type TransUnit_TransUnit struct {
    analyzingStateQueue *LnsList
    importModuleInfo *FrontInterface_ImportModuleInfo
    validMutControl bool
    moduleName string
    moduleType *Ast_TypeInfo
    globalScope *Ast_Scope
    topScope *Ast_Scope
    moduleScope *Ast_Scope
    macroScope LnsAny
    scope *Ast_Scope
    tentativeSymbol *TransUnit_TentativeSymbol
    typeId2ClassMap *LnsMap
    typeInfo2ClassNode *LnsMap
    parser *Parser_DefaultPushbackParser
    commentCtrl *Parser_CommentCtrl
    errMessList *LnsList
    warnMessList *LnsList
    importModule2ModuleInfo *LnsMap
    importModuleName2ModuleInfo *LnsMap
    importModuleSet *LnsSet
    importedAliasMap *LnsMap
    subfileList *LnsList
    has__func__Symbol *LnsSet
    analyzeMode LnsInt
    analyzePos *Types_Position
    analyzeModule string
    helperInfo *Nodes_LuneHelperInfo
    provideNode LnsAny
    nodeManager *Nodes_NodeManager
    loopScopeQueue *LnsList
    protoFuncMap *LnsMap
    protoClassMap *LnsMap
    targetLuaVer *LuaVer_LuaVerInfo
    moduleId *FrontInterface_ModuleId
    ignoreToCheckSymbol_ bool
    ctrl_info *Types_TransCtrlInfo
    macroCtrl *Macro_MacroCtrl
    typeNameCtrl *Ast_TypeNameCtrl
    scopeAccess LnsInt
    closureFunList *LnsList
    advertisedTypeSet *LnsSet
    accessSymbolSetQueue *TransUnit_AccessSymbolSetQueue
    processInfo *Ast_ProcessInfo
    FP TransUnit_TransUnitMtd
}
func TransUnit_TransUnit2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_TransUnit).FP
}
type TransUnit_TransUnitDownCast interface {
    ToTransUnit_TransUnit() *TransUnit_TransUnit
}
func TransUnit_TransUnitDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_TransUnitDownCast)
    if ok { return work.ToTransUnit_TransUnit() }
    return nil
}
func (obj *TransUnit_TransUnit) ToTransUnit_TransUnit() *TransUnit_TransUnit {
    return obj
}
func NewTransUnit_TransUnit(arg1 *FrontInterface_ModuleId, arg2 *FrontInterface_ImportModuleInfo, arg3 *Nodes_MacroEval, arg4 LnsAny, arg5 LnsAny, arg6 LnsAny, arg7 *LuaVer_LuaVerInfo, arg8 *Types_TransCtrlInfo) *TransUnit_TransUnit {
    obj := &TransUnit_TransUnit{}
    obj.FP = obj
    obj.InitTransUnit_TransUnit(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *TransUnit_TransUnit) Get_errMessList() *LnsList{ return self.errMessList }
func (self *TransUnit_TransUnit) Get_warnMessList() *LnsList{ return self.warnMessList }
// 624: DeclConstr
func (self *TransUnit_TransUnit) InitTransUnit_TransUnit(moduleId *FrontInterface_ModuleId,importModuleInfo *FrontInterface_ImportModuleInfo,macroEval *Nodes_MacroEval,analyzeModule LnsAny,mode LnsAny,pos LnsAny,targetLuaVer *LuaVer_LuaVerInfo,ctrl_info *Types_TransCtrlInfo) {
    self.importedAliasMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.importModuleSet = NewLnsSet([]LnsAny{})
    
    self.processInfo = Ast_createProcessInfo(ctrl_info.ValidLuaval)
    
    self.accessSymbolSetQueue = NewTransUnit_AccessSymbolSetQueue()
    
    self.advertisedTypeSet = NewLnsSet([]LnsAny{})
    
    self.closureFunList = NewLnsList([]LnsAny{})
    
    self.scopeAccess = Ast_ScopeAccess__Normal
    
    self.macroCtrl = NewMacro_MacroCtrl(macroEval)
    
    self.typeNameCtrl = Ast_defaultTypeNameCtrl
    
    self.protoClassMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.analyzingStateQueue = NewLnsList([]LnsAny{})
    
    self.ctrl_info = ctrl_info
    
    self.ignoreToCheckSymbol_ = false
    
    self.moduleId = moduleId
    
    self.helperInfo = NewNodes_LuneHelperInfo()
    
    self.targetLuaVer = targetLuaVer
    
    self.importModuleInfo = importModuleInfo
    
    self.protoFuncMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.loopScopeQueue = NewLnsList([]LnsAny{})
    
    self.has__func__Symbol = NewLnsSet([]LnsAny{})
    
    self.nodeManager = NewNodes_NodeManager()
    
    self.importModuleName2ModuleInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.importModule2ModuleInfo = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.macroScope = nil
    
    self.validMutControl = true
    
    self.moduleName = ""
    
    self.moduleType = Ast_headTypeInfo
    
    self.parser = NewParser_DefaultPushbackParser(&NewParser_DummyParser().Parser_Parser)
    
    self.subfileList = NewLnsList([]LnsAny{})
    
    self.globalScope = NewAst_Scope(self.processInfo, Ast_rootScope, false, nil, nil)
    
    self.scope = NewAst_Scope(self.processInfo, self.globalScope, true, nil, nil)
    
    self.topScope = self.scope
    
    self.moduleScope = self.scope
    
    self.tentativeSymbol = NewTransUnit_TentativeSymbol(nil, self.globalScope, self.moduleScope, false, nil)
    
    self.typeId2ClassMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.typeInfo2ClassNode = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.commentCtrl = NewParser_CommentCtrl()
    
    self.errMessList = NewLnsList([]LnsAny{})
    
    self.warnMessList = NewLnsList([]LnsAny{})
    
    self.analyzeMode = Lns_unwrapDefault( mode, TransUnit_AnalyzeMode__Compile).(LnsInt)
    
    self.analyzePos = Lns_unwrapDefault( pos, self.FP.CreatePosition(0, 0)).(*Types_Position)
    
    self.analyzeModule = Lns_unwrapDefault( analyzeModule, "").(string)
    
    self.provideNode = nil
    
}



// 691: decl @lune.@base.@TransUnit.TransUnit.pushAnalyzingState
func (self *TransUnit_TransUnit) pushAnalyzingState(state LnsInt) {
    self.analyzingStateQueue.Insert(state)
}

// 695: decl @lune.@base.@TransUnit.TransUnit.popAnalyzingState
func (self *TransUnit_TransUnit) popAnalyzingState() {
    if self.analyzingStateQueue.Len() == 0{
        self.FP.Error("underflow analyzingStateQueue")
    }
    self.analyzingStateQueue.Remove(nil)
}

// 702: decl @lune.@base.@TransUnit.TransUnit.inAnalyzingState
func (self *TransUnit_TransUnit) inAnalyzingState(state LnsInt) bool {
    if self.analyzingStateQueue.Len() > 0{
        return self.analyzingStateQueue.GetAt(self.analyzingStateQueue.Len()).(LnsInt) == state
    }
    return false
}

// 714: decl @lune.@base.@TransUnit.TransUnit.addErrMess
func (self *TransUnit_TransUnit) addErrMess(pos *Types_Position,mess string) {
    self.errMessList.Insert(Lns_getVM().String_format("%s:%d:%d: error: %s", []LnsAny{pos.StreamName, pos.LineNo, pos.Column, mess}))
}

// 720: decl @lune.@base.@TransUnit.TransUnit.addWarnMess
func (self *TransUnit_TransUnit) addWarnMess(pos *Types_Position,mess string) {
    self.warnMessList.Insert(Lns_getVM().String_format("%s:%d:%d: warning: %s", []LnsAny{self.parser.FP.GetStreamName(), pos.LineNo, pos.Column, mess}))
}

// 726: decl @lune.@base.@TransUnit.TransUnit.addWarnErrMess
func (self *TransUnit_TransUnit) addWarnErrMess(pos *Types_Position,err bool,mess string) {
    if err{
        self.FP.addErrMess(pos, mess)
    } else { 
        self.FP.addWarnMess(pos, mess)
    }
}

// 739: decl @lune.@base.@TransUnit.TransUnit.pushScope
func (self *TransUnit_TransUnit) pushScope(classFlag bool,baseInfo LnsAny,interfaceList LnsAny) *Ast_Scope {
    self.scope = Ast_TypeInfo_createScope(self.processInfo, self.scope, classFlag, baseInfo, interfaceList)
    
    return self.scope
}

// 748: decl @lune.@base.@TransUnit.TransUnit.popScope
func (self *TransUnit_TransUnit) popScope() {
    self.scope = self.scope.FP.Get_parent()
    
}

// 752: decl @lune.@base.@TransUnit.TransUnit.prepareTentativeSymbol
func (self *TransUnit_TransUnit) prepareTentativeSymbol(scope *Ast_Scope,loopFlag bool,list LnsAny) {
    self.tentativeSymbol = NewTransUnit_TentativeSymbol(self.tentativeSymbol, scope, self.moduleScope, loopFlag, list)
    
}

// 759: decl @lune.@base.@TransUnit.TransUnit.checkAccesSym
func (self *TransUnit_TransUnit) checkAccesSym() {
    for _, _accessSym := range( self.tentativeSymbol.FP.Get_accessSymList().Items ) {
        accessSym := _accessSym.(TransUnit_AccessSymPosDownCast).ToTransUnit_AccessSymPos()
        var symbolInfo *Ast_SymbolInfo
        symbolInfo = accessSym.FP.Get_symbol()
        if Lns_op_not(symbolInfo.FP.Get_hasValueFlag()){
            self.FP.addErrMess(accessSym.FP.Get_pos(), Lns_getVM().String_format("This can't access variable have no value -- %s", []LnsAny{symbolInfo.FP.Get_name()}))
        }
    }
    self.tentativeSymbol.FP.ClearAccessSym()
}

// 772: decl @lune.@base.@TransUnit.TransUnit.finishTentativeSymbol
func (self *TransUnit_TransUnit) finishTentativeSymbol(complete bool) {
    self.FP.checkAccesSym()
    var tentativeSymbol *TransUnit_TentativeSymbol
    tentativeSymbol = self.tentativeSymbol
    self.tentativeSymbol = Lns_unwrap( tentativeSymbol.FP.Finish(complete)).(*TransUnit_TentativeSymbol)
    
    {
        var errSymMap *LnsMap
        errSymMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _symbolInfo := range( tentativeSymbol.FP.Get_initSymSet().Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if tentativeSymbol.FP.Get_scope().FP.Get_parent() == symbolInfo.FP.Get_scope(){
                if Lns_op_not(symbolInfo.FP.Get_hasValueFlag()){
                    errSymMap.Set(symbolInfo.FP.Get_name(),symbolInfo)
                }
            }
        }
        {
            __collection2618 := errSymMap
            __sorted2618 := __collection2618.CreateKeyListStr()
            __sorted2618.Sort( LnsItemKindStr, nil )
            for _, ___key2618 := range( __sorted2618.Items ) {
                symbolInfo := __collection2618.Items[ ___key2618 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.addErrMess(self.parser.FP.GetLastPos(), Lns_getVM().String_format("There is the case no initialized value for '%s'", []LnsAny{symbolInfo.FP.Get_name()}))
            }
        }
    }
    if tentativeSymbol.FP.Get_scope().FP.Get_validCheckingUnaccess(){
        {
            __collection2714 := tentativeSymbol.FP.Get_scope().FP.Get_symbol2SymbolInfoMap()
            __sorted2714 := __collection2714.CreateKeyListStr()
            __sorted2714.Sort( LnsItemKindStr, nil )
            for _, ___key2714 := range( __sorted2714.Items ) {
                symbolInfo := __collection2714.Items[ ___key2714 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(symbolInfo.FP.Get_posForModToRef())) ||
                        Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_posForModToRef() != symbolInfo.FP.Get_posForLatestMod()) ).(bool))) &&
                    Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Var) ||
                        Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Fun) ).(bool))) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(symbolInfo.FP.Get_accessMode()))) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Lns_getVM().String_find(symbolInfo.FP.Get_name(),"^_", nil, nil)))) ).(bool)){
                    {
                        _pos := Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_posForLatestMod()) ||
                            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_pos()) )
                        if _pos != nil {
                            pos := _pos.(*Types_Position)
                            self.FP.addWarnMess(pos, Lns_getVM().String_format("'%s' var isn't accessed", []LnsAny{symbolInfo.FP.Get_name()}))
                        }
                    }
                }
            }
        }
    }
}

// 810: decl @lune.@base.@TransUnit.TransUnit.mergeTentativeSymbol
func (self *TransUnit_TransUnit) mergeTentativeSymbol(scope *Ast_Scope) {
    self.FP.checkAccesSym()
    self.tentativeSymbol.FP.NewSet(scope)
}

// 816: decl @lune.@base.@TransUnit.TransUnit.getCurrentClass
func (self *TransUnit_TransUnit) getCurrentClass() *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var scope *Ast_Scope
    scope = self.scope
    for {
        {
            __exp := scope.FP.Get_ownerTypeInfo()
            if __exp != nil {
                _exp := __exp.(*Ast_TypeInfo)
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
                    Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__Module) ||
                    Lns_GetEnv().SetStackVal( _exp.FP.Get_kind() == Ast_TypeInfoKind__IF) ).(bool){
                    return _exp
                }
            }
        }
        scope = scope.FP.Get_parent()
        
        if scope.FP.IsRoot(){ break }
    }
    return typeInfo
}

// 830: decl @lune.@base.@TransUnit.TransUnit.getCurrentNamespaceTypeInfo
func (self *TransUnit_TransUnit) getCurrentNamespaceTypeInfo() *Ast_TypeInfo {
    return self.scope.FP.GetNamespaceTypeInfo()
}

// 843: decl @lune.@base.@TransUnit.TransUnit.getCurrentNamespaceScope
func (self *TransUnit_TransUnit) getCurrentNamespaceScope() *Ast_Scope {
    return Lns_unwrap( self.FP.getCurrentNamespaceTypeInfo().FP.Get_scope()).(*Ast_Scope)
}

// 847: decl @lune.@base.@TransUnit.TransUnit.pushModule
func (self *TransUnit_TransUnit) pushModule(externalFlag bool,name string,mutable bool) *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var modName string
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(name,"^@", nil, nil))){
        modName = name
        
    } else { 
        modName = Lns_getVM().String_format("@%s", []LnsAny{name})
        
    }
    {
        __exp := self.scope.FP.GetTypeInfoChild(modName)
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            typeInfo = _exp
            
            self.scope = Lns_unwrap( Ast_getScope(typeInfo)).(*Ast_Scope)
            
        } else {
            var parentInfo *Ast_TypeInfo
            parentInfo = self.FP.getCurrentNamespaceTypeInfo()
            var parentScope *Ast_Scope
            parentScope = self.scope
            var scope *Ast_Scope
            scope = self.FP.pushScope(true, nil, nil)
            typeInfo = self.processInfo.FP.CreateModule(scope, parentInfo, externalFlag, modName, mutable)
            
            var existSym LnsAny
            _,existSym = parentScope.FP.AddClass(self.processInfo, modName, nil, typeInfo)
            if existSym != nil{
                existSym_5922 := existSym.(*Ast_SymbolInfo)
                self.FP.addErrMess(self.parser.FP.GetLastPos(), Lns_getVM().String_format("module symbols exist -- %s.%s -- %s.%s", []LnsAny{existSym_5922.FP.Get_namespaceTypeInfo().FP.GetFullName(self.typeNameCtrl, parentScope.FP, false), existSym_5922.FP.Get_name(), parentInfo.FP.GetFullName(self.typeNameCtrl, parentScope.FP, false), modName}))
            }
        }
    }
    if Lns_op_not(self.typeId2ClassMap.Items[typeInfo.FP.Get_typeId()]){
        var namespace *Nodes_NamespaceInfo
        namespace = NewNodes_NamespaceInfo(modName, self.scope, typeInfo)
        self.typeId2ClassMap.Set(typeInfo.FP.Get_typeId(),namespace)
    }
    return typeInfo
}

// 890: decl @lune.@base.@TransUnit.TransUnit.popModule
func (self *TransUnit_TransUnit) popModule() {
    self.FP.popScope()
}

// 894: decl @lune.@base.@TransUnit.TransUnit.pushExtModule
func (self *TransUnit_TransUnit) pushExtModule(externalFlag bool,name string,accessMode LnsInt,pos *Types_Position,lazy bool,lang LnsInt,requirePath string) *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var modName string
    modName = name
    {
        __exp := self.scope.FP.GetTypeInfoChild(modName)
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            self.FP.addErrMess(pos, Lns_getVM().String_format("multiple define -- %s", []LnsAny{name}))
            return _exp
        } else {
            var parentInfo *Ast_TypeInfo
            parentInfo = self.FP.getCurrentNamespaceTypeInfo()
            var parentScope *Ast_Scope
            parentScope = self.scope
            var scope *Ast_Scope
            scope = self.FP.pushScope(true, nil, nil)
            typeInfo = self.processInfo.FP.CreateExtModule(scope, parentInfo, externalFlag, accessMode, name, lang, requirePath)
            
            parentScope.FP.AddExtModule(self.processInfo, name, pos, typeInfo, lazy, lang)
        }
    }
    if Lns_op_not(self.typeId2ClassMap.Items[typeInfo.FP.Get_typeId()]){
        var namespace *Nodes_NamespaceInfo
        namespace = NewNodes_NamespaceInfo(modName, self.scope, typeInfo)
        self.typeId2ClassMap.Set(typeInfo.FP.Get_typeId(),namespace)
    }
    return typeInfo
}

// 926: decl @lune.@base.@TransUnit.TransUnit.pushClassScope
func (self *TransUnit_TransUnit) pushClassScope(errPos *Types_Position,classTypeInfo *Ast_TypeInfo) {
    if self.scope != Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(classTypeInfo.FP.Get_scope()) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_parent()})){
        var classParentName string
        var classParentTypeId LnsInt
        {
            _parentType := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(classTypeInfo.FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_parent()})&&
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_ownerTypeInfo()}))
            if _parentType != nil {
                parentType := _parentType.(*Ast_TypeInfo)
                classParentName = parentType.FP.GetFullName(self.typeNameCtrl, Lns_unwrap( parentType.FP.Get_scope()).(*Ast_Scope).FP, false)
                
                classParentTypeId = parentType.FP.Get_typeId()
                
            } else {
                classParentName = "nil"
                
                classParentTypeId = -1
                
            }
        }
        self.FP.addErrMess(errPos, Lns_getVM().String_format("This class does not exist in this scope. -- %s -- %s(%d), %s(%d)", []LnsAny{classTypeInfo.FP.GetTxt(nil, nil, nil), Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.scope.FP.Get_ownerTypeInfo()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.GetFullName(self.typeNameCtrl, self.scope.FP, false)})/* 942:15 */), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.scope.FP.Get_ownerTypeInfo()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_typeId()}))) ||
            Lns_GetEnv().SetStackVal( -1) ).(LnsInt), classParentName, classParentTypeId}))
    }
    self.scope = Lns_unwrap( Ast_getScope(classTypeInfo)).(*Ast_Scope)
    
}

// 955: decl @lune.@base.@TransUnit.TransUnit.pushClass
func (self *TransUnit_TransUnit) pushClass(errPos *Types_Position,mode LnsInt,abstractFlag bool,baseInfo LnsAny,interfaceList LnsAny,genTypeList LnsAny,externalFlag bool,name string,allowMultiple bool,accessMode LnsInt,defNamespace LnsAny) *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    {
        __exp := self.scope.FP.GetTypeInfo(name, self.scope, true, Ast_ScopeAccess__Normal)
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            typeInfo = _exp
            
        }
    }
    if typeInfo != Ast_headTypeInfo{
        if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(typeInfo.FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_parent()})) != self.scope{
            self.FP.addErrMess(errPos, Lns_getVM().String_format("multiple class(%s)", []LnsAny{typeInfo.FP.GetTxt(self.typeNameCtrl, nil, nil)}))
            self.FP.Error("stop by error")
        }
    }
    if typeInfo != Ast_headTypeInfo{
        if typeInfo.FP.Get_abstractFlag() != abstractFlag{
            self.FP.addErrMess(errPos, Lns_getVM().String_format("mismatch class(%s) abstract for prototpye", []LnsAny{typeInfo.FP.GetTxt(self.typeNameCtrl, nil, nil)}))
        }
        if typeInfo.FP.Get_accessMode() != accessMode{
            self.FP.addErrMess(errPos, Lns_getVM().String_format("mismatch class(%s) accessmode(%s) for prototpye accessmode(%s)", []LnsAny{typeInfo.FP.GetTxt(self.typeNameCtrl, nil, nil), Ast_AccessMode_getTxt( accessMode), Ast_AccessMode_getTxt( typeInfo.FP.Get_accessMode())}))
        }
        if baseInfo != nil{
            baseInfo_5986 := baseInfo.(*Ast_TypeInfo)
            if typeInfo.FP.Get_baseTypeInfo() != baseInfo_5986{
                self.FP.addErrMess(errPos, Lns_getVM().String_format("mismatch class(%s) base class(%s) for prototpye base class(%s)", []LnsAny{typeInfo.FP.GetTxt(self.typeNameCtrl, nil, nil), baseInfo_5986.FP.GetTxt(nil, nil, nil), typeInfo.FP.Get_baseTypeInfo().FP.GetTxt(nil, nil, nil)}))
            }
        } else {
            if typeInfo.FP.HasBase(){
                self.FP.addErrMess(errPos, Lns_getVM().String_format("mismatch class(%s) base class(None) for prototpye base class(%s)", []LnsAny{typeInfo.FP.GetTxt(self.typeNameCtrl, nil, nil), typeInfo.FP.Get_baseTypeInfo().FP.GetTxt(nil, nil, nil)}))
            }
        }
        var compareList func(protoList *LnsList,typeList *LnsList,message string)
        compareList = func(protoList *LnsList,typeList *LnsList,message string) {
            if protoList.Len() == typeList.Len(){
                for _index, _protoType := range( protoList.Items ) {
                    index := _index + 1
                    protoType := _protoType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if protoType != typeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(){
                        self.FP.addErrMess(errPos, Lns_getVM().String_format("mismatch class(%s) %s(%s) for prototpye %s(%s)", []LnsAny{typeInfo.FP.GetTxt(self.typeNameCtrl, nil, nil), message, typeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.GetTxt(self.typeNameCtrl, nil, nil), message, protoType.FP.GetTxt(nil, nil, nil)}))
                    }
                }
            } else { 
                self.FP.addErrMess(errPos, Lns_getVM().String_format("mismatch class(%s) %s(%d) for prototpye %s(%d)", []LnsAny{typeInfo.FP.GetTxt(self.typeNameCtrl, nil, nil), message, typeList.Len(), message, protoList.Len()}))
            }
        }
        compareList(typeInfo.FP.Get_interfaceList(), Lns_unwrapDefault( interfaceList, NewLnsList([]LnsAny{})).(*LnsList), "interface")
        compareList(typeInfo.FP.Get_itemTypeInfoList(), Lns_unwrapDefault( genTypeList, NewLnsList([]LnsAny{})).(*LnsList), "generics")
        self.scope = Lns_unwrap( Ast_getScope(typeInfo)).(*Ast_Scope)
        
        if _switch3961 := (typeInfo.FP.Get_kind()); _switch3961 == Ast_TypeInfoKind__Class {
            if mode == TransUnit_DeclClassMode__Interface{
                self.FP.addErrMess(errPos, Lns_getVM().String_format("define interface already -- %s", []LnsAny{name}))
                Util_printStackTrace()
            }
        } else if _switch3961 == Ast_TypeInfoKind__IF {
            if mode != TransUnit_DeclClassMode__Interface{
                self.FP.addErrMess(errPos, Lns_getVM().String_format("define class already -- %s", []LnsAny{name}))
                Util_printStackTrace()
            }
        }
    } else { 
        var parentInfo *Ast_TypeInfo
        parentInfo = self.FP.getCurrentNamespaceTypeInfo()
        var parentScope *Ast_Scope
        parentScope = self.scope
        var scope *Ast_Scope
        scope = self.FP.pushScope(true, baseInfo, interfaceList)
        var workGenTypeList *LnsList
        if genTypeList != nil{
            genTypeList_6012 := genTypeList.(*LnsList)
            workGenTypeList = genTypeList_6012
            
        } else {
            workGenTypeList = NewLnsList([]LnsAny{})
            
        }
        typeInfo = self.processInfo.FP.CreateClass(mode != TransUnit_DeclClassMode__Interface, abstractFlag, scope, baseInfo, interfaceList, workGenTypeList, parentInfo, externalFlag, accessMode, name)
        
        parentScope.FP.AddClassLazy(self.processInfo, name, errPos, typeInfo, mode == TransUnit_DeclClassMode__LazyModule)
    }
    if genTypeList != nil{
        genTypeList_6015 := genTypeList.(*LnsList)
        for _, _genType := range( genTypeList_6015.Items ) {
            genType := _genType.(Ast_AlternateTypeInfoDownCast).ToAst_AlternateTypeInfo()
            self.scope.FP.AddAlternate(self.processInfo, accessMode, genType.FP.Get_txt(), errPos, &genType.Ast_TypeInfo)
        }
    }
    var namespace *Nodes_NamespaceInfo
    
    {
        _namespace := defNamespace
        if _namespace == nil{
            namespace = NewNodes_NamespaceInfo(name, self.scope, typeInfo)
            
        } else {
            namespace = _namespace.(*Nodes_NamespaceInfo)
        }
    }
    self.typeId2ClassMap.Set(typeInfo.FP.Get_typeId(),namespace)
    return typeInfo
}

// 1097: decl @lune.@base.@TransUnit.TransUnit.popClass
func (self *TransUnit_TransUnit) popClass() {
    self.FP.popScope()
}

















// 1146: decl @lune.@base.@TransUnit.TransUnit.isTargetTokenPos
func (self *TransUnit_TransUnit) isTargetTokenPos(txt string,pos *Types_Position) bool {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.analyzePos.LineNo == pos.LineNo) &&
        Lns_GetEnv().SetStackVal( self.analyzePos.Column >= pos.Column) &&
        Lns_GetEnv().SetStackVal( self.analyzePos.Column <= pos.Column + len(txt)) ).(bool)){
        return true
    }
    return false
}

// 1156: decl @lune.@base.@TransUnit.TransUnit.isTargetToken
func (self *TransUnit_TransUnit) isTargetToken(token *Types_Token) bool {
    return self.FP.isTargetTokenPos(token.Txt, token.Pos)
}

// 1160: decl @lune.@base.@TransUnit.TransUnit.dumpSymbolType
func (self *TransUnit_TransUnit) dumpSymbolType(name string,typeInfo *Ast_TypeInfo) {
    var writer *Writer_JSON
    writer = NewWriter_JSON(Lns_io_stdout)
    writer.FP.StartParent("lunescript", false)
    writer.FP.StartParent("inquire", false)
    writer.FP.Write("name", name)
    writer.FP.Write("type", typeInfo.FP.GetTxt(self.typeNameCtrl, nil, nil))
    writer.FP.Write("typeKind", Ast_TypeInfoKind_getTxt( typeInfo.FP.Get_kind()))
    writer.FP.Write("static", Lns_getVM().String_format("%s", []LnsAny{typeInfo.FP.Get_staticFlag()}))
    writer.FP.Write("display", typeInfo.FP.Get_display_stirng_with(typeInfo.FP.Get_rawTxt(), nil))
    writer.FP.EndElement()
    writer.FP.EndElement()
    writer.FP.Fin()
    Lns_getVM().OS_exit(0)
}

// 1176: decl @lune.@base.@TransUnit.TransUnit.errorShadowing
func (self *TransUnit_TransUnit) errorShadowing(pos *Types_Position,symbolInfo LnsAny) {
    if symbolInfo != nil{
        symbolInfo_6143 := symbolInfo.(*Ast_SymbolInfo)
        var symPos LnsAny
        symPos = symbolInfo_6143.FP.Get_pos()
        if symPos != nil{
            symPos_6146 := symPos.(*Types_Position)
            var mess string
            mess = Lns_getVM().String_format("This symbol is shadowed from %d:%d -- %s", []LnsAny{pos.LineNo, pos.Column, symbolInfo_6143.FP.Get_name()})
            self.FP.addWarnErrMess(symPos_6146, Lns_op_not(self.ctrl_info.WarningShadowing), mess)
        }
        var mess string
        mess = Lns_getVM().String_format("shadowing symbol of %s -- %s", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symPos) &&
            Lns_GetEnv().SetStackVal( Lns_getVM().String_format("%s:%s", []LnsAny{Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(symPos) && 
            Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Types_Position).LineNo)), Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(symPos) && 
            Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Types_Position).Column))})) ||
            Lns_GetEnv().SetStackVal( "external") ).(string), symbolInfo_6143.FP.Get_name()})
        self.FP.addWarnErrMess(pos, Lns_op_not(self.ctrl_info.WarningShadowing), mess)
    }
}

// 1193: decl @lune.@base.@TransUnit.TransUnit.checkShadowing
func (self *TransUnit_TransUnit) checkShadowing(pos *Types_Position,name string,scope *Ast_Scope) {
    if name == "_"{
        return 
    }
    var symbolInfo *Ast_SymbolInfo
    
    {
        _symbolInfo := self.scope.FP.GetSymbolTypeInfo(name, scope, self.moduleScope, self.scopeAccess)
        if _symbolInfo == nil{
            return 
        } else {
            symbolInfo = _symbolInfo.(*Ast_SymbolInfo)
        }
    }
    self.FP.errorShadowing(pos, symbolInfo)
}

// 1207: decl @lune.@base.@TransUnit.TransUnit.addLocalVar
func (self *TransUnit_TransUnit) addLocalVar(pos *Types_Position,argFlag bool,canBeLeft bool,name string,typeInfo *Ast_TypeInfo,mutable LnsInt,allowShadow LnsAny) *Ast_SymbolInfo {
    if Lns_op_not(allowShadow){
        self.FP.checkShadowing(pos, name, self.scope)
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.analyzeMode == TransUnit_AnalyzeMode__Inquire) &&
        Lns_GetEnv().SetStackVal( self.FP.isTargetTokenPos(name, pos)) ).(bool)){
        self.FP.dumpSymbolType(name, typeInfo)
    }
    return Lns_unwrap( Lns_car(self.scope.FP.AddLocalVar(self.processInfo, argFlag, canBeLeft, name, pos, typeInfo, mutable))).(*Ast_SymbolInfo)
}


// 1289: decl @lune.@base.@TransUnit.TransUnit.createModifier
func (self *TransUnit_TransUnit) CreateModifier(typeInfo *Ast_TypeInfo,mutMode LnsInt) *Ast_TypeInfo {
    if Lns_op_not(self.validMutControl){
        return typeInfo
    }
    return self.processInfo.FP.CreateModifier(typeInfo, mutMode)
}

// 1298: decl @lune.@base.@TransUnit.TransUnit.createExtType
func (self *TransUnit_TransUnit) createExtType(pos *Types_Position,typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    switch _exp5482 := self.processInfo.FP.CreateLuaval(typeInfo, true).(type) {
    case *Ast_LuavalResult__OK:
    work := _exp5482.Val1
        return work
    case *Ast_LuavalResult__Err:
    err := _exp5482.Val1
        self.FP.addErrMess(pos, Lns_getVM().String_format("not support -- %s", []LnsAny{err}))
        return typeInfo
    }
// insert a dummy
    return nil
}

// 2190: decl @lune.@base.@TransUnit.TransUnit.checkThreading
func (self *TransUnit_TransUnit) checkThreading(pos *Types_Position) bool {
    var curClass *Ast_TypeInfo
    curClass = self.FP.getCurrentClass()
    if curClass != Ast_headTypeInfo{
        if curClass.FP.IsInheritFrom(self.processInfo, TransUnit_builtinFunc.Lnsthread_, nil){
            return true
        }
    }
    if self.helperInfo.PragmaSet.Has(LuneControl_Pragma__use_async_Obj){
        self.FP.addErrMess(pos, "This operation need perform on thread")
    }
    return false
}

// 2207: decl @lune.@base.@TransUnit.TransUnit.registBuiltInScope
func (self *TransUnit_TransUnit) registBuiltInScope() {
    if TransUnit_readyBuiltin{
        return 
    }
    TransUnit_readyBuiltin = true
    
    var builtInInfo *LnsList
    builtInInfo = TransUnit_getBuiltInInfo_3512_()
    var getTypeInfo func(typeName string) *Ast_TypeInfo
    getTypeInfo = func(typeName string) *Ast_TypeInfo {
        if _switch20155 := typeName; _switch20155 == "_T" {
            return Ast_builtinTypeBox.FP.Get_boxingType()
        } else if _switch20155 == "_T!" {
            return Ast_builtinTypeBox.FP.Get_boxingType().FP.Get_nilableTypeInfo()
        }
        var getTypeInfoFromScope func(scope *Ast_Scope,symbol string,genTypeList *LnsList) LnsAny
        getTypeInfoFromScope = func(scope *Ast_Scope,symbol string,genTypeList *LnsList) LnsAny {
            var typeInfo *Ast_TypeInfo
            
            {
                _typeInfo := scope.FP.GetTypeInfo(symbol, scope, false, self.scopeAccess)
                if _typeInfo == nil{
                    return nil
                } else {
                    typeInfo = _typeInfo.(*Ast_TypeInfo)
                }
            }
            var validGenType bool
            validGenType = false
            for _, _genType := range( genTypeList.Items ) {
                genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if genType.FP.Get_kind() != Ast_TypeInfoKind__Alternate{
                    validGenType = true
                    
                    break
                }
            }
            if validGenType{
                if _switch20425 := typeInfo.FP.Get_kind(); _switch20425 == Ast_TypeInfoKind__Map {
                    if genTypeList.Len() != 2{
                        Util_err(Lns_getVM().String_format("illegal map param -- %d", []LnsAny{genTypeList.Len()}))
                    }
                    var keyType *Ast_TypeInfo
                    keyType = genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    var valType *Ast_TypeInfo
                    valType = genTypeList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    return self.processInfo.FP.CreateMap(Ast_AccessMode__Pub, typeInfo.FP.Get_parentInfo(), keyType, valType, typeInfo.FP.Get_mutMode())
                } else if _switch20425 == Ast_TypeInfoKind__Ext {
                    if genTypeList.Len() != 1{
                        Util_err(Lns_getVM().String_format("illegal param -- %d", []LnsAny{genTypeList.Len()}))
                    }
                    switch _exp20355 := self.processInfo.FP.CreateLuaval(genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), true).(type) {
                    case *Ast_LuavalResult__OK:
                    workType := _exp20355.Val1
                        if self.ctrl_info.ValidLuaval{
                            return workType
                        }
                        return genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    case *Ast_LuavalResult__Err:
                    mess := _exp20355.Val1
                        Util_err(mess)
                    }
                } else if _switch20425 == Ast_TypeInfoKind__DDD {
                    if genTypeList.Len() != 1{
                        Util_err(Lns_getVM().String_format("illegal map param -- %d", []LnsAny{genTypeList.Len()}))
                    }
                    return &self.processInfo.FP.CreateDDD(genTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), true, false).Ast_TypeInfo
                } else {
                    Util_err(Lns_getVM().String_format("not support type -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}))
                }
            }
            return typeInfo
        }
        var mutable bool
        mutable = true
        if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(typeName,"^&", nil, nil))){
            mutable = false
            
            typeName = TransUnit_convExp20471(Lns_2DDD(Lns_getVM().String_gsub(typeName,"^&", "")))
            
        }
        var genTypeList *LnsList
        genTypeList = NewLnsList([]LnsAny{})
        var endIndex LnsAny
        _,endIndex = TransUnit_convExp20496(Lns_2DDD(Lns_getVM().String_find(typeName,"[%w%.]+<", nil, nil)))
        var suffix string
        suffix = ""
        if endIndex != nil{
            endIndex_7109 := endIndex.(LnsInt)
            var genTypeName string
            genTypeName = Lns_getVM().String_sub(typeName,endIndex_7109 + 1, nil)
            for  {
                {
                    _tailIndex := TransUnit_convExp20589(Lns_2DDD(Lns_getVM().String_find(genTypeName,"[,>]", nil, nil)))
                    if _tailIndex != nil {
                        tailIndex := _tailIndex.(LnsInt)
                        var genType *Ast_TypeInfo
                        genType = getTypeInfo(Lns_getVM().String_sub(genTypeName,1, tailIndex - 1))
                        genTypeList.Insert(Ast_TypeInfo2Stem(genType))
                        genTypeName = Lns_getVM().String_sub(genTypeName,tailIndex + 1, nil)
                        
                    } else {
                        suffix = Lns_getVM().String_sub(genTypeName,1, nil)
                        
                        break
                    }
                }
            }
            typeName = Lns_getVM().String_sub(typeName,1, endIndex_7109 - 1) + suffix
            
        }
        var typeInfo *Ast_TypeInfo
        typeInfo = Ast_headTypeInfo
        if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(typeName,"!$", nil, nil))){
            var orgTypeName string
            orgTypeName = TransUnit_convExp20644(Lns_2DDD(Lns_getVM().String_gsub(typeName,"!$", "")))
            {
                __exp := getTypeInfoFromScope(self.scope, orgTypeName, genTypeList)
                if __exp != nil {
                    _exp := __exp.(*Ast_TypeInfo)
                    typeInfo = _exp
                    
                } else {
                    Util_err(Lns_getVM().String_format("not found builtin -- %s", []LnsAny{orgTypeName}))
                }
            }
            typeInfo = typeInfo.FP.Get_nilableTypeInfo()
            
        } else { 
            {
                __exp := getTypeInfoFromScope(self.scope, typeName, genTypeList)
                if __exp != nil {
                    _exp := __exp.(*Ast_TypeInfo)
                    typeInfo = _exp
                    
                } else {
                    Util_err(Lns_getVM().String_format("not found builtin -- %s", []LnsAny{typeName}))
                }
            }
        }
        if mutable{
            return typeInfo
        }
        typeInfo = self.FP.CreateModifier(typeInfo, Ast_MutMode__IMut)
        
        return typeInfo
    }
    var processField func(name string,fieldName string,info *LnsMap,parentInfo *Ast_TypeInfo)
    processField = func(name string,fieldName string,info *LnsMap,parentInfo *Ast_TypeInfo) {
        if self.targetLuaVer.FP.IsSupport(Lns_getVM().String_format("%s.%s", []LnsAny{name, fieldName})){
            if Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
            info.Items["type"]) && 
            Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsList).GetAt(1))) == "var"{
                var symbol *Ast_SymbolInfo
                symbol = Lns_unwrap( Lns_car(self.scope.FP.Add(self.processInfo, Ast_SymbolKind__Var, false, true, fieldName, nil, getTypeInfo(Lns_unwrap( Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
                info.Items["typeInfo"]) && 
                Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsList).GetAt(1)))).(string)), Ast_AccessMode__Pub, true, Ast_MutMode__Mut, true, false))).(*Ast_SymbolInfo)
                TransUnit_setupBuiltinTypeInfo_3435_(name, fieldName, symbol)
            } else { 
                var argTypeList *LnsList
                argTypeList = NewLnsList([]LnsAny{})
                for _, _argType := range( Lns_unwrap( info.Items["arg"]).(*LnsList).Items ) {
                    argType := _argType.(string)
                    argTypeList.Insert(Ast_TypeInfo2Stem(getTypeInfo(argType)))
                }
                var retTypeList *LnsList
                retTypeList = NewLnsList([]LnsAny{})
                for _, _retType := range( Lns_unwrap( info.Items["ret"]).(*LnsList).Items ) {
                    retType := _retType.(string)
                    var retTypeInfo *Ast_TypeInfo
                    retTypeInfo = getTypeInfo(retType)
                    retTypeList.Insert(Ast_TypeInfo2Stem(retTypeInfo))
                }
                var funcType LnsAny
                funcType = Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
                info.Items["type"]) && 
                Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsList).GetAt(1)))
                var mutable bool
                var staticFlag bool
                var kind LnsInt
                var symbolKind LnsInt
                var abstractFlag bool
                var accessMode LnsInt
                if _switch21198 := funcType; _switch21198 == "method" || _switch21198 == "mut" {
                    staticFlag = false
                    
                    kind = Ast_TypeInfoKind__Method
                    
                    symbolKind = Ast_SymbolKind__Mtd
                    
                    abstractFlag = false
                    
                    accessMode = Ast_AccessMode__Pub
                    
                    mutable = funcType == "mut"
                    
                } else if _switch21198 == "abstract" {
                    staticFlag = false
                    
                    kind = Ast_TypeInfoKind__Method
                    
                    symbolKind = Ast_SymbolKind__Mtd
                    
                    abstractFlag = true
                    
                    accessMode = Ast_AccessMode__Pro
                    
                    mutable = true
                    
                } else if _switch21198 == "macro" {
                    staticFlag = true
                    
                    kind = Ast_TypeInfoKind__Macro
                    
                    symbolKind = Ast_SymbolKind__Mac
                    
                    abstractFlag = false
                    
                    accessMode = Ast_AccessMode__Pub
                    
                    mutable = false
                    
                } else if _switch21198 == "formfunc" {
                    staticFlag = true
                    
                    kind = Ast_TypeInfoKind__FormFunc
                    
                    symbolKind = Ast_SymbolKind__Typ
                    
                    abstractFlag = false
                    
                    accessMode = Ast_AccessMode__Pub
                    
                    mutable = false
                    
                } else {
                    staticFlag = true
                    
                    kind = Ast_TypeInfoKind__Func
                    
                    symbolKind = Ast_SymbolKind__Fun
                    
                    abstractFlag = false
                    
                    accessMode = Ast_AccessMode__Pub
                    
                    mutable = false
                    
                }
                self.FP.pushScope(false, nil, nil)
                var typeInfo *Ast_TypeInfo
                typeInfo = self.processInfo.FP.CreateFunc(abstractFlag, true, self.scope, kind, parentInfo, false, true, staticFlag, accessMode, fieldName, nil, argTypeList, retTypeList, mutable)
                self.FP.popScope()
                TransUnit_builtinFunc.FP.Get_allFuncTypeSet().Add(Ast_TypeInfo2Stem(typeInfo))
                Ast_addBuiltin(typeInfo)
                if typeInfo.FP.Get_nilableTypeInfo() != Ast_headTypeInfo{
                    Ast_addBuiltin(typeInfo.FP.Get_nilableTypeInfo())
                }
                var symInfo *Ast_SymbolInfo
                symInfo = Lns_unwrap( Lns_car(self.scope.FP.Add(self.processInfo, symbolKind, false, kind == Ast_TypeInfoKind__Func, fieldName, nil, typeInfo, accessMode, staticFlag, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( mutable) &&
                    Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
                    Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, false))).(*Ast_SymbolInfo)
                TransUnit_setupBuiltinTypeInfo_3435_(name, fieldName, symInfo)
            }
        }
    }
    self.scope = Ast_rootScope
    
    var builtinModuleName2Scope *LnsMap
    builtinModuleName2Scope = NewLnsMap( map[LnsAny]LnsAny{})
    var mapType *Ast_TypeInfo
    mapType = self.processInfo.FP.CreateMap(Ast_AccessMode__Pub, Ast_headTypeInfo, Ast_builtinTypeString, Ast_builtinTypeStem, Ast_MutMode__Mut)
    self.scope.FP.AddVar(self.processInfo, Ast_AccessMode__Global, "_ENV", nil, mapType, Ast_MutMode__IMutRe, true)
    self.scope.FP.AddVar(self.processInfo, Ast_AccessMode__Global, "_G", nil, mapType, Ast_MutMode__IMutRe, true)
    self.scope.FP.AddVar(self.processInfo, Ast_AccessMode__Global, "__line__", nil, Ast_builtinTypeInt, Ast_MutMode__IMut, true)
    for _, _builtinClassInfo := range( builtInInfo.Items ) {
        builtinClassInfo := _builtinClassInfo.(*LnsMap)
        for _className, _name2FieldInfo := range( builtinClassInfo.Items ) {
            className := _className.(string)
            name2FieldInfo := _name2FieldInfo.(*LnsMap)
            var name string
            name = className
            var genTypeList *LnsList
            genTypeList = NewLnsList([]LnsAny{})
            if _switch21698 := className; _switch21698 == "List<T>" {
                name = "List"
                
                TransUnit_registBuiltInScope__processCopyAlterList_4630_(genTypeList, Ast_builtinTypeList.FP.Get_itemTypeInfoList())
            } else if _switch21698 == "Array<T>" {
                name = "Array"
                
                TransUnit_registBuiltInScope__processCopyAlterList_4630_(genTypeList, Ast_builtinTypeArray.FP.Get_itemTypeInfoList())
            } else if _switch21698 == "Set<T>" {
                name = "Set"
                
                TransUnit_registBuiltInScope__processCopyAlterList_4630_(genTypeList, Ast_builtinTypeSet.FP.Get_itemTypeInfoList())
            } else if _switch21698 == "Nilable<_T>" {
                name = "Nilable"
                
                TransUnit_registBuiltInScope__processCopyAlterList_4630_(genTypeList, Ast_builtinTypeBox.FP.Get_itemTypeInfoList())
            } else {
                if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(className,"<", nil, nil))){
                    name = ""
                    
                    {
                        _form21694, _param21694, _prev21694 := Lns_getVM().String_gmatch(className,"[^<>,%s]+")
                        for {
                            _work21694 := _form21694.(*Lns_luaValue).Call( Lns_2DDD( _param21694, _prev21694 ) )
                            _prev21694 = Lns_getFromMulti(_work21694,0)
                            if Lns_IsNil( _prev21694 ) { break }
                            token := _prev21694.(string)
                            if len(name) == 0{
                                name = token
                                
                            } else { 
                                genTypeList.Insert(Ast_AlternateTypeInfo2Stem(self.processInfo.FP.CreateAlternate(true, genTypeList.Len() + 1, token, Ast_AccessMode__Pri, self.moduleType, nil, nil)))
                            }
                        }
                    }
                }
            }
            var parentInfo *Ast_TypeInfo
            parentInfo = Ast_headTypeInfo
            if name != ""{
                var classKind LnsInt
                classKind = TransUnit_DeclClassMode__Class
                {
                    _types := Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
                    name2FieldInfo.Items["__attrib"]) && 
                    Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsMap).Items["type"]))
                    if _types != nil {
                        types := _types.(*LnsList)
                        if types.Len() > 0{
                            if _switch21757 := types.GetAt(1).(string); _switch21757 == "interface" {
                                classKind = TransUnit_DeclClassMode__Interface
                                
                            } else if _switch21757 == "module" {
                                classKind = TransUnit_DeclClassMode__Module
                                
                            }
                        }
                    }
                }
                var interfaceList *LnsList
                interfaceList = NewLnsList([]LnsAny{})
                {
                    __exp := Lns_GetEnv().NilAccFin( Lns_GetEnv().NilAccPush(
                    name2FieldInfo.Items["__attrib"]) && 
                    Lns_GetEnv().NilAccPush( Lns_GetEnv().NilAccPop().(*LnsMap).Items["inplements"]))
                    if __exp != nil {
                        _exp := __exp.(*LnsList)
                        for _, _ifname := range( _exp.Items ) {
                            ifname := _ifname.(string)
                            var ifType *Ast_TypeInfo
                            ifType = getTypeInfo(ifname)
                            interfaceList.Insert(Ast_TypeInfo2Stem(ifType))
                        }
                    }
                }
                if _switch21933 := classKind; _switch21933 == TransUnit_DeclClassMode__Class || _switch21933 == TransUnit_DeclClassMode__Interface {
                    var declMode LnsInt
                    if classKind == TransUnit_DeclClassMode__Class{
                        declMode = TransUnit_DeclClassMode__Class
                        
                    } else { 
                        declMode = TransUnit_DeclClassMode__Interface
                        
                    }
                    parentInfo = self.FP.pushClass(self.parser.FP.Get_currentToken().Pos, declMode, false, nil, interfaceList, genTypeList, true, name, true, Ast_AccessMode__Pub, nil)
                    
                    TransUnit_builtinFunc.FP.RegisterClass(parentInfo)
                } else if _switch21933 == TransUnit_DeclClassMode__Module {
                    parentInfo = self.FP.pushModule(true, name, true)
                    
                    self.scope.FP.Get_parent().FP.Add(self.processInfo, Ast_SymbolKind__Typ, false, false, name, nil, parentInfo, Ast_AccessMode__Local, true, Ast_MutMode__Mut, true, false)
                }
                Ast_addBuiltin(parentInfo)
                Ast_addBuiltin(parentInfo.FP.Get_nilableTypeInfo())
            }
            if Lns_op_not(builtinModuleName2Scope.Items[name]){
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( name != "") &&
                    Lns_GetEnv().SetStackVal( getTypeInfo(name)) )){
                    builtinModuleName2Scope.Set(name,self.scope)
                }
                {
                    __collection22012 := name2FieldInfo
                    __sorted22012 := __collection22012.CreateKeyListStr()
                    __sorted22012.Sort( LnsItemKindStr, nil )
                    for _, _fieldName := range( __sorted22012.Items ) {
                        info := __collection22012.Items[ _fieldName ].(*LnsMap)
                        fieldName := _fieldName.(string)
                        if _switch22010 := fieldName; _switch22010 == "__attrib" {
                        } else {
                            processField(name, fieldName, info, parentInfo)
                        }
                    }
                }
            }
            if name != ""{
                self.FP.popClass()
            }
        }
    }
    var threadSafeSet *LnsSet
    threadSafeSet = NewLnsSet([]LnsAny{Ast_TypeInfo2Stem(TransUnit_builtinFunc.Lns_error), Ast_TypeInfo2Stem(TransUnit_builtinFunc.Lns_print), Ast_TypeInfo2Stem(TransUnit_builtinFunc.Lns_type), Ast_TypeInfo2Stem(TransUnit_builtinFunc.Lns_tonumber), Ast_TypeInfo2Stem(TransUnit_builtinFunc.Io_open), Ast_TypeInfo2Stem(TransUnit_builtinFunc.Set_has), Ast_TypeInfo2Stem(TransUnit_builtinFunc.Set_add)})
    for _typeInfo := range( TransUnit_builtinFunc.FP.Get_allFuncTypeSet().Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(threadSafeSet.Has(Ast_TypeInfo2Stem(typeInfo))){
            TransUnit_builtinFunc.FP.Get_needThreadingTypes().Add(Ast_TypeInfo2Stem(typeInfo))
        }
    }
    self.scope = self.topScope
    
}

// 2598: decl @lune.@base.@TransUnit.TransUnit.errorAt
func (self *TransUnit_TransUnit) ErrorAt(pos *Types_Position,mess string) {
    self.FP.addErrMess(pos, mess)
    for _, _errmess := range( self.errMessList.Items ) {
        errmess := _errmess.(string)
        Util_errorLog(errmess)
    }
    for _, _warnmess := range( self.warnMessList.Items ) {
        warnmess := _warnmess.(string)
        Util_errorLog(warnmess)
    }
    if self.macroCtrl.FP.Get_analyzeInfo().FP.Get_mode() != Nodes_MacroMode__None{
        Lns_print([]LnsAny{"------ near code -----", Nodes_MacroMode_getTxt( self.macroCtrl.FP.Get_analyzeInfo().FP.Get_mode())})
        Lns_print([]LnsAny{self.parser.FP.GetNearCode()})
        Lns_print([]LnsAny{"------"})
    }
    Util_err("has error")
}

// 2616: decl @lune.@base.@TransUnit.TransUnit.error
func (self *TransUnit_TransUnit) Error(mess string) {
    self.FP.ErrorAt(self.parser.FP.GetLastPos(), mess)
}

// 2620: decl @lune.@base.@TransUnit.TransUnit.createNoneNode
func (self *TransUnit_TransUnit) createNoneNode(pos *Types_Position) *Nodes_Node {
    return &Nodes_NoneNode_create(self.nodeManager, pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})).Nodes_Node
}

// 2627: decl @lune.@base.@TransUnit.TransUnit.pushbackToken
func (self *TransUnit_TransUnit) PushbackToken(token *Types_Token) {
    self.parser.FP.PushbackToken(token)
}

// 2632: decl @lune.@base.@TransUnit.TransUnit.newPushback
func (self *TransUnit_TransUnit) NewPushback(tokenList *LnsList) {
    self.parser.FP.NewPushback(tokenList)
}

// 2636: decl @lune.@base.@TransUnit.TransUnit.getStreamName
func (self *TransUnit_TransUnit) GetStreamName() string {
    return self.parser.FP.GetStreamName()
}

// 2640: decl @lune.@base.@TransUnit.TransUnit.createPosition
func (self *TransUnit_TransUnit) CreatePosition(lineNo LnsInt,column LnsInt) *Types_Position {
    return self.parser.FP.CreatePosition(lineNo, column)
}

// 2644: decl @lune.@base.@TransUnit.TransUnit.getTokenNoErr
func (self *TransUnit_TransUnit) GetTokenNoErr() *Types_Token {
    var token *Types_Token
    var commentList *LnsList
    commentList = NewLnsList([]LnsAny{})
    var workToken *Types_Token
    workToken = self.parser.FP.GetTokenNoErr()
    for workToken.Kind == Types_TokenKind__Cmnt {
        commentList.Insert(Types_Token2Stem(workToken))
        workToken = self.parser.FP.GetTokenNoErr()
        
    }
    if workToken.Kind != Types_TokenKind__Eof{
        token = workToken
        
        if self.macroCtrl.FP.Get_analyzeInfo().FP.Get_mode() != Nodes_MacroMode__None{
            token = self.macroCtrl.FP.ExpandMacroVal(self.typeNameCtrl, self.scope, self.FP, token)
            
        }
        if Lns_op_not(self.ctrl_info.CompatComment){
            self.commentCtrl.FP.AddDirect(commentList)
        }
    } else { 
        token = Parser_getEofToken()
        
        self.commentCtrl.FP.AddDirect(commentList)
    }
    if token.FP.Get_commentList().Len() > 0{
        self.commentCtrl.FP.Add(token)
    }
    return token
}

// 2678: decl @lune.@base.@TransUnit.TransUnit.getToken
func (self *TransUnit_TransUnit) getToken(allowEof LnsAny) *Types_Token {
    var token *Types_Token
    token = self.FP.GetTokenNoErr()
    if token == Parser_getEofToken(){
        if Lns_isCondTrue( allowEof){
            return Parser_getEofToken()
        }
        self.FP.Error("EOF")
    }
    return token
}

// 2690: decl @lune.@base.@TransUnit.TransUnit.pushback
func (self *TransUnit_TransUnit) Pushback() {
    self.parser.FP.Pushback()
}

// 2694: decl @lune.@base.@TransUnit.TransUnit.pushbackStr
func (self *TransUnit_TransUnit) PushbackStr(name string,statement string) {
    self.parser.FP.PushbackStr(name, statement)
}

// 2720: decl @lune.@base.@TransUnit.TransUnit.checkSymbol
func (self *TransUnit_TransUnit) checkSymbol(token *Types_Token,mode LnsInt) *Types_Token {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( token.Kind != Types_TokenKind__Symb) &&
        Lns_GetEnv().SetStackVal( token.Kind != Types_TokenKind__Kywd) &&
        Lns_GetEnv().SetStackVal( token.Kind != Types_TokenKind__Type) ).(bool)){
        self.FP.addErrMess(token.Pos, Lns_getVM().String_format("illegal symbol -- '%s'", []LnsAny{token.Txt}))
    }
    var frontChar LnsInt
    frontChar = LnsInt(token.Txt[1-1])
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_SymbolMode__Must_) &&
        Lns_GetEnv().SetStackVal( frontChar != 95) ).(bool)){
        self.FP.addErrMess(token.Pos, Lns_getVM().String_format("macro name must begin with '_' -- '%s'", []LnsAny{token.Txt}))
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode != TransUnit_SymbolMode__Must_) &&
        Lns_GetEnv().SetStackVal( frontChar == 95) ).(bool)){
        if Lns_op_not(self.ignoreToCheckSymbol_){
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( mode == TransUnit_SymbolMode__MustNot_Or_) &&
                Lns_GetEnv().SetStackVal( token.Txt == "_") ).(bool)){
            } else if Lns_op_not(TransUnit_specialSymbolSet.Has(token.Txt)){
                self.FP.addErrMess(token.Pos, Lns_getVM().String_format("symbol must not begin with '_' -- '%s'", []LnsAny{token.Txt}))
            }
        }
    } else if TransUnit_builtinKeywordSet.Has(token.Txt){
        self.FP.addErrMess(token.Pos, Lns_getVM().String_format("this symbol is special keyword -- %s", []LnsAny{token.Txt}))
    } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Parser_isLuaKeyword(token.Txt)) ||
        Lns_GetEnv().SetStackVal( Parser_isOp2(token.Txt)) ||
        Lns_GetEnv().SetStackVal( Parser_isOp1(token.Txt)) ).(bool){
        self.FP.addErrMess(token.Pos, Lns_getVM().String_format("this symbol is lua keyword -- %s", []LnsAny{token.Txt}))
    }
    return token
}

// 2756: decl @lune.@base.@TransUnit.TransUnit.getSymbolToken
func (self *TransUnit_TransUnit) getSymbolToken(mode LnsInt) *Types_Token {
    return self.FP.checkSymbol(self.FP.getToken(nil), mode)
}

// 2761: decl @lune.@base.@TransUnit.TransUnit.checkToken
func (self *TransUnit_TransUnit) checkToken(token *Types_Token,txt string) *Types_Token {
    if token.Txt != txt{
        self.FP.Error(Lns_getVM().String_format("not found -- %s. expects %s", []LnsAny{txt, token.Txt}))
    }
    return token
}

// 2768: decl @lune.@base.@TransUnit.TransUnit.checkNextToken
func (self *TransUnit_TransUnit) checkNextToken(txt string) *Types_Token {
    return self.FP.checkToken(self.FP.getToken(nil), txt)
}

// 2778: decl @lune.@base.@TransUnit.TransUnit.getContinueToken
func (self *TransUnit_TransUnit) getContinueToken()(*Types_Token, bool) {
    var token *Types_Token
    token = self.FP.getToken(nil)
    return token, token.Consecutive
}

// 2784: decl @lune.@base.@TransUnit.TransUnit.analyzeStatementList
func (self *TransUnit_TransUnit) analyzeStatementList(stmtList *LnsList,termTxt LnsAny)(LnsAny, LnsInt) {
    var breakKind LnsInt
    breakKind = Nodes_BreakKind__None
    if stmtList.Len() > 0{
        breakKind = stmtList.GetAt(stmtList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.GetBreakKind(Nodes_CheckBreakMode__Normal)
        
    }
    var parser2lastLineMap *LnsMap
    parser2lastLineMap = NewLnsMap( map[LnsAny]LnsAny{})
    var getLastLineNo func() LnsInt
    getLastLineNo = func() LnsInt {
        {
            _lastLineNo := parser2lastLineMap.Items[self.parser.FP]
            if _lastLineNo != nil {
                lastLineNo := _lastLineNo.(LnsInt)
                return lastLineNo
            }
        }
        return self.parser.FP.GetLastPos().LineNo
    }
    var setLastLineNo func(lineNo LnsInt)
    setLastLineNo = func(lineNo LnsInt) {
        parser2lastLineMap.Set(self.parser.FP,lineNo)
    }
    var lastStatement LnsAny
    lastStatement = nil
    var lastLineNo LnsInt
    lastLineNo = getLastLineNo()
    var setTailComment func(statement LnsAny) LnsInt
    setTailComment = func(statement LnsAny) LnsInt {
        var blank LnsInt
        var commentList *LnsList
        commentList = self.commentCtrl.FP.Get_commentList()
        if commentList.Len() > 0{
            if lastStatement != nil{
                lastStatement_7363 := lastStatement.(*Nodes_Node)
                var tailComment LnsAny
                tailComment = nil
                for _, _comment := range( commentList.Items ) {
                    comment := _comment.(Types_TokenDownCast).ToTypes_Token()
                    if comment.Pos.LineNo == lastStatement_7363.FP.Get_pos().LineNo{
                        if Lns_op_not(tailComment){
                            lastStatement_7363.FP.Set_tailComment(comment)
                            tailComment = comment
                            
                        } else { 
                        }
                    }
                }
                if Lns_isCondTrue( tailComment){
                    commentList.Remove(1)
                }
            }
        }
        if commentList.Len() > 0{
            blank = commentList.GetAt(1).(Types_TokenDownCast).ToTypes_Token().Pos.LineNo - commentList.GetAt(1).(Types_TokenDownCast).ToTypes_Token().FP.GetLineCount() - lastLineNo
            
        } else { 
            if statement != nil{
                statement_7374 := statement.(*Nodes_Node)
                blank = statement_7374.FP.Get_pos().LineNo - lastLineNo
                
            } else {
                blank = self.parser.FP.GetLastPos().LineNo - lastLineNo
                
            }
        }
        return blank
    }
    for  {
        lastLineNo = getLastLineNo()
        
        var statement LnsAny
        statement = self.FP.analyzeStatement(termTxt)
        if statement != nil{
            statement_7379 := statement.(*Nodes_Node)
            if breakKind != Nodes_BreakKind__None{
                if statement_7379.FP.Get_kind() != Nodes_NodeKind_get_BlankLine(){
                    self.FP.addErrMess(statement_7379.FP.Get_pos(), Lns_getVM().String_format("This statement is not reached -- %s", []LnsAny{Nodes_BreakKind_getTxt( breakKind)}))
                }
            }
            var blank LnsInt
            blank = setTailComment(statement_7379)
            if blank > 1{
                stmtList.Insert(Nodes_BlankLineNode2Stem(Nodes_BlankLineNode_create(self.nodeManager, self.FP.CreatePosition(lastLineNo + 1, 0), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), blank - 1)))
            }
            setLastLineNo(self.parser.FP.GetLastPos().LineNo)
            stmtList.Insert(Nodes_Node2Stem(statement_7379))
            lastStatement = statement_7379
            
            if statement_7379.FP.Get_kind() != Nodes_NodeKind_get_BlankLine(){
                breakKind = statement_7379.FP.GetBreakKind(Nodes_CheckBreakMode__Normal)
                
            }
            statement_7379.FP.AddComment(self.commentCtrl.FP.Get_commentList())
            self.commentCtrl.FP.Clear()
        } else {
            setTailComment(nil)
            break
        }
    }
    return lastStatement, lastLineNo
}

// 2894: decl @lune.@base.@TransUnit.TransUnit.analyzeStatementListSubfile
func (self *TransUnit_TransUnit) analyzeStatementListSubfile(stmtList *LnsList) LnsAny {
    var statement LnsAny
    statement = self.FP.analyzeStatement(nil)
    {
        __exp := statement
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            if _exp.FP.Get_kind() != Nodes_NodeKind_get_Subfile(){
                self.FP.Error("subfile must have 'subfile' declaration at top.")
            }
        } else {
            self.FP.Error("subfile must have 'subfile' declaration at top.")
        }
    }
    return Lns_car(self.FP.analyzeStatementList(stmtList, nil))
}

// 2909: decl @lune.@base.@TransUnit.TransUnit.supportLang
func (self *TransUnit_TransUnit) supportLang(lang string) bool {
    for _pragma := range( self.helperInfo.PragmaSet.Items ) {
        pragma := _pragma
        switch _exp23552 := pragma.(type) {
        case *LuneControl_Pragma__limit_conv_code:
        codeSet := _exp23552.Val1
            return codeSet.Has(lang)
        }
    }
    return true
}

// 2920: decl @lune.@base.@TransUnit.TransUnit.analyzeLuneControl
func (self *TransUnit_TransUnit) analyzeLuneControl(firstToken *Types_Token) LnsAny {
    var node LnsAny
    node = nil
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var pragma LnsAny
    if _switch23916 := (nextToken.Txt); _switch23916 == "disable_mut_control" {
        self.validMutControl = false
        
        pragma = LuneControl_Pragma__disable_mut_control_Obj
        
    } else if _switch23916 == "ignore_symbol_" {
        self.ignoreToCheckSymbol_ = true
        
        pragma = LuneControl_Pragma__ignore_symbol__Obj
        
    } else if _switch23916 == "load__lune_module" {
        pragma = LuneControl_Pragma__load__lune_module_Obj
        
    } else if _switch23916 == "limit_conv_code" {
        var codeSet *LnsSet
        codeSet = NewLnsSet([]LnsAny{})
        for  {
            var token *Types_Token
            token = self.FP.getToken(nil)
            if token.Txt == ";"{
                self.FP.Pushback()
                break
            }
            {
                _code := LuneControl_Code__from(token.Txt)
                if _code != nil {
                    code := _code.(string)
                    codeSet.Add(code)
                } else {
                    self.FP.addErrMess(token.Pos, Lns_getVM().String_format("illegal code -- '%s'", []LnsAny{token.Txt}))
                }
            }
        }
        pragma = &LuneControl_Pragma__limit_conv_code{codeSet}
        
    } else if _switch23916 == "use_async" {
        pragma = LuneControl_Pragma__use_async_Obj
        
    } else if _switch23916 == "run_async_pipe" {
        if Lns_op_not(self.helperInfo.PragmaSet.Has(LuneControl_Pragma__use_async_Obj)){
            self.FP.addErrMess(nextToken.Pos, "must set '_lune_control use_async'")
        }
        var nowMethod *Ast_TypeInfo
        nowMethod = self.FP.getCurrentNamespaceTypeInfo()
        var nowClass *Ast_TypeInfo
        nowClass = nowMethod.FP.Get_parentInfo()
        var valid bool
        valid = false
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( nowMethod.FP.Get_kind() == Ast_TypeInfoKind__Method) &&
            Lns_GetEnv().SetStackVal( Ast_isClass(nowClass)) ).(bool)){
            {
                _loopMethod := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(nowClass.FP.Get_scope()) && 
                Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.GetTypeInfoChild("loop")})/* 2968:34 */)
                if _loopMethod != nil {
                    loopMethod := _loopMethod.(*Ast_TypeInfo)
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( loopMethod.FP.Get_kind() == Ast_TypeInfoKind__Method) &&
                        Lns_GetEnv().SetStackVal( loopMethod.FP.Get_argTypeInfoList().Len() == 0) ).(bool)){
                        valid = true
                        
                    }
                }
            }
        }
        if valid{
            pragma = LuneControl_Pragma__run_async_pipe_Obj
            
        } else { 
            self.FP.addErrMess(nextToken.Pos, "this option only use in method of the class have loop method.")
            return nil
        }
    } else {
        self.FP.addErrMess(nextToken.Pos, Lns_getVM().String_format("unknown option -- %s", []LnsAny{nextToken.Txt}))
        self.FP.checkNextToken(";")
        return nil
    }
    node = Nodes_LuneControlNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), pragma)
    
    self.helperInfo.PragmaSet.Add(pragma)
    self.FP.checkNextToken(";")
    return node
}

// 3014: decl @lune.@base.@TransUnit.TransUnit.analyzeBlock
func (self *TransUnit_TransUnit) analyzeBlock(blockKind LnsInt,tentativeMode LnsInt,scope LnsAny,refAccessSymPosList LnsAny) *Nodes_BlockNode {
    var token *Types_Token
    token = self.FP.checkNextToken("{")
    var backScope *Ast_Scope
    backScope = self.scope
    if scope != nil{
        scope_7458 := scope.(*Ast_Scope)
        self.scope = scope_7458
        
    } else {
        self.FP.pushScope(false, nil, nil)
    }
    var blockScope *Ast_Scope
    blockScope = self.scope
    blockScope.FP.AddIgnoredVar(self.processInfo)
    if _switch24106 := tentativeMode; _switch24106 == TransUnit_TentativeMode__Loop {
        self.FP.prepareTentativeSymbol(self.scope, true, refAccessSymPosList)
    } else if _switch24106 == TransUnit_TentativeMode__Simple || _switch24106 == TransUnit_TentativeMode__Start || _switch24106 == TransUnit_TentativeMode__Ignore {
        self.FP.prepareTentativeSymbol(self.scope, false, nil)
    } else if _switch24106 == TransUnit_TentativeMode__Merge || _switch24106 == TransUnit_TentativeMode__Finish {
        self.FP.mergeTentativeSymbol(self.scope)
    }
    var loopFlag bool
    loopFlag = false
    if _switch24146 := blockKind; _switch24146 == Nodes_BlockKind__For || _switch24146 == Nodes_BlockKind__Apply || _switch24146 == Nodes_BlockKind__While || _switch24146 == Nodes_BlockKind__Repeat || _switch24146 == Nodes_BlockKind__Foreach {
        loopFlag = true
        
        self.loopScopeQueue.Insert(Ast_Scope2Stem(self.scope))
    }
    var stmtList *LnsList
    stmtList = NewLnsList([]LnsAny{})
    self.FP.analyzeStatementList(stmtList, "}")
    self.FP.checkNextToken("}")
    if loopFlag{
        self.loopScopeQueue.Remove(nil)
    }
    if scope != nil{
        self.scope = backScope
        
    } else {
        self.FP.popScope()
    }
    var node *Nodes_BlockNode
    node = Nodes_BlockNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), blockKind, blockScope, stmtList)
    if node.FP.GetBreakKind(Nodes_CheckBreakMode__Normal) != Nodes_BreakKind__None{
        self.tentativeSymbol.FP.Skip()
    }
    if blockKind != Nodes_BlockKind__Repeat{
        if _switch24297 := tentativeMode; _switch24297 == TransUnit_TentativeMode__Simple || _switch24297 == TransUnit_TentativeMode__Finish {
            self.FP.finishTentativeSymbol(true)
        } else if _switch24297 == TransUnit_TentativeMode__Ignore || _switch24297 == TransUnit_TentativeMode__Loop {
            self.FP.finishTentativeSymbol(false)
        }
    }
    return node
}

// 3111: decl @lune.@base.@TransUnit.TransUnit.processImport
func (self *TransUnit_TransUnit) processImport(modulePath string,depth LnsInt)(*Lns_luaValue, *LnsMap, *Nodes_ModuleInfo) {
    __func__ := "@lune.@base.@TransUnit.TransUnit.processImport"
    Log_log(Log_Level__Info, __func__, 3114, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s -> %s start", []LnsAny{self.moduleType.FP.GetTxt(self.typeNameCtrl, nil, nil), modulePath})
    }))
    
    if Lns_op_not(self.importModuleInfo.FP.Add(modulePath)){
        self.FP.Error(Lns_getVM().String_format("recursive import: %s -> %s", []LnsAny{self.importModuleInfo.FP.GetFull(), modulePath}))
    }
    {
        _moduleInfo := self.importModuleName2ModuleInfo.Items[modulePath]
        if _moduleInfo != nil {
            moduleInfo := _moduleInfo.(*Nodes_ModuleInfo)
            {
                _metaInfoStem := FrontInterface_loadMeta(self.importModuleInfo, modulePath)
                if _metaInfoStem != nil {
                    metaInfoStem := _metaInfoStem
                    Log_log(Log_Level__Info, __func__, 3126, Log_CreateMessage(func() string {
                        return Lns_getVM().String_format("%s already", []LnsAny{modulePath})
                    }))
                    
                    var metaInfo *Lns_luaValue
                    metaInfo = metaInfoStem.(*Lns_luaValue)
                    var typeId2TypeInfo *LnsMap
                    typeId2TypeInfo = moduleInfo.FP.Get_importId2localTypeInfoMap()
                    self.importModuleInfo.FP.Remove()
                    for _key, _val := range( moduleInfo.FP.Get_importedAliasMap().Items ) {
                        key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        val := _val.(Ast_AliasTypeInfoDownCast).ToAst_AliasTypeInfo()
                        self.importedAliasMap.Set(key,val)
                    }
                    return metaInfo, typeId2TypeInfo, moduleInfo
                }
            }
            self.FP.Error("failed to load meta -- " + modulePath)
        }
    }
    var nameList *LnsList
    nameList = NewLnsList([]LnsAny{})
    {
        _form24626, _param24626, _prev24626 := Lns_getVM().String_gmatch(modulePath, "[^%./:]+")
        for {
            _work24626 := _form24626.(*Lns_luaValue).Call( Lns_2DDD( _param24626, _prev24626 ) )
            _prev24626 = Lns_getFromMulti(_work24626,0)
            if Lns_IsNil( _prev24626 ) { break }
            txt := _prev24626.(string)
            nameList.Insert(txt)
        }
    }
    var metaInfoStem LnsAny
    
    {
        _metaInfoStem := FrontInterface_loadMeta(self.importModuleInfo, modulePath)
        if _metaInfoStem == nil{
            self.FP.Error("failed to load meta -- " + modulePath)
        } else {
            metaInfoStem = _metaInfoStem
        }
    }
    var metaInfo *Lns_luaValue
    metaInfo = metaInfoStem.(*Lns_luaValue)
    Log_log(Log_Level__Debug, __func__, 3150, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s processing", []LnsAny{modulePath})
    }))
    
    var dependLibId2DependInfo *LnsMap
    dependLibId2DependInfo = NewLnsMap( map[LnsAny]LnsAny{})
    {
        _exp24776 := metaInfo.GetAt( "__dependModuleMap" ).(*Lns_luaValue)
        _sorted24776 := Lns_getVM().SortMapKeyList( _exp24776 )
        _index24776, _key24776 := _sorted24776.Get1stFromMap()
        for _index24776 != nil {
            dependName := _key24776.(string)
            dependInfo := _exp24776.GetAt( _key24776 ).(*Lns_luaValue)
            if Lns_isCondTrue( dependInfo.GetAt("use")){
                var metaTypeId2TypeInfoMap *LnsMap
                _,metaTypeId2TypeInfoMap = TransUnit_convExp24743(Lns_2DDD(self.FP.processImport(dependName, depth + 1)))
                var typeId LnsInt
                typeId = Lns_forceCastInt((Lns_unwrap( dependInfo.GetAt("typeId"))))
                dependLibId2DependInfo.Set(typeId,NewTransUnit_DependModuleInfo(typeId, metaTypeId2TypeInfoMap))
            }
            _index24776, _key24776 = _sorted24776.NextFromMap( _index24776 )
        }
    }
    var typeId2TypeInfo *LnsMap
    typeId2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    typeId2TypeInfo.Set(Ast_rootTypeId,Ast_headTypeInfo)
    var typeId2Scope *LnsMap
    typeId2Scope = NewLnsMap( map[LnsAny]LnsAny{})
    typeId2Scope.Set(Ast_rootTypeId,self.scope)
    typeId2TypeInfo.Set(TransUnit_builtinFunc.Lnsthread_.FP.Get_typeId(),TransUnit_builtinFunc.Lnsthread_)
    typeId2Scope.Set(TransUnit_builtinFunc.Lnsthread_.FP.Get_typeId(),TransUnit_builtinFunc.Lnsthread_.FP.Get_scope())
    {
        _exp24917 := metaInfo.GetAt( "__dependIdMap" ).(*Lns_luaValue)
        _key24917, _val24917 := _exp24917.Get1stFromMap()
        for _key24917 != nil {
            typeId := _key24917.(LnsInt)
            dependIdInfo := _val24917.(*Lns_luaValue)
            var dependInfo *TransUnit_DependModuleInfo
            dependInfo = Lns_unwrap( dependLibId2DependInfo.Items[Lns_unwrap( dependIdInfo.GetAt(1)).(LnsInt)]).(*TransUnit_DependModuleInfo)
            var typeInfo *Ast_TypeInfo
            typeInfo = dependInfo.FP.GetTypeInfo(Lns_unwrap( dependIdInfo.GetAt(2)).(LnsInt))
            typeId2TypeInfo.Set(typeId,typeInfo)
            {
                __exp := Ast_getScope(typeInfo)
                if __exp != nil {
                    _exp := __exp.(*Ast_Scope)
                    typeId2Scope.Set(typeId,_exp)
                }
            }
            _key24917, _val24917 = _exp24917.NextFromMap( _key24917 )
        }
    }
    var moduleTypeInfo *Ast_TypeInfo
    moduleTypeInfo = Ast_headTypeInfo
    for _index, _moduleName := range( nameList.Items ) {
        index := _index + 1
        moduleName := _moduleName.(string)
        var mutable bool
        mutable = false
        if index == nameList.Len(){
            mutable = metaInfo.GetAt( "__moduleMutable" ).(bool)
            
        }
        moduleTypeInfo = self.FP.pushModule(true, moduleName, mutable)
        
    }
    for range( nameList.Items ) {
        self.FP.popModule()
    }
    for _, _symbolInfo := range( Ast_getSym2builtInTypeMap().Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        typeId2TypeInfo.Set(symbolInfo.FP.Get_typeInfo().FP.Get_typeId(),symbolInfo.FP.Get_typeInfo())
    }
    for _, _builtinTypeInfo := range( Ast_getBuiltInTypeIdMap().Items ) {
        builtinTypeInfo := _builtinTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        typeId2TypeInfo.Set(builtinTypeInfo.FP.Get_typeId(),builtinTypeInfo)
    }
    var newId2OldIdMap *LnsMap
    newId2OldIdMap = NewLnsMap( map[LnsAny]LnsAny{})
    var _typeInfoList *LnsList
    _typeInfoList = NewLnsList([]LnsAny{})
    var id2atomMap *LnsMap
    id2atomMap = NewLnsMap( map[LnsAny]LnsAny{})
    var _typeInfoNormalList *LnsList
    _typeInfoNormalList = NewLnsList([]LnsAny{})
    {
        _exp25514 := metaInfo.GetAt( "__typeInfoList" ).(*Lns_luaValue)
        _key25514, _val25514 := _exp25514.Get1stFromMap()
        for _key25514 != nil {
            atomInfoLua := _val25514.(*Lns_luaValue)
            var workAtomInfo LnsAny
            
            {
                _workAtomInfo := Lns_getVM().ExpandLuavalMap(atomInfoLua)
                if _workAtomInfo == nil{
                    self.FP.Error("illegal atomInfo")
                } else {
                    workAtomInfo = _workAtomInfo
                }
            }
            var atomInfo *LnsMap
            atomInfo = workAtomInfo.(*LnsMap)
            {
                _skind := atomInfo.Items["skind"]
                if _skind != nil {
                    skind := _skind
                    var actInfo LnsAny
                    actInfo = nil
                    var mess LnsAny
                    mess = nil
                    var kind LnsInt
                    kind = Lns_unwrap( Ast_SerializeKind__from(Lns_forceCastInt(skind))).(LnsInt)
                    if _switch25435 := kind; _switch25435 == Ast_SerializeKind__Enum {
                        actInfo, mess = TransUnit_convExp25147(Lns_2DDD(TransUnit__TypeInfoEnum__fromMap_2193_(atomInfo,nil)))
                        
                    } else if _switch25435 == Ast_SerializeKind__Alge {
                        actInfo, mess = TransUnit_convExp25170(Lns_2DDD(TransUnit__TypeInfoAlge__fromMap_2252_(atomInfo,nil)))
                        
                        self.helperInfo.UseAlge = true
                        
                    } else if _switch25435 == Ast_SerializeKind__Module {
                        actInfo, mess = TransUnit_convExp25202(Lns_2DDD(TransUnit__TypeInfoModule__fromMap_2070_(atomInfo,nil)))
                        
                    } else if _switch25435 == Ast_SerializeKind__Normal {
                        var workInfo LnsAny
                        workInfo, mess = TransUnit__TypeInfoNormal__fromMap_2162_(atomInfo,nil)
                        
                        if workInfo != nil{
                            workInfo_7564 := workInfo.(*TransUnit__TypeInfoNormal)
                            _typeInfoNormalList.Insert(TransUnit__TypeInfoNormal2Stem(workInfo_7564))
                        }
                        actInfo = TransUnit__TypeInfoDownCastF(workInfo)
                        
                    } else if _switch25435 == Ast_SerializeKind__Nilable {
                        actInfo, mess = TransUnit_convExp25271(Lns_2DDD(TransUnit__TypeInfoNilable__fromMap_1855_(atomInfo,nil)))
                        
                    } else if _switch25435 == Ast_SerializeKind__Alias {
                        actInfo, mess = TransUnit_convExp25294(Lns_2DDD(TransUnit__TypeInfoAlias__fromMap_1879_(atomInfo,nil)))
                        
                    } else if _switch25435 == Ast_SerializeKind__DDD {
                        actInfo, mess = TransUnit_convExp25317(Lns_2DDD(TransUnit__TypeInfoDDD__fromMap_1903_(atomInfo,nil)))
                        
                    } else if _switch25435 == Ast_SerializeKind__Alternate {
                        actInfo, mess = TransUnit_convExp25340(Lns_2DDD(TransUnit__TypeInfoAlternate__fromMap_1937_(atomInfo,nil)))
                        
                    } else if _switch25435 == Ast_SerializeKind__Generic {
                        actInfo, mess = TransUnit_convExp25363(Lns_2DDD(TransUnit__TypeInfoGeneric__fromMap_1971_(atomInfo,nil)))
                        
                    } else if _switch25435 == Ast_SerializeKind__Modifier {
                        actInfo, mess = TransUnit_convExp25386(Lns_2DDD(TransUnit__TypeInfoModifier__fromMap_2043_(atomInfo,nil)))
                        
                    } else if _switch25435 == Ast_SerializeKind__Box {
                        actInfo, mess = TransUnit_convExp25409(Lns_2DDD(TransUnit__TypeInfoBox__fromMap_1995_(atomInfo,nil)))
                        
                    } else if _switch25435 == Ast_SerializeKind__Ext {
                        actInfo, mess = TransUnit_convExp25432(Lns_2DDD(TransUnit__TypeInfoExt__fromMap_2019_(atomInfo,nil)))
                        
                    }
                    if actInfo != nil{
                        actInfo_7574 := actInfo.(*TransUnit__TypeInfo)
                        _typeInfoList.Insert(TransUnit__TypeInfo2Stem(actInfo_7574))
                        id2atomMap.Set(actInfo_7574.TypeId,actInfo_7574)
                    } else {
                        for _key, _val := range( atomInfo.Items ) {
                            key := _key.(string)
                            val := _val
                            Util_errorLog(Lns_getVM().String_format("table: %s:%s", []LnsAny{key, val}))
                        }
                        if mess != nil{
                            mess_7580 := mess.(string)
                            Util_errorLog(mess_7580)
                        }
                        Util_err(Lns_getVM().String_format("_TypeInfo.%s._fromMap error", []LnsAny{Ast_SerializeKind_getTxt( kind)}))
                    }
                }
            }
            _key25514, _val25514 = _exp25514.NextFromMap( _key25514 )
        }
    }
    var orgId2MacroTypeInfo *LnsMap
    orgId2MacroTypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    var lazyModuleSet *LnsSet
    lazyModuleSet = NewLnsSet([]LnsAny{})
    {
        _exp25554 := metaInfo.GetAt( "__lazyModuleList" ).(*Lns_luaValue)
        _key25554, _val25554 := _exp25554.Get1stFromMap()
        for _key25554 != nil {
            typeId := _val25554.(LnsInt)
            lazyModuleSet.Add(typeId)
            _key25554, _val25554 = _exp25554.NextFromMap( _key25554 )
        }
    }
    var importParam *TransUnit_ImportParam
    importParam = NewTransUnit_ImportParam(self.parser.FP.GetLastPos(), self.FP, self.processInfo, typeId2Scope, typeId2TypeInfo, NewLnsMap( map[LnsAny]LnsAny{}), lazyModuleSet, metaInfo, self.scope, moduleTypeInfo, self.scopeAccess, id2atomMap)
    for _, _atomInfo := range( _typeInfoList.Items ) {
        atomInfo := _atomInfo.(TransUnit__TypeInfoDownCast).ToTransUnit__TypeInfo()
        var newTypeInfo LnsAny
        var errMess LnsAny
        newTypeInfo,errMess = atomInfo.FP.CreateTypeInfoCache(importParam)
        {
            __exp := errMess
            if __exp != nil {
                _exp := __exp.(string)
                Util_err(Lns_getVM().String_format("Failed to createType -- %s: %s(%d): %s", []LnsAny{modulePath, Ast_SerializeKind_getTxt( atomInfo.Skind), atomInfo.TypeId, _exp}))
            }
        }
        if newTypeInfo != nil{
            newTypeInfo_7593 := newTypeInfo.(*Ast_TypeInfo)
            if newTypeInfo_7593.FP.Get_kind() == Ast_TypeInfoKind__Macro{
                orgId2MacroTypeInfo.Set(atomInfo.TypeId,newTypeInfo_7593)
            }
            if newTypeInfo_7593.FP.Get_kind() == Ast_TypeInfoKind__Set{
                self.helperInfo.UseSet = true
                
            }
            if newTypeInfo_7593.FP.Get_accessMode() == Ast_AccessMode__Global{
                if _switch25787 := newTypeInfo_7593.FP.Get_kind(); _switch25787 == Ast_TypeInfoKind__IF || _switch25787 == Ast_TypeInfoKind__Class {
                    self.globalScope.FP.AddClass(self.processInfo, newTypeInfo_7593.FP.Get_rawTxt(), nil, newTypeInfo_7593)
                } else if _switch25787 == Ast_TypeInfoKind__Func {
                    self.globalScope.FP.AddFunc(self.processInfo, nil, newTypeInfo_7593, Ast_AccessMode__Global, newTypeInfo_7593.FP.Get_staticFlag(), Ast_TypeInfo_isMut(newTypeInfo_7593))
                } else if _switch25787 == Ast_TypeInfoKind__Enum {
                    self.globalScope.FP.AddEnum(self.processInfo, Ast_AccessMode__Global, newTypeInfo_7593.FP.Get_rawTxt(), nil, newTypeInfo_7593)
                } else if _switch25787 == Ast_TypeInfoKind__Nilable {
                } else {
                    Util_err(Lns_getVM().String_format("%s: not support kind -- %s", []LnsAny{__func__, Ast_TypeInfoKind_getTxt( newTypeInfo_7593.FP.Get_kind())}))
                }
            }
        }
    }
    for _, _atomInfo := range( _typeInfoNormalList.Items ) {
        atomInfo := _atomInfo.(TransUnit__TypeInfoNormalDownCast).ToTransUnit__TypeInfoNormal()
        if atomInfo.Children.Len() > 0{
            var scope *Ast_Scope
            scope = Lns_unwrap( typeId2Scope.Items[atomInfo.TypeId]).(*Ast_Scope)
            for _, _childId := range( atomInfo.Children.Items ) {
                childId := _childId.(LnsInt)
                var typeInfo *Ast_TypeInfo
                
                {
                    _typeInfo := typeId2TypeInfo.Items[childId]
                    if _typeInfo == nil{
                        Util_err(Lns_getVM().String_format("not found childId -- %s, %d, %s(%d)", []LnsAny{modulePath, childId, atomInfo.Txt, atomInfo.TypeId}))
                    } else {
                        typeInfo = _typeInfo.(*Ast_TypeInfo)
                    }
                }
                var symbolKind LnsInt
                symbolKind = Ast_SymbolKind__Typ
                var addFlag bool
                addFlag = true
                if _switch25919 := typeInfo.FP.Get_kind(); _switch25919 == Ast_TypeInfoKind__Func {
                    symbolKind = Ast_SymbolKind__Fun
                    
                } else if _switch25919 == Ast_TypeInfoKind__Form || _switch25919 == Ast_TypeInfoKind__FormFunc {
                    symbolKind = Ast_SymbolKind__Typ
                    
                } else if _switch25919 == Ast_TypeInfoKind__Method {
                    symbolKind = Ast_SymbolKind__Mtd
                    
                } else if _switch25919 == Ast_TypeInfoKind__Class || _switch25919 == Ast_TypeInfoKind__Module {
                    symbolKind = Ast_SymbolKind__Typ
                    
                } else if _switch25919 == Ast_TypeInfoKind__Enum {
                    addFlag = false
                    
                }
                if addFlag{
                    scope.FP.Add(self.processInfo, symbolKind, false, typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Func, typeInfo.FP.GetTxt(nil, nil, nil), nil, typeInfo, typeInfo.FP.Get_accessMode(), typeInfo.FP.Get_staticFlag(), typeInfo.FP.Get_mutMode(), true, false)
                }
            }
        }
    }
    for _typeId, _typeInfo := range( typeId2TypeInfo.Items ) {
        typeId := _typeId.(LnsInt)
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        newId2OldIdMap.Set(typeInfo,typeId)
    }
    var registMember func(classTypeId LnsInt)
    registMember = func(classTypeId LnsInt) {
        __func__ := "@lune.@base.@TransUnit.TransUnit.processImport.registMember"
        if Lns_isCondTrue( metaInfo.GetAt( "__dependIdMap" ).(*Lns_luaValue).GetAt(classTypeId)){
            return 
        }
        var classTypeInfo *Ast_TypeInfo
        classTypeInfo = Lns_unwrap( typeId2TypeInfo.Items[classTypeId]).(*Ast_TypeInfo)
        if _switch26358 := (classTypeInfo.FP.Get_kind()); _switch26358 == Ast_TypeInfoKind__Class || _switch26358 == Ast_TypeInfoKind__ExtModule {
            self.FP.pushClassScope(self.parser.FP.GetLastPos(), classTypeInfo)
            {
                __exp := metaInfo.GetAt( "__typeId2ClassInfoMap" ).(*Lns_luaValue).GetAt(classTypeId)
                if __exp != nil {
                    _exp := __exp.(*Lns_luaValue)
                    var classInfo LnsAny
                    
                    {
                        _classInfo := Lns_getVM().ExpandLuavalMap(_exp)
                        if _classInfo == nil{
                            self.FP.Error("illegal val")
                        } else {
                            classInfo = _classInfo
                        }
                    }
                    for _fieldName, _fieldInfo := range( classInfo.(*LnsMap).Items ) {
                        fieldName := _fieldName.(string)
                        fieldInfo := _fieldInfo.(*LnsMap)
                        {
                            _typeId := fieldInfo.Items["typeId"]
                            if _typeId != nil {
                                typeId := _typeId
                                var fieldTypeInfo *Ast_TypeInfo
                                fieldTypeInfo = Lns_unwrap( typeId2TypeInfo.Items[Lns_forceCastInt(typeId)]).(*Ast_TypeInfo)
                                _ = TransUnit_convExp26152(Lns_2DDD(self.scope.FP.AddMember(self.processInfo, fieldName, nil, fieldTypeInfo, Lns_unwrap( Ast_AccessMode__from(Lns_forceCastInt((Lns_unwrap( fieldInfo.Items["accessMode"]))))).(LnsInt), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                    Lns_GetEnv().SetStackVal( fieldInfo.Items["staticFlag"]) &&
                                    Lns_GetEnv().SetStackVal( true) ||
                                    Lns_GetEnv().SetStackVal( false) ).(bool), Lns_unwrap( Ast_MutMode__from(Lns_forceCastInt((Lns_unwrap( fieldInfo.Items["mutMode"]))))).(LnsInt))))
                            } else {
                                self.FP.Error("not found fieldInfo.typeId")
                            }
                        }
                    }
                } else {
                    self.FP.Error(Lns_getVM().String_format("not found class -- %s: %d, %s", []LnsAny{modulePath, classTypeId, classTypeInfo.FP.GetTxt(nil, nil, nil)}))
                }
            }
        } else if _switch26358 == Ast_TypeInfoKind__Module {
            self.FP.pushModule(true, classTypeInfo.FP.GetTxt(nil, nil, nil), Ast_TypeInfo_isMut(classTypeInfo))
            Log_log(Log_Level__Debug, __func__, 3423, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("push module -- %s, %s, %d, %d, %d", []LnsAny{classTypeInfo.FP.GetTxt(nil, nil, nil), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.scope.FP.Get_ownerTypeInfo()) && 
                    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.GetFullName(Ast_defaultTypeNameCtrl, self.scope.FP, false)})/* 1:72 */)) ||
                    Lns_GetEnv().SetStackVal( "nil") ).(string), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.scope.FP.Get_ownerTypeInfo()) && 
                    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_typeId()}))) ||
                    Lns_GetEnv().SetStackVal( -1) ).(LnsInt), classTypeInfo.FP.Get_typeId(), self.scope.FP.Get_parent().FP.Get_scopeId()})
            }))
            
        }
        for _, _child := range( classTypeInfo.FP.Get_children().Items ) {
            child := _child.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( child.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( child.FP.Get_kind() == Ast_TypeInfoKind__ExtModule) ||
                Lns_GetEnv().SetStackVal( child.FP.Get_kind() == Ast_TypeInfoKind__Module) ||
                Lns_GetEnv().SetStackVal( child.FP.Get_kind() == Ast_TypeInfoKind__IF) ).(bool){
                var oldId LnsAny
                oldId = newId2OldIdMap.Items[child]
                if Lns_isCondTrue( oldId){
                    registMember(Lns_unwrap( oldId).(LnsInt))
                }
            }
        }
        if _switch26444 := classTypeInfo.FP.Get_kind(); _switch26444 == Ast_TypeInfoKind__Class || _switch26444 == Ast_TypeInfoKind__ExtModule {
            self.FP.popClass()
        } else if _switch26444 == Ast_TypeInfoKind__Module {
            self.FP.popModule()
        }
    }
    for _, _atomInfo := range( _typeInfoList.Items ) {
        atomInfo := _atomInfo.(TransUnit__TypeInfoDownCast).ToTransUnit__TypeInfo()
        {
            _workInfo := TransUnit__TypeInfoNormalDownCastF(atomInfo.FP)
            if _workInfo != nil {
                workInfo := _workInfo.(*TransUnit__TypeInfoNormal)
                if workInfo.ParentId == Ast_rootTypeId{
                    registMember(atomInfo.TypeId)
                }
            } else {
                {
                    _workInfo := TransUnit__TypeInfoModuleDownCastF(atomInfo.FP)
                    if _workInfo != nil {
                        workInfo := _workInfo.(*TransUnit__TypeInfoModule)
                        if workInfo.ParentId == Ast_rootTypeId{
                            registMember(atomInfo.TypeId)
                        }
                    }
                }
            }
        }
    }
    for _index, _moduleName := range( nameList.Items ) {
        index := _index + 1
        moduleName := _moduleName.(string)
        var mutable bool
        mutable = false
        if index == nameList.Len(){
            mutable = metaInfo.GetAt( "__moduleMutable" ).(bool)
            
        }
        self.FP.pushModule(true, moduleName, mutable)
    }
    {
        _exp26602 := metaInfo.GetAt( "__varName2InfoMap" ).(*Lns_luaValue)
        _key26602, _val26602 := _exp26602.Get1stFromMap()
        for _key26602 != nil {
            varName := _key26602.(string)
            varInfo := _val26602.(*Lns_luaValue)
            {
                _typeId := varInfo.GetAt("typeId")
                if _typeId != nil {
                    typeId := _typeId
                    self.scope.FP.AddStaticVar(self.processInfo, false, true, varName, nil, Lns_unwrap( typeId2TypeInfo.Items[Lns_forceCastInt(typeId)]).(*Ast_TypeInfo), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( varInfo.GetAt("mutable")) &&
                        Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
                        Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt))
                } else {
                    self.FP.Error("illegal varInfo.typeId")
                }
            }
            _key26602, _val26602 = _exp26602.NextFromMap( _key26602 )
        }
    }
    {
        _exp26634 := metaInfo.GetAt( "__macroName2InfoMap" ).(*Lns_luaValue)
        _key26634, _val26634 := _exp26634.Get1stFromMap()
        for _key26634 != nil {
            orgTypeId := _key26634.(LnsInt)
            macroInfoStem := _val26634
            self.macroCtrl.FP.ImportMacro(self.processInfo, Lns_getVM().ExpandLuavalMap(macroInfoStem), Lns_unwrap( orgId2MacroTypeInfo.Items[orgTypeId]).(*Ast_TypeInfo), typeId2TypeInfo)
            _key26634, _val26634 = _exp26634.NextFromMap( _key26634 )
        }
    }
    for range( nameList.Items ) {
        self.FP.popModule()
    }
    if depth == 1{
        for _key, _val := range( importParam.ImportedAliasMap.Items ) {
            key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            val := _val.(Ast_AliasTypeInfoDownCast).ToAst_AliasTypeInfo()
            self.importedAliasMap.Set(key,val)
        }
    }
    var moduleInfo *Nodes_ModuleInfo
    moduleInfo = NewNodes_ModuleInfo(modulePath, nameList.GetAt(nameList.Len()).(string), newId2OldIdMap, FrontInterface_ModuleId_createIdFromTxt(metaInfo.GetAt( "__buildId" ).(string)), importParam.ImportedAliasMap)
    self.importModule2ModuleInfo.Set(moduleTypeInfo,moduleInfo)
    self.importModuleName2ModuleInfo.Set(modulePath,moduleInfo)
    self.importModuleInfo.FP.Remove()
    Log_log(Log_Level__Info, __func__, 3523, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("%s complete", []LnsAny{modulePath})
    }))
    
    return metaInfo, typeId2TypeInfo, moduleInfo
}

// 3529: decl @lune.@base.@TransUnit.TransUnit.analyzeImportFor
func (self *TransUnit_TransUnit) analyzeImportFor(pos *Types_Position,modulePath string,assignName string,assigned bool,lazyLoad LnsInt) *Nodes_Node {
    var backupScope *Ast_Scope
    backupScope = self.scope
    self.scope = self.topScope
    
    self.processInfo.FP.SwitchIdProvier(Ast_IdType__Ext)
    var metaInfo *Lns_luaValue
    var typeId2TypeInfo *LnsMap
    var moduleInfo *Nodes_ModuleInfo
    metaInfo,typeId2TypeInfo,moduleInfo = self.FP.processImport(modulePath, 1)
    self.processInfo.FP.SwitchIdProvier(Ast_IdType__Base)
    self.scope = backupScope
    
    var moduleTypeInfo *Ast_TypeInfo
    moduleTypeInfo = Lns_unwrap( typeId2TypeInfo.Items[metaInfo.GetAt( "__moduleTypeId" ).(LnsInt)]).(*Ast_TypeInfo)
    self.scope.FP.AddModule(moduleTypeInfo, moduleInfo.FP.Assign(assignName).FP)
    var moduleSymbolKind LnsInt
    moduleSymbolKind = Lns_unwrap( Ast_SymbolKind__from(metaInfo.GetAt( "__moduleSymbolKind" ).(LnsInt))).(LnsInt)
    var moduleSymbolInfo LnsAny
    var shadowing LnsAny
    moduleSymbolInfo,shadowing = self.scope.FP.Add(self.processInfo, moduleSymbolKind, false, false, assignName, pos, moduleTypeInfo, Ast_AccessMode__Local, true, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( metaInfo.GetAt( "__moduleMutable" ).(bool)) &&
        Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
        Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, lazyLoad != Nodes_LazyLoad__Off)
    if moduleSymbolInfo != nil{
        moduleSymbolInfo_7703 := moduleSymbolInfo.(*Ast_SymbolInfo)
        return &Nodes_ImportNode_create(self.nodeManager, pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(moduleTypeInfo)}), modulePath, lazyLoad, assignName, assigned, moduleSymbolInfo_7703, moduleTypeInfo).Nodes_Node
    }
    self.FP.errorShadowing(pos, shadowing)
    return self.FP.createNoneNode(pos)
}

// 3580: decl @lune.@base.@TransUnit.TransUnit.analyzeImport
func (self *TransUnit_TransUnit) analyzeImport(opeToken *Types_Token) *Nodes_Node {
    var lazyLoad LnsInt
    if self.FP.getToken(nil).Txt == "."{
        var modeToken *Types_Token
        modeToken = self.FP.getToken(nil)
        if _switch27111 := modeToken.Txt; _switch27111 == "l" {
            lazyLoad = Nodes_LazyLoad__On
            
        } else if _switch27111 == "d" {
            lazyLoad = Nodes_LazyLoad__Off
            
        } else {
            lazyLoad = Nodes_LazyLoad__Off
            
            self.FP.Error(Lns_getVM().String_format("illegal import mode -- %s", []LnsAny{modeToken.Txt}))
        }
    } else { 
        self.FP.Pushback()
        if self.ctrl_info.DefaultLazy{
            lazyLoad = Nodes_LazyLoad__Auto
            
        } else { 
            lazyLoad = Nodes_LazyLoad__Off
            
        }
    }
    if lazyLoad != Nodes_LazyLoad__Off{
        self.helperInfo.UseLazyLoad = true
        
    }
    var moduleToken *Types_Token
    moduleToken = self.FP.getToken(nil)
    var modulePath string
    modulePath = moduleToken.Txt
    var nextToken *Types_Token
    nextToken = moduleToken
    for  {
        nextToken = self.FP.getToken(nil)
        
        if _switch27250 := nextToken.Txt; _switch27250 == "." || _switch27250 == "/" || _switch27250 == ":" {
            var demilit string
            demilit = nextToken.Txt
            nextToken = self.FP.getToken(nil)
            
            moduleToken = nextToken
            
            modulePath = Lns_getVM().String_format("%s%s%s", []LnsAny{modulePath, demilit, moduleToken.Txt})
            
        } else {
            break
        }
    }
    var assignName *Types_Token
    assignName = moduleToken
    var assigned bool
    if nextToken.Txt == "as"{
        assignName = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
        
        nextToken = self.FP.getToken(nil)
        
        assigned = true
        
    } else { 
        assigned = false
        
    }
    self.FP.checkToken(nextToken, ";")
    var node *Nodes_Node
    node = self.FP.analyzeImportFor(opeToken.Pos, modulePath, assignName.Txt, assigned, lazyLoad)
    self.importModuleSet.Add(Ast_TypeInfo2Stem(node.FP.Get_expType()))
    return node
}

// 3646: decl @lune.@base.@TransUnit.TransUnit.analyzeTestCase
func (self *TransUnit_TransUnit) analyzeTestCase(firstToken *Types_Token) *Nodes_TestCaseNode {
    var newScope *Ast_Scope
    newScope = self.FP.pushScope(false, nil, nil)
    var importNode *Nodes_Node
    importNode = self.FP.analyzeImportFor(firstToken.Pos, "lune.base.Testing", "__t", false, Nodes_LazyLoad__Off)
    var nameToken *Types_Token
    nameToken = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
    self.FP.checkNextToken("(")
    var ctrlToken *Types_Token
    ctrlToken = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
    var ctrlName string
    ctrlName = ctrlToken.Txt
    self.FP.checkNextToken(")")
    var moduleType *Ast_TypeInfo
    moduleType = importNode.FP.Get_expType()
    var ctrlType *Ast_TypeInfo
    
    {
        _ctrlType := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(moduleType.FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.GetTypeInfoChild("Ctrl")})/* 3662:20 */)
        if _ctrlType == nil{
            self.FP.Error("not found Testing.Ctrl class")
        } else {
            ctrlType = _ctrlType.(*Ast_TypeInfo)
        }
    }
    self.FP.addLocalVar(ctrlToken.Pos, true, false, ctrlToken.Txt, ctrlType, Ast_MutMode__IMut, false)
    self.scopeAccess = Ast_ScopeAccess__Full
    
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(Nodes_BlockKind__Test, TransUnit_TentativeMode__Ignore, newScope, nil)
    self.scopeAccess = Ast_ScopeAccess__Normal
    
    self.FP.popScope()
    return Nodes_TestCaseNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), nameToken, importNode, ctrlName, block)
}

// 3682: decl @lune.@base.@TransUnit.TransUnit.analyzeTest
func (self *TransUnit_TransUnit) analyzeTest(firstToken *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    if nextToken.Txt != "{"{
        self.FP.Pushback()
        return &self.FP.analyzeTestCase(firstToken).Nodes_Node
    }
    self.FP.checkToken(nextToken, "{")
    var stmtList *LnsList
    stmtList = NewLnsList([]LnsAny{})
    self.FP.analyzeStatementList(stmtList, "}")
    self.FP.checkNextToken("}")
    return &Nodes_TestBlockNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), stmtList).Nodes_Node
}

// 3701: decl @lune.@base.@TransUnit.TransUnit.analyzeSubfile
func (self *TransUnit_TransUnit) analyzeSubfile(token *Types_Token) *Nodes_SubfileNode {
    if self.scope != self.moduleScope{
        self.FP.Error("'module' must be top scope.")
    }
    var mode *Types_Token
    mode = self.FP.getToken(nil)
    var moduleName string
    moduleName = ""
    for  {
        var nextToken *Types_Token
        nextToken = self.FP.getToken(nil)
        if nextToken.Txt == ";"{
            break
        }
        if moduleName == ""{
            moduleName = nextToken.Txt
            
        } else { 
            moduleName = Lns_getVM().String_format("%s%s", []LnsAny{moduleName, nextToken.Txt})
            
        }
    }
    var usePath LnsAny
    usePath = nil
    if moduleName == ""{
        self.FP.addErrMess(token.Pos, "illegal subfile")
    } else { 
        if mode.Txt == "use"{
            usePath = moduleName
            
            if Lns_isCondTrue( FrontInterface_searchModule(moduleName)){
                self.subfileList.Insert(moduleName)
            } else { 
                self.FP.addErrMess(token.Pos, Lns_getVM().String_format("not found subfile -- %s", []LnsAny{moduleName}))
            }
        } else if mode.Txt == "owner"{
            if self.moduleName != moduleName{
                self.FP.addErrMess(token.Pos, Lns_getVM().String_format("illegal owner module -- %s, %s", []LnsAny{moduleName, self.moduleName}))
            }
        } else { 
            self.FP.addErrMess(mode.Pos, Lns_getVM().String_format("illegal module mode -- %s", []LnsAny{mode.Txt}))
        }
    }
    return Nodes_SubfileNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), usePath)
}

// 3754: decl @lune.@base.@TransUnit.TransUnit.analyzeEnvLock
func (self *TransUnit_TransUnit) analyzeEnvLock(token *Types_Token) *Nodes_Node {
    return &Nodes_EnvNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), self.FP.analyzeBlock(Nodes_BlockKind__Env, TransUnit_TentativeMode__Simple, nil, nil)).Nodes_Node
}

// 3760: decl @lune.@base.@TransUnit.TransUnit.analyzeIf
func (self *TransUnit_TransUnit) analyzeIf(token *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    var continueFlag bool
    nextToken,continueFlag = self.FP.getContinueToken()
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( continueFlag) &&
        Lns_GetEnv().SetStackVal( nextToken.Txt == "!") ).(bool)){
        return &self.FP.analyzeIfUnwrap(token).Nodes_Node
    }
    self.FP.Pushback()
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    var ifExp *Nodes_Node
    ifExp = self.FP.analyzeExpOneRVal(false, false, nil, nil)
    list.Insert(Nodes_IfStmtInfo2Stem(NewNodes_IfStmtInfo(Nodes_IfKind__If, ifExp, self.FP.analyzeBlock(Nodes_BlockKind__If, TransUnit_TentativeMode__Start, nil, nil))))
    var checkCond func(condExp *Nodes_Node)
    checkCond = func(condExp *Nodes_Node) {
        if _switch28159 := condExp.FP.Get_expType().FP.Get_kind(); _switch28159 == Ast_TypeInfoKind__Nilable || _switch28159 == Ast_TypeInfoKind__Stem {
        } else if _switch28159 == Ast_TypeInfoKind__Prim {
            if Lns_op_not(condExp.FP.Get_expType().FP.Equals(self.processInfo, Ast_builtinTypeBool, nil, nil)){
                self.FP.addErrMess(condExp.FP.Get_pos(), Lns_getVM().String_format("This exp never be false -- %s", []LnsAny{condExp.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
            }
        } else {
            self.FP.addErrMess(condExp.FP.Get_pos(), Lns_getVM().String_format("This exp never be false -- %s", []LnsAny{condExp.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
        }
    }
    checkCond(ifExp)
    nextToken = self.FP.getToken(true)
    
    if nextToken.Txt == "elseif"{
        for nextToken.Txt == "elseif" {
            var condExp *Nodes_Node
            condExp = self.FP.analyzeExpOneRVal(false, false, nil, nil)
            list.Insert(Nodes_IfStmtInfo2Stem(NewNodes_IfStmtInfo(Nodes_IfKind__ElseIf, condExp, self.FP.analyzeBlock(Nodes_BlockKind__Elseif, TransUnit_TentativeMode__Merge, nil, nil))))
            checkCond(condExp)
            nextToken = self.FP.getToken(true)
            
        }
    }
    if nextToken.Txt == "else"{
        list.Insert(Nodes_IfStmtInfo2Stem(NewNodes_IfStmtInfo(Nodes_IfKind__Else, self.FP.createNoneNode(nextToken.Pos), self.FP.analyzeBlock(Nodes_BlockKind__Else, TransUnit_TentativeMode__Finish, nil, nil))))
    } else { 
        self.FP.finishTentativeSymbol(false)
        self.FP.Pushback()
    }
    return &Nodes_IfNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), list).Nodes_Node
}

// 3822: decl @lune.@base.@TransUnit.TransUnit.processCaseDefault
func (self *TransUnit_TransUnit) processCaseDefault(firstToken *Types_Token,caseKind LnsInt,nextToken *Types_Token,hasCase bool)(LnsAny, bool) {
    var keyword string
    keyword = TransUnit_convExp28383(Lns_2DDD(Lns_getVM().String_gsub(firstToken.Txt,"_", "")))
    var fullKeyword string
    fullKeyword = Lns_getVM().String_format("_%s", []LnsAny{keyword})
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( firstToken.Txt == fullKeyword) &&
        Lns_GetEnv().SetStackVal( caseKind != Nodes_CaseKind__MustFull) ).(bool)){
        self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("This '%s' hasn't enough 'case' condition.", []LnsAny{keyword}))
    }
    var defaultBlock LnsAny
    defaultBlock = nil
    var failSafeDefault bool
    failSafeDefault = false
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( nextToken.Txt == "default") ||
        Lns_GetEnv().SetStackVal( nextToken.Txt == "_default") ).(bool){
        if firstToken.Txt == fullKeyword{
            self.FP.addErrMess(nextToken.Pos, Lns_getVM().String_format("'_%s' can't have default.", []LnsAny{keyword}))
        }
        if nextToken.Txt == "_default"{
            failSafeDefault = true
            
        } else if caseKind == Nodes_CaseKind__Full{
            self.FP.addWarnMess(nextToken.Pos, Lns_getVM().String_format("This '%s' has full case. This 'default' is no reach.", []LnsAny{keyword}))
        }
        defaultBlock = self.FP.analyzeBlock(Nodes_BlockKind__Default, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(hasCase)) &&
            Lns_GetEnv().SetStackVal( TransUnit_TentativeMode__Simple) ||
            Lns_GetEnv().SetStackVal( TransUnit_TentativeMode__Finish) ).(LnsInt), nil, nil)
        
    } else { 
        if hasCase{
            self.FP.finishTentativeSymbol(caseKind != Nodes_CaseKind__Lack)
        }
        self.FP.Pushback()
    }
    self.FP.checkNextToken("}")
    if Lns_op_not(hasCase){
        self.FP.addWarnMess(firstToken.Pos, Lns_getVM().String_format("'%s' should have 'case' blocks.", []LnsAny{keyword}))
        if Lns_isCondTrue( defaultBlock){
            self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("'%s' must have 'case' blocks when have 'default' block.", []LnsAny{keyword}))
        }
    }
    return defaultBlock, failSafeDefault
}

// 3874: decl @lune.@base.@TransUnit.TransUnit.analyzeSwitch
func (self *TransUnit_TransUnit) analyzeSwitch(firstToken *Types_Token) *Nodes_SwitchNode {
    var exp *Nodes_Node
    exp = self.FP.analyzeExpOneRVal(false, false, nil, nil)
    self.FP.checkNextToken("{")
    var caseList *LnsList
    caseList = NewLnsList([]LnsAny{})
    var condObjSet *LnsSet
    condObjSet = NewLnsSet([]LnsAny{})
    var condSymIdSet *LnsSet
    condSymIdSet = NewLnsSet([]LnsAny{})
    var hasNilCond bool
    hasNilCond = false
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var firstFlag bool
    firstFlag = true
    for (nextToken.Txt == "case") {
        self.FP.checkToken(nextToken, "case")
        var condexpList *Nodes_ExpListNode
        condexpList = self.FP.analyzeExpList(false, false, false, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp.FP.Get_expType())}), true)
        var condBock *Nodes_BlockNode
        condBock = self.FP.analyzeBlock(Nodes_BlockKind__Switch, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( firstFlag) &&
            Lns_GetEnv().SetStackVal( TransUnit_TentativeMode__Start) ||
            Lns_GetEnv().SetStackVal( TransUnit_TentativeMode__Merge) ).(LnsInt), nil, nil)
        if firstFlag{
            firstFlag = false
            
        }
        for _, _condExp := range( condexpList.FP.Get_expList().Items ) {
            condExp := _condExp.(Nodes_NodeDownCast).ToNodes_Node()
            if condExp.FP.Get_expType().FP.Get_nilable(){
                if hasNilCond{
                    self.FP.addWarnMess(condExp.FP.Get_pos(), "multiple case with nil or nilable")
                } else { 
                    hasNilCond = true
                    
                }
            }
            {
                _condLiteral := TransUnit_convExp28928(Lns_2DDD(condExp.FP.GetLiteral()))
                if _condLiteral != nil {
                    condLiteral := _condLiteral
                    var literalObj LnsAny
                    literalObj = Nodes_getLiteralObj(condLiteral)
                    if literalObj != nil{
                        literalObj_7850 := literalObj
                        if condObjSet.Has(literalObj_7850){
                            self.FP.addErrMess(condExp.FP.Get_pos(), Lns_getVM().String_format("multiple case exp -- %s", []LnsAny{literalObj_7850}))
                        } else { 
                            condObjSet.Add(literalObj_7850)
                        }
                    }
                } else {
                    {
                        _expRef := Nodes_ExpRefNodeDownCastF(condExp.FP)
                        if _expRef != nil {
                            expRef := _expRef.(*Nodes_ExpRefNode)
                            var symInfo *Ast_SymbolInfo
                            symInfo = expRef.FP.Get_symbolInfo()
                            if condSymIdSet.Has(symInfo.FP.Get_symbolId()){
                                self.FP.addErrMess(condExp.FP.Get_pos(), Lns_getVM().String_format("multiple case exp -- %s", []LnsAny{symInfo.FP.Get_name()}))
                            } else { 
                                condSymIdSet.Add(symInfo.FP.Get_symbolId())
                            }
                        }
                    }
                }
            }
            if Lns_op_not(Lns_car(exp.FP.Get_expType().FP.CanEvalWith(self.processInfo, condExp.FP.Get_expType(), Ast_CanEvalType__Equal, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)){
                self.FP.addErrMess(condExp.FP.Get_pos(), Lns_getVM().String_format("case exp unmatch type -- %s <- %s", []LnsAny{exp.FP.Get_expType().FP.GetTxt(nil, nil, nil), condExp.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
            }
        }
        caseList.Insert(Nodes_CaseInfo2Stem(NewNodes_CaseInfo(condexpList, condBock)))
        nextToken = self.FP.getToken(nil)
        
    }
    var caseKind LnsInt
    {
        _enumType := Ast_EnumTypeInfoDownCastF(exp.FP.Get_expType().FP.Get_srcTypeInfo().FP.Get_aliasSrc().FP)
        if _enumType != nil {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            var miss bool
            miss = false
            for _, _enumVal := range( enumType.FP.Get_name2EnumValInfo().Items ) {
                enumVal := _enumVal.(Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
                if Lns_op_not(condObjSet.Has(Ast_getEnumLiteralVal(enumVal.FP.Get_val()))){
                    miss = true
                    
                    break
                }
            }
            if Lns_op_not(miss){
                if firstToken.Txt == "_switch"{
                    caseKind = Nodes_CaseKind__MustFull
                    
                } else { 
                    caseKind = Nodes_CaseKind__Full
                    
                }
            } else { 
                caseKind = Nodes_CaseKind__Lack
                
            }
        } else {
            caseKind = Nodes_CaseKind__Lack
            
            if firstToken.Txt == "_switch"{
                self.FP.addErrMess(exp.FP.Get_pos(), "The condition of '_switch' must be enum.")
            }
        }
    }
    var defaultBlock LnsAny
    var failSafeDefault bool
    defaultBlock,failSafeDefault = self.FP.processCaseDefault(firstToken, caseKind, nextToken, caseList.Len() != 0)
    return Nodes_SwitchNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), exp, caseList, defaultBlock, caseKind, failSafeDefault)
}

// 3983: decl @lune.@base.@TransUnit.TransUnit.analyzeMatch
func (self *TransUnit_TransUnit) analyzeMatch(firstToken *Types_Token) *Nodes_MatchNode {
    var exp *Nodes_Node
    exp = self.FP.analyzeExpOneRVal(false, false, nil, nil)
    var algeTypeInfo *Ast_AlgeTypeInfo
    
    {
        _algeTypeInfo := Ast_AlgeTypeInfoDownCastF(exp.FP.Get_expType().FP.Get_srcTypeInfo().FP)
        if _algeTypeInfo == nil{
            self.FP.Error(Lns_getVM().String_format("match must have alge value -- %s", []LnsAny{exp.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
        } else {
            algeTypeInfo = _algeTypeInfo.(*Ast_AlgeTypeInfo)
        }
    }
    self.FP.checkNextToken("{")
    var caseList *LnsList
    caseList = NewLnsList([]LnsAny{})
    var algeValNameSet *LnsSet
    algeValNameSet = NewLnsSet([]LnsAny{})
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var firstFlag bool
    firstFlag = true
    for (nextToken.Txt == "case") {
        self.FP.checkNextToken(".")
        var valNameToken *Types_Token
        valNameToken = self.FP.getToken(nil)
        self.FP.checkAlgeComp(valNameToken, algeTypeInfo)
        var valInfo *Ast_AlgeValInfo
        
        {
            _valInfo := algeTypeInfo.FP.GetValInfo(valNameToken.Txt)
            if _valInfo == nil{
                self.FP.Error(Lns_getVM().String_format("not found val -- %s", []LnsAny{valNameToken.Txt}))
            } else {
                valInfo = _valInfo.(*Ast_AlgeValInfo)
            }
        }
        if algeValNameSet.Has(valNameToken.Txt){
            self.FP.addErrMess(valNameToken.Pos, Lns_getVM().String_format("multiple val -- %s", []LnsAny{valNameToken.Txt}))
        }
        algeValNameSet.Add(valNameToken.Txt)
        var valParamNameList *LnsList
        valParamNameList = NewLnsList([]LnsAny{})
        nextToken = self.FP.getToken(nil)
        
        var blockScope *Ast_Scope
        blockScope = self.FP.pushScope(false, nil, nil)
        if nextToken.Txt == "("{
            for _, _paramType := range( valInfo.FP.Get_typeList().Items ) {
                paramType := _paramType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var paramName *Types_Token
                paramName = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_Or_)
                self.FP.checkShadowing(paramName.Pos, paramName.Txt, self.scope)
                var workType *Ast_TypeInfo
                workType = paramType
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(paramType)) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(exp.FP.Get_expType()))) ).(bool)){
                    workType = self.FP.CreateModifier(workType, Ast_MutMode__IMut)
                    
                }
                valParamNameList.Insert(Ast_SymbolInfo2Stem(self.FP.addLocalVar(paramName.Pos, false, false, paramName.Txt, workType, Ast_MutMode__IMut, nil)))
                nextToken = self.FP.getToken(nil)
                
                if nextToken.Txt != ","{
                    break
                }
            }
            self.FP.checkToken(nextToken, ")")
        } else { 
            self.FP.Pushback()
        }
        if valParamNameList.Len() != valInfo.FP.Get_typeList().Len(){
            self.FP.addErrMess(valNameToken.Pos, Lns_getVM().String_format("unmatch param -- %d != %d", []LnsAny{valParamNameList.Len(), valInfo.FP.Get_typeList().Len()}))
        }
        var mode LnsInt
        mode = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( firstFlag) &&
            Lns_GetEnv().SetStackVal( TransUnit_TentativeMode__Start) ||
            Lns_GetEnv().SetStackVal( TransUnit_TentativeMode__Merge) ).(LnsInt)
        var block *Nodes_BlockNode
        block = self.FP.analyzeBlock(Nodes_BlockKind__Match, mode, blockScope, nil)
        if firstFlag{
            firstFlag = false
            
        }
        self.FP.popScope()
        var matchCase *Nodes_MatchCase
        matchCase = NewNodes_MatchCase(valInfo, valParamNameList, block)
        caseList.Insert(Nodes_MatchCase2Stem(matchCase))
        nextToken = self.FP.getToken(nil)
        
    }
    var caseKind LnsInt
    if algeValNameSet.Len() == algeTypeInfo.FP.Get_valInfoNum(){
        if firstToken.Txt == "_match"{
            caseKind = Nodes_CaseKind__MustFull
            
        } else { 
            caseKind = Nodes_CaseKind__Full
            
        }
    } else { 
        caseKind = Nodes_CaseKind__Lack
        
    }
    var defaultBlock LnsAny
    var failSafeDefault bool
    defaultBlock,failSafeDefault = self.FP.processCaseDefault(firstToken, caseKind, nextToken, caseList.Len() != 0)
    return Nodes_MatchNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), exp, algeTypeInfo, caseList, Nodes_NodeDownCastF(defaultBlock), caseKind, failSafeDefault)
}

// 4078: decl @lune.@base.@TransUnit.TransUnit.analyzeWhile
func (self *TransUnit_TransUnit) analyzeWhile(token *Types_Token) *Nodes_WhileNode {
    var refAccessSymPosList *LnsList
    refAccessSymPosList = TransUnit_clearThePosForModToRef_1103_(self.scope, self.moduleScope)
    var cond *Nodes_Node
    cond = self.FP.analyzeExpOneRVal(false, false, nil, nil)
    var infinit bool
    infinit = false
    if cond.FP.Get_expType() == Ast_builtinTypeBool{
        {
            _literal := TransUnit_convExp29876(Lns_2DDD(cond.FP.GetLiteral()))
            if _literal != nil {
                literal := _literal
                switch _exp29874 := literal.(type) {
                case *Nodes_Literal__Bool:
                val := _exp29874.Val1
                    infinit = val
                    
                }
            }
        }
    } else if Lns_op_not(cond.FP.Get_expType().FP.Get_nilable()){
        infinit = true
        
    }
    return Nodes_WhileNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), cond, infinit, self.FP.analyzeBlock(Nodes_BlockKind__While, TransUnit_TentativeMode__Loop, nil, refAccessSymPosList))
}

// 4102: decl @lune.@base.@TransUnit.TransUnit.analyzeRepeat
func (self *TransUnit_TransUnit) analyzeRepeat(token *Types_Token) *Nodes_RepeatNode {
    var scope *Ast_Scope
    scope = self.FP.pushScope(false, nil, nil)
    var node *Nodes_RepeatNode
    node = Nodes_RepeatNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), self.FP.analyzeBlock(Nodes_BlockKind__Repeat, TransUnit_TentativeMode__Loop, scope, nil), self.FP.analyzeExpOneRVal(false, false, nil, nil))
    self.FP.finishTentativeSymbol(false)
    self.FP.popScope()
    self.FP.checkNextToken(";")
    return node
}

// 4121: decl @lune.@base.@TransUnit.TransUnit.analyzeFor
func (self *TransUnit_TransUnit) analyzeFor(firstToken *Types_Token) *Nodes_ForNode {
    var scope *Ast_Scope
    scope = self.FP.pushScope(false, nil, nil)
    var val *Types_Token
    val = self.FP.getToken(nil)
    if val.Kind != Types_TokenKind__Symb{
        self.FP.Error("not symbol")
    }
    self.FP.checkNextToken("=")
    var exp1 *Nodes_Node
    exp1 = self.FP.analyzeExpOneRVal(false, false, nil, nil)
    if Lns_op_not(Ast_isNumberType(exp1.FP.Get_expType())){
        self.FP.addErrMess(exp1.FP.Get_pos(), Lns_getVM().String_format("exp1 is not number -- %s", []LnsAny{exp1.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
    }
    self.FP.checkNextToken(",")
    var exp2 *Nodes_Node
    exp2 = self.FP.analyzeExpOneRVal(false, false, nil, nil)
    if Lns_op_not(Ast_isNumberType(exp2.FP.Get_expType())){
        self.FP.addErrMess(exp2.FP.Get_pos(), Lns_getVM().String_format("exp2 is not number -- %s", []LnsAny{exp2.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
    }
    var token *Types_Token
    token = self.FP.getToken(nil)
    var exp3 LnsAny
    exp3 = nil
    if token.Txt == ","{
        exp3 = self.FP.analyzeExpOneRVal(false, false, nil, nil)
        
        {
            __exp := exp3
            if __exp != nil {
                _exp := __exp.(*Nodes_Node)
                if Lns_op_not(Ast_isNumberType(_exp.FP.Get_expType())){
                    self.FP.addErrMess(_exp.FP.Get_pos(), Lns_getVM().String_format("exp is not number -- %s", []LnsAny{_exp.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
                }
            }
        }
    } else { 
        self.FP.Pushback()
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( exp1.FP.Get_expType() == Ast_builtinTypeInt) &&
        Lns_GetEnv().SetStackVal( exp2.FP.Get_expType() == Ast_builtinTypeReal) ||
        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(exp3) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_Node).FP.Get_expType()})) == Ast_builtinTypeReal) ).(bool){
        exp1 = &Nodes_ExpCastNode_create(self.nodeManager, exp1.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeReal)}), exp1, Ast_builtinTypeReal, Nodes_CastKind__Force).Nodes_Node
        
    }
    var symbolInfo *Ast_SymbolInfo
    symbolInfo = self.FP.addLocalVar(val.Pos, false, true, val.Txt, exp1.FP.Get_expType(), Ast_MutMode__IMut, nil)
    var node *Nodes_ForNode
    node = Nodes_ForNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), self.FP.analyzeBlock(Nodes_BlockKind__For, TransUnit_TentativeMode__Loop, scope, nil), symbolInfo, exp1, exp2, exp3)
    self.FP.popScope()
    return node
}

// 4181: decl @lune.@base.@TransUnit.TransUnit.analyzeApply
func (self *TransUnit_TransUnit) analyzeApply(token *Types_Token) *Nodes_ApplyNode {
    var scope *Ast_Scope
    scope = self.FP.pushScope(false, nil, nil)
    var varList *LnsList
    varList = NewLnsList([]LnsAny{})
    var nextToken *Types_Token
    nextToken = Parser_getEofToken()
    for {
        var _var *Types_Token
        _var = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_Or_)
        if _var.Kind != Types_TokenKind__Symb{
            self.FP.Error("illegal symbol")
        }
        varList.Insert(Types_Token2Stem(_var))
        nextToken = self.FP.getToken(nil)
        
        if nextToken.Txt != ","{ break }
    }
    self.FP.checkToken(nextToken, "of")
    var expListNode *Nodes_ExpListNode
    expListNode = self.FP.analyzeExpList(false, false, false, nil, nil, nil)
    var itFunc *Ast_TypeInfo
    itFunc = Ast_builtinTypeNone
    var expTypeList *LnsList
    expTypeList = expListNode.FP.Get_expTypeList()
    if expTypeList.Len() < 3{
        self.FP.addErrMess(expListNode.FP.Get_pos(), Lns_getVM().String_format("apply must have 3 values -- %s", []LnsAny{expTypeList.Len()}))
    } else { 
        itFunc = expTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_extedType()
        
    }
    var itemTypeList *LnsList
    itemTypeList = NewLnsList([]LnsAny{})
    var defaultItemType *Ast_TypeInfo
    defaultItemType = Ast_builtinTypeStem_
    var readyFlag bool
    readyFlag = false
    {
        _callNode := Nodes_ExpCallNodeDownCastF(Nodes_getUnwraped(expListNode.FP.Get_expList().GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()).FP)
        if _callNode != nil {
            callNode := _callNode.(*Nodes_ExpCallNode)
            var callFuncType *Ast_TypeInfo
            callFuncType = callNode.FP.Get_func().FP.Get_expType()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( callFuncType.FP.Equals(self.processInfo, TransUnit_builtinFunc.Str_gmatch, nil, nil)) ||
                Lns_GetEnv().SetStackVal( callFuncType.FP.Equals(self.processInfo, TransUnit_builtinFunc.String_gmatch, nil, nil)) ).(bool){
                itemTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeString))
                defaultItemType = Ast_builtinTypeString.FP.Get_nilableTypeInfo()
                
                readyFlag = true
                
            }
        }
    }
    if Lns_op_not(readyFlag){
        if _switch30736 := itFunc.FP.Get_kind(); _switch30736 == Ast_TypeInfoKind__Func || _switch30736 == Ast_TypeInfoKind__FormFunc || _switch30736 == Ast_TypeInfoKind__Form {
        } else {
            self.FP.addErrMess(expListNode.FP.Get_pos(), Lns_getVM().String_format("The 1st value must be iterator function. -- %s", []LnsAny{itFunc.FP.GetTxt(nil, nil, nil)}))
        }
        if itFunc.FP.Get_argTypeInfoList().Len() != 2{
            self.FP.addErrMess(expListNode.FP.Get_pos(), Lns_getVM().String_format("iterator function must has two arguments. -- %s", []LnsAny{itFunc.FP.GetTxt(nil, nil, nil)}))
        } else { 
            var arg2Type *Ast_TypeInfo
            arg2Type = itFunc.FP.Get_argTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(arg2Type.FP.Get_nilable()){
                self.FP.addErrMess(expListNode.FP.Get_pos(), Lns_getVM().String_format("the 2nd argument of iterator function must be nilable. -- %s", []LnsAny{itFunc.FP.GetTxt(nil, nil, nil)}))
            }
        }
        if itFunc.FP.Get_retTypeInfoList().Len() == 0{
            self.FP.addErrMess(expListNode.FP.Get_pos(), "iterator function must return value.")
        } else { 
            var iteRetType *Ast_TypeInfo
            iteRetType = itFunc.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(iteRetType.FP.Get_nilable()){
                self.FP.addErrMess(expListNode.FP.Get_pos(), "iterator function must return nilable type at 1st.")
            }
        }
        for _index, _itemType := range( itFunc.FP.Get_retTypeInfoList().Items ) {
            index := _index + 1
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var workType *Ast_TypeInfo
            workType = itemType
            if index == 1{
                if itemType.FP.Get_nilable(){
                    workType = workType.FP.Get_nonnilableType()
                    
                }
            }
            itemTypeList.Insert(Ast_TypeInfo2Stem(workType))
        }
    }
    var varSymList *LnsList
    varSymList = NewLnsList([]LnsAny{})
    for _index, __var := range( varList.Items ) {
        index := _index + 1
        _var := __var.(Types_TokenDownCast).ToTypes_Token()
        var itemType *Ast_TypeInfo
        itemType = defaultItemType
        if index <= itemTypeList.Len(){
            itemType = itemTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            
        }
        varSymList.Insert(Ast_SymbolInfo2Stem(self.FP.addLocalVar(_var.Pos, false, true, _var.Txt, itemType, Ast_MutMode__IMut, nil)))
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(Nodes_BlockKind__Apply, TransUnit_TentativeMode__Loop, scope, nil)
    self.FP.popScope()
    return Nodes_ApplyNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), varSymList, expListNode, block)
}

// 4291: decl @lune.@base.@TransUnit.TransUnit.convToExtTypeList
func (self *TransUnit_TransUnit) convToExtTypeList(pos *Types_Position,typeInfo *Ast_TypeInfo,list *LnsList) *LnsList {
    if typeInfo.FP.Get_nonnilableType().FP.Get_kind() != Ast_TypeInfoKind__Ext{
        return list
    }
    var newList LnsAny
    var mess string
    newList,mess = Ast_convToExtTypeList(self.processInfo, list)
    if newList != nil{
        newList_8017 := newList.(*LnsList)
        return newList_8017
    }
    self.FP.addErrMess(pos, mess)
    return list
}

// 4307: decl @lune.@base.@TransUnit.TransUnit.analyzeForeach
func (self *TransUnit_TransUnit) analyzeForeach(token *Types_Token,sortFlag bool) *Nodes_Node {
    var scope *Ast_Scope
    scope = self.FP.pushScope(false, nil, nil)
    var mainSymToken *Types_Token
    mainSymToken = Parser_getEofToken()
    var subSymToken LnsAny
    subSymToken = nil
    var mainSym *Ast_SymbolInfo
    var subSym LnsAny
    subSym = nil
    var nextToken *Types_Token
    nextToken = Parser_getEofToken()
    {
        var _from31231 LnsInt = 1
        var _to31231 LnsInt = 2
        for _work31231 := _from31231; _work31231 <= _to31231; _work31231++ {
            index := _work31231
            var symbol *Types_Token
            symbol = self.FP.getToken(nil)
            if symbol.Kind != Types_TokenKind__Symb{
                self.FP.Error("illegal symbol")
            }
            if index == 1{
                mainSymToken = symbol
                
            } else { 
                subSymToken = symbol
                
            }
            nextToken = self.FP.getToken(nil)
            
            if nextToken.Txt != ","{
                break
            }
        }
    }
    self.FP.checkToken(nextToken, "in")
    var exp *Nodes_Node
    exp = self.FP.analyzeExpOneRVal(false, false, nil, nil)
    var expType *Ast_TypeInfo
    expType = exp.FP.Get_expType().FP.Get_extedType()
    var itemTypeInfoList *LnsList
    itemTypeInfoList = self.FP.convToExtTypeList(token.Pos, exp.FP.Get_expType(), expType.FP.Get_itemTypeInfoList())
    if _switch31505 := expType.FP.Get_kind(); _switch31505 == Ast_TypeInfoKind__Map {
        mainSym = self.FP.addLocalVar(mainSymToken.Pos, false, true, mainSymToken.Txt, itemTypeInfoList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil)
        
        {
            __exp := subSymToken
            if __exp != nil {
                _exp := __exp.(*Types_Token)
                subSym = self.FP.addLocalVar(_exp.Pos, false, true, _exp.Txt, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil)
                
            }
        }
    } else if _switch31505 == Ast_TypeInfoKind__Set {
        if subSymToken != nil{
            subSymToken_8045 := subSymToken.(*Types_Token)
            self.FP.addErrMess(subSymToken_8045.Pos, "Set can't use index")
        }
        mainSym = self.FP.addLocalVar(mainSymToken.Pos, false, true, mainSymToken.Txt, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil)
        
    } else if _switch31505 == Ast_TypeInfoKind__List || _switch31505 == Ast_TypeInfoKind__Array {
        if sortFlag{
            self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("'%s' doesn't support forsort.", []LnsAny{Ast_TypeInfoKind_getTxt( expType.FP.Get_kind())}))
        }
        mainSym = self.FP.addLocalVar(mainSymToken.Pos, false, true, mainSymToken.Txt, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil)
        
        {
            __exp := subSymToken
            if __exp != nil {
                _exp := __exp.(*Types_Token)
                subSym = self.FP.addLocalVar(_exp.Pos, false, false, _exp.Txt, Ast_builtinTypeInt, Ast_MutMode__IMut, nil)
                
            }
        }
    } else {
        self.FP.ErrorAt(exp.FP.Get_pos(), Lns_getVM().String_format("unknown kind type of exp for foreach -- %s", []LnsAny{expType.FP.GetTxt(nil, nil, nil)}))
    }
    var seqSym LnsAny
    seqSym = nil
    {
        _refNode := Nodes_ExpRefNodeDownCastF(exp.FP)
        if _refNode != nil {
            refNode := _refNode.(*Nodes_ExpRefNode)
            var seqSymbol *Ast_SymbolInfo
            seqSymbol = refNode.FP.Get_symbolInfo()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( seqSymbol.FP.Get_mutable()) ||
                Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(seqSymbol.FP.Get_typeInfo())) ).(bool){
                scope.FP.AddOverrideImut(seqSymbol)
                seqSym = seqSymbol.FP.Get_name()
                
            }
        }
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(Nodes_BlockKind__Foreach, TransUnit_TentativeMode__Loop, scope, nil)
    if seqSym != nil{
        seqSym_8058 := seqSym.(string)
        scope.FP.Remove(seqSym_8058)
    }
    self.FP.popScope()
    var threading bool
    if exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Ext{
        threading = self.FP.checkThreading(token.Pos)
        
    } else { 
        threading = false
        
    }
    if sortFlag{
        return &Nodes_ForsortNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), mainSym, subSym, exp, threading, block, sortFlag).Nodes_Node
    } else { 
        return &Nodes_ForeachNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), mainSym, subSym, exp, threading, block).Nodes_Node
    }
// insert a dummy
    return nil
}

// 4428: decl @lune.@base.@TransUnit.TransUnit.analyzeProvide
func (self *TransUnit_TransUnit) analyzeProvide(firstToken *Types_Token) *Nodes_ProvideNode {
    var token *Types_Token
    token = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
    var symbolNode *Nodes_Node
    symbolNode = self.FP.analyzeExpSymbol(firstToken, token, TransUnit_ExpSymbolMode__Symbol, nil, true, false)
    self.FP.checkNextToken(";")
    var symbolInfoList *LnsList
    symbolInfoList = symbolNode.FP.GetSymbolInfo()
    if symbolInfoList.Len() != 1{
        self.FP.Error("'provide' must be symbol.")
    }
    var symbolInfo *Ast_SymbolInfo
    symbolInfo = symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
    var node *Nodes_ProvideNode
    node = Nodes_ProvideNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), symbolInfo)
    if Lns_isCondTrue( self.provideNode){
        self.FP.addErrMess(firstToken.Pos, "multiple provide")
    }
    self.provideNode = node
    
    if symbolInfo.FP.Get_accessMode() != Ast_AccessMode__Pub{
        self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("provide variable must be 'pub'.  -- %s", []LnsAny{symbolInfo.FP.Get_accessMode()}))
    }
    return node
}

// 4459: decl @lune.@base.@TransUnit.TransUnit.analyzeScope
func (self *TransUnit_TransUnit) analyzeScope(firstToken *Types_Token) *Nodes_ScopeNode {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var scopeKind LnsInt
    if _switch31927 := nextToken.Txt; _switch31927 == "root" {
        scopeKind = Nodes_ScopeKind__Root
        
    } else {
        self.FP.Error(Lns_getVM().String_format("illegal scope kind. -- %s", []LnsAny{nextToken.Txt}))
    }
    var symList *LnsList
    symList = NewLnsList([]LnsAny{})
    nextToken = self.FP.getToken(nil)
    
    if nextToken.Txt == "("{
        nextToken = self.FP.getToken(nil)
        
        for nextToken.Txt != ")" {
            var symbolNode *Nodes_Node
            symbolNode = self.FP.analyzeExpSymbol(nextToken, nextToken, TransUnit_ExpSymbolMode__Symbol, nil, true, false)
            var workSymList *LnsList
            workSymList = symbolNode.FP.GetSymbolInfo()
            if workSymList.Len() > 0{
                symList.Insert(Ast_SymbolInfo2Stem(workSymList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()))
            }
            nextToken = self.FP.getToken(nil)
            
            if _switch32074 := nextToken.Txt; _switch32074 == ")" {
            } else if _switch32074 == "," {
                nextToken = self.FP.getToken(nil)
                
            } else {
                self.FP.Error(Lns_getVM().String_format("illegal token: expects ')' or ',' but -- %s", []LnsAny{nextToken.Txt}))
            }
        }
    } else { 
        self.FP.Pushback()
    }
    var bakScope *Ast_Scope
    bakScope = self.scope
    self.scope = self.topScope
    
    self.FP.pushScope(false, nil, nil)
    for _, _symInfo := range( symList.Items ) {
        symInfo := _symInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        self.scope.FP.AddAlias(self.processInfo, symInfo.FP.Get_name(), nextToken.Pos, false, symInfo.FP.Get_accessMode(), symInfo.FP.Get_typeInfo().FP.Get_parentInfo(), symInfo)
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(Nodes_BlockKind__Block, TransUnit_TentativeMode__Simple, nil, nil)
    var node *Nodes_ScopeNode
    node = Nodes_ScopeNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), scopeKind, self.scope, symList, block)
    self.scope = bakScope
    
    return node
}

// 34: decl @lune.@base.@TransUnit.TransUnit.analyzeRefType
func (self *TransUnit_TransUnit) analyzeRefType(accessMode LnsInt,allowDDD bool,parentPub bool) *Nodes_RefTypeNode {
    var firstToken *Types_Token
    firstToken = self.FP.getToken(nil)
    var token *Types_Token
    token = firstToken
    var refFlag bool
    refFlag = false
    if token.Txt == "&"{
        refFlag = true
        
        token = self.FP.getToken(nil)
        
    }
    var mutFlag bool
    mutFlag = false
    if token.Txt == "mut"{
        mutFlag = true
        
        token = self.FP.getToken(nil)
        
    }
    var name *Nodes_Node
    if token.Txt == "..."{
        var dddSym *Ast_SymbolInfo
        dddSym = Lns_unwrap( self.moduleScope.FP.GetSymbolInfo("...", self.moduleScope, true, Ast_ScopeAccess__Normal)).(*Ast_SymbolInfo)
        name = &Nodes_ExpRefNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddSym.FP.Get_typeInfo())}), &NewAst_AccessSymbolInfo(dddSym, Ast_OverrideMut__None_Obj, true).Ast_SymbolInfo).Nodes_Node
        
    } else { 
        name = self.FP.analyzeExpSymbol(firstToken, token, TransUnit_ExpSymbolMode__Symbol, nil, true, false)
        
        var symbolList *LnsList
        symbolList = name.FP.GetSymbolInfo()
        if symbolList.Len() > 0{
            var symbol *Ast_SymbolInfo
            symbol = symbolList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if symbol.FP.Get_kind() != Ast_SymbolKind__Typ{
                self.FP.addErrMess(name.FP.Get_pos(), Lns_getVM().String_format("illegal type -- %s", []LnsAny{symbol.FP.Get_name()}))
            }
        } else { 
            self.FP.addErrMess(name.FP.Get_pos(), Lns_getVM().String_format("illegal symbol node -- %s", []LnsAny{Nodes_getNodeKindName(name.FP.Get_kind())}))
        }
    }
    return self.FP.analyzeRefTypeWithSymbol(accessMode, allowDDD, refFlag, mutFlag, name, parentPub)
}

// 84: decl @lune.@base.@TransUnit.TransUnit.analyzeRefTypeWithSymbol
func (self *TransUnit_TransUnit) analyzeRefTypeWithSymbol(accessMode LnsInt,allowDDD bool,refFlag bool,mutFlag bool,symbolNode *Nodes_Node,parentPub bool) *Nodes_RefTypeNode {
    var typeInfo *Ast_TypeInfo
    typeInfo = symbolNode.FP.Get_expType()
    {
        _aliasType := Ast_AliasTypeInfoDownCastF(typeInfo.FP)
        if _aliasType != nil {
            aliasType := _aliasType.(*Ast_AliasTypeInfo)
            var aliasSrc *Ast_TypeInfo
            aliasSrc = aliasType.FP.Get_aliasSrcTypeInfo()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(self.importModuleSet.Has(Ast_TypeInfo2Stem(aliasSrc.FP.GetModule())))) &&
                Lns_GetEnv().SetStackVal( self.moduleType.FP.Get_parentInfo() != aliasSrc.FP.GetModule().FP.Get_parentInfo()) ).(bool)){
                self.FP.addErrMess(symbolNode.FP.Get_pos(), Lns_getVM().String_format("must import '%s' for this alias -- %s", []LnsAny{aliasType.FP.GetModule().FP.GetFullName(self.typeNameCtrl, self.scope.FP, false), symbolNode.FP.GetSymbolInfo().GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_name()}))
            }
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( parentPub) &&
        Lns_GetEnv().SetStackVal( Ast_isPubToExternal(accessMode)) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(typeInfo.FP.Get_accessMode()))) ).(bool)){
        self.FP.addErrMess(symbolNode.FP.Get_pos(), Lns_getVM().String_format("This type must be public. -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}))
    }
    var continueToken *Types_Token
    var continueFlag bool
    continueToken,continueFlag = self.FP.getContinueToken()
    var token *Types_Token
    token = continueToken
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( continueFlag) &&
        Lns_GetEnv().SetStackVal( token.Txt == "!") ).(bool)){
        typeInfo = typeInfo.FP.Get_nilableTypeInfo()
        
        token = self.FP.getToken(nil)
        
    }
    var itemNodeList *LnsList
    itemNodeList = NewLnsList([]LnsAny{})
    var arrayMode string
    arrayMode = "no"
    for  {
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == "[") ||
            Lns_GetEnv().SetStackVal( token.Txt == "[@") ).(bool){
            if token.Txt == "["{
                arrayMode = "list"
                
                typeInfo = self.processInfo.FP.CreateList(accessMode, self.FP.getCurrentClass(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), Ast_MutMode__Mut)
                
            } else { 
                arrayMode = "array"
                
                typeInfo = self.processInfo.FP.CreateArray(accessMode, self.FP.getCurrentClass(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), Ast_MutMode__Mut)
                
            }
            token = self.FP.getToken(nil)
            
            if token.Txt != "]"{
                self.FP.Pushback()
                self.FP.checkNextToken("]")
            }
            itemNodeList.Insert(Nodes_Node2Stem(symbolNode))
        } else if token.Txt == "<"{
            var genericList *LnsList
            genericList = NewLnsList([]LnsAny{})
            var nextToken *Types_Token
            nextToken = Parser_getEofToken()
            for {
                var typeExp *Nodes_RefTypeNode
                typeExp = self.FP.analyzeRefType(accessMode, false, parentPub)
                itemNodeList.Insert(Nodes_RefTypeNode2Stem(typeExp))
                genericList.Insert(Ast_TypeInfo2Stem(typeExp.FP.Get_expType()))
                nextToken = self.FP.getToken(nil)
                
                if nextToken.Txt != ","{ break }
            }
            self.FP.checkToken(nextToken, ">")
            var checkAlternateTypeCount func(count LnsInt) bool
            checkAlternateTypeCount = func(count LnsInt) bool {
                if genericList.Len() != count{
                    self.FP.addErrMess(symbolNode.FP.Get_pos(), Lns_getVM().String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()}))
                    return false
                }
                return true
            }
            if _switch33407 := typeInfo.FP.Get_kind(); _switch33407 == Ast_TypeInfoKind__Map {
                if genericList.Len() != 2{
                    self.FP.addErrMess(symbolNode.FP.Get_pos(), "Key or value type is unknown")
                    typeInfo = self.processInfo.FP.CreateMap(accessMode, self.FP.getCurrentClass(), Ast_builtinTypeStem, Ast_builtinTypeStem, Ast_MutMode__Mut)
                    
                } else { 
                    typeInfo = self.processInfo.FP.CreateMap(accessMode, self.FP.getCurrentClass(), genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), genericList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__Mut)
                    
                }
            } else if _switch33407 == Ast_TypeInfoKind__List {
                if checkAlternateTypeCount(1){
                    typeInfo = self.processInfo.FP.CreateList(accessMode, self.FP.getCurrentClass(), genericList, Ast_MutMode__Mut)
                    
                }
            } else if _switch33407 == Ast_TypeInfoKind__Array {
                if checkAlternateTypeCount(1){
                    typeInfo = self.processInfo.FP.CreateArray(accessMode, self.FP.getCurrentClass(), genericList, Ast_MutMode__Mut)
                    
                }
            } else if _switch33407 == Ast_TypeInfoKind__Set {
                if checkAlternateTypeCount(1){
                    typeInfo = self.processInfo.FP.CreateSet(accessMode, self.FP.getCurrentClass(), genericList, Ast_MutMode__Mut)
                    
                }
            } else if _switch33407 == Ast_TypeInfoKind__DDD {
                if checkAlternateTypeCount(1){
                    typeInfo = &self.processInfo.FP.CreateDDD(genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), false, false).Ast_TypeInfo
                    
                }
            } else if _switch33407 == Ast_TypeInfoKind__Class || _switch33407 == Ast_TypeInfoKind__IF {
                if checkAlternateTypeCount(typeInfo.FP.Get_itemTypeInfoList().Len()){
                    for _, _itemType := range( genericList.Items ) {
                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if itemType.FP.Get_nilable(){
                            self.FP.addErrMess(symbolNode.FP.Get_pos(), Lns_getVM().String_format("can't use nilable type -- %s", []LnsAny{itemType.FP.GetTxt(nil, nil, nil)}))
                        }
                    }
                    for _index, _itemType := range( typeInfo.FP.Get_itemTypeInfoList().Items ) {
                        index := _index + 1
                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( itemType.FP.HasBase()) &&
                            Lns_GetEnv().SetStackVal( Lns_op_not(genericList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.IsInheritFrom(self.processInfo, itemType.FP.Get_baseTypeInfo(), nil))) ).(bool)){
                            self.FP.addErrMess(symbolNode.FP.Get_pos(), Lns_getVM().String_format("'%s' of %s doesn't inherit '%s'", []LnsAny{itemType.FP.GetTxt(nil, nil, nil), typeInfo.FP.GetTxt(nil, nil, nil), itemType.FP.Get_baseTypeInfo().FP.GetTxt(nil, nil, nil)}))
                        }
                    }
                    typeInfo = &self.processInfo.FP.CreateGeneric(typeInfo, genericList, self.moduleType).Ast_TypeInfo
                    
                }
            } else if _switch33407 == Ast_TypeInfoKind__Box {
                if checkAlternateTypeCount(1){
                    typeInfo = self.processInfo.FP.CreateBox(accessMode, genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                    
                }
            } else if _switch33407 == Ast_TypeInfoKind__Ext {
                if checkAlternateTypeCount(1){
                    typeInfo = self.FP.createExtType(symbolNode.FP.Get_pos(), genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                    
                }
            } else {
                self.FP.Error(Lns_getVM().String_format("not support generic: %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}))
            }
        } else { 
            self.FP.Pushback()
            break
        }
        token = self.FP.getToken(nil)
        
    }
    if token.Txt == "!"{
        typeInfo = typeInfo.FP.Get_nilableTypeInfo()
        
        self.FP.getToken(nil)
    }
    if Lns_op_not(allowDDD){
        if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__DDD{
            self.FP.addErrMess(symbolNode.FP.Get_pos(), Lns_getVM().String_format("invalid type. -- '%s'", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}))
        }
    }
    if refFlag{
        typeInfo = self.FP.CreateModifier(typeInfo, Ast_MutMode__IMut)
        
    }
    if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Module{
        self.FP.addErrMess(symbolNode.FP.Get_pos(), Lns_getVM().String_format("module can't use as Type. -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}))
    }
    return Nodes_RefTypeNode_create(self.nodeManager, symbolNode.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), symbolNode, itemNodeList, refFlag, mutFlag, arrayMode)
}

// 282: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclArgList
func (self *TransUnit_TransUnit) analyzeDeclArgList(accessMode LnsInt,scope *Ast_Scope,argList *LnsList,parentPub bool) *Types_Token {
    var nextToken *Types_Token
    nextToken = Parser_noneToken
    var hasDDDFlag bool
    hasDDDFlag = false
    for {
        nextToken = self.FP.getToken(nil)
        
        if nextToken.Txt == ")"{
            break
        }
        if hasDDDFlag{
            self.FP.addErrMess(nextToken.Pos, "Argument exists after '...'.")
        }
        var mutable LnsInt
        mutable = Ast_MutMode__IMut
        if nextToken.Txt == "mut"{
            mutable = Ast_MutMode__Mut
            
            nextToken = self.FP.getToken(nil)
            
        }
        var argName *Types_Token
        argName = nextToken
        if argName.Txt == "..."{
            hasDDDFlag = true
            
            var workToken *Types_Token
            var flag bool
            workToken,flag = self.FP.getContinueToken()
            self.FP.Pushback()
            var dddTypeInfo *Ast_TypeInfo
            dddTypeInfo = Ast_builtinTypeDDD
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( flag) &&
                Lns_GetEnv().SetStackVal( workToken.Txt == "<") ).(bool)){
                self.FP.PushbackToken(nextToken)
                var refTypeNode *Nodes_RefTypeNode
                refTypeNode = self.FP.analyzeRefType(accessMode, true, parentPub)
                dddTypeInfo = refTypeNode.FP.Get_expType()
                
            }
            argList.Insert(Nodes_DeclArgDDDNode2Stem(Nodes_DeclArgDDDNode_create(self.nodeManager, argName.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddTypeInfo)}))))
            scope.FP.AddLocalVar(self.processInfo, true, true, argName.Txt, argName.Pos, dddTypeInfo, Ast_MutMode__IMut)
        } else { 
            argName = self.FP.checkSymbol(argName, TransUnit_SymbolMode__MustNot_)
            
            self.FP.checkShadowing(argName.Pos, argName.Txt, scope)
            self.FP.checkNextToken(":")
            var refType *Nodes_RefTypeNode
            refType = self.FP.analyzeRefType(accessMode, false, parentPub)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( refType.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Class) &&
                Lns_GetEnv().SetStackVal( refType.FP.Get_expType().FP.Get_itemTypeInfoList().Len() > 0) ).(bool)){
                var argType *Ast_TypeInfo
                argType = refType.FP.Get_expType().FP.Get_srcTypeInfo()
                if Lns_op_not(Ast_isGenericType(argType)){
                    self.FP.addErrMess(refType.FP.Get_pos(), Lns_getVM().String_format("can't use this type without <T>. please use %s.", []LnsAny{argType.FP.GetTxt(nil, nil, nil)}))
                }
            }
            {
                _symbolInfo := TransUnit_convExp33994(Lns_2DDD(scope.FP.AddLocalVar(self.processInfo, true, true, argName.Txt, argName.Pos, refType.FP.Get_expType(), mutable)))
                if _symbolInfo != nil {
                    symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                    var arg *Nodes_DeclArgNode
                    arg = Nodes_DeclArgNode_create(self.nodeManager, argName.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), refType.FP.Get_expTypeList(), argName, symbolInfo, refType)
                    argList.Insert(Nodes_DeclArgNode2Stem(arg))
                }
            }
        }
        nextToken = self.FP.getToken(nil)
        
        if nextToken.Txt != ","{ break }
    }
    self.FP.checkToken(nextToken, ")")
    return nextToken
}

// 368: decl @lune.@base.@TransUnit.TransUnit.checkOverrideMethod
func (self *TransUnit_TransUnit) checkOverrideMethod(overrideType *Ast_TypeInfo,typeInfo *Ast_TypeInfo) *LnsList {
    var accessMode LnsInt
    accessMode = typeInfo.FP.Get_accessMode()
    var funcName string
    funcName = typeInfo.FP.Get_rawTxt()
    var altTypeList *LnsList
    altTypeList = typeInfo.FP.Get_itemTypeInfoList()
    var alt2typeMap *LnsMap
    alt2typeMap = typeInfo.FP.Get_parentInfo().FP.CreateAlt2typeMap(false)
    var errList *LnsList
    errList = NewLnsList([]LnsAny{})
    var addErr func(mess string)
    addErr = func(mess string) {
        var fullName string
        fullName = Lns_getVM().String_format("%s.%s", []LnsAny{typeInfo.FP.Get_parentInfo().FP.Get_rawTxt(), typeInfo.FP.Get_rawTxt()})
        errList.Insert(Lns_getVM().String_format("%s: %s: %s -- %s", []LnsAny{fullName, mess, typeInfo.FP.Get_display_stirng(), typeInfo.FP.Get_display_stirng()}))
    }
    if overrideType.FP.Get_accessMode() != accessMode{
        var mess string
        mess = Lns_getVM().String_format("mismatch override accessMode -- %s,%s,%s", []LnsAny{funcName, Ast_AccessMode_getTxt( overrideType.FP.Get_accessMode()), Ast_AccessMode_getTxt( accessMode)})
        addErr(mess)
    }
    if overrideType.FP.Get_staticFlag() != typeInfo.FP.Get_staticFlag(){
        addErr("mismatch override staticFlag -- " + funcName)
    }
    if overrideType.FP.Get_kind() != Ast_TypeInfoKind__Method{
        addErr(Lns_getVM().String_format("mismatch override kind -- %s, %d", []LnsAny{funcName, overrideType.FP.Get_kind()}))
    }
    if overrideType.FP.Get_mutMode() != typeInfo.FP.Get_mutMode(){
        addErr(Lns_getVM().String_format("mismatch mutable -- %s", []LnsAny{funcName}))
    }
    if overrideType.FP.Get_itemTypeInfoList().Len() != altTypeList.Len(){
        var mess string
        mess = Lns_getVM().String_format("mismatch altTypeList -- %d, %d", []LnsAny{overrideType.FP.Get_itemTypeInfoList().Len(), altTypeList.Len()})
        addErr(mess)
    } else { 
        for _index, _alterType := range( overrideType.FP.Get_itemTypeInfoList().Items ) {
            index := _index + 1
            alterType := _alterType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            alt2typeMap.Set(alterType,altTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        }
    }
    var matchFlag bool
    var err LnsAny
    matchFlag,err = overrideType.FP.CanEvalWith(self.processInfo, typeInfo, Ast_CanEvalType__SetEq, alt2typeMap)
    if Lns_op_not(matchFlag){
        if err != nil{
            err_8252 := err.(string)
            addErr(Lns_getVM().String_format("mismatch method type -- %s", []LnsAny{err_8252}))
        } else {
            addErr("mismatch method type")
        }
    }
    for _index, _retType := range( overrideType.FP.Get_retTypeInfoList().Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if typeInfo.FP.Get_retTypeInfoList().Len() >= index{
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( retType.FP.Get_nonnilableType().FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
                Lns_GetEnv().SetStackVal( typeInfo.FP.Get_retTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nonnilableType().FP.Get_kind() != Ast_TypeInfoKind__Alternate) ).(bool)){
                var mess string
                mess = Lns_getVM().String_format("not support to override the method has generics at return type. -- %s", []LnsAny{funcName})
                addErr(mess)
            }
        }
    }
    return errList
}

// 439: decl @lune.@base.@TransUnit.TransUnit.checkOverriededMethodOfAllClass
func (self *TransUnit_TransUnit) checkOverriededMethodOfAllClass() {
    var process func(pos *Types_Position,alt2typeMap *LnsMap,classScope *Ast_Scope,superScope *Ast_Scope)
    process = func(pos *Types_Position,alt2typeMap *LnsMap,classScope *Ast_Scope,superScope *Ast_Scope) {
        superScope.FP.FilterTypeInfoField(true, classScope, self.scopeAccess, Ast_filterForm(func(symbolInfo *Ast_SymbolInfo) bool {
            if symbolInfo.FP.Get_name() == "__init"{
                return true
            }
            var implimented bool
            implimented = true
            if symbolInfo.FP.Get_typeInfo().FP.Get_kind() == Ast_TypeInfoKind__Method{
                {
                    _impMethodType := classScope.FP.GetTypeInfoField(symbolInfo.FP.Get_name(), true, classScope, self.scopeAccess)
                    if _impMethodType != nil {
                        impMethodType := _impMethodType.(*Ast_TypeInfo)
                        if impMethodType == symbolInfo.FP.Get_typeInfo(){
                            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_typeInfo().FP.Get_abstractFlag()) &&
                                Lns_GetEnv().SetStackVal( classScope != superScope) ).(bool)){
                                implimented = false
                                
                            }
                        } else { 
                            for _, _err := range( self.FP.checkOverrideMethod(symbolInfo.FP.Get_typeInfo(), impMethodType).Items ) {
                                err := _err.(string)
                                self.FP.addErrMess(pos, err)
                            }
                        }
                    } else {
                        implimented = false
                        
                    }
                }
            }
            if Lns_op_not(implimented){
                self.FP.addErrMess(pos, Lns_getVM().String_format("not implements method -- %s.%s at %s", []LnsAny{Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(superScope.FP.Get_ownerTypeInfo()) && 
                Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.GetTxt(nil, nil, nil)})/* 473:35 */), symbolInfo.FP.Get_name(), Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(classScope.FP.Get_ownerTypeInfo()) && 
                Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.GetTxt(nil, nil, nil)})/* 475:35 */)}))
            }
            return true
        }))
    }
    var typeId2DeclClassNode *LnsMap
    typeId2DeclClassNode = NewLnsMap( map[LnsAny]LnsAny{})
    for _classTypeInfo, _classNode := range( self.typeInfo2ClassNode.Items ) {
        classTypeInfo := _classTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        classNode := _classNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        typeId2DeclClassNode.Set(classTypeInfo.FP.Get_typeId(),classNode)
    }
    {
        __collection34735 := typeId2DeclClassNode
        __sorted34735 := __collection34735.CreateKeyListInt()
        __sorted34735.Sort( LnsItemKindInt, nil )
        for _, ___key34735 := range( __sorted34735.Items ) {
            classNode := __collection34735.Items[ ___key34735 ].(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            var classTypeInfo *Ast_TypeInfo
            classTypeInfo = classNode.FP.Get_expType()
            var workTypeInfo *Ast_TypeInfo
            workTypeInfo = classTypeInfo
            var alt2typeMap *LnsMap
            alt2typeMap = classTypeInfo.FP.CreateAlt2typeMap(false)
            for {
                if Lns_op_not(classTypeInfo.FP.Get_abstractFlag()){
                    if workTypeInfo != Ast_headTypeInfo{
                        process(classNode.FP.Get_pos(), alt2typeMap, Lns_unwrap( classTypeInfo.FP.Get_scope()).(*Ast_Scope), Lns_unwrap( workTypeInfo.FP.Get_scope()).(*Ast_Scope))
                    }
                }
                for _, _ifType := range( workTypeInfo.FP.Get_interfaceList().Items ) {
                    ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if ifType != Ast_builtinTypeMapping{
                        process(classNode.FP.Get_pos(), alt2typeMap, Lns_unwrap( classTypeInfo.FP.Get_scope()).(*Ast_Scope), Lns_unwrap( ifType.FP.Get_scope()).(*Ast_Scope))
                    }
                }
                workTypeInfo = workTypeInfo.FP.Get_baseTypeInfo()
                
                if workTypeInfo == Ast_headTypeInfo{ break }
            }
        }
    }
}

// 517: decl @lune.@base.@TransUnit.TransUnit.createAST
func (self *TransUnit_TransUnit) CreateAST(parser *Parser_Parser,macroFlag bool,moduleName LnsAny) *TransUnit_ASTInfo {
    self.moduleName = Lns_unwrapDefault( moduleName, "").(string)
    
    var backProcessInfo *Ast_ProcessInfo
    backProcessInfo = self.processInfo
    self.processInfo = Ast_getRootProcessInfo()
    
    self.FP.registBuiltInScope()
    self.processInfo = backProcessInfo
    
    var moduleTypeInfo *Ast_TypeInfo
    moduleTypeInfo = Ast_headTypeInfo
    var moduleSymbolKind LnsInt
    moduleSymbolKind = Ast_SymbolKind__Typ
    if moduleName != nil{
        moduleName_8322 := moduleName.(string)
        {
            _form34861, _param34861, _prev34861 := Lns_getVM().String_gmatch(moduleName_8322, "[^%.]+")
            for {
                _work34861 := _form34861.(*Lns_luaValue).Call( Lns_2DDD( _param34861, _prev34861 ) )
                _prev34861 = Lns_getFromMulti(_work34861,0)
                if Lns_IsNil( _prev34861 ) { break }
                txt := _prev34861.(string)
                moduleTypeInfo = self.FP.pushModule(false, txt, true)
                
            }
        }
    }
    self.moduleScope = self.scope
    
    self.moduleType = moduleTypeInfo
    
    self.moduleScope.FP.AddVar(self.processInfo, Ast_AccessMode__Global, "__mod__", nil, Ast_builtinTypeString, Ast_MutMode__IMut, true)
    self.typeNameCtrl = NewAst_TypeNameCtrl(moduleTypeInfo)
    
    self.parser = NewParser_DefaultPushbackParser(parser)
    
    self.scope.FP.AddIgnoredVar(self.processInfo)
    var ast *Nodes_Node
    var lastStatement LnsAny
    lastStatement = nil
    if macroFlag{
        ast = &self.FP.analyzeBlock(Nodes_BlockKind__Macro, TransUnit_TentativeMode__Ignore, nil, nil).Nodes_Node
        
    } else { 
        var children *LnsList
        children = NewLnsList([]LnsAny{})
        var lastLineNo LnsInt
        lastStatement, lastLineNo = self.FP.analyzeStatementList(children, nil)
        
        var statement *Nodes_BlankLineNode
        statement = Nodes_BlankLineNode_create(self.nodeManager, self.FP.CreatePosition(lastLineNo + 1, 0), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), 0)
        statement.FP.AddComment(self.commentCtrl.FP.Get_commentList())
        self.commentCtrl.FP.Clear()
        children.Insert(Nodes_BlankLineNode2Stem(statement))
        var token *Types_Token
        token = self.FP.GetTokenNoErr()
        if token != Parser_getEofToken(){
            self.FP.Error(Lns_getVM().String_format("%s:%d:%d:(%s) not eof -- %s", []LnsAny{self.parser.FP.GetStreamName(), token.Pos.LineNo, token.Pos.Column, Types_TokenKind_getTxt( token.Kind), token.Txt}))
        }
        for _, _subModule := range( self.subfileList.Items ) {
            subModule := _subModule.(string)
            var file string
            
            {
                _file := FrontInterface_searchModule(subModule)
                if _file == nil{
                    self.FP.Error(Lns_getVM().String_format("not found subfile -- %s", []LnsAny{subModule}))
                } else {
                    file = _file.(string)
                }
            }
            if self.scope != self.moduleScope{
                self.FP.Error("scope does not close")
            }
            var subParser *Parser_StreamParser
            
            {
                _subParser := Parser_StreamParser_create(file, false, subModule)
                if _subParser == nil{
                    self.FP.Error(Lns_getVM().String_format("open error -- %s", []LnsAny{file}))
                } else {
                    subParser = _subParser.(*Parser_StreamParser)
                }
            }
            self.parser = NewParser_DefaultPushbackParser(&subParser.Parser_Parser)
            
            lastStatement = self.FP.analyzeStatementListSubfile(children)
            
            token = self.FP.GetTokenNoErr()
            
            if token != Parser_getEofToken(){
                Util_err(Lns_getVM().String_format("unknown:%d:%d:(%s) %s", []LnsAny{token.Pos.LineNo, token.Pos.Column, Types_TokenKind_getTxt( token.Kind), token.Txt}))
            }
        }
        self.FP.checkOverriededMethodOfAllClass()
        var rootNode *Nodes_RootNode
        rootNode = Nodes_RootNode_create(self.nodeManager, self.FP.CreatePosition(0, 0), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), children, self.moduleScope, self.macroCtrl.FP.Get_useModuleMacroSet(), self.moduleId, self.processInfo, moduleTypeInfo, nil, self.helperInfo, self.nodeManager, self.importModule2ModuleInfo, self.macroCtrl.FP.Get_typeId2MacroInfo(), self.typeId2ClassMap)
        ast = &rootNode.Nodes_Node
        
        {
            __exp := self.provideNode
            if __exp != nil {
                _exp := __exp.(*Nodes_ProvideNode)
                if lastStatement != &_exp.Nodes_Node{
                    self.FP.addErrMess(_exp.FP.Get_pos(), "'provide' must be last.")
                }
                rootNode.FP.Set_provide(_exp)
                moduleSymbolKind = _exp.FP.Get_symbol().FP.Get_kind()
                
            }
        }
        TransUnit_ClosureFun_checkList_1280_(self.closureFunList)
    }
    if moduleName != nil{
        moduleName_8349 := moduleName.(string)
        {
            _form35433, _param35433, _prev35433 := Lns_getVM().String_gmatch(moduleName_8349, "[^%.]+")
            for {
                _work35433 := _form35433.(*Lns_luaValue).Call( Lns_2DDD( _param35433, _prev35433 ) )
                _prev35433 = Lns_getFromMulti(_work35433,0)
                if Lns_IsNil( _prev35433 ) { break }
                self.FP.popModule()
            }
        }
    }
    {
        __collection35525 := TransUnit_createAST__createId2proto_5304_(self.protoFuncMap)
        __sorted35525 := __collection35525.CreateKeyListInt()
        __sorted35525.Sort( LnsItemKindInt, nil )
        for _, ___key35525 := range( __sorted35525.Items ) {
            protoType := __collection35525.Items[ ___key35525 ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            self.FP.addErrMess(Lns_unwrap( self.protoFuncMap.Items[protoType]).(*Types_Position), Lns_getVM().String_format("This function doesn't have body. -- %s", []LnsAny{protoType.FP.GetTxt(nil, nil, nil)}))
        }
    }
    {
        __collection35563 := TransUnit_createAST__createId2proto_5304_(self.protoClassMap)
        __sorted35563 := __collection35563.CreateKeyListInt()
        __sorted35563.Sort( LnsItemKindInt, nil )
        for _, ___key35563 := range( __sorted35563.Items ) {
            protoType := __collection35563.Items[ ___key35563 ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            self.FP.addErrMess(Lns_unwrap( self.protoClassMap.Items[protoType]).(*Types_Position), Lns_getVM().String_format("This class doesn't have body. -- %s", []LnsAny{protoType.FP.GetTxt(nil, nil, nil)}))
        }
    }
    for _, _mess := range( self.warnMessList.Items ) {
        mess := _mess.(string)
        Util_errorLog(mess)
    }
    if self.errMessList.Len() > 0{
        for _, _mess := range( self.errMessList.Items ) {
            mess := _mess.(string)
            Util_errorLog(mess)
        }
        Util_err("has error")
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.ctrl_info.StopByWarning) &&
        Lns_GetEnv().SetStackVal( self.warnMessList.Len() > 0) ).(bool)){
        Util_err("has error")
    }
    if _switch35651 := self.analyzeMode; _switch35651 == TransUnit_AnalyzeMode__Diag || _switch35651 == TransUnit_AnalyzeMode__Complete || _switch35651 == TransUnit_AnalyzeMode__Inquire {
        Lns_getVM().OS_exit(0)
    }
    return NewTransUnit_ASTInfo(ast, moduleTypeInfo, moduleSymbolKind, self.processInfo)
}

// 674: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclMacroSub
func (self *TransUnit_TransUnit) analyzeDeclMacroSub(accessMode LnsInt,firstToken *Types_Token,nameToken *Types_Token,macroScope *Ast_Scope,parentType *Ast_TypeInfo,workArgList *LnsList) *Nodes_DeclMacroNode {
    if self.macroCtrl.FP.Get_isDeclaringMacro(){
        self.FP.Error("can't declare macro in the macro.")
    }
    self.macroCtrl.FP.StartDecl()
    var pubFlag bool
    pubFlag = false
    if _switch35767 := accessMode; _switch35767 == Ast_AccessMode__Pub {
        pubFlag = true
        
    } else if _switch35767 == Ast_AccessMode__Local || _switch35767 == Ast_AccessMode__None {
    } else {
        self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("macro not support this access mode. -- %s", []LnsAny{Ast_AccessMode_getTxt( accessMode)}))
    }
    var argList *LnsList
    argList = NewLnsList([]LnsAny{})
    var argTypeList *LnsList
    argTypeList = NewLnsList([]LnsAny{})
    for _, _argNode := range( workArgList.Items ) {
        argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
        {
            __exp := Nodes_DeclArgNodeDownCastF(argNode.FP)
            if __exp != nil {
                _exp := __exp.(*Nodes_DeclArgNode)
                argList.Insert(Nodes_DeclArgNode2Stem(_exp))
            } else {
                self.FP.Error("macro argument can not use '...'.")
            }
        }
        var argType *Ast_TypeInfo
        argType = argNode.FP.Get_expType()
        argTypeList.Insert(Ast_TypeInfo2Stem(argType))
    }
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var retTypeList *LnsList
    if nextToken.Txt == ":"{
        retTypeList = self.FP.analyzeRefType(accessMode, true, false).FP.Get_expTypeList()
        
        self.FP.checkNextToken("{")
    } else { 
        retTypeList = NewLnsList([]LnsAny{})
        
        self.FP.checkToken(nextToken, "{")
    }
    nextToken = self.FP.getToken(nil)
    
    var stmtNode LnsAny
    stmtNode = nil
    if nextToken.Txt == "{"{
        self.macroScope = macroScope
        
        macroScope.FP.Set_validCheckingUnaccess(false)
        var funcType *Ast_TypeInfo
        funcType = self.processInfo.FP.CreateFunc(false, true, nil, Ast_TypeInfoKind__Func, Ast_headTypeInfo, false, true, true, Ast_AccessMode__Global, "_lnsLoad", nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString), Ast_TypeInfo2Stem(Ast_builtinTypeString)}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeStem)}), false)
        macroScope.FP.AddLocalVar(self.processInfo, false, false, "_lnsLoad", nil, funcType, Ast_MutMode__IMut)
        var macroLocalVarType *Ast_TypeInfo
        macroLocalVarType = self.processInfo.FP.CreateMap(Ast_AccessMode__Local, self.moduleType, Ast_builtinTypeString, Ast_builtinTypeStem, Ast_MutMode__Mut)
        macroScope.FP.AddLocalVar(self.processInfo, false, true, "__var", nil, macroLocalVarType, Ast_MutMode__IMut)
        var stmtList *LnsList
        stmtList = NewLnsList([]LnsAny{})
        self.FP.prepareTentativeSymbol(self.scope, false, nil)
        self.FP.analyzeStatementList(stmtList, "}")
        stmtNode = Nodes_BlockNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), Nodes_BlockKind__Macro, macroScope, stmtList)
        
        self.FP.checkNextToken("}")
        self.FP.finishTentativeSymbol(true)
        self.macroScope = nil
        
    } else { 
        self.FP.Pushback()
    }
    var tokenList *LnsList
    tokenList = NewLnsList([]LnsAny{})
    var braceCount LnsInt
    braceCount = 0
    for  {
        nextToken = self.FP.getToken(nil)
        
        if nextToken.Txt == "{"{
            braceCount = braceCount + 1
            
        } else if nextToken.Txt == "}"{
            if braceCount == 0{
                break
            }
            braceCount = braceCount - 1
            
        }
        tokenList.Insert(Types_Token2Stem(nextToken))
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = self.processInfo.FP.CreateFunc(false, false, macroScope, Ast_TypeInfoKind__Macro, parentType, false, false, true, accessMode, nameToken.Txt, nil, argTypeList, retTypeList, nil)
    var declMacroInfo *Nodes_DeclMacroInfo
    declMacroInfo = NewNodes_DeclMacroInfo(pubFlag, nameToken, argList, stmtNode, tokenList)
    var node *Nodes_DeclMacroNode
    node = Nodes_DeclMacroNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), declMacroInfo)
    self.macroCtrl.FP.Regist(self.processInfo, node, macroScope)
    return node
}

// 825: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclMacro
func (self *TransUnit_TransUnit) analyzeDeclMacro(accessMode LnsInt,firstToken *Types_Token) *Nodes_DeclMacroNode {
    var nameToken *Types_Token
    nameToken = self.FP.getSymbolToken(TransUnit_SymbolMode__Must_)
    self.FP.checkNextToken("(")
    var scope *Ast_Scope
    scope = Ast_TypeInfo_createScope(self.processInfo, self.topScope, false, nil, nil)
    var workArgList *LnsList
    workArgList = NewLnsList([]LnsAny{})
    self.FP.analyzeDeclArgList(accessMode, scope, workArgList, false)
    var parentInfo *Ast_TypeInfo
    parentInfo = self.FP.getCurrentNamespaceTypeInfo()
    var backScope *Ast_Scope
    backScope = self.scope
    self.scope = scope
    
    self.scope.FP.AddIgnoredVar(self.processInfo)
    var node *Nodes_DeclMacroNode
    node = self.FP.analyzeDeclMacroSub(accessMode, firstToken, nameToken, scope, parentInfo, workArgList)
    self.scope = backScope
    
    var existSym LnsAny
    _,existSym = self.scope.FP.AddMacro(self.processInfo, nameToken.Pos, node.FP.Get_expType(), accessMode)
    if Lns_isCondTrue( existSym){
        self.FP.addErrMess(nameToken.Pos, Lns_getVM().String_format("multiple define symbol -- %s", []LnsAny{nameToken.Txt}))
    }
    return node
}

// 870: decl @lune.@base.@TransUnit.TransUnit.analyzeExtend
func (self *TransUnit_TransUnit) analyzeExtend(accessMode LnsInt,firstPos *Types_Position)(*Types_Token, LnsAny, *LnsList, *LnsMap) {
    var baseRef LnsAny
    baseRef = nil
    var interfaceList *LnsList
    interfaceList = NewLnsList([]LnsAny{})
    var ifAlt2typeMap *LnsMap
    ifAlt2typeMap = NewLnsMap( map[LnsAny]LnsAny{})
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    if nextToken.Txt != "("{
        self.FP.Pushback()
        var workBaseRefType *Nodes_RefTypeNode
        workBaseRefType = self.FP.analyzeRefType(accessMode, false, Ast_isPubToExternal(accessMode))
        baseRef = &workBaseRefType.Nodes_Node
        
        var baseType *Ast_TypeInfo
        baseType = workBaseRefType.FP.Get_expType()
        if baseType.FP.Get_kind() != Ast_TypeInfoKind__Class{
            self.FP.addErrMess(workBaseRefType.FP.Get_pos(), Lns_getVM().String_format("%s is not class.", []LnsAny{baseType.FP.GetTxt(nil, nil, nil)}))
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Ast_isPubToExternal(accessMode)) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(baseType.FP.Get_accessMode()))) ).(bool)){
            self.FP.addErrMess(workBaseRefType.FP.Get_pos(), Lns_getVM().String_format("%s can't be external symbol.", []LnsAny{baseType.FP.GetTxt(nil, nil, nil)}))
        }
        nextToken = self.FP.getToken(nil)
        
    }
    if nextToken.Txt == "("{
        for  {
            nextToken = self.FP.getToken(nil)
            
            if nextToken.Txt == ")"{
                break
            }
            self.FP.Pushback()
            var ifTypeNode *Nodes_RefTypeNode
            ifTypeNode = self.FP.analyzeRefType(accessMode, false, Ast_isPubToExternal(accessMode))
            var ifType *Ast_TypeInfo
            ifType = ifTypeNode.FP.Get_expType()
            if ifType.FP.Get_kind() != Ast_TypeInfoKind__IF{
                self.FP.Error(Lns_getVM().String_format("%s is not interface -- %d", []LnsAny{ifType.FP.GetTxt(nil, nil, nil), ifType.FP.Get_kind()}))
            }
            if Ast_isGenericType(ifType){
                for _altType, _genType := range( ifType.FP.CreateAlt2typeMap(false).Items ) {
                    altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    ifAlt2typeMap.Set(altType,genType)
                }
            }
            interfaceList.Insert(Ast_TypeInfo2Stem(ifType))
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Ast_isPubToExternal(accessMode)) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(ifType.FP.Get_accessMode()))) ).(bool)){
                self.FP.addErrMess(ifTypeNode.FP.Get_pos(), Lns_getVM().String_format("%s can't be external symbol.", []LnsAny{ifType.FP.GetTxt(nil, nil, nil)}))
            }
            nextToken = self.FP.getToken(nil)
            
            if nextToken.Txt != ","{
                if nextToken.Txt == ")"{
                    break
                }
                self.FP.Error("illegal token")
            }
        }
        nextToken = self.FP.getToken(nil)
        
    }
    var symbol2TypeInfo *LnsMap
    symbol2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    for _, _ifType := range( interfaceList.Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(ifType.FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.FilterTypeInfoField(true, self.scope, self.scopeAccess, Ast_filterForm(func(symbolInfo *Ast_SymbolInfo) bool {
            {
                _ifFuncType := symbol2TypeInfo.Items[symbolInfo.FP.Get_name()]
                if _ifFuncType != nil {
                    ifFuncType := _ifFuncType.(*Ast_TypeInfo)
                    var ret bool
                    var mess LnsAny
                    ret,mess = ifFuncType.FP.CanEvalWith(self.processInfo, symbolInfo.FP.Get_typeInfo(), Ast_CanEvalType__SetOp, ifAlt2typeMap)
                    if Lns_op_not(ret){
                        self.FP.addErrMess(firstPos, Lns_getVM().String_format("mismatch method type -- %s.%s, %s.%s\n%s", []LnsAny{symbolInfo.FP.Get_typeInfo().FP.Get_parentInfo().FP.GetTxt(nil, nil, nil), symbolInfo.FP.Get_name(), ifFuncType.FP.Get_parentInfo().FP.GetTxt(nil, nil, nil), ifFuncType.FP.GetTxt(nil, nil, nil), mess}))
                    }
                } else {
                    symbol2TypeInfo.Set(symbolInfo.FP.Get_name(),symbolInfo.FP.Get_typeInfo())
                }
            }
            return true
        }))})/* 941:7 */)
    }
    var baseTypeInfo LnsAny
    baseTypeInfo = nil
    if baseRef != nil{
        baseRef_8470 := baseRef.(*Nodes_Node)
        baseTypeInfo = baseRef_8470.FP.Get_expType()
        
    }
    return nextToken, baseTypeInfo, interfaceList, ifAlt2typeMap
}

// 971: decl @lune.@base.@TransUnit.TransUnit.analyzePushClass
func (self *TransUnit_TransUnit) analyzePushClass(mode LnsInt,abstractFlag bool,firstToken *Types_Token,name *Types_Token,allowMultiple bool,requirePath LnsAny,moduleLang LnsAny,accessMode LnsInt,altTypeList *LnsList)(*Types_Token, *Ast_TypeInfo) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Ast_isPubToExternal(accessMode)) &&
        Lns_GetEnv().SetStackVal( self.moduleScope != self.scope) ).(bool)){
        self.FP.addErrMess(firstToken.Pos, "The public class must declare at top scope.")
    }
    var tempScope *Ast_Scope
    tempScope = self.FP.pushScope(false, nil, nil)
    for _, _altType := range( altTypeList.Items ) {
        altType := _altType.(Ast_AlternateTypeInfoDownCast).ToAst_AlternateTypeInfo()
        tempScope.FP.AddAlternate(self.processInfo, accessMode, altType.FP.Get_rawTxt(), name.Pos, &altType.Ast_TypeInfo)
    }
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var baseTypeInfo LnsAny
    baseTypeInfo = nil
    var interfaceList LnsAny
    interfaceList = nil
    if nextToken.Txt == "extend"{
        nextToken, baseTypeInfo, interfaceList = TransUnit_convExp37295(Lns_2DDD(self.FP.analyzeExtend(accessMode, firstToken.Pos)))
        
        if baseTypeInfo != nil{
            baseTypeInfo_8493 := baseTypeInfo.(*Ast_TypeInfo)
            {
                _initTypeInfo := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(baseTypeInfo_8493.FP.Get_scope()) && 
                Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.GetTypeInfoChild("__init")})/* 999:33 */)
                if _initTypeInfo != nil {
                    initTypeInfo := _initTypeInfo.(*Ast_TypeInfo)
                    if initTypeInfo.FP.Get_accessMode() == Ast_AccessMode__Pri{
                        self.FP.addErrMess(firstToken.Pos, "The access mode of '__init' is 'pri'.")
                    }
                }
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( baseTypeInfo_8493.FP.IsInheritFrom(self.processInfo, TransUnit_builtinFunc.Lnsthread_, nil)) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(self.helperInfo.PragmaSet.Has(LuneControl_Pragma__use_async_Obj))) ).(bool)){
                self.FP.addErrMess(nextToken.Pos, "must set '_lune_control use_async'")
            }
        }
    }
    self.FP.popScope()
    var classTypeInfo *Ast_TypeInfo
    if _switch37474 := mode; _switch37474 == TransUnit_DeclClassMode__Module || _switch37474 == TransUnit_DeclClassMode__LazyModule {
        _ = self.scope
        classTypeInfo = self.FP.pushExtModule(false, name.Txt, accessMode, name.Pos, mode == TransUnit_DeclClassMode__LazyModule, Lns_unwrap( moduleLang).(LnsInt), (Lns_unwrap( requirePath).(*Types_Token)).FP.GetExcludedDelimitTxt())
        
    } else if _switch37474 == TransUnit_DeclClassMode__Class || _switch37474 == TransUnit_DeclClassMode__Interface {
        classTypeInfo = self.FP.pushClass(firstToken.Pos, mode, abstractFlag, baseTypeInfo, interfaceList, altTypeList, false, name.Txt, allowMultiple, accessMode, nil)
        
    }
    return nextToken, classTypeInfo
}

// 1037: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclAlternateType
func (self *TransUnit_TransUnit) analyzeDeclAlternateType(belongClassFlag bool,token *Types_Token,accessMode LnsInt)(*Types_Token, *LnsList) {
    var altTypeList *LnsList
    altTypeList = NewLnsList([]LnsAny{})
    var nextToken *Types_Token
    nextToken = token
    var altNameSet *LnsSet
    altNameSet = NewLnsSet([]LnsAny{})
    var altIndex LnsInt
    altIndex = 0
    for  {
        altIndex = altIndex + 1
        
        var genericSymToken *Types_Token
        genericSymToken = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
        if Lns_isCondTrue( self.scope.FP.GetTypeInfo(genericSymToken.Txt, self.scope, true, self.scopeAccess)){
            self.FP.addErrMess(genericSymToken.Pos, Lns_getVM().String_format("shadowing Type -- %s", []LnsAny{genericSymToken.Txt}))
        } else { 
            if altNameSet.Has(genericSymToken.Txt){
                self.FP.addErrMess(genericSymToken.Pos, Lns_getVM().String_format("multiple Type -- %s", []LnsAny{genericSymToken.Txt}))
            } else { 
                altNameSet.Add(genericSymToken.Txt)
            }
        }
        var workToken *Types_Token
        workToken = self.FP.getToken(nil)
        if workToken.Txt == "!"{
            self.FP.addErrMess(workToken.Pos, "not support nilable")
            workToken = self.FP.getToken(nil)
            
        }
        var baseTypeInfo LnsAny
        baseTypeInfo = nil
        var interfaceList *LnsList
        interfaceList = NewLnsList([]LnsAny{})
        if workToken.Txt == ":"{
            workToken, baseTypeInfo, interfaceList = TransUnit_convExp37723(Lns_2DDD(self.FP.analyzeExtend(accessMode, token.Pos)))
            
        }
        var altType *Ast_AlternateTypeInfo
        altType = self.processInfo.FP.CreateAlternate(belongClassFlag, altIndex, genericSymToken.Txt, accessMode, self.moduleType, baseTypeInfo, interfaceList)
        altTypeList.Insert(Ast_AlternateTypeInfo2Stem(altType))
        if workToken.Txt == ">"{
            nextToken = self.FP.getToken(nil)
            
            break
        }
        self.FP.checkToken(workToken, ",")
    }
    return nextToken, altTypeList
}

// 1089: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclProto
func (self *TransUnit_TransUnit) analyzeDeclProto(accessMode LnsInt,firstToken *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var abstractFlag bool
    abstractFlag = false
    if nextToken.Txt == "abstract"{
        abstractFlag = true
        
        nextToken = self.FP.getToken(nil)
        
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( nextToken.Txt == "class") ||
        Lns_GetEnv().SetStackVal( nextToken.Txt == "interface") ).(bool){
        var name *Types_Token
        name = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
        var altTypeList *LnsList
        altTypeList = NewLnsList([]LnsAny{})
        {
            var workToken *Types_Token
            workToken = self.FP.getToken(nil)
            if workToken.Txt == "<"{
                workToken, altTypeList = self.FP.analyzeDeclAlternateType(true, workToken, accessMode)
                
            }
            self.FP.PushbackToken(workToken)
        }
        if accessMode == Ast_AccessMode__Local{
            accessMode = Ast_AccessMode__Pri
            
        }
        var declMode LnsInt
        if nextToken.Txt == "class"{
            declMode = TransUnit_DeclClassMode__Class
            
        } else { 
            declMode = TransUnit_DeclClassMode__Interface
            
        }
        var classTypeInfo *Ast_TypeInfo
        nextToken, classTypeInfo = self.FP.analyzePushClass(declMode, abstractFlag, firstToken, name, false, nil, nil, accessMode, altTypeList)
        
        self.protoClassMap.Set(classTypeInfo,firstToken.Pos)
        self.FP.popClass()
        self.FP.checkToken(nextToken, ";")
        return &Nodes_ProtoClassNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo)}), name).Nodes_Node
    }
    self.FP.Error("illegal proto")
// insert a dummy
    return nil
}

// 1141: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclEnum
func (self *TransUnit_TransUnit) analyzeDeclEnum(accessMode LnsInt,firstToken *Types_Token) *Nodes_DeclEnumNode {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Ast_isPubToExternal(accessMode)) &&
        Lns_GetEnv().SetStackVal( self.scope != self.moduleScope) &&
        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self.scope.FP.Get_ownerTypeInfo()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_kind()})) != Ast_TypeInfoKind__Class) ).(bool)){
        self.FP.addErrMess(firstToken.Pos, "can't external at inner scope.")
    }
    var name *Types_Token
    name = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
    self.FP.checkNextToken("{")
    var valueList *LnsList
    valueList = NewLnsList([]LnsAny{})
    var scope *Ast_Scope
    scope = self.FP.pushScope(true, nil, nil)
    var workEnumTypeInfo LnsAny
    workEnumTypeInfo = nil
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var number LnsReal
    number = 0.0
    var prevValTypeInfo *Ast_TypeInfo
    prevValTypeInfo = Ast_headTypeInfo
    var valTypeInfo *Ast_TypeInfo
    valTypeInfo = Ast_headTypeInfo
    for nextToken.Txt != "}" {
        var valName *Types_Token
        valName = self.FP.checkSymbol(nextToken, TransUnit_SymbolMode__MustNot_)
        nextToken = self.FP.getToken(nil)
        
        var enumVal LnsAny
        enumVal = &Ast_EnumLiteral__Real{number}
        if _switch38287 := (prevValTypeInfo); _switch38287 == Ast_builtinTypeReal {
        } else if _switch38287 == Ast_builtinTypeInt || _switch38287 == Ast_headTypeInfo {
            enumVal = &Ast_EnumLiteral__Int{(LnsInt)(number)}
            
        }
        if nextToken.Txt == "="{
            var exp *Nodes_Node
            exp = self.FP.analyzeExpOneRVal(false, false, nil, nil)
            var literal LnsAny
            var mess LnsAny
            literal,mess = exp.FP.GetLiteral()
            if literal != nil{
                literal_8571 := literal
                switch _exp38409 := literal_8571.(type) {
                case *Nodes_Literal__Int:
                val := _exp38409.Val1
                    enumVal = &Ast_EnumLiteral__Int{val}
                    
                    number = (LnsReal)(val)
                    
                    valTypeInfo = Ast_builtinTypeInt
                    
                case *Nodes_Literal__Real:
                val := _exp38409.Val1
                    enumVal = &Ast_EnumLiteral__Real{val}
                    
                    number = val
                    
                    valTypeInfo = Ast_builtinTypeReal
                    
                case *Nodes_Literal__Str:
                val := _exp38409.Val1
                    enumVal = &Ast_EnumLiteral__Str{val}
                    
                    valTypeInfo = Ast_builtinTypeString
                    
                default:
                    self.FP.Error(Lns_getVM().String_format("illegal enum val -- %s", []LnsAny{literal_8571.(LnsAlgeVal).GetTxt()}))
                }
            } else {
                self.FP.Error(Lns_getVM().String_format("illegal enum val -- %s", []LnsAny{mess}))
            }
            nextToken = self.FP.getToken(nil)
            
        } else { 
            if _switch38497 := (prevValTypeInfo); _switch38497 == Ast_headTypeInfo {
                valTypeInfo = Ast_builtinTypeInt
                
            } else if _switch38497 == Ast_builtinTypeInt || _switch38497 == Ast_builtinTypeReal {
                valTypeInfo = prevValTypeInfo
                
            } else {
                self.FP.addErrMess(valName.Pos, Lns_getVM().String_format("illegal enum val type -- %s", []LnsAny{valTypeInfo.FP.GetTxt(nil, nil, nil)}))
            }
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( prevValTypeInfo != Ast_headTypeInfo) &&
            Lns_GetEnv().SetStackVal( prevValTypeInfo != valTypeInfo) ).(bool)){
            self.FP.addErrMess(valName.Pos, Lns_getVM().String_format("multiple enum val type. %s, %s", []LnsAny{valTypeInfo.FP.GetTxt(nil, nil, nil), prevValTypeInfo.FP.GetTxt(nil, nil, nil)}))
        }
        prevValTypeInfo = valTypeInfo
        
        if Lns_op_not(workEnumTypeInfo){
            workEnumTypeInfo = self.processInfo.FP.CreateEnum(scope, self.FP.getCurrentNamespaceTypeInfo(), false, accessMode, name.Txt, valTypeInfo)
            
        }
        if workEnumTypeInfo != nil{
            workEnumTypeInfo_8587 := workEnumTypeInfo.(*Ast_EnumTypeInfo)
            scope.FP.AddEnumVal(self.processInfo, valName.Txt, valName.Pos, &workEnumTypeInfo_8587.Ast_TypeInfo)
            var enumValInfo *Ast_EnumValInfo
            enumValInfo = NewAst_EnumValInfo(valName.Txt, enumVal)
            valueList.Insert(Types_Token2Stem(valName))
            workEnumTypeInfo_8587.FP.AddEnumValInfo(enumValInfo)
        }
        if nextToken.Txt == "}"{
            break
        }
        self.FP.checkToken(nextToken, ",")
        nextToken = self.FP.getToken(nil)
        
        number = number + LnsReal(1)
        
    }
    var enumTypeInfo *Ast_EnumTypeInfo
    
    {
        _enumTypeInfo := workEnumTypeInfo
        if _enumTypeInfo == nil{
            enumTypeInfo = self.processInfo.FP.CreateEnum(scope, self.FP.getCurrentNamespaceTypeInfo(), false, accessMode, name.Txt, Ast_builtinTypeNone)
            
        } else {
            enumTypeInfo = _enumTypeInfo.(*Ast_EnumTypeInfo)
        }
    }
    self.FP.popScope()
    var shadowing LnsAny
    _,shadowing = self.scope.FP.AddEnum(self.processInfo, accessMode, name.Txt, name.Pos, &enumTypeInfo.Ast_TypeInfo)
    self.FP.errorShadowing(name.Pos, shadowing)
    return Nodes_DeclEnumNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_EnumTypeInfo2Stem(enumTypeInfo)}), enumTypeInfo, accessMode, name, valueList, scope)
}

// 1277: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclAlge
func (self *TransUnit_TransUnit) analyzeDeclAlge(accessMode LnsInt,firstToken *Types_Token) *Nodes_DeclAlgeNode {
    self.helperInfo.UseAlge = true
    
    var name *Types_Token
    name = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
    self.FP.checkNextToken("{")
    var scope *Ast_Scope
    scope = self.scope
    var algeScope *Ast_Scope
    algeScope = self.FP.pushScope(true, nil, nil)
    var algeTypeInfo *Ast_AlgeTypeInfo
    algeTypeInfo = self.processInfo.FP.CreateAlge(algeScope, self.FP.getCurrentNamespaceTypeInfo(), false, accessMode, name.Txt)
    var shadowing LnsAny
    _,shadowing = scope.FP.AddAlge(self.processInfo, accessMode, name.Txt, name.Pos, &algeTypeInfo.Ast_TypeInfo)
    self.FP.errorShadowing(name.Pos, shadowing)
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    for nextToken.Txt != "}" {
        var valName *Types_Token
        valName = self.FP.checkSymbol(nextToken, TransUnit_SymbolMode__MustNot_)
        if Lns_isCondTrue( algeTypeInfo.FP.GetValInfo(valName.Txt)){
            self.FP.addErrMess(valName.Pos, Lns_getVM().String_format("multiple symbole -- %s", []LnsAny{valName.Txt}))
        }
        nextToken = self.FP.getToken(nil)
        
        var typeInfoList *LnsList
        typeInfoList = NewLnsList([]LnsAny{})
        if nextToken.Txt == "("{
            for  {
                _ = self.FP.getToken(nil)
                var workToken2 *Types_Token
                workToken2 = self.FP.getToken(nil)
                if workToken2.Txt != ":"{
                    self.FP.Pushback()
                    self.FP.Pushback()
                }
                var typeNode *Nodes_RefTypeNode
                typeNode = self.FP.analyzeRefType(Ast_AccessMode__Pub, false, Ast_isPubToExternal(accessMode))
                if Lns_isCondTrue( self.protoClassMap.Items[typeNode.FP.Get_expType()]){
                    self.FP.addErrMess(typeNode.FP.Get_pos(), Lns_getVM().String_format("can't use the prototype class -- %s", []LnsAny{typeNode.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
                }
                typeInfoList.Insert(Ast_TypeInfo2Stem(typeNode.FP.Get_expType()))
                nextToken = self.FP.getToken(nil)
                
                if nextToken.Txt != ","{
                    self.FP.checkToken(nextToken, ")")
                    nextToken = self.FP.getToken(nil)
                    
                    break
                }
            }
        }
        algeScope.FP.AddAlgeVal(self.processInfo, valName.Txt, valName.Pos, &algeTypeInfo.Ast_TypeInfo)
        var algeValInfo *Ast_AlgeValInfo
        algeValInfo = NewAst_AlgeValInfo(valName.Txt, typeInfoList, algeTypeInfo)
        algeTypeInfo.FP.AddValInfo(algeValInfo)
        if nextToken.Txt == "}"{
            break
        }
        self.FP.checkToken(nextToken, ",")
        nextToken = self.FP.getToken(nil)
        
    }
    self.FP.popScope()
    return Nodes_DeclAlgeNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_AlgeTypeInfo2Stem(algeTypeInfo)}), accessMode, algeTypeInfo, algeScope)
}

// 1353: decl @lune.@base.@TransUnit.TransUnit.analyzeAlias
func (self *TransUnit_TransUnit) analyzeAlias(accessMode LnsInt,firstToken *Types_Token) *Nodes_AliasNode {
    if self.scope != self.moduleScope{
        self.FP.addErrMess(firstToken.Pos, "alias must use at top scope.")
    }
    var newToken *Types_Token
    newToken = self.FP.getToken(nil)
    self.FP.checkNextToken("=")
    var srcToken *Types_Token
    srcToken = self.FP.getToken(nil)
    var symbolNode *Nodes_Node
    symbolNode = self.FP.analyzeExpSymbol(firstToken, srcToken, TransUnit_ExpSymbolMode__Symbol, nil, true, false)
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = Ast_builtinTypeNone
    var symbolInfoList *LnsList
    symbolInfoList = symbolNode.FP.GetSymbolInfo()
    var newSymbolInfo *Ast_SymbolInfo
    newSymbolInfo = Ast_dummySymbol
    if symbolInfoList.Len() >= 1{
        var symbolInfo *Ast_SymbolInfo
        symbolInfo = symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(newToken.Txt,"^_", nil, nil))) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Lns_getVM().String_find(srcToken.Txt,"^_", nil, nil)))) ||
            Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Lns_getVM().String_find(newToken.Txt,"^_", nil, nil)))) &&
            Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(srcToken.Txt,"^_", nil, nil))) )){
            self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("alias symbol unmatch. %s %s", []LnsAny{newToken.Txt, newToken.Txt}))
        } else { 
            if _switch39546 := symbolInfo.FP.Get_kind(); _switch39546 == Ast_SymbolKind__Typ || _switch39546 == Ast_SymbolKind__Fun {
                var aliasSymbolInfo LnsAny
                var shadowing LnsAny
                aliasSymbolInfo,shadowing = self.scope.FP.AddAlias(self.processInfo, newToken.Txt, newToken.Pos, false, accessMode, self.moduleType, symbolInfo)
                if aliasSymbolInfo != nil{
                    aliasSymbolInfo_8643 := aliasSymbolInfo.(*Ast_SymbolInfo)
                    newTypeInfo = aliasSymbolInfo_8643.FP.Get_typeInfo()
                    
                    newSymbolInfo = aliasSymbolInfo_8643
                    
                } else {
                    self.FP.errorShadowing(newToken.Pos, shadowing)
                }
            } else {
                self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("can alias symbol -- %s. (%s)", []LnsAny{srcToken.Txt, Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind())}))
            }
        }
    } else { 
        self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("not found symbold -- %s", []LnsAny{srcToken.Txt}))
    }
    self.FP.checkNextToken(";")
    return Nodes_AliasNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(newTypeInfo)}), newSymbolInfo, symbolNode, newTypeInfo)
}

// 1417: decl @lune.@base.@TransUnit.TransUnit.analyzeRetTypeList
func (self *TransUnit_TransUnit) analyzeRetTypeList(pubToExtFlag bool,accessMode LnsInt,token *Types_Token,parentPub bool)(*LnsList, *Types_Token) {
    var retTypeInfoList *LnsList
    retTypeInfoList = NewLnsList([]LnsAny{})
    if token.Txt == ":"{
        var hasDDDFlag bool
        hasDDDFlag = false
        for  {
            var refTypeNode *Nodes_RefTypeNode
            refTypeNode = self.FP.analyzeRefType(accessMode, true, parentPub)
            if hasDDDFlag{
                self.FP.addErrMess(refTypeNode.FP.Get_pos(), "Type exists after '...'.")
            }
            var retType *Ast_TypeInfo
            retType = refTypeNode.FP.Get_expType()
            if retType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                hasDDDFlag = true
                
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( pubToExtFlag) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(retType.FP.Get_accessMode()))) ).(bool)){
                self.FP.addErrMess(refTypeNode.FP.Get_pos(), Lns_getVM().String_format("this is not public type -- %s", []LnsAny{retType.FP.GetTxt(nil, nil, nil)}))
            }
            retTypeInfoList.Insert(Ast_TypeInfo2Stem(retType))
            token = self.FP.getToken(nil)
            
            if token.Txt != ","{
                break
            }
        }
    }
    return retTypeInfoList, token
}

// 1450: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclForm
func (self *TransUnit_TransUnit) analyzeDeclForm(accessMode LnsInt,firstToken *Types_Token) *Nodes_DeclFormNode {
    var name *Types_Token
    name = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.scope != self.moduleScope) &&
        Lns_GetEnv().SetStackVal( Ast_isPubToExternal(accessMode)) ).(bool)){
        self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("You must declare this form at the outside scope. -- %s", []LnsAny{name.Txt}))
    }
    self.FP.checkNextToken("(")
    var argList *LnsList
    argList = NewLnsList([]LnsAny{})
    var funcBodyScope *Ast_Scope
    funcBodyScope = self.FP.pushScope(false, nil, nil)
    var nextToken *Types_Token
    nextToken = self.FP.analyzeDeclArgList(accessMode, funcBodyScope, argList, Ast_isPubToExternal(accessMode))
    self.FP.checkToken(nextToken, ")")
    var retTypeList *LnsList
    retTypeList = NewLnsList([]LnsAny{})
    nextToken = self.FP.getToken(nil)
    
    retTypeList, nextToken = self.FP.analyzeRetTypeList(Ast_isPubToExternal(accessMode), accessMode, nextToken, Ast_isPubToExternal(accessMode))
    
    self.FP.checkToken(nextToken, ";")
    self.FP.popScope()
    var argTypeInfoList *LnsList
    argTypeInfoList = NewLnsList([]LnsAny{})
    for _, _argNode := range( argList.Items ) {
        argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
        argTypeInfoList.Insert(Ast_TypeInfo2Stem(argNode.FP.Get_expType()))
    }
    var formType *Ast_TypeInfo
    formType = self.processInfo.FP.CreateFunc(false, false, nil, Ast_TypeInfoKind__FormFunc, self.FP.getCurrentNamespaceTypeInfo(), false, false, true, accessMode, name.Txt, nil, argTypeInfoList, retTypeList, false)
    self.scope.FP.AddForm(self.processInfo, name.Pos, formType, accessMode)
    return Nodes_DeclFormNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(formType)}), argList)
}

// 1501: decl @lune.@base.@TransUnit.TransUnit.analyzeDecl
func (self *TransUnit_TransUnit) analyzeDecl(accessMode LnsInt,staticFlag bool,firstToken *Types_Token,token *Types_Token) LnsAny {
    if Lns_op_not(staticFlag){
        if token.Txt == "static"{
            staticFlag = true
            
            token = self.FP.getToken(nil)
            
        }
    }
    var overrideFlag bool
    overrideFlag = false
    if token.Txt == "override"{
        overrideFlag = true
        
        token = self.FP.getToken(nil)
        
    }
    var abstractFlag bool
    abstractFlag = false
    if token.Txt == "abstract"{
        abstractFlag = true
        
        token = self.FP.getToken(nil)
        
    }
    if token.Txt == "let"{
        return self.FP.analyzeDeclVar(Nodes_DeclVarMode__Let, accessMode, firstToken)
    } else if token.Txt == "fn"{
        var nextToken *Types_Token
        nextToken = self.FP.getToken(nil)
        self.FP.Pushback()
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( nextToken.Kind == Types_TokenKind__Symb) ||
            Lns_GetEnv().SetStackVal( Ast_isPubToExternal(accessMode)) ||
            Lns_GetEnv().SetStackVal( staticFlag) ||
            Lns_GetEnv().SetStackVal( overrideFlag) ||
            Lns_GetEnv().SetStackVal( abstractFlag) ).(bool){
            return self.FP.analyzeDeclFunc(TransUnit_DeclFuncMode__Func, abstractFlag, overrideFlag, accessMode, staticFlag, nil, firstToken, nil)
        }
    } else if token.Txt == "class"{
        return &self.FP.analyzeDeclClass(abstractFlag, accessMode, firstToken, TransUnit_DeclClassMode__Class).Nodes_Node
    } else if token.Txt == "interface"{
        return &self.FP.analyzeDeclClass(true, accessMode, firstToken, TransUnit_DeclClassMode__Interface).Nodes_Node
    } else if token.Txt == "module"{
        return &self.FP.analyzeDeclClass(false, accessMode, firstToken, TransUnit_DeclClassMode__Module).Nodes_Node
    } else if token.Txt == "proto"{
        return self.FP.analyzeDeclProto(accessMode, firstToken)
    } else if token.Txt == "macro"{
        return &self.FP.analyzeDeclMacro(accessMode, firstToken).Nodes_Node
    } else if token.Txt == "enum"{
        return &self.FP.analyzeDeclEnum(accessMode, firstToken).Nodes_Node
    } else if token.Txt == "alge"{
        return &self.FP.analyzeDeclAlge(accessMode, firstToken).Nodes_Node
    } else if token.Txt == "form"{
        return &self.FP.analyzeDeclForm(accessMode, firstToken).Nodes_Node
    } else if token.Txt == "alias"{
        return &self.FP.analyzeAlias(accessMode, firstToken).Nodes_Node
    } else if token.Txt == "__test"{
        return self.FP.analyzeTest(firstToken)
    }
    return nil
}

// 1579: decl @lune.@base.@TransUnit.TransUnit.checkPublic
func (self *TransUnit_TransUnit) checkPublic(pos *Types_Position,typeInfo *Ast_TypeInfo) {
    var checkedTypeSet *LnsSet
    checkedTypeSet = NewLnsSet([]LnsAny{})
    var checkPub func(workType *Ast_TypeInfo)
    checkPub = func(workType *Ast_TypeInfo) {
        if checkedTypeSet.Has(Ast_TypeInfo2Stem(workType)){
            return 
        }
        checkedTypeSet.Add(Ast_TypeInfo2Stem(workType))
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( workType.FP.Get_kind() != Ast_TypeInfoKind__Array) &&
            Lns_GetEnv().SetStackVal( workType.FP.Get_kind() != Ast_TypeInfoKind__List) &&
            Lns_GetEnv().SetStackVal( workType.FP.Get_kind() != Ast_TypeInfoKind__Set) &&
            Lns_GetEnv().SetStackVal( workType.FP.Get_kind() != Ast_TypeInfoKind__Map) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(workType.FP.Get_accessMode()))) ).(bool)){
            self.FP.addErrMess(pos, Lns_getVM().String_format("not public this type -- %s", []LnsAny{workType.FP.GetTxt(nil, nil, nil)}))
        } else { 
            for _, _itemType := range( workType.FP.Get_itemTypeInfoList().Items ) {
                itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                checkPub(itemType)
            }
        }
    }
    checkPub(typeInfo)
}

// 1601: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclMember
func (self *TransUnit_TransUnit) analyzeDeclMember(classTypeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,firstToken *Types_Token) *Nodes_DeclMemberNode {
    __func__ := "@lune.@base.@TransUnit.TransUnit.analyzeDeclMember"
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var mutMode LnsInt
    mutMode = Ast_MutMode__IMut
    if _switch40733 := nextToken.Txt; _switch40733 == "mut" {
        mutMode = Ast_MutMode__Mut
        
        nextToken = self.FP.getToken(nil)
        
    } else if _switch40733 == "allmut" {
        mutMode = Ast_MutMode__AllMut
        
        nextToken = self.FP.getToken(nil)
        
    }
    var varName *Types_Token
    varName = self.FP.checkSymbol(nextToken, TransUnit_SymbolMode__MustNot_)
    var token *Types_Token
    token = self.FP.getToken(nil)
    var refType *Nodes_RefTypeNode
    refType = self.FP.analyzeRefType(accessMode, false, Ast_isPubToExternal(classTypeInfo.FP.Get_accessMode()))
    token = self.FP.getToken(nil)
    
    var getterMode LnsInt
    getterMode = Ast_AccessMode__None
    var getterRetType *Ast_TypeInfo
    getterRetType = refType.FP.Get_expType()
    var getterMutable bool
    getterMutable = true
    var setterMode LnsInt
    setterMode = Ast_AccessMode__None
    if token.Txt == "{"{
        var analyzeAccessorMode func()(LnsInt, *Ast_TypeInfo, *Types_Token)
        analyzeAccessorMode = func()(LnsInt, *Ast_TypeInfo, *Types_Token) {
            var retType *Ast_TypeInfo
            retType = Ast_headTypeInfo
            var mode LnsInt
            mode = Ast_AccessMode__None
            var workToken *Types_Token
            workToken = self.FP.getToken(nil)
            if _switch40991 := workToken.Txt; _switch40991 == "pub" || _switch40991 == "pri" || _switch40991 == "pro" || _switch40991 == "local" {
                mode = Lns_unwrap( Ast_txt2AccessMode(workToken.Txt)).(LnsInt)
                
                workToken = self.FP.getToken(nil)
                
                if workToken.Txt == "&"{
                    getterMutable = false
                    
                    workToken = self.FP.getToken(nil)
                    
                }
                if workToken.Txt == ":"{
                    var typeNode *Nodes_RefTypeNode
                    typeNode = self.FP.analyzeRefType(mode, false, Ast_isPubToExternal(classTypeInfo.FP.Get_accessMode()))
                    retType = typeNode.FP.Get_expType()
                    
                    workToken = self.FP.getToken(nil)
                    
                }
            } else if _switch40991 == "non" {
                workToken = self.FP.getToken(nil)
                
            } else {
                self.FP.addErrMess(workToken.Pos, Lns_getVM().String_format("access mode is invalid -- %s", []LnsAny{workToken.Txt}))
            }
            return mode, retType, workToken
        }
        {
            var workRetType *Ast_TypeInfo
            getterMode, workRetType, nextToken = analyzeAccessorMode()
            
            if workRetType != Ast_headTypeInfo{
                if Lns_op_not(Lns_car(workRetType.FP.CanEvalWith(self.processInfo, getterRetType, Ast_CanEvalType__SetOp, classTypeInfo.FP.CreateAlt2typeMap(false))).(bool)){
                    self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("getter type mismatch -- %s <- %s", []LnsAny{workRetType.FP.GetTxt(nil, nil, nil), getterRetType.FP.GetTxt(nil, nil, nil)}))
                }
                getterRetType = workRetType
                
            }
        }
        if nextToken.Txt == ","{
            var dummyRetType *Ast_TypeInfo
            setterMode, dummyRetType, nextToken = analyzeAccessorMode()
            
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( setterMode != Ast_AccessMode__None) &&
                Lns_GetEnv().SetStackVal( mutMode == Ast_MutMode__IMut) ).(bool)){
                self.FP.addErrMess(varName.Pos, Lns_getVM().String_format("This member can't have setter, this member is immutable. -- %s", []LnsAny{varName.Txt}))
            }
            Log_log(Log_Level__Debug, __func__, 1687, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("%s", []LnsAny{dummyRetType})
            }))
            
        }
        self.FP.checkToken(nextToken, "}")
        token = self.FP.getToken(nil)
        
    }
    self.FP.checkToken(token, ";")
    var typeInfo *Ast_TypeInfo
    typeInfo = refType.FP.Get_expType()
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(typeInfo)) &&
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_mutMode() != mutMode) ).(bool)){
        typeInfo = self.FP.CreateModifier(typeInfo, mutMode)
        
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(getterRetType)) &&
        Lns_GetEnv().SetStackVal( getterRetType.FP.Get_mutMode() != mutMode) ).(bool)){
        getterRetType = self.FP.CreateModifier(getterRetType, mutMode)
        
    }
    if Ast_isPubToExternal(classTypeInfo.FP.Get_accessMode()){
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Ast_isPubToExternal(accessMode)) ||
            Lns_GetEnv().SetStackVal( Ast_isPubToExternal(setterMode)) ).(bool){
            self.FP.checkPublic(refType.FP.Get_pos(), typeInfo)
        }
        if Ast_isPubToExternal(getterMode){
            self.FP.checkPublic(refType.FP.Get_pos(), getterRetType)
        }
    }
    var symbolInfo LnsAny
    var shadowing LnsAny
    symbolInfo,shadowing = self.scope.FP.AddMember(self.processInfo, varName.Txt, varName.Pos, typeInfo, accessMode, staticFlag, mutMode)
    var workSym *Ast_SymbolInfo
    workSym = Lns_unwrap( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( symbolInfo) ||
        Lns_GetEnv().SetStackVal( shadowing) ))).(*Ast_SymbolInfo)
    if shadowing != nil{
        shadowing_8776 := shadowing.(*Ast_SymbolInfo)
        self.FP.errorShadowing(varName.Pos, shadowing_8776)
    }
    return Nodes_DeclMemberNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), varName, refType, workSym, classTypeInfo, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode)
}

// 1731: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclMethod
func (self *TransUnit_TransUnit) analyzeDeclMethod(classTypeInfo *Ast_TypeInfo,declFuncMode LnsInt,abstractFlag bool,overrideFlag bool,accessMode LnsInt,staticFlag bool,firstToken *Types_Token,name *Types_Token) *Nodes_Node {
    var node *Nodes_Node
    node = self.FP.analyzeDeclFunc(declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, classTypeInfo, name, name)
    return node
}

// 1752: decl @lune.@base.@TransUnit.TransUnit.addDefaultConstructor
func (self *TransUnit_TransUnit) addDefaultConstructor(pos *Types_Position,classTypeInfo *Ast_TypeInfo,classScope *Ast_Scope,memberNodeList *LnsList,methodNameSet *LnsSet,oldFlag bool) {
    if Lns_isCondTrue( classScope.FP.GetTypeInfoChild("__init")){
        self.FP.addErrMess(pos, "already declare __init().")
    }
    var argTypeList *LnsList
    argTypeList = NewLnsList([]LnsAny{})
    if classTypeInfo.FP.Get_baseTypeInfo() != Ast_headTypeInfo{
        var superScope *Ast_Scope
        superScope = Lns_unwrap( classTypeInfo.FP.Get_baseTypeInfo().FP.Get_scope()).(*Ast_Scope)
        var superTypeInfo *Ast_TypeInfo
        superTypeInfo = Lns_unwrap( superScope.FP.GetTypeInfoChild("__init")).(*Ast_TypeInfo)
        for _, _argType := range( superTypeInfo.FP.Get_argTypeInfoList().Items ) {
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if oldFlag{
                if Lns_op_not(argType.FP.Get_nilable()){
                    self.FP.addErrMess(pos, "not found '__init' decl.")
                }
            } else { 
                argTypeList.Insert(Ast_TypeInfo2Stem(argType))
            }
        }
    }
    for _, _memberNode := range( memberNodeList.Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if Lns_op_not(memberNode.FP.Get_staticFlag()){
            argTypeList.Insert(Ast_TypeInfo2Stem(memberNode.FP.Get_expType()))
        }
    }
    if Ast_isPubToExternal(classTypeInfo.FP.Get_accessMode()){
        for _, _memberType := range( argTypeList.Items ) {
            memberType := _memberType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(Ast_isPubToExternal(memberType.FP.Get_accessMode())){
                self.FP.addErrMess(pos, Lns_getVM().String_format("The type must be 'pub' becaue using in __init(). -- %s:%s", []LnsAny{memberType.FP.GetTxt(nil, nil, nil), Ast_AccessMode_getTxt( memberType.FP.Get_accessMode())}))
            }
        }
    }
    var ctorScope *Ast_Scope
    ctorScope = self.FP.pushScope(false, nil, nil)
    var initTypeInfo *Ast_TypeInfo
    initTypeInfo = self.processInfo.FP.CreateFunc(false, false, ctorScope, Ast_TypeInfoKind__Method, classTypeInfo, true, false, false, Ast_AccessMode__Pub, "__init", nil, argTypeList, NewLnsList([]LnsAny{}), nil)
    if oldFlag{
        ctorScope.FP.AddVar(self.processInfo, Ast_AccessMode__Pri, "", nil, Ast_headTypeInfo, Ast_MutMode__IMut, true)
    }
    self.FP.popScope()
    classScope.FP.AddMethod(self.processInfo, pos, initTypeInfo, Ast_AccessMode__Pub, false, false)
    methodNameSet.Add("__init")
    for _, _memberNode := range( memberNodeList.Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if Lns_op_not(memberNode.FP.Get_staticFlag()){
            memberNode.FP.Get_symbolInfo().FP.UpdateValue(memberNode.FP.Get_symbolInfo().FP.Get_posForLatestMod())
        }
    }
}

// 1828: decl @lune.@base.@TransUnit.TransUnit.analyzeFuncBlock
func (self *TransUnit_TransUnit) analyzeFuncBlock(analyzingState LnsInt,firstToken *Types_Token,classTypeInfo LnsAny,funcTypeInfo *Ast_TypeInfo,funcName string,funcBodyScope *Ast_Scope,retTypeInfoList *LnsList) *Nodes_BlockNode {
    if Lns_op_not(funcTypeInfo.FP.Get_staticFlag()){
        if Lns_isCondTrue( classTypeInfo){
            {
                _overrideType := self.scope.FP.Get_parent().FP.GetTypeInfoField(funcName, false, funcBodyScope, self.scopeAccess)
                if _overrideType != nil {
                    overrideType := _overrideType.(*Ast_TypeInfo)
                    if Lns_op_not(overrideType.FP.Get_abstractFlag()){
                        funcBodyScope.FP.Add(self.processInfo, Ast_SymbolKind__Fun, false, false, "super", nil, overrideType, Ast_AccessMode__Local, false, Ast_MutMode__IMut, true, false)
                    }
                }
            }
        }
    }
    self.FP.pushAnalyzingState(analyzingState)
    var body *Nodes_BlockNode
    body = self.FP.analyzeBlock(Nodes_BlockKind__Func, TransUnit_TentativeMode__Ignore, funcBodyScope, nil)
    self.FP.popAnalyzingState()
    if retTypeInfoList.Len() != 0{
        var breakKind LnsInt
        breakKind = body.FP.GetBreakKind(Nodes_CheckBreakMode__Return)
        if retTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeNeverRet{
            if _switch42032 := breakKind; _switch42032 == Nodes_BreakKind__Return || _switch42032 == Nodes_BreakKind__NeverRet {
            } else {
                self.FP.addErrMess(firstToken.Pos, "This funcion doesn't have return.")
            }
        } else { 
            if breakKind != Nodes_BreakKind__NeverRet{
                self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("This funcion must be never return. -- %s", []LnsAny{Nodes_BreakKind_getTxt( breakKind)}))
            }
        }
    }
    return body
}

// 1875: decl @lune.@base.@TransUnit.TransUnit.addAccessor
func (self *TransUnit_TransUnit) addAccessor(memberNode *Nodes_DeclMemberNode,methodNameSet *LnsSet,classScope *Ast_Scope,classTypeInfo *Ast_TypeInfo) {
    var memberType *Ast_TypeInfo
    memberType = memberNode.FP.Get_expType()
    var memberName *Types_Token
    memberName = memberNode.FP.Get_name()
    var getterName string
    getterName = "get_" + memberName.Txt
    var accessMode LnsInt
    accessMode = memberNode.FP.Get_getterMode()
    var typeKind LnsInt
    if memberNode.FP.Get_staticFlag(){
        typeKind = Ast_TypeInfoKind__Func
        
    } else { 
        typeKind = Ast_TypeInfoKind__Method
        
    }
    if accessMode != Ast_AccessMode__None{
        if Lns_isCondTrue( classScope.FP.GetTypeInfoChild(getterName)){
            self.FP.addErrMess(memberName.Pos, Lns_getVM().String_format("exist -- %s.%s", []LnsAny{classTypeInfo.FP.Get_rawTxt(), getterName}))
        } else { 
            var mutable bool
            mutable = memberNode.FP.Get_getterMutable()
            var getterMemberType *Ast_TypeInfo
            getterMemberType = memberNode.FP.Get_getterRetType()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(getterMemberType)) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(mutable)) ).(bool)){
                getterMemberType = self.FP.CreateModifier(getterMemberType, Ast_MutMode__IMut)
                
            }
            var retTypeInfo *Ast_TypeInfo
            retTypeInfo = self.processInfo.FP.CreateFunc(false, false, self.FP.pushScope(false, nil, nil), typeKind, classTypeInfo, false, false, memberNode.FP.Get_staticFlag(), accessMode, getterName, nil, NewLnsList([]LnsAny{}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(getterMemberType)}), nil)
            self.FP.popScope()
            classScope.FP.AddMethod(self.processInfo, memberName.Pos, retTypeInfo, accessMode, memberNode.FP.Get_staticFlag(), false)
            methodNameSet.Add(getterName)
        }
    }
    var setterName string
    setterName = "set_" + memberName.Txt
    accessMode = memberNode.FP.Get_setterMode()
    
    if memberNode.FP.Get_setterMode() != Ast_AccessMode__None{
        if Lns_isCondTrue( classScope.FP.GetTypeInfoChild(setterName)){
            self.FP.addErrMess(memberName.Pos, Lns_getVM().String_format("exist -- %s.%s", []LnsAny{classTypeInfo.FP.Get_rawTxt(), setterName}))
        } else { 
            var mutable bool
            if memberNode.FP.Get_symbolInfo().FP.Get_mutMode() != Ast_MutMode__AllMut{
                mutable = true
                
            } else { 
                mutable = false
                
            }
            classScope.FP.AddMethod(self.processInfo, memberName.Pos, self.processInfo.FP.CreateFunc(false, false, self.FP.pushScope(false, nil, nil), typeKind, classTypeInfo, false, false, memberNode.FP.Get_staticFlag(), accessMode, setterName, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(memberType)}), nil, mutable), accessMode, memberNode.FP.Get_staticFlag(), true)
            self.FP.popScope()
            methodNameSet.Add(setterName)
        }
    }
}

// 1958: decl @lune.@base.@TransUnit.TransUnit.analyzeClassBody
func (self *TransUnit_TransUnit) analyzeClassBody(hasProto bool,classAccessMode LnsInt,firstToken *Types_Token,mode LnsInt,gluePrefix LnsAny,classTypeInfo *Ast_TypeInfo,name *Types_Token,moduleLang LnsAny,moduleName LnsAny,lazyLoad LnsInt,nextToken *Types_Token)(*Nodes_DeclClassNode, *Types_Token, *LnsSet) {
    var memberName2Node *LnsMap
    memberName2Node = NewLnsMap( map[LnsAny]LnsAny{})
    var allStmtList *LnsList
    allStmtList = NewLnsList([]LnsAny{})
    var declStmtList *LnsList
    declStmtList = NewLnsList([]LnsAny{})
    var fieldList *LnsList
    fieldList = NewLnsList([]LnsAny{})
    var memberList *LnsList
    memberList = NewLnsList([]LnsAny{})
    var methodNameSet *LnsSet
    methodNameSet = NewLnsSet([]LnsAny{})
    var initBlockInfo *Nodes_ClassInitBlockInfo
    initBlockInfo = NewNodes_ClassInitBlockInfo(nil)
    var advertiseList *LnsList
    advertiseList = NewLnsList([]LnsAny{})
    var trustList *LnsList
    trustList = NewLnsList([]LnsAny{})
    var uninitMemberList *LnsList
    uninitMemberList = NewLnsList([]LnsAny{})
    var node *Nodes_DeclClassNode
    node = Nodes_DeclClassNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo)}), classAccessMode, name, hasProto, gluePrefix, moduleName, moduleLang, lazyLoad, false, allStmtList, declStmtList, fieldList, memberList, self.scope, initBlockInfo, advertiseList, trustList, uninitMemberList, NewLnsSet([]LnsAny{}))
    self.typeInfo2ClassNode.Set(classTypeInfo,node)
    var alreadyCtorFlag bool
    alreadyCtorFlag = false
    var hasInitBlock bool
    hasInitBlock = false
    var hasStaticMember bool
    hasStaticMember = false
    var classScope *Ast_Scope
    classScope = self.scope
    var processLet func(token *Types_Token,staticFlag bool,accessMode LnsInt,alreadyFlag bool)
    processLet = func(token *Types_Token,staticFlag bool,accessMode LnsInt,alreadyFlag bool) {
        if staticFlag{
            hasStaticMember = true
            
        }
        if mode == TransUnit_DeclClassMode__Interface{
            self.FP.addErrMess(token.Pos, "interface can not have member")
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(staticFlag)) &&
            Lns_GetEnv().SetStackVal( alreadyFlag) ).(bool)){
            self.FP.addErrMess(token.Pos, "member can't declare after '__init' method.")
        } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( staticFlag) &&
            Lns_GetEnv().SetStackVal( hasInitBlock) ).(bool)){
            self.FP.addErrMess(token.Pos, "static member can't declare after '__init' block.")
        }
        var memberNode *Nodes_DeclMemberNode
        memberNode = self.FP.analyzeDeclMember(classTypeInfo, accessMode, staticFlag, token)
        allStmtList.Insert(Nodes_DeclMemberNode2Stem(memberNode))
        fieldList.Insert(Nodes_DeclMemberNode2Stem(memberNode))
        memberList.Insert(Nodes_DeclMemberNode2Stem(memberNode))
        memberName2Node.Set(memberNode.FP.Get_name().Txt,memberNode)
        self.FP.addAccessor(memberNode, methodNameSet, classScope, classTypeInfo)
    }
    var checkInitializeMember func(staticFlag bool,pos LnsAny)
    checkInitializeMember = func(staticFlag bool,pos LnsAny) {
        for _memberName, _memberNode := range( memberName2Node.Items ) {
            memberName := _memberName.(string)
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if memberNode.FP.Get_staticFlag() == staticFlag{
                var symbolInfo *Ast_SymbolInfo
                symbolInfo = Lns_unwrap( self.scope.FP.GetSymbolInfoChild(memberName)).(*Ast_SymbolInfo)
                var typeInfo *Ast_TypeInfo
                typeInfo = symbolInfo.FP.Get_typeInfo()
                if Lns_op_not(symbolInfo.FP.Get_hasValueFlag()){
                    var msg string
                    if staticFlag{
                        msg = Lns_getVM().String_format("Set member -- %s", []LnsAny{memberName})
                        
                    } else { 
                        msg = Lns_getVM().String_format("Set member -- %s.%s", []LnsAny{name.Txt, memberName})
                        
                    }
                    if Lns_op_not(typeInfo.FP.Get_nilable()){
                        self.FP.addErrMess(Lns_unwrapDefault( pos, memberNode.FP.Get_pos()).(*Types_Position), msg)
                    } else { 
                        uninitMemberList.Insert(Ast_SymbolInfo2Stem(symbolInfo))
                        self.FP.addWarnMess(Lns_unwrapDefault( pos, memberNode.FP.Get_pos()).(*Types_Position), msg)
                    }
                }
            }
        }
    }
    var processFn func(token *Types_Token,staticFlag bool,accessMode LnsInt,abstractFlag bool,overrideFlag bool)
    processFn = func(token *Types_Token,staticFlag bool,accessMode LnsInt,abstractFlag bool,overrideFlag bool) {
        var nameToken *Types_Token
        nameToken = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
        var declFuncMode LnsInt
        declFuncMode = TransUnit_DeclFuncMode__Class
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( mode == TransUnit_DeclClassMode__Module) ||
            Lns_GetEnv().SetStackVal( mode == TransUnit_DeclClassMode__LazyModule) ).(bool){
            if Lns_isCondTrue( gluePrefix){
                declFuncMode = TransUnit_DeclFuncMode__Glue
                
            } else { 
                declFuncMode = TransUnit_DeclFuncMode__Module
                
            }
        }
        if nameToken.Txt == "__init"{
            for _, _symbolInfo := range( self.scope.FP.Get_symbol2SymbolInfoMap().Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if Lns_op_not(symbolInfo.FP.Get_staticFlag()){
                    symbolInfo.FP.ClearValue()
                }
            }
        }
        var methodNode *Nodes_Node
        methodNode = self.FP.analyzeDeclMethod(classTypeInfo, declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, token, nameToken)
        allStmtList.Insert(Nodes_Node2Stem(methodNode))
        fieldList.Insert(Nodes_Node2Stem(methodNode))
        methodNameSet.Add(nameToken.Txt)
        if nameToken.Txt == "__init"{
            alreadyCtorFlag = true
            
            checkInitializeMember(false, methodNode.FP.Get_pos())
        }
    }
    var processInitBlock func(token *Types_Token)
    processInitBlock = func(token *Types_Token) {
        if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(classTypeInfo.FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_parent()})) != self.moduleScope{
            self.FP.addErrMess(token.Pos, "The '__init' block only support at the top level classes.")
        }
        if mode != TransUnit_DeclClassMode__Class{
            self.FP.Error(Lns_getVM().String_format("%s can not have __init block.", []LnsAny{mode}))
        }
        hasInitBlock = true
        
        for _, _symbolInfo := range( self.scope.FP.Get_symbol2SymbolInfoMap().Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if symbolInfo.FP.Get_staticFlag(){
                symbolInfo.FP.ClearValue()
            }
        }
        var parentScope *Ast_Scope
        parentScope = self.scope
        var initBlockScope *Ast_Scope
        initBlockScope = self.FP.pushScope(false, nil, nil)
        self.FP.prepareTentativeSymbol(initBlockScope, false, nil)
        var ininame string
        ininame = "___init"
        var funcTypeInfo *Ast_TypeInfo
        funcTypeInfo = self.processInfo.FP.CreateFunc(false, false, initBlockScope, Ast_TypeInfoKind__Func, classTypeInfo, false, false, true, Ast_AccessMode__Pri, ininame, nil, nil, nil, false)
        var funcSym LnsAny
        var shadowing LnsAny
        funcSym,shadowing = parentScope.FP.AddFunc(self.processInfo, token.Pos, funcTypeInfo, Ast_AccessMode__Pri, true, true)
        var block *Nodes_BlockNode
        block = self.FP.analyzeFuncBlock(TransUnit_AnalyzingState__InitBlock, token, classTypeInfo, funcTypeInfo, ininame, initBlockScope, funcTypeInfo.FP.Get_retTypeInfoList())
        var info *Nodes_DeclFuncInfo
        info = NewNodes_DeclFuncInfo(Nodes_FuncKind__InitBlock, classTypeInfo, node, token, Lns_unwrap( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( funcSym) ||
            Lns_GetEnv().SetStackVal( shadowing) ))).(*Ast_SymbolInfo), NewLnsList([]LnsAny{}), true, Ast_AccessMode__Pri, block, NewLnsList([]LnsAny{}), false, false)
        var initBlockNode *Nodes_DeclMethodNode
        initBlockNode = Nodes_DeclMethodNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(funcTypeInfo)}), info)
        initBlockInfo.FP.Set_func(initBlockNode)
        allStmtList.Insert(Nodes_DeclMethodNode2Stem(initBlockNode))
        self.FP.popScope()
        self.FP.finishTentativeSymbol(false)
    }
    var processAdvertise func()
    processAdvertise = func() {
        var memberToken *Types_Token
        memberToken = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
        nextToken = self.FP.getToken(nil)
        
        var prefix string
        prefix = ""
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( nextToken.Txt != ";") &&
            Lns_GetEnv().SetStackVal( nextToken.Txt != "{") ).(bool)){
            prefix = nextToken.Txt
            
            nextToken = self.FP.getToken(nil)
            
        }
        self.FP.checkToken(nextToken, ";")
        var memberNode *Nodes_DeclMemberNode
        
        {
            _memberNode := memberName2Node.Items[memberToken.Txt]
            if _memberNode == nil{
                self.FP.Error(Lns_getVM().String_format("not found member -- %s", []LnsAny{memberToken.Txt}))
            } else {
                memberNode = _memberNode.(*Nodes_DeclMemberNode)
            }
        }
        var advInfo *Nodes_AdvertiseInfo
        advInfo = NewNodes_AdvertiseInfo(memberNode, prefix, memberToken.Pos)
        advertiseList.Insert(Nodes_AdvertiseInfo2Stem(advInfo))
        allStmtList.Insert(Nodes_DeclAdvertiseNode2Stem(Nodes_DeclAdvertiseNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), advInfo)))
        self.advertisedTypeSet.Add(Ast_TypeInfo2Stem(memberNode.FP.Get_expType().FP.Get_srcTypeInfo().FP.Get_genSrcTypeInfo()))
    }
    var processEnum func(token *Types_Token,accessMode LnsInt)
    processEnum = func(token *Types_Token,accessMode LnsInt) {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( accessMode != Ast_AccessMode__Pri) &&
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( classAccessMode == Ast_AccessMode__Pri) ||
                Lns_GetEnv().SetStackVal( classAccessMode == Ast_AccessMode__Local) ).(bool))) ).(bool)){
            self.FP.addErrMess(token.Pos, Lns_getVM().String_format("unmatch access mode, class('%s') and enum('%s')", []LnsAny{Ast_AccessMode_getTxt( classAccessMode), Ast_AccessMode_getTxt( accessMode)}))
        }
        var enumNode *Nodes_DeclEnumNode
        enumNode = self.FP.analyzeDeclEnum(accessMode, token)
        allStmtList.Insert(Nodes_DeclEnumNode2Stem(enumNode))
        declStmtList.Insert(Nodes_DeclEnumNode2Stem(enumNode))
    }
    var processLuneControl func()
    processLuneControl = func() {
        nextToken = self.FP.getToken(nil)
        
        var pragma LnsAny
        if _switch43845 := nextToken.Txt; _switch43845 == "default__init" {
            pragma = LuneControl_Pragma__default__init_Obj
            
            alreadyCtorFlag = true
            
            self.FP.addDefaultConstructor(nextToken.Pos, classTypeInfo, self.scope, memberList, methodNameSet, false)
        } else if _switch43845 == "default__init_old" {
            pragma = LuneControl_Pragma__default__init_old_Obj
            
            alreadyCtorFlag = true
            
            self.FP.addDefaultConstructor(nextToken.Pos, classTypeInfo, self.scope, memberList, methodNameSet, true)
            node.FP.SetHasOldCtor()
        } else {
            self.FP.Error(Lns_getVM().String_format("unknown option -- %s", []LnsAny{nextToken.Txt}))
        }
        self.FP.checkNextToken(";")
        var ctrlNode *Nodes_LuneControlNode
        ctrlNode = Nodes_LuneControlNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), pragma)
        self.helperInfo.PragmaSet.Add(pragma)
        allStmtList.Insert(Nodes_LuneControlNode2Stem(ctrlNode))
    }
    var processClassFields func(inMacro bool)
    processClassFields = func(inMacro bool) {
        for  {
            var token *Types_Token
            token = self.FP.getToken(inMacro)
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( token.Kind == Types_TokenKind__Eof) ||
                Lns_GetEnv().SetStackVal( token.Txt == "}") ).(bool){
                break
            }
            var accessMode LnsInt
            
            {
                _accessMode := Ast_txt2AccessMode(token.Txt)
                if _accessMode == nil{
                    accessMode = Ast_AccessMode__Pri
                    
                } else {
                    accessMode = _accessMode.(LnsInt)
                    token = self.FP.getToken(nil)
                    
                }
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( mode == TransUnit_DeclClassMode__Interface) &&
                Lns_GetEnv().SetStackVal( accessMode != Ast_AccessMode__Pub) ).(bool)){
                self.FP.addErrMess(token.Pos, "interface's fields must be 'pub'.")
            }
            var staticFlag bool
            staticFlag = false
            if token.Txt == "static"{
                staticFlag = true
                
                token = self.FP.getToken(nil)
                
            }
            var overrideFlag bool
            overrideFlag = false
            if token.Txt == "override"{
                overrideFlag = true
                
                token = self.FP.getToken(nil)
                
            }
            var abstractFlag bool
            abstractFlag = false
            if token.Txt == "abstract"{
                abstractFlag = true
                
                token = self.FP.getToken(nil)
                
            } else if mode == TransUnit_DeclClassMode__Interface{
                abstractFlag = true
                
            }
            if token.Txt == "let"{
                processLet(token, staticFlag, accessMode, alreadyCtorFlag)
            } else if token.Txt == "fn"{
                processFn(token, staticFlag, accessMode, abstractFlag, overrideFlag)
            } else if token.Txt == "__init"{
                processInitBlock(token)
            } else if token.Txt == "advertise"{
                processAdvertise()
            } else if token.Txt == ";"{
            } else if token.Txt == "enum"{
                processEnum(token, accessMode)
            } else if token.Txt == "_lune_control"{
                processLuneControl()
            } else { 
                {
                    _symbolInfo := self.scope.FP.GetSymbolInfo(token.Txt, self.scope, false, self.scopeAccess)
                    if _symbolInfo != nil {
                        symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                        if symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mac{
                            self.FP.checkNextToken("(")
                            var argList LnsAny
                            _,argList = self.FP.prepareExpCall(token.Pos, symbolInfo.FP.Get_typeInfo(), NewLnsList([]LnsAny{}), Ast_headTypeInfo)
                            self.FP.evalMacroOp(token, symbolInfo.FP.Get_typeInfo(), argList, Macro_EvalMacroCallback(func() {
                                processClassFields(true)
                            }))
                            self.FP.checkNextToken(";")
                        } else { 
                            self.FP.Error("illegal field")
                        }
                    } else {
                        self.FP.Error("illegal field")
                    }
                }
            }
        }
    }
    processClassFields(false)
    if _switch44379 := mode; _switch44379 == TransUnit_DeclClassMode__Module || _switch44379 == TransUnit_DeclClassMode__LazyModule {
    } else {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( hasStaticMember) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(hasInitBlock)) ).(bool)){
            self.FP.addErrMess(node.FP.Get_pos(), Lns_getVM().String_format("This class (%s) need __init block for initialize static members.", []LnsAny{name.Txt}))
        }
        checkInitializeMember(true, nil)
    }
    return node, nextToken, methodNameSet
}

// 2326: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclClass
func (self *TransUnit_TransUnit) analyzeDeclClass(classAbstructFlag bool,classAccessMode LnsInt,firstToken *Types_Token,mode LnsInt) *Nodes_DeclClassNode {
    if mode == TransUnit_DeclClassMode__Module{
        if self.FP.getToken(nil).Txt == "."{
            if _switch44457 := self.FP.getToken(nil).Txt; _switch44457 == "l" {
                mode = TransUnit_DeclClassMode__LazyModule
                
            } else if _switch44457 == "d" {
                mode = TransUnit_DeclClassMode__Module
                
            }
        } else { 
            self.FP.Pushback()
            if self.ctrl_info.DefaultLazy{
                mode = TransUnit_DeclClassMode__LazyModule
                
            }
        }
    }
    if mode == TransUnit_DeclClassMode__LazyModule{
        self.helperInfo.UseLazyRequire = true
        
    }
    if _switch44559 := mode; _switch44559 == TransUnit_DeclClassMode__Module || _switch44559 == TransUnit_DeclClassMode__LazyModule {
    } else {
        if _switch44557 := self.FP.getCurrentNamespaceTypeInfo().FP.Get_kind(); _switch44557 == Ast_TypeInfoKind__IF || _switch44557 == Ast_TypeInfoKind__Class || _switch44557 == Ast_TypeInfoKind__Module {
        } else if _switch44557 == Ast_TypeInfoKind__Func || _switch44557 == Ast_TypeInfoKind__Method {
            if _switch44543 := classAccessMode; _switch44543 == Ast_AccessMode__Pub || _switch44543 == Ast_AccessMode__Global {
                self.FP.addErrMess(firstToken.Pos, "Class can't declare on here.")
            }
        } else {
            self.FP.addErrMess(firstToken.Pos, "Class can't declare on here.")
        }
    }
    var name *Types_Token
    name = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
    var altTypeList *LnsList
    altTypeList = NewLnsList([]LnsAny{})
    {
        var nextToken *Types_Token
        nextToken = self.FP.getToken(nil)
        if nextToken.Txt == "<"{
            nextToken, altTypeList = self.FP.analyzeDeclAlternateType(true, nextToken, classAccessMode)
            
        }
        self.FP.PushbackToken(nextToken)
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( altTypeList.Len() > 0) &&
            Lns_GetEnv().SetStackVal( mode != TransUnit_DeclClassMode__Class) ).(bool)){
            self.FP.addErrMess(name.Pos, Lns_getVM().String_format("Only class can use the generics. -- %s ", []LnsAny{name.Txt}))
        }
    }
    if classAccessMode == Ast_AccessMode__Local{
        classAccessMode = Ast_AccessMode__Pri
        
    }
    var moduleName LnsAny
    moduleName = nil
    var gluePrefix LnsAny
    gluePrefix = nil
    var moduleLang LnsAny
    moduleLang = nil
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_DeclClassMode__Module) ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_DeclClassMode__LazyModule) ).(bool){
        self.FP.checkNextToken("require")
        moduleName = self.FP.getToken(nil)
        
        var nextToken *Types_Token
        nextToken = self.FP.getToken(nil)
        if nextToken.Txt == "of"{
            var langToken *Types_Token
            langToken = self.FP.getToken(nil)
            if langToken.Kind != Types_TokenKind__Str{
                self.FP.Error(Lns_getVM().String_format("it's not a string -- %s", []LnsAny{langToken.Txt}))
            }
            if langToken.FP.GetExcludedDelimitTxt() != ""{
                for _, _langId := range( Types_Lang_get__allList().Items ) {
                    langId := _langId.(LnsInt)
                    {
                        __exp := Types_Lang__from(LnsInt(langId))
                        if __exp != nil {
                            _exp := __exp.(LnsInt)
                            var ldName string
                            ldName = TransUnit_convExp44812(Lns_2DDD(Lns_getVM().String_gsub(Types_Lang_getTxt( _exp),".*%.", "")))
                            if ldName == langToken.Txt{
                                moduleLang = _exp
                                
                                break
                            }
                        }
                    }
                }
                if moduleLang == nil{
                    self.FP.ErrorAt(langToken.Pos, Lns_getVM().String_format("invalid lang -- %s", []LnsAny{langToken.Txt}))
                }
            } else { 
                moduleLang = Types_Lang__Same
                
            }
            nextToken = self.FP.getToken(nil)
            
        }
        if nextToken.Txt == "glue"{
            gluePrefix = self.FP.getToken(nil).FP.GetExcludedDelimitTxt()
            
        } else { 
            self.FP.Pushback()
        }
    }
    var existSymbolInfo LnsAny
    existSymbolInfo = self.scope.FP.GetSymbolTypeInfo(name.Txt, self.scope, self.scope, self.scopeAccess)
    var nextToken *Types_Token
    var classTypeInfo *Ast_TypeInfo
    nextToken,classTypeInfo = self.FP.analyzePushClass(mode, classAbstructFlag, firstToken, name, true, moduleName, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( moduleLang) ||
        Lns_GetEnv().SetStackVal( Types_Lang__Same) ).(LnsInt), classAccessMode, altTypeList)
    var hasProto bool
    if Lns_isCondTrue( self.protoClassMap.Items[classTypeInfo]){
        self.protoClassMap.Set(classTypeInfo,nil)
        hasProto = true
        
    } else { 
        hasProto = false
        
        if Lns_isCondTrue( existSymbolInfo){
            self.FP.addErrMess(name.Pos, Lns_getVM().String_format("already declare symbol -- %s", []LnsAny{name.Txt}))
        }
    }
    var classScope *Ast_Scope
    classScope = self.scope
    self.FP.checkToken(nextToken, "{")
    var mapType *Ast_TypeInfo
    mapType = self.processInfo.FP.CreateMap(Ast_AccessMode__Pub, classTypeInfo, Ast_builtinTypeString, self.FP.CreateModifier(Ast_builtinTypeStem, Ast_MutMode__IMut), Ast_MutMode__IMut)
    if classTypeInfo.FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, nil){
        self.helperInfo.HasMappingClassDef = true
        
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( classTypeInfo.FP.Get_baseTypeInfo() != Ast_headTypeInfo) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(classTypeInfo.FP.Get_baseTypeInfo().FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, nil))) ).(bool)){
            self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("must extend Mapping at %s", []LnsAny{classTypeInfo.FP.Get_baseTypeInfo().FP.GetTxt(nil, nil, nil)}))
        }
        var toMapFuncTypeInfo *Ast_TypeInfo
        toMapFuncTypeInfo = self.processInfo.FP.CreateFunc(false, false, nil, Ast_TypeInfoKind__Method, classTypeInfo, true, false, false, Ast_AccessMode__Pub, "_toMap", nil, NewLnsList([]LnsAny{}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(mapType)}), false)
        classScope.FP.AddMethod(self.processInfo, nil, toMapFuncTypeInfo, Ast_AccessMode__Pub, false, false)
    }
    var lazyLoad LnsInt
    if _switch45243 := mode; _switch45243 == TransUnit_DeclClassMode__LazyModule {
        lazyLoad = Nodes_LazyLoad__On
        
    } else if _switch45243 == TransUnit_DeclClassMode__Module || _switch45243 == TransUnit_DeclClassMode__Class || _switch45243 == TransUnit_DeclClassMode__Interface {
        lazyLoad = Nodes_LazyLoad__Off
        
    }
    var node *Nodes_DeclClassNode
    var methodNameSet *LnsSet
    node,_,methodNameSet = self.FP.analyzeClassBody(hasProto, classAccessMode, firstToken, mode, gluePrefix, classTypeInfo, name, moduleLang, moduleName, lazyLoad, nextToken)
    var ctorAccessMode LnsInt
    ctorAccessMode = Ast_AccessMode__Pub
    {
        _ctorTypeInfo := classScope.FP.GetTypeInfoChild("__init")
        if _ctorTypeInfo != nil {
            ctorTypeInfo := _ctorTypeInfo.(*Ast_TypeInfo)
            ctorAccessMode = ctorTypeInfo.FP.Get_accessMode()
            
        } else {
            self.FP.addDefaultConstructor(firstToken.Pos, classTypeInfo, classScope, node.FP.Get_memberList(), methodNameSet, false)
        }
    }
    for _, _advertiseInfo := range( node.FP.Get_advertiseList().Items ) {
        advertiseInfo := _advertiseInfo.(Nodes_AdvertiseInfoDownCast).ToNodes_AdvertiseInfo()
        var memberType *Ast_TypeInfo
        memberType = advertiseInfo.FP.Get_member().FP.Get_expType()
        if _switch45461 := memberType.FP.Get_kind(); _switch45461 == Ast_TypeInfoKind__Class || _switch45461 == Ast_TypeInfoKind__IF {
            for _, _mtdName := range( Ast_getAllMethodName(memberType, Ast_MethodKind__Object).FP.Get_list().Items ) {
                mtdName := _mtdName.(string)
                var scope *Ast_Scope
                scope = Lns_unwrap( memberType.FP.Get_scope()).(*Ast_Scope)
                var child *Ast_TypeInfo
                child = Lns_unwrap( scope.FP.GetTypeInfoField(mtdName, true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
                if child.FP.Get_accessMode() != Ast_AccessMode__Pri{
                    var childName string
                    childName = advertiseInfo.FP.Get_prefix() + child.FP.GetTxt(nil, nil, nil)
                    if Lns_op_not(methodNameSet.Has(childName)){
                        var impMtdType *Ast_TypeInfo
                        impMtdType = self.processInfo.FP.CreateAdvertiseMethodFrom(classTypeInfo, child)
                        classScope.FP.AddMethod(self.processInfo, advertiseInfo.FP.Get_pos(), impMtdType, child.FP.Get_accessMode(), child.FP.Get_staticFlag(), false)
                    }
                }
            }
        } else {
            self.FP.Error(Lns_getVM().String_format("advertise member type is illegal -- %s", []LnsAny{advertiseInfo.FP.Get_member().FP.Get_name()}))
        }
    }
    if classTypeInfo.FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, nil){
        var checkedTypeMap *LnsMap
        checkedTypeMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _, _memberNode := range( node.FP.Get_memberList().Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            var memberType *Ast_TypeInfo
            memberType = memberNode.FP.Get_expType()
            if Lns_op_not(Ast_NormalTypeInfo_isAvailableMapping(self.processInfo, memberType, checkedTypeMap)){
                self.FP.addErrMess(memberNode.FP.Get_pos(), Lns_getVM().String_format("member type is not Mapping -- %s", []LnsAny{memberType.FP.GetTxt(nil, nil, nil)}))
            } else if memberType.FP.Get_kind() == Ast_TypeInfoKind__IF{
                self.FP.addErrMess(memberNode.FP.Get_pos(), Lns_getVM().String_format("Mapping class has not the interface type member. -- %s", []LnsAny{memberNode.FP.Get_name().Txt}))
            } else if memberType.FP.Get_abstractFlag(){
                self.FP.addErrMess(memberNode.FP.Get_pos(), Lns_getVM().String_format("Mapping class has not the abstract class member. -- %s", []LnsAny{memberNode.FP.Get_name().Txt}))
            }
        }
        var fromMapFuncTypeInfo *Ast_TypeInfo
        fromMapFuncTypeInfo = self.processInfo.FP.CreateFunc(false, false, nil, Ast_TypeInfoKind__Func, classTypeInfo, true, false, true, Ast_AccessMode__Pub, "_fromMap", nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(mapType.FP.Get_nilableTypeInfo())}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo.FP.Get_nilableTypeInfo()), Ast_TypeInfo2Stem(Ast_builtinTypeString.FP.Get_nilableTypeInfo())}), true)
        classScope.FP.AddMethod(self.processInfo, nil, fromMapFuncTypeInfo, ctorAccessMode, true, false)
        var fromStemFuncTypeInfo *Ast_TypeInfo
        fromStemFuncTypeInfo = self.processInfo.FP.CreateFunc(false, false, nil, Ast_TypeInfoKind__Func, classTypeInfo, true, false, true, Ast_AccessMode__Pub, "_fromStem", nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeStem_)}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo.FP.Get_nilableTypeInfo()), Ast_TypeInfo2Stem(Ast_builtinTypeString.FP.Get_nilableTypeInfo())}), true)
        classScope.FP.AddMethod(self.processInfo, nil, fromStemFuncTypeInfo, ctorAccessMode, true, false)
    }
    if classTypeInfo.FP.IsInheritFrom(self.processInfo, Ast_builtinTypeAsyncItem, nil){
        if Lns_op_not(classTypeInfo.FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, nil)){
            self.FP.addErrMess(firstToken.Pos, "__AsyncItem implemented class must inherit Mapping.")
        }
        var pipeType *Ast_GenericTypeInfo
        pipeType = self.processInfo.FP.CreateGeneric(TransUnit_builtinFunc.G__pipe_, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo)}), self.moduleType)
        var createPipeFuncTypeInfo *Ast_TypeInfo
        createPipeFuncTypeInfo = self.processInfo.FP.CreateFunc(false, false, nil, Ast_TypeInfoKind__Func, classTypeInfo, true, false, true, Ast_AccessMode__Pub, "_createPipe", nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(pipeType.FP.Get_nilableTypeInfo())}), true)
        classScope.FP.AddMethod(self.processInfo, nil, createPipeFuncTypeInfo, Ast_AccessMode__Pub, true, false)
    }
    self.FP.popClass()
    return node
}

// 2617: decl @lune.@base.@TransUnit.TransUnit.addMethod
func (self *TransUnit_TransUnit) addMethod(classTypeInfo *Ast_TypeInfo,methodNode *Nodes_Node,name string) {
    var classNodeInfo *Nodes_DeclClassNode
    classNodeInfo = Lns_unwrap( self.typeInfo2ClassNode.Items[classTypeInfo]).(*Nodes_DeclClassNode)
    classNodeInfo.FP.Get_outerMethodSet().Add(name)
    classNodeInfo.FP.Get_fieldList().Insert(Nodes_Node2Stem(methodNode))
}

// 2626: decl @lune.@base.@TransUnit.TransUnit.processAddFunc
func (self *TransUnit_TransUnit) processAddFunc(isFunc bool,parentScope *Ast_Scope,name *Types_Token,typeInfo *Ast_TypeInfo,alt2typeMap *LnsMap) *Ast_SymbolInfo {
    var accessMode LnsInt
    accessMode = typeInfo.FP.Get_accessMode()
    if accessMode == Ast_AccessMode__Global{
        parentScope = self.globalScope
        
    }
    var hasPrototype bool
    {
        _prottype := parentScope.FP.GetTypeInfoChild(typeInfo.FP.Get_rawTxt())
        if _prottype != nil {
            prottype := _prottype.(*Ast_TypeInfo)
            var argTypeList *LnsList
            argTypeList = typeInfo.FP.Get_argTypeInfoList()
            var retTypeInfoList *LnsList
            retTypeInfoList = typeInfo.FP.Get_retTypeInfoList()
            var matched bool
            matched = true
            {
                var matchFlag LnsInt
                var err string
                matchFlag,err = Ast_TypeInfo_checkMatchType(self.processInfo, prottype.FP.Get_argTypeInfoList(), argTypeList, false, nil, alt2typeMap)
                if matchFlag != Ast_MatchType__Match{
                    self.FP.addErrMess(name.Pos, "mismatch functype param: " + err)
                    matched = false
                    
                }
            }
            {
                var matchFlag LnsInt
                var err string
                matchFlag,err = Ast_TypeInfo_checkMatchType(self.processInfo, prottype.FP.Get_retTypeInfoList(), retTypeInfoList, false, nil, alt2typeMap)
                if matchFlag != Ast_MatchType__Match{
                    self.FP.addErrMess(name.Pos, "mismatch functype ret: " + err)
                    matched = false
                    
                }
            }
            {
                var matchFlag bool
                var err LnsAny
                matchFlag,err = typeInfo.FP.CanEvalWith(self.processInfo, prottype, Ast_CanEvalType__SetOp, alt2typeMap)
                if Lns_op_not(matchFlag){
                    if Lns_isCondTrue( err){
                        self.FP.addErrMess(name.Pos, Lns_getVM().String_format("mismatch functype -- %s", []LnsAny{err}))
                    } else { 
                        self.FP.addErrMess(name.Pos, Lns_getVM().String_format("mismatch functype -- %s / %s", []LnsAny{typeInfo.FP.Get_display_stirng(), prottype.FP.Get_display_stirng()}))
                    }
                    matched = false
                    
                }
            }
            if Lns_isCondTrue( self.protoFuncMap.Items[prottype]){
                hasPrototype = true
                
                self.protoFuncMap.Set(prottype,nil)
            } else { 
                hasPrototype = false
                
                if Lns_op_not(prottype.FP.Get_autoFlag()){
                    self.FP.addErrMess(name.Pos, Lns_getVM().String_format("multiple define -- %s", []LnsAny{name.Txt}))
                }
            }
            if matched{
                (Lns_unwrap( Ast_NormalTypeInfoDownCastF(prottype.FP)).(*Ast_NormalTypeInfo)).FP.SwitchScopeTo(Lns_unwrap( typeInfo.FP.Get_scope()).(*Ast_Scope))
                typeInfo = prottype
                
            }
        } else {
            hasPrototype = false
            
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Method) &&
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_accessMode() != Ast_AccessMode__Pri) ).(bool)){
        var classType *Ast_TypeInfo
        classType = typeInfo.FP.Get_parentInfo()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( self.advertisedTypeSet.Has(Ast_TypeInfo2Stem(classType))) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(hasPrototype)) ).(bool)){
            self.FP.addErrMess(name.Pos, Lns_getVM().String_format("This class(%s) is used by advertise. You must declare the prototype of this method.", []LnsAny{classType.FP.GetTxt(nil, nil, nil)}))
        }
    }
    var staticFlag bool
    staticFlag = typeInfo.FP.Get_staticFlag()
    var mutable bool
    mutable = Ast_TypeInfo_isMut(typeInfo)
    var funcSym LnsAny
    var shadowing LnsAny
    if isFunc{
        funcSym, shadowing = parentScope.FP.AddFunc(self.processInfo, name.Pos, typeInfo, accessMode, staticFlag, mutable)
        
        self.FP.errorShadowing(name.Pos, shadowing)
    } else { 
        funcSym, shadowing = parentScope.FP.AddMethod(self.processInfo, name.Pos, typeInfo, accessMode, staticFlag, mutable)
        
    }
    return Lns_unwrap( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( funcSym) ||
        Lns_GetEnv().SetStackVal( shadowing) ))).(*Ast_SymbolInfo)
}

// 2723: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclFunc
func (self *TransUnit_TransUnit) analyzeDeclFunc(declFuncMode LnsInt,abstractFlag bool,overrideFlag bool,accessMode LnsInt,staticFlag bool,classTypeInfo LnsAny,firstToken *Types_Token,name LnsAny) *Nodes_Node {
    var token *Types_Token
    token = self.FP.getToken(nil)
    {
        __exp := name
        if __exp != nil {
            _exp := __exp.(*Types_Token)
            if _exp.Txt != "__main"{
                name = self.FP.checkSymbol(_exp, TransUnit_SymbolMode__MustNot_)
                
            }
        } else {
            if token.Txt != "("{
                if token.Txt != "__main"{
                    name = self.FP.checkSymbol(token, TransUnit_SymbolMode__MustNot_)
                    
                } else { 
                    name = token
                    
                }
                token = self.FP.getToken(nil)
                
            }
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(name)) &&
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Ast_isPubToExternal(accessMode)) ||
            Lns_GetEnv().SetStackVal( abstractFlag) ||
            Lns_GetEnv().SetStackVal( overrideFlag) ||
            Lns_GetEnv().SetStackVal( staticFlag) ).(bool))) ).(bool)){
        self.FP.addErrMess(firstToken.Pos, "The anonymous function must be local.")
    }
    var needPopFlag bool
    needPopFlag = false
    if token.Txt == "."{
        needPopFlag = true
        
        if name != nil{
            name_9210 := name.(*Types_Token)
            var className string
            className = name_9210.Txt
            classTypeInfo = self.scope.FP.GetTypeInfoChild(className)
            
            if classTypeInfo != nil{
                classTypeInfo_9213 := classTypeInfo.(*Ast_TypeInfo)
                self.FP.pushClassScope(name_9210.Pos, classTypeInfo_9213)
            } else {
                self.FP.Error(Lns_getVM().String_format("not found class -- %s", []LnsAny{className}))
            }
        } else {
            self.FP.Error("can't use '.' for any function name")
        }
        name = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
        
        token = self.FP.getToken(nil)
        
        if accessMode == Ast_AccessMode__Local{
            accessMode = Ast_AccessMode__Pri
            
        }
    }
    var isCtorFlag bool
    isCtorFlag = false
    var kind LnsInt
    kind = Nodes_NodeKind_get_DeclConstr()
    var typeKind LnsInt
    typeKind = Ast_TypeInfoKind__Func
    if classTypeInfo != nil{
        if Lns_op_not(staticFlag){
            typeKind = Ast_TypeInfoKind__Method
            
        }
        if _switch46851 := (Lns_unwrap( name).(*Types_Token)).Txt; _switch46851 == "__init" {
            isCtorFlag = true
            
            kind = Nodes_NodeKind_get_DeclConstr()
            
            for _, _symbolInfo := range( self.scope.FP.Get_symbol2SymbolInfoMap().Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if Lns_op_not(symbolInfo.FP.Get_staticFlag()){
                    symbolInfo.FP.ClearValue()
                }
            }
        } else if _switch46851 == "__free" {
            kind = Nodes_NodeKind_get_DeclDestr()
            
            if Lns_op_not(self.targetLuaVer.FP.Get_canUseMetaGc()){
                self.FP.addErrMess(firstToken.Pos, "this lua version is not support __free.")
            }
        } else {
            kind = Nodes_NodeKind_get_DeclMethod()
            
        }
    } else {
        kind = Nodes_NodeKind_get_DeclFunc()
        
        if Lns_op_not(staticFlag){
            staticFlag = true
            
        }
    }
    var orgStaticFlag bool
    orgStaticFlag = staticFlag
    if declFuncMode == TransUnit_DeclFuncMode__Module{
        staticFlag = true
        
    }
    var parentScope *Ast_Scope
    parentScope = self.scope
    var funcBodyScope *Ast_Scope
    funcBodyScope = self.FP.pushScope(false, nil, nil)
    var altTypeList *LnsList
    altTypeList = NewLnsList([]LnsAny{})
    if token.Txt == "<"{
        token, altTypeList = self.FP.analyzeDeclAlternateType(false, token, accessMode)
        
        for _, _altType := range( altTypeList.Items ) {
            altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            funcBodyScope.FP.AddAlternate(self.processInfo, accessMode, altType.FP.Get_rawTxt(), firstToken.Pos, altType)
        }
    }
    self.FP.checkToken(token, "(")
    var parentPub bool
    if classTypeInfo != nil{
        classTypeInfo_9242 := classTypeInfo.(*Ast_TypeInfo)
        parentPub = Ast_isPubToExternal(classTypeInfo_9242.FP.Get_accessMode())
        
    } else {
        parentPub = Ast_isPubToExternal(accessMode)
        
    }
    var argList *LnsList
    argList = NewLnsList([]LnsAny{})
    token = self.FP.analyzeDeclArgList(accessMode, funcBodyScope, argList, parentPub)
    
    self.FP.checkToken(token, ")")
    token = self.FP.getToken(nil)
    
    var mutable bool
    mutable = false
    if token.Txt == "mut"{
        token = self.FP.getToken(nil)
        
        mutable = true
        
    }
    var pubToExtFlag bool
    pubToExtFlag = Ast_isPubToExternal(accessMode)
    var argTypeList *LnsList
    argTypeList = NewLnsList([]LnsAny{})
    for _, _argNode := range( argList.Items ) {
        argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
        argTypeList.Insert(Ast_TypeInfo2Stem(argNode.FP.Get_expType()))
    }
    var alt2typeMap *LnsMap
    alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false)
    if classTypeInfo != nil{
        classTypeInfo_9253 := classTypeInfo.(*Ast_TypeInfo)
        alt2typeMap = classTypeInfo_9253.FP.CreateAlt2typeMap(false)
        
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( kind == Nodes_NodeKind_get_DeclMethod()) ||
            Lns_GetEnv().SetStackVal( kind == Nodes_NodeKind_get_DeclConstr()) ||
            Lns_GetEnv().SetStackVal( kind == Nodes_NodeKind_get_DeclDestr()) ).(bool){
            var workClass *Ast_TypeInfo
            workClass = classTypeInfo_9253
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( kind == Nodes_NodeKind_get_DeclConstr()) ||
                Lns_GetEnv().SetStackVal( kind == Nodes_NodeKind_get_DeclDestr()) ).(bool){
                mutable = true
                
            }
            if Lns_op_not(Ast_isPubToExternal(workClass.FP.Get_accessMode())){
                pubToExtFlag = false
                
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(workClass)) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(mutable)) ).(bool)){
                workClass = self.FP.CreateModifier(workClass, Ast_MutMode__IMut)
                
            }
            if Lns_op_not(staticFlag){
                self.scope.FP.Add(self.processInfo, Ast_SymbolKind__Arg, false, true, "self", nil, workClass, Ast_AccessMode__Pri, false, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( mutable) &&
                    Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
                    Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, false)
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(workClass.FP.Get_abstractFlag())) &&
                Lns_GetEnv().SetStackVal( abstractFlag) ).(bool)){
                self.FP.addErrMess(firstToken.Pos, "no abstract class does not have abstract method")
            }
        }
    }
    var retTypeInfoList *LnsList
    retTypeInfoList = NewLnsList([]LnsAny{})
    retTypeInfoList, token = self.FP.analyzeRetTypeList(pubToExtFlag, accessMode, token, parentPub)
    
    var namespaceInfo *Ast_TypeInfo
    namespaceInfo = self.FP.getCurrentNamespaceTypeInfo()
    var funcName string
    funcName = ""
    if name != nil{
        name_9265 := name.(*Types_Token)
        funcName = name_9265.Txt
        
        if kind == Nodes_NodeKind_get_DeclFunc(){
            if _switch47438 := accessMode; _switch47438 == Ast_AccessMode__Pub || _switch47438 == Ast_AccessMode__Global {
                if parentScope != self.moduleScope{
                    self.FP.addErrMess(firstToken.Pos, "'global' or 'pub' function must exist top scope.")
                }
            }
        }
    }
    var typeInfo *Ast_TypeInfo
    var funcSym LnsAny
    {
        var workTypeInfo *Ast_TypeInfo
        workTypeInfo = self.processInfo.FP.CreateFunc(abstractFlag, false, funcBodyScope, typeKind, namespaceInfo, false, false, staticFlag, accessMode, funcName, altTypeList, argTypeList, retTypeInfoList, mutable)
        if name != nil{
            name_9274 := name.(*Types_Token)
            var workSym *Ast_SymbolInfo
            workSym = self.FP.processAddFunc(kind == Nodes_NodeKind_get_DeclFunc(), funcBodyScope.FP.Get_parent(), name_9274, workTypeInfo, alt2typeMap)
            typeInfo = workSym.FP.Get_typeInfo()
            
            funcSym = workSym
            
            if name_9274.Txt == "__main"{
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( typeInfo.FP.Get_argTypeInfoList().Len() != 1) ||
                    Lns_GetEnv().SetStackVal( typeInfo.FP.Get_argTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() != Ast_TypeInfoKind__List) ||
                    Lns_GetEnv().SetStackVal( typeInfo.FP.Get_argTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeString) ||
                    Lns_GetEnv().SetStackVal( typeInfo.FP.Get_retTypeInfoList().Len() != 1) ||
                    Lns_GetEnv().SetStackVal( typeInfo.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeInt) ).(bool){
                    var mess string
                    mess = Lns_getVM().String_format("'__main' function's type has to be __main( argList:List<str> ) : int -- %s", []LnsAny{typeInfo.FP.Get_display_stirng()})
                    self.FP.addErrMess(name_9274.Pos, mess)
                }
            }
        } else {
            typeInfo = workTypeInfo
            
            funcSym = nil
            
        }
    }
    if overrideFlag{
        if Lns_op_not(name){
            self.FP.addErrMess(firstToken.Pos, "can't override anonymous func")
        }
        if TransUnit_CantOverrideMethods.Has(funcName){
            self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("This method can't override -- %s", []LnsAny{funcName}))
        }
        {
            _overrideType := self.scope.FP.Get_parent().FP.GetTypeInfoField(funcName, false, funcBodyScope, self.scopeAccess)
            if _overrideType != nil {
                overrideType := _overrideType.(*Ast_TypeInfo)
                for _, _err := range( self.FP.checkOverrideMethod(overrideType, typeInfo).Items ) {
                    err := _err.(string)
                    self.FP.addErrMess(firstToken.Pos, err)
                }
            } else {
                self.FP.addErrMess(firstToken.Pos, "not found override -- " + funcName)
            }
        }
    } else { 
        if name != nil{
            name_9290 := name.(*Types_Token)
            if Lns_op_not(TransUnit_CantOverrideMethods.Has(name_9290.Txt)){
                if Lns_isCondTrue( self.scope.FP.Get_parent().FP.GetTypeInfoField(name_9290.Txt, false, funcBodyScope, Ast_ScopeAccess__Full)){
                    self.FP.addErrMess(firstToken.Pos, "mismatch override --" + funcName)
                } else { 
                    {
                        _ifFunc := self.scope.FP.Get_parent().FP.GetSymbolInfoIfField(name_9290.Txt, funcBodyScope, Ast_ScopeAccess__Full)
                        if _ifFunc != nil {
                            ifFunc := _ifFunc.(*Ast_SymbolInfo)
                            if Lns_op_not(Lns_car(ifFunc.FP.Get_typeInfo().FP.CanEvalWith(self.processInfo, typeInfo, Ast_CanEvalType__SetEq, alt2typeMap)).(bool)){
                                self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("mismatch method type -- %s", []LnsAny{funcName}))
                            }
                        }
                    }
                }
            }
        }
    }
    var node *Nodes_Node
    node = self.FP.createNoneNode(firstToken.Pos)
    var body LnsAny
    body = nil
    if token.Txt == ";"{
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( declFuncMode == TransUnit_DeclFuncMode__Module) ||
            Lns_GetEnv().SetStackVal( declFuncMode == TransUnit_DeclFuncMode__Glue) ).(bool){
        } else { 
            if Lns_op_not(abstractFlag){
                self.protoFuncMap.Set(typeInfo,firstToken.Pos)
            }
            if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(classTypeInfo) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_kind()})) == Ast_TypeInfoKind__IF{
            } else { 
                if kind == Nodes_NodeKind_get_DeclMethod(){
                    kind = Nodes_NodeKind_get_ProtoMethod()
                    
                }
            }
        }
    } else { 
        if abstractFlag{
            self.FP.addErrMess(token.Pos, "abstract method can't have body.")
        }
        self.FP.Pushback()
        var analyzingState LnsInt
        if isCtorFlag{
            analyzingState = TransUnit_AnalyzingState__Constructor
            
        } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( staticFlag) &&
            Lns_GetEnv().SetStackVal( classTypeInfo) )){
            analyzingState = TransUnit_AnalyzingState__ClassMethod
            
        } else { 
            analyzingState = TransUnit_AnalyzingState__Func
            
        }
        funcBodyScope.FP.AddLocalVar(self.processInfo, false, false, "__func__", firstToken.Pos, Ast_builtinTypeString, Ast_MutMode__IMut)
        var workBody *Nodes_BlockNode
        workBody = self.FP.analyzeFuncBlock(analyzingState, firstToken, classTypeInfo, typeInfo, funcName, funcBodyScope, typeInfo.FP.Get_retTypeInfoList())
        body = workBody
        
        if isCtorFlag{
            if classTypeInfo != nil{
                classTypeInfo_9315 := classTypeInfo.(*Ast_TypeInfo)
                if classTypeInfo_9315.FP.Get_baseTypeInfo() != Ast_headTypeInfo{
                    var needCall bool
                    needCall = true
                    for _, _stmt := range( workBody.FP.Get_stmtList().Items ) {
                        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
                        if _switch48094 := stmt.FP.Get_kind(); _switch48094 == Nodes_nodeKindEnum__ExpCallSuperCtor {
                            needCall = false
                            
                        } else if _switch48094 == Nodes_nodeKindEnum__BlankLine {
                        } else {
                            break
                        }
                    }
                    if needCall{
                        self.FP.addErrMess(workBody.FP.Get_pos(), "__init must call super() with first.")
                    }
                }
            }
        }
    }
    var createDeclFuncInfo func(funcKind LnsInt) *Nodes_DeclFuncInfo
    createDeclFuncInfo = func(funcKind LnsInt) *Nodes_DeclFuncInfo {
        var classDeclNode LnsAny
        if classTypeInfo != nil{
            classTypeInfo_9330 := classTypeInfo.(*Ast_TypeInfo)
            classDeclNode = self.typeInfo2ClassNode.Items[classTypeInfo_9330]
            
        } else {
            classDeclNode = nil
            
        }
        return NewNodes_DeclFuncInfo(funcKind, classTypeInfo, classDeclNode, name, funcSym, argList, orgStaticFlag, accessMode, body, retTypeInfoList, self.has__func__Symbol.Has(Ast_TypeInfo2Stem(typeInfo)), overrideFlag)
    }
    if _switch48446 := (kind); _switch48446 == Nodes_NodeKind_get_DeclConstr() {
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(Nodes_FuncKind__Ctor)
        node = &Nodes_DeclConstrNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
        
    } else if _switch48446 == Nodes_NodeKind_get_DeclDestr() {
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(Nodes_FuncKind__Dstr)
        node = &Nodes_DeclDestrNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
        
    } else if _switch48446 == Nodes_NodeKind_get_DeclMethod() {
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(Nodes_FuncKind__Mtd)
        node = &Nodes_DeclMethodNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
        
    } else if _switch48446 == Nodes_NodeKind_get_ProtoMethod() {
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(Nodes_FuncKind__Mtd)
        node = &Nodes_ProtoMethodNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
        
    } else if _switch48446 == Nodes_NodeKind_get_DeclFunc() {
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(Nodes_FuncKind__Func)
        node = &Nodes_DeclFuncNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
        
    } else {
        self.FP.Error(Lns_getVM().String_format("illegal kind -- %d", []LnsAny{kind}))
    }
    self.has__func__Symbol.Del(Ast_TypeInfo2Stem(typeInfo))
    self.FP.popScope()
    if needPopFlag{
        self.FP.addMethod(Lns_unwrap( classTypeInfo).(*Ast_TypeInfo), node, funcName)
        self.FP.popClass()
    }
    return node
}

// 3158: decl @lune.@base.@TransUnit.TransUnit.createExpListNode
func (self *TransUnit_TransUnit) createExpListNode(orgExpList *Nodes_ExpListNode,newExpList *LnsList) *Nodes_ExpListNode {
    var newExpTypeList *LnsList
    newExpTypeList = NewLnsList([]LnsAny{})
    for _, _expNode := range( newExpList.Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        newExpTypeList.Insert(Ast_TypeInfo2Stem(expNode.FP.Get_expType()))
    }
    if newExpList.GetAt(newExpList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_expTypeList().Len() > 1{
        self.FP.addErrMess(orgExpList.FP.Get_pos(), Lns_getVM().String_format("illegal exp -- %d", []LnsAny{newExpList.GetAt(newExpList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_expTypeList().Len()}))
    }
    {
        _mRetIndex := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(orgExpList.FP.Get_mRetExp()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_MRetExp).FP.Get_index()}))
        if _mRetIndex != nil {
            mRetIndex := _mRetIndex.(LnsInt)
            if mRetIndex > newExpList.Len(){
                self.FP.addErrMess(orgExpList.FP.Get_pos(), Lns_getVM().String_format("over index -- %d", []LnsAny{mRetIndex}))
            }
        }
    }
    return Nodes_ExpListNode_create(self.nodeManager, orgExpList.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), newExpTypeList, newExpList, orgExpList.FP.Get_mRetExp(), orgExpList.FP.Get_followOn())
}

// 3192: decl @lune.@base.@TransUnit.TransUnit.checkLiteralEmptyCollection
func (self *TransUnit_TransUnit) checkLiteralEmptyCollection(pos *Types_Position,symbolName string,expType *Ast_TypeInfo) {
    for _, _itemType := range( expType.FP.Get_itemTypeInfoList().Items ) {
        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if itemType == Ast_builtinTypeNone{
            self.FP.addErrMess(pos, Lns_getVM().String_format("must set the item type of Collection. -- %s:%s", []LnsAny{symbolName, expType.FP.Get_srcTypeInfo().FP.GetTxt(self.typeNameCtrl, nil, nil)}))
            break
        }
    }
}

// 3207: decl @lune.@base.@TransUnit.TransUnit.accessSymbol
func (self *TransUnit_TransUnit) accessSymbol(symbolInfo *Ast_SymbolInfo,canLeftExp bool) {
    if symbolInfo.FP.Get_kind() == Ast_SymbolKind__Fun{
        self.scope.FP.AccessSymbol(self.moduleScope, symbolInfo)
        symbolInfo.FP.Set_posForModToRef(symbolInfo.FP.Get_posForLatestMod())
        if self.scope.FP.IsClosureAccess(self.moduleScope, symbolInfo){
            self.closureFunList.Insert(TransUnit_ClosureFun2Stem(NewTransUnit_ClosureFun(symbolInfo, self.scope)))
        }
        {
            _scope := symbolInfo.FP.Get_typeInfo().FP.Get_scope()
            if _scope != nil {
                scope := _scope.(*Ast_Scope)
                for _, _symInfo := range( scope.FP.Get_closureSymList().Items ) {
                    symInfo := _symInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symInfo.FP.Set_posForModToRef(symInfo.FP.Get_posForLatestMod())
                }
            }
        }
    } else { 
        self.scope.FP.AccessSymbol(self.moduleScope, symbolInfo)
        if canLeftExp{
            self.accessSymbolSetQueue.FP.Add(symbolInfo)
        } else { 
            symbolInfo.FP.Set_posForModToRef(symbolInfo.FP.Get_posForLatestMod())
        }
    }
}

// 3245: decl @lune.@base.@TransUnit.TransUnit.analyzeInitExp
func (self *TransUnit_TransUnit) analyzeInitExp(firstPos *Types_Position,accessMode LnsInt,unwrapFlag bool,letVarList *LnsList,typeInfoList *LnsList)(*LnsList, *LnsList, *LnsList, LnsAny) {
    var expList LnsAny
    expList = nil
    var expectTypeList *LnsList
    expectTypeList = NewLnsList([]LnsAny{})
    for _, _varInfo := range( letVarList.Items ) {
        varInfo := _varInfo.(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
        expectTypeList.Insert(Ast_TypeInfo2Stem(Lns_unwrapDefault( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(varInfo.VarType) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_RefTypeNode).FP.Get_expType()})), Ast_builtinTypeNone).(*Ast_TypeInfo)))
    }
    expList = self.FP.analyzeExpList(false, false, false, nil, expectTypeList, nil)
    
    if Lns_op_not(expList){
        self.FP.Error("expList is nil")
    }
    var orgExpTypeList *LnsList
    orgExpTypeList = NewLnsList([]LnsAny{})
    if expList != nil{
        expList_9403 := expList.(*Nodes_ExpListNode)
        if unwrapFlag{
            var hasNilable bool
            hasNilable = false
            for _index, _ := range( letVarList.Items ) {
                index := _index + 1
                if expList_9403.FP.GetExpTypeAt(index).FP.Get_nilable(){
                    hasNilable = true
                    
                    break
                }
            }
            if Lns_op_not(hasNilable){
                self.FP.addWarnMess(firstPos, "has no nilable")
            }
        }
        var workList *Nodes_ExpListNode
        workList = expList_9403
        var updateExpList bool
        updateExpList = false
        var newExpList *LnsList
        newExpList = NewLnsList([]LnsAny{})
        for _index, _exp := range( workList.FP.Get_expList().Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            newExpList.Insert(Nodes_Node2Stem(exp))
            if Lns_op_not(exp.FP.CanBeRight(self.processInfo)){
                self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("this node(%d) can not be r-value. -- %s", []LnsAny{index, Nodes_getNodeKindName(exp.FP.Get_kind())}))
            }
        }
        var expTypeList *LnsList
        expTypeList = NewLnsList([]LnsAny{})
        for _index, _expType := range( workList.FP.Get_expTypeList().Items ) {
            index := _index + 1
            expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( index == workList.FP.Get_expTypeList().Len()) &&
                Lns_GetEnv().SetStackVal( expType.FP.Get_kind() == Ast_TypeInfoKind__DDD) ).(bool)){
                var dddItemType *Ast_TypeInfo
                dddItemType = Ast_builtinTypeStem_
                if expType.FP.Get_itemTypeInfoList().Len() > 0{
                    dddItemType = expType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo()
                    
                }
                {
                    var _from49265 LnsInt = index
                    var _to49265 LnsInt = letVarList.Len()
                    for _work49265 := _from49265; _work49265 <= _to49265; _work49265++ {
                        subIndex := _work49265
                        var argType *Ast_TypeInfo
                        argType = typeInfoList.GetAt(subIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        var checkType *Ast_TypeInfo
                        checkType = dddItemType
                        if unwrapFlag{
                            checkType = dddItemType.FP.Get_nonnilableType()
                            
                        }
                        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( Lns_op_not(argType.FP.Equals(self.processInfo, Ast_builtinTypeEmpty, nil, nil))) &&
                            Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(argType.FP.CanEvalWith(self.processInfo, checkType, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool)){
                            self.FP.addErrMess(firstPos, Lns_getVM().String_format("unmatch value type (index = %d) %s <- %s", []LnsAny{subIndex, argType.FP.GetTxt(self.typeNameCtrl, nil, nil), dddItemType.FP.GetTxt(nil, nil, nil)}))
                        }
                        expTypeList.Insert(Ast_TypeInfo2Stem(checkType))
                        orgExpTypeList.Insert(Ast_TypeInfo2Stem(dddItemType))
                    }
                }
            } else { 
                var expTypeInfo *Ast_TypeInfo
                expTypeInfo = expType
                if expType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                    var itemList *LnsList
                    itemList = expType.FP.Get_itemTypeInfoList()
                    if itemList.Len() > 0{
                        expTypeInfo = itemList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        
                    } else { 
                        expTypeInfo = Ast_builtinTypeStem_
                        
                    }
                }
                orgExpTypeList.Insert(Ast_TypeInfo2Stem(expTypeInfo))
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( expTypeInfo == Ast_builtinTypeNil) &&
                    Lns_GetEnv().SetStackVal( index <= typeInfoList.Len()) ).(bool)){
                    orgExpTypeList.Set(index,typeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo())
                }
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( unwrapFlag) &&
                    Lns_GetEnv().SetStackVal( expTypeInfo.FP.Get_nilable()) ).(bool)){
                    expTypeInfo = expTypeInfo.FP.Get_nonnilableType()
                    
                }
                if index <= typeInfoList.Len(){
                    var varType *Ast_TypeInfo
                    varType = typeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    var alt2typeMap *LnsMap
                    alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false)
                    if varType.FP.Get_kind() == Ast_TypeInfoKind__Box{
                        alt2typeMap = varType.FP.CreateAlt2typeMap(true)
                        
                    }
                    Ast_CanEvalCtrlTypeInfo_setupNeedAutoBoxing(alt2typeMap, self.processInfo)
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(varType.FP.Equals(self.processInfo, Ast_builtinTypeEmpty, nil, nil))) &&
                        Lns_GetEnv().SetStackVal( Lns_op_not((Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( unwrapFlag) &&
                            Lns_GetEnv().SetStackVal( expTypeInfo.FP.Equals(self.processInfo, Ast_builtinTypeNil, nil, nil)) ).(bool)))) ).(bool)){
                        var canEval bool
                        var mess LnsAny
                        canEval,mess = varType.FP.CanEvalWith(self.processInfo, expTypeInfo, Ast_CanEvalType__SetOp, alt2typeMap)
                        if Lns_op_not(canEval){
                            self.FP.addErrMess(firstPos, Lns_getVM().String_format("unmatch value type (index:%d) %s <- %s%s", []LnsAny{index, varType.FP.GetTxt(self.typeNameCtrl, nil, nil), expTypeInfo.FP.GetTxt(self.typeNameCtrl, nil, nil), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                Lns_GetEnv().SetStackVal( mess) &&
                                Lns_GetEnv().SetStackVal( Lns_getVM().String_format(" -- %s", []LnsAny{mess})) ||
                                Lns_GetEnv().SetStackVal( "") ).(string)}))
                        }
                    }
                    if varType.FP.Get_kind() == Ast_TypeInfoKind__Box{
                        typeInfoList.Set(index,self.processInfo.FP.CreateBox(accessMode, expTypeInfo))
                    }
                    if Ast_CanEvalCtrlTypeInfo_canAutoBoxing(varType, expTypeInfo){
                        updateExpList = true
                        
                        var exp *Nodes_Node
                        exp = newExpList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                        newExpList.Set(index,&Nodes_BoxingNode_create(self.nodeManager, exp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(varType)}), exp).Nodes_Node)
                        if Lns_op_not(Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(alt2typeMap, 1)){
                            self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("auto boxing error %s <- %s", []LnsAny{varType.FP.GetTxt(nil, nil, nil), expTypeInfo.FP.GetTxt(nil, nil, nil)}))
                        }
                    } else { 
                        if Lns_op_not(Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(alt2typeMap, 0)){
                            self.FP.addErrMess(newExpList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_pos(), Lns_getVM().String_format("illegal auto boxing error %s <- %s", []LnsAny{varType.FP.GetTxt(nil, nil, nil), expTypeInfo.FP.GetTxt(nil, nil, nil)}))
                        }
                    }
                }
                expTypeList.Insert(Ast_TypeInfo2Stem(expTypeInfo))
            }
        }
        if updateExpList{
            workList = self.FP.createExpListNode(workList, newExpList)
            
        }
        {
            var alt2typeMap *LnsMap
            alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false)
            {
                __exp := self.FP.checkImplicitCast(alt2typeMap, true, typeInfoList, workList, TransUnit_checkImplicitCastCallback_1263_(TransUnit_analyzeInitExp___anonymous_5892_))
                if __exp != nil {
                    _exp := __exp.(*Nodes_ExpListNode)
                    workList = _exp
                    
                }
            }
        }
        for _index, _varType := range( typeInfoList.Items ) {
            index := _index + 1
            varType := _varType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index > expTypeList.Len(){
                if Lns_op_not(varType.FP.Get_nilable()){
                    self.FP.addErrMess(firstPos, Lns_getVM().String_format("unmatch value type (index:%d) %s <- nil", []LnsAny{index, varType.FP.GetTxt(self.typeNameCtrl, nil, nil)}))
                }
            }
        }
        for _index, _typeInfo := range( expTypeList.Items ) {
            index := _index + 1
            typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( typeInfoList.Len() < index) ||
                Lns_GetEnv().SetStackVal( typeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Equals(self.processInfo, Ast_builtinTypeEmpty, nil, nil)) ).(bool){
                var workPos *Types_Position
                var workType *Ast_TypeInfo
                var workName string
                if index <= letVarList.Len(){
                    var letVar *TransUnit_LetVarInfo
                    letVar = letVarList.GetAt(index).(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
                    workPos = letVar.VarName.Pos
                    
                    workName = letVar.VarName.Txt
                    
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(typeInfo)) &&
                        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isMutable(letVar.Mutable))) ).(bool)){
                        workType = self.FP.CreateModifier(typeInfo, Ast_MutMode__IMutRe)
                        
                    } else { 
                        workType = typeInfo
                        
                    }
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( self.analyzeMode == TransUnit_AnalyzeMode__Inquire) &&
                        Lns_GetEnv().SetStackVal( self.FP.isTargetToken(letVar.VarName)) ).(bool)){
                        self.FP.dumpSymbolType(letVar.VarName.Txt, workType)
                    }
                } else { 
                    workType = typeInfo
                    
                    workPos = firstPos
                    
                    workName = ""
                    
                }
                typeInfoList.Set(index,workType)
                if _switch50083 := workType.FP.Get_kind(); _switch50083 == Ast_TypeInfoKind__Func {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( expTypeList.Len() != 1) ||
                        Lns_GetEnv().SetStackVal( workType.FP.Get_rawTxt() != "") ).(bool){
                        self.FP.addErrMess(firstPos, Lns_getVM().String_format("must set the type of variable for function. -- %s", []LnsAny{workName}))
                    }
                } else if _switch50083 == Ast_TypeInfoKind__List || _switch50083 == Ast_TypeInfoKind__Array || _switch50083 == Ast_TypeInfoKind__Set || _switch50083 == Ast_TypeInfoKind__Map {
                    self.FP.checkLiteralEmptyCollection(workPos, workName, workType)
                }
            }
        }
        return typeInfoList, letVarList, orgExpTypeList, workList
    }
    return typeInfoList, letVarList, orgExpTypeList, nil
}

// 3506: decl @lune.@base.@TransUnit.TransUnit.analyzeLetAndInitExp
func (self *TransUnit_TransUnit) analyzeLetAndInitExp(firstPos *Types_Position,letFlag bool,initMutable LnsInt,accessMode LnsInt,unwrapFlag bool)(*LnsList, *LnsList, *LnsList, LnsAny) {
    var typeInfoList *LnsList
    typeInfoList = NewLnsList([]LnsAny{})
    var letVarList *LnsList
    letVarList = NewLnsList([]LnsAny{})
    var nextToken *Types_Token
    nextToken = Parser_getEofToken()
    if letFlag{
        for {
            var mutable LnsInt
            mutable = initMutable
            nextToken = self.FP.getToken(nil)
            
            if nextToken.Txt == "mut"{
                mutable = Ast_MutMode__Mut
                
                nextToken = self.FP.getToken(nil)
                
            }
            var varName *Types_Token
            varName = self.FP.checkSymbol(nextToken, TransUnit_SymbolMode__MustNot_Or_)
            nextToken = self.FP.getToken(nil)
            
            var typeInfo *Ast_TypeInfo
            typeInfo = Ast_builtinTypeEmpty
            if nextToken.Txt == ":"{
                var refType *Nodes_RefTypeNode
                refType = self.FP.analyzeRefType(accessMode, false, Ast_isPubToExternal(accessMode))
                letVarList.Insert(TransUnit_LetVarInfo2Stem(NewTransUnit_LetVarInfo(mutable, varName, refType)))
                typeInfo = refType.FP.Get_expType()
                
                nextToken = self.FP.getToken(nil)
                
            } else { 
                letVarList.Insert(TransUnit_LetVarInfo2Stem(NewTransUnit_LetVarInfo(Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Ast_isMutable(mutable)) &&
                    Lns_GetEnv().SetStackVal( mutable) ||
                    Lns_GetEnv().SetStackVal( Ast_MutMode__IMutRe) ).(LnsInt), varName, nil)))
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(typeInfo.FP.Equals(self.processInfo, Ast_builtinTypeEmpty, nil, nil))) &&
                Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(typeInfo)) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isMutable(mutable))) ).(bool)){
                typeInfo = self.FP.CreateModifier(typeInfo, Ast_MutMode__IMutRe)
                
            }
            typeInfoList.Insert(Ast_TypeInfo2Stem(typeInfo))
            if nextToken.Txt != ","{ break }
        }
    } else { 
        for  {
            var symbolToken *Types_Token
            symbolToken = self.FP.getToken(nil)
            nextToken = self.FP.getToken(nil)
            
            var verSym LnsAny
            verSym = self.scope.FP.GetSymbolTypeInfo(symbolToken.Txt, self.scope, self.moduleScope, self.scopeAccess)
            if verSym != nil{
                verSym_9510 := verSym.(*Ast_SymbolInfo)
                letVarList.Insert(TransUnit_LetVarInfo2Stem(NewTransUnit_LetVarInfo(verSym_9510.FP.Get_mutMode(), symbolToken, nil)))
                typeInfoList.Insert(Ast_TypeInfo2Stem(verSym_9510.FP.Get_typeInfo()))
            } else {
                self.FP.addErrMess(symbolToken.Pos, Lns_getVM().String_format("not found symbol -- %s", []LnsAny{symbolToken.Txt}))
            }
            if nextToken.Txt != ","{
                break
            }
        }
    }
    if nextToken.Txt != "="{
        self.FP.Pushback()
        var orgExpTypeList *LnsList
        orgExpTypeList = NewLnsList([]LnsAny{})
        return typeInfoList, letVarList, orgExpTypeList, nil
    }
    return self.FP.analyzeInitExp(firstPos, accessMode, unwrapFlag, letVarList, typeInfoList)
}

// 3581: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclVar
func (self *TransUnit_TransUnit) analyzeDeclVar(mode LnsInt,accessMode LnsInt,firstToken *Types_Token) *Nodes_Node {
    var unwrapFlag bool
    unwrapFlag = false
    var token *Types_Token
    var continueFlag bool
    token,continueFlag = self.FP.getContinueToken()
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( continueFlag) &&
        Lns_GetEnv().SetStackVal( token.Txt == "!") ).(bool)){
        unwrapFlag = true
        
    } else { 
        self.FP.Pushback()
        if mode != Nodes_DeclVarMode__Let{
            Util_log("need '!'")
        }
    }
    if accessMode == Ast_AccessMode__Pub{
        if self.scope != self.moduleScope{
            self.FP.addErrMess(firstToken.Pos, "'pub' variable must exist top scope.")
        }
    }
    var typeInfoList *LnsList
    var letVarList *LnsList
    var orgExpTypeList *LnsList
    var expList LnsAny
    typeInfoList,letVarList,orgExpTypeList,expList = self.FP.analyzeLetAndInitExp(firstToken.Pos, mode == Nodes_DeclVarMode__Let, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode == Nodes_DeclVarMode__Sync) &&
        Lns_GetEnv().SetStackVal( Ast_MutMode__Mut) ||
        Lns_GetEnv().SetStackVal( Ast_MutMode__IMut) ).(LnsInt), accessMode, unwrapFlag)
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode == Nodes_DeclVarMode__Let) &&
        Lns_GetEnv().SetStackVal( typeInfoList.Len() == 1) ).(bool)){
        if expList != nil{
            expList_9536 := expList.(*Nodes_ExpListNode)
            var typeInfo *Ast_TypeInfo
            typeInfo = typeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var letVaInfo *TransUnit_LetVarInfo
            letVaInfo = letVarList.GetAt(1).(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( expList_9536.FP.Get_expList().Len() == 1) &&
                Lns_GetEnv().SetStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Func) ).(bool)){
                var valExp *Nodes_Node
                valExp = expList_9536.FP.Get_expList().GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
                {
                    _macroExp := Nodes_ExpMacroExpNodeDownCastF(valExp.FP)
                    if _macroExp != nil {
                        macroExp := _macroExp.(*Nodes_ExpMacroExpNode)
                        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( macroExp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Func) &&
                            Lns_GetEnv().SetStackVal( macroExp.FP.Get_stmtList().Len() == 1) ).(bool)){
                            valExp = macroExp.FP.Get_stmtList().GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
                            
                        }
                    }
                }
                {
                    _declNode := Nodes_DeclFuncNodeDownCastF(valExp.FP)
                    if _declNode != nil {
                        declNode := _declNode.(*Nodes_DeclFuncNode)
                        if Lns_op_not(declNode.FP.Get_declInfo().FP.Get_name()){
                            if Ast_isMutable(letVaInfo.Mutable){
                                self.FP.addErrMess(letVaInfo.VarName.Pos, Lns_getVM().String_format("Any function can't be mutable. -- %s", []LnsAny{letVaInfo.VarName.Txt}))
                            }
                            var letVarInfo *TransUnit_LetVarInfo
                            letVarInfo = letVarList.GetAt(1).(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
                            var newTypeInfo *Ast_TypeInfo
                            newTypeInfo = self.processInfo.FP.CreateFunc(typeInfo.FP.Get_abstractFlag(), false, typeInfo.FP.Get_scope(), typeInfo.FP.Get_kind(), typeInfo.FP.Get_parentInfo(), false, false, typeInfo.FP.Get_staticFlag(), accessMode, letVarInfo.VarName.Txt, typeInfo.FP.Get_itemTypeInfoList(), typeInfo.FP.Get_argTypeInfoList(), typeInfo.FP.Get_retTypeInfoList(), Ast_TypeInfo_isMut(typeInfo))
                            var funcSym *Ast_SymbolInfo
                            funcSym = self.FP.processAddFunc(true, self.scope, letVarInfo.VarName, newTypeInfo, Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false))
                            self.nodeManager.FP.DelNode(&declNode.Nodes_Node)
                            var declInfo *Nodes_DeclFuncInfo
                            declInfo = Nodes_DeclFuncInfo_createFrom(declNode.FP.Get_declInfo(), letVarInfo.VarName, funcSym)
                            return &Nodes_DeclFuncNode_create(self.nodeManager, declNode.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(newTypeInfo)}), declInfo).Nodes_Node
                        }
                    }
                }
            }
        }
    }
    var syncScope *Ast_Scope
    syncScope = self.scope
    if mode == Nodes_DeclVarMode__Sync{
        syncScope = self.FP.pushScope(false, nil, nil)
        
    }
    var symbolInfoList *LnsList
    symbolInfoList = NewLnsList([]LnsAny{})
    var varList *LnsList
    varList = NewLnsList([]LnsAny{})
    var syncSymbolList *LnsList
    syncSymbolList = NewLnsList([]LnsAny{})
    for _index, _letVarInfo := range( letVarList.Items ) {
        index := _index + 1
        letVarInfo := _letVarInfo.(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
        var varName *Types_Token
        varName = letVarInfo.VarName
        var typeInfo *Ast_TypeInfo
        typeInfo = typeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var varInfo *Nodes_VarInfo
        varInfo = NewNodes_VarInfo(varName, letVarInfo.VarType, typeInfo)
        varList.Insert(Nodes_VarInfo2Stem(varInfo))
        if Ast_isPubToExternal(accessMode){
            self.FP.checkPublic(varName.Pos, typeInfo)
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(letVarInfo.VarType)) &&
            Lns_GetEnv().SetStackVal( typeInfo.FP.Equals(self.processInfo, Ast_builtinTypeNil, nil, nil)) ).(bool)){
            self.FP.addErrMess(varName.Pos, Lns_getVM().String_format("need type -- %s", []LnsAny{varName.Txt}))
        }
        if mode == Nodes_DeclVarMode__Sync{
            {
                _symInfo := self.scope.FP.GetSymbolInfo(varName.Txt, self.scope, true, self.scopeAccess)
                if _symInfo != nil {
                    symInfo := _symInfo.(*Ast_SymbolInfo)
                    syncSymbolList.Insert(Ast_SymbolInfo2Stem(symInfo))
                }
            }
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( mode == Nodes_DeclVarMode__Let) ||
            Lns_GetEnv().SetStackVal( mode == Nodes_DeclVarMode__Sync) ).(bool){
            if mode == Nodes_DeclVarMode__Let{
                self.FP.checkShadowing(varName.Pos, varName.Txt, self.scope)
            }
            var orgExpType *Ast_TypeInfo
            orgExpType = Ast_builtinTypeStem_
            if Lns_op_not(unwrapFlag){
                orgExpType = Ast_builtinTypeEmpty
                
            }
            if index <= orgExpTypeList.Len(){
                orgExpType = orgExpTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                
            }
            var hasValue bool
            hasValue = false
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(unwrapFlag)) &&
                Lns_GetEnv().SetStackVal( orgExpType != Ast_builtinTypeEmpty) ||
                Lns_GetEnv().SetStackVal( unwrapFlag) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(orgExpType.FP.Get_nilable())) ).(bool){
                hasValue = true
                
            }
            self.scope.FP.AddVar(self.processInfo, accessMode, varName.Txt, varName.Pos, typeInfo, letVarInfo.Mutable, hasValue)
        }
        symbolInfoList.Insert(Ast_SymbolInfo2Stem(Lns_unwrap( self.scope.FP.GetSymbolInfo(varName.Txt, self.scope, true, self.scopeAccess)).(*Ast_SymbolInfo)))
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode != Nodes_DeclVarMode__Sync) &&
        Lns_GetEnv().SetStackVal( self.macroScope) )){
        self.macroCtrl.FP.RegistVar(symbolInfoList)
    }
    var unwrapBlock LnsAny
    unwrapBlock = nil
    var thenBlock LnsAny
    thenBlock = nil
    if unwrapFlag{
        var scope *Ast_Scope
        scope = self.FP.pushScope(false, nil, nil)
        for _index, _letVarInfo := range( letVarList.Items ) {
            index := _index + 1
            letVarInfo := _letVarInfo.(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
            if letVarInfo.VarName.Txt != "_"{
                self.FP.addLocalVar(letVarInfo.VarName.Pos, false, true, "_" + letVarInfo.VarName.Txt, orgExpTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil)
            }
        }
        unwrapBlock = self.FP.analyzeBlock(Nodes_BlockKind__LetUnwrap, TransUnit_TentativeMode__Start, scope, nil)
        
        self.FP.popScope()
        if unwrapBlock != nil{
            unwrapBlock_9585 := unwrapBlock.(*Nodes_BlockNode)
            if _switch51597 := mode; _switch51597 == Nodes_DeclVarMode__Let || _switch51597 == Nodes_DeclVarMode__Sync {
                var breakKind LnsInt
                breakKind = unwrapBlock_9585.FP.GetBreakKind(Nodes_CheckBreakMode__Normal)
                for _, _symbolInfo := range( symbolInfoList.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    if breakKind != Nodes_BreakKind__None{
                        self.tentativeSymbol.FP.CheckAndExclude(symbolInfo)
                        symbolInfo.FP.UpdateValue(symbolInfo.FP.Get_pos())
                    } else { 
                        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_name() != "_") &&
                            Lns_GetEnv().SetStackVal( Lns_op_not(self.tentativeSymbol.FP.CheckAndExclude(symbolInfo))) ).(bool)){
                            if Lns_op_not(symbolInfo.FP.Get_hasValueFlag()){
                                self.FP.addErrMess(unwrapBlock_9585.FP.Get_pos(), "This variable isn't set -- " + (symbolInfo.FP.Get_name()))
                            }
                        }
                    }
                }
            } else if _switch51597 == Nodes_DeclVarMode__Unwrap {
                for _, _symbolInfo := range( symbolInfoList.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.UpdateValue(firstToken.Pos)
                }
            }
        }
        token = self.FP.getToken(true)
        
        if token.Txt == "then"{
            thenBlock = self.FP.analyzeBlock(Nodes_BlockKind__LetUnwrapThenDo, TransUnit_TentativeMode__Finish, scope, nil)
            
        } else { 
            self.FP.Pushback()
            self.FP.finishTentativeSymbol(true)
        }
    }
    var syncBlock LnsAny
    syncBlock = nil
    if mode == Nodes_DeclVarMode__Sync{
        self.FP.checkNextToken("do")
        syncBlock = self.FP.analyzeBlock(Nodes_BlockKind__LetUnwrapThenDo, TransUnit_TentativeMode__Simple, syncScope, nil)
        
        self.FP.popScope()
    }
    self.FP.checkNextToken(";")
    var node *Nodes_DeclVarNode
    node = Nodes_DeclVarNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), mode, accessMode, false, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncSymbolList, syncBlock)
    return &node.Nodes_Node
}

// 3810: decl @lune.@base.@TransUnit.TransUnit.analyzeIfUnwrap
func (self *TransUnit_TransUnit) analyzeIfUnwrap(firstToken *Types_Token) *Nodes_IfUnwrapNode {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var typeInfoList *LnsList
    typeInfoList = NewLnsList([]LnsAny{})
    var varNameList *LnsList
    varNameList = NewLnsList([]LnsAny{})
    var expList *Nodes_ExpListNode
    var varList *LnsList
    varList = NewLnsList([]LnsAny{})
    var workTypeInfoList *LnsList
    var letVarList *LnsList
    var workExpList LnsAny
    if nextToken.Txt == "let"{
        workTypeInfoList, letVarList, _, workExpList = self.FP.analyzeLetAndInitExp(firstToken.Pos, true, Ast_MutMode__IMut, Ast_AccessMode__Local, true)
        
    } else { 
        self.FP.Pushback()
        var tmpTypeInfoList *LnsList
        tmpTypeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeEmpty)})
        var tmpLetVarList *LnsList
        tmpLetVarList = NewLnsList([]LnsAny{TransUnit_LetVarInfo2Stem(NewTransUnit_LetVarInfo(Ast_MutMode__Mut, NewTypes_Token(Types_TokenKind__Symb, "_exp", firstToken.Pos, false, nil), nil))})
        workTypeInfoList, letVarList, _, workExpList = self.FP.analyzeInitExp(firstToken.Pos, Ast_AccessMode__Local, true, tmpLetVarList, tmpTypeInfoList)
        
    }
    typeInfoList = workTypeInfoList
    
    if workExpList != nil{
        workExpList_9620 := workExpList.(*Nodes_ExpListNode)
        expList = workExpList_9620
        
    } else {
        self.FP.addErrMess(nextToken.Pos, "if! let has illegal init val.")
        self.FP.Error("system error")
    }
    for _, _varInfo := range( letVarList.Items ) {
        varInfo := _varInfo.(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
        varNameList.Insert(Types_Token2Stem(varInfo.VarName))
    }
    var scope *Ast_Scope
    scope = self.FP.pushScope(false, nil, nil)
    for _index, _expType := range( typeInfoList.Items ) {
        index := _index + 1
        expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > varNameList.Len(){
            break
        }
        var varName *Types_Token
        varName = varNameList.GetAt(index).(Types_TokenDownCast).ToTypes_Token()
        varList.Insert(Ast_SymbolInfo2Stem(self.FP.addLocalVar(varName.Pos, false, true, varName.Txt, expType, Ast_MutMode__IMut, nil)))
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(Nodes_BlockKind__IfUnwrap, TransUnit_TentativeMode__Start, scope, nil)
    self.FP.popScope()
    var elseBlock LnsAny
    elseBlock = nil
    nextToken = self.FP.getToken(true)
    
    if nextToken.Txt == "else"{
        elseBlock = self.FP.analyzeBlock(Nodes_BlockKind__Else, TransUnit_TentativeMode__Finish, nil, nil)
        
    } else { 
        self.FP.finishTentativeSymbol(false)
        self.FP.Pushback()
    }
    var hasCond bool
    hasCond = false
    for _index, _expNode := range( expList.FP.Get_expList().Items ) {
        index := _index + 1
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        if index != expList.FP.Get_expList().Len(){
            if Ast_isConditionalbe(self.processInfo, expNode.FP.Get_expType()){
                hasCond = true
                
                break
            }
        } else { 
            for _, _expType := range( expNode.FP.Get_expTypeList().Items ) {
                expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if Ast_isConditionalbe(self.processInfo, expType){
                    hasCond = true
                    
                    break
                }
            }
        }
    }
    if Lns_op_not(hasCond){
        self.FP.addErrMess(firstToken.Pos, "This condition never be false")
    }
    for _, _varSym := range( varList.Items ) {
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if _switch52274 := varSym.FP.Get_name(); _switch52274 == "_" || _switch52274 == "_exp" {
        } else {
            if Lns_op_not(varSym.FP.Get_posForModToRef()){
                self.FP.addWarnMess(Lns_unwrap( varSym.FP.Get_pos()).(*Types_Position), Lns_getVM().String_format("This symbol has no referer -- %s", []LnsAny{varSym.FP.Get_name()}))
            }
        }
    }
    return Nodes_IfUnwrapNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), varList, expList, block, elseBlock)
}

// 3910: decl @lune.@base.@TransUnit.TransUnit.analyzeWhen
func (self *TransUnit_TransUnit) analyzeWhen(firstToken *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    var continueFlag bool
    nextToken,continueFlag = self.FP.getContinueToken()
    if Lns_op_not((Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( continueFlag) &&
        Lns_GetEnv().SetStackVal( nextToken.Txt == "!") ).(bool))){
        self.FP.Pushback()
        self.FP.addErrMess(nextToken.Pos, "'when' need '!'")
    }
    var symListNode *Nodes_ExpListNode
    symListNode = self.FP.analyzeExpList(false, false, false, nil, nil, nil)
    var scope *Ast_Scope
    scope = self.FP.pushScope(false, nil, nil)
    var symPairList *LnsList
    symPairList = NewLnsList([]LnsAny{})
    for _, _expNode := range( symListNode.FP.Get_expList().Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        {
            _refNode := Nodes_ExpRefNodeDownCastF(expNode.FP)
            if _refNode != nil {
                refNode := _refNode.(*Nodes_ExpRefNode)
                if expNode.FP.Get_expType().FP.Get_nilable(){
                    var symbolInfo *Ast_SymbolInfo
                    symbolInfo = refNode.FP.Get_symbolInfo()
                    var newSymbolInfo *Ast_SymbolInfo
                    newSymbolInfo = self.FP.addLocalVar(firstToken.Pos, false, expNode.FP.CanBeLeft(), refNode.FP.Get_symbolInfo().FP.Get_name(), expNode.FP.Get_expType().FP.Get_nonnilableType(), Ast_MutMode__IMut, true)
                    symPairList.Insert(Nodes_UnwrapSymbolPair2Stem(NewNodes_UnwrapSymbolPair(symbolInfo, newSymbolInfo)))
                } else { 
                    self.FP.addErrMess(expNode.FP.Get_pos(), Lns_getVM().String_format("This type isn't nilable. -- %s", []LnsAny{expNode.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
                }
            } else {
                self.FP.addErrMess(expNode.FP.Get_pos(), "'when' support only local variables or arguments.")
            }
        }
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(Nodes_BlockKind__When, TransUnit_TentativeMode__Start, scope, nil)
    self.FP.popScope()
    var elseBlock LnsAny
    elseBlock = nil
    nextToken = self.FP.getToken(true)
    
    if nextToken.Txt == "else"{
        elseBlock = self.FP.analyzeBlock(Nodes_BlockKind__Else, TransUnit_TentativeMode__Finish, nil, nil)
        
    } else { 
        self.FP.finishTentativeSymbol(false)
        self.FP.Pushback()
    }
    return &Nodes_WhenNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), symPairList, block, elseBlock).Nodes_Node
}

// 27: decl @lune.@base.@TransUnit.TransUnit.MultiTo1
func (self *TransUnit_TransUnit) MultiTo1(exp *Nodes_Node) *Nodes_Node {
    if exp.FP.Get_expTypeList().Len() > 1{
        exp = &Nodes_ExpMultiTo1Node_create(self.nodeManager, exp.FP.Get_pos(), exp.FP.Get_macroArgFlag(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp.FP.Get_expType())}), exp).Nodes_Node
        
    } else { 
        {
            _dddType := Ast_DDDTypeInfoDownCastF(exp.FP.Get_expType().FP)
            if _dddType != nil {
                dddType := _dddType.(*Ast_DDDTypeInfo)
                exp = &Nodes_ExpMultiTo1Node_create(self.nodeManager, exp.FP.Get_pos(), exp.FP.Get_macroArgFlag(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddType.FP.Get_typeInfo().FP.Get_nilableTypeInfo())}), exp).Nodes_Node
                
            }
        }
    }
    return exp
}

// 46: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOneRVal
func (self *TransUnit_TransUnit) analyzeExpOneRVal(allowNoneType bool,skipOp2Flag bool,opLevel LnsAny,expectType LnsAny) *Nodes_Node {
    var exp *Nodes_Node
    exp = self.FP.analyzeExp(allowNoneType, skipOp2Flag, false, opLevel, expectType)
    exp = self.FP.MultiTo1(exp)
    
    if expectType != nil{
        expectType_9693 := expectType.(*Ast_TypeInfo)
        if _switch52844 := expectType_9693.FP.Get_kind(); _switch52844 == Ast_TypeInfoKind__IF || _switch52844 == Ast_TypeInfoKind__Class {
            var expOrgType *Ast_TypeInfo
            expOrgType = exp.FP.Get_expType().FP.Get_nonnilableType().FP.Get_srcTypeInfo()
            var exceptOrgType *Ast_TypeInfo
            exceptOrgType = expectType_9693.FP.Get_nonnilableType().FP.Get_srcTypeInfo()
            if expOrgType.FP.IsInheritFrom(self.processInfo, exceptOrgType, nil){
                exp = &Nodes_ExpCastNode_create(self.nodeManager, exp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(expectType_9693)}), exp, expectType_9693, Nodes_CastKind__Implicit).Nodes_Node
                
            }
        }
    }
    return exp
}

// 77: decl @lune.@base.@TransUnit.TransUnit.createExpList
func (self *TransUnit_TransUnit) createExpList(pos *Types_Position,expTypeList *LnsList,expList *LnsList,followOn bool,abbrNode LnsAny) *Nodes_ExpListNode {
    var workList *LnsList
    workList = NewLnsList([]LnsAny{})
    var mRetExp LnsAny
    mRetExp = nil
    if expList.Len() > 0{
        for _index, _exp := range( expList.Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( expTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeMultiExp) &&
                Lns_GetEnv().SetStackVal( Nodes_hasMultiValNode(exp)) ).(bool)){
                if index != expList.Len(){
                    var firstType *Ast_TypeInfo
                    {
                        _dddType := Ast_DDDTypeInfoDownCastF(exp.FP.Get_expType().FP)
                        if _dddType != nil {
                            dddType := _dddType.(*Ast_DDDTypeInfo)
                            firstType = dddType.FP.Get_typeInfo().FP.Get_nilableTypeInfo()
                            
                        } else {
                            firstType = exp.FP.Get_expType()
                            
                        }
                    }
                    workList.Insert(Nodes_ExpMultiTo1Node2Stem(Nodes_ExpMultiTo1Node_create(self.nodeManager, exp.FP.Get_pos(), exp.FP.Get_macroArgFlag(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(firstType)}), exp)))
                } else { 
                    mRetExp = NewNodes_MRetExp(exp, index)
                    
                    for _listIndex, _expType := range( exp.FP.Get_expTypeList().Items ) {
                        listIndex := _listIndex + 1
                        expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if listIndex != 1{
                            expTypeList.Insert(Ast_TypeInfo2Stem(expType))
                        }
                    }
                    for _retIndex, _retType := range( exp.FP.Get_expTypeList().Items ) {
                        retIndex := _retIndex + 1
                        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if retIndex == 1{
                            workList.Insert(Nodes_ExpMRetNode2Stem(Nodes_ExpMRetNode_create(self.nodeManager, exp.FP.Get_pos(), exp.FP.Get_macroArgFlag(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(retType)}), exp)))
                        } else { 
                            workList.Insert(Nodes_ExpAccessMRetNode2Stem(Nodes_ExpAccessMRetNode_create(self.nodeManager, exp.FP.Get_pos(), exp.FP.Get_macroArgFlag(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(retType)}), exp, retIndex)))
                        }
                    }
                }
            } else { 
                workList.Insert(Nodes_Node2Stem(exp))
            }
        }
    }
    if abbrNode != nil{
        abbrNode_9731 := abbrNode.(*Nodes_AbbrNode)
        workList.Insert(Nodes_AbbrNode2Stem(abbrNode_9731))
    }
    return Nodes_ExpListNode_create(self.nodeManager, pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), expTypeList, workList, mRetExp, followOn)
}

// 169: decl @lune.@base.@TransUnit.TransUnit.analyzeExpList
func (self *TransUnit_TransUnit) analyzeExpList(allowNoneType bool,skipOp2Flag bool,canLeftExp bool,expNode LnsAny,expectTypeList LnsAny,contExpect LnsAny) *Nodes_ExpListNode {
    var expList *LnsList
    expList = NewLnsList([]LnsAny{})
    var pos LnsAny
    pos = nil
    var expTypeList *LnsList
    expTypeList = NewLnsList([]LnsAny{})
    if expNode != nil{
        expNode_9746 := expNode.(*Nodes_Node)
        pos = expNode_9746.FP.Get_pos()
        
        expList.Insert(Nodes_Node2Stem(expNode_9746))
        expTypeList.Insert(Ast_TypeInfo2Stem(expNode_9746.FP.Get_expType()))
    }
    var index LnsInt
    index = 1
    var abbrNode LnsAny
    abbrNode = nil
    var followOn bool
    followOn = false
    for {
        var expectType LnsAny
        expectType = nil
        var allowNoneTypeOne bool
        allowNoneTypeOne = allowNoneType
        if expectTypeList != nil{
            expectTypeList_9754 := expectTypeList.(*LnsList)
            if expectTypeList_9754.Len() > 0{
                var checkIndex LnsInt
                checkIndex = index
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( index > expectTypeList_9754.Len()) &&
                    Lns_GetEnv().SetStackVal( contExpect) )){
                    checkIndex = expectTypeList_9754.Len()
                    
                }
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( checkIndex <= expectTypeList_9754.Len()) &&
                    Lns_GetEnv().SetStackVal( expectTypeList_9754.GetAt(checkIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeNone) ).(bool)){
                    var worktype *Ast_TypeInfo
                    worktype = expectTypeList_9754.GetAt(checkIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    expectType = worktype
                    
                    if worktype == Ast_builtinTypeExp{
                        allowNoneTypeOne = true
                        
                    }
                }
            }
        }
        var exp *Nodes_Node
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( self.macroCtrl.FP.Get_analyzeInfo().FP.Get_mode() == Nodes_MacroMode__AnalyzeArg) &&
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( expectType == Ast_builtinTypeExp) ||
                Lns_GetEnv().SetStackVal( expectType == Ast_builtinTypeMultiExp) ).(bool))) ).(bool)){
            self.macroCtrl.FP.SwitchMacroMode()
            exp = self.FP.analyzeExp(allowNoneTypeOne, skipOp2Flag, canLeftExp, 0, expectType)
            
            self.macroCtrl.FP.RestoreMacroMode()
        } else { 
            exp = self.FP.analyzeExp(allowNoneTypeOne, skipOp2Flag, canLeftExp, 0, expectType)
            
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(allowNoneTypeOne)) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(exp.FP.CanBeRight(self.processInfo))) ).(bool)){
            self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("This arg can't be r-value. -- %s", []LnsAny{Nodes_getNodeKindName(exp.FP.Get_kind())}))
        }
        if Lns_op_not(pos){
            pos = exp.FP.Get_pos()
            
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( expectType == Ast_builtinTypeExp) ||
            Lns_GetEnv().SetStackVal( expectType == Ast_builtinTypeMultiExp) ).(bool){
            exp = &Nodes_ExpMacroArgExpNode_create(self.nodeManager, exp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), exp.FP.Get_expTypeList(), Macro_nodeToCodeTxt(exp, self.moduleType)).Nodes_Node
            
            expTypeList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( expectType).(*Ast_TypeInfo)))
        } else { 
            expTypeList.Insert(Ast_TypeInfo2Stem(exp.FP.Get_expType()))
        }
        expList.Insert(Nodes_Node2Stem(exp))
        var token *Types_Token
        token = self.FP.getToken(true)
        if token.Txt == "**"{
            if Lns_op_not(Nodes_hasMultiValNode(exp)){
                self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("This arg(%d) doesn't have multiple value. It must not use '**'", []LnsAny{index}))
            }
            followOn = true
            
            token = self.FP.getToken(true)
            
        }
        if token.Txt == "##"{
            if exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__DDD{
                self.FP.addErrMess(token.Pos, "'##' can't use with '...'")
            }
            abbrNode = Nodes_AbbrNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeAbbr)}))
            
            self.FP.getToken(nil)
            break
        }
        index = index + 1
        
        if self.macroCtrl.FP.Get_analyzeInfo().FP.EqualsArgTypeList(expectTypeList){
            self.macroCtrl.FP.Get_analyzeInfo().FP.NextArg()
        }
        if token.Txt != ","{ break }
    }
    var expListNode *Nodes_ExpListNode
    expListNode = self.FP.createExpList(Lns_unwrapDefault( pos, self.FP.CreatePosition(0, 0)).(*Types_Position), expTypeList, expList, followOn, abbrNode)
    if Lns_op_not(allowNoneType){
        for _expIndex, _expType := range( expTypeList.Items ) {
            expIndex := _expIndex + 1
            expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if expType == Ast_builtinTypeNone{
                self.FP.addErrMess(Lns_unwrapDefault( pos, self.FP.CreatePosition(0, 0)).(*Types_Position), Lns_getVM().String_format("The type of exp(%d) is None!!", []LnsAny{expIndex}))
            }
        }
    }
    self.FP.Pushback()
    return expListNode
}

// 291: decl @lune.@base.@TransUnit.TransUnit.checkSymbolHavingValue
func (self *TransUnit_TransUnit) checkSymbolHavingValue(pos *Types_Position,symbolInfoList *LnsList) {
    for _, _symbolInfo := range( symbolInfoList.Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if symbolInfo.FP.Get_kind() == Ast_SymbolKind__Var{
            if Lns_op_not(symbolInfo.FP.Get_hasValueFlag()){
                self.FP.addErrMess(pos, Lns_getVM().String_format("this variable have no value. -- %s", []LnsAny{symbolInfo.FP.Get_name()}))
            }
        }
    }
}

// 305: decl @lune.@base.@TransUnit.TransUnit.analyzeExpRefItem
func (self *TransUnit_TransUnit) analyzeExpRefItem(token *Types_Token,exp *Nodes_Node,nilAccess bool) *Nodes_Node {
    self.FP.checkSymbolHavingValue(exp.FP.Get_pos(), exp.FP.GetSymbolInfo())
    var expType *Ast_TypeInfo
    expType = exp.FP.Get_expType()
    expType = expType.FP.Get_extedType()
    
    if nilAccess{
        if Lns_op_not(exp.FP.Get_expType().FP.Get_nilable()){
            self.FP.addWarnMess(token.Pos, Lns_getVM().String_format("This is not nilable. -- %s", []LnsAny{expType.FP.GetTxt(nil, nil, nil)}))
            nilAccess = false
            
        } else { 
            expType = expType.FP.Get_nonnilableType()
            
        }
    } else if exp.FP.Get_expType().FP.Get_nilable(){
        self.FP.addErrMess(token.Pos, Lns_getVM().String_format("could not access with []. Use '$[]'-- %s", []LnsAny{exp.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
    }
    var expectItemType LnsAny
    expectItemType = nil
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_builtinTypeStem_
    var indexTypeInfo *Ast_TypeInfo
    indexTypeInfo = Ast_builtinTypeInt
    if expType.FP.Get_kind() == Ast_TypeInfoKind__Map{
        var itemTypeList *LnsList
        itemTypeList = expType.FP.Get_itemTypeInfoList()
        typeInfo = itemTypeList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        
        indexTypeInfo = itemTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        
        expectItemType = itemTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(typeInfo.FP.Equals(self.processInfo, Ast_builtinTypeStem_, nil, nil))) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(typeInfo.FP.Get_nilable())) ).(bool)){
            typeInfo = typeInfo.FP.Get_nilableTypeInfo()
            
        }
    } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( expType.FP.Get_kind() == Ast_TypeInfoKind__Array) ||
        Lns_GetEnv().SetStackVal( expType.FP.Get_kind() == Ast_TypeInfoKind__List) ).(bool){
        typeInfo = expType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        
    } else if expType.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil){
        typeInfo = Ast_builtinTypeInt
        
    } else if expType.FP.Equals(self.processInfo, Ast_builtinTypeStem, nil, nil){
        indexTypeInfo = Ast_builtinTypeStem
        
        typeInfo = Ast_builtinTypeStem_
        
    } else { 
        self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("could not access with []. -- %s", []LnsAny{exp.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
    }
    if nilAccess{
        self.helperInfo.UseNilAccess = true
        
        if Lns_op_not(typeInfo.FP.Get_nilable()){
            typeInfo = typeInfo.FP.Get_nilableTypeInfo()
            
        }
    }
    if Ast_TypeInfo_isMut(typeInfo){
        if expType.FP.Get_mutMode() == Ast_MutMode__IMutRe{
            typeInfo = self.FP.CreateModifier(typeInfo, Ast_MutMode__IMutRe)
            
        }
    }
    if Ast_isExtType(exp.FP.Get_expType().FP.Get_nonnilableType()){
        typeInfo = self.FP.createExtType(exp.FP.Get_pos(), typeInfo)
        
    }
    var indexExp *Nodes_Node
    indexExp = self.FP.analyzeExpOneRVal(false, false, nil, expectItemType)
    if Lns_op_not(indexExp.FP.CanBeRight(self.processInfo)){
        self.FP.addErrMess(indexExp.FP.Get_pos(), "This node can't use index")
    }
    if Lns_op_not(Lns_car(indexTypeInfo.FP.CanEvalWith(self.processInfo, indexExp.FP.Get_expType(), Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)){
        self.FP.addErrMess(indexExp.FP.Get_pos(), Lns_getVM().String_format("unmatch index type -- %s, %s", []LnsAny{indexTypeInfo.FP.GetTxt(nil, nil, nil), indexExp.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
    }
    self.FP.checkNextToken("]")
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( expType.FP.Get_kind() == Ast_TypeInfoKind__Array) ||
        Lns_GetEnv().SetStackVal( expType.FP.Get_kind() == Ast_TypeInfoKind__List) ).(bool){
        {
            _indexLit := TransUnit_convExp54460(Lns_2DDD(indexExp.FP.GetLiteral()))
            if _indexLit != nil {
                indexLit := _indexLit
                switch _exp54458 := indexLit.(type) {
                case *Nodes_Literal__Int:
                val := _exp54458.Val1
                    if val <= 0{
                        self.FP.addWarnMess(indexExp.FP.Get_pos(), Lns_getVM().String_format("index <= -1 (%d)", []LnsAny{val}))
                    }
                }
            }
        }
    }
    var threading bool
    threading = false
    if nilAccess{
        threading = self.FP.checkThreading(token.Pos)
        
    }
    return &Nodes_ExpRefItemNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), exp, nilAccess, threading, nil, indexExp).Nodes_Node
}

// 439: decl @lune.@base.@TransUnit.TransUnit.checkImplicitCast
func (self *TransUnit_TransUnit) checkImplicitCast(alt2typeMap *LnsMap,validCastToGen bool,dstTypeList *LnsList,expListNode *Nodes_ExpListNode,callback TransUnit_checkImplicitCastCallback_1263_) LnsAny {
    var expNodeList *LnsList
    expNodeList = expListNode.FP.Get_expList()
    var newExpNodeList *LnsList
    newExpNodeList = NewLnsList([]LnsAny{})
    var expTypeList *LnsList
    expTypeList = NewLnsList([]LnsAny{})
    var hasModNode bool
    hasModNode = false
    var process func(index LnsInt,dstType *Ast_TypeInfo,expNode *Nodes_Node,workNode *Nodes_Node)(*Nodes_Node, bool)
    process = func(index LnsInt,dstType *Ast_TypeInfo,expNode *Nodes_Node,workNode *Nodes_Node)(*Nodes_Node, bool) {
        {
            _repNode := callback(dstType, expNode)
            if _repNode != nil {
                repNode := _repNode.(*Nodes_Node)
                if Lns_op_not(hasModNode){
                    hasModNode = true
                    
                }
                workNode = repNode
                
            } else {
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( expNode.FP.Get_expType() != Ast_builtinTypeNil) &&
                    Lns_GetEnv().SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(dstType.FP.Get_srcTypeInfo().FP.Equals(self.processInfo, expNode.FP.Get_expType().FP.Get_srcTypeInfo(), nil, nil))) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(dstType.FP.Get_nonnilableType().FP.Get_srcTypeInfo().FP.Equals(self.processInfo, expNode.FP.Get_expType().FP.Get_srcTypeInfo(), nil, nil))) ).(bool)){
                    if expNode.FP.Get_kind() != Nodes_NodeKind_get_Abbr(){
                        if Lns_op_not(hasModNode){
                            hasModNode = true
                            
                        }
                        if dstType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                            var argList *LnsList
                            argList = NewLnsList([]LnsAny{})
                            var argTypeList *LnsList
                            argTypeList = NewLnsList([]LnsAny{})
                            {
                                var _from54809 LnsInt = index
                                var _to54809 LnsInt = expNodeList.Len()
                                for _work54809 := _from54809; _work54809 <= _to54809; _work54809++ {
                                    workIndex := _work54809
                                    var appNode *Nodes_Node
                                    appNode = expNodeList.GetAt(workIndex).(Nodes_NodeDownCast).ToNodes_Node()
                                    argList.Insert(Nodes_Node2Stem(appNode))
                                    if workIndex == expNodeList.Len(){
                                        for _, _expType := range( appNode.FP.Get_expTypeList().Items ) {
                                            expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                                            argTypeList.Insert(Ast_TypeInfo2Stem(expType))
                                        }
                                    } else { 
                                        argTypeList.Insert(Ast_TypeInfo2Stem(appNode.FP.Get_expType()))
                                    }
                                }
                            }
                            var mRetExp LnsAny
                            {
                                _workMRetExp := expListNode.FP.Get_mRetExp()
                                if _workMRetExp != nil {
                                    workMRetExp := _workMRetExp.(*Nodes_MRetExp)
                                    mRetExp = NewNodes_MRetExp(workMRetExp.FP.Get_exp(), workMRetExp.FP.Get_index() - index + 1)
                                    
                                    if argList.Len() == 0{
                                        return &Nodes_ExpSubDDDNode_create(self.nodeManager, expNode.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), workMRetExp.FP.Get_exp().FP.Get_expTypeList(), workMRetExp.FP.Get_exp(), index - workMRetExp.FP.Get_index()).Nodes_Node, true
                                    }
                                } else {
                                    mRetExp = nil
                                    
                                }
                            }
                            var newExpListNode *Nodes_ExpListNode
                            newExpListNode = Nodes_ExpListNode_create(self.nodeManager, expNode.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), argTypeList, argList, mRetExp, expListNode.FP.Get_followOn())
                            workNode = &Nodes_ExpToDDDNode_create(self.nodeManager, expNode.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dstType)}), newExpListNode).Nodes_Node
                            
                            return workNode, true
                        } else { 
                            var castType *Ast_TypeInfo
                            if validCastToGen{
                                castType = Lns_unwrapDefault( alt2typeMap.Items[dstType], dstType).(*Ast_TypeInfo)
                                
                            } else { 
                                castType = dstType
                                
                            }
                            workNode = &Nodes_ExpCastNode_create(self.nodeManager, expNode.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(expNode.FP.Get_expType())}), expNode, castType, Nodes_CastKind__Implicit).Nodes_Node
                            
                        }
                    }
                }
            }
        }
        return workNode, false
    }
    for _index, _expNode := range( expNodeList.Items ) {
        index := _index + 1
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        var workNode *Nodes_Node
        workNode = expNode
        var stopFlag bool
        stopFlag = false
        if dstTypeList.Len() >= index{
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( index == expNodeList.Len()) &&
                Lns_GetEnv().SetStackVal( expNode.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__DDD) ).(bool)){
                {
                    var _from55140 LnsInt = index
                    var _to55140 LnsInt = dstTypeList.Len()
                    for _work55140 := _from55140; _work55140 <= _to55140; _work55140++ {
                        dstIndex := _work55140
                        workNode = expNode
                        
                        workNode, stopFlag = process(dstIndex, dstTypeList.GetAt(dstIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), expNode, workNode)
                        
                        newExpNodeList.Insert(Nodes_Node2Stem(workNode))
                        expTypeList.Insert(Ast_TypeInfo2Stem(workNode.FP.Get_expType()))
                    }
                }
                break
            } else { 
                workNode, stopFlag = process(index, dstTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), expNode, workNode)
                
            }
        }
        newExpNodeList.Insert(Nodes_Node2Stem(workNode))
        expTypeList.Insert(Ast_TypeInfo2Stem(workNode.FP.Get_expType()))
        if stopFlag{
            break
        }
    }
    if Lns_op_not(hasModNode){
        return nil
    }
    var newMRetExp LnsAny
    newMRetExp = nil
    {
        _mRetExp := expListNode.FP.Get_mRetExp()
        if _mRetExp != nil {
            mRetExp := _mRetExp.(*Nodes_MRetExp)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( mRetExp.FP.Get_index() <= dstTypeList.Len()) &&
                Lns_GetEnv().SetStackVal( dstTypeList.GetAt(mRetExp.FP.Get_index()).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() != Ast_TypeInfoKind__DDD) ).(bool)){
                newMRetExp = mRetExp
                
            }
        }
    }
    return Nodes_ExpListNode_create(self.nodeManager, expListNode.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), expTypeList, newExpNodeList, newMRetExp, expListNode.FP.Get_followOn())
}

// 604: decl @lune.@base.@TransUnit.TransUnit.checkMatchType
func (self *TransUnit_TransUnit) checkMatchType(message string,pos *Types_Position,dstTypeList *LnsList,expListNode LnsAny,allowDstShort bool,warnForFollow bool,workAlt2typeMap LnsAny)(LnsInt, *LnsMap, LnsAny, *LnsList) {
    var expNodeList *LnsList
    
    {
        _expNodeList := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expListNode) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_ExpListNode).FP.Get_expList()}))
        if _expNodeList == nil{
            expNodeList = NewLnsList([]LnsAny{})
            
        } else {
            expNodeList = _expNodeList.(*LnsList)
        }
    }
    var warnForFollowSrcIndex LnsAny
    warnForFollowSrcIndex = nil
    var expTypeList *LnsList
    expTypeList = NewLnsList([]LnsAny{})
    var workExpNodeList *LnsList
    workExpNodeList = expNodeList
    var hasAbbr bool
    hasAbbr = false
    if expNodeList.Len() > 0{
        if expNodeList.GetAt(expNodeList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_kind() == Nodes_NodeKind_get_Abbr(){
            hasAbbr = true
            
            var workList *LnsList
            workList = NewLnsList([]LnsAny{})
            for _, _node := range( expNodeList.Items ) {
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                workList.Insert(Nodes_Node2Stem(node))
            }
            workList.Remove(nil)
            workExpNodeList = workList
            
        }
    }
    var realExpNum LnsInt
    realExpNum = -1
    for _index, _expNode := range( workExpNodeList.Items ) {
        index := _index + 1
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( realExpNum == -1) &&
            Lns_GetEnv().SetStackVal( expNode.FP.Get_kind() == Nodes_NodeKind_get_ExpAccessMRet()) ).(bool)){
            realExpNum = index - 1
            
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( index == workExpNodeList.Len()) &&
            Lns_GetEnv().SetStackVal( Nodes_ExpMacroArgExpNode2Stem(Nodes_ExpMacroArgExpNodeDownCastF(expNode.FP)) == nil) ).(bool)){
            for _, _expType := range( expNode.FP.Get_expTypeList().Items ) {
                expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                expTypeList.Insert(Ast_TypeInfo2Stem(expType))
            }
        } else { 
            expTypeList.Insert(Ast_TypeInfo2Stem(expNode.FP.Get_expType()))
        }
    }
    if realExpNum == -1{
        realExpNum = workExpNodeList.Len()
        
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( warnForFollow) &&
        Lns_GetEnv().SetStackVal( expTypeList.Len() > realExpNum) ).(bool)){
        warnForFollowSrcIndex = realExpNum + 1
        
    }
    if hasAbbr{
        expTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeAbbr))
    }
    var alt2typeMap *LnsMap
    if workAlt2typeMap != nil{
        workAlt2typeMap_9930 := workAlt2typeMap.(*LnsMap)
        alt2typeMap = workAlt2typeMap_9930
        
    } else {
        alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false)
        
    }
    Ast_CanEvalCtrlTypeInfo_setupNeedAutoBoxing(alt2typeMap, self.processInfo)
    var result LnsInt
    var mess string
    result,mess = Ast_TypeInfo_checkMatchType(self.processInfo, dstTypeList, expTypeList, allowDstShort, warnForFollowSrcIndex, alt2typeMap)
    if _switch55725 := result; _switch55725 == Ast_MatchType__Error {
        self.FP.addErrMess(pos, Lns_getVM().String_format("%s: %s", []LnsAny{message, mess}))
    } else if _switch55725 == Ast_MatchType__Warn {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(self.ctrl_info.CheckingDefineAbbr)) &&
            Lns_GetEnv().SetStackVal( Code_isMessageOf(Code_ID__nothing_define_abbr, mess)) ).(bool)){
        } else { 
            self.FP.addWarnMess(pos, Lns_getVM().String_format("%s: %s", []LnsAny{message, mess}))
        }
    }
    if expListNode != nil{
        expListNode_9939 := expListNode.(*Nodes_ExpListNode)
        var autoBoxingCount LnsInt
        autoBoxingCount = 0
        var hasImplictCast bool
        hasImplictCast = false
        var newExpListNode LnsAny
        if result != Ast_MatchType__Error{
            {
                _workList := self.FP.checkImplicitCast(alt2typeMap, false, dstTypeList, expListNode_9939, TransUnit_checkImplicitCastCallback_1263_(func(dstType *Ast_TypeInfo,expNode *Nodes_Node) LnsAny {
                    if Ast_CanEvalCtrlTypeInfo_canAutoBoxing(dstType, expNode.FP.Get_expType()){
                        autoBoxingCount = autoBoxingCount + 1
                        
                        return &Nodes_BoxingNode_create(self.nodeManager, expNode.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dstType)}), expNode).Nodes_Node
                    }
                    return nil
                }))
                if _workList != nil {
                    workList := _workList.(*Nodes_ExpListNode)
                    newExpListNode = workList
                    
                    hasImplictCast = true
                    
                } else {
                    newExpListNode = nil
                    
                }
            }
        } else { 
            newExpListNode = nil
            
        }
        if autoBoxingCount > 0{
            if Lns_op_not(Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(alt2typeMap, autoBoxingCount)){
                self.FP.addErrMess(pos, Lns_getVM().String_format("illegal auto boxing error -- %d", []LnsAny{autoBoxingCount}))
            }
            return result, alt2typeMap, newExpListNode, expTypeList
        } else if Ast_CanEvalCtrlTypeInfo_hasNeedAutoBoxing(alt2typeMap){
            self.FP.addErrMess(pos, "not support auto boxing")
        }
        if hasImplictCast{
            return result, alt2typeMap, newExpListNode, expTypeList
        }
    }
    if Lns_op_not(Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(alt2typeMap, 0)){
        self.FP.addErrMess(pos, "can't auto boxing error")
    }
    return result, alt2typeMap, nil, expTypeList
}

// 768: decl @lune.@base.@TransUnit.TransUnit.checkMatchValType
func (self *TransUnit_TransUnit) checkMatchValType(pos *Types_Position,funcTypeInfo *Ast_TypeInfo,expList LnsAny,genericTypeList *LnsList,genericsClass LnsAny)(LnsInt, *LnsMap, LnsAny) {
    var argTypeList *LnsList
    argTypeList = funcTypeInfo.FP.Get_argTypeInfoList()
    if funcTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Ext{
        var extTypeList LnsAny
        var mess string
        extTypeList,mess = Ast_convToExtTypeList(self.processInfo, argTypeList)
        if extTypeList != nil{
            extTypeList_9972 := extTypeList.(*LnsList)
            argTypeList = extTypeList_9972
            
        } else {
            self.FP.addErrMess(pos, Lns_getVM().String_format("not support argType on Luaval -- %s", []LnsAny{mess}))
        }
    }
    if _switch56185 := funcTypeInfo; _switch56185 == TransUnit_builtinFunc.List_insert || _switch56185 == TransUnit_builtinFunc.Set_add || _switch56185 == TransUnit_builtinFunc.Set_del {
    } else if _switch56185 == TransUnit_builtinFunc.List_sort {
        _ = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false)
        var callback *Ast_TypeInfo
        callback = self.processInfo.FP.CreateFunc(false, false, nil, Ast_TypeInfoKind__Func, Ast_headTypeInfo, false, false, true, Ast_AccessMode__Pri, "sort", nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()), Ast_TypeInfo2Stem(genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeBool)}), false)
        argTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(callback.FP.Get_nilableTypeInfo())})
        
    } else if _switch56185 == TransUnit_builtinFunc.List_remove {
    }
    var warnForFollow bool
    warnForFollow = true
    if expList != nil{
        expList_9981 := expList.(*Nodes_ExpListNode)
        if expList_9981.FP.Get_followOn(){
            warnForFollow = false
            
        }
    }
    var alt2typeMap *LnsMap
    if funcTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Method{
        if genericsClass != nil{
            genericsClass_9986 := genericsClass.(*Ast_TypeInfo)
            if funcTypeInfo.FP.Get_rawTxt() == "__init"{
                alt2typeMap = genericsClass_9986.FP.CreateAlt2typeMap(true)
                
            } else { 
                if funcTypeInfo.FP.Get_itemTypeInfoList().Len() == 0{
                    alt2typeMap = genericsClass_9986.FP.CreateAlt2typeMap(false)
                    
                } else { 
                    alt2typeMap = genericsClass_9986.FP.CreateAlt2typeMap(true)
                    
                    for _, _itemType := range( genericsClass_9986.FP.Get_itemTypeInfoList().Items ) {
                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( itemType.FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
                            Lns_GetEnv().SetStackVal( Lns_op_not(alt2typeMap.Items[itemType])) ).(bool)){
                            alt2typeMap.Set(itemType,Ast_builtinTypeNone)
                        }
                    }
                }
            }
        } else {
            self.FP.Error("none class")
        }
    } else { 
        alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(funcTypeInfo.FP.Get_itemTypeInfoList().Len() > 0)
        
    }
    var matchResult LnsInt
    var newExpNodeList LnsAny
    matchResult,_,newExpNodeList = TransUnit_convExp56376(Lns_2DDD(self.FP.checkMatchType(funcTypeInfo.FP.GetTxt(nil, nil, nil), pos, argTypeList, expList, false, warnForFollow, alt2typeMap)))
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( expList) &&
        Lns_GetEnv().SetStackVal( newExpNodeList) )){
        return matchResult, alt2typeMap, newExpNodeList
    }
    return matchResult, alt2typeMap, expList
}

// 844: decl @lune.@base.@TransUnit.TransUnit.analyzeListItems
func (self *TransUnit_TransUnit) analyzeListItems(firstPos *Types_Position,nextToken *Types_Token,termTxt string,expectTypeList LnsAny)(LnsAny, *Ast_TypeInfo) {
    var expList LnsAny
    expList = nil
    var itemCommonType LnsAny
    itemCommonType = &Ast_CommonType__Normal{Ast_builtinTypeNone}
    if nextToken.Txt != termTxt{
        self.FP.Pushback()
        expList = self.FP.analyzeExpList(false, false, false, nil, expectTypeList, expectTypeList != nil)
        
        self.FP.checkNextToken(termTxt)
        var nodeList *LnsList
        nodeList = (Lns_unwrap( expList).(*Nodes_ExpListNode)).FP.Get_expList()
        for _, _exp := range( nodeList.Items ) {
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            itemCommonType = Ast_TypeInfo_getCommonTypeCombo(self.processInfo, itemCommonType, &Ast_CommonType__Normal{exp.FP.Get_expType()}, Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false))
            
        }
    }
    var itemTypeInfo *Ast_TypeInfo
    switch _exp56568 := itemCommonType.(type) {
    case *Ast_CommonType__Normal:
    info := _exp56568.Val1
        itemTypeInfo = info
        
    case *Ast_CommonType__Combine:
    info := _exp56568.Val1
        itemTypeInfo = info.FP.Get_typeInfo(self.processInfo)
        
    }
    if itemTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__DDD{
        if itemTypeInfo.FP.Get_itemTypeInfoList().Len() > 0{
            itemTypeInfo = itemTypeInfo.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            
        } else { 
            itemTypeInfo = Ast_builtinTypeStem_
            
        }
    }
    if Lns_op_not(expectTypeList){
        var expTypeList *LnsList
        expTypeList = NewLnsList([]LnsAny{})
        if expList != nil{
            expList_10025 := expList.(*Nodes_ExpListNode)
            for _index, _expNode := range( expList_10025.FP.Get_expList().Items ) {
                index := _index + 1
                expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                if index == expList_10025.FP.Get_expList().Len(){
                    if expNode.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__DDD{
                        expTypeList.Insert(Ast_TypeInfo2Stem(expNode.FP.Get_expType()))
                    } else { 
                        {
                            var _from56676 LnsInt = 1
                            var _to56676 LnsInt = expNode.FP.Get_expTypeList().Len()
                            for _work56676 := _from56676; _work56676 <= _to56676; _work56676++ {
                                expTypeList.Insert(Ast_TypeInfo2Stem(itemTypeInfo))
                            }
                        }
                    }
                } else { 
                    expTypeList.Insert(Ast_TypeInfo2Stem(itemTypeInfo))
                }
            }
        }
        var workExpList LnsAny
        _,_,workExpList = TransUnit_convExp56721(Lns_2DDD(self.FP.checkMatchType("List constructor", firstPos, expTypeList, expList, false, false, nil)))
        if workExpList != nil{
            workExpList_10039 := workExpList.(*Nodes_ExpListNode)
            expList = workExpList_10039
            
        }
    }
    return expList, itemTypeInfo
}

// 913: decl @lune.@base.@TransUnit.TransUnit.analyzeListConst
func (self *TransUnit_TransUnit) analyzeListConst(token *Types_Token,expectType LnsAny) *Nodes_Node {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var expectTypeList LnsAny
    expectTypeList = nil
    if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expectType) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_kind()})) == Ast_TypeInfoKind__List{
        {
            _itemTypeInfoList := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expectType) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_itemTypeInfoList()}))
            if _itemTypeInfoList != nil {
                itemTypeInfoList := _itemTypeInfoList.(*LnsList)
                expectTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
                
            }
        }
    }
    var expList LnsAny
    var itemTypeInfo *Ast_TypeInfo
    expList,itemTypeInfo = self.FP.analyzeListItems(token.Pos, nextToken, "]", expectTypeList)
    var typeInfoList *LnsList
    typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})
    if token.Txt == "["{
        typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.processInfo.FP.CreateList(Ast_AccessMode__Local, self.FP.getCurrentClass(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfo)}), Ast_MutMode__Mut))})
        
        return &Nodes_LiteralListNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), typeInfoList, expList).Nodes_Node
    } else { 
        typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.processInfo.FP.CreateArray(Ast_AccessMode__Local, self.FP.getCurrentClass(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfo)}), Ast_MutMode__Mut))})
        
        return &Nodes_LiteralArrayNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), typeInfoList, expList).Nodes_Node
    }
// insert a dummy
    return nil
}

// 946: decl @lune.@base.@TransUnit.TransUnit.analyzeSetConst
func (self *TransUnit_TransUnit) analyzeSetConst(token *Types_Token,expectType LnsAny) *Nodes_Node {
    self.helperInfo.UseSet = true
    
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var expectTypeList LnsAny
    expectTypeList = nil
    if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expectType) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_kind()})) == Ast_TypeInfoKind__Set{
        {
            _itemTypeInfoList := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expectType) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_itemTypeInfoList()}))
            if _itemTypeInfoList != nil {
                itemTypeInfoList := _itemTypeInfoList.(*LnsList)
                expectTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
                
            }
        }
    }
    var expList LnsAny
    var itemTypeInfo *Ast_TypeInfo
    expList,itemTypeInfo = self.FP.analyzeListItems(token.Pos, nextToken, ")", expectTypeList)
    if itemTypeInfo.FP.Get_nilable(){
        if expList != nil{
            expList_10071 := expList.(*Nodes_ExpListNode)
            for _, _exp := range( expList_10071.FP.Get_expList().Items ) {
                exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
                var expType *Ast_TypeInfo
                expType = exp.FP.Get_expType()
                if expType.FP.Get_nilable(){
                    self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("'Set' object can't store nilable. -- %s", []LnsAny{expType.FP.GetTxt(nil, nil, nil)}))
                }
            }
        }
    }
    var typeInfoList *LnsList
    typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})
    typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.processInfo.FP.CreateSet(Ast_AccessMode__Local, self.FP.getCurrentClass(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfo)}), Ast_MutMode__Mut))})
    
    return &Nodes_LiteralSetNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), typeInfoList, expList).Nodes_Node
}

// 986: decl @lune.@base.@TransUnit.TransUnit.analyzeMapConst
func (self *TransUnit_TransUnit) analyzeMapConst(token *Types_Token,expectType LnsAny) *Nodes_LiteralMapNode {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var _map *LnsMap
    _map = NewLnsMap( map[LnsAny]LnsAny{})
    var pairList *LnsList
    pairList = NewLnsList([]LnsAny{})
    var keyTypeInfo *Ast_TypeInfo
    keyTypeInfo = Ast_builtinTypeNone
    var valTypeInfo *Ast_TypeInfo
    valTypeInfo = Ast_builtinTypeNone
    var getMapKeyValType func(pos *Types_Position,keyFlag bool,typeInfo *Ast_TypeInfo,expType *Ast_TypeInfo) *Ast_TypeInfo
    getMapKeyValType = func(pos *Types_Position,keyFlag bool,typeInfo *Ast_TypeInfo,expType *Ast_TypeInfo) *Ast_TypeInfo {
        if expType.FP.Get_nilable(){
            if keyFlag{
                self.FP.addErrMess(pos, Lns_getVM().String_format("map key can't set a nilable -- %s", []LnsAny{expType.FP.GetTxt(nil, nil, nil)}))
            }
            if expType.FP.Equals(self.processInfo, Ast_builtinTypeNil, nil, nil){
                return typeInfo
            }
            expType = expType.FP.Get_nonnilableType()
            
        }
        return Ast_TypeInfo_getCommonType(self.processInfo, typeInfo, expType, Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false))
    }
    var expectKeyType LnsAny
    expectKeyType = nil
    var expectValType LnsAny
    expectValType = nil
    if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expectType) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_kind()})) == Ast_TypeInfoKind__Map{
        {
            _itemTypeInfoList := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expectType) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_itemTypeInfoList()}))
            if _itemTypeInfoList != nil {
                itemTypeInfoList := _itemTypeInfoList.(*LnsList)
                expectKeyType = itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                
                expectValType = itemTypeInfoList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                
            }
        }
    }
    for  {
        if nextToken.Txt == "}"{
            break
        }
        self.FP.Pushback()
        var key *Nodes_Node
        key = self.FP.analyzeExpOneRVal(false, false, nil, expectKeyType)
        keyTypeInfo = getMapKeyValType(key.FP.Get_pos(), true, keyTypeInfo, key.FP.Get_expType())
        
        self.FP.checkNextToken(":")
        var val *Nodes_Node
        val = self.FP.analyzeExpOneRVal(false, false, nil, expectValType)
        valTypeInfo = getMapKeyValType(val.FP.Get_pos(), false, valTypeInfo, val.FP.Get_expType())
        
        pairList.Insert(Nodes_PairItem2Stem(NewNodes_PairItem(key, val)))
        _map.Set(key,val)
        nextToken = self.FP.getToken(nil)
        
        if nextToken.Txt != ","{
            break
        }
        nextToken = self.FP.getToken(nil)
        
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = self.processInfo.FP.CreateMap(Ast_AccessMode__Local, self.FP.getCurrentClass(), keyTypeInfo, valTypeInfo, Ast_MutMode__Mut)
    self.FP.checkToken(nextToken, "}")
    return Nodes_LiteralMapNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), _map, pairList)
}

// 1074: decl @lune.@base.@TransUnit.TransUnit.evalMacroOp
func (self *TransUnit_TransUnit) evalMacroOp(firstToken *Types_Token,macroTypeInfo *Ast_TypeInfo,expList LnsAny,evalMacroCallback Macro_EvalMacroCallback) {
    var parser LnsAny
    var mess LnsAny
    parser,mess = self.macroCtrl.FP.EvalMacroOp(self.parser.FP.GetStreamName(), firstToken, macroTypeInfo, expList)
    var bakParser *Parser_DefaultPushbackParser
    bakParser = self.parser
    if parser != nil{
        parser_10121 := parser.(*Parser_Parser)
        self.parser = NewParser_DefaultPushbackParser(parser_10121)
        
    } else {
        self.FP.Error(Lns_unwrap( mess).(string))
    }
    self.macroCtrl.FP.StartExpandMode(firstToken.Pos.LineNo, macroTypeInfo, evalMacroCallback)
    var nextToken *Types_Token
    nextToken = self.FP.GetTokenNoErr()
    self.parser = bakParser
    
    if nextToken != Parser_getEofToken(){
        self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("remain macro expand-statement token -- '%s'(%d:%d)", []LnsAny{nextToken.Txt, nextToken.Pos.LineNo, nextToken.Pos.Column}))
        if Lns_op_not(macroTypeInfo.FP.Get_externalFlag()){
            self.FP.addErrMess(nextToken.Pos, Lns_getVM().String_format("remain macro expand-statement token -- '%s'", []LnsAny{nextToken.Txt}))
        }
    }
}

// 1113: decl @lune.@base.@TransUnit.TransUnit.evalMacro
func (self *TransUnit_TransUnit) evalMacro(firstToken *Types_Token,macroTypeInfo *Ast_TypeInfo,expList LnsAny) *Nodes_ExpMacroExpNode {
    var stmtList *LnsList
    stmtList = NewLnsList([]LnsAny{})
    self.FP.evalMacroOp(firstToken, macroTypeInfo, expList, Macro_EvalMacroCallback(func() {
        if macroTypeInfo.FP.Get_retTypeInfoList().Len() == 0{
            self.FP.analyzeStatementList(stmtList, "}")
        } else { 
            stmtList.Insert(Nodes_Node2Stem(self.FP.analyzeExp(false, false, false, nil, nil)))
        }
    }))
    var expTypeList *LnsList
    expTypeList = macroTypeInfo.FP.Get_retTypeInfoList()
    if macroTypeInfo.FP.Get_retTypeInfoList().Len() > 0{
        var macroRetTypeList *LnsList
        macroRetTypeList = macroTypeInfo.FP.Get_retTypeInfoList()
        if stmtList.Len() == 1{
            var node *Nodes_Node
            node = stmtList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
            var retType *Ast_TypeInfo
            retType = macroRetTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if retType.FP.Equals(self.processInfo, Ast_builtinTypeMultiExp, nil, nil){
                expTypeList = node.FP.Get_expTypeList()
                
            } else if retType.FP.Equals(self.processInfo, Ast_builtinTypeExp, nil, nil){
                if node.FP.Get_expTypeList().Len() == 1{
                    expTypeList = node.FP.Get_expTypeList()
                    
                } else { 
                    self.FP.addErrMess(firstToken.Pos, "__exp can't return multiple values. use __exps.")
                }
            } else if node.FP.Get_expTypeList().Len() == 1{
                if retType.FP.Equals(self.processInfo, node.FP.Get_expType(), nil, nil){
                    expTypeList = node.FP.Get_expTypeList()
                    
                } else { 
                    self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("mismatch type -- %s != %s", []LnsAny{macroRetTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.GetTxt(nil, nil, nil), node.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
                }
            } else { 
                self.FP.addErrMess(firstToken.Pos, "macro can't return multiple values.")
            }
        } else { 
            self.FP.addErrMess(firstToken.Pos, "macro to return value must be one statemnt.")
        }
    } else { 
        expTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})
        
    }
    return Nodes_ExpMacroExpNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), expTypeList, stmtList)
}

// 1253: decl @lune.@base.@TransUnit.TransUnit.checkStringFormat
func (self *TransUnit_TransUnit) checkStringFormat(pos *Types_Position,formatTxt string,argTypeList *LnsList) {
    var opList *LnsList
    opList = TransUnit_findForm(formatTxt)
    if opList.Len() != argTypeList.Len(){
        self.FP.addErrMess(pos, Lns_getVM().String_format("argument number is mismatch -- %d != %d (%s)", []LnsAny{opList.Len(), argTypeList.Len(), Lns_getVM().String_sub(formatTxt,1, 20)}))
        return 
    }
    for _index, _op := range( opList.Items ) {
        index := _index + 1
        op := _op.(string)
        var argType *Ast_TypeInfo
        argType = argTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var match LnsInt
        var reqType *Ast_TypeInfo
        match,reqType = TransUnit_isMatchStringFormatType(op, argType, self.targetLuaVer)
        if match == TransUnit_FormType__Unmatch{
            var mess string
            mess = Lns_getVM().String_format("type must be %s -- %s", []LnsAny{reqType.FP.GetTxt(nil, nil, nil), argType.FP.GetTxt(nil, nil, nil)})
            self.FP.addErrMess(pos, Lns_getVM().String_format("argument(%d) %s", []LnsAny{index, mess}))
        }
    }
}

// 1292: decl @lune.@base.@TransUnit.TransUnit.prepareExpCall
func (self *TransUnit_TransUnit) prepareExpCall(position *Types_Position,funcTypeInfo *Ast_TypeInfo,genericTypeList *LnsList,genericsClass *Ast_TypeInfo)(*LnsMap, LnsAny) {
    if funcTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Macro{
        self.macroCtrl.FP.StartAnalyzeArgMode(funcTypeInfo)
    }
    var work *Types_Token
    work = self.FP.getToken(nil)
    var argList LnsAny
    argList = nil
    if work.Txt != ")"{
        self.FP.Pushback()
        argList = self.FP.analyzeExpList(false, false, false, nil, funcTypeInfo.FP.Get_argTypeInfoList(), nil)
        
        self.FP.checkNextToken(")")
        if argList != nil{
            argList_10229 := argList.(*Nodes_ExpListNode)
            for _, _argNode := range( argList_10229.FP.Get_expList().Items ) {
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_op_not(argNode.FP.CanBeRight(self.processInfo))) &&
                    Lns_GetEnv().SetStackVal( argNode.FP.Get_kind() != Nodes_NodeKind_get_Abbr()) ).(bool)){
                    self.FP.addErrMess(argNode.FP.Get_pos(), Lns_getVM().String_format("this node can't be r-value. -- %s", []LnsAny{Nodes_getNodeKindName(argNode.FP.Get_kind())}))
                }
            }
        }
    }
    var matchResult LnsInt
    var alt2typeMap *LnsMap
    var workArgList LnsAny
    matchResult,alt2typeMap,workArgList = self.FP.checkMatchValType(position, funcTypeInfo, argList, genericTypeList, genericsClass)
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Macro) &&
        Lns_GetEnv().SetStackVal( matchResult == Ast_MatchType__Error) ).(bool)){
        self.FP.Error(Lns_getVM().String_format("unmatch macro arguments. -- %s", []LnsAny{funcTypeInfo.FP.GetTxt(nil, nil, nil)}))
    }
    if funcTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Macro{
        self.macroCtrl.FP.FinishMacroMode()
    }
    return alt2typeMap, workArgList
}

// 1341: decl @lune.@base.@TransUnit.TransUnit.checkArgForStringForm
func (self *TransUnit_TransUnit) checkArgForStringForm(firstToken *Types_Token,argList *Nodes_ExpListNode) {
    var formArgTypeList *LnsList
    formArgTypeList = NewLnsList([]LnsAny{})
    var formatTxt string
    formatTxt = ""
    if argList.FP.Get_expList().Len() > 0{
        var argNode *Nodes_Node
        argNode = argList.FP.Get_expList().GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
        if argNode.FP.Get_kind() != Nodes_NodeKind_get_LiteralString(){
            return 
        }
        {
            _literal := TransUnit_convExp58829(Lns_2DDD(argNode.FP.GetLiteral()))
            if _literal != nil {
                literal := _literal
                switch _exp58827 := literal.(type) {
                case *Nodes_Literal__Str:
                val := _exp58827.Val1
                    formatTxt = val
                    
                }
            }
        }
    }
    if argList.FP.Get_expList().Len() > 1{
        {
            _toDDDNode := Nodes_ExpToDDDNodeDownCastF(argList.FP.Get_expList().GetAt(2).(Nodes_NodeDownCast).ToNodes_Node().FP)
            if _toDDDNode != nil {
                toDDDNode := _toDDDNode.(*Nodes_ExpToDDDNode)
                for _, _workNode := range( toDDDNode.FP.Get_expList().FP.Get_expList().Items ) {
                    workNode := _workNode.(Nodes_NodeDownCast).ToNodes_Node()
                    formArgTypeList.Insert(Ast_TypeInfo2Stem(workNode.FP.Get_expType()))
                }
            } else {
                self.FP.Error(Lns_getVM().String_format("illegal node -- %s", []LnsAny{Nodes_getNodeKindName(argList.FP.Get_expList().GetAt(2).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_kind())}))
            }
        }
    }
    self.FP.checkStringFormat(firstToken.Pos, formatTxt, formArgTypeList)
}

// 1378: decl @lune.@base.@TransUnit.TransUnit.checkArgForSort
func (self *TransUnit_TransUnit) checkArgForSort(firstToken *Types_Token,genericTypeList *LnsList,argList *Nodes_ExpListNode) {
    if argList.FP.Get_expTypeList().Len() > 0{
        var callback *Ast_TypeInfo
        callback = argList.FP.Get_expTypeList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if callback == Ast_builtinTypeAbbr{
            return 
        }
        if callback.FP.Get_retTypeInfoList().Len() != 1{
            self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("The callback's to return value of sort() must have a value. -- %d", []LnsAny{callback.FP.Get_retTypeInfoList().Len()}))
            return 
        }
        if Lns_op_not(Ast_builtinTypeBool.FP.Equals(self.processInfo, callback.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), nil, nil)){
            self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("The callback's return type of sort() must be bool. -- '%s'", []LnsAny{callback.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.GetTxt(nil, nil, nil)}))
        }
        if callback.FP.Get_argTypeInfoList().Len() != 2{
            self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("The callback's argument must have 2 arguments. -- '%s'", []LnsAny{callback.FP.Get_display_stirng()}))
        }
        if genericTypeList.Len() == 1{
            for _index, _argType := range( callback.FP.Get_argTypeInfoList().Items ) {
                index := _index + 1
                argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if Lns_op_not(genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Equals(self.processInfo, argType, nil, nil)){
                    self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("The callback's argument(%d) type must be -- '%s'", []LnsAny{index, genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.GetTxt(nil, nil, nil)}))
                }
            }
        } else { 
            self.FP.addErrMess(firstToken.Pos, "The generics of the list is illegal")
        }
    }
}

// 1428: decl @lune.@base.@TransUnit.TransUnit.processFunc
func (self *TransUnit_TransUnit) processFunc(firstToken *Types_Token,nextToken *Types_Token,refFieldNode LnsAny,funcExp *Nodes_Node,funcType *Ast_TypeInfo,alt2typeMap *LnsMap,genericTypeList *LnsList,genericsClass *Ast_TypeInfo,argList LnsAny) *Nodes_Node {
    var funcSymbol LnsAny
    var symbolInfoList *LnsList
    symbolInfoList = funcExp.FP.GetSymbolInfo()
    if symbolInfoList.Len() > 0{
        var symbol *Ast_SymbolInfo
        symbol = symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if symbol.FP.Get_kind() == Ast_SymbolKind__Typ{
            self.FP.addErrMess(funcExp.FP.Get_pos(), Lns_getVM().String_format("can't call any Type. -- %s", []LnsAny{symbol.FP.Get_name()}))
        }
        funcSymbol = symbol
        
    } else { 
        funcSymbol = nil
        
    }
    var funcTypeInfo *Ast_TypeInfo
    funcTypeInfo = funcExp.FP.Get_expType().FP.Get_srcTypeInfo()
    var nilAccess bool
    if nextToken.Txt == "$("{
        if funcTypeInfo.FP.Get_nilable(){
            funcTypeInfo = funcTypeInfo.FP.Get_nonnilableType()
            
            nilAccess = true
            
        } else { 
            self.FP.addWarnMess(funcExp.FP.Get_pos(), Lns_getVM().String_format("This is not nilable. -- %s", []LnsAny{funcTypeInfo.FP.GetTxt(nil, nil, nil)}))
            nilAccess = false
            
        }
    } else { 
        nilAccess = false
        
    }
    if _switch59415 := (funcTypeInfo.FP.Get_kind()); _switch59415 == Ast_TypeInfoKind__Method || _switch59415 == Ast_TypeInfoKind__Func || _switch59415 == Ast_TypeInfoKind__Form || _switch59415 == Ast_TypeInfoKind__FormFunc {
    } else {
        {
            _extType := Ast_ExtTypeInfoDownCastF(funcTypeInfo.FP)
            if _extType != nil {
                extType := _extType.(*Ast_ExtTypeInfo)
                if _switch59387 := (extType.FP.Get_extedType().FP.Get_kind()); _switch59387 == Ast_TypeInfoKind__Method || _switch59387 == Ast_TypeInfoKind__Func || _switch59387 == Ast_TypeInfoKind__Form || _switch59387 == Ast_TypeInfoKind__FormFunc {
                } else {
                    self.FP.Error(Lns_getVM().String_format("can't call the type -- %s, %s", []LnsAny{funcTypeInfo.FP.GetTxt(nil, nil, nil), Ast_TypeInfoKind_getTxt( funcTypeInfo.FP.Get_kind())}))
                }
            } else {
                self.FP.Error(Lns_getVM().String_format("can't call the type -- %s, %s", []LnsAny{funcTypeInfo.FP.GetTxt(nil, nil, nil), Ast_TypeInfoKind_getTxt( funcTypeInfo.FP.Get_kind())}))
            }
        }
    }
    var retTypeInfoList *LnsList
    retTypeInfoList = NewLnsList([]LnsAny{})
    for _index, _retType := range( funcTypeInfo.FP.Get_retTypeInfoList().Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        retTypeInfoList.Insert(Ast_TypeInfo2Stem(retType))
        {
            _applyType := retType.FP.ApplyGeneric(alt2typeMap, self.moduleType)
            if _applyType != nil {
                applyType := _applyType.(*Ast_TypeInfo)
                retTypeInfoList.Set(index,applyType)
            } else {
                if funcTypeInfo == TransUnit_builtinFunc.List_remove{
                    retTypeInfoList.Set(index,genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo())
                } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Func) &&
                    Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_rawTxt() == "_fromMap") ||
                        Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_rawTxt() == "_fromStem") ).(bool))) &&
                    Lns_GetEnv().SetStackVal( genericsClass.FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, alt2typeMap)) ).(bool)){
                    retTypeInfoList.Set(index,genericsClass.FP.Get_nilableTypeInfo())
                } else { 
                    self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("not support generics yet. -- %s", []LnsAny{retType.FP.GetTxt(nil, nil, nil)}))
                }
            }
        }
    }
    if refFieldNode != nil{
        refFieldNode_10321 := refFieldNode.(*Nodes_RefFieldNode)
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Equals(self.processInfo, TransUnit_builtinFunc.List_unpack, nil, nil)) ||
            Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Equals(self.processInfo, TransUnit_builtinFunc.Array_unpack, nil, nil)) ).(bool){
            var prefixType *Ast_TypeInfo
            prefixType = refFieldNode_10321.FP.Get_prefix().FP.Get_expType()
            if prefixType.FP.Get_itemTypeInfoList().Len() > 0{
                var dddType *Ast_DDDTypeInfo
                dddType = self.processInfo.FP.CreateDDD(prefixType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), false, false)
                retTypeInfoList = NewLnsList([]LnsAny{})
                
                retTypeInfoList.Insert(Ast_DDDTypeInfo2Stem(dddType))
            }
        }
    }
    if nilAccess{
        var retList *LnsList
        retList = NewLnsList([]LnsAny{})
        for _, _retType := range( retTypeInfoList.Items ) {
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if retType.FP.Get_nilable(){
                retList.Insert(Ast_TypeInfo2Stem(retType))
            } else { 
                retList.Insert(Ast_TypeInfo2Stem(retType.FP.Get_nilableTypeInfo()))
            }
        }
        retTypeInfoList = retList
        
        self.helperInfo.UseNilAccess = true
        
    }
    var errorFuncFlag bool
    errorFuncFlag = false
    if retTypeInfoList.Len() > 0{
        var retType *Ast_TypeInfo
        retType = retTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if retType == Ast_builtinTypeNeverRet{
            errorFuncFlag = true
            
        }
    }
    if argList != nil{
        argList_10337 := argList.(*Nodes_ExpListNode)
        if _switch59787 := funcTypeInfo; _switch59787 == TransUnit_builtinFunc.String_format {
            self.FP.checkArgForStringForm(firstToken, argList_10337)
        } else if _switch59787 == TransUnit_builtinFunc.List_sort || _switch59787 == TransUnit_builtinFunc.Array_sort {
            self.FP.checkArgForSort(firstToken, genericTypeList, argList_10337)
        }
    }
    if funcTypeInfo.FP.Equals(self.processInfo, TransUnit_builtinFunc.Lns__kind, nil, nil){
        {
            _expList := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(argList) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_ExpListNode).FP.Get_expList()}))
            if _expList != nil {
                expList := _expList.(*LnsList)
                if expList.Len() > 0{
                    return &Nodes_LuneKindNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), expList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()).Nodes_Node
                }
            }
        }
        return &Nodes_LuneKindNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), self.FP.createNoneNode(firstToken.Pos)).Nodes_Node
    }
    if funcSymbol != nil{
        funcSymbol_10345 := funcSymbol.(*Ast_SymbolInfo)
        if funcSymbol_10345.FP.Get_name() == "super"{
            return &Nodes_ExpCallSuperNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), retTypeInfoList, funcSymbol_10345.FP.Get_typeInfo().FP.Get_parentInfo(), funcSymbol_10345.FP.Get_typeInfo(), argList).Nodes_Node
        }
    }
    if funcType.FP.Get_kind() == Ast_TypeInfoKind__Ext{
        var work LnsAny
        var err string
        work,err = Ast_convToExtTypeList(self.processInfo, retTypeInfoList)
        if work != nil{
            work_10351 := work.(*LnsList)
            retTypeInfoList = work_10351
            
        } else {
            self.FP.addErrMess(firstToken.Pos, err)
        }
    }
    var threading bool
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( nilAccess) ||
        Lns_GetEnv().SetStackVal( funcType.FP.Get_kind() == Ast_TypeInfoKind__Ext) ||
        Lns_GetEnv().SetStackVal( TransUnit_builtinFunc.FP.Get_needThreadingTypes().Has(Ast_TypeInfo2Stem(funcType.FP.Get_nonnilableType().FP.Get_srcTypeInfo()))) ).(bool){
        threading = self.FP.checkThreading(firstToken.Pos)
        
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( threading) &&
            Lns_GetEnv().SetStackVal( funcType.FP.Get_kind() == Ast_TypeInfoKind__Ext) ).(bool)){
            self.FP.addErrMess(firstToken.Pos, "not support to use Luaval on thread.")
        }
    } else { 
        threading = false
        
    }
    return &Nodes_ExpCallNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), retTypeInfoList, funcExp, errorFuncFlag, nilAccess, threading, argList).Nodes_Node
}

// 1625: decl @lune.@base.@TransUnit.TransUnit.analyzeExpCall
func (self *TransUnit_TransUnit) analyzeExpCall(firstToken *Types_Token,funcExp *Nodes_Node,nextToken *Types_Token)(*Nodes_Node, *Types_Token) {
    self.FP.checkSymbolHavingValue(funcExp.FP.Get_pos(), funcExp.FP.GetSymbolInfo())
    var funcTypeInfo *Ast_TypeInfo
    funcTypeInfo = funcExp.FP.Get_expType().FP.Get_nonnilableType()
    var genericTypeList *LnsList
    genericTypeList = funcTypeInfo.FP.Get_itemTypeInfoList()
    var refFieldNode LnsAny
    refFieldNode = nil
    var genericsClass *Ast_TypeInfo
    genericsClass = Ast_headTypeInfo
    {
        _refField := Nodes_RefFieldNodeDownCastF(funcExp.FP)
        if _refField != nil {
            refField := _refField.(*Nodes_RefFieldNode)
            refFieldNode = refField
            
            var classType *Ast_TypeInfo
            classType = refField.FP.Get_prefix().FP.Get_expType()
            genericsClass = classType
            
            if funcTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Method{
                genericTypeList = classType.FP.Get_itemTypeInfoList()
                
            }
        }
    }
    var alt2typeMap *LnsMap
    var argList LnsAny
    alt2typeMap,argList = self.FP.prepareExpCall(funcExp.FP.Get_pos(), funcTypeInfo, genericTypeList, genericsClass)
    if funcTypeInfo.FP.Equals(self.processInfo, TransUnit_builtinFunc.List_insert, nil, nil){
        if argList != nil{
            argList_10376 := argList.(*Nodes_ExpListNode)
            if argList_10376.FP.Get_expType().FP.Get_nilable(){
                self.FP.addErrMess(argList_10376.FP.Get_pos(), "list can't insert nilable")
            }
        }
    }
    if funcTypeInfo.FP.Equals(self.processInfo, TransUnit_builtinFunc.Set_add, nil, nil){
        if argList != nil{
            argList_10380 := argList.(*Nodes_ExpListNode)
            if argList_10380.FP.Get_expType().FP.Get_nilable(){
                self.FP.addErrMess(argList_10380.FP.Get_pos(), "set can't add nilable")
            }
        }
    } else if funcTypeInfo.FP.Equals(self.processInfo, TransUnit_builtinFunc.List_remove, nil, nil){
        if genericTypeList.Len() > 0{
            if genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilable(){
                self.FP.addWarnMess(funcExp.FP.Get_pos(), "remove() is dangerous for nilable's list.")
            }
        }
    }
    if funcTypeInfo.FP.Get_rawTxt() == ""{
        self.FP.addErrMess(funcExp.FP.Get_pos(), "can't directly call the declared function.")
    }
    var exp *Nodes_Node
    if funcTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Macro{
        exp = &self.FP.evalMacro(firstToken, funcTypeInfo, argList).Nodes_Node
        
        if self.FP.checkThreading(firstToken.Pos){
            self.FP.addErrMess(firstToken.Pos, "not support to use a macro")
        }
    } else { 
        exp = self.FP.processFunc(firstToken, nextToken, refFieldNode, funcExp, funcTypeInfo, alt2typeMap, genericTypeList, genericsClass, argList)
        
    }
    return exp, self.FP.GetTokenNoErr()
}

// 1700: decl @lune.@base.@TransUnit.TransUnit.analyzeExpCast
func (self *TransUnit_TransUnit) analyzeExpCast(firstToken *Types_Token,opTxt string,exp *Nodes_Node) *Nodes_Node {
    var castTypeNode *Nodes_RefTypeNode
    castTypeNode = self.FP.analyzeRefType(Ast_AccessMode__Local, false, false)
    var castType *Ast_TypeInfo
    castType = castTypeNode.FP.Get_expType()
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Ext) &&
        Lns_GetEnv().SetStackVal( castType.FP.Get_kind() != Ast_TypeInfoKind__Ext) &&
        Lns_GetEnv().SetStackVal( castType.FP.Get_kind() != Ast_TypeInfoKind__Stem) ).(bool)){
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( opTxt == "@@") ||
            Lns_GetEnv().SetStackVal( opTxt == "@@=") ).(bool){
            castType = self.FP.CreateModifier(castType, Ast_MutMode__IMut)
            
        }
        castType = self.FP.createExtType(castTypeNode.FP.Get_pos(), castType)
        
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( castType.FP.Get_kind() == Ast_TypeInfoKind__Form) &&
        Lns_GetEnv().SetStackVal( exp.FP.Get_expType().FP.Get_nonnilableType().FP.Get_kind() == Ast_TypeInfoKind__Stem) ).(bool)){
        if self.FP.supportLang(LuneControl_Code__C){
            self.FP.addWarnMess(castTypeNode.FP.Get_pos(), "not support cast from stem to form for transcompiling to c-lang.")
        }
    }
    var expType *Ast_TypeInfo
    expType = exp.FP.Get_expType()
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( opTxt == "@@@") ||
        Lns_GetEnv().SetStackVal( opTxt == "@@=") ).(bool){
        if castType.FP.Get_itemTypeInfoList().Len() > 0{
            self.FP.addErrMess(castTypeNode.FP.Get_pos(), Lns_getVM().String_format("not support cast for generics class yet -- %s", []LnsAny{castType.FP.GetTxt(nil, nil, nil)}))
        }
        if _switch60663 := castType.FP.Get_kind(); _switch60663 == Ast_TypeInfoKind__IF || _switch60663 == Ast_TypeInfoKind__Class || _switch60663 == Ast_TypeInfoKind__Prim {
        } else {
            if opTxt != "@@="{
                self.FP.addErrMess(castTypeNode.FP.Get_pos(), Lns_getVM().String_format("not support cast -- %s", []LnsAny{castType.FP.GetTxt(nil, nil, nil)}))
            }
        }
        if opTxt == "@@="{
            var orgExpType *Ast_TypeInfo
            orgExpType = expType.FP.Get_extedType()
            var orgCastType *Ast_TypeInfo
            orgCastType = castType.FP.Get_extedType()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( orgCastType.FP.Get_kind() != Ast_TypeInfoKind__IF) &&
                Lns_GetEnv().SetStackVal( orgCastType.FP.Get_kind() != Ast_TypeInfoKind__Class) ).(bool)){
                self.FP.addErrMess(castTypeNode.FP.Get_pos(), Lns_getVM().String_format("'@@=' cast must be class or interface. -- %s", []LnsAny{castType.FP.GetTxt(nil, nil, nil)}))
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( orgExpType.FP.Get_srcTypeInfo() != Ast_builtinTypeStem) &&
                Lns_GetEnv().SetStackVal( orgExpType.FP.Get_kind() != Ast_TypeInfoKind__IF) &&
                Lns_GetEnv().SetStackVal( orgExpType.FP.Get_kind() != Ast_TypeInfoKind__Class) ).(bool)){
                self.FP.addErrMess(castTypeNode.FP.Get_pos(), Lns_getVM().String_format("'@@=' cast must be class or interface. -- %s", []LnsAny{castType.FP.GetTxt(nil, nil, nil)}))
            }
            if Lns_op_not(Ast_isStruct(orgCastType)){
                self.FP.addErrMess(castTypeNode.FP.Get_pos(), Lns_getVM().String_format("'@@=' cast type can't use class has method -- %s", []LnsAny{castType.FP.GetTxt(nil, nil, nil)}))
            }
        }
    } else { 
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( castType != Ast_builtinTypeString) &&
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( castType.FP.Get_kind() == Ast_TypeInfoKind__IF) ||
                Lns_GetEnv().SetStackVal( castType.FP.Get_kind() == Ast_TypeInfoKind__Class) ).(bool))) ).(bool)){
            self.FP.addWarnMess(castTypeNode.FP.Get_pos(), Lns_getVM().String_format("use '@@@' cast for class or interface. -- %s", []LnsAny{castType.FP.GetTxt(nil, nil, nil)}))
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( opTxt != "@@@") &&
        Lns_GetEnv().SetStackVal( expType.FP.Get_nilable()) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(castType.FP.Get_nilable())) ).(bool)){
        self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("can't cast from nilable to not nilable  -- %s->%s", []LnsAny{expType.FP.GetTxt(nil, nil, nil), castType.FP.GetTxt(nil, nil, nil)}))
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(expType))) &&
        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(castType)) ).(bool)){
        castType = self.FP.CreateModifier(castType, Ast_MutMode__IMut)
        
    }
    if Lns_car(castType.FP.CanEvalWith(self.processInfo, expType, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( castType.FP.Get_kind() != Ast_TypeInfoKind__Ext) &&
            Lns_GetEnv().SetStackVal( Lns_op_not((Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( castType.FP.Get_kind() == Ast_TypeInfoKind__Stem) &&
                Lns_GetEnv().SetStackVal( Ast_isExtType(expType)) ).(bool)))) ).(bool)){
            self.FP.addWarnMess(castTypeNode.FP.Get_pos(), Lns_getVM().String_format("This cast isn't need. (%s <- %s)", []LnsAny{castType.FP.GetTxt(self.typeNameCtrl, nil, nil), expType.FP.GetTxt(self.typeNameCtrl, nil, nil)}))
        }
    } else if Lns_op_not(Lns_car(expType.FP.CanEvalWith(self.processInfo, castType, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)){
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isNumberType(expType))) ||
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isNumberType(castType))) ).(bool){
            self.FP.addErrMess(castTypeNode.FP.Get_pos(), Lns_getVM().String_format("This type can't cast. (%s <- %s)", []LnsAny{castType.FP.GetTxt(self.typeNameCtrl, nil, nil), expType.FP.GetTxt(self.typeNameCtrl, nil, nil)}))
        }
    }
    if opTxt == "@@@"{
        castType = castType.FP.Get_nilableTypeInfo()
        
    }
    return &Nodes_ExpCastNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(castType)}), self.nodeManager.FP.MultiTo1(exp), castType, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( opTxt != "@@@") &&
        Lns_GetEnv().SetStackVal( Nodes_CastKind__Force) ||
        Lns_GetEnv().SetStackVal( Nodes_CastKind__Normal) ).(LnsInt)).Nodes_Node
}

// 1826: decl @lune.@base.@TransUnit.TransUnit.analyzeExpCont
func (self *TransUnit_TransUnit) analyzeExpCont(firstToken *Types_Token,exp *Nodes_Node,skipFlag bool,canLeftExp bool) *Nodes_Node {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(true)
    if nextToken.Kind == Types_TokenKind__Eof{
        return exp
    }
    if Lns_op_not(skipFlag){
        for {
            var matchFlag bool
            matchFlag = false
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( nextToken.Txt == "[") ||
                Lns_GetEnv().SetStackVal( nextToken.Txt == "$[") ).(bool){
                matchFlag = true
                
                exp = self.FP.analyzeExpRefItem(nextToken, exp, nextToken.Txt == "$[")
                
                nextToken = self.FP.getToken(nil)
                
            }
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( nextToken.Txt == "(") ||
                Lns_GetEnv().SetStackVal( nextToken.Txt == "$(") ).(bool){
                matchFlag = true
                
                exp, nextToken = self.FP.analyzeExpCall(firstToken, exp, nextToken)
                
            }
            if Lns_op_not(matchFlag){ break }
        }
        if _switch61346 := nextToken.Txt; _switch61346 == "@@" || _switch61346 == "@@@" || _switch61346 == "@@=" {
            exp = self.FP.analyzeExpCast(firstToken, nextToken.Txt, exp)
            
            nextToken = self.FP.getToken(nil)
            
        }
    }
    if _switch61466 := nextToken.Txt; _switch61466 == "." {
        return self.FP.analyzeExpSymbol(firstToken, self.FP.getToken(nil), TransUnit_ExpSymbolMode__Field, exp, skipFlag, canLeftExp)
    } else if _switch61466 == "$." {
        return self.FP.analyzeExpSymbol(firstToken, self.FP.getToken(nil), TransUnit_ExpSymbolMode__FieldNil, exp, skipFlag, canLeftExp)
    } else if _switch61466 == ".$" {
        return self.FP.analyzeExpSymbol(firstToken, self.FP.getToken(nil), TransUnit_ExpSymbolMode__Get, exp, skipFlag, canLeftExp)
    } else if _switch61466 == "$.$" {
        return self.FP.analyzeExpSymbol(firstToken, self.FP.getToken(nil), TransUnit_ExpSymbolMode__GetNil, exp, skipFlag, canLeftExp)
    }
    self.FP.Pushback()
    return exp
}

// 1883: decl @lune.@base.@TransUnit.TransUnit.analyzeAccessClassField
func (self *TransUnit_TransUnit) analyzeAccessClassField(classTypeInfo *Ast_TypeInfo,mode LnsInt,token *Types_Token)(*Ast_TypeInfo, LnsAny, bool) {
    if _switch61541 := classTypeInfo.FP.Get_kind(); _switch61541 == Ast_TypeInfoKind__List {
        classTypeInfo = Ast_builtinTypeList
        
    } else if _switch61541 == Ast_TypeInfoKind__Array {
        classTypeInfo = Ast_builtinTypeArray
        
    } else if _switch61541 == Ast_TypeInfoKind__Set {
        classTypeInfo = Ast_builtinTypeSet
        
    }
    var className string
    className = classTypeInfo.FP.GetTxt(nil, nil, nil)
    var classScope *Ast_Scope
    
    {
        _classScope := classTypeInfo.FP.Get_scope()
        if _classScope == nil{
            self.FP.Error(Lns_getVM().String_format("not found field: %s, %s", []LnsAny{className, classTypeInfo}))
        } else {
            classScope = _classScope.(*Ast_Scope)
        }
    }
    var symbolInfo LnsAny
    symbolInfo = nil
    var fieldTypeInfo LnsAny
    fieldTypeInfo = nil
    var getterFlag bool
    getterFlag = false
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__Get) ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__GetNil) ).(bool){
        var fieldSymbolInfo LnsAny
        fieldSymbolInfo = classScope.FP.GetSymbolInfo(Lns_getVM().String_format("get_%s", []LnsAny{token.Txt}), self.scope, false, self.scopeAccess)
        if fieldSymbolInfo != nil{
            fieldSymbolInfo_10464 := fieldSymbolInfo.(*Ast_SymbolInfo)
            if (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( fieldSymbolInfo_10464.FP.Get_kind() == Ast_SymbolKind__Mtd) ||
                Lns_GetEnv().SetStackVal( fieldSymbolInfo_10464.FP.Get_kind() == Ast_SymbolKind__Fun) ).(bool)){
                var retTypeList *LnsList
                retTypeList = fieldSymbolInfo_10464.FP.Get_typeInfo().FP.Get_retTypeInfoList()
                symbolInfo = fieldSymbolInfo_10464
                
                if retTypeList.Len() > 0{
                    {
                        _applyedType := retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.ApplyGeneric(classTypeInfo.FP.CreateAlt2typeMap(false), self.moduleType)
                        if _applyedType != nil {
                            applyedType := _applyedType.(*Ast_TypeInfo)
                            fieldTypeInfo = applyedType
                            
                        } else {
                            fieldTypeInfo = retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                            
                        }
                    }
                }
                if fieldSymbolInfo_10464.FP.Get_typeInfo().FP.Get_argTypeInfoList().Len() > 0{
                    self.FP.addErrMess(token.Pos, Lns_getVM().String_format("can't use '$' with -- %s", []LnsAny{fieldSymbolInfo_10464.FP.Get_typeInfo().FP.GetTxt(nil, nil, nil)}))
                }
                getterFlag = true
                
            }
        }
    }
    if Lns_op_not(symbolInfo){
        symbolInfo = classScope.FP.GetSymbolInfoField(token.Txt, true, self.scope, self.scopeAccess)
        
        if Lns_op_not(symbolInfo){
            symbolInfo = classScope.FP.GetSymbolInfoIfField(token.Txt, self.scope, self.scopeAccess)
            
        }
        if symbolInfo != nil{
            symbolInfo_10475 := symbolInfo.(*Ast_SymbolInfo)
            fieldTypeInfo = symbolInfo_10475.FP.Get_typeInfo()
            
        }
    }
    if Lns_op_not(fieldTypeInfo){
        for _name, _val := range( classScope.FP.Get_symbol2SymbolInfoMap().Items ) {
            name := _name.(string)
            val := _val.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            Util_log(Lns_getVM().String_format("debug: %s, %s", []LnsAny{name, val}))
        }
        self.FP.Error(Lns_getVM().String_format("not found field typeInfo: %s.%s", []LnsAny{className, token.Txt}))
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = Lns_unwrapDefault( fieldTypeInfo, Ast_builtinTypeNone).(*Ast_TypeInfo)
    if symbolInfo != nil{
        symbolInfo_10482 := symbolInfo.(*Ast_SymbolInfo)
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( self.FP.inAnalyzingState(TransUnit_AnalyzingState__InitBlock)) ||
            Lns_GetEnv().SetStackVal( self.FP.inAnalyzingState(TransUnit_AnalyzingState__ClassMethod)) ).(bool){
            var errorMess LnsAny
            errorMess = nil
            if Lns_isCondTrue( self.protoFuncMap.Items[symbolInfo_10482.FP.Get_typeInfo()]){
                errorMess = Lns_getVM().String_format("It can't call prototype function from static -- %s", []LnsAny{symbolInfo_10482.FP.Get_name()})
                
            }
            if errorMess != nil{
                errorMess_10487 := errorMess.(string)
                self.FP.addErrMess(token.Pos, errorMess_10487)
            }
        } else if self.FP.inAnalyzingState(TransUnit_AnalyzingState__Constructor){
            var errorMess LnsAny
            errorMess = nil
            if Lns_isCondTrue( self.protoFuncMap.Items[symbolInfo_10482.FP.Get_typeInfo()]){
                errorMess = "It can't call prototype function from '__init'"
                
            } else { 
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( symbolInfo_10482.FP.Get_typeInfo().FP.Get_kind() == Ast_TypeInfoKind__Method) &&
                    Lns_GetEnv().SetStackVal( symbolInfo_10482.FP.Get_scope() == classScope) ).(bool)){
                    for _, _val := range( classScope.FP.Get_symbol2SymbolInfoMap().Items ) {
                        val := _val.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( val.FP.Get_kind() == Ast_SymbolKind__Mbr) &&
                            Lns_GetEnv().SetStackVal( Lns_op_not(val.FP.Get_staticFlag())) ).(bool)){
                            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                Lns_GetEnv().SetStackVal( Lns_op_not(val.FP.Get_hasValueFlag())) &&
                                Lns_GetEnv().SetStackVal( Lns_op_not(val.FP.Get_typeInfo().FP.Get_nilable())) ).(bool)){
                                errorMess = Lns_getVM().String_format("Set member(%s) before to access the method-- %s", []LnsAny{val.FP.Get_name(), symbolInfo_10482.FP.Get_name()})
                                
                                break
                            }
                        }
                    }
                }
            }
            if errorMess != nil{
                errorMess_10498 := errorMess.(string)
                self.FP.addErrMess(token.Pos, errorMess_10498)
            }
        }
    }
    return typeInfo, symbolInfo, getterFlag
}

// 2002: decl @lune.@base.@TransUnit.TransUnit.dumpComp
func (self *TransUnit_TransUnit) dumpComp(writer Writer_Writer,pattern string,symbolInfo *Ast_SymbolInfo,getterFlag bool) bool {
    var symbol string
    symbol = symbolInfo.FP.Get_name()
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( pattern == "") ||
        Lns_GetEnv().SetStackVal( Lns_car(Lns_getVM().String_find(symbol,pattern, nil, nil))) )){
        if getterFlag{
            writer.StartParent("candidate", false)
            var typeInfo *Ast_TypeInfo
            typeInfo = symbolInfo.FP.Get_typeInfo()
            writer.Write("type", Lns_getVM().String_format("%s", []LnsAny{Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind())}))
            if _switch62217 := (symbolInfo.FP.Get_kind()); _switch62217 == Ast_SymbolKind__Mtd || _switch62217 == Ast_SymbolKind__Fun || _switch62217 == Ast_SymbolKind__Mac {
                writer.Write("displayTxt", Lns_getVM().String_format("$%s", []LnsAny{Lns_car(Lns_getVM().String_gsub(typeInfo.FP.Get_rawTxt(),"^get_", "")).(string)}))
            } else if _switch62217 == Ast_SymbolKind__Mbr {
                writer.Write("displayTxt", Lns_getVM().String_format("$%s: %s", []LnsAny{symbolInfo.FP.Get_name(), typeInfo.FP.GetTxt(nil, nil, nil)}))
            }
        } else { 
            writer.StartParent("candidate", false)
            var typeInfo *Ast_TypeInfo
            typeInfo = symbolInfo.FP.Get_typeInfo()
            writer.Write("type", Lns_getVM().String_format("%s", []LnsAny{Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind())}))
            if _switch62439 := (symbolInfo.FP.Get_kind()); _switch62439 == Ast_SymbolKind__Fun || _switch62439 == Ast_SymbolKind__Mtd || _switch62439 == Ast_SymbolKind__Mac {
                writer.Write("displayTxt", typeInfo.FP.Get_display_stirng_with(symbolInfo.FP.Get_name(), nil))
            } else if _switch62439 == Ast_SymbolKind__Mbr || _switch62439 == Ast_SymbolKind__Var || _switch62439 == Ast_SymbolKind__Arg {
                var name string
                name = symbolInfo.FP.Get_name()
                {
                    _algeTypeInfo := Ast_AlgeTypeInfoDownCastF(typeInfo.FP)
                    if _algeTypeInfo != nil {
                        algeTypeInfo := _algeTypeInfo.(*Ast_AlgeTypeInfo)
                        {
                            _valInfo := algeTypeInfo.FP.GetValInfo(name)
                            if _valInfo != nil {
                                valInfo := _valInfo.(*Ast_AlgeValInfo)
                                if valInfo.FP.Get_typeList().Len() > 0{
                                    name = Lns_getVM().String_format("%s(", []LnsAny{name})
                                    
                                    for _index, _itemType := range( valInfo.FP.Get_typeList().Items ) {
                                        index := _index + 1
                                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                                        if index > 1{
                                            name = name + ","
                                            
                                        }
                                        name = name + itemType.FP.Get_display_stirng()
                                        
                                    }
                                    name = name + ")"
                                    
                                }
                            }
                        }
                    }
                }
                writer.Write("displayTxt", Lns_getVM().String_format("%s: %s", []LnsAny{name, typeInfo.FP.Get_display_stirng()}))
            } else if _switch62439 == Ast_SymbolKind__Typ {
                writer.Write("displayTxt", Lns_getVM().String_format("%s", []LnsAny{Lns_car(Lns_getVM().String_gsub(typeInfo.FP.Get_display_stirng(),"@", "")).(string)}))
            }
        }
        writer.EndElement()
    }
    return true
}

// 2063: decl @lune.@base.@TransUnit.TransUnit.dumpFieldComp
func (self *TransUnit_TransUnit) dumpFieldComp(writer Writer_Writer,isPrefixType bool,prefixTypeInfo *Ast_TypeInfo,pattern string,getterPattern LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = prefixTypeInfo
    var scope *Ast_Scope
    
    {
        _scope := typeInfo.FP.Get_scope()
        if _scope == nil{
            return 
        } else {
            scope = _scope.(*Ast_Scope)
        }
    }
    scope.FP.FilterTypeInfoField(true, self.scope, self.scopeAccess, Ast_filterForm(func(symbolInfo *Ast_SymbolInfo) bool {
        if (isPrefixType){
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(symbolInfo.FP.Get_staticFlag())) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(symbolInfo.FP.Get_typeInfo().FP.Get_staticFlag())) &&
                Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() != Ast_SymbolKind__Typ) ).(bool)){
                return true
            }
        } else if symbolInfo.FP.Get_staticFlag(){
            return true
        }
        var symbol string
        symbol = symbolInfo.FP.Get_name()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symbol != "__init") &&
            Lns_GetEnv().SetStackVal( symbol != "__free") &&
            Lns_GetEnv().SetStackVal( symbol != "self") ).(bool)){
            if getterPattern != nil{
                getterPattern_10550 := getterPattern.(string)
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mtd) ||
                    Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Fun) ).(bool){
                    var retList *LnsList
                    retList = symbolInfo.FP.Get_typeInfo().FP.Get_retTypeInfoList()
                    if retList.Len() == 1{
                        return self.FP.dumpComp(writer, getterPattern_10550, symbolInfo, true)
                    }
                }
                return true
            }
            return self.FP.dumpComp(writer, pattern, symbolInfo, false)
        }
        return true
    }))
}

// 2107: decl @lune.@base.@TransUnit.TransUnit.dumpSymbolComp
func (self *TransUnit_TransUnit) dumpSymbolComp(writer Writer_Writer,scope *Ast_Scope,pattern string) {
    scope.FP.FilterSymbolTypeInfo(scope, self.moduleScope, self.scopeAccess, Ast_filterForm(func(symbolInfo *Ast_SymbolInfo) bool {
        return self.FP.dumpComp(writer, pattern, symbolInfo, false)
    }))
}

// 2117: decl @lune.@base.@TransUnit.TransUnit.checkComp
func (self *TransUnit_TransUnit) checkComp(token *Types_Token,callback TransUnit_checkCompForm_6433_) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.analyzeMode == TransUnit_AnalyzeMode__Complete) &&
        Lns_GetEnv().SetStackVal( self.FP.isTargetToken(token)) ).(bool)){
        var currentModule string
        currentModule = TransUnit_convExp62753(Lns_2DDD(Lns_getVM().String_gsub(self.parser.FP.GetStreamName(),"%.lns", "")))
        currentModule = TransUnit_convExp62768(Lns_2DDD(Lns_getVM().String_gsub(currentModule,".*/", "")))
        
        var target string
        target = TransUnit_convExp62783(Lns_2DDD(Lns_getVM().String_gsub(self.analyzeModule,"[^%.]+%.", "")))
        if currentModule == target{
            var jsonWriter *Writer_JSON
            jsonWriter = NewWriter_JSON(Lns_io_stdout)
            jsonWriter.FP.StartParent("lunescript", false)
            var prefix string
            prefix = TransUnit_convExp62823(Lns_2DDD(Lns_getVM().String_gsub(token.Txt,"lune$", "")))
            jsonWriter.FP.Write("prefix", prefix)
            jsonWriter.FP.StartParent("candidateList", true)
            callback(jsonWriter, prefix)
            jsonWriter.FP.EndElement()
            jsonWriter.FP.EndElement()
            jsonWriter.FP.Fin()
            Lns_getVM().OS_exit(0)
        }
    }
}

// 2141: decl @lune.@base.@TransUnit.TransUnit.checkFieldComp
func (self *TransUnit_TransUnit) checkFieldComp(getterFlag bool,token *Types_Token,prefixExp *Nodes_Node) {
    if self.analyzeMode != TransUnit_AnalyzeMode__Complete{
        return 
    }
    var prefixSymbolInfoList *LnsList
    prefixSymbolInfoList = prefixExp.FP.GetSymbolInfo()
    var prefixSymbolInfo LnsAny
    prefixSymbolInfo = nil
    if prefixSymbolInfoList.Len() == 1{
        prefixSymbolInfo = prefixSymbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        
    }
    self.FP.checkComp(token, TransUnit_checkCompForm_6433_(func(jsonWriter *Writer_JSON,prefix string) {
        var getterPattern LnsAny
        getterPattern = nil
        if getterFlag{
            getterPattern = "^get_" + prefix
            
        }
        var isPrefixType bool
        isPrefixType = false
        {
            __exp := prefixSymbolInfo
            if __exp != nil {
                _exp := __exp.(*Ast_SymbolInfo)
                isPrefixType = _exp.FP.Get_kind() == Ast_SymbolKind__Typ
                
            }
        }
        self.FP.dumpFieldComp(jsonWriter.FP, isPrefixType, prefixExp.FP.Get_expType(), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( prefix == "") &&
            Lns_GetEnv().SetStackVal( "") ||
            Lns_GetEnv().SetStackVal( "^" + prefix) ).(string), getterPattern)
    }))
}

// 2171: decl @lune.@base.@TransUnit.TransUnit.checkEnumComp
func (self *TransUnit_TransUnit) checkEnumComp(token *Types_Token,enumTypeInfo *Ast_EnumTypeInfo) {
    if self.analyzeMode != TransUnit_AnalyzeMode__Complete{
        return 
    }
    self.FP.checkComp(token, TransUnit_checkCompForm_6433_(func(jsonWriter *Writer_JSON,prefix string) {
        var scope *Ast_Scope
        
        {
            _scope := enumTypeInfo.FP.Get_scope()
            if _scope == nil{
                return 
            } else {
                scope = _scope.(*Ast_Scope)
            }
        }
        var pattern string
        pattern = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( prefix == "") &&
            Lns_GetEnv().SetStackVal( "") ||
            Lns_GetEnv().SetStackVal( "^" + prefix) ).(string)
        scope.FP.FilterTypeInfoField(true, self.scope, self.scopeAccess, Ast_filterForm(func(symbolInfo *Ast_SymbolInfo) bool {
            if symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mbr{
                return self.FP.dumpComp(jsonWriter.FP, pattern, symbolInfo, false)
            }
            return true
        }))
    }))
}

// 2197: decl @lune.@base.@TransUnit.TransUnit.checkAlgeComp
func (self *TransUnit_TransUnit) checkAlgeComp(token *Types_Token,algeTypeInfo *Ast_AlgeTypeInfo) {
    if self.analyzeMode != TransUnit_AnalyzeMode__Complete{
        return 
    }
    self.FP.checkComp(token, TransUnit_checkCompForm_6433_(func(jsonWriter *Writer_JSON,prefix string) {
        self.FP.dumpFieldComp(jsonWriter.FP, true, &algeTypeInfo.Ast_TypeInfo, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( prefix == "") &&
            Lns_GetEnv().SetStackVal( "") ||
            Lns_GetEnv().SetStackVal( "^" + prefix) ).(string), nil)
    }))
}

// 2214: decl @lune.@base.@TransUnit.TransUnit.checkSymbolComp
func (self *TransUnit_TransUnit) checkSymbolComp(token *Types_Token) {
    self.FP.checkComp(token, TransUnit_checkCompForm_6433_(func(jsonWriter *Writer_JSON,prefix string) {
        self.FP.dumpSymbolComp(jsonWriter.FP, self.scope, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( prefix == "") &&
            Lns_GetEnv().SetStackVal( "") ||
            Lns_GetEnv().SetStackVal( "^" + prefix) ).(string))
    }))
}

// 2225: decl @lune.@base.@TransUnit.TransUnit.analyzeExpField
func (self *TransUnit_TransUnit) analyzeExpField(firstToken *Types_Token,fieldToken *Types_Token,mode LnsInt,prefixExp *Nodes_Node) *Nodes_Node {
    if prefixExp.FP.Get_expTypeList().Len() > 1{
        prefixExp = &Nodes_ExpMultiTo1Node_create(self.nodeManager, prefixExp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(prefixExp.FP.Get_expType())}), prefixExp).Nodes_Node
        
    }
    var threading bool
    var accessNil bool
    accessNil = false
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__FieldNil) ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__GetNil) ).(bool){
        accessNil = true
        
        if Lns_op_not(prefixExp.FP.Get_expType().FP.Get_nilable()){
            self.FP.addWarnMess(prefixExp.FP.Get_pos(), Lns_getVM().String_format("This is not nilable. -- %s", []LnsAny{prefixExp.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
        }
        threading = self.FP.checkThreading(firstToken.Pos)
        
    } else { 
        threading = false
        
    }
    if self.macroCtrl.FP.Get_analyzeInfo().FP.IsAnalyzingSymArg(){
        if accessNil{
            self.helperInfo.UseNilAccess = true
            
        }
        return &Nodes_RefFieldNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeSymbol)}), fieldToken, nil, accessNil, threading, prefixExp).Nodes_Node
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_builtinTypeStem_
    var prefixExpType *Ast_TypeInfo
    prefixExpType = prefixExp.FP.Get_expType()
    if accessNil{
        if prefixExpType.FP.Get_nilable(){
            prefixExpType = prefixExpType.FP.Get_nonnilableType()
            
            if prefixExpType.FP.Get_srcTypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Box{
                self.FP.addErrMess(prefixExp.FP.Get_pos(), "Nilable can't support '$.' access yet")
            }
        } else { 
            accessNil = false
            
        }
    }
    var extFlag bool
    if Ast_isExtType(prefixExpType){
        extFlag = true
        
        prefixExpType = prefixExpType.FP.Get_extedType()
        
    } else { 
        extFlag = false
        
    }
    self.FP.checkFieldComp(Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__Get) ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__GetNil) ).(bool), fieldToken, prefixExp)
    if accessNil{
        self.helperInfo.UseNilAccess = true
        
        if _switch63650 := prefixExpType.FP.Get_kind(); _switch63650 == Ast_TypeInfoKind__Set || _switch63650 == Ast_TypeInfoKind__Enum || _switch63650 == Ast_TypeInfoKind__Alge {
            self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("%s does not support $.", []LnsAny{prefixExpType.FP.GetTxt(nil, nil, nil)}))
        }
    }
    var prefixSymbolInfoList *LnsList
    prefixSymbolInfoList = prefixExp.FP.GetSymbolInfo()
    self.FP.checkSymbolHavingValue(prefixExp.FP.Get_pos(), prefixSymbolInfoList)
    var getterTypeInfo LnsAny
    getterTypeInfo = nil
    var symbolInfo LnsAny
    symbolInfo = nil
    if _switch64350 := prefixExpType.FP.Get_kind(); _switch64350 == Ast_TypeInfoKind__Class || _switch64350 == Ast_TypeInfoKind__Module || _switch64350 == Ast_TypeInfoKind__ExtModule || _switch64350 == Ast_TypeInfoKind__IF || _switch64350 == Ast_TypeInfoKind__List || _switch64350 == Ast_TypeInfoKind__Array || _switch64350 == Ast_TypeInfoKind__Set || _switch64350 == Ast_TypeInfoKind__Box || _switch64350 == Ast_TypeInfoKind__Alternate {
        var getterFlag bool
        getterFlag = false
        typeInfo, symbolInfo, getterFlag = self.FP.analyzeAccessClassField(prefixExpType, mode, fieldToken)
        
        if getterFlag{
            {
                __exp := symbolInfo
                if __exp != nil {
                    _exp := __exp.(*Ast_SymbolInfo)
                    getterTypeInfo = _exp.FP.Get_typeInfo()
                    
                }
            }
        }
    } else if _switch64350 == Ast_TypeInfoKind__Enum || _switch64350 == Ast_TypeInfoKind__Alge {
        var scope *Ast_Scope
        scope = Lns_unwrap( prefixExpType.FP.Get_scope()).(*Ast_Scope)
        var fieldName string
        fieldName = fieldToken.Txt
        var symbolInfoList *LnsList
        symbolInfoList = prefixExp.FP.GetSymbolInfo()
        var isTypeSymbol bool
        isTypeSymbol = false
        if symbolInfoList.Len() > 0{
            if symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_kind() == Ast_SymbolKind__Typ{
                isTypeSymbol = true
                
            }
        }
        if mode == TransUnit_ExpSymbolMode__Get{
            var moduleType *Ast_TypeInfo
            moduleType = prefixExpType.FP.GetModule()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(moduleType.FP.Equals(self.processInfo, self.moduleType, nil, nil))) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(self.scope.FP.GetModuleInfo(moduleType))) ).(bool)){
                if Lns_op_not(self.importedAliasMap.Items[prefixExpType]){
                    self.FP.addErrMess(fieldToken.Pos, Lns_getVM().String_format("need to import module -- %s", []LnsAny{prefixExpType.FP.GetModule().FP.GetTxt(nil, nil, nil)}))
                }
            }
            fieldName = "get_" + fieldName
            
            {
                _funcSymbol := scope.FP.GetSymbolInfoChild(fieldName)
                if _funcSymbol != nil {
                    funcSymbol := _funcSymbol.(*Ast_SymbolInfo)
                    symbolInfo = funcSymbol
                    
                    var funcType *Ast_TypeInfo
                    funcType = funcSymbol.FP.Get_typeInfo()
                    if funcType.FP.Get_staticFlag() != isTypeSymbol{
                        self.FP.addErrMess(prefixExp.FP.Get_pos(), Lns_getVM().String_format("Can't access -- %s, %s", []LnsAny{fieldName, isTypeSymbol}))
                    }
                    var retTypeList *LnsList
                    retTypeList = funcType.FP.Get_retTypeInfoList()
                    if retTypeList.Len() == 0{
                        self.FP.addErrMess(fieldToken.Pos, Lns_getVM().String_format("The func (%s) doesn't return value.", []LnsAny{funcType.FP.GetTxt(nil, nil, nil)}))
                    } else { 
                        typeInfo = retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        
                    }
                } else {
                    self.FP.addErrMess(fieldToken.Pos, Lns_getVM().String_format("not found -- %s.", []LnsAny{fieldName}))
                    typeInfo = Ast_builtinTypeNone
                    
                }
            }
            getterTypeInfo = Ast_headTypeInfo
            
        } else { 
            {
                __exp := scope.FP.GetTypeInfoChild(fieldName)
                if __exp != nil {
                    _exp := __exp.(*Ast_TypeInfo)
                    typeInfo = _exp
                    
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Enum) ||
                        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Alge) ).(bool){
                        if Lns_op_not(isTypeSymbol){
                            self.FP.addErrMess(fieldToken.Pos, Lns_getVM().String_format("can't access field -- %s", []LnsAny{fieldToken.Txt}))
                        }
                    }
                } else {
                    self.FP.addErrMess(fieldToken.Pos, Lns_getVM().String_format("not found field -- %s", []LnsAny{fieldToken.Txt}))
                    typeInfo = Ast_builtinTypeInt
                    
                }
            }
        }
    } else if _switch64350 == Ast_TypeInfoKind__Map {
        var work *Ast_TypeInfo
        work = prefixExpType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(work.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil)){
            self.FP.addErrMess(fieldToken.Pos, Lns_getVM().String_format("map key type is not str. (%s)", []LnsAny{work.FP.GetTxt(nil, nil, nil)}))
        }
        typeInfo = prefixExpType.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        
        if Lns_op_not(typeInfo.FP.Get_nilable()){
            typeInfo = typeInfo.FP.Get_nilableTypeInfo()
            
        }
        if extFlag{
            typeInfo = self.FP.createExtType(fieldToken.Pos, typeInfo)
            
        }
        return &Nodes_ExpRefItemNode_create(self.nodeManager, fieldToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), prefixExp, accessNil, threading, fieldToken.Txt, nil).Nodes_Node
    } else {
        if prefixExpType.FP.Equals(self.processInfo, Ast_builtinTypeStem, nil, nil){
            typeInfo = Ast_builtinTypeStem_
            
            if extFlag{
                typeInfo = self.FP.createExtType(fieldToken.Pos, typeInfo)
                
            }
            return &Nodes_ExpRefItemNode_create(self.nodeManager, fieldToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), prefixExp, accessNil, threading, fieldToken.Txt, nil).Nodes_Node
        } else { 
            self.FP.Error(Lns_getVM().String_format("illegal type -- %s, %s", []LnsAny{prefixExpType.FP.GetTxt(nil, nil, nil), Ast_TypeInfoKind_getTxt( prefixExpType.FP.Get_kind())}))
        }
    }
    if Lns_op_not(symbolInfo){
        var prefixScope LnsAny
        prefixScope = prefixExpType.FP.Get_scope()
        {
            __exp := prefixScope
            if __exp != nil {
                _exp := __exp.(*Ast_Scope)
                symbolInfo = _exp.FP.GetSymbolInfoField(fieldToken.Txt, true, self.scope, self.scopeAccess)
                
            }
        }
    }
    if symbolInfo != nil{
        symbolInfo_10712 := symbolInfo.(*Ast_SymbolInfo)
        if prefixSymbolInfoList.Len() == 1{
            var prefixSymbolInfo *Ast_SymbolInfo
            prefixSymbolInfo = prefixSymbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if prefixSymbolInfo.FP.Get_kind() == Ast_SymbolKind__Typ{
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_op_not(symbolInfo_10712.FP.Get_staticFlag())) &&
                    Lns_GetEnv().SetStackVal( symbolInfo_10712.FP.Get_kind() != Ast_SymbolKind__Typ) ).(bool)){
                    self.FP.addErrMess(fieldToken.Pos, Lns_getVM().String_format("Type can't access this symbol. -- %s", []LnsAny{symbolInfo_10712.FP.Get_name()}))
                }
            } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( symbolInfo_10712.FP.Get_staticFlag()) &&
                Lns_GetEnv().SetStackVal( symbolInfo_10712.FP.Get_typeInfo().FP.Get_kind() != Ast_TypeInfoKind__Method) ).(bool)){
                self.FP.addErrMess(fieldToken.Pos, Lns_getVM().String_format("can't access this symbol. -- %s", []LnsAny{fieldToken.Txt}))
            }
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(prefixExpType))) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(symbolInfo_10712.FP.Get_staticFlag())) &&
            Lns_GetEnv().SetStackVal( symbolInfo_10712.FP.Get_kind() == Ast_SymbolKind__Mtd) &&
            Lns_GetEnv().SetStackVal( symbolInfo_10712.FP.Get_mutable()) ).(bool)){
            self.FP.addErrMess(fieldToken.Pos, Lns_getVM().String_format("can't access mutable method. -- %s.%s", []LnsAny{prefixExpType.FP.GetTxt(nil, nil, nil), fieldToken.Txt}))
        }
    }
    var accessSymbolInfo LnsAny
    accessSymbolInfo = nil
    var symbolMutMode LnsInt
    symbolMutMode = typeInfo.FP.Get_mutMode()
    if symbolInfo != nil{
        symbolInfo_10722 := symbolInfo.(*Ast_SymbolInfo)
        var workSymInfo *Ast_AccessSymbolInfo
        workSymInfo = NewAst_AccessSymbolInfo(symbolInfo_10722, &Ast_OverrideMut__Prefix{prefixExpType}, Lns_op_not(accessNil))
        if Lns_op_not(getterTypeInfo){
            typeInfo = workSymInfo.FP.Get_typeInfo()
            
        }
        accessSymbolInfo = workSymInfo
        
        if _switch64609 := mode; _switch64609 == TransUnit_ExpSymbolMode__Field || _switch64609 == TransUnit_ExpSymbolMode__FieldNil {
            symbolMutMode = symbolInfo_10722.FP.Get_mutMode()
            
        }
    }
    if accessNil{
        if Lns_op_not(typeInfo.FP.Get_nilable()){
            typeInfo = typeInfo.FP.Get_nilableTypeInfo()
            
        }
        self.helperInfo.UseNilAccess = true
        
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(prefixExpType))) &&
        Lns_GetEnv().SetStackVal( symbolMutMode == Ast_MutMode__Mut) ).(bool)){
        typeInfo = self.FP.CreateModifier(typeInfo, Ast_MutMode__IMut)
        
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Equals(self.processInfo, TransUnit_builtinFunc.List_unpack, nil, nil)) ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Equals(self.processInfo, TransUnit_builtinFunc.Array_unpack, nil, nil)) ).(bool){
        self.helperInfo.UseUnpack = true
        
    }
    {
        _expRef := Nodes_ExpRefNodeDownCastF(prefixExp.FP)
        if _expRef != nil {
            expRef := _expRef.(*Nodes_ExpRefNode)
            var prefixSym *Ast_SymbolInfo
            prefixSym = expRef.FP.Get_symbolInfo()
            var prefixType *Ast_TypeInfo
            prefixType = prefixSym.FP.Get_typeInfo()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( prefixSym.FP.Get_kind() == Ast_SymbolKind__Typ) &&
                Lns_GetEnv().SetStackVal( prefixType.FP.Get_kind() == Ast_TypeInfoKind__Class) &&
                Lns_GetEnv().SetStackVal( prefixType.FP.Get_itemTypeInfoList().Len() > 0) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isGenericType(prefixType))) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(self.scope.FP.IsInnerOf(Lns_unwrap( prefixType.FP.Get_scope()).(*Ast_Scope)))) ).(bool)){
                var accessErr bool
                accessErr = false
                if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Func{
                    var altSet *LnsSet
                    altSet = NewLnsSet([]LnsAny{})
                    for _, _argType := range( typeInfo.FP.Get_argTypeInfoList().Items ) {
                        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        var orgType *Ast_TypeInfo
                        orgType = argType.FP.Get_nonnilableType().FP.Get_srcTypeInfo()
                        if orgType.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
                            altSet.Add(Ast_TypeInfo2Stem(orgType))
                        }
                    }
                    for _, _itemType := range( prefixType.FP.Get_itemTypeInfoList().Items ) {
                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if Lns_op_not(altSet.Has(Ast_TypeInfo2Stem(itemType.FP.Get_nonnilableType().FP.Get_srcTypeInfo()))){
                            accessErr = true
                            
                            break
                        }
                    }
                } else { 
                    accessErr = true
                    
                }
                if accessErr{
                    self.FP.addErrMess(prefixExp.FP.Get_pos(), Lns_getVM().String_format("can't access this class(%s) without '<>'.", []LnsAny{prefixType.FP.GetTxt(nil, nil, nil)}))
                }
            }
        }
    }
    if extFlag{
        typeInfo = self.FP.createExtType(firstToken.Pos, typeInfo)
        
        if self.FP.checkThreading(firstToken.Pos){
            self.FP.addErrMess(firstToken.Pos, "not support to use Luaval on thread.")
        }
    }
    {
        __exp := getterTypeInfo
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            return &Nodes_GetFieldNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), fieldToken, Ast_SymbolInfoDownCastF(accessSymbolInfo), accessNil, threading, prefixExp, _exp).Nodes_Node
        } else {
            return &Nodes_RefFieldNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), fieldToken, Ast_SymbolInfoDownCastF(accessSymbolInfo), accessNil, threading, prefixExp).Nodes_Node
        }
    }
// insert a dummy
    return nil
}

// 2563: decl @lune.@base.@TransUnit.TransUnit.analyzeNewAlge
func (self *TransUnit_TransUnit) analyzeNewAlge(firstToken *Types_Token,algeTypeInfo *Ast_AlgeTypeInfo,prefix LnsAny) *Nodes_NewAlgeValNode {
    var symbolToken *Types_Token
    symbolToken = self.FP.getSymbolToken(TransUnit_SymbolMode__MustNot_)
    self.FP.checkAlgeComp(symbolToken, algeTypeInfo)
    {
        _valInfo := algeTypeInfo.FP.GetValInfo(symbolToken.Txt)
        if _valInfo != nil {
            valInfo := _valInfo.(*Ast_AlgeValInfo)
            var argList *LnsList
            argList = NewLnsList([]LnsAny{})
            var argListNode LnsAny
            if valInfo.FP.Get_typeList().Len() > 0{
                self.FP.checkNextToken("(")
                argListNode = self.FP.analyzeExpList(false, false, false, nil, valInfo.FP.Get_typeList(), nil)
                
                argList = (Lns_unwrap( argListNode).(*Nodes_ExpListNode)).FP.Get_expList()
                
                self.FP.checkNextToken(")")
            } else { 
                argListNode = nil
                
            }
            {
                _, _, _newExpNodeList := TransUnit_convExp65190(Lns_2DDD(self.FP.checkMatchType("call", symbolToken.Pos, valInfo.FP.Get_typeList(), argListNode, false, true, nil)))
                if _newExpNodeList != nil {
                    newExpNodeList := _newExpNodeList.(*Nodes_ExpListNode)
                    argList = newExpNodeList.FP.Get_expList()
                    
                }
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( algeTypeInfo.FP.Get_externalFlag()) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(self.scope.FP.GetModuleInfo(algeTypeInfo.FP.GetModule().FP.Get_srcTypeInfo()))) ).(bool)){
                var fullname string
                fullname = algeTypeInfo.FP.GetFullName(self.typeNameCtrl, self.scope.FP, true)
                self.FP.addErrMess(firstToken.Pos, Lns_getVM().String_format("This module not import -- %s", []LnsAny{fullname}))
            }
            return Nodes_NewAlgeValNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_AlgeTypeInfo2Stem(algeTypeInfo)}), symbolToken, prefix, algeTypeInfo, valInfo, argList)
        } else {
            self.FP.addErrMess(symbolToken.Pos, Lns_getVM().String_format("not found Alge -- %s", []LnsAny{symbolToken.Txt}))
            return Nodes_NewAlgeValNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_AlgeTypeInfo2Stem(algeTypeInfo)}), symbolToken, prefix, algeTypeInfo, NewAst_AlgeValInfo("", NewLnsList([]LnsAny{}), algeTypeInfo), NewLnsList([]LnsAny{}))
        }
    }
// insert a dummy
    return nil
}

// 2616: decl @lune.@base.@TransUnit.TransUnit.analyzeExpSymbol
func (self *TransUnit_TransUnit) analyzeExpSymbol(firstToken *Types_Token,symbolToken *Types_Token,mode LnsInt,prefixExp LnsAny,skipFlag bool,canLeftExp bool) *Nodes_Node {
    var exp *Nodes_Node
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__Field) ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__Get) ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__FieldNil) ||
        Lns_GetEnv().SetStackVal( mode == TransUnit_ExpSymbolMode__GetNil) ).(bool){
        if prefixExp != nil{
            prefixExp_10786 := prefixExp.(*Nodes_Node)
            exp = self.FP.analyzeExpField(firstToken, symbolToken, mode, prefixExp_10786)
            
            var expType *Ast_TypeInfo
            expType = exp.FP.Get_expType()
            if prefixExp_10786.FP.Get_expType().FP.IsModule(){
                {
                    _algeType := Ast_AlgeTypeInfoDownCastF(expType.FP)
                    if _algeType != nil {
                        algeType := _algeType.(*Ast_AlgeTypeInfo)
                        var nextToken *Types_Token
                        nextToken = self.FP.getToken(nil)
                        if nextToken.Txt == "."{
                            return &self.FP.analyzeNewAlge(firstToken, algeType, exp).Nodes_Node
                        }
                        self.FP.Pushback()
                    }
                }
            }
        } else {
            Util_err("prefixExp is nil")
        }
    } else if mode == TransUnit_ExpSymbolMode__Symbol{
        if self.macroCtrl.FP.Get_analyzeInfo().FP.IsAnalyzingSymArg(){
            exp = &Nodes_LiteralSymbolNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeSymbol)}), symbolToken).Nodes_Node
            
        } else { 
            self.FP.checkSymbolComp(symbolToken)
            var symbolInfo *Ast_SymbolInfo
            
            {
                _symbolInfo := self.scope.FP.GetSymbolTypeInfo(symbolToken.Txt, self.scope, self.moduleScope, self.scopeAccess)
                if _symbolInfo == nil{
                    if self.analyzeMode != TransUnit_AnalyzeMode__Diag{
                        var work *Ast_Scope
                        work = self.scope
                        for  {
                            Lns_print([]LnsAny{work, self.globalScope, Ast_rootScope})
                            if work == work.FP.Get_parent(){
                                break
                            }
                            work = work.FP.Get_parent()
                            
                        }
                        self.scope.FP.FilterSymbolTypeInfo(self.scope, self.moduleScope, self.scopeAccess, Ast_filterForm(TransUnit_analyzeExpSymbol___anonymous_6521_))
                    }
                    self.FP.Error("not found type -- " + symbolToken.Txt)
                } else {
                    symbolInfo = _symbolInfo.(*Ast_SymbolInfo)
                }
            }
            self.FP.accessSymbol(symbolInfo, canLeftExp)
            var typeInfo *Ast_TypeInfo
            typeInfo = symbolInfo.FP.Get_typeInfo()
            if _switch65812 := symbolInfo.FP.Get_kind(); _switch65812 == Ast_SymbolKind__Typ {
                {
                    _algeType := Ast_AlgeTypeInfoDownCastF(typeInfo.FP)
                    if _algeType != nil {
                        algeType := _algeType.(*Ast_AlgeTypeInfo)
                        var nextToken *Types_Token
                        nextToken = self.FP.getToken(nil)
                        if nextToken.Txt == "."{
                            return &self.FP.analyzeNewAlge(firstToken, algeType, nil).Nodes_Node
                        }
                        self.FP.Pushback()
                    }
                }
            } else if _switch65812 == Ast_SymbolKind__Var {
                self.tentativeSymbol.FP.AddAccessSym(symbolInfo)
                if Lns_op_not(symbolInfo.FP.Get_hasValueFlag()){
                    var nsTypeInfo *Ast_TypeInfo
                    nsTypeInfo = self.FP.getCurrentNamespaceTypeInfo()
                    if Lns_op_not(symbolInfo.FP.Get_scope().FP.IsInnerOf(Lns_unwrap( nsTypeInfo.FP.Get_scope()).(*Ast_Scope))){
                        self.tentativeSymbol.FP.AddAccessSymPos(NewTransUnit_AccessSymPos(symbolInfo, firstToken.Pos))
                    }
                }
            }
            if typeInfo.FP.Equals(self.processInfo, Ast_builtinTypeSymbol, nil, nil){
                skipFlag = true
                
            }
            if typeInfo.FP.Equals(self.processInfo, TransUnit_builtinFunc.Lns__load, nil, nil){
                self.helperInfo.UseLoad = true
                
            }
            if _switch65949 := symbolToken.Txt; _switch65949 == "__func__" {
                var funcTypeInfo *Ast_TypeInfo
                funcTypeInfo = self.FP.getCurrentNamespaceTypeInfo()
                self.has__func__Symbol.Add(Ast_TypeInfo2Stem(funcTypeInfo))
            } else if _switch65949 == "_G" || _switch65949 == "_ENV" {
                var valid bool
                valid = false
                for _pragma := range( self.helperInfo.PragmaSet.Items ) {
                    pragma := _pragma
                    switch _exp65928 := pragma.(type) {
                    case *LuneControl_Pragma__limit_conv_code:
                    codeSet := _exp65928.Val1
                        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( codeSet.Len() == 1) &&
                            Lns_GetEnv().SetStackVal( codeSet.Has(LuneControl_Code__Lua)) ).(bool)){
                            valid = true
                            
                            break
                        }
                    }
                }
                if Lns_op_not(valid){
                    self.FP.addErrMess(firstToken.Pos, "'_G' and '_ENV' only can access with transcompiling to lua.")
                }
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( typeInfo.FP.Get_nonnilableType().FP.Get_kind() == Ast_TypeInfoKind__Ext) &&
                Lns_GetEnv().SetStackVal( self.FP.checkThreading(symbolToken.Pos)) ).(bool)){
                self.FP.addErrMess(symbolToken.Pos, "not support to use Luaval on thread.")
            }
            exp = &Nodes_ExpRefNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), &NewAst_AccessSymbolInfo(symbolInfo, Ast_OverrideMut__None_Obj, true).Ast_SymbolInfo).Nodes_Node
            
        }
    } else if mode == TransUnit_ExpSymbolMode__Fn{
        exp = self.FP.analyzeDeclFunc(TransUnit_DeclFuncMode__Func, false, false, Ast_AccessMode__Local, false, nil, symbolToken, nil)
        
    } else { 
        self.FP.Error(Lns_getVM().String_format("illegal mode -- %s", []LnsAny{mode}))
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.analyzeMode == TransUnit_AnalyzeMode__Inquire) &&
        Lns_GetEnv().SetStackVal( self.FP.isTargetToken(symbolToken)) ).(bool)){
        self.FP.dumpSymbolType(symbolToken.Txt, exp.FP.Get_expType())
    }
    return self.FP.analyzeExpCont(firstToken, exp, skipFlag, canLeftExp)
}

// 2778: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOpSet
func (self *TransUnit_TransUnit) analyzeExpOpSet(exp *Nodes_Node,opeToken *Types_Token,expectTypeList *LnsList) *Nodes_Node {
    exp.FP.SetLValue()
    if Lns_op_not(exp.FP.CanBeLeft()){
        self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("this node can not be l-value. -- %s", []LnsAny{Nodes_getNodeKindName(exp.FP.Get_kind())}))
    }
    var listRefItemNode LnsAny
    listRefItemNode = nil
    {
        _symNodeList := Nodes_ExpListNodeDownCastF(exp.FP)
        if _symNodeList != nil {
            symNodeList := _symNodeList.(*Nodes_ExpListNode)
            for _, _symNode := range( symNodeList.FP.Get_expList().Items ) {
                symNode := _symNode.(Nodes_NodeDownCast).ToNodes_Node()
                {
                    _refItemNode := TransUnit_analyzeExpOpSet__process_6535_(symNode)
                    if _refItemNode != nil {
                        refItemNode := _refItemNode.(*Nodes_ExpRefItemNode)
                        listRefItemNode = refItemNode
                        
                        if symNodeList.FP.Get_expList().Len() > 1{
                            self.FP.addErrMess(refItemNode.FP.Get_pos(), "When left-value includes 'list[i]', left-value must be single.")
                        }
                        break
                    }
                }
            }
        } else {
            listRefItemNode = TransUnit_analyzeExpOpSet__process_6535_(exp)
            
        }
    }
    var expList *Nodes_ExpListNode
    expList = self.FP.analyzeExpList(false, false, false, nil, expectTypeList, nil)
    if Lns_op_not(expList.FP.CanBeRight(self.processInfo)){
        self.FP.addErrMess(expList.FP.Get_pos(), Lns_getVM().String_format("this node can not be r-value. -- %s", []LnsAny{Nodes_getNodeKindName(expList.FP.Get_kind())}))
    }
    var workList LnsAny
    var expTypeList *LnsList
    _,_,workList,expTypeList = self.FP.checkMatchType("= operator", opeToken.Pos, exp.FP.Get_expTypeList(), expList, true, false, nil)
    if workList != nil{
        workList_10865 := workList.(*Nodes_ExpListNode)
        expList = workList_10865
        
    }
    var initSymSet *LnsSet
    initSymSet = NewLnsSet([]LnsAny{})
    var symbolList *LnsList
    symbolList = NewLnsList([]LnsAny{})
    for _index, _symbolInfo := range( exp.FP.GetSymbolInfo().Items ) {
        index := _index + 1
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        symbolList.Insert(Ast_SymbolInfo2Stem(symbolInfo))
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(symbolInfo.FP.Get_mutable())) &&
            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_hasValueFlag()) ).(bool)){
            if self.validMutControl{
                self.FP.addErrMess(opeToken.Pos, Lns_getVM().String_format("this is not mutable variable. -- %s", []LnsAny{symbolInfo.FP.Get_name()}))
            }
        }
        self.tentativeSymbol.FP.ModSym(symbolInfo)
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( index <= expTypeList.Len()) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(symbolInfo.FP.Get_hasValueFlag())) ).(bool)){
            if _switch66659 := symbolInfo.FP.Get_kind(); _switch66659 == Ast_SymbolKind__Var {
                if symbolInfo.FP.Get_typeInfo() == Ast_builtinTypeEmpty{
                    var expType *Ast_TypeInfo
                    expType = expTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if _switch66590 := expType.FP.Get_kind(); _switch66590 == Ast_TypeInfoKind__DDD {
                        if expType.FP.Get_itemTypeInfoList().Len() > 0{
                            expType = expType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo()
                            
                        }
                    } else if _switch66590 == Ast_TypeInfoKind__List || _switch66590 == Ast_TypeInfoKind__Array || _switch66590 == Ast_TypeInfoKind__Set || _switch66590 == Ast_TypeInfoKind__Map {
                        var workPos *Types_Position
                        if index <= expList.FP.Get_expList().Len(){
                            workPos = expList.FP.Get_expList().GetAt(index).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_pos()
                            
                        } else { 
                            workPos = opeToken.Pos
                            
                        }
                        self.FP.checkLiteralEmptyCollection(workPos, symbolInfo.FP.Get_name(), expType)
                    }
                    symbolInfo.FP.Set_typeInfo(expType)
                }
                if Lns_op_not(self.tentativeSymbol.FP.Regist(symbolInfo, exp.FP.Get_pos())){
                    self.FP.addErrMess(opeToken.Pos, Lns_getVM().String_format("can't access in this scope. -- %s", []LnsAny{symbolInfo.FP.Get_name()}))
                }
                initSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
            } else if _switch66659 == Ast_SymbolKind__Mbr {
                initSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
            }
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() != Ast_SymbolKind__Var) ||
            Lns_GetEnv().SetStackVal( self.scope.FP.GetNamespaceTypeInfo() == symbolInfo.FP.Get_scope().FP.GetNamespaceTypeInfo()) ).(bool){
            symbolInfo.FP.UpdateValue(exp.FP.Get_pos())
        }
    }
    if listRefItemNode != nil{
        listRefItemNode_10887 := listRefItemNode.(*Nodes_ExpRefItemNode)
        var index LnsAny
        {
            _indexNode := listRefItemNode_10887.FP.Get_index()
            if _indexNode != nil {
                indexNode := _indexNode.(*Nodes_Node)
                index = &Nodes_IndexVal__NodeIdx{indexNode}
                
            } else {
                index = &Nodes_IndexVal__SymIdx{Lns_unwrap( listRefItemNode_10887.FP.Get_symbol()).(string)}
                
            }
        }
        return &Nodes_ExpSetItemNode_create(self.nodeManager, exp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), listRefItemNode_10887.FP.Get_val(), index, &expList.Nodes_Node).Nodes_Node
    }
    return &Nodes_ExpSetValNode_create(self.nodeManager, exp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), exp, expList, symbolList, initSymSet).Nodes_Node
}

// 2917: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOpEquals
func (self *TransUnit_TransUnit) analyzeExpOpEquals(pos *Types_Position,opToken *Types_Token,exp1 *Nodes_Node,exp2 *Nodes_Node)(*Nodes_Node, *Nodes_Node) {
    var exp1Type *Ast_TypeInfo
    exp1Type = exp1.FP.Get_expType()
    var exp2Type *Ast_TypeInfo
    exp2Type = exp2.FP.Get_expType()
    if Lns_isCondTrue( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(exp1Type.FP.CanEvalWith(self.processInfo, exp2Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(exp2Type.FP.CanEvalWith(self.processInfo, exp1Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool))){
        self.FP.addErrMess(opToken.Pos, Lns_getVM().String_format("not compatible type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(self.typeNameCtrl, nil, nil), exp2Type.FP.GetTxt(self.typeNameCtrl, nil, nil)}))
        return exp1, exp2
    }
    {
        __exp := Nodes_NewAlgeValNodeDownCastF(exp1.FP)
        if __exp != nil {
            _exp := __exp.(*Nodes_NewAlgeValNode)
            if _exp.FP.Get_paramList().Len() > 0{
                self.FP.addErrMess(exp1.FP.Get_pos(), "can't compare alge.")
                return exp1, exp2
            }
        }
    }
    {
        __exp := Nodes_NewAlgeValNodeDownCastF(exp2.FP)
        if __exp != nil {
            _exp := __exp.(*Nodes_NewAlgeValNode)
            if _exp.FP.Get_paramList().Len() > 0{
                self.FP.addErrMess(exp2.FP.Get_pos(), "can't compare alge.")
                return exp1, exp2
            }
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeBool, nil, nil)) &&
        Lns_GetEnv().SetStackVal( exp2Type.FP.Equals(self.processInfo, Ast_builtinTypeBool, nil, nil)) &&
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( exp1.FP.Get_kind() == Nodes_NodeKind_get_LiteralBool()) ||
            Lns_GetEnv().SetStackVal( exp2.FP.Get_kind() == Nodes_NodeKind_get_LiteralBool()) ).(bool))) ).(bool)){
        self.FP.addWarnMess(exp1.FP.Get_pos(), "this operation is deprecated.")
        return exp1, exp2
    }
    var nonNilType1 *Ast_TypeInfo
    nonNilType1 = TransUnit_analyzeExpOpEquals__getType_6565_(exp1Type)
    var nonNilType2 *Ast_TypeInfo
    nonNilType2 = TransUnit_analyzeExpOpEquals__getType_6565_(exp2Type)
    if nonNilType1 != nonNilType2{
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( nonNilType1.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
            Lns_GetEnv().SetStackVal( nonNilType1.FP.Get_kind() == Ast_TypeInfoKind__IF) ).(bool){
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( nonNilType2.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( nonNilType2.FP.Get_kind() == Ast_TypeInfoKind__IF) ).(bool){
                if nonNilType1.FP.IsInheritFrom(self.processInfo, nonNilType2, nil){
                    exp1 = &Nodes_ExpCastNode_create(self.nodeManager, exp1.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp2Type.FP.Get_nonnilableType())}), exp1, exp2Type, Nodes_CastKind__Implicit).Nodes_Node
                    
                } else { 
                    exp2 = &Nodes_ExpCastNode_create(self.nodeManager, exp2.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp1Type.FP.Get_nonnilableType())}), exp2, exp1Type, Nodes_CastKind__Implicit).Nodes_Node
                    
                }
            } else { 
                exp1 = &Nodes_ExpCastNode_create(self.nodeManager, exp1.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp1Type)}), exp1, Ast_builtinTypeStem, Nodes_CastKind__Implicit).Nodes_Node
                
            }
        } else { 
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( nonNilType2.FP.Get_kind() == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( nonNilType2.FP.Get_kind() == Ast_TypeInfoKind__IF) ).(bool){
                exp2 = &Nodes_ExpCastNode_create(self.nodeManager, exp2.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp2Type)}), exp2, Ast_builtinTypeStem, Nodes_CastKind__Implicit).Nodes_Node
                
            }
        }
    }
    return exp1, exp2
}

// 3008: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOp2
func (self *TransUnit_TransUnit) analyzeExpOp2(firstToken *Types_Token,exp *Nodes_Node,prevOpLevel LnsAny) *Nodes_Node {
    for  {
        var opToken *Types_Token
        opToken = self.FP.GetTokenNoErr()
        var opTxt string
        opTxt = opToken.Txt
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( opToken.Txt == "@@") ||
            Lns_GetEnv().SetStackVal( opToken.Txt == "@@@") ||
            Lns_GetEnv().SetStackVal( opToken.Txt == "@@=") ).(bool){
            exp = self.FP.analyzeExpCast(firstToken, opTxt, exp)
            
        } else if opToken.Kind == Types_TokenKind__Ope{
            if Parser_isOp2(opTxt){
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( opTxt != "=") &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(exp.FP.CanBeRight(self.processInfo))) ).(bool)){
                    self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("This can't evaluate for '%s' -- %s", []LnsAny{opTxt, Nodes_getNodeKindName(exp.FP.Get_kind())}))
                }
                var opLevel LnsInt
                
                {
                    _opLevel := TransUnit_op2levelMap.Items[opTxt]
                    if _opLevel == nil{
                        self.FP.Error(Lns_getVM().String_format("unknown op -- %s %s", []LnsAny{opTxt, prevOpLevel}))
                    } else {
                        opLevel = _opLevel.(LnsInt)
                    }
                }
                {
                    __exp := prevOpLevel
                    if __exp != nil {
                        _exp := __exp.(LnsInt)
                        if opLevel <= _exp{
                            self.FP.Pushback()
                            return exp
                        }
                    }
                }
                var expectTypeList *LnsList
                expectTypeList = NewLnsList([]LnsAny{})
                for _, _exp1Type := range( exp.FP.Get_expTypeList().Items ) {
                    exp1Type := _exp1Type.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    var prefixExpType *Ast_TypeInfo
                    prefixExpType = exp1Type
                    if prefixExpType.FP.Get_nilable(){
                        prefixExpType = prefixExpType.FP.Get_nonnilableType()
                        
                    }
                    var expectType *Ast_TypeInfo
                    expectType = Ast_builtinTypeNone
                    {
                        __exp := Ast_EnumTypeInfoDownCastF(prefixExpType.FP.Get_srcTypeInfo().FP.Get_aliasSrc().FP)
                        if __exp != nil {
                            _exp := __exp.(*Ast_EnumTypeInfo)
                            expectType = &_exp.Ast_TypeInfo
                            
                        }
                    }
                    {
                        __exp := Ast_AlgeTypeInfoDownCastF(prefixExpType.FP.Get_srcTypeInfo().FP)
                        if __exp != nil {
                            _exp := __exp.(*Ast_AlgeTypeInfo)
                            expectType = &_exp.Ast_TypeInfo
                            
                        }
                    }
                    expectTypeList.Insert(Ast_TypeInfo2Stem(expectType))
                }
                if opTxt == "="{
                    return self.FP.analyzeExpOpSet(exp, opToken, expectTypeList)
                }
                exp = self.FP.MultiTo1(exp)
                
                var exp2 *Nodes_Node
                exp2 = self.FP.analyzeExp(false, false, false, opLevel, expectTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                exp2 = self.FP.MultiTo1(exp2)
                
                if Lns_op_not(exp2.FP.CanBeRight(self.processInfo)){
                    self.FP.addErrMess(exp2.FP.Get_pos(), Lns_getVM().String_format("This can't evaluate for '%s' -- %s", []LnsAny{opTxt, Nodes_getNodeKindName(exp2.FP.Get_kind())}))
                }
                var retType *Ast_TypeInfo
                retType = Ast_builtinTypeNone
                if Lns_op_not(exp2.FP.CanBeRight(self.processInfo)){
                    self.FP.addErrMess(exp2.FP.Get_pos(), Lns_getVM().String_format("this node can not be r-value. -- %s", []LnsAny{Nodes_getNodeKindName(exp2.FP.Get_kind())}))
                }
                var exp1Type *Ast_TypeInfo
                exp1Type = exp.FP.Get_expType()
                var exp2Type *Ast_TypeInfo
                exp2Type = exp2.FP.Get_expType()
                if exp1Type.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                    {
                        _dddType := Ast_DDDTypeInfoDownCastF(exp1Type.FP)
                        if _dddType != nil {
                            dddType := _dddType.(*Ast_DDDTypeInfo)
                            exp = &Nodes_ExpMultiTo1Node_create(self.nodeManager, exp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddType.FP.Get_typeInfo().FP.Get_nilableTypeInfo())}), exp).Nodes_Node
                            
                        }
                    }
                }
                if exp2Type.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                    {
                        _dddType := Ast_DDDTypeInfoDownCastF(exp2Type.FP)
                        if _dddType != nil {
                            dddType := _dddType.(*Ast_DDDTypeInfo)
                            exp2 = &Nodes_ExpMultiTo1Node_create(self.nodeManager, exp2.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddType.FP.Get_typeInfo().FP.Get_nilableTypeInfo())}), exp2).Nodes_Node
                            
                        }
                    }
                }
                if _switch69164 := opTxt; _switch69164 == "or" {
                    var is3op bool
                    {
                        _opExpType := Ast_AndExpTypeInfoDownCastF(exp1Type.FP)
                        if _opExpType != nil {
                            opExpType := _opExpType.(*Ast_AndExpTypeInfo)
                            exp1Type = opExpType.FP.Get_exp2()
                            
                            is3op = true
                            
                        } else {
                            is3op = false
                            
                        }
                    }
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeBool, nil, nil))) &&
                        Lns_GetEnv().SetStackVal( Lns_op_not(exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeStem, nil, nil))) &&
                        Lns_GetEnv().SetStackVal( Lns_op_not(exp1Type.FP.Get_nilable())) ).(bool)){
                        if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush((Nodes_ExpOp2NodeDownCastF(exp.FP))) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_ExpOp2Node).FP.Get_op()})&&
                        Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Types_Token).Txt)) == "and"{
                        } else { 
                            self.FP.addWarnMess(exp.FP.Get_pos(), Lns_getVM().String_format("this value never be 'false' -- %s", []LnsAny{exp1Type.FP.GetTxt(nil, nil, nil)}))
                        }
                    }
                    if exp1Type.FP.Equals(self.processInfo, exp2Type, nil, nil){
                        retType = exp1Type
                        
                    } else if Lns_car(exp1Type.FP.CanEvalWith(self.processInfo, exp2Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
                        retType = exp1Type
                        
                    } else if Lns_car(exp2Type.FP.CanEvalWith(self.processInfo, exp1Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
                        retType = exp2Type
                        
                    } else if exp2Type.FP.Equals(self.processInfo, Ast_builtinTypeNil, nil, nil){
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( is3op) ||
                            Lns_GetEnv().SetStackVal( exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeBool, nil, nil)) ).(bool){
                            retType = exp1Type.FP.Get_nilableTypeInfo()
                            
                        } else { 
                            retType = exp1Type
                            
                        }
                    } else if exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeNil, nil, nil){
                        retType = exp2Type
                        
                    } else { 
                        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( exp1Type.FP.Get_nilable()) &&
                            Lns_GetEnv().SetStackVal( exp2Type.FP.Get_nilable()) ).(bool)){
                            retType = Ast_builtinTypeStem_
                            
                        } else if exp2Type.FP.Get_nilable(){
                            retType = Ast_builtinTypeStem_
                            
                        } else if exp1Type.FP.Get_nilable(){
                            retType = Ast_builtinTypeStem
                            
                        } else { 
                            retType = Ast_builtinTypeStem
                            
                        }
                    }
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( retType.FP.Get_nilable()) &&
                        Lns_GetEnv().SetStackVal( Lns_op_not(exp2Type.FP.Get_nilable())) ).(bool)){
                        retType = retType.FP.Get_nonnilableType()
                        
                    }
                } else if _switch69164 == "and" {
                    _ = self.FP.getToken(nil)
                    self.FP.Pushback()
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeBool, nil, nil))) &&
                        Lns_GetEnv().SetStackVal( Lns_op_not(exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeStem, nil, nil))) &&
                        Lns_GetEnv().SetStackVal( Lns_op_not(exp1Type.FP.Get_nilable())) ).(bool)){
                        self.FP.addWarnMess(exp.FP.Get_pos(), "this value never be 'false'")
                    } else if exp2.FP.Get_kind() == Nodes_NodeKind_get_LiteralBool(){
                        {
                            _literal := TransUnit_convExp68291(Lns_2DDD(exp2.FP.GetLiteral()))
                            if _literal != nil {
                                literal := _literal
                                if Lns_op_not(Nodes_getLiteralObj(literal)){
                                    self.FP.addErrMess(exp2.FP.Get_pos(), "this value never be 'true'")
                                }
                            }
                        }
                    }
                    if exp1Type.FP.Get_nilable(){
                        if exp2Type.FP.Get_nilable(){
                            retType = exp2Type
                            
                        } else { 
                            retType = exp2Type.FP.Get_nilableTypeInfo()
                            
                        }
                    } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeBool, nil, nil)) ||
                        Lns_GetEnv().SetStackVal( exp2Type.FP.Equals(self.processInfo, Ast_builtinTypeBool, nil, nil)) ).(bool){
                        if Lns_car(exp1Type.FP.CanEvalWith(self.processInfo, exp2Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
                            retType = exp1Type
                            
                        } else if Lns_car(exp2Type.FP.CanEvalWith(self.processInfo, exp1Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
                            retType = exp2Type
                            
                        } else { 
                            if exp2Type.FP.Get_nilable(){
                                retType = Ast_builtinTypeStem_
                                
                            } else { 
                                retType = Ast_builtinTypeStem
                                
                            }
                        }
                    } else if exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeStem, nil, nil){
                        retType = Ast_builtinTypeStem
                        
                    } else { 
                        retType = exp2Type
                        
                    }
                    retType = &NewAst_AndExpTypeInfo(self.processInfo, exp1Type, exp2Type, retType).Ast_TypeInfo
                    
                } else if _switch69164 == "<" || _switch69164 == ">" || _switch69164 == "<=" || _switch69164 == ">=" {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_car(Ast_builtinTypeString.FP.CanEvalWith(self.processInfo, exp1Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) &&
                        Lns_GetEnv().SetStackVal( Lns_car(Ast_builtinTypeString.FP.CanEvalWith(self.processInfo, exp2Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ||
                        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(self.processInfo, exp1Type, Ast_CanEvalType__Comp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ||
                            Lns_GetEnv().SetStackVal( Lns_car(Ast_builtinTypeReal.FP.CanEvalWith(self.processInfo, exp1Type, Ast_CanEvalType__Comp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ).(bool))) &&
                        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(self.processInfo, exp2Type, Ast_CanEvalType__Comp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ||
                            Lns_GetEnv().SetStackVal( Lns_car(Ast_builtinTypeReal.FP.CanEvalWith(self.processInfo, exp2Type, Ast_CanEvalType__Comp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ).(bool))) ).(bool){
                    } else { 
                        self.FP.addErrMess(opToken.Pos, Lns_getVM().String_format("no numeric type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(self.typeNameCtrl, nil, nil), exp2Type.FP.GetTxt(self.typeNameCtrl, nil, nil)}))
                    }
                    retType = Ast_builtinTypeBool
                    
                } else if _switch69164 == "~=" || _switch69164 == "==" {
                    exp, exp2 = self.FP.analyzeExpOpEquals(firstToken.Pos, opToken, exp, exp2)
                    
                    retType = Ast_builtinTypeBool
                    
                } else if _switch69164 == "^" || _switch69164 == "|" || _switch69164 == "~" || _switch69164 == "&" || _switch69164 == "|<<" || _switch69164 == "|>>" {
                    if self.targetLuaVer.FP.Get_hasBitOp() == LuaVer_BitOp__Cant{
                        self.FP.addErrMess(opToken.Pos, "this lua version can't use bit operand.")
                    }
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(self.processInfo, exp1Type, Ast_CanEvalType__Logical, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(self.processInfo, exp2Type, Ast_CanEvalType__Logical, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool){
                        self.FP.addErrMess(opToken.Pos, Lns_getVM().String_format("no int type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(nil, nil, nil), exp2Type.FP.GetTxt(nil, nil, nil)}))
                    }
                    retType = Ast_builtinTypeInt
                    
                } else if _switch69164 == ".." {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil))) ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(exp2Type.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil))) ).(bool){
                        self.FP.addErrMess(opToken.Pos, Lns_getVM().String_format("no string type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(nil, nil, nil), exp2Type.FP.GetTxt(nil, nil, nil)}))
                    }
                    retType = Ast_builtinTypeString
                    
                } else if _switch69164 == "+" || _switch69164 == "-" || _switch69164 == "*" || _switch69164 == "/" || _switch69164 == "%" {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(self.processInfo, exp1Type, Ast_CanEvalType__Math, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) &&
                            Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeReal.FP.CanEvalWith(self.processInfo, exp1Type, Ast_CanEvalType__Math, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool))) ||
                        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(self.processInfo, exp2Type, Ast_CanEvalType__Math, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) &&
                            Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeReal.FP.CanEvalWith(self.processInfo, exp2Type, Ast_CanEvalType__Math, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool))) ).(bool){
                        self.FP.addErrMess(opToken.Pos, Lns_getVM().String_format("no numeric type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(nil, nil, nil), exp2Type.FP.GetTxt(nil, nil, nil)}))
                    }
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeReal, nil, nil)) ||
                        Lns_GetEnv().SetStackVal( exp2Type.FP.Equals(self.processInfo, Ast_builtinTypeReal, nil, nil)) ).(bool){
                        retType = Ast_builtinTypeReal
                        
                        if exp1Type.FP.Equals(self.processInfo, Ast_builtinTypeInt, nil, nil){
                            exp = &Nodes_ExpCastNode_create(self.nodeManager, exp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeReal)}), exp, Ast_builtinTypeReal, Nodes_CastKind__Implicit).Nodes_Node
                            
                        } else if exp2Type.FP.Equals(self.processInfo, Ast_builtinTypeInt, nil, nil){
                            exp2 = &Nodes_ExpCastNode_create(self.nodeManager, exp2.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeReal)}), exp2, Ast_builtinTypeReal, Nodes_CastKind__Implicit).Nodes_Node
                            
                        }
                    } else { 
                        retType = Ast_builtinTypeInt
                        
                    }
                } else {
                    self.FP.Error("unknown op " + opTxt)
                }
                var threading bool
                if _switch69197 := opTxt; _switch69197 == "and" || _switch69197 == "or" {
                    threading = self.FP.checkThreading(firstToken.Pos)
                    
                } else {
                    threading = false
                    
                }
                exp = &Nodes_ExpOp2Node_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(retType)}), opToken, threading, self.nodeManager.FP.MultiTo1(exp), self.nodeManager.FP.MultiTo1(exp2)).Nodes_Node
                
            } else { 
                self.FP.Error("illegal op")
            }
        } else { 
            self.FP.Pushback()
            return exp
        }
    }
// insert a dummy
    return nil
}

// 3369: decl @lune.@base.@TransUnit.TransUnit.analyzeExpMacroStat
func (self *TransUnit_TransUnit) analyzeExpMacroStat(firstToken *Types_Token) *Nodes_ExpMacroStatNode {
    var expStrList *LnsList
    expStrList = NewLnsList([]LnsAny{})
    self.FP.checkNextToken("{")
    var braceCount LnsInt
    braceCount = 0
    var prevToken *Types_Token
    prevToken = firstToken
    var errMessList *LnsList
    errMessList = NewLnsList([]LnsAny{})
    for  {
        var token *Types_Token
        token = self.FP.getToken(nil)
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == ",,") ||
            Lns_GetEnv().SetStackVal( token.Txt == ",,,") ||
            Lns_GetEnv().SetStackVal( token.Txt == ",,,,") ).(bool){
            var exp *Nodes_Node
            exp = self.FP.analyzeExp(false, true, false, Lns_unwrap( TransUnit_op1levelMap.Items[token.Txt]).(LnsInt), nil)
            var literalStr *Nodes_LiteralStringNode
            literalStr = self.macroCtrl.FP.ExpandSymbol(self.FP, token, exp, self.nodeManager, errMessList)
            for _, _errMess := range( errMessList.Items ) {
                errMess := _errMess.(Macro_ErrorMessDownCast).ToMacro_ErrorMess()
                self.FP.addErrMess(errMess.Pos, errMess.Mess)
            }
            expStrList.Insert(Nodes_LiteralStringNode2Stem(literalStr))
        } else { 
            if token.Txt == "{"{
                braceCount = braceCount + 1
                
            } else if token.Txt == "}"{
                if braceCount == 0{
                    break
                }
                braceCount = braceCount - 1
                
            }
            var format string
            format = "' %s'"
            var consecutive bool
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( prevToken == firstToken) ||
                Lns_GetEnv().SetStackVal( token.Consecutive) ).(bool){
                format = "'%s'"
                
                consecutive = true
                
            } else { 
                consecutive = false
                
            }
            var newToken *Types_Token
            newToken = NewTypes_Token(token.Kind, Lns_getVM().String_format(format, []LnsAny{token.Txt}), token.Pos, consecutive, nil)
            var literalStr *Nodes_LiteralStringNode
            literalStr = Nodes_LiteralStringNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), newToken, nil, nil, false)
            expStrList.Insert(Nodes_LiteralStringNode2Stem(literalStr))
        }
        prevToken = token
        
    }
    return Nodes_ExpMacroStatNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeStat)}), expStrList)
}

// 3429: decl @lune.@base.@TransUnit.TransUnit.analyzeSuper
func (self *TransUnit_TransUnit) analyzeSuper(firstToken *Types_Token) *Nodes_Node {
    self.FP.checkNextToken("(")
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var expList LnsAny
    expList = nil
    if nextToken.Txt != ")"{
        self.FP.Pushback()
        expList = self.FP.analyzeExpList(false, false, false, nil, nil, nil)
        
        self.FP.checkNextToken(")")
    }
    self.FP.checkNextToken(";")
    var classType *Ast_TypeInfo
    classType = self.FP.getCurrentClass()
    var currentFunc *Ast_TypeInfo
    currentFunc = self.FP.getCurrentNamespaceTypeInfo()
    if currentFunc.FP.Get_kind() == Ast_TypeInfoKind__Method{
        var superType *Ast_TypeInfo
        superType = classType.FP.Get_baseTypeInfo()
        if superType.FP.Equals(self.processInfo, Ast_headTypeInfo, nil, nil){
            self.FP.addErrMess(firstToken.Pos, "This class doesn't have super-class.")
        } else { 
            if currentFunc.FP.Get_rawTxt() == "__init"{
                var superScope *Ast_Scope
                
                {
                    _superScope := superType.FP.Get_scope()
                    if _superScope == nil{
                        self.FP.Error("not found super scope")
                    } else {
                        superScope = _superScope.(*Ast_Scope)
                    }
                }
                var superCtorType *Ast_TypeInfo
                
                {
                    _superCtorType := superScope.FP.GetTypeInfoChild("__init")
                    if _superCtorType == nil{
                        self.FP.Error("not found super '__init'")
                    } else {
                        superCtorType = _superCtorType.(*Ast_TypeInfo)
                    }
                }
                self.FP.checkMatchValType(firstToken.Pos, superCtorType, expList, NewLnsList([]LnsAny{}), classType)
                return &Nodes_ExpCallSuperCtorNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), superType, superCtorType, expList).Nodes_Node
            } else { 
                {
                    _superFunc := (Lns_unwrap( superType.FP.Get_scope()).(*Ast_Scope)).FP.GetTypeInfoField(currentFunc.FP.Get_rawTxt(), true, self.scope, self.scopeAccess)
                    if _superFunc != nil {
                        superFunc := _superFunc.(*Ast_TypeInfo)
                        if superFunc.FP.Get_abstractFlag(){
                            self.FP.addErrMess(firstToken.Pos, "super is abstract.")
                        }
                        self.FP.checkMatchValType(firstToken.Pos, superFunc, expList, NewLnsList([]LnsAny{}), classType)
                        var exp *Nodes_ExpCallSuperNode
                        exp = Nodes_ExpCallSuperNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), superType, superFunc, expList)
                        return &Nodes_StmtExpNode_create(self.nodeManager, exp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), exp.FP.Get_expTypeList(), &exp.Nodes_Node).Nodes_Node
                    }
                }
                self.FP.addErrMess(firstToken.Pos, "this is not override method.")
                return self.FP.createNoneNode(firstToken.Pos)
            }
        }
    }
    self.FP.addErrMess(firstToken.Pos, "super can't call here.")
    return self.FP.createNoneNode(firstToken.Pos)
}

// 3494: decl @lune.@base.@TransUnit.TransUnit.analyzeUnwrap
func (self *TransUnit_TransUnit) analyzeUnwrap(firstToken *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    var continueFlag bool
    nextToken,continueFlag = self.FP.getContinueToken()
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(continueFlag)) ||
        Lns_GetEnv().SetStackVal( nextToken.Txt != "!") ).(bool){
        self.FP.Pushback()
        self.FP.PushbackToken(firstToken)
        var exp *Nodes_Node
        exp = self.FP.analyzeExp(false, false, false, nil, nil)
        self.FP.checkNextToken(";")
        if Lns_op_not(exp.FP.Get_expType().FP.Get_nilable()){
            self.FP.addErrMess(exp.FP.Get_pos(), "this value is not nilable.")
        }
        return &Nodes_StmtExpNode_create(self.nodeManager, nextToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), exp).Nodes_Node
    }
    self.FP.Pushback()
    return self.FP.analyzeDeclVar(Nodes_DeclVarMode__Unwrap, Ast_AccessMode__Local, firstToken)
}

// 3514: decl @lune.@base.@TransUnit.TransUnit.analyzeExpUnwrap
func (self *TransUnit_TransUnit) analyzeExpUnwrap(firstToken *Types_Token) *Nodes_Node {
    var expNode *Nodes_Node
    expNode = self.FP.analyzeExpOneRVal(false, true, nil, nil)
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var insNode LnsAny
    insNode = nil
    if nextToken.Txt == "default"{
        insNode = self.FP.analyzeExpOneRVal(false, false, nil, nil)
        
    } else { 
        self.FP.Pushback()
    }
    var unwrapType *Ast_TypeInfo
    unwrapType = Ast_builtinTypeStem_
    var expType *Ast_TypeInfo
    expType = expNode.FP.Get_expType()
    if Lns_op_not(expType.FP.Get_nilable()){
        unwrapType = expType
        
        self.FP.addErrMess(expNode.FP.Get_pos(), Lns_getVM().String_format("this exp is not nilable -- %s", []LnsAny{expType.FP.GetTxt(nil, nil, nil)}))
    } else if expType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
        if expType.FP.Get_itemTypeInfoList().Len() > 0{
            unwrapType = expType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nonnilableType()
            
        } else { 
            unwrapType = Ast_builtinTypeStem
            
        }
    } else { 
        unwrapType = expType.FP.Get_nonnilableType()
        
    }
    if insNode != nil{
        insNode_11109 := insNode.(*Nodes_Node)
        var insType *Ast_TypeInfo
        insType = insNode_11109.FP.Get_expType()
        if insType.FP.Get_nilable(){
            self.FP.addErrMess(insNode_11109.FP.Get_pos(), Lns_getVM().String_format("default can't use nilable -- %s", []LnsAny{insType.FP.GetTxt(nil, nil, nil)}))
        }
        var alt2type *LnsMap
        alt2type = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(false)
        if Lns_op_not(Lns_car(unwrapType.FP.CanEvalWith(self.processInfo, insType, Ast_CanEvalType__SetOp, alt2type)).(bool)){
            if Lns_op_not(Lns_car(insType.FP.CanEvalWith(self.processInfo, unwrapType, Ast_CanEvalType__SetOp, alt2type)).(bool)){
                unwrapType = Ast_builtinTypeStem
                
            } else { 
                unwrapType = insType
                
            }
        }
    }
    self.helperInfo.UseUnwrapExp = true
    
    if Ast_isExtType(expType.FP.Get_nonnilableType()){
        switch _exp70479 := self.processInfo.FP.CreateLuaval(unwrapType, false).(type) {
        case *Ast_LuavalResult__OK:
        work := _exp70479.Val1
            unwrapType = work
            
        case *Ast_LuavalResult__Err:
        err := _exp70479.Val1
            self.FP.addErrMess(firstToken.Pos, err)
        }
    }
    return &Nodes_ExpUnwrapNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(unwrapType)}), expNode, insNode).Nodes_Node
}

// 3597: decl @lune.@base.@TransUnit.TransUnit.analyzeStrConst
func (self *TransUnit_TransUnit) analyzeStrConst(firstToken *Types_Token,token *Types_Token) *Nodes_Node {
    var exp *Nodes_Node
    var nextToken *Types_Token
    nextToken = self.FP.getToken(true)
    if nextToken.Kind != Types_TokenKind__Eof{
        var param LnsAny
        var dddParam LnsAny
        if nextToken.Txt == "("{
            var argNodeList *Nodes_ExpListNode
            argNodeList = self.FP.analyzeExpList(false, false, false, nil, nil, nil)
            param = argNodeList
            
            var workExpList LnsAny
            _,_,workExpList = TransUnit_convExp70613(Lns_2DDD(self.FP.checkMatchType("str constructor", firstToken.Pos, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeDDD)}), argNodeList, false, false, nil)))
            if workExpList != nil{
                workExpList_11139 := workExpList.(*Nodes_ExpListNode)
                dddParam = workExpList_11139
                
            } else {
                dddParam = nil
                
            }
            self.FP.checkNextToken(")")
            nextToken = self.FP.getToken(true)
            
            if param != nil{
                param_11142 := param.(*Nodes_ExpListNode)
                self.FP.checkStringFormat(token.Pos, token.Txt, param_11142.FP.Get_expTypeList())
            }
        } else { 
            param = nil
            
            dddParam = nil
            
        }
        var threading bool
        if Lns_isCondTrue( param){
            threading = self.FP.checkThreading(firstToken.Pos)
            
        } else { 
            threading = false
            
        }
        var workExp *Nodes_Node
        workExp = &Nodes_LiteralStringNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), token, param, dddParam, threading).Nodes_Node
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( nextToken.Txt == "[") ||
            Lns_GetEnv().SetStackVal( nextToken.Txt == "$[") ).(bool){
            exp = self.FP.analyzeExpRefItem(nextToken, workExp, nextToken.Txt == "$[")
            
        } else { 
            exp = workExp
            
            if nextToken.Kind != Types_TokenKind__Eof{
                self.FP.Pushback()
            }
        }
    } else { 
        exp = &Nodes_LiteralStringNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), token, nil, nil, false).Nodes_Node
        
    }
    return exp
}

// 3661: decl @lune.@base.@TransUnit.TransUnit.analyzeExp
func (self *TransUnit_TransUnit) analyzeExp(allowNoneType bool,skipOp2Flag bool,canLeftExp bool,prevOpLevel LnsAny,expectType LnsAny) *Nodes_Node {
    var firstToken *Types_Token
    firstToken = self.FP.getToken(nil)
    var processsExpectExp func(token *Types_Token,orgExpectType *Ast_TypeInfo) *Nodes_Node
    processsExpectExp = func(token *Types_Token,orgExpectType *Ast_TypeInfo) *Nodes_Node {
        {
            _enumTyepInfo := Ast_EnumTypeInfoDownCastF(orgExpectType.FP.Get_srcTypeInfo().FP.Get_aliasSrc().FP)
            if _enumTyepInfo != nil {
                enumTyepInfo := _enumTyepInfo.(*Ast_EnumTypeInfo)
                var nextToken *Types_Token
                nextToken = self.FP.getToken(nil)
                self.FP.checkEnumComp(nextToken, enumTyepInfo)
                {
                    _valInfo := enumTyepInfo.FP.GetEnumValInfo(nextToken.Txt)
                    if _valInfo != nil {
                        valInfo := _valInfo.(*Ast_EnumValInfo)
                        var aliasType LnsAny
                        aliasType = nil
                        var expType *Ast_TypeInfo
                        expType = &enumTyepInfo.Ast_TypeInfo
                        aliasType = self.importedAliasMap.Items[&enumTyepInfo.Ast_TypeInfo]
                        
                        if aliasType != nil{
                            aliasType_11175 := aliasType.(*Ast_AliasTypeInfo)
                            expType = &aliasType_11175.Ast_TypeInfo
                            
                        }
                        if orgExpectType.FP.Get_externalFlag(){
                            if Lns_op_not(self.importModuleSet.Has(Ast_TypeInfo2Stem(orgExpectType.FP.GetModule()))){
                                if Lns_op_not(aliasType){
                                    var fullname string
                                    fullname = orgExpectType.FP.GetFullName(self.typeNameCtrl, self.scope.FP, true)
                                    self.FP.addErrMess(token.Pos, Lns_getVM().String_format("need to import module -- %s (%s)", []LnsAny{fullname, enumTyepInfo.FP.GetTxt(nil, nil, nil)}))
                                }
                            }
                        }
                        var exp *Nodes_Node
                        exp = &Nodes_ExpOmitEnumNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(expType)}), nextToken, valInfo, aliasType, enumTyepInfo).Nodes_Node
                        return self.FP.analyzeExpCont(firstToken, exp, false, canLeftExp)
                    }
                }
                self.FP.Error(Lns_getVM().String_format("illegal enum val -- %s.%s", []LnsAny{orgExpectType.FP.GetTxt(nil, nil, nil), nextToken.Txt}))
            }
        }
        {
            _algeTyepInfo := Ast_AlgeTypeInfoDownCastF(orgExpectType.FP.Get_srcTypeInfo().FP)
            if _algeTyepInfo != nil {
                algeTyepInfo := _algeTyepInfo.(*Ast_AlgeTypeInfo)
                return &self.FP.analyzeNewAlge(firstToken, algeTyepInfo, nil).Nodes_Node
            }
        }
        self.FP.Error(Lns_getVM().String_format("illegal type for '.' -- %s", []LnsAny{orgExpectType.FP.GetTxt(nil, nil, nil)}))
    // insert a dummy
        return nil
    }
    var processsNewExp func(token *Types_Token) *Nodes_Node
    processsNewExp = func(token *Types_Token) *Nodes_Node {
        var exp *Nodes_Node
        exp = &self.FP.analyzeRefType(Ast_AccessMode__Local, false, false).Nodes_Node
        var classTypeInfo *Ast_TypeInfo
        classTypeInfo = exp.FP.Get_expType()
        if _switch71279 := classTypeInfo.FP.Get_kind(); _switch71279 == Ast_TypeInfoKind__Class || _switch71279 == Ast_TypeInfoKind__IF {
            if classTypeInfo.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil){
                self.FP.Error(Lns_getVM().String_format("'new' can't use this type -- %s", []LnsAny{classTypeInfo.FP.GetTxt(nil, nil, nil)}))
            }
        } else {
            self.FP.Error(Lns_getVM().String_format("'new' can't use this type -- %s", []LnsAny{classTypeInfo.FP.GetTxt(nil, nil, nil)}))
        }
        if classTypeInfo.FP.Get_externalFlag(){
            if _switch71311 := classTypeInfo.FP.Get_accessMode(); _switch71311 == Ast_AccessMode__Pri || _switch71311 == Ast_AccessMode__Local {
                self.FP.addErrMess(token.Pos, Lns_getVM().String_format("Can't access -- %s", []LnsAny{Ast_AccessMode_getTxt( classTypeInfo.FP.Get_accessMode())}))
            }
        }
        if classTypeInfo.FP.Get_abstractFlag(){
            self.FP.addErrMess(token.Pos, "abstract class can't new")
        }
        var classScope LnsAny
        classScope = classTypeInfo.FP.Get_scope()
        var initTypeInfo *Ast_TypeInfo
        
        {
            _initTypeInfo := (Lns_unwrap( classScope).(*Ast_Scope)).FP.GetTypeInfoChild("__init")
            if _initTypeInfo == nil{
                self.FP.Error("not found __init")
            } else {
                initTypeInfo = _initTypeInfo.(*Ast_TypeInfo)
            }
        }
        self.FP.checkNextToken("(")
        var nextToken *Types_Token
        nextToken = self.FP.getToken(nil)
        var argList LnsAny
        argList = nil
        if nextToken.Txt != ")"{
            self.FP.Pushback()
            argList = self.FP.analyzeExpList(false, false, false, nil, initTypeInfo.FP.Get_argTypeInfoList(), nil)
            
            self.FP.checkNextToken(")")
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( initTypeInfo.FP.Get_accessMode() == Ast_AccessMode__Pub) ||
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( initTypeInfo.FP.Get_accessMode() == Ast_AccessMode__Pro) &&
                Lns_GetEnv().SetStackVal( self.scope.FP.GetClassTypeInfo().FP.IsInheritFrom(self.processInfo, classTypeInfo, nil)) ).(bool))) ||
            Lns_GetEnv().SetStackVal( (self.scope.FP.GetClassTypeInfo() == classTypeInfo)) ||
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( initTypeInfo.FP.Get_accessMode() == Ast_AccessMode__Local) &&
                Lns_GetEnv().SetStackVal( initTypeInfo.FP.GetModule() == self.moduleType) ).(bool))) ).(bool){
        } else { 
            self.FP.addErrMess(token.Pos, Lns_getVM().String_format("can't access to __init of %s", []LnsAny{classTypeInfo.FP.GetTxt(nil, nil, nil)}))
        }
        var alt2type *LnsMap
        var newArgList LnsAny
        _,alt2type,newArgList = self.FP.checkMatchValType(exp.FP.Get_pos(), initTypeInfo, argList, classTypeInfo.FP.Get_itemTypeInfoList(), classTypeInfo)
        if classTypeInfo.FP.Get_itemTypeInfoList().Len() > 0{
            if classTypeInfo.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Alternate{
                var genTypeList *LnsList
                genTypeList = NewLnsList([]LnsAny{})
                var detect bool
                detect = true
                for _, _altType := range( classTypeInfo.FP.Get_itemTypeInfoList().Items ) {
                    altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    {
                        __exp := alt2type.Items[altType]
                        if __exp != nil {
                            _exp := __exp.(*Ast_TypeInfo)
                            genTypeList.Insert(Ast_TypeInfo2Stem(_exp))
                        } else {
                            self.FP.addErrMess(token.Pos, Lns_getVM().String_format("Can't new generic class. -- %s", []LnsAny{classTypeInfo.FP.GetTxt(nil, nil, nil)}))
                            detect = false
                            
                            break
                        }
                    }
                }
                if detect{
                    classTypeInfo = &self.processInfo.FP.CreateGeneric(classTypeInfo, genTypeList, self.moduleType).Ast_TypeInfo
                    
                }
            }
        }
        exp = &Nodes_ExpNewNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo)}), exp, initTypeInfo, newArgList).Nodes_Node
        
        exp = self.FP.analyzeExpCont(firstToken, exp, false, canLeftExp)
        
        return exp
    }
    var processOp1 func(token *Types_Token)(*Nodes_Node, bool)
    processOp1 = func(token *Types_Token)(*Nodes_Node, bool) {
        if token.Txt == "`"{
            return &self.FP.analyzeExpMacroStat(token).Nodes_Node, false
        }
        var exp *Nodes_Node
        exp = self.FP.analyzeExpOneRVal(false, true, Lns_unwrap( TransUnit_op1levelMap.Items[token.Txt]).(LnsInt), nil)
        var typeInfo *Ast_TypeInfo
        typeInfo = Ast_builtinTypeNone
        var macroExpFlag bool
        macroExpFlag = false
        var expType *Ast_TypeInfo
        expType = exp.FP.Get_expType()
        if expType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
            self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("... can't evaluate for '%s'.", []LnsAny{token.Txt}))
        }
        if _switch72198 := (token.Txt); _switch72198 == "-" {
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(expType.FP.Equals(self.processInfo, Ast_builtinTypeInt, nil, nil))) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(expType.FP.Equals(self.processInfo, Ast_builtinTypeReal, nil, nil))) ).(bool)){
                self.FP.addErrMess(token.Pos, Lns_getVM().String_format("unmatch type for \"-\" -- %s", []LnsAny{expType.FP.GetTxt(nil, nil, nil)}))
            }
            typeInfo = expType
            
        } else if _switch72198 == "#" {
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( expType.FP.Get_extedType().FP.Get_kind() != Ast_TypeInfoKind__List) &&
                Lns_GetEnv().SetStackVal( expType.FP.Get_extedType().FP.Get_kind() != Ast_TypeInfoKind__Array) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeString.FP.CanEvalWith(self.processInfo, expType, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool)){
                self.FP.addErrMess(token.Pos, Lns_getVM().String_format("unmatch type for \"#\" -- %s", []LnsAny{expType.FP.GetTxt(nil, nil, nil)}))
            }
            typeInfo = Ast_builtinTypeInt
            
        } else if _switch72198 == "not" {
            typeInfo = Ast_builtinTypeBool
            
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(expType.FP.Get_nilable())) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(expType.FP.Equals(self.processInfo, Ast_builtinTypeBool, nil, nil))) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(expType.FP.Equals(self.processInfo, Ast_builtinTypeStem, nil, nil))) &&
                Lns_GetEnv().SetStackVal( expType.FP.Get_kind() != Ast_TypeInfoKind__DDD) ).(bool)){
                self.FP.addErrMess(token.Pos, "this 'not' operand never be false")
            }
        } else if _switch72198 == ",," {
            macroExpFlag = true
            
            typeInfo = expType
            
        } else if _switch72198 == ",,," {
            macroExpFlag = true
            
            if Lns_op_not(expType.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil)){
                self.FP.Error("unmatch ,,, type, need string type")
            }
            typeInfo = Ast_builtinTypeSymbol
            
        } else if _switch72198 == ",,,," {
            macroExpFlag = true
            
            if Lns_op_not(expType.FP.Equals(self.processInfo, Ast_builtinTypeSymbol, nil, nil)){
                self.FP.Error("unmatch ,,, type, need symbol type")
            }
            typeInfo = Ast_builtinTypeString
            
        } else if _switch72198 == "`" {
            typeInfo = Ast_builtinTypeNone
            
        } else if _switch72198 == "~" {
            if Lns_op_not(expType.FP.Equals(self.processInfo, Ast_builtinTypeInt, nil, nil)){
                self.FP.addErrMess(token.Pos, Lns_getVM().String_format("unmatch type for \"~\" -- %s", []LnsAny{expType.FP.GetTxt(nil, nil, nil)}))
            }
            typeInfo = Ast_builtinTypeInt
            
        } else {
            self.FP.Error("unknown op1")
        }
        if macroExpFlag{
            var nextToken *Types_Token
            nextToken = self.FP.getToken(true)
            if nextToken.Txt != "~~"{
                self.FP.Pushback()
            }
        }
        exp = &Nodes_ExpOp1Node_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), token, self.macroCtrl.FP.Get_analyzeInfo().FP.Get_mode(), self.nodeManager.FP.MultiTo1(exp)).Nodes_Node
        
        return self.FP.analyzeExpOp2(firstToken, exp, prevOpLevel), true
    }
    var token *Types_Token
    token = firstToken
    var exp *Nodes_Node
    exp = self.FP.createNoneNode(firstToken.Pos)
    if token.Txt == "##"{
        if allowNoneType{
            self.FP.addErrMess(token.Pos, "illeal syntax -- ##")
        }
        return &Nodes_AbbrNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeAbbr)})).Nodes_Node
    }
    if token.Kind == Types_TokenKind__Dlmt{
        if token.Txt == "."{
            if expectType != nil{
                expectType_11252 := expectType.(*Ast_TypeInfo)
                var orgExpectType *Ast_TypeInfo
                orgExpectType = expectType_11252
                if orgExpectType.FP.Get_nilable(){
                    orgExpectType = orgExpectType.FP.Get_nonnilableType()
                    
                }
                exp = processsExpectExp(token, orgExpectType)
                
            } else {
                self.FP.Error("illegal '.'")
            }
        } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == "[") ||
            Lns_GetEnv().SetStackVal( token.Txt == "[@") ).(bool){
            exp = self.FP.analyzeListConst(token, expectType)
            
        } else if token.Txt == "(@"{
            exp = self.FP.analyzeSetConst(token, expectType)
            
        } else if token.Txt == "{"{
            exp = &self.FP.analyzeMapConst(token, expectType).Nodes_Node
            
        } else if token.Txt == "("{
            exp = self.FP.analyzeExp(false, false, false, nil, nil)
            
            self.FP.checkNextToken(")")
            if Lns_op_not(exp.FP.CanBeRight(self.processInfo)){
                self.FP.addErrMess(exp.FP.Get_pos(), Lns_getVM().String_format("can't be r-value in paren. -- %s", []LnsAny{Nodes_getNodeKindName(exp.FP.Get_kind())}))
            }
            exp = &Nodes_ExpParenNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp.FP.Get_expType())}), exp).Nodes_Node
            
            exp = self.FP.analyzeExpCont(firstToken, exp, false, canLeftExp)
            
        }
    }
    if token.Txt == "new"{
        exp = processsNewExp(token)
        
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( token.Kind == Types_TokenKind__Ope) &&
        Lns_GetEnv().SetStackVal( Parser_isOp1(token.Txt)) ).(bool)){
        var workExp *Nodes_Node
        var fin bool
        workExp,fin = processOp1(token)
        if fin{
            return workExp
        }
        exp = workExp
        
    }
    if token.Kind == Types_TokenKind__Int{
        exp = &Nodes_LiteralIntNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), token, Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(token.Txt, nil), 0)))).Nodes_Node
        
    } else if token.Kind == Types_TokenKind__Real{
        exp = &Nodes_LiteralRealNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeReal)}), token, (Lns_unwrapDefault( Lns_tonumber(token.Txt, nil), 0.0).(LnsReal))).Nodes_Node
        
    } else if token.Kind == Types_TokenKind__Char{
        var num LnsInt
        if len(token.Txt) == 1{
            num = LnsInt(token.Txt[1-1])
            
        } else { 
            num = Lns_unwrap( TransUnit_quotedChar2Code.Items[Lns_getVM().String_sub(token.Txt,2, 2)]).(LnsInt)
            
        }
        exp = &Nodes_LiteralCharNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeChar)}), token, num).Nodes_Node
        
    } else if token.Kind == Types_TokenKind__Str{
        exp = self.FP.analyzeStrConst(firstToken, token)
        
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( token.Kind == Types_TokenKind__Symb) &&
        Lns_GetEnv().SetStackVal( token.Txt == "__line__") ).(bool)){
        var lineNo LnsInt
        lineNo = token.Pos.LineNo
        if self.macroCtrl.FP.Get_analyzeInfo().FP.Get_mode() != Nodes_MacroMode__None{
            lineNo = self.macroCtrl.FP.Get_macroCallLineNo()
            
        }
        exp = &Nodes_LiteralIntNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), NewTypes_Token(Types_TokenKind__Int, Lns_getVM().String_format("%d", []LnsAny{lineNo}), token.Pos, false, nil), token.Pos.LineNo).Nodes_Node
        
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( token.Kind == Types_TokenKind__Kywd) &&
        Lns_GetEnv().SetStackVal( token.Txt == "fn") ).(bool)){
        exp = self.FP.analyzeExpSymbol(firstToken, token, TransUnit_ExpSymbolMode__Fn, nil, false, false)
        
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( token.Kind == Types_TokenKind__Kywd) &&
        Lns_GetEnv().SetStackVal( token.Txt == "unwrap") ).(bool)){
        exp = self.FP.analyzeExpUnwrap(token)
        
    } else if token.Kind == Types_TokenKind__Symb{
        exp = self.FP.analyzeExpSymbol(firstToken, token, TransUnit_ExpSymbolMode__Symbol, nil, false, canLeftExp)
        
        var symbolInfoList *LnsList
        symbolInfoList = exp.FP.GetSymbolInfo()
        if symbolInfoList.Len() == 1{
            var symbolInfo *Ast_SymbolInfo
            symbolInfo = symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if symbolInfo.FP.Get_kind() == Ast_SymbolKind__Typ{
                exp = &self.FP.analyzeRefTypeWithSymbol(Ast_AccessMode__Local, false, false, false, exp, false).Nodes_Node
                
                var workToken *Types_Token
                workToken = self.FP.getToken(nil)
                if workToken.Txt == "."{
                    exp = self.FP.analyzeExpSymbol(firstToken, self.FP.getToken(nil), TransUnit_ExpSymbolMode__Field, exp, false, canLeftExp)
                    
                } else { 
                    self.FP.Pushback()
                }
            }
        }
    } else if token.Kind == Types_TokenKind__Type{
        var symbolTypeInfo *Ast_SymbolInfo
        
        {
            _symbolTypeInfo := Ast_getSym2builtInTypeMap().Items[token.Txt]
            if _symbolTypeInfo == nil{
                self.FP.Error(Lns_getVM().String_format("unknown type -- %s", []LnsAny{token.Txt}))
            } else {
                symbolTypeInfo = _symbolTypeInfo.(*Ast_SymbolInfo)
            }
        }
        exp = &Nodes_ExpRefNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), &NewAst_AccessSymbolInfo(symbolTypeInfo, Ast_OverrideMut__None_Obj, false).Ast_SymbolInfo).Nodes_Node
        
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( token.Kind == Types_TokenKind__Kywd) &&
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == "true") ||
            Lns_GetEnv().SetStackVal( token.Txt == "false") ).(bool))) ).(bool)){
        exp = &Nodes_LiteralBoolNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeBool)}), token).Nodes_Node
        
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( token.Kind == Types_TokenKind__Kywd) &&
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == "nil") ||
            Lns_GetEnv().SetStackVal( token.Txt == "null") ).(bool))) ).(bool)){
        exp = &Nodes_LiteralNilNode_create(self.nodeManager, firstToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNil)})).Nodes_Node
        
    }
    if exp.FP.Get_kind() == Nodes_NodeKind_get_None(){
        self.FP.Error("illegal exp")
    }
    if skipOp2Flag{
        return exp
    }
    return self.FP.analyzeExpOp2(firstToken, exp, prevOpLevel)
}

// 27: decl @lune.@base.@TransUnit.TransUnit.analyzeReturn
func (self *TransUnit_TransUnit) analyzeReturn(token *Types_Token) *Nodes_ReturnNode {
    var expList LnsAny
    expList = nil
    var funcTypeInfo *Ast_TypeInfo
    funcTypeInfo = self.FP.getCurrentNamespaceTypeInfo()
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( funcTypeInfo == Ast_headTypeInfo) ||
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_kind() != Ast_TypeInfoKind__Func) &&
            Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_kind() != Ast_TypeInfoKind__Method) ).(bool))) ).(bool){
        self.FP.addErrMess(token.Pos, "'return' could not use here")
    }
    var nextToken *Types_Token
    nextToken = self.FP.getToken(nil)
    var retTypeList *LnsList
    retTypeList = funcTypeInfo.FP.Get_retTypeInfoList()
    if nextToken.Txt != ";"{
        self.FP.Pushback()
        expList = self.FP.analyzeExpList(false, false, false, nil, retTypeList, nil)
        
        self.FP.checkNextToken(";")
    }
    if funcTypeInfo.FP.GetTxt(nil, nil, nil) == "__init"{
        self.FP.addErrMess(token.Pos, "__init method can't return")
    }
    {
        _workList := expList
        if _workList != nil {
            workList := _workList.(*Nodes_ExpListNode)
            {
                _, _, _newExpNodeList := TransUnit_convExp73619(Lns_2DDD(self.FP.checkMatchType("return", token.Pos, retTypeList, workList, false, Lns_op_not(workList.FP.Get_followOn()), nil)))
                if _newExpNodeList != nil {
                    newExpNodeList := _newExpNodeList.(*Nodes_ExpListNode)
                    expList = newExpNodeList
                    
                }
            }
        } else {
            if retTypeList.Len() != 0{
                if retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() != Ast_TypeInfoKind__DDD{
                    self.FP.addErrMess(token.Pos, Lns_getVM().String_format("no return value -- need values(%d)", []LnsAny{retTypeList.Len()}))
                }
            }
        }
    }
    return Nodes_ReturnNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), retTypeList, expList)
}

// 74: decl @lune.@base.@TransUnit.TransUnit.analyzeStatement
func (self *TransUnit_TransUnit) analyzeStatement(termTxt LnsAny) LnsAny {
    self.commentCtrl.FP.Push()
    var token *Types_Token
    token = self.FP.GetTokenNoErr()
    var statement LnsAny
    statement = nil
    if token.Kind == Types_TokenKind__Sheb{
        statement = &Nodes_ShebangNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), token.Txt).Nodes_Node
        
    }
    if token == Parser_getEofToken(){
        return nil
    }
    if Lns_op_not(statement){
        statement = self.FP.analyzeDecl(Ast_AccessMode__Local, false, token, token)
        
    }
    if Lns_op_not(statement){
        if token.Txt == termTxt{
            if self.commentCtrl.FP.Get_commentList().Len() > 0{
                var commentToken *Types_Token
                commentToken = self.commentCtrl.FP.Get_commentList().GetAt(1).(Types_TokenDownCast).ToTypes_Token()
                statement = &Nodes_BlankLineNode_create(self.nodeManager, commentToken.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), 0).Nodes_Node
                
            }
            self.FP.Pushback()
            self.commentCtrl.FP.Pop()
            return statement
        } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == "pub") ||
            Lns_GetEnv().SetStackVal( token.Txt == "pro") ||
            Lns_GetEnv().SetStackVal( token.Txt == "pri") ||
            Lns_GetEnv().SetStackVal( token.Txt == "global") ||
            Lns_GetEnv().SetStackVal( token.Txt == "static") ).(bool){
            var accessMode LnsInt
            
            {
                _accessMode := Ast_txt2AccessMode(token.Txt)
                if _accessMode == nil{
                    accessMode = Ast_AccessMode__Pri
                    
                } else {
                    accessMode = _accessMode.(LnsInt)
                }
            }
            var staticFlag bool
            staticFlag = (token.Txt == "static")
            var nextToken *Types_Token
            nextToken = token
            if token.Txt != "static"{
                nextToken = self.FP.getToken(nil)
                
            }
            statement = self.FP.analyzeDecl(accessMode, staticFlag, token, nextToken)
            
            if Lns_op_not(statement){
                self.FP.addErrMess(nextToken.Pos, Lns_getVM().String_format("This token is illegal -- %s", []LnsAny{nextToken.Txt}))
            }
        } else if token.Txt == "{"{
            self.FP.Pushback()
            statement = &self.FP.analyzeBlock(Nodes_BlockKind__Block, TransUnit_TentativeMode__Simple, nil, nil).Nodes_Node
            
        } else if token.Txt == "super"{
            statement = self.FP.analyzeSuper(token)
            
        } else if token.Txt == "__envLock"{
            statement = self.FP.analyzeEnvLock(token)
            
        } else if token.Txt == "if"{
            statement = self.FP.analyzeIf(token)
            
        } else if token.Txt == "when"{
            statement = self.FP.analyzeWhen(token)
            
        } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == "switch") ||
            Lns_GetEnv().SetStackVal( token.Txt == "_switch") ).(bool){
            statement = &self.FP.analyzeSwitch(token).Nodes_Node
            
        } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == "match") ||
            Lns_GetEnv().SetStackVal( token.Txt == "_match") ).(bool){
            statement = &self.FP.analyzeMatch(token).Nodes_Node
            
        } else if token.Txt == "while"{
            statement = &self.FP.analyzeWhile(token).Nodes_Node
            
        } else if token.Txt == "repeat"{
            statement = &self.FP.analyzeRepeat(token).Nodes_Node
            
        } else if token.Txt == "for"{
            statement = &self.FP.analyzeFor(token).Nodes_Node
            
        } else if token.Txt == "apply"{
            statement = &self.FP.analyzeApply(token).Nodes_Node
            
        } else if token.Txt == "foreach"{
            statement = self.FP.analyzeForeach(token, false)
            
        } else if token.Txt == "forsort"{
            statement = self.FP.analyzeForeach(token, true)
            
        } else if token.Txt == "return"{
            statement = &self.FP.analyzeReturn(token).Nodes_Node
            
        } else if token.Txt == "break"{
            self.FP.checkNextToken(";")
            statement = &Nodes_BreakNode_create(self.nodeManager, token.Pos, self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})).Nodes_Node
            
            if self.loopScopeQueue.Len() == 0{
                self.FP.addErrMess(token.Pos, "no loop syntax.")
            }
        } else if token.Txt == "unwrap"{
            statement = self.FP.analyzeUnwrap(token)
            
        } else if token.Txt == "sync"{
            statement = self.FP.analyzeDeclVar(Nodes_DeclVarMode__Sync, Ast_AccessMode__Local, token)
            
        } else if token.Txt == "__scope"{
            statement = &self.FP.analyzeScope(token).Nodes_Node
            
        } else if token.Txt == "import"{
            statement = self.FP.analyzeImport(token)
            
        } else if token.Txt == "subfile"{
            statement = &self.FP.analyzeSubfile(token).Nodes_Node
            
        } else if token.Txt == "_lune_control"{
            {
                __exp := self.FP.analyzeLuneControl(token)
                if __exp != nil {
                    _exp := __exp.(*Nodes_LuneControlNode)
                    statement = &_exp.Nodes_Node
                    
                } else {
                    statement = self.FP.createNoneNode(token.Pos)
                    
                }
            }
        } else if token.Txt == "provide"{
            statement = &self.FP.analyzeProvide(token).Nodes_Node
            
        } else if token.Txt == ";"{
            statement = self.FP.createNoneNode(token.Pos)
            
        } else if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( token.Txt == ",,") ||
            Lns_GetEnv().SetStackVal( token.Txt == ",,,") ||
            Lns_GetEnv().SetStackVal( token.Txt == ",,,,") ).(bool){
            self.FP.Error(Lns_getVM().String_format("illegal macro op -- %s", []LnsAny{token.Txt}))
        } else { 
            self.FP.Pushback()
            self.accessSymbolSetQueue.FP.Push()
            var exp *Nodes_Node
            exp = self.FP.analyzeExp(true, false, true, nil, nil)
            var nextToken *Types_Token
            nextToken = self.FP.getToken(nil)
            if nextToken.Txt == ","{
                var expList *Nodes_ExpListNode
                expList = self.FP.analyzeExpList(true, true, true, exp, nil, nil)
                exp = self.FP.analyzeExpOp2(token, &expList.Nodes_Node, nil)
                
                nextToken = self.FP.getToken(nil)
                
            }
            self.FP.checkToken(nextToken, ";")
            {
                _setNode := Nodes_ExpSetValNodeDownCastF(exp.FP)
                if _setNode != nil {
                    setNode := _setNode.(*Nodes_ExpSetValNode)
                    var set *LnsSet
                    set = NewLnsSet([]LnsAny{})
                    for _, _symbol := range( setNode.FP.Get_LeftSymList().Items ) {
                        symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                        set.Add(Ast_SymbolInfo2Stem(symbol.FP.GetOrg()))
                    }
                    self.accessSymbolSetQueue.FP.Pop(set)
                } else {
                    self.accessSymbolSetQueue.FP.Pop(nil)
                }
            }
            statement = &Nodes_StmtExpNode_create(self.nodeManager, exp.FP.Get_pos(), self.macroCtrl.FP.IsInAnalyzeArgMode(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), exp).Nodes_Node
            
        }
    }
    if statement != nil{
        statement_11377 := statement.(*Nodes_Node)
        if Lns_op_not(statement_11377.FP.CanBeStatement()){
            self.FP.addErrMess(statement_11377.FP.Get_pos(), Lns_getVM().String_format("This node can't be statement. -- %s", []LnsAny{Nodes_getNodeKindName(statement_11377.FP.Get_kind())}))
        }
    }
    self.commentCtrl.FP.Pop()
    return statement
}


// declaration Class -- ImportParam
type TransUnit_ImportParamMtd interface {
    GetTypeInfo(arg1 LnsInt)(LnsAny, LnsAny)
}
type TransUnit_ImportParam struct {
    Pos *Types_Position
    TransUnit TransUnit_TransUnitIF
    ProcessInfo *Ast_ProcessInfo
    TypeId2Scope *LnsMap
    TypeId2TypeInfo *LnsMap
    ImportedAliasMap *LnsMap
    LazyModuleSet *LnsSet
    MetaInfo *Lns_luaValue
    Scope *Ast_Scope
    ModuleTypeInfo *Ast_TypeInfo
    ScopeAccess LnsInt
    typeId2AtomMap *LnsMap
    FP TransUnit_ImportParamMtd
}
func TransUnit_ImportParam2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_ImportParam).FP
}
type TransUnit_ImportParamDownCast interface {
    ToTransUnit_ImportParam() *TransUnit_ImportParam
}
func TransUnit_ImportParamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_ImportParamDownCast)
    if ok { return work.ToTransUnit_ImportParam() }
    return nil
}
func (obj *TransUnit_ImportParam) ToTransUnit_ImportParam() *TransUnit_ImportParam {
    return obj
}
func NewTransUnit_ImportParam(arg1 *Types_Position, arg2 TransUnit_TransUnitIF, arg3 *Ast_ProcessInfo, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsSet, arg8 *Lns_luaValue, arg9 *Ast_Scope, arg10 *Ast_TypeInfo, arg11 LnsInt, arg12 *LnsMap) *TransUnit_ImportParam {
    obj := &TransUnit_ImportParam{}
    obj.FP = obj
    obj.InitTransUnit_ImportParam(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
    return obj
}
func (self *TransUnit_ImportParam) InitTransUnit_ImportParam(arg1 *Types_Position, arg2 TransUnit_TransUnitIF, arg3 *Ast_ProcessInfo, arg4 *LnsMap, arg5 *LnsMap, arg6 *LnsMap, arg7 *LnsSet, arg8 *Lns_luaValue, arg9 *Ast_Scope, arg10 *Ast_TypeInfo, arg11 LnsInt, arg12 *LnsMap) {
    self.Pos = arg1
    self.TransUnit = arg2
    self.ProcessInfo = arg3
    self.TypeId2Scope = arg4
    self.TypeId2TypeInfo = arg5
    self.ImportedAliasMap = arg6
    self.LazyModuleSet = arg7
    self.MetaInfo = arg8
    self.Scope = arg9
    self.ModuleTypeInfo = arg10
    self.ScopeAccess = arg11
    self.typeId2AtomMap = arg12
}
// 1360: decl @lune.@base.@TransUnit.ImportParam.getTypeInfo
func (self *TransUnit_ImportParam) GetTypeInfo(typeId LnsInt)(LnsAny, LnsAny) {
    {
        _typeInfo := self.TypeId2TypeInfo.Items[typeId]
        if _typeInfo != nil {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo, nil
        }
    }
    {
        _atom := self.typeId2AtomMap.Items[typeId]
        if _atom != nil {
            atom := _atom.(*TransUnit__TypeInfo)
            var typeInfo LnsAny
            var mess LnsAny
            typeInfo,mess = atom.FP.CreateTypeInfoCache(self)
            if typeInfo != nil{
                typeInfo_6265 := typeInfo.(*Ast_TypeInfo)
                self.TypeId2TypeInfo.Set(typeId,typeInfo_6265)
            }
            return typeInfo, mess
        }
    }
    return nil, nil
}


// declaration Class -- _TypeInfo
type TransUnit__TypeInfoMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfo struct {
    Skind LnsInt
    TypeId LnsInt
    FP TransUnit__TypeInfoMtd
}
func TransUnit__TypeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfo).FP
}
type TransUnit__TypeInfoDownCast interface {
    ToTransUnit__TypeInfo() *TransUnit__TypeInfo
}
func TransUnit__TypeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoDownCast)
    if ok { return work.ToTransUnit__TypeInfo() }
    return nil
}
func (obj *TransUnit__TypeInfo) ToTransUnit__TypeInfo() *TransUnit__TypeInfo {
    return obj
}
func (self *TransUnit__TypeInfo) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["skind"] = Lns_ToCollection( self.Skind )
    obj.Items["typeId"] = Lns_ToCollection( self.TypeId )
    return obj
}
func (self *TransUnit__TypeInfo) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfo_FromMapMain( newObj *TransUnit__TypeInfo, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["skind"], false, nil); !ok {
       return false,nil,"skind:" + mess.(string)
    } else {
       newObj.Skind = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["typeId"], false, nil); !ok {
       return false,nil,"typeId:" + mess.(string)
    } else {
       newObj.TypeId = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 1333: DeclConstr
func (self *TransUnit__TypeInfo) InitTransUnit__TypeInfo() {
    self.TypeId = Ast_rootTypeId
    
    self.Skind = Ast_SerializeKind__Normal
    
}


// 1348: decl @lune.@base.@TransUnit._TypeInfo.createTypeInfoCache
func (self *TransUnit__TypeInfo) CreateTypeInfoCache(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    {
        _typeInfo := param.TypeId2TypeInfo.Items[self.TypeId]
        if _typeInfo != nil {
            typeInfo := _typeInfo.(*Ast_TypeInfo)
            return typeInfo, nil
        }
    }
    var typeInfo LnsAny
    var mess LnsAny
    typeInfo,mess = self.FP.CreateTypeInfo(param)
    if typeInfo != nil{
        typeInfo_6250 := typeInfo.(*Ast_TypeInfo)
        param.TypeId2TypeInfo.Set(self.TypeId,typeInfo_6250)
    }
    return typeInfo, mess
}


// declaration Class -- _TypeInfoNilable
type TransUnit__TypeInfoNilableMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoNilable struct {
    TransUnit__TypeInfo
    OrgTypeId LnsInt
    FP TransUnit__TypeInfoNilableMtd
}
func TransUnit__TypeInfoNilable2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoNilable).FP
}
type TransUnit__TypeInfoNilableDownCast interface {
    ToTransUnit__TypeInfoNilable() *TransUnit__TypeInfoNilable
}
func TransUnit__TypeInfoNilableDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoNilableDownCast)
    if ok { return work.ToTransUnit__TypeInfoNilable() }
    return nil
}
func (obj *TransUnit__TypeInfoNilable) ToTransUnit__TypeInfoNilable() *TransUnit__TypeInfoNilable {
    return obj
}
func NewTransUnit__TypeInfoNilable(arg1 LnsInt) *TransUnit__TypeInfoNilable {
    obj := &TransUnit__TypeInfoNilable{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoNilable(arg1)
    return obj
}
func (self *TransUnit__TypeInfoNilable) InitTransUnit__TypeInfoNilable(arg1 LnsInt) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.OrgTypeId = arg1
}
func (self *TransUnit__TypeInfoNilable) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["orgTypeId"] = Lns_ToCollection( self.OrgTypeId )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoNilable) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoNilable__fromMap_1855_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoNilable_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoNilable__fromStem_1858_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoNilable_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoNilable_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoNilable_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoNilable_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoNilable{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoNilable_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoNilable_FromMapMain( newObj *TransUnit__TypeInfoNilable, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["orgTypeId"], false, nil); !ok {
       return false,nil,"orgTypeId:" + mess.(string)
    } else {
       newObj.OrgTypeId = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 1377: decl @lune.@base.@TransUnit._TypeInfoNilable.createTypeInfo
func (self *TransUnit__TypeInfoNilable) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    var orgTypeInfo *Ast_TypeInfo
    orgTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.OrgTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = orgTypeInfo.FP.Get_nilableTypeInfo()
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoAlias
type TransUnit__TypeInfoAliasMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoAlias struct {
    TransUnit__TypeInfo
    ParentId LnsInt
    rawTxt string
    srcTypeId LnsInt
    FP TransUnit__TypeInfoAliasMtd
}
func TransUnit__TypeInfoAlias2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoAlias).FP
}
type TransUnit__TypeInfoAliasDownCast interface {
    ToTransUnit__TypeInfoAlias() *TransUnit__TypeInfoAlias
}
func TransUnit__TypeInfoAliasDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoAliasDownCast)
    if ok { return work.ToTransUnit__TypeInfoAlias() }
    return nil
}
func (obj *TransUnit__TypeInfoAlias) ToTransUnit__TypeInfoAlias() *TransUnit__TypeInfoAlias {
    return obj
}
func NewTransUnit__TypeInfoAlias(arg1 LnsInt, arg2 string, arg3 LnsInt) *TransUnit__TypeInfoAlias {
    obj := &TransUnit__TypeInfoAlias{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoAlias(arg1, arg2, arg3)
    return obj
}
func (self *TransUnit__TypeInfoAlias) InitTransUnit__TypeInfoAlias(arg1 LnsInt, arg2 string, arg3 LnsInt) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.ParentId = arg1
    self.rawTxt = arg2
    self.srcTypeId = arg3
}
func (self *TransUnit__TypeInfoAlias) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["rawTxt"] = Lns_ToCollection( self.rawTxt )
    obj.Items["srcTypeId"] = Lns_ToCollection( self.srcTypeId )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoAlias) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoAlias__fromMap_1879_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoAlias_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoAlias__fromStem_1882_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoAlias_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoAlias_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoAlias_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoAlias_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoAlias{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoAlias_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoAlias_FromMapMain( newObj *TransUnit__TypeInfoAlias, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["rawTxt"], false, nil); !ok {
       return false,nil,"rawTxt:" + mess.(string)
    } else {
       newObj.rawTxt = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["srcTypeId"], false, nil); !ok {
       return false,nil,"srcTypeId:" + mess.(string)
    } else {
       newObj.srcTypeId = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 1393: decl @lune.@base.@TransUnit._TypeInfoAlias.createTypeInfo
func (self *TransUnit__TypeInfoAlias) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@TransUnit._TypeInfoAlias.createTypeInfo"
    var srcTypeInfo *Ast_TypeInfo
    srcTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.srcTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_AliasTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateAlias(param.ProcessInfo, self.rawTxt, true, Ast_AccessMode__Pub, param.ModuleTypeInfo, srcTypeInfo)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    var parentScope *Ast_Scope
    
    {
        _parentScope := param.TypeId2Scope.Items[self.ParentId]
        if _parentScope == nil{
            return nil, Lns_getVM().String_format("%s: not found parentScope %s %s", []LnsAny{__func__, self.ParentId, self.rawTxt})
        } else {
            parentScope = _parentScope.(*Ast_Scope)
        }
    }
    parentScope.FP.AddAliasForType(param.ProcessInfo, self.rawTxt, nil, &newTypeInfo.Ast_TypeInfo)
    param.ImportedAliasMap.Set(srcTypeInfo,newTypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoDDD
type TransUnit__TypeInfoDDDMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoDDD struct {
    TransUnit__TypeInfo
    ParentId LnsInt
    ItemTypeId LnsInt
    ExtTypeFlag bool
    FP TransUnit__TypeInfoDDDMtd
}
func TransUnit__TypeInfoDDD2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoDDD).FP
}
type TransUnit__TypeInfoDDDDownCast interface {
    ToTransUnit__TypeInfoDDD() *TransUnit__TypeInfoDDD
}
func TransUnit__TypeInfoDDDDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoDDDDownCast)
    if ok { return work.ToTransUnit__TypeInfoDDD() }
    return nil
}
func (obj *TransUnit__TypeInfoDDD) ToTransUnit__TypeInfoDDD() *TransUnit__TypeInfoDDD {
    return obj
}
func NewTransUnit__TypeInfoDDD(arg1 LnsInt, arg2 LnsInt, arg3 bool) *TransUnit__TypeInfoDDD {
    obj := &TransUnit__TypeInfoDDD{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoDDD(arg1, arg2, arg3)
    return obj
}
func (self *TransUnit__TypeInfoDDD) InitTransUnit__TypeInfoDDD(arg1 LnsInt, arg2 LnsInt, arg3 bool) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.ParentId = arg1
    self.ItemTypeId = arg2
    self.ExtTypeFlag = arg3
}
func (self *TransUnit__TypeInfoDDD) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["itemTypeId"] = Lns_ToCollection( self.ItemTypeId )
    obj.Items["extTypeFlag"] = Lns_ToCollection( self.ExtTypeFlag )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoDDD) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoDDD__fromMap_1903_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoDDD_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoDDD__fromStem_1906_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoDDD_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoDDD_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoDDD_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoDDD_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoDDD{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoDDD_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoDDD_FromMapMain( newObj *TransUnit__TypeInfoDDD, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["itemTypeId"], false, nil); !ok {
       return false,nil,"itemTypeId:" + mess.(string)
    } else {
       newObj.ItemTypeId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["extTypeFlag"], false, nil); !ok {
       return false,nil,"extTypeFlag:" + mess.(string)
    } else {
       newObj.ExtTypeFlag = conv.(bool)
    }
    return true, newObj, nil
}
// 1417: decl @lune.@base.@TransUnit._TypeInfoDDD.createTypeInfo
func (self *TransUnit__TypeInfoDDD) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    var itemTypeInfo *Ast_TypeInfo
    itemTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.ItemTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_DDDTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateDDD(itemTypeInfo, true, self.ExtTypeFlag)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoAlternate
type TransUnit__TypeInfoAlternateMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoAlternate struct {
    TransUnit__TypeInfo
    ParentId LnsInt
    Txt string
    AccessMode LnsInt
    BaseId LnsInt
    IfList *LnsList
    BelongClassFlag bool
    AltIndex LnsInt
    FP TransUnit__TypeInfoAlternateMtd
}
func TransUnit__TypeInfoAlternate2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoAlternate).FP
}
type TransUnit__TypeInfoAlternateDownCast interface {
    ToTransUnit__TypeInfoAlternate() *TransUnit__TypeInfoAlternate
}
func TransUnit__TypeInfoAlternateDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoAlternateDownCast)
    if ok { return work.ToTransUnit__TypeInfoAlternate() }
    return nil
}
func (obj *TransUnit__TypeInfoAlternate) ToTransUnit__TypeInfoAlternate() *TransUnit__TypeInfoAlternate {
    return obj
}
func NewTransUnit__TypeInfoAlternate(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 LnsInt, arg5 *LnsList, arg6 bool, arg7 LnsInt) *TransUnit__TypeInfoAlternate {
    obj := &TransUnit__TypeInfoAlternate{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoAlternate(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *TransUnit__TypeInfoAlternate) InitTransUnit__TypeInfoAlternate(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 LnsInt, arg5 *LnsList, arg6 bool, arg7 LnsInt) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.ParentId = arg1
    self.Txt = arg2
    self.AccessMode = arg3
    self.BaseId = arg4
    self.IfList = arg5
    self.BelongClassFlag = arg6
    self.AltIndex = arg7
}
func (self *TransUnit__TypeInfoAlternate) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["baseId"] = Lns_ToCollection( self.BaseId )
    obj.Items["ifList"] = Lns_ToCollection( self.IfList )
    obj.Items["belongClassFlag"] = Lns_ToCollection( self.BelongClassFlag )
    obj.Items["altIndex"] = Lns_ToCollection( self.AltIndex )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoAlternate) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoAlternate__fromMap_1937_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoAlternate_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoAlternate__fromStem_1940_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoAlternate_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoAlternate_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoAlternate_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoAlternate_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoAlternate{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoAlternate_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoAlternate_FromMapMain( newObj *TransUnit__TypeInfoAlternate, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["baseId"], false, nil); !ok {
       return false,nil,"baseId:" + mess.(string)
    } else {
       newObj.BaseId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["ifList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"ifList:" + mess.(string)
    } else {
       newObj.IfList = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["belongClassFlag"], false, nil); !ok {
       return false,nil,"belongClassFlag:" + mess.(string)
    } else {
       newObj.BelongClassFlag = conv.(bool)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["altIndex"], false, nil); !ok {
       return false,nil,"altIndex:" + mess.(string)
    } else {
       newObj.AltIndex = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 1438: decl @lune.@base.@TransUnit._TypeInfoAlternate.createTypeInfo
func (self *TransUnit__TypeInfoAlternate) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    var baseInfo *Ast_TypeInfo
    baseInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.BaseId))).(*Ast_TypeInfo)
    var interfaceList *LnsList
    interfaceList = NewLnsList([]LnsAny{})
    for _, _ifTypeId := range( self.IfList.Items ) {
        ifTypeId := _ifTypeId.(LnsInt)
        interfaceList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfo(ifTypeId))).(*Ast_TypeInfo)))
    }
    var newTypeInfo *Ast_AlternateTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateAlternate(self.BelongClassFlag, self.AltIndex, self.Txt, self.AccessMode, param.ModuleTypeInfo, baseInfo, interfaceList)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoGeneric
type TransUnit__TypeInfoGenericMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoGeneric struct {
    TransUnit__TypeInfo
    GenSrcTypeId LnsInt
    GenTypeList *LnsList
    FP TransUnit__TypeInfoGenericMtd
}
func TransUnit__TypeInfoGeneric2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoGeneric).FP
}
type TransUnit__TypeInfoGenericDownCast interface {
    ToTransUnit__TypeInfoGeneric() *TransUnit__TypeInfoGeneric
}
func TransUnit__TypeInfoGenericDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoGenericDownCast)
    if ok { return work.ToTransUnit__TypeInfoGeneric() }
    return nil
}
func (obj *TransUnit__TypeInfoGeneric) ToTransUnit__TypeInfoGeneric() *TransUnit__TypeInfoGeneric {
    return obj
}
func NewTransUnit__TypeInfoGeneric(arg1 LnsInt, arg2 *LnsList) *TransUnit__TypeInfoGeneric {
    obj := &TransUnit__TypeInfoGeneric{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoGeneric(arg1, arg2)
    return obj
}
func (self *TransUnit__TypeInfoGeneric) InitTransUnit__TypeInfoGeneric(arg1 LnsInt, arg2 *LnsList) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.GenSrcTypeId = arg1
    self.GenTypeList = arg2
}
func (self *TransUnit__TypeInfoGeneric) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["genSrcTypeId"] = Lns_ToCollection( self.GenSrcTypeId )
    obj.Items["genTypeList"] = Lns_ToCollection( self.GenTypeList )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoGeneric) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoGeneric__fromMap_1971_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoGeneric_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoGeneric__fromStem_1974_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoGeneric_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoGeneric_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoGeneric_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoGeneric_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoGeneric{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoGeneric_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoGeneric_FromMapMain( newObj *TransUnit__TypeInfoGeneric, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["genSrcTypeId"], false, nil); !ok {
       return false,nil,"genSrcTypeId:" + mess.(string)
    } else {
       newObj.GenSrcTypeId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["genTypeList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"genTypeList:" + mess.(string)
    } else {
       newObj.GenTypeList = conv.(*LnsList)
    }
    return true, newObj, nil
}
// 1458: decl @lune.@base.@TransUnit._TypeInfoGeneric.createTypeInfo
func (self *TransUnit__TypeInfoGeneric) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    var genSrcTypeInfo *Ast_TypeInfo
    genSrcTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.GenSrcTypeId))).(*Ast_TypeInfo)
    var genTypeList *LnsList
    genTypeList = NewLnsList([]LnsAny{})
    for _, _typeId := range( self.GenTypeList.Items ) {
        typeId := _typeId.(LnsInt)
        genTypeList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfo(typeId))).(*Ast_TypeInfo)))
    }
    var newTypeInfo *Ast_GenericTypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateGeneric(genSrcTypeInfo, genTypeList, param.ModuleTypeInfo)
    param.TypeId2TypeInfo.Set(self.TypeId,&newTypeInfo.Ast_TypeInfo)
    param.TypeId2Scope.Set(self.TypeId,Ast_getScope(&newTypeInfo.Ast_TypeInfo))
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoBox
type TransUnit__TypeInfoBoxMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoBox struct {
    TransUnit__TypeInfo
    AccessMode LnsInt
    BoxingType LnsInt
    FP TransUnit__TypeInfoBoxMtd
}
func TransUnit__TypeInfoBox2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoBox).FP
}
type TransUnit__TypeInfoBoxDownCast interface {
    ToTransUnit__TypeInfoBox() *TransUnit__TypeInfoBox
}
func TransUnit__TypeInfoBoxDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoBoxDownCast)
    if ok { return work.ToTransUnit__TypeInfoBox() }
    return nil
}
func (obj *TransUnit__TypeInfoBox) ToTransUnit__TypeInfoBox() *TransUnit__TypeInfoBox {
    return obj
}
func NewTransUnit__TypeInfoBox(arg1 LnsInt, arg2 LnsInt) *TransUnit__TypeInfoBox {
    obj := &TransUnit__TypeInfoBox{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoBox(arg1, arg2)
    return obj
}
func (self *TransUnit__TypeInfoBox) InitTransUnit__TypeInfoBox(arg1 LnsInt, arg2 LnsInt) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.AccessMode = arg1
    self.BoxingType = arg2
}
func (self *TransUnit__TypeInfoBox) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["boxingType"] = Lns_ToCollection( self.BoxingType )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoBox) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoBox__fromMap_1995_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoBox_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoBox__fromStem_1998_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoBox_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoBox_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoBox_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoBox_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoBox{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoBox_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoBox_FromMapMain( newObj *TransUnit__TypeInfoBox, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["boxingType"], false, nil); !ok {
       return false,nil,"boxingType:" + mess.(string)
    } else {
       newObj.BoxingType = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 1477: decl @lune.@base.@TransUnit._TypeInfoBox.createTypeInfo
func (self *TransUnit__TypeInfoBox) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    var boxingType *Ast_TypeInfo
    boxingType = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.BoxingType))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = param.ProcessInfo.FP.CreateBox(self.AccessMode, boxingType)
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoExt
type TransUnit__TypeInfoExtMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoExt struct {
    TransUnit__TypeInfo
    ExtedTypeId LnsInt
    FP TransUnit__TypeInfoExtMtd
}
func TransUnit__TypeInfoExt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoExt).FP
}
type TransUnit__TypeInfoExtDownCast interface {
    ToTransUnit__TypeInfoExt() *TransUnit__TypeInfoExt
}
func TransUnit__TypeInfoExtDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoExtDownCast)
    if ok { return work.ToTransUnit__TypeInfoExt() }
    return nil
}
func (obj *TransUnit__TypeInfoExt) ToTransUnit__TypeInfoExt() *TransUnit__TypeInfoExt {
    return obj
}
func NewTransUnit__TypeInfoExt(arg1 LnsInt) *TransUnit__TypeInfoExt {
    obj := &TransUnit__TypeInfoExt{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoExt(arg1)
    return obj
}
func (self *TransUnit__TypeInfoExt) InitTransUnit__TypeInfoExt(arg1 LnsInt) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.ExtedTypeId = arg1
}
func (self *TransUnit__TypeInfoExt) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["extedTypeId"] = Lns_ToCollection( self.ExtedTypeId )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoExt) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoExt__fromMap_2019_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoExt_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoExt__fromStem_2022_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoExt_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoExt_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoExt_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoExt_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoExt{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoExt_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoExt_FromMapMain( newObj *TransUnit__TypeInfoExt, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["extedTypeId"], false, nil); !ok {
       return false,nil,"extedTypeId:" + mess.(string)
    } else {
       newObj.ExtedTypeId = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 1489: decl @lune.@base.@TransUnit._TypeInfoExt.createTypeInfo
func (self *TransUnit__TypeInfoExt) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    var extedType *Ast_TypeInfo
    extedType = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.ExtedTypeId))).(*Ast_TypeInfo)
    var newTypeInfo *Ast_TypeInfo
    switch _exp6389 := param.ProcessInfo.FP.CreateLuaval(extedType, true).(type) {
    case *Ast_LuavalResult__OK:
    extType := _exp6389.Val1
        newTypeInfo = extType
        
    case *Ast_LuavalResult__Err:
    mess := _exp6389.Val1
        Util_err(mess)
    }
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoModifier
type TransUnit__TypeInfoModifierMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoModifier struct {
    TransUnit__TypeInfo
    SrcTypeId LnsInt
    MutMode LnsInt
    FP TransUnit__TypeInfoModifierMtd
}
func TransUnit__TypeInfoModifier2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoModifier).FP
}
type TransUnit__TypeInfoModifierDownCast interface {
    ToTransUnit__TypeInfoModifier() *TransUnit__TypeInfoModifier
}
func TransUnit__TypeInfoModifierDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoModifierDownCast)
    if ok { return work.ToTransUnit__TypeInfoModifier() }
    return nil
}
func (obj *TransUnit__TypeInfoModifier) ToTransUnit__TypeInfoModifier() *TransUnit__TypeInfoModifier {
    return obj
}
func NewTransUnit__TypeInfoModifier(arg1 LnsInt, arg2 LnsInt) *TransUnit__TypeInfoModifier {
    obj := &TransUnit__TypeInfoModifier{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoModifier(arg1, arg2)
    return obj
}
func (self *TransUnit__TypeInfoModifier) InitTransUnit__TypeInfoModifier(arg1 LnsInt, arg2 LnsInt) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.SrcTypeId = arg1
    self.MutMode = arg2
}
func (self *TransUnit__TypeInfoModifier) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["srcTypeId"] = Lns_ToCollection( self.SrcTypeId )
    obj.Items["mutMode"] = Lns_ToCollection( self.MutMode )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoModifier) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoModifier__fromMap_2043_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoModifier_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoModifier__fromStem_2046_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoModifier_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoModifier_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoModifier_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoModifier_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoModifier{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoModifier_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoModifier_FromMapMain( newObj *TransUnit__TypeInfoModifier, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["srcTypeId"], false, nil); !ok {
       return false,nil,"srcTypeId:" + mess.(string)
    } else {
       newObj.SrcTypeId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["mutMode"], false, nil); !ok {
       return false,nil,"mutMode:" + mess.(string)
    } else {
       newObj.MutMode = conv.(LnsInt)
    }
    return true, newObj, nil
}
// 1511: decl @lune.@base.@TransUnit._TypeInfoModifier.createTypeInfo
func (self *TransUnit__TypeInfoModifier) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    var srcTypeInfo *Ast_TypeInfo
    
    {
        _srcTypeInfo := TransUnit_convExp6455(Lns_2DDD(param.FP.GetTypeInfo(self.SrcTypeId)))
        if _srcTypeInfo == nil{
            return nil, Lns_getVM().String_format("not found srcType -- %d", []LnsAny{self.SrcTypeId})
        } else {
            srcTypeInfo = _srcTypeInfo.(*Ast_TypeInfo)
        }
    }
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = param.TransUnit.CreateModifier(srcTypeInfo, self.MutMode)
    param.TypeId2TypeInfo.Set(self.TypeId,newTypeInfo)
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoModule
type TransUnit__TypeInfoModuleMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoModule struct {
    TransUnit__TypeInfo
    ParentId LnsInt
    Txt string
    FP TransUnit__TypeInfoModuleMtd
}
func TransUnit__TypeInfoModule2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoModule).FP
}
type TransUnit__TypeInfoModuleDownCast interface {
    ToTransUnit__TypeInfoModule() *TransUnit__TypeInfoModule
}
func TransUnit__TypeInfoModuleDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoModuleDownCast)
    if ok { return work.ToTransUnit__TypeInfoModule() }
    return nil
}
func (obj *TransUnit__TypeInfoModule) ToTransUnit__TypeInfoModule() *TransUnit__TypeInfoModule {
    return obj
}
func NewTransUnit__TypeInfoModule(arg1 LnsInt, arg2 string) *TransUnit__TypeInfoModule {
    obj := &TransUnit__TypeInfoModule{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoModule(arg1, arg2)
    return obj
}
func (self *TransUnit__TypeInfoModule) InitTransUnit__TypeInfoModule(arg1 LnsInt, arg2 string) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.ParentId = arg1
    self.Txt = arg2
}
func (self *TransUnit__TypeInfoModule) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoModule) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoModule__fromMap_2070_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoModule_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoModule__fromStem_2073_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoModule_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoModule_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoModule_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoModule_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoModule{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoModule_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoModule_FromMapMain( newObj *TransUnit__TypeInfoModule, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    return true, newObj, nil
}
// 1527: decl @lune.@base.@TransUnit._TypeInfoModule.createTypeInfo
func (self *TransUnit__TypeInfoModule) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@TransUnit._TypeInfoModule.createTypeInfo"
    var parentInfo *Ast_TypeInfo
    parentInfo = Ast_headTypeInfo
    if self.ParentId != Ast_rootTypeId{
        var workTypeInfo *Ast_TypeInfo
        
        {
            _workTypeInfo := TransUnit_convExp6553(Lns_2DDD(param.FP.GetTypeInfo(self.ParentId)))
            if _workTypeInfo == nil{
                Util_err(Lns_getVM().String_format("not found parentInfo %s %s", []LnsAny{self.ParentId, self.Txt}))
            } else {
                workTypeInfo = _workTypeInfo.(*Ast_TypeInfo)
            }
        }
        parentInfo = workTypeInfo
        
    }
    var parentScope *Ast_Scope
    
    {
        _parentScope := param.TypeId2Scope.Items[self.ParentId]
        if _parentScope == nil{
            return nil, Lns_getVM().String_format("%s: not found parentScope %s %s", []LnsAny{__func__, self.ParentId, self.Txt})
        } else {
            parentScope = _parentScope.(*Ast_Scope)
        }
    }
    var newTypeInfo LnsAny
    newTypeInfo = parentScope.FP.GetTypeInfoChild(self.Txt)
    {
        __exp := newTypeInfo
        if __exp != nil {
            _exp := __exp.(*Ast_TypeInfo)
            param.TypeId2Scope.Set(self.TypeId,Ast_getScope(_exp))
            if Lns_op_not(_exp.FP.Get_scope()){
                return nil, Lns_getVM().String_format("not found scope %s %s %s %s %s", []LnsAny{parentScope, self.ParentId, self.TypeId, self.Txt, _exp.FP.GetTxt(nil, nil, nil)})
            }
            param.TypeId2TypeInfo.Set(self.TypeId,_exp)
        } else {
            var scope *Ast_Scope
            scope = NewAst_Scope(param.ProcessInfo, parentScope, true, nil, nil)
            var mutable bool
            mutable = false
            if self.TypeId == param.MetaInfo.GetAt( "__moduleTypeId" ).(LnsInt){
                mutable = param.MetaInfo.GetAt( "__moduleMutable" ).(bool)
                
            }
            var workTypeInfo *Ast_TypeInfo
            workTypeInfo = param.ProcessInfo.FP.CreateModule(scope, parentInfo, true, self.Txt, mutable)
            newTypeInfo = workTypeInfo
            
            param.TypeId2Scope.Set(self.TypeId,scope)
            param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
            parentScope.FP.AddClass(param.ProcessInfo, self.Txt, nil, workTypeInfo)
            Log_log(Log_Level__Info, __func__, 1569, Log_CreateMessage(func() string {
                return Lns_getVM().String_format("new module -- %s, %s, %d, %d", []LnsAny{self.Txt, workTypeInfo.FP.GetFullName(Ast_defaultTypeNameCtrl, parentScope.FP, false), workTypeInfo.FP.Get_typeId(), parentScope.FP.Get_scopeId()})
            }))
            
        }
    }
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoNormal
type TransUnit__TypeInfoNormalMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoNormal struct {
    TransUnit__TypeInfo
    ParentId LnsInt
    AbstractFlag bool
    BaseId LnsInt
    Txt string
    StaticFlag bool
    AccessMode LnsInt
    Kind LnsInt
    MutMode LnsInt
    IfList *LnsList
    ItemTypeId *LnsList
    ArgTypeId *LnsList
    RetTypeId *LnsList
    Children *LnsList
    ModuleLang LnsAny
    RequirePath LnsAny
    FP TransUnit__TypeInfoNormalMtd
}
func TransUnit__TypeInfoNormal2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoNormal).FP
}
type TransUnit__TypeInfoNormalDownCast interface {
    ToTransUnit__TypeInfoNormal() *TransUnit__TypeInfoNormal
}
func TransUnit__TypeInfoNormalDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoNormalDownCast)
    if ok { return work.ToTransUnit__TypeInfoNormal() }
    return nil
}
func (obj *TransUnit__TypeInfoNormal) ToTransUnit__TypeInfoNormal() *TransUnit__TypeInfoNormal {
    return obj
}
func NewTransUnit__TypeInfoNormal(arg1 LnsInt, arg2 bool, arg3 LnsInt, arg4 string, arg5 bool, arg6 LnsInt, arg7 LnsInt, arg8 LnsInt, arg9 *LnsList, arg10 *LnsList, arg11 *LnsList, arg12 *LnsList, arg13 *LnsList, arg14 LnsAny, arg15 LnsAny) *TransUnit__TypeInfoNormal {
    obj := &TransUnit__TypeInfoNormal{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoNormal(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15)
    return obj
}
func (self *TransUnit__TypeInfoNormal) InitTransUnit__TypeInfoNormal(arg1 LnsInt, arg2 bool, arg3 LnsInt, arg4 string, arg5 bool, arg6 LnsInt, arg7 LnsInt, arg8 LnsInt, arg9 *LnsList, arg10 *LnsList, arg11 *LnsList, arg12 *LnsList, arg13 *LnsList, arg14 LnsAny, arg15 LnsAny) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.ParentId = arg1
    self.AbstractFlag = arg2
    self.BaseId = arg3
    self.Txt = arg4
    self.StaticFlag = arg5
    self.AccessMode = arg6
    self.Kind = arg7
    self.MutMode = arg8
    self.IfList = arg9
    self.ItemTypeId = arg10
    self.ArgTypeId = arg11
    self.RetTypeId = arg12
    self.Children = arg13
    self.ModuleLang = arg14
    self.RequirePath = arg15
}
func (self *TransUnit__TypeInfoNormal) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["abstractFlag"] = Lns_ToCollection( self.AbstractFlag )
    obj.Items["baseId"] = Lns_ToCollection( self.BaseId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    obj.Items["staticFlag"] = Lns_ToCollection( self.StaticFlag )
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["kind"] = Lns_ToCollection( self.Kind )
    obj.Items["mutMode"] = Lns_ToCollection( self.MutMode )
    obj.Items["ifList"] = Lns_ToCollection( self.IfList )
    obj.Items["itemTypeId"] = Lns_ToCollection( self.ItemTypeId )
    obj.Items["argTypeId"] = Lns_ToCollection( self.ArgTypeId )
    obj.Items["retTypeId"] = Lns_ToCollection( self.RetTypeId )
    obj.Items["children"] = Lns_ToCollection( self.Children )
    obj.Items["moduleLang"] = Lns_ToCollection( self.ModuleLang )
    obj.Items["requirePath"] = Lns_ToCollection( self.RequirePath )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoNormal) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoNormal__fromMap_2162_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoNormal_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoNormal__fromStem_2165_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoNormal_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoNormal_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoNormal_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoNormal_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoNormal{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoNormal_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoNormal_FromMapMain( newObj *TransUnit__TypeInfoNormal, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["abstractFlag"], false, nil); !ok {
       return false,nil,"abstractFlag:" + mess.(string)
    } else {
       newObj.AbstractFlag = conv.(bool)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["baseId"], false, nil); !ok {
       return false,nil,"baseId:" + mess.(string)
    } else {
       newObj.BaseId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["staticFlag"], false, nil); !ok {
       return false,nil,"staticFlag:" + mess.(string)
    } else {
       newObj.StaticFlag = conv.(bool)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["kind"], false, nil); !ok {
       return false,nil,"kind:" + mess.(string)
    } else {
       newObj.Kind = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["mutMode"], false, nil); !ok {
       return false,nil,"mutMode:" + mess.(string)
    } else {
       newObj.MutMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["ifList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"ifList:" + mess.(string)
    } else {
       newObj.IfList = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["itemTypeId"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"itemTypeId:" + mess.(string)
    } else {
       newObj.ItemTypeId = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["argTypeId"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"argTypeId:" + mess.(string)
    } else {
       newObj.ArgTypeId = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["retTypeId"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"retTypeId:" + mess.(string)
    } else {
       newObj.RetTypeId = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["children"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"children:" + mess.(string)
    } else {
       newObj.Children = conv.(*LnsList)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["moduleLang"], true, nil); !ok {
       return false,nil,"moduleLang:" + mess.(string)
    } else {
       newObj.ModuleLang = conv
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["requirePath"], true, nil); !ok {
       return false,nil,"requirePath:" + mess.(string)
    } else {
       newObj.RequirePath = conv
    }
    return true, newObj, nil
}
// 1597: decl @lune.@base.@TransUnit._TypeInfoNormal.createTypeInfo
func (self *TransUnit__TypeInfoNormal) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    __func__ := "@lune.@base.@TransUnit._TypeInfoNormal.createTypeInfo"
    var newTypeInfo LnsAny
    newTypeInfo = nil
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.ParentId != Ast_rootTypeId) ||
        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_getBuiltInTypeIdMap().Items[self.TypeId])) ||
        Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__List) ||
        Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Array) ||
        Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Map) ||
        Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Set) ).(bool){
        var parentInfo *Ast_TypeInfo
        parentInfo = Ast_headTypeInfo
        if self.ParentId != Ast_rootTypeId{
            var workTypeInfo *Ast_TypeInfo
            
            {
                _workTypeInfo := TransUnit_convExp7058(Lns_2DDD(param.FP.GetTypeInfo(self.ParentId)))
                if _workTypeInfo == nil{
                    return nil, Lns_getVM().String_format("not found parentInfo %s %s", []LnsAny{self.ParentId, self.Txt})
                } else {
                    workTypeInfo = _workTypeInfo.(*Ast_TypeInfo)
                }
            }
            parentInfo = workTypeInfo
            
        }
        var itemTypeInfo *LnsList
        itemTypeInfo = NewLnsList([]LnsAny{})
        for _, _typeId := range( self.ItemTypeId.Items ) {
            typeId := _typeId.(LnsInt)
            itemTypeInfo.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfo(typeId))).(*Ast_TypeInfo)))
        }
        var argTypeInfo *LnsList
        argTypeInfo = NewLnsList([]LnsAny{})
        for _index, _typeId := range( self.ArgTypeId.Items ) {
            index := _index + 1
            typeId := _typeId.(LnsInt)
            var argType LnsAny
            var mess LnsAny
            argType,mess = param.FP.GetTypeInfo(typeId)
            if argType != nil{
                argType_6469 := argType.(*Ast_TypeInfo)
                argTypeInfo.Insert(Ast_TypeInfo2Stem(argType_6469))
            } else {
                var errmess string
                errmess = Lns_getVM().String_format("not found arg (index:%d) -- %s.%s, %d, %d. %s", []LnsAny{index, parentInfo.FP.GetTxt(nil, nil, nil), self.Txt, typeId, self.ArgTypeId.Len(), mess})
                return nil, errmess
            }
        }
        var retTypeInfo *LnsList
        retTypeInfo = NewLnsList([]LnsAny{})
        for _, _typeId := range( self.RetTypeId.Items ) {
            typeId := _typeId.(LnsInt)
            retTypeInfo.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfo(typeId))).(*Ast_TypeInfo)))
        }
        var baseInfo *Ast_TypeInfo
        baseInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.BaseId))).(*Ast_TypeInfo)
        var interfaceList *LnsList
        interfaceList = NewLnsList([]LnsAny{})
        for _, _ifTypeId := range( self.IfList.Items ) {
            ifTypeId := _ifTypeId.(LnsInt)
            interfaceList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfo(ifTypeId))).(*Ast_TypeInfo)))
        }
        var parentScope *Ast_Scope
        
        {
            _parentScope := param.TypeId2Scope.Items[self.ParentId]
            if _parentScope == nil{
                return nil, Lns_getVM().String_format("%s: not found parentScope %s %s", []LnsAny{__func__, self.ParentId, self.Txt})
            } else {
                parentScope = _parentScope.(*Ast_Scope)
            }
        }
        if self.Txt != ""{
            newTypeInfo = parentScope.FP.GetTypeInfoChild(self.Txt)
            
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( newTypeInfo) &&
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__ExtModule) ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__IF) ).(bool))) )){
            {
                __exp := newTypeInfo
                if __exp != nil {
                    _exp := __exp.(*Ast_TypeInfo)
                    param.TypeId2Scope.Set(self.TypeId,Ast_getScope(_exp))
                    if Lns_op_not(_exp.FP.Get_scope()){
                        Util_err(Lns_getVM().String_format("not found scope %s %s %s %s %s", []LnsAny{parentScope, self.ParentId, self.TypeId, self.Txt, _exp.FP.GetTxt(nil, nil, nil)}))
                    }
                    param.TypeId2TypeInfo.Set(self.TypeId,_exp)
                }
            }
        } else { 
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Class) ||
                Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__IF) ).(bool){
                Log_log(Log_Level__Debug, __func__, 1676, Log_CreateMessage(func() string {
                    return Lns_getVM().String_format("new type -- %d, %s -- %s, %d", []LnsAny{self.ParentId, self.Txt, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentScope.FP.Get_ownerTypeInfo()) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.GetFullName(Ast_defaultTypeNameCtrl, parentScope.FP, false)})/* 1:65 */)) ||
                        Lns_GetEnv().SetStackVal( "nil") ).(string), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentScope.FP.Get_ownerTypeInfo()) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_typeId()}))) ||
                        Lns_GetEnv().SetStackVal( -1) ).(LnsInt)})
                }))
                
                var baseScope *Ast_Scope
                baseScope = Lns_unwrap( param.TypeId2Scope.Items[self.BaseId]).(*Ast_Scope)
                var ifScopeList *LnsList
                ifScopeList = NewLnsList([]LnsAny{})
                for _, _ifType := range( interfaceList.Items ) {
                    ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    ifScopeList.Insert(Ast_Scope2Stem(Lns_unwrap( ifType.FP.Get_scope()).(*Ast_Scope)))
                }
                var scope *Ast_Scope
                scope = NewAst_Scope(param.ProcessInfo, parentScope, true, baseScope, ifScopeList)
                var altTypeList *LnsList
                altTypeList = NewLnsList([]LnsAny{})
                for _, _itemType := range( itemTypeInfo.Items ) {
                    itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    altTypeList.Insert(Ast_AlternateTypeInfo2Stem(Lns_unwrap( (Ast_AlternateTypeInfoDownCastF(itemType.FP))).(*Ast_AlternateTypeInfo)))
                }
                var workTypeInfo *Ast_TypeInfo
                workTypeInfo = param.ProcessInfo.FP.CreateClass(self.Kind == Ast_TypeInfoKind__Class, self.AbstractFlag, scope, baseInfo, interfaceList, altTypeList, parentInfo, true, Ast_AccessMode__Pub, self.Txt)
                parentScope.FP.AddClassLazy(param.ProcessInfo, self.Txt, nil, workTypeInfo, param.LazyModuleSet.Has(self.TypeId))
                newTypeInfo = workTypeInfo
                
                param.TypeId2Scope.Set(self.TypeId,scope)
                param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
            } else if self.Kind == Ast_TypeInfoKind__ExtModule{
                Log_log(Log_Level__Debug, __func__, 1713, Log_CreateMessage(func() string {
                    return Lns_getVM().String_format("new type -- %d, %s -- %s, %d", []LnsAny{self.ParentId, self.Txt, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentScope.FP.Get_ownerTypeInfo()) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.GetFullName(Ast_defaultTypeNameCtrl, parentScope.FP, false)})/* 1:65 */)) ||
                        Lns_GetEnv().SetStackVal( "nil") ).(string), Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(parentScope.FP.Get_ownerTypeInfo()) && 
                        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_typeId()}))) ||
                        Lns_GetEnv().SetStackVal( -1) ).(LnsInt)})
                }))
                
                var scope *Ast_Scope
                scope = NewAst_Scope(param.ProcessInfo, parentScope, true, nil, NewLnsList([]LnsAny{}))
                var workTypeInfo *Ast_TypeInfo
                workTypeInfo = param.ProcessInfo.FP.CreateExtModule(scope, parentInfo, true, Ast_AccessMode__Pub, self.Txt, Lns_unwrap( self.ModuleLang).(LnsInt), Lns_unwrap( self.RequirePath).(string))
                parentScope.FP.AddExtModule(param.ProcessInfo, self.Txt, nil, workTypeInfo, param.LazyModuleSet.Has(self.TypeId), Lns_unwrap( self.ModuleLang).(LnsInt))
                newTypeInfo = workTypeInfo
                
                param.TypeId2Scope.Set(self.TypeId,scope)
                param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
            } else { 
                var scope LnsAny
                scope = nil
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Func) ||
                    Lns_GetEnv().SetStackVal( self.Kind == Ast_TypeInfoKind__Method) ).(bool){
                    scope = NewAst_Scope(param.ProcessInfo, parentScope, false, nil, nil)
                    
                }
                var typeInfoKind LnsInt
                typeInfoKind = self.Kind
                var accessMode LnsInt
                accessMode = self.AccessMode
                var workTypeInfo *Ast_TypeInfo
                workTypeInfo = Ast_NormalTypeInfo_create(param.ProcessInfo, accessMode, self.AbstractFlag, scope, baseInfo, parentInfo, self.StaticFlag, typeInfoKind, self.Txt, itemTypeInfo, argTypeInfo, retTypeInfo, self.MutMode)
                newTypeInfo = workTypeInfo
                
                param.TypeId2TypeInfo.Set(self.TypeId,workTypeInfo)
                if _switch8214 := self.Kind; _switch8214 == Ast_TypeInfoKind__Func || _switch8214 == Ast_TypeInfoKind__Method || _switch8214 == Ast_TypeInfoKind__Macro || _switch8214 == Ast_TypeInfoKind__Form || _switch8214 == Ast_TypeInfoKind__FormFunc {
                    var symbolKind LnsInt
                    symbolKind = Ast_SymbolKind__Fun
                    if _switch8152 := self.Kind; _switch8152 == Ast_TypeInfoKind__Method {
                        symbolKind = Ast_SymbolKind__Mtd
                        
                    } else if _switch8152 == Ast_TypeInfoKind__Macro {
                        symbolKind = Ast_SymbolKind__Mac
                        
                    } else if _switch8152 == Ast_TypeInfoKind__Form || _switch8152 == Ast_TypeInfoKind__FormFunc {
                        symbolKind = Ast_SymbolKind__Typ
                        
                    }
                    var workParentScope *Ast_Scope
                    workParentScope = Lns_unwrap( param.TypeId2Scope.Items[self.ParentId]).(*Ast_Scope)
                    workParentScope.FP.Add(param.ProcessInfo, symbolKind, false, self.Kind == Ast_TypeInfoKind__Func, self.Txt, nil, workTypeInfo, accessMode, self.StaticFlag, Ast_MutMode__IMut, true, false)
                    param.TypeId2Scope.Set(self.TypeId,scope)
                }
            }
        }
    } else { 
        newTypeInfo = param.Scope.FP.GetTypeInfo(self.Txt, param.Scope, false, param.ScopeAccess)
        
        if Lns_op_not(newTypeInfo){
            for _key, _val := range( self.FP.ToMap().Items ) {
                key := _key.(string)
                val := _val
                Util_errorLog(Lns_getVM().String_format("error: illegal self %s:%s", []LnsAny{key, val}))
            }
        }
        param.TypeId2TypeInfo.Set(self.TypeId,Lns_unwrap( newTypeInfo).(*Ast_TypeInfo))
    }
    return newTypeInfo, nil
}


// declaration Class -- _TypeInfoEnum
type TransUnit__TypeInfoEnumMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoEnum struct {
    TransUnit__TypeInfo
    ParentId LnsInt
    Txt string
    AccessMode LnsInt
    ValTypeId LnsInt
    EnumValList *LnsMap
    FP TransUnit__TypeInfoEnumMtd
}
func TransUnit__TypeInfoEnum2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoEnum).FP
}
type TransUnit__TypeInfoEnumDownCast interface {
    ToTransUnit__TypeInfoEnum() *TransUnit__TypeInfoEnum
}
func TransUnit__TypeInfoEnumDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoEnumDownCast)
    if ok { return work.ToTransUnit__TypeInfoEnum() }
    return nil
}
func (obj *TransUnit__TypeInfoEnum) ToTransUnit__TypeInfoEnum() *TransUnit__TypeInfoEnum {
    return obj
}
func NewTransUnit__TypeInfoEnum(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 LnsInt, arg5 *LnsMap) *TransUnit__TypeInfoEnum {
    obj := &TransUnit__TypeInfoEnum{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoEnum(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *TransUnit__TypeInfoEnum) InitTransUnit__TypeInfoEnum(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 LnsInt, arg5 *LnsMap) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.ParentId = arg1
    self.Txt = arg2
    self.AccessMode = arg3
    self.ValTypeId = arg4
    self.EnumValList = arg5
}
func (self *TransUnit__TypeInfoEnum) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["valTypeId"] = Lns_ToCollection( self.ValTypeId )
    obj.Items["enumValList"] = Lns_ToCollection( self.EnumValList )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoEnum) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoEnum__fromMap_2193_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoEnum_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoEnum__fromStem_2196_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoEnum_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoEnum_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoEnum_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoEnum_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoEnum{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoEnum_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoEnum_FromMapMain( newObj *TransUnit__TypeInfoEnum, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["valTypeId"], false, nil); !ok {
       return false,nil,"valTypeId:" + mess.(string)
    } else {
       newObj.ValTypeId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToLnsMapSub( objMap.Items["enumValList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToStrSub, false,nil},Lns_ToObjParam{
            Lns_ToStemSub, false,nil}}); !ok {
       return false,nil,"enumValList:" + mess.(string)
    } else {
       newObj.EnumValList = conv.(*LnsMap)
    }
    return true, newObj, nil
}
// 1805: decl @lune.@base.@TransUnit._TypeInfoEnum.createTypeInfo
func (self *TransUnit__TypeInfoEnum) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    var accessMode LnsInt
    accessMode = Lns_unwrap( Ast_AccessMode__from(LnsInt(self.AccessMode))).(LnsInt)
    var parentInfo *Ast_TypeInfo
    parentInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.ParentId))).(*Ast_TypeInfo)
    var parentScope *Ast_Scope
    parentScope = Lns_unwrap( Ast_getScope(parentInfo)).(*Ast_Scope)
    var scope *Ast_Scope
    scope = NewAst_Scope(param.ProcessInfo, parentScope, true, nil, nil)
    param.TypeId2Scope.Set(self.TypeId,scope)
    var valTypeInfo *Ast_TypeInfo
    valTypeInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.ValTypeId))).(*Ast_TypeInfo)
    var enumTypeInfo *Ast_EnumTypeInfo
    enumTypeInfo = param.ProcessInfo.FP.CreateEnum(scope, parentInfo, true, accessMode, self.Txt, valTypeInfo)
    var newTypeInfo *Ast_EnumTypeInfo
    newTypeInfo = enumTypeInfo
    param.TypeId2TypeInfo.Set(self.TypeId,&enumTypeInfo.Ast_TypeInfo)
    var getEnumLiteral func(val LnsAny) LnsAny
    getEnumLiteral = func(val LnsAny) LnsAny {
        if _switch8511 := valTypeInfo; _switch8511 == Ast_builtinTypeInt {
            return &Ast_EnumLiteral__Int{Lns_forceCastInt(val)}
        } else if _switch8511 == Ast_builtinTypeReal {
            return &Ast_EnumLiteral__Real{Lns_forceCastReal(val)}
        } else if _switch8511 == Ast_builtinTypeString {
            return &Ast_EnumLiteral__Str{val.(string)}
        }
        return nil
    }
    for _valName, _valData := range( self.EnumValList.Items ) {
        valName := _valName.(string)
        valData := _valData
        var val LnsAny
        
        {
            _val := getEnumLiteral(valData)
            if _val == nil{
                return nil, Lns_getVM().String_format("unknown enum val type -- %s", []LnsAny{valTypeInfo.FP.GetTxt(nil, nil, nil)})
            } else {
                val = _val
            }
        }
        enumTypeInfo.FP.AddEnumValInfo(NewAst_EnumValInfo(valName, val))
        scope.FP.AddEnumVal(param.ProcessInfo, valName, nil, &enumTypeInfo.Ast_TypeInfo)
    }
    parentScope.FP.AddEnum(param.ProcessInfo, accessMode, self.Txt, nil, &enumTypeInfo.Ast_TypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- _TypeInfoAlgeVal
type TransUnit__TypeInfoAlgeValMtd interface {
    ToMap() *LnsMap
}
type TransUnit__TypeInfoAlgeVal struct {
    Name string
    TypeList *LnsList
    FP TransUnit__TypeInfoAlgeValMtd
}
func TransUnit__TypeInfoAlgeVal2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoAlgeVal).FP
}
type TransUnit__TypeInfoAlgeValDownCast interface {
    ToTransUnit__TypeInfoAlgeVal() *TransUnit__TypeInfoAlgeVal
}
func TransUnit__TypeInfoAlgeValDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoAlgeValDownCast)
    if ok { return work.ToTransUnit__TypeInfoAlgeVal() }
    return nil
}
func (obj *TransUnit__TypeInfoAlgeVal) ToTransUnit__TypeInfoAlgeVal() *TransUnit__TypeInfoAlgeVal {
    return obj
}
func NewTransUnit__TypeInfoAlgeVal(arg1 string, arg2 *LnsList) *TransUnit__TypeInfoAlgeVal {
    obj := &TransUnit__TypeInfoAlgeVal{}
    obj.FP = obj
    obj.InitTransUnit__TypeInfoAlgeVal(arg1, arg2)
    return obj
}
func (self *TransUnit__TypeInfoAlgeVal) InitTransUnit__TypeInfoAlgeVal(arg1 string, arg2 *LnsList) {
    self.Name = arg1
    self.TypeList = arg2
}
func (self *TransUnit__TypeInfoAlgeVal) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["name"] = Lns_ToCollection( self.Name )
    obj.Items["typeList"] = Lns_ToCollection( self.TypeList )
    return obj
}
func (self *TransUnit__TypeInfoAlgeVal) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoAlgeVal__fromMap_2217_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoAlgeVal_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoAlgeVal__fromStem_2220_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoAlgeVal_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoAlgeVal_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoAlgeVal_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoAlgeVal_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoAlgeVal{}
    newObj.FP = newObj
    return TransUnit__TypeInfoAlgeVal_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoAlgeVal_FromMapMain( newObj *TransUnit__TypeInfoAlgeVal, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["name"], false, nil); !ok {
       return false,nil,"name:" + mess.(string)
    } else {
       newObj.Name = conv.(string)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["typeList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            Lns_ToIntSub, false,nil}}); !ok {
       return false,nil,"typeList:" + mess.(string)
    } else {
       newObj.TypeList = conv.(*LnsList)
    }
    return true, newObj, nil
}

// declaration Class -- _TypeInfoAlge
type TransUnit__TypeInfoAlgeMtd interface {
    ToMap() *LnsMap
    CreateTypeInfo(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
    CreateTypeInfoCache(arg1 *TransUnit_ImportParam)(LnsAny, LnsAny)
}
type TransUnit__TypeInfoAlge struct {
    TransUnit__TypeInfo
    ParentId LnsInt
    Txt string
    AccessMode LnsInt
    AlgeValList *LnsList
    FP TransUnit__TypeInfoAlgeMtd
}
func TransUnit__TypeInfoAlge2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit__TypeInfoAlge).FP
}
type TransUnit__TypeInfoAlgeDownCast interface {
    ToTransUnit__TypeInfoAlge() *TransUnit__TypeInfoAlge
}
func TransUnit__TypeInfoAlgeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit__TypeInfoAlgeDownCast)
    if ok { return work.ToTransUnit__TypeInfoAlge() }
    return nil
}
func (obj *TransUnit__TypeInfoAlge) ToTransUnit__TypeInfoAlge() *TransUnit__TypeInfoAlge {
    return obj
}
func NewTransUnit__TypeInfoAlge(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *LnsList) *TransUnit__TypeInfoAlge {
    obj := &TransUnit__TypeInfoAlge{}
    obj.FP = obj
    obj.TransUnit__TypeInfo.FP = obj
    obj.InitTransUnit__TypeInfoAlge(arg1, arg2, arg3, arg4)
    return obj
}
func (self *TransUnit__TypeInfoAlge) InitTransUnit__TypeInfoAlge(arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 *LnsList) {
    self.TransUnit__TypeInfo.InitTransUnit__TypeInfo( )
    self.ParentId = arg1
    self.Txt = arg2
    self.AccessMode = arg3
    self.AlgeValList = arg4
}
func (self *TransUnit__TypeInfoAlge) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["parentId"] = Lns_ToCollection( self.ParentId )
    obj.Items["txt"] = Lns_ToCollection( self.Txt )
    obj.Items["accessMode"] = Lns_ToCollection( self.AccessMode )
    obj.Items["algeValList"] = Lns_ToCollection( self.AlgeValList )
    return self.TransUnit__TypeInfo.ToMapSetup( obj )
}
func (self *TransUnit__TypeInfoAlge) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func TransUnit__TypeInfoAlge__fromMap_2252_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoAlge_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoAlge__fromStem_2255_(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return TransUnit__TypeInfoAlge_FromMap( arg1, paramList )
}
func TransUnit__TypeInfoAlge_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := TransUnit__TypeInfoAlge_FromMapSub(obj,false, paramList);
    return conv,mess
}
func TransUnit__TypeInfoAlge_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &TransUnit__TypeInfoAlge{}
    newObj.FP = newObj
    newObj.TransUnit__TypeInfo.FP = newObj
    return TransUnit__TypeInfoAlge_FromMapMain( newObj, objMap, paramList )
}
func TransUnit__TypeInfoAlge_FromMapMain( newObj *TransUnit__TypeInfoAlge, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,_,mess := TransUnit__TypeInfo_FromMapMain( &newObj.TransUnit__TypeInfo, objMap, paramList ); !ok {
        return false,nil,mess
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["parentId"], false, nil); !ok {
       return false,nil,"parentId:" + mess.(string)
    } else {
       newObj.ParentId = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToStrSub( objMap.Items["txt"], false, nil); !ok {
       return false,nil,"txt:" + mess.(string)
    } else {
       newObj.Txt = conv.(string)
    }
    if ok,conv,mess := Lns_ToIntSub( objMap.Items["accessMode"], false, nil); !ok {
       return false,nil,"accessMode:" + mess.(string)
    } else {
       newObj.AccessMode = conv.(LnsInt)
    }
    if ok,conv,mess := Lns_ToListSub( objMap.Items["algeValList"], false, []Lns_ToObjParam{Lns_ToObjParam{
            TransUnit__TypeInfoAlgeVal_FromMapSub, false,nil}}); !ok {
       return false,nil,"algeValList:" + mess.(string)
    } else {
       newObj.AlgeValList = conv.(*LnsList)
    }
    return true, newObj, nil
}
// 1858: decl @lune.@base.@TransUnit._TypeInfoAlge.createTypeInfo
func (self *TransUnit__TypeInfoAlge) CreateTypeInfo(param *TransUnit_ImportParam)(LnsAny, LnsAny) {
    var accessMode LnsInt
    accessMode = Lns_unwrap( Ast_AccessMode__from(LnsInt(self.AccessMode))).(LnsInt)
    var parentInfo *Ast_TypeInfo
    parentInfo = Lns_unwrap( Lns_car(param.FP.GetTypeInfo(self.ParentId))).(*Ast_TypeInfo)
    var parentScope *Ast_Scope
    parentScope = Lns_unwrap( Ast_getScope(parentInfo)).(*Ast_Scope)
    var scope *Ast_Scope
    scope = NewAst_Scope(param.ProcessInfo, parentScope, true, nil, nil)
    param.TypeId2Scope.Set(self.TypeId,scope)
    var algeTypeInfo *Ast_AlgeTypeInfo
    algeTypeInfo = param.ProcessInfo.FP.CreateAlge(scope, parentInfo, true, accessMode, self.Txt)
    var newTypeInfo *Ast_AlgeTypeInfo
    newTypeInfo = algeTypeInfo
    param.TypeId2TypeInfo.Set(self.TypeId,&algeTypeInfo.Ast_TypeInfo)
    for _, _valInfo := range( self.AlgeValList.Items ) {
        valInfo := _valInfo.(TransUnit__TypeInfoAlgeValDownCast).ToTransUnit__TypeInfoAlgeVal()
        var typeInfoList *LnsList
        typeInfoList = NewLnsList([]LnsAny{})
        for _, _orgTypeId := range( valInfo.TypeList.Items ) {
            orgTypeId := _orgTypeId.(LnsInt)
            typeInfoList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( Lns_car(param.FP.GetTypeInfo(orgTypeId))).(*Ast_TypeInfo)))
        }
        var algeVal *Ast_AlgeValInfo
        algeVal = NewAst_AlgeValInfo(valInfo.Name, typeInfoList, algeTypeInfo)
        scope.FP.AddAlgeVal(param.ProcessInfo, valInfo.Name, nil, &algeTypeInfo.Ast_TypeInfo)
        algeTypeInfo.FP.AddValInfo(algeVal)
    }
    parentScope.FP.AddAlge(param.ProcessInfo, accessMode, self.Txt, nil, &algeTypeInfo.Ast_TypeInfo)
    return &newTypeInfo.Ast_TypeInfo, nil
}


// declaration Class -- BuiltinFuncType
type TransUnit_BuiltinFuncTypeMtd interface {
    Get_allClass() *LnsList
    Get_allFuncTypeSet() *LnsSet
    Get_allSymbol() *LnsList
    Get_allSymbolSet() *LnsSet
    Get_needThreadingTypes() *LnsSet
    Register(arg1 *Ast_SymbolInfo)
    RegisterClass(arg1 *Ast_TypeInfo)
}
type TransUnit_BuiltinFuncType struct {
    Lns_ *Ast_TypeInfo
    Lns__fcall *Ast_TypeInfo
    Lns__fcall_sym *Ast_SymbolInfo
    Lns__kind *Ast_TypeInfo
    Lns__kind_sym *Ast_SymbolInfo
    Lns__load *Ast_TypeInfo
    Lns__load_sym *Ast_SymbolInfo
    Lns_collectgarbage *Ast_TypeInfo
    Lns_collectgarbage_sym *Ast_SymbolInfo
    Lns_error *Ast_TypeInfo
    Lns_error_sym *Ast_SymbolInfo
    Lns_expandLuavalMap *Ast_TypeInfo
    Lns_expandLuavalMap_sym *Ast_SymbolInfo
    Lns_loadfile *Ast_TypeInfo
    Lns_loadfile_sym *Ast_SymbolInfo
    Lns_print *Ast_TypeInfo
    Lns_print_sym *Ast_SymbolInfo
    Lns_require *Ast_TypeInfo
    Lns_require_sym *Ast_SymbolInfo
    Lns_tonumber *Ast_TypeInfo
    Lns_tonumber_sym *Ast_SymbolInfo
    Lns_tostring *Ast_TypeInfo
    Lns_tostring_sym *Ast_SymbolInfo
    Lns_type *Ast_TypeInfo
    Lns_type_sym *Ast_SymbolInfo
    Istream_ *Ast_TypeInfo
    Istream___attrib *Ast_TypeInfo
    Istream___attrib_sym *Ast_SymbolInfo
    Istream_close *Ast_TypeInfo
    Istream_close_sym *Ast_SymbolInfo
    Istream_read *Ast_TypeInfo
    Istream_read_sym *Ast_SymbolInfo
    Ostream_ *Ast_TypeInfo
    Ostream___attrib *Ast_TypeInfo
    Ostream___attrib_sym *Ast_SymbolInfo
    Ostream_close *Ast_TypeInfo
    Ostream_close_sym *Ast_SymbolInfo
    Ostream_flush *Ast_TypeInfo
    Ostream_flush_sym *Ast_SymbolInfo
    Ostream_write *Ast_TypeInfo
    Ostream_write_sym *Ast_SymbolInfo
    Luastream_ *Ast_TypeInfo
    Luastream___attrib *Ast_TypeInfo
    Luastream___attrib_sym *Ast_SymbolInfo
    Luastream_close *Ast_TypeInfo
    Luastream_close_sym *Ast_SymbolInfo
    Luastream_flush *Ast_TypeInfo
    Luastream_flush_sym *Ast_SymbolInfo
    Luastream_read *Ast_TypeInfo
    Luastream_read_sym *Ast_SymbolInfo
    Luastream_seek *Ast_TypeInfo
    Luastream_seek_sym *Ast_SymbolInfo
    Luastream_write *Ast_TypeInfo
    Luastream_write_sym *Ast_SymbolInfo
    Mapping_ *Ast_TypeInfo
    Mapping___attrib *Ast_TypeInfo
    Mapping___attrib_sym *Ast_SymbolInfo
    Mapping__toMap *Ast_TypeInfo
    Mapping__toMap_sym *Ast_SymbolInfo
    G__pipe_ *Ast_TypeInfo
    G__pipe_get *Ast_TypeInfo
    G__pipe_get_sym *Ast_SymbolInfo
    G__pipe_put *Ast_TypeInfo
    G__pipe_put_sym *Ast_SymbolInfo
    Lnsthread_ *Ast_TypeInfo
    Lnsthread___init *Ast_TypeInfo
    Lnsthread___init_sym *Ast_SymbolInfo
    Lnsthread_loop *Ast_TypeInfo
    Lnsthread_loop_sym *Ast_SymbolInfo
    Io_ *Ast_TypeInfo
    Io___attrib *Ast_TypeInfo
    Io___attrib_sym *Ast_SymbolInfo
    Io_open *Ast_TypeInfo
    Io_open_sym *Ast_SymbolInfo
    Io_popen *Ast_TypeInfo
    Io_popen_sym *Ast_SymbolInfo
    Io_stderr *Ast_TypeInfo
    Io_stderr_sym *Ast_SymbolInfo
    Io_stdin *Ast_TypeInfo
    Io_stdin_sym *Ast_SymbolInfo
    Io_stdout *Ast_TypeInfo
    Io_stdout_sym *Ast_SymbolInfo
    Package_ *Ast_TypeInfo
    Package___attrib *Ast_TypeInfo
    Package___attrib_sym *Ast_SymbolInfo
    Package_path *Ast_TypeInfo
    Package_path_sym *Ast_SymbolInfo
    Package_searchpath *Ast_TypeInfo
    Package_searchpath_sym *Ast_SymbolInfo
    Os_ *Ast_TypeInfo
    Os___attrib *Ast_TypeInfo
    Os___attrib_sym *Ast_SymbolInfo
    Os_clock *Ast_TypeInfo
    Os_clock_sym *Ast_SymbolInfo
    Os_date *Ast_TypeInfo
    Os_date_sym *Ast_SymbolInfo
    Os_difftime *Ast_TypeInfo
    Os_difftime_sym *Ast_SymbolInfo
    Os_exit *Ast_TypeInfo
    Os_exit_sym *Ast_SymbolInfo
    Os_remove *Ast_TypeInfo
    Os_remove_sym *Ast_SymbolInfo
    Os_rename *Ast_TypeInfo
    Os_rename_sym *Ast_SymbolInfo
    Os_time *Ast_TypeInfo
    Os_time_sym *Ast_SymbolInfo
    String_ *Ast_TypeInfo
    String___attrib *Ast_TypeInfo
    String___attrib_sym *Ast_SymbolInfo
    String_byte *Ast_TypeInfo
    String_byte_sym *Ast_SymbolInfo
    String_dump *Ast_TypeInfo
    String_dump_sym *Ast_SymbolInfo
    String_find *Ast_TypeInfo
    String_find_sym *Ast_SymbolInfo
    String_format *Ast_TypeInfo
    String_format_sym *Ast_SymbolInfo
    String_gmatch *Ast_TypeInfo
    String_gmatch_sym *Ast_SymbolInfo
    String_gsub *Ast_TypeInfo
    String_gsub_sym *Ast_SymbolInfo
    String_lower *Ast_TypeInfo
    String_lower_sym *Ast_SymbolInfo
    String_rep *Ast_TypeInfo
    String_rep_sym *Ast_SymbolInfo
    String_reverse *Ast_TypeInfo
    String_reverse_sym *Ast_SymbolInfo
    String_sub *Ast_TypeInfo
    String_sub_sym *Ast_SymbolInfo
    String_upper *Ast_TypeInfo
    String_upper_sym *Ast_SymbolInfo
    Str_ *Ast_TypeInfo
    Str___attrib *Ast_TypeInfo
    Str___attrib_sym *Ast_SymbolInfo
    Str_byte *Ast_TypeInfo
    Str_byte_sym *Ast_SymbolInfo
    Str_find *Ast_TypeInfo
    Str_find_sym *Ast_SymbolInfo
    Str_format *Ast_TypeInfo
    Str_format_sym *Ast_SymbolInfo
    Str_gmatch *Ast_TypeInfo
    Str_gmatch_sym *Ast_SymbolInfo
    Str_gsub *Ast_TypeInfo
    Str_gsub_sym *Ast_SymbolInfo
    Str_lower *Ast_TypeInfo
    Str_lower_sym *Ast_SymbolInfo
    Str_rep *Ast_TypeInfo
    Str_rep_sym *Ast_SymbolInfo
    Str_reverse *Ast_TypeInfo
    Str_reverse_sym *Ast_SymbolInfo
    Str_sub *Ast_TypeInfo
    Str_sub_sym *Ast_SymbolInfo
    Str_upper *Ast_TypeInfo
    Str_upper_sym *Ast_SymbolInfo
    List_ *Ast_TypeInfo
    List___less *Ast_TypeInfo
    List___less_sym *Ast_SymbolInfo
    List_insert *Ast_TypeInfo
    List_insert_sym *Ast_SymbolInfo
    List_remove *Ast_TypeInfo
    List_remove_sym *Ast_SymbolInfo
    List_sort *Ast_TypeInfo
    List_sort_sym *Ast_SymbolInfo
    List_unpack *Ast_TypeInfo
    List_unpack_sym *Ast_SymbolInfo
    Array_ *Ast_TypeInfo
    Array___less *Ast_TypeInfo
    Array___less_sym *Ast_SymbolInfo
    Array_sort *Ast_TypeInfo
    Array_sort_sym *Ast_SymbolInfo
    Array_unpack *Ast_TypeInfo
    Array_unpack_sym *Ast_SymbolInfo
    Set_ *Ast_TypeInfo
    Set_add *Ast_TypeInfo
    Set_add_sym *Ast_SymbolInfo
    Set_and *Ast_TypeInfo
    Set_and_sym *Ast_SymbolInfo
    Set_clone *Ast_TypeInfo
    Set_clone_sym *Ast_SymbolInfo
    Set_del *Ast_TypeInfo
    Set_del_sym *Ast_SymbolInfo
    Set_has *Ast_TypeInfo
    Set_has_sym *Ast_SymbolInfo
    Set_len *Ast_TypeInfo
    Set_len_sym *Ast_SymbolInfo
    Set_or *Ast_TypeInfo
    Set_or_sym *Ast_SymbolInfo
    Set_sub *Ast_TypeInfo
    Set_sub_sym *Ast_SymbolInfo
    Math_ *Ast_TypeInfo
    Math___attrib *Ast_TypeInfo
    Math___attrib_sym *Ast_SymbolInfo
    Math_random *Ast_TypeInfo
    Math_random_sym *Ast_SymbolInfo
    Math_randomseed *Ast_TypeInfo
    Math_randomseed_sym *Ast_SymbolInfo
    Debug_ *Ast_TypeInfo
    Debug___attrib *Ast_TypeInfo
    Debug___attrib_sym *Ast_SymbolInfo
    Debug_getinfo *Ast_TypeInfo
    Debug_getinfo_sym *Ast_SymbolInfo
    Debug_getlocal *Ast_TypeInfo
    Debug_getlocal_sym *Ast_SymbolInfo
    Nilable_ *Ast_TypeInfo
    Nilable_val *Ast_TypeInfo
    Nilable_val_sym *Ast_SymbolInfo
    allSymbol *LnsList
    allClass *LnsList
    allFuncTypeSet *LnsSet
    allSymbolSet *LnsSet
    needThreadingTypes *LnsSet
    FP TransUnit_BuiltinFuncTypeMtd
}
func TransUnit_BuiltinFuncType2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_BuiltinFuncType).FP
}
type TransUnit_BuiltinFuncTypeDownCast interface {
    ToTransUnit_BuiltinFuncType() *TransUnit_BuiltinFuncType
}
func TransUnit_BuiltinFuncTypeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_BuiltinFuncTypeDownCast)
    if ok { return work.ToTransUnit_BuiltinFuncType() }
    return nil
}
func (obj *TransUnit_BuiltinFuncType) ToTransUnit_BuiltinFuncType() *TransUnit_BuiltinFuncType {
    return obj
}
func NewTransUnit_BuiltinFuncType() *TransUnit_BuiltinFuncType {
    obj := &TransUnit_BuiltinFuncType{}
    obj.FP = obj
    obj.InitTransUnit_BuiltinFuncType()
    return obj
}
func (self *TransUnit_BuiltinFuncType) Get_allSymbol() *LnsList{ return self.allSymbol }
func (self *TransUnit_BuiltinFuncType) Get_allClass() *LnsList{ return self.allClass }
func (self *TransUnit_BuiltinFuncType) Get_allFuncTypeSet() *LnsSet{ return self.allFuncTypeSet }
func (self *TransUnit_BuiltinFuncType) Get_allSymbolSet() *LnsSet{ return self.allSymbolSet }
func (self *TransUnit_BuiltinFuncType) Get_needThreadingTypes() *LnsSet{ return self.needThreadingTypes }
// 1966: DeclConstr
func (self *TransUnit_BuiltinFuncType) InitTransUnit_BuiltinFuncType() {
    self.Lns_ = Ast_headTypeInfo
    
    self.Lns__fcall = Ast_headTypeInfo
    
    self.Lns__fcall_sym = Ast_dummySymbol
    
    self.Lns__kind = Ast_headTypeInfo
    
    self.Lns__kind_sym = Ast_dummySymbol
    
    self.Lns__load = Ast_headTypeInfo
    
    self.Lns__load_sym = Ast_dummySymbol
    
    self.Lns_collectgarbage = Ast_headTypeInfo
    
    self.Lns_collectgarbage_sym = Ast_dummySymbol
    
    self.Lns_error = Ast_headTypeInfo
    
    self.Lns_error_sym = Ast_dummySymbol
    
    self.Lns_expandLuavalMap = Ast_headTypeInfo
    
    self.Lns_expandLuavalMap_sym = Ast_dummySymbol
    
    self.Lns_loadfile = Ast_headTypeInfo
    
    self.Lns_loadfile_sym = Ast_dummySymbol
    
    self.Lns_print = Ast_headTypeInfo
    
    self.Lns_print_sym = Ast_dummySymbol
    
    self.Lns_require = Ast_headTypeInfo
    
    self.Lns_require_sym = Ast_dummySymbol
    
    self.Lns_tonumber = Ast_headTypeInfo
    
    self.Lns_tonumber_sym = Ast_dummySymbol
    
    self.Lns_tostring = Ast_headTypeInfo
    
    self.Lns_tostring_sym = Ast_dummySymbol
    
    self.Lns_type = Ast_headTypeInfo
    
    self.Lns_type_sym = Ast_dummySymbol
    
    self.Istream_ = Ast_headTypeInfo
    
    self.Istream___attrib = Ast_headTypeInfo
    
    self.Istream___attrib_sym = Ast_dummySymbol
    
    self.Istream_close = Ast_headTypeInfo
    
    self.Istream_close_sym = Ast_dummySymbol
    
    self.Istream_read = Ast_headTypeInfo
    
    self.Istream_read_sym = Ast_dummySymbol
    
    self.Ostream_ = Ast_headTypeInfo
    
    self.Ostream___attrib = Ast_headTypeInfo
    
    self.Ostream___attrib_sym = Ast_dummySymbol
    
    self.Ostream_close = Ast_headTypeInfo
    
    self.Ostream_close_sym = Ast_dummySymbol
    
    self.Ostream_flush = Ast_headTypeInfo
    
    self.Ostream_flush_sym = Ast_dummySymbol
    
    self.Ostream_write = Ast_headTypeInfo
    
    self.Ostream_write_sym = Ast_dummySymbol
    
    self.Luastream_ = Ast_headTypeInfo
    
    self.Luastream___attrib = Ast_headTypeInfo
    
    self.Luastream___attrib_sym = Ast_dummySymbol
    
    self.Luastream_close = Ast_headTypeInfo
    
    self.Luastream_close_sym = Ast_dummySymbol
    
    self.Luastream_flush = Ast_headTypeInfo
    
    self.Luastream_flush_sym = Ast_dummySymbol
    
    self.Luastream_read = Ast_headTypeInfo
    
    self.Luastream_read_sym = Ast_dummySymbol
    
    self.Luastream_seek = Ast_headTypeInfo
    
    self.Luastream_seek_sym = Ast_dummySymbol
    
    self.Luastream_write = Ast_headTypeInfo
    
    self.Luastream_write_sym = Ast_dummySymbol
    
    self.Mapping_ = Ast_headTypeInfo
    
    self.Mapping___attrib = Ast_headTypeInfo
    
    self.Mapping___attrib_sym = Ast_dummySymbol
    
    self.Mapping__toMap = Ast_headTypeInfo
    
    self.Mapping__toMap_sym = Ast_dummySymbol
    
    self.G__pipe_ = Ast_headTypeInfo
    
    self.G__pipe_get = Ast_headTypeInfo
    
    self.G__pipe_get_sym = Ast_dummySymbol
    
    self.G__pipe_put = Ast_headTypeInfo
    
    self.G__pipe_put_sym = Ast_dummySymbol
    
    self.Lnsthread_ = Ast_headTypeInfo
    
    self.Lnsthread___init = Ast_headTypeInfo
    
    self.Lnsthread___init_sym = Ast_dummySymbol
    
    self.Lnsthread_loop = Ast_headTypeInfo
    
    self.Lnsthread_loop_sym = Ast_dummySymbol
    
    self.Io_ = Ast_headTypeInfo
    
    self.Io___attrib = Ast_headTypeInfo
    
    self.Io___attrib_sym = Ast_dummySymbol
    
    self.Io_open = Ast_headTypeInfo
    
    self.Io_open_sym = Ast_dummySymbol
    
    self.Io_popen = Ast_headTypeInfo
    
    self.Io_popen_sym = Ast_dummySymbol
    
    self.Io_stderr = Ast_headTypeInfo
    
    self.Io_stderr_sym = Ast_dummySymbol
    
    self.Io_stdin = Ast_headTypeInfo
    
    self.Io_stdin_sym = Ast_dummySymbol
    
    self.Io_stdout = Ast_headTypeInfo
    
    self.Io_stdout_sym = Ast_dummySymbol
    
    self.Package_ = Ast_headTypeInfo
    
    self.Package___attrib = Ast_headTypeInfo
    
    self.Package___attrib_sym = Ast_dummySymbol
    
    self.Package_path = Ast_headTypeInfo
    
    self.Package_path_sym = Ast_dummySymbol
    
    self.Package_searchpath = Ast_headTypeInfo
    
    self.Package_searchpath_sym = Ast_dummySymbol
    
    self.Os_ = Ast_headTypeInfo
    
    self.Os___attrib = Ast_headTypeInfo
    
    self.Os___attrib_sym = Ast_dummySymbol
    
    self.Os_clock = Ast_headTypeInfo
    
    self.Os_clock_sym = Ast_dummySymbol
    
    self.Os_date = Ast_headTypeInfo
    
    self.Os_date_sym = Ast_dummySymbol
    
    self.Os_difftime = Ast_headTypeInfo
    
    self.Os_difftime_sym = Ast_dummySymbol
    
    self.Os_exit = Ast_headTypeInfo
    
    self.Os_exit_sym = Ast_dummySymbol
    
    self.Os_remove = Ast_headTypeInfo
    
    self.Os_remove_sym = Ast_dummySymbol
    
    self.Os_rename = Ast_headTypeInfo
    
    self.Os_rename_sym = Ast_dummySymbol
    
    self.Os_time = Ast_headTypeInfo
    
    self.Os_time_sym = Ast_dummySymbol
    
    self.String_ = Ast_headTypeInfo
    
    self.String___attrib = Ast_headTypeInfo
    
    self.String___attrib_sym = Ast_dummySymbol
    
    self.String_byte = Ast_headTypeInfo
    
    self.String_byte_sym = Ast_dummySymbol
    
    self.String_dump = Ast_headTypeInfo
    
    self.String_dump_sym = Ast_dummySymbol
    
    self.String_find = Ast_headTypeInfo
    
    self.String_find_sym = Ast_dummySymbol
    
    self.String_format = Ast_headTypeInfo
    
    self.String_format_sym = Ast_dummySymbol
    
    self.String_gmatch = Ast_headTypeInfo
    
    self.String_gmatch_sym = Ast_dummySymbol
    
    self.String_gsub = Ast_headTypeInfo
    
    self.String_gsub_sym = Ast_dummySymbol
    
    self.String_lower = Ast_headTypeInfo
    
    self.String_lower_sym = Ast_dummySymbol
    
    self.String_rep = Ast_headTypeInfo
    
    self.String_rep_sym = Ast_dummySymbol
    
    self.String_reverse = Ast_headTypeInfo
    
    self.String_reverse_sym = Ast_dummySymbol
    
    self.String_sub = Ast_headTypeInfo
    
    self.String_sub_sym = Ast_dummySymbol
    
    self.String_upper = Ast_headTypeInfo
    
    self.String_upper_sym = Ast_dummySymbol
    
    self.Str_ = Ast_headTypeInfo
    
    self.Str___attrib = Ast_headTypeInfo
    
    self.Str___attrib_sym = Ast_dummySymbol
    
    self.Str_byte = Ast_headTypeInfo
    
    self.Str_byte_sym = Ast_dummySymbol
    
    self.Str_find = Ast_headTypeInfo
    
    self.Str_find_sym = Ast_dummySymbol
    
    self.Str_format = Ast_headTypeInfo
    
    self.Str_format_sym = Ast_dummySymbol
    
    self.Str_gmatch = Ast_headTypeInfo
    
    self.Str_gmatch_sym = Ast_dummySymbol
    
    self.Str_gsub = Ast_headTypeInfo
    
    self.Str_gsub_sym = Ast_dummySymbol
    
    self.Str_lower = Ast_headTypeInfo
    
    self.Str_lower_sym = Ast_dummySymbol
    
    self.Str_rep = Ast_headTypeInfo
    
    self.Str_rep_sym = Ast_dummySymbol
    
    self.Str_reverse = Ast_headTypeInfo
    
    self.Str_reverse_sym = Ast_dummySymbol
    
    self.Str_sub = Ast_headTypeInfo
    
    self.Str_sub_sym = Ast_dummySymbol
    
    self.Str_upper = Ast_headTypeInfo
    
    self.Str_upper_sym = Ast_dummySymbol
    
    self.List_ = Ast_headTypeInfo
    
    self.List___less = Ast_headTypeInfo
    
    self.List___less_sym = Ast_dummySymbol
    
    self.List_insert = Ast_headTypeInfo
    
    self.List_insert_sym = Ast_dummySymbol
    
    self.List_remove = Ast_headTypeInfo
    
    self.List_remove_sym = Ast_dummySymbol
    
    self.List_sort = Ast_headTypeInfo
    
    self.List_sort_sym = Ast_dummySymbol
    
    self.List_unpack = Ast_headTypeInfo
    
    self.List_unpack_sym = Ast_dummySymbol
    
    self.Array_ = Ast_headTypeInfo
    
    self.Array___less = Ast_headTypeInfo
    
    self.Array___less_sym = Ast_dummySymbol
    
    self.Array_sort = Ast_headTypeInfo
    
    self.Array_sort_sym = Ast_dummySymbol
    
    self.Array_unpack = Ast_headTypeInfo
    
    self.Array_unpack_sym = Ast_dummySymbol
    
    self.Set_ = Ast_headTypeInfo
    
    self.Set_add = Ast_headTypeInfo
    
    self.Set_add_sym = Ast_dummySymbol
    
    self.Set_and = Ast_headTypeInfo
    
    self.Set_and_sym = Ast_dummySymbol
    
    self.Set_clone = Ast_headTypeInfo
    
    self.Set_clone_sym = Ast_dummySymbol
    
    self.Set_del = Ast_headTypeInfo
    
    self.Set_del_sym = Ast_dummySymbol
    
    self.Set_has = Ast_headTypeInfo
    
    self.Set_has_sym = Ast_dummySymbol
    
    self.Set_len = Ast_headTypeInfo
    
    self.Set_len_sym = Ast_dummySymbol
    
    self.Set_or = Ast_headTypeInfo
    
    self.Set_or_sym = Ast_dummySymbol
    
    self.Set_sub = Ast_headTypeInfo
    
    self.Set_sub_sym = Ast_dummySymbol
    
    self.Math_ = Ast_headTypeInfo
    
    self.Math___attrib = Ast_headTypeInfo
    
    self.Math___attrib_sym = Ast_dummySymbol
    
    self.Math_random = Ast_headTypeInfo
    
    self.Math_random_sym = Ast_dummySymbol
    
    self.Math_randomseed = Ast_headTypeInfo
    
    self.Math_randomseed_sym = Ast_dummySymbol
    
    self.Debug_ = Ast_headTypeInfo
    
    self.Debug___attrib = Ast_headTypeInfo
    
    self.Debug___attrib_sym = Ast_dummySymbol
    
    self.Debug_getinfo = Ast_headTypeInfo
    
    self.Debug_getinfo_sym = Ast_dummySymbol
    
    self.Debug_getlocal = Ast_headTypeInfo
    
    self.Debug_getlocal_sym = Ast_dummySymbol
    
    self.Nilable_ = Ast_headTypeInfo
    
    self.Nilable_val = Ast_headTypeInfo
    
    self.Nilable_val_sym = Ast_dummySymbol
    
    self.allSymbol = NewLnsList([]LnsAny{})
    
    self.allClass = NewLnsList([]LnsAny{})
    
    self.allSymbolSet = NewLnsSet([]LnsAny{})
    
    self.allFuncTypeSet = NewLnsSet([]LnsAny{})
    
    self.needThreadingTypes = NewLnsSet([]LnsAny{})
    
}

// 1976: decl @lune.@base.@TransUnit.BuiltinFuncType.register
func (self *TransUnit_BuiltinFuncType) Register(symbolInfo *Ast_SymbolInfo) {
    self.allSymbol.Insert(Ast_SymbolInfo2Stem(symbolInfo))
    self.allSymbolSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
}

// 1980: decl @lune.@base.@TransUnit.BuiltinFuncType.registerClass
func (self *TransUnit_BuiltinFuncType) RegisterClass(classInfo *Ast_TypeInfo) {
    self.allClass.Insert(Ast_TypeInfo2Stem(classInfo))
    if _switch14697 := classInfo.FP.Get_rawTxt(); _switch14697 == "" {
        self.Lns_ = classInfo
        
    } else if _switch14697 == "iStream" {
        self.Istream_ = classInfo
        
    } else if _switch14697 == "oStream" {
        self.Ostream_ = classInfo
        
    } else if _switch14697 == "luaStream" {
        self.Luastream_ = classInfo
        
    } else if _switch14697 == "Mapping" {
        self.Mapping_ = classInfo
        
    } else if _switch14697 == "__pipe" {
        self.G__pipe_ = classInfo
        
    } else if _switch14697 == "LnsThread" {
        self.Lnsthread_ = classInfo
        
    } else if _switch14697 == "io" {
        self.Io_ = classInfo
        
    } else if _switch14697 == "package" {
        self.Package_ = classInfo
        
    } else if _switch14697 == "os" {
        self.Os_ = classInfo
        
    } else if _switch14697 == "string" {
        self.String_ = classInfo
        
    } else if _switch14697 == "str" {
        self.Str_ = classInfo
        
    } else if _switch14697 == "List" {
        self.List_ = classInfo
        
    } else if _switch14697 == "Array" {
        self.Array_ = classInfo
        
    } else if _switch14697 == "Set" {
        self.Set_ = classInfo
        
    } else if _switch14697 == "math" {
        self.Math_ = classInfo
        
    } else if _switch14697 == "debug" {
        self.Debug_ = classInfo
        
    } else if _switch14697 == "Nilable" {
        self.Nilable_ = classInfo
        
    }
}


// declaration Class -- DependModuleInfo
type TransUnit_DependModuleInfoMtd interface {
    GetTypeInfo(arg1 LnsInt) *Ast_TypeInfo
}
type TransUnit_DependModuleInfo struct {
    id LnsInt
    metaTypeId2TypeInfoMap *LnsMap
    FP TransUnit_DependModuleInfoMtd
}
func TransUnit_DependModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_DependModuleInfo).FP
}
type TransUnit_DependModuleInfoDownCast interface {
    ToTransUnit_DependModuleInfo() *TransUnit_DependModuleInfo
}
func TransUnit_DependModuleInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_DependModuleInfoDownCast)
    if ok { return work.ToTransUnit_DependModuleInfo() }
    return nil
}
func (obj *TransUnit_DependModuleInfo) ToTransUnit_DependModuleInfo() *TransUnit_DependModuleInfo {
    return obj
}
func NewTransUnit_DependModuleInfo(arg1 LnsInt, arg2 *LnsMap) *TransUnit_DependModuleInfo {
    obj := &TransUnit_DependModuleInfo{}
    obj.FP = obj
    obj.InitTransUnit_DependModuleInfo(arg1, arg2)
    return obj
}
func (self *TransUnit_DependModuleInfo) InitTransUnit_DependModuleInfo(arg1 LnsInt, arg2 *LnsMap) {
    self.id = arg1
    self.metaTypeId2TypeInfoMap = arg2
}
// 3105: decl @lune.@base.@TransUnit.DependModuleInfo.getTypeInfo
func (self *TransUnit_DependModuleInfo) GetTypeInfo(metaTypeId LnsInt) *Ast_TypeInfo {
    return Lns_unwrap( self.metaTypeId2TypeInfoMap.Items[metaTypeId]).(*Ast_TypeInfo)
}


// declaration Class -- ASTInfo
type TransUnit_ASTInfoMtd interface {
    Get_moduleSymbolKind() LnsInt
    Get_moduleTypeInfo() *Ast_TypeInfo
    Get_node() *Nodes_Node
    Get_processInfo() *Ast_ProcessInfo
}
type TransUnit_ASTInfo struct {
    node *Nodes_Node
    moduleTypeInfo *Ast_TypeInfo
    moduleSymbolKind LnsInt
    processInfo *Ast_ProcessInfo
    FP TransUnit_ASTInfoMtd
}
func TransUnit_ASTInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_ASTInfo).FP
}
type TransUnit_ASTInfoDownCast interface {
    ToTransUnit_ASTInfo() *TransUnit_ASTInfo
}
func TransUnit_ASTInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_ASTInfoDownCast)
    if ok { return work.ToTransUnit_ASTInfo() }
    return nil
}
func (obj *TransUnit_ASTInfo) ToTransUnit_ASTInfo() *TransUnit_ASTInfo {
    return obj
}
func NewTransUnit_ASTInfo(arg1 *Nodes_Node, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *Ast_ProcessInfo) *TransUnit_ASTInfo {
    obj := &TransUnit_ASTInfo{}
    obj.FP = obj
    obj.InitTransUnit_ASTInfo(arg1, arg2, arg3, arg4)
    return obj
}
func (self *TransUnit_ASTInfo) InitTransUnit_ASTInfo(arg1 *Nodes_Node, arg2 *Ast_TypeInfo, arg3 LnsInt, arg4 *Ast_ProcessInfo) {
    self.node = arg1
    self.moduleTypeInfo = arg2
    self.moduleSymbolKind = arg3
    self.processInfo = arg4
}
func (self *TransUnit_ASTInfo) Get_node() *Nodes_Node{ return self.node }
func (self *TransUnit_ASTInfo) Get_moduleTypeInfo() *Ast_TypeInfo{ return self.moduleTypeInfo }
func (self *TransUnit_ASTInfo) Get_moduleSymbolKind() LnsInt{ return self.moduleSymbolKind }
func (self *TransUnit_ASTInfo) Get_processInfo() *Ast_ProcessInfo{ return self.processInfo }

// declaration Class -- LetVarInfo
type TransUnit_LetVarInfoMtd interface {
}
type TransUnit_LetVarInfo struct {
    Mutable LnsInt
    VarName *Types_Token
    VarType LnsAny
    FP TransUnit_LetVarInfoMtd
}
func TransUnit_LetVarInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_LetVarInfo).FP
}
type TransUnit_LetVarInfoDownCast interface {
    ToTransUnit_LetVarInfo() *TransUnit_LetVarInfo
}
func TransUnit_LetVarInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_LetVarInfoDownCast)
    if ok { return work.ToTransUnit_LetVarInfo() }
    return nil
}
func (obj *TransUnit_LetVarInfo) ToTransUnit_LetVarInfo() *TransUnit_LetVarInfo {
    return obj
}
func NewTransUnit_LetVarInfo(arg1 LnsInt, arg2 *Types_Token, arg3 LnsAny) *TransUnit_LetVarInfo {
    obj := &TransUnit_LetVarInfo{}
    obj.FP = obj
    obj.InitTransUnit_LetVarInfo(arg1, arg2, arg3)
    return obj
}
func (self *TransUnit_LetVarInfo) InitTransUnit_LetVarInfo(arg1 LnsInt, arg2 *Types_Token, arg3 LnsAny) {
    self.Mutable = arg1
    self.VarName = arg2
    self.VarType = arg3
}

func Lns_TransUnit_init() {
    if init_TransUnit { return }
    init_TransUnit = true
    TransUnit__mod__ = "@lune.@base.@TransUnit"
    Lns_InitMod()
    Lns_Types_init()
    Lns_Meta_init()
    Lns_Parser_init()
    Lns_Util_init()
    Lns_Ast_init()
    Lns_Nodes_init()
    Lns_Writer_init()
    Lns_frontInterface_init()
    Lns_LuaVer_init()
    Lns_Option_init()
    Lns_Code_init()
    Lns_Log_init()
    Lns_LuneControl_init()
    Lns_Macro_init()
    TransUnit_op2levelMap = NewLnsMap( map[LnsAny]LnsAny{})
    TransUnit_op1levelMap = NewLnsMap( map[LnsAny]LnsAny{})
    _ = 0
    {
        var opLevelBase LnsInt
        opLevelBase = 0
        var regOpLevel func(opnum LnsInt,opList *LnsList)
        regOpLevel = func(opnum LnsInt,opList *LnsList) {
            opLevelBase = opLevelBase + 1
            
            if opnum == 1{
                for _, _op := range( opList.Items ) {
                    op := _op.(string)
                    TransUnit_op1levelMap.Set(op,opLevelBase)
                }
            } else { 
                for _, _op := range( opList.Items ) {
                    op := _op.(string)
                    TransUnit_op2levelMap.Set(op,opLevelBase)
                }
            }
        }
        regOpLevel(2, NewLnsList([]LnsAny{"="}))
        regOpLevel(2, NewLnsList([]LnsAny{"or"}))
        regOpLevel(2, NewLnsList([]LnsAny{"and"}))
        regOpLevel(2, NewLnsList([]LnsAny{"<", ">", "<=", ">=", "~=", "=="}))
        regOpLevel(2, NewLnsList([]LnsAny{"|"}))
        regOpLevel(2, NewLnsList([]LnsAny{"~"}))
        regOpLevel(2, NewLnsList([]LnsAny{"&"}))
        regOpLevel(2, NewLnsList([]LnsAny{"|<<", "|>>"}))
        regOpLevel(2, NewLnsList([]LnsAny{".."}))
        regOpLevel(2, NewLnsList([]LnsAny{"+", "-"}))
        regOpLevel(2, NewLnsList([]LnsAny{"*", "/", "//", "%"}))
        regOpLevel(1, NewLnsList([]LnsAny{"`", ",,", ",,,", ",,,,"}))
        regOpLevel(1, NewLnsList([]LnsAny{"not", "#", "-", "~"}))
        regOpLevel(1, NewLnsList([]LnsAny{"^"}))
        TransUnit_opTopLevel = opLevelBase + 1
        
    }
    TransUnit_quotedChar2Code = NewLnsMap( map[LnsAny]LnsAny{})
    TransUnit_quotedChar2Code.Set("a",7)
    TransUnit_quotedChar2Code.Set("b",8)
    TransUnit_quotedChar2Code.Set("t",9)
    TransUnit_quotedChar2Code.Set("n",10)
    TransUnit_quotedChar2Code.Set("v",11)
    TransUnit_quotedChar2Code.Set("f",12)
    TransUnit_quotedChar2Code.Set("r",13)
    TransUnit_quotedChar2Code.Set("\\",LnsInt(92))
    TransUnit_quotedChar2Code.Set("\"",LnsInt(34))
    TransUnit_quotedChar2Code.Set("'",LnsInt(39))
    TransUnit_builtinFunc = NewTransUnit_BuiltinFuncType()
    
    TransUnit_readyBuiltin = false
    TransUnit_specialSymbolSet = NewLnsSet([]LnsAny{"__init", "__free", "__", "_exp"})
    TransUnit_builtinKeywordSet = NewLnsSet([]LnsAny{"self", "super"})
    TransUnit_CantOverrideMethods = NewLnsSet([]LnsAny{"__init", "__free"})
}
func init() {
    init_TransUnit = false
}
