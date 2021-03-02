// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_convGo bool
var convGo__mod__ string
// decl enum -- ProcessMode 
type convGo_ProcessMode = LnsInt
const convGo_ProcessMode__DeclClass = 1
const convGo_ProcessMode__DeclTopScopeVar = 0
const convGo_ProcessMode__Main = 3
const convGo_ProcessMode__NonClosureFuncDecl = 2
var convGo_ProcessModeList_ = NewLnsList( []LnsAny {
  convGo_ProcessMode__DeclTopScopeVar,
  convGo_ProcessMode__DeclClass,
  convGo_ProcessMode__NonClosureFuncDecl,
  convGo_ProcessMode__Main,
})
func convGo_ProcessMode_get__allList_1029_() *LnsList{
    return convGo_ProcessModeList_
}
var convGo_ProcessModeMap_ = map[LnsInt]string {
  convGo_ProcessMode__DeclClass: "ProcessMode.DeclClass",
  convGo_ProcessMode__DeclTopScopeVar: "ProcessMode.DeclTopScopeVar",
  convGo_ProcessMode__Main: "ProcessMode.Main",
  convGo_ProcessMode__NonClosureFuncDecl: "ProcessMode.NonClosureFuncDecl",
}
func convGo_ProcessMode__from_1022_(arg1 LnsInt) LnsAny{
    if _, ok := convGo_ProcessModeMap_[arg1]; ok { return arg1 }
    return nil
}

func convGo_ProcessMode_getTxt(arg1 LnsInt) string {
    return convGo_ProcessModeMap_[arg1];
}
// decl enum -- ExpListKind 
type convGo_ExpListKind = LnsInt
const convGo_ExpListKind__Conv = 2
const convGo_ExpListKind__Direct = 0
const convGo_ExpListKind__Slice = 1
var convGo_ExpListKindList_ = NewLnsList( []LnsAny {
  convGo_ExpListKind__Direct,
  convGo_ExpListKind__Slice,
  convGo_ExpListKind__Conv,
})
func convGo_ExpListKind_get__allList_1314_() *LnsList{
    return convGo_ExpListKindList_
}
var convGo_ExpListKindMap_ = map[LnsInt]string {
  convGo_ExpListKind__Conv: "ExpListKind.Conv",
  convGo_ExpListKind__Direct: "ExpListKind.Direct",
  convGo_ExpListKind__Slice: "ExpListKind.Slice",
}
func convGo_ExpListKind__from_1307_(arg1 LnsInt) LnsAny{
    if _, ok := convGo_ExpListKindMap_[arg1]; ok { return arg1 }
    return nil
}

func convGo_ExpListKind_getTxt(arg1 LnsInt) string {
    return convGo_ExpListKindMap_[arg1];
}
var convGo_MaxNilAccNum LnsInt
var convGo_ignoreNodeInInnerBlockSet *LnsSet
var convGo_golanKeywordSet *LnsSet
var convGo_type2LnsItemKindMap *LnsMap
var convGo_type2FromStemNameMap *LnsMap
var convGo_op2map *LnsMap
// decl alge -- SymbolKind
type convGo_SymbolKind = LnsAny
type convGo_SymbolKind__Arg struct{
}
var convGo_SymbolKind__Arg_Obj = &convGo_SymbolKind__Arg{}
func (self *convGo_SymbolKind__Arg) GetTxt() string {
return "SymbolKind.Arg"
}
type convGo_SymbolKind__Func struct{
Val1 *Ast_TypeInfo
}
func (self *convGo_SymbolKind__Func) GetTxt() string {
return "SymbolKind.Func"
}
type convGo_SymbolKind__Member struct{
Val1 bool
}
func (self *convGo_SymbolKind__Member) GetTxt() string {
return "SymbolKind.Member"
}
type convGo_SymbolKind__Normal struct{
}
var convGo_SymbolKind__Normal_Obj = &convGo_SymbolKind__Normal{}
func (self *convGo_SymbolKind__Normal) GetTxt() string {
return "SymbolKind.Normal"
}
type convGo_SymbolKind__Static struct{
Val1 *Ast_TypeInfo
}
func (self *convGo_SymbolKind__Static) GetTxt() string {
return "SymbolKind.Static"
}
type convGo_SymbolKind__Type struct{
Val1 *Ast_TypeInfo
Val2 bool
}
func (self *convGo_SymbolKind__Type) GetTxt() string {
return "SymbolKind.Type"
}
type convGo_SymbolKind__Var struct{
Val1 *Ast_SymbolInfo
}
func (self *convGo_SymbolKind__Var) GetTxt() string {
return "SymbolKind.Var"
}
// decl alge -- FuncInfo
type convGo_FuncInfo = LnsAny
type convGo_FuncInfo__DeclInfo struct{
Val1 *Nodes_Node
Val2 *Nodes_DeclFuncInfo
}
func (self *convGo_FuncInfo__DeclInfo) GetTxt() string {
return "FuncInfo.DeclInfo"
}
type convGo_FuncInfo__Type struct{
Val1 *Ast_TypeInfo
}
func (self *convGo_FuncInfo__Type) GetTxt() string {
return "FuncInfo.Type"
}
type convGo_FuncInfo__WithClass struct{
Val1 *Ast_TypeInfo
Val2 *Ast_TypeInfo
}
func (self *convGo_FuncInfo__WithClass) GetTxt() string {
return "FuncInfo.WithClass"
}
// decl alge -- CallKind
type convGo_CallKind = LnsAny
type convGo_CallKind__BuiltinCall struct{
Val1 string
Val2 string
}
func (self *convGo_CallKind__BuiltinCall) GetTxt() string {
return "CallKind.BuiltinCall"
}
type convGo_CallKind__FormCall struct{
}
var convGo_CallKind__FormCall_Obj = &convGo_CallKind__FormCall{}
func (self *convGo_CallKind__FormCall) GetTxt() string {
return "CallKind.FormCall"
}
type convGo_CallKind__LuaCall struct{
}
var convGo_CallKind__LuaCall_Obj = &convGo_CallKind__LuaCall{}
func (self *convGo_CallKind__LuaCall) GetTxt() string {
return "CallKind.LuaCall"
}
type convGo_CallKind__Normal struct{
}
var convGo_CallKind__Normal_Obj = &convGo_CallKind__Normal{}
func (self *convGo_CallKind__Normal) GetTxt() string {
return "CallKind.Normal"
}
type convGo_CallKind__RunLoaded struct{
}
var convGo_CallKind__RunLoaded_Obj = &convGo_CallKind__RunLoaded{}
func (self *convGo_CallKind__RunLoaded) GetTxt() string {
return "CallKind.RunLoaded"
}
type convGo_CallKind__RuntimeCall struct{
Val1 *Nodes_Node
}
func (self *convGo_CallKind__RuntimeCall) GetTxt() string {
return "CallKind.RuntimeCall"
}
type convGo_CallKind__SortCall struct{
Val1 *Ast_TypeInfo
}
func (self *convGo_CallKind__SortCall) GetTxt() string {
return "CallKind.SortCall"
}
type convFilter_processRoot__ProcNode_1446_ func (arg1 *Nodes_Node)
// for 146
func convGo_convExp493(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 590
func convGo_convExp2781(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 594
func convGo_convExp2826(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 340
func convGo_convExp1365(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 341
func convGo_convExp1394(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 382
func convGo_convExp1568(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1540
func convGo_convExp7504(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1558
func convGo_convExp7610(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1560
func convGo_convExp7649(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1563
func convGo_convExp7687(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1580
func convGo_convExp7790(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1587
func convGo_convExp7869(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1832
func convGo_convExp10445(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 5160
func convGo_convExp29208(arg1 []LnsAny) bool {
    return Lns_getFromMulti( arg1, 0 ).(bool)
}
// 92: decl @lune.@base.@convGo.isMain
func convGo_isMain_1062_(funcType *Ast_TypeInfo) bool {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( funcType.FP.Get_kind() == Ast_TypeInfoKind__Func) &&
        Lns_GetEnv().SetStackVal( funcType.FP.Get_rawTxt() == "__main") &&
        Lns_GetEnv().SetStackVal( funcType.FP.Get_accessMode() == Ast_AccessMode__Pub) ).(bool)){
        return true
    }
    return false
}

// 211: decl @lune.@base.@convGo.filter
func convGo_filter_1171_(node *Nodes_Node,filter *convGo_convFilter,parent *Nodes_Node) {
    node.FP.ProcessFilter(&filter.Nodes_Filter, ConvGo_Opt2Stem(NewConvGo_Opt(parent)))
}

// 217: decl @lune.@base.@convGo.isAnyType
func convGo_isAnyType_1174_(typeInfo *Ast_TypeInfo) bool {
    var work *Ast_TypeInfo
    work = typeInfo.FP.Get_srcTypeInfo()
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_nilable()) ||
        Lns_GetEnv().SetStackVal( work == Ast_builtinTypeStem) ).(bool){
        return true
    }
    if _switch902 := typeInfo.FP.Get_kind(); _switch902 == Ast_TypeInfoKind__Stem || _switch902 == Ast_TypeInfoKind__Alge {
        return true
    } else if _switch902 == Ast_TypeInfoKind__Alternate {
        if typeInfo.FP.HasBase(){
            return convGo_isAnyType_1174_(typeInfo.FP.Get_baseTypeInfo())
        }
        return true
    } else if _switch902 == Ast_TypeInfoKind__Ext {
        if typeInfo.FP.Get_extedType().FP.Get_kind() == Ast_TypeInfoKind__Stem{
            return true
        }
    }
    return false
}

// 241: decl @lune.@base.@convGo.isClosure
func convGo_isClosure_1177_(typeInfo *Ast_TypeInfo) bool {
    var scope *Ast_Scope
    
    {
        _scope := typeInfo.FP.Get_scope()
        if _scope == nil{
            return false
        } else {
            scope = _scope.(*Ast_Scope)
        }
    }
    return scope.FP.Get_closureSymList().Len() > 0
}

// 276: decl @lune.@base.@convGo.concatGLSym
func convGo_concatGLSym_1189_(name string,external bool) string {
    if external{
        var frontChar LnsInt
        frontChar = LnsInt(name[1-1])
        var front string
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( frontChar >= 97) &&
            Lns_GetEnv().SetStackVal( frontChar <= 122) ).(bool)){
            front = Lns_getVM().String_format("%c", []LnsAny{65 + frontChar - 97})
            
        } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( frontChar >= 65) &&
            Lns_GetEnv().SetStackVal( frontChar <= 90) ).(bool)){
            front = Lns_getVM().String_sub(name,1, 1)
            
        } else { 
            front = Lns_getVM().String_format("G%c", []LnsAny{frontChar})
            
        }
        return front + Lns_getVM().String_sub(name,2, nil)
    }
    return name
}

// 296: decl @lune.@base.@convGo.isInnerDeclType
func convGo_isInnerDeclType_1192_(typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__FormFunc{
        return typeInfo.FP.Get_parentInfo().FP.Get_kind() != Ast_TypeInfoKind__Module
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_parentInfo().FP.Get_kind() != Ast_TypeInfoKind__Module) ||
        Lns_GetEnv().SetStackVal( Ast_TypeInfo2Stem(Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(typeInfo.FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_parent()})&&
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_ownerTypeInfo()}))) == nil) ).(bool){
        return true
    }
    return false
}

// 316: decl @lune.@base.@convGo.getModuleName
func convGo_getModuleName_1198_(typeInfo *Ast_TypeInfo) string {
    if typeInfo == Ast_headTypeInfo{
        return ""
    }
    return Lns_car(Lns_getVM().String_gsub(typeInfo.FP.GetModule().FP.Get_rawTxt(),"@", "")).(string)
}

// 587: decl @lune.@base.@convGo.str2gostr
func convGo_str2gostr_1254_(txt string) string {
    var work string
    work = txt
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(work, "^```", nil, nil))){
        work = convGo_convExp2781(Lns_2DDD(Lns_getVM().String_gsub(Lns_getVM().String_sub(work,4, -4),"^\n", "")))
        
        work = Parser_quoteStr(work)
        
    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(work, "^'", nil, nil))){
        work = convGo_convExp2826(Lns_2DDD(Lns_getVM().String_gsub(Parser_convFromRawToStr(work),"\"", "\\\"")))
        
        work = Lns_getVM().String_format("\"%s\"", []LnsAny{work})
        
    }
    return work
}

// 602: decl @lune.@base.@convGo.getOrgTypeInfo
func convGo_getOrgTypeInfo_1257_(typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    {
        _enumType := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo().FP.Get_nonnilableType().FP.Get_aliasSrc().FP)
        if _enumType != nil {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            return enumType.FP.Get_valTypeInfo()
        }
    }
    return typeInfo.FP.Get_srcTypeInfo().FP.Get_nonnilableType()
}


// 738: decl @lune.@base.@convGo.needConvFormFunc
func convGo_needConvFormFunc_1293_(node *Nodes_ExpCastNode) bool {
    var castType *Ast_TypeInfo
    castType = node.FP.Get_castType().FP.Get_extedType()
    if castType.FP.Get_kind() != Ast_TypeInfoKind__FormFunc{
        return false
    }
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_exp().FP.Get_expType().FP.Get_extedType()
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( castType.FP.Get_argTypeInfoList().Len() != funcType.FP.Get_argTypeInfoList().Len()) ||
        Lns_GetEnv().SetStackVal( castType.FP.Get_retTypeInfoList().Len() != funcType.FP.Get_retTypeInfoList().Len()) ).(bool){
        return true
    }
    for _index, _argType := range( castType.FP.Get_argTypeInfoList().Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if argType != funcType.FP.Get_argTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(){
            return true
        }
    }
    for _index, _retType := range( castType.FP.Get_retTypeInfoList().Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if retType != funcType.FP.Get_retTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(){
            return true
        }
    }
    return false
}

// 769: decl @lune.@base.@convGo.needConvCast
func convGo_needConvCast_1296_(dstType *Ast_TypeInfo,srcType *Ast_TypeInfo) bool {
    if _switch3765 := dstType.FP.Get_kind(); _switch3765 == Ast_TypeInfoKind__Nilable {
        return convGo_needConvCast_1296_(dstType.FP.Get_nonnilableType(), srcType)
    } else if _switch3765 == Ast_TypeInfoKind__Class {
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( dstType == Ast_builtinTypeString) ||
            Lns_GetEnv().SetStackVal( srcType.FP.Get_genSrcTypeInfo().FP.Get_srcTypeInfo().FP.Get_nonnilableType() == dstType.FP.Get_genSrcTypeInfo().FP.Get_srcTypeInfo().FP.Get_nonnilableType()) ).(bool){
            return false
        } else { 
            return true
        }
    } else if _switch3765 == Ast_TypeInfoKind__IF {
        return false
    } else if _switch3765 == Ast_TypeInfoKind__FormFunc {
        return true
    } else if _switch3765 == Ast_TypeInfoKind__Alternate {
        if Lns_op_not(dstType.FP.HasBase()){
            return false
        } else { 
            return convGo_needConvCast_1296_(dstType.FP.Get_baseTypeInfo(), srcType)
        }
    } else if _switch3765 == Ast_TypeInfoKind__Form {
        return true
    } else if _switch3765 == Ast_TypeInfoKind__Prim {
        if Lns_op_not(dstType.FP.Get_nilable()){
            if _switch3729 := dstType; _switch3729 == Ast_builtinTypeInt {
                return true
            } else if _switch3729 == Ast_builtinTypeReal {
                return true
            } else {
                return false
            }
        } else { 
            return false
        }
    } else {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( srcType.FP.Get_kind() == Ast_TypeInfoKind__Class) &&
            Lns_GetEnv().SetStackVal( srcType != Ast_builtinTypeString) ).(bool)){
            return true
        } else { 
            return false
        }
    }
// insert a dummy
    return false
}

// 943: decl @lune.@base.@convGo.getExpListKind
func convGo_getExpListKind_1321_(dstTypeList *LnsList,node *Nodes_ExpListNode) LnsInt {
    if dstTypeList.Len() < node.FP.Get_expList().Len(){
        if dstTypeList.GetAt(dstTypeList.Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() != Ast_TypeInfoKind__DDD{
            return convGo_ExpListKind__Conv
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( dstTypeList.Len() > 1) &&
        Lns_GetEnv().SetStackVal( node.FP.Get_mRetExp()) )){
        for _, _exp := range( node.FP.Get_expList().Items ) {
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            {
                _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
                if _castNode != nil {
                    castNode := _castNode.(*Nodes_ExpCastNode)
                    if convGo_needConvCast_1296_(castNode.FP.Get_castType(), castNode.FP.Get_exp().FP.Get_expType()){
                        return convGo_ExpListKind__Conv
                    }
                }
            }
        }
    }
    var lastExp *Nodes_Node
    lastExp = node.FP.Get_expList().GetAt(node.FP.Get_expList().Len()).(Nodes_NodeDownCast).ToNodes_Node()
    var hasAbbr bool
    if lastExp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr{
        hasAbbr = true
        
        if node.FP.Get_expList().Len() < 2{
            return convGo_ExpListKind__Direct
        }
        lastExp = node.FP.Get_expList().GetAt(node.FP.Get_expList().Len() - 1).(Nodes_NodeDownCast).ToNodes_Node()
        
    } else { 
        hasAbbr = false
        
    }
    if Lns_isCondTrue( Nodes_ExpToDDDNodeDownCastF(lastExp.FP)){
        var mRetExp *Nodes_MRetExp
        
        {
            _mRetExp := node.FP.Get_mRetExp()
            if _mRetExp == nil{
                return convGo_ExpListKind__Slice
            } else {
                mRetExp = _mRetExp.(*Nodes_MRetExp)
            }
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( mRetExp.FP.Get_index() == 1) &&
            Lns_GetEnv().SetStackVal( dstTypeList.GetAt(mRetExp.FP.Get_index()).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() == Ast_TypeInfoKind__DDD) ).(bool)){
            return convGo_ExpListKind__Slice
        }
        return convGo_ExpListKind__Conv
    }
    if lastExp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__DDD{
        var mRetExp *Nodes_MRetExp
        
        {
            _mRetExp := node.FP.Get_mRetExp()
            if _mRetExp == nil{
                return convGo_ExpListKind__Slice
            } else {
                mRetExp = _mRetExp.(*Nodes_MRetExp)
            }
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( mRetExp.FP.Get_index() == 1) &&
            Lns_GetEnv().SetStackVal( dstTypeList.GetAt(mRetExp.FP.Get_index()).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() == Ast_TypeInfoKind__DDD) ).(bool)){
            return convGo_ExpListKind__Direct
        }
    } else { 
        var mRetExp *Nodes_MRetExp
        
        {
            _mRetExp := node.FP.Get_mRetExp()
            if _mRetExp == nil{
                return convGo_ExpListKind__Direct
            } else {
                mRetExp = _mRetExp.(*Nodes_MRetExp)
            }
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(hasAbbr)) &&
            Lns_GetEnv().SetStackVal( mRetExp.FP.Get_index() == 1) ).(bool)){
            return convGo_ExpListKind__Direct
        }
    }
    return convGo_ExpListKind__Conv
}


// 1211: decl @lune.@base.@convGo.isRetGenerics
func convGo_isRetGenerics_1358_(node *Nodes_ExpCallNode) bool {
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_func().FP.Get_expType()
    for _index, _retType := range( funcType.FP.Get_retTypeInfoList().Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( retType.FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(convGo_isAnyType_1174_(node.FP.Get_expTypeList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))) ).(bool)){
            return true
        }
    }
    return false
}

































// 3327: decl @lune.@base.@convGo.getLnsItemKind
func convGo_getLnsItemKind_1815_(typeInfo *Ast_TypeInfo) string {
    {
        _kind := convGo_type2LnsItemKindMap.Items[typeInfo]
        if _kind != nil {
            kind := _kind.(string)
            return kind
        }
    }
    return "LnsItemKindStem"
}





// 4791: decl @lune.@base.@convGo.convFilter.processAndOr.isAndOr
func convFilter_processAndOr__isAndOr_1999_(exp *Nodes_Node) bool {
    {
        _parentNode := Nodes_ExpOp2NodeDownCastF(exp.FP)
        if _parentNode != nil {
            parentNode := _parentNode.(*Nodes_ExpOp2Node)
            if _switch27032 := parentNode.FP.Get_op().Txt; _switch27032 == "and" || _switch27032 == "or" {
                return true
            }
        }
    }
    return false
}

// 5405: decl @lune.@base.@convGo.createFilter
func ConvGo_createFilter(enableTest bool,streamName string,stream Lns_oStream,ast *TransUnit_ASTInfo,option *ConvGo_Option) *Nodes_Filter {
    return &NewconvGo_convFilter(enableTest, streamName, stream, ast, option).Nodes_Filter
}

// 5415: decl @lune.@base.@convGo.Ast2Code
func convGo_Ast2Code_2098_(streamName string,ast *TransUnit_ASTInfo) string {
    var stream *Util_memStream
    stream = NewUtil_memStream()
    var conv *Nodes_Filter
    conv = ConvGo_createFilter(false, streamName, stream.FP, ast, NewConvGo_Option("lns", "lnsgo", ""))
    ast.FP.Get_node().FP.ProcessFilter(conv, ConvGo_Opt2Stem(NewConvGo_Opt(ast.FP.Get_node())))
    return stream.FP.Get_txt()
}

// declaration Class -- Opt
type ConvGo_OptMtd interface {
}
type ConvGo_Opt struct {
    Parent *Nodes_Node
    FP ConvGo_OptMtd
}
func ConvGo_Opt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvGo_Opt).FP
}
type ConvGo_OptDownCast interface {
    ToConvGo_Opt() *ConvGo_Opt
}
func ConvGo_OptDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvGo_OptDownCast)
    if ok { return work.ToConvGo_Opt() }
    return nil
}
func (obj *ConvGo_Opt) ToConvGo_Opt() *ConvGo_Opt {
    return obj
}
func NewConvGo_Opt(arg1 *Nodes_Node) *ConvGo_Opt {
    obj := &ConvGo_Opt{}
    obj.FP = obj
    obj.InitConvGo_Opt(arg1)
    return obj
}
// 40: DeclConstr
func (self *ConvGo_Opt) InitConvGo_Opt(parent *Nodes_Node) {
    self.Parent = parent
    
}


// declaration Class -- Env
type convGo_EnvMtd interface {
    GetEnv(arg1 bool) string
    GetLuavm(arg1 bool) string
    Pop()
    Push(arg1 *Nodes_Node)
}
type convGo_Env struct {
    stack *LnsList
    FP convGo_EnvMtd
}
func convGo_Env2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convGo_Env).FP
}
type convGo_EnvDownCast interface {
    ToconvGo_Env() *convGo_Env
}
func convGo_EnvDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convGo_EnvDownCast)
    if ok { return work.ToconvGo_Env() }
    return nil
}
func (obj *convGo_Env) ToconvGo_Env() *convGo_Env {
    return obj
}
func NewconvGo_Env() *convGo_Env {
    obj := &convGo_Env{}
    obj.FP = obj
    obj.InitconvGo_Env()
    return obj
}
// 59: DeclConstr
func (self *convGo_Env) InitconvGo_Env() {
    self.stack = NewLnsList([]LnsAny{})
    
}

// 62: decl @lune.@base.@convGo.Env.push
func (self *convGo_Env) Push(node *Nodes_Node) {
    self.stack.Insert(Nodes_Node2Stem(node))
}

// 65: decl @lune.@base.@convGo.Env.pop
func (self *convGo_Env) Pop() {
    self.stack.Remove(nil)
}

// 68: decl @lune.@base.@convGo.Env.getLuavm
func (self *convGo_Env) GetLuavm(threading bool) string {
    if self.stack.Len() == 0{
        if threading{
            return "self.LnsEnv.LuaVM"
        } else { 
            return "Lns_getVM()"
        }
    } else { 
        return "Lns_getVMSync()"
    }
// insert a dummy
    return ""
}

// 79: decl @lune.@base.@convGo.Env.getEnv
func (self *convGo_Env) GetEnv(threading bool) string {
    if self.stack.Len() == 0{
        if threading{
            return "self.LnsEnv"
        } else { 
            return "Lns_GetEnv()"
        }
    } else { 
        return "Lns_GetEnvSync()"
    }
// insert a dummy
    return ""
}


// declaration Class -- Option
type ConvGo_OptionMtd interface {
    Get_appName() string
    Get_mainModule() string
    Get_packageName() string
}
type ConvGo_Option struct {
    packageName string
    appName string
    mainModule string
    FP ConvGo_OptionMtd
}
func ConvGo_Option2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvGo_Option).FP
}
type ConvGo_OptionDownCast interface {
    ToConvGo_Option() *ConvGo_Option
}
func ConvGo_OptionDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvGo_OptionDownCast)
    if ok { return work.ToConvGo_Option() }
    return nil
}
func (obj *ConvGo_Option) ToConvGo_Option() *ConvGo_Option {
    return obj
}
func NewConvGo_Option(arg1 string, arg2 string, arg3 string) *ConvGo_Option {
    obj := &ConvGo_Option{}
    obj.FP = obj
    obj.InitConvGo_Option(arg1, arg2, arg3)
    return obj
}
func (self *ConvGo_Option) InitConvGo_Option(arg1 string, arg2 string, arg3 string) {
    self.packageName = arg1
    self.appName = arg2
    self.mainModule = arg3
}
func (self *ConvGo_Option) Get_packageName() string{ return self.packageName }
func (self *ConvGo_Option) Get_appName() string{ return self.appName }
func (self *ConvGo_Option) Get_mainModule() string{ return self.mainModule }

// declaration Class -- convFilter
type convGo_convFilterMtd interface {
    concatSymWithType(arg1 string, arg2 *Ast_TypeInfo) string
    DefaultProcess(arg1 *Nodes_Node, arg2 LnsAny)
    expList2Slice(arg1 *Nodes_ExpListNode, arg2 bool)
    getAccessorSym(arg1 LnsInt, arg2 string) string
    getAlgeSymbol(arg1 *Ast_AlgeValInfo) string
    getCanonicalName(arg1 *Ast_TypeInfo, arg2 bool) string
    getConstrSymbol(arg1 *Ast_TypeInfo) string
    getConv2formName(arg1 *Nodes_Node) string
    getConvExpName(arg1 LnsInt, arg2 *Nodes_ExpListNode) string
    getConvGenericsName(arg1 *Nodes_Node) string
    getEnumGetTxtSym(arg1 *Ast_EnumTypeInfo) string
    getFromStemName(arg1 *Ast_TypeInfo) string
    GetFull(arg1 *Ast_TypeInfo, arg2 bool) string
    getFuncSymbol(arg1 *Ast_TypeInfo) string
    getModuleSym(arg1 *Ast_TypeInfo, arg2 bool) string
    getSymbol(arg1 LnsAny, arg2 string) string
    getSymbolSym(arg1 *Ast_SymbolInfo) string
    getTypeSymbol(arg1 *Ast_TypeInfo) string
    getTypeSymbolWithPrefix(arg1 *Ast_TypeInfo) string
    getVM(arg1 bool, arg2 *Ast_TypeInfo) LnsAny
    Get_moduleInfoManager() Ast_ModuleInfoManager
    Get_optStack() *LnsList
    Get_typeNameCtrl() *Ast_TypeNameCtrl
    IsExtSymbol(arg1 *Ast_SymbolInfo) bool
    IsExtType(arg1 *Ast_TypeInfo) bool
    IsPubSym(arg1 *Ast_SymbolInfo) bool
    IsPubType(arg1 *Ast_TypeInfo) bool
    isSameModDir(arg1 *Ast_TypeInfo) bool
    isSamePackage(arg1 *Ast_TypeInfo) bool
    isSamePackageExtModule(arg1 *Ast_TypeInfo) bool
    needPrefix(arg1 *Ast_TypeInfo) bool
    OutputAccessor(arg1 *Nodes_DeclClassNode)
    outputAdvertise(arg1 *Nodes_DeclClassNode)
    outputAlter2MapFunc(arg1 *LnsMap)
    outputAny2Type(arg1 *Ast_TypeInfo)
    outputAsyncItem(arg1 *Nodes_DeclClassNode)
    OutputCallPrefix(arg1 bool, arg2 LnsInt, arg3 *Nodes_Node, arg4 LnsAny, arg5 *Ast_SymbolInfo)(bool, LnsAny)
    outputCastReceiver(arg1 *Nodes_DeclClassNode)
    outputClassType(arg1 *Nodes_DeclClassNode)
    outputConstructor(arg1 *Nodes_DeclClassNode)
    outputConvExt(arg1 *Nodes_Node)
    outputConvItemType(arg1 *Ast_TypeInfo, arg2 LnsAny)
    outputConvItemTypeList(arg1 *LnsList, arg2 LnsAny)
    outputConvToForm(arg1 *Nodes_ExpCastNode)
    OutputDeclFunc(arg1 LnsAny) *convGo_FuncConv
    outputDeclFuncArg(arg1 *Ast_TypeInfo)
    OutputDeclFuncInfo(arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo)
    outputDownCast(arg1 *Nodes_DeclClassNode)
    outputDummyAbstractMethod(arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo)
    outputDummyAbstractMethodOfClass(arg1 *Ast_TypeInfo)
    outputDummyReturn(arg1 *LnsList)
    outputForeachLua(arg1 *Nodes_Node, arg2 bool, arg3 *Nodes_Node, arg4 *Ast_SymbolInfo, arg5 LnsAny, arg6 *Nodes_BlockNode)
    outputIFMethods(arg1 *Nodes_DeclClassNode)
    outputImplicitCast(arg1 *Ast_TypeInfo, arg2 *Nodes_Node, arg3 *Nodes_ExpCastNode)
    outputImport(arg1 *Nodes_ImportNode)
    outputInterfaceType(arg1 *Nodes_DeclClassNode)
    outputLetVar(arg1 *Nodes_DeclVarNode)
    outputMapping(arg1 *Nodes_DeclClassNode)
    outputMethodIF(arg1 *Nodes_DeclClassNode)
    outputModule(arg1 *Ast_TypeInfo, arg2 bool)
    outputModuleImport(arg1 *Nodes_DeclClassNode)
    outputNewSetup(arg1 string, arg2 *Ast_TypeInfo)
    outputNilAccCall(arg1 *Nodes_ExpCallNode)
    outputRetType(arg1 *LnsList)
    outputStaticMember(arg1 *Nodes_DeclClassNode)
    outputStem2Type(arg1 *Ast_TypeInfo)
    outputSymbol(arg1 LnsAny, arg2 string)
    outputToStem(arg1 *Nodes_DeclClassNode)
    outputTopScopeVar(arg1 *Nodes_DeclVarNode)
    PopIndent()
    popProcessMode()
    ProcessAbbr(arg1 *Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(arg1 *Nodes_AliasNode, arg2 LnsAny)
    processAndOr(arg1 *Nodes_ExpOp2Node, arg2 string, arg3 *Nodes_Node)
    ProcessApply(arg1 *Nodes_ApplyNode, arg2 LnsAny)
    ProcessBlankLine(arg1 *Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(arg1 *Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(arg1 *Nodes_BreakNode, arg2 LnsAny)
    processCond(arg1 *Nodes_Node, arg2 *Nodes_Node)
    processConvExp(arg1 LnsInt, arg2 *LnsList, arg3 LnsAny)
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
    ProcessEnv(arg1 *Nodes_EnvNode, arg2 LnsAny)
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
    processGenericsCall(arg1 *Nodes_ExpCallNode)
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
    processSetFromExpList(arg1 string, arg2 *LnsList, arg3 *Nodes_ExpListNode)
    ProcessShebang(arg1 *Nodes_ShebangNode, arg2 LnsAny)
    ProcessStmtExp(arg1 *Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(arg1 *Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(arg1 *Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(arg1 *Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(arg1 *Nodes_TestCaseNode, arg2 LnsAny)
    ProcessUnboxing(arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(arg1 *Nodes_WhileNode, arg2 LnsAny)
    PushIndent(arg1 LnsAny)
    pushProcessMode(arg1 LnsInt)
    ReturnToSource()
    SwitchToHeader()
    type2gotype(arg1 *Ast_TypeInfo) string
    type2gotypeOrg(arg1 *Ast_TypeInfo, arg2 bool) string
    Write(arg1 string)
    Writeln(arg1 string)
}
type convGo_convFilter struct {
    Nodes_Filter
    stream *Util_SimpleSourceOStream
    processMode LnsInt
    processModeStack *LnsList
    moduleTypeInfo *Ast_TypeInfo
    moduleScope *Ast_Scope
    builtin2runtime *LnsMap
    type2gotypeMap *LnsMap
    nodeManager *Nodes_NodeManager
    enableTest bool
    processInfo *Ast_ProcessInfo
    noneNode *Nodes_NoneNode
    option *ConvGo_Option
    modDir string
    module2PackSym *LnsMap
    env *convGo_Env
    FP convGo_convFilterMtd
}
func convGo_convFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convGo_convFilter).FP
}
type convGo_convFilterDownCast interface {
    ToconvGo_convFilter() *convGo_convFilter
}
func convGo_convFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convGo_convFilterDownCast)
    if ok { return work.ToconvGo_convFilter() }
    return nil
}
func (obj *convGo_convFilter) ToconvGo_convFilter() *convGo_convFilter {
    return obj
}
func NewconvGo_convFilter(arg1 bool, arg2 string, arg3 Lns_oStream, arg4 *TransUnit_ASTInfo, arg5 *ConvGo_Option) *convGo_convFilter {
    obj := &convGo_convFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvGo_convFilter(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *convGo_convFilter) PopIndent() {
self.stream. FP.PopIndent( )
}
func (self *convGo_convFilter) PushIndent(arg1 LnsAny) {
self.stream. FP.PushIndent( arg1)
}
func (self *convGo_convFilter) ReturnToSource() {
self.stream. FP.ReturnToSource( )
}
func (self *convGo_convFilter) SwitchToHeader() {
self.stream. FP.SwitchToHeader( )
}
func (self *convGo_convFilter) Write(arg1 string) {
self.stream. FP.Write( arg1)
}
func (self *convGo_convFilter) Writeln(arg1 string) {
self.stream. FP.Writeln( arg1)
}
// 125: DeclConstr
func (self *convGo_convFilter) InitconvGo_convFilter(enableTest bool,streamName string,stream Lns_oStream,ast *TransUnit_ASTInfo,option *ConvGo_Option) {
    self.InitNodes_Filter(true, ast.FP.Get_moduleTypeInfo(), ast.FP.Get_moduleTypeInfo().FP.Get_scope())
    self.env = NewconvGo_Env()
    
    self.option = option
    
    self.processInfo = ast.FP.Get_processInfo()
    
    self.stream = NewUtil_SimpleSourceOStream(stream, nil, 4)
    
    self.processMode = convGo_ProcessMode__Main
    
    self.processModeStack = NewLnsList([]LnsAny{})
    
    self.moduleTypeInfo = ast.FP.Get_moduleTypeInfo()
    
    self.moduleScope = Lns_unwrap( ast.FP.Get_moduleTypeInfo().FP.Get_scope()).(*Ast_Scope)
    
    self.builtin2runtime = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.type2gotypeMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.nodeManager = NewNodes_NodeManager()
    
    self.enableTest = enableTest
    
    self.module2PackSym = NewLnsMap( map[LnsAny]LnsAny{})
    
    var modDir string
    modDir = self.moduleTypeInfo.FP.GetParentFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), false)
    self.modDir = convGo_convExp493(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(modDir,"@", "")).(string),"%.$", "")))
    
    self.noneNode = Nodes_NoneNode_create(self.nodeManager, Parser_noneToken.Pos, false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}))
    
}

