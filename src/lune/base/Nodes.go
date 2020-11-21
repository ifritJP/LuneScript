// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_Nodes bool
var Nodes__mod__ string
// decl enum -- BreakKind 
const Nodes_BreakKind__Break = 1
const Nodes_BreakKind__NeverRet = 3
const Nodes_BreakKind__None = 0
const Nodes_BreakKind__Return = 2
var Nodes_BreakKindList_ = NewLnsList( []LnsAny {
  Nodes_BreakKind__None,
  Nodes_BreakKind__Break,
  Nodes_BreakKind__Return,
  Nodes_BreakKind__NeverRet,
})
func Nodes_BreakKind_get__allList() *LnsList{
    return Nodes_BreakKindList_
}
var Nodes_BreakKindMap_ = map[LnsInt]string {
  Nodes_BreakKind__Break: "BreakKind.Break",
  Nodes_BreakKind__NeverRet: "BreakKind.NeverRet",
  Nodes_BreakKind__None: "BreakKind.None",
  Nodes_BreakKind__Return: "BreakKind.Return",
}
func Nodes_BreakKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_BreakKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_BreakKind_getTxt(arg1 LnsInt) string {
    return Nodes_BreakKindMap_[arg1];
}
// decl enum -- CheckBreakMode 
const Nodes_CheckBreakMode__IgnoreFlow = 2
const Nodes_CheckBreakMode__IgnoreFlowReturn = 3
const Nodes_CheckBreakMode__Normal = 0
const Nodes_CheckBreakMode__Return = 1
var Nodes_CheckBreakModeList_ = NewLnsList( []LnsAny {
  Nodes_CheckBreakMode__Normal,
  Nodes_CheckBreakMode__Return,
  Nodes_CheckBreakMode__IgnoreFlow,
  Nodes_CheckBreakMode__IgnoreFlowReturn,
})
func Nodes_CheckBreakMode_get__allList() *LnsList{
    return Nodes_CheckBreakModeList_
}
var Nodes_CheckBreakModeMap_ = map[LnsInt]string {
  Nodes_CheckBreakMode__IgnoreFlow: "CheckBreakMode.IgnoreFlow",
  Nodes_CheckBreakMode__IgnoreFlowReturn: "CheckBreakMode.IgnoreFlowReturn",
  Nodes_CheckBreakMode__Normal: "CheckBreakMode.Normal",
  Nodes_CheckBreakMode__Return: "CheckBreakMode.Return",
}
func Nodes_CheckBreakMode__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_CheckBreakModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_CheckBreakMode_getTxt(arg1 LnsInt) string {
    return Nodes_CheckBreakModeMap_[arg1];
}
// decl enum -- NodeVisitMode 
const Nodes_NodeVisitMode__Child = 0
const Nodes_NodeVisitMode__End = 2
const Nodes_NodeVisitMode__Next = 1
var Nodes_NodeVisitModeList_ = NewLnsList( []LnsAny {
  Nodes_NodeVisitMode__Child,
  Nodes_NodeVisitMode__Next,
  Nodes_NodeVisitMode__End,
})
func Nodes_NodeVisitMode_get__allList() *LnsList{
    return Nodes_NodeVisitModeList_
}
var Nodes_NodeVisitModeMap_ = map[LnsInt]string {
  Nodes_NodeVisitMode__Child: "NodeVisitMode.Child",
  Nodes_NodeVisitMode__End: "NodeVisitMode.End",
  Nodes_NodeVisitMode__Next: "NodeVisitMode.Next",
}
func Nodes_NodeVisitMode__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_NodeVisitModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_NodeVisitMode_getTxt(arg1 LnsInt) string {
    return Nodes_NodeVisitModeMap_[arg1];
}
// decl enum -- LazyLoad 
const Nodes_LazyLoad__Auto = 2
const Nodes_LazyLoad__Off = 0
const Nodes_LazyLoad__On = 1
var Nodes_LazyLoadList_ = NewLnsList( []LnsAny {
  Nodes_LazyLoad__Off,
  Nodes_LazyLoad__On,
  Nodes_LazyLoad__Auto,
})
func Nodes_LazyLoad_get__allList() *LnsList{
    return Nodes_LazyLoadList_
}
var Nodes_LazyLoadMap_ = map[LnsInt]string {
  Nodes_LazyLoad__Auto: "LazyLoad.Auto",
  Nodes_LazyLoad__Off: "LazyLoad.Off",
  Nodes_LazyLoad__On: "LazyLoad.On",
}
func Nodes_LazyLoad__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_LazyLoadMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_LazyLoad_getTxt(arg1 LnsInt) string {
    return Nodes_LazyLoadMap_[arg1];
}
// decl enum -- BlockKind 
const Nodes_BlockKind__Apply = 8
const Nodes_BlockKind__Block = 13
const Nodes_BlockKind__Default = 12
const Nodes_BlockKind__Else = 2
const Nodes_BlockKind__Elseif = 1
const Nodes_BlockKind__For = 7
const Nodes_BlockKind__Foreach = 9
const Nodes_BlockKind__Func = 11
const Nodes_BlockKind__If = 0
const Nodes_BlockKind__IfUnwrap = 17
const Nodes_BlockKind__LetUnwrap = 15
const Nodes_BlockKind__LetUnwrapThenDo = 16
const Nodes_BlockKind__Macro = 14
const Nodes_BlockKind__Match = 5
const Nodes_BlockKind__Repeat = 6
const Nodes_BlockKind__Switch = 4
const Nodes_BlockKind__Test = 19
const Nodes_BlockKind__When = 18
const Nodes_BlockKind__While = 3
var Nodes_BlockKindList_ = NewLnsList( []LnsAny {
  Nodes_BlockKind__If,
  Nodes_BlockKind__Elseif,
  Nodes_BlockKind__Else,
  Nodes_BlockKind__While,
  Nodes_BlockKind__Switch,
  Nodes_BlockKind__Match,
  Nodes_BlockKind__Repeat,
  Nodes_BlockKind__For,
  Nodes_BlockKind__Apply,
  Nodes_BlockKind__Foreach,
  Nodes_BlockKind__Macro,
  Nodes_BlockKind__Func,
  Nodes_BlockKind__Default,
  Nodes_BlockKind__Block,
  Nodes_BlockKind__Macro,
  Nodes_BlockKind__LetUnwrap,
  Nodes_BlockKind__LetUnwrapThenDo,
  Nodes_BlockKind__IfUnwrap,
  Nodes_BlockKind__When,
  Nodes_BlockKind__Test,
})
func Nodes_BlockKind_get__allList() *LnsList{
    return Nodes_BlockKindList_
}
var Nodes_BlockKindMap_ = map[LnsInt]string {
  Nodes_BlockKind__Apply: "BlockKind.Apply",
  Nodes_BlockKind__Block: "BlockKind.Block",
  Nodes_BlockKind__Default: "BlockKind.Default",
  Nodes_BlockKind__Else: "BlockKind.Else",
  Nodes_BlockKind__Elseif: "BlockKind.Elseif",
  Nodes_BlockKind__For: "BlockKind.For",
  Nodes_BlockKind__Foreach: "BlockKind.Foreach",
  Nodes_BlockKind__Func: "BlockKind.Func",
  Nodes_BlockKind__If: "BlockKind.If",
  Nodes_BlockKind__IfUnwrap: "BlockKind.IfUnwrap",
  Nodes_BlockKind__LetUnwrap: "BlockKind.LetUnwrap",
  Nodes_BlockKind__LetUnwrapThenDo: "BlockKind.LetUnwrapThenDo",
  Nodes_BlockKind__Macro: "BlockKind.Macro",
  Nodes_BlockKind__Match: "BlockKind.Match",
  Nodes_BlockKind__Repeat: "BlockKind.Repeat",
  Nodes_BlockKind__Switch: "BlockKind.Switch",
  Nodes_BlockKind__Test: "BlockKind.Test",
  Nodes_BlockKind__When: "BlockKind.When",
  Nodes_BlockKind__While: "BlockKind.While",
}
func Nodes_BlockKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_BlockKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_BlockKind_getTxt(arg1 LnsInt) string {
    return Nodes_BlockKindMap_[arg1];
}
// decl enum -- ScopeKind 
const Nodes_ScopeKind__Root = 0
var Nodes_ScopeKindList_ = NewLnsList( []LnsAny {
  Nodes_ScopeKind__Root,
})
func Nodes_ScopeKind_get__allList() *LnsList{
    return Nodes_ScopeKindList_
}
var Nodes_ScopeKindMap_ = map[LnsInt]string {
  Nodes_ScopeKind__Root: "ScopeKind.Root",
}
func Nodes_ScopeKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_ScopeKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_ScopeKind_getTxt(arg1 LnsInt) string {
    return Nodes_ScopeKindMap_[arg1];
}
// decl enum -- IfKind 
const Nodes_IfKind__Else = 2
const Nodes_IfKind__ElseIf = 1
const Nodes_IfKind__If = 0
var Nodes_IfKindList_ = NewLnsList( []LnsAny {
  Nodes_IfKind__If,
  Nodes_IfKind__ElseIf,
  Nodes_IfKind__Else,
})
func Nodes_IfKind_get__allList() *LnsList{
    return Nodes_IfKindList_
}
var Nodes_IfKindMap_ = map[LnsInt]string {
  Nodes_IfKind__Else: "IfKind.Else",
  Nodes_IfKind__ElseIf: "IfKind.ElseIf",
  Nodes_IfKind__If: "IfKind.If",
}
func Nodes_IfKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_IfKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_IfKind_getTxt(arg1 LnsInt) string {
    return Nodes_IfKindMap_[arg1];
}
// decl enum -- CaseKind 
const Nodes_CaseKind__Full = 1
const Nodes_CaseKind__Lack = 0
const Nodes_CaseKind__MustFull = 2
var Nodes_CaseKindList_ = NewLnsList( []LnsAny {
  Nodes_CaseKind__Lack,
  Nodes_CaseKind__Full,
  Nodes_CaseKind__MustFull,
})
func Nodes_CaseKind_get__allList() *LnsList{
    return Nodes_CaseKindList_
}
var Nodes_CaseKindMap_ = map[LnsInt]string {
  Nodes_CaseKind__Full: "CaseKind.Full",
  Nodes_CaseKind__Lack: "CaseKind.Lack",
  Nodes_CaseKind__MustFull: "CaseKind.MustFull",
}
func Nodes_CaseKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_CaseKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_CaseKind_getTxt(arg1 LnsInt) string {
    return Nodes_CaseKindMap_[arg1];
}
// decl enum -- CastKind 
const Nodes_CastKind__Force = 1
const Nodes_CastKind__Implicit = 2
const Nodes_CastKind__Normal = 0
var Nodes_CastKindList_ = NewLnsList( []LnsAny {
  Nodes_CastKind__Normal,
  Nodes_CastKind__Force,
  Nodes_CastKind__Implicit,
})
func Nodes_CastKind_get__allList() *LnsList{
    return Nodes_CastKindList_
}
var Nodes_CastKindMap_ = map[LnsInt]string {
  Nodes_CastKind__Force: "CastKind.Force",
  Nodes_CastKind__Implicit: "CastKind.Implicit",
  Nodes_CastKind__Normal: "CastKind.Normal",
}
func Nodes_CastKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_CastKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_CastKind_getTxt(arg1 LnsInt) string {
    return Nodes_CastKindMap_[arg1];
}
// decl enum -- MacroMode 
const Nodes_MacroMode__AnalyzeArg = 2
const Nodes_MacroMode__Expand = 1
const Nodes_MacroMode__None = 0
var Nodes_MacroModeList_ = NewLnsList( []LnsAny {
  Nodes_MacroMode__None,
  Nodes_MacroMode__Expand,
  Nodes_MacroMode__AnalyzeArg,
})
func Nodes_MacroMode_get__allList() *LnsList{
    return Nodes_MacroModeList_
}
var Nodes_MacroModeMap_ = map[LnsInt]string {
  Nodes_MacroMode__AnalyzeArg: "MacroMode.AnalyzeArg",
  Nodes_MacroMode__Expand: "MacroMode.Expand",
  Nodes_MacroMode__None: "MacroMode.None",
}
func Nodes_MacroMode__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_MacroModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_MacroMode_getTxt(arg1 LnsInt) string {
    return Nodes_MacroModeMap_[arg1];
}
// decl enum -- MacroStatKind 
const Nodes_MacroStatKind__Exp = 1
const Nodes_MacroStatKind__Stat = 0
var Nodes_MacroStatKindList_ = NewLnsList( []LnsAny {
  Nodes_MacroStatKind__Stat,
  Nodes_MacroStatKind__Exp,
})
func Nodes_MacroStatKind_get__allList() *LnsList{
    return Nodes_MacroStatKindList_
}
var Nodes_MacroStatKindMap_ = map[LnsInt]string {
  Nodes_MacroStatKind__Exp: "MacroStatKind.Exp",
  Nodes_MacroStatKind__Stat: "MacroStatKind.Stat",
}
func Nodes_MacroStatKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_MacroStatKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_MacroStatKind_getTxt(arg1 LnsInt) string {
    return Nodes_MacroStatKindMap_[arg1];
}
// decl enum -- DeclVarMode 
const Nodes_DeclVarMode__Let = 0
const Nodes_DeclVarMode__Sync = 1
const Nodes_DeclVarMode__Unwrap = 2
var Nodes_DeclVarModeList_ = NewLnsList( []LnsAny {
  Nodes_DeclVarMode__Let,
  Nodes_DeclVarMode__Sync,
  Nodes_DeclVarMode__Unwrap,
})
func Nodes_DeclVarMode_get__allList() *LnsList{
    return Nodes_DeclVarModeList_
}
var Nodes_DeclVarModeMap_ = map[LnsInt]string {
  Nodes_DeclVarMode__Let: "DeclVarMode.Let",
  Nodes_DeclVarMode__Sync: "DeclVarMode.Sync",
  Nodes_DeclVarMode__Unwrap: "DeclVarMode.Unwrap",
}
func Nodes_DeclVarMode__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_DeclVarModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_DeclVarMode_getTxt(arg1 LnsInt) string {
    return Nodes_DeclVarModeMap_[arg1];
}
// decl enum -- FuncKind 
const Nodes_FuncKind__Ctor = 2
const Nodes_FuncKind__Dstr = 3
const Nodes_FuncKind__Func = 0
const Nodes_FuncKind__InitBlock = 4
const Nodes_FuncKind__Mtd = 1
var Nodes_FuncKindList_ = NewLnsList( []LnsAny {
  Nodes_FuncKind__Func,
  Nodes_FuncKind__Mtd,
  Nodes_FuncKind__Ctor,
  Nodes_FuncKind__Dstr,
  Nodes_FuncKind__InitBlock,
})
func Nodes_FuncKind_get__allList() *LnsList{
    return Nodes_FuncKindList_
}
var Nodes_FuncKindMap_ = map[LnsInt]string {
  Nodes_FuncKind__Ctor: "FuncKind.Ctor",
  Nodes_FuncKind__Dstr: "FuncKind.Dstr",
  Nodes_FuncKind__Func: "FuncKind.Func",
  Nodes_FuncKind__InitBlock: "FuncKind.InitBlock",
  Nodes_FuncKind__Mtd: "FuncKind.Mtd",
}
func Nodes_FuncKind__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_FuncKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_FuncKind_getTxt(arg1 LnsInt) string {
    return Nodes_FuncKindMap_[arg1];
}
// decl enum -- nodeKindEnum 
const Nodes_nodeKindEnum__Abbr = 73
const Nodes_nodeKindEnum__Alias = 48
const Nodes_nodeKindEnum__Apply = 15
const Nodes_nodeKindEnum__BlankLine = 2
const Nodes_nodeKindEnum__Block = 7
const Nodes_nodeKindEnum__Boxing = 74
const Nodes_nodeKindEnum__Break = 19
const Nodes_nodeKindEnum__ConvStat = 1
const Nodes_nodeKindEnum__DeclAdvertise = 61
const Nodes_nodeKindEnum__DeclAlge = 65
const Nodes_nodeKindEnum__DeclArg = 59
const Nodes_nodeKindEnum__DeclArgDDD = 60
const Nodes_nodeKindEnum__DeclClass = 63
const Nodes_nodeKindEnum__DeclConstr = 54
const Nodes_nodeKindEnum__DeclDestr = 55
const Nodes_nodeKindEnum__DeclEnum = 64
const Nodes_nodeKindEnum__DeclForm = 50
const Nodes_nodeKindEnum__DeclFunc = 51
const Nodes_nodeKindEnum__DeclMacro = 70
const Nodes_nodeKindEnum__DeclMember = 58
const Nodes_nodeKindEnum__DeclMethod = 52
const Nodes_nodeKindEnum__DeclVar = 49
const Nodes_nodeKindEnum__ExpAccessMRet = 37
const Nodes_nodeKindEnum__ExpCall = 35
const Nodes_nodeKindEnum__ExpCallSuper = 57
const Nodes_nodeKindEnum__ExpCallSuperCtor = 56
const Nodes_nodeKindEnum__ExpCast = 30
const Nodes_nodeKindEnum__ExpList = 10
const Nodes_nodeKindEnum__ExpMRet = 36
const Nodes_nodeKindEnum__ExpMacroArgExp = 42
const Nodes_nodeKindEnum__ExpMacroExp = 40
const Nodes_nodeKindEnum__ExpMacroStat = 41
const Nodes_nodeKindEnum__ExpMacroStatList = 44
const Nodes_nodeKindEnum__ExpMultiTo1 = 38
const Nodes_nodeKindEnum__ExpNew = 21
const Nodes_nodeKindEnum__ExpOmitEnum = 45
const Nodes_nodeKindEnum__ExpOp1 = 33
const Nodes_nodeKindEnum__ExpOp2 = 26
const Nodes_nodeKindEnum__ExpParen = 39
const Nodes_nodeKindEnum__ExpRef = 23
const Nodes_nodeKindEnum__ExpRefItem = 34
const Nodes_nodeKindEnum__ExpSetItem = 25
const Nodes_nodeKindEnum__ExpSetVal = 24
const Nodes_nodeKindEnum__ExpSubDDD = 32
const Nodes_nodeKindEnum__ExpToDDD = 31
const Nodes_nodeKindEnum__ExpUnwrap = 22
const Nodes_nodeKindEnum__For = 14
const Nodes_nodeKindEnum__Foreach = 16
const Nodes_nodeKindEnum__Forsort = 17
const Nodes_nodeKindEnum__GetField = 47
const Nodes_nodeKindEnum__If = 9
const Nodes_nodeKindEnum__IfUnwrap = 28
const Nodes_nodeKindEnum__Import = 4
const Nodes_nodeKindEnum__LiteralArray = 80
const Nodes_nodeKindEnum__LiteralBool = 85
const Nodes_nodeKindEnum__LiteralChar = 77
const Nodes_nodeKindEnum__LiteralInt = 78
const Nodes_nodeKindEnum__LiteralList = 81
const Nodes_nodeKindEnum__LiteralMap = 83
const Nodes_nodeKindEnum__LiteralNil = 76
const Nodes_nodeKindEnum__LiteralReal = 79
const Nodes_nodeKindEnum__LiteralSet = 82
const Nodes_nodeKindEnum__LiteralString = 84
const Nodes_nodeKindEnum__LiteralSymbol = 86
const Nodes_nodeKindEnum__LuneControl = 67
const Nodes_nodeKindEnum__LuneKind = 69
const Nodes_nodeKindEnum__Match = 68
const Nodes_nodeKindEnum__NewAlgeVal = 66
const Nodes_nodeKindEnum__None = 0
const Nodes_nodeKindEnum__ProtoClass = 62
const Nodes_nodeKindEnum__ProtoMethod = 53
const Nodes_nodeKindEnum__Provide = 20
const Nodes_nodeKindEnum__RefField = 46
const Nodes_nodeKindEnum__RefType = 6
const Nodes_nodeKindEnum__Repeat = 13
const Nodes_nodeKindEnum__Return = 18
const Nodes_nodeKindEnum__Root = 5
const Nodes_nodeKindEnum__Scope = 8
const Nodes_nodeKindEnum__StmtExp = 43
const Nodes_nodeKindEnum__Subfile = 3
const Nodes_nodeKindEnum__Switch = 11
const Nodes_nodeKindEnum__TestBlock = 72
const Nodes_nodeKindEnum__TestCase = 71
const Nodes_nodeKindEnum__Unboxing = 75
const Nodes_nodeKindEnum__UnwrapSet = 27
const Nodes_nodeKindEnum__When = 29
const Nodes_nodeKindEnum__While = 12
var Nodes_nodeKindEnumList_ = NewLnsList( []LnsAny {
  Nodes_nodeKindEnum__None,
  Nodes_nodeKindEnum__ConvStat,
  Nodes_nodeKindEnum__BlankLine,
  Nodes_nodeKindEnum__Subfile,
  Nodes_nodeKindEnum__Import,
  Nodes_nodeKindEnum__Root,
  Nodes_nodeKindEnum__RefType,
  Nodes_nodeKindEnum__Block,
  Nodes_nodeKindEnum__Scope,
  Nodes_nodeKindEnum__If,
  Nodes_nodeKindEnum__ExpList,
  Nodes_nodeKindEnum__Switch,
  Nodes_nodeKindEnum__While,
  Nodes_nodeKindEnum__Repeat,
  Nodes_nodeKindEnum__For,
  Nodes_nodeKindEnum__Apply,
  Nodes_nodeKindEnum__Foreach,
  Nodes_nodeKindEnum__Forsort,
  Nodes_nodeKindEnum__Return,
  Nodes_nodeKindEnum__Break,
  Nodes_nodeKindEnum__Provide,
  Nodes_nodeKindEnum__ExpNew,
  Nodes_nodeKindEnum__ExpUnwrap,
  Nodes_nodeKindEnum__ExpRef,
  Nodes_nodeKindEnum__ExpSetVal,
  Nodes_nodeKindEnum__ExpSetItem,
  Nodes_nodeKindEnum__ExpOp2,
  Nodes_nodeKindEnum__UnwrapSet,
  Nodes_nodeKindEnum__IfUnwrap,
  Nodes_nodeKindEnum__When,
  Nodes_nodeKindEnum__ExpCast,
  Nodes_nodeKindEnum__ExpToDDD,
  Nodes_nodeKindEnum__ExpSubDDD,
  Nodes_nodeKindEnum__ExpOp1,
  Nodes_nodeKindEnum__ExpRefItem,
  Nodes_nodeKindEnum__ExpCall,
  Nodes_nodeKindEnum__ExpMRet,
  Nodes_nodeKindEnum__ExpAccessMRet,
  Nodes_nodeKindEnum__ExpMultiTo1,
  Nodes_nodeKindEnum__ExpParen,
  Nodes_nodeKindEnum__ExpMacroExp,
  Nodes_nodeKindEnum__ExpMacroStat,
  Nodes_nodeKindEnum__ExpMacroArgExp,
  Nodes_nodeKindEnum__StmtExp,
  Nodes_nodeKindEnum__ExpMacroStatList,
  Nodes_nodeKindEnum__ExpOmitEnum,
  Nodes_nodeKindEnum__RefField,
  Nodes_nodeKindEnum__GetField,
  Nodes_nodeKindEnum__Alias,
  Nodes_nodeKindEnum__DeclVar,
  Nodes_nodeKindEnum__DeclForm,
  Nodes_nodeKindEnum__DeclFunc,
  Nodes_nodeKindEnum__DeclMethod,
  Nodes_nodeKindEnum__ProtoMethod,
  Nodes_nodeKindEnum__DeclConstr,
  Nodes_nodeKindEnum__DeclDestr,
  Nodes_nodeKindEnum__ExpCallSuperCtor,
  Nodes_nodeKindEnum__ExpCallSuper,
  Nodes_nodeKindEnum__DeclMember,
  Nodes_nodeKindEnum__DeclArg,
  Nodes_nodeKindEnum__DeclArgDDD,
  Nodes_nodeKindEnum__DeclAdvertise,
  Nodes_nodeKindEnum__ProtoClass,
  Nodes_nodeKindEnum__DeclClass,
  Nodes_nodeKindEnum__DeclEnum,
  Nodes_nodeKindEnum__DeclAlge,
  Nodes_nodeKindEnum__NewAlgeVal,
  Nodes_nodeKindEnum__LuneControl,
  Nodes_nodeKindEnum__Match,
  Nodes_nodeKindEnum__LuneKind,
  Nodes_nodeKindEnum__DeclMacro,
  Nodes_nodeKindEnum__TestCase,
  Nodes_nodeKindEnum__TestBlock,
  Nodes_nodeKindEnum__Abbr,
  Nodes_nodeKindEnum__Boxing,
  Nodes_nodeKindEnum__Unboxing,
  Nodes_nodeKindEnum__LiteralNil,
  Nodes_nodeKindEnum__LiteralChar,
  Nodes_nodeKindEnum__LiteralInt,
  Nodes_nodeKindEnum__LiteralReal,
  Nodes_nodeKindEnum__LiteralArray,
  Nodes_nodeKindEnum__LiteralList,
  Nodes_nodeKindEnum__LiteralSet,
  Nodes_nodeKindEnum__LiteralMap,
  Nodes_nodeKindEnum__LiteralString,
  Nodes_nodeKindEnum__LiteralBool,
  Nodes_nodeKindEnum__LiteralSymbol,
})
func Nodes_nodeKindEnum_get__allList() *LnsList{
    return Nodes_nodeKindEnumList_
}
var Nodes_nodeKindEnumMap_ = map[LnsInt]string {
  Nodes_nodeKindEnum__Abbr: "nodeKindEnum.Abbr",
  Nodes_nodeKindEnum__Alias: "nodeKindEnum.Alias",
  Nodes_nodeKindEnum__Apply: "nodeKindEnum.Apply",
  Nodes_nodeKindEnum__BlankLine: "nodeKindEnum.BlankLine",
  Nodes_nodeKindEnum__Block: "nodeKindEnum.Block",
  Nodes_nodeKindEnum__Boxing: "nodeKindEnum.Boxing",
  Nodes_nodeKindEnum__Break: "nodeKindEnum.Break",
  Nodes_nodeKindEnum__ConvStat: "nodeKindEnum.ConvStat",
  Nodes_nodeKindEnum__DeclAdvertise: "nodeKindEnum.DeclAdvertise",
  Nodes_nodeKindEnum__DeclAlge: "nodeKindEnum.DeclAlge",
  Nodes_nodeKindEnum__DeclArg: "nodeKindEnum.DeclArg",
  Nodes_nodeKindEnum__DeclArgDDD: "nodeKindEnum.DeclArgDDD",
  Nodes_nodeKindEnum__DeclClass: "nodeKindEnum.DeclClass",
  Nodes_nodeKindEnum__DeclConstr: "nodeKindEnum.DeclConstr",
  Nodes_nodeKindEnum__DeclDestr: "nodeKindEnum.DeclDestr",
  Nodes_nodeKindEnum__DeclEnum: "nodeKindEnum.DeclEnum",
  Nodes_nodeKindEnum__DeclForm: "nodeKindEnum.DeclForm",
  Nodes_nodeKindEnum__DeclFunc: "nodeKindEnum.DeclFunc",
  Nodes_nodeKindEnum__DeclMacro: "nodeKindEnum.DeclMacro",
  Nodes_nodeKindEnum__DeclMember: "nodeKindEnum.DeclMember",
  Nodes_nodeKindEnum__DeclMethod: "nodeKindEnum.DeclMethod",
  Nodes_nodeKindEnum__DeclVar: "nodeKindEnum.DeclVar",
  Nodes_nodeKindEnum__ExpAccessMRet: "nodeKindEnum.ExpAccessMRet",
  Nodes_nodeKindEnum__ExpCall: "nodeKindEnum.ExpCall",
  Nodes_nodeKindEnum__ExpCallSuper: "nodeKindEnum.ExpCallSuper",
  Nodes_nodeKindEnum__ExpCallSuperCtor: "nodeKindEnum.ExpCallSuperCtor",
  Nodes_nodeKindEnum__ExpCast: "nodeKindEnum.ExpCast",
  Nodes_nodeKindEnum__ExpList: "nodeKindEnum.ExpList",
  Nodes_nodeKindEnum__ExpMRet: "nodeKindEnum.ExpMRet",
  Nodes_nodeKindEnum__ExpMacroArgExp: "nodeKindEnum.ExpMacroArgExp",
  Nodes_nodeKindEnum__ExpMacroExp: "nodeKindEnum.ExpMacroExp",
  Nodes_nodeKindEnum__ExpMacroStat: "nodeKindEnum.ExpMacroStat",
  Nodes_nodeKindEnum__ExpMacroStatList: "nodeKindEnum.ExpMacroStatList",
  Nodes_nodeKindEnum__ExpMultiTo1: "nodeKindEnum.ExpMultiTo1",
  Nodes_nodeKindEnum__ExpNew: "nodeKindEnum.ExpNew",
  Nodes_nodeKindEnum__ExpOmitEnum: "nodeKindEnum.ExpOmitEnum",
  Nodes_nodeKindEnum__ExpOp1: "nodeKindEnum.ExpOp1",
  Nodes_nodeKindEnum__ExpOp2: "nodeKindEnum.ExpOp2",
  Nodes_nodeKindEnum__ExpParen: "nodeKindEnum.ExpParen",
  Nodes_nodeKindEnum__ExpRef: "nodeKindEnum.ExpRef",
  Nodes_nodeKindEnum__ExpRefItem: "nodeKindEnum.ExpRefItem",
  Nodes_nodeKindEnum__ExpSetItem: "nodeKindEnum.ExpSetItem",
  Nodes_nodeKindEnum__ExpSetVal: "nodeKindEnum.ExpSetVal",
  Nodes_nodeKindEnum__ExpSubDDD: "nodeKindEnum.ExpSubDDD",
  Nodes_nodeKindEnum__ExpToDDD: "nodeKindEnum.ExpToDDD",
  Nodes_nodeKindEnum__ExpUnwrap: "nodeKindEnum.ExpUnwrap",
  Nodes_nodeKindEnum__For: "nodeKindEnum.For",
  Nodes_nodeKindEnum__Foreach: "nodeKindEnum.Foreach",
  Nodes_nodeKindEnum__Forsort: "nodeKindEnum.Forsort",
  Nodes_nodeKindEnum__GetField: "nodeKindEnum.GetField",
  Nodes_nodeKindEnum__If: "nodeKindEnum.If",
  Nodes_nodeKindEnum__IfUnwrap: "nodeKindEnum.IfUnwrap",
  Nodes_nodeKindEnum__Import: "nodeKindEnum.Import",
  Nodes_nodeKindEnum__LiteralArray: "nodeKindEnum.LiteralArray",
  Nodes_nodeKindEnum__LiteralBool: "nodeKindEnum.LiteralBool",
  Nodes_nodeKindEnum__LiteralChar: "nodeKindEnum.LiteralChar",
  Nodes_nodeKindEnum__LiteralInt: "nodeKindEnum.LiteralInt",
  Nodes_nodeKindEnum__LiteralList: "nodeKindEnum.LiteralList",
  Nodes_nodeKindEnum__LiteralMap: "nodeKindEnum.LiteralMap",
  Nodes_nodeKindEnum__LiteralNil: "nodeKindEnum.LiteralNil",
  Nodes_nodeKindEnum__LiteralReal: "nodeKindEnum.LiteralReal",
  Nodes_nodeKindEnum__LiteralSet: "nodeKindEnum.LiteralSet",
  Nodes_nodeKindEnum__LiteralString: "nodeKindEnum.LiteralString",
  Nodes_nodeKindEnum__LiteralSymbol: "nodeKindEnum.LiteralSymbol",
  Nodes_nodeKindEnum__LuneControl: "nodeKindEnum.LuneControl",
  Nodes_nodeKindEnum__LuneKind: "nodeKindEnum.LuneKind",
  Nodes_nodeKindEnum__Match: "nodeKindEnum.Match",
  Nodes_nodeKindEnum__NewAlgeVal: "nodeKindEnum.NewAlgeVal",
  Nodes_nodeKindEnum__None: "nodeKindEnum.None",
  Nodes_nodeKindEnum__ProtoClass: "nodeKindEnum.ProtoClass",
  Nodes_nodeKindEnum__ProtoMethod: "nodeKindEnum.ProtoMethod",
  Nodes_nodeKindEnum__Provide: "nodeKindEnum.Provide",
  Nodes_nodeKindEnum__RefField: "nodeKindEnum.RefField",
  Nodes_nodeKindEnum__RefType: "nodeKindEnum.RefType",
  Nodes_nodeKindEnum__Repeat: "nodeKindEnum.Repeat",
  Nodes_nodeKindEnum__Return: "nodeKindEnum.Return",
  Nodes_nodeKindEnum__Root: "nodeKindEnum.Root",
  Nodes_nodeKindEnum__Scope: "nodeKindEnum.Scope",
  Nodes_nodeKindEnum__StmtExp: "nodeKindEnum.StmtExp",
  Nodes_nodeKindEnum__Subfile: "nodeKindEnum.Subfile",
  Nodes_nodeKindEnum__Switch: "nodeKindEnum.Switch",
  Nodes_nodeKindEnum__TestBlock: "nodeKindEnum.TestBlock",
  Nodes_nodeKindEnum__TestCase: "nodeKindEnum.TestCase",
  Nodes_nodeKindEnum__Unboxing: "nodeKindEnum.Unboxing",
  Nodes_nodeKindEnum__UnwrapSet: "nodeKindEnum.UnwrapSet",
  Nodes_nodeKindEnum__When: "nodeKindEnum.When",
  Nodes_nodeKindEnum__While: "nodeKindEnum.While",
}
func Nodes_nodeKindEnum__from(arg1 LnsInt) LnsAny{
    if _, ok := Nodes_nodeKindEnumMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_nodeKindEnum_getTxt(arg1 LnsInt) string {
    return Nodes_nodeKindEnumMap_[arg1];
}
var Nodes_nodeKind2NameMap *LnsMap
var Nodes_nodeKindSeed LnsInt
type Nodes_Literal__ARRAY struct{
Val1 *LnsList
}
func (self *Nodes_Literal__ARRAY) GetTxt() string {
return "Literal.ARRAY"
}
type Nodes_Literal__Bool struct{
Val1 bool
}
func (self *Nodes_Literal__Bool) GetTxt() string {
return "Literal.Bool"
}
type Nodes_Literal__Field struct{
Val1 *LnsList
}
func (self *Nodes_Literal__Field) GetTxt() string {
return "Literal.Field"
}
type Nodes_Literal__Int struct{
Val1 LnsInt
}
func (self *Nodes_Literal__Int) GetTxt() string {
return "Literal.Int"
}
type Nodes_Literal__LIST struct{
Val1 *LnsList
}
func (self *Nodes_Literal__LIST) GetTxt() string {
return "Literal.LIST"
}
type Nodes_Literal__MAP struct{
Val1 *LnsMap
}
func (self *Nodes_Literal__MAP) GetTxt() string {
return "Literal.MAP"
}
type Nodes_Literal__Nil struct{
}
var Nodes_Literal__Nil_Obj = &Nodes_Literal__Nil{}
func (self *Nodes_Literal__Nil) GetTxt() string {
return "Literal.Nil"
}
type Nodes_Literal__Real struct{
Val1 LnsReal
}
func (self *Nodes_Literal__Real) GetTxt() string {
return "Literal.Real"
}
type Nodes_Literal__SET struct{
Val1 *LnsList
}
func (self *Nodes_Literal__SET) GetTxt() string {
return "Literal.SET"
}
type Nodes_Literal__Str struct{
Val1 string
}
func (self *Nodes_Literal__Str) GetTxt() string {
return "Literal.Str"
}
type Nodes_Literal__Symbol struct{
Val1 string
}
func (self *Nodes_Literal__Symbol) GetTxt() string {
return "Literal.Symbol"
}
type Nodes_IndexVal__NodeIdx struct{
Val1 *Nodes_Node
}
func (self *Nodes_IndexVal__NodeIdx) GetTxt() string {
return "IndexVal.NodeIdx"
}
type Nodes_IndexVal__SymIdx struct{
Val1 string
}
func (self *Nodes_IndexVal__SymIdx) GetTxt() string {
return "IndexVal.SymIdx"
}
type Nodes_NodeVisitor func (arg1 *Nodes_Node,arg2 *Nodes_Node,arg3 string,arg4 LnsInt) LnsInt
type Nodes_macroStatmentProc func (arg1 *LnsMap) *LnsMap
// for 2429
func Nodes_convExp42888(arg1 string, arg2 []LnsAny) (string, []LnsAny) {
    return arg1, Lns_2DDD( arg2[0:])
}
// for 503
func Nodes_convExp1367(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2373
func Nodes_convExp42626(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2530
func Nodes_convExp43351(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2590
func Nodes_convExp43627(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2640
func Nodes_convExp43758(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}

// 136: decl @lune.@base.@Nodes.getLiteralObj
func Nodes_getLiteralObj(obj LnsAny) LnsAny {
    switch _exp382 := obj.(type) {
    case *Nodes_Literal__Nil:
        return nil
    case *Nodes_Literal__Int:
    val := _exp382.Val1
        return val
    case *Nodes_Literal__Real:
    val := _exp382.Val1
        return val
    case *Nodes_Literal__Str:
    val := _exp382.Val1
        return val
    case *Nodes_Literal__Bool:
    val := _exp382.Val1
        return val
    case *Nodes_Literal__Symbol:
    val := _exp382.Val1
        return val
    case *Nodes_Literal__Field:
    val := _exp382.Val1
        return val
    case *Nodes_Literal__LIST:
    val := _exp382.Val1
        return val
    case *Nodes_Literal__ARRAY:
    val := _exp382.Val1
        return val
    case *Nodes_Literal__SET:
    val := _exp382.Val1
        return val
    case *Nodes_Literal__MAP:
    val := _exp382.Val1
        return val
    }
// insert a dummy
    return nil
}

// 334: decl @lune.@base.@Nodes.regKind
func Nodes_regKind_1338_(name string) LnsInt {
    var kind LnsInt
    kind = Nodes_nodeKindSeed
    Nodes_nodeKindSeed = Nodes_nodeKindSeed + 1
    
    Nodes_nodeKind2NameMap.Set(kind,name)
    return kind
}

// 341: decl @lune.@base.@Nodes.getNodeKindName
func Nodes_getNodeKindName(kind LnsInt) string {
    return Lns_unwrap( Nodes_nodeKind2NameMap.Items[kind]).(string)
}


// 962: decl @lune.@base.@Nodes.getBreakKindForStmtList
func Nodes_getBreakKindForStmtList_2517_(checkMode LnsInt,stmtList *LnsList) LnsInt {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        var kind LnsInt
        kind = Nodes_BreakKind__None
        for _, _stmt := range( stmtList.Items ) {
            stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
            if stmt.FP.Get_kind() != Nodes_NodeKind_get_BlankLine(){
                var work LnsInt
                work = stmt.FP.GetBreakKind(checkMode)
                if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                    if work == Nodes_BreakKind__Return{
                        return Nodes_BreakKind__Return
                    }
                    if work == Nodes_BreakKind__NeverRet{
                        return Nodes_BreakKind__NeverRet
                    }
                } else { 
                    if _switch6462 := work; _switch6462 == Nodes_BreakKind__None {
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                            Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                            if false{
                                return Nodes_BreakKind__None
                            }
                        }
                    } else {
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                            Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                            kind = work
                            
                        }
                    }
                }
                
            }
        }
        return kind
    }
    if stmtList.Len() > 0{
        {
            var _from6519 LnsInt = stmtList.Len()
            var _to6519 LnsInt = 1
            _work6519 := _from6519
            _delta6519 := -1
            for {
                if _delta6519 > 0 {
                   if _work6519 > _to6519 { break }
                } else {
                   if _work6519 < _to6519 { break }
                }
                index := _work6519
                var stmt *Nodes_Node
                stmt = stmtList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                if stmt.FP.Get_kind() != Nodes_NodeKind_get_BlankLine(){
                    return stmt.FP.GetBreakKind(checkMode)
                }
                _work6519 += _delta6519
            }
        }
    }
    return Nodes_BreakKind__None
}

// 2098: decl @lune.@base.@Nodes.Node.getSymbolInfo.processExpNode
func Node_getSymbolInfo__processExpNode_9477_(node *Nodes_Node) *LnsList {
    if _switch41475 := (node.FP.Get_kind()); _switch41475 == Nodes_NodeKind_get_ExpRef() {
        return NewLnsList([]LnsAny{Ast_SymbolInfo2Stem((Lns_unwrap( (Nodes_ExpRefNodeDownCastF(node.FP))).(*Nodes_ExpRefNode)).FP.Get_symbolInfo())})
    } else if _switch41475 == Nodes_NodeKind_get_RefField() {
        {
            _refFieldNode := Nodes_RefFieldNodeDownCastF(node.FP)
            if _refFieldNode != nil {
                refFieldNode := _refFieldNode.(*Nodes_RefFieldNode)
                if refFieldNode.FP.Get_nilAccess(){
                    return NewLnsList([]LnsAny{})
                }
                {
                    __exp := refFieldNode.FP.Get_symbolInfo()
                    if __exp != nil {
                        _exp := __exp.(*Ast_SymbolInfo)
                        return NewLnsList([]LnsAny{Ast_SymbolInfo2Stem(_exp)})
                    }
                }
            }
        }
    } else if _switch41475 == Nodes_NodeKind_get_ExpList() {
        {
            _expListNode := Nodes_ExpListNodeDownCastF(node.FP)
            if _expListNode != nil {
                expListNode := _expListNode.(*Nodes_ExpListNode)
                var list *LnsList
                list = NewLnsList([]LnsAny{})
                for _index, _expNode := range( expListNode.FP.Get_expList().Items ) {
                    index := _index + 1
                    expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                    if index == expListNode.FP.Get_expList().Len(){
                        for _, _symbolInfo := range( Node_getSymbolInfo__processExpNode_9477_(expNode).Items ) {
                            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            list.Insert(Ast_SymbolInfo2Stem(symbolInfo))
                        }
                    } else { 
                        for _, _symbolInfo := range( Node_getSymbolInfo__processExpNode_9477_(expNode).Items ) {
                            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            list.Insert(Ast_SymbolInfo2Stem(symbolInfo))
                            break
                        }
                    }
                }
                return list
            }
        }
    } else if _switch41475 == Nodes_NodeKind_get_RefType() {
        {
            _refTypeNode := Nodes_RefTypeNodeDownCastF(node.FP)
            if _refTypeNode != nil {
                refTypeNode := _refTypeNode.(*Nodes_RefTypeNode)
                return refTypeNode.FP.Get_name().FP.GetSymbolInfo()
            }
        }
    }
    return NewLnsList([]LnsAny{})
}

// 2474: decl @lune.@base.@Nodes.enumLiteral2Literal
func Nodes_enumLiteral2Literal_9640_(obj LnsAny)(LnsAny, LnsAny) {
    switch _exp43148 := obj.(type) {
    case *Ast_EnumLiteral__Int:
    val := _exp43148.Val1
        return &Nodes_Literal__Int{val}, nil
    case *Ast_EnumLiteral__Real:
    val := _exp43148.Val1
        return &Nodes_Literal__Real{val}, nil
    case *Ast_EnumLiteral__Str:
    val := _exp43148.Val1
        return &Nodes_Literal__Str{val}, nil
    }
// insert a dummy
    return nil,nil
}

// 2785: decl @lune.@base.@Nodes.hasMultiValNode
func Nodes_hasMultiValNode(node *Nodes_Node) bool {
    return Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( node.FP.Get_expTypeList().Len() > 1) ||
        Lns_GetEnv().SetStackVal( node.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__DDD) ).(bool)
}

// 2806: decl @lune.@base.@Nodes.getUnwraped
func Nodes_getUnwraped(node *Nodes_Node) *Nodes_Node {
    {
        _work := Nodes_ExpMRetNodeDownCastF(node.FP)
        if _work != nil {
            work := _work.(*Nodes_ExpMRetNode)
            return Nodes_getUnwraped(work.FP.Get_mRet())
        }
    }
    {
        _work := Nodes_ExpParenNodeDownCastF(node.FP)
        if _work != nil {
            work := _work.(*Nodes_ExpParenNode)
            return Nodes_getUnwraped(work.FP.Get_exp())
        }
    }
    return node
}

// 2816: decl @lune.@base.@Nodes.getCastUnwraped
func Nodes_getCastUnwraped(node *Nodes_Node) *Nodes_Node {
    {
        _work := Nodes_ExpCastNodeDownCastF(node.FP)
        if _work != nil {
            work := _work.(*Nodes_ExpCastNode)
            return Nodes_getUnwraped(work.FP.Get_exp())
        }
    }
    return node
}

// declaration Class -- SimpleModuleInfoManager
type Nodes_SimpleModuleInfoManagerMtd interface {
    GetModuleInfo(arg1 *Ast_TypeInfo) LnsAny
    Pop()
    Push(arg1 Ast_ModuleInfoManager)
}
type Nodes_SimpleModuleInfoManager struct {
    ModuleInfoManager Ast_ModuleInfoManager
    moduleInfoManagerHist *LnsList
    FP Nodes_SimpleModuleInfoManagerMtd
}
func Nodes_SimpleModuleInfoManager2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_SimpleModuleInfoManager).FP
}
type Nodes_SimpleModuleInfoManagerDownCast interface {
    ToNodes_SimpleModuleInfoManager() *Nodes_SimpleModuleInfoManager
}
func Nodes_SimpleModuleInfoManagerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_SimpleModuleInfoManagerDownCast)
    if ok { return work.ToNodes_SimpleModuleInfoManager() }
    return nil
}
func (obj *Nodes_SimpleModuleInfoManager) ToNodes_SimpleModuleInfoManager() *Nodes_SimpleModuleInfoManager {
    return obj
}
func NewNodes_SimpleModuleInfoManager(arg1 LnsAny) *Nodes_SimpleModuleInfoManager {
    obj := &Nodes_SimpleModuleInfoManager{}
    obj.FP = obj
    obj.InitNodes_SimpleModuleInfoManager(arg1)
    return obj
}
func (self *Nodes_SimpleModuleInfoManager) GetModuleInfo(arg1 *Ast_TypeInfo) LnsAny {
    return self.ModuleInfoManager. GetModuleInfo( arg1)
}
// 36: DeclConstr
func (self *Nodes_SimpleModuleInfoManager) InitNodes_SimpleModuleInfoManager(moduleInfoManager LnsAny) {
    if moduleInfoManager != nil{
        moduleInfoManager_2343 := moduleInfoManager.(Ast_ModuleInfoManager)
        self.ModuleInfoManager = moduleInfoManager_2343
        
    } else {
        self.ModuleInfoManager = Ast_DummyModuleInfoManager_get_instance().FP
        
    }
    self.moduleInfoManagerHist = NewLnsList([]LnsAny{})
    
}

// 46: decl @lune.@base.@Nodes.SimpleModuleInfoManager.push
func (self *Nodes_SimpleModuleInfoManager) Push(moduleInfoManager Ast_ModuleInfoManager) {
    self.moduleInfoManagerHist.Insert(self.ModuleInfoManager)
    self.ModuleInfoManager = moduleInfoManager
    
}

// 51: decl @lune.@base.@Nodes.SimpleModuleInfoManager.pop
func (self *Nodes_SimpleModuleInfoManager) Pop() {
    self.ModuleInfoManager = self.moduleInfoManagerHist.GetAt(self.moduleInfoManagerHist.Len()).(Ast_ModuleInfoManager)
    
    self.moduleInfoManagerHist.Remove(nil)
}


// declaration Class -- Filter
type Nodes_FilterMtd interface {
    DefaultProcess(arg1 *Nodes_Node, arg2 LnsAny)
    Get_moduleInfoManager() Ast_ModuleInfoManager
    Get_optStack() *LnsList
    Get_typeNameCtrl() *Ast_TypeNameCtrl
    popOpt(arg1 LnsAny)
    ProcessAbbr(arg1 *Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(arg1 *Nodes_AliasNode, arg2 LnsAny)
    ProcessApply(arg1 *Nodes_ApplyNode, arg2 LnsAny)
    ProcessBlankLine(arg1 *Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(arg1 *Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(arg1 *Nodes_BreakNode, arg2 LnsAny)
    ProcessConvStat(arg1 *Nodes_ConvStatNode, arg2 LnsAny)
    ProcessDeclAdvertise(arg1 *Nodes_DeclAdvertiseNode, arg2 LnsAny)
    ProcessDeclAlge(arg1 *Nodes_DeclAlgeNode, arg2 LnsAny)
    ProcessDeclArg(arg1 *Nodes_DeclArgNode, arg2 LnsAny)
    ProcessDeclArgDDD(arg1 *Nodes_DeclArgDDDNode, arg2 LnsAny)
    ProcessDeclClass(arg1 *Nodes_DeclClassNode, arg2 LnsAny)
    ProcessDeclConstr(arg1 *Nodes_DeclConstrNode, arg2 LnsAny)
    ProcessDeclDestr(arg1 *Nodes_DeclDestrNode, arg2 LnsAny)
    ProcessDeclEnum(arg1 *Nodes_DeclEnumNode, arg2 LnsAny)
    ProcessDeclForm(arg1 *Nodes_DeclFormNode, arg2 LnsAny)
    ProcessDeclFunc(arg1 *Nodes_DeclFuncNode, arg2 LnsAny)
    ProcessDeclMacro(arg1 *Nodes_DeclMacroNode, arg2 LnsAny)
    ProcessDeclMember(arg1 *Nodes_DeclMemberNode, arg2 LnsAny)
    ProcessDeclMethod(arg1 *Nodes_DeclMethodNode, arg2 LnsAny)
    ProcessDeclVar(arg1 *Nodes_DeclVarNode, arg2 LnsAny)
    ProcessExpAccessMRet(arg1 *Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(arg1 *Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(arg1 *Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(arg1 *Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(arg1 *Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpList(arg1 *Nodes_ExpListNode, arg2 LnsAny)
    ProcessExpMRet(arg1 *Nodes_ExpMRetNode, arg2 LnsAny)
    ProcessExpMacroArgExp(arg1 *Nodes_ExpMacroArgExpNode, arg2 LnsAny)
    ProcessExpMacroExp(arg1 *Nodes_ExpMacroExpNode, arg2 LnsAny)
    ProcessExpMacroStat(arg1 *Nodes_ExpMacroStatNode, arg2 LnsAny)
    ProcessExpMacroStatList(arg1 *Nodes_ExpMacroStatListNode, arg2 LnsAny)
    ProcessExpMultiTo1(arg1 *Nodes_ExpMultiTo1Node, arg2 LnsAny)
    ProcessExpNew(arg1 *Nodes_ExpNewNode, arg2 LnsAny)
    ProcessExpOmitEnum(arg1 *Nodes_ExpOmitEnumNode, arg2 LnsAny)
    ProcessExpOp1(arg1 *Nodes_ExpOp1Node, arg2 LnsAny)
    ProcessExpOp2(arg1 *Nodes_ExpOp2Node, arg2 LnsAny)
    ProcessExpParen(arg1 *Nodes_ExpParenNode, arg2 LnsAny)
    ProcessExpRef(arg1 *Nodes_ExpRefNode, arg2 LnsAny)
    ProcessExpRefItem(arg1 *Nodes_ExpRefItemNode, arg2 LnsAny)
    ProcessExpSetItem(arg1 *Nodes_ExpSetItemNode, arg2 LnsAny)
    ProcessExpSetVal(arg1 *Nodes_ExpSetValNode, arg2 LnsAny)
    ProcessExpSubDDD(arg1 *Nodes_ExpSubDDDNode, arg2 LnsAny)
    ProcessExpToDDD(arg1 *Nodes_ExpToDDDNode, arg2 LnsAny)
    ProcessExpUnwrap(arg1 *Nodes_ExpUnwrapNode, arg2 LnsAny)
    ProcessFor(arg1 *Nodes_ForNode, arg2 LnsAny)
    ProcessForeach(arg1 *Nodes_ForeachNode, arg2 LnsAny)
    ProcessForsort(arg1 *Nodes_ForsortNode, arg2 LnsAny)
    ProcessGetField(arg1 *Nodes_GetFieldNode, arg2 LnsAny)
    ProcessIf(arg1 *Nodes_IfNode, arg2 LnsAny)
    ProcessIfUnwrap(arg1 *Nodes_IfUnwrapNode, arg2 LnsAny)
    ProcessImport(arg1 *Nodes_ImportNode, arg2 LnsAny)
    ProcessLiteralArray(arg1 *Nodes_LiteralArrayNode, arg2 LnsAny)
    ProcessLiteralBool(arg1 *Nodes_LiteralBoolNode, arg2 LnsAny)
    ProcessLiteralChar(arg1 *Nodes_LiteralCharNode, arg2 LnsAny)
    ProcessLiteralInt(arg1 *Nodes_LiteralIntNode, arg2 LnsAny)
    ProcessLiteralList(arg1 *Nodes_LiteralListNode, arg2 LnsAny)
    ProcessLiteralMap(arg1 *Nodes_LiteralMapNode, arg2 LnsAny)
    ProcessLiteralNil(arg1 *Nodes_LiteralNilNode, arg2 LnsAny)
    ProcessLiteralReal(arg1 *Nodes_LiteralRealNode, arg2 LnsAny)
    ProcessLiteralSet(arg1 *Nodes_LiteralSetNode, arg2 LnsAny)
    ProcessLiteralString(arg1 *Nodes_LiteralStringNode, arg2 LnsAny)
    ProcessLiteralSymbol(arg1 *Nodes_LiteralSymbolNode, arg2 LnsAny)
    ProcessLuneControl(arg1 *Nodes_LuneControlNode, arg2 LnsAny)
    ProcessLuneKind(arg1 *Nodes_LuneKindNode, arg2 LnsAny)
    ProcessMatch(arg1 *Nodes_MatchNode, arg2 LnsAny)
    ProcessNewAlgeVal(arg1 *Nodes_NewAlgeValNode, arg2 LnsAny)
    ProcessNone(arg1 *Nodes_NoneNode, arg2 LnsAny)
    ProcessProtoClass(arg1 *Nodes_ProtoClassNode, arg2 LnsAny)
    ProcessProtoMethod(arg1 *Nodes_ProtoMethodNode, arg2 LnsAny)
    ProcessProvide(arg1 *Nodes_ProvideNode, arg2 LnsAny)
    ProcessRefField(arg1 *Nodes_RefFieldNode, arg2 LnsAny)
    ProcessRefType(arg1 *Nodes_RefTypeNode, arg2 LnsAny)
    ProcessRepeat(arg1 *Nodes_RepeatNode, arg2 LnsAny)
    ProcessReturn(arg1 *Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(arg1 *Nodes_RootNode, arg2 LnsAny)
    ProcessScope(arg1 *Nodes_ScopeNode, arg2 LnsAny)
    ProcessStmtExp(arg1 *Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(arg1 *Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(arg1 *Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(arg1 *Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(arg1 *Nodes_TestCaseNode, arg2 LnsAny)
    ProcessUnboxing(arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(arg1 *Nodes_WhileNode, arg2 LnsAny)
    pushOpt(arg1 LnsAny)
}
type Nodes_Filter struct {
    moduleInfoManager *Nodes_SimpleModuleInfoManager
    typeNameCtrl *Ast_TypeNameCtrl
    errorOnDefault bool
    optStack *LnsList
    FP Nodes_FilterMtd
}
func Nodes_Filter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_Filter).FP
}
type Nodes_FilterDownCast interface {
    ToNodes_Filter() *Nodes_Filter
}
func Nodes_FilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_FilterDownCast)
    if ok { return work.ToNodes_Filter() }
    return nil
}
func (obj *Nodes_Filter) ToNodes_Filter() *Nodes_Filter {
    return obj
}
func NewNodes_Filter(arg1 bool, arg2 LnsAny, arg3 LnsAny) *Nodes_Filter {
    obj := &Nodes_Filter{}
    obj.FP = obj
    obj.InitNodes_Filter(arg1, arg2, arg3)
    return obj
}
func (self *Nodes_Filter) Get_typeNameCtrl() *Ast_TypeNameCtrl{ return self.typeNameCtrl }
func (self *Nodes_Filter) Get_optStack() *LnsList{ return self.optStack }
// 65: DeclConstr
func (self *Nodes_Filter) InitNodes_Filter(errorOnDefault bool,moduleTypeInfo LnsAny,moduleInfoManager LnsAny) {
    self.errorOnDefault = errorOnDefault
    
    self.moduleInfoManager = NewNodes_SimpleModuleInfoManager(moduleInfoManager)
    
    var process func() *Ast_TypeNameCtrl
    process = func() *Ast_TypeNameCtrl {
        if moduleTypeInfo != nil{
            moduleTypeInfo_2373 := moduleTypeInfo.(*Ast_TypeInfo)
            return NewAst_TypeNameCtrl(moduleTypeInfo_2373)
        }
        return Ast_defaultTypeNameCtrl
    }
    self.typeNameCtrl = process()
    
    self.optStack = NewLnsList([]LnsAny{})
    
}

// 81: decl @lune.@base.@Nodes.Filter.pushOpt
func (self *Nodes_Filter) pushOpt(_opt LnsAny) {
    opt := _opt
    self.optStack.Insert(opt)
}

// 84: decl @lune.@base.@Nodes.Filter.popOpt
func (self *Nodes_Filter) popOpt(_opt LnsAny) {
    self.optStack.Remove(nil)
}

// 88: decl @lune.@base.@Nodes.Filter.get_moduleInfoManager
func (self *Nodes_Filter) Get_moduleInfoManager() Ast_ModuleInfoManager {
    return self.moduleInfoManager.FP
}

// 345: decl @lune.@base.@Nodes.Filter.defaultProcess
func (self *Nodes_Filter) DefaultProcess(node *Nodes_Node,_opt LnsAny) {
    if self.errorOnDefault{
        Util_err(Lns_getVM().String_format("not implement yet -- %s", []LnsAny{Nodes_getNodeKindName(node.FP.Get_kind())}))
    }
}

// 1: decl @lune.@base.@Nodes.Filter.processNone
func (self *Nodes_Filter) ProcessNone(node *Nodes_NoneNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processConvStat
func (self *Nodes_Filter) ProcessConvStat(node *Nodes_ConvStatNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processBlankLine
func (self *Nodes_Filter) ProcessBlankLine(node *Nodes_BlankLineNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processSubfile
func (self *Nodes_Filter) ProcessSubfile(node *Nodes_SubfileNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processImport
func (self *Nodes_Filter) ProcessImport(node *Nodes_ImportNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processRoot
func (self *Nodes_Filter) ProcessRoot(node *Nodes_RootNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processRefType
func (self *Nodes_Filter) ProcessRefType(node *Nodes_RefTypeNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}


// 1: decl @lune.@base.@Nodes.Filter.processScope
func (self *Nodes_Filter) ProcessScope(node *Nodes_ScopeNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processIf
func (self *Nodes_Filter) ProcessIf(node *Nodes_IfNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpList
func (self *Nodes_Filter) ProcessExpList(node *Nodes_ExpListNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processSwitch
func (self *Nodes_Filter) ProcessSwitch(node *Nodes_SwitchNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processWhile
func (self *Nodes_Filter) ProcessWhile(node *Nodes_WhileNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processRepeat
func (self *Nodes_Filter) ProcessRepeat(node *Nodes_RepeatNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processFor
func (self *Nodes_Filter) ProcessFor(node *Nodes_ForNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processApply
func (self *Nodes_Filter) ProcessApply(node *Nodes_ApplyNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processForeach
func (self *Nodes_Filter) ProcessForeach(node *Nodes_ForeachNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processForsort
func (self *Nodes_Filter) ProcessForsort(node *Nodes_ForsortNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processReturn
func (self *Nodes_Filter) ProcessReturn(node *Nodes_ReturnNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processBreak
func (self *Nodes_Filter) ProcessBreak(node *Nodes_BreakNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processProvide
func (self *Nodes_Filter) ProcessProvide(node *Nodes_ProvideNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpNew
func (self *Nodes_Filter) ProcessExpNew(node *Nodes_ExpNewNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpUnwrap
func (self *Nodes_Filter) ProcessExpUnwrap(node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpRef
func (self *Nodes_Filter) ProcessExpRef(node *Nodes_ExpRefNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpSetVal
func (self *Nodes_Filter) ProcessExpSetVal(node *Nodes_ExpSetValNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpSetItem
func (self *Nodes_Filter) ProcessExpSetItem(node *Nodes_ExpSetItemNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpOp2
func (self *Nodes_Filter) ProcessExpOp2(node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processUnwrapSet
func (self *Nodes_Filter) ProcessUnwrapSet(node *Nodes_UnwrapSetNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processIfUnwrap
func (self *Nodes_Filter) ProcessIfUnwrap(node *Nodes_IfUnwrapNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processWhen
func (self *Nodes_Filter) ProcessWhen(node *Nodes_WhenNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpCast
func (self *Nodes_Filter) ProcessExpCast(node *Nodes_ExpCastNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpToDDD
func (self *Nodes_Filter) ProcessExpToDDD(node *Nodes_ExpToDDDNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpSubDDD
func (self *Nodes_Filter) ProcessExpSubDDD(node *Nodes_ExpSubDDDNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpOp1
func (self *Nodes_Filter) ProcessExpOp1(node *Nodes_ExpOp1Node,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpRefItem
func (self *Nodes_Filter) ProcessExpRefItem(node *Nodes_ExpRefItemNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpCall
func (self *Nodes_Filter) ProcessExpCall(node *Nodes_ExpCallNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpMRet
func (self *Nodes_Filter) ProcessExpMRet(node *Nodes_ExpMRetNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpAccessMRet
func (self *Nodes_Filter) ProcessExpAccessMRet(node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpMultiTo1
func (self *Nodes_Filter) ProcessExpMultiTo1(node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpParen
func (self *Nodes_Filter) ProcessExpParen(node *Nodes_ExpParenNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpMacroExp
func (self *Nodes_Filter) ProcessExpMacroExp(node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpMacroStat
func (self *Nodes_Filter) ProcessExpMacroStat(node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpMacroArgExp
func (self *Nodes_Filter) ProcessExpMacroArgExp(node *Nodes_ExpMacroArgExpNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processStmtExp
func (self *Nodes_Filter) ProcessStmtExp(node *Nodes_StmtExpNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpMacroStatList
func (self *Nodes_Filter) ProcessExpMacroStatList(node *Nodes_ExpMacroStatListNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpOmitEnum
func (self *Nodes_Filter) ProcessExpOmitEnum(node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processRefField
func (self *Nodes_Filter) ProcessRefField(node *Nodes_RefFieldNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processGetField
func (self *Nodes_Filter) ProcessGetField(node *Nodes_GetFieldNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processAlias
func (self *Nodes_Filter) ProcessAlias(node *Nodes_AliasNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclVar
func (self *Nodes_Filter) ProcessDeclVar(node *Nodes_DeclVarNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclForm
func (self *Nodes_Filter) ProcessDeclForm(node *Nodes_DeclFormNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclFunc
func (self *Nodes_Filter) ProcessDeclFunc(node *Nodes_DeclFuncNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclMethod
func (self *Nodes_Filter) ProcessDeclMethod(node *Nodes_DeclMethodNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processProtoMethod
func (self *Nodes_Filter) ProcessProtoMethod(node *Nodes_ProtoMethodNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclConstr
func (self *Nodes_Filter) ProcessDeclConstr(node *Nodes_DeclConstrNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclDestr
func (self *Nodes_Filter) ProcessDeclDestr(node *Nodes_DeclDestrNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpCallSuperCtor
func (self *Nodes_Filter) ProcessExpCallSuperCtor(node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processExpCallSuper
func (self *Nodes_Filter) ProcessExpCallSuper(node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclMember
func (self *Nodes_Filter) ProcessDeclMember(node *Nodes_DeclMemberNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclArg
func (self *Nodes_Filter) ProcessDeclArg(node *Nodes_DeclArgNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclArgDDD
func (self *Nodes_Filter) ProcessDeclArgDDD(node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclAdvertise
func (self *Nodes_Filter) ProcessDeclAdvertise(node *Nodes_DeclAdvertiseNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processProtoClass
func (self *Nodes_Filter) ProcessProtoClass(node *Nodes_ProtoClassNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclClass
func (self *Nodes_Filter) ProcessDeclClass(node *Nodes_DeclClassNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclEnum
func (self *Nodes_Filter) ProcessDeclEnum(node *Nodes_DeclEnumNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclAlge
func (self *Nodes_Filter) ProcessDeclAlge(node *Nodes_DeclAlgeNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processNewAlgeVal
func (self *Nodes_Filter) ProcessNewAlgeVal(node *Nodes_NewAlgeValNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLuneControl
func (self *Nodes_Filter) ProcessLuneControl(node *Nodes_LuneControlNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processMatch
func (self *Nodes_Filter) ProcessMatch(node *Nodes_MatchNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLuneKind
func (self *Nodes_Filter) ProcessLuneKind(node *Nodes_LuneKindNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processDeclMacro
func (self *Nodes_Filter) ProcessDeclMacro(node *Nodes_DeclMacroNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processTestCase
func (self *Nodes_Filter) ProcessTestCase(node *Nodes_TestCaseNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processTestBlock
func (self *Nodes_Filter) ProcessTestBlock(node *Nodes_TestBlockNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processAbbr
func (self *Nodes_Filter) ProcessAbbr(node *Nodes_AbbrNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processBoxing
func (self *Nodes_Filter) ProcessBoxing(node *Nodes_BoxingNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processUnboxing
func (self *Nodes_Filter) ProcessUnboxing(node *Nodes_UnboxingNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralNil
func (self *Nodes_Filter) ProcessLiteralNil(node *Nodes_LiteralNilNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralChar
func (self *Nodes_Filter) ProcessLiteralChar(node *Nodes_LiteralCharNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralInt
func (self *Nodes_Filter) ProcessLiteralInt(node *Nodes_LiteralIntNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralReal
func (self *Nodes_Filter) ProcessLiteralReal(node *Nodes_LiteralRealNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralArray
func (self *Nodes_Filter) ProcessLiteralArray(node *Nodes_LiteralArrayNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralList
func (self *Nodes_Filter) ProcessLiteralList(node *Nodes_LiteralListNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralSet
func (self *Nodes_Filter) ProcessLiteralSet(node *Nodes_LiteralSetNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralMap
func (self *Nodes_Filter) ProcessLiteralMap(node *Nodes_LiteralMapNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralString
func (self *Nodes_Filter) ProcessLiteralString(node *Nodes_LiteralStringNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralBool
func (self *Nodes_Filter) ProcessLiteralBool(node *Nodes_LiteralBoolNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 1: decl @lune.@base.@Nodes.Filter.processLiteralSymbol
func (self *Nodes_Filter) ProcessLiteralSymbol(node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(opt)
    self.FP.DefaultProcess(&node.Nodes_Node, opt)
    self.FP.popOpt(opt)
}

// 2768: decl @lune.@base.@Nodes.Filter.processBlockSub
func (self *Nodes_Filter) ProcessBlockSub(node *Nodes_BlockNode,_opt LnsAny) {
}

// 2771: decl @lune.@base.@Nodes.Filter.processBlock
func (self *Nodes_Filter) ProcessBlock(node *Nodes_BlockNode,_opt LnsAny) {
    opt := _opt
    self.moduleInfoManager.FP.Push(node.FP.Get_scope().FP)
    self.FP.ProcessBlockSub(node, opt)
    self.moduleInfoManager.FP.Pop()
}


// declaration Class -- Node
type Nodes_NodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_Node struct {
    id LnsInt
    kind LnsInt
    pos *Types_Position
    expTypeList *LnsList
    commentList LnsAny
    tailComment LnsAny
    IsLValue bool
    macroArgFlag bool
    FP Nodes_NodeMtd
}
func Nodes_Node2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_Node).FP
}
type Nodes_NodeDownCast interface {
    ToNodes_Node() *Nodes_Node
}
func Nodes_NodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_NodeDownCast)
    if ok { return work.ToNodes_Node() }
    return nil
}
func (obj *Nodes_Node) ToNodes_Node() *Nodes_Node {
    return obj
}
func (self *Nodes_Node) Get_id() LnsInt{ return self.id }
func (self *Nodes_Node) Get_kind() LnsInt{ return self.kind }
func (self *Nodes_Node) Get_pos() *Types_Position{ return self.pos }
func (self *Nodes_Node) Get_expTypeList() *LnsList{ return self.expTypeList }
func (self *Nodes_Node) Get_commentList() LnsAny{ return self.commentList }
func (self *Nodes_Node) Get_tailComment() LnsAny{ return self.tailComment }
func (self *Nodes_Node) Set_tailComment(arg1 LnsAny){ self.tailComment = arg1 }
func (self *Nodes_Node) Get_isLValue() bool{ return self.IsLValue }
func (self *Nodes_Node) Get_macroArgFlag() bool{ return self.macroArgFlag }
// 207: decl @lune.@base.@Nodes.Node.get_effectivePos
func (self *Nodes_Node) Get_effectivePos() *Types_Position {
    return self.pos
}

// 211: decl @lune.@base.@Nodes.Node.setLValue
func (self *Nodes_Node) SetLValue() {
    self.IsLValue = true
    
}

// 215: decl @lune.@base.@Nodes.Node.getPrefix
func (self *Nodes_Node) GetPrefix() LnsAny {
    return nil
}

// 219: DeclConstr
func (self *Nodes_Node) InitNodes_Node(id LnsInt,kind LnsInt,pos *Types_Position,macroArgFlag bool,expTypeList *LnsList) {
    self.IsLValue = false
    
    self.id = id
    
    self.kind = kind
    
    self.pos = pos
    
    self.expTypeList = expTypeList
    
    self.commentList = nil
    
    self.tailComment = nil
    
    self.macroArgFlag = macroArgFlag
    
}

// 232: decl @lune.@base.@Nodes.Node.addComment
func (self *Nodes_Node) AddComment(commentList *LnsList) {
    if commentList.Len() != 0{
        var workList *LnsList
        {
            __exp := self.commentList
            if __exp != nil {
                _exp := __exp.(*LnsList)
                workList = _exp
                
            } else {
                workList = NewLnsList([]LnsAny{})
                
                self.commentList = workList
                
            }
        }
        for _, _comment := range( commentList.Items ) {
            comment := _comment.(Types_TokenDownCast).ToTypes_Token()
            workList.Insert(Types_Token2Stem(comment))
        }
    }
}

// 248: decl @lune.@base.@Nodes.Node.get_expType
func (self *Nodes_Node) Get_expType() *Ast_TypeInfo {
    if self.expTypeList.Len() == 0{
        return Ast_builtinTypeNone
    }
    return self.expTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
}

// 255: decl @lune.@base.@Nodes.Node.addTokenList
func (self *Nodes_Node) AddTokenList(list *LnsList,kind LnsInt,txt string) {
    list.Insert(Types_Token2Stem(NewTypes_Token(kind, txt, self.pos, false, nil)))
}

// 259: decl @lune.@base.@Nodes.Node.setupLiteralTokenList
func (self *Nodes_Node) SetupLiteralTokenList(list *LnsList) bool {
    return false
}

// 263: decl @lune.@base.@Nodes.Node.getLiteral
func (self *Nodes_Node) GetLiteral()(LnsAny, LnsAny) {
    return nil, nil
}

// 266: decl @lune.@base.@Nodes.Node.processFilter
func (self *Nodes_Node) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
}

// 269: decl @lune.@base.@Nodes.Node.canBeLeft
func (self *Nodes_Node) CanBeLeft() bool {
    return false
}

// 273: decl @lune.@base.@Nodes.Node.canBeRight
func (self *Nodes_Node) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 277: decl @lune.@base.@Nodes.Node.canBeStatement
func (self *Nodes_Node) CanBeStatement() bool {
    return false
}

// 287: decl @lune.@base.@Nodes.Node.getBreakKind
func (self *Nodes_Node) GetBreakKind(checkMode LnsInt) LnsInt {
    return Nodes_BreakKind__None
}


// 294: decl @lune.@base.@Nodes.Node.hasNilAccess
func (self *Nodes_Node) HasNilAccess() bool {
    return false
}

// 298: decl @lune.@base.@Nodes.Node.isThreading
func (self *Nodes_Node) IsThreading() bool {
    return false
}


// 2097: decl @lune.@base.@Nodes.Node.getSymbolInfo
func (self *Nodes_Node) GetSymbolInfo() *LnsList {
    return Node_getSymbolInfo__processExpNode_9477_(self)
}


// declaration Class -- NamespaceInfo
type Nodes_NamespaceInfoMtd interface {
}
type Nodes_NamespaceInfo struct {
    Name string
    Scope *Ast_Scope
    TypeInfo *Ast_TypeInfo
    FP Nodes_NamespaceInfoMtd
}
func Nodes_NamespaceInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_NamespaceInfo).FP
}
type Nodes_NamespaceInfoDownCast interface {
    ToNodes_NamespaceInfo() *Nodes_NamespaceInfo
}
func Nodes_NamespaceInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_NamespaceInfoDownCast)
    if ok { return work.ToNodes_NamespaceInfo() }
    return nil
}
func (obj *Nodes_NamespaceInfo) ToNodes_NamespaceInfo() *Nodes_NamespaceInfo {
    return obj
}
func NewNodes_NamespaceInfo(arg1 string, arg2 *Ast_Scope, arg3 *Ast_TypeInfo) *Nodes_NamespaceInfo {
    obj := &Nodes_NamespaceInfo{}
    obj.FP = obj
    obj.InitNodes_NamespaceInfo(arg1, arg2, arg3)
    return obj
}
func (self *Nodes_NamespaceInfo) InitNodes_NamespaceInfo(arg1 string, arg2 *Ast_Scope, arg3 *Ast_TypeInfo) {
    self.Name = arg1
    self.Scope = arg2
    self.TypeInfo = arg3
}

// declaration Class -- DeclMacroInfo
type Nodes_DeclMacroInfoMtd interface {
    Get_argList() *LnsList
    Get_name() *Types_Token
    Get_pubFlag() bool
    Get_stmtBlock() LnsAny
    Get_tokenList() *LnsList
}
type Nodes_DeclMacroInfo struct {
    pubFlag bool
    name *Types_Token
    argList *LnsList
    stmtBlock LnsAny
    tokenList *LnsList
    FP Nodes_DeclMacroInfoMtd
}
func Nodes_DeclMacroInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclMacroInfo).FP
}
type Nodes_DeclMacroInfoDownCast interface {
    ToNodes_DeclMacroInfo() *Nodes_DeclMacroInfo
}
func Nodes_DeclMacroInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclMacroInfoDownCast)
    if ok { return work.ToNodes_DeclMacroInfo() }
    return nil
}
func (obj *Nodes_DeclMacroInfo) ToNodes_DeclMacroInfo() *Nodes_DeclMacroInfo {
    return obj
}
func NewNodes_DeclMacroInfo(arg1 bool, arg2 *Types_Token, arg3 *LnsList, arg4 LnsAny, arg5 *LnsList) *Nodes_DeclMacroInfo {
    obj := &Nodes_DeclMacroInfo{}
    obj.FP = obj
    obj.InitNodes_DeclMacroInfo(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_DeclMacroInfo) InitNodes_DeclMacroInfo(arg1 bool, arg2 *Types_Token, arg3 *LnsList, arg4 LnsAny, arg5 *LnsList) {
    self.pubFlag = arg1
    self.name = arg2
    self.argList = arg3
    self.stmtBlock = arg4
    self.tokenList = arg5
}
func (self *Nodes_DeclMacroInfo) Get_pubFlag() bool{ return self.pubFlag }
func (self *Nodes_DeclMacroInfo) Get_name() *Types_Token{ return self.name }
func (self *Nodes_DeclMacroInfo) Get_argList() *LnsList{ return self.argList }
func (self *Nodes_DeclMacroInfo) Get_stmtBlock() LnsAny{ return self.stmtBlock }
func (self *Nodes_DeclMacroInfo) Get_tokenList() *LnsList{ return self.tokenList }

// declaration Class -- NodeManager
type Nodes_NodeManagerMtd interface {
    MultiTo1(arg1 *Nodes_Node) *Nodes_Node
    AddNode(arg1 *Nodes_Node)
    DelNode(arg1 *Nodes_Node)
    GetAbbrNodeList() *LnsList
    GetAliasNodeList() *LnsList
    GetApplyNodeList() *LnsList
    GetBlankLineNodeList() *LnsList
    GetBlockNodeList() *LnsList
    GetBoxingNodeList() *LnsList
    GetBreakNodeList() *LnsList
    GetConvStatNodeList() *LnsList
    GetDeclAdvertiseNodeList() *LnsList
    GetDeclAlgeNodeList() *LnsList
    GetDeclArgDDDNodeList() *LnsList
    GetDeclArgNodeList() *LnsList
    GetDeclClassNodeList() *LnsList
    GetDeclConstrNodeList() *LnsList
    GetDeclDestrNodeList() *LnsList
    GetDeclEnumNodeList() *LnsList
    GetDeclFormNodeList() *LnsList
    GetDeclFuncNodeList() *LnsList
    GetDeclMacroNodeList() *LnsList
    GetDeclMemberNodeList() *LnsList
    GetDeclMethodNodeList() *LnsList
    GetDeclVarNodeList() *LnsList
    GetExpAccessMRetNodeList() *LnsList
    GetExpCallNodeList() *LnsList
    GetExpCallSuperCtorNodeList() *LnsList
    GetExpCallSuperNodeList() *LnsList
    GetExpCastNodeList() *LnsList
    GetExpListNodeList() *LnsList
    GetExpMRetNodeList() *LnsList
    GetExpMacroArgExpNodeList() *LnsList
    GetExpMacroExpNodeList() *LnsList
    GetExpMacroStatListNodeList() *LnsList
    GetExpMacroStatNodeList() *LnsList
    GetExpMultiTo1NodeList() *LnsList
    GetExpNewNodeList() *LnsList
    GetExpOmitEnumNodeList() *LnsList
    GetExpOp1NodeList() *LnsList
    GetExpOp2NodeList() *LnsList
    GetExpParenNodeList() *LnsList
    GetExpRefItemNodeList() *LnsList
    GetExpRefNodeList() *LnsList
    GetExpSetItemNodeList() *LnsList
    GetExpSetValNodeList() *LnsList
    GetExpSubDDDNodeList() *LnsList
    GetExpToDDDNodeList() *LnsList
    GetExpUnwrapNodeList() *LnsList
    GetForNodeList() *LnsList
    GetForeachNodeList() *LnsList
    GetForsortNodeList() *LnsList
    GetGetFieldNodeList() *LnsList
    GetIfNodeList() *LnsList
    GetIfUnwrapNodeList() *LnsList
    GetImportNodeList() *LnsList
    GetList(arg1 LnsInt) *LnsList
    GetLiteralArrayNodeList() *LnsList
    GetLiteralBoolNodeList() *LnsList
    GetLiteralCharNodeList() *LnsList
    GetLiteralIntNodeList() *LnsList
    GetLiteralListNodeList() *LnsList
    GetLiteralMapNodeList() *LnsList
    GetLiteralNilNodeList() *LnsList
    GetLiteralRealNodeList() *LnsList
    GetLiteralSetNodeList() *LnsList
    GetLiteralStringNodeList() *LnsList
    GetLiteralSymbolNodeList() *LnsList
    GetLuneControlNodeList() *LnsList
    GetLuneKindNodeList() *LnsList
    GetMatchNodeList() *LnsList
    GetNewAlgeValNodeList() *LnsList
    GetNoneNodeList() *LnsList
    GetProtoClassNodeList() *LnsList
    GetProtoMethodNodeList() *LnsList
    GetProvideNodeList() *LnsList
    GetRefFieldNodeList() *LnsList
    GetRefTypeNodeList() *LnsList
    GetRepeatNodeList() *LnsList
    GetReturnNodeList() *LnsList
    GetRootNodeList() *LnsList
    GetScopeNodeList() *LnsList
    GetStmtExpNodeList() *LnsList
    GetSubfileNodeList() *LnsList
    GetSwitchNodeList() *LnsList
    GetTestBlockNodeList() *LnsList
    GetTestCaseNodeList() *LnsList
    GetUnboxingNodeList() *LnsList
    GetUnwrapSetNodeList() *LnsList
    GetWhenNodeList() *LnsList
    GetWhileNodeList() *LnsList
    NextId() LnsInt
}
type Nodes_NodeManager struct {
    nodeKind2NodeList *LnsMap
    idSeed LnsInt
    FP Nodes_NodeManagerMtd
}
func Nodes_NodeManager2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_NodeManager).FP
}
type Nodes_NodeManagerDownCast interface {
    ToNodes_NodeManager() *Nodes_NodeManager
}
func Nodes_NodeManagerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_NodeManagerDownCast)
    if ok { return work.ToNodes_NodeManager() }
    return nil
}
func (obj *Nodes_NodeManager) ToNodes_NodeManager() *Nodes_NodeManager {
    return obj
}
func NewNodes_NodeManager() *Nodes_NodeManager {
    obj := &Nodes_NodeManager{}
    obj.FP = obj
    obj.InitNodes_NodeManager()
    return obj
}
// 405: DeclConstr
func (self *Nodes_NodeManager) InitNodes_NodeManager() {
    self.idSeed = 0
    
    self.nodeKind2NodeList = NewLnsMap( map[LnsAny]LnsAny{})
    
    for _kind, _ := range( Nodes_nodeKind2NameMap.Items ) {
        kind := _kind.(LnsInt)
        if Lns_op_not(self.nodeKind2NodeList.Items[kind]){
            self.nodeKind2NodeList.Set(kind,NewLnsList([]LnsAny{}))
        }
    }
}

// 415: decl @lune.@base.@Nodes.NodeManager.nextId
func (self *Nodes_NodeManager) NextId() LnsInt {
    self.idSeed = self.idSeed + 1
    
    return self.idSeed
}

// 420: decl @lune.@base.@Nodes.NodeManager.getList
func (self *Nodes_NodeManager) GetList(kind LnsInt) *LnsList {
    var list *LnsList
    
    {
        _list := self.nodeKind2NodeList.Items[kind]
        if _list == nil{
            return NewLnsList([]LnsAny{})
        } else {
            list = _list.(*LnsList)
        }
    }
    return list
}

// 426: decl @lune.@base.@Nodes.NodeManager.addNode
func (self *Nodes_NodeManager) AddNode(node *Nodes_Node) {
    var list *LnsList
    
    {
        _list := self.nodeKind2NodeList.Items[node.FP.Get_kind()]
        if _list == nil{
            list = NewLnsList([]LnsAny{})
            
            self.nodeKind2NodeList.Set(node.FP.Get_kind(),list)
        } else {
            list = _list.(*LnsList)
        }
    }
    list.Insert(Nodes_Node2Stem(node))
}

// 433: decl @lune.@base.@Nodes.NodeManager.delNode
func (self *Nodes_NodeManager) DelNode(node *Nodes_Node) {
    var list *LnsList
    
    {
        _list := self.nodeKind2NodeList.Items[node.FP.Get_kind()]
        if _list == nil{
            return 
        } else {
            list = _list.(*LnsList)
        }
    }
    var findIndex LnsInt
    findIndex = -1
    for _index, _item := range( list.Items ) {
        index := _index + 1
        item := _item.(Nodes_NodeDownCast).ToNodes_Node()
        if item == node{
            findIndex = index
            
            break
        }
    }
    if findIndex != -1{
        list.Remove(findIndex)
    }
}

// 1: decl @lune.@base.@Nodes.NodeManager.getNoneNodeList
func (self *Nodes_NodeManager) GetNoneNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(0))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getConvStatNodeList
func (self *Nodes_NodeManager) GetConvStatNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(1))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getBlankLineNodeList
func (self *Nodes_NodeManager) GetBlankLineNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(2))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getSubfileNodeList
func (self *Nodes_NodeManager) GetSubfileNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(3))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getImportNodeList
func (self *Nodes_NodeManager) GetImportNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(4))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getRootNodeList
func (self *Nodes_NodeManager) GetRootNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(5))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getRefTypeNodeList
func (self *Nodes_NodeManager) GetRefTypeNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(6))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getBlockNodeList
func (self *Nodes_NodeManager) GetBlockNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(7))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getScopeNodeList
func (self *Nodes_NodeManager) GetScopeNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(8))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getIfNodeList
func (self *Nodes_NodeManager) GetIfNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(9))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpListNodeList
func (self *Nodes_NodeManager) GetExpListNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(10))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getSwitchNodeList
func (self *Nodes_NodeManager) GetSwitchNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(11))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getWhileNodeList
func (self *Nodes_NodeManager) GetWhileNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(12))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getRepeatNodeList
func (self *Nodes_NodeManager) GetRepeatNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(13))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getForNodeList
func (self *Nodes_NodeManager) GetForNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(14))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getApplyNodeList
func (self *Nodes_NodeManager) GetApplyNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(15))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getForeachNodeList
func (self *Nodes_NodeManager) GetForeachNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(16))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getForsortNodeList
func (self *Nodes_NodeManager) GetForsortNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(17))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getReturnNodeList
func (self *Nodes_NodeManager) GetReturnNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(18))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getBreakNodeList
func (self *Nodes_NodeManager) GetBreakNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(19))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getProvideNodeList
func (self *Nodes_NodeManager) GetProvideNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(20))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpNewNodeList
func (self *Nodes_NodeManager) GetExpNewNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(21))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpUnwrapNodeList
func (self *Nodes_NodeManager) GetExpUnwrapNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(22))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpRefNodeList
func (self *Nodes_NodeManager) GetExpRefNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(23))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpSetValNodeList
func (self *Nodes_NodeManager) GetExpSetValNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(24))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpSetItemNodeList
func (self *Nodes_NodeManager) GetExpSetItemNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(25))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpOp2NodeList
func (self *Nodes_NodeManager) GetExpOp2NodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(26))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getUnwrapSetNodeList
func (self *Nodes_NodeManager) GetUnwrapSetNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(27))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getIfUnwrapNodeList
func (self *Nodes_NodeManager) GetIfUnwrapNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(28))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getWhenNodeList
func (self *Nodes_NodeManager) GetWhenNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(29))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpCastNodeList
func (self *Nodes_NodeManager) GetExpCastNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(30))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpToDDDNodeList
func (self *Nodes_NodeManager) GetExpToDDDNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(31))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpSubDDDNodeList
func (self *Nodes_NodeManager) GetExpSubDDDNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(32))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpOp1NodeList
func (self *Nodes_NodeManager) GetExpOp1NodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(33))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpRefItemNodeList
func (self *Nodes_NodeManager) GetExpRefItemNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(34))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpCallNodeList
func (self *Nodes_NodeManager) GetExpCallNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(35))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpMRetNodeList
func (self *Nodes_NodeManager) GetExpMRetNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(36))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpAccessMRetNodeList
func (self *Nodes_NodeManager) GetExpAccessMRetNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(37))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpMultiTo1NodeList
func (self *Nodes_NodeManager) GetExpMultiTo1NodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(38))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpParenNodeList
func (self *Nodes_NodeManager) GetExpParenNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(39))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpMacroExpNodeList
func (self *Nodes_NodeManager) GetExpMacroExpNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(40))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpMacroStatNodeList
func (self *Nodes_NodeManager) GetExpMacroStatNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(41))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpMacroArgExpNodeList
func (self *Nodes_NodeManager) GetExpMacroArgExpNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(42))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getStmtExpNodeList
func (self *Nodes_NodeManager) GetStmtExpNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(43))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpMacroStatListNodeList
func (self *Nodes_NodeManager) GetExpMacroStatListNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(44))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpOmitEnumNodeList
func (self *Nodes_NodeManager) GetExpOmitEnumNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(45))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getRefFieldNodeList
func (self *Nodes_NodeManager) GetRefFieldNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(46))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getGetFieldNodeList
func (self *Nodes_NodeManager) GetGetFieldNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(47))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getAliasNodeList
func (self *Nodes_NodeManager) GetAliasNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(48))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclVarNodeList
func (self *Nodes_NodeManager) GetDeclVarNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(49))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclFormNodeList
func (self *Nodes_NodeManager) GetDeclFormNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(50))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclFuncNodeList
func (self *Nodes_NodeManager) GetDeclFuncNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(51))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclMethodNodeList
func (self *Nodes_NodeManager) GetDeclMethodNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(52))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getProtoMethodNodeList
func (self *Nodes_NodeManager) GetProtoMethodNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(53))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclConstrNodeList
func (self *Nodes_NodeManager) GetDeclConstrNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(54))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclDestrNodeList
func (self *Nodes_NodeManager) GetDeclDestrNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(55))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpCallSuperCtorNodeList
func (self *Nodes_NodeManager) GetExpCallSuperCtorNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(56))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getExpCallSuperNodeList
func (self *Nodes_NodeManager) GetExpCallSuperNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(57))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclMemberNodeList
func (self *Nodes_NodeManager) GetDeclMemberNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(58))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclArgNodeList
func (self *Nodes_NodeManager) GetDeclArgNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(59))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclArgDDDNodeList
func (self *Nodes_NodeManager) GetDeclArgDDDNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(60))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclAdvertiseNodeList
func (self *Nodes_NodeManager) GetDeclAdvertiseNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(61))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getProtoClassNodeList
func (self *Nodes_NodeManager) GetProtoClassNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(62))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclClassNodeList
func (self *Nodes_NodeManager) GetDeclClassNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(63))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclEnumNodeList
func (self *Nodes_NodeManager) GetDeclEnumNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(64))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclAlgeNodeList
func (self *Nodes_NodeManager) GetDeclAlgeNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(65))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getNewAlgeValNodeList
func (self *Nodes_NodeManager) GetNewAlgeValNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(66))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLuneControlNodeList
func (self *Nodes_NodeManager) GetLuneControlNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(67))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getMatchNodeList
func (self *Nodes_NodeManager) GetMatchNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(68))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLuneKindNodeList
func (self *Nodes_NodeManager) GetLuneKindNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(69))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getDeclMacroNodeList
func (self *Nodes_NodeManager) GetDeclMacroNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(70))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getTestCaseNodeList
func (self *Nodes_NodeManager) GetTestCaseNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(71))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getTestBlockNodeList
func (self *Nodes_NodeManager) GetTestBlockNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(72))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getAbbrNodeList
func (self *Nodes_NodeManager) GetAbbrNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(73))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getBoxingNodeList
func (self *Nodes_NodeManager) GetBoxingNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(74))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getUnboxingNodeList
func (self *Nodes_NodeManager) GetUnboxingNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(75))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralNilNodeList
func (self *Nodes_NodeManager) GetLiteralNilNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(76))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralCharNodeList
func (self *Nodes_NodeManager) GetLiteralCharNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(77))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralIntNodeList
func (self *Nodes_NodeManager) GetLiteralIntNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(78))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralRealNodeList
func (self *Nodes_NodeManager) GetLiteralRealNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(79))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralArrayNodeList
func (self *Nodes_NodeManager) GetLiteralArrayNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(80))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralListNodeList
func (self *Nodes_NodeManager) GetLiteralListNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(81))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralSetNodeList
func (self *Nodes_NodeManager) GetLiteralSetNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(82))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralMapNodeList
func (self *Nodes_NodeManager) GetLiteralMapNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(83))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralStringNodeList
func (self *Nodes_NodeManager) GetLiteralStringNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(84))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralBoolNodeList
func (self *Nodes_NodeManager) GetLiteralBoolNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(85))
}

// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralSymbolNodeList
func (self *Nodes_NodeManager) GetLiteralSymbolNodeList() *LnsList {
    return (*LnsList)(self.FP.GetList(86))
}

// 2755: decl @lune.@base.@Nodes.NodeManager.MultiTo1
func (self *Nodes_NodeManager) MultiTo1(node *Nodes_Node) *Nodes_Node {
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType()
    if node.FP.Get_expTypeList().Len() > 1{
        return &Nodes_ExpMultiTo1Node_create(self, node.FP.Get_pos(), node.FP.Get_macroArgFlag(), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(expType)}), node).Nodes_Node
    } else if expType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
        return &Nodes_ExpMultiTo1Node_create(self, node.FP.Get_pos(), node.FP.Get_macroArgFlag(), expType.FP.Get_itemTypeInfoList(), node).Nodes_Node
    }
    return node
}


// declaration Class -- NodeKind
type Nodes_NodeKindMtd interface {
}
type Nodes_NodeKind struct {
    FP Nodes_NodeKindMtd
}
func Nodes_NodeKind2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_NodeKind).FP
}
type Nodes_NodeKindDownCast interface {
    ToNodes_NodeKind() *Nodes_NodeKind
}
func Nodes_NodeKindDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_NodeKindDownCast)
    if ok { return work.ToNodes_NodeKind() }
    return nil
}
func (obj *Nodes_NodeKind) ToNodes_NodeKind() *Nodes_NodeKind {
    return obj
}
func NewNodes_NodeKind() *Nodes_NodeKind {
    obj := &Nodes_NodeKind{}
    obj.FP = obj
    obj.InitNodes_NodeKind()
    return obj
}
func (self *Nodes_NodeKind) InitNodes_NodeKind() {
}
// 667: decl @lune.@base.@Nodes.NodeKind.get_None
func Nodes_NodeKind_get_None() LnsInt {
    return 0
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ConvStat
func Nodes_NodeKind_get_ConvStat() LnsInt {
    return 1
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_BlankLine
func Nodes_NodeKind_get_BlankLine() LnsInt {
    return 2
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Subfile
func Nodes_NodeKind_get_Subfile() LnsInt {
    return 3
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Import
func Nodes_NodeKind_get_Import() LnsInt {
    return 4
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Root
func Nodes_NodeKind_get_Root() LnsInt {
    return 5
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_RefType
func Nodes_NodeKind_get_RefType() LnsInt {
    return 6
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Block
func Nodes_NodeKind_get_Block() LnsInt {
    return 7
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Scope
func Nodes_NodeKind_get_Scope() LnsInt {
    return 8
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_If
func Nodes_NodeKind_get_If() LnsInt {
    return 9
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpList
func Nodes_NodeKind_get_ExpList() LnsInt {
    return 10
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Switch
func Nodes_NodeKind_get_Switch() LnsInt {
    return 11
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_While
func Nodes_NodeKind_get_While() LnsInt {
    return 12
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Repeat
func Nodes_NodeKind_get_Repeat() LnsInt {
    return 13
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_For
func Nodes_NodeKind_get_For() LnsInt {
    return 14
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Apply
func Nodes_NodeKind_get_Apply() LnsInt {
    return 15
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Foreach
func Nodes_NodeKind_get_Foreach() LnsInt {
    return 16
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Forsort
func Nodes_NodeKind_get_Forsort() LnsInt {
    return 17
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Return
func Nodes_NodeKind_get_Return() LnsInt {
    return 18
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Break
func Nodes_NodeKind_get_Break() LnsInt {
    return 19
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Provide
func Nodes_NodeKind_get_Provide() LnsInt {
    return 20
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpNew
func Nodes_NodeKind_get_ExpNew() LnsInt {
    return 21
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpUnwrap
func Nodes_NodeKind_get_ExpUnwrap() LnsInt {
    return 22
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpRef
func Nodes_NodeKind_get_ExpRef() LnsInt {
    return 23
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpSetVal
func Nodes_NodeKind_get_ExpSetVal() LnsInt {
    return 24
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpSetItem
func Nodes_NodeKind_get_ExpSetItem() LnsInt {
    return 25
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpOp2
func Nodes_NodeKind_get_ExpOp2() LnsInt {
    return 26
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_UnwrapSet
func Nodes_NodeKind_get_UnwrapSet() LnsInt {
    return 27
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_IfUnwrap
func Nodes_NodeKind_get_IfUnwrap() LnsInt {
    return 28
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_When
func Nodes_NodeKind_get_When() LnsInt {
    return 29
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpCast
func Nodes_NodeKind_get_ExpCast() LnsInt {
    return 30
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpToDDD
func Nodes_NodeKind_get_ExpToDDD() LnsInt {
    return 31
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpSubDDD
func Nodes_NodeKind_get_ExpSubDDD() LnsInt {
    return 32
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpOp1
func Nodes_NodeKind_get_ExpOp1() LnsInt {
    return 33
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpRefItem
func Nodes_NodeKind_get_ExpRefItem() LnsInt {
    return 34
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpCall
func Nodes_NodeKind_get_ExpCall() LnsInt {
    return 35
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpMRet
func Nodes_NodeKind_get_ExpMRet() LnsInt {
    return 36
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpAccessMRet
func Nodes_NodeKind_get_ExpAccessMRet() LnsInt {
    return 37
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpMultiTo1
func Nodes_NodeKind_get_ExpMultiTo1() LnsInt {
    return 38
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpParen
func Nodes_NodeKind_get_ExpParen() LnsInt {
    return 39
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpMacroExp
func Nodes_NodeKind_get_ExpMacroExp() LnsInt {
    return 40
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpMacroStat
func Nodes_NodeKind_get_ExpMacroStat() LnsInt {
    return 41
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpMacroArgExp
func Nodes_NodeKind_get_ExpMacroArgExp() LnsInt {
    return 42
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_StmtExp
func Nodes_NodeKind_get_StmtExp() LnsInt {
    return 43
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpMacroStatList
func Nodes_NodeKind_get_ExpMacroStatList() LnsInt {
    return 44
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpOmitEnum
func Nodes_NodeKind_get_ExpOmitEnum() LnsInt {
    return 45
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_RefField
func Nodes_NodeKind_get_RefField() LnsInt {
    return 46
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_GetField
func Nodes_NodeKind_get_GetField() LnsInt {
    return 47
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Alias
func Nodes_NodeKind_get_Alias() LnsInt {
    return 48
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclVar
func Nodes_NodeKind_get_DeclVar() LnsInt {
    return 49
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclForm
func Nodes_NodeKind_get_DeclForm() LnsInt {
    return 50
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclFunc
func Nodes_NodeKind_get_DeclFunc() LnsInt {
    return 51
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclMethod
func Nodes_NodeKind_get_DeclMethod() LnsInt {
    return 52
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ProtoMethod
func Nodes_NodeKind_get_ProtoMethod() LnsInt {
    return 53
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclConstr
func Nodes_NodeKind_get_DeclConstr() LnsInt {
    return 54
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclDestr
func Nodes_NodeKind_get_DeclDestr() LnsInt {
    return 55
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpCallSuperCtor
func Nodes_NodeKind_get_ExpCallSuperCtor() LnsInt {
    return 56
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ExpCallSuper
func Nodes_NodeKind_get_ExpCallSuper() LnsInt {
    return 57
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclMember
func Nodes_NodeKind_get_DeclMember() LnsInt {
    return 58
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclArg
func Nodes_NodeKind_get_DeclArg() LnsInt {
    return 59
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclArgDDD
func Nodes_NodeKind_get_DeclArgDDD() LnsInt {
    return 60
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclAdvertise
func Nodes_NodeKind_get_DeclAdvertise() LnsInt {
    return 61
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_ProtoClass
func Nodes_NodeKind_get_ProtoClass() LnsInt {
    return 62
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclClass
func Nodes_NodeKind_get_DeclClass() LnsInt {
    return 63
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclEnum
func Nodes_NodeKind_get_DeclEnum() LnsInt {
    return 64
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclAlge
func Nodes_NodeKind_get_DeclAlge() LnsInt {
    return 65
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_NewAlgeVal
func Nodes_NodeKind_get_NewAlgeVal() LnsInt {
    return 66
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LuneControl
func Nodes_NodeKind_get_LuneControl() LnsInt {
    return 67
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Match
func Nodes_NodeKind_get_Match() LnsInt {
    return 68
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LuneKind
func Nodes_NodeKind_get_LuneKind() LnsInt {
    return 69
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_DeclMacro
func Nodes_NodeKind_get_DeclMacro() LnsInt {
    return 70
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_TestCase
func Nodes_NodeKind_get_TestCase() LnsInt {
    return 71
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_TestBlock
func Nodes_NodeKind_get_TestBlock() LnsInt {
    return 72
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Abbr
func Nodes_NodeKind_get_Abbr() LnsInt {
    return 73
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Boxing
func Nodes_NodeKind_get_Boxing() LnsInt {
    return 74
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_Unboxing
func Nodes_NodeKind_get_Unboxing() LnsInt {
    return 75
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralNil
func Nodes_NodeKind_get_LiteralNil() LnsInt {
    return 76
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralChar
func Nodes_NodeKind_get_LiteralChar() LnsInt {
    return 77
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralInt
func Nodes_NodeKind_get_LiteralInt() LnsInt {
    return 78
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralReal
func Nodes_NodeKind_get_LiteralReal() LnsInt {
    return 79
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralArray
func Nodes_NodeKind_get_LiteralArray() LnsInt {
    return 80
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralList
func Nodes_NodeKind_get_LiteralList() LnsInt {
    return 81
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralSet
func Nodes_NodeKind_get_LiteralSet() LnsInt {
    return 82
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralMap
func Nodes_NodeKind_get_LiteralMap() LnsInt {
    return 83
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralString
func Nodes_NodeKind_get_LiteralString() LnsInt {
    return 84
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralBool
func Nodes_NodeKind_get_LiteralBool() LnsInt {
    return 85
}

// 667: decl @lune.@base.@Nodes.NodeKind.get_LiteralSymbol
func Nodes_NodeKind_get_LiteralSymbol() LnsInt {
    return 86
}


// declaration Class -- NoneNode
type Nodes_NoneNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_NoneNode struct {
    Nodes_Node
    FP Nodes_NoneNodeMtd
}
func Nodes_NoneNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_NoneNode).FP
}
type Nodes_NoneNodeDownCast interface {
    ToNodes_NoneNode() *Nodes_NoneNode
}
func Nodes_NoneNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_NoneNodeDownCast)
    if ok { return work.ToNodes_NoneNode() }
    return nil
}
func (obj *Nodes_NoneNode) ToNodes_NoneNode() *Nodes_NoneNode {
    return obj
}
func NewNodes_NoneNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList) *Nodes_NoneNode {
    obj := &Nodes_NoneNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_NoneNode(arg1, arg2, arg3, arg4)
    return obj
}
// 1: decl @lune.@base.@Nodes.NoneNode.processFilter
func (self *Nodes_NoneNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessNone(self, opt)
}

// 1: decl @lune.@base.@Nodes.NoneNode.canBeRight
func (self *Nodes_NoneNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.NoneNode.canBeLeft
func (self *Nodes_NoneNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.NoneNode.canBeStatement
func (self *Nodes_NoneNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_NoneNode) InitNodes_NoneNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(id, 0, pos, macroArgFlag, typeList)
}

// 681: decl @lune.@base.@Nodes.NoneNode.create
func Nodes_NoneNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList) *Nodes_NoneNode {
    var node *Nodes_NoneNode
    node = NewNodes_NoneNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.NoneNode.visit
func (self *Nodes_NoneNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- ConvStatNode
type Nodes_ConvStatNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_txt() string
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ConvStatNode struct {
    Nodes_Node
    txt string
    FP Nodes_ConvStatNodeMtd
}
func Nodes_ConvStatNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ConvStatNode).FP
}
type Nodes_ConvStatNodeDownCast interface {
    ToNodes_ConvStatNode() *Nodes_ConvStatNode
}
func Nodes_ConvStatNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ConvStatNodeDownCast)
    if ok { return work.ToNodes_ConvStatNode() }
    return nil
}
func (obj *Nodes_ConvStatNode) ToNodes_ConvStatNode() *Nodes_ConvStatNode {
    return obj
}
func NewNodes_ConvStatNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 string) *Nodes_ConvStatNode {
    obj := &Nodes_ConvStatNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ConvStatNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ConvStatNode) Get_txt() string{ return self.txt }
// 1: decl @lune.@base.@Nodes.ConvStatNode.processFilter
func (self *Nodes_ConvStatNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessConvStat(self, opt)
}

// 1: decl @lune.@base.@Nodes.ConvStatNode.canBeRight
func (self *Nodes_ConvStatNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ConvStatNode.canBeLeft
func (self *Nodes_ConvStatNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ConvStatNode.canBeStatement
func (self *Nodes_ConvStatNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ConvStatNode) InitNodes_ConvStatNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,txt string) {
    self.InitNodes_Node(id, 1, pos, macroArgFlag, typeList)
    self.txt = txt
    
}

// 681: decl @lune.@base.@Nodes.ConvStatNode.create
func Nodes_ConvStatNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,txt string) *Nodes_ConvStatNode {
    var node *Nodes_ConvStatNode
    node = NewNodes_ConvStatNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, txt)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ConvStatNode.visit
func (self *Nodes_ConvStatNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- BlankLineNode
type Nodes_BlankLineNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_lineNum() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_BlankLineNode struct {
    Nodes_Node
    lineNum LnsInt
    FP Nodes_BlankLineNodeMtd
}
func Nodes_BlankLineNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_BlankLineNode).FP
}
type Nodes_BlankLineNodeDownCast interface {
    ToNodes_BlankLineNode() *Nodes_BlankLineNode
}
func Nodes_BlankLineNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_BlankLineNodeDownCast)
    if ok { return work.ToNodes_BlankLineNode() }
    return nil
}
func (obj *Nodes_BlankLineNode) ToNodes_BlankLineNode() *Nodes_BlankLineNode {
    return obj
}
func NewNodes_BlankLineNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsInt) *Nodes_BlankLineNode {
    obj := &Nodes_BlankLineNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_BlankLineNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_BlankLineNode) Get_lineNum() LnsInt{ return self.lineNum }
// 1: decl @lune.@base.@Nodes.BlankLineNode.processFilter
func (self *Nodes_BlankLineNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessBlankLine(self, opt)
}

// 1: decl @lune.@base.@Nodes.BlankLineNode.canBeRight
func (self *Nodes_BlankLineNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.BlankLineNode.canBeLeft
func (self *Nodes_BlankLineNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.BlankLineNode.canBeStatement
func (self *Nodes_BlankLineNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_BlankLineNode) InitNodes_BlankLineNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,lineNum LnsInt) {
    self.InitNodes_Node(id, 2, pos, macroArgFlag, typeList)
    self.lineNum = lineNum
    
}

// 681: decl @lune.@base.@Nodes.BlankLineNode.create
func Nodes_BlankLineNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,lineNum LnsInt) *Nodes_BlankLineNode {
    var node *Nodes_BlankLineNode
    node = NewNodes_BlankLineNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, lineNum)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.BlankLineNode.visit
func (self *Nodes_BlankLineNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- SubfileNode
type Nodes_SubfileNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_usePath() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_SubfileNode struct {
    Nodes_Node
    usePath LnsAny
    FP Nodes_SubfileNodeMtd
}
func Nodes_SubfileNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_SubfileNode).FP
}
type Nodes_SubfileNodeDownCast interface {
    ToNodes_SubfileNode() *Nodes_SubfileNode
}
func Nodes_SubfileNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_SubfileNodeDownCast)
    if ok { return work.ToNodes_SubfileNode() }
    return nil
}
func (obj *Nodes_SubfileNode) ToNodes_SubfileNode() *Nodes_SubfileNode {
    return obj
}
func NewNodes_SubfileNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsAny) *Nodes_SubfileNode {
    obj := &Nodes_SubfileNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_SubfileNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_SubfileNode) Get_usePath() LnsAny{ return self.usePath }
// 1: decl @lune.@base.@Nodes.SubfileNode.processFilter
func (self *Nodes_SubfileNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessSubfile(self, opt)
}

// 1: decl @lune.@base.@Nodes.SubfileNode.canBeRight
func (self *Nodes_SubfileNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.SubfileNode.canBeLeft
func (self *Nodes_SubfileNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.SubfileNode.canBeStatement
func (self *Nodes_SubfileNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_SubfileNode) InitNodes_SubfileNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,usePath LnsAny) {
    self.InitNodes_Node(id, 3, pos, macroArgFlag, typeList)
    self.usePath = usePath
    
}

// 681: decl @lune.@base.@Nodes.SubfileNode.create
func Nodes_SubfileNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,usePath LnsAny) *Nodes_SubfileNode {
    var node *Nodes_SubfileNode
    node = NewNodes_SubfileNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, usePath)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.SubfileNode.visit
func (self *Nodes_SubfileNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- ImportNode
type Nodes_ImportNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_assignName() string
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_lazy() LnsInt
    Get_macroArgFlag() bool
    Get_modulePath() string
    Get_moduleTypeInfo() *Ast_TypeInfo
    Get_pos() *Types_Position
    Get_symbolInfo() *Ast_SymbolInfo
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ImportNode struct {
    Nodes_Node
    modulePath string
    lazy LnsInt
    assignName string
    symbolInfo *Ast_SymbolInfo
    moduleTypeInfo *Ast_TypeInfo
    FP Nodes_ImportNodeMtd
}
func Nodes_ImportNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ImportNode).FP
}
type Nodes_ImportNodeDownCast interface {
    ToNodes_ImportNode() *Nodes_ImportNode
}
func Nodes_ImportNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ImportNodeDownCast)
    if ok { return work.ToNodes_ImportNode() }
    return nil
}
func (obj *Nodes_ImportNode) ToNodes_ImportNode() *Nodes_ImportNode {
    return obj
}
func NewNodes_ImportNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 string, arg6 LnsInt, arg7 string, arg8 *Ast_SymbolInfo, arg9 *Ast_TypeInfo) *Nodes_ImportNode {
    obj := &Nodes_ImportNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ImportNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ImportNode) Get_modulePath() string{ return self.modulePath }
func (self *Nodes_ImportNode) Get_lazy() LnsInt{ return self.lazy }
func (self *Nodes_ImportNode) Get_assignName() string{ return self.assignName }
func (self *Nodes_ImportNode) Get_symbolInfo() *Ast_SymbolInfo{ return self.symbolInfo }
func (self *Nodes_ImportNode) Get_moduleTypeInfo() *Ast_TypeInfo{ return self.moduleTypeInfo }
// 1: decl @lune.@base.@Nodes.ImportNode.processFilter
func (self *Nodes_ImportNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessImport(self, opt)
}

// 1: decl @lune.@base.@Nodes.ImportNode.canBeRight
func (self *Nodes_ImportNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ImportNode.canBeLeft
func (self *Nodes_ImportNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ImportNode.canBeStatement
func (self *Nodes_ImportNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ImportNode) InitNodes_ImportNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,modulePath string,lazy LnsInt,assignName string,symbolInfo *Ast_SymbolInfo,moduleTypeInfo *Ast_TypeInfo) {
    self.InitNodes_Node(id, 4, pos, macroArgFlag, typeList)
    self.modulePath = modulePath
    
    self.lazy = lazy
    
    self.assignName = assignName
    
    self.symbolInfo = symbolInfo
    
    self.moduleTypeInfo = moduleTypeInfo
    
}

// 681: decl @lune.@base.@Nodes.ImportNode.create
func Nodes_ImportNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,modulePath string,lazy LnsInt,assignName string,symbolInfo *Ast_SymbolInfo,moduleTypeInfo *Ast_TypeInfo) *Nodes_ImportNode {
    var node *Nodes_ImportNode
    node = NewNodes_ImportNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, modulePath, lazy, assignName, symbolInfo, moduleTypeInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ImportNode.visit
func (self *Nodes_ImportNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- LuneHelperInfo
type Nodes_LuneHelperInfoMtd interface {
}
type Nodes_LuneHelperInfo struct {
    UseNilAccess bool
    UseUnwrapExp bool
    HasMappingClassDef bool
    UseLoad bool
    UseUnpack bool
    UseAlge bool
    UseSet bool
    CallAnonymous bool
    PragmaSet *LnsSet
    UseLazyLoad bool
    UseLazyRequire bool
    FP Nodes_LuneHelperInfoMtd
}
func Nodes_LuneHelperInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LuneHelperInfo).FP
}
type Nodes_LuneHelperInfoDownCast interface {
    ToNodes_LuneHelperInfo() *Nodes_LuneHelperInfo
}
func Nodes_LuneHelperInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LuneHelperInfoDownCast)
    if ok { return work.ToNodes_LuneHelperInfo() }
    return nil
}
func (obj *Nodes_LuneHelperInfo) ToNodes_LuneHelperInfo() *Nodes_LuneHelperInfo {
    return obj
}
func NewNodes_LuneHelperInfo() *Nodes_LuneHelperInfo {
    obj := &Nodes_LuneHelperInfo{}
    obj.FP = obj
    obj.InitNodes_LuneHelperInfo()
    return obj
}
// 740: DeclConstr
func (self *Nodes_LuneHelperInfo) InitNodes_LuneHelperInfo() {
    self.UseNilAccess = false
    
    self.UseUnwrapExp = false
    
    self.HasMappingClassDef = false
    
    self.UseLoad = false
    
    self.UseUnpack = false
    
    self.UseAlge = false
    
    self.UseSet = false
    
    self.CallAnonymous = false
    
    self.PragmaSet = NewLnsSet([]LnsAny{})
    
    self.UseLazyLoad = false
    
    self.UseLazyRequire = false
    
}


// declaration Class -- ModuleInfo
type Nodes_ModuleInfoMtd interface {
    Assign(arg1 string) *Nodes_ModuleInfo
    Get_assignName() string
    Get_fullName() string
    Get_importId2localTypeInfoMap() *LnsMap
    Get_importedAliasMap() *LnsMap
    Get_localTypeInfo2importIdMap() *LnsMap
    Get_moduleId() *FrontInterface_ModuleId
    Get_modulePath() string
}
type Nodes_ModuleInfo struct {
    fullName string
    localTypeInfo2importIdMap *LnsMap
    importId2localTypeInfoMap *LnsMap
    assignName string
    moduleId *FrontInterface_ModuleId
    importedAliasMap *LnsMap
    FP Nodes_ModuleInfoMtd
}
func Nodes_ModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ModuleInfo).FP
}
type Nodes_ModuleInfoDownCast interface {
    ToNodes_ModuleInfo() *Nodes_ModuleInfo
}
func Nodes_ModuleInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ModuleInfoDownCast)
    if ok { return work.ToNodes_ModuleInfo() }
    return nil
}
func (obj *Nodes_ModuleInfo) ToNodes_ModuleInfo() *Nodes_ModuleInfo {
    return obj
}
func NewNodes_ModuleInfo(arg1 string, arg2 string, arg3 *LnsMap, arg4 *FrontInterface_ModuleId, arg5 *LnsMap) *Nodes_ModuleInfo {
    obj := &Nodes_ModuleInfo{}
    obj.FP = obj
    obj.InitNodes_ModuleInfo(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ModuleInfo) Get_fullName() string{ return self.fullName }
func (self *Nodes_ModuleInfo) Get_localTypeInfo2importIdMap() *LnsMap{ return self.localTypeInfo2importIdMap }
func (self *Nodes_ModuleInfo) Get_importId2localTypeInfoMap() *LnsMap{ return self.importId2localTypeInfoMap }
func (self *Nodes_ModuleInfo) Get_assignName() string{ return self.assignName }
func (self *Nodes_ModuleInfo) Get_moduleId() *FrontInterface_ModuleId{ return self.moduleId }
func (self *Nodes_ModuleInfo) Get_importedAliasMap() *LnsMap{ return self.importedAliasMap }
// 763: DeclConstr
func (self *Nodes_ModuleInfo) InitNodes_ModuleInfo(fullName string,assignName string,idMap *LnsMap,moduleId *FrontInterface_ModuleId,importedAliasMap *LnsMap) {
    self.moduleId = moduleId
    
    self.fullName = fullName
    
    self.assignName = assignName
    
    self.localTypeInfo2importIdMap = idMap
    
    self.importId2localTypeInfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    for _typeInfo, _importId := range( idMap.Items ) {
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        importId := _importId.(LnsInt)
        self.importId2localTypeInfoMap.Set(importId,typeInfo)
    }
    self.importedAliasMap = importedAliasMap
    
}

// 778: decl @lune.@base.@Nodes.ModuleInfo.get_modulePath
func (self *Nodes_ModuleInfo) Get_modulePath() string {
    return self.fullName
}

// 782: decl @lune.@base.@Nodes.ModuleInfo.assign
func (self *Nodes_ModuleInfo) Assign(assignName string) *Nodes_ModuleInfo {
    return NewNodes_ModuleInfo(self.fullName, assignName, self.localTypeInfo2importIdMap, self.moduleId, self.importedAliasMap)
}


// declaration Class -- MacroValInfo
type Nodes_MacroValInfoMtd interface {
}
type Nodes_MacroValInfo struct {
    Val LnsAny
    TypeInfo *Ast_TypeInfo
    ArgNode LnsAny
    FP Nodes_MacroValInfoMtd
}
func Nodes_MacroValInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_MacroValInfo).FP
}
type Nodes_MacroValInfoDownCast interface {
    ToNodes_MacroValInfo() *Nodes_MacroValInfo
}
func Nodes_MacroValInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_MacroValInfoDownCast)
    if ok { return work.ToNodes_MacroValInfo() }
    return nil
}
func (obj *Nodes_MacroValInfo) ToNodes_MacroValInfo() *Nodes_MacroValInfo {
    return obj
}
func NewNodes_MacroValInfo(arg1 LnsAny, arg2 *Ast_TypeInfo, arg3 LnsAny) *Nodes_MacroValInfo {
    obj := &Nodes_MacroValInfo{}
    obj.FP = obj
    obj.InitNodes_MacroValInfo(arg1, arg2, arg3)
    return obj
}
func (self *Nodes_MacroValInfo) InitNodes_MacroValInfo(arg1 LnsAny, arg2 *Ast_TypeInfo, arg3 LnsAny) {
    self.Val = arg1
    self.TypeInfo = arg2
    self.ArgNode = arg3
}

// declaration Class -- MacroArgInfo
type Nodes_MacroArgInfoMtd interface {
    Get_name() string
    Get_typeInfo() *Ast_TypeInfo
}
type Nodes_MacroArgInfo struct {
    name string
    typeInfo *Ast_TypeInfo
    FP Nodes_MacroArgInfoMtd
}
func Nodes_MacroArgInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_MacroArgInfo).FP
}
type Nodes_MacroArgInfoDownCast interface {
    ToNodes_MacroArgInfo() *Nodes_MacroArgInfo
}
func Nodes_MacroArgInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_MacroArgInfoDownCast)
    if ok { return work.ToNodes_MacroArgInfo() }
    return nil
}
func (obj *Nodes_MacroArgInfo) ToNodes_MacroArgInfo() *Nodes_MacroArgInfo {
    return obj
}
func NewNodes_MacroArgInfo(arg1 string, arg2 *Ast_TypeInfo) *Nodes_MacroArgInfo {
    obj := &Nodes_MacroArgInfo{}
    obj.FP = obj
    obj.InitNodes_MacroArgInfo(arg1, arg2)
    return obj
}
func (self *Nodes_MacroArgInfo) InitNodes_MacroArgInfo(arg1 string, arg2 *Ast_TypeInfo) {
    self.name = arg1
    self.typeInfo = arg2
}
func (self *Nodes_MacroArgInfo) Get_name() string{ return self.name }
func (self *Nodes_MacroArgInfo) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }

// declaration Class -- MacroInfo
type Nodes_MacroInfoMtd interface {
    GetArgList() *LnsList
    GetTokenList() *LnsList
    Get_name() string
}
type Nodes_MacroInfo struct {
    _func *Lns_luaValue
    Symbol2MacroValInfoMap *LnsMap
    FP Nodes_MacroInfoMtd
}
func Nodes_MacroInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_MacroInfo).FP
}
type Nodes_MacroInfoDownCast interface {
    ToNodes_MacroInfo() *Nodes_MacroInfo
}
func Nodes_MacroInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_MacroInfoDownCast)
    if ok { return work.ToNodes_MacroInfo() }
    return nil
}
func (obj *Nodes_MacroInfo) ToNodes_MacroInfo() *Nodes_MacroInfo {
    return obj
}
func (self *Nodes_MacroInfo) InitNodes_MacroInfo(arg1 *Lns_luaValue, arg2 *LnsMap) {
    self._func = arg1
    self.Symbol2MacroValInfoMap = arg2
}




// declaration Class -- RootNode
type Nodes_RootNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_children() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_importModule2moduleInfo() *LnsMap
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_luneHelperInfo() *Nodes_LuneHelperInfo
    Get_macroArgFlag() bool
    Get_moduleId() *FrontInterface_ModuleId
    Get_moduleScope() *Ast_Scope
    Get_moduleTypeInfo() *Ast_TypeInfo
    Get_nodeManager() *Nodes_NodeManager
    Get_pos() *Types_Position
    Get_processInfo() *Ast_ProcessInfo
    Get_provideNode() LnsAny
    Get_tailComment() LnsAny
    Get_typeId2ClassMap() *LnsMap
    Get_typeId2MacroInfo() *LnsMap
    Get_useModuleMacroSet() *LnsSet
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_provide(arg1 *Nodes_ProvideNode)
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_RootNode struct {
    Nodes_Node
    children *LnsList
    moduleScope *Ast_Scope
    useModuleMacroSet *LnsSet
    moduleId *FrontInterface_ModuleId
    processInfo *Ast_ProcessInfo
    moduleTypeInfo *Ast_TypeInfo
    provideNode LnsAny
    luneHelperInfo *Nodes_LuneHelperInfo
    nodeManager *Nodes_NodeManager
    importModule2moduleInfo *LnsMap
    typeId2MacroInfo *LnsMap
    typeId2ClassMap *LnsMap
    FP Nodes_RootNodeMtd
}
func Nodes_RootNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_RootNode).FP
}
type Nodes_RootNodeDownCast interface {
    ToNodes_RootNode() *Nodes_RootNode
}
func Nodes_RootNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_RootNodeDownCast)
    if ok { return work.ToNodes_RootNode() }
    return nil
}
func (obj *Nodes_RootNode) ToNodes_RootNode() *Nodes_RootNode {
    return obj
}
func NewNodes_RootNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList, arg6 *Ast_Scope, arg7 *LnsSet, arg8 *FrontInterface_ModuleId, arg9 *Ast_ProcessInfo, arg10 *Ast_TypeInfo, arg11 LnsAny, arg12 *Nodes_LuneHelperInfo, arg13 *Nodes_NodeManager, arg14 *LnsMap, arg15 *LnsMap, arg16 *LnsMap) *Nodes_RootNode {
    obj := &Nodes_RootNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_RootNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16)
    return obj
}
func (self *Nodes_RootNode) Get_children() *LnsList{ return self.children }
func (self *Nodes_RootNode) Get_moduleScope() *Ast_Scope{ return self.moduleScope }
func (self *Nodes_RootNode) Get_useModuleMacroSet() *LnsSet{ return self.useModuleMacroSet }
func (self *Nodes_RootNode) Get_moduleId() *FrontInterface_ModuleId{ return self.moduleId }
func (self *Nodes_RootNode) Get_processInfo() *Ast_ProcessInfo{ return self.processInfo }
func (self *Nodes_RootNode) Get_moduleTypeInfo() *Ast_TypeInfo{ return self.moduleTypeInfo }
func (self *Nodes_RootNode) Get_provideNode() LnsAny{ return self.provideNode }
func (self *Nodes_RootNode) Get_luneHelperInfo() *Nodes_LuneHelperInfo{ return self.luneHelperInfo }
func (self *Nodes_RootNode) Get_nodeManager() *Nodes_NodeManager{ return self.nodeManager }
func (self *Nodes_RootNode) Get_importModule2moduleInfo() *LnsMap{ return self.importModule2moduleInfo }
func (self *Nodes_RootNode) Get_typeId2MacroInfo() *LnsMap{ return self.typeId2MacroInfo }
func (self *Nodes_RootNode) Get_typeId2ClassMap() *LnsMap{ return self.typeId2ClassMap }
// 1: decl @lune.@base.@Nodes.RootNode.processFilter
func (self *Nodes_RootNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRoot(self, opt)
}

// 1: decl @lune.@base.@Nodes.RootNode.canBeRight
func (self *Nodes_RootNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.RootNode.canBeLeft
func (self *Nodes_RootNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.RootNode.canBeStatement
func (self *Nodes_RootNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_RootNode) InitNodes_RootNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,children *LnsList,moduleScope *Ast_Scope,useModuleMacroSet *LnsSet,moduleId *FrontInterface_ModuleId,processInfo *Ast_ProcessInfo,moduleTypeInfo *Ast_TypeInfo,provideNode LnsAny,luneHelperInfo *Nodes_LuneHelperInfo,nodeManager *Nodes_NodeManager,importModule2moduleInfo *LnsMap,typeId2MacroInfo *LnsMap,typeId2ClassMap *LnsMap) {
    self.InitNodes_Node(id, 5, pos, macroArgFlag, typeList)
    self.children = children
    
    self.moduleScope = moduleScope
    
    self.useModuleMacroSet = useModuleMacroSet
    
    self.moduleId = moduleId
    
    self.processInfo = processInfo
    
    self.moduleTypeInfo = moduleTypeInfo
    
    self.provideNode = provideNode
    
    self.luneHelperInfo = luneHelperInfo
    
    self.nodeManager = nodeManager
    
    self.importModule2moduleInfo = importModule2moduleInfo
    
    self.typeId2MacroInfo = typeId2MacroInfo
    
    self.typeId2ClassMap = typeId2ClassMap
    
}

// 681: decl @lune.@base.@Nodes.RootNode.create
func Nodes_RootNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,children *LnsList,moduleScope *Ast_Scope,useModuleMacroSet *LnsSet,moduleId *FrontInterface_ModuleId,processInfo *Ast_ProcessInfo,moduleTypeInfo *Ast_TypeInfo,provideNode LnsAny,luneHelperInfo *Nodes_LuneHelperInfo,nodeManager *Nodes_NodeManager,importModule2moduleInfo *LnsMap,typeId2MacroInfo *LnsMap,typeId2ClassMap *LnsMap) *Nodes_RootNode {
    var node *Nodes_RootNode
    node = NewNodes_RootNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, children, moduleScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.RootNode.visit
func (self *Nodes_RootNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.children
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch4877 := visitor(child, &self.Nodes_Node, "children", depth); _switch4877 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch4877 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    {
        {
            _child := self.provideNode
            if _child != nil {
                child := _child.(*Nodes_ProvideNode)
                if _switch4936 := visitor(&child.Nodes_Node, &self.Nodes_Node, "provideNode", depth); _switch4936 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch4936 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 874: decl @lune.@base.@Nodes.RootNode.set_provide
func (self *Nodes_RootNode) Set_provide(node *Nodes_ProvideNode) {
    self.provideNode = node
    
}


// declaration Class -- RefTypeNode
type Nodes_RefTypeNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_array() string
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_itemNodeList() *LnsList
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_mutFlag() bool
    Get_name() *Nodes_Node
    Get_pos() *Types_Position
    Get_refFlag() bool
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_RefTypeNode struct {
    Nodes_Node
    name *Nodes_Node
    itemNodeList *LnsList
    refFlag bool
    mutFlag bool
    array string
    FP Nodes_RefTypeNodeMtd
}
func Nodes_RefTypeNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_RefTypeNode).FP
}
type Nodes_RefTypeNodeDownCast interface {
    ToNodes_RefTypeNode() *Nodes_RefTypeNode
}
func Nodes_RefTypeNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_RefTypeNodeDownCast)
    if ok { return work.ToNodes_RefTypeNode() }
    return nil
}
func (obj *Nodes_RefTypeNode) ToNodes_RefTypeNode() *Nodes_RefTypeNode {
    return obj
}
func NewNodes_RefTypeNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 *LnsList, arg7 bool, arg8 bool, arg9 string) *Nodes_RefTypeNode {
    obj := &Nodes_RefTypeNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_RefTypeNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_RefTypeNode) Get_name() *Nodes_Node{ return self.name }
func (self *Nodes_RefTypeNode) Get_itemNodeList() *LnsList{ return self.itemNodeList }
func (self *Nodes_RefTypeNode) Get_refFlag() bool{ return self.refFlag }
func (self *Nodes_RefTypeNode) Get_mutFlag() bool{ return self.mutFlag }
func (self *Nodes_RefTypeNode) Get_array() string{ return self.array }
// 1: decl @lune.@base.@Nodes.RefTypeNode.processFilter
func (self *Nodes_RefTypeNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRefType(self, opt)
}

// 1: decl @lune.@base.@Nodes.RefTypeNode.canBeRight
func (self *Nodes_RefTypeNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.RefTypeNode.canBeLeft
func (self *Nodes_RefTypeNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.RefTypeNode.canBeStatement
func (self *Nodes_RefTypeNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_RefTypeNode) InitNodes_RefTypeNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Nodes_Node,itemNodeList *LnsList,refFlag bool,mutFlag bool,array string) {
    self.InitNodes_Node(id, 6, pos, macroArgFlag, typeList)
    self.name = name
    
    self.itemNodeList = itemNodeList
    
    self.refFlag = refFlag
    
    self.mutFlag = mutFlag
    
    self.array = array
    
}

// 681: decl @lune.@base.@Nodes.RefTypeNode.create
func Nodes_RefTypeNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Nodes_Node,itemNodeList *LnsList,refFlag bool,mutFlag bool,array string) *Nodes_RefTypeNode {
    var node *Nodes_RefTypeNode
    node = NewNodes_RefTypeNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, name, itemNodeList, refFlag, mutFlag, array)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.RefTypeNode.visit
func (self *Nodes_RefTypeNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.name
        if _switch5412 := visitor(child, &self.Nodes_Node, "name", depth); _switch5412 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch5412 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var list *LnsList
        list = self.itemNodeList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch5470 := visitor(child, &self.Nodes_Node, "itemNodeList", depth); _switch5470 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch5470 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}


// declaration Class -- BlockNode
type Nodes_BlockNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_blockKind() LnsInt
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_scope() *Ast_Scope
    Get_stmtList() *LnsList
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_BlockNode struct {
    Nodes_Node
    blockKind LnsInt
    scope *Ast_Scope
    stmtList *LnsList
    FP Nodes_BlockNodeMtd
}
func Nodes_BlockNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_BlockNode).FP
}
type Nodes_BlockNodeDownCast interface {
    ToNodes_BlockNode() *Nodes_BlockNode
}
func Nodes_BlockNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_BlockNodeDownCast)
    if ok { return work.ToNodes_BlockNode() }
    return nil
}
func (obj *Nodes_BlockNode) ToNodes_BlockNode() *Nodes_BlockNode {
    return obj
}
func NewNodes_BlockNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsInt, arg6 *Ast_Scope, arg7 *LnsList) *Nodes_BlockNode {
    obj := &Nodes_BlockNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_BlockNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_BlockNode) Get_blockKind() LnsInt{ return self.blockKind }
func (self *Nodes_BlockNode) Get_scope() *Ast_Scope{ return self.scope }
func (self *Nodes_BlockNode) Get_stmtList() *LnsList{ return self.stmtList }
// 1: decl @lune.@base.@Nodes.BlockNode.processFilter
func (self *Nodes_BlockNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessBlock(self, opt)
}

// 1: decl @lune.@base.@Nodes.BlockNode.canBeRight
func (self *Nodes_BlockNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.BlockNode.canBeLeft
func (self *Nodes_BlockNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.BlockNode.canBeStatement
func (self *Nodes_BlockNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_BlockNode) InitNodes_BlockNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,blockKind LnsInt,scope *Ast_Scope,stmtList *LnsList) {
    self.InitNodes_Node(id, 7, pos, macroArgFlag, typeList)
    self.blockKind = blockKind
    
    self.scope = scope
    
    self.stmtList = stmtList
    
}

// 681: decl @lune.@base.@Nodes.BlockNode.create
func Nodes_BlockNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,blockKind LnsInt,scope *Ast_Scope,stmtList *LnsList) *Nodes_BlockNode {
    var node *Nodes_BlockNode
    node = NewNodes_BlockNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, blockKind, scope, stmtList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.BlockNode.visit
func (self *Nodes_BlockNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.stmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch5855 := visitor(child, &self.Nodes_Node, "stmtList", depth); _switch5855 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch5855 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}

// 987: decl @lune.@base.@Nodes.BlockNode.getBreakKind
func (self *Nodes_BlockNode) GetBreakKind(checkMode LnsInt) LnsInt {
    return Nodes_getBreakKindForStmtList_2517_(checkMode, self.stmtList)
}


// declaration Class -- ScopeNode
type Nodes_ScopeNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_scope() *Ast_Scope
    Get_scopeKind() LnsInt
    Get_symbolList() *LnsList
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ScopeNode struct {
    Nodes_Node
    scopeKind LnsInt
    scope *Ast_Scope
    symbolList *LnsList
    block *Nodes_BlockNode
    FP Nodes_ScopeNodeMtd
}
func Nodes_ScopeNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ScopeNode).FP
}
type Nodes_ScopeNodeDownCast interface {
    ToNodes_ScopeNode() *Nodes_ScopeNode
}
func Nodes_ScopeNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ScopeNodeDownCast)
    if ok { return work.ToNodes_ScopeNode() }
    return nil
}
func (obj *Nodes_ScopeNode) ToNodes_ScopeNode() *Nodes_ScopeNode {
    return obj
}
func NewNodes_ScopeNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsInt, arg6 *Ast_Scope, arg7 *LnsList, arg8 *Nodes_BlockNode) *Nodes_ScopeNode {
    obj := &Nodes_ScopeNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ScopeNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_ScopeNode) Get_scopeKind() LnsInt{ return self.scopeKind }
func (self *Nodes_ScopeNode) Get_scope() *Ast_Scope{ return self.scope }
func (self *Nodes_ScopeNode) Get_symbolList() *LnsList{ return self.symbolList }
func (self *Nodes_ScopeNode) Get_block() *Nodes_BlockNode{ return self.block }
// 1: decl @lune.@base.@Nodes.ScopeNode.processFilter
func (self *Nodes_ScopeNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessScope(self, opt)
}

// 1: decl @lune.@base.@Nodes.ScopeNode.canBeRight
func (self *Nodes_ScopeNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ScopeNode.canBeLeft
func (self *Nodes_ScopeNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ScopeNode.canBeStatement
func (self *Nodes_ScopeNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ScopeNode) InitNodes_ScopeNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,scopeKind LnsInt,scope *Ast_Scope,symbolList *LnsList,block *Nodes_BlockNode) {
    self.InitNodes_Node(id, 8, pos, macroArgFlag, typeList)
    self.scopeKind = scopeKind
    
    self.scope = scope
    
    self.symbolList = symbolList
    
    self.block = block
    
}

// 681: decl @lune.@base.@Nodes.ScopeNode.create
func Nodes_ScopeNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,scopeKind LnsInt,scope *Ast_Scope,symbolList *LnsList,block *Nodes_BlockNode) *Nodes_ScopeNode {
    var node *Nodes_ScopeNode
    node = NewNodes_ScopeNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, scopeKind, scope, symbolList, block)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ScopeNode.visit
func (self *Nodes_ScopeNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch6299 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch6299 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch6299 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- IfStmtInfo
type Nodes_IfStmtInfoMtd interface {
    Get_block() *Nodes_BlockNode
    Get_exp() *Nodes_Node
    Get_kind() LnsInt
}
type Nodes_IfStmtInfo struct {
    kind LnsInt
    exp *Nodes_Node
    block *Nodes_BlockNode
    FP Nodes_IfStmtInfoMtd
}
func Nodes_IfStmtInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_IfStmtInfo).FP
}
type Nodes_IfStmtInfoDownCast interface {
    ToNodes_IfStmtInfo() *Nodes_IfStmtInfo
}
func Nodes_IfStmtInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_IfStmtInfoDownCast)
    if ok { return work.ToNodes_IfStmtInfo() }
    return nil
}
func (obj *Nodes_IfStmtInfo) ToNodes_IfStmtInfo() *Nodes_IfStmtInfo {
    return obj
}
func NewNodes_IfStmtInfo(arg1 LnsInt, arg2 *Nodes_Node, arg3 *Nodes_BlockNode) *Nodes_IfStmtInfo {
    obj := &Nodes_IfStmtInfo{}
    obj.FP = obj
    obj.InitNodes_IfStmtInfo(arg1, arg2, arg3)
    return obj
}
func (self *Nodes_IfStmtInfo) InitNodes_IfStmtInfo(arg1 LnsInt, arg2 *Nodes_Node, arg3 *Nodes_BlockNode) {
    self.kind = arg1
    self.exp = arg2
    self.block = arg3
}
func (self *Nodes_IfStmtInfo) Get_kind() LnsInt{ return self.kind }
func (self *Nodes_IfStmtInfo) Get_exp() *Nodes_Node{ return self.exp }
func (self *Nodes_IfStmtInfo) Get_block() *Nodes_BlockNode{ return self.block }

// declaration Class -- IfNode
type Nodes_IfNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_stmtList() *LnsList
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_IfNode struct {
    Nodes_Node
    stmtList *LnsList
    FP Nodes_IfNodeMtd
}
func Nodes_IfNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_IfNode).FP
}
type Nodes_IfNodeDownCast interface {
    ToNodes_IfNode() *Nodes_IfNode
}
func Nodes_IfNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_IfNodeDownCast)
    if ok { return work.ToNodes_IfNode() }
    return nil
}
func (obj *Nodes_IfNode) ToNodes_IfNode() *Nodes_IfNode {
    return obj
}
func NewNodes_IfNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList) *Nodes_IfNode {
    obj := &Nodes_IfNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_IfNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_IfNode) Get_stmtList() *LnsList{ return self.stmtList }
// 1: decl @lune.@base.@Nodes.IfNode.processFilter
func (self *Nodes_IfNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessIf(self, opt)
}

// 1: decl @lune.@base.@Nodes.IfNode.canBeRight
func (self *Nodes_IfNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.IfNode.canBeLeft
func (self *Nodes_IfNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.IfNode.canBeStatement
func (self *Nodes_IfNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_IfNode) InitNodes_IfNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) {
    self.InitNodes_Node(id, 9, pos, macroArgFlag, typeList)
    self.stmtList = stmtList
    
}

// 681: decl @lune.@base.@Nodes.IfNode.create
func Nodes_IfNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) *Nodes_IfNode {
    var node *Nodes_IfNode
    node = NewNodes_IfNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, stmtList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.IfNode.visit
func (self *Nodes_IfNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 1007: decl @lune.@base.@Nodes.IfNode.getBreakKind
func (self *Nodes_IfNode) GetBreakKind(checkMode LnsInt) LnsInt {
    var hasElseFlag bool
    hasElseFlag = false
    var kind LnsInt
    kind = Nodes_BreakKind__None
    for _, _stmtInfo := range( self.stmtList.Items ) {
        stmtInfo := _stmtInfo.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        var work LnsInt
        work = stmtInfo.FP.Get_block().FP.GetBreakKind(checkMode)
        if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
            if work == Nodes_BreakKind__Return{
                return Nodes_BreakKind__Return
            }
            if work == Nodes_BreakKind__NeverRet{
                return Nodes_BreakKind__NeverRet
            }
        } else { 
            if _switch6975 := work; _switch6975 == Nodes_BreakKind__None {
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                    Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                    if true{
                        return Nodes_BreakKind__None
                    }
                }
            } else {
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                    Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                    kind = work
                    
                }
            }
        }
        
        if stmtInfo.FP.Get_kind() == Nodes_IfKind__Else{
            hasElseFlag = true
            
        }
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( hasElseFlag) ||
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
            Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool))) ).(bool){
        return kind
    }
    return Nodes_BreakKind__None
}


// declaration Class -- MRetExp
type Nodes_MRetExpMtd interface {
    Get_exp() *Nodes_Node
    Get_index() LnsInt
}
type Nodes_MRetExp struct {
    exp *Nodes_Node
    index LnsInt
    FP Nodes_MRetExpMtd
}
func Nodes_MRetExp2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_MRetExp).FP
}
type Nodes_MRetExpDownCast interface {
    ToNodes_MRetExp() *Nodes_MRetExp
}
func Nodes_MRetExpDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_MRetExpDownCast)
    if ok { return work.ToNodes_MRetExp() }
    return nil
}
func (obj *Nodes_MRetExp) ToNodes_MRetExp() *Nodes_MRetExp {
    return obj
}
func NewNodes_MRetExp(arg1 *Nodes_Node, arg2 LnsInt) *Nodes_MRetExp {
    obj := &Nodes_MRetExp{}
    obj.FP = obj
    obj.InitNodes_MRetExp(arg1, arg2)
    return obj
}
func (self *Nodes_MRetExp) InitNodes_MRetExp(arg1 *Nodes_Node, arg2 LnsInt) {
    self.exp = arg1
    self.index = arg2
}
func (self *Nodes_MRetExp) Get_exp() *Nodes_Node{ return self.exp }
func (self *Nodes_MRetExp) Get_index() LnsInt{ return self.index }

// declaration Class -- ExpListNode
type Nodes_ExpListNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetExpTypeAt(arg1 LnsInt) *Ast_TypeInfo
    GetExpTypeNoDDDAt(arg1 LnsInt) *Ast_TypeInfo
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() *LnsList
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_followOn() bool
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_mRetExp() LnsAny
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpListNode struct {
    Nodes_Node
    expList *LnsList
    mRetExp LnsAny
    followOn bool
    FP Nodes_ExpListNodeMtd
}
func Nodes_ExpListNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpListNode).FP
}
type Nodes_ExpListNodeDownCast interface {
    ToNodes_ExpListNode() *Nodes_ExpListNode
}
func Nodes_ExpListNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpListNodeDownCast)
    if ok { return work.ToNodes_ExpListNode() }
    return nil
}
func (obj *Nodes_ExpListNode) ToNodes_ExpListNode() *Nodes_ExpListNode {
    return obj
}
func NewNodes_ExpListNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList, arg6 LnsAny, arg7 bool) *Nodes_ExpListNode {
    obj := &Nodes_ExpListNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpListNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpListNode) Get_expList() *LnsList{ return self.expList }
func (self *Nodes_ExpListNode) Get_mRetExp() LnsAny{ return self.mRetExp }
func (self *Nodes_ExpListNode) Get_followOn() bool{ return self.followOn }
// 1: decl @lune.@base.@Nodes.ExpListNode.processFilter
func (self *Nodes_ExpListNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpList(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpListNode.canBeStatement
func (self *Nodes_ExpListNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpListNode) InitNodes_ExpListNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList *LnsList,mRetExp LnsAny,followOn bool) {
    self.InitNodes_Node(id, 10, pos, macroArgFlag, typeList)
    self.expList = expList
    
    self.mRetExp = mRetExp
    
    self.followOn = followOn
    
}

// 681: decl @lune.@base.@Nodes.ExpListNode.create
func Nodes_ExpListNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList *LnsList,mRetExp LnsAny,followOn bool) *Nodes_ExpListNode {
    var node *Nodes_ExpListNode
    node = NewNodes_ExpListNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, expList, mRetExp, followOn)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpListNode.visit
func (self *Nodes_ExpListNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.expList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch7407 := visitor(child, &self.Nodes_Node, "expList", depth); _switch7407 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch7407 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}

// 1039: decl @lune.@base.@Nodes.ExpListNode.canBeLeft
func (self *Nodes_ExpListNode) CanBeLeft() bool {
    for _, _expNode := range( self.FP.Get_expList().Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(expNode.FP.CanBeLeft()){
            return false
        }
    }
    return true
}

// 1047: decl @lune.@base.@Nodes.ExpListNode.canBeRight
func (self *Nodes_ExpListNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    for _, _expNode := range( self.FP.Get_expList().Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(expNode.FP.CanBeRight(processInfo)){
            return false
        }
    }
    return true
}

// 1055: decl @lune.@base.@Nodes.ExpListNode.setLValue
func (self *Nodes_ExpListNode) SetLValue() {
    for _, _expNode := range( self.FP.Get_expList().Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        expNode.IsLValue = true
        
    }
}

// 1072: decl @lune.@base.@Nodes.ExpListNode.getExpTypeAt
func (self *Nodes_ExpListNode) GetExpTypeAt(index LnsInt) *Ast_TypeInfo {
    if index > self.FP.Get_expTypeList().Len(){
        var lastExpType *Ast_TypeInfo
        lastExpType = self.FP.Get_expTypeList().GetAt(self.FP.Get_expTypeList().Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        {
            _dddTypeInfo := Ast_DDDTypeInfoDownCastF(lastExpType.FP)
            if _dddTypeInfo != nil {
                dddTypeInfo := _dddTypeInfo.(*Ast_DDDTypeInfo)
                return dddTypeInfo.FP.Get_typeInfo().FP.Get_nilableTypeInfo()
            }
        }
        return Ast_builtinTypeNil
    }
    return self.FP.Get_expTypeList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
}

// 1092: decl @lune.@base.@Nodes.ExpListNode.getExpTypeNoDDDAt
func (self *Nodes_ExpListNode) GetExpTypeNoDDDAt(index LnsInt) *Ast_TypeInfo {
    if index >= self.FP.Get_expTypeList().Len(){
        var lastExpType *Ast_TypeInfo
        lastExpType = self.FP.Get_expTypeList().GetAt(self.FP.Get_expTypeList().Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        {
            _dddTypeInfo := Ast_DDDTypeInfoDownCastF(lastExpType.FP)
            if _dddTypeInfo != nil {
                dddTypeInfo := _dddTypeInfo.(*Ast_DDDTypeInfo)
                return dddTypeInfo.FP.Get_typeInfo().FP.Get_nilableTypeInfo()
            }
        }
        if index == self.FP.Get_expTypeList().Len(){
            return lastExpType
        }
        return Ast_builtinTypeNil
    }
    return self.FP.Get_expTypeList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
}


// declaration Class -- CaseInfo
type Nodes_CaseInfoMtd interface {
    Get_block() *Nodes_BlockNode
    Get_expList() *Nodes_ExpListNode
}
type Nodes_CaseInfo struct {
    expList *Nodes_ExpListNode
    block *Nodes_BlockNode
    FP Nodes_CaseInfoMtd
}
func Nodes_CaseInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_CaseInfo).FP
}
type Nodes_CaseInfoDownCast interface {
    ToNodes_CaseInfo() *Nodes_CaseInfo
}
func Nodes_CaseInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_CaseInfoDownCast)
    if ok { return work.ToNodes_CaseInfo() }
    return nil
}
func (obj *Nodes_CaseInfo) ToNodes_CaseInfo() *Nodes_CaseInfo {
    return obj
}
func NewNodes_CaseInfo(arg1 *Nodes_ExpListNode, arg2 *Nodes_BlockNode) *Nodes_CaseInfo {
    obj := &Nodes_CaseInfo{}
    obj.FP = obj
    obj.InitNodes_CaseInfo(arg1, arg2)
    return obj
}
func (self *Nodes_CaseInfo) InitNodes_CaseInfo(arg1 *Nodes_ExpListNode, arg2 *Nodes_BlockNode) {
    self.expList = arg1
    self.block = arg2
}
func (self *Nodes_CaseInfo) Get_expList() *Nodes_ExpListNode{ return self.expList }
func (self *Nodes_CaseInfo) Get_block() *Nodes_BlockNode{ return self.block }

// declaration Class -- SwitchNode
type Nodes_SwitchNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_caseKind() LnsInt
    Get_caseList() *LnsList
    Get_commentList() LnsAny
    Get_default() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_failSafeDefault() bool
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_SwitchNode struct {
    Nodes_Node
    exp *Nodes_Node
    caseList *LnsList
    _default LnsAny
    caseKind LnsInt
    failSafeDefault bool
    FP Nodes_SwitchNodeMtd
}
func Nodes_SwitchNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_SwitchNode).FP
}
type Nodes_SwitchNodeDownCast interface {
    ToNodes_SwitchNode() *Nodes_SwitchNode
}
func Nodes_SwitchNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_SwitchNodeDownCast)
    if ok { return work.ToNodes_SwitchNode() }
    return nil
}
func (obj *Nodes_SwitchNode) ToNodes_SwitchNode() *Nodes_SwitchNode {
    return obj
}
func NewNodes_SwitchNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 *LnsList, arg7 LnsAny, arg8 LnsInt, arg9 bool) *Nodes_SwitchNode {
    obj := &Nodes_SwitchNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_SwitchNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_SwitchNode) Get_exp() *Nodes_Node{ return self.exp }
func (self *Nodes_SwitchNode) Get_caseList() *LnsList{ return self.caseList }
func (self *Nodes_SwitchNode) Get_default() LnsAny{ return self._default }
func (self *Nodes_SwitchNode) Get_caseKind() LnsInt{ return self.caseKind }
func (self *Nodes_SwitchNode) Get_failSafeDefault() bool{ return self.failSafeDefault }
// 1: decl @lune.@base.@Nodes.SwitchNode.processFilter
func (self *Nodes_SwitchNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessSwitch(self, opt)
}

// 1: decl @lune.@base.@Nodes.SwitchNode.canBeRight
func (self *Nodes_SwitchNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.SwitchNode.canBeLeft
func (self *Nodes_SwitchNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.SwitchNode.canBeStatement
func (self *Nodes_SwitchNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_SwitchNode) InitNodes_SwitchNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,caseList *LnsList,_default LnsAny,caseKind LnsInt,failSafeDefault bool) {
    self.InitNodes_Node(id, 11, pos, macroArgFlag, typeList)
    self.exp = exp
    
    self.caseList = caseList
    
    self._default = _default
    
    self.caseKind = caseKind
    
    self.failSafeDefault = failSafeDefault
    
}

// 681: decl @lune.@base.@Nodes.SwitchNode.create
func Nodes_SwitchNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,caseList *LnsList,_default LnsAny,caseKind LnsInt,failSafeDefault bool) *Nodes_SwitchNode {
    var node *Nodes_SwitchNode
    node = NewNodes_SwitchNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp, caseList, _default, caseKind, failSafeDefault)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.SwitchNode.visit
func (self *Nodes_SwitchNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch8084 := visitor(child, &self.Nodes_Node, "exp", depth); _switch8084 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch8084 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self._default
            if _child != nil {
                child := _child.(*Nodes_BlockNode)
                if _switch8140 := visitor(&child.Nodes_Node, &self.Nodes_Node, "default", depth); _switch8140 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch8140 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1129: decl @lune.@base.@Nodes.SwitchNode.getBreakKind
func (self *Nodes_SwitchNode) GetBreakKind(checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = Nodes_BreakKind__None
    var fullCase bool
    fullCase = self.caseKind != Nodes_CaseKind__Lack
    for _, _caseInfo := range( self.caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        var work LnsInt
        work = caseInfo.FP.Get_block().FP.GetBreakKind(checkMode)
        var goNext bool
        goNext = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( (work == Nodes_BreakKind__None)) ||
            Lns_GetEnv().SetStackVal( Lns_op_not(fullCase)) ).(bool)
        if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
            if work == Nodes_BreakKind__Return{
                return Nodes_BreakKind__Return
            }
            if work == Nodes_BreakKind__NeverRet{
                return Nodes_BreakKind__NeverRet
            }
        } else { 
            if _switch8292 := work; _switch8292 == Nodes_BreakKind__None {
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                    Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                    if goNext{
                        return Nodes_BreakKind__None
                    }
                }
            } else {
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                    Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                    kind = work
                    
                }
            }
        }
        
    }
    {
        _block := self._default
        if _block != nil {
            block := _block.(*Nodes_BlockNode)
            var work LnsInt
            work = block.FP.GetBreakKind(checkMode)
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch8404 := work; _switch8404 == Nodes_BreakKind__None {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                        Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                        kind = work
                        
                    }
                }
            }
            
            return kind
        }
    }
    if fullCase{
        return kind
    }
    return Nodes_BreakKind__None
}


// declaration Class -- WhileNode
type Nodes_WhileNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_infinit() bool
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_WhileNode struct {
    Nodes_Node
    exp *Nodes_Node
    infinit bool
    block *Nodes_BlockNode
    FP Nodes_WhileNodeMtd
}
func Nodes_WhileNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_WhileNode).FP
}
type Nodes_WhileNodeDownCast interface {
    ToNodes_WhileNode() *Nodes_WhileNode
}
func Nodes_WhileNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_WhileNodeDownCast)
    if ok { return work.ToNodes_WhileNode() }
    return nil
}
func (obj *Nodes_WhileNode) ToNodes_WhileNode() *Nodes_WhileNode {
    return obj
}
func NewNodes_WhileNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 bool, arg7 *Nodes_BlockNode) *Nodes_WhileNode {
    obj := &Nodes_WhileNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_WhileNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_WhileNode) Get_exp() *Nodes_Node{ return self.exp }
func (self *Nodes_WhileNode) Get_infinit() bool{ return self.infinit }
func (self *Nodes_WhileNode) Get_block() *Nodes_BlockNode{ return self.block }
// 1: decl @lune.@base.@Nodes.WhileNode.processFilter
func (self *Nodes_WhileNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessWhile(self, opt)
}

// 1: decl @lune.@base.@Nodes.WhileNode.canBeRight
func (self *Nodes_WhileNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.WhileNode.canBeLeft
func (self *Nodes_WhileNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.WhileNode.canBeStatement
func (self *Nodes_WhileNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_WhileNode) InitNodes_WhileNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,infinit bool,block *Nodes_BlockNode) {
    self.InitNodes_Node(id, 12, pos, macroArgFlag, typeList)
    self.exp = exp
    
    self.infinit = infinit
    
    self.block = block
    
}

// 681: decl @lune.@base.@Nodes.WhileNode.create
func Nodes_WhileNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,infinit bool,block *Nodes_BlockNode) *Nodes_WhileNode {
    var node *Nodes_WhileNode
    node = NewNodes_WhileNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp, infinit, block)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.WhileNode.visit
func (self *Nodes_WhileNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch8825 := visitor(child, &self.Nodes_Node, "exp", depth); _switch8825 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch8825 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch8882 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch8882 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch8882 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 2154: decl @lune.@base.@Nodes.WhileNode.getBreakKind
func (self *Nodes_WhileNode) GetBreakKind(checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        var kind LnsInt
        kind = Nodes_BreakKind__None
        for _, _stmt := range( self.block.FP.Get_stmtList().Items ) {
            stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
            if stmt.FP.Get_kind() != Nodes_NodeKind_get_BlankLine(){
                var work LnsInt
                work = stmt.FP.GetBreakKind(checkMode)
                if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                    if work == Nodes_BreakKind__Return{
                        return Nodes_BreakKind__Return
                    }
                    if work == Nodes_BreakKind__NeverRet{
                        return Nodes_BreakKind__NeverRet
                    }
                } else { 
                    if _switch41628 := work; _switch41628 == Nodes_BreakKind__None {
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                            Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                            if false{
                                return Nodes_BreakKind__None
                            }
                        }
                    } else {
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                            Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                            kind = work
                            
                        }
                    }
                }
                
            }
        }
        if kind == Nodes_BreakKind__Break{
            return Nodes_BreakKind__None
        }
        return kind
    } else { 
        if Lns_op_not(self.infinit){
            return Nodes_BreakKind__None
        }
        var mode LnsInt
        mode = Nodes_CheckBreakMode__IgnoreFlow
        var kind LnsInt
        kind = Nodes_BreakKind__None
        for _, _stmt := range( self.block.FP.Get_stmtList().Items ) {
            stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
            if stmt.FP.Get_kind() != Nodes_NodeKind_get_BlankLine(){
                var work LnsInt
                work = stmt.FP.GetBreakKind(mode)
                if mode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                    if work == Nodes_BreakKind__Return{
                        return Nodes_BreakKind__Return
                    }
                    if work == Nodes_BreakKind__NeverRet{
                        return Nodes_BreakKind__NeverRet
                    }
                } else { 
                    if _switch41787 := work; _switch41787 == Nodes_BreakKind__None {
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( mode == Nodes_CheckBreakMode__Normal) ||
                            Lns_GetEnv().SetStackVal( mode == Nodes_CheckBreakMode__Return) ).(bool){
                            if false{
                                return Nodes_BreakKind__None
                            }
                        }
                    } else {
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                            Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                            kind = work
                            
                        }
                    }
                }
                
            }
        }
        if kind == Nodes_BreakKind__Break{
            return Nodes_BreakKind__None
        }
        if kind == Nodes_BreakKind__Return{
            return Nodes_BreakKind__Return
        }
        return Nodes_BreakKind__NeverRet
    }
// insert a dummy
    return 0
}


// declaration Class -- RepeatNode
type Nodes_RepeatNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_RepeatNode struct {
    Nodes_Node
    block *Nodes_BlockNode
    exp *Nodes_Node
    FP Nodes_RepeatNodeMtd
}
func Nodes_RepeatNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_RepeatNode).FP
}
type Nodes_RepeatNodeDownCast interface {
    ToNodes_RepeatNode() *Nodes_RepeatNode
}
func Nodes_RepeatNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_RepeatNodeDownCast)
    if ok { return work.ToNodes_RepeatNode() }
    return nil
}
func (obj *Nodes_RepeatNode) ToNodes_RepeatNode() *Nodes_RepeatNode {
    return obj
}
func NewNodes_RepeatNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_BlockNode, arg6 *Nodes_Node) *Nodes_RepeatNode {
    obj := &Nodes_RepeatNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_RepeatNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Nodes_RepeatNode) Get_block() *Nodes_BlockNode{ return self.block }
func (self *Nodes_RepeatNode) Get_exp() *Nodes_Node{ return self.exp }
// 1: decl @lune.@base.@Nodes.RepeatNode.processFilter
func (self *Nodes_RepeatNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRepeat(self, opt)
}

// 1: decl @lune.@base.@Nodes.RepeatNode.canBeRight
func (self *Nodes_RepeatNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.RepeatNode.canBeLeft
func (self *Nodes_RepeatNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.RepeatNode.canBeStatement
func (self *Nodes_RepeatNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_RepeatNode) InitNodes_RepeatNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,block *Nodes_BlockNode,exp *Nodes_Node) {
    self.InitNodes_Node(id, 13, pos, macroArgFlag, typeList)
    self.block = block
    
    self.exp = exp
    
}

// 681: decl @lune.@base.@Nodes.RepeatNode.create
func Nodes_RepeatNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,block *Nodes_BlockNode,exp *Nodes_Node) *Nodes_RepeatNode {
    var node *Nodes_RepeatNode
    node = NewNodes_RepeatNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, block, exp)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.RepeatNode.visit
func (self *Nodes_RepeatNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch9255 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch9255 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9255 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_Node
        child = self.exp
        if _switch9311 := visitor(child, &self.Nodes_Node, "exp", depth); _switch9311 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9311 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1153: decl @lune.@base.@Nodes.RepeatNode.getBreakKind
func (self *Nodes_RepeatNode) GetBreakKind(checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(checkMode)
    }
    return Nodes_BreakKind__None
}


// declaration Class -- ForNode
type Nodes_ForNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_delta() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_init() *Nodes_Node
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_to() *Nodes_Node
    Get_val() *Ast_SymbolInfo
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ForNode struct {
    Nodes_Node
    block *Nodes_BlockNode
    val *Ast_SymbolInfo
    init *Nodes_Node
    to *Nodes_Node
    delta LnsAny
    FP Nodes_ForNodeMtd
}
func Nodes_ForNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ForNode).FP
}
type Nodes_ForNodeDownCast interface {
    ToNodes_ForNode() *Nodes_ForNode
}
func Nodes_ForNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ForNodeDownCast)
    if ok { return work.ToNodes_ForNode() }
    return nil
}
func (obj *Nodes_ForNode) ToNodes_ForNode() *Nodes_ForNode {
    return obj
}
func NewNodes_ForNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_BlockNode, arg6 *Ast_SymbolInfo, arg7 *Nodes_Node, arg8 *Nodes_Node, arg9 LnsAny) *Nodes_ForNode {
    obj := &Nodes_ForNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ForNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ForNode) Get_block() *Nodes_BlockNode{ return self.block }
func (self *Nodes_ForNode) Get_val() *Ast_SymbolInfo{ return self.val }
func (self *Nodes_ForNode) Get_init() *Nodes_Node{ return self.init }
func (self *Nodes_ForNode) Get_to() *Nodes_Node{ return self.to }
func (self *Nodes_ForNode) Get_delta() LnsAny{ return self.delta }
// 1: decl @lune.@base.@Nodes.ForNode.processFilter
func (self *Nodes_ForNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessFor(self, opt)
}

// 1: decl @lune.@base.@Nodes.ForNode.canBeRight
func (self *Nodes_ForNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ForNode.canBeLeft
func (self *Nodes_ForNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ForNode.canBeStatement
func (self *Nodes_ForNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ForNode) InitNodes_ForNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,block *Nodes_BlockNode,val *Ast_SymbolInfo,init *Nodes_Node,to *Nodes_Node,delta LnsAny) {
    self.InitNodes_Node(id, 14, pos, macroArgFlag, typeList)
    self.block = block
    
    self.val = val
    
    self.init = init
    
    self.to = to
    
    self.delta = delta
    
}

// 681: decl @lune.@base.@Nodes.ForNode.create
func Nodes_ForNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,block *Nodes_BlockNode,val *Ast_SymbolInfo,init *Nodes_Node,to *Nodes_Node,delta LnsAny) *Nodes_ForNode {
    var node *Nodes_ForNode
    node = NewNodes_ForNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, block, val, init, to, delta)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ForNode.visit
func (self *Nodes_ForNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch9815 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch9815 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9815 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_Node
        child = self.init
        if _switch9871 := visitor(child, &self.Nodes_Node, "init", depth); _switch9871 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9871 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_Node
        child = self.to
        if _switch9927 := visitor(child, &self.Nodes_Node, "to", depth); _switch9927 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch9927 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.delta
            if _child != nil {
                child := _child.(*Nodes_Node)
                if _switch9982 := visitor(child, &self.Nodes_Node, "delta", depth); _switch9982 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch9982 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1153: decl @lune.@base.@Nodes.ForNode.getBreakKind
func (self *Nodes_ForNode) GetBreakKind(checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(checkMode)
    }
    return Nodes_BreakKind__None
}


// declaration Class -- ApplyNode
type Nodes_ApplyNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() *Nodes_ExpListNode
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_varList() *LnsList
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ApplyNode struct {
    Nodes_Node
    varList *LnsList
    expList *Nodes_ExpListNode
    block *Nodes_BlockNode
    FP Nodes_ApplyNodeMtd
}
func Nodes_ApplyNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ApplyNode).FP
}
type Nodes_ApplyNodeDownCast interface {
    ToNodes_ApplyNode() *Nodes_ApplyNode
}
func Nodes_ApplyNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ApplyNodeDownCast)
    if ok { return work.ToNodes_ApplyNode() }
    return nil
}
func (obj *Nodes_ApplyNode) ToNodes_ApplyNode() *Nodes_ApplyNode {
    return obj
}
func NewNodes_ApplyNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList, arg6 *Nodes_ExpListNode, arg7 *Nodes_BlockNode) *Nodes_ApplyNode {
    obj := &Nodes_ApplyNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ApplyNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ApplyNode) Get_varList() *LnsList{ return self.varList }
func (self *Nodes_ApplyNode) Get_expList() *Nodes_ExpListNode{ return self.expList }
func (self *Nodes_ApplyNode) Get_block() *Nodes_BlockNode{ return self.block }
// 1: decl @lune.@base.@Nodes.ApplyNode.processFilter
func (self *Nodes_ApplyNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessApply(self, opt)
}

// 1: decl @lune.@base.@Nodes.ApplyNode.canBeRight
func (self *Nodes_ApplyNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ApplyNode.canBeLeft
func (self *Nodes_ApplyNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ApplyNode.canBeStatement
func (self *Nodes_ApplyNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ApplyNode) InitNodes_ApplyNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,varList *LnsList,expList *Nodes_ExpListNode,block *Nodes_BlockNode) {
    self.InitNodes_Node(id, 15, pos, macroArgFlag, typeList)
    self.varList = varList
    
    self.expList = expList
    
    self.block = block
    
}

// 681: decl @lune.@base.@Nodes.ApplyNode.create
func Nodes_ApplyNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,varList *LnsList,expList *Nodes_ExpListNode,block *Nodes_BlockNode) *Nodes_ApplyNode {
    var node *Nodes_ApplyNode
    node = NewNodes_ApplyNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, varList, expList, block)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ApplyNode.visit
func (self *Nodes_ApplyNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_ExpListNode
        child = self.expList
        if _switch10440 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch10440 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch10440 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch10497 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch10497 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch10497 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1153: decl @lune.@base.@Nodes.ApplyNode.getBreakKind
func (self *Nodes_ApplyNode) GetBreakKind(checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(checkMode)
    }
    return Nodes_BreakKind__None
}


// declaration Class -- ForeachNode
type Nodes_ForeachNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_key() LnsAny
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_threading() bool
    Get_val() *Ast_SymbolInfo
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ForeachNode struct {
    Nodes_Node
    val *Ast_SymbolInfo
    key LnsAny
    exp *Nodes_Node
    threading bool
    block *Nodes_BlockNode
    FP Nodes_ForeachNodeMtd
}
func Nodes_ForeachNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ForeachNode).FP
}
type Nodes_ForeachNodeDownCast interface {
    ToNodes_ForeachNode() *Nodes_ForeachNode
}
func Nodes_ForeachNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ForeachNodeDownCast)
    if ok { return work.ToNodes_ForeachNode() }
    return nil
}
func (obj *Nodes_ForeachNode) ToNodes_ForeachNode() *Nodes_ForeachNode {
    return obj
}
func NewNodes_ForeachNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Ast_SymbolInfo, arg6 LnsAny, arg7 *Nodes_Node, arg8 bool, arg9 *Nodes_BlockNode) *Nodes_ForeachNode {
    obj := &Nodes_ForeachNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ForeachNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ForeachNode) Get_val() *Ast_SymbolInfo{ return self.val }
func (self *Nodes_ForeachNode) Get_key() LnsAny{ return self.key }
func (self *Nodes_ForeachNode) Get_exp() *Nodes_Node{ return self.exp }
func (self *Nodes_ForeachNode) Get_threading() bool{ return self.threading }
func (self *Nodes_ForeachNode) Get_block() *Nodes_BlockNode{ return self.block }
// 1: decl @lune.@base.@Nodes.ForeachNode.isThreading
func (self *Nodes_ForeachNode) IsThreading() bool {
    return self.threading
}

// 1: decl @lune.@base.@Nodes.ForeachNode.processFilter
func (self *Nodes_ForeachNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessForeach(self, opt)
}

// 1: decl @lune.@base.@Nodes.ForeachNode.canBeRight
func (self *Nodes_ForeachNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ForeachNode.canBeLeft
func (self *Nodes_ForeachNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ForeachNode.canBeStatement
func (self *Nodes_ForeachNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ForeachNode) InitNodes_ForeachNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Ast_SymbolInfo,key LnsAny,exp *Nodes_Node,threading bool,block *Nodes_BlockNode) {
    self.InitNodes_Node(id, 16, pos, macroArgFlag, typeList)
    self.val = val
    
    self.key = key
    
    self.exp = exp
    
    self.threading = threading
    
    self.block = block
    
}

// 681: decl @lune.@base.@Nodes.ForeachNode.create
func Nodes_ForeachNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Ast_SymbolInfo,key LnsAny,exp *Nodes_Node,threading bool,block *Nodes_BlockNode) *Nodes_ForeachNode {
    var node *Nodes_ForeachNode
    node = NewNodes_ForeachNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, val, key, exp, threading, block)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ForeachNode.visit
func (self *Nodes_ForeachNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch11013 := visitor(child, &self.Nodes_Node, "exp", depth); _switch11013 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch11013 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch11070 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch11070 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch11070 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1153: decl @lune.@base.@Nodes.ForeachNode.getBreakKind
func (self *Nodes_ForeachNode) GetBreakKind(checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(checkMode)
    }
    return Nodes_BreakKind__None
}


// declaration Class -- ForsortNode
type Nodes_ForsortNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_key() LnsAny
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_sort() bool
    Get_tailComment() LnsAny
    Get_threading() bool
    Get_val() *Ast_SymbolInfo
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ForsortNode struct {
    Nodes_Node
    val *Ast_SymbolInfo
    key LnsAny
    exp *Nodes_Node
    threading bool
    block *Nodes_BlockNode
    sort bool
    FP Nodes_ForsortNodeMtd
}
func Nodes_ForsortNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ForsortNode).FP
}
type Nodes_ForsortNodeDownCast interface {
    ToNodes_ForsortNode() *Nodes_ForsortNode
}
func Nodes_ForsortNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ForsortNodeDownCast)
    if ok { return work.ToNodes_ForsortNode() }
    return nil
}
func (obj *Nodes_ForsortNode) ToNodes_ForsortNode() *Nodes_ForsortNode {
    return obj
}
func NewNodes_ForsortNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Ast_SymbolInfo, arg6 LnsAny, arg7 *Nodes_Node, arg8 bool, arg9 *Nodes_BlockNode, arg10 bool) *Nodes_ForsortNode {
    obj := &Nodes_ForsortNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ForsortNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_ForsortNode) Get_val() *Ast_SymbolInfo{ return self.val }
func (self *Nodes_ForsortNode) Get_key() LnsAny{ return self.key }
func (self *Nodes_ForsortNode) Get_exp() *Nodes_Node{ return self.exp }
func (self *Nodes_ForsortNode) Get_threading() bool{ return self.threading }
func (self *Nodes_ForsortNode) Get_block() *Nodes_BlockNode{ return self.block }
func (self *Nodes_ForsortNode) Get_sort() bool{ return self.sort }
// 1: decl @lune.@base.@Nodes.ForsortNode.isThreading
func (self *Nodes_ForsortNode) IsThreading() bool {
    return self.threading
}

// 1: decl @lune.@base.@Nodes.ForsortNode.processFilter
func (self *Nodes_ForsortNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessForsort(self, opt)
}

// 1: decl @lune.@base.@Nodes.ForsortNode.canBeRight
func (self *Nodes_ForsortNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ForsortNode.canBeLeft
func (self *Nodes_ForsortNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ForsortNode.canBeStatement
func (self *Nodes_ForsortNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ForsortNode) InitNodes_ForsortNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Ast_SymbolInfo,key LnsAny,exp *Nodes_Node,threading bool,block *Nodes_BlockNode,sort bool) {
    self.InitNodes_Node(id, 17, pos, macroArgFlag, typeList)
    self.val = val
    
    self.key = key
    
    self.exp = exp
    
    self.threading = threading
    
    self.block = block
    
    self.sort = sort
    
}

// 681: decl @lune.@base.@Nodes.ForsortNode.create
func Nodes_ForsortNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Ast_SymbolInfo,key LnsAny,exp *Nodes_Node,threading bool,block *Nodes_BlockNode,sort bool) *Nodes_ForsortNode {
    var node *Nodes_ForsortNode
    node = NewNodes_ForsortNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, val, key, exp, threading, block, sort)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ForsortNode.visit
func (self *Nodes_ForsortNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch11613 := visitor(child, &self.Nodes_Node, "exp", depth); _switch11613 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch11613 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch11670 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch11670 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch11670 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1153: decl @lune.@base.@Nodes.ForsortNode.getBreakKind
func (self *Nodes_ForsortNode) GetBreakKind(checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(checkMode)
    }
    return Nodes_BreakKind__None
}


// declaration Class -- ReturnNode
type Nodes_ReturnNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() LnsAny
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ReturnNode struct {
    Nodes_Node
    expList LnsAny
    FP Nodes_ReturnNodeMtd
}
func Nodes_ReturnNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ReturnNode).FP
}
type Nodes_ReturnNodeDownCast interface {
    ToNodes_ReturnNode() *Nodes_ReturnNode
}
func Nodes_ReturnNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ReturnNodeDownCast)
    if ok { return work.ToNodes_ReturnNode() }
    return nil
}
func (obj *Nodes_ReturnNode) ToNodes_ReturnNode() *Nodes_ReturnNode {
    return obj
}
func NewNodes_ReturnNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsAny) *Nodes_ReturnNode {
    obj := &Nodes_ReturnNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ReturnNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ReturnNode) Get_expList() LnsAny{ return self.expList }
// 1: decl @lune.@base.@Nodes.ReturnNode.processFilter
func (self *Nodes_ReturnNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessReturn(self, opt)
}

// 1: decl @lune.@base.@Nodes.ReturnNode.canBeRight
func (self *Nodes_ReturnNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ReturnNode.canBeLeft
func (self *Nodes_ReturnNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ReturnNode.canBeStatement
func (self *Nodes_ReturnNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ReturnNode) InitNodes_ReturnNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList LnsAny) {
    self.InitNodes_Node(id, 18, pos, macroArgFlag, typeList)
    self.expList = expList
    
}

// 681: decl @lune.@base.@Nodes.ReturnNode.create
func Nodes_ReturnNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList LnsAny) *Nodes_ReturnNode {
    var node *Nodes_ReturnNode
    node = NewNodes_ReturnNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ReturnNode.visit
func (self *Nodes_ReturnNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch12062 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch12062 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch12062 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1222: decl @lune.@base.@Nodes.ReturnNode.getBreakKind
func (self *Nodes_ReturnNode) GetBreakKind(checkMode LnsInt) LnsInt {
    return Nodes_BreakKind__Return
}


// declaration Class -- BreakNode
type Nodes_BreakNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_BreakNode struct {
    Nodes_Node
    FP Nodes_BreakNodeMtd
}
func Nodes_BreakNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_BreakNode).FP
}
type Nodes_BreakNodeDownCast interface {
    ToNodes_BreakNode() *Nodes_BreakNode
}
func Nodes_BreakNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_BreakNodeDownCast)
    if ok { return work.ToNodes_BreakNode() }
    return nil
}
func (obj *Nodes_BreakNode) ToNodes_BreakNode() *Nodes_BreakNode {
    return obj
}
func NewNodes_BreakNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList) *Nodes_BreakNode {
    obj := &Nodes_BreakNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_BreakNode(arg1, arg2, arg3, arg4)
    return obj
}
// 1: decl @lune.@base.@Nodes.BreakNode.processFilter
func (self *Nodes_BreakNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessBreak(self, opt)
}

// 1: decl @lune.@base.@Nodes.BreakNode.canBeRight
func (self *Nodes_BreakNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.BreakNode.canBeLeft
func (self *Nodes_BreakNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.BreakNode.canBeStatement
func (self *Nodes_BreakNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_BreakNode) InitNodes_BreakNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(id, 19, pos, macroArgFlag, typeList)
}

// 681: decl @lune.@base.@Nodes.BreakNode.create
func Nodes_BreakNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList) *Nodes_BreakNode {
    var node *Nodes_BreakNode
    node = NewNodes_BreakNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.BreakNode.visit
func (self *Nodes_BreakNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 1228: decl @lune.@base.@Nodes.BreakNode.getBreakKind
func (self *Nodes_BreakNode) GetBreakKind(checkMode LnsInt) LnsInt {
    return Nodes_BreakKind__Break
}


// declaration Class -- ProvideNode
type Nodes_ProvideNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_symbol() *Ast_SymbolInfo
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ProvideNode struct {
    Nodes_Node
    symbol *Ast_SymbolInfo
    FP Nodes_ProvideNodeMtd
}
func Nodes_ProvideNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ProvideNode).FP
}
type Nodes_ProvideNodeDownCast interface {
    ToNodes_ProvideNode() *Nodes_ProvideNode
}
func Nodes_ProvideNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ProvideNodeDownCast)
    if ok { return work.ToNodes_ProvideNode() }
    return nil
}
func (obj *Nodes_ProvideNode) ToNodes_ProvideNode() *Nodes_ProvideNode {
    return obj
}
func NewNodes_ProvideNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Ast_SymbolInfo) *Nodes_ProvideNode {
    obj := &Nodes_ProvideNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ProvideNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ProvideNode) Get_symbol() *Ast_SymbolInfo{ return self.symbol }
// 1: decl @lune.@base.@Nodes.ProvideNode.processFilter
func (self *Nodes_ProvideNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessProvide(self, opt)
}

// 1: decl @lune.@base.@Nodes.ProvideNode.canBeRight
func (self *Nodes_ProvideNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ProvideNode.canBeLeft
func (self *Nodes_ProvideNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ProvideNode.canBeStatement
func (self *Nodes_ProvideNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ProvideNode) InitNodes_ProvideNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,symbol *Ast_SymbolInfo) {
    self.InitNodes_Node(id, 20, pos, macroArgFlag, typeList)
    self.symbol = symbol
    
}

// 681: decl @lune.@base.@Nodes.ProvideNode.create
func Nodes_ProvideNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,symbol *Ast_SymbolInfo) *Nodes_ProvideNode {
    var node *Nodes_ProvideNode
    node = NewNodes_ProvideNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, symbol)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ProvideNode.visit
func (self *Nodes_ProvideNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- ExpNewNode
type Nodes_ExpNewNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_argList() LnsAny
    Get_commentList() LnsAny
    Get_ctorTypeInfo() *Ast_TypeInfo
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_symbol() *Nodes_Node
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpNewNode struct {
    Nodes_Node
    symbol *Nodes_Node
    ctorTypeInfo *Ast_TypeInfo
    argList LnsAny
    FP Nodes_ExpNewNodeMtd
}
func Nodes_ExpNewNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpNewNode).FP
}
type Nodes_ExpNewNodeDownCast interface {
    ToNodes_ExpNewNode() *Nodes_ExpNewNode
}
func Nodes_ExpNewNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpNewNodeDownCast)
    if ok { return work.ToNodes_ExpNewNode() }
    return nil
}
func (obj *Nodes_ExpNewNode) ToNodes_ExpNewNode() *Nodes_ExpNewNode {
    return obj
}
func NewNodes_ExpNewNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 *Ast_TypeInfo, arg7 LnsAny) *Nodes_ExpNewNode {
    obj := &Nodes_ExpNewNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpNewNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpNewNode) Get_symbol() *Nodes_Node{ return self.symbol }
func (self *Nodes_ExpNewNode) Get_ctorTypeInfo() *Ast_TypeInfo{ return self.ctorTypeInfo }
func (self *Nodes_ExpNewNode) Get_argList() LnsAny{ return self.argList }
// 1: decl @lune.@base.@Nodes.ExpNewNode.processFilter
func (self *Nodes_ExpNewNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpNew(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpNewNode.canBeRight
func (self *Nodes_ExpNewNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpNewNode.canBeLeft
func (self *Nodes_ExpNewNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpNewNode.canBeStatement
func (self *Nodes_ExpNewNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ExpNewNode) InitNodes_ExpNewNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,symbol *Nodes_Node,ctorTypeInfo *Ast_TypeInfo,argList LnsAny) {
    self.InitNodes_Node(id, 21, pos, macroArgFlag, typeList)
    self.symbol = symbol
    
    self.ctorTypeInfo = ctorTypeInfo
    
    self.argList = argList
    
}

// 681: decl @lune.@base.@Nodes.ExpNewNode.create
func Nodes_ExpNewNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,symbol *Nodes_Node,ctorTypeInfo *Ast_TypeInfo,argList LnsAny) *Nodes_ExpNewNode {
    var node *Nodes_ExpNewNode
    node = NewNodes_ExpNewNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, symbol, ctorTypeInfo, argList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpNewNode.visit
func (self *Nodes_ExpNewNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.symbol
        if _switch13043 := visitor(child, &self.Nodes_Node, "symbol", depth); _switch13043 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch13043 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.argList
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch13099 := visitor(&child.Nodes_Node, &self.Nodes_Node, "argList", depth); _switch13099 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch13099 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- ExpUnwrapNode
type Nodes_ExpUnwrapNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_default() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpUnwrapNode struct {
    Nodes_Node
    exp *Nodes_Node
    _default LnsAny
    FP Nodes_ExpUnwrapNodeMtd
}
func Nodes_ExpUnwrapNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpUnwrapNode).FP
}
type Nodes_ExpUnwrapNodeDownCast interface {
    ToNodes_ExpUnwrapNode() *Nodes_ExpUnwrapNode
}
func Nodes_ExpUnwrapNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpUnwrapNodeDownCast)
    if ok { return work.ToNodes_ExpUnwrapNode() }
    return nil
}
func (obj *Nodes_ExpUnwrapNode) ToNodes_ExpUnwrapNode() *Nodes_ExpUnwrapNode {
    return obj
}
func NewNodes_ExpUnwrapNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 LnsAny) *Nodes_ExpUnwrapNode {
    obj := &Nodes_ExpUnwrapNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpUnwrapNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Nodes_ExpUnwrapNode) Get_exp() *Nodes_Node{ return self.exp }
func (self *Nodes_ExpUnwrapNode) Get_default() LnsAny{ return self._default }
// 1: decl @lune.@base.@Nodes.ExpUnwrapNode.processFilter
func (self *Nodes_ExpUnwrapNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpUnwrap(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpUnwrapNode.canBeRight
func (self *Nodes_ExpUnwrapNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpUnwrapNode.canBeLeft
func (self *Nodes_ExpUnwrapNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpUnwrapNode.canBeStatement
func (self *Nodes_ExpUnwrapNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpUnwrapNode) InitNodes_ExpUnwrapNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,_default LnsAny) {
    self.InitNodes_Node(id, 22, pos, macroArgFlag, typeList)
    self.exp = exp
    
    self._default = _default
    
}

// 681: decl @lune.@base.@Nodes.ExpUnwrapNode.create
func Nodes_ExpUnwrapNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,_default LnsAny) *Nodes_ExpUnwrapNode {
    var node *Nodes_ExpUnwrapNode
    node = NewNodes_ExpUnwrapNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp, _default)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpUnwrapNode.visit
func (self *Nodes_ExpUnwrapNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch13473 := visitor(child, &self.Nodes_Node, "exp", depth); _switch13473 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch13473 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self._default
            if _child != nil {
                child := _child.(*Nodes_Node)
                if _switch13528 := visitor(child, &self.Nodes_Node, "default", depth); _switch13528 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch13528 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- ExpRefNode
type Nodes_ExpRefNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_symbolInfo() *Ast_SymbolInfo
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpRefNode struct {
    Nodes_Node
    symbolInfo *Ast_SymbolInfo
    FP Nodes_ExpRefNodeMtd
}
func Nodes_ExpRefNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpRefNode).FP
}
type Nodes_ExpRefNodeDownCast interface {
    ToNodes_ExpRefNode() *Nodes_ExpRefNode
}
func Nodes_ExpRefNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpRefNodeDownCast)
    if ok { return work.ToNodes_ExpRefNode() }
    return nil
}
func (obj *Nodes_ExpRefNode) ToNodes_ExpRefNode() *Nodes_ExpRefNode {
    return obj
}
func NewNodes_ExpRefNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Ast_SymbolInfo) *Nodes_ExpRefNode {
    obj := &Nodes_ExpRefNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpRefNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ExpRefNode) Get_symbolInfo() *Ast_SymbolInfo{ return self.symbolInfo }
// 1: decl @lune.@base.@Nodes.ExpRefNode.processFilter
func (self *Nodes_ExpRefNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpRef(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpRefNode.canBeStatement
func (self *Nodes_ExpRefNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpRefNode) InitNodes_ExpRefNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,symbolInfo *Ast_SymbolInfo) {
    self.InitNodes_Node(id, 23, pos, macroArgFlag, typeList)
    self.symbolInfo = symbolInfo
    
}

// 681: decl @lune.@base.@Nodes.ExpRefNode.create
func Nodes_ExpRefNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,symbolInfo *Ast_SymbolInfo) *Nodes_ExpRefNode {
    var node *Nodes_ExpRefNode
    node = NewNodes_ExpRefNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, symbolInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpRefNode.visit
func (self *Nodes_ExpRefNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 1252: decl @lune.@base.@Nodes.ExpRefNode.canBeLeft
func (self *Nodes_ExpRefNode) CanBeLeft() bool {
    return self.FP.Get_symbolInfo().FP.Get_canBeLeft()
}

// 1256: decl @lune.@base.@Nodes.ExpRefNode.canBeRight
func (self *Nodes_ExpRefNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.FP.Get_symbolInfo().FP.Get_canBeRight()) &&
        Lns_GetEnv().SetStackVal( self.FP.Get_symbolInfo().FP.Get_hasValueFlag()) ).(bool)
}

// 2551: decl @lune.@base.@Nodes.ExpRefNode.getLiteral
func (self *Nodes_ExpRefNode) GetLiteral()(LnsAny, LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = self.symbolInfo.FP.Get_typeInfo()
    {
        _enumTypeInfo := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_aliasSrc().FP)
        if _enumTypeInfo != nil {
            enumTypeInfo := _enumTypeInfo.(*Ast_EnumTypeInfo)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mbr) &&
                Lns_GetEnv().SetStackVal( self.symbolInfo.FP.Get_namespaceTypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Enum) ).(bool)){
                var enumval *Ast_EnumValInfo
                enumval = Lns_unwrap( enumTypeInfo.FP.GetEnumValInfo(self.symbolInfo.FP.Get_name())).(*Ast_EnumValInfo)
                return Nodes_enumLiteral2Literal_9640_(enumval.FP.Get_val())
            }
        }
    }
    return nil, Lns_getVM().String_format("unsupport refnode -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)})
}


// declaration Class -- ExpSetValNode
type Nodes_ExpSetValNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_LeftSymList() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp1() *Nodes_Node
    Get_exp2() *Nodes_ExpListNode
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_initSymSet() *LnsSet
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpSetValNode struct {
    Nodes_Node
    exp1 *Nodes_Node
    exp2 *Nodes_ExpListNode
    LeftSymList *LnsList
    initSymSet *LnsSet
    FP Nodes_ExpSetValNodeMtd
}
func Nodes_ExpSetValNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpSetValNode).FP
}
type Nodes_ExpSetValNodeDownCast interface {
    ToNodes_ExpSetValNode() *Nodes_ExpSetValNode
}
func Nodes_ExpSetValNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpSetValNodeDownCast)
    if ok { return work.ToNodes_ExpSetValNode() }
    return nil
}
func (obj *Nodes_ExpSetValNode) ToNodes_ExpSetValNode() *Nodes_ExpSetValNode {
    return obj
}
func NewNodes_ExpSetValNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 *Nodes_ExpListNode, arg7 *LnsList, arg8 *LnsSet) *Nodes_ExpSetValNode {
    obj := &Nodes_ExpSetValNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpSetValNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_ExpSetValNode) Get_exp1() *Nodes_Node{ return self.exp1 }
func (self *Nodes_ExpSetValNode) Get_exp2() *Nodes_ExpListNode{ return self.exp2 }
func (self *Nodes_ExpSetValNode) Get_LeftSymList() *LnsList{ return self.LeftSymList }
func (self *Nodes_ExpSetValNode) Get_initSymSet() *LnsSet{ return self.initSymSet }
// 1: decl @lune.@base.@Nodes.ExpSetValNode.processFilter
func (self *Nodes_ExpSetValNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpSetVal(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpSetValNode.canBeRight
func (self *Nodes_ExpSetValNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpSetValNode.canBeLeft
func (self *Nodes_ExpSetValNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpSetValNode.canBeStatement
func (self *Nodes_ExpSetValNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ExpSetValNode) InitNodes_ExpSetValNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp1 *Nodes_Node,exp2 *Nodes_ExpListNode,LeftSymList *LnsList,initSymSet *LnsSet) {
    self.InitNodes_Node(id, 24, pos, macroArgFlag, typeList)
    self.exp1 = exp1
    
    self.exp2 = exp2
    
    self.LeftSymList = LeftSymList
    
    self.initSymSet = initSymSet
    
}

// 681: decl @lune.@base.@Nodes.ExpSetValNode.create
func Nodes_ExpSetValNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp1 *Nodes_Node,exp2 *Nodes_ExpListNode,LeftSymList *LnsList,initSymSet *LnsSet) *Nodes_ExpSetValNode {
    var node *Nodes_ExpSetValNode
    node = NewNodes_ExpSetValNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp1, exp2, LeftSymList, initSymSet)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpSetValNode.visit
func (self *Nodes_ExpSetValNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp1
        if _switch14277 := visitor(child, &self.Nodes_Node, "exp1", depth); _switch14277 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch14277 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_ExpListNode
        child = self.exp2
        if _switch14334 := visitor(&child.Nodes_Node, &self.Nodes_Node, "exp2", depth); _switch14334 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch14334 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpSetItemNode
type Nodes_ExpSetItemNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp2() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_index() LnsAny
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_val() *Nodes_Node
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpSetItemNode struct {
    Nodes_Node
    val *Nodes_Node
    index LnsAny
    exp2 *Nodes_Node
    FP Nodes_ExpSetItemNodeMtd
}
func Nodes_ExpSetItemNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpSetItemNode).FP
}
type Nodes_ExpSetItemNodeDownCast interface {
    ToNodes_ExpSetItemNode() *Nodes_ExpSetItemNode
}
func Nodes_ExpSetItemNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpSetItemNodeDownCast)
    if ok { return work.ToNodes_ExpSetItemNode() }
    return nil
}
func (obj *Nodes_ExpSetItemNode) ToNodes_ExpSetItemNode() *Nodes_ExpSetItemNode {
    return obj
}
func NewNodes_ExpSetItemNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 LnsAny, arg7 *Nodes_Node) *Nodes_ExpSetItemNode {
    obj := &Nodes_ExpSetItemNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpSetItemNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpSetItemNode) Get_val() *Nodes_Node{ return self.val }
func (self *Nodes_ExpSetItemNode) Get_index() LnsAny{ return self.index }
func (self *Nodes_ExpSetItemNode) Get_exp2() *Nodes_Node{ return self.exp2 }
// 1: decl @lune.@base.@Nodes.ExpSetItemNode.processFilter
func (self *Nodes_ExpSetItemNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpSetItem(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpSetItemNode.canBeRight
func (self *Nodes_ExpSetItemNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpSetItemNode.canBeLeft
func (self *Nodes_ExpSetItemNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpSetItemNode.canBeStatement
func (self *Nodes_ExpSetItemNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ExpSetItemNode) InitNodes_ExpSetItemNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,index LnsAny,exp2 *Nodes_Node) {
    self.InitNodes_Node(id, 25, pos, macroArgFlag, typeList)
    self.val = val
    
    self.index = index
    
    self.exp2 = exp2
    
}

// 681: decl @lune.@base.@Nodes.ExpSetItemNode.create
func Nodes_ExpSetItemNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,index LnsAny,exp2 *Nodes_Node) *Nodes_ExpSetItemNode {
    var node *Nodes_ExpSetItemNode
    node = NewNodes_ExpSetItemNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, val, index, exp2)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpSetItemNode.visit
func (self *Nodes_ExpSetItemNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.val
        if _switch14739 := visitor(child, &self.Nodes_Node, "val", depth); _switch14739 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch14739 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_Node
        child = self.exp2
        if _switch14795 := visitor(child, &self.Nodes_Node, "exp2", depth); _switch14795 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch14795 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpOp2Node
type Nodes_ExpOp2NodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    getValType(arg1 *Nodes_Node)(bool, LnsInt, LnsReal, string, *Ast_TypeInfo)
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp1() *Nodes_Node
    Get_exp2() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_op() *Types_Token
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_threading() bool
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpOp2Node struct {
    Nodes_Node
    op *Types_Token
    threading bool
    exp1 *Nodes_Node
    exp2 *Nodes_Node
    FP Nodes_ExpOp2NodeMtd
}
func Nodes_ExpOp2Node2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpOp2Node).FP
}
type Nodes_ExpOp2NodeDownCast interface {
    ToNodes_ExpOp2Node() *Nodes_ExpOp2Node
}
func Nodes_ExpOp2NodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpOp2NodeDownCast)
    if ok { return work.ToNodes_ExpOp2Node() }
    return nil
}
func (obj *Nodes_ExpOp2Node) ToNodes_ExpOp2Node() *Nodes_ExpOp2Node {
    return obj
}
func NewNodes_ExpOp2Node(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 bool, arg7 *Nodes_Node, arg8 *Nodes_Node) *Nodes_ExpOp2Node {
    obj := &Nodes_ExpOp2Node{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpOp2Node(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_ExpOp2Node) Get_op() *Types_Token{ return self.op }
func (self *Nodes_ExpOp2Node) Get_threading() bool{ return self.threading }
func (self *Nodes_ExpOp2Node) Get_exp1() *Nodes_Node{ return self.exp1 }
func (self *Nodes_ExpOp2Node) Get_exp2() *Nodes_Node{ return self.exp2 }
// 1: decl @lune.@base.@Nodes.ExpOp2Node.isThreading
func (self *Nodes_ExpOp2Node) IsThreading() bool {
    return self.threading
}

// 1: decl @lune.@base.@Nodes.ExpOp2Node.processFilter
func (self *Nodes_ExpOp2Node) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpOp2(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpOp2Node.canBeRight
func (self *Nodes_ExpOp2Node) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpOp2Node.canBeLeft
func (self *Nodes_ExpOp2Node) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpOp2Node.canBeStatement
func (self *Nodes_ExpOp2Node) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpOp2Node) InitNodes_ExpOp2Node(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,op *Types_Token,threading bool,exp1 *Nodes_Node,exp2 *Nodes_Node) {
    self.InitNodes_Node(id, 26, pos, macroArgFlag, typeList)
    self.op = op
    
    self.threading = threading
    
    self.exp1 = exp1
    
    self.exp2 = exp2
    
}

// 681: decl @lune.@base.@Nodes.ExpOp2Node.create
func Nodes_ExpOp2Node_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,op *Types_Token,threading bool,exp1 *Nodes_Node,exp2 *Nodes_Node) *Nodes_ExpOp2Node {
    var node *Nodes_ExpOp2Node
    node = NewNodes_ExpOp2Node(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, op, threading, exp1, exp2)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpOp2Node.visit
func (self *Nodes_ExpOp2Node) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp1
        if _switch15234 := visitor(child, &self.Nodes_Node, "exp1", depth); _switch15234 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch15234 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_Node
        child = self.exp2
        if _switch15290 := visitor(child, &self.Nodes_Node, "exp2", depth); _switch15290 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch15290 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 2589: decl @lune.@base.@Nodes.ExpOp2Node.getValType
func (self *Nodes_ExpOp2Node) getValType(node *Nodes_Node)(bool, LnsInt, LnsReal, string, *Ast_TypeInfo) {
    var literal LnsAny
    
    {
        _literal := Nodes_convExp43627(Lns_2DDD(node.FP.GetLiteral()))
        if _literal == nil{
            return false, 0, 0.0, "", Ast_headTypeInfo
        } else {
            literal = _literal
        }
    }
    var intVal LnsInt
    var realVal LnsReal
    var strVal string
    intVal,realVal,strVal = 0, 0.0, ""
    var retTypeInfo *Ast_TypeInfo
    retTypeInfo = Ast_builtinTypeNone
    switch _exp43726 := literal.(type) {
    case *Nodes_Literal__Int:
    val := _exp43726.Val1
        intVal = val
        
        realVal = (LnsReal)(val)
        
        retTypeInfo = Ast_builtinTypeInt
        
    case *Nodes_Literal__Real:
    val := _exp43726.Val1
        realVal = val
        
        intVal = (LnsInt)(val)
        
        retTypeInfo = Ast_builtinTypeReal
        
    case *Nodes_Literal__Str:
    val := _exp43726.Val1
        strVal = val
        
        retTypeInfo = Ast_builtinTypeString
        
    default:
        return false, 0, 0.0, "", Ast_headTypeInfo
    }
    return true, intVal, realVal, strVal, retTypeInfo
}

// 2639: decl @lune.@base.@Nodes.ExpOp2Node.setupLiteralTokenList
func (self *Nodes_ExpOp2Node) SetupLiteralTokenList(list *LnsList) bool {
    var literal LnsAny
    literal = Nodes_convExp43758(Lns_2DDD(self.FP.GetLiteral()))
    if literal != nil{
        literal_10704 := literal
        switch _exp43826 := literal_10704.(type) {
        case *Nodes_Literal__Int:
        val := _exp43826.Val1
            self.FP.AddTokenList(list, Types_TokenKind__Int, Lns_getVM().String_format("%d", []LnsAny{val}))
        case *Nodes_Literal__Real:
        val := _exp43826.Val1
            self.FP.AddTokenList(list, Types_TokenKind__Real, Lns_getVM().String_format("%g", []LnsAny{val}))
        case *Nodes_Literal__Str:
        val := _exp43826.Val1
            self.FP.AddTokenList(list, Types_TokenKind__Str, Lns_getVM().String_format("%q", []LnsAny{val}))
        default:
            return false
        }
        return true
    } else {
        return false
    }
// insert a dummy
    return false
}

// 2664: decl @lune.@base.@Nodes.ExpOp2Node.getLiteral
func (self *Nodes_ExpOp2Node) GetLiteral()(LnsAny, LnsAny) {
    var ret1 bool
    var int1 LnsInt
    var real1 LnsReal
    var str1 string
    var type1 *Ast_TypeInfo
    ret1,int1,real1,str1,type1 = self.FP.getValType(self.FP.Get_exp1())
    var ret2 bool
    var int2 LnsInt
    var real2 LnsReal
    var str2 string
    var type2 *Ast_TypeInfo
    ret2,int2,real2,str2,type2 = self.FP.getValType(self.FP.Get_exp2())
    if Lns_op_not(ret1){
        return nil, Lns_getVM().String_format("not support literal -- %s", []LnsAny{Nodes_getNodeKindName(self.FP.Get_exp1().FP.Get_kind())})
    }
    if Lns_op_not(ret2){
        return nil, Lns_getVM().String_format("not support literal -- %s", []LnsAny{Nodes_getNodeKindName(self.FP.Get_exp2().FP.Get_kind())})
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( type1 == Ast_builtinTypeInt) ||
            Lns_GetEnv().SetStackVal( type1 == Ast_builtinTypeReal) ).(bool))) &&
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( type2 == Ast_builtinTypeInt) ||
            Lns_GetEnv().SetStackVal( type2 == Ast_builtinTypeReal) ).(bool))) ).(bool)){
        var retType *Ast_TypeInfo
        retType = Ast_builtinTypeInt
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( type1 == Ast_builtinTypeReal) ||
            Lns_GetEnv().SetStackVal( type2 == Ast_builtinTypeReal) ).(bool){
            retType = Ast_builtinTypeReal
            
        }
        if _switch44135 := (self.op.Txt); _switch44135 == "+" {
            if retType == Ast_builtinTypeInt{
                return &Nodes_Literal__Int{int1 + int2}, nil
            }
            return &Nodes_Literal__Real{real1 + real2}, nil
        } else if _switch44135 == "-" {
            if retType == Ast_builtinTypeInt{
                return &Nodes_Literal__Int{int1 - int2}, nil
            }
            return &Nodes_Literal__Real{real1 - real2}, nil
        } else if _switch44135 == "*" {
            if retType == Ast_builtinTypeInt{
                return &Nodes_Literal__Int{int1 * int2}, nil
            }
            return &Nodes_Literal__Real{real1 * real2}, nil
        } else if _switch44135 == "/" {
            if retType == Ast_builtinTypeInt{
                return &Nodes_Literal__Int{int1 / int2}, nil
            }
            return &Nodes_Literal__Real{real1 / real2}, nil
        }
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( type1 == Ast_builtinTypeString) &&
        Lns_GetEnv().SetStackVal( type2 == Ast_builtinTypeString) ).(bool)){
        if self.op.Txt == ".."{
            return &Nodes_Literal__Str{str1 + str2}, nil
        }
    }
    var mess string
    mess = Lns_getVM().String_format("not support literal operation -- %s %s %s", []LnsAny{type1.FP.GetTxt(nil, nil, nil), self.op.Txt, type2.FP.GetTxt(nil, nil, nil)})
    return nil, mess
}


// declaration Class -- UnwrapSetNode
type Nodes_UnwrapSetNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_dstExpList() *Nodes_ExpListNode
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_srcExpList() *Nodes_ExpListNode
    Get_tailComment() LnsAny
    Get_unwrapBlock() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_UnwrapSetNode struct {
    Nodes_Node
    dstExpList *Nodes_ExpListNode
    srcExpList *Nodes_ExpListNode
    unwrapBlock LnsAny
    FP Nodes_UnwrapSetNodeMtd
}
func Nodes_UnwrapSetNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_UnwrapSetNode).FP
}
type Nodes_UnwrapSetNodeDownCast interface {
    ToNodes_UnwrapSetNode() *Nodes_UnwrapSetNode
}
func Nodes_UnwrapSetNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_UnwrapSetNodeDownCast)
    if ok { return work.ToNodes_UnwrapSetNode() }
    return nil
}
func (obj *Nodes_UnwrapSetNode) ToNodes_UnwrapSetNode() *Nodes_UnwrapSetNode {
    return obj
}
func NewNodes_UnwrapSetNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_ExpListNode, arg6 *Nodes_ExpListNode, arg7 LnsAny) *Nodes_UnwrapSetNode {
    obj := &Nodes_UnwrapSetNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_UnwrapSetNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_UnwrapSetNode) Get_dstExpList() *Nodes_ExpListNode{ return self.dstExpList }
func (self *Nodes_UnwrapSetNode) Get_srcExpList() *Nodes_ExpListNode{ return self.srcExpList }
func (self *Nodes_UnwrapSetNode) Get_unwrapBlock() LnsAny{ return self.unwrapBlock }
// 1: decl @lune.@base.@Nodes.UnwrapSetNode.processFilter
func (self *Nodes_UnwrapSetNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessUnwrapSet(self, opt)
}

// 1: decl @lune.@base.@Nodes.UnwrapSetNode.canBeRight
func (self *Nodes_UnwrapSetNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.UnwrapSetNode.canBeLeft
func (self *Nodes_UnwrapSetNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.UnwrapSetNode.canBeStatement
func (self *Nodes_UnwrapSetNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_UnwrapSetNode) InitNodes_UnwrapSetNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,dstExpList *Nodes_ExpListNode,srcExpList *Nodes_ExpListNode,unwrapBlock LnsAny) {
    self.InitNodes_Node(id, 27, pos, macroArgFlag, typeList)
    self.dstExpList = dstExpList
    
    self.srcExpList = srcExpList
    
    self.unwrapBlock = unwrapBlock
    
}

// 681: decl @lune.@base.@Nodes.UnwrapSetNode.create
func Nodes_UnwrapSetNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,dstExpList *Nodes_ExpListNode,srcExpList *Nodes_ExpListNode,unwrapBlock LnsAny) *Nodes_UnwrapSetNode {
    var node *Nodes_UnwrapSetNode
    node = NewNodes_UnwrapSetNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, dstExpList, srcExpList, unwrapBlock)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.UnwrapSetNode.visit
func (self *Nodes_UnwrapSetNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_ExpListNode
        child = self.dstExpList
        if _switch15690 := visitor(&child.Nodes_Node, &self.Nodes_Node, "dstExpList", depth); _switch15690 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch15690 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_ExpListNode
        child = self.srcExpList
        if _switch15747 := visitor(&child.Nodes_Node, &self.Nodes_Node, "srcExpList", depth); _switch15747 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch15747 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.unwrapBlock
            if _child != nil {
                child := _child.(*Nodes_BlockNode)
                if _switch15803 := visitor(&child.Nodes_Node, &self.Nodes_Node, "unwrapBlock", depth); _switch15803 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch15803 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- IfUnwrapNode
type Nodes_IfUnwrapNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() *Nodes_ExpListNode
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_nilBlock() LnsAny
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_varSymList() *LnsList
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_IfUnwrapNode struct {
    Nodes_Node
    varSymList *LnsList
    expList *Nodes_ExpListNode
    block *Nodes_BlockNode
    nilBlock LnsAny
    FP Nodes_IfUnwrapNodeMtd
}
func Nodes_IfUnwrapNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_IfUnwrapNode).FP
}
type Nodes_IfUnwrapNodeDownCast interface {
    ToNodes_IfUnwrapNode() *Nodes_IfUnwrapNode
}
func Nodes_IfUnwrapNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_IfUnwrapNodeDownCast)
    if ok { return work.ToNodes_IfUnwrapNode() }
    return nil
}
func (obj *Nodes_IfUnwrapNode) ToNodes_IfUnwrapNode() *Nodes_IfUnwrapNode {
    return obj
}
func NewNodes_IfUnwrapNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList, arg6 *Nodes_ExpListNode, arg7 *Nodes_BlockNode, arg8 LnsAny) *Nodes_IfUnwrapNode {
    obj := &Nodes_IfUnwrapNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_IfUnwrapNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_IfUnwrapNode) Get_varSymList() *LnsList{ return self.varSymList }
func (self *Nodes_IfUnwrapNode) Get_expList() *Nodes_ExpListNode{ return self.expList }
func (self *Nodes_IfUnwrapNode) Get_block() *Nodes_BlockNode{ return self.block }
func (self *Nodes_IfUnwrapNode) Get_nilBlock() LnsAny{ return self.nilBlock }
// 1: decl @lune.@base.@Nodes.IfUnwrapNode.processFilter
func (self *Nodes_IfUnwrapNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessIfUnwrap(self, opt)
}

// 1: decl @lune.@base.@Nodes.IfUnwrapNode.canBeRight
func (self *Nodes_IfUnwrapNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.IfUnwrapNode.canBeLeft
func (self *Nodes_IfUnwrapNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.IfUnwrapNode.canBeStatement
func (self *Nodes_IfUnwrapNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_IfUnwrapNode) InitNodes_IfUnwrapNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,varSymList *LnsList,expList *Nodes_ExpListNode,block *Nodes_BlockNode,nilBlock LnsAny) {
    self.InitNodes_Node(id, 28, pos, macroArgFlag, typeList)
    self.varSymList = varSymList
    
    self.expList = expList
    
    self.block = block
    
    self.nilBlock = nilBlock
    
}

// 681: decl @lune.@base.@Nodes.IfUnwrapNode.create
func Nodes_IfUnwrapNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,varSymList *LnsList,expList *Nodes_ExpListNode,block *Nodes_BlockNode,nilBlock LnsAny) *Nodes_IfUnwrapNode {
    var node *Nodes_IfUnwrapNode
    node = NewNodes_IfUnwrapNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, varSymList, expList, block, nilBlock)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.IfUnwrapNode.visit
func (self *Nodes_IfUnwrapNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_ExpListNode
        child = self.expList
        if _switch16241 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch16241 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch16241 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch16298 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch16298 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch16298 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.nilBlock
            if _child != nil {
                child := _child.(*Nodes_BlockNode)
                if _switch16354 := visitor(&child.Nodes_Node, &self.Nodes_Node, "nilBlock", depth); _switch16354 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch16354 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1311: decl @lune.@base.@Nodes.IfUnwrapNode.getBreakKind
func (self *Nodes_IfUnwrapNode) GetBreakKind(checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = self.block.FP.GetBreakKind(checkMode)
    var work LnsInt
    work = kind
    if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
        if work == Nodes_BreakKind__Return{
            return Nodes_BreakKind__Return
        }
        if work == Nodes_BreakKind__NeverRet{
            return Nodes_BreakKind__NeverRet
        }
    } else { 
        if _switch16481 := work; _switch16481 == Nodes_BreakKind__None {
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                if true{
                    return Nodes_BreakKind__None
                }
            }
        } else {
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                kind = work
                
            }
        }
    }
    
    {
        _block := self.nilBlock
        if _block != nil {
            block := _block.(*Nodes_BlockNode)
            work = block.FP.GetBreakKind(checkMode)
            
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch16594 := work; _switch16594 == Nodes_BreakKind__None {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                        Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                        kind = work
                        
                    }
                }
            }
            
            return kind
        }
    }
    return Nodes_BreakKind__None
}


// declaration Class -- UnwrapSymbolPair
type Nodes_UnwrapSymbolPairMtd interface {
    Get_dst() *Ast_SymbolInfo
    Get_src() *Ast_SymbolInfo
}
type Nodes_UnwrapSymbolPair struct {
    src *Ast_SymbolInfo
    dst *Ast_SymbolInfo
    FP Nodes_UnwrapSymbolPairMtd
}
func Nodes_UnwrapSymbolPair2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_UnwrapSymbolPair).FP
}
type Nodes_UnwrapSymbolPairDownCast interface {
    ToNodes_UnwrapSymbolPair() *Nodes_UnwrapSymbolPair
}
func Nodes_UnwrapSymbolPairDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_UnwrapSymbolPairDownCast)
    if ok { return work.ToNodes_UnwrapSymbolPair() }
    return nil
}
func (obj *Nodes_UnwrapSymbolPair) ToNodes_UnwrapSymbolPair() *Nodes_UnwrapSymbolPair {
    return obj
}
func NewNodes_UnwrapSymbolPair(arg1 *Ast_SymbolInfo, arg2 *Ast_SymbolInfo) *Nodes_UnwrapSymbolPair {
    obj := &Nodes_UnwrapSymbolPair{}
    obj.FP = obj
    obj.InitNodes_UnwrapSymbolPair(arg1, arg2)
    return obj
}
func (self *Nodes_UnwrapSymbolPair) InitNodes_UnwrapSymbolPair(arg1 *Ast_SymbolInfo, arg2 *Ast_SymbolInfo) {
    self.src = arg1
    self.dst = arg2
}
func (self *Nodes_UnwrapSymbolPair) Get_src() *Ast_SymbolInfo{ return self.src }
func (self *Nodes_UnwrapSymbolPair) Get_dst() *Ast_SymbolInfo{ return self.dst }

// declaration Class -- WhenNode
type Nodes_WhenNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_elseBlock() LnsAny
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_symPairList() *LnsList
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_WhenNode struct {
    Nodes_Node
    symPairList *LnsList
    block *Nodes_BlockNode
    elseBlock LnsAny
    FP Nodes_WhenNodeMtd
}
func Nodes_WhenNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_WhenNode).FP
}
type Nodes_WhenNodeDownCast interface {
    ToNodes_WhenNode() *Nodes_WhenNode
}
func Nodes_WhenNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_WhenNodeDownCast)
    if ok { return work.ToNodes_WhenNode() }
    return nil
}
func (obj *Nodes_WhenNode) ToNodes_WhenNode() *Nodes_WhenNode {
    return obj
}
func NewNodes_WhenNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList, arg6 *Nodes_BlockNode, arg7 LnsAny) *Nodes_WhenNode {
    obj := &Nodes_WhenNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_WhenNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_WhenNode) Get_symPairList() *LnsList{ return self.symPairList }
func (self *Nodes_WhenNode) Get_block() *Nodes_BlockNode{ return self.block }
func (self *Nodes_WhenNode) Get_elseBlock() LnsAny{ return self.elseBlock }
// 1: decl @lune.@base.@Nodes.WhenNode.processFilter
func (self *Nodes_WhenNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessWhen(self, opt)
}

// 1: decl @lune.@base.@Nodes.WhenNode.canBeRight
func (self *Nodes_WhenNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.WhenNode.canBeLeft
func (self *Nodes_WhenNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.WhenNode.canBeStatement
func (self *Nodes_WhenNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_WhenNode) InitNodes_WhenNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,symPairList *LnsList,block *Nodes_BlockNode,elseBlock LnsAny) {
    self.InitNodes_Node(id, 29, pos, macroArgFlag, typeList)
    self.symPairList = symPairList
    
    self.block = block
    
    self.elseBlock = elseBlock
    
}

// 681: decl @lune.@base.@Nodes.WhenNode.create
func Nodes_WhenNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,symPairList *LnsList,block *Nodes_BlockNode,elseBlock LnsAny) *Nodes_WhenNode {
    var node *Nodes_WhenNode
    node = NewNodes_WhenNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, symPairList, block, elseBlock)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.WhenNode.visit
func (self *Nodes_WhenNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch17015 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch17015 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch17015 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.elseBlock
            if _child != nil {
                child := _child.(*Nodes_BlockNode)
                if _switch17071 := visitor(&child.Nodes_Node, &self.Nodes_Node, "elseBlock", depth); _switch17071 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch17071 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1337: decl @lune.@base.@Nodes.WhenNode.getBreakKind
func (self *Nodes_WhenNode) GetBreakKind(checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = self.block.FP.GetBreakKind(checkMode)
    var work LnsInt
    work = kind
    if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
        if work == Nodes_BreakKind__Return{
            return Nodes_BreakKind__Return
        }
        if work == Nodes_BreakKind__NeverRet{
            return Nodes_BreakKind__NeverRet
        }
    } else { 
        if _switch17198 := work; _switch17198 == Nodes_BreakKind__None {
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                if true{
                    return Nodes_BreakKind__None
                }
            }
        } else {
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                kind = work
                
            }
        }
    }
    
    {
        _block := self.elseBlock
        if _block != nil {
            block := _block.(*Nodes_BlockNode)
            work = block.FP.GetBreakKind(checkMode)
            
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch17311 := work; _switch17311 == Nodes_BreakKind__None {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                        Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                        kind = work
                        
                    }
                }
            }
            
            return kind
        }
    }
    return Nodes_BreakKind__None
}


// declaration Class -- ExpCastNode
type Nodes_ExpCastNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_castKind() LnsInt
    Get_castType() *Ast_TypeInfo
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpCastNode struct {
    Nodes_Node
    exp *Nodes_Node
    castType *Ast_TypeInfo
    castKind LnsInt
    FP Nodes_ExpCastNodeMtd
}
func Nodes_ExpCastNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpCastNode).FP
}
type Nodes_ExpCastNodeDownCast interface {
    ToNodes_ExpCastNode() *Nodes_ExpCastNode
}
func Nodes_ExpCastNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpCastNodeDownCast)
    if ok { return work.ToNodes_ExpCastNode() }
    return nil
}
func (obj *Nodes_ExpCastNode) ToNodes_ExpCastNode() *Nodes_ExpCastNode {
    return obj
}
func NewNodes_ExpCastNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 *Ast_TypeInfo, arg7 LnsInt) *Nodes_ExpCastNode {
    obj := &Nodes_ExpCastNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpCastNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpCastNode) Get_exp() *Nodes_Node{ return self.exp }
func (self *Nodes_ExpCastNode) Get_castType() *Ast_TypeInfo{ return self.castType }
func (self *Nodes_ExpCastNode) Get_castKind() LnsInt{ return self.castKind }
// 1: decl @lune.@base.@Nodes.ExpCastNode.processFilter
func (self *Nodes_ExpCastNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCast(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpCastNode.canBeRight
func (self *Nodes_ExpCastNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpCastNode.canBeLeft
func (self *Nodes_ExpCastNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpCastNode.canBeStatement
func (self *Nodes_ExpCastNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpCastNode) InitNodes_ExpCastNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,castType *Ast_TypeInfo,castKind LnsInt) {
    self.InitNodes_Node(id, 30, pos, macroArgFlag, typeList)
    self.exp = exp
    
    self.castType = castType
    
    self.castKind = castKind
    
}

// 681: decl @lune.@base.@Nodes.ExpCastNode.create
func Nodes_ExpCastNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,castType *Ast_TypeInfo,castKind LnsInt) *Nodes_ExpCastNode {
    var node *Nodes_ExpCastNode
    node = NewNodes_ExpCastNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp, castType, castKind)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpCastNode.visit
func (self *Nodes_ExpCastNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch17720 := visitor(child, &self.Nodes_Node, "exp", depth); _switch17720 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch17720 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1372: decl @lune.@base.@Nodes.ExpCastNode.getPrefix
func (self *Nodes_ExpCastNode) GetPrefix() LnsAny {
    return self.exp.FP.GetPrefix()
}

// 1375: decl @lune.@base.@Nodes.ExpCastNode.getLiteral
func (self *Nodes_ExpCastNode) GetLiteral()(LnsAny, LnsAny) {
    return self.exp.FP.GetLiteral()
}

// 1378: decl @lune.@base.@Nodes.ExpCastNode.setupLiteralTokenList
func (self *Nodes_ExpCastNode) SetupLiteralTokenList(list *LnsList) bool {
    return self.exp.FP.SetupLiteralTokenList(list)
}


// declaration Class -- ExpToDDDNode
type Nodes_ExpToDDDNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() *Nodes_ExpListNode
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpToDDDNode struct {
    Nodes_Node
    expList *Nodes_ExpListNode
    FP Nodes_ExpToDDDNodeMtd
}
func Nodes_ExpToDDDNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpToDDDNode).FP
}
type Nodes_ExpToDDDNodeDownCast interface {
    ToNodes_ExpToDDDNode() *Nodes_ExpToDDDNode
}
func Nodes_ExpToDDDNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpToDDDNodeDownCast)
    if ok { return work.ToNodes_ExpToDDDNode() }
    return nil
}
func (obj *Nodes_ExpToDDDNode) ToNodes_ExpToDDDNode() *Nodes_ExpToDDDNode {
    return obj
}
func NewNodes_ExpToDDDNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_ExpListNode) *Nodes_ExpToDDDNode {
    obj := &Nodes_ExpToDDDNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpToDDDNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ExpToDDDNode) Get_expList() *Nodes_ExpListNode{ return self.expList }
// 1: decl @lune.@base.@Nodes.ExpToDDDNode.processFilter
func (self *Nodes_ExpToDDDNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpToDDD(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpToDDDNode.canBeRight
func (self *Nodes_ExpToDDDNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpToDDDNode.canBeLeft
func (self *Nodes_ExpToDDDNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpToDDDNode.canBeStatement
func (self *Nodes_ExpToDDDNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpToDDDNode) InitNodes_ExpToDDDNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList *Nodes_ExpListNode) {
    self.InitNodes_Node(id, 31, pos, macroArgFlag, typeList)
    self.expList = expList
    
}

// 681: decl @lune.@base.@Nodes.ExpToDDDNode.create
func Nodes_ExpToDDDNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList *Nodes_ExpListNode) *Nodes_ExpToDDDNode {
    var node *Nodes_ExpToDDDNode
    node = NewNodes_ExpToDDDNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpToDDDNode.visit
func (self *Nodes_ExpToDDDNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_ExpListNode
        child = self.expList
        if _switch18116 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch18116 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch18116 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpSubDDDNode
type Nodes_ExpSubDDDNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_remainIndex() LnsInt
    Get_src() *Nodes_Node
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpSubDDDNode struct {
    Nodes_Node
    src *Nodes_Node
    remainIndex LnsInt
    FP Nodes_ExpSubDDDNodeMtd
}
func Nodes_ExpSubDDDNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpSubDDDNode).FP
}
type Nodes_ExpSubDDDNodeDownCast interface {
    ToNodes_ExpSubDDDNode() *Nodes_ExpSubDDDNode
}
func Nodes_ExpSubDDDNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpSubDDDNodeDownCast)
    if ok { return work.ToNodes_ExpSubDDDNode() }
    return nil
}
func (obj *Nodes_ExpSubDDDNode) ToNodes_ExpSubDDDNode() *Nodes_ExpSubDDDNode {
    return obj
}
func NewNodes_ExpSubDDDNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 LnsInt) *Nodes_ExpSubDDDNode {
    obj := &Nodes_ExpSubDDDNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpSubDDDNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Nodes_ExpSubDDDNode) Get_src() *Nodes_Node{ return self.src }
func (self *Nodes_ExpSubDDDNode) Get_remainIndex() LnsInt{ return self.remainIndex }
// 1: decl @lune.@base.@Nodes.ExpSubDDDNode.processFilter
func (self *Nodes_ExpSubDDDNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpSubDDD(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpSubDDDNode.canBeRight
func (self *Nodes_ExpSubDDDNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpSubDDDNode.canBeLeft
func (self *Nodes_ExpSubDDDNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpSubDDDNode.canBeStatement
func (self *Nodes_ExpSubDDDNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpSubDDDNode) InitNodes_ExpSubDDDNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,src *Nodes_Node,remainIndex LnsInt) {
    self.InitNodes_Node(id, 32, pos, macroArgFlag, typeList)
    self.src = src
    
    self.remainIndex = remainIndex
    
}

// 681: decl @lune.@base.@Nodes.ExpSubDDDNode.create
func Nodes_ExpSubDDDNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,src *Nodes_Node,remainIndex LnsInt) *Nodes_ExpSubDDDNode {
    var node *Nodes_ExpSubDDDNode
    node = NewNodes_ExpSubDDDNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, src, remainIndex)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpSubDDDNode.visit
func (self *Nodes_ExpSubDDDNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.src
        if _switch18488 := visitor(child, &self.Nodes_Node, "src", depth); _switch18488 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch18488 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpOp1Node
type Nodes_ExpOp1NodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_macroMode() LnsInt
    Get_op() *Types_Token
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpOp1Node struct {
    Nodes_Node
    op *Types_Token
    macroMode LnsInt
    exp *Nodes_Node
    FP Nodes_ExpOp1NodeMtd
}
func Nodes_ExpOp1Node2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpOp1Node).FP
}
type Nodes_ExpOp1NodeDownCast interface {
    ToNodes_ExpOp1Node() *Nodes_ExpOp1Node
}
func Nodes_ExpOp1NodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpOp1NodeDownCast)
    if ok { return work.ToNodes_ExpOp1Node() }
    return nil
}
func (obj *Nodes_ExpOp1Node) ToNodes_ExpOp1Node() *Nodes_ExpOp1Node {
    return obj
}
func NewNodes_ExpOp1Node(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 LnsInt, arg7 *Nodes_Node) *Nodes_ExpOp1Node {
    obj := &Nodes_ExpOp1Node{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpOp1Node(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpOp1Node) Get_op() *Types_Token{ return self.op }
func (self *Nodes_ExpOp1Node) Get_macroMode() LnsInt{ return self.macroMode }
func (self *Nodes_ExpOp1Node) Get_exp() *Nodes_Node{ return self.exp }
// 1: decl @lune.@base.@Nodes.ExpOp1Node.processFilter
func (self *Nodes_ExpOp1Node) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpOp1(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpOp1Node.canBeRight
func (self *Nodes_ExpOp1Node) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpOp1Node.canBeLeft
func (self *Nodes_ExpOp1Node) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpOp1Node.canBeStatement
func (self *Nodes_ExpOp1Node) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpOp1Node) InitNodes_ExpOp1Node(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,op *Types_Token,macroMode LnsInt,exp *Nodes_Node) {
    self.InitNodes_Node(id, 33, pos, macroArgFlag, typeList)
    self.op = op
    
    self.macroMode = macroMode
    
    self.exp = exp
    
}

// 681: decl @lune.@base.@Nodes.ExpOp1Node.create
func Nodes_ExpOp1Node_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,op *Types_Token,macroMode LnsInt,exp *Nodes_Node) *Nodes_ExpOp1Node {
    var node *Nodes_ExpOp1Node
    node = NewNodes_ExpOp1Node(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, op, macroMode, exp)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpOp1Node.visit
func (self *Nodes_ExpOp1Node) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch18892 := visitor(child, &self.Nodes_Node, "exp", depth); _switch18892 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch18892 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpRefItemNode
type Nodes_ExpRefItemNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_index() LnsAny
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_nilAccess() bool
    Get_pos() *Types_Position
    Get_symbol() LnsAny
    Get_tailComment() LnsAny
    Get_threading() bool
    Get_val() *Nodes_Node
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpRefItemNode struct {
    Nodes_Node
    val *Nodes_Node
    nilAccess bool
    threading bool
    symbol LnsAny
    index LnsAny
    FP Nodes_ExpRefItemNodeMtd
}
func Nodes_ExpRefItemNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpRefItemNode).FP
}
type Nodes_ExpRefItemNodeDownCast interface {
    ToNodes_ExpRefItemNode() *Nodes_ExpRefItemNode
}
func Nodes_ExpRefItemNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpRefItemNodeDownCast)
    if ok { return work.ToNodes_ExpRefItemNode() }
    return nil
}
func (obj *Nodes_ExpRefItemNode) ToNodes_ExpRefItemNode() *Nodes_ExpRefItemNode {
    return obj
}
func NewNodes_ExpRefItemNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 bool, arg7 bool, arg8 LnsAny, arg9 LnsAny) *Nodes_ExpRefItemNode {
    obj := &Nodes_ExpRefItemNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpRefItemNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ExpRefItemNode) Get_val() *Nodes_Node{ return self.val }
func (self *Nodes_ExpRefItemNode) Get_nilAccess() bool{ return self.nilAccess }
func (self *Nodes_ExpRefItemNode) Get_threading() bool{ return self.threading }
func (self *Nodes_ExpRefItemNode) Get_symbol() LnsAny{ return self.symbol }
func (self *Nodes_ExpRefItemNode) Get_index() LnsAny{ return self.index }
// 1: decl @lune.@base.@Nodes.ExpRefItemNode.hasNilAccess
func (self *Nodes_ExpRefItemNode) HasNilAccess() bool {
    return self.nilAccess
}

// 1: decl @lune.@base.@Nodes.ExpRefItemNode.isThreading
func (self *Nodes_ExpRefItemNode) IsThreading() bool {
    return self.threading
}

// 1: decl @lune.@base.@Nodes.ExpRefItemNode.processFilter
func (self *Nodes_ExpRefItemNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpRefItem(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpRefItemNode.canBeRight
func (self *Nodes_ExpRefItemNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpRefItemNode.canBeStatement
func (self *Nodes_ExpRefItemNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpRefItemNode) InitNodes_ExpRefItemNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,nilAccess bool,threading bool,symbol LnsAny,index LnsAny) {
    self.InitNodes_Node(id, 34, pos, macroArgFlag, typeList)
    self.val = val
    
    self.nilAccess = nilAccess
    
    self.threading = threading
    
    self.symbol = symbol
    
    self.index = index
    
}

// 681: decl @lune.@base.@Nodes.ExpRefItemNode.create
func Nodes_ExpRefItemNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,nilAccess bool,threading bool,symbol LnsAny,index LnsAny) *Nodes_ExpRefItemNode {
    var node *Nodes_ExpRefItemNode
    node = NewNodes_ExpRefItemNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, val, nilAccess, threading, symbol, index)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpRefItemNode.visit
func (self *Nodes_ExpRefItemNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.val
        if _switch19356 := visitor(child, &self.Nodes_Node, "val", depth); _switch19356 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch19356 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.index
            if _child != nil {
                child := _child.(*Nodes_Node)
                if _switch19411 := visitor(child, &self.Nodes_Node, "index", depth); _switch19411 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch19411 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1417: decl @lune.@base.@Nodes.ExpRefItemNode.getPrefix
func (self *Nodes_ExpRefItemNode) GetPrefix() LnsAny {
    return self.val
}

// 1420: decl @lune.@base.@Nodes.ExpRefItemNode.canBeLeft
func (self *Nodes_ExpRefItemNode) CanBeLeft() bool {
    if self.val.FP.Get_expType() == Ast_builtinTypeStem{
        return false
    }
    return Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Ast_TypeInfo_isMut(self.FP.Get_val().FP.Get_expType())) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(self.nilAccess)) ).(bool)
}


// declaration Class -- ExpCallNode
type Nodes_ExpCallNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_argList() LnsAny
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_errorFunc() bool
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_func() *Nodes_Node
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_nilAccess() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_threading() bool
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpCallNode struct {
    Nodes_Node
    _func *Nodes_Node
    errorFunc bool
    nilAccess bool
    threading bool
    argList LnsAny
    FP Nodes_ExpCallNodeMtd
}
func Nodes_ExpCallNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpCallNode).FP
}
type Nodes_ExpCallNodeDownCast interface {
    ToNodes_ExpCallNode() *Nodes_ExpCallNode
}
func Nodes_ExpCallNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpCallNodeDownCast)
    if ok { return work.ToNodes_ExpCallNode() }
    return nil
}
func (obj *Nodes_ExpCallNode) ToNodes_ExpCallNode() *Nodes_ExpCallNode {
    return obj
}
func NewNodes_ExpCallNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 bool, arg7 bool, arg8 bool, arg9 LnsAny) *Nodes_ExpCallNode {
    obj := &Nodes_ExpCallNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpCallNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ExpCallNode) Get_func() *Nodes_Node{ return self._func }
func (self *Nodes_ExpCallNode) Get_errorFunc() bool{ return self.errorFunc }
func (self *Nodes_ExpCallNode) Get_nilAccess() bool{ return self.nilAccess }
func (self *Nodes_ExpCallNode) Get_threading() bool{ return self.threading }
func (self *Nodes_ExpCallNode) Get_argList() LnsAny{ return self.argList }
// 1: decl @lune.@base.@Nodes.ExpCallNode.hasNilAccess
func (self *Nodes_ExpCallNode) HasNilAccess() bool {
    return self.nilAccess
}

// 1: decl @lune.@base.@Nodes.ExpCallNode.isThreading
func (self *Nodes_ExpCallNode) IsThreading() bool {
    return self.threading
}

// 1: decl @lune.@base.@Nodes.ExpCallNode.processFilter
func (self *Nodes_ExpCallNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCall(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpCallNode.canBeLeft
func (self *Nodes_ExpCallNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpCallNode.canBeStatement
func (self *Nodes_ExpCallNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ExpCallNode) InitNodes_ExpCallNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,_func *Nodes_Node,errorFunc bool,nilAccess bool,threading bool,argList LnsAny) {
    self.InitNodes_Node(id, 35, pos, macroArgFlag, typeList)
    self._func = _func
    
    self.errorFunc = errorFunc
    
    self.nilAccess = nilAccess
    
    self.threading = threading
    
    self.argList = argList
    
}

// 681: decl @lune.@base.@Nodes.ExpCallNode.create
func Nodes_ExpCallNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,_func *Nodes_Node,errorFunc bool,nilAccess bool,threading bool,argList LnsAny) *Nodes_ExpCallNode {
    var node *Nodes_ExpCallNode
    node = NewNodes_ExpCallNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, _func, errorFunc, nilAccess, threading, argList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpCallNode.visit
func (self *Nodes_ExpCallNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self._func
        if _switch19920 := visitor(child, &self.Nodes_Node, "func", depth); _switch19920 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch19920 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.argList
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch19976 := visitor(&child.Nodes_Node, &self.Nodes_Node, "argList", depth); _switch19976 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch19976 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1435: decl @lune.@base.@Nodes.ExpCallNode.get_effectivePos
func (self *Nodes_ExpCallNode) Get_effectivePos() *Types_Position {
    return self._func.FP.Get_effectivePos()
}

// 1438: decl @lune.@base.@Nodes.ExpCallNode.getPrefix
func (self *Nodes_ExpCallNode) GetPrefix() LnsAny {
    return self._func
}

// 1442: decl @lune.@base.@Nodes.ExpCallNode.canBeRight
func (self *Nodes_ExpCallNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    var expType *Ast_TypeInfo
    expType = self.FP.Get_expType()
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( expType.FP.Equals(processInfo, Ast_builtinTypeNone, nil, nil)) ||
        Lns_GetEnv().SetStackVal( expType.FP.Equals(processInfo, Ast_builtinTypeNeverRet, nil, nil)) ).(bool){
        return false
    }
    return true
}

// 1453: decl @lune.@base.@Nodes.ExpCallNode.getBreakKind
func (self *Nodes_ExpCallNode) GetBreakKind(checkMode LnsInt) LnsInt {
    if self.errorFunc{
        return Nodes_BreakKind__NeverRet
    }
    return Nodes_BreakKind__None
}


// declaration Class -- ExpMRetNode
type Nodes_ExpMRetNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_mRet() *Nodes_Node
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpMRetNode struct {
    Nodes_Node
    mRet *Nodes_Node
    FP Nodes_ExpMRetNodeMtd
}
func Nodes_ExpMRetNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpMRetNode).FP
}
type Nodes_ExpMRetNodeDownCast interface {
    ToNodes_ExpMRetNode() *Nodes_ExpMRetNode
}
func Nodes_ExpMRetNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpMRetNodeDownCast)
    if ok { return work.ToNodes_ExpMRetNode() }
    return nil
}
func (obj *Nodes_ExpMRetNode) ToNodes_ExpMRetNode() *Nodes_ExpMRetNode {
    return obj
}
func NewNodes_ExpMRetNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node) *Nodes_ExpMRetNode {
    obj := &Nodes_ExpMRetNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMRetNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ExpMRetNode) Get_mRet() *Nodes_Node{ return self.mRet }
// 1: decl @lune.@base.@Nodes.ExpMRetNode.processFilter
func (self *Nodes_ExpMRetNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMRet(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpMRetNode.canBeRight
func (self *Nodes_ExpMRetNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpMRetNode.canBeLeft
func (self *Nodes_ExpMRetNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpMRetNode.canBeStatement
func (self *Nodes_ExpMRetNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpMRetNode) InitNodes_ExpMRetNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,mRet *Nodes_Node) {
    self.InitNodes_Node(id, 36, pos, macroArgFlag, typeList)
    self.mRet = mRet
    
}

// 681: decl @lune.@base.@Nodes.ExpMRetNode.create
func Nodes_ExpMRetNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,mRet *Nodes_Node) *Nodes_ExpMRetNode {
    var node *Nodes_ExpMRetNode
    node = NewNodes_ExpMRetNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, mRet)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpMRetNode.visit
func (self *Nodes_ExpMRetNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.mRet
        if _switch20415 := visitor(child, &self.Nodes_Node, "mRet", depth); _switch20415 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch20415 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1464: decl @lune.@base.@Nodes.ExpMRetNode.getPrefix
func (self *Nodes_ExpMRetNode) GetPrefix() LnsAny {
    return self.mRet.FP.GetPrefix()
}


// declaration Class -- ExpAccessMRetNode
type Nodes_ExpAccessMRetNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_index() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_mRet() *Nodes_Node
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpAccessMRetNode struct {
    Nodes_Node
    mRet *Nodes_Node
    index LnsInt
    FP Nodes_ExpAccessMRetNodeMtd
}
func Nodes_ExpAccessMRetNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpAccessMRetNode).FP
}
type Nodes_ExpAccessMRetNodeDownCast interface {
    ToNodes_ExpAccessMRetNode() *Nodes_ExpAccessMRetNode
}
func Nodes_ExpAccessMRetNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpAccessMRetNodeDownCast)
    if ok { return work.ToNodes_ExpAccessMRetNode() }
    return nil
}
func (obj *Nodes_ExpAccessMRetNode) ToNodes_ExpAccessMRetNode() *Nodes_ExpAccessMRetNode {
    return obj
}
func NewNodes_ExpAccessMRetNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 LnsInt) *Nodes_ExpAccessMRetNode {
    obj := &Nodes_ExpAccessMRetNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpAccessMRetNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Nodes_ExpAccessMRetNode) Get_mRet() *Nodes_Node{ return self.mRet }
func (self *Nodes_ExpAccessMRetNode) Get_index() LnsInt{ return self.index }
// 1: decl @lune.@base.@Nodes.ExpAccessMRetNode.processFilter
func (self *Nodes_ExpAccessMRetNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpAccessMRet(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpAccessMRetNode.canBeRight
func (self *Nodes_ExpAccessMRetNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpAccessMRetNode.canBeLeft
func (self *Nodes_ExpAccessMRetNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpAccessMRetNode.canBeStatement
func (self *Nodes_ExpAccessMRetNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpAccessMRetNode) InitNodes_ExpAccessMRetNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,mRet *Nodes_Node,index LnsInt) {
    self.InitNodes_Node(id, 37, pos, macroArgFlag, typeList)
    self.mRet = mRet
    
    self.index = index
    
}

// 681: decl @lune.@base.@Nodes.ExpAccessMRetNode.create
func Nodes_ExpAccessMRetNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,mRet *Nodes_Node,index LnsInt) *Nodes_ExpAccessMRetNode {
    var node *Nodes_ExpAccessMRetNode
    node = NewNodes_ExpAccessMRetNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, mRet, index)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpAccessMRetNode.visit
func (self *Nodes_ExpAccessMRetNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.mRet
        if _switch20799 := visitor(child, &self.Nodes_Node, "mRet", depth); _switch20799 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch20799 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1473: decl @lune.@base.@Nodes.ExpAccessMRetNode.getPrefix
func (self *Nodes_ExpAccessMRetNode) GetPrefix() LnsAny {
    return self.mRet.FP.GetPrefix()
}


// declaration Class -- ExpMultiTo1Node
type Nodes_ExpMultiTo1NodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpMultiTo1Node struct {
    Nodes_Node
    exp *Nodes_Node
    FP Nodes_ExpMultiTo1NodeMtd
}
func Nodes_ExpMultiTo1Node2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpMultiTo1Node).FP
}
type Nodes_ExpMultiTo1NodeDownCast interface {
    ToNodes_ExpMultiTo1Node() *Nodes_ExpMultiTo1Node
}
func Nodes_ExpMultiTo1NodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpMultiTo1NodeDownCast)
    if ok { return work.ToNodes_ExpMultiTo1Node() }
    return nil
}
func (obj *Nodes_ExpMultiTo1Node) ToNodes_ExpMultiTo1Node() *Nodes_ExpMultiTo1Node {
    return obj
}
func NewNodes_ExpMultiTo1Node(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node) *Nodes_ExpMultiTo1Node {
    obj := &Nodes_ExpMultiTo1Node{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMultiTo1Node(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ExpMultiTo1Node) Get_exp() *Nodes_Node{ return self.exp }
// 1: decl @lune.@base.@Nodes.ExpMultiTo1Node.processFilter
func (self *Nodes_ExpMultiTo1Node) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMultiTo1(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpMultiTo1Node.canBeRight
func (self *Nodes_ExpMultiTo1Node) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpMultiTo1Node.canBeLeft
func (self *Nodes_ExpMultiTo1Node) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpMultiTo1Node.canBeStatement
func (self *Nodes_ExpMultiTo1Node) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpMultiTo1Node) InitNodes_ExpMultiTo1Node(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(id, 38, pos, macroArgFlag, typeList)
    self.exp = exp
    
}

// 681: decl @lune.@base.@Nodes.ExpMultiTo1Node.create
func Nodes_ExpMultiTo1Node_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_ExpMultiTo1Node {
    var node *Nodes_ExpMultiTo1Node
    node = NewNodes_ExpMultiTo1Node(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpMultiTo1Node.visit
func (self *Nodes_ExpMultiTo1Node) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch21156 := visitor(child, &self.Nodes_Node, "exp", depth); _switch21156 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch21156 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1480: decl @lune.@base.@Nodes.ExpMultiTo1Node.getPrefix
func (self *Nodes_ExpMultiTo1Node) GetPrefix() LnsAny {
    return self.exp.FP.GetPrefix()
}


// declaration Class -- ExpParenNode
type Nodes_ExpParenNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpParenNode struct {
    Nodes_Node
    exp *Nodes_Node
    FP Nodes_ExpParenNodeMtd
}
func Nodes_ExpParenNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpParenNode).FP
}
type Nodes_ExpParenNodeDownCast interface {
    ToNodes_ExpParenNode() *Nodes_ExpParenNode
}
func Nodes_ExpParenNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpParenNodeDownCast)
    if ok { return work.ToNodes_ExpParenNode() }
    return nil
}
func (obj *Nodes_ExpParenNode) ToNodes_ExpParenNode() *Nodes_ExpParenNode {
    return obj
}
func NewNodes_ExpParenNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node) *Nodes_ExpParenNode {
    obj := &Nodes_ExpParenNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpParenNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ExpParenNode) Get_exp() *Nodes_Node{ return self.exp }
// 1: decl @lune.@base.@Nodes.ExpParenNode.processFilter
func (self *Nodes_ExpParenNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpParen(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpParenNode.canBeRight
func (self *Nodes_ExpParenNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpParenNode.canBeLeft
func (self *Nodes_ExpParenNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpParenNode.canBeStatement
func (self *Nodes_ExpParenNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpParenNode) InitNodes_ExpParenNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(id, 39, pos, macroArgFlag, typeList)
    self.exp = exp
    
}

// 681: decl @lune.@base.@Nodes.ExpParenNode.create
func Nodes_ExpParenNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_ExpParenNode {
    var node *Nodes_ExpParenNode
    node = NewNodes_ExpParenNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpParenNode.visit
func (self *Nodes_ExpParenNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch21513 := visitor(child, &self.Nodes_Node, "exp", depth); _switch21513 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch21513 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1487: decl @lune.@base.@Nodes.ExpParenNode.getPrefix
func (self *Nodes_ExpParenNode) GetPrefix() LnsAny {
    return self.exp.FP.GetPrefix()
}

// 1491: decl @lune.@base.@Nodes.ExpParenNode.getSymbolInfo
func (self *Nodes_ExpParenNode) GetSymbolInfo() *LnsList {
    return self.exp.FP.GetSymbolInfo()
}


// declaration Class -- ExpMacroExpNode
type Nodes_ExpMacroExpNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_stmtList() *LnsList
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpMacroExpNode struct {
    Nodes_Node
    stmtList *LnsList
    FP Nodes_ExpMacroExpNodeMtd
}
func Nodes_ExpMacroExpNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpMacroExpNode).FP
}
type Nodes_ExpMacroExpNodeDownCast interface {
    ToNodes_ExpMacroExpNode() *Nodes_ExpMacroExpNode
}
func Nodes_ExpMacroExpNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpMacroExpNodeDownCast)
    if ok { return work.ToNodes_ExpMacroExpNode() }
    return nil
}
func (obj *Nodes_ExpMacroExpNode) ToNodes_ExpMacroExpNode() *Nodes_ExpMacroExpNode {
    return obj
}
func NewNodes_ExpMacroExpNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList) *Nodes_ExpMacroExpNode {
    obj := &Nodes_ExpMacroExpNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMacroExpNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ExpMacroExpNode) Get_stmtList() *LnsList{ return self.stmtList }
// 1: decl @lune.@base.@Nodes.ExpMacroExpNode.processFilter
func (self *Nodes_ExpMacroExpNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMacroExp(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpMacroExpNode.canBeLeft
func (self *Nodes_ExpMacroExpNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpMacroExpNode.canBeStatement
func (self *Nodes_ExpMacroExpNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ExpMacroExpNode) InitNodes_ExpMacroExpNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) {
    self.InitNodes_Node(id, 40, pos, macroArgFlag, typeList)
    self.stmtList = stmtList
    
}

// 681: decl @lune.@base.@Nodes.ExpMacroExpNode.create
func Nodes_ExpMacroExpNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) *Nodes_ExpMacroExpNode {
    var node *Nodes_ExpMacroExpNode
    node = NewNodes_ExpMacroExpNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, stmtList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpMacroExpNode.visit
func (self *Nodes_ExpMacroExpNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.stmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch21881 := visitor(child, &self.Nodes_Node, "stmtList", depth); _switch21881 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch21881 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}

// 1501: decl @lune.@base.@Nodes.ExpMacroExpNode.canBeRight
func (self *Nodes_ExpMacroExpNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return self.FP.Get_expType() != Ast_builtinTypeNone
}

// 1505: decl @lune.@base.@Nodes.ExpMacroExpNode.getBreakKind
func (self *Nodes_ExpMacroExpNode) GetBreakKind(checkMode LnsInt) LnsInt {
    return Nodes_getBreakKindForStmtList_2517_(checkMode, self.stmtList)
}


// declaration Class -- ExpMacroStatNode
type Nodes_ExpMacroStatNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expStrList() *LnsList
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpMacroStatNode struct {
    Nodes_Node
    expStrList *LnsList
    FP Nodes_ExpMacroStatNodeMtd
}
func Nodes_ExpMacroStatNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpMacroStatNode).FP
}
type Nodes_ExpMacroStatNodeDownCast interface {
    ToNodes_ExpMacroStatNode() *Nodes_ExpMacroStatNode
}
func Nodes_ExpMacroStatNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpMacroStatNodeDownCast)
    if ok { return work.ToNodes_ExpMacroStatNode() }
    return nil
}
func (obj *Nodes_ExpMacroStatNode) ToNodes_ExpMacroStatNode() *Nodes_ExpMacroStatNode {
    return obj
}
func NewNodes_ExpMacroStatNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList) *Nodes_ExpMacroStatNode {
    obj := &Nodes_ExpMacroStatNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMacroStatNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ExpMacroStatNode) Get_expStrList() *LnsList{ return self.expStrList }
// 1: decl @lune.@base.@Nodes.ExpMacroStatNode.processFilter
func (self *Nodes_ExpMacroStatNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMacroStat(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpMacroStatNode.canBeRight
func (self *Nodes_ExpMacroStatNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpMacroStatNode.canBeLeft
func (self *Nodes_ExpMacroStatNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpMacroStatNode.canBeStatement
func (self *Nodes_ExpMacroStatNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpMacroStatNode) InitNodes_ExpMacroStatNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expStrList *LnsList) {
    self.InitNodes_Node(id, 41, pos, macroArgFlag, typeList)
    self.expStrList = expStrList
    
}

// 681: decl @lune.@base.@Nodes.ExpMacroStatNode.create
func Nodes_ExpMacroStatNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expStrList *LnsList) *Nodes_ExpMacroStatNode {
    var node *Nodes_ExpMacroStatNode
    node = NewNodes_ExpMacroStatNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, expStrList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpMacroStatNode.visit
func (self *Nodes_ExpMacroStatNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.expStrList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch22278 := visitor(child, &self.Nodes_Node, "expStrList", depth); _switch22278 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch22278 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}

// 2526: decl @lune.@base.@Nodes.ExpMacroStatNode.getLiteral
func (self *Nodes_ExpMacroStatNode) GetLiteral()(LnsAny, LnsAny) {
    var txt string
    txt = ""
    for _, _token := range( self.expStrList.Items ) {
        token := _token.(Nodes_NodeDownCast).ToNodes_Node()
        var literal LnsAny
        literal = Nodes_convExp43351(Lns_2DDD(token.FP.GetLiteral()))
        if literal != nil{
            literal_10645 := literal
            switch _exp43373 := literal_10645.(type) {
            case *Nodes_Literal__Str:
            work := _exp43373.Val1
                txt = Lns_getVM().String_format("%s%s", []LnsAny{txt, work})
                
            }
        } else {
            return nil, Lns_getVM().String_format("illegal literal -- %s", []LnsAny{Nodes_getNodeKindName(token.FP.Get_kind())})
        }
    }
    return &Nodes_Literal__Str{txt}, nil
}


// declaration Class -- ExpMacroArgExpNode
type Nodes_ExpMacroArgExpNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_codeTxt() string
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpMacroArgExpNode struct {
    Nodes_Node
    codeTxt string
    FP Nodes_ExpMacroArgExpNodeMtd
}
func Nodes_ExpMacroArgExpNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpMacroArgExpNode).FP
}
type Nodes_ExpMacroArgExpNodeDownCast interface {
    ToNodes_ExpMacroArgExpNode() *Nodes_ExpMacroArgExpNode
}
func Nodes_ExpMacroArgExpNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpMacroArgExpNodeDownCast)
    if ok { return work.ToNodes_ExpMacroArgExpNode() }
    return nil
}
func (obj *Nodes_ExpMacroArgExpNode) ToNodes_ExpMacroArgExpNode() *Nodes_ExpMacroArgExpNode {
    return obj
}
func NewNodes_ExpMacroArgExpNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 string) *Nodes_ExpMacroArgExpNode {
    obj := &Nodes_ExpMacroArgExpNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMacroArgExpNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ExpMacroArgExpNode) Get_codeTxt() string{ return self.codeTxt }
// 1: decl @lune.@base.@Nodes.ExpMacroArgExpNode.processFilter
func (self *Nodes_ExpMacroArgExpNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMacroArgExp(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpMacroArgExpNode.canBeRight
func (self *Nodes_ExpMacroArgExpNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpMacroArgExpNode.canBeLeft
func (self *Nodes_ExpMacroArgExpNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpMacroArgExpNode.canBeStatement
func (self *Nodes_ExpMacroArgExpNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpMacroArgExpNode) InitNodes_ExpMacroArgExpNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,codeTxt string) {
    self.InitNodes_Node(id, 42, pos, macroArgFlag, typeList)
    self.codeTxt = codeTxt
    
}

// 681: decl @lune.@base.@Nodes.ExpMacroArgExpNode.create
func Nodes_ExpMacroArgExpNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,codeTxt string) *Nodes_ExpMacroArgExpNode {
    var node *Nodes_ExpMacroArgExpNode
    node = NewNodes_ExpMacroArgExpNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, codeTxt)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpMacroArgExpNode.visit
func (self *Nodes_ExpMacroArgExpNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 2545: decl @lune.@base.@Nodes.ExpMacroArgExpNode.getLiteral
func (self *Nodes_ExpMacroArgExpNode) GetLiteral()(LnsAny, LnsAny) {
    return &Nodes_Literal__Str{self.FP.Get_codeTxt()}, nil
}


// declaration Class -- StmtExpNode
type Nodes_StmtExpNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_StmtExpNode struct {
    Nodes_Node
    exp *Nodes_Node
    FP Nodes_StmtExpNodeMtd
}
func Nodes_StmtExpNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_StmtExpNode).FP
}
type Nodes_StmtExpNodeDownCast interface {
    ToNodes_StmtExpNode() *Nodes_StmtExpNode
}
func Nodes_StmtExpNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_StmtExpNodeDownCast)
    if ok { return work.ToNodes_StmtExpNode() }
    return nil
}
func (obj *Nodes_StmtExpNode) ToNodes_StmtExpNode() *Nodes_StmtExpNode {
    return obj
}
func NewNodes_StmtExpNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node) *Nodes_StmtExpNode {
    obj := &Nodes_StmtExpNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_StmtExpNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_StmtExpNode) Get_exp() *Nodes_Node{ return self.exp }
// 1: decl @lune.@base.@Nodes.StmtExpNode.processFilter
func (self *Nodes_StmtExpNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessStmtExp(self, opt)
}

// 1: decl @lune.@base.@Nodes.StmtExpNode.canBeRight
func (self *Nodes_StmtExpNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.StmtExpNode.canBeLeft
func (self *Nodes_StmtExpNode) CanBeLeft() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_StmtExpNode) InitNodes_StmtExpNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(id, 43, pos, macroArgFlag, typeList)
    self.exp = exp
    
}

// 681: decl @lune.@base.@Nodes.StmtExpNode.create
func Nodes_StmtExpNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_StmtExpNode {
    var node *Nodes_StmtExpNode
    node = NewNodes_StmtExpNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.StmtExpNode.visit
func (self *Nodes_StmtExpNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch22905 := visitor(child, &self.Nodes_Node, "exp", depth); _switch22905 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch22905 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1526: decl @lune.@base.@Nodes.StmtExpNode.canBeStatement
func (self *Nodes_StmtExpNode) CanBeStatement() bool {
    return self.FP.Get_exp().FP.CanBeStatement()
}

// 1530: decl @lune.@base.@Nodes.StmtExpNode.getBreakKind
func (self *Nodes_StmtExpNode) GetBreakKind(checkMode LnsInt) LnsInt {
    return self.FP.Get_exp().FP.GetBreakKind(checkMode)
}


// declaration Class -- ExpMacroStatListNode
type Nodes_ExpMacroStatListNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpMacroStatListNode struct {
    Nodes_Node
    exp *Nodes_Node
    FP Nodes_ExpMacroStatListNodeMtd
}
func Nodes_ExpMacroStatListNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpMacroStatListNode).FP
}
type Nodes_ExpMacroStatListNodeDownCast interface {
    ToNodes_ExpMacroStatListNode() *Nodes_ExpMacroStatListNode
}
func Nodes_ExpMacroStatListNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpMacroStatListNodeDownCast)
    if ok { return work.ToNodes_ExpMacroStatListNode() }
    return nil
}
func (obj *Nodes_ExpMacroStatListNode) ToNodes_ExpMacroStatListNode() *Nodes_ExpMacroStatListNode {
    return obj
}
func NewNodes_ExpMacroStatListNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node) *Nodes_ExpMacroStatListNode {
    obj := &Nodes_ExpMacroStatListNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMacroStatListNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ExpMacroStatListNode) Get_exp() *Nodes_Node{ return self.exp }
// 1: decl @lune.@base.@Nodes.ExpMacroStatListNode.processFilter
func (self *Nodes_ExpMacroStatListNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMacroStatList(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpMacroStatListNode.canBeRight
func (self *Nodes_ExpMacroStatListNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpMacroStatListNode.canBeLeft
func (self *Nodes_ExpMacroStatListNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpMacroStatListNode.canBeStatement
func (self *Nodes_ExpMacroStatListNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpMacroStatListNode) InitNodes_ExpMacroStatListNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(id, 44, pos, macroArgFlag, typeList)
    self.exp = exp
    
}

// 681: decl @lune.@base.@Nodes.ExpMacroStatListNode.create
func Nodes_ExpMacroStatListNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_ExpMacroStatListNode {
    var node *Nodes_ExpMacroStatListNode
    node = NewNodes_ExpMacroStatListNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpMacroStatListNode.visit
func (self *Nodes_ExpMacroStatListNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch23281 := visitor(child, &self.Nodes_Node, "exp", depth); _switch23281 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch23281 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- ExpOmitEnumNode
type Nodes_ExpOmitEnumNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_aliasType() LnsAny
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_enumTypeInfo() *Ast_EnumTypeInfo
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_valInfo() *Ast_EnumValInfo
    Get_valToken() *Types_Token
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpOmitEnumNode struct {
    Nodes_Node
    valToken *Types_Token
    valInfo *Ast_EnumValInfo
    aliasType LnsAny
    enumTypeInfo *Ast_EnumTypeInfo
    FP Nodes_ExpOmitEnumNodeMtd
}
func Nodes_ExpOmitEnumNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpOmitEnumNode).FP
}
type Nodes_ExpOmitEnumNodeDownCast interface {
    ToNodes_ExpOmitEnumNode() *Nodes_ExpOmitEnumNode
}
func Nodes_ExpOmitEnumNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpOmitEnumNodeDownCast)
    if ok { return work.ToNodes_ExpOmitEnumNode() }
    return nil
}
func (obj *Nodes_ExpOmitEnumNode) ToNodes_ExpOmitEnumNode() *Nodes_ExpOmitEnumNode {
    return obj
}
func NewNodes_ExpOmitEnumNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 *Ast_EnumValInfo, arg7 LnsAny, arg8 *Ast_EnumTypeInfo) *Nodes_ExpOmitEnumNode {
    obj := &Nodes_ExpOmitEnumNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpOmitEnumNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_ExpOmitEnumNode) Get_valToken() *Types_Token{ return self.valToken }
func (self *Nodes_ExpOmitEnumNode) Get_valInfo() *Ast_EnumValInfo{ return self.valInfo }
func (self *Nodes_ExpOmitEnumNode) Get_aliasType() LnsAny{ return self.aliasType }
func (self *Nodes_ExpOmitEnumNode) Get_enumTypeInfo() *Ast_EnumTypeInfo{ return self.enumTypeInfo }
// 1: decl @lune.@base.@Nodes.ExpOmitEnumNode.processFilter
func (self *Nodes_ExpOmitEnumNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpOmitEnum(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpOmitEnumNode.canBeRight
func (self *Nodes_ExpOmitEnumNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpOmitEnumNode.canBeLeft
func (self *Nodes_ExpOmitEnumNode) CanBeLeft() bool {
    return true
}

// 1: decl @lune.@base.@Nodes.ExpOmitEnumNode.canBeStatement
func (self *Nodes_ExpOmitEnumNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_ExpOmitEnumNode) InitNodes_ExpOmitEnumNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,valToken *Types_Token,valInfo *Ast_EnumValInfo,aliasType LnsAny,enumTypeInfo *Ast_EnumTypeInfo) {
    self.InitNodes_Node(id, 45, pos, macroArgFlag, typeList)
    self.valToken = valToken
    
    self.valInfo = valInfo
    
    self.aliasType = aliasType
    
    self.enumTypeInfo = enumTypeInfo
    
}

// 681: decl @lune.@base.@Nodes.ExpOmitEnumNode.create
func Nodes_ExpOmitEnumNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,valToken *Types_Token,valInfo *Ast_EnumValInfo,aliasType LnsAny,enumTypeInfo *Ast_EnumTypeInfo) *Nodes_ExpOmitEnumNode {
    var node *Nodes_ExpOmitEnumNode
    node = NewNodes_ExpOmitEnumNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, valToken, valInfo, aliasType, enumTypeInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpOmitEnumNode.visit
func (self *Nodes_ExpOmitEnumNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 2570: decl @lune.@base.@Nodes.ExpOmitEnumNode.getLiteral
func (self *Nodes_ExpOmitEnumNode) GetLiteral()(LnsAny, LnsAny) {
    var enumval *Ast_EnumValInfo
    enumval = self.valInfo
    return Nodes_enumLiteral2Literal_9640_(enumval.FP.Get_val())
}

// 2578: decl @lune.@base.@Nodes.ExpOmitEnumNode.setupLiteralTokenList
func (self *Nodes_ExpOmitEnumNode) SetupLiteralTokenList(list *LnsList) bool {
    var enumval *Ast_EnumValInfo
    enumval = self.valInfo
    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, ".")
    self.FP.AddTokenList(list, Types_TokenKind__Symb, Lns_car(Lns_getVM().String_gsub(enumval.FP.Get_val().(LnsAlgeVal).GetTxt(),".*%.", "")).(string))
    return true
}


// declaration Class -- RefFieldNode
type Nodes_RefFieldNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_field() *Types_Token
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_nilAccess() bool
    Get_pos() *Types_Position
    Get_prefix() *Nodes_Node
    Get_symbolInfo() LnsAny
    Get_tailComment() LnsAny
    Get_threading() bool
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_RefFieldNode struct {
    Nodes_Node
    field *Types_Token
    symbolInfo LnsAny
    nilAccess bool
    threading bool
    prefix *Nodes_Node
    FP Nodes_RefFieldNodeMtd
}
func Nodes_RefFieldNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_RefFieldNode).FP
}
type Nodes_RefFieldNodeDownCast interface {
    ToNodes_RefFieldNode() *Nodes_RefFieldNode
}
func Nodes_RefFieldNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_RefFieldNodeDownCast)
    if ok { return work.ToNodes_RefFieldNode() }
    return nil
}
func (obj *Nodes_RefFieldNode) ToNodes_RefFieldNode() *Nodes_RefFieldNode {
    return obj
}
func NewNodes_RefFieldNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 LnsAny, arg7 bool, arg8 bool, arg9 *Nodes_Node) *Nodes_RefFieldNode {
    obj := &Nodes_RefFieldNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_RefFieldNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_RefFieldNode) Get_field() *Types_Token{ return self.field }
func (self *Nodes_RefFieldNode) Get_symbolInfo() LnsAny{ return self.symbolInfo }
func (self *Nodes_RefFieldNode) Get_nilAccess() bool{ return self.nilAccess }
func (self *Nodes_RefFieldNode) Get_threading() bool{ return self.threading }
func (self *Nodes_RefFieldNode) Get_prefix() *Nodes_Node{ return self.prefix }
// 1: decl @lune.@base.@Nodes.RefFieldNode.hasNilAccess
func (self *Nodes_RefFieldNode) HasNilAccess() bool {
    return self.nilAccess
}

// 1: decl @lune.@base.@Nodes.RefFieldNode.isThreading
func (self *Nodes_RefFieldNode) IsThreading() bool {
    return self.threading
}

// 1: decl @lune.@base.@Nodes.RefFieldNode.processFilter
func (self *Nodes_RefFieldNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRefField(self, opt)
}

// 1: decl @lune.@base.@Nodes.RefFieldNode.canBeStatement
func (self *Nodes_RefFieldNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_RefFieldNode) InitNodes_RefFieldNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,field *Types_Token,symbolInfo LnsAny,nilAccess bool,threading bool,prefix *Nodes_Node) {
    self.InitNodes_Node(id, 46, pos, macroArgFlag, typeList)
    self.field = field
    
    self.symbolInfo = symbolInfo
    
    self.nilAccess = nilAccess
    
    self.threading = threading
    
    self.prefix = prefix
    
}

// 681: decl @lune.@base.@Nodes.RefFieldNode.create
func Nodes_RefFieldNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,field *Types_Token,symbolInfo LnsAny,nilAccess bool,threading bool,prefix *Nodes_Node) *Nodes_RefFieldNode {
    var node *Nodes_RefFieldNode
    node = NewNodes_RefFieldNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, threading, prefix)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.RefFieldNode.visit
func (self *Nodes_RefFieldNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.prefix
        if _switch24119 := visitor(child, &self.Nodes_Node, "prefix", depth); _switch24119 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch24119 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1554: decl @lune.@base.@Nodes.RefFieldNode.get_effectivePos
func (self *Nodes_RefFieldNode) Get_effectivePos() *Types_Position {
    return self.field.Pos
}

// 1557: decl @lune.@base.@Nodes.RefFieldNode.getPrefix
func (self *Nodes_RefFieldNode) GetPrefix() LnsAny {
    return self.prefix
}

// 1560: decl @lune.@base.@Nodes.RefFieldNode.canBeLeft
func (self *Nodes_RefFieldNode) CanBeLeft() bool {
    {
        __exp := self.FP.Get_symbolInfo()
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_canBeLeft()
        }
    }
    return false
}

// 1575: decl @lune.@base.@Nodes.RefFieldNode.canBeRight
func (self *Nodes_RefFieldNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    {
        __exp := self.FP.Get_symbolInfo()
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_canBeRight()
        }
    }
    return true
}

// 2489: decl @lune.@base.@Nodes.RefFieldNode.getLiteral
func (self *Nodes_RefFieldNode) GetLiteral()(LnsAny, LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = self.FP.Get_expType()
    {
        _enumTypeInfo := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_aliasSrc().FP)
        if _enumTypeInfo != nil {
            enumTypeInfo := _enumTypeInfo.(*Ast_EnumTypeInfo)
            if Lns_isCondTrue( Ast_EnumTypeInfoDownCastF(self.prefix.FP.Get_expType().FP.Get_aliasSrc().FP)){
                var enumval *Ast_EnumValInfo
                enumval = Lns_unwrap( enumTypeInfo.FP.GetEnumValInfo(self.field.Txt)).(*Ast_EnumValInfo)
                return Nodes_enumLiteral2Literal_9640_(enumval.FP.Get_val())
            }
        }
    }
    var tokenList *LnsList
    tokenList = NewLnsList([]LnsAny{})
    var literal LnsAny
    var mess LnsAny
    literal,mess = self.prefix.FP.GetLiteral()
    if literal != nil{
        literal_10625 := literal
        switch _exp43272 := literal_10625.(type) {
        case *Nodes_Literal__Symbol:
        symbol := _exp43272.Val1
            tokenList.Insert(symbol)
        case *Nodes_Literal__Field:
        symList := _exp43272.Val1
            for _, _symbol := range( symList.Items ) {
                symbol := _symbol.(string)
                tokenList.Insert(symbol)
            }
        default:
            return nil, Lns_getVM().String_format("not support -- %s", []LnsAny{literal_10625.(LnsAlgeVal).GetTxt()})
        }
        if self.nilAccess{
            tokenList.Insert("$.")
        } else { 
            tokenList.Insert(".")
        }
        tokenList.Insert(self.field.Txt)
        return &Nodes_Literal__Field{tokenList}, nil
    }
    return nil, mess
}


// declaration Class -- GetFieldNode
type Nodes_GetFieldNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_field() *Types_Token
    Get_getterTypeInfo() *Ast_TypeInfo
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_nilAccess() bool
    Get_pos() *Types_Position
    Get_prefix() *Nodes_Node
    Get_symbolInfo() LnsAny
    Get_tailComment() LnsAny
    Get_threading() bool
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_GetFieldNode struct {
    Nodes_Node
    field *Types_Token
    symbolInfo LnsAny
    nilAccess bool
    threading bool
    prefix *Nodes_Node
    getterTypeInfo *Ast_TypeInfo
    FP Nodes_GetFieldNodeMtd
}
func Nodes_GetFieldNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_GetFieldNode).FP
}
type Nodes_GetFieldNodeDownCast interface {
    ToNodes_GetFieldNode() *Nodes_GetFieldNode
}
func Nodes_GetFieldNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_GetFieldNodeDownCast)
    if ok { return work.ToNodes_GetFieldNode() }
    return nil
}
func (obj *Nodes_GetFieldNode) ToNodes_GetFieldNode() *Nodes_GetFieldNode {
    return obj
}
func NewNodes_GetFieldNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 LnsAny, arg7 bool, arg8 bool, arg9 *Nodes_Node, arg10 *Ast_TypeInfo) *Nodes_GetFieldNode {
    obj := &Nodes_GetFieldNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_GetFieldNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_GetFieldNode) Get_field() *Types_Token{ return self.field }
func (self *Nodes_GetFieldNode) Get_symbolInfo() LnsAny{ return self.symbolInfo }
func (self *Nodes_GetFieldNode) Get_nilAccess() bool{ return self.nilAccess }
func (self *Nodes_GetFieldNode) Get_threading() bool{ return self.threading }
func (self *Nodes_GetFieldNode) Get_prefix() *Nodes_Node{ return self.prefix }
func (self *Nodes_GetFieldNode) Get_getterTypeInfo() *Ast_TypeInfo{ return self.getterTypeInfo }
// 1: decl @lune.@base.@Nodes.GetFieldNode.hasNilAccess
func (self *Nodes_GetFieldNode) HasNilAccess() bool {
    return self.nilAccess
}

// 1: decl @lune.@base.@Nodes.GetFieldNode.isThreading
func (self *Nodes_GetFieldNode) IsThreading() bool {
    return self.threading
}

// 1: decl @lune.@base.@Nodes.GetFieldNode.processFilter
func (self *Nodes_GetFieldNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessGetField(self, opt)
}

// 1: decl @lune.@base.@Nodes.GetFieldNode.canBeRight
func (self *Nodes_GetFieldNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.GetFieldNode.canBeStatement
func (self *Nodes_GetFieldNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_GetFieldNode) InitNodes_GetFieldNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,field *Types_Token,symbolInfo LnsAny,nilAccess bool,threading bool,prefix *Nodes_Node,getterTypeInfo *Ast_TypeInfo) {
    self.InitNodes_Node(id, 47, pos, macroArgFlag, typeList)
    self.field = field
    
    self.symbolInfo = symbolInfo
    
    self.nilAccess = nilAccess
    
    self.threading = threading
    
    self.prefix = prefix
    
    self.getterTypeInfo = getterTypeInfo
    
}

// 681: decl @lune.@base.@Nodes.GetFieldNode.create
func Nodes_GetFieldNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,field *Types_Token,symbolInfo LnsAny,nilAccess bool,threading bool,prefix *Nodes_Node,getterTypeInfo *Ast_TypeInfo) *Nodes_GetFieldNode {
    var node *Nodes_GetFieldNode
    node = NewNodes_GetFieldNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, field, symbolInfo, nilAccess, threading, prefix, getterTypeInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.GetFieldNode.visit
func (self *Nodes_GetFieldNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.prefix
        if _switch24685 := visitor(child, &self.Nodes_Node, "prefix", depth); _switch24685 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch24685 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}

// 1594: decl @lune.@base.@Nodes.GetFieldNode.get_effectivePos
func (self *Nodes_GetFieldNode) Get_effectivePos() *Types_Position {
    return self.field.Pos
}

// 1597: decl @lune.@base.@Nodes.GetFieldNode.canBeLeft
func (self *Nodes_GetFieldNode) CanBeLeft() bool {
    {
        __exp := self.FP.Get_symbolInfo()
        if __exp != nil {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_canBeLeft()
        }
    }
    return false
}

// 1606: decl @lune.@base.@Nodes.GetFieldNode.getPrefix
func (self *Nodes_GetFieldNode) GetPrefix() LnsAny {
    return self.prefix
}


// declaration Class -- AliasNode
type Nodes_AliasNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_newSymbol() *Ast_SymbolInfo
    Get_pos() *Types_Position
    Get_srcNode() *Nodes_Node
    Get_tailComment() LnsAny
    Get_typeInfo() *Ast_TypeInfo
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_AliasNode struct {
    Nodes_Node
    newSymbol *Ast_SymbolInfo
    srcNode *Nodes_Node
    typeInfo *Ast_TypeInfo
    FP Nodes_AliasNodeMtd
}
func Nodes_AliasNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_AliasNode).FP
}
type Nodes_AliasNodeDownCast interface {
    ToNodes_AliasNode() *Nodes_AliasNode
}
func Nodes_AliasNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_AliasNodeDownCast)
    if ok { return work.ToNodes_AliasNode() }
    return nil
}
func (obj *Nodes_AliasNode) ToNodes_AliasNode() *Nodes_AliasNode {
    return obj
}
func NewNodes_AliasNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Ast_SymbolInfo, arg6 *Nodes_Node, arg7 *Ast_TypeInfo) *Nodes_AliasNode {
    obj := &Nodes_AliasNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_AliasNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_AliasNode) Get_newSymbol() *Ast_SymbolInfo{ return self.newSymbol }
func (self *Nodes_AliasNode) Get_srcNode() *Nodes_Node{ return self.srcNode }
func (self *Nodes_AliasNode) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }
// 1: decl @lune.@base.@Nodes.AliasNode.processFilter
func (self *Nodes_AliasNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessAlias(self, opt)
}

// 1: decl @lune.@base.@Nodes.AliasNode.canBeRight
func (self *Nodes_AliasNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.AliasNode.canBeLeft
func (self *Nodes_AliasNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.AliasNode.canBeStatement
func (self *Nodes_AliasNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_AliasNode) InitNodes_AliasNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,newSymbol *Ast_SymbolInfo,srcNode *Nodes_Node,typeInfo *Ast_TypeInfo) {
    self.InitNodes_Node(id, 48, pos, macroArgFlag, typeList)
    self.newSymbol = newSymbol
    
    self.srcNode = srcNode
    
    self.typeInfo = typeInfo
    
}

// 681: decl @lune.@base.@Nodes.AliasNode.create
func Nodes_AliasNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,newSymbol *Ast_SymbolInfo,srcNode *Nodes_Node,typeInfo *Ast_TypeInfo) *Nodes_AliasNode {
    var node *Nodes_AliasNode
    node = NewNodes_AliasNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, newSymbol, srcNode, typeInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.AliasNode.visit
func (self *Nodes_AliasNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.srcNode
        if _switch25132 := visitor(child, &self.Nodes_Node, "srcNode", depth); _switch25132 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch25132 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- VarInfo
type Nodes_VarInfoMtd interface {
    Get_actualType() *Ast_TypeInfo
    Get_name() *Types_Token
    Get_refType() LnsAny
}
type Nodes_VarInfo struct {
    name *Types_Token
    refType LnsAny
    actualType *Ast_TypeInfo
    FP Nodes_VarInfoMtd
}
func Nodes_VarInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_VarInfo).FP
}
type Nodes_VarInfoDownCast interface {
    ToNodes_VarInfo() *Nodes_VarInfo
}
func Nodes_VarInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_VarInfoDownCast)
    if ok { return work.ToNodes_VarInfo() }
    return nil
}
func (obj *Nodes_VarInfo) ToNodes_VarInfo() *Nodes_VarInfo {
    return obj
}
func NewNodes_VarInfo(arg1 *Types_Token, arg2 LnsAny, arg3 *Ast_TypeInfo) *Nodes_VarInfo {
    obj := &Nodes_VarInfo{}
    obj.FP = obj
    obj.InitNodes_VarInfo(arg1, arg2, arg3)
    return obj
}
func (self *Nodes_VarInfo) InitNodes_VarInfo(arg1 *Types_Token, arg2 LnsAny, arg3 *Ast_TypeInfo) {
    self.name = arg1
    self.refType = arg2
    self.actualType = arg3
}
func (self *Nodes_VarInfo) Get_name() *Types_Token{ return self.name }
func (self *Nodes_VarInfo) Get_refType() LnsAny{ return self.refType }
func (self *Nodes_VarInfo) Get_actualType() *Ast_TypeInfo{ return self.actualType }

// declaration Class -- DeclVarNode
type Nodes_DeclVarNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_accessMode() LnsInt
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() LnsAny
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_mode() LnsInt
    Get_pos() *Types_Position
    Get_staticFlag() bool
    Get_symbolInfoList() *LnsList
    Get_syncBlock() LnsAny
    Get_syncVarList() *LnsList
    Get_tailComment() LnsAny
    Get_thenBlock() LnsAny
    Get_typeInfoList() *LnsList
    Get_unwrapBlock() LnsAny
    Get_unwrapFlag() bool
    Get_varList() *LnsList
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclVarNode struct {
    Nodes_Node
    mode LnsInt
    accessMode LnsInt
    staticFlag bool
    varList *LnsList
    expList LnsAny
    symbolInfoList *LnsList
    typeInfoList *LnsList
    unwrapFlag bool
    unwrapBlock LnsAny
    thenBlock LnsAny
    syncVarList *LnsList
    syncBlock LnsAny
    FP Nodes_DeclVarNodeMtd
}
func Nodes_DeclVarNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclVarNode).FP
}
type Nodes_DeclVarNodeDownCast interface {
    ToNodes_DeclVarNode() *Nodes_DeclVarNode
}
func Nodes_DeclVarNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclVarNodeDownCast)
    if ok { return work.ToNodes_DeclVarNode() }
    return nil
}
func (obj *Nodes_DeclVarNode) ToNodes_DeclVarNode() *Nodes_DeclVarNode {
    return obj
}
func NewNodes_DeclVarNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsInt, arg6 LnsInt, arg7 bool, arg8 *LnsList, arg9 LnsAny, arg10 *LnsList, arg11 *LnsList, arg12 bool, arg13 LnsAny, arg14 LnsAny, arg15 *LnsList, arg16 LnsAny) *Nodes_DeclVarNode {
    obj := &Nodes_DeclVarNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclVarNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16)
    return obj
}
func (self *Nodes_DeclVarNode) Get_mode() LnsInt{ return self.mode }
func (self *Nodes_DeclVarNode) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Nodes_DeclVarNode) Get_staticFlag() bool{ return self.staticFlag }
func (self *Nodes_DeclVarNode) Get_varList() *LnsList{ return self.varList }
func (self *Nodes_DeclVarNode) Get_expList() LnsAny{ return self.expList }
func (self *Nodes_DeclVarNode) Get_symbolInfoList() *LnsList{ return self.symbolInfoList }
func (self *Nodes_DeclVarNode) Get_typeInfoList() *LnsList{ return self.typeInfoList }
func (self *Nodes_DeclVarNode) Get_unwrapFlag() bool{ return self.unwrapFlag }
func (self *Nodes_DeclVarNode) Get_unwrapBlock() LnsAny{ return self.unwrapBlock }
func (self *Nodes_DeclVarNode) Get_thenBlock() LnsAny{ return self.thenBlock }
func (self *Nodes_DeclVarNode) Get_syncVarList() *LnsList{ return self.syncVarList }
func (self *Nodes_DeclVarNode) Get_syncBlock() LnsAny{ return self.syncBlock }
// 1: decl @lune.@base.@Nodes.DeclVarNode.processFilter
func (self *Nodes_DeclVarNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclVar(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclVarNode.canBeRight
func (self *Nodes_DeclVarNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclVarNode.canBeLeft
func (self *Nodes_DeclVarNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclVarNode.canBeStatement
func (self *Nodes_DeclVarNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclVarNode) InitNodes_DeclVarNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,mode LnsInt,accessMode LnsInt,staticFlag bool,varList *LnsList,expList LnsAny,symbolInfoList *LnsList,typeInfoList *LnsList,unwrapFlag bool,unwrapBlock LnsAny,thenBlock LnsAny,syncVarList *LnsList,syncBlock LnsAny) {
    self.InitNodes_Node(id, 49, pos, macroArgFlag, typeList)
    self.mode = mode
    
    self.accessMode = accessMode
    
    self.staticFlag = staticFlag
    
    self.varList = varList
    
    self.expList = expList
    
    self.symbolInfoList = symbolInfoList
    
    self.typeInfoList = typeInfoList
    
    self.unwrapFlag = unwrapFlag
    
    self.unwrapBlock = unwrapBlock
    
    self.thenBlock = thenBlock
    
    self.syncVarList = syncVarList
    
    self.syncBlock = syncBlock
    
}

// 681: decl @lune.@base.@Nodes.DeclVarNode.create
func Nodes_DeclVarNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,mode LnsInt,accessMode LnsInt,staticFlag bool,varList *LnsList,expList LnsAny,symbolInfoList *LnsList,typeInfoList *LnsList,unwrapFlag bool,unwrapBlock LnsAny,thenBlock LnsAny,syncVarList *LnsList,syncBlock LnsAny) *Nodes_DeclVarNode {
    var node *Nodes_DeclVarNode
    node = NewNodes_DeclVarNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclVarNode.visit
func (self *Nodes_DeclVarNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch25825 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch25825 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch25825 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    {
        {
            _child := self.unwrapBlock
            if _child != nil {
                child := _child.(*Nodes_BlockNode)
                if _switch25883 := visitor(&child.Nodes_Node, &self.Nodes_Node, "unwrapBlock", depth); _switch25883 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch25883 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    {
        {
            _child := self.thenBlock
            if _child != nil {
                child := _child.(*Nodes_BlockNode)
                if _switch25941 := visitor(&child.Nodes_Node, &self.Nodes_Node, "thenBlock", depth); _switch25941 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch25941 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    {
        {
            _child := self.syncBlock
            if _child != nil {
                child := _child.(*Nodes_BlockNode)
                if _switch25999 := visitor(&child.Nodes_Node, &self.Nodes_Node, "syncBlock", depth); _switch25999 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch25999 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1647: decl @lune.@base.@Nodes.DeclVarNode.getBreakKind
func (self *Nodes_DeclVarNode) GetBreakKind(checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = Nodes_BreakKind__None
    var work LnsInt
    work = Nodes_BreakKind__None
    {
        _block := self.unwrapBlock
        if _block != nil {
            block := _block.(*Nodes_BlockNode)
            work = block.FP.GetBreakKind(checkMode)
            
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch26138 := work; _switch26138 == Nodes_BreakKind__None {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                        Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                        kind = work
                        
                    }
                }
            }
            
            {
                _thenBlock := self.thenBlock
                if _thenBlock != nil {
                    thenBlock := _thenBlock.(*Nodes_BlockNode)
                    work = thenBlock.FP.GetBreakKind(checkMode)
                    
                    if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                        if work == Nodes_BreakKind__Return{
                            return Nodes_BreakKind__Return
                        }
                        if work == Nodes_BreakKind__NeverRet{
                            return Nodes_BreakKind__NeverRet
                        }
                    } else { 
                        if _switch26251 := work; _switch26251 == Nodes_BreakKind__None {
                            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                                Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                                if true{
                                    return Nodes_BreakKind__None
                                }
                            }
                        } else {
                            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                                Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                                kind = work
                                
                            }
                        }
                    }
                    
                    {
                        _syncBlock := self.syncBlock
                        if _syncBlock != nil {
                            syncBlock := _syncBlock.(*Nodes_BlockNode)
                            work = syncBlock.FP.GetBreakKind(checkMode)
                            
                            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                                if work == Nodes_BreakKind__Return{
                                    return Nodes_BreakKind__Return
                                }
                                if work == Nodes_BreakKind__NeverRet{
                                    return Nodes_BreakKind__NeverRet
                                }
                            } else { 
                                if _switch26364 := work; _switch26364 == Nodes_BreakKind__None {
                                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                                        if true{
                                            return Nodes_BreakKind__None
                                        }
                                    }
                                } else {
                                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                                        Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                                        Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                                        kind = work
                                        
                                    }
                                }
                            }
                            
                        }
                    }
                    return kind
                }
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
                Lns_GetEnv().SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
                return kind
            }
        }
    }
    return Nodes_BreakKind__None
}


// declaration Class -- DeclFormNode
type Nodes_DeclFormNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_argList() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclFormNode struct {
    Nodes_Node
    argList *LnsList
    FP Nodes_DeclFormNodeMtd
}
func Nodes_DeclFormNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclFormNode).FP
}
type Nodes_DeclFormNodeDownCast interface {
    ToNodes_DeclFormNode() *Nodes_DeclFormNode
}
func Nodes_DeclFormNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclFormNodeDownCast)
    if ok { return work.ToNodes_DeclFormNode() }
    return nil
}
func (obj *Nodes_DeclFormNode) ToNodes_DeclFormNode() *Nodes_DeclFormNode {
    return obj
}
func NewNodes_DeclFormNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList) *Nodes_DeclFormNode {
    obj := &Nodes_DeclFormNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclFormNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_DeclFormNode) Get_argList() *LnsList{ return self.argList }
// 1: decl @lune.@base.@Nodes.DeclFormNode.processFilter
func (self *Nodes_DeclFormNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclForm(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclFormNode.canBeRight
func (self *Nodes_DeclFormNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclFormNode.canBeLeft
func (self *Nodes_DeclFormNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclFormNode.canBeStatement
func (self *Nodes_DeclFormNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclFormNode) InitNodes_DeclFormNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,argList *LnsList) {
    self.InitNodes_Node(id, 50, pos, macroArgFlag, typeList)
    self.argList = argList
    
}

// 681: decl @lune.@base.@Nodes.DeclFormNode.create
func Nodes_DeclFormNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,argList *LnsList) *Nodes_DeclFormNode {
    var node *Nodes_DeclFormNode
    node = NewNodes_DeclFormNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, argList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclFormNode.visit
func (self *Nodes_DeclFormNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.argList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch26743 := visitor(child, &self.Nodes_Node, "argList", depth); _switch26743 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch26743 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}


// declaration Class -- DeclFuncInfo
type Nodes_DeclFuncInfoMtd interface {
    Get_accessMode() LnsInt
    Get_argList() *LnsList
    Get_body() LnsAny
    Get_classTypeInfo() LnsAny
    Get_declClassNode() LnsAny
    Get_has__func__Symbol() bool
    Get_kind() LnsInt
    Get_name() LnsAny
    Get_overrideFlag() bool
    Get_retTypeInfoList() *LnsList
    Get_staticFlag() bool
    Get_symbol() LnsAny
}
type Nodes_DeclFuncInfo struct {
    kind LnsInt
    classTypeInfo LnsAny
    declClassNode LnsAny
    name LnsAny
    symbol LnsAny
    argList *LnsList
    staticFlag bool
    accessMode LnsInt
    body LnsAny
    retTypeInfoList *LnsList
    has__func__Symbol bool
    overrideFlag bool
    FP Nodes_DeclFuncInfoMtd
}
func Nodes_DeclFuncInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclFuncInfo).FP
}
type Nodes_DeclFuncInfoDownCast interface {
    ToNodes_DeclFuncInfo() *Nodes_DeclFuncInfo
}
func Nodes_DeclFuncInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclFuncInfoDownCast)
    if ok { return work.ToNodes_DeclFuncInfo() }
    return nil
}
func (obj *Nodes_DeclFuncInfo) ToNodes_DeclFuncInfo() *Nodes_DeclFuncInfo {
    return obj
}
func NewNodes_DeclFuncInfo(arg1 LnsInt, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 *LnsList, arg7 bool, arg8 LnsInt, arg9 LnsAny, arg10 *LnsList, arg11 bool, arg12 bool) *Nodes_DeclFuncInfo {
    obj := &Nodes_DeclFuncInfo{}
    obj.FP = obj
    obj.InitNodes_DeclFuncInfo(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
    return obj
}
func (self *Nodes_DeclFuncInfo) InitNodes_DeclFuncInfo(arg1 LnsInt, arg2 LnsAny, arg3 LnsAny, arg4 LnsAny, arg5 LnsAny, arg6 *LnsList, arg7 bool, arg8 LnsInt, arg9 LnsAny, arg10 *LnsList, arg11 bool, arg12 bool) {
    self.kind = arg1
    self.classTypeInfo = arg2
    self.declClassNode = arg3
    self.name = arg4
    self.symbol = arg5
    self.argList = arg6
    self.staticFlag = arg7
    self.accessMode = arg8
    self.body = arg9
    self.retTypeInfoList = arg10
    self.has__func__Symbol = arg11
    self.overrideFlag = arg12
}
func (self *Nodes_DeclFuncInfo) Get_kind() LnsInt{ return self.kind }
func (self *Nodes_DeclFuncInfo) Get_classTypeInfo() LnsAny{ return self.classTypeInfo }
func (self *Nodes_DeclFuncInfo) Get_declClassNode() LnsAny{ return self.declClassNode }
func (self *Nodes_DeclFuncInfo) Get_name() LnsAny{ return self.name }
func (self *Nodes_DeclFuncInfo) Get_symbol() LnsAny{ return self.symbol }
func (self *Nodes_DeclFuncInfo) Get_argList() *LnsList{ return self.argList }
func (self *Nodes_DeclFuncInfo) Get_staticFlag() bool{ return self.staticFlag }
func (self *Nodes_DeclFuncInfo) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Nodes_DeclFuncInfo) Get_body() LnsAny{ return self.body }
func (self *Nodes_DeclFuncInfo) Get_retTypeInfoList() *LnsList{ return self.retTypeInfoList }
func (self *Nodes_DeclFuncInfo) Get_has__func__Symbol() bool{ return self.has__func__Symbol }
func (self *Nodes_DeclFuncInfo) Get_overrideFlag() bool{ return self.overrideFlag }
// 1751: decl @lune.@base.@Nodes.DeclFuncInfo.createFrom
func Nodes_DeclFuncInfo_createFrom(info *Nodes_DeclFuncInfo,name *Types_Token,symbol *Ast_SymbolInfo) *Nodes_DeclFuncInfo {
    return NewNodes_DeclFuncInfo(info.FP.Get_kind(), info.classTypeInfo, info.declClassNode, name, symbol, info.argList, info.staticFlag, info.accessMode, info.body, info.retTypeInfoList, info.has__func__Symbol, info.overrideFlag)
}


// declaration Class -- DeclFuncNode
type Nodes_DeclFuncNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_declInfo() *Nodes_DeclFuncInfo
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclFuncNode struct {
    Nodes_Node
    declInfo *Nodes_DeclFuncInfo
    FP Nodes_DeclFuncNodeMtd
}
func Nodes_DeclFuncNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclFuncNode).FP
}
type Nodes_DeclFuncNodeDownCast interface {
    ToNodes_DeclFuncNode() *Nodes_DeclFuncNode
}
func Nodes_DeclFuncNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclFuncNodeDownCast)
    if ok { return work.ToNodes_DeclFuncNode() }
    return nil
}
func (obj *Nodes_DeclFuncNode) ToNodes_DeclFuncNode() *Nodes_DeclFuncNode {
    return obj
}
func NewNodes_DeclFuncNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_DeclFuncInfo) *Nodes_DeclFuncNode {
    obj := &Nodes_DeclFuncNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclFuncNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_DeclFuncNode) Get_declInfo() *Nodes_DeclFuncInfo{ return self.declInfo }
// 1: decl @lune.@base.@Nodes.DeclFuncNode.processFilter
func (self *Nodes_DeclFuncNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclFunc(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclFuncNode.canBeLeft
func (self *Nodes_DeclFuncNode) CanBeLeft() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_DeclFuncNode) InitNodes_DeclFuncNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(id, 51, pos, macroArgFlag, typeList)
    self.declInfo = declInfo
    
}

// 681: decl @lune.@base.@Nodes.DeclFuncNode.create
func Nodes_DeclFuncNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_DeclFuncNode {
    var node *Nodes_DeclFuncNode
    node = NewNodes_DeclFuncNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclFuncNode.visit
func (self *Nodes_DeclFuncNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _body := self.declInfo.FP.Get_body()
            if _body != nil {
                body := _body.(*Nodes_BlockNode)
                if _switch27177 := visitor(&body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch27177 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(body.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch27177 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1764: decl @lune.@base.@Nodes.DeclFuncNode.canBeRight
func (self *Nodes_DeclFuncNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return Types_Token2Stem(self.declInfo.FP.Get_name()) == nil
}

// 1768: decl @lune.@base.@Nodes.DeclFuncNode.canBeStatement
func (self *Nodes_DeclFuncNode) CanBeStatement() bool {
    return Lns_op_not((Types_Token2Stem(self.declInfo.FP.Get_name()) == nil))
}


// declaration Class -- DeclMethodNode
type Nodes_DeclMethodNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_declInfo() *Nodes_DeclFuncInfo
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclMethodNode struct {
    Nodes_Node
    declInfo *Nodes_DeclFuncInfo
    FP Nodes_DeclMethodNodeMtd
}
func Nodes_DeclMethodNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclMethodNode).FP
}
type Nodes_DeclMethodNodeDownCast interface {
    ToNodes_DeclMethodNode() *Nodes_DeclMethodNode
}
func Nodes_DeclMethodNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclMethodNodeDownCast)
    if ok { return work.ToNodes_DeclMethodNode() }
    return nil
}
func (obj *Nodes_DeclMethodNode) ToNodes_DeclMethodNode() *Nodes_DeclMethodNode {
    return obj
}
func NewNodes_DeclMethodNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_DeclFuncInfo) *Nodes_DeclMethodNode {
    obj := &Nodes_DeclMethodNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclMethodNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_DeclMethodNode) Get_declInfo() *Nodes_DeclFuncInfo{ return self.declInfo }
// 1: decl @lune.@base.@Nodes.DeclMethodNode.processFilter
func (self *Nodes_DeclMethodNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclMethod(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclMethodNode.canBeRight
func (self *Nodes_DeclMethodNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclMethodNode.canBeLeft
func (self *Nodes_DeclMethodNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclMethodNode.canBeStatement
func (self *Nodes_DeclMethodNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclMethodNode) InitNodes_DeclMethodNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(id, 52, pos, macroArgFlag, typeList)
    self.declInfo = declInfo
    
}

// 681: decl @lune.@base.@Nodes.DeclMethodNode.create
func Nodes_DeclMethodNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_DeclMethodNode {
    var node *Nodes_DeclMethodNode
    node = NewNodes_DeclMethodNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclMethodNode.visit
func (self *Nodes_DeclMethodNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _body := self.declInfo.FP.Get_body()
            if _body != nil {
                body := _body.(*Nodes_BlockNode)
                if _switch27562 := visitor(&body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch27562 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(body.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch27562 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- ProtoMethodNode
type Nodes_ProtoMethodNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_declInfo() *Nodes_DeclFuncInfo
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ProtoMethodNode struct {
    Nodes_Node
    declInfo *Nodes_DeclFuncInfo
    FP Nodes_ProtoMethodNodeMtd
}
func Nodes_ProtoMethodNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ProtoMethodNode).FP
}
type Nodes_ProtoMethodNodeDownCast interface {
    ToNodes_ProtoMethodNode() *Nodes_ProtoMethodNode
}
func Nodes_ProtoMethodNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ProtoMethodNodeDownCast)
    if ok { return work.ToNodes_ProtoMethodNode() }
    return nil
}
func (obj *Nodes_ProtoMethodNode) ToNodes_ProtoMethodNode() *Nodes_ProtoMethodNode {
    return obj
}
func NewNodes_ProtoMethodNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_DeclFuncInfo) *Nodes_ProtoMethodNode {
    obj := &Nodes_ProtoMethodNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ProtoMethodNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ProtoMethodNode) Get_declInfo() *Nodes_DeclFuncInfo{ return self.declInfo }
// 1: decl @lune.@base.@Nodes.ProtoMethodNode.processFilter
func (self *Nodes_ProtoMethodNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessProtoMethod(self, opt)
}

// 1: decl @lune.@base.@Nodes.ProtoMethodNode.canBeRight
func (self *Nodes_ProtoMethodNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ProtoMethodNode.canBeLeft
func (self *Nodes_ProtoMethodNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ProtoMethodNode.canBeStatement
func (self *Nodes_ProtoMethodNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ProtoMethodNode) InitNodes_ProtoMethodNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(id, 53, pos, macroArgFlag, typeList)
    self.declInfo = declInfo
    
}

// 681: decl @lune.@base.@Nodes.ProtoMethodNode.create
func Nodes_ProtoMethodNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_ProtoMethodNode {
    var node *Nodes_ProtoMethodNode
    node = NewNodes_ProtoMethodNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ProtoMethodNode.visit
func (self *Nodes_ProtoMethodNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _body := self.declInfo.FP.Get_body()
            if _body != nil {
                body := _body.(*Nodes_BlockNode)
                if _switch27909 := visitor(&body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch27909 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(body.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch27909 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- DeclConstrNode
type Nodes_DeclConstrNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_declInfo() *Nodes_DeclFuncInfo
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclConstrNode struct {
    Nodes_Node
    declInfo *Nodes_DeclFuncInfo
    FP Nodes_DeclConstrNodeMtd
}
func Nodes_DeclConstrNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclConstrNode).FP
}
type Nodes_DeclConstrNodeDownCast interface {
    ToNodes_DeclConstrNode() *Nodes_DeclConstrNode
}
func Nodes_DeclConstrNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclConstrNodeDownCast)
    if ok { return work.ToNodes_DeclConstrNode() }
    return nil
}
func (obj *Nodes_DeclConstrNode) ToNodes_DeclConstrNode() *Nodes_DeclConstrNode {
    return obj
}
func NewNodes_DeclConstrNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_DeclFuncInfo) *Nodes_DeclConstrNode {
    obj := &Nodes_DeclConstrNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclConstrNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_DeclConstrNode) Get_declInfo() *Nodes_DeclFuncInfo{ return self.declInfo }
// 1: decl @lune.@base.@Nodes.DeclConstrNode.processFilter
func (self *Nodes_DeclConstrNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclConstr(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclConstrNode.canBeRight
func (self *Nodes_DeclConstrNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclConstrNode.canBeLeft
func (self *Nodes_DeclConstrNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclConstrNode.canBeStatement
func (self *Nodes_DeclConstrNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclConstrNode) InitNodes_DeclConstrNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(id, 54, pos, macroArgFlag, typeList)
    self.declInfo = declInfo
    
}

// 681: decl @lune.@base.@Nodes.DeclConstrNode.create
func Nodes_DeclConstrNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_DeclConstrNode {
    var node *Nodes_DeclConstrNode
    node = NewNodes_DeclConstrNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclConstrNode.visit
func (self *Nodes_DeclConstrNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _body := self.declInfo.FP.Get_body()
            if _body != nil {
                body := _body.(*Nodes_BlockNode)
                if _switch28256 := visitor(&body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch28256 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(body.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch28256 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- DeclDestrNode
type Nodes_DeclDestrNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_declInfo() *Nodes_DeclFuncInfo
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclDestrNode struct {
    Nodes_Node
    declInfo *Nodes_DeclFuncInfo
    FP Nodes_DeclDestrNodeMtd
}
func Nodes_DeclDestrNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclDestrNode).FP
}
type Nodes_DeclDestrNodeDownCast interface {
    ToNodes_DeclDestrNode() *Nodes_DeclDestrNode
}
func Nodes_DeclDestrNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclDestrNodeDownCast)
    if ok { return work.ToNodes_DeclDestrNode() }
    return nil
}
func (obj *Nodes_DeclDestrNode) ToNodes_DeclDestrNode() *Nodes_DeclDestrNode {
    return obj
}
func NewNodes_DeclDestrNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_DeclFuncInfo) *Nodes_DeclDestrNode {
    obj := &Nodes_DeclDestrNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclDestrNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_DeclDestrNode) Get_declInfo() *Nodes_DeclFuncInfo{ return self.declInfo }
// 1: decl @lune.@base.@Nodes.DeclDestrNode.processFilter
func (self *Nodes_DeclDestrNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclDestr(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclDestrNode.canBeRight
func (self *Nodes_DeclDestrNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclDestrNode.canBeLeft
func (self *Nodes_DeclDestrNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclDestrNode.canBeStatement
func (self *Nodes_DeclDestrNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclDestrNode) InitNodes_DeclDestrNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(id, 55, pos, macroArgFlag, typeList)
    self.declInfo = declInfo
    
}

// 681: decl @lune.@base.@Nodes.DeclDestrNode.create
func Nodes_DeclDestrNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_DeclDestrNode {
    var node *Nodes_DeclDestrNode
    node = NewNodes_DeclDestrNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclDestrNode.visit
func (self *Nodes_DeclDestrNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _body := self.declInfo.FP.Get_body()
            if _body != nil {
                body := _body.(*Nodes_BlockNode)
                if _switch28603 := visitor(&body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch28603 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(body.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch28603 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- ExpCallSuperCtorNode
type Nodes_ExpCallSuperCtorNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() LnsAny
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_methodType() *Ast_TypeInfo
    Get_pos() *Types_Position
    Get_superType() *Ast_TypeInfo
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpCallSuperCtorNode struct {
    Nodes_Node
    superType *Ast_TypeInfo
    methodType *Ast_TypeInfo
    expList LnsAny
    FP Nodes_ExpCallSuperCtorNodeMtd
}
func Nodes_ExpCallSuperCtorNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpCallSuperCtorNode).FP
}
type Nodes_ExpCallSuperCtorNodeDownCast interface {
    ToNodes_ExpCallSuperCtorNode() *Nodes_ExpCallSuperCtorNode
}
func Nodes_ExpCallSuperCtorNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpCallSuperCtorNodeDownCast)
    if ok { return work.ToNodes_ExpCallSuperCtorNode() }
    return nil
}
func (obj *Nodes_ExpCallSuperCtorNode) ToNodes_ExpCallSuperCtorNode() *Nodes_ExpCallSuperCtorNode {
    return obj
}
func NewNodes_ExpCallSuperCtorNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Ast_TypeInfo, arg6 *Ast_TypeInfo, arg7 LnsAny) *Nodes_ExpCallSuperCtorNode {
    obj := &Nodes_ExpCallSuperCtorNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpCallSuperCtorNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpCallSuperCtorNode) Get_superType() *Ast_TypeInfo{ return self.superType }
func (self *Nodes_ExpCallSuperCtorNode) Get_methodType() *Ast_TypeInfo{ return self.methodType }
func (self *Nodes_ExpCallSuperCtorNode) Get_expList() LnsAny{ return self.expList }
// 1: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.processFilter
func (self *Nodes_ExpCallSuperCtorNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCallSuperCtor(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.canBeRight
func (self *Nodes_ExpCallSuperCtorNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.canBeLeft
func (self *Nodes_ExpCallSuperCtorNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.canBeStatement
func (self *Nodes_ExpCallSuperCtorNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ExpCallSuperCtorNode) InitNodes_ExpCallSuperCtorNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,superType *Ast_TypeInfo,methodType *Ast_TypeInfo,expList LnsAny) {
    self.InitNodes_Node(id, 56, pos, macroArgFlag, typeList)
    self.superType = superType
    
    self.methodType = methodType
    
    self.expList = expList
    
}

// 681: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.create
func Nodes_ExpCallSuperCtorNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,superType *Ast_TypeInfo,methodType *Ast_TypeInfo,expList LnsAny) *Nodes_ExpCallSuperCtorNode {
    var node *Nodes_ExpCallSuperCtorNode
    node = NewNodes_ExpCallSuperCtorNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, superType, methodType, expList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.visit
func (self *Nodes_ExpCallSuperCtorNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch29009 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch29009 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch29009 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- ExpCallSuperNode
type Nodes_ExpCallSuperNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() LnsAny
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_methodType() *Ast_TypeInfo
    Get_pos() *Types_Position
    Get_superType() *Ast_TypeInfo
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ExpCallSuperNode struct {
    Nodes_Node
    superType *Ast_TypeInfo
    methodType *Ast_TypeInfo
    expList LnsAny
    FP Nodes_ExpCallSuperNodeMtd
}
func Nodes_ExpCallSuperNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExpCallSuperNode).FP
}
type Nodes_ExpCallSuperNodeDownCast interface {
    ToNodes_ExpCallSuperNode() *Nodes_ExpCallSuperNode
}
func Nodes_ExpCallSuperNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExpCallSuperNodeDownCast)
    if ok { return work.ToNodes_ExpCallSuperNode() }
    return nil
}
func (obj *Nodes_ExpCallSuperNode) ToNodes_ExpCallSuperNode() *Nodes_ExpCallSuperNode {
    return obj
}
func NewNodes_ExpCallSuperNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Ast_TypeInfo, arg6 *Ast_TypeInfo, arg7 LnsAny) *Nodes_ExpCallSuperNode {
    obj := &Nodes_ExpCallSuperNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpCallSuperNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpCallSuperNode) Get_superType() *Ast_TypeInfo{ return self.superType }
func (self *Nodes_ExpCallSuperNode) Get_methodType() *Ast_TypeInfo{ return self.methodType }
func (self *Nodes_ExpCallSuperNode) Get_expList() LnsAny{ return self.expList }
// 1: decl @lune.@base.@Nodes.ExpCallSuperNode.processFilter
func (self *Nodes_ExpCallSuperNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCallSuper(self, opt)
}

// 1: decl @lune.@base.@Nodes.ExpCallSuperNode.canBeLeft
func (self *Nodes_ExpCallSuperNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ExpCallSuperNode.canBeStatement
func (self *Nodes_ExpCallSuperNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ExpCallSuperNode) InitNodes_ExpCallSuperNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,superType *Ast_TypeInfo,methodType *Ast_TypeInfo,expList LnsAny) {
    self.InitNodes_Node(id, 57, pos, macroArgFlag, typeList)
    self.superType = superType
    
    self.methodType = methodType
    
    self.expList = expList
    
}

// 681: decl @lune.@base.@Nodes.ExpCallSuperNode.create
func Nodes_ExpCallSuperNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,superType *Ast_TypeInfo,methodType *Ast_TypeInfo,expList LnsAny) *Nodes_ExpCallSuperNode {
    var node *Nodes_ExpCallSuperNode
    node = NewNodes_ExpCallSuperNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, superType, methodType, expList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ExpCallSuperNode.visit
func (self *Nodes_ExpCallSuperNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch29403 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch29403 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch29403 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1802: decl @lune.@base.@Nodes.ExpCallSuperNode.canBeRight
func (self *Nodes_ExpCallSuperNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return self.FP.Get_expType() != Ast_builtinTypeNone
}


// declaration Class -- DeclMemberNode
type Nodes_DeclMemberNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_accessMode() LnsInt
    Get_classType() *Ast_TypeInfo
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_getterMode() LnsInt
    Get_getterMutable() bool
    Get_getterRetType() *Ast_TypeInfo
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_name() *Types_Token
    Get_pos() *Types_Position
    Get_refType() *Nodes_RefTypeNode
    Get_setterMode() LnsInt
    Get_staticFlag() bool
    Get_symbolInfo() *Ast_SymbolInfo
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclMemberNode struct {
    Nodes_Node
    name *Types_Token
    refType *Nodes_RefTypeNode
    symbolInfo *Ast_SymbolInfo
    classType *Ast_TypeInfo
    staticFlag bool
    accessMode LnsInt
    getterMutable bool
    getterMode LnsInt
    getterRetType *Ast_TypeInfo
    setterMode LnsInt
    FP Nodes_DeclMemberNodeMtd
}
func Nodes_DeclMemberNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclMemberNode).FP
}
type Nodes_DeclMemberNodeDownCast interface {
    ToNodes_DeclMemberNode() *Nodes_DeclMemberNode
}
func Nodes_DeclMemberNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclMemberNodeDownCast)
    if ok { return work.ToNodes_DeclMemberNode() }
    return nil
}
func (obj *Nodes_DeclMemberNode) ToNodes_DeclMemberNode() *Nodes_DeclMemberNode {
    return obj
}
func NewNodes_DeclMemberNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 *Nodes_RefTypeNode, arg7 *Ast_SymbolInfo, arg8 *Ast_TypeInfo, arg9 bool, arg10 LnsInt, arg11 bool, arg12 LnsInt, arg13 *Ast_TypeInfo, arg14 LnsInt) *Nodes_DeclMemberNode {
    obj := &Nodes_DeclMemberNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclMemberNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14)
    return obj
}
func (self *Nodes_DeclMemberNode) Get_name() *Types_Token{ return self.name }
func (self *Nodes_DeclMemberNode) Get_refType() *Nodes_RefTypeNode{ return self.refType }
func (self *Nodes_DeclMemberNode) Get_symbolInfo() *Ast_SymbolInfo{ return self.symbolInfo }
func (self *Nodes_DeclMemberNode) Get_classType() *Ast_TypeInfo{ return self.classType }
func (self *Nodes_DeclMemberNode) Get_staticFlag() bool{ return self.staticFlag }
func (self *Nodes_DeclMemberNode) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Nodes_DeclMemberNode) Get_getterMutable() bool{ return self.getterMutable }
func (self *Nodes_DeclMemberNode) Get_getterMode() LnsInt{ return self.getterMode }
func (self *Nodes_DeclMemberNode) Get_getterRetType() *Ast_TypeInfo{ return self.getterRetType }
func (self *Nodes_DeclMemberNode) Get_setterMode() LnsInt{ return self.setterMode }
// 1: decl @lune.@base.@Nodes.DeclMemberNode.processFilter
func (self *Nodes_DeclMemberNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclMember(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclMemberNode.canBeRight
func (self *Nodes_DeclMemberNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclMemberNode.canBeLeft
func (self *Nodes_DeclMemberNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclMemberNode.canBeStatement
func (self *Nodes_DeclMemberNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclMemberNode) InitNodes_DeclMemberNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token,refType *Nodes_RefTypeNode,symbolInfo *Ast_SymbolInfo,classType *Ast_TypeInfo,staticFlag bool,accessMode LnsInt,getterMutable bool,getterMode LnsInt,getterRetType *Ast_TypeInfo,setterMode LnsInt) {
    self.InitNodes_Node(id, 58, pos, macroArgFlag, typeList)
    self.name = name
    
    self.refType = refType
    
    self.symbolInfo = symbolInfo
    
    self.classType = classType
    
    self.staticFlag = staticFlag
    
    self.accessMode = accessMode
    
    self.getterMutable = getterMutable
    
    self.getterMode = getterMode
    
    self.getterRetType = getterRetType
    
    self.setterMode = setterMode
    
}

// 681: decl @lune.@base.@Nodes.DeclMemberNode.create
func Nodes_DeclMemberNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token,refType *Nodes_RefTypeNode,symbolInfo *Ast_SymbolInfo,classType *Ast_TypeInfo,staticFlag bool,accessMode LnsInt,getterMutable bool,getterMode LnsInt,getterRetType *Ast_TypeInfo,setterMode LnsInt) *Nodes_DeclMemberNode {
    var node *Nodes_DeclMemberNode
    node = NewNodes_DeclMemberNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, name, refType, symbolInfo, classType, staticFlag, accessMode, getterMutable, getterMode, getterRetType, setterMode)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclMemberNode.visit
func (self *Nodes_DeclMemberNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_RefTypeNode
        child = self.refType
        if _switch30033 := visitor(&child.Nodes_Node, &self.Nodes_Node, "refType", depth); _switch30033 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch30033 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- DeclArgNode
type Nodes_DeclArgNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_argType() LnsAny
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_name() *Types_Token
    Get_pos() *Types_Position
    Get_symbolInfo() *Ast_SymbolInfo
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclArgNode struct {
    Nodes_Node
    name *Types_Token
    symbolInfo *Ast_SymbolInfo
    argType LnsAny
    FP Nodes_DeclArgNodeMtd
}
func Nodes_DeclArgNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclArgNode).FP
}
type Nodes_DeclArgNodeDownCast interface {
    ToNodes_DeclArgNode() *Nodes_DeclArgNode
}
func Nodes_DeclArgNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclArgNodeDownCast)
    if ok { return work.ToNodes_DeclArgNode() }
    return nil
}
func (obj *Nodes_DeclArgNode) ToNodes_DeclArgNode() *Nodes_DeclArgNode {
    return obj
}
func NewNodes_DeclArgNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 *Ast_SymbolInfo, arg7 LnsAny) *Nodes_DeclArgNode {
    obj := &Nodes_DeclArgNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclArgNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_DeclArgNode) Get_name() *Types_Token{ return self.name }
func (self *Nodes_DeclArgNode) Get_symbolInfo() *Ast_SymbolInfo{ return self.symbolInfo }
func (self *Nodes_DeclArgNode) Get_argType() LnsAny{ return self.argType }
// 1: decl @lune.@base.@Nodes.DeclArgNode.processFilter
func (self *Nodes_DeclArgNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclArg(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclArgNode.canBeRight
func (self *Nodes_DeclArgNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclArgNode.canBeLeft
func (self *Nodes_DeclArgNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclArgNode.canBeStatement
func (self *Nodes_DeclArgNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_DeclArgNode) InitNodes_DeclArgNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token,symbolInfo *Ast_SymbolInfo,argType LnsAny) {
    self.InitNodes_Node(id, 59, pos, macroArgFlag, typeList)
    self.name = name
    
    self.symbolInfo = symbolInfo
    
    self.argType = argType
    
}

// 681: decl @lune.@base.@Nodes.DeclArgNode.create
func Nodes_DeclArgNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token,symbolInfo *Ast_SymbolInfo,argType LnsAny) *Nodes_DeclArgNode {
    var node *Nodes_DeclArgNode
    node = NewNodes_DeclArgNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, name, symbolInfo, argType)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclArgNode.visit
func (self *Nodes_DeclArgNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.argType
            if _child != nil {
                child := _child.(*Nodes_RefTypeNode)
                if _switch30438 := visitor(&child.Nodes_Node, &self.Nodes_Node, "argType", depth); _switch30438 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch30438 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}


// declaration Class -- DeclArgDDDNode
type Nodes_DeclArgDDDNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclArgDDDNode struct {
    Nodes_Node
    FP Nodes_DeclArgDDDNodeMtd
}
func Nodes_DeclArgDDDNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclArgDDDNode).FP
}
type Nodes_DeclArgDDDNodeDownCast interface {
    ToNodes_DeclArgDDDNode() *Nodes_DeclArgDDDNode
}
func Nodes_DeclArgDDDNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclArgDDDNodeDownCast)
    if ok { return work.ToNodes_DeclArgDDDNode() }
    return nil
}
func (obj *Nodes_DeclArgDDDNode) ToNodes_DeclArgDDDNode() *Nodes_DeclArgDDDNode {
    return obj
}
func NewNodes_DeclArgDDDNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList) *Nodes_DeclArgDDDNode {
    obj := &Nodes_DeclArgDDDNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclArgDDDNode(arg1, arg2, arg3, arg4)
    return obj
}
// 1: decl @lune.@base.@Nodes.DeclArgDDDNode.processFilter
func (self *Nodes_DeclArgDDDNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclArgDDD(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclArgDDDNode.canBeRight
func (self *Nodes_DeclArgDDDNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclArgDDDNode.canBeLeft
func (self *Nodes_DeclArgDDDNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclArgDDDNode.canBeStatement
func (self *Nodes_DeclArgDDDNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_DeclArgDDDNode) InitNodes_DeclArgDDDNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(id, 60, pos, macroArgFlag, typeList)
}

// 681: decl @lune.@base.@Nodes.DeclArgDDDNode.create
func Nodes_DeclArgDDDNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList) *Nodes_DeclArgDDDNode {
    var node *Nodes_DeclArgDDDNode
    node = NewNodes_DeclArgDDDNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclArgDDDNode.visit
func (self *Nodes_DeclArgDDDNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- AdvertiseInfo
type Nodes_AdvertiseInfoMtd interface {
    Get_member() *Nodes_DeclMemberNode
    Get_pos() *Types_Position
    Get_prefix() string
}
type Nodes_AdvertiseInfo struct {
    member *Nodes_DeclMemberNode
    prefix string
    pos *Types_Position
    FP Nodes_AdvertiseInfoMtd
}
func Nodes_AdvertiseInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_AdvertiseInfo).FP
}
type Nodes_AdvertiseInfoDownCast interface {
    ToNodes_AdvertiseInfo() *Nodes_AdvertiseInfo
}
func Nodes_AdvertiseInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_AdvertiseInfoDownCast)
    if ok { return work.ToNodes_AdvertiseInfo() }
    return nil
}
func (obj *Nodes_AdvertiseInfo) ToNodes_AdvertiseInfo() *Nodes_AdvertiseInfo {
    return obj
}
func NewNodes_AdvertiseInfo(arg1 *Nodes_DeclMemberNode, arg2 string, arg3 *Types_Position) *Nodes_AdvertiseInfo {
    obj := &Nodes_AdvertiseInfo{}
    obj.FP = obj
    obj.InitNodes_AdvertiseInfo(arg1, arg2, arg3)
    return obj
}
func (self *Nodes_AdvertiseInfo) InitNodes_AdvertiseInfo(arg1 *Nodes_DeclMemberNode, arg2 string, arg3 *Types_Position) {
    self.member = arg1
    self.prefix = arg2
    self.pos = arg3
}
func (self *Nodes_AdvertiseInfo) Get_member() *Nodes_DeclMemberNode{ return self.member }
func (self *Nodes_AdvertiseInfo) Get_prefix() string{ return self.prefix }
func (self *Nodes_AdvertiseInfo) Get_pos() *Types_Position{ return self.pos }

// declaration Class -- DeclAdvertiseNode
type Nodes_DeclAdvertiseNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_advInfo() *Nodes_AdvertiseInfo
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclAdvertiseNode struct {
    Nodes_Node
    advInfo *Nodes_AdvertiseInfo
    FP Nodes_DeclAdvertiseNodeMtd
}
func Nodes_DeclAdvertiseNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclAdvertiseNode).FP
}
type Nodes_DeclAdvertiseNodeDownCast interface {
    ToNodes_DeclAdvertiseNode() *Nodes_DeclAdvertiseNode
}
func Nodes_DeclAdvertiseNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclAdvertiseNodeDownCast)
    if ok { return work.ToNodes_DeclAdvertiseNode() }
    return nil
}
func (obj *Nodes_DeclAdvertiseNode) ToNodes_DeclAdvertiseNode() *Nodes_DeclAdvertiseNode {
    return obj
}
func NewNodes_DeclAdvertiseNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_AdvertiseInfo) *Nodes_DeclAdvertiseNode {
    obj := &Nodes_DeclAdvertiseNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclAdvertiseNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_DeclAdvertiseNode) Get_advInfo() *Nodes_AdvertiseInfo{ return self.advInfo }
// 1: decl @lune.@base.@Nodes.DeclAdvertiseNode.processFilter
func (self *Nodes_DeclAdvertiseNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclAdvertise(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclAdvertiseNode.canBeRight
func (self *Nodes_DeclAdvertiseNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclAdvertiseNode.canBeLeft
func (self *Nodes_DeclAdvertiseNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclAdvertiseNode.canBeStatement
func (self *Nodes_DeclAdvertiseNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_DeclAdvertiseNode) InitNodes_DeclAdvertiseNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,advInfo *Nodes_AdvertiseInfo) {
    self.InitNodes_Node(id, 61, pos, macroArgFlag, typeList)
    self.advInfo = advInfo
    
}

// 681: decl @lune.@base.@Nodes.DeclAdvertiseNode.create
func Nodes_DeclAdvertiseNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,advInfo *Nodes_AdvertiseInfo) *Nodes_DeclAdvertiseNode {
    var node *Nodes_DeclAdvertiseNode
    node = NewNodes_DeclAdvertiseNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, advInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclAdvertiseNode.visit
func (self *Nodes_DeclAdvertiseNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- ProtoClassNode
type Nodes_ProtoClassNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_name() *Types_Token
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_ProtoClassNode struct {
    Nodes_Node
    name *Types_Token
    FP Nodes_ProtoClassNodeMtd
}
func Nodes_ProtoClassNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ProtoClassNode).FP
}
type Nodes_ProtoClassNodeDownCast interface {
    ToNodes_ProtoClassNode() *Nodes_ProtoClassNode
}
func Nodes_ProtoClassNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ProtoClassNodeDownCast)
    if ok { return work.ToNodes_ProtoClassNode() }
    return nil
}
func (obj *Nodes_ProtoClassNode) ToNodes_ProtoClassNode() *Nodes_ProtoClassNode {
    return obj
}
func NewNodes_ProtoClassNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token) *Nodes_ProtoClassNode {
    obj := &Nodes_ProtoClassNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ProtoClassNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_ProtoClassNode) Get_name() *Types_Token{ return self.name }
// 1: decl @lune.@base.@Nodes.ProtoClassNode.processFilter
func (self *Nodes_ProtoClassNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessProtoClass(self, opt)
}

// 1: decl @lune.@base.@Nodes.ProtoClassNode.canBeRight
func (self *Nodes_ProtoClassNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ProtoClassNode.canBeLeft
func (self *Nodes_ProtoClassNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.ProtoClassNode.canBeStatement
func (self *Nodes_ProtoClassNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_ProtoClassNode) InitNodes_ProtoClassNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token) {
    self.InitNodes_Node(id, 62, pos, macroArgFlag, typeList)
    self.name = name
    
}

// 681: decl @lune.@base.@Nodes.ProtoClassNode.create
func Nodes_ProtoClassNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token) *Nodes_ProtoClassNode {
    var node *Nodes_ProtoClassNode
    node = NewNodes_ProtoClassNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, name)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.ProtoClassNode.visit
func (self *Nodes_ProtoClassNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- ClassInitBlockInfo
type Nodes_ClassInitBlockInfoMtd interface {
    Get_func() LnsAny
    Set_func(arg1 LnsAny)
}
type Nodes_ClassInitBlockInfo struct {
    _func LnsAny
    FP Nodes_ClassInitBlockInfoMtd
}
func Nodes_ClassInitBlockInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ClassInitBlockInfo).FP
}
type Nodes_ClassInitBlockInfoDownCast interface {
    ToNodes_ClassInitBlockInfo() *Nodes_ClassInitBlockInfo
}
func Nodes_ClassInitBlockInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ClassInitBlockInfoDownCast)
    if ok { return work.ToNodes_ClassInitBlockInfo() }
    return nil
}
func (obj *Nodes_ClassInitBlockInfo) ToNodes_ClassInitBlockInfo() *Nodes_ClassInitBlockInfo {
    return obj
}
func NewNodes_ClassInitBlockInfo(arg1 LnsAny) *Nodes_ClassInitBlockInfo {
    obj := &Nodes_ClassInitBlockInfo{}
    obj.FP = obj
    obj.InitNodes_ClassInitBlockInfo(arg1)
    return obj
}
func (self *Nodes_ClassInitBlockInfo) InitNodes_ClassInitBlockInfo(arg1 LnsAny) {
    self._func = arg1
}
func (self *Nodes_ClassInitBlockInfo) Get_func() LnsAny{ return self._func }
func (self *Nodes_ClassInitBlockInfo) Set_func(arg1 LnsAny){ self._func = arg1 }

// declaration Class -- DeclClassNode
type Nodes_DeclClassNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    CreateMethodNameSetWithoutAdv() *LnsSet
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_accessMode() LnsInt
    Get_advertiseList() *LnsList
    Get_allStmtList() *LnsList
    Get_commentList() LnsAny
    Get_declStmtList() *LnsList
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_fieldList() *LnsList
    Get_gluePrefix() LnsAny
    Get_hasOldCtor() bool
    Get_hasPrototype() bool
    Get_id() LnsInt
    Get_initBlock() *Nodes_ClassInitBlockInfo
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_lazyLoad() LnsInt
    Get_macroArgFlag() bool
    Get_memberList() *LnsList
    Get_moduleName() LnsAny
    Get_name() *Types_Token
    Get_outerMethodSet() *LnsSet
    Get_pos() *Types_Position
    Get_scope() *Ast_Scope
    Get_tailComment() LnsAny
    Get_trustList() *LnsList
    Get_uninitMemberList() *LnsList
    HasNilAccess() bool
    HasUserInit() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetHasOldCtor()
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclClassNode struct {
    Nodes_Node
    accessMode LnsInt
    name *Types_Token
    hasPrototype bool
    gluePrefix LnsAny
    moduleName LnsAny
    lazyLoad LnsInt
    hasOldCtor bool
    allStmtList *LnsList
    declStmtList *LnsList
    fieldList *LnsList
    memberList *LnsList
    scope *Ast_Scope
    initBlock *Nodes_ClassInitBlockInfo
    advertiseList *LnsList
    trustList *LnsList
    uninitMemberList *LnsList
    outerMethodSet *LnsSet
    FP Nodes_DeclClassNodeMtd
}
func Nodes_DeclClassNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclClassNode).FP
}
type Nodes_DeclClassNodeDownCast interface {
    ToNodes_DeclClassNode() *Nodes_DeclClassNode
}
func Nodes_DeclClassNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclClassNodeDownCast)
    if ok { return work.ToNodes_DeclClassNode() }
    return nil
}
func (obj *Nodes_DeclClassNode) ToNodes_DeclClassNode() *Nodes_DeclClassNode {
    return obj
}
func NewNodes_DeclClassNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsInt, arg6 *Types_Token, arg7 bool, arg8 LnsAny, arg9 LnsAny, arg10 LnsInt, arg11 bool, arg12 *LnsList, arg13 *LnsList, arg14 *LnsList, arg15 *LnsList, arg16 *Ast_Scope, arg17 *Nodes_ClassInitBlockInfo, arg18 *LnsList, arg19 *LnsList, arg20 *LnsList, arg21 *LnsSet) *Nodes_DeclClassNode {
    obj := &Nodes_DeclClassNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclClassNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20, arg21)
    return obj
}
func (self *Nodes_DeclClassNode) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Nodes_DeclClassNode) Get_name() *Types_Token{ return self.name }
func (self *Nodes_DeclClassNode) Get_hasPrototype() bool{ return self.hasPrototype }
func (self *Nodes_DeclClassNode) Get_gluePrefix() LnsAny{ return self.gluePrefix }
func (self *Nodes_DeclClassNode) Get_moduleName() LnsAny{ return self.moduleName }
func (self *Nodes_DeclClassNode) Get_lazyLoad() LnsInt{ return self.lazyLoad }
func (self *Nodes_DeclClassNode) Get_hasOldCtor() bool{ return self.hasOldCtor }
func (self *Nodes_DeclClassNode) Get_allStmtList() *LnsList{ return self.allStmtList }
func (self *Nodes_DeclClassNode) Get_declStmtList() *LnsList{ return self.declStmtList }
func (self *Nodes_DeclClassNode) Get_fieldList() *LnsList{ return self.fieldList }
func (self *Nodes_DeclClassNode) Get_memberList() *LnsList{ return self.memberList }
func (self *Nodes_DeclClassNode) Get_scope() *Ast_Scope{ return self.scope }
func (self *Nodes_DeclClassNode) Get_initBlock() *Nodes_ClassInitBlockInfo{ return self.initBlock }
func (self *Nodes_DeclClassNode) Get_advertiseList() *LnsList{ return self.advertiseList }
func (self *Nodes_DeclClassNode) Get_trustList() *LnsList{ return self.trustList }
func (self *Nodes_DeclClassNode) Get_uninitMemberList() *LnsList{ return self.uninitMemberList }
func (self *Nodes_DeclClassNode) Get_outerMethodSet() *LnsSet{ return self.outerMethodSet }
// 1: decl @lune.@base.@Nodes.DeclClassNode.processFilter
func (self *Nodes_DeclClassNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclClass(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclClassNode.canBeRight
func (self *Nodes_DeclClassNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclClassNode.canBeLeft
func (self *Nodes_DeclClassNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclClassNode.canBeStatement
func (self *Nodes_DeclClassNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclClassNode) InitNodes_DeclClassNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,accessMode LnsInt,name *Types_Token,hasPrototype bool,gluePrefix LnsAny,moduleName LnsAny,lazyLoad LnsInt,hasOldCtor bool,allStmtList *LnsList,declStmtList *LnsList,fieldList *LnsList,memberList *LnsList,scope *Ast_Scope,initBlock *Nodes_ClassInitBlockInfo,advertiseList *LnsList,trustList *LnsList,uninitMemberList *LnsList,outerMethodSet *LnsSet) {
    self.InitNodes_Node(id, 63, pos, macroArgFlag, typeList)
    self.accessMode = accessMode
    
    self.name = name
    
    self.hasPrototype = hasPrototype
    
    self.gluePrefix = gluePrefix
    
    self.moduleName = moduleName
    
    self.lazyLoad = lazyLoad
    
    self.hasOldCtor = hasOldCtor
    
    self.allStmtList = allStmtList
    
    self.declStmtList = declStmtList
    
    self.fieldList = fieldList
    
    self.memberList = memberList
    
    self.scope = scope
    
    self.initBlock = initBlock
    
    self.advertiseList = advertiseList
    
    self.trustList = trustList
    
    self.uninitMemberList = uninitMemberList
    
    self.outerMethodSet = outerMethodSet
    
}

// 681: decl @lune.@base.@Nodes.DeclClassNode.create
func Nodes_DeclClassNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,accessMode LnsInt,name *Types_Token,hasPrototype bool,gluePrefix LnsAny,moduleName LnsAny,lazyLoad LnsInt,hasOldCtor bool,allStmtList *LnsList,declStmtList *LnsList,fieldList *LnsList,memberList *LnsList,scope *Ast_Scope,initBlock *Nodes_ClassInitBlockInfo,advertiseList *LnsList,trustList *LnsList,uninitMemberList *LnsList,outerMethodSet *LnsSet) *Nodes_DeclClassNode {
    var node *Nodes_DeclClassNode
    node = NewNodes_DeclClassNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, accessMode, name, hasPrototype, gluePrefix, moduleName, lazyLoad, hasOldCtor, allStmtList, declStmtList, fieldList, memberList, scope, initBlock, advertiseList, trustList, uninitMemberList, outerMethodSet)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclClassNode.visit
func (self *Nodes_DeclClassNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.allStmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch32142 := visitor(child, &self.Nodes_Node, "allStmtList", depth); _switch32142 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch32142 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    {
        var list *LnsList
        list = self.declStmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch32203 := visitor(child, &self.Nodes_Node, "declStmtList", depth); _switch32203 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch32203 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    {
        var list *LnsList
        list = self.fieldList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch32264 := visitor(child, &self.Nodes_Node, "fieldList", depth); _switch32264 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch32264 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    {
        var list *LnsList
        list = self.memberList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if _switch32326 := visitor(&child.Nodes_Node, &self.Nodes_Node, "memberList", depth); _switch32326 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch32326 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}

// 1884: decl @lune.@base.@Nodes.DeclClassNode.createMethodNameSetWithoutAdv
func (self *Nodes_DeclClassNode) CreateMethodNameSetWithoutAdv() *LnsSet {
    var methodNameSet *LnsSet
    methodNameSet = NewLnsSet([]LnsAny{})
    if self.FP.Get_expType().FP.Get_kind() != Ast_TypeInfoKind__IF{
        for _, _field := range( self.FP.Get_fieldList().Items ) {
            field := _field.(Nodes_NodeDownCast).ToNodes_Node()
            if field.FP.Get_kind() == Nodes_NodeKind_get_DeclConstr(){
                methodNameSet.Add("__init")
            }
            if field.FP.Get_kind() == Nodes_NodeKind_get_DeclDestr(){
                methodNameSet.Add("__free")
            }
            {
                _methodNode := Nodes_DeclMethodNodeDownCastF(field.FP)
                if _methodNode != nil {
                    methodNode := _methodNode.(*Nodes_DeclMethodNode)
                    var methodNameToken *Types_Token
                    methodNameToken = Lns_unwrap( methodNode.FP.Get_declInfo().FP.Get_name()).(*Types_Token)
                    methodNameSet.Add(methodNameToken.Txt)
                }
            }
        }
        for _, _memberNode := range( self.FP.Get_memberList().Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            var memberName string
            memberName = memberNode.FP.Get_name().Txt
            if memberNode.FP.Get_getterMode() != Ast_AccessMode__None{
                methodNameSet.Add("get_" + memberName)
            }
            if memberNode.FP.Get_setterMode() != Ast_AccessMode__None{
                methodNameSet.Add("set_" + memberName)
            }
        }
    }
    return methodNameSet
}

// 1912: decl @lune.@base.@Nodes.DeclClassNode.setHasOldCtor
func (self *Nodes_DeclClassNode) SetHasOldCtor() {
    self.hasOldCtor = true
    
}

// 1916: decl @lune.@base.@Nodes.DeclClassNode.hasUserInit
func (self *Nodes_DeclClassNode) HasUserInit() bool {
    var scope *Ast_Scope
    scope = Lns_unwrap( self.FP.Get_expType().FP.Get_scope()).(*Ast_Scope)
    var initFuncType *Ast_TypeInfo
    initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField("__init", true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
    return Lns_op_not(initFuncType.FP.Get_autoFlag())
}


// declaration Class -- DeclEnumNode
type Nodes_DeclEnumNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_accessMode() LnsInt
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_enumType() *Ast_EnumTypeInfo
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_name() *Types_Token
    Get_pos() *Types_Position
    Get_scope() *Ast_Scope
    Get_tailComment() LnsAny
    Get_valueNameList() *LnsList
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclEnumNode struct {
    Nodes_Node
    enumType *Ast_EnumTypeInfo
    accessMode LnsInt
    name *Types_Token
    valueNameList *LnsList
    scope *Ast_Scope
    FP Nodes_DeclEnumNodeMtd
}
func Nodes_DeclEnumNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclEnumNode).FP
}
type Nodes_DeclEnumNodeDownCast interface {
    ToNodes_DeclEnumNode() *Nodes_DeclEnumNode
}
func Nodes_DeclEnumNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclEnumNodeDownCast)
    if ok { return work.ToNodes_DeclEnumNode() }
    return nil
}
func (obj *Nodes_DeclEnumNode) ToNodes_DeclEnumNode() *Nodes_DeclEnumNode {
    return obj
}
func NewNodes_DeclEnumNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Ast_EnumTypeInfo, arg6 LnsInt, arg7 *Types_Token, arg8 *LnsList, arg9 *Ast_Scope) *Nodes_DeclEnumNode {
    obj := &Nodes_DeclEnumNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclEnumNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_DeclEnumNode) Get_enumType() *Ast_EnumTypeInfo{ return self.enumType }
func (self *Nodes_DeclEnumNode) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Nodes_DeclEnumNode) Get_name() *Types_Token{ return self.name }
func (self *Nodes_DeclEnumNode) Get_valueNameList() *LnsList{ return self.valueNameList }
func (self *Nodes_DeclEnumNode) Get_scope() *Ast_Scope{ return self.scope }
// 1: decl @lune.@base.@Nodes.DeclEnumNode.processFilter
func (self *Nodes_DeclEnumNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclEnum(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclEnumNode.canBeRight
func (self *Nodes_DeclEnumNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclEnumNode.canBeLeft
func (self *Nodes_DeclEnumNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclEnumNode.canBeStatement
func (self *Nodes_DeclEnumNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclEnumNode) InitNodes_DeclEnumNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,enumType *Ast_EnumTypeInfo,accessMode LnsInt,name *Types_Token,valueNameList *LnsList,scope *Ast_Scope) {
    self.InitNodes_Node(id, 64, pos, macroArgFlag, typeList)
    self.enumType = enumType
    
    self.accessMode = accessMode
    
    self.name = name
    
    self.valueNameList = valueNameList
    
    self.scope = scope
    
}

// 681: decl @lune.@base.@Nodes.DeclEnumNode.create
func Nodes_DeclEnumNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,enumType *Ast_EnumTypeInfo,accessMode LnsInt,name *Types_Token,valueNameList *LnsList,scope *Ast_Scope) *Nodes_DeclEnumNode {
    var node *Nodes_DeclEnumNode
    node = NewNodes_DeclEnumNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, enumType, accessMode, name, valueNameList, scope)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclEnumNode.visit
func (self *Nodes_DeclEnumNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- DeclAlgeNode
type Nodes_DeclAlgeNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_accessMode() LnsInt
    Get_algeType() *Ast_AlgeTypeInfo
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_scope() *Ast_Scope
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclAlgeNode struct {
    Nodes_Node
    accessMode LnsInt
    algeType *Ast_AlgeTypeInfo
    scope *Ast_Scope
    FP Nodes_DeclAlgeNodeMtd
}
func Nodes_DeclAlgeNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclAlgeNode).FP
}
type Nodes_DeclAlgeNodeDownCast interface {
    ToNodes_DeclAlgeNode() *Nodes_DeclAlgeNode
}
func Nodes_DeclAlgeNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclAlgeNodeDownCast)
    if ok { return work.ToNodes_DeclAlgeNode() }
    return nil
}
func (obj *Nodes_DeclAlgeNode) ToNodes_DeclAlgeNode() *Nodes_DeclAlgeNode {
    return obj
}
func NewNodes_DeclAlgeNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsInt, arg6 *Ast_AlgeTypeInfo, arg7 *Ast_Scope) *Nodes_DeclAlgeNode {
    obj := &Nodes_DeclAlgeNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclAlgeNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_DeclAlgeNode) Get_accessMode() LnsInt{ return self.accessMode }
func (self *Nodes_DeclAlgeNode) Get_algeType() *Ast_AlgeTypeInfo{ return self.algeType }
func (self *Nodes_DeclAlgeNode) Get_scope() *Ast_Scope{ return self.scope }
// 1: decl @lune.@base.@Nodes.DeclAlgeNode.processFilter
func (self *Nodes_DeclAlgeNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclAlge(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclAlgeNode.canBeRight
func (self *Nodes_DeclAlgeNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclAlgeNode.canBeLeft
func (self *Nodes_DeclAlgeNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclAlgeNode.canBeStatement
func (self *Nodes_DeclAlgeNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclAlgeNode) InitNodes_DeclAlgeNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,accessMode LnsInt,algeType *Ast_AlgeTypeInfo,scope *Ast_Scope) {
    self.InitNodes_Node(id, 65, pos, macroArgFlag, typeList)
    self.accessMode = accessMode
    
    self.algeType = algeType
    
    self.scope = scope
    
}

// 681: decl @lune.@base.@Nodes.DeclAlgeNode.create
func Nodes_DeclAlgeNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,accessMode LnsInt,algeType *Ast_AlgeTypeInfo,scope *Ast_Scope) *Nodes_DeclAlgeNode {
    var node *Nodes_DeclAlgeNode
    node = NewNodes_DeclAlgeNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, accessMode, algeType, scope)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclAlgeNode.visit
func (self *Nodes_DeclAlgeNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- NewAlgeValNode
type Nodes_NewAlgeValNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_algeTypeInfo() *Ast_AlgeTypeInfo
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_name() *Types_Token
    Get_paramList() *LnsList
    Get_pos() *Types_Position
    Get_prefix() LnsAny
    Get_tailComment() LnsAny
    Get_valInfo() *Ast_AlgeValInfo
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_NewAlgeValNode struct {
    Nodes_Node
    name *Types_Token
    prefix LnsAny
    algeTypeInfo *Ast_AlgeTypeInfo
    valInfo *Ast_AlgeValInfo
    paramList *LnsList
    FP Nodes_NewAlgeValNodeMtd
}
func Nodes_NewAlgeValNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_NewAlgeValNode).FP
}
type Nodes_NewAlgeValNodeDownCast interface {
    ToNodes_NewAlgeValNode() *Nodes_NewAlgeValNode
}
func Nodes_NewAlgeValNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_NewAlgeValNodeDownCast)
    if ok { return work.ToNodes_NewAlgeValNode() }
    return nil
}
func (obj *Nodes_NewAlgeValNode) ToNodes_NewAlgeValNode() *Nodes_NewAlgeValNode {
    return obj
}
func NewNodes_NewAlgeValNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 LnsAny, arg7 *Ast_AlgeTypeInfo, arg8 *Ast_AlgeValInfo, arg9 *LnsList) *Nodes_NewAlgeValNode {
    obj := &Nodes_NewAlgeValNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_NewAlgeValNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_NewAlgeValNode) Get_name() *Types_Token{ return self.name }
func (self *Nodes_NewAlgeValNode) Get_prefix() LnsAny{ return self.prefix }
func (self *Nodes_NewAlgeValNode) Get_algeTypeInfo() *Ast_AlgeTypeInfo{ return self.algeTypeInfo }
func (self *Nodes_NewAlgeValNode) Get_valInfo() *Ast_AlgeValInfo{ return self.valInfo }
func (self *Nodes_NewAlgeValNode) Get_paramList() *LnsList{ return self.paramList }
// 1: decl @lune.@base.@Nodes.NewAlgeValNode.processFilter
func (self *Nodes_NewAlgeValNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessNewAlgeVal(self, opt)
}

// 1: decl @lune.@base.@Nodes.NewAlgeValNode.canBeRight
func (self *Nodes_NewAlgeValNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.NewAlgeValNode.canBeLeft
func (self *Nodes_NewAlgeValNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.NewAlgeValNode.canBeStatement
func (self *Nodes_NewAlgeValNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_NewAlgeValNode) InitNodes_NewAlgeValNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token,prefix LnsAny,algeTypeInfo *Ast_AlgeTypeInfo,valInfo *Ast_AlgeValInfo,paramList *LnsList) {
    self.InitNodes_Node(id, 66, pos, macroArgFlag, typeList)
    self.name = name
    
    self.prefix = prefix
    
    self.algeTypeInfo = algeTypeInfo
    
    self.valInfo = valInfo
    
    self.paramList = paramList
    
}

// 681: decl @lune.@base.@Nodes.NewAlgeValNode.create
func Nodes_NewAlgeValNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token,prefix LnsAny,algeTypeInfo *Ast_AlgeTypeInfo,valInfo *Ast_AlgeValInfo,paramList *LnsList) *Nodes_NewAlgeValNode {
    var node *Nodes_NewAlgeValNode
    node = NewNodes_NewAlgeValNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, name, prefix, algeTypeInfo, valInfo, paramList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.NewAlgeValNode.visit
func (self *Nodes_NewAlgeValNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.prefix
            if _child != nil {
                child := _child.(*Nodes_Node)
                if _switch33766 := visitor(child, &self.Nodes_Node, "prefix", depth); _switch33766 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch33766 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    {
        var list *LnsList
        list = self.paramList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch33826 := visitor(child, &self.Nodes_Node, "paramList", depth); _switch33826 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch33826 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}


// declaration Class -- LuneControlNode
type Nodes_LuneControlNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_pragma() LnsAny
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LuneControlNode struct {
    Nodes_Node
    pragma LnsAny
    FP Nodes_LuneControlNodeMtd
}
func Nodes_LuneControlNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LuneControlNode).FP
}
type Nodes_LuneControlNodeDownCast interface {
    ToNodes_LuneControlNode() *Nodes_LuneControlNode
}
func Nodes_LuneControlNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LuneControlNodeDownCast)
    if ok { return work.ToNodes_LuneControlNode() }
    return nil
}
func (obj *Nodes_LuneControlNode) ToNodes_LuneControlNode() *Nodes_LuneControlNode {
    return obj
}
func NewNodes_LuneControlNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsAny) *Nodes_LuneControlNode {
    obj := &Nodes_LuneControlNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LuneControlNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_LuneControlNode) Get_pragma() LnsAny{ return self.pragma }
// 1: decl @lune.@base.@Nodes.LuneControlNode.processFilter
func (self *Nodes_LuneControlNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLuneControl(self, opt)
}

// 1: decl @lune.@base.@Nodes.LuneControlNode.canBeRight
func (self *Nodes_LuneControlNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LuneControlNode.canBeLeft
func (self *Nodes_LuneControlNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LuneControlNode.canBeStatement
func (self *Nodes_LuneControlNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_LuneControlNode) InitNodes_LuneControlNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,pragma LnsAny) {
    self.InitNodes_Node(id, 67, pos, macroArgFlag, typeList)
    self.pragma = pragma
    
}

// 681: decl @lune.@base.@Nodes.LuneControlNode.create
func Nodes_LuneControlNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,pragma LnsAny) *Nodes_LuneControlNode {
    var node *Nodes_LuneControlNode
    node = NewNodes_LuneControlNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, pragma)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LuneControlNode.visit
func (self *Nodes_LuneControlNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- MatchCase
type Nodes_MatchCaseMtd interface {
    Get_block() *Nodes_BlockNode
    Get_valInfo() *Ast_AlgeValInfo
    Get_valParamNameList() *LnsList
}
type Nodes_MatchCase struct {
    valInfo *Ast_AlgeValInfo
    valParamNameList *LnsList
    block *Nodes_BlockNode
    FP Nodes_MatchCaseMtd
}
func Nodes_MatchCase2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_MatchCase).FP
}
type Nodes_MatchCaseDownCast interface {
    ToNodes_MatchCase() *Nodes_MatchCase
}
func Nodes_MatchCaseDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_MatchCaseDownCast)
    if ok { return work.ToNodes_MatchCase() }
    return nil
}
func (obj *Nodes_MatchCase) ToNodes_MatchCase() *Nodes_MatchCase {
    return obj
}
func NewNodes_MatchCase(arg1 *Ast_AlgeValInfo, arg2 *LnsList, arg3 *Nodes_BlockNode) *Nodes_MatchCase {
    obj := &Nodes_MatchCase{}
    obj.FP = obj
    obj.InitNodes_MatchCase(arg1, arg2, arg3)
    return obj
}
func (self *Nodes_MatchCase) InitNodes_MatchCase(arg1 *Ast_AlgeValInfo, arg2 *LnsList, arg3 *Nodes_BlockNode) {
    self.valInfo = arg1
    self.valParamNameList = arg2
    self.block = arg3
}
func (self *Nodes_MatchCase) Get_valInfo() *Ast_AlgeValInfo{ return self.valInfo }
func (self *Nodes_MatchCase) Get_valParamNameList() *LnsList{ return self.valParamNameList }
func (self *Nodes_MatchCase) Get_block() *Nodes_BlockNode{ return self.block }

// declaration Class -- MatchNode
type Nodes_MatchNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_algeTypeInfo() *Ast_AlgeTypeInfo
    Get_caseKind() LnsInt
    Get_caseList() *LnsList
    Get_commentList() LnsAny
    Get_defaultBlock() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_failSafeDefault() bool
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_val() *Nodes_Node
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_MatchNode struct {
    Nodes_Node
    val *Nodes_Node
    algeTypeInfo *Ast_AlgeTypeInfo
    caseList *LnsList
    defaultBlock LnsAny
    caseKind LnsInt
    failSafeDefault bool
    FP Nodes_MatchNodeMtd
}
func Nodes_MatchNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_MatchNode).FP
}
type Nodes_MatchNodeDownCast interface {
    ToNodes_MatchNode() *Nodes_MatchNode
}
func Nodes_MatchNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_MatchNodeDownCast)
    if ok { return work.ToNodes_MatchNode() }
    return nil
}
func (obj *Nodes_MatchNode) ToNodes_MatchNode() *Nodes_MatchNode {
    return obj
}
func NewNodes_MatchNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node, arg6 *Ast_AlgeTypeInfo, arg7 *LnsList, arg8 LnsAny, arg9 LnsInt, arg10 bool) *Nodes_MatchNode {
    obj := &Nodes_MatchNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_MatchNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_MatchNode) Get_val() *Nodes_Node{ return self.val }
func (self *Nodes_MatchNode) Get_algeTypeInfo() *Ast_AlgeTypeInfo{ return self.algeTypeInfo }
func (self *Nodes_MatchNode) Get_caseList() *LnsList{ return self.caseList }
func (self *Nodes_MatchNode) Get_defaultBlock() LnsAny{ return self.defaultBlock }
func (self *Nodes_MatchNode) Get_caseKind() LnsInt{ return self.caseKind }
func (self *Nodes_MatchNode) Get_failSafeDefault() bool{ return self.failSafeDefault }
// 1: decl @lune.@base.@Nodes.MatchNode.processFilter
func (self *Nodes_MatchNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessMatch(self, opt)
}

// 1: decl @lune.@base.@Nodes.MatchNode.canBeRight
func (self *Nodes_MatchNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.MatchNode.canBeLeft
func (self *Nodes_MatchNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.MatchNode.canBeStatement
func (self *Nodes_MatchNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_MatchNode) InitNodes_MatchNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,algeTypeInfo *Ast_AlgeTypeInfo,caseList *LnsList,defaultBlock LnsAny,caseKind LnsInt,failSafeDefault bool) {
    self.InitNodes_Node(id, 68, pos, macroArgFlag, typeList)
    self.val = val
    
    self.algeTypeInfo = algeTypeInfo
    
    self.caseList = caseList
    
    self.defaultBlock = defaultBlock
    
    self.caseKind = caseKind
    
    self.failSafeDefault = failSafeDefault
    
}

// 681: decl @lune.@base.@Nodes.MatchNode.create
func Nodes_MatchNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,algeTypeInfo *Ast_AlgeTypeInfo,caseList *LnsList,defaultBlock LnsAny,caseKind LnsInt,failSafeDefault bool) *Nodes_MatchNode {
    var node *Nodes_MatchNode
    node = NewNodes_MatchNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, val, algeTypeInfo, caseList, defaultBlock, caseKind, failSafeDefault)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.MatchNode.visit
func (self *Nodes_MatchNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.val
        if _switch34624 := visitor(child, &self.Nodes_Node, "val", depth); _switch34624 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch34624 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        {
            _child := self.defaultBlock
            if _child != nil {
                child := _child.(*Nodes_Node)
                if _switch34679 := visitor(child, &self.Nodes_Node, "defaultBlock", depth); _switch34679 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch34679 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 1967: decl @lune.@base.@Nodes.MatchNode.getBreakKind
func (self *Nodes_MatchNode) GetBreakKind(checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = Nodes_BreakKind__None
    var fullCase bool
    fullCase = self.caseKind != Nodes_CaseKind__Lack
    for _, _caseInfo := range( self.caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        var work LnsInt
        work = caseInfo.FP.Get_block().FP.GetBreakKind(checkMode)
        var goNext bool
        goNext = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( (work == Nodes_BreakKind__None)) ||
            Lns_GetEnv().SetStackVal( Lns_op_not(fullCase)) ).(bool)
        if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
            if work == Nodes_BreakKind__Return{
                return Nodes_BreakKind__Return
            }
            if work == Nodes_BreakKind__NeverRet{
                return Nodes_BreakKind__NeverRet
            }
        } else { 
            if _switch34832 := work; _switch34832 == Nodes_BreakKind__None {
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                    Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                    if goNext{
                        return Nodes_BreakKind__None
                    }
                }
            } else {
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                    Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                    kind = work
                    
                }
            }
        }
        
    }
    {
        _block := self.defaultBlock
        if _block != nil {
            block := _block.(*Nodes_Node)
            var work LnsInt
            work = block.FP.GetBreakKind(checkMode)
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch34944 := work; _switch34944 == Nodes_BreakKind__None {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        Lns_GetEnv().SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( kind == Nodes_BreakKind__None) ||
                        Lns_GetEnv().SetStackVal( kind > work) ).(bool){
                        kind = work
                        
                    }
                }
            }
            
            return kind
        }
    }
    if fullCase{
        return kind
    }
    return Nodes_BreakKind__None
}


// declaration Class -- LuneKindNode
type Nodes_LuneKindNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_exp() *Nodes_Node
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LuneKindNode struct {
    Nodes_Node
    exp *Nodes_Node
    FP Nodes_LuneKindNodeMtd
}
func Nodes_LuneKindNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LuneKindNode).FP
}
type Nodes_LuneKindNodeDownCast interface {
    ToNodes_LuneKindNode() *Nodes_LuneKindNode
}
func Nodes_LuneKindNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LuneKindNodeDownCast)
    if ok { return work.ToNodes_LuneKindNode() }
    return nil
}
func (obj *Nodes_LuneKindNode) ToNodes_LuneKindNode() *Nodes_LuneKindNode {
    return obj
}
func NewNodes_LuneKindNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node) *Nodes_LuneKindNode {
    obj := &Nodes_LuneKindNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LuneKindNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_LuneKindNode) Get_exp() *Nodes_Node{ return self.exp }
// 1: decl @lune.@base.@Nodes.LuneKindNode.processFilter
func (self *Nodes_LuneKindNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLuneKind(self, opt)
}

// 1: decl @lune.@base.@Nodes.LuneKindNode.canBeRight
func (self *Nodes_LuneKindNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LuneKindNode.canBeLeft
func (self *Nodes_LuneKindNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LuneKindNode.canBeStatement
func (self *Nodes_LuneKindNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LuneKindNode) InitNodes_LuneKindNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(id, 69, pos, macroArgFlag, typeList)
    self.exp = exp
    
}

// 681: decl @lune.@base.@Nodes.LuneKindNode.create
func Nodes_LuneKindNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_LuneKindNode {
    var node *Nodes_LuneKindNode
    node = NewNodes_LuneKindNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LuneKindNode.visit
func (self *Nodes_LuneKindNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if _switch35302 := visitor(child, &self.Nodes_Node, "exp", depth); _switch35302 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch35302 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- DeclMacroNode
type Nodes_DeclMacroNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_declInfo() *Nodes_DeclMacroInfo
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_DeclMacroNode struct {
    Nodes_Node
    declInfo *Nodes_DeclMacroInfo
    FP Nodes_DeclMacroNodeMtd
}
func Nodes_DeclMacroNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclMacroNode).FP
}
type Nodes_DeclMacroNodeDownCast interface {
    ToNodes_DeclMacroNode() *Nodes_DeclMacroNode
}
func Nodes_DeclMacroNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclMacroNodeDownCast)
    if ok { return work.ToNodes_DeclMacroNode() }
    return nil
}
func (obj *Nodes_DeclMacroNode) ToNodes_DeclMacroNode() *Nodes_DeclMacroNode {
    return obj
}
func NewNodes_DeclMacroNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_DeclMacroInfo) *Nodes_DeclMacroNode {
    obj := &Nodes_DeclMacroNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclMacroNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_DeclMacroNode) Get_declInfo() *Nodes_DeclMacroInfo{ return self.declInfo }
// 1: decl @lune.@base.@Nodes.DeclMacroNode.processFilter
func (self *Nodes_DeclMacroNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclMacro(self, opt)
}

// 1: decl @lune.@base.@Nodes.DeclMacroNode.canBeRight
func (self *Nodes_DeclMacroNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclMacroNode.canBeLeft
func (self *Nodes_DeclMacroNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.DeclMacroNode.canBeStatement
func (self *Nodes_DeclMacroNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_DeclMacroNode) InitNodes_DeclMacroNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclMacroInfo) {
    self.InitNodes_Node(id, 70, pos, macroArgFlag, typeList)
    self.declInfo = declInfo
    
}

// 681: decl @lune.@base.@Nodes.DeclMacroNode.create
func Nodes_DeclMacroNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclMacroInfo) *Nodes_DeclMacroNode {
    var node *Nodes_DeclMacroNode
    node = NewNodes_DeclMacroNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.DeclMacroNode.visit
func (self *Nodes_DeclMacroNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- MacroEval
type Nodes_MacroEvalMtd interface {
    Eval(arg1 *Ast_ProcessInfo, arg2 *Nodes_DeclMacroNode) *Lns_luaValue
    EvalFromCode(arg1 *Ast_ProcessInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) *Lns_luaValue
}
type Nodes_MacroEval struct {
    FP Nodes_MacroEvalMtd
}
func Nodes_MacroEval2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_MacroEval).FP
}
type Nodes_MacroEvalDownCast interface {
    ToNodes_MacroEval() *Nodes_MacroEval
}
func Nodes_MacroEvalDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_MacroEvalDownCast)
    if ok { return work.ToNodes_MacroEval() }
    return nil
}
func (obj *Nodes_MacroEval) ToNodes_MacroEval() *Nodes_MacroEval {
    return obj
}
func (self *Nodes_MacroEval) InitNodes_MacroEval() {
}



// declaration Class -- TestCaseNode
type Nodes_TestCaseNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_block() *Nodes_BlockNode
    Get_commentList() LnsAny
    Get_ctrlName() string
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_impNode() *Nodes_Node
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_name() *Types_Token
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_TestCaseNode struct {
    Nodes_Node
    name *Types_Token
    impNode *Nodes_Node
    ctrlName string
    block *Nodes_BlockNode
    FP Nodes_TestCaseNodeMtd
}
func Nodes_TestCaseNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_TestCaseNode).FP
}
type Nodes_TestCaseNodeDownCast interface {
    ToNodes_TestCaseNode() *Nodes_TestCaseNode
}
func Nodes_TestCaseNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_TestCaseNodeDownCast)
    if ok { return work.ToNodes_TestCaseNode() }
    return nil
}
func (obj *Nodes_TestCaseNode) ToNodes_TestCaseNode() *Nodes_TestCaseNode {
    return obj
}
func NewNodes_TestCaseNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 *Nodes_Node, arg7 string, arg8 *Nodes_BlockNode) *Nodes_TestCaseNode {
    obj := &Nodes_TestCaseNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_TestCaseNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_TestCaseNode) Get_name() *Types_Token{ return self.name }
func (self *Nodes_TestCaseNode) Get_impNode() *Nodes_Node{ return self.impNode }
func (self *Nodes_TestCaseNode) Get_ctrlName() string{ return self.ctrlName }
func (self *Nodes_TestCaseNode) Get_block() *Nodes_BlockNode{ return self.block }
// 1: decl @lune.@base.@Nodes.TestCaseNode.processFilter
func (self *Nodes_TestCaseNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessTestCase(self, opt)
}

// 1: decl @lune.@base.@Nodes.TestCaseNode.canBeRight
func (self *Nodes_TestCaseNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.TestCaseNode.canBeLeft
func (self *Nodes_TestCaseNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.TestCaseNode.canBeStatement
func (self *Nodes_TestCaseNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_TestCaseNode) InitNodes_TestCaseNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token,impNode *Nodes_Node,ctrlName string,block *Nodes_BlockNode) {
    self.InitNodes_Node(id, 71, pos, macroArgFlag, typeList)
    self.name = name
    
    self.impNode = impNode
    
    self.ctrlName = ctrlName
    
    self.block = block
    
}

// 681: decl @lune.@base.@Nodes.TestCaseNode.create
func Nodes_TestCaseNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,name *Types_Token,impNode *Nodes_Node,ctrlName string,block *Nodes_BlockNode) *Nodes_TestCaseNode {
    var node *Nodes_TestCaseNode
    node = NewNodes_TestCaseNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, name, impNode, ctrlName, block)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.TestCaseNode.visit
func (self *Nodes_TestCaseNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.impNode
        if _switch36055 := visitor(child, &self.Nodes_Node, "impNode", depth); _switch36055 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch36055 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if _switch36112 := visitor(&child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch36112 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch36112 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- TestBlockNode
type Nodes_TestBlockNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_stmtList() *LnsList
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsInnerPos(arg1 *Types_Position) bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_TestBlockNode struct {
    Nodes_Node
    stmtList *LnsList
    FP Nodes_TestBlockNodeMtd
}
func Nodes_TestBlockNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_TestBlockNode).FP
}
type Nodes_TestBlockNodeDownCast interface {
    ToNodes_TestBlockNode() *Nodes_TestBlockNode
}
func Nodes_TestBlockNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_TestBlockNodeDownCast)
    if ok { return work.ToNodes_TestBlockNode() }
    return nil
}
func (obj *Nodes_TestBlockNode) ToNodes_TestBlockNode() *Nodes_TestBlockNode {
    return obj
}
func NewNodes_TestBlockNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsList) *Nodes_TestBlockNode {
    obj := &Nodes_TestBlockNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_TestBlockNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_TestBlockNode) Get_stmtList() *LnsList{ return self.stmtList }
// 1: decl @lune.@base.@Nodes.TestBlockNode.processFilter
func (self *Nodes_TestBlockNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessTestBlock(self, opt)
}

// 1: decl @lune.@base.@Nodes.TestBlockNode.canBeRight
func (self *Nodes_TestBlockNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return false
}

// 1: decl @lune.@base.@Nodes.TestBlockNode.canBeLeft
func (self *Nodes_TestBlockNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.TestBlockNode.canBeStatement
func (self *Nodes_TestBlockNode) CanBeStatement() bool {
    return true
}

// 676: DeclConstr
func (self *Nodes_TestBlockNode) InitNodes_TestBlockNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) {
    self.InitNodes_Node(id, 72, pos, macroArgFlag, typeList)
    self.stmtList = stmtList
    
}

// 681: decl @lune.@base.@Nodes.TestBlockNode.create
func Nodes_TestBlockNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) *Nodes_TestBlockNode {
    var node *Nodes_TestBlockNode
    node = NewNodes_TestBlockNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, stmtList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.TestBlockNode.visit
func (self *Nodes_TestBlockNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var list *LnsList
        list = self.stmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch36465 := visitor(child, &self.Nodes_Node, "stmtList", depth); _switch36465 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                    return false
                }
            } else if _switch36465 == Nodes_NodeVisitMode__End {
                return false
            }
        }
    }
    return true
}

// 2020: decl @lune.@base.@Nodes.TestBlockNode.isInnerPos
func (self *Nodes_TestBlockNode) IsInnerPos(pos *Types_Position) bool {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.FP.Get_pos().StreamName == pos.StreamName) &&
        Lns_GetEnv().SetStackVal( self.FP.Get_pos().LineNo < pos.LineNo) &&
        Lns_GetEnv().SetStackVal( self.FP.Get_stmtList().Len() > 0) &&
        Lns_GetEnv().SetStackVal( self.FP.Get_stmtList().GetAt(1).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_pos().LineNo >= pos.LineNo) ).(bool)){
        return true
    }
    return false
}


// declaration Class -- AbbrNode
type Nodes_AbbrNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_AbbrNode struct {
    Nodes_Node
    FP Nodes_AbbrNodeMtd
}
func Nodes_AbbrNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_AbbrNode).FP
}
type Nodes_AbbrNodeDownCast interface {
    ToNodes_AbbrNode() *Nodes_AbbrNode
}
func Nodes_AbbrNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_AbbrNodeDownCast)
    if ok { return work.ToNodes_AbbrNode() }
    return nil
}
func (obj *Nodes_AbbrNode) ToNodes_AbbrNode() *Nodes_AbbrNode {
    return obj
}
func NewNodes_AbbrNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList) *Nodes_AbbrNode {
    obj := &Nodes_AbbrNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_AbbrNode(arg1, arg2, arg3, arg4)
    return obj
}
// 1: decl @lune.@base.@Nodes.AbbrNode.processFilter
func (self *Nodes_AbbrNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessAbbr(self, opt)
}

// 1: decl @lune.@base.@Nodes.AbbrNode.canBeRight
func (self *Nodes_AbbrNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.AbbrNode.canBeLeft
func (self *Nodes_AbbrNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.AbbrNode.canBeStatement
func (self *Nodes_AbbrNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_AbbrNode) InitNodes_AbbrNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(id, 73, pos, macroArgFlag, typeList)
}

// 681: decl @lune.@base.@Nodes.AbbrNode.create
func Nodes_AbbrNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList) *Nodes_AbbrNode {
    var node *Nodes_AbbrNode
    node = NewNodes_AbbrNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.AbbrNode.visit
func (self *Nodes_AbbrNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}


// declaration Class -- BoxingNode
type Nodes_BoxingNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_src() *Nodes_Node
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_BoxingNode struct {
    Nodes_Node
    src *Nodes_Node
    FP Nodes_BoxingNodeMtd
}
func Nodes_BoxingNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_BoxingNode).FP
}
type Nodes_BoxingNodeDownCast interface {
    ToNodes_BoxingNode() *Nodes_BoxingNode
}
func Nodes_BoxingNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_BoxingNodeDownCast)
    if ok { return work.ToNodes_BoxingNode() }
    return nil
}
func (obj *Nodes_BoxingNode) ToNodes_BoxingNode() *Nodes_BoxingNode {
    return obj
}
func NewNodes_BoxingNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node) *Nodes_BoxingNode {
    obj := &Nodes_BoxingNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_BoxingNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_BoxingNode) Get_src() *Nodes_Node{ return self.src }
// 1: decl @lune.@base.@Nodes.BoxingNode.processFilter
func (self *Nodes_BoxingNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessBoxing(self, opt)
}

// 1: decl @lune.@base.@Nodes.BoxingNode.canBeRight
func (self *Nodes_BoxingNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.BoxingNode.canBeLeft
func (self *Nodes_BoxingNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.BoxingNode.canBeStatement
func (self *Nodes_BoxingNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_BoxingNode) InitNodes_BoxingNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,src *Nodes_Node) {
    self.InitNodes_Node(id, 74, pos, macroArgFlag, typeList)
    self.src = src
    
}

// 681: decl @lune.@base.@Nodes.BoxingNode.create
func Nodes_BoxingNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,src *Nodes_Node) *Nodes_BoxingNode {
    var node *Nodes_BoxingNode
    node = NewNodes_BoxingNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, src)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.BoxingNode.visit
func (self *Nodes_BoxingNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.src
        if _switch37136 := visitor(child, &self.Nodes_Node, "src", depth); _switch37136 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch37136 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- UnboxingNode
type Nodes_UnboxingNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_src() *Nodes_Node
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_UnboxingNode struct {
    Nodes_Node
    src *Nodes_Node
    FP Nodes_UnboxingNodeMtd
}
func Nodes_UnboxingNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_UnboxingNode).FP
}
type Nodes_UnboxingNodeDownCast interface {
    ToNodes_UnboxingNode() *Nodes_UnboxingNode
}
func Nodes_UnboxingNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_UnboxingNodeDownCast)
    if ok { return work.ToNodes_UnboxingNode() }
    return nil
}
func (obj *Nodes_UnboxingNode) ToNodes_UnboxingNode() *Nodes_UnboxingNode {
    return obj
}
func NewNodes_UnboxingNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Nodes_Node) *Nodes_UnboxingNode {
    obj := &Nodes_UnboxingNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_UnboxingNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_UnboxingNode) Get_src() *Nodes_Node{ return self.src }
// 1: decl @lune.@base.@Nodes.UnboxingNode.processFilter
func (self *Nodes_UnboxingNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessUnboxing(self, opt)
}

// 1: decl @lune.@base.@Nodes.UnboxingNode.canBeRight
func (self *Nodes_UnboxingNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.UnboxingNode.canBeLeft
func (self *Nodes_UnboxingNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.UnboxingNode.canBeStatement
func (self *Nodes_UnboxingNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_UnboxingNode) InitNodes_UnboxingNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,src *Nodes_Node) {
    self.InitNodes_Node(id, 75, pos, macroArgFlag, typeList)
    self.src = src
    
}

// 681: decl @lune.@base.@Nodes.UnboxingNode.create
func Nodes_UnboxingNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,src *Nodes_Node) *Nodes_UnboxingNode {
    var node *Nodes_UnboxingNode
    node = NewNodes_UnboxingNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, src)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.UnboxingNode.visit
func (self *Nodes_UnboxingNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var child *Nodes_Node
        child = self.src
        if _switch37481 := visitor(child, &self.Nodes_Node, "src", depth); _switch37481 == Nodes_NodeVisitMode__Child {
            if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                return false
            }
        } else if _switch37481 == Nodes_NodeVisitMode__End {
            return false
        }
    }
    return true
}


// declaration Class -- LiteralNilNode
type Nodes_LiteralNilNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralNilNode struct {
    Nodes_Node
    FP Nodes_LiteralNilNodeMtd
}
func Nodes_LiteralNilNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralNilNode).FP
}
type Nodes_LiteralNilNodeDownCast interface {
    ToNodes_LiteralNilNode() *Nodes_LiteralNilNode
}
func Nodes_LiteralNilNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralNilNodeDownCast)
    if ok { return work.ToNodes_LiteralNilNode() }
    return nil
}
func (obj *Nodes_LiteralNilNode) ToNodes_LiteralNilNode() *Nodes_LiteralNilNode {
    return obj
}
func NewNodes_LiteralNilNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList) *Nodes_LiteralNilNode {
    obj := &Nodes_LiteralNilNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralNilNode(arg1, arg2, arg3, arg4)
    return obj
}
// 1: decl @lune.@base.@Nodes.LiteralNilNode.processFilter
func (self *Nodes_LiteralNilNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralNil(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralNilNode.canBeRight
func (self *Nodes_LiteralNilNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralNilNode.canBeLeft
func (self *Nodes_LiteralNilNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralNilNode.canBeStatement
func (self *Nodes_LiteralNilNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralNilNode) InitNodes_LiteralNilNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(id, 76, pos, macroArgFlag, typeList)
}

// 681: decl @lune.@base.@Nodes.LiteralNilNode.create
func Nodes_LiteralNilNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList) *Nodes_LiteralNilNode {
    var node *Nodes_LiteralNilNode
    node = NewNodes_LiteralNilNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralNilNode.visit
func (self *Nodes_LiteralNilNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 2214: decl @lune.@base.@Nodes.LiteralNilNode.getLiteral
func (self *Nodes_LiteralNilNode) GetLiteral()(LnsAny, LnsAny) {
    return Nodes_Literal__Nil_Obj, nil
}

// 2217: decl @lune.@base.@Nodes.LiteralNilNode.setupLiteralTokenList
func (self *Nodes_LiteralNilNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Symb, "nil")
    return true
}


// declaration Class -- LiteralCharNode
type Nodes_LiteralCharNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_num() LnsInt
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_token() *Types_Token
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralCharNode struct {
    Nodes_Node
    token *Types_Token
    num LnsInt
    FP Nodes_LiteralCharNodeMtd
}
func Nodes_LiteralCharNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralCharNode).FP
}
type Nodes_LiteralCharNodeDownCast interface {
    ToNodes_LiteralCharNode() *Nodes_LiteralCharNode
}
func Nodes_LiteralCharNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralCharNodeDownCast)
    if ok { return work.ToNodes_LiteralCharNode() }
    return nil
}
func (obj *Nodes_LiteralCharNode) ToNodes_LiteralCharNode() *Nodes_LiteralCharNode {
    return obj
}
func NewNodes_LiteralCharNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 LnsInt) *Nodes_LiteralCharNode {
    obj := &Nodes_LiteralCharNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralCharNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Nodes_LiteralCharNode) Get_token() *Types_Token{ return self.token }
func (self *Nodes_LiteralCharNode) Get_num() LnsInt{ return self.num }
// 1: decl @lune.@base.@Nodes.LiteralCharNode.processFilter
func (self *Nodes_LiteralCharNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralChar(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralCharNode.canBeRight
func (self *Nodes_LiteralCharNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralCharNode.canBeLeft
func (self *Nodes_LiteralCharNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralCharNode.canBeStatement
func (self *Nodes_LiteralCharNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralCharNode) InitNodes_LiteralCharNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsInt) {
    self.InitNodes_Node(id, 77, pos, macroArgFlag, typeList)
    self.token = token
    
    self.num = num
    
}

// 681: decl @lune.@base.@Nodes.LiteralCharNode.create
func Nodes_LiteralCharNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsInt) *Nodes_LiteralCharNode {
    var node *Nodes_LiteralCharNode
    node = NewNodes_LiteralCharNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, token, num)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralCharNode.visit
func (self *Nodes_LiteralCharNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 2223: decl @lune.@base.@Nodes.LiteralCharNode.getLiteral
func (self *Nodes_LiteralCharNode) GetLiteral()(LnsAny, LnsAny) {
    return &Nodes_Literal__Int{self.num}, nil
}

// 2226: decl @lune.@base.@Nodes.LiteralCharNode.setupLiteralTokenList
func (self *Nodes_LiteralCharNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Char, Lns_getVM().String_format("%d", []LnsAny{self.num}))
    return true
}


// declaration Class -- LiteralIntNode
type Nodes_LiteralIntNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_num() LnsInt
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_token() *Types_Token
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralIntNode struct {
    Nodes_Node
    token *Types_Token
    num LnsInt
    FP Nodes_LiteralIntNodeMtd
}
func Nodes_LiteralIntNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralIntNode).FP
}
type Nodes_LiteralIntNodeDownCast interface {
    ToNodes_LiteralIntNode() *Nodes_LiteralIntNode
}
func Nodes_LiteralIntNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralIntNodeDownCast)
    if ok { return work.ToNodes_LiteralIntNode() }
    return nil
}
func (obj *Nodes_LiteralIntNode) ToNodes_LiteralIntNode() *Nodes_LiteralIntNode {
    return obj
}
func NewNodes_LiteralIntNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 LnsInt) *Nodes_LiteralIntNode {
    obj := &Nodes_LiteralIntNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralIntNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Nodes_LiteralIntNode) Get_token() *Types_Token{ return self.token }
func (self *Nodes_LiteralIntNode) Get_num() LnsInt{ return self.num }
// 1: decl @lune.@base.@Nodes.LiteralIntNode.processFilter
func (self *Nodes_LiteralIntNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralInt(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralIntNode.canBeRight
func (self *Nodes_LiteralIntNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralIntNode.canBeLeft
func (self *Nodes_LiteralIntNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralIntNode.canBeStatement
func (self *Nodes_LiteralIntNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralIntNode) InitNodes_LiteralIntNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsInt) {
    self.InitNodes_Node(id, 78, pos, macroArgFlag, typeList)
    self.token = token
    
    self.num = num
    
}

// 681: decl @lune.@base.@Nodes.LiteralIntNode.create
func Nodes_LiteralIntNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsInt) *Nodes_LiteralIntNode {
    var node *Nodes_LiteralIntNode
    node = NewNodes_LiteralIntNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, token, num)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralIntNode.visit
func (self *Nodes_LiteralIntNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 2231: decl @lune.@base.@Nodes.LiteralIntNode.getLiteral
func (self *Nodes_LiteralIntNode) GetLiteral()(LnsAny, LnsAny) {
    return &Nodes_Literal__Int{self.num}, nil
}

// 2234: decl @lune.@base.@Nodes.LiteralIntNode.setupLiteralTokenList
func (self *Nodes_LiteralIntNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Int, Lns_getVM().String_format("%d", []LnsAny{self.num}))
    return true
}


// declaration Class -- LiteralRealNode
type Nodes_LiteralRealNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_num() LnsReal
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_token() *Types_Token
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralRealNode struct {
    Nodes_Node
    token *Types_Token
    num LnsReal
    FP Nodes_LiteralRealNodeMtd
}
func Nodes_LiteralRealNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralRealNode).FP
}
type Nodes_LiteralRealNodeDownCast interface {
    ToNodes_LiteralRealNode() *Nodes_LiteralRealNode
}
func Nodes_LiteralRealNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralRealNodeDownCast)
    if ok { return work.ToNodes_LiteralRealNode() }
    return nil
}
func (obj *Nodes_LiteralRealNode) ToNodes_LiteralRealNode() *Nodes_LiteralRealNode {
    return obj
}
func NewNodes_LiteralRealNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 LnsReal) *Nodes_LiteralRealNode {
    obj := &Nodes_LiteralRealNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralRealNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Nodes_LiteralRealNode) Get_token() *Types_Token{ return self.token }
func (self *Nodes_LiteralRealNode) Get_num() LnsReal{ return self.num }
// 1: decl @lune.@base.@Nodes.LiteralRealNode.processFilter
func (self *Nodes_LiteralRealNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralReal(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralRealNode.canBeRight
func (self *Nodes_LiteralRealNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralRealNode.canBeLeft
func (self *Nodes_LiteralRealNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralRealNode.canBeStatement
func (self *Nodes_LiteralRealNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralRealNode) InitNodes_LiteralRealNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsReal) {
    self.InitNodes_Node(id, 79, pos, macroArgFlag, typeList)
    self.token = token
    
    self.num = num
    
}

// 681: decl @lune.@base.@Nodes.LiteralRealNode.create
func Nodes_LiteralRealNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsReal) *Nodes_LiteralRealNode {
    var node *Nodes_LiteralRealNode
    node = NewNodes_LiteralRealNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, token, num)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralRealNode.visit
func (self *Nodes_LiteralRealNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 2239: decl @lune.@base.@Nodes.LiteralRealNode.getLiteral
func (self *Nodes_LiteralRealNode) GetLiteral()(LnsAny, LnsAny) {
    return &Nodes_Literal__Real{self.num}, nil
}

// 2242: decl @lune.@base.@Nodes.LiteralRealNode.setupLiteralTokenList
func (self *Nodes_LiteralRealNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Real, Lns_getVM().String_format("%g", []LnsAny{self.num}))
    return true
}


// declaration Class -- LiteralArrayNode
type Nodes_LiteralArrayNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() LnsAny
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralArrayNode struct {
    Nodes_Node
    expList LnsAny
    FP Nodes_LiteralArrayNodeMtd
}
func Nodes_LiteralArrayNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralArrayNode).FP
}
type Nodes_LiteralArrayNodeDownCast interface {
    ToNodes_LiteralArrayNode() *Nodes_LiteralArrayNode
}
func Nodes_LiteralArrayNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralArrayNodeDownCast)
    if ok { return work.ToNodes_LiteralArrayNode() }
    return nil
}
func (obj *Nodes_LiteralArrayNode) ToNodes_LiteralArrayNode() *Nodes_LiteralArrayNode {
    return obj
}
func NewNodes_LiteralArrayNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsAny) *Nodes_LiteralArrayNode {
    obj := &Nodes_LiteralArrayNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralArrayNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_LiteralArrayNode) Get_expList() LnsAny{ return self.expList }
// 1: decl @lune.@base.@Nodes.LiteralArrayNode.processFilter
func (self *Nodes_LiteralArrayNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralArray(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralArrayNode.canBeRight
func (self *Nodes_LiteralArrayNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralArrayNode.canBeLeft
func (self *Nodes_LiteralArrayNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralArrayNode.canBeStatement
func (self *Nodes_LiteralArrayNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralArrayNode) InitNodes_LiteralArrayNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList LnsAny) {
    self.InitNodes_Node(id, 80, pos, macroArgFlag, typeList)
    self.expList = expList
    
}

// 681: decl @lune.@base.@Nodes.LiteralArrayNode.create
func Nodes_LiteralArrayNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList LnsAny) *Nodes_LiteralArrayNode {
    var node *Nodes_LiteralArrayNode
    node = NewNodes_LiteralArrayNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralArrayNode.visit
func (self *Nodes_LiteralArrayNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch39041 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch39041 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch39041 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 2247: decl @lune.@base.@Nodes.LiteralArrayNode.getLiteral
func (self *Nodes_LiteralArrayNode) GetLiteral()(LnsAny, LnsAny) {
    var literalList *LnsList
    literalList = NewLnsList([]LnsAny{})
    {
        __exp := self.expList
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            for _, _node := range( _exp.FP.Get_expList().Items ) {
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                var literal LnsAny
                var mess LnsAny
                literal,mess = node.FP.GetLiteral()
                if literal != nil{
                    literal_10432 := literal
                    literalList.Insert(literal_10432)
                } else {
                    return nil, mess
                }
            }
        }
    }
    return &Nodes_Literal__ARRAY{literalList}, nil
}

// 2263: decl @lune.@base.@Nodes.LiteralArrayNode.setupLiteralTokenList
func (self *Nodes_LiteralArrayNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, "[@")
    {
        __exp := self.expList
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            for _index, _node := range( _exp.FP.Get_expList().Items ) {
                index := _index + 1
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, ",")
                }
                if Lns_op_not(node.FP.SetupLiteralTokenList(list)){
                    return false
                }
            }
        }
    }
    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, "]")
    return true
}


// declaration Class -- LiteralListNode
type Nodes_LiteralListNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() LnsAny
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralListNode struct {
    Nodes_Node
    expList LnsAny
    FP Nodes_LiteralListNodeMtd
}
func Nodes_LiteralListNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralListNode).FP
}
type Nodes_LiteralListNodeDownCast interface {
    ToNodes_LiteralListNode() *Nodes_LiteralListNode
}
func Nodes_LiteralListNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralListNodeDownCast)
    if ok { return work.ToNodes_LiteralListNode() }
    return nil
}
func (obj *Nodes_LiteralListNode) ToNodes_LiteralListNode() *Nodes_LiteralListNode {
    return obj
}
func NewNodes_LiteralListNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsAny) *Nodes_LiteralListNode {
    obj := &Nodes_LiteralListNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralListNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_LiteralListNode) Get_expList() LnsAny{ return self.expList }
// 1: decl @lune.@base.@Nodes.LiteralListNode.processFilter
func (self *Nodes_LiteralListNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralList(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralListNode.canBeRight
func (self *Nodes_LiteralListNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralListNode.canBeLeft
func (self *Nodes_LiteralListNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralListNode.canBeStatement
func (self *Nodes_LiteralListNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralListNode) InitNodes_LiteralListNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList LnsAny) {
    self.InitNodes_Node(id, 81, pos, macroArgFlag, typeList)
    self.expList = expList
    
}

// 681: decl @lune.@base.@Nodes.LiteralListNode.create
func Nodes_LiteralListNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList LnsAny) *Nodes_LiteralListNode {
    var node *Nodes_LiteralListNode
    node = NewNodes_LiteralListNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralListNode.visit
func (self *Nodes_LiteralListNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch39388 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch39388 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch39388 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 2280: decl @lune.@base.@Nodes.LiteralListNode.getLiteral
func (self *Nodes_LiteralListNode) GetLiteral()(LnsAny, LnsAny) {
    var literalList *LnsList
    literalList = NewLnsList([]LnsAny{})
    {
        __exp := self.expList
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            for _, _node := range( _exp.FP.Get_expList().Items ) {
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                var literal LnsAny
                var mess LnsAny
                literal,mess = node.FP.GetLiteral()
                if literal != nil{
                    literal_10460 := literal
                    literalList.Insert(literal_10460)
                } else {
                    return nil, mess
                }
            }
        }
    }
    return &Nodes_Literal__LIST{literalList}, nil
}

// 2296: decl @lune.@base.@Nodes.LiteralListNode.setupLiteralTokenList
func (self *Nodes_LiteralListNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, "[")
    {
        __exp := self.expList
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            for _index, _node := range( _exp.FP.Get_expList().Items ) {
                index := _index + 1
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, ",")
                }
                if Lns_op_not(node.FP.SetupLiteralTokenList(list)){
                    return false
                }
            }
        }
    }
    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, "]")
    return true
}


// declaration Class -- LiteralSetNode
type Nodes_LiteralSetNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expList() LnsAny
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralSetNode struct {
    Nodes_Node
    expList LnsAny
    FP Nodes_LiteralSetNodeMtd
}
func Nodes_LiteralSetNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralSetNode).FP
}
type Nodes_LiteralSetNodeDownCast interface {
    ToNodes_LiteralSetNode() *Nodes_LiteralSetNode
}
func Nodes_LiteralSetNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralSetNodeDownCast)
    if ok { return work.ToNodes_LiteralSetNode() }
    return nil
}
func (obj *Nodes_LiteralSetNode) ToNodes_LiteralSetNode() *Nodes_LiteralSetNode {
    return obj
}
func NewNodes_LiteralSetNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 LnsAny) *Nodes_LiteralSetNode {
    obj := &Nodes_LiteralSetNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralSetNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_LiteralSetNode) Get_expList() LnsAny{ return self.expList }
// 1: decl @lune.@base.@Nodes.LiteralSetNode.processFilter
func (self *Nodes_LiteralSetNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralSet(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralSetNode.canBeRight
func (self *Nodes_LiteralSetNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralSetNode.canBeLeft
func (self *Nodes_LiteralSetNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralSetNode.canBeStatement
func (self *Nodes_LiteralSetNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralSetNode) InitNodes_LiteralSetNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList LnsAny) {
    self.InitNodes_Node(id, 82, pos, macroArgFlag, typeList)
    self.expList = expList
    
}

// 681: decl @lune.@base.@Nodes.LiteralSetNode.create
func Nodes_LiteralSetNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,expList LnsAny) *Nodes_LiteralSetNode {
    var node *Nodes_LiteralSetNode
    node = NewNodes_LiteralSetNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralSetNode.visit
func (self *Nodes_LiteralSetNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.expList
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch39735 := visitor(&child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch39735 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch39735 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 2314: decl @lune.@base.@Nodes.LiteralSetNode.getLiteral
func (self *Nodes_LiteralSetNode) GetLiteral()(LnsAny, LnsAny) {
    var literalList *LnsList
    literalList = NewLnsList([]LnsAny{})
    {
        __exp := self.expList
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            for _, _node := range( _exp.FP.Get_expList().Items ) {
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                var literal LnsAny
                var mess LnsAny
                literal,mess = node.FP.GetLiteral()
                if literal != nil{
                    literal_10488 := literal
                    literalList.Insert(literal_10488)
                } else {
                    return nil, mess
                }
            }
        }
    }
    return &Nodes_Literal__SET{literalList}, nil
}

// 2330: decl @lune.@base.@Nodes.LiteralSetNode.setupLiteralTokenList
func (self *Nodes_LiteralSetNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, "(@")
    {
        __exp := self.expList
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            for _index, _node := range( _exp.FP.Get_expList().Items ) {
                index := _index + 1
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, ",")
                }
                if Lns_op_not(node.FP.SetupLiteralTokenList(list)){
                    return false
                }
            }
        }
    }
    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, ")")
    return true
}


// declaration Class -- PairItem
type Nodes_PairItemMtd interface {
    Get_key() *Nodes_Node
    Get_val() *Nodes_Node
}
type Nodes_PairItem struct {
    key *Nodes_Node
    val *Nodes_Node
    FP Nodes_PairItemMtd
}
func Nodes_PairItem2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_PairItem).FP
}
type Nodes_PairItemDownCast interface {
    ToNodes_PairItem() *Nodes_PairItem
}
func Nodes_PairItemDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_PairItemDownCast)
    if ok { return work.ToNodes_PairItem() }
    return nil
}
func (obj *Nodes_PairItem) ToNodes_PairItem() *Nodes_PairItem {
    return obj
}
func NewNodes_PairItem(arg1 *Nodes_Node, arg2 *Nodes_Node) *Nodes_PairItem {
    obj := &Nodes_PairItem{}
    obj.FP = obj
    obj.InitNodes_PairItem(arg1, arg2)
    return obj
}
func (self *Nodes_PairItem) InitNodes_PairItem(arg1 *Nodes_Node, arg2 *Nodes_Node) {
    self.key = arg1
    self.val = arg2
}
func (self *Nodes_PairItem) Get_key() *Nodes_Node{ return self.key }
func (self *Nodes_PairItem) Get_val() *Nodes_Node{ return self.val }

// declaration Class -- LiteralMapNode
type Nodes_LiteralMapNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_map() *LnsMap
    Get_pairList() *LnsList
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralMapNode struct {
    Nodes_Node
    _map *LnsMap
    pairList *LnsList
    FP Nodes_LiteralMapNodeMtd
}
func Nodes_LiteralMapNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralMapNode).FP
}
type Nodes_LiteralMapNodeDownCast interface {
    ToNodes_LiteralMapNode() *Nodes_LiteralMapNode
}
func Nodes_LiteralMapNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralMapNodeDownCast)
    if ok { return work.ToNodes_LiteralMapNode() }
    return nil
}
func (obj *Nodes_LiteralMapNode) ToNodes_LiteralMapNode() *Nodes_LiteralMapNode {
    return obj
}
func NewNodes_LiteralMapNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *LnsMap, arg6 *LnsList) *Nodes_LiteralMapNode {
    obj := &Nodes_LiteralMapNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralMapNode(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *Nodes_LiteralMapNode) Get_map() *LnsMap{ return self._map }
func (self *Nodes_LiteralMapNode) Get_pairList() *LnsList{ return self.pairList }
// 1: decl @lune.@base.@Nodes.LiteralMapNode.processFilter
func (self *Nodes_LiteralMapNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralMap(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralMapNode.canBeRight
func (self *Nodes_LiteralMapNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralMapNode.canBeLeft
func (self *Nodes_LiteralMapNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralMapNode.canBeStatement
func (self *Nodes_LiteralMapNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralMapNode) InitNodes_LiteralMapNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,_map *LnsMap,pairList *LnsList) {
    self.InitNodes_Node(id, 83, pos, macroArgFlag, typeList)
    self._map = _map
    
    self.pairList = pairList
    
}

// 681: decl @lune.@base.@Nodes.LiteralMapNode.create
func Nodes_LiteralMapNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,_map *LnsMap,pairList *LnsList) *Nodes_LiteralMapNode {
    var node *Nodes_LiteralMapNode
    node = NewNodes_LiteralMapNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, _map, pairList)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralMapNode.visit
func (self *Nodes_LiteralMapNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        var _map *LnsMap
        _map = self._map
        for _key, _val := range( _map.Items ) {
            key := _key.(Nodes_NodeDownCast).ToNodes_Node()
            val := _val.(Nodes_NodeDownCast).ToNodes_Node()
            {
                var child *Nodes_Node
                child = key
                if _switch40141 := visitor(child, &self.Nodes_Node, "map", depth); _switch40141 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch40141 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
            {
                var child *Nodes_Node
                child = val
                if _switch40196 := visitor(child, &self.Nodes_Node, "map", depth); _switch40196 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch40196 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 2347: decl @lune.@base.@Nodes.LiteralMapNode.getLiteral
func (self *Nodes_LiteralMapNode) GetLiteral()(LnsAny, LnsAny) {
    var litMap *LnsMap
    litMap = NewLnsMap( map[LnsAny]LnsAny{})
    for _key, _val := range( self._map.Items ) {
        key := _key.(Nodes_NodeDownCast).ToNodes_Node()
        val := _val.(Nodes_NodeDownCast).ToNodes_Node()
        var keyLiteral LnsAny
        var keyMess LnsAny
        keyLiteral,keyMess = key.FP.GetLiteral()
        var valLiteral LnsAny
        var valMess LnsAny
        valLiteral,valMess = val.FP.GetLiteral()
        if keyLiteral != nil && valLiteral != nil{
            keyLiteral_10517 := keyLiteral
            valLiteral_10518 := valLiteral
            litMap.Set(keyLiteral_10517,valLiteral_10518)
        } else {
            if Lns_op_not(keyLiteral){
                return nil, keyMess
            }
            if Lns_op_not(valLiteral){
                return nil, valMess
            }
        }
    }
    return &Nodes_Literal__MAP{litMap}, nil
}

// 2367: decl @lune.@base.@Nodes.LiteralMapNode.setupLiteralTokenList
func (self *Nodes_LiteralMapNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, "{")
    var lit2valNode *LnsMap
    lit2valNode = NewLnsMap( map[LnsAny]LnsAny{})
    for _key, _ := range( self._map.Items ) {
        key := _key.(Nodes_NodeDownCast).ToNodes_Node()
        var literal LnsAny
        literal = Nodes_convExp42626(Lns_2DDD(key.FP.GetLiteral()))
        if literal != nil{
            literal_10534 := literal
            switch _exp42670 := literal_10534.(type) {
            case *Nodes_Literal__Int:
            param := _exp42670.Val1
                lit2valNode.Set(param,key)
            case *Nodes_Literal__Str:
            param := _exp42670.Val1
                lit2valNode.Set(param,key)
            case *Nodes_Literal__Real:
            param := _exp42670.Val1
                lit2valNode.Set(param,key)
            default:
                return false
            }
        }
    }
    {
        __collection42737 := lit2valNode
        __sorted42737 := __collection42737.CreateKeyListStem()
        __sorted42737.Sort( LnsItemKindStem, nil )
        for _, __ := range( __sorted42737.Items ) {
            key := __collection42737.Items[ __ ].(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(key.FP.SetupLiteralTokenList(list)){
                return false
            }
            self.FP.AddTokenList(list, Types_TokenKind__Dlmt, ":")
            if Lns_op_not(Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(self._map.Items[key]) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_Node).FP.SetupLiteralTokenList(list)})/* 2397:14 */)){
                return false
            }
            self.FP.AddTokenList(list, Types_TokenKind__Dlmt, ",")
        }
    }
    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, "}")
    return true
}


// declaration Class -- LiteralStringNode
type Nodes_LiteralStringNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_dddParam() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_orgParam() LnsAny
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_threading() bool
    Get_token() *Types_Token
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralStringNode struct {
    Nodes_Node
    token *Types_Token
    orgParam LnsAny
    dddParam LnsAny
    threading bool
    FP Nodes_LiteralStringNodeMtd
}
func Nodes_LiteralStringNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralStringNode).FP
}
type Nodes_LiteralStringNodeDownCast interface {
    ToNodes_LiteralStringNode() *Nodes_LiteralStringNode
}
func Nodes_LiteralStringNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralStringNodeDownCast)
    if ok { return work.ToNodes_LiteralStringNode() }
    return nil
}
func (obj *Nodes_LiteralStringNode) ToNodes_LiteralStringNode() *Nodes_LiteralStringNode {
    return obj
}
func NewNodes_LiteralStringNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token, arg6 LnsAny, arg7 LnsAny, arg8 bool) *Nodes_LiteralStringNode {
    obj := &Nodes_LiteralStringNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralStringNode(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_LiteralStringNode) Get_token() *Types_Token{ return self.token }
func (self *Nodes_LiteralStringNode) Get_orgParam() LnsAny{ return self.orgParam }
func (self *Nodes_LiteralStringNode) Get_dddParam() LnsAny{ return self.dddParam }
func (self *Nodes_LiteralStringNode) Get_threading() bool{ return self.threading }
// 1: decl @lune.@base.@Nodes.LiteralStringNode.isThreading
func (self *Nodes_LiteralStringNode) IsThreading() bool {
    return self.threading
}

// 1: decl @lune.@base.@Nodes.LiteralStringNode.processFilter
func (self *Nodes_LiteralStringNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralString(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralStringNode.canBeRight
func (self *Nodes_LiteralStringNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralStringNode.canBeLeft
func (self *Nodes_LiteralStringNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralStringNode.canBeStatement
func (self *Nodes_LiteralStringNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralStringNode) InitNodes_LiteralStringNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token,orgParam LnsAny,dddParam LnsAny,threading bool) {
    self.InitNodes_Node(id, 84, pos, macroArgFlag, typeList)
    self.token = token
    
    self.orgParam = orgParam
    
    self.dddParam = dddParam
    
    self.threading = threading
    
}

// 681: decl @lune.@base.@Nodes.LiteralStringNode.create
func Nodes_LiteralStringNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token,orgParam LnsAny,dddParam LnsAny,threading bool) *Nodes_LiteralStringNode {
    var node *Nodes_LiteralStringNode
    node = NewNodes_LiteralStringNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, token, orgParam, dddParam, threading)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralStringNode.visit
func (self *Nodes_LiteralStringNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    {
        {
            _child := self.orgParam
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch40638 := visitor(&child.Nodes_Node, &self.Nodes_Node, "orgParam", depth); _switch40638 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch40638 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    {
        {
            _child := self.dddParam
            if _child != nil {
                child := _child.(*Nodes_ExpListNode)
                if _switch40696 := visitor(&child.Nodes_Node, &self.Nodes_Node, "dddParam", depth); _switch40696 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(visitor, depth + 1)){
                        return false
                    }
                } else if _switch40696 == Nodes_NodeVisitMode__End {
                    return false
                }
            }
        }
    }
    return true
}

// 2407: decl @lune.@base.@Nodes.LiteralStringNode.getLiteral
func (self *Nodes_LiteralStringNode) GetLiteral()(LnsAny, LnsAny) {
    var txt string
    txt = self.token.Txt
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(txt, "^```", nil, nil))){
        txt = Lns_getVM().String_sub(txt,4, -4)
        
    } else { 
        txt = Lns_getVM().String_sub(txt,2, -2)
        
    }
    {
        _param := self.FP.Get_orgParam()
        if _param != nil {
            param := _param.(*Nodes_ExpListNode)
            var argList *LnsList
            argList = param.FP.Get_expList()
            var paramList *LnsList
            paramList = NewLnsList([]LnsAny{})
            for _, _argNode := range( argList.Items ) {
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                var arg LnsAny
                var mess LnsAny
                arg,mess = argNode.FP.GetLiteral()
                if arg != nil{
                    arg_10564 := arg
                    paramList.Set(paramList.Len() + 1,Nodes_getLiteralObj(arg_10564))
                } else {
                    return nil, mess
                }
            }
            txt = Lns_getVM().String_format(Nodes_convExp42888(txt, Lns_2DDD(paramList.Unpack())))
            
        }
    }
    return &Nodes_Literal__Str{txt}, nil
}

// 2434: decl @lune.@base.@Nodes.LiteralStringNode.setupLiteralTokenList
func (self *Nodes_LiteralStringNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Str, self.token.Txt)
    {
        _param := self.FP.Get_orgParam()
        if _param != nil {
            param := _param.(*Nodes_ExpListNode)
            self.FP.AddTokenList(list, Types_TokenKind__Dlmt, "(")
            for _index, _argNode := range( param.FP.Get_expList().Items ) {
                index := _index + 1
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.AddTokenList(list, Types_TokenKind__Dlmt, ",")
                }
                if Lns_op_not(argNode.FP.SetupLiteralTokenList(list)){
                    return false
                }
            }
            self.FP.AddTokenList(list, Types_TokenKind__Dlmt, ")")
        }
    }
    return true
}


// declaration Class -- LiteralBoolNode
type Nodes_LiteralBoolNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_token() *Types_Token
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralBoolNode struct {
    Nodes_Node
    token *Types_Token
    FP Nodes_LiteralBoolNodeMtd
}
func Nodes_LiteralBoolNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralBoolNode).FP
}
type Nodes_LiteralBoolNodeDownCast interface {
    ToNodes_LiteralBoolNode() *Nodes_LiteralBoolNode
}
func Nodes_LiteralBoolNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralBoolNodeDownCast)
    if ok { return work.ToNodes_LiteralBoolNode() }
    return nil
}
func (obj *Nodes_LiteralBoolNode) ToNodes_LiteralBoolNode() *Nodes_LiteralBoolNode {
    return obj
}
func NewNodes_LiteralBoolNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token) *Nodes_LiteralBoolNode {
    obj := &Nodes_LiteralBoolNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralBoolNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_LiteralBoolNode) Get_token() *Types_Token{ return self.token }
// 1: decl @lune.@base.@Nodes.LiteralBoolNode.processFilter
func (self *Nodes_LiteralBoolNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralBool(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralBoolNode.canBeRight
func (self *Nodes_LiteralBoolNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralBoolNode.canBeLeft
func (self *Nodes_LiteralBoolNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralBoolNode.canBeStatement
func (self *Nodes_LiteralBoolNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralBoolNode) InitNodes_LiteralBoolNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token) {
    self.InitNodes_Node(id, 85, pos, macroArgFlag, typeList)
    self.token = token
    
}

// 681: decl @lune.@base.@Nodes.LiteralBoolNode.create
func Nodes_LiteralBoolNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token) *Nodes_LiteralBoolNode {
    var node *Nodes_LiteralBoolNode
    node = NewNodes_LiteralBoolNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, token)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralBoolNode.visit
func (self *Nodes_LiteralBoolNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 2452: decl @lune.@base.@Nodes.LiteralBoolNode.getLiteral
func (self *Nodes_LiteralBoolNode) GetLiteral()(LnsAny, LnsAny) {
    return &Nodes_Literal__Bool{self.token.Txt == "true"}, nil
}

// 2456: decl @lune.@base.@Nodes.LiteralBoolNode.setupLiteralTokenList
func (self *Nodes_LiteralBoolNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Kywd, self.token.Txt)
    return true
}


// declaration Class -- LiteralSymbolNode
type Nodes_LiteralSymbolNodeMtd interface {
    AddComment(arg1 *LnsList)
    AddTokenList(arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft() bool
    CanBeRight(arg1 *Ast_ProcessInfo) bool
    CanBeStatement() bool
    GetBreakKind(arg1 LnsInt) LnsInt
    GetLiteral()(LnsAny, LnsAny)
    GetPrefix() LnsAny
    GetSymbolInfo() *LnsList
    Get_commentList() LnsAny
    Get_effectivePos() *Types_Position
    Get_expType() *Ast_TypeInfo
    Get_expTypeList() *LnsList
    Get_id() LnsInt
    Get_isLValue() bool
    Get_kind() LnsInt
    Get_macroArgFlag() bool
    Get_pos() *Types_Position
    Get_tailComment() LnsAny
    Get_token() *Types_Token
    HasNilAccess() bool
    IsThreading() bool
    ProcessFilter(arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue()
    Set_tailComment(arg1 LnsAny)
    SetupLiteralTokenList(arg1 *LnsList) bool
    Visit(arg1 Nodes_NodeVisitor, arg2 LnsInt) bool
}
type Nodes_LiteralSymbolNode struct {
    Nodes_Node
    token *Types_Token
    FP Nodes_LiteralSymbolNodeMtd
}
func Nodes_LiteralSymbolNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_LiteralSymbolNode).FP
}
type Nodes_LiteralSymbolNodeDownCast interface {
    ToNodes_LiteralSymbolNode() *Nodes_LiteralSymbolNode
}
func Nodes_LiteralSymbolNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_LiteralSymbolNodeDownCast)
    if ok { return work.ToNodes_LiteralSymbolNode() }
    return nil
}
func (obj *Nodes_LiteralSymbolNode) ToNodes_LiteralSymbolNode() *Nodes_LiteralSymbolNode {
    return obj
}
func NewNodes_LiteralSymbolNode(arg1 LnsInt, arg2 *Types_Position, arg3 bool, arg4 *LnsList, arg5 *Types_Token) *Nodes_LiteralSymbolNode {
    obj := &Nodes_LiteralSymbolNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralSymbolNode(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_LiteralSymbolNode) Get_token() *Types_Token{ return self.token }
// 1: decl @lune.@base.@Nodes.LiteralSymbolNode.processFilter
func (self *Nodes_LiteralSymbolNode) ProcessFilter(filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralSymbol(self, opt)
}

// 1: decl @lune.@base.@Nodes.LiteralSymbolNode.canBeRight
func (self *Nodes_LiteralSymbolNode) CanBeRight(processInfo *Ast_ProcessInfo) bool {
    return true
}

// 1: decl @lune.@base.@Nodes.LiteralSymbolNode.canBeLeft
func (self *Nodes_LiteralSymbolNode) CanBeLeft() bool {
    return false
}

// 1: decl @lune.@base.@Nodes.LiteralSymbolNode.canBeStatement
func (self *Nodes_LiteralSymbolNode) CanBeStatement() bool {
    return false
}

// 676: DeclConstr
func (self *Nodes_LiteralSymbolNode) InitNodes_LiteralSymbolNode(id LnsInt,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token) {
    self.InitNodes_Node(id, 86, pos, macroArgFlag, typeList)
    self.token = token
    
}

// 681: decl @lune.@base.@Nodes.LiteralSymbolNode.create
func Nodes_LiteralSymbolNode_create(nodeMan *Nodes_NodeManager,pos *Types_Position,macroArgFlag bool,typeList *LnsList,token *Types_Token) *Nodes_LiteralSymbolNode {
    var node *Nodes_LiteralSymbolNode
    node = NewNodes_LiteralSymbolNode(nodeMan.FP.NextId(), pos, macroArgFlag, typeList, token)
    nodeMan.FP.AddNode(&node.Nodes_Node)
    return node
}

// 690: decl @lune.@base.@Nodes.LiteralSymbolNode.visit
func (self *Nodes_LiteralSymbolNode) Visit(visitor Nodes_NodeVisitor,depth LnsInt) bool {
    return true
}

// 2463: decl @lune.@base.@Nodes.LiteralSymbolNode.getLiteral
func (self *Nodes_LiteralSymbolNode) GetLiteral()(LnsAny, LnsAny) {
    return &Nodes_Literal__Symbol{self.token.Txt}, nil
}

// 2467: decl @lune.@base.@Nodes.LiteralSymbolNode.setupLiteralTokenList
func (self *Nodes_LiteralSymbolNode) SetupLiteralTokenList(list *LnsList) bool {
    self.FP.AddTokenList(list, Types_TokenKind__Symb, self.token.Txt)
    return true
}


// declaration Class -- DefMacroInfo
type Nodes_DefMacroInfoMtd interface {
    GetArgList() *LnsList
    GetTokenList() *LnsList
    Get_name() string
}
type Nodes_DefMacroInfo struct {
    Nodes_MacroInfo
    DeclInfo *Nodes_DeclMacroInfo
    argList *LnsList
    FP Nodes_DefMacroInfoMtd
}
func Nodes_DefMacroInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DefMacroInfo).FP
}
type Nodes_DefMacroInfoDownCast interface {
    ToNodes_DefMacroInfo() *Nodes_DefMacroInfo
}
func Nodes_DefMacroInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DefMacroInfoDownCast)
    if ok { return work.ToNodes_DefMacroInfo() }
    return nil
}
func (obj *Nodes_DefMacroInfo) ToNodes_DefMacroInfo() *Nodes_DefMacroInfo {
    return obj
}
func NewNodes_DefMacroInfo(arg1 *Lns_luaValue, arg2 *Nodes_DeclMacroInfo, arg3 *LnsMap) *Nodes_DefMacroInfo {
    obj := &Nodes_DefMacroInfo{}
    obj.FP = obj
    obj.Nodes_MacroInfo.FP = obj
    obj.InitNodes_DefMacroInfo(arg1, arg2, arg3)
    return obj
}
// 2727: decl @lune.@base.@Nodes.DefMacroInfo.get_name
func (self *Nodes_DefMacroInfo) Get_name() string {
    return self.DeclInfo.FP.Get_name().Txt
}

// 2731: decl @lune.@base.@Nodes.DefMacroInfo.getArgList
func (self *Nodes_DefMacroInfo) GetArgList() *LnsList {
    return self.argList
}

// 2734: decl @lune.@base.@Nodes.DefMacroInfo.getTokenList
func (self *Nodes_DefMacroInfo) GetTokenList() *LnsList {
    return self.DeclInfo.FP.Get_tokenList()
}

// 2738: DeclConstr
func (self *Nodes_DefMacroInfo) InitNodes_DefMacroInfo(_func *Lns_luaValue,declInfo *Nodes_DeclMacroInfo,symbol2MacroValInfoMap *LnsMap) {
    self.InitNodes_MacroInfo(_func, symbol2MacroValInfoMap)
    self.DeclInfo = declInfo
    
    self.argList = NewLnsList([]LnsAny{})
    
    for _, _argNode := range( declInfo.FP.Get_argList().Items ) {
        argNode := _argNode.(Nodes_DeclArgNodeDownCast).ToNodes_DeclArgNode()
        if argNode.FP.Get_kind() == Nodes_NodeKind_get_DeclArg(){
            var argType *Ast_TypeInfo
            argType = argNode.FP.Get_expType()
            var argName string
            argName = argNode.FP.Get_name().Txt
            self.argList.Insert(Nodes_MacroArgInfo2Stem(NewNodes_MacroArgInfo(argName, argType)))
        }
    }
}


func Lns_Nodes_init() {
    if init_Nodes { return }
    init_Nodes = true
    Nodes__mod__ = "@lune.@base.@Nodes"
    Lns_InitMod()
    Lns_Parser_init()
    Lns_Util_init()
    Lns_frontInterface_init()
    Lns_Ast_init()
    Lns_LuneControl_init()
    Nodes_nodeKind2NameMap = NewLnsMap( map[LnsAny]LnsAny{})
    Nodes_nodeKindSeed = 0
    Nodes_regKind_1338_("None")
    
    Nodes_regKind_1338_("ConvStat")
    
    Nodes_regKind_1338_("BlankLine")
    
    Nodes_regKind_1338_("Subfile")
    
    Nodes_regKind_1338_("Import")
    
    Nodes_regKind_1338_("Root")
    
    Nodes_regKind_1338_("RefType")
    
    Nodes_regKind_1338_("Block")
    
    Nodes_regKind_1338_("Scope")
    
    Nodes_regKind_1338_("If")
    
    Nodes_regKind_1338_("ExpList")
    
    Nodes_regKind_1338_("Switch")
    
    Nodes_regKind_1338_("While")
    
    Nodes_regKind_1338_("Repeat")
    
    
    Nodes_regKind_1338_("For")
    
    
    Nodes_regKind_1338_("Apply")
    
    
    Nodes_regKind_1338_("Foreach")
    
    
    Nodes_regKind_1338_("Forsort")
    
    
    Nodes_regKind_1338_("Return")
    
    Nodes_regKind_1338_("Break")
    
    Nodes_regKind_1338_("Provide")
    
    Nodes_regKind_1338_("ExpNew")
    
    Nodes_regKind_1338_("ExpUnwrap")
    
    Nodes_regKind_1338_("ExpRef")
    
    Nodes_regKind_1338_("ExpSetVal")
    
    Nodes_regKind_1338_("ExpSetItem")
    
    Nodes_regKind_1338_("ExpOp2")
    
    Nodes_regKind_1338_("UnwrapSet")
    
    Nodes_regKind_1338_("IfUnwrap")
    
    Nodes_regKind_1338_("When")
    
    Nodes_regKind_1338_("ExpCast")
    
    Nodes_regKind_1338_("ExpToDDD")
    
    Nodes_regKind_1338_("ExpSubDDD")
    
    Nodes_regKind_1338_("ExpOp1")
    
    Nodes_regKind_1338_("ExpRefItem")
    
    Nodes_regKind_1338_("ExpCall")
    
    Nodes_regKind_1338_("ExpMRet")
    
    Nodes_regKind_1338_("ExpAccessMRet")
    
    Nodes_regKind_1338_("ExpMultiTo1")
    
    Nodes_regKind_1338_("ExpParen")
    
    Nodes_regKind_1338_("ExpMacroExp")
    
    Nodes_regKind_1338_("ExpMacroStat")
    
    Nodes_regKind_1338_("ExpMacroArgExp")
    
    Nodes_regKind_1338_("StmtExp")
    
    Nodes_regKind_1338_("ExpMacroStatList")
    
    Nodes_regKind_1338_("ExpOmitEnum")
    
    Nodes_regKind_1338_("RefField")
    
    Nodes_regKind_1338_("GetField")
    
    Nodes_regKind_1338_("Alias")
    
    Nodes_regKind_1338_("DeclVar")
    
    Nodes_regKind_1338_("DeclForm")
    
    Nodes_regKind_1338_("DeclFunc")
    
    Nodes_regKind_1338_("DeclMethod")
    
    Nodes_regKind_1338_("ProtoMethod")
    
    Nodes_regKind_1338_("DeclConstr")
    
    Nodes_regKind_1338_("DeclDestr")
    
    Nodes_regKind_1338_("ExpCallSuperCtor")
    
    Nodes_regKind_1338_("ExpCallSuper")
    
    Nodes_regKind_1338_("DeclMember")
    
    Nodes_regKind_1338_("DeclArg")
    
    Nodes_regKind_1338_("DeclArgDDD")
    
    Nodes_regKind_1338_("DeclAdvertise")
    
    Nodes_regKind_1338_("ProtoClass")
    
    Nodes_regKind_1338_("DeclClass")
    
    Nodes_regKind_1338_("DeclEnum")
    
    Nodes_regKind_1338_("DeclAlge")
    
    Nodes_regKind_1338_("NewAlgeVal")
    
    Nodes_regKind_1338_("LuneControl")
    
    Nodes_regKind_1338_("Match")
    
    Nodes_regKind_1338_("LuneKind")
    
    Nodes_regKind_1338_("DeclMacro")
    
    Nodes_regKind_1338_("TestCase")
    
    Nodes_regKind_1338_("TestBlock")
    
    Nodes_regKind_1338_("Abbr")
    
    Nodes_regKind_1338_("Boxing")
    
    Nodes_regKind_1338_("Unboxing")
    
    Nodes_regKind_1338_("LiteralNil")
    
    Nodes_regKind_1338_("LiteralChar")
    
    Nodes_regKind_1338_("LiteralInt")
    
    Nodes_regKind_1338_("LiteralReal")
    
    Nodes_regKind_1338_("LiteralArray")
    
    Nodes_regKind_1338_("LiteralList")
    
    Nodes_regKind_1338_("LiteralSet")
    
    Nodes_regKind_1338_("LiteralMap")
    
    Nodes_regKind_1338_("LiteralString")
    
    Nodes_regKind_1338_("LiteralBool")
    
    Nodes_regKind_1338_("LiteralSymbol")
    
    
}
func init() {
    init_Nodes = false
}
