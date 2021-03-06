// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Nodes bool
var Nodes__mod__ string
// decl enum -- BreakKind 
type Nodes_BreakKind = LnsInt
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
func Nodes_BreakKind_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_BreakKindList_
}
var Nodes_BreakKindMap_ = map[LnsInt]string {
  Nodes_BreakKind__Break: "BreakKind.Break",
  Nodes_BreakKind__NeverRet: "BreakKind.NeverRet",
  Nodes_BreakKind__None: "BreakKind.None",
  Nodes_BreakKind__Return: "BreakKind.Return",
}
func Nodes_BreakKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_BreakKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_BreakKind_getTxt(arg1 LnsInt) string {
    return Nodes_BreakKindMap_[arg1];
}
// decl enum -- CheckBreakMode 
type Nodes_CheckBreakMode = LnsInt
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
func Nodes_CheckBreakMode_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_CheckBreakModeList_
}
var Nodes_CheckBreakModeMap_ = map[LnsInt]string {
  Nodes_CheckBreakMode__IgnoreFlow: "CheckBreakMode.IgnoreFlow",
  Nodes_CheckBreakMode__IgnoreFlowReturn: "CheckBreakMode.IgnoreFlowReturn",
  Nodes_CheckBreakMode__Normal: "CheckBreakMode.Normal",
  Nodes_CheckBreakMode__Return: "CheckBreakMode.Return",
}
func Nodes_CheckBreakMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_CheckBreakModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_CheckBreakMode_getTxt(arg1 LnsInt) string {
    return Nodes_CheckBreakModeMap_[arg1];
}
// decl enum -- NodeVisitMode 
type Nodes_NodeVisitMode = LnsInt
const Nodes_NodeVisitMode__Child = 0
const Nodes_NodeVisitMode__End = 2
const Nodes_NodeVisitMode__Next = 1
var Nodes_NodeVisitModeList_ = NewLnsList( []LnsAny {
  Nodes_NodeVisitMode__Child,
  Nodes_NodeVisitMode__Next,
  Nodes_NodeVisitMode__End,
})
func Nodes_NodeVisitMode_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_NodeVisitModeList_
}
var Nodes_NodeVisitModeMap_ = map[LnsInt]string {
  Nodes_NodeVisitMode__Child: "NodeVisitMode.Child",
  Nodes_NodeVisitMode__End: "NodeVisitMode.End",
  Nodes_NodeVisitMode__Next: "NodeVisitMode.Next",
}
func Nodes_NodeVisitMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_NodeVisitModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_NodeVisitMode_getTxt(arg1 LnsInt) string {
    return Nodes_NodeVisitModeMap_[arg1];
}
// decl enum -- LazyLoad 
type Nodes_LazyLoad = LnsInt
const Nodes_LazyLoad__Auto = 2
const Nodes_LazyLoad__Off = 0
const Nodes_LazyLoad__On = 1
var Nodes_LazyLoadList_ = NewLnsList( []LnsAny {
  Nodes_LazyLoad__Off,
  Nodes_LazyLoad__On,
  Nodes_LazyLoad__Auto,
})
func Nodes_LazyLoad_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_LazyLoadList_
}
var Nodes_LazyLoadMap_ = map[LnsInt]string {
  Nodes_LazyLoad__Auto: "LazyLoad.Auto",
  Nodes_LazyLoad__Off: "LazyLoad.Off",
  Nodes_LazyLoad__On: "LazyLoad.On",
}
func Nodes_LazyLoad__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_LazyLoadMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_LazyLoad_getTxt(arg1 LnsInt) string {
    return Nodes_LazyLoadMap_[arg1];
}
// decl enum -- BlockKind 
type Nodes_BlockKind = LnsInt
const Nodes_BlockKind__Apply = 8
const Nodes_BlockKind__AsyncLock = 20
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
  Nodes_BlockKind__AsyncLock,
})
func Nodes_BlockKind_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_BlockKindList_
}
var Nodes_BlockKindMap_ = map[LnsInt]string {
  Nodes_BlockKind__Apply: "BlockKind.Apply",
  Nodes_BlockKind__AsyncLock: "BlockKind.AsyncLock",
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
func Nodes_BlockKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_BlockKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_BlockKind_getTxt(arg1 LnsInt) string {
    return Nodes_BlockKindMap_[arg1];
}
// decl enum -- ScopeKind 
type Nodes_ScopeKind = LnsInt
const Nodes_ScopeKind__Root = 0
var Nodes_ScopeKindList_ = NewLnsList( []LnsAny {
  Nodes_ScopeKind__Root,
})
func Nodes_ScopeKind_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_ScopeKindList_
}
var Nodes_ScopeKindMap_ = map[LnsInt]string {
  Nodes_ScopeKind__Root: "ScopeKind.Root",
}
func Nodes_ScopeKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_ScopeKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_ScopeKind_getTxt(arg1 LnsInt) string {
    return Nodes_ScopeKindMap_[arg1];
}
// decl enum -- IfKind 
type Nodes_IfKind = LnsInt
const Nodes_IfKind__Else = 2
const Nodes_IfKind__ElseIf = 1
const Nodes_IfKind__If = 0
var Nodes_IfKindList_ = NewLnsList( []LnsAny {
  Nodes_IfKind__If,
  Nodes_IfKind__ElseIf,
  Nodes_IfKind__Else,
})
func Nodes_IfKind_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_IfKindList_
}
var Nodes_IfKindMap_ = map[LnsInt]string {
  Nodes_IfKind__Else: "IfKind.Else",
  Nodes_IfKind__ElseIf: "IfKind.ElseIf",
  Nodes_IfKind__If: "IfKind.If",
}
func Nodes_IfKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_IfKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_IfKind_getTxt(arg1 LnsInt) string {
    return Nodes_IfKindMap_[arg1];
}
// decl enum -- CaseKind 
type Nodes_CaseKind = LnsInt
const Nodes_CaseKind__Full = 1
const Nodes_CaseKind__Lack = 0
const Nodes_CaseKind__MustFull = 2
var Nodes_CaseKindList_ = NewLnsList( []LnsAny {
  Nodes_CaseKind__Lack,
  Nodes_CaseKind__Full,
  Nodes_CaseKind__MustFull,
})
func Nodes_CaseKind_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_CaseKindList_
}
var Nodes_CaseKindMap_ = map[LnsInt]string {
  Nodes_CaseKind__Full: "CaseKind.Full",
  Nodes_CaseKind__Lack: "CaseKind.Lack",
  Nodes_CaseKind__MustFull: "CaseKind.MustFull",
}
func Nodes_CaseKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_CaseKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_CaseKind_getTxt(arg1 LnsInt) string {
    return Nodes_CaseKindMap_[arg1];
}
// decl enum -- CastKind 
type Nodes_CastKind = LnsInt
const Nodes_CastKind__Force = 1
const Nodes_CastKind__Implicit = 2
const Nodes_CastKind__Normal = 0
var Nodes_CastKindList_ = NewLnsList( []LnsAny {
  Nodes_CastKind__Normal,
  Nodes_CastKind__Force,
  Nodes_CastKind__Implicit,
})
func Nodes_CastKind_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_CastKindList_
}
var Nodes_CastKindMap_ = map[LnsInt]string {
  Nodes_CastKind__Force: "CastKind.Force",
  Nodes_CastKind__Implicit: "CastKind.Implicit",
  Nodes_CastKind__Normal: "CastKind.Normal",
}
func Nodes_CastKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_CastKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_CastKind_getTxt(arg1 LnsInt) string {
    return Nodes_CastKindMap_[arg1];
}
// decl enum -- MacroMode 
type Nodes_MacroMode = LnsInt
const Nodes_MacroMode__AnalyzeArg = 2
const Nodes_MacroMode__Expand = 1
const Nodes_MacroMode__None = 0
var Nodes_MacroModeList_ = NewLnsList( []LnsAny {
  Nodes_MacroMode__None,
  Nodes_MacroMode__Expand,
  Nodes_MacroMode__AnalyzeArg,
})
func Nodes_MacroMode_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_MacroModeList_
}
var Nodes_MacroModeMap_ = map[LnsInt]string {
  Nodes_MacroMode__AnalyzeArg: "MacroMode.AnalyzeArg",
  Nodes_MacroMode__Expand: "MacroMode.Expand",
  Nodes_MacroMode__None: "MacroMode.None",
}
func Nodes_MacroMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_MacroModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_MacroMode_getTxt(arg1 LnsInt) string {
    return Nodes_MacroModeMap_[arg1];
}
// decl enum -- MacroStatKind 
type Nodes_MacroStatKind = LnsInt
const Nodes_MacroStatKind__Exp = 1
const Nodes_MacroStatKind__Stat = 0
var Nodes_MacroStatKindList_ = NewLnsList( []LnsAny {
  Nodes_MacroStatKind__Stat,
  Nodes_MacroStatKind__Exp,
})
func Nodes_MacroStatKind_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_MacroStatKindList_
}
var Nodes_MacroStatKindMap_ = map[LnsInt]string {
  Nodes_MacroStatKind__Exp: "MacroStatKind.Exp",
  Nodes_MacroStatKind__Stat: "MacroStatKind.Stat",
}
func Nodes_MacroStatKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_MacroStatKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_MacroStatKind_getTxt(arg1 LnsInt) string {
    return Nodes_MacroStatKindMap_[arg1];
}
// decl enum -- DeclVarMode 
type Nodes_DeclVarMode = LnsInt
const Nodes_DeclVarMode__Let = 0
const Nodes_DeclVarMode__Sync = 1
const Nodes_DeclVarMode__Unwrap = 2
var Nodes_DeclVarModeList_ = NewLnsList( []LnsAny {
  Nodes_DeclVarMode__Let,
  Nodes_DeclVarMode__Sync,
  Nodes_DeclVarMode__Unwrap,
})
func Nodes_DeclVarMode_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_DeclVarModeList_
}
var Nodes_DeclVarModeMap_ = map[LnsInt]string {
  Nodes_DeclVarMode__Let: "DeclVarMode.Let",
  Nodes_DeclVarMode__Sync: "DeclVarMode.Sync",
  Nodes_DeclVarMode__Unwrap: "DeclVarMode.Unwrap",
}
func Nodes_DeclVarMode__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_DeclVarModeMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_DeclVarMode_getTxt(arg1 LnsInt) string {
    return Nodes_DeclVarModeMap_[arg1];
}
// decl enum -- FuncKind 
type Nodes_FuncKind = LnsInt
const Nodes_FuncKind__Ctor = 2
const Nodes_FuncKind__Dstr = 3
const Nodes_FuncKind__Form = 5
const Nodes_FuncKind__Func = 0
const Nodes_FuncKind__InitBlock = 4
const Nodes_FuncKind__Mtd = 1
var Nodes_FuncKindList_ = NewLnsList( []LnsAny {
  Nodes_FuncKind__Func,
  Nodes_FuncKind__Mtd,
  Nodes_FuncKind__Ctor,
  Nodes_FuncKind__Dstr,
  Nodes_FuncKind__InitBlock,
  Nodes_FuncKind__Form,
})
func Nodes_FuncKind_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_FuncKindList_
}
var Nodes_FuncKindMap_ = map[LnsInt]string {
  Nodes_FuncKind__Ctor: "FuncKind.Ctor",
  Nodes_FuncKind__Dstr: "FuncKind.Dstr",
  Nodes_FuncKind__Form: "FuncKind.Form",
  Nodes_FuncKind__Func: "FuncKind.Func",
  Nodes_FuncKind__InitBlock: "FuncKind.InitBlock",
  Nodes_FuncKind__Mtd: "FuncKind.Mtd",
}
func Nodes_FuncKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_FuncKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_FuncKind_getTxt(arg1 LnsInt) string {
    return Nodes_FuncKindMap_[arg1];
}
// decl enum -- LockKind 
type Nodes_LockKind = LnsInt
const Nodes_LockKind__AsyncLock = 0
const Nodes_LockKind__LuaDepend = 3
const Nodes_LockKind__LuaGo = 1
const Nodes_LockKind__LuaLock = 2
var Nodes_LockKindList_ = NewLnsList( []LnsAny {
  Nodes_LockKind__AsyncLock,
  Nodes_LockKind__LuaGo,
  Nodes_LockKind__LuaLock,
  Nodes_LockKind__LuaDepend,
})
func Nodes_LockKind_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_LockKindList_
}
var Nodes_LockKindMap_ = map[LnsInt]string {
  Nodes_LockKind__AsyncLock: "LockKind.AsyncLock",
  Nodes_LockKind__LuaDepend: "LockKind.LuaDepend",
  Nodes_LockKind__LuaGo: "LockKind.LuaGo",
  Nodes_LockKind__LuaLock: "LockKind.LuaLock",
}
func Nodes_LockKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_LockKindMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_LockKind_getTxt(arg1 LnsInt) string {
    return Nodes_LockKindMap_[arg1];
}
// decl enum -- nodeKindEnum 
type Nodes_nodeKindEnum = LnsInt
const Nodes_nodeKindEnum__Abbr = 76
const Nodes_nodeKindEnum__Alias = 49
const Nodes_nodeKindEnum__Apply = 16
const Nodes_nodeKindEnum__AsyncLock = 59
const Nodes_nodeKindEnum__BlankLine = 3
const Nodes_nodeKindEnum__Block = 8
const Nodes_nodeKindEnum__Boxing = 77
const Nodes_nodeKindEnum__Break = 20
const Nodes_nodeKindEnum__ConvStat = 2
const Nodes_nodeKindEnum__DeclAdvertise = 64
const Nodes_nodeKindEnum__DeclAlge = 68
const Nodes_nodeKindEnum__DeclArg = 62
const Nodes_nodeKindEnum__DeclArgDDD = 63
const Nodes_nodeKindEnum__DeclClass = 66
const Nodes_nodeKindEnum__DeclConstr = 55
const Nodes_nodeKindEnum__DeclDestr = 56
const Nodes_nodeKindEnum__DeclEnum = 67
const Nodes_nodeKindEnum__DeclForm = 51
const Nodes_nodeKindEnum__DeclFunc = 52
const Nodes_nodeKindEnum__DeclMacro = 73
const Nodes_nodeKindEnum__DeclMember = 61
const Nodes_nodeKindEnum__DeclMethod = 53
const Nodes_nodeKindEnum__DeclVar = 50
const Nodes_nodeKindEnum__ExpAccessMRet = 38
const Nodes_nodeKindEnum__ExpCall = 36
const Nodes_nodeKindEnum__ExpCallSuper = 58
const Nodes_nodeKindEnum__ExpCallSuperCtor = 57
const Nodes_nodeKindEnum__ExpCast = 31
const Nodes_nodeKindEnum__ExpList = 11
const Nodes_nodeKindEnum__ExpMRet = 37
const Nodes_nodeKindEnum__ExpMacroArgExp = 43
const Nodes_nodeKindEnum__ExpMacroExp = 41
const Nodes_nodeKindEnum__ExpMacroStat = 42
const Nodes_nodeKindEnum__ExpMacroStatList = 45
const Nodes_nodeKindEnum__ExpMultiTo1 = 39
const Nodes_nodeKindEnum__ExpNew = 22
const Nodes_nodeKindEnum__ExpOmitEnum = 46
const Nodes_nodeKindEnum__ExpOp1 = 34
const Nodes_nodeKindEnum__ExpOp2 = 27
const Nodes_nodeKindEnum__ExpParen = 40
const Nodes_nodeKindEnum__ExpRef = 24
const Nodes_nodeKindEnum__ExpRefItem = 35
const Nodes_nodeKindEnum__ExpSetItem = 26
const Nodes_nodeKindEnum__ExpSetVal = 25
const Nodes_nodeKindEnum__ExpSubDDD = 33
const Nodes_nodeKindEnum__ExpToDDD = 32
const Nodes_nodeKindEnum__ExpUnwrap = 23
const Nodes_nodeKindEnum__For = 15
const Nodes_nodeKindEnum__Foreach = 17
const Nodes_nodeKindEnum__Forsort = 18
const Nodes_nodeKindEnum__GetField = 48
const Nodes_nodeKindEnum__If = 10
const Nodes_nodeKindEnum__IfUnwrap = 29
const Nodes_nodeKindEnum__Import = 5
const Nodes_nodeKindEnum__LiteralArray = 83
const Nodes_nodeKindEnum__LiteralBool = 88
const Nodes_nodeKindEnum__LiteralChar = 80
const Nodes_nodeKindEnum__LiteralInt = 81
const Nodes_nodeKindEnum__LiteralList = 84
const Nodes_nodeKindEnum__LiteralMap = 86
const Nodes_nodeKindEnum__LiteralNil = 79
const Nodes_nodeKindEnum__LiteralReal = 82
const Nodes_nodeKindEnum__LiteralSet = 85
const Nodes_nodeKindEnum__LiteralString = 87
const Nodes_nodeKindEnum__LiteralSymbol = 89
const Nodes_nodeKindEnum__LuneControl = 70
const Nodes_nodeKindEnum__LuneKind = 72
const Nodes_nodeKindEnum__Match = 71
const Nodes_nodeKindEnum__NewAlgeVal = 69
const Nodes_nodeKindEnum__None = 0
const Nodes_nodeKindEnum__ProtoClass = 65
const Nodes_nodeKindEnum__ProtoMethod = 54
const Nodes_nodeKindEnum__Provide = 21
const Nodes_nodeKindEnum__RefField = 47
const Nodes_nodeKindEnum__RefType = 7
const Nodes_nodeKindEnum__Repeat = 14
const Nodes_nodeKindEnum__Request = 60
const Nodes_nodeKindEnum__Return = 19
const Nodes_nodeKindEnum__Root = 6
const Nodes_nodeKindEnum__Scope = 9
const Nodes_nodeKindEnum__Shebang = 1
const Nodes_nodeKindEnum__StmtExp = 44
const Nodes_nodeKindEnum__Subfile = 4
const Nodes_nodeKindEnum__Switch = 12
const Nodes_nodeKindEnum__TestBlock = 75
const Nodes_nodeKindEnum__TestCase = 74
const Nodes_nodeKindEnum__Unboxing = 78
const Nodes_nodeKindEnum__UnwrapSet = 28
const Nodes_nodeKindEnum__When = 30
const Nodes_nodeKindEnum__While = 13
var Nodes_nodeKindEnumList_ = NewLnsList( []LnsAny {
  Nodes_nodeKindEnum__None,
  Nodes_nodeKindEnum__Shebang,
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
  Nodes_nodeKindEnum__AsyncLock,
  Nodes_nodeKindEnum__Request,
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
func Nodes_nodeKindEnum_get__allList(_env *LnsEnv) *LnsList{
    return Nodes_nodeKindEnumList_
}
var Nodes_nodeKindEnumMap_ = map[LnsInt]string {
  Nodes_nodeKindEnum__Abbr: "nodeKindEnum.Abbr",
  Nodes_nodeKindEnum__Alias: "nodeKindEnum.Alias",
  Nodes_nodeKindEnum__Apply: "nodeKindEnum.Apply",
  Nodes_nodeKindEnum__AsyncLock: "nodeKindEnum.AsyncLock",
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
  Nodes_nodeKindEnum__Request: "nodeKindEnum.Request",
  Nodes_nodeKindEnum__Return: "nodeKindEnum.Return",
  Nodes_nodeKindEnum__Root: "nodeKindEnum.Root",
  Nodes_nodeKindEnum__Scope: "nodeKindEnum.Scope",
  Nodes_nodeKindEnum__Shebang: "nodeKindEnum.Shebang",
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
func Nodes_nodeKindEnum__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Nodes_nodeKindEnumMap_[arg1]; ok { return arg1 }
    return nil
}

func Nodes_nodeKindEnum_getTxt(arg1 LnsInt) string {
    return Nodes_nodeKindEnumMap_[arg1];
}
var Nodes_nodeKind2NameMapWork *LnsMap
var Nodes_nodeKind2NameMap *LnsMap
var Nodes_nodeKindSeed LnsInt
// decl alge -- Literal
type Nodes_Literal = LnsAny
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
// decl alge -- MacroStatement
type Nodes_MacroStatement = LnsAny
type Nodes_MacroStatement__Import struct{
Val1 string
Val2 *LnsList
Val3 LnsAny
Val4 LnsAny
}
func (self *Nodes_MacroStatement__Import) GetTxt() string {
return "MacroStatement.Import"
}
type Nodes_MacroStatement__Local struct{
Val1 *Lns_luaValue
}
func (self *Nodes_MacroStatement__Local) GetTxt() string {
return "MacroStatement.Local"
}
// decl alge -- IndexVal
type Nodes_IndexVal = LnsAny
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
type Nodes_NodeVisitor func (_env *LnsEnv, arg1 *Nodes_Node,arg2 *Nodes_Node,arg3 string,arg4 LnsInt) LnsInt
type Nodes_macroStatmentProc func (_env *LnsEnv, arg1 *LnsMap) *LnsMap
// for 2591
func Nodes_convExp0_29460(arg1 string, arg2 []LnsAny) (string, []LnsAny) {
    return arg1, Lns_2DDD( arg2[0:])
}
// for 580
func Nodes_convExp0_947(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2413
func Nodes_convExp0_28696(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2446
func Nodes_convExp0_28833(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2480
func Nodes_convExp0_28970(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2512
func Nodes_convExp0_29104(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2513
func Nodes_convExp0_29112(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2535
func Nodes_convExp0_29202(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2583
func Nodes_convExp0_29413(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2661
func Nodes_convExp0_29736(arg1 []LnsAny) (LnsAny, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ), Lns_getFromMulti( arg1, 1 )
}
// for 2692
func Nodes_convExp0_29856(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2752
func Nodes_convExp0_30087(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2802
func Nodes_convExp0_30211(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2828
func Nodes_convExp0_30309(arg1 []LnsAny) (bool, LnsInt, LnsReal, string, *Ast_TypeInfo) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 ).(LnsInt), Lns_getFromMulti( arg1, 2 ).(LnsReal), Lns_getFromMulti( arg1, 3 ).(string), Lns_getFromMulti( arg1, 4 ).(*Ast_TypeInfo)
}
// for 2829
func Nodes_convExp0_30325(arg1 []LnsAny) (bool, LnsInt, LnsReal, string, *Ast_TypeInfo) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 ).(LnsInt), Lns_getFromMulti( arg1, 2 ).(LnsReal), Lns_getFromMulti( arg1, 3 ).(string), Lns_getFromMulti( arg1, 4 ).(*Ast_TypeInfo)
}

// 144: decl @lune.@base.@Nodes.getLiteralObj
func Nodes_getLiteralObj(_env *LnsEnv, obj LnsAny) LnsAny {
    switch _matchExp0 := obj.(type) {
    case *Nodes_Literal__Nil:
        return nil
    case *Nodes_Literal__Int:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__Real:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__Str:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__Bool:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__Symbol:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__Field:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__LIST:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__ARRAY:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__SET:
    val := _matchExp0.Val1
        return val
    case *Nodes_Literal__MAP:
    val := _matchExp0.Val1
        return val
    }
// insert a dummy
    return nil
}

// 397: decl @lune.@base.@Nodes.regKind
func Nodes_regKind_17_(_env *LnsEnv, name string) LnsInt {
    var kind LnsInt
    kind = Nodes_nodeKindSeed
    Nodes_nodeKindSeed = Nodes_nodeKindSeed + 1
    Nodes_nodeKind2NameMapWork.Set(kind,name)
    return kind
}

// 404: decl @lune.@base.@Nodes.getNodeKindName
func Nodes_getNodeKindName(_env *LnsEnv, kind LnsInt) string {
    return Lns_unwrap( Nodes_nodeKind2NameMap.Get(kind)).(string)
}


// 963: decl @lune.@base.@Nodes.getBreakKindForStmtList
func Nodes_getBreakKindForStmtList_42_(_env *LnsEnv, checkMode LnsInt,stmtList *LnsList) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        var kind LnsInt
        kind = Nodes_BreakKind__None
        for _, _stmt := range( stmtList.Items ) {
            stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
            if stmt.FP.Get_kind(_env) != Nodes_NodeKind_get_BlankLine(_env){
                var work LnsInt
                work = stmt.FP.GetBreakKind(_env, checkMode)
                if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                    if work == Nodes_BreakKind__Return{
                        return Nodes_BreakKind__Return
                    }
                    if work == Nodes_BreakKind__NeverRet{
                        return Nodes_BreakKind__NeverRet
                    }
                } else { 
                    if _switch0 := work; _switch0 == Nodes_BreakKind__None {
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                            _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                            if false{
                                return Nodes_BreakKind__None
                            }
                        }
                    } else {
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                            _env.SetStackVal( kind > work) ).(bool){
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
            var _forFrom0 LnsInt = stmtList.Len()
            var _forTo0 LnsInt = 1
            _forWork0 := _forFrom0
            _forDelta0 := -1
            for {
                if _forDelta0 > 0 {
                   if _forWork0 > _forTo0 { break }
                } else {
                   if _forWork0 < _forTo0 { break }
                }
                index := _forWork0
                var stmt *Nodes_Node
                stmt = stmtList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                if stmt.FP.Get_kind(_env) != Nodes_NodeKind_get_BlankLine(_env){
                    return stmt.FP.GetBreakKind(_env, checkMode)
                }
                _forWork0 += _forDelta0
            }
        }
    }
    return Nodes_BreakKind__None
}

// 2636: decl @lune.@base.@Nodes.enumLiteral2Literal
func Nodes_enumLiteral2Literal_142_(_env *LnsEnv, obj LnsAny)(LnsAny, LnsAny) {
    switch _matchExp0 := obj.(type) {
    case *Ast_EnumLiteral__Int:
    val := _matchExp0.Val1
        return &Nodes_Literal__Int{val}, nil
    case *Ast_EnumLiteral__Real:
    val := _matchExp0.Val1
        return &Nodes_Literal__Real{val}, nil
    case *Ast_EnumLiteral__Str:
    val := _matchExp0.Val1
        return &Nodes_Literal__Str{val}, nil
    }
// insert a dummy
    return nil,nil
}

// 2963: decl @lune.@base.@Nodes.hasMultiValNode
func Nodes_hasMultiValNode(_env *LnsEnv, node *Nodes_Node) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_expTypeList(_env).Len() > 1) ||
        _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)
}

// 2984: decl @lune.@base.@Nodes.getUnwraped
func Nodes_getUnwraped(_env *LnsEnv, node *Nodes_Node) *Nodes_Node {
    {
        _work := Nodes_ExpMRetNodeDownCastF(node.FP)
        if !Lns_IsNil( _work ) {
            work := _work.(*Nodes_ExpMRetNode)
            return Nodes_getUnwraped(_env, work.FP.Get_mRet(_env))
        }
    }
    {
        _work := Nodes_ExpParenNodeDownCastF(node.FP)
        if !Lns_IsNil( _work ) {
            work := _work.(*Nodes_ExpParenNode)
            return Nodes_getUnwraped(_env, work.FP.Get_exp(_env))
        }
    }
    return node
}

// 2994: decl @lune.@base.@Nodes.getCastUnwraped
func Nodes_getCastUnwraped(_env *LnsEnv, node *Nodes_Node) *Nodes_Node {
    {
        _work := Nodes_ExpCastNodeDownCastF(node.FP)
        if !Lns_IsNil( _work ) {
            work := _work.(*Nodes_ExpCastNode)
            return Nodes_getUnwraped(_env, work.FP.Get_exp(_env))
        }
    }
    return node
}

// 2260: decl @lune.@base.@Nodes.Node.getSymbolInfo.processExpNode
func Node_getSymbolInfo__processExpNode_0_(_env *LnsEnv, node *Nodes_Node) *LnsList {
    if _switch0 := (node.FP.Get_kind(_env)); _switch0 == Nodes_NodeKind_get_ExpRef(_env) {
        return NewLnsList([]LnsAny{Ast_SymbolInfo2Stem((Lns_unwrap( (Nodes_ExpRefNodeDownCastF(node.FP))).(*Nodes_ExpRefNode)).FP.Get_symbolInfo(_env))})
    } else if _switch0 == Nodes_NodeKind_get_RefField(_env) {
        {
            _refFieldNode := Nodes_RefFieldNodeDownCastF(node.FP)
            if !Lns_IsNil( _refFieldNode ) {
                refFieldNode := _refFieldNode.(*Nodes_RefFieldNode)
                if refFieldNode.FP.Get_nilAccess(_env){
                    return NewLnsList([]LnsAny{})
                }
                {
                    __exp := refFieldNode.FP.Get_symbolInfo(_env)
                    if !Lns_IsNil( __exp ) {
                        _exp := __exp.(*Ast_SymbolInfo)
                        return NewLnsList([]LnsAny{Ast_SymbolInfo2Stem(_exp)})
                    }
                }
            }
        }
    } else if _switch0 == Nodes_NodeKind_get_ExpList(_env) {
        {
            _expListNode := Nodes_ExpListNodeDownCastF(node.FP)
            if !Lns_IsNil( _expListNode ) {
                expListNode := _expListNode.(*Nodes_ExpListNode)
                var list *LnsList
                list = NewLnsList([]LnsAny{})
                for _index, _expNode := range( expListNode.FP.Get_expList(_env).Items ) {
                    index := _index + 1
                    expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                    if index == expListNode.FP.Get_expList(_env).Len(){
                        for _, _symbolInfo := range( Node_getSymbolInfo__processExpNode_0_(_env, expNode).Items ) {
                            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            list.Insert(Ast_SymbolInfo2Stem(symbolInfo))
                        }
                    } else { 
                        for _, _symbolInfo := range( Node_getSymbolInfo__processExpNode_0_(_env, expNode).Items ) {
                            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            list.Insert(Ast_SymbolInfo2Stem(symbolInfo))
                            break
                        }
                    }
                }
                return list
            }
        }
    } else if _switch0 == Nodes_NodeKind_get_RefType(_env) {
        {
            _refTypeNode := Nodes_RefTypeNodeDownCastF(node.FP)
            if !Lns_IsNil( _refTypeNode ) {
                refTypeNode := _refTypeNode.(*Nodes_RefTypeNode)
                return refTypeNode.FP.Get_name(_env).FP.GetSymbolInfo(_env)
            }
        }
    }
    return NewLnsList([]LnsAny{})
}

// declaration Class -- SimpleModuleInfoManager
type Nodes_SimpleModuleInfoManagerMtd interface {
    GetModuleInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny
    Pop(_env *LnsEnv)
    Push(_env *LnsEnv, arg1 Ast_ModuleInfoManager)
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
func NewNodes_SimpleModuleInfoManager(_env *LnsEnv, arg1 LnsAny) *Nodes_SimpleModuleInfoManager {
    obj := &Nodes_SimpleModuleInfoManager{}
    obj.FP = obj
    obj.InitNodes_SimpleModuleInfoManager(_env, arg1)
    return obj
}
// advertise -- 35
func (self *Nodes_SimpleModuleInfoManager) GetModuleInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny {
    return self.ModuleInfoManager. GetModuleInfo( _env, arg1)
}
// 40: DeclConstr
func (self *Nodes_SimpleModuleInfoManager) InitNodes_SimpleModuleInfoManager(_env *LnsEnv, moduleInfoManager LnsAny) {
    if moduleInfoManager != nil{
        moduleInfoManager_21 := moduleInfoManager.(Ast_ModuleInfoManager)
        self.ModuleInfoManager = moduleInfoManager_21
    } else {
        self.ModuleInfoManager = Ast_DummyModuleInfoManager_get_instance(_env).FP
    }
    self.moduleInfoManagerHist = NewLnsList([]LnsAny{})
}


// declaration Class -- Filter
type Nodes_FilterMtd interface {
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    GetFull(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl
    popOpt(_env *LnsEnv, arg1 LnsAny)
    ProcessAbbr(_env *LnsEnv, arg1 *Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(_env *LnsEnv, arg1 *Nodes_AliasNode, arg2 LnsAny)
    ProcessApply(_env *LnsEnv, arg1 *Nodes_ApplyNode, arg2 LnsAny)
    ProcessAsyncLock(_env *LnsEnv, arg1 *Nodes_AsyncLockNode, arg2 LnsAny)
    ProcessBlankLine(_env *LnsEnv, arg1 *Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(_env *LnsEnv, arg1 *Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(_env *LnsEnv, arg1 *Nodes_BreakNode, arg2 LnsAny)
    ProcessConvStat(_env *LnsEnv, arg1 *Nodes_ConvStatNode, arg2 LnsAny)
    ProcessDeclAdvertise(_env *LnsEnv, arg1 *Nodes_DeclAdvertiseNode, arg2 LnsAny)
    ProcessDeclAlge(_env *LnsEnv, arg1 *Nodes_DeclAlgeNode, arg2 LnsAny)
    ProcessDeclArg(_env *LnsEnv, arg1 *Nodes_DeclArgNode, arg2 LnsAny)
    ProcessDeclArgDDD(_env *LnsEnv, arg1 *Nodes_DeclArgDDDNode, arg2 LnsAny)
    ProcessDeclClass(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 LnsAny)
    ProcessDeclConstr(_env *LnsEnv, arg1 *Nodes_DeclConstrNode, arg2 LnsAny)
    ProcessDeclDestr(_env *LnsEnv, arg1 *Nodes_DeclDestrNode, arg2 LnsAny)
    ProcessDeclEnum(_env *LnsEnv, arg1 *Nodes_DeclEnumNode, arg2 LnsAny)
    ProcessDeclForm(_env *LnsEnv, arg1 *Nodes_DeclFormNode, arg2 LnsAny)
    ProcessDeclFunc(_env *LnsEnv, arg1 *Nodes_DeclFuncNode, arg2 LnsAny)
    ProcessDeclMacro(_env *LnsEnv, arg1 *Nodes_DeclMacroNode, arg2 LnsAny)
    ProcessDeclMember(_env *LnsEnv, arg1 *Nodes_DeclMemberNode, arg2 LnsAny)
    ProcessDeclMethod(_env *LnsEnv, arg1 *Nodes_DeclMethodNode, arg2 LnsAny)
    ProcessDeclVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode, arg2 LnsAny)
    ProcessExpAccessMRet(_env *LnsEnv, arg1 *Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(_env *LnsEnv, arg1 *Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(_env *LnsEnv, arg1 *Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(_env *LnsEnv, arg1 *Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpList(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 LnsAny)
    ProcessExpMRet(_env *LnsEnv, arg1 *Nodes_ExpMRetNode, arg2 LnsAny)
    ProcessExpMacroArgExp(_env *LnsEnv, arg1 *Nodes_ExpMacroArgExpNode, arg2 LnsAny)
    ProcessExpMacroExp(_env *LnsEnv, arg1 *Nodes_ExpMacroExpNode, arg2 LnsAny)
    ProcessExpMacroStat(_env *LnsEnv, arg1 *Nodes_ExpMacroStatNode, arg2 LnsAny)
    ProcessExpMacroStatList(_env *LnsEnv, arg1 *Nodes_ExpMacroStatListNode, arg2 LnsAny)
    ProcessExpMultiTo1(_env *LnsEnv, arg1 *Nodes_ExpMultiTo1Node, arg2 LnsAny)
    ProcessExpNew(_env *LnsEnv, arg1 *Nodes_ExpNewNode, arg2 LnsAny)
    ProcessExpOmitEnum(_env *LnsEnv, arg1 *Nodes_ExpOmitEnumNode, arg2 LnsAny)
    ProcessExpOp1(_env *LnsEnv, arg1 *Nodes_ExpOp1Node, arg2 LnsAny)
    ProcessExpOp2(_env *LnsEnv, arg1 *Nodes_ExpOp2Node, arg2 LnsAny)
    ProcessExpParen(_env *LnsEnv, arg1 *Nodes_ExpParenNode, arg2 LnsAny)
    ProcessExpRef(_env *LnsEnv, arg1 *Nodes_ExpRefNode, arg2 LnsAny)
    ProcessExpRefItem(_env *LnsEnv, arg1 *Nodes_ExpRefItemNode, arg2 LnsAny)
    ProcessExpSetItem(_env *LnsEnv, arg1 *Nodes_ExpSetItemNode, arg2 LnsAny)
    ProcessExpSetVal(_env *LnsEnv, arg1 *Nodes_ExpSetValNode, arg2 LnsAny)
    ProcessExpSubDDD(_env *LnsEnv, arg1 *Nodes_ExpSubDDDNode, arg2 LnsAny)
    ProcessExpToDDD(_env *LnsEnv, arg1 *Nodes_ExpToDDDNode, arg2 LnsAny)
    ProcessExpUnwrap(_env *LnsEnv, arg1 *Nodes_ExpUnwrapNode, arg2 LnsAny)
    ProcessFor(_env *LnsEnv, arg1 *Nodes_ForNode, arg2 LnsAny)
    ProcessForeach(_env *LnsEnv, arg1 *Nodes_ForeachNode, arg2 LnsAny)
    ProcessForsort(_env *LnsEnv, arg1 *Nodes_ForsortNode, arg2 LnsAny)
    ProcessGetField(_env *LnsEnv, arg1 *Nodes_GetFieldNode, arg2 LnsAny)
    ProcessIf(_env *LnsEnv, arg1 *Nodes_IfNode, arg2 LnsAny)
    ProcessIfUnwrap(_env *LnsEnv, arg1 *Nodes_IfUnwrapNode, arg2 LnsAny)
    ProcessImport(_env *LnsEnv, arg1 *Nodes_ImportNode, arg2 LnsAny)
    ProcessLiteralArray(_env *LnsEnv, arg1 *Nodes_LiteralArrayNode, arg2 LnsAny)
    ProcessLiteralBool(_env *LnsEnv, arg1 *Nodes_LiteralBoolNode, arg2 LnsAny)
    ProcessLiteralChar(_env *LnsEnv, arg1 *Nodes_LiteralCharNode, arg2 LnsAny)
    ProcessLiteralInt(_env *LnsEnv, arg1 *Nodes_LiteralIntNode, arg2 LnsAny)
    ProcessLiteralList(_env *LnsEnv, arg1 *Nodes_LiteralListNode, arg2 LnsAny)
    ProcessLiteralMap(_env *LnsEnv, arg1 *Nodes_LiteralMapNode, arg2 LnsAny)
    ProcessLiteralNil(_env *LnsEnv, arg1 *Nodes_LiteralNilNode, arg2 LnsAny)
    ProcessLiteralReal(_env *LnsEnv, arg1 *Nodes_LiteralRealNode, arg2 LnsAny)
    ProcessLiteralSet(_env *LnsEnv, arg1 *Nodes_LiteralSetNode, arg2 LnsAny)
    ProcessLiteralString(_env *LnsEnv, arg1 *Nodes_LiteralStringNode, arg2 LnsAny)
    ProcessLiteralSymbol(_env *LnsEnv, arg1 *Nodes_LiteralSymbolNode, arg2 LnsAny)
    ProcessLuneControl(_env *LnsEnv, arg1 *Nodes_LuneControlNode, arg2 LnsAny)
    ProcessLuneKind(_env *LnsEnv, arg1 *Nodes_LuneKindNode, arg2 LnsAny)
    ProcessMatch(_env *LnsEnv, arg1 *Nodes_MatchNode, arg2 LnsAny)
    ProcessNewAlgeVal(_env *LnsEnv, arg1 *Nodes_NewAlgeValNode, arg2 LnsAny)
    ProcessNone(_env *LnsEnv, arg1 *Nodes_NoneNode, arg2 LnsAny)
    ProcessProtoClass(_env *LnsEnv, arg1 *Nodes_ProtoClassNode, arg2 LnsAny)
    ProcessProtoMethod(_env *LnsEnv, arg1 *Nodes_ProtoMethodNode, arg2 LnsAny)
    ProcessProvide(_env *LnsEnv, arg1 *Nodes_ProvideNode, arg2 LnsAny)
    ProcessRefField(_env *LnsEnv, arg1 *Nodes_RefFieldNode, arg2 LnsAny)
    ProcessRefType(_env *LnsEnv, arg1 *Nodes_RefTypeNode, arg2 LnsAny)
    ProcessRepeat(_env *LnsEnv, arg1 *Nodes_RepeatNode, arg2 LnsAny)
    ProcessRequest(_env *LnsEnv, arg1 *Nodes_RequestNode, arg2 LnsAny)
    ProcessReturn(_env *LnsEnv, arg1 *Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(_env *LnsEnv, arg1 *Nodes_RootNode, arg2 LnsAny)
    ProcessScope(_env *LnsEnv, arg1 *Nodes_ScopeNode, arg2 LnsAny)
    ProcessShebang(_env *LnsEnv, arg1 *Nodes_ShebangNode, arg2 LnsAny)
    ProcessStmtExp(_env *LnsEnv, arg1 *Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(_env *LnsEnv, arg1 *Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(_env *LnsEnv, arg1 *Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(_env *LnsEnv, arg1 *Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(_env *LnsEnv, arg1 *Nodes_TestCaseNode, arg2 LnsAny)
    ProcessUnboxing(_env *LnsEnv, arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(_env *LnsEnv, arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(_env *LnsEnv, arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(_env *LnsEnv, arg1 *Nodes_WhileNode, arg2 LnsAny)
    pushOpt(_env *LnsEnv, arg1 LnsAny)
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
func NewNodes_Filter(_env *LnsEnv, arg1 bool, arg2 LnsAny, arg3 LnsAny) *Nodes_Filter {
    obj := &Nodes_Filter{}
    obj.FP = obj
    obj.InitNodes_Filter(_env, arg1, arg2, arg3)
    return obj
}
func (self *Nodes_Filter) Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl{ return self.typeNameCtrl }
func (self *Nodes_Filter) Get_optStack(_env *LnsEnv) *LnsList{ return self.optStack }
// 69: DeclConstr
func (self *Nodes_Filter) InitNodes_Filter(_env *LnsEnv, errorOnDefault bool,moduleTypeInfo LnsAny,moduleInfoManager LnsAny) {
    self.errorOnDefault = errorOnDefault
    self.moduleInfoManager = NewNodes_SimpleModuleInfoManager(_env, moduleInfoManager)
    var process func(_env *LnsEnv) *Ast_TypeNameCtrl
    process = func(_env *LnsEnv) *Ast_TypeNameCtrl {
        if moduleTypeInfo != nil{
            moduleTypeInfo_51 := moduleTypeInfo.(*Ast_TypeInfo)
            return NewAst_TypeNameCtrl(_env, moduleTypeInfo_51)
        }
        return Ast_defaultTypeNameCtrl
    }
    self.typeNameCtrl = process(_env)
    self.optStack = NewLnsList([]LnsAny{})
}



// declaration Class -- Node
type Nodes_NodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_Node struct {
    managerId LnsInt
    id LnsInt
    kind LnsInt
    pos Types_Position
    expTypeList *LnsList
    commentList LnsAny
    tailComment LnsAny
    IsLValue bool
    macroArgFlag bool
    inTestBlock bool
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
func (self *Nodes_Node) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *Nodes_Node) Get_pos(_env *LnsEnv) Types_Position{ return self.pos }
func (self *Nodes_Node) Get_expTypeList(_env *LnsEnv) *LnsList{ return self.expTypeList }
func (self *Nodes_Node) Get_commentList(_env *LnsEnv) LnsAny{ return self.commentList }
func (self *Nodes_Node) Get_tailComment(_env *LnsEnv) LnsAny{ return self.tailComment }
func (self *Nodes_Node) Set_tailComment(_env *LnsEnv, arg1 LnsAny){ self.tailComment = arg1 }
func (self *Nodes_Node) Get_isLValue(_env *LnsEnv) bool{ return self.IsLValue }
func (self *Nodes_Node) Get_macroArgFlag(_env *LnsEnv) bool{ return self.macroArgFlag }
func (self *Nodes_Node) Get_inTestBlock(_env *LnsEnv) bool{ return self.inTestBlock }
// 271: DeclConstr
func (self *Nodes_Node) InitNodes_Node(_env *LnsEnv, managerId LnsInt,id LnsInt,kind LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,expTypeList *LnsList) {
    self.IsLValue = false
    self.managerId = managerId
    self.id = id
    self.kind = kind
    self.pos = pos
    self.inTestBlock = inTestBlock
    self.expTypeList = expTypeList
    self.commentList = nil
    self.tailComment = nil
    self.macroArgFlag = macroArgFlag
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
func NewNodes_NamespaceInfo(_env *LnsEnv, arg1 string, arg2 *Ast_Scope, arg3 *Ast_TypeInfo) *Nodes_NamespaceInfo {
    obj := &Nodes_NamespaceInfo{}
    obj.FP = obj
    obj.InitNodes_NamespaceInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Nodes_NamespaceInfo) InitNodes_NamespaceInfo(_env *LnsEnv, arg1 string, arg2 *Ast_Scope, arg3 *Ast_TypeInfo) {
    self.Name = arg1
    self.Scope = arg2
    self.TypeInfo = arg3
}

// declaration Class -- DeclMacroInfo
type Nodes_DeclMacroInfoMtd interface {
    Get_argList(_env *LnsEnv) *LnsList
    Get_name(_env *LnsEnv) *Types_Token
    Get_pubFlag(_env *LnsEnv) bool
    Get_stmtBlock(_env *LnsEnv) LnsAny
    Get_tokenList(_env *LnsEnv) *LnsList
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
func NewNodes_DeclMacroInfo(_env *LnsEnv, arg1 bool, arg2 *Types_Token, arg3 *LnsList, arg4 LnsAny, arg5 *LnsList) *Nodes_DeclMacroInfo {
    obj := &Nodes_DeclMacroInfo{}
    obj.FP = obj
    obj.InitNodes_DeclMacroInfo(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *Nodes_DeclMacroInfo) InitNodes_DeclMacroInfo(_env *LnsEnv, arg1 bool, arg2 *Types_Token, arg3 *LnsList, arg4 LnsAny, arg5 *LnsList) {
    self.pubFlag = arg1
    self.name = arg2
    self.argList = arg3
    self.stmtBlock = arg4
    self.tokenList = arg5
}
func (self *Nodes_DeclMacroInfo) Get_pubFlag(_env *LnsEnv) bool{ return self.pubFlag }
func (self *Nodes_DeclMacroInfo) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_DeclMacroInfo) Get_argList(_env *LnsEnv) *LnsList{ return self.argList }
func (self *Nodes_DeclMacroInfo) Get_stmtBlock(_env *LnsEnv) LnsAny{ return self.stmtBlock }
func (self *Nodes_DeclMacroInfo) Get_tokenList(_env *LnsEnv) *LnsList{ return self.tokenList }

// declaration Class -- NodeManager
type Nodes_NodeManagerMtd interface {
    MultiTo1(_env *LnsEnv, arg1 *Nodes_Node) *Nodes_Node
    AddFrom(_env *LnsEnv, arg1 *Nodes_NodeManager)
    AddNode(_env *LnsEnv, arg1 *Nodes_Node)
    DelNode(_env *LnsEnv, arg1 *Nodes_Node)
    GetAbbrNodeList(_env *LnsEnv) *LnsList
    GetAliasNodeList(_env *LnsEnv) *LnsList
    GetApplyNodeList(_env *LnsEnv) *LnsList
    GetAsyncLockNodeList(_env *LnsEnv) *LnsList
    GetBlankLineNodeList(_env *LnsEnv) *LnsList
    GetBlockNodeList(_env *LnsEnv) *LnsList
    GetBoxingNodeList(_env *LnsEnv) *LnsList
    GetBreakNodeList(_env *LnsEnv) *LnsList
    GetConvStatNodeList(_env *LnsEnv) *LnsList
    GetDeclAdvertiseNodeList(_env *LnsEnv) *LnsList
    GetDeclAlgeNodeList(_env *LnsEnv) *LnsList
    GetDeclArgDDDNodeList(_env *LnsEnv) *LnsList
    GetDeclArgNodeList(_env *LnsEnv) *LnsList
    GetDeclClassNodeList(_env *LnsEnv) *LnsList
    GetDeclConstrNodeList(_env *LnsEnv) *LnsList
    GetDeclDestrNodeList(_env *LnsEnv) *LnsList
    GetDeclEnumNodeList(_env *LnsEnv) *LnsList
    GetDeclFormNodeList(_env *LnsEnv) *LnsList
    GetDeclFuncNodeList(_env *LnsEnv) *LnsList
    GetDeclMacroNodeList(_env *LnsEnv) *LnsList
    GetDeclMemberNodeList(_env *LnsEnv) *LnsList
    GetDeclMethodNodeList(_env *LnsEnv) *LnsList
    GetDeclVarNodeList(_env *LnsEnv) *LnsList
    GetExpAccessMRetNodeList(_env *LnsEnv) *LnsList
    GetExpCallNodeList(_env *LnsEnv) *LnsList
    GetExpCallSuperCtorNodeList(_env *LnsEnv) *LnsList
    GetExpCallSuperNodeList(_env *LnsEnv) *LnsList
    GetExpCastNodeList(_env *LnsEnv) *LnsList
    GetExpListNodeList(_env *LnsEnv) *LnsList
    GetExpMRetNodeList(_env *LnsEnv) *LnsList
    GetExpMacroArgExpNodeList(_env *LnsEnv) *LnsList
    GetExpMacroExpNodeList(_env *LnsEnv) *LnsList
    GetExpMacroStatListNodeList(_env *LnsEnv) *LnsList
    GetExpMacroStatNodeList(_env *LnsEnv) *LnsList
    GetExpMultiTo1NodeList(_env *LnsEnv) *LnsList
    GetExpNewNodeList(_env *LnsEnv) *LnsList
    GetExpOmitEnumNodeList(_env *LnsEnv) *LnsList
    GetExpOp1NodeList(_env *LnsEnv) *LnsList
    GetExpOp2NodeList(_env *LnsEnv) *LnsList
    GetExpParenNodeList(_env *LnsEnv) *LnsList
    GetExpRefItemNodeList(_env *LnsEnv) *LnsList
    GetExpRefNodeList(_env *LnsEnv) *LnsList
    GetExpSetItemNodeList(_env *LnsEnv) *LnsList
    GetExpSetValNodeList(_env *LnsEnv) *LnsList
    GetExpSubDDDNodeList(_env *LnsEnv) *LnsList
    GetExpToDDDNodeList(_env *LnsEnv) *LnsList
    GetExpUnwrapNodeList(_env *LnsEnv) *LnsList
    GetForNodeList(_env *LnsEnv) *LnsList
    GetForeachNodeList(_env *LnsEnv) *LnsList
    GetForsortNodeList(_env *LnsEnv) *LnsList
    GetGetFieldNodeList(_env *LnsEnv) *LnsList
    GetIfNodeList(_env *LnsEnv) *LnsList
    GetIfUnwrapNodeList(_env *LnsEnv) *LnsList
    GetImportNodeList(_env *LnsEnv) *LnsList
    GetList(_env *LnsEnv, arg1 LnsInt) *LnsList
    GetLiteralArrayNodeList(_env *LnsEnv) *LnsList
    GetLiteralBoolNodeList(_env *LnsEnv) *LnsList
    GetLiteralCharNodeList(_env *LnsEnv) *LnsList
    GetLiteralIntNodeList(_env *LnsEnv) *LnsList
    GetLiteralListNodeList(_env *LnsEnv) *LnsList
    GetLiteralMapNodeList(_env *LnsEnv) *LnsList
    GetLiteralNilNodeList(_env *LnsEnv) *LnsList
    GetLiteralRealNodeList(_env *LnsEnv) *LnsList
    GetLiteralSetNodeList(_env *LnsEnv) *LnsList
    GetLiteralStringNodeList(_env *LnsEnv) *LnsList
    GetLiteralSymbolNodeList(_env *LnsEnv) *LnsList
    GetLuneControlNodeList(_env *LnsEnv) *LnsList
    GetLuneKindNodeList(_env *LnsEnv) *LnsList
    GetMatchNodeList(_env *LnsEnv) *LnsList
    GetNewAlgeValNodeList(_env *LnsEnv) *LnsList
    GetNoneNodeList(_env *LnsEnv) *LnsList
    GetProtoClassNodeList(_env *LnsEnv) *LnsList
    GetProtoMethodNodeList(_env *LnsEnv) *LnsList
    GetProvideNodeList(_env *LnsEnv) *LnsList
    GetRefFieldNodeList(_env *LnsEnv) *LnsList
    GetRefTypeNodeList(_env *LnsEnv) *LnsList
    GetRepeatNodeList(_env *LnsEnv) *LnsList
    GetRequestNodeList(_env *LnsEnv) *LnsList
    GetReturnNodeList(_env *LnsEnv) *LnsList
    GetRootNodeList(_env *LnsEnv) *LnsList
    GetScopeNodeList(_env *LnsEnv) *LnsList
    GetShebangNodeList(_env *LnsEnv) *LnsList
    GetStmtExpNodeList(_env *LnsEnv) *LnsList
    GetSubfileNodeList(_env *LnsEnv) *LnsList
    GetSwitchNodeList(_env *LnsEnv) *LnsList
    GetTestBlockNodeList(_env *LnsEnv) *LnsList
    GetTestCaseNodeList(_env *LnsEnv) *LnsList
    GetUnboxingNodeList(_env *LnsEnv) *LnsList
    GetUnwrapSetNodeList(_env *LnsEnv) *LnsList
    GetWhenNodeList(_env *LnsEnv) *LnsList
    GetWhileNodeList(_env *LnsEnv) *LnsList
    Get_managerId(_env *LnsEnv) LnsInt
    NextId(_env *LnsEnv) LnsInt
    Set_managerId(_env *LnsEnv, arg1 LnsInt)
}
type Nodes_NodeManager struct {
    nodeKind2NodeList *LnsMap
    idSeed LnsInt
    managerId LnsInt
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
func NewNodes_NodeManager(_env *LnsEnv) *Nodes_NodeManager {
    obj := &Nodes_NodeManager{}
    obj.FP = obj
    obj.InitNodes_NodeManager(_env)
    return obj
}
func (self *Nodes_NodeManager) Get_managerId(_env *LnsEnv) LnsInt{ return self.managerId }
func (self *Nodes_NodeManager) Set_managerId(_env *LnsEnv, arg1 LnsInt){ self.managerId = arg1 }
// 469: DeclConstr
func (self *Nodes_NodeManager) InitNodes_NodeManager(_env *LnsEnv) {
    self.managerId = 0
    self.idSeed = 0
    self.nodeKind2NodeList = NewLnsMap( map[LnsAny]LnsAny{})
    for _kind, _ := range( Nodes_nodeKind2NameMap.Items ) {
        kind := _kind.(LnsInt)
        if Lns_op_not(self.nodeKind2NodeList.Get(kind)){
            self.nodeKind2NodeList.Set(kind,NewLnsList([]LnsAny{}))
        }
    }
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
func NewNodes_NodeKind(_env *LnsEnv) *Nodes_NodeKind {
    obj := &Nodes_NodeKind{}
    obj.FP = obj
    obj.InitNodes_NodeKind(_env)
    return obj
}
func (self *Nodes_NodeKind) InitNodes_NodeKind(_env *LnsEnv) {
}

// declaration Class -- NoneNode
type Nodes_NoneNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_NoneNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList) *Nodes_NoneNode {
    obj := &Nodes_NoneNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_NoneNode(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
// 734: DeclConstr
func (self *Nodes_NoneNode) InitNodes_NoneNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 0, pos, inTestBlock, macroArgFlag, typeList)
}


// declaration Class -- ShebangNode
type Nodes_ShebangNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_cmd(_env *LnsEnv) string
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ShebangNode struct {
    Nodes_Node
    cmd string
    FP Nodes_ShebangNodeMtd
}
func Nodes_ShebangNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ShebangNode).FP
}
type Nodes_ShebangNodeDownCast interface {
    ToNodes_ShebangNode() *Nodes_ShebangNode
}
func Nodes_ShebangNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ShebangNodeDownCast)
    if ok { return work.ToNodes_ShebangNode() }
    return nil
}
func (obj *Nodes_ShebangNode) ToNodes_ShebangNode() *Nodes_ShebangNode {
    return obj
}
func NewNodes_ShebangNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 string) *Nodes_ShebangNode {
    obj := &Nodes_ShebangNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ShebangNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ShebangNode) Get_cmd(_env *LnsEnv) string{ return self.cmd }
// 734: DeclConstr
func (self *Nodes_ShebangNode) InitNodes_ShebangNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,cmd string) {
    self.InitNodes_Node(_env, managerId, id, 1, pos, inTestBlock, macroArgFlag, typeList)
    self.cmd = cmd
}


// declaration Class -- ConvStatNode
type Nodes_ConvStatNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_txt(_env *LnsEnv) string
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ConvStatNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 string) *Nodes_ConvStatNode {
    obj := &Nodes_ConvStatNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ConvStatNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ConvStatNode) Get_txt(_env *LnsEnv) string{ return self.txt }
// 734: DeclConstr
func (self *Nodes_ConvStatNode) InitNodes_ConvStatNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,txt string) {
    self.InitNodes_Node(_env, managerId, id, 2, pos, inTestBlock, macroArgFlag, typeList)
    self.txt = txt
}


// declaration Class -- BlankLineNode
type Nodes_BlankLineNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_lineNum(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_BlankLineNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt) *Nodes_BlankLineNode {
    obj := &Nodes_BlankLineNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_BlankLineNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_BlankLineNode) Get_lineNum(_env *LnsEnv) LnsInt{ return self.lineNum }
// 734: DeclConstr
func (self *Nodes_BlankLineNode) InitNodes_BlankLineNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,lineNum LnsInt) {
    self.InitNodes_Node(_env, managerId, id, 3, pos, inTestBlock, macroArgFlag, typeList)
    self.lineNum = lineNum
}


// declaration Class -- SubfileNode
type Nodes_SubfileNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_usePath(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_SubfileNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsAny) *Nodes_SubfileNode {
    obj := &Nodes_SubfileNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_SubfileNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_SubfileNode) Get_usePath(_env *LnsEnv) LnsAny{ return self.usePath }
// 734: DeclConstr
func (self *Nodes_SubfileNode) InitNodes_SubfileNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,usePath LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 4, pos, inTestBlock, macroArgFlag, typeList)
    self.usePath = usePath
}


// declaration Class -- ImportInfo
type Nodes_ImportInfoMtd interface {
    Get_assignName(_env *LnsEnv) string
    Get_assigned(_env *LnsEnv) bool
    Get_lazy(_env *LnsEnv) LnsInt
    Get_modulePath(_env *LnsEnv) string
    Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_pos(_env *LnsEnv) Types_Position
    Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo
}
type Nodes_ImportInfo struct {
    pos Types_Position
    modulePath string
    lazy LnsInt
    assignName string
    assigned bool
    symbolInfo *Ast_SymbolInfo
    moduleTypeInfo *Ast_TypeInfo
    FP Nodes_ImportInfoMtd
}
func Nodes_ImportInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ImportInfo).FP
}
type Nodes_ImportInfoDownCast interface {
    ToNodes_ImportInfo() *Nodes_ImportInfo
}
func Nodes_ImportInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ImportInfoDownCast)
    if ok { return work.ToNodes_ImportInfo() }
    return nil
}
func (obj *Nodes_ImportInfo) ToNodes_ImportInfo() *Nodes_ImportInfo {
    return obj
}
func NewNodes_ImportInfo(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 LnsInt, arg4 string, arg5 bool, arg6 *Ast_SymbolInfo, arg7 *Ast_TypeInfo) *Nodes_ImportInfo {
    obj := &Nodes_ImportInfo{}
    obj.FP = obj
    obj.InitNodes_ImportInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ImportInfo) InitNodes_ImportInfo(_env *LnsEnv, arg1 Types_Position, arg2 string, arg3 LnsInt, arg4 string, arg5 bool, arg6 *Ast_SymbolInfo, arg7 *Ast_TypeInfo) {
    self.pos = arg1
    self.modulePath = arg2
    self.lazy = arg3
    self.assignName = arg4
    self.assigned = arg5
    self.symbolInfo = arg6
    self.moduleTypeInfo = arg7
}
func (self *Nodes_ImportInfo) Get_pos(_env *LnsEnv) Types_Position{ return self.pos }
func (self *Nodes_ImportInfo) Get_modulePath(_env *LnsEnv) string{ return self.modulePath }
func (self *Nodes_ImportInfo) Get_lazy(_env *LnsEnv) LnsInt{ return self.lazy }
func (self *Nodes_ImportInfo) Get_assignName(_env *LnsEnv) string{ return self.assignName }
func (self *Nodes_ImportInfo) Get_assigned(_env *LnsEnv) bool{ return self.assigned }
func (self *Nodes_ImportInfo) Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbolInfo }
func (self *Nodes_ImportInfo) Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.moduleTypeInfo }

// declaration Class -- ImportNode
type Nodes_ImportNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_info(_env *LnsEnv) *Nodes_ImportInfo
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ImportNode struct {
    Nodes_Node
    info *Nodes_ImportInfo
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
func NewNodes_ImportNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_ImportInfo) *Nodes_ImportNode {
    obj := &Nodes_ImportNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ImportNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ImportNode) Get_info(_env *LnsEnv) *Nodes_ImportInfo{ return self.info }
// 734: DeclConstr
func (self *Nodes_ImportNode) InitNodes_ImportNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,info *Nodes_ImportInfo) {
    self.InitNodes_Node(_env, managerId, id, 5, pos, inTestBlock, macroArgFlag, typeList)
    self.info = info
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
func NewNodes_MacroValInfo(_env *LnsEnv, arg1 LnsAny, arg2 *Ast_TypeInfo, arg3 LnsAny) *Nodes_MacroValInfo {
    obj := &Nodes_MacroValInfo{}
    obj.FP = obj
    obj.InitNodes_MacroValInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Nodes_MacroValInfo) InitNodes_MacroValInfo(_env *LnsEnv, arg1 LnsAny, arg2 *Ast_TypeInfo, arg3 LnsAny) {
    self.Val = arg1
    self.TypeInfo = arg2
    self.ArgNode = arg3
}

// declaration Class -- MacroArgInfo
type Nodes_MacroArgInfoMtd interface {
    Get_name(_env *LnsEnv) string
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
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
func NewNodes_MacroArgInfo(_env *LnsEnv, arg1 string, arg2 *Ast_TypeInfo) *Nodes_MacroArgInfo {
    obj := &Nodes_MacroArgInfo{}
    obj.FP = obj
    obj.InitNodes_MacroArgInfo(_env, arg1, arg2)
    return obj
}
func (self *Nodes_MacroArgInfo) InitNodes_MacroArgInfo(_env *LnsEnv, arg1 string, arg2 *Ast_TypeInfo) {
    self.name = arg1
    self.typeInfo = arg2
}
func (self *Nodes_MacroArgInfo) Get_name(_env *LnsEnv) string{ return self.name }
func (self *Nodes_MacroArgInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }

// declaration Class -- MacroInfo
type Nodes_MacroInfoMtd interface {
    GetArgList(_env *LnsEnv) *LnsList
    GetTokenList(_env *LnsEnv) *LnsList
    Get_func(_env *LnsEnv) LnsAny
    Get_name(_env *LnsEnv) string
    Get_symbol2MacroValInfoMap(_env *LnsEnv) *LnsMap
    Set_func(_env *LnsEnv, arg1 LnsAny)
}
type Nodes_MacroInfo struct {
    _func LnsAny
    symbol2MacroValInfoMap *LnsMap
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
func (self *Nodes_MacroInfo) InitNodes_MacroInfo(_env *LnsEnv, arg1 LnsAny, arg2 *LnsMap) {
    self._func = arg1
    self.symbol2MacroValInfoMap = arg2
}
func (self *Nodes_MacroInfo) Get_func(_env *LnsEnv) LnsAny{ return self._func }
func (self *Nodes_MacroInfo) Set_func(_env *LnsEnv, arg1 LnsAny){ self._func = arg1 }
func (self *Nodes_MacroInfo) Get_symbol2MacroValInfoMap(_env *LnsEnv) *LnsMap{ return self.symbol2MacroValInfoMap }




// declaration Class -- RootNode
type Nodes_RootNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_children(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_globalScope(_env *LnsEnv) *Ast_Scope
    Get_importModule2moduleInfo(_env *LnsEnv) *LnsMap
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_luneHelperInfo(_env *LnsEnv) *FrontInterface_LuneHelperInfo
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId
    Get_moduleScope(_env *LnsEnv) *Ast_Scope
    Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_nodeManager(_env *LnsEnv) *Nodes_NodeManager
    Get_pos(_env *LnsEnv) Types_Position
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_provideNode(_env *LnsEnv) LnsAny
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_typeId2ClassMap(_env *LnsEnv) *LnsMap
    Get_typeId2MacroInfo(_env *LnsEnv) *LnsMap
    Get_useModuleMacroSet(_env *LnsEnv) *LnsSet
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_provide(_env *LnsEnv, arg1 *Nodes_ProvideNode)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_RootNode struct {
    Nodes_Node
    children *LnsList
    moduleScope *Ast_Scope
    globalScope *Ast_Scope
    useModuleMacroSet *LnsSet
    moduleId *FrontInterface_ModuleId
    processInfo *Ast_ProcessInfo
    moduleTypeInfo *Ast_TypeInfo
    provideNode LnsAny
    luneHelperInfo *FrontInterface_LuneHelperInfo
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
func NewNodes_RootNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *LnsList, arg8 *Ast_Scope, arg9 *Ast_Scope, arg10 *LnsSet, arg11 *FrontInterface_ModuleId, arg12 *Ast_ProcessInfo, arg13 *Ast_TypeInfo, arg14 LnsAny, arg15 *FrontInterface_LuneHelperInfo, arg16 *Nodes_NodeManager, arg17 *LnsMap, arg18 *LnsMap, arg19 *LnsMap) *Nodes_RootNode {
    obj := &Nodes_RootNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_RootNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19)
    return obj
}
func (self *Nodes_RootNode) Get_children(_env *LnsEnv) *LnsList{ return self.children }
func (self *Nodes_RootNode) Get_moduleScope(_env *LnsEnv) *Ast_Scope{ return self.moduleScope }
func (self *Nodes_RootNode) Get_globalScope(_env *LnsEnv) *Ast_Scope{ return self.globalScope }
func (self *Nodes_RootNode) Get_useModuleMacroSet(_env *LnsEnv) *LnsSet{ return self.useModuleMacroSet }
func (self *Nodes_RootNode) Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId{ return self.moduleId }
func (self *Nodes_RootNode) Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo{ return self.processInfo }
func (self *Nodes_RootNode) Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.moduleTypeInfo }
func (self *Nodes_RootNode) Get_provideNode(_env *LnsEnv) LnsAny{ return self.provideNode }
func (self *Nodes_RootNode) Get_luneHelperInfo(_env *LnsEnv) *FrontInterface_LuneHelperInfo{ return self.luneHelperInfo }
func (self *Nodes_RootNode) Get_nodeManager(_env *LnsEnv) *Nodes_NodeManager{ return self.nodeManager }
func (self *Nodes_RootNode) Get_importModule2moduleInfo(_env *LnsEnv) *LnsMap{ return self.importModule2moduleInfo }
func (self *Nodes_RootNode) Get_typeId2MacroInfo(_env *LnsEnv) *LnsMap{ return self.typeId2MacroInfo }
func (self *Nodes_RootNode) Get_typeId2ClassMap(_env *LnsEnv) *LnsMap{ return self.typeId2ClassMap }
// 734: DeclConstr
func (self *Nodes_RootNode) InitNodes_RootNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,children *LnsList,moduleScope *Ast_Scope,globalScope *Ast_Scope,useModuleMacroSet *LnsSet,moduleId *FrontInterface_ModuleId,processInfo *Ast_ProcessInfo,moduleTypeInfo *Ast_TypeInfo,provideNode LnsAny,luneHelperInfo *FrontInterface_LuneHelperInfo,nodeManager *Nodes_NodeManager,importModule2moduleInfo *LnsMap,typeId2MacroInfo *LnsMap,typeId2ClassMap *LnsMap) {
    self.InitNodes_Node(_env, managerId, id, 6, pos, inTestBlock, macroArgFlag, typeList)
    self.children = children
    self.moduleScope = moduleScope
    self.globalScope = globalScope
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


// declaration Class -- RefTypeNode
type Nodes_RefTypeNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_array(_env *LnsEnv) string
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_itemNodeList(_env *LnsEnv) *LnsList
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_mutMode(_env *LnsEnv) LnsAny
    Get_name(_env *LnsEnv) *Nodes_Node
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_RefTypeNode struct {
    Nodes_Node
    name *Nodes_Node
    itemNodeList *LnsList
    mutMode LnsAny
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
func NewNodes_RefTypeNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 *LnsList, arg9 LnsAny, arg10 string) *Nodes_RefTypeNode {
    obj := &Nodes_RefTypeNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_RefTypeNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_RefTypeNode) Get_name(_env *LnsEnv) *Nodes_Node{ return self.name }
func (self *Nodes_RefTypeNode) Get_itemNodeList(_env *LnsEnv) *LnsList{ return self.itemNodeList }
func (self *Nodes_RefTypeNode) Get_mutMode(_env *LnsEnv) LnsAny{ return self.mutMode }
func (self *Nodes_RefTypeNode) Get_array(_env *LnsEnv) string{ return self.array }
// 734: DeclConstr
func (self *Nodes_RefTypeNode) InitNodes_RefTypeNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Nodes_Node,itemNodeList *LnsList,mutMode LnsAny,array string) {
    self.InitNodes_Node(_env, managerId, id, 7, pos, inTestBlock, macroArgFlag, typeList)
    self.name = name
    self.itemNodeList = itemNodeList
    self.mutMode = mutMode
    self.array = array
}


// declaration Class -- BlockNode
type Nodes_BlockNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_blockKind(_env *LnsEnv) LnsInt
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_stmtList(_env *LnsEnv) *LnsList
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_BlockNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Ast_Scope, arg9 *LnsList) *Nodes_BlockNode {
    obj := &Nodes_BlockNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_BlockNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_BlockNode) Get_blockKind(_env *LnsEnv) LnsInt{ return self.blockKind }
func (self *Nodes_BlockNode) Get_scope(_env *LnsEnv) *Ast_Scope{ return self.scope }
func (self *Nodes_BlockNode) Get_stmtList(_env *LnsEnv) *LnsList{ return self.stmtList }
// 734: DeclConstr
func (self *Nodes_BlockNode) InitNodes_BlockNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,blockKind LnsInt,scope *Ast_Scope,stmtList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 8, pos, inTestBlock, macroArgFlag, typeList)
    self.blockKind = blockKind
    self.scope = scope
    self.stmtList = stmtList
}


// declaration Class -- ScopeNode
type Nodes_ScopeNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_scopeKind(_env *LnsEnv) LnsInt
    Get_symbolList(_env *LnsEnv) *LnsList
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ScopeNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Ast_Scope, arg9 *LnsList, arg10 *Nodes_BlockNode) *Nodes_ScopeNode {
    obj := &Nodes_ScopeNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ScopeNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_ScopeNode) Get_scopeKind(_env *LnsEnv) LnsInt{ return self.scopeKind }
func (self *Nodes_ScopeNode) Get_scope(_env *LnsEnv) *Ast_Scope{ return self.scope }
func (self *Nodes_ScopeNode) Get_symbolList(_env *LnsEnv) *LnsList{ return self.symbolList }
func (self *Nodes_ScopeNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
// 734: DeclConstr
func (self *Nodes_ScopeNode) InitNodes_ScopeNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,scopeKind LnsInt,scope *Ast_Scope,symbolList *LnsList,block *Nodes_BlockNode) {
    self.InitNodes_Node(_env, managerId, id, 9, pos, inTestBlock, macroArgFlag, typeList)
    self.scopeKind = scopeKind
    self.scope = scope
    self.symbolList = symbolList
    self.block = block
}


// declaration Class -- IfStmtInfo
type Nodes_IfStmtInfoMtd interface {
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_kind(_env *LnsEnv) LnsInt
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
func NewNodes_IfStmtInfo(_env *LnsEnv, arg1 LnsInt, arg2 *Nodes_Node, arg3 *Nodes_BlockNode) *Nodes_IfStmtInfo {
    obj := &Nodes_IfStmtInfo{}
    obj.FP = obj
    obj.InitNodes_IfStmtInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Nodes_IfStmtInfo) InitNodes_IfStmtInfo(_env *LnsEnv, arg1 LnsInt, arg2 *Nodes_Node, arg3 *Nodes_BlockNode) {
    self.kind = arg1
    self.exp = arg2
    self.block = arg3
}
func (self *Nodes_IfStmtInfo) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *Nodes_IfStmtInfo) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
func (self *Nodes_IfStmtInfo) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }

// declaration Class -- IfNode
type Nodes_IfNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_stmtList(_env *LnsEnv) *LnsList
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_IfNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *LnsList) *Nodes_IfNode {
    obj := &Nodes_IfNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_IfNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_IfNode) Get_stmtList(_env *LnsEnv) *LnsList{ return self.stmtList }
// 734: DeclConstr
func (self *Nodes_IfNode) InitNodes_IfNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 10, pos, inTestBlock, macroArgFlag, typeList)
    self.stmtList = stmtList
}


// declaration Class -- MRetExp
type Nodes_MRetExpMtd interface {
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_index(_env *LnsEnv) LnsInt
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
func NewNodes_MRetExp(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsInt) *Nodes_MRetExp {
    obj := &Nodes_MRetExp{}
    obj.FP = obj
    obj.InitNodes_MRetExp(_env, arg1, arg2)
    return obj
}
func (self *Nodes_MRetExp) InitNodes_MRetExp(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsInt) {
    self.exp = arg1
    self.index = arg2
}
func (self *Nodes_MRetExp) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
func (self *Nodes_MRetExp) Get_index(_env *LnsEnv) LnsInt{ return self.index }

// declaration Class -- ExpListNode
type Nodes_ExpListNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetExpTypeAt(_env *LnsEnv, arg1 LnsInt) *Ast_TypeInfo
    GetExpTypeNoDDDAt(_env *LnsEnv, arg1 LnsInt) *Ast_TypeInfo
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) *LnsList
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_followOn(_env *LnsEnv) bool
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_mRetExp(_env *LnsEnv) LnsAny
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpListNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *LnsList, arg8 LnsAny, arg9 bool) *Nodes_ExpListNode {
    obj := &Nodes_ExpListNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpListNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ExpListNode) Get_expList(_env *LnsEnv) *LnsList{ return self.expList }
func (self *Nodes_ExpListNode) Get_mRetExp(_env *LnsEnv) LnsAny{ return self.mRetExp }
func (self *Nodes_ExpListNode) Get_followOn(_env *LnsEnv) bool{ return self.followOn }
// 734: DeclConstr
func (self *Nodes_ExpListNode) InitNodes_ExpListNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList *LnsList,mRetExp LnsAny,followOn bool) {
    self.InitNodes_Node(_env, managerId, id, 11, pos, inTestBlock, macroArgFlag, typeList)
    self.expList = expList
    self.mRetExp = mRetExp
    self.followOn = followOn
}


// declaration Class -- CaseInfo
type Nodes_CaseInfoMtd interface {
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_expList(_env *LnsEnv) *Nodes_ExpListNode
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
func NewNodes_CaseInfo(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 *Nodes_BlockNode) *Nodes_CaseInfo {
    obj := &Nodes_CaseInfo{}
    obj.FP = obj
    obj.InitNodes_CaseInfo(_env, arg1, arg2)
    return obj
}
func (self *Nodes_CaseInfo) InitNodes_CaseInfo(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 *Nodes_BlockNode) {
    self.expList = arg1
    self.block = arg2
}
func (self *Nodes_CaseInfo) Get_expList(_env *LnsEnv) *Nodes_ExpListNode{ return self.expList }
func (self *Nodes_CaseInfo) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }

// declaration Class -- SwitchNode
type Nodes_SwitchNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_caseKind(_env *LnsEnv) LnsInt
    Get_caseList(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_default(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_failSafeDefault(_env *LnsEnv) bool
    Get_idInNS(_env *LnsEnv) LnsInt
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_SwitchNode struct {
    Nodes_Node
    idInNS LnsInt
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
func NewNodes_SwitchNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Nodes_Node, arg9 *LnsList, arg10 LnsAny, arg11 LnsInt, arg12 bool) *Nodes_SwitchNode {
    obj := &Nodes_SwitchNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_SwitchNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
    return obj
}
func (self *Nodes_SwitchNode) Get_idInNS(_env *LnsEnv) LnsInt{ return self.idInNS }
func (self *Nodes_SwitchNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
func (self *Nodes_SwitchNode) Get_caseList(_env *LnsEnv) *LnsList{ return self.caseList }
func (self *Nodes_SwitchNode) Get_default(_env *LnsEnv) LnsAny{ return self._default }
func (self *Nodes_SwitchNode) Get_caseKind(_env *LnsEnv) LnsInt{ return self.caseKind }
func (self *Nodes_SwitchNode) Get_failSafeDefault(_env *LnsEnv) bool{ return self.failSafeDefault }
// 734: DeclConstr
func (self *Nodes_SwitchNode) InitNodes_SwitchNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,exp *Nodes_Node,caseList *LnsList,_default LnsAny,caseKind LnsInt,failSafeDefault bool) {
    self.InitNodes_Node(_env, managerId, id, 12, pos, inTestBlock, macroArgFlag, typeList)
    self.idInNS = idInNS
    self.exp = exp
    self.caseList = caseList
    self._default = _default
    self.caseKind = caseKind
    self.failSafeDefault = failSafeDefault
}


// declaration Class -- WhileNode
type Nodes_WhileNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_infinit(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_WhileNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 bool, arg9 *Nodes_BlockNode) *Nodes_WhileNode {
    obj := &Nodes_WhileNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_WhileNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_WhileNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
func (self *Nodes_WhileNode) Get_infinit(_env *LnsEnv) bool{ return self.infinit }
func (self *Nodes_WhileNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
// 734: DeclConstr
func (self *Nodes_WhileNode) InitNodes_WhileNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,infinit bool,block *Nodes_BlockNode) {
    self.InitNodes_Node(_env, managerId, id, 13, pos, inTestBlock, macroArgFlag, typeList)
    self.exp = exp
    self.infinit = infinit
    self.block = block
}


// declaration Class -- RepeatNode
type Nodes_RepeatNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_RepeatNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_BlockNode, arg8 *Nodes_Node) *Nodes_RepeatNode {
    obj := &Nodes_RepeatNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_RepeatNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_RepeatNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
func (self *Nodes_RepeatNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
// 734: DeclConstr
func (self *Nodes_RepeatNode) InitNodes_RepeatNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,block *Nodes_BlockNode,exp *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 14, pos, inTestBlock, macroArgFlag, typeList)
    self.block = block
    self.exp = exp
}


// declaration Class -- ForNode
type Nodes_ForNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_delta(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_idInNS(_env *LnsEnv) LnsInt
    Get_inTestBlock(_env *LnsEnv) bool
    Get_init(_env *LnsEnv) *Nodes_Node
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_to(_env *LnsEnv) *Nodes_Node
    Get_val(_env *LnsEnv) *Ast_SymbolInfo
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ForNode struct {
    Nodes_Node
    idInNS LnsInt
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
func NewNodes_ForNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Nodes_BlockNode, arg9 *Ast_SymbolInfo, arg10 *Nodes_Node, arg11 *Nodes_Node, arg12 LnsAny) *Nodes_ForNode {
    obj := &Nodes_ForNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ForNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
    return obj
}
func (self *Nodes_ForNode) Get_idInNS(_env *LnsEnv) LnsInt{ return self.idInNS }
func (self *Nodes_ForNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
func (self *Nodes_ForNode) Get_val(_env *LnsEnv) *Ast_SymbolInfo{ return self.val }
func (self *Nodes_ForNode) Get_init(_env *LnsEnv) *Nodes_Node{ return self.init }
func (self *Nodes_ForNode) Get_to(_env *LnsEnv) *Nodes_Node{ return self.to }
func (self *Nodes_ForNode) Get_delta(_env *LnsEnv) LnsAny{ return self.delta }
// 734: DeclConstr
func (self *Nodes_ForNode) InitNodes_ForNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,block *Nodes_BlockNode,val *Ast_SymbolInfo,init *Nodes_Node,to *Nodes_Node,delta LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 15, pos, inTestBlock, macroArgFlag, typeList)
    self.idInNS = idInNS
    self.block = block
    self.val = val
    self.init = init
    self.to = to
    self.delta = delta
}


// declaration Class -- ApplyNode
type Nodes_ApplyNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) *Nodes_ExpListNode
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_idInNS(_env *LnsEnv) LnsInt
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_varList(_env *LnsEnv) *LnsList
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ApplyNode struct {
    Nodes_Node
    idInNS LnsInt
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
func NewNodes_ApplyNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *LnsList, arg9 *Nodes_ExpListNode, arg10 *Nodes_BlockNode) *Nodes_ApplyNode {
    obj := &Nodes_ApplyNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ApplyNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_ApplyNode) Get_idInNS(_env *LnsEnv) LnsInt{ return self.idInNS }
func (self *Nodes_ApplyNode) Get_varList(_env *LnsEnv) *LnsList{ return self.varList }
func (self *Nodes_ApplyNode) Get_expList(_env *LnsEnv) *Nodes_ExpListNode{ return self.expList }
func (self *Nodes_ApplyNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
// 734: DeclConstr
func (self *Nodes_ApplyNode) InitNodes_ApplyNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,varList *LnsList,expList *Nodes_ExpListNode,block *Nodes_BlockNode) {
    self.InitNodes_Node(_env, managerId, id, 16, pos, inTestBlock, macroArgFlag, typeList)
    self.idInNS = idInNS
    self.varList = varList
    self.expList = expList
    self.block = block
}


// declaration Class -- ForeachNode
type Nodes_ForeachNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_idInNS(_env *LnsEnv) LnsInt
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_key(_env *LnsEnv) LnsAny
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_val(_env *LnsEnv) *Ast_SymbolInfo
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ForeachNode struct {
    Nodes_Node
    idInNS LnsInt
    val *Ast_SymbolInfo
    key LnsAny
    exp *Nodes_Node
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
func NewNodes_ForeachNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Ast_SymbolInfo, arg9 LnsAny, arg10 *Nodes_Node, arg11 *Nodes_BlockNode) *Nodes_ForeachNode {
    obj := &Nodes_ForeachNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ForeachNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
    return obj
}
func (self *Nodes_ForeachNode) Get_idInNS(_env *LnsEnv) LnsInt{ return self.idInNS }
func (self *Nodes_ForeachNode) Get_val(_env *LnsEnv) *Ast_SymbolInfo{ return self.val }
func (self *Nodes_ForeachNode) Get_key(_env *LnsEnv) LnsAny{ return self.key }
func (self *Nodes_ForeachNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
func (self *Nodes_ForeachNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
// 734: DeclConstr
func (self *Nodes_ForeachNode) InitNodes_ForeachNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,val *Ast_SymbolInfo,key LnsAny,exp *Nodes_Node,block *Nodes_BlockNode) {
    self.InitNodes_Node(_env, managerId, id, 17, pos, inTestBlock, macroArgFlag, typeList)
    self.idInNS = idInNS
    self.val = val
    self.key = key
    self.exp = exp
    self.block = block
}


// declaration Class -- ForsortNode
type Nodes_ForsortNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_idInNS(_env *LnsEnv) LnsInt
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_key(_env *LnsEnv) LnsAny
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_sort(_env *LnsEnv) bool
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_val(_env *LnsEnv) *Ast_SymbolInfo
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ForsortNode struct {
    Nodes_Node
    idInNS LnsInt
    val *Ast_SymbolInfo
    key LnsAny
    exp *Nodes_Node
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
func NewNodes_ForsortNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Ast_SymbolInfo, arg9 LnsAny, arg10 *Nodes_Node, arg11 *Nodes_BlockNode, arg12 bool) *Nodes_ForsortNode {
    obj := &Nodes_ForsortNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ForsortNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
    return obj
}
func (self *Nodes_ForsortNode) Get_idInNS(_env *LnsEnv) LnsInt{ return self.idInNS }
func (self *Nodes_ForsortNode) Get_val(_env *LnsEnv) *Ast_SymbolInfo{ return self.val }
func (self *Nodes_ForsortNode) Get_key(_env *LnsEnv) LnsAny{ return self.key }
func (self *Nodes_ForsortNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
func (self *Nodes_ForsortNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
func (self *Nodes_ForsortNode) Get_sort(_env *LnsEnv) bool{ return self.sort }
// 734: DeclConstr
func (self *Nodes_ForsortNode) InitNodes_ForsortNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,val *Ast_SymbolInfo,key LnsAny,exp *Nodes_Node,block *Nodes_BlockNode,sort bool) {
    self.InitNodes_Node(_env, managerId, id, 18, pos, inTestBlock, macroArgFlag, typeList)
    self.idInNS = idInNS
    self.val = val
    self.key = key
    self.exp = exp
    self.block = block
    self.sort = sort
}


// declaration Class -- ReturnNode
type Nodes_ReturnNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) LnsAny
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ReturnNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsAny) *Nodes_ReturnNode {
    obj := &Nodes_ReturnNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ReturnNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ReturnNode) Get_expList(_env *LnsEnv) LnsAny{ return self.expList }
// 734: DeclConstr
func (self *Nodes_ReturnNode) InitNodes_ReturnNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 19, pos, inTestBlock, macroArgFlag, typeList)
    self.expList = expList
}


// declaration Class -- BreakNode
type Nodes_BreakNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_BreakNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList) *Nodes_BreakNode {
    obj := &Nodes_BreakNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_BreakNode(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
// 734: DeclConstr
func (self *Nodes_BreakNode) InitNodes_BreakNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 20, pos, inTestBlock, macroArgFlag, typeList)
}


// declaration Class -- ProvideNode
type Nodes_ProvideNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_symbol(_env *LnsEnv) *Ast_SymbolInfo
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ProvideNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Ast_SymbolInfo) *Nodes_ProvideNode {
    obj := &Nodes_ProvideNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ProvideNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ProvideNode) Get_symbol(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbol }
// 734: DeclConstr
func (self *Nodes_ProvideNode) InitNodes_ProvideNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,symbol *Ast_SymbolInfo) {
    self.InitNodes_Node(_env, managerId, id, 21, pos, inTestBlock, macroArgFlag, typeList)
    self.symbol = symbol
}


// declaration Class -- ExpNewNode
type Nodes_ExpNewNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_argList(_env *LnsEnv) LnsAny
    Get_commentList(_env *LnsEnv) LnsAny
    Get_ctorTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_symbol(_env *LnsEnv) *Nodes_Node
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpNewNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 *Ast_TypeInfo, arg9 LnsAny) *Nodes_ExpNewNode {
    obj := &Nodes_ExpNewNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpNewNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ExpNewNode) Get_symbol(_env *LnsEnv) *Nodes_Node{ return self.symbol }
func (self *Nodes_ExpNewNode) Get_ctorTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.ctorTypeInfo }
func (self *Nodes_ExpNewNode) Get_argList(_env *LnsEnv) LnsAny{ return self.argList }
// 734: DeclConstr
func (self *Nodes_ExpNewNode) InitNodes_ExpNewNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,symbol *Nodes_Node,ctorTypeInfo *Ast_TypeInfo,argList LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 22, pos, inTestBlock, macroArgFlag, typeList)
    self.symbol = symbol
    self.ctorTypeInfo = ctorTypeInfo
    self.argList = argList
}


// declaration Class -- ExpUnwrapNode
type Nodes_ExpUnwrapNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_default(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpUnwrapNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 LnsAny) *Nodes_ExpUnwrapNode {
    obj := &Nodes_ExpUnwrapNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpUnwrapNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_ExpUnwrapNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
func (self *Nodes_ExpUnwrapNode) Get_default(_env *LnsEnv) LnsAny{ return self._default }
// 734: DeclConstr
func (self *Nodes_ExpUnwrapNode) InitNodes_ExpUnwrapNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,_default LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 23, pos, inTestBlock, macroArgFlag, typeList)
    self.exp = exp
    self._default = _default
}


// declaration Class -- ExpRefNode
type Nodes_ExpRefNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpRefNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Ast_SymbolInfo) *Nodes_ExpRefNode {
    obj := &Nodes_ExpRefNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpRefNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpRefNode) Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbolInfo }
// 734: DeclConstr
func (self *Nodes_ExpRefNode) InitNodes_ExpRefNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,symbolInfo *Ast_SymbolInfo) {
    self.InitNodes_Node(_env, managerId, id, 24, pos, inTestBlock, macroArgFlag, typeList)
    self.symbolInfo = symbolInfo
}


// declaration Class -- ExpSetValNode
type Nodes_ExpSetValNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_LeftSymList(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp1(_env *LnsEnv) *Nodes_Node
    Get_exp2(_env *LnsEnv) *Nodes_ExpListNode
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_initSymSet(_env *LnsEnv) *LnsSet
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpSetValNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 *Nodes_ExpListNode, arg9 *LnsList, arg10 *LnsSet) *Nodes_ExpSetValNode {
    obj := &Nodes_ExpSetValNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpSetValNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_ExpSetValNode) Get_exp1(_env *LnsEnv) *Nodes_Node{ return self.exp1 }
func (self *Nodes_ExpSetValNode) Get_exp2(_env *LnsEnv) *Nodes_ExpListNode{ return self.exp2 }
func (self *Nodes_ExpSetValNode) Get_LeftSymList(_env *LnsEnv) *LnsList{ return self.LeftSymList }
func (self *Nodes_ExpSetValNode) Get_initSymSet(_env *LnsEnv) *LnsSet{ return self.initSymSet }
// 734: DeclConstr
func (self *Nodes_ExpSetValNode) InitNodes_ExpSetValNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp1 *Nodes_Node,exp2 *Nodes_ExpListNode,LeftSymList *LnsList,initSymSet *LnsSet) {
    self.InitNodes_Node(_env, managerId, id, 25, pos, inTestBlock, macroArgFlag, typeList)
    self.exp1 = exp1
    self.exp2 = exp2
    self.LeftSymList = LeftSymList
    self.initSymSet = initSymSet
}


// declaration Class -- ExpSetItemNode
type Nodes_ExpSetItemNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp2(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_index(_env *LnsEnv) LnsAny
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_val(_env *LnsEnv) *Nodes_Node
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpSetItemNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 LnsAny, arg9 *Nodes_Node) *Nodes_ExpSetItemNode {
    obj := &Nodes_ExpSetItemNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpSetItemNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ExpSetItemNode) Get_val(_env *LnsEnv) *Nodes_Node{ return self.val }
func (self *Nodes_ExpSetItemNode) Get_index(_env *LnsEnv) LnsAny{ return self.index }
func (self *Nodes_ExpSetItemNode) Get_exp2(_env *LnsEnv) *Nodes_Node{ return self.exp2 }
// 734: DeclConstr
func (self *Nodes_ExpSetItemNode) InitNodes_ExpSetItemNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,index LnsAny,exp2 *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 26, pos, inTestBlock, macroArgFlag, typeList)
    self.val = val
    self.index = index
    self.exp2 = exp2
}


// declaration Class -- ExpOp2Node
type Nodes_ExpOp2NodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    getValType(_env *LnsEnv, arg1 *Nodes_Node)(bool, LnsInt, LnsReal, string, *Ast_TypeInfo)
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp1(_env *LnsEnv) *Nodes_Node
    Get_exp2(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_op(_env *LnsEnv) *Types_Token
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ExpOp2Node struct {
    Nodes_Node
    op *Types_Token
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
func NewNodes_ExpOp2Node(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 *Nodes_Node, arg9 *Nodes_Node) *Nodes_ExpOp2Node {
    obj := &Nodes_ExpOp2Node{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpOp2Node(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ExpOp2Node) Get_op(_env *LnsEnv) *Types_Token{ return self.op }
func (self *Nodes_ExpOp2Node) Get_exp1(_env *LnsEnv) *Nodes_Node{ return self.exp1 }
func (self *Nodes_ExpOp2Node) Get_exp2(_env *LnsEnv) *Nodes_Node{ return self.exp2 }
// 734: DeclConstr
func (self *Nodes_ExpOp2Node) InitNodes_ExpOp2Node(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,op *Types_Token,exp1 *Nodes_Node,exp2 *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 27, pos, inTestBlock, macroArgFlag, typeList)
    self.op = op
    self.exp1 = exp1
    self.exp2 = exp2
}


// declaration Class -- UnwrapSetNode
type Nodes_UnwrapSetNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_dstExpList(_env *LnsEnv) *Nodes_ExpListNode
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_srcExpList(_env *LnsEnv) *Nodes_ExpListNode
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_unwrapBlock(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_UnwrapSetNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_ExpListNode, arg8 *Nodes_ExpListNode, arg9 LnsAny) *Nodes_UnwrapSetNode {
    obj := &Nodes_UnwrapSetNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_UnwrapSetNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_UnwrapSetNode) Get_dstExpList(_env *LnsEnv) *Nodes_ExpListNode{ return self.dstExpList }
func (self *Nodes_UnwrapSetNode) Get_srcExpList(_env *LnsEnv) *Nodes_ExpListNode{ return self.srcExpList }
func (self *Nodes_UnwrapSetNode) Get_unwrapBlock(_env *LnsEnv) LnsAny{ return self.unwrapBlock }
// 734: DeclConstr
func (self *Nodes_UnwrapSetNode) InitNodes_UnwrapSetNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,dstExpList *Nodes_ExpListNode,srcExpList *Nodes_ExpListNode,unwrapBlock LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 28, pos, inTestBlock, macroArgFlag, typeList)
    self.dstExpList = dstExpList
    self.srcExpList = srcExpList
    self.unwrapBlock = unwrapBlock
}


// declaration Class -- IfUnwrapNode
type Nodes_IfUnwrapNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) *Nodes_ExpListNode
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_nilBlock(_env *LnsEnv) LnsAny
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_varSymList(_env *LnsEnv) *LnsList
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_IfUnwrapNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *LnsList, arg8 *Nodes_ExpListNode, arg9 *Nodes_BlockNode, arg10 LnsAny) *Nodes_IfUnwrapNode {
    obj := &Nodes_IfUnwrapNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_IfUnwrapNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_IfUnwrapNode) Get_varSymList(_env *LnsEnv) *LnsList{ return self.varSymList }
func (self *Nodes_IfUnwrapNode) Get_expList(_env *LnsEnv) *Nodes_ExpListNode{ return self.expList }
func (self *Nodes_IfUnwrapNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
func (self *Nodes_IfUnwrapNode) Get_nilBlock(_env *LnsEnv) LnsAny{ return self.nilBlock }
// 734: DeclConstr
func (self *Nodes_IfUnwrapNode) InitNodes_IfUnwrapNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,varSymList *LnsList,expList *Nodes_ExpListNode,block *Nodes_BlockNode,nilBlock LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 29, pos, inTestBlock, macroArgFlag, typeList)
    self.varSymList = varSymList
    self.expList = expList
    self.block = block
    self.nilBlock = nilBlock
}


// declaration Class -- UnwrapSymbolPair
type Nodes_UnwrapSymbolPairMtd interface {
    Get_dst(_env *LnsEnv) *Ast_SymbolInfo
    Get_src(_env *LnsEnv) *Ast_SymbolInfo
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
func NewNodes_UnwrapSymbolPair(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 *Ast_SymbolInfo) *Nodes_UnwrapSymbolPair {
    obj := &Nodes_UnwrapSymbolPair{}
    obj.FP = obj
    obj.InitNodes_UnwrapSymbolPair(_env, arg1, arg2)
    return obj
}
func (self *Nodes_UnwrapSymbolPair) InitNodes_UnwrapSymbolPair(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 *Ast_SymbolInfo) {
    self.src = arg1
    self.dst = arg2
}
func (self *Nodes_UnwrapSymbolPair) Get_src(_env *LnsEnv) *Ast_SymbolInfo{ return self.src }
func (self *Nodes_UnwrapSymbolPair) Get_dst(_env *LnsEnv) *Ast_SymbolInfo{ return self.dst }

// declaration Class -- WhenNode
type Nodes_WhenNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_elseBlock(_env *LnsEnv) LnsAny
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_symPairList(_env *LnsEnv) *LnsList
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_WhenNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *LnsList, arg8 *Nodes_BlockNode, arg9 LnsAny) *Nodes_WhenNode {
    obj := &Nodes_WhenNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_WhenNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_WhenNode) Get_symPairList(_env *LnsEnv) *LnsList{ return self.symPairList }
func (self *Nodes_WhenNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
func (self *Nodes_WhenNode) Get_elseBlock(_env *LnsEnv) LnsAny{ return self.elseBlock }
// 734: DeclConstr
func (self *Nodes_WhenNode) InitNodes_WhenNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,symPairList *LnsList,block *Nodes_BlockNode,elseBlock LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 30, pos, inTestBlock, macroArgFlag, typeList)
    self.symPairList = symPairList
    self.block = block
    self.elseBlock = elseBlock
}


// declaration Class -- ExpCastNode
type Nodes_ExpCastNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_castKind(_env *LnsEnv) LnsInt
    Get_castOpe(_env *LnsEnv) string
    Get_castType(_env *LnsEnv) *Ast_TypeInfo
    Get_castTypeNode(_env *LnsEnv) LnsAny
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ExpCastNode struct {
    Nodes_Node
    exp *Nodes_Node
    castType *Ast_TypeInfo
    castTypeNode LnsAny
    castOpe string
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
func NewNodes_ExpCastNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 *Ast_TypeInfo, arg9 LnsAny, arg10 string, arg11 LnsInt) *Nodes_ExpCastNode {
    obj := &Nodes_ExpCastNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpCastNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
    return obj
}
func (self *Nodes_ExpCastNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
func (self *Nodes_ExpCastNode) Get_castType(_env *LnsEnv) *Ast_TypeInfo{ return self.castType }
func (self *Nodes_ExpCastNode) Get_castTypeNode(_env *LnsEnv) LnsAny{ return self.castTypeNode }
func (self *Nodes_ExpCastNode) Get_castOpe(_env *LnsEnv) string{ return self.castOpe }
func (self *Nodes_ExpCastNode) Get_castKind(_env *LnsEnv) LnsInt{ return self.castKind }
// 734: DeclConstr
func (self *Nodes_ExpCastNode) InitNodes_ExpCastNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,castType *Ast_TypeInfo,castTypeNode LnsAny,castOpe string,castKind LnsInt) {
    self.InitNodes_Node(_env, managerId, id, 31, pos, inTestBlock, macroArgFlag, typeList)
    self.exp = exp
    self.castType = castType
    self.castTypeNode = castTypeNode
    self.castOpe = castOpe
    self.castKind = castKind
}


// declaration Class -- ExpToDDDNode
type Nodes_ExpToDDDNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) *Nodes_ExpListNode
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpToDDDNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_ExpListNode) *Nodes_ExpToDDDNode {
    obj := &Nodes_ExpToDDDNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpToDDDNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpToDDDNode) Get_expList(_env *LnsEnv) *Nodes_ExpListNode{ return self.expList }
// 734: DeclConstr
func (self *Nodes_ExpToDDDNode) InitNodes_ExpToDDDNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList *Nodes_ExpListNode) {
    self.InitNodes_Node(_env, managerId, id, 32, pos, inTestBlock, macroArgFlag, typeList)
    self.expList = expList
}


// declaration Class -- ExpSubDDDNode
type Nodes_ExpSubDDDNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_remainIndex(_env *LnsEnv) LnsInt
    Get_src(_env *LnsEnv) *Nodes_Node
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpSubDDDNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 LnsInt) *Nodes_ExpSubDDDNode {
    obj := &Nodes_ExpSubDDDNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpSubDDDNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_ExpSubDDDNode) Get_src(_env *LnsEnv) *Nodes_Node{ return self.src }
func (self *Nodes_ExpSubDDDNode) Get_remainIndex(_env *LnsEnv) LnsInt{ return self.remainIndex }
// 734: DeclConstr
func (self *Nodes_ExpSubDDDNode) InitNodes_ExpSubDDDNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,src *Nodes_Node,remainIndex LnsInt) {
    self.InitNodes_Node(_env, managerId, id, 33, pos, inTestBlock, macroArgFlag, typeList)
    self.src = src
    self.remainIndex = remainIndex
}


// declaration Class -- ExpOp1Node
type Nodes_ExpOp1NodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_macroMode(_env *LnsEnv) LnsInt
    Get_op(_env *LnsEnv) *Types_Token
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpOp1Node(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 LnsInt, arg9 *Nodes_Node) *Nodes_ExpOp1Node {
    obj := &Nodes_ExpOp1Node{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpOp1Node(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ExpOp1Node) Get_op(_env *LnsEnv) *Types_Token{ return self.op }
func (self *Nodes_ExpOp1Node) Get_macroMode(_env *LnsEnv) LnsInt{ return self.macroMode }
func (self *Nodes_ExpOp1Node) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
// 734: DeclConstr
func (self *Nodes_ExpOp1Node) InitNodes_ExpOp1Node(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,op *Types_Token,macroMode LnsInt,exp *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 34, pos, inTestBlock, macroArgFlag, typeList)
    self.op = op
    self.macroMode = macroMode
    self.exp = exp
}


// declaration Class -- ExpRefItemNode
type Nodes_ExpRefItemNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_index(_env *LnsEnv) LnsAny
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_nilAccess(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_symbol(_env *LnsEnv) LnsAny
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_val(_env *LnsEnv) *Nodes_Node
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ExpRefItemNode struct {
    Nodes_Node
    val *Nodes_Node
    nilAccess bool
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
func NewNodes_ExpRefItemNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 bool, arg9 LnsAny, arg10 LnsAny) *Nodes_ExpRefItemNode {
    obj := &Nodes_ExpRefItemNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpRefItemNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_ExpRefItemNode) Get_val(_env *LnsEnv) *Nodes_Node{ return self.val }
func (self *Nodes_ExpRefItemNode) Get_nilAccess(_env *LnsEnv) bool{ return self.nilAccess }
func (self *Nodes_ExpRefItemNode) Get_symbol(_env *LnsEnv) LnsAny{ return self.symbol }
func (self *Nodes_ExpRefItemNode) Get_index(_env *LnsEnv) LnsAny{ return self.index }
// 734: DeclConstr
func (self *Nodes_ExpRefItemNode) InitNodes_ExpRefItemNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,nilAccess bool,symbol LnsAny,index LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 35, pos, inTestBlock, macroArgFlag, typeList)
    self.val = val
    self.nilAccess = nilAccess
    self.symbol = symbol
    self.index = index
}


// declaration Class -- ExpCallNode
type Nodes_ExpCallNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_argList(_env *LnsEnv) LnsAny
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_errorFunc(_env *LnsEnv) bool
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_func(_env *LnsEnv) *Nodes_Node
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_nilAccess(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ExpCallNode struct {
    Nodes_Node
    _func *Nodes_Node
    errorFunc bool
    nilAccess bool
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
func NewNodes_ExpCallNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 bool, arg9 bool, arg10 LnsAny) *Nodes_ExpCallNode {
    obj := &Nodes_ExpCallNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpCallNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_ExpCallNode) Get_func(_env *LnsEnv) *Nodes_Node{ return self._func }
func (self *Nodes_ExpCallNode) Get_errorFunc(_env *LnsEnv) bool{ return self.errorFunc }
func (self *Nodes_ExpCallNode) Get_nilAccess(_env *LnsEnv) bool{ return self.nilAccess }
func (self *Nodes_ExpCallNode) Get_argList(_env *LnsEnv) LnsAny{ return self.argList }
// 734: DeclConstr
func (self *Nodes_ExpCallNode) InitNodes_ExpCallNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,_func *Nodes_Node,errorFunc bool,nilAccess bool,argList LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 36, pos, inTestBlock, macroArgFlag, typeList)
    self._func = _func
    self.errorFunc = errorFunc
    self.nilAccess = nilAccess
    self.argList = argList
}


// declaration Class -- ExpMRetNode
type Nodes_ExpMRetNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_mRet(_env *LnsEnv) *Nodes_Node
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpMRetNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node) *Nodes_ExpMRetNode {
    obj := &Nodes_ExpMRetNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMRetNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpMRetNode) Get_mRet(_env *LnsEnv) *Nodes_Node{ return self.mRet }
// 734: DeclConstr
func (self *Nodes_ExpMRetNode) InitNodes_ExpMRetNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,mRet *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 37, pos, inTestBlock, macroArgFlag, typeList)
    self.mRet = mRet
}


// declaration Class -- ExpAccessMRetNode
type Nodes_ExpAccessMRetNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_index(_env *LnsEnv) LnsInt
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_mRet(_env *LnsEnv) *Nodes_Node
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpAccessMRetNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 LnsInt) *Nodes_ExpAccessMRetNode {
    obj := &Nodes_ExpAccessMRetNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpAccessMRetNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_ExpAccessMRetNode) Get_mRet(_env *LnsEnv) *Nodes_Node{ return self.mRet }
func (self *Nodes_ExpAccessMRetNode) Get_index(_env *LnsEnv) LnsInt{ return self.index }
// 734: DeclConstr
func (self *Nodes_ExpAccessMRetNode) InitNodes_ExpAccessMRetNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,mRet *Nodes_Node,index LnsInt) {
    self.InitNodes_Node(_env, managerId, id, 38, pos, inTestBlock, macroArgFlag, typeList)
    self.mRet = mRet
    self.index = index
}


// declaration Class -- ExpMultiTo1Node
type Nodes_ExpMultiTo1NodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpMultiTo1Node(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node) *Nodes_ExpMultiTo1Node {
    obj := &Nodes_ExpMultiTo1Node{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMultiTo1Node(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpMultiTo1Node) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
// 734: DeclConstr
func (self *Nodes_ExpMultiTo1Node) InitNodes_ExpMultiTo1Node(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 39, pos, inTestBlock, macroArgFlag, typeList)
    self.exp = exp
}


// declaration Class -- ExpParenNode
type Nodes_ExpParenNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpParenNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node) *Nodes_ExpParenNode {
    obj := &Nodes_ExpParenNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpParenNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpParenNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
// 734: DeclConstr
func (self *Nodes_ExpParenNode) InitNodes_ExpParenNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 40, pos, inTestBlock, macroArgFlag, typeList)
    self.exp = exp
}


// declaration Class -- ExpMacroExpNode
type Nodes_ExpMacroExpNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_macroType(_env *LnsEnv) *Ast_TypeInfo
    Get_pos(_env *LnsEnv) Types_Position
    Get_stmtList(_env *LnsEnv) *LnsList
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ExpMacroExpNode struct {
    Nodes_Node
    macroType *Ast_TypeInfo
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
func NewNodes_ExpMacroExpNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Ast_TypeInfo, arg8 *LnsList) *Nodes_ExpMacroExpNode {
    obj := &Nodes_ExpMacroExpNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMacroExpNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_ExpMacroExpNode) Get_macroType(_env *LnsEnv) *Ast_TypeInfo{ return self.macroType }
func (self *Nodes_ExpMacroExpNode) Get_stmtList(_env *LnsEnv) *LnsList{ return self.stmtList }
// 734: DeclConstr
func (self *Nodes_ExpMacroExpNode) InitNodes_ExpMacroExpNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,macroType *Ast_TypeInfo,stmtList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 41, pos, inTestBlock, macroArgFlag, typeList)
    self.macroType = macroType
    self.stmtList = stmtList
}


// declaration Class -- ExpMacroStatNode
type Nodes_ExpMacroStatNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expStrList(_env *LnsEnv) *LnsList
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpMacroStatNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *LnsList) *Nodes_ExpMacroStatNode {
    obj := &Nodes_ExpMacroStatNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMacroStatNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpMacroStatNode) Get_expStrList(_env *LnsEnv) *LnsList{ return self.expStrList }
// 734: DeclConstr
func (self *Nodes_ExpMacroStatNode) InitNodes_ExpMacroStatNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expStrList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 42, pos, inTestBlock, macroArgFlag, typeList)
    self.expStrList = expStrList
}


// declaration Class -- ExpMacroArgExpNode
type Nodes_ExpMacroArgExpNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_codeTxt(_env *LnsEnv) string
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpMacroArgExpNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 string) *Nodes_ExpMacroArgExpNode {
    obj := &Nodes_ExpMacroArgExpNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMacroArgExpNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpMacroArgExpNode) Get_codeTxt(_env *LnsEnv) string{ return self.codeTxt }
// 734: DeclConstr
func (self *Nodes_ExpMacroArgExpNode) InitNodes_ExpMacroArgExpNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,codeTxt string) {
    self.InitNodes_Node(_env, managerId, id, 43, pos, inTestBlock, macroArgFlag, typeList)
    self.codeTxt = codeTxt
}


// declaration Class -- StmtExpNode
type Nodes_StmtExpNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_StmtExpNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node) *Nodes_StmtExpNode {
    obj := &Nodes_StmtExpNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_StmtExpNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_StmtExpNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
// 734: DeclConstr
func (self *Nodes_StmtExpNode) InitNodes_StmtExpNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 44, pos, inTestBlock, macroArgFlag, typeList)
    self.exp = exp
}


// declaration Class -- ExpMacroStatListNode
type Nodes_ExpMacroStatListNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpMacroStatListNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node) *Nodes_ExpMacroStatListNode {
    obj := &Nodes_ExpMacroStatListNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpMacroStatListNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ExpMacroStatListNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
// 734: DeclConstr
func (self *Nodes_ExpMacroStatListNode) InitNodes_ExpMacroStatListNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 45, pos, inTestBlock, macroArgFlag, typeList)
    self.exp = exp
}


// declaration Class -- ExpOmitEnumNode
type Nodes_ExpOmitEnumNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_aliasType(_env *LnsEnv) LnsAny
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_enumTypeInfo(_env *LnsEnv) *Ast_EnumTypeInfo
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_valInfo(_env *LnsEnv) *Ast_EnumValInfo
    Get_valToken(_env *LnsEnv) *Types_Token
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpOmitEnumNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 *Ast_EnumValInfo, arg9 LnsAny, arg10 *Ast_EnumTypeInfo) *Nodes_ExpOmitEnumNode {
    obj := &Nodes_ExpOmitEnumNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpOmitEnumNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_ExpOmitEnumNode) Get_valToken(_env *LnsEnv) *Types_Token{ return self.valToken }
func (self *Nodes_ExpOmitEnumNode) Get_valInfo(_env *LnsEnv) *Ast_EnumValInfo{ return self.valInfo }
func (self *Nodes_ExpOmitEnumNode) Get_aliasType(_env *LnsEnv) LnsAny{ return self.aliasType }
func (self *Nodes_ExpOmitEnumNode) Get_enumTypeInfo(_env *LnsEnv) *Ast_EnumTypeInfo{ return self.enumTypeInfo }
// 734: DeclConstr
func (self *Nodes_ExpOmitEnumNode) InitNodes_ExpOmitEnumNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,valToken *Types_Token,valInfo *Ast_EnumValInfo,aliasType LnsAny,enumTypeInfo *Ast_EnumTypeInfo) {
    self.InitNodes_Node(_env, managerId, id, 46, pos, inTestBlock, macroArgFlag, typeList)
    self.valToken = valToken
    self.valInfo = valInfo
    self.aliasType = aliasType
    self.enumTypeInfo = enumTypeInfo
}


// declaration Class -- RefFieldNode
type Nodes_RefFieldNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_field(_env *LnsEnv) *Types_Token
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_nilAccess(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_prefix(_env *LnsEnv) *Nodes_Node
    Get_symbolInfo(_env *LnsEnv) LnsAny
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_RefFieldNode struct {
    Nodes_Node
    field *Types_Token
    symbolInfo LnsAny
    nilAccess bool
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
func NewNodes_RefFieldNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 LnsAny, arg9 bool, arg10 *Nodes_Node) *Nodes_RefFieldNode {
    obj := &Nodes_RefFieldNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_RefFieldNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_RefFieldNode) Get_field(_env *LnsEnv) *Types_Token{ return self.field }
func (self *Nodes_RefFieldNode) Get_symbolInfo(_env *LnsEnv) LnsAny{ return self.symbolInfo }
func (self *Nodes_RefFieldNode) Get_nilAccess(_env *LnsEnv) bool{ return self.nilAccess }
func (self *Nodes_RefFieldNode) Get_prefix(_env *LnsEnv) *Nodes_Node{ return self.prefix }
// 734: DeclConstr
func (self *Nodes_RefFieldNode) InitNodes_RefFieldNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,field *Types_Token,symbolInfo LnsAny,nilAccess bool,prefix *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 47, pos, inTestBlock, macroArgFlag, typeList)
    self.field = field
    self.symbolInfo = symbolInfo
    self.nilAccess = nilAccess
    self.prefix = prefix
}


// declaration Class -- GetFieldNode
type Nodes_GetFieldNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_field(_env *LnsEnv) *Types_Token
    Get_getterTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_nilAccess(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_prefix(_env *LnsEnv) *Nodes_Node
    Get_symbolInfo(_env *LnsEnv) LnsAny
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_GetFieldNode struct {
    Nodes_Node
    field *Types_Token
    symbolInfo LnsAny
    nilAccess bool
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
func NewNodes_GetFieldNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 LnsAny, arg9 bool, arg10 *Nodes_Node, arg11 *Ast_TypeInfo) *Nodes_GetFieldNode {
    obj := &Nodes_GetFieldNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_GetFieldNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
    return obj
}
func (self *Nodes_GetFieldNode) Get_field(_env *LnsEnv) *Types_Token{ return self.field }
func (self *Nodes_GetFieldNode) Get_symbolInfo(_env *LnsEnv) LnsAny{ return self.symbolInfo }
func (self *Nodes_GetFieldNode) Get_nilAccess(_env *LnsEnv) bool{ return self.nilAccess }
func (self *Nodes_GetFieldNode) Get_prefix(_env *LnsEnv) *Nodes_Node{ return self.prefix }
func (self *Nodes_GetFieldNode) Get_getterTypeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.getterTypeInfo }
// 734: DeclConstr
func (self *Nodes_GetFieldNode) InitNodes_GetFieldNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,field *Types_Token,symbolInfo LnsAny,nilAccess bool,prefix *Nodes_Node,getterTypeInfo *Ast_TypeInfo) {
    self.InitNodes_Node(_env, managerId, id, 48, pos, inTestBlock, macroArgFlag, typeList)
    self.field = field
    self.symbolInfo = symbolInfo
    self.nilAccess = nilAccess
    self.prefix = prefix
    self.getterTypeInfo = getterTypeInfo
}


// declaration Class -- AliasNode
type Nodes_AliasNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_newSymbol(_env *LnsEnv) *Ast_SymbolInfo
    Get_pos(_env *LnsEnv) Types_Position
    Get_srcNode(_env *LnsEnv) *Nodes_Node
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_AliasNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Ast_SymbolInfo, arg8 *Nodes_Node, arg9 *Ast_TypeInfo) *Nodes_AliasNode {
    obj := &Nodes_AliasNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_AliasNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_AliasNode) Get_newSymbol(_env *LnsEnv) *Ast_SymbolInfo{ return self.newSymbol }
func (self *Nodes_AliasNode) Get_srcNode(_env *LnsEnv) *Nodes_Node{ return self.srcNode }
func (self *Nodes_AliasNode) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
// 734: DeclConstr
func (self *Nodes_AliasNode) InitNodes_AliasNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,newSymbol *Ast_SymbolInfo,srcNode *Nodes_Node,typeInfo *Ast_TypeInfo) {
    self.InitNodes_Node(_env, managerId, id, 49, pos, inTestBlock, macroArgFlag, typeList)
    self.newSymbol = newSymbol
    self.srcNode = srcNode
    self.typeInfo = typeInfo
}


// declaration Class -- VarInfo
type Nodes_VarInfoMtd interface {
    Get_actualType(_env *LnsEnv) *Ast_TypeInfo
    Get_name(_env *LnsEnv) *Types_Token
    Get_refType(_env *LnsEnv) LnsAny
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
func NewNodes_VarInfo(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny, arg3 *Ast_TypeInfo) *Nodes_VarInfo {
    obj := &Nodes_VarInfo{}
    obj.FP = obj
    obj.InitNodes_VarInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Nodes_VarInfo) InitNodes_VarInfo(_env *LnsEnv, arg1 *Types_Token, arg2 LnsAny, arg3 *Ast_TypeInfo) {
    self.name = arg1
    self.refType = arg2
    self.actualType = arg3
}
func (self *Nodes_VarInfo) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_VarInfo) Get_refType(_env *LnsEnv) LnsAny{ return self.refType }
func (self *Nodes_VarInfo) Get_actualType(_env *LnsEnv) *Ast_TypeInfo{ return self.actualType }

// declaration Class -- DeclVarNode
type Nodes_DeclVarNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) LnsAny
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_mode(_env *LnsEnv) LnsInt
    Get_pos(_env *LnsEnv) Types_Position
    Get_staticFlag(_env *LnsEnv) bool
    Get_symbolInfoList(_env *LnsEnv) *LnsList
    Get_syncBlock(_env *LnsEnv) LnsAny
    Get_syncVarList(_env *LnsEnv) *LnsList
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_thenBlock(_env *LnsEnv) LnsAny
    Get_typeInfoList(_env *LnsEnv) *LnsList
    Get_unwrapBlock(_env *LnsEnv) LnsAny
    Get_unwrapFlag(_env *LnsEnv) bool
    Get_varList(_env *LnsEnv) *LnsList
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclVarNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 LnsInt, arg9 bool, arg10 *LnsList, arg11 LnsAny, arg12 *LnsList, arg13 *LnsList, arg14 bool, arg15 LnsAny, arg16 LnsAny, arg17 *LnsList, arg18 LnsAny) *Nodes_DeclVarNode {
    obj := &Nodes_DeclVarNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclVarNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18)
    return obj
}
func (self *Nodes_DeclVarNode) Get_mode(_env *LnsEnv) LnsInt{ return self.mode }
func (self *Nodes_DeclVarNode) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Nodes_DeclVarNode) Get_staticFlag(_env *LnsEnv) bool{ return self.staticFlag }
func (self *Nodes_DeclVarNode) Get_varList(_env *LnsEnv) *LnsList{ return self.varList }
func (self *Nodes_DeclVarNode) Get_expList(_env *LnsEnv) LnsAny{ return self.expList }
func (self *Nodes_DeclVarNode) Get_symbolInfoList(_env *LnsEnv) *LnsList{ return self.symbolInfoList }
func (self *Nodes_DeclVarNode) Get_typeInfoList(_env *LnsEnv) *LnsList{ return self.typeInfoList }
func (self *Nodes_DeclVarNode) Get_unwrapFlag(_env *LnsEnv) bool{ return self.unwrapFlag }
func (self *Nodes_DeclVarNode) Get_unwrapBlock(_env *LnsEnv) LnsAny{ return self.unwrapBlock }
func (self *Nodes_DeclVarNode) Get_thenBlock(_env *LnsEnv) LnsAny{ return self.thenBlock }
func (self *Nodes_DeclVarNode) Get_syncVarList(_env *LnsEnv) *LnsList{ return self.syncVarList }
func (self *Nodes_DeclVarNode) Get_syncBlock(_env *LnsEnv) LnsAny{ return self.syncBlock }
// 734: DeclConstr
func (self *Nodes_DeclVarNode) InitNodes_DeclVarNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,mode LnsInt,accessMode LnsInt,staticFlag bool,varList *LnsList,expList LnsAny,symbolInfoList *LnsList,typeInfoList *LnsList,unwrapFlag bool,unwrapBlock LnsAny,thenBlock LnsAny,syncVarList *LnsList,syncBlock LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 50, pos, inTestBlock, macroArgFlag, typeList)
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


// declaration Class -- DeclFuncInfo
type Nodes_DeclFuncInfoMtd interface {
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_argList(_env *LnsEnv) *LnsList
    Get_asyncMode(_env *LnsEnv) LnsAny
    Get_body(_env *LnsEnv) LnsAny
    Get_classTypeInfo(_env *LnsEnv) LnsAny
    Get_declClassNode(_env *LnsEnv) LnsAny
    Get_has__func__Symbol(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_name(_env *LnsEnv) LnsAny
    Get_outsizeOfClass(_env *LnsEnv) bool
    Get_overrideFlag(_env *LnsEnv) bool
    Get_retTypeInfoList(_env *LnsEnv) *LnsList
    Get_retTypeNodeList(_env *LnsEnv) *LnsList
    Get_staticFlag(_env *LnsEnv) bool
    Get_stmtNum(_env *LnsEnv) LnsInt
    Get_symbol(_env *LnsEnv) LnsAny
    Set_body(_env *LnsEnv, arg1 LnsAny)
    Set_has__func__Symbol(_env *LnsEnv, arg1 bool)
    Set_stmtNum(_env *LnsEnv, arg1 LnsInt)
}
type Nodes_DeclFuncInfo struct {
    kind LnsInt
    classTypeInfo LnsAny
    declClassNode LnsAny
    outsizeOfClass bool
    name LnsAny
    symbol LnsAny
    argList *LnsList
    staticFlag bool
    accessMode LnsInt
    asyncMode LnsAny
    body LnsAny
    retTypeInfoList *LnsList
    retTypeNodeList *LnsList
    has__func__Symbol bool
    overrideFlag bool
    stmtNum LnsInt
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
func NewNodes_DeclFuncInfo(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 *LnsList, arg8 bool, arg9 LnsInt, arg10 LnsAny, arg11 LnsAny, arg12 *LnsList, arg13 *LnsList, arg14 bool, arg15 bool, arg16 LnsInt) *Nodes_DeclFuncInfo {
    obj := &Nodes_DeclFuncInfo{}
    obj.FP = obj
    obj.InitNodes_DeclFuncInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16)
    return obj
}
func (self *Nodes_DeclFuncInfo) InitNodes_DeclFuncInfo(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 LnsAny, arg4 bool, arg5 LnsAny, arg6 LnsAny, arg7 *LnsList, arg8 bool, arg9 LnsInt, arg10 LnsAny, arg11 LnsAny, arg12 *LnsList, arg13 *LnsList, arg14 bool, arg15 bool, arg16 LnsInt) {
    self.kind = arg1
    self.classTypeInfo = arg2
    self.declClassNode = arg3
    self.outsizeOfClass = arg4
    self.name = arg5
    self.symbol = arg6
    self.argList = arg7
    self.staticFlag = arg8
    self.accessMode = arg9
    self.asyncMode = arg10
    self.body = arg11
    self.retTypeInfoList = arg12
    self.retTypeNodeList = arg13
    self.has__func__Symbol = arg14
    self.overrideFlag = arg15
    self.stmtNum = arg16
}
func (self *Nodes_DeclFuncInfo) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *Nodes_DeclFuncInfo) Get_classTypeInfo(_env *LnsEnv) LnsAny{ return self.classTypeInfo }
func (self *Nodes_DeclFuncInfo) Get_declClassNode(_env *LnsEnv) LnsAny{ return self.declClassNode }
func (self *Nodes_DeclFuncInfo) Get_outsizeOfClass(_env *LnsEnv) bool{ return self.outsizeOfClass }
func (self *Nodes_DeclFuncInfo) Get_name(_env *LnsEnv) LnsAny{ return self.name }
func (self *Nodes_DeclFuncInfo) Get_symbol(_env *LnsEnv) LnsAny{ return self.symbol }
func (self *Nodes_DeclFuncInfo) Get_argList(_env *LnsEnv) *LnsList{ return self.argList }
func (self *Nodes_DeclFuncInfo) Get_staticFlag(_env *LnsEnv) bool{ return self.staticFlag }
func (self *Nodes_DeclFuncInfo) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Nodes_DeclFuncInfo) Get_asyncMode(_env *LnsEnv) LnsAny{ return self.asyncMode }
func (self *Nodes_DeclFuncInfo) Get_body(_env *LnsEnv) LnsAny{ return self.body }
func (self *Nodes_DeclFuncInfo) Set_body(_env *LnsEnv, arg1 LnsAny){ self.body = arg1 }
func (self *Nodes_DeclFuncInfo) Get_retTypeInfoList(_env *LnsEnv) *LnsList{ return self.retTypeInfoList }
func (self *Nodes_DeclFuncInfo) Get_retTypeNodeList(_env *LnsEnv) *LnsList{ return self.retTypeNodeList }
func (self *Nodes_DeclFuncInfo) Get_has__func__Symbol(_env *LnsEnv) bool{ return self.has__func__Symbol }
func (self *Nodes_DeclFuncInfo) Set_has__func__Symbol(_env *LnsEnv, arg1 bool){ self.has__func__Symbol = arg1 }
func (self *Nodes_DeclFuncInfo) Get_overrideFlag(_env *LnsEnv) bool{ return self.overrideFlag }
func (self *Nodes_DeclFuncInfo) Get_stmtNum(_env *LnsEnv) LnsInt{ return self.stmtNum }
func (self *Nodes_DeclFuncInfo) Set_stmtNum(_env *LnsEnv, arg1 LnsInt){ self.stmtNum = arg1 }

// declaration Class -- DeclFormNode
type Nodes_DeclFormNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_DeclFormNode struct {
    Nodes_Node
    declInfo *Nodes_DeclFuncInfo
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
func NewNodes_DeclFormNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_DeclFuncInfo) *Nodes_DeclFormNode {
    obj := &Nodes_DeclFormNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclFormNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_DeclFormNode) Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo{ return self.declInfo }
// 734: DeclConstr
func (self *Nodes_DeclFormNode) InitNodes_DeclFormNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(_env, managerId, id, 51, pos, inTestBlock, macroArgFlag, typeList)
    self.declInfo = declInfo
}


// declaration Class -- DeclFuncNode
type Nodes_DeclFuncNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclFuncNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_DeclFuncInfo) *Nodes_DeclFuncNode {
    obj := &Nodes_DeclFuncNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclFuncNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_DeclFuncNode) Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo{ return self.declInfo }
// 734: DeclConstr
func (self *Nodes_DeclFuncNode) InitNodes_DeclFuncNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(_env, managerId, id, 52, pos, inTestBlock, macroArgFlag, typeList)
    self.declInfo = declInfo
}


// declaration Class -- DeclMethodNode
type Nodes_DeclMethodNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclMethodNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_DeclFuncInfo) *Nodes_DeclMethodNode {
    obj := &Nodes_DeclMethodNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclMethodNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_DeclMethodNode) Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo{ return self.declInfo }
// 734: DeclConstr
func (self *Nodes_DeclMethodNode) InitNodes_DeclMethodNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(_env, managerId, id, 53, pos, inTestBlock, macroArgFlag, typeList)
    self.declInfo = declInfo
}


// declaration Class -- ProtoMethodNode
type Nodes_ProtoMethodNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ProtoMethodNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_DeclFuncInfo) *Nodes_ProtoMethodNode {
    obj := &Nodes_ProtoMethodNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ProtoMethodNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_ProtoMethodNode) Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo{ return self.declInfo }
// 734: DeclConstr
func (self *Nodes_ProtoMethodNode) InitNodes_ProtoMethodNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(_env, managerId, id, 54, pos, inTestBlock, macroArgFlag, typeList)
    self.declInfo = declInfo
}


// declaration Class -- DeclConstrNode
type Nodes_DeclConstrNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclConstrNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_DeclFuncInfo) *Nodes_DeclConstrNode {
    obj := &Nodes_DeclConstrNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclConstrNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_DeclConstrNode) Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo{ return self.declInfo }
// 734: DeclConstr
func (self *Nodes_DeclConstrNode) InitNodes_DeclConstrNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(_env, managerId, id, 55, pos, inTestBlock, macroArgFlag, typeList)
    self.declInfo = declInfo
}


// declaration Class -- DeclDestrNode
type Nodes_DeclDestrNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclDestrNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_DeclFuncInfo) *Nodes_DeclDestrNode {
    obj := &Nodes_DeclDestrNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclDestrNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_DeclDestrNode) Get_declInfo(_env *LnsEnv) *Nodes_DeclFuncInfo{ return self.declInfo }
// 734: DeclConstr
func (self *Nodes_DeclDestrNode) InitNodes_DeclDestrNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) {
    self.InitNodes_Node(_env, managerId, id, 56, pos, inTestBlock, macroArgFlag, typeList)
    self.declInfo = declInfo
}


// declaration Class -- ExpCallSuperCtorNode
type Nodes_ExpCallSuperCtorNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) LnsAny
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_methodType(_env *LnsEnv) *Ast_TypeInfo
    Get_pos(_env *LnsEnv) Types_Position
    Get_superType(_env *LnsEnv) *Ast_TypeInfo
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpCallSuperCtorNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Ast_TypeInfo, arg8 *Ast_TypeInfo, arg9 LnsAny) *Nodes_ExpCallSuperCtorNode {
    obj := &Nodes_ExpCallSuperCtorNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpCallSuperCtorNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ExpCallSuperCtorNode) Get_superType(_env *LnsEnv) *Ast_TypeInfo{ return self.superType }
func (self *Nodes_ExpCallSuperCtorNode) Get_methodType(_env *LnsEnv) *Ast_TypeInfo{ return self.methodType }
func (self *Nodes_ExpCallSuperCtorNode) Get_expList(_env *LnsEnv) LnsAny{ return self.expList }
// 734: DeclConstr
func (self *Nodes_ExpCallSuperCtorNode) InitNodes_ExpCallSuperCtorNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,superType *Ast_TypeInfo,methodType *Ast_TypeInfo,expList LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 57, pos, inTestBlock, macroArgFlag, typeList)
    self.superType = superType
    self.methodType = methodType
    self.expList = expList
}


// declaration Class -- ExpCallSuperNode
type Nodes_ExpCallSuperNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) LnsAny
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_methodType(_env *LnsEnv) *Ast_TypeInfo
    Get_pos(_env *LnsEnv) Types_Position
    Get_superType(_env *LnsEnv) *Ast_TypeInfo
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_ExpCallSuperNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Ast_TypeInfo, arg8 *Ast_TypeInfo, arg9 LnsAny) *Nodes_ExpCallSuperNode {
    obj := &Nodes_ExpCallSuperNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ExpCallSuperNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_ExpCallSuperNode) Get_superType(_env *LnsEnv) *Ast_TypeInfo{ return self.superType }
func (self *Nodes_ExpCallSuperNode) Get_methodType(_env *LnsEnv) *Ast_TypeInfo{ return self.methodType }
func (self *Nodes_ExpCallSuperNode) Get_expList(_env *LnsEnv) LnsAny{ return self.expList }
// 734: DeclConstr
func (self *Nodes_ExpCallSuperNode) InitNodes_ExpCallSuperNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,superType *Ast_TypeInfo,methodType *Ast_TypeInfo,expList LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 58, pos, inTestBlock, macroArgFlag, typeList)
    self.superType = superType
    self.methodType = methodType
    self.expList = expList
}


// declaration Class -- AsyncLockNode
type Nodes_AsyncLockNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_lockKind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_AsyncLockNode struct {
    Nodes_Node
    lockKind LnsInt
    block *Nodes_BlockNode
    FP Nodes_AsyncLockNodeMtd
}
func Nodes_AsyncLockNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_AsyncLockNode).FP
}
type Nodes_AsyncLockNodeDownCast interface {
    ToNodes_AsyncLockNode() *Nodes_AsyncLockNode
}
func Nodes_AsyncLockNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_AsyncLockNodeDownCast)
    if ok { return work.ToNodes_AsyncLockNode() }
    return nil
}
func (obj *Nodes_AsyncLockNode) ToNodes_AsyncLockNode() *Nodes_AsyncLockNode {
    return obj
}
func NewNodes_AsyncLockNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Nodes_BlockNode) *Nodes_AsyncLockNode {
    obj := &Nodes_AsyncLockNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_AsyncLockNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_AsyncLockNode) Get_lockKind(_env *LnsEnv) LnsInt{ return self.lockKind }
func (self *Nodes_AsyncLockNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
// 734: DeclConstr
func (self *Nodes_AsyncLockNode) InitNodes_AsyncLockNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,lockKind LnsInt,block *Nodes_BlockNode) {
    self.InitNodes_Node(_env, managerId, id, 59, pos, inTestBlock, macroArgFlag, typeList)
    self.lockKind = lockKind
    self.block = block
}


// declaration Class -- RequestNode
type Nodes_RequestNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_processor(_env *LnsEnv) *Nodes_Node
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_RequestNode struct {
    Nodes_Node
    processor *Nodes_Node
    exp *Nodes_Node
    FP Nodes_RequestNodeMtd
}
func Nodes_RequestNode2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_RequestNode).FP
}
type Nodes_RequestNodeDownCast interface {
    ToNodes_RequestNode() *Nodes_RequestNode
}
func Nodes_RequestNodeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_RequestNodeDownCast)
    if ok { return work.ToNodes_RequestNode() }
    return nil
}
func (obj *Nodes_RequestNode) ToNodes_RequestNode() *Nodes_RequestNode {
    return obj
}
func NewNodes_RequestNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node, arg8 *Nodes_Node) *Nodes_RequestNode {
    obj := &Nodes_RequestNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_RequestNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_RequestNode) Get_processor(_env *LnsEnv) *Nodes_Node{ return self.processor }
func (self *Nodes_RequestNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
// 734: DeclConstr
func (self *Nodes_RequestNode) InitNodes_RequestNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,processor *Nodes_Node,exp *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 60, pos, inTestBlock, macroArgFlag, typeList)
    self.processor = processor
    self.exp = exp
}


// declaration Class -- DeclMemberNode
type Nodes_DeclMemberNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetGetterSym(_env *LnsEnv) LnsAny
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSetterSym(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_classType(_env *LnsEnv) *Ast_TypeInfo
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_getterMode(_env *LnsEnv) LnsInt
    Get_getterMutable(_env *LnsEnv) bool
    Get_getterRetType(_env *LnsEnv) *Ast_TypeInfo
    Get_getterToken(_env *LnsEnv) LnsAny
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) *Types_Token
    Get_pos(_env *LnsEnv) Types_Position
    Get_refType(_env *LnsEnv) *Nodes_RefTypeNode
    Get_setterMode(_env *LnsEnv) LnsInt
    Get_setterToken(_env *LnsEnv) LnsAny
    Get_staticFlag(_env *LnsEnv) bool
    Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
    getterToken LnsAny
    getterRetType *Ast_TypeInfo
    setterMode LnsInt
    setterToken LnsAny
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
func NewNodes_DeclMemberNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 *Nodes_RefTypeNode, arg9 *Ast_SymbolInfo, arg10 *Ast_TypeInfo, arg11 bool, arg12 LnsInt, arg13 bool, arg14 LnsInt, arg15 LnsAny, arg16 *Ast_TypeInfo, arg17 LnsInt, arg18 LnsAny) *Nodes_DeclMemberNode {
    obj := &Nodes_DeclMemberNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclMemberNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18)
    return obj
}
func (self *Nodes_DeclMemberNode) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_DeclMemberNode) Get_refType(_env *LnsEnv) *Nodes_RefTypeNode{ return self.refType }
func (self *Nodes_DeclMemberNode) Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbolInfo }
func (self *Nodes_DeclMemberNode) Get_classType(_env *LnsEnv) *Ast_TypeInfo{ return self.classType }
func (self *Nodes_DeclMemberNode) Get_staticFlag(_env *LnsEnv) bool{ return self.staticFlag }
func (self *Nodes_DeclMemberNode) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Nodes_DeclMemberNode) Get_getterMutable(_env *LnsEnv) bool{ return self.getterMutable }
func (self *Nodes_DeclMemberNode) Get_getterMode(_env *LnsEnv) LnsInt{ return self.getterMode }
func (self *Nodes_DeclMemberNode) Get_getterToken(_env *LnsEnv) LnsAny{ return self.getterToken }
func (self *Nodes_DeclMemberNode) Get_getterRetType(_env *LnsEnv) *Ast_TypeInfo{ return self.getterRetType }
func (self *Nodes_DeclMemberNode) Get_setterMode(_env *LnsEnv) LnsInt{ return self.setterMode }
func (self *Nodes_DeclMemberNode) Get_setterToken(_env *LnsEnv) LnsAny{ return self.setterToken }
// 734: DeclConstr
func (self *Nodes_DeclMemberNode) InitNodes_DeclMemberNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,refType *Nodes_RefTypeNode,symbolInfo *Ast_SymbolInfo,classType *Ast_TypeInfo,staticFlag bool,accessMode LnsInt,getterMutable bool,getterMode LnsInt,getterToken LnsAny,getterRetType *Ast_TypeInfo,setterMode LnsInt,setterToken LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 61, pos, inTestBlock, macroArgFlag, typeList)
    self.name = name
    self.refType = refType
    self.symbolInfo = symbolInfo
    self.classType = classType
    self.staticFlag = staticFlag
    self.accessMode = accessMode
    self.getterMutable = getterMutable
    self.getterMode = getterMode
    self.getterToken = getterToken
    self.getterRetType = getterRetType
    self.setterMode = setterMode
    self.setterToken = setterToken
}


// declaration Class -- DeclArgNode
type Nodes_DeclArgNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_argType(_env *LnsEnv) LnsAny
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) *Types_Token
    Get_pos(_env *LnsEnv) Types_Position
    Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclArgNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 *Ast_SymbolInfo, arg9 LnsAny) *Nodes_DeclArgNode {
    obj := &Nodes_DeclArgNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclArgNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_DeclArgNode) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_DeclArgNode) Get_symbolInfo(_env *LnsEnv) *Ast_SymbolInfo{ return self.symbolInfo }
func (self *Nodes_DeclArgNode) Get_argType(_env *LnsEnv) LnsAny{ return self.argType }
// 734: DeclConstr
func (self *Nodes_DeclArgNode) InitNodes_DeclArgNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,symbolInfo *Ast_SymbolInfo,argType LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 62, pos, inTestBlock, macroArgFlag, typeList)
    self.name = name
    self.symbolInfo = symbolInfo
    self.argType = argType
}


// declaration Class -- DeclArgDDDNode
type Nodes_DeclArgDDDNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclArgDDDNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList) *Nodes_DeclArgDDDNode {
    obj := &Nodes_DeclArgDDDNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclArgDDDNode(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
// 734: DeclConstr
func (self *Nodes_DeclArgDDDNode) InitNodes_DeclArgDDDNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 63, pos, inTestBlock, macroArgFlag, typeList)
}


// declaration Class -- AdvertiseInfo
type Nodes_AdvertiseInfoMtd interface {
    Get_member(_env *LnsEnv) *Nodes_DeclMemberNode
    Get_pos(_env *LnsEnv) Types_Position
    Get_prefix(_env *LnsEnv) string
}
type Nodes_AdvertiseInfo struct {
    member *Nodes_DeclMemberNode
    prefix string
    pos Types_Position
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
func NewNodes_AdvertiseInfo(_env *LnsEnv, arg1 *Nodes_DeclMemberNode, arg2 string, arg3 Types_Position) *Nodes_AdvertiseInfo {
    obj := &Nodes_AdvertiseInfo{}
    obj.FP = obj
    obj.InitNodes_AdvertiseInfo(_env, arg1, arg2, arg3)
    return obj
}
func (self *Nodes_AdvertiseInfo) InitNodes_AdvertiseInfo(_env *LnsEnv, arg1 *Nodes_DeclMemberNode, arg2 string, arg3 Types_Position) {
    self.member = arg1
    self.prefix = arg2
    self.pos = arg3
}
func (self *Nodes_AdvertiseInfo) Get_member(_env *LnsEnv) *Nodes_DeclMemberNode{ return self.member }
func (self *Nodes_AdvertiseInfo) Get_prefix(_env *LnsEnv) string{ return self.prefix }
func (self *Nodes_AdvertiseInfo) Get_pos(_env *LnsEnv) Types_Position{ return self.pos }

// declaration Class -- DeclAdvertiseNode
type Nodes_DeclAdvertiseNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_advInfo(_env *LnsEnv) *Nodes_AdvertiseInfo
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclAdvertiseNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_AdvertiseInfo) *Nodes_DeclAdvertiseNode {
    obj := &Nodes_DeclAdvertiseNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclAdvertiseNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_DeclAdvertiseNode) Get_advInfo(_env *LnsEnv) *Nodes_AdvertiseInfo{ return self.advInfo }
// 734: DeclConstr
func (self *Nodes_DeclAdvertiseNode) InitNodes_DeclAdvertiseNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,advInfo *Nodes_AdvertiseInfo) {
    self.InitNodes_Node(_env, managerId, id, 64, pos, inTestBlock, macroArgFlag, typeList)
    self.advInfo = advInfo
}


// declaration Class -- ClassInheritInfo
type Nodes_ClassInheritInfoMtd interface {
    Get_base(_env *LnsEnv) LnsAny
    Get_impliments(_env *LnsEnv) *LnsList
    Visit(_env *LnsEnv, arg1 *Nodes_Node, arg2 Nodes_NodeVisitor, arg3 LnsInt, arg4 *LnsSet) bool
}
type Nodes_ClassInheritInfo struct {
    base LnsAny
    impliments *LnsList
    FP Nodes_ClassInheritInfoMtd
}
func Nodes_ClassInheritInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ClassInheritInfo).FP
}
type Nodes_ClassInheritInfoDownCast interface {
    ToNodes_ClassInheritInfo() *Nodes_ClassInheritInfo
}
func Nodes_ClassInheritInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ClassInheritInfoDownCast)
    if ok { return work.ToNodes_ClassInheritInfo() }
    return nil
}
func (obj *Nodes_ClassInheritInfo) ToNodes_ClassInheritInfo() *Nodes_ClassInheritInfo {
    return obj
}
func NewNodes_ClassInheritInfo(_env *LnsEnv, arg1 LnsAny, arg2 *LnsList) *Nodes_ClassInheritInfo {
    obj := &Nodes_ClassInheritInfo{}
    obj.FP = obj
    obj.InitNodes_ClassInheritInfo(_env, arg1, arg2)
    return obj
}
func (self *Nodes_ClassInheritInfo) InitNodes_ClassInheritInfo(_env *LnsEnv, arg1 LnsAny, arg2 *LnsList) {
    self.base = arg1
    self.impliments = arg2
}
func (self *Nodes_ClassInheritInfo) Get_base(_env *LnsEnv) LnsAny{ return self.base }
func (self *Nodes_ClassInheritInfo) Get_impliments(_env *LnsEnv) *LnsList{ return self.impliments }

// declaration Class -- ProtoClassNode
type Nodes_ProtoClassNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_inheritInfo(_env *LnsEnv) *Nodes_ClassInheritInfo
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) *Types_Token
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_ProtoClassNode struct {
    Nodes_Node
    name *Types_Token
    inheritInfo *Nodes_ClassInheritInfo
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
func NewNodes_ProtoClassNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 *Nodes_ClassInheritInfo) *Nodes_ProtoClassNode {
    obj := &Nodes_ProtoClassNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_ProtoClassNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_ProtoClassNode) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_ProtoClassNode) Get_inheritInfo(_env *LnsEnv) *Nodes_ClassInheritInfo{ return self.inheritInfo }
// 734: DeclConstr
func (self *Nodes_ProtoClassNode) InitNodes_ProtoClassNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,inheritInfo *Nodes_ClassInheritInfo) {
    self.InitNodes_Node(_env, managerId, id, 65, pos, inTestBlock, macroArgFlag, typeList)
    self.name = name
    self.inheritInfo = inheritInfo
}


// declaration Class -- ClassInitBlockInfo
type Nodes_ClassInitBlockInfoMtd interface {
    Get_func(_env *LnsEnv) LnsAny
    Set_func(_env *LnsEnv, arg1 LnsAny)
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
func NewNodes_ClassInitBlockInfo(_env *LnsEnv, arg1 LnsAny) *Nodes_ClassInitBlockInfo {
    obj := &Nodes_ClassInitBlockInfo{}
    obj.FP = obj
    obj.InitNodes_ClassInitBlockInfo(_env, arg1)
    return obj
}
func (self *Nodes_ClassInitBlockInfo) InitNodes_ClassInitBlockInfo(_env *LnsEnv, arg1 LnsAny) {
    self._func = arg1
}
func (self *Nodes_ClassInitBlockInfo) Get_func(_env *LnsEnv) LnsAny{ return self._func }
func (self *Nodes_ClassInitBlockInfo) Set_func(_env *LnsEnv, arg1 LnsAny){ self._func = arg1 }

// declaration Class -- DeclClassNode
type Nodes_DeclClassNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    CreateMethodNameSetWithoutAdv(_env *LnsEnv) *LnsSet
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_advertiseList(_env *LnsEnv) *LnsList
    Get_allStmtList(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_declStmtList(_env *LnsEnv) *LnsList
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_fieldList(_env *LnsEnv) *LnsList
    Get_gluePrefix(_env *LnsEnv) LnsAny
    Get_hasOldCtor(_env *LnsEnv) bool
    Get_hasPrototype(_env *LnsEnv) bool
    Get_inTestBlock(_env *LnsEnv) bool
    Get_inheritInfo(_env *LnsEnv) *Nodes_ClassInheritInfo
    Get_initBlock(_env *LnsEnv) *Nodes_ClassInitBlockInfo
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_lang(_env *LnsEnv) LnsAny
    Get_lazyLoad(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_memberList(_env *LnsEnv) *LnsList
    Get_moduleName(_env *LnsEnv) LnsAny
    Get_name(_env *LnsEnv) *Types_Token
    Get_outerMethodSet(_env *LnsEnv) *LnsSet
    Get_pos(_env *LnsEnv) Types_Position
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_trustList(_env *LnsEnv) *LnsList
    Get_uninitMemberList(_env *LnsEnv) *LnsList
    HasNilAccess(_env *LnsEnv) bool
    HasUserInit(_env *LnsEnv) bool
    IsModule(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetHasOldCtor(_env *LnsEnv)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_DeclClassNode struct {
    Nodes_Node
    accessMode LnsInt
    name *Types_Token
    inheritInfo *Nodes_ClassInheritInfo
    hasPrototype bool
    gluePrefix LnsAny
    moduleName LnsAny
    lang LnsAny
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
func NewNodes_DeclClassNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Types_Token, arg9 *Nodes_ClassInheritInfo, arg10 bool, arg11 LnsAny, arg12 LnsAny, arg13 LnsAny, arg14 LnsInt, arg15 bool, arg16 *LnsList, arg17 *LnsList, arg18 *LnsList, arg19 *LnsList, arg20 *Ast_Scope, arg21 *Nodes_ClassInitBlockInfo, arg22 *LnsList, arg23 *LnsList, arg24 *LnsList, arg25 *LnsSet) *Nodes_DeclClassNode {
    obj := &Nodes_DeclClassNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclClassNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16, arg17, arg18, arg19, arg20, arg21, arg22, arg23, arg24, arg25)
    return obj
}
func (self *Nodes_DeclClassNode) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Nodes_DeclClassNode) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_DeclClassNode) Get_inheritInfo(_env *LnsEnv) *Nodes_ClassInheritInfo{ return self.inheritInfo }
func (self *Nodes_DeclClassNode) Get_hasPrototype(_env *LnsEnv) bool{ return self.hasPrototype }
func (self *Nodes_DeclClassNode) Get_gluePrefix(_env *LnsEnv) LnsAny{ return self.gluePrefix }
func (self *Nodes_DeclClassNode) Get_moduleName(_env *LnsEnv) LnsAny{ return self.moduleName }
func (self *Nodes_DeclClassNode) Get_lang(_env *LnsEnv) LnsAny{ return self.lang }
func (self *Nodes_DeclClassNode) Get_lazyLoad(_env *LnsEnv) LnsInt{ return self.lazyLoad }
func (self *Nodes_DeclClassNode) Get_hasOldCtor(_env *LnsEnv) bool{ return self.hasOldCtor }
func (self *Nodes_DeclClassNode) Get_allStmtList(_env *LnsEnv) *LnsList{ return self.allStmtList }
func (self *Nodes_DeclClassNode) Get_declStmtList(_env *LnsEnv) *LnsList{ return self.declStmtList }
func (self *Nodes_DeclClassNode) Get_fieldList(_env *LnsEnv) *LnsList{ return self.fieldList }
func (self *Nodes_DeclClassNode) Get_memberList(_env *LnsEnv) *LnsList{ return self.memberList }
func (self *Nodes_DeclClassNode) Get_scope(_env *LnsEnv) *Ast_Scope{ return self.scope }
func (self *Nodes_DeclClassNode) Get_initBlock(_env *LnsEnv) *Nodes_ClassInitBlockInfo{ return self.initBlock }
func (self *Nodes_DeclClassNode) Get_advertiseList(_env *LnsEnv) *LnsList{ return self.advertiseList }
func (self *Nodes_DeclClassNode) Get_trustList(_env *LnsEnv) *LnsList{ return self.trustList }
func (self *Nodes_DeclClassNode) Get_uninitMemberList(_env *LnsEnv) *LnsList{ return self.uninitMemberList }
func (self *Nodes_DeclClassNode) Get_outerMethodSet(_env *LnsEnv) *LnsSet{ return self.outerMethodSet }
// 734: DeclConstr
func (self *Nodes_DeclClassNode) InitNodes_DeclClassNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,accessMode LnsInt,name *Types_Token,inheritInfo *Nodes_ClassInheritInfo,hasPrototype bool,gluePrefix LnsAny,moduleName LnsAny,lang LnsAny,lazyLoad LnsInt,hasOldCtor bool,allStmtList *LnsList,declStmtList *LnsList,fieldList *LnsList,memberList *LnsList,scope *Ast_Scope,initBlock *Nodes_ClassInitBlockInfo,advertiseList *LnsList,trustList *LnsList,uninitMemberList *LnsList,outerMethodSet *LnsSet) {
    self.InitNodes_Node(_env, managerId, id, 66, pos, inTestBlock, macroArgFlag, typeList)
    self.accessMode = accessMode
    self.name = name
    self.inheritInfo = inheritInfo
    self.hasPrototype = hasPrototype
    self.gluePrefix = gluePrefix
    self.moduleName = moduleName
    self.lang = lang
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


// declaration Class -- DeclEnumNode
type Nodes_DeclEnumNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_enumType(_env *LnsEnv) *Ast_EnumTypeInfo
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) *Types_Token
    Get_pos(_env *LnsEnv) Types_Position
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_valueNameList(_env *LnsEnv) *LnsList
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclEnumNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Ast_EnumTypeInfo, arg8 LnsInt, arg9 *Types_Token, arg10 *LnsList, arg11 *Ast_Scope) *Nodes_DeclEnumNode {
    obj := &Nodes_DeclEnumNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclEnumNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
    return obj
}
func (self *Nodes_DeclEnumNode) Get_enumType(_env *LnsEnv) *Ast_EnumTypeInfo{ return self.enumType }
func (self *Nodes_DeclEnumNode) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Nodes_DeclEnumNode) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_DeclEnumNode) Get_valueNameList(_env *LnsEnv) *LnsList{ return self.valueNameList }
func (self *Nodes_DeclEnumNode) Get_scope(_env *LnsEnv) *Ast_Scope{ return self.scope }
// 734: DeclConstr
func (self *Nodes_DeclEnumNode) InitNodes_DeclEnumNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,enumType *Ast_EnumTypeInfo,accessMode LnsInt,name *Types_Token,valueNameList *LnsList,scope *Ast_Scope) {
    self.InitNodes_Node(_env, managerId, id, 67, pos, inTestBlock, macroArgFlag, typeList)
    self.enumType = enumType
    self.accessMode = accessMode
    self.name = name
    self.valueNameList = valueNameList
    self.scope = scope
}


// declaration Class -- AlgeValParamInfo
type Nodes_AlgeValParamInfoMtd interface {
    Get_name(_env *LnsEnv) LnsAny
    Get_typeRef(_env *LnsEnv) *Nodes_RefTypeNode
}
type Nodes_AlgeValParamInfo struct {
    name LnsAny
    typeRef *Nodes_RefTypeNode
    FP Nodes_AlgeValParamInfoMtd
}
func Nodes_AlgeValParamInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_AlgeValParamInfo).FP
}
type Nodes_AlgeValParamInfoDownCast interface {
    ToNodes_AlgeValParamInfo() *Nodes_AlgeValParamInfo
}
func Nodes_AlgeValParamInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_AlgeValParamInfoDownCast)
    if ok { return work.ToNodes_AlgeValParamInfo() }
    return nil
}
func (obj *Nodes_AlgeValParamInfo) ToNodes_AlgeValParamInfo() *Nodes_AlgeValParamInfo {
    return obj
}
func NewNodes_AlgeValParamInfo(_env *LnsEnv, arg1 LnsAny, arg2 *Nodes_RefTypeNode) *Nodes_AlgeValParamInfo {
    obj := &Nodes_AlgeValParamInfo{}
    obj.FP = obj
    obj.InitNodes_AlgeValParamInfo(_env, arg1, arg2)
    return obj
}
func (self *Nodes_AlgeValParamInfo) InitNodes_AlgeValParamInfo(_env *LnsEnv, arg1 LnsAny, arg2 *Nodes_RefTypeNode) {
    self.name = arg1
    self.typeRef = arg2
}
func (self *Nodes_AlgeValParamInfo) Get_name(_env *LnsEnv) LnsAny{ return self.name }
func (self *Nodes_AlgeValParamInfo) Get_typeRef(_env *LnsEnv) *Nodes_RefTypeNode{ return self.typeRef }

// declaration Class -- DeclAlgeValInfo
type Nodes_DeclAlgeValInfoMtd interface {
    Get_paramList(_env *LnsEnv) *LnsList
    Get_valSym(_env *LnsEnv) *Ast_SymbolInfo
}
type Nodes_DeclAlgeValInfo struct {
    valSym *Ast_SymbolInfo
    paramList *LnsList
    FP Nodes_DeclAlgeValInfoMtd
}
func Nodes_DeclAlgeValInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_DeclAlgeValInfo).FP
}
type Nodes_DeclAlgeValInfoDownCast interface {
    ToNodes_DeclAlgeValInfo() *Nodes_DeclAlgeValInfo
}
func Nodes_DeclAlgeValInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_DeclAlgeValInfoDownCast)
    if ok { return work.ToNodes_DeclAlgeValInfo() }
    return nil
}
func (obj *Nodes_DeclAlgeValInfo) ToNodes_DeclAlgeValInfo() *Nodes_DeclAlgeValInfo {
    return obj
}
func NewNodes_DeclAlgeValInfo(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 *LnsList) *Nodes_DeclAlgeValInfo {
    obj := &Nodes_DeclAlgeValInfo{}
    obj.FP = obj
    obj.InitNodes_DeclAlgeValInfo(_env, arg1, arg2)
    return obj
}
func (self *Nodes_DeclAlgeValInfo) InitNodes_DeclAlgeValInfo(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 *LnsList) {
    self.valSym = arg1
    self.paramList = arg2
}
func (self *Nodes_DeclAlgeValInfo) Get_valSym(_env *LnsEnv) *Ast_SymbolInfo{ return self.valSym }
func (self *Nodes_DeclAlgeValInfo) Get_paramList(_env *LnsEnv) *LnsList{ return self.paramList }

// declaration Class -- DeclAlgeNode
type Nodes_DeclAlgeNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_accessMode(_env *LnsEnv) LnsInt
    Get_algeType(_env *LnsEnv) *Ast_AlgeTypeInfo
    Get_algeValList(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) *Types_Token
    Get_pos(_env *LnsEnv) Types_Position
    Get_scope(_env *LnsEnv) *Ast_Scope
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_DeclAlgeNode struct {
    Nodes_Node
    accessMode LnsInt
    algeType *Ast_AlgeTypeInfo
    name *Types_Token
    algeValList *LnsList
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
func NewNodes_DeclAlgeNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Ast_AlgeTypeInfo, arg9 *Types_Token, arg10 *LnsList, arg11 *Ast_Scope) *Nodes_DeclAlgeNode {
    obj := &Nodes_DeclAlgeNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclAlgeNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
    return obj
}
func (self *Nodes_DeclAlgeNode) Get_accessMode(_env *LnsEnv) LnsInt{ return self.accessMode }
func (self *Nodes_DeclAlgeNode) Get_algeType(_env *LnsEnv) *Ast_AlgeTypeInfo{ return self.algeType }
func (self *Nodes_DeclAlgeNode) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_DeclAlgeNode) Get_algeValList(_env *LnsEnv) *LnsList{ return self.algeValList }
func (self *Nodes_DeclAlgeNode) Get_scope(_env *LnsEnv) *Ast_Scope{ return self.scope }
// 734: DeclConstr
func (self *Nodes_DeclAlgeNode) InitNodes_DeclAlgeNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,accessMode LnsInt,algeType *Ast_AlgeTypeInfo,name *Types_Token,algeValList *LnsList,scope *Ast_Scope) {
    self.InitNodes_Node(_env, managerId, id, 68, pos, inTestBlock, macroArgFlag, typeList)
    self.accessMode = accessMode
    self.algeType = algeType
    self.name = name
    self.algeValList = algeValList
    self.scope = scope
}


// declaration Class -- NewAlgeValNode
type Nodes_NewAlgeValNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_algeTypeInfo(_env *LnsEnv) *Ast_AlgeTypeInfo
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) *Types_Token
    Get_paramList(_env *LnsEnv) *LnsList
    Get_pos(_env *LnsEnv) Types_Position
    Get_prefix(_env *LnsEnv) LnsAny
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_valInfo(_env *LnsEnv) *Ast_AlgeValInfo
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_NewAlgeValNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 LnsAny, arg9 *Ast_AlgeTypeInfo, arg10 *Ast_AlgeValInfo, arg11 *LnsList) *Nodes_NewAlgeValNode {
    obj := &Nodes_NewAlgeValNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_NewAlgeValNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
    return obj
}
func (self *Nodes_NewAlgeValNode) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_NewAlgeValNode) Get_prefix(_env *LnsEnv) LnsAny{ return self.prefix }
func (self *Nodes_NewAlgeValNode) Get_algeTypeInfo(_env *LnsEnv) *Ast_AlgeTypeInfo{ return self.algeTypeInfo }
func (self *Nodes_NewAlgeValNode) Get_valInfo(_env *LnsEnv) *Ast_AlgeValInfo{ return self.valInfo }
func (self *Nodes_NewAlgeValNode) Get_paramList(_env *LnsEnv) *LnsList{ return self.paramList }
// 734: DeclConstr
func (self *Nodes_NewAlgeValNode) InitNodes_NewAlgeValNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,prefix LnsAny,algeTypeInfo *Ast_AlgeTypeInfo,valInfo *Ast_AlgeValInfo,paramList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 69, pos, inTestBlock, macroArgFlag, typeList)
    self.name = name
    self.prefix = prefix
    self.algeTypeInfo = algeTypeInfo
    self.valInfo = valInfo
    self.paramList = paramList
}


// declaration Class -- LuneControlNode
type Nodes_LuneControlNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_pragma(_env *LnsEnv) LnsAny
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LuneControlNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsAny) *Nodes_LuneControlNode {
    obj := &Nodes_LuneControlNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LuneControlNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_LuneControlNode) Get_pragma(_env *LnsEnv) LnsAny{ return self.pragma }
// 734: DeclConstr
func (self *Nodes_LuneControlNode) InitNodes_LuneControlNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,pragma LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 70, pos, inTestBlock, macroArgFlag, typeList)
    self.pragma = pragma
}


// declaration Class -- MatchCase
type Nodes_MatchCaseMtd interface {
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_valExpRef(_env *LnsEnv) *Nodes_ExpRefNode
    Get_valInfo(_env *LnsEnv) *Ast_AlgeValInfo
    Get_valParamNameList(_env *LnsEnv) *LnsList
}
type Nodes_MatchCase struct {
    valInfo *Ast_AlgeValInfo
    valExpRef *Nodes_ExpRefNode
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
func NewNodes_MatchCase(_env *LnsEnv, arg1 *Ast_AlgeValInfo, arg2 *Nodes_ExpRefNode, arg3 *LnsList, arg4 *Nodes_BlockNode) *Nodes_MatchCase {
    obj := &Nodes_MatchCase{}
    obj.FP = obj
    obj.InitNodes_MatchCase(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *Nodes_MatchCase) InitNodes_MatchCase(_env *LnsEnv, arg1 *Ast_AlgeValInfo, arg2 *Nodes_ExpRefNode, arg3 *LnsList, arg4 *Nodes_BlockNode) {
    self.valInfo = arg1
    self.valExpRef = arg2
    self.valParamNameList = arg3
    self.block = arg4
}
func (self *Nodes_MatchCase) Get_valInfo(_env *LnsEnv) *Ast_AlgeValInfo{ return self.valInfo }
func (self *Nodes_MatchCase) Get_valExpRef(_env *LnsEnv) *Nodes_ExpRefNode{ return self.valExpRef }
func (self *Nodes_MatchCase) Get_valParamNameList(_env *LnsEnv) *LnsList{ return self.valParamNameList }
func (self *Nodes_MatchCase) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }

// declaration Class -- MatchNode
type Nodes_MatchNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_algeTypeInfo(_env *LnsEnv) *Ast_AlgeTypeInfo
    Get_caseKind(_env *LnsEnv) LnsInt
    Get_caseList(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_defaultBlock(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_failSafeDefault(_env *LnsEnv) bool
    Get_idInNS(_env *LnsEnv) LnsInt
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_val(_env *LnsEnv) *Nodes_Node
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_MatchNode struct {
    Nodes_Node
    idInNS LnsInt
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
func NewNodes_MatchNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsInt, arg8 *Nodes_Node, arg9 *Ast_AlgeTypeInfo, arg10 *LnsList, arg11 LnsAny, arg12 LnsInt, arg13 bool) *Nodes_MatchNode {
    obj := &Nodes_MatchNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_MatchNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
    return obj
}
func (self *Nodes_MatchNode) Get_idInNS(_env *LnsEnv) LnsInt{ return self.idInNS }
func (self *Nodes_MatchNode) Get_val(_env *LnsEnv) *Nodes_Node{ return self.val }
func (self *Nodes_MatchNode) Get_algeTypeInfo(_env *LnsEnv) *Ast_AlgeTypeInfo{ return self.algeTypeInfo }
func (self *Nodes_MatchNode) Get_caseList(_env *LnsEnv) *LnsList{ return self.caseList }
func (self *Nodes_MatchNode) Get_defaultBlock(_env *LnsEnv) LnsAny{ return self.defaultBlock }
func (self *Nodes_MatchNode) Get_caseKind(_env *LnsEnv) LnsInt{ return self.caseKind }
func (self *Nodes_MatchNode) Get_failSafeDefault(_env *LnsEnv) bool{ return self.failSafeDefault }
// 734: DeclConstr
func (self *Nodes_MatchNode) InitNodes_MatchNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,val *Nodes_Node,algeTypeInfo *Ast_AlgeTypeInfo,caseList *LnsList,defaultBlock LnsAny,caseKind LnsInt,failSafeDefault bool) {
    self.InitNodes_Node(_env, managerId, id, 71, pos, inTestBlock, macroArgFlag, typeList)
    self.idInNS = idInNS
    self.val = val
    self.algeTypeInfo = algeTypeInfo
    self.caseList = caseList
    self.defaultBlock = defaultBlock
    self.caseKind = caseKind
    self.failSafeDefault = failSafeDefault
}


// declaration Class -- LuneKindNode
type Nodes_LuneKindNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_exp(_env *LnsEnv) *Nodes_Node
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LuneKindNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node) *Nodes_LuneKindNode {
    obj := &Nodes_LuneKindNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LuneKindNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_LuneKindNode) Get_exp(_env *LnsEnv) *Nodes_Node{ return self.exp }
// 734: DeclConstr
func (self *Nodes_LuneKindNode) InitNodes_LuneKindNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 72, pos, inTestBlock, macroArgFlag, typeList)
    self.exp = exp
}


// declaration Class -- DeclMacroNode
type Nodes_DeclMacroNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_declInfo(_env *LnsEnv) *Nodes_DeclMacroInfo
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_DeclMacroNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_DeclMacroInfo) *Nodes_DeclMacroNode {
    obj := &Nodes_DeclMacroNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_DeclMacroNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_DeclMacroNode) Get_declInfo(_env *LnsEnv) *Nodes_DeclMacroInfo{ return self.declInfo }
// 734: DeclConstr
func (self *Nodes_DeclMacroNode) InitNodes_DeclMacroNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclMacroInfo) {
    self.InitNodes_Node(_env, managerId, id, 73, pos, inTestBlock, macroArgFlag, typeList)
    self.declInfo = declInfo
}


// declaration Class -- MacroEval
type Nodes_MacroEvalMtd interface {
    EvalFromCodeToLuaCode(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 string, arg3 *LnsList, arg4 LnsAny) string
    EvalToLuaCode(_env *LnsEnv, arg1 *Ast_ProcessInfo, arg2 *Nodes_DeclMacroNode) string
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
func (self *Nodes_MacroEval) InitNodes_MacroEval(_env *LnsEnv) {
}



// declaration Class -- TestCaseNode
type Nodes_TestCaseNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_block(_env *LnsEnv) *Nodes_BlockNode
    Get_commentList(_env *LnsEnv) LnsAny
    Get_ctrlName(_env *LnsEnv) string
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_impNode(_env *LnsEnv) *Nodes_Node
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_name(_env *LnsEnv) *Types_Token
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_TestCaseNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 *Nodes_Node, arg9 string, arg10 *Nodes_BlockNode) *Nodes_TestCaseNode {
    obj := &Nodes_TestCaseNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_TestCaseNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return obj
}
func (self *Nodes_TestCaseNode) Get_name(_env *LnsEnv) *Types_Token{ return self.name }
func (self *Nodes_TestCaseNode) Get_impNode(_env *LnsEnv) *Nodes_Node{ return self.impNode }
func (self *Nodes_TestCaseNode) Get_ctrlName(_env *LnsEnv) string{ return self.ctrlName }
func (self *Nodes_TestCaseNode) Get_block(_env *LnsEnv) *Nodes_BlockNode{ return self.block }
// 734: DeclConstr
func (self *Nodes_TestCaseNode) InitNodes_TestCaseNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,impNode *Nodes_Node,ctrlName string,block *Nodes_BlockNode) {
    self.InitNodes_Node(_env, managerId, id, 74, pos, inTestBlock, macroArgFlag, typeList)
    self.name = name
    self.impNode = impNode
    self.ctrlName = ctrlName
    self.block = block
}


// declaration Class -- TestBlockNode
type Nodes_TestBlockNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_stmtList(_env *LnsEnv) *LnsList
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    IsInnerPos(_env *LnsEnv, arg1 Types_Position) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_TestBlockNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *LnsList) *Nodes_TestBlockNode {
    obj := &Nodes_TestBlockNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_TestBlockNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_TestBlockNode) Get_stmtList(_env *LnsEnv) *LnsList{ return self.stmtList }
// 734: DeclConstr
func (self *Nodes_TestBlockNode) InitNodes_TestBlockNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 75, pos, inTestBlock, macroArgFlag, typeList)
    self.stmtList = stmtList
}


// declaration Class -- AbbrNode
type Nodes_AbbrNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_AbbrNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList) *Nodes_AbbrNode {
    obj := &Nodes_AbbrNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_AbbrNode(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
// 734: DeclConstr
func (self *Nodes_AbbrNode) InitNodes_AbbrNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 76, pos, inTestBlock, macroArgFlag, typeList)
}


// declaration Class -- BoxingNode
type Nodes_BoxingNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_src(_env *LnsEnv) *Nodes_Node
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_BoxingNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node) *Nodes_BoxingNode {
    obj := &Nodes_BoxingNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_BoxingNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_BoxingNode) Get_src(_env *LnsEnv) *Nodes_Node{ return self.src }
// 734: DeclConstr
func (self *Nodes_BoxingNode) InitNodes_BoxingNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,src *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 77, pos, inTestBlock, macroArgFlag, typeList)
    self.src = src
}


// declaration Class -- UnboxingNode
type Nodes_UnboxingNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_src(_env *LnsEnv) *Nodes_Node
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_UnboxingNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Nodes_Node) *Nodes_UnboxingNode {
    obj := &Nodes_UnboxingNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_UnboxingNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_UnboxingNode) Get_src(_env *LnsEnv) *Nodes_Node{ return self.src }
// 734: DeclConstr
func (self *Nodes_UnboxingNode) InitNodes_UnboxingNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,src *Nodes_Node) {
    self.InitNodes_Node(_env, managerId, id, 78, pos, inTestBlock, macroArgFlag, typeList)
    self.src = src
}


// declaration Class -- LiteralNilNode
type Nodes_LiteralNilNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralNilNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList) *Nodes_LiteralNilNode {
    obj := &Nodes_LiteralNilNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralNilNode(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
// 734: DeclConstr
func (self *Nodes_LiteralNilNode) InitNodes_LiteralNilNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 79, pos, inTestBlock, macroArgFlag, typeList)
}


// declaration Class -- LiteralCharNode
type Nodes_LiteralCharNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_num(_env *LnsEnv) LnsInt
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_token(_env *LnsEnv) *Types_Token
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralCharNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 LnsInt) *Nodes_LiteralCharNode {
    obj := &Nodes_LiteralCharNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralCharNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_LiteralCharNode) Get_token(_env *LnsEnv) *Types_Token{ return self.token }
func (self *Nodes_LiteralCharNode) Get_num(_env *LnsEnv) LnsInt{ return self.num }
// 734: DeclConstr
func (self *Nodes_LiteralCharNode) InitNodes_LiteralCharNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsInt) {
    self.InitNodes_Node(_env, managerId, id, 80, pos, inTestBlock, macroArgFlag, typeList)
    self.token = token
    self.num = num
}


// declaration Class -- LiteralIntNode
type Nodes_LiteralIntNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_num(_env *LnsEnv) LnsInt
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_token(_env *LnsEnv) *Types_Token
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralIntNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 LnsInt) *Nodes_LiteralIntNode {
    obj := &Nodes_LiteralIntNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralIntNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_LiteralIntNode) Get_token(_env *LnsEnv) *Types_Token{ return self.token }
func (self *Nodes_LiteralIntNode) Get_num(_env *LnsEnv) LnsInt{ return self.num }
// 734: DeclConstr
func (self *Nodes_LiteralIntNode) InitNodes_LiteralIntNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsInt) {
    self.InitNodes_Node(_env, managerId, id, 81, pos, inTestBlock, macroArgFlag, typeList)
    self.token = token
    self.num = num
}


// declaration Class -- LiteralRealNode
type Nodes_LiteralRealNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_num(_env *LnsEnv) LnsReal
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_token(_env *LnsEnv) *Types_Token
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralRealNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 LnsReal) *Nodes_LiteralRealNode {
    obj := &Nodes_LiteralRealNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralRealNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_LiteralRealNode) Get_token(_env *LnsEnv) *Types_Token{ return self.token }
func (self *Nodes_LiteralRealNode) Get_num(_env *LnsEnv) LnsReal{ return self.num }
// 734: DeclConstr
func (self *Nodes_LiteralRealNode) InitNodes_LiteralRealNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsReal) {
    self.InitNodes_Node(_env, managerId, id, 82, pos, inTestBlock, macroArgFlag, typeList)
    self.token = token
    self.num = num
}


// declaration Class -- LiteralArrayNode
type Nodes_LiteralArrayNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) LnsAny
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralArrayNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsAny) *Nodes_LiteralArrayNode {
    obj := &Nodes_LiteralArrayNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralArrayNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_LiteralArrayNode) Get_expList(_env *LnsEnv) LnsAny{ return self.expList }
// 734: DeclConstr
func (self *Nodes_LiteralArrayNode) InitNodes_LiteralArrayNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 83, pos, inTestBlock, macroArgFlag, typeList)
    self.expList = expList
}


// declaration Class -- LiteralListNode
type Nodes_LiteralListNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) LnsAny
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralListNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsAny) *Nodes_LiteralListNode {
    obj := &Nodes_LiteralListNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralListNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_LiteralListNode) Get_expList(_env *LnsEnv) LnsAny{ return self.expList }
// 734: DeclConstr
func (self *Nodes_LiteralListNode) InitNodes_LiteralListNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 84, pos, inTestBlock, macroArgFlag, typeList)
    self.expList = expList
}


// declaration Class -- LiteralSetNode
type Nodes_LiteralSetNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expList(_env *LnsEnv) LnsAny
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralSetNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 LnsAny) *Nodes_LiteralSetNode {
    obj := &Nodes_LiteralSetNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralSetNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_LiteralSetNode) Get_expList(_env *LnsEnv) LnsAny{ return self.expList }
// 734: DeclConstr
func (self *Nodes_LiteralSetNode) InitNodes_LiteralSetNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 85, pos, inTestBlock, macroArgFlag, typeList)
    self.expList = expList
}


// declaration Class -- PairItem
type Nodes_PairItemMtd interface {
    Get_key(_env *LnsEnv) *Nodes_Node
    Get_val(_env *LnsEnv) *Nodes_Node
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
func NewNodes_PairItem(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_Node) *Nodes_PairItem {
    obj := &Nodes_PairItem{}
    obj.FP = obj
    obj.InitNodes_PairItem(_env, arg1, arg2)
    return obj
}
func (self *Nodes_PairItem) InitNodes_PairItem(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_Node) {
    self.key = arg1
    self.val = arg2
}
func (self *Nodes_PairItem) Get_key(_env *LnsEnv) *Nodes_Node{ return self.key }
func (self *Nodes_PairItem) Get_val(_env *LnsEnv) *Nodes_Node{ return self.val }

// declaration Class -- LiteralMapNode
type Nodes_LiteralMapNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_map(_env *LnsEnv) *LnsMap
    Get_pairList(_env *LnsEnv) *LnsList
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralMapNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *LnsMap, arg8 *LnsList) *Nodes_LiteralMapNode {
    obj := &Nodes_LiteralMapNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralMapNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    return obj
}
func (self *Nodes_LiteralMapNode) Get_map(_env *LnsEnv) *LnsMap{ return self._map }
func (self *Nodes_LiteralMapNode) Get_pairList(_env *LnsEnv) *LnsList{ return self.pairList }
// 734: DeclConstr
func (self *Nodes_LiteralMapNode) InitNodes_LiteralMapNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,_map *LnsMap,pairList *LnsList) {
    self.InitNodes_Node(_env, managerId, id, 86, pos, inTestBlock, macroArgFlag, typeList)
    self._map = _map
    self.pairList = pairList
}


// declaration Class -- LiteralStringNode
type Nodes_LiteralStringNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_dddParam(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_orgParam(_env *LnsEnv) LnsAny
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_token(_env *LnsEnv) *Types_Token
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
}
type Nodes_LiteralStringNode struct {
    Nodes_Node
    token *Types_Token
    orgParam LnsAny
    dddParam LnsAny
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
func NewNodes_LiteralStringNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token, arg8 LnsAny, arg9 LnsAny) *Nodes_LiteralStringNode {
    obj := &Nodes_LiteralStringNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralStringNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
    return obj
}
func (self *Nodes_LiteralStringNode) Get_token(_env *LnsEnv) *Types_Token{ return self.token }
func (self *Nodes_LiteralStringNode) Get_orgParam(_env *LnsEnv) LnsAny{ return self.orgParam }
func (self *Nodes_LiteralStringNode) Get_dddParam(_env *LnsEnv) LnsAny{ return self.dddParam }
// 734: DeclConstr
func (self *Nodes_LiteralStringNode) InitNodes_LiteralStringNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token,orgParam LnsAny,dddParam LnsAny) {
    self.InitNodes_Node(_env, managerId, id, 87, pos, inTestBlock, macroArgFlag, typeList)
    self.token = token
    self.orgParam = orgParam
    self.dddParam = dddParam
}


// declaration Class -- LiteralBoolNode
type Nodes_LiteralBoolNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_token(_env *LnsEnv) *Types_Token
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralBoolNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token) *Nodes_LiteralBoolNode {
    obj := &Nodes_LiteralBoolNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralBoolNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_LiteralBoolNode) Get_token(_env *LnsEnv) *Types_Token{ return self.token }
// 734: DeclConstr
func (self *Nodes_LiteralBoolNode) InitNodes_LiteralBoolNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token) {
    self.InitNodes_Node(_env, managerId, id, 88, pos, inTestBlock, macroArgFlag, typeList)
    self.token = token
}


// declaration Class -- LiteralSymbolNode
type Nodes_LiteralSymbolNodeMtd interface {
    AddComment(_env *LnsEnv, arg1 *LnsList)
    AddTokenList(_env *LnsEnv, arg1 *LnsList, arg2 LnsInt, arg3 string)
    CanBeLeft(_env *LnsEnv) bool
    CanBeRight(_env *LnsEnv, arg1 *Ast_ProcessInfo) bool
    CanBeStatement(_env *LnsEnv) bool
    Comp(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetBreakKind(_env *LnsEnv, arg1 LnsInt) LnsInt
    GetIdTxt(_env *LnsEnv) string
    GetLiteral(_env *LnsEnv)(LnsAny, LnsAny)
    GetPrefix(_env *LnsEnv) LnsAny
    GetSymbolInfo(_env *LnsEnv) *LnsList
    Get_commentList(_env *LnsEnv) LnsAny
    Get_effectivePos(_env *LnsEnv) Types_Position
    Get_expType(_env *LnsEnv) *Ast_TypeInfo
    Get_expTypeList(_env *LnsEnv) *LnsList
    Get_inTestBlock(_env *LnsEnv) bool
    Get_isLValue(_env *LnsEnv) bool
    Get_kind(_env *LnsEnv) LnsInt
    Get_macroArgFlag(_env *LnsEnv) bool
    Get_pos(_env *LnsEnv) Types_Position
    Get_tailComment(_env *LnsEnv) LnsAny
    Get_token(_env *LnsEnv) *Types_Token
    HasNilAccess(_env *LnsEnv) bool
    ProcessFilter(_env *LnsEnv, arg1 *Nodes_Filter, arg2 LnsAny)
    SetLValue(_env *LnsEnv)
    Set_tailComment(_env *LnsEnv, arg1 LnsAny)
    SetupLiteralTokenList(_env *LnsEnv, arg1 *LnsList) bool
    Visit(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
    VisitSub(_env *LnsEnv, arg1 Nodes_NodeVisitor, arg2 LnsInt, arg3 *LnsSet) bool
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
func NewNodes_LiteralSymbolNode(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 Types_Position, arg4 bool, arg5 bool, arg6 *LnsList, arg7 *Types_Token) *Nodes_LiteralSymbolNode {
    obj := &Nodes_LiteralSymbolNode{}
    obj.FP = obj
    obj.Nodes_Node.FP = obj
    obj.InitNodes_LiteralSymbolNode(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Nodes_LiteralSymbolNode) Get_token(_env *LnsEnv) *Types_Token{ return self.token }
// 734: DeclConstr
func (self *Nodes_LiteralSymbolNode) InitNodes_LiteralSymbolNode(_env *LnsEnv, managerId LnsInt,id LnsInt,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token) {
    self.InitNodes_Node(_env, managerId, id, 89, pos, inTestBlock, macroArgFlag, typeList)
    self.token = token
}


// declaration Class -- DefMacroInfo
type Nodes_DefMacroInfoMtd interface {
    GetArgList(_env *LnsEnv) *LnsList
    GetTokenList(_env *LnsEnv) *LnsList
    Get_func(_env *LnsEnv) LnsAny
    Get_name(_env *LnsEnv) string
    Get_symbol2MacroValInfoMap(_env *LnsEnv) *LnsMap
    Set_func(_env *LnsEnv, arg1 LnsAny)
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
func NewNodes_DefMacroInfo(_env *LnsEnv, arg1 LnsAny, arg2 *Nodes_DeclMacroInfo, arg3 *LnsMap) *Nodes_DefMacroInfo {
    obj := &Nodes_DefMacroInfo{}
    obj.FP = obj
    obj.Nodes_MacroInfo.FP = obj
    obj.InitNodes_DefMacroInfo(_env, arg1, arg2, arg3)
    return obj
}
// 2900: DeclConstr
func (self *Nodes_DefMacroInfo) InitNodes_DefMacroInfo(_env *LnsEnv, _func LnsAny,declInfo *Nodes_DeclMacroInfo,symbol2MacroValInfoMap *LnsMap) {
    self.InitNodes_MacroInfo(_env, _func, symbol2MacroValInfoMap)
    self.DeclInfo = declInfo
    self.argList = NewLnsList([]LnsAny{})
    for _, _argNode := range( declInfo.FP.Get_argList(_env).Items ) {
        argNode := _argNode.(Nodes_DeclArgNodeDownCast).ToNodes_DeclArgNode()
        if argNode.FP.Get_kind(_env) == Nodes_NodeKind_get_DeclArg(_env){
            var argType *Ast_TypeInfo
            argType = argNode.FP.Get_expType(_env)
            var argName string
            argName = argNode.FP.Get_name(_env).Txt
            self.argList.Insert(Nodes_MacroArgInfo2Stem(NewNodes_MacroArgInfo(_env, argName, argType)))
        }
    }
}


// declaration Class -- ExportInfo
type Nodes_ExportInfoMtd interface {
    Assign(_env *LnsEnv, arg1 string) *FrontInterface_ExportInfo
    GetTypeInfo(_env *LnsEnv, arg1 LnsInt) LnsAny
    Get_assignName(_env *LnsEnv) string
    Get_fullName(_env *LnsEnv) string
    Get_globalSymbolList(_env *LnsEnv) *LnsList
    Get_importId2localTypeInfoMap(_env *LnsEnv) *LnsMap
    Get_importedAliasMap(_env *LnsEnv) *LnsMap
    Get_moduleId(_env *LnsEnv) *FrontInterface_ModuleId
    Get_modulePath(_env *LnsEnv) string
    Get_moduleTypeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_processInfo(_env *LnsEnv) *Ast_ProcessInfo
    Get_provideInfo(_env *LnsEnv) *FrontInterface_ModuleProvideInfo
    Get_streamName(_env *LnsEnv) string
    Get_typeId2DefMacroInfo(_env *LnsEnv) *LnsMap
    Set_importId2localTypeInfoMap(_env *LnsEnv, arg1 *LnsMap)
}
type Nodes_ExportInfo struct {
    FrontInterface_ExportInfo
    typeId2DefMacroInfo *LnsMap
    FP Nodes_ExportInfoMtd
}
func Nodes_ExportInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Nodes_ExportInfo).FP
}
type Nodes_ExportInfoDownCast interface {
    ToNodes_ExportInfo() *Nodes_ExportInfo
}
func Nodes_ExportInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Nodes_ExportInfoDownCast)
    if ok { return work.ToNodes_ExportInfo() }
    return nil
}
func (obj *Nodes_ExportInfo) ToNodes_ExportInfo() *Nodes_ExportInfo {
    return obj
}
func NewNodes_ExportInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *FrontInterface_ModuleProvideInfo, arg3 *Ast_ProcessInfo, arg4 *LnsList, arg5 *LnsMap, arg6 *FrontInterface_ModuleId, arg7 string, arg8 string, arg9 string, arg10 *LnsMap, arg11 *LnsMap) *Nodes_ExportInfo {
    obj := &Nodes_ExportInfo{}
    obj.FP = obj
    obj.FrontInterface_ExportInfo.FP = obj
    obj.InitNodes_ExportInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
    return obj
}
func (self *Nodes_ExportInfo) InitNodes_ExportInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *FrontInterface_ModuleProvideInfo, arg3 *Ast_ProcessInfo, arg4 *LnsList, arg5 *LnsMap, arg6 *FrontInterface_ModuleId, arg7 string, arg8 string, arg9 string, arg10 *LnsMap, arg11 *LnsMap) {
    self.FrontInterface_ExportInfo.InitFrontInterface_ExportInfo( _env, arg1,arg2,arg3,arg4,arg5,arg6,arg7,arg8,arg9,arg10)
    self.typeId2DefMacroInfo = arg11
}
func (self *Nodes_ExportInfo) Get_typeId2DefMacroInfo(_env *LnsEnv) *LnsMap{ return self.typeId2DefMacroInfo }

func Lns_Nodes_init(_env *LnsEnv) {
    if init_Nodes { return }
    init_Nodes = true
    Nodes__mod__ = "@lune.@base.@Nodes"
    Lns_InitMod()
    Lns_Parser_init(_env)
    Lns_Util_init(_env)
    Lns_frontInterface_init(_env)
    Lns_Ast_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Types_init(_env)
    Nodes_nodeKind2NameMapWork = NewLnsMap( map[LnsAny]LnsAny{})
    Nodes_nodeKind2NameMap = Nodes_nodeKind2NameMapWork
    Nodes_nodeKindSeed = 0
    Nodes_regKind_17_(_env, "None")
    
    Nodes_regKind_17_(_env, "Shebang")
    
    Nodes_regKind_17_(_env, "ConvStat")
    
    Nodes_regKind_17_(_env, "BlankLine")
    
    Nodes_regKind_17_(_env, "Subfile")
    
    Nodes_regKind_17_(_env, "Import")
    
    Nodes_regKind_17_(_env, "Root")
    
    Nodes_regKind_17_(_env, "RefType")
    
    Nodes_regKind_17_(_env, "Block")
    
    Nodes_regKind_17_(_env, "Scope")
    
    Nodes_regKind_17_(_env, "If")
    
    Nodes_regKind_17_(_env, "ExpList")
    
    Nodes_regKind_17_(_env, "Switch")
    
    Nodes_regKind_17_(_env, "While")
    
    Nodes_regKind_17_(_env, "Repeat")
    
    
    Nodes_regKind_17_(_env, "For")
    
    
    Nodes_regKind_17_(_env, "Apply")
    
    
    Nodes_regKind_17_(_env, "Foreach")
    
    
    Nodes_regKind_17_(_env, "Forsort")
    
    
    Nodes_regKind_17_(_env, "Return")
    
    Nodes_regKind_17_(_env, "Break")
    
    Nodes_regKind_17_(_env, "Provide")
    
    Nodes_regKind_17_(_env, "ExpNew")
    
    Nodes_regKind_17_(_env, "ExpUnwrap")
    
    Nodes_regKind_17_(_env, "ExpRef")
    
    Nodes_regKind_17_(_env, "ExpSetVal")
    
    Nodes_regKind_17_(_env, "ExpSetItem")
    
    Nodes_regKind_17_(_env, "ExpOp2")
    
    Nodes_regKind_17_(_env, "UnwrapSet")
    
    Nodes_regKind_17_(_env, "IfUnwrap")
    
    Nodes_regKind_17_(_env, "When")
    
    Nodes_regKind_17_(_env, "ExpCast")
    
    Nodes_regKind_17_(_env, "ExpToDDD")
    
    Nodes_regKind_17_(_env, "ExpSubDDD")
    
    Nodes_regKind_17_(_env, "ExpOp1")
    
    Nodes_regKind_17_(_env, "ExpRefItem")
    
    Nodes_regKind_17_(_env, "ExpCall")
    
    Nodes_regKind_17_(_env, "ExpMRet")
    
    Nodes_regKind_17_(_env, "ExpAccessMRet")
    
    Nodes_regKind_17_(_env, "ExpMultiTo1")
    
    Nodes_regKind_17_(_env, "ExpParen")
    
    Nodes_regKind_17_(_env, "ExpMacroExp")
    
    Nodes_regKind_17_(_env, "ExpMacroStat")
    
    Nodes_regKind_17_(_env, "ExpMacroArgExp")
    
    Nodes_regKind_17_(_env, "StmtExp")
    
    Nodes_regKind_17_(_env, "ExpMacroStatList")
    
    Nodes_regKind_17_(_env, "ExpOmitEnum")
    
    Nodes_regKind_17_(_env, "RefField")
    
    Nodes_regKind_17_(_env, "GetField")
    
    Nodes_regKind_17_(_env, "Alias")
    
    Nodes_regKind_17_(_env, "DeclVar")
    
    Nodes_regKind_17_(_env, "DeclForm")
    
    Nodes_regKind_17_(_env, "DeclFunc")
    
    Nodes_regKind_17_(_env, "DeclMethod")
    
    Nodes_regKind_17_(_env, "ProtoMethod")
    
    Nodes_regKind_17_(_env, "DeclConstr")
    
    Nodes_regKind_17_(_env, "DeclDestr")
    
    Nodes_regKind_17_(_env, "ExpCallSuperCtor")
    
    Nodes_regKind_17_(_env, "ExpCallSuper")
    
    Nodes_regKind_17_(_env, "AsyncLock")
    
    Nodes_regKind_17_(_env, "Request")
    
    Nodes_regKind_17_(_env, "DeclMember")
    
    Nodes_regKind_17_(_env, "DeclArg")
    
    Nodes_regKind_17_(_env, "DeclArgDDD")
    
    Nodes_regKind_17_(_env, "DeclAdvertise")
    
    Nodes_regKind_17_(_env, "ProtoClass")
    
    Nodes_regKind_17_(_env, "DeclClass")
    
    Nodes_regKind_17_(_env, "DeclEnum")
    
    Nodes_regKind_17_(_env, "DeclAlge")
    
    Nodes_regKind_17_(_env, "NewAlgeVal")
    
    Nodes_regKind_17_(_env, "LuneControl")
    
    Nodes_regKind_17_(_env, "Match")
    
    Nodes_regKind_17_(_env, "LuneKind")
    
    Nodes_regKind_17_(_env, "DeclMacro")
    
    Nodes_regKind_17_(_env, "TestCase")
    
    Nodes_regKind_17_(_env, "TestBlock")
    
    Nodes_regKind_17_(_env, "Abbr")
    
    Nodes_regKind_17_(_env, "Boxing")
    
    Nodes_regKind_17_(_env, "Unboxing")
    
    Nodes_regKind_17_(_env, "LiteralNil")
    
    Nodes_regKind_17_(_env, "LiteralChar")
    
    Nodes_regKind_17_(_env, "LiteralInt")
    
    Nodes_regKind_17_(_env, "LiteralReal")
    
    Nodes_regKind_17_(_env, "LiteralArray")
    
    Nodes_regKind_17_(_env, "LiteralList")
    
    Nodes_regKind_17_(_env, "LiteralSet")
    
    Nodes_regKind_17_(_env, "LiteralMap")
    
    Nodes_regKind_17_(_env, "LiteralString")
    
    Nodes_regKind_17_(_env, "LiteralBool")
    
    Nodes_regKind_17_(_env, "LiteralSymbol")
    
    
}
func init() {
    init_Nodes = false
}
// 50: decl @lune.@base.@Nodes.SimpleModuleInfoManager.push
func (self *Nodes_SimpleModuleInfoManager) Push(_env *LnsEnv, moduleInfoManager Ast_ModuleInfoManager) {
    self.moduleInfoManagerHist.Insert(self.ModuleInfoManager)
    self.ModuleInfoManager = moduleInfoManager
}
// 55: decl @lune.@base.@Nodes.SimpleModuleInfoManager.pop
func (self *Nodes_SimpleModuleInfoManager) Pop(_env *LnsEnv) {
    self.ModuleInfoManager = self.moduleInfoManagerHist.GetAt(self.moduleInfoManagerHist.Len()).(Ast_ModuleInfoManager)
    self.moduleInfoManagerHist.Remove(nil)
}
// 85: decl @lune.@base.@Nodes.Filter.pushOpt
func (self *Nodes_Filter) pushOpt(_env *LnsEnv, _opt LnsAny) {
    opt := _opt
    self.optStack.Insert(opt)
}
// 88: decl @lune.@base.@Nodes.Filter.popOpt
func (self *Nodes_Filter) popOpt(_env *LnsEnv, _opt LnsAny) {
    self.optStack.Remove(nil)
}
// 92: decl @lune.@base.@Nodes.Filter.get_moduleInfoManager
func (self *Nodes_Filter) Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager {
    return self.moduleInfoManager.FP
}
// 96: decl @lune.@base.@Nodes.Filter.getFull
func (self *Nodes_Filter) GetFull(_env *LnsEnv, typeInfo *Ast_TypeInfo,localFlag bool) string {
    return typeInfo.FP.GetFullName(_env, self.typeNameCtrl, self.moduleInfoManager.FP, localFlag)
}
// 408: decl @lune.@base.@Nodes.Filter.defaultProcess
func (self *Nodes_Filter) DefaultProcess(_env *LnsEnv, node *Nodes_Node,_opt LnsAny) {
    if self.errorOnDefault{
        Util_err(_env, _env.GetVM().String_format("not implement yet -- %s", []LnsAny{Nodes_getNodeKindName(_env, node.FP.Get_kind(_env))}))
    }
}
// 1: decl @lune.@base.@Nodes.Filter.processNone
func (self *Nodes_Filter) ProcessNone(_env *LnsEnv, node *Nodes_NoneNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processShebang
func (self *Nodes_Filter) ProcessShebang(_env *LnsEnv, node *Nodes_ShebangNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processConvStat
func (self *Nodes_Filter) ProcessConvStat(_env *LnsEnv, node *Nodes_ConvStatNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processBlankLine
func (self *Nodes_Filter) ProcessBlankLine(_env *LnsEnv, node *Nodes_BlankLineNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processSubfile
func (self *Nodes_Filter) ProcessSubfile(_env *LnsEnv, node *Nodes_SubfileNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processImport
func (self *Nodes_Filter) ProcessImport(_env *LnsEnv, node *Nodes_ImportNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processRoot
func (self *Nodes_Filter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processRefType
func (self *Nodes_Filter) ProcessRefType(_env *LnsEnv, node *Nodes_RefTypeNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processScope
func (self *Nodes_Filter) ProcessScope(_env *LnsEnv, node *Nodes_ScopeNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processIf
func (self *Nodes_Filter) ProcessIf(_env *LnsEnv, node *Nodes_IfNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpList
func (self *Nodes_Filter) ProcessExpList(_env *LnsEnv, node *Nodes_ExpListNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processSwitch
func (self *Nodes_Filter) ProcessSwitch(_env *LnsEnv, node *Nodes_SwitchNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processWhile
func (self *Nodes_Filter) ProcessWhile(_env *LnsEnv, node *Nodes_WhileNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processRepeat
func (self *Nodes_Filter) ProcessRepeat(_env *LnsEnv, node *Nodes_RepeatNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processFor
func (self *Nodes_Filter) ProcessFor(_env *LnsEnv, node *Nodes_ForNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processApply
func (self *Nodes_Filter) ProcessApply(_env *LnsEnv, node *Nodes_ApplyNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processForeach
func (self *Nodes_Filter) ProcessForeach(_env *LnsEnv, node *Nodes_ForeachNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processForsort
func (self *Nodes_Filter) ProcessForsort(_env *LnsEnv, node *Nodes_ForsortNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processReturn
func (self *Nodes_Filter) ProcessReturn(_env *LnsEnv, node *Nodes_ReturnNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processBreak
func (self *Nodes_Filter) ProcessBreak(_env *LnsEnv, node *Nodes_BreakNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processProvide
func (self *Nodes_Filter) ProcessProvide(_env *LnsEnv, node *Nodes_ProvideNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpNew
func (self *Nodes_Filter) ProcessExpNew(_env *LnsEnv, node *Nodes_ExpNewNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpUnwrap
func (self *Nodes_Filter) ProcessExpUnwrap(_env *LnsEnv, node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpRef
func (self *Nodes_Filter) ProcessExpRef(_env *LnsEnv, node *Nodes_ExpRefNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpSetVal
func (self *Nodes_Filter) ProcessExpSetVal(_env *LnsEnv, node *Nodes_ExpSetValNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpSetItem
func (self *Nodes_Filter) ProcessExpSetItem(_env *LnsEnv, node *Nodes_ExpSetItemNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpOp2
func (self *Nodes_Filter) ProcessExpOp2(_env *LnsEnv, node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processUnwrapSet
func (self *Nodes_Filter) ProcessUnwrapSet(_env *LnsEnv, node *Nodes_UnwrapSetNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processIfUnwrap
func (self *Nodes_Filter) ProcessIfUnwrap(_env *LnsEnv, node *Nodes_IfUnwrapNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processWhen
func (self *Nodes_Filter) ProcessWhen(_env *LnsEnv, node *Nodes_WhenNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpCast
func (self *Nodes_Filter) ProcessExpCast(_env *LnsEnv, node *Nodes_ExpCastNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpToDDD
func (self *Nodes_Filter) ProcessExpToDDD(_env *LnsEnv, node *Nodes_ExpToDDDNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpSubDDD
func (self *Nodes_Filter) ProcessExpSubDDD(_env *LnsEnv, node *Nodes_ExpSubDDDNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpOp1
func (self *Nodes_Filter) ProcessExpOp1(_env *LnsEnv, node *Nodes_ExpOp1Node,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpRefItem
func (self *Nodes_Filter) ProcessExpRefItem(_env *LnsEnv, node *Nodes_ExpRefItemNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpCall
func (self *Nodes_Filter) ProcessExpCall(_env *LnsEnv, node *Nodes_ExpCallNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpMRet
func (self *Nodes_Filter) ProcessExpMRet(_env *LnsEnv, node *Nodes_ExpMRetNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpAccessMRet
func (self *Nodes_Filter) ProcessExpAccessMRet(_env *LnsEnv, node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpMultiTo1
func (self *Nodes_Filter) ProcessExpMultiTo1(_env *LnsEnv, node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpParen
func (self *Nodes_Filter) ProcessExpParen(_env *LnsEnv, node *Nodes_ExpParenNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpMacroExp
func (self *Nodes_Filter) ProcessExpMacroExp(_env *LnsEnv, node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpMacroStat
func (self *Nodes_Filter) ProcessExpMacroStat(_env *LnsEnv, node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpMacroArgExp
func (self *Nodes_Filter) ProcessExpMacroArgExp(_env *LnsEnv, node *Nodes_ExpMacroArgExpNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processStmtExp
func (self *Nodes_Filter) ProcessStmtExp(_env *LnsEnv, node *Nodes_StmtExpNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpMacroStatList
func (self *Nodes_Filter) ProcessExpMacroStatList(_env *LnsEnv, node *Nodes_ExpMacroStatListNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpOmitEnum
func (self *Nodes_Filter) ProcessExpOmitEnum(_env *LnsEnv, node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processRefField
func (self *Nodes_Filter) ProcessRefField(_env *LnsEnv, node *Nodes_RefFieldNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processGetField
func (self *Nodes_Filter) ProcessGetField(_env *LnsEnv, node *Nodes_GetFieldNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processAlias
func (self *Nodes_Filter) ProcessAlias(_env *LnsEnv, node *Nodes_AliasNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclVar
func (self *Nodes_Filter) ProcessDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclForm
func (self *Nodes_Filter) ProcessDeclForm(_env *LnsEnv, node *Nodes_DeclFormNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclFunc
func (self *Nodes_Filter) ProcessDeclFunc(_env *LnsEnv, node *Nodes_DeclFuncNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclMethod
func (self *Nodes_Filter) ProcessDeclMethod(_env *LnsEnv, node *Nodes_DeclMethodNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processProtoMethod
func (self *Nodes_Filter) ProcessProtoMethod(_env *LnsEnv, node *Nodes_ProtoMethodNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclConstr
func (self *Nodes_Filter) ProcessDeclConstr(_env *LnsEnv, node *Nodes_DeclConstrNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclDestr
func (self *Nodes_Filter) ProcessDeclDestr(_env *LnsEnv, node *Nodes_DeclDestrNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpCallSuperCtor
func (self *Nodes_Filter) ProcessExpCallSuperCtor(_env *LnsEnv, node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processExpCallSuper
func (self *Nodes_Filter) ProcessExpCallSuper(_env *LnsEnv, node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processAsyncLock
func (self *Nodes_Filter) ProcessAsyncLock(_env *LnsEnv, node *Nodes_AsyncLockNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processRequest
func (self *Nodes_Filter) ProcessRequest(_env *LnsEnv, node *Nodes_RequestNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclMember
func (self *Nodes_Filter) ProcessDeclMember(_env *LnsEnv, node *Nodes_DeclMemberNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclArg
func (self *Nodes_Filter) ProcessDeclArg(_env *LnsEnv, node *Nodes_DeclArgNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclArgDDD
func (self *Nodes_Filter) ProcessDeclArgDDD(_env *LnsEnv, node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclAdvertise
func (self *Nodes_Filter) ProcessDeclAdvertise(_env *LnsEnv, node *Nodes_DeclAdvertiseNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processProtoClass
func (self *Nodes_Filter) ProcessProtoClass(_env *LnsEnv, node *Nodes_ProtoClassNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclClass
func (self *Nodes_Filter) ProcessDeclClass(_env *LnsEnv, node *Nodes_DeclClassNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclEnum
func (self *Nodes_Filter) ProcessDeclEnum(_env *LnsEnv, node *Nodes_DeclEnumNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclAlge
func (self *Nodes_Filter) ProcessDeclAlge(_env *LnsEnv, node *Nodes_DeclAlgeNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processNewAlgeVal
func (self *Nodes_Filter) ProcessNewAlgeVal(_env *LnsEnv, node *Nodes_NewAlgeValNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLuneControl
func (self *Nodes_Filter) ProcessLuneControl(_env *LnsEnv, node *Nodes_LuneControlNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processMatch
func (self *Nodes_Filter) ProcessMatch(_env *LnsEnv, node *Nodes_MatchNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLuneKind
func (self *Nodes_Filter) ProcessLuneKind(_env *LnsEnv, node *Nodes_LuneKindNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processDeclMacro
func (self *Nodes_Filter) ProcessDeclMacro(_env *LnsEnv, node *Nodes_DeclMacroNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processTestCase
func (self *Nodes_Filter) ProcessTestCase(_env *LnsEnv, node *Nodes_TestCaseNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processTestBlock
func (self *Nodes_Filter) ProcessTestBlock(_env *LnsEnv, node *Nodes_TestBlockNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processAbbr
func (self *Nodes_Filter) ProcessAbbr(_env *LnsEnv, node *Nodes_AbbrNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processBoxing
func (self *Nodes_Filter) ProcessBoxing(_env *LnsEnv, node *Nodes_BoxingNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processUnboxing
func (self *Nodes_Filter) ProcessUnboxing(_env *LnsEnv, node *Nodes_UnboxingNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralNil
func (self *Nodes_Filter) ProcessLiteralNil(_env *LnsEnv, node *Nodes_LiteralNilNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralChar
func (self *Nodes_Filter) ProcessLiteralChar(_env *LnsEnv, node *Nodes_LiteralCharNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralInt
func (self *Nodes_Filter) ProcessLiteralInt(_env *LnsEnv, node *Nodes_LiteralIntNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralReal
func (self *Nodes_Filter) ProcessLiteralReal(_env *LnsEnv, node *Nodes_LiteralRealNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralArray
func (self *Nodes_Filter) ProcessLiteralArray(_env *LnsEnv, node *Nodes_LiteralArrayNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralList
func (self *Nodes_Filter) ProcessLiteralList(_env *LnsEnv, node *Nodes_LiteralListNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralSet
func (self *Nodes_Filter) ProcessLiteralSet(_env *LnsEnv, node *Nodes_LiteralSetNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralMap
func (self *Nodes_Filter) ProcessLiteralMap(_env *LnsEnv, node *Nodes_LiteralMapNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralString
func (self *Nodes_Filter) ProcessLiteralString(_env *LnsEnv, node *Nodes_LiteralStringNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralBool
func (self *Nodes_Filter) ProcessLiteralBool(_env *LnsEnv, node *Nodes_LiteralBoolNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 1: decl @lune.@base.@Nodes.Filter.processLiteralSymbol
func (self *Nodes_Filter) ProcessLiteralSymbol(_env *LnsEnv, node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    opt := _opt
    self.FP.pushOpt(_env, opt)
    self.FP.DefaultProcess(_env, &node.Nodes_Node, opt)
    self.FP.popOpt(_env, opt)
}
// 2946: decl @lune.@base.@Nodes.Filter.processBlockSub
func (self *Nodes_Filter) ProcessBlockSub(_env *LnsEnv, node *Nodes_BlockNode,_opt LnsAny) {
}
// 2949: decl @lune.@base.@Nodes.Filter.processBlock
func (self *Nodes_Filter) ProcessBlock(_env *LnsEnv, node *Nodes_BlockNode,_opt LnsAny) {
    opt := _opt
    self.moduleInfoManager.FP.Push(_env, node.FP.Get_scope(_env).FP)
    self.FP.ProcessBlockSub(_env, node, opt)
    self.moduleInfoManager.FP.Pop(_env)
}
// 239: decl @lune.@base.@Nodes.Node.comp
func (self *Nodes_Node) Comp(_env *LnsEnv, node *Nodes_Node) LnsInt {
    if self.managerId < node.managerId{
        return -1
    }
    if self.managerId > node.managerId{
        return 1
    }
    if self.id < node.id{
        return -1
    }
    if self.id > node.id{
        return 1
    }
    return 0
}
// 255: decl @lune.@base.@Nodes.Node.getIdTxt
func (self *Nodes_Node) GetIdTxt(_env *LnsEnv) string {
    return _env.GetVM().String_format("%d_%d", []LnsAny{self.managerId, self.id})
}
// 259: decl @lune.@base.@Nodes.Node.get_effectivePos
func (self *Nodes_Node) Get_effectivePos(_env *LnsEnv) Types_Position {
    return self.pos
}
// 263: decl @lune.@base.@Nodes.Node.setLValue
func (self *Nodes_Node) SetLValue(_env *LnsEnv) {
    self.IsLValue = true
}
// 267: decl @lune.@base.@Nodes.Node.getPrefix
func (self *Nodes_Node) GetPrefix(_env *LnsEnv) LnsAny {
    return nil
}
// 287: decl @lune.@base.@Nodes.Node.addComment
func (self *Nodes_Node) AddComment(_env *LnsEnv, commentList *LnsList) {
    if commentList.Len() != 0{
        var workList *LnsList
        {
            __exp := self.commentList
            if !Lns_IsNil( __exp ) {
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
// 303: decl @lune.@base.@Nodes.Node.get_expType
func (self *Nodes_Node) Get_expType(_env *LnsEnv) *Ast_TypeInfo {
    if self.expTypeList.Len() == 0{
        return Ast_builtinTypeNone
    }
    return self.expTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
}
// 310: decl @lune.@base.@Nodes.Node.addTokenList
func (self *Nodes_Node) AddTokenList(_env *LnsEnv, list *LnsList,kind LnsInt,txt string) {
    list.Insert(Types_Token2Stem(NewTypes_Token(_env, kind, txt, self.pos, false, nil)))
}
// 314: decl @lune.@base.@Nodes.Node.setupLiteralTokenList
func (self *Nodes_Node) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    return false
}
// 318: decl @lune.@base.@Nodes.Node.getLiteral
func (self *Nodes_Node) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    return nil, nil
}
// 321: decl @lune.@base.@Nodes.Node.processFilter
func (self *Nodes_Node) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
}
// 324: decl @lune.@base.@Nodes.Node.canBeLeft
func (self *Nodes_Node) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 328: decl @lune.@base.@Nodes.Node.canBeRight
func (self *Nodes_Node) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 332: decl @lune.@base.@Nodes.Node.canBeStatement
func (self *Nodes_Node) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 342: decl @lune.@base.@Nodes.Node.getBreakKind
func (self *Nodes_Node) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    return Nodes_BreakKind__None
}
// 349: decl @lune.@base.@Nodes.Node.hasNilAccess
func (self *Nodes_Node) HasNilAccess(_env *LnsEnv) bool {
    return false
}
// 360: decl @lune.@base.@Nodes.Node.visitSub
func (self *Nodes_Node) VisitSub(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return true
}
// 2259: decl @lune.@base.@Nodes.Node.getSymbolInfo
func (self *Nodes_Node) GetSymbolInfo(_env *LnsEnv) *LnsList {
    return Node_getSymbolInfo__processExpNode_0_(_env, self)
}
// 480: decl @lune.@base.@Nodes.NodeManager.nextId
func (self *Nodes_NodeManager) NextId(_env *LnsEnv) LnsInt {
    self.idSeed = self.idSeed + 1
    return self.idSeed
}
// 485: decl @lune.@base.@Nodes.NodeManager.getList
func (self *Nodes_NodeManager) GetList(_env *LnsEnv, kind LnsInt) *LnsList {
    var list *LnsList
    
    {
        _list := self.nodeKind2NodeList.Get(kind)
        if _list == nil{
            return NewLnsList([]LnsAny{})
        } else {
            list = _list.(*LnsList)
        }
    }
    return list
}
// 491: decl @lune.@base.@Nodes.NodeManager.addNode
func (self *Nodes_NodeManager) AddNode(_env *LnsEnv, node *Nodes_Node) {
    var list *LnsList
    
    {
        _list := self.nodeKind2NodeList.Get(node.FP.Get_kind(_env))
        if _list == nil{
            list = NewLnsList([]LnsAny{})
            self.nodeKind2NodeList.Set(node.FP.Get_kind(_env),list)
        } else {
            list = _list.(*LnsList)
        }
    }
    list.Insert(Nodes_Node2Stem(node))
}
// 498: decl @lune.@base.@Nodes.NodeManager.delNode
func (self *Nodes_NodeManager) DelNode(_env *LnsEnv, node *Nodes_Node) {
    var list *LnsList
    
    {
        _list := self.nodeKind2NodeList.Get(node.FP.Get_kind(_env))
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
// 514: decl @lune.@base.@Nodes.NodeManager.addFrom
func (self *Nodes_NodeManager) AddFrom(_env *LnsEnv, nodeManager *Nodes_NodeManager) {
    for _nodeKind, _list := range( nodeManager.nodeKind2NodeList.Items ) {
        nodeKind := _nodeKind.(LnsInt)
        list := _list.(*LnsList)
        var dstList *LnsList
        
        {
            _dstList := self.nodeKind2NodeList.Get(nodeKind)
            if _dstList == nil{
                dstList = NewLnsList([]LnsAny{})
                self.nodeKind2NodeList.Set(nodeKind,dstList)
            } else {
                dstList = _dstList.(*LnsList)
            }
        }
        for _, _node := range( list.Items ) {
            node := _node.(Nodes_NodeDownCast).ToNodes_Node()
            dstList.Insert(Nodes_Node2Stem(node))
        }
    }
}
// 1: decl @lune.@base.@Nodes.NodeManager.getNoneNodeList
func (self *Nodes_NodeManager) GetNoneNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 0))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getShebangNodeList
func (self *Nodes_NodeManager) GetShebangNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 1))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getConvStatNodeList
func (self *Nodes_NodeManager) GetConvStatNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 2))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getBlankLineNodeList
func (self *Nodes_NodeManager) GetBlankLineNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 3))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getSubfileNodeList
func (self *Nodes_NodeManager) GetSubfileNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 4))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getImportNodeList
func (self *Nodes_NodeManager) GetImportNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 5))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getRootNodeList
func (self *Nodes_NodeManager) GetRootNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 6))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getRefTypeNodeList
func (self *Nodes_NodeManager) GetRefTypeNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 7))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getBlockNodeList
func (self *Nodes_NodeManager) GetBlockNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 8))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getScopeNodeList
func (self *Nodes_NodeManager) GetScopeNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 9))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getIfNodeList
func (self *Nodes_NodeManager) GetIfNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 10))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpListNodeList
func (self *Nodes_NodeManager) GetExpListNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 11))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getSwitchNodeList
func (self *Nodes_NodeManager) GetSwitchNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 12))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getWhileNodeList
func (self *Nodes_NodeManager) GetWhileNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 13))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getRepeatNodeList
func (self *Nodes_NodeManager) GetRepeatNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 14))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getForNodeList
func (self *Nodes_NodeManager) GetForNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 15))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getApplyNodeList
func (self *Nodes_NodeManager) GetApplyNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 16))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getForeachNodeList
func (self *Nodes_NodeManager) GetForeachNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 17))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getForsortNodeList
func (self *Nodes_NodeManager) GetForsortNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 18))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getReturnNodeList
func (self *Nodes_NodeManager) GetReturnNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 19))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getBreakNodeList
func (self *Nodes_NodeManager) GetBreakNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 20))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getProvideNodeList
func (self *Nodes_NodeManager) GetProvideNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 21))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpNewNodeList
func (self *Nodes_NodeManager) GetExpNewNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 22))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpUnwrapNodeList
func (self *Nodes_NodeManager) GetExpUnwrapNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 23))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpRefNodeList
func (self *Nodes_NodeManager) GetExpRefNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 24))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpSetValNodeList
func (self *Nodes_NodeManager) GetExpSetValNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 25))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpSetItemNodeList
func (self *Nodes_NodeManager) GetExpSetItemNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 26))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpOp2NodeList
func (self *Nodes_NodeManager) GetExpOp2NodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 27))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getUnwrapSetNodeList
func (self *Nodes_NodeManager) GetUnwrapSetNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 28))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getIfUnwrapNodeList
func (self *Nodes_NodeManager) GetIfUnwrapNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 29))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getWhenNodeList
func (self *Nodes_NodeManager) GetWhenNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 30))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpCastNodeList
func (self *Nodes_NodeManager) GetExpCastNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 31))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpToDDDNodeList
func (self *Nodes_NodeManager) GetExpToDDDNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 32))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpSubDDDNodeList
func (self *Nodes_NodeManager) GetExpSubDDDNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 33))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpOp1NodeList
func (self *Nodes_NodeManager) GetExpOp1NodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 34))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpRefItemNodeList
func (self *Nodes_NodeManager) GetExpRefItemNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 35))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpCallNodeList
func (self *Nodes_NodeManager) GetExpCallNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 36))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpMRetNodeList
func (self *Nodes_NodeManager) GetExpMRetNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 37))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpAccessMRetNodeList
func (self *Nodes_NodeManager) GetExpAccessMRetNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 38))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpMultiTo1NodeList
func (self *Nodes_NodeManager) GetExpMultiTo1NodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 39))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpParenNodeList
func (self *Nodes_NodeManager) GetExpParenNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 40))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpMacroExpNodeList
func (self *Nodes_NodeManager) GetExpMacroExpNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 41))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpMacroStatNodeList
func (self *Nodes_NodeManager) GetExpMacroStatNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 42))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpMacroArgExpNodeList
func (self *Nodes_NodeManager) GetExpMacroArgExpNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 43))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getStmtExpNodeList
func (self *Nodes_NodeManager) GetStmtExpNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 44))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpMacroStatListNodeList
func (self *Nodes_NodeManager) GetExpMacroStatListNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 45))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpOmitEnumNodeList
func (self *Nodes_NodeManager) GetExpOmitEnumNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 46))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getRefFieldNodeList
func (self *Nodes_NodeManager) GetRefFieldNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 47))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getGetFieldNodeList
func (self *Nodes_NodeManager) GetGetFieldNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 48))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getAliasNodeList
func (self *Nodes_NodeManager) GetAliasNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 49))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclVarNodeList
func (self *Nodes_NodeManager) GetDeclVarNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 50))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclFormNodeList
func (self *Nodes_NodeManager) GetDeclFormNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 51))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclFuncNodeList
func (self *Nodes_NodeManager) GetDeclFuncNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 52))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclMethodNodeList
func (self *Nodes_NodeManager) GetDeclMethodNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 53))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getProtoMethodNodeList
func (self *Nodes_NodeManager) GetProtoMethodNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 54))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclConstrNodeList
func (self *Nodes_NodeManager) GetDeclConstrNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 55))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclDestrNodeList
func (self *Nodes_NodeManager) GetDeclDestrNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 56))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpCallSuperCtorNodeList
func (self *Nodes_NodeManager) GetExpCallSuperCtorNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 57))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getExpCallSuperNodeList
func (self *Nodes_NodeManager) GetExpCallSuperNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 58))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getAsyncLockNodeList
func (self *Nodes_NodeManager) GetAsyncLockNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 59))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getRequestNodeList
func (self *Nodes_NodeManager) GetRequestNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 60))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclMemberNodeList
func (self *Nodes_NodeManager) GetDeclMemberNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 61))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclArgNodeList
func (self *Nodes_NodeManager) GetDeclArgNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 62))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclArgDDDNodeList
func (self *Nodes_NodeManager) GetDeclArgDDDNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 63))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclAdvertiseNodeList
func (self *Nodes_NodeManager) GetDeclAdvertiseNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 64))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getProtoClassNodeList
func (self *Nodes_NodeManager) GetProtoClassNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 65))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclClassNodeList
func (self *Nodes_NodeManager) GetDeclClassNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 66))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclEnumNodeList
func (self *Nodes_NodeManager) GetDeclEnumNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 67))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclAlgeNodeList
func (self *Nodes_NodeManager) GetDeclAlgeNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 68))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getNewAlgeValNodeList
func (self *Nodes_NodeManager) GetNewAlgeValNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 69))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLuneControlNodeList
func (self *Nodes_NodeManager) GetLuneControlNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 70))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getMatchNodeList
func (self *Nodes_NodeManager) GetMatchNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 71))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLuneKindNodeList
func (self *Nodes_NodeManager) GetLuneKindNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 72))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getDeclMacroNodeList
func (self *Nodes_NodeManager) GetDeclMacroNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 73))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getTestCaseNodeList
func (self *Nodes_NodeManager) GetTestCaseNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 74))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getTestBlockNodeList
func (self *Nodes_NodeManager) GetTestBlockNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 75))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getAbbrNodeList
func (self *Nodes_NodeManager) GetAbbrNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 76))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getBoxingNodeList
func (self *Nodes_NodeManager) GetBoxingNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 77))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getUnboxingNodeList
func (self *Nodes_NodeManager) GetUnboxingNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 78))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralNilNodeList
func (self *Nodes_NodeManager) GetLiteralNilNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 79))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralCharNodeList
func (self *Nodes_NodeManager) GetLiteralCharNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 80))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralIntNodeList
func (self *Nodes_NodeManager) GetLiteralIntNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 81))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralRealNodeList
func (self *Nodes_NodeManager) GetLiteralRealNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 82))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralArrayNodeList
func (self *Nodes_NodeManager) GetLiteralArrayNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 83))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralListNodeList
func (self *Nodes_NodeManager) GetLiteralListNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 84))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralSetNodeList
func (self *Nodes_NodeManager) GetLiteralSetNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 85))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralMapNodeList
func (self *Nodes_NodeManager) GetLiteralMapNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 86))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralStringNodeList
func (self *Nodes_NodeManager) GetLiteralStringNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 87))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralBoolNodeList
func (self *Nodes_NodeManager) GetLiteralBoolNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 88))
}
// 1: decl @lune.@base.@Nodes.NodeManager.getLiteralSymbolNodeList
func (self *Nodes_NodeManager) GetLiteralSymbolNodeList(_env *LnsEnv) *LnsList {
    return (*LnsList)(self.FP.GetList(_env, 89))
}
// 2932: decl @lune.@base.@Nodes.NodeManager.MultiTo1
func (self *Nodes_NodeManager) MultiTo1(_env *LnsEnv, node *Nodes_Node) *Nodes_Node {
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType(_env)
    if node.FP.Get_expTypeList(_env).Len() > 1{
        return &Nodes_ExpMultiTo1Node_create(_env, self, node.FP.Get_pos(_env), node.FP.Get_inTestBlock(_env), node.FP.Get_macroArgFlag(_env), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(expType)}), node).Nodes_Node
    } else if expType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        return &Nodes_ExpMultiTo1Node_create(_env, self, node.FP.Get_pos(_env), node.FP.Get_inTestBlock(_env), node.FP.Get_macroArgFlag(_env), expType.FP.Get_itemTypeInfoList(_env), node).Nodes_Node
    }
    return node
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_None
func Nodes_NodeKind_get_None(_env *LnsEnv) LnsInt {
    return 0
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Shebang
func Nodes_NodeKind_get_Shebang(_env *LnsEnv) LnsInt {
    return 1
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ConvStat
func Nodes_NodeKind_get_ConvStat(_env *LnsEnv) LnsInt {
    return 2
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_BlankLine
func Nodes_NodeKind_get_BlankLine(_env *LnsEnv) LnsInt {
    return 3
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Subfile
func Nodes_NodeKind_get_Subfile(_env *LnsEnv) LnsInt {
    return 4
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Import
func Nodes_NodeKind_get_Import(_env *LnsEnv) LnsInt {
    return 5
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Root
func Nodes_NodeKind_get_Root(_env *LnsEnv) LnsInt {
    return 6
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_RefType
func Nodes_NodeKind_get_RefType(_env *LnsEnv) LnsInt {
    return 7
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Block
func Nodes_NodeKind_get_Block(_env *LnsEnv) LnsInt {
    return 8
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Scope
func Nodes_NodeKind_get_Scope(_env *LnsEnv) LnsInt {
    return 9
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_If
func Nodes_NodeKind_get_If(_env *LnsEnv) LnsInt {
    return 10
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpList
func Nodes_NodeKind_get_ExpList(_env *LnsEnv) LnsInt {
    return 11
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Switch
func Nodes_NodeKind_get_Switch(_env *LnsEnv) LnsInt {
    return 12
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_While
func Nodes_NodeKind_get_While(_env *LnsEnv) LnsInt {
    return 13
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Repeat
func Nodes_NodeKind_get_Repeat(_env *LnsEnv) LnsInt {
    return 14
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_For
func Nodes_NodeKind_get_For(_env *LnsEnv) LnsInt {
    return 15
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Apply
func Nodes_NodeKind_get_Apply(_env *LnsEnv) LnsInt {
    return 16
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Foreach
func Nodes_NodeKind_get_Foreach(_env *LnsEnv) LnsInt {
    return 17
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Forsort
func Nodes_NodeKind_get_Forsort(_env *LnsEnv) LnsInt {
    return 18
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Return
func Nodes_NodeKind_get_Return(_env *LnsEnv) LnsInt {
    return 19
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Break
func Nodes_NodeKind_get_Break(_env *LnsEnv) LnsInt {
    return 20
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Provide
func Nodes_NodeKind_get_Provide(_env *LnsEnv) LnsInt {
    return 21
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpNew
func Nodes_NodeKind_get_ExpNew(_env *LnsEnv) LnsInt {
    return 22
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpUnwrap
func Nodes_NodeKind_get_ExpUnwrap(_env *LnsEnv) LnsInt {
    return 23
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpRef
func Nodes_NodeKind_get_ExpRef(_env *LnsEnv) LnsInt {
    return 24
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpSetVal
func Nodes_NodeKind_get_ExpSetVal(_env *LnsEnv) LnsInt {
    return 25
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpSetItem
func Nodes_NodeKind_get_ExpSetItem(_env *LnsEnv) LnsInt {
    return 26
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpOp2
func Nodes_NodeKind_get_ExpOp2(_env *LnsEnv) LnsInt {
    return 27
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_UnwrapSet
func Nodes_NodeKind_get_UnwrapSet(_env *LnsEnv) LnsInt {
    return 28
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_IfUnwrap
func Nodes_NodeKind_get_IfUnwrap(_env *LnsEnv) LnsInt {
    return 29
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_When
func Nodes_NodeKind_get_When(_env *LnsEnv) LnsInt {
    return 30
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpCast
func Nodes_NodeKind_get_ExpCast(_env *LnsEnv) LnsInt {
    return 31
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpToDDD
func Nodes_NodeKind_get_ExpToDDD(_env *LnsEnv) LnsInt {
    return 32
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpSubDDD
func Nodes_NodeKind_get_ExpSubDDD(_env *LnsEnv) LnsInt {
    return 33
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpOp1
func Nodes_NodeKind_get_ExpOp1(_env *LnsEnv) LnsInt {
    return 34
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpRefItem
func Nodes_NodeKind_get_ExpRefItem(_env *LnsEnv) LnsInt {
    return 35
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpCall
func Nodes_NodeKind_get_ExpCall(_env *LnsEnv) LnsInt {
    return 36
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpMRet
func Nodes_NodeKind_get_ExpMRet(_env *LnsEnv) LnsInt {
    return 37
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpAccessMRet
func Nodes_NodeKind_get_ExpAccessMRet(_env *LnsEnv) LnsInt {
    return 38
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpMultiTo1
func Nodes_NodeKind_get_ExpMultiTo1(_env *LnsEnv) LnsInt {
    return 39
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpParen
func Nodes_NodeKind_get_ExpParen(_env *LnsEnv) LnsInt {
    return 40
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpMacroExp
func Nodes_NodeKind_get_ExpMacroExp(_env *LnsEnv) LnsInt {
    return 41
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpMacroStat
func Nodes_NodeKind_get_ExpMacroStat(_env *LnsEnv) LnsInt {
    return 42
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpMacroArgExp
func Nodes_NodeKind_get_ExpMacroArgExp(_env *LnsEnv) LnsInt {
    return 43
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_StmtExp
func Nodes_NodeKind_get_StmtExp(_env *LnsEnv) LnsInt {
    return 44
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpMacroStatList
func Nodes_NodeKind_get_ExpMacroStatList(_env *LnsEnv) LnsInt {
    return 45
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpOmitEnum
func Nodes_NodeKind_get_ExpOmitEnum(_env *LnsEnv) LnsInt {
    return 46
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_RefField
func Nodes_NodeKind_get_RefField(_env *LnsEnv) LnsInt {
    return 47
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_GetField
func Nodes_NodeKind_get_GetField(_env *LnsEnv) LnsInt {
    return 48
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Alias
func Nodes_NodeKind_get_Alias(_env *LnsEnv) LnsInt {
    return 49
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclVar
func Nodes_NodeKind_get_DeclVar(_env *LnsEnv) LnsInt {
    return 50
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclForm
func Nodes_NodeKind_get_DeclForm(_env *LnsEnv) LnsInt {
    return 51
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclFunc
func Nodes_NodeKind_get_DeclFunc(_env *LnsEnv) LnsInt {
    return 52
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclMethod
func Nodes_NodeKind_get_DeclMethod(_env *LnsEnv) LnsInt {
    return 53
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ProtoMethod
func Nodes_NodeKind_get_ProtoMethod(_env *LnsEnv) LnsInt {
    return 54
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclConstr
func Nodes_NodeKind_get_DeclConstr(_env *LnsEnv) LnsInt {
    return 55
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclDestr
func Nodes_NodeKind_get_DeclDestr(_env *LnsEnv) LnsInt {
    return 56
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpCallSuperCtor
func Nodes_NodeKind_get_ExpCallSuperCtor(_env *LnsEnv) LnsInt {
    return 57
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ExpCallSuper
func Nodes_NodeKind_get_ExpCallSuper(_env *LnsEnv) LnsInt {
    return 58
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_AsyncLock
func Nodes_NodeKind_get_AsyncLock(_env *LnsEnv) LnsInt {
    return 59
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Request
func Nodes_NodeKind_get_Request(_env *LnsEnv) LnsInt {
    return 60
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclMember
func Nodes_NodeKind_get_DeclMember(_env *LnsEnv) LnsInt {
    return 61
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclArg
func Nodes_NodeKind_get_DeclArg(_env *LnsEnv) LnsInt {
    return 62
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclArgDDD
func Nodes_NodeKind_get_DeclArgDDD(_env *LnsEnv) LnsInt {
    return 63
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclAdvertise
func Nodes_NodeKind_get_DeclAdvertise(_env *LnsEnv) LnsInt {
    return 64
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_ProtoClass
func Nodes_NodeKind_get_ProtoClass(_env *LnsEnv) LnsInt {
    return 65
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclClass
func Nodes_NodeKind_get_DeclClass(_env *LnsEnv) LnsInt {
    return 66
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclEnum
func Nodes_NodeKind_get_DeclEnum(_env *LnsEnv) LnsInt {
    return 67
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclAlge
func Nodes_NodeKind_get_DeclAlge(_env *LnsEnv) LnsInt {
    return 68
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_NewAlgeVal
func Nodes_NodeKind_get_NewAlgeVal(_env *LnsEnv) LnsInt {
    return 69
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LuneControl
func Nodes_NodeKind_get_LuneControl(_env *LnsEnv) LnsInt {
    return 70
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Match
func Nodes_NodeKind_get_Match(_env *LnsEnv) LnsInt {
    return 71
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LuneKind
func Nodes_NodeKind_get_LuneKind(_env *LnsEnv) LnsInt {
    return 72
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_DeclMacro
func Nodes_NodeKind_get_DeclMacro(_env *LnsEnv) LnsInt {
    return 73
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_TestCase
func Nodes_NodeKind_get_TestCase(_env *LnsEnv) LnsInt {
    return 74
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_TestBlock
func Nodes_NodeKind_get_TestBlock(_env *LnsEnv) LnsInt {
    return 75
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Abbr
func Nodes_NodeKind_get_Abbr(_env *LnsEnv) LnsInt {
    return 76
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Boxing
func Nodes_NodeKind_get_Boxing(_env *LnsEnv) LnsInt {
    return 77
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_Unboxing
func Nodes_NodeKind_get_Unboxing(_env *LnsEnv) LnsInt {
    return 78
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralNil
func Nodes_NodeKind_get_LiteralNil(_env *LnsEnv) LnsInt {
    return 79
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralChar
func Nodes_NodeKind_get_LiteralChar(_env *LnsEnv) LnsInt {
    return 80
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralInt
func Nodes_NodeKind_get_LiteralInt(_env *LnsEnv) LnsInt {
    return 81
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralReal
func Nodes_NodeKind_get_LiteralReal(_env *LnsEnv) LnsInt {
    return 82
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralArray
func Nodes_NodeKind_get_LiteralArray(_env *LnsEnv) LnsInt {
    return 83
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralList
func Nodes_NodeKind_get_LiteralList(_env *LnsEnv) LnsInt {
    return 84
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralSet
func Nodes_NodeKind_get_LiteralSet(_env *LnsEnv) LnsInt {
    return 85
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralMap
func Nodes_NodeKind_get_LiteralMap(_env *LnsEnv) LnsInt {
    return 86
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralString
func Nodes_NodeKind_get_LiteralString(_env *LnsEnv) LnsInt {
    return 87
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralBool
func Nodes_NodeKind_get_LiteralBool(_env *LnsEnv) LnsInt {
    return 88
}
// 725: decl @lune.@base.@Nodes.NodeKind.get_LiteralSymbol
func Nodes_NodeKind_get_LiteralSymbol(_env *LnsEnv) LnsInt {
    return 89
}
// 1: decl @lune.@base.@Nodes.NoneNode.processFilter
func (self *Nodes_NoneNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessNone(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.NoneNode.canBeRight
func (self *Nodes_NoneNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.NoneNode.canBeLeft
func (self *Nodes_NoneNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.NoneNode.canBeStatement
func (self *Nodes_NoneNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.NoneNode.create
func Nodes_NoneNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) *Nodes_NoneNode {
    var node *Nodes_NoneNode
    node = NewNodes_NoneNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.NoneNode.visit
func (self *Nodes_NoneNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ShebangNode.processFilter
func (self *Nodes_ShebangNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessShebang(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ShebangNode.canBeRight
func (self *Nodes_ShebangNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ShebangNode.canBeLeft
func (self *Nodes_ShebangNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ShebangNode.canBeStatement
func (self *Nodes_ShebangNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ShebangNode.create
func Nodes_ShebangNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,cmd string) *Nodes_ShebangNode {
    var node *Nodes_ShebangNode
    node = NewNodes_ShebangNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, cmd)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ShebangNode.visit
func (self *Nodes_ShebangNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ConvStatNode.processFilter
func (self *Nodes_ConvStatNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessConvStat(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ConvStatNode.canBeRight
func (self *Nodes_ConvStatNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ConvStatNode.canBeLeft
func (self *Nodes_ConvStatNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ConvStatNode.canBeStatement
func (self *Nodes_ConvStatNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ConvStatNode.create
func Nodes_ConvStatNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,txt string) *Nodes_ConvStatNode {
    var node *Nodes_ConvStatNode
    node = NewNodes_ConvStatNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, txt)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ConvStatNode.visit
func (self *Nodes_ConvStatNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.BlankLineNode.processFilter
func (self *Nodes_BlankLineNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessBlankLine(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.BlankLineNode.canBeRight
func (self *Nodes_BlankLineNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.BlankLineNode.canBeLeft
func (self *Nodes_BlankLineNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.BlankLineNode.canBeStatement
func (self *Nodes_BlankLineNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.BlankLineNode.create
func Nodes_BlankLineNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,lineNum LnsInt) *Nodes_BlankLineNode {
    var node *Nodes_BlankLineNode
    node = NewNodes_BlankLineNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, lineNum)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.BlankLineNode.visit
func (self *Nodes_BlankLineNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.SubfileNode.processFilter
func (self *Nodes_SubfileNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessSubfile(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.SubfileNode.canBeRight
func (self *Nodes_SubfileNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.SubfileNode.canBeLeft
func (self *Nodes_SubfileNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.SubfileNode.canBeStatement
func (self *Nodes_SubfileNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.SubfileNode.create
func Nodes_SubfileNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,usePath LnsAny) *Nodes_SubfileNode {
    var node *Nodes_SubfileNode
    node = NewNodes_SubfileNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, usePath)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.SubfileNode.visit
func (self *Nodes_SubfileNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ImportNode.processFilter
func (self *Nodes_ImportNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessImport(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ImportNode.canBeRight
func (self *Nodes_ImportNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ImportNode.canBeLeft
func (self *Nodes_ImportNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ImportNode.canBeStatement
func (self *Nodes_ImportNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ImportNode.create
func Nodes_ImportNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,info *Nodes_ImportInfo) *Nodes_ImportNode {
    var node *Nodes_ImportNode
    node = NewNodes_ImportNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, info)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ImportNode.visit
func (self *Nodes_ImportNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.RootNode.processFilter
func (self *Nodes_RootNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRoot(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.RootNode.canBeRight
func (self *Nodes_RootNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.RootNode.canBeLeft
func (self *Nodes_RootNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.RootNode.canBeStatement
func (self *Nodes_RootNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.RootNode.create
func Nodes_RootNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,children *LnsList,moduleScope *Ast_Scope,globalScope *Ast_Scope,useModuleMacroSet *LnsSet,moduleId *FrontInterface_ModuleId,processInfo *Ast_ProcessInfo,moduleTypeInfo *Ast_TypeInfo,provideNode LnsAny,luneHelperInfo *FrontInterface_LuneHelperInfo,nodeManager *Nodes_NodeManager,importModule2moduleInfo *LnsMap,typeId2MacroInfo *LnsMap,typeId2ClassMap *LnsMap) *Nodes_RootNode {
    var node *Nodes_RootNode
    node = NewNodes_RootNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, children, moduleScope, globalScope, useModuleMacroSet, moduleId, processInfo, moduleTypeInfo, provideNode, luneHelperInfo, nodeManager, importModule2moduleInfo, typeId2MacroInfo, typeId2ClassMap)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.RootNode.visit
func (self *Nodes_RootNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var list *LnsList
        list = self.children
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch0 := visitor(_env, child, &self.Nodes_Node, "children", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    {
        {
            _child := self.provideNode
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ProvideNode)
                if Lns_op_not(alreadySet.Has(Nodes_ProvideNode2Stem(child))){
                    alreadySet.Add(Nodes_ProvideNode2Stem(child))
                    if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "provideNode", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 875: decl @lune.@base.@Nodes.RootNode.set_provide
func (self *Nodes_RootNode) Set_provide(_env *LnsEnv, node *Nodes_ProvideNode) {
    self.provideNode = node
}
// 1: decl @lune.@base.@Nodes.RefTypeNode.processFilter
func (self *Nodes_RefTypeNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRefType(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.RefTypeNode.canBeRight
func (self *Nodes_RefTypeNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.RefTypeNode.canBeLeft
func (self *Nodes_RefTypeNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.RefTypeNode.canBeStatement
func (self *Nodes_RefTypeNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.RefTypeNode.create
func Nodes_RefTypeNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Nodes_Node,itemNodeList *LnsList,mutMode LnsAny,array string) *Nodes_RefTypeNode {
    var node *Nodes_RefTypeNode
    node = NewNodes_RefTypeNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, name, itemNodeList, mutMode, array)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.RefTypeNode.visit
func (self *Nodes_RefTypeNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.name
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "name", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var list *LnsList
        list = self.itemNodeList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch1 := visitor(_env, child, &self.Nodes_Node, "itemNodeList", depth); _switch1 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch1 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch1 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.BlockNode.processFilter
func (self *Nodes_BlockNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessBlock(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.BlockNode.canBeRight
func (self *Nodes_BlockNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.BlockNode.canBeLeft
func (self *Nodes_BlockNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.BlockNode.canBeStatement
func (self *Nodes_BlockNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.BlockNode.create
func Nodes_BlockNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,blockKind LnsInt,scope *Ast_Scope,stmtList *LnsList) *Nodes_BlockNode {
    var node *Nodes_BlockNode
    node = NewNodes_BlockNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, blockKind, scope, stmtList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.BlockNode.visit
func (self *Nodes_BlockNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var list *LnsList
        list = self.stmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch0 := visitor(_env, child, &self.Nodes_Node, "stmtList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 988: decl @lune.@base.@Nodes.BlockNode.getBreakKind
func (self *Nodes_BlockNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    return Nodes_getBreakKindForStmtList_42_(_env, checkMode, self.stmtList)
}
// 1: decl @lune.@base.@Nodes.ScopeNode.processFilter
func (self *Nodes_ScopeNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessScope(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ScopeNode.canBeRight
func (self *Nodes_ScopeNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ScopeNode.canBeLeft
func (self *Nodes_ScopeNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ScopeNode.canBeStatement
func (self *Nodes_ScopeNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ScopeNode.create
func Nodes_ScopeNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,scopeKind LnsInt,scope *Ast_Scope,symbolList *LnsList,block *Nodes_BlockNode) *Nodes_ScopeNode {
    var node *Nodes_ScopeNode
    node = NewNodes_ScopeNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, scopeKind, scope, symbolList, block)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ScopeNode.visit
func (self *Nodes_ScopeNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.IfNode.processFilter
func (self *Nodes_IfNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessIf(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.IfNode.canBeRight
func (self *Nodes_IfNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.IfNode.canBeLeft
func (self *Nodes_IfNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.IfNode.canBeStatement
func (self *Nodes_IfNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.IfNode.create
func Nodes_IfNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) *Nodes_IfNode {
    var node *Nodes_IfNode
    node = NewNodes_IfNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, stmtList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.IfNode.visit
func (self *Nodes_IfNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1008: decl @lune.@base.@Nodes.IfNode.getBreakKind
func (self *Nodes_IfNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    var hasElseFlag bool
    hasElseFlag = false
    var kind LnsInt
    kind = Nodes_BreakKind__None
    for _, _stmtInfo := range( self.stmtList.Items ) {
        stmtInfo := _stmtInfo.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        var work LnsInt
        work = stmtInfo.FP.Get_block(_env).FP.GetBreakKind(_env, checkMode)
        if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
            if work == Nodes_BreakKind__Return{
                return Nodes_BreakKind__Return
            }
            if work == Nodes_BreakKind__NeverRet{
                return Nodes_BreakKind__NeverRet
            }
        } else { 
            if _switch0 := work; _switch0 == Nodes_BreakKind__None {
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                    _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                    if true{
                        return Nodes_BreakKind__None
                    }
                }
            } else {
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                    _env.SetStackVal( kind > work) ).(bool){
                    kind = work
                }
            }
        }
        
        if stmtInfo.FP.Get_kind(_env) == Nodes_IfKind__Else{
            hasElseFlag = true
        }
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( hasElseFlag) ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
            _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool))) ).(bool){
        return kind
    }
    return Nodes_BreakKind__None
}
// 1025: decl @lune.@base.@Nodes.IfNode.visitSub
func (self *Nodes_IfNode) VisitSub(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    for _, _stmt := range( self.FP.Get_stmtList(_env).Items ) {
        stmt := _stmt.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(stmt.FP.Get_exp(_env)))){
            alreadySet.Add(Nodes_Node2Stem(stmt.FP.Get_exp(_env)))
            if _switch0 := visitor(_env, stmt.FP.Get_exp(_env), &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(stmt.FP.Get_exp(_env).FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(stmt.FP.Get_block(_env)))){
            alreadySet.Add(Nodes_BlockNode2Stem(stmt.FP.Get_block(_env)))
            if _switch1 := visitor(_env, &stmt.FP.Get_block(_env).Nodes_Node, &self.Nodes_Node, "block", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(stmt.FP.Get_block(_env).FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return true
}
// 1: decl @lune.@base.@Nodes.ExpListNode.processFilter
func (self *Nodes_ExpListNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpList(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpListNode.canBeStatement
func (self *Nodes_ExpListNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpListNode.create
func Nodes_ExpListNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList *LnsList,mRetExp LnsAny,followOn bool) *Nodes_ExpListNode {
    var node *Nodes_ExpListNode
    node = NewNodes_ExpListNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, expList, mRetExp, followOn)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpListNode.visit
func (self *Nodes_ExpListNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var list *LnsList
        list = self.expList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch0 := visitor(_env, child, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1050: decl @lune.@base.@Nodes.ExpListNode.canBeLeft
func (self *Nodes_ExpListNode) CanBeLeft(_env *LnsEnv) bool {
    for _, _expNode := range( self.FP.Get_expList(_env).Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(expNode.FP.CanBeLeft(_env)){
            return false
        }
    }
    return true
}
// 1058: decl @lune.@base.@Nodes.ExpListNode.canBeRight
func (self *Nodes_ExpListNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    for _, _expNode := range( self.FP.Get_expList(_env).Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(expNode.FP.CanBeRight(_env, processInfo)){
            return false
        }
    }
    return true
}
// 1066: decl @lune.@base.@Nodes.ExpListNode.setLValue
func (self *Nodes_ExpListNode) SetLValue(_env *LnsEnv) {
    for _, _expNode := range( self.FP.Get_expList(_env).Items ) {
        expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
        expNode.IsLValue = true
    }
}
// 1083: decl @lune.@base.@Nodes.ExpListNode.getExpTypeAt
func (self *Nodes_ExpListNode) GetExpTypeAt(_env *LnsEnv, index LnsInt) *Ast_TypeInfo {
    if index > self.FP.Get_expTypeList(_env).Len(){
        var lastExpType *Ast_TypeInfo
        lastExpType = self.FP.Get_expTypeList(_env).GetAt(self.FP.Get_expTypeList(_env).Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        {
            _dddTypeInfo := Ast_DDDTypeInfoDownCastF(lastExpType.FP)
            if !Lns_IsNil( _dddTypeInfo ) {
                dddTypeInfo := _dddTypeInfo.(*Ast_DDDTypeInfo)
                return dddTypeInfo.FP.Get_typeInfo(_env).FP.Get_nilableTypeInfo(_env)
            }
        }
        return Ast_builtinTypeNil
    }
    return self.FP.Get_expTypeList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
}
// 1103: decl @lune.@base.@Nodes.ExpListNode.getExpTypeNoDDDAt
func (self *Nodes_ExpListNode) GetExpTypeNoDDDAt(_env *LnsEnv, index LnsInt) *Ast_TypeInfo {
    if index >= self.FP.Get_expTypeList(_env).Len(){
        var lastExpType *Ast_TypeInfo
        lastExpType = self.FP.Get_expTypeList(_env).GetAt(self.FP.Get_expTypeList(_env).Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        {
            _dddTypeInfo := Ast_DDDTypeInfoDownCastF(lastExpType.FP)
            if !Lns_IsNil( _dddTypeInfo ) {
                dddTypeInfo := _dddTypeInfo.(*Ast_DDDTypeInfo)
                return dddTypeInfo.FP.Get_typeInfo(_env).FP.Get_nilableTypeInfo(_env)
            }
        }
        if index == self.FP.Get_expTypeList(_env).Len(){
            return lastExpType
        }
        return Ast_builtinTypeNil
    }
    return self.FP.Get_expTypeList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
}
// 1: decl @lune.@base.@Nodes.SwitchNode.processFilter
func (self *Nodes_SwitchNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessSwitch(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.SwitchNode.canBeRight
func (self *Nodes_SwitchNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.SwitchNode.canBeLeft
func (self *Nodes_SwitchNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.SwitchNode.canBeStatement
func (self *Nodes_SwitchNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.SwitchNode.create
func Nodes_SwitchNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,exp *Nodes_Node,caseList *LnsList,_default LnsAny,caseKind LnsInt,failSafeDefault bool) *Nodes_SwitchNode {
    var node *Nodes_SwitchNode
    node = NewNodes_SwitchNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, idInNS, exp, caseList, _default, caseKind, failSafeDefault)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.SwitchNode.visit
func (self *Nodes_SwitchNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self._default
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
                    alreadySet.Add(Nodes_BlockNode2Stem(child))
                    if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "default", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1141: decl @lune.@base.@Nodes.SwitchNode.getBreakKind
func (self *Nodes_SwitchNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = Nodes_BreakKind__None
    var fullCase bool
    fullCase = self.caseKind != Nodes_CaseKind__Lack
    for _, _caseInfo := range( self.caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        var work LnsInt
        work = caseInfo.FP.Get_block(_env).FP.GetBreakKind(_env, checkMode)
        var goNext bool
        goNext = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( (work == Nodes_BreakKind__None)) ||
            _env.SetStackVal( Lns_op_not(fullCase)) ).(bool)
        if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
            if work == Nodes_BreakKind__Return{
                return Nodes_BreakKind__Return
            }
            if work == Nodes_BreakKind__NeverRet{
                return Nodes_BreakKind__NeverRet
            }
        } else { 
            if _switch0 := work; _switch0 == Nodes_BreakKind__None {
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                    _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                    if goNext{
                        return Nodes_BreakKind__None
                    }
                }
            } else {
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                    _env.SetStackVal( kind > work) ).(bool){
                    kind = work
                }
            }
        }
        
    }
    {
        _block := self._default
        if !Lns_IsNil( _block ) {
            block := _block.(*Nodes_BlockNode)
            var work LnsInt
            work = block.FP.GetBreakKind(_env, checkMode)
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch1 := work; _switch1 == Nodes_BreakKind__None {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                        _env.SetStackVal( kind > work) ).(bool){
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
    if kind == Nodes_BreakKind__Break{
        return kind
    }
    return Nodes_BreakKind__None
}
// 1170: decl @lune.@base.@Nodes.SwitchNode.visitSub
func (self *Nodes_SwitchNode) VisitSub(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    for _, _caseInfo := range( self.caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(caseInfo.FP.Get_expList(_env)))){
            alreadySet.Add(Nodes_ExpListNode2Stem(caseInfo.FP.Get_expList(_env)))
            if _switch0 := visitor(_env, &caseInfo.FP.Get_expList(_env).Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(caseInfo.FP.Get_expList(_env).FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(caseInfo.FP.Get_block(_env)))){
            alreadySet.Add(Nodes_BlockNode2Stem(caseInfo.FP.Get_block(_env)))
            if _switch1 := visitor(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, &self.Nodes_Node, "block", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(caseInfo.FP.Get_block(_env).FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return true
}
// 1: decl @lune.@base.@Nodes.WhileNode.processFilter
func (self *Nodes_WhileNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessWhile(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.WhileNode.canBeRight
func (self *Nodes_WhileNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.WhileNode.canBeLeft
func (self *Nodes_WhileNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.WhileNode.canBeStatement
func (self *Nodes_WhileNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.WhileNode.create
func Nodes_WhileNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,infinit bool,block *Nodes_BlockNode) *Nodes_WhileNode {
    var node *Nodes_WhileNode
    node = NewNodes_WhileNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, exp, infinit, block)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.WhileNode.visit
func (self *Nodes_WhileNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2316: decl @lune.@base.@Nodes.WhileNode.getBreakKind
func (self *Nodes_WhileNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        var kind LnsInt
        kind = Nodes_BreakKind__None
        for _, _stmt := range( self.block.FP.Get_stmtList(_env).Items ) {
            stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
            if stmt.FP.Get_kind(_env) != Nodes_NodeKind_get_BlankLine(_env){
                var work LnsInt
                work = stmt.FP.GetBreakKind(_env, checkMode)
                if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                    if work == Nodes_BreakKind__Return{
                        return Nodes_BreakKind__Return
                    }
                    if work == Nodes_BreakKind__NeverRet{
                        return Nodes_BreakKind__NeverRet
                    }
                } else { 
                    if _switch0 := work; _switch0 == Nodes_BreakKind__None {
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                            _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                            if false{
                                return Nodes_BreakKind__None
                            }
                        }
                    } else {
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                            _env.SetStackVal( kind > work) ).(bool){
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
        for _, _stmt := range( self.block.FP.Get_stmtList(_env).Items ) {
            stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
            if stmt.FP.Get_kind(_env) != Nodes_NodeKind_get_BlankLine(_env){
                var work LnsInt
                work = stmt.FP.GetBreakKind(_env, mode)
                if mode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                    if work == Nodes_BreakKind__Return{
                        return Nodes_BreakKind__Return
                    }
                    if work == Nodes_BreakKind__NeverRet{
                        return Nodes_BreakKind__NeverRet
                    }
                } else { 
                    if _switch1 := work; _switch1 == Nodes_BreakKind__None {
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( mode == Nodes_CheckBreakMode__Normal) ||
                            _env.SetStackVal( mode == Nodes_CheckBreakMode__Return) ).(bool){
                            if false{
                                return Nodes_BreakKind__None
                            }
                        }
                    } else {
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                            _env.SetStackVal( kind > work) ).(bool){
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
// 1: decl @lune.@base.@Nodes.RepeatNode.processFilter
func (self *Nodes_RepeatNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRepeat(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.RepeatNode.canBeRight
func (self *Nodes_RepeatNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.RepeatNode.canBeLeft
func (self *Nodes_RepeatNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.RepeatNode.canBeStatement
func (self *Nodes_RepeatNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.RepeatNode.create
func Nodes_RepeatNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,block *Nodes_BlockNode,exp *Nodes_Node) *Nodes_RepeatNode {
    var node *Nodes_RepeatNode
    node = NewNodes_RepeatNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, block, exp)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.RepeatNode.visit
func (self *Nodes_RepeatNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch1 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1183: decl @lune.@base.@Nodes.RepeatNode.getBreakKind
func (self *Nodes_RepeatNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(_env, checkMode)
    }
    return Nodes_BreakKind__None
}
// 1: decl @lune.@base.@Nodes.ForNode.processFilter
func (self *Nodes_ForNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessFor(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ForNode.canBeRight
func (self *Nodes_ForNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ForNode.canBeLeft
func (self *Nodes_ForNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ForNode.canBeStatement
func (self *Nodes_ForNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ForNode.create
func Nodes_ForNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,block *Nodes_BlockNode,val *Ast_SymbolInfo,init *Nodes_Node,to *Nodes_Node,delta LnsAny) *Nodes_ForNode {
    var node *Nodes_ForNode
    node = NewNodes_ForNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, idInNS, block, val, init, to, delta)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ForNode.visit
func (self *Nodes_ForNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_Node
        child = self.init
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch1 := visitor(_env, child, &self.Nodes_Node, "init", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_Node
        child = self.to
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch2 := visitor(_env, child, &self.Nodes_Node, "to", depth); _switch2 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch2 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch2 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self.delta
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_Node)
                if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                    alreadySet.Add(Nodes_Node2Stem(child))
                    if _switch3 := visitor(_env, child, &self.Nodes_Node, "delta", depth); _switch3 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch3 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch3 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1183: decl @lune.@base.@Nodes.ForNode.getBreakKind
func (self *Nodes_ForNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(_env, checkMode)
    }
    return Nodes_BreakKind__None
}
// 1: decl @lune.@base.@Nodes.ApplyNode.processFilter
func (self *Nodes_ApplyNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessApply(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ApplyNode.canBeRight
func (self *Nodes_ApplyNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ApplyNode.canBeLeft
func (self *Nodes_ApplyNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ApplyNode.canBeStatement
func (self *Nodes_ApplyNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ApplyNode.create
func Nodes_ApplyNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,varList *LnsList,expList *Nodes_ExpListNode,block *Nodes_BlockNode) *Nodes_ApplyNode {
    var node *Nodes_ApplyNode
    node = NewNodes_ApplyNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, idInNS, varList, expList, block)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ApplyNode.visit
func (self *Nodes_ApplyNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_ExpListNode
        child = self.expList
        if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
            alreadySet.Add(Nodes_ExpListNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1183: decl @lune.@base.@Nodes.ApplyNode.getBreakKind
func (self *Nodes_ApplyNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(_env, checkMode)
    }
    return Nodes_BreakKind__None
}
// 1: decl @lune.@base.@Nodes.ForeachNode.processFilter
func (self *Nodes_ForeachNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessForeach(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ForeachNode.canBeRight
func (self *Nodes_ForeachNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ForeachNode.canBeLeft
func (self *Nodes_ForeachNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ForeachNode.canBeStatement
func (self *Nodes_ForeachNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ForeachNode.create
func Nodes_ForeachNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,val *Ast_SymbolInfo,key LnsAny,exp *Nodes_Node,block *Nodes_BlockNode) *Nodes_ForeachNode {
    var node *Nodes_ForeachNode
    node = NewNodes_ForeachNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, idInNS, val, key, exp, block)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ForeachNode.visit
func (self *Nodes_ForeachNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1183: decl @lune.@base.@Nodes.ForeachNode.getBreakKind
func (self *Nodes_ForeachNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(_env, checkMode)
    }
    return Nodes_BreakKind__None
}
// 1: decl @lune.@base.@Nodes.ForsortNode.processFilter
func (self *Nodes_ForsortNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessForsort(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ForsortNode.canBeRight
func (self *Nodes_ForsortNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ForsortNode.canBeLeft
func (self *Nodes_ForsortNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ForsortNode.canBeStatement
func (self *Nodes_ForsortNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ForsortNode.create
func Nodes_ForsortNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,val *Ast_SymbolInfo,key LnsAny,exp *Nodes_Node,block *Nodes_BlockNode,sort bool) *Nodes_ForsortNode {
    var node *Nodes_ForsortNode
    node = NewNodes_ForsortNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, idInNS, val, key, exp, block, sort)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ForsortNode.visit
func (self *Nodes_ForsortNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1183: decl @lune.@base.@Nodes.ForsortNode.getBreakKind
func (self *Nodes_ForsortNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
        _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
        return self.block.FP.GetBreakKind(_env, checkMode)
    }
    return Nodes_BreakKind__None
}
// 1: decl @lune.@base.@Nodes.ReturnNode.processFilter
func (self *Nodes_ReturnNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessReturn(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ReturnNode.canBeRight
func (self *Nodes_ReturnNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ReturnNode.canBeLeft
func (self *Nodes_ReturnNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ReturnNode.canBeStatement
func (self *Nodes_ReturnNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ReturnNode.create
func Nodes_ReturnNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList LnsAny) *Nodes_ReturnNode {
    var node *Nodes_ReturnNode
    node = NewNodes_ReturnNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ReturnNode.visit
func (self *Nodes_ReturnNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.expList
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1254: decl @lune.@base.@Nodes.ReturnNode.getBreakKind
func (self *Nodes_ReturnNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    return Nodes_BreakKind__Return
}
// 1: decl @lune.@base.@Nodes.BreakNode.processFilter
func (self *Nodes_BreakNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessBreak(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.BreakNode.canBeRight
func (self *Nodes_BreakNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.BreakNode.canBeLeft
func (self *Nodes_BreakNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.BreakNode.canBeStatement
func (self *Nodes_BreakNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.BreakNode.create
func Nodes_BreakNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) *Nodes_BreakNode {
    var node *Nodes_BreakNode
    node = NewNodes_BreakNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.BreakNode.visit
func (self *Nodes_BreakNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1260: decl @lune.@base.@Nodes.BreakNode.getBreakKind
func (self *Nodes_BreakNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    return Nodes_BreakKind__Break
}
// 1: decl @lune.@base.@Nodes.ProvideNode.processFilter
func (self *Nodes_ProvideNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessProvide(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ProvideNode.canBeRight
func (self *Nodes_ProvideNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ProvideNode.canBeLeft
func (self *Nodes_ProvideNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ProvideNode.canBeStatement
func (self *Nodes_ProvideNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ProvideNode.create
func Nodes_ProvideNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,symbol *Ast_SymbolInfo) *Nodes_ProvideNode {
    var node *Nodes_ProvideNode
    node = NewNodes_ProvideNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, symbol)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ProvideNode.visit
func (self *Nodes_ProvideNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpNewNode.processFilter
func (self *Nodes_ExpNewNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpNew(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpNewNode.canBeRight
func (self *Nodes_ExpNewNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpNewNode.canBeLeft
func (self *Nodes_ExpNewNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpNewNode.canBeStatement
func (self *Nodes_ExpNewNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ExpNewNode.create
func Nodes_ExpNewNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,symbol *Nodes_Node,ctorTypeInfo *Ast_TypeInfo,argList LnsAny) *Nodes_ExpNewNode {
    var node *Nodes_ExpNewNode
    node = NewNodes_ExpNewNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, symbol, ctorTypeInfo, argList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpNewNode.visit
func (self *Nodes_ExpNewNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.symbol
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "symbol", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self.argList
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "argList", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpUnwrapNode.processFilter
func (self *Nodes_ExpUnwrapNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpUnwrap(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpUnwrapNode.canBeRight
func (self *Nodes_ExpUnwrapNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpUnwrapNode.canBeLeft
func (self *Nodes_ExpUnwrapNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpUnwrapNode.canBeStatement
func (self *Nodes_ExpUnwrapNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpUnwrapNode.create
func Nodes_ExpUnwrapNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,_default LnsAny) *Nodes_ExpUnwrapNode {
    var node *Nodes_ExpUnwrapNode
    node = NewNodes_ExpUnwrapNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, exp, _default)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpUnwrapNode.visit
func (self *Nodes_ExpUnwrapNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self._default
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_Node)
                if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                    alreadySet.Add(Nodes_Node2Stem(child))
                    if _switch1 := visitor(_env, child, &self.Nodes_Node, "default", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpRefNode.processFilter
func (self *Nodes_ExpRefNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpRef(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpRefNode.canBeStatement
func (self *Nodes_ExpRefNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpRefNode.create
func Nodes_ExpRefNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,symbolInfo *Ast_SymbolInfo) *Nodes_ExpRefNode {
    var node *Nodes_ExpRefNode
    node = NewNodes_ExpRefNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, symbolInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpRefNode.visit
func (self *Nodes_ExpRefNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1284: decl @lune.@base.@Nodes.ExpRefNode.canBeLeft
func (self *Nodes_ExpRefNode) CanBeLeft(_env *LnsEnv) bool {
    return self.FP.Get_symbolInfo(_env).FP.Get_canBeLeft(_env)
}
// 1288: decl @lune.@base.@Nodes.ExpRefNode.canBeRight
func (self *Nodes_ExpRefNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.Get_symbolInfo(_env).FP.Get_canBeRight(_env)) &&
        _env.SetStackVal( self.FP.Get_symbolInfo(_env).FP.Get_hasValueFlag(_env)) ).(bool)
}
// 2713: decl @lune.@base.@Nodes.ExpRefNode.getLiteral
func (self *Nodes_ExpRefNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = self.symbolInfo.FP.Get_typeInfo(_env)
    {
        _enumTypeInfo := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_aliasSrc(_env).FP)
        if !Lns_IsNil( _enumTypeInfo ) {
            enumTypeInfo := _enumTypeInfo.(*Ast_EnumTypeInfo)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) &&
                _env.SetStackVal( self.symbolInfo.FP.Get_namespaceTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) ).(bool)){
                var enumval *Ast_EnumValInfo
                enumval = Lns_unwrap( enumTypeInfo.FP.GetEnumValInfo(_env, self.symbolInfo.FP.Get_name(_env))).(*Ast_EnumValInfo)
                return Nodes_enumLiteral2Literal_142_(_env, enumval.FP.Get_val(_env))
            }
        }
    }
    return nil, _env.GetVM().String_format("unsupport refnode -- %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)})
}
// 1: decl @lune.@base.@Nodes.ExpSetValNode.processFilter
func (self *Nodes_ExpSetValNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpSetVal(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpSetValNode.canBeRight
func (self *Nodes_ExpSetValNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpSetValNode.canBeLeft
func (self *Nodes_ExpSetValNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpSetValNode.canBeStatement
func (self *Nodes_ExpSetValNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ExpSetValNode.create
func Nodes_ExpSetValNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp1 *Nodes_Node,exp2 *Nodes_ExpListNode,LeftSymList *LnsList,initSymSet *LnsSet) *Nodes_ExpSetValNode {
    var node *Nodes_ExpSetValNode
    node = NewNodes_ExpSetValNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, exp1, exp2, LeftSymList, initSymSet)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpSetValNode.visit
func (self *Nodes_ExpSetValNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp1
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp1", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_ExpListNode
        child = self.exp2
        if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
            alreadySet.Add(Nodes_ExpListNode2Stem(child))
            if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "exp2", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpSetItemNode.processFilter
func (self *Nodes_ExpSetItemNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpSetItem(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpSetItemNode.canBeRight
func (self *Nodes_ExpSetItemNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpSetItemNode.canBeLeft
func (self *Nodes_ExpSetItemNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpSetItemNode.canBeStatement
func (self *Nodes_ExpSetItemNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ExpSetItemNode.create
func Nodes_ExpSetItemNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,index LnsAny,exp2 *Nodes_Node) *Nodes_ExpSetItemNode {
    var node *Nodes_ExpSetItemNode
    node = NewNodes_ExpSetItemNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, val, index, exp2)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpSetItemNode.visit
func (self *Nodes_ExpSetItemNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.val
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "val", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_Node
        child = self.exp2
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch1 := visitor(_env, child, &self.Nodes_Node, "exp2", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpOp2Node.processFilter
func (self *Nodes_ExpOp2Node) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpOp2(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpOp2Node.canBeRight
func (self *Nodes_ExpOp2Node) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpOp2Node.canBeLeft
func (self *Nodes_ExpOp2Node) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpOp2Node.canBeStatement
func (self *Nodes_ExpOp2Node) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpOp2Node.create
func Nodes_ExpOp2Node_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,op *Types_Token,exp1 *Nodes_Node,exp2 *Nodes_Node) *Nodes_ExpOp2Node {
    var node *Nodes_ExpOp2Node
    node = NewNodes_ExpOp2Node(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, op, exp1, exp2)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpOp2Node.visit
func (self *Nodes_ExpOp2Node) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp1
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp1", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_Node
        child = self.exp2
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch1 := visitor(_env, child, &self.Nodes_Node, "exp2", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2751: decl @lune.@base.@Nodes.ExpOp2Node.getValType
func (self *Nodes_ExpOp2Node) getValType(_env *LnsEnv, node *Nodes_Node)(bool, LnsInt, LnsReal, string, *Ast_TypeInfo) {
    var literal LnsAny
    
    {
        _literal := Nodes_convExp0_30087(Lns_2DDD(node.FP.GetLiteral(_env)))
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
    switch _matchExp0 := literal.(type) {
    case *Nodes_Literal__Int:
    val := _matchExp0.Val1
        intVal = val
        realVal = (LnsReal)(val)
        retTypeInfo = Ast_builtinTypeInt
    case *Nodes_Literal__Real:
    val := _matchExp0.Val1
        realVal = val
        intVal = (LnsInt)(val)
        retTypeInfo = Ast_builtinTypeReal
    case *Nodes_Literal__Str:
    val := _matchExp0.Val1
        strVal = val
        retTypeInfo = Ast_builtinTypeString
    default:
        return false, 0, 0.0, "", Ast_headTypeInfo
    }
    return true, intVal, realVal, strVal, retTypeInfo
}
// 2801: decl @lune.@base.@Nodes.ExpOp2Node.setupLiteralTokenList
func (self *Nodes_ExpOp2Node) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    var literal LnsAny
    literal = Nodes_convExp0_30211(Lns_2DDD(self.FP.GetLiteral(_env)))
    if literal != nil{
        literal_7324 := literal
        switch _matchExp0 := literal_7324.(type) {
        case *Nodes_Literal__Int:
        val := _matchExp0.Val1
            self.FP.AddTokenList(_env, list, Types_TokenKind__Int, _env.GetVM().String_format("%d", []LnsAny{val}))
        case *Nodes_Literal__Real:
        val := _matchExp0.Val1
            self.FP.AddTokenList(_env, list, Types_TokenKind__Real, _env.GetVM().String_format("%g", []LnsAny{val}))
        case *Nodes_Literal__Str:
        val := _matchExp0.Val1
            self.FP.AddTokenList(_env, list, Types_TokenKind__Str, _env.GetVM().String_format("%q", []LnsAny{val}))
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
// 2826: decl @lune.@base.@Nodes.ExpOp2Node.getLiteral
func (self *Nodes_ExpOp2Node) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var ret1 bool
    var int1 LnsInt
    var real1 LnsReal
    var str1 string
    var type1 *Ast_TypeInfo
    ret1,int1,real1,str1,type1 = self.FP.getValType(_env, self.FP.Get_exp1(_env))
    var ret2 bool
    var int2 LnsInt
    var real2 LnsReal
    var str2 string
    var type2 *Ast_TypeInfo
    ret2,int2,real2,str2,type2 = self.FP.getValType(_env, self.FP.Get_exp2(_env))
    if Lns_op_not(ret1){
        return nil, _env.GetVM().String_format("not support literal -- %s", []LnsAny{Nodes_getNodeKindName(_env, self.FP.Get_exp1(_env).FP.Get_kind(_env))})
    }
    if Lns_op_not(ret2){
        return nil, _env.GetVM().String_format("not support literal -- %s", []LnsAny{Nodes_getNodeKindName(_env, self.FP.Get_exp2(_env).FP.Get_kind(_env))})
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( type1 == Ast_builtinTypeInt) ||
            _env.SetStackVal( type1 == Ast_builtinTypeReal) ).(bool))) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( type2 == Ast_builtinTypeInt) ||
            _env.SetStackVal( type2 == Ast_builtinTypeReal) ).(bool))) ).(bool)){
        var retType *Ast_TypeInfo
        retType = Ast_builtinTypeInt
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( type1 == Ast_builtinTypeReal) ||
            _env.SetStackVal( type2 == Ast_builtinTypeReal) ).(bool){
            retType = Ast_builtinTypeReal
        }
        if _switch0 := (self.op.Txt); _switch0 == "+" {
            if retType == Ast_builtinTypeInt{
                return &Nodes_Literal__Int{int1 + int2}, nil
            }
            return &Nodes_Literal__Real{real1 + real2}, nil
        } else if _switch0 == "-" {
            if retType == Ast_builtinTypeInt{
                return &Nodes_Literal__Int{int1 - int2}, nil
            }
            return &Nodes_Literal__Real{real1 - real2}, nil
        } else if _switch0 == "*" {
            if retType == Ast_builtinTypeInt{
                return &Nodes_Literal__Int{int1 * int2}, nil
            }
            return &Nodes_Literal__Real{real1 * real2}, nil
        } else if _switch0 == "/" {
            if retType == Ast_builtinTypeInt{
                return &Nodes_Literal__Int{int1 / int2}, nil
            }
            return &Nodes_Literal__Real{real1 / real2}, nil
        }
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( type1 == Ast_builtinTypeString) &&
        _env.SetStackVal( type2 == Ast_builtinTypeString) ).(bool)){
        if self.op.Txt == ".."{
            return &Nodes_Literal__Str{str1 + str2}, nil
        }
    }
    var mess string
    mess = _env.GetVM().String_format("not support literal operation -- %s %s %s", []LnsAny{type1.FP.GetTxt(_env, nil, nil, nil), self.op.Txt, type2.FP.GetTxt(_env, nil, nil, nil)})
    return nil, mess
}
// 1: decl @lune.@base.@Nodes.UnwrapSetNode.processFilter
func (self *Nodes_UnwrapSetNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessUnwrapSet(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.UnwrapSetNode.canBeRight
func (self *Nodes_UnwrapSetNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.UnwrapSetNode.canBeLeft
func (self *Nodes_UnwrapSetNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.UnwrapSetNode.canBeStatement
func (self *Nodes_UnwrapSetNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.UnwrapSetNode.create
func Nodes_UnwrapSetNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,dstExpList *Nodes_ExpListNode,srcExpList *Nodes_ExpListNode,unwrapBlock LnsAny) *Nodes_UnwrapSetNode {
    var node *Nodes_UnwrapSetNode
    node = NewNodes_UnwrapSetNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, dstExpList, srcExpList, unwrapBlock)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.UnwrapSetNode.visit
func (self *Nodes_UnwrapSetNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_ExpListNode
        child = self.dstExpList
        if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
            alreadySet.Add(Nodes_ExpListNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "dstExpList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_ExpListNode
        child = self.srcExpList
        if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
            alreadySet.Add(Nodes_ExpListNode2Stem(child))
            if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "srcExpList", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self.unwrapBlock
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
                    alreadySet.Add(Nodes_BlockNode2Stem(child))
                    if _switch2 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "unwrapBlock", depth); _switch2 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch2 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch2 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.IfUnwrapNode.processFilter
func (self *Nodes_IfUnwrapNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessIfUnwrap(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.IfUnwrapNode.canBeRight
func (self *Nodes_IfUnwrapNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.IfUnwrapNode.canBeLeft
func (self *Nodes_IfUnwrapNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.IfUnwrapNode.canBeStatement
func (self *Nodes_IfUnwrapNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.IfUnwrapNode.create
func Nodes_IfUnwrapNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,varSymList *LnsList,expList *Nodes_ExpListNode,block *Nodes_BlockNode,nilBlock LnsAny) *Nodes_IfUnwrapNode {
    var node *Nodes_IfUnwrapNode
    node = NewNodes_IfUnwrapNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, varSymList, expList, block, nilBlock)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.IfUnwrapNode.visit
func (self *Nodes_IfUnwrapNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_ExpListNode
        child = self.expList
        if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
            alreadySet.Add(Nodes_ExpListNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self.nilBlock
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
                    alreadySet.Add(Nodes_BlockNode2Stem(child))
                    if _switch2 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "nilBlock", depth); _switch2 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch2 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch2 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1342: decl @lune.@base.@Nodes.IfUnwrapNode.getBreakKind
func (self *Nodes_IfUnwrapNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = self.block.FP.GetBreakKind(_env, checkMode)
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
        if _switch0 := work; _switch0 == Nodes_BreakKind__None {
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                if true{
                    return Nodes_BreakKind__None
                }
            }
        } else {
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                _env.SetStackVal( kind > work) ).(bool){
                kind = work
            }
        }
    }
    
    {
        _block := self.nilBlock
        if !Lns_IsNil( _block ) {
            block := _block.(*Nodes_BlockNode)
            work = block.FP.GetBreakKind(_env, checkMode)
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch1 := work; _switch1 == Nodes_BreakKind__None {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                        _env.SetStackVal( kind > work) ).(bool){
                        kind = work
                    }
                }
            }
            
            return kind
        }
    }
    return Nodes_BreakKind__None
}
// 1: decl @lune.@base.@Nodes.WhenNode.processFilter
func (self *Nodes_WhenNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessWhen(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.WhenNode.canBeRight
func (self *Nodes_WhenNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.WhenNode.canBeLeft
func (self *Nodes_WhenNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.WhenNode.canBeStatement
func (self *Nodes_WhenNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.WhenNode.create
func Nodes_WhenNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,symPairList *LnsList,block *Nodes_BlockNode,elseBlock LnsAny) *Nodes_WhenNode {
    var node *Nodes_WhenNode
    node = NewNodes_WhenNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, symPairList, block, elseBlock)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.WhenNode.visit
func (self *Nodes_WhenNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self.elseBlock
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
                    alreadySet.Add(Nodes_BlockNode2Stem(child))
                    if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "elseBlock", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1368: decl @lune.@base.@Nodes.WhenNode.getBreakKind
func (self *Nodes_WhenNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = self.block.FP.GetBreakKind(_env, checkMode)
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
        if _switch0 := work; _switch0 == Nodes_BreakKind__None {
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                if true{
                    return Nodes_BreakKind__None
                }
            }
        } else {
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                _env.SetStackVal( kind > work) ).(bool){
                kind = work
            }
        }
    }
    
    {
        _block := self.elseBlock
        if !Lns_IsNil( _block ) {
            block := _block.(*Nodes_BlockNode)
            work = block.FP.GetBreakKind(_env, checkMode)
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch1 := work; _switch1 == Nodes_BreakKind__None {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                        _env.SetStackVal( kind > work) ).(bool){
                        kind = work
                    }
                }
            }
            
            return kind
        }
    }
    return Nodes_BreakKind__None
}
// 1: decl @lune.@base.@Nodes.ExpCastNode.processFilter
func (self *Nodes_ExpCastNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCast(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpCastNode.canBeRight
func (self *Nodes_ExpCastNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpCastNode.canBeLeft
func (self *Nodes_ExpCastNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpCastNode.canBeStatement
func (self *Nodes_ExpCastNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpCastNode.create
func Nodes_ExpCastNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node,castType *Ast_TypeInfo,castTypeNode LnsAny,castOpe string,castKind LnsInt) *Nodes_ExpCastNode {
    var node *Nodes_ExpCastNode
    node = NewNodes_ExpCastNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, exp, castType, castTypeNode, castOpe, castKind)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpCastNode.visit
func (self *Nodes_ExpCastNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self.castTypeNode
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_RefTypeNode)
                if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(child))){
                    alreadySet.Add(Nodes_RefTypeNode2Stem(child))
                    if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "castTypeNode", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1407: decl @lune.@base.@Nodes.ExpCastNode.getPrefix
func (self *Nodes_ExpCastNode) GetPrefix(_env *LnsEnv) LnsAny {
    return self.exp.FP.GetPrefix(_env)
}
// 1410: decl @lune.@base.@Nodes.ExpCastNode.getLiteral
func (self *Nodes_ExpCastNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    return self.exp.FP.GetLiteral(_env)
}
// 1413: decl @lune.@base.@Nodes.ExpCastNode.setupLiteralTokenList
func (self *Nodes_ExpCastNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    return self.exp.FP.SetupLiteralTokenList(_env, list)
}
// 1: decl @lune.@base.@Nodes.ExpToDDDNode.processFilter
func (self *Nodes_ExpToDDDNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpToDDD(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpToDDDNode.canBeRight
func (self *Nodes_ExpToDDDNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpToDDDNode.canBeLeft
func (self *Nodes_ExpToDDDNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpToDDDNode.canBeStatement
func (self *Nodes_ExpToDDDNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpToDDDNode.create
func Nodes_ExpToDDDNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList *Nodes_ExpListNode) *Nodes_ExpToDDDNode {
    var node *Nodes_ExpToDDDNode
    node = NewNodes_ExpToDDDNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpToDDDNode.visit
func (self *Nodes_ExpToDDDNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_ExpListNode
        child = self.expList
        if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
            alreadySet.Add(Nodes_ExpListNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpSubDDDNode.processFilter
func (self *Nodes_ExpSubDDDNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpSubDDD(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpSubDDDNode.canBeRight
func (self *Nodes_ExpSubDDDNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpSubDDDNode.canBeLeft
func (self *Nodes_ExpSubDDDNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpSubDDDNode.canBeStatement
func (self *Nodes_ExpSubDDDNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpSubDDDNode.create
func Nodes_ExpSubDDDNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,src *Nodes_Node,remainIndex LnsInt) *Nodes_ExpSubDDDNode {
    var node *Nodes_ExpSubDDDNode
    node = NewNodes_ExpSubDDDNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, src, remainIndex)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpSubDDDNode.visit
func (self *Nodes_ExpSubDDDNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.src
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "src", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpOp1Node.processFilter
func (self *Nodes_ExpOp1Node) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpOp1(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpOp1Node.canBeRight
func (self *Nodes_ExpOp1Node) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpOp1Node.canBeLeft
func (self *Nodes_ExpOp1Node) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpOp1Node.canBeStatement
func (self *Nodes_ExpOp1Node) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpOp1Node.create
func Nodes_ExpOp1Node_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,op *Types_Token,macroMode LnsInt,exp *Nodes_Node) *Nodes_ExpOp1Node {
    var node *Nodes_ExpOp1Node
    node = NewNodes_ExpOp1Node(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, op, macroMode, exp)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpOp1Node.visit
func (self *Nodes_ExpOp1Node) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpRefItemNode.hasNilAccess
func (self *Nodes_ExpRefItemNode) HasNilAccess(_env *LnsEnv) bool {
    return self.nilAccess
}
// 1: decl @lune.@base.@Nodes.ExpRefItemNode.processFilter
func (self *Nodes_ExpRefItemNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpRefItem(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpRefItemNode.canBeRight
func (self *Nodes_ExpRefItemNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpRefItemNode.canBeStatement
func (self *Nodes_ExpRefItemNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpRefItemNode.create
func Nodes_ExpRefItemNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,val *Nodes_Node,nilAccess bool,symbol LnsAny,index LnsAny) *Nodes_ExpRefItemNode {
    var node *Nodes_ExpRefItemNode
    node = NewNodes_ExpRefItemNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, val, nilAccess, symbol, index)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpRefItemNode.visit
func (self *Nodes_ExpRefItemNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.val
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "val", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self.index
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_Node)
                if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                    alreadySet.Add(Nodes_Node2Stem(child))
                    if _switch1 := visitor(_env, child, &self.Nodes_Node, "index", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1451: decl @lune.@base.@Nodes.ExpRefItemNode.getPrefix
func (self *Nodes_ExpRefItemNode) GetPrefix(_env *LnsEnv) LnsAny {
    return self.val
}
// 1454: decl @lune.@base.@Nodes.ExpRefItemNode.canBeLeft
func (self *Nodes_ExpRefItemNode) CanBeLeft(_env *LnsEnv) bool {
    if self.val.FP.Get_expType(_env) == Ast_builtinTypeStem{
        return false
    }
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Ast_TypeInfo_isMut(_env, self.FP.Get_val(_env).FP.Get_expType(_env))) &&
        _env.SetStackVal( Lns_op_not(self.nilAccess)) ).(bool)
}
// 1: decl @lune.@base.@Nodes.ExpCallNode.hasNilAccess
func (self *Nodes_ExpCallNode) HasNilAccess(_env *LnsEnv) bool {
    return self.nilAccess
}
// 1: decl @lune.@base.@Nodes.ExpCallNode.processFilter
func (self *Nodes_ExpCallNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCall(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpCallNode.canBeLeft
func (self *Nodes_ExpCallNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpCallNode.canBeStatement
func (self *Nodes_ExpCallNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ExpCallNode.create
func Nodes_ExpCallNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,_func *Nodes_Node,errorFunc bool,nilAccess bool,argList LnsAny) *Nodes_ExpCallNode {
    var node *Nodes_ExpCallNode
    node = NewNodes_ExpCallNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, _func, errorFunc, nilAccess, argList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpCallNode.visit
func (self *Nodes_ExpCallNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self._func
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "func", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self.argList
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "argList", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1468: decl @lune.@base.@Nodes.ExpCallNode.get_effectivePos
func (self *Nodes_ExpCallNode) Get_effectivePos(_env *LnsEnv) Types_Position {
    return self._func.FP.Get_effectivePos(_env)
}
// 1471: decl @lune.@base.@Nodes.ExpCallNode.getPrefix
func (self *Nodes_ExpCallNode) GetPrefix(_env *LnsEnv) LnsAny {
    return self._func
}
// 1475: decl @lune.@base.@Nodes.ExpCallNode.canBeRight
func (self *Nodes_ExpCallNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    var expType *Ast_TypeInfo
    expType = self.FP.Get_expType(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( expType.FP.Equals(_env, processInfo, Ast_builtinTypeNone, nil, nil)) ||
        _env.SetStackVal( expType.FP.Equals(_env, processInfo, Ast_builtinTypeNeverRet, nil, nil)) ).(bool){
        return false
    }
    return true
}
// 1486: decl @lune.@base.@Nodes.ExpCallNode.getBreakKind
func (self *Nodes_ExpCallNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    if self.errorFunc{
        return Nodes_BreakKind__NeverRet
    }
    return Nodes_BreakKind__None
}
// 1: decl @lune.@base.@Nodes.ExpMRetNode.processFilter
func (self *Nodes_ExpMRetNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMRet(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpMRetNode.canBeRight
func (self *Nodes_ExpMRetNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpMRetNode.canBeLeft
func (self *Nodes_ExpMRetNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpMRetNode.canBeStatement
func (self *Nodes_ExpMRetNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpMRetNode.create
func Nodes_ExpMRetNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,mRet *Nodes_Node) *Nodes_ExpMRetNode {
    var node *Nodes_ExpMRetNode
    node = NewNodes_ExpMRetNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, mRet)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpMRetNode.visit
func (self *Nodes_ExpMRetNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.mRet
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "mRet", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1497: decl @lune.@base.@Nodes.ExpMRetNode.getPrefix
func (self *Nodes_ExpMRetNode) GetPrefix(_env *LnsEnv) LnsAny {
    return self.mRet.FP.GetPrefix(_env)
}
// 1: decl @lune.@base.@Nodes.ExpAccessMRetNode.processFilter
func (self *Nodes_ExpAccessMRetNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpAccessMRet(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpAccessMRetNode.canBeRight
func (self *Nodes_ExpAccessMRetNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpAccessMRetNode.canBeLeft
func (self *Nodes_ExpAccessMRetNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpAccessMRetNode.canBeStatement
func (self *Nodes_ExpAccessMRetNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpAccessMRetNode.create
func Nodes_ExpAccessMRetNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,mRet *Nodes_Node,index LnsInt) *Nodes_ExpAccessMRetNode {
    var node *Nodes_ExpAccessMRetNode
    node = NewNodes_ExpAccessMRetNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, mRet, index)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpAccessMRetNode.visit
func (self *Nodes_ExpAccessMRetNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.mRet
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "mRet", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1506: decl @lune.@base.@Nodes.ExpAccessMRetNode.getPrefix
func (self *Nodes_ExpAccessMRetNode) GetPrefix(_env *LnsEnv) LnsAny {
    return self.mRet.FP.GetPrefix(_env)
}
// 1: decl @lune.@base.@Nodes.ExpMultiTo1Node.processFilter
func (self *Nodes_ExpMultiTo1Node) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMultiTo1(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpMultiTo1Node.canBeRight
func (self *Nodes_ExpMultiTo1Node) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpMultiTo1Node.canBeLeft
func (self *Nodes_ExpMultiTo1Node) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpMultiTo1Node.canBeStatement
func (self *Nodes_ExpMultiTo1Node) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpMultiTo1Node.create
func Nodes_ExpMultiTo1Node_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_ExpMultiTo1Node {
    var node *Nodes_ExpMultiTo1Node
    node = NewNodes_ExpMultiTo1Node(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpMultiTo1Node.visit
func (self *Nodes_ExpMultiTo1Node) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1513: decl @lune.@base.@Nodes.ExpMultiTo1Node.getPrefix
func (self *Nodes_ExpMultiTo1Node) GetPrefix(_env *LnsEnv) LnsAny {
    return self.exp.FP.GetPrefix(_env)
}
// 1: decl @lune.@base.@Nodes.ExpParenNode.processFilter
func (self *Nodes_ExpParenNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpParen(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpParenNode.canBeRight
func (self *Nodes_ExpParenNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpParenNode.canBeLeft
func (self *Nodes_ExpParenNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpParenNode.canBeStatement
func (self *Nodes_ExpParenNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpParenNode.create
func Nodes_ExpParenNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_ExpParenNode {
    var node *Nodes_ExpParenNode
    node = NewNodes_ExpParenNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpParenNode.visit
func (self *Nodes_ExpParenNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1520: decl @lune.@base.@Nodes.ExpParenNode.getPrefix
func (self *Nodes_ExpParenNode) GetPrefix(_env *LnsEnv) LnsAny {
    return self.exp.FP.GetPrefix(_env)
}
// 1524: decl @lune.@base.@Nodes.ExpParenNode.getSymbolInfo
func (self *Nodes_ExpParenNode) GetSymbolInfo(_env *LnsEnv) *LnsList {
    return self.exp.FP.GetSymbolInfo(_env)
}
// 1: decl @lune.@base.@Nodes.ExpMacroExpNode.processFilter
func (self *Nodes_ExpMacroExpNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMacroExp(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpMacroExpNode.canBeLeft
func (self *Nodes_ExpMacroExpNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpMacroExpNode.canBeStatement
func (self *Nodes_ExpMacroExpNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ExpMacroExpNode.create
func Nodes_ExpMacroExpNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,macroType *Ast_TypeInfo,stmtList *LnsList) *Nodes_ExpMacroExpNode {
    var node *Nodes_ExpMacroExpNode
    node = NewNodes_ExpMacroExpNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, macroType, stmtList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpMacroExpNode.visit
func (self *Nodes_ExpMacroExpNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var list *LnsList
        list = self.stmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch0 := visitor(_env, child, &self.Nodes_Node, "stmtList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1535: decl @lune.@base.@Nodes.ExpMacroExpNode.canBeRight
func (self *Nodes_ExpMacroExpNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return self.FP.Get_expType(_env) != Ast_builtinTypeNone
}
// 1539: decl @lune.@base.@Nodes.ExpMacroExpNode.getBreakKind
func (self *Nodes_ExpMacroExpNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    return Nodes_getBreakKindForStmtList_42_(_env, checkMode, self.stmtList)
}
// 1: decl @lune.@base.@Nodes.ExpMacroStatNode.processFilter
func (self *Nodes_ExpMacroStatNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMacroStat(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpMacroStatNode.canBeRight
func (self *Nodes_ExpMacroStatNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpMacroStatNode.canBeLeft
func (self *Nodes_ExpMacroStatNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpMacroStatNode.canBeStatement
func (self *Nodes_ExpMacroStatNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpMacroStatNode.create
func Nodes_ExpMacroStatNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expStrList *LnsList) *Nodes_ExpMacroStatNode {
    var node *Nodes_ExpMacroStatNode
    node = NewNodes_ExpMacroStatNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, expStrList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpMacroStatNode.visit
func (self *Nodes_ExpMacroStatNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var list *LnsList
        list = self.expStrList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch0 := visitor(_env, child, &self.Nodes_Node, "expStrList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2688: decl @lune.@base.@Nodes.ExpMacroStatNode.getLiteral
func (self *Nodes_ExpMacroStatNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var txt string
    txt = ""
    for _, _token := range( self.expStrList.Items ) {
        token := _token.(Nodes_NodeDownCast).ToNodes_Node()
        var literal LnsAny
        literal = Nodes_convExp0_29856(Lns_2DDD(token.FP.GetLiteral(_env)))
        if literal != nil{
            literal_7286 := literal
            switch _matchExp0 := literal_7286.(type) {
            case *Nodes_Literal__Str:
            work := _matchExp0.Val1
                txt = _env.GetVM().String_format("%s%s", []LnsAny{txt, work})
            }
        } else {
            return nil, _env.GetVM().String_format("illegal literal -- %s", []LnsAny{Nodes_getNodeKindName(_env, token.FP.Get_kind(_env))})
        }
    }
    return &Nodes_Literal__Str{txt}, nil
}
// 1: decl @lune.@base.@Nodes.ExpMacroArgExpNode.processFilter
func (self *Nodes_ExpMacroArgExpNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMacroArgExp(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpMacroArgExpNode.canBeRight
func (self *Nodes_ExpMacroArgExpNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpMacroArgExpNode.canBeLeft
func (self *Nodes_ExpMacroArgExpNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpMacroArgExpNode.canBeStatement
func (self *Nodes_ExpMacroArgExpNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpMacroArgExpNode.create
func Nodes_ExpMacroArgExpNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,codeTxt string) *Nodes_ExpMacroArgExpNode {
    var node *Nodes_ExpMacroArgExpNode
    node = NewNodes_ExpMacroArgExpNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, codeTxt)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpMacroArgExpNode.visit
func (self *Nodes_ExpMacroArgExpNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2707: decl @lune.@base.@Nodes.ExpMacroArgExpNode.getLiteral
func (self *Nodes_ExpMacroArgExpNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    return &Nodes_Literal__Str{self.FP.Get_codeTxt(_env)}, nil
}
// 1: decl @lune.@base.@Nodes.StmtExpNode.processFilter
func (self *Nodes_StmtExpNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessStmtExp(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.StmtExpNode.canBeRight
func (self *Nodes_StmtExpNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.StmtExpNode.canBeLeft
func (self *Nodes_StmtExpNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.StmtExpNode.create
func Nodes_StmtExpNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_StmtExpNode {
    var node *Nodes_StmtExpNode
    node = NewNodes_StmtExpNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.StmtExpNode.visit
func (self *Nodes_StmtExpNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1560: decl @lune.@base.@Nodes.StmtExpNode.canBeStatement
func (self *Nodes_StmtExpNode) CanBeStatement(_env *LnsEnv) bool {
    return self.FP.Get_exp(_env).FP.CanBeStatement(_env)
}
// 1564: decl @lune.@base.@Nodes.StmtExpNode.getBreakKind
func (self *Nodes_StmtExpNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    return self.FP.Get_exp(_env).FP.GetBreakKind(_env, checkMode)
}
// 1: decl @lune.@base.@Nodes.ExpMacroStatListNode.processFilter
func (self *Nodes_ExpMacroStatListNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpMacroStatList(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpMacroStatListNode.canBeRight
func (self *Nodes_ExpMacroStatListNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpMacroStatListNode.canBeLeft
func (self *Nodes_ExpMacroStatListNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpMacroStatListNode.canBeStatement
func (self *Nodes_ExpMacroStatListNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpMacroStatListNode.create
func Nodes_ExpMacroStatListNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_ExpMacroStatListNode {
    var node *Nodes_ExpMacroStatListNode
    node = NewNodes_ExpMacroStatListNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpMacroStatListNode.visit
func (self *Nodes_ExpMacroStatListNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpOmitEnumNode.processFilter
func (self *Nodes_ExpOmitEnumNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpOmitEnum(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpOmitEnumNode.canBeRight
func (self *Nodes_ExpOmitEnumNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpOmitEnumNode.canBeLeft
func (self *Nodes_ExpOmitEnumNode) CanBeLeft(_env *LnsEnv) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.ExpOmitEnumNode.canBeStatement
func (self *Nodes_ExpOmitEnumNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.ExpOmitEnumNode.create
func Nodes_ExpOmitEnumNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,valToken *Types_Token,valInfo *Ast_EnumValInfo,aliasType LnsAny,enumTypeInfo *Ast_EnumTypeInfo) *Nodes_ExpOmitEnumNode {
    var node *Nodes_ExpOmitEnumNode
    node = NewNodes_ExpOmitEnumNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, valToken, valInfo, aliasType, enumTypeInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpOmitEnumNode.visit
func (self *Nodes_ExpOmitEnumNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2732: decl @lune.@base.@Nodes.ExpOmitEnumNode.getLiteral
func (self *Nodes_ExpOmitEnumNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var enumval *Ast_EnumValInfo
    enumval = self.valInfo
    return Nodes_enumLiteral2Literal_142_(_env, enumval.FP.Get_val(_env))
}
// 2740: decl @lune.@base.@Nodes.ExpOmitEnumNode.setupLiteralTokenList
func (self *Nodes_ExpOmitEnumNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    var enumval *Ast_EnumValInfo
    enumval = self.valInfo
    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, ".")
    self.FP.AddTokenList(_env, list, Types_TokenKind__Symb, Lns_car(_env.GetVM().String_gsub(enumval.FP.Get_val(_env).(LnsAlgeVal).GetTxt(),".*%.", "")).(string))
    return true
}
// 1: decl @lune.@base.@Nodes.RefFieldNode.hasNilAccess
func (self *Nodes_RefFieldNode) HasNilAccess(_env *LnsEnv) bool {
    return self.nilAccess
}
// 1: decl @lune.@base.@Nodes.RefFieldNode.processFilter
func (self *Nodes_RefFieldNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRefField(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.RefFieldNode.canBeStatement
func (self *Nodes_RefFieldNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.RefFieldNode.create
func Nodes_RefFieldNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,field *Types_Token,symbolInfo LnsAny,nilAccess bool,prefix *Nodes_Node) *Nodes_RefFieldNode {
    var node *Nodes_RefFieldNode
    node = NewNodes_RefFieldNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.RefFieldNode.visit
func (self *Nodes_RefFieldNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.prefix
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "prefix", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1587: decl @lune.@base.@Nodes.RefFieldNode.get_effectivePos
func (self *Nodes_RefFieldNode) Get_effectivePos(_env *LnsEnv) Types_Position {
    return self.field.Pos
}
// 1590: decl @lune.@base.@Nodes.RefFieldNode.getPrefix
func (self *Nodes_RefFieldNode) GetPrefix(_env *LnsEnv) LnsAny {
    return self.prefix
}
// 1593: decl @lune.@base.@Nodes.RefFieldNode.canBeLeft
func (self *Nodes_RefFieldNode) CanBeLeft(_env *LnsEnv) bool {
    {
        __exp := self.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            return _exp.FP.Get_canBeLeft(_env)
        }
    }
    return false
}
// 1608: decl @lune.@base.@Nodes.RefFieldNode.canBeRight
func (self *Nodes_RefFieldNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    {
        __exp := self.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Ast_SymbolInfo)
            {
                _prefix := Nodes_ExpRefNodeDownCastF(self.FP.Get_prefix(_env).FP)
                if !Lns_IsNil( _prefix ) {
                    prefix := _prefix.(*Nodes_ExpRefNode)
                    if prefix.FP.Get_symbolInfo(_env).FP.Get_name(_env) == "self"{
                        return _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( _exp.FP.Get_canBeRight(_env)) &&
                            _env.SetStackVal( _exp.FP.Get_hasValueFlag(_env)) ).(bool)
                    }
                }
            }
            return _exp.FP.Get_canBeRight(_env)
        }
    }
    return true
}
// 2651: decl @lune.@base.@Nodes.RefFieldNode.getLiteral
func (self *Nodes_RefFieldNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var typeInfo *Ast_TypeInfo
    typeInfo = self.FP.Get_expType(_env)
    {
        _enumTypeInfo := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_aliasSrc(_env).FP)
        if !Lns_IsNil( _enumTypeInfo ) {
            enumTypeInfo := _enumTypeInfo.(*Ast_EnumTypeInfo)
            if Lns_isCondTrue( Ast_EnumTypeInfoDownCastF(self.prefix.FP.Get_expType(_env).FP.Get_aliasSrc(_env).FP)){
                var enumval *Ast_EnumValInfo
                enumval = Lns_unwrap( enumTypeInfo.FP.GetEnumValInfo(_env, self.field.Txt)).(*Ast_EnumValInfo)
                return Nodes_enumLiteral2Literal_142_(_env, enumval.FP.Get_val(_env))
            }
        }
    }
    var tokenList *LnsList
    tokenList = NewLnsList([]LnsAny{})
    var literal LnsAny
    var mess LnsAny
    literal,mess = self.prefix.FP.GetLiteral(_env)
    if literal != nil{
        literal_7269 := literal
        switch _matchExp0 := literal_7269.(type) {
        case *Nodes_Literal__Symbol:
        symbol := _matchExp0.Val1
            tokenList.Insert(symbol)
        case *Nodes_Literal__Field:
        symList := _matchExp0.Val1
            for _, _symbol := range( symList.Items ) {
                symbol := _symbol.(string)
                tokenList.Insert(symbol)
            }
        default:
            return nil, _env.GetVM().String_format("not support -- %s", []LnsAny{literal_7269.(LnsAlgeVal).GetTxt()})
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
// 1: decl @lune.@base.@Nodes.GetFieldNode.hasNilAccess
func (self *Nodes_GetFieldNode) HasNilAccess(_env *LnsEnv) bool {
    return self.nilAccess
}
// 1: decl @lune.@base.@Nodes.GetFieldNode.processFilter
func (self *Nodes_GetFieldNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessGetField(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.GetFieldNode.canBeRight
func (self *Nodes_GetFieldNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.GetFieldNode.canBeLeft
func (self *Nodes_GetFieldNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.GetFieldNode.canBeStatement
func (self *Nodes_GetFieldNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.GetFieldNode.create
func Nodes_GetFieldNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,field *Types_Token,symbolInfo LnsAny,nilAccess bool,prefix *Nodes_Node,getterTypeInfo *Ast_TypeInfo) *Nodes_GetFieldNode {
    var node *Nodes_GetFieldNode
    node = NewNodes_GetFieldNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, field, symbolInfo, nilAccess, prefix, getterTypeInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.GetFieldNode.visit
func (self *Nodes_GetFieldNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.prefix
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "prefix", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1632: decl @lune.@base.@Nodes.GetFieldNode.get_effectivePos
func (self *Nodes_GetFieldNode) Get_effectivePos(_env *LnsEnv) Types_Position {
    return self.field.Pos
}
// 1635: decl @lune.@base.@Nodes.GetFieldNode.getPrefix
func (self *Nodes_GetFieldNode) GetPrefix(_env *LnsEnv) LnsAny {
    return self.prefix
}
// 1: decl @lune.@base.@Nodes.AliasNode.processFilter
func (self *Nodes_AliasNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessAlias(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.AliasNode.canBeRight
func (self *Nodes_AliasNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.AliasNode.canBeLeft
func (self *Nodes_AliasNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.AliasNode.canBeStatement
func (self *Nodes_AliasNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.AliasNode.create
func Nodes_AliasNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,newSymbol *Ast_SymbolInfo,srcNode *Nodes_Node,typeInfo *Ast_TypeInfo) *Nodes_AliasNode {
    var node *Nodes_AliasNode
    node = NewNodes_AliasNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, newSymbol, srcNode, typeInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.AliasNode.visit
func (self *Nodes_AliasNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.srcNode
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "srcNode", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclVarNode.processFilter
func (self *Nodes_DeclVarNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclVar(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclVarNode.canBeRight
func (self *Nodes_DeclVarNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclVarNode.canBeLeft
func (self *Nodes_DeclVarNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclVarNode.canBeStatement
func (self *Nodes_DeclVarNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclVarNode.create
func Nodes_DeclVarNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,mode LnsInt,accessMode LnsInt,staticFlag bool,varList *LnsList,expList LnsAny,symbolInfoList *LnsList,typeInfoList *LnsList,unwrapFlag bool,unwrapBlock LnsAny,thenBlock LnsAny,syncVarList *LnsList,syncBlock LnsAny) *Nodes_DeclVarNode {
    var node *Nodes_DeclVarNode
    node = NewNodes_DeclVarNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, mode, accessMode, staticFlag, varList, expList, symbolInfoList, typeInfoList, unwrapFlag, unwrapBlock, thenBlock, syncVarList, syncBlock)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclVarNode.visit
func (self *Nodes_DeclVarNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.expList
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    {
        {
            _child := self.unwrapBlock
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
                    alreadySet.Add(Nodes_BlockNode2Stem(child))
                    if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "unwrapBlock", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    {
        {
            _child := self.thenBlock
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
                    alreadySet.Add(Nodes_BlockNode2Stem(child))
                    if _switch2 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "thenBlock", depth); _switch2 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch2 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch2 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    {
        {
            _child := self.syncBlock
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
                    alreadySet.Add(Nodes_BlockNode2Stem(child))
                    if _switch3 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "syncBlock", depth); _switch3 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch3 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch3 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1676: decl @lune.@base.@Nodes.DeclVarNode.getBreakKind
func (self *Nodes_DeclVarNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = Nodes_BreakKind__None
    var work LnsInt
    work = Nodes_BreakKind__None
    {
        _block := self.unwrapBlock
        if !Lns_IsNil( _block ) {
            block := _block.(*Nodes_BlockNode)
            work = block.FP.GetBreakKind(_env, checkMode)
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch0 := work; _switch0 == Nodes_BreakKind__None {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                        _env.SetStackVal( kind > work) ).(bool){
                        kind = work
                    }
                }
            }
            
            {
                _thenBlock := self.thenBlock
                if !Lns_IsNil( _thenBlock ) {
                    thenBlock := _thenBlock.(*Nodes_BlockNode)
                    work = thenBlock.FP.GetBreakKind(_env, checkMode)
                    if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                        if work == Nodes_BreakKind__Return{
                            return Nodes_BreakKind__Return
                        }
                        if work == Nodes_BreakKind__NeverRet{
                            return Nodes_BreakKind__NeverRet
                        }
                    } else { 
                        if _switch1 := work; _switch1 == Nodes_BreakKind__None {
                            if _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                                _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                                if true{
                                    return Nodes_BreakKind__None
                                }
                            }
                        } else {
                            if _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                                _env.SetStackVal( kind > work) ).(bool){
                                kind = work
                            }
                        }
                    }
                    
                    {
                        _syncBlock := self.syncBlock
                        if !Lns_IsNil( _syncBlock ) {
                            syncBlock := _syncBlock.(*Nodes_BlockNode)
                            work = syncBlock.FP.GetBreakKind(_env, checkMode)
                            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                                if work == Nodes_BreakKind__Return{
                                    return Nodes_BreakKind__Return
                                }
                                if work == Nodes_BreakKind__NeverRet{
                                    return Nodes_BreakKind__NeverRet
                                }
                            } else { 
                                if _switch2 := work; _switch2 == Nodes_BreakKind__None {
                                    if _env.PopVal( _env.IncStack() ||
                                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                                        if true{
                                            return Nodes_BreakKind__None
                                        }
                                    }
                                } else {
                                    if _env.PopVal( _env.IncStack() ||
                                        _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                                        _env.SetStackVal( kind > work) ).(bool){
                                        kind = work
                                    }
                                }
                            }
                            
                        }
                    }
                    return kind
                }
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Normal) &&
                _env.SetStackVal( checkMode != Nodes_CheckBreakMode__Return) ).(bool)){
                return kind
            }
        }
    }
    return Nodes_BreakKind__None
}
// 1745: decl @lune.@base.@Nodes.DeclVarNode.visitSub
func (self *Nodes_DeclVarNode) VisitSub(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    for _, _varInfo := range( self.varList.Items ) {
        varInfo := _varInfo.(Nodes_VarInfoDownCast).ToNodes_VarInfo()
        {
            _refTypeNode := varInfo.FP.Get_refType(_env)
            if !Lns_IsNil( _refTypeNode ) {
                refTypeNode := _refTypeNode.(*Nodes_RefTypeNode)
                if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(refTypeNode))){
                    alreadySet.Add(Nodes_RefTypeNode2Stem(refTypeNode))
                    if _switch0 := visitor(_env, &refTypeNode.Nodes_Node, &self.Nodes_Node, "refType", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(refTypeNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return true
}
// 1796: decl @lune.@base.@Nodes.DeclFuncInfo.createFrom
func Nodes_DeclFuncInfo_createFrom(_env *LnsEnv, info *Nodes_DeclFuncInfo,name *Types_Token,symbol *Ast_SymbolInfo) *Nodes_DeclFuncInfo {
    return NewNodes_DeclFuncInfo(_env, info.FP.Get_kind(_env), info.classTypeInfo, info.declClassNode, info.outsizeOfClass, name, symbol, info.argList, info.staticFlag, info.accessMode, info.asyncMode, info.body, info.retTypeInfoList, info.retTypeNodeList, info.has__func__Symbol, info.overrideFlag, info.stmtNum)
}
// 1: decl @lune.@base.@Nodes.DeclFormNode.processFilter
func (self *Nodes_DeclFormNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclForm(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclFormNode.canBeRight
func (self *Nodes_DeclFormNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclFormNode.canBeLeft
func (self *Nodes_DeclFormNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclFormNode.canBeStatement
func (self *Nodes_DeclFormNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclFormNode.create
func Nodes_DeclFormNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_DeclFormNode {
    var node *Nodes_DeclFormNode
    node = NewNodes_DeclFormNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclFormNode.visit
func (self *Nodes_DeclFormNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        for _, _argNode := range( self.declInfo.FP.Get_argList(_env).Items ) {
            argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(argNode))){
                alreadySet.Add(Nodes_Node2Stem(argNode))
                if _switch0 := visitor(_env, argNode, &self.Nodes_Node, "arg", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(argNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        for _, _retTypeNode := range( self.declInfo.FP.Get_retTypeNodeList(_env).Items ) {
            retTypeNode := _retTypeNode.(Nodes_RefTypeNodeDownCast).ToNodes_RefTypeNode()
            if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(retTypeNode))){
                alreadySet.Add(Nodes_RefTypeNode2Stem(retTypeNode))
                if _switch1 := visitor(_env, &retTypeNode.Nodes_Node, &self.Nodes_Node, "retType", depth); _switch1 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(retTypeNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch1 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch1 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        {
            _body := self.declInfo.FP.Get_body(_env)
            if !Lns_IsNil( _body ) {
                body := _body.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(body))){
                    alreadySet.Add(Nodes_BlockNode2Stem(body))
                    if _switch2 := visitor(_env, &body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch2 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(body.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch2 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch2 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclFuncNode.processFilter
func (self *Nodes_DeclFuncNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclFunc(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclFuncNode.canBeLeft
func (self *Nodes_DeclFuncNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.DeclFuncNode.create
func Nodes_DeclFuncNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_DeclFuncNode {
    var node *Nodes_DeclFuncNode
    node = NewNodes_DeclFuncNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclFuncNode.visit
func (self *Nodes_DeclFuncNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        for _, _argNode := range( self.declInfo.FP.Get_argList(_env).Items ) {
            argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(argNode))){
                alreadySet.Add(Nodes_Node2Stem(argNode))
                if _switch0 := visitor(_env, argNode, &self.Nodes_Node, "arg", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(argNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        for _, _retTypeNode := range( self.declInfo.FP.Get_retTypeNodeList(_env).Items ) {
            retTypeNode := _retTypeNode.(Nodes_RefTypeNodeDownCast).ToNodes_RefTypeNode()
            if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(retTypeNode))){
                alreadySet.Add(Nodes_RefTypeNode2Stem(retTypeNode))
                if _switch1 := visitor(_env, &retTypeNode.Nodes_Node, &self.Nodes_Node, "retType", depth); _switch1 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(retTypeNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch1 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch1 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        {
            _body := self.declInfo.FP.Get_body(_env)
            if !Lns_IsNil( _body ) {
                body := _body.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(body))){
                    alreadySet.Add(Nodes_BlockNode2Stem(body))
                    if _switch2 := visitor(_env, &body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch2 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(body.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch2 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch2 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1813: decl @lune.@base.@Nodes.DeclFuncNode.canBeRight
func (self *Nodes_DeclFuncNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return Types_Token2Stem(self.declInfo.FP.Get_name(_env)) == nil
}
// 1817: decl @lune.@base.@Nodes.DeclFuncNode.canBeStatement
func (self *Nodes_DeclFuncNode) CanBeStatement(_env *LnsEnv) bool {
    return Lns_op_not((Types_Token2Stem(self.declInfo.FP.Get_name(_env)) == nil))
}
// 1: decl @lune.@base.@Nodes.DeclMethodNode.processFilter
func (self *Nodes_DeclMethodNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclMethod(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclMethodNode.canBeRight
func (self *Nodes_DeclMethodNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclMethodNode.canBeLeft
func (self *Nodes_DeclMethodNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclMethodNode.canBeStatement
func (self *Nodes_DeclMethodNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclMethodNode.create
func Nodes_DeclMethodNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_DeclMethodNode {
    var node *Nodes_DeclMethodNode
    node = NewNodes_DeclMethodNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclMethodNode.visit
func (self *Nodes_DeclMethodNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        for _, _argNode := range( self.declInfo.FP.Get_argList(_env).Items ) {
            argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(argNode))){
                alreadySet.Add(Nodes_Node2Stem(argNode))
                if _switch0 := visitor(_env, argNode, &self.Nodes_Node, "arg", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(argNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        for _, _retTypeNode := range( self.declInfo.FP.Get_retTypeNodeList(_env).Items ) {
            retTypeNode := _retTypeNode.(Nodes_RefTypeNodeDownCast).ToNodes_RefTypeNode()
            if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(retTypeNode))){
                alreadySet.Add(Nodes_RefTypeNode2Stem(retTypeNode))
                if _switch1 := visitor(_env, &retTypeNode.Nodes_Node, &self.Nodes_Node, "retType", depth); _switch1 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(retTypeNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch1 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch1 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        {
            _body := self.declInfo.FP.Get_body(_env)
            if !Lns_IsNil( _body ) {
                body := _body.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(body))){
                    alreadySet.Add(Nodes_BlockNode2Stem(body))
                    if _switch2 := visitor(_env, &body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch2 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(body.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch2 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch2 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ProtoMethodNode.processFilter
func (self *Nodes_ProtoMethodNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessProtoMethod(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ProtoMethodNode.canBeRight
func (self *Nodes_ProtoMethodNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ProtoMethodNode.canBeLeft
func (self *Nodes_ProtoMethodNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ProtoMethodNode.canBeStatement
func (self *Nodes_ProtoMethodNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ProtoMethodNode.create
func Nodes_ProtoMethodNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_ProtoMethodNode {
    var node *Nodes_ProtoMethodNode
    node = NewNodes_ProtoMethodNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ProtoMethodNode.visit
func (self *Nodes_ProtoMethodNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        for _, _argNode := range( self.declInfo.FP.Get_argList(_env).Items ) {
            argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(argNode))){
                alreadySet.Add(Nodes_Node2Stem(argNode))
                if _switch0 := visitor(_env, argNode, &self.Nodes_Node, "arg", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(argNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        for _, _retTypeNode := range( self.declInfo.FP.Get_retTypeNodeList(_env).Items ) {
            retTypeNode := _retTypeNode.(Nodes_RefTypeNodeDownCast).ToNodes_RefTypeNode()
            if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(retTypeNode))){
                alreadySet.Add(Nodes_RefTypeNode2Stem(retTypeNode))
                if _switch1 := visitor(_env, &retTypeNode.Nodes_Node, &self.Nodes_Node, "retType", depth); _switch1 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(retTypeNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch1 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch1 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        {
            _body := self.declInfo.FP.Get_body(_env)
            if !Lns_IsNil( _body ) {
                body := _body.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(body))){
                    alreadySet.Add(Nodes_BlockNode2Stem(body))
                    if _switch2 := visitor(_env, &body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch2 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(body.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch2 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch2 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclConstrNode.processFilter
func (self *Nodes_DeclConstrNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclConstr(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclConstrNode.canBeRight
func (self *Nodes_DeclConstrNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclConstrNode.canBeLeft
func (self *Nodes_DeclConstrNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclConstrNode.canBeStatement
func (self *Nodes_DeclConstrNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclConstrNode.create
func Nodes_DeclConstrNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_DeclConstrNode {
    var node *Nodes_DeclConstrNode
    node = NewNodes_DeclConstrNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclConstrNode.visit
func (self *Nodes_DeclConstrNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        for _, _argNode := range( self.declInfo.FP.Get_argList(_env).Items ) {
            argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(argNode))){
                alreadySet.Add(Nodes_Node2Stem(argNode))
                if _switch0 := visitor(_env, argNode, &self.Nodes_Node, "arg", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(argNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        for _, _retTypeNode := range( self.declInfo.FP.Get_retTypeNodeList(_env).Items ) {
            retTypeNode := _retTypeNode.(Nodes_RefTypeNodeDownCast).ToNodes_RefTypeNode()
            if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(retTypeNode))){
                alreadySet.Add(Nodes_RefTypeNode2Stem(retTypeNode))
                if _switch1 := visitor(_env, &retTypeNode.Nodes_Node, &self.Nodes_Node, "retType", depth); _switch1 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(retTypeNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch1 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch1 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        {
            _body := self.declInfo.FP.Get_body(_env)
            if !Lns_IsNil( _body ) {
                body := _body.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(body))){
                    alreadySet.Add(Nodes_BlockNode2Stem(body))
                    if _switch2 := visitor(_env, &body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch2 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(body.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch2 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch2 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclDestrNode.processFilter
func (self *Nodes_DeclDestrNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclDestr(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclDestrNode.canBeRight
func (self *Nodes_DeclDestrNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclDestrNode.canBeLeft
func (self *Nodes_DeclDestrNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclDestrNode.canBeStatement
func (self *Nodes_DeclDestrNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclDestrNode.create
func Nodes_DeclDestrNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclFuncInfo) *Nodes_DeclDestrNode {
    var node *Nodes_DeclDestrNode
    node = NewNodes_DeclDestrNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclDestrNode.visit
func (self *Nodes_DeclDestrNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        for _, _argNode := range( self.declInfo.FP.Get_argList(_env).Items ) {
            argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(argNode))){
                alreadySet.Add(Nodes_Node2Stem(argNode))
                if _switch0 := visitor(_env, argNode, &self.Nodes_Node, "arg", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(argNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        for _, _retTypeNode := range( self.declInfo.FP.Get_retTypeNodeList(_env).Items ) {
            retTypeNode := _retTypeNode.(Nodes_RefTypeNodeDownCast).ToNodes_RefTypeNode()
            if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(retTypeNode))){
                alreadySet.Add(Nodes_RefTypeNode2Stem(retTypeNode))
                if _switch1 := visitor(_env, &retTypeNode.Nodes_Node, &self.Nodes_Node, "retType", depth); _switch1 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(retTypeNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch1 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch1 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
        {
            _body := self.declInfo.FP.Get_body(_env)
            if !Lns_IsNil( _body ) {
                body := _body.(*Nodes_BlockNode)
                if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(body))){
                    alreadySet.Add(Nodes_BlockNode2Stem(body))
                    if _switch2 := visitor(_env, &body.Nodes_Node, &self.Nodes_Node, "declInfo", depth); _switch2 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(body.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch2 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch2 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.processFilter
func (self *Nodes_ExpCallSuperCtorNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCallSuperCtor(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.canBeRight
func (self *Nodes_ExpCallSuperCtorNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.canBeLeft
func (self *Nodes_ExpCallSuperCtorNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.canBeStatement
func (self *Nodes_ExpCallSuperCtorNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.create
func Nodes_ExpCallSuperCtorNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,superType *Ast_TypeInfo,methodType *Ast_TypeInfo,expList LnsAny) *Nodes_ExpCallSuperCtorNode {
    var node *Nodes_ExpCallSuperCtorNode
    node = NewNodes_ExpCallSuperCtorNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, superType, methodType, expList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpCallSuperCtorNode.visit
func (self *Nodes_ExpCallSuperCtorNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.expList
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.ExpCallSuperNode.processFilter
func (self *Nodes_ExpCallSuperNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessExpCallSuper(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ExpCallSuperNode.canBeLeft
func (self *Nodes_ExpCallSuperNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ExpCallSuperNode.canBeStatement
func (self *Nodes_ExpCallSuperNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ExpCallSuperNode.create
func Nodes_ExpCallSuperNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,superType *Ast_TypeInfo,methodType *Ast_TypeInfo,expList LnsAny) *Nodes_ExpCallSuperNode {
    var node *Nodes_ExpCallSuperNode
    node = NewNodes_ExpCallSuperNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, superType, methodType, expList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ExpCallSuperNode.visit
func (self *Nodes_ExpCallSuperNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.expList
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1851: decl @lune.@base.@Nodes.ExpCallSuperNode.canBeRight
func (self *Nodes_ExpCallSuperNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return self.FP.Get_expType(_env) != Ast_builtinTypeNone
}
// 1: decl @lune.@base.@Nodes.AsyncLockNode.processFilter
func (self *Nodes_AsyncLockNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessAsyncLock(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.AsyncLockNode.canBeRight
func (self *Nodes_AsyncLockNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.AsyncLockNode.canBeLeft
func (self *Nodes_AsyncLockNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.AsyncLockNode.canBeStatement
func (self *Nodes_AsyncLockNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.AsyncLockNode.create
func Nodes_AsyncLockNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,lockKind LnsInt,block *Nodes_BlockNode) *Nodes_AsyncLockNode {
    var node *Nodes_AsyncLockNode
    node = NewNodes_AsyncLockNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, lockKind, block)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.AsyncLockNode.visit
func (self *Nodes_AsyncLockNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.RequestNode.processFilter
func (self *Nodes_RequestNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessRequest(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.RequestNode.canBeRight
func (self *Nodes_RequestNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.RequestNode.canBeLeft
func (self *Nodes_RequestNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.RequestNode.canBeStatement
func (self *Nodes_RequestNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.RequestNode.create
func Nodes_RequestNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,processor *Nodes_Node,exp *Nodes_Node) *Nodes_RequestNode {
    var node *Nodes_RequestNode
    node = NewNodes_RequestNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, processor, exp)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.RequestNode.visit
func (self *Nodes_RequestNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.processor
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "processor", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch1 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclMemberNode.processFilter
func (self *Nodes_DeclMemberNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclMember(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclMemberNode.canBeRight
func (self *Nodes_DeclMemberNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclMemberNode.canBeLeft
func (self *Nodes_DeclMemberNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclMemberNode.canBeStatement
func (self *Nodes_DeclMemberNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclMemberNode.create
func Nodes_DeclMemberNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,refType *Nodes_RefTypeNode,symbolInfo *Ast_SymbolInfo,classType *Ast_TypeInfo,staticFlag bool,accessMode LnsInt,getterMutable bool,getterMode LnsInt,getterToken LnsAny,getterRetType *Ast_TypeInfo,setterMode LnsInt,setterToken LnsAny) *Nodes_DeclMemberNode {
    var node *Nodes_DeclMemberNode
    node = NewNodes_DeclMemberNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, name, refType, symbolInfo, classType, staticFlag, accessMode, getterMutable, getterMode, getterToken, getterRetType, setterMode, setterToken)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclMemberNode.visit
func (self *Nodes_DeclMemberNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_RefTypeNode
        child = self.refType
        if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(child))){
            alreadySet.Add(Nodes_RefTypeNode2Stem(child))
            if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "refType", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1890: decl @lune.@base.@Nodes.DeclMemberNode.getGetterSym
func (self *Nodes_DeclMemberNode) GetGetterSym(_env *LnsEnv) LnsAny {
    if self.getterMode != Ast_AccessMode__None{
        return _env.NilAccFin(_env.NilAccPush(self.classType.FP.Get_scope(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(_env, _env.GetVM().String_format("get_%s", []LnsAny{self.name.Txt}))})/* 1892:14 */)
    }
    return nil
}
// 1897: decl @lune.@base.@Nodes.DeclMemberNode.getSetterSym
func (self *Nodes_DeclMemberNode) GetSetterSym(_env *LnsEnv) LnsAny {
    if self.setterMode != Ast_AccessMode__None{
        return _env.NilAccFin(_env.NilAccPush(self.classType.FP.Get_scope(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(_env, _env.GetVM().String_format("set_%s", []LnsAny{self.name.Txt}))})/* 1899:14 */)
    }
    return nil
}
// 1: decl @lune.@base.@Nodes.DeclArgNode.processFilter
func (self *Nodes_DeclArgNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclArg(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclArgNode.canBeRight
func (self *Nodes_DeclArgNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclArgNode.canBeLeft
func (self *Nodes_DeclArgNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclArgNode.canBeStatement
func (self *Nodes_DeclArgNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.DeclArgNode.create
func Nodes_DeclArgNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,symbolInfo *Ast_SymbolInfo,argType LnsAny) *Nodes_DeclArgNode {
    var node *Nodes_DeclArgNode
    node = NewNodes_DeclArgNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, name, symbolInfo, argType)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclArgNode.visit
func (self *Nodes_DeclArgNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.argType
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_RefTypeNode)
                if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(child))){
                    alreadySet.Add(Nodes_RefTypeNode2Stem(child))
                    if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "argType", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclArgDDDNode.processFilter
func (self *Nodes_DeclArgDDDNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclArgDDD(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclArgDDDNode.canBeRight
func (self *Nodes_DeclArgDDDNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclArgDDDNode.canBeLeft
func (self *Nodes_DeclArgDDDNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclArgDDDNode.canBeStatement
func (self *Nodes_DeclArgDDDNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.DeclArgDDDNode.create
func Nodes_DeclArgDDDNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) *Nodes_DeclArgDDDNode {
    var node *Nodes_DeclArgDDDNode
    node = NewNodes_DeclArgDDDNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclArgDDDNode.visit
func (self *Nodes_DeclArgDDDNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclAdvertiseNode.processFilter
func (self *Nodes_DeclAdvertiseNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclAdvertise(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclAdvertiseNode.canBeRight
func (self *Nodes_DeclAdvertiseNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclAdvertiseNode.canBeLeft
func (self *Nodes_DeclAdvertiseNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclAdvertiseNode.canBeStatement
func (self *Nodes_DeclAdvertiseNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.DeclAdvertiseNode.create
func Nodes_DeclAdvertiseNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,advInfo *Nodes_AdvertiseInfo) *Nodes_DeclAdvertiseNode {
    var node *Nodes_DeclAdvertiseNode
    node = NewNodes_DeclAdvertiseNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, advInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclAdvertiseNode.visit
func (self *Nodes_DeclAdvertiseNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1925: decl @lune.@base.@Nodes.ClassInheritInfo.visit
func (self *Nodes_ClassInheritInfo) Visit(_env *LnsEnv, parent *Nodes_Node,visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        _base := self.base
        if !Lns_IsNil( _base ) {
            base := _base.(*Nodes_RefTypeNode)
            if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(base))){
                alreadySet.Add(Nodes_RefTypeNode2Stem(base))
                if _switch0 := visitor(_env, &base.Nodes_Node, parent, "base", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(base.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    for _, _ifTypeNode := range( self.impliments.Items ) {
        ifTypeNode := _ifTypeNode.(Nodes_RefTypeNodeDownCast).ToNodes_RefTypeNode()
        if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(ifTypeNode))){
            alreadySet.Add(Nodes_RefTypeNode2Stem(ifTypeNode))
            if _switch1 := visitor(_env, &ifTypeNode.Nodes_Node, parent, "ifTypeNode", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(ifTypeNode.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return true
}
// 1: decl @lune.@base.@Nodes.ProtoClassNode.processFilter
func (self *Nodes_ProtoClassNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessProtoClass(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.ProtoClassNode.canBeRight
func (self *Nodes_ProtoClassNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ProtoClassNode.canBeLeft
func (self *Nodes_ProtoClassNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.ProtoClassNode.canBeStatement
func (self *Nodes_ProtoClassNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.ProtoClassNode.create
func Nodes_ProtoClassNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,inheritInfo *Nodes_ClassInheritInfo) *Nodes_ProtoClassNode {
    var node *Nodes_ProtoClassNode
    node = NewNodes_ProtoClassNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, name, inheritInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.ProtoClassNode.visit
func (self *Nodes_ProtoClassNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1942: decl @lune.@base.@Nodes.ProtoClassNode.visitSub
func (self *Nodes_ProtoClassNode) VisitSub(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.inheritInfo.FP.Visit(_env, &self.Nodes_Node, visitor, depth, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclClassNode.processFilter
func (self *Nodes_DeclClassNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclClass(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclClassNode.canBeRight
func (self *Nodes_DeclClassNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclClassNode.canBeLeft
func (self *Nodes_DeclClassNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclClassNode.canBeStatement
func (self *Nodes_DeclClassNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclClassNode.create
func Nodes_DeclClassNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,accessMode LnsInt,name *Types_Token,inheritInfo *Nodes_ClassInheritInfo,hasPrototype bool,gluePrefix LnsAny,moduleName LnsAny,lang LnsAny,lazyLoad LnsInt,hasOldCtor bool,allStmtList *LnsList,declStmtList *LnsList,fieldList *LnsList,memberList *LnsList,scope *Ast_Scope,initBlock *Nodes_ClassInitBlockInfo,advertiseList *LnsList,trustList *LnsList,uninitMemberList *LnsList,outerMethodSet *LnsSet) *Nodes_DeclClassNode {
    var node *Nodes_DeclClassNode
    node = NewNodes_DeclClassNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, accessMode, name, inheritInfo, hasPrototype, gluePrefix, moduleName, lang, lazyLoad, hasOldCtor, allStmtList, declStmtList, fieldList, memberList, scope, initBlock, advertiseList, trustList, uninitMemberList, outerMethodSet)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclClassNode.visit
func (self *Nodes_DeclClassNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var list *LnsList
        list = self.allStmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch0 := visitor(_env, child, &self.Nodes_Node, "allStmtList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    {
        var list *LnsList
        list = self.declStmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch1 := visitor(_env, child, &self.Nodes_Node, "declStmtList", depth); _switch1 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch1 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch1 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    {
        var list *LnsList
        list = self.fieldList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch2 := visitor(_env, child, &self.Nodes_Node, "fieldList", depth); _switch2 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch2 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch2 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    {
        var list *LnsList
        list = self.memberList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if Lns_op_not(alreadySet.Has(Nodes_DeclMemberNode2Stem(child))){
                alreadySet.Add(Nodes_DeclMemberNode2Stem(child))
                if _switch3 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "memberList", depth); _switch3 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch3 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch3 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1993: decl @lune.@base.@Nodes.DeclClassNode.visitSub
func (self *Nodes_DeclClassNode) VisitSub(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.inheritInfo.FP.Visit(_env, &self.Nodes_Node, visitor, depth, alreadySet)
}
// 1999: decl @lune.@base.@Nodes.DeclClassNode.isModule
func (self *Nodes_DeclClassNode) IsModule(_env *LnsEnv) bool {
    return Types_Token2Stem(self.moduleName) != nil
}
// 2006: decl @lune.@base.@Nodes.DeclClassNode.createMethodNameSetWithoutAdv
func (self *Nodes_DeclClassNode) CreateMethodNameSetWithoutAdv(_env *LnsEnv) *LnsSet {
    var methodNameSet *LnsSet
    methodNameSet = NewLnsSet([]LnsAny{})
    if self.FP.Get_expType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__IF{
        for _, _field := range( self.FP.Get_fieldList(_env).Items ) {
            field := _field.(Nodes_NodeDownCast).ToNodes_Node()
            if field.FP.Get_kind(_env) == Nodes_NodeKind_get_DeclConstr(_env){
                methodNameSet.Add("__init")
            }
            if field.FP.Get_kind(_env) == Nodes_NodeKind_get_DeclDestr(_env){
                methodNameSet.Add("__free")
            }
            {
                _methodNode := Nodes_DeclMethodNodeDownCastF(field.FP)
                if !Lns_IsNil( _methodNode ) {
                    methodNode := _methodNode.(*Nodes_DeclMethodNode)
                    var methodNameToken *Types_Token
                    methodNameToken = Lns_unwrap( methodNode.FP.Get_declInfo(_env).FP.Get_name(_env)).(*Types_Token)
                    methodNameSet.Add(methodNameToken.Txt)
                }
            }
        }
        for _, _memberNode := range( self.FP.Get_memberList(_env).Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            var memberName string
            memberName = memberNode.FP.Get_name(_env).Txt
            if memberNode.FP.Get_getterMode(_env) != Ast_AccessMode__None{
                methodNameSet.Add("get_" + memberName)
            }
            if memberNode.FP.Get_setterMode(_env) != Ast_AccessMode__None{
                methodNameSet.Add("set_" + memberName)
            }
        }
    }
    return methodNameSet
}
// 2034: decl @lune.@base.@Nodes.DeclClassNode.setHasOldCtor
func (self *Nodes_DeclClassNode) SetHasOldCtor(_env *LnsEnv) {
    self.hasOldCtor = true
}
// 2038: decl @lune.@base.@Nodes.DeclClassNode.hasUserInit
func (self *Nodes_DeclClassNode) HasUserInit(_env *LnsEnv) bool {
    var scope *Ast_Scope
    scope = Lns_unwrap( self.FP.Get_expType(_env).FP.Get_scope(_env)).(*Ast_Scope)
    var initFuncType *Ast_TypeInfo
    initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField(_env, "__init", true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
    return Lns_op_not(initFuncType.FP.Get_autoFlag(_env))
}
// 1: decl @lune.@base.@Nodes.DeclEnumNode.processFilter
func (self *Nodes_DeclEnumNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclEnum(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclEnumNode.canBeRight
func (self *Nodes_DeclEnumNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclEnumNode.canBeLeft
func (self *Nodes_DeclEnumNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclEnumNode.canBeStatement
func (self *Nodes_DeclEnumNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclEnumNode.create
func Nodes_DeclEnumNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,enumType *Ast_EnumTypeInfo,accessMode LnsInt,name *Types_Token,valueNameList *LnsList,scope *Ast_Scope) *Nodes_DeclEnumNode {
    var node *Nodes_DeclEnumNode
    node = NewNodes_DeclEnumNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, enumType, accessMode, name, valueNameList, scope)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclEnumNode.visit
func (self *Nodes_DeclEnumNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclAlgeNode.processFilter
func (self *Nodes_DeclAlgeNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclAlge(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclAlgeNode.canBeRight
func (self *Nodes_DeclAlgeNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclAlgeNode.canBeLeft
func (self *Nodes_DeclAlgeNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclAlgeNode.canBeStatement
func (self *Nodes_DeclAlgeNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclAlgeNode.create
func Nodes_DeclAlgeNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,accessMode LnsInt,algeType *Ast_AlgeTypeInfo,name *Types_Token,algeValList *LnsList,scope *Ast_Scope) *Nodes_DeclAlgeNode {
    var node *Nodes_DeclAlgeNode
    node = NewNodes_DeclAlgeNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, accessMode, algeType, name, algeValList, scope)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclAlgeNode.visit
func (self *Nodes_DeclAlgeNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2068: decl @lune.@base.@Nodes.DeclAlgeNode.visitSub
func (self *Nodes_DeclAlgeNode) VisitSub(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    for _, _valInfo := range( self.algeValList.Items ) {
        valInfo := _valInfo.(Nodes_DeclAlgeValInfoDownCast).ToNodes_DeclAlgeValInfo()
        for _, _paramInfo := range( valInfo.FP.Get_paramList(_env).Items ) {
            paramInfo := _paramInfo.(Nodes_AlgeValParamInfoDownCast).ToNodes_AlgeValParamInfo()
            if Lns_op_not(alreadySet.Has(Nodes_RefTypeNode2Stem(paramInfo.FP.Get_typeRef(_env)))){
                alreadySet.Add(Nodes_RefTypeNode2Stem(paramInfo.FP.Get_typeRef(_env)))
                if _switch0 := visitor(_env, &paramInfo.FP.Get_typeRef(_env).Nodes_Node, &self.Nodes_Node, "typeRef", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(paramInfo.FP.Get_typeRef(_env).FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    return true
}
// 1: decl @lune.@base.@Nodes.NewAlgeValNode.processFilter
func (self *Nodes_NewAlgeValNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessNewAlgeVal(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.NewAlgeValNode.canBeRight
func (self *Nodes_NewAlgeValNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.NewAlgeValNode.canBeLeft
func (self *Nodes_NewAlgeValNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.NewAlgeValNode.canBeStatement
func (self *Nodes_NewAlgeValNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.NewAlgeValNode.create
func Nodes_NewAlgeValNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,prefix LnsAny,algeTypeInfo *Ast_AlgeTypeInfo,valInfo *Ast_AlgeValInfo,paramList *LnsList) *Nodes_NewAlgeValNode {
    var node *Nodes_NewAlgeValNode
    node = NewNodes_NewAlgeValNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, name, prefix, algeTypeInfo, valInfo, paramList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.NewAlgeValNode.visit
func (self *Nodes_NewAlgeValNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.prefix
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_Node)
                if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                    alreadySet.Add(Nodes_Node2Stem(child))
                    if _switch0 := visitor(_env, child, &self.Nodes_Node, "prefix", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    {
        var list *LnsList
        list = self.paramList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch1 := visitor(_env, child, &self.Nodes_Node, "paramList", depth); _switch1 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch1 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch1 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.LuneControlNode.processFilter
func (self *Nodes_LuneControlNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLuneControl(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LuneControlNode.canBeRight
func (self *Nodes_LuneControlNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LuneControlNode.canBeLeft
func (self *Nodes_LuneControlNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LuneControlNode.canBeStatement
func (self *Nodes_LuneControlNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.LuneControlNode.create
func Nodes_LuneControlNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,pragma LnsAny) *Nodes_LuneControlNode {
    var node *Nodes_LuneControlNode
    node = NewNodes_LuneControlNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, pragma)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LuneControlNode.visit
func (self *Nodes_LuneControlNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.MatchNode.processFilter
func (self *Nodes_MatchNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessMatch(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.MatchNode.canBeRight
func (self *Nodes_MatchNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.MatchNode.canBeLeft
func (self *Nodes_MatchNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.MatchNode.canBeStatement
func (self *Nodes_MatchNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.MatchNode.create
func Nodes_MatchNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,idInNS LnsInt,val *Nodes_Node,algeTypeInfo *Ast_AlgeTypeInfo,caseList *LnsList,defaultBlock LnsAny,caseKind LnsInt,failSafeDefault bool) *Nodes_MatchNode {
    var node *Nodes_MatchNode
    node = NewNodes_MatchNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, idInNS, val, algeTypeInfo, caseList, defaultBlock, caseKind, failSafeDefault)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.MatchNode.visit
func (self *Nodes_MatchNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.val
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "val", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        {
            _child := self.defaultBlock
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_Node)
                if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                    alreadySet.Add(Nodes_Node2Stem(child))
                    if _switch1 := visitor(_env, child, &self.Nodes_Node, "defaultBlock", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2113: decl @lune.@base.@Nodes.MatchNode.getBreakKind
func (self *Nodes_MatchNode) GetBreakKind(_env *LnsEnv, checkMode LnsInt) LnsInt {
    var kind LnsInt
    kind = Nodes_BreakKind__None
    var fullCase bool
    fullCase = self.caseKind != Nodes_CaseKind__Lack
    for _, _caseInfo := range( self.caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        var work LnsInt
        work = caseInfo.FP.Get_block(_env).FP.GetBreakKind(_env, checkMode)
        var goNext bool
        goNext = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( (work == Nodes_BreakKind__None)) ||
            _env.SetStackVal( Lns_op_not(fullCase)) ).(bool)
        if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
            if work == Nodes_BreakKind__Return{
                return Nodes_BreakKind__Return
            }
            if work == Nodes_BreakKind__NeverRet{
                return Nodes_BreakKind__NeverRet
            }
        } else { 
            if _switch0 := work; _switch0 == Nodes_BreakKind__None {
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                    _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                    if goNext{
                        return Nodes_BreakKind__None
                    }
                }
            } else {
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                    _env.SetStackVal( kind > work) ).(bool){
                    kind = work
                }
            }
        }
        
    }
    {
        _block := self.defaultBlock
        if !Lns_IsNil( _block ) {
            block := _block.(*Nodes_Node)
            var work LnsInt
            work = block.FP.GetBreakKind(_env, checkMode)
            if checkMode == Nodes_CheckBreakMode__IgnoreFlowReturn{
                if work == Nodes_BreakKind__Return{
                    return Nodes_BreakKind__Return
                }
                if work == Nodes_BreakKind__NeverRet{
                    return Nodes_BreakKind__NeverRet
                }
            } else { 
                if _switch1 := work; _switch1 == Nodes_BreakKind__None {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Normal) ||
                        _env.SetStackVal( checkMode == Nodes_CheckBreakMode__Return) ).(bool){
                        if true{
                            return Nodes_BreakKind__None
                        }
                    }
                } else {
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( kind == Nodes_BreakKind__None) ||
                        _env.SetStackVal( kind > work) ).(bool){
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
// 2134: decl @lune.@base.@Nodes.MatchNode.visitSub
func (self *Nodes_MatchNode) VisitSub(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    for _, _caseInfo := range( self.caseList.Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        if Lns_op_not(alreadySet.Has(Nodes_ExpRefNode2Stem(caseInfo.FP.Get_valExpRef(_env)))){
            alreadySet.Add(Nodes_ExpRefNode2Stem(caseInfo.FP.Get_valExpRef(_env)))
            if _switch0 := visitor(_env, &caseInfo.FP.Get_valExpRef(_env).Nodes_Node, &self.Nodes_Node, "valExpRef", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(caseInfo.FP.Get_valExpRef(_env).FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(caseInfo.FP.Get_block(_env)))){
            alreadySet.Add(Nodes_BlockNode2Stem(caseInfo.FP.Get_block(_env)))
            if _switch1 := visitor(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, &self.Nodes_Node, "block", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(caseInfo.FP.Get_block(_env).FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return true
}
// 1: decl @lune.@base.@Nodes.LuneKindNode.processFilter
func (self *Nodes_LuneKindNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLuneKind(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LuneKindNode.canBeRight
func (self *Nodes_LuneKindNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LuneKindNode.canBeLeft
func (self *Nodes_LuneKindNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LuneKindNode.canBeStatement
func (self *Nodes_LuneKindNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LuneKindNode.create
func Nodes_LuneKindNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,exp *Nodes_Node) *Nodes_LuneKindNode {
    var node *Nodes_LuneKindNode
    node = NewNodes_LuneKindNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, exp)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LuneKindNode.visit
func (self *Nodes_LuneKindNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.exp
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "exp", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.DeclMacroNode.processFilter
func (self *Nodes_DeclMacroNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessDeclMacro(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.DeclMacroNode.canBeRight
func (self *Nodes_DeclMacroNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclMacroNode.canBeLeft
func (self *Nodes_DeclMacroNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.DeclMacroNode.canBeStatement
func (self *Nodes_DeclMacroNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.DeclMacroNode.create
func Nodes_DeclMacroNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,declInfo *Nodes_DeclMacroInfo) *Nodes_DeclMacroNode {
    var node *Nodes_DeclMacroNode
    node = NewNodes_DeclMacroNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, declInfo)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.DeclMacroNode.visit
func (self *Nodes_DeclMacroNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.TestCaseNode.processFilter
func (self *Nodes_TestCaseNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessTestCase(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.TestCaseNode.canBeRight
func (self *Nodes_TestCaseNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.TestCaseNode.canBeLeft
func (self *Nodes_TestCaseNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.TestCaseNode.canBeStatement
func (self *Nodes_TestCaseNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.TestCaseNode.create
func Nodes_TestCaseNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,name *Types_Token,impNode *Nodes_Node,ctrlName string,block *Nodes_BlockNode) *Nodes_TestCaseNode {
    var node *Nodes_TestCaseNode
    node = NewNodes_TestCaseNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, name, impNode, ctrlName, block)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.TestCaseNode.visit
func (self *Nodes_TestCaseNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.impNode
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "impNode", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    {
        var child *Nodes_BlockNode
        child = self.block
        if Lns_op_not(alreadySet.Has(Nodes_BlockNode2Stem(child))){
            alreadySet.Add(Nodes_BlockNode2Stem(child))
            if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "block", depth); _switch1 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch1 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch1 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.TestBlockNode.processFilter
func (self *Nodes_TestBlockNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessTestBlock(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.TestBlockNode.canBeRight
func (self *Nodes_TestBlockNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.TestBlockNode.canBeLeft
func (self *Nodes_TestBlockNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.TestBlockNode.canBeStatement
func (self *Nodes_TestBlockNode) CanBeStatement(_env *LnsEnv) bool {
    return true
}
// 740: decl @lune.@base.@Nodes.TestBlockNode.create
func Nodes_TestBlockNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,stmtList *LnsList) *Nodes_TestBlockNode {
    var node *Nodes_TestBlockNode
    node = NewNodes_TestBlockNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, stmtList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.TestBlockNode.visit
func (self *Nodes_TestBlockNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var list *LnsList
        list = self.stmtList
        for _, _child := range( list.Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                alreadySet.Add(Nodes_Node2Stem(child))
                if _switch0 := visitor(_env, child, &self.Nodes_Node, "stmtList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                    if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                        return false
                    }
                } else if _switch0 == Nodes_NodeVisitMode__End {
                    return false
                } else if _switch0 == Nodes_NodeVisitMode__Next {
                }
            }
            
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2183: decl @lune.@base.@Nodes.TestBlockNode.isInnerPos
func (self *Nodes_TestBlockNode) IsInnerPos(_env *LnsEnv, pos Types_Position) bool {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.Get_pos(_env).StreamName == pos.StreamName) &&
        _env.SetStackVal( self.FP.Get_pos(_env).LineNo < pos.LineNo) &&
        _env.SetStackVal( self.FP.Get_stmtList(_env).Len() > 0) &&
        _env.SetStackVal( self.FP.Get_stmtList(_env).GetAt(1).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_pos(_env).LineNo >= pos.LineNo) ).(bool)){
        return true
    }
    return false
}
// 1: decl @lune.@base.@Nodes.AbbrNode.processFilter
func (self *Nodes_AbbrNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessAbbr(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.AbbrNode.canBeRight
func (self *Nodes_AbbrNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.AbbrNode.canBeLeft
func (self *Nodes_AbbrNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.AbbrNode.canBeStatement
func (self *Nodes_AbbrNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.AbbrNode.create
func Nodes_AbbrNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) *Nodes_AbbrNode {
    var node *Nodes_AbbrNode
    node = NewNodes_AbbrNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.AbbrNode.visit
func (self *Nodes_AbbrNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.BoxingNode.processFilter
func (self *Nodes_BoxingNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessBoxing(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.BoxingNode.canBeRight
func (self *Nodes_BoxingNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.BoxingNode.canBeLeft
func (self *Nodes_BoxingNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.BoxingNode.canBeStatement
func (self *Nodes_BoxingNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.BoxingNode.create
func Nodes_BoxingNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,src *Nodes_Node) *Nodes_BoxingNode {
    var node *Nodes_BoxingNode
    node = NewNodes_BoxingNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, src)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.BoxingNode.visit
func (self *Nodes_BoxingNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.src
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "src", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.UnboxingNode.processFilter
func (self *Nodes_UnboxingNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessUnboxing(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.UnboxingNode.canBeRight
func (self *Nodes_UnboxingNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.UnboxingNode.canBeLeft
func (self *Nodes_UnboxingNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.UnboxingNode.canBeStatement
func (self *Nodes_UnboxingNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.UnboxingNode.create
func Nodes_UnboxingNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,src *Nodes_Node) *Nodes_UnboxingNode {
    var node *Nodes_UnboxingNode
    node = NewNodes_UnboxingNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, src)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.UnboxingNode.visit
func (self *Nodes_UnboxingNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var child *Nodes_Node
        child = self.src
        if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
            alreadySet.Add(Nodes_Node2Stem(child))
            if _switch0 := visitor(_env, child, &self.Nodes_Node, "src", depth); _switch0 == Nodes_NodeVisitMode__Child {
                if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                    return false
                }
            } else if _switch0 == Nodes_NodeVisitMode__End {
                return false
            } else if _switch0 == Nodes_NodeVisitMode__Next {
            }
        }
        
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 1: decl @lune.@base.@Nodes.LiteralNilNode.processFilter
func (self *Nodes_LiteralNilNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralNil(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralNilNode.canBeRight
func (self *Nodes_LiteralNilNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralNilNode.canBeLeft
func (self *Nodes_LiteralNilNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralNilNode.canBeStatement
func (self *Nodes_LiteralNilNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralNilNode.create
func Nodes_LiteralNilNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList) *Nodes_LiteralNilNode {
    var node *Nodes_LiteralNilNode
    node = NewNodes_LiteralNilNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralNilNode.visit
func (self *Nodes_LiteralNilNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2376: decl @lune.@base.@Nodes.LiteralNilNode.getLiteral
func (self *Nodes_LiteralNilNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    return Nodes_Literal__Nil_Obj, nil
}
// 2379: decl @lune.@base.@Nodes.LiteralNilNode.setupLiteralTokenList
func (self *Nodes_LiteralNilNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Symb, "nil")
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralCharNode.processFilter
func (self *Nodes_LiteralCharNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralChar(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralCharNode.canBeRight
func (self *Nodes_LiteralCharNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralCharNode.canBeLeft
func (self *Nodes_LiteralCharNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralCharNode.canBeStatement
func (self *Nodes_LiteralCharNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralCharNode.create
func Nodes_LiteralCharNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsInt) *Nodes_LiteralCharNode {
    var node *Nodes_LiteralCharNode
    node = NewNodes_LiteralCharNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, token, num)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralCharNode.visit
func (self *Nodes_LiteralCharNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2385: decl @lune.@base.@Nodes.LiteralCharNode.getLiteral
func (self *Nodes_LiteralCharNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    return &Nodes_Literal__Int{self.num}, nil
}
// 2388: decl @lune.@base.@Nodes.LiteralCharNode.setupLiteralTokenList
func (self *Nodes_LiteralCharNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Char, _env.GetVM().String_format("%d", []LnsAny{self.num}))
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralIntNode.processFilter
func (self *Nodes_LiteralIntNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralInt(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralIntNode.canBeRight
func (self *Nodes_LiteralIntNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralIntNode.canBeLeft
func (self *Nodes_LiteralIntNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralIntNode.canBeStatement
func (self *Nodes_LiteralIntNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralIntNode.create
func Nodes_LiteralIntNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsInt) *Nodes_LiteralIntNode {
    var node *Nodes_LiteralIntNode
    node = NewNodes_LiteralIntNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, token, num)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralIntNode.visit
func (self *Nodes_LiteralIntNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2393: decl @lune.@base.@Nodes.LiteralIntNode.getLiteral
func (self *Nodes_LiteralIntNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    return &Nodes_Literal__Int{self.num}, nil
}
// 2396: decl @lune.@base.@Nodes.LiteralIntNode.setupLiteralTokenList
func (self *Nodes_LiteralIntNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Int, _env.GetVM().String_format("%d", []LnsAny{self.num}))
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralRealNode.processFilter
func (self *Nodes_LiteralRealNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralReal(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralRealNode.canBeRight
func (self *Nodes_LiteralRealNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralRealNode.canBeLeft
func (self *Nodes_LiteralRealNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralRealNode.canBeStatement
func (self *Nodes_LiteralRealNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralRealNode.create
func Nodes_LiteralRealNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token,num LnsReal) *Nodes_LiteralRealNode {
    var node *Nodes_LiteralRealNode
    node = NewNodes_LiteralRealNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, token, num)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralRealNode.visit
func (self *Nodes_LiteralRealNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2401: decl @lune.@base.@Nodes.LiteralRealNode.getLiteral
func (self *Nodes_LiteralRealNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    return &Nodes_Literal__Real{self.num}, nil
}
// 2404: decl @lune.@base.@Nodes.LiteralRealNode.setupLiteralTokenList
func (self *Nodes_LiteralRealNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Real, _env.GetVM().String_format("%g", []LnsAny{self.num}))
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralArrayNode.processFilter
func (self *Nodes_LiteralArrayNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralArray(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralArrayNode.canBeRight
func (self *Nodes_LiteralArrayNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralArrayNode.canBeLeft
func (self *Nodes_LiteralArrayNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralArrayNode.canBeStatement
func (self *Nodes_LiteralArrayNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralArrayNode.create
func Nodes_LiteralArrayNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList LnsAny) *Nodes_LiteralArrayNode {
    var node *Nodes_LiteralArrayNode
    node = NewNodes_LiteralArrayNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralArrayNode.visit
func (self *Nodes_LiteralArrayNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.expList
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2409: decl @lune.@base.@Nodes.LiteralArrayNode.getLiteral
func (self *Nodes_LiteralArrayNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var literalList *LnsList
    literalList = NewLnsList([]LnsAny{})
    {
        __exp := self.expList
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            for _, _node := range( _exp.FP.Get_expList(_env).Items ) {
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                var literal LnsAny
                var mess LnsAny
                literal,mess = node.FP.GetLiteral(_env)
                if literal != nil{
                    literal_7128 := literal
                    literalList.Insert(literal_7128)
                } else {
                    return nil, mess
                }
            }
        }
    }
    return &Nodes_Literal__ARRAY{literalList}, nil
}
// 2425: decl @lune.@base.@Nodes.LiteralArrayNode.setupLiteralTokenList
func (self *Nodes_LiteralArrayNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, "[@")
    {
        __exp := self.expList
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            for _index, _node := range( _exp.FP.Get_expList(_env).Items ) {
                index := _index + 1
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, ",")
                }
                if Lns_op_not(node.FP.SetupLiteralTokenList(_env, list)){
                    return false
                }
            }
        }
    }
    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, "]")
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralListNode.processFilter
func (self *Nodes_LiteralListNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralList(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralListNode.canBeRight
func (self *Nodes_LiteralListNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralListNode.canBeLeft
func (self *Nodes_LiteralListNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralListNode.canBeStatement
func (self *Nodes_LiteralListNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralListNode.create
func Nodes_LiteralListNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList LnsAny) *Nodes_LiteralListNode {
    var node *Nodes_LiteralListNode
    node = NewNodes_LiteralListNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralListNode.visit
func (self *Nodes_LiteralListNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.expList
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2442: decl @lune.@base.@Nodes.LiteralListNode.getLiteral
func (self *Nodes_LiteralListNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var literalList *LnsList
    literalList = NewLnsList([]LnsAny{})
    {
        __exp := self.expList
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            for _, _node := range( _exp.FP.Get_expList(_env).Items ) {
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                var literal LnsAny
                var mess LnsAny
                literal,mess = node.FP.GetLiteral(_env)
                if literal != nil{
                    literal_7149 := literal
                    literalList.Insert(literal_7149)
                } else {
                    return nil, mess
                }
            }
        }
    }
    return &Nodes_Literal__LIST{literalList}, nil
}
// 2458: decl @lune.@base.@Nodes.LiteralListNode.setupLiteralTokenList
func (self *Nodes_LiteralListNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, "[")
    {
        __exp := self.expList
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            for _index, _node := range( _exp.FP.Get_expList(_env).Items ) {
                index := _index + 1
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, ",")
                }
                if Lns_op_not(node.FP.SetupLiteralTokenList(_env, list)){
                    return false
                }
            }
        }
    }
    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, "]")
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralSetNode.processFilter
func (self *Nodes_LiteralSetNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralSet(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralSetNode.canBeRight
func (self *Nodes_LiteralSetNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralSetNode.canBeLeft
func (self *Nodes_LiteralSetNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralSetNode.canBeStatement
func (self *Nodes_LiteralSetNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralSetNode.create
func Nodes_LiteralSetNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,expList LnsAny) *Nodes_LiteralSetNode {
    var node *Nodes_LiteralSetNode
    node = NewNodes_LiteralSetNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, expList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralSetNode.visit
func (self *Nodes_LiteralSetNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.expList
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "expList", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2476: decl @lune.@base.@Nodes.LiteralSetNode.getLiteral
func (self *Nodes_LiteralSetNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var literalList *LnsList
    literalList = NewLnsList([]LnsAny{})
    {
        __exp := self.expList
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            for _, _node := range( _exp.FP.Get_expList(_env).Items ) {
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                var literal LnsAny
                var mess LnsAny
                literal,mess = node.FP.GetLiteral(_env)
                if literal != nil{
                    literal_7170 := literal
                    literalList.Insert(literal_7170)
                } else {
                    return nil, mess
                }
            }
        }
    }
    return &Nodes_Literal__SET{literalList}, nil
}
// 2492: decl @lune.@base.@Nodes.LiteralSetNode.setupLiteralTokenList
func (self *Nodes_LiteralSetNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, "(@")
    {
        __exp := self.expList
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ExpListNode)
            for _index, _node := range( _exp.FP.Get_expList(_env).Items ) {
                index := _index + 1
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, ",")
                }
                if Lns_op_not(node.FP.SetupLiteralTokenList(_env, list)){
                    return false
                }
            }
        }
    }
    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, ")")
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralMapNode.processFilter
func (self *Nodes_LiteralMapNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralMap(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralMapNode.canBeRight
func (self *Nodes_LiteralMapNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralMapNode.canBeLeft
func (self *Nodes_LiteralMapNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralMapNode.canBeStatement
func (self *Nodes_LiteralMapNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralMapNode.create
func Nodes_LiteralMapNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,_map *LnsMap,pairList *LnsList) *Nodes_LiteralMapNode {
    var node *Nodes_LiteralMapNode
    node = NewNodes_LiteralMapNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, _map, pairList)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralMapNode.visit
func (self *Nodes_LiteralMapNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        var _map *LnsMap
        _map = self._map
        for _key, _val := range( _map.Items ) {
            key := _key.(Nodes_NodeDownCast).ToNodes_Node()
            val := _val.(Nodes_NodeDownCast).ToNodes_Node()
            {
                var child *Nodes_Node
                child = key
                if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                    alreadySet.Add(Nodes_Node2Stem(child))
                    if _switch0 := visitor(_env, child, &self.Nodes_Node, "map", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
            {
                var child *Nodes_Node
                child = val
                if Lns_op_not(alreadySet.Has(Nodes_Node2Stem(child))){
                    alreadySet.Add(Nodes_Node2Stem(child))
                    if _switch1 := visitor(_env, child, &self.Nodes_Node, "map", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2509: decl @lune.@base.@Nodes.LiteralMapNode.getLiteral
func (self *Nodes_LiteralMapNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var litMap *LnsMap
    litMap = NewLnsMap( map[LnsAny]LnsAny{})
    for _key, _val := range( self._map.Items ) {
        key := _key.(Nodes_NodeDownCast).ToNodes_Node()
        val := _val.(Nodes_NodeDownCast).ToNodes_Node()
        var keyLiteral LnsAny
        var keyMess LnsAny
        keyLiteral,keyMess = key.FP.GetLiteral(_env)
        var valLiteral LnsAny
        var valMess LnsAny
        valLiteral,valMess = val.FP.GetLiteral(_env)
        if keyLiteral != nil && valLiteral != nil{
            keyLiteral_7192 := keyLiteral
            valLiteral_7193 := valLiteral
            litMap.Set(keyLiteral_7192,valLiteral_7193)
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
// 2529: decl @lune.@base.@Nodes.LiteralMapNode.setupLiteralTokenList
func (self *Nodes_LiteralMapNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, "{")
    var lit2valNode *LnsMap
    lit2valNode = NewLnsMap( map[LnsAny]LnsAny{})
    for _key, _ := range( self._map.Items ) {
        key := _key.(Nodes_NodeDownCast).ToNodes_Node()
        var literal LnsAny
        literal = Nodes_convExp0_29202(Lns_2DDD(key.FP.GetLiteral(_env)))
        if literal != nil{
            literal_7205 := literal
            switch _matchExp0 := literal_7205.(type) {
            case *Nodes_Literal__Int:
            param := _matchExp0.Val1
                lit2valNode.Set(param,key)
            case *Nodes_Literal__Str:
            param := _matchExp0.Val1
                lit2valNode.Set(param,key)
            case *Nodes_Literal__Real:
            param := _matchExp0.Val1
                lit2valNode.Set(param,key)
            default:
                return false
            }
        }
    }
    {
        __forsortCollection0 := lit2valNode
        __forsortSorted0 := __forsortCollection0.CreateKeyListStem()
        __forsortSorted0.Sort( _env, LnsItemKindStem, nil )
        for _, __ := range( __forsortSorted0.Items ) {
            key := __forsortCollection0.Items[ __ ].(Nodes_NodeDownCast).ToNodes_Node()
            if Lns_op_not(key.FP.SetupLiteralTokenList(_env, list)){
                return false
            }
            self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, ":")
            if Lns_op_not(_env.NilAccFin(_env.NilAccPush(self._map.Get(key)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_Node).FP.SetupLiteralTokenList(_env, list)})/* 2559:14 */)){
                return false
            }
            self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, ",")
        }
    }
    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, "}")
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralStringNode.processFilter
func (self *Nodes_LiteralStringNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralString(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralStringNode.canBeRight
func (self *Nodes_LiteralStringNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralStringNode.canBeLeft
func (self *Nodes_LiteralStringNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralStringNode.canBeStatement
func (self *Nodes_LiteralStringNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralStringNode.create
func Nodes_LiteralStringNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token,orgParam LnsAny,dddParam LnsAny) *Nodes_LiteralStringNode {
    var node *Nodes_LiteralStringNode
    node = NewNodes_LiteralStringNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, token, orgParam, dddParam)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralStringNode.visit
func (self *Nodes_LiteralStringNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    {
        {
            _child := self.orgParam
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch0 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "orgParam", depth); _switch0 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch0 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch0 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    {
        {
            _child := self.dddParam
            if !Lns_IsNil( _child ) {
                child := _child.(*Nodes_ExpListNode)
                if Lns_op_not(alreadySet.Has(Nodes_ExpListNode2Stem(child))){
                    alreadySet.Add(Nodes_ExpListNode2Stem(child))
                    if _switch1 := visitor(_env, &child.Nodes_Node, &self.Nodes_Node, "dddParam", depth); _switch1 == Nodes_NodeVisitMode__Child {
                        if Lns_op_not(child.FP.Visit(_env, visitor, depth + 1, alreadySet)){
                            return false
                        }
                    } else if _switch1 == Nodes_NodeVisitMode__End {
                        return false
                    } else if _switch1 == Nodes_NodeVisitMode__Next {
                    }
                }
                
            }
        }
    }
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2569: decl @lune.@base.@Nodes.LiteralStringNode.getLiteral
func (self *Nodes_LiteralStringNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    var txt string
    txt = self.token.Txt
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(txt, "^```", nil, nil))){
        txt = _env.GetVM().String_sub(txt,4, -4)
    } else { 
        txt = _env.GetVM().String_sub(txt,2, -2)
    }
    {
        _param := self.FP.Get_orgParam(_env)
        if !Lns_IsNil( _param ) {
            param := _param.(*Nodes_ExpListNode)
            var argList *LnsList
            argList = param.FP.Get_expList(_env)
            var paramList *LnsList
            paramList = NewLnsList([]LnsAny{})
            for _, _argNode := range( argList.Items ) {
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                var arg LnsAny
                var mess LnsAny
                arg,mess = argNode.FP.GetLiteral(_env)
                if arg != nil{
                    arg_7232 := arg
                    paramList.Set(paramList.Len() + 1,Nodes_getLiteralObj(_env, arg_7232))
                } else {
                    return nil, mess
                }
            }
            txt = _env.GetVM().String_format(Nodes_convExp0_29460(txt, Lns_2DDD(paramList.Unpack())))
        }
    }
    return &Nodes_Literal__Str{txt}, nil
}
// 2596: decl @lune.@base.@Nodes.LiteralStringNode.setupLiteralTokenList
func (self *Nodes_LiteralStringNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Str, self.token.Txt)
    {
        _param := self.FP.Get_orgParam(_env)
        if !Lns_IsNil( _param ) {
            param := _param.(*Nodes_ExpListNode)
            self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, "(")
            for _index, _argNode := range( param.FP.Get_expList(_env).Items ) {
                index := _index + 1
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                if index > 1{
                    self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, ",")
                }
                if Lns_op_not(argNode.FP.SetupLiteralTokenList(_env, list)){
                    return false
                }
            }
            self.FP.AddTokenList(_env, list, Types_TokenKind__Dlmt, ")")
        }
    }
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralBoolNode.processFilter
func (self *Nodes_LiteralBoolNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralBool(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralBoolNode.canBeRight
func (self *Nodes_LiteralBoolNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralBoolNode.canBeLeft
func (self *Nodes_LiteralBoolNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralBoolNode.canBeStatement
func (self *Nodes_LiteralBoolNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralBoolNode.create
func Nodes_LiteralBoolNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token) *Nodes_LiteralBoolNode {
    var node *Nodes_LiteralBoolNode
    node = NewNodes_LiteralBoolNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, token)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralBoolNode.visit
func (self *Nodes_LiteralBoolNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2614: decl @lune.@base.@Nodes.LiteralBoolNode.getLiteral
func (self *Nodes_LiteralBoolNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    return &Nodes_Literal__Bool{self.token.Txt == "true"}, nil
}
// 2618: decl @lune.@base.@Nodes.LiteralBoolNode.setupLiteralTokenList
func (self *Nodes_LiteralBoolNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Kywd, self.token.Txt)
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralSymbolNode.processFilter
func (self *Nodes_LiteralSymbolNode) ProcessFilter(_env *LnsEnv, filter *Nodes_Filter,_opt LnsAny) {
    opt := _opt
    filter.FP.ProcessLiteralSymbol(_env, self, opt)
}
// 1: decl @lune.@base.@Nodes.LiteralSymbolNode.canBeRight
func (self *Nodes_LiteralSymbolNode) CanBeRight(_env *LnsEnv, processInfo *Ast_ProcessInfo) bool {
    return true
}
// 1: decl @lune.@base.@Nodes.LiteralSymbolNode.canBeLeft
func (self *Nodes_LiteralSymbolNode) CanBeLeft(_env *LnsEnv) bool {
    return false
}
// 1: decl @lune.@base.@Nodes.LiteralSymbolNode.canBeStatement
func (self *Nodes_LiteralSymbolNode) CanBeStatement(_env *LnsEnv) bool {
    return false
}
// 740: decl @lune.@base.@Nodes.LiteralSymbolNode.create
func Nodes_LiteralSymbolNode_create(_env *LnsEnv, nodeMan *Nodes_NodeManager,pos Types_Position,inTestBlock bool,macroArgFlag bool,typeList *LnsList,token *Types_Token) *Nodes_LiteralSymbolNode {
    var node *Nodes_LiteralSymbolNode
    node = NewNodes_LiteralSymbolNode(_env, nodeMan.FP.Get_managerId(_env), nodeMan.FP.NextId(_env), pos, inTestBlock, macroArgFlag, typeList, token)
    nodeMan.FP.AddNode(_env, &node.Nodes_Node)
    return node
}
// 751: decl @lune.@base.@Nodes.LiteralSymbolNode.visit
func (self *Nodes_LiteralSymbolNode) Visit(_env *LnsEnv, visitor Nodes_NodeVisitor,depth LnsInt,alreadySet *LnsSet) bool {
    return self.FP.VisitSub(_env, visitor, depth + 1, alreadySet)
}
// 2625: decl @lune.@base.@Nodes.LiteralSymbolNode.getLiteral
func (self *Nodes_LiteralSymbolNode) GetLiteral(_env *LnsEnv)(LnsAny, LnsAny) {
    return &Nodes_Literal__Symbol{self.token.Txt}, nil
}
// 2629: decl @lune.@base.@Nodes.LiteralSymbolNode.setupLiteralTokenList
func (self *Nodes_LiteralSymbolNode) SetupLiteralTokenList(_env *LnsEnv, list *LnsList) bool {
    self.FP.AddTokenList(_env, list, Types_TokenKind__Symb, self.token.Txt)
    return true
}
// 2889: decl @lune.@base.@Nodes.DefMacroInfo.get_name
func (self *Nodes_DefMacroInfo) Get_name(_env *LnsEnv) string {
    return self.DeclInfo.FP.Get_name(_env).Txt
}
// 2893: decl @lune.@base.@Nodes.DefMacroInfo.getArgList
func (self *Nodes_DefMacroInfo) GetArgList(_env *LnsEnv) *LnsList {
    return self.argList
}
// 2896: decl @lune.@base.@Nodes.DefMacroInfo.getTokenList
func (self *Nodes_DefMacroInfo) GetTokenList(_env *LnsEnv) *LnsList {
    return self.DeclInfo.FP.Get_tokenList(_env)
}
// 2920: decl @lune.@base.@Nodes.ExportInfo.assign
func (self *Nodes_ExportInfo) Assign(_env *LnsEnv, assignName string) *FrontInterface_ExportInfo {
    var info *Nodes_ExportInfo
    info = NewNodes_ExportInfo(_env, self.FP.Get_moduleTypeInfo(_env), self.FP.Get_provideInfo(_env), self.FP.Get_processInfo(_env), self.FP.Get_globalSymbolList(_env), self.FP.Get_importedAliasMap(_env), self.FP.Get_moduleId(_env), self.FP.Get_fullName(_env), assignName, self.FP.Get_streamName(_env), NewLnsMap( map[LnsAny]LnsAny{}), self.typeId2DefMacroInfo)
    info.FP.Set_importId2localTypeInfoMap(_env, self.FP.Get_importId2localTypeInfoMap(_env))
    return &info.FrontInterface_ExportInfo
}