// 153: decl @lune.@base.@convGo.convFilter.getVM
func (self *convGo_convFilter) getVM(threading bool,typeInfo *Ast_TypeInfo) LnsAny {
    var txt string
    
    {
        _txt := self.builtin2runtime.Items[typeInfo]
        if _txt == nil{
            return nil
        } else {
            txt = _txt.(string)
        }
    }
    return Lns_car(Lns_getVM().String_gsub(txt,"GETVM", self.env.FP.GetLuavm(threading))).(string)
}

// 160: decl @lune.@base.@convGo.convFilter.pushProcessMode
func (self *convGo_convFilter) pushProcessMode(mode LnsInt) {
    self.processModeStack.Insert(self.processMode)
    self.processMode = mode
    
}

// 164: decl @lune.@base.@convGo.convFilter.popProcessMode
func (self *convGo_convFilter) popProcessMode() {
    self.processMode = self.processModeStack.GetAt(self.processModeStack.Len()).(LnsInt)
    
    self.processModeStack.Remove(nil)
}

// 169: decl @lune.@base.@convGo.convFilter.isPubType
func (self *convGo_convFilter) IsPubType(typeInfo *Ast_TypeInfo) bool {
    if Ast_isPubToExternal(typeInfo.FP.Get_accessMode()){
        if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Func{
            if _switch664 := typeInfo.FP.Get_parentInfo().FP.Get_kind(); _switch664 == Ast_TypeInfoKind__Class || _switch664 == Ast_TypeInfoKind__Enum {
                return self.FP.IsPubType(typeInfo.FP.Get_parentInfo())
            }
        }
        return true
    }
    return typeInfo.FP.GetModule() != self.moduleTypeInfo
}

// 182: decl @lune.@base.@convGo.convFilter.isPubSym
func (self *convGo_convFilter) IsPubSym(symbol *Ast_SymbolInfo) bool {
    if Ast_isPubToExternal(symbol.FP.Get_accessMode()){
        return true
    }
    return symbol.FP.GetModule() != self.moduleTypeInfo
}

// 189: decl @lune.@base.@convGo.convFilter.isExtType
func (self *convGo_convFilter) IsExtType(typeInfo *Ast_TypeInfo) bool {
    return typeInfo.FP.GetModule() != self.moduleTypeInfo
}

// 192: decl @lune.@base.@convGo.convFilter.isExtSymbol
func (self *convGo_convFilter) IsExtSymbol(symbol *Ast_SymbolInfo) bool {
    return symbol.FP.GetModule() != self.moduleTypeInfo
}

// 309: decl @lune.@base.@convGo.convFilter.getCanonicalName
func (self *convGo_convFilter) getCanonicalName(typeInfo *Ast_TypeInfo,localFlag bool) string {
    var enumName string
    enumName = typeInfo.FP.GetFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), localFlag)
    return Lns_getVM().String_format("%s", []LnsAny{Lns_car(Lns_getVM().String_gsub(enumName,"&", "")).(string)})
}

// 323: decl @lune.@base.@convGo.convFilter.concatSymWithType
func (self *convGo_convFilter) concatSymWithType(name string,typeInfo *Ast_TypeInfo) string {
    var modName string
    modName = convGo_getModuleName_1198_(typeInfo.FP.GetModule())
    var typeName string
    if modName == ""{
        typeName = name
        
    } else { 
        typeName = Lns_getVM().String_format("%s_%s", []LnsAny{modName, name})
        
    }
    return convGo_concatGLSym_1189_(typeName, self.FP.IsPubType(typeInfo))
}

// 334: decl @lune.@base.@convGo.convFilter.isSamePackageExtModule
func (self *convGo_convFilter) isSamePackageExtModule(typeInfo *Ast_TypeInfo) bool {
    var extModuleType *Ast_NormalTypeInfo
    
    {
        _extModuleType := Ast_NormalTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo().FP.Get_nonnilableType().FP.Get_aliasSrc().FP)
        if _extModuleType == nil{
            Util_err(Lns_getVM().String_format("illegal type -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}))
        } else {
            extModuleType = _extModuleType.(*Ast_NormalTypeInfo)
        }
    }
    var requireParent string
    requireParent = convGo_convExp1365(Lns_2DDD(Lns_getVM().String_gsub(extModuleType.FP.Get_requirePath(),"[^%.]+$", "")))
    var moduleParent string
    moduleParent = convGo_convExp1394(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(self.FP.GetFull(self.moduleTypeInfo, false),"[^@%.]+$", "")).(string),"@", "")))
    return requireParent == moduleParent
}

// 346: decl @lune.@base.@convGo.convFilter.isSameModDir
func (self *convGo_convFilter) isSameModDir(moduleTypeInfo *Ast_TypeInfo) bool {
    if moduleTypeInfo.FP.Get_parentInfo() == self.moduleTypeInfo.FP.Get_parentInfo(){
        return true
    }
    return false
}

// 353: decl @lune.@base.@convGo.convFilter.isSamePackage
func (self *convGo_convFilter) isSamePackage(typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__ExtModule{
        return self.FP.isSamePackageExtModule(typeInfo)
    }
    return self.FP.isSameModDir(typeInfo)
}

// 360: decl @lune.@base.@convGo.convFilter.needPrefix
func (self *convGo_convFilter) needPrefix(typeInfo *Ast_TypeInfo) bool {
    if Ast_isBuiltin(typeInfo.FP.Get_typeId()){
        return false
    }
    return Lns_op_not(self.FP.isSamePackage(typeInfo))
}

// 368: decl @lune.@base.@convGo.convFilter.getSymbol
func (self *convGo_convFilter) getSymbol(kind LnsAny,name string) string {
    __func__ := "@lune.@base.@convGo.convFilter.getSymbol"
    var symbolName string
    if convGo_golanKeywordSet.Has(name){
        symbolName = Lns_getVM().String_format("_%s", []LnsAny{name})
        
    } else { 
        symbolName = name
        
    }
    switch _exp2235 := kind.(type) {
    case *convGo_SymbolKind__Arg:
        return symbolName
    case *convGo_SymbolKind__Var:
    symbolInfo := _exp2235.Val1
        var modName string
        modName = convGo_convExp1568(Lns_2DDD(Lns_getVM().String_gsub(self.moduleTypeInfo.FP.Get_rawTxt(),"@", "")))
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symbolInfo.FP.GetModule() != self.moduleTypeInfo) &&
            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_staticFlag()) ).(bool)){
            symbolName = Lns_getVM().String_format("%s_%s", []LnsAny{convGo_getModuleName_1198_(symbolInfo.FP.GetModule()), symbolInfo.FP.Get_name()})
            
        } else if name == "__mod__"{
            symbolName = Lns_getVM().String_format("%s__mod__", []LnsAny{modName})
            
        } else if symbolInfo.FP.Get_scope() == self.moduleScope{
            symbolName = convGo_concatGLSym_1189_(Lns_getVM().String_format("%s_", []LnsAny{modName}) + symbolName, Ast_isPubToExternal(symbolInfo.FP.Get_accessMode()))
            
        } else if Lns_op_not(symbolInfo.FP.Get_posForModToRef()){
            if symbolName != "__func__"{
                symbolName = "_"
                
            }
        }
        if self.FP.needPrefix(symbolInfo.FP.GetModule()){
            symbolName = Lns_getVM().String_format("%s.%s", []LnsAny{convGo_getModuleName_1198_(symbolInfo.FP.GetModule()), symbolName})
            
        }
    case *convGo_SymbolKind__Member:
    external := _exp2235.Val1
        symbolName = convGo_concatGLSym_1189_(symbolName, external)
        
    case *convGo_SymbolKind__Func:
    typeInfo := _exp2235.Val1
        if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Method{
            if _switch1763 := symbolName; _switch1763 == "_toMap" {
                return "ToMap"
            } else {
                symbolName = convGo_concatGLSym_1189_(symbolName, Ast_isPubToExternal(typeInfo.FP.Get_accessMode()))
                
            }
        } else { 
            var prefix LnsAny
            prefix = nil
            if _switch2034 := typeInfo.FP.Get_parentInfo().FP.Get_kind(); _switch2034 == Ast_TypeInfoKind__Module || _switch2034 == Ast_TypeInfoKind__Func || _switch2034 == Ast_TypeInfoKind__Method {
                if convGo_isInnerDeclType_1192_(typeInfo){
                    if Lns_op_not(convGo_isClosure_1177_(typeInfo)){
                        var parentName string
                        parentName = typeInfo.FP.GetParentFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), true)
                        symbolName = Lns_getVM().String_format("%s_%s_%d_", []LnsAny{Lns_car(Lns_getVM().String_gsub(parentName,"[%.@]", "_")).(string), symbolName, typeInfo.FP.Get_typeId()})
                        
                    }
                } else { 
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isBuiltin(typeInfo.FP.Get_typeId()))) &&
                        Lns_GetEnv().SetStackVal( Lns_op_not(self.FP.IsPubType(typeInfo))) ).(bool)){
                        symbolName = Lns_getVM().String_format("%s_%d_", []LnsAny{symbolName, typeInfo.FP.Get_typeId()})
                        
                    }
                    symbolName = self.FP.concatSymWithType(symbolName, typeInfo)
                    
                }
            } else if _switch2034 == Ast_TypeInfoKind__Enum || _switch2034 == Ast_TypeInfoKind__Class {
                var parentInfo *Ast_TypeInfo
                parentInfo = typeInfo.FP.Get_parentInfo()
                symbolName = Lns_getVM().String_format("%s_%s", []LnsAny{self.FP.getSymbol(&convGo_SymbolKind__Type{parentInfo, false}, parentInfo.FP.Get_rawTxt()), symbolName})
                
                if Lns_op_not(self.FP.IsPubType(typeInfo)){
                    symbolName = Lns_getVM().String_format("%s_%d_", []LnsAny{symbolName, typeInfo.FP.Get_typeId()})
                    
                }
            } else if _switch2034 == Ast_TypeInfoKind__ExtModule {
                symbolName = convGo_concatGLSym_1189_(symbolName, true)
                
                if Lns_op_not(self.FP.isSamePackageExtModule(typeInfo.FP.Get_parentInfo())){
                    prefix = typeInfo.FP.Get_parentInfo().FP.Get_rawTxt()
                    
                }
            } else {
                Util_err(Lns_getVM().String_format("%s: not support -- %s:%s", []LnsAny{__func__, typeInfo.FP.GetFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), true), Ast_TypeInfoKind_getTxt( typeInfo.FP.Get_parentInfo().FP.Get_kind())}))
            }
            if Lns_op_not(prefix){
                if self.FP.needPrefix(typeInfo.FP.GetModule()){
                    prefix = convGo_getModuleName_1198_(typeInfo)
                    
                }
            }
            if prefix != nil{
                prefix_5804 := prefix.(string)
                symbolName = Lns_getVM().String_format("%s.%s", []LnsAny{prefix_5804, symbolName})
                
            }
        }
    case *convGo_SymbolKind__Type:
    typeInfo := _exp2235.Val1
    needPrefix := _exp2235.Val2
        if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__FormFunc{
            return self.FP.getSymbol(&convGo_SymbolKind__Func{typeInfo}, symbolName)
        }
        var workName string
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( convGo_isInnerDeclType_1192_(typeInfo)) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isBuiltin(typeInfo.FP.Get_typeId()))) ).(bool)){
            workName = Lns_getVM().String_format("%s%d", []LnsAny{name, typeInfo.FP.Get_typeId()})
            
        } else { 
            workName = symbolName
            
        }
        symbolName = self.FP.concatSymWithType(workName, typeInfo)
        
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( needPrefix) &&
            Lns_GetEnv().SetStackVal( self.FP.needPrefix(typeInfo.FP.GetModule())) ).(bool)){
            symbolName = Lns_getVM().String_format("%s.%s", []LnsAny{convGo_getModuleName_1198_(typeInfo), symbolName})
            
        }
    case *convGo_SymbolKind__Static:
    typeInfo := _exp2235.Val1
        var workName string
        workName = self.FP.getSymbol(&convGo_SymbolKind__Type{typeInfo, true}, typeInfo.FP.Get_rawTxt())
        symbolName = Lns_getVM().String_format("%s__%s", []LnsAny{workName, name})
        
    case *convGo_SymbolKind__Normal:
    }
    return symbolName
}

// 498: decl @lune.@base.@convGo.convFilter.getTypeSymbol
func (self *convGo_convFilter) getTypeSymbol(typeInfo *Ast_TypeInfo) string {
    var orgType *Ast_TypeInfo
    orgType = typeInfo.FP.Get_srcTypeInfo().FP.Get_nonnilableType().FP.Get_aliasSrc()
    return self.FP.getSymbol(&convGo_SymbolKind__Type{orgType, false}, orgType.FP.Get_rawTxt())
}

// 503: decl @lune.@base.@convGo.convFilter.getTypeSymbolWithPrefix
func (self *convGo_convFilter) getTypeSymbolWithPrefix(typeInfo *Ast_TypeInfo) string {
    var orgType *Ast_TypeInfo
    orgType = typeInfo.FP.Get_srcTypeInfo().FP.Get_nonnilableType().FP.Get_aliasSrc()
    return self.FP.getSymbol(&convGo_SymbolKind__Type{orgType, true}, orgType.FP.Get_rawTxt())
}

// 509: decl @lune.@base.@convGo.convFilter.getConstrSymbol
func (self *convGo_convFilter) getConstrSymbol(typeInfo *Ast_TypeInfo) string {
    return Lns_getVM().String_format("Init%s", []LnsAny{self.FP.getTypeSymbol(typeInfo)})
}

// 513: decl @lune.@base.@convGo.convFilter.getFuncSymbol
func (self *convGo_convFilter) getFuncSymbol(typeInfo *Ast_TypeInfo) string {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Method) &&
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_staticFlag()) ).(bool)){
        return self.FP.getSymbol(&convGo_SymbolKind__Static{typeInfo.FP.Get_parentInfo()}, typeInfo.FP.Get_rawTxt())
    }
    return self.FP.getSymbol(&convGo_SymbolKind__Func{typeInfo}, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_rawTxt() == "") &&
        Lns_GetEnv().SetStackVal( "_anonymous") ||
        Lns_GetEnv().SetStackVal( typeInfo.FP.Get_rawTxt()) ).(string))
}

// 522: decl @lune.@base.@convGo.convFilter.getAlgeSymbol
func (self *convGo_convFilter) getAlgeSymbol(valInfo *Ast_AlgeValInfo) string {
    return self.FP.getSymbol(&convGo_SymbolKind__Static{&valInfo.FP.Get_algeTpye().Ast_TypeInfo}, valInfo.FP.Get_name())
}

// 526: decl @lune.@base.@convGo.convFilter.getSymbolSym
func (self *convGo_convFilter) getSymbolSym(symbolInfo *Ast_SymbolInfo) string {
    if _switch2591 := symbolInfo.FP.Get_kind(); _switch2591 == Ast_SymbolKind__Fun || _switch2591 == Ast_SymbolKind__Mtd {
        return self.FP.getFuncSymbol(symbolInfo.FP.Get_typeInfo())
    } else if _switch2591 == Ast_SymbolKind__Arg {
        return self.FP.getSymbol(convGo_SymbolKind__Arg_Obj, symbolInfo.FP.Get_name())
    } else if _switch2591 == Ast_SymbolKind__Var {
        return self.FP.getSymbol(&convGo_SymbolKind__Var{symbolInfo}, symbolInfo.FP.Get_name())
    } else if _switch2591 == Ast_SymbolKind__Mbr {
        if symbolInfo.FP.Get_staticFlag(){
            return self.FP.getSymbol(&convGo_SymbolKind__Static{symbolInfo.FP.Get_namespaceTypeInfo()}, symbolInfo.FP.Get_name())
        }
        return self.FP.getSymbol(&convGo_SymbolKind__Member{Ast_isPubToExternal(symbolInfo.FP.Get_accessMode())}, symbolInfo.FP.Get_name())
    } else if _switch2591 == Ast_SymbolKind__Typ {
        if Lns_isCondTrue( Ast_AliasTypeInfoDownCastF(symbolInfo.FP.Get_typeInfo().FP)){
            return self.FP.getSymbol(&convGo_SymbolKind__Var{symbolInfo}, symbolInfo.FP.Get_name())
        }
        return self.FP.getTypeSymbol(symbolInfo.FP.Get_typeInfo())
    } else {
        Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind())}))
    }
// insert a dummy
    return ""
}

// 557: decl @lune.@base.@convGo.convFilter.getAccessorSym
func (self *convGo_convFilter) getAccessorSym(accessMode LnsInt,name string) string {
    return self.FP.getSymbol(&convGo_SymbolKind__Member{Ast_isPubToExternal(accessMode)}, name)
}

// 561: decl @lune.@base.@convGo.convFilter.outputSymbol
func (self *convGo_convFilter) outputSymbol(kind LnsAny,name string) {
    self.FP.Write(self.FP.getSymbol(kind, name))
}

// 565: decl @lune.@base.@convGo.convFilter.getConv2formName
func (self *convGo_convFilter) getConv2formName(node *Nodes_Node) string {
    return Lns_getVM().String_format("conv2Form%d", []LnsAny{node.FP.Get_id()})
}

// 569: decl @lune.@base.@convGo.convFilter.getConvGenericsName
func (self *convGo_convFilter) getConvGenericsName(node *Nodes_Node) string {
    return Lns_getVM().String_format("lns_convGenerics%d", []LnsAny{node.FP.Get_id()})
}

// 574: decl @lune.@base.@convGo.convFilter.getModuleSym
func (self *convGo_convFilter) getModuleSym(moduleTypeInfo *Ast_TypeInfo,addDot bool) string {
    {
        _packSym := self.module2PackSym.Items[moduleTypeInfo]
        if _packSym != nil {
            packSym := _packSym.(string)
            if addDot{
                return Lns_getVM().String_format("%s.", []LnsAny{packSym})
            }
            return packSym
        }
    }
    return ""
}

// 609: decl @lune.@base.@convGo.convFilter.type2gotypeOrg
func (self *convGo_convFilter) type2gotypeOrg(typeInfo *Ast_TypeInfo,addClassAster bool) string {
    if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__DDD{
        return "[]LnsAny"
    }
    if convGo_isAnyType_1174_(typeInfo){
        return "LnsAny"
    }
    var orgType *Ast_TypeInfo
    orgType = convGo_getOrgTypeInfo_1257_(typeInfo)
    {
        _goType := self.type2gotypeMap.Items[orgType]
        if _goType != nil {
            goType := _goType.(string)
            return goType
        }
    }
    if _switch3083 := orgType.FP.Get_kind(); _switch3083 == Ast_TypeInfoKind__Ext {
        if orgType.FP.Get_extedType().FP.Get_kind() == Ast_TypeInfoKind__Stem{
            return "LnsAny"
        }
        return "*Lns_luaValue"
    } else if _switch3083 == Ast_TypeInfoKind__List || _switch3083 == Ast_TypeInfoKind__Array {
        return "*LnsList"
    } else if _switch3083 == Ast_TypeInfoKind__Set {
        return "*LnsSet"
    } else if _switch3083 == Ast_TypeInfoKind__Map {
        return "*LnsMap"
    } else if _switch3083 == Ast_TypeInfoKind__Form {
        return "LnsForm"
    } else if _switch3083 == Ast_TypeInfoKind__FormFunc {
        return self.FP.getFuncSymbol(typeInfo)
    } else if _switch3083 == Ast_TypeInfoKind__Class {
        if typeInfo.FP.Get_genSrcTypeInfo() == TransUnit_getBuiltinFunc().G__pipe_{
            return "*Lns__pipe"
        }
        var symbol string
        symbol = self.FP.getTypeSymbolWithPrefix(typeInfo)
        if addClassAster{
            return "*" + symbol
        }
        return symbol
    } else if _switch3083 == Ast_TypeInfoKind__IF {
        return self.FP.getTypeSymbolWithPrefix(typeInfo)
    } else if _switch3083 == Ast_TypeInfoKind__Alternate {
        return self.FP.type2gotypeOrg(typeInfo.FP.Get_baseTypeInfo(), addClassAster)
    }
    Util_err(Lns_getVM().String_format("not support yet -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}))
// insert a dummy
    return ""
}

// 667: decl @lune.@base.@convGo.convFilter.type2gotype
func (self *convGo_convFilter) type2gotype(typeInfo *Ast_TypeInfo) string {
    return self.FP.type2gotypeOrg(typeInfo, true)
}

// 681: decl @lune.@base.@convGo.convFilter.outputAny2Type
func (self *convGo_convFilter) outputAny2Type(dstType *Ast_TypeInfo) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(convGo_isAnyType_1174_(dstType))) &&
        Lns_GetEnv().SetStackVal( dstType.FP.Get_kind() != Ast_TypeInfoKind__Alternate) ).(bool)){
        self.FP.Write(Lns_getVM().String_format(".(%s)", []LnsAny{self.FP.type2gotype(dstType)}))
    }
}

// 688: decl @lune.@base.@convGo.convFilter.outputStem2Type
func (self *convGo_convFilter) outputStem2Type(dstType *Ast_TypeInfo) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( dstType.FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
        Lns_GetEnv().SetStackVal( dstType.FP.HasBase()) ).(bool)){
        self.FP.Write(Lns_getVM().String_format(".(%s)", []LnsAny{self.FP.type2gotype(dstType)}))
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( dstType.FP.Get_kind() == Ast_TypeInfoKind__Class) &&
        Lns_GetEnv().SetStackVal( dstType != Ast_builtinTypeString) ).(bool)){
        self.FP.Write(Lns_getVM().String_format(".(%sDownCast).To%s()", []LnsAny{self.FP.getTypeSymbol(dstType), self.FP.getTypeSymbol(dstType)}))
    } else { 
        self.FP.outputAny2Type(dstType)
    }
}

// 713: decl @lune.@base.@convGo.convFilter.processBlankLine
func (self *convGo_convFilter) ProcessBlankLine(node *Nodes_BlankLineNode,_opt LnsAny) {
}

// 713: decl @lune.@base.@convGo.convFilter.processNone
func (self *convGo_convFilter) ProcessNone(node *Nodes_NoneNode,_opt LnsAny) {
}

// 722: decl @lune.@base.@convGo.convFilter.processImport
func (self *convGo_convFilter) ProcessImport(node *Nodes_ImportNode,_opt LnsAny) {
    if node.FP.Get_modulePath() == "lune.base.Depend"{
        self.FP.Writeln("Lns_LuaVer_init()")
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(self.FP.isSameModDir(node.FP.Get_moduleTypeInfo()))) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isBuiltin(node.FP.Get_moduleTypeInfo().FP.Get_typeId()))) ).(bool)){
        self.FP.Write(Lns_getVM().String_format("%s.", []LnsAny{convGo_getModuleName_1198_(node.FP.Get_moduleTypeInfo())}))
    }
    self.FP.Writeln(Lns_getVM().String_format("Lns_%s_init()", []LnsAny{convGo_getModuleName_1198_(node.FP.Get_moduleTypeInfo())}))
}

// 830: decl @lune.@base.@convGo.convFilter.outputImplicitCast
func (self *convGo_convFilter) outputImplicitCast(castType *Ast_TypeInfo,node *Nodes_Node,parent *Nodes_ExpCastNode) {
    if _switch4429 := castType.FP.Get_kind(); _switch4429 == Ast_TypeInfoKind__Nilable {
        self.FP.outputImplicitCast(castType.FP.Get_nonnilableType(), node, parent)
    } else if _switch4429 == Ast_TypeInfoKind__Class {
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( castType == Ast_builtinTypeString) ||
            Lns_GetEnv().SetStackVal( node.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Alternate) ||
            Lns_GetEnv().SetStackVal( node.FP.Get_expType().FP.Get_genSrcTypeInfo().FP.Get_srcTypeInfo().FP.Get_nonnilableType() == castType.FP.Get_genSrcTypeInfo().FP.Get_srcTypeInfo().FP.Get_nonnilableType()) ).(bool){
            convGo_filter_1171_(node, self, &parent.Nodes_Node)
        } else { 
            if convGo_isAnyType_1174_(node.FP.Get_expType()){
                self.FP.Write(Lns_getVM().String_format("%sDownCastF(", []LnsAny{self.FP.getTypeSymbol(castType)}))
                convGo_filter_1171_(node, self, &parent.Nodes_Node)
                self.FP.Write(")")
            } else { 
                self.FP.Write("&")
                convGo_filter_1171_(node, self, &parent.Nodes_Node)
                self.FP.Write(Lns_getVM().String_format(".%s", []LnsAny{self.FP.getTypeSymbol(castType)}))
            }
        }
    } else if _switch4429 == Ast_TypeInfoKind__IF {
        convGo_filter_1171_(node, self, &parent.Nodes_Node)
        if Ast_isClass(node.FP.Get_expType()){
            self.FP.Write(".FP")
        }
    } else if _switch4429 == Ast_TypeInfoKind__FormFunc {
        _ = node.FP.Get_expType()
        if convGo_needConvFormFunc_1293_(parent){
            self.FP.Write(Lns_getVM().String_format("%s(", []LnsAny{self.FP.getConv2formName(&parent.Nodes_Node)}))
            convGo_filter_1171_(node, self, &parent.Nodes_Node)
            self.FP.Write(")")
        } else { 
            self.FP.Write(Lns_getVM().String_format("%s(", []LnsAny{self.FP.getTypeSymbol(castType)}))
            convGo_filter_1171_(node, self, &parent.Nodes_Node)
            self.FP.Write(")")
        }
    } else if _switch4429 == Ast_TypeInfoKind__Alternate {
        if Lns_op_not(castType.FP.HasBase()){
            if Ast_isClass(node.FP.Get_expType().FP.Get_nonnilableType()){
                self.FP.Write(Lns_getVM().String_format("%s2Stem(", []LnsAny{self.FP.getTypeSymbolWithPrefix(node.FP.Get_expType().FP.Get_nonnilableType())}))
                convGo_filter_1171_(node, self, &parent.Nodes_Node)
                self.FP.Write(")")
            } else { 
                convGo_filter_1171_(node, self, &parent.Nodes_Node)
            }
        } else { 
            self.FP.outputImplicitCast(castType.FP.Get_baseTypeInfo(), node, parent)
        }
    } else if _switch4429 == Ast_TypeInfoKind__Form {
        self.FP.Write(Lns_getVM().String_format("%s(", []LnsAny{self.FP.getConv2formName(&parent.Nodes_Node)}))
        convGo_filter_1171_(node, self, &parent.Nodes_Node)
        self.FP.Write(")")
    } else if _switch4429 == Ast_TypeInfoKind__Prim {
        if Lns_op_not(node.FP.Get_expType().FP.Get_nilable()){
            if _switch4311 := castType; _switch4311 == Ast_builtinTypeInt {
                self.FP.Write("LnsInt(")
                convGo_filter_1171_(node, self, &parent.Nodes_Node)
                self.FP.Write(")")
            } else if _switch4311 == Ast_builtinTypeReal {
                self.FP.Write("LnsReal(")
                convGo_filter_1171_(node, self, &parent.Nodes_Node)
                self.FP.Write(")")
            } else {
                convGo_filter_1171_(node, self, &parent.Nodes_Node)
            }
        } else { 
            convGo_filter_1171_(node, self, &parent.Nodes_Node)
        }
    } else {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( node.FP.Get_expType().FP.Get_nilable()) &&
            Lns_GetEnv().SetStackVal( Ast_isClass(node.FP.Get_expType().FP.Get_nonnilableType())) ).(bool)){
            self.FP.Write(Lns_getVM().String_format("%s2Stem(", []LnsAny{self.FP.getTypeSymbolWithPrefix(node.FP.Get_expType().FP.Get_nonnilableType())}))
            convGo_filter_1171_(node, self, &parent.Nodes_Node)
            self.FP.Write(")")
        } else { 
            convGo_filter_1171_(node, self, &parent.Nodes_Node)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( node.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Class) &&
                Lns_GetEnv().SetStackVal( node.FP.Get_expType() != Ast_builtinTypeString) ).(bool)){
                self.FP.Write(".FP")
            }
        }
    }
}

// 1011: decl @lune.@base.@convGo.convFilter.getConvExpName
func (self *convGo_convFilter) getConvExpName(nodeId LnsInt,argListNode *Nodes_ExpListNode) string {
    return Lns_getVM().String_format("%s_convExp%d", []LnsAny{Lns_car(Lns_getVM().String_gsub(self.moduleTypeInfo.FP.Get_rawTxt(),"@", "")).(string), nodeId})
}

