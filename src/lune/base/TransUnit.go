// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_TransUnit bool
var TransUnit__mod__ string
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
func TransUnit_DeclFuncMode_get__allList_2_(_env *LnsEnv) *LnsList{
    return TransUnit_DeclFuncModeList_
}
var TransUnit_DeclFuncModeMap_ = map[LnsInt]string {
  TransUnit_DeclFuncMode__Class: "DeclFuncMode.Class",
  TransUnit_DeclFuncMode__Func: "DeclFuncMode.Func",
  TransUnit_DeclFuncMode__Glue: "DeclFuncMode.Glue",
  TransUnit_DeclFuncMode__Module: "DeclFuncMode.Module",
}
func TransUnit_DeclFuncMode__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
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
func TransUnit_ExpSymbolMode_get__allList_2_(_env *LnsEnv) *LnsList{
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
func TransUnit_ExpSymbolMode__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
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
func TransUnit_AnalyzeMode_get__allList(_env *LnsEnv) *LnsList{
    return TransUnit_AnalyzeModeList_
}
var TransUnit_AnalyzeModeMap_ = map[LnsInt]string {
  TransUnit_AnalyzeMode__Compile: "AnalyzeMode.Compile",
  TransUnit_AnalyzeMode__Complete: "AnalyzeMode.Complete",
  TransUnit_AnalyzeMode__Diag: "AnalyzeMode.Diag",
  TransUnit_AnalyzeMode__Inquire: "AnalyzeMode.Inquire",
}
func TransUnit_AnalyzeMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_AnalyzeModeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_AnalyzeMode_getTxt(arg1 LnsInt) string {
    return TransUnit_AnalyzeModeMap_[arg1];
}
// decl enum -- DefaultAsyncMode 
type TransUnit_DefaultAsyncMode = LnsInt
const TransUnit_DefaultAsyncMode__AsyncAll = 2
const TransUnit_DefaultAsyncMode__AsyncFunc = 1
const TransUnit_DefaultAsyncMode__NoAsync = 0
var TransUnit_DefaultAsyncModeList_ = NewLnsList( []LnsAny {
  TransUnit_DefaultAsyncMode__NoAsync,
  TransUnit_DefaultAsyncMode__AsyncFunc,
  TransUnit_DefaultAsyncMode__AsyncAll,
})
func TransUnit_DefaultAsyncMode_get__allList_2_(_env *LnsEnv) *LnsList{
    return TransUnit_DefaultAsyncModeList_
}
var TransUnit_DefaultAsyncModeMap_ = map[LnsInt]string {
  TransUnit_DefaultAsyncMode__AsyncAll: "DefaultAsyncMode.AsyncAll",
  TransUnit_DefaultAsyncMode__AsyncFunc: "DefaultAsyncMode.AsyncFunc",
  TransUnit_DefaultAsyncMode__NoAsync: "DefaultAsyncMode.NoAsync",
}
func TransUnit_DefaultAsyncMode__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_DefaultAsyncModeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_DefaultAsyncMode_getTxt(arg1 LnsInt) string {
    return TransUnit_DefaultAsyncModeMap_[arg1];
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
func TransUnit_AnalyzingState_get__allList_2_(_env *LnsEnv) *LnsList{
    return TransUnit_AnalyzingStateList_
}
var TransUnit_AnalyzingStateMap_ = map[LnsInt]string {
  TransUnit_AnalyzingState__ClassMethod: "AnalyzingState.ClassMethod",
  TransUnit_AnalyzingState__Constructor: "AnalyzingState.Constructor",
  TransUnit_AnalyzingState__Func: "AnalyzingState.Func",
  TransUnit_AnalyzingState__InitBlock: "AnalyzingState.InitBlock",
  TransUnit_AnalyzingState__Other: "AnalyzingState.Other",
}
func TransUnit_AnalyzingState__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_AnalyzingStateMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_AnalyzingState_getTxt(arg1 LnsInt) string {
    return TransUnit_AnalyzingStateMap_[arg1];
}
// decl enum -- AnalyzePhase 
type TransUnit_AnalyzePhase = LnsInt
const TransUnit_AnalyzePhase__Main = 2
const TransUnit_AnalyzePhase__Meta = 0
const TransUnit_AnalyzePhase__Runner = 1
var TransUnit_AnalyzePhaseList_ = NewLnsList( []LnsAny {
  TransUnit_AnalyzePhase__Meta,
  TransUnit_AnalyzePhase__Runner,
  TransUnit_AnalyzePhase__Main,
})
func TransUnit_AnalyzePhase_get__allList_2_(_env *LnsEnv) *LnsList{
    return TransUnit_AnalyzePhaseList_
}
var TransUnit_AnalyzePhaseMap_ = map[LnsInt]string {
  TransUnit_AnalyzePhase__Main: "AnalyzePhase.Main",
  TransUnit_AnalyzePhase__Meta: "AnalyzePhase.Meta",
  TransUnit_AnalyzePhase__Runner: "AnalyzePhase.Runner",
}
func TransUnit_AnalyzePhase__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_AnalyzePhaseMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_AnalyzePhase_getTxt(arg1 LnsInt) string {
    return TransUnit_AnalyzePhaseMap_[arg1];
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
func TransUnit_SymbolMode_get__allList_2_(_env *LnsEnv) *LnsList{
    return TransUnit_SymbolModeList_
}
var TransUnit_SymbolModeMap_ = map[LnsInt]string {
  TransUnit_SymbolMode__MustNot_: "SymbolMode.MustNot_",
  TransUnit_SymbolMode__MustNot_Or_: "SymbolMode.MustNot_Or_",
  TransUnit_SymbolMode__Must_: "SymbolMode.Must_",
}
func TransUnit_SymbolMode__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
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
func TransUnit_TentativeMode_get__allList_2_(_env *LnsEnv) *LnsList{
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
func TransUnit_TentativeMode__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
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
func TransUnit_FormType_get__allList(_env *LnsEnv) *LnsList{
    return TransUnit_FormTypeList_
}
var TransUnit_FormTypeMap_ = map[LnsInt]string {
  TransUnit_FormType__Match: "FormType.Match",
  TransUnit_FormType__NeedConv: "FormType.NeedConv",
  TransUnit_FormType__Unmatch: "FormType.Unmatch",
}
func TransUnit_FormType__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := TransUnit_FormTypeMap_[arg1]; ok { return arg1 }
    return nil
}

func TransUnit_FormType_getTxt(arg1 LnsInt) string {
    return TransUnit_FormTypeMap_[arg1];
}
var TransUnit_opTopLevel LnsInt
var TransUnit_op2levelMap *LnsMap
var TransUnit_op1levelMap *LnsMap
var TransUnit_quotedChar2Code *LnsMap
var TransUnit_specialSymbolSet *LnsSet
var TransUnit_builtinKeywordSet *LnsSet
var TransUnit_CantOverrideMethods *LnsSet
var TransUnit_unsupportStatement *LnsSet
type TransUnit_checkImplicitCastCallback_10_ func (_env *LnsEnv, arg1 *Ast_TypeInfo,arg2 *Nodes_Node) LnsAny
type TransUnit_ReadyExportInfo func (_env *LnsEnv, arg1 *FrontInterface_ExportInfo)
type TransUnit_checkCompForm_31_ func (_env *LnsEnv, arg1 *Writer_JSON,arg2 string)
// for 2822
func TransUnit_convExp0_5096(arg1 []LnsAny) (LnsAny, LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 )
}
// for 3468
func TransUnit_convExp0_8416(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 57
func TransUnit_convExp0_13806(arg1 []LnsAny) (LnsAny, LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 )
}
// for 2150
func TransUnit_convExp1_6721(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2354
func TransUnit_convExp1_7795(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 368
func TransUnit_convExp2_1550(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 408
func TransUnit_convExp3_10766(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1217
func TransUnit_convExp4_3546(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1213
func TransUnit_convExp4_3548(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1383
func TransUnit_convExp4_4186(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 3540
func TransUnit_convExp0_8804(arg1 []LnsAny) (*Nodes_Node, *Nodes_Node) {
    return Lns_getFromMulti( arg1, 0 ).(*Nodes_Node), Lns_getFromMulti( arg1, 1 ).(*Nodes_Node)
}
// for 4059
func TransUnit_convExp0_11713(arg1 []LnsAny) *Ast_TypeInfo {
    return &Lns_getFromMulti( arg1, 0 ).(*Ast_GenericTypeInfo).Ast_TypeInfo
}
// for 781
func TransUnit_convExp0_18367(arg1 []LnsAny) (LnsAny, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// for 253
func TransUnit_convExp2_878(arg1 []LnsAny) *Ast_TypeInfo {
    return &Lns_getFromMulti( arg1, 0 ).(*Ast_GenericTypeInfo).Ast_TypeInfo
}
// for 970
func TransUnit_convExp2_4136(arg1 []LnsAny) (*Types_Token, LnsAny, LnsAny, *LnsMap, *Nodes_ClassInheritInfo) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 ), Lns_getFromMulti( arg1, 3 ).(*LnsMap), Lns_getFromMulti( arg1, 4 ).(*Nodes_ClassInheritInfo)
}
// for 1041
func TransUnit_convExp2_4528(arg1 []LnsAny) (*Types_Token, LnsAny, *LnsList) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 ).(*LnsList)
}
// for 1078
func TransUnit_convExp2_4718(arg1 []LnsAny) (*Types_Token, *LnsList) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(*LnsList)
}
// for 1098
func TransUnit_convExp2_4820(arg1 []LnsAny) (*Types_Token, *TransUnitIF_NSInfo, *Nodes_ClassInheritInfo) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(*TransUnitIF_NSInfo), Lns_getFromMulti( arg1, 2 ).(*Nodes_ClassInheritInfo)
}
// for 1510
func TransUnit_convExp2_7009(arg1 []LnsAny) (*Types_Token, LnsInt, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(LnsInt), Lns_getFromMulti( arg1, 2 )
}
// for 1515
func TransUnit_convExp2_7059(arg1 []LnsAny) (*LnsList, *Types_Token, *LnsList) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsList), Lns_getFromMulti( arg1, 1 ).(*Types_Token), Lns_getFromMulti( arg1, 2 ).(*LnsList)
}
// for 1747
func TransUnit_convExp2_8283(arg1 []LnsAny) (LnsInt, *Ast_TypeInfo, LnsAny, *Types_Token) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*Ast_TypeInfo), Lns_getFromMulti( arg1, 2 ), Lns_getFromMulti( arg1, 3 ).(*Types_Token)
}
// for 1762
func TransUnit_convExp2_8379(arg1 []LnsAny) (LnsInt, *Ast_TypeInfo, LnsAny, *Types_Token) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*Ast_TypeInfo), Lns_getFromMulti( arg1, 2 ), Lns_getFromMulti( arg1, 3 ).(*Types_Token)
}
// for 2508
func TransUnit_convExp3_207(arg1 []LnsAny) (*Types_Token, *LnsList) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(*LnsList)
}
// for 2863
func TransUnit_convExp3_2062(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2867
func TransUnit_convExp3_2102(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 3009
func TransUnit_convExp3_2645(arg1 []LnsAny) (*Types_Token, *LnsList) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(*LnsList)
}
// for 3037
func TransUnit_convExp3_2791(arg1 []LnsAny) (*Types_Token, LnsInt, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(LnsInt), Lns_getFromMulti( arg1, 2 )
}
// for 3099
func TransUnit_convExp3_3115(arg1 []LnsAny) (*LnsList, *Types_Token, *LnsList) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsList), Lns_getFromMulti( arg1, 1 ).(*Types_Token), Lns_getFromMulti( arg1, 2 ).(*LnsList)
}
// for 3138
func TransUnit_convExp3_3284(arg1 []LnsAny) (*Ast_SymbolInfo, *TransUnitIF_NSInfo) {
    return Lns_getFromMulti( arg1, 0 ).(*Ast_SymbolInfo), Lns_getFromMulti( arg1, 1 ).(*TransUnitIF_NSInfo)
}
// for 4085
func TransUnit_convExp3_7879(arg1 []LnsAny) (*LnsList, *LnsList, *LnsList, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsList), Lns_getFromMulti( arg1, 1 ).(*LnsList), Lns_getFromMulti( arg1, 2 ).(*LnsList), Lns_getFromMulti( arg1, 3 )
}
// for 4094
func TransUnit_convExp3_7967(arg1 []LnsAny) (*LnsList, *LnsList, *LnsList, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsList), Lns_getFromMulti( arg1, 1 ).(*LnsList), Lns_getFromMulti( arg1, 2 ).(*LnsList), Lns_getFromMulti( arg1, 3 )
}
// for 548
func TransUnit_convExp4_568(arg1 []LnsAny) (*Nodes_Node, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Nodes_Node), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 556
func TransUnit_convExp4_620(arg1 []LnsAny) (*Nodes_Node, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Nodes_Node), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 1932
func TransUnit_convExp4_6740(arg1 []LnsAny) (*Nodes_Node, *Types_Token) {
    return Lns_getFromMulti( arg1, 0 ).(*Nodes_Node), Lns_getFromMulti( arg1, 1 ).(*Types_Token)
}
// for 2228
func TransUnit_convExp4_8235(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2519
func TransUnit_convExp4_9535(arg1 []LnsAny) (*Ast_TypeInfo, LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Ast_TypeInfo), Lns_getFromMulti( arg1, 1 ), Lns_getFromMulti( arg1, 2 ).(bool)
}
// for 836
func TransUnit_convExp0_878(_env *LnsEnv, arg1 []LnsAny) (*LnsEnv, *Types_TransCtrlInfo, *Ast_ProcessInfo) {
    return _env, Lns_getFromMulti( arg1, 0 ).(*Types_TransCtrlInfo), Lns_getFromMulti( arg1, 1 ).(*Ast_ProcessInfo)
}
// for 1262
func TransUnit_convExp0_1977(arg1 []LnsAny) (*LnsMap, *LnsMap) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsMap), Lns_getFromMulti( arg1, 1 ).(*LnsMap)
}
// for 3094
func TransUnit_convExp0_6442(arg1 []LnsAny) (LnsInt, *LnsMap, LnsAny, *LnsList) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*LnsMap), Lns_getFromMulti( arg1, 2 ), Lns_getFromMulti( arg1, 3 ).(*LnsList)
}
// for 3760
func TransUnit_convExp0_10138(arg1 []LnsAny) (*Types_Token, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 3876
func TransUnit_convExp0_10695(arg1 []LnsAny) (LnsInt, *LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*LnsMap), Lns_getFromMulti( arg1, 2 )
}
// for 4038
func TransUnit_convExp0_11601(arg1 []LnsAny) (LnsInt, *LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*LnsMap), Lns_getFromMulti( arg1, 2 )
}
// for 4238
func TransUnit_convExp0_12784(arg1 []LnsAny) (*Nodes_Node, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Nodes_Node), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 236
func TransUnit_convExp0_15583(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 262
func TransUnit_convExp0_15722(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2004
func TransUnit_convExp1_5818(arg1 []LnsAny) (*Types_Token, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 2071
func TransUnit_convExp1_6184(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2218
func TransUnit_convExp1_6954(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 2333
func TransUnit_convExp1_7679(arg1 []LnsAny) (LnsAny, bool) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 2572
func TransUnit_convExp1_8995(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 136
func TransUnit_convExp2_178(arg1 []LnsAny) (*Types_Token, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 331
func TransUnit_convExp2_1258(arg1 []LnsAny) (*Types_Token, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 438
func TransUnit_convExp2_1883(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 813
func TransUnit_convExp2_3301(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 910
func TransUnit_convExp2_3849(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 1045
func TransUnit_convExp2_4557(arg1 []LnsAny) *Ast_AlternateTypeInfo {
    return Lns_getFromMulti( arg1, 0 ).(*Ast_AlternateTypeInfo)
}
// for 1158
func TransUnit_convExp2_5120(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1241
func TransUnit_convExp2_5570(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1270
func TransUnit_convExp2_5720(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1324
func TransUnit_convExp2_6067(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1382
func TransUnit_convExp2_6424(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1538
func TransUnit_convExp2_7186(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1809
func TransUnit_convExp2_8715(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2243
func TransUnit_convExp2_10613(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2412
func TransUnit_convExp2_11628(arg1 []LnsAny) (*LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsMap), Lns_getFromMulti( arg1, 1 )
}
// for 2542
func TransUnit_convExp3_408(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2570
func TransUnit_convExp3_563(arg1 []LnsAny) (*Types_Token, *TransUnitIF_NSInfo, *Nodes_ClassInheritInfo) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(*TransUnitIF_NSInfo), Lns_getFromMulti( arg1, 2 ).(*Nodes_ClassInheritInfo)
}
// for 2626
func TransUnit_convExp3_882(arg1 []LnsAny) (*Nodes_DeclClassNode, *Types_Token, *LnsSet) {
    return Lns_getFromMulti( arg1, 0 ).(*Nodes_DeclClassNode), Lns_getFromMulti( arg1, 1 ).(*Types_Token), Lns_getFromMulti( arg1, 2 ).(*LnsSet)
}
// for 2731
func TransUnit_convExp3_1381(arg1 []LnsAny) *Ast_GenericTypeInfo {
    return Lns_getFromMulti( arg1, 0 ).(*Ast_GenericTypeInfo)
}
// for 2778
func TransUnit_convExp3_1586(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 2787
func TransUnit_convExp3_1639(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 2796
func TransUnit_convExp3_1686(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 3610
func TransUnit_convExp3_5465(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 3838
func TransUnit_convExp3_6569(arg1 []LnsAny) (*Types_Token, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 3856
func TransUnit_convExp3_6675(arg1 []LnsAny) (*LnsList, *LnsList, *LnsList, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsList), Lns_getFromMulti( arg1, 1 ).(*LnsList), Lns_getFromMulti( arg1, 2 ).(*LnsList), Lns_getFromMulti( arg1, 3 )
}
// for 3892
func TransUnit_convExp3_6912(arg1 []LnsAny) *Ast_SymbolInfo {
    return Lns_getFromMulti( arg1, 0 ).(*Ast_SymbolInfo)
}
// for 4180
func TransUnit_convExp3_8350(arg1 []LnsAny) (*Types_Token, bool) {
    return Lns_getFromMulti( arg1, 0 ).(*Types_Token), Lns_getFromMulti( arg1, 1 ).(bool)
}
// for 678
func TransUnit_convExp4_1063(arg1 []LnsAny) (LnsInt, string) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 786
func TransUnit_convExp4_1474(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 856
func TransUnit_convExp4_1837(arg1 []LnsAny) (LnsInt, *LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*LnsMap), Lns_getFromMulti( arg1, 2 )
}
// for 925
func TransUnit_convExp4_2161(arg1 []LnsAny) (LnsInt, *LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*LnsMap), Lns_getFromMulti( arg1, 2 )
}
// for 947
func TransUnit_convExp4_2247(arg1 []LnsAny) (LnsAny, *Ast_TypeInfo) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(*Ast_TypeInfo)
}
// for 983
func TransUnit_convExp4_2453(arg1 []LnsAny) (LnsAny, *Ast_TypeInfo) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(*Ast_TypeInfo)
}
// for 1100
func TransUnit_convExp4_3009(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 1292
func TransUnit_convExp4_3819(arg1 []LnsAny) (LnsInt, *Ast_TypeInfo) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*Ast_TypeInfo)
}
// for 1355
func TransUnit_convExp4_4063(arg1 []LnsAny) (LnsInt, *LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*LnsMap), Lns_getFromMulti( arg1, 2 )
}
// for 1643
func TransUnit_convExp4_5343(arg1 []LnsAny) (LnsAny, string) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 ).(string)
}
// for 1741
func TransUnit_convExp4_5689(arg1 []LnsAny) (*LnsMap, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(*LnsMap), Lns_getFromMulti( arg1, 1 )
}
// for 2227
func TransUnit_convExp4_8220(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2229
func TransUnit_convExp4_8250(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2233
func TransUnit_convExp4_8290(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
type TransUnit_DeclClassMode = TransUnitIF_DeclClassMode
// 110: decl @lune.@base.@TransUnit.clearThePosForModToRef
func TransUnit_clearThePosForModToRef_7_(_env *LnsEnv, scope *Ast_Scope,moduleScope *Ast_Scope) *LnsList {
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    scope.FP.FilterSymbolTypeInfo(_env, scope, moduleScope, Ast_ScopeAccess__Normal, Ast_filterForm(func(_env *LnsEnv, symInfo *Ast_SymbolInfo) bool {
        if symInfo.FP.Get_kind(_env) == Ast_SymbolKind__Var{
            list.Insert(TransUnit_RefAccessSymPos2Stem(NewTransUnit_RefAccessSymPos(_env, symInfo, symInfo.FP.Get_posForModToRef(_env))))
            symInfo.FP.Set_posForModToRef(_env, nil)
        }
        return true
    }))
    return list
}

// 1224: decl @lune.@base.@TransUnit.setupOpLevel
func TransUnit_setupOpLevel_21_(_env *LnsEnv)(*LnsMap, *LnsMap) {
    var op2levelMapWork *LnsMap
    op2levelMapWork = NewLnsMap( map[LnsAny]LnsAny{})
    var op1levelMapWork *LnsMap
    op1levelMapWork = NewLnsMap( map[LnsAny]LnsAny{})
    var opLevelBase LnsInt
    opLevelBase = 0
    var regOpLevel func(_env *LnsEnv, opnum LnsInt,opList *LnsList)
    regOpLevel = func(_env *LnsEnv, opnum LnsInt,opList *LnsList) {
        opLevelBase = opLevelBase + 1
        if opnum == 1{
            for _, _op := range( opList.Items ) {
                op := _op.(string)
                op1levelMapWork.Set(op,opLevelBase)
            }
        } else { 
            for _, _op := range( opList.Items ) {
                op := _op.(string)
                op2levelMapWork.Set(op,opLevelBase)
            }
        }
    }
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"="}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"or"}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"and"}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"<", ">", "<=", ">=", "~=", "=="}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"|"}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"~"}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"&"}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"|<<", "|>>"}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{".."}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"+", "-"}))
    regOpLevel(_env, 2, NewLnsList([]LnsAny{"*", "/", "//", "%"}))
    regOpLevel(_env, 1, NewLnsList([]LnsAny{"`", ",,", ",,,", ",,,,"}))
    regOpLevel(_env, 1, NewLnsList([]LnsAny{"not", "#", "-", "~"}))
    regOpLevel(_env, 1, NewLnsList([]LnsAny{"^"}))
    TransUnit_opTopLevel = opLevelBase + 1
    return op2levelMapWork, op1levelMapWork
}

// 536: decl @lune.@base.@TransUnit.getAnalyzingState
func TransUnit_getAnalyzingState_24_(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_rawTxt(_env) == "__init") &&
        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) ).(bool)){
        return TransUnit_AnalyzingState__Constructor
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_staticFlag(_env)) &&
        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) ).(bool)){
        return TransUnit_AnalyzingState__ClassMethod
    } else { 
        return TransUnit_AnalyzingState__Func
    }
// insert a dummy
    return 0
}

// 548: decl @lune.@base.@TransUnit.getFirstStmt
func TransUnit_getFirstStmt_25_(_env *LnsEnv, stmtList *LnsList) LnsAny {
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        if _switch0 := stmt.FP.Get_kind(_env); _switch0 == Nodes_nodeKindEnum__BlankLine {
        } else {
            return stmt
        }
    }
    return nil
}

// 1206: decl @lune.@base.@TransUnit.findForm
func TransUnit_findForm(_env *LnsEnv, format string) *LnsList {
    var remain string
    remain = _env.GetVM().String_replace(format,"%%", "")
    var opList *LnsList
    opList = NewLnsList([]LnsAny{})
    for  {
        var pos LnsAny
        var endPos LnsAny
        pos,endPos = nil, nil
        {
            _index, _endIndex := TransUnit_convExp4_3548(Lns_2DDD(_env.GetVM().String_find(remain,"^%%%-?[%d]*%a", nil, nil)))
            if !Lns_IsNil( _index ) && !Lns_IsNil( _endIndex ) {
                index := _index.(LnsInt)
                endIndex := _endIndex.(LnsInt)
                pos, endPos = index, endIndex
            } else {
                {
                    _index, _endIndex := TransUnit_convExp4_3546(Lns_2DDD(_env.GetVM().String_find(remain,"[^%%]%%%-?[%d]*%a", nil, nil)))
                    if !Lns_IsNil( _index ) && !Lns_IsNil( _endIndex ) {
                        index := _index.(LnsInt)
                        endIndex := _endIndex.(LnsInt)
                        pos, endPos = index + 1, endIndex
                    }
                }
            }
        }
        if pos != nil && endPos != nil{
            pos_281 := pos.(LnsInt)
            endPos_282 := endPos.(LnsInt)
            var op string
            op = _env.GetVM().String_sub(remain,pos_281, endPos_282)
            opList.Insert(op)
            remain = _env.GetVM().String_sub(remain,endPos_282 + 1, nil)
        } else {
            break
        }
    }
    return opList
}

// 1239: decl @lune.@base.@TransUnit.isMatchStringFormatType
func TransUnit_isMatchStringFormatType(_env *LnsEnv, opKind string,argType *Ast_TypeInfo,luaVer *LuaVer_LuaVerInfo)(LnsInt, *Ast_TypeInfo) {
    {
        _enumType := Ast_EnumTypeInfoDownCastF(argType.FP.Get_srcTypeInfo(_env).FP.Get_aliasSrc(_env).FP)
        if !Lns_IsNil( _enumType ) {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            argType = enumType.FP.Get_valTypeInfo(_env)
        }
    }
    if _switch0 := LnsInt(opKind[len(opKind)-1]); _switch0 == 115 {
        if argType.FP.Get_srcTypeInfo(_env) != Ast_builtinTypeString{
            if Lns_op_not(luaVer.FP.Get_canFormStem2Str(_env)){
                return TransUnit_FormType__NeedConv, Ast_builtinTypeString
            }
        }
    } else if _switch0 == 113 {
        if argType.FP.Get_srcTypeInfo(_env) != Ast_builtinTypeString{
            return TransUnit_FormType__Unmatch, Ast_builtinTypeString
        }
    } else if _switch0 == 65 || _switch0 == 97 || _switch0 == 69 || _switch0 == 101 || _switch0 == 102 || _switch0 == 71 || _switch0 == 103 {
        if argType.FP.Get_srcTypeInfo(_env) != Ast_builtinTypeReal{
            return TransUnit_FormType__Unmatch, Ast_builtinTypeReal
        }
    } else {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( argType.FP.Get_srcTypeInfo(_env) != Ast_builtinTypeInt) &&
            _env.SetStackVal( argType.FP.Get_srcTypeInfo(_env) != Ast_builtinTypeChar) ).(bool)){
            return TransUnit_FormType__Unmatch, Ast_builtinTypeInt
        }
    }
    return TransUnit_FormType__Match, Ast_builtinTypeNone
}

func TransUnit_analyzeExpSymbol___anonymous_0_(_env *LnsEnv, workSymbolInfo *Ast_SymbolInfo) bool {
    Lns_print([]LnsAny{"sym", workSymbolInfo.FP.Get_name(_env)})
    return true
}
// 3046: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOpSet.process
func TransUnit_analyzeExpOpSet__process_0_(_env *LnsEnv, lValNode *Nodes_Node) LnsAny {
    var refItemNode *Nodes_ExpRefItemNode
    
    {
        _refItemNode := Nodes_ExpRefItemNodeDownCastF(lValNode.FP)
        if _refItemNode == nil{
            return nil
        } else {
            refItemNode = _refItemNode.(*Nodes_ExpRefItemNode)
        }
    }
    if _switch0 := refItemNode.FP.Get_val(_env).FP.Get_expType(_env).FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Map {
        return refItemNode
    }
    return nil
}

// 3222: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOpEquals.getType
func TransUnit_analyzeExpOpEquals__getType_0_(_env *LnsEnv, typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    var workType *Ast_TypeInfo
    workType = typeInfo.FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( workType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
        _env.SetStackVal( workType.FP.HasBase(_env)) ).(bool)){
        return workType.FP.Get_baseTypeInfo(_env)
    }
    return workType
}







// 885: decl @lune.@base.@TransUnit.TransUnitCtrl.createAST.createId2proto
func TransUnitCtrl_createAST__createId2proto_2_(_env *LnsEnv, _map *LnsMap) *LnsMap {
    var id2proto *LnsMap
    id2proto = NewLnsMap( map[LnsAny]LnsAny{})
    for _protoType, _nsInfo := range( _map.Items ) {
        protoType := _protoType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        nsInfo := _nsInfo.(TransUnitIF_NSInfoDownCast).ToTransUnitIF_NSInfo()
        id2proto.Set(protoType.FP.Get_typeId(_env).Id,nsInfo)
    }
    return id2proto
}


// 778: decl @lune.@base.@TransUnit.TransUnit.mergeFrom.cloneMessList
func TransUnit_mergeFrom__cloneMessList_0_(_env *LnsEnv, dstList *LnsList,srcList *LnsList) {
    var set *LnsSet
    set = NewLnsSet([]LnsAny{})
    for _, _info := range( dstList.Items ) {
        info := _info.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
        set.Add(info.Mess)
    }
    for _, _info := range( srcList.Items ) {
        info := _info.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
        if Lns_op_not(set.Has(info.Mess)){
            dstList.Insert(TransUnitIF_ErrMess2Stem(info))
        }
    }
}


// 1504: decl @lune.@base.@TransUnit.TransUnit.getDefaultAsync.process
func TransUnit_getDefaultAsync__process_0_(_env *LnsEnv, defaultAsyncMode LnsInt) LnsInt {
    if _switch0 := defaultAsyncMode; _switch0 == TransUnit_DefaultAsyncMode__AsyncAll || _switch0 == TransUnit_DefaultAsyncMode__AsyncFunc {
        return Ast_Async__Async
    } else if _switch0 == TransUnit_DefaultAsyncMode__NoAsync {
        return Ast_Async__Noasync
    }
// insert a dummy
    return 0
}
























func TransUnit_analyzeInitExp___anonymous_0_(_env *LnsEnv, dstType *Ast_TypeInfo,expNode *Nodes_Node) LnsAny {
    return nil
}











// declaration Class -- AccessSymPos
type TransUnit_AccessSymPosMtd interface {
    Get_pos(_env *LnsEnv) Types_Position
    Get_symbol(_env *LnsEnv) *Ast_SymbolInfo
}
type TransUnit_AccessSymPos struct {
    symbol *Ast_SymbolInfo
    pos Types_Position
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
func NewTransUnit_AccessSymPos(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 Types_Position) *TransUnit_AccessSymPos {
    obj := &TransUnit_AccessSymPos{}
    obj.FP = obj
    obj.InitTransUnit_AccessSymPos(_env, arg1, arg2)
    return obj
}
func (self *TransUnit_AccessSymPos) InitTransUnit_AccessSymPos(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 Types_Position) {
    self.symbol = arg1
    self.pos = arg2
}
func (self *TransUnit_AccessSymPos) Get_symbol(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbol }
func (self *TransUnit_AccessSymPos) Get_pos(_env *LnsEnv) Types_Position{ return self.pos }

// declaration Class -- RefAccessSymPos
type TransUnit_RefAccessSymPosMtd interface {
    Get_pos(_env *LnsEnv) LnsAny
    Get_symbol(_env *LnsEnv) *Ast_SymbolInfo
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
func NewTransUnit_RefAccessSymPos(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 LnsAny) *TransUnit_RefAccessSymPos {
    obj := &TransUnit_RefAccessSymPos{}
    obj.FP = obj
    obj.InitTransUnit_RefAccessSymPos(_env, arg1, arg2)
    return obj
}
func (self *TransUnit_RefAccessSymPos) InitTransUnit_RefAccessSymPos(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 LnsAny) {
    self.symbol = arg1
    self.pos = arg2
}
func (self *TransUnit_RefAccessSymPos) Get_symbol(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbol }
func (self *TransUnit_RefAccessSymPos) Get_pos(_env *LnsEnv) LnsAny{ return self.pos }

// declaration Class -- TentativeSymbol
type TransUnit_TentativeSymbolMtd interface {
    AddAccessSym(_env *LnsEnv, arg1 *Ast_SymbolInfo)
    AddAccessSymPos(_env *LnsEnv, arg1 *TransUnit_AccessSymPos)
    CheckAndExclude(_env *LnsEnv, arg1 *Ast_SymbolInfo) bool
    ClearAccessSym(_env *LnsEnv)
    Finish(_env *LnsEnv, arg1 bool) LnsAny
    Get_accessSymList(_env *LnsEnv) *LnsList
    Get_initSymSet(_env *LnsEnv) *LnsSet
    Get_scope(_env *LnsEnv) *Ast_Scope
    merge(_env *LnsEnv, arg1 bool) bool
    ModSym(_env *LnsEnv, arg1 *Ast_SymbolInfo)
    NewSet(_env *LnsEnv, arg1 *Ast_Scope)
    Regist(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 Types_Position) bool
    Skip(_env *LnsEnv)
    SyncPos(_env *LnsEnv)
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
func NewTransUnit_TentativeSymbol(_env *LnsEnv, arg1 LnsAny, arg2 *Ast_Scope, arg3 *Ast_Scope, arg4 bool, arg5 LnsAny) *TransUnit_TentativeSymbol {
    obj := &TransUnit_TentativeSymbol{}
    obj.FP = obj
    obj.InitTransUnit_TentativeSymbol(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *TransUnit_TentativeSymbol) Get_initSymSet(_env *LnsEnv) *LnsSet{ return self.initSymSet }
func (self *TransUnit_TentativeSymbol) Get_scope(_env *LnsEnv) *Ast_Scope{ return self.scope }
func (self *TransUnit_TentativeSymbol) Get_accessSymList(_env *LnsEnv) *LnsList{ return self.accessSymList }
// 170: DeclConstr
func (self *TransUnit_TentativeSymbol) InitTransUnit_TentativeSymbol(_env *LnsEnv, parent LnsAny,scope *Ast_Scope,moduleScope *Ast_Scope,loopFlag bool,refAccessSymPosList LnsAny) {
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
        refAccessSymPosList_106 := refAccessSymPosList.(*LnsList)
        list = refAccessSymPosList_106
    } else {
        if loopFlag{
            list = TransUnit_clearThePosForModToRef_7_(_env, scope, moduleScope)
        } else { 
            list = NewLnsList([]LnsAny{})
        }
    }
    self.sym2posForModToRef = list
}


// declaration Class -- ClosureFun
type TransUnit_ClosureFunMtd interface {
    Check(_env *LnsEnv) bool
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
func NewTransUnit_ClosureFun(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 *Ast_Scope) *TransUnit_ClosureFun {
    obj := &TransUnit_ClosureFun{}
    obj.FP = obj
    obj.InitTransUnit_ClosureFun(_env, arg1, arg2)
    return obj
}
func (self *TransUnit_ClosureFun) InitTransUnit_ClosureFun(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 *Ast_Scope) {
    self.symbol = arg1
    self.fromScope = arg2
}

// declaration Class -- AccessSymbolSet
type TransUnit_AccessSymbolSetMtd interface {
    Add(_env *LnsEnv, arg1 *Ast_SymbolInfo)
    ApplyPos(_env *LnsEnv, arg1 LnsAny)
    Clone(_env *LnsEnv) *TransUnit_AccessSymbolSet
    Get_accessSym2Pos(_env *LnsEnv) *LnsMap
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
func NewTransUnit_AccessSymbolSet(_env *LnsEnv) *TransUnit_AccessSymbolSet {
    obj := &TransUnit_AccessSymbolSet{}
    obj.FP = obj
    obj.InitTransUnit_AccessSymbolSet(_env)
    return obj
}
func (self *TransUnit_AccessSymbolSet) Get_accessSym2Pos(_env *LnsEnv) *LnsMap{ return self.accessSym2Pos }
// 479: DeclConstr
func (self *TransUnit_AccessSymbolSet) InitTransUnit_AccessSymbolSet(_env *LnsEnv) {
    self.accessSym2Pos = NewLnsMap( map[LnsAny]LnsAny{})
}


// declaration Class -- AccessSymbolSetQueue
type TransUnit_AccessSymbolSetQueueMtd interface {
    Add(_env *LnsEnv, arg1 *Ast_SymbolInfo)
    GetMap(_env *LnsEnv) *LnsMap
    Pop(_env *LnsEnv, arg1 LnsAny)
    Push(_env *LnsEnv)
    setupFrom(_env *LnsEnv, arg1 *TransUnit_AccessSymbolSetQueue)
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
func NewTransUnit_AccessSymbolSetQueue(_env *LnsEnv) *TransUnit_AccessSymbolSetQueue {
    obj := &TransUnit_AccessSymbolSetQueue{}
    obj.FP = obj
    obj.InitTransUnit_AccessSymbolSetQueue(_env)
    return obj
}
// 507: DeclConstr
func (self *TransUnit_AccessSymbolSetQueue) InitTransUnit_AccessSymbolSetQueue(_env *LnsEnv) {
    self.queue = NewLnsList([]LnsAny{})
}


// declaration Class -- FuncBlockInfo
type TransUnit_FuncBlockInfoMtd interface {
    Get_declFuncInfo(_env *LnsEnv) *Nodes_DeclFuncInfo
    Get_declPos(_env *LnsEnv) Types_Position
    Get_funcScope(_env *LnsEnv) *Ast_Scope
    Get_funcType(_env *LnsEnv) *Ast_TypeInfo
    Get_nsInfo(_env *LnsEnv) *TransUnitIF_NSInfo
    Get_orgPos(_env *LnsEnv) LnsAny
    Get_tokenList(_env *LnsEnv) *LnsList
    Get_typeDataAccessor(_env *LnsEnv) Ast_TypeDataAccessor
    set_funcScope(_env *LnsEnv, arg1 *Ast_Scope)
}
type TransUnit_FuncBlockInfo struct {
    declFuncInfo *Nodes_DeclFuncInfo
    nsInfo *TransUnitIF_NSInfo
    funcType *Ast_TypeInfo
    typeDataAccessor Ast_TypeDataAccessor
    funcScope *Ast_Scope
    tokenList *LnsList
    declPos Types_Position
    orgPos LnsAny
    FP TransUnit_FuncBlockInfoMtd
}
func TransUnit_FuncBlockInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_FuncBlockInfo).FP
}
type TransUnit_FuncBlockInfoDownCast interface {
    ToTransUnit_FuncBlockInfo() *TransUnit_FuncBlockInfo
}
func TransUnit_FuncBlockInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_FuncBlockInfoDownCast)
    if ok { return work.ToTransUnit_FuncBlockInfo() }
    return nil
}
func (obj *TransUnit_FuncBlockInfo) ToTransUnit_FuncBlockInfo() *TransUnit_FuncBlockInfo {
    return obj
}
func NewTransUnit_FuncBlockInfo(_env *LnsEnv, arg1 *Nodes_DeclFuncInfo, arg2 *TransUnitIF_NSInfo, arg3 *Ast_TypeInfo, arg4 Ast_TypeDataAccessor, arg5 *Ast_Scope, arg6 *LnsList, arg7 Types_Position, arg8 LnsAny) *TransUnit_FuncBlockInfo {
    obj := &TransUnit_FuncBlockInfo{}
    obj.FP = obj
    obj.InitTransUnit_FuncBlockInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *TransUnit_FuncBlockInfo) InitTransUnit_FuncBlockInfo(_env *LnsEnv, arg1 *Nodes_DeclFuncInfo, arg2 *TransUnitIF_NSInfo, arg3 *Ast_TypeInfo, arg4 Ast_TypeDataAccessor, arg5 *Ast_Scope, arg6 *LnsList, arg7 Types_Position, arg8 LnsAny) {
    self.declFuncInfo = arg1
    self.nsInfo = arg2
    self.funcType = arg3
    self.typeDataAccessor = arg4
    self.funcScope = arg5
    self.tokenList = arg6
    self.declPos = arg7
    self.orgPos = arg8
}
func (self *TransUnit_FuncBlockInfo) Get_declFuncInfo(_env *LnsEnv) *Nodes_DeclFuncInfo{ return self.declFuncInfo }
func (self *TransUnit_FuncBlockInfo) Get_nsInfo(_env *LnsEnv) *TransUnitIF_NSInfo{ return self.nsInfo }
func (self *TransUnit_FuncBlockInfo) Get_funcType(_env *LnsEnv) *Ast_TypeInfo{ return self.funcType }
func (self *TransUnit_FuncBlockInfo) Get_typeDataAccessor(_env *LnsEnv) Ast_TypeDataAccessor{ return self.typeDataAccessor }
func (self *TransUnit_FuncBlockInfo) Get_funcScope(_env *LnsEnv) *Ast_Scope{ return self.funcScope }
func (self *TransUnit_FuncBlockInfo) set_funcScope(_env *LnsEnv, arg1 *Ast_Scope){ self.funcScope = arg1 }
func (self *TransUnit_FuncBlockInfo) Get_tokenList(_env *LnsEnv) *LnsList{ return self.tokenList }
func (self *TransUnit_FuncBlockInfo) Get_declPos(_env *LnsEnv) Types_Position{ return self.declPos }
func (self *TransUnit_FuncBlockInfo) Get_orgPos(_env *LnsEnv) LnsAny{ return self.orgPos }
func TransUnit_FuncBlockInfo__createPipe_10_(_env *LnsEnv, arg1 LnsInt) LnsAny{
   return NewLnspipe( arg1 )
}

type TransUnit_FuncBlockCtlIF interface {
        GetNext(_env *LnsEnv) LnsAny
}
func Lns_cast2TransUnit_FuncBlockCtlIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(TransUnit_FuncBlockCtlIF); ok { 
        return obj
    }
    return nil
}

// declaration Class -- ListFuncBlockCtl
type TransUnit_ListFuncBlockCtlMtd interface {
    GetNext(_env *LnsEnv) LnsAny
}
type TransUnit_ListFuncBlockCtl struct {
    list *LnsList
    pos LnsInt
    FP TransUnit_ListFuncBlockCtlMtd
}
func TransUnit_ListFuncBlockCtl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_ListFuncBlockCtl).FP
}
type TransUnit_ListFuncBlockCtlDownCast interface {
    ToTransUnit_ListFuncBlockCtl() *TransUnit_ListFuncBlockCtl
}
func TransUnit_ListFuncBlockCtlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_ListFuncBlockCtlDownCast)
    if ok { return work.ToTransUnit_ListFuncBlockCtl() }
    return nil
}
func (obj *TransUnit_ListFuncBlockCtl) ToTransUnit_ListFuncBlockCtl() *TransUnit_ListFuncBlockCtl {
    return obj
}
func NewTransUnit_ListFuncBlockCtl(_env *LnsEnv, arg1 *LnsList) *TransUnit_ListFuncBlockCtl {
    obj := &TransUnit_ListFuncBlockCtl{}
    obj.FP = obj
    obj.InitTransUnit_ListFuncBlockCtl(_env, arg1)
    return obj
}
// 550: DeclConstr
func (self *TransUnit_ListFuncBlockCtl) InitTransUnit_ListFuncBlockCtl(_env *LnsEnv, list *LnsList) {
    self.list = list
    self.pos = 0
}


// declaration Class -- PipeFuncBlockCtl
type TransUnit_PipeFuncBlockCtlMtd interface {
    GetNext(_env *LnsEnv) LnsAny
}
type TransUnit_PipeFuncBlockCtl struct {
    pipe *Lns__pipe
    FP TransUnit_PipeFuncBlockCtlMtd
}
func TransUnit_PipeFuncBlockCtl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_PipeFuncBlockCtl).FP
}
type TransUnit_PipeFuncBlockCtlDownCast interface {
    ToTransUnit_PipeFuncBlockCtl() *TransUnit_PipeFuncBlockCtl
}
func TransUnit_PipeFuncBlockCtlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_PipeFuncBlockCtlDownCast)
    if ok { return work.ToTransUnit_PipeFuncBlockCtl() }
    return nil
}
func (obj *TransUnit_PipeFuncBlockCtl) ToTransUnit_PipeFuncBlockCtl() *TransUnit_PipeFuncBlockCtl {
    return obj
}
func NewTransUnit_PipeFuncBlockCtl(_env *LnsEnv, arg1 *Lns__pipe) *TransUnit_PipeFuncBlockCtl {
    obj := &TransUnit_PipeFuncBlockCtl{}
    obj.FP = obj
    obj.InitTransUnit_PipeFuncBlockCtl(_env, arg1)
    return obj
}
// 567: DeclConstr
func (self *TransUnit_PipeFuncBlockCtl) InitTransUnit_PipeFuncBlockCtl(_env *LnsEnv, pipe *Lns__pipe) {
    self.pipe = pipe
}


// declaration Class -- FuncBlockResult
type TransUnit_FuncBlockResultMtd interface {
    Get_body(_env *LnsEnv) *Nodes_BlockNode
    Get_funcBlockInfo(_env *LnsEnv) *TransUnit_FuncBlockInfo
    Get_has_func_sym(_env *LnsEnv) bool
    Get_stmtNum(_env *LnsEnv) LnsInt
}
type TransUnit_FuncBlockResult struct {
    funcBlockInfo *TransUnit_FuncBlockInfo
    body *Nodes_BlockNode
    has_func_sym bool
    stmtNum LnsInt
    FP TransUnit_FuncBlockResultMtd
}
func TransUnit_FuncBlockResult2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_FuncBlockResult).FP
}
type TransUnit_FuncBlockResultDownCast interface {
    ToTransUnit_FuncBlockResult() *TransUnit_FuncBlockResult
}
func TransUnit_FuncBlockResultDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_FuncBlockResultDownCast)
    if ok { return work.ToTransUnit_FuncBlockResult() }
    return nil
}
func (obj *TransUnit_FuncBlockResult) ToTransUnit_FuncBlockResult() *TransUnit_FuncBlockResult {
    return obj
}
func NewTransUnit_FuncBlockResult(_env *LnsEnv, arg1 *TransUnit_FuncBlockInfo, arg2 *Nodes_BlockNode, arg3 bool, arg4 LnsInt) *TransUnit_FuncBlockResult {
    obj := &TransUnit_FuncBlockResult{}
    obj.FP = obj
    obj.InitTransUnit_FuncBlockResult(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *TransUnit_FuncBlockResult) InitTransUnit_FuncBlockResult(_env *LnsEnv, arg1 *TransUnit_FuncBlockInfo, arg2 *Nodes_BlockNode, arg3 bool, arg4 LnsInt) {
    self.funcBlockInfo = arg1
    self.body = arg2
    self.has_func_sym = arg3
    self.stmtNum = arg4
}
func (self *TransUnit_FuncBlockResult) Get_funcBlockInfo(_env *LnsEnv) *TransUnit_FuncBlockInfo{ return self.funcBlockInfo }
func (self *TransUnit_FuncBlockResult) Get_body(_env *LnsEnv) *Nodes_BlockNode{ return self.body }
func (self *TransUnit_FuncBlockResult) Get_has_func_sym(_env *LnsEnv) bool{ return self.has_func_sym }
func (self *TransUnit_FuncBlockResult) Get_stmtNum(_env *LnsEnv) LnsInt{ return self.stmtNum }

// declaration Class -- TransUnit
type TransUnit_TransUnitMtd interface {
    MultiTo1(_env *LnsEnv, arg1 *Nodes_Node) *Nodes_Node
    accessSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 bool)
    addAccessor(_env *LnsEnv, arg1 *Nodes_DeclMemberNode, arg2 *LnsSet, arg3 *Ast_Scope, arg4 *Ast_TypeInfo, arg5 Ast_TypeDataAccessor)
    addDefaultConstructor(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 Ast_TypeDataAccessor, arg4 *Ast_Scope, arg5 *LnsList, arg6 *LnsSet, arg7 bool)
    AddErrMess(_env *LnsEnv, arg1 Types_Position, arg2 string)
    addFuncBlockInfoList(_env *LnsEnv, arg1 *TransUnit_FuncBlockInfo)
    AddLocalVar(_env *LnsEnv, arg1 Types_Position, arg2 bool, arg3 bool, arg4 string, arg5 *Ast_TypeInfo, arg6 LnsInt, arg7 LnsAny) *Ast_SymbolInfo
    addMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Nodes_Node, arg3 string)
    addWarnErrMess(_env *LnsEnv, arg1 Types_Position, arg2 bool, arg3 string)
    addWarnMess(_env *LnsEnv, arg1 Types_Position, arg2 string)
    analyzeAccessClassField(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 *Types_Token)(*Ast_TypeInfo, LnsAny, bool)
    analyzeAlias(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_AliasNode
    analyzeApply(_env *LnsEnv, arg1 *Types_Token) *Nodes_ApplyNode
    analyzeAsyncLock(_env *LnsEnv, arg1 *Types_Token, arg2 LnsInt) *Nodes_Node
    analyzeBlock(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsAny, arg4 LnsAny) *Nodes_BlockNode
    analyzeClassBody(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 LnsInt, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 Ast_TypeDataAccessor, arg8 *Types_Token, arg9 LnsAny, arg10 LnsAny, arg11 LnsInt, arg12 *Types_Token, arg13 *Nodes_ClassInheritInfo)(*Nodes_DeclClassNode, *Types_Token, *LnsSet)
    analyzeDecl(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token) LnsAny
    analyzeDeclAlge(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclAlgeNode
    analyzeDeclAlternateType(_env *LnsEnv, arg1 bool, arg2 *Types_Token, arg3 LnsInt)(*Types_Token, *LnsList)
    analyzeDeclArgList(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_Scope, arg3 *LnsList, arg4 bool) *Types_Token
    analyzeDeclClass(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 LnsInt) *Nodes_DeclClassNode
    analyzeDeclEnum(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclEnumNode
    analyzeDeclForm(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclFormNode
    analyzeDeclFunc(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 bool, arg4 bool, arg5 LnsInt, arg6 bool, arg7 LnsAny, arg8 *Types_Token, arg9 LnsAny) *Nodes_Node
    analyzeDeclMacro(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclMacroNode
    analyzeDeclMacroSub(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token, arg3 *Types_Token, arg4 *Ast_Scope, arg5 *Ast_TypeInfo, arg6 Ast_TypeDataAccessor, arg7 *LnsList) *Nodes_DeclMacroNode
    analyzeDeclMember(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool, arg4 *Types_Token) *Nodes_DeclMemberNode
    analyzeDeclMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool, arg4 bool, arg5 LnsInt, arg6 bool, arg7 *Types_Token, arg8 *Types_Token) *Nodes_Node
    analyzeDeclProto(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_Node
    AnalyzeDeclToken(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token) LnsAny
    analyzeDeclVar(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 *Types_Token) *Nodes_Node
    analyzeExp(_env *LnsEnv, arg1 bool, arg2 bool, arg3 bool, arg4 LnsAny, arg5 LnsAny) *Nodes_Node
    analyzeExpCall(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 *Types_Token)(*Nodes_Node, *Types_Token)
    analyzeExpCast(_env *LnsEnv, arg1 *Types_Token, arg2 string, arg3 *Nodes_Node) *Nodes_Node
    analyzeExpCont(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 bool, arg4 bool) *Nodes_Node
    analyzeExpField(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token, arg3 LnsInt, arg4 *Nodes_Node) *Nodes_Node
    analyzeExpList(_env *LnsEnv, arg1 bool, arg2 bool, arg3 bool, arg4 LnsAny, arg5 LnsAny, arg6 LnsAny) *Nodes_ExpListNode
    analyzeExpMacroStat(_env *LnsEnv, arg1 *Types_Token) *Nodes_ExpMacroStatNode
    analyzeExpOneRVal(_env *LnsEnv, arg1 bool, arg2 bool, arg3 LnsAny, arg4 LnsAny) *Nodes_Node
    analyzeExpOp2(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 LnsAny) *Nodes_Node
    analyzeExpOpEquals(_env *LnsEnv, arg1 Types_Position, arg2 *Types_Token, arg3 *Nodes_Node, arg4 *Nodes_Node)(*Nodes_Node, *Nodes_Node)
    analyzeExpOpSet(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Types_Token, arg3 *LnsList) *Nodes_Node
    analyzeExpRefItem(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 bool) *Nodes_Node
    analyzeExpSymbol(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token, arg3 LnsInt, arg4 LnsAny, arg5 bool, arg6 bool) *Nodes_Node
    analyzeExpUnwrap(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeExtend(_env *LnsEnv, arg1 LnsInt, arg2 Types_Position)(*Types_Token, LnsAny, *LnsList, *LnsMap, *Nodes_ClassInheritInfo)
    analyzeFor(_env *LnsEnv, arg1 *Types_Token) *Nodes_ForNode
    analyzeForeach(_env *LnsEnv, arg1 *Types_Token, arg2 bool) *Nodes_Node
    analyzeFuncBlock(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 string, arg6 *Ast_Scope, arg7 *LnsList) *Nodes_BlockNode
    analyzeIf(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeIfUnwrap(_env *LnsEnv, arg1 *Types_Token) *Nodes_IfUnwrapNode
    analyzeInitExp(_env *LnsEnv, arg1 Types_Position, arg2 LnsInt, arg3 bool, arg4 *LnsList, arg5 *LnsList)(*LnsList, *LnsList, *LnsList, LnsAny)
    analyzeLetAndInitExp(_env *LnsEnv, arg1 Types_Position, arg2 bool, arg3 LnsInt, arg4 LnsInt, arg5 bool)(*LnsList, *LnsList, *LnsList, LnsAny)
    analyzeListConst(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny) *Nodes_Node
    analyzeListItems(_env *LnsEnv, arg1 Types_Position, arg2 *Types_Token, arg3 string, arg4 LnsAny)(LnsAny, *Ast_TypeInfo)
    analyzeLuneControl(_env *LnsEnv, arg1 *Types_Token) LnsAny
    AnalyzeLuneControlToken(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token) LnsAny
    analyzeMapConst(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny) *Nodes_LiteralMapNode
    analyzeMatch(_env *LnsEnv, arg1 *Types_Token) *Nodes_MatchNode
    analyzeNewAlge(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_AlgeTypeInfo, arg3 LnsAny) *Nodes_NewAlgeValNode
    analyzePushClass(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsInt, arg9 *LnsList)(*Types_Token, *TransUnitIF_NSInfo, *Nodes_ClassInheritInfo)
    analyzeRefType(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 bool) *Nodes_RefTypeNode
    analyzeRefTypeWithSymbol(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 LnsAny, arg4 *Nodes_Node, arg5 bool) *Nodes_RefTypeNode
    analyzeRepeat(_env *LnsEnv, arg1 *Types_Token) *Nodes_RepeatNode
    analyzeRequest(_env *LnsEnv, arg1 *Types_Token) *Nodes_RequestNode
    analyzeRetTypeList(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 bool)(*LnsList, *Types_Token, *LnsList)
    analyzeReturn(_env *LnsEnv, arg1 *Types_Token) *Nodes_ReturnNode
    analyzeScope(_env *LnsEnv, arg1 *Types_Token) *Nodes_ScopeNode
    analyzeSetConst(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny) *Nodes_Node
    analyzeStatement(_env *LnsEnv, arg1 LnsAny) LnsAny
    analyzeStatementList(_env *LnsEnv, arg1 *LnsList, arg2 bool, arg3 LnsAny)(LnsAny, LnsInt)
    analyzeStatementListSubfile(_env *LnsEnv, arg1 *LnsList) LnsAny
    AnalyzeStatementToken(_env *LnsEnv, arg1 *Types_Token) LnsAny
    analyzeStrConst(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token) *Nodes_Node
    analyzeSuper(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeSwitch(_env *LnsEnv, arg1 *Types_Token) *Nodes_SwitchNode
    analyzeUnwrap(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeWhen(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeWhile(_env *LnsEnv, arg1 *Types_Token) *Nodes_WhileNode
    canBeAsyncParam(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    checkAccesSym(_env *LnsEnv)
    checkAlgeComp(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_AlgeTypeInfo)
    checkArgForSort(_env *LnsEnv, arg1 *Types_Token, arg2 *LnsList, arg3 *Nodes_ExpListNode)
    checkArgForStringForm(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_ExpListNode)
    checkAsyncField(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 Types_Position)
    checkAsyncSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 Types_Position)
    checkComp(_env *LnsEnv, arg1 *Types_Token, arg2 TransUnit_checkCompForm_31_)
    checkEnumComp(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_EnumTypeInfo)
    checkFieldComp(_env *LnsEnv, arg1 bool, arg2 *Types_Token, arg3 *Nodes_Node)
    checkImplicitCast(_env *LnsEnv, arg1 *LnsMap, arg2 bool, arg3 *LnsList, arg4 *Nodes_ExpListNode, arg5 TransUnit_checkImplicitCastCallback_10_) LnsAny
    checkLiteralEmptyCollection(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 *Ast_TypeInfo)
    checkMatchType(_env *LnsEnv, arg1 string, arg2 Types_Position, arg3 *LnsList, arg4 LnsAny, arg5 bool, arg6 bool, arg7 LnsAny, arg8 bool)(LnsInt, *LnsMap, LnsAny, *LnsList)
    checkMatchValType(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 *LnsList, arg5 LnsAny)(LnsInt, *LnsMap, LnsAny)
    checkNextToken(_env *LnsEnv, arg1 string) *Types_Token
    checkNoasyncType(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo)
    checkOverrideMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo) *LnsList
    checkOverriededMethodOfAllClass(_env *LnsEnv)
    checkPublic(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo)
    checkShadowing(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 *Ast_Scope)
    checkStringFormat(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 *LnsList)
    checkSymbol(_env *LnsEnv, arg1 *Types_Token, arg2 LnsInt) *Types_Token
    checkSymbolComp(_env *LnsEnv, arg1 *Types_Token)
    checkSymbolHavingValue(_env *LnsEnv, arg1 Types_Position, arg2 *LnsList)
    checkToken(_env *LnsEnv, arg1 *Types_Token, arg2 string) *Types_Token
    convToExtTypeList(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *LnsList) *LnsList
    createDummyNS(_env *LnsEnv, arg1 *Ast_Scope, arg2 Types_Position, arg3 LnsInt)
    createExpList(_env *LnsEnv, arg1 Types_Position, arg2 *LnsList, arg3 *LnsList, arg4 bool, arg5 LnsAny) *Nodes_ExpListNode
    createExpListNode(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 *LnsList) *Nodes_ExpListNode
    createExtType(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo) *Ast_TypeInfo
    createModifier(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt) *Ast_TypeInfo
    createNoneNode(_env *LnsEnv, arg1 Types_Position) *Nodes_Node
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    declFuncPostProcess(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsAny, arg3 *Nodes_BlockNode, arg4 *Ast_Scope)
    dumpComp(_env *LnsEnv, arg1 Writer_Writer, arg2 string, arg3 *Ast_SymbolInfo, arg4 bool) bool
    dumpFieldComp(_env *LnsEnv, arg1 Writer_Writer, arg2 bool, arg3 *Ast_TypeInfo, arg4 string, arg5 LnsAny)
    dumpSymbolComp(_env *LnsEnv, arg1 Writer_Writer, arg2 *Ast_Scope, arg3 string)
    dumpSymbolType(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 *Ast_TypeInfo)
    Error(_env *LnsEnv, arg1 string)
    ErrorAt(_env *LnsEnv, arg1 Types_Position, arg2 string)
    errorShadowing(_env *LnsEnv, arg1 Types_Position, arg2 LnsAny)
    ErrorShadowingOp(_env *LnsEnv, arg1 Types_Position, arg2 LnsAny, arg3 bool)
    evalMacro(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_TypeInfo, arg3 LnsAny) *Nodes_ExpMacroExpNode
    evalMacroOp(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 Macro_EvalMacroCallback)
    finishTentativeSymbol(_env *LnsEnv, arg1 bool)
    getContinueToken(_env *LnsEnv)(*Types_Token, bool)
    getCurrentClass(_env *LnsEnv) *Ast_TypeInfo
    GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    getDefaultAsync(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny) LnsInt
    GetLatestPos(_env *LnsEnv) Types_Position
    getLineNo(_env *LnsEnv, arg1 *Types_Token) Types_Position
    getMutableAsync(_env *LnsEnv, arg1 *Types_Token)(*Types_Token, LnsInt, LnsAny)
    GetNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) *TransUnitIF_NSInfo
    GetStreamName(_env *LnsEnv) string
    getSymbolToken(_env *LnsEnv, arg1 LnsInt) *Types_Token
    getToken(_env *LnsEnv, arg1 LnsAny) *Types_Token
    GetTokenNoErr(_env *LnsEnv, arg1 LnsAny) *Types_Token
    Get_curNsInfo(_env *LnsEnv) *TransUnitIF_NSInfo
    Get_errMessList(_env *LnsEnv) *LnsList
    Get_globalScope(_env *LnsEnv) *Ast_Scope
    Get_importedAliasMap(_env *LnsEnv) *LnsMap
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_scopeRO(_env *LnsEnv) *Ast_Scope
    Get_warnMessList(_env *LnsEnv) *LnsList
    inAnalyzingState(_env *LnsEnv, arg1 LnsInt) bool
    isTargetToken(_env *LnsEnv, arg1 *Types_Token) bool
    isTargetTokenPos(_env *LnsEnv, arg1 string, arg2 Types_Position) bool
    mergeFrom(_env *LnsEnv, arg1 *TransUnit_TransUnit, arg2 *LnsMap)
    mergeTentativeSymbol(_env *LnsEnv, arg1 *Ast_Scope)
    NewNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Types_Position) *TransUnitIF_NSInfo
    NewNSInfoWithTypeData(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Ast_TypeDataAccessor, arg3 Types_Position) *TransUnitIF_NSInfo
    NewPushback(_env *LnsEnv, arg1 *LnsList)
    popAnalyzingState(_env *LnsEnv)
    PopClass(_env *LnsEnv)
    PopModule(_env *LnsEnv)
    PopScope(_env *LnsEnv)
    prepareExpCall(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 *Ast_TypeInfo)(*LnsMap, LnsAny)
    prepareTentativeSymbol(_env *LnsEnv, arg1 *Ast_Scope, arg2 bool, arg3 LnsAny)
    processAddFunc(_env *LnsEnv, arg1 bool, arg2 *Ast_Scope, arg3 *Types_Token, arg4 *Ast_TypeInfo, arg5 *LnsMap)(*Ast_SymbolInfo, *TransUnitIF_NSInfo)
    processCaseDefault(_env *LnsEnv, arg1 *Types_Token, arg2 LnsInt, arg3 *Types_Token, arg4 bool)(LnsAny, bool)
    processFunc(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token, arg3 LnsAny, arg4 *Nodes_Node, arg5 *Ast_TypeInfo, arg6 *LnsMap, arg7 *LnsList, arg8 *Ast_TypeInfo, arg9 LnsAny) *Nodes_Node
    processFuncBlockInfo(_env *LnsEnv, arg1 TransUnit_FuncBlockCtlIF, arg2 string) *LnsMap
    pushAnalyzingState(_env *LnsEnv, arg1 LnsInt)
    PushClass(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *TransUnitIF_NSInfo
    PushClassLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *Ast_TypeInfo
    PushClassScope(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
    pushExtModule(_env *LnsEnv, arg1 bool, arg2 string, arg3 LnsInt, arg4 Types_Position, arg5 bool, arg6 LnsInt, arg7 string) *TransUnitIF_NSInfo
    PushModule(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *TransUnitIF_NSInfo
    PushModuleLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *Ast_TypeInfo
    PushScope(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny) *Ast_Scope
    Pushback(_env *LnsEnv)
    PushbackStr(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 string, arg4 Types_Position)
    PushbackToken(_env *LnsEnv, arg1 *Types_Token)
    SetParser(_env *LnsEnv, arg1 *Parser_DefaultPushbackParser)
    SetScope(_env *LnsEnv, arg1 *Ast_Scope, arg2 LnsInt)
    Set_curNsInfo(_env *LnsEnv, arg1 *TransUnitIF_NSInfo)
    setup(_env *LnsEnv, arg1 *TransUnit_TransUnit)
    skipAndCreateDummyBlock(_env *LnsEnv) *Nodes_BlockNode
    skipBlock(_env *LnsEnv, arg1 bool) *LnsList
    supportLang(_env *LnsEnv, arg1 string) bool
}
type TransUnit_TransUnit struct {
    TransUnitIF_TransUnitBase
    analyzeMode LnsInt
    analyzePos Types_Position
    analyzeModule string
    TargetLuaVer *LuaVer_LuaVerInfo
    ModuleId *FrontInterface_ModuleId
    macroEval *Nodes_MacroEval
    builtinFunc *Builtin_BuiltinFuncType
    ValidMutControl bool
    moduleName string
    moduleType *Ast_TypeInfo
    IgnoreToCheckSymbol_ bool
    defaultAsyncMode LnsInt
    BaseDir LnsAny
    analyzePhase LnsInt
    Modifier *TransUnitIF_Modifier
    TopScope *Ast_Scope
    macroScope LnsAny
    tentativeSymbol *TransUnit_TentativeSymbol
    Parser *Parser_DefaultPushbackParser
    ScopeAccess LnsInt
    analyzingStaticMethodArgsScope LnsAny
    inTestBlock bool
    FuncBlockInfoLinkNo LnsAny
    ModuleScope *Ast_Scope
    Has__func__Symbol *LnsSet
    NodeManager *Nodes_NodeManager
    closureFunList *LnsList
    commentCtrl *Parser_CommentCtrl
    typeInfo2ClassNode *LnsMap
    analyzingStateQueue *LnsList
    MacroCtrl *Macro_MacroCtrl
    advertisedTypeSet *LnsSet
    accessSymbolSetQueue *TransUnit_AccessSymbolSetQueue
    class2defaultAsyncMode *LnsMap
    HelperInfo *FrontInterface_LuneHelperInfo
    WarnMessList *LnsList
    ImportModuleSet *LnsSet
    ImportedAliasMap *LnsMap
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
func NewTransUnit_TransUnit(_env *LnsEnv, arg1 *FrontInterface_ModuleId, arg2 *FrontInterface_ImportModuleInfo, arg3 *Nodes_MacroEval, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 *LuaVer_LuaVerInfo, arg9 *Types_TransCtrlInfo, arg10 *Builtin_BuiltinFuncType) *TransUnit_TransUnit {
    obj := &TransUnit_TransUnit{}
    obj.FP = obj
    obj.TransUnitIF_TransUnitBase.FP = obj
    obj.InitTransUnit_TransUnit(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *TransUnit_TransUnit) Get_warnMessList(_env *LnsEnv) *LnsList{ return self.WarnMessList }
func (self *TransUnit_TransUnit) Get_importedAliasMap(_env *LnsEnv) *LnsMap{ return self.ImportedAliasMap }
// 828: DeclConstr
func (self *TransUnit_TransUnit) InitTransUnit_TransUnit(_env *LnsEnv, moduleId *FrontInterface_ModuleId,importModuleInfo *FrontInterface_ImportModuleInfo,macroEval *Nodes_MacroEval,enableMultiPhase bool,analyzeModule LnsAny,mode LnsAny,pos LnsAny,targetLuaVer *LuaVer_LuaVerInfo,ctrl_info *Types_TransCtrlInfo,builtinFunc *Builtin_BuiltinFuncType) {
    self.InitTransUnitIF_TransUnitBase(TransUnit_convExp0_878(_env, Lns_2DDD(TransUnit_TransUnit_getSuperParam_5_(_env, ctrl_info))))
    self.FuncBlockInfoLinkNo = nil
    var phase LnsInt
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( ctrl_info.ValidMultiPhaseTransUnit) &&
        _env.SetStackVal( enableMultiPhase) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( mode == nil) ||
            _env.SetStackVal( mode == TransUnit_AnalyzeMode__Compile) ).(bool))) ).(bool)){
        phase = TransUnit_AnalyzePhase__Meta
    } else { 
        phase = TransUnit_AnalyzePhase__Main
    }
    self.analyzePhase = phase
    self.inTestBlock = false
    self.BaseDir = nil
    self.builtinFunc = builtinFunc
    self.analyzingStaticMethodArgsScope = nil
    self.class2defaultAsyncMode = NewLnsMap( map[LnsAny]LnsAny{})
    var defaultAsyncMode LnsInt
    if ctrl_info.DefaultAsync{
        defaultAsyncMode = TransUnit_DefaultAsyncMode__AsyncAll
    } else { 
        defaultAsyncMode = TransUnit_DefaultAsyncMode__NoAsync
    }
    self.defaultAsyncMode = defaultAsyncMode
    self.ImportedAliasMap = NewLnsMap( map[LnsAny]LnsAny{})
    self.ImportModuleSet = NewLnsSet([]LnsAny{})
    self.accessSymbolSetQueue = NewTransUnit_AccessSymbolSetQueue(_env)
    self.advertisedTypeSet = NewLnsSet([]LnsAny{})
    self.closureFunList = NewLnsList([]LnsAny{})
    self.ScopeAccess = Ast_ScopeAccess__Normal
    self.macroEval = macroEval
    self.MacroCtrl = NewMacro_MacroCtrl(_env, macroEval, ctrl_info.ValidMacroAsync)
    self.analyzingStateQueue = NewLnsList([]LnsAny{})
    self.IgnoreToCheckSymbol_ = false
    self.ModuleId = moduleId
    self.HelperInfo = NewFrontInterface_LuneHelperInfo(_env)
    self.TargetLuaVer = targetLuaVer
    self.Has__func__Symbol = NewLnsSet([]LnsAny{})
    self.NodeManager = NewNodes_NodeManager(_env)
    self.macroScope = nil
    self.ValidMutControl = true
    self.Modifier = NewTransUnitIF_Modifier(_env, self.ValidMutControl, self.ProcessInfo)
    self.moduleName = ""
    self.moduleType = Ast_headTypeInfo
    self.Parser = NewParser_DefaultPushbackParser(_env, &NewParser_DummyParser(_env).Parser_Parser)
    self.TopScope = self.FP.Get_scope(_env)
    self.ModuleScope = self.FP.Get_scope(_env)
    self.tentativeSymbol = NewTransUnit_TentativeSymbol(_env, nil, self.GlobalScope, self.ModuleScope, false, nil)
    self.typeInfo2ClassNode = NewLnsMap( map[LnsAny]LnsAny{})
    self.commentCtrl = NewParser_CommentCtrl(_env)
    self.WarnMessList = NewLnsList([]LnsAny{})
    self.analyzeMode = Lns_unwrapDefault( mode, TransUnit_AnalyzeMode__Compile).(LnsInt)
    self.analyzePos = Lns_unwrapDefault( pos, self.FP.CreatePosition(_env, 0, 0)).(Types_Position)
    self.analyzeModule = Lns_unwrapDefault( analyzeModule, "").(string)
}




















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
func NewTransUnit_LetVarInfo(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token, arg3 LnsAny) *TransUnit_LetVarInfo {
    obj := &TransUnit_LetVarInfo{}
    obj.FP = obj
    obj.InitTransUnit_LetVarInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *TransUnit_LetVarInfo) InitTransUnit_LetVarInfo(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token, arg3 LnsAny) {
    self.Mutable = arg1
    self.VarName = arg2
    self.VarType = arg3
}

// declaration Class -- TransUnitRunner
type TransUnit_TransUnitRunnerMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    MultiTo1(_env *LnsEnv, arg1 *Nodes_Node) *Nodes_Node
    accessSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 bool)
    addAccessor(_env *LnsEnv, arg1 *Nodes_DeclMemberNode, arg2 *LnsSet, arg3 *Ast_Scope, arg4 *Ast_TypeInfo, arg5 Ast_TypeDataAccessor)
    addDefaultConstructor(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 Ast_TypeDataAccessor, arg4 *Ast_Scope, arg5 *LnsList, arg6 *LnsSet, arg7 bool)
    AddErrMess(_env *LnsEnv, arg1 Types_Position, arg2 string)
    addFuncBlockInfoList(_env *LnsEnv, arg1 *TransUnit_FuncBlockInfo)
    AddLocalVar(_env *LnsEnv, arg1 Types_Position, arg2 bool, arg3 bool, arg4 string, arg5 *Ast_TypeInfo, arg6 LnsInt, arg7 LnsAny) *Ast_SymbolInfo
    analyzeAccessClassField(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 *Types_Token)(*Ast_TypeInfo, LnsAny, bool)
    analyzeAlias(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_AliasNode
    analyzeApply(_env *LnsEnv, arg1 *Types_Token) *Nodes_ApplyNode
    analyzeAsyncLock(_env *LnsEnv, arg1 *Types_Token, arg2 LnsInt) *Nodes_Node
    analyzeBlock(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsAny, arg4 LnsAny) *Nodes_BlockNode
    analyzeClassBody(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 LnsInt, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 Ast_TypeDataAccessor, arg8 *Types_Token, arg9 LnsAny, arg10 LnsAny, arg11 LnsInt, arg12 *Types_Token, arg13 *Nodes_ClassInheritInfo)(*Nodes_DeclClassNode, *Types_Token, *LnsSet)
    analyzeDecl(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token) LnsAny
    analyzeDeclAlge(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclAlgeNode
    analyzeDeclAlternateType(_env *LnsEnv, arg1 bool, arg2 *Types_Token, arg3 LnsInt)(*Types_Token, *LnsList)
    analyzeDeclArgList(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_Scope, arg3 *LnsList, arg4 bool) *Types_Token
    analyzeDeclClass(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 LnsInt) *Nodes_DeclClassNode
    analyzeDeclEnum(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclEnumNode
    analyzeDeclForm(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclFormNode
    analyzeDeclFunc(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 bool, arg4 bool, arg5 LnsInt, arg6 bool, arg7 LnsAny, arg8 *Types_Token, arg9 LnsAny) *Nodes_Node
    analyzeDeclMacro(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclMacroNode
    analyzeDeclMacroSub(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token, arg3 *Types_Token, arg4 *Ast_Scope, arg5 *Ast_TypeInfo, arg6 Ast_TypeDataAccessor, arg7 *LnsList) *Nodes_DeclMacroNode
    analyzeDeclMember(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool, arg4 *Types_Token) *Nodes_DeclMemberNode
    analyzeDeclMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool, arg4 bool, arg5 LnsInt, arg6 bool, arg7 *Types_Token, arg8 *Types_Token) *Nodes_Node
    analyzeDeclProto(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_Node
    AnalyzeDeclToken(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token) LnsAny
    analyzeDeclVar(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 *Types_Token) *Nodes_Node
    analyzeExp(_env *LnsEnv, arg1 bool, arg2 bool, arg3 bool, arg4 LnsAny, arg5 LnsAny) *Nodes_Node
    analyzeExpCall(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 *Types_Token)(*Nodes_Node, *Types_Token)
    analyzeExpCast(_env *LnsEnv, arg1 *Types_Token, arg2 string, arg3 *Nodes_Node) *Nodes_Node
    analyzeExpCont(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 bool, arg4 bool) *Nodes_Node
    analyzeExpField(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token, arg3 LnsInt, arg4 *Nodes_Node) *Nodes_Node
    analyzeExpList(_env *LnsEnv, arg1 bool, arg2 bool, arg3 bool, arg4 LnsAny, arg5 LnsAny, arg6 LnsAny) *Nodes_ExpListNode
    analyzeExpMacroStat(_env *LnsEnv, arg1 *Types_Token) *Nodes_ExpMacroStatNode
    analyzeExpOneRVal(_env *LnsEnv, arg1 bool, arg2 bool, arg3 LnsAny, arg4 LnsAny) *Nodes_Node
    analyzeExpOp2(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 LnsAny) *Nodes_Node
    analyzeExpOpEquals(_env *LnsEnv, arg1 Types_Position, arg2 *Types_Token, arg3 *Nodes_Node, arg4 *Nodes_Node)(*Nodes_Node, *Nodes_Node)
    analyzeExpOpSet(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Types_Token, arg3 *LnsList) *Nodes_Node
    analyzeExpRefItem(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 bool) *Nodes_Node
    analyzeExpSymbol(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token, arg3 LnsInt, arg4 LnsAny, arg5 bool, arg6 bool) *Nodes_Node
    analyzeExpUnwrap(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeExtend(_env *LnsEnv, arg1 LnsInt, arg2 Types_Position)(*Types_Token, LnsAny, *LnsList, *LnsMap, *Nodes_ClassInheritInfo)
    analyzeFor(_env *LnsEnv, arg1 *Types_Token) *Nodes_ForNode
    analyzeForeach(_env *LnsEnv, arg1 *Types_Token, arg2 bool) *Nodes_Node
    analyzeFuncBlock(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 string, arg6 *Ast_Scope, arg7 *LnsList) *Nodes_BlockNode
    analyzeIf(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeIfUnwrap(_env *LnsEnv, arg1 *Types_Token) *Nodes_IfUnwrapNode
    analyzeInitExp(_env *LnsEnv, arg1 Types_Position, arg2 LnsInt, arg3 bool, arg4 *LnsList, arg5 *LnsList)(*LnsList, *LnsList, *LnsList, LnsAny)
    analyzeLetAndInitExp(_env *LnsEnv, arg1 Types_Position, arg2 bool, arg3 LnsInt, arg4 LnsInt, arg5 bool)(*LnsList, *LnsList, *LnsList, LnsAny)
    analyzeListConst(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny) *Nodes_Node
    analyzeListItems(_env *LnsEnv, arg1 Types_Position, arg2 *Types_Token, arg3 string, arg4 LnsAny)(LnsAny, *Ast_TypeInfo)
    analyzeLuneControl(_env *LnsEnv, arg1 *Types_Token) LnsAny
    AnalyzeLuneControlToken(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token) LnsAny
    analyzeMapConst(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny) *Nodes_LiteralMapNode
    analyzeMatch(_env *LnsEnv, arg1 *Types_Token) *Nodes_MatchNode
    analyzeNewAlge(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_AlgeTypeInfo, arg3 LnsAny) *Nodes_NewAlgeValNode
    analyzePushClass(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsInt, arg9 *LnsList)(*Types_Token, *TransUnitIF_NSInfo, *Nodes_ClassInheritInfo)
    analyzeRefType(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 bool) *Nodes_RefTypeNode
    analyzeRefTypeWithSymbol(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 LnsAny, arg4 *Nodes_Node, arg5 bool) *Nodes_RefTypeNode
    analyzeRepeat(_env *LnsEnv, arg1 *Types_Token) *Nodes_RepeatNode
    analyzeRequest(_env *LnsEnv, arg1 *Types_Token) *Nodes_RequestNode
    analyzeRetTypeList(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 bool)(*LnsList, *Types_Token, *LnsList)
    analyzeReturn(_env *LnsEnv, arg1 *Types_Token) *Nodes_ReturnNode
    analyzeScope(_env *LnsEnv, arg1 *Types_Token) *Nodes_ScopeNode
    analyzeSetConst(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny) *Nodes_Node
    analyzeStatement(_env *LnsEnv, arg1 LnsAny) LnsAny
    analyzeStatementList(_env *LnsEnv, arg1 *LnsList, arg2 bool, arg3 LnsAny)(LnsAny, LnsInt)
    analyzeStatementListSubfile(_env *LnsEnv, arg1 *LnsList) LnsAny
    AnalyzeStatementToken(_env *LnsEnv, arg1 *Types_Token) LnsAny
    analyzeStrConst(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token) *Nodes_Node
    analyzeSuper(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeSwitch(_env *LnsEnv, arg1 *Types_Token) *Nodes_SwitchNode
    analyzeUnwrap(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeWhen(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeWhile(_env *LnsEnv, arg1 *Types_Token) *Nodes_WhileNode
    canBeAsyncParam(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    checkAlgeComp(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_AlgeTypeInfo)
    checkArgForSort(_env *LnsEnv, arg1 *Types_Token, arg2 *LnsList, arg3 *Nodes_ExpListNode)
    checkArgForStringForm(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_ExpListNode)
    checkAsyncField(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 Types_Position)
    checkAsyncSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 Types_Position)
    checkComp(_env *LnsEnv, arg1 *Types_Token, arg2 TransUnit_checkCompForm_31_)
    checkEnumComp(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_EnumTypeInfo)
    checkFieldComp(_env *LnsEnv, arg1 bool, arg2 *Types_Token, arg3 *Nodes_Node)
    checkImplicitCast(_env *LnsEnv, arg1 *LnsMap, arg2 bool, arg3 *LnsList, arg4 *Nodes_ExpListNode, arg5 TransUnit_checkImplicitCastCallback_10_) LnsAny
    checkLiteralEmptyCollection(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 *Ast_TypeInfo)
    checkMatchType(_env *LnsEnv, arg1 string, arg2 Types_Position, arg3 *LnsList, arg4 LnsAny, arg5 bool, arg6 bool, arg7 LnsAny, arg8 bool)(LnsInt, *LnsMap, LnsAny, *LnsList)
    checkMatchValType(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 *LnsList, arg5 LnsAny)(LnsInt, *LnsMap, LnsAny)
    checkNextToken(_env *LnsEnv, arg1 string) *Types_Token
    checkNoasyncType(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo)
    checkOverrideMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo) *LnsList
    checkOverriededMethodOfAllClass(_env *LnsEnv)
    checkPublic(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo)
    checkStringFormat(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 *LnsList)
    checkSymbol(_env *LnsEnv, arg1 *Types_Token, arg2 LnsInt) *Types_Token
    checkSymbolComp(_env *LnsEnv, arg1 *Types_Token)
    checkSymbolHavingValue(_env *LnsEnv, arg1 Types_Position, arg2 *LnsList)
    checkToken(_env *LnsEnv, arg1 *Types_Token, arg2 string) *Types_Token
    convToExtTypeList(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *LnsList) *LnsList
    createExpList(_env *LnsEnv, arg1 Types_Position, arg2 *LnsList, arg3 *LnsList, arg4 bool, arg5 LnsAny) *Nodes_ExpListNode
    createExpListNode(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 *LnsList) *Nodes_ExpListNode
    createExtType(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo) *Ast_TypeInfo
    createModifier(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt) *Ast_TypeInfo
    createNoneNode(_env *LnsEnv, arg1 Types_Position) *Nodes_Node
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    declFuncPostProcess(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsAny, arg3 *Nodes_BlockNode, arg4 *Ast_Scope)
    dumpComp(_env *LnsEnv, arg1 Writer_Writer, arg2 string, arg3 *Ast_SymbolInfo, arg4 bool) bool
    dumpFieldComp(_env *LnsEnv, arg1 Writer_Writer, arg2 bool, arg3 *Ast_TypeInfo, arg4 string, arg5 LnsAny)
    dumpSymbolComp(_env *LnsEnv, arg1 Writer_Writer, arg2 *Ast_Scope, arg3 string)
    Error(_env *LnsEnv, arg1 string)
    ErrorAt(_env *LnsEnv, arg1 Types_Position, arg2 string)
    ErrorShadowingOp(_env *LnsEnv, arg1 Types_Position, arg2 LnsAny, arg3 bool)
    evalMacro(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_TypeInfo, arg3 LnsAny) *Nodes_ExpMacroExpNode
    evalMacroOp(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 Macro_EvalMacroCallback)
    Get(_env *LnsEnv) *LnsMap
    getContinueToken(_env *LnsEnv)(*Types_Token, bool)
    GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    getDefaultAsync(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny) LnsInt
    GetLatestPos(_env *LnsEnv) Types_Position
    getMutableAsync(_env *LnsEnv, arg1 *Types_Token)(*Types_Token, LnsInt, LnsAny)
    GetNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) *TransUnitIF_NSInfo
    GetStreamName(_env *LnsEnv) string
    getSymbolToken(_env *LnsEnv, arg1 LnsInt) *Types_Token
    getToken(_env *LnsEnv, arg1 LnsAny) *Types_Token
    GetTokenNoErr(_env *LnsEnv, arg1 LnsAny) *Types_Token
    Get_curNsInfo(_env *LnsEnv) *TransUnitIF_NSInfo
    Get_errMessList(_env *LnsEnv) *LnsList
    Get_globalScope(_env *LnsEnv) *Ast_Scope
    Get_importedAliasMap(_env *LnsEnv) *LnsMap
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_scopeRO(_env *LnsEnv) *Ast_Scope
    Get_warnMessList(_env *LnsEnv) *LnsList
    mergeFrom(_env *LnsEnv, arg1 *TransUnit_TransUnit, arg2 *LnsMap)
    NewNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Types_Position) *TransUnitIF_NSInfo
    NewNSInfoWithTypeData(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Ast_TypeDataAccessor, arg3 Types_Position) *TransUnitIF_NSInfo
    NewPushback(_env *LnsEnv, arg1 *LnsList)
    PopClass(_env *LnsEnv)
    PopModule(_env *LnsEnv)
    PopScope(_env *LnsEnv)
    prepareExpCall(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 *Ast_TypeInfo)(*LnsMap, LnsAny)
    processAddFunc(_env *LnsEnv, arg1 bool, arg2 *Ast_Scope, arg3 *Types_Token, arg4 *Ast_TypeInfo, arg5 *LnsMap)(*Ast_SymbolInfo, *TransUnitIF_NSInfo)
    processCaseDefault(_env *LnsEnv, arg1 *Types_Token, arg2 LnsInt, arg3 *Types_Token, arg4 bool)(LnsAny, bool)
    processFunc(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token, arg3 LnsAny, arg4 *Nodes_Node, arg5 *Ast_TypeInfo, arg6 *LnsMap, arg7 *LnsList, arg8 *Ast_TypeInfo, arg9 LnsAny) *Nodes_Node
    processFuncBlockInfo(_env *LnsEnv, arg1 TransUnit_FuncBlockCtlIF, arg2 string) *LnsMap
    PushClass(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *TransUnitIF_NSInfo
    PushClassLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *Ast_TypeInfo
    PushClassScope(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
    PushModule(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *TransUnitIF_NSInfo
    PushModuleLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *Ast_TypeInfo
    PushScope(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny) *Ast_Scope
    Pushback(_env *LnsEnv)
    PushbackStr(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 string, arg4 Types_Position)
    PushbackToken(_env *LnsEnv, arg1 *Types_Token)
    Run(_env *LnsEnv)
    SetParser(_env *LnsEnv, arg1 *Parser_DefaultPushbackParser)
    SetScope(_env *LnsEnv, arg1 *Ast_Scope, arg2 LnsInt)
    Set_curNsInfo(_env *LnsEnv, arg1 *TransUnitIF_NSInfo)
    setup(_env *LnsEnv, arg1 *TransUnit_TransUnit)
    skipAndCreateDummyBlock(_env *LnsEnv) *Nodes_BlockNode
    skipBlock(_env *LnsEnv, arg1 bool) *LnsList
    supportLang(_env *LnsEnv, arg1 string) bool
    WaitToSetup(_env *LnsEnv)
}
type TransUnit_TransUnitRunner struct {
    TransUnit_TransUnit
    _syncFlag *Lns_syncFlag
    funcBlockCtl TransUnit_FuncBlockCtlIF
    resultMap *LnsMap
    srcTranUnit *TransUnit_TransUnit
    alreadyToSetup LnsAny
    FP TransUnit_TransUnitRunnerMtd
}
func (self *TransUnit_TransUnitRunner) GetLnsSyncFlag() *Lns_syncFlag { return self._syncFlag }
func TransUnit_TransUnitRunner2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_TransUnitRunner).FP
}
type TransUnit_TransUnitRunnerDownCast interface {
    ToTransUnit_TransUnitRunner() *TransUnit_TransUnitRunner
}
func TransUnit_TransUnitRunnerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_TransUnitRunnerDownCast)
    if ok { return work.ToTransUnit_TransUnitRunner() }
    return nil
}
func (obj *TransUnit_TransUnitRunner) ToTransUnit_TransUnitRunner() *TransUnit_TransUnitRunner {
    return obj
}
func NewTransUnit_TransUnitRunner(_env *LnsEnv, arg1 *TransUnit_TransUnit, arg2 *FrontInterface_ModuleId, arg3 *FrontInterface_ImportModuleInfo, arg4 *Nodes_MacroEval, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsAny, arg9 *LuaVer_LuaVerInfo, arg10 *Types_TransCtrlInfo, arg11 *Builtin_BuiltinFuncType, arg12 *LnsList, arg13 LnsInt) *TransUnit_TransUnitRunner {
    obj := &TransUnit_TransUnitRunner{}
    obj.FP = obj
    obj.TransUnit_TransUnit.FP = obj
    obj.TransUnitIF_TransUnitBase.FP = obj
    obj.InitTransUnit_TransUnitRunner(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
    return obj
}
// 16: DeclConstr
func (self *TransUnit_TransUnitRunner) InitTransUnit_TransUnitRunner(_env *LnsEnv, srcTranUnit *TransUnit_TransUnit,moduleId *FrontInterface_ModuleId,importModuleInfo *FrontInterface_ImportModuleInfo,macroEval *Nodes_MacroEval,enableMultiPhase bool,analyzeModule LnsAny,mode LnsAny,pos LnsAny,targetLuaVer *LuaVer_LuaVerInfo,ctrl_info *Types_TransCtrlInfo,builtinFunc *Builtin_BuiltinFuncType,list *LnsList,managerId LnsInt) {
    self._syncFlag = &Lns_syncFlag{}
    self.InitTransUnit_TransUnit(_env, moduleId, importModuleInfo, macroEval, enableMultiPhase, analyzeModule, mode, pos, targetLuaVer, ctrl_info, builtinFunc)
    self.funcBlockCtl = NewTransUnit_ListFuncBlockCtl(_env, list).FP
    self.resultMap = NewLnsMap( map[LnsAny]LnsAny{})
    self.NodeManager.FP.Set_managerId(_env, managerId)
    self.srcTranUnit = srcTranUnit
    self.alreadyToSetup = LnsCreateSyncFlag(_env)
}


// declaration Class -- TransUnitCtrl
type TransUnit_TransUnitCtrlMtd interface {
    MultiTo1(_env *LnsEnv, arg1 *Nodes_Node) *Nodes_Node
    accessSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 bool)
    addAccessor(_env *LnsEnv, arg1 *Nodes_DeclMemberNode, arg2 *LnsSet, arg3 *Ast_Scope, arg4 *Ast_TypeInfo, arg5 Ast_TypeDataAccessor)
    addDefaultConstructor(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 Ast_TypeDataAccessor, arg4 *Ast_Scope, arg5 *LnsList, arg6 *LnsSet, arg7 bool)
    AddErrMess(_env *LnsEnv, arg1 Types_Position, arg2 string)
    addFuncBlockInfoList(_env *LnsEnv, arg1 *TransUnit_FuncBlockInfo)
    AddLocalVar(_env *LnsEnv, arg1 Types_Position, arg2 bool, arg3 bool, arg4 string, arg5 *Ast_TypeInfo, arg6 LnsInt, arg7 LnsAny) *Ast_SymbolInfo
    analyzeAccessClassField(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 *Types_Token)(*Ast_TypeInfo, LnsAny, bool)
    analyzeAlias(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_AliasNode
    analyzeApply(_env *LnsEnv, arg1 *Types_Token) *Nodes_ApplyNode
    analyzeAsyncLock(_env *LnsEnv, arg1 *Types_Token, arg2 LnsInt) *Nodes_Node
    analyzeBlock(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 LnsAny, arg4 LnsAny) *Nodes_BlockNode
    analyzeClassBody(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 LnsInt, arg5 LnsAny, arg6 *Ast_TypeInfo, arg7 Ast_TypeDataAccessor, arg8 *Types_Token, arg9 LnsAny, arg10 LnsAny, arg11 LnsInt, arg12 *Types_Token, arg13 *Nodes_ClassInheritInfo)(*Nodes_DeclClassNode, *Types_Token, *LnsSet)
    analyzeDecl(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token) LnsAny
    analyzeDeclAlge(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclAlgeNode
    analyzeDeclAlternateType(_env *LnsEnv, arg1 bool, arg2 *Types_Token, arg3 LnsInt)(*Types_Token, *LnsList)
    analyzeDeclArgList(_env *LnsEnv, arg1 LnsInt, arg2 *Ast_Scope, arg3 *LnsList, arg4 bool) *Types_Token
    analyzeDeclClass(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 LnsInt) *Nodes_DeclClassNode
    analyzeDeclEnum(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclEnumNode
    analyzeDeclForm(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclFormNode
    analyzeDeclFunc(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 bool, arg4 bool, arg5 LnsInt, arg6 bool, arg7 LnsAny, arg8 *Types_Token, arg9 LnsAny) *Nodes_Node
    analyzeDeclMacro(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_DeclMacroNode
    analyzeDeclMacroSub(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token, arg3 *Types_Token, arg4 *Ast_Scope, arg5 *Ast_TypeInfo, arg6 Ast_TypeDataAccessor, arg7 *LnsList) *Nodes_DeclMacroNode
    analyzeDeclMember(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool, arg4 *Types_Token) *Nodes_DeclMemberNode
    analyzeDeclMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 bool, arg4 bool, arg5 LnsInt, arg6 bool, arg7 *Types_Token, arg8 *Types_Token) *Nodes_Node
    analyzeDeclProto(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token) *Nodes_Node
    AnalyzeDeclToken(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token) LnsAny
    analyzeDeclVar(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 *Types_Token) *Nodes_Node
    analyzeExp(_env *LnsEnv, arg1 bool, arg2 bool, arg3 bool, arg4 LnsAny, arg5 LnsAny) *Nodes_Node
    analyzeExpCall(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 *Types_Token)(*Nodes_Node, *Types_Token)
    analyzeExpCast(_env *LnsEnv, arg1 *Types_Token, arg2 string, arg3 *Nodes_Node) *Nodes_Node
    analyzeExpCont(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 bool, arg4 bool) *Nodes_Node
    analyzeExpField(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token, arg3 LnsInt, arg4 *Nodes_Node) *Nodes_Node
    analyzeExpList(_env *LnsEnv, arg1 bool, arg2 bool, arg3 bool, arg4 LnsAny, arg5 LnsAny, arg6 LnsAny) *Nodes_ExpListNode
    analyzeExpMacroStat(_env *LnsEnv, arg1 *Types_Token) *Nodes_ExpMacroStatNode
    analyzeExpOneRVal(_env *LnsEnv, arg1 bool, arg2 bool, arg3 LnsAny, arg4 LnsAny) *Nodes_Node
    analyzeExpOp2(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 LnsAny) *Nodes_Node
    analyzeExpOpEquals(_env *LnsEnv, arg1 Types_Position, arg2 *Types_Token, arg3 *Nodes_Node, arg4 *Nodes_Node)(*Nodes_Node, *Nodes_Node)
    analyzeExpOpSet(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Types_Token, arg3 *LnsList) *Nodes_Node
    analyzeExpRefItem(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_Node, arg3 bool) *Nodes_Node
    analyzeExpSymbol(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token, arg3 LnsInt, arg4 LnsAny, arg5 bool, arg6 bool) *Nodes_Node
    analyzeExpUnwrap(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeExtend(_env *LnsEnv, arg1 LnsInt, arg2 Types_Position)(*Types_Token, LnsAny, *LnsList, *LnsMap, *Nodes_ClassInheritInfo)
    analyzeFor(_env *LnsEnv, arg1 *Types_Token) *Nodes_ForNode
    analyzeForeach(_env *LnsEnv, arg1 *Types_Token, arg2 bool) *Nodes_Node
    analyzeFuncBlock(_env *LnsEnv, arg1 LnsInt, arg2 *Types_Token, arg3 LnsAny, arg4 *Ast_TypeInfo, arg5 string, arg6 *Ast_Scope, arg7 *LnsList) *Nodes_BlockNode
    analyzeIf(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeIfUnwrap(_env *LnsEnv, arg1 *Types_Token) *Nodes_IfUnwrapNode
    analyzeImport(_env *LnsEnv, arg1 *Types_Token) *Nodes_ImportNode
    analyzeImportFor(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 string, arg4 bool, arg5 LnsInt) *Nodes_ImportNode
    analyzeInitExp(_env *LnsEnv, arg1 Types_Position, arg2 LnsInt, arg3 bool, arg4 *LnsList, arg5 *LnsList)(*LnsList, *LnsList, *LnsList, LnsAny)
    analyzeLetAndInitExp(_env *LnsEnv, arg1 Types_Position, arg2 bool, arg3 LnsInt, arg4 LnsInt, arg5 bool)(*LnsList, *LnsList, *LnsList, LnsAny)
    analyzeListConst(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny) *Nodes_Node
    analyzeListItems(_env *LnsEnv, arg1 Types_Position, arg2 *Types_Token, arg3 string, arg4 LnsAny)(LnsAny, *Ast_TypeInfo)
    analyzeLuneControl(_env *LnsEnv, arg1 *Types_Token) LnsAny
    AnalyzeLuneControlToken(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token) LnsAny
    analyzeMapConst(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny) *Nodes_LiteralMapNode
    analyzeMatch(_env *LnsEnv, arg1 *Types_Token) *Nodes_MatchNode
    analyzeNewAlge(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_AlgeTypeInfo, arg3 LnsAny) *Nodes_NewAlgeValNode
    analyzeProvide(_env *LnsEnv, arg1 *Types_Token) *Nodes_ProvideNode
    analyzePushClass(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 *Types_Token, arg4 *Types_Token, arg5 bool, arg6 LnsAny, arg7 LnsAny, arg8 LnsInt, arg9 *LnsList)(*Types_Token, *TransUnitIF_NSInfo, *Nodes_ClassInheritInfo)
    analyzeRefType(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 bool) *Nodes_RefTypeNode
    analyzeRefTypeWithSymbol(_env *LnsEnv, arg1 LnsInt, arg2 bool, arg3 LnsAny, arg4 *Nodes_Node, arg5 bool) *Nodes_RefTypeNode
    analyzeRepeat(_env *LnsEnv, arg1 *Types_Token) *Nodes_RepeatNode
    analyzeRequest(_env *LnsEnv, arg1 *Types_Token) *Nodes_RequestNode
    analyzeRetTypeList(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Types_Token, arg4 bool)(*LnsList, *Types_Token, *LnsList)
    analyzeReturn(_env *LnsEnv, arg1 *Types_Token) *Nodes_ReturnNode
    analyzeScope(_env *LnsEnv, arg1 *Types_Token) *Nodes_ScopeNode
    analyzeSetConst(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny) *Nodes_Node
    analyzeStatement(_env *LnsEnv, arg1 LnsAny) LnsAny
    analyzeStatementList(_env *LnsEnv, arg1 *LnsList, arg2 bool, arg3 LnsAny)(LnsAny, LnsInt)
    analyzeStatementListSubfile(_env *LnsEnv, arg1 *LnsList) LnsAny
    AnalyzeStatementToken(_env *LnsEnv, arg1 *Types_Token) LnsAny
    analyzeStrConst(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token) *Nodes_Node
    analyzeSubfile(_env *LnsEnv, arg1 *Types_Token) *Nodes_SubfileNode
    analyzeSuper(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeSwitch(_env *LnsEnv, arg1 *Types_Token) *Nodes_SwitchNode
    analyzeTest(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeTestCase(_env *LnsEnv, arg1 *Types_Token) *Nodes_TestCaseNode
    analyzeUnwrap(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeWhen(_env *LnsEnv, arg1 *Types_Token) *Nodes_Node
    analyzeWhile(_env *LnsEnv, arg1 *Types_Token) *Nodes_WhileNode
    canBeAsyncParam(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    checkAlgeComp(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_AlgeTypeInfo)
    checkArgForSort(_env *LnsEnv, arg1 *Types_Token, arg2 *LnsList, arg3 *Nodes_ExpListNode)
    checkArgForStringForm(_env *LnsEnv, arg1 *Types_Token, arg2 *Nodes_ExpListNode)
    checkAsyncField(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 Types_Position)
    checkAsyncSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 Types_Position)
    checkComp(_env *LnsEnv, arg1 *Types_Token, arg2 TransUnit_checkCompForm_31_)
    checkEnumComp(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_EnumTypeInfo)
    checkFieldComp(_env *LnsEnv, arg1 bool, arg2 *Types_Token, arg3 *Nodes_Node)
    checkImplicitCast(_env *LnsEnv, arg1 *LnsMap, arg2 bool, arg3 *LnsList, arg4 *Nodes_ExpListNode, arg5 TransUnit_checkImplicitCastCallback_10_) LnsAny
    checkLiteralEmptyCollection(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 *Ast_TypeInfo)
    checkMatchType(_env *LnsEnv, arg1 string, arg2 Types_Position, arg3 *LnsList, arg4 LnsAny, arg5 bool, arg6 bool, arg7 LnsAny, arg8 bool)(LnsInt, *LnsMap, LnsAny, *LnsList)
    checkMatchValType(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 *LnsList, arg5 LnsAny)(LnsInt, *LnsMap, LnsAny)
    checkNextToken(_env *LnsEnv, arg1 string) *Types_Token
    checkNoasyncType(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo)
    checkOverrideMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo) *LnsList
    checkOverriededMethodOfAllClass(_env *LnsEnv)
    checkPublic(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo)
    checkStringFormat(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 *LnsList)
    checkSymbol(_env *LnsEnv, arg1 *Types_Token, arg2 LnsInt) *Types_Token
    checkSymbolComp(_env *LnsEnv, arg1 *Types_Token)
    checkSymbolHavingValue(_env *LnsEnv, arg1 Types_Position, arg2 *LnsList)
    checkToken(_env *LnsEnv, arg1 *Types_Token, arg2 string) *Types_Token
    convToExtTypeList(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *LnsList) *LnsList
    CreateAST(_env *LnsEnv, arg1 LnsAny, arg2 bool, arg3 LnsAny, arg4 LnsAny, arg5 bool, arg6 LnsAny, arg7 LnsAny) *AstInfo_ASTInfo
    createExpList(_env *LnsEnv, arg1 Types_Position, arg2 *LnsList, arg3 *LnsList, arg4 bool, arg5 LnsAny) *Nodes_ExpListNode
    createExpListNode(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 *LnsList) *Nodes_ExpListNode
    createExtType(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo) *Ast_TypeInfo
    createModifier(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt) *Ast_TypeInfo
    createNoneNode(_env *LnsEnv, arg1 Types_Position) *Nodes_Node
    CreatePosition(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) Types_Position
    declFuncPostProcess(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsAny, arg3 *Nodes_BlockNode, arg4 *Ast_Scope)
    dumpComp(_env *LnsEnv, arg1 Writer_Writer, arg2 string, arg3 *Ast_SymbolInfo, arg4 bool) bool
    dumpFieldComp(_env *LnsEnv, arg1 Writer_Writer, arg2 bool, arg3 *Ast_TypeInfo, arg4 string, arg5 LnsAny)
    dumpSymbolComp(_env *LnsEnv, arg1 Writer_Writer, arg2 *Ast_Scope, arg3 string)
    Error(_env *LnsEnv, arg1 string)
    ErrorAt(_env *LnsEnv, arg1 Types_Position, arg2 string)
    ErrorShadowingOp(_env *LnsEnv, arg1 Types_Position, arg2 LnsAny, arg3 bool)
    evalMacro(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_TypeInfo, arg3 LnsAny) *Nodes_ExpMacroExpNode
    evalMacroOp(_env *LnsEnv, arg1 *Types_Token, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 Macro_EvalMacroCallback)
    getContinueToken(_env *LnsEnv)(*Types_Token, bool)
    GetCurrentNamespaceTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    getDefaultAsync(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny) LnsInt
    GetLatestPos(_env *LnsEnv) Types_Position
    getMutableAsync(_env *LnsEnv, arg1 *Types_Token)(*Types_Token, LnsInt, LnsAny)
    GetNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) *TransUnitIF_NSInfo
    GetStreamName(_env *LnsEnv) string
    getSymbolToken(_env *LnsEnv, arg1 LnsInt) *Types_Token
    getToken(_env *LnsEnv, arg1 LnsAny) *Types_Token
    GetTokenNoErr(_env *LnsEnv, arg1 LnsAny) *Types_Token
    Get_curNsInfo(_env *LnsEnv) *TransUnitIF_NSInfo
    Get_errMessList(_env *LnsEnv) *LnsList
    Get_globalScope(_env *LnsEnv) *Ast_Scope
    Get_importedAliasMap(_env *LnsEnv) *LnsMap
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_scopeRO(_env *LnsEnv) *Ast_Scope
    Get_warnMessList(_env *LnsEnv) *LnsList
    mergeFrom(_env *LnsEnv, arg1 *TransUnit_TransUnit, arg2 *LnsMap)
    NewNSInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Types_Position) *TransUnitIF_NSInfo
    NewNSInfoWithTypeData(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 Ast_TypeDataAccessor, arg3 Types_Position) *TransUnitIF_NSInfo
    NewPushback(_env *LnsEnv, arg1 *LnsList)
    PopClass(_env *LnsEnv)
    PopModule(_env *LnsEnv)
    PopScope(_env *LnsEnv)
    prepareExpCall(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *LnsList, arg4 *Ast_TypeInfo)(*LnsMap, LnsAny)
    processAddFunc(_env *LnsEnv, arg1 bool, arg2 *Ast_Scope, arg3 *Types_Token, arg4 *Ast_TypeInfo, arg5 *LnsMap)(*Ast_SymbolInfo, *TransUnitIF_NSInfo)
    processCaseDefault(_env *LnsEnv, arg1 *Types_Token, arg2 LnsInt, arg3 *Types_Token, arg4 bool)(LnsAny, bool)
    processFunc(_env *LnsEnv, arg1 *Types_Token, arg2 *Types_Token, arg3 LnsAny, arg4 *Nodes_Node, arg5 *Ast_TypeInfo, arg6 *LnsMap, arg7 *LnsList, arg8 *Ast_TypeInfo, arg9 LnsAny) *Nodes_Node
    processFuncBlock(_env *LnsEnv, arg1 string)
    processFuncBlockInfo(_env *LnsEnv, arg1 TransUnit_FuncBlockCtlIF, arg2 string) *LnsMap
    PushClass(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *TransUnitIF_NSInfo
    PushClassLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 Types_Position, arg3 LnsInt, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 bool, arg9 string, arg10 bool, arg11 LnsInt, arg12 LnsAny) *Ast_TypeInfo
    PushClassScope(_env *LnsEnv, arg1 Types_Position, arg2 *Ast_TypeInfo, arg3 *Ast_Scope)
    PushModule(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *TransUnitIF_NSInfo
    PushModuleLow(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 bool, arg3 string, arg4 bool) *Ast_TypeInfo
    PushScope(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny) *Ast_Scope
    Pushback(_env *LnsEnv)
    PushbackStr(_env *LnsEnv, arg1 LnsAny, arg2 string, arg3 string, arg4 Types_Position)
    PushbackToken(_env *LnsEnv, arg1 *Types_Token)
    SetParser(_env *LnsEnv, arg1 *Parser_DefaultPushbackParser)
    SetScope(_env *LnsEnv, arg1 *Ast_Scope, arg2 LnsInt)
    Set_curNsInfo(_env *LnsEnv, arg1 *TransUnitIF_NSInfo)
    setup(_env *LnsEnv, arg1 *TransUnit_TransUnit)
    skipAndCreateDummyBlock(_env *LnsEnv) *Nodes_BlockNode
    skipBlock(_env *LnsEnv, arg1 bool) *LnsList
    supportLang(_env *LnsEnv, arg1 string) bool
}
type TransUnit_TransUnitCtrl struct {
    TransUnit_TransUnit
    macroEval *Nodes_MacroEval
    importCtrl LnsAny
    importModuleInfo *FrontInterface_ImportModuleInfo
    stdinFile LnsAny
    funcBlockInfoList *LnsList
    totalFuncBlockTokenNum LnsInt
    subfileList *LnsList
    provideNode LnsAny
    FP TransUnit_TransUnitCtrlMtd
}
func TransUnit_TransUnitCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*TransUnit_TransUnitCtrl).FP
}
type TransUnit_TransUnitCtrlDownCast interface {
    ToTransUnit_TransUnitCtrl() *TransUnit_TransUnitCtrl
}
func TransUnit_TransUnitCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(TransUnit_TransUnitCtrlDownCast)
    if ok { return work.ToTransUnit_TransUnitCtrl() }
    return nil
}
func (obj *TransUnit_TransUnitCtrl) ToTransUnit_TransUnitCtrl() *TransUnit_TransUnitCtrl {
    return obj
}
func NewTransUnit_TransUnitCtrl(_env *LnsEnv, arg1 *FrontInterface_ModuleId, arg2 *FrontInterface_ImportModuleInfo, arg3 *Nodes_MacroEval, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 LnsAny, arg8 *LuaVer_LuaVerInfo, arg9 *Types_TransCtrlInfo, arg10 *Builtin_BuiltinFuncType) *TransUnit_TransUnitCtrl {
    obj := &TransUnit_TransUnitCtrl{}
    obj.FP = obj
    obj.TransUnit_TransUnit.FP = obj
    obj.TransUnitIF_TransUnitBase.FP = obj
    obj.InitTransUnit_TransUnitCtrl(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
// 102: DeclConstr
func (self *TransUnit_TransUnitCtrl) InitTransUnit_TransUnitCtrl(_env *LnsEnv, moduleId *FrontInterface_ModuleId,importModuleInfo *FrontInterface_ImportModuleInfo,macroEval *Nodes_MacroEval,enableMultiPhase bool,analyzeModule LnsAny,mode LnsAny,pos LnsAny,targetLuaVer *LuaVer_LuaVerInfo,ctrl_info *Types_TransCtrlInfo,builtinFunc *Builtin_BuiltinFuncType) {
    self.InitTransUnit_TransUnit(_env, moduleId, importModuleInfo, macroEval, enableMultiPhase, analyzeModule, mode, pos, targetLuaVer, ctrl_info, builtinFunc)
    self.totalFuncBlockTokenNum = 0
    self.macroEval = macroEval
    self.importModuleInfo = importModuleInfo.FP.Clone(_env)
    self.importCtrl = nil
    self.stdinFile = nil
    self.subfileList = NewLnsList([]LnsAny{})
    self.provideNode = nil
    self.funcBlockInfoList = NewLnsList([]LnsAny{})
}



func Lns_TransUnit_init(_env *LnsEnv) {
    if init_TransUnit { return }
    init_TransUnit = true
    TransUnit__mod__ = "@lune.@base.@TransUnit"
    Lns_InitMod()
    Lns_Types_init(_env)
    Lns_Meta_init(_env)
    Lns_Parser_init(_env)
    Lns_Util_init(_env)
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Writer_init(_env)
    Lns_frontInterface_init(_env)
    Lns_LuaVer_init(_env)
    Lns_Option_init(_env)
    Lns_Code_init(_env)
    Lns_Log_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Macro_init(_env)
    Lns_TransUnitIF_init(_env)
    Lns_Builtin_init(_env)
    Lns_Import_init(_env)
    Lns_AstInfo_init(_env)
    TransUnit_opTopLevel = 0
    TransUnit_op2levelMap,TransUnit_op1levelMap = TransUnit_setupOpLevel_21_(_env)
    TransUnit_quotedChar2Code = NewLnsMap( map[LnsAny]LnsAny{"a":7,"b":8,"t":9,"n":10,"v":11,"f":12,"r":13,"\\":92,"\"":34,"'":39,})
    TransUnit_specialSymbolSet = NewLnsSet([]LnsAny{"__init", "__free", "__", "_exp"})
    TransUnit_builtinKeywordSet = NewLnsSet([]LnsAny{"self", "super"})
    TransUnit_CantOverrideMethods = NewLnsSet([]LnsAny{"__init", "__free"})
    TransUnit_unsupportStatement = NewLnsSet([]LnsAny{"import", "subfile", "provide"})
}
func init() {
    init_TransUnit = false
}
// 199: decl @lune.@base.@TransUnit.TentativeSymbol.modSym
func (self *TransUnit_TentativeSymbol) ModSym(_env *LnsEnv, accessSym *Ast_SymbolInfo) {
    self.modSymbolSet.Add(Ast_SymbolInfo2Stem(accessSym.FP.GetOrg(_env)))
}
// 202: decl @lune.@base.@TransUnit.TentativeSymbol.addAccessSym
func (self *TransUnit_TentativeSymbol) AddAccessSym(_env *LnsEnv, accessSym *Ast_SymbolInfo) {
    self.accessSymSet.Add(Ast_SymbolInfo2Stem(accessSym.FP.GetOrg(_env)))
}
// 211: decl @lune.@base.@TransUnit.TentativeSymbol.addAccessSymPos
func (self *TransUnit_TentativeSymbol) AddAccessSymPos(_env *LnsEnv, accessSym *TransUnit_AccessSymPos) {
    self.accessSymList.Insert(TransUnit_AccessSymPos2Stem(accessSym))
}
// 217: decl @lune.@base.@TransUnit.TentativeSymbol.clearAccessSym
func (self *TransUnit_TentativeSymbol) ClearAccessSym(_env *LnsEnv) {
    if self.accessSymList.Len() != 0{
        self.accessSymList = NewLnsList([]LnsAny{})
    }
}
// 229: decl @lune.@base.@TransUnit.TentativeSymbol.checkAndExclude
func (self *TransUnit_TentativeSymbol) CheckAndExclude(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) bool {
    symbolInfo = symbolInfo.FP.GetOrg(_env)
    if self.symbolSet.Has(Ast_SymbolInfo2Stem(symbolInfo)){
        self.symbolSet.Del(Ast_SymbolInfo2Stem(symbolInfo))
        return true
    }
    return false
}
// 241: decl @lune.@base.@TransUnit.TentativeSymbol.regist
func (self *TransUnit_TentativeSymbol) Regist(_env *LnsEnv, symbolInfo *Ast_SymbolInfo,pos Types_Position) bool {
    self.symbolSet.Add(Ast_SymbolInfo2Stem(symbolInfo.FP.GetOrg(_env)))
    self.initSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo.FP.GetOrg(_env)))
    symbolInfo.FP.Set_hasValueFlag(_env, true)
    if self.FP.Get_scope(_env).FP.IsInnerOf(_env, symbolInfo.FP.Get_scope(_env)){
        if Lns_op_not(symbolInfo.FP.Get_mutable(_env)){
            var work *TransUnit_TentativeSymbol
            work = self
            for  {
                {
                    __exp := work.parent
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*TransUnit_TentativeSymbol)
                        if work.scope == symbolInfo.FP.Get_scope(_env){
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
// 275: decl @lune.@base.@TransUnit.TentativeSymbol.skip
func (self *TransUnit_TentativeSymbol) Skip(_env *LnsEnv) {
    for _symbolInfo := range( self.symbolSet.Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        symbolInfo.FP.ClearValue(_env)
    }
    self.skipFlag = true
}
// 288: decl @lune.@base.@TransUnit.TentativeSymbol.merge
func (self *TransUnit_TentativeSymbol) merge(_env *LnsEnv, finishFlag bool) bool {
    if self.skipFlag{
        self.skipFlag = false
        {
            _other := self.oldSymbolSet
            if !Lns_IsNil( _other ) {
                other := _other.(*LnsSet)
                self.symbolSet = other.Clone()
            }
        }
        if finishFlag{
            for _symbolInfo := range( self.symbolSet.Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                symbolInfo.FP.UpdateValue(_env, symbolInfo.FP.Get_posForLatestMod(_env))
            }
        }
        return self.oldSymbolSet != nil
    }
    {
        _other := self.oldSymbolSet
        if !Lns_IsNil( _other ) {
            other := _other.(*LnsSet)
            var mergedSet *LnsSet
            mergedSet = self.symbolSet.Clone().And(other)
            if finishFlag{
                for _symbolInfo := range( self.symbolSet.Clone().Or(other).Sub(mergedSet).Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.ClearValue(_env)
                }
            } else { 
                for _symbolInfo := range( self.symbolSet.Clone().Or(other).Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.ClearValue(_env)
                }
            }
            self.symbolSet = mergedSet
        } else {
            if Lns_op_not(finishFlag){
                for _symbolInfo := range( self.symbolSet.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.ClearValue(_env)
                }
            }
        }
    }
    return true
}
// 325: decl @lune.@base.@TransUnit.TentativeSymbol.syncPos
func (self *TransUnit_TentativeSymbol) SyncPos(_env *LnsEnv) {
    if self.loopFlag{
        for _, _info := range( self.sym2posForModToRef.Items ) {
            info := _info.(TransUnit_RefAccessSymPosDownCast).ToTransUnit_RefAccessSymPos()
            var symbol *Ast_SymbolInfo
            symbol = info.FP.Get_symbol(_env)
            if Lns_isCondTrue( symbol.FP.Get_posForModToRef(_env)){
                symbol.FP.Set_posForModToRef(_env, symbol.FP.Get_posForLatestMod(_env))
            } else { 
                symbol.FP.Set_posForModToRef(_env, info.FP.Get_pos(_env))
            }
        }
    }
}
// 350: decl @lune.@base.@TransUnit.TentativeSymbol.finish
func (self *TransUnit_TentativeSymbol) Finish(_env *LnsEnv, complete bool) LnsAny {
    self.FP.SyncPos(_env)
    self.FP.merge(_env, true)
    {
        _parent := self.parent
        if !Lns_IsNil( _parent ) {
            parent := _parent.(*TransUnit_TentativeSymbol)
            if complete{
                for _symbolInfo := range( self.symbolSet.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    if symbolInfo.FP.Get_hasValueFlag(_env){
                        if parent.scope.FP.IsInnerOf(_env, symbolInfo.FP.Get_scope(_env)){
                            parent.symbolSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
                        }
                    }
                }
            } else { 
                for _symbolInfo := range( self.symbolSet.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.ClearValue(_env)
                }
            }
            for _symbolInfo := range( self.initSymSet.Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if symbolInfo.FP.Get_scope(_env) != self.FP.Get_scope(_env){
                    parent.initSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
                }
            }
            
            for _symbolInfo := range( self.accessSymSet.Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if symbolInfo.FP.Get_scope(_env) != self.FP.Get_scope(_env){
                    parent.accessSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
                }
            }
            
            for _symbolInfo := range( self.modSymbolSet.Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if symbolInfo.FP.Get_scope(_env) != self.FP.Get_scope(_env){
                    parent.modSymbolSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
                }
            }
            
            return parent
        }
    }
    return nil
}
// 386: decl @lune.@base.@TransUnit.TentativeSymbol.newSet
func (self *TransUnit_TentativeSymbol) NewSet(_env *LnsEnv, scope *Ast_Scope) {
    if self.FP.merge(_env, false){
        self.oldSymbolSet = self.symbolSet
    }
    self.symbolSet = NewLnsSet([]LnsAny{})
    self.scope = scope
}
// 428: decl @lune.@base.@TransUnit.ClosureFun.check
func (self *TransUnit_ClosureFun) Check(_env *LnsEnv) bool {
    {
        __exp := _env.NilAccFin(_env.NilAccPush(self.symbol.FP.Get_typeInfo(_env).FP.Get_scope(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_closureSymList(_env)}))
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*LnsList)
            if _exp.Len() > 0{
                self.fromScope.FP.SetClosure(_env, self.symbol)
                return true
            }
        }
    }
    return false
}
// 441: decl @lune.@base.@TransUnit.ClosureFun.checkList
func TransUnit_ClosureFun_checkList_1_(_env *LnsEnv, list *LnsList) {
    var workList *LnsList
    workList = list
    var remainList *LnsList
    remainList = NewLnsList([]LnsAny{})
    for  {
        var update bool
        update = false
        for _, _closureFun := range( workList.Items ) {
            closureFun := _closureFun.(TransUnit_ClosureFunDownCast).ToTransUnit_ClosureFun()
            if closureFun.FP.Check(_env){
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
// 482: decl @lune.@base.@TransUnit.AccessSymbolSet.add
func (self *TransUnit_AccessSymbolSet) Add(_env *LnsEnv, symbol *Ast_SymbolInfo) {
    self.accessSym2Pos.Set(symbol,symbol.FP.Get_posForLatestMod(_env))
}
// 485: decl @lune.@base.@TransUnit.AccessSymbolSet.applyPos
func (self *TransUnit_AccessSymbolSet) ApplyPos(_env *LnsEnv, excludeSymList LnsAny) {
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
        if Lns_op_not(set.Has(Ast_SymbolInfo2Stem(symbol.FP.GetOrg(_env)))){
            symbol.FP.Set_posForModToRef(_env, pos)
        }
    }
}
// 496: decl @lune.@base.@TransUnit.AccessSymbolSet.clone
func (self *TransUnit_AccessSymbolSet) Clone(_env *LnsEnv) *TransUnit_AccessSymbolSet {
    var obj *TransUnit_AccessSymbolSet
    obj = NewTransUnit_AccessSymbolSet(_env)
    for _key, _val := range( self.accessSym2Pos.Items ) {
        key := _key.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        val := _val.(Types_PositionDownCast).ToTypes_Position()
        obj.accessSym2Pos.Set(key,val)
    }
    return obj
}
// 510: decl @lune.@base.@TransUnit.AccessSymbolSetQueue.push
func (self *TransUnit_AccessSymbolSetQueue) Push(_env *LnsEnv) {
    self.queue.Insert(TransUnit_AccessSymbolSet2Stem(NewTransUnit_AccessSymbolSet(_env)))
}
// 513: decl @lune.@base.@TransUnit.AccessSymbolSetQueue.pop
func (self *TransUnit_AccessSymbolSetQueue) Pop(_env *LnsEnv, symbolList LnsAny) {
    self.queue.GetAt(self.queue.Len()).(TransUnit_AccessSymbolSetDownCast).ToTransUnit_AccessSymbolSet().FP.ApplyPos(_env, symbolList)
    self.queue.Remove(nil)
}
// 517: decl @lune.@base.@TransUnit.AccessSymbolSetQueue.add
func (self *TransUnit_AccessSymbolSetQueue) Add(_env *LnsEnv, symbol *Ast_SymbolInfo) {
    self.queue.GetAt(self.queue.Len()).(TransUnit_AccessSymbolSetDownCast).ToTransUnit_AccessSymbolSet().FP.Add(_env, symbol)
}
// 520: decl @lune.@base.@TransUnit.AccessSymbolSetQueue.getMap
func (self *TransUnit_AccessSymbolSetQueue) GetMap(_env *LnsEnv) *LnsMap {
    return self.queue.GetAt(self.queue.Len()).(TransUnit_AccessSymbolSetDownCast).ToTransUnit_AccessSymbolSet().FP.Get_accessSym2Pos(_env)
}
// 524: decl @lune.@base.@TransUnit.AccessSymbolSetQueue.setupFrom
func (self *TransUnit_AccessSymbolSetQueue) setupFrom(_env *LnsEnv, src *TransUnit_AccessSymbolSetQueue) {
    for _, _symbol := range( src.queue.Items ) {
        symbol := _symbol.(TransUnit_AccessSymbolSetDownCast).ToTransUnit_AccessSymbolSet()
        self.queue.Insert(TransUnit_AccessSymbolSet2Stem(symbol.FP.Clone(_env)))
    }
}
// 555: decl @lune.@base.@TransUnit.ListFuncBlockCtl.getNext
func (self *TransUnit_ListFuncBlockCtl) GetNext(_env *LnsEnv) LnsAny {
    if self.pos >= self.list.Len(){
        return nil
    }
    self.pos = self.pos + 1
    return self.list.GetAt(self.pos).(TransUnit_FuncBlockInfoDownCast).ToTransUnit_FuncBlockInfo()
}
// 571: decl @lune.@base.@TransUnit.PipeFuncBlockCtl.getNext
func (self *TransUnit_PipeFuncBlockCtl) GetNext(_env *LnsEnv) LnsAny {
    return self.pipe.Get(_env)
}
// 649: decl @lune.@base.@TransUnit.TransUnit.setParser
func (self *TransUnit_TransUnit) SetParser(_env *LnsEnv, parser *Parser_DefaultPushbackParser) {
    self.Parser = parser
}
// 719: decl @lune.@base.@TransUnit.TransUnit.setup
func (self *TransUnit_TransUnit) setup(_env *LnsEnv, src *TransUnit_TransUnit) {
    {
        for _key, _val := range( src.TypeId2ClassMap.Items ) {
            key := _key.(Ast_IdInfoDownCast).ToAst_IdInfo()
            val := _val.(Nodes_NamespaceInfoDownCast).ToNodes_NamespaceInfo()
            self.TypeId2ClassMap.Set(key,val)
        }
    }
    
    for _, _val := range( src.ErrMessList.Items ) {
        val := _val.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
        self.ErrMessList.Insert(TransUnitIF_ErrMess2Stem(val))
    }
    self.GlobalScope = &NewAst_ScopeWithRef(_env, self.ProcessInfo, self.GlobalScope, src.GlobalScope, Ast_ScopeKind__Other, nil, nil).Ast_Scope
    {
        for _key, _val := range( src.TypeId2ClassMap.Items ) {
            key := _key.(Ast_IdInfoDownCast).ToAst_IdInfo()
            val := _val.(Nodes_NamespaceInfoDownCast).ToNodes_NamespaceInfo()
            self.TypeId2ClassMap.Set(key,val)
        }
    }
    
    for _typeInfo, _nsInfo := range( src.NsInfoMap.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        nsInfo := _nsInfo.(TransUnitIF_NSInfoDownCast).ToTransUnitIF_NSInfo()
        self.NsInfoMap.Set(typeInfo,nsInfo.FP.Duplicate(_env))
    }
    self.defaultAsyncMode = src.defaultAsyncMode
    self.ValidMutControl = src.ValidMutControl
    self.moduleName = src.moduleName
    self.moduleType = src.moduleType
    self.IgnoreToCheckSymbol_ = src.IgnoreToCheckSymbol_
    self.BaseDir = src.BaseDir
    self.analyzePhase = src.analyzePhase
    self.TypeNameCtrl = NewAst_TypeNameCtrl(_env, self.moduleType)
    self.MacroCtrl = src.MacroCtrl.FP.Clone(_env)
    self.advertisedTypeSet.Or(src.advertisedTypeSet)
    self.accessSymbolSetQueue.FP.setupFrom(_env, src.accessSymbolSetQueue)
    self.HelperInfo.PragmaSet.Or(src.HelperInfo.PragmaSet)
    for _, _mess := range( src.WarnMessList.Items ) {
        mess := _mess.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
        self.WarnMessList.Insert(TransUnitIF_ErrMess2Stem(mess))
    }
    self.ImportModuleSet.Or(src.ImportModuleSet)
    {
        for _key, _val := range( src.ImportedAliasMap.Items ) {
            key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            val := _val.(Ast_AliasTypeInfoDownCast).ToAst_AliasTypeInfo()
            self.ImportedAliasMap.Set(key,val)
        }
    }
    
}
// 759: decl @lune.@base.@TransUnit.TransUnit.mergeFrom
func (self *TransUnit_TransUnit) mergeFrom(_env *LnsEnv, src *TransUnit_TransUnit,funcBlockResultMap *LnsMap) {
    self.NodeManager.FP.AddFrom(_env, src.NodeManager)
    for _, _val := range( src.closureFunList.Items ) {
        val := _val.(TransUnit_ClosureFunDownCast).ToTransUnit_ClosureFun()
        self.closureFunList.Insert(TransUnit_ClosureFun2Stem(val))
    }
    self.MacroCtrl.FP.MergeFrom(_env, src.MacroCtrl)
    {
        for _key, _val := range( src.typeInfo2ClassNode.Items ) {
            key := _key.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            val := _val.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            self.typeInfo2ClassNode.Set(key,val)
        }
    }
    
    self.HelperInfo.FP.MergeFrom(_env, src.HelperInfo)
    TransUnit_mergeFrom__cloneMessList_0_(_env, self.WarnMessList, src.WarnMessList)
    TransUnit_mergeFrom__cloneMessList_0_(_env, self.ErrMessList, src.ErrMessList)
    for _typeInfo, _nsInfo := range( src.NsInfoMap.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        nsInfo := _nsInfo.(TransUnitIF_NSInfoDownCast).ToTransUnitIF_NSInfo()
        var dstInfo *TransUnitIF_NSInfo
        
        {
            _dstInfo := self.NsInfoMap.Get(typeInfo)
            if _dstInfo == nil{
                dstInfo = self.FP.NewNSInfoWithTypeData(_env, nsInfo.FP.Get_typeInfo(_env), nsInfo.FP.Get_typeDataAccessor(_env), nsInfo.FP.Get_pos(_env))
                self.NsInfoMap.Set(typeInfo,dstInfo)
            } else {
                dstInfo = _dstInfo.(*TransUnitIF_NSInfo)
            }
        }
        var dstChildren *LnsList
        dstChildren = dstInfo.FP.Get_typeDataAccessor(_env).Get_typeData(_env).FP.Get_children(_env)
        var srcChildren *LnsList
        srcChildren = nsInfo.FP.Get_typeDataAccessor(_env).Get_typeData(_env).FP.Get_children(_env)
        if dstChildren.Len() == 0{
            for _, _child := range( srcChildren.Items ) {
                child := _child.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                dstChildren.Insert(Ast_TypeInfo2Stem(child))
            }
        } else if dstChildren.Len() < srcChildren.Len(){
            {
                var _forFrom0 LnsInt = dstChildren.Len()
                var _forTo0 LnsInt = srcChildren.Len()
                for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                    index := _forWork0
                    dstChildren.Insert(Ast_TypeInfo2Stem(srcChildren.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))
                }
            }
        }
    }
}
// 819: decl @lune.@base.@TransUnit.TransUnit.getSuperParam
func TransUnit_TransUnit_getSuperParam_5_(_env *LnsEnv, ctrl_info *Types_TransCtrlInfo)(*Types_TransCtrlInfo, *Ast_ProcessInfo) {
    var processInfo *Ast_ProcessInfo
    processInfo = Ast_createProcessInfo(_env, ctrl_info.ValidCheckingMutable, ctrl_info.ValidLuaval, ctrl_info.ValidAstDetailError)
    return ctrl_info, processInfo
}
// 903: decl @lune.@base.@TransUnit.TransUnit.getLatestPos
func (self *TransUnit_TransUnit) GetLatestPos(_env *LnsEnv) Types_Position {
    return self.Parser.FP.GetLastPos(_env)
}
// 907: decl @lune.@base.@TransUnit.TransUnit.pushAnalyzingState
func (self *TransUnit_TransUnit) pushAnalyzingState(_env *LnsEnv, state LnsInt) {
    self.analyzingStateQueue.Insert(state)
}
// 911: decl @lune.@base.@TransUnit.TransUnit.popAnalyzingState
func (self *TransUnit_TransUnit) popAnalyzingState(_env *LnsEnv) {
    if self.analyzingStateQueue.Len() == 0{
        self.FP.Error(_env, "underflow analyzingStateQueue")
    }
    self.analyzingStateQueue.Remove(nil)
}
// 918: decl @lune.@base.@TransUnit.TransUnit.inAnalyzingState
func (self *TransUnit_TransUnit) inAnalyzingState(_env *LnsEnv, state LnsInt) bool {
    if self.analyzingStateQueue.Len() > 0{
        return self.analyzingStateQueue.GetAt(self.analyzingStateQueue.Len()).(LnsInt) == state
    }
    return false
}
// 925: decl @lune.@base.@TransUnit.TransUnit.addWarnMess
func (self *TransUnit_TransUnit) addWarnMess(_env *LnsEnv, pos Types_Position,mess string) {
    self.WarnMessList.Insert(TransUnitIF_ErrMess2Stem(NewTransUnitIF_ErrMess(_env, _env.GetVM().String_format("%s: warning: %s", []LnsAny{pos.GetDisplayTxt(_env), mess}), pos)))
}
// 930: decl @lune.@base.@TransUnit.TransUnit.addWarnErrMess
func (self *TransUnit_TransUnit) addWarnErrMess(_env *LnsEnv, pos Types_Position,err bool,mess string) {
    if err{
        self.FP.AddErrMess(_env, pos, mess)
    } else { 
        self.FP.addWarnMess(_env, pos, mess)
    }
}
// 940: decl @lune.@base.@TransUnit.TransUnit.createDummyNS
func (self *TransUnit_TransUnit) createDummyNS(_env *LnsEnv, scope *Ast_Scope,pos Types_Position,asyncMode LnsInt) {
    var dummyType *Ast_NormalTypeInfo
    dummyType = self.ProcessInfo.FP.CreateDummyNameSpace(_env, scope, Ast_headTypeInfo, asyncMode)
    self.FP.NewNSInfo(_env, &dummyType.Ast_TypeInfo, pos)
}
// 948: decl @lune.@base.@TransUnit.TransUnit.prepareTentativeSymbol
func (self *TransUnit_TransUnit) prepareTentativeSymbol(_env *LnsEnv, scope *Ast_Scope,loopFlag bool,list LnsAny) {
    self.tentativeSymbol = NewTransUnit_TentativeSymbol(_env, self.tentativeSymbol, scope, self.ModuleScope, loopFlag, list)
}
// 955: decl @lune.@base.@TransUnit.TransUnit.checkAccesSym
func (self *TransUnit_TransUnit) checkAccesSym(_env *LnsEnv) {
    for _, _accessSym := range( self.tentativeSymbol.FP.Get_accessSymList(_env).Items ) {
        accessSym := _accessSym.(TransUnit_AccessSymPosDownCast).ToTransUnit_AccessSymPos()
        var symbolInfo *Ast_SymbolInfo
        symbolInfo = accessSym.FP.Get_symbol(_env)
        if Lns_op_not(symbolInfo.FP.Get_hasValueFlag(_env)){
            self.FP.AddErrMess(_env, accessSym.FP.Get_pos(_env), _env.GetVM().String_format("This can't access variable have no value -- %s", []LnsAny{symbolInfo.FP.Get_name(_env)}))
        }
    }
    self.tentativeSymbol.FP.ClearAccessSym(_env)
}
// 968: decl @lune.@base.@TransUnit.TransUnit.finishTentativeSymbol
func (self *TransUnit_TransUnit) finishTentativeSymbol(_env *LnsEnv, complete bool) {
    self.FP.checkAccesSym(_env)
    var tentativeSymbol *TransUnit_TentativeSymbol
    tentativeSymbol = self.tentativeSymbol
    self.tentativeSymbol = Lns_unwrap( tentativeSymbol.FP.Finish(_env, complete)).(*TransUnit_TentativeSymbol)
    {
        var errSymMap *LnsMap
        errSymMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _symbolInfo := range( tentativeSymbol.FP.Get_initSymSet(_env).Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if tentativeSymbol.FP.Get_scope(_env).FP.Get_parent(_env) == symbolInfo.FP.Get_scope(_env){
                if Lns_op_not(symbolInfo.FP.Get_hasValueFlag(_env)){
                    errSymMap.Set(symbolInfo.FP.Get_name(_env),symbolInfo)
                }
            }
        }
        {
            __forsortCollection0 := errSymMap
            __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
            __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
            for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
                symbolInfo := __forsortCollection0.Items[ ___forsortKey0 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.AddErrMess(_env, self.Parser.FP.GetLastPos(_env), _env.GetVM().String_format("There is the case no initialized value for '%s'", []LnsAny{symbolInfo.FP.Get_name(_env)}))
            }
        }
    }
    if tentativeSymbol.FP.Get_scope(_env).FP.Get_validCheckingUnaccess(_env){
        {
            __forsortCollection1 := tentativeSymbol.FP.Get_scope(_env).FP.Get_symbol2SymbolInfoMap(_env)
            __forsortSorted1 := __forsortCollection1.CreateKeyListStr()
            __forsortSorted1.Sort( _env, LnsItemKindStr, nil )
            for _, ___forsortKey1 := range( __forsortSorted1.Items ) {
                symbolInfo := __forsortCollection1.Items[ ___forsortKey1 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(symbolInfo.FP.Get_posForModToRef(_env))) ||
                        _env.SetStackVal( symbolInfo.FP.Get_posForModToRef(_env) != symbolInfo.FP.Get_posForLatestMod(_env)) ).(bool))) &&
                    _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Var) ||
                        _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Fun) ).(bool))) &&
                    _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, symbolInfo.FP.Get_accessMode(_env)))) &&
                    _env.SetStackVal( Lns_op_not(Lns_car(_env.GetVM().String_find(symbolInfo.FP.Get_name(_env),"^_", nil, nil)))) ).(bool)){
                    {
                        _pos := _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( symbolInfo.FP.Get_posForLatestMod(_env)) ||
                            _env.SetStackVal( symbolInfo.FP.Get_pos(_env)) )
                        if !Lns_IsNil( _pos ) {
                            pos := _pos.(Types_Position)
                            self.FP.addWarnMess(_env, pos, _env.GetVM().String_format("'%s' var isn't accessed", []LnsAny{symbolInfo.FP.Get_name(_env)}))
                        }
                    }
                }
            }
        }
    }
}
// 1006: decl @lune.@base.@TransUnit.TransUnit.mergeTentativeSymbol
func (self *TransUnit_TransUnit) mergeTentativeSymbol(_env *LnsEnv, scope *Ast_Scope) {
    self.FP.checkAccesSym(_env)
    self.tentativeSymbol.FP.NewSet(_env, scope)
}
// 1012: decl @lune.@base.@TransUnit.TransUnit.getCurrentClass
func (self *TransUnit_TransUnit) getCurrentClass(_env *LnsEnv) *Ast_TypeInfo {
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_headTypeInfo
    var scope *Ast_Scope
    scope = self.FP.Get_scope(_env)
    for {
        {
            __exp := scope.FP.Get_ownerTypeInfo(_env)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_TypeInfo)
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
                    _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__Module) ||
                    _env.SetStackVal( _exp.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ).(bool){
                    return _exp
                }
            }
        }
        scope = scope.FP.Get_parent(_env)
        if scope.FP.IsRoot(_env){ break }
    }
    return typeInfo
}
// 1027: decl @lune.@base.@TransUnit.TransUnit.pushExtModule
func (self *TransUnit_TransUnit) pushExtModule(_env *LnsEnv, externalFlag bool,name string,accessMode LnsInt,pos Types_Position,lazy bool,lang LnsInt,requirePath string) *TransUnitIF_NSInfo {
    var modName string
    modName = name
    {
        __ := self.FP.Get_scope(_env).FP.GetTypeInfoChild(_env, modName)
        if !Lns_IsNil( __ ) {
            self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("multiple define -- %s", []LnsAny{name}))
        }
    }
    var parentNsInfo *TransUnitIF_NSInfo
    parentNsInfo = self.FP.Get_curNsInfo(_env)
    var parentInfo *Ast_TypeInfo
    parentInfo = parentNsInfo.FP.Get_typeInfo(_env)
    var parentScope *Ast_Scope
    parentScope = self.FP.Get_scope(_env)
    var scope *Ast_Scope
    scope = self.FP.PushScope(_env, Ast_ScopeKind__Module, nil, nil)
    var typeInfo *Ast_TypeInfo
    typeInfo = self.ProcessInfo.FP.CreateExtModule(_env, scope, parentInfo, parentNsInfo.FP.Get_typeDataAccessor(_env), externalFlag, accessMode, name, lang, requirePath)
    parentScope.FP.AddExtModule(_env, self.ProcessInfo, name, pos, typeInfo, lazy, lang)
    if Lns_op_not(self.TypeId2ClassMap.Get(typeInfo.FP.Get_typeId(_env))){
        var namespace *Nodes_NamespaceInfo
        namespace = NewNodes_NamespaceInfo(_env, modName, self.FP.Get_scope(_env), typeInfo)
        self.TypeId2ClassMap.Set(typeInfo.FP.Get_typeId(_env),namespace)
    }
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.NewNSInfo(_env, typeInfo, pos)
    self.FP.Set_curNsInfo(_env, nsInfo)
    return nsInfo
}
// 1108: decl @lune.@base.@TransUnit.TransUnit.isTargetTokenPos
func (self *TransUnit_TransUnit) isTargetTokenPos(_env *LnsEnv, txt string,pos Types_Position) bool {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.analyzePos.LineNo == pos.LineNo) &&
        _env.SetStackVal( self.analyzePos.Column >= pos.Column) &&
        _env.SetStackVal( self.analyzePos.Column <= pos.Column + len(txt)) ).(bool)){
        return true
    }
    return false
}
// 1118: decl @lune.@base.@TransUnit.TransUnit.isTargetToken
func (self *TransUnit_TransUnit) isTargetToken(_env *LnsEnv, token *Types_Token) bool {
    return self.FP.isTargetTokenPos(_env, token.Txt, token.Pos)
}
// 1122: decl @lune.@base.@TransUnit.TransUnit.dumpSymbolType
func (self *TransUnit_TransUnit) dumpSymbolType(_env *LnsEnv, accessMode LnsInt,name string,typeInfo *Ast_TypeInfo) {
    var writer *Writer_JSON
    writer = NewWriter_JSON(_env, Lns_io_stdout)
    writer.FP.StartParent(_env, "lunescript", false)
    writer.FP.StartParent(_env, "inquire", false)
    writer.FP.Write(_env, "access", Ast_accessMode2txt(_env, accessMode))
    writer.FP.Write(_env, "name", name)
    writer.FP.Write(_env, "type", typeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil))
    writer.FP.Write(_env, "typeKind", Ast_TypeInfoKind_getTxt( typeInfo.FP.Get_kind(_env)))
    writer.FP.Write(_env, "static", _env.GetVM().String_format("%s", []LnsAny{typeInfo.FP.Get_staticFlag(_env)}))
    writer.FP.Write(_env, "display", typeInfo.FP.Get_display_stirng_with(_env, typeInfo.FP.Get_rawTxt(_env), nil))
    writer.FP.EndElement(_env)
    writer.FP.EndElement(_env)
    writer.FP.Fin(_env)
    _env.GetVM().OS_exit(0)
}
// 1138: decl @lune.@base.@TransUnit.TransUnit.errorShadowingOp
func (self *TransUnit_TransUnit) ErrorShadowingOp(_env *LnsEnv, pos Types_Position,symbolInfo LnsAny,errFlag bool) {
    if symbolInfo != nil{
        symbolInfo_232 := symbolInfo.(*Ast_SymbolInfo)
        var symPos LnsAny
        symPos = symbolInfo_232.FP.Get_pos(_env)
        if symPos != nil{
            symPos_235 := symPos.(Types_Position)
            var mess string
            mess = _env.GetVM().String_format("This symbol is shadowed from %d:%d -- %s", []LnsAny{pos.LineNo, pos.Column, symbolInfo_232.FP.Get_name(_env)})
            self.FP.addWarnErrMess(_env, symPos_235, errFlag, mess)
        }
        var mess string
        mess = _env.GetVM().String_format("shadowing symbol of %s -- %s", []LnsAny{_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( symPos) &&
            _env.SetStackVal( _env.GetVM().String_format("%s:%s", []LnsAny{_env.NilAccFin(_env.NilAccPush(symPos) && 
            _env.NilAccPush(_env.NilAccPop().(Types_Position).LineNo)), _env.NilAccFin(_env.NilAccPush(symPos) && 
            _env.NilAccPush(_env.NilAccPop().(Types_Position).Column))})) ||
            _env.SetStackVal( "external") ).(string), symbolInfo_232.FP.Get_name(_env)})
        self.FP.addWarnErrMess(_env, pos, errFlag, mess)
    }
}
// 1156: decl @lune.@base.@TransUnit.TransUnit.errorShadowing
func (self *TransUnit_TransUnit) errorShadowing(_env *LnsEnv, pos Types_Position,symbolInfo LnsAny) {
    self.FP.ErrorShadowingOp(_env, pos, symbolInfo, Lns_op_not(self.Ctrl_info.WarningShadowing))
}
// 1160: decl @lune.@base.@TransUnit.TransUnit.checkShadowing
func (self *TransUnit_TransUnit) checkShadowing(_env *LnsEnv, pos Types_Position,name string,scope *Ast_Scope) {
    if name == "_"{
        return 
    }
    var symbolInfo *Ast_SymbolInfo
    
    {
        _symbolInfo := self.FP.Get_scope(_env).FP.GetSymbolTypeInfo(_env, name, scope, self.ModuleScope, self.ScopeAccess)
        if _symbolInfo == nil{
            return 
        } else {
            symbolInfo = _symbolInfo.(*Ast_SymbolInfo)
        }
    }
    self.FP.errorShadowing(_env, pos, symbolInfo)
}
// 1174: decl @lune.@base.@TransUnit.TransUnit.addLocalVar
func (self *TransUnit_TransUnit) AddLocalVar(_env *LnsEnv, pos Types_Position,argFlag bool,canBeLeft bool,name string,typeInfo *Ast_TypeInfo,mutable LnsInt,allowShadow LnsAny) *Ast_SymbolInfo {
    if Lns_op_not(allowShadow){
        self.FP.checkShadowing(_env, pos, name, self.FP.Get_scope(_env))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.analyzeMode == TransUnit_AnalyzeMode__Inquire) &&
        _env.SetStackVal( self.FP.isTargetTokenPos(_env, name, pos)) ).(bool)){
        self.FP.dumpSymbolType(_env, Ast_AccessMode__Local, name, typeInfo)
    }
    return Lns_unwrap( Lns_car(self.FP.Get_scope(_env).FP.AddLocalVar(_env, self.ProcessInfo, argFlag, canBeLeft, name, pos, typeInfo, mutable))).(*Ast_SymbolInfo)
}
// 1197: decl @lune.@base.@TransUnit.TransUnit.getLineNo
func (self *TransUnit_TransUnit) getLineNo(_env *LnsEnv, token *Types_Token) Types_Position {
    var pos Types_Position
    pos = token.Pos
    {
        _work := self.FuncBlockInfoLinkNo
        if !Lns_IsNil( _work ) {
            work := _work.(Types_Position)
            pos = work
        }
    }
    if self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.Get_mode(_env) != Nodes_MacroMode__None{
        pos = Lns_unwrap( self.MacroCtrl.FP.Get_macroCallLineNo(_env)).(Types_Position)
    }
    return pos
}
// 1209: decl @lune.@base.@TransUnit.TransUnit.canBeAsyncParam
func (self *TransUnit_TransUnit) canBeAsyncParam(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if _switch0 := typeInfo.FP.Get_nilableTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.Get_genSrcTypeInfo(_env); _switch0 == self.builtinFunc.G__pipe_ || _switch0 == self.builtinFunc.G__lns_sync_flag_ {
        return true
    }
    return Ast_TypeInfo_canBeAsyncParam(_env, typeInfo)
}
// 1279: decl @lune.@base.@TransUnit.TransUnit.createModifier
func (self *TransUnit_TransUnit) createModifier(_env *LnsEnv, typeInfo *Ast_TypeInfo,mutMode LnsInt) *Ast_TypeInfo {
    if Lns_op_not(self.ValidMutControl){
        return typeInfo
    }
    return self.ProcessInfo.FP.CreateModifier(_env, typeInfo, mutMode)
}
// 1288: decl @lune.@base.@TransUnit.TransUnit.createExtType
func (self *TransUnit_TransUnit) createExtType(_env *LnsEnv, pos Types_Position,typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    switch _matchExp0 := self.ProcessInfo.FP.CreateLuaval(_env, typeInfo, true).(type) {
    case *Ast_LuavalResult__OK:
    work := _matchExp0.Val1
        return work
    case *Ast_LuavalResult__Err:
    err := _matchExp0.Val1
        self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("not support -- %s", []LnsAny{err}))
        return typeInfo
    }
// insert a dummy
    return nil
}
// 1302: decl @lune.@base.@TransUnit.TransUnit.errorAt
func (self *TransUnit_TransUnit) ErrorAt(_env *LnsEnv, pos Types_Position,mess string) {
    self.FP.AddErrMess(_env, pos, mess)
    for _, _errmess := range( TransUnitIF_sortMess(_env, self.ErrMessList).Items ) {
        errmess := _errmess.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
        Util_errorLog(_env, errmess.FP.Get_mess(_env))
    }
    for _, _warnmess := range( TransUnitIF_sortMess(_env, self.WarnMessList).Items ) {
        warnmess := _warnmess.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
        Util_errorLog(_env, warnmess.FP.Get_mess(_env))
    }
    if self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.Get_mode(_env) != Nodes_MacroMode__None{
        Lns_print([]LnsAny{"------ near code -----", Nodes_MacroMode_getTxt( self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.Get_mode(_env))})
        Lns_print([]LnsAny{self.Parser.FP.GetNearCode(_env)})
        Lns_print([]LnsAny{"------"})
    }
    Util_err(_env, "has error")
}
// 1320: decl @lune.@base.@TransUnit.TransUnit.createNoneNode
func (self *TransUnit_TransUnit) createNoneNode(_env *LnsEnv, pos Types_Position) *Nodes_Node {
    return &Nodes_NoneNode_create(_env, self.NodeManager, pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})).Nodes_Node
}
// 1326: decl @lune.@base.@TransUnit.TransUnit.pushbackToken
func (self *TransUnit_TransUnit) PushbackToken(_env *LnsEnv, token *Types_Token) {
    self.Parser.FP.PushbackToken(_env, token)
}
// 1331: decl @lune.@base.@TransUnit.TransUnit.newPushback
func (self *TransUnit_TransUnit) NewPushback(_env *LnsEnv, tokenList *LnsList) {
    self.Parser.FP.NewPushback(_env, tokenList)
}
// 1335: decl @lune.@base.@TransUnit.TransUnit.getStreamName
func (self *TransUnit_TransUnit) GetStreamName(_env *LnsEnv) string {
    return self.Parser.FP.GetStreamName(_env)
}
// 1339: decl @lune.@base.@TransUnit.TransUnit.createPosition
func (self *TransUnit_TransUnit) CreatePosition(_env *LnsEnv, lineNo LnsInt,column LnsInt) Types_Position {
    return self.Parser.FP.CreatePosition(_env, lineNo, column)
}
// 1343: decl @lune.@base.@TransUnit.TransUnit.getTokenNoErr
func (self *TransUnit_TransUnit) GetTokenNoErr(_env *LnsEnv, skipFlag LnsAny) *Types_Token {
    var token *Types_Token
    var commentList *LnsList
    commentList = NewLnsList([]LnsAny{})
    var workToken *Types_Token
    workToken = self.Parser.FP.GetTokenNoErr(_env, nil)
    for workToken.Kind == Types_TokenKind__Cmnt {
        commentList.Insert(Types_Token2Stem(workToken))
        workToken = self.Parser.FP.GetTokenNoErr(_env, nil)
    }
    if workToken.Kind != Types_TokenKind__Eof{
        token = workToken
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.Get_mode(_env) != Nodes_MacroMode__None) &&
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(skipFlag)) ||
                _env.SetStackVal( Lns_op_not(self.inTestBlock)) ||
                _env.SetStackVal( self.Ctrl_info.Testing) ).(bool))) ).(bool)){
            token = self.MacroCtrl.FP.ExpandMacroVal(_env, self.TypeNameCtrl, self.FP.Get_scope(_env), self.FP, token)
        }
        if Lns_op_not(self.Ctrl_info.CompatComment){
            self.commentCtrl.FP.AddDirect(_env, commentList)
        }
    } else { 
        token = Parser_getEofToken(_env)
        self.commentCtrl.FP.AddDirect(_env, commentList)
    }
    if token.FP.Get_commentList(_env).Len() > 0{
        self.commentCtrl.FP.Add(_env, token)
    }
    return token
}
// 1380: decl @lune.@base.@TransUnit.TransUnit.getToken
func (self *TransUnit_TransUnit) getToken(_env *LnsEnv, allowEof LnsAny) *Types_Token {
    var token *Types_Token
    token = self.FP.GetTokenNoErr(_env, nil)
    if token == Parser_getEofToken(_env){
        if Lns_isCondTrue( allowEof){
            return token
        }
        self.FP.Error(_env, "EOF")
    }
    return token
}
// 1392: decl @lune.@base.@TransUnit.TransUnit.pushback
func (self *TransUnit_TransUnit) Pushback(_env *LnsEnv) {
    self.Parser.FP.Pushback(_env)
}
// 1396: decl @lune.@base.@TransUnit.TransUnit.pushbackStr
func (self *TransUnit_TransUnit) PushbackStr(_env *LnsEnv, asyncParse LnsAny,name string,statement string,pos Types_Position) {
    var async bool
    
    {
        _async := asyncParse
        if _async == nil{
            if self.Ctrl_info.MacroAsyncParseStmtLen == 0{
                async = false
            } else { 
                async = len(statement) >= self.Ctrl_info.MacroAsyncParseStmtLen
            }
        } else {
            async = _async.(bool)
        }
    }
    self.Parser.FP.PushbackStr(_env, async, name, statement, pos)
}
// 1432: decl @lune.@base.@TransUnit.TransUnit.checkSymbol
func (self *TransUnit_TransUnit) checkSymbol(_env *LnsEnv, token *Types_Token,mode LnsInt) *Types_Token {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Kind != Types_TokenKind__Symb) &&
        _env.SetStackVal( token.Kind != Types_TokenKind__Kywd) &&
        _env.SetStackVal( token.Kind != Types_TokenKind__Type) ).(bool)){
        self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("illegal symbol -- '%s'", []LnsAny{token.Txt}))
    }
    var frontChar LnsInt
    frontChar = LnsInt(token.Txt[1-1])
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode == TransUnit_SymbolMode__Must_) &&
        _env.SetStackVal( frontChar != 95) ).(bool)){
        self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("macro name must begin with '_' -- '%s'", []LnsAny{token.Txt}))
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode != TransUnit_SymbolMode__Must_) &&
        _env.SetStackVal( frontChar == 95) ).(bool)){
        if Lns_op_not(self.IgnoreToCheckSymbol_){
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( mode == TransUnit_SymbolMode__MustNot_Or_) &&
                _env.SetStackVal( token.Txt == "_") ).(bool)){
            } else if Lns_op_not(TransUnit_specialSymbolSet.Has(token.Txt)){
                self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("symbol must not begin with '_' -- '%s'", []LnsAny{token.Txt}))
            }
        }
    } else if TransUnit_builtinKeywordSet.Has(token.Txt){
        self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("this symbol is special keyword -- %s", []LnsAny{token.Txt}))
    } else if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Parser_isLuaKeyword(_env, token.Txt)) ||
        _env.SetStackVal( Parser_isOp2(_env, token.Txt)) ||
        _env.SetStackVal( Parser_isOp1(_env, token.Txt)) ).(bool){
        self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("this symbol is lua keyword -- %s", []LnsAny{token.Txt}))
    }
    return token
}
// 1468: decl @lune.@base.@TransUnit.TransUnit.getSymbolToken
func (self *TransUnit_TransUnit) getSymbolToken(_env *LnsEnv, mode LnsInt) *Types_Token {
    return self.FP.checkSymbol(_env, self.FP.getToken(_env, nil), mode)
}
// 1473: decl @lune.@base.@TransUnit.TransUnit.checkToken
func (self *TransUnit_TransUnit) checkToken(_env *LnsEnv, token *Types_Token,txt string) *Types_Token {
    if token.Txt != txt{
        self.FP.Error(_env, _env.GetVM().String_format("not found -- %s. expects %s", []LnsAny{txt, token.Txt}))
    }
    return token
}
// 1480: decl @lune.@base.@TransUnit.TransUnit.checkNextToken
func (self *TransUnit_TransUnit) checkNextToken(_env *LnsEnv, txt string) *Types_Token {
    return self.FP.checkToken(_env, self.FP.getToken(_env, nil), txt)
}
// 1490: decl @lune.@base.@TransUnit.TransUnit.getContinueToken
func (self *TransUnit_TransUnit) getContinueToken(_env *LnsEnv)(*Types_Token, bool) {
    var token *Types_Token
    token = self.FP.getToken(_env, nil)
    return token, token.Consecutive
}
// 1496: decl @lune.@base.@TransUnit.TransUnit.getDefaultAsync
func (self *TransUnit_TransUnit) getDefaultAsync(_env *LnsEnv, kind LnsInt,classTypeInfo LnsAny,asyncMode LnsAny) LnsInt {
    if asyncMode != nil{
        asyncMode_327 := asyncMode.(LnsInt)
        return asyncMode_327
    }
    if classTypeInfo != nil{
        classTypeInfo_335 := classTypeInfo.(*Ast_TypeInfo)
        {
            __exp := self.class2defaultAsyncMode.Get(classTypeInfo_335)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(LnsInt)
                return TransUnit_getDefaultAsync__process_0_(_env, _exp)
            }
        }
    }
    if _switch0 := kind; _switch0 == Ast_TypeInfoKind__Method {
        if self.defaultAsyncMode == TransUnit_DefaultAsyncMode__AsyncAll{
            return Ast_Async__Async
        }
    } else if _switch0 == Ast_TypeInfoKind__Func || _switch0 == Ast_TypeInfoKind__FormFunc {
        return TransUnit_getDefaultAsync__process_0_(_env, self.defaultAsyncMode)
    }
    return Ast_Async__Noasync
}
// 1543: decl @lune.@base.@TransUnit.TransUnit.analyzeStatementList
func (self *TransUnit_TransUnit) analyzeStatementList(_env *LnsEnv, stmtList *LnsList,firstSwitchingParser bool,termTxt LnsAny)(LnsAny, LnsInt) {
    var breakKind LnsInt
    breakKind = Nodes_BreakKind__None
    if stmtList.Len() > 0{
        breakKind = stmtList.GetAt(stmtList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.GetBreakKind(_env, Nodes_CheckBreakMode__Normal)
    }
    var parser2lastLineMap *LnsMap
    parser2lastLineMap = NewLnsMap( map[LnsAny]LnsAny{})
    var getLastLineNo func(_env *LnsEnv) LnsInt
    getLastLineNo = func(_env *LnsEnv) LnsInt {
        {
            _lastLineNo := parser2lastLineMap.Get(self.Parser.FP)
            if !Lns_IsNil( _lastLineNo ) {
                lastLineNo := _lastLineNo.(LnsInt)
                return lastLineNo
            }
        }
        return self.Parser.FP.GetLastPos(_env).LineNo
    }
    var setLastLineNo func(_env *LnsEnv, lineNo LnsInt)
    setLastLineNo = func(_env *LnsEnv, lineNo LnsInt) {
        parser2lastLineMap.Set(self.Parser.FP,lineNo)
    }
    var lastStatement LnsAny
    lastStatement = nil
    var lastLineNo LnsInt
    lastLineNo = getLastLineNo(_env)
    var setTailComment func(_env *LnsEnv, statement LnsAny) LnsInt
    setTailComment = func(_env *LnsEnv, statement LnsAny) LnsInt {
        var blank LnsInt
        var commentList *LnsList
        commentList = self.commentCtrl.FP.Get_commentList(_env)
        if commentList.Len() > 0{
            if lastStatement != nil{
                lastStatement_364 := lastStatement.(*Nodes_Node)
                var tailComment LnsAny
                tailComment = nil
                for _, _comment := range( commentList.Items ) {
                    comment := _comment.(Types_TokenDownCast).ToTypes_Token()
                    if comment.Pos.LineNo == lastStatement_364.FP.Get_pos(_env).LineNo{
                        if Lns_op_not(tailComment){
                            lastStatement_364.FP.Set_tailComment(_env, comment)
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
            blank = commentList.GetAt(1).(Types_TokenDownCast).ToTypes_Token().Pos.LineNo - commentList.GetAt(1).(Types_TokenDownCast).ToTypes_Token().FP.GetLineCount(_env) - lastLineNo
        } else { 
            if statement != nil{
                statement_375 := statement.(*Nodes_Node)
                blank = statement_375.FP.Get_pos(_env).LineNo - lastLineNo
            } else {
                blank = self.Parser.FP.GetLastPos(_env).LineNo - lastLineNo
            }
        }
        return blank
    }
    for  {
        lastLineNo = getLastLineNo(_env)
        var statement LnsAny
        statement = self.FP.analyzeStatement(_env, termTxt)
        if statement != nil{
            statement_380 := statement.(*Nodes_Node)
            if breakKind != Nodes_BreakKind__None{
                if statement_380.FP.Get_kind(_env) != Nodes_NodeKind_get_BlankLine(_env){
                    self.FP.AddErrMess(_env, statement_380.FP.Get_pos(_env), _env.GetVM().String_format("This statement is not reached -- %s", []LnsAny{Nodes_BreakKind_getTxt( breakKind)}))
                }
            }
            var blank LnsInt
            blank = setTailComment(_env, statement_380)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( blank > 1) &&
                _env.SetStackVal( Lns_op_not(firstSwitchingParser)) ).(bool)){
                stmtList.Insert(Nodes_BlankLineNode2Stem(Nodes_BlankLineNode_create(_env, self.NodeManager, self.FP.CreatePosition(_env, lastLineNo + 1, 0), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), blank - 1)))
            }
            setLastLineNo(_env, self.Parser.FP.GetLastPos(_env).LineNo)
            if firstSwitchingParser{
                firstSwitchingParser = false
            }
            stmtList.Insert(Nodes_Node2Stem(statement_380))
            lastStatement = statement_380
            if statement_380.FP.Get_kind(_env) != Nodes_NodeKind_get_BlankLine(_env){
                breakKind = statement_380.FP.GetBreakKind(_env, Nodes_CheckBreakMode__Normal)
            }
            statement_380.FP.AddComment(_env, self.commentCtrl.FP.Get_commentList(_env))
            self.commentCtrl.FP.Clear(_env)
        } else {
            setTailComment(_env, nil)
            break
        }
    }
    return lastStatement, lastLineNo
}
// 1657: decl @lune.@base.@TransUnit.TransUnit.analyzeStatementListSubfile
func (self *TransUnit_TransUnit) analyzeStatementListSubfile(_env *LnsEnv, stmtList *LnsList) LnsAny {
    var statement LnsAny
    statement = self.FP.analyzeStatement(_env, nil)
    {
        __exp := statement
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_Node)
            if _exp.FP.Get_kind(_env) != Nodes_NodeKind_get_Subfile(_env){
                self.FP.Error(_env, "subfile must have 'subfile' declaration at top.")
            }
        } else {
            self.FP.Error(_env, "subfile must have 'subfile' declaration at top.")
        }
    }
    return Lns_car(self.FP.analyzeStatementList(_env, stmtList, true, nil))
}
// 1674: decl @lune.@base.@TransUnit.TransUnit.supportLang
func (self *TransUnit_TransUnit) supportLang(_env *LnsEnv, lang string) bool {
    for _pragma := range( self.HelperInfo.PragmaSet.Items ) {
        pragma := _pragma
        switch _matchExp0 := pragma.(type) {
        case *LuneControl_Pragma__limit_conv_code:
        codeSet := _matchExp0.Val1
            return codeSet.Has(lang)
        }
    }
    return true
}
// 1685: decl @lune.@base.@TransUnit.TransUnit.analyzeLuneControlToken
func (self *TransUnit_TransUnit) AnalyzeLuneControlToken(_env *LnsEnv, firstToken *Types_Token,controlToken *Types_Token) LnsAny {
    return nil
}
// 1691: decl @lune.@base.@TransUnit.TransUnit.analyzeLuneControl
func (self *TransUnit_TransUnit) analyzeLuneControl(_env *LnsEnv, firstToken *Types_Token) LnsAny {
    var node LnsAny
    node = nil
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var pragma LnsAny
    {
        _work := self.FP.AnalyzeLuneControlToken(_env, firstToken, nextToken)
        if !Lns_IsNil( _work ) {
            work := _work
            pragma = work
        } else {
            if _switch0 := (nextToken.Txt); _switch0 == "load__lune_module" {
                pragma = LuneControl_Pragma__load__lune_module_Obj
            } else if _switch0 == "run_async_pipe" {
                if Lns_op_not(self.HelperInfo.PragmaSet.Has(LuneControl_Pragma__use_async_Obj)){
                    self.FP.AddErrMess(_env, nextToken.Pos, "must set '_lune_control use_async'")
                }
                var nowMethod *Ast_TypeInfo
                nowMethod = self.FP.GetCurrentNamespaceTypeInfo(_env)
                var nowClass *Ast_TypeInfo
                nowClass = nowMethod.FP.Get_parentInfo(_env)
                var valid bool
                valid = false
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( nowMethod.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) &&
                    _env.SetStackVal( Ast_isClass(_env, nowClass)) ).(bool)){
                    {
                        _loopMethod := _env.NilAccFin(_env.NilAccPush(nowClass.FP.Get_scope(_env)) && 
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetTypeInfoChild(_env, "loop")})/* 1744:37 */)
                        if !Lns_IsNil( _loopMethod ) {
                            loopMethod := _loopMethod.(*Ast_TypeInfo)
                            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( loopMethod.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) &&
                                _env.SetStackVal( loopMethod.FP.Get_argTypeInfoList(_env).Len() == 0) ).(bool)){
                                valid = true
                            }
                        }
                    }
                }
                if valid{
                    pragma = LuneControl_Pragma__run_async_pipe_Obj
                } else { 
                    self.FP.AddErrMess(_env, nextToken.Pos, "this option only use in method of the class have loop method.")
                    return nil
                }
            } else {
                self.FP.AddErrMess(_env, nextToken.Pos, _env.GetVM().String_format("unknown option -- %s", []LnsAny{nextToken.Txt}))
                self.FP.checkNextToken(_env, ";")
                return nil
            }
        }
    }
    node = Nodes_LuneControlNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), pragma)
    self.HelperInfo.PragmaSet.Add(pragma)
    self.FP.checkNextToken(_env, ";")
    return node
}
// 1807: decl @lune.@base.@TransUnit.TransUnit.analyzeBlock
func (self *TransUnit_TransUnit) analyzeBlock(_env *LnsEnv, blockKind LnsInt,tentativeMode LnsInt,scope LnsAny,refAccessSymPosList LnsAny) *Nodes_BlockNode {
    var backScope *Ast_Scope
    backScope = self.FP.Get_scope(_env)
    if scope != nil{
        scope_423 := scope.(*Ast_Scope)
        self.FP.SetScope(_env, scope_423, TransUnitIF_SetNSInfo__FromScope)
    } else {
        self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    }
    var blockScope *Ast_Scope
    blockScope = self.FP.Get_scope(_env)
    blockScope.FP.AddIgnoredVar(_env, self.ProcessInfo)
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.Get_curNsInfo(_env)
    if _switch0 := tentativeMode; _switch0 == TransUnit_TentativeMode__Loop {
        self.FP.prepareTentativeSymbol(_env, self.FP.Get_scope(_env), true, refAccessSymPosList)
    } else if _switch0 == TransUnit_TentativeMode__Simple || _switch0 == TransUnit_TentativeMode__Start || _switch0 == TransUnit_TentativeMode__Ignore {
        self.FP.prepareTentativeSymbol(_env, self.FP.Get_scope(_env), false, nil)
    } else if _switch0 == TransUnit_TentativeMode__Merge || _switch0 == TransUnit_TentativeMode__Finish {
        self.FP.mergeTentativeSymbol(_env, self.FP.Get_scope(_env))
    }
    var loopFlag bool
    loopFlag = false
    if _switch1 := blockKind; _switch1 == Nodes_BlockKind__For || _switch1 == Nodes_BlockKind__Apply || _switch1 == Nodes_BlockKind__While || _switch1 == Nodes_BlockKind__Repeat || _switch1 == Nodes_BlockKind__Foreach {
        loopFlag = true
        nsInfo.FP.Get_loopScopeQueue(_env).Insert(Ast_Scope2Stem(self.FP.Get_scope(_env)))
    }
    var stmtList *LnsList
    stmtList = NewLnsList([]LnsAny{})
    var token *Types_Token
    token = self.FP.checkNextToken(_env, "{")
    self.FP.analyzeStatementList(_env, stmtList, false, "}")
    nsInfo.FP.AddStmtNum(_env, stmtList.Len())
    self.FP.checkNextToken(_env, "}")
    if loopFlag{
        nsInfo.FP.Get_loopScopeQueue(_env).Remove(nil)
    }
    if scope != nil{
        self.FP.SetScope(_env, backScope, TransUnitIF_SetNSInfo__FromScope)
    } else {
        self.FP.PopScope(_env)
    }
    var node *Nodes_BlockNode
    node = Nodes_BlockNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), blockKind, blockScope, stmtList)
    if node.FP.GetBreakKind(_env, Nodes_CheckBreakMode__Normal) != Nodes_BreakKind__None{
        self.tentativeSymbol.FP.Skip(_env)
    }
    if blockKind != Nodes_BlockKind__Repeat{
        if _switch2 := tentativeMode; _switch2 == TransUnit_TentativeMode__Simple || _switch2 == TransUnit_TentativeMode__Finish {
            self.FP.finishTentativeSymbol(_env, true)
        } else if _switch2 == TransUnit_TentativeMode__Ignore || _switch2 == TransUnit_TentativeMode__Loop {
            self.FP.finishTentativeSymbol(_env, false)
        }
    }
    return node
}
// 1892: decl @lune.@base.@TransUnit.TransUnit.skipBlock
func (self *TransUnit_TransUnit) skipBlock(_env *LnsEnv, recordToken bool) *LnsList {
    var blockDepth LnsInt
    blockDepth = 0
    var tokenList *LnsList
    tokenList = NewLnsList([]LnsAny{})
    for  {
        var token *Types_Token
        token = self.FP.GetTokenNoErr(_env, Lns_op_not(recordToken))
        if recordToken{
            tokenList.Insert(Types_Token2Stem(token))
        }
        if token.Kind == Types_TokenKind__Eof{
            self.FP.Error(_env, "EOF")
        }
        if _switch0 := token.Txt; _switch0 == "{" {
            blockDepth = blockDepth + 1
        } else if _switch0 == "}" {
            blockDepth = blockDepth - 1
            if blockDepth <= 0{
                break
            }
        } else {
        }
    }
    if blockDepth < 0{
        self.FP.Error(_env, "mismatch '}'")
    }
    return tokenList
}
// 1925: decl @lune.@base.@TransUnit.TransUnit.skipAndCreateDummyBlock
func (self *TransUnit_TransUnit) skipAndCreateDummyBlock(_env *LnsEnv) *Nodes_BlockNode {
    var blockToken *Types_Token
    blockToken = self.FP.checkNextToken(_env, "{")
    self.FP.Pushback(_env)
    var blockScope *Ast_Scope
    blockScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var stmtList *LnsList
    stmtList = NewLnsList([]LnsAny{})
    self.FP.PopScope(_env)
    self.FP.skipBlock(_env, false)
    return Nodes_BlockNode_create(_env, self.NodeManager, blockToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), Nodes_BlockKind__Func, blockScope, stmtList)
}
// 1941: decl @lune.@base.@TransUnit.TransUnit.analyzeRequest
func (self *TransUnit_TransUnit) analyzeRequest(_env *LnsEnv, reqToken *Types_Token) *Nodes_RequestNode {
    var processor *Nodes_Node
    processor = self.FP.analyzeExp(_env, false, true, false, nil, Ast_builtinTypeProcessor)
    if processor.FP.Get_expType(_env) != Ast_builtinTypeProcessor{
        self.FP.AddErrMess(_env, processor.FP.Get_pos(_env), _env.GetVM().String_format("It must be 'processor'. -- %s", []LnsAny{processor.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
    }
    var exp *Nodes_Node
    exp = self.FP.analyzeExp(_env, false, true, false, nil, nil)
    return Nodes_RequestNode_create(_env, self.NodeManager, reqToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), exp.FP.Get_expTypeList(_env), processor, exp)
}
// 1957: decl @lune.@base.@TransUnit.TransUnit.analyzeAsyncLock
func (self *TransUnit_TransUnit) analyzeAsyncLock(_env *LnsEnv, asyncToken *Types_Token,lockKind LnsInt) *Nodes_Node {
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.GetNSInfo(_env, self.FP.GetCurrentNamespaceTypeInfo(_env))
    if _switch0 := lockKind; _switch0 == Nodes_LockKind__AsyncLock {
        if nsInfo.FP.IsNoasync(_env){
            self.FP.AddErrMess(_env, asyncToken.Pos, _env.GetVM().String_format("can't use __asyncLock on __noasync. -- %s", []LnsAny{nsInfo.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
    } else if _switch0 == Nodes_LockKind__LuaLock || _switch0 == Nodes_LockKind__LuaDepend {
        if nsInfo.FP.IsNoasync(_env){
            self.FP.AddErrMess(_env, asyncToken.Pos, _env.GetVM().String_format("can't use __luaDepend or __luaLock on __noasync. -- %s", []LnsAny{nsInfo.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
    } else if _switch0 == Nodes_LockKind__LuaGo {
        if Lns_op_not(nsInfo.FP.IsNoasync(_env)){
            self.FP.AddErrMess(_env, asyncToken.Pos, _env.GetVM().String_format("can't use __luago on __async. -- %s", []LnsAny{nsInfo.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    nsInfo.FP.IncLock(_env, lockKind)
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(_env, Nodes_BlockKind__AsyncLock, TransUnit_TentativeMode__Simple, nil, nil)
    nsInfo.FP.DecLock(_env)
    return &Nodes_AsyncLockNode_create(_env, self.NodeManager, asyncToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), lockKind, block).Nodes_Node
}
// 2001: decl @lune.@base.@TransUnit.TransUnit.analyzeIf
func (self *TransUnit_TransUnit) analyzeIf(_env *LnsEnv, token *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    var continueFlag bool
    nextToken,continueFlag = self.FP.getContinueToken(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( continueFlag) &&
        _env.SetStackVal( nextToken.Txt == "!") ).(bool)){
        return &self.FP.analyzeIfUnwrap(_env, token).Nodes_Node
    }
    self.FP.Pushback(_env)
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    var ifExp *Nodes_Node
    ifExp = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
    list.Insert(Nodes_IfStmtInfo2Stem(NewNodes_IfStmtInfo(_env, Nodes_IfKind__If, ifExp, self.FP.analyzeBlock(_env, Nodes_BlockKind__If, TransUnit_TentativeMode__Start, nil, nil))))
    var checkCond func(_env *LnsEnv, condExp *Nodes_Node)
    checkCond = func(_env *LnsEnv, condExp *Nodes_Node) {
        if _switch0 := condExp.FP.Get_expType(_env).FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Nilable || _switch0 == Ast_TypeInfoKind__Stem {
        } else if _switch0 == Ast_TypeInfoKind__Prim {
            if Lns_op_not(condExp.FP.Get_expType(_env).FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeBool, nil, nil)){
                self.FP.AddErrMess(_env, condExp.FP.Get_pos(_env), _env.GetVM().String_format("This exp never be false -- %s", []LnsAny{condExp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
            }
        } else {
            self.FP.AddErrMess(_env, condExp.FP.Get_pos(_env), _env.GetVM().String_format("This exp never be false -- %s", []LnsAny{condExp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    checkCond(_env, ifExp)
    nextToken = self.FP.getToken(_env, true)
    if nextToken.Txt == "elseif"{
        for nextToken.Txt == "elseif" {
            var condExp *Nodes_Node
            condExp = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
            list.Insert(Nodes_IfStmtInfo2Stem(NewNodes_IfStmtInfo(_env, Nodes_IfKind__ElseIf, condExp, self.FP.analyzeBlock(_env, Nodes_BlockKind__Elseif, TransUnit_TentativeMode__Merge, nil, nil))))
            checkCond(_env, condExp)
            nextToken = self.FP.getToken(_env, true)
        }
    }
    if nextToken.Txt == "else"{
        list.Insert(Nodes_IfStmtInfo2Stem(NewNodes_IfStmtInfo(_env, Nodes_IfKind__Else, self.FP.createNoneNode(_env, nextToken.Pos), self.FP.analyzeBlock(_env, Nodes_BlockKind__Else, TransUnit_TentativeMode__Finish, nil, nil))))
    } else { 
        self.FP.finishTentativeSymbol(_env, false)
        self.FP.Pushback(_env)
    }
    return &Nodes_IfNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), list).Nodes_Node
}
// 2066: decl @lune.@base.@TransUnit.TransUnit.processCaseDefault
func (self *TransUnit_TransUnit) processCaseDefault(_env *LnsEnv, firstToken *Types_Token,caseKind LnsInt,nextToken *Types_Token,hasCase bool)(LnsAny, bool) {
    var keyword string
    keyword = TransUnit_convExp1_6184(Lns_2DDD(_env.GetVM().String_gsub(firstToken.Txt,"_", "")))
    var fullKeyword string
    fullKeyword = _env.GetVM().String_format("_%s", []LnsAny{keyword})
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( firstToken.Txt == fullKeyword) &&
        _env.SetStackVal( caseKind != Nodes_CaseKind__MustFull) ).(bool)){
        self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("This '%s' hasn't enough 'case' condition.", []LnsAny{keyword}))
    }
    var defaultBlock LnsAny
    defaultBlock = nil
    var failSafeDefault bool
    failSafeDefault = false
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( nextToken.Txt == "default") ||
        _env.SetStackVal( nextToken.Txt == "_default") ).(bool){
        if firstToken.Txt == fullKeyword{
            self.FP.AddErrMess(_env, nextToken.Pos, _env.GetVM().String_format("'_%s' can't have default.", []LnsAny{keyword}))
        }
        if nextToken.Txt == "_default"{
            failSafeDefault = true
        } else if caseKind == Nodes_CaseKind__Full{
            self.FP.addWarnMess(_env, nextToken.Pos, _env.GetVM().String_format("This '%s' has full case. This 'default' is no reach.", []LnsAny{keyword}))
        }
        defaultBlock = self.FP.analyzeBlock(_env, Nodes_BlockKind__Default, _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(hasCase)) &&
            _env.SetStackVal( TransUnit_TentativeMode__Simple) ||
            _env.SetStackVal( TransUnit_TentativeMode__Finish) ).(LnsInt), nil, nil)
    } else { 
        if hasCase{
            self.FP.finishTentativeSymbol(_env, caseKind != Nodes_CaseKind__Lack)
        }
        self.FP.Pushback(_env)
    }
    self.FP.checkNextToken(_env, "}")
    if Lns_op_not(hasCase){
        self.FP.addWarnMess(_env, firstToken.Pos, _env.GetVM().String_format("'%s' should have 'case' blocks.", []LnsAny{keyword}))
        if Lns_isCondTrue( defaultBlock){
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("'%s' must have 'case' blocks when have 'default' block.", []LnsAny{keyword}))
        }
    }
    return defaultBlock, failSafeDefault
}
// 2118: decl @lune.@base.@TransUnit.TransUnit.analyzeSwitch
func (self *TransUnit_TransUnit) analyzeSwitch(_env *LnsEnv, firstToken *Types_Token) *Nodes_SwitchNode {
    var exp *Nodes_Node
    exp = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
    self.FP.checkNextToken(_env, "{")
    var caseList *LnsList
    caseList = NewLnsList([]LnsAny{})
    var condObjSet *LnsSet
    condObjSet = NewLnsSet([]LnsAny{})
    var condSymIdSet *LnsSet
    condSymIdSet = NewLnsSet([]LnsAny{})
    var hasNilCond bool
    hasNilCond = false
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var firstFlag bool
    firstFlag = true
    for (nextToken.Txt == "case") {
        self.FP.checkToken(_env, nextToken, "case")
        var condexpList *Nodes_ExpListNode
        condexpList = self.FP.analyzeExpList(_env, false, false, false, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp.FP.Get_expType(_env))}), true)
        var condBock *Nodes_BlockNode
        condBock = self.FP.analyzeBlock(_env, Nodes_BlockKind__Switch, _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( firstFlag) &&
            _env.SetStackVal( TransUnit_TentativeMode__Start) ||
            _env.SetStackVal( TransUnit_TentativeMode__Merge) ).(LnsInt), nil, nil)
        if firstFlag{
            firstFlag = false
        }
        for _, _condExp := range( condexpList.FP.Get_expList(_env).Items ) {
            condExp := _condExp.(Nodes_NodeDownCast).ToNodes_Node()
            if condExp.FP.Get_expType(_env).FP.Get_nilable(_env){
                if hasNilCond{
                    self.FP.addWarnMess(_env, condExp.FP.Get_pos(_env), "multiple case with nil or nilable")
                } else { 
                    hasNilCond = true
                }
            }
            {
                _condLiteral := TransUnit_convExp1_6721(Lns_2DDD(condExp.FP.GetLiteral(_env)))
                if !Lns_IsNil( _condLiteral ) {
                    condLiteral := _condLiteral
                    var literalObj LnsAny
                    literalObj = Nodes_getLiteralObj(_env, condLiteral)
                    if literalObj != nil{
                        literalObj_526 := literalObj
                        if condObjSet.Has(literalObj_526){
                            self.FP.AddErrMess(_env, condExp.FP.Get_pos(_env), _env.GetVM().String_format("multiple case exp -- %s", []LnsAny{literalObj_526}))
                        } else { 
                            condObjSet.Add(literalObj_526)
                        }
                    }
                } else {
                    {
                        _expRef := Nodes_ExpRefNodeDownCastF(condExp.FP)
                        if !Lns_IsNil( _expRef ) {
                            expRef := _expRef.(*Nodes_ExpRefNode)
                            var symInfo *Ast_SymbolInfo
                            symInfo = expRef.FP.Get_symbolInfo(_env)
                            if condSymIdSet.Has(symInfo.FP.Get_symbolId(_env)){
                                self.FP.AddErrMess(_env, condExp.FP.Get_pos(_env), _env.GetVM().String_format("multiple case exp -- %s", []LnsAny{symInfo.FP.Get_name(_env)}))
                            } else { 
                                condSymIdSet.Add(symInfo.FP.Get_symbolId(_env))
                            }
                        }
                    }
                }
            }
            if Lns_op_not(Lns_car(exp.FP.Get_expType(_env).FP.CanEvalWith(_env, self.ProcessInfo, condExp.FP.Get_expType(_env), Ast_CanEvalType__Equal, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)){
                self.FP.AddErrMess(_env, condExp.FP.Get_pos(_env), _env.GetVM().String_format("case exp unmatch type -- %s <- %s", []LnsAny{exp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil), condExp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
            }
        }
        caseList.Insert(Nodes_CaseInfo2Stem(NewNodes_CaseInfo(_env, condexpList, condBock)))
        nextToken = self.FP.getToken(_env, nil)
    }
    var caseKind LnsInt
    {
        _enumType := Ast_EnumTypeInfoDownCastF(exp.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env).FP.Get_aliasSrc(_env).FP)
        if !Lns_IsNil( _enumType ) {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            var miss bool
            miss = false
            for _, _enumVal := range( enumType.FP.Get_name2EnumValInfo(_env).Items ) {
                enumVal := _enumVal.(Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
                if Lns_op_not(condObjSet.Has(Ast_getEnumLiteralVal(_env, enumVal.FP.Get_val(_env)))){
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
                self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), "The condition of '_switch' must be enum.")
            }
        }
    }
    var defaultBlock LnsAny
    var failSafeDefault bool
    defaultBlock,failSafeDefault = self.FP.processCaseDefault(_env, firstToken, caseKind, nextToken, caseList.Len() != 0)
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.Get_curNsInfo(_env)
    return Nodes_SwitchNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), nsInfo.FP.GetNextStmtId(_env, TransUnitIF_StmtKind__Switch), exp, caseList, defaultBlock, caseKind, failSafeDefault)
}
// 2230: decl @lune.@base.@TransUnit.TransUnit.analyzeMatch
func (self *TransUnit_TransUnit) analyzeMatch(_env *LnsEnv, firstToken *Types_Token) *Nodes_MatchNode {
    var exp *Nodes_Node
    exp = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
    var algeTypeInfo *Ast_AlgeTypeInfo
    
    {
        _algeTypeInfo := Ast_AlgeTypeInfoDownCastF(exp.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env).FP)
        if _algeTypeInfo == nil{
            self.FP.Error(_env, _env.GetVM().String_format("match must have alge value -- %s", []LnsAny{exp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
        } else {
            algeTypeInfo = _algeTypeInfo.(*Ast_AlgeTypeInfo)
        }
    }
    if Lns_op_not(self.moduleType.FP.Equals(_env, self.ProcessInfo, algeTypeInfo.FP.GetModule(_env), nil, nil)){
        if Lns_op_not(self.ImportModuleSet.Has(Ast_TypeInfo2Stem(algeTypeInfo.FP.GetModule(_env)))){
            var fullname string
            fullname = algeTypeInfo.FP.GetFullName(_env, self.TypeNameCtrl, self.FP.Get_scope(_env).FP, true)
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("need to import module -- %s (%s)", []LnsAny{fullname, algeTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    self.FP.checkNextToken(_env, "{")
    var caseList *LnsList
    caseList = NewLnsList([]LnsAny{})
    var algeValNameSet *LnsSet
    algeValNameSet = NewLnsSet([]LnsAny{})
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var firstFlag bool
    firstFlag = true
    for (nextToken.Txt == "case") {
        self.FP.checkNextToken(_env, ".")
        var valNameToken *Types_Token
        valNameToken = self.FP.getToken(_env, nil)
        self.FP.checkAlgeComp(_env, valNameToken, algeTypeInfo)
        var valInfo *Ast_AlgeValInfo
        
        {
            _valInfo := algeTypeInfo.FP.GetValInfo(_env, valNameToken.Txt)
            if _valInfo == nil{
                self.FP.Error(_env, _env.GetVM().String_format("not found val -- %s", []LnsAny{valNameToken.Txt}))
            } else {
                valInfo = _valInfo.(*Ast_AlgeValInfo)
            }
        }
        if algeValNameSet.Has(valNameToken.Txt){
            self.FP.AddErrMess(_env, valNameToken.Pos, _env.GetVM().String_format("multiple val -- %s", []LnsAny{valNameToken.Txt}))
        }
        algeValNameSet.Add(valNameToken.Txt)
        var valParamNameList *LnsList
        valParamNameList = NewLnsList([]LnsAny{})
        nextToken = self.FP.getToken(_env, nil)
        var blockScope *Ast_Scope
        blockScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
        if nextToken.Txt == "("{
            for _, _paramType := range( valInfo.FP.Get_typeList(_env).Items ) {
                paramType := _paramType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var paramName *Types_Token
                paramName = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_Or_)
                self.FP.checkShadowing(_env, paramName.Pos, paramName.Txt, self.FP.Get_scope(_env))
                var workType *Ast_TypeInfo
                workType = paramType
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( Ast_TypeInfo_isMut(_env, paramType)) &&
                    _env.SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(_env, exp.FP.Get_expType(_env)))) ).(bool)){
                    workType = self.FP.createModifier(_env, workType, Ast_MutMode__IMut)
                }
                valParamNameList.Insert(Ast_SymbolInfo2Stem(self.FP.AddLocalVar(_env, paramName.Pos, false, false, paramName.Txt, workType, Ast_MutMode__IMut, nil)))
                nextToken = self.FP.getToken(_env, nil)
                if nextToken.Txt != ","{
                    break
                }
            }
            self.FP.checkToken(_env, nextToken, ")")
        } else { 
            self.FP.Pushback(_env)
        }
        if valParamNameList.Len() != valInfo.FP.Get_typeList(_env).Len(){
            self.FP.AddErrMess(_env, valNameToken.Pos, _env.GetVM().String_format("unmatch param -- %d != %d", []LnsAny{valParamNameList.Len(), valInfo.FP.Get_typeList(_env).Len()}))
        }
        var mode LnsInt
        mode = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( firstFlag) &&
            _env.SetStackVal( TransUnit_TentativeMode__Start) ||
            _env.SetStackVal( TransUnit_TentativeMode__Merge) ).(LnsInt)
        var block *Nodes_BlockNode
        block = self.FP.analyzeBlock(_env, Nodes_BlockKind__Match, mode, blockScope, nil)
        if firstFlag{
            firstFlag = false
        }
        self.FP.PopScope(_env)
        var valRefNode *Nodes_ExpRefNode
        valRefNode = Nodes_ExpRefNode_create(_env, self.NodeManager, valNameToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_AlgeTypeInfo2Stem(valInfo.FP.Get_algeTpye(_env))}), valInfo.FP.Get_symbolInfo(_env))
        var matchCase *Nodes_MatchCase
        matchCase = NewNodes_MatchCase(_env, valInfo, valRefNode, valParamNameList, block)
        caseList.Insert(Nodes_MatchCase2Stem(matchCase))
        nextToken = self.FP.getToken(_env, nil)
    }
    var caseKind LnsInt
    if algeValNameSet.Len() == algeTypeInfo.FP.Get_valInfoNum(_env){
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
    defaultBlock,failSafeDefault = self.FP.processCaseDefault(_env, firstToken, caseKind, nextToken, caseList.Len() != 0)
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.Get_curNsInfo(_env)
    return Nodes_MatchNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), nsInfo.FP.GetNextStmtId(_env, TransUnitIF_StmtKind__Match), exp, algeTypeInfo, caseList, Nodes_NodeDownCastF(defaultBlock), caseKind, failSafeDefault)
}
// 2347: decl @lune.@base.@TransUnit.TransUnit.analyzeWhile
func (self *TransUnit_TransUnit) analyzeWhile(_env *LnsEnv, token *Types_Token) *Nodes_WhileNode {
    var refAccessSymPosList *LnsList
    refAccessSymPosList = TransUnit_clearThePosForModToRef_7_(_env, self.FP.Get_scope(_env), self.ModuleScope)
    var cond *Nodes_Node
    cond = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
    var infinit bool
    infinit = false
    if cond.FP.Get_expType(_env) == Ast_builtinTypeBool{
        {
            _literal := TransUnit_convExp1_7795(Lns_2DDD(cond.FP.GetLiteral(_env)))
            if !Lns_IsNil( _literal ) {
                literal := _literal
                switch _matchExp0 := literal.(type) {
                case *Nodes_Literal__Bool:
                val := _matchExp0.Val1
                    infinit = val
                }
            }
        }
    } else if Lns_op_not(cond.FP.Get_expType(_env).FP.Get_nilable(_env)){
        infinit = true
    }
    return Nodes_WhileNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), cond, infinit, self.FP.analyzeBlock(_env, Nodes_BlockKind__While, TransUnit_TentativeMode__Loop, nil, refAccessSymPosList))
}
// 2372: decl @lune.@base.@TransUnit.TransUnit.analyzeRepeat
func (self *TransUnit_TransUnit) analyzeRepeat(_env *LnsEnv, token *Types_Token) *Nodes_RepeatNode {
    var scope *Ast_Scope
    scope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var node *Nodes_RepeatNode
    node = Nodes_RepeatNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), self.FP.analyzeBlock(_env, Nodes_BlockKind__Repeat, TransUnit_TentativeMode__Loop, scope, nil), self.FP.analyzeExpOneRVal(_env, false, false, nil, nil))
    self.FP.finishTentativeSymbol(_env, false)
    self.FP.PopScope(_env)
    self.FP.checkNextToken(_env, ";")
    return node
}
// 2391: decl @lune.@base.@TransUnit.TransUnit.analyzeFor
func (self *TransUnit_TransUnit) analyzeFor(_env *LnsEnv, firstToken *Types_Token) *Nodes_ForNode {
    var scope *Ast_Scope
    scope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var val *Types_Token
    val = self.FP.getToken(_env, nil)
    if val.Kind != Types_TokenKind__Symb{
        self.FP.Error(_env, "not symbol")
    }
    self.FP.checkNextToken(_env, "=")
    var exp1 *Nodes_Node
    exp1 = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
    if Lns_op_not(Ast_isNumberType(_env, exp1.FP.Get_expType(_env))){
        self.FP.AddErrMess(_env, exp1.FP.Get_pos(_env), _env.GetVM().String_format("exp1 is not number -- %s", []LnsAny{exp1.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
    }
    self.FP.checkNextToken(_env, ",")
    var exp2 *Nodes_Node
    exp2 = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
    if Lns_op_not(Ast_isNumberType(_env, exp2.FP.Get_expType(_env))){
        self.FP.AddErrMess(_env, exp2.FP.Get_pos(_env), _env.GetVM().String_format("exp2 is not number -- %s", []LnsAny{exp2.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
    }
    var token *Types_Token
    token = self.FP.getToken(_env, nil)
    var exp3 LnsAny
    exp3 = nil
    if token.Txt == ","{
        exp3 = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
        {
            __exp := exp3
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_Node)
                if Lns_op_not(Ast_isNumberType(_env, _exp.FP.Get_expType(_env))){
                    self.FP.AddErrMess(_env, _exp.FP.Get_pos(_env), _env.GetVM().String_format("exp is not number -- %s", []LnsAny{_exp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        }
    } else { 
        self.FP.Pushback(_env)
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( exp1.FP.Get_expType(_env) == Ast_builtinTypeInt) &&
        _env.SetStackVal( exp2.FP.Get_expType(_env) == Ast_builtinTypeReal) ||
        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(exp3) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_Node).FP.Get_expType(_env)})) == Ast_builtinTypeReal) ).(bool){
        exp1 = &Nodes_ExpCastNode_create(_env, self.NodeManager, exp1.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeReal)}), exp1, Ast_builtinTypeReal, nil, "@@", Nodes_CastKind__Force).Nodes_Node
    }
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.Get_curNsInfo(_env)
    var symbolInfo *Ast_SymbolInfo
    symbolInfo = nsInfo.FP.RegisterSym(_env, self.FP.AddLocalVar(_env, val.Pos, false, true, val.Txt, exp1.FP.Get_expType(_env), Ast_MutMode__IMut, nil))
    var node *Nodes_ForNode
    node = Nodes_ForNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), nsInfo.FP.GetNextStmtId(_env, TransUnitIF_StmtKind__For), self.FP.analyzeBlock(_env, Nodes_BlockKind__For, TransUnit_TentativeMode__Loop, scope, nil), symbolInfo, exp1, exp2, exp3)
    self.FP.PopScope(_env)
    return node
}
// 2453: decl @lune.@base.@TransUnit.TransUnit.analyzeApply
func (self *TransUnit_TransUnit) analyzeApply(_env *LnsEnv, token *Types_Token) *Nodes_ApplyNode {
    var scope *Ast_Scope
    scope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var varList *LnsList
    varList = NewLnsList([]LnsAny{})
    var nextToken *Types_Token
    nextToken = Parser_getEofToken(_env)
    for {
        var _var *Types_Token
        _var = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_Or_)
        if _var.Kind != Types_TokenKind__Symb{
            self.FP.Error(_env, "illegal symbol")
        }
        varList.Insert(Types_Token2Stem(_var))
        nextToken = self.FP.getToken(_env, nil)
        if nextToken.Txt != ","{ break }
    }
    self.FP.checkToken(_env, nextToken, "of")
    var expListNode *Nodes_ExpListNode
    expListNode = self.FP.analyzeExpList(_env, false, false, false, nil, nil, nil)
    var itFunc *Ast_TypeInfo
    itFunc = Ast_builtinTypeNone
    var expTypeList *LnsList
    expTypeList = expListNode.FP.Get_expTypeList(_env)
    if expTypeList.Len() < 3{
        self.FP.AddErrMess(_env, expListNode.FP.Get_pos(_env), _env.GetVM().String_format("apply must have 3 values -- %s", []LnsAny{expTypeList.Len()}))
    } else { 
        itFunc = expTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_extedType(_env)
    }
    var itemTypeList *LnsList
    itemTypeList = NewLnsList([]LnsAny{})
    var defaultItemType *Ast_TypeInfo
    defaultItemType = Ast_builtinTypeStem_
    var readyFlag bool
    readyFlag = false
    {
        _callNode := Nodes_ExpCallNodeDownCastF(Nodes_getUnwraped(_env, expListNode.FP.Get_expList(_env).GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()).FP)
        if !Lns_IsNil( _callNode ) {
            callNode := _callNode.(*Nodes_ExpCallNode)
            var callFuncType *Ast_TypeInfo
            callFuncType = callNode.FP.Get_func(_env).FP.Get_expType(_env)
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( callFuncType.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.Str_gmatch, nil, nil)) ||
                _env.SetStackVal( callFuncType.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.String_gmatch, nil, nil)) ).(bool){
                itemTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeString))
                defaultItemType = Ast_builtinTypeString.FP.Get_nilableTypeInfo(_env)
                readyFlag = true
            }
        }
    }
    if Lns_op_not(readyFlag){
        if _switch0 := itFunc.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Func || _switch0 == Ast_TypeInfoKind__FormFunc || _switch0 == Ast_TypeInfoKind__Form {
        } else {
            self.FP.AddErrMess(_env, expListNode.FP.Get_pos(_env), _env.GetVM().String_format("The 1st value must be iterator function. -- %s", []LnsAny{itFunc.FP.GetTxt(_env, nil, nil, nil)}))
        }
        if itFunc.FP.Get_argTypeInfoList(_env).Len() != 2{
            self.FP.AddErrMess(_env, expListNode.FP.Get_pos(_env), _env.GetVM().String_format("iterator function must has two arguments. -- %s", []LnsAny{itFunc.FP.GetTxt(_env, nil, nil, nil)}))
        } else { 
            var arg2Type *Ast_TypeInfo
            arg2Type = itFunc.FP.Get_argTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(arg2Type.FP.Get_nilable(_env)){
                self.FP.AddErrMess(_env, expListNode.FP.Get_pos(_env), _env.GetVM().String_format("the 2nd argument of iterator function must be nilable. -- %s", []LnsAny{itFunc.FP.GetTxt(_env, nil, nil, nil)}))
            }
        }
        if itFunc.FP.Get_retTypeInfoList(_env).Len() == 0{
            self.FP.AddErrMess(_env, expListNode.FP.Get_pos(_env), "iterator function must return value.")
        } else { 
            var iteRetType *Ast_TypeInfo
            iteRetType = itFunc.FP.Get_retTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(iteRetType.FP.Get_nilable(_env)){
                self.FP.AddErrMess(_env, expListNode.FP.Get_pos(_env), "iterator function must return nilable type at 1st.")
            }
        }
        for _index, _itemType := range( itFunc.FP.Get_retTypeInfoList(_env).Items ) {
            index := _index + 1
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var workType *Ast_TypeInfo
            workType = itemType
            if index == 1{
                if itemType.FP.Get_nilable(_env){
                    workType = workType.FP.Get_nonnilableType(_env)
                }
            }
            itemTypeList.Insert(Ast_TypeInfo2Stem(workType))
        }
    }
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.Get_curNsInfo(_env)
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
        varSymList.Insert(Ast_SymbolInfo2Stem(nsInfo.FP.RegisterSym(_env, self.FP.AddLocalVar(_env, _var.Pos, false, true, _var.Txt, itemType, Ast_MutMode__IMut, nil))))
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(_env, Nodes_BlockKind__Apply, TransUnit_TentativeMode__Loop, scope, nil)
    self.FP.PopScope(_env)
    return Nodes_ApplyNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), nsInfo.FP.GetNextStmtId(_env, TransUnitIF_StmtKind__Apply), varSymList, expListNode, block)
}
// 2565: decl @lune.@base.@TransUnit.TransUnit.convToExtTypeList
func (self *TransUnit_TransUnit) convToExtTypeList(_env *LnsEnv, pos Types_Position,typeInfo *Ast_TypeInfo,list *LnsList) *LnsList {
    if typeInfo.FP.Get_nonnilableType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Ext{
        return list
    }
    var newList LnsAny
    var mess string
    newList,mess = Ast_convToExtTypeList(_env, self.ProcessInfo, list)
    if newList != nil{
        newList_675 := newList.(*LnsList)
        return newList_675
    }
    self.FP.AddErrMess(_env, pos, mess)
    return list
}
// 2581: decl @lune.@base.@TransUnit.TransUnit.analyzeForeach
func (self *TransUnit_TransUnit) analyzeForeach(_env *LnsEnv, token *Types_Token,sortFlag bool) *Nodes_Node {
    var scope *Ast_Scope
    scope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var mainSymToken *Types_Token
    mainSymToken = Parser_getEofToken(_env)
    var subSymToken LnsAny
    subSymToken = nil
    var mainSym *Ast_SymbolInfo
    var subSym LnsAny
    subSym = nil
    var nextToken *Types_Token
    nextToken = Parser_getEofToken(_env)
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = 2
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            index := _forWork0
            var symbol *Types_Token
            symbol = self.FP.getToken(_env, nil)
            if symbol.Kind != Types_TokenKind__Symb{
                self.FP.Error(_env, "illegal symbol")
            }
            if index == 1{
                mainSymToken = symbol
            } else { 
                subSymToken = symbol
            }
            nextToken = self.FP.getToken(_env, nil)
            if nextToken.Txt != ","{
                break
            }
        }
    }
    self.FP.checkToken(_env, nextToken, "in")
    var exp *Nodes_Node
    exp = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
    if exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        var nsInfo *TransUnitIF_NSInfo
        nsInfo = self.FP.Get_curNsInfo(_env)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.IsAnalyzingBlockArg(_env))) &&
            _env.SetStackVal( Lns_op_not(nsInfo.FP.CanAccessLuaval(_env))) ).(bool)){
            self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("can't access the luaval without __luago. -- %s in %s", []LnsAny{exp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil), nsInfo.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    var checkSortType func(_env *LnsEnv, sortKeyType *Ast_TypeInfo)
    checkSortType = func(_env *LnsEnv, sortKeyType *Ast_TypeInfo) {
        if sortFlag{
            if _switch0 := sortKeyType.FP.Get_srcTypeInfo(_env).FP.Get_extedType(_env); _switch0 == Ast_builtinTypeString || _switch0 == Ast_builtinTypeInt || _switch0 == Ast_builtinTypeReal || _switch0 == Ast_builtinTypeStem {
            } else {
                self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("This type can't use forsort -- %s", []LnsAny{sortKeyType.FP.GetTxt(_env, nil, nil, nil)}))
            }
        }
    }
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.Get_curNsInfo(_env)
    var expType *Ast_TypeInfo
    expType = exp.FP.Get_expType(_env).FP.Get_extedType(_env)
    var itemTypeInfoList *LnsList
    itemTypeInfoList = self.FP.convToExtTypeList(_env, token.Pos, exp.FP.Get_expType(_env), expType.FP.Get_itemTypeInfoList(_env))
    if _switch0 := expType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Map {
        mainSym = nsInfo.FP.RegisterSym(_env, self.FP.AddLocalVar(_env, mainSymToken.Pos, false, true, mainSymToken.Txt, itemTypeInfoList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil))
        {
            __exp := subSymToken
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Types_Token)
                subSym = nsInfo.FP.RegisterSym(_env, self.FP.AddLocalVar(_env, _exp.Pos, false, true, _exp.Txt, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil))
            }
        }
        checkSortType(_env, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
    } else if _switch0 == Ast_TypeInfoKind__Set {
        if subSymToken != nil{
            subSymToken_709 := subSymToken.(*Types_Token)
            self.FP.AddErrMess(_env, subSymToken_709.Pos, "Set can't use index")
        }
        mainSym = nsInfo.FP.RegisterSym(_env, self.FP.AddLocalVar(_env, mainSymToken.Pos, false, true, mainSymToken.Txt, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil))
        checkSortType(_env, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
    } else if _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        if sortFlag{
            self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("'%s' doesn't support forsort.", []LnsAny{Ast_TypeInfoKind_getTxt( expType.FP.Get_kind(_env))}))
        }
        mainSym = nsInfo.FP.RegisterSym(_env, self.FP.AddLocalVar(_env, mainSymToken.Pos, false, true, mainSymToken.Txt, itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil))
        {
            __exp := subSymToken
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Types_Token)
                subSym = nsInfo.FP.RegisterSym(_env, self.FP.AddLocalVar(_env, _exp.Pos, false, false, _exp.Txt, Ast_builtinTypeInt, Ast_MutMode__IMut, nil))
            }
        }
    } else {
        self.FP.ErrorAt(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("unknown kind type of exp for foreach -- %s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil)}))
    }
    var seqSym LnsAny
    seqSym = nil
    {
        _refNode := Nodes_ExpRefNodeDownCastF(exp.FP)
        if !Lns_IsNil( _refNode ) {
            refNode := _refNode.(*Nodes_ExpRefNode)
            var seqSymbol *Ast_SymbolInfo
            seqSymbol = refNode.FP.Get_symbolInfo(_env)
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( seqSymbol.FP.Get_mutable(_env)) ||
                _env.SetStackVal( Ast_TypeInfo_isMut(_env, seqSymbol.FP.Get_typeInfo(_env))) ).(bool){
                scope.FP.AddOverrideImut(_env, self.ProcessInfo, seqSymbol)
                seqSym = seqSymbol.FP.Get_name(_env)
            }
        }
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(_env, Nodes_BlockKind__Foreach, TransUnit_TentativeMode__Loop, scope, nil)
    if seqSym != nil{
        seqSym_722 := seqSym.(string)
        scope.FP.Remove(_env, seqSym_722)
    }
    self.FP.PopScope(_env)
    if sortFlag{
        return &Nodes_ForsortNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), nsInfo.FP.GetNextStmtId(_env, TransUnitIF_StmtKind__Forsort), mainSym, subSym, exp, block, sortFlag).Nodes_Node
    } else { 
        return &Nodes_ForeachNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), nsInfo.FP.GetNextStmtId(_env, TransUnitIF_StmtKind__Foreach), mainSym, subSym, exp, block).Nodes_Node
    }
// insert a dummy
    return nil
}
// 2720: decl @lune.@base.@TransUnit.TransUnit.analyzeScope
func (self *TransUnit_TransUnit) analyzeScope(_env *LnsEnv, firstToken *Types_Token) *Nodes_ScopeNode {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var scopeKind LnsInt
    if _switch0 := nextToken.Txt; _switch0 == "root" {
        scopeKind = Nodes_ScopeKind__Root
    } else {
        self.FP.Error(_env, _env.GetVM().String_format("illegal scope kind. -- %s", []LnsAny{nextToken.Txt}))
    }
    var symList *LnsList
    symList = NewLnsList([]LnsAny{})
    nextToken = self.FP.getToken(_env, nil)
    if nextToken.Txt == "("{
        nextToken = self.FP.getToken(_env, nil)
        for nextToken.Txt != ")" {
            var symbolNode *Nodes_Node
            symbolNode = self.FP.analyzeExpSymbol(_env, nextToken, nextToken, TransUnit_ExpSymbolMode__Symbol, nil, true, false)
            var workSymList *LnsList
            workSymList = symbolNode.FP.GetSymbolInfo(_env)
            if workSymList.Len() > 0{
                symList.Insert(Ast_SymbolInfo2Stem(workSymList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()))
            }
            nextToken = self.FP.getToken(_env, nil)
            if _switch1 := nextToken.Txt; _switch1 == ")" {
            } else if _switch1 == "," {
                nextToken = self.FP.getToken(_env, nil)
            } else {
                self.FP.Error(_env, _env.GetVM().String_format("illegal token: expects ')' or ',' but -- %s", []LnsAny{nextToken.Txt}))
            }
        }
    } else { 
        self.FP.Pushback(_env)
    }
    var asyncMode LnsInt
    asyncMode = self.FP.getDefaultAsync(_env, Ast_TypeInfoKind__Func, nil, nil)
    var bakScope *Ast_Scope
    bakScope = self.FP.Get_scope(_env)
    self.FP.SetScope(_env, self.TopScope, TransUnitIF_SetNSInfo__FromScope)
    var localScope *Ast_Scope
    localScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    self.FP.createDummyNS(_env, localScope, nextToken.Pos, asyncMode)
    for _, _symInfo := range( symList.Items ) {
        symInfo := _symInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        self.FP.Get_scope(_env).FP.AddAlias(_env, self.ProcessInfo, symInfo.FP.Get_name(_env), nextToken.Pos, false, symInfo.FP.Get_accessMode(_env), symInfo.FP.Get_typeInfo(_env).FP.Get_parentInfo(_env), symInfo)
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(_env, Nodes_BlockKind__Block, TransUnit_TentativeMode__Simple, nil, nil)
    var node *Nodes_ScopeNode
    node = Nodes_ScopeNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), scopeKind, self.FP.Get_scope(_env), symList, block)
    self.FP.SetScope(_env, bakScope, TransUnitIF_SetNSInfo__FromScope)
    return node
}
// 34: decl @lune.@base.@TransUnit.TransUnit.analyzeRefType
func (self *TransUnit_TransUnit) analyzeRefType(_env *LnsEnv, accessMode LnsInt,allowDDD bool,parentPub bool) *Nodes_RefTypeNode {
    var firstToken *Types_Token
    firstToken = self.FP.getToken(_env, nil)
    var token *Types_Token
    token = firstToken
    var mutMode LnsAny
    if _switch0 := token.Txt; _switch0 == "&" {
        mutMode = Ast_MutMode__IMut
        token = self.FP.getToken(_env, nil)
    } else if _switch0 == "allmut" {
        mutMode = Ast_MutMode__AllMut
        token = self.FP.getToken(_env, nil)
    } else if _switch0 == "#" {
        mutMode = Ast_MutMode__Depend
        token = self.FP.getToken(_env, nil)
    } else {
        mutMode = nil
    }
    var name *Nodes_Node
    if token.Txt == "..."{
        var dddSym *Ast_SymbolInfo
        dddSym = Lns_unwrap( self.ModuleScope.FP.GetSymbolInfo(_env, "...", self.ModuleScope, true, Ast_ScopeAccess__Normal)).(*Ast_SymbolInfo)
        name = &Nodes_ExpRefNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddSym.FP.Get_typeInfo(_env))}), &NewAst_AccessSymbolInfo(_env, self.ProcessInfo, dddSym, Ast_OverrideMut__None_Obj, true).Ast_SymbolInfo).Nodes_Node
    } else { 
        name = self.FP.analyzeExpSymbol(_env, firstToken, token, TransUnit_ExpSymbolMode__Symbol, nil, true, false)
        var symbolList *LnsList
        symbolList = name.FP.GetSymbolInfo(_env)
        if symbolList.Len() > 0{
            var symbol *Ast_SymbolInfo
            symbol = symbolList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if symbol.FP.Get_kind(_env) != Ast_SymbolKind__Typ{
                self.FP.AddErrMess(_env, name.FP.Get_pos(_env), _env.GetVM().String_format("illegal type -- %s", []LnsAny{symbol.FP.Get_name(_env)}))
            }
        } else { 
            self.FP.AddErrMess(_env, name.FP.Get_pos(_env), _env.GetVM().String_format("illegal symbol node -- %s", []LnsAny{Nodes_getNodeKindName(_env, name.FP.Get_kind(_env))}))
        }
    }
    return self.FP.analyzeRefTypeWithSymbol(_env, accessMode, allowDDD, mutMode, name, parentPub)
}
// 92: decl @lune.@base.@TransUnit.TransUnit.analyzeRefTypeWithSymbol
func (self *TransUnit_TransUnit) analyzeRefTypeWithSymbol(_env *LnsEnv, accessMode LnsInt,allowDDD bool,mutMode LnsAny,symbolNode *Nodes_Node,parentPub bool) *Nodes_RefTypeNode {
    var typeInfo *Ast_TypeInfo
    typeInfo = symbolNode.FP.Get_expType(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Set) &&
        _env.SetStackVal( Lns_op_not(self.HelperInfo.UseSet)) ).(bool)){
        self.HelperInfo.UseSet = true
    }
    {
        _aliasType := Ast_AliasTypeInfoDownCastF(typeInfo.FP)
        if !Lns_IsNil( _aliasType ) {
            aliasType := _aliasType.(*Ast_AliasTypeInfo)
            var aliasSrc *Ast_TypeInfo
            aliasSrc = aliasType.FP.Get_aliasSrcTypeInfo(_env)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(self.ImportModuleSet.Has(Ast_TypeInfo2Stem(aliasSrc.FP.GetModule(_env))))) &&
                _env.SetStackVal( self.moduleType.FP.Get_parentInfo(_env) != aliasSrc.FP.GetModule(_env).FP.Get_parentInfo(_env)) ).(bool)){
                self.FP.AddErrMess(_env, symbolNode.FP.Get_pos(_env), _env.GetVM().String_format("must import '%s' for this alias -- %s (%s,%s)", []LnsAny{aliasSrc.FP.GetModule(_env).FP.GetFullName(_env, self.TypeNameCtrl, self.FP.Get_scope(_env).FP, false), symbolNode.FP.GetSymbolInfo(_env).GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_name(_env), self.moduleType.FP.Get_typeId(_env).Id, aliasSrc.FP.GetModule(_env).FP.Get_typeId(_env).Id}))
            }
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( parentPub) &&
        _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) &&
        _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, typeInfo.FP.Get_accessMode(_env)))) ).(bool)){
        self.FP.AddErrMess(_env, symbolNode.FP.Get_pos(_env), _env.GetVM().String_format("This type must be public. -- %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
    }
    var continueToken *Types_Token
    var continueFlag bool
    continueToken,continueFlag = self.FP.getContinueToken(_env)
    var token *Types_Token
    token = continueToken
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( continueFlag) &&
        _env.SetStackVal( token.Txt == "!") ).(bool)){
        typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
        token = self.FP.getToken(_env, nil)
    }
    var itemNodeList *LnsList
    itemNodeList = NewLnsList([]LnsAny{})
    var arrayMode string
    arrayMode = "no"
    for  {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Txt == "[") ||
            _env.SetStackVal( token.Txt == "[@") ).(bool){
            if token.Txt == "["{
                arrayMode = "list"
                typeInfo = self.ProcessInfo.FP.CreateList(_env, accessMode, self.FP.getCurrentClass(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), Ast_MutMode__Mut)
            } else { 
                arrayMode = "array"
                typeInfo = self.ProcessInfo.FP.CreateArray(_env, accessMode, self.FP.getCurrentClass(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), Ast_MutMode__Mut)
            }
            token = self.FP.getToken(_env, nil)
            if token.Txt != "]"{
                self.FP.Pushback(_env)
                self.FP.checkNextToken(_env, "]")
            }
            itemNodeList.Insert(Nodes_Node2Stem(symbolNode))
        } else if token.Txt == "<"{
            var genericList *LnsList
            genericList = NewLnsList([]LnsAny{})
            var nextToken *Types_Token
            nextToken = Parser_getEofToken(_env)
            for {
                var typeExp *Nodes_RefTypeNode
                typeExp = self.FP.analyzeRefType(_env, accessMode, false, parentPub)
                itemNodeList.Insert(Nodes_RefTypeNode2Stem(typeExp))
                genericList.Insert(Ast_TypeInfo2Stem(typeExp.FP.Get_expType(_env)))
                if typeExp.FP.Get_expType(_env).FP.Get_mutMode(_env) == Ast_MutMode__Depend{
                    self.FP.AddErrMess(_env, typeExp.FP.Get_pos(_env), _env.GetVM().String_format("type parameter can't have dep attribute. -- %s", []LnsAny{typeExp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
                nextToken = self.FP.getToken(_env, nil)
                if nextToken.Txt != ","{ break }
            }
            self.FP.checkToken(_env, nextToken, ">")
            var checkAlternateTypeCount func(_env *LnsEnv, count LnsInt) bool
            checkAlternateTypeCount = func(_env *LnsEnv, count LnsInt) bool {
                if genericList.Len() != count{
                    self.FP.AddErrMess(_env, symbolNode.FP.Get_pos(_env), _env.GetVM().String_format("generic type count is unmatch. -- %d", []LnsAny{genericList.Len()}))
                    return false
                }
                return true
            }
            if _switch0 := typeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Map {
                if genericList.Len() != 2{
                    self.FP.AddErrMess(_env, symbolNode.FP.Get_pos(_env), "Key or value type is unknown")
                    typeInfo = self.ProcessInfo.FP.CreateMap(_env, accessMode, self.FP.getCurrentClass(_env), Ast_builtinTypeStem, Ast_builtinTypeStem, Ast_MutMode__Mut)
                } else { 
                    typeInfo = self.ProcessInfo.FP.CreateMap(_env, accessMode, self.FP.getCurrentClass(_env), genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), genericList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__Mut)
                }
            } else if _switch0 == Ast_TypeInfoKind__List {
                if checkAlternateTypeCount(_env, 1){
                    typeInfo = self.ProcessInfo.FP.CreateList(_env, accessMode, self.FP.getCurrentClass(_env), genericList, Ast_MutMode__Mut)
                }
            } else if _switch0 == Ast_TypeInfoKind__Array {
                if checkAlternateTypeCount(_env, 1){
                    typeInfo = self.ProcessInfo.FP.CreateArray(_env, accessMode, self.FP.getCurrentClass(_env), genericList, Ast_MutMode__Mut)
                }
            } else if _switch0 == Ast_TypeInfoKind__Set {
                if checkAlternateTypeCount(_env, 1){
                    typeInfo = self.ProcessInfo.FP.CreateSet(_env, accessMode, self.FP.getCurrentClass(_env), genericList, Ast_MutMode__Mut)
                }
            } else if _switch0 == Ast_TypeInfoKind__DDD {
                if checkAlternateTypeCount(_env, 1){
                    typeInfo = &self.ProcessInfo.FP.CreateDDD(_env, genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), false, false).Ast_TypeInfo
                }
            } else if _switch0 == Ast_TypeInfoKind__Class || _switch0 == Ast_TypeInfoKind__IF {
                if checkAlternateTypeCount(_env, typeInfo.FP.Get_itemTypeInfoList(_env).Len()){
                    for _, _itemType := range( genericList.Items ) {
                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if itemType.FP.Get_nilable(_env){
                            self.FP.AddErrMess(_env, symbolNode.FP.Get_pos(_env), _env.GetVM().String_format("can't use nilable type -- %s", []LnsAny{itemType.FP.GetTxt(_env, nil, nil, nil)}))
                        }
                    }
                    for _index, _itemType := range( typeInfo.FP.Get_itemTypeInfoList(_env).Items ) {
                        index := _index + 1
                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( itemType.FP.HasBase(_env)) &&
                            _env.SetStackVal( Lns_op_not(genericList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.IsInheritFrom(_env, self.ProcessInfo, itemType.FP.Get_baseTypeInfo(_env), nil))) ).(bool)){
                            self.FP.AddErrMess(_env, symbolNode.FP.Get_pos(_env), _env.GetVM().String_format("'%s' of %s doesn't inherit '%s'", []LnsAny{itemType.FP.GetTxt(_env, nil, nil, nil), typeInfo.FP.GetTxt(_env, nil, nil, nil), itemType.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
                        }
                    }
                    typeInfo = TransUnit_convExp2_878(Lns_2DDD(self.ProcessInfo.FP.CreateGeneric(_env, typeInfo, genericList, self.moduleType)))
                }
            } else if _switch0 == Ast_TypeInfoKind__Box {
                if checkAlternateTypeCount(_env, 1){
                    typeInfo = self.ProcessInfo.FP.CreateBox(_env, accessMode, genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                }
            } else if _switch0 == Ast_TypeInfoKind__Ext {
                if checkAlternateTypeCount(_env, 1){
                    typeInfo = self.FP.createExtType(_env, symbolNode.FP.Get_pos(_env), genericList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                }
            } else {
                self.FP.Error(_env, _env.GetVM().String_format("not support generic: %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            }
        } else { 
            self.FP.Pushback(_env)
            break
        }
        token = self.FP.getToken(_env, nil)
    }
    if token.Txt == "!"{
        typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
        self.FP.getToken(_env, nil)
    }
    if Lns_op_not(allowDDD){
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            self.FP.AddErrMess(_env, symbolNode.FP.Get_pos(_env), _env.GetVM().String_format("invalid type. -- '%s'", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    if mutMode != nil{
        mutMode_60 := mutMode.(LnsInt)
        if typeInfo.FP.Get_mutMode(_env) != mutMode_60{
            typeInfo = self.FP.createModifier(_env, typeInfo, mutMode_60)
        }
    }
    if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Module{
        self.FP.AddErrMess(_env, symbolNode.FP.Get_pos(_env), _env.GetVM().String_format("module can't use as Type. -- %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
    }
    return Nodes_RefTypeNode_create(_env, self.NodeManager, symbolNode.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), symbolNode, itemNodeList, mutMode, arrayMode)
}
// 306: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclArgList
func (self *TransUnit_TransUnit) analyzeDeclArgList(_env *LnsEnv, accessMode LnsInt,scope *Ast_Scope,argList *LnsList,parentPub bool) *Types_Token {
    var nextToken *Types_Token
    nextToken = Parser_noneToken
    var hasDDDFlag bool
    hasDDDFlag = false
    for {
        nextToken = self.FP.getToken(_env, nil)
        if nextToken.Txt == ")"{
            break
        }
        if hasDDDFlag{
            self.FP.AddErrMess(_env, nextToken.Pos, "Argument exists after '...'.")
        }
        var mutable LnsInt
        mutable = Ast_MutMode__IMut
        if nextToken.Txt == "mut"{
            mutable = Ast_MutMode__Mut
            nextToken = self.FP.getToken(_env, nil)
        }
        var argName *Types_Token
        argName = nextToken
        if argName.Txt == "..."{
            hasDDDFlag = true
            var workToken *Types_Token
            var flag bool
            workToken,flag = self.FP.getContinueToken(_env)
            self.FP.Pushback(_env)
            var dddTypeInfo *Ast_TypeInfo
            dddTypeInfo = Ast_builtinTypeDDD
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( flag) &&
                _env.SetStackVal( workToken.Txt == "<") ).(bool)){
                self.FP.PushbackToken(_env, nextToken)
                var refTypeNode *Nodes_RefTypeNode
                refTypeNode = self.FP.analyzeRefType(_env, accessMode, true, parentPub)
                dddTypeInfo = refTypeNode.FP.Get_expType(_env)
            }
            argList.Insert(Nodes_DeclArgDDDNode2Stem(Nodes_DeclArgDDDNode_create(_env, self.NodeManager, argName.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddTypeInfo)}))))
            scope.FP.AddLocalVar(_env, self.ProcessInfo, true, true, argName.Txt, argName.Pos, dddTypeInfo, Ast_MutMode__IMut)
        } else { 
            argName = self.FP.checkSymbol(_env, argName, TransUnit_SymbolMode__MustNot_)
            self.FP.checkShadowing(_env, argName.Pos, argName.Txt, scope)
            self.FP.checkNextToken(_env, ":")
            var refType *Nodes_RefTypeNode
            refType = self.FP.analyzeRefType(_env, accessMode, false, parentPub)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( refType.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
                _env.SetStackVal( refType.FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env).Len() > 0) ).(bool)){
                var argType *Ast_TypeInfo
                argType = refType.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env)
                if Lns_op_not(Ast_isGenericType(_env, argType)){
                    self.FP.AddErrMess(_env, refType.FP.Get_pos(_env), _env.GetVM().String_format("can't use this type without <T>. please use %s.", []LnsAny{argType.FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
            {
                _symbolInfo := TransUnit_convExp2_1550(Lns_2DDD(scope.FP.AddLocalVar(_env, self.ProcessInfo, true, true, argName.Txt, argName.Pos, refType.FP.Get_expType(_env), mutable)))
                if !Lns_IsNil( _symbolInfo ) {
                    symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                    var arg *Nodes_DeclArgNode
                    arg = Nodes_DeclArgNode_create(_env, self.NodeManager, argName.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), refType.FP.Get_expTypeList(_env), argName, symbolInfo, refType)
                    argList.Insert(Nodes_DeclArgNode2Stem(arg))
                }
            }
        }
        nextToken = self.FP.getToken(_env, nil)
        if nextToken.Txt != ","{ break }
    }
    self.FP.checkToken(_env, nextToken, ")")
    return nextToken
}
// 388: decl @lune.@base.@TransUnit.TransUnit.checkOverrideMethod
func (self *TransUnit_TransUnit) checkOverrideMethod(_env *LnsEnv, overrideType *Ast_TypeInfo,typeInfo *Ast_TypeInfo) *LnsList {
    var accessMode LnsInt
    accessMode = typeInfo.FP.Get_accessMode(_env)
    var funcName string
    funcName = typeInfo.FP.Get_rawTxt(_env)
    var altTypeList *LnsList
    altTypeList = typeInfo.FP.Get_itemTypeInfoList(_env)
    var alt2typeMap *LnsMap
    alt2typeMap = typeInfo.FP.Get_parentInfo(_env).FP.CreateAlt2typeMap(_env, false)
    var errList *LnsList
    errList = NewLnsList([]LnsAny{})
    var addErr func(_env *LnsEnv, mess string)
    addErr = func(_env *LnsEnv, mess string) {
        var fullName string
        fullName = _env.GetVM().String_format("%s.%s", []LnsAny{typeInfo.FP.Get_parentInfo(_env).FP.Get_rawTxt(_env), typeInfo.FP.Get_rawTxt(_env)})
        errList.Insert(_env.GetVM().String_format("%s: %s: %s -- %s", []LnsAny{fullName, mess, typeInfo.FP.Get_display_stirng(_env), typeInfo.FP.Get_display_stirng(_env)}))
    }
    if self.Ctrl_info.ValidAsyncCtrl{
        if overrideType.FP.Get_asyncMode(_env) != typeInfo.FP.Get_asyncMode(_env){
            addErr(_env, _env.GetVM().String_format("mismatch asyncMode --  %s, %s", []LnsAny{Ast_Async_getTxt( overrideType.FP.Get_asyncMode(_env)), Ast_Async_getTxt( typeInfo.FP.Get_asyncMode(_env))}))
        }
    }
    if overrideType.FP.Get_accessMode(_env) != accessMode{
        var mess string
        mess = _env.GetVM().String_format("mismatch override accessMode -- %s,%s", []LnsAny{Ast_AccessMode_getTxt( overrideType.FP.Get_accessMode(_env)), Ast_AccessMode_getTxt( accessMode)})
        addErr(_env, mess)
    }
    if overrideType.FP.Get_staticFlag(_env) != typeInfo.FP.Get_staticFlag(_env){
        addErr(_env, "mismatch override staticFlag -- " + funcName)
    }
    if overrideType.FP.Get_kind(_env) != Ast_TypeInfoKind__Method{
        addErr(_env, _env.GetVM().String_format("mismatch override kind -- %s, %d", []LnsAny{funcName, overrideType.FP.Get_kind(_env)}))
    }
    if overrideType.FP.Get_mutMode(_env) != typeInfo.FP.Get_mutMode(_env){
        addErr(_env, _env.GetVM().String_format("mismatch mutable -- %s", []LnsAny{funcName}))
    }
    if overrideType.FP.Get_itemTypeInfoList(_env).Len() != altTypeList.Len(){
        var mess string
        mess = _env.GetVM().String_format("mismatch altTypeList -- %d, %d", []LnsAny{overrideType.FP.Get_itemTypeInfoList(_env).Len(), altTypeList.Len()})
        addErr(_env, mess)
    } else { 
        for _index, _alterType := range( overrideType.FP.Get_itemTypeInfoList(_env).Items ) {
            index := _index + 1
            alterType := _alterType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            alt2typeMap.Set(alterType,altTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
        }
    }
    var matchFlag bool
    var err LnsAny
    matchFlag,err = overrideType.FP.CanEvalWith(_env, self.ProcessInfo, typeInfo, Ast_CanEvalType__SetEq, alt2typeMap)
    if Lns_op_not(matchFlag){
        if err != nil{
            err_114 := err.(string)
            addErr(_env, _env.GetVM().String_format("mismatch method type -- %s", []LnsAny{err_114}))
        } else {
            addErr(_env, "mismatch method type")
        }
    }
    for _index, _retType := range( overrideType.FP.Get_retTypeInfoList(_env).Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if typeInfo.FP.Get_retTypeInfoList(_env).Len() >= index{
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( retType.FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                _env.SetStackVal( typeInfo.FP.Get_retTypeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nonnilableType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Alternate) ).(bool)){
                var mess string
                mess = _env.GetVM().String_format("not support to override the method has generics at return type. -- %s", []LnsAny{funcName})
                addErr(_env, mess)
            }
        }
    }
    return errList
}
// 465: decl @lune.@base.@TransUnit.TransUnit.checkOverriededMethodOfAllClass
func (self *TransUnit_TransUnit) checkOverriededMethodOfAllClass(_env *LnsEnv) {
    var process func(_env *LnsEnv, pos Types_Position,alt2typeMap *LnsMap,classScope *Ast_Scope,superScope *Ast_Scope)
    process = func(_env *LnsEnv, pos Types_Position,alt2typeMap *LnsMap,classScope *Ast_Scope,superScope *Ast_Scope) {
        superScope.FP.FilterTypeInfoField(_env, true, classScope, self.ScopeAccess, Ast_filterForm(func(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) bool {
            if symbolInfo.FP.Get_name(_env) == "__init"{
                return true
            }
            var implimented bool
            implimented = true
            if symbolInfo.FP.Get_typeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
                {
                    _impMethodType := classScope.FP.GetTypeInfoField(_env, symbolInfo.FP.Get_name(_env), true, classScope, self.ScopeAccess)
                    if !Lns_IsNil( _impMethodType ) {
                        impMethodType := _impMethodType.(*Ast_TypeInfo)
                        if impMethodType == symbolInfo.FP.Get_typeInfo(_env){
                            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( symbolInfo.FP.Get_typeInfo(_env).FP.Get_abstractFlag(_env)) &&
                                _env.SetStackVal( classScope != superScope) ).(bool)){
                                implimented = false
                            }
                        } else { 
                            for _, _err := range( self.FP.checkOverrideMethod(_env, symbolInfo.FP.Get_typeInfo(_env), impMethodType).Items ) {
                                err := _err.(string)
                                self.FP.AddErrMess(_env, pos, err)
                            }
                        }
                    } else {
                        implimented = false
                    }
                }
            }
            if Lns_op_not(implimented){
                self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("not implements method -- %s.%s at %s", []LnsAny{_env.NilAccFin(_env.NilAccPush(superScope.FP.Get_ownerTypeInfo(_env)) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetTxt(_env, nil, nil, nil)})/* 499:35 */), symbolInfo.FP.Get_name(_env), _env.NilAccFin(_env.NilAccPush(classScope.FP.Get_ownerTypeInfo(_env)) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.GetTxt(_env, nil, nil, nil)})/* 501:35 */)}))
            }
            return true
        }))
    }
    var typeId2DeclClassNode *LnsMap
    typeId2DeclClassNode = NewLnsMap( map[LnsAny]LnsAny{})
    for _classTypeInfo, _classNode := range( self.typeInfo2ClassNode.Items ) {
        classTypeInfo := _classTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        classNode := _classNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        typeId2DeclClassNode.Set(classTypeInfo.FP.Get_typeId(_env).Id,classNode)
    }
    {
        __forsortCollection0 := typeId2DeclClassNode
        __forsortSorted0 := __forsortCollection0.CreateKeyListInt()
        __forsortSorted0.Sort( _env, LnsItemKindInt, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            classNode := __forsortCollection0.Items[ ___forsortKey0 ].(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            var classTypeInfo *Ast_TypeInfo
            classTypeInfo = classNode.FP.Get_expType(_env)
            var workTypeInfo *Ast_TypeInfo
            workTypeInfo = classTypeInfo
            var alt2typeMap *LnsMap
            alt2typeMap = classTypeInfo.FP.CreateAlt2typeMap(_env, false)
            for {
                if Lns_op_not(classTypeInfo.FP.Get_abstractFlag(_env)){
                    if workTypeInfo != Ast_headTypeInfo{
                        process(_env, classNode.FP.Get_pos(_env), alt2typeMap, Lns_unwrap( classTypeInfo.FP.Get_scope(_env)).(*Ast_Scope), Lns_unwrap( workTypeInfo.FP.Get_scope(_env)).(*Ast_Scope))
                    }
                }
                for _, _ifType := range( workTypeInfo.FP.Get_interfaceList(_env).Items ) {
                    ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if ifType != Ast_builtinTypeMapping{
                        process(_env, classNode.FP.Get_pos(_env), alt2typeMap, Lns_unwrap( classTypeInfo.FP.Get_scope(_env)).(*Ast_Scope), Lns_unwrap( ifType.FP.Get_scope(_env)).(*Ast_Scope))
                    }
                }
                workTypeInfo = workTypeInfo.FP.Get_baseTypeInfo(_env)
                if workTypeInfo == Ast_headTypeInfo{ break }
            }
        }
    }
}
// 561: decl @lune.@base.@TransUnit.TransUnit.declFuncPostProcess
func (self *TransUnit_TransUnit) declFuncPostProcess(_env *LnsEnv, typeInfo *Ast_TypeInfo,classTypeInfo LnsAny,workBody *Nodes_BlockNode,funcBodyScope *Ast_Scope) {
    if funcBodyScope.FP.Get_closureSymList(_env).Len() > 0{
        funcBodyScope.FP.Set_hasClosureAccess(_env, true)
    }
    if classTypeInfo != nil{
        classTypeInfo_172 := classTypeInfo.(*Ast_TypeInfo)
        var isCtorFlag bool
        isCtorFlag = typeInfo.FP.Get_rawTxt(_env) == "__init"
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( isCtorFlag) &&
            _env.SetStackVal( classTypeInfo_172.FP.HasBase(_env)) ).(bool)){
            var needCall bool
            needCall = true
            {
                _firstNode := TransUnit_getFirstStmt_25_(_env, workBody.FP.Get_stmtList(_env))
                if !Lns_IsNil( _firstNode ) {
                    firstNode := _firstNode.(*Nodes_Node)
                    if firstNode.FP.Get_kind(_env) == Nodes_nodeKindEnum__ExpCallSuperCtor{
                        needCall = false
                    }
                }
            }
            if needCall{
                self.FP.AddErrMess(_env, workBody.FP.Get_pos(_env), "__init must call super() with first.")
            }
        }
    }
}
// 614: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclMacroSub
func (self *TransUnit_TransUnit) analyzeDeclMacroSub(_env *LnsEnv, accessMode LnsInt,firstToken *Types_Token,nameToken *Types_Token,macroScope *Ast_Scope,parentType *Ast_TypeInfo,typeDataAccessor Ast_TypeDataAccessor,workArgList *LnsList) *Nodes_DeclMacroNode {
    if self.MacroCtrl.FP.Get_isDeclaringMacro(_env){
        self.FP.Error(_env, "can't declare the macro in the macro.")
    }
    var pubFlag bool
    pubFlag = false
    if _switch0 := accessMode; _switch0 == Ast_AccessMode__Pub {
        pubFlag = true
    } else if _switch0 == Ast_AccessMode__Local || _switch0 == Ast_AccessMode__None {
    } else {
        self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("macro not support this access mode. -- %s", []LnsAny{Ast_AccessMode_getTxt( accessMode)}))
    }
    var argList *LnsList
    argList = NewLnsList([]LnsAny{})
    var argTypeList *LnsList
    argTypeList = NewLnsList([]LnsAny{})
    for _, _argNode := range( workArgList.Items ) {
        argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
        {
            __exp := Nodes_DeclArgNodeDownCastF(argNode.FP)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_DeclArgNode)
                argList.Insert(Nodes_DeclArgNode2Stem(_exp))
            } else {
                self.FP.Error(_env, "macro argument can not use '...'.")
            }
        }
        var argType *Ast_TypeInfo
        argType = argNode.FP.Get_expType(_env)
        argTypeList.Insert(Ast_TypeInfo2Stem(argType))
    }
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var retTypeList *LnsList
    if nextToken.Txt == ":"{
        retTypeList = self.FP.analyzeRefType(_env, accessMode, true, false).FP.Get_expTypeList(_env)
        self.FP.checkNextToken(_env, "{")
    } else { 
        retTypeList = NewLnsList([]LnsAny{})
        self.FP.checkToken(_env, nextToken, "{")
    }
    var typeInfo *Ast_NormalTypeInfo
    typeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, macroScope, Ast_TypeInfoKind__Macro, parentType, typeDataAccessor, false, false, true, accessMode, nameToken.Txt, Ast_Async__Async, nil, argTypeList, retTypeList, Ast_MutMode__IMut)
    self.MacroCtrl.FP.StartDecl(_env, &typeInfo.Ast_TypeInfo)
    nextToken = self.FP.getToken(_env, nil)
    var stmtNode LnsAny
    stmtNode = nil
    if nextToken.Txt == "{"{
        self.macroScope = macroScope
        macroScope.FP.Set_validCheckingUnaccess(_env, false)
        var funcType *Ast_TypeInfo
        funcType = Ast_builtinTypeLnsLoad
        macroScope.FP.AddLocalVar(_env, self.ProcessInfo, false, false, "_lnsLoad", nil, funcType, Ast_MutMode__IMut)
        var macroLocalVarType *Ast_TypeInfo
        macroLocalVarType = self.ProcessInfo.FP.CreateMap(_env, Ast_AccessMode__Local, self.moduleType, Ast_builtinTypeString, Ast_builtinTypeStem, Ast_MutMode__Mut)
        if Lns_op_not(pubFlag){
            macroScope.FP.AddLocalVar(_env, self.ProcessInfo, false, true, "__var", nil, macroLocalVarType, Ast_MutMode__IMut)
        }
        var stmtList *LnsList
        stmtList = NewLnsList([]LnsAny{})
        self.FP.prepareTentativeSymbol(_env, self.FP.Get_scope(_env), false, nil)
        self.FP.analyzeStatementList(_env, stmtList, false, "}")
        if stmtList.Len() > 0{
            stmtNode = Nodes_BlockNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), Nodes_BlockKind__Macro, macroScope, stmtList)
        }
        self.FP.checkNextToken(_env, "}")
        self.FP.finishTentativeSymbol(_env, true)
        self.macroScope = nil
    } else { 
        self.FP.Pushback(_env)
    }
    var tokenList *LnsList
    tokenList = NewLnsList([]LnsAny{})
    var braceCount LnsInt
    braceCount = 0
    for  {
        nextToken = self.FP.getToken(_env, nil)
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
    var declMacroInfo *Nodes_DeclMacroInfo
    declMacroInfo = NewNodes_DeclMacroInfo(_env, pubFlag, nameToken, argList, stmtNode, tokenList)
    var node *Nodes_DeclMacroNode
    node = Nodes_DeclMacroNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_NormalTypeInfo2Stem(typeInfo)}), declMacroInfo)
    {
        __exp := self.MacroCtrl.FP.Regist(_env, self.ProcessInfo, node, macroScope, self.BaseDir)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(string)
            self.FP.ErrorAt(_env, nameToken.Pos, _exp)
        }
    }
    return node
}
// 772: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclMacro
func (self *TransUnit_TransUnit) analyzeDeclMacro(_env *LnsEnv, accessMode LnsInt,firstToken *Types_Token) *Nodes_DeclMacroNode {
    var nameToken *Types_Token
    nameToken = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__Must_)
    self.FP.checkNextToken(_env, "(")
    var scope *Ast_Scope
    scope = Ast_TypeInfo_createScope(_env, self.ProcessInfo, self.TopScope, Ast_ScopeKind__Other, nil, nil)
    self.FP.createDummyNS(_env, scope, nameToken.Pos, Ast_Async__Noasync)
    var workArgList *LnsList
    workArgList = NewLnsList([]LnsAny{})
    self.FP.analyzeDeclArgList(_env, accessMode, scope, workArgList, false)
    var parentNsInfo *TransUnitIF_NSInfo
    parentNsInfo = self.FP.Get_curNsInfo(_env)
    var parentInfo *Ast_TypeInfo
    parentInfo = parentNsInfo.FP.Get_typeInfo(_env)
    var backScope *Ast_Scope
    backScope = self.FP.Get_scope(_env)
    self.FP.SetScope(_env, scope, TransUnitIF_SetNSInfo__FromScope)
    self.FP.Get_scope(_env).FP.AddIgnoredVar(_env, self.ProcessInfo)
    var node *Nodes_DeclMacroNode
    node = self.FP.analyzeDeclMacroSub(_env, accessMode, firstToken, nameToken, scope, parentInfo, parentNsInfo.FP.Get_typeDataAccessor(_env), workArgList)
    self.FP.SetScope(_env, backScope, TransUnitIF_SetNSInfo__FromScope)
    var existSym LnsAny
    _,existSym = self.FP.Get_scope(_env).FP.AddMacro(_env, self.ProcessInfo, nameToken.Pos, node.FP.Get_expType(_env), accessMode)
    if Lns_isCondTrue( existSym){
        self.FP.AddErrMess(_env, nameToken.Pos, _env.GetVM().String_format("multiple define symbol -- %s", []LnsAny{nameToken.Txt}))
    }
    return node
}
// 823: decl @lune.@base.@TransUnit.TransUnit.analyzeExtend
func (self *TransUnit_TransUnit) analyzeExtend(_env *LnsEnv, accessMode LnsInt,firstPos Types_Position)(*Types_Token, LnsAny, *LnsList, *LnsMap, *Nodes_ClassInheritInfo) {
    var baseRef LnsAny
    baseRef = nil
    var interfaceList *LnsList
    interfaceList = NewLnsList([]LnsAny{})
    var ifAlt2typeMap *LnsMap
    ifAlt2typeMap = NewLnsMap( map[LnsAny]LnsAny{})
    var ifRefList *LnsList
    ifRefList = NewLnsList([]LnsAny{})
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    if nextToken.Txt != "("{
        self.FP.Pushback(_env)
        var workBaseRefType *Nodes_RefTypeNode
        workBaseRefType = self.FP.analyzeRefType(_env, accessMode, false, Ast_isPubToExternal(_env, accessMode))
        baseRef = workBaseRefType
        var baseType *Ast_TypeInfo
        baseType = workBaseRefType.FP.Get_expType(_env)
        if baseType.FP.Get_kind(_env) != Ast_TypeInfoKind__Class{
            self.FP.AddErrMess(_env, workBaseRefType.FP.Get_pos(_env), _env.GetVM().String_format("%s is not class.", []LnsAny{baseType.FP.GetTxt(_env, nil, nil, nil)}))
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) &&
            _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, baseType.FP.Get_accessMode(_env)))) ).(bool)){
            self.FP.AddErrMess(_env, workBaseRefType.FP.Get_pos(_env), _env.GetVM().String_format("%s can't be external symbol.", []LnsAny{baseType.FP.GetTxt(_env, nil, nil, nil)}))
        }
        nextToken = self.FP.getToken(_env, nil)
    }
    if nextToken.Txt == "("{
        for  {
            nextToken = self.FP.getToken(_env, nil)
            if nextToken.Txt == ")"{
                break
            }
            self.FP.Pushback(_env)
            var ifTypeNode *Nodes_RefTypeNode
            ifTypeNode = self.FP.analyzeRefType(_env, accessMode, false, Ast_isPubToExternal(_env, accessMode))
            ifRefList.Insert(Nodes_RefTypeNode2Stem(ifTypeNode))
            var ifType *Ast_TypeInfo
            ifType = ifTypeNode.FP.Get_expType(_env)
            if ifType.FP.Get_kind(_env) != Ast_TypeInfoKind__IF{
                self.FP.Error(_env, _env.GetVM().String_format("%s is not interface -- %d", []LnsAny{ifType.FP.GetTxt(_env, nil, nil, nil), ifType.FP.Get_kind(_env)}))
            }
            if Ast_isGenericType(_env, ifType){
                for _altType, _genType := range( ifType.FP.CreateAlt2typeMap(_env, false).Items ) {
                    altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    ifAlt2typeMap.Set(altType,genType)
                }
            }
            interfaceList.Insert(Ast_TypeInfo2Stem(ifType))
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) &&
                _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, ifType.FP.Get_accessMode(_env)))) ).(bool)){
                self.FP.AddErrMess(_env, ifTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("%s can't be external symbol.", []LnsAny{ifType.FP.GetTxt(_env, nil, nil, nil)}))
            }
            nextToken = self.FP.getToken(_env, nil)
            if nextToken.Txt != ","{
                if nextToken.Txt == ")"{
                    break
                }
                self.FP.Error(_env, "illegal token")
            }
        }
        nextToken = self.FP.getToken(_env, nil)
    }
    var symbol2TypeInfo *LnsMap
    symbol2TypeInfo = NewLnsMap( map[LnsAny]LnsAny{})
    for _, _ifType := range( interfaceList.Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if ifType == Ast_builtinTypeAbsImmut{
            if baseRef != nil{
                baseRef_257 := baseRef.(*Nodes_RefTypeNode)
                self.FP.AddErrMess(_env, baseRef_257.FP.Get_pos(_env), _env.GetVM().String_format("__absimmut can't extend the class -- %s", []LnsAny{baseRef_257.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
            }
        }
        _env.NilAccFin(_env.NilAccPush(ifType.FP.Get_scope(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.FilterTypeInfoField(_env, true, self.FP.Get_scope(_env), self.ScopeAccess, Ast_filterForm(func(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) bool {
            if symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mtd{
                {
                    _ifFuncType := symbol2TypeInfo.Get(symbolInfo.FP.Get_name(_env))
                    if !Lns_IsNil( _ifFuncType ) {
                        ifFuncType := _ifFuncType.(*Ast_TypeInfo)
                        var ret bool
                        var mess LnsAny
                        ret,mess = ifFuncType.FP.CanEvalWith(_env, self.ProcessInfo, symbolInfo.FP.Get_typeInfo(_env), Ast_CanEvalType__SetOp, ifAlt2typeMap)
                        if Lns_op_not(ret){
                            self.FP.AddErrMess(_env, firstPos, _env.GetVM().String_format("mismatch method type -- %s.%s, %s.%s\n%s", []LnsAny{symbolInfo.FP.Get_typeInfo(_env).FP.Get_parentInfo(_env).FP.GetTxt(_env, nil, nil, nil), symbolInfo.FP.Get_name(_env), ifFuncType.FP.Get_parentInfo(_env).FP.GetTxt(_env, nil, nil, nil), ifFuncType.FP.GetTxt(_env, nil, nil, nil), mess}))
                        }
                    } else {
                        symbol2TypeInfo.Set(symbolInfo.FP.Get_name(_env),symbolInfo.FP.Get_typeInfo(_env))
                    }
                }
            }
            return true
        }))})/* 905:7 */)
    }
    var baseTypeInfo LnsAny
    baseTypeInfo = nil
    if baseRef != nil{
        baseRef_270 := baseRef.(*Nodes_RefTypeNode)
        baseTypeInfo = baseRef_270.FP.Get_expType(_env)
        if baseRef_270.FP.Get_expType(_env).FP.IsInheritFrom(_env, self.ProcessInfo, Ast_builtinTypeAbsImmut, nil){
            self.FP.AddErrMess(_env, baseRef_270.FP.Get_pos(_env), _env.GetVM().String_format("can't extend the __absimmut. (%s)", []LnsAny{baseRef_270.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    return nextToken, baseTypeInfo, interfaceList, ifAlt2typeMap, NewNodes_ClassInheritInfo(_env, baseRef, ifRefList)
}
// 945: decl @lune.@base.@TransUnit.TransUnit.analyzePushClass
func (self *TransUnit_TransUnit) analyzePushClass(_env *LnsEnv, mode LnsInt,abstractFlag bool,firstToken *Types_Token,name *Types_Token,allowMultiple bool,requirePath LnsAny,moduleLang LnsAny,accessMode LnsInt,altTypeList *LnsList)(*Types_Token, *TransUnitIF_NSInfo, *Nodes_ClassInheritInfo) {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) &&
        _env.SetStackVal( self.ModuleScope != self.FP.Get_scope(_env)) ).(bool)){
        self.FP.AddErrMess(_env, firstToken.Pos, "The public class must declare at top scope.")
    }
    var tempScope *Ast_Scope
    tempScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    for _, _altType := range( altTypeList.Items ) {
        altType := _altType.(Ast_AlternateTypeInfoDownCast).ToAst_AlternateTypeInfo()
        tempScope.FP.AddAlternate(_env, self.ProcessInfo, accessMode, altType.FP.Get_rawTxt(_env), name.Pos, &altType.Ast_TypeInfo)
    }
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var baseTypeInfo LnsAny
    baseTypeInfo = nil
    var interfaceList LnsAny
    interfaceList = nil
    var inheritInfo *Nodes_ClassInheritInfo
    if nextToken.Txt == "extend"{
        nextToken, baseTypeInfo, interfaceList, _, inheritInfo = self.FP.analyzeExtend(_env, accessMode, firstToken.Pos)
        if baseTypeInfo != nil{
            baseTypeInfo_283 := baseTypeInfo.(*Ast_TypeInfo)
            {
                _initTypeInfo := _env.NilAccFin(_env.NilAccPush(baseTypeInfo_283.FP.Get_scope(_env)) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetTypeInfoChild(_env, "__init")})/* 974:33 */)
                if !Lns_IsNil( _initTypeInfo ) {
                    initTypeInfo := _initTypeInfo.(*Ast_TypeInfo)
                    if initTypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pri{
                        self.FP.AddErrMess(_env, firstToken.Pos, "The access mode of '__init' is 'pri'.")
                    }
                }
            }
        }
    } else { 
        inheritInfo = NewNodes_ClassInheritInfo(_env, nil, NewLnsList([]LnsAny{}))
    }
    self.FP.PopScope(_env)
    var nsInfo *TransUnitIF_NSInfo
    if _switch0 := mode; _switch0 == TransUnitIF_DeclClassMode__Module || _switch0 == TransUnitIF_DeclClassMode__LazyModule {
        _ = self.FP.Get_scope(_env)
        nsInfo = self.FP.pushExtModule(_env, false, name.Txt, accessMode, name.Pos, mode == TransUnitIF_DeclClassMode__LazyModule, Lns_unwrap( moduleLang).(LnsInt), (Lns_unwrap( requirePath).(*Types_Token)).FP.GetExcludedDelimitTxt(_env))
    } else if _switch0 == TransUnitIF_DeclClassMode__Class || _switch0 == TransUnitIF_DeclClassMode__Interface {
        nsInfo = self.FP.PushClass(_env, self.ProcessInfo, firstToken.Pos, mode, abstractFlag, baseTypeInfo, interfaceList, altTypeList, false, name.Txt, allowMultiple, accessMode, nil)
    }
    return nextToken, nsInfo, inheritInfo
}
// 1007: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclAlternateType
func (self *TransUnit_TransUnit) analyzeDeclAlternateType(_env *LnsEnv, belongClassFlag bool,token *Types_Token,accessMode LnsInt)(*Types_Token, *LnsList) {
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
        genericSymToken = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
        if Lns_isCondTrue( self.FP.Get_scope(_env).FP.GetTypeInfo(_env, genericSymToken.Txt, self.FP.Get_scope(_env), false, self.ScopeAccess)){
            self.FP.AddErrMess(_env, genericSymToken.Pos, _env.GetVM().String_format("shadowing Type -- %s", []LnsAny{genericSymToken.Txt}))
        } else { 
            if altNameSet.Has(genericSymToken.Txt){
                self.FP.AddErrMess(_env, genericSymToken.Pos, _env.GetVM().String_format("multiple Type -- %s", []LnsAny{genericSymToken.Txt}))
            } else { 
                altNameSet.Add(genericSymToken.Txt)
            }
        }
        var workToken *Types_Token
        workToken = self.FP.getToken(_env, nil)
        if workToken.Txt == "!"{
            self.FP.AddErrMess(_env, workToken.Pos, "not support nilable")
            workToken = self.FP.getToken(_env, nil)
        }
        var baseTypeInfo LnsAny
        baseTypeInfo = nil
        var interfaceList *LnsList
        interfaceList = NewLnsList([]LnsAny{})
        if workToken.Txt == ":"{
            workToken, baseTypeInfo, interfaceList = TransUnit_convExp2_4528(Lns_2DDD(self.FP.analyzeExtend(_env, accessMode, token.Pos)))
        }
        var altType *Ast_AlternateTypeInfo
        altType = TransUnit_convExp2_4557(Lns_2DDD(self.ProcessInfo.FP.CreateAlternate(_env, belongClassFlag, altIndex, genericSymToken.Txt, accessMode, self.moduleType, baseTypeInfo, interfaceList)))
        altTypeList.Insert(Ast_AlternateTypeInfo2Stem(altType))
        if workToken.Txt == ">"{
            nextToken = self.FP.getToken(_env, nil)
            break
        }
        self.FP.checkToken(_env, workToken, ",")
    }
    return nextToken, altTypeList
}
// 1059: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclProto
func (self *TransUnit_TransUnit) analyzeDeclProto(_env *LnsEnv, accessMode LnsInt,firstToken *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var abstractFlag bool
    abstractFlag = false
    if nextToken.Txt == "abstract"{
        abstractFlag = true
        nextToken = self.FP.getToken(_env, nil)
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( nextToken.Txt == "class") ||
        _env.SetStackVal( nextToken.Txt == "interface") ).(bool){
        var name *Types_Token
        name = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
        var altTypeList *LnsList
        altTypeList = NewLnsList([]LnsAny{})
        {
            var workToken *Types_Token
            workToken = self.FP.getToken(_env, nil)
            if workToken.Txt == "<"{
                workToken, altTypeList = self.FP.analyzeDeclAlternateType(_env, true, workToken, accessMode)
            }
            self.FP.PushbackToken(_env, workToken)
        }
        if accessMode == Ast_AccessMode__Local{
            accessMode = Ast_AccessMode__Pri
        }
        var declMode LnsInt
        if nextToken.Txt == "class"{
            declMode = TransUnitIF_DeclClassMode__Class
        } else { 
            declMode = TransUnitIF_DeclClassMode__Interface
            abstractFlag = true
        }
        var classTypeInfo *Ast_TypeInfo
        var inheritInfo *Nodes_ClassInheritInfo
        var nsInfo *TransUnitIF_NSInfo
        nextToken, nsInfo, inheritInfo = self.FP.analyzePushClass(_env, declMode, abstractFlag, firstToken, name, false, nil, nil, accessMode, altTypeList)
        classTypeInfo = nsInfo.FP.Get_typeInfo(_env)
        nsInfo.FP.Set_nobody(_env, true)
        self.FP.PopClass(_env)
        self.FP.checkToken(_env, nextToken, ";")
        return &Nodes_ProtoClassNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo)}), name, inheritInfo).Nodes_Node
    }
    self.FP.Error(_env, "illegal proto")
// insert a dummy
    return nil
}
// 1115: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclEnum
func (self *TransUnit_TransUnit) analyzeDeclEnum(_env *LnsEnv, accessMode LnsInt,firstToken *Types_Token) *Nodes_DeclEnumNode {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) &&
        _env.SetStackVal( self.FP.Get_scope(_env) != self.ModuleScope) &&
        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(self.FP.Get_scope(_env).FP.Get_ownerTypeInfo(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) != Ast_TypeInfoKind__Class) ).(bool)){
        self.FP.AddErrMess(_env, firstToken.Pos, "can't external at inner scope.")
    }
    var name *Types_Token
    name = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
    self.FP.checkNextToken(_env, "{")
    var valueList *LnsList
    valueList = NewLnsList([]LnsAny{})
    var scope *Ast_Scope
    scope = self.FP.PushScope(_env, Ast_ScopeKind__Class, nil, nil)
    var workEnumTypeInfo LnsAny
    workEnumTypeInfo = nil
    var parentNsInfo *TransUnitIF_NSInfo
    parentNsInfo = self.FP.Get_curNsInfo(_env)
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var number LnsReal
    number = 0.0
    var prevValTypeInfo *Ast_TypeInfo
    prevValTypeInfo = Ast_headTypeInfo
    var valTypeInfo *Ast_TypeInfo
    valTypeInfo = Ast_headTypeInfo
    for nextToken.Txt != "}" {
        var valName *Types_Token
        valName = self.FP.checkSymbol(_env, nextToken, TransUnit_SymbolMode__MustNot_)
        nextToken = self.FP.getToken(_env, nil)
        var enumVal LnsAny
        enumVal = &Ast_EnumLiteral__Real{number}
        if _switch0 := (prevValTypeInfo); _switch0 == Ast_builtinTypeReal {
        } else if _switch0 == Ast_builtinTypeInt || _switch0 == Ast_headTypeInfo {
            enumVal = &Ast_EnumLiteral__Int{(LnsInt)(number)}
        }
        if nextToken.Txt == "="{
            var exp *Nodes_Node
            exp = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
            var literal LnsAny
            var mess LnsAny
            literal,mess = exp.FP.GetLiteral(_env)
            if literal != nil{
                literal_348 := literal
                switch _matchExp0 := literal_348.(type) {
                case *Nodes_Literal__Int:
                val := _matchExp0.Val1
                    enumVal = &Ast_EnumLiteral__Int{val}
                    number = (LnsReal)(val)
                    valTypeInfo = Ast_builtinTypeInt
                case *Nodes_Literal__Real:
                val := _matchExp0.Val1
                    enumVal = &Ast_EnumLiteral__Real{val}
                    number = val
                    valTypeInfo = Ast_builtinTypeReal
                case *Nodes_Literal__Str:
                val := _matchExp0.Val1
                    enumVal = &Ast_EnumLiteral__Str{val}
                    valTypeInfo = Ast_builtinTypeString
                default:
                    self.FP.Error(_env, _env.GetVM().String_format("illegal enum val -- %s", []LnsAny{literal_348.(LnsAlgeVal).GetTxt()}))
                }
            } else {
                self.FP.Error(_env, _env.GetVM().String_format("illegal enum val -- %s", []LnsAny{mess}))
            }
            nextToken = self.FP.getToken(_env, nil)
        } else { 
            if _switch1 := (prevValTypeInfo); _switch1 == Ast_headTypeInfo {
                valTypeInfo = Ast_builtinTypeInt
            } else if _switch1 == Ast_builtinTypeInt || _switch1 == Ast_builtinTypeReal {
                valTypeInfo = prevValTypeInfo
            } else {
                self.FP.AddErrMess(_env, valName.Pos, _env.GetVM().String_format("illegal enum val type -- %s", []LnsAny{valTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( prevValTypeInfo != Ast_headTypeInfo) &&
            _env.SetStackVal( prevValTypeInfo != valTypeInfo) ).(bool)){
            self.FP.AddErrMess(_env, valName.Pos, _env.GetVM().String_format("multiple enum val type. %s, %s", []LnsAny{valTypeInfo.FP.GetTxt(_env, nil, nil, nil), prevValTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
        }
        prevValTypeInfo = valTypeInfo
        if Lns_op_not(workEnumTypeInfo){
            workEnumTypeInfo = self.ProcessInfo.FP.CreateEnum(_env, scope, parentNsInfo.FP.Get_typeInfo(_env), parentNsInfo.FP.Get_typeDataAccessor(_env), false, accessMode, name.Txt, valTypeInfo)
        }
        if workEnumTypeInfo != nil{
            workEnumTypeInfo_364 := workEnumTypeInfo.(*Ast_EnumTypeInfo)
            var evalValSym *Ast_SymbolInfo
            evalValSym = Lns_unwrap( Lns_car(scope.FP.AddEnumVal(_env, self.ProcessInfo, valName.Txt, valName.Pos, &workEnumTypeInfo_364.Ast_TypeInfo))).(*Ast_SymbolInfo)
            var enumValInfo *Ast_EnumValInfo
            enumValInfo = NewAst_EnumValInfo(_env, valName.Txt, enumVal, evalValSym)
            valueList.Insert(Types_Token2Stem(valName))
            workEnumTypeInfo_364.FP.AddEnumValInfo(_env, enumValInfo)
        }
        if nextToken.Txt == "}"{
            break
        }
        self.FP.checkToken(_env, nextToken, ",")
        nextToken = self.FP.getToken(_env, nil)
        number = number + LnsReal(1)
    }
    var enumTypeInfo *Ast_EnumTypeInfo
    
    {
        _enumTypeInfo := workEnumTypeInfo
        if _enumTypeInfo == nil{
            enumTypeInfo = self.ProcessInfo.FP.CreateEnum(_env, scope, parentNsInfo.FP.Get_typeInfo(_env), parentNsInfo.FP.Get_typeDataAccessor(_env), false, accessMode, name.Txt, Ast_builtinTypeNone)
        } else {
            enumTypeInfo = _enumTypeInfo.(*Ast_EnumTypeInfo)
        }
    }
    self.FP.PopScope(_env)
    var shadowing LnsAny
    _,shadowing = self.FP.Get_scope(_env).FP.AddEnum(_env, self.ProcessInfo, accessMode, name.Txt, name.Pos, &enumTypeInfo.Ast_TypeInfo)
    self.FP.errorShadowing(_env, name.Pos, shadowing)
    return Nodes_DeclEnumNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_EnumTypeInfo2Stem(enumTypeInfo)}), enumTypeInfo, accessMode, name, valueList, scope)
}
// 1254: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclAlge
func (self *TransUnit_TransUnit) analyzeDeclAlge(_env *LnsEnv, accessMode LnsInt,firstToken *Types_Token) *Nodes_DeclAlgeNode {
    self.HelperInfo.UseAlge = true
    var name *Types_Token
    name = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
    self.FP.checkNextToken(_env, "{")
    var scope *Ast_Scope
    scope = self.FP.Get_scope(_env)
    var algeScope *Ast_Scope
    algeScope = self.FP.PushScope(_env, Ast_ScopeKind__Class, nil, nil)
    var parentNsInfo *TransUnitIF_NSInfo
    parentNsInfo = self.FP.Get_curNsInfo(_env)
    var algeTypeInfo *Ast_AlgeTypeInfo
    algeTypeInfo = self.ProcessInfo.FP.CreateAlge(_env, algeScope, parentNsInfo.FP.Get_typeInfo(_env), parentNsInfo.FP.Get_typeDataAccessor(_env), false, accessMode, name.Txt)
    var shadowing LnsAny
    _,shadowing = scope.FP.AddAlge(_env, self.ProcessInfo, accessMode, name.Txt, name.Pos, &algeTypeInfo.Ast_TypeInfo)
    self.FP.NewNSInfo(_env, &algeTypeInfo.Ast_TypeInfo, name.Pos)
    self.FP.errorShadowing(_env, name.Pos, shadowing)
    var algeValList *LnsList
    algeValList = NewLnsList([]LnsAny{})
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    for nextToken.Txt != "}" {
        var valName *Types_Token
        valName = self.FP.checkSymbol(_env, nextToken, TransUnit_SymbolMode__MustNot_)
        if Lns_isCondTrue( algeTypeInfo.FP.GetValInfo(_env, valName.Txt)){
            self.FP.AddErrMess(_env, valName.Pos, _env.GetVM().String_format("multiple symbole -- %s", []LnsAny{valName.Txt}))
        }
        nextToken = self.FP.getToken(_env, nil)
        var paramList *LnsList
        paramList = NewLnsList([]LnsAny{})
        var typeInfoList *LnsList
        typeInfoList = NewLnsList([]LnsAny{})
        if nextToken.Txt == "("{
            for  {
                var paramNameToken LnsAny
                var workToken1 *Types_Token
                workToken1 = self.FP.getToken(_env, nil)
                var workToken2 *Types_Token
                workToken2 = self.FP.getToken(_env, nil)
                if workToken2.Txt != ":"{
                    self.FP.Pushback(_env)
                    self.FP.Pushback(_env)
                    paramNameToken = nil
                } else { 
                    paramNameToken = workToken1
                }
                var typeNode *Nodes_RefTypeNode
                typeNode = self.FP.analyzeRefType(_env, Ast_AccessMode__Pub, false, Ast_isPubToExternal(_env, accessMode))
                if Lns_isCondTrue( _env.NilAccFin(_env.NilAccPush(self.NsInfoMap.Get(typeNode.FP.Get_expType(_env))) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*TransUnitIF_NSInfo).FP.Get_nobody(_env)}))){
                    self.FP.AddErrMess(_env, typeNode.FP.Get_pos(_env), _env.GetVM().String_format("can't use the prototype class -- %s", []LnsAny{typeNode.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
                typeInfoList.Insert(Ast_TypeInfo2Stem(typeNode.FP.Get_expType(_env)))
                paramList.Insert(Nodes_AlgeValParamInfo2Stem(NewNodes_AlgeValParamInfo(_env, paramNameToken, typeNode)))
                nextToken = self.FP.getToken(_env, nil)
                if nextToken.Txt != ","{
                    self.FP.checkToken(_env, nextToken, ")")
                    nextToken = self.FP.getToken(_env, nil)
                    break
                }
            }
        }
        var workAlgeValSym LnsAny
        var algeValSymShadow LnsAny
        workAlgeValSym,algeValSymShadow = algeScope.FP.AddAlgeVal(_env, self.ProcessInfo, valName.Txt, valName.Pos, &algeTypeInfo.Ast_TypeInfo)
        self.FP.errorShadowing(_env, valName.Pos, algeValSymShadow)
        var algeValSym *Ast_SymbolInfo
        algeValSym = Lns_unwrap( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( workAlgeValSym) ||
            _env.SetStackVal( shadowing) ))).(*Ast_SymbolInfo)
        var algeValInfo *Ast_AlgeValInfo
        algeValInfo = NewAst_AlgeValInfo(_env, valName.Txt, typeInfoList, algeTypeInfo, algeValSym)
        algeTypeInfo.FP.AddValInfo(_env, algeValInfo)
        algeValList.Insert(Nodes_DeclAlgeValInfo2Stem(NewNodes_DeclAlgeValInfo(_env, algeValSym, paramList)))
        if nextToken.Txt == "}"{
            break
        }
        self.FP.checkToken(_env, nextToken, ",")
        nextToken = self.FP.getToken(_env, nil)
    }
    self.FP.PopScope(_env)
    return Nodes_DeclAlgeNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_AlgeTypeInfo2Stem(algeTypeInfo)}), accessMode, algeTypeInfo, name, algeValList, algeScope)
}
// 1350: decl @lune.@base.@TransUnit.TransUnit.analyzeAlias
func (self *TransUnit_TransUnit) analyzeAlias(_env *LnsEnv, accessMode LnsInt,firstToken *Types_Token) *Nodes_AliasNode {
    if self.FP.Get_scope(_env) != self.ModuleScope{
        self.FP.AddErrMess(_env, firstToken.Pos, "alias must use at top scope.")
    }
    var newToken *Types_Token
    newToken = self.FP.getToken(_env, nil)
    self.FP.checkNextToken(_env, "=")
    var srcToken *Types_Token
    srcToken = self.FP.getToken(_env, nil)
    var symbolNode *Nodes_Node
    symbolNode = self.FP.analyzeExpSymbol(_env, srcToken, srcToken, TransUnit_ExpSymbolMode__Symbol, nil, true, false)
    var newTypeInfo *Ast_TypeInfo
    newTypeInfo = Ast_builtinTypeNone
    var symbolInfoList *LnsList
    symbolInfoList = symbolNode.FP.GetSymbolInfo(_env)
    var newSymbolInfo *Ast_SymbolInfo
    newSymbolInfo = Ast_dummySymbol
    if symbolInfoList.Len() >= 1{
        var symbolInfo *Ast_SymbolInfo
        symbolInfo = symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_car(_env.GetVM().String_find(newToken.Txt,"^_", nil, nil))) &&
            _env.SetStackVal( Lns_op_not(Lns_car(_env.GetVM().String_find(srcToken.Txt,"^_", nil, nil)))) ||
            _env.SetStackVal( Lns_op_not(Lns_car(_env.GetVM().String_find(newToken.Txt,"^_", nil, nil)))) &&
            _env.SetStackVal( Lns_car(_env.GetVM().String_find(srcToken.Txt,"^_", nil, nil))) )){
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("alias symbol unmatch. %s %s", []LnsAny{newToken.Txt, newToken.Txt}))
        } else { 
            if _switch0 := symbolInfo.FP.Get_kind(_env); _switch0 == Ast_SymbolKind__Typ || _switch0 == Ast_SymbolKind__Fun {
                var aliasSymbolInfo LnsAny
                var shadowing LnsAny
                aliasSymbolInfo,shadowing = self.FP.Get_scope(_env).FP.AddAlias(_env, self.ProcessInfo, newToken.Txt, newToken.Pos, false, accessMode, self.moduleType, symbolInfo)
                if aliasSymbolInfo != nil{
                    aliasSymbolInfo_419 := aliasSymbolInfo.(*Ast_SymbolInfo)
                    newTypeInfo = aliasSymbolInfo_419.FP.Get_typeInfo(_env)
                    newSymbolInfo = aliasSymbolInfo_419
                } else {
                    self.FP.errorShadowing(_env, newToken.Pos, shadowing)
                }
            } else {
                self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("can alias symbol -- %s. (%s)", []LnsAny{srcToken.Txt, Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind(_env))}))
            }
        }
    } else { 
        self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("not found symbold -- %s", []LnsAny{srcToken.Txt}))
    }
    self.FP.checkNextToken(_env, ";")
    return Nodes_AliasNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(newTypeInfo)}), newSymbolInfo, symbolNode, newTypeInfo)
}
// 1415: decl @lune.@base.@TransUnit.TransUnit.analyzeRetTypeList
func (self *TransUnit_TransUnit) analyzeRetTypeList(_env *LnsEnv, pubToExtFlag bool,accessMode LnsInt,token *Types_Token,parentPub bool)(*LnsList, *Types_Token, *LnsList) {
    var retTypeInfoList *LnsList
    retTypeInfoList = NewLnsList([]LnsAny{})
    var retTypeNodeList *LnsList
    retTypeNodeList = NewLnsList([]LnsAny{})
    if token.Txt == ":"{
        var hasDDDFlag bool
        hasDDDFlag = false
        for  {
            var refTypeNode *Nodes_RefTypeNode
            refTypeNode = self.FP.analyzeRefType(_env, accessMode, true, parentPub)
            if hasDDDFlag{
                self.FP.AddErrMess(_env, refTypeNode.FP.Get_pos(_env), "Type exists after '...'.")
            }
            var retType *Ast_TypeInfo
            retType = refTypeNode.FP.Get_expType(_env)
            if retType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                hasDDDFlag = true
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( pubToExtFlag) &&
                _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, retType.FP.Get_accessMode(_env)))) ).(bool)){
                self.FP.AddErrMess(_env, refTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("this is not public type -- %s", []LnsAny{retType.FP.GetTxt(_env, nil, nil, nil)}))
            }
            retTypeInfoList.Insert(Ast_TypeInfo2Stem(retType))
            retTypeNodeList.Insert(Nodes_RefTypeNode2Stem(refTypeNode))
            token = self.FP.getToken(_env, nil)
            if token.Txt != ","{
                break
            }
        }
    }
    return retTypeInfoList, token, retTypeNodeList
}
// 1450: decl @lune.@base.@TransUnit.TransUnit.getMutableAsync
func (self *TransUnit_TransUnit) getMutableAsync(_env *LnsEnv, token *Types_Token)(*Types_Token, LnsInt, LnsAny) {
    var asyncMode LnsAny
    asyncMode = nil
    if _switch0 := token.Txt; _switch0 == "__async" {
        asyncMode = Ast_Async__Async
        token = self.FP.getToken(_env, nil)
    } else if _switch0 == "__noasync" {
        asyncMode = Ast_Async__Noasync
        token = self.FP.getToken(_env, nil)
    } else if _switch0 == "__trans" {
        asyncMode = Ast_Async__Transient
        token = self.FP.getToken(_env, nil)
    }
    var mutMode LnsInt
    if _switch1 := token.Txt; _switch1 == "mut" {
        mutMode = Ast_MutMode__Mut
        token = self.FP.getToken(_env, nil)
    } else if _switch1 == "dep" {
        mutMode = Ast_MutMode__Depend
        token = self.FP.getToken(_env, nil)
    } else {
        mutMode = Ast_MutMode__IMut
    }
    return token, mutMode, asyncMode
}
// 1487: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclForm
func (self *TransUnit_TransUnit) analyzeDeclForm(_env *LnsEnv, accessMode LnsInt,firstToken *Types_Token) *Nodes_DeclFormNode {
    var name *Types_Token
    name = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.Get_scope(_env) != self.ModuleScope) &&
        _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) ).(bool)){
        self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("You must declare this form at the outside scope. -- %s", []LnsAny{name.Txt}))
    }
    self.FP.checkNextToken(_env, "(")
    var argList *LnsList
    argList = NewLnsList([]LnsAny{})
    var funcBodyScope *Ast_Scope
    funcBodyScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var nextToken *Types_Token
    nextToken = self.FP.analyzeDeclArgList(_env, accessMode, funcBodyScope, argList, Ast_isPubToExternal(_env, accessMode))
    self.FP.checkToken(_env, nextToken, ")")
    var asyncMode LnsAny
    nextToken, _, asyncMode = self.FP.getMutableAsync(_env, self.FP.getToken(_env, nil))
    var retTypeList *LnsList
    retTypeList = NewLnsList([]LnsAny{})
    var retNodeList *LnsList
    retTypeList, nextToken, retNodeList = self.FP.analyzeRetTypeList(_env, Ast_isPubToExternal(_env, accessMode), accessMode, nextToken, Ast_isPubToExternal(_env, accessMode))
    self.FP.checkToken(_env, nextToken, ";")
    self.FP.PopScope(_env)
    var argTypeInfoList *LnsList
    argTypeInfoList = NewLnsList([]LnsAny{})
    for _, _argNode := range( argList.Items ) {
        argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
        argTypeInfoList.Insert(Ast_TypeInfo2Stem(argNode.FP.Get_expType(_env)))
    }
    var parentNsInfo *TransUnitIF_NSInfo
    parentNsInfo = self.FP.Get_curNsInfo(_env)
    var formType *Ast_NormalTypeInfo
    formType = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, nil, Ast_TypeInfoKind__FormFunc, parentNsInfo.FP.Get_typeInfo(_env), parentNsInfo.FP.Get_typeDataAccessor(_env), false, false, true, accessMode, name.Txt, self.FP.getDefaultAsync(_env, Ast_TypeInfoKind__FormFunc, self.FP.getCurrentClass(_env), asyncMode), nil, argTypeInfoList, retTypeList, Ast_MutMode__IMut)
    var formSymbol LnsAny
    var shadowing LnsAny
    formSymbol,shadowing = self.FP.Get_scope(_env).FP.AddForm(_env, self.ProcessInfo, name.Pos, &formType.Ast_TypeInfo, accessMode)
    self.FP.errorShadowing(_env, name.Pos, shadowing)
    var declFuncInfo *Nodes_DeclFuncInfo
    declFuncInfo = NewNodes_DeclFuncInfo(_env, Nodes_FuncKind__Form, nil, nil, false, name, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( formSymbol) ||
        _env.SetStackVal( shadowing) ), argList, false, accessMode, asyncMode, nil, retTypeList, retNodeList, false, false, 0)
    return Nodes_DeclFormNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_NormalTypeInfo2Stem(formType)}), declFuncInfo)
}
// 1552: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclToken
func (self *TransUnit_TransUnit) AnalyzeDeclToken(_env *LnsEnv, accessMode LnsInt,staticFlag bool,firstToken *Types_Token,token *Types_Token) LnsAny {
    return nil
}
// 1559: decl @lune.@base.@TransUnit.TransUnit.analyzeDecl
func (self *TransUnit_TransUnit) analyzeDecl(_env *LnsEnv, accessMode LnsInt,staticFlag bool,firstToken *Types_Token,token *Types_Token) LnsAny {
    if Lns_op_not(staticFlag){
        if token.Txt == "static"{
            staticFlag = true
            token = self.FP.getToken(_env, nil)
        }
    }
    var overrideFlag bool
    overrideFlag = false
    if token.Txt == "override"{
        overrideFlag = true
        token = self.FP.getToken(_env, nil)
    }
    var abstractFlag bool
    abstractFlag = false
    if token.Txt == "abstract"{
        abstractFlag = true
        token = self.FP.getToken(_env, nil)
    }
    if token.Txt == "let"{
        return self.FP.analyzeDeclVar(_env, Nodes_DeclVarMode__Let, accessMode, firstToken)
    } else if token.Txt == "fn"{
        var nextToken *Types_Token
        nextToken = self.FP.getToken(_env, nil)
        self.FP.Pushback(_env)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( nextToken.Kind == Types_TokenKind__Symb) ||
            _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) ||
            _env.SetStackVal( staticFlag) ||
            _env.SetStackVal( overrideFlag) ||
            _env.SetStackVal( abstractFlag) ).(bool){
            return self.FP.analyzeDeclFunc(_env, TransUnit_DeclFuncMode__Func, false, abstractFlag, overrideFlag, accessMode, staticFlag, nil, firstToken, nil)
        }
    } else if token.Txt == "class"{
        return &self.FP.analyzeDeclClass(_env, abstractFlag, accessMode, firstToken, TransUnitIF_DeclClassMode__Class).Nodes_Node
    } else if token.Txt == "interface"{
        return &self.FP.analyzeDeclClass(_env, true, accessMode, firstToken, TransUnitIF_DeclClassMode__Interface).Nodes_Node
    } else if token.Txt == "module"{
        return &self.FP.analyzeDeclClass(_env, false, accessMode, firstToken, TransUnitIF_DeclClassMode__Module).Nodes_Node
    } else if token.Txt == "proto"{
        return self.FP.analyzeDeclProto(_env, accessMode, firstToken)
    } else if token.Txt == "macro"{
        return &self.FP.analyzeDeclMacro(_env, accessMode, firstToken).Nodes_Node
    } else if token.Txt == "enum"{
        return &self.FP.analyzeDeclEnum(_env, accessMode, firstToken).Nodes_Node
    } else if token.Txt == "alge"{
        return &self.FP.analyzeDeclAlge(_env, accessMode, firstToken).Nodes_Node
    } else if token.Txt == "form"{
        return &self.FP.analyzeDeclForm(_env, accessMode, firstToken).Nodes_Node
    } else if token.Txt == "alias"{
        return &self.FP.analyzeAlias(_env, accessMode, firstToken).Nodes_Node
    }
    return self.FP.AnalyzeDeclToken(_env, accessMode, staticFlag, firstToken, token)
}
// 1638: decl @lune.@base.@TransUnit.TransUnit.checkPublic
func (self *TransUnit_TransUnit) checkPublic(_env *LnsEnv, pos Types_Position,typeInfo *Ast_TypeInfo) {
    var checkedTypeSet *LnsSet
    checkedTypeSet = NewLnsSet([]LnsAny{})
    var checkPub func(_env *LnsEnv, workType *Ast_TypeInfo)
    checkPub = func(_env *LnsEnv, workType *Ast_TypeInfo) {
        if checkedTypeSet.Has(Ast_TypeInfo2Stem(workType)){
            return 
        }
        checkedTypeSet.Add(Ast_TypeInfo2Stem(workType))
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( workType.FP.Get_kind(_env) != Ast_TypeInfoKind__Array) &&
            _env.SetStackVal( workType.FP.Get_kind(_env) != Ast_TypeInfoKind__List) &&
            _env.SetStackVal( workType.FP.Get_kind(_env) != Ast_TypeInfoKind__Set) &&
            _env.SetStackVal( workType.FP.Get_kind(_env) != Ast_TypeInfoKind__Map) &&
            _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, workType.FP.Get_accessMode(_env)))) ).(bool)){
            self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("not public this type -- %s", []LnsAny{workType.FP.GetTxt(_env, nil, nil, nil)}))
        } else { 
            for _, _itemType := range( workType.FP.Get_itemTypeInfoList(_env).Items ) {
                itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                checkPub(_env, itemType)
            }
        }
    }
    checkPub(_env, typeInfo)
}
// 1660: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclMember
func (self *TransUnit_TransUnit) analyzeDeclMember(_env *LnsEnv, classTypeInfo *Ast_TypeInfo,accessMode LnsInt,staticFlag bool,firstToken *Types_Token) *Nodes_DeclMemberNode {
    __func__ := "@lune.@base.@TransUnit.TransUnit.analyzeDeclMember"
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var mutMode LnsInt
    mutMode = Ast_MutMode__IMut
    if _switch0 := nextToken.Txt; _switch0 == "mut" {
        mutMode = Ast_MutMode__Mut
        nextToken = self.FP.getToken(_env, nil)
    } else if _switch0 == "allmut" {
        mutMode = Ast_MutMode__AllMut
        nextToken = self.FP.getToken(_env, nil)
    }
    var varName *Types_Token
    varName = self.FP.checkSymbol(_env, nextToken, TransUnit_SymbolMode__MustNot_)
    var token *Types_Token
    token = self.FP.getToken(_env, nil)
    var refType *Nodes_RefTypeNode
    refType = self.FP.analyzeRefType(_env, accessMode, false, Ast_isPubToExternal(_env, classTypeInfo.FP.Get_accessMode(_env)))
    token = self.FP.getToken(_env, nil)
    if classTypeInfo.FP.IsInheritFrom(_env, self.ProcessInfo, Ast_builtinTypeAbsImmut, nil){
        if mutMode != Ast_MutMode__IMut{
            self.FP.AddErrMess(_env, varName.Pos, _env.GetVM().String_format("__absimmut can't have mutable member. -- %s", []LnsAny{varName.Txt}))
        }
        if Lns_op_not(Ast_TypeInfo_canBeAbsImmutMember(_env, refType.FP.Get_expType(_env))){
            self.FP.AddErrMess(_env, refType.FP.Get_pos(_env), _env.GetVM().String_format("__absimmut can't have member of mutable type. -- %s", []LnsAny{refType.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    if refType.FP.Get_expType(_env).FP.Get_asyncMode(_env) == Ast_Async__Transient{
        self.FP.AddErrMess(_env, refType.FP.Get_pos(_env), _env.GetVM().String_format("can't hold with the type of __trans. -- %s", []LnsAny{varName.Txt}))
    }
    var getterMode LnsInt
    getterMode = Ast_AccessMode__None
    var getterRetType *Ast_TypeInfo
    getterRetType = refType.FP.Get_expType(_env)
    var getterToken LnsAny
    getterToken = nil
    var getterMutable LnsInt
    getterMutable = Ast_MutMode__Mut
    var setterMode LnsInt
    setterMode = Ast_AccessMode__None
    var setterToken LnsAny
    setterToken = nil
    if token.Txt == "{"{
        var analyzeAccessorMode func(_env *LnsEnv)(LnsInt, *Ast_TypeInfo, *Types_Token, *Types_Token)
        analyzeAccessorMode = func(_env *LnsEnv)(LnsInt, *Ast_TypeInfo, *Types_Token, *Types_Token) {
            var retType *Ast_TypeInfo
            retType = Ast_headTypeInfo
            var mode LnsInt
            mode = Ast_AccessMode__None
            var accessorToken *Types_Token
            accessorToken = self.FP.getToken(_env, nil)
            var workToken *Types_Token
            workToken = accessorToken
            if _switch0 := workToken.Txt; _switch0 == "pub" || _switch0 == "pri" || _switch0 == "pro" || _switch0 == "local" {
                mode = Lns_unwrap( Ast_txt2AccessMode(_env, workToken.Txt)).(LnsInt)
                workToken = self.FP.getToken(_env, nil)
                if workToken.Txt == "&"{
                    getterMutable = Ast_MutMode__IMut
                    workToken = self.FP.getToken(_env, nil)
                }
                if workToken.Txt == ":"{
                    var typeNode *Nodes_RefTypeNode
                    typeNode = self.FP.analyzeRefType(_env, mode, false, Ast_isPubToExternal(_env, classTypeInfo.FP.Get_accessMode(_env)))
                    retType = typeNode.FP.Get_expType(_env)
                    workToken = self.FP.getToken(_env, nil)
                }
            } else if _switch0 == "non" {
                workToken = self.FP.getToken(_env, nil)
            } else {
                self.FP.AddErrMess(_env, workToken.Pos, _env.GetVM().String_format("access mode is invalid -- %s", []LnsAny{workToken.Txt}))
            }
            return mode, retType, accessorToken, workToken
        }
        {
            var workRetType *Ast_TypeInfo
            getterMode, workRetType, getterToken, nextToken = analyzeAccessorMode(_env)
            if workRetType != Ast_headTypeInfo{
                if Lns_op_not(Lns_car(workRetType.FP.CanEvalWith(_env, self.ProcessInfo, getterRetType, Ast_CanEvalType__SetOp, classTypeInfo.FP.CreateAlt2typeMap(_env, false))).(bool)){
                    self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("getter type mismatch -- %s <- %s", []LnsAny{workRetType.FP.GetTxt(_env, nil, nil, nil), getterRetType.FP.GetTxt(_env, nil, nil, nil)}))
                }
                getterRetType = workRetType
            }
        }
        if nextToken.Txt == ","{
            var dummyRetType *Ast_TypeInfo
            setterMode, dummyRetType, setterToken, nextToken = analyzeAccessorMode(_env)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( setterMode != Ast_AccessMode__None) &&
                _env.SetStackVal( mutMode == Ast_MutMode__IMut) ).(bool)){
                self.FP.AddErrMess(_env, varName.Pos, _env.GetVM().String_format("This member can't have setter, this member is immutable. -- %s", []LnsAny{varName.Txt}))
            }
            Log_log(_env, Log_Level__Debug, __func__, 1771, Log_CreateMessage(func(_env *LnsEnv) string {
                return _env.GetVM().String_format("%s", []LnsAny{dummyRetType})
            }))
            
        }
        self.FP.checkToken(_env, nextToken, "}")
        token = self.FP.getToken(_env, nil)
    }
    self.FP.checkToken(_env, token, ";")
    var typeInfo *Ast_TypeInfo
    typeInfo = refType.FP.Get_expType(_env)
    if self.Ctrl_info.LegacyMutableControl{
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Ast_TypeInfo_isMut(_env, typeInfo)) &&
            _env.SetStackVal( typeInfo.FP.Get_mutMode(_env) != mutMode) ).(bool)){
            typeInfo = self.FP.createModifier(_env, typeInfo, mutMode)
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Ast_TypeInfo_isMut(_env, getterRetType)) &&
            _env.SetStackVal( getterRetType.FP.Get_mutMode(_env) != mutMode) ).(bool)){
            getterRetType = self.FP.createModifier(_env, getterRetType, mutMode)
        }
    } else { 
        if Ast_TypeInfo_isMut(_env, getterRetType){
            if getterMutable == Ast_MutMode__AllMut{
                getterRetType = self.FP.createModifier(_env, getterRetType, Ast_MutMode__AllMut)
            } else if getterMutable == Ast_MutMode__IMut{
                getterRetType = self.FP.createModifier(_env, getterRetType, Ast_MutMode__IMut)
            }
        }
    }
    if Ast_isPubToExternal(_env, classTypeInfo.FP.Get_accessMode(_env)){
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) ||
            _env.SetStackVal( Ast_isPubToExternal(_env, setterMode)) ).(bool){
            self.FP.checkPublic(_env, refType.FP.Get_pos(_env), typeInfo)
        }
        if Ast_isPubToExternal(_env, getterMode){
            self.FP.checkPublic(_env, refType.FP.Get_pos(_env), getterRetType)
        }
    }
    var symbolInfo LnsAny
    var shadowing LnsAny
    symbolInfo,shadowing = self.FP.Get_scope(_env).FP.AddMember(_env, self.ProcessInfo, varName.Txt, varName.Pos, typeInfo, accessMode, staticFlag, mutMode)
    var workSym *Ast_SymbolInfo
    workSym = Lns_unwrap( (_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( symbolInfo) ||
        _env.SetStackVal( shadowing) ))).(*Ast_SymbolInfo)
    if shadowing != nil{
        shadowing_549 := shadowing.(*Ast_SymbolInfo)
        self.FP.errorShadowing(_env, varName.Pos, shadowing_549)
    }
    return Nodes_DeclMemberNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), varName, refType, workSym, classTypeInfo, staticFlag, accessMode, getterMutable != Ast_MutMode__IMut, getterMode, getterToken, getterRetType, setterMode, setterToken)
}
// 1826: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclMethod
func (self *TransUnit_TransUnit) analyzeDeclMethod(_env *LnsEnv, classTypeInfo *Ast_TypeInfo,declFuncMode LnsInt,abstractFlag bool,overrideFlag bool,accessMode LnsInt,staticFlag bool,firstToken *Types_Token,name *Types_Token) *Nodes_Node {
    var node *Nodes_Node
    node = self.FP.analyzeDeclFunc(_env, declFuncMode, false, abstractFlag, overrideFlag, accessMode, staticFlag, classTypeInfo, name, name)
    return node
}
// 1847: decl @lune.@base.@TransUnit.TransUnit.addDefaultConstructor
func (self *TransUnit_TransUnit) addDefaultConstructor(_env *LnsEnv, pos Types_Position,classTypeInfo *Ast_TypeInfo,typeDataAccessor Ast_TypeDataAccessor,classScope *Ast_Scope,memberNodeList *LnsList,methodNameSet *LnsSet,oldFlag bool) {
    if Lns_isCondTrue( classScope.FP.GetTypeInfoChild(_env, "__init")){
        self.FP.AddErrMess(_env, pos, "already declare __init().")
    }
    var argTypeList *LnsList
    argTypeList = NewLnsList([]LnsAny{})
    if classTypeInfo.FP.Get_baseTypeInfo(_env) != Ast_headTypeInfo{
        var superScope *Ast_Scope
        superScope = Lns_unwrap( classTypeInfo.FP.Get_baseTypeInfo(_env).FP.Get_scope(_env)).(*Ast_Scope)
        var superTypeInfo *Ast_TypeInfo
        superTypeInfo = Lns_unwrap( superScope.FP.GetTypeInfoChild(_env, "__init")).(*Ast_TypeInfo)
        for _, _argType := range( superTypeInfo.FP.Get_argTypeInfoList(_env).Items ) {
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if oldFlag{
                if Lns_op_not(argType.FP.Get_nilable(_env)){
                    self.FP.AddErrMess(_env, pos, "not found '__init' decl.")
                }
            } else { 
                argTypeList.Insert(Ast_TypeInfo2Stem(argType))
            }
        }
    }
    for _, _memberNode := range( memberNodeList.Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
            argTypeList.Insert(Ast_TypeInfo2Stem(memberNode.FP.Get_expType(_env)))
        }
    }
    if Ast_isPubToExternal(_env, classTypeInfo.FP.Get_accessMode(_env)){
        for _, _memberType := range( argTypeList.Items ) {
            memberType := _memberType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_op_not(Ast_isPubToExternal(_env, memberType.FP.Get_accessMode(_env))){
                self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("The type must be 'pub' becaue using in __init(). -- %s:%s", []LnsAny{memberType.FP.GetTxt(_env, nil, nil, nil), Ast_AccessMode_getTxt( memberType.FP.Get_accessMode(_env))}))
            }
        }
    }
    var ctorScope *Ast_Scope
    ctorScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var initTypeInfo *Ast_NormalTypeInfo
    initTypeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, ctorScope, Ast_TypeInfoKind__Method, classTypeInfo, typeDataAccessor, true, false, false, Ast_AccessMode__Pub, "__init", Ast_Async__Async, nil, argTypeList, NewLnsList([]LnsAny{}), Ast_MutMode__IMut)
    if oldFlag{
        ctorScope.FP.AddVar(_env, self.ProcessInfo, Ast_AccessMode__Pri, "", nil, Ast_headTypeInfo, Ast_MutMode__IMut, true)
    }
    self.FP.PopScope(_env)
    classScope.FP.AddMethod(_env, self.ProcessInfo, pos, &initTypeInfo.Ast_TypeInfo, Ast_AccessMode__Pub, false)
    methodNameSet.Add("__init")
    for _, _memberNode := range( memberNodeList.Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
            memberNode.FP.Get_symbolInfo(_env).FP.UpdateValue(_env, memberNode.FP.Get_symbolInfo(_env).FP.Get_posForLatestMod(_env))
        }
    }
}
// 1924: decl @lune.@base.@TransUnit.TransUnit.analyzeFuncBlock
func (self *TransUnit_TransUnit) analyzeFuncBlock(_env *LnsEnv, analyzingState LnsInt,firstToken *Types_Token,classTypeInfo LnsAny,funcTypeInfo *Ast_TypeInfo,funcName string,funcBodyScope *Ast_Scope,retTypeInfoList *LnsList) *Nodes_BlockNode {
    if Lns_op_not(funcTypeInfo.FP.Get_staticFlag(_env)){
        if classTypeInfo != nil{
            classTypeInfo_579 := classTypeInfo.(*Ast_TypeInfo)
            {
                _overrideType := _env.NilAccFin(_env.NilAccPush(classTypeInfo_579.FP.Get_scope(_env)) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetTypeInfoField(_env, funcName, false, funcBodyScope, self.ScopeAccess)})/* 1931:33 */)
                if !Lns_IsNil( _overrideType ) {
                    overrideType := _overrideType.(*Ast_TypeInfo)
                    if Lns_op_not(overrideType.FP.Get_abstractFlag(_env)){
                        funcBodyScope.FP.Add(_env, self.ProcessInfo, Ast_SymbolKind__Fun, false, false, "super", nil, overrideType, Ast_AccessMode__Local, false, Ast_MutMode__IMut, true, false)
                    }
                }
            }
        }
    }
    var body *Nodes_BlockNode
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(self.inTestBlock)) ||
        _env.SetStackVal( self.Ctrl_info.Testing) ).(bool){
        self.FP.pushAnalyzingState(_env, analyzingState)
        body = self.FP.analyzeBlock(_env, Nodes_BlockKind__Func, TransUnit_TentativeMode__Ignore, funcBodyScope, nil)
        self.FP.popAnalyzingState(_env)
        if retTypeInfoList.Len() != 0{
            var breakKind LnsInt
            breakKind = body.FP.GetBreakKind(_env, Nodes_CheckBreakMode__Return)
            if retTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeNeverRet{
                if _switch0 := breakKind; _switch0 == Nodes_BreakKind__Return || _switch0 == Nodes_BreakKind__NeverRet {
                } else {
                    self.FP.AddErrMess(_env, firstToken.Pos, "This funcion doesn't have return.")
                }
            } else { 
                if breakKind != Nodes_BreakKind__NeverRet{
                    self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("This funcion must be never return. -- %s", []LnsAny{Nodes_BreakKind_getTxt( breakKind)}))
                }
            }
        }
    } else { 
        body = self.FP.skipAndCreateDummyBlock(_env)
    }
    return body
}
// 1991: decl @lune.@base.@TransUnit.TransUnit.addAccessor
func (self *TransUnit_TransUnit) addAccessor(_env *LnsEnv, memberNode *Nodes_DeclMemberNode,methodNameSet *LnsSet,classScope *Ast_Scope,classTypeInfo *Ast_TypeInfo,typeDataAccessor Ast_TypeDataAccessor) {
    var memberType *Ast_TypeInfo
    memberType = memberNode.FP.Get_expType(_env)
    var memberName *Types_Token
    memberName = memberNode.FP.Get_name(_env)
    var getterName string
    getterName = "get_" + memberName.Txt
    var accessMode LnsInt
    accessMode = memberNode.FP.Get_getterMode(_env)
    var typeKind LnsInt
    if memberNode.FP.Get_staticFlag(_env){
        typeKind = Ast_TypeInfoKind__Func
    } else { 
        typeKind = Ast_TypeInfoKind__Method
    }
    var asyncMode LnsInt
    if memberType.FP.Get_mutMode(_env) == Ast_MutMode__AllMut{
        asyncMode = Ast_Async__Noasync
    } else { 
        asyncMode = Ast_Async__Async
    }
    if accessMode != Ast_AccessMode__None{
        if Lns_isCondTrue( classScope.FP.GetTypeInfoChild(_env, getterName)){
            self.FP.AddErrMess(_env, memberName.Pos, _env.GetVM().String_format("exist -- %s.%s", []LnsAny{classTypeInfo.FP.Get_rawTxt(_env), getterName}))
        } else { 
            var mutable bool
            mutable = memberNode.FP.Get_getterMutable(_env)
            var getterMemberType *Ast_TypeInfo
            getterMemberType = memberNode.FP.Get_getterRetType(_env)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Ast_TypeInfo_isMut(_env, getterMemberType)) &&
                _env.SetStackVal( Lns_op_not(mutable)) ).(bool)){
                getterMemberType = self.FP.createModifier(_env, getterMemberType, Ast_MutMode__IMut)
            }
            var retTypeInfo *Ast_NormalTypeInfo
            retTypeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil), typeKind, classTypeInfo, typeDataAccessor, false, false, memberNode.FP.Get_staticFlag(_env), accessMode, getterName, asyncMode, nil, NewLnsList([]LnsAny{}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(getterMemberType)}), Ast_MutMode__IMut)
            self.FP.PopScope(_env)
            classScope.FP.AddMethod(_env, self.ProcessInfo, memberName.Pos, &retTypeInfo.Ast_TypeInfo, accessMode, memberNode.FP.Get_staticFlag(_env))
            methodNameSet.Add(getterName)
        }
    }
    var setterName string
    setterName = "set_" + memberName.Txt
    accessMode = memberNode.FP.Get_setterMode(_env)
    if memberNode.FP.Get_setterMode(_env) != Ast_AccessMode__None{
        if Lns_isCondTrue( classScope.FP.GetTypeInfoChild(_env, setterName)){
            self.FP.AddErrMess(_env, memberName.Pos, _env.GetVM().String_format("exist -- %s.%s", []LnsAny{classTypeInfo.FP.Get_rawTxt(_env), setterName}))
        } else { 
            var mutMode LnsInt
            if memberNode.FP.Get_symbolInfo(_env).FP.Get_mutMode(_env) != Ast_MutMode__AllMut{
                mutMode = Ast_MutMode__Mut
            } else { 
                mutMode = Ast_MutMode__IMut
            }
            classScope.FP.AddMethod(_env, self.ProcessInfo, memberName.Pos, &self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil), typeKind, classTypeInfo, typeDataAccessor, false, false, memberNode.FP.Get_staticFlag(_env), accessMode, setterName, asyncMode, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(memberType)}), nil, mutMode).Ast_TypeInfo, accessMode, memberNode.FP.Get_staticFlag(_env))
            self.FP.PopScope(_env)
            methodNameSet.Add(setterName)
        }
    }
}
// 2086: decl @lune.@base.@TransUnit.TransUnit.analyzeClassBody
func (self *TransUnit_TransUnit) analyzeClassBody(_env *LnsEnv, hasProto bool,classAccessMode LnsInt,firstToken *Types_Token,mode LnsInt,gluePrefix LnsAny,classTypeInfo *Ast_TypeInfo,typeDataAccessor Ast_TypeDataAccessor,name *Types_Token,moduleLang LnsAny,moduleName LnsAny,lazyLoad LnsInt,nextToken *Types_Token,inheritInfo *Nodes_ClassInheritInfo)(*Nodes_DeclClassNode, *Types_Token, *LnsSet) {
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
    initBlockInfo = NewNodes_ClassInitBlockInfo(_env, nil)
    var advertiseList *LnsList
    advertiseList = NewLnsList([]LnsAny{})
    var trustList *LnsList
    trustList = NewLnsList([]LnsAny{})
    var uninitMemberList *LnsList
    uninitMemberList = NewLnsList([]LnsAny{})
    var node *Nodes_DeclClassNode
    node = Nodes_DeclClassNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo)}), classAccessMode, name, inheritInfo, hasProto, gluePrefix, moduleName, moduleLang, lazyLoad, false, allStmtList, declStmtList, fieldList, memberList, self.FP.Get_scope(_env), initBlockInfo, advertiseList, trustList, uninitMemberList, NewLnsSet([]LnsAny{}))
    self.typeInfo2ClassNode.Set(classTypeInfo,node)
    var alreadyCtorFlag bool
    alreadyCtorFlag = false
    var hasInitBlock bool
    hasInitBlock = false
    var hasStaticMember bool
    hasStaticMember = false
    var classScope *Ast_Scope
    classScope = self.FP.Get_scope(_env)
    var processLet func(_env *LnsEnv, token *Types_Token,staticFlag bool,accessMode LnsInt,alreadyFlag bool)
    processLet = func(_env *LnsEnv, token *Types_Token,staticFlag bool,accessMode LnsInt,alreadyFlag bool) {
        if staticFlag{
            hasStaticMember = true
        }
        if mode == TransUnitIF_DeclClassMode__Interface{
            self.FP.AddErrMess(_env, token.Pos, "interface can not have member")
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(staticFlag)) &&
            _env.SetStackVal( alreadyFlag) ).(bool)){
            self.FP.AddErrMess(_env, token.Pos, "member can't declare after '__init' method.")
        } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( staticFlag) &&
            _env.SetStackVal( hasInitBlock) ).(bool)){
            self.FP.AddErrMess(_env, token.Pos, "static member can't declare after '__init' block.")
        }
        var memberNode *Nodes_DeclMemberNode
        memberNode = self.FP.analyzeDeclMember(_env, classTypeInfo, accessMode, staticFlag, token)
        allStmtList.Insert(Nodes_DeclMemberNode2Stem(memberNode))
        fieldList.Insert(Nodes_DeclMemberNode2Stem(memberNode))
        memberList.Insert(Nodes_DeclMemberNode2Stem(memberNode))
        memberName2Node.Set(memberNode.FP.Get_name(_env).Txt,memberNode)
        self.FP.addAccessor(_env, memberNode, methodNameSet, classScope, classTypeInfo, typeDataAccessor)
    }
    var checkInitializeMember func(_env *LnsEnv, staticFlag bool,pos LnsAny)
    checkInitializeMember = func(_env *LnsEnv, staticFlag bool,pos LnsAny) {
        for _memberName, _memberNode := range( memberName2Node.Items ) {
            memberName := _memberName.(string)
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if memberNode.FP.Get_staticFlag(_env) == staticFlag{
                var symbolInfo *Ast_SymbolInfo
                symbolInfo = Lns_unwrap( self.FP.Get_scope(_env).FP.GetSymbolInfoChild(_env, memberName)).(*Ast_SymbolInfo)
                var typeInfo *Ast_TypeInfo
                typeInfo = symbolInfo.FP.Get_typeInfo(_env)
                if Lns_op_not(symbolInfo.FP.Get_hasValueFlag(_env)){
                    var msg string
                    if staticFlag{
                        msg = _env.GetVM().String_format("Set member -- %s", []LnsAny{memberName})
                    } else { 
                        msg = _env.GetVM().String_format("Set member -- %s.%s", []LnsAny{name.Txt, memberName})
                    }
                    if Lns_op_not(typeInfo.FP.Get_nilable(_env)){
                        self.FP.AddErrMess(_env, Lns_unwrapDefault( pos, memberNode.FP.Get_pos(_env)).(Types_Position), msg)
                    } else { 
                        uninitMemberList.Insert(Ast_SymbolInfo2Stem(symbolInfo))
                        self.FP.addWarnMess(_env, Lns_unwrapDefault( pos, memberNode.FP.Get_pos(_env)).(Types_Position), msg)
                    }
                }
            }
        }
    }
    var processFn func(_env *LnsEnv, token *Types_Token,staticFlag bool,accessMode LnsInt,abstractFlag bool,overrideFlag bool)
    processFn = func(_env *LnsEnv, token *Types_Token,staticFlag bool,accessMode LnsInt,abstractFlag bool,overrideFlag bool) {
        var nameToken *Types_Token
        nameToken = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
        var declFuncMode LnsInt
        declFuncMode = TransUnit_DeclFuncMode__Class
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( mode == TransUnitIF_DeclClassMode__Module) ||
            _env.SetStackVal( mode == TransUnitIF_DeclClassMode__LazyModule) ).(bool){
            if Lns_isCondTrue( gluePrefix){
                declFuncMode = TransUnit_DeclFuncMode__Glue
            } else { 
                declFuncMode = TransUnit_DeclFuncMode__Module
            }
        }
        if nameToken.Txt == "__init"{
            for _, _symbolInfo := range( self.FP.Get_scope(_env).FP.Get_symbol2SymbolInfoMap(_env).Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if Lns_op_not(symbolInfo.FP.Get_staticFlag(_env)){
                    symbolInfo.FP.ClearValue(_env)
                }
            }
        }
        var methodNode *Nodes_Node
        methodNode = self.FP.analyzeDeclMethod(_env, classTypeInfo, declFuncMode, abstractFlag, overrideFlag, accessMode, staticFlag, token, nameToken)
        allStmtList.Insert(Nodes_Node2Stem(methodNode))
        fieldList.Insert(Nodes_Node2Stem(methodNode))
        methodNameSet.Add(nameToken.Txt)
        if nameToken.Txt == "__init"{
            alreadyCtorFlag = true
            checkInitializeMember(_env, false, methodNode.FP.Get_pos(_env))
        }
    }
    var processInitBlock func(_env *LnsEnv, token *Types_Token)
    processInitBlock = func(_env *LnsEnv, token *Types_Token) {
        if _env.NilAccFin(_env.NilAccPush(classTypeInfo.FP.Get_scope(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_parent(_env)})) != self.ModuleScope{
            self.FP.AddErrMess(_env, token.Pos, "The '__init' block only support at the top level classes.")
        }
        if mode != TransUnitIF_DeclClassMode__Class{
            self.FP.Error(_env, _env.GetVM().String_format("%s can not have __init block.", []LnsAny{mode}))
        }
        hasInitBlock = true
        for _, _symbolInfo := range( self.FP.Get_scope(_env).FP.Get_symbol2SymbolInfoMap(_env).Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if symbolInfo.FP.Get_staticFlag(_env){
                symbolInfo.FP.ClearValue(_env)
            }
        }
        var parentScope *Ast_Scope
        parentScope = self.FP.Get_scope(_env)
        var initBlockScope *Ast_Scope
        initBlockScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
        self.FP.prepareTentativeSymbol(_env, initBlockScope, false, nil)
        var ininame string
        ininame = "___init"
        var funcTypeInfo *Ast_NormalTypeInfo
        funcTypeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, initBlockScope, Ast_TypeInfoKind__Func, classTypeInfo, typeDataAccessor, false, false, true, Ast_AccessMode__Pri, ininame, Ast_Async__Noasync, nil, nil, nil, Ast_MutMode__IMut)
        var funcSym LnsAny
        var shadowing LnsAny
        funcSym,shadowing = parentScope.FP.AddFunc(_env, self.ProcessInfo, token.Pos, &funcTypeInfo.Ast_TypeInfo, Ast_AccessMode__Pri, true, true)
        var nsInfo *TransUnitIF_NSInfo
        nsInfo = self.FP.NewNSInfo(_env, &funcTypeInfo.Ast_TypeInfo, token.Pos)
        var block *Nodes_BlockNode
        block = self.FP.analyzeFuncBlock(_env, TransUnit_AnalyzingState__InitBlock, token, classTypeInfo, &funcTypeInfo.Ast_TypeInfo, ininame, initBlockScope, funcTypeInfo.FP.Get_retTypeInfoList(_env))
        var info *Nodes_DeclFuncInfo
        info = NewNodes_DeclFuncInfo(_env, Nodes_FuncKind__InitBlock, classTypeInfo, node, true, token, Lns_unwrap( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( funcSym) ||
            _env.SetStackVal( shadowing) ))).(*Ast_SymbolInfo), NewLnsList([]LnsAny{}), true, Ast_AccessMode__Pri, nil, block, NewLnsList([]LnsAny{}), NewLnsList([]LnsAny{}), false, false, nsInfo.FP.Get_stmtNum(_env))
        var initBlockNode *Nodes_DeclMethodNode
        initBlockNode = Nodes_DeclMethodNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_NormalTypeInfo2Stem(funcTypeInfo)}), info)
        initBlockInfo.FP.Set_func(_env, initBlockNode)
        allStmtList.Insert(Nodes_DeclMethodNode2Stem(initBlockNode))
        self.FP.PopScope(_env)
        self.FP.finishTentativeSymbol(_env, false)
    }
    var processAdvertise func(_env *LnsEnv)
    processAdvertise = func(_env *LnsEnv) {
        var memberToken *Types_Token
        memberToken = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
        nextToken = self.FP.getToken(_env, nil)
        var prefix string
        prefix = ""
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( nextToken.Txt != ";") &&
            _env.SetStackVal( nextToken.Txt != "{") ).(bool)){
            prefix = nextToken.Txt
            nextToken = self.FP.getToken(_env, nil)
        }
        self.FP.checkToken(_env, nextToken, ";")
        var memberNode *Nodes_DeclMemberNode
        
        {
            _memberNode := memberName2Node.Get(memberToken.Txt)
            if _memberNode == nil{
                self.FP.Error(_env, _env.GetVM().String_format("not found member -- %s", []LnsAny{memberToken.Txt}))
            } else {
                memberNode = _memberNode.(*Nodes_DeclMemberNode)
            }
        }
        var advInfo *Nodes_AdvertiseInfo
        advInfo = NewNodes_AdvertiseInfo(_env, memberNode, prefix, memberToken.Pos)
        advertiseList.Insert(Nodes_AdvertiseInfo2Stem(advInfo))
        allStmtList.Insert(Nodes_DeclAdvertiseNode2Stem(Nodes_DeclAdvertiseNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), advInfo)))
        self.advertisedTypeSet.Add(Ast_TypeInfo2Stem(memberNode.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env).FP.Get_genSrcTypeInfo(_env)))
    }
    var processEnum func(_env *LnsEnv, token *Types_Token,accessMode LnsInt)
    processEnum = func(_env *LnsEnv, token *Types_Token,accessMode LnsInt) {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( accessMode != Ast_AccessMode__Pri) &&
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( classAccessMode == Ast_AccessMode__Pri) ||
                _env.SetStackVal( classAccessMode == Ast_AccessMode__Local) ).(bool))) ).(bool)){
            self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("unmatch access mode, class('%s') and enum('%s')", []LnsAny{Ast_AccessMode_getTxt( classAccessMode), Ast_AccessMode_getTxt( accessMode)}))
        }
        var enumNode *Nodes_DeclEnumNode
        enumNode = self.FP.analyzeDeclEnum(_env, accessMode, token)
        allStmtList.Insert(Nodes_DeclEnumNode2Stem(enumNode))
        declStmtList.Insert(Nodes_DeclEnumNode2Stem(enumNode))
    }
    var processLuneControl func(_env *LnsEnv)
    processLuneControl = func(_env *LnsEnv) {
        nextToken = self.FP.getToken(_env, nil)
        var pragma LnsAny
        if _switch0 := nextToken.Txt; _switch0 == "default__init" {
            pragma = LuneControl_Pragma__default__init_Obj
            alreadyCtorFlag = true
            self.FP.addDefaultConstructor(_env, nextToken.Pos, classTypeInfo, typeDataAccessor, self.FP.Get_scope(_env), memberList, methodNameSet, false)
        } else if _switch0 == "default__init_old" {
            pragma = LuneControl_Pragma__default__init_old_Obj
            alreadyCtorFlag = true
            self.FP.addDefaultConstructor(_env, nextToken.Pos, classTypeInfo, typeDataAccessor, self.FP.Get_scope(_env), memberList, methodNameSet, true)
            node.FP.SetHasOldCtor(_env)
        } else if _switch0 == "default_async_this_class" {
            pragma = LuneControl_Pragma__default_async_this_class_Obj
            self.class2defaultAsyncMode.Set(self.FP.GetCurrentNamespaceTypeInfo(_env),TransUnit_DefaultAsyncMode__AsyncAll)
        } else if _switch0 == "default_noasync_this_class" {
            pragma = LuneControl_Pragma__default_noasync_this_class_Obj
            self.class2defaultAsyncMode.Set(self.FP.GetCurrentNamespaceTypeInfo(_env),TransUnit_DefaultAsyncMode__NoAsync)
        } else {
            self.FP.Error(_env, _env.GetVM().String_format("unknown option -- %s", []LnsAny{nextToken.Txt}))
        }
        self.FP.checkNextToken(_env, ";")
        var ctrlNode *Nodes_LuneControlNode
        ctrlNode = Nodes_LuneControlNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), pragma)
        self.HelperInfo.PragmaSet.Add(pragma)
        allStmtList.Insert(Nodes_LuneControlNode2Stem(ctrlNode))
    }
    var processClassFields func(_env *LnsEnv, inMacro bool)
    processClassFields = func(_env *LnsEnv, inMacro bool) {
        for  {
            var token *Types_Token
            token = self.FP.getToken(_env, inMacro)
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( token.Kind == Types_TokenKind__Eof) ||
                _env.SetStackVal( token.Txt == "}") ).(bool){
                break
            }
            var accessMode LnsInt
            
            {
                _accessMode := Ast_txt2AccessMode(_env, token.Txt)
                if _accessMode == nil{
                    accessMode = Ast_AccessMode__Pri
                } else {
                    accessMode = _accessMode.(LnsInt)
                    token = self.FP.getToken(_env, nil)
                }
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( mode == TransUnitIF_DeclClassMode__Interface) &&
                _env.SetStackVal( accessMode != Ast_AccessMode__Pub) ).(bool)){
                self.FP.AddErrMess(_env, token.Pos, "interface's fields must be 'pub'.")
            }
            var staticFlag bool
            staticFlag = false
            if token.Txt == "static"{
                staticFlag = true
                token = self.FP.getToken(_env, nil)
            }
            var overrideFlag bool
            overrideFlag = false
            if token.Txt == "override"{
                overrideFlag = true
                token = self.FP.getToken(_env, nil)
            }
            var abstractFlag bool
            abstractFlag = false
            if token.Txt == "abstract"{
                abstractFlag = true
                token = self.FP.getToken(_env, nil)
            } else if mode == TransUnitIF_DeclClassMode__Interface{
                abstractFlag = true
            }
            if token.Txt == "let"{
                processLet(_env, token, staticFlag, accessMode, alreadyCtorFlag)
            } else if token.Txt == "fn"{
                processFn(_env, token, staticFlag, accessMode, abstractFlag, overrideFlag)
            } else if token.Txt == "__init"{
                processInitBlock(_env, token)
            } else if token.Txt == "advertise"{
                processAdvertise(_env)
            } else if token.Txt == ";"{
            } else if token.Txt == "enum"{
                processEnum(_env, token, accessMode)
            } else if token.Txt == "_lune_control"{
                processLuneControl(_env)
            } else { 
                {
                    _symbolInfo := self.FP.Get_scope(_env).FP.GetSymbolInfo(_env, token.Txt, self.FP.Get_scope(_env), false, self.ScopeAccess)
                    if !Lns_IsNil( _symbolInfo ) {
                        symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                        if symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mac{
                            self.FP.checkNextToken(_env, "(")
                            var argList LnsAny
                            _,argList = self.FP.prepareExpCall(_env, token.Pos, symbolInfo.FP.Get_typeInfo(_env), NewLnsList([]LnsAny{}), Ast_headTypeInfo)
                            self.FP.evalMacroOp(_env, token, symbolInfo.FP.Get_typeInfo(_env), argList, Macro_EvalMacroCallback(func(_env *LnsEnv) {
                                processClassFields(_env, true)
                            }))
                            self.FP.checkNextToken(_env, ";")
                        } else { 
                            self.FP.Error(_env, "illegal field")
                        }
                    } else {
                        self.FP.Error(_env, "illegal field")
                    }
                }
            }
        }
    }
    processClassFields(_env, false)
    if _switch0 := mode; _switch0 == TransUnitIF_DeclClassMode__Module || _switch0 == TransUnitIF_DeclClassMode__LazyModule {
    } else {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( hasStaticMember) &&
            _env.SetStackVal( Lns_op_not(hasInitBlock)) ).(bool)){
            self.FP.AddErrMess(_env, node.FP.Get_pos(_env), _env.GetVM().String_format("This class (%s) need __init block for initialize static members.", []LnsAny{name.Txt}))
        }
        checkInitializeMember(_env, true, nil)
    }
    return node, nextToken, methodNameSet
}
// 2454: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclClass
func (self *TransUnit_TransUnit) analyzeDeclClass(_env *LnsEnv, classAbstructFlag bool,classAccessMode LnsInt,firstToken *Types_Token,mode LnsInt) *Nodes_DeclClassNode {
    if mode == TransUnitIF_DeclClassMode__Module{
        if self.FP.getToken(_env, nil).Txt == "."{
            if _switch0 := self.FP.getToken(_env, nil).Txt; _switch0 == "l" {
                mode = TransUnitIF_DeclClassMode__LazyModule
            } else if _switch0 == "d" {
                mode = TransUnitIF_DeclClassMode__Module
            }
        } else { 
            self.FP.Pushback(_env)
            if self.Ctrl_info.DefaultLazy{
                mode = TransUnitIF_DeclClassMode__LazyModule
            }
        }
    }
    if mode == TransUnitIF_DeclClassMode__LazyModule{
        self.HelperInfo.UseLazyRequire = true
    }
    if _switch3 := mode; _switch3 == TransUnitIF_DeclClassMode__Module || _switch3 == TransUnitIF_DeclClassMode__LazyModule {
    } else {
        if _switch2 := self.FP.GetCurrentNamespaceTypeInfo(_env).FP.Get_kind(_env); _switch2 == Ast_TypeInfoKind__IF || _switch2 == Ast_TypeInfoKind__Class || _switch2 == Ast_TypeInfoKind__Module {
        } else if _switch2 == Ast_TypeInfoKind__Func || _switch2 == Ast_TypeInfoKind__Method {
            if _switch1 := classAccessMode; _switch1 == Ast_AccessMode__Pub || _switch1 == Ast_AccessMode__Global {
                self.FP.AddErrMess(_env, firstToken.Pos, "Class can't declare on here.")
            }
        } else {
            self.FP.AddErrMess(_env, firstToken.Pos, "Class can't declare on here.")
        }
    }
    var name *Types_Token
    name = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
    var altTypeList *LnsList
    altTypeList = NewLnsList([]LnsAny{})
    {
        var nextToken *Types_Token
        nextToken = self.FP.getToken(_env, nil)
        if nextToken.Txt == "<"{
            nextToken, altTypeList = self.FP.analyzeDeclAlternateType(_env, true, nextToken, classAccessMode)
        }
        self.FP.PushbackToken(_env, nextToken)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( altTypeList.Len() > 0) &&
            _env.SetStackVal( mode != TransUnitIF_DeclClassMode__Class) ).(bool)){
            self.FP.AddErrMess(_env, name.Pos, _env.GetVM().String_format("Only class can use the generics. -- %s ", []LnsAny{name.Txt}))
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
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode == TransUnitIF_DeclClassMode__Module) ||
        _env.SetStackVal( mode == TransUnitIF_DeclClassMode__LazyModule) ).(bool){
        self.FP.checkNextToken(_env, "require")
        moduleName = self.FP.getToken(_env, nil)
        var nextToken *Types_Token
        nextToken = self.FP.getToken(_env, nil)
        if nextToken.Txt == "of"{
            var langToken *Types_Token
            langToken = self.FP.getToken(_env, nil)
            if langToken.Kind != Types_TokenKind__Str{
                self.FP.Error(_env, _env.GetVM().String_format("it's not a string -- %s", []LnsAny{langToken.Txt}))
            }
            var langIdToken string
            langIdToken = langToken.FP.GetExcludedDelimitTxt(_env)
            if langIdToken != ""{
                for _, _langId := range( Types_Lang_get__allList(_env).Items ) {
                    langId := _langId.(LnsInt)
                    {
                        __exp := Types_Lang__from(_env, LnsInt(langId))
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(LnsInt)
                            var ldName string
                            ldName = TransUnit_convExp3_408(Lns_2DDD(_env.GetVM().String_gsub(Types_Lang_getTxt( _exp),".*%.", "")))
                            if ldName == langIdToken{
                                moduleLang = _exp
                                break
                            }
                        }
                    }
                }
                if moduleLang == nil{
                    self.FP.ErrorAt(_env, langToken.Pos, _env.GetVM().String_format("invalid lang -- %s", []LnsAny{langToken.Txt}))
                }
            } else { 
                moduleLang = Types_Lang__Same
            }
            nextToken = self.FP.getToken(_env, nil)
        }
        if nextToken.Txt == "glue"{
            gluePrefix = self.FP.getToken(_env, nil).FP.GetExcludedDelimitTxt(_env)
        } else { 
            self.FP.Pushback(_env)
        }
    }
    var existSymbolInfo LnsAny
    existSymbolInfo = self.FP.Get_scope(_env).FP.GetSymbolTypeInfo(_env, name.Txt, self.FP.Get_scope(_env), self.FP.Get_scope(_env), self.ScopeAccess)
    var nextToken *Types_Token
    var nsInfo *TransUnitIF_NSInfo
    var inheritInfo *Nodes_ClassInheritInfo
    nextToken,nsInfo,inheritInfo = self.FP.analyzePushClass(_env, mode, classAbstructFlag, firstToken, name, true, moduleName, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( moduleLang) ||
        _env.SetStackVal( Types_Lang__Same) ).(LnsInt), classAccessMode, altTypeList)
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = nsInfo.FP.Get_typeInfo(_env)
    var typeDataAccessor Ast_TypeDataAccessor
    typeDataAccessor = nsInfo.FP.Get_typeDataAccessor(_env)
    var hasProto bool
    if nsInfo.FP.Get_nobody(_env){
        nsInfo.FP.Set_nobody(_env, false)
        hasProto = true
    } else { 
        hasProto = false
        if Lns_isCondTrue( existSymbolInfo){
            self.FP.AddErrMess(_env, name.Pos, _env.GetVM().String_format("already declare symbol -- %s", []LnsAny{name.Txt}))
        }
    }
    var classScope *Ast_Scope
    classScope = self.FP.Get_scope(_env)
    self.FP.checkToken(_env, nextToken, "{")
    var mapType *Ast_TypeInfo
    mapType = self.ProcessInfo.FP.CreateMap(_env, Ast_AccessMode__Pub, classTypeInfo, Ast_builtinTypeString, self.FP.createModifier(_env, Ast_builtinTypeStem, Ast_MutMode__IMut), Ast_MutMode__IMut)
    if classTypeInfo.FP.IsInheritFrom(_env, self.ProcessInfo, Ast_builtinTypeMapping, nil){
        self.HelperInfo.HasMappingClassDef = true
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( classTypeInfo.FP.Get_baseTypeInfo(_env) != Ast_headTypeInfo) &&
            _env.SetStackVal( Lns_op_not(classTypeInfo.FP.Get_baseTypeInfo(_env).FP.IsInheritFrom(_env, self.ProcessInfo, Ast_builtinTypeMapping, nil))) ).(bool)){
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("must extend Mapping at %s", []LnsAny{classTypeInfo.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
        var toMapFuncTypeInfo *Ast_NormalTypeInfo
        toMapFuncTypeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, nil, Ast_TypeInfoKind__Method, classTypeInfo, typeDataAccessor, true, false, false, Ast_AccessMode__Pub, "_toMap", Ast_Async__Async, nil, NewLnsList([]LnsAny{}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(mapType)}), Ast_MutMode__IMut)
        classScope.FP.AddMethod(_env, self.ProcessInfo, nil, &toMapFuncTypeInfo.Ast_TypeInfo, Ast_AccessMode__Pub, false)
    }
    var lazyLoad LnsInt
    if _switch4 := mode; _switch4 == TransUnitIF_DeclClassMode__LazyModule {
        lazyLoad = Nodes_LazyLoad__On
    } else if _switch4 == TransUnitIF_DeclClassMode__Module || _switch4 == TransUnitIF_DeclClassMode__Class || _switch4 == TransUnitIF_DeclClassMode__Interface {
        lazyLoad = Nodes_LazyLoad__Off
    }
    var node *Nodes_DeclClassNode
    var methodNameSet *LnsSet
    node,_,methodNameSet = self.FP.analyzeClassBody(_env, hasProto, classAccessMode, firstToken, mode, gluePrefix, classTypeInfo, typeDataAccessor, name, moduleLang, moduleName, lazyLoad, nextToken, inheritInfo)
    var ctorAccessMode LnsInt
    ctorAccessMode = Ast_AccessMode__Pub
    {
        _ctorTypeInfo := classScope.FP.GetTypeInfoChild(_env, "__init")
        if !Lns_IsNil( _ctorTypeInfo ) {
            ctorTypeInfo := _ctorTypeInfo.(*Ast_TypeInfo)
            ctorAccessMode = ctorTypeInfo.FP.Get_accessMode(_env)
        } else {
            self.FP.addDefaultConstructor(_env, firstToken.Pos, classTypeInfo, typeDataAccessor, classScope, node.FP.Get_memberList(_env), methodNameSet, false)
        }
    }
    for _, _advertiseInfo := range( node.FP.Get_advertiseList(_env).Items ) {
        advertiseInfo := _advertiseInfo.(Nodes_AdvertiseInfoDownCast).ToNodes_AdvertiseInfo()
        var memberType *Ast_TypeInfo
        memberType = advertiseInfo.FP.Get_member(_env).FP.Get_expType(_env)
        if _switch5 := memberType.FP.Get_kind(_env); _switch5 == Ast_TypeInfoKind__Class || _switch5 == Ast_TypeInfoKind__IF {
            for _, _mtdName := range( Ast_getAllMethodName(_env, memberType, Ast_MethodKind__Object).FP.Get_list(_env).Items ) {
                mtdName := _mtdName.(string)
                var scope *Ast_Scope
                scope = Lns_unwrap( memberType.FP.Get_scope(_env)).(*Ast_Scope)
                var child *Ast_TypeInfo
                child = Lns_unwrap( scope.FP.GetTypeInfoField(_env, mtdName, true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
                if child.FP.Get_accessMode(_env) != Ast_AccessMode__Pri{
                    var childName string
                    childName = advertiseInfo.FP.Get_prefix(_env) + child.FP.GetTxt(_env, nil, nil, nil)
                    if Lns_op_not(methodNameSet.Has(childName)){
                        var impMtdType *Ast_TypeInfo
                        impMtdType = self.ProcessInfo.FP.CreateAdvertiseMethodFrom(_env, classTypeInfo, typeDataAccessor, child)
                        classScope.FP.AddMethod(_env, self.ProcessInfo, advertiseInfo.FP.Get_pos(_env), impMtdType, child.FP.Get_accessMode(_env), child.FP.Get_staticFlag(_env))
                    }
                }
            }
        } else {
            self.FP.Error(_env, _env.GetVM().String_format("advertise member type is illegal -- %s", []LnsAny{advertiseInfo.FP.Get_member(_env).FP.Get_name(_env)}))
        }
    }
    if classTypeInfo.FP.IsInheritFrom(_env, self.ProcessInfo, Ast_builtinTypeMapping, nil){
        var checkedTypeMap *LnsMap
        checkedTypeMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            var memberType *Ast_TypeInfo
            memberType = memberNode.FP.Get_expType(_env)
            if Lns_op_not(Ast_NormalTypeInfo_isAvailableMapping(_env, self.ProcessInfo, memberType, checkedTypeMap)){
                self.FP.AddErrMess(_env, memberNode.FP.Get_pos(_env), _env.GetVM().String_format("member type is not Mapping -- %s", []LnsAny{memberType.FP.GetTxt(_env, nil, nil, nil)}))
            } else if memberType.FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
                self.FP.AddErrMess(_env, memberNode.FP.Get_pos(_env), _env.GetVM().String_format("Mapping class has not the interface type member. -- %s", []LnsAny{memberNode.FP.Get_name(_env).Txt}))
            } else if memberType.FP.Get_abstractFlag(_env){
                self.FP.AddErrMess(_env, memberNode.FP.Get_pos(_env), _env.GetVM().String_format("Mapping class has not the abstract class member. -- %s", []LnsAny{memberNode.FP.Get_name(_env).Txt}))
            }
        }
        var fromMapFuncTypeInfo *Ast_NormalTypeInfo
        fromMapFuncTypeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, nil, Ast_TypeInfoKind__Func, classTypeInfo, typeDataAccessor, true, false, true, Ast_AccessMode__Pub, "_fromMap", Ast_Async__Async, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(mapType.FP.Get_nilableTypeInfo(_env))}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo.FP.Get_nilableTypeInfo(_env)), Ast_TypeInfo2Stem(Ast_builtinTypeString.FP.Get_nilableTypeInfo(_env))}), Ast_MutMode__IMut)
        classScope.FP.AddMethod(_env, self.ProcessInfo, nil, &fromMapFuncTypeInfo.Ast_TypeInfo, ctorAccessMode, true)
        var fromStemFuncTypeInfo *Ast_NormalTypeInfo
        fromStemFuncTypeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, nil, Ast_TypeInfoKind__Func, classTypeInfo, typeDataAccessor, true, false, true, Ast_AccessMode__Pub, "_fromStem", Ast_Async__Async, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeStem_)}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo.FP.Get_nilableTypeInfo(_env)), Ast_TypeInfo2Stem(Ast_builtinTypeString.FP.Get_nilableTypeInfo(_env))}), Ast_MutMode__IMut)
        classScope.FP.AddMethod(_env, self.ProcessInfo, nil, &fromStemFuncTypeInfo.Ast_TypeInfo, ctorAccessMode, true)
    }
    if classTypeInfo.FP.IsInheritFrom(_env, self.ProcessInfo, Ast_builtinTypeAsyncItem, nil){
        var pipeType *Ast_GenericTypeInfo
        pipeType = TransUnit_convExp3_1381(Lns_2DDD(self.ProcessInfo.FP.CreateGeneric(_env, self.builtinFunc.G__pipe_, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo)}), self.moduleType)))
        var createPipeFuncTypeInfo *Ast_NormalTypeInfo
        createPipeFuncTypeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, nil, Ast_TypeInfoKind__Func, classTypeInfo, typeDataAccessor, true, false, true, Ast_AccessMode__Pub, "_createPipe", Ast_Async__Async, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(pipeType.FP.Get_nilableTypeInfo(_env))}), Ast_MutMode__IMut)
        classScope.FP.AddMethod(_env, self.ProcessInfo, nil, &createPipeFuncTypeInfo.Ast_TypeInfo, Ast_AccessMode__Pub, true)
    }
    self.FP.PopClass(_env)
    return node
}
// 2748: decl @lune.@base.@TransUnit.TransUnit.addMethod
func (self *TransUnit_TransUnit) addMethod(_env *LnsEnv, classTypeInfo *Ast_TypeInfo,methodNode *Nodes_Node,name string) {
    var classNodeInfo *Nodes_DeclClassNode
    classNodeInfo = Lns_unwrap( self.typeInfo2ClassNode.Get(classTypeInfo)).(*Nodes_DeclClassNode)
    classNodeInfo.FP.Get_outerMethodSet(_env).Add(name)
    classNodeInfo.FP.Get_fieldList(_env).Insert(Nodes_Node2Stem(methodNode))
}
// 2757: decl @lune.@base.@TransUnit.TransUnit.processAddFunc
func (self *TransUnit_TransUnit) processAddFunc(_env *LnsEnv, isFunc bool,parentScope *Ast_Scope,name *Types_Token,typeInfoMut *Ast_TypeInfo,alt2typeMap *LnsMap)(*Ast_SymbolInfo, *TransUnitIF_NSInfo) {
    var accessMode LnsInt
    accessMode = typeInfoMut.FP.Get_accessMode(_env)
    var typeDataAccessor Ast_TypeDataAccessor
    typeDataAccessor = typeInfoMut.FP
    var typeInfo *Ast_TypeInfo
    typeInfo = typeInfoMut
    if accessMode == Ast_AccessMode__Global{
        parentScope = self.GlobalScope
    }
    var hasPrototype bool
    var nsInfo *TransUnitIF_NSInfo
    {
        _prottype := parentScope.FP.GetTypeInfoChild(_env, typeInfo.FP.Get_rawTxt(_env))
        if !Lns_IsNil( _prottype ) {
            prottype := _prottype.(*Ast_TypeInfo)
            var argTypeList *LnsList
            argTypeList = typeInfo.FP.Get_argTypeInfoList(_env)
            var retTypeInfoList *LnsList
            retTypeInfoList = typeInfo.FP.Get_retTypeInfoList(_env)
            var matched bool
            matched = true
            {
                var matchFlag LnsInt
                var err string
                matchFlag,err = Ast_TypeInfo_checkMatchType(_env, self.ProcessInfo, prottype.FP.Get_argTypeInfoList(_env), argTypeList, false, nil, alt2typeMap)
                if matchFlag != Ast_MatchType__Match{
                    self.FP.AddErrMess(_env, name.Pos, "mismatch functype param: " + err)
                    matched = false
                }
            }
            {
                var matchFlag LnsInt
                var err string
                matchFlag,err = Ast_TypeInfo_checkMatchType(_env, self.ProcessInfo, prottype.FP.Get_retTypeInfoList(_env), retTypeInfoList, false, nil, alt2typeMap)
                if matchFlag != Ast_MatchType__Match{
                    self.FP.AddErrMess(_env, name.Pos, "mismatch functype ret: " + err)
                    matched = false
                }
            }
            {
                var matchFlag bool
                var err LnsAny
                matchFlag,err = typeInfo.FP.CanEvalWith(_env, self.ProcessInfo, prottype, Ast_CanEvalType__SetOp, alt2typeMap)
                if Lns_op_not(matchFlag){
                    if Lns_isCondTrue( err){
                        self.FP.AddErrMess(_env, name.Pos, _env.GetVM().String_format("mismatch functype -- %s", []LnsAny{err}))
                    } else { 
                        self.FP.AddErrMess(_env, name.Pos, _env.GetVM().String_format("mismatch functype -- %s / %s", []LnsAny{typeInfo.FP.Get_display_stirng(_env), prottype.FP.Get_display_stirng(_env)}))
                    }
                    matched = false
                }
            }
            {
                if prottype.FP.Get_asyncMode(_env) != typeInfo.FP.Get_asyncMode(_env){
                    self.FP.AddErrMess(_env, name.Pos, _env.GetVM().String_format("mismatch async -- %s / %s", []LnsAny{Ast_Async_getTxt( prottype.FP.Get_asyncMode(_env)), Ast_Async_getTxt( typeInfo.FP.Get_asyncMode(_env))}))
                    matched = false
                }
            }
            {
                _workNsInfo := self.NsInfoMap.Get(prottype)
                if !Lns_IsNil( _workNsInfo ) {
                    workNsInfo := _workNsInfo.(*TransUnitIF_NSInfo)
                    nsInfo = workNsInfo
                    if nsInfo.FP.Get_nobody(_env){
                        hasPrototype = true
                        nsInfo.FP.Set_nobody(_env, false)
                    } else { 
                        hasPrototype = false
                    }
                } else {
                    hasPrototype = false
                    nsInfo = self.FP.NewNSInfoWithTypeData(_env, typeInfo, typeDataAccessor, name.Pos)
                }
            }
            if matched{
                (Lns_unwrap( Ast_NormalTypeInfoDownCastF(nsInfo.FP.Get_typeInfo(_env).FP)).(*Ast_NormalTypeInfo)).FP.SwitchScopeTo(_env, Lns_unwrap( typeInfo.FP.Get_scope(_env)).(*Ast_Scope))
                typeInfo = nsInfo.FP.Get_typeInfo(_env)
                typeDataAccessor = nsInfo.FP.Get_typeDataAccessor(_env)
            }
            if Lns_op_not(hasPrototype){
                if Lns_op_not(prottype.FP.Get_autoFlag(_env)){
                    self.FP.AddErrMess(_env, name.Pos, _env.GetVM().String_format("multiple define -- %s", []LnsAny{name.Txt}))
                }
            }
        } else {
            hasPrototype = false
            nsInfo = self.FP.NewNSInfoWithTypeData(_env, typeInfo, typeDataAccessor, name.Pos)
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) &&
        _env.SetStackVal( typeInfo.FP.Get_accessMode(_env) != Ast_AccessMode__Pri) ).(bool)){
        var classType *Ast_TypeInfo
        classType = typeInfo.FP.Get_parentInfo(_env)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.advertisedTypeSet.Has(Ast_TypeInfo2Stem(classType))) &&
            _env.SetStackVal( Lns_op_not(hasPrototype)) ).(bool)){
            self.FP.AddErrMess(_env, name.Pos, _env.GetVM().String_format("This class(%s) is used by advertise. You must declare the prototype of this method.", []LnsAny{classType.FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    var staticFlag bool
    staticFlag = typeInfo.FP.Get_staticFlag(_env)
    var mutable bool
    mutable = Ast_TypeInfo_isMut(_env, typeInfo)
    var funcSym LnsAny
    var shadowing LnsAny
    if isFunc{
        funcSym, shadowing = parentScope.FP.AddFunc(_env, self.ProcessInfo, name.Pos, typeInfo, accessMode, staticFlag, mutable)
        self.FP.errorShadowing(_env, name.Pos, shadowing)
    } else { 
        funcSym, shadowing = parentScope.FP.AddMethod(_env, self.ProcessInfo, name.Pos, typeInfo, accessMode, staticFlag)
    }
    self.FP.Set_curNsInfo(_env, nsInfo)
    return Lns_unwrap( (_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( funcSym) ||
        _env.SetStackVal( shadowing) ))).(*Ast_SymbolInfo), nsInfo
}
// 2875: decl @lune.@base.@TransUnit.TransUnit.addFuncBlockInfoList
func (self *TransUnit_TransUnit) addFuncBlockInfoList(_env *LnsEnv, funcBlockInfo *TransUnit_FuncBlockInfo) {
    Util_err(_env, "not implemented")
}
// 2887: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclFunc
func (self *TransUnit_TransUnit) analyzeDeclFunc(_env *LnsEnv, declFuncMode LnsInt,asyncLocked bool,abstractFlag bool,overrideFlag bool,accessMode LnsInt,staticFlag bool,classTypeInfo LnsAny,firstToken *Types_Token,name LnsAny) *Nodes_Node {
    var token *Types_Token
    token = self.FP.getToken(_env, nil)
    {
        __exp := name
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Types_Token)
            if _exp.Txt != "__main"{
                name = self.FP.checkSymbol(_env, _exp, TransUnit_SymbolMode__MustNot_)
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( declFuncMode == TransUnit_DeclFuncMode__Func) &&
                _env.SetStackVal( _exp.Txt == "main") ).(bool)){
                self.FP.addWarnMess(_env, _exp.Pos, "LuneScript's main function is __main.")
            }
        } else {
            if token.Txt != "("{
                if token.Txt != "__main"{
                    name = self.FP.checkSymbol(_env, token, TransUnit_SymbolMode__MustNot_)
                } else { 
                    name = token
                }
                token = self.FP.getToken(_env, nil)
            }
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(name)) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) ||
            _env.SetStackVal( abstractFlag) ||
            _env.SetStackVal( overrideFlag) ||
            _env.SetStackVal( staticFlag) ).(bool))) ).(bool)){
        self.FP.AddErrMess(_env, firstToken.Pos, "The anonymous function must be local.")
    }
    var outsizeOfClass bool
    outsizeOfClass = Ast_TypeInfo2Stem(classTypeInfo) == nil
    var needPopFlag bool
    needPopFlag = false
    if token.Txt == "."{
        needPopFlag = true
        if name != nil{
            name_156 := name.(*Types_Token)
            var className string
            className = name_156.Txt
            classTypeInfo = self.FP.Get_scope(_env).FP.GetTypeInfoChild(_env, className)
            if classTypeInfo != nil{
                classTypeInfo_159 := classTypeInfo.(*Ast_TypeInfo)
                self.FP.PushClassScope(_env, name_156.Pos, classTypeInfo_159, Lns_unwrap( self.Namespace2Scope.Get(classTypeInfo_159)).(*Ast_Scope))
            } else {
                self.FP.Error(_env, _env.GetVM().String_format("not found class -- %s", []LnsAny{className}))
            }
        } else {
            self.FP.Error(_env, "can't use '.' for any function name")
        }
        name = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
        token = self.FP.getToken(_env, nil)
    }
    var kind LnsInt
    kind = Nodes_NodeKind_get_DeclConstr(_env)
    var typeKind LnsInt
    typeKind = Ast_TypeInfoKind__Func
    if classTypeInfo != nil{
        if Lns_op_not(staticFlag){
            typeKind = Ast_TypeInfoKind__Method
        }
        if _switch0 := (Lns_unwrap( name).(*Types_Token)).Txt; _switch0 == "__init" {
            kind = Nodes_NodeKind_get_DeclConstr(_env)
            for _, _symbolInfo := range( self.FP.Get_scope(_env).FP.Get_symbol2SymbolInfoMap(_env).Items ) {
                symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if Lns_op_not(symbolInfo.FP.Get_staticFlag(_env)){
                    symbolInfo.FP.ClearValue(_env)
                }
            }
        } else if _switch0 == "__free" {
            kind = Nodes_NodeKind_get_DeclDestr(_env)
            if Lns_op_not(self.TargetLuaVer.FP.Get_canUseMetaGc(_env)){
                self.FP.AddErrMess(_env, firstToken.Pos, "this lua version is not support __free.")
            }
        } else {
            kind = Nodes_NodeKind_get_DeclMethod(_env)
        }
    } else {
        kind = Nodes_NodeKind_get_DeclFunc(_env)
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
    parentScope = self.FP.Get_scope(_env)
    var funcBodyScope *Ast_Scope
    funcBodyScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( staticFlag) &&
        _env.SetStackVal( classTypeInfo) )){
        self.analyzingStaticMethodArgsScope = funcBodyScope
    }
    var altTypeList *LnsList
    altTypeList = NewLnsList([]LnsAny{})
    if token.Txt == "<"{
        token, altTypeList = self.FP.analyzeDeclAlternateType(_env, false, token, accessMode)
        for _, _altType := range( altTypeList.Items ) {
            altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            funcBodyScope.FP.AddAlternate(_env, self.ProcessInfo, accessMode, altType.FP.Get_rawTxt(_env), firstToken.Pos, altType)
        }
    }
    self.FP.checkToken(_env, token, "(")
    var parentPub bool
    if classTypeInfo != nil{
        classTypeInfo_187 := classTypeInfo.(*Ast_TypeInfo)
        parentPub = Ast_isPubToExternal(_env, classTypeInfo_187.FP.Get_accessMode(_env))
    } else {
        parentPub = Ast_isPubToExternal(_env, accessMode)
    }
    var argList *LnsList
    argList = NewLnsList([]LnsAny{})
    token = self.FP.analyzeDeclArgList(_env, accessMode, funcBodyScope, argList, parentPub)
    self.FP.checkToken(_env, token, ")")
    token = self.FP.getToken(_env, nil)
    var mutMode LnsInt
    var asyncMode LnsAny
    token, mutMode, asyncMode = self.FP.getMutableAsync(_env, token)
    var pubToExtFlag bool
    pubToExtFlag = Ast_isPubToExternal(_env, accessMode)
    var argTypeList *LnsList
    argTypeList = NewLnsList([]LnsAny{})
    for _, _argNode := range( argList.Items ) {
        argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
        argTypeList.Insert(Ast_TypeInfo2Stem(argNode.FP.Get_expType(_env)))
    }
    var alt2typeMap *LnsMap
    alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false)
    if classTypeInfo != nil{
        classTypeInfo_198 := classTypeInfo.(*Ast_TypeInfo)
        alt2typeMap = classTypeInfo_198.FP.CreateAlt2typeMap(_env, false)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( kind == Nodes_NodeKind_get_DeclMethod(_env)) ||
            _env.SetStackVal( kind == Nodes_NodeKind_get_DeclConstr(_env)) ||
            _env.SetStackVal( kind == Nodes_NodeKind_get_DeclDestr(_env)) ).(bool){
            var workClass *Ast_TypeInfo
            workClass = classTypeInfo_198
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( kind == Nodes_NodeKind_get_DeclConstr(_env)) ||
                _env.SetStackVal( kind == Nodes_NodeKind_get_DeclDestr(_env)) ).(bool){
                mutMode = Ast_MutMode__Mut
            }
            if Lns_op_not(Ast_isPubToExternal(_env, workClass.FP.Get_accessMode(_env))){
                pubToExtFlag = false
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Ast_TypeInfo_isMut(_env, workClass)) &&
                _env.SetStackVal( mutMode != Ast_MutMode__Mut) ).(bool)){
                workClass = self.FP.createModifier(_env, workClass, mutMode)
            }
            if Lns_op_not(staticFlag){
                self.FP.Get_scope(_env).FP.Add(_env, self.ProcessInfo, Ast_SymbolKind__Arg, false, true, "self", nil, workClass, Ast_AccessMode__Pri, false, mutMode, true, false)
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(workClass.FP.Get_abstractFlag(_env))) &&
                _env.SetStackVal( abstractFlag) ).(bool)){
                self.FP.AddErrMess(_env, firstToken.Pos, "no abstract class does not have abstract method")
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( classTypeInfo_198.FP.IsInheritFrom(_env, self.ProcessInfo, Ast_builtinTypeRunner, nil)) &&
            _env.SetStackVal( Ast_isPubToExternal(_env, accessMode)) ).(bool)){
            for _index, _argNode := range( argList.Items ) {
                index := _index + 1
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                if Lns_op_not(self.FP.canBeAsyncParam(_env, argNode.FP.Get_expType(_env))){
                    self.FP.AddErrMess(_env, argNode.FP.Get_pos(_env), _env.GetVM().String_format("__Runner can't have the mutable argument with public method. -- %d: %s", []LnsAny{index, argNode.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        }
    }
    var retTypeInfoList *LnsList
    var retTypeNodeList *LnsList
    retTypeInfoList, token, retTypeNodeList = self.FP.analyzeRetTypeList(_env, pubToExtFlag, accessMode, token, parentPub)
    var parentNsInfo *TransUnitIF_NSInfo
    parentNsInfo = self.FP.Get_curNsInfo(_env)
    var funcName string
    funcName = ""
    if name != nil{
        name_216 := name.(*Types_Token)
        funcName = name_216.Txt
        if kind == Nodes_NodeKind_get_DeclFunc(_env){
            if _switch1 := accessMode; _switch1 == Ast_AccessMode__Pub || _switch1 == Ast_AccessMode__Global {
                if parentScope != self.ModuleScope{
                    self.FP.AddErrMess(_env, firstToken.Pos, "'global' or 'pub' function must exist top scope.")
                }
            }
        }
    }
    var typeInfo *Ast_TypeInfo
    var typeDataAccessor Ast_TypeDataAccessor
    var funcSym LnsAny
    var nsInfo *TransUnitIF_NSInfo
    {
        var workTypeInfo *Ast_NormalTypeInfo
        workTypeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, abstractFlag, false, funcBodyScope, typeKind, parentNsInfo.FP.Get_typeInfo(_env), parentNsInfo.FP.Get_typeDataAccessor(_env), false, false, staticFlag, accessMode, funcName, self.FP.getDefaultAsync(_env, typeKind, _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( classTypeInfo) ||
            _env.SetStackVal( self.FP.getCurrentClass(_env)) ).(*Ast_TypeInfo), asyncMode), altTypeList, argTypeList, retTypeInfoList, mutMode)
        if name != nil{
            name_227 := name.(*Types_Token)
            var workSym *Ast_SymbolInfo
            workSym, nsInfo = self.FP.processAddFunc(_env, kind == Nodes_NodeKind_get_DeclFunc(_env), funcBodyScope.FP.Get_outerScope(_env), name_227, &workTypeInfo.Ast_TypeInfo, alt2typeMap)
            typeDataAccessor = nsInfo.FP.Get_typeDataAccessor(_env)
            typeInfo = nsInfo.FP.Get_typeInfo(_env)
            funcSym = workSym
            if name_227.Txt == "__main"{
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( typeInfo.FP.Get_argTypeInfoList(_env).Len() != 1) ||
                    _env.SetStackVal( typeInfo.FP.Get_argTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind(_env) != Ast_TypeInfoKind__List) ||
                    _env.SetStackVal( typeInfo.FP.Get_argTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeString) ||
                    _env.SetStackVal( typeInfo.FP.Get_retTypeInfoList(_env).Len() != 1) ||
                    _env.SetStackVal( typeInfo.FP.Get_retTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeInt) ).(bool){
                    var mess string
                    mess = _env.GetVM().String_format("'__main' function's type has to be __main( argList:List<str> ) : int -- %s", []LnsAny{typeInfo.FP.Get_display_stirng(_env)})
                    self.FP.AddErrMess(_env, name_227.Pos, mess)
                }
            }
        } else {
            typeInfo = &workTypeInfo.Ast_TypeInfo
            typeDataAccessor = workTypeInfo.FP
            funcSym = nil
            nsInfo = self.FP.NewNSInfo(_env, &workTypeInfo.Ast_TypeInfo, firstToken.Pos)
        }
        self.Namespace2Scope.Set(typeInfo,funcBodyScope)
    }
    if asyncLocked{
        nsInfo.FP.IncLock(_env, Nodes_LockKind__AsyncLock)
    }
    if overrideFlag{
        if Lns_op_not(name){
            self.FP.AddErrMess(_env, firstToken.Pos, "can't override anonymous func")
        }
        if TransUnit_CantOverrideMethods.Has(funcName){
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("This method can't override -- %s", []LnsAny{funcName}))
        }
        {
            _overrideType := self.FP.Get_scope(_env).FP.Get_parent(_env).FP.GetTypeInfoField(_env, funcName, false, funcBodyScope, self.ScopeAccess)
            if !Lns_IsNil( _overrideType ) {
                overrideType := _overrideType.(*Ast_TypeInfo)
                for _, _err := range( self.FP.checkOverrideMethod(_env, overrideType, typeInfo).Items ) {
                    err := _err.(string)
                    self.FP.AddErrMess(_env, firstToken.Pos, err)
                }
            } else {
                self.FP.AddErrMess(_env, firstToken.Pos, "not found override -- " + funcName)
            }
        }
    } else { 
        if name != nil{
            name_244 := name.(*Types_Token)
            if Lns_op_not(TransUnit_CantOverrideMethods.Has(name_244.Txt)){
                if Lns_isCondTrue( self.FP.Get_scope(_env).FP.Get_parent(_env).FP.GetTypeInfoField(_env, name_244.Txt, false, funcBodyScope, Ast_ScopeAccess__Full)){
                    self.FP.AddErrMess(_env, firstToken.Pos, "mismatch override --" + funcName)
                } else { 
                    {
                        _ifFunc := self.FP.Get_scope(_env).FP.Get_parent(_env).FP.GetSymbolInfoIfField(_env, name_244.Txt, funcBodyScope, Ast_ScopeAccess__Full)
                        if !Lns_IsNil( _ifFunc ) {
                            ifFunc := _ifFunc.(*Ast_SymbolInfo)
                            if Lns_op_not(Lns_car(ifFunc.FP.Get_typeInfo(_env).FP.CanEvalWith(_env, self.ProcessInfo, typeInfo, Ast_CanEvalType__SetEq, alt2typeMap)).(bool)){
                                self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("mismatch method type -- %s", []LnsAny{funcName}))
                            }
                        }
                    }
                }
            }
        }
    }
    var node *Nodes_Node
    node = self.FP.createNoneNode(_env, firstToken.Pos)
    var body LnsAny
    body = nil
    var blockTokenList *LnsList
    blockTokenList = NewLnsList([]LnsAny{})
    var needFuncBlockInfo bool
    needFuncBlockInfo = false
    var inMacroExpand bool
    inMacroExpand = self.MacroCtrl.FP.IsInExpandMode(_env)
    var stmtNum LnsInt
    if token.Txt == ";"{
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( declFuncMode == TransUnit_DeclFuncMode__Module) ||
            _env.SetStackVal( declFuncMode == TransUnit_DeclFuncMode__Glue) ).(bool){
        } else { 
            if Lns_op_not(abstractFlag){
                nsInfo.FP.Set_nobody(_env, true)
            }
            if _env.NilAccFin(_env.NilAccPush(classTypeInfo) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) == Ast_TypeInfoKind__IF{
            } else { 
                if kind == Nodes_NodeKind_get_DeclMethod(_env){
                    kind = Nodes_NodeKind_get_ProtoMethod(_env)
                }
            }
        }
        stmtNum = 0
    } else { 
        if abstractFlag{
            self.FP.AddErrMess(_env, token.Pos, "abstract method can't have body.")
        }
        self.FP.Pushback(_env)
        var analyzingState LnsInt
        analyzingState = TransUnit_getAnalyzingState_24_(_env, typeInfo)
        funcBodyScope.FP.AddLocalVar(_env, self.ProcessInfo, false, false, "__func__", firstToken.Pos, Ast_builtinTypeString, Ast_MutMode__IMut)
        if _switch2 := self.analyzePhase; _switch2 == TransUnit_AnalyzePhase__Meta {
            var declScope LnsAny
            if classTypeInfo != nil{
                classTypeInfo_270 := classTypeInfo.(*Ast_TypeInfo)
                declScope = (Lns_unwrap( classTypeInfo_270.FP.Get_scope(_env)).(*Ast_Scope)).FP.Get_parent(_env)
            } else {
                declScope = (Lns_unwrap( typeInfo.FP.Get_scope(_env)).(*Ast_Scope)).FP.Get_parent(_env)
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( analyzingState == TransUnit_AnalyzingState__Constructor) ||
                _env.SetStackVal( self.MacroCtrl.FP.Get_isDeclaringMacro(_env)) ||
                _env.SetStackVal( declScope != self.ModuleScope) ||
                _env.SetStackVal( Types_Token2Stem(name) == nil) ).(bool){
            } else { 
                needFuncBlockInfo = true
            }
        } else if _switch2 == TransUnit_AnalyzePhase__Main || _switch2 == TransUnit_AnalyzePhase__Runner {
        }
        if needFuncBlockInfo{
            blockTokenList = self.FP.skipBlock(_env, true)
            stmtNum = blockTokenList.Len()
        } else { 
            var workBody *Nodes_BlockNode
            workBody = self.FP.analyzeFuncBlock(_env, analyzingState, firstToken, classTypeInfo, typeInfo, funcName, funcBodyScope, typeInfo.FP.Get_retTypeInfoList(_env))
            body = workBody
            self.FP.declFuncPostProcess(_env, typeInfo, classTypeInfo, workBody, funcBodyScope)
            stmtNum = nsInfo.FP.Get_stmtNum(_env)
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( staticFlag) &&
        _env.SetStackVal( classTypeInfo) )){
        self.analyzingStaticMethodArgsScope = nil
    }
    var createDeclFuncInfo func(_env *LnsEnv, funcKind LnsInt) *Nodes_DeclFuncInfo
    createDeclFuncInfo = func(_env *LnsEnv, funcKind LnsInt) *Nodes_DeclFuncInfo {
        var classDeclNode LnsAny
        if classTypeInfo != nil{
            classTypeInfo_285 := classTypeInfo.(*Ast_TypeInfo)
            classDeclNode = self.typeInfo2ClassNode.Get(classTypeInfo_285)
        } else {
            classDeclNode = nil
        }
        var declFuncInfo *Nodes_DeclFuncInfo
        declFuncInfo = NewNodes_DeclFuncInfo(_env, funcKind, classTypeInfo, classDeclNode, outsizeOfClass, name, funcSym, argList, orgStaticFlag, accessMode, asyncMode, body, retTypeInfoList, retTypeNodeList, self.Has__func__Symbol.Has(Ast_TypeInfo2Stem(typeInfo)), overrideFlag, stmtNum)
        if needFuncBlockInfo{
            var pos Types_Position
            pos = self.FP.getLineNo(_env, blockTokenList.GetAt(1).(Types_TokenDownCast).ToTypes_Token())
            var orgPos LnsAny
            if inMacroExpand{
                orgPos = pos
            } else { 
                orgPos = nil
            }
            var funcBlockInfo *TransUnit_FuncBlockInfo
            funcBlockInfo = NewTransUnit_FuncBlockInfo(_env, declFuncInfo, nsInfo, typeInfo, typeDataAccessor, funcBodyScope, blockTokenList, pos, orgPos)
            self.FP.addFuncBlockInfoList(_env, funcBlockInfo)
        }
        return declFuncInfo
    }
    if _switch3 := (kind); _switch3 == Nodes_NodeKind_get_DeclConstr(_env) {
        if Lns_isCondTrue( _env.NilAccFin(_env.NilAccPush(classTypeInfo) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.IsInheritFrom(_env, self.ProcessInfo, Ast_builtinTypeRunner, nil)})/* 3355:13 */)){
            if typeInfo.FP.Get_asyncMode(_env) != Ast_Async__Async{
                var message string
                message = _env.GetVM().String_format("The constructor of the class extended __Runner must be __async. -- %s", []LnsAny{typeInfo.FP.Get_parentInfo(_env).FP.GetTxt(_env, nil, nil, nil)})
                self.FP.AddErrMess(_env, firstToken.Pos, message)
            }
        }
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(_env, Nodes_FuncKind__Ctor)
        node = &Nodes_DeclConstrNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
    } else if _switch3 == Nodes_NodeKind_get_DeclDestr(_env) {
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(_env, Nodes_FuncKind__Dstr)
        node = &Nodes_DeclDestrNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
    } else if _switch3 == Nodes_NodeKind_get_DeclMethod(_env) {
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(_env, Nodes_FuncKind__Mtd)
        node = &Nodes_DeclMethodNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
    } else if _switch3 == Nodes_NodeKind_get_ProtoMethod(_env) {
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(_env, Nodes_FuncKind__Mtd)
        node = &Nodes_ProtoMethodNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
    } else if _switch3 == Nodes_NodeKind_get_DeclFunc(_env) {
        var info *Nodes_DeclFuncInfo
        info = createDeclFuncInfo(_env, Nodes_FuncKind__Func)
        node = &Nodes_DeclFuncNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), info).Nodes_Node
    } else {
        self.FP.Error(_env, _env.GetVM().String_format("illegal kind -- %d", []LnsAny{kind}))
    }
    self.Has__func__Symbol.Del(Ast_TypeInfo2Stem(typeInfo))
    self.FP.PopScope(_env)
    if needPopFlag{
        self.FP.addMethod(_env, Lns_unwrap( classTypeInfo).(*Ast_TypeInfo), node, funcName)
        self.FP.PopClass(_env)
    }
    return node
}
// 3415: decl @lune.@base.@TransUnit.TransUnit.createExpListNode
func (self *TransUnit_TransUnit) createExpListNode(_env *LnsEnv, orgExpList *Nodes_ExpListNode,newExpList *LnsList) *Nodes_ExpListNode {
    var newExpTypeList *LnsList
    newExpTypeList = NewLnsList([]LnsAny{})
    for _, _expNode := range( newExpList.Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        newExpTypeList.Insert(Ast_TypeInfo2Stem(expNode.FP.Get_expType(_env)))
    }
    if newExpList.GetAt(newExpList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_expTypeList(_env).Len() > 1{
        self.FP.AddErrMess(_env, orgExpList.FP.Get_pos(_env), _env.GetVM().String_format("illegal exp -- %d", []LnsAny{newExpList.GetAt(newExpList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_expTypeList(_env).Len()}))
    }
    {
        _mRetIndex := _env.NilAccFin(_env.NilAccPush(orgExpList.FP.Get_mRetExp(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
        if !Lns_IsNil( _mRetIndex ) {
            mRetIndex := _mRetIndex.(LnsInt)
            if mRetIndex > newExpList.Len(){
                self.FP.AddErrMess(_env, orgExpList.FP.Get_pos(_env), _env.GetVM().String_format("over index -- %d", []LnsAny{mRetIndex}))
            }
        }
    }
    return Nodes_ExpListNode_create(_env, self.NodeManager, orgExpList.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), newExpTypeList, newExpList, orgExpList.FP.Get_mRetExp(_env), orgExpList.FP.Get_followOn(_env))
}
// 3443: decl @lune.@base.@TransUnit.TransUnit.checkLiteralEmptyCollection
func (self *TransUnit_TransUnit) checkLiteralEmptyCollection(_env *LnsEnv, pos Types_Position,symbolName string,expType *Ast_TypeInfo) {
    for _, _itemType := range( expType.FP.Get_itemTypeInfoList(_env).Items ) {
        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if itemType == Ast_builtinTypeNone{
            self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("must set the item type of Collection. -- %s:%s", []LnsAny{symbolName, expType.FP.Get_srcTypeInfo(_env).FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
            break
        }
    }
}
// 3458: decl @lune.@base.@TransUnit.TransUnit.accessSymbol
func (self *TransUnit_TransUnit) accessSymbol(_env *LnsEnv, symbolInfo *Ast_SymbolInfo,canLeftExp bool) {
    if symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Fun{
        self.FP.Get_scope(_env).FP.AccessSymbol(_env, self.ModuleScope, symbolInfo)
        symbolInfo.FP.Set_posForModToRef(_env, symbolInfo.FP.Get_posForLatestMod(_env))
        if self.FP.Get_scope(_env).FP.IsClosureAccess(_env, self.ModuleScope, symbolInfo){
            self.closureFunList.Insert(TransUnit_ClosureFun2Stem(NewTransUnit_ClosureFun(_env, symbolInfo, self.FP.Get_scope(_env))))
        }
        {
            _scope := symbolInfo.FP.Get_typeInfo(_env).FP.Get_scope(_env)
            if !Lns_IsNil( _scope ) {
                scope := _scope.(*Ast_Scope)
                scope.FP.UpdateClosureRefPos(_env)
            }
        }
    } else { 
        self.FP.Get_scope(_env).FP.AccessSymbol(_env, self.ModuleScope, symbolInfo)
        if canLeftExp{
            self.accessSymbolSetQueue.FP.Add(_env, symbolInfo)
        } else { 
            symbolInfo.FP.Set_posForModToRef(_env, symbolInfo.FP.Get_posForLatestMod(_env))
        }
    }
}
// 3494: decl @lune.@base.@TransUnit.TransUnit.analyzeInitExp
func (self *TransUnit_TransUnit) analyzeInitExp(_env *LnsEnv, firstPos Types_Position,accessMode LnsInt,unwrapFlag bool,letVarList *LnsList,typeInfoList *LnsList)(*LnsList, *LnsList, *LnsList, LnsAny) {
    var expList LnsAny
    expList = nil
    var expectTypeList *LnsList
    expectTypeList = NewLnsList([]LnsAny{})
    for _, _varInfo := range( letVarList.Items ) {
        varInfo := _varInfo.(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
        expectTypeList.Insert(Ast_TypeInfo2Stem(Lns_unwrapDefault( _env.NilAccFin(_env.NilAccPush(varInfo.VarType) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_RefTypeNode).FP.Get_expType(_env)})), Ast_builtinTypeNone).(*Ast_TypeInfo)))
    }
    expList = self.FP.analyzeExpList(_env, false, false, false, nil, expectTypeList, nil)
    if Lns_op_not(expList){
        self.FP.Error(_env, "expList is nil")
    }
    var orgExpTypeList *LnsList
    orgExpTypeList = NewLnsList([]LnsAny{})
    if expList != nil{
        expList_337 := expList.(*Nodes_ExpListNode)
        if unwrapFlag{
            var hasNilable bool
            hasNilable = false
            for _index, _ := range( letVarList.Items ) {
                index := _index + 1
                if expList_337.FP.GetExpTypeAt(_env, index).FP.Get_nilable(_env){
                    hasNilable = true
                    break
                }
            }
            if Lns_op_not(hasNilable){
                self.FP.addWarnMess(_env, firstPos, "has no nilable")
            }
        }
        var workList *Nodes_ExpListNode
        workList = expList_337
        var updateExpList bool
        updateExpList = false
        var newExpList *LnsList
        newExpList = NewLnsList([]LnsAny{})
        for _index, _exp := range( workList.FP.Get_expList(_env).Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            newExpList.Insert(Nodes_Node2Stem(exp))
            if Lns_op_not(exp.FP.CanBeRight(_env, self.ProcessInfo)){
                self.FP.AddErrMess(_env, exp.FP.Get_effectivePos(_env), _env.GetVM().String_format("this node(%d) can not be r-value. -- %s", []LnsAny{index, Nodes_getNodeKindName(_env, exp.FP.Get_kind(_env))}))
            }
        }
        var expTypeList *LnsList
        expTypeList = NewLnsList([]LnsAny{})
        for _index, _expType := range( workList.FP.Get_expTypeList(_env).Items ) {
            index := _index + 1
            expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( index == workList.FP.Get_expTypeList(_env).Len()) &&
                _env.SetStackVal( expType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
                var dddItemType *Ast_TypeInfo
                dddItemType = Ast_builtinTypeStem_
                if expType.FP.Get_itemTypeInfoList(_env).Len() > 0{
                    dddItemType = expType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo(_env)
                }
                {
                    var _forFrom0 LnsInt = index
                    var _forTo0 LnsInt = letVarList.Len()
                    for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                        subIndex := _forWork0
                        var argType *Ast_TypeInfo
                        argType = typeInfoList.GetAt(subIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        var checkType *Ast_TypeInfo
                        checkType = dddItemType
                        if unwrapFlag{
                            checkType = dddItemType.FP.Get_nonnilableType(_env)
                        }
                        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( Lns_op_not(argType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeEmpty, nil, nil))) &&
                            _env.SetStackVal( Lns_op_not(Lns_car(argType.FP.CanEvalWith(_env, self.ProcessInfo, checkType, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool)){
                            self.FP.AddErrMess(_env, firstPos, _env.GetVM().String_format("unmatch value type (index = %d) %s <- %s", []LnsAny{subIndex, argType.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), dddItemType.FP.GetTxt(_env, nil, nil, nil)}))
                        }
                        expTypeList.Insert(Ast_TypeInfo2Stem(checkType))
                        orgExpTypeList.Insert(Ast_TypeInfo2Stem(dddItemType))
                    }
                }
            } else { 
                var expTypeInfo *Ast_TypeInfo
                expTypeInfo = expType
                if expType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    var itemList *LnsList
                    itemList = expType.FP.Get_itemTypeInfoList(_env)
                    if itemList.Len() > 0{
                        expTypeInfo = itemList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    } else { 
                        expTypeInfo = Ast_builtinTypeStem_
                    }
                }
                orgExpTypeList.Insert(Ast_TypeInfo2Stem(expTypeInfo))
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( expTypeInfo == Ast_builtinTypeNil) &&
                    _env.SetStackVal( index <= typeInfoList.Len()) ).(bool)){
                    orgExpTypeList.Set(index,typeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo(_env))
                }
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( unwrapFlag) &&
                    _env.SetStackVal( expTypeInfo.FP.Get_nilable(_env)) ).(bool)){
                    expTypeInfo = expTypeInfo.FP.Get_nonnilableType(_env)
                }
                if index <= typeInfoList.Len(){
                    var varType *Ast_TypeInfo
                    varType = typeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    var alt2typeMap *LnsMap
                    alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false)
                    if varType.FP.Get_kind(_env) == Ast_TypeInfoKind__Box{
                        alt2typeMap = varType.FP.CreateAlt2typeMap(_env, true)
                    }
                    Ast_CanEvalCtrlTypeInfo_setupNeedAutoBoxing(_env, alt2typeMap, self.ProcessInfo)
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(varType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeEmpty, nil, nil))) &&
                        _env.SetStackVal( Lns_op_not((_env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( unwrapFlag) &&
                            _env.SetStackVal( expTypeInfo.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeNil, nil, nil)) ).(bool)))) ).(bool)){
                        var canEval bool
                        var mess LnsAny
                        canEval,mess = varType.FP.CanEvalWith(_env, self.ProcessInfo, expTypeInfo, Ast_CanEvalType__SetOp, alt2typeMap)
                        if Lns_op_not(canEval){
                            self.FP.AddErrMess(_env, firstPos, _env.GetVM().String_format("unmatch value type (index:%d) %s <- %s%s", []LnsAny{index, varType.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), expTypeInfo.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( mess) &&
                                _env.SetStackVal( _env.GetVM().String_format(" -- %s", []LnsAny{mess})) ||
                                _env.SetStackVal( "") ).(string)}))
                        }
                    }
                    if varType.FP.Get_kind(_env) == Ast_TypeInfoKind__Box{
                        typeInfoList.Set(index,self.ProcessInfo.FP.CreateBox(_env, accessMode, expTypeInfo))
                    }
                    if Ast_CanEvalCtrlTypeInfo_canAutoBoxing(_env, varType, expTypeInfo){
                        updateExpList = true
                        var exp *Nodes_Node
                        exp = newExpList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                        newExpList.Set(index,&Nodes_BoxingNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(varType)}), exp).Nodes_Node)
                        if Lns_op_not(Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(_env, alt2typeMap, 1)){
                            self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("auto boxing error %s <- %s", []LnsAny{varType.FP.GetTxt(_env, nil, nil, nil), expTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
                        }
                    } else { 
                        if Lns_op_not(Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(_env, alt2typeMap, 0)){
                            self.FP.AddErrMess(_env, newExpList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_pos(_env), _env.GetVM().String_format("illegal auto boxing error %s <- %s", []LnsAny{varType.FP.GetTxt(_env, nil, nil, nil), expTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
                        }
                    }
                }
                expTypeList.Insert(Ast_TypeInfo2Stem(expTypeInfo))
            }
        }
        if updateExpList{
            workList = self.FP.createExpListNode(_env, workList, newExpList)
        }
        {
            var alt2typeMap *LnsMap
            alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false)
            {
                __exp := self.FP.checkImplicitCast(_env, alt2typeMap, true, typeInfoList, workList, TransUnit_checkImplicitCastCallback_10_(TransUnit_analyzeInitExp___anonymous_0_))
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Nodes_ExpListNode)
                    workList = _exp
                }
            }
        }
        for _index, _varType := range( typeInfoList.Items ) {
            index := _index + 1
            varType := _varType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index > expTypeList.Len(){
                if Lns_op_not(varType.FP.Get_nilable(_env)){
                    self.FP.AddErrMess(_env, firstPos, _env.GetVM().String_format("unmatch value type (index:%d) %s <- nil", []LnsAny{index, varType.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
                }
            }
        }
        for _index, _typeInfo := range( expTypeList.Items ) {
            index := _index + 1
            typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( typeInfoList.Len() < index) ||
                _env.SetStackVal( typeInfoList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeEmpty, nil, nil)) ).(bool){
                var workPos Types_Position
                var workType *Ast_TypeInfo
                var workName string
                if index <= letVarList.Len(){
                    var letVar *TransUnit_LetVarInfo
                    letVar = letVarList.GetAt(index).(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
                    workPos = letVar.VarName.Pos
                    workName = letVar.VarName.Txt
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Ast_TypeInfo_isMut(_env, typeInfo)) &&
                        _env.SetStackVal( Lns_op_not(Ast_isMutable(_env, letVar.Mutable))) ).(bool)){
                        workType = self.FP.createModifier(_env, typeInfo, Ast_MutMode__IMutRe)
                    } else { 
                        workType = typeInfo
                    }
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( self.analyzeMode == TransUnit_AnalyzeMode__Inquire) &&
                        _env.SetStackVal( self.FP.isTargetToken(_env, letVar.VarName)) ).(bool)){
                        self.FP.dumpSymbolType(_env, accessMode, letVar.VarName.Txt, workType)
                    }
                } else { 
                    workType = typeInfo
                    workPos = firstPos
                    workName = ""
                }
                typeInfoList.Set(index,workType)
                if _switch0 := workType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Func {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( expTypeList.Len() != 1) ||
                        _env.SetStackVal( workType.FP.Get_rawTxt(_env) != "") ).(bool){
                        self.FP.AddErrMess(_env, firstPos, _env.GetVM().String_format("must set the type of variable for function. -- %s", []LnsAny{workName}))
                    }
                } else if _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array || _switch0 == Ast_TypeInfoKind__Set || _switch0 == Ast_TypeInfoKind__Map {
                    self.FP.checkLiteralEmptyCollection(_env, workPos, workName, workType)
                }
            }
        }
        return typeInfoList, letVarList, orgExpTypeList, workList
    }
    return typeInfoList, letVarList, orgExpTypeList, nil
}
// 3755: decl @lune.@base.@TransUnit.TransUnit.analyzeLetAndInitExp
func (self *TransUnit_TransUnit) analyzeLetAndInitExp(_env *LnsEnv, firstPos Types_Position,letFlag bool,initMutable LnsInt,accessMode LnsInt,unwrapFlag bool)(*LnsList, *LnsList, *LnsList, LnsAny) {
    var typeInfoList *LnsList
    typeInfoList = NewLnsList([]LnsAny{})
    var letVarList *LnsList
    letVarList = NewLnsList([]LnsAny{})
    var nextToken *Types_Token
    nextToken = Parser_getEofToken(_env)
    if letFlag{
        for {
            var mutable LnsInt
            mutable = initMutable
            nextToken = self.FP.getToken(_env, nil)
            if _switch0 := nextToken.Txt; _switch0 == "mut" {
                mutable = Ast_MutMode__Mut
                nextToken = self.FP.getToken(_env, nil)
            }
            var varName *Types_Token
            varName = self.FP.checkSymbol(_env, nextToken, TransUnit_SymbolMode__MustNot_Or_)
            nextToken = self.FP.getToken(_env, nil)
            var typeInfo *Ast_TypeInfo
            typeInfo = Ast_builtinTypeEmpty
            if nextToken.Txt == ":"{
                var refType *Nodes_RefTypeNode
                refType = self.FP.analyzeRefType(_env, accessMode, false, Ast_isPubToExternal(_env, accessMode))
                letVarList.Insert(TransUnit_LetVarInfo2Stem(NewTransUnit_LetVarInfo(_env, mutable, varName, refType)))
                typeInfo = refType.FP.Get_expType(_env)
                nextToken = self.FP.getToken(_env, nil)
            } else { 
                letVarList.Insert(TransUnit_LetVarInfo2Stem(NewTransUnit_LetVarInfo(_env, _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( Ast_isMutable(_env, mutable)) &&
                    _env.SetStackVal( mutable) ||
                    _env.SetStackVal( Ast_MutMode__IMutRe) ).(LnsInt), varName, nil)))
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(typeInfo.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeEmpty, nil, nil))) &&
                _env.SetStackVal( Ast_TypeInfo_isMut(_env, typeInfo)) &&
                _env.SetStackVal( Lns_op_not(Ast_isMutable(_env, mutable))) &&
                _env.SetStackVal( self.Ctrl_info.LegacyMutableControl) ).(bool)){
                typeInfo = self.FP.createModifier(_env, typeInfo, Ast_MutMode__IMutRe)
            }
            typeInfoList.Insert(Ast_TypeInfo2Stem(typeInfo))
            if nextToken.Txt != ","{ break }
        }
    } else { 
        for  {
            var symbolToken *Types_Token
            symbolToken = self.FP.getToken(_env, nil)
            nextToken = self.FP.getToken(_env, nil)
            var verSym LnsAny
            verSym = self.FP.Get_scope(_env).FP.GetSymbolTypeInfo(_env, symbolToken.Txt, self.FP.Get_scope(_env), self.ModuleScope, self.ScopeAccess)
            if verSym != nil{
                verSym_436 := verSym.(*Ast_SymbolInfo)
                letVarList.Insert(TransUnit_LetVarInfo2Stem(NewTransUnit_LetVarInfo(_env, verSym_436.FP.Get_mutMode(_env), symbolToken, nil)))
                typeInfoList.Insert(Ast_TypeInfo2Stem(verSym_436.FP.Get_typeInfo(_env)))
            } else {
                self.FP.AddErrMess(_env, symbolToken.Pos, _env.GetVM().String_format("not found symbol -- %s", []LnsAny{symbolToken.Txt}))
            }
            if nextToken.Txt != ","{
                break
            }
        }
    }
    if nextToken.Txt != "="{
        self.FP.Pushback(_env)
        var orgExpTypeList *LnsList
        orgExpTypeList = NewLnsList([]LnsAny{})
        return typeInfoList, letVarList, orgExpTypeList, nil
    }
    return self.FP.analyzeInitExp(_env, firstPos, accessMode, unwrapFlag, letVarList, typeInfoList)
}
// 3833: decl @lune.@base.@TransUnit.TransUnit.analyzeDeclVar
func (self *TransUnit_TransUnit) analyzeDeclVar(_env *LnsEnv, mode LnsInt,accessMode LnsInt,firstToken *Types_Token) *Nodes_Node {
    var unwrapFlag bool
    unwrapFlag = false
    var token *Types_Token
    var continueFlag bool
    token,continueFlag = self.FP.getContinueToken(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( continueFlag) &&
        _env.SetStackVal( token.Txt == "!") ).(bool)){
        unwrapFlag = true
    } else { 
        self.FP.Pushback(_env)
        if mode != Nodes_DeclVarMode__Let{
            Util_log(_env, "need '!'")
        }
    }
    if accessMode == Ast_AccessMode__Pub{
        if self.FP.Get_scope(_env) != self.ModuleScope{
            self.FP.AddErrMess(_env, firstToken.Pos, "'pub' variable must exist top scope.")
        }
    }
    var typeInfoList *LnsList
    var letVarList *LnsList
    var orgExpTypeList *LnsList
    var expList LnsAny
    typeInfoList,letVarList,orgExpTypeList,expList = self.FP.analyzeLetAndInitExp(_env, firstToken.Pos, mode == Nodes_DeclVarMode__Let, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode == Nodes_DeclVarMode__Sync) &&
        _env.SetStackVal( Ast_MutMode__Mut) ||
        _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt), accessMode, unwrapFlag)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode == Nodes_DeclVarMode__Let) &&
        _env.SetStackVal( typeInfoList.Len() == 1) ).(bool)){
        if expList != nil{
            expList_456 := expList.(*Nodes_ExpListNode)
            var typeInfo *Ast_TypeInfo
            typeInfo = typeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var letVaInfo *TransUnit_LetVarInfo
            letVaInfo = letVarList.GetAt(1).(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( expList_456.FP.Get_expList(_env).Len() == 1) &&
                _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Func) ).(bool)){
                var valExp *Nodes_Node
                valExp = expList_456.FP.Get_expList(_env).GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
                {
                    _macroExp := Nodes_ExpMacroExpNodeDownCastF(valExp.FP)
                    if !Lns_IsNil( _macroExp ) {
                        macroExp := _macroExp.(*Nodes_ExpMacroExpNode)
                        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( macroExp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Func) &&
                            _env.SetStackVal( macroExp.FP.Get_stmtList(_env).Len() == 1) ).(bool)){
                            valExp = macroExp.FP.Get_stmtList(_env).GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
                        }
                    }
                }
                {
                    _declNode := Nodes_DeclFuncNodeDownCastF(valExp.FP)
                    if !Lns_IsNil( _declNode ) {
                        declNode := _declNode.(*Nodes_DeclFuncNode)
                        if Lns_op_not(declNode.FP.Get_declInfo(_env).FP.Get_name(_env)){
                            if Ast_isMutable(_env, letVaInfo.Mutable){
                                self.FP.AddErrMess(_env, letVaInfo.VarName.Pos, _env.GetVM().String_format("Any function can't be mutable. -- %s", []LnsAny{letVaInfo.VarName.Txt}))
                            }
                            var letVarInfo *TransUnit_LetVarInfo
                            letVarInfo = letVarList.GetAt(1).(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
                            var newTypeInfo *Ast_NormalTypeInfo
                            newTypeInfo = self.ProcessInfo.FP.CreateFuncAsync(_env, typeInfo.FP.Get_abstractFlag(_env), false, Lns_unwrap( self.Namespace2Scope.Get(typeInfo)).(*Ast_Scope), typeInfo.FP.Get_kind(_env), typeInfo.FP.Get_parentInfo(_env), self.FP.GetNSInfo(_env, typeInfo.FP.Get_parentInfo(_env)).FP.Get_typeDataAccessor(_env), false, false, typeInfo.FP.Get_staticFlag(_env), accessMode, letVarInfo.VarName.Txt, typeInfo.FP.Get_asyncMode(_env), typeInfo.FP.Get_itemTypeInfoList(_env), typeInfo.FP.Get_argTypeInfoList(_env), typeInfo.FP.Get_retTypeInfoList(_env), typeInfo.FP.Get_mutMode(_env))
                            var funcSym *Ast_SymbolInfo
                            funcSym = TransUnit_convExp3_6912(Lns_2DDD(self.FP.processAddFunc(_env, true, self.FP.Get_scope(_env), letVarInfo.VarName, &newTypeInfo.Ast_TypeInfo, Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false))))
                            self.NodeManager.FP.DelNode(_env, &declNode.Nodes_Node)
                            var declInfo *Nodes_DeclFuncInfo
                            declInfo = Nodes_DeclFuncInfo_createFrom(_env, declNode.FP.Get_declInfo(_env), letVarInfo.VarName, funcSym)
                            return &Nodes_DeclFuncNode_create(_env, self.NodeManager, declNode.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_NormalTypeInfo2Stem(newTypeInfo)}), declInfo).Nodes_Node
                        }
                    }
                }
            }
        }
    }
    var syncScope *Ast_Scope
    syncScope = self.FP.Get_scope(_env)
    if mode == Nodes_DeclVarMode__Sync{
        syncScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    }
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.Get_curNsInfo(_env)
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
        varInfo = NewNodes_VarInfo(_env, varName, letVarInfo.VarType, typeInfo)
        varList.Insert(Nodes_VarInfo2Stem(varInfo))
        if Ast_isPubToExternal(_env, accessMode){
            self.FP.checkPublic(_env, varName.Pos, typeInfo)
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(letVarInfo.VarType)) &&
            _env.SetStackVal( typeInfo.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeNil, nil, nil)) ).(bool)){
            self.FP.AddErrMess(_env, varName.Pos, _env.GetVM().String_format("need type -- %s", []LnsAny{varName.Txt}))
        }
        if mode == Nodes_DeclVarMode__Sync{
            {
                _symInfo := self.FP.Get_scope(_env).FP.GetSymbolInfo(_env, varName.Txt, self.FP.Get_scope(_env), true, self.ScopeAccess)
                if !Lns_IsNil( _symInfo ) {
                    symInfo := _symInfo.(*Ast_SymbolInfo)
                    syncSymbolList.Insert(Ast_SymbolInfo2Stem(symInfo))
                }
            }
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( mode == Nodes_DeclVarMode__Let) ||
            _env.SetStackVal( mode == Nodes_DeclVarMode__Sync) ).(bool){
            if mode == Nodes_DeclVarMode__Let{
                self.FP.checkShadowing(_env, varName.Pos, varName.Txt, self.FP.Get_scope(_env))
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
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(unwrapFlag)) &&
                _env.SetStackVal( orgExpType != Ast_builtinTypeEmpty) ||
                _env.SetStackVal( unwrapFlag) &&
                _env.SetStackVal( Lns_op_not(orgExpType.FP.Get_nilable(_env))) ).(bool){
                hasValue = true
            }
            self.FP.Get_scope(_env).FP.AddVar(_env, self.ProcessInfo, accessMode, varName.Txt, varName.Pos, typeInfo, letVarInfo.Mutable, hasValue)
            if typeInfo.FP.Get_asyncMode(_env) == Ast_Async__Transient{
                self.FP.AddErrMess(_env, varName.Pos, _env.GetVM().String_format("can't set the __trans type -- index:%d, %s", []LnsAny{index, typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            }
        }
        symbolInfoList.Insert(Ast_SymbolInfo2Stem(nsInfo.FP.RegisterSym(_env, Lns_unwrap( self.FP.Get_scope(_env).FP.GetSymbolInfo(_env, varName.Txt, self.FP.Get_scope(_env), true, self.ScopeAccess)).(*Ast_SymbolInfo))))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode != Nodes_DeclVarMode__Sync) &&
        _env.SetStackVal( self.macroScope) )){
        self.MacroCtrl.FP.RegistVar(_env, symbolInfoList)
    }
    var unwrapBlock LnsAny
    unwrapBlock = nil
    var thenBlock LnsAny
    thenBlock = nil
    if unwrapFlag{
        var scope *Ast_Scope
        scope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
        for _index, _letVarInfo := range( letVarList.Items ) {
            index := _index + 1
            letVarInfo := _letVarInfo.(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
            if letVarInfo.VarName.Txt != "_"{
                self.FP.AddLocalVar(_env, letVarInfo.VarName.Pos, false, true, "_" + letVarInfo.VarName.Txt, orgExpTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut, nil)
            }
        }
        unwrapBlock = self.FP.analyzeBlock(_env, Nodes_BlockKind__LetUnwrap, TransUnit_TentativeMode__Start, scope, nil)
        self.FP.PopScope(_env)
        if unwrapBlock != nil{
            unwrapBlock_507 := unwrapBlock.(*Nodes_BlockNode)
            if _switch0 := mode; _switch0 == Nodes_DeclVarMode__Let || _switch0 == Nodes_DeclVarMode__Sync {
                var breakKind LnsInt
                breakKind = unwrapBlock_507.FP.GetBreakKind(_env, Nodes_CheckBreakMode__Normal)
                for _, _symbolInfo := range( symbolInfoList.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    if breakKind != Nodes_BreakKind__None{
                        self.tentativeSymbol.FP.CheckAndExclude(_env, symbolInfo)
                        symbolInfo.FP.UpdateValue(_env, symbolInfo.FP.Get_pos(_env))
                    } else { 
                        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( symbolInfo.FP.Get_name(_env) != "_") &&
                            _env.SetStackVal( Lns_op_not(self.tentativeSymbol.FP.CheckAndExclude(_env, symbolInfo))) ).(bool)){
                            if Lns_op_not(symbolInfo.FP.Get_hasValueFlag(_env)){
                                self.FP.AddErrMess(_env, unwrapBlock_507.FP.Get_pos(_env), "This variable isn't set -- " + (symbolInfo.FP.Get_name(_env)))
                            }
                        }
                    }
                }
            } else if _switch0 == Nodes_DeclVarMode__Unwrap {
                for _, _symbolInfo := range( symbolInfoList.Items ) {
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    symbolInfo.FP.UpdateValue(_env, firstToken.Pos)
                }
            }
        }
        token = self.FP.getToken(_env, true)
        if token.Txt == "then"{
            thenBlock = self.FP.analyzeBlock(_env, Nodes_BlockKind__LetUnwrapThenDo, TransUnit_TentativeMode__Finish, scope, nil)
        } else { 
            self.FP.Pushback(_env)
            self.FP.finishTentativeSymbol(_env, true)
        }
    }
    var syncBlock LnsAny
    syncBlock = nil
    if mode == Nodes_DeclVarMode__Sync{
        self.FP.checkNextToken(_env, "do")
        syncBlock = self.FP.analyzeBlock(_env, Nodes_BlockKind__LetUnwrapThenDo, TransUnit_TentativeMode__Simple, syncScope, nil)
        self.FP.PopScope(_env)
    }
    self.FP.checkNextToken(_env, ";")
    var node *Nodes_DeclVarNode
    node = Nodes_DeclVarNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), mode, accessMode, false, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncSymbolList, syncBlock)
    return &node.Nodes_Node
}
// 4075: decl @lune.@base.@TransUnit.TransUnit.analyzeIfUnwrap
func (self *TransUnit_TransUnit) analyzeIfUnwrap(_env *LnsEnv, firstToken *Types_Token) *Nodes_IfUnwrapNode {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
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
        workTypeInfoList, letVarList, _, workExpList = self.FP.analyzeLetAndInitExp(_env, firstToken.Pos, true, Ast_MutMode__IMut, Ast_AccessMode__Local, true)
    } else { 
        self.FP.Pushback(_env)
        var tmpTypeInfoList *LnsList
        tmpTypeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeEmpty)})
        var tmpLetVarList *LnsList
        tmpLetVarList = NewLnsList([]LnsAny{TransUnit_LetVarInfo2Stem(NewTransUnit_LetVarInfo(_env, Ast_MutMode__Mut, NewTypes_Token(_env, Types_TokenKind__Symb, "_exp", firstToken.Pos, false, nil), nil))})
        workTypeInfoList, letVarList, _, workExpList = self.FP.analyzeInitExp(_env, firstToken.Pos, Ast_AccessMode__Local, true, tmpLetVarList, tmpTypeInfoList)
    }
    typeInfoList = workTypeInfoList
    if workExpList != nil{
        workExpList_538 := workExpList.(*Nodes_ExpListNode)
        expList = workExpList_538
    } else {
        self.FP.AddErrMess(_env, nextToken.Pos, "if! let has illegal init val.")
        self.FP.Error(_env, "system error")
    }
    for _, _varInfo := range( letVarList.Items ) {
        varInfo := _varInfo.(TransUnit_LetVarInfoDownCast).ToTransUnit_LetVarInfo()
        varNameList.Insert(Types_Token2Stem(varInfo.VarName))
    }
    var scope *Ast_Scope
    scope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.Get_curNsInfo(_env)
    for _index, _expType := range( typeInfoList.Items ) {
        index := _index + 1
        expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > varNameList.Len(){
            break
        }
        var varName *Types_Token
        varName = varNameList.GetAt(index).(Types_TokenDownCast).ToTypes_Token()
        varList.Insert(Ast_SymbolInfo2Stem(nsInfo.FP.RegisterSym(_env, self.FP.AddLocalVar(_env, varName.Pos, false, true, varName.Txt, expType, Ast_MutMode__IMut, nil))))
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(_env, Nodes_BlockKind__IfUnwrap, TransUnit_TentativeMode__Start, scope, nil)
    self.FP.PopScope(_env)
    var elseBlock LnsAny
    elseBlock = nil
    nextToken = self.FP.getToken(_env, true)
    if nextToken.Txt == "else"{
        elseBlock = self.FP.analyzeBlock(_env, Nodes_BlockKind__Else, TransUnit_TentativeMode__Finish, nil, nil)
    } else { 
        self.FP.finishTentativeSymbol(_env, false)
        self.FP.Pushback(_env)
    }
    var hasCond bool
    hasCond = false
    for _index, _expNode := range( expList.FP.Get_expList(_env).Items ) {
        index := _index + 1
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        if index != expList.FP.Get_expList(_env).Len(){
            if Ast_isConditionalbe(_env, self.ProcessInfo, expNode.FP.Get_expType(_env)){
                hasCond = true
                break
            }
        } else { 
            for _, _expType := range( expNode.FP.Get_expTypeList(_env).Items ) {
                expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if Ast_isConditionalbe(_env, self.ProcessInfo, expType){
                    hasCond = true
                    break
                }
            }
        }
    }
    if Lns_op_not(hasCond){
        self.FP.AddErrMess(_env, firstToken.Pos, "This condition never be false")
    }
    for _, _varSym := range( varList.Items ) {
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if _switch0 := varSym.FP.Get_name(_env); _switch0 == "_" || _switch0 == "_exp" {
        } else {
            if Lns_op_not(varSym.FP.Get_posForModToRef(_env)){
                self.FP.addWarnMess(_env, Lns_unwrap( varSym.FP.Get_pos(_env)).(Types_Position), _env.GetVM().String_format("This symbol has no referer -- %s", []LnsAny{varSym.FP.Get_name(_env)}))
            }
        }
    }
    return Nodes_IfUnwrapNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), varList, expList, block, elseBlock)
}
// 4178: decl @lune.@base.@TransUnit.TransUnit.analyzeWhen
func (self *TransUnit_TransUnit) analyzeWhen(_env *LnsEnv, firstToken *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    var continueFlag bool
    nextToken,continueFlag = self.FP.getContinueToken(_env)
    if Lns_op_not((_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( continueFlag) &&
        _env.SetStackVal( nextToken.Txt == "!") ).(bool))){
        self.FP.Pushback(_env)
        self.FP.AddErrMess(_env, nextToken.Pos, "'when' need '!'")
    }
    var symListNode *Nodes_ExpListNode
    symListNode = self.FP.analyzeExpList(_env, false, false, false, nil, nil, nil)
    var scope *Ast_Scope
    scope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var symPairList *LnsList
    symPairList = NewLnsList([]LnsAny{})
    for _, _expNode := range( symListNode.FP.Get_expList(_env).Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        {
            _refNode := Nodes_ExpRefNodeDownCastF(expNode.FP)
            if !Lns_IsNil( _refNode ) {
                refNode := _refNode.(*Nodes_ExpRefNode)
                if expNode.FP.Get_expType(_env).FP.Get_nilable(_env){
                    var symbolInfo *Ast_SymbolInfo
                    symbolInfo = refNode.FP.Get_symbolInfo(_env)
                    var newSymbolInfo *Ast_SymbolInfo
                    newSymbolInfo = self.FP.AddLocalVar(_env, firstToken.Pos, false, expNode.FP.CanBeLeft(_env), refNode.FP.Get_symbolInfo(_env).FP.Get_name(_env), expNode.FP.Get_expType(_env).FP.Get_nonnilableType(_env), Ast_MutMode__IMut, true)
                    symPairList.Insert(Nodes_UnwrapSymbolPair2Stem(NewNodes_UnwrapSymbolPair(_env, symbolInfo, newSymbolInfo)))
                } else { 
                    self.FP.AddErrMess(_env, expNode.FP.Get_pos(_env), _env.GetVM().String_format("This type isn't nilable. -- %s", []LnsAny{expNode.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            } else {
                self.FP.AddErrMess(_env, expNode.FP.Get_pos(_env), "'when' support only local variables or arguments.")
            }
        }
    }
    var block *Nodes_BlockNode
    block = self.FP.analyzeBlock(_env, Nodes_BlockKind__When, TransUnit_TentativeMode__Start, scope, nil)
    self.FP.PopScope(_env)
    var elseBlock LnsAny
    elseBlock = nil
    nextToken = self.FP.getToken(_env, true)
    if nextToken.Txt == "else"{
        elseBlock = self.FP.analyzeBlock(_env, Nodes_BlockKind__Else, TransUnit_TentativeMode__Finish, nil, nil)
    } else { 
        self.FP.finishTentativeSymbol(_env, false)
        self.FP.Pushback(_env)
    }
    return &Nodes_WhenNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), symPairList, block, elseBlock).Nodes_Node
}
// 4239: decl @lune.@base.@TransUnit.TransUnit.processFuncBlockInfo
func (self *TransUnit_TransUnit) processFuncBlockInfo(_env *LnsEnv, funcBlockCtlIF TransUnit_FuncBlockCtlIF,streamName string) *LnsMap {
    var resultMap *LnsMap
    resultMap = NewLnsMap( map[LnsAny]LnsAny{})
    var bakParser *Parser_DefaultPushbackParser
    bakParser = self.Parser
    var outerScope *Ast_Scope
    outerScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    for  {
        var funcBlockInfo *TransUnit_FuncBlockInfo
        
        {
            _funcBlockInfo := funcBlockCtlIF.GetNext(_env)
            if _funcBlockInfo == nil{
                break
            } else {
                funcBlockInfo = _funcBlockInfo.(*TransUnit_FuncBlockInfo)
            }
        }
        var typeInfo *Ast_TypeInfo
        typeInfo = funcBlockInfo.FP.Get_funcType(_env)
        self.Parser = NewParser_DefaultPushbackParser(_env, &NewParser_TokenListParser(_env, funcBlockInfo.FP.Get_tokenList(_env), streamName, funcBlockInfo.FP.Get_tokenList(_env).GetAt(1).(Types_TokenDownCast).ToTypes_Token().Pos.OrgPos).Parser_Parser)
        var declFuncInfo *Nodes_DeclFuncInfo
        declFuncInfo = funcBlockInfo.FP.Get_declFuncInfo(_env)
        var classTypeInfo LnsAny
        classTypeInfo = declFuncInfo.FP.Get_classTypeInfo(_env)
        self.FuncBlockInfoLinkNo = funcBlockInfo.FP.Get_orgPos(_env)
        var funcBodyScope *Ast_ScopeWithRef
        funcBodyScope = NewAst_ScopeWithRef(_env, self.ProcessInfo, outerScope, funcBlockInfo.FP.Get_funcScope(_env), Ast_ScopeKind__Other, nil, nil)
        self.FP.SetScope(_env, &funcBodyScope.Ast_Scope, TransUnitIF_SetNSInfo__FromScope)
        var workBody *Nodes_BlockNode
        workBody = self.FP.analyzeFuncBlock(_env, TransUnit_getAnalyzingState_24_(_env, typeInfo), funcBlockInfo.FP.Get_tokenList(_env).GetAt(1).(Types_TokenDownCast).ToTypes_Token(), classTypeInfo, typeInfo, typeInfo.FP.Get_rawTxt(_env), &funcBodyScope.Ast_Scope, typeInfo.FP.Get_retTypeInfoList(_env))
        self.FP.declFuncPostProcess(_env, typeInfo, classTypeInfo, workBody, &funcBodyScope.Ast_Scope)
        self.FuncBlockInfoLinkNo = nil
        var has_func_sym bool
        has_func_sym = self.Has__func__Symbol.Has(Ast_TypeInfo2Stem(typeInfo))
        if has_func_sym{
            self.Has__func__Symbol.Del(Ast_TypeInfo2Stem(typeInfo))
        }
        resultMap.Set(funcBlockInfo,NewTransUnit_FuncBlockResult(_env, funcBlockInfo, workBody, has_func_sym, self.FP.Get_curNsInfo(_env).FP.Get_stmtNum(_env)))
    }
    self.FP.PopScope(_env)
    self.Parser = bakParser
    return resultMap
}
// 27: decl @lune.@base.@TransUnit.TransUnit.MultiTo1
func (self *TransUnit_TransUnit) MultiTo1(_env *LnsEnv, exp *Nodes_Node) *Nodes_Node {
    if exp.FP.Get_expTypeList(_env).Len() > 1{
        exp = &Nodes_ExpMultiTo1Node_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, exp.FP.Get_macroArgFlag(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp.FP.Get_expType(_env))}), exp).Nodes_Node
    } else { 
        {
            _dddType := Ast_DDDTypeInfoDownCastF(exp.FP.Get_expType(_env).FP)
            if !Lns_IsNil( _dddType ) {
                dddType := _dddType.(*Ast_DDDTypeInfo)
                exp = &Nodes_ExpMultiTo1Node_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, exp.FP.Get_macroArgFlag(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddType.FP.Get_typeInfo(_env).FP.Get_nilableTypeInfo(_env))}), exp).Nodes_Node
            }
        }
    }
    return exp
}
// 47: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOneRVal
func (self *TransUnit_TransUnit) analyzeExpOneRVal(_env *LnsEnv, allowNoneType bool,skipOp2Flag bool,opLevel LnsAny,expectType LnsAny) *Nodes_Node {
    var exp *Nodes_Node
    exp = self.FP.analyzeExp(_env, allowNoneType, skipOp2Flag, false, opLevel, expectType)
    exp = self.FP.MultiTo1(_env, exp)
    if expectType != nil{
        expectType_612 := expectType.(*Ast_TypeInfo)
        if _switch0 := expectType_612.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__IF || _switch0 == Ast_TypeInfoKind__Class {
            var expOrgType *Ast_TypeInfo
            expOrgType = exp.FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env)
            var exceptOrgType *Ast_TypeInfo
            exceptOrgType = expectType_612.FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env)
            if expOrgType.FP.IsInheritFrom(_env, self.ProcessInfo, exceptOrgType, nil){
                exp = &Nodes_ExpCastNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(expectType_612)}), exp, expectType_612, nil, "", Nodes_CastKind__Implicit).Nodes_Node
            }
        }
    }
    return exp
}
// 78: decl @lune.@base.@TransUnit.TransUnit.createExpList
func (self *TransUnit_TransUnit) createExpList(_env *LnsEnv, pos Types_Position,expTypeList *LnsList,expList *LnsList,followOn bool,abbrNode LnsAny) *Nodes_ExpListNode {
    var workList *LnsList
    workList = NewLnsList([]LnsAny{})
    var mRetExp LnsAny
    mRetExp = nil
    if expList.Len() > 0{
        for _index, _exp := range( expList.Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( expTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeMultiExp) &&
                _env.SetStackVal( Nodes_hasMultiValNode(_env, exp)) ).(bool)){
                if index != expList.Len(){
                    var firstType *Ast_TypeInfo
                    {
                        _dddType := Ast_DDDTypeInfoDownCastF(exp.FP.Get_expType(_env).FP)
                        if !Lns_IsNil( _dddType ) {
                            dddType := _dddType.(*Ast_DDDTypeInfo)
                            firstType = dddType.FP.Get_typeInfo(_env).FP.Get_nilableTypeInfo(_env)
                        } else {
                            firstType = exp.FP.Get_expType(_env)
                        }
                    }
                    workList.Insert(Nodes_ExpMultiTo1Node2Stem(Nodes_ExpMultiTo1Node_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, exp.FP.Get_macroArgFlag(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(firstType)}), exp)))
                } else { 
                    mRetExp = NewNodes_MRetExp(_env, exp, index)
                    for _listIndex, _expType := range( exp.FP.Get_expTypeList(_env).Items ) {
                        listIndex := _listIndex + 1
                        expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if listIndex != 1{
                            expTypeList.Insert(Ast_TypeInfo2Stem(expType))
                        }
                    }
                    for _retIndex, _retType := range( exp.FP.Get_expTypeList(_env).Items ) {
                        retIndex := _retIndex + 1
                        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if retIndex == 1{
                            workList.Insert(Nodes_ExpMRetNode2Stem(Nodes_ExpMRetNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, exp.FP.Get_macroArgFlag(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(retType)}), exp)))
                        } else { 
                            workList.Insert(Nodes_ExpAccessMRetNode2Stem(Nodes_ExpAccessMRetNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, exp.FP.Get_macroArgFlag(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(retType)}), exp, retIndex)))
                        }
                    }
                }
            } else { 
                workList.Insert(Nodes_Node2Stem(exp))
            }
        }
    }
    if abbrNode != nil{
        abbrNode_642 := abbrNode.(*Nodes_AbbrNode)
        workList.Insert(Nodes_AbbrNode2Stem(abbrNode_642))
    }
    return Nodes_ExpListNode_create(_env, self.NodeManager, pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), expTypeList, workList, mRetExp, followOn)
}
// 157: decl @lune.@base.@TransUnit.TransUnit.analyzeExpList
func (self *TransUnit_TransUnit) analyzeExpList(_env *LnsEnv, allowNoneType bool,skipOp2Flag bool,canLeftExp bool,expNode LnsAny,expectTypeList LnsAny,contExpect LnsAny) *Nodes_ExpListNode {
    var expList *LnsList
    expList = NewLnsList([]LnsAny{})
    var pos LnsAny
    pos = nil
    var expTypeList *LnsList
    expTypeList = NewLnsList([]LnsAny{})
    if expNode != nil{
        expNode_648 := expNode.(*Nodes_Node)
        pos = expNode_648.FP.Get_pos(_env)
        expList.Insert(Nodes_Node2Stem(expNode_648))
        expTypeList.Insert(Ast_TypeInfo2Stem(expNode_648.FP.Get_expType(_env)))
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
            expectTypeList_656 := expectTypeList.(*LnsList)
            if expectTypeList_656.Len() > 0{
                var checkIndex LnsInt
                checkIndex = index
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( index > expectTypeList_656.Len()) &&
                    _env.SetStackVal( contExpect) )){
                    checkIndex = expectTypeList_656.Len()
                }
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( checkIndex <= expectTypeList_656.Len()) &&
                    _env.SetStackVal( expectTypeList_656.GetAt(checkIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeNone) ).(bool)){
                    var worktype *Ast_TypeInfo
                    worktype = expectTypeList_656.GetAt(checkIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    expectType = worktype
                    if worktype == Ast_builtinTypeExp{
                        allowNoneTypeOne = true
                    }
                }
            }
        }
        var exp *Nodes_Node
        if self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.Get_mode(_env) == Nodes_MacroMode__AnalyzeArg{
            if _switch0 := expectType; _switch0 == Ast_builtinTypeExp || _switch0 == Ast_builtinTypeMultiExp {
                self.MacroCtrl.FP.SwitchMacroMode(_env)
                exp = self.FP.analyzeExp(_env, allowNoneTypeOne, skipOp2Flag, canLeftExp, 0, expectType)
                self.MacroCtrl.FP.RestoreMacroMode(_env)
            } else if _switch0 == Ast_builtinTypeBlockArg {
                exp = &self.FP.analyzeBlock(_env, Nodes_BlockKind__Macro, TransUnit_TentativeMode__Ignore, nil, nil).Nodes_Node
            } else {
                exp = self.FP.analyzeExp(_env, allowNoneTypeOne, skipOp2Flag, canLeftExp, 0, expectType)
            }
        } else { 
            exp = self.FP.analyzeExp(_env, allowNoneTypeOne, skipOp2Flag, canLeftExp, 0, expectType)
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(allowNoneTypeOne)) &&
            _env.SetStackVal( expectType != Ast_builtinTypeBlockArg) &&
            _env.SetStackVal( Lns_op_not(exp.FP.CanBeRight(_env, self.ProcessInfo))) ).(bool)){
            self.FP.AddErrMess(_env, exp.FP.Get_effectivePos(_env), _env.GetVM().String_format("This arg can't be r-value. -- %s", []LnsAny{Nodes_getNodeKindName(_env, exp.FP.Get_kind(_env))}))
        }
        if Lns_op_not(pos){
            pos = exp.FP.Get_pos(_env)
        }
        if _switch1 := expectType; _switch1 == Ast_builtinTypeExp || _switch1 == Ast_builtinTypeMultiExp || _switch1 == Ast_builtinTypeBlockArg {
            exp = &Nodes_ExpMacroArgExpNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), exp.FP.Get_expTypeList(_env), Macro_nodeToCodeTxt(_env, exp, self.moduleType)).Nodes_Node
            expTypeList.Insert(Ast_TypeInfo2Stem(Lns_unwrap( expectType).(*Ast_TypeInfo)))
        } else {
            expTypeList.Insert(Ast_TypeInfo2Stem(exp.FP.Get_expType(_env)))
        }
        expList.Insert(Nodes_Node2Stem(exp))
        var token *Types_Token
        token = self.FP.getToken(_env, true)
        if token.Txt == "**"{
            if Lns_op_not(Nodes_hasMultiValNode(_env, exp)){
                self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("This arg(%d) doesn't have multiple value. It must not use '**'", []LnsAny{index}))
            }
            followOn = true
            token = self.FP.getToken(_env, true)
        }
        if token.Txt == "##"{
            if exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                self.FP.AddErrMess(_env, token.Pos, "'##' can't use with '...'")
            }
            abbrNode = Nodes_AbbrNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeAbbr)}))
            self.FP.getToken(_env, nil)
            break
        }
        index = index + 1
        if self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.EqualsArgTypeList(_env, expectTypeList){
            self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.NextArg(_env)
        }
        if token.Txt != ","{ break }
    }
    var expListNode *Nodes_ExpListNode
    expListNode = self.FP.createExpList(_env, Lns_unwrapDefault( pos, self.FP.CreatePosition(_env, 0, 0)).(Types_Position), expTypeList, expList, followOn, abbrNode)
    if Lns_op_not(allowNoneType){
        for _expIndex, _expType := range( expTypeList.Items ) {
            expIndex := _expIndex + 1
            expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if expType == Ast_builtinTypeNone{
                self.FP.AddErrMess(_env, Lns_unwrapDefault( pos, self.FP.CreatePosition(_env, 0, 0)).(Types_Position), _env.GetVM().String_format("The type of exp(%d) is None!!", []LnsAny{expIndex}))
            }
        }
    }
    self.FP.Pushback(_env)
    return expListNode
}
// 295: decl @lune.@base.@TransUnit.TransUnit.checkSymbolHavingValue
func (self *TransUnit_TransUnit) checkSymbolHavingValue(_env *LnsEnv, pos Types_Position,symbolInfoList *LnsList) {
    for _, _symbolInfo := range( symbolInfoList.Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Var{
            if Lns_op_not(symbolInfo.FP.Get_hasValueFlag(_env)){
                self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("this variable have no value. -- %s", []LnsAny{symbolInfo.FP.Get_name(_env)}))
            }
        }
    }
}
// 309: decl @lune.@base.@TransUnit.TransUnit.analyzeExpRefItem
func (self *TransUnit_TransUnit) analyzeExpRefItem(_env *LnsEnv, token *Types_Token,exp *Nodes_Node,nilAccess bool) *Nodes_Node {
    self.FP.checkSymbolHavingValue(_env, exp.FP.Get_pos(_env), exp.FP.GetSymbolInfo(_env))
    var expType *Ast_TypeInfo
    expType = exp.FP.Get_expType(_env)
    expType = expType.FP.Get_extedType(_env)
    if nilAccess{
        if Lns_op_not(exp.FP.Get_expType(_env).FP.Get_nilable(_env)){
            self.FP.addWarnMess(_env, token.Pos, _env.GetVM().String_format("This is not nilable. -- %s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil)}))
            nilAccess = false
        } else { 
            expType = expType.FP.Get_nonnilableType(_env)
        }
    } else if exp.FP.Get_expType(_env).FP.Get_nilable(_env){
        self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("could not access with []. Use '$[]'-- %s", []LnsAny{exp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
    }
    var expectItemType LnsAny
    expectItemType = nil
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_builtinTypeStem_
    var indexTypeInfo *Ast_TypeInfo
    indexTypeInfo = Ast_builtinTypeInt
    if expType.FP.Get_kind(_env) == Ast_TypeInfoKind__Map{
        var itemTypeList *LnsList
        itemTypeList = expType.FP.Get_itemTypeInfoList(_env)
        typeInfo = itemTypeList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        indexTypeInfo = itemTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        expectItemType = itemTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(typeInfo.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeStem_, nil, nil))) &&
            _env.SetStackVal( Lns_op_not(typeInfo.FP.Get_nilable(_env))) ).(bool)){
            typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
        }
    } else if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( expType.FP.Get_kind(_env) == Ast_TypeInfoKind__Array) ||
        _env.SetStackVal( expType.FP.Get_kind(_env) == Ast_TypeInfoKind__List) ).(bool){
        typeInfo = expType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
    } else if expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeString, nil, nil){
        typeInfo = Ast_builtinTypeInt
    } else if expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeStem, nil, nil){
        indexTypeInfo = Ast_builtinTypeStem
        typeInfo = Ast_builtinTypeStem_
    } else { 
        self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("could not access with []. -- %s", []LnsAny{exp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
    }
    if nilAccess{
        self.HelperInfo.UseNilAccess = true
        if Lns_op_not(typeInfo.FP.Get_nilable(_env)){
            typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
        }
    }
    if Ast_TypeInfo_isMut(_env, typeInfo){
        if expType.FP.Get_mutMode(_env) == Ast_MutMode__IMutRe{
            typeInfo = self.FP.createModifier(_env, typeInfo, Ast_MutMode__IMutRe)
        }
    }
    if Ast_isExtType(_env, exp.FP.Get_expType(_env).FP.Get_nonnilableType(_env)){
        typeInfo = self.FP.createExtType(_env, exp.FP.Get_pos(_env), typeInfo)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.IsAnalyzingBlockArg(_env))) &&
            _env.SetStackVal( Lns_op_not(self.FP.Get_curNsInfo(_env).FP.CanAccessLuaval(_env))) ).(bool)){
            self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("can't access Luaval without __luago. -- %s", []LnsAny{exp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    var indexExp *Nodes_Node
    indexExp = self.FP.analyzeExpOneRVal(_env, false, false, nil, expectItemType)
    if Lns_op_not(indexExp.FP.CanBeRight(_env, self.ProcessInfo)){
        self.FP.AddErrMess(_env, indexExp.FP.Get_pos(_env), "This node can't use index")
    }
    if Lns_op_not(Lns_car(indexTypeInfo.FP.CanEvalWith(_env, self.ProcessInfo, indexExp.FP.Get_expType(_env), Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)){
        self.FP.AddErrMess(_env, indexExp.FP.Get_pos(_env), _env.GetVM().String_format("unmatch index type -- %s, %s", []LnsAny{indexTypeInfo.FP.GetTxt(_env, nil, nil, nil), indexExp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
    }
    self.FP.checkNextToken(_env, "]")
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( expType.FP.Get_kind(_env) == Ast_TypeInfoKind__Array) ||
        _env.SetStackVal( expType.FP.Get_kind(_env) == Ast_TypeInfoKind__List) ).(bool){
        {
            _indexLit := TransUnit_convExp3_10766(Lns_2DDD(indexExp.FP.GetLiteral(_env)))
            if !Lns_IsNil( _indexLit ) {
                indexLit := _indexLit
                switch _matchExp0 := indexLit.(type) {
                case *Nodes_Literal__Int:
                val := _matchExp0.Val1
                    if val <= 0{
                        self.FP.addWarnMess(_env, indexExp.FP.Get_pos(_env), _env.GetVM().String_format("index <= -1 (%d)", []LnsAny{val}))
                    }
                }
            }
        }
    }
    return &Nodes_ExpRefItemNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), exp, nilAccess, nil, indexExp).Nodes_Node
}
// 438: decl @lune.@base.@TransUnit.TransUnit.checkImplicitCast
func (self *TransUnit_TransUnit) checkImplicitCast(_env *LnsEnv, alt2typeMap *LnsMap,validCastToGen bool,dstTypeList *LnsList,expListNode *Nodes_ExpListNode,callback TransUnit_checkImplicitCastCallback_10_) LnsAny {
    var expNodeList *LnsList
    expNodeList = expListNode.FP.Get_expList(_env)
    var newExpNodeList *LnsList
    newExpNodeList = NewLnsList([]LnsAny{})
    var expTypeList *LnsList
    expTypeList = NewLnsList([]LnsAny{})
    var hasModNode bool
    hasModNode = false
    var process func(_env *LnsEnv, index LnsInt,dstType *Ast_TypeInfo,expNode *Nodes_Node,workNode *Nodes_Node)(*Nodes_Node, bool)
    process = func(_env *LnsEnv, index LnsInt,dstType *Ast_TypeInfo,expNode *Nodes_Node,workNode *Nodes_Node)(*Nodes_Node, bool) {
        {
            _repNode := callback(_env, dstType, expNode)
            if !Lns_IsNil( _repNode ) {
                repNode := _repNode.(*Nodes_Node)
                if Lns_op_not(hasModNode){
                    hasModNode = true
                }
                workNode = repNode
            } else {
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( expNode.FP.Get_expType(_env) != Ast_builtinTypeNil) &&
                    _env.SetStackVal( dstType != Ast_builtinTypeEmpty) &&
                    _env.SetStackVal( Lns_op_not(dstType.FP.Get_srcTypeInfo(_env).FP.Equals(_env, self.ProcessInfo, expNode.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env), nil, nil))) &&
                    _env.SetStackVal( Lns_op_not(dstType.FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env).FP.Equals(_env, self.ProcessInfo, expNode.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env), nil, nil))) ).(bool)){
                    if expNode.FP.Get_kind(_env) != Nodes_NodeKind_get_Abbr(_env){
                        if Lns_op_not(hasModNode){
                            hasModNode = true
                        }
                        if dstType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                            var argList *LnsList
                            argList = NewLnsList([]LnsAny{})
                            var argTypeList *LnsList
                            argTypeList = NewLnsList([]LnsAny{})
                            {
                                var _forFrom0 LnsInt = index
                                var _forTo0 LnsInt = expNodeList.Len()
                                for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                                    workIndex := _forWork0
                                    var appNode *Nodes_Node
                                    appNode = expNodeList.GetAt(workIndex).(Nodes_NodeDownCast).ToNodes_Node()
                                    argList.Insert(Nodes_Node2Stem(appNode))
                                    if workIndex == expNodeList.Len(){
                                        for _, _expType := range( appNode.FP.Get_expTypeList(_env).Items ) {
                                            expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                                            argTypeList.Insert(Ast_TypeInfo2Stem(expType))
                                        }
                                    } else { 
                                        argTypeList.Insert(Ast_TypeInfo2Stem(appNode.FP.Get_expType(_env)))
                                    }
                                }
                            }
                            var mRetExp LnsAny
                            {
                                _workMRetExp := expListNode.FP.Get_mRetExp(_env)
                                if !Lns_IsNil( _workMRetExp ) {
                                    workMRetExp := _workMRetExp.(*Nodes_MRetExp)
                                    mRetExp = NewNodes_MRetExp(_env, workMRetExp.FP.Get_exp(_env), workMRetExp.FP.Get_index(_env) - index + 1)
                                    if argList.Len() == 0{
                                        return &Nodes_ExpSubDDDNode_create(_env, self.NodeManager, expNode.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), workMRetExp.FP.Get_exp(_env).FP.Get_expTypeList(_env), workMRetExp.FP.Get_exp(_env), index - workMRetExp.FP.Get_index(_env)).Nodes_Node, true
                                    }
                                } else {
                                    mRetExp = nil
                                }
                            }
                            var newExpListNode *Nodes_ExpListNode
                            newExpListNode = Nodes_ExpListNode_create(_env, self.NodeManager, expNode.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), argTypeList, argList, mRetExp, expListNode.FP.Get_followOn(_env))
                            workNode = &Nodes_ExpToDDDNode_create(_env, self.NodeManager, expNode.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dstType)}), newExpListNode).Nodes_Node
                            return workNode, true
                        } else { 
                            var castType *Ast_TypeInfo
                            if validCastToGen{
                                castType = Lns_unwrapDefault( alt2typeMap.Get(dstType), dstType).(*Ast_TypeInfo)
                            } else { 
                                castType = dstType
                            }
                            workNode = &Nodes_ExpCastNode_create(_env, self.NodeManager, expNode.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(expNode.FP.Get_expType(_env))}), expNode, castType, nil, "", Nodes_CastKind__Implicit).Nodes_Node
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
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( index == expNodeList.Len()) &&
                _env.SetStackVal( expNode.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
                {
                    var _forFrom0 LnsInt = index
                    var _forTo0 LnsInt = dstTypeList.Len()
                    for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                        dstIndex := _forWork0
                        workNode = expNode
                        workNode, stopFlag = process(_env, dstIndex, dstTypeList.GetAt(dstIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), expNode, workNode)
                        newExpNodeList.Insert(Nodes_Node2Stem(workNode))
                        expTypeList.Insert(Ast_TypeInfo2Stem(workNode.FP.Get_expType(_env)))
                    }
                }
                break
            } else { 
                workNode, stopFlag = process(_env, index, dstTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), expNode, workNode)
            }
        }
        newExpNodeList.Insert(Nodes_Node2Stem(workNode))
        expTypeList.Insert(Ast_TypeInfo2Stem(workNode.FP.Get_expType(_env)))
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
        _mRetExp := expListNode.FP.Get_mRetExp(_env)
        if !Lns_IsNil( _mRetExp ) {
            mRetExp := _mRetExp.(*Nodes_MRetExp)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( mRetExp.FP.Get_index(_env) <= dstTypeList.Len()) &&
                _env.SetStackVal( dstTypeList.GetAt(mRetExp.FP.Get_index(_env)).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind(_env) != Ast_TypeInfoKind__DDD) ).(bool)){
                newMRetExp = mRetExp
            }
        }
    }
    return Nodes_ExpListNode_create(_env, self.NodeManager, expListNode.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), expTypeList, newExpNodeList, newMRetExp, expListNode.FP.Get_followOn(_env))
}
// 605: decl @lune.@base.@TransUnit.TransUnit.checkMatchType
func (self *TransUnit_TransUnit) checkMatchType(_env *LnsEnv, message string,pos Types_Position,dstTypeList *LnsList,expListNode LnsAny,allowDstShort bool,warnForFollow bool,workAlt2typeMap LnsAny,validImplicitCast bool)(LnsInt, *LnsMap, LnsAny, *LnsList) {
    var expNodeList *LnsList
    
    {
        _expNodeList := _env.NilAccFin(_env.NilAccPush(expListNode) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_ExpListNode).FP.Get_expList(_env)}))
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
        if expNodeList.GetAt(expNodeList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_kind(_env) == Nodes_NodeKind_get_Abbr(_env){
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
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( realExpNum == -1) &&
            _env.SetStackVal( expNode.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpAccessMRet(_env)) ).(bool)){
            realExpNum = index - 1
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( index == workExpNodeList.Len()) &&
            _env.SetStackVal( Nodes_ExpMacroArgExpNode2Stem(Nodes_ExpMacroArgExpNodeDownCastF(expNode.FP)) == nil) ).(bool)){
            for _, _expType := range( expNode.FP.Get_expTypeList(_env).Items ) {
                expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                expTypeList.Insert(Ast_TypeInfo2Stem(expType))
            }
        } else { 
            expTypeList.Insert(Ast_TypeInfo2Stem(expNode.FP.Get_expType(_env)))
        }
    }
    if realExpNum == -1{
        realExpNum = workExpNodeList.Len()
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( warnForFollow) &&
        _env.SetStackVal( expTypeList.Len() > realExpNum) ).(bool)){
        warnForFollowSrcIndex = realExpNum + 1
    }
    if hasAbbr{
        expTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeAbbr))
    }
    var alt2typeMap *LnsMap
    if workAlt2typeMap != nil{
        workAlt2typeMap_82 := workAlt2typeMap.(*LnsMap)
        alt2typeMap = workAlt2typeMap_82
    } else {
        alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false)
    }
    Ast_CanEvalCtrlTypeInfo_setupNeedAutoBoxing(_env, alt2typeMap, self.ProcessInfo)
    var result LnsInt
    var mess string
    result,mess = Ast_TypeInfo_checkMatchType(_env, self.ProcessInfo, dstTypeList, expTypeList, allowDstShort, warnForFollowSrcIndex, alt2typeMap)
    if _switch0 := result; _switch0 == Ast_MatchType__Error {
        self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("%s: %s", []LnsAny{message, mess}))
    } else if _switch0 == Ast_MatchType__Warn {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(self.Ctrl_info.CheckingDefineAbbr)) &&
            _env.SetStackVal( Code_isMessageOf(_env, Code_ID__nothing_define_abbr, mess)) ).(bool)){
        } else { 
            self.FP.addWarnMess(_env, pos, _env.GetVM().String_format("%s: %s", []LnsAny{message, mess}))
        }
    }
    if expListNode != nil{
        expListNode_91 := expListNode.(*Nodes_ExpListNode)
        var autoBoxingCount LnsInt
        autoBoxingCount = 0
        var hasImplictCast bool
        hasImplictCast = false
        var newExpListNode LnsAny
        if result != Ast_MatchType__Error{
            if validImplicitCast{
                {
                    _workList := self.FP.checkImplicitCast(_env, alt2typeMap, false, dstTypeList, expListNode_91, TransUnit_checkImplicitCastCallback_10_(func(_env *LnsEnv, dstType *Ast_TypeInfo,expNode *Nodes_Node) LnsAny {
                        if Ast_CanEvalCtrlTypeInfo_canAutoBoxing(_env, dstType, expNode.FP.Get_expType(_env)){
                            autoBoxingCount = autoBoxingCount + 1
                            return &Nodes_BoxingNode_create(_env, self.NodeManager, expNode.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dstType)}), expNode).Nodes_Node
                        }
                        return nil
                    }))
                    if !Lns_IsNil( _workList ) {
                        workList := _workList.(*Nodes_ExpListNode)
                        newExpListNode = workList
                        hasImplictCast = true
                    } else {
                        newExpListNode = nil
                    }
                }
            } else { 
                newExpListNode = Nodes_ExpListNode_create(_env, self.NodeManager, expListNode_91.FP.Get_pos(_env), expListNode_91.FP.Get_inTestBlock(_env), expListNode_91.FP.Get_macroArgFlag(_env), expListNode_91.FP.Get_expTypeList(_env), expListNode_91.FP.Get_expList(_env), expListNode_91.FP.Get_mRetExp(_env), expListNode_91.FP.Get_followOn(_env))
                hasImplictCast = false
            }
        } else { 
            newExpListNode = nil
        }
        if autoBoxingCount > 0{
            if Lns_op_not(Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(_env, alt2typeMap, autoBoxingCount)){
                self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("illegal auto boxing error -- %d", []LnsAny{autoBoxingCount}))
            }
            return result, alt2typeMap, newExpListNode, expTypeList
        } else if Ast_CanEvalCtrlTypeInfo_hasNeedAutoBoxing(_env, alt2typeMap){
            self.FP.AddErrMess(_env, pos, "not support auto boxing")
        }
        if hasImplictCast{
            return result, alt2typeMap, newExpListNode, expTypeList
        }
    }
    if Lns_op_not(Ast_CanEvalCtrlTypeInfo_finishNeedAutoBoxing(_env, alt2typeMap, 0)){
        self.FP.AddErrMess(_env, pos, "can't auto boxing error")
    }
    return result, alt2typeMap, nil, expTypeList
}
// 779: decl @lune.@base.@TransUnit.TransUnit.checkMatchValType
func (self *TransUnit_TransUnit) checkMatchValType(_env *LnsEnv, pos Types_Position,funcTypeInfo *Ast_TypeInfo,expList LnsAny,genericTypeList *LnsList,genericsClass LnsAny)(LnsInt, *LnsMap, LnsAny) {
    var argTypeList *LnsList
    argTypeList = funcTypeInfo.FP.Get_argTypeInfoList(_env)
    if funcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        var extTypeList LnsAny
        var mess string
        extTypeList,mess = Ast_convToExtTypeList(_env, self.ProcessInfo, argTypeList)
        if extTypeList != nil{
            extTypeList_118 := extTypeList.(*LnsList)
            argTypeList = extTypeList_118
        } else {
            self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("not support argType on Luaval -- %s", []LnsAny{mess}))
        }
    }
    var validImplicitCast bool
    validImplicitCast = true
    if _switch0 := funcTypeInfo; _switch0 == self.builtinFunc.List_insert || _switch0 == self.builtinFunc.Set_add || _switch0 == self.builtinFunc.Set_del {
    } else if _switch0 == self.builtinFunc.List_sort {
        _ = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false)
        var callback *Ast_NormalTypeInfo
        callback = self.ProcessInfo.FP.CreateFuncAsync(_env, false, false, nil, Ast_TypeInfoKind__FormFunc, self.ProcessInfo.FP.Get_dummyParentType(_env), self.ProcessInfo.FP.Get_dummyParentType(_env).FP, false, false, true, Ast_AccessMode__Pri, "sort", Ast_Async__Async, nil, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()), Ast_TypeInfo2Stem(genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeBool)}), Ast_MutMode__IMut)
        argTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(callback.FP.Get_nilableTypeInfo(_env))})
        validImplicitCast = false
    } else if _switch0 == self.builtinFunc.List_remove {
    } else if _switch0 == self.builtinFunc.Lns___run {
        self.HelperInfo.UseRun = true
    }
    var warnForFollow bool
    warnForFollow = true
    if expList != nil{
        expList_129 := expList.(*Nodes_ExpListNode)
        if expList_129.FP.Get_followOn(_env){
            warnForFollow = false
        }
    }
    var alt2typeMap *LnsMap
    if funcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
        if genericsClass != nil{
            genericsClass_134 := genericsClass.(*Ast_TypeInfo)
            if funcTypeInfo.FP.Get_rawTxt(_env) == "__init"{
                alt2typeMap = genericsClass_134.FP.CreateAlt2typeMap(_env, true)
            } else { 
                if funcTypeInfo.FP.Get_itemTypeInfoList(_env).Len() == 0{
                    alt2typeMap = genericsClass_134.FP.CreateAlt2typeMap(_env, false)
                } else { 
                    alt2typeMap = genericsClass_134.FP.CreateAlt2typeMap(_env, true)
                    for _, _itemType := range( genericsClass_134.FP.Get_itemTypeInfoList(_env).Items ) {
                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( itemType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                            _env.SetStackVal( Lns_op_not(alt2typeMap.Get(itemType))) ).(bool)){
                            alt2typeMap.Set(itemType,Ast_builtinTypeNone)
                        }
                    }
                }
            }
        } else {
            self.FP.Error(_env, "none class")
        }
    } else { 
        alt2typeMap = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, funcTypeInfo.FP.Get_itemTypeInfoList(_env).Len() > 0)
    }
    var matchResult LnsInt
    var newExpNodeList LnsAny
    matchResult,_,newExpNodeList = TransUnit_convExp4_1837(Lns_2DDD(self.FP.checkMatchType(_env, funcTypeInfo.FP.GetTxt(_env, nil, nil, nil), pos, argTypeList, expList, false, warnForFollow, alt2typeMap, validImplicitCast)))
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( expList) &&
        _env.SetStackVal( newExpNodeList) )){
        return matchResult, alt2typeMap, newExpNodeList
    }
    return matchResult, alt2typeMap, expList
}
// 866: decl @lune.@base.@TransUnit.TransUnit.analyzeListItems
func (self *TransUnit_TransUnit) analyzeListItems(_env *LnsEnv, firstPos Types_Position,nextToken *Types_Token,termTxt string,expectTypeList LnsAny)(LnsAny, *Ast_TypeInfo) {
    var expList LnsAny
    expList = nil
    var itemCommonType LnsAny
    itemCommonType = &Ast_CommonType__Normal{Ast_builtinTypeNone}
    if nextToken.Txt != termTxt{
        self.FP.Pushback(_env)
        expList = self.FP.analyzeExpList(_env, false, false, false, nil, expectTypeList, expectTypeList != nil)
        self.FP.checkNextToken(_env, termTxt)
        var nodeList *LnsList
        nodeList = (Lns_unwrap( expList).(*Nodes_ExpListNode)).FP.Get_expList(_env)
        for _, _exp := range( nodeList.Items ) {
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            itemCommonType = Ast_TypeInfo_getCommonTypeCombo(_env, self.ProcessInfo, itemCommonType, &Ast_CommonType__Normal{exp.FP.Get_expType(_env)}, Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false))
        }
    }
    var itemTypeInfo *Ast_TypeInfo
    switch _matchExp0 := itemCommonType.(type) {
    case *Ast_CommonType__Normal:
    info := _matchExp0.Val1
        itemTypeInfo = info
    case *Ast_CommonType__Combine:
    info := _matchExp0.Val1
        itemTypeInfo = info.FP.Get_typeInfo(_env, self.ProcessInfo)
    }
    if itemTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        if itemTypeInfo.FP.Get_itemTypeInfoList(_env).Len() > 0{
            itemTypeInfo = itemTypeInfo.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        } else { 
            itemTypeInfo = Ast_builtinTypeStem_
        }
    }
    if Lns_op_not(expectTypeList){
        var expTypeList *LnsList
        expTypeList = NewLnsList([]LnsAny{})
        if expList != nil{
            expList_166 := expList.(*Nodes_ExpListNode)
            for _index, _expNode := range( expList_166.FP.Get_expList(_env).Items ) {
                index := _index + 1
                expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                if index == expList_166.FP.Get_expList(_env).Len(){
                    if expNode.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                        expTypeList.Insert(Ast_TypeInfo2Stem(expNode.FP.Get_expType(_env)))
                    } else { 
                        {
                            var _forFrom0 LnsInt = 1
                            var _forTo0 LnsInt = expNode.FP.Get_expTypeList(_env).Len()
                            for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
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
        _,_,workExpList = TransUnit_convExp4_2161(Lns_2DDD(self.FP.checkMatchType(_env, "List constructor", firstPos, expTypeList, expList, false, false, nil, true)))
        if workExpList != nil{
            workExpList_180 := workExpList.(*Nodes_ExpListNode)
            expList = workExpList_180
        }
    }
    return expList, itemTypeInfo
}
// 935: decl @lune.@base.@TransUnit.TransUnit.analyzeListConst
func (self *TransUnit_TransUnit) analyzeListConst(_env *LnsEnv, token *Types_Token,expectType LnsAny) *Nodes_Node {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var expectTypeList LnsAny
    expectTypeList = nil
    if _env.NilAccFin(_env.NilAccPush(expectType) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) == Ast_TypeInfoKind__List{
        {
            _itemTypeInfoList := _env.NilAccFin(_env.NilAccPush(expectType) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_itemTypeInfoList(_env)}))
            if !Lns_IsNil( _itemTypeInfoList ) {
                itemTypeInfoList := _itemTypeInfoList.(*LnsList)
                expectTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
            }
        }
    }
    var expList LnsAny
    var itemTypeInfo *Ast_TypeInfo
    expList,itemTypeInfo = self.FP.analyzeListItems(_env, token.Pos, nextToken, "]", expectTypeList)
    var typeInfoList *LnsList
    typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})
    if token.Txt == "["{
        typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.ProcessInfo.FP.CreateList(_env, Ast_AccessMode__Local, self.FP.getCurrentClass(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfo)}), Ast_MutMode__Mut))})
        return &Nodes_LiteralListNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), typeInfoList, expList).Nodes_Node
    } else { 
        typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.ProcessInfo.FP.CreateArray(_env, Ast_AccessMode__Local, self.FP.getCurrentClass(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfo)}), Ast_MutMode__Mut))})
        return &Nodes_LiteralArrayNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), typeInfoList, expList).Nodes_Node
    }
// insert a dummy
    return nil
}
// 968: decl @lune.@base.@TransUnit.TransUnit.analyzeSetConst
func (self *TransUnit_TransUnit) analyzeSetConst(_env *LnsEnv, token *Types_Token,expectType LnsAny) *Nodes_Node {
    self.HelperInfo.UseSet = true
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var expectTypeList LnsAny
    expectTypeList = nil
    if _env.NilAccFin(_env.NilAccPush(expectType) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) == Ast_TypeInfoKind__Set{
        {
            _itemTypeInfoList := _env.NilAccFin(_env.NilAccPush(expectType) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_itemTypeInfoList(_env)}))
            if !Lns_IsNil( _itemTypeInfoList ) {
                itemTypeInfoList := _itemTypeInfoList.(*LnsList)
                expectTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
            }
        }
    }
    var expList LnsAny
    var itemTypeInfo *Ast_TypeInfo
    expList,itemTypeInfo = self.FP.analyzeListItems(_env, token.Pos, nextToken, ")", expectTypeList)
    if itemTypeInfo.FP.Get_nilable(_env){
        if expList != nil{
            expList_202 := expList.(*Nodes_ExpListNode)
            for _, _exp := range( expList_202.FP.Get_expList(_env).Items ) {
                exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
                var expType *Ast_TypeInfo
                expType = exp.FP.Get_expType(_env)
                if expType.FP.Get_nilable(_env){
                    self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("'Set' object can't store nilable. -- %s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        }
    }
    var typeInfoList *LnsList
    typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})
    typeInfoList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(self.ProcessInfo.FP.CreateSet(_env, Ast_AccessMode__Local, self.FP.getCurrentClass(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(itemTypeInfo)}), Ast_MutMode__Mut))})
    return &Nodes_LiteralSetNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), typeInfoList, expList).Nodes_Node
}
// 1008: decl @lune.@base.@TransUnit.TransUnit.analyzeMapConst
func (self *TransUnit_TransUnit) analyzeMapConst(_env *LnsEnv, token *Types_Token,expectType LnsAny) *Nodes_LiteralMapNode {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var _map *LnsMap
    _map = NewLnsMap( map[LnsAny]LnsAny{})
    var pairList *LnsList
    pairList = NewLnsList([]LnsAny{})
    var keyTypeInfo *Ast_TypeInfo
    keyTypeInfo = Ast_builtinTypeNone
    var valTypeInfo *Ast_TypeInfo
    valTypeInfo = Ast_builtinTypeNone
    var getMapKeyValType func(_env *LnsEnv, pos Types_Position,keyFlag bool,typeInfo *Ast_TypeInfo,expType *Ast_TypeInfo) *Ast_TypeInfo
    getMapKeyValType = func(_env *LnsEnv, pos Types_Position,keyFlag bool,typeInfo *Ast_TypeInfo,expType *Ast_TypeInfo) *Ast_TypeInfo {
        if expType.FP.Get_nilable(_env){
            if keyFlag{
                self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("map key can't set a nilable -- %s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil)}))
            }
            if expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeNil, nil, nil){
                return typeInfo
            }
            expType = expType.FP.Get_nonnilableType(_env)
        }
        return Ast_TypeInfo_getCommonType(_env, self.ProcessInfo, typeInfo, expType, Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false))
    }
    var expectKeyType LnsAny
    expectKeyType = nil
    var expectValType LnsAny
    expectValType = nil
    if _env.NilAccFin(_env.NilAccPush(expectType) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) == Ast_TypeInfoKind__Map{
        {
            _itemTypeInfoList := _env.NilAccFin(_env.NilAccPush(expectType) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_itemTypeInfoList(_env)}))
            if !Lns_IsNil( _itemTypeInfoList ) {
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
        self.FP.Pushback(_env)
        var key *Nodes_Node
        key = self.FP.analyzeExpOneRVal(_env, false, false, nil, expectKeyType)
        keyTypeInfo = getMapKeyValType(_env, key.FP.Get_pos(_env), true, keyTypeInfo, key.FP.Get_expType(_env))
        self.FP.checkNextToken(_env, ":")
        var val *Nodes_Node
        val = self.FP.analyzeExpOneRVal(_env, false, false, nil, expectValType)
        valTypeInfo = getMapKeyValType(_env, val.FP.Get_pos(_env), false, valTypeInfo, val.FP.Get_expType(_env))
        pairList.Insert(Nodes_PairItem2Stem(NewNodes_PairItem(_env, key, val)))
        _map.Set(key,val)
        nextToken = self.FP.getToken(_env, nil)
        if nextToken.Txt != ","{
            break
        }
        nextToken = self.FP.getToken(_env, nil)
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = self.ProcessInfo.FP.CreateMap(_env, Ast_AccessMode__Local, self.FP.getCurrentClass(_env), keyTypeInfo, valTypeInfo, Ast_MutMode__Mut)
    self.FP.checkToken(_env, nextToken, "}")
    return Nodes_LiteralMapNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), _map, pairList)
}
// 1096: decl @lune.@base.@TransUnit.TransUnit.evalMacroOp
func (self *TransUnit_TransUnit) evalMacroOp(_env *LnsEnv, firstToken *Types_Token,macroTypeInfo *Ast_TypeInfo,expList LnsAny,evalMacroCallback Macro_EvalMacroCallback) {
    var parser LnsAny
    var mess LnsAny
    parser,mess = self.MacroCtrl.FP.EvalMacroOp(_env, self.moduleType, self.Parser.FP.GetStreamName(_env), firstToken, macroTypeInfo, expList)
    var bakParser *Parser_DefaultPushbackParser
    bakParser = self.Parser
    if parser != nil{
        parser_240 := parser.(*Parser_Parser)
        self.FP.SetParser(_env, NewParser_DefaultPushbackParser(_env, parser_240))
    } else {
        self.FP.Error(_env, Lns_unwrap( mess).(string))
    }
    self.MacroCtrl.FP.StartExpandMode(_env, firstToken.Pos, macroTypeInfo, evalMacroCallback)
    var nextToken *Types_Token
    nextToken = self.FP.GetTokenNoErr(_env, nil)
    self.FP.SetParser(_env, bakParser)
    if nextToken != Parser_getEofToken(_env){
        self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("remain macro expand-statement token -- '%s'(%d:%d)", []LnsAny{nextToken.Txt, nextToken.Pos.LineNo, nextToken.Pos.Column}))
        if Lns_op_not(macroTypeInfo.FP.Get_externalFlag(_env)){
            self.FP.AddErrMess(_env, nextToken.Pos, _env.GetVM().String_format("remain macro expand-statement token -- '%s'", []LnsAny{nextToken.Txt}))
        }
    }
}
// 1136: decl @lune.@base.@TransUnit.TransUnit.evalMacro
func (self *TransUnit_TransUnit) evalMacro(_env *LnsEnv, firstToken *Types_Token,macroTypeInfo *Ast_TypeInfo,expList LnsAny) *Nodes_ExpMacroExpNode {
    var stmtList *LnsList
    stmtList = NewLnsList([]LnsAny{})
    self.FP.evalMacroOp(_env, firstToken, macroTypeInfo, expList, Macro_EvalMacroCallback(func(_env *LnsEnv) {
        if macroTypeInfo.FP.Get_retTypeInfoList(_env).Len() == 0{
            self.FP.analyzeStatementList(_env, stmtList, true, "}")
        } else { 
            stmtList.Insert(Nodes_Node2Stem(self.FP.analyzeExp(_env, false, false, false, nil, nil)))
        }
    }))
    var expTypeList *LnsList
    expTypeList = macroTypeInfo.FP.Get_retTypeInfoList(_env)
    if macroTypeInfo.FP.Get_retTypeInfoList(_env).Len() > 0{
        var macroRetTypeList *LnsList
        macroRetTypeList = macroTypeInfo.FP.Get_retTypeInfoList(_env)
        if stmtList.Len() == 1{
            var node *Nodes_Node
            node = stmtList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
            var retType *Ast_TypeInfo
            retType = macroRetTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if retType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeMultiExp, nil, nil){
                expTypeList = node.FP.Get_expTypeList(_env)
            } else if retType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeExp, nil, nil){
                if node.FP.Get_expTypeList(_env).Len() == 1{
                    expTypeList = node.FP.Get_expTypeList(_env)
                } else { 
                    self.FP.AddErrMess(_env, firstToken.Pos, "__exp can't return multiple values. use __exps.")
                }
            } else if node.FP.Get_expTypeList(_env).Len() == 1{
                if retType.FP.Equals(_env, self.ProcessInfo, node.FP.Get_expType(_env), nil, nil){
                    expTypeList = node.FP.Get_expTypeList(_env)
                } else { 
                    self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("mismatch type -- %s != %s", []LnsAny{macroRetTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.GetTxt(_env, nil, nil, nil), node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            } else { 
                self.FP.AddErrMess(_env, firstToken.Pos, "macro can't return multiple values.")
            }
        } else { 
            self.FP.AddErrMess(_env, firstToken.Pos, "macro to return value must be one statemnt.")
        }
    } else { 
        expTypeList = NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})
    }
    return Nodes_ExpMacroExpNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), expTypeList, macroTypeInfo, stmtList)
}
// 1277: decl @lune.@base.@TransUnit.TransUnit.checkStringFormat
func (self *TransUnit_TransUnit) checkStringFormat(_env *LnsEnv, pos Types_Position,formatTxt string,argTypeList *LnsList) {
    var opList *LnsList
    opList = TransUnit_findForm(_env, formatTxt)
    if opList.Len() != argTypeList.Len(){
        self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("argument number is mismatch -- %d != %d (%s)", []LnsAny{opList.Len(), argTypeList.Len(), _env.GetVM().String_sub(formatTxt,1, 20)}))
        return 
    }
    for _index, _op := range( opList.Items ) {
        index := _index + 1
        op := _op.(string)
        var argType *Ast_TypeInfo
        argType = argTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var match LnsInt
        var reqType *Ast_TypeInfo
        match,reqType = TransUnit_isMatchStringFormatType(_env, op, argType, self.TargetLuaVer)
        if match == TransUnit_FormType__Unmatch{
            var mess string
            mess = _env.GetVM().String_format("type must be %s -- %s", []LnsAny{reqType.FP.GetTxt(_env, nil, nil, nil), argType.FP.GetTxt(_env, nil, nil, nil)})
            self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("argument(%d) %s", []LnsAny{index, mess}))
        }
    }
}
// 1316: decl @lune.@base.@TransUnit.TransUnit.prepareExpCall
func (self *TransUnit_TransUnit) prepareExpCall(_env *LnsEnv, position Types_Position,funcTypeInfo *Ast_TypeInfo,genericTypeList *LnsList,genericsClass *Ast_TypeInfo)(*LnsMap, LnsAny) {
    if funcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Macro{
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.MacroCtrl.FP.IsUsing__var(_env, funcTypeInfo)) &&
            _env.SetStackVal( funcTypeInfo.FP.Get_parentInfo(_env) != self.FP.Get_curNsInfo(_env).FP.Get_typeInfo(_env)) ).(bool)){
            self.FP.AddErrMess(_env, position, _env.GetVM().String_format("The macro (%s) only use at the same scope declared it, because it's access __ver.", []LnsAny{funcTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
        }
        self.MacroCtrl.FP.StartAnalyzeArgMode(_env, funcTypeInfo)
    }
    var work *Types_Token
    work = self.FP.getToken(_env, nil)
    var argList LnsAny
    argList = nil
    if work.Txt != ")"{
        self.FP.Pushback(_env)
        argList = self.FP.analyzeExpList(_env, false, false, false, nil, funcTypeInfo.FP.Get_argTypeInfoList(_env), nil)
        self.FP.checkNextToken(_env, ")")
        if argList != nil{
            argList_315 := argList.(*Nodes_ExpListNode)
            for _, _argNode := range( argList_315.FP.Get_expList(_env).Items ) {
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( Lns_op_not(argNode.FP.CanBeRight(_env, self.ProcessInfo))) &&
                    _env.SetStackVal( argNode.FP.Get_kind(_env) != Nodes_NodeKind_get_Abbr(_env)) ).(bool)){
                    self.FP.AddErrMess(_env, argNode.FP.Get_effectivePos(_env), _env.GetVM().String_format("this node can't be r-value. -- %s", []LnsAny{Nodes_getNodeKindName(_env, argNode.FP.Get_kind(_env))}))
                }
            }
        }
    }
    var matchResult LnsInt
    var alt2typeMap *LnsMap
    var workArgList LnsAny
    matchResult,alt2typeMap,workArgList = self.FP.checkMatchValType(_env, position, funcTypeInfo, argList, genericTypeList, genericsClass)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( funcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Macro) &&
        _env.SetStackVal( matchResult == Ast_MatchType__Error) ).(bool)){
        self.FP.Error(_env, _env.GetVM().String_format("unmatch macro arguments. -- %s", []LnsAny{funcTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
    }
    if funcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Macro{
        self.MacroCtrl.FP.FinishMacroMode(_env)
    }
    return alt2typeMap, workArgList
}
// 1373: decl @lune.@base.@TransUnit.TransUnit.checkArgForStringForm
func (self *TransUnit_TransUnit) checkArgForStringForm(_env *LnsEnv, firstToken *Types_Token,argList *Nodes_ExpListNode) {
    var formArgTypeList *LnsList
    formArgTypeList = NewLnsList([]LnsAny{})
    var formatTxt string
    formatTxt = ""
    if argList.FP.Get_expList(_env).Len() > 0{
        var argNode *Nodes_Node
        argNode = argList.FP.Get_expList(_env).GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
        if argNode.FP.Get_kind(_env) != Nodes_NodeKind_get_LiteralString(_env){
            return 
        }
        {
            _literal := TransUnit_convExp4_4186(Lns_2DDD(argNode.FP.GetLiteral(_env)))
            if !Lns_IsNil( _literal ) {
                literal := _literal
                switch _matchExp0 := literal.(type) {
                case *Nodes_Literal__Str:
                val := _matchExp0.Val1
                    formatTxt = val
                }
            }
        }
    }
    if argList.FP.Get_expList(_env).Len() > 1{
        {
            _toDDDNode := Nodes_ExpToDDDNodeDownCastF(argList.FP.Get_expList(_env).GetAt(2).(Nodes_NodeDownCast).ToNodes_Node().FP)
            if !Lns_IsNil( _toDDDNode ) {
                toDDDNode := _toDDDNode.(*Nodes_ExpToDDDNode)
                for _, _workNode := range( toDDDNode.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
                    workNode := _workNode.(Nodes_NodeDownCast).ToNodes_Node()
                    formArgTypeList.Insert(Ast_TypeInfo2Stem(workNode.FP.Get_expType(_env)))
                }
            } else {
                self.FP.Error(_env, _env.GetVM().String_format("illegal node -- %s", []LnsAny{Nodes_getNodeKindName(_env, argList.FP.Get_expList(_env).GetAt(2).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_kind(_env))}))
            }
        }
    }
    self.FP.checkStringFormat(_env, firstToken.Pos, formatTxt, formArgTypeList)
}
// 1410: decl @lune.@base.@TransUnit.TransUnit.checkArgForSort
func (self *TransUnit_TransUnit) checkArgForSort(_env *LnsEnv, firstToken *Types_Token,genericTypeList *LnsList,argList *Nodes_ExpListNode) {
    if argList.FP.Get_expTypeList(_env).Len() > 0{
        var callback *Ast_TypeInfo
        callback = argList.FP.Get_expTypeList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if callback == Ast_builtinTypeAbbr{
            return 
        }
        if callback.FP.Get_retTypeInfoList(_env).Len() != 1{
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("The callback's to return value of sort() must have a value. -- %d", []LnsAny{callback.FP.Get_retTypeInfoList(_env).Len()}))
            return 
        }
        if Lns_op_not(Ast_builtinTypeBool.FP.Equals(_env, self.ProcessInfo, callback.FP.Get_retTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), nil, nil)){
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("The callback's return type of sort() must be bool. -- '%s'", []LnsAny{callback.FP.Get_retTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.GetTxt(_env, nil, nil, nil)}))
        }
        if callback.FP.Get_argTypeInfoList(_env).Len() != 2{
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("The callback's argument must have 2 arguments. -- '%s'", []LnsAny{callback.FP.Get_display_stirng(_env)}))
        }
        if genericTypeList.Len() == 1{
            for _index, _argType := range( callback.FP.Get_argTypeInfoList(_env).Items ) {
                index := _index + 1
                argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if Lns_op_not(genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Equals(_env, self.ProcessInfo, argType, nil, nil)){
                    self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("The callback's argument(%d) type must be -- '%s'", []LnsAny{index, genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        } else { 
            self.FP.AddErrMess(_env, firstToken.Pos, "The generics of the list is illegal")
        }
    }
}
// 1460: decl @lune.@base.@TransUnit.TransUnit.processFunc
func (self *TransUnit_TransUnit) processFunc(_env *LnsEnv, firstToken *Types_Token,nextToken *Types_Token,refFieldNode LnsAny,funcExp *Nodes_Node,funcType *Ast_TypeInfo,alt2typeMap *LnsMap,genericTypeList *LnsList,genericsClass *Ast_TypeInfo,argList LnsAny) *Nodes_Node {
    var funcSymbol LnsAny
    var symbolInfoList *LnsList
    symbolInfoList = funcExp.FP.GetSymbolInfo(_env)
    if symbolInfoList.Len() > 0{
        var symbol *Ast_SymbolInfo
        symbol = symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if symbol.FP.Get_kind(_env) == Ast_SymbolKind__Typ{
            self.FP.AddErrMess(_env, funcExp.FP.Get_pos(_env), _env.GetVM().String_format("can't call any Type. -- %s", []LnsAny{symbol.FP.Get_name(_env)}))
        }
        funcSymbol = symbol
    } else { 
        funcSymbol = nil
    }
    var funcTypeInfo *Ast_TypeInfo
    funcTypeInfo = funcExp.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env)
    var nilAccess bool
    if nextToken.Txt == "$("{
        if funcTypeInfo.FP.Get_nilable(_env){
            funcTypeInfo = funcTypeInfo.FP.Get_nonnilableType(_env)
            nilAccess = true
        } else { 
            self.FP.addWarnMess(_env, funcExp.FP.Get_pos(_env), _env.GetVM().String_format("This is not nilable. -- %s", []LnsAny{funcTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            nilAccess = false
        }
    } else { 
        nilAccess = false
    }
    if _switch1 := (funcTypeInfo.FP.Get_kind(_env)); _switch1 == Ast_TypeInfoKind__Method || _switch1 == Ast_TypeInfoKind__Func || _switch1 == Ast_TypeInfoKind__Form || _switch1 == Ast_TypeInfoKind__FormFunc {
    } else {
        {
            _extType := Ast_ExtTypeInfoDownCastF(funcTypeInfo.FP)
            if !Lns_IsNil( _extType ) {
                extType := _extType.(*Ast_ExtTypeInfo)
                if _switch0 := (extType.FP.Get_extedType(_env).FP.Get_kind(_env)); _switch0 == Ast_TypeInfoKind__Method || _switch0 == Ast_TypeInfoKind__Func || _switch0 == Ast_TypeInfoKind__Form || _switch0 == Ast_TypeInfoKind__FormFunc {
                } else {
                    self.FP.Error(_env, _env.GetVM().String_format("can't call the type -- %s, %s", []LnsAny{funcTypeInfo.FP.GetTxt(_env, nil, nil, nil), Ast_TypeInfoKind_getTxt( funcTypeInfo.FP.Get_kind(_env))}))
                }
            } else {
                self.FP.Error(_env, _env.GetVM().String_format("can't call the type -- %s, %s", []LnsAny{funcTypeInfo.FP.GetTxt(_env, nil, nil, nil), Ast_TypeInfoKind_getTxt( funcTypeInfo.FP.Get_kind(_env))}))
            }
        }
    }
    var prefixMutMode LnsInt
    if refFieldNode != nil{
        refFieldNode_375 := refFieldNode.(*Nodes_RefFieldNode)
        prefixMutMode = refFieldNode_375.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_mutMode(_env)
    } else {
        prefixMutMode = Ast_MutMode__Mut
    }
    var retTypeInfoList *LnsList
    retTypeInfoList = NewLnsList([]LnsAny{})
    for _, _retType := range( funcTypeInfo.FP.Get_retTypeInfoList(_env).Items ) {
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var workType *Ast_TypeInfo
        workType = retType
        {
            _applyType := retType.FP.ApplyGeneric(_env, self.ProcessInfo, alt2typeMap, self.moduleType)
            if !Lns_IsNil( _applyType ) {
                applyType := _applyType.(*Ast_TypeInfo)
                workType = applyType
            } else {
                if funcTypeInfo == self.builtinFunc.List_remove{
                    workType = genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo(_env)
                } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( funcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Func) &&
                    _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( funcTypeInfo.FP.Get_rawTxt(_env) == "_fromMap") ||
                        _env.SetStackVal( funcTypeInfo.FP.Get_rawTxt(_env) == "_fromStem") ).(bool))) &&
                    _env.SetStackVal( genericsClass.FP.IsInheritFrom(_env, self.ProcessInfo, Ast_builtinTypeMapping, alt2typeMap)) ).(bool)){
                    workType = genericsClass.FP.Get_nilableTypeInfo(_env)
                } else { 
                    self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("not support generics yet. -- %s", []LnsAny{retType.FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        }
        if retType.FP.Get_mutMode(_env) == Ast_MutMode__Depend{
            if prefixMutMode == Ast_MutMode__Mut{
                retTypeInfoList.Insert(Ast_TypeInfo2Stem(workType.FP.Get_srcTypeInfo(_env)))
            } else { 
                retTypeInfoList.Insert(Ast_TypeInfo2Stem(self.FP.createModifier(_env, workType.FP.Get_srcTypeInfo(_env), Ast_MutMode__IMut)))
            }
        } else { 
            retTypeInfoList.Insert(Ast_TypeInfo2Stem(workType))
        }
    }
    if refFieldNode != nil{
        refFieldNode_392 := refFieldNode.(*Nodes_RefFieldNode)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( funcTypeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.List_unpack, nil, nil)) ||
            _env.SetStackVal( funcTypeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.Array_unpack, nil, nil)) ).(bool){
            var prefixType *Ast_TypeInfo
            prefixType = refFieldNode_392.FP.Get_prefix(_env).FP.Get_expType(_env)
            if prefixType.FP.Get_itemTypeInfoList(_env).Len() > 0{
                var dddType *Ast_DDDTypeInfo
                dddType = self.ProcessInfo.FP.CreateDDD(_env, prefixType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), false, false)
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
            if retType.FP.Get_nilable(_env){
                retList.Insert(Ast_TypeInfo2Stem(retType))
            } else { 
                retList.Insert(Ast_TypeInfo2Stem(retType.FP.Get_nilableTypeInfo(_env)))
            }
        }
        retTypeInfoList = retList
        self.HelperInfo.UseNilAccess = true
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
        argList_408 := argList.(*Nodes_ExpListNode)
        if _switch2 := funcTypeInfo; _switch2 == self.builtinFunc.String_format {
            self.FP.checkArgForStringForm(_env, firstToken, argList_408)
        } else if _switch2 == self.builtinFunc.List_sort || _switch2 == self.builtinFunc.Array_sort {
            self.FP.checkArgForSort(_env, firstToken, genericTypeList, argList_408)
        }
    }
    if funcTypeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.Lns__kind, nil, nil){
        {
            _expList := _env.NilAccFin(_env.NilAccPush(argList) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_ExpListNode).FP.Get_expList(_env)}))
            if !Lns_IsNil( _expList ) {
                expList := _expList.(*LnsList)
                if expList.Len() > 0{
                    return &Nodes_LuneKindNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), expList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()).Nodes_Node
                }
            }
        }
        return &Nodes_LuneKindNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), self.FP.createNoneNode(_env, firstToken.Pos)).Nodes_Node
    }
    if funcSymbol != nil{
        funcSymbol_416 := funcSymbol.(*Ast_SymbolInfo)
        if funcSymbol_416.FP.Get_name(_env) == "super"{
            return &Nodes_ExpCallSuperNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), retTypeInfoList, funcSymbol_416.FP.Get_typeInfo(_env).FP.Get_parentInfo(_env), funcSymbol_416.FP.Get_typeInfo(_env), argList).Nodes_Node
        }
    }
    if funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        var work LnsAny
        var err string
        work,err = Ast_convToExtTypeList(_env, self.ProcessInfo, retTypeInfoList)
        if work != nil{
            work_422 := work.(*LnsList)
            retTypeInfoList = work_422
        } else {
            self.FP.AddErrMess(_env, firstToken.Pos, err)
        }
    }
    return &Nodes_ExpCallNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), retTypeInfoList, funcExp, errorFuncFlag, nilAccess, argList).Nodes_Node
}
// 1660: decl @lune.@base.@TransUnit.TransUnit.checkNoasyncType
func (self *TransUnit_TransUnit) checkNoasyncType(_env *LnsEnv, pos Types_Position,funcTypeInfo *Ast_TypeInfo) {
    if self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.IsAnalyzingBlockArg(_env){
        return 
    }
    var isExt bool
    isExt = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( funcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext) ||
        _env.SetStackVal( self.builtinFunc.FP.IsLuavalFunc(_env, funcTypeInfo)) ).(bool)
    if isExt{
        var curType *Ast_TypeInfo
        curType = self.FP.GetCurrentNamespaceTypeInfo(_env)
        if Lns_op_not(self.FP.GetNSInfo(_env, curType).FP.CanAccessLuaval(_env)){
            self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("can't access Luaval function without __luago. -- %s on %s", []LnsAny{funcTypeInfo.FP.GetTxt(_env, nil, nil, nil), curType.FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    if funcTypeInfo.FP.Get_asyncMode(_env) == Ast_Async__Noasync{
        var curType *Ast_TypeInfo
        curType = self.FP.GetCurrentNamespaceTypeInfo(_env)
        if _switch1 := curType.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Func || _switch1 == Ast_TypeInfoKind__Method {
            if _switch0 := curType.FP.Get_asyncMode(_env); _switch0 == Ast_Async__Async || _switch0 == Ast_Async__Transient {
                if Lns_op_not(self.FP.GetNSInfo(_env, curType).FP.CanAccessNoasync(_env)){
                    self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("can't access noasync function in async. -- %s on %s", []LnsAny{funcTypeInfo.FP.GetTxt(_env, nil, nil, nil), curType.FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        }
    }
}
// 1709: decl @lune.@base.@TransUnit.TransUnit.analyzeExpCall
func (self *TransUnit_TransUnit) analyzeExpCall(_env *LnsEnv, firstToken *Types_Token,funcExp *Nodes_Node,nextToken *Types_Token)(*Nodes_Node, *Types_Token) {
    self.FP.checkSymbolHavingValue(_env, funcExp.FP.Get_pos(_env), funcExp.FP.GetSymbolInfo(_env))
    var funcTypeInfo *Ast_TypeInfo
    funcTypeInfo = funcExp.FP.Get_expType(_env).FP.Get_nonnilableType(_env)
    self.FP.checkNoasyncType(_env, funcExp.FP.Get_effectivePos(_env), funcTypeInfo)
    if funcTypeInfo == Ast_builtinTypeLnsLoad{
        self.MacroCtrl.FP.SetToUseLnsLoad(_env)
    }
    var genericTypeList *LnsList
    genericTypeList = funcTypeInfo.FP.Get_itemTypeInfoList(_env)
    var refFieldNode LnsAny
    refFieldNode = nil
    var genericsClass *Ast_TypeInfo
    genericsClass = Ast_headTypeInfo
    {
        _refField := Nodes_RefFieldNodeDownCastF(funcExp.FP)
        if !Lns_IsNil( _refField ) {
            refField := _refField.(*Nodes_RefFieldNode)
            refFieldNode = refField
            var classType *Ast_TypeInfo
            classType = refField.FP.Get_prefix(_env).FP.Get_expType(_env)
            genericsClass = classType
            if funcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
                genericTypeList = classType.FP.Get_itemTypeInfoList(_env)
            }
        }
    }
    var alt2typeMap *LnsMap
    var argList LnsAny
    alt2typeMap,argList = self.FP.prepareExpCall(_env, funcExp.FP.Get_pos(_env), funcTypeInfo, genericTypeList, genericsClass)
    if funcTypeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.List_insert, nil, nil){
        if argList != nil{
            argList_449 := argList.(*Nodes_ExpListNode)
            if argList_449.FP.Get_expType(_env).FP.Get_nilable(_env){
                self.FP.AddErrMess(_env, argList_449.FP.Get_pos(_env), "list can't insert nilable")
            }
        }
    }
    if funcTypeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.Set_add, nil, nil){
        if argList != nil{
            argList_453 := argList.(*Nodes_ExpListNode)
            if argList_453.FP.Get_expType(_env).FP.Get_nilable(_env){
                self.FP.AddErrMess(_env, argList_453.FP.Get_pos(_env), "set can't add nilable")
            }
        }
    } else if funcTypeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.List_remove, nil, nil){
        if genericTypeList.Len() > 0{
            if genericTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilable(_env){
                self.FP.addWarnMess(_env, funcExp.FP.Get_pos(_env), "remove() is dangerous for nilable's list.")
            }
        }
    }
    if funcTypeInfo.FP.Get_rawTxt(_env) == ""{
        self.FP.AddErrMess(_env, funcExp.FP.Get_pos(_env), "can't directly call the declared function.")
    }
    var exp *Nodes_Node
    if funcTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Macro{
        exp = &self.FP.evalMacro(_env, firstToken, funcTypeInfo, argList).Nodes_Node
    } else { 
        exp = self.FP.processFunc(_env, firstToken, nextToken, refFieldNode, funcExp, funcTypeInfo, alt2typeMap, genericTypeList, genericsClass, argList)
    }
    return exp, self.FP.GetTokenNoErr(_env, nil)
}
// 1785: decl @lune.@base.@TransUnit.TransUnit.analyzeExpCast
func (self *TransUnit_TransUnit) analyzeExpCast(_env *LnsEnv, firstToken *Types_Token,opTxt string,exp *Nodes_Node) *Nodes_Node {
    var castTypeNode *Nodes_RefTypeNode
    castTypeNode = self.FP.analyzeRefType(_env, Ast_AccessMode__Local, false, false)
    var castType *Ast_TypeInfo
    castType = castTypeNode.FP.Get_expType(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext) &&
        _env.SetStackVal( castType.FP.Get_kind(_env) != Ast_TypeInfoKind__Ext) &&
        _env.SetStackVal( castType.FP.Get_kind(_env) != Ast_TypeInfoKind__Stem) ).(bool)){
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( opTxt == "@@") ||
            _env.SetStackVal( opTxt == "@@=") ).(bool){
            castType = self.FP.createModifier(_env, castType, Ast_MutMode__IMut)
        }
        castType = self.FP.createExtType(_env, castTypeNode.FP.Get_pos(_env), castType)
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( castType.FP.Get_kind(_env) == Ast_TypeInfoKind__Form) &&
        _env.SetStackVal( exp.FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Stem) ).(bool)){
        if self.FP.supportLang(_env, LuneControl_Code__C){
            self.FP.addWarnMess(_env, castTypeNode.FP.Get_pos(_env), "not support cast from stem to form for transcompiling to c-lang.")
        }
    }
    var expType *Ast_TypeInfo
    expType = exp.FP.Get_expType(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( opTxt == "@@@") ||
        _env.SetStackVal( opTxt == "@@=") ).(bool){
        if castType.FP.Get_itemTypeInfoList(_env).Len() > 0{
            self.FP.AddErrMess(_env, castTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("not support cast for generics class yet -- %s", []LnsAny{castType.FP.GetTxt(_env, nil, nil, nil)}))
        }
        if _switch0 := castType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__IF || _switch0 == Ast_TypeInfoKind__Class || _switch0 == Ast_TypeInfoKind__Prim {
        } else {
            if opTxt != "@@="{
                self.FP.AddErrMess(_env, castTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("not support cast -- %s", []LnsAny{castType.FP.GetTxt(_env, nil, nil, nil)}))
            }
        }
        if opTxt == "@@="{
            var orgExpType *Ast_TypeInfo
            orgExpType = expType.FP.Get_extedType(_env)
            var orgCastType *Ast_TypeInfo
            orgCastType = castType.FP.Get_extedType(_env)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( orgCastType.FP.Get_kind(_env) != Ast_TypeInfoKind__IF) &&
                _env.SetStackVal( orgCastType.FP.Get_kind(_env) != Ast_TypeInfoKind__Class) ).(bool)){
                self.FP.AddErrMess(_env, castTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("'@@=' cast must be class or interface. -- %s", []LnsAny{castType.FP.GetTxt(_env, nil, nil, nil)}))
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( orgExpType.FP.Get_srcTypeInfo(_env) != Ast_builtinTypeStem) &&
                _env.SetStackVal( orgExpType.FP.Get_kind(_env) != Ast_TypeInfoKind__IF) &&
                _env.SetStackVal( orgExpType.FP.Get_kind(_env) != Ast_TypeInfoKind__Class) ).(bool)){
                self.FP.AddErrMess(_env, castTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("'@@=' cast must be class or interface. -- %s", []LnsAny{castType.FP.GetTxt(_env, nil, nil, nil)}))
            }
            if Lns_op_not(Ast_isStruct(_env, orgCastType)){
                self.FP.AddErrMess(_env, castTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("'@@=' cast type can't use class has method -- %s", []LnsAny{castType.FP.GetTxt(_env, nil, nil, nil)}))
            }
        }
    } else { 
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( castType != Ast_builtinTypeString) &&
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( castType.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ||
                _env.SetStackVal( castType.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ).(bool))) ).(bool)){
            self.FP.addWarnMess(_env, castTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("use '@@@' cast for class or interface. -- %s", []LnsAny{castType.FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( opTxt != "@@@") &&
        _env.SetStackVal( expType.FP.Get_nilable(_env)) &&
        _env.SetStackVal( Lns_op_not(castType.FP.Get_nilable(_env))) ).(bool)){
        self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("can't cast from nilable to not nilable  -- %s->%s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil), castType.FP.GetTxt(_env, nil, nil, nil)}))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(_env, expType))) &&
        _env.SetStackVal( Ast_TypeInfo_isMut(_env, castType)) ).(bool)){
        castType = self.FP.createModifier(_env, castType, Ast_MutMode__IMut)
    }
    if Lns_car(castType.FP.CanEvalWith(_env, self.ProcessInfo, expType, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( castType.FP.Get_kind(_env) != Ast_TypeInfoKind__Ext) &&
            _env.SetStackVal( Lns_op_not((_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( castType.FP.Get_kind(_env) == Ast_TypeInfoKind__Stem) &&
                _env.SetStackVal( Ast_isExtType(_env, expType)) ).(bool)))) ).(bool)){
            self.FP.addWarnMess(_env, castTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("This cast isn't need. (%s <- %s)", []LnsAny{castType.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), expType.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
        }
    } else if Lns_op_not(Lns_car(expType.FP.CanEvalWith(_env, self.ProcessInfo, castType, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)){
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(Ast_isNumberType(_env, expType))) ||
            _env.SetStackVal( Lns_op_not(Ast_isNumberType(_env, castType))) ).(bool){
            self.FP.AddErrMess(_env, castTypeNode.FP.Get_pos(_env), _env.GetVM().String_format("This type can't cast. (%s <- %s)", []LnsAny{castType.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), expType.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
        }
    }
    if opTxt == "@@@"{
        castType = castType.FP.Get_nilableTypeInfo(_env)
    }
    return &Nodes_ExpCastNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(castType)}), self.NodeManager.FP.MultiTo1(_env, exp), castType, castTypeNode, opTxt, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( opTxt != "@@@") &&
        _env.SetStackVal( Nodes_CastKind__Force) ||
        _env.SetStackVal( Nodes_CastKind__Normal) ).(LnsInt)).Nodes_Node
}
// 1912: decl @lune.@base.@TransUnit.TransUnit.analyzeExpCont
func (self *TransUnit_TransUnit) analyzeExpCont(_env *LnsEnv, firstToken *Types_Token,exp *Nodes_Node,skipFlag bool,canLeftExp bool) *Nodes_Node {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, true)
    if nextToken.Kind == Types_TokenKind__Eof{
        return exp
    }
    if Lns_op_not(skipFlag){
        for {
            var matchFlag bool
            matchFlag = false
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( nextToken.Txt == "[") ||
                _env.SetStackVal( nextToken.Txt == "$[") ).(bool){
                matchFlag = true
                exp = self.FP.analyzeExpRefItem(_env, nextToken, exp, nextToken.Txt == "$[")
                nextToken = self.FP.getToken(_env, nil)
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( nextToken.Txt == "(") ||
                _env.SetStackVal( nextToken.Txt == "$(") ).(bool){
                matchFlag = true
                exp, nextToken = self.FP.analyzeExpCall(_env, firstToken, exp, nextToken)
            }
            if Lns_op_not(matchFlag){ break }
        }
        if _switch0 := nextToken.Txt; _switch0 == "@@" || _switch0 == "@@@" || _switch0 == "@@=" {
            exp = self.FP.analyzeExpCast(_env, firstToken, nextToken.Txt, exp)
            nextToken = self.FP.getToken(_env, nil)
        }
    }
    if _switch1 := nextToken.Txt; _switch1 == "." {
        return self.FP.analyzeExpSymbol(_env, firstToken, self.FP.getToken(_env, nil), TransUnit_ExpSymbolMode__Field, exp, skipFlag, canLeftExp)
    } else if _switch1 == "$." {
        return self.FP.analyzeExpSymbol(_env, firstToken, self.FP.getToken(_env, nil), TransUnit_ExpSymbolMode__FieldNil, exp, skipFlag, canLeftExp)
    } else if _switch1 == ".$" {
        return self.FP.analyzeExpSymbol(_env, firstToken, self.FP.getToken(_env, nil), TransUnit_ExpSymbolMode__Get, exp, skipFlag, canLeftExp)
    } else if _switch1 == "$.$" {
        return self.FP.analyzeExpSymbol(_env, firstToken, self.FP.getToken(_env, nil), TransUnit_ExpSymbolMode__GetNil, exp, skipFlag, canLeftExp)
    }
    self.FP.Pushback(_env)
    return exp
}
// 1969: decl @lune.@base.@TransUnit.TransUnit.analyzeAccessClassField
func (self *TransUnit_TransUnit) analyzeAccessClassField(_env *LnsEnv, classTypeInfo *Ast_TypeInfo,mode LnsInt,token *Types_Token)(*Ast_TypeInfo, LnsAny, bool) {
    if _switch0 := classTypeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List {
        classTypeInfo = Ast_builtinTypeList
    } else if _switch0 == Ast_TypeInfoKind__Array {
        classTypeInfo = Ast_builtinTypeArray
    } else if _switch0 == Ast_TypeInfoKind__Set {
        classTypeInfo = Ast_builtinTypeSet
    } else if _switch0 == Ast_TypeInfoKind__Box {
        classTypeInfo = &Ast_builtinTypeBox.Ast_TypeInfo
    }
    var className string
    className = classTypeInfo.FP.GetTxt(_env, nil, nil, nil)
    var classScope *Ast_Scope
    
    {
        _classScope := classTypeInfo.FP.Get_scope(_env)
        if _classScope == nil{
            self.FP.Error(_env, _env.GetVM().String_format("not found field: %s, %s", []LnsAny{className, classTypeInfo}))
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
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__Get) ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__GetNil) ).(bool){
        var fieldSymbolInfo LnsAny
        fieldSymbolInfo = classScope.FP.GetSymbolInfo(_env, _env.GetVM().String_format("get_%s", []LnsAny{token.Txt}), self.FP.Get_scope(_env), false, self.ScopeAccess)
        if fieldSymbolInfo != nil{
            fieldSymbolInfo_518 := fieldSymbolInfo.(*Ast_SymbolInfo)
            if (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( fieldSymbolInfo_518.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) ||
                _env.SetStackVal( fieldSymbolInfo_518.FP.Get_kind(_env) == Ast_SymbolKind__Fun) ).(bool)){
                var retTypeList *LnsList
                retTypeList = fieldSymbolInfo_518.FP.Get_typeInfo(_env).FP.Get_retTypeInfoList(_env)
                symbolInfo = fieldSymbolInfo_518
                if retTypeList.Len() > 0{
                    {
                        _applyedType := retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.ApplyGeneric(_env, self.ProcessInfo, classTypeInfo.FP.CreateAlt2typeMap(_env, false), self.moduleType)
                        if !Lns_IsNil( _applyedType ) {
                            applyedType := _applyedType.(*Ast_TypeInfo)
                            fieldTypeInfo = applyedType
                        } else {
                            fieldTypeInfo = retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        }
                    }
                }
                if fieldSymbolInfo_518.FP.Get_typeInfo(_env).FP.Get_argTypeInfoList(_env).Len() > 0{
                    self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("can't use '$' with -- %s", []LnsAny{fieldSymbolInfo_518.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
                getterFlag = true
            }
        }
    }
    if Lns_op_not(symbolInfo){
        symbolInfo = classScope.FP.GetSymbolInfoField(_env, token.Txt, true, self.FP.Get_scope(_env), self.ScopeAccess)
        if Lns_op_not(symbolInfo){
            symbolInfo = classScope.FP.GetSymbolInfoIfField(_env, token.Txt, self.FP.Get_scope(_env), self.ScopeAccess)
        }
        if symbolInfo != nil{
            symbolInfo_529 := symbolInfo.(*Ast_SymbolInfo)
            fieldTypeInfo = symbolInfo_529.FP.Get_typeInfo(_env)
        }
    }
    if Lns_op_not(fieldTypeInfo){
        for _name, _val := range( classScope.FP.Get_symbol2SymbolInfoMap(_env).Items ) {
            name := _name.(string)
            val := _val.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            Util_errorLog(_env, _env.GetVM().String_format("debug: %s, %s", []LnsAny{name, val}))
        }
        Util_errorLog(_env, _env.GetVM().String_format("class, scope: -- %s, %s", []LnsAny{classTypeInfo, classScope}))
        self.FP.Error(_env, _env.GetVM().String_format("not found field typeInfo: %s.%s -- %s", []LnsAny{classTypeInfo.FP.GetFullName(_env, self.TypeNameCtrl, self.FP.Get_scope(_env).FP, false), token.Txt, Ast_TypeInfoKind_getTxt( classTypeInfo.FP.Get_kind(_env))}))
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = Lns_unwrapDefault( fieldTypeInfo, Ast_builtinTypeNone).(*Ast_TypeInfo)
    if symbolInfo != nil{
        symbolInfo_536 := symbolInfo.(*Ast_SymbolInfo)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.FP.inAnalyzingState(_env, TransUnit_AnalyzingState__InitBlock)) ||
            _env.SetStackVal( self.FP.inAnalyzingState(_env, TransUnit_AnalyzingState__ClassMethod)) ).(bool){
            var errorMess LnsAny
            errorMess = nil
            if Lns_isCondTrue( _env.NilAccFin(_env.NilAccPush(self.NsInfoMap.Get(symbolInfo_536.FP.Get_typeInfo(_env))) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*TransUnitIF_NSInfo).FP.Get_nobody(_env)}))){
                errorMess = _env.GetVM().String_format("It can't call prototype function from static -- %s", []LnsAny{symbolInfo_536.FP.Get_name(_env)})
            }
            if errorMess != nil{
                errorMess_541 := errorMess.(string)
                self.FP.AddErrMess(_env, token.Pos, errorMess_541)
            }
        } else if self.FP.inAnalyzingState(_env, TransUnit_AnalyzingState__Constructor){
            var errorMess LnsAny
            errorMess = nil
            if Lns_isCondTrue( _env.NilAccFin(_env.NilAccPush(self.NsInfoMap.Get(symbolInfo_536.FP.Get_typeInfo(_env))) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*TransUnitIF_NSInfo).FP.Get_nobody(_env)}))){
                errorMess = "It can't call prototype function from '__init'"
            } else { 
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( symbolInfo_536.FP.Get_typeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Method) &&
                    _env.SetStackVal( symbolInfo_536.FP.Get_scope(_env) == classScope) ).(bool)){
                    for _, _val := range( classScope.FP.Get_symbol2SymbolInfoMap(_env).Items ) {
                        val := _val.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( val.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) &&
                            _env.SetStackVal( Lns_op_not(val.FP.Get_staticFlag(_env))) ).(bool)){
                            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( Lns_op_not(val.FP.Get_hasValueFlag(_env))) &&
                                _env.SetStackVal( Lns_op_not(val.FP.Get_typeInfo(_env).FP.Get_nilable(_env))) ).(bool)){
                                errorMess = _env.GetVM().String_format("Set member(%s) before to access the method-- %s", []LnsAny{val.FP.Get_name(_env), symbolInfo_536.FP.Get_name(_env)})
                                break
                            }
                        }
                    }
                }
            }
            if errorMess != nil{
                errorMess_552 := errorMess.(string)
                self.FP.AddErrMess(_env, token.Pos, errorMess_552)
            }
        }
    }
    return typeInfo, symbolInfo, getterFlag
}
// 2092: decl @lune.@base.@TransUnit.TransUnit.dumpComp
func (self *TransUnit_TransUnit) dumpComp(_env *LnsEnv, writer Writer_Writer,pattern string,symbolInfo *Ast_SymbolInfo,getterFlag bool) bool {
    var symbol string
    symbol = symbolInfo.FP.Get_name(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( pattern == "") ||
        _env.SetStackVal( Lns_car(_env.GetVM().String_find(symbol,pattern, nil, nil))) )){
        if getterFlag{
            writer.StartParent(_env, "candidate", false)
            var typeInfo *Ast_TypeInfo
            typeInfo = symbolInfo.FP.Get_typeInfo(_env)
            writer.Write(_env, "type", _env.GetVM().String_format("%s", []LnsAny{Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind(_env))}))
            if _switch0 := (symbolInfo.FP.Get_kind(_env)); _switch0 == Ast_SymbolKind__Mtd || _switch0 == Ast_SymbolKind__Fun || _switch0 == Ast_SymbolKind__Mac {
                writer.Write(_env, "displayTxt", _env.GetVM().String_format("$%s", []LnsAny{Lns_car(_env.GetVM().String_gsub(typeInfo.FP.Get_rawTxt(_env),"^get_", "")).(string)}))
            } else if _switch0 == Ast_SymbolKind__Mbr {
                writer.Write(_env, "displayTxt", _env.GetVM().String_format("$%s: %s", []LnsAny{symbolInfo.FP.Get_name(_env), typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            }
        } else { 
            writer.StartParent(_env, "candidate", false)
            var typeInfo *Ast_TypeInfo
            typeInfo = symbolInfo.FP.Get_typeInfo(_env)
            writer.Write(_env, "type", _env.GetVM().String_format("%s", []LnsAny{Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind(_env))}))
            if _switch1 := (symbolInfo.FP.Get_kind(_env)); _switch1 == Ast_SymbolKind__Fun || _switch1 == Ast_SymbolKind__Mtd || _switch1 == Ast_SymbolKind__Mac {
                writer.Write(_env, "displayTxt", typeInfo.FP.Get_display_stirng_with(_env, symbolInfo.FP.Get_name(_env), nil))
            } else if _switch1 == Ast_SymbolKind__Mbr || _switch1 == Ast_SymbolKind__Var || _switch1 == Ast_SymbolKind__Arg {
                var name string
                name = symbolInfo.FP.Get_name(_env)
                {
                    _algeTypeInfo := Ast_AlgeTypeInfoDownCastF(typeInfo.FP)
                    if !Lns_IsNil( _algeTypeInfo ) {
                        algeTypeInfo := _algeTypeInfo.(*Ast_AlgeTypeInfo)
                        {
                            _valInfo := algeTypeInfo.FP.GetValInfo(_env, name)
                            if !Lns_IsNil( _valInfo ) {
                                valInfo := _valInfo.(*Ast_AlgeValInfo)
                                if valInfo.FP.Get_typeList(_env).Len() > 0{
                                    name = _env.GetVM().String_format("%s(", []LnsAny{name})
                                    for _index, _itemType := range( valInfo.FP.Get_typeList(_env).Items ) {
                                        index := _index + 1
                                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                                        if index > 1{
                                            name = name + ","
                                        }
                                        name = name + itemType.FP.Get_display_stirng(_env)
                                    }
                                    name = name + ")"
                                }
                            }
                        }
                    }
                }
                writer.Write(_env, "displayTxt", _env.GetVM().String_format("%s: %s", []LnsAny{name, typeInfo.FP.Get_display_stirng(_env)}))
            } else if _switch1 == Ast_SymbolKind__Typ {
                writer.Write(_env, "displayTxt", _env.GetVM().String_format("%s", []LnsAny{Lns_car(_env.GetVM().String_gsub(typeInfo.FP.Get_display_stirng(_env),"@", "")).(string)}))
            }
        }
        writer.EndElement(_env)
    }
    return true
}
// 2153: decl @lune.@base.@TransUnit.TransUnit.dumpFieldComp
func (self *TransUnit_TransUnit) dumpFieldComp(_env *LnsEnv, writer Writer_Writer,isPrefixType bool,prefixTypeInfo *Ast_TypeInfo,pattern string,getterPattern LnsAny) {
    var typeInfo *Ast_TypeInfo
    if _switch0 := prefixTypeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List {
        typeInfo = Ast_builtinTypeList
    } else if _switch0 == Ast_TypeInfoKind__Array {
        typeInfo = Ast_builtinTypeArray
    } else if _switch0 == Ast_TypeInfoKind__Set {
        typeInfo = Ast_builtinTypeSet
    } else if _switch0 == Ast_TypeInfoKind__Box {
        typeInfo = &Ast_builtinTypeBox.Ast_TypeInfo
    } else {
        typeInfo = prefixTypeInfo
    }
    var scope *Ast_Scope
    
    {
        _scope := typeInfo.FP.Get_scope(_env)
        if _scope == nil{
            return 
        } else {
            scope = _scope.(*Ast_Scope)
        }
    }
    scope.FP.FilterTypeInfoField(_env, true, self.FP.Get_scopeRO(_env), self.ScopeAccess, Ast_filterForm(func(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) bool {
        if (isPrefixType){
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(symbolInfo.FP.Get_staticFlag(_env))) &&
                _env.SetStackVal( Lns_op_not(symbolInfo.FP.Get_typeInfo(_env).FP.Get_staticFlag(_env))) &&
                _env.SetStackVal( symbolInfo.FP.Get_kind(_env) != Ast_SymbolKind__Typ) ).(bool)){
                return true
            }
        } else if symbolInfo.FP.Get_staticFlag(_env){
            return true
        }
        var symbol string
        symbol = symbolInfo.FP.Get_name(_env)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( symbol != "__init") &&
            _env.SetStackVal( symbol != "__free") &&
            _env.SetStackVal( symbol != "self") ).(bool)){
            if getterPattern != nil{
                getterPattern_594 := getterPattern.(string)
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) ||
                    _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Fun) ).(bool){
                    var retList *LnsList
                    retList = symbolInfo.FP.Get_typeInfo(_env).FP.Get_retTypeInfoList(_env)
                    if retList.Len() == 1{
                        return self.FP.dumpComp(_env, writer, getterPattern_594, symbolInfo, true)
                    }
                }
                return true
            }
            return self.FP.dumpComp(_env, writer, pattern, symbolInfo, false)
        }
        return true
    }))
}
// 2214: decl @lune.@base.@TransUnit.TransUnit.dumpSymbolComp
func (self *TransUnit_TransUnit) dumpSymbolComp(_env *LnsEnv, writer Writer_Writer,scope *Ast_Scope,pattern string) {
    scope.FP.FilterSymbolTypeInfo(_env, scope, self.ModuleScope, self.ScopeAccess, Ast_filterForm(func(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) bool {
        return self.FP.dumpComp(_env, writer, pattern, symbolInfo, false)
    }))
}
// 2224: decl @lune.@base.@TransUnit.TransUnit.checkComp
func (self *TransUnit_TransUnit) checkComp(_env *LnsEnv, token *Types_Token,callback TransUnit_checkCompForm_31_) {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.analyzeMode == TransUnit_AnalyzeMode__Complete) &&
        _env.SetStackVal( self.FP.isTargetToken(_env, token)) ).(bool)){
        var currentModule string
        currentModule = TransUnit_convExp4_8220(Lns_2DDD(_env.GetVM().String_gsub(self.Parser.FP.GetStreamName(_env),"%.lns", "")))
        currentModule = TransUnit_convExp4_8235(Lns_2DDD(_env.GetVM().String_gsub(currentModule,".*/", "")))
        var target string
        target = TransUnit_convExp4_8250(Lns_2DDD(_env.GetVM().String_gsub(self.analyzeModule,"[^%.]+%.", "")))
        if currentModule == target{
            var jsonWriter *Writer_JSON
            jsonWriter = NewWriter_JSON(_env, Lns_io_stdout)
            jsonWriter.FP.StartParent(_env, "lunescript", false)
            var prefix string
            prefix = TransUnit_convExp4_8290(Lns_2DDD(_env.GetVM().String_gsub(token.Txt,"lune$", "")))
            jsonWriter.FP.Write(_env, "prefix", prefix)
            jsonWriter.FP.StartParent(_env, "candidateList", true)
            callback(_env, jsonWriter, prefix)
            jsonWriter.FP.EndElement(_env)
            jsonWriter.FP.EndElement(_env)
            jsonWriter.FP.Fin(_env)
            _env.GetVM().OS_exit(0)
        }
    }
}
// 2248: decl @lune.@base.@TransUnit.TransUnit.checkFieldComp
func (self *TransUnit_TransUnit) checkFieldComp(_env *LnsEnv, getterFlag bool,token *Types_Token,prefixExp *Nodes_Node) {
    if self.analyzeMode != TransUnit_AnalyzeMode__Complete{
        return 
    }
    var prefixSymbolInfoList *LnsList
    prefixSymbolInfoList = prefixExp.FP.GetSymbolInfo(_env)
    var prefixSymbolInfo LnsAny
    prefixSymbolInfo = nil
    if prefixSymbolInfoList.Len() == 1{
        prefixSymbolInfo = prefixSymbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
    }
    self.FP.checkComp(_env, token, TransUnit_checkCompForm_31_(func(_env *LnsEnv, jsonWriter *Writer_JSON,prefix string) {
        var getterPattern LnsAny
        getterPattern = nil
        if getterFlag{
            getterPattern = "^get_" + prefix
        }
        var isPrefixType bool
        isPrefixType = false
        {
            __exp := prefixSymbolInfo
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_SymbolInfo)
                isPrefixType = _exp.FP.Get_kind(_env) == Ast_SymbolKind__Typ
            }
        }
        self.FP.dumpFieldComp(_env, jsonWriter.FP, isPrefixType, prefixExp.FP.Get_expType(_env), _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( prefix == "") &&
            _env.SetStackVal( "") ||
            _env.SetStackVal( "^" + prefix) ).(string), getterPattern)
    }))
}
// 2278: decl @lune.@base.@TransUnit.TransUnit.checkEnumComp
func (self *TransUnit_TransUnit) checkEnumComp(_env *LnsEnv, token *Types_Token,enumTypeInfo *Ast_EnumTypeInfo) {
    if self.analyzeMode != TransUnit_AnalyzeMode__Complete{
        return 
    }
    self.FP.checkComp(_env, token, TransUnit_checkCompForm_31_(func(_env *LnsEnv, jsonWriter *Writer_JSON,prefix string) {
        var scope *Ast_Scope
        
        {
            _scope := enumTypeInfo.FP.Get_scope(_env)
            if _scope == nil{
                return 
            } else {
                scope = _scope.(*Ast_Scope)
            }
        }
        var pattern string
        pattern = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( prefix == "") &&
            _env.SetStackVal( "") ||
            _env.SetStackVal( "^" + prefix) ).(string)
        scope.FP.FilterTypeInfoField(_env, true, self.FP.Get_scopeRO(_env), self.ScopeAccess, Ast_filterForm(func(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) bool {
            if symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mbr{
                return self.FP.dumpComp(_env, jsonWriter.FP, pattern, symbolInfo, false)
            }
            return true
        }))
    }))
}
// 2304: decl @lune.@base.@TransUnit.TransUnit.checkAlgeComp
func (self *TransUnit_TransUnit) checkAlgeComp(_env *LnsEnv, token *Types_Token,algeTypeInfo *Ast_AlgeTypeInfo) {
    if self.analyzeMode != TransUnit_AnalyzeMode__Complete{
        return 
    }
    self.FP.checkComp(_env, token, TransUnit_checkCompForm_31_(func(_env *LnsEnv, jsonWriter *Writer_JSON,prefix string) {
        self.FP.dumpFieldComp(_env, jsonWriter.FP, true, &algeTypeInfo.Ast_TypeInfo, _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( prefix == "") &&
            _env.SetStackVal( "") ||
            _env.SetStackVal( "^" + prefix) ).(string), nil)
    }))
}
// 2320: decl @lune.@base.@TransUnit.TransUnit.checkAsyncSymbol
func (self *TransUnit_TransUnit) checkAsyncSymbol(_env *LnsEnv, symbolInfo *Ast_SymbolInfo,pos Types_Position) {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(self.FP.supportLang(_env, LuneControl_Code__Go))) ||
        _env.SetStackVal( Lns_op_not(self.FP.supportLang(_env, LuneControl_Code__C))) ).(bool){
        return 
    }
    var curNs *Ast_TypeInfo
    curNs = self.FP.GetCurrentNamespaceTypeInfo(_env)
    if _switch0 := curNs.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Func || _switch0 == Ast_TypeInfoKind__Method {
    } else {
        return 
    }
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.GetNSInfo(_env, curNs)
    var warn bool
    warn = false
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(nsInfo.FP.CanAccessNoasync(_env))) &&
        _env.SetStackVal( symbolInfo.FP.Get_name(_env) != "self") ).(bool)){
        if _switch1 := symbolInfo.FP.Get_kind(_env); _switch1 == Ast_SymbolKind__Mbr {
            if Lns_op_not(Ast_isPrimitive(_env, symbolInfo.FP.Get_typeInfo(_env))){
                warn = true
            }
        } else if _switch1 == Ast_SymbolKind__Arg || _switch1 == Ast_SymbolKind__Var {
            if Lns_op_not(symbolInfo.FP.Get_scope(_env).FP.IsInnerOf(_env, Lns_unwrap( curNs.FP.Get_scope(_env)).(*Ast_Scope))){
                warn = true
            }
        }
    }
    if warn{
        if _switch2 := curNs.FP.Get_kind(_env); _switch2 == Ast_TypeInfoKind__Func {
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(nsInfo.FP.CanAccessNoasync(_env))) &&
                _env.SetStackVal( curNs.FP.Get_asyncMode(_env) != Ast_Async__Transient) ).(bool)){
                if Lns_op_not(self.FP.canBeAsyncParam(_env, symbolInfo.FP.Get_typeInfo(_env))){
                    self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("can't access the mutable type's symbol(%s) from async (%s).", []LnsAny{symbolInfo.FP.Get_name(_env), nsInfo.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        } else if _switch2 == Ast_TypeInfoKind__Method {
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(nsInfo.FP.CanAccessNoasync(_env))) ||
                _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( symbolInfo.FP.Get_staticFlag(_env)) &&
                    _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) ).(bool))) ).(bool){
                if Lns_op_not(self.FP.canBeAsyncParam(_env, symbolInfo.FP.Get_typeInfo(_env))){
                    self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("can't access the mutable type's symbol(%s) from async (%s).", []LnsAny{symbolInfo.FP.Get_name(_env), nsInfo.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        }
    }
}
// 2383: decl @lune.@base.@TransUnit.TransUnit.checkAsyncField
func (self *TransUnit_TransUnit) checkAsyncField(_env *LnsEnv, symbolInfo *Ast_SymbolInfo,pos Types_Position) {
    var curNs *Ast_TypeInfo
    curNs = self.FP.GetCurrentNamespaceTypeInfo(_env)
    var warn bool
    warn = false
    if curNs.FP.Get_asyncMode(_env) == Ast_Async__Async{
        if Lns_op_not(Ast_isPrimitive(_env, symbolInfo.FP.Get_typeInfo(_env))){
            warn = true
        }
    }
    if warn{
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( symbolInfo.FP.Get_staticFlag(_env)) &&
                    _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) ).(bool))) ||
                _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Var) ).(bool))) &&
            _env.SetStackVal( Lns_op_not(self.builtinFunc.FP.Get_allSymbolSet(_env).Has(Ast_SymbolInfo2Stem(symbolInfo.FP.GetOrg(_env))))) &&
            _env.SetStackVal( Lns_op_not(self.FP.canBeAsyncParam(_env, symbolInfo.FP.Get_typeInfo(_env)))) ).(bool)){
            self.FP.AddErrMess(_env, pos, _env.GetVM().String_format("can't access the mutable symbol(%s) from async (%s).", []LnsAny{symbolInfo.FP.Get_name(_env), curNs.FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
}
// 2413: decl @lune.@base.@TransUnit.TransUnit.checkSymbolComp
func (self *TransUnit_TransUnit) checkSymbolComp(_env *LnsEnv, token *Types_Token) {
    self.FP.checkComp(_env, token, TransUnit_checkCompForm_31_(func(_env *LnsEnv, jsonWriter *Writer_JSON,prefix string) {
        self.FP.dumpSymbolComp(_env, jsonWriter.FP, self.FP.Get_scopeRO(_env), _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( prefix == "") &&
            _env.SetStackVal( "") ||
            _env.SetStackVal( "^" + prefix) ).(string))
    }))
}
// 2424: decl @lune.@base.@TransUnit.TransUnit.analyzeExpField
func (self *TransUnit_TransUnit) analyzeExpField(_env *LnsEnv, firstToken *Types_Token,fieldToken *Types_Token,mode LnsInt,prefixExp *Nodes_Node) *Nodes_Node {
    if prefixExp.FP.Get_expTypeList(_env).Len() > 1{
        prefixExp = &Nodes_ExpMultiTo1Node_create(_env, self.NodeManager, prefixExp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(prefixExp.FP.Get_expType(_env))}), prefixExp).Nodes_Node
    }
    var accessNil bool
    accessNil = false
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__FieldNil) ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__GetNil) ).(bool){
        accessNil = true
        if Lns_op_not(prefixExp.FP.Get_expType(_env).FP.Get_nilable(_env)){
            self.FP.addWarnMess(_env, prefixExp.FP.Get_pos(_env), _env.GetVM().String_format("This is not nilable. -- %s", []LnsAny{prefixExp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    if self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.IsAnalyzingSymArg(_env){
        if accessNil{
            self.HelperInfo.UseNilAccess = true
        }
        return &Nodes_RefFieldNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeSymbol)}), fieldToken, nil, accessNil, prefixExp).Nodes_Node
    }
    var typeInfo *Ast_TypeInfo
    typeInfo = Ast_builtinTypeStem_
    var prefixExpType *Ast_TypeInfo
    prefixExpType = prefixExp.FP.Get_expType(_env)
    if accessNil{
        if prefixExpType.FP.Get_nilable(_env){
            prefixExpType = prefixExpType.FP.Get_nonnilableType(_env)
            if prefixExpType.FP.Get_srcTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Box{
                self.FP.AddErrMess(_env, prefixExp.FP.Get_pos(_env), "Nilable can't support '$.' access yet")
            }
        } else { 
            accessNil = false
        }
    }
    var extFlag bool
    if Ast_isExtType(_env, prefixExpType){
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.IsAnalyzingBlockArg(_env))) &&
            _env.SetStackVal( Lns_op_not(self.FP.Get_curNsInfo(_env).FP.CanAccessLuaval(_env))) ).(bool)){
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("can't access Luaval without __luago. -- %s", []LnsAny{prefixExp.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
        }
        extFlag = true
        prefixExpType = prefixExpType.FP.Get_extedType(_env)
    } else { 
        extFlag = false
    }
    self.FP.checkFieldComp(_env, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__Get) ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__GetNil) ).(bool), fieldToken, prefixExp)
    if accessNil{
        self.HelperInfo.UseNilAccess = true
        if _switch0 := prefixExpType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Set || _switch0 == Ast_TypeInfoKind__Enum || _switch0 == Ast_TypeInfoKind__Alge || _switch0 == Ast_TypeInfoKind__Box {
            self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("%s does not support $.", []LnsAny{prefixExpType.FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    var prefixSymbolInfoList *LnsList
    prefixSymbolInfoList = prefixExp.FP.GetSymbolInfo(_env)
    self.FP.checkSymbolHavingValue(_env, prefixExp.FP.Get_pos(_env), prefixSymbolInfoList)
    var getterTypeInfo LnsAny
    getterTypeInfo = nil
    var symbolInfo LnsAny
    symbolInfo = nil
    if _switch1 := prefixExpType.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Class || _switch1 == Ast_TypeInfoKind__Module || _switch1 == Ast_TypeInfoKind__ExtModule || _switch1 == Ast_TypeInfoKind__IF || _switch1 == Ast_TypeInfoKind__List || _switch1 == Ast_TypeInfoKind__Array || _switch1 == Ast_TypeInfoKind__Set || _switch1 == Ast_TypeInfoKind__Box || _switch1 == Ast_TypeInfoKind__Alternate {
        var getterFlag bool
        getterFlag = false
        typeInfo, symbolInfo, getterFlag = self.FP.analyzeAccessClassField(_env, prefixExpType, mode, fieldToken)
        if getterFlag{
            {
                __exp := symbolInfo
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Ast_SymbolInfo)
                    getterTypeInfo = _exp.FP.Get_typeInfo(_env)
                    self.FP.checkNoasyncType(_env, fieldToken.Pos, _exp.FP.Get_typeInfo(_env))
                }
            }
        }
    } else if _switch1 == Ast_TypeInfoKind__Enum || _switch1 == Ast_TypeInfoKind__Alge {
        var scope *Ast_Scope
        scope = Lns_unwrap( prefixExpType.FP.Get_scope(_env)).(*Ast_Scope)
        var fieldName string
        fieldName = fieldToken.Txt
        var symbolInfoList *LnsList
        symbolInfoList = prefixExp.FP.GetSymbolInfo(_env)
        var isTypeSymbol bool
        isTypeSymbol = false
        if symbolInfoList.Len() > 0{
            if symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_kind(_env) == Ast_SymbolKind__Typ{
                isTypeSymbol = true
            }
        }
        if mode == TransUnit_ExpSymbolMode__Get{
            var moduleType *Ast_TypeInfo
            moduleType = prefixExpType.FP.GetModule(_env)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(moduleType.FP.Equals(_env, self.ProcessInfo, self.moduleType, nil, nil))) &&
                _env.SetStackVal( Lns_op_not(self.FP.Get_scope(_env).FP.GetModuleInfo(_env, moduleType))) ).(bool)){
                if Lns_op_not(self.ImportedAliasMap.Get(prefixExpType)){
                    self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("need to import module -- %s", []LnsAny{prefixExpType.FP.GetModule(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
            fieldName = "get_" + fieldName
            {
                _funcSymbol := scope.FP.GetSymbolInfoChild(_env, fieldName)
                if !Lns_IsNil( _funcSymbol ) {
                    funcSymbol := _funcSymbol.(*Ast_SymbolInfo)
                    symbolInfo = funcSymbol
                    var funcType *Ast_TypeInfo
                    funcType = funcSymbol.FP.Get_typeInfo(_env)
                    if funcType.FP.Get_staticFlag(_env) != isTypeSymbol{
                        self.FP.AddErrMess(_env, prefixExp.FP.Get_pos(_env), _env.GetVM().String_format("Can't access -- %s, %s", []LnsAny{fieldName, isTypeSymbol}))
                    }
                    var retTypeList *LnsList
                    retTypeList = funcType.FP.Get_retTypeInfoList(_env)
                    if retTypeList.Len() == 0{
                        self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("The func (%s) doesn't return value.", []LnsAny{funcType.FP.GetTxt(_env, nil, nil, nil)}))
                    } else { 
                        typeInfo = retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    }
                } else {
                    self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("not found -- %s.", []LnsAny{fieldName}))
                    typeInfo = Ast_builtinTypeNone
                }
            }
            getterTypeInfo = Ast_headTypeInfo
        } else { 
            {
                __exp := scope.FP.GetTypeInfoChild(_env, fieldName)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Ast_TypeInfo)
                    typeInfo = _exp
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) ||
                        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Alge) ).(bool){
                        if Lns_op_not(isTypeSymbol){
                            self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("can't access field -- %s", []LnsAny{fieldToken.Txt}))
                        }
                    }
                } else {
                    self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("not found field -- %s", []LnsAny{fieldToken.Txt}))
                    typeInfo = Ast_builtinTypeInt
                }
            }
        }
    } else if _switch1 == Ast_TypeInfoKind__Map {
        var work *Ast_TypeInfo
        work = prefixExpType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(work.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeString, nil, nil)){
            self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("map key type is not str. (%s)", []LnsAny{work.FP.GetTxt(_env, nil, nil, nil)}))
        }
        typeInfo = prefixExpType.FP.Get_itemTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(typeInfo.FP.Get_nilable(_env)){
            typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
        }
        if extFlag{
            typeInfo = self.FP.createExtType(_env, fieldToken.Pos, typeInfo)
        }
        return &Nodes_ExpRefItemNode_create(_env, self.NodeManager, fieldToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), prefixExp, accessNil, fieldToken.Txt, nil).Nodes_Node
    } else {
        if prefixExpType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeStem, nil, nil){
            typeInfo = Ast_builtinTypeStem_
            if extFlag{
                typeInfo = self.FP.createExtType(_env, fieldToken.Pos, typeInfo)
            }
            return &Nodes_ExpRefItemNode_create(_env, self.NodeManager, fieldToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), prefixExp, accessNil, fieldToken.Txt, nil).Nodes_Node
        } else { 
            self.FP.Error(_env, _env.GetVM().String_format("illegal type -- %s, %s", []LnsAny{prefixExpType.FP.GetTxt(_env, nil, nil, nil), Ast_TypeInfoKind_getTxt( prefixExpType.FP.Get_kind(_env))}))
        }
    }
    if Lns_op_not(symbolInfo){
        var prefixScope LnsAny
        prefixScope = prefixExpType.FP.Get_scope(_env)
        {
            __exp := prefixScope
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Ast_Scope)
                symbolInfo = _exp.FP.GetSymbolInfoField(_env, fieldToken.Txt, true, self.FP.Get_scope(_env), self.ScopeAccess)
            }
        }
    }
    if symbolInfo != nil{
        symbolInfo_740 := symbolInfo.(*Ast_SymbolInfo)
        if prefixSymbolInfoList.Len() == 1{
            var prefixSymbolInfo *Ast_SymbolInfo
            prefixSymbolInfo = prefixSymbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if prefixSymbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Typ{
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( prefixSymbolInfo.FP.Get_typeInfo(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Module) &&
                    _env.SetStackVal( Lns_op_not(symbolInfo_740.FP.Get_staticFlag(_env))) &&
                    _env.SetStackVal( symbolInfo_740.FP.Get_kind(_env) != Ast_SymbolKind__Typ) ).(bool)){
                    self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("Type can't access this symbol. -- %s", []LnsAny{symbolInfo_740.FP.Get_name(_env)}))
                }
            } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( symbolInfo_740.FP.Get_staticFlag(_env)) &&
                _env.SetStackVal( symbolInfo_740.FP.Get_typeInfo(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Method) ).(bool)){
                self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("can't access this symbol. -- %s", []LnsAny{fieldToken.Txt}))
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(Ast_TypeInfo_isMut(_env, prefixExpType))) &&
            _env.SetStackVal( Lns_op_not(symbolInfo_740.FP.Get_staticFlag(_env))) &&
            _env.SetStackVal( symbolInfo_740.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) &&
            _env.SetStackVal( symbolInfo_740.FP.Get_mutMode(_env) == Ast_MutMode__Mut) ).(bool)){
            self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("can't access mutable method. -- %s.%s", []LnsAny{prefixExpType.FP.GetTxt(_env, nil, nil, nil), fieldToken.Txt}))
        }
        if symbolInfo_740.FP.Get_typeInfo(_env).FP.Get_mutMode(_env) == Ast_MutMode__AllMut{
            var curType *Ast_TypeInfo
            curType = self.FP.GetCurrentNamespaceTypeInfo(_env)
            if _switch3 := curType.FP.Get_kind(_env); _switch3 == Ast_TypeInfoKind__Func || _switch3 == Ast_TypeInfoKind__Method {
                if _switch2 := curType.FP.Get_asyncMode(_env); _switch2 == Ast_Async__Async || _switch2 == Ast_Async__Transient {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( curType.FP.Get_kind(_env) != Ast_TypeInfoKind__Method) ||
                        _env.SetStackVal( curType.FP.Get_rawTxt(_env) != "__init") ).(bool){
                        self.FP.AddErrMess(_env, fieldToken.Pos, _env.GetVM().String_format("can't access allmut type's field(%s) in async function.", []LnsAny{symbolInfo_740.FP.Get_name(_env)}))
                    }
                }
            }
        }
    }
    var accessSymbolInfo LnsAny
    accessSymbolInfo = nil
    var symbolMutMode LnsInt
    symbolMutMode = typeInfo.FP.Get_mutMode(_env)
    if symbolInfo != nil{
        symbolInfo_755 := symbolInfo.(*Ast_SymbolInfo)
        var workSymInfo *Ast_AccessSymbolInfo
        workSymInfo = NewAst_AccessSymbolInfo(_env, self.ProcessInfo, symbolInfo_755, &Ast_OverrideMut__Prefix{prefixExpType}, Lns_op_not(accessNil))
        if Lns_op_not(getterTypeInfo){
            typeInfo = workSymInfo.FP.Get_typeInfo(_env)
        }
        accessSymbolInfo = workSymInfo
        if _switch4 := mode; _switch4 == TransUnit_ExpSymbolMode__Field || _switch4 == TransUnit_ExpSymbolMode__FieldNil {
            symbolMutMode = symbolInfo_755.FP.Get_mutMode(_env)
        }
    }
    if accessNil{
        if Lns_op_not(typeInfo.FP.Get_nilable(_env)){
            typeInfo = typeInfo.FP.Get_nilableTypeInfo(_env)
        }
        self.HelperInfo.UseNilAccess = true
    }
    if Lns_op_not(Ast_TypeInfo_isMut(_env, prefixExpType)){
        if self.Ctrl_info.LegacyMutableControl{
            if symbolMutMode == Ast_MutMode__Mut{
                typeInfo = self.FP.createModifier(_env, typeInfo, Ast_MutMode__IMut)
            }
        } else { 
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( typeInfo.FP.Get_mutMode(_env) == Ast_MutMode__Mut) &&
                _env.SetStackVal( typeInfo.FP.Get_mutMode(_env) != Ast_MutMode__AllMut) ).(bool)){
                typeInfo = self.FP.createModifier(_env, typeInfo, Ast_MutMode__IMut)
            }
        }
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.List_unpack, nil, nil)) ||
        _env.SetStackVal( typeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.Array_unpack, nil, nil)) ).(bool){
        self.HelperInfo.UseUnpack = true
    }
    if typeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.Str_replace, nil, nil){
        self.HelperInfo.UseStrReplace = true
    }
    {
        _expRef := Nodes_ExpRefNodeDownCastF(prefixExp.FP)
        if !Lns_IsNil( _expRef ) {
            expRef := _expRef.(*Nodes_ExpRefNode)
            var prefixSym *Ast_SymbolInfo
            prefixSym = expRef.FP.Get_symbolInfo(_env)
            var prefixType *Ast_TypeInfo
            prefixType = prefixSym.FP.Get_typeInfo(_env)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( prefixSym.FP.Get_kind(_env) == Ast_SymbolKind__Typ) &&
                _env.SetStackVal( prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
                _env.SetStackVal( prefixType.FP.Get_itemTypeInfoList(_env).Len() > 0) &&
                _env.SetStackVal( Lns_op_not(Ast_isGenericType(_env, prefixType))) &&
                _env.SetStackVal( Lns_op_not(self.FP.Get_scope(_env).FP.IsInnerOf(_env, Lns_unwrap( prefixType.FP.Get_scope(_env)).(*Ast_Scope)))) ).(bool)){
                var accessErr bool
                accessErr = false
                if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Func{
                    var altSet *LnsSet
                    altSet = NewLnsSet([]LnsAny{})
                    for _, _argType := range( typeInfo.FP.Get_argTypeInfoList(_env).Items ) {
                        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        var orgType *Ast_TypeInfo
                        orgType = argType.FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env)
                        if orgType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                            altSet.Add(Ast_TypeInfo2Stem(orgType))
                        }
                    }
                    for _, _itemType := range( prefixType.FP.Get_itemTypeInfoList(_env).Items ) {
                        itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        if Lns_op_not(altSet.Has(Ast_TypeInfo2Stem(itemType.FP.Get_nonnilableType(_env).FP.Get_srcTypeInfo(_env)))){
                            accessErr = true
                            break
                        }
                    }
                } else { 
                    accessErr = true
                }
                if accessErr{
                    self.FP.AddErrMess(_env, prefixExp.FP.Get_pos(_env), _env.GetVM().String_format("can't access this class(%s) without '<>'.", []LnsAny{prefixType.FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        }
    }
    if extFlag{
        typeInfo = self.FP.createExtType(_env, firstToken.Pos, typeInfo)
    }
    {
        __exp := getterTypeInfo
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_TypeInfo)
            return &Nodes_GetFieldNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), fieldToken, Ast_SymbolInfoDownCastF(accessSymbolInfo), accessNil, prefixExp, _exp).Nodes_Node
        } else {
            if accessSymbolInfo != nil{
                accessSymbolInfo_790 := accessSymbolInfo.(*Ast_AccessSymbolInfo)
                self.FP.checkAsyncField(_env, &accessSymbolInfo_790.Ast_SymbolInfo, fieldToken.Pos)
            }
            return &Nodes_RefFieldNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), fieldToken, Ast_SymbolInfoDownCastF(accessSymbolInfo), accessNil, prefixExp).Nodes_Node
        }
    }
// insert a dummy
    return nil
}
// 2800: decl @lune.@base.@TransUnit.TransUnit.analyzeNewAlge
func (self *TransUnit_TransUnit) analyzeNewAlge(_env *LnsEnv, firstToken *Types_Token,algeTypeInfo *Ast_AlgeTypeInfo,prefix LnsAny) *Nodes_NewAlgeValNode {
    var symbolToken *Types_Token
    symbolToken = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
    self.FP.checkAlgeComp(_env, symbolToken, algeTypeInfo)
    {
        _valInfo := algeTypeInfo.FP.GetValInfo(_env, symbolToken.Txt)
        if !Lns_IsNil( _valInfo ) {
            valInfo := _valInfo.(*Ast_AlgeValInfo)
            var argList *LnsList
            argList = NewLnsList([]LnsAny{})
            var argListNode LnsAny
            if valInfo.FP.Get_typeList(_env).Len() > 0{
                self.FP.checkNextToken(_env, "(")
                argListNode = self.FP.analyzeExpList(_env, false, false, false, nil, valInfo.FP.Get_typeList(_env), nil)
                argList = (Lns_unwrap( argListNode).(*Nodes_ExpListNode)).FP.Get_expList(_env)
                self.FP.checkNextToken(_env, ")")
            } else { 
                argListNode = nil
            }
            {
                _, _, _newExpNodeList := TransUnit_convExp0_5096(Lns_2DDD(self.FP.checkMatchType(_env, "call", symbolToken.Pos, valInfo.FP.Get_typeList(_env), argListNode, false, true, nil, true)))
                if !Lns_IsNil( _newExpNodeList ) {
                    newExpNodeList := _newExpNodeList.(*Nodes_ExpListNode)
                    argList = newExpNodeList.FP.Get_expList(_env)
                }
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( algeTypeInfo.FP.Get_externalFlag(_env)) &&
                _env.SetStackVal( Lns_op_not(self.FP.Get_scope(_env).FP.GetModuleInfo(_env, algeTypeInfo.FP.GetModule(_env).FP.Get_srcTypeInfo(_env)))) ).(bool)){
                var fullname string
                fullname = algeTypeInfo.FP.GetFullName(_env, self.TypeNameCtrl, self.FP.Get_scope(_env).FP, true)
                self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("This module not import -- %s", []LnsAny{fullname}))
            }
            return Nodes_NewAlgeValNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_AlgeTypeInfo2Stem(algeTypeInfo)}), symbolToken, prefix, algeTypeInfo, valInfo, argList)
        } else {
            var dummySymbol LnsAny
            dummySymbol = _env.NilAccFin(_env.NilAccPush(algeTypeInfo.FP.Get_parentInfo(_env).FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(_env, algeTypeInfo.FP.Get_rawTxt(_env))})/* 2844:10 */)
            self.FP.AddErrMess(_env, symbolToken.Pos, _env.GetVM().String_format("not found Alge -- %s", []LnsAny{symbolToken.Txt}))
            return Nodes_NewAlgeValNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_AlgeTypeInfo2Stem(algeTypeInfo)}), symbolToken, prefix, algeTypeInfo, NewAst_AlgeValInfo(_env, "", NewLnsList([]LnsAny{}), algeTypeInfo, Lns_unwrap( dummySymbol).(*Ast_SymbolInfo)), NewLnsList([]LnsAny{}))
        }
    }
// insert a dummy
    return nil
}
// 2854: decl @lune.@base.@TransUnit.TransUnit.analyzeExpSymbol
func (self *TransUnit_TransUnit) analyzeExpSymbol(_env *LnsEnv, firstToken *Types_Token,symbolToken *Types_Token,mode LnsInt,prefixExp LnsAny,skipFlag bool,canLeftExp bool) *Nodes_Node {
    var exp *Nodes_Node
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__Field) ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__Get) ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__FieldNil) ||
        _env.SetStackVal( mode == TransUnit_ExpSymbolMode__GetNil) ).(bool){
        if prefixExp != nil{
            prefixExp_1543 := prefixExp.(*Nodes_Node)
            exp = self.FP.analyzeExpField(_env, firstToken, symbolToken, mode, prefixExp_1543)
            var expType *Ast_TypeInfo
            expType = exp.FP.Get_expType(_env)
            if prefixExp_1543.FP.Get_expType(_env).FP.IsModule(_env){
                {
                    _algeType := Ast_AlgeTypeInfoDownCastF(expType.FP)
                    if !Lns_IsNil( _algeType ) {
                        algeType := _algeType.(*Ast_AlgeTypeInfo)
                        var nextToken *Types_Token
                        nextToken = self.FP.getToken(_env, nil)
                        if nextToken.Txt == "."{
                            return &self.FP.analyzeNewAlge(_env, firstToken, algeType, exp).Nodes_Node
                        }
                        self.FP.Pushback(_env)
                    }
                }
            }
        } else {
            Util_err(_env, "prefixExp is nil")
        }
    } else if mode == TransUnit_ExpSymbolMode__Symbol{
        if self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.IsAnalyzingSymArg(_env){
            exp = &Nodes_LiteralSymbolNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeSymbol)}), symbolToken).Nodes_Node
        } else { 
            self.FP.checkSymbolComp(_env, symbolToken)
            var symbolInfo *Ast_SymbolInfo
            
            {
                _symbolInfo := self.FP.Get_scope(_env).FP.GetSymbolTypeInfo(_env, symbolToken.Txt, self.FP.Get_scope(_env), self.ModuleScope, self.ScopeAccess)
                if _symbolInfo == nil{
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( self.MacroCtrl.FP.Get_isDeclaringMacro(_env)) &&
                        _env.SetStackVal( symbolToken.Txt == "__var") &&
                        _env.SetStackVal( self.analyzePhase == TransUnit_AnalyzePhase__Runner) ).(bool)){
                        self.FP.Error(_env, "declare '_lune_control use_macro_special_var'.")
                    } else { 
                        if self.analyzeMode != TransUnit_AnalyzeMode__Diag{
                            var work *Ast_Scope
                            work = self.FP.Get_scope(_env)
                            for  {
                                Lns_print([]LnsAny{work, self.ModuleScope})
                                if work == work.FP.Get_parent(_env){
                                    break
                                }
                                work = work.FP.Get_parent(_env)
                            }
                            self.FP.Get_scope(_env).FP.FilterSymbolTypeInfo(_env, self.FP.Get_scope(_env), self.ModuleScope, self.ScopeAccess, Ast_filterForm(TransUnit_analyzeExpSymbol___anonymous_0_))
                        }
                        self.FP.Error(_env, "not found type -- " + symbolToken.Txt)
                    }
                } else {
                    symbolInfo = _symbolInfo.(*Ast_SymbolInfo)
                }
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.MacroCtrl.FP.Get_isDeclaringMacro(_env)) &&
                _env.SetStackVal( symbolToken.Txt == "__var") ).(bool)){
                self.MacroCtrl.FP.SetUsing__var(_env)
            }
            self.FP.accessSymbol(_env, symbolInfo, canLeftExp)
            self.FP.checkAsyncSymbol(_env, symbolInfo, firstToken.Pos)
            var typeInfo *Ast_TypeInfo
            typeInfo = symbolInfo.FP.Get_typeInfo(_env)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Typ) &&
                _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                _env.SetStackVal( Ast_Scope2Stem(self.analyzingStaticMethodArgsScope) != nil) &&
                _env.SetStackVal( symbolInfo.FP.Get_scope(_env) != self.analyzingStaticMethodArgsScope) ).(bool)){
                self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("The static method can't use the type parameter. -- %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            }
            if _switch0 := symbolInfo.FP.Get_kind(_env); _switch0 == Ast_SymbolKind__Typ {
                {
                    _algeType := Ast_AlgeTypeInfoDownCastF(typeInfo.FP)
                    if !Lns_IsNil( _algeType ) {
                        algeType := _algeType.(*Ast_AlgeTypeInfo)
                        var nextToken *Types_Token
                        nextToken = self.FP.getToken(_env, nil)
                        if nextToken.Txt == "."{
                            return &self.FP.analyzeNewAlge(_env, firstToken, algeType, nil).Nodes_Node
                        }
                        self.FP.Pushback(_env)
                    }
                }
            } else if _switch0 == Ast_SymbolKind__Var {
                self.tentativeSymbol.FP.AddAccessSym(_env, symbolInfo)
                if Lns_op_not(symbolInfo.FP.Get_hasValueFlag(_env)){
                    var nsTypeInfo *Ast_TypeInfo
                    nsTypeInfo = self.FP.GetCurrentNamespaceTypeInfo(_env)
                    if Lns_op_not(symbolInfo.FP.Get_scope(_env).FP.IsInnerOf(_env, Lns_unwrap( nsTypeInfo.FP.Get_scope(_env)).(*Ast_Scope))){
                        self.tentativeSymbol.FP.AddAccessSymPos(_env, NewTransUnit_AccessSymPos(_env, symbolInfo, firstToken.Pos))
                    }
                }
            }
            if typeInfo.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeSymbol, nil, nil){
                skipFlag = true
            }
            if typeInfo.FP.Equals(_env, self.ProcessInfo, self.builtinFunc.Lns__load, nil, nil){
                self.HelperInfo.UseLoad = true
            }
            if _switch1 := symbolToken.Txt; _switch1 == "__func__" {
                var funcTypeInfo *Ast_TypeInfo
                funcTypeInfo = self.FP.GetCurrentNamespaceTypeInfo(_env)
                self.Has__func__Symbol.Add(Ast_TypeInfo2Stem(funcTypeInfo))
            } else if _switch1 == "_G" || _switch1 == "_ENV" {
                var valid bool
                valid = false
                for _pragma := range( self.HelperInfo.PragmaSet.Items ) {
                    pragma := _pragma
                    switch _matchExp0 := pragma.(type) {
                    case *LuneControl_Pragma__limit_conv_code:
                    codeSet := _matchExp0.Val1
                        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( codeSet.Len() == 1) &&
                            _env.SetStackVal( codeSet.Has(LuneControl_Code__Lua)) ).(bool)){
                            valid = true
                            break
                        }
                    }
                }
                if Lns_op_not(valid){
                    self.FP.AddErrMess(_env, firstToken.Pos, "'_G' and '_ENV' only can access with transcompiling to lua.")
                }
            } else if _switch1 == "_" {
                if Lns_op_not(canLeftExp){
                    self.FP.AddErrMess(_env, firstToken.Pos, "It can't access the symbol '_'.")
                }
            }
            exp = &Nodes_ExpRefNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), &NewAst_AccessSymbolInfo(_env, self.ProcessInfo, symbolInfo, Ast_OverrideMut__None_Obj, true).Ast_SymbolInfo).Nodes_Node
        }
    } else if mode == TransUnit_ExpSymbolMode__Fn{
        exp = self.FP.analyzeDeclFunc(_env, TransUnit_DeclFuncMode__Func, false, false, false, Ast_AccessMode__Local, false, nil, symbolToken, nil)
    } else { 
        self.FP.Error(_env, _env.GetVM().String_format("illegal mode -- %s", []LnsAny{mode}))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.analyzeMode == TransUnit_AnalyzeMode__Inquire) &&
        _env.SetStackVal( self.FP.isTargetToken(_env, symbolToken)) ).(bool)){
        var accessMode LnsInt
        accessMode = Ast_AccessMode__None
        for _, _symbolInfo := range( exp.FP.GetSymbolInfo(_env).Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            accessMode = symbolInfo.FP.Get_accessMode(_env)
        }
        self.FP.dumpSymbolType(_env, accessMode, symbolToken.Txt, exp.FP.Get_expType(_env))
    }
    return self.FP.analyzeExpCont(_env, firstToken, exp, skipFlag, canLeftExp)
}
// 3034: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOpSet
func (self *TransUnit_TransUnit) analyzeExpOpSet(_env *LnsEnv, exp *Nodes_Node,opeToken *Types_Token,expectTypeList *LnsList) *Nodes_Node {
    exp.FP.SetLValue(_env)
    if Lns_op_not(exp.FP.CanBeLeft(_env)){
        self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("this node can not be l-value. -- %s", []LnsAny{Nodes_getNodeKindName(_env, exp.FP.Get_kind(_env))}))
    }
    var listRefItemNode LnsAny
    listRefItemNode = nil
    {
        _symNodeList := Nodes_ExpListNodeDownCastF(exp.FP)
        if !Lns_IsNil( _symNodeList ) {
            symNodeList := _symNodeList.(*Nodes_ExpListNode)
            for _, _symNode := range( symNodeList.FP.Get_expList(_env).Items ) {
                symNode := _symNode.(Nodes_NodeDownCast).ToNodes_Node()
                {
                    _refItemNode := TransUnit_analyzeExpOpSet__process_0_(_env, symNode)
                    if !Lns_IsNil( _refItemNode ) {
                        refItemNode := _refItemNode.(*Nodes_ExpRefItemNode)
                        listRefItemNode = refItemNode
                        if symNodeList.FP.Get_expList(_env).Len() > 1{
                            self.FP.AddErrMess(_env, refItemNode.FP.Get_pos(_env), "When left-value includes 'list[i]', left-value must be single.")
                        }
                        break
                    }
                }
            }
        } else {
            listRefItemNode = TransUnit_analyzeExpOpSet__process_0_(_env, exp)
        }
    }
    var expList *Nodes_ExpListNode
    expList = self.FP.analyzeExpList(_env, false, false, false, nil, expectTypeList, nil)
    for _index, _expType := range( expList.FP.Get_expTypeList(_env).Items ) {
        index := _index + 1
        expType := _expType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if expType.FP.Get_asyncMode(_env) == Ast_Async__Transient{
            self.FP.AddErrMess(_env, expList.FP.Get_pos(_env), _env.GetVM().String_format("can't set the __trans type -- index:%d, %s", []LnsAny{index, expType.FP.GetTxt(_env, nil, nil, nil)}))
        }
    }
    if Lns_op_not(expList.FP.CanBeRight(_env, self.ProcessInfo)){
        self.FP.AddErrMess(_env, expList.FP.Get_effectivePos(_env), _env.GetVM().String_format("this node can not be r-value. -- %s", []LnsAny{Nodes_getNodeKindName(_env, expList.FP.Get_kind(_env))}))
    }
    var workList LnsAny
    var expTypeList *LnsList
    _,_,workList,expTypeList = self.FP.checkMatchType(_env, "= operator", opeToken.Pos, exp.FP.Get_expTypeList(_env), expList, true, false, nil, true)
    if workList != nil{
        workList_1628 := workList.(*Nodes_ExpListNode)
        expList = workList_1628
    }
    var initSymSet *LnsSet
    initSymSet = NewLnsSet([]LnsAny{})
    var symbolList *LnsList
    symbolList = NewLnsList([]LnsAny{})
    for _index, _symbolInfo := range( exp.FP.GetSymbolInfo(_env).Items ) {
        index := _index + 1
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        symbolList.Insert(Ast_SymbolInfo2Stem(symbolInfo))
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(symbolInfo.FP.Get_mutable(_env))) &&
            _env.SetStackVal( symbolInfo.FP.Get_hasValueFlag(_env)) ).(bool)){
            if self.ValidMutControl{
                self.FP.AddErrMess(_env, opeToken.Pos, _env.GetVM().String_format("this is not mutable variable. -- %s", []LnsAny{symbolInfo.FP.Get_name(_env)}))
            }
        }
        self.tentativeSymbol.FP.ModSym(_env, symbolInfo)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( index <= expTypeList.Len()) &&
            _env.SetStackVal( Lns_op_not(symbolInfo.FP.Get_hasValueFlag(_env))) ).(bool)){
            if _switch1 := symbolInfo.FP.Get_kind(_env); _switch1 == Ast_SymbolKind__Var {
                if symbolInfo.FP.Get_typeInfo(_env) == Ast_builtinTypeEmpty{
                    var expType *Ast_TypeInfo
                    expType = expTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if _switch0 := expType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__DDD {
                        if expType.FP.Get_itemTypeInfoList(_env).Len() > 0{
                            expType = expType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nilableTypeInfo(_env)
                        }
                    } else if _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array || _switch0 == Ast_TypeInfoKind__Set || _switch0 == Ast_TypeInfoKind__Map {
                        var workPos Types_Position
                        if index <= expList.FP.Get_expList(_env).Len(){
                            workPos = expList.FP.Get_expList(_env).GetAt(index).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_pos(_env)
                        } else { 
                            workPos = opeToken.Pos
                        }
                        self.FP.checkLiteralEmptyCollection(_env, workPos, symbolInfo.FP.Get_name(_env), expType)
                    }
                    symbolInfo.FP.Set_typeInfo(_env, expType)
                }
                if Lns_op_not(self.tentativeSymbol.FP.Regist(_env, symbolInfo, exp.FP.Get_pos(_env))){
                    self.FP.AddErrMess(_env, opeToken.Pos, _env.GetVM().String_format("can't access in this scope. -- %s", []LnsAny{symbolInfo.FP.Get_name(_env)}))
                }
                initSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
            } else if _switch1 == Ast_SymbolKind__Mbr {
                initSymSet.Add(Ast_SymbolInfo2Stem(symbolInfo))
            }
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( symbolInfo.FP.Get_kind(_env) != Ast_SymbolKind__Var) ||
            _env.SetStackVal( self.FP.Get_scope(_env).FP.GetNamespaceTypeInfo(_env) == symbolInfo.FP.Get_scope(_env).FP.GetNamespaceTypeInfo(_env)) ).(bool){
            symbolInfo.FP.UpdateValue(_env, exp.FP.Get_pos(_env))
        }
    }
    if listRefItemNode != nil{
        listRefItemNode_1650 := listRefItemNode.(*Nodes_ExpRefItemNode)
        var index LnsAny
        {
            _indexNode := listRefItemNode_1650.FP.Get_index(_env)
            if !Lns_IsNil( _indexNode ) {
                indexNode := _indexNode.(*Nodes_Node)
                index = &Nodes_IndexVal__NodeIdx{indexNode}
            } else {
                index = &Nodes_IndexVal__SymIdx{Lns_unwrap( listRefItemNode_1650.FP.Get_symbol(_env)).(string)}
            }
        }
        return &Nodes_ExpSetItemNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), listRefItemNode_1650.FP.Get_val(_env), index, &expList.Nodes_Node).Nodes_Node
    }
    return &Nodes_ExpSetValNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), exp, expList, symbolList, initSymSet).Nodes_Node
}
// 3182: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOpEquals
func (self *TransUnit_TransUnit) analyzeExpOpEquals(_env *LnsEnv, pos Types_Position,opToken *Types_Token,exp1 *Nodes_Node,exp2 *Nodes_Node)(*Nodes_Node, *Nodes_Node) {
    var exp1Type *Ast_TypeInfo
    exp1Type = exp1.FP.Get_expType(_env)
    var exp2Type *Ast_TypeInfo
    exp2Type = exp2.FP.Get_expType(_env)
    if Lns_isCondTrue( (_env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(Lns_car(exp1Type.FP.CanEvalWith(_env, self.ProcessInfo, exp2Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) &&
        _env.SetStackVal( Lns_op_not(Lns_car(exp2Type.FP.CanEvalWith(_env, self.ProcessInfo, exp1Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool))){
        self.FP.AddErrMess(_env, opToken.Pos, _env.GetVM().String_format("not compatible type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), exp2Type.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
        return exp1, exp2
    }
    {
        __exp := Nodes_NewAlgeValNodeDownCastF(exp1.FP)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_NewAlgeValNode)
            if _exp.FP.Get_paramList(_env).Len() > 0{
                self.FP.AddErrMess(_env, exp1.FP.Get_pos(_env), "can't compare alge.")
                return exp1, exp2
            }
        }
    }
    {
        __exp := Nodes_NewAlgeValNodeDownCastF(exp2.FP)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_NewAlgeValNode)
            if _exp.FP.Get_paramList(_env).Len() > 0{
                self.FP.AddErrMess(_env, exp2.FP.Get_pos(_env), "can't compare alge.")
                return exp1, exp2
            }
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeBool, nil, nil)) &&
        _env.SetStackVal( exp2Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeBool, nil, nil)) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( exp1.FP.Get_kind(_env) == Nodes_NodeKind_get_LiteralBool(_env)) ||
            _env.SetStackVal( exp2.FP.Get_kind(_env) == Nodes_NodeKind_get_LiteralBool(_env)) ).(bool))) ).(bool)){
        self.FP.addWarnMess(_env, exp1.FP.Get_pos(_env), "this operation is deprecated.")
        return exp1, exp2
    }
    var nonNilType1 *Ast_TypeInfo
    nonNilType1 = TransUnit_analyzeExpOpEquals__getType_0_(_env, exp1Type)
    var nonNilType2 *Ast_TypeInfo
    nonNilType2 = TransUnit_analyzeExpOpEquals__getType_0_(_env, exp2Type)
    if nonNilType1 != nonNilType2{
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( nonNilType1.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
            _env.SetStackVal( nonNilType1.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ).(bool){
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( nonNilType2.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( nonNilType2.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ).(bool){
                if nonNilType1.FP.IsInheritFrom(_env, self.ProcessInfo, nonNilType2, nil){
                    exp1 = &Nodes_ExpCastNode_create(_env, self.NodeManager, exp1.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp2Type.FP.Get_nonnilableType(_env))}), exp1, exp2Type, nil, "", Nodes_CastKind__Implicit).Nodes_Node
                } else { 
                    exp2 = &Nodes_ExpCastNode_create(_env, self.NodeManager, exp2.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp1Type.FP.Get_nonnilableType(_env))}), exp2, exp1Type, nil, "", Nodes_CastKind__Implicit).Nodes_Node
                }
            } else { 
                exp1 = &Nodes_ExpCastNode_create(_env, self.NodeManager, exp1.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp1Type)}), exp1, Ast_builtinTypeStem, nil, "", Nodes_CastKind__Implicit).Nodes_Node
            }
        } else { 
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( nonNilType2.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) ||
                _env.SetStackVal( nonNilType2.FP.Get_kind(_env) == Ast_TypeInfoKind__IF) ).(bool){
                exp2 = &Nodes_ExpCastNode_create(_env, self.NodeManager, exp2.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp2Type)}), exp2, Ast_builtinTypeStem, nil, "", Nodes_CastKind__Implicit).Nodes_Node
            }
        }
    }
    return exp1, exp2
}
// 3273: decl @lune.@base.@TransUnit.TransUnit.analyzeExpOp2
func (self *TransUnit_TransUnit) analyzeExpOp2(_env *LnsEnv, firstToken *Types_Token,exp *Nodes_Node,prevOpLevel LnsAny) *Nodes_Node {
    for  {
        var opToken *Types_Token
        opToken = self.FP.GetTokenNoErr(_env, nil)
        var opTxt string
        opTxt = opToken.Txt
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( opToken.Txt == "@@") ||
            _env.SetStackVal( opToken.Txt == "@@@") ||
            _env.SetStackVal( opToken.Txt == "@@=") ).(bool){
            exp = self.FP.analyzeExpCast(_env, firstToken, opTxt, exp)
        } else if opToken.Kind == Types_TokenKind__Ope{
            if Parser_isOp2(_env, opTxt){
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( opTxt != "=") &&
                    _env.SetStackVal( Lns_op_not(exp.FP.CanBeRight(_env, self.ProcessInfo))) ).(bool)){
                    self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("This can't evaluate for '%s' -- %s", []LnsAny{opTxt, Nodes_getNodeKindName(_env, exp.FP.Get_kind(_env))}))
                }
                var opLevel LnsInt
                
                {
                    _opLevel := TransUnit_op2levelMap.Get(opTxt)
                    if _opLevel == nil{
                        self.FP.Error(_env, _env.GetVM().String_format("unknown op -- %s %s", []LnsAny{opTxt, prevOpLevel}))
                    } else {
                        opLevel = _opLevel.(LnsInt)
                    }
                }
                {
                    __exp := prevOpLevel
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(LnsInt)
                        if opLevel <= _exp{
                            self.FP.Pushback(_env)
                            return exp
                        }
                    }
                }
                var expectTypeList *LnsList
                expectTypeList = NewLnsList([]LnsAny{})
                for _, _exp1Type := range( exp.FP.Get_expTypeList(_env).Items ) {
                    exp1Type := _exp1Type.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    var prefixExpType *Ast_TypeInfo
                    prefixExpType = exp1Type
                    if prefixExpType.FP.Get_nilable(_env){
                        prefixExpType = prefixExpType.FP.Get_nonnilableType(_env)
                    }
                    var expectType *Ast_TypeInfo
                    expectType = Ast_builtinTypeNone
                    {
                        __exp := Ast_EnumTypeInfoDownCastF(prefixExpType.FP.Get_srcTypeInfo(_env).FP.Get_aliasSrc(_env).FP)
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(*Ast_EnumTypeInfo)
                            expectType = &_exp.Ast_TypeInfo
                        }
                    }
                    {
                        __exp := Ast_AlgeTypeInfoDownCastF(prefixExpType.FP.Get_srcTypeInfo(_env).FP)
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(*Ast_AlgeTypeInfo)
                            expectType = &_exp.Ast_TypeInfo
                        }
                    }
                    expectTypeList.Insert(Ast_TypeInfo2Stem(expectType))
                }
                if expectTypeList.Len() == 0{
                    self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("This exp have no value -- %s", []LnsAny{Nodes_getNodeKindName(_env, exp.FP.Get_kind(_env))}))
                    expectTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeNone))
                }
                if opTxt == "="{
                    return self.FP.analyzeExpOpSet(_env, exp, opToken, expectTypeList)
                }
                exp = self.FP.MultiTo1(_env, exp)
                var exp2 *Nodes_Node
                exp2 = self.FP.analyzeExp(_env, false, false, false, opLevel, expectTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                exp2 = self.FP.MultiTo1(_env, exp2)
                if Lns_op_not(exp2.FP.CanBeRight(_env, self.ProcessInfo)){
                    self.FP.AddErrMess(_env, exp2.FP.Get_pos(_env), _env.GetVM().String_format("This can't evaluate for '%s' -- %s", []LnsAny{opTxt, Nodes_getNodeKindName(_env, exp2.FP.Get_kind(_env))}))
                }
                var retType *Ast_TypeInfo
                retType = Ast_builtinTypeNone
                if Lns_op_not(exp2.FP.CanBeRight(_env, self.ProcessInfo)){
                    self.FP.AddErrMess(_env, exp2.FP.Get_effectivePos(_env), _env.GetVM().String_format("this node can not be r-value. -- %s", []LnsAny{Nodes_getNodeKindName(_env, exp2.FP.Get_kind(_env))}))
                }
                var exp1Type *Ast_TypeInfo
                exp1Type = exp.FP.Get_expType(_env)
                var exp2Type *Ast_TypeInfo
                exp2Type = exp2.FP.Get_expType(_env)
                if exp1Type.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    {
                        _dddType := Ast_DDDTypeInfoDownCastF(exp1Type.FP)
                        if !Lns_IsNil( _dddType ) {
                            dddType := _dddType.(*Ast_DDDTypeInfo)
                            exp = &Nodes_ExpMultiTo1Node_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddType.FP.Get_typeInfo(_env).FP.Get_nilableTypeInfo(_env))}), exp).Nodes_Node
                        }
                    }
                }
                if exp2Type.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    {
                        _dddType := Ast_DDDTypeInfoDownCastF(exp2Type.FP)
                        if !Lns_IsNil( _dddType ) {
                            dddType := _dddType.(*Ast_DDDTypeInfo)
                            exp2 = &Nodes_ExpMultiTo1Node_create(_env, self.NodeManager, exp2.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(dddType.FP.Get_typeInfo(_env).FP.Get_nilableTypeInfo(_env))}), exp2).Nodes_Node
                        }
                    }
                }
                if _switch0 := opTxt; _switch0 == "or" {
                    var is3op bool
                    {
                        _opExpType := Ast_AndExpTypeInfoDownCastF(exp1Type.FP)
                        if !Lns_IsNil( _opExpType ) {
                            opExpType := _opExpType.(*Ast_AndExpTypeInfo)
                            exp1Type = opExpType.FP.Get_exp2(_env)
                            is3op = true
                        } else {
                            is3op = false
                        }
                    }
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeBool, nil, nil))) &&
                        _env.SetStackVal( Lns_op_not(exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeStem, nil, nil))) &&
                        _env.SetStackVal( Lns_op_not(exp1Type.FP.Get_nilable(_env))) ).(bool)){
                        if _env.NilAccFin(_env.NilAccPush((Nodes_ExpOp2NodeDownCastF(exp.FP))) && 
                        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_ExpOp2Node).FP.Get_op(_env)})&&
                        _env.NilAccPush(_env.NilAccPop().(*Types_Token).Txt)) == "and"{
                        } else { 
                            self.FP.addWarnMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("this value never be 'false' -- %s", []LnsAny{exp1Type.FP.GetTxt(_env, nil, nil, nil)}))
                        }
                    }
                    if exp1Type.FP.Equals(_env, self.ProcessInfo, exp2Type, nil, nil){
                        retType = exp1Type
                    } else if Lns_car(exp1Type.FP.CanEvalWith(_env, self.ProcessInfo, exp2Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
                        retType = exp1Type
                    } else if Lns_car(exp2Type.FP.CanEvalWith(_env, self.ProcessInfo, exp1Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
                        retType = exp2Type
                    } else if exp2Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeNil, nil, nil){
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( is3op) ||
                            _env.SetStackVal( exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeBool, nil, nil)) ).(bool){
                            retType = exp1Type.FP.Get_nilableTypeInfo(_env)
                        } else { 
                            retType = exp1Type
                        }
                    } else if exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeNil, nil, nil){
                        retType = exp2Type
                    } else { 
                        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( exp1Type.FP.Get_nilable(_env)) &&
                            _env.SetStackVal( exp2Type.FP.Get_nilable(_env)) ).(bool)){
                            retType = Ast_builtinTypeStem_
                        } else if exp2Type.FP.Get_nilable(_env){
                            retType = Ast_builtinTypeStem_
                        } else if exp1Type.FP.Get_nilable(_env){
                            retType = Ast_builtinTypeStem
                        } else { 
                            retType = Ast_builtinTypeStem
                        }
                    }
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( retType.FP.Get_nilable(_env)) &&
                        _env.SetStackVal( Lns_op_not(exp2Type.FP.Get_nilable(_env))) ).(bool)){
                        retType = retType.FP.Get_nonnilableType(_env)
                    }
                } else if _switch0 == "and" {
                    _ = self.FP.getToken(_env, nil)
                    self.FP.Pushback(_env)
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeBool, nil, nil))) &&
                        _env.SetStackVal( Lns_op_not(exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeStem, nil, nil))) &&
                        _env.SetStackVal( Lns_op_not(exp1Type.FP.Get_nilable(_env))) ).(bool)){
                        self.FP.addWarnMess(_env, exp.FP.Get_pos(_env), "this value never be 'false'")
                    } else if exp2.FP.Get_kind(_env) == Nodes_NodeKind_get_LiteralBool(_env){
                        {
                            _literal := TransUnit_convExp0_8416(Lns_2DDD(exp2.FP.GetLiteral(_env)))
                            if !Lns_IsNil( _literal ) {
                                literal := _literal
                                if Lns_op_not(Nodes_getLiteralObj(_env, literal)){
                                    self.FP.AddErrMess(_env, exp2.FP.Get_pos(_env), "this value never be 'true'")
                                }
                            }
                        }
                    }
                    if exp1Type.FP.Get_nilable(_env){
                        if exp2Type.FP.Get_nilable(_env){
                            retType = exp2Type
                        } else { 
                            retType = exp2Type.FP.Get_nilableTypeInfo(_env)
                        }
                    } else if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeBool, nil, nil)) ||
                        _env.SetStackVal( exp2Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeBool, nil, nil)) ).(bool){
                        if Lns_car(exp1Type.FP.CanEvalWith(_env, self.ProcessInfo, exp2Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
                            retType = exp1Type
                        } else if Lns_car(exp2Type.FP.CanEvalWith(_env, self.ProcessInfo, exp1Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool){
                            retType = exp2Type
                        } else { 
                            if exp2Type.FP.Get_nilable(_env){
                                retType = Ast_builtinTypeStem_
                            } else { 
                                retType = Ast_builtinTypeStem
                            }
                        }
                    } else if exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeStem, nil, nil){
                        retType = Ast_builtinTypeStem
                    } else { 
                        retType = exp2Type
                    }
                    retType = &NewAst_AndExpTypeInfo(_env, self.ProcessInfo, exp1Type, exp2Type, retType).Ast_TypeInfo
                } else if _switch0 == "<" || _switch0 == ">" || _switch0 == "<=" || _switch0 == ">=" {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_car(Ast_builtinTypeString.FP.CanEvalWith(_env, self.ProcessInfo, exp1Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) &&
                        _env.SetStackVal( Lns_car(Ast_builtinTypeString.FP.CanEvalWith(_env, self.ProcessInfo, exp2Type, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ||
                        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(_env, self.ProcessInfo, exp1Type, Ast_CanEvalType__Comp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ||
                            _env.SetStackVal( Lns_car(Ast_builtinTypeReal.FP.CanEvalWith(_env, self.ProcessInfo, exp1Type, Ast_CanEvalType__Comp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ).(bool))) &&
                        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(_env, self.ProcessInfo, exp2Type, Ast_CanEvalType__Comp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ||
                            _env.SetStackVal( Lns_car(Ast_builtinTypeReal.FP.CanEvalWith(_env, self.ProcessInfo, exp2Type, Ast_CanEvalType__Comp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool)) ).(bool))) ).(bool){
                    } else { 
                        self.FP.AddErrMess(_env, opToken.Pos, _env.GetVM().String_format("no numeric type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil), exp2Type.FP.GetTxt(_env, self.TypeNameCtrl, nil, nil)}))
                    }
                    retType = Ast_builtinTypeBool
                } else if _switch0 == "~=" || _switch0 == "==" {
                    exp, exp2 = self.FP.analyzeExpOpEquals(_env, firstToken.Pos, opToken, exp, exp2)
                    retType = Ast_builtinTypeBool
                } else if _switch0 == "^" || _switch0 == "|" || _switch0 == "~" || _switch0 == "&" || _switch0 == "|<<" || _switch0 == "|>>" {
                    if self.TargetLuaVer.FP.Get_hasBitOp(_env) == LuaVer_BitOp__Cant{
                        self.FP.AddErrMess(_env, opToken.Pos, "this lua version can't use bit operand.")
                    }
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(_env, self.ProcessInfo, exp1Type, Ast_CanEvalType__Logical, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ||
                        _env.SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(_env, self.ProcessInfo, exp2Type, Ast_CanEvalType__Logical, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool){
                        self.FP.AddErrMess(_env, opToken.Pos, _env.GetVM().String_format("no int type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(_env, nil, nil, nil), exp2Type.FP.GetTxt(_env, nil, nil, nil)}))
                    }
                    retType = Ast_builtinTypeInt
                } else if _switch0 == ".." {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeString, nil, nil))) ||
                        _env.SetStackVal( Lns_op_not(exp2Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeString, nil, nil))) ).(bool){
                        self.FP.AddErrMess(_env, opToken.Pos, _env.GetVM().String_format("no string type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(_env, nil, nil, nil), exp2Type.FP.GetTxt(_env, nil, nil, nil)}))
                    }
                    retType = Ast_builtinTypeString
                } else if _switch0 == "+" || _switch0 == "-" || _switch0 == "*" || _switch0 == "/" || _switch0 == "%" {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(_env, self.ProcessInfo, exp1Type, Ast_CanEvalType__Math, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) &&
                            _env.SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeReal.FP.CanEvalWith(_env, self.ProcessInfo, exp1Type, Ast_CanEvalType__Math, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool))) ||
                        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeInt.FP.CanEvalWith(_env, self.ProcessInfo, exp2Type, Ast_CanEvalType__Math, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) &&
                            _env.SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeReal.FP.CanEvalWith(_env, self.ProcessInfo, exp2Type, Ast_CanEvalType__Math, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool))) ).(bool){
                        self.FP.AddErrMess(_env, opToken.Pos, _env.GetVM().String_format("no numeric type '%s' or '%s'", []LnsAny{exp1Type.FP.GetTxt(_env, nil, nil, nil), exp2Type.FP.GetTxt(_env, nil, nil, nil)}))
                    }
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeReal, nil, nil)) ||
                        _env.SetStackVal( exp2Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeReal, nil, nil)) ).(bool){
                        retType = Ast_builtinTypeReal
                        if exp1Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeInt, nil, nil){
                            exp = &Nodes_ExpCastNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeReal)}), exp, Ast_builtinTypeReal, nil, "", Nodes_CastKind__Implicit).Nodes_Node
                        } else if exp2Type.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeInt, nil, nil){
                            exp2 = &Nodes_ExpCastNode_create(_env, self.NodeManager, exp2.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeReal)}), exp2, Ast_builtinTypeReal, nil, "", Nodes_CastKind__Implicit).Nodes_Node
                        }
                    } else { 
                        retType = Ast_builtinTypeInt
                    }
                } else {
                    self.FP.Error(_env, "unknown op " + opTxt)
                }
                exp = &Nodes_ExpOp2Node_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(retType)}), opToken, self.NodeManager.FP.MultiTo1(_env, exp), self.NodeManager.FP.MultiTo1(_env, exp2)).Nodes_Node
            } else { 
                self.FP.Error(_env, "illegal op")
            }
        } else { 
            self.FP.Pushback(_env)
            return exp
        }
    }
// insert a dummy
    return nil
}
// 3631: decl @lune.@base.@TransUnit.TransUnit.analyzeExpMacroStat
func (self *TransUnit_TransUnit) analyzeExpMacroStat(_env *LnsEnv, firstToken *Types_Token) *Nodes_ExpMacroStatNode {
    var expStrList *LnsList
    expStrList = NewLnsList([]LnsAny{})
    self.FP.checkNextToken(_env, "{")
    var braceCount LnsInt
    braceCount = 0
    var prevToken *Types_Token
    prevToken = firstToken
    var errMessList *LnsList
    errMessList = NewLnsList([]LnsAny{})
    for  {
        var token *Types_Token
        token = self.FP.getToken(_env, nil)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Txt == ",,") ||
            _env.SetStackVal( token.Txt == ",,,") ||
            _env.SetStackVal( token.Txt == ",,,,") ).(bool){
            var exp *Nodes_Node
            exp = self.FP.analyzeExp(_env, false, true, false, Lns_unwrap( TransUnit_op1levelMap.Get(token.Txt)).(LnsInt), nil)
            var literalStr *Nodes_LiteralStringNode
            literalStr = self.MacroCtrl.FP.ExpandSymbol(_env, self.FP, self.inTestBlock, token, exp, self.NodeManager, errMessList)
            for _, _errMess := range( errMessList.Items ) {
                errMess := _errMess.(Macro_ErrorMessDownCast).ToMacro_ErrorMess()
                self.FP.AddErrMess(_env, errMess.Pos, errMess.Mess)
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
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( prevToken == firstToken) ||
                _env.SetStackVal( token.Consecutive) ).(bool){
                format = "'%s'"
                consecutive = true
            } else { 
                consecutive = false
            }
            var newToken *Types_Token
            newToken = NewTypes_Token(_env, token.Kind, _env.GetVM().String_format(format, []LnsAny{token.Txt}), token.Pos, consecutive, nil)
            var literalStr *Nodes_LiteralStringNode
            literalStr = Nodes_LiteralStringNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), newToken, nil, nil)
            expStrList.Insert(Nodes_LiteralStringNode2Stem(literalStr))
        }
        prevToken = token
    }
    return Nodes_ExpMacroStatNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeStat)}), expStrList)
}
// 3694: decl @lune.@base.@TransUnit.TransUnit.analyzeSuper
func (self *TransUnit_TransUnit) analyzeSuper(_env *LnsEnv, firstToken *Types_Token) *Nodes_Node {
    self.FP.checkNextToken(_env, "(")
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var expList LnsAny
    expList = nil
    if nextToken.Txt != ")"{
        self.FP.Pushback(_env)
        expList = self.FP.analyzeExpList(_env, false, false, false, nil, nil, nil)
        self.FP.checkNextToken(_env, ")")
    }
    self.FP.checkNextToken(_env, ";")
    var classType *Ast_TypeInfo
    classType = self.FP.getCurrentClass(_env)
    var currentFunc *Ast_TypeInfo
    currentFunc = self.FP.GetCurrentNamespaceTypeInfo(_env)
    if currentFunc.FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
        var superType *Ast_TypeInfo
        superType = classType.FP.Get_baseTypeInfo(_env)
        if superType.FP.Equals(_env, self.ProcessInfo, Ast_headTypeInfo, nil, nil){
            self.FP.AddErrMess(_env, firstToken.Pos, "This class doesn't have super-class.")
        } else { 
            if currentFunc.FP.Get_rawTxt(_env) == "__init"{
                var superScope *Ast_Scope
                
                {
                    _superScope := superType.FP.Get_scope(_env)
                    if _superScope == nil{
                        self.FP.Error(_env, "not found super scope")
                    } else {
                        superScope = _superScope.(*Ast_Scope)
                    }
                }
                var superCtorType *Ast_TypeInfo
                
                {
                    _superCtorType := superScope.FP.GetTypeInfoChild(_env, "__init")
                    if _superCtorType == nil{
                        self.FP.Error(_env, "not found super '__init'")
                    } else {
                        superCtorType = _superCtorType.(*Ast_TypeInfo)
                    }
                }
                self.FP.checkMatchValType(_env, firstToken.Pos, superCtorType, expList, NewLnsList([]LnsAny{}), classType)
                return &Nodes_ExpCallSuperCtorNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), superType, superCtorType, expList).Nodes_Node
            } else { 
                {
                    _superFunc := (Lns_unwrap( superType.FP.Get_scope(_env)).(*Ast_Scope)).FP.GetTypeInfoField(_env, currentFunc.FP.Get_rawTxt(_env), true, self.FP.Get_scope(_env), self.ScopeAccess)
                    if !Lns_IsNil( _superFunc ) {
                        superFunc := _superFunc.(*Ast_TypeInfo)
                        if superFunc.FP.Get_abstractFlag(_env){
                            self.FP.AddErrMess(_env, firstToken.Pos, "super is abstract.")
                        }
                        self.FP.checkMatchValType(_env, firstToken.Pos, superFunc, expList, NewLnsList([]LnsAny{}), classType)
                        var exp *Nodes_ExpCallSuperNode
                        exp = Nodes_ExpCallSuperNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), superType, superFunc, expList)
                        return &Nodes_StmtExpNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), exp.FP.Get_expTypeList(_env), &exp.Nodes_Node).Nodes_Node
                    }
                }
                self.FP.AddErrMess(_env, firstToken.Pos, "this is not override method.")
                return self.FP.createNoneNode(_env, firstToken.Pos)
            }
        }
    }
    self.FP.AddErrMess(_env, firstToken.Pos, "super can't call here.")
    return self.FP.createNoneNode(_env, firstToken.Pos)
}
// 3759: decl @lune.@base.@TransUnit.TransUnit.analyzeUnwrap
func (self *TransUnit_TransUnit) analyzeUnwrap(_env *LnsEnv, firstToken *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    var continueFlag bool
    nextToken,continueFlag = self.FP.getContinueToken(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(continueFlag)) ||
        _env.SetStackVal( nextToken.Txt != "!") ).(bool){
        self.FP.Pushback(_env)
        self.FP.PushbackToken(_env, firstToken)
        var exp *Nodes_Node
        exp = self.FP.analyzeExp(_env, false, false, false, nil, nil)
        self.FP.checkNextToken(_env, ";")
        if Lns_op_not(exp.FP.Get_expType(_env).FP.Get_nilable(_env)){
            self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), "this value is not nilable.")
        }
        return &Nodes_StmtExpNode_create(_env, self.NodeManager, nextToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), exp).Nodes_Node
    }
    self.FP.Pushback(_env)
    return self.FP.analyzeDeclVar(_env, Nodes_DeclVarMode__Unwrap, Ast_AccessMode__Local, firstToken)
}
// 3779: decl @lune.@base.@TransUnit.TransUnit.analyzeExpUnwrap
func (self *TransUnit_TransUnit) analyzeExpUnwrap(_env *LnsEnv, firstToken *Types_Token) *Nodes_Node {
    var expNode *Nodes_Node
    expNode = self.FP.analyzeExpOneRVal(_env, false, true, nil, nil)
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var insNode LnsAny
    insNode = nil
    if nextToken.Txt == "default"{
        insNode = self.FP.analyzeExpOneRVal(_env, false, false, nil, nil)
    } else { 
        self.FP.Pushback(_env)
    }
    var unwrapType *Ast_TypeInfo
    unwrapType = Ast_builtinTypeStem_
    var expType *Ast_TypeInfo
    expType = expNode.FP.Get_expType(_env)
    if Lns_op_not(expType.FP.Get_nilable(_env)){
        unwrapType = expType
        self.FP.AddErrMess(_env, expNode.FP.Get_pos(_env), _env.GetVM().String_format("this exp is not nilable -- %s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil)}))
    } else if expType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        if expType.FP.Get_itemTypeInfoList(_env).Len() > 0{
            unwrapType = expType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_nonnilableType(_env)
        } else { 
            unwrapType = Ast_builtinTypeStem
        }
    } else { 
        unwrapType = expType.FP.Get_nonnilableType(_env)
    }
    if insNode != nil{
        insNode_1841 := insNode.(*Nodes_Node)
        var insType *Ast_TypeInfo
        insType = insNode_1841.FP.Get_expType(_env)
        if insType.FP.Get_nilable(_env){
            self.FP.AddErrMess(_env, insNode_1841.FP.Get_pos(_env), _env.GetVM().String_format("default can't use nilable -- %s", []LnsAny{insType.FP.GetTxt(_env, nil, nil, nil)}))
        }
        var alt2type *LnsMap
        alt2type = Ast_CanEvalCtrlTypeInfo_createDefaultAlt2typeMap(_env, false)
        if Lns_op_not(Lns_car(unwrapType.FP.CanEvalWith(_env, self.ProcessInfo, insType, Ast_CanEvalType__SetOp, alt2type)).(bool)){
            if Lns_op_not(Lns_car(insType.FP.CanEvalWith(_env, self.ProcessInfo, unwrapType, Ast_CanEvalType__SetOp, alt2type)).(bool)){
                unwrapType = Ast_builtinTypeStem
            } else { 
                unwrapType = insType
            }
        }
    }
    self.HelperInfo.UseUnwrapExp = true
    if Ast_isExtType(_env, expType.FP.Get_nonnilableType(_env)){
        switch _matchExp0 := self.ProcessInfo.FP.CreateLuaval(_env, unwrapType, false).(type) {
        case *Ast_LuavalResult__OK:
        work := _matchExp0.Val1
            unwrapType = work
        case *Ast_LuavalResult__Err:
        err := _matchExp0.Val1
            self.FP.AddErrMess(_env, firstToken.Pos, err)
        }
    }
    return &Nodes_ExpUnwrapNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(unwrapType)}), expNode, insNode).Nodes_Node
}
// 3862: decl @lune.@base.@TransUnit.TransUnit.analyzeStrConst
func (self *TransUnit_TransUnit) analyzeStrConst(_env *LnsEnv, firstToken *Types_Token,token *Types_Token) *Nodes_Node {
    var exp *Nodes_Node
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, true)
    if nextToken.Kind != Types_TokenKind__Eof{
        var param LnsAny
        var dddParam LnsAny
        if nextToken.Txt == "("{
            var argNodeList *Nodes_ExpListNode
            argNodeList = self.FP.analyzeExpList(_env, false, false, false, nil, nil, nil)
            param = argNodeList
            var workExpList LnsAny
            _,_,workExpList = TransUnit_convExp0_10695(Lns_2DDD(self.FP.checkMatchType(_env, "str constructor", firstToken.Pos, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeDDD)}), argNodeList, false, false, nil, true)))
            if workExpList != nil{
                workExpList_1866 := workExpList.(*Nodes_ExpListNode)
                dddParam = workExpList_1866
            } else {
                dddParam = nil
            }
            self.FP.checkNextToken(_env, ")")
            nextToken = self.FP.getToken(_env, true)
            if param != nil{
                param_1869 := param.(*Nodes_ExpListNode)
                self.FP.checkStringFormat(_env, token.Pos, token.Txt, param_1869.FP.Get_expTypeList(_env))
            }
        } else { 
            param = nil
            dddParam = nil
        }
        var workExp *Nodes_Node
        workExp = &Nodes_LiteralStringNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), token, param, dddParam).Nodes_Node
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( nextToken.Txt == "[") ||
            _env.SetStackVal( nextToken.Txt == "$[") ).(bool){
            exp = self.FP.analyzeExpRefItem(_env, nextToken, workExp, nextToken.Txt == "$[")
        } else { 
            exp = workExp
            if nextToken.Kind != Types_TokenKind__Eof{
                self.FP.Pushback(_env)
            }
        }
    } else { 
        exp = &Nodes_LiteralStringNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeString)}), token, nil, nil).Nodes_Node
    }
    return exp
}
// 3921: decl @lune.@base.@TransUnit.TransUnit.analyzeExp
func (self *TransUnit_TransUnit) analyzeExp(_env *LnsEnv, allowNoneType bool,skipOp2Flag bool,canLeftExp bool,prevOpLevel LnsAny,expectType LnsAny) *Nodes_Node {
    var firstToken *Types_Token
    firstToken = self.FP.getToken(_env, nil)
    var processsExpectExp func(_env *LnsEnv, token *Types_Token,orgExpectType *Ast_TypeInfo) *Nodes_Node
    processsExpectExp = func(_env *LnsEnv, token *Types_Token,orgExpectType *Ast_TypeInfo) *Nodes_Node {
        {
            _enumTypeInfo := Ast_EnumTypeInfoDownCastF(orgExpectType.FP.Get_srcTypeInfo(_env).FP.Get_aliasSrc(_env).FP)
            if !Lns_IsNil( _enumTypeInfo ) {
                enumTypeInfo := _enumTypeInfo.(*Ast_EnumTypeInfo)
                var nextToken *Types_Token
                nextToken = self.FP.getToken(_env, nil)
                self.FP.checkEnumComp(_env, nextToken, enumTypeInfo)
                {
                    _valInfo := enumTypeInfo.FP.GetEnumValInfo(_env, nextToken.Txt)
                    if !Lns_IsNil( _valInfo ) {
                        valInfo := _valInfo.(*Ast_EnumValInfo)
                        var aliasType LnsAny
                        aliasType = nil
                        var expType *Ast_TypeInfo
                        expType = &enumTypeInfo.Ast_TypeInfo
                        aliasType = self.ImportedAliasMap.Get(&enumTypeInfo.Ast_TypeInfo)
                        if aliasType != nil{
                            aliasType_1891 := aliasType.(*Ast_AliasTypeInfo)
                            expType = &aliasType_1891.Ast_TypeInfo
                        }
                        if Lns_op_not(self.moduleType.FP.Equals(_env, self.ProcessInfo, orgExpectType.FP.GetModule(_env), nil, nil)){
                            if Lns_op_not(self.ImportModuleSet.Has(Ast_TypeInfo2Stem(orgExpectType.FP.GetModule(_env)))){
                                if Lns_op_not(aliasType){
                                    var fullname string
                                    fullname = orgExpectType.FP.GetFullName(_env, self.TypeNameCtrl, self.FP.Get_scope(_env).FP, true)
                                    self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("need to import module -- %s (%s)", []LnsAny{fullname, enumTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
                                }
                            }
                        }
                        var exp *Nodes_Node
                        exp = &Nodes_ExpOmitEnumNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(expType)}), nextToken, valInfo, aliasType, enumTypeInfo).Nodes_Node
                        return self.FP.analyzeExpCont(_env, firstToken, exp, false, canLeftExp)
                    }
                }
                self.FP.Error(_env, _env.GetVM().String_format("illegal enum val -- %s.%s", []LnsAny{orgExpectType.FP.GetTxt(_env, nil, nil, nil), nextToken.Txt}))
            }
        }
        {
            _algeTyepInfo := Ast_AlgeTypeInfoDownCastF(orgExpectType.FP.Get_srcTypeInfo(_env).FP)
            if !Lns_IsNil( _algeTyepInfo ) {
                algeTyepInfo := _algeTyepInfo.(*Ast_AlgeTypeInfo)
                return &self.FP.analyzeNewAlge(_env, firstToken, algeTyepInfo, nil).Nodes_Node
            }
        }
        self.FP.Error(_env, _env.GetVM().String_format("illegal type for '.' -- %s", []LnsAny{orgExpectType.FP.GetTxt(_env, nil, nil, nil)}))
    // insert a dummy
        return nil
    }
    var processsNewExp func(_env *LnsEnv, token *Types_Token) *Nodes_Node
    processsNewExp = func(_env *LnsEnv, token *Types_Token) *Nodes_Node {
        var exp *Nodes_Node
        exp = &self.FP.analyzeRefType(_env, Ast_AccessMode__Local, false, false).Nodes_Node
        var classTypeInfo *Ast_TypeInfo
        classTypeInfo = exp.FP.Get_expType(_env)
        if _switch0 := classTypeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Class || _switch0 == Ast_TypeInfoKind__IF {
            if classTypeInfo.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeString, nil, nil){
                self.FP.Error(_env, _env.GetVM().String_format("'new' can't use this type -- %s", []LnsAny{classTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
            }
        } else {
            self.FP.Error(_env, _env.GetVM().String_format("'new' can't use this type -- %s", []LnsAny{classTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
        }
        if classTypeInfo.FP.Get_externalFlag(_env){
            if _switch1 := classTypeInfo.FP.Get_accessMode(_env); _switch1 == Ast_AccessMode__Pri || _switch1 == Ast_AccessMode__Local {
                self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("Can't access -- %s", []LnsAny{Ast_AccessMode_getTxt( classTypeInfo.FP.Get_accessMode(_env))}))
            }
        }
        if classTypeInfo.FP.Get_abstractFlag(_env){
            self.FP.AddErrMess(_env, token.Pos, "abstract class can't new")
        }
        var classScope LnsAny
        classScope = classTypeInfo.FP.Get_scope(_env)
        var initTypeInfo *Ast_TypeInfo
        
        {
            _initTypeInfo := (Lns_unwrap( classScope).(*Ast_Scope)).FP.GetTypeInfoChild(_env, "__init")
            if _initTypeInfo == nil{
                self.FP.Error(_env, "not found __init")
            } else {
                initTypeInfo = _initTypeInfo.(*Ast_TypeInfo)
            }
        }
        self.FP.checkNoasyncType(_env, token.Pos, initTypeInfo)
        self.FP.checkNextToken(_env, "(")
        var nextToken *Types_Token
        nextToken = self.FP.getToken(_env, nil)
        var argList LnsAny
        argList = nil
        if nextToken.Txt != ")"{
            self.FP.Pushback(_env)
            argList = self.FP.analyzeExpList(_env, false, false, false, nil, initTypeInfo.FP.Get_argTypeInfoList(_env), nil)
            self.FP.checkNextToken(_env, ")")
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( initTypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) ||
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( initTypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Pro) &&
                _env.SetStackVal( self.FP.Get_scope(_env).FP.GetClassTypeInfo(_env).FP.IsInheritFrom(_env, self.ProcessInfo, classTypeInfo, nil)) ).(bool))) ||
            _env.SetStackVal( (self.FP.Get_scope(_env).FP.GetClassTypeInfo(_env) == classTypeInfo)) ||
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( initTypeInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Local) &&
                _env.SetStackVal( initTypeInfo.FP.GetModule(_env) == self.moduleType) ).(bool))) ).(bool){
        } else { 
            self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("can't access to __init of %s", []LnsAny{classTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
        }
        var alt2type *LnsMap
        var newArgList LnsAny
        _,alt2type,newArgList = self.FP.checkMatchValType(_env, exp.FP.Get_pos(_env), initTypeInfo, argList, classTypeInfo.FP.Get_itemTypeInfoList(_env), classTypeInfo)
        if classTypeInfo.FP.Get_itemTypeInfoList(_env).Len() > 0{
            if classTypeInfo.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                var genTypeList *LnsList
                genTypeList = NewLnsList([]LnsAny{})
                var detect bool
                detect = true
                for _, _altType := range( classTypeInfo.FP.Get_itemTypeInfoList(_env).Items ) {
                    altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    {
                        __exp := alt2type.Get(altType)
                        if !Lns_IsNil( __exp ) {
                            _exp := __exp.(*Ast_TypeInfo)
                            genTypeList.Insert(Ast_TypeInfo2Stem(_exp))
                        } else {
                            self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("Can't new generic class. -- %s", []LnsAny{classTypeInfo.FP.GetTxt(_env, nil, nil, nil)}))
                            detect = false
                            break
                        }
                    }
                }
                if detect{
                    classTypeInfo = TransUnit_convExp0_11713(Lns_2DDD(self.ProcessInfo.FP.CreateGeneric(_env, classTypeInfo, genTypeList, self.moduleType)))
                }
            }
        }
        exp = &Nodes_ExpNewNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classTypeInfo)}), exp, initTypeInfo, newArgList).Nodes_Node
        exp = self.FP.analyzeExpCont(_env, firstToken, exp, false, canLeftExp)
        return exp
    }
    var processOp1 func(_env *LnsEnv, token *Types_Token)(*Nodes_Node, bool)
    processOp1 = func(_env *LnsEnv, token *Types_Token)(*Nodes_Node, bool) {
        if token.Txt == "`"{
            return &self.FP.analyzeExpMacroStat(_env, token).Nodes_Node, false
        }
        var exp *Nodes_Node
        exp = self.FP.analyzeExpOneRVal(_env, false, true, Lns_unwrap( TransUnit_op1levelMap.Get(token.Txt)).(LnsInt), nil)
        var typeInfo *Ast_TypeInfo
        typeInfo = Ast_builtinTypeNone
        var macroExpFlag bool
        macroExpFlag = false
        var expType *Ast_TypeInfo
        expType = exp.FP.Get_expType(_env)
        if expType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            self.FP.AddErrMess(_env, exp.FP.Get_pos(_env), _env.GetVM().String_format("... can't evaluate for '%s'.", []LnsAny{token.Txt}))
        }
        if _switch0 := (token.Txt); _switch0 == "-" {
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeInt, nil, nil))) &&
                _env.SetStackVal( Lns_op_not(expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeReal, nil, nil))) ).(bool)){
                self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("unmatch type for \"-\" -- %s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil)}))
            }
            typeInfo = expType
        } else if _switch0 == "#" {
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( expType.FP.Get_extedType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__List) &&
                _env.SetStackVal( expType.FP.Get_extedType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Array) &&
                _env.SetStackVal( Lns_op_not(Lns_car(Ast_builtinTypeString.FP.CanEvalWith(_env, self.ProcessInfo, expType, Ast_CanEvalType__SetOp, NewLnsMap( map[LnsAny]LnsAny{}))).(bool))) ).(bool)){
                self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("unmatch type for \"#\" -- %s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil)}))
            }
            if expType.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( Lns_op_not(self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.IsAnalyzingBlockArg(_env))) &&
                    _env.SetStackVal( Lns_op_not(self.FP.Get_curNsInfo(_env).FP.CanAccessLuaval(_env))) ).(bool)){
                    self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("can't access the Luaval with '#' without __luago. -- %s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
            typeInfo = Ast_builtinTypeInt
        } else if _switch0 == "not" {
            typeInfo = Ast_builtinTypeBool
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(expType.FP.Get_nilable(_env))) &&
                _env.SetStackVal( Lns_op_not(expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeBool, nil, nil))) &&
                _env.SetStackVal( Lns_op_not(expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeStem, nil, nil))) &&
                _env.SetStackVal( expType.FP.Get_kind(_env) != Ast_TypeInfoKind__DDD) ).(bool)){
                self.FP.AddErrMess(_env, token.Pos, "this 'not' operand never be false")
            }
        } else if _switch0 == ",," {
            macroExpFlag = true
            typeInfo = expType
        } else if _switch0 == ",,," {
            macroExpFlag = true
            if Lns_op_not(expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeString, nil, nil)){
                self.FP.Error(_env, "unmatch ,,, type, need string type")
            }
            typeInfo = Ast_builtinTypeSymbol
        } else if _switch0 == ",,,," {
            macroExpFlag = true
            if Lns_op_not(expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeSymbol, nil, nil)){
                self.FP.Error(_env, "unmatch ,,, type, need symbol type")
            }
            typeInfo = Ast_builtinTypeString
        } else if _switch0 == "`" {
            typeInfo = Ast_builtinTypeNone
        } else if _switch0 == "~" {
            if Lns_op_not(expType.FP.Equals(_env, self.ProcessInfo, Ast_builtinTypeInt, nil, nil)){
                self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("unmatch type for \"~\" -- %s", []LnsAny{expType.FP.GetTxt(_env, nil, nil, nil)}))
            }
            typeInfo = Ast_builtinTypeInt
        } else {
            self.FP.Error(_env, "unknown op1")
        }
        if macroExpFlag{
            var nextToken *Types_Token
            nextToken = self.FP.getToken(_env, true)
            if nextToken.Txt != "~~"{
                self.FP.Pushback(_env)
            }
        }
        exp = &Nodes_ExpOp1Node_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(typeInfo)}), token, self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.Get_mode(_env), self.NodeManager.FP.MultiTo1(_env, exp)).Nodes_Node
        return self.FP.analyzeExpOp2(_env, firstToken, exp, prevOpLevel), true
    }
    var token *Types_Token
    token = firstToken
    var exp *Nodes_Node
    exp = self.FP.createNoneNode(_env, firstToken.Pos)
    if token.Txt == "##"{
        if allowNoneType{
            self.FP.AddErrMess(_env, token.Pos, "illeal syntax -- ##")
        }
        return &Nodes_AbbrNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeAbbr)})).Nodes_Node
    }
    if token.Kind == Types_TokenKind__Dlmt{
        if token.Txt == "."{
            if expectType != nil{
                expectType_1970 := expectType.(*Ast_TypeInfo)
                var orgExpectType *Ast_TypeInfo
                orgExpectType = expectType_1970
                if orgExpectType.FP.Get_nilable(_env){
                    orgExpectType = orgExpectType.FP.Get_nonnilableType(_env)
                }
                exp = processsExpectExp(_env, token, orgExpectType)
            } else {
                self.FP.Error(_env, "illegal '.'")
            }
        } else if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Txt == "[") ||
            _env.SetStackVal( token.Txt == "[@") ).(bool){
            exp = self.FP.analyzeListConst(_env, token, expectType)
        } else if token.Txt == "(@"{
            exp = self.FP.analyzeSetConst(_env, token, expectType)
        } else if token.Txt == "{"{
            exp = &self.FP.analyzeMapConst(_env, token, expectType).Nodes_Node
        } else if token.Txt == "("{
            exp = self.FP.analyzeExp(_env, false, false, false, nil, nil)
            self.FP.checkNextToken(_env, ")")
            if Lns_op_not(exp.FP.CanBeRight(_env, self.ProcessInfo)){
                self.FP.AddErrMess(_env, exp.FP.Get_effectivePos(_env), _env.GetVM().String_format("can't be r-value in paren. -- %s", []LnsAny{Nodes_getNodeKindName(_env, exp.FP.Get_kind(_env))}))
            }
            exp = &Nodes_ExpParenNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp.FP.Get_expType(_env))}), exp).Nodes_Node
            exp = self.FP.analyzeExpCont(_env, firstToken, exp, false, canLeftExp)
        }
    }
    if token.Txt == "new"{
        exp = processsNewExp(_env, token)
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Kind == Types_TokenKind__Ope) &&
        _env.SetStackVal( Parser_isOp1(_env, token.Txt)) ).(bool)){
        var workExp *Nodes_Node
        var fin bool
        workExp,fin = processOp1(_env, token)
        if fin{
            return workExp
        }
        exp = workExp
    }
    if token.Kind == Types_TokenKind__Int{
        exp = &Nodes_LiteralIntNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), token, Lns_forceCastInt((Lns_unwrapDefault( Lns_tonumber(token.Txt, nil), 0)))).Nodes_Node
    } else if token.Kind == Types_TokenKind__Real{
        exp = &Nodes_LiteralRealNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeReal)}), token, (Lns_unwrapDefault( Lns_tonumber(token.Txt, nil), 0.0).(LnsReal))).Nodes_Node
    } else if token.Kind == Types_TokenKind__Char{
        var num LnsInt
        if len(token.Txt) == 1{
            num = LnsInt(token.Txt[1-1])
        } else { 
            num = Lns_unwrap( TransUnit_quotedChar2Code.Get(_env.GetVM().String_sub(token.Txt,2, 2))).(LnsInt)
        }
        exp = &Nodes_LiteralCharNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeChar)}), token, num).Nodes_Node
    } else if token.Kind == Types_TokenKind__Str{
        exp = self.FP.analyzeStrConst(_env, firstToken, token)
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Kind == Types_TokenKind__Symb) &&
        _env.SetStackVal( token.Txt == "__line__") ).(bool)){
        var pos Types_Position
        pos = self.FP.getLineNo(_env, token)
        exp = &Nodes_LiteralIntNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeInt)}), NewTypes_Token(_env, Types_TokenKind__Int, _env.GetVM().String_format("%d", []LnsAny{pos.LineNo}), token.Pos, false, nil), token.Pos.LineNo).Nodes_Node
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Kind == Types_TokenKind__Kywd) &&
        _env.SetStackVal( token.Txt == "fn") ).(bool)){
        exp = self.FP.analyzeExpSymbol(_env, firstToken, token, TransUnit_ExpSymbolMode__Fn, nil, false, false)
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Kind == Types_TokenKind__Kywd) &&
        _env.SetStackVal( token.Txt == "unwrap") ).(bool)){
        exp = self.FP.analyzeExpUnwrap(_env, token)
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Kind == Types_TokenKind__Kywd) &&
        _env.SetStackVal( token.Txt == "__request") ).(bool)){
        exp = &self.FP.analyzeRequest(_env, token).Nodes_Node
    } else if token.Kind == Types_TokenKind__Symb{
        exp = self.FP.analyzeExpSymbol(_env, firstToken, token, TransUnit_ExpSymbolMode__Symbol, nil, false, canLeftExp)
        var symbolInfoList *LnsList
        symbolInfoList = exp.FP.GetSymbolInfo(_env)
        if symbolInfoList.Len() == 1{
            var symbolInfo *Ast_SymbolInfo
            symbolInfo = symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Typ{
                exp = &self.FP.analyzeRefTypeWithSymbol(_env, Ast_AccessMode__Local, false, nil, exp, false).Nodes_Node
                var workToken *Types_Token
                workToken = self.FP.getToken(_env, nil)
                if workToken.Txt == "."{
                    exp = self.FP.analyzeExpSymbol(_env, firstToken, self.FP.getToken(_env, nil), TransUnit_ExpSymbolMode__Field, exp, false, canLeftExp)
                } else { 
                    self.FP.Pushback(_env)
                }
            }
        }
    } else if token.Kind == Types_TokenKind__Type{
        var symbolTypeInfo *Ast_SymbolInfo
        
        {
            _symbolTypeInfo := Ast_getSym2builtInTypeMap(_env).Get(token.Txt)
            if _symbolTypeInfo == nil{
                self.FP.Error(_env, _env.GetVM().String_format("unknown type -- %s", []LnsAny{token.Txt}))
            } else {
                symbolTypeInfo = _symbolTypeInfo.(*Ast_SymbolInfo)
            }
        }
        exp = &Nodes_ExpRefNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), &NewAst_AccessSymbolInfo(_env, self.ProcessInfo, symbolTypeInfo, Ast_OverrideMut__None_Obj, false).Ast_SymbolInfo).Nodes_Node
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Kind == Types_TokenKind__Kywd) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Txt == "true") ||
            _env.SetStackVal( token.Txt == "false") ).(bool))) ).(bool)){
        exp = &Nodes_LiteralBoolNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeBool)}), token).Nodes_Node
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( token.Kind == Types_TokenKind__Kywd) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Txt == "nil") ||
            _env.SetStackVal( token.Txt == "null") ).(bool))) ).(bool)){
        exp = &Nodes_LiteralNilNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNil)})).Nodes_Node
    }
    if exp.FP.Get_kind(_env) == Nodes_NodeKind_get_None(_env){
        self.FP.Error(_env, "illegal exp")
    }
    if skipOp2Flag{
        return exp
    }
    return self.FP.analyzeExpOp2(_env, firstToken, exp, prevOpLevel)
}
// 27: decl @lune.@base.@TransUnit.TransUnit.analyzeReturn
func (self *TransUnit_TransUnit) analyzeReturn(_env *LnsEnv, token *Types_Token) *Nodes_ReturnNode {
    var expList LnsAny
    expList = nil
    var funcTypeInfo *Ast_TypeInfo
    funcTypeInfo = self.FP.GetCurrentNamespaceTypeInfo(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(Ast_TypeInfo_hasParent(_env, funcTypeInfo))) ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( funcTypeInfo.FP.Get_kind(_env) != Ast_TypeInfoKind__Func) &&
            _env.SetStackVal( funcTypeInfo.FP.Get_kind(_env) != Ast_TypeInfoKind__Method) ).(bool))) ).(bool){
        self.FP.AddErrMess(_env, token.Pos, "'return' could not use here")
    }
    var nsInfo *TransUnitIF_NSInfo
    nsInfo = self.FP.GetNSInfo(_env, funcTypeInfo)
    if nsInfo.FP.IsLockedAsync(_env){
        self.FP.AddErrMess(_env, token.Pos, "can't use 'return' in the __asyncLock.")
    }
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    var retTypeList *LnsList
    retTypeList = funcTypeInfo.FP.Get_retTypeInfoList(_env)
    if nextToken.Txt != ";"{
        self.FP.Pushback(_env)
        expList = self.FP.analyzeExpList(_env, false, false, false, nil, retTypeList, nil)
        self.FP.checkNextToken(_env, ";")
    }
    if funcTypeInfo.FP.GetTxt(_env, nil, nil, nil) == "__init"{
        self.FP.AddErrMess(_env, token.Pos, "__init method can't return")
    }
    {
        _workList := expList
        if !Lns_IsNil( _workList ) {
            workList := _workList.(*Nodes_ExpListNode)
            {
                _, _, _newExpNodeList := TransUnit_convExp0_13806(Lns_2DDD(self.FP.checkMatchType(_env, "return", token.Pos, retTypeList, workList, false, Lns_op_not(workList.FP.Get_followOn(_env)), nil, true)))
                if !Lns_IsNil( _newExpNodeList ) {
                    newExpNodeList := _newExpNodeList.(*Nodes_ExpListNode)
                    expList = newExpNodeList
                }
            }
        } else {
            if retTypeList.Len() != 0{
                if retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind(_env) != Ast_TypeInfoKind__DDD{
                    self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("no return value -- need values(%d)", []LnsAny{retTypeList.Len()}))
                }
            }
        }
    }
    return Nodes_ReturnNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), retTypeList, expList)
}
// 78: decl @lune.@base.@TransUnit.TransUnit.analyzeStatementToken
func (self *TransUnit_TransUnit) AnalyzeStatementToken(_env *LnsEnv, token *Types_Token) LnsAny {
    return nil
}
// 84: decl @lune.@base.@TransUnit.TransUnit.analyzeStatement
func (self *TransUnit_TransUnit) analyzeStatement(_env *LnsEnv, termTxt LnsAny) LnsAny {
    self.commentCtrl.FP.Push(_env)
    var token *Types_Token
    token = self.FP.GetTokenNoErr(_env, nil)
    var statement LnsAny
    statement = nil
    if token.Kind == Types_TokenKind__Sheb{
        statement = &Nodes_ShebangNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), token.Txt).Nodes_Node
    }
    if token == Parser_getEofToken(_env){
        return nil
    }
    if Lns_op_not(statement){
        statement = self.FP.analyzeDecl(_env, Ast_AccessMode__Local, false, token, token)
    }
    if Lns_op_not(statement){
        if token.Txt == termTxt{
            if self.commentCtrl.FP.Get_commentList(_env).Len() > 0{
                var commentToken *Types_Token
                commentToken = self.commentCtrl.FP.Get_commentList(_env).GetAt(1).(Types_TokenDownCast).ToTypes_Token()
                statement = &Nodes_BlankLineNode_create(_env, self.NodeManager, commentToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), 0).Nodes_Node
            }
            self.FP.Pushback(_env)
            self.commentCtrl.FP.Pop(_env)
            return statement
        } else if Lns_isCondTrue( Ast_txt2AccessMode(_env, token.Txt)){
            var accessMode LnsInt
            
            {
                _accessMode := Ast_txt2AccessMode(_env, token.Txt)
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
                nextToken = self.FP.getToken(_env, nil)
            }
            statement = self.FP.analyzeDecl(_env, accessMode, staticFlag, token, nextToken)
            if Lns_op_not(statement){
                self.FP.AddErrMess(_env, nextToken.Pos, _env.GetVM().String_format("This token is illegal -- %s", []LnsAny{nextToken.Txt}))
            }
        } else if token.Txt == "{"{
            self.FP.Pushback(_env)
            statement = &self.FP.analyzeBlock(_env, Nodes_BlockKind__Block, TransUnit_TentativeMode__Simple, nil, nil).Nodes_Node
        } else if token.Txt == "super"{
            statement = self.FP.analyzeSuper(_env, token)
        } else if token.Txt == "__asyncLock"{
            statement = self.FP.analyzeAsyncLock(_env, token, Nodes_LockKind__AsyncLock)
        } else if token.Txt == "__luago"{
            statement = self.FP.analyzeAsyncLock(_env, token, Nodes_LockKind__LuaGo)
        } else if token.Txt == "__luaLock"{
            statement = self.FP.analyzeAsyncLock(_env, token, Nodes_LockKind__LuaLock)
        } else if token.Txt == "__luaDepend"{
            statement = self.FP.analyzeAsyncLock(_env, token, Nodes_LockKind__LuaDepend)
        } else if token.Txt == "if"{
            statement = self.FP.analyzeIf(_env, token)
        } else if token.Txt == "when"{
            statement = self.FP.analyzeWhen(_env, token)
        } else if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Txt == "switch") ||
            _env.SetStackVal( token.Txt == "_switch") ).(bool){
            statement = &self.FP.analyzeSwitch(_env, token).Nodes_Node
        } else if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Txt == "match") ||
            _env.SetStackVal( token.Txt == "_match") ).(bool){
            statement = &self.FP.analyzeMatch(_env, token).Nodes_Node
        } else if token.Txt == "while"{
            statement = &self.FP.analyzeWhile(_env, token).Nodes_Node
        } else if token.Txt == "repeat"{
            statement = &self.FP.analyzeRepeat(_env, token).Nodes_Node
        } else if token.Txt == "for"{
            statement = &self.FP.analyzeFor(_env, token).Nodes_Node
        } else if token.Txt == "apply"{
            statement = &self.FP.analyzeApply(_env, token).Nodes_Node
        } else if token.Txt == "foreach"{
            statement = self.FP.analyzeForeach(_env, token, false)
        } else if token.Txt == "forsort"{
            statement = self.FP.analyzeForeach(_env, token, true)
        } else if token.Txt == "return"{
            statement = &self.FP.analyzeReturn(_env, token).Nodes_Node
        } else if token.Txt == "break"{
            self.FP.checkNextToken(_env, ";")
            statement = &Nodes_BreakNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)})).Nodes_Node
            if Lns_op_not(self.FP.Get_curNsInfo(_env).FP.CanBreak(_env)){
                self.FP.AddErrMess(_env, token.Pos, "no loop syntax.")
            }
        } else if token.Txt == "unwrap"{
            statement = self.FP.analyzeUnwrap(_env, token)
        } else if token.Txt == "sync"{
            statement = self.FP.analyzeDeclVar(_env, Nodes_DeclVarMode__Sync, Ast_AccessMode__Local, token)
        } else if token.Txt == "__scope"{
            statement = &self.FP.analyzeScope(_env, token).Nodes_Node
        } else if token.Txt == "_lune_control"{
            {
                __exp := self.FP.analyzeLuneControl(_env, token)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Nodes_LuneControlNode)
                    statement = &_exp.Nodes_Node
                } else {
                    statement = self.FP.createNoneNode(_env, token.Pos)
                }
            }
        } else if token.Txt == ";"{
            statement = self.FP.createNoneNode(_env, token.Pos)
        } else if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( token.Txt == ",,") ||
            _env.SetStackVal( token.Txt == ",,,") ||
            _env.SetStackVal( token.Txt == ",,,,") ).(bool){
            self.FP.Error(_env, _env.GetVM().String_format("illegal macro op -- %s", []LnsAny{token.Txt}))
        } else { 
            statement = self.FP.AnalyzeStatementToken(_env, token)
            if Lns_op_not(statement){
                self.FP.Pushback(_env)
                self.accessSymbolSetQueue.FP.Push(_env)
                var exp *Nodes_Node
                exp = self.FP.analyzeExp(_env, true, false, true, nil, nil)
                var nextToken *Types_Token
                nextToken = self.FP.getToken(_env, nil)
                if nextToken.Txt == ","{
                    var expList *Nodes_ExpListNode
                    expList = self.FP.analyzeExpList(_env, true, true, true, exp, nil, nil)
                    exp = self.FP.analyzeExpOp2(_env, token, &expList.Nodes_Node, nil)
                    nextToken = self.FP.getToken(_env, nil)
                }
                self.FP.checkToken(_env, nextToken, ";")
                {
                    _setNode := Nodes_ExpSetValNodeDownCastF(exp.FP)
                    if !Lns_IsNil( _setNode ) {
                        setNode := _setNode.(*Nodes_ExpSetValNode)
                        var set *LnsSet
                        set = NewLnsSet([]LnsAny{})
                        for _, _symbol := range( setNode.FP.Get_LeftSymList(_env).Items ) {
                            symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            set.Add(Ast_SymbolInfo2Stem(symbol.FP.GetOrg(_env)))
                        }
                        self.accessSymbolSetQueue.FP.Pop(_env, set)
                    } else {
                        self.accessSymbolSetQueue.FP.Pop(_env, nil)
                    }
                }
                statement = &Nodes_StmtExpNode_create(_env, self.NodeManager, exp.FP.Get_pos(_env), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), exp).Nodes_Node
            }
        }
    }
    if statement != nil{
        statement_2091 := statement.(*Nodes_Node)
        if Lns_op_not(statement_2091.FP.CanBeStatement(_env)){
            self.FP.AddErrMess(_env, statement_2091.FP.Get_pos(_env), _env.GetVM().String_format("This node can't be statement. -- %s", []LnsAny{Nodes_getNodeKindName(_env, statement_2091.FP.Get_kind(_env))}))
        }
    }
    self.commentCtrl.FP.Pop(_env)
    return statement
}
// 38: decl @lune.@base.@TransUnit.TransUnitRunner.run
func (self *TransUnit_TransUnitRunner) Run(_env *LnsEnv) {
    self.FP.setup(_env, self.srcTranUnit)
    {
        __exp := self.alreadyToSetup
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(G__lns_Flag)
            _exp.Set(_env)
        }
    }
    self.resultMap = self.FP.processFuncBlockInfo(_env, self.funcBlockCtl, self.Parser.FP.GetStreamName(_env))
}
// 48: decl @lune.@base.@TransUnit.TransUnitRunner.waitToSetup
func (self *TransUnit_TransUnitRunner) WaitToSetup(_env *LnsEnv) {
    {
        __exp := self.alreadyToSetup
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(G__lns_Flag)
            _exp.Wait(_env)
        }
    }
}
// 54: decl @lune.@base.@TransUnit.TransUnitRunner.get
func (self *TransUnit_TransUnitRunner) Get(_env *LnsEnv) *LnsMap {
    LnsJoin(_env, self.FP)
    return self.resultMap
}
// 59: decl @lune.@base.@TransUnit.TransUnitRunner.analyzeStatementToken
func (self *TransUnit_TransUnitRunner) AnalyzeStatementToken(_env *LnsEnv, token *Types_Token) LnsAny {
    if TransUnit_unsupportStatement.Has(token.Txt){
        self.FP.ErrorAt(_env, token.Pos, _env.GetVM().String_format("unsupport the '%s' statement on the multi phase ast. ", []LnsAny{token.Txt}) + "please declare '_lune_control single_phase_ast'")
    }
    return nil
}
// 133: decl @lune.@base.@TransUnit.TransUnitCtrl.analyzeLuneControlToken
func (self *TransUnit_TransUnitCtrl) AnalyzeLuneControlToken(_env *LnsEnv, firstToken *Types_Token,controlToken *Types_Token) LnsAny {
    var pragma LnsAny
    if _switch0 := controlToken.Txt; _switch0 == "disable_mut_control" {
        self.ValidMutControl = false
        self.Modifier.FP.Set_validMutControl(_env, false)
        pragma = LuneControl_Pragma__disable_mut_control_Obj
    } else if _switch0 == "ignore_symbol_" {
        self.IgnoreToCheckSymbol_ = true
        pragma = LuneControl_Pragma__ignore_symbol__Obj
    } else if _switch0 == "limit_conv_code" {
        var codeSet *LnsSet
        codeSet = NewLnsSet([]LnsAny{})
        for  {
            var token *Types_Token
            token = self.FP.getToken(_env, nil)
            if token.Txt == ";"{
                self.FP.Pushback(_env)
                break
            }
            {
                _code := LuneControl_Code__from(_env, token.Txt)
                if !Lns_IsNil( _code ) {
                    code := _code.(string)
                    codeSet.Add(code)
                } else {
                    self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("illegal code -- '%s'", []LnsAny{token.Txt}))
                }
            }
        }
        pragma = &LuneControl_Pragma__limit_conv_code{codeSet}
    } else if _switch0 == "use_async" {
        pragma = LuneControl_Pragma__use_async_Obj
    } else if _switch0 == "default_async_func" {
        pragma = LuneControl_Pragma__default_async_func_Obj
        self.defaultAsyncMode = TransUnit_DefaultAsyncMode__AsyncFunc
    } else if _switch0 == "default_async_all" {
        pragma = LuneControl_Pragma__default_async_all_Obj
        self.defaultAsyncMode = TransUnit_DefaultAsyncMode__AsyncAll
    } else if _switch0 == "use_macro_special_var" {
        pragma = LuneControl_Pragma__use_macro_special_var_Obj
        self.analyzePhase = TransUnit_AnalyzePhase__Main
    } else if _switch0 == "single_phase_ast" {
        pragma = LuneControl_Pragma__single_phase_ast_Obj
        self.analyzePhase = TransUnit_AnalyzePhase__Main
    } else {
        pragma = nil
    }
    return pragma
}
// 191: decl @lune.@base.@TransUnit.TransUnitCtrl.analyzeImportFor
func (self *TransUnit_TransUnitCtrl) analyzeImportFor(_env *LnsEnv, pos Types_Position,modulePath string,assignName string,assigned bool,lazyLoad LnsInt) *Nodes_ImportNode {
    var backupScope *Ast_Scope
    backupScope = self.FP.Get_scope(_env)
    self.FP.SetScope(_env, self.TopScope, TransUnitIF_SetNSInfo__FromScope)
    var macroMode string
    var nearCode LnsAny
    if self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.Get_mode(_env) != Nodes_MacroMode__None{
        macroMode = Nodes_MacroMode_getTxt( self.MacroCtrl.FP.Get_analyzeInfo(_env).FP.Get_mode(_env))
        nearCode = self.Parser.FP.GetNearCode(_env)
    } else { 
        macroMode = ""
        nearCode = nil
    }
    var importObj *Import_Import
    
    {
        _importObj := self.importCtrl
        if _importObj == nil{
            importObj = NewImport_Import(_env, self.FP.GetLatestPos(_env), self.importModuleInfo, self.moduleType, self.MacroCtrl, self.TypeNameCtrl, self.ImportedAliasMap, self.BaseDir, self.ValidMutControl)
            self.importCtrl = importObj
        } else {
            importObj = _importObj.(*Import_Import)
        }
    }
    var moduleLoaderParam *Import_ModuleLoaderParam
    moduleLoaderParam = NewImport_ModuleLoaderParam(_env, self.Ctrl_info, self.ProcessInfo, self.FP.GetLatestPos(_env), macroMode, nearCode, self.ValidMutControl, self.macroEval)
    var moduleLoader *Import_ModuleLoader
    moduleLoader = importObj.FP.ProcessImport(_env, modulePath, moduleLoaderParam)
    var exportInfo *FrontInterface_ExportInfo
    Lns_LockEnvSync( _env, 235, func () {
        var work LnsAny
        var err string
        work,err = importObj.FP.LoadModuleInfo(_env, moduleLoader)
        if work != nil{
            work_2138 := work.(*FrontInterface_ExportInfo)
            exportInfo = work_2138
        } else {
            self.FP.Error(_env, err)
        }
    })
    for _, _symbol := range( exportInfo.FP.Get_globalSymbolList(_env).Items ) {
        symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        self.GlobalScope.FP.AddSymbolInfo(_env, self.ProcessInfo, symbol)
    }
    self.FP.SetScope(_env, backupScope, TransUnitIF_SetNSInfo__FromScope)
    var provideInfo *FrontInterface_ModuleProvideInfo
    provideInfo = exportInfo.FP.Get_provideInfo(_env)
    var moduleTypeInfo *Ast_TypeInfo
    moduleTypeInfo = provideInfo.FP.Get_typeInfo(_env)
    self.FP.Get_scope(_env).FP.AddModule(_env, moduleTypeInfo, exportInfo.FP.Assign(_env, assignName).FP)
    var moduleSymbolKind LnsInt
    moduleSymbolKind = provideInfo.FP.Get_symbolKind(_env)
    var moduleSymbolInfo LnsAny
    var shadowing LnsAny
    moduleSymbolInfo,shadowing = self.FP.Get_scope(_env).FP.Add(_env, self.ProcessInfo, moduleSymbolKind, false, false, assignName, pos, moduleTypeInfo, Ast_AccessMode__Local, true, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( provideInfo.FP.Get_mutable(_env)) &&
        _env.SetStackVal( Ast_MutMode__Mut) ||
        _env.SetStackVal( Ast_MutMode__IMut) ).(LnsInt), true, lazyLoad != Nodes_LazyLoad__Off)
    if shadowing != nil{
        shadowing_2148 := shadowing.(*Ast_SymbolInfo)
        var err bool
        err = shadowing_2148.FP.Get_typeInfo(_env) != moduleTypeInfo
        self.FP.ErrorShadowingOp(_env, pos, shadowing_2148, err)
        if err{
            self.FP.Error(_env, _env.GetVM().String_format("failed to import -- %s", []LnsAny{modulePath}))
        }
        moduleSymbolInfo = shadowing_2148
    }
    if moduleSymbolInfo != nil{
        moduleSymbolInfo_2152 := moduleSymbolInfo.(*Ast_SymbolInfo)
        var info *Nodes_ImportInfo
        info = NewNodes_ImportInfo(_env, pos, modulePath, lazyLoad, assignName, assigned, moduleSymbolInfo_2152, moduleTypeInfo)
        return Nodes_ImportNode_create(_env, self.NodeManager, pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(moduleTypeInfo)}), info)
    }
    self.FP.Error(_env, _env.GetVM().String_format("failed to import -- %s", []LnsAny{modulePath}))
// insert a dummy
    return nil
}
// 288: decl @lune.@base.@TransUnit.TransUnitCtrl.analyzeImport
func (self *TransUnit_TransUnitCtrl) analyzeImport(_env *LnsEnv, opeToken *Types_Token) *Nodes_ImportNode {
    var lazyLoad LnsInt
    if self.FP.getToken(_env, nil).Txt == "."{
        var modeToken *Types_Token
        modeToken = self.FP.getToken(_env, nil)
        if _switch0 := modeToken.Txt; _switch0 == "l" {
            lazyLoad = Nodes_LazyLoad__On
        } else if _switch0 == "d" {
            lazyLoad = Nodes_LazyLoad__Off
        } else {
            lazyLoad = Nodes_LazyLoad__Off
            self.FP.Error(_env, _env.GetVM().String_format("illegal import mode -- %s", []LnsAny{modeToken.Txt}))
        }
    } else { 
        self.FP.Pushback(_env)
        if self.Ctrl_info.DefaultLazy{
            lazyLoad = Nodes_LazyLoad__Auto
        } else { 
            lazyLoad = Nodes_LazyLoad__Off
        }
    }
    if lazyLoad != Nodes_LazyLoad__Off{
        self.HelperInfo.UseLazyLoad = true
    }
    var moduleToken *Types_Token
    moduleToken = self.FP.getToken(_env, nil)
    var modulePath string
    modulePath = moduleToken.Txt
    var nextToken *Types_Token
    nextToken = moduleToken
    for  {
        nextToken = self.FP.getToken(_env, nil)
        if _switch1 := nextToken.Txt; _switch1 == "." || _switch1 == "/" || _switch1 == ":" {
            var demilit string
            demilit = nextToken.Txt
            nextToken = self.FP.getToken(_env, nil)
            moduleToken = nextToken
            modulePath = _env.GetVM().String_format("%s%s%s", []LnsAny{modulePath, demilit, moduleToken.Txt})
        } else {
            break
        }
    }
    var assignName *Types_Token
    assignName = moduleToken
    var assigned bool
    if nextToken.Txt == "as"{
        assignName = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
        nextToken = self.FP.getToken(_env, nil)
        assigned = true
    } else { 
        assigned = false
    }
    self.FP.checkToken(_env, nextToken, ";")
    var node *Nodes_ImportNode
    Lns_LockEnvSync( _env, 348, func () {
        node = self.FP.analyzeImportFor(_env, opeToken.Pos, modulePath, assignName.Txt, assigned, lazyLoad)
    })
    self.ImportModuleSet.Add(Ast_TypeInfo2Stem(node.FP.Get_expType(_env)))
    return node
}
// 357: decl @lune.@base.@TransUnit.TransUnitCtrl.analyzeTestCase
func (self *TransUnit_TransUnitCtrl) analyzeTestCase(_env *LnsEnv, firstToken *Types_Token) *Nodes_TestCaseNode {
    var newScope *Ast_Scope
    newScope = self.FP.PushScope(_env, Ast_ScopeKind__Other, nil, nil)
    var importNode *Nodes_ImportNode
    Lns_LockEnvSync( _env, 363, func () {
        importNode = self.FP.analyzeImportFor(_env, firstToken.Pos, "lune.base.Testing", "__t", false, Nodes_LazyLoad__Off)
    })
    var nameToken *Types_Token
    nameToken = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
    self.FP.checkNextToken(_env, "(")
    var ctrlToken *Types_Token
    ctrlToken = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
    var ctrlName string
    ctrlName = ctrlToken.Txt
    self.FP.checkNextToken(_env, ")")
    var moduleType *Ast_TypeInfo
    moduleType = importNode.FP.Get_expType(_env)
    var ctrlType *Ast_TypeInfo
    
    {
        _ctrlType := _env.NilAccFin(_env.NilAccPush(moduleType.FP.Get_scope(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetTypeInfoChild(_env, "Ctrl")})/* 377:20 */)
        if _ctrlType == nil{
            self.FP.Error(_env, "not found Testing.Ctrl class")
        } else {
            ctrlType = _ctrlType.(*Ast_TypeInfo)
        }
    }
    self.FP.AddLocalVar(_env, ctrlToken.Pos, true, false, ctrlToken.Txt, ctrlType, Ast_MutMode__IMut, false)
    self.ScopeAccess = Ast_ScopeAccess__Full
    self.inTestBlock = true
    var block *Nodes_BlockNode
    if self.Ctrl_info.Testing{
        block = self.FP.analyzeBlock(_env, Nodes_BlockKind__Test, TransUnit_TentativeMode__Ignore, newScope, nil)
    } else { 
        block = self.FP.skipAndCreateDummyBlock(_env)
    }
    self.inTestBlock = false
    self.ScopeAccess = Ast_ScopeAccess__Normal
    self.FP.PopScope(_env)
    return Nodes_TestCaseNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), nameToken, &importNode.Nodes_Node, ctrlName, block)
}
// 406: decl @lune.@base.@TransUnit.TransUnitCtrl.analyzeTest
func (self *TransUnit_TransUnitCtrl) analyzeTest(_env *LnsEnv, firstToken *Types_Token) *Nodes_Node {
    var nextToken *Types_Token
    nextToken = self.FP.getToken(_env, nil)
    if nextToken.Txt != "{"{
        self.FP.Pushback(_env)
        return &self.FP.analyzeTestCase(_env, firstToken).Nodes_Node
    }
    self.FP.checkToken(_env, nextToken, "{")
    self.inTestBlock = true
    var stmtList *LnsList
    stmtList = NewLnsList([]LnsAny{})
    self.FP.analyzeStatementList(_env, stmtList, false, "}")
    self.FP.checkNextToken(_env, "}")
    self.inTestBlock = false
    return &Nodes_TestBlockNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), stmtList).Nodes_Node
}
// 431: decl @lune.@base.@TransUnit.TransUnitCtrl.analyzeDeclToken
func (self *TransUnit_TransUnitCtrl) AnalyzeDeclToken(_env *LnsEnv, accessMode LnsInt,staticFlag bool,firstToken *Types_Token,token *Types_Token) LnsAny {
    if token.Txt == "__test"{
        return self.FP.analyzeTest(_env, firstToken)
    }
    return nil
}
// 442: decl @lune.@base.@TransUnit.TransUnitCtrl.analyzeSubfile
func (self *TransUnit_TransUnitCtrl) analyzeSubfile(_env *LnsEnv, token *Types_Token) *Nodes_SubfileNode {
    if self.FP.Get_scope(_env) != self.ModuleScope{
        self.FP.Error(_env, "'module' must be top scope.")
    }
    var mode *Types_Token
    mode = self.FP.getToken(_env, nil)
    var moduleName string
    moduleName = ""
    for  {
        var nextToken *Types_Token
        nextToken = self.FP.getToken(_env, nil)
        if nextToken.Txt == ";"{
            break
        }
        if moduleName == ""{
            moduleName = nextToken.Txt
        } else { 
            moduleName = _env.GetVM().String_format("%s%s", []LnsAny{moduleName, nextToken.Txt})
        }
    }
    var usePath LnsAny
    usePath = nil
    if moduleName == ""{
        self.FP.AddErrMess(_env, token.Pos, "illegal subfile")
    } else { 
        if mode.Txt == "use"{
            usePath = moduleName
            if Lns_isCondTrue( FrontInterface_searchModule(_env, moduleName, self.BaseDir, nil)){
                self.subfileList.Insert(moduleName)
            } else { 
                self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("not found subfile -- %s", []LnsAny{moduleName}))
            }
        } else if mode.Txt == "owner"{
            if Lns_car(FrontInterface_getLuaModulePath(_env, self.moduleName, self.BaseDir)).(string) != moduleName{
                self.FP.AddErrMess(_env, token.Pos, _env.GetVM().String_format("illegal owner module -- %s, %s", []LnsAny{moduleName, self.moduleName}))
            }
        } else { 
            self.FP.AddErrMess(_env, mode.Pos, _env.GetVM().String_format("illegal module mode -- %s", []LnsAny{mode.Txt}))
        }
    }
    return Nodes_SubfileNode_create(_env, self.NodeManager, token.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), usePath)
}
// 499: decl @lune.@base.@TransUnit.TransUnitCtrl.analyzeProvide
func (self *TransUnit_TransUnitCtrl) analyzeProvide(_env *LnsEnv, firstToken *Types_Token) *Nodes_ProvideNode {
    var token *Types_Token
    token = self.FP.getSymbolToken(_env, TransUnit_SymbolMode__MustNot_)
    var symbolNode *Nodes_Node
    symbolNode = self.FP.analyzeExpSymbol(_env, firstToken, token, TransUnit_ExpSymbolMode__Symbol, nil, true, false)
    self.FP.checkNextToken(_env, ";")
    var symbolInfoList *LnsList
    symbolInfoList = symbolNode.FP.GetSymbolInfo(_env)
    if symbolInfoList.Len() != 1{
        self.FP.Error(_env, "'provide' must be symbol.")
    }
    var symbolInfo *Ast_SymbolInfo
    symbolInfo = symbolInfoList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
    var node *Nodes_ProvideNode
    node = Nodes_ProvideNode_create(_env, self.NodeManager, firstToken.Pos, self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), symbolInfo)
    if Lns_isCondTrue( self.provideNode){
        self.FP.AddErrMess(_env, firstToken.Pos, "multiple provide")
    }
    self.provideNode = node
    if symbolInfo.FP.Get_accessMode(_env) != Ast_AccessMode__Pub{
        self.FP.AddErrMess(_env, firstToken.Pos, _env.GetVM().String_format("provide variable must be 'pub'.  -- %s", []LnsAny{symbolInfo.FP.Get_accessMode(_env)}))
    }
    return node
}
// 530: decl @lune.@base.@TransUnit.TransUnitCtrl.analyzeStatementToken
func (self *TransUnit_TransUnitCtrl) AnalyzeStatementToken(_env *LnsEnv, token *Types_Token) LnsAny {
    var statement LnsAny
    if token.Txt == "import"{
        statement = &self.FP.analyzeImport(_env, token).Nodes_Node
    } else if token.Txt == "subfile"{
        Lns_LockEnvSync( _env, 536, func () {
            statement = &self.FP.analyzeSubfile(_env, token).Nodes_Node
        })
    } else if token.Txt == "provide"{
        statement = &self.FP.analyzeProvide(_env, token).Nodes_Node
    } else { 
        statement = nil
    }
    return statement
}
// 547: decl @lune.@base.@TransUnit.TransUnitCtrl.addFuncBlockInfoList
func (self *TransUnit_TransUnitCtrl) addFuncBlockInfoList(_env *LnsEnv, funcBlockInfo *TransUnit_FuncBlockInfo) {
    self.totalFuncBlockTokenNum = self.totalFuncBlockTokenNum + funcBlockInfo.FP.Get_tokenList(_env).Len()
    self.funcBlockInfoList.Insert(TransUnit_FuncBlockInfo2Stem(funcBlockInfo))
}
// 553: decl @lune.@base.@TransUnit.TransUnitCtrl.processFuncBlock
func (self *TransUnit_TransUnitCtrl) processFuncBlock(_env *LnsEnv, streamName string) {
    var runnerList *LnsList
    runnerList = NewLnsList([]LnsAny{})
    var resultMap *LnsMap
    resultMap = NewLnsMap( map[LnsAny]LnsAny{})
    var noRunnerList *LnsList
    noRunnerList = NewLnsList([]LnsAny{})
    var noRunnerTypeSet *LnsSet
    noRunnerTypeSet = NewLnsSet([]LnsAny{})
    var lastPartList LnsAny
    lastPartList = nil
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.funcBlockInfoList.Len() < 20) ||
        _env.SetStackVal( self.totalFuncBlockTokenNum < 2000) ).(bool){
        LnsLog(_env, _env.GetVM().String_format("disable processing async AST. %d, %d", []LnsAny{self.funcBlockInfoList.Len(), self.totalFuncBlockTokenNum}))
    } else { 
        var divCount LnsInt
        divCount = self.Ctrl_info.ThreadPerUnitThread
        if divCount > 0{
            self.analyzePhase = TransUnit_AnalyzePhase__Runner
            var maxTokenCount LnsInt
            maxTokenCount = (LnsInt)((LnsReal((self.totalFuncBlockTokenNum / divCount)) * 0.9))
            var offset LnsInt
            offset = 1
            var len LnsInt
            len = self.funcBlockInfoList.Len()
            {
                var _forFrom0 LnsInt = 1
                var _forTo0 LnsInt = divCount
                for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                    managerId := _forWork0
                    var list *LnsList
                    list = NewLnsList([]LnsAny{})
                    if managerId == divCount{
                        {
                            var _forFrom1 LnsInt = offset
                            var _forTo1 LnsInt = len
                            for _forWork1 := _forFrom1; _forWork1 <= _forTo1; _forWork1++ {
                                index := _forWork1
                                var funcBlockInfo *TransUnit_FuncBlockInfo
                                funcBlockInfo = self.funcBlockInfoList.GetAt(index).(TransUnit_FuncBlockInfoDownCast).ToTransUnit_FuncBlockInfo()
                                list.Insert(TransUnit_FuncBlockInfo2Stem(funcBlockInfo))
                            }
                        }
                        lastPartList = list
                    } else { 
                        var tokenCount LnsInt
                        tokenCount = 0
                        for offset <= len {
                            var funcBlockInfo *TransUnit_FuncBlockInfo
                            funcBlockInfo = self.funcBlockInfoList.GetAt(offset).(TransUnit_FuncBlockInfoDownCast).ToTransUnit_FuncBlockInfo()
                            offset = offset + 1
                            list.Insert(TransUnit_FuncBlockInfo2Stem(funcBlockInfo))
                            tokenCount = tokenCount + funcBlockInfo.FP.Get_tokenList(_env).Len()
                            if tokenCount >= maxTokenCount{
                                break
                            }
                        }
                        var runner *TransUnit_TransUnitRunner
                        runner = NewTransUnit_TransUnitRunner(_env, &self.TransUnit_TransUnit, self.ModuleId, self.importModuleInfo, self.macroEval, false, self.moduleName, TransUnit_AnalyzeMode__Compile, nil, self.TargetLuaVer, self.Ctrl_info, self.builtinFunc, list, managerId)
                        if LnsRun(_env, runner.FP, 2, _env.GetVM().String_format("astMain -- %s", []LnsAny{streamName})){
                            runnerList.Insert(TransUnit_TransUnitRunner2Stem(runner))
                        } else { 
                            noRunnerList.Insert(list)
                        }
                    }
                }
            }
        }
    }
    if runnerList.Len() == 0{
        self.analyzePhase = TransUnit_AnalyzePhase__Main
        resultMap = self.FP.processFuncBlockInfo(_env, NewTransUnit_ListFuncBlockCtl(_env, self.funcBlockInfoList).FP, self.Parser.FP.GetStreamName(_env))
    } else { 
        for _, _runner := range( runnerList.Items ) {
            runner := _runner.(TransUnit_TransUnitRunnerDownCast).ToTransUnit_TransUnitRunner()
            runner.FP.WaitToSetup(_env)
        }
        if lastPartList != nil{
            lastPartList_2269 := lastPartList.(*LnsList)
            resultMap = self.FP.processFuncBlockInfo(_env, NewTransUnit_ListFuncBlockCtl(_env, lastPartList_2269).FP, self.Parser.FP.GetStreamName(_env))
            LnsLog(_env, "process lastpart")
        }
        for _, _runner := range( runnerList.Items ) {
            runner := _runner.(TransUnit_TransUnitRunnerDownCast).ToTransUnit_TransUnitRunner()
            var workMap *LnsMap
            workMap = runner.FP.Get(_env)
            self.FP.mergeFrom(_env, &runner.TransUnit_TransUnit, resultMap)
            for _key, _result := range( workMap.Items ) {
                key := _key.(TransUnit_FuncBlockInfoDownCast).ToTransUnit_FuncBlockInfo()
                result := _result.(TransUnit_FuncBlockResultDownCast).ToTransUnit_FuncBlockResult()
                resultMap.Set(key,result)
            }
        }
        LnsLog(_env, "merged AST.")
        for _, _noRunner := range( noRunnerList.Items ) {
            noRunner := _noRunner.(*LnsList)
            var workMap *LnsMap
            workMap = self.FP.processFuncBlockInfo(_env, NewTransUnit_ListFuncBlockCtl(_env, noRunner).FP, self.Parser.FP.GetStreamName(_env))
            for _funcBlockInfo, _ := range( workMap.Items ) {
                funcBlockInfo := _funcBlockInfo.(TransUnit_FuncBlockInfoDownCast).ToTransUnit_FuncBlockInfo()
                noRunnerTypeSet.Add(Ast_TypeInfo2Stem(funcBlockInfo.FP.Get_funcType(_env)))
            }
        }
    }
    for _, _funcBlockInfo := range( self.funcBlockInfoList.Items ) {
        funcBlockInfo := _funcBlockInfo.(TransUnit_FuncBlockInfoDownCast).ToTransUnit_FuncBlockInfo()
        {
            _result := resultMap.Get(funcBlockInfo)
            if !Lns_IsNil( _result ) {
                result := _result.(*TransUnit_FuncBlockResult)
                var declFuncInfo *Nodes_DeclFuncInfo
                declFuncInfo = funcBlockInfo.FP.Get_declFuncInfo(_env)
                declFuncInfo.FP.Set_body(_env, result.FP.Get_body(_env))
                declFuncInfo.FP.Set_has__func__Symbol(_env, result.FP.Get_has_func_sym(_env))
                declFuncInfo.FP.Set_stmtNum(_env, result.FP.Get_stmtNum(_env))
            } else {
                if Lns_op_not(noRunnerTypeSet.Has(Ast_TypeInfo2Stem(funcBlockInfo.FP.Get_funcType(_env)))){
                    Util_err(_env, _env.GetVM().String_format("not found result -- %s", []LnsAny{funcBlockInfo.FP.Get_funcType(_env).FP.GetTxt(_env, nil, nil, nil)}))
                }
            }
        }
    }
}
// 675: decl @lune.@base.@TransUnit.TransUnitCtrl.createAST
func (self *TransUnit_TransUnitCtrl) CreateAST(_env *LnsEnv, parserSrc LnsAny,asyncParse bool,baseDir LnsAny,stdinFile LnsAny,macroFlag bool,moduleName LnsAny,readyExportInfo LnsAny) *AstInfo_ASTInfo {
    __func__ := "@lune.@base.@TransUnit.TransUnitCtrl.createAST"
    var parser *Parser_Parser
    parser = Parser_createParserFrom(_env, parserSrc, asyncParse, stdinFile)
    var streamName string
    streamName = parser.FP.GetStreamName(_env)
    self.stdinFile = stdinFile
    self.BaseDir = baseDir
    Log_log(_env, Log_Level__Log, __func__, 687, Log_CreateMessage(func(_env *LnsEnv) string {
        __func__ := "@lune.@base.@TransUnit.TransUnitCtrl.createAST.<anonymous>"
        return _env.GetVM().String_format("%s start -- %s on %s, %s, %s", []LnsAny{__func__, parser.FP.GetStreamName(_env), baseDir, macroFlag, TransUnit_AnalyzePhase_getTxt( self.analyzePhase)})
    }))
    
    self.moduleName = Lns_unwrapDefault( moduleName, "").(string)
    if moduleName != nil{
        moduleName_2295 := moduleName.(string)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.importModuleInfo.FP.Len(_env) == 0) &&
            _env.SetStackVal( Lns_op_not(self.importModuleInfo.FP.Add(_env, moduleName_2295))) ).(bool)){
            self.FP.Error(_env, _env.GetVM().String_format("already imported -- %s", []LnsAny{moduleName_2295}))
        }
    }
    var moduleTypeInfo *Ast_TypeInfo
    moduleTypeInfo = Ast_headTypeInfo
    Lns_LockEnvSync( _env, 703, func () {
            if moduleName != nil{
                moduleName_2301 := moduleName.(string)
                {
                    _applyForm0, _applyParam0, _applyPrev0 := _env.GetVM().String_gmatch(Lns_car(FrontInterface_getLuaModulePath(_env, moduleName_2301, baseDir)).(string), "[^%.]+")
                    for {
                        _applyWork0 := _applyForm0.(*Lns_luaValue).Call( Lns_2DDD( _applyParam0, _applyPrev0 ) )
                        _applyPrev0 = Lns_getFromMulti(_applyWork0,0)
                        if Lns_IsNil( _applyPrev0 ) { break }
                        txt := _applyPrev0.(string)
                        moduleTypeInfo = self.FP.PushModule(_env, self.ProcessInfo, false, txt, true).FP.Get_typeInfo(_env)
                    }
                }
            }
    })
    self.ModuleScope = self.FP.Get_scope(_env)
    self.FP.Get_scope(_env).FP.AddVar(_env, self.ProcessInfo, Ast_AccessMode__Global, "__mod__", nil, Ast_builtinTypeString, Ast_MutMode__IMut, true)
    self.moduleType = moduleTypeInfo
    self.TypeNameCtrl = NewAst_TypeNameCtrl(_env, moduleTypeInfo)
    self.FP.SetParser(_env, NewParser_DefaultPushbackParser(_env, parser))
    self.FP.Get_scope(_env).FP.AddIgnoredVar(_env, self.ProcessInfo)
    var ast *Nodes_Node
    var exportInfo *Nodes_ExportInfo
    var createExportInfo func(_env *LnsEnv, moduleSymboInfo LnsAny,globalSymbolList *LnsList,processInfo *Ast_ProcessInfo) *Nodes_ExportInfo
    createExportInfo = func(_env *LnsEnv, moduleSymboInfo LnsAny,globalSymbolList *LnsList,processInfo *Ast_ProcessInfo) *Nodes_ExportInfo {
        __func__ := "@lune.@base.@TransUnit.TransUnitCtrl.createAST.createExportInfo"
        var provideInfo *FrontInterface_ModuleProvideInfo
        if moduleSymboInfo != nil{
            moduleSymboInfo_2314 := moduleSymboInfo.(*Ast_SymbolInfo)
            provideInfo = NewFrontInterface_ModuleProvideInfo(_env, moduleSymboInfo_2314.FP.Get_typeInfo(_env), moduleSymboInfo_2314.FP.Get_kind(_env), moduleSymboInfo_2314.FP.Get_mutable(_env))
        } else {
            provideInfo = NewFrontInterface_ModuleProvideInfo(_env, moduleTypeInfo, Ast_SymbolKind__Typ, false)
        }
        var importedAliasMap *LnsMap
        importedAliasMap = NewLnsMap( map[LnsAny]LnsAny{})
        for _, _node := range( self.NodeManager.FP.GetAliasNodeList(_env).Items ) {
            node := _node.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
            importedAliasMap.Set(node.FP.Get_typeInfo(_env),Ast_AliasTypeInfoDownCastF(node.FP.Get_expType(_env).FP))
        }
        var workExportInfo *Nodes_ExportInfo
        workExportInfo = NewNodes_ExportInfo(_env, moduleTypeInfo, provideInfo, processInfo, globalSymbolList, importedAliasMap, self.ModuleId, self.moduleName, moduleTypeInfo.FP.Get_rawTxt(_env), streamName, NewLnsMap( map[LnsAny]LnsAny{}), self.MacroCtrl.FP.Get_declPubMacroInfoMap(_env))
        LnsLog(_env, _env.GetVM().String_format("ready meta -- %d, %d", []LnsAny{self.funcBlockInfoList.Len(), self.totalFuncBlockTokenNum}))
        Log_log(_env, Log_Level__Log, __func__, 761, Log_CreateMessage(func(_env *LnsEnv) string {
            return _env.GetVM().String_format("ready meta -- %s, %d, %s, %s", []LnsAny{streamName, self.Parser.FP.GetUsedTokenListLen(_env), moduleTypeInfo, moduleTypeInfo.FP.Get_scope(_env)})
        }))
        
        if readyExportInfo != nil{
            readyExportInfo_2323 := readyExportInfo.(TransUnit_ReadyExportInfo)
            readyExportInfo_2323(_env, &workExportInfo.FrontInterface_ExportInfo)
        }
        return workExportInfo
    }
    var lastStatement LnsAny
    lastStatement = nil
    if macroFlag{
        ast = &self.FP.analyzeBlock(_env, Nodes_BlockKind__Macro, TransUnit_TentativeMode__Ignore, nil, nil).Nodes_Node
        exportInfo = createExportInfo(_env, nil, NewLnsList([]LnsAny{}), self.ProcessInfo)
    } else { 
        var children *LnsList
        children = NewLnsList([]LnsAny{})
        var lastLineNo LnsInt
        lastStatement, lastLineNo = self.FP.analyzeStatementList(_env, children, false, nil)
        var statement *Nodes_BlankLineNode
        statement = Nodes_BlankLineNode_create(_env, self.NodeManager, self.FP.CreatePosition(_env, lastLineNo + 1, 0), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), 0)
        statement.FP.AddComment(_env, self.commentCtrl.FP.Get_commentList(_env))
        self.commentCtrl.FP.Clear(_env)
        children.Insert(Nodes_BlankLineNode2Stem(statement))
        var token *Types_Token
        token = self.FP.GetTokenNoErr(_env, nil)
        if token != Parser_getEofToken(_env){
            self.FP.Error(_env, _env.GetVM().String_format("%s:%d:%d:(%s) not eof -- %s", []LnsAny{self.Parser.FP.GetStreamName(_env), token.Pos.LineNo, token.Pos.Column, Types_TokenKind_getTxt( token.Kind), token.Txt}))
        }
        for _, _subModule := range( self.subfileList.Items ) {
            subModule := _subModule.(string)
            var file string
            Lns_LockEnvSync( _env, 802, func () {
                {
                    __exp := FrontInterface_searchModule(_env, subModule, self.BaseDir, nil)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(string)
                        file = _exp
                    } else {
                        self.FP.Error(_env, _env.GetVM().String_format("not found subfile -- %s", []LnsAny{subModule}))
                    }
                }
            })
            if self.FP.Get_scope(_env) != self.ModuleScope{
                self.FP.Error(_env, "scope does not close")
            }
            var subParser *Parser_StreamParser
            subParser = Parser_StreamParser_create(_env, &Types_ParserSrc__LnsPath{file, subModule, nil}, true, self.stdinFile, nil)
            self.FP.SetParser(_env, NewParser_DefaultPushbackParser(_env, &subParser.Parser_Parser))
            lastStatement = self.FP.analyzeStatementListSubfile(_env, children)
            token = self.FP.GetTokenNoErr(_env, nil)
            if token != Parser_getEofToken(_env){
                Util_err(_env, _env.GetVM().String_format("unknown:%d:%d:(%s) %s", []LnsAny{token.Pos.LineNo, token.Pos.Column, Types_TokenKind_getTxt( token.Kind), token.Txt}))
            }
        }
        var globalSymbolList *LnsList
        globalSymbolList = NewLnsList([]LnsAny{})
        for _, _node := range( children.Items ) {
            node := _node.(Nodes_NodeDownCast).ToNodes_Node()
            {
                _workNode := Nodes_DeclVarNodeDownCastF(node.FP)
                if !Lns_IsNil( _workNode ) {
                    workNode := _workNode.(*Nodes_DeclVarNode)
                    for _, _symbolInfo := range( workNode.FP.Get_symbolInfoList(_env).Items ) {
                        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                        if symbolInfo.FP.Get_accessMode(_env) == Ast_AccessMode__Global{
                            globalSymbolList.Insert(Ast_SymbolInfo2Stem(symbolInfo))
                        }
                    }
                }
            }
            {
                _workNode := Nodes_DeclFuncNodeDownCastF(node.FP)
                if !Lns_IsNil( _workNode ) {
                    workNode := _workNode.(*Nodes_DeclFuncNode)
                    if workNode.FP.Get_declInfo(_env).FP.Get_accessMode(_env) == Ast_AccessMode__Global{
                        {
                            _symbolInfo := workNode.FP.Get_declInfo(_env).FP.Get_symbol(_env)
                            if !Lns_IsNil( _symbolInfo ) {
                                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                                globalSymbolList.Insert(Ast_SymbolInfo2Stem(symbolInfo))
                            }
                        }
                    }
                }
            }
        }
        var moduleSymboInfo LnsAny
        {
            __exp := self.provideNode
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(*Nodes_ProvideNode)
                if lastStatement != &_exp.Nodes_Node{
                    self.FP.AddErrMess(_env, _exp.FP.Get_pos(_env), "'provide' must be last.")
                }
                moduleSymboInfo = _exp.FP.Get_symbol(_env)
            } else {
                moduleSymboInfo = nil
            }
        }
        exportInfo = createExportInfo(_env, moduleSymboInfo, globalSymbolList, self.ProcessInfo.FP.Duplicate(_env))
        self.FP.processFuncBlock(_env, streamName)
        self.FP.checkOverriededMethodOfAllClass(_env)
        var rootNode *Nodes_RootNode
        rootNode = Nodes_RootNode_create(_env, self.NodeManager, self.FP.CreatePosition(_env, 0, 0), self.inTestBlock, self.MacroCtrl.FP.IsInAnalyzeArgMode(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}), children, self.ModuleScope, self.GlobalScope, self.MacroCtrl.FP.Get_useModuleMacroSet(_env), self.ModuleId, self.ProcessInfo, moduleTypeInfo, self.provideNode, self.HelperInfo, self.NodeManager, Lns_unwrapDefault( _env.NilAccFin(_env.NilAccPush(self.importCtrl) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Import_Import).FP.Get_importModule2ExportInfo(_env)})), NewLnsMap( map[LnsAny]LnsAny{})).(*LnsMap), self.MacroCtrl.FP.Get_declPubMacroInfoMap(_env), self.TypeId2ClassMap)
        ast = &rootNode.Nodes_Node
        TransUnit_ClosureFun_checkList_1_(_env, self.closureFunList)
    }
    if moduleName != nil{
        moduleName_2362 := moduleName.(string)
        for range( Util_splitStr(_env, moduleName_2362, "[^%.]+").Items ) {
            self.FP.PopModule(_env)
        }
    }
    {
        __forsortCollection0 := TransUnitCtrl_createAST__createId2proto_2_(_env, self.NsInfoMap)
        __forsortSorted0 := __forsortCollection0.CreateKeyListInt()
        __forsortSorted0.Sort( _env, LnsItemKindInt, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            nsInfo := __forsortCollection0.Items[ ___forsortKey0 ].(TransUnitIF_NSInfoDownCast).ToTransUnitIF_NSInfo()
            if nsInfo.FP.Get_nobody(_env){
                var protoType *Ast_TypeInfo
                protoType = nsInfo.FP.Get_typeInfo(_env)
                var mess string
                if _switch0 := protoType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Class || _switch0 == Ast_TypeInfoKind__IF {
                    mess = _env.GetVM().String_format("This class doesn't have body. -- %s", []LnsAny{protoType.FP.GetTxt(_env, nil, nil, nil)})
                } else {
                    mess = _env.GetVM().String_format("This function doesn't have body. -- %s", []LnsAny{protoType.FP.GetTxt(_env, nil, nil, nil)})
                }
                self.FP.AddErrMess(_env, Lns_unwrap( _env.NilAccFin(_env.NilAccPush(self.NsInfoMap.Get(protoType)) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*TransUnitIF_NSInfo).FP.Get_pos(_env)}))).(Types_Position), mess)
            }
        }
    }
    for _, _mess := range( TransUnitIF_sortMess(_env, self.WarnMessList).Items ) {
        mess := _mess.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
        Util_errorLog(_env, mess.FP.Get_mess(_env))
    }
    if self.ErrMessList.Len() > 0{
        for _, _mess := range( TransUnitIF_sortMess(_env, self.ErrMessList).Items ) {
            mess := _mess.(TransUnitIF_ErrMessDownCast).ToTransUnitIF_ErrMess()
            Util_errorLog(_env, mess.FP.Get_mess(_env))
        }
        Util_err(_env, "has error")
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.Ctrl_info.StopByWarning) &&
        _env.SetStackVal( self.WarnMessList.Len() > 0) ).(bool)){
        Util_err(_env, "has error")
    }
    if _switch1 := self.analyzeMode; _switch1 == TransUnit_AnalyzeMode__Diag || _switch1 == TransUnit_AnalyzeMode__Complete || _switch1 == TransUnit_AnalyzeMode__Inquire {
        _env.GetVM().OS_exit(0)
    }
    return NewAstInfo_ASTInfo(_env, ast, &exportInfo.FrontInterface_ExportInfo, parser.FP.GetStreamName(_env), self.builtinFunc)
}