// 1016: decl @lune.@base.@convGo.convFilter.processConvExp
func (self *convGo_convFilter) processConvExp(nodeId LnsInt,dstTypeList *LnsList,argListNode LnsAny) {
    var argList *Nodes_ExpListNode
    
    {
        _argList := argListNode
        if _argList == nil{
            return 
        } else {
            argList = _argList.(*Nodes_ExpListNode)
        }
    }
    if convGo_getExpListKind_1321_(dstTypeList, argList) != convGo_ExpListKind__Conv{
        return 
    }
    var expList *LnsList
    expList = argList.FP.Get_expList()
    var mRetIndex LnsAny
    mRetIndex = Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(argList.FP.Get_mRetExp()) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_MRetExp).FP.Get_index()}))
    if Lns_op_not(mRetIndex){
        if expList.GetAt(expList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr{
        } else { 
            return 
        }
    }
    var workList *LnsList
    workList = NewLnsList([]LnsAny{})
    for _, _exp := range( expList.Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        var workExp *Nodes_Node
        workExp = Nodes_getUnwraped(exp)
        if workExp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr{
            break
        }
        workList.Insert(Nodes_Node2Stem(workExp))
    }
    expList = workList
    
    self.FP.Writeln(Lns_getVM().String_format("// for %d", []LnsAny{argList.FP.Get_pos().LineNo}))
    self.FP.Write(Lns_getVM().String_format("func %s(", []LnsAny{self.FP.getConvExpName(nodeId, argList)}))
    for _index, _argExp := range( expList.Items ) {
        index := _index + 1
        argExp := _argExp.(Nodes_NodeDownCast).ToNodes_Node()
        {
            _exp2ddd := Nodes_ExpToDDDNodeDownCastF(argExp.FP)
            if _exp2ddd != nil {
                exp2ddd := _exp2ddd.(*Nodes_ExpToDDDNode)
                for _, _exp := range( exp2ddd.FP.Get_expList().FP.Get_expList().Items ) {
                    exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
                    if index != 1{
                        self.FP.Write(", ")
                    }
                    self.FP.Write(Lns_getVM().String_format("arg%d ", []LnsAny{index}))
                    self.FP.Write(self.FP.type2gotype(exp.FP.Get_expType()))
                }
            } else {
                if index != 1{
                    self.FP.Write(", ")
                }
                if mRetIndex == index{
                    self.FP.Write(Lns_getVM().String_format("arg%d []LnsAny", []LnsAny{index}))
                    break
                } else { 
                    self.FP.Write(Lns_getVM().String_format("arg%d ", []LnsAny{index}))
                    {
                        _castNode := Nodes_ExpCastNodeDownCastF(argExp.FP)
                        if _castNode != nil {
                            castNode := _castNode.(*Nodes_ExpCastNode)
                            self.FP.Write(self.FP.type2gotype(castNode.FP.Get_castType()))
                        } else {
                            self.FP.Write(self.FP.type2gotype(argExp.FP.Get_expType()))
                        }
                    }
                }
            }
        }
    }
    self.FP.Write(") ")
    var getRetType func(retType *Ast_TypeInfo,index LnsInt) *Ast_TypeInfo
    getRetType = func(retType *Ast_TypeInfo,index LnsInt) *Ast_TypeInfo {
        if retType == Ast_builtinTypeEmpty{
            return argList.FP.GetExpTypeNoDDDAt(index)
        }
        return retType
    }
    var retTypeList *LnsList
    retTypeList = NewLnsList([]LnsAny{})
    for _index, _dstType := range( dstTypeList.Items ) {
        index := _index + 1
        dstType := _dstType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        retTypeList.Insert(Ast_TypeInfo2Stem(getRetType(dstType, index)))
    }
    if retTypeList.Len() >= 2{
        self.FP.Write("(")
        for _index, _retType := range( retTypeList.Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(", ")
            }
            self.FP.Write(self.FP.type2gotype(getRetType(retType, index)))
        }
        self.FP.Writeln(") {")
    } else if retTypeList.Len() == 1{
        self.FP.Write(self.FP.type2gotype(getRetType(retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), 1)))
        self.FP.Writeln(" {")
    } else { 
        self.FP.Writeln(" {")
    }
    self.FP.Write("    return ")
    if mRetIndex != nil{
        mRetIndex_6157 := mRetIndex.(LnsInt)
        var restIndex LnsAny
        restIndex = nil
        for _index, _retType := range( retTypeList.Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(", ")
            }
            if retType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                restIndex = index
                
                break
            }
            if index >= mRetIndex_6157{
                var wrote bool
                wrote = false
                if index <= expList.Len(){
                    var exp *Nodes_Node
                    exp = expList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                    {
                        _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
                        if _castNode != nil {
                            castNode := _castNode.(*Nodes_ExpCastNode)
                            var statNode *Nodes_ConvStatNode
                            statNode = Nodes_ConvStatNode_create(self.nodeManager, exp.FP.Get_pos(), false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp.FP.Get_expType())}), Lns_getVM().String_format("Lns_getFromMulti( arg%d, 0 )", []LnsAny{index}))
                            self.FP.outputImplicitCast(castNode.FP.Get_castType(), &statNode.Nodes_Node, castNode)
                            wrote = true
                            
                        }
                    }
                }
                if Lns_op_not(wrote){
                    self.FP.Write(Lns_getVM().String_format("Lns_getFromMulti( arg%d, %d )", []LnsAny{mRetIndex_6157, index - mRetIndex_6157}))
                    self.FP.outputAny2Type(retType)
                }
            } else { 
                self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
            }
        }
        if restIndex != nil{
            restIndex_6174 := restIndex.(LnsInt)
            self.FP.Write("Lns_2DDD( ")
            for _index, _ := range( expList.Items ) {
                index := _index + 1
                if index >= restIndex_6174{
                    if index < expList.Len(){
                        self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
                    } else { 
                        self.FP.Write(Lns_getVM().String_format("arg%d[%d:]", []LnsAny{mRetIndex_6157, index - mRetIndex_6157}))
                        break
                    }
                }
            }
            self.FP.Writeln(")")
        } else {
            self.FP.Writeln("")
        }
    } else {
        for _index, _ := range( retTypeList.Items ) {
            index := _index + 1
            if index != 1{
                self.FP.Write(", ")
            }
            if index <= expList.Len(){
                self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
            } else { 
                self.FP.Write("nil")
            }
        }
        self.FP.Writeln("")
    }
    self.FP.Writeln("}")
}

// 1181: decl @lune.@base.@convGo.convFilter.outputNilAccCall
func (self *convGo_convFilter) outputNilAccCall(node *Nodes_ExpCallNode) {
    if Lns_op_not(node.FP.HasNilAccess()){
        return 
    }
    if node.FP.Get_expTypeList().Len() > convGo_MaxNilAccNum{
        var anys string
        anys = "LnsAny"
        var nils string
        nils = "nil"
        var lists string
        lists = "list[0]"
        {
            var _from5690 LnsInt = 2
            var _to5690 LnsInt = node.FP.Get_expTypeList().Len()
            for _work5690 := _from5690; _work5690 <= _to5690; _work5690++ {
                count := _work5690
                anys = Lns_getVM().String_format("%s,LnsAny", []LnsAny{anys})
                
                nils = Lns_getVM().String_format("%s,nil", []LnsAny{nils})
                
                lists = Lns_getVM().String_format("%s,list[%d]", []LnsAny{lists, count - 1})
                
            }
        }
        var name string
        name = Lns_getVM().String_format("%s_%d", []LnsAny{Lns_car(Lns_getVM().String_gsub(self.moduleTypeInfo.FP.Get_rawTxt(),"@", "")).(string), node.FP.Get_id()})
        self.FP.Write(Lns_getVM().String_format("func lns_NilAccCall_%s( env *LnsEnv, call func () (%s) ) bool {\n    return env.NilAccPush( Lns_2DDD( call() ) )\n}\nfunc lns_NilAccFinCall_%s( ret LnsAny ) (%s) {\n    if Lns_IsNil( ret ) {\n        return %s\n    }\n    list := ret.([]LnsAny)\n    return %s\n}\n", []LnsAny{name, anys, name, anys, nils, lists}))
    }
}

// 1222: decl @lune.@base.@convGo.convFilter.processGenericsCall
func (self *convGo_convFilter) processGenericsCall(node *Nodes_ExpCallNode) {
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(convGo_isRetGenerics_1358_(node))) ||
        Lns_GetEnv().SetStackVal( node.FP.Get_expTypeList().Len() < 2) ).(bool){
        return 
    }
    var srcTypeList *LnsList
    srcTypeList = node.FP.Get_func().FP.Get_expType().FP.Get_retTypeInfoList()
    var dstTypeList *LnsList
    dstTypeList = node.FP.Get_expTypeList()
    var srcTxt string
    srcTxt = Lns_getVM().String_format("arg1 %s", []LnsAny{self.FP.type2gotype(srcTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
    var dstTxt string
    dstTxt = Lns_getVM().String_format("%s", []LnsAny{self.FP.type2gotype(dstTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
    {
        var _from5898 LnsInt = 2
        var _to5898 LnsInt = srcTypeList.Len()
        for _work5898 := _from5898; _work5898 <= _to5898; _work5898++ {
            index := _work5898
            srcTxt = Lns_getVM().String_format("%s,arg%d %s", []LnsAny{srcTxt, index, self.FP.type2gotype(srcTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
            
        }
    }
    {
        var _from5930 LnsInt = 2
        var _to5930 LnsInt = dstTypeList.Len()
        for _work5930 := _from5930; _work5930 <= _to5930; _work5930++ {
            index := _work5930
            dstTxt = Lns_getVM().String_format("%s,%s", []LnsAny{dstTxt, self.FP.type2gotype(dstTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
            
        }
    }
    self.FP.Writeln(Lns_getVM().String_format("func %s(%s) (%s) {", []LnsAny{self.FP.getConvGenericsName(&node.Nodes_Node), srcTxt, dstTxt}))
    self.FP.PushIndent(nil)
    self.FP.Write("return ")
    for _index, _dstType := range( dstTypeList.Items ) {
        index := _index + 1
        dstType := _dstType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(", ")
        }
        if index > srcTypeList.Len(){
            self.FP.Write("nil")
        } else { 
            self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
            var srcType *Ast_TypeInfo
            srcType = srcTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if srcType.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
                self.FP.outputAny2Type(dstType)
            }
        }
    }
    self.FP.Writeln("")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 1266: decl @lune.@base.@convGo.convFilter.outputRetType
func (self *convGo_convFilter) outputRetType(retTypeList *LnsList) {
    if _switch6194 := retTypeList.Len(); _switch6194 == 0 {
        self.FP.Write("")
    } else if _switch6194 == 1 {
        if retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeNeverRet{
            self.FP.Write(" " + self.FP.type2gotype(retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))
        }
    } else {
        self.FP.Write("(")
        for _index, _retType := range( retTypeList.Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(", ")
            }
            self.FP.Write(self.FP.type2gotype(retType))
        }
        self.FP.Write(")")
    }
}

// 1299: decl @lune.@base.@convGo.convFilter.outputDeclFunc
func (self *convGo_convFilter) OutputDeclFunc(funcInfo LnsAny) *convGo_FuncConv {
    var typeInfo *Ast_TypeInfo
    var name LnsAny
    var prefixType *Ast_TypeInfo
    var extFlag bool
    switch _exp6383 := funcInfo.(type) {
    case *convGo_FuncInfo__DeclInfo:
    node := _exp6383.Val1
    workDeclInfo := _exp6383.Val2
        extFlag = false
        
        typeInfo = node.FP.Get_expType()
        
        prefixType = typeInfo.FP.Get_parentInfo()
        
        if Lns_op_not(workDeclInfo.FP.Get_name()){
            if self.processMode == convGo_ProcessMode__NonClosureFuncDecl{
                name = "_anonymous"
                
            } else { 
                name = nil
                
            }
        } else { 
            name = typeInfo.FP.Get_rawTxt()
            
        }
    case *convGo_FuncInfo__Type:
    workTypeInfo := _exp6383.Val1
        extFlag = workTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__Ext
        
        typeInfo = workTypeInfo
        
        prefixType = typeInfo.FP.Get_parentInfo()
        
        name = typeInfo.FP.Get_rawTxt()
        
    case *convGo_FuncInfo__WithClass:
    classType := _exp6383.Val1
    methodType := _exp6383.Val2
        extFlag = false
        
        typeInfo = methodType
        
        prefixType = classType
        
        name = typeInfo.FP.Get_rawTxt()
        
    }
    if convGo_isClosure_1177_(typeInfo){
        self.FP.Write("func")
    } else { 
        if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__Method{
            self.FP.Write("func ")
            self.FP.Write("(self *")
            self.FP.Write(self.FP.getTypeSymbol(prefixType))
            self.FP.Write(") ")
        } else { 
            self.FP.Write("func ")
        }
        if typeInfo.FP.Get_extedType().FP.Get_kind() != Ast_TypeInfoKind__FormFunc{
            if name != nil{
                name_6288 := name.(string)
                self.FP.outputSymbol(&convGo_SymbolKind__Func{typeInfo}, name_6288)
            }
        }
    }
    self.FP.Write("(")
    var workType *Ast_TypeInfo
    
    {
        _workType := typeInfo.FP.GetOverridingType()
        if _workType == nil{
            workType = typeInfo
            
        } else {
            workType = _workType.(*Ast_TypeInfo)
        }
    }
    var retTypeList *LnsList
    if extFlag{
        retTypeList = Lns_unwrap( Lns_car(Ast_convToExtTypeList(self.processInfo, workType.FP.Get_retTypeInfoList()))).(*LnsList)
        
    } else { 
        retTypeList = workType.FP.Get_retTypeInfoList()
        
    }
    var funcConv *convGo_FuncConv
    funcConv = NewconvGo_FuncConv(retTypeList)
    switch _exp6767 := funcInfo.(type) {
    case *convGo_FuncInfo__DeclInfo:
    node := _exp6767.Val1
    declInfo := _exp6767.Val2
        for _index, _arg := range( declInfo.FP.Get_argList().Items ) {
            index := _index + 1
            arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
            if index != 1{
                self.FP.Write(",")
            }
            var argType *Ast_TypeInfo
            argType = workType.FP.Get_argTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if argType.FP.Get_nonnilableType().FP.Get_kind() == Ast_TypeInfoKind__Alternate{
                {
                    _argNode := Nodes_DeclArgNodeDownCastF(arg.FP)
                    if _argNode != nil {
                        argNode := _argNode.(*Nodes_DeclArgNode)
                        var argName string
                        argName = self.FP.getSymbolSym(argNode.FP.Get_symbolInfo())
                        self.FP.Write(Lns_getVM().String_format("_%s ", []LnsAny{argName}))
                        self.FP.Write(self.FP.type2gotype(argType))
                        funcConv.FP.Get_argList().Insert(Ast_SymbolInfo2Stem(argNode.FP.Get_symbolInfo()))
                    } else {
                        convGo_filter_1171_(arg, self, node)
                    }
                }
            } else { 
                convGo_filter_1171_(arg, self, node)
            }
        }
    case *convGo_FuncInfo__Type:
        for _index, _argType := range( workType.FP.Get_argTypeInfoList().Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(",")
            }
            self.FP.Write(Lns_getVM().String_format("arg%d %s", []LnsAny{index, self.FP.type2gotype(argType)}))
        }
    case *convGo_FuncInfo__WithClass:
        for _index, _argType := range( workType.FP.Get_argTypeInfoList().Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(",")
            }
            self.FP.Write(Lns_getVM().String_format("arg%d %s", []LnsAny{index, self.FP.type2gotype(argType)}))
        }
    }
    self.FP.Write(")")
    self.FP.outputRetType(retTypeList)
    return funcConv
}

// 1417: decl @lune.@base.@convGo.convFilter.outputConvToForm
func (self *convGo_convFilter) outputConvToForm(node *Nodes_ExpCastNode) {
    var castType *Ast_TypeInfo
    castType = node.FP.Get_castType().FP.Get_nonnilableType().FP.Get_extedType()
    if castType.FP.Get_kind() != Ast_TypeInfoKind__Form{
        return 
    }
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_exp().FP.Get_expType().FP.Get_extedType()
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( node.FP.Get_exp().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Ext) &&
        Lns_GetEnv().SetStackVal( funcType.FP.Get_srcTypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Form) ).(bool)){
        self.FP.Writeln(Lns_getVM().String_format("      \nfunc %s( luaform LnsAny ) LnsForm {\n    return func (argList []LnsAny) []LnsAny {\n        return %s.RunLoadedfunc( luaform.(*Lns_luaValue), argList )\n    }\n}", []LnsAny{self.FP.getConv2formName(&node.Nodes_Node), self.env.FP.GetLuavm(node.FP.Get_exp().FP.IsThreading())}))
        return 
    }
    self.FP.Writeln(Lns_getVM().String_format("// for %d: %s", []LnsAny{node.FP.Get_pos().LineNo, Nodes_getNodeKindName(node.FP.Get_kind())}))
    self.FP.Write(Lns_getVM().String_format("func %s( src func (", []LnsAny{self.FP.getConv2formName(&node.Nodes_Node)}))
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList().Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(", ")
        }
        self.FP.Write(Lns_getVM().String_format("arg%d %s", []LnsAny{index, self.FP.type2gotype(argType)}))
    }
    self.FP.Write(")")
    self.FP.outputRetType(funcType.FP.Get_retTypeInfoList())
    self.FP.Writeln(") LnsForm {")
    self.FP.PushIndent(nil)
    self.FP.Writeln("return func (argList []LnsAny) []LnsAny {")
    self.FP.PushIndent(nil)
    if funcType.FP.Get_retTypeInfoList().Len() > 0{
        self.FP.Write("return ")
        if funcType.FP.Get_argTypeInfoList().Len() > 0{
            self.FP.Write("Lns_2DDD(")
        }
    }
    self.FP.Write("src(")
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList().Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(", ")
        }
        if argType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
            self.FP.Write(Lns_getVM().String_format("argList[ %d: ]", []LnsAny{index - 1}))
            break
        }
        self.FP.Write(Lns_getVM().String_format("Lns_getFromMulti( argList, %d )", []LnsAny{index - 1}))
    }
    self.FP.Write(")")
    if funcType.FP.Get_retTypeInfoList().Len() > 0{
        if funcType.FP.Get_argTypeInfoList().Len() > 0{
            self.FP.Writeln(")")
        } else { 
            self.FP.Writeln("")
        }
    } else { 
        self.FP.Writeln("")
        self.FP.Writeln("return []LnsAny{}")
    }
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 1490: decl @lune.@base.@convGo.convFilter.processConvStat
func (self *convGo_convFilter) ProcessConvStat(node *Nodes_ConvStatNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_txt())
}

// 1496: decl @lune.@base.@convGo.convFilter.outputTopScopeVar
func (self *convGo_convFilter) outputTopScopeVar(node *Nodes_DeclVarNode) {
    for _, _symbolInfo := range( node.FP.Get_symbolInfoList().Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_scope() == self.moduleScope) &&
            Lns_GetEnv().SetStackVal( node.FP.Get_mode() == Nodes_DeclVarMode__Let) ).(bool)){
            self.FP.Writeln(Lns_getVM().String_format("var %s %s", []LnsAny{self.FP.getSymbolSym(symbolInfo), self.FP.type2gotype(symbolInfo.FP.Get_typeInfo())}))
        }
    }
}

// 1505: decl @lune.@base.@convGo.convFilter.outputConvExt
func (self *convGo_convFilter) outputConvExt(funcNode *Nodes_Node) {
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(funcNode.FP)
        if _fieldNode != nil {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            if fieldNode.FP.Get_prefix().FP.Get_expType().FP.Get_nonnilableType().FP.Get_kind() != Ast_TypeInfoKind__Ext{
                return 
            }
        } else {
            return 
        }
    }
    self.FP.Write(Lns_getVM().String_format("func Lns_callExt%d( args []LnsAny ) (", []LnsAny{funcNode.FP.Get_id()}))
    for _index, _retType := range( funcNode.FP.Get_expType().FP.Get_retTypeInfoList().Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(",")
        }
        self.FP.Write(self.FP.type2gotype(retType))
    }
    self.FP.Writeln(") {")
    self.FP.Write("    return ")
    for _index, _ := range( funcNode.FP.Get_expType().FP.Get_retTypeInfoList().Items ) {
        index := _index + 1
        if index > 1{
            self.FP.Write(",")
        }
        self.FP.Write(Lns_getVM().String_format("Lns_getFromMulti( args, %d )", []LnsAny{index - 1}))
    }
    self.FP.Writeln("")
    self.FP.Writeln("}")
}

// 1532: decl @lune.@base.@convGo.convFilter.outputModule
func (self *convGo_convFilter) outputModule(moduleTypeInfo *Ast_TypeInfo,addDot bool) {
    self.FP.Write(self.FP.getModuleSym(moduleTypeInfo, addDot))
}

// 1547: decl @lune.@base.@convGo.convFilter.outputModuleImport
func (self *convGo_convFilter) outputModuleImport(node *Nodes_DeclClassNode) {
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( node.FP.Get_expType().FP.Get_kind() != Ast_TypeInfoKind__ExtModule) ||
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( node.FP.Get_lang() != Types_Lang__Go) &&
            Lns_GetEnv().SetStackVal( node.FP.Get_lang() != Types_Lang__Same) ).(bool))) ||
        Lns_GetEnv().SetStackVal( self.FP.isSamePackageExtModule(node.FP.Get_expType())) ).(bool){
        return 
    }
    var normalType *Ast_NormalTypeInfo
    normalType = Lns_unwrap( Ast_NormalTypeInfoDownCastF(node.FP.Get_expType().FP)).(*Ast_NormalTypeInfo)
    self.FP.Write(Lns_getVM().String_format("import %s ", []LnsAny{node.FP.Get_expType().FP.Get_rawTxt()}))
    var mod string
    mod = convGo_convExp7610(Lns_2DDD(Lns_getVM().String_gsub(normalType.FP.Get_requirePath(),"%.[^%.]+$", "")))
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(mod,"^go/", nil, nil))){
        var workMod string
        workMod = convGo_convExp7649(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(mod,"^go/", "")).(string),"%.", "/")).(string),":", ".")))
        self.FP.Writeln(Lns_getVM().String_format("\"%s\"", []LnsAny{workMod}))
    } else { 
        var path string
        path = convGo_convExp7687(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(normalType.FP.Get_requirePath(),"%.", "/")).(string),":", ".")))
        self.FP.Writeln(Lns_getVM().String_format("\"%s/%s\"", []LnsAny{self.option.FP.Get_appName(), path}))
    }
}

// 1569: decl @lune.@base.@convGo.convFilter.outputImport
func (self *convGo_convFilter) outputImport(node *Nodes_ImportNode) {
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.FP.isSameModDir(node.FP.Get_moduleTypeInfo())) ||
        Lns_GetEnv().SetStackVal( Ast_isBuiltin(node.FP.Get_moduleTypeInfo().FP.Get_typeId())) ).(bool){
        return 
    }
    self.FP.Write("import ")
    var packSym string
    packSym = node.FP.Get_assignName()
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(node.FP.Get_modulePath(),"^go/", nil, nil))){
        var workMod string
        workMod = convGo_convExp7790(Lns_2DDD(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(node.FP.Get_modulePath(),"^go/", "")).(string),"%.", "/")).(string),":", ".")))
        self.FP.Writeln(Lns_getVM().String_format("%s \"%s\"", []LnsAny{packSym, Lns_car(Lns_getVM().String_gsub(workMod,"/[^/]+$", "")).(string)}))
    } else { 
        var modulePath string
        var count LnsInt
        modulePath,count = Lns_getVM().String_gsub(node.FP.Get_modulePath(),"([^%.]+)%.[^%.]+$", "/%1")
        if count == 0{
            self.FP.Writeln(Lns_getVM().String_format("%s \"%s\"", []LnsAny{packSym, self.option.appName}))
        } else { 
            var modDir string
            modDir = convGo_convExp7869(Lns_2DDD(Lns_getVM().String_gsub(modulePath,"/", "")))
            self.FP.Writeln(Lns_getVM().String_format("%s \"%s/%s\"", []LnsAny{packSym, self.option.appName, Lns_car(Lns_getVM().String_gsub(modDir,"%.", "/")).(string)}))
        }
    }
    self.module2PackSym.Set(node.FP.Get_moduleTypeInfo(),packSym)
}

// 1596: decl @lune.@base.@convGo.convFilter.processRoot
func (self *convGo_convFilter) ProcessRoot(node *Nodes_RootNode,_opt LnsAny) {
    for _pragma := range( node.FP.Get_luneHelperInfo().PragmaSet.Items ) {
        pragma := _pragma
        switch _exp7975 := pragma.(type) {
        case *LuneControl_Pragma__limit_conv_code:
        codeSet := _exp7975.Val1
            if Lns_op_not(codeSet.Has(LuneControl_Code__Go)){
                self.FP.Writeln("// This code is transcompiled by LuneScript.")
                self.FP.Writeln(Lns_getVM().String_format("package %s", []LnsAny{self.option.packageName}))
                return 
            }
        }
    }
    var isUsingInTest func(aNode *Nodes_Node) bool
    isUsingInTest = func(aNode *Nodes_Node) bool {
        for _, _testBlock := range( node.FP.Get_nodeManager().FP.GetTestBlockNodeList().Items ) {
            testBlock := _testBlock.(Nodes_TestBlockNodeDownCast).ToNodes_TestBlockNode()
            if testBlock.FP.IsInnerPos(aNode.FP.Get_pos()){
                return true
            }
        }
        return false
    }
    var builtinFuncs *TransUnit_BuiltinFuncType
    builtinFuncs = TransUnit_getBuiltinFunc()
    var builtin2runtime *LnsMap
    builtin2runtime = NewLnsMap( map[LnsAny]LnsAny{builtinFuncs.Str_gsub:"GETVM.String_gsub",builtinFuncs.String_gsub:"GETVM.String_gsub",builtinFuncs.Str_find:"GETVM.String_find",builtinFuncs.String_find:"GETVM.String_find",builtinFuncs.Str_byte:"GETVM.String_byte",builtinFuncs.String_byte:"GETVM.String_byte",builtinFuncs.Str_format:"GETVM.String_format",builtinFuncs.String_format:"GETVM.String_format",builtinFuncs.Str_rep:"GETVM.String_rep",builtinFuncs.String_rep:"GETVM.String_rep",builtinFuncs.Str_gmatch:"GETVM.String_gmatch",builtinFuncs.String_gmatch:"GETVM.String_gmatch",builtinFuncs.Str_sub:"GETVM.String_sub",builtinFuncs.String_sub:"GETVM.String_sub",builtinFuncs.Str_lower:"GETVM.String_lower",builtinFuncs.String_lower:"GETVM.String_lower",builtinFuncs.Str_upper:"GETVM.String_upper",builtinFuncs.String_upper:"GETVM.String_upper",builtinFuncs.Str_reverse:"GETVM.String_reverse",builtinFuncs.String_reverse:"GETVM.String_reverse",Ast_builtinTypeNone:"",})
    
    builtin2runtime.Set(builtinFuncs.Lns_error,"panic")
    builtin2runtime.Set(builtinFuncs.Lns_print,"Lns_print")
    builtin2runtime.Set(builtinFuncs.Lns_type,"Lns_type")
    builtin2runtime.Set(builtinFuncs.Lns_require,"Lns_require")
    builtin2runtime.Set(builtinFuncs.Lns_tonumber,"Lns_tonumber")
    builtin2runtime.Set(builtinFuncs.Lns__load,"GETVM.Load")
    builtin2runtime.Set(builtinFuncs.Lns_loadfile,"GETVM.Loadfile")
    builtin2runtime.Set(builtinFuncs.Lns_expandLuavalMap,"GETVM.ExpandLuavalMap")
    builtin2runtime.Set(builtinFuncs.String_dump,"GETVM.String_dump")
    builtin2runtime.Set(builtinFuncs.Io_open,"Lns_io_open")
    builtin2runtime.Set(builtinFuncs.Io_popen,"GETVM.IO_popen")
    builtin2runtime.Set(builtinFuncs.Package_searchpath,"GETVM.Package_searchpath")
    builtin2runtime.Set(builtinFuncs.Os_clock,"GETVM.OS_clock")
    builtin2runtime.Set(builtinFuncs.Os_exit,"GETVM.OS_exit")
    builtin2runtime.Set(builtinFuncs.Os_remove,"GETVM.OS_remove")
    builtin2runtime.Set(builtinFuncs.Os_date,"GETVM.OS_date")
    builtin2runtime.Set(builtinFuncs.Os_time,"GETVM.OS_time")
    builtin2runtime.Set(builtinFuncs.Os_difftime,"GETVM.OS_difftime")
    builtin2runtime.Set(builtinFuncs.Os_rename,"GETVM.OS_rename")
    builtin2runtime.Set(builtinFuncs.Math_random,"GETVM.Math_random")
    builtin2runtime.Set(builtinFuncs.Math_randomseed,"GETVM.Math_randomseed")
    self.builtin2runtime = builtin2runtime
    
    self.type2gotypeMap = NewLnsMap( map[LnsAny]LnsAny{Ast_builtinTypeInt:"LnsInt",Ast_builtinTypeReal:"LnsReal",Ast_builtinTypeStem:"LnsAny",Ast_builtinTypeString:"string",Ast_builtinTypeBool:"bool",builtinFuncs.Ostream_:"Lns_oStream",builtinFuncs.Istream_:"Lns_iStream",builtinFuncs.Luastream_:"Lns_luaStream",})
    
    self.FP.Writeln("// This code is transcompiled by LuneScript.")
    self.FP.Writeln(Lns_getVM().String_format("package %s", []LnsAny{self.option.packageName}))
    self.FP.Writeln("import . \"github.com/ifritJP/LuneScript/src/lune/base/runtime_go\"")
    for _, _workNode := range( node.FP.Get_nodeManager().FP.GetImportNodeList().Items ) {
        workNode := _workNode.(Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        self.FP.outputImport(workNode)
    }
    for _, _workNode := range( node.FP.Get_nodeManager().FP.GetDeclClassNodeList().Items ) {
        workNode := _workNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        self.FP.outputModuleImport(workNode)
    }
    var initModVar string
    initModVar = Lns_getVM().String_format("init_%s", []LnsAny{convGo_getModuleName_1198_(node.FP.Get_moduleTypeInfo())})
    self.FP.Writeln(Lns_getVM().String_format("var %s bool", []LnsAny{initModVar}))
    self.FP.pushProcessMode(convGo_ProcessMode__DeclTopScopeVar)
    var modSym *Ast_SymbolInfo
    modSym = Lns_unwrap( self.moduleScope.FP.GetSymbolInfoChild("__mod__")).(*Ast_SymbolInfo)
    self.FP.Writeln(Lns_getVM().String_format("var %s string", []LnsAny{self.FP.getSymbolSym(modSym)}))
    {
        var procNode func(workNode *Nodes_Node)
        procNode = func(workNode *Nodes_Node) {
            convGo_filter_1171_(workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetDeclEnumNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_DeclEnumNodeDownCast).ToNodes_DeclEnumNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(&tmpNode.Nodes_Node)
            }
        }
    }
    
    for _, _workNode := range( node.FP.Get_nodeManager().FP.GetDeclVarNodeList().Items ) {
        workNode := _workNode.(Nodes_DeclVarNodeDownCast).ToNodes_DeclVarNode()
        self.FP.outputTopScopeVar(workNode)
    }
    self.FP.popProcessMode()
    {
        var procNode func(workNode *Nodes_Node)
        procNode = func(workNode *Nodes_Node) {
            convGo_filter_1171_(workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetDeclAlgeNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(&tmpNode.Nodes_Node)
            }
        }
    }
    
    {
        var procNode func(workNode *Nodes_Node)
        procNode = func(workNode *Nodes_Node) {
            convGo_filter_1171_(workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetDeclFormNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_DeclFormNodeDownCast).ToNodes_DeclFormNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(&tmpNode.Nodes_Node)
            }
        }
    }
    
    {
        var procNode func(workNode *Nodes_ExpCallNode)
        procNode = func(workNode *Nodes_ExpCallNode) {
            self.FP.processGenericsCall(workNode)
            self.FP.outputNilAccCall(workNode)
            self.FP.outputConvExt(workNode.FP.Get_func())
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetExpCallNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCallNodeDownCast).ToNodes_ExpCallNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(tmpNode)
            }
        }
    }
    
    {
        var procNode func(workNode *Nodes_ExpCastNode)
        procNode = func(workNode *Nodes_ExpCastNode) {
            self.FP.outputConvToForm(workNode)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetExpCastNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCastNodeDownCast).ToNodes_ExpCastNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(tmpNode)
            }
        }
    }
    
    {
        var procNode func(workNode *Nodes_Node)
        procNode = func(workNode *Nodes_Node) {
            convGo_filter_1171_(workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetTestCaseNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_TestCaseNodeDownCast).ToNodes_TestCaseNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(&tmpNode.Nodes_Node)
            }
        }
    }
    
    {
        var procNode func(workNode *Nodes_IfUnwrapNode)
        procNode = func(workNode *Nodes_IfUnwrapNode) {
            var symTypeList *LnsList
            symTypeList = NewLnsList([]LnsAny{})
            {
                var _from9459 LnsInt = 1
                var _to9459 LnsInt = workNode.FP.Get_varSymList().Len()
                for _work9459 := _from9459; _work9459 <= _to9459; _work9459++ {
                    symTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeStem_))
                }
            }
            self.FP.processConvExp(workNode.FP.Get_id(), symTypeList, workNode.FP.Get_expList())
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetIfUnwrapNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_IfUnwrapNodeDownCast).ToNodes_IfUnwrapNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(tmpNode)
            }
        }
    }
    
    {
        var procNode func(workNode *Nodes_ExpSetValNode)
        procNode = func(workNode *Nodes_ExpSetValNode) {
            self.FP.processConvExp(workNode.FP.Get_id(), workNode.FP.Get_exp1().FP.Get_expTypeList(), workNode.FP.Get_exp2())
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetExpSetValNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_ExpSetValNodeDownCast).ToNodes_ExpSetValNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(tmpNode)
            }
        }
    }
    
    {
        var procNode func(workNode *Nodes_ExpCallNode)
        procNode = func(workNode *Nodes_ExpCallNode) {
            self.FP.processConvExp(workNode.FP.Get_id(), workNode.FP.Get_func().FP.Get_expType().FP.Get_argTypeInfoList(), workNode.FP.Get_argList())
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetExpCallNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCallNodeDownCast).ToNodes_ExpCallNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(tmpNode)
            }
        }
    }
    
    {
        var procNode func(workNode *Nodes_DeclVarNode)
        procNode = func(workNode *Nodes_DeclVarNode) {
            var typeList *LnsList
            typeList = NewLnsList([]LnsAny{})
            {
                _expList := workNode.FP.Get_expList()
                if _expList != nil {
                    expList := _expList.(*Nodes_ExpListNode)
                    for _index, _symbolInfo := range( workNode.FP.Get_symbolInfoList().Items ) {
                        index := _index + 1
                        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( workNode.FP.Get_mode() == Nodes_DeclVarMode__Let) ||
                            Lns_GetEnv().SetStackVal( workNode.FP.Get_mode() == Nodes_DeclVarMode__Unwrap) ).(bool){
                            if workNode.FP.Get_unwrapFlag(){
                                typeList.Insert(Ast_TypeInfo2Stem(expList.FP.GetExpTypeNoDDDAt(index)))
                            } else { 
                                typeList.Insert(Ast_TypeInfo2Stem(symbolInfo.FP.Get_typeInfo()))
                            }
                        }
                    }
                    self.FP.processConvExp(workNode.FP.Get_id(), typeList, expList)
                }
            }
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetDeclVarNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_DeclVarNodeDownCast).ToNodes_DeclVarNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(tmpNode)
            }
        }
    }
    
    {
        var procNode func(workNode *Nodes_AliasNode)
        procNode = func(workNode *Nodes_AliasNode) {
            self.FP.Writeln(Lns_getVM().String_format("type %s = %s", []LnsAny{self.FP.getSymbolSym(workNode.FP.Get_newSymbol()), self.FP.getTypeSymbol(workNode.FP.Get_typeInfo().FP.Get_aliasSrc())}))
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetAliasNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(tmpNode)
            }
        }
    }
    
    self.FP.pushProcessMode(convGo_ProcessMode__NonClosureFuncDecl)
    {
        var procNode func(workNode *Nodes_Node)
        procNode = func(workNode *Nodes_Node) {
            convGo_filter_1171_(workNode, self, &node.Nodes_Node)
            self.FP.Writeln("")
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetDeclFuncNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(&tmpNode.Nodes_Node)
            }
        }
    }
    
    self.FP.popProcessMode()
    self.FP.pushProcessMode(convGo_ProcessMode__DeclClass)
    {
        var procNode func(workNode *Nodes_Node)
        procNode = func(workNode *Nodes_Node) {
            convGo_filter_1171_(workNode, self, &node.Nodes_Node)
            self.FP.Writeln("")
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager().FP.GetDeclClassNodeList().Items ) {
            tmpNode := _tmpNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( self.enableTest) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(isUsingInTest(&tmpNode.Nodes_Node))) ).(bool){
                procNode(&tmpNode.Nodes_Node)
            }
        }
    }
    
    self.FP.popProcessMode()
    self.FP.Writeln(Lns_getVM().String_format("func Lns_%s_init() {", []LnsAny{convGo_getModuleName_1198_(node.FP.Get_moduleTypeInfo())}))
    self.FP.PushIndent(nil)
    self.FP.Writeln(Lns_getVM().String_format("if %s { return }", []LnsAny{initModVar}))
    self.FP.Writeln(Lns_getVM().String_format("%s = true", []LnsAny{initModVar}))
    self.FP.Writeln(Lns_getVM().String_format("%s = \"%s\"", []LnsAny{self.FP.getSymbolSym(modSym), node.FP.Get_moduleTypeInfo().FP.GetFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), nil)}))
    self.FP.Writeln("Lns_InitMod()")
    var modulePath string
    modulePath = convGo_convExp10445(Lns_2DDD(Lns_getVM().String_gsub(node.FP.Get_moduleTypeInfo().FP.GetFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), nil),"@", "")))
    var moduleName string
    moduleName = convGo_getModuleName_1198_(node.FP.Get_moduleTypeInfo())
    if self.enableTest{
        for _, _workNode := range( node.FP.Get_nodeManager().FP.GetTestCaseNodeList().Items ) {
            workNode := _workNode.(Nodes_TestCaseNodeDownCast).ToNodes_TestCaseNode()
            self.FP.Writeln(Lns_getVM().String_format("Testing_registerTestcase( \"%s\", \"%s\", lns_test_%s_%s )", []LnsAny{modulePath, workNode.FP.Get_name().Txt, moduleName, workNode.FP.Get_name().Txt}))
        }
    }
    for _, _child := range( node.FP.Get_children().Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(convGo_ignoreNodeInInnerBlockSet.Has(child.FP.Get_kind())){
            convGo_filter_1171_(child, self, &node.Nodes_Node)
        }
    }
    self.FP.PopIndent()
    self.FP.Writeln("}")
    if self.option.mainModule == self.FP.getCanonicalName(self.moduleTypeInfo, false){
        var hasMain bool
        hasMain = false
        for _, _workNode := range( node.FP.Get_nodeManager().FP.GetDeclFuncNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            if convGo_isMain_1062_(workNode.FP.Get_expType()){
                hasMain = true
                
                break
            }
        }
        if Lns_op_not(hasMain){
            self.FP.Writeln(Lns_getVM().String_format("func %s___main(args *LnsList) LnsInt {", []LnsAny{convGo_getModuleName_1198_(self.moduleTypeInfo)}))
            self.FP.Writeln(Lns_getVM().String_format("Lns_%s_init()", []LnsAny{convGo_getModuleName_1198_(self.moduleTypeInfo)}))
            self.FP.Writeln("return 0")
            self.FP.Writeln("}")
        }
    }
    self.FP.Writeln("func init() {")
    self.FP.PushIndent(nil)
    self.FP.Writeln(Lns_getVM().String_format("%s = false", []LnsAny{initModVar}))
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 1881: decl @lune.@base.@convGo.convFilter.processSubfile
func (self *convGo_convFilter) ProcessSubfile(node *Nodes_SubfileNode,_opt LnsAny) {
}

// 1887: decl @lune.@base.@convGo.convFilter.processEnv
func (self *convGo_convFilter) ProcessEnv(node *Nodes_EnvNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Writeln("Lns_LockEnvSync()")
    self.FP.PopIndent()
    self.env.FP.Push(&node.Nodes_Node)
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.env.FP.Pop()
    self.FP.PushIndent(nil)
    self.FP.Writeln("Lns_UnlockEnvSync()")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 1908: decl @lune.@base.@convGo.convFilter.processBlockSub
func (self *convGo_convFilter) ProcessBlockSub(node *Nodes_BlockNode,_opt LnsAny) {
    if _switch10833 := node.FP.Get_blockKind(); _switch10833 == Nodes_BlockKind__Block {
        self.FP.Writeln("{")
    }
    self.FP.pushProcessMode(convGo_ProcessMode__Main)
    self.FP.PushIndent(nil)
    for _, _child := range( node.FP.Get_stmtList().Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(convGo_ignoreNodeInInnerBlockSet.Has(child.FP.Get_kind())){
            convGo_filter_1171_(child, self, &node.Nodes_Node)
        }
    }
    self.FP.PopIndent()
    self.FP.popProcessMode()
    if _switch10908 := node.FP.Get_blockKind(); _switch10908 == Nodes_BlockKind__Block {
        self.FP.Writeln("}")
    }
}

// 1934: decl @lune.@base.@convGo.convFilter.expList2Slice
func (self *convGo_convFilter) expList2Slice(subList *Nodes_ExpListNode,toStem bool) {
    var processExp func(exp *Nodes_Node)
    processExp = func(exp *Nodes_Node) {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( toStem) &&
            Lns_GetEnv().SetStackVal( Ast_isClass(exp.FP.Get_expType())) ).(bool)){
            self.FP.Write(Lns_getVM().String_format("%s2Stem(", []LnsAny{self.FP.getTypeSymbolWithPrefix(exp.FP.Get_expType().FP.Get_nonnilableType())}))
            convGo_filter_1171_(Nodes_getCastUnwraped(exp), self, &subList.Nodes_Node)
            self.FP.Write(")")
        } else { 
            convGo_filter_1171_(exp, self, &subList.Nodes_Node)
        }
    }
    var mRetIndex LnsAny
    mRetIndex = Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(subList.FP.Get_mRetExp()) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_MRetExp).FP.Get_index()}))
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( mRetIndex) &&
        Lns_GetEnv().SetStackVal( mRetIndex == 1) )){
        var subExp *Nodes_Node
        subExp = subList.FP.Get_expList().GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
        if subExp.FP.Get_expType().FP.Get_kind() != Ast_TypeInfoKind__DDD{
            self.FP.Write("Lns_2DDD(")
            processExp(subExp)
            self.FP.Write(")")
        } else { 
            processExp(subExp)
        }
    } else { 
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( mRetIndex) &&
            Lns_GetEnv().SetStackVal( mRetIndex != 1) )){
            self.FP.Write("append( ")
        }
        self.FP.Write("[]LnsAny{")
        for _subIndex, _subExp := range( subList.FP.Get_expList().Items ) {
            subIndex := _subIndex + 1
            subExp := _subExp.(Nodes_NodeDownCast).ToNodes_Node()
            if mRetIndex == subIndex{
                if mRetIndex != 1{
                    self.FP.Write("}, ")
                }
                if subExp.FP.Get_expType().FP.Get_kind() != Ast_TypeInfoKind__DDD{
                    self.FP.Write("Lns_2DDD(")
                    processExp(subExp)
                    self.FP.Write(")")
                } else { 
                    processExp(subExp)
                }
                self.FP.Write("...")
                break
            }
            if subIndex != 1{
                self.FP.Write(", ")
            }
            processExp(subExp)
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( mRetIndex) &&
            Lns_GetEnv().SetStackVal( mRetIndex != 1) )){
            self.FP.Write(")")
        } else { 
            self.FP.Write("}")
        }
    }
}

// 1997: decl @lune.@base.@convGo.convFilter.processSetFromExpList
func (self *convGo_convFilter) processSetFromExpList(convArgFuncName string,dstTypeList *LnsList,expListNode *Nodes_ExpListNode) {
    if _switch11717 := convGo_getExpListKind_1321_(dstTypeList, expListNode); _switch11717 == convGo_ExpListKind__Conv {
        self.FP.Write(Lns_getVM().String_format("%s(", []LnsAny{convArgFuncName}))
        var mRetIndex LnsAny
        mRetIndex = Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expListNode.FP.Get_mRetExp()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_MRetExp).FP.Get_index()}))
        for _index, _exp := range( expListNode.FP.Get_expList().Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            if exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr{
                break
            }
            if index != 1{
                self.FP.Write(", ")
            }
            if mRetIndex == index{
                self.FP.Write("Lns_2DDD(")
                convGo_filter_1171_(Nodes_getCastUnwraped(exp), self, &expListNode.Nodes_Node)
                self.FP.Write(")")
                break
            }
            convGo_filter_1171_(exp, self, &expListNode.Nodes_Node)
            if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expListNode.FP.Get_mRetExp()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_MRetExp).FP.Get_index()})) == index{
                break
            }
        }
        self.FP.Write(")")
    } else if _switch11717 == convGo_ExpListKind__Slice {
        for _index, _argType := range( dstTypeList.Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(", ")
            }
            if expListNode.FP.Get_expList().Len() >= index{
                var argExp *Nodes_Node
                argExp = expListNode.FP.Get_expList().GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                {
                    _exp2ddd := Nodes_ExpToDDDNodeDownCastF(argExp.FP)
                    if _exp2ddd != nil {
                        exp2ddd := _exp2ddd.(*Nodes_ExpToDDDNode)
                        self.FP.expList2Slice(exp2ddd.FP.Get_expList(), false)
                    } else {
                        if argExp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr{
                            if argType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                                self.FP.Write("[]LnsAny{}")
                            } else { 
                                self.FP.Write("nil")
                            }
                        } else { 
                            convGo_filter_1171_(argExp, self, &expListNode.Nodes_Node)
                        }
                    }
                }
            } else { 
                self.FP.Write("[]LnsAny{}")
            }
        }
    } else if _switch11717 == convGo_ExpListKind__Direct {
        var mRetIndex LnsAny
        mRetIndex = Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(expListNode.FP.Get_mRetExp()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_MRetExp).FP.Get_index()}))
        for _index, _dstType := range( dstTypeList.Items ) {
            index := _index + 1
            dstType := _dstType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if mRetIndex == index - 1{
                break
            }
            if index != 1{
                self.FP.Write(", ")
            }
            var exp *Nodes_Node
            if expListNode.FP.Get_expList().Len() < index{
                exp = &self.noneNode.Nodes_Node
                
            } else { 
                exp = expListNode.FP.Get_expList().GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( index == dstTypeList.Len()) &&
                Lns_GetEnv().SetStackVal( dstType.FP.Get_kind() == Ast_TypeInfoKind__DDD) ).(bool)){
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( expListNode.FP.Get_expList().Len() < index) ||
                    Lns_GetEnv().SetStackVal( exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr) ).(bool){
                    self.FP.Write("[]LnsAny{}")
                } else { 
                    convGo_filter_1171_(exp, self, &expListNode.Nodes_Node)
                }
            } else { 
                if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( expListNode.FP.Get_expList().Len() < index) ||
                    Lns_GetEnv().SetStackVal( exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr) ).(bool){
                    self.FP.Write("nil")
                } else if expListNode.FP.Get_expTypeList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() == Ast_TypeInfoKind__DDD{
                    self.FP.Write("Lns_car(")
                    convGo_filter_1171_(exp, self, &expListNode.Nodes_Node)
                    self.FP.Write(")")
                } else { 
                    convGo_filter_1171_(exp, self, &expListNode.Nodes_Node)
                }
            }
        }
    }
}

// 2099: decl @lune.@base.@convGo.convFilter.processStmtExp
func (self *convGo_convFilter) ProcessStmtExp(node *Nodes_StmtExpNode,_opt LnsAny) {
    convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Writeln("")
}

// 2106: decl @lune.@base.@convGo.convFilter.processDeclAlge
func (self *convGo_convFilter) ProcessDeclAlge(node *Nodes_DeclAlgeNode,_opt LnsAny) {
    self.FP.Writeln(Lns_getVM().String_format("// decl alge -- %s", []LnsAny{node.FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
    self.FP.Writeln(Lns_getVM().String_format("type %s = LnsAny", []LnsAny{self.FP.getTypeSymbol(&node.FP.Get_algeType().Ast_TypeInfo)}))
    {
        __collection11932 := node.FP.Get_algeType().FP.Get_valInfoMap()
        __sorted11932 := __collection11932.CreateKeyListStr()
        __sorted11932.Sort( LnsItemKindStr, nil )
        for _, ___key11932 := range( __sorted11932.Items ) {
            valInfo := __collection11932.Items[ ___key11932 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            var algeSym string
            algeSym = self.FP.getAlgeSymbol(valInfo)
            self.FP.Writeln(Lns_getVM().String_format("type %s struct{", []LnsAny{algeSym}))
            for _index, _paramType := range( valInfo.FP.Get_typeList().Items ) {
                index := _index + 1
                paramType := _paramType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                self.FP.Writeln(Lns_getVM().String_format("Val%d %s", []LnsAny{index, self.FP.type2gotype(paramType)}))
            }
            self.FP.Writeln("}")
            if valInfo.FP.Get_typeList().Len() == 0{
                self.FP.Writeln(Lns_getVM().String_format("var %s_Obj = &%s{}", []LnsAny{algeSym, algeSym}))
            }
            self.FP.Writeln(Lns_getVM().String_format("func (self *%s) GetTxt() string {", []LnsAny{algeSym}))
            self.FP.Writeln(Lns_getVM().String_format("return \"%s.%s\"", []LnsAny{node.FP.Get_algeType().FP.Get_rawTxt(), valInfo.FP.Get_name()}))
            self.FP.Writeln("}")
        }
    }
}

// 2128: decl @lune.@base.@convGo.convFilter.processNewAlgeVal
func (self *convGo_convFilter) ProcessNewAlgeVal(node *Nodes_NewAlgeValNode,_opt LnsAny) {
    var algeSym string
    algeSym = self.FP.getAlgeSymbol(node.FP.Get_valInfo())
    if node.FP.Get_valInfo().FP.Get_typeList().Len() == 0{
        self.FP.Write(Lns_getVM().String_format("%s_Obj", []LnsAny{algeSym}))
    } else { 
        self.FP.Write(Lns_getVM().String_format("&%s{", []LnsAny{algeSym}))
        for _index, _param := range( node.FP.Get_paramList().Items ) {
            index := _index + 1
            param := _param.(Nodes_NodeDownCast).ToNodes_Node()
            if index > 1{
                self.FP.Write(", ")
            }
            convGo_filter_1171_(param, self, &node.Nodes_Node)
        }
        self.FP.Write("}")
    }
}

// 2147: decl @lune.@base.@convGo.convFilter.processDeclMember
func (self *convGo_convFilter) ProcessDeclMember(node *Nodes_DeclMemberNode,_opt LnsAny) {
    self.FP.outputSymbol(&convGo_SymbolKind__Member{Ast_isPubToExternal(node.FP.Get_accessMode())}, node.FP.Get_name().Txt)
    self.FP.Write(Lns_getVM().String_format(" %s", []LnsAny{self.FP.type2gotype(node.FP.Get_refType().FP.Get_expType())}))
    self.FP.Writeln("")
}

// 2156: decl @lune.@base.@convGo.convFilter.processExpMacroExp
func (self *convGo_convFilter) ProcessExpMacroExp(node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    for _, _stmt := range( node.FP.Get_stmtList().Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(convGo_ignoreNodeInInnerBlockSet.Has(stmt.FP.Get_kind())){
            convGo_filter_1171_(stmt, self, &node.Nodes_Node)
        }
    }
}

// 2166: decl @lune.@base.@convGo.convFilter.processDeclMacro
func (self *convGo_convFilter) ProcessDeclMacro(node *Nodes_DeclMacroNode,_opt LnsAny) {
}

// 2172: decl @lune.@base.@convGo.convFilter.processExpMacroStat
func (self *convGo_convFilter) ProcessExpMacroStat(node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpMacroStat"
    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
}

// 2178: decl @lune.@base.@convGo.convFilter.outputDeclFuncArg
func (self *convGo_convFilter) outputDeclFuncArg(funcType *Ast_TypeInfo) {
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList().Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index != 1{
            self.FP.Write(", ")
        }
        self.FP.Write(Lns_getVM().String_format("arg%d ", []LnsAny{index}))
        self.FP.Write(self.FP.type2gotype(argType))
    }
}

// 2189: decl @lune.@base.@convGo.convFilter.processDeclConstr
func (self *convGo_convFilter) ProcessDeclConstr(node *Nodes_DeclConstrNode,_opt LnsAny) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType().FP.Get_parentInfo()
    var className string
    className = self.FP.getTypeSymbol(classType)
    self.FP.Writeln(Lns_getVM().String_format("// %d: %s", []LnsAny{node.FP.Get_pos().LineNo, Nodes_getNodeKindName(node.FP.Get_kind())}))
    self.FP.Write(Lns_getVM().String_format("func (self *%s) %s(", []LnsAny{className, self.FP.getConstrSymbol(classType)}))
    for _index, _arg := range( node.FP.Get_declInfo().FP.Get_argList().Items ) {
        index := _index + 1
        arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
        if index != 1{
            self.FP.Write(",")
        }
        convGo_filter_1171_(arg, self, &node.Nodes_Node)
    }
    self.FP.Writeln(") {")
    convGo_filter_1171_(&Lns_unwrap( node.FP.Get_declInfo().FP.Get_body()).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln("}")
}

// 2209: decl @lune.@base.@convGo.convFilter.processDeclDestr
func (self *convGo_convFilter) ProcessDeclDestr(node *Nodes_DeclDestrNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processDeclDestr"
    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
}

// 2215: decl @lune.@base.@convGo.convFilter.processExpCallSuperCtor
func (self *convGo_convFilter) ProcessExpCallSuperCtor(node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("self.%s(", []LnsAny{self.FP.getConstrSymbol(node.FP.Get_superType())}))
    {
        _argList := node.FP.Get_expList()
        if _argList != nil {
            argList := _argList.(*Nodes_ExpListNode)
            self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), argList), node.FP.Get_methodType().FP.Get_argTypeInfoList(), argList)
        }
    }
    self.FP.Writeln(")")
}

// 2226: decl @lune.@base.@convGo.convFilter.processExpCallSuper
func (self *convGo_convFilter) ProcessExpCallSuper(node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("self.%s.%s(", []LnsAny{self.FP.getTypeSymbol(node.FP.Get_methodType().FP.Get_parentInfo()), self.FP.getFuncSymbol(node.FP.Get_methodType())}))
    {
        _argList := node.FP.Get_expList()
        if _argList != nil {
            argList := _argList.(*Nodes_ExpListNode)
            self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), argList), node.FP.Get_methodType().FP.Get_argTypeInfoList(), argList)
        }
    }
    self.FP.Write(")")
}

// 2238: decl @lune.@base.@convGo.convFilter.outputDummyReturn
func (self *convGo_convFilter) outputDummyReturn(retTypeInfoList *LnsList) {
    self.FP.Writeln("// insert a dummy")
    self.FP.Write("    return ")
    for _index, _retType := range( retTypeInfoList.Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(",")
        }
        if convGo_isAnyType_1174_(retType){
            self.FP.Write("nil")
        } else { 
            if _switch12668 := convGo_getOrgTypeInfo_1257_(retType); _switch12668 == Ast_builtinTypeInt {
                self.FP.Write("0")
            } else if _switch12668 == Ast_builtinTypeReal {
                self.FP.Write("0.0")
            } else if _switch12668 == Ast_builtinTypeBool {
                self.FP.Write("false")
            } else if _switch12668 == Ast_builtinTypeString {
                self.FP.Write("\"\"")
            } else {
                self.FP.Write("nil")
            }
        }
    }
    self.FP.Writeln("")
}

// 2270: decl @lune.@base.@convGo.convFilter.outputDeclFuncInfo
func (self *convGo_convFilter) OutputDeclFuncInfo(node *Nodes_Node,declInfo *Nodes_DeclFuncInfo) {
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_expType()
    if funcType.FP.Get_abstractFlag(){
        return 
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( declInfo.FP.Get_name()) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(convGo_isClosure_1177_(funcType))) )){
        self.FP.Writeln(Lns_getVM().String_format("// %d: decl %s", []LnsAny{node.FP.Get_pos().LineNo, self.FP.getCanonicalName(funcType, false)}))
    }
    var convFunc *convGo_FuncConv
    convFunc = self.FP.OutputDeclFunc(&convGo_FuncInfo__DeclInfo{node, declInfo})
    self.FP.Writeln(" {")
    self.FP.PushIndent(nil)
    if convGo_isMain_1062_(funcType){
        self.FP.Writeln(Lns_getVM().String_format("Lns_%s_init()", []LnsAny{convGo_getModuleName_1198_(self.moduleTypeInfo)}))
    }
    if declInfo.FP.Get_has__func__Symbol(){
        var nameSpace string
        nameSpace = self.FP.getCanonicalName(funcType.FP.Get_parentInfo(), false)
        var funcName string
        {
            _name := declInfo.FP.Get_name()
            if _name != nil {
                name := _name.(*Types_Token)
                funcName = name.Txt
                
            } else {
                funcName = "<anonymous>"
                
            }
        }
        var funcSym_ *Ast_SymbolInfo
        funcSym_ = Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(funcType.FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild("__func__")})/* 2299:29 */)).(*Ast_SymbolInfo)
        self.FP.Writeln(Lns_getVM().String_format("%s := \"%s.%s\"", []LnsAny{self.FP.getSymbolSym(funcSym_), nameSpace, funcName}))
    }
    for _, _convArg := range( convFunc.FP.Get_argList().Items ) {
        convArg := _convArg.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if Lns_isCondTrue( convArg.FP.Get_posForModToRef()){
            self.FP.Write(Lns_getVM().String_format("%s := _%s", []LnsAny{self.FP.getSymbolSym(convArg), self.FP.getSymbolSym(convArg)}))
            self.FP.outputAny2Type(convArg.FP.Get_typeInfo())
            self.FP.Writeln("")
        }
    }
    self.FP.PopIndent()
    {
        _body := declInfo.FP.Get_body()
        if _body != nil {
            body := _body.(*Nodes_BlockNode)
            convGo_filter_1171_(&body.Nodes_Node, self, node)
            var retTypeInfoList *LnsList
            retTypeInfoList = funcType.FP.Get_retTypeInfoList()
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( retTypeInfoList.Len() > 0) &&
                Lns_GetEnv().SetStackVal( retTypeInfoList.GetAt(retTypeInfoList.Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeNeverRet) ).(bool)){
                var needReturn bool
                needReturn = false
                {
                    var _from13029 LnsInt = body.FP.Get_stmtList().Len()
                    var _to13029 LnsInt = 1
                    _work13029 := _from13029
                    _delta13029 := -1
                    for {
                        if _delta13029 > 0 {
                           if _work13029 > _to13029 { break }
                        } else {
                           if _work13029 < _to13029 { break }
                        }
                        index := _work13029
                        var statment *Nodes_Node
                        statment = body.FP.Get_stmtList().GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                        if _switch13027 := statment.FP.Get_kind(); _switch13027 == Nodes_NodeKind_get_BlankLine() {
                        } else if _switch13027 == Nodes_NodeKind_get_Return() {
                            break
                        } else {
                            needReturn = true
                            
                            break
                        }
                        _work13029 += _delta13029
                    }
                }
                if needReturn{
                    self.FP.outputDummyReturn(funcType.FP.Get_retTypeInfoList())
                }
            }
        }
    }
    self.FP.Write("}")
    if Lns_isCondTrue( declInfo.FP.Get_name()){
        self.FP.Writeln("")
    }
}

// 2352: decl @lune.@base.@convGo.convFilter.processDeclMethod
func (self *convGo_convFilter) ProcessDeclMethod(node *Nodes_DeclMethodNode,_opt LnsAny) {
    self.FP.OutputDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo())
}

// 2358: decl @lune.@base.@convGo.convFilter.processProtoMethod
func (self *convGo_convFilter) ProcessProtoMethod(node *Nodes_ProtoMethodNode,_opt LnsAny) {
}

// 2364: decl @lune.@base.@convGo.convFilter.getEnumGetTxtSym
func (self *convGo_convFilter) getEnumGetTxtSym(enumType *Ast_EnumTypeInfo) string {
    return Lns_getVM().String_format("%s_getTxt", []LnsAny{self.FP.getTypeSymbol(&enumType.Ast_TypeInfo)})
}

// 2368: decl @lune.@base.@convGo.convFilter.processDeclEnum
func (self *convGo_convFilter) ProcessDeclEnum(node *Nodes_DeclEnumNode,_opt LnsAny) {
    if _switch13675 := self.processMode; _switch13675 == convGo_ProcessMode__DeclTopScopeVar {
        self.FP.Writeln(Lns_getVM().String_format("// decl enum -- %s ", []LnsAny{node.FP.Get_enumType().FP.GetTxt(nil, nil, nil)}))
        self.FP.Writeln(Lns_getVM().String_format("type %s = %s", []LnsAny{self.FP.getTypeSymbol(&node.FP.Get_enumType().Ast_TypeInfo), self.FP.type2gotype(&node.FP.Get_enumType().Ast_TypeInfo)}))
        {
            __collection13295 := node.FP.Get_enumType().FP.Get_name2EnumValInfo()
            __sorted13295 := __collection13295.CreateKeyListStr()
            __sorted13295.Sort( LnsItemKindStr, nil )
            for _, ___key13295 := range( __sorted13295.Items ) {
                valInfo := __collection13295.Items[ ___key13295 ].(Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
                self.FP.Write(Lns_getVM().String_format("const %s = ", []LnsAny{self.FP.getSymbol(&convGo_SymbolKind__Static{&node.FP.Get_enumType().Ast_TypeInfo}, valInfo.FP.Get_name())}))
                switch _exp13285 := valInfo.FP.Get_val().(type) {
                case *Ast_EnumLiteral__Int:
                val := _exp13285.Val1
                    self.FP.Write(Lns_getVM().String_format("%d", []LnsAny{val}))
                case *Ast_EnumLiteral__Real:
                val := _exp13285.Val1
                    self.FP.Write(Lns_getVM().String_format("%g", []LnsAny{val}))
                case *Ast_EnumLiteral__Str:
                val := _exp13285.Val1
                    self.FP.Write(convGo_str2gostr_1254_(Lns_getVM().String_format("\"%s\"", []LnsAny{val})))
                }
                self.FP.Writeln("")
            }
        }
        var listName string
        listName = Lns_getVM().String_format("%sList_", []LnsAny{self.FP.getTypeSymbol(&node.FP.Get_enumType().Ast_TypeInfo)})
        self.FP.Writeln(Lns_getVM().String_format("var %s = NewLnsList( []LnsAny {", []LnsAny{listName}))
        for _, _valName := range( node.FP.Get_valueNameList().Items ) {
            valName := _valName.(Types_TokenDownCast).ToTypes_Token()
            self.FP.Writeln(Lns_getVM().String_format("  %s,", []LnsAny{self.FP.getSymbol(&convGo_SymbolKind__Static{&node.FP.Get_enumType().Ast_TypeInfo}, valName.Txt)}))
        }
        self.FP.Writeln("})")
        var scope *Ast_Scope
        scope = Lns_unwrap( node.FP.Get_enumType().FP.Get_scope()).(*Ast_Scope)
        self.FP.OutputDeclFunc(&convGo_FuncInfo__Type{Lns_unwrap( scope.FP.GetTypeInfoChild("get__allList")).(*Ast_TypeInfo)})
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        self.FP.Writeln(Lns_getVM().String_format("return %s", []LnsAny{listName}))
        self.FP.PopIndent()
        self.FP.Writeln("}")
        var mapName string
        mapName = Lns_getVM().String_format("%sMap_", []LnsAny{self.FP.getTypeSymbol(&node.FP.Get_enumType().Ast_TypeInfo)})
        self.FP.Writeln(Lns_getVM().String_format("var %s = map[%s]string {", []LnsAny{mapName, self.FP.type2gotype(node.FP.Get_enumType().FP.Get_valTypeInfo())}))
        {
            __collection13521 := node.FP.Get_enumType().FP.Get_name2EnumValInfo()
            __sorted13521 := __collection13521.CreateKeyListStr()
            __sorted13521.Sort( LnsItemKindStr, nil )
            for _, ___key13521 := range( __sorted13521.Items ) {
                valInfo := __collection13521.Items[ ___key13521 ].(Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
                self.FP.Writeln(Lns_getVM().String_format("  %s: \"%s.%s\",", []LnsAny{self.FP.getSymbol(&convGo_SymbolKind__Static{&node.FP.Get_enumType().Ast_TypeInfo}, valInfo.FP.Get_name()), node.FP.Get_expType().FP.Get_rawTxt(), valInfo.FP.Get_name()}))
            }
        }
        self.FP.Writeln("}")
        self.FP.OutputDeclFunc(&convGo_FuncInfo__Type{Lns_unwrap( scope.FP.GetTypeInfoChild("_from")).(*Ast_TypeInfo)})
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        self.FP.Writeln(Lns_getVM().String_format("if _, ok := %s[arg1]; ok { return arg1 }", []LnsAny{mapName}))
        self.FP.Writeln("return nil")
        self.FP.PopIndent()
        self.FP.Writeln("}")
        self.FP.Writeln("")
        self.FP.Writeln(Lns_getVM().String_format("func %s(arg1 %s) string {", []LnsAny{self.FP.getEnumGetTxtSym(node.FP.Get_enumType()), self.FP.type2gotype(node.FP.Get_enumType().FP.Get_valTypeInfo())}))
        self.FP.PushIndent(nil)
        self.FP.Writeln(Lns_getVM().String_format("return %s[arg1];", []LnsAny{mapName}))
        self.FP.PopIndent()
        self.FP.Writeln("}")
    } else {
    }
}

// 2447: decl @lune.@base.@convGo.convFilter.processUnwrapSet
func (self *convGo_convFilter) ProcessUnwrapSet(node *Nodes_UnwrapSetNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processUnwrapSet"
    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
}

// 2453: decl @lune.@base.@convGo.convFilter.processIfUnwrap
func (self *convGo_convFilter) ProcessIfUnwrap(node *Nodes_IfUnwrapNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    var tempTypeList *LnsList
    tempTypeList = NewLnsList([]LnsAny{})
    for _index, _varSym := range( node.FP.Get_varSymList().Items ) {
        index := _index + 1
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(", ")
        }
        if varSym.FP.Get_name() == "_"{
            self.FP.Write("_")
        } else { 
            self.FP.Write("_" + self.FP.getSymbolSym(varSym))
        }
        tempTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeStem_))
    }
    if convGo_getExpListKind_1321_(tempTypeList, node.FP.Get_expList()) == convGo_ExpListKind__Direct{
        {
            var _from13840 LnsInt = node.FP.Get_varSymList().Len() + 1
            var _to13840 LnsInt = node.FP.Get_expList().FP.Get_expTypeList().Len()
            for _work13840 := _from13840; _work13840 <= _to13840; _work13840++ {
                self.FP.Write(", _")
            }
        }
    }
    self.FP.Write(" := ")
    self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), node.FP.Get_expList()), tempTypeList, node.FP.Get_expList())
    self.FP.Writeln("")
    self.FP.Write("if ")
    var hasSym bool
    hasSym = false
    for _, _varSym := range( node.FP.Get_varSymList().Items ) {
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if varSym.FP.Get_name() != "_"{
            if hasSym{
                self.FP.Write(" && ")
            }
            self.FP.Write(Lns_getVM().String_format("_%s != nil", []LnsAny{self.FP.getSymbolSym(varSym)}))
            hasSym = true
            
        }
    }
    self.FP.Writeln(" {")
    self.FP.PushIndent(nil)
    for _index, _varSym := range( node.FP.Get_varSymList().Items ) {
        index := _index + 1
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if varSym.FP.Get_name() != "_"{
            if varSym.FP.HasAccess(){
                self.FP.Write(Lns_getVM().String_format("%s := _%s", []LnsAny{self.FP.getSymbolSym(varSym), self.FP.getSymbolSym(varSym)}))
                if node.FP.Get_expList().FP.GetExpTypeNoDDDAt(index).FP.Get_nilable(){
                    self.FP.outputAny2Type(varSym.FP.Get_typeInfo())
                }
                self.FP.Writeln("")
            }
        }
    }
    self.FP.PopIndent()
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    {
        _nilBlock := node.FP.Get_nilBlock()
        if _nilBlock != nil {
            nilBlock := _nilBlock.(*Nodes_BlockNode)
            self.FP.Writeln("} else {")
            convGo_filter_1171_(&nilBlock.Nodes_Node, self, &node.Nodes_Node)
            self.FP.Writeln("}")
        } else {
            self.FP.Writeln("}")
        }
    }
    self.FP.PopIndent()
    self.FP.Write("}")
    self.FP.Writeln("")
}

// 2531: decl @lune.@base.@convGo.convFilter.outputLetVar
func (self *convGo_convFilter) outputLetVar(node *Nodes_DeclVarNode) {
    var declVar func()
    declVar = func() {
        if node.FP.Get_symbolInfoList().GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_scope() == self.moduleScope{
            return 
        }
        for _, _symbolInfo := range( node.FP.Get_symbolInfoList().Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if Lns_isCondTrue( symbolInfo.FP.Get_posForModToRef()){
                self.FP.Writeln(Lns_getVM().String_format("var %s %s", []LnsAny{self.FP.getSymbolSym(symbolInfo), self.FP.type2gotype(symbolInfo.FP.Get_typeInfo())}))
            }
        }
    }
    if node.FP.Get_unwrapFlag(){
        {
            _expList, _unwrapBlock := node.FP.Get_expList(), node.FP.Get_unwrapBlock()
            if _expList != nil && _unwrapBlock != nil {
                expList := _expList.(*Nodes_ExpListNode)
                unwrapBlock := _unwrapBlock.(*Nodes_BlockNode)
                if node.FP.Get_symbolInfoList().GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_scope() != self.moduleScope{
                    declVar()
                }
                self.FP.Writeln("")
                self.FP.Writeln("{")
                self.FP.PushIndent(nil)
                for _index, _varInfo := range( node.FP.Get_varList().Items ) {
                    index := _index + 1
                    varInfo := _varInfo.(Nodes_VarInfoDownCast).ToNodes_VarInfo()
                    if index != 1{
                        self.FP.Write(", ")
                    }
                    self.FP.Write(Lns_getVM().String_format("_%s", []LnsAny{varInfo.FP.Get_name().Txt}))
                }
                var tmpVarTypeList *LnsList
                tmpVarTypeList = NewLnsList([]LnsAny{})
                for _index, _ := range( node.FP.Get_symbolInfoList().Items ) {
                    index := _index + 1
                    tmpVarTypeList.Insert(Ast_TypeInfo2Stem(expList.FP.GetExpTypeNoDDDAt(index)))
                }
                if convGo_getExpListKind_1321_(tmpVarTypeList, expList) == convGo_ExpListKind__Direct{
                    {
                        var _from14337 LnsInt = tmpVarTypeList.Len() + 1
                        var _to14337 LnsInt = expList.FP.Get_expTypeList().Len()
                        for _work14337 := _from14337; _work14337 <= _to14337; _work14337++ {
                            self.FP.Write(", _")
                        }
                    }
                }
                self.FP.Write(" := ")
                self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), expList), tmpVarTypeList, expList)
                self.FP.Writeln("")
                self.FP.Write("if ")
                var hasCond bool
                hasCond = false
                for _index, _varInfo := range( node.FP.Get_varList().Items ) {
                    index := _index + 1
                    varInfo := _varInfo.(Nodes_VarInfoDownCast).ToNodes_VarInfo()
                    if expList.FP.GetExpTypeNoDDDAt(index).FP.Get_nilable(){
                        if hasCond{
                            self.FP.Write(" || ")
                        }
                        self.FP.Write(Lns_getVM().String_format("_%s == nil", []LnsAny{varInfo.FP.Get_name().Txt}))
                        hasCond = true
                        
                    }
                }
                self.FP.Writeln("{")
                convGo_filter_1171_(&unwrapBlock.Nodes_Node, self, &node.Nodes_Node)
                {
                    _thenBlock := node.FP.Get_thenBlock()
                    if _thenBlock != nil {
                        thenBlock := _thenBlock.(*Nodes_BlockNode)
                        self.FP.Writeln("} else {")
                        self.FP.PushIndent(nil)
                        for _index, _varInfo := range( node.FP.Get_symbolInfoList().Items ) {
                            index := _index + 1
                            varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            self.FP.Write(Lns_getVM().String_format("%s = _%s", []LnsAny{self.FP.getSymbolSym(varInfo), varInfo.FP.Get_name()}))
                            if expList.FP.GetExpTypeNoDDDAt(index).FP.Get_nilable(){
                                self.FP.outputAny2Type(varInfo.FP.Get_typeInfo())
                            }
                            self.FP.Writeln("")
                        }
                        self.FP.PopIndent()
                        convGo_filter_1171_(&thenBlock.Nodes_Node, self, &node.Nodes_Node)
                        self.FP.Writeln("}")
                    } else {
                        self.FP.Writeln("} else {")
                        self.FP.PushIndent(nil)
                        for _index, _varInfo := range( node.FP.Get_symbolInfoList().Items ) {
                            index := _index + 1
                            varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            self.FP.Write(Lns_getVM().String_format("%s = _%s", []LnsAny{self.FP.getSymbolSym(varInfo), varInfo.FP.Get_name()}))
                            if expList.FP.GetExpTypeNoDDDAt(index).FP.Get_nilable(){
                                self.FP.outputAny2Type(varInfo.FP.Get_typeInfo())
                            }
                            self.FP.Writeln("")
                        }
                        self.FP.PopIndent()
                        self.FP.Writeln("}")
                    }
                }
                self.FP.PopIndent()
                self.FP.Writeln("}")
            }
        }
    } else { 
        declVar()
        {
            _expList := node.FP.Get_expList()
            if _expList != nil {
                expList := _expList.(*Nodes_ExpListNode)
                var varTypeList *LnsList
                varTypeList = NewLnsList([]LnsAny{})
                for _index, _symbolInfo := range( node.FP.Get_symbolInfoList().Items ) {
                    index := _index + 1
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    varTypeList.Insert(Ast_TypeInfo2Stem(symbolInfo.FP.Get_typeInfo()))
                    if index > 1{
                        self.FP.Write(",")
                    }
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_posForModToRef()) ||
                        Lns_GetEnv().SetStackVal( Ast_isPubToExternal(symbolInfo.FP.Get_accessMode())) )){
                        self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{self.FP.getSymbolSym(symbolInfo)}))
                    } else { 
                        self.FP.Write("_")
                    }
                }
                if convGo_getExpListKind_1321_(varTypeList, expList) == convGo_ExpListKind__Direct{
                    {
                        var _from14793 LnsInt = varTypeList.Len() + 1
                        var _to14793 LnsInt = expList.FP.Get_expTypeList().Len()
                        for _work14793 := _from14793; _work14793 <= _to14793; _work14793++ {
                            self.FP.Write(", _")
                        }
                    }
                }
                self.FP.Write(" = ")
                self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), expList), varTypeList, expList)
                self.FP.Writeln("")
            }
        }
    }
}

// 2660: decl @lune.@base.@convGo.convFilter.processDeclVar
func (self *convGo_convFilter) ProcessDeclVar(node *Nodes_DeclVarNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processDeclVar"
    if _switch15288 := node.FP.Get_mode(); _switch15288 == Nodes_DeclVarMode__Let {
        self.FP.outputLetVar(node)
    } else if _switch15288 == Nodes_DeclVarMode__Unwrap {
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        for _, _varSym := range( node.FP.Get_symbolInfoList().Items ) {
            varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            self.FP.Writeln(Lns_getVM().String_format("var _%s LnsAny", []LnsAny{varSym.FP.Get_name()}))
        }
        var expList *Nodes_ExpListNode
        
        {
            _expList := node.FP.Get_expList()
            if _expList == nil{
                Util_err("illegal")
            } else {
                expList = _expList.(*Nodes_ExpListNode)
            }
        }
        var setVals func()
        setVals = func() {
            for _index, _varSym := range( node.FP.Get_symbolInfoList().Items ) {
                index := _index + 1
                varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.Write(Lns_getVM().String_format("%s = _%s", []LnsAny{self.FP.getSymbolSym(varSym), varSym.FP.Get_name()}))
                if expList.FP.GetExpTypeNoDDDAt(index).FP.Get_nilable(){
                    self.FP.outputAny2Type(varSym.FP.Get_typeInfo())
                }
                self.FP.Writeln("")
            }
        }
        var typeList *LnsList
        typeList = NewLnsList([]LnsAny{})
        for _index, _varSym := range( node.FP.Get_symbolInfoList().Items ) {
            index := _index + 1
            varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            typeList.Insert(Ast_TypeInfo2Stem(varSym.FP.Get_typeInfo()))
            if index > 1{
                self.FP.Write(",")
            }
            self.FP.Write(Lns_getVM().String_format("_%s", []LnsAny{varSym.FP.Get_name()}))
        }
        if convGo_getExpListKind_1321_(typeList, expList) == convGo_ExpListKind__Direct{
            {
                var _from15067 LnsInt = node.FP.Get_symbolInfoList().Len() + 1
                var _to15067 LnsInt = expList.FP.Get_expTypeList().Len()
                for _work15067 := _from15067; _work15067 <= _to15067; _work15067++ {
                    self.FP.Write(",_")
                }
            }
        }
        self.FP.Write(" = ")
        self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), expList), typeList, expList)
        self.FP.Writeln("")
        self.FP.Write("if ")
        var hasCond bool
        hasCond = false
        for _index, _varSym := range( node.FP.Get_symbolInfoList().Items ) {
            index := _index + 1
            varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if expList.FP.GetExpTypeNoDDDAt(index).FP.Get_nilable(){
                if hasCond{
                    self.FP.Write(" || ")
                }
                self.FP.Write(Lns_getVM().String_format("_%s == nil", []LnsAny{varSym.FP.Get_name()}))
                hasCond = true
                
            }
        }
        self.FP.Writeln(" {")
        convGo_filter_1171_(&Lns_unwrap( node.FP.Get_unwrapBlock()).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
        {
            _thenBlock := node.FP.Get_thenBlock()
            if _thenBlock != nil {
                thenBlock := _thenBlock.(*Nodes_BlockNode)
                self.FP.Writeln("} else {")
                self.FP.PushIndent(nil)
                setVals()
                self.FP.PopIndent()
                convGo_filter_1171_(&thenBlock.Nodes_Node, self, &node.Nodes_Node)
                self.FP.Writeln("}")
            } else {
                self.FP.Writeln("}")
                setVals()
            }
        }
        self.FP.PopIndent()
        self.FP.Writeln("}")
    } else {
        Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
    }
}

// 2740: decl @lune.@base.@convGo.convFilter.processWhen
func (self *convGo_convFilter) ProcessWhen(node *Nodes_WhenNode,_opt LnsAny) {
    self.FP.Write("if ")
    for _index, _symPair := range( node.FP.Get_symPairList().Items ) {
        index := _index + 1
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        if index > 1{
            self.FP.Write(" && ")
        }
        self.FP.Write(Lns_getVM().String_format("%s != nil", []LnsAny{self.FP.getSymbolSym(symPair.FP.Get_src())}))
        symPair.FP.Get_dst().FP.Set_convModuleParam(true)
    }
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    for _, _symPair := range( node.FP.Get_symPairList().Items ) {
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        if Lns_isCondTrue( symPair.FP.Get_dst().FP.Get_posForModToRef()){
            self.FP.Write(Lns_getVM().String_format("%s_%d := %s", []LnsAny{symPair.FP.Get_dst().FP.Get_name(), symPair.FP.Get_dst().FP.Get_symbolId(), self.FP.getSymbolSym(symPair.FP.Get_src())}))
            self.FP.outputAny2Type(symPair.FP.Get_dst().FP.Get_typeInfo())
            self.FP.Writeln("")
        }
    }
    self.FP.PopIndent()
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    {
        _elseBlock := node.FP.Get_elseBlock()
        if _elseBlock != nil {
            elseBlock := _elseBlock.(*Nodes_BlockNode)
            self.FP.Writeln("} else {")
            convGo_filter_1171_(&elseBlock.Nodes_Node, self, &node.Nodes_Node)
            self.FP.Write("}")
        } else {
            self.FP.Write("}")
        }
    }
    self.FP.Writeln("")
}

// 2777: decl @lune.@base.@convGo.convFilter.processDeclArg
func (self *convGo_convFilter) ProcessDeclArg(node *Nodes_DeclArgNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("%s ", []LnsAny{self.FP.getSymbolSym(node.FP.Get_symbolInfo())}))
    convGo_filter_1171_(&Lns_unwrap( node.FP.Get_argType()).(*Nodes_RefTypeNode).Nodes_Node, self, &node.Nodes_Node)
}

// 2784: decl @lune.@base.@convGo.convFilter.processDeclArgDDD
func (self *convGo_convFilter) ProcessDeclArgDDD(node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    self.FP.Write("ddd []LnsAny")
}

// 2790: decl @lune.@base.@convGo.convFilter.processExpSubDDD
func (self *convGo_convFilter) ProcessExpSubDDD(node *Nodes_ExpSubDDDNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpSubDDD"
    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
}

// 2797: decl @lune.@base.@convGo.convFilter.processDeclForm
func (self *convGo_convFilter) ProcessDeclForm(node *Nodes_DeclFormNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("type %s ", []LnsAny{self.FP.getFuncSymbol(node.FP.Get_expType())}))
    self.FP.OutputDeclFunc(&convGo_FuncInfo__Type{node.FP.Get_expType()})
    self.FP.Writeln("")
}

// 2806: decl @lune.@base.@convGo.convFilter.processDeclFunc
func (self *convGo_convFilter) ProcessDeclFunc(node *Nodes_DeclFuncNode,_opt LnsAny) {
    {
        _funcSym := node.FP.Get_declInfo().FP.Get_symbol()
        if _funcSym != nil {
            funcSym := _funcSym.(*Ast_SymbolInfo)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(funcSym.FP.Get_posForModToRef())) &&
                Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(funcSym.FP.Get_accessMode()))) ).(bool)){
                return 
            }
        }
    }
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_expType()
    if (self.processMode == convGo_ProcessMode__NonClosureFuncDecl) == convGo_isClosure_1177_(node.FP.Get_expType()){
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( self.processMode != convGo_ProcessMode__NonClosureFuncDecl) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(node.FP.Get_declInfo().FP.Get_symbol())) ).(bool)){
            self.FP.Write(self.FP.getFuncSymbol(funcType))
        }
        return 
    }
    if convGo_isClosure_1177_(funcType){
        {
            _funcSym := node.FP.Get_declInfo().FP.Get_symbol()
            if _funcSym != nil {
                funcSym := _funcSym.(*Ast_SymbolInfo)
                self.FP.Write("var ")
                self.FP.outputSymbol(&convGo_SymbolKind__Func{funcType}, funcSym.FP.Get_name())
                self.FP.Write(" ")
                self.FP.OutputDeclFunc(&convGo_FuncInfo__DeclInfo{&node.Nodes_Node, node.FP.Get_declInfo()})
                self.FP.Writeln("")
                self.FP.outputSymbol(&convGo_SymbolKind__Func{funcType}, funcSym.FP.Get_name())
                self.FP.Write(" = ")
            }
        }
    }
    self.FP.OutputDeclFuncInfo(&node.Nodes_Node, node.FP.Get_declInfo())
}

// 2838: decl @lune.@base.@convGo.convFilter.processRefType
func (self *convGo_convFilter) ProcessRefType(node *Nodes_RefTypeNode,_opt LnsAny) {
    self.FP.Write(self.FP.type2gotype(node.FP.Get_expType()))
}

// 2860: decl @lune.@base.@convGo.convFilter.processCond
func (self *convGo_convFilter) processCond(node *Nodes_Node,parent *Nodes_Node) {
    if node.FP.Get_expType() != Ast_builtinTypeBool{
        self.FP.Write("Lns_isCondTrue( ")
        convGo_filter_1171_(node, self, parent)
        self.FP.Write(")")
    } else { 
        convGo_filter_1171_(node, self, parent)
    }
}

// 2871: decl @lune.@base.@convGo.convFilter.processIf
func (self *convGo_convFilter) ProcessIf(node *Nodes_IfNode,_opt LnsAny) {
    for _, _stmt := range( node.FP.Get_stmtList().Items ) {
        stmt := _stmt.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if _switch16071 := stmt.FP.Get_kind(); _switch16071 == Nodes_IfKind__If {
            self.FP.Write("if ")
            self.FP.processCond(stmt.FP.Get_exp(), &node.Nodes_Node)
            self.FP.Writeln("{")
            convGo_filter_1171_(&stmt.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
        } else if _switch16071 == Nodes_IfKind__ElseIf {
            self.FP.Write("} else if ")
            self.FP.processCond(stmt.FP.Get_exp(), &node.Nodes_Node)
            self.FP.Writeln("{")
            convGo_filter_1171_(&stmt.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
        } else if _switch16071 == Nodes_IfKind__Else {
            self.FP.Writeln("} else { ")
            convGo_filter_1171_(&stmt.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln("}")
}

// 2897: decl @lune.@base.@convGo.convFilter.processSwitch
func (self *convGo_convFilter) ProcessSwitch(node *Nodes_SwitchNode,_opt LnsAny) {
    var valName string
    valName = Lns_getVM().String_format("_switch%d", []LnsAny{node.FP.Get_id()})
    self.FP.Write(Lns_getVM().String_format("if %s := ", []LnsAny{valName}))
    convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Write("; ")
    for _caseIndex, _caseNode := range( node.FP.Get_caseList().Items ) {
        caseIndex := _caseIndex + 1
        caseNode := _caseNode.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        if caseIndex != 1{
            self.FP.Write("} else if ")
        }
        for _index, _exp := range( caseNode.FP.Get_expList().FP.Get_expList().Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            if index != 1{
                self.FP.Write(" || ")
            }
            self.FP.Write(Lns_getVM().String_format("%s == ", []LnsAny{valName}))
            convGo_filter_1171_(exp, self, &caseNode.FP.Get_expList().Nodes_Node)
        }
        self.FP.Writeln(" {")
        convGo_filter_1171_(&caseNode.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    }
    {
        _defaultBlock := node.FP.Get_default()
        if _defaultBlock != nil {
            defaultBlock := _defaultBlock.(*Nodes_BlockNode)
            self.FP.Writeln("} else {")
            convGo_filter_1171_(&defaultBlock.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln("}")
}

// 2930: decl @lune.@base.@convGo.convFilter.processMatch
func (self *convGo_convFilter) ProcessMatch(node *Nodes_MatchNode,_opt LnsAny) {
    var hasAccessing func() bool
    hasAccessing = func() bool {
        for _, _caseInfo := range( node.FP.Get_caseList().Items ) {
            caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
            for _, _symbol := range( caseInfo.FP.Get_valParamNameList().Items ) {
                symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if Lns_isCondTrue( symbol.FP.Get_posForModToRef()){
                    return true
                }
            }
        }
        return false
    }
    var val string
    if hasAccessing(){
        val = Lns_getVM().String_format("_exp%d", []LnsAny{node.FP.Get_id()})
        
        self.FP.Write(Lns_getVM().String_format("switch %s := ", []LnsAny{val}))
    } else { 
        val = ""
        
        self.FP.Write("switch ")
    }
    convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
    self.FP.Writeln(".(type) {")
    for _, _caseInfo := range( node.FP.Get_caseList().Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        self.FP.Writeln(Lns_getVM().String_format("case *%s:", []LnsAny{self.FP.getAlgeSymbol(caseInfo.FP.Get_valInfo())}))
        for _index, _symbol := range( caseInfo.FP.Get_valParamNameList().Items ) {
            index := _index + 1
            symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if Lns_isCondTrue( symbol.FP.Get_posForModToRef()){
                self.FP.Writeln(Lns_getVM().String_format("%s := %s.Val%d", []LnsAny{self.FP.getSymbolSym(symbol), val, index}))
            }
        }
        convGo_filter_1171_(&caseInfo.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    }
    {
        _defaultBlock := node.FP.Get_defaultBlock()
        if _defaultBlock != nil {
            defaultBlock := _defaultBlock.(*Nodes_Node)
            self.FP.Writeln("default:")
            convGo_filter_1171_(defaultBlock, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln("}")
}

// 2970: decl @lune.@base.@convGo.convFilter.processWhile
func (self *convGo_convFilter) ProcessWhile(node *Nodes_WhileNode,_opt LnsAny) {
    self.FP.Write("for ")
    if Lns_op_not(node.FP.Get_infinit()){
        self.FP.processCond(node.FP.Get_exp(), &node.Nodes_Node)
    }
    self.FP.Writeln(" {")
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln("}")
}

// 2984: decl @lune.@base.@convGo.convFilter.processRepeat
func (self *convGo_convFilter) ProcessRepeat(node *Nodes_RepeatNode,_opt LnsAny) {
    self.FP.Writeln("for {")
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.PushIndent(nil)
    self.FP.Write("if ")
    self.FP.processCond(node.FP.Get_exp(), &node.Nodes_Node)
    self.FP.Writeln("{ break }")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 3002: decl @lune.@base.@convGo.convFilter.processFor
func (self *convGo_convFilter) ProcessFor(node *Nodes_ForNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    var fromSym string
    fromSym = Lns_getVM().String_format("_from%d", []LnsAny{node.FP.Get_id()})
    var toSym string
    toSym = Lns_getVM().String_format("_to%d", []LnsAny{node.FP.Get_id()})
    var deltaSym string
    deltaSym = Lns_getVM().String_format("_delta%d", []LnsAny{node.FP.Get_id()})
    var workSym string
    workSym = Lns_getVM().String_format("_work%d", []LnsAny{node.FP.Get_id()})
    var valType string
    valType = Lns_getVM().String_format("%s", []LnsAny{self.FP.type2gotype(node.FP.Get_init().FP.Get_expType())})
    self.FP.Write(Lns_getVM().String_format("var %s %s = ", []LnsAny{fromSym, valType}))
    convGo_filter_1171_(node.FP.Get_init(), self, &node.Nodes_Node)
    self.FP.Writeln("")
    self.FP.Write(Lns_getVM().String_format("var %s %s = ", []LnsAny{toSym, valType}))
    convGo_filter_1171_(node.FP.Get_to(), self, &node.Nodes_Node)
    self.FP.Writeln("")
    {
        _delta := node.FP.Get_delta()
        if _delta != nil {
            delta := _delta.(*Nodes_Node)
            self.FP.Writeln(Lns_getVM().String_format("%s := %s", []LnsAny{workSym, fromSym}))
            self.FP.Write(Lns_getVM().String_format("%s := ", []LnsAny{deltaSym}))
            convGo_filter_1171_(delta, self, &node.Nodes_Node)
            self.FP.Writeln("")
            self.FP.Writeln("for {")
            self.FP.PushIndent(nil)
            self.FP.Writeln(Lns_getVM().String_format("if %s > 0 {", []LnsAny{deltaSym}))
            self.FP.Writeln(Lns_getVM().String_format("   if %s > %s { break }", []LnsAny{workSym, toSym}))
            self.FP.Writeln("} else {")
            self.FP.Writeln(Lns_getVM().String_format("   if %s < %s { break }", []LnsAny{workSym, toSym}))
            self.FP.Writeln("}")
            self.FP.PopIndent()
        } else {
            self.FP.Writeln(Lns_getVM().String_format("for %s := %s; %s <= %s; %s++ {", []LnsAny{workSym, fromSym, workSym, toSym, workSym}))
        }
    }
    self.FP.PushIndent(nil)
    if Lns_isCondTrue( node.FP.Get_val().FP.Get_posForModToRef()){
        self.FP.Writeln(Lns_getVM().String_format("%s := %s", []LnsAny{node.FP.Get_val().FP.Get_name(), workSym}))
    }
    self.FP.PopIndent()
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    if Lns_isCondTrue( node.FP.Get_delta()){
        self.FP.PushIndent(nil)
        self.FP.Writeln(Lns_getVM().String_format("%s += %s", []LnsAny{workSym, deltaSym}))
        self.FP.PopIndent()
    }
    self.FP.Writeln("}")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 3066: decl @lune.@base.@convGo.convFilter.processApply
func (self *convGo_convFilter) ProcessApply(node *Nodes_ApplyNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    var formSym string
    formSym = Lns_getVM().String_format("_form%d", []LnsAny{node.FP.Get_id()})
    var paramSym string
    paramSym = Lns_getVM().String_format("_param%d", []LnsAny{node.FP.Get_id()})
    var prevSym string
    prevSym = Lns_getVM().String_format("_prev%d", []LnsAny{node.FP.Get_id()})
    if Lns_isCondTrue( node.FP.Get_expList().FP.Get_mRetExp()){
        self.FP.Write(Lns_getVM().String_format("%s, %s, %s := ", []LnsAny{formSym, paramSym, prevSym}))
        convGo_filter_1171_(&node.FP.Get_expList().Nodes_Node, self, &node.Nodes_Node)
        self.FP.Writeln("")
    } else { 
        self.FP.Write(Lns_getVM().String_format("%s, %s, %s := ", []LnsAny{formSym, paramSym, prevSym}))
        convGo_filter_1171_(node.FP.Get_expList().FP.Get_expList().GetAt(1).(Nodes_NodeDownCast).ToNodes_Node(), self, &node.FP.Get_expList().Nodes_Node)
        self.FP.Write(",")
        convGo_filter_1171_(node.FP.Get_expList().FP.Get_expList().GetAt(2).(Nodes_NodeDownCast).ToNodes_Node(), self, &node.FP.Get_expList().Nodes_Node)
        self.FP.Write(", LnsAny(")
        convGo_filter_1171_(node.FP.Get_expList().FP.Get_expList().GetAt(3).(Nodes_NodeDownCast).ToNodes_Node(), self, &node.FP.Get_expList().Nodes_Node)
        self.FP.Writeln(")")
    }
    self.FP.Writeln("for {")
    self.FP.PushIndent(nil)
    var setTxt string
    setTxt = prevSym
    {
        var _from17393 LnsInt = 2
        var _to17393 LnsInt = node.FP.Get_varList().Len()
        for _work17393 := _from17393; _work17393 <= _to17393; _work17393++ {
            index := _work17393
            var symInfo *Ast_SymbolInfo
            symInfo = node.FP.Get_varList().GetAt(index).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            self.FP.Writeln(Lns_getVM().String_format("var %s %s", []LnsAny{self.FP.getSymbolSym(symInfo), self.FP.type2gotype(symInfo.FP.Get_typeInfo())}))
            setTxt = Lns_getVM().String_format("%s, %s", []LnsAny{setTxt, self.FP.getSymbolSym(symInfo)})
            
        }
    }
    if node.FP.Get_expList().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Ext{
        var workSym string
        workSym = Lns_getVM().String_format("_work%d", []LnsAny{node.FP.Get_id()})
        self.FP.Writeln(Lns_getVM().String_format("%s := %s.(*Lns_luaValue).Call( Lns_2DDD( %s, %s ) )", []LnsAny{workSym, formSym, paramSym, prevSym}))
        self.FP.Write(Lns_getVM().String_format("%s = ", []LnsAny{setTxt}))
        for _index, _ := range( node.FP.Get_varList().Items ) {
            index := _index + 1
            if index > 1{
                self.FP.Write(",")
            }
            self.FP.Write(Lns_getVM().String_format("Lns_getFromMulti(%s,%d)", []LnsAny{workSym, index - 1}))
        }
        self.FP.Writeln("")
    } else { 
        self.FP.Writeln(Lns_getVM().String_format("%s = %s( %s, %s )", []LnsAny{setTxt, formSym, paramSym, prevSym}))
    }
    self.FP.Writeln(Lns_getVM().String_format("if Lns_IsNil( %s ) { break }", []LnsAny{prevSym}))
    var topSymInfo *Ast_SymbolInfo
    topSymInfo = node.FP.Get_varList().GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
    if topSymInfo.FP.Get_name() != "_"{
        self.FP.Writeln(Lns_getVM().String_format("%s := %s.(%s)", []LnsAny{self.FP.getSymbolSym(topSymInfo), prevSym, self.FP.type2gotype(topSymInfo.FP.Get_typeInfo())}))
    }
    self.FP.PopIndent()
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln("}")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 3137: decl @lune.@base.@convGo.convFilter.outputForeachLua
func (self *convGo_convFilter) outputForeachLua(node *Nodes_Node,sortFlag bool,exp *Nodes_Node,val *Ast_SymbolInfo,key LnsAny,block *Nodes_BlockNode) {
    __func__ := "@lune.@base.@convGo.convFilter.outputForeachLua"
    if _switch18164 := exp.FP.Get_expType().FP.Get_extedType().FP.Get_kind(); _switch18164 == Ast_TypeInfoKind__List || _switch18164 == Ast_TypeInfoKind__Map {
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        var tmpExp string
        tmpExp = Lns_getVM().String_format("_exp%d", []LnsAny{node.FP.Get_id()})
        var tmpKey string
        tmpKey = Lns_getVM().String_format("_key%d", []LnsAny{node.FP.Get_id()})
        var tmpVal string
        tmpVal = Lns_getVM().String_format("_val%d", []LnsAny{node.FP.Get_id()})
        var tmpIndex string
        tmpIndex = Lns_getVM().String_format("_index%d", []LnsAny{node.FP.Get_id()})
        var sorted string
        sorted = Lns_getVM().String_format("_sorted%d", []LnsAny{node.FP.Get_id()})
        self.FP.Write(Lns_getVM().String_format("%s := ", []LnsAny{tmpExp}))
        convGo_filter_1171_(exp, self, node)
        self.FP.Writeln("")
        var tmpValName string
        if val.FP.HasAccess(){
            tmpValName = tmpVal
            
        } else { 
            tmpValName = "_"
            
        }
        if sortFlag{
            self.FP.Writeln(Lns_getVM().String_format("%s := %s.SortMapKeyList( %s )", []LnsAny{sorted, self.env.FP.GetLuavm(node.FP.IsThreading()), tmpExp}))
            self.FP.Writeln(Lns_getVM().String_format("%s, %s := %s.Get1stFromMap()", []LnsAny{tmpIndex, tmpKey, sorted}))
            self.FP.Writeln(Lns_getVM().String_format("for %s != nil {", []LnsAny{tmpIndex}))
            self.FP.PushIndent(nil)
        } else { 
            self.FP.Writeln(Lns_getVM().String_format("%s, %s := %s.Get1stFromMap()", []LnsAny{tmpKey, tmpValName, tmpExp}))
            self.FP.Writeln(Lns_getVM().String_format("for %s != nil {", []LnsAny{tmpKey}))
            self.FP.PushIndent(nil)
        }
        {
            _keySym := key
            if _keySym != nil {
                keySym := _keySym.(*Ast_SymbolInfo)
                if keySym.FP.HasAccess(){
                    self.FP.Write(Lns_getVM().String_format("%s := %s", []LnsAny{self.FP.getSymbolSym(keySym), tmpKey}))
                    self.FP.outputAny2Type(keySym.FP.Get_typeInfo())
                    self.FP.Writeln("")
                }
            }
        }
        if val.FP.HasAccess(){
            if sortFlag{
                self.FP.Write(Lns_getVM().String_format("%s := %s.GetAt( %s )", []LnsAny{self.FP.getSymbolSym(val), tmpExp, tmpKey}))
            } else { 
                self.FP.Write(Lns_getVM().String_format("%s := %s", []LnsAny{self.FP.getSymbolSym(val), tmpVal}))
            }
            self.FP.outputAny2Type(val.FP.Get_typeInfo())
            self.FP.Writeln("")
        }
        self.FP.PopIndent()
        convGo_filter_1171_(&block.Nodes_Node, self, node)
        self.FP.PushIndent(nil)
        if Lns_op_not(sortFlag){
            self.FP.Write(Lns_getVM().String_format("%s, %s = ", []LnsAny{tmpKey, tmpValName}))
            self.FP.Writeln(Lns_getVM().String_format("%s.NextFromMap( %s )", []LnsAny{tmpExp, tmpKey}))
        } else { 
            self.FP.Write(Lns_getVM().String_format("%s, %s = ", []LnsAny{tmpIndex, tmpKey}))
            self.FP.Writeln(Lns_getVM().String_format("%s.NextFromMap( %s )", []LnsAny{sorted, tmpIndex}))
        }
        self.FP.PopIndent()
        self.FP.Writeln("}")
        self.FP.PopIndent()
        self.FP.Writeln("}")
    } else {
        Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
    }
}

// 3217: decl @lune.@base.@convGo.convFilter.processForeach
func (self *convGo_convFilter) ProcessForeach(node *Nodes_ForeachNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processForeach"
    if node.FP.Get_exp().FP.Get_expType().FP.Get_srcTypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Ext{
        self.FP.outputForeachLua(&node.Nodes_Node, false, node.FP.Get_exp(), node.FP.Get_val(), node.FP.Get_key(), node.FP.Get_block())
        return 
    }
    var hasAccessLoopSym LnsAny
    hasAccessLoopSym = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(node.FP.Get_key()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_SymbolInfo).FP.Get_posForModToRef()}))) ||
        Lns_GetEnv().SetStackVal( node.FP.Get_val().FP.Get_posForModToRef()) )
    self.FP.Write("for ")
    var loopExpType *Ast_TypeInfo
    loopExpType = node.FP.Get_exp().FP.Get_expType()
    if _switch19057 := loopExpType.FP.Get_kind(); _switch19057 == Ast_TypeInfoKind__List || _switch19057 == Ast_TypeInfoKind__Array {
        var valName string
        valName = self.FP.getSymbolSym(node.FP.Get_val())
        var itemType *Ast_TypeInfo
        itemType = loopExpType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_isCondTrue( hasAccessLoopSym){
            {
                _key := node.FP.Get_key()
                if _key != nil {
                    key := _key.(*Ast_SymbolInfo)
                    if key.FP.Get_name() != "_"{
                        self.FP.Write(Lns_getVM().String_format("_%s", []LnsAny{key.FP.Get_name()}))
                    } else { 
                        self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{key.FP.Get_name()}))
                    }
                    
                    self.FP.Write(", ")
                } else {
                    self.FP.Write("_, ")
                }
            }
            if valName != "_"{
                self.FP.Write(Lns_getVM().String_format("_%s", []LnsAny{valName}))
            } else { 
                self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{valName}))
            }
            
            self.FP.Write(" := ")
        }
        self.FP.Write("range( ")
        convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
        self.FP.Writeln(".Items ) {")
        self.FP.PushIndent(nil)
        {
            _key := node.FP.Get_key()
            if _key != nil {
                key := _key.(*Ast_SymbolInfo)
                if Lns_isCondTrue( key.FP.Get_posForModToRef()){
                    self.FP.Writeln(Lns_getVM().String_format("%s := _%s + 1", []LnsAny{self.FP.getSymbolSym(key), key.FP.Get_name()}))
                }
            }
        }
        if valName != "_"{
            self.FP.Write(Lns_getVM().String_format("%s := _%s", []LnsAny{valName, valName}))
            self.FP.outputStem2Type(itemType)
            self.FP.Writeln("")
        }
        
        self.FP.PopIndent()
    } else if _switch19057 == Ast_TypeInfoKind__Map {
        var valName string
        valName = self.FP.getSymbolSym(node.FP.Get_val())
        var itemType *Ast_TypeInfo
        itemType = loopExpType.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var keyType *Ast_TypeInfo
        keyType = loopExpType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_isCondTrue( hasAccessLoopSym){
            {
                _key := node.FP.Get_key()
                if _key != nil {
                    key := _key.(*Ast_SymbolInfo)
                    if key.FP.Get_name() != "_"{
                        self.FP.Write(Lns_getVM().String_format("_%s", []LnsAny{key.FP.Get_name()}))
                    } else { 
                        self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{key.FP.Get_name()}))
                    }
                    
                    self.FP.Write(", ")
                } else {
                    self.FP.Write("_, ")
                }
            }
            if valName != "_"{
                self.FP.Write(Lns_getVM().String_format("_%s", []LnsAny{valName}))
            } else { 
                self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{valName}))
            }
            
            self.FP.Write(" := ")
        }
        self.FP.Write("range( ")
        convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
        self.FP.Writeln(".Items ) {")
        self.FP.PushIndent(nil)
        {
            _key := node.FP.Get_key()
            if _key != nil {
                key := _key.(*Ast_SymbolInfo)
                if key.FP.Get_name() != "_"{
                    self.FP.Write(Lns_getVM().String_format("%s := _%s", []LnsAny{key.FP.Get_name(), key.FP.Get_name()}))
                    self.FP.outputStem2Type(keyType)
                    self.FP.Writeln("")
                }
                
            }
        }
        if valName != "_"{
            self.FP.Write(Lns_getVM().String_format("%s := _%s", []LnsAny{valName, valName}))
            self.FP.outputStem2Type(itemType)
            self.FP.Writeln("")
        }
        
        self.FP.PopIndent()
    } else if _switch19057 == Ast_TypeInfoKind__Set {
        var valType *Ast_TypeInfo
        valType = loopExpType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var valName string
        valName = self.FP.getSymbolSym(node.FP.Get_val())
        if Lns_isCondTrue( hasAccessLoopSym){
            if valName != "_"{
                self.FP.Write(Lns_getVM().String_format("_%s", []LnsAny{valName}))
            } else { 
                self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{valName}))
            }
            
            self.FP.Write(" := ")
        }
        self.FP.Write("range( ")
        convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
        self.FP.Writeln(".Items ) {")
        self.FP.PushIndent(nil)
        if valName != "_"{
            self.FP.Write(Lns_getVM().String_format("%s := _%s", []LnsAny{valName, valName}))
            self.FP.outputStem2Type(valType)
            self.FP.Writeln("")
        }
        
        self.FP.PopIndent()
    } else {
        Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
    }
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.Write("}")
    self.FP.Writeln("")
}

// 3335: decl @lune.@base.@convGo.convFilter.processForsort
func (self *convGo_convFilter) ProcessForsort(node *Nodes_ForsortNode,_opt LnsAny) {
    if node.FP.Get_exp().FP.Get_expType().FP.Get_srcTypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Ext{
        self.FP.outputForeachLua(&node.Nodes_Node, true, node.FP.Get_exp(), node.FP.Get_val(), node.FP.Get_key(), node.FP.Get_block())
        return 
    }
    var keySym LnsAny
    var valSym LnsAny
    var keyTypeInfo *Ast_TypeInfo
    keyTypeInfo = node.FP.Get_exp().FP.Get_expType().FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
    if node.FP.Get_exp().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Set{
        keySym = node.FP.Get_val()
        
        valSym = node.FP.Get_key()
        
    } else { 
        keySym = node.FP.Get_key()
        
        valSym = node.FP.Get_val()
        
    }
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    var collSym string
    collSym = Lns_getVM().String_format("__collection%d", []LnsAny{node.FP.Get_id()})
    self.FP.Write(Lns_getVM().String_format("%s := ", []LnsAny{collSym}))
    convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Writeln("")
    var sortSym string
    sortSym = Lns_getVM().String_format("__sorted%d", []LnsAny{node.FP.Get_id()})
    self.FP.Write(Lns_getVM().String_format("%s := %s.", []LnsAny{sortSym, collSym}))
    if _switch19387 := keyTypeInfo; _switch19387 == Ast_builtinTypeInt {
        self.FP.Writeln("CreateKeyListInt()")
    } else if _switch19387 == Ast_builtinTypeReal {
        self.FP.Writeln("CreateKeyListReal()")
    } else if _switch19387 == Ast_builtinTypeString {
        self.FP.Writeln("CreateKeyListStr()")
    } else {
        self.FP.Writeln("CreateKeyListStem()")
    }
    self.FP.Writeln(Lns_getVM().String_format("%s.Sort( %s, nil )", []LnsAny{sortSym, convGo_getLnsItemKind_1815_(keyTypeInfo)}))
    self.FP.Write("for _, ")
    var key string
    key = Lns_getVM().String_format("__key%d", []LnsAny{node.FP.Get_id()})
    if keySym != nil{
        keySym_7403 := keySym.(*Ast_SymbolInfo)
        key = Lns_getVM().String_format("%s", []LnsAny{self.FP.getSymbolSym(keySym_7403)})
        
    }
    self.FP.Write(Lns_getVM().String_format("_%s", []LnsAny{key}))
    self.FP.Writeln(Lns_getVM().String_format(" := range( %s.Items ) {", []LnsAny{sortSym}))
    self.FP.PushIndent(nil)
    if valSym != nil{
        valSym_7405 := valSym.(*Ast_SymbolInfo)
        if Lns_isCondTrue( valSym_7405.FP.Get_posForModToRef()){
            self.FP.Write(Lns_getVM().String_format("%s := %s.Items[ _%s ]", []LnsAny{self.FP.getSymbolSym(valSym_7405), collSym, key}))
            self.FP.outputStem2Type(valSym_7405.FP.Get_typeInfo())
            self.FP.Writeln("")
        }
    }
    if keySym != nil{
        keySym_7408 := keySym.(*Ast_SymbolInfo)
        if Lns_isCondTrue( keySym_7408.FP.Get_posForModToRef()){
            self.FP.Write(Lns_getVM().String_format("%s := _%s", []LnsAny{key, key}))
            self.FP.outputStem2Type(keySym_7408.FP.Get_typeInfo())
            self.FP.Writeln("")
        }
    }
    self.FP.PopIndent()
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln("}")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 3413: decl @lune.@base.@convGo.convFilter.processExpUnwrap
func (self *convGo_convFilter) ProcessExpUnwrap(node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    {
        _def := node.FP.Get_default()
        if _def != nil {
            def := _def.(*Nodes_Node)
            self.FP.Write("Lns_unwrapDefault( ")
            convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(", ")
            convGo_filter_1171_(def, self, &node.Nodes_Node)
        } else {
            self.FP.Write("Lns_unwrap( ")
            convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
        }
    }
    self.FP.Write(")")
    self.FP.outputAny2Type(node.FP.Get_expType())
}

// 3429: decl @lune.@base.@convGo.convFilter.processExpToDDD
func (self *convGo_convFilter) ProcessExpToDDD(node *Nodes_ExpToDDDNode,_opt LnsAny) {
    if Lns_isCondTrue( node.FP.Get_expList().FP.Get_mRetExp()){
        convGo_filter_1171_(&node.FP.Get_expList().Nodes_Node, self, &node.Nodes_Node)
    } else { 
        self.FP.Write("[]LnsAny{ ")
        convGo_filter_1171_(&node.FP.Get_expList().Nodes_Node, self, &node.Nodes_Node)
        self.FP.Write("}")
    }
}

// 3441: decl @lune.@base.@convGo.convFilter.processExpNew
func (self *convGo_convFilter) ProcessExpNew(node *Nodes_ExpNewNode,_opt LnsAny) {
    var className string
    className = self.FP.getTypeSymbol(node.FP.Get_expType())
    if self.FP.isSamePackage(node.FP.Get_expType().FP.GetModule()){
        self.FP.Write(Lns_getVM().String_format("New%s(", []LnsAny{className}))
    } else { 
        {
            _refTypeNode := Nodes_RefTypeNodeDownCastF(node.FP.Get_symbol().FP)
            if _refTypeNode != nil {
                refTypeNode := _refTypeNode.(*Nodes_RefTypeNode)
                {
                    _refNode := Nodes_RefFieldNodeDownCastF(refTypeNode.FP.Get_name().FP)
                    if _refNode != nil {
                        refNode := _refNode.(*Nodes_RefFieldNode)
                        convGo_filter_1171_(refNode.FP.Get_prefix(), self, &node.Nodes_Node)
                        self.FP.Write(".")
                    }
                }
            }
        }
        self.FP.Write(Lns_getVM().String_format("New%s(", []LnsAny{className}))
    }
    {
        _argList := node.FP.Get_argList()
        if _argList != nil {
            argList := _argList.(*Nodes_ExpListNode)
            var scope *Ast_Scope
            scope = Lns_unwrap( node.FP.Get_expType().FP.Get_scope()).(*Ast_Scope)
            var initFuncType *Ast_TypeInfo
            initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField("__init", true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
            self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), argList), initFuncType.FP.Get_argTypeInfoList(), argList)
        }
    }
    self.FP.Write(")")
}

// 3467: decl @lune.@base.@convGo.convFilter.outputIFMethods
func (self *convGo_convFilter) outputIFMethods(node *Nodes_DeclClassNode) {
    self.FP.PushIndent(nil)
    var name2MtdType *LnsMap
    name2MtdType = NewLnsMap( map[LnsAny]LnsAny{})
    var scope *Ast_Scope
    scope = Lns_unwrap( node.FP.Get_expType().FP.Get_scope()).(*Ast_Scope)
    scope.FP.FilterTypeInfoField(true, scope, Ast_ScopeAccess__Normal, Ast_filterForm(func(symbolInfo *Ast_SymbolInfo) bool {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mtd) &&
            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_name() != "__init") &&
            Lns_GetEnv().SetStackVal( Lns_op_not(symbolInfo.FP.Get_staticFlag())) ).(bool)){
            name2MtdType.Set(symbolInfo.FP.Get_name(),symbolInfo.FP.Get_typeInfo())
        }
        return true
    }))
    {
        __collection20146 := name2MtdType
        __sorted20146 := __collection20146.CreateKeyListStr()
        __sorted20146.Sort( LnsItemKindStr, nil )
        for _, _name := range( __sorted20146.Items ) {
            typeInfo := __collection20146.Items[ _name ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            name := _name.(string)
            self.FP.Write(Lns_getVM().String_format("%s(", []LnsAny{self.FP.getSymbol(&convGo_SymbolKind__Func{typeInfo}, name)}))
            for _index, _argType := range( typeInfo.FP.Get_argTypeInfoList().Items ) {
                index := _index + 1
                argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if index != 1{
                    self.FP.Write(", ")
                }
                self.FP.Write(Lns_getVM().String_format("arg%d %s", []LnsAny{index, self.FP.type2gotype(argType)}))
            }
            self.FP.Write(")")
            self.FP.outputRetType(typeInfo.FP.Get_retTypeInfoList())
            self.FP.Writeln("")
        }
    }
    self.FP.PopIndent()
}

// 3497: decl @lune.@base.@convGo.convFilter.outputMethodIF
func (self *convGo_convFilter) outputMethodIF(node *Nodes_DeclClassNode) {
    self.FP.Write("type ")
    self.FP.Write(self.FP.getTypeSymbol(node.FP.Get_expType()))
    self.FP.Writeln("Mtd interface {")
    self.FP.outputIFMethods(node)
    self.FP.Writeln("}")
}

// 3507: decl @lune.@base.@convGo.convFilter.outputInterfaceType
func (self *convGo_convFilter) outputInterfaceType(node *Nodes_DeclClassNode) {
    self.FP.Writeln(Lns_getVM().String_format("type %s interface {", []LnsAny{self.FP.getTypeSymbol(node.FP.Get_expType())}))
    self.FP.PushIndent(nil)
    self.FP.outputIFMethods(node)
    self.FP.PopIndent()
    self.FP.Writeln("}")
    var typeName string
    typeName = self.FP.type2gotype(node.FP.Get_expType())
    self.FP.Writeln(Lns_getVM().String_format("func Lns_cast2%s( obj LnsAny ) LnsAny {", []LnsAny{typeName}))
    self.FP.Writeln(Lns_getVM().String_format("    if _, ok := obj.(%s); ok { ", []LnsAny{typeName}))
    self.FP.Writeln("        return obj")
    self.FP.Writeln("    }")
    self.FP.Writeln("    return nil")
    self.FP.Writeln("}")
}

// 3532: decl @lune.@base.@convGo.convFilter.outputClassType
func (self *convGo_convFilter) outputClassType(node *Nodes_DeclClassNode) {
    self.FP.Write("type ")
    self.FP.Write(self.FP.getTypeSymbol(node.FP.Get_expType()))
    self.FP.Writeln(" struct {")
    self.FP.PushIndent(nil)
    if node.FP.Get_expType().FP.HasBase(){
        var superClassType *Ast_TypeInfo
        superClassType = node.FP.Get_expType().FP.Get_baseTypeInfo()
        self.FP.Writeln(self.FP.getTypeSymbol(superClassType))
    }
    for _, _memberNode := range( node.FP.Get_memberList().Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        convGo_filter_1171_(&memberNode.Nodes_Node, self, &node.Nodes_Node)
    }
    self.FP.Write("FP ")
    self.FP.Write(self.FP.getTypeSymbol(node.FP.Get_expType()))
    self.FP.Writeln("Mtd")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 3561: decl @lune.@base.@convGo.convFilter.outputToStem
func (self *convGo_convFilter) outputToStem(node *Nodes_DeclClassNode) {
    self.FP.Writeln(Lns_getVM().String_format("func %s2Stem( obj LnsAny ) LnsAny {", []LnsAny{self.FP.getTypeSymbolWithPrefix(node.FP.Get_expType())}))
    self.FP.PushIndent(nil)
    self.FP.Writeln("if obj == nil {")
    self.FP.Writeln("    return nil")
    self.FP.Writeln("}")
    self.FP.Writeln(Lns_getVM().String_format("return obj.(%s).FP", []LnsAny{self.FP.type2gotype(node.FP.Get_expType())}))
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 3574: decl @lune.@base.@convGo.convFilter.outputDownCast
func (self *convGo_convFilter) outputDownCast(node *Nodes_DeclClassNode) {
    var symbol string
    symbol = self.FP.getTypeSymbol(node.FP.Get_expType())
    self.FP.Write("type ")
    self.FP.Writeln(Lns_getVM().String_format("%sDownCast interface {", []LnsAny{symbol}))
    self.FP.PushIndent(nil)
    self.FP.Write("To")
    self.FP.Write(self.FP.getTypeSymbol(node.FP.Get_expType()))
    self.FP.Write("() *")
    self.FP.Write(self.FP.getTypeSymbol(node.FP.Get_expType()))
    self.FP.Writeln("")
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.FP.Writeln(Lns_getVM().String_format("func %sDownCastF( multi ...LnsAny ) LnsAny {", []LnsAny{symbol}))
    self.FP.PushIndent(nil)
    self.FP.Writeln("if len( multi ) == 0 { return nil }")
    self.FP.Writeln("obj := multi[ 0 ]")
    self.FP.Writeln("if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }")
    self.FP.Writeln(Lns_getVM().String_format("work, ok := obj.(%sDownCast)", []LnsAny{symbol}))
    self.FP.Writeln(Lns_getVM().String_format("if ok { return work.To%s() }", []LnsAny{symbol}))
    self.FP.Writeln("return nil")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 3599: decl @lune.@base.@convGo.convFilter.outputCastReceiver
func (self *convGo_convFilter) outputCastReceiver(node *Nodes_DeclClassNode) {
    self.FP.Write("func (obj *")
    self.FP.Write(self.FP.getTypeSymbol(node.FP.Get_expType()))
    self.FP.Write(") To")
    self.FP.Write(self.FP.getTypeSymbol(node.FP.Get_expType()))
    self.FP.Write("() *")
    self.FP.Write(self.FP.getTypeSymbol(node.FP.Get_expType()))
    self.FP.Writeln(" {")
    self.FP.PushIndent(nil)
    self.FP.Writeln("return obj")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 3613: decl @lune.@base.@convGo.convFilter.outputNewSetup
func (self *convGo_convFilter) outputNewSetup(objName string,classType *Ast_TypeInfo) {
    var className string
    className = self.FP.getTypeSymbol(classType)
    self.FP.Writeln(Lns_getVM().String_format("%s := &%s{}", []LnsAny{objName, className}))
    self.FP.Writeln(Lns_getVM().String_format("%s.FP = %s", []LnsAny{objName, objName}))
    {
        var workType *Ast_TypeInfo
        workType = classType
        for workType.FP.HasBase() {
            workType = workType.FP.Get_baseTypeInfo()
            
            var superName string
            superName = self.FP.getTypeSymbol(workType)
            self.FP.Writeln(Lns_getVM().String_format("%s.%s.FP = %s", []LnsAny{objName, superName, objName}))
        }
    }
}

// 3630: decl @lune.@base.@convGo.convFilter.outputConstructor
func (self *convGo_convFilter) outputConstructor(node *Nodes_DeclClassNode) {
    var scope *Ast_Scope
    scope = Lns_unwrap( node.FP.Get_expType().FP.Get_scope()).(*Ast_Scope)
    var initFuncType *Ast_TypeInfo
    initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField("__init", true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
    var className string
    className = self.FP.getTypeSymbol(node.FP.Get_expType())
    var ctorName string
    ctorName = self.FP.getConstrSymbol(node.FP.Get_expType())
    if Lns_op_not(node.FP.Get_expType().FP.Get_abstractFlag()){
        self.FP.Write(Lns_getVM().String_format("func New%s(", []LnsAny{className}))
        self.FP.outputDeclFuncArg(initFuncType)
        self.FP.Writeln(Lns_getVM().String_format(") *%s {", []LnsAny{className}))
        self.FP.PushIndent(nil)
        self.FP.outputNewSetup("obj", node.FP.Get_expType())
        self.FP.Write(Lns_getVM().String_format("obj.%s(", []LnsAny{ctorName}))
        for _index, _ := range( initFuncType.FP.Get_argTypeInfoList().Items ) {
            index := _index + 1
            if index != 1{
                self.FP.Write(", ")
            }
            self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
        }
        self.FP.Writeln(")")
        self.FP.Writeln("return obj")
        self.FP.PopIndent()
        self.FP.Writeln("}")
    }
    if Lns_op_not(node.FP.HasUserInit()){
        self.FP.Write(Lns_getVM().String_format("func (self *%s) %s(", []LnsAny{className, ctorName}))
        self.FP.outputDeclFuncArg(initFuncType)
        self.FP.Writeln(") {")
        self.FP.PushIndent(nil)
        var superArgNum LnsInt
        if node.FP.Get_expType().FP.HasBase(){
            var superType *Ast_TypeInfo
            superType = node.FP.Get_expType().FP.Get_baseTypeInfo()
            var baseScope *Ast_Scope
            baseScope = Lns_unwrap( superType.FP.Get_scope()).(*Ast_Scope)
            var baseInitFuncType *Ast_TypeInfo
            baseInitFuncType = Lns_unwrap( baseScope.FP.GetTypeInfoField("__init", true, baseScope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
            var superName string
            superName = self.FP.getTypeSymbol(superType)
            self.FP.Write(Lns_getVM().String_format("self.%s.%s( ", []LnsAny{superName, self.FP.getConstrSymbol(superType)}))
            {
                var _from21355 LnsInt = 1
                var _to21355 LnsInt = baseInitFuncType.FP.Get_argTypeInfoList().Len()
                for _work21355 := _from21355; _work21355 <= _to21355; _work21355++ {
                    index := _work21355
                    if index != 1{
                        self.FP.Write(",")
                    }
                    if node.FP.Get_hasOldCtor(){
                        self.FP.Write("nil")
                    } else { 
                        self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
                    }
                }
            }
            self.FP.Writeln(")")
            if node.FP.Get_hasOldCtor(){
                superArgNum = 0
                
            } else { 
                superArgNum = baseInitFuncType.FP.Get_argTypeInfoList().Len()
                
            }
        } else { 
            superArgNum = 0
            
        }
        var argIndex LnsInt
        argIndex = superArgNum + 1
        for _, _memberNode := range( node.FP.Get_memberList().Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if Lns_op_not(memberNode.FP.Get_staticFlag()){
                self.FP.Writeln(Lns_getVM().String_format("self.%s = arg%d", []LnsAny{self.FP.getSymbol(&convGo_SymbolKind__Member{Ast_isPubToExternal(memberNode.FP.Get_accessMode())}, memberNode.FP.Get_name().Txt), argIndex}))
                argIndex = argIndex + 1
                
            }
        }
        self.FP.PopIndent()
        self.FP.Writeln("}")
    }
}

// 3717: decl @lune.@base.@convGo.convFilter.outputAccessor
func (self *convGo_convFilter) OutputAccessor(node *Nodes_DeclClassNode) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType()
    if classType.FP.Get_kind() == Ast_TypeInfoKind__IF{
        return 
    }
    for _, _memberNode := range( node.FP.Get_memberList().Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var memberNameToken *Types_Token
        memberNameToken = memberNode.FP.Get_name()
        var memberName string
        memberName = memberNameToken.Txt
        var memberSym *Ast_SymbolInfo
        memberSym = memberNode.FP.Get_symbolInfo()
        if memberNode.FP.Get_getterMode() != Ast_AccessMode__None{
            var getterName string
            getterName = Lns_getVM().String_format("get_%s", []LnsAny{memberName})
            var getterSym *Ast_SymbolInfo
            getterSym = Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(classType.FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(getterName)})/* 3731:33 */)).(*Ast_SymbolInfo)
            self.FP.OutputDeclFunc(&convGo_FuncInfo__Type{getterSym.FP.Get_typeInfo()})
            var retType *Ast_TypeInfo
            retType = getterSym.FP.Get_typeInfo().FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_srcTypeInfo()
            self.FP.Write("{ return ")
            var prefix string
            if memberSym.FP.Get_staticFlag(){
                prefix = ""
                
            } else { 
                prefix = "self."
                
            }
            if retType.FP.Get_srcTypeInfo() != memberSym.FP.Get_typeInfo().FP.Get_srcTypeInfo(){
                if retType.FP.Get_kind() == Ast_TypeInfoKind__IF{
                    if Ast_isClass(memberSym.FP.Get_typeInfo().FP.Get_srcTypeInfo()){
                        self.FP.Write(Lns_getVM().String_format("%s%s.FP", []LnsAny{prefix, self.FP.getSymbolSym(memberSym)}))
                    }
                } else if Ast_isClass(retType){
                    self.FP.Write(Lns_getVM().String_format("&%s%s.%s", []LnsAny{prefix, self.FP.getSymbolSym(memberSym), self.FP.getTypeSymbol(retType)}))
                } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( retType.FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
                    Lns_GetEnv().SetStackVal( retType.FP.HasBase()) ).(bool)){
                    self.FP.Write(Lns_getVM().String_format("%s%s.FP", []LnsAny{prefix, self.FP.getSymbolSym(memberSym)}))
                    self.FP.outputStem2Type(retType)
                } else { 
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( convGo_isAnyType_1174_(retType)) &&
                        Lns_GetEnv().SetStackVal( Ast_isClass(memberSym.FP.Get_typeInfo())) ).(bool)){
                        self.FP.Write(Lns_getVM().String_format("%s%s.FP", []LnsAny{prefix, self.FP.getSymbolSym(memberSym)}))
                    } else { 
                        Util_err("not support")
                    }
                }
            } else { 
                self.FP.Write(Lns_getVM().String_format("%s%s", []LnsAny{prefix, self.FP.getSymbolSym(memberSym)}))
            }
            self.FP.Writeln(" }")
        }
        if memberNode.FP.Get_setterMode() != Ast_AccessMode__None{
            var setterName string
            setterName = Lns_getVM().String_format("set_%s", []LnsAny{memberName})
            var setterSym *Ast_SymbolInfo
            setterSym = Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(classType.FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(setterName)})/* 3771:33 */)).(*Ast_SymbolInfo)
            self.FP.OutputDeclFunc(&convGo_FuncInfo__Type{setterSym.FP.Get_typeInfo()})
            self.FP.Write(Lns_getVM().String_format("{ self.%s = arg1 ", []LnsAny{self.FP.getSymbolSym(memberSym)}))
            self.FP.Writeln("}")
        }
    }
}

// 3779: decl @lune.@base.@convGo.convFilter.outputStaticMember
func (self *convGo_convFilter) outputStaticMember(node *Nodes_DeclClassNode) {
    if self.processMode == convGo_ProcessMode__DeclClass{
        for _, _memberNode := range( node.FP.Get_memberList().Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if memberNode.FP.Get_staticFlag(){
                self.FP.Writeln(Lns_getVM().String_format("var %s %s", []LnsAny{self.FP.getSymbol(&convGo_SymbolKind__Static{node.FP.Get_expType()}, memberNode.FP.Get_name().Txt), self.FP.type2gotype(memberNode.FP.Get_expType())}))
            }
        }
        {
            _initBlock := node.FP.Get_initBlock().FP.Get_func()
            if _initBlock != nil {
                initBlock := _initBlock.(*Nodes_DeclMethodNode)
                convGo_filter_1171_(&initBlock.Nodes_Node, self, &node.Nodes_Node)
                self.FP.Writeln("")
            }
        }
    } else { 
        {
            _initBlock := node.FP.Get_initBlock().FP.Get_func()
            if _initBlock != nil {
                initBlock := _initBlock.(*Nodes_DeclMethodNode)
                self.FP.Writeln(Lns_getVM().String_format("%s()", []LnsAny{self.FP.getFuncSymbol(initBlock.FP.Get_expType())}))
            }
        }
    }
}

// 3808: decl @lune.@base.@convGo.convFilter.getFromStemName
func (self *convGo_convFilter) getFromStemName(typeInfo *Ast_TypeInfo) string {
    __func__ := "@lune.@base.@convGo.convFilter.getFromStemName"
    var workTypeInfo *Ast_TypeInfo
    workTypeInfo = convGo_getOrgTypeInfo_1257_(typeInfo)
    {
        _name := convGo_type2FromStemNameMap.Items[workTypeInfo]
        if _name != nil {
            name := _name.(string)
            return name
        }
    }
    if _switch22135 := workTypeInfo.FP.Get_kind(); _switch22135 == Ast_TypeInfoKind__List || _switch22135 == Ast_TypeInfoKind__Array {
        return "Lns_ToList"
    } else if _switch22135 == Ast_TypeInfoKind__Set {
        return "Lns_ToSet"
    } else if _switch22135 == Ast_TypeInfoKind__Map {
        return "Lns_ToLnsMap"
    } else if _switch22135 == Ast_TypeInfoKind__Class {
        return Lns_getVM().String_format("%s_FromMap", []LnsAny{self.FP.getTypeSymbol(workTypeInfo)})
    } else {
        Util_err(Lns_getVM().String_format("%s: not support -- %s", []LnsAny{__func__, Ast_TypeInfoKind_getTxt( workTypeInfo.FP.Get_kind())}))
    }
// insert a dummy
    return ""
}


// 3835: decl @lune.@base.@convGo.convFilter.outputConvItemType
func (self *convGo_convFilter) outputConvItemType(typeInfo *Ast_TypeInfo,alt2type LnsAny) {
    var workTypeInfo *Ast_TypeInfo
    workTypeInfo = typeInfo.FP.Get_srcTypeInfo().FP.Get_nonnilableType()
    if typeInfo.FP.Get_srcTypeInfo().FP.Get_nonnilableType().FP.Get_kind() == Ast_TypeInfoKind__Alternate{
        if alt2type != nil{
            alt2type_7617 := alt2type.(*LnsMap)
            {
                _alt := alt2type_7617.Items[workTypeInfo]
                if _alt != nil {
                    alt := _alt.(*Ast_TypeInfo)
                    workTypeInfo = alt
                    
                }
            }
        }
    }
    {
        _altType := Ast_AlternateTypeInfoDownCastF(workTypeInfo.FP)
        if _altType != nil {
            altType := _altType.(*Ast_AlternateTypeInfo)
            self.FP.Write(Lns_getVM().String_format("paramList[%d]", []LnsAny{altType.FP.Get_altIndex() - 1}))
        } else {
            self.FP.Writeln("Lns_ToObjParam{")
            self.FP.PushIndent(nil)
            self.FP.Write(Lns_getVM().String_format("%sSub, %s,", []LnsAny{self.FP.getFromStemName(workTypeInfo), typeInfo.FP.Get_nilable()}))
            self.FP.outputConvItemTypeList(workTypeInfo.FP.Get_itemTypeInfoList(), alt2type)
            self.FP.PopIndent()
            self.FP.Write("}")
        }
    }
}

// 3860: decl @lune.@base.@convGo.convFilter.outputConvItemTypeList
func (self *convGo_convFilter) outputConvItemTypeList(itemTypeInfoList *LnsList,alt2type LnsAny) {
    if itemTypeInfoList.Len() > 0{
        self.FP.Write("[]Lns_ToObjParam{")
        self.FP.PushIndent(nil)
        for _index, _itemType := range( itemTypeInfoList.Items ) {
            index := _index + 1
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index > 1{
                self.FP.Write(",")
            }
            self.FP.outputConvItemType(itemType, alt2type)
        }
        self.FP.PopIndent()
        self.FP.Write("}")
    } else { 
        self.FP.Write("nil")
    }
}

// 3880: decl @lune.@base.@convGo.convFilter.outputAlter2MapFunc
func (self *convGo_convFilter) outputAlter2MapFunc(alt2Map *LnsMap) {
    __func__ := "@lune.@base.@convGo.convFilter.outputAlter2MapFunc"
    self.FP.Write("{")
    for _altType, _assinType := range( alt2Map.Items ) {
        altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        assinType := _assinType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if altType.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
            if assinType.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
                Util_err(Lns_getVM().String_format("not support: %s", []LnsAny{__func__}))
            } else { 
                self.FP.outputConvItemType(assinType, alt2Map)
            }
        }
    }
    self.FP.Write("}")
}

// 3899: decl @lune.@base.@convGo.convFilter.outputAsyncItem
func (self *convGo_convFilter) outputAsyncItem(node *Nodes_DeclClassNode) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType()
    var classScope *Ast_Scope
    classScope = Lns_unwrap( classType.FP.Get_scope()).(*Ast_Scope)
    var createSym *Ast_SymbolInfo
    createSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild("_createPipe")).(*Ast_SymbolInfo)
    self.FP.OutputDeclFunc(&convGo_FuncInfo__WithClass{classType, createSym.FP.Get_typeInfo()})
    self.FP.Writeln("{")
    self.FP.Writeln("   return NewLnspipe( arg1 )")
    self.FP.Writeln("}")
}

// 3909: decl @lune.@base.@convGo.convFilter.outputMapping
func (self *convGo_convFilter) outputMapping(node *Nodes_DeclClassNode) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType()
    var className string
    className = self.FP.getTypeSymbol(classType)
    self.FP.Writeln(Lns_getVM().String_format("func (self *%s) ToMapSetup( obj *LnsMap ) *LnsMap {", []LnsAny{className}))
    self.FP.PushIndent(nil)
    for _, _memberNode := range( node.FP.Get_memberList().Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if Lns_op_not(memberNode.FP.Get_staticFlag()){
            self.FP.Writeln(Lns_getVM().String_format("obj.Items[\"%s\"] = Lns_ToCollection( self.%s )", []LnsAny{memberNode.FP.Get_symbolInfo().FP.Get_name(), self.FP.getSymbolSym(memberNode.FP.Get_symbolInfo())}))
        }
    }
    if classType.FP.HasBase(){
        self.FP.Writeln(Lns_getVM().String_format("return self.%s.ToMapSetup( obj )", []LnsAny{self.FP.getTypeSymbol(classType.FP.Get_baseTypeInfo())}))
    } else { 
        self.FP.Writeln("return obj")
    }
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.FP.Writeln(Lns_getVM().String_format("func (self *%s) ToMap() *LnsMap {", []LnsAny{className}))
    self.FP.Writeln("    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )")
    self.FP.Writeln("}")
    var fromStemName string
    fromStemName = self.FP.getFromStemName(classType)
    var classScope *Ast_Scope
    classScope = Lns_unwrap( classType.FP.Get_scope()).(*Ast_Scope)
    if Lns_op_not(classType.FP.Get_abstractFlag()){
        {
            var fromMapSym *Ast_SymbolInfo
            fromMapSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild("_fromMap")).(*Ast_SymbolInfo)
            self.FP.Writeln(Lns_getVM().String_format("func %s(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){", []LnsAny{self.FP.getSymbolSym(fromMapSym)}))
            self.FP.Writeln(Lns_getVM().String_format("   return %s( arg1, paramList )", []LnsAny{fromStemName}))
            self.FP.Writeln("}")
        }
        {
            var fromStemSym *Ast_SymbolInfo
            fromStemSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild("_fromStem")).(*Ast_SymbolInfo)
            self.FP.Writeln(Lns_getVM().String_format("func %s(arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){", []LnsAny{self.FP.getSymbolSym(fromStemSym)}))
            self.FP.Writeln(Lns_getVM().String_format("   return %s( arg1, paramList )", []LnsAny{fromStemName}))
            self.FP.Writeln("}")
        }
        self.FP.Writeln(Lns_getVM().String_format("func %s( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {", []LnsAny{fromStemName}))
        self.FP.PushIndent(nil)
        self.FP.Writeln(Lns_getVM().String_format("_,conv,mess := %sSub(obj,false, paramList);", []LnsAny{fromStemName}))
        self.FP.Writeln("return conv,mess")
        self.FP.PopIndent()
        self.FP.Writeln("}")
        self.FP.Writeln(Lns_getVM().String_format("func %sSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {", []LnsAny{fromStemName}))
        self.FP.PushIndent(nil)
        self.FP.Writeln("var objMap *LnsMap")
        self.FP.Writeln("if work, ok := obj.(*LnsMap); !ok {")
        self.FP.Writeln("   return false, nil, \"no map -- \" + Lns_ToString(obj)")
        self.FP.Writeln("} else {")
        self.FP.Writeln("   objMap = work")
        self.FP.Writeln("}")
        self.FP.outputNewSetup("newObj", classType)
        self.FP.Writeln(Lns_getVM().String_format("return %sMain( newObj, objMap, paramList )", []LnsAny{fromStemName}))
        self.FP.PopIndent()
        self.FP.Writeln("}")
    }
    self.FP.Writeln(Lns_getVM().String_format("func %sMain( newObj %s, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {", []LnsAny{fromStemName, self.FP.type2gotype(classType)}))
    self.FP.PushIndent(nil)
    if classType.FP.HasBase(){
        self.FP.Writeln(Lns_getVM().String_format("if ok,_,mess := %sMain( &newObj.%s, objMap, paramList ); !ok {", []LnsAny{self.FP.getFromStemName(classType.FP.Get_baseTypeInfo()), self.FP.getTypeSymbol(classType.FP.Get_baseTypeInfo())}))
        self.FP.Writeln("    return false,nil,mess")
        self.FP.Writeln("}")
    }
    for _, _memberNode := range( node.FP.Get_memberList().Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if Lns_op_not(memberNode.FP.Get_staticFlag()){
            var memberName string
            memberName = self.FP.getSymbolSym(memberNode.FP.Get_symbolInfo())
            self.FP.Write("if ok,conv,mess := ")
            if memberNode.FP.Get_expType().FP.Get_nonnilableType().FP.Get_kind() == Ast_TypeInfoKind__Alternate{
                for _index, _itemType := range( classType.FP.Get_itemTypeInfoList().Items ) {
                    index := _index + 1
                    itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if itemType == memberNode.FP.Get_expType().FP.Get_srcTypeInfo(){
                        self.FP.Write(Lns_getVM().String_format("paramList[%d].Func( objMap.Items[\"%s\"], %s, paramList[%d].Child", []LnsAny{index - 1, memberNode.FP.Get_symbolInfo().FP.Get_name(), memberNode.FP.Get_expType().FP.Get_nilable(), index - 1}))
                    }
                }
            } else { 
                self.FP.Write(Lns_getVM().String_format("%sSub( objMap.Items[\"%s\"], %s, ", []LnsAny{self.FP.getFromStemName(memberNode.FP.Get_expType()), memberNode.FP.Get_symbolInfo().FP.Get_name(), memberNode.FP.Get_expType().FP.Get_nilable()}))
                self.FP.outputConvItemTypeList(memberNode.FP.Get_expType().FP.Get_itemTypeInfoList(), nil)
            }
            self.FP.Writeln("); !ok {")
            self.FP.Writeln(Lns_getVM().String_format("   return false,nil,\"%s:\" + mess.(string)", []LnsAny{memberNode.FP.Get_symbolInfo().FP.Get_name()}))
            self.FP.Writeln("} else {")
            self.FP.Write(Lns_getVM().String_format("   newObj.%s = conv", []LnsAny{memberName}))
            self.FP.outputAny2Type(memberNode.FP.Get_expType())
            self.FP.Writeln("")
            self.FP.Writeln("}")
        }
    }
    self.FP.Writeln("return true, newObj, nil")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 4035: decl @lune.@base.@convGo.convFilter.outputDummyAbstractMethod
func (self *convGo_convFilter) outputDummyAbstractMethod(classType *Ast_TypeInfo,methodType *Ast_TypeInfo) {
    self.FP.OutputDeclFunc(&convGo_FuncInfo__WithClass{classType, methodType})
    self.FP.Writeln("{")
    self.FP.outputDummyReturn(methodType.FP.Get_retTypeInfoList())
    self.FP.Writeln("}")
}

// 4046: decl @lune.@base.@convGo.convFilter.outputDummyAbstractMethodOfClass
func (self *convGo_convFilter) outputDummyAbstractMethodOfClass(classType *Ast_TypeInfo) {
    var name2typeMap *LnsMap
    name2typeMap = NewLnsMap( map[LnsAny]LnsAny{})
    var scope *Ast_Scope
    scope = Lns_unwrap( classType.FP.Get_scope()).(*Ast_Scope)
    scope.FP.FilterSymbolInfoIfField(scope, Ast_ScopeAccess__Normal, Ast_filterForm(func(symbolInfo *Ast_SymbolInfo) bool {
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mtd) &&
            Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_typeInfo().FP.Get_abstractFlag()) ).(bool)){
            {
                _methodType := scope.FP.GetTypeInfoChild(symbolInfo.FP.Get_name())
                if _methodType != nil {
                    methodType := _methodType.(*Ast_TypeInfo)
                    if methodType.FP.Get_abstractFlag(){
                        name2typeMap.Set(symbolInfo.FP.Get_name(),symbolInfo.FP.Get_typeInfo())
                    }
                } else {
                    name2typeMap.Set(symbolInfo.FP.Get_name(),symbolInfo.FP.Get_typeInfo())
                }
            }
        }
        return true
    }))
    {
        __collection23471 := name2typeMap
        __sorted23471 := __collection23471.CreateKeyListStr()
        __sorted23471.Sort( LnsItemKindStr, nil )
        for _, ___key23471 := range( __sorted23471.Items ) {
            methodType := __collection23471.Items[ ___key23471 ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            self.FP.outputDummyAbstractMethod(classType, methodType)
        }
    }
}

// 4076: decl @lune.@base.@convGo.convFilter.outputAdvertise
func (self *convGo_convFilter) outputAdvertise(node *Nodes_DeclClassNode) {
    __func__ := "@lune.@base.@convGo.convFilter.outputAdvertise"
    var methodNameSet *LnsSet
    methodNameSet = node.FP.CreateMethodNameSetWithoutAdv()
    for _, _adv := range( node.FP.Get_advertiseList().Items ) {
        adv := _adv.(Nodes_AdvertiseInfoDownCast).ToNodes_AdvertiseInfo()
        if adv.FP.Get_prefix() != ""{
            Util_err(Lns_getVM().String_format("%s: not support advertise with prefix", []LnsAny{__func__}))
        }
        {
            _scope := adv.FP.Get_member().FP.Get_expType().FP.Get_scope()
            if _scope != nil {
                scope := _scope.(*Ast_Scope)
                scope.FP.FilterTypeInfoField(true, scope, Ast_ScopeAccess__Normal, Ast_filterForm(func(symbol *Ast_SymbolInfo) bool {
                    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( symbol.FP.Get_kind() == Ast_SymbolKind__Mtd) &&
                        Lns_GetEnv().SetStackVal( symbol.FP.Get_name() != "__init") &&
                        Lns_GetEnv().SetStackVal( Lns_op_not(methodNameSet.Has(symbol.FP.Get_name()))) &&
                        Lns_GetEnv().SetStackVal( Lns_op_not(symbol.FP.Get_staticFlag())) ).(bool)){
                        var funcType *Ast_TypeInfo
                        funcType = symbol.FP.Get_typeInfo()
                        self.FP.OutputDeclFunc(&convGo_FuncInfo__WithClass{node.FP.Get_expType(), funcType})
                        self.FP.Writeln(" {")
                        if funcType.FP.Get_retTypeInfoList().Len() > 0{
                            self.FP.Write("    return ")
                        }
                        self.FP.Write(Lns_getVM().String_format("self.%s. ", []LnsAny{self.FP.getSymbolSym(adv.FP.Get_member().FP.Get_symbolInfo())}))
                        if adv.FP.Get_member().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Class{
                            self.FP.Write("FP.")
                        }
                        self.FP.Write(Lns_getVM().String_format("%s( ", []LnsAny{self.FP.getSymbolSym(symbol)}))
                        for _index, _ := range( funcType.FP.Get_argTypeInfoList().Items ) {
                            index := _index + 1
                            if index > 1{
                                self.FP.Write(",")
                            }
                            self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
                        }
                        self.FP.Writeln(")")
                        self.FP.Writeln("}")
                    }
                    return true
                }))
            }
        }
    }
}

// 4117: decl @lune.@base.@convGo.convFilter.processProtoClass
func (self *convGo_convFilter) ProcessProtoClass(node *Nodes_ProtoClassNode,_opt LnsAny) {
}

// 4122: decl @lune.@base.@convGo.convFilter.processDeclClass
func (self *convGo_convFilter) ProcessDeclClass(node *Nodes_DeclClassNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processDeclClass"
    if node.FP.IsModule(){
        return 
    }
    if self.processMode == convGo_ProcessMode__DeclClass{
        if _switch24001 := node.FP.Get_expType().FP.Get_kind(); _switch24001 == Ast_TypeInfoKind__Class {
            self.FP.Writeln(Lns_getVM().String_format("// declaration Class -- %s", []LnsAny{node.FP.Get_expType().FP.Get_rawTxt()}))
            self.FP.outputStaticMember(node)
            self.FP.outputMethodIF(node)
            self.FP.outputClassType(node)
            self.FP.outputToStem(node)
            self.FP.outputDownCast(node)
            self.FP.outputCastReceiver(node)
            self.FP.outputConstructor(node)
            self.FP.OutputAccessor(node)
            self.FP.outputDummyAbstractMethodOfClass(node.FP.Get_expType())
            self.FP.outputAdvertise(node)
            if node.FP.Get_expType().FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, nil){
                self.FP.outputMapping(node)
            }
            if node.FP.Get_expType().FP.IsInheritFrom(self.processInfo, Ast_builtinTypeAsyncItem, nil){
                self.FP.outputAsyncItem(node)
            }
            for _, _fieldNode := range( node.FP.Get_fieldList().Items ) {
                fieldNode := _fieldNode.(Nodes_NodeDownCast).ToNodes_Node()
                if Lns_isCondTrue( Nodes_DeclMemberNodeDownCastF(fieldNode.FP)){
                } else { 
                    convGo_filter_1171_(fieldNode, self, &node.Nodes_Node)
                    self.FP.Writeln("")
                }
            }
        } else if _switch24001 == Ast_TypeInfoKind__IF {
            self.FP.outputInterfaceType(node)
        } else {
            Util_err(Lns_getVM().String_format("%s: not support -- %s:%d", []LnsAny{__func__, Ast_TypeInfoKind_getTxt( node.FP.Get_expType().FP.Get_kind()), node.FP.Get_pos().LineNo}))
        }
    } else { 
        self.FP.outputStaticMember(node)
    }
}

// 4209: decl @lune.@base.@convGo.convFilter.outputCallPrefix
func (self *convGo_convFilter) OutputCallPrefix(threading bool,callId LnsInt,node *Nodes_Node,prefixNode LnsAny,funcSymbol *Ast_SymbolInfo)(bool, LnsAny) {
    var funcType *Ast_TypeInfo
    funcType = funcSymbol.FP.Get_typeInfo()
    var nilAccName string
    nilAccName = Lns_getVM().String_format("%s_%d", []LnsAny{Lns_car(Lns_getVM().String_gsub(self.moduleTypeInfo.FP.Get_rawTxt(),"@", "")).(string), callId})
    var callKind LnsAny
    callKind = convGo_CallKind__Normal_Obj
    var extCallFlag bool
    if prefixNode != nil{
        prefixNode_7778 := prefixNode.(*Nodes_Node)
        extCallFlag = prefixNode_7778.FP.Get_expType().FP.Get_nonnilableType().FP.Get_kind() == Ast_TypeInfoKind__Ext
        
    } else {
        extCallFlag = false
        
    }
    if extCallFlag{
        callKind = convGo_CallKind__LuaCall_Obj
        
    }
    var getEnvTxt string
    getEnvTxt = self.env.FP.GetEnv(threading)
    var processNilAcc func(workPrefixNode *Nodes_Node)
    processNilAcc = func(workPrefixNode *Nodes_Node) {
        if Lns_op_not(node.FP.HasNilAccess()){
            return 
        }
        var retNum LnsInt
        retNum = funcType.FP.Get_retTypeInfoList().Len()
        if _switch24288 := retNum; _switch24288 == 0 {
            self.FP.Write(Lns_getVM().String_format("Lns_NilAccCall0( %s, func () {", []LnsAny{getEnvTxt}))
        } else if _switch24288 == 1 {
            self.FP.Write(Lns_getVM().String_format("Lns_NilAccCall1( %s, func () LnsAny { return ", []LnsAny{getEnvTxt}))
        } else {
            if retNum <= convGo_MaxNilAccNum{
                var anys string
                anys = "LnsAny"
                {
                    var _from24223 LnsInt = 2
                    var _to24223 LnsInt = retNum
                    for _work24223 := _from24223; _work24223 <= _to24223; _work24223++ {
                        anys = Lns_getVM().String_format("%s,LnsAny", []LnsAny{anys})
                        
                    }
                }
                self.FP.Write(Lns_getVM().String_format("Lns_NilAccCall%d( %s, func () (%s) { return ", []LnsAny{retNum, getEnvTxt, anys}))
            } else { 
                var args string
                args = "LnsAny"
                {
                    var _from24266 LnsInt = 2
                    var _to24266 LnsInt = retNum
                    for _work24266 := _from24266; _work24266 <= _to24266; _work24266++ {
                        args = Lns_getVM().String_format("%s,LnsAny", []LnsAny{args})
                        
                    }
                }
                self.FP.Write(Lns_getVM().String_format("lns_NilAccCall_%s( %s, func () (%s) { return ", []LnsAny{nilAccName, getEnvTxt, args}))
            }
        }
        if extCallFlag{
            if funcType.FP.Get_retTypeInfoList().Len() > 1{
                self.FP.Write(Lns_getVM().String_format("Lns_callExt%d( ", []LnsAny{node.FP.Get_id()}))
            }
        }
        self.FP.Write(Lns_getVM().String_format("%s.NilAccPop().(%s)", []LnsAny{getEnvTxt, self.FP.type2gotype(workPrefixNode.FP.Get_expType().FP.Get_nonnilableType())}))
    }
    var closeParen bool
    closeParen = false
    if prefixNode != nil{
        prefixNode_7803 := prefixNode.(*Nodes_Node)
        if node.FP.HasNilAccess(){
            if funcType.FP.Get_retTypeInfoList().Len() >= 2{
                if funcType.FP.Get_retTypeInfoList().Len() <= convGo_MaxNilAccNum{
                    self.FP.Write(Lns_getVM().String_format("Lns_NilAccFinCall%d( ", []LnsAny{funcType.FP.Get_retTypeInfoList().Len()}))
                } else { 
                    self.FP.Write(Lns_getVM().String_format("lns_NilAccFinCall_%s(", []LnsAny{nilAccName}))
                }
                closeParen = true
                
            }
        }
        var prefixType *Ast_TypeInfo
        prefixType = prefixNode_7803.FP.Get_expType().FP.Get_nonnilableType()
        if prefixType == Ast_builtinTypeString{
            if node.FP.HasNilAccess(){
                Util_err("not support nilAccName")
            }
            {
                _runtime := self.FP.getVM(threading, funcType)
                if _runtime != nil {
                    runtime := _runtime.(string)
                    callKind = &convGo_CallKind__RuntimeCall{prefixNode_7803}
                    
                    self.FP.Write(runtime)
                }
            }
        } else { 
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(funcType.FP.Get_staticFlag())) ||
                Lns_GetEnv().SetStackVal( funcSymbol.FP.Get_kind() == Ast_SymbolKind__Mbr) ).(bool){
                if node.FP.HasNilAccess(){
                    if Lns_op_not(prefixNode_7803.FP.HasNilAccess()){
                        self.FP.Write(Lns_getVM().String_format("%s.NilAccFin(", []LnsAny{getEnvTxt}))
                        self.FP.Write(Lns_getVM().String_format("%s.NilAccPush(", []LnsAny{getEnvTxt}))
                        convGo_filter_1171_(prefixNode_7803, self, node)
                        self.FP.Writeln(") && ")
                    } else { 
                        convGo_filter_1171_(prefixNode_7803, self, node)
                        self.FP.Writeln("&&")
                    }
                } else { 
                    if extCallFlag{
                        if funcType.FP.Get_retTypeInfoList().Len() > 1{
                            self.FP.Write(Lns_getVM().String_format("Lns_callExt%d( ", []LnsAny{node.FP.Get_id()}))
                        }
                    }
                    convGo_filter_1171_(prefixNode_7803, self, node)
                }
            } else { 
                self.FP.outputModule(funcType.FP.GetModule(), true)
            }
            processNilAcc(prefixNode_7803)
            if prefixType.FP.Get_kind() == Ast_TypeInfoKind__Ext{
                self.FP.Write(Lns_getVM().String_format(".CallMethod( \"%s\", Lns_2DDD", []LnsAny{funcSymbol.FP.Get_name()}))
            } else { 
                var prefixKind LnsInt
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( prefixType.FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
                    Lns_GetEnv().SetStackVal( prefixType.FP.HasBase()) ).(bool)){
                    prefixKind = prefixType.FP.Get_baseTypeInfo().FP.Get_kind()
                    
                } else { 
                    prefixKind = prefixType.FP.Get_kind()
                    
                }
                if Ast_isBuiltin(funcType.FP.Get_typeId()){
                    if _switch24798 := prefixKind; _switch24798 == Ast_TypeInfoKind__Class {
                        self.FP.Write(Lns_getVM().String_format(".FP.%s", []LnsAny{self.FP.getSymbolSym(funcSymbol)}))
                    } else {
                        var builtinFuncs *TransUnit_BuiltinFuncType
                        builtinFuncs = TransUnit_getBuiltinFunc()
                        {
                            _runtime := self.FP.getVM(threading, funcType)
                            if _runtime != nil {
                                runtime := _runtime.(string)
                                self.FP.Write(runtime)
                            } else {
                                if _switch24775 := funcType; _switch24775 == builtinFuncs.List_sort || _switch24775 == builtinFuncs.Array_sort {
                                    callKind = &convGo_CallKind__SortCall{prefixType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()}
                                    
                                }
                                self.FP.Write(Lns_getVM().String_format(".%s", []LnsAny{self.FP.getSymbolSym(funcSymbol)}))
                            }
                        }
                    }
                } else { 
                    if _switch24884 := funcType.FP.Get_kind(); _switch24884 == Ast_TypeInfoKind__Method {
                        if _switch24852 := prefixKind; _switch24852 == Ast_TypeInfoKind__Class {
                            self.FP.Write(Lns_getVM().String_format(".FP.%s", []LnsAny{self.FP.getSymbolSym(funcSymbol)}))
                        } else {
                            self.FP.Write(Lns_getVM().String_format(".%s", []LnsAny{self.FP.getSymbolSym(funcSymbol)}))
                        }
                    } else {
                        if funcSymbol.FP.Get_kind() == Ast_SymbolKind__Mbr{
                            self.FP.Write(".")
                        }
                        self.FP.Write(self.FP.getSymbolSym(funcSymbol))
                    }
                }
            }
        }
    }
    return closeParen, callKind
}

// 4386: decl @lune.@base.@convGo.convFilter.processExpCall
func (self *convGo_convFilter) ProcessExpCall(node *Nodes_ExpCallNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpCall"
    opt := _opt.(*ConvGo_Opt)
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_func().FP.Get_expType().FP.Get_nonnilableType()
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( funcType.FP.Get_kind() == Ast_TypeInfoKind__Method) &&
        Lns_GetEnv().SetStackVal( funcType.FP.Get_parentInfo().FP.Get_kind() == Ast_TypeInfoKind__Enum) &&
        Lns_GetEnv().SetStackVal( funcType.FP.Get_rawTxt() == "get__txt") ).(bool)){
        self.FP.Write(Lns_getVM().String_format("%s(", []LnsAny{self.FP.getEnumGetTxtSym(Lns_unwrap( Ast_EnumTypeInfoDownCastF(funcType.FP.Get_parentInfo().FP.Get_aliasSrc().FP)).(*Ast_EnumTypeInfo))}))
        var fieldNode *Nodes_RefFieldNode
        
        {
            _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func().FP)
            if _fieldNode == nil{
                Util_err("not support -- __func__")
            } else {
                fieldNode = _fieldNode.(*Nodes_RefFieldNode)
            }
        }
        convGo_filter_1171_(fieldNode.FP.Get_prefix(), self, &node.Nodes_Node)
        self.FP.Write(")")
        return 
    }
    var retGenerics bool
    if opt.Parent.FP.Get_kind() == Nodes_NodeKind_get_StmtExp(){
        retGenerics = false
        
    } else { 
        retGenerics = convGo_isRetGenerics_1358_(node)
        
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( retGenerics) &&
            Lns_GetEnv().SetStackVal( funcType.FP.Get_retTypeInfoList().Len() != 1) ).(bool)){
            self.FP.Write(Lns_getVM().String_format("%s(", []LnsAny{self.FP.getConvGenericsName(&node.Nodes_Node)}))
        }
    }
    var closeParen bool
    closeParen = false
    var callKind LnsAny
    callKind = convGo_CallKind__Normal_Obj
    var withPrefix bool
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func().FP)
        if _fieldNode != nil {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            withPrefix = true
            
            closeParen, callKind = self.FP.OutputCallPrefix(node.FP.IsThreading(), node.FP.Get_id(), &fieldNode.Nodes_Node, fieldNode.FP.Get_prefix(), Lns_unwrap( fieldNode.FP.Get_symbolInfo()).(*Ast_SymbolInfo))
            
        } else {
            withPrefix = false
            
            if Ast_isBuiltin(funcType.FP.Get_typeId()){
                _ = TransUnit_getBuiltinFunc()
                {
                    _runtime := self.FP.getVM(node.FP.IsThreading(), funcType)
                    if _runtime != nil {
                        runtime := _runtime.(string)
                        self.FP.Write(runtime)
                    } else {
                        if _switch25238 := funcType.FP.Get_srcTypeInfo(); _switch25238 == Ast_builtinTypeForm {
                            convGo_filter_1171_(node.FP.Get_func(), self, &node.Nodes_Node)
                            callKind = convGo_CallKind__FormCall_Obj
                            
                        } else {
                            Util_err(Lns_getVM().String_format("%s: not support -- %s:%d", []LnsAny{__func__, funcType.FP.GetTxt(nil, nil, nil), node.FP.Get_pos().LineNo}))
                        }
                    }
                }
            } else { 
                if funcType.FP.Get_kind() == Ast_TypeInfoKind__Ext{
                    self.FP.Write(Lns_getVM().String_format("%s.RunLoadedfunc", []LnsAny{self.env.FP.GetLuavm(node.FP.IsThreading())}))
                    callKind = convGo_CallKind__RunLoaded_Obj
                    
                } else { 
                    convGo_filter_1171_(node.FP.Get_func(), self, &node.Nodes_Node)
                }
            }
        }
    }
    self.FP.Write("(")
    var closeTxt LnsAny
    closeTxt = nil
    switch _exp25473 := callKind.(type) {
    case *convGo_CallKind__RuntimeCall:
    prefixNode := _exp25473.Val1
        convGo_filter_1171_(prefixNode, self, node.FP.Get_func())
        if Lns_isCondTrue( node.FP.Get_argList()){
            self.FP.Write(",")
        }
    case *convGo_CallKind__FormCall:
        self.FP.Write("Lns_2DDD(")
    case *convGo_CallKind__RunLoaded:
        convGo_filter_1171_(node.FP.Get_func(), self, &node.Nodes_Node)
        self.FP.Write(",")
        if Lns_op_not(node.FP.Get_argList()){
            self.FP.Write("[]LnsAny{}")
        } else { 
            self.FP.Write("Lns_2DDD(")
            closeTxt = ")"
            
        }
    case *convGo_CallKind__SortCall:
    typeInfo := _exp25473.Val1
        self.FP.Write(Lns_getVM().String_format("%s, ", []LnsAny{convGo_getLnsItemKind_1815_(typeInfo)}))
    case *convGo_CallKind__BuiltinCall:
    packName := _exp25473.Val1
    funcname := _exp25473.Val2
        closeTxt = "}"
        
        self.FP.Write(Lns_getVM().String_format("\"%s\", \"%s\"", []LnsAny{packName, funcname}))
        if Lns_isCondTrue( node.FP.Get_argList()){
            self.FP.Write(", []LnsAny{")
        }
    case *convGo_CallKind__LuaCall:
        closeTxt = ")"
        
    }
    {
        _argList := node.FP.Get_argList()
        if _argList != nil {
            argList := _argList.(*Nodes_ExpListNode)
            self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), argList), funcType.FP.Get_argTypeInfoList(), argList)
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( funcType.FP.Get_kind() == Ast_TypeInfoKind__Func) &&
        Lns_GetEnv().SetStackVal( funcType.FP.Get_staticFlag()) &&
        Lns_GetEnv().SetStackVal( funcType.FP.Get_parentInfo().FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, nil)) &&
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( funcType.FP.Get_rawTxt() == "_fromMap") ||
            Lns_GetEnv().SetStackVal( funcType.FP.Get_rawTxt() == "_fromStem") ).(bool))) ).(bool)){
        var fieldNode *Nodes_RefFieldNode
        
        {
            _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func().FP)
            if _fieldNode == nil{
                Util_err("not support -- __func__")
            } else {
                fieldNode = _fieldNode.(*Nodes_RefFieldNode)
            }
        }
        self.FP.Write(",")
        self.FP.outputConvItemTypeList(funcType.FP.Get_parentInfo().FP.Get_itemTypeInfoList(), fieldNode.FP.Get_prefix().FP.Get_expType().FP.CreateAlt2typeMap(false))
    }
    if closeTxt != nil{
        closeTxt_7895 := closeTxt.(string)
        self.FP.Write(closeTxt_7895)
    }
    self.FP.Write(")")
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( callKind == convGo_CallKind__LuaCall_Obj) ||
        Lns_GetEnv().SetStackVal( callKind == convGo_CallKind__RunLoaded_Obj) ).(bool){
        if funcType.FP.Get_retTypeInfoList().Len() == 1{
            if opt.Parent.FP.Get_kind() != Nodes_NodeKind_get_StmtExp(){
                self.FP.Write("[0]")
                var retTypeList *LnsList
                retTypeList = Lns_unwrap( Lns_car(Ast_convToExtTypeList(self.processInfo, funcType.FP.Get_retTypeInfoList()))).(*LnsList)
                self.FP.outputAny2Type(retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
            }
        } else if funcType.FP.Get_retTypeInfoList().Len() > 1{
            self.FP.Write(")")
        }
    }
    if retGenerics{
        if funcType.FP.Get_retTypeInfoList().Len() == 1{
            var retType *Ast_TypeInfo
            retType = funcType.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if convGo_isAnyType_1174_(retType){
                if Lns_op_not(convGo_isAnyType_1174_(node.FP.Get_expType())){
                    self.FP.outputAny2Type(node.FP.Get_expType())
                }
            } else if retType.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
                if retType.FP.Get_srcTypeInfo() != node.FP.Get_expType().FP.Get_srcTypeInfo(){
                    self.FP.Write(".FP")
                    self.FP.outputStem2Type(node.FP.Get_expType())
                }
            }
        } else { 
            self.FP.Write(")")
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( withPrefix) &&
        Lns_GetEnv().SetStackVal( node.FP.HasNilAccess()) ).(bool)){
        self.FP.Write("})")
        self.FP.Write(Lns_getVM().String_format("/* %d:%d */", []LnsAny{node.FP.Get_pos().LineNo, node.FP.Get_pos().Column}))
        if opt.Parent.FP.HasNilAccess(){
        } else { 
            self.FP.Write(")")
        }
        if closeParen{
            self.FP.Write(")")
        }
    }
    if callKind == convGo_CallKind__FormCall_Obj{
        self.FP.Write(")")
    }
}

// 4579: decl @lune.@base.@convGo.convFilter.processExpMRet
func (self *convGo_convFilter) ProcessExpMRet(node *Nodes_ExpMRetNode,_opt LnsAny) {
    convGo_filter_1171_(node.FP.Get_mRet(), self, &node.Nodes_Node)
}

// 4586: decl @lune.@base.@convGo.convFilter.processExpAccessMRet
func (self *convGo_convFilter) ProcessExpAccessMRet(node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
}

// 4594: decl @lune.@base.@convGo.convFilter.processExpList
func (self *convGo_convFilter) ProcessExpList(node *Nodes_ExpListNode,_opt LnsAny) {
    for _index, _exp := range( node.FP.Get_expList().Items ) {
        index := _index + 1
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        if index != 1{
            self.FP.Write(", ")
        }
        {
            _mRetExp := node.FP.Get_mRetExp()
            if _mRetExp != nil {
                mRetExp := _mRetExp.(*Nodes_MRetExp)
                if mRetExp.FP.Get_index() == index{
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( index == 1) ||
                        Lns_GetEnv().SetStackVal( exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__DDD) ).(bool){
                        convGo_filter_1171_(exp, self, &node.Nodes_Node)
                    } else { 
                        self.FP.Write("Lns_2DDD(")
                        convGo_filter_1171_(exp, self, &node.Nodes_Node)
                        self.FP.Write(")")
                    }
                    break
                }
            }
        }
        convGo_filter_1171_(exp, self, &node.Nodes_Node)
    }
}

// 4618: decl @lune.@base.@convGo.convFilter.processExpOp1
func (self *convGo_convFilter) ProcessExpOp1(node *Nodes_ExpOp1Node,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpOp1"
    if _switch26286 := node.FP.Get_op().Txt; _switch26286 == "~" {
        self.FP.Write("^")
        convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
    } else if _switch26286 == "+" || _switch26286 == "-" {
        self.FP.Write(node.FP.Get_op().Txt)
        convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
    } else if _switch26286 == "not" {
        self.FP.Write("Lns_op_not(")
        convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
        self.FP.Write(")")
    } else if _switch26286 == "#" {
        if _switch26265 := node.FP.Get_exp().FP.Get_expType().FP.Get_kind(); _switch26265 == Ast_TypeInfoKind__List {
            convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(".Len()")
        } else if _switch26265 == Ast_TypeInfoKind__Ext {
            if _switch26232 := node.FP.Get_exp().FP.Get_expType().FP.Get_extedType().FP.Get_kind(); _switch26232 == Ast_TypeInfoKind__List {
                convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
                self.FP.Write(".Len()")
            } else {
                Util_err(Lns_getVM().String_format("%s: not support -- %s", []LnsAny{__func__, node.FP.Get_exp().FP.Get_expType().FP.GetTxt(nil, nil, nil)}))
            }
        } else {
            self.FP.Write("len(")
            convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(")")
        }
    } else {
        Util_err(Lns_getVM().String_format("%s: not support -- %s", []LnsAny{__func__, node.FP.Get_op().Txt}))
    }
}

// 4666: decl @lune.@base.@convGo.convFilter.processExpMultiTo1
func (self *convGo_convFilter) ProcessExpMultiTo1(node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    self.FP.Write("Lns_car(")
    convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Write(")")
    self.FP.outputAny2Type(node.FP.Get_expType())
}

// 4678: decl @lune.@base.@convGo.convFilter.processExpCast
func (self *convGo_convFilter) ProcessExpCast(node *Nodes_ExpCastNode,_opt LnsAny) {
    if _switch26693 := node.FP.Get_castKind(); _switch26693 == Nodes_CastKind__Force {
        if convGo_isAnyType_1174_(node.FP.Get_exp().FP.Get_expType()){
            if _switch26459 := node.FP.Get_castType(); _switch26459 == Ast_builtinTypeInt {
                self.FP.Write("Lns_forceCastInt(")
                convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
                self.FP.Write(")")
            } else if _switch26459 == Ast_builtinTypeReal {
                self.FP.Write("Lns_forceCastReal(")
                convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
                self.FP.Write(")")
            } else {
                convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
                self.FP.outputAny2Type(node.FP.Get_castType())
            }
        } else { 
            self.FP.Write(Lns_getVM().String_format("(%s)(", []LnsAny{self.FP.type2gotype(node.FP.Get_castType())}))
            convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(")")
        }
    } else if _switch26693 == Nodes_CastKind__Implicit {
        if node.FP.Get_exp().FP.Get_expTypeList().Len() > 1{
            convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
        } else { 
            self.FP.outputImplicitCast(node.FP.Get_castType(), node.FP.Get_exp(), node)
        }
    } else if _switch26693 == Nodes_CastKind__Normal {
        var typeName string
        typeName = self.FP.getTypeSymbol(node.FP.Get_castType())
        var castType *Ast_TypeInfo
        castType = node.FP.Get_castType().FP.Get_nonnilableType()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( castType.FP.Get_kind() == Ast_TypeInfoKind__Class) &&
            Lns_GetEnv().SetStackVal( castType != Ast_builtinTypeString) ).(bool)){
            self.FP.Write(Lns_getVM().String_format("%sDownCastF(", []LnsAny{typeName}))
            convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( node.FP.Get_exp().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Class) &&
                Lns_GetEnv().SetStackVal( node.FP.Get_exp().FP.Get_expType() != Ast_builtinTypeString) ).(bool)){
                self.FP.Write(".FP")
            }
            self.FP.Write(")")
        } else { 
            self.FP.Write(Lns_getVM().String_format("Lns_cast2%s( ", []LnsAny{self.FP.type2gotype(castType)}))
            convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
            self.FP.Write(")")
        }
    }
}

// 4738: decl @lune.@base.@convGo.convFilter.processExpParen
func (self *convGo_convFilter) ProcessExpParen(node *Nodes_ExpParenNode,_opt LnsAny) {
    if node.FP.Get_exp().FP.Get_expTypeList().Len() >= 2{
        self.FP.Write("Lns_car(")
        convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
        self.FP.Write(")")
        self.FP.outputAny2Type(node.FP.Get_expType())
    } else { 
        self.FP.Write("(")
        convGo_filter_1171_(node.FP.Get_exp(), self, &node.Nodes_Node)
        self.FP.Write(")")
    }
}

// 4755: decl @lune.@base.@convGo.convFilter.processExpSetVal
func (self *convGo_convFilter) ProcessExpSetVal(node *Nodes_ExpSetValNode,_opt LnsAny) {
    convGo_filter_1171_(node.FP.Get_exp1(), self, &node.Nodes_Node)
    if convGo_getExpListKind_1321_(node.FP.Get_exp1().FP.Get_expTypeList(), node.FP.Get_exp2()) == convGo_ExpListKind__Direct{
        {
            var _from26850 LnsInt = node.FP.Get_exp1().FP.Get_expTypeList().Len() + 1
            var _to26850 LnsInt = node.FP.Get_exp2().FP.Get_expTypeList().Len()
            for _work26850 := _from26850; _work26850 <= _to26850; _work26850++ {
                self.FP.Write(",_")
            }
        }
    }
    self.FP.Write(" = ")
    self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), node.FP.Get_exp2()), node.FP.Get_exp1().FP.Get_expTypeList(), node.FP.Get_exp2())
    self.FP.Writeln("")
}

// 4770: decl @lune.@base.@convGo.convFilter.processExpSetItem
func (self *convGo_convFilter) ProcessExpSetItem(node *Nodes_ExpSetItemNode,_opt LnsAny) {
    convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
    self.FP.Write(".Set(")
    switch _exp26958 := node.FP.Get_index().(type) {
    case *Nodes_IndexVal__NodeIdx:
    index := _exp26958.Val1
        convGo_filter_1171_(index, self, &node.Nodes_Node)
    case *Nodes_IndexVal__SymIdx:
    index := _exp26958.Val1
        self.FP.Write(Lns_getVM().String_format("\"%s\"", []LnsAny{index}))
    }
    self.FP.Write(",")
    convGo_filter_1171_(node.FP.Get_exp2(), self, &node.Nodes_Node)
    self.FP.Write(")")
}

// 4788: decl @lune.@base.@convGo.convFilter.processAndOr
func (self *convGo_convFilter) processAndOr(node *Nodes_ExpOp2Node,opTxt string,parent *Nodes_Node) {
    var getEnvTxt string
    getEnvTxt = self.env.FP.GetEnv(node.FP.IsThreading())
    var firstFlag bool
    firstFlag = Lns_op_not(convFilter_processAndOr__isAndOr_1999_(parent))
    if firstFlag{
        self.FP.Writeln(Lns_getVM().String_format("%s.PopVal( %s.IncStack() ||", []LnsAny{getEnvTxt, getEnvTxt}))
        self.FP.PushIndent(nil)
    }
    var opCC string
    if opTxt == "and"{
        opCC = "&&"
        
    } else { 
        opCC = "||"
        
    }
    if convFilter_processAndOr__isAndOr_1999_(node.FP.Get_exp1()){
        convGo_filter_1171_(node.FP.Get_exp1(), self, &node.Nodes_Node)
    } else { 
        self.FP.Write(Lns_getVM().String_format("%s.SetStackVal( ", []LnsAny{getEnvTxt}))
        convGo_filter_1171_(node.FP.Get_exp1(), self, &node.Nodes_Node)
        self.FP.Write(") ")
    }
    self.FP.Writeln(opCC)
    if convFilter_processAndOr__isAndOr_1999_(node.FP.Get_exp2()){
        convGo_filter_1171_(node.FP.Get_exp2(), self, &node.Nodes_Node)
    } else { 
        self.FP.Write(Lns_getVM().String_format("%s.SetStackVal( ", []LnsAny{getEnvTxt}))
        convGo_filter_1171_(node.FP.Get_exp2(), self, &node.Nodes_Node)
        self.FP.Write(") ")
    }
    if firstFlag{
        self.FP.Write(")")
        self.FP.outputAny2Type(node.FP.Get_expType())
        self.FP.PopIndent()
    }
}

// 4849: decl @lune.@base.@convGo.convFilter.processExpOp2
func (self *convGo_convFilter) ProcessExpOp2(node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt.(*ConvGo_Opt)
    var opTxt string
    opTxt = node.FP.Get_op().Txt
    if _switch27512 := opTxt; _switch27512 == "and" || _switch27512 == "or" {
        self.FP.processAndOr(node, opTxt, opt.Parent)
    } else if _switch27512 == ".." {
        convGo_filter_1171_(node.FP.Get_exp1(), self, &node.Nodes_Node)
        self.FP.Write(" + ")
        convGo_filter_1171_(node.FP.Get_exp2(), self, &node.Nodes_Node)
    } else {
        {
            __exp := Ast_bitBinOpMap.Items[opTxt]
            if __exp != nil {
                _exp := __exp.(LnsInt)
                if _switch27399 := _exp; _switch27399 == Ast_BitOpKind__LShift {
                    opTxt = "<<"
                    
                } else if _switch27399 == Ast_BitOpKind__RShift {
                    opTxt = ">>"
                    
                }
                convGo_filter_1171_(node.FP.Get_exp1(), self, &node.Nodes_Node)
                self.FP.Write(" " + opTxt + " ")
                convGo_filter_1171_(node.FP.Get_exp2(), self, &node.Nodes_Node)
            } else {
                convGo_filter_1171_(node.FP.Get_exp1(), self, &node.Nodes_Node)
                {
                    _op := convGo_op2map.Items[opTxt]
                    if _op != nil {
                        op := _op.(string)
                        self.FP.Write(Lns_getVM().String_format(" %s ", []LnsAny{op}))
                    } else {
                        self.FP.Write(Lns_getVM().String_format(" %s ", []LnsAny{opTxt}))
                    }
                }
                convGo_filter_1171_(node.FP.Get_exp2(), self, &node.Nodes_Node)
            }
        }
    }
}

// 4894: decl @lune.@base.@convGo.convFilter.processExpRef
func (self *convGo_convFilter) ProcessExpRef(node *Nodes_ExpRefNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpRef"
    if node.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__DDD{
        self.FP.Write("ddd")
    } else { 
        if Lns_isCondTrue( node.FP.Get_symbolInfo().FP.Get_convModuleParam()){
            self.FP.Write(Lns_getVM().String_format("%s_%d", []LnsAny{node.FP.Get_symbolInfo().FP.Get_name(), node.FP.Get_symbolInfo().FP.Get_symbolId()}))
        } else { 
            if _switch27723 := node.FP.Get_symbolInfo().FP.Get_kind(); _switch27723 == Ast_SymbolKind__Var || _switch27723 == Ast_SymbolKind__Arg {
                self.FP.Write(self.FP.getSymbolSym(node.FP.Get_symbolInfo()))
            } else if _switch27723 == Ast_SymbolKind__Fun {
                if Ast_isBuiltin(node.FP.Get_expType().FP.Get_typeId()){
                    var builtinFunc *TransUnit_BuiltinFuncType
                    builtinFunc = TransUnit_getBuiltinFunc()
                    if _switch27645 := node.FP.Get_expType(); _switch27645 == builtinFunc.Lns_print {
                        self.FP.Write("Lns_print")
                    } else {
                        Util_err(Lns_getVM().String_format("%s: not support -- %s", []LnsAny{__func__, node.FP.Get_symbolInfo().FP.Get_name()}))
                    }
                } else { 
                    self.FP.Write(self.FP.getSymbol(&convGo_SymbolKind__Func{node.FP.Get_expType()}, node.FP.Get_symbolInfo().FP.Get_name()))
                }
            } else if _switch27723 == Ast_SymbolKind__Typ {
                if node.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Module{
                    self.FP.Write(node.FP.Get_symbolInfo().FP.Get_name())
                } else { 
                    self.FP.Write(self.FP.getTypeSymbol(node.FP.Get_expType()))
                }
            } else {
                self.FP.Write(node.FP.Get_symbolInfo().FP.Get_name())
            }
        }
    }
}

// 4942: decl @lune.@base.@convGo.convFilter.processExpRefItem
func (self *convGo_convFilter) ProcessExpRefItem(node *Nodes_ExpRefItemNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpRefItem"
    var getEnvTxt string
    getEnvTxt = self.env.FP.GetEnv(node.FP.IsThreading())
    var prefixType *Ast_TypeInfo
    prefixType = node.FP.Get_val().FP.Get_expType().FP.Get_nonnilableType()
    if _switch28520 := prefixType.FP.Get_kind(); _switch28520 == Ast_TypeInfoKind__Ext {
        if node.FP.Get_nilAccess(){
            self.FP.Write(Lns_getVM().String_format("%s.NilAccFin( %s.NilAccPush( ", []LnsAny{getEnvTxt, getEnvTxt}))
            convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
            self.FP.Write(") && ")
            self.FP.Write(Lns_getVM().String_format("%s.NilAccPush( %s.NilAccPop().(*Lns_luaValue)", []LnsAny{getEnvTxt, getEnvTxt}))
        } else { 
            convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
            if prefixType.FP.Get_extedType().FP.Get_kind() == Ast_TypeInfoKind__Stem{
                self.FP.Write(".(*Lns_luaValue)")
            }
        }
        self.FP.Write(".GetAt(")
        {
            _index := node.FP.Get_index()
            if _index != nil {
                index := _index.(*Nodes_Node)
                convGo_filter_1171_(index, self, &node.Nodes_Node)
            } else {
                self.FP.Write(Lns_getVM().String_format("\"%s\"", []LnsAny{convGo_str2gostr_1254_(Lns_unwrap( node.FP.Get_symbol()).(string))}))
            }
        }
        self.FP.Write(")")
        self.FP.outputStem2Type(node.FP.Get_expType())
    } else if _switch28520 == Ast_TypeInfoKind__List || _switch28520 == Ast_TypeInfoKind__Array {
        if node.FP.Get_nilAccess(){
            if Lns_op_not(node.FP.Get_val().FP.HasNilAccess()){
                self.FP.Writeln(Lns_getVM().String_format("%s.NilAccFin( %s.NilAccPush(", []LnsAny{getEnvTxt, getEnvTxt}))
                convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
                self.FP.Writeln(") && ")
            } else { 
                convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
                self.FP.Writeln("&&")
            }
            self.FP.Write(Lns_getVM().String_format("%s.NilAccPush( %s.NilAccPop().(*LnsList)", []LnsAny{getEnvTxt, getEnvTxt}))
            self.FP.Write(".GetAt(")
            convGo_filter_1171_(Lns_unwrap( node.FP.Get_index()).(*Nodes_Node), self, &node.Nodes_Node)
            self.FP.Write(")")
            self.FP.outputStem2Type(node.FP.Get_expType())
            self.FP.Write("))")
        } else { 
            convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
            self.FP.Write(".GetAt(")
            convGo_filter_1171_(Lns_unwrap( node.FP.Get_index()).(*Nodes_Node), self, &node.Nodes_Node)
            self.FP.Write(")")
            self.FP.outputStem2Type(node.FP.Get_expType())
        }
    } else if _switch28520 == Ast_TypeInfoKind__Map {
        if node.FP.Get_nilAccess(){
            if Lns_op_not(node.FP.Get_val().FP.HasNilAccess()){
                self.FP.Writeln(Lns_getVM().String_format("%s.NilAccFin( %s.NilAccPush(", []LnsAny{getEnvTxt, getEnvTxt}))
                convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
                self.FP.Writeln(") && ")
            } else { 
                convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
                self.FP.Writeln("&&")
            }
            self.FP.Write(Lns_getVM().String_format("%s.NilAccPush( %s.NilAccPop().(*LnsMap)", []LnsAny{getEnvTxt, getEnvTxt}))
        } else { 
            convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
            if prefixType.FP.Get_kind() == Ast_TypeInfoKind__Stem{
                self.FP.Write(".(*LnsMap)")
            }
        }
        self.FP.Write(".Items[")
        {
            _index := node.FP.Get_index()
            if _index != nil {
                index := _index.(*Nodes_Node)
                convGo_filter_1171_(index, self, &node.Nodes_Node)
            } else {
                self.FP.Write(Lns_getVM().String_format("\"%s\"", []LnsAny{convGo_str2gostr_1254_(Lns_unwrap( node.FP.Get_symbol()).(string))}))
            }
        }
        self.FP.Write("]")
        self.FP.outputStem2Type(node.FP.Get_expType())
        if node.FP.Get_nilAccess(){
            self.FP.Write("))")
        }
    } else if _switch28520 == Ast_TypeInfoKind__Stem {
        self.FP.Write("Lns_FromStemGetAt(")
        convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
        self.FP.Write(",")
        {
            _index := node.FP.Get_index()
            if _index != nil {
                index := _index.(*Nodes_Node)
                convGo_filter_1171_(index, self, &node.Nodes_Node)
            } else {
                self.FP.Write(Lns_getVM().String_format("\"%s\"", []LnsAny{convGo_str2gostr_1254_(Lns_unwrap( node.FP.Get_symbol()).(string))}))
            }
        }
        self.FP.Write(Lns_getVM().String_format(", %s )", []LnsAny{node.FP.Get_nilAccess()}))
        self.FP.outputStem2Type(node.FP.Get_expType())
    } else {
        if prefixType == Ast_builtinTypeString{
            self.FP.Write("LnsInt(")
            convGo_filter_1171_(node.FP.Get_val(), self, &node.Nodes_Node)
            self.FP.Write("[")
            convGo_filter_1171_(Lns_unwrap( node.FP.Get_index()).(*Nodes_Node), self, &node.Nodes_Node)
            self.FP.Write("-1])")
        } else { 
            Util_err(Lns_getVM().String_format("%s: not support -- %d, %s", []LnsAny{__func__, node.FP.Get_pos().LineNo, Ast_TypeInfoKind_getTxt( prefixType.FP.Get_kind())}))
        }
    }
}

// 5053: decl @lune.@base.@convGo.convFilter.processRefField
func (self *convGo_convFilter) ProcessRefField(node *Nodes_RefFieldNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processRefField"
    opt := _opt.(*ConvGo_Opt)
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( node.FP.Get_prefix().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Enum) &&
        Lns_GetEnv().SetStackVal( node.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Enum) ).(bool)){
        self.FP.Write(self.FP.getSymbol(&convGo_SymbolKind__Static{node.FP.Get_expType()}, node.FP.Get_field().Txt))
        return 
    }
    {
        _symbol := node.FP.Get_symbolInfo()
        if _symbol != nil {
            symbol := _symbol.(*Ast_SymbolInfo)
            var orgSym *Ast_SymbolInfo
            orgSym = symbol.FP.GetOrg()
            var builtinFuncs *TransUnit_BuiltinFuncType
            builtinFuncs = TransUnit_getBuiltinFunc()
            if builtinFuncs.FP.Get_allSymbolSet().Has(Ast_SymbolInfo2Stem(orgSym)){
                self.FP.Write(Lns_getVM().String_format("Lns_%s_%s", []LnsAny{Lns_car(Lns_getVM().String_gsub(node.FP.Get_prefix().FP.Get_expType().FP.Get_rawTxt(),"@", "")).(string), orgSym.FP.Get_name()}))
                return 
            }
            if symbol.FP.Get_staticFlag(){
                self.FP.outputModule(symbol.FP.GetModule(), true)
                self.FP.Write(self.FP.getSymbolSym(symbol))
                return 
            }
        }
    }
    var getEnvTxt string
    getEnvTxt = self.env.FP.GetEnv(node.FP.IsThreading())
    var openParenNum LnsInt
    if Lns_op_not(node.FP.HasNilAccess()){
        openParenNum = 0
        
        if Lns_op_not(node.FP.Get_prefix().FP.HasNilAccess()){
            convGo_filter_1171_(node.FP.Get_prefix(), self, &node.Nodes_Node)
        } else { 
            Util_err(Lns_getVM().String_format("%s: not support", []LnsAny{__func__}))
        }
    } else { 
        if Lns_op_not(node.FP.Get_prefix().FP.HasNilAccess()){
            self.FP.Write(Lns_getVM().String_format("%s.NilAccFin(", []LnsAny{getEnvTxt}))
            self.FP.Write(Lns_getVM().String_format("%s.NilAccPush(", []LnsAny{getEnvTxt}))
            convGo_filter_1171_(node.FP.Get_prefix(), self, &node.Nodes_Node)
            self.FP.Writeln(") && ")
        } else { 
            convGo_filter_1171_(node.FP.Get_prefix(), self, &node.Nodes_Node)
            self.FP.Writeln("&&")
        }
        self.FP.Write(Lns_getVM().String_format("%s.NilAccPush(", []LnsAny{getEnvTxt}))
        if opt.Parent.FP.HasNilAccess(){
            openParenNum = 1
            
        } else { 
            openParenNum = 2
            
        }
        self.FP.Write(Lns_getVM().String_format("%s.NilAccPop().(%s)", []LnsAny{getEnvTxt, self.FP.type2gotype(node.FP.Get_prefix().FP.Get_expType().FP.Get_nonnilableType())}))
    }
    self.FP.Write(".")
    {
        _symbol := node.FP.Get_symbolInfo()
        if _symbol != nil {
            symbol := _symbol.(*Ast_SymbolInfo)
            if node.FP.Get_prefix().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Ext{
                self.FP.Write(Lns_getVM().String_format("GetAt( \"%s\" )", []LnsAny{symbol.FP.Get_name()}))
                self.FP.outputAny2Type(node.FP.Get_expType())
            } else { 
                self.FP.Write(self.FP.getSymbolSym(symbol))
                var orgSym *Ast_SymbolInfo
                orgSym = symbol.FP.GetOrg()
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( orgSym.FP.Get_kind() == Ast_SymbolKind__Mbr) &&
                    Lns_GetEnv().SetStackVal( orgSym.FP.Get_typeInfo().FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
                    Lns_GetEnv().SetStackVal( convGo_isAnyType_1174_(orgSym.FP.Get_typeInfo())) ).(bool)){
                    self.FP.outputAny2Type(node.FP.Get_expType())
                }
            }
        } else {
            Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
        }
    }
    {
        var _from29009 LnsInt = 1
        var _to29009 LnsInt = openParenNum
        for _work29009 := _from29009; _work29009 <= _to29009; _work29009++ {
            self.FP.Write(")")
        }
    }
}

// 5132: decl @lune.@base.@convGo.convFilter.processExpOmitEnum
func (self *convGo_convFilter) ProcessExpOmitEnum(node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    self.FP.Write(self.FP.getSymbol(&convGo_SymbolKind__Static{node.FP.Get_expType().FP.Get_aliasSrc()}, node.FP.Get_valInfo().FP.Get_name()))
}

// 5138: decl @lune.@base.@convGo.convFilter.processGetField
func (self *convGo_convFilter) ProcessGetField(node *Nodes_GetFieldNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processGetField"
    opt := _opt.(*ConvGo_Opt)
    {
        _symbolInfo := node.FP.Get_symbolInfo()
        if _symbolInfo != nil {
            symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mtd) &&
                Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_name() == "get__txt") ).(bool)){
                {
                    _enumType := Ast_EnumTypeInfoDownCastF(symbolInfo.FP.Get_namespaceTypeInfo().FP)
                    if _enumType != nil {
                        enumType := _enumType.(*Ast_EnumTypeInfo)
                        self.FP.Write(Lns_getVM().String_format("%s( ", []LnsAny{self.FP.getEnumGetTxtSym(enumType)}))
                        convGo_filter_1171_(node.FP.Get_prefix(), self, &node.Nodes_Node)
                        self.FP.Write(")")
                        return 
                    }
                }
                if Lns_isCondTrue( Ast_AlgeTypeInfoDownCastF(symbolInfo.FP.Get_namespaceTypeInfo().FP)){
                    convGo_filter_1171_(node.FP.Get_prefix(), self, &node.Nodes_Node)
                    self.FP.Write(".(LnsAlgeVal).GetTxt()")
                    return 
                }
            }
            if symbolInfo.FP.Get_staticFlag(){
                self.FP.Write(self.FP.getSymbolSym(symbolInfo))
                self.FP.Write("()")
            } else { 
                var closeParen bool
                closeParen = convGo_convExp29208(Lns_2DDD(self.FP.OutputCallPrefix(node.FP.IsThreading(), node.FP.Get_id(), &node.Nodes_Node, node.FP.Get_prefix(), symbolInfo)))
                self.FP.Write("()")
                var retType *Ast_TypeInfo
                retType = symbolInfo.FP.Get_typeInfo().FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( retType.FP.Get_kind() == Ast_TypeInfoKind__Alternate) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(retType.FP.HasBase())) ).(bool)){
                    self.FP.outputAny2Type(node.FP.Get_expType())
                }
                if node.FP.HasNilAccess(){
                    self.FP.Write("})")
                    if opt.Parent.FP.HasNilAccess(){
                    } else { 
                        self.FP.Write(")")
                    }
                }
                if closeParen{
                    self.FP.Write(")")
                }
            }
        } else {
            Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
        }
    }
}

// 5187: decl @lune.@base.@convGo.convFilter.processReturn
func (self *convGo_convFilter) ProcessReturn(node *Nodes_ReturnNode,_opt LnsAny) {
    self.FP.Write("return ")
    {
        _expList := node.FP.Get_expList()
        if _expList != nil {
            expList := _expList.(*Nodes_ExpListNode)
            convGo_filter_1171_(&expList.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln("")
}

// 5197: decl @lune.@base.@convGo.convFilter.processTestCase
func (self *convGo_convFilter) ProcessTestCase(node *Nodes_TestCaseNode,_opt LnsAny) {
    if Lns_op_not(self.enableTest){
        return 
    }
    self.FP.Writeln(Lns_getVM().String_format("func lns_test_%s_%s( %s *Testing_Ctrl ) {", []LnsAny{convGo_getModuleName_1198_(self.moduleTypeInfo), node.FP.Get_name().Txt, node.FP.Get_ctrlName()}))
    convGo_filter_1171_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln("}")
}

// 5211: decl @lune.@base.@convGo.convFilter.processTestBlock
func (self *convGo_convFilter) ProcessTestBlock(node *Nodes_TestBlockNode,_opt LnsAny) {
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList()
    for _, _statement := range( stmtList.Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(convGo_ignoreNodeInInnerBlockSet.Has(statement.FP.Get_kind())){
            convGo_filter_1171_(statement, self, &node.Nodes_Node)
        }
    }
}

// 5223: decl @lune.@base.@convGo.convFilter.processProvide
func (self *convGo_convFilter) ProcessProvide(node *Nodes_ProvideNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processProvide"
    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
}

// 5229: decl @lune.@base.@convGo.convFilter.processAlias
func (self *convGo_convFilter) ProcessAlias(node *Nodes_AliasNode,_opt LnsAny) {
}

// 5234: decl @lune.@base.@convGo.convFilter.processBoxing
func (self *convGo_convFilter) ProcessBoxing(node *Nodes_BoxingNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processBoxing"
    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
}

// 5240: decl @lune.@base.@convGo.convFilter.processUnboxing
func (self *convGo_convFilter) ProcessUnboxing(node *Nodes_UnboxingNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processUnboxing"
    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
}

// 5246: decl @lune.@base.@convGo.convFilter.processLiteralList
func (self *convGo_convFilter) ProcessLiteralList(node *Nodes_LiteralListNode,_opt LnsAny) {
    self.FP.Write("NewLnsList(")
    {
        _expList := node.FP.Get_expList()
        if _expList != nil {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.expList2Slice(expList, true)
        } else {
            self.FP.Write("[]LnsAny{}")
        }
    }
    self.FP.Write(")")
}

// 5259: decl @lune.@base.@convGo.convFilter.processLiteralSet
func (self *convGo_convFilter) ProcessLiteralSet(node *Nodes_LiteralSetNode,_opt LnsAny) {
    self.FP.Write("NewLnsSet(")
    {
        _expList := node.FP.Get_expList()
        if _expList != nil {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.expList2Slice(expList, true)
        } else {
            self.FP.Write("[]LnsAny{}")
        }
    }
    self.FP.Write(")")
}

// 5273: decl @lune.@base.@convGo.convFilter.processLiteralMap
func (self *convGo_convFilter) ProcessLiteralMap(node *Nodes_LiteralMapNode,_opt LnsAny) {
    var hasNilable bool
    hasNilable = false
    self.FP.Write("NewLnsMap( map[LnsAny]LnsAny{")
    for _, _pair := range( node.FP.Get_pairList().Items ) {
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( pair.FP.Get_key().FP.Get_kind() == Nodes_NodeKind_get_LiteralNil()) ||
            Lns_GetEnv().SetStackVal( pair.FP.Get_val().FP.Get_kind() == Nodes_NodeKind_get_LiteralNil()) ).(bool){
        } else { 
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( pair.FP.Get_key().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Nilable) ||
                Lns_GetEnv().SetStackVal( pair.FP.Get_val().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Nilable) ).(bool){
                hasNilable = true
                
            }
            convGo_filter_1171_(pair.FP.Get_key(), self, &node.Nodes_Node)
            self.FP.Write(":")
            convGo_filter_1171_(pair.FP.Get_val(), self, &node.Nodes_Node)
            self.FP.Write(",")
        }
    }
    self.FP.Write("})")
    if hasNilable{
        self.FP.Write(".Correct()")
    }
}

// 5301: decl @lune.@base.@convGo.convFilter.processLiteralArray
func (self *convGo_convFilter) ProcessLiteralArray(node *Nodes_LiteralArrayNode,_opt LnsAny) {
    self.FP.Write("NewLnsList(")
    {
        _expList := node.FP.Get_expList()
        if _expList != nil {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.expList2Slice(expList, true)
        } else {
            self.FP.Write("[]LnsAny{}")
        }
    }
    self.FP.Write(")")
}

// 5315: decl @lune.@base.@convGo.convFilter.processLiteralChar
func (self *convGo_convFilter) ProcessLiteralChar(node *Nodes_LiteralCharNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("%d", []LnsAny{node.FP.Get_num()}))
}

// 5321: decl @lune.@base.@convGo.convFilter.processLiteralInt
func (self *convGo_convFilter) ProcessLiteralInt(node *Nodes_LiteralIntNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 5327: decl @lune.@base.@convGo.convFilter.processLiteralReal
func (self *convGo_convFilter) ProcessLiteralReal(node *Nodes_LiteralRealNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 5333: decl @lune.@base.@convGo.convFilter.processLiteralString
func (self *convGo_convFilter) ProcessLiteralString(node *Nodes_LiteralStringNode,_opt LnsAny) {
    var txt string
    txt = node.FP.Get_token().Txt
    {
        _expList := node.FP.Get_dddParam()
        if _expList != nil {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.Write(Lns_getVM().String_format("%s.String_format(%s, ", []LnsAny{self.env.FP.GetLuavm(node.FP.IsThreading()), convGo_str2gostr_1254_(txt)}))
            self.FP.processSetFromExpList(self.FP.getConvExpName(node.FP.Get_id(), expList), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeDDD)}), expList)
            self.FP.Write(")")
        } else {
            self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{convGo_str2gostr_1254_(txt)}))
        }
    }
}

// 5348: decl @lune.@base.@convGo.convFilter.processLiteralBool
func (self *convGo_convFilter) ProcessLiteralBool(node *Nodes_LiteralBoolNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 5354: decl @lune.@base.@convGo.convFilter.processLiteralNil
func (self *convGo_convFilter) ProcessLiteralNil(node *Nodes_LiteralNilNode,_opt LnsAny) {
    self.FP.Write("nil")
}

// 5360: decl @lune.@base.@convGo.convFilter.processBreak
func (self *convGo_convFilter) ProcessBreak(node *Nodes_BreakNode,_opt LnsAny) {
    self.FP.Write("break")
    self.FP.Writeln("")
}

// 5367: decl @lune.@base.@convGo.convFilter.processLiteralSymbol
func (self *convGo_convFilter) ProcessLiteralSymbol(node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processLiteralSymbol"
    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{__func__}))
}

// 5373: decl @lune.@base.@convGo.convFilter.processLuneControl
func (self *convGo_convFilter) ProcessLuneControl(node *Nodes_LuneControlNode,_opt LnsAny) {
    switch node.FP.Get_pragma().(type) {
    case *LuneControl_Pragma__default__init:
    case *LuneControl_Pragma__default__init_old:
    case *LuneControl_Pragma__disable_mut_control:
    case *LuneControl_Pragma__ignore_symbol_:
    case *LuneControl_Pragma__load__lune_module:
    case *LuneControl_Pragma__limit_conv_code:
    case *LuneControl_Pragma__use_async:
    case *LuneControl_Pragma__run_async_pipe:
        self.FP.Writeln("go self.FP.Loop()")
    }
}

// 5397: decl @lune.@base.@convGo.convFilter.processAbbr
func (self *convGo_convFilter) ProcessAbbr(node *Nodes_AbbrNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processAbbr"
    Util_err(Lns_getVM().String_format("not support -- %s:%d", []LnsAny{__func__, node.FP.Get_pos().LineNo}))
}


// declaration Class -- FuncConv
type convGo_FuncConvMtd interface {
    Get_argList() *LnsList
    Get_retList() *LnsList
}
type convGo_FuncConv struct {
    argList *LnsList
    retList *LnsList
    FP convGo_FuncConvMtd
}
func convGo_FuncConv2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convGo_FuncConv).FP
}
type convGo_FuncConvDownCast interface {
    ToconvGo_FuncConv() *convGo_FuncConv
}
func convGo_FuncConvDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convGo_FuncConvDownCast)
    if ok { return work.ToconvGo_FuncConv() }
    return nil
}
func (obj *convGo_FuncConv) ToconvGo_FuncConv() *convGo_FuncConv {
    return obj
}
func NewconvGo_FuncConv(arg1 *LnsList) *convGo_FuncConv {
    obj := &convGo_FuncConv{}
    obj.FP = obj
    obj.InitconvGo_FuncConv(arg1)
    return obj
}
func (self *convGo_FuncConv) Get_argList() *LnsList{ return self.argList }
func (self *convGo_FuncConv) Get_retList() *LnsList{ return self.retList }
// 1293: DeclConstr
func (self *convGo_FuncConv) InitconvGo_FuncConv(retList *LnsList) {
    self.argList = NewLnsList([]LnsAny{})
    
    self.retList = retList
    
}


func Lns_convGo_init() {
    if init_convGo { return }
    init_convGo = true
    convGo__mod__ = "@lune.@base.@convGo"
    Lns_InitMod()
    Lns_Ver_init()
    Lns_Ast_init()
    Lns_Nodes_init()
    Lns_Util_init()
    Lns_TransUnit_init()
    Lns_LuaVer_init()
    Lns_Parser_init()
    Lns_LuneControl_init()
    Lns_Types_init()
    convGo_MaxNilAccNum = 3
    convGo_ignoreNodeInInnerBlockSet = NewLnsSet([]LnsAny{Nodes_NodeKind_get_DeclAlge(), Nodes_NodeKind_get_DeclEnum(), Nodes_NodeKind_get_DeclMethod(), Nodes_NodeKind_get_DeclForm(), Nodes_NodeKind_get_DeclMacro(), Nodes_NodeKind_get_TestCase()})
    convGo_golanKeywordSet = NewLnsSet([]LnsAny{"func", "select", "defer", "go", "chan", "package", "const", "fallthrough", "range", "continue", "var", "map", "default"})
    
    
    convGo_type2LnsItemKindMap = NewLnsMap( map[LnsAny]LnsAny{Ast_builtinTypeInt:"LnsItemKindInt",Ast_builtinTypeReal:"LnsItemKindReal",Ast_builtinTypeString:"LnsItemKindStr",})
    convGo_type2FromStemNameMap = NewLnsMap( map[LnsAny]LnsAny{Ast_builtinTypeInt:"Lns_ToInt",Ast_builtinTypeReal:"Lns_ToReal",Ast_builtinTypeBool:"Lns_ToBool",Ast_builtinTypeString:"Lns_ToStr",Ast_builtinTypeStem:"Lns_ToStem",})
    convGo_op2map = NewLnsMap( map[LnsAny]LnsAny{"..":"+","~=":"!=",})
    Lns_convLua_init()
}
func init() {
    init_convGo = false
}
