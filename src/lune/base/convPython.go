// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_convPython bool
var convPython__mod__ string
// decl enum -- ProcessMode 
type convPython_ProcessMode = LnsInt
const convPython_ProcessMode__DeclClass = 1
const convPython_ProcessMode__DeclTopScopeVar = 0
const convPython_ProcessMode__Main = 3
const convPython_ProcessMode__NonClosureFuncDecl = 2
var convPython_ProcessModeList_ = NewLnsList( []LnsAny {
  convPython_ProcessMode__DeclTopScopeVar,
  convPython_ProcessMode__DeclClass,
  convPython_ProcessMode__NonClosureFuncDecl,
  convPython_ProcessMode__Main,
})
func convPython_ProcessMode_get__allList_2_(_env *LnsEnv) *LnsList{
    return convPython_ProcessModeList_
}
var convPython_ProcessModeMap_ = map[LnsInt]string {
  convPython_ProcessMode__DeclClass: "ProcessMode.DeclClass",
  convPython_ProcessMode__DeclTopScopeVar: "ProcessMode.DeclTopScopeVar",
  convPython_ProcessMode__Main: "ProcessMode.Main",
  convPython_ProcessMode__NonClosureFuncDecl: "ProcessMode.NonClosureFuncDecl",
}
func convPython_ProcessMode__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := convPython_ProcessModeMap_[arg1]; ok { return arg1 }
    return nil
}

func convPython_ProcessMode_getTxt(arg1 LnsInt) string {
    return convPython_ProcessModeMap_[arg1];
}
// decl enum -- ClassAsterMode 
type convPython_ClassAsterMode = LnsInt
const convPython_ClassAsterMode__Force = 2
const convPython_ClassAsterMode__None = 0
const convPython_ClassAsterMode__Normal = 1
var convPython_ClassAsterModeList_ = NewLnsList( []LnsAny {
  convPython_ClassAsterMode__None,
  convPython_ClassAsterMode__Normal,
  convPython_ClassAsterMode__Force,
})
func convPython_ClassAsterMode_get__allList_2_(_env *LnsEnv) *LnsList{
    return convPython_ClassAsterModeList_
}
var convPython_ClassAsterModeMap_ = map[LnsInt]string {
  convPython_ClassAsterMode__Force: "ClassAsterMode.Force",
  convPython_ClassAsterMode__None: "ClassAsterMode.None",
  convPython_ClassAsterMode__Normal: "ClassAsterMode.Normal",
}
func convPython_ClassAsterMode__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := convPython_ClassAsterModeMap_[arg1]; ok { return arg1 }
    return nil
}

func convPython_ClassAsterMode_getTxt(arg1 LnsInt) string {
    return convPython_ClassAsterModeMap_[arg1];
}
// decl enum -- ExpListKind 
type convPython_ExpListKind = LnsInt
const convPython_ExpListKind__Conv = 2
const convPython_ExpListKind__Direct = 0
const convPython_ExpListKind__Slice = 1
var convPython_ExpListKindList_ = NewLnsList( []LnsAny {
  convPython_ExpListKind__Direct,
  convPython_ExpListKind__Slice,
  convPython_ExpListKind__Conv,
})
func convPython_ExpListKind_get__allList_2_(_env *LnsEnv) *LnsList{
    return convPython_ExpListKindList_
}
var convPython_ExpListKindMap_ = map[LnsInt]string {
  convPython_ExpListKind__Conv: "ExpListKind.Conv",
  convPython_ExpListKind__Direct: "ExpListKind.Direct",
  convPython_ExpListKind__Slice: "ExpListKind.Slice",
}
func convPython_ExpListKind__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := convPython_ExpListKindMap_[arg1]; ok { return arg1 }
    return nil
}

func convPython_ExpListKind_getTxt(arg1 LnsInt) string {
    return convPython_ExpListKindMap_[arg1];
}
var convPython_MaxNilAccNum LnsInt
var convPython_literalNil string
var convPython_ignoreNodeInInnerBlockSetForNoTest *LnsSet2_[LnsInt]
var convPython_ignoreNodeInInnerBlockSetForTest *LnsSet2_[LnsInt]
var convPython_pythonLangKeywordSet *LnsSet2_[string]
var convPython_type2LnsItemKindMap *LnsMap2_[*Ast_TypeInfo,string]
var convPython_type2FromStemNameMap *LnsMap2_[*Ast_TypeInfo,string]
var convPython_op2map *LnsMap2_[string,string]
// decl alge -- SymbolKind
type convPython_SymbolKind = LnsAny
type convPython_SymbolKind__Arg struct{
}
var convPython_SymbolKind__Arg_Obj = &convPython_SymbolKind__Arg{}
func (self *convPython_SymbolKind__Arg) GetTxt() string {
return "SymbolKind.Arg"
}
type convPython_SymbolKind__Func struct{
Val1 *Ast_TypeInfo
}
func (self *convPython_SymbolKind__Func) GetTxt() string {
return "SymbolKind.Func"
}
type convPython_SymbolKind__Member struct{
Val1 bool
}
func (self *convPython_SymbolKind__Member) GetTxt() string {
return "SymbolKind.Member"
}
type convPython_SymbolKind__Normal struct{
}
var convPython_SymbolKind__Normal_Obj = &convPython_SymbolKind__Normal{}
func (self *convPython_SymbolKind__Normal) GetTxt() string {
return "SymbolKind.Normal"
}
type convPython_SymbolKind__Static struct{
Val1 *Ast_TypeInfo
}
func (self *convPython_SymbolKind__Static) GetTxt() string {
return "SymbolKind.Static"
}
type convPython_SymbolKind__Type struct{
Val1 *Ast_TypeInfo
Val2 bool
}
func (self *convPython_SymbolKind__Type) GetTxt() string {
return "SymbolKind.Type"
}
type convPython_SymbolKind__Var struct{
Val1 *Ast_SymbolInfo
}
func (self *convPython_SymbolKind__Var) GetTxt() string {
return "SymbolKind.Var"
}
// decl alge -- FuncInfo
type convPython_FuncInfo = LnsAny
type convPython_FuncInfo__DeclInfo struct{
Val1 *Nodes_Node
Val2 *Nodes_DeclFuncInfo
}
func (self *convPython_FuncInfo__DeclInfo) GetTxt() string {
return "FuncInfo.DeclInfo"
}
type convPython_FuncInfo__Type struct{
Val1 *Ast_TypeInfo
}
func (self *convPython_FuncInfo__Type) GetTxt() string {
return "FuncInfo.Type"
}
type convPython_FuncInfo__WithClass struct{
Val1 *Ast_TypeInfo
Val2 *Ast_TypeInfo
}
func (self *convPython_FuncInfo__WithClass) GetTxt() string {
return "FuncInfo.WithClass"
}
// decl alge -- CallKind
type convPython_CallKind = LnsAny
type convPython_CallKind__BuiltinCall struct{
}
var convPython_CallKind__BuiltinCall_Obj = &convPython_CallKind__BuiltinCall{}
func (self *convPython_CallKind__BuiltinCall) GetTxt() string {
return "CallKind.BuiltinCall"
}
type convPython_CallKind__BuiltinCallEnv struct{
}
var convPython_CallKind__BuiltinCallEnv_Obj = &convPython_CallKind__BuiltinCallEnv{}
func (self *convPython_CallKind__BuiltinCallEnv) GetTxt() string {
return "CallKind.BuiltinCallEnv"
}
type convPython_CallKind__FormCall struct{
}
var convPython_CallKind__FormCall_Obj = &convPython_CallKind__FormCall{}
func (self *convPython_CallKind__FormCall) GetTxt() string {
return "CallKind.FormCall"
}
type convPython_CallKind__LuaCall struct{
}
var convPython_CallKind__LuaCall_Obj = &convPython_CallKind__LuaCall{}
func (self *convPython_CallKind__LuaCall) GetTxt() string {
return "CallKind.LuaCall"
}
type convPython_CallKind__Normal struct{
}
var convPython_CallKind__Normal_Obj = &convPython_CallKind__Normal{}
func (self *convPython_CallKind__Normal) GetTxt() string {
return "CallKind.Normal"
}
type convPython_CallKind__RunLoaded struct{
}
var convPython_CallKind__RunLoaded_Obj = &convPython_CallKind__RunLoaded{}
func (self *convPython_CallKind__RunLoaded) GetTxt() string {
return "CallKind.RunLoaded"
}
type convPython_CallKind__RuntimeCall struct{
Val1 *Nodes_Node
}
func (self *convPython_CallKind__RuntimeCall) GetTxt() string {
return "CallKind.RuntimeCall"
}
type convPython_CallKind__SortCall struct{
Val1 *Ast_TypeInfo
}
func (self *convPython_CallKind__SortCall) GetTxt() string {
return "CallKind.SortCall"
}
type convPython_convFilter_processRoot__ProcNode_1_ func (_env *LnsEnv, arg1 *Nodes_Node)
// for 217
func convPython_convExp0_574(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 685
func convPython_convExp1_2011(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 4970
func convPython_convExp0_4804(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 427
func convPython_convExp1_716(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 428
func convPython_convExp1_745(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 5781
func convPython_convExp0_9177(arg1 []LnsAny) bool {
    return Lns_getFromMulti( arg1, 0 ).(bool)
}
// 95: decl @lune.@base.@convPython.isMain
func convPython_isMain_3_(_env *LnsEnv, funcType *Ast_TypeInfo) bool {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Func) &&
        _env.SetStackVal( funcType.FP.Get_rawTxt(_env) == "__main") &&
        _env.SetStackVal( funcType.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) ).(bool)){
        return true
    }
    return false
}

// 297: decl @lune.@base.@convPython.getAddEnvArg
func convPython_getAddEnvArg_6_(_env *LnsEnv, argLen LnsInt,addEnvArg bool) string {
    if addEnvArg{
        var txt string
        txt = "_env"
        if argLen > 0{
            return txt + ", "
        }
        return txt
    }
    return ""
}

// 308: decl @lune.@base.@convPython.filter
func convPython_filter_7_(_env *LnsEnv, node *Nodes_Node,filter *convPython_convFilter,parent *Nodes_Node) {
    node.FP.ProcessFilter(_env, &filter.Nodes_Filter, ConvPython_Opt2Stem(NewConvPython_Opt(_env, parent)))
}

// 314: decl @lune.@base.@convPython.isAnyType
func convPython_isAnyType_8_(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    var work *Ast_TypeInfo
    work = typeInfo.FP.Get_srcTypeInfo(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_nilable(_env)) ||
        _env.SetStackVal( work == Ast_builtinTypeStem) ).(bool){
        return true
    }
    if _switch0 := typeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Stem || _switch0 == Ast_TypeInfoKind__Alge {
        return true
    } else if _switch0 == Ast_TypeInfoKind__Alternate {
        if typeInfo.FP.HasBase(_env){
            return convPython_isAnyType_8_(_env, typeInfo.FP.Get_baseTypeInfo(_env))
        }
        return true
    } else if _switch0 == Ast_TypeInfoKind__Ext {
        if typeInfo.FP.Get_extedType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Stem{
            return true
        }
    }
    return false
}

// 338: decl @lune.@base.@convPython.isClosure
func convPython_isClosure_9_(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    var scope *Ast_Scope
    
    {
        _scope := typeInfo.FP.Get_scope(_env)
        if _scope == nil{
            return false
        } else {
            scope = _scope.(*Ast_Scope)
        }
    }
    return scope.FP.Get_hasClosureAccess(_env)
}

// 373: decl @lune.@base.@convPython.concatGLSym
func convPython_concatGLSym_11_(_env *LnsEnv, name string,external bool) string {
    return name
}

// 380: decl @lune.@base.@convPython.isInnerDeclType
func convPython_isInnerDeclType_12_(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__FormFunc{
        return typeInfo.FP.Get_parentInfo(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Module
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_parentInfo(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Module) ||
        _env.SetStackVal( Ast_TypeInfo2Stem(_env.NilAccFin(_env.NilAccPush(typeInfo.FP.Get_scope(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_parent(_env)})&&
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.Get_ownerTypeInfo(_env)}))) == nil) ).(bool){
        return true
    }
    return false
}

// 456: decl @lune.@base.@convPython.normalizeSym
func convPython_normalizeSym_13_(_env *LnsEnv, name string) string {
    if convPython_pythonLangKeywordSet.Has(name){
        return _env.GetVM().String_format("_%s", Lns_2DDD(name))
    }
    return name
}

// 682: decl @lune.@base.@convPython.str2gostr
func convPython_str2gostr_14_(_env *LnsEnv, txt string) string {
    var work string
    work = txt
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(work, "^```", nil, nil))){
        work = convPython_convExp1_2011(Lns_2DDD(_env.GetVM().String_gsub(_env.GetVM().String_sub(work,4, -4),"^\n", "")))
        work = Tokenizer_quoteStr(_env, work)
    } else if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(work, "^'", nil, nil))){
        work = Str_replace(_env, Tokenizer_convFromRawToStr(_env, work), "\"", "\\\"")
        work = _env.GetVM().String_format("\"%s\"", Lns_2DDD(work))
    }
    return work
}

// 697: decl @lune.@base.@convPython.getOrgTypeInfo
func convPython_getOrgTypeInfo_15_(_env *LnsEnv, typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    {
        _enumType := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env).FP)
        if !Lns_IsNil( _enumType ) {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            return enumType.FP.Get_valTypeInfo(_env)
        }
    }
    return typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)
}


// 868: decl @lune.@base.@convPython.needConvCast
func convPython_needConvCast_19_(_env *LnsEnv, dstType *Ast_TypeInfo,srcType *Ast_TypeInfo) bool {
    if _switch1 := dstType.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Nilable {
        return convPython_needConvCast_19_(_env, dstType.FP.Get_nonnilableType(_env), srcType)
    } else if _switch1 == Ast_TypeInfoKind__Class {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( dstType == Ast_builtinTypeString) ||
            _env.SetStackVal( srcType.FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env) == dstType.FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)) ).(bool){
            return false
        } else { 
            return true
        }
    } else if _switch1 == Ast_TypeInfoKind__IF {
        return false
    } else if _switch1 == Ast_TypeInfoKind__FormFunc {
        return true
    } else if _switch1 == Ast_TypeInfoKind__Alternate {
        if Lns_op_not(dstType.FP.HasBase(_env)){
            return false
        } else { 
            return convPython_needConvCast_19_(_env, dstType.FP.Get_baseTypeInfo(_env), srcType)
        }
    } else if _switch1 == Ast_TypeInfoKind__Form {
        return true
    } else if _switch1 == Ast_TypeInfoKind__Prim {
        if Lns_op_not(dstType.FP.Get_nilable(_env)){
            if _switch0 := dstType; _switch0 == Ast_builtinTypeInt {
                return true
            } else if _switch0 == Ast_builtinTypeReal {
                return true
            } else {
                return false
            }
        } else { 
            return false
        }
    } else {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( srcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
            _env.SetStackVal( srcType != Ast_builtinTypeString) ).(bool)){
            return true
        } else { 
            return false
        }
    }
// insert a dummy
    return false
}

// 1042: decl @lune.@base.@convPython.getExpListKind
func convPython_getExpListKind_21_(_env *LnsEnv, dstTypeList *LnsList2_[*Ast_TypeInfo],node *Nodes_ExpListNode,addEnvArg bool) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( addEnvArg) &&
        _env.SetStackVal( node.FP.Get_mRetExp(_env)) )){
        return convPython_ExpListKind__Conv
    }
    if dstTypeList.Len() < node.FP.Get_expList(_env).Len(){
        if dstTypeList.GetAt(dstTypeList.Len()).FP.Get_kind(_env) != Ast_TypeInfoKind__DDD{
            return convPython_ExpListKind__Conv
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( dstTypeList.Len() > 1) &&
        _env.SetStackVal( node.FP.Get_mRetExp(_env)) )){
        for _, _exp := range( node.FP.Get_expList(_env).Items ) {
            exp := _exp
            {
                _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
                if !Lns_IsNil( _castNode ) {
                    castNode := _castNode.(*Nodes_ExpCastNode)
                    if convPython_needConvCast_19_(_env, castNode.FP.Get_castType(_env), castNode.FP.Get_exp(_env).FP.Get_expType(_env)){
                        return convPython_ExpListKind__Conv
                    }
                }
            }
        }
    }
    var lastExp *Nodes_Node
    lastExp = node.FP.Get_expList(_env).GetAt(node.FP.Get_expList(_env).Len())
    var hasAbbr bool
    if lastExp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
        hasAbbr = true
        if node.FP.Get_expList(_env).Len() < 2{
            return convPython_ExpListKind__Direct
        }
        lastExp = node.FP.Get_expList(_env).GetAt(node.FP.Get_expList(_env).Len() - 1)
    } else { 
        hasAbbr = false
    }
    if Lns_isCondTrue( Nodes_ExpToDDDNodeDownCastF(lastExp.FP)){
        var mRetExp *Nodes_MRetExp
        
        {
            _mRetExp := node.FP.Get_mRetExp(_env)
            if _mRetExp == nil{
                return convPython_ExpListKind__Slice
            } else {
                mRetExp = _mRetExp.(*Nodes_MRetExp)
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( mRetExp.FP.Get_index(_env) == 1) &&
            _env.SetStackVal( dstTypeList.GetAt(mRetExp.FP.Get_index(_env)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
            return convPython_ExpListKind__Slice
        }
        return convPython_ExpListKind__Conv
    }
    if lastExp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        var mRetExp *Nodes_MRetExp
        
        {
            _mRetExp := node.FP.Get_mRetExp(_env)
            if _mRetExp == nil{
                return convPython_ExpListKind__Slice
            } else {
                mRetExp = _mRetExp.(*Nodes_MRetExp)
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( mRetExp.FP.Get_index(_env) == 1) &&
            _env.SetStackVal( dstTypeList.GetAt(mRetExp.FP.Get_index(_env)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
            return convPython_ExpListKind__Direct
        }
    } else { 
        var mRetExp *Nodes_MRetExp
        
        {
            _mRetExp := node.FP.Get_mRetExp(_env)
            if _mRetExp == nil{
                return convPython_ExpListKind__Direct
            } else {
                mRetExp = _mRetExp.(*Nodes_MRetExp)
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(hasAbbr)) &&
            _env.SetStackVal( mRetExp.FP.Get_index(_env) == 1) ).(bool)){
            return convPython_ExpListKind__Direct
        }
    }
    return convPython_ExpListKind__Conv
}


// 1339: decl @lune.@base.@convPython.isRetGenerics
func convPython_isRetGenerics_22_(_env *LnsEnv, node *Nodes_ExpCallNode) bool {
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_func(_env).FP.Get_expType(_env)
    for _index, _retType := range( funcType.FP.Get_retTypeInfoList(_env).Items ) {
        index := _index + 1
        retType := _retType
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( retType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
            _env.SetStackVal( Lns_op_not(convPython_isAnyType_8_(_env, node.FP.Get_expTypeList(_env).GetAt(index)))) ).(bool)){
            return true
        }
    }
    return false
}




















// 3754: decl @lune.@base.@convPython.getLnsItemKind
func convPython_getLnsItemKind_28_(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    {
        _kind := convPython_type2LnsItemKindMap.Get(typeInfo)
        if !Lns_IsNil( _kind ) {
            kind := _kind.(string)
            return kind
        }
    }
    return "LnsItemKindStem"
}





// 5373: decl @lune.@base.@convPython.convFilter.processAndOr.isAndOr
func convPython_convFilter_processAndOr__isAndOr_0_(_env *LnsEnv, exp *Nodes_Node) bool {
    {
        _parentNode := Nodes_ExpOp2NodeDownCastF(exp.FP)
        if !Lns_IsNil( _parentNode ) {
            parentNode := _parentNode.(*Nodes_ExpOp2Node)
            if _switch0 := parentNode.FP.Get_op(_env).Txt; _switch0 == "and" || _switch0 == "or" {
                return true
            }
        }
    }
    return false
}

// 6033: decl @lune.@base.@convPython.createFilter
func ConvPython_createFilter(_env *LnsEnv, enableTest bool,streamName string,stream Lns_oStream,ast *AstInfo_ASTInfo,option *ConvPython_Option) *Nodes_Filter {
    return &NewconvPython_convFilter(_env, enableTest, streamName, stream, ast, option).Nodes_Filter
}
















// declaration Class -- Opt
type ConvPython_OptMtd interface {
}
type ConvPython_Opt struct {
    Parent *Nodes_Node
    FP ConvPython_OptMtd
}
func ConvPython_Opt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvPython_Opt).FP
}
func ConvPython_Opt_toSlice(slice []LnsAny) []*ConvPython_Opt {
    ret := make([]*ConvPython_Opt, len(slice))
    for index, val := range slice {
        ret[index] = val.(ConvPython_OptDownCast).ToConvPython_Opt()
    }
    return ret
}
type ConvPython_OptDownCast interface {
    ToConvPython_Opt() *ConvPython_Opt
}
func ConvPython_OptDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvPython_OptDownCast)
    if ok { return work.ToConvPython_Opt() }
    return nil
}
func (obj *ConvPython_Opt) ToConvPython_Opt() *ConvPython_Opt {
    return obj
}
func NewConvPython_Opt(_env *LnsEnv, arg1 *Nodes_Node) *ConvPython_Opt {
    obj := &ConvPython_Opt{}
    obj.FP = obj
    obj.InitConvPython_Opt(_env, arg1)
    return obj
}
// 51: DeclConstr
func (self *ConvPython_Opt) InitConvPython_Opt(_env *LnsEnv, parent *Nodes_Node) {
    self.Parent = parent
}


// declaration Class -- Env
type convPython_EnvMtd interface {
    getCommonVm(_env *LnsEnv) string
    getEnv(_env *LnsEnv) string
    getLuavm(_env *LnsEnv) string
}
type convPython_Env struct {
    addEnvArg bool
    FP convPython_EnvMtd
}
func convPython_Env2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convPython_Env).FP
}
func convPython_Env_toSlice(slice []LnsAny) []*convPython_Env {
    ret := make([]*convPython_Env, len(slice))
    for index, val := range slice {
        ret[index] = val.(convPython_EnvDownCast).ToconvPython_Env()
    }
    return ret
}
type convPython_EnvDownCast interface {
    ToconvPython_Env() *convPython_Env
}
func convPython_EnvDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convPython_EnvDownCast)
    if ok { return work.ToconvPython_Env() }
    return nil
}
func (obj *convPython_Env) ToconvPython_Env() *convPython_Env {
    return obj
}
func NewconvPython_Env(_env *LnsEnv, arg1 bool) *convPython_Env {
    obj := &convPython_Env{}
    obj.FP = obj
    obj.InitconvPython_Env(_env, arg1)
    return obj
}
// 70: DeclConstr
func (self *convPython_Env) InitconvPython_Env(_env *LnsEnv, addEnvArg bool) {
    self.addEnvArg = addEnvArg
}


// declaration Class -- Option
type ConvPython_OptionMtd interface {
    Get_addEnvArg(_env *LnsEnv) bool
    Get_appName(_env *LnsEnv) string
    Get_mainModule(_env *LnsEnv) string
    Get_packageName(_env *LnsEnv) string
    Get_runnerNum(_env *LnsEnv) LnsInt
}
type ConvPython_Option struct {
    packageName string
    appName string
    mainModule string
    addEnvArg bool
    runnerNum LnsInt
    FP ConvPython_OptionMtd
}
func ConvPython_Option2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvPython_Option).FP
}
func ConvPython_Option_toSlice(slice []LnsAny) []*ConvPython_Option {
    ret := make([]*ConvPython_Option, len(slice))
    for index, val := range slice {
        ret[index] = val.(ConvPython_OptionDownCast).ToConvPython_Option()
    }
    return ret
}
type ConvPython_OptionDownCast interface {
    ToConvPython_Option() *ConvPython_Option
}
func ConvPython_OptionDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvPython_OptionDownCast)
    if ok { return work.ToConvPython_Option() }
    return nil
}
func (obj *ConvPython_Option) ToConvPython_Option() *ConvPython_Option {
    return obj
}
func NewConvPython_Option(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 bool, arg5 LnsInt) *ConvPython_Option {
    obj := &ConvPython_Option{}
    obj.FP = obj
    obj.InitConvPython_Option(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvPython_Option) InitConvPython_Option(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 bool, arg5 LnsInt) {
    self.packageName = arg1
    self.appName = arg2
    self.mainModule = arg3
    self.addEnvArg = arg4
    self.runnerNum = arg5
}
func (self *ConvPython_Option) Get_packageName(_env *LnsEnv) string{ return self.packageName }
func (self *ConvPython_Option) Get_appName(_env *LnsEnv) string{ return self.appName }
func (self *ConvPython_Option) Get_mainModule(_env *LnsEnv) string{ return self.mainModule }
func (self *ConvPython_Option) Get_addEnvArg(_env *LnsEnv) bool{ return self.addEnvArg }
func (self *ConvPython_Option) Get_runnerNum(_env *LnsEnv) LnsInt{ return self.runnerNum }

// declaration Class -- convFilter
type convPython_convFilterMtd interface {
    CheckExpOp3(_env *LnsEnv, arg1 *Nodes_ExpOp2Node, arg2 *ConvPython_Opt)
    concatSymWithType(_env *LnsEnv, arg1 string, arg2 *Ast_TypeInfo) string
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    expList2Slice(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 bool)
    getAccessorSym(_env *LnsEnv, arg1 LnsInt, arg2 string) string
    getAlgeSymbol(_env *LnsEnv, arg1 *Ast_AlgeValInfo) string
    getCanonicalName(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    getConstrSymbol(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    getConv2formName(_env *LnsEnv, arg1 *Nodes_Node) string
    getConvExpName(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_ExpListNode) string
    getConvGenericsName(_env *LnsEnv, arg1 *Nodes_Node) string
    getEnumGetTxtSym(_env *LnsEnv, arg1 *Ast_EnumTypeInfo) string
    getEnvArgDecl(_env *LnsEnv, arg1 LnsInt) string
    getFromStemName(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    GetFull(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    getFuncSymbol(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    getModuleName(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    getModuleSym(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    getSymbol(_env *LnsEnv, arg1 LnsAny, arg2 string) string
    getSymbolSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) string
    getTypeSymbol(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    getTypeSymbolWithPrefix(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    getVM(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny
    Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_stepIndent(_env *LnsEnv) LnsInt
    Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl
    IsExtSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo) bool
    IsExtType(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isImplementedRunner(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isInheritAbsImmut(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsPubSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) bool
    IsPubType(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isSameModDir(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isSamePackage(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isSamePackageExtModule(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    needConvFormFunc(_env *LnsEnv, arg1 *Nodes_ExpCastNode) bool
    needPrefix(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    OutputAccessor(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputAdvertise(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputAlter2MapFunc(_env *LnsEnv, arg1 *LnsMap2_[*Ast_TypeInfo,*Ast_TypeInfo])
    outputAny2Type(_env *LnsEnv, arg1 *Ast_TypeInfo)
    outputAsyncItem(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    OutputCallPrefix(_env *LnsEnv, arg1 string, arg2 *Nodes_Node, arg3 LnsAny, arg4 *Ast_SymbolInfo)(bool, LnsAny)
    outputCastReceiver(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputClassType(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 bool)
    outputConstructor(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 bool)
    outputConvExt(_env *LnsEnv, arg1 *Nodes_Node)
    outputConvItemType(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsAny)
    outputConvItemTypeList(_env *LnsEnv, arg1 *LnsList2_[*Ast_TypeInfo], arg2 LnsAny)
    outputConvToForm(_env *LnsEnv, arg1 *Nodes_ExpCastNode)
    OutputDeclFunc(_env *LnsEnv, arg1 bool, arg2 LnsAny) *convPython_FuncConv
    outputDeclFuncArg(_env *LnsEnv, arg1 *Ast_TypeInfo)
    OutputDeclFuncInfo(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo)
    outputDownCast(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputDummyAbstractMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo)
    outputDummyAbstractMethodOfClass(_env *LnsEnv, arg1 *Ast_TypeInfo)
    outputDummyReturn(_env *LnsEnv, arg1 *LnsList2_[*Ast_TypeInfo])
    outputForeachLua(_env *LnsEnv, arg1 *Nodes_Node, arg2 bool, arg3 *Nodes_Node, arg4 *Ast_SymbolInfo, arg5 LnsAny, arg6 *Nodes_BlockNode)
    outputIFMethods(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputImplicitCast(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Nodes_Node, arg3 *Nodes_ExpCastNode)
    outputImport(_env *LnsEnv, arg1 *Nodes_ImportNode)
    outputInterfaceType(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputLetVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode)
    outputMapping(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 bool)
    outputMethodIF(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputModule(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool)
    outputModuleImport(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputNewSetup(_env *LnsEnv, arg1 string, arg2 *Ast_TypeInfo, arg3 bool)
    outputNilAccCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode)
    outputRetType(_env *LnsEnv, arg1 *LnsList2_[*Ast_TypeInfo])
    outputStaticMember(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputStem2Type(_env *LnsEnv, arg1 *Ast_TypeInfo)
    outputSymbol(_env *LnsEnv, arg1 LnsAny, arg2 string)
    outputToStem(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 bool)
    outputTopScopeVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode)
    PopIndent(_env *LnsEnv)
    popProcessMode(_env *LnsEnv)
    ProcessAbbr(_env *LnsEnv, arg1 *Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(_env *LnsEnv, arg1 *Nodes_AliasNode, arg2 LnsAny)
    processAndOr(_env *LnsEnv, arg1 *Nodes_ExpOp2Node, arg2 string, arg3 *Nodes_Node)
    ProcessApply(_env *LnsEnv, arg1 *Nodes_ApplyNode, arg2 LnsAny)
    ProcessAsyncLock(_env *LnsEnv, arg1 *Nodes_AsyncLockNode, arg2 LnsAny)
    ProcessBlankLine(_env *LnsEnv, arg1 *Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(_env *LnsEnv, arg1 *Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(_env *LnsEnv, arg1 *Nodes_BreakNode, arg2 LnsAny)
    processCond(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_Node)
    ProcessCondRet(_env *LnsEnv, arg1 *Nodes_CondRetNode, arg2 LnsAny)
    ProcessCondRetList(_env *LnsEnv, arg1 *Nodes_CondRetListNode, arg2 LnsAny)
    processConvExp(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsList2_[*Ast_TypeInfo], arg3 LnsAny, arg4 bool)
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
    ProcessDeclTuple(_env *LnsEnv, arg1 *Nodes_DeclTupleNode, arg2 LnsAny)
    ProcessDeclVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode, arg2 LnsAny)
    ProcessExpAccessMRet(_env *LnsEnv, arg1 *Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(_env *LnsEnv, arg1 *Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(_env *LnsEnv, arg1 *Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(_env *LnsEnv, arg1 *Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpExpandTuple(_env *LnsEnv, arg1 *Nodes_ExpExpandTupleNode, arg2 LnsAny)
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
    processGenericsCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode)
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
    processMethodAsync(_env *LnsEnv, arg1 *LnsList) *LnsList2_[*convPython_ConvRunner]
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
    processSetFromExpList(_env *LnsEnv, arg1 string, arg2 *LnsList2_[*Ast_TypeInfo], arg3 *Nodes_ExpListNode, arg4 bool)
    ProcessShebang(_env *LnsEnv, arg1 *Nodes_ShebangNode, arg2 LnsAny)
    ProcessStmtExp(_env *LnsEnv, arg1 *Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(_env *LnsEnv, arg1 *Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(_env *LnsEnv, arg1 *Nodes_SwitchNode, arg2 LnsAny)
    ProcessTestBlock(_env *LnsEnv, arg1 *Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(_env *LnsEnv, arg1 *Nodes_TestCaseNode, arg2 LnsAny)
    ProcessTupleConst(_env *LnsEnv, arg1 *Nodes_TupleConstNode, arg2 LnsAny)
    ProcessUnboxing(_env *LnsEnv, arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(_env *LnsEnv, arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(_env *LnsEnv, arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(_env *LnsEnv, arg1 *Nodes_WhileNode, arg2 LnsAny)
    PushIndent(_env *LnsEnv, arg1 LnsAny)
    pushProcessMode(_env *LnsEnv, arg1 LnsInt)
    ReturnToSource(_env *LnsEnv)
    Setup(_env *LnsEnv)
    SwitchToHeader(_env *LnsEnv)
    type2gotype(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    type2gotypeOrg(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt) string
    Write(_env *LnsEnv, arg1 string)
    WriteRaw(_env *LnsEnv, arg1 string)
    Writeln(_env *LnsEnv, arg1 string)
}
type convPython_convFilter struct {
    Nodes_Filter
    moduleTypeInfo *Ast_TypeInfo
    moduleScope *Ast_Scope
    builtin2code *LnsMap2_[*Ast_SymbolInfo,string]
    enableTest bool
    noneNode *Nodes_NoneNode
    option *ConvPython_Option
    modDir string
    env *convPython_Env
    builtinFuncs *Builtin_BuiltinFuncType
    ignoreNodeInInnerBlockSet *LnsSet2_[LnsInt]
    ast *AstInfo_ASTInfo
    processMode LnsInt
    builtin2runtime *LnsMap2_[*Ast_TypeInfo,string]
    builtin2runtimeEnv *LnsMap2_[*Ast_TypeInfo,string]
    type2gotypeMap *LnsMap2_[*Ast_TypeInfo,string]
    stream *Util_SimpleSourceOStream
    orgStream Lns_oStream
    streamName string
    moduleType2SymbolMap *LnsMap2_[*Ast_TypeInfo,*Ast_SymbolInfo]
    module2PackSym *LnsMap2_[*Ast_TypeInfo,string]
    nodeManager *Nodes_NodeManager
    processInfo *Ast_ProcessInfo
    processModeStack *LnsList2_[LnsInt]
    FP convPython_convFilterMtd
}
func convPython_convFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convPython_convFilter).FP
}
func convPython_convFilter_toSlice(slice []LnsAny) []*convPython_convFilter {
    ret := make([]*convPython_convFilter, len(slice))
    for index, val := range slice {
        ret[index] = val.(convPython_convFilterDownCast).ToconvPython_convFilter()
    }
    return ret
}
type convPython_convFilterDownCast interface {
    ToconvPython_convFilter() *convPython_convFilter
}
func convPython_convFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convPython_convFilterDownCast)
    if ok { return work.ToconvPython_convFilter() }
    return nil
}
func (obj *convPython_convFilter) ToconvPython_convFilter() *convPython_convFilter {
    return obj
}
func NewconvPython_convFilter(_env *LnsEnv, arg1 bool, arg2 string, arg3 Lns_oStream, arg4 *AstInfo_ASTInfo, arg5 *ConvPython_Option) *convPython_convFilter {
    obj := &convPython_convFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvPython_convFilter(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
// advertise -- 136
func (self *convPython_convFilter) Get_stepIndent(_env *LnsEnv) LnsInt {
    return self.stream. FP.Get_stepIndent( _env)
}
// advertise -- 136
func (self *convPython_convFilter) PopIndent(_env *LnsEnv) {
self.stream. FP.PopIndent( _env)
}
// advertise -- 136
func (self *convPython_convFilter) PushIndent(_env *LnsEnv, arg1 LnsAny) {
self.stream. FP.PushIndent( _env, arg1)
}
// advertise -- 136
func (self *convPython_convFilter) ReturnToSource(_env *LnsEnv) {
self.stream. FP.ReturnToSource( _env)
}
// advertise -- 136
func (self *convPython_convFilter) SwitchToHeader(_env *LnsEnv) {
self.stream. FP.SwitchToHeader( _env)
}
// advertise -- 136
func (self *convPython_convFilter) Write(_env *LnsEnv, arg1 string) {
self.stream. FP.Write( _env, arg1)
}
// advertise -- 136
func (self *convPython_convFilter) WriteRaw(_env *LnsEnv, arg1 string) {
self.stream. FP.WriteRaw( _env, arg1)
}
// advertise -- 136
func (self *convPython_convFilter) Writeln(_env *LnsEnv, arg1 string) {
self.stream. FP.Writeln( _env, arg1)
}
// 180: DeclConstr
func (self *convPython_convFilter) InitconvPython_convFilter(_env *LnsEnv, enableTest bool,streamName string,stream Lns_oStream,ast *AstInfo_ASTInfo,option *ConvPython_Option) {
    self.InitNodes_Filter(_env, true, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env))
    var ignoreNodeInInnerBlockSet *LnsSet2_[LnsInt]
    if enableTest{
        ignoreNodeInInnerBlockSet = convPython_ignoreNodeInInnerBlockSetForTest
    } else { 
        ignoreNodeInInnerBlockSet = convPython_ignoreNodeInInnerBlockSetForNoTest
    }
    self.ignoreNodeInInnerBlockSet = ignoreNodeInInnerBlockSet
    self.streamName = streamName
    self.orgStream = stream
    self.ast = ast
    self.builtinFuncs = ast.FP.Get_builtinFunc(_env)
    self.moduleType2SymbolMap = NewLnsMap2_[*Ast_TypeInfo,*Ast_SymbolInfo]( map[*Ast_TypeInfo]*Ast_SymbolInfo{})
    self.env = NewconvPython_Env(_env, option.FP.Get_addEnvArg(_env))
    self.option = option
    self.processInfo = ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env).FP.Clone(_env)
    self.stream = NewUtil_SimpleSourceOStream(_env, stream, nil, 4)
    self.processMode = convPython_ProcessMode__Main
    self.processModeStack = NewLnsList2_[LnsInt]([]LnsInt{})
    self.moduleTypeInfo = ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env)
    self.moduleScope = Lns_unwrap( ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env)).(*Ast_Scope)
    self.builtin2runtime = NewLnsMap2_[*Ast_TypeInfo,string]( map[*Ast_TypeInfo]string{})
    self.builtin2runtimeEnv = NewLnsMap2_[*Ast_TypeInfo,string]( map[*Ast_TypeInfo]string{})
    self.type2gotypeMap = NewLnsMap2_[*Ast_TypeInfo,string]( map[*Ast_TypeInfo]string{})
    self.nodeManager = NewNodes_NodeManager(_env)
    self.enableTest = enableTest
    self.module2PackSym = NewLnsMap2_[*Ast_TypeInfo,string]( map[*Ast_TypeInfo]string{})
    var modDir string
    modDir = self.moduleTypeInfo.FP.GetParentFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), false)
    self.modDir = convPython_convExp0_574(Lns_2DDD(_env.GetVM().String_gsub(Str_replace(_env, modDir, "@", ""),"%.$", "")))
    self.noneNode = Nodes_NoneNode_create(_env, self.nodeManager, Tokenizer_noneToken.Pos, false, false, NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeNone)))
    self.builtin2code = NewLnsMap2_[*Ast_SymbolInfo,string]( map[*Ast_SymbolInfo]string{self.builtinFuncs.G__lns_runmode_Sync_sym:_env.GetVM().String_format("%d", Lns_2DDD(0)),self.builtinFuncs.G__lns_runmode_Queue_sym:_env.GetVM().String_format("%d", Lns_2DDD(1)),self.builtinFuncs.G__lns_runmode_Skip_sym:_env.GetVM().String_format("%d", Lns_2DDD(2)),self.builtinFuncs.G__lns_capability_async_sym:"true",self.builtinFuncs.Io_stdout_sym:"sys.stdout",})
}



// declaration Class -- FuncConv
type convPython_FuncConvMtd interface {
    Get_argList(_env *LnsEnv) *LnsList2_[*Ast_SymbolInfo]
    Get_retList(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]
}
type convPython_FuncConv struct {
    argList *LnsList2_[*Ast_SymbolInfo]
    retList *LnsList2_[*Ast_TypeInfo]
    FP convPython_FuncConvMtd
}
func convPython_FuncConv2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convPython_FuncConv).FP
}
func convPython_FuncConv_toSlice(slice []LnsAny) []*convPython_FuncConv {
    ret := make([]*convPython_FuncConv, len(slice))
    for index, val := range slice {
        ret[index] = val.(convPython_FuncConvDownCast).ToconvPython_FuncConv()
    }
    return ret
}
type convPython_FuncConvDownCast interface {
    ToconvPython_FuncConv() *convPython_FuncConv
}
func convPython_FuncConvDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convPython_FuncConvDownCast)
    if ok { return work.ToconvPython_FuncConv() }
    return nil
}
func (obj *convPython_FuncConv) ToconvPython_FuncConv() *convPython_FuncConv {
    return obj
}
func NewconvPython_FuncConv(_env *LnsEnv, arg1 *LnsList2_[*Ast_TypeInfo]) *convPython_FuncConv {
    obj := &convPython_FuncConv{}
    obj.FP = obj
    obj.InitconvPython_FuncConv(_env, arg1)
    return obj
}
func (self *convPython_FuncConv) Get_argList(_env *LnsEnv) *LnsList2_[*Ast_SymbolInfo]{ return self.argList }
func (self *convPython_FuncConv) Get_retList(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]{ return self.retList }
// 1400: DeclConstr
func (self *convPython_FuncConv) InitconvPython_FuncConv(_env *LnsEnv, retList *LnsList2_[*Ast_TypeInfo]) {
    self.argList = NewLnsList2_[*Ast_SymbolInfo]([]*Ast_SymbolInfo{})
    self.retList = retList
}


// declaration Class -- ProcessDeclMethodItem
type convPython_ProcessDeclMethodItemMtd interface {
    Get_classNode(_env *LnsEnv) *Nodes_DeclClassNode
    Get_fieldNode(_env *LnsEnv) *Nodes_DeclMethodNode
}
type convPython_ProcessDeclMethodItem struct {
    classNode *Nodes_DeclClassNode
    fieldNode *Nodes_DeclMethodNode
    FP convPython_ProcessDeclMethodItemMtd
}
func convPython_ProcessDeclMethodItem2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convPython_ProcessDeclMethodItem).FP
}
func convPython_ProcessDeclMethodItem_toSlice(slice []LnsAny) []*convPython_ProcessDeclMethodItem {
    ret := make([]*convPython_ProcessDeclMethodItem, len(slice))
    for index, val := range slice {
        ret[index] = val.(convPython_ProcessDeclMethodItemDownCast).ToconvPython_ProcessDeclMethodItem()
    }
    return ret
}
type convPython_ProcessDeclMethodItemDownCast interface {
    ToconvPython_ProcessDeclMethodItem() *convPython_ProcessDeclMethodItem
}
func convPython_ProcessDeclMethodItemDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convPython_ProcessDeclMethodItemDownCast)
    if ok { return work.ToconvPython_ProcessDeclMethodItem() }
    return nil
}
func (obj *convPython_ProcessDeclMethodItem) ToconvPython_ProcessDeclMethodItem() *convPython_ProcessDeclMethodItem {
    return obj
}
func NewconvPython_ProcessDeclMethodItem(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 *Nodes_DeclMethodNode) *convPython_ProcessDeclMethodItem {
    obj := &convPython_ProcessDeclMethodItem{}
    obj.FP = obj
    obj.InitconvPython_ProcessDeclMethodItem(_env, arg1, arg2)
    return obj
}
func (self *convPython_ProcessDeclMethodItem) InitconvPython_ProcessDeclMethodItem(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 *Nodes_DeclMethodNode) {
    self.classNode = arg1
    self.fieldNode = arg2
}
func (self *convPython_ProcessDeclMethodItem) Get_classNode(_env *LnsEnv) *Nodes_DeclClassNode{ return self.classNode }
func (self *convPython_ProcessDeclMethodItem) Get_fieldNode(_env *LnsEnv) *Nodes_DeclMethodNode{ return self.fieldNode }

// declaration Class -- ConvRunner
type convPython_ConvRunnerMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    CheckExpOp3(_env *LnsEnv, arg1 *Nodes_ExpOp2Node, arg2 *ConvPython_Opt)
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    GetFull(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    GetResult(_env *LnsEnv) string
    Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_stepIndent(_env *LnsEnv) LnsInt
    Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl
    IsExtSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo) bool
    IsExtType(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsPubSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) bool
    IsPubType(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    OutputAccessor(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    OutputCallPrefix(_env *LnsEnv, arg1 string, arg2 *Nodes_Node, arg3 LnsAny, arg4 *Ast_SymbolInfo)(bool, LnsAny)
    OutputDeclFunc(_env *LnsEnv, arg1 bool, arg2 LnsAny) *convPython_FuncConv
    OutputDeclFuncInfo(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo)
    PopIndent(_env *LnsEnv)
    popProcessMode(_env *LnsEnv)
    ProcessAbbr(_env *LnsEnv, arg1 *Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(_env *LnsEnv, arg1 *Nodes_AliasNode, arg2 LnsAny)
    ProcessApply(_env *LnsEnv, arg1 *Nodes_ApplyNode, arg2 LnsAny)
    ProcessAsyncLock(_env *LnsEnv, arg1 *Nodes_AsyncLockNode, arg2 LnsAny)
    ProcessBlankLine(_env *LnsEnv, arg1 *Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(_env *LnsEnv, arg1 *Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(_env *LnsEnv, arg1 *Nodes_BreakNode, arg2 LnsAny)
    ProcessCondRet(_env *LnsEnv, arg1 *Nodes_CondRetNode, arg2 LnsAny)
    ProcessCondRetList(_env *LnsEnv, arg1 *Nodes_CondRetListNode, arg2 LnsAny)
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
    ProcessDeclTuple(_env *LnsEnv, arg1 *Nodes_DeclTupleNode, arg2 LnsAny)
    ProcessDeclVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode, arg2 LnsAny)
    ProcessExpAccessMRet(_env *LnsEnv, arg1 *Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 LnsAny)
    ProcessExpCallSuper(_env *LnsEnv, arg1 *Nodes_ExpCallSuperNode, arg2 LnsAny)
    ProcessExpCallSuperCtor(_env *LnsEnv, arg1 *Nodes_ExpCallSuperCtorNode, arg2 LnsAny)
    ProcessExpCast(_env *LnsEnv, arg1 *Nodes_ExpCastNode, arg2 LnsAny)
    ProcessExpExpandTuple(_env *LnsEnv, arg1 *Nodes_ExpExpandTupleNode, arg2 LnsAny)
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
    ProcessTupleConst(_env *LnsEnv, arg1 *Nodes_TupleConstNode, arg2 LnsAny)
    ProcessUnboxing(_env *LnsEnv, arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(_env *LnsEnv, arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessWhen(_env *LnsEnv, arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(_env *LnsEnv, arg1 *Nodes_WhileNode, arg2 LnsAny)
    PushIndent(_env *LnsEnv, arg1 LnsAny)
    pushProcessMode(_env *LnsEnv, arg1 LnsInt)
    ReturnToSource(_env *LnsEnv)
    Run(_env *LnsEnv)
    Setup(_env *LnsEnv)
    SwitchToHeader(_env *LnsEnv)
    Write(_env *LnsEnv, arg1 string)
    WriteRaw(_env *LnsEnv, arg1 string)
    Writeln(_env *LnsEnv, arg1 string)
}
type convPython_ConvRunner struct {
    convPython_convFilter
    _syncFlag *Lns_syncFlag
    declMethodItemList *LnsList2_[*convPython_ProcessDeclMethodItem]
    FP convPython_ConvRunnerMtd
}
func (self *convPython_ConvRunner) GetLnsSyncFlag() *Lns_syncFlag { return self._syncFlag }
func convPython_ConvRunner2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convPython_ConvRunner).FP
}
func convPython_ConvRunner_toSlice(slice []LnsAny) []*convPython_ConvRunner {
    ret := make([]*convPython_ConvRunner, len(slice))
    for index, val := range slice {
        ret[index] = val.(convPython_ConvRunnerDownCast).ToconvPython_ConvRunner()
    }
    return ret
}
type convPython_ConvRunnerDownCast interface {
    ToconvPython_ConvRunner() *convPython_ConvRunner
}
func convPython_ConvRunnerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convPython_ConvRunnerDownCast)
    if ok { return work.ToconvPython_ConvRunner() }
    return nil
}
func (obj *convPython_ConvRunner) ToconvPython_ConvRunner() *convPython_ConvRunner {
    return obj
}
func NewconvPython_ConvRunner(_env *LnsEnv, arg1 bool, arg2 *AstInfo_ASTInfo, arg3 *ConvPython_Option, arg4 *LnsList2_[*convPython_ProcessDeclMethodItem]) *convPython_ConvRunner {
    obj := &convPython_ConvRunner{}
    obj.FP = obj
    obj.convPython_convFilter.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvPython_ConvRunner(_env, arg1, arg2, arg3, arg4)
    return obj
}
// 1813: DeclConstr
func (self *convPython_ConvRunner) InitconvPython_ConvRunner(_env *LnsEnv, enableTest bool,ast *AstInfo_ASTInfo,option *ConvPython_Option,declMethodItemList *LnsList2_[*convPython_ProcessDeclMethodItem]) {
    self._syncFlag = &Lns_syncFlag{}
    self.InitconvPython_convFilter(_env, enableTest, "", NewUtil_memStream(_env), ast, option)
    self.declMethodItemList = declMethodItemList
}


func Lns_convPython_init(_env *LnsEnv) {
    if init_convPython { return }
    init_convPython = true
    convPython__mod__ = "@lune.@base.@convPython"
    Lns_InitMod()
    Lns_Ver_init(_env)
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Util_init(_env)
    Lns_AstInfo_init(_env)
    Lns_LuaVer_init(_env)
    Lns_Tokenizer_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Types_init(_env)
    Lns_Option_init(_env)
    Lns_Builtin_init(_env)
    Lns_Str_init(_env)
    convPython_MaxNilAccNum = 3
    convPython_literalNil = "None"
    convPython_ignoreNodeInInnerBlockSetForNoTest = NewLnsSet2_[LnsInt](Lns_2DDDGen[LnsInt](Nodes_NodeKind_get_DeclAlge(_env), Nodes_NodeKind_get_DeclEnum(_env), Nodes_NodeKind_get_DeclMethod(_env), Nodes_NodeKind_get_DeclForm(_env), Nodes_NodeKind_get_DeclMacro(_env), Nodes_NodeKind_get_TestBlock(_env), Nodes_NodeKind_get_TestCase(_env)))
    convPython_ignoreNodeInInnerBlockSetForTest = NewLnsSet2_[LnsInt](Lns_2DDDGen[LnsInt](Nodes_NodeKind_get_DeclAlge(_env), Nodes_NodeKind_get_DeclEnum(_env), Nodes_NodeKind_get_DeclMethod(_env), Nodes_NodeKind_get_DeclForm(_env), Nodes_NodeKind_get_DeclMacro(_env), Nodes_NodeKind_get_TestCase(_env)))
    convPython_pythonLangKeywordSet = NewLnsSet2_[string](Lns_2DDDGen[string]("func", "select", "defer", "go", "chan", "package", "const", "fallthrough", "range", "continue", "var", "map", "default"))
    
    
    convPython_type2LnsItemKindMap = NewLnsMap2_[*Ast_TypeInfo,string]( map[*Ast_TypeInfo]string{Ast_builtinTypeInt:"LnsItemKindInt",Ast_builtinTypeReal:"LnsItemKindReal",Ast_builtinTypeString:"LnsItemKindStr",})
    convPython_type2FromStemNameMap = NewLnsMap2_[*Ast_TypeInfo,string]( map[*Ast_TypeInfo]string{Ast_builtinTypeInt:"Lns_ToInt",Ast_builtinTypeReal:"Lns_ToReal",Ast_builtinTypeBool:"Lns_ToBool",Ast_builtinTypeString:"Lns_ToStr",Ast_builtinTypeStem:"Lns_ToStem",})
    convPython_op2map = NewLnsMap2_[string,string]( map[string]string{"..":"+","~=":"!=",})
}
func init() {
    init_convPython = false
}
// 73: decl @lune.@base.@convPython.Env.getCommonVm
func (self *convPython_Env) getCommonVm(_env *LnsEnv) string {
    if self.addEnvArg{
        return "_env.GetVM()"
    }
    return "Lns_getVM()"
}
// 80: decl @lune.@base.@convPython.Env.getLuavm
func (self *convPython_Env) getLuavm(_env *LnsEnv) string {
    if self.addEnvArg{
        return "_env.GetVM()"
    }
    return "Lns_getVM()"
}
// 87: decl @lune.@base.@convPython.Env.getEnv
func (self *convPython_Env) getEnv(_env *LnsEnv) string {
    if self.addEnvArg{
        return "_env"
    }
    return "Lns_GetEnv()"
}
// 234: decl @lune.@base.@convPython.convFilter.getVM
func (self *convPython_convFilter) getVM(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny {
    var txt string
    
    {
        _txt := self.builtin2runtime.Get(typeInfo)
        if _txt == nil{
            return nil
        } else {
            txt = _txt.(string)
        }
    }
    var vmTxt string
    if typeInfo.FP.Get_asyncMode(_env) == Ast_Async__Noasync{
        vmTxt = self.env.FP.getCommonVm(_env)
    } else { 
        vmTxt = self.env.FP.getLuavm(_env)
    }
    return (Str_replace(_env, txt, "GETVM", vmTxt))
}
// 247: decl @lune.@base.@convPython.convFilter.pushProcessMode
func (self *convPython_convFilter) pushProcessMode(_env *LnsEnv, mode LnsInt) {
    self.processModeStack.Insert(self.processMode)
    self.processMode = mode
}
// 251: decl @lune.@base.@convPython.convFilter.popProcessMode
func (self *convPython_convFilter) popProcessMode(_env *LnsEnv) {
    self.processMode = self.processModeStack.GetAt(self.processModeStack.Len())
    self.processModeStack.Remove(nil)
}
// 256: decl @lune.@base.@convPython.convFilter.isPubType
func (self *convPython_convFilter) IsPubType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if Ast_isPubToExternal(_env, typeInfo.FP.Get_accessMode(_env)){
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Func{
            if _switch0 := typeInfo.FP.Get_parentInfo(_env).FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Class || _switch0 == Ast_TypeInfoKind__Enum {
                return self.FP.IsPubType(_env, typeInfo.FP.Get_parentInfo(_env))
            }
        }
        return true
    }
    return Lns_op_not(typeInfo.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil))
}
// 269: decl @lune.@base.@convPython.convFilter.isPubSym
func (self *convPython_convFilter) IsPubSym(_env *LnsEnv, symbol *Ast_SymbolInfo) bool {
    if Ast_isPubToExternal(_env, symbol.FP.Get_accessMode(_env)){
        return true
    }
    return Lns_op_not(symbol.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil))
}
// 276: decl @lune.@base.@convPython.convFilter.isExtType
func (self *convPython_convFilter) IsExtType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    return Lns_op_not(typeInfo.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil))
}
// 279: decl @lune.@base.@convPython.convFilter.isExtSymbol
func (self *convPython_convFilter) IsExtSymbol(_env *LnsEnv, symbol *Ast_SymbolInfo) bool {
    return Lns_op_not(symbol.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil))
}
// 286: decl @lune.@base.@convPython.convFilter.getEnvArgDecl
func (self *convPython_convFilter) getEnvArgDecl(_env *LnsEnv, argLen LnsInt) string {
    if self.option.FP.Get_addEnvArg(_env){
        var txt string
        txt = "_env *LnsEnv"
        if argLen > 0{
            return txt + ", "
        }
        return txt
    }
    return ""
}
// 393: decl @lune.@base.@convPython.convFilter.isInheritAbsImmut
func (self *convPython_convFilter) isInheritAbsImmut(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    return typeInfo.FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeAbsImmut, nil)
}
// 397: decl @lune.@base.@convPython.convFilter.getCanonicalName
func (self *convPython_convFilter) getCanonicalName(_env *LnsEnv, typeInfo *Ast_TypeInfo,localFlag bool) string {
    var enumName string
    enumName = typeInfo.FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), localFlag)
    return _env.GetVM().String_format("%s", Lns_2DDD((Str_replace(_env, enumName, "&", ""))))
}
// 404: decl @lune.@base.@convPython.convFilter.getModuleName
func (self *convPython_convFilter) getModuleName(_env *LnsEnv, typeInfo *Ast_TypeInfo,assign bool) string {
    if Lns_op_not(Ast_TypeInfo_hasParent(_env, typeInfo)){
        return ""
    }
    var moduleType *Ast_TypeInfo
    moduleType = typeInfo.FP.GetModule(_env)
    if assign{
        {
            _symbolInfo := self.moduleType2SymbolMap.Get(moduleType)
            if !Lns_IsNil( _symbolInfo ) {
                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                return symbolInfo.FP.Get_name(_env)
            }
        }
    }
    return (Str_replace(_env, moduleType.FP.Get_rawTxt(_env), "@", ""))
}
// 417: decl @lune.@base.@convPython.convFilter.concatSymWithType
func (self *convPython_convFilter) concatSymWithType(_env *LnsEnv, name string,typeInfo *Ast_TypeInfo) string {
    return convPython_concatGLSym_11_(_env, name, self.FP.IsPubType(_env, typeInfo))
}
// 421: decl @lune.@base.@convPython.convFilter.isSamePackageExtModule
func (self *convPython_convFilter) isSamePackageExtModule(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    var extModuleType *Ast_NormalTypeInfo
    
    {
        _extModuleType := Ast_NormalTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env).FP)
        if _extModuleType == nil{
            Util_err(_env, _env.GetVM().String_format("illegal type -- %s", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            extModuleType = _extModuleType.(*Ast_NormalTypeInfo)
        }
    }
    var requireParent string
    requireParent = convPython_convExp1_716(Lns_2DDD(_env.GetVM().String_gsub(extModuleType.FP.Get_requirePath(_env),"[^%.]+$", "")))
    var moduleParent string
    moduleParent = convPython_convExp1_745(Lns_2DDD(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(self.FP.GetFull(_env, self.moduleTypeInfo, false),"[^@%.]+$", "")).(string),"@", "")))
    return requireParent == moduleParent
}
// 433: decl @lune.@base.@convPython.convFilter.isSameModDir
func (self *convPython_convFilter) isSameModDir(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo) bool {
    if moduleTypeInfo.FP.Get_parentInfo(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo.FP.Get_parentInfo(_env), nil, nil){
        return true
    }
    return false
}
// 442: decl @lune.@base.@convPython.convFilter.isSamePackage
func (self *convPython_convFilter) isSamePackage(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__ExtModule{
        return self.FP.isSamePackageExtModule(_env, typeInfo)
    }
    return self.FP.isSameModDir(_env, typeInfo)
}
// 449: decl @lune.@base.@convPython.convFilter.needPrefix
func (self *convPython_convFilter) needPrefix(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if Ast_isBuiltin(_env, typeInfo.FP.Get_typeId(_env).Id){
        return false
    }
    return Lns_op_not(self.FP.isSamePackage(_env, typeInfo))
}
// 463: decl @lune.@base.@convPython.convFilter.getSymbol
func (self *convPython_convFilter) getSymbol(_env *LnsEnv, kind LnsAny,name string) string {
    __func__ := "@lune.@base.@convPython.convFilter.getSymbol"
    var symbolName string
    symbolName = convPython_normalizeSym_13_(_env, name)
    switch _matchExp0 := kind.(type) {
    case *convPython_SymbolKind__Arg:
        return symbolName
    case *convPython_SymbolKind__Var:
        symbolInfo := _matchExp0.Val1
        var modName string
        modName = Str_replace(_env, self.moduleTypeInfo.FP.Get_rawTxt(_env), "@", "")
        if Lns_op_not(symbolInfo.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil)){
            symbolName = _env.GetVM().String_format("%s_%s", Lns_2DDD(self.FP.getModuleName(_env, symbolInfo.FP.GetModule(_env), false), symbolInfo.FP.Get_name(_env)))
        } else if name == "__mod__"{
            symbolName = name
        } else if symbolInfo.FP.Get_scope(_env) == self.moduleScope{
            symbolName = convPython_concatGLSym_11_(_env, _env.GetVM().String_format("%s_", Lns_2DDD(modName)) + symbolName, Ast_isPubToExternal(_env, symbolInfo.FP.Get_accessMode(_env)))
        } else if Lns_op_not(symbolInfo.FP.Get_posForModToRef(_env)){
            if symbolName != "__func__"{
                symbolName = "_"
            }
        }
        if self.FP.needPrefix(_env, symbolInfo.FP.GetModule(_env)){
            symbolName = _env.GetVM().String_format("%s.%s", Lns_2DDD(self.FP.getModuleName(_env, symbolInfo.FP.GetModule(_env), true), symbolName))
        }
    case *convPython_SymbolKind__Member:
        external := _matchExp0.Val1
        symbolName = convPython_concatGLSym_11_(_env, symbolName, external)
    case *convPython_SymbolKind__Func:
        typeInfo := _matchExp0.Val1
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
            if _switch0 := symbolName; _switch0 == "_toMap" {
                return "ToMap"
            } else {
                symbolName = convPython_concatGLSym_11_(_env, symbolName, Ast_isPubToExternal(_env, typeInfo.FP.Get_accessMode(_env)))
            }
        } else { 
            var prefix LnsAny
            prefix = nil
            if _switch1 := typeInfo.FP.Get_parentInfo(_env).FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Module || _switch1 == Ast_TypeInfoKind__Func || _switch1 == Ast_TypeInfoKind__Method {
                if convPython_isInnerDeclType_12_(_env, typeInfo){
                    if Lns_op_not(convPython_isClosure_9_(_env, typeInfo)){
                        var parentName string
                        parentName = typeInfo.FP.GetParentFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), true)
                        symbolName = _env.GetVM().String_format("%s_%s_%d_", Lns_2DDD(Lns_car(_env.GetVM().String_gsub(parentName,"[%.@]", "_")).(string), symbolName, typeInfo.FP.Get_childId(_env)))
                    }
                } else { 
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(Ast_isBuiltin(_env, typeInfo.FP.Get_typeId(_env).Id))) &&
                        _env.SetStackVal( Lns_op_not(self.FP.IsPubType(_env, typeInfo))) ).(bool)){
                        symbolName = _env.GetVM().String_format("%s_%d_", Lns_2DDD(symbolName, typeInfo.FP.Get_childId(_env)))
                    }
                    symbolName = self.FP.concatSymWithType(_env, symbolName, typeInfo)
                }
            } else if _switch1 == Ast_TypeInfoKind__Enum || _switch1 == Ast_TypeInfoKind__Class {
                var parentInfo *Ast_TypeInfo
                parentInfo = typeInfo.FP.Get_parentInfo(_env)
                symbolName = _env.GetVM().String_format("%s_%s", Lns_2DDD(self.FP.getSymbol(_env, &convPython_SymbolKind__Type{parentInfo, false}, parentInfo.FP.Get_rawTxt(_env)), symbolName))
                if Lns_op_not(self.FP.IsPubType(_env, typeInfo)){
                    symbolName = _env.GetVM().String_format("%s_%d_", Lns_2DDD(symbolName, typeInfo.FP.Get_childId(_env)))
                }
            } else if _switch1 == Ast_TypeInfoKind__ExtModule {
                symbolName = convPython_concatGLSym_11_(_env, symbolName, true)
                if Lns_op_not(self.FP.isSamePackageExtModule(_env, typeInfo.FP.Get_parentInfo(_env))){
                    prefix = typeInfo.FP.Get_parentInfo(_env).FP.Get_rawTxt(_env)
                }
            } else {
                Util_err(_env, _env.GetVM().String_format("%s: not support -- %s:%s", Lns_2DDD(__func__, typeInfo.FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), true), Ast_TypeInfoKind_getTxt( typeInfo.FP.Get_parentInfo(_env).FP.Get_kind(_env)))))
            }
            if Lns_op_not(prefix){
                if self.FP.needPrefix(_env, typeInfo.FP.GetModule(_env)){
                    prefix = self.FP.getModuleName(_env, typeInfo, true)
                }
            }
            if prefix != nil{
                prefix_110 := prefix.(string)
                symbolName = _env.GetVM().String_format("%s.%s", Lns_2DDD(prefix_110, symbolName))
            }
        }
    case *convPython_SymbolKind__Type:
        typeInfo := _matchExp0.Val1
        needPrefix := _matchExp0.Val2
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__FormFunc{
            return self.FP.getSymbol(_env, &convPython_SymbolKind__Func{typeInfo}, symbolName)
        }
        var workName string
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( convPython_isInnerDeclType_12_(_env, typeInfo)) &&
            _env.SetStackVal( Lns_op_not(Ast_isBuiltin(_env, typeInfo.FP.Get_typeId(_env).Id))) ).(bool)){
            workName = _env.GetVM().String_format("%s%d", Lns_2DDD(name, typeInfo.FP.Get_typeId(_env).Id))
        } else { 
            workName = symbolName
        }
        symbolName = self.FP.concatSymWithType(_env, workName, typeInfo)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( needPrefix) &&
            _env.SetStackVal( self.FP.needPrefix(_env, typeInfo.FP.GetModule(_env))) ).(bool)){
            symbolName = _env.GetVM().String_format("%s.%s", Lns_2DDD(self.FP.getModuleName(_env, typeInfo, true), symbolName))
        }
    case *convPython_SymbolKind__Static:
        typeInfo := _matchExp0.Val1
        var workName string
        workName = self.FP.getSymbol(_env, &convPython_SymbolKind__Type{typeInfo, true}, typeInfo.FP.Get_rawTxt(_env))
        symbolName = _env.GetVM().String_format("%s__%s", Lns_2DDD(workName, name))
    case *convPython_SymbolKind__Normal:
    }
    return symbolName
}
// 588: decl @lune.@base.@convPython.convFilter.getTypeSymbol
func (self *convPython_convFilter) getTypeSymbol(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    var orgType *Ast_TypeInfo
    orgType = typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env)
    return self.FP.getSymbol(_env, &convPython_SymbolKind__Type{orgType, false}, orgType.FP.Get_rawTxt(_env))
}
// 598: decl @lune.@base.@convPython.convFilter.getTypeSymbolWithPrefix
func (self *convPython_convFilter) getTypeSymbolWithPrefix(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    var orgType *Ast_TypeInfo
    orgType = typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env)
    return self.FP.getSymbol(_env, &convPython_SymbolKind__Type{orgType, true}, orgType.FP.Get_rawTxt(_env))
}
// 604: decl @lune.@base.@convPython.convFilter.getConstrSymbol
func (self *convPython_convFilter) getConstrSymbol(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    return _env.GetVM().String_format("Init%s", Lns_2DDD(self.FP.getTypeSymbol(_env, typeInfo)))
}
// 608: decl @lune.@base.@convPython.convFilter.getFuncSymbol
func (self *convPython_convFilter) getFuncSymbol(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) &&
        _env.SetStackVal( typeInfo.FP.Get_staticFlag(_env)) ).(bool)){
        return self.FP.getSymbol(_env, &convPython_SymbolKind__Static{typeInfo.FP.Get_parentInfo(_env)}, typeInfo.FP.Get_rawTxt(_env))
    }
    return self.FP.getSymbol(_env, &convPython_SymbolKind__Func{typeInfo}, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_rawTxt(_env) == "") &&
        _env.SetStackVal( "_anonymous") ||
        _env.SetStackVal( typeInfo.FP.Get_rawTxt(_env)) ).(string))
}
// 617: decl @lune.@base.@convPython.convFilter.getAlgeSymbol
func (self *convPython_convFilter) getAlgeSymbol(_env *LnsEnv, valInfo *Ast_AlgeValInfo) string {
    return self.FP.getSymbol(_env, &convPython_SymbolKind__Static{&valInfo.FP.Get_algeTpye(_env).Ast_TypeInfo}, valInfo.FP.Get_name(_env))
}
// 621: decl @lune.@base.@convPython.convFilter.getSymbolSym
func (self *convPython_convFilter) getSymbolSym(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) string {
    if _switch0 := symbolInfo.FP.Get_kind(_env); _switch0 == Ast_SymbolKind__Fun || _switch0 == Ast_SymbolKind__Mtd {
        return self.FP.getFuncSymbol(_env, symbolInfo.FP.Get_typeInfo(_env))
    } else if _switch0 == Ast_SymbolKind__Arg {
        return self.FP.getSymbol(_env, convPython_SymbolKind__Arg_Obj, symbolInfo.FP.Get_name(_env))
    } else if _switch0 == Ast_SymbolKind__Var {
        return self.FP.getSymbol(_env, &convPython_SymbolKind__Var{symbolInfo}, symbolInfo.FP.Get_name(_env))
    } else if _switch0 == Ast_SymbolKind__Mbr {
        if symbolInfo.FP.Get_staticFlag(_env){
            return self.FP.getSymbol(_env, &convPython_SymbolKind__Static{symbolInfo.FP.Get_namespaceTypeInfo(_env)}, symbolInfo.FP.Get_name(_env))
        }
        return self.FP.getSymbol(_env, &convPython_SymbolKind__Member{Ast_isPubToExternal(_env, symbolInfo.FP.Get_accessMode(_env))}, symbolInfo.FP.Get_name(_env))
    } else if _switch0 == Ast_SymbolKind__Typ {
        if Lns_isCondTrue( Ast_AliasTypeInfoDownCastF(symbolInfo.FP.Get_typeInfo(_env).FP)){
            return self.FP.getSymbol(_env, &convPython_SymbolKind__Var{symbolInfo}, symbolInfo.FP.Get_name(_env))
        }
        return self.FP.getTypeSymbol(_env, symbolInfo.FP.Get_typeInfo(_env))
    } else {
        Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind(_env)))))
    }
// insert a dummy
    return ""
}
// 652: decl @lune.@base.@convPython.convFilter.getAccessorSym
func (self *convPython_convFilter) getAccessorSym(_env *LnsEnv, accessMode LnsInt,name string) string {
    return self.FP.getSymbol(_env, &convPython_SymbolKind__Member{Ast_isPubToExternal(_env, accessMode)}, name)
}
// 656: decl @lune.@base.@convPython.convFilter.outputSymbol
func (self *convPython_convFilter) outputSymbol(_env *LnsEnv, kind LnsAny,name string) {
    self.FP.WriteRaw(_env, self.FP.getSymbol(_env, kind, name))
}
// 660: decl @lune.@base.@convPython.convFilter.getConv2formName
func (self *convPython_convFilter) getConv2formName(_env *LnsEnv, node *Nodes_Node) string {
    return _env.GetVM().String_format("conv2Form%s", Lns_2DDD(node.FP.GetIdTxt(_env)))
}
// 664: decl @lune.@base.@convPython.convFilter.getConvGenericsName
func (self *convPython_convFilter) getConvGenericsName(_env *LnsEnv, node *Nodes_Node) string {
    return _env.GetVM().String_format("lns_convGenerics%s", Lns_2DDD(node.FP.GetIdTxt(_env)))
}
// 669: decl @lune.@base.@convPython.convFilter.getModuleSym
func (self *convPython_convFilter) getModuleSym(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,addDot bool) string {
    {
        _packSym := self.module2PackSym.Get(moduleTypeInfo)
        if !Lns_IsNil( _packSym ) {
            packSym := _packSym.(string)
            if addDot{
                return _env.GetVM().String_format("%s.", Lns_2DDD(packSym))
            }
            return packSym
        }
    }
    return ""
}
// 710: decl @lune.@base.@convPython.convFilter.type2gotypeOrg
func (self *convPython_convFilter) type2gotypeOrg(_env *LnsEnv, typeInfo *Ast_TypeInfo,mode LnsInt) string {
    if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        return "[]LnsAny"
    }
    if convPython_isAnyType_8_(_env, typeInfo){
        return "LnsAny"
    }
    var orgType *Ast_TypeInfo
    orgType = convPython_getOrgTypeInfo_15_(_env, typeInfo)
    {
        _goType := self.type2gotypeMap.Get(orgType)
        if !Lns_IsNil( _goType ) {
            goType := _goType.(string)
            return goType
        }
    }
    if _switch0 := orgType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Ext {
        if orgType.FP.Get_extedType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Stem{
            return "LnsAny"
        }
        return "*Lns_luaValue"
    } else if _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        return "*LnsList"
    } else if _switch0 == Ast_TypeInfoKind__Set {
        return "*LnsSet"
    } else if _switch0 == Ast_TypeInfoKind__Map {
        return "*LnsMap"
    } else if _switch0 == Ast_TypeInfoKind__Form {
        return "LnsForm"
    } else if _switch0 == Ast_TypeInfoKind__FormFunc {
        return self.FP.getFuncSymbol(_env, typeInfo)
    } else if _switch0 == Ast_TypeInfoKind__Class {
        if typeInfo.FP.Get_genSrcTypeInfo(_env) == self.builtinFuncs.G__pipe_{
            return "*Lns__pipe"
        }
        var symbol string
        symbol = self.FP.getTypeSymbolWithPrefix(_env, typeInfo)
        if mode != convPython_ClassAsterMode__None{
            if self.FP.isInheritAbsImmut(_env, typeInfo){
                if mode == convPython_ClassAsterMode__Force{
                    return "*" + symbol
                }
            } else { 
                return "*" + symbol
            }
        }
        return symbol
    } else if _switch0 == Ast_TypeInfoKind__IF {
        return self.FP.getTypeSymbolWithPrefix(_env, typeInfo)
    } else if _switch0 == Ast_TypeInfoKind__Alternate {
        return self.FP.type2gotypeOrg(_env, typeInfo.FP.Get_baseTypeInfo(_env), mode)
    }
    Util_err(_env, _env.GetVM().String_format("not support yet -- %s", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil))))
// insert a dummy
    return ""
}
// 774: decl @lune.@base.@convPython.convFilter.type2gotype
func (self *convPython_convFilter) type2gotype(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    return self.FP.type2gotypeOrg(_env, typeInfo, convPython_ClassAsterMode__Normal)
}
// 788: decl @lune.@base.@convPython.convFilter.outputAny2Type
func (self *convPython_convFilter) outputAny2Type(_env *LnsEnv, dstType *Ast_TypeInfo) {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(convPython_isAnyType_8_(_env, dstType))) &&
        _env.SetStackVal( dstType.FP.Get_kind(_env) != Ast_TypeInfoKind__Alternate) ).(bool)){
        self.FP.WriteRaw(_env, _env.GetVM().String_format(".(%s)", Lns_2DDD(self.FP.type2gotype(_env, dstType))))
    }
}
// 795: decl @lune.@base.@convPython.convFilter.outputStem2Type
func (self *convPython_convFilter) outputStem2Type(_env *LnsEnv, dstType *Ast_TypeInfo) {
}
// 821: decl @lune.@base.@convPython.convFilter.processBlankLine
func (self *convPython_convFilter) ProcessBlankLine(_env *LnsEnv, node *Nodes_BlankLineNode,_opt LnsAny) {
}
// 821: decl @lune.@base.@convPython.convFilter.processNone
func (self *convPython_convFilter) ProcessNone(_env *LnsEnv, node *Nodes_NoneNode,_opt LnsAny) {
}
// 830: decl @lune.@base.@convPython.convFilter.processImport
func (self *convPython_convFilter) ProcessImport(_env *LnsEnv, node *Nodes_ImportNode,_opt LnsAny) {
}
// 835: decl @lune.@base.@convPython.convFilter.needConvFormFunc
func (self *convPython_convFilter) needConvFormFunc(_env *LnsEnv, node *Nodes_ExpCastNode) bool {
    var castType *Ast_TypeInfo
    castType = node.FP.Get_castType(_env).FP.Get_extedType(_env).FP.Get_nonnilableType(_env)
    if castType.FP.Get_kind(_env) != Ast_TypeInfoKind__FormFunc{
        return false
    }
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_extedType(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( castType.FP.Get_argTypeInfoList(_env).Len() != funcType.FP.Get_argTypeInfoList(_env).Len()) ||
        _env.SetStackVal( castType.FP.Get_retTypeInfoList(_env).Len() != funcType.FP.Get_retTypeInfoList(_env).Len()) ).(bool){
        return true
    }
    for _index, _argType := range( castType.FP.Get_argTypeInfoList(_env).Items ) {
        index := _index + 1
        argType := _argType
        if Lns_op_not(argType.FP.Equals(_env, self.processInfo, funcType.FP.Get_argTypeInfoList(_env).GetAt(index), nil, nil)){
            return true
        }
    }
    for _index, _retType := range( castType.FP.Get_retTypeInfoList(_env).Items ) {
        index := _index + 1
        retType := _retType
        if Lns_op_not(retType.FP.Equals(_env, self.processInfo, funcType.FP.Get_retTypeInfoList(_env).GetAt(index), nil, nil)){
            return true
        }
    }
    return false
}
// 929: decl @lune.@base.@convPython.convFilter.outputImplicitCast
func (self *convPython_convFilter) outputImplicitCast(_env *LnsEnv, castType *Ast_TypeInfo,node *Nodes_Node,parent *Nodes_ExpCastNode) {
    if _switch1 := castType.FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Nilable {
        self.FP.outputImplicitCast(_env, castType.FP.Get_nonnilableType(_env), node, parent)
    } else if _switch1 == Ast_TypeInfoKind__Class {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( castType == Ast_builtinTypeString) ||
            _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) ||
            _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env) == castType.FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)) ).(bool){
            convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
        } else { 
            if convPython_isAnyType_8_(_env, node.FP.Get_expType(_env)){
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%sDownCastF(", Lns_2DDD(self.FP.getTypeSymbolWithPrefix(_env, castType))))
                convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
                self.FP.WriteRaw(_env, ")")
            } else { 
                self.FP.WriteRaw(_env, "&")
                convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
                self.FP.WriteRaw(_env, _env.GetVM().String_format(".%s", Lns_2DDD(self.FP.getTypeSymbol(_env, castType))))
            }
        }
    } else if _switch1 == Ast_TypeInfoKind__IF {
        convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
        if Ast_isClass(_env, node.FP.Get_expType(_env)){
            self.FP.WriteRaw(_env, ".FP")
        }
    } else if _switch1 == Ast_TypeInfoKind__FormFunc {
        _ = node.FP.Get_expType(_env)
        if self.FP.needConvFormFunc(_env, parent){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s(", Lns_2DDD(self.FP.getConv2formName(_env, &parent.Nodes_Node))))
            convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        } else { 
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s(", Lns_2DDD(self.FP.getTypeSymbol(_env, castType))))
            convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        }
    } else if _switch1 == Ast_TypeInfoKind__Alternate {
        if Lns_op_not(castType.FP.HasBase(_env)){
            if Ast_isClass(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env)){
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s2Stem(", Lns_2DDD(self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env)))))
                convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
                self.FP.WriteRaw(_env, ")")
            } else { 
                convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
            }
        } else { 
            self.FP.outputImplicitCast(_env, castType.FP.Get_baseTypeInfo(_env), node, parent)
        }
    } else if _switch1 == Ast_TypeInfoKind__Form {
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s(", Lns_2DDD(self.FP.getConv2formName(_env, &parent.Nodes_Node))))
        convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
        self.FP.WriteRaw(_env, ")")
    } else if _switch1 == Ast_TypeInfoKind__Prim {
        if Lns_op_not(node.FP.Get_expType(_env).FP.Get_nilable(_env)){
            if _switch0 := castType; _switch0 == Ast_builtinTypeInt {
                self.FP.WriteRaw(_env, "LnsInt(")
                convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
                self.FP.WriteRaw(_env, ")")
            } else if _switch0 == Ast_builtinTypeReal {
                self.FP.WriteRaw(_env, "LnsReal(")
                convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
                self.FP.WriteRaw(_env, ")")
            } else {
                convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
            }
        } else { 
            convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
        }
    } else {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_nilable(_env)) &&
            _env.SetStackVal( Ast_isClass(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env))) ).(bool)){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s2Stem(", Lns_2DDD(self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env)))))
            convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        } else { 
            convPython_filter_7_(_env, node, self, &parent.Nodes_Node)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
                _env.SetStackVal( node.FP.Get_expType(_env) != Ast_builtinTypeString) ).(bool)){
                self.FP.WriteRaw(_env, ".FP")
            }
        }
    }
}
// 1115: decl @lune.@base.@convPython.convFilter.getConvExpName
func (self *convPython_convFilter) getConvExpName(_env *LnsEnv, node *Nodes_Node,argListNode *Nodes_ExpListNode) string {
    return _env.GetVM().String_format("%s_convExp%s", Lns_2DDD(Str_replace(_env, self.moduleTypeInfo.FP.Get_rawTxt(_env), "@", ""), node.FP.GetIdTxt(_env)))
}
// 1121: decl @lune.@base.@convPython.convFilter.processConvExp
func (self *convPython_convFilter) processConvExp(_env *LnsEnv, node *Nodes_Node,dstTypeList *LnsList2_[*Ast_TypeInfo],argListNode LnsAny,hasRetEnv bool) {
    var argList *Nodes_ExpListNode
    
    {
        _argList := argListNode
        if _argList == nil{
            return 
        } else {
            argList = _argList.(*Nodes_ExpListNode)
        }
    }
    if convPython_getExpListKind_21_(_env, dstTypeList, argList, self.option.FP.Get_addEnvArg(_env)) != convPython_ExpListKind__Conv{
        return 
    }
    var expList *LnsList2_[*Nodes_Node]
    expList = argList.FP.Get_expList(_env)
    var mRetIndex LnsAny
    mRetIndex = _env.NilAccFin(_env.NilAccPush(argList.FP.Get_mRetExp(_env)) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
    if Lns_op_not(mRetIndex){
        if expList.GetAt(expList.Len()).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
        } else { 
            return 
        }
    }
    var workList *LnsList2_[*Nodes_Node]
    workList = NewLnsList2_[*Nodes_Node]([]*Nodes_Node{})
    for _, _exp := range( expList.Items ) {
        exp := _exp
        var workExp *Nodes_Node
        workExp = Nodes_getUnwraped(_env, exp)
        if workExp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
            break
        }
        workList.Insert(workExp)
    }
    expList = workList
    self.FP.Writeln(_env, _env.GetVM().String_format("# for %d", Lns_2DDD(argList.FP.Get_pos(_env).LineNo)))
    self.FP.WriteRaw(_env, _env.GetVM().String_format("func %s(", Lns_2DDD(self.FP.getConvExpName(_env, node, argList))))
    if hasRetEnv{
        self.FP.WriteRaw(_env, self.FP.getEnvArgDecl(_env, expList.Len()))
    }
    for _index, _argExp := range( expList.Items ) {
        index := _index + 1
        argExp := _argExp
        {
            _exp2ddd := Nodes_ExpToDDDNodeDownCastF(argExp.FP)
            if !Lns_IsNil( _exp2ddd ) {
                exp2ddd := _exp2ddd.(*Nodes_ExpToDDDNode)
                for _, _exp := range( exp2ddd.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
                    exp := _exp
                    if index != 1{
                        self.FP.WriteRaw(_env, ", ")
                    }
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d ", Lns_2DDD(index)))
                    self.FP.WriteRaw(_env, self.FP.type2gotype(_env, exp.FP.Get_expType(_env)))
                }
            } else {
                if index != 1{
                    self.FP.WriteRaw(_env, ", ")
                }
                if mRetIndex == index{
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d []LnsAny", Lns_2DDD(index)))
                    break
                } else { 
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d ", Lns_2DDD(index)))
                    {
                        _castNode := Nodes_ExpCastNodeDownCastF(argExp.FP)
                        if !Lns_IsNil( _castNode ) {
                            castNode := _castNode.(*Nodes_ExpCastNode)
                            self.FP.WriteRaw(_env, self.FP.type2gotype(_env, castNode.FP.Get_castType(_env)))
                        } else {
                            self.FP.WriteRaw(_env, self.FP.type2gotype(_env, argExp.FP.Get_expType(_env)))
                        }
                    }
                }
            }
        }
    }
    self.FP.WriteRaw(_env, ") ")
    var convPython_getRetType func(_env *LnsEnv, retType *Ast_TypeInfo,index LnsInt) *Ast_TypeInfo
    convPython_getRetType = func(_env *LnsEnv, retType *Ast_TypeInfo,index LnsInt) *Ast_TypeInfo {
        if retType == Ast_builtinTypeEmpty{
            return argList.FP.GetExpTypeNoDDDAt(_env, index)
        }
        return retType
    }
    var retTypeList *LnsList2_[*Ast_TypeInfo]
    retTypeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    for _index, _dstType := range( dstTypeList.Items ) {
        index := _index + 1
        dstType := _dstType
        retTypeList.Insert(convPython_getRetType(_env, dstType, index))
    }
    var needRetEnv bool
    needRetEnv = false
    if retTypeList.Len() >= 2{
        self.FP.WriteRaw(_env, "(")
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( hasRetEnv) &&
            _env.SetStackVal( self.option.FP.Get_addEnvArg(_env)) ).(bool)){
            self.FP.WriteRaw(_env, "*LnsEnv, ")
            needRetEnv = true
        }
        for _index, _retType := range( retTypeList.Items ) {
            index := _index + 1
            retType := _retType
            if index != 1{
                self.FP.WriteRaw(_env, ", ")
            }
            self.FP.WriteRaw(_env, self.FP.type2gotype(_env, convPython_getRetType(_env, retType, index)))
        }
        self.FP.Writeln(_env, ") {")
    } else if retTypeList.Len() == 1{
        self.FP.WriteRaw(_env, self.FP.type2gotype(_env, convPython_getRetType(_env, retTypeList.GetAt(1), 1)))
        self.FP.Writeln(_env, " {")
    } else { 
        self.FP.Writeln(_env, " {")
    }
    self.FP.WriteRaw(_env, "    return ")
    if needRetEnv{
        self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, retTypeList.Len(), self.option.FP.Get_addEnvArg(_env)))
    }
    if mRetIndex != nil{
        mRetIndex_335 := mRetIndex.(LnsInt)
        var restIndex LnsAny
        restIndex = nil
        for _index, _retType := range( retTypeList.Items ) {
            index := _index + 1
            retType := _retType
            if index != 1{
                self.FP.WriteRaw(_env, ", ")
            }
            if retType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                restIndex = index
                break
            }
            if index >= mRetIndex_335{
                var valTxt string
                valTxt = _env.GetVM().String_format("Lns_getFromMulti( arg%d, %d )", Lns_2DDD(mRetIndex_335, index - mRetIndex_335))
                var wrote bool
                wrote = false
                if index <= expList.Len(){
                    var exp *Nodes_Node
                    exp = expList.GetAt(index)
                    {
                        _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
                        if !Lns_IsNil( _castNode ) {
                            castNode := _castNode.(*Nodes_ExpCastNode)
                            var srcTxt string
                            if castNode.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_nilable(_env){
                                srcTxt = valTxt
                            } else { 
                                srcTxt = _env.GetVM().String_format("%s.(%s)", Lns_2DDD(valTxt, self.FP.type2gotype(_env, castNode.FP.Get_exp(_env).FP.Get_expType(_env))))
                            }
                            var statNode *Nodes_ConvStatNode
                            statNode = Nodes_ConvStatNode_create(_env, self.nodeManager, exp.FP.Get_pos(_env), false, false, NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](exp.FP.Get_expType(_env))), srcTxt)
                            self.FP.outputImplicitCast(_env, castNode.FP.Get_castType(_env), &statNode.Nodes_Node, castNode)
                            wrote = true
                        }
                    }
                }
                if Lns_op_not(wrote){
                    self.FP.WriteRaw(_env, valTxt)
                    self.FP.outputAny2Type(_env, retType)
                }
            } else { 
                self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d", Lns_2DDD(index)))
            }
        }
        if restIndex != nil{
            restIndex_356 := restIndex.(LnsInt)
            self.FP.WriteRaw(_env, "Lns_2DDD( ")
            for _index, _ := range( expList.Items ) {
                index := _index + 1
                if index >= restIndex_356{
                    if index < expList.Len(){
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d", Lns_2DDD(index)))
                    } else { 
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d[%d:]", Lns_2DDD(mRetIndex_335, index - mRetIndex_335)))
                        break
                    }
                }
            }
            self.FP.Writeln(_env, ")")
        } else {
            self.FP.Writeln(_env, "")
        }
    } else {
        for _index, _ := range( retTypeList.Items ) {
            index := _index + 1
            if index != 1{
                self.FP.WriteRaw(_env, ", ")
            }
            if index <= expList.Len(){
                self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d", Lns_2DDD(index)))
            } else { 
                self.FP.WriteRaw(_env, convPython_literalNil)
            }
        }
        self.FP.Writeln(_env, "")
    }
    self.FP.Writeln(_env, "}")
}
// 1308: decl @lune.@base.@convPython.convFilter.outputNilAccCall
func (self *convPython_convFilter) outputNilAccCall(_env *LnsEnv, node *Nodes_ExpCallNode) {
    if Lns_op_not(node.FP.HasNilAccess(_env)){
        return 
    }
    if node.FP.Get_expTypeList(_env).Len() > convPython_MaxNilAccNum{
        var anys string
        anys = "LnsAny"
        var nils string
        nils = convPython_literalNil
        var lists string
        lists = "list[0]"
        {
            var _forFrom0 LnsInt = 2
            var _forTo0 LnsInt = node.FP.Get_expTypeList(_env).Len()
            for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                count := _forWork0
                anys = _env.GetVM().String_format("%s,LnsAny", Lns_2DDD(anys))
                nils = _env.GetVM().String_format("%s,%s", Lns_2DDD(nils, convPython_literalNil))
                lists = _env.GetVM().String_format("%s,list[%d]", Lns_2DDD(lists, count - 1))
            }
        }
        var name string
        name = _env.GetVM().String_format("%s_%s", Lns_2DDD(Str_replace(_env, self.moduleTypeInfo.FP.Get_rawTxt(_env), "@", ""), node.FP.GetIdTxt(_env)))
        self.FP.WriteRaw(_env, _env.GetVM().String_format("func lns_NilAccCall_%s( env *LnsEnv, call func () (%s) ) bool {\n    return env.NilAccPush( Lns_2DDD( call() ) )\n}\nfunc lns_NilAccFinCall_%s( ret LnsAny ) (%s) {\n    if Lns_IsNil( ret ) {\n        return %s\n    }\n    list := ret.([]LnsAny)\n    return %s\n}\n", Lns_2DDD(name, anys, name, anys, nils, lists)))
    }
}
// 1350: decl @lune.@base.@convPython.convFilter.processGenericsCall
func (self *convPython_convFilter) processGenericsCall(_env *LnsEnv, node *Nodes_ExpCallNode) {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(convPython_isRetGenerics_22_(_env, node))) ||
        _env.SetStackVal( node.FP.Get_expTypeList(_env).Len() < 2) ).(bool){
        return 
    }
    var srcTypeList *LnsList2_[*Ast_TypeInfo]
    srcTypeList = node.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_retTypeInfoList(_env)
    var dstTypeList *LnsList2_[*Ast_TypeInfo]
    dstTypeList = node.FP.Get_expTypeList(_env)
    var srcTxt string
    srcTxt = _env.GetVM().String_format("arg1 %s", Lns_2DDD(self.FP.type2gotype(_env, srcTypeList.GetAt(1))))
    var dstTxt string
    dstTxt = _env.GetVM().String_format("%s", Lns_2DDD(self.FP.type2gotype(_env, dstTypeList.GetAt(1))))
    {
        var _forFrom0 LnsInt = 2
        var _forTo0 LnsInt = srcTypeList.Len()
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            index := _forWork0
            srcTxt = _env.GetVM().String_format("%s,arg%d %s", Lns_2DDD(srcTxt, index, self.FP.type2gotype(_env, srcTypeList.GetAt(index))))
        }
    }
    {
        var _forFrom1 LnsInt = 2
        var _forTo1 LnsInt = dstTypeList.Len()
        for _forWork1 := _forFrom1; _forWork1 <= _forTo1; _forWork1++ {
            index := _forWork1
            dstTxt = _env.GetVM().String_format("%s,%s", Lns_2DDD(dstTxt, self.FP.type2gotype(_env, dstTypeList.GetAt(index))))
        }
    }
    self.FP.Writeln(_env, _env.GetVM().String_format("func %s(%s) (%s) {", Lns_2DDD(self.FP.getConvGenericsName(_env, &node.Nodes_Node), srcTxt, dstTxt)))
    self.FP.PushIndent(_env, nil)
    self.FP.WriteRaw(_env, "return ")
    for _index, _dstType := range( dstTypeList.Items ) {
        index := _index + 1
        dstType := _dstType
        if index > 1{
            self.FP.WriteRaw(_env, ", ")
        }
        if index > srcTypeList.Len(){
            self.FP.WriteRaw(_env, convPython_literalNil)
        } else { 
            self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d", Lns_2DDD(index)))
            var srcType *Ast_TypeInfo
            srcType = srcTypeList.GetAt(index)
            if srcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                self.FP.outputAny2Type(_env, dstType)
            }
        }
    }
    self.FP.Writeln(_env, "")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 1394: decl @lune.@base.@convPython.convFilter.outputRetType
func (self *convPython_convFilter) outputRetType(_env *LnsEnv, retTypeList *LnsList2_[*Ast_TypeInfo]) {
}
// 1406: decl @lune.@base.@convPython.convFilter.outputDeclFunc
func (self *convPython_convFilter) OutputDeclFunc(_env *LnsEnv, addEnvArg bool,funcInfo LnsAny) *convPython_FuncConv {
    var typeInfo *Ast_TypeInfo
    var name LnsAny
    var prefixType *Ast_TypeInfo
    var extFlag bool
    switch _matchExp0 := funcInfo.(type) {
    case *convPython_FuncInfo__DeclInfo:
        node := _matchExp0.Val1
        workDeclInfo := _matchExp0.Val2
        extFlag = false
        typeInfo = node.FP.Get_expType(_env)
        prefixType = typeInfo.FP.Get_parentInfo(_env)
        if Lns_op_not(workDeclInfo.FP.Get_name(_env)){
            if self.processMode == convPython_ProcessMode__NonClosureFuncDecl{
                name = "_anonymous"
            } else { 
                name = nil
            }
        } else { 
            name = typeInfo.FP.Get_rawTxt(_env)
        }
    case *convPython_FuncInfo__Type:
        workTypeInfo := _matchExp0.Val1
        extFlag = workTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext
        typeInfo = workTypeInfo
        prefixType = typeInfo.FP.Get_parentInfo(_env)
        name = typeInfo.FP.Get_rawTxt(_env)
    case *convPython_FuncInfo__WithClass:
        classType := _matchExp0.Val1
        methodType := _matchExp0.Val2
        extFlag = false
        typeInfo = methodType
        prefixType = classType
        name = typeInfo.FP.Get_rawTxt(_env)
    }
    if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
        self.FP.WriteRaw(_env, "def ")
        self.FP.WriteRaw(_env, "(self ")
        self.FP.Write(_env, self.FP.type2gotype(_env, prefixType))
        self.FP.WriteRaw(_env, ") ")
    } else { 
        self.FP.WriteRaw(_env, "def ")
    }
    if typeInfo.FP.Get_extedType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__FormFunc{
        if name != nil{
            name_40 := name.(string)
            self.FP.outputSymbol(_env, &convPython_SymbolKind__Func{typeInfo}, name_40)
        }
    }
    self.FP.WriteRaw(_env, "(")
    var workType *Ast_TypeInfo
    
    {
        _workType := typeInfo.FP.GetOverridingType(_env)
        if _workType == nil{
            workType = typeInfo
        } else {
            workType = _workType.(*Ast_TypeInfo)
        }
    }
    var retTypeList *LnsList2_[*Ast_TypeInfo]
    if extFlag{
        retTypeList = Lns_unwrap( Lns_car(Ast_convToExtTypeList(_env, self.processInfo, workType.FP.Get_retTypeInfoList(_env)))).(*LnsList2_[*Ast_TypeInfo])
    } else { 
        retTypeList = workType.FP.Get_retTypeInfoList(_env)
    }
    var funcConv *convPython_FuncConv
    funcConv = NewconvPython_FuncConv(_env, retTypeList)
    if addEnvArg{
        self.FP.WriteRaw(_env, self.FP.getEnvArgDecl(_env, workType.FP.Get_argTypeInfoList(_env).Len()))
    }
    switch _matchExp1 := funcInfo.(type) {
    case *convPython_FuncInfo__DeclInfo:
        node := _matchExp1.Val1
        declInfo := _matchExp1.Val2
        for _index, _arg := range( declInfo.FP.Get_argList(_env).Items ) {
            index := _index + 1
            arg := _arg
            if index != 1{
                self.FP.WriteRaw(_env, ",")
            }
            var argType *Ast_TypeInfo
            argType = workType.FP.Get_argTypeInfoList(_env).GetAt(index)
            if argType.FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                {
                    _argNode := Nodes_DeclArgNodeDownCastF(arg.FP)
                    if !Lns_IsNil( _argNode ) {
                        argNode := _argNode.(*Nodes_DeclArgNode)
                        var argName string
                        argName = self.FP.getSymbolSym(_env, argNode.FP.Get_symbolInfo(_env))
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("_%s ", Lns_2DDD(argName)))
                        self.FP.WriteRaw(_env, self.FP.type2gotype(_env, argType))
                        funcConv.FP.Get_argList(_env).Insert(argNode.FP.Get_symbolInfo(_env))
                    } else {
                        convPython_filter_7_(_env, arg, self, node)
                    }
                }
            } else { 
                convPython_filter_7_(_env, arg, self, node)
            }
        }
    case *convPython_FuncInfo__Type:
        for _index, _argType := range( workType.FP.Get_argTypeInfoList(_env).Items ) {
            index := _index + 1
            argType := _argType
            if index != 1{
                self.FP.WriteRaw(_env, ",")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d %s", Lns_2DDD(index, self.FP.type2gotype(_env, argType))))
        }
    case *convPython_FuncInfo__WithClass:
        for _index, _argType := range( workType.FP.Get_argTypeInfoList(_env).Items ) {
            index := _index + 1
            argType := _argType
            if index != 1{
                self.FP.WriteRaw(_env, ",")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d %s", Lns_2DDD(index, self.FP.type2gotype(_env, argType))))
        }
    }
    self.FP.WriteRaw(_env, ")")
    self.FP.outputRetType(_env, retTypeList)
    return funcConv
}
// 1531: decl @lune.@base.@convPython.convFilter.outputConvToForm
func (self *convPython_convFilter) outputConvToForm(_env *LnsEnv, node *Nodes_ExpCastNode) {
    var castType *Ast_TypeInfo
    castType = node.FP.Get_castType(_env).FP.Get_nonnilableType(_env).FP.Get_extedType(_env)
    if castType.FP.Get_kind(_env) != Ast_TypeInfoKind__Form{
        return 
    }
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_extedType(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext) &&
        _env.SetStackVal( funcType.FP.Get_srcTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Form) ).(bool)){
        self.FP.Writeln(_env, _env.GetVM().String_format("      \nfunc %s( luaform LnsAny ) LnsForm {\n    return func (argList []LnsAny) []LnsAny {\n        return %s.RunLoadedfunc( luaform.(*Lns_luaValue), argList )\n    }\n}", Lns_2DDD(self.FP.getConv2formName(_env, &node.Nodes_Node), self.env.FP.getCommonVm(_env))))
        return 
    }
    self.FP.Writeln(_env, _env.GetVM().String_format("# for %d: %s", Lns_2DDD(node.FP.Get_pos(_env).LineNo, Nodes_getNodeKindName(_env, node.FP.Get_kind(_env)))))
    self.FP.WriteRaw(_env, _env.GetVM().String_format("func %s( src func (%s", Lns_2DDD(self.FP.getConv2formName(_env, &node.Nodes_Node), self.FP.getEnvArgDecl(_env, funcType.FP.Get_argTypeInfoList(_env).Len()))))
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList(_env).Items ) {
        index := _index + 1
        argType := _argType
        if index > 1{
            self.FP.WriteRaw(_env, ", ")
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d %s", Lns_2DDD(index, self.FP.type2gotype(_env, argType))))
    }
    self.FP.WriteRaw(_env, ")")
    self.FP.outputRetType(_env, funcType.FP.Get_retTypeInfoList(_env))
    self.FP.Writeln(_env, ") LnsForm {")
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, _env.GetVM().String_format("return func (%s argList []LnsAny) []LnsAny {", Lns_2DDD(self.FP.getEnvArgDecl(_env, 1))))
    self.FP.PushIndent(_env, nil)
    if funcType.FP.Get_retTypeInfoList(_env).Len() > 0{
        self.FP.Write(_env, "return ")
        if funcType.FP.Get_argTypeInfoList(_env).Len() > 0{
            self.FP.Write(_env, "Lns_2DDD(")
        }
    }
    self.FP.WriteRaw(_env, "src(")
    self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, funcType.FP.Get_argTypeInfoList(_env).Len(), self.option.FP.Get_addEnvArg(_env)))
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList(_env).Items ) {
        index := _index + 1
        argType := _argType
        if index > 1{
            self.FP.WriteRaw(_env, ", ")
        }
        if argType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            self.FP.WriteRaw(_env, _env.GetVM().String_format("argList[ %d: ]", Lns_2DDD(index - 1)))
            break
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_getFromMulti( argList, %d )", Lns_2DDD(index - 1)))
    }
    self.FP.WriteRaw(_env, ")")
    if funcType.FP.Get_retTypeInfoList(_env).Len() > 0{
        if funcType.FP.Get_argTypeInfoList(_env).Len() > 0{
            self.FP.Writeln(_env, ")")
        } else { 
            self.FP.Writeln(_env, "")
        }
    } else { 
        self.FP.Writeln(_env, "")
        self.FP.Writeln(_env, "return []LnsAny{}")
    }
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 1608: decl @lune.@base.@convPython.convFilter.processConvStat
func (self *convPython_convFilter) ProcessConvStat(_env *LnsEnv, node *Nodes_ConvStatNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, node.FP.Get_txt(_env))
}
// 1614: decl @lune.@base.@convPython.convFilter.outputTopScopeVar
func (self *convPython_convFilter) outputTopScopeVar(_env *LnsEnv, node *Nodes_DeclVarNode) {
    for _, _symbolInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
        symbolInfo := _symbolInfo
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( symbolInfo.FP.Get_scope(_env) == self.moduleScope) &&
            _env.SetStackVal( node.FP.Get_mode(_env) == Nodes_DeclVarMode__Let) ).(bool)){
            self.FP.Writeln(_env, _env.GetVM().String_format("var %s %s", Lns_2DDD(self.FP.getSymbolSym(_env, symbolInfo), self.FP.type2gotype(_env, symbolInfo.FP.Get_typeInfo(_env)))))
        }
    }
}
// 1623: decl @lune.@base.@convPython.convFilter.outputConvExt
func (self *convPython_convFilter) outputConvExt(_env *LnsEnv, funcNode *Nodes_Node) {
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(funcNode.FP)
        if !Lns_IsNil( _fieldNode ) {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            if fieldNode.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__Ext{
                return 
            }
        } else {
            return 
        }
    }
    self.FP.WriteRaw(_env, _env.GetVM().String_format("func Lns_callExt%s( args []LnsAny ) (", Lns_2DDD(funcNode.FP.GetIdTxt(_env))))
    for _index, _retType := range( funcNode.FP.Get_expType(_env).FP.Get_retTypeInfoList(_env).Items ) {
        index := _index + 1
        retType := _retType
        if index > 1{
            self.FP.WriteRaw(_env, ",")
        }
        self.FP.WriteRaw(_env, self.FP.type2gotype(_env, retType))
    }
    self.FP.Writeln(_env, ") {")
    self.FP.WriteRaw(_env, "    return ")
    for _index, _ := range( funcNode.FP.Get_expType(_env).FP.Get_retTypeInfoList(_env).Items ) {
        index := _index + 1
        if index > 1{
            self.FP.WriteRaw(_env, ",")
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_getFromMulti( args, %d )", Lns_2DDD(index - 1)))
    }
    self.FP.Writeln(_env, "")
    self.FP.Writeln(_env, "}")
}
// 1650: decl @lune.@base.@convPython.convFilter.outputModule
func (self *convPython_convFilter) outputModule(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,addDot bool) {
    self.FP.WriteRaw(_env, self.FP.getModuleSym(_env, moduleTypeInfo, addDot))
}
// 1666: decl @lune.@base.@convPython.convFilter.outputModuleImport
func (self *convPython_convFilter) outputModuleImport(_env *LnsEnv, node *Nodes_DeclClassNode) {
}
// 1688: decl @lune.@base.@convPython.convFilter.outputImport
func (self *convPython_convFilter) outputImport(_env *LnsEnv, node *Nodes_ImportNode) {
    var info *Nodes_ImportInfo
    info = node.FP.Get_info(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.isSameModDir(_env, info.FP.Get_moduleTypeInfo(_env))) ||
        _env.SetStackVal( Ast_isBuiltin(_env, info.FP.Get_moduleTypeInfo(_env).FP.Get_typeId(_env).Id)) ).(bool){
        return 
    }
}
// 1716: decl @lune.@base.@convPython.convFilter.setup
func (self *convPython_convFilter) Setup(_env *LnsEnv) {
    var builtin2runtime *LnsMap2_[*Ast_TypeInfo,string]
    builtin2runtime = NewLnsMap2_[*Ast_TypeInfo,string]( map[*Ast_TypeInfo]string{self.builtinFuncs.Str_gsub:"GETVM.String_gsub",self.builtinFuncs.String_gsub:"GETVM.String_gsub",self.builtinFuncs.Str_find:"GETVM.String_find",self.builtinFuncs.String_find:"GETVM.String_find",self.builtinFuncs.Str_byte:"GETVM.String_byte",self.builtinFuncs.String_byte:"GETVM.String_byte",self.builtinFuncs.Str_format:"GETVM.String_format",self.builtinFuncs.String_format:"GETVM.String_format",self.builtinFuncs.Str_rep:"GETVM.String_rep",self.builtinFuncs.String_rep:"GETVM.String_rep",self.builtinFuncs.Str_gmatch:"GETVM.String_gmatch",self.builtinFuncs.String_gmatch:"GETVM.String_gmatch",self.builtinFuncs.Str_sub:"GETVM.String_sub",self.builtinFuncs.String_sub:"GETVM.String_sub",self.builtinFuncs.Str_lower:"GETVM.String_lower",self.builtinFuncs.String_lower:"GETVM.String_lower",self.builtinFuncs.Str_upper:"GETVM.String_upper",self.builtinFuncs.String_upper:"GETVM.String_upper",self.builtinFuncs.Str_reverse:"GETVM.String_reverse",self.builtinFuncs.String_reverse:"GETVM.String_reverse",Ast_builtinTypeNone:"",})
    
    builtin2runtime.Set(self.builtinFuncs.Str_replace,"GETVM.String_replace")
    builtin2runtime.Set(self.builtinFuncs.Lns_error,"panic")
    builtin2runtime.Set(self.builtinFuncs.Lns_print,"print")
    builtin2runtime.Set(self.builtinFuncs.Lns_type,"Lns_type")
    builtin2runtime.Set(self.builtinFuncs.Lns_require,"Lns_require")
    builtin2runtime.Set(self.builtinFuncs.Lns_tonumber,"Lns_tonumber")
    builtin2runtime.Set(self.builtinFuncs.Lns__load,"GETVM.Load")
    builtin2runtime.Set(self.builtinFuncs.Lns_loadfile,"GETVM.Loadfile")
    builtin2runtime.Set(self.builtinFuncs.Lns_expandLuavalMap,"GETVM.ExpandLuavalMap")
    builtin2runtime.Set(self.builtinFuncs.String_dump,"GETVM.String_dump")
    builtin2runtime.Set(self.builtinFuncs.Io_open,"Lns_io_open")
    builtin2runtime.Set(self.builtinFuncs.Io_popen,"GETVM.IO_popen")
    builtin2runtime.Set(self.builtinFuncs.Package_searchpath,"GETVM.Package_searchpath")
    builtin2runtime.Set(self.builtinFuncs.Os_clock,"GETVM.OS_clock")
    builtin2runtime.Set(self.builtinFuncs.Os_exit,"GETVM.OS_exit")
    builtin2runtime.Set(self.builtinFuncs.Os_remove,"GETVM.OS_remove")
    builtin2runtime.Set(self.builtinFuncs.Os_date,"GETVM.OS_date")
    builtin2runtime.Set(self.builtinFuncs.Os_time,"GETVM.OS_time")
    builtin2runtime.Set(self.builtinFuncs.Os_difftime,"GETVM.OS_difftime")
    builtin2runtime.Set(self.builtinFuncs.Os_rename,"GETVM.OS_rename")
    builtin2runtime.Set(self.builtinFuncs.Math_random,"GETVM.Math_random")
    builtin2runtime.Set(self.builtinFuncs.Math_randomseed,"GETVM.Math_randomseed")
    self.builtin2runtime = builtin2runtime
    self.builtin2runtimeEnv = NewLnsMap2_[*Ast_TypeInfo,string]( map[*Ast_TypeInfo]string{self.builtinFuncs.G__lns_runtime_log:"LnsLog",self.builtinFuncs.G__lns_runtime_enableLog:"LnsStartRunnerLog",self.builtinFuncs.G__lns_runtime_dumpLog:"LnsDumpRunnerLog",self.builtinFuncs.G__lns_sync_createFlag:"LnsCreateSyncFlag",self.builtinFuncs.G__lns_sync_createProcesser:"LnsCreateProcessor",self.builtinFuncs.G_list_insert:".append",self.builtinFuncs.G__list_insert:".append",})
    self.type2gotypeMap = NewLnsMap2_[*Ast_TypeInfo,string]( map[*Ast_TypeInfo]string{Ast_builtinTypeInt:"LnsInt",Ast_builtinTypeReal:"LnsReal",Ast_builtinTypeStem:"LnsAny",Ast_builtinTypeString:"string",Ast_builtinTypeBool:"bool",Ast_builtinTypeProcessor:"*LnsProcessor",self.builtinFuncs.Ostream_:"Lns_oStream",self.builtinFuncs.Istream_:"Lns_iStream",self.builtinFuncs.Luastream_:"Lns_luaStream",})
}
// 1842: decl @lune.@base.@convPython.convFilter.processMethodAsync
func (self *convPython_convFilter) processMethodAsync(_env *LnsEnv, nodeList *LnsList) *LnsList2_[*convPython_ConvRunner] {
    var totalStmtNum LnsInt
    totalStmtNum = 0
    var declMethodNodeList *LnsList2_[*convPython_ProcessDeclMethodItem]
    declMethodNodeList = NewLnsList2_[*convPython_ProcessDeclMethodItem]([]*convPython_ProcessDeclMethodItem{})
    for _, _workNode := range( nodeList.Items ) {
        workNode := _workNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.enableTest) ||
            _env.SetStackVal( Lns_op_not(workNode.FP.Get_inTestBlock(_env))) &&
            _env.SetStackVal( Lns_op_not(workNode.FP.IsModule(_env))) ).(bool){
            if _switch0 := workNode.FP.Get_expType(_env).FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Class {
                for _, _fieldNode := range( workNode.FP.Get_fieldList(_env).Items ) {
                    fieldNode := _fieldNode
                    {
                        _declMethodNode := Nodes_DeclMethodNodeDownCastF(fieldNode.FP)
                        if !Lns_IsNil( _declMethodNode ) {
                            declMethodNode := _declMethodNode.(*Nodes_DeclMethodNode)
                            declMethodNodeList.Insert(NewconvPython_ProcessDeclMethodItem(_env, workNode, declMethodNode))
                            totalStmtNum = totalStmtNum + declMethodNode.FP.Get_declInfo(_env).FP.Get_stmtNum(_env)
                        }
                    }
                }
            }
        }
    }
    LnsLog(_env, _env.GetVM().String_format("class fields (%d, %d)", Lns_2DDD(declMethodNodeList.Len(), totalStmtNum)))
    var runnerList *LnsList2_[*convPython_ConvRunner]
    runnerList = NewLnsList2_[*convPython_ConvRunner]([]*convPython_ConvRunner{})
    var divCount LnsInt
    divCount = self.option.runnerNum
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( totalStmtNum > 1000) &&
        _env.SetStackVal( divCount > 0) ).(bool)){
        var maxStmtCount LnsInt
        maxStmtCount = (totalStmtNum + divCount - 1) / divCount
        var offset LnsInt
        offset = 1
        var len LnsInt
        len = declMethodNodeList.Len()
        {
            var _forFrom0 LnsInt = 1
            var _forTo0 LnsInt = divCount
            for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                var list *LnsList2_[*convPython_ProcessDeclMethodItem]
                list = NewLnsList2_[*convPython_ProcessDeclMethodItem]([]*convPython_ProcessDeclMethodItem{})
                var stmtCount LnsInt
                stmtCount = 0
                for offset <= len {
                    var declFieldInfo *convPython_ProcessDeclMethodItem
                    declFieldInfo = declMethodNodeList.GetAt(offset)
                    offset = offset + 1
                    list.Insert(declFieldInfo)
                    var declMethodNode *Nodes_DeclMethodNode
                    declMethodNode = declFieldInfo.FP.Get_fieldNode(_env)
                    stmtCount = stmtCount + declMethodNode.FP.Get_declInfo(_env).FP.Get_stmtNum(_env)
                    if stmtCount >= maxStmtCount{
                        break
                    }
                }
                var runner *convPython_ConvRunner
                runner = NewconvPython_ConvRunner(_env, self.enableTest, self.ast, self.option, list)
                runnerList.Insert(runner)
                if Lns_op_not(LnsRun(_env, runner.FP, 2, _env.GetVM().String_format("convGo Field - %s", Lns_2DDD(self.streamName)))){
                    runner.FP.Run(_env)
                }
            }
        }
    } else { 
        self.FP.pushProcessMode(_env, convPython_ProcessMode__DeclClass)
        LnsLog(_env, _env.GetVM().String_format("class fields div (%d)", Lns_2DDD(declMethodNodeList.Len())))
        for _, _info := range( declMethodNodeList.Items ) {
            info := _info
            convPython_filter_7_(_env, &info.FP.Get_fieldNode(_env).Nodes_Node, self, &info.FP.Get_classNode(_env).Nodes_Node)
        }
        self.FP.popProcessMode(_env)
    }
    return runnerList
}
// 1910: decl @lune.@base.@convPython.convFilter.processRoot
func (self *convPython_convFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    for _, _importNode := range( node.FP.Get_nodeManager(_env).FP.GetImportNodeList(_env).Items ) {
        importNode := _importNode.(Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        var info *Nodes_ImportInfo
        info = importNode.FP.Get_info(_env)
        self.moduleType2SymbolMap.Set(info.FP.Get_moduleTypeInfo(_env),info.FP.Get_symbolInfo(_env))
    }
    for _pragma := range( node.FP.Get_luneHelperInfo(_env).PragmaSet.Items ) {
        pragma := _pragma
        switch _matchExp0 := pragma.(type) {
        case *LuneControl_Pragma__limit_conv_code:
            codeSet := _matchExp0.Val1
            if Lns_op_not(codeSet.Has(LuneControl_Code__Python)){
                self.FP.Writeln(_env, "# This code is transcompiled by LuneScript.")
                self.FP.Writeln(_env, _env.GetVM().String_format("package %s", Lns_2DDD(self.option.packageName)))
                return 
            }
        }
    }
    var convPython_isUsingInTest func(_env *LnsEnv, aNode *Nodes_Node) bool
    convPython_isUsingInTest = func(_env *LnsEnv, aNode *Nodes_Node) bool {
        for _, _testBlock := range( node.FP.Get_nodeManager(_env).FP.GetTestBlockNodeList(_env).Items ) {
            testBlock := _testBlock.(Nodes_TestBlockNodeDownCast).ToNodes_TestBlockNode()
            if testBlock.FP.IsInnerPos(_env, aNode.FP.Get_pos(_env)){
                return true
            }
        }
        return false
    }
    self.FP.Setup(_env)
    self.FP.Writeln(_env, "# -*- coding:utf-8; -*-")
    self.FP.Writeln(_env, "# This code is transcompiled by LuneScript.")
    self.FP.Writeln(_env, "import sys")
    for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetImportNodeList(_env).Items ) {
        workNode := _workNode.(Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        self.FP.outputImport(_env, workNode)
    }
    for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclClassNodeList(_env).Items ) {
        workNode := _workNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        self.FP.outputModuleImport(_env, workNode)
    }
    self.FP.pushProcessMode(_env, convPython_ProcessMode__DeclTopScopeVar)
    var modSym *Ast_SymbolInfo
    modSym = Lns_unwrap( self.moduleScope.FP.GetSymbolInfoChild(_env, "__mod__")).(*Ast_SymbolInfo)
    self.FP.Writeln(_env, _env.GetVM().String_format("%s = \"%s\"", Lns_2DDD(self.FP.getSymbolSym(_env, modSym), node.FP.Get_moduleTypeInfo(_env).FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), nil))))
    if node.FP.Get_luneHelperInfo(_env).UseUnwrapExp{
        self.FP.Writeln(_env, "def _lnsUnwrap( val, default ):\n   if val is None:\n       if default == None:\n           raise ValueError()\n       return default\n   return val\n")
    }
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_Node)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convPython_filter_7_(_env, workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclEnumNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclEnumNodeDownCast).ToNodes_DeclEnumNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclVarNodeList(_env).Items ) {
        workNode := _workNode.(Nodes_DeclVarNodeDownCast).ToNodes_DeclVarNode()
        self.FP.outputTopScopeVar(_env, workNode)
    }
    self.FP.popProcessMode(_env)
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_Node)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convPython_filter_7_(_env, workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclAlgeNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_Node)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convPython_filter_7_(_env, workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclFormNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclFormNodeDownCast).ToNodes_DeclFormNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_ExpCallNode)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_ExpCallNode) {
            self.FP.processGenericsCall(_env, workNode)
            self.FP.outputNilAccCall(_env, workNode)
            self.FP.outputConvExt(_env, workNode.FP.Get_func(_env))
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpCallNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCallNodeDownCast).ToNodes_ExpCallNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_ExpCastNode)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_ExpCastNode) {
            self.FP.outputConvToForm(_env, workNode)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpCastNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCastNodeDownCast).ToNodes_ExpCastNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_Node)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convPython_filter_7_(_env, workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetTestCaseNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_TestCaseNodeDownCast).ToNodes_TestCaseNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_IfUnwrapNode)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_IfUnwrapNode) {
            var symTypeList *LnsList2_[*Ast_TypeInfo]
            symTypeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
            {
                var _forFrom0 LnsInt = 1
                var _forTo0 LnsInt = workNode.FP.Get_varSymList(_env).Len()
                for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                    symTypeList.Insert(Ast_builtinTypeStem_)
                }
            }
            self.FP.processConvExp(_env, &workNode.Nodes_Node, symTypeList, workNode.FP.Get_expList(_env), false)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetIfUnwrapNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_IfUnwrapNodeDownCast).ToNodes_IfUnwrapNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_ExpSetValNode)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_ExpSetValNode) {
            self.FP.processConvExp(_env, &workNode.Nodes_Node, workNode.FP.Get_exp1(_env).FP.Get_expTypeList(_env), workNode.FP.Get_exp2(_env), false)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpSetValNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpSetValNodeDownCast).ToNodes_ExpSetValNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_ExpCallNode)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_ExpCallNode) {
            var needEnv bool
            needEnv = Lns_op_not(Ast_isBuiltin(_env, workNode.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_typeId(_env).Id))
            self.FP.processConvExp(_env, &workNode.Nodes_Node, workNode.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_argTypeInfoList(_env), workNode.FP.Get_argList(_env), needEnv)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpCallNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCallNodeDownCast).ToNodes_ExpCallNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_ExpCallSuperCtorNode)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_ExpCallSuperCtorNode) {
            var needEnv bool
            needEnv = Lns_op_not(Ast_isBuiltin(_env, workNode.FP.Get_methodType(_env).FP.Get_typeId(_env).Id))
            self.FP.processConvExp(_env, &workNode.Nodes_Node, workNode.FP.Get_methodType(_env).FP.Get_argTypeInfoList(_env), workNode.FP.Get_expList(_env), needEnv)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpCallSuperCtorNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCallSuperCtorNodeDownCast).ToNodes_ExpCallSuperCtorNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_ExpCallSuperNode)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_ExpCallSuperNode) {
            var needEnv bool
            needEnv = Lns_op_not(Ast_isBuiltin(_env, workNode.FP.Get_methodType(_env).FP.Get_typeId(_env).Id))
            self.FP.processConvExp(_env, &workNode.Nodes_Node, workNode.FP.Get_methodType(_env).FP.Get_argTypeInfoList(_env), workNode.FP.Get_expList(_env), needEnv)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpCallSuperNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCallSuperNodeDownCast).ToNodes_ExpCallSuperNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_AliasNode)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_AliasNode) {
            self.FP.Writeln(_env, _env.GetVM().String_format("type %s = %s", Lns_2DDD(self.FP.getSymbolSym(_env, workNode.FP.Get_newSymbol(_env)), self.FP.getTypeSymbol(_env, workNode.FP.Get_typeInfo(_env).FP.Get_aliasSrc(_env)))))
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetAliasNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, tmpNode)
            }
        }
    }
    
    self.FP.pushProcessMode(_env, convPython_ProcessMode__NonClosureFuncDecl)
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_Node)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convPython_filter_7_(_env, workNode, self, &node.Nodes_Node)
            self.FP.Writeln(_env, "")
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclFuncNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    self.FP.popProcessMode(_env)
    var runnerList *LnsList2_[*convPython_ConvRunner]
    runnerList = self.FP.processMethodAsync(_env, node.FP.Get_nodeManager(_env).FP.GetDeclClassNodeList(_env))
    self.FP.pushProcessMode(_env, convPython_ProcessMode__DeclClass)
    {
        var convPython_procNode func(_env *LnsEnv, workNode *Nodes_Node)
        convPython_procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convPython_filter_7_(_env, workNode, self, &node.Nodes_Node)
            self.FP.Writeln(_env, "")
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclClassNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(convPython_isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                convPython_procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    self.FP.popProcessMode(_env)
    for _, _child := range( node.FP.Get_children(_env).Items ) {
        child := _child
        if Lns_op_not(self.ignoreNodeInInnerBlockSet.Has(LnsInt(child.FP.Get_kind(_env)))){
            convPython_filter_7_(_env, child, self, &node.Nodes_Node)
        }
    }
    if self.option.mainModule == Lns_car(_env.GetVM().String_gsub(self.FP.getCanonicalName(_env, self.moduleTypeInfo, false),"@", "")).(string){
        var hasMain bool
        hasMain = false
        for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclFuncNodeList(_env).Items ) {
            workNode := _workNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            if convPython_isMain_3_(_env, workNode.FP.Get_expType(_env)){
                hasMain = true
                break
            }
        }
        if Lns_op_not(hasMain){
            var callArgs string
            if self.option.FP.Get_addEnvArg(_env){
                callArgs = "_env"
            } else { 
                callArgs = ""
            }
            var moduleSym string
            moduleSym = self.FP.getModuleName(_env, self.moduleTypeInfo, false)
            self.FP.Writeln(_env, _env.GetVM().String_format("func %s___main( %sargs *LnsList ) LnsInt {", Lns_2DDD(convPython_concatGLSym_11_(_env, moduleSym, true), self.FP.getEnvArgDecl(_env, 1))))
            self.FP.Writeln(_env, _env.GetVM().String_format("Lns_%s_init( %s )", Lns_2DDD(moduleSym, callArgs)))
            self.FP.Writeln(_env, "return 0")
            self.FP.Writeln(_env, "}")
        }
    }
    for _, _runner := range( runnerList.Items ) {
        runner := _runner
        self.FP.WriteRaw(_env, runner.FP.GetResult(_env))
    }
    if self.enableTest{
        self.FP.Writeln(_env, "if __name__ == '__main__':\n    unittest.main()\n")
    }
}
// 2175: decl @lune.@base.@convPython.convFilter.processSubfile
func (self *convPython_convFilter) ProcessSubfile(_env *LnsEnv, node *Nodes_SubfileNode,_opt LnsAny) {
}
// 2181: decl @lune.@base.@convPython.convFilter.processRequest
func (self *convPython_convFilter) ProcessRequest(_env *LnsEnv, node *Nodes_RequestNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "func() ")
    self.FP.outputRetType(_env, node.FP.Get_expTypeList(_env))
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    var retVars *LnsList2_[string]
    retVars = NewLnsList2_[string]([]string{})
    for _index, _retType := range( node.FP.Get_expTypeList(_env).Items ) {
        index := _index + 1
        retType := _retType
        var varSym string
        varSym = _env.GetVM().String_format("ret%d", Lns_2DDD(index))
        retVars.Insert(varSym)
        self.FP.Writeln(_env, _env.GetVM().String_format("var %s %s", Lns_2DDD(varSym, self.FP.type2gotype(_env, retType))))
    }
    convPython_filter_7_(_env, node.FP.Get_processor(_env), self, &node.Nodes_Node)
    self.FP.Writeln(_env, ".Request( _env, func( _env *LnsEnv ) {")
    self.FP.PushIndent(_env, nil)
    if retVars.Len() > 0{
        for _index, _varSym := range( retVars.Items ) {
            index := _index + 1
            varSym := _varSym
            self.FP.Writeln(_env, varSym)
            if index != retVars.Len(){
                self.FP.WriteRaw(_env, ",")
            }
        }
        self.FP.WriteRaw(_env, "=")
    }
    convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "")
    self.FP.WriteRaw(_env, "}")
    self.FP.Writeln(_env, ")")
    self.FP.PopIndent(_env)
    if retVars.Len() > 0{
        self.FP.WriteRaw(_env, "return ")
        for _index, _varSym := range( retVars.Items ) {
            index := _index + 1
            varSym := _varSym
            self.FP.Writeln(_env, varSym)
            if index != retVars.Len(){
                self.FP.WriteRaw(_env, ",")
            }
        }
        self.FP.Writeln(_env, "")
    }
    self.FP.WriteRaw(_env, "}()")
}
// 2229: decl @lune.@base.@convPython.convFilter.processAsyncLock
func (self *convPython_convFilter) ProcessAsyncLock(_env *LnsEnv, node *Nodes_AsyncLockNode,_opt LnsAny) {
    if _switch0 := node.FP.Get_lockKind(_env); _switch0 == Nodes_LockKind__AsyncLock || _switch0 == Nodes_LockKind__LuaLock {
        self.FP.Writeln(_env, _env.GetVM().String_format("Lns_LockEnvSync( %s, %d, func () {", Lns_2DDD(self.env.FP.getEnv(_env), node.FP.Get_pos(_env).LineNo)))
        convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        self.FP.Writeln(_env, "})")
    } else if _switch0 == Nodes_LockKind__LuaDepend || _switch0 == Nodes_LockKind__LuaGo {
        convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    }
}
// 2247: decl @lune.@base.@convPython.convFilter.processBlockSub
func (self *convPython_convFilter) ProcessBlockSub(_env *LnsEnv, node *Nodes_BlockNode,_opt LnsAny) {
    if _switch0 := node.FP.Get_blockKind(_env); _switch0 == Nodes_BlockKind__Block {
        self.FP.Writeln(_env, "{")
    } else {
        self.FP.PushIndent(_env, nil)
    }
    self.FP.pushProcessMode(_env, convPython_ProcessMode__Main)
    for _, _child := range( node.FP.Get_stmtList(_env).Items ) {
        child := _child
        if Lns_op_not(self.ignoreNodeInInnerBlockSet.Has(LnsInt(child.FP.Get_kind(_env)))){
            convPython_filter_7_(_env, child, self, &node.Nodes_Node)
        }
    }
    self.FP.popProcessMode(_env)
    if _switch1 := node.FP.Get_blockKind(_env); _switch1 == Nodes_BlockKind__Block {
        self.FP.Writeln(_env, "}")
    } else {
        self.FP.PopIndent(_env)
    }
}
// 2277: decl @lune.@base.@convPython.convFilter.expList2Slice
func (self *convPython_convFilter) expList2Slice(_env *LnsEnv, subList *Nodes_ExpListNode,toStem bool) {
    for _, _exp := range( subList.FP.Get_expList(_env).Items ) {
        exp := _exp
        convPython_filter_7_(_env, exp, self, &subList.Nodes_Node)
    }
}
// 2343: decl @lune.@base.@convPython.convFilter.processSetFromExpList
func (self *convPython_convFilter) processSetFromExpList(_env *LnsEnv, convArgFuncName string,dstTypeList *LnsList2_[*Ast_TypeInfo],expListNode *Nodes_ExpListNode,addEnvArg bool) {
    if _switch0 := convPython_getExpListKind_21_(_env, dstTypeList, expListNode, addEnvArg); _switch0 == convPython_ExpListKind__Conv {
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s(", Lns_2DDD(convArgFuncName)))
        var mRetIndex LnsAny
        mRetIndex = _env.NilAccFin(_env.NilAccPush(expListNode.FP.Get_mRetExp(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
        self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, expListNode.FP.Get_expList(_env).Len(), addEnvArg))
        for _index, _exp := range( expListNode.FP.Get_expList(_env).Items ) {
            index := _index + 1
            exp := _exp
            if exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
                break
            }
            if index != 1{
                self.FP.WriteRaw(_env, ", ")
            }
            if mRetIndex == index{
                self.FP.WriteRaw(_env, "Lns_2DDD(")
                convPython_filter_7_(_env, Nodes_getCastUnwraped(_env, exp), self, &expListNode.Nodes_Node)
                self.FP.WriteRaw(_env, ")")
                break
            }
            convPython_filter_7_(_env, exp, self, &expListNode.Nodes_Node)
            if _env.NilAccFin(_env.NilAccPush(expListNode.FP.Get_mRetExp(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)})) == index{
                break
            }
        }
        self.FP.WriteRaw(_env, ")")
    } else if _switch0 == convPython_ExpListKind__Slice {
        self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, dstTypeList.Len(), addEnvArg))
        for _index, _argType := range( dstTypeList.Items ) {
            index := _index + 1
            argType := _argType
            if index != 1{
                self.FP.WriteRaw(_env, ", ")
            }
            if expListNode.FP.Get_expList(_env).Len() >= index{
                var argExp *Nodes_Node
                argExp = expListNode.FP.Get_expList(_env).GetAt(index)
                {
                    _exp2ddd := Nodes_ExpToDDDNodeDownCastF(argExp.FP)
                    if !Lns_IsNil( _exp2ddd ) {
                        exp2ddd := _exp2ddd.(*Nodes_ExpToDDDNode)
                        self.FP.expList2Slice(_env, exp2ddd.FP.Get_expList(_env), false)
                    } else {
                        if argExp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
                            if argType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                                self.FP.WriteRaw(_env, "[]LnsAny{}")
                            } else { 
                                self.FP.WriteRaw(_env, convPython_literalNil)
                            }
                        } else { 
                            convPython_filter_7_(_env, argExp, self, &expListNode.Nodes_Node)
                        }
                    }
                }
            } else { 
                self.FP.WriteRaw(_env, "[]LnsAny{}")
            }
        }
    } else if _switch0 == convPython_ExpListKind__Direct {
        self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, dstTypeList.Len(), addEnvArg))
        var mRetIndex LnsAny
        mRetIndex = _env.NilAccFin(_env.NilAccPush(expListNode.FP.Get_mRetExp(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
        for _index, _dstType := range( dstTypeList.Items ) {
            index := _index + 1
            dstType := _dstType
            if mRetIndex == index - 1{
                break
            }
            if index != 1{
                self.FP.WriteRaw(_env, ", ")
            }
            var exp *Nodes_Node
            if expListNode.FP.Get_expList(_env).Len() < index{
                exp = &self.noneNode.Nodes_Node
            } else { 
                exp = expListNode.FP.Get_expList(_env).GetAt(index)
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( index == dstTypeList.Len()) &&
                _env.SetStackVal( dstType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( expListNode.FP.Get_expList(_env).Len() < index) ||
                    _env.SetStackVal( exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr) ).(bool){
                    self.FP.WriteRaw(_env, "[]LnsAny{}")
                } else { 
                    convPython_filter_7_(_env, exp, self, &expListNode.Nodes_Node)
                }
            } else { 
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( expListNode.FP.Get_expList(_env).Len() < index) ||
                    _env.SetStackVal( exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr) ).(bool){
                    self.FP.WriteRaw(_env, convPython_literalNil)
                } else if expListNode.FP.Get_expTypeList(_env).GetAt(index).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    self.FP.WriteRaw(_env, "Lns_car(")
                    convPython_filter_7_(_env, exp, self, &expListNode.Nodes_Node)
                    self.FP.WriteRaw(_env, ")")
                } else { 
                    convPython_filter_7_(_env, exp, self, &expListNode.Nodes_Node)
                }
            }
        }
    }
}
// 2450: decl @lune.@base.@convPython.convFilter.processStmtExp
func (self *convPython_convFilter) ProcessStmtExp(_env *LnsEnv, node *Nodes_StmtExpNode,_opt LnsAny) {
    convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Writeln(_env, "")
}
// 2457: decl @lune.@base.@convPython.convFilter.processDeclAlge
func (self *convPython_convFilter) ProcessDeclAlge(_env *LnsEnv, node *Nodes_DeclAlgeNode,_opt LnsAny) {
    self.FP.Writeln(_env, _env.GetVM().String_format("# decl alge -- %s", Lns_2DDD(node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))))
    self.FP.Writeln(_env, _env.GetVM().String_format("type %s = LnsAny", Lns_2DDD(self.FP.getTypeSymbol(_env, &node.FP.Get_algeType(_env).Ast_TypeInfo))))
    {
        __forsortCollection0 := node.FP.Get_algeType(_env).FP.Get_valInfoMap(_env)
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            valInfo := __forsortCollection0.Items[ ___forsortKey0 ]
            var algeSym string
            algeSym = self.FP.getAlgeSymbol(_env, valInfo)
            self.FP.Writeln(_env, _env.GetVM().String_format("type %s struct{", Lns_2DDD(algeSym)))
            for _index, _paramType := range( valInfo.FP.Get_typeList(_env).Items ) {
                index := _index + 1
                paramType := _paramType
                self.FP.Writeln(_env, _env.GetVM().String_format("Val%d %s", Lns_2DDD(index, self.FP.type2gotype(_env, paramType))))
            }
            self.FP.Writeln(_env, "}")
            if valInfo.FP.Get_typeList(_env).Len() == 0{
                self.FP.Writeln(_env, _env.GetVM().String_format("var %s_Obj = &%s{}", Lns_2DDD(algeSym, algeSym)))
            }
            self.FP.Writeln(_env, _env.GetVM().String_format("func (self *%s) GetTxt() string {", Lns_2DDD(algeSym)))
            self.FP.Writeln(_env, _env.GetVM().String_format("return \"%s.%s\"", Lns_2DDD(node.FP.Get_algeType(_env).FP.Get_rawTxt(_env), valInfo.FP.Get_name(_env))))
            self.FP.Writeln(_env, "}")
        }
    }
}
// 2479: decl @lune.@base.@convPython.convFilter.processNewAlgeVal
func (self *convPython_convFilter) ProcessNewAlgeVal(_env *LnsEnv, node *Nodes_NewAlgeValNode,_opt LnsAny) {
    var algeSym string
    algeSym = self.FP.getAlgeSymbol(_env, node.FP.Get_valInfo(_env))
    if node.FP.Get_valInfo(_env).FP.Get_typeList(_env).Len() == 0{
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s_Obj", Lns_2DDD(algeSym)))
    } else { 
        self.FP.WriteRaw(_env, _env.GetVM().String_format("&%s{", Lns_2DDD(algeSym)))
        for _index, _param := range( node.FP.Get_paramList(_env).Items ) {
            index := _index + 1
            param := _param
            if index > 1{
                self.FP.WriteRaw(_env, ", ")
            }
            convPython_filter_7_(_env, param, self, &node.Nodes_Node)
        }
        self.FP.WriteRaw(_env, "}")
    }
}
// 2498: decl @lune.@base.@convPython.convFilter.processDeclMember
func (self *convPython_convFilter) ProcessDeclMember(_env *LnsEnv, node *Nodes_DeclMemberNode,_opt LnsAny) {
    self.FP.outputSymbol(_env, &convPython_SymbolKind__Member{Ast_isPubToExternal(_env, node.FP.Get_accessMode(_env))}, node.FP.Get_name(_env).Txt)
    self.FP.WriteRaw(_env, _env.GetVM().String_format(" %s", Lns_2DDD(self.FP.type2gotype(_env, node.FP.Get_refType(_env).FP.Get_expType(_env)))))
    self.FP.Writeln(_env, "")
}
// 2507: decl @lune.@base.@convPython.convFilter.processExpMacroArgExp
func (self *convPython_convFilter) ProcessExpMacroArgExp(_env *LnsEnv, node *Nodes_ExpMacroArgExpNode,_opt LnsAny) {
    convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
}
// 2514: decl @lune.@base.@convPython.convFilter.processExpMacroExp
func (self *convPython_convFilter) ProcessExpMacroExp(_env *LnsEnv, node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    var macroParentFull string
    macroParentFull = node.FP.Get_macroType(_env).FP.GetParentFullName(_env, self.FP.Get_typeNameCtrl(_env), nil, false)
    if macroParentFull == "@lune.@base.@Testing."{
        self.FP.WriteRaw(_env, "self.")
        if _switch0 := node.FP.Get_macroType(_env).FP.GetTxt(_env, nil, nil, nil); _switch0 == "_eq" {
            self.FP.WriteRaw(_env, "assertEqual( ")
            {
                _expList := node.FP.Get_expList(_env)
                if !Lns_IsNil( _expList ) {
                    expList := _expList.(*Nodes_ExpListNode)
                    convPython_filter_7_(_env, expList.FP.Get_expList(_env).GetAt(2), self, &node.Nodes_Node)
                    self.FP.Write(_env, ", ")
                    convPython_filter_7_(_env, expList.FP.Get_expList(_env).GetAt(3), self, &node.Nodes_Node)
                    self.FP.WriteRaw(_env, " )")
                }
            }
        } else {
            Util_err(_env, _env.GetVM().String_format("not support macro -- %s", Lns_2DDD(node.FP.Get_macroType(_env).FP.GetTxt(_env, nil, nil, nil))))
        }
    } else { 
        for _, _stmt := range( node.FP.Get_stmtList(_env).Items ) {
            stmt := _stmt
            if Lns_op_not(self.ignoreNodeInInnerBlockSet.Has(LnsInt(stmt.FP.Get_kind(_env)))){
                convPython_filter_7_(_env, stmt, self, &node.Nodes_Node)
            }
        }
    }
}
// 2545: decl @lune.@base.@convPython.convFilter.processDeclMacro
func (self *convPython_convFilter) ProcessDeclMacro(_env *LnsEnv, node *Nodes_DeclMacroNode,_opt LnsAny) {
}
// 2551: decl @lune.@base.@convPython.convFilter.processExpMacroStat
func (self *convPython_convFilter) ProcessExpMacroStat(_env *LnsEnv, node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processExpMacroStat"
    Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
}
// 2557: decl @lune.@base.@convPython.convFilter.outputDeclFuncArg
func (self *convPython_convFilter) outputDeclFuncArg(_env *LnsEnv, funcType *Ast_TypeInfo) {
    self.FP.WriteRaw(_env, self.FP.getEnvArgDecl(_env, funcType.FP.Get_argTypeInfoList(_env).Len()))
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList(_env).Items ) {
        index := _index + 1
        argType := _argType
        if index != 1{
            self.FP.WriteRaw(_env, ", ")
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d ", Lns_2DDD(index)))
        self.FP.WriteRaw(_env, self.FP.type2gotype(_env, argType))
    }
}
// 2568: decl @lune.@base.@convPython.convFilter.isImplementedRunner
func (self *convPython_convFilter) isImplementedRunner(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    for _, _ifType := range( typeInfo.FP.Get_interfaceList(_env).Items ) {
        ifType := _ifType
        if ifType == self.builtinFuncs.G__runner_{
            return true
        }
    }
    return false
}
// 2577: decl @lune.@base.@convPython.convFilter.processDeclConstr
func (self *convPython_convFilter) ProcessDeclConstr(_env *LnsEnv, node *Nodes_DeclConstrNode,_opt LnsAny) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType(_env).FP.Get_parentInfo(_env)
    var className string
    className = self.FP.getTypeSymbol(_env, classType)
    self.FP.Writeln(_env, _env.GetVM().String_format("# %d: %s", Lns_2DDD(node.FP.Get_pos(_env).LineNo, Nodes_getNodeKindName(_env, node.FP.Get_kind(_env)))))
    self.FP.WriteRaw(_env, _env.GetVM().String_format("func (self *%s) %s(", Lns_2DDD(className, self.FP.getConstrSymbol(_env, classType))))
    self.FP.WriteRaw(_env, self.FP.getEnvArgDecl(_env, node.FP.Get_declInfo(_env).FP.Get_argList(_env).Len()))
    for _index, _arg := range( node.FP.Get_declInfo(_env).FP.Get_argList(_env).Items ) {
        index := _index + 1
        arg := _arg
        if index != 1{
            self.FP.WriteRaw(_env, ",")
        }
        convPython_filter_7_(_env, arg, self, &node.Nodes_Node)
    }
    self.FP.Writeln(_env, ") {")
    if self.FP.isImplementedRunner(_env, classType){
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, "self._syncFlag = &Lns_syncFlag{}")
        self.FP.PopIndent(_env)
    }
    convPython_filter_7_(_env, &Lns_unwrap( node.FP.Get_declInfo(_env).FP.Get_body(_env)).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "}")
}
// 2606: decl @lune.@base.@convPython.convFilter.processDeclDestr
func (self *convPython_convFilter) ProcessDeclDestr(_env *LnsEnv, node *Nodes_DeclDestrNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processDeclDestr"
    Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
}
// 2612: decl @lune.@base.@convPython.convFilter.processExpCallSuperCtor
func (self *convPython_convFilter) ProcessExpCallSuperCtor(_env *LnsEnv, node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("self.%s(", Lns_2DDD(self.FP.getConstrSymbol(_env, node.FP.Get_superType(_env)))))
    {
        _argList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, argList), node.FP.Get_methodType(_env).FP.Get_argTypeInfoList(_env), argList, self.option.FP.Get_addEnvArg(_env))
        } else {
            self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, 0, self.option.FP.Get_addEnvArg(_env)))
        }
    }
    self.FP.Writeln(_env, ")")
}
// 2626: decl @lune.@base.@convPython.convFilter.processExpCallSuper
func (self *convPython_convFilter) ProcessExpCallSuper(_env *LnsEnv, node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("self.%s.%s(", Lns_2DDD(self.FP.getTypeSymbol(_env, node.FP.Get_methodType(_env).FP.Get_parentInfo(_env)), self.FP.getFuncSymbol(_env, node.FP.Get_methodType(_env)))))
    {
        _argList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, argList), node.FP.Get_methodType(_env).FP.Get_argTypeInfoList(_env), argList, self.option.FP.Get_addEnvArg(_env))
        } else {
            self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, 0, self.option.FP.Get_addEnvArg(_env)))
        }
    }
    self.FP.WriteRaw(_env, ")")
}
// 2641: decl @lune.@base.@convPython.convFilter.outputDummyReturn
func (self *convPython_convFilter) outputDummyReturn(_env *LnsEnv, retTypeInfoList *LnsList2_[*Ast_TypeInfo]) {
}
// 2675: decl @lune.@base.@convPython.convFilter.outputDeclFuncInfo
func (self *convPython_convFilter) OutputDeclFuncInfo(_env *LnsEnv, node *Nodes_Node,declInfo *Nodes_DeclFuncInfo) {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(self.enableTest)) &&
        _env.SetStackVal( node.FP.Get_inTestBlock(_env)) ).(bool)){
        return 
    }
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_expType(_env)
    if funcType.FP.Get_abstractFlag(_env){
        return 
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( declInfo.FP.Get_name(_env)) &&
        _env.SetStackVal( Lns_op_not(convPython_isClosure_9_(_env, funcType))) )){
        self.FP.Writeln(_env, _env.GetVM().String_format("# %d: decl %s", Lns_2DDD(node.FP.Get_pos(_env).LineNo, self.FP.getCanonicalName(_env, funcType, false))))
    }
    var convFunc *convPython_FuncConv
    convFunc = self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convPython_FuncInfo__DeclInfo{node, declInfo})
    self.FP.Writeln(_env, ":")
    if declInfo.FP.Get_has__func__Symbol(_env){
        var nameSpace string
        nameSpace = self.FP.getCanonicalName(_env, funcType.FP.Get_parentInfo(_env), false)
        var funcName string
        {
            _name := declInfo.FP.Get_name(_env)
            if !Lns_IsNil( _name ) {
                name := _name.(*Types_Token)
                funcName = name.Txt
            } else {
                funcName = "<anonymous>"
            }
        }
        var funcSym_ *Ast_SymbolInfo
        funcSym_ = Lns_unwrap( _env.NilAccFin(_env.NilAccPush(funcType.FP.Get_scope(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(_env, "__func__")})/* 2712:29 */)).(*Ast_SymbolInfo)
        self.FP.Writeln(_env, _env.GetVM().String_format("%s := \"%s.%s\"", Lns_2DDD(self.FP.getSymbolSym(_env, funcSym_), nameSpace, funcName)))
    }
    for _, _convArg := range( convFunc.FP.Get_argList(_env).Items ) {
        convArg := _convArg
        if Lns_isCondTrue( convArg.FP.Get_posForModToRef(_env)){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := _%s", Lns_2DDD(self.FP.getSymbolSym(_env, convArg), self.FP.getSymbolSym(_env, convArg))))
            self.FP.outputAny2Type(_env, convArg.FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
    }
    {
        _body := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( _body ) {
            body := _body.(*Nodes_BlockNode)
            convPython_filter_7_(_env, &body.Nodes_Node, self, node)
            var retTypeInfoList *LnsList2_[*Ast_TypeInfo]
            retTypeInfoList = funcType.FP.Get_retTypeInfoList(_env)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( retTypeInfoList.Len() > 0) &&
                _env.SetStackVal( retTypeInfoList.GetAt(retTypeInfoList.Len()) != Ast_builtinTypeNeverRet) ).(bool)){
                var needReturn bool
                needReturn = false
                {
                    var _forFrom0 LnsInt = body.FP.Get_stmtList(_env).Len()
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
                        var statment *Nodes_Node
                        statment = body.FP.Get_stmtList(_env).GetAt(index)
                        if _switch0 := statment.FP.Get_kind(_env); _switch0 == Nodes_NodeKind_get_BlankLine(_env) {
                        } else if _switch0 == Nodes_NodeKind_get_Return(_env) {
                            break
                        } else {
                            needReturn = true
                            break
                        }
                        _forWork0 += _forDelta0
                    }
                }
                if needReturn{
                    self.FP.outputDummyReturn(_env, funcType.FP.Get_retTypeInfoList(_env))
                }
            }
        }
    }
    if Lns_isCondTrue( declInfo.FP.Get_name(_env)){
        self.FP.Writeln(_env, "")
    }
}
// 2765: decl @lune.@base.@convPython.convFilter.processDeclMethod
func (self *convPython_convFilter) ProcessDeclMethod(_env *LnsEnv, node *Nodes_DeclMethodNode,_opt LnsAny) {
    self.FP.OutputDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env))
}
// 2771: decl @lune.@base.@convPython.convFilter.processProtoMethod
func (self *convPython_convFilter) ProcessProtoMethod(_env *LnsEnv, node *Nodes_ProtoMethodNode,_opt LnsAny) {
}
// 2777: decl @lune.@base.@convPython.convFilter.getEnumGetTxtSym
func (self *convPython_convFilter) getEnumGetTxtSym(_env *LnsEnv, enumType *Ast_EnumTypeInfo) string {
    return _env.GetVM().String_format("%s_getTxt", Lns_2DDD(self.FP.getTypeSymbol(_env, &enumType.Ast_TypeInfo)))
}
// 2781: decl @lune.@base.@convPython.convFilter.processDeclEnum
func (self *convPython_convFilter) ProcessDeclEnum(_env *LnsEnv, node *Nodes_DeclEnumNode,_opt LnsAny) {
    if _switch0 := self.processMode; _switch0 == convPython_ProcessMode__DeclTopScopeVar {
        self.FP.Writeln(_env, _env.GetVM().String_format("# decl enum -- %s ", Lns_2DDD(node.FP.Get_enumType(_env).FP.GetTxt(_env, nil, nil, nil))))
        self.FP.Writeln(_env, _env.GetVM().String_format("type %s = %s", Lns_2DDD(self.FP.getTypeSymbol(_env, &node.FP.Get_enumType(_env).Ast_TypeInfo), self.FP.type2gotype(_env, &node.FP.Get_enumType(_env).Ast_TypeInfo))))
        {
            __forsortCollection0 := node.FP.Get_enumType(_env).FP.Get_name2EnumValInfo(_env)
            __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
            __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
            for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
                valInfo := __forsortCollection0.Items[ ___forsortKey0 ]
                self.FP.WriteRaw(_env, _env.GetVM().String_format("const %s = ", Lns_2DDD(self.FP.getSymbol(_env, &convPython_SymbolKind__Static{&node.FP.Get_enumType(_env).Ast_TypeInfo}, valInfo.FP.Get_name(_env)))))
                switch _matchExp0 := valInfo.FP.Get_val(_env).(type) {
                case *Ast_EnumLiteral__Int:
                    val := _matchExp0.Val1
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("%d", Lns_2DDD(val)))
                case *Ast_EnumLiteral__Real:
                    val := _matchExp0.Val1
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("%g", Lns_2DDD(val)))
                case *Ast_EnumLiteral__Str:
                    val := _matchExp0.Val1
                    self.FP.WriteRaw(_env, convPython_str2gostr_14_(_env, _env.GetVM().String_format("\"%s\"", Lns_2DDD(val))))
                }
                self.FP.Writeln(_env, "")
            }
        }
        var listName string
        listName = _env.GetVM().String_format("%sList_", Lns_2DDD(self.FP.getTypeSymbol(_env, &node.FP.Get_enumType(_env).Ast_TypeInfo)))
        self.FP.Writeln(_env, _env.GetVM().String_format("var %s = NewLnsList( []LnsAny {", Lns_2DDD(listName)))
        for _, _valName := range( node.FP.Get_valueNameList(_env).Items ) {
            valName := _valName
            self.FP.Writeln(_env, _env.GetVM().String_format("  %s,", Lns_2DDD(self.FP.getSymbol(_env, &convPython_SymbolKind__Static{&node.FP.Get_enumType(_env).Ast_TypeInfo}, valName.Txt))))
        }
        self.FP.Writeln(_env, "})")
        var scope *Ast_Scope
        scope = Lns_unwrap( node.FP.Get_enumType(_env).FP.Get_scope(_env)).(*Ast_Scope)
        self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convPython_FuncInfo__Type{Lns_unwrap( scope.FP.GetTypeInfoChild(_env, "get__allList")).(*Ast_TypeInfo)})
        self.FP.Writeln(_env, "{")
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, _env.GetVM().String_format("return %s", Lns_2DDD(listName)))
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
        var mapName string
        mapName = _env.GetVM().String_format("%sMap_", Lns_2DDD(self.FP.getTypeSymbol(_env, &node.FP.Get_enumType(_env).Ast_TypeInfo)))
        self.FP.Writeln(_env, _env.GetVM().String_format("var %s = map[%s]string {", Lns_2DDD(mapName, self.FP.type2gotype(_env, node.FP.Get_enumType(_env).FP.Get_valTypeInfo(_env)))))
        {
            __forsortCollection1 := node.FP.Get_enumType(_env).FP.Get_name2EnumValInfo(_env)
            __forsortSorted1 := __forsortCollection1.CreateKeyListStr()
            __forsortSorted1.Sort( _env, LnsItemKindStr, nil )
            for _, ___forsortKey1 := range( __forsortSorted1.Items ) {
                valInfo := __forsortCollection1.Items[ ___forsortKey1 ]
                self.FP.Writeln(_env, _env.GetVM().String_format("  %s: \"%s.%s\",", Lns_2DDD(self.FP.getSymbol(_env, &convPython_SymbolKind__Static{&node.FP.Get_enumType(_env).Ast_TypeInfo}, valInfo.FP.Get_name(_env)), node.FP.Get_expType(_env).FP.Get_rawTxt(_env), valInfo.FP.Get_name(_env))))
            }
        }
        self.FP.Writeln(_env, "}")
        self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convPython_FuncInfo__Type{Lns_unwrap( scope.FP.GetTypeInfoChild(_env, "_from")).(*Ast_TypeInfo)})
        self.FP.Writeln(_env, "{")
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, _env.GetVM().String_format("if _, ok := %s[arg1]; ok { return arg1 }", Lns_2DDD(mapName)))
        self.FP.Writeln(_env, _env.GetVM().String_format("return %s", Lns_2DDD(convPython_literalNil)))
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
        self.FP.Writeln(_env, "")
        self.FP.Writeln(_env, _env.GetVM().String_format("func %s(arg1 %s) string {", Lns_2DDD(self.FP.getEnumGetTxtSym(_env, node.FP.Get_enumType(_env)), self.FP.type2gotype(_env, node.FP.Get_enumType(_env).FP.Get_valTypeInfo(_env)))))
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, _env.GetVM().String_format("return %s[arg1];", Lns_2DDD(mapName)))
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    } else {
    }
}
// 2863: decl @lune.@base.@convPython.convFilter.processUnwrapSet
func (self *convPython_convFilter) ProcessUnwrapSet(_env *LnsEnv, node *Nodes_UnwrapSetNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processUnwrapSet"
    Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
}
// 2869: decl @lune.@base.@convPython.convFilter.processIfUnwrap
func (self *convPython_convFilter) ProcessIfUnwrap(_env *LnsEnv, node *Nodes_IfUnwrapNode,_opt LnsAny) {
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    var tempTypeList *LnsList2_[*Ast_TypeInfo]
    tempTypeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    for _index, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        index := _index + 1
        varSym := _varSym
        if index > 1{
            self.FP.WriteRaw(_env, ", ")
        }
        if varSym.FP.Get_name(_env) == "_"{
            self.FP.WriteRaw(_env, "_")
        } else { 
            self.FP.WriteRaw(_env, "_" + self.FP.getSymbolSym(_env, varSym))
        }
        tempTypeList.Insert(Ast_builtinTypeStem_)
    }
    if convPython_getExpListKind_21_(_env, tempTypeList, node.FP.Get_expList(_env), self.option.FP.Get_addEnvArg(_env)) == convPython_ExpListKind__Direct{
        {
            var _forFrom0 LnsInt = node.FP.Get_varSymList(_env).Len() + 1
            var _forTo0 LnsInt = node.FP.Get_expList(_env).FP.Get_expTypeList(_env).Len()
            for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                self.FP.WriteRaw(_env, ", _")
            }
        }
    }
    self.FP.WriteRaw(_env, " := ")
    self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, node.FP.Get_expList(_env)), tempTypeList, node.FP.Get_expList(_env), false)
    self.FP.Writeln(_env, "")
    self.FP.WriteRaw(_env, "if ")
    var hasSym bool
    hasSym = false
    for _, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        varSym := _varSym
        if varSym.FP.Get_name(_env) != "_"{
            if hasSym{
                self.FP.WriteRaw(_env, " && ")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("!Lns_IsNil( _%s )", Lns_2DDD(self.FP.getSymbolSym(_env, varSym))))
            hasSym = true
        }
    }
    self.FP.Writeln(_env, " {")
    self.FP.PushIndent(_env, nil)
    for _index, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        index := _index + 1
        varSym := _varSym
        if varSym.FP.Get_name(_env) != "_"{
            if varSym.FP.HasAccess(_env){
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := _%s", Lns_2DDD(self.FP.getSymbolSym(_env, varSym), self.FP.getSymbolSym(_env, varSym))))
                if node.FP.Get_expList(_env).FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                    self.FP.outputAny2Type(_env, varSym.FP.Get_typeInfo(_env))
                }
                self.FP.Writeln(_env, "")
            }
        }
    }
    self.FP.PopIndent(_env)
    convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    {
        _nilBlock := node.FP.Get_nilBlock(_env)
        if !Lns_IsNil( _nilBlock ) {
            nilBlock := _nilBlock.(*Nodes_BlockNode)
            self.FP.Writeln(_env, "} else {")
            convPython_filter_7_(_env, &nilBlock.Nodes_Node, self, &node.Nodes_Node)
            self.FP.Writeln(_env, "}")
        } else {
            self.FP.Writeln(_env, "}")
        }
    }
    self.FP.PopIndent(_env)
    self.FP.WriteRaw(_env, "}")
    self.FP.Writeln(_env, "")
}
// 2947: decl @lune.@base.@convPython.convFilter.outputLetVar
func (self *convPython_convFilter) outputLetVar(_env *LnsEnv, node *Nodes_DeclVarNode) {
    var convPython_declVar func(_env *LnsEnv)
    convPython_declVar = func(_env *LnsEnv) {
        if node.FP.Get_symbolInfoList(_env).GetAt(1).FP.Get_scope(_env) == self.moduleScope{
            return 
        }
        for _, _symbolInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
            symbolInfo := _symbolInfo
            if Lns_isCondTrue( symbolInfo.FP.Get_posForModToRef(_env)){
                self.FP.Writeln(_env, _env.GetVM().String_format("%s = None", Lns_2DDD(self.FP.getSymbolSym(_env, symbolInfo))))
            }
        }
    }
    if node.FP.Get_unwrapFlag(_env){
        {
            _expList, _unwrapBlock := node.FP.Get_expList(_env), node.FP.Get_unwrapBlock(_env)
            if !Lns_IsNil( _expList ) && !Lns_IsNil( _unwrapBlock ) {
                expList := _expList.(*Nodes_ExpListNode)
                unwrapBlock := _unwrapBlock.(*Nodes_BlockNode)
                if node.FP.Get_symbolInfoList(_env).GetAt(1).FP.Get_scope(_env) != self.moduleScope{
                    convPython_declVar(_env)
                }
                self.FP.Writeln(_env, "")
                self.FP.Writeln(_env, "{")
                self.FP.PushIndent(_env, nil)
                for _index, _varInfo := range( node.FP.Get_varList(_env).Items ) {
                    index := _index + 1
                    varInfo := _varInfo
                    if index != 1{
                        self.FP.WriteRaw(_env, ", ")
                    }
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("_%s", Lns_2DDD(convPython_normalizeSym_13_(_env, varInfo.FP.Get_name(_env).Txt))))
                }
                var tmpVarTypeList *LnsList2_[*Ast_TypeInfo]
                tmpVarTypeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
                for _index, _ := range( node.FP.Get_symbolInfoList(_env).Items ) {
                    index := _index + 1
                    tmpVarTypeList.Insert(expList.FP.GetExpTypeNoDDDAt(_env, index))
                }
                if convPython_getExpListKind_21_(_env, tmpVarTypeList, expList, self.option.FP.Get_addEnvArg(_env)) == convPython_ExpListKind__Direct{
                    {
                        var _forFrom0 LnsInt = tmpVarTypeList.Len() + 1
                        var _forTo0 LnsInt = expList.FP.Get_expTypeList(_env).Len()
                        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                            self.FP.WriteRaw(_env, ", _")
                        }
                    }
                }
                self.FP.WriteRaw(_env, " := ")
                self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, expList), tmpVarTypeList, expList, false)
                self.FP.Writeln(_env, "")
                self.FP.WriteRaw(_env, "if ")
                var hasCond bool
                hasCond = false
                for _index, _varInfo := range( node.FP.Get_varList(_env).Items ) {
                    index := _index + 1
                    varInfo := _varInfo
                    if expList.FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                        if hasCond{
                            self.FP.WriteRaw(_env, " || ")
                        }
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("_%s == %s", Lns_2DDD(convPython_normalizeSym_13_(_env, varInfo.FP.Get_name(_env).Txt), convPython_literalNil)))
                        hasCond = true
                    }
                }
                self.FP.Writeln(_env, "{")
                convPython_filter_7_(_env, &unwrapBlock.Nodes_Node, self, &node.Nodes_Node)
                {
                    _thenBlock := node.FP.Get_thenBlock(_env)
                    if !Lns_IsNil( _thenBlock ) {
                        thenBlock := _thenBlock.(*Nodes_BlockNode)
                        self.FP.Writeln(_env, "} else {")
                        self.FP.PushIndent(_env, nil)
                        for _index, _varInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
                            index := _index + 1
                            varInfo := _varInfo
                            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s = _%s", Lns_2DDD(self.FP.getSymbolSym(_env, varInfo), convPython_normalizeSym_13_(_env, varInfo.FP.Get_name(_env)))))
                            if expList.FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                                self.FP.outputAny2Type(_env, varInfo.FP.Get_typeInfo(_env))
                            }
                            self.FP.Writeln(_env, "")
                        }
                        self.FP.PopIndent(_env)
                        convPython_filter_7_(_env, &thenBlock.Nodes_Node, self, &node.Nodes_Node)
                        self.FP.Writeln(_env, "}")
                    } else {
                        self.FP.Writeln(_env, "} else {")
                        self.FP.PushIndent(_env, nil)
                        for _index, _varInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
                            index := _index + 1
                            varInfo := _varInfo
                            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s = _%s", Lns_2DDD(self.FP.getSymbolSym(_env, varInfo), convPython_normalizeSym_13_(_env, varInfo.FP.Get_name(_env)))))
                            if expList.FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                                self.FP.outputAny2Type(_env, varInfo.FP.Get_typeInfo(_env))
                            }
                            self.FP.Writeln(_env, "")
                        }
                        self.FP.PopIndent(_env)
                        self.FP.Writeln(_env, "}")
                    }
                }
                self.FP.PopIndent(_env)
                self.FP.Writeln(_env, "}")
            }
        }
    } else { 
        {
            _expList := node.FP.Get_expList(_env)
            if !Lns_IsNil( _expList ) {
                expList := _expList.(*Nodes_ExpListNode)
                if node.FP.Get_symbolInfoList(_env).Len() == 1{
                    var symbolInfo *Ast_SymbolInfo
                    symbolInfo = node.FP.Get_symbolInfoList(_env).GetAt(1)
                    self.FP.WriteRaw(_env, symbolInfo.FP.Get_name(_env))
                    self.FP.WriteRaw(_env, " = ")
                    self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, expList), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](symbolInfo.FP.Get_typeInfo(_env))), expList, false)
                    self.FP.Writeln(_env, "")
                } else { 
                    convPython_declVar(_env)
                    var varTypeList *LnsList2_[*Ast_TypeInfo]
                    varTypeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
                    for _index, _symbolInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
                        index := _index + 1
                        symbolInfo := _symbolInfo
                        varTypeList.Insert(symbolInfo.FP.Get_typeInfo(_env))
                        if index > 1{
                            self.FP.WriteRaw(_env, ",")
                        }
                        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( symbolInfo.FP.Get_scope(_env) == self.moduleScope) ||
                            _env.SetStackVal( symbolInfo.FP.Get_posForModToRef(_env)) ||
                            _env.SetStackVal( Ast_isPubToExternal(_env, symbolInfo.FP.Get_accessMode(_env))) )){
                            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s", Lns_2DDD(self.FP.getSymbolSym(_env, symbolInfo))))
                        } else { 
                            self.FP.WriteRaw(_env, "_")
                        }
                    }
                    if convPython_getExpListKind_21_(_env, varTypeList, expList, self.option.FP.Get_addEnvArg(_env)) == convPython_ExpListKind__Direct{
                        {
                            var _forFrom1 LnsInt = varTypeList.Len() + 1
                            var _forTo1 LnsInt = expList.FP.Get_expTypeList(_env).Len()
                            for _forWork1 := _forFrom1; _forWork1 <= _forTo1; _forWork1++ {
                                self.FP.WriteRaw(_env, ", _")
                            }
                        }
                    }
                    self.FP.WriteRaw(_env, " = ")
                    self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, expList), varTypeList, expList, false)
                    self.FP.Writeln(_env, "")
                }
            } else {
                convPython_declVar(_env)
            }
        }
    }
}
// 3091: decl @lune.@base.@convPython.convFilter.processDeclVar
func (self *convPython_convFilter) ProcessDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,_opt LnsAny) {
    if _switch0 := node.FP.Get_mode(_env); _switch0 == Nodes_DeclVarMode__Let {
        self.FP.outputLetVar(_env, node)
    } else if _switch0 == Nodes_DeclVarMode__Unwrap {
        self.FP.Writeln(_env, "{")
        self.FP.PushIndent(_env, nil)
        for _, _varSym := range( node.FP.Get_symbolInfoList(_env).Items ) {
            varSym := _varSym
            self.FP.Writeln(_env, _env.GetVM().String_format("var _%s LnsAny", Lns_2DDD(varSym.FP.Get_name(_env))))
        }
        var expList *Nodes_ExpListNode
        
        {
            _expList := node.FP.Get_expList(_env)
            if _expList == nil{
                Util_err(_env, "illegal")
            } else {
                expList = _expList.(*Nodes_ExpListNode)
            }
        }
        var convPython_setVals func(_env *LnsEnv)
        convPython_setVals = func(_env *LnsEnv) {
            for _index, _varSym := range( node.FP.Get_symbolInfoList(_env).Items ) {
                index := _index + 1
                varSym := _varSym
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s = _%s", Lns_2DDD(self.FP.getSymbolSym(_env, varSym), varSym.FP.Get_name(_env))))
                if expList.FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                    self.FP.outputAny2Type(_env, varSym.FP.Get_typeInfo(_env))
                }
                self.FP.Writeln(_env, "")
            }
        }
        var typeList *LnsList2_[*Ast_TypeInfo]
        typeList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
        for _index, _varSym := range( node.FP.Get_symbolInfoList(_env).Items ) {
            index := _index + 1
            varSym := _varSym
            typeList.Insert(varSym.FP.Get_typeInfo(_env))
            if index > 1{
                self.FP.WriteRaw(_env, ",")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("_%s", Lns_2DDD(varSym.FP.Get_name(_env))))
        }
        if convPython_getExpListKind_21_(_env, typeList, expList, self.option.FP.Get_addEnvArg(_env)) == convPython_ExpListKind__Direct{
            {
                var _forFrom0 LnsInt = node.FP.Get_symbolInfoList(_env).Len() + 1
                var _forTo0 LnsInt = expList.FP.Get_expTypeList(_env).Len()
                for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                    self.FP.WriteRaw(_env, ",_")
                }
            }
        }
        self.FP.WriteRaw(_env, " = ")
        self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, expList), typeList, expList, false)
        self.FP.Writeln(_env, "")
        self.FP.WriteRaw(_env, "if ")
        var hasCond bool
        hasCond = false
        for _index, _varSym := range( node.FP.Get_symbolInfoList(_env).Items ) {
            index := _index + 1
            varSym := _varSym
            if expList.FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                if hasCond{
                    self.FP.WriteRaw(_env, " || ")
                }
                self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_IsNil( _%s )", Lns_2DDD(varSym.FP.Get_name(_env))))
                hasCond = true
            }
        }
        self.FP.Writeln(_env, " {")
        convPython_filter_7_(_env, &Lns_unwrap( node.FP.Get_unwrapBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
        {
            _thenBlock := node.FP.Get_thenBlock(_env)
            if !Lns_IsNil( _thenBlock ) {
                thenBlock := _thenBlock.(*Nodes_BlockNode)
                self.FP.Writeln(_env, "} else {")
                self.FP.PushIndent(_env, nil)
                convPython_setVals(_env)
                self.FP.PopIndent(_env)
                convPython_filter_7_(_env, &thenBlock.Nodes_Node, self, &node.Nodes_Node)
                self.FP.Writeln(_env, "}")
            } else {
                self.FP.Writeln(_env, "}")
                convPython_setVals(_env)
            }
        }
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    }
}
// 3168: decl @lune.@base.@convPython.convFilter.processWhen
func (self *convPython_convFilter) ProcessWhen(_env *LnsEnv, node *Nodes_WhenNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "if ")
    for _index, _symPair := range( node.FP.Get_symPairList(_env).Items ) {
        index := _index + 1
        symPair := _symPair
        if index > 1{
            self.FP.WriteRaw(_env, " && ")
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s != %s", Lns_2DDD(self.FP.getSymbolSym(_env, symPair.FP.Get_src(_env)), convPython_literalNil)))
        symPair.FP.Get_dst(_env).FP.Set_convModuleParam(_env, true)
    }
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    for _, _symPair := range( node.FP.Get_symPairList(_env).Items ) {
        symPair := _symPair
        if Lns_isCondTrue( symPair.FP.Get_dst(_env).FP.Get_posForModToRef(_env)){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s_%d := %s", Lns_2DDD(symPair.FP.Get_dst(_env).FP.Get_name(_env), symPair.FP.Get_dst(_env).FP.Get_symbolId(_env), self.FP.getSymbolSym(_env, symPair.FP.Get_src(_env)))))
            self.FP.outputAny2Type(_env, symPair.FP.Get_dst(_env).FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
    }
    self.FP.PopIndent(_env)
    convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    {
        _elseBlock := node.FP.Get_elseBlock(_env)
        if !Lns_IsNil( _elseBlock ) {
            elseBlock := _elseBlock.(*Nodes_BlockNode)
            self.FP.Writeln(_env, "} else {")
            convPython_filter_7_(_env, &elseBlock.Nodes_Node, self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, "}")
        } else {
            self.FP.WriteRaw(_env, "}")
        }
    }
    self.FP.Writeln(_env, "")
}
// 3205: decl @lune.@base.@convPython.convFilter.processDeclArg
func (self *convPython_convFilter) ProcessDeclArg(_env *LnsEnv, node *Nodes_DeclArgNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s ", Lns_2DDD(self.FP.getSymbolSym(_env, node.FP.Get_symbolInfo(_env)))))
}
// 3212: decl @lune.@base.@convPython.convFilter.processDeclArgDDD
func (self *convPython_convFilter) ProcessDeclArgDDD(_env *LnsEnv, node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "*ddd")
}
// 3218: decl @lune.@base.@convPython.convFilter.processExpSubDDD
func (self *convPython_convFilter) ProcessExpSubDDD(_env *LnsEnv, node *Nodes_ExpSubDDDNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processExpSubDDD"
    Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
}
// 3225: decl @lune.@base.@convPython.convFilter.processDeclForm
func (self *convPython_convFilter) ProcessDeclForm(_env *LnsEnv, node *Nodes_DeclFormNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("type %s ", Lns_2DDD(self.FP.getFuncSymbol(_env, node.FP.Get_expType(_env)))))
    self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convPython_FuncInfo__Type{node.FP.Get_expType(_env)})
    self.FP.Writeln(_env, "")
}
// 3234: decl @lune.@base.@convPython.convFilter.processDeclFunc
func (self *convPython_convFilter) ProcessDeclFunc(_env *LnsEnv, node *Nodes_DeclFuncNode,_opt LnsAny) {
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_expType(_env)
    if (self.processMode == convPython_ProcessMode__NonClosureFuncDecl) == convPython_isClosure_9_(_env, node.FP.Get_expType(_env)){
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.processMode != convPython_ProcessMode__NonClosureFuncDecl) &&
            _env.SetStackVal( Lns_op_not(node.FP.Get_declInfo(_env).FP.Get_symbol(_env))) ).(bool)){
            self.FP.WriteRaw(_env, self.FP.getFuncSymbol(_env, funcType))
        }
        return 
    }
    self.FP.OutputDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env))
}
// 3267: decl @lune.@base.@convPython.convFilter.processRefType
func (self *convPython_convFilter) ProcessRefType(_env *LnsEnv, node *Nodes_RefTypeNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, self.FP.type2gotype(_env, node.FP.Get_expType(_env)))
}
// 3289: decl @lune.@base.@convPython.convFilter.processCond
func (self *convPython_convFilter) processCond(_env *LnsEnv, node *Nodes_Node,parent *Nodes_Node) {
    if node.FP.Get_expType(_env).FP.Get_nonnilableType(_env) != Ast_builtinTypeBool{
        self.FP.WriteRaw(_env, "Lns_isCondTrue( ")
        convPython_filter_7_(_env, node, self, parent)
        self.FP.WriteRaw(_env, ")")
    } else { 
        convPython_filter_7_(_env, node, self, parent)
    }
}
// 3300: decl @lune.@base.@convPython.convFilter.processIf
func (self *convPython_convFilter) ProcessIf(_env *LnsEnv, node *Nodes_IfNode,_opt LnsAny) {
    for _, _stmt := range( node.FP.Get_stmtList(_env).Items ) {
        stmt := _stmt
        if _switch0 := stmt.FP.Get_kind(_env); _switch0 == Nodes_IfKind__If {
            self.FP.WriteRaw(_env, "if ")
            self.FP.processCond(_env, stmt.FP.Get_exp(_env), &node.Nodes_Node)
            self.FP.Writeln(_env, ":")
            convPython_filter_7_(_env, &stmt.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        } else if _switch0 == Nodes_IfKind__ElseIf {
            self.FP.WriteRaw(_env, "else if ")
            self.FP.processCond(_env, stmt.FP.Get_exp(_env), &node.Nodes_Node)
            self.FP.Writeln(_env, ":")
            convPython_filter_7_(_env, &stmt.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        } else if _switch0 == Nodes_IfKind__Else {
            self.FP.Writeln(_env, "else:")
            convPython_filter_7_(_env, &stmt.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        }
    }
}
// 3325: decl @lune.@base.@convPython.convFilter.processSwitch
func (self *convPython_convFilter) ProcessSwitch(_env *LnsEnv, node *Nodes_SwitchNode,_opt LnsAny) {
    var nodeIndex LnsInt
    nodeIndex = node.FP.Get_idInNS(_env)
    var valName string
    valName = _env.GetVM().String_format("_switch%d", Lns_2DDD(nodeIndex))
    self.FP.WriteRaw(_env, _env.GetVM().String_format("if %s := ", Lns_2DDD(valName)))
    convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, "; ")
    for _caseIndex, _caseNode := range( node.FP.Get_caseList(_env).Items ) {
        caseIndex := _caseIndex + 1
        caseNode := _caseNode
        if caseIndex != 1{
            self.FP.WriteRaw(_env, "} else if ")
        }
        for _index, _exp := range( caseNode.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
            index := _index + 1
            exp := _exp
            if index != 1{
                self.FP.WriteRaw(_env, " || ")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s == ", Lns_2DDD(valName)))
            convPython_filter_7_(_env, exp, self, &caseNode.FP.Get_expList(_env).Nodes_Node)
        }
        self.FP.Writeln(_env, " {")
        convPython_filter_7_(_env, &caseNode.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    }
    {
        _defaultBlock := node.FP.Get_default(_env)
        if !Lns_IsNil( _defaultBlock ) {
            defaultBlock := _defaultBlock.(*Nodes_BlockNode)
            self.FP.Writeln(_env, "} else {")
            convPython_filter_7_(_env, &defaultBlock.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln(_env, "}")
}
// 3360: decl @lune.@base.@convPython.convFilter.processMatch
func (self *convPython_convFilter) ProcessMatch(_env *LnsEnv, node *Nodes_MatchNode,_opt LnsAny) {
    var convPython_hasAccessing func(_env *LnsEnv) bool
    convPython_hasAccessing = func(_env *LnsEnv) bool {
        for _, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
            caseInfo := _caseInfo
            for _, _symbol := range( caseInfo.FP.Get_valParamNameList(_env).Items ) {
                symbol := _symbol
                if Lns_isCondTrue( symbol.FP.Get_posForModToRef(_env)){
                    return true
                }
            }
        }
        return false
    }
    var val string
    var nodeIndex LnsInt
    nodeIndex = node.FP.Get_idInNS(_env)
    if convPython_hasAccessing(_env){
        val = _env.GetVM().String_format("_matchExp%d", Lns_2DDD(nodeIndex))
        self.FP.WriteRaw(_env, _env.GetVM().String_format("switch %s := ", Lns_2DDD(val)))
    } else { 
        val = ""
        self.FP.WriteRaw(_env, "switch ")
    }
    convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
    self.FP.Writeln(_env, ".(type) {")
    for _, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
        caseInfo := _caseInfo
        self.FP.Writeln(_env, _env.GetVM().String_format("case *%s:", Lns_2DDD(self.FP.getAlgeSymbol(_env, caseInfo.FP.Get_valInfo(_env)))))
        for _index, _symbol := range( caseInfo.FP.Get_valParamNameList(_env).Items ) {
            index := _index + 1
            symbol := _symbol
            if Lns_isCondTrue( symbol.FP.Get_posForModToRef(_env)){
                self.FP.Writeln(_env, _env.GetVM().String_format("%s := %s.Val%d", Lns_2DDD(self.FP.getSymbolSym(_env, symbol), val, index)))
            }
        }
        convPython_filter_7_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    }
    {
        _defaultBlock := node.FP.Get_defaultBlock(_env)
        if !Lns_IsNil( _defaultBlock ) {
            defaultBlock := _defaultBlock.(*Nodes_Node)
            self.FP.Writeln(_env, "default:")
            convPython_filter_7_(_env, defaultBlock, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln(_env, "}")
}
// 3402: decl @lune.@base.@convPython.convFilter.processWhile
func (self *convPython_convFilter) ProcessWhile(_env *LnsEnv, node *Nodes_WhileNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "for ")
    if Lns_op_not(node.FP.Get_infinit(_env)){
        self.FP.processCond(_env, node.FP.Get_exp(_env), &node.Nodes_Node)
    }
    self.FP.Writeln(_env, " {")
    convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "}")
}
// 3416: decl @lune.@base.@convPython.convFilter.processRepeat
func (self *convPython_convFilter) ProcessRepeat(_env *LnsEnv, node *Nodes_RepeatNode,_opt LnsAny) {
    self.FP.Writeln(_env, "for {")
    convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.PushIndent(_env, nil)
    self.FP.WriteRaw(_env, "if ")
    self.FP.processCond(_env, node.FP.Get_exp(_env), &node.Nodes_Node)
    self.FP.Writeln(_env, "{ break }")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 3434: decl @lune.@base.@convPython.convFilter.processFor
func (self *convPython_convFilter) ProcessFor(_env *LnsEnv, node *Nodes_ForNode,_opt LnsAny) {
    self.FP.Write(_env, _env.GetVM().String_format("for %s in range( ", Lns_2DDD(node.FP.Get_val(_env).FP.Get_name(_env))))
    convPython_filter_7_(_env, node.FP.Get_init(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, ", ")
    convPython_filter_7_(_env, node.FP.Get_to(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, " + 1")
    {
        _delta := node.FP.Get_delta(_env)
        if !Lns_IsNil( _delta ) {
            delta := _delta.(*Nodes_Node)
            self.FP.Write(_env, ", ")
            convPython_filter_7_(_env, delta, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln(_env, "):")
    convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
}
// 3481: decl @lune.@base.@convPython.convFilter.processApply
func (self *convPython_convFilter) ProcessApply(_env *LnsEnv, node *Nodes_ApplyNode,_opt LnsAny) {
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    var nodeIndex LnsInt
    nodeIndex = node.FP.Get_idInNS(_env)
    var formSym string
    formSym = _env.GetVM().String_format("_applyForm%d", Lns_2DDD(nodeIndex))
    var paramSym string
    paramSym = _env.GetVM().String_format("_applyParam%d", Lns_2DDD(nodeIndex))
    var prevSym string
    prevSym = _env.GetVM().String_format("_applyPrev%d", Lns_2DDD(nodeIndex))
    if Lns_isCondTrue( node.FP.Get_expList(_env).FP.Get_mRetExp(_env)){
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s, %s, %s := ", Lns_2DDD(formSym, paramSym, prevSym)))
        convPython_filter_7_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, &node.Nodes_Node)
        self.FP.Writeln(_env, "")
    } else { 
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s, %s, %s := ", Lns_2DDD(formSym, paramSym, prevSym)))
        convPython_filter_7_(_env, node.FP.Get_expList(_env).FP.Get_expList(_env).GetAt(1), self, &node.FP.Get_expList(_env).Nodes_Node)
        self.FP.WriteRaw(_env, ",")
        convPython_filter_7_(_env, node.FP.Get_expList(_env).FP.Get_expList(_env).GetAt(2), self, &node.FP.Get_expList(_env).Nodes_Node)
        self.FP.WriteRaw(_env, ", LnsAny(")
        convPython_filter_7_(_env, node.FP.Get_expList(_env).FP.Get_expList(_env).GetAt(3), self, &node.FP.Get_expList(_env).Nodes_Node)
        self.FP.Writeln(_env, ")")
    }
    self.FP.Writeln(_env, "for {")
    self.FP.PushIndent(_env, nil)
    var setTxt string
    setTxt = prevSym
    {
        var _forFrom0 LnsInt = 2
        var _forTo0 LnsInt = node.FP.Get_varList(_env).Len()
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            index := _forWork0
            var symInfo *Ast_SymbolInfo
            symInfo = node.FP.Get_varList(_env).GetAt(index)
            self.FP.Writeln(_env, _env.GetVM().String_format("var %s %s", Lns_2DDD(self.FP.getSymbolSym(_env, symInfo), self.FP.type2gotype(_env, symInfo.FP.Get_typeInfo(_env)))))
            setTxt = _env.GetVM().String_format("%s, %s", Lns_2DDD(setTxt, self.FP.getSymbolSym(_env, symInfo)))
        }
    }
    if node.FP.Get_expList(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        var workSym string
        workSym = _env.GetVM().String_format("_applyWork%d", Lns_2DDD(nodeIndex))
        self.FP.Writeln(_env, _env.GetVM().String_format("%s := %s.(*Lns_luaValue).Call( Lns_2DDD( %s, %s ) )", Lns_2DDD(workSym, formSym, paramSym, prevSym)))
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s = ", Lns_2DDD(setTxt)))
        for _index, _ := range( node.FP.Get_varList(_env).Items ) {
            index := _index + 1
            if index > 1{
                self.FP.WriteRaw(_env, ",")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_getFromMulti(%s,%d)", Lns_2DDD(workSym, index - 1)))
        }
        self.FP.Writeln(_env, "")
    } else { 
        self.FP.Writeln(_env, _env.GetVM().String_format("%s = %s( %s %s, %s )", Lns_2DDD(setTxt, formSym, convPython_getAddEnvArg_6_(_env, 2, self.option.FP.Get_addEnvArg(_env)), paramSym, prevSym)))
    }
    self.FP.Writeln(_env, _env.GetVM().String_format("if Lns_IsNil( %s ) { break }", Lns_2DDD(prevSym)))
    var topSymInfo *Ast_SymbolInfo
    topSymInfo = node.FP.Get_varList(_env).GetAt(1)
    if topSymInfo.FP.Get_name(_env) != "_"{
        self.FP.Writeln(_env, _env.GetVM().String_format("%s := %s.(%s)", Lns_2DDD(self.FP.getSymbolSym(_env, topSymInfo), prevSym, self.FP.type2gotype(_env, topSymInfo.FP.Get_typeInfo(_env)))))
    }
    self.FP.PopIndent(_env)
    convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "}")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 3557: decl @lune.@base.@convPython.convFilter.outputForeachLua
func (self *convPython_convFilter) outputForeachLua(_env *LnsEnv, node *Nodes_Node,sortFlag bool,exp *Nodes_Node,val *Ast_SymbolInfo,key LnsAny,block *Nodes_BlockNode) {
    __func__ := "@lune.@base.@convPython.convFilter.outputForeachLua"
    var nodeIndex LnsInt
    {
        __exp := Nodes_ForeachNodeDownCastF(node.FP)
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Nodes_ForeachNode)
            nodeIndex = _exp.FP.Get_idInNS(_env)
        } else {
            {
                __exp := Nodes_ForsortNodeDownCastF(node.FP)
                if !Lns_IsNil( __exp ) {
                    _exp := __exp.(*Nodes_ForsortNode)
                    nodeIndex = _exp.FP.Get_idInNS(_env)
                } else {
                    Util_err(_env, _env.GetVM().String_format("illegal node -- %s", Lns_2DDD(Nodes_getNodeKindName(_env, node.FP.Get_kind(_env)))))
                }
            }
        }
    }
    if _switch0 := exp.FP.Get_expType(_env).FP.Get_extedType(_env).FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Map {
        self.FP.Writeln(_env, "{")
        self.FP.PushIndent(_env, nil)
        var tmpExp string
        tmpExp = _env.GetVM().String_format("_foreachExp%d", Lns_2DDD(nodeIndex))
        var tmpKey string
        tmpKey = _env.GetVM().String_format("_foreachKey%d", Lns_2DDD(nodeIndex))
        var tmpVal string
        tmpVal = _env.GetVM().String_format("_foreachVal%d", Lns_2DDD(nodeIndex))
        var tmpIndex string
        tmpIndex = _env.GetVM().String_format("_foreachIndex%d", Lns_2DDD(nodeIndex))
        var sorted string
        sorted = _env.GetVM().String_format("_foreachSorted%d", Lns_2DDD(nodeIndex))
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := ", Lns_2DDD(tmpExp)))
        convPython_filter_7_(_env, exp, self, node)
        self.FP.Writeln(_env, "")
        var tmpValName string
        if val.FP.HasAccess(_env){
            tmpValName = tmpVal
        } else { 
            tmpValName = "_"
        }
        if sortFlag{
            self.FP.Writeln(_env, _env.GetVM().String_format("%s := %s.SortMapKeyList( %s )", Lns_2DDD(sorted, self.env.FP.getCommonVm(_env), tmpExp)))
            self.FP.Writeln(_env, _env.GetVM().String_format("%s, %s := %s.Get1stFromMap()", Lns_2DDD(tmpIndex, tmpKey, sorted)))
            self.FP.Writeln(_env, _env.GetVM().String_format("for %s != %s {", Lns_2DDD(tmpIndex, convPython_literalNil)))
            self.FP.PushIndent(_env, nil)
        } else { 
            self.FP.Writeln(_env, _env.GetVM().String_format("%s, %s := %s.Get1stFromMap()", Lns_2DDD(tmpKey, tmpValName, tmpExp)))
            self.FP.Writeln(_env, _env.GetVM().String_format("for %s != %s {", Lns_2DDD(tmpKey, convPython_literalNil)))
            self.FP.PushIndent(_env, nil)
        }
        {
            _keySym := key
            if !Lns_IsNil( _keySym ) {
                keySym := _keySym.(*Ast_SymbolInfo)
                if keySym.FP.HasAccess(_env){
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := %s", Lns_2DDD(self.FP.getSymbolSym(_env, keySym), tmpKey)))
                    self.FP.outputAny2Type(_env, keySym.FP.Get_typeInfo(_env))
                    self.FP.Writeln(_env, "")
                }
            }
        }
        if val.FP.HasAccess(_env){
            if sortFlag{
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := %s.GetAt( %s )", Lns_2DDD(self.FP.getSymbolSym(_env, val), tmpExp, tmpKey)))
            } else { 
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := %s", Lns_2DDD(self.FP.getSymbolSym(_env, val), tmpVal)))
            }
            self.FP.outputAny2Type(_env, val.FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
        self.FP.PopIndent(_env)
        convPython_filter_7_(_env, &block.Nodes_Node, self, node)
        self.FP.PushIndent(_env, nil)
        if Lns_op_not(sortFlag){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s, %s = ", Lns_2DDD(tmpKey, tmpValName)))
            self.FP.Writeln(_env, _env.GetVM().String_format("%s.NextFromMap( %s )", Lns_2DDD(tmpExp, tmpKey)))
        } else { 
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s, %s = ", Lns_2DDD(tmpIndex, tmpKey)))
            self.FP.Writeln(_env, _env.GetVM().String_format("%s.NextFromMap( %s )", Lns_2DDD(sorted, tmpIndex)))
        }
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    } else {
        Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
    }
}
// 3650: decl @lune.@base.@convPython.convFilter.processForeach
func (self *convPython_convFilter) ProcessForeach(_env *LnsEnv, node *Nodes_ForeachNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processForeach"
    if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_srcTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        self.FP.outputForeachLua(_env, &node.Nodes_Node, false, node.FP.Get_exp(_env), node.FP.Get_val(_env), node.FP.Get_key(_env), node.FP.Get_block(_env))
        return 
    }
    var hasAccessLoopSym LnsAny
    hasAccessLoopSym = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(node.FP.Get_key(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_SymbolInfo).FP.Get_posForModToRef(_env)}))) ||
        _env.SetStackVal( node.FP.Get_val(_env).FP.Get_posForModToRef(_env)) )
    self.FP.WriteRaw(_env, "for ")
    var loopExpType *Ast_TypeInfo
    loopExpType = node.FP.Get_exp(_env).FP.Get_expType(_env)
    if _switch0 := loopExpType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        var valName string
        valName = self.FP.getSymbolSym(_env, node.FP.Get_val(_env))
        _ = loopExpType.FP.Get_itemTypeInfoList(_env).GetAt(1)
        if Lns_isCondTrue( hasAccessLoopSym){
            {
                _key := node.FP.Get_key(_env)
                if !Lns_IsNil( _key ) {
                    key := _key.(*Ast_SymbolInfo)
                    if key.FP.Get_name(_env) != "_"{
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("_%s", Lns_2DDD(key.FP.Get_name(_env))))
                    } else { 
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s", Lns_2DDD(key.FP.Get_name(_env))))
                    }
                    
                    self.FP.WriteRaw(_env, ", ")
                }
            }
            self.FP.WriteRaw(_env, valName)
            self.FP.WriteRaw(_env, " in ")
        }
        convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.Writeln(_env, ":")
        self.FP.PushIndent(_env, nil)
        {
            _key := node.FP.Get_key(_env)
            if !Lns_IsNil( _key ) {
                key := _key.(*Ast_SymbolInfo)
                if Lns_isCondTrue( key.FP.Get_posForModToRef(_env)){
                    self.FP.Writeln(_env, _env.GetVM().String_format("%s = _%s + 1", Lns_2DDD(self.FP.getSymbolSym(_env, key), key.FP.Get_name(_env))))
                }
            }
        }
        self.FP.PopIndent(_env)
    } else if _switch0 == Ast_TypeInfoKind__Map {
        var valName string
        valName = self.FP.getSymbolSym(_env, node.FP.Get_val(_env))
        var itemType *Ast_TypeInfo
        itemType = loopExpType.FP.Get_itemTypeInfoList(_env).GetAt(2)
        var keyType *Ast_TypeInfo
        keyType = loopExpType.FP.Get_itemTypeInfoList(_env).GetAt(1)
        if Lns_isCondTrue( hasAccessLoopSym){
            {
                _key := node.FP.Get_key(_env)
                if !Lns_IsNil( _key ) {
                    key := _key.(*Ast_SymbolInfo)
                    if key.FP.Get_name(_env) != "_"{
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("_%s", Lns_2DDD(key.FP.Get_name(_env))))
                    } else { 
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s", Lns_2DDD(key.FP.Get_name(_env))))
                    }
                    
                    self.FP.WriteRaw(_env, ", ")
                } else {
                    self.FP.WriteRaw(_env, "_, ")
                }
            }
            if valName != "_"{
                self.FP.WriteRaw(_env, _env.GetVM().String_format("_%s", Lns_2DDD(valName)))
            } else { 
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s", Lns_2DDD(valName)))
            }
            
            self.FP.WriteRaw(_env, " := ")
        }
        self.FP.WriteRaw(_env, "range( ")
        convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.Writeln(_env, ".Items ) {")
        self.FP.PushIndent(_env, nil)
        {
            _key := node.FP.Get_key(_env)
            if !Lns_IsNil( _key ) {
                key := _key.(*Ast_SymbolInfo)
                if key.FP.Get_name(_env) != "_"{
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := _%s", Lns_2DDD(key.FP.Get_name(_env), key.FP.Get_name(_env))))
                    self.FP.outputStem2Type(_env, keyType)
                    self.FP.Writeln(_env, "")
                }
                
            }
        }
        if valName != "_"{
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := _%s", Lns_2DDD(valName, valName)))
            self.FP.outputStem2Type(_env, itemType)
            self.FP.Writeln(_env, "")
        }
        
        self.FP.PopIndent(_env)
    } else if _switch0 == Ast_TypeInfoKind__Set {
        var valType *Ast_TypeInfo
        valType = loopExpType.FP.Get_itemTypeInfoList(_env).GetAt(1)
        var valName string
        valName = self.FP.getSymbolSym(_env, node.FP.Get_val(_env))
        if Lns_isCondTrue( hasAccessLoopSym){
            if valName != "_"{
                self.FP.WriteRaw(_env, _env.GetVM().String_format("_%s", Lns_2DDD(valName)))
            } else { 
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s", Lns_2DDD(valName)))
            }
            
            self.FP.WriteRaw(_env, " := ")
        }
        self.FP.WriteRaw(_env, "range( ")
        convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.Writeln(_env, ".Items ) {")
        self.FP.PushIndent(_env, nil)
        if valName != "_"{
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := _%s", Lns_2DDD(valName, valName)))
            self.FP.outputStem2Type(_env, valType)
            self.FP.Writeln(_env, "")
        }
        
        self.FP.PopIndent(_env)
    } else {
        Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
    }
    convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "")
}
// 3762: decl @lune.@base.@convPython.convFilter.processForsort
func (self *convPython_convFilter) ProcessForsort(_env *LnsEnv, node *Nodes_ForsortNode,_opt LnsAny) {
    if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_srcTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        self.FP.outputForeachLua(_env, &node.Nodes_Node, true, node.FP.Get_exp(_env), node.FP.Get_val(_env), node.FP.Get_key(_env), node.FP.Get_block(_env))
        return 
    }
    var keySym LnsAny
    var valSym LnsAny
    var keyTypeInfo *Ast_TypeInfo
    keyTypeInfo = node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env).GetAt(1)
    if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Set{
        keySym = node.FP.Get_val(_env)
        valSym = node.FP.Get_key(_env)
    } else { 
        keySym = node.FP.Get_key(_env)
        valSym = node.FP.Get_val(_env)
    }
    var nodeIndex LnsInt
    nodeIndex = node.FP.Get_idInNS(_env)
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    var collSym string
    collSym = _env.GetVM().String_format("__forsortCollection%d", Lns_2DDD(nodeIndex))
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := ", Lns_2DDD(collSym)))
    convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Writeln(_env, "")
    var sortSym string
    sortSym = _env.GetVM().String_format("__forsortSorted%d", Lns_2DDD(nodeIndex))
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := %s.", Lns_2DDD(sortSym, collSym)))
    if _switch0 := keyTypeInfo; _switch0 == Ast_builtinTypeInt {
        self.FP.Writeln(_env, "CreateKeyListInt()")
    } else if _switch0 == Ast_builtinTypeReal {
        self.FP.Writeln(_env, "CreateKeyListReal()")
    } else if _switch0 == Ast_builtinTypeString {
        self.FP.Writeln(_env, "CreateKeyListStr()")
    } else {
        self.FP.Writeln(_env, "CreateKeyListStem()")
    }
    self.FP.Writeln(_env, _env.GetVM().String_format("%s.Sort( %s%s, nil )", Lns_2DDD(sortSym, convPython_getAddEnvArg_6_(_env, 2, self.option.FP.Get_addEnvArg(_env)), convPython_getLnsItemKind_28_(_env, keyTypeInfo))))
    self.FP.WriteRaw(_env, "for _, ")
    var key string
    key = _env.GetVM().String_format("__forsortKey%d", Lns_2DDD(nodeIndex))
    if keySym != nil{
        keySym_95 := keySym.(*Ast_SymbolInfo)
        key = _env.GetVM().String_format("%s", Lns_2DDD(self.FP.getSymbolSym(_env, keySym_95)))
    }
    self.FP.WriteRaw(_env, _env.GetVM().String_format("_%s", Lns_2DDD(key)))
    self.FP.Writeln(_env, _env.GetVM().String_format(" := range( %s.Items ) {", Lns_2DDD(sortSym)))
    self.FP.PushIndent(_env, nil)
    if valSym != nil{
        valSym_97 := valSym.(*Ast_SymbolInfo)
        if Lns_isCondTrue( valSym_97.FP.Get_posForModToRef(_env)){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := %s.Items[ _%s ]", Lns_2DDD(self.FP.getSymbolSym(_env, valSym_97), collSym, key)))
            self.FP.outputStem2Type(_env, valSym_97.FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
    }
    if keySym != nil{
        keySym_100 := keySym.(*Ast_SymbolInfo)
        if Lns_isCondTrue( keySym_100.FP.Get_posForModToRef(_env)){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := _%s", Lns_2DDD(key, key)))
            self.FP.outputStem2Type(_env, keySym_100.FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
    }
    self.FP.PopIndent(_env)
    convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "}")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 3847: decl @lune.@base.@convPython.convFilter.processExpUnwrap
func (self *convPython_convFilter) ProcessExpUnwrap(_env *LnsEnv, node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "_lnsUnwrap( ")
    convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, ", ")
    {
        _def := node.FP.Get_default(_env)
        if !Lns_IsNil( _def ) {
            def := _def.(*Nodes_Node)
            convPython_filter_7_(_env, def, self, &node.Nodes_Node)
        } else {
            self.FP.WriteRaw(_env, "None")
        }
    }
    self.FP.WriteRaw(_env, ")")
}
// 3861: decl @lune.@base.@convPython.convFilter.processExpToDDD
func (self *convPython_convFilter) ProcessExpToDDD(_env *LnsEnv, node *Nodes_ExpToDDDNode,_opt LnsAny) {
    if Lns_isCondTrue( node.FP.Get_expList(_env).FP.Get_mRetExp(_env)){
        convPython_filter_7_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, &node.Nodes_Node)
    } else { 
        convPython_filter_7_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, &node.Nodes_Node)
    }
}
// 3873: decl @lune.@base.@convPython.convFilter.processExpNew
func (self *convPython_convFilter) ProcessExpNew(_env *LnsEnv, node *Nodes_ExpNewNode,_opt LnsAny) {
    var className string
    className = self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env))
    if self.FP.isSamePackage(_env, node.FP.Get_expType(_env).FP.GetModule(_env)){
        self.FP.WriteRaw(_env, _env.GetVM().String_format("New%s(", Lns_2DDD(className)))
    } else { 
        {
            _refTypeNode := Nodes_RefTypeNodeDownCastF(node.FP.Get_symbol(_env).FP)
            if !Lns_IsNil( _refTypeNode ) {
                refTypeNode := _refTypeNode.(*Nodes_RefTypeNode)
                {
                    _refNode := Nodes_RefFieldNodeDownCastF(refTypeNode.FP.Get_typeNode(_env).FP)
                    if !Lns_IsNil( _refNode ) {
                        refNode := _refNode.(*Nodes_RefFieldNode)
                        convPython_filter_7_(_env, refNode.FP.Get_prefix(_env), self, &node.Nodes_Node)
                        self.FP.WriteRaw(_env, ".")
                    }
                }
            }
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format("New%s(", Lns_2DDD(className)))
    }
    {
        _argList := node.FP.Get_argList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            var scope *Ast_Scope
            scope = Lns_unwrap( node.FP.Get_expType(_env).FP.Get_scope(_env)).(*Ast_Scope)
            var initFuncType *Ast_TypeInfo
            initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField(_env, "__init", true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
            self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, argList), initFuncType.FP.Get_argTypeInfoList(_env), argList, self.option.FP.Get_addEnvArg(_env))
        } else {
            self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, 0, self.option.FP.Get_addEnvArg(_env)))
        }
    }
    self.FP.WriteRaw(_env, ")")
}
// 3901: decl @lune.@base.@convPython.convFilter.outputIFMethods
func (self *convPython_convFilter) outputIFMethods(_env *LnsEnv, node *Nodes_DeclClassNode) {
    self.FP.PushIndent(_env, nil)
    if node.FP.Get_expType(_env).FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeRunner, nil){
        self.FP.Writeln(_env, "GetLnsSyncFlag() *Lns_syncFlag")
    }
    var name2MtdType *LnsMap2_[string,*Ast_TypeInfo]
    name2MtdType = NewLnsMap2_[string,*Ast_TypeInfo]( map[string]*Ast_TypeInfo{})
    var scope *Ast_Scope
    scope = Lns_unwrap( node.FP.Get_expType(_env).FP.Get_scope(_env)).(*Ast_Scope)
    scope.FP.FilterTypeInfoField(_env, true, scope, Ast_ScopeAccess__Normal, Ast_filterForm(func(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) bool {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) &&
            _env.SetStackVal( symbolInfo.FP.Get_name(_env) != "__init") &&
            _env.SetStackVal( Lns_op_not(symbolInfo.FP.Get_staticFlag(_env))) ).(bool)){
            name2MtdType.Set(symbolInfo.FP.Get_name(_env),symbolInfo.FP.Get_typeInfo(_env))
        }
        return true
    }))
    {
        __forsortCollection0 := name2MtdType
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, _name := range( __forsortSorted0.Items ) {
            typeInfo := __forsortCollection0.Items[ _name ]
            name := _name
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s(", Lns_2DDD(self.FP.getSymbol(_env, &convPython_SymbolKind__Func{typeInfo}, name))))
            if name != "_toMap"{
                self.FP.WriteRaw(_env, self.FP.getEnvArgDecl(_env, typeInfo.FP.Get_argTypeInfoList(_env).Len()))
            }
            for _index, _argType := range( typeInfo.FP.Get_argTypeInfoList(_env).Items ) {
                index := _index + 1
                argType := _argType
                if index != 1{
                    self.FP.WriteRaw(_env, ", ")
                }
                self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d %s", Lns_2DDD(index, self.FP.type2gotype(_env, argType))))
            }
            self.FP.WriteRaw(_env, ")")
            self.FP.outputRetType(_env, typeInfo.FP.Get_retTypeInfoList(_env))
            self.FP.Writeln(_env, "")
        }
    }
    self.FP.PopIndent(_env)
}
// 3943: decl @lune.@base.@convPython.convFilter.outputMethodIF
func (self *convPython_convFilter) outputMethodIF(_env *LnsEnv, node *Nodes_DeclClassNode) {
    self.FP.WriteRaw(_env, "type ")
    self.FP.WriteRaw(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
    self.FP.Writeln(_env, "Mtd interface {")
    self.FP.outputIFMethods(_env, node)
    self.FP.Writeln(_env, "}")
}
// 3953: decl @lune.@base.@convPython.convFilter.outputInterfaceType
func (self *convPython_convFilter) outputInterfaceType(_env *LnsEnv, node *Nodes_DeclClassNode) {
    self.FP.Writeln(_env, _env.GetVM().String_format("type %s interface {", Lns_2DDD(self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))))
    self.FP.PushIndent(_env, nil)
    self.FP.outputIFMethods(_env, node)
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
    var typeName string
    typeName = self.FP.type2gotype(_env, node.FP.Get_expType(_env))
    self.FP.Writeln(_env, _env.GetVM().String_format("func Lns_cast2%s( obj LnsAny ) LnsAny {", Lns_2DDD(typeName)))
    self.FP.Writeln(_env, _env.GetVM().String_format("    if _, ok := obj.(%s); ok { ", Lns_2DDD(typeName)))
    self.FP.Writeln(_env, "        return obj")
    self.FP.Writeln(_env, "    }")
    self.FP.Writeln(_env, _env.GetVM().String_format("    return %s", Lns_2DDD(convPython_literalNil)))
    self.FP.Writeln(_env, "}")
}
// 3976: decl @lune.@base.@convPython.convFilter.outputClassType
func (self *convPython_convFilter) outputClassType(_env *LnsEnv, node *Nodes_DeclClassNode,absImmutFlag bool) {
    self.FP.WriteRaw(_env, "type ")
    self.FP.WriteRaw(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
    self.FP.Writeln(_env, " struct {")
    self.FP.PushIndent(_env, nil)
    if node.FP.Get_expType(_env).FP.HasBase(_env){
        var superClassType *Ast_TypeInfo
        superClassType = node.FP.Get_expType(_env).FP.Get_baseTypeInfo(_env)
        self.FP.Writeln(_env, self.FP.getTypeSymbolWithPrefix(_env, superClassType))
    }
    var hasRunner bool
    hasRunner = self.FP.isImplementedRunner(_env, node.FP.Get_expType(_env))
    if hasRunner{
        self.FP.Writeln(_env, "_syncFlag *Lns_syncFlag")
    }
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode
        convPython_filter_7_(_env, &memberNode.Nodes_Node, self, &node.Nodes_Node)
    }
    if Lns_op_not(absImmutFlag){
        self.FP.WriteRaw(_env, "FP ")
        self.FP.WriteRaw(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
        self.FP.Writeln(_env, "Mtd")
    }
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
    if hasRunner{
        self.FP.Writeln(_env, _env.GetVM().String_format("func (self *%s) GetLnsSyncFlag() *Lns_syncFlag { return self._syncFlag }", Lns_2DDD(self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))))
    }
}
// 4014: decl @lune.@base.@convPython.convFilter.outputToStem
func (self *convPython_convFilter) outputToStem(_env *LnsEnv, node *Nodes_DeclClassNode,absImmutFlag bool) {
    self.FP.Writeln(_env, _env.GetVM().String_format("func %s2Stem( obj LnsAny ) LnsAny {", Lns_2DDD(self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_expType(_env)))))
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, _env.GetVM().String_format("if obj == %s {", Lns_2DDD(convPython_literalNil)))
    self.FP.Writeln(_env, _env.GetVM().String_format("    return %s", Lns_2DDD(convPython_literalNil)))
    self.FP.Writeln(_env, "}")
    self.FP.WriteRaw(_env, _env.GetVM().String_format("return obj.(%s)", Lns_2DDD(self.FP.type2gotype(_env, node.FP.Get_expType(_env)))))
    if Lns_op_not(absImmutFlag){
        self.FP.Writeln(_env, ".FP")
    } else { 
        self.FP.Writeln(_env, "")
    }
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 4032: decl @lune.@base.@convPython.convFilter.outputDownCast
func (self *convPython_convFilter) outputDownCast(_env *LnsEnv, node *Nodes_DeclClassNode) {
    var symbol string
    symbol = self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env))
    self.FP.WriteRaw(_env, "type ")
    self.FP.Writeln(_env, _env.GetVM().String_format("%sDownCast interface {", Lns_2DDD(symbol)))
    self.FP.PushIndent(_env, nil)
    self.FP.WriteRaw(_env, "To")
    self.FP.WriteRaw(_env, symbol)
    self.FP.WriteRaw(_env, "() ")
    self.FP.WriteRaw(_env, self.FP.type2gotype(_env, node.FP.Get_expType(_env)))
    self.FP.Writeln(_env, "")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
    self.FP.Writeln(_env, _env.GetVM().String_format("func %sDownCastF( multi ...LnsAny ) LnsAny {", Lns_2DDD(symbol)))
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, _env.GetVM().String_format("if len( multi ) == 0 { return %s }", Lns_2DDD(convPython_literalNil)))
    self.FP.Writeln(_env, "obj := multi[ 0 ]")
    self.FP.Writeln(_env, "if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }")
    self.FP.Writeln(_env, _env.GetVM().String_format("work, ok := obj.(%sDownCast)", Lns_2DDD(self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_expType(_env)))))
    self.FP.Writeln(_env, _env.GetVM().String_format("if ok { return work.To%s() }", Lns_2DDD(symbol)))
    self.FP.Writeln(_env, _env.GetVM().String_format("return %s", Lns_2DDD(convPython_literalNil)))
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 4058: decl @lune.@base.@convPython.convFilter.outputCastReceiver
func (self *convPython_convFilter) outputCastReceiver(_env *LnsEnv, node *Nodes_DeclClassNode) {
    var gotype string
    gotype = self.FP.type2gotype(_env, node.FP.Get_expType(_env))
    self.FP.WriteRaw(_env, "func (obj ")
    self.FP.WriteRaw(_env, gotype)
    self.FP.WriteRaw(_env, ") To")
    self.FP.WriteRaw(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
    self.FP.WriteRaw(_env, "() ")
    self.FP.WriteRaw(_env, gotype)
    self.FP.Writeln(_env, " {")
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, "return obj")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 4073: decl @lune.@base.@convPython.convFilter.outputNewSetup
func (self *convPython_convFilter) outputNewSetup(_env *LnsEnv, objName string,classType *Ast_TypeInfo,absImmutFlag bool) {
    var className string
    className = self.FP.getTypeSymbol(_env, classType)
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s := ", Lns_2DDD(objName)))
    if Lns_op_not(absImmutFlag){
        self.FP.WriteRaw(_env, "&")
    }
    self.FP.Writeln(_env, _env.GetVM().String_format("%s{}", Lns_2DDD(className)))
    if Lns_op_not(absImmutFlag){
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.FP = ", Lns_2DDD(objName)))
        self.FP.Writeln(_env, _env.GetVM().String_format("%s", Lns_2DDD(objName)))
    }
    {
        var workType *Ast_TypeInfo
        workType = classType
        for workType.FP.HasBase(_env) {
            workType = workType.FP.Get_baseTypeInfo(_env)
            var superName string
            superName = self.FP.getTypeSymbol(_env, workType)
            self.FP.Writeln(_env, _env.GetVM().String_format("%s.%s.FP = %s", Lns_2DDD(objName, superName, objName)))
        }
    }
}
// 4098: decl @lune.@base.@convPython.convFilter.outputConstructor
func (self *convPython_convFilter) outputConstructor(_env *LnsEnv, node *Nodes_DeclClassNode,absImmutFlag bool) {
    var scope *Ast_Scope
    scope = Lns_unwrap( node.FP.Get_expType(_env).FP.Get_scope(_env)).(*Ast_Scope)
    var initFuncType *Ast_TypeInfo
    initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField(_env, "__init", true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
    var className string
    className = self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env))
    var ctorName string
    ctorName = self.FP.getConstrSymbol(_env, node.FP.Get_expType(_env))
    var goType string
    goType = self.FP.type2gotype(_env, node.FP.Get_expType(_env))
    if Lns_op_not(node.FP.Get_expType(_env).FP.Get_abstractFlag(_env)){
        self.FP.WriteRaw(_env, _env.GetVM().String_format("func New%s(", Lns_2DDD(className)))
        self.FP.outputDeclFuncArg(_env, initFuncType)
        self.FP.Writeln(_env, _env.GetVM().String_format(") %s {", Lns_2DDD(goType)))
        self.FP.PushIndent(_env, nil)
        self.FP.outputNewSetup(_env, "obj", node.FP.Get_expType(_env), absImmutFlag)
        self.FP.WriteRaw(_env, _env.GetVM().String_format("obj.%s(", Lns_2DDD(ctorName)))
        self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, initFuncType.FP.Get_argTypeInfoList(_env).Len(), self.option.FP.Get_addEnvArg(_env)))
        for _index, _ := range( initFuncType.FP.Get_argTypeInfoList(_env).Items ) {
            index := _index + 1
            if index != 1{
                self.FP.WriteRaw(_env, ", ")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d", Lns_2DDD(index)))
        }
        self.FP.Writeln(_env, ")")
        self.FP.Writeln(_env, "return obj")
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    }
    if Lns_op_not(node.FP.HasUserInit(_env)){
        self.FP.WriteRaw(_env, _env.GetVM().String_format("func (self *%s) %s(", Lns_2DDD(className, ctorName)))
        self.FP.outputDeclFuncArg(_env, initFuncType)
        self.FP.Writeln(_env, ") {")
        self.FP.PushIndent(_env, nil)
        var superArgNum LnsInt
        if node.FP.Get_expType(_env).FP.HasBase(_env){
            var superType *Ast_TypeInfo
            superType = node.FP.Get_expType(_env).FP.Get_baseTypeInfo(_env)
            var baseScope *Ast_Scope
            baseScope = Lns_unwrap( superType.FP.Get_scope(_env)).(*Ast_Scope)
            var baseInitFuncType *Ast_TypeInfo
            baseInitFuncType = Lns_unwrap( baseScope.FP.GetTypeInfoField(_env, "__init", true, baseScope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
            var superName string
            superName = self.FP.getTypeSymbol(_env, superType)
            self.FP.WriteRaw(_env, _env.GetVM().String_format("self.%s.%s( ", Lns_2DDD(superName, self.FP.getConstrSymbol(_env, superType))))
            self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, baseInitFuncType.FP.Get_argTypeInfoList(_env).Len(), self.option.FP.Get_addEnvArg(_env)))
            {
                var _forFrom0 LnsInt = 1
                var _forTo0 LnsInt = baseInitFuncType.FP.Get_argTypeInfoList(_env).Len()
                for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                    index := _forWork0
                    if index != 1{
                        self.FP.WriteRaw(_env, ",")
                    }
                    if node.FP.Get_hasOldCtor(_env){
                        self.FP.WriteRaw(_env, convPython_literalNil)
                    } else { 
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d", Lns_2DDD(index)))
                    }
                }
            }
            self.FP.Writeln(_env, ")")
            if node.FP.Get_hasOldCtor(_env){
                superArgNum = 0
            } else { 
                superArgNum = baseInitFuncType.FP.Get_argTypeInfoList(_env).Len()
            }
        } else { 
            superArgNum = 0
        }
        if self.FP.isImplementedRunner(_env, node.FP.Get_expType(_env)){
            self.FP.Writeln(_env, "self._syncFlag = &Lns_syncFlag{}")
        }
        var argIndex LnsInt
        argIndex = superArgNum + 1
        for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
            memberNode := _memberNode
            if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
                self.FP.Writeln(_env, _env.GetVM().String_format("self.%s = arg%d", Lns_2DDD(self.FP.getSymbol(_env, &convPython_SymbolKind__Member{Ast_isPubToExternal(_env, memberNode.FP.Get_accessMode(_env))}, memberNode.FP.Get_name(_env).Txt), argIndex)))
                argIndex = argIndex + 1
            }
        }
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    }
}
// 4194: decl @lune.@base.@convPython.convFilter.outputAccessor
func (self *convPython_convFilter) OutputAccessor(_env *LnsEnv, node *Nodes_DeclClassNode) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType(_env)
    if classType.FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
        return 
    }
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode
        var memberNameToken *Types_Token
        memberNameToken = memberNode.FP.Get_name(_env)
        var memberName string
        memberName = memberNameToken.Txt
        var memberSym *Ast_SymbolInfo
        memberSym = memberNode.FP.Get_symbolInfo(_env)
        if memberNode.FP.Get_getterMode(_env) != Ast_AccessMode__None{
            var getterName string
            getterName = _env.GetVM().String_format("get_%s", Lns_2DDD(memberName))
            var getterSym *Ast_SymbolInfo
            getterSym = Lns_unwrap( _env.NilAccFin(_env.NilAccPush(classType.FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(_env, getterName)})/* 4208:33 */)).(*Ast_SymbolInfo)
            self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convPython_FuncInfo__Type{getterSym.FP.Get_typeInfo(_env)})
            var retType *Ast_TypeInfo
            retType = getterSym.FP.Get_typeInfo(_env).FP.Get_retTypeInfoList(_env).GetAt(1).FP.Get_srcTypeInfo(_env)
            self.FP.WriteRaw(_env, "{ return ")
            var prefix string
            if memberSym.FP.Get_staticFlag(_env){
                prefix = ""
            } else { 
                prefix = "self."
            }
            if retType.FP.Get_srcTypeInfo(_env) != memberSym.FP.Get_typeInfo(_env).FP.Get_srcTypeInfo(_env){
                if retType.FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
                    if Ast_isClass(_env, memberSym.FP.Get_typeInfo(_env).FP.Get_srcTypeInfo(_env)){
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s%s.FP", Lns_2DDD(prefix, self.FP.getSymbolSym(_env, memberSym))))
                    }
                } else if Ast_isClass(_env, retType){
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("&%s%s.%s", Lns_2DDD(prefix, self.FP.getSymbolSym(_env, memberSym), self.FP.getTypeSymbol(_env, retType))))
                } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( retType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                    _env.SetStackVal( retType.FP.HasBase(_env)) ).(bool)){
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s%s.FP", Lns_2DDD(prefix, self.FP.getSymbolSym(_env, memberSym))))
                    self.FP.outputStem2Type(_env, retType)
                } else { 
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( convPython_isAnyType_8_(_env, retType)) &&
                        _env.SetStackVal( Ast_isClass(_env, memberSym.FP.Get_typeInfo(_env))) ).(bool)){
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s%s.FP", Lns_2DDD(prefix, self.FP.getSymbolSym(_env, memberSym))))
                    } else { 
                        Util_err(_env, "not support")
                    }
                }
            } else { 
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s%s", Lns_2DDD(prefix, self.FP.getSymbolSym(_env, memberSym))))
            }
            self.FP.Writeln(_env, " }")
        }
        if memberNode.FP.Get_setterMode(_env) != Ast_AccessMode__None{
            var setterName string
            setterName = _env.GetVM().String_format("set_%s", Lns_2DDD(memberName))
            var setterSym *Ast_SymbolInfo
            setterSym = Lns_unwrap( _env.NilAccFin(_env.NilAccPush(classType.FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(_env, setterName)})/* 4249:33 */)).(*Ast_SymbolInfo)
            self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convPython_FuncInfo__Type{setterSym.FP.Get_typeInfo(_env)})
            self.FP.WriteRaw(_env, _env.GetVM().String_format("{ self.%s = arg1 ", Lns_2DDD(self.FP.getSymbolSym(_env, memberSym))))
            self.FP.Writeln(_env, "}")
        }
    }
}
// 4258: decl @lune.@base.@convPython.convFilter.outputStaticMember
func (self *convPython_convFilter) outputStaticMember(_env *LnsEnv, node *Nodes_DeclClassNode) {
    if self.processMode == convPython_ProcessMode__DeclClass{
        for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
            memberNode := _memberNode
            if memberNode.FP.Get_staticFlag(_env){
                self.FP.Writeln(_env, _env.GetVM().String_format("var %s %s", Lns_2DDD(self.FP.getSymbol(_env, &convPython_SymbolKind__Static{node.FP.Get_expType(_env)}, memberNode.FP.Get_name(_env).Txt), self.FP.type2gotype(_env, memberNode.FP.Get_expType(_env)))))
            }
        }
        {
            _initBlock := node.FP.Get_initBlock(_env).FP.Get_func(_env)
            if !Lns_IsNil( _initBlock ) {
                initBlock := _initBlock.(*Nodes_DeclMethodNode)
                convPython_filter_7_(_env, &initBlock.Nodes_Node, self, &node.Nodes_Node)
                self.FP.Writeln(_env, "")
            }
        }
    } else { 
        {
            _initBlock := node.FP.Get_initBlock(_env).FP.Get_func(_env)
            if !Lns_IsNil( _initBlock ) {
                initBlock := _initBlock.(*Nodes_DeclMethodNode)
                self.FP.Writeln(_env, _env.GetVM().String_format("%s(%s)", Lns_2DDD(self.FP.getFuncSymbol(_env, initBlock.FP.Get_expType(_env)), convPython_getAddEnvArg_6_(_env, 0, self.option.FP.Get_addEnvArg(_env)))))
            }
        }
    }
}
// 4290: decl @lune.@base.@convPython.convFilter.getFromStemName
func (self *convPython_convFilter) getFromStemName(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    __func__ := "@lune.@base.@convPython.convFilter.getFromStemName"
    var workTypeInfo *Ast_TypeInfo
    workTypeInfo = convPython_getOrgTypeInfo_15_(_env, typeInfo)
    {
        _name := convPython_type2FromStemNameMap.Get(workTypeInfo)
        if !Lns_IsNil( _name ) {
            name := _name.(string)
            return name
        }
    }
    if _switch0 := workTypeInfo.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        return "Lns_ToList"
    } else if _switch0 == Ast_TypeInfoKind__Set {
        return "Lns_ToSet"
    } else if _switch0 == Ast_TypeInfoKind__Map {
        return "Lns_ToLnsMap"
    } else if _switch0 == Ast_TypeInfoKind__Class {
        return _env.GetVM().String_format("%s_FromMap", Lns_2DDD(self.FP.getTypeSymbol(_env, workTypeInfo)))
    } else {
        Util_err(_env, _env.GetVM().String_format("%s: not support -- %s", Lns_2DDD(__func__, Ast_TypeInfoKind_getTxt( workTypeInfo.FP.Get_kind(_env)))))
    }
// insert a dummy
    return ""
}
// 4317: decl @lune.@base.@convPython.convFilter.outputConvItemType
func (self *convPython_convFilter) outputConvItemType(_env *LnsEnv, typeInfo *Ast_TypeInfo,alt2type LnsAny) {
    var workTypeInfo *Ast_TypeInfo
    workTypeInfo = typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)
    if typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
        if alt2type != nil{
            alt2type_249 := alt2type.(*LnsMap2_[*Ast_TypeInfo,*Ast_TypeInfo])
            {
                _alt := alt2type_249.Get(workTypeInfo)
                if !Lns_IsNil( _alt ) {
                    alt := _alt.(*Ast_TypeInfo)
                    workTypeInfo = alt
                }
            }
        }
    }
    {
        _altType := Ast_AlternateTypeInfoDownCastF(workTypeInfo.FP)
        if !Lns_IsNil( _altType ) {
            altType := _altType.(*Ast_AlternateTypeInfo)
            self.FP.WriteRaw(_env, _env.GetVM().String_format("paramList[%d]", Lns_2DDD(altType.FP.Get_altIndex(_env) - 1)))
        } else {
            self.FP.Writeln(_env, "Lns_ToObjParam{")
            self.FP.PushIndent(_env, nil)
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%sSub, %s,", Lns_2DDD(self.FP.getFromStemName(_env, workTypeInfo), typeInfo.FP.Get_nilable(_env))))
            self.FP.outputConvItemTypeList(_env, workTypeInfo.FP.Get_itemTypeInfoList(_env), alt2type)
            self.FP.PopIndent(_env)
            self.FP.WriteRaw(_env, "}")
        }
    }
}
// 4342: decl @lune.@base.@convPython.convFilter.outputConvItemTypeList
func (self *convPython_convFilter) outputConvItemTypeList(_env *LnsEnv, itemTypeInfoList *LnsList2_[*Ast_TypeInfo],alt2type LnsAny) {
    if itemTypeInfoList.Len() > 0{
        self.FP.WriteRaw(_env, "[]Lns_ToObjParam{")
        self.FP.PushIndent(_env, nil)
        for _index, _itemType := range( itemTypeInfoList.Items ) {
            index := _index + 1
            itemType := _itemType
            if index > 1{
                self.FP.WriteRaw(_env, ",")
            }
            self.FP.outputConvItemType(_env, itemType, alt2type)
        }
        self.FP.PopIndent(_env)
        self.FP.WriteRaw(_env, "}")
    } else { 
        self.FP.WriteRaw(_env, convPython_literalNil)
    }
}
// 4362: decl @lune.@base.@convPython.convFilter.outputAlter2MapFunc
func (self *convPython_convFilter) outputAlter2MapFunc(_env *LnsEnv, alt2Map *LnsMap2_[*Ast_TypeInfo,*Ast_TypeInfo]) {
    __func__ := "@lune.@base.@convPython.convFilter.outputAlter2MapFunc"
    self.FP.WriteRaw(_env, "{")
    for _altType, _assinType := range( alt2Map.Items ) {
        altType := _altType
        assinType := _assinType
        if altType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
            if assinType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                Util_err(_env, _env.GetVM().String_format("not support: %s", Lns_2DDD(__func__)))
            } else { 
                self.FP.outputConvItemType(_env, assinType, alt2Map)
            }
        }
    }
    self.FP.WriteRaw(_env, "}")
}
// 4381: decl @lune.@base.@convPython.convFilter.outputAsyncItem
func (self *convPython_convFilter) outputAsyncItem(_env *LnsEnv, node *Nodes_DeclClassNode) {
}
// 4392: decl @lune.@base.@convPython.convFilter.outputMapping
func (self *convPython_convFilter) outputMapping(_env *LnsEnv, node *Nodes_DeclClassNode,absImmutFlag bool) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType(_env)
    var className string
    className = self.FP.getTypeSymbol(_env, classType)
    self.FP.Writeln(_env, _env.GetVM().String_format("func (self *%s) ToMapSetup( obj *LnsMap ) *LnsMap {", Lns_2DDD(className)))
    self.FP.PushIndent(_env, nil)
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode
        if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
            self.FP.Writeln(_env, _env.GetVM().String_format("obj.Items[\"%s\"] = Lns_ToCollection( self.%s )", Lns_2DDD(memberNode.FP.Get_symbolInfo(_env).FP.Get_name(_env), self.FP.getSymbolSym(_env, memberNode.FP.Get_symbolInfo(_env)))))
        }
    }
    if classType.FP.HasBase(_env){
        self.FP.Writeln(_env, _env.GetVM().String_format("return self.%s.ToMapSetup( obj )", Lns_2DDD(self.FP.getTypeSymbol(_env, classType.FP.Get_baseTypeInfo(_env)))))
    } else { 
        self.FP.Writeln(_env, "return obj")
    }
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
    self.FP.Writeln(_env, _env.GetVM().String_format("func (self *%s) ToMap() *LnsMap {", Lns_2DDD(className)))
    self.FP.Writeln(_env, "    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )")
    self.FP.Writeln(_env, "}")
    var fromStemName string
    fromStemName = self.FP.getFromStemName(_env, classType)
    var classScope *Ast_Scope
    classScope = Lns_unwrap( classType.FP.Get_scope(_env)).(*Ast_Scope)
    if Lns_op_not(classType.FP.Get_abstractFlag(_env)){
        var envStr string
        envStr = convPython_getAddEnvArg_6_(_env, 1, self.option.FP.Get_addEnvArg(_env))
        {
            var fromMapSym *Ast_SymbolInfo
            fromMapSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild(_env, "_fromMap")).(*Ast_SymbolInfo)
            self.FP.Writeln(_env, _env.GetVM().String_format("func %s(%s arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){", Lns_2DDD(self.FP.getSymbolSym(_env, fromMapSym), envStr)))
            self.FP.Writeln(_env, _env.GetVM().String_format("   return %s( arg1, paramList )", Lns_2DDD(fromStemName)))
            self.FP.Writeln(_env, "}")
        }
        {
            var fromStemSym *Ast_SymbolInfo
            fromStemSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild(_env, "_fromStem")).(*Ast_SymbolInfo)
            self.FP.Writeln(_env, _env.GetVM().String_format("func %s(%s arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){", Lns_2DDD(self.FP.getSymbolSym(_env, fromStemSym), envStr)))
            self.FP.Writeln(_env, _env.GetVM().String_format("   return %s( arg1, paramList )", Lns_2DDD(fromStemName)))
            self.FP.Writeln(_env, "}")
        }
        self.FP.Writeln(_env, _env.GetVM().String_format("func %s( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {", Lns_2DDD(fromStemName)))
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, _env.GetVM().String_format("_,conv,mess := %sSub(obj,false, paramList);", Lns_2DDD(fromStemName)))
        self.FP.Writeln(_env, "return conv,mess")
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
        self.FP.Writeln(_env, _env.GetVM().String_format("func %sSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {", Lns_2DDD(fromStemName)))
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, "var objMap *LnsMap")
        self.FP.Writeln(_env, "if work, ok := obj.(*LnsMap); !ok {")
        self.FP.Writeln(_env, _env.GetVM().String_format("   return false, %s, \"no map -- \" + Lns_ToString(obj)", Lns_2DDD(convPython_literalNil)))
        self.FP.Writeln(_env, "} else {")
        self.FP.Writeln(_env, "   objMap = work")
        self.FP.Writeln(_env, "}")
        self.FP.outputNewSetup(_env, "newObj", classType, absImmutFlag)
        self.FP.Writeln(_env, _env.GetVM().String_format("return %sMain( newObj, objMap, paramList )", Lns_2DDD(fromStemName)))
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    }
    self.FP.Writeln(_env, _env.GetVM().String_format("func %sMain( newObj %s, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {", Lns_2DDD(fromStemName, self.FP.type2gotype(_env, classType))))
    self.FP.PushIndent(_env, nil)
    if classType.FP.HasBase(_env){
        self.FP.Writeln(_env, _env.GetVM().String_format("if ok,_,mess := %sMain( &newObj.%s, objMap, paramList ); !ok {", Lns_2DDD(self.FP.getFromStemName(_env, classType.FP.Get_baseTypeInfo(_env)), self.FP.getTypeSymbol(_env, classType.FP.Get_baseTypeInfo(_env)))))
        self.FP.Writeln(_env, _env.GetVM().String_format("    return false,%s,mess", Lns_2DDD(convPython_literalNil)))
        self.FP.Writeln(_env, "}")
    }
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode
        if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
            var memberName string
            memberName = self.FP.getSymbolSym(_env, memberNode.FP.Get_symbolInfo(_env))
            self.FP.WriteRaw(_env, "if ok,conv,mess := ")
            if memberNode.FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                for _index, _itemType := range( classType.FP.Get_itemTypeInfoList(_env).Items ) {
                    index := _index + 1
                    itemType := _itemType
                    if itemType == memberNode.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env){
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("paramList[%d].Func( objMap.Items[\"%s\"], %s, paramList[%d].Child", Lns_2DDD(index - 1, memberNode.FP.Get_symbolInfo(_env).FP.Get_name(_env), memberNode.FP.Get_expType(_env).FP.Get_nilable(_env), index - 1)))
                    }
                }
            } else { 
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%sSub( objMap.Items[\"%s\"], %s, ", Lns_2DDD(self.FP.getFromStemName(_env, memberNode.FP.Get_expType(_env)), memberNode.FP.Get_symbolInfo(_env).FP.Get_name(_env), memberNode.FP.Get_expType(_env).FP.Get_nilable(_env))))
                self.FP.outputConvItemTypeList(_env, memberNode.FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env), nil)
            }
            self.FP.Writeln(_env, "); !ok {")
            self.FP.Writeln(_env, _env.GetVM().String_format("   return false,%s,\"%s:\" + mess.(string)", Lns_2DDD(convPython_literalNil, memberNode.FP.Get_symbolInfo(_env).FP.Get_name(_env))))
            self.FP.Writeln(_env, "} else {")
            self.FP.WriteRaw(_env, _env.GetVM().String_format("   newObj.%s = conv", Lns_2DDD(memberName)))
            self.FP.outputAny2Type(_env, memberNode.FP.Get_expType(_env))
            self.FP.Writeln(_env, "")
            self.FP.Writeln(_env, "}")
        }
    }
    self.FP.Writeln(_env, _env.GetVM().String_format("return true, newObj, %s", Lns_2DDD(convPython_literalNil)))
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}
// 4520: decl @lune.@base.@convPython.convFilter.outputDummyAbstractMethod
func (self *convPython_convFilter) outputDummyAbstractMethod(_env *LnsEnv, classType *Ast_TypeInfo,methodType *Ast_TypeInfo) {
    self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convPython_FuncInfo__WithClass{classType, methodType})
    self.FP.Writeln(_env, "{")
    self.FP.outputDummyReturn(_env, methodType.FP.Get_retTypeInfoList(_env))
    self.FP.Writeln(_env, "}")
}
// 4532: decl @lune.@base.@convPython.convFilter.outputDummyAbstractMethodOfClass
func (self *convPython_convFilter) outputDummyAbstractMethodOfClass(_env *LnsEnv, classType *Ast_TypeInfo) {
    var name2typeMap *LnsMap2_[string,*Ast_TypeInfo]
    name2typeMap = NewLnsMap2_[string,*Ast_TypeInfo]( map[string]*Ast_TypeInfo{})
    var scope *Ast_Scope
    scope = Lns_unwrap( classType.FP.Get_scope(_env)).(*Ast_Scope)
    scope.FP.FilterSymbolInfoIfField(_env, scope, Ast_ScopeAccess__Normal, Ast_filterForm(func(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) bool {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) &&
            _env.SetStackVal( symbolInfo.FP.Get_typeInfo(_env).FP.Get_abstractFlag(_env)) ).(bool)){
            {
                _methodType := scope.FP.GetTypeInfoField(_env, symbolInfo.FP.Get_name(_env), true, scope, Ast_ScopeAccess__Normal)
                if !Lns_IsNil( _methodType ) {
                    methodType := _methodType.(*Ast_TypeInfo)
                    if methodType.FP.Get_abstractFlag(_env){
                        name2typeMap.Set(symbolInfo.FP.Get_name(_env),symbolInfo.FP.Get_typeInfo(_env))
                    }
                } else {
                    name2typeMap.Set(symbolInfo.FP.Get_name(_env),symbolInfo.FP.Get_typeInfo(_env))
                }
            }
        }
        return true
    }))
    {
        __forsortCollection0 := name2typeMap
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
            methodType := __forsortCollection0.Items[ ___forsortKey0 ]
            self.FP.outputDummyAbstractMethod(_env, classType, methodType)
        }
    }
}
// 4565: decl @lune.@base.@convPython.convFilter.outputAdvertise
func (self *convPython_convFilter) outputAdvertise(_env *LnsEnv, node *Nodes_DeclClassNode) {
    __func__ := "@lune.@base.@convPython.convFilter.outputAdvertise"
    var methodNameSet *LnsSet2_[string]
    methodNameSet = node.FP.CreateMethodNameSetWithoutAdv(_env)
    for _, _adv := range( node.FP.Get_advertiseList(_env).Items ) {
        adv := _adv
        if adv.FP.Get_prefix(_env) != ""{
            Util_err(_env, _env.GetVM().String_format("%s: not support advertise with prefix", Lns_2DDD(__func__)))
        }
        {
            _scope := adv.FP.Get_member(_env).FP.Get_expType(_env).FP.Get_scope(_env)
            if !Lns_IsNil( _scope ) {
                scope := _scope.(*Ast_Scope)
                scope.FP.FilterTypeInfoField(_env, true, scope, Ast_ScopeAccess__Normal, Ast_filterForm(func(_env *LnsEnv, symbol *Ast_SymbolInfo) bool {
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( symbol.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) &&
                        _env.SetStackVal( symbol.FP.Get_name(_env) != "__init") &&
                        _env.SetStackVal( Lns_op_not(methodNameSet.Has(symbol.FP.Get_name(_env)))) &&
                        _env.SetStackVal( Lns_op_not(symbol.FP.Get_staticFlag(_env))) ).(bool)){
                        var funcType *Ast_TypeInfo
                        funcType = symbol.FP.Get_typeInfo(_env)
                        self.FP.Writeln(_env, _env.GetVM().String_format("# advertise -- %d", Lns_2DDD(node.FP.Get_pos(_env).LineNo)))
                        self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convPython_FuncInfo__WithClass{node.FP.Get_expType(_env), funcType})
                        self.FP.Writeln(_env, " {")
                        if funcType.FP.Get_retTypeInfoList(_env).Len() > 0{
                            self.FP.WriteRaw(_env, "    return ")
                        }
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("self.%s. ", Lns_2DDD(self.FP.getSymbolSym(_env, adv.FP.Get_member(_env).FP.Get_symbolInfo(_env)))))
                        if adv.FP.Get_member(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Class{
                            self.FP.WriteRaw(_env, "FP.")
                        }
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s( ", Lns_2DDD(self.FP.getSymbolSym(_env, symbol))))
                        self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, funcType.FP.Get_argTypeInfoList(_env).Len(), self.option.FP.Get_addEnvArg(_env)))
                        for _index, _ := range( funcType.FP.Get_argTypeInfoList(_env).Items ) {
                            index := _index + 1
                            if index > 1{
                                self.FP.WriteRaw(_env, ",")
                            }
                            self.FP.WriteRaw(_env, _env.GetVM().String_format("arg%d", Lns_2DDD(index)))
                        }
                        self.FP.Writeln(_env, ")")
                        self.FP.Writeln(_env, "}")
                    }
                    return true
                }))
            }
        }
    }
}
// 4610: decl @lune.@base.@convPython.convFilter.processProtoClass
func (self *convPython_convFilter) ProcessProtoClass(_env *LnsEnv, node *Nodes_ProtoClassNode,_opt LnsAny) {
}
// 4615: decl @lune.@base.@convPython.convFilter.processDeclClass
func (self *convPython_convFilter) ProcessDeclClass(_env *LnsEnv, node *Nodes_DeclClassNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processDeclClass"
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(self.enableTest)) &&
        _env.SetStackVal( node.FP.Get_inTestBlock(_env)) ).(bool)){
        return 
    }
    if node.FP.IsModule(_env){
        return 
    }
    if self.processMode == convPython_ProcessMode__DeclClass{
        if _switch1 := node.FP.Get_expType(_env).FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__Class {
            self.FP.Writeln(_env, _env.GetVM().String_format("# declaration Class -- %s", Lns_2DDD(node.FP.Get_expType(_env).FP.Get_rawTxt(_env))))
            var absImmutFlag bool
            absImmutFlag = self.FP.isInheritAbsImmut(_env, node.FP.Get_expType(_env))
            self.FP.outputStaticMember(_env, node)
            self.FP.outputMethodIF(_env, node)
            self.FP.outputClassType(_env, node, absImmutFlag)
            self.FP.outputToStem(_env, node, absImmutFlag)
            self.FP.outputDownCast(_env, node)
            self.FP.outputCastReceiver(_env, node)
            self.FP.outputConstructor(_env, node, absImmutFlag)
            self.FP.OutputAccessor(_env, node)
            self.FP.outputDummyAbstractMethodOfClass(_env, node.FP.Get_expType(_env))
            self.FP.outputAdvertise(_env, node)
            if node.FP.Get_expType(_env).FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeMapping, nil){
                self.FP.outputMapping(_env, node, absImmutFlag)
            }
            if node.FP.Get_expType(_env).FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeAsyncItem, nil){
                self.FP.outputAsyncItem(_env, node)
            }
            for _, _fieldNode := range( node.FP.Get_fieldList(_env).Items ) {
                fieldNode := _fieldNode
                if _switch0 := fieldNode.FP.Get_kind(_env); _switch0 == Nodes_NodeKind_get_DeclMember(_env) || _switch0 == Nodes_NodeKind_get_DeclMethod(_env) {
                } else {
                    convPython_filter_7_(_env, fieldNode, self, &node.Nodes_Node)
                    self.FP.Writeln(_env, "")
                }
            }
        } else if _switch1 == Ast_TypeInfoKind__IF {
            self.FP.outputInterfaceType(_env, node)
        } else {
            Util_err(_env, _env.GetVM().String_format("%s: not support -- %s:%d", Lns_2DDD(__func__, Ast_TypeInfoKind_getTxt( node.FP.Get_expType(_env).FP.Get_kind(_env)), node.FP.Get_pos(_env).LineNo)))
        }
    } else { 
        self.FP.outputStaticMember(_env, node)
    }
}
// 4713: decl @lune.@base.@convPython.convFilter.outputCallPrefix
func (self *convPython_convFilter) OutputCallPrefix(_env *LnsEnv, callId string,node *Nodes_Node,prefixNode LnsAny,funcSymbol *Ast_SymbolInfo)(bool, LnsAny) {
    var funcType *Ast_TypeInfo
    funcType = funcSymbol.FP.Get_typeInfo(_env)
    var nilAccName string
    nilAccName = _env.GetVM().String_format("%s_%s", Lns_2DDD(Str_replace(_env, self.moduleTypeInfo.FP.Get_rawTxt(_env), "@", ""), callId))
    var callKind LnsAny
    callKind = convPython_CallKind__Normal_Obj
    var extCallFlag bool
    if prefixNode != nil{
        prefixNode_1111 := prefixNode.(*Nodes_Node)
        extCallFlag = prefixNode_1111.FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext
    } else {
        extCallFlag = false
    }
    if extCallFlag{
        callKind = convPython_CallKind__LuaCall_Obj
    }
    var getEnvTxt string
    getEnvTxt = self.env.FP.getEnv(_env)
    var convPython_processNilAcc func(_env *LnsEnv, workPrefixNode *Nodes_Node)
    convPython_processNilAcc = func(_env *LnsEnv, workPrefixNode *Nodes_Node) {
        if Lns_op_not(node.FP.HasNilAccess(_env)){
            return 
        }
        var retNum LnsInt
        retNum = funcType.FP.Get_retTypeInfoList(_env).Len()
        if _switch0 := retNum; _switch0 == 0 {
            self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_NilAccCall0( %s, func () {", Lns_2DDD(getEnvTxt)))
        } else if _switch0 == 1 {
            self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_NilAccCall1( %s, func () LnsAny { return ", Lns_2DDD(getEnvTxt)))
        } else {
            if retNum <= convPython_MaxNilAccNum{
                var anys string
                anys = "LnsAny"
                {
                    var _forFrom0 LnsInt = 2
                    var _forTo0 LnsInt = retNum
                    for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                        anys = _env.GetVM().String_format("%s,LnsAny", Lns_2DDD(anys))
                    }
                }
                self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_NilAccCall%d( %s, func () (%s) { return ", Lns_2DDD(retNum, getEnvTxt, anys)))
            } else { 
                var args string
                args = "LnsAny"
                {
                    var _forFrom1 LnsInt = 2
                    var _forTo1 LnsInt = retNum
                    for _forWork1 := _forFrom1; _forWork1 <= _forTo1; _forWork1++ {
                        args = _env.GetVM().String_format("%s,LnsAny", Lns_2DDD(args))
                    }
                }
                self.FP.WriteRaw(_env, _env.GetVM().String_format("lns_NilAccCall_%s( %s, func () (%s) { return ", Lns_2DDD(nilAccName, getEnvTxt, args)))
            }
        }
        if extCallFlag{
            if funcType.FP.Get_retTypeInfoList(_env).Len() > 1{
                self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_callExt%s( ", Lns_2DDD(node.FP.GetIdTxt(_env))))
            }
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccPop().(%s)", Lns_2DDD(getEnvTxt, self.FP.type2gotype(_env, workPrefixNode.FP.Get_expType(_env).FP.Get_nonnilableType(_env)))))
    }
    var closeParen bool
    closeParen = false
    if prefixNode != nil{
        prefixNode_1136 := prefixNode.(*Nodes_Node)
        if node.FP.HasNilAccess(_env){
            if funcType.FP.Get_retTypeInfoList(_env).Len() >= 2{
                if funcType.FP.Get_retTypeInfoList(_env).Len() <= convPython_MaxNilAccNum{
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_NilAccFinCall%d( ", Lns_2DDD(funcType.FP.Get_retTypeInfoList(_env).Len())))
                } else { 
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("lns_NilAccFinCall_%s(", Lns_2DDD(nilAccName)))
                }
                closeParen = true
            }
        }
        var prefixType *Ast_TypeInfo
        prefixType = prefixNode_1136.FP.Get_expType(_env).FP.Get_nonnilableType(_env)
        if prefixType == Ast_builtinTypeString{
            if node.FP.HasNilAccess(_env){
                Util_err(_env, "not support nilAccName")
            }
            {
                _runtime := self.FP.getVM(_env, funcType)
                if !Lns_IsNil( _runtime ) {
                    runtime := _runtime.(string)
                    callKind = &convPython_CallKind__RuntimeCall{prefixNode_1136}
                    self.FP.WriteRaw(_env, runtime)
                }
            }
        } else { 
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(funcType.FP.Get_staticFlag(_env))) ||
                _env.SetStackVal( funcSymbol.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) ).(bool){
                if node.FP.HasNilAccess(_env){
                    if Lns_op_not(prefixNode_1136.FP.HasNilAccess(_env)){
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccFin(", Lns_2DDD(getEnvTxt)))
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccPush(", Lns_2DDD(getEnvTxt)))
                        convPython_filter_7_(_env, prefixNode_1136, self, node)
                        self.FP.Writeln(_env, ") && ")
                    } else { 
                        convPython_filter_7_(_env, prefixNode_1136, self, node)
                        self.FP.Writeln(_env, "&&")
                    }
                } else { 
                    if extCallFlag{
                        if funcType.FP.Get_retTypeInfoList(_env).Len() > 1{
                            self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_callExt%s( ", Lns_2DDD(node.FP.GetIdTxt(_env))))
                        }
                    }
                    convPython_filter_7_(_env, prefixNode_1136, self, node)
                }
            } else { 
            }
            convPython_processNilAcc(_env, prefixNode_1136)
            if prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                self.FP.WriteRaw(_env, _env.GetVM().String_format(".CallMethod( \"%s\", Lns_2DDD", Lns_2DDD(funcSymbol.FP.Get_name(_env))))
            } else { 
                var prefixKind LnsInt
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                    _env.SetStackVal( prefixType.FP.HasBase(_env)) ).(bool)){
                    prefixKind = prefixType.FP.Get_baseTypeInfo(_env).FP.Get_kind(_env)
                } else { 
                    prefixKind = prefixType.FP.Get_kind(_env)
                }
                if Ast_isBuiltin(_env, funcType.FP.Get_typeId(_env).Id){
                    {
                        _runtime := self.builtin2runtimeEnv.Get(funcType)
                        if !Lns_IsNil( _runtime ) {
                            runtime := _runtime.(string)
                            self.FP.WriteRaw(_env, runtime)
                            callKind = convPython_CallKind__BuiltinCallEnv_Obj
                        } else {
                            if _switch1 := prefixKind; _switch1 == Ast_TypeInfoKind__Class {
                                if self.FP.isInheritAbsImmut(_env, prefixType){
                                    self.FP.WriteRaw(_env, ".FP")
                                }
                                self.FP.WriteRaw(_env, _env.GetVM().String_format(".%s", Lns_2DDD(self.FP.getSymbolSym(_env, funcSymbol))))
                            } else {
                                {
                                    _runtime := self.FP.getVM(_env, funcType)
                                    if !Lns_IsNil( _runtime ) {
                                        runtime := _runtime.(string)
                                        self.FP.WriteRaw(_env, runtime)
                                        callKind = convPython_CallKind__BuiltinCall_Obj
                                    } else {
                                        if _switch0 := funcType; _switch0 == self.builtinFuncs.G_list_sort || _switch0 == self.builtinFuncs.G__list_sort || _switch0 == self.builtinFuncs.Array_sort {
                                            callKind = &convPython_CallKind__SortCall{prefixType.FP.Get_itemTypeInfoList(_env).GetAt(1)}
                                        }
                                        self.FP.WriteRaw(_env, _env.GetVM().String_format(".%s", Lns_2DDD(self.FP.getSymbolSym(_env, funcSymbol))))
                                    }
                                }
                            }
                        }
                    }
                } else { 
                    if _switch3 := funcType.FP.Get_kind(_env); _switch3 == Ast_TypeInfoKind__Method {
                        if _switch2 := prefixKind; _switch2 == Ast_TypeInfoKind__Class {
                            if Lns_op_not(self.FP.isInheritAbsImmut(_env, prefixType)){
                                self.FP.WriteRaw(_env, ".FP")
                            }
                            self.FP.WriteRaw(_env, _env.GetVM().String_format(".%s", Lns_2DDD(self.FP.getSymbolSym(_env, funcSymbol))))
                            if funcSymbol.FP.Get_name(_env) == "_toMap"{
                                callKind = convPython_CallKind__BuiltinCall_Obj
                            }
                        } else {
                            self.FP.WriteRaw(_env, _env.GetVM().String_format(".%s", Lns_2DDD(self.FP.getSymbolSym(_env, funcSymbol))))
                        }
                    } else {
                        if funcSymbol.FP.Get_kind(_env) == Ast_SymbolKind__Mbr{
                            self.FP.WriteRaw(_env, ".")
                        }
                        self.FP.WriteRaw(_env, self.FP.getSymbolSym(_env, funcSymbol))
                    }
                }
            }
        }
    }
    return closeParen, callKind
}
// 4908: decl @lune.@base.@convPython.convFilter.processExpCall
func (self *convPython_convFilter) ProcessExpCall(_env *LnsEnv, node *Nodes_ExpCallNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processExpCall"
    opt := _opt.(*ConvPython_Opt)
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_nonnilableType(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) &&
        _env.SetStackVal( funcType.FP.Get_parentInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) &&
        _env.SetStackVal( funcType.FP.Get_rawTxt(_env) == "get__txt") ).(bool)){
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s(", Lns_2DDD(self.FP.getEnumGetTxtSym(_env, Lns_unwrap( Ast_EnumTypeInfoDownCastF(funcType.FP.Get_parentInfo(_env).FP.Get_aliasSrc(_env).FP)).(*Ast_EnumTypeInfo)))))
        var fieldNode *Nodes_RefFieldNode
        
        {
            _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func(_env).FP)
            if _fieldNode == nil{
                Util_err(_env, "not support -- __func__")
            } else {
                fieldNode = _fieldNode.(*Nodes_RefFieldNode)
            }
        }
        convPython_filter_7_(_env, fieldNode.FP.Get_prefix(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, ")")
        return 
    }
    var retGenerics bool
    if opt.Parent.FP.Get_kind(_env) == Nodes_NodeKind_get_StmtExp(_env){
        retGenerics = false
    } else { 
        retGenerics = convPython_isRetGenerics_22_(_env, node)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( retGenerics) &&
            _env.SetStackVal( funcType.FP.Get_retTypeInfoList(_env).Len() != 1) ).(bool)){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s(", Lns_2DDD(self.FP.getConvGenericsName(_env, &node.Nodes_Node))))
        }
    }
    var addEnvArg bool
    addEnvArg = self.option.FP.Get_addEnvArg(_env)
    var closeParen bool
    closeParen = false
    var callKind LnsAny
    callKind = convPython_CallKind__Normal_Obj
    var withPrefix bool
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func(_env).FP)
        if !Lns_IsNil( _fieldNode ) {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            if _switch0 := fieldNode.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Map || _switch0 == Ast_TypeInfoKind__Set || _switch0 == Ast_TypeInfoKind__Array {
                addEnvArg = false
            }
            if funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.RunLoadedfunc(", Lns_2DDD(self.env.FP.getCommonVm(_env))))
                addEnvArg = false
            }
            withPrefix = true
            closeParen, callKind = self.FP.OutputCallPrefix(_env, node.FP.GetIdTxt(_env), &fieldNode.Nodes_Node, fieldNode.FP.Get_prefix(_env), Lns_unwrap( fieldNode.FP.Get_symbolInfo(_env)).(*Ast_SymbolInfo))
            if funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                self.FP.WriteRaw(_env, ", ")
            } else { 
                self.FP.WriteRaw(_env, "(")
            }
        } else {
            withPrefix = false
            if Ast_isBuiltin(_env, funcType.FP.Get_typeId(_env).Id){
                {
                    _runtime := self.FP.getVM(_env, funcType)
                    if !Lns_IsNil( _runtime ) {
                        runtime := _runtime.(string)
                        self.FP.WriteRaw(_env, runtime)
                        addEnvArg = false
                    } else {
                        if _switch1 := funcType.FP.Get_srcTypeInfo(_env); _switch1 == Ast_builtinTypeForm {
                            convPython_filter_7_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
                            callKind = convPython_CallKind__FormCall_Obj
                        } else if _switch1 == self.builtinFuncs.Lns___run || _switch1 == self.builtinFuncs.Lns___join {
                            convPython_filter_7_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
                        } else {
                            Util_err(_env, _env.GetVM().String_format("%s: not support -- %s:%d", Lns_2DDD(__func__, funcType.FP.GetTxt(_env, nil, nil, nil), node.FP.Get_pos(_env).LineNo)))
                        }
                    }
                }
            } else { 
                if funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.RunLoadedfunc", Lns_2DDD(self.env.FP.getCommonVm(_env))))
                    callKind = convPython_CallKind__RunLoaded_Obj
                } else { 
                    convPython_filter_7_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
                }
            }
            self.FP.WriteRaw(_env, "(")
        }
    }
    var skipArg bool
    skipArg = false
    var closeTxt LnsAny
    closeTxt = nil
    switch _matchExp0 := callKind.(type) {
    case *convPython_CallKind__RuntimeCall:
        prefixNode := _matchExp0.Val1
        convPython_filter_7_(_env, prefixNode, self, node.FP.Get_func(_env))
        if Lns_isCondTrue( node.FP.Get_argList(_env)){
            self.FP.WriteRaw(_env, ",")
        }
        addEnvArg = false
    case *convPython_CallKind__FormCall:
        self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, 1, addEnvArg))
        self.FP.WriteRaw(_env, "Lns_2DDD(")
        addEnvArg = false
    case *convPython_CallKind__RunLoaded:
        convPython_filter_7_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, ",")
        if Lns_op_not(node.FP.Get_argList(_env)){
            self.FP.WriteRaw(_env, "[]LnsAny{}")
        } else { 
            self.FP.WriteRaw(_env, "Lns_2DDD(")
            closeTxt = ")"
        }
        addEnvArg = false
    case *convPython_CallKind__SortCall:
        typeInfo := _matchExp0.Val1
        self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, 2, self.option.FP.Get_addEnvArg(_env)))
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s, ", Lns_2DDD(convPython_getLnsItemKind_28_(_env, typeInfo))))
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                if argList.FP.Get_expType(_env).FP.Get_argTypeInfoList(_env).Len() == 2{
                    skipArg = true
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("LnsComp(func ( %sval1, val2 LnsAny ) bool {", Lns_2DDD(self.FP.getEnvArgDecl(_env, 2))))
                    self.FP.WriteRaw(_env, "return ")
                    self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, argList), funcType.FP.Get_argTypeInfoList(_env), argList, false)
                    self.FP.WriteRaw(_env, _env.GetVM().String_format("( %s", Lns_2DDD(convPython_getAddEnvArg_6_(_env, 2, self.option.FP.Get_addEnvArg(_env)))))
                    self.FP.WriteRaw(_env, "val1")
                    self.FP.outputStem2Type(_env, argList.FP.Get_expType(_env).FP.Get_argTypeInfoList(_env).GetAt(1))
                    self.FP.WriteRaw(_env, ", val2")
                    self.FP.outputStem2Type(_env, argList.FP.Get_expType(_env).FP.Get_argTypeInfoList(_env).GetAt(1))
                    self.FP.WriteRaw(_env, ")})")
                }
            }
        }
    case *convPython_CallKind__BuiltinCall:
        addEnvArg = false
    case *convPython_CallKind__BuiltinCallEnv:
    case *convPython_CallKind__LuaCall:
        closeTxt = ")"
    }
    if Lns_op_not(skipArg){
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, argList), funcType.FP.Get_argTypeInfoList(_env), argList, addEnvArg)
            } else {
                self.FP.WriteRaw(_env, convPython_getAddEnvArg_6_(_env, 0, addEnvArg))
            }
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Func) &&
        _env.SetStackVal( funcType.FP.Get_staticFlag(_env)) &&
        _env.SetStackVal( funcType.FP.Get_parentInfo(_env).FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeMapping, nil)) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( funcType.FP.Get_rawTxt(_env) == "_fromMap") ||
            _env.SetStackVal( funcType.FP.Get_rawTxt(_env) == "_fromStem") ).(bool))) ).(bool)){
        var fieldNode *Nodes_RefFieldNode
        
        {
            _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func(_env).FP)
            if _fieldNode == nil{
                Util_err(_env, "not support -- __func__")
            } else {
                fieldNode = _fieldNode.(*Nodes_RefFieldNode)
            }
        }
        self.FP.WriteRaw(_env, ",")
        self.FP.outputConvItemTypeList(_env, funcType.FP.Get_parentInfo(_env).FP.Get_itemTypeInfoList(_env), fieldNode.FP.Get_prefix(_env).FP.Get_expType(_env).FP.CreateAlt2typeMap(_env, false))
    }
    if closeTxt != nil{
        closeTxt_1237 := closeTxt.(string)
        self.FP.WriteRaw(_env, closeTxt_1237)
    }
    self.FP.WriteRaw(_env, ")")
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( callKind == convPython_CallKind__LuaCall_Obj) ||
        _env.SetStackVal( callKind == convPython_CallKind__RunLoaded_Obj) ).(bool){
        if funcType.FP.Get_retTypeInfoList(_env).Len() == 1{
            if opt.Parent.FP.Get_kind(_env) != Nodes_NodeKind_get_StmtExp(_env){
                self.FP.WriteRaw(_env, "[0]")
                var retTypeList *LnsList2_[*Ast_TypeInfo]
                retTypeList = Lns_unwrap( Lns_car(Ast_convToExtTypeList(_env, self.processInfo, funcType.FP.Get_retTypeInfoList(_env)))).(*LnsList2_[*Ast_TypeInfo])
                self.FP.outputAny2Type(_env, retTypeList.GetAt(1))
            }
        } else if funcType.FP.Get_retTypeInfoList(_env).Len() > 1{
            self.FP.WriteRaw(_env, ")")
        }
    }
    if retGenerics{
        if funcType.FP.Get_retTypeInfoList(_env).Len() == 1{
            var retType *Ast_TypeInfo
            retType = funcType.FP.Get_retTypeInfoList(_env).GetAt(1)
            if convPython_isAnyType_8_(_env, retType){
                if Lns_op_not(convPython_isAnyType_8_(_env, node.FP.Get_expType(_env))){
                    self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
                }
            } else if retType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                if retType.FP.Get_srcTypeInfo(_env) != node.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env){
                    self.FP.WriteRaw(_env, ".FP")
                    self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
                }
            }
        } else { 
            self.FP.WriteRaw(_env, ")")
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( withPrefix) &&
        _env.SetStackVal( node.FP.HasNilAccess(_env)) ).(bool)){
        self.FP.WriteRaw(_env, "})")
        self.FP.WriteRaw(_env, _env.GetVM().String_format("/* %d:%d */", Lns_2DDD(node.FP.Get_pos(_env).LineNo, node.FP.Get_pos(_env).Column)))
        if opt.Parent.FP.HasNilAccess(_env){
        } else { 
            self.FP.WriteRaw(_env, ")")
        }
        if closeParen{
            self.FP.WriteRaw(_env, ")")
        }
    }
    if callKind == convPython_CallKind__FormCall_Obj{
        self.FP.WriteRaw(_env, ")")
    }
}
// 5159: decl @lune.@base.@convPython.convFilter.processExpMRet
func (self *convPython_convFilter) ProcessExpMRet(_env *LnsEnv, node *Nodes_ExpMRetNode,_opt LnsAny) {
    convPython_filter_7_(_env, node.FP.Get_mRet(_env), self, &node.Nodes_Node)
}
// 5166: decl @lune.@base.@convPython.convFilter.processExpAccessMRet
func (self *convPython_convFilter) ProcessExpAccessMRet(_env *LnsEnv, node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
}
// 5174: decl @lune.@base.@convPython.convFilter.processExpList
func (self *convPython_convFilter) ProcessExpList(_env *LnsEnv, node *Nodes_ExpListNode,_opt LnsAny) {
    for _index, _exp := range( node.FP.Get_expList(_env).Items ) {
        index := _index + 1
        exp := _exp
        if index != 1{
            self.FP.WriteRaw(_env, ", ")
        }
        {
            _mRetExp := node.FP.Get_mRetExp(_env)
            if !Lns_IsNil( _mRetExp ) {
                mRetExp := _mRetExp.(*Nodes_MRetExp)
                if mRetExp.FP.Get_index(_env) == index{
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( index == 1) ||
                        _env.SetStackVal( exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool){
                        convPython_filter_7_(_env, exp, self, &node.Nodes_Node)
                    } else { 
                        self.FP.WriteRaw(_env, "Lns_2DDD(")
                        convPython_filter_7_(_env, exp, self, &node.Nodes_Node)
                        self.FP.WriteRaw(_env, ")")
                    }
                    break
                }
            }
        }
        convPython_filter_7_(_env, exp, self, &node.Nodes_Node)
    }
}
// 5198: decl @lune.@base.@convPython.convFilter.processExpOp1
func (self *convPython_convFilter) ProcessExpOp1(_env *LnsEnv, node *Nodes_ExpOp1Node,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processExpOp1"
    if _switch2 := node.FP.Get_op(_env).Txt; _switch2 == "~" {
        self.FP.WriteRaw(_env, "^")
        convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    } else if _switch2 == "+" || _switch2 == "-" {
        self.FP.WriteRaw(_env, node.FP.Get_op(_env).Txt)
        convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    } else if _switch2 == "not" {
        self.FP.WriteRaw(_env, "Lns_op_not(")
        convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, ")")
    } else if _switch2 == "#" {
        if _switch1 := node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env); _switch1 == Ast_TypeInfoKind__List {
            self.FP.WriteRaw(_env, "len(")
            convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        } else if _switch1 == Ast_TypeInfoKind__Ext {
            if _switch0 := node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_extedType(_env).FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List {
                convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
                self.FP.WriteRaw(_env, ".Len()")
            } else {
                Util_err(_env, _env.GetVM().String_format("%s: not support -- %s", Lns_2DDD(__func__, node.FP.Get_exp(_env).FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))))
            }
        } else {
            self.FP.WriteRaw(_env, "len(")
            convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        }
    } else {
        Util_err(_env, _env.GetVM().String_format("%s: not support -- %s", Lns_2DDD(__func__, node.FP.Get_op(_env).Txt)))
    }
}
// 5247: decl @lune.@base.@convPython.convFilter.processExpMultiTo1
func (self *convPython_convFilter) ProcessExpMultiTo1(_env *LnsEnv, node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    self.FP.WriteRaw(_env, "Lns_car(")
    convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, ")")
    self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
}
// 5259: decl @lune.@base.@convPython.convFilter.processExpCast
func (self *convPython_convFilter) ProcessExpCast(_env *LnsEnv, node *Nodes_ExpCastNode,_opt LnsAny) {
    if _switch1 := node.FP.Get_castKind(_env); _switch1 == Nodes_CastKind__Force {
        if convPython_isAnyType_8_(_env, node.FP.Get_exp(_env).FP.Get_expType(_env)){
            if _switch0 := node.FP.Get_castType(_env); _switch0 == Ast_builtinTypeInt {
                self.FP.WriteRaw(_env, "Lns_forceCastInt(")
                convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
                self.FP.WriteRaw(_env, ")")
            } else if _switch0 == Ast_builtinTypeReal {
                self.FP.WriteRaw(_env, "Lns_forceCastReal(")
                convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
                self.FP.WriteRaw(_env, ")")
            } else {
                convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
                self.FP.outputAny2Type(_env, node.FP.Get_castType(_env))
            }
        } else { 
            self.FP.WriteRaw(_env, _env.GetVM().String_format("(%s)(", Lns_2DDD(self.FP.type2gotype(_env, node.FP.Get_castType(_env)))))
            convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        }
    } else if _switch1 == Nodes_CastKind__Implicit {
        if node.FP.Get_exp(_env).FP.Get_expTypeList(_env).Len() > 1{
            convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else { 
            self.FP.outputImplicitCast(_env, node.FP.Get_castType(_env), node.FP.Get_exp(_env), node)
        }
    } else if _switch1 == Nodes_CastKind__Normal {
        var typeName string
        typeName = self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_castType(_env))
        var castType *Ast_TypeInfo
        castType = node.FP.Get_castType(_env).FP.Get_nonnilableType(_env)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( castType.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
            _env.SetStackVal( castType != Ast_builtinTypeString) ).(bool)){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%sDownCastF(", Lns_2DDD(typeName)))
            convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
                _env.SetStackVal( node.FP.Get_exp(_env).FP.Get_expType(_env) != Ast_builtinTypeString) ).(bool)){
                self.FP.WriteRaw(_env, ".FP")
            }
            self.FP.WriteRaw(_env, ")")
        } else { 
            self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_cast2%s( ", Lns_2DDD(self.FP.type2gotype(_env, castType))))
            convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        }
    }
}
// 5319: decl @lune.@base.@convPython.convFilter.processExpParen
func (self *convPython_convFilter) ProcessExpParen(_env *LnsEnv, node *Nodes_ExpParenNode,_opt LnsAny) {
    if node.FP.Get_exp(_env).FP.Get_expTypeList(_env).Len() >= 2{
        self.FP.WriteRaw(_env, "Lns_car(")
        convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, ")")
        self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
    } else { 
        self.FP.WriteRaw(_env, "(")
        convPython_filter_7_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, ")")
    }
}
// 5336: decl @lune.@base.@convPython.convFilter.processExpSetVal
func (self *convPython_convFilter) ProcessExpSetVal(_env *LnsEnv, node *Nodes_ExpSetValNode,_opt LnsAny) {
    convPython_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
    if convPython_getExpListKind_21_(_env, node.FP.Get_exp1(_env).FP.Get_expTypeList(_env), node.FP.Get_exp2(_env), self.option.FP.Get_addEnvArg(_env)) == convPython_ExpListKind__Direct{
        {
            var _forFrom0 LnsInt = node.FP.Get_exp1(_env).FP.Get_expTypeList(_env).Len() + 1
            var _forTo0 LnsInt = node.FP.Get_exp2(_env).FP.Get_expTypeList(_env).Len()
            for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                self.FP.WriteRaw(_env, ",_")
            }
        }
    }
    self.FP.WriteRaw(_env, " = ")
    self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, &node.Nodes_Node, node.FP.Get_exp2(_env)), node.FP.Get_exp1(_env).FP.Get_expTypeList(_env), node.FP.Get_exp2(_env), false)
}
// 5353: decl @lune.@base.@convPython.convFilter.processExpSetItem
func (self *convPython_convFilter) ProcessExpSetItem(_env *LnsEnv, node *Nodes_ExpSetItemNode,_opt LnsAny) {
    convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
    self.FP.WriteRaw(_env, "[")
    switch _matchExp0 := node.FP.Get_index(_env).(type) {
    case *Nodes_IndexVal__NodeIdx:
        index := _matchExp0.Val1
        convPython_filter_7_(_env, index, self, &node.Nodes_Node)
    case *Nodes_IndexVal__SymIdx:
        index := _matchExp0.Val1
        self.FP.WriteRaw(_env, _env.GetVM().String_format("\"%s\"", Lns_2DDD(index)))
    }
    self.FP.WriteRaw(_env, "] = ")
    convPython_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
}
// 5370: decl @lune.@base.@convPython.convFilter.processAndOr
func (self *convPython_convFilter) processAndOr(_env *LnsEnv, node *Nodes_ExpOp2Node,opTxt string,parent *Nodes_Node) {
    var getEnvTxt string
    getEnvTxt = self.env.FP.getEnv(_env)
    var firstFlag bool
    firstFlag = Lns_op_not(convPython_convFilter_processAndOr__isAndOr_0_(_env, parent))
    if firstFlag{
        self.FP.Writeln(_env, _env.GetVM().String_format("%s.PopVal( %s.IncStack() ||", Lns_2DDD(getEnvTxt, getEnvTxt)))
        self.FP.PushIndent(_env, nil)
    }
    var opCC string
    if opTxt == "and"{
        opCC = "&&"
    } else { 
        opCC = "||"
    }
    if convPython_convFilter_processAndOr__isAndOr_0_(_env, node.FP.Get_exp1(_env)){
        convPython_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
    } else { 
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.SetStackVal( ", Lns_2DDD(getEnvTxt)))
        convPython_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, ") ")
    }
    self.FP.Writeln(_env, opCC)
    if convPython_convFilter_processAndOr__isAndOr_0_(_env, node.FP.Get_exp2(_env)){
        convPython_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
    } else { 
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.SetStackVal( ", Lns_2DDD(getEnvTxt)))
        convPython_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, ") ")
    }
    if firstFlag{
        self.FP.WriteRaw(_env, ")")
        self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
        self.FP.PopIndent(_env)
    }
}
// 5425: decl @lune.@base.@convPython.convFilter.checkExpOp3
func (self *convPython_convFilter) CheckExpOp3(_env *LnsEnv, node *Nodes_ExpOp2Node,opt *ConvPython_Opt) {
    {
        _andExp := Nodes_ExpOp2NodeDownCastF(node.FP.Get_exp1(_env).FP)
        if !Lns_IsNil( _andExp ) {
            andExp := _andExp.(*Nodes_ExpOp2Node)
            if andExp.FP.Get_op(_env).Txt == "and"{
                convPython_filter_7_(_env, andExp.FP.Get_exp2(_env), self, &node.Nodes_Node)
                self.FP.WriteRaw(_env, " if ")
                self.FP.processCond(_env, andExp.FP.Get_exp1(_env), &node.Nodes_Node)
                self.FP.WriteRaw(_env, " else ")
                convPython_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
                return 
            }
        }
    }
    self.FP.processAndOr(_env, node, node.FP.Get_op(_env).Txt, opt.Parent)
}
// 5444: decl @lune.@base.@convPython.convFilter.processExpOp2
func (self *convPython_convFilter) ProcessExpOp2(_env *LnsEnv, node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt.(*ConvPython_Opt)
    var opTxt string
    opTxt = node.FP.Get_op(_env).Txt
    if _switch1 := opTxt; _switch1 == "or" {
        self.FP.CheckExpOp3(_env, node, opt)
    } else if _switch1 == "and" {
        self.FP.processAndOr(_env, node, opTxt, opt.Parent)
    } else if _switch1 == ".." {
        convPython_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, " + ")
        convPython_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
    } else if _switch1 == "/" {
        convPython_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
        if node.FP.Get_expType(_env).FP.Equals(_env, self.processInfo, Ast_builtinTypeInt, nil, nil){
            self.FP.WriteRaw(_env, " // ")
        } else { 
            self.FP.WriteRaw(_env, " / ")
        }
        convPython_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
    } else {
        {
            __exp := Ast_bitBinOpMap.Get(opTxt)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(LnsInt)
                if _switch0 := _exp; _switch0 == Ast_BitOpKind__LShift {
                    opTxt = "<<"
                } else if _switch0 == Ast_BitOpKind__RShift {
                    opTxt = ">>"
                }
                convPython_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
                self.FP.WriteRaw(_env, " " + opTxt + " ")
                convPython_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
            } else {
                convPython_filter_7_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
                {
                    _op := convPython_op2map.Get(opTxt)
                    if !Lns_IsNil( _op ) {
                        op := _op.(string)
                        self.FP.WriteRaw(_env, _env.GetVM().String_format(" %s ", Lns_2DDD(op)))
                    } else {
                        self.FP.WriteRaw(_env, _env.GetVM().String_format(" %s ", Lns_2DDD(opTxt)))
                    }
                }
                convPython_filter_7_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
            }
        }
    }
}
// 5501: decl @lune.@base.@convPython.convFilter.processExpRef
func (self *convPython_convFilter) ProcessExpRef(_env *LnsEnv, node *Nodes_ExpRefNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processExpRef"
    if node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        self.FP.WriteRaw(_env, "ddd")
    } else { 
        if Lns_isCondTrue( node.FP.Get_symbolInfo(_env).FP.Get_convModuleParam(_env)){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s_%d", Lns_2DDD(node.FP.Get_symbolInfo(_env).FP.Get_name(_env), node.FP.Get_symbolInfo(_env).FP.Get_symbolId(_env))))
        } else { 
            if _switch1 := node.FP.Get_symbolInfo(_env).FP.Get_kind(_env); _switch1 == Ast_SymbolKind__Var || _switch1 == Ast_SymbolKind__Arg {
                self.FP.WriteRaw(_env, self.FP.getSymbolSym(_env, node.FP.Get_symbolInfo(_env)))
            } else if _switch1 == Ast_SymbolKind__Fun {
                if Ast_isBuiltin(_env, node.FP.Get_expType(_env).FP.Get_typeId(_env).Id){
                    var builtinFunc *Builtin_BuiltinFuncType
                    builtinFunc = self.builtinFuncs
                    if _switch0 := node.FP.Get_expType(_env); _switch0 == builtinFunc.Lns_print {
                        self.FP.WriteRaw(_env, "print")
                    } else if _switch0 == builtinFunc.Lns___run {
                        self.FP.WriteRaw(_env, "LnsRun")
                    } else if _switch0 == builtinFunc.Lns___join {
                        self.FP.WriteRaw(_env, "LnsJoin")
                    } else {
                        Util_err(_env, _env.GetVM().String_format("%s: not support -- %s", Lns_2DDD(__func__, node.FP.Get_symbolInfo(_env).FP.Get_name(_env))))
                    }
                } else { 
                    self.FP.WriteRaw(_env, self.FP.getSymbol(_env, &convPython_SymbolKind__Func{node.FP.Get_expType(_env)}, node.FP.Get_symbolInfo(_env).FP.Get_name(_env)))
                }
            } else if _switch1 == Ast_SymbolKind__Typ {
                if node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Module{
                    self.FP.WriteRaw(_env, node.FP.Get_symbolInfo(_env).FP.Get_name(_env))
                } else { 
                    self.FP.WriteRaw(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
                }
            } else {
                self.FP.WriteRaw(_env, node.FP.Get_symbolInfo(_env).FP.Get_name(_env))
            }
        }
    }
}
// 5554: decl @lune.@base.@convPython.convFilter.processExpRefItem
func (self *convPython_convFilter) ProcessExpRefItem(_env *LnsEnv, node *Nodes_ExpRefItemNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processExpRefItem"
    var getEnvTxt string
    getEnvTxt = self.env.FP.getEnv(_env)
    var prefixType *Ast_TypeInfo
    prefixType = node.FP.Get_val(_env).FP.Get_expType(_env).FP.Get_nonnilableType(_env)
    if _switch0 := prefixType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Ext {
        if node.FP.Get_nilAccess(_env){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccFin( %s.NilAccPush( ", Lns_2DDD(getEnvTxt, getEnvTxt)))
            convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ") && ")
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccPush( %s.NilAccPop().(*Lns_luaValue)", Lns_2DDD(getEnvTxt, getEnvTxt)))
        } else { 
            convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            if prefixType.FP.Get_extedType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Stem{
                self.FP.WriteRaw(_env, ".(*Lns_luaValue)")
            }
        }
        self.FP.WriteRaw(_env, ".GetAt(")
        {
            _index := node.FP.Get_index(_env)
            if !Lns_IsNil( _index ) {
                index := _index.(*Nodes_Node)
                convPython_filter_7_(_env, index, self, &node.Nodes_Node)
            } else {
                self.FP.WriteRaw(_env, _env.GetVM().String_format("\"%s\"", Lns_2DDD(convPython_str2gostr_14_(_env, Lns_unwrap( node.FP.Get_symbol(_env)).(string)))))
            }
        }
        self.FP.WriteRaw(_env, ")")
        self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
    } else if _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        if node.FP.Get_nilAccess(_env){
            if Lns_op_not(node.FP.Get_val(_env).FP.HasNilAccess(_env)){
                self.FP.Writeln(_env, _env.GetVM().String_format("%s.NilAccFin( %s.NilAccPush(", Lns_2DDD(getEnvTxt, getEnvTxt)))
                convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
                self.FP.Writeln(_env, ") && ")
            } else { 
                convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
                self.FP.Writeln(_env, "&&")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccPush( %s.NilAccPop().(*LnsList)", Lns_2DDD(getEnvTxt, getEnvTxt)))
            self.FP.WriteRaw(_env, ".GetAt(")
            convPython_filter_7_(_env, Lns_unwrap( node.FP.Get_index(_env)).(*Nodes_Node), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
            self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
            self.FP.WriteRaw(_env, "))")
        } else { 
            convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, "[")
            convPython_filter_7_(_env, Lns_unwrap( node.FP.Get_index(_env)).(*Nodes_Node), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, "-1]")
            self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
        }
    } else if _switch0 == Ast_TypeInfoKind__Map {
        if node.FP.Get_nilAccess(_env){
            if Lns_op_not(node.FP.Get_val(_env).FP.HasNilAccess(_env)){
                self.FP.Writeln(_env, _env.GetVM().String_format("%s.NilAccFin( %s.NilAccPush(", Lns_2DDD(getEnvTxt, getEnvTxt)))
                convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
                self.FP.Writeln(_env, ") && ")
            } else { 
                convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
                self.FP.Writeln(_env, "&&")
            }
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccPush( %s.NilAccPop().(*LnsMap)", Lns_2DDD(getEnvTxt, getEnvTxt)))
        } else { 
            convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            if prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Stem{
                self.FP.WriteRaw(_env, ".(*LnsMap)")
            }
        }
        self.FP.WriteRaw(_env, ".get(")
        {
            _index := node.FP.Get_index(_env)
            if !Lns_IsNil( _index ) {
                index := _index.(*Nodes_Node)
                convPython_filter_7_(_env, index, self, &node.Nodes_Node)
            } else {
                self.FP.WriteRaw(_env, _env.GetVM().String_format("\"%s\"", Lns_2DDD(convPython_str2gostr_14_(_env, Lns_unwrap( node.FP.Get_symbol(_env)).(string)))))
            }
        }
        self.FP.WriteRaw(_env, ")")
        self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
        if node.FP.Get_nilAccess(_env){
            self.FP.WriteRaw(_env, "))")
        }
    } else if _switch0 == Ast_TypeInfoKind__Stem {
        self.FP.WriteRaw(_env, "Lns_FromStemGetAt(")
        convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
        self.FP.WriteRaw(_env, ",")
        {
            _index := node.FP.Get_index(_env)
            if !Lns_IsNil( _index ) {
                index := _index.(*Nodes_Node)
                convPython_filter_7_(_env, index, self, &node.Nodes_Node)
            } else {
                self.FP.WriteRaw(_env, _env.GetVM().String_format("\"%s\"", Lns_2DDD(convPython_str2gostr_14_(_env, Lns_unwrap( node.FP.Get_symbol(_env)).(string)))))
            }
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format(", %s )", Lns_2DDD(node.FP.Get_nilAccess(_env))))
        self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
    } else {
        if prefixType == Ast_builtinTypeString{
            self.FP.WriteRaw(_env, "ord(")
            convPython_filter_7_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, "[")
            convPython_filter_7_(_env, Lns_unwrap( node.FP.Get_index(_env)).(*Nodes_Node), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, "-1])")
        } else { 
            Util_err(_env, _env.GetVM().String_format("%s: not support -- %d, %s", Lns_2DDD(__func__, node.FP.Get_pos(_env).LineNo, Ast_TypeInfoKind_getTxt( prefixType.FP.Get_kind(_env)))))
        }
    }
}
// 5665: decl @lune.@base.@convPython.convFilter.processRefField
func (self *convPython_convFilter) ProcessRefField(_env *LnsEnv, node *Nodes_RefFieldNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processRefField"
    opt := _opt.(*ConvPython_Opt)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) &&
        _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) ).(bool)){
        self.FP.WriteRaw(_env, self.FP.getSymbol(_env, &convPython_SymbolKind__Static{node.FP.Get_expType(_env)}, node.FP.Get_field(_env).Txt))
        return 
    }
    {
        _symbol := node.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( _symbol ) {
            symbol := _symbol.(*Ast_SymbolInfo)
            var orgSym *Ast_SymbolInfo
            orgSym = symbol.FP.GetOrg(_env)
            {
                _code := self.builtin2code.Get(orgSym)
                if !Lns_IsNil( _code ) {
                    code := _code.(string)
                    self.FP.WriteRaw(_env, code)
                    return 
                }
            }
            if self.builtinFuncs.FP.Get_allSymbolSet(_env).Has(orgSym){
                self.FP.WriteRaw(_env, _env.GetVM().String_format("Lns_%s_%s", Lns_2DDD(Str_replace(_env, node.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_rawTxt(_env), "@", ""), orgSym.FP.Get_name(_env))))
                return 
            }
            if symbol.FP.Get_staticFlag(_env){
                self.FP.WriteRaw(_env, self.FP.getSymbolSym(_env, symbol))
                return 
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(symbol.FP.Get_scope(_env).FP.Get_ownerTypeInfo(_env)) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) == Ast_TypeInfoKind__Module) &&
                _env.SetStackVal( symbol.FP.Get_kind(_env) == Ast_SymbolKind__Var) ).(bool)){
                self.FP.WriteRaw(_env, self.FP.getSymbolSym(_env, symbol))
                return 
            }
        }
    }
    var getEnvTxt string
    getEnvTxt = self.env.FP.getEnv(_env)
    var openParenNum LnsInt
    if Lns_op_not(node.FP.HasNilAccess(_env)){
        openParenNum = 0
        if Lns_op_not(node.FP.Get_prefix(_env).FP.HasNilAccess(_env)){
            convPython_filter_7_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
        } else { 
            Util_err(_env, _env.GetVM().String_format("%s: not support", Lns_2DDD(__func__)))
        }
    } else { 
        if Lns_op_not(node.FP.Get_prefix(_env).FP.HasNilAccess(_env)){
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccFin(", Lns_2DDD(getEnvTxt)))
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccPush(", Lns_2DDD(getEnvTxt)))
            convPython_filter_7_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
            self.FP.Writeln(_env, ") && ")
        } else { 
            convPython_filter_7_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
            self.FP.Writeln(_env, "&&")
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccPush(", Lns_2DDD(getEnvTxt)))
        if opt.Parent.FP.HasNilAccess(_env){
            openParenNum = 1
        } else { 
            openParenNum = 2
        }
        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s.NilAccPop().(%s)", Lns_2DDD(getEnvTxt, self.FP.type2gotype(_env, node.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_nonnilableType(_env)))))
    }
    self.FP.WriteRaw(_env, ".")
    {
        _symbol := node.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( _symbol ) {
            symbol := _symbol.(*Ast_SymbolInfo)
            if node.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                self.FP.WriteRaw(_env, _env.GetVM().String_format("GetAt( \"%s\" )", Lns_2DDD(symbol.FP.Get_name(_env))))
                self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
            } else { 
                self.FP.WriteRaw(_env, self.FP.getSymbolSym(_env, symbol))
                var orgSym *Ast_SymbolInfo
                orgSym = symbol.FP.GetOrg(_env)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( orgSym.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) &&
                    _env.SetStackVal( orgSym.FP.Get_typeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                    _env.SetStackVal( convPython_isAnyType_8_(_env, orgSym.FP.Get_typeInfo(_env))) ).(bool)){
                    self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
                }
            }
        } else {
            Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
        }
    }
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = openParenNum
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            self.FP.WriteRaw(_env, ")")
        }
    }
}
// 5753: decl @lune.@base.@convPython.convFilter.processExpOmitEnum
func (self *convPython_convFilter) ProcessExpOmitEnum(_env *LnsEnv, node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, self.FP.getSymbol(_env, &convPython_SymbolKind__Static{node.FP.Get_expType(_env).FP.Get_aliasSrc(_env)}, node.FP.Get_valInfo(_env).FP.Get_name(_env)))
}
// 5759: decl @lune.@base.@convPython.convFilter.processGetField
func (self *convPython_convFilter) ProcessGetField(_env *LnsEnv, node *Nodes_GetFieldNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processGetField"
    opt := _opt.(*ConvPython_Opt)
    {
        _symbolInfo := node.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( _symbolInfo ) {
            symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) &&
                _env.SetStackVal( symbolInfo.FP.Get_name(_env) == "get__txt") ).(bool)){
                {
                    _enumType := Ast_EnumTypeInfoDownCastF(symbolInfo.FP.Get_namespaceTypeInfo(_env).FP)
                    if !Lns_IsNil( _enumType ) {
                        enumType := _enumType.(*Ast_EnumTypeInfo)
                        self.FP.WriteRaw(_env, _env.GetVM().String_format("%s( ", Lns_2DDD(self.FP.getEnumGetTxtSym(_env, enumType))))
                        convPython_filter_7_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
                        self.FP.WriteRaw(_env, ")")
                        return 
                    }
                }
                if Lns_isCondTrue( Ast_AlgeTypeInfoDownCastF(symbolInfo.FP.Get_namespaceTypeInfo(_env).FP)){
                    convPython_filter_7_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
                    self.FP.WriteRaw(_env, ".(LnsAlgeVal).GetTxt()")
                    return 
                }
            }
            if symbolInfo.FP.Get_staticFlag(_env){
                self.FP.WriteRaw(_env, self.FP.getSymbolSym(_env, symbolInfo))
                self.FP.WriteRaw(_env, _env.GetVM().String_format("(%s)", Lns_2DDD(convPython_getAddEnvArg_6_(_env, 0, self.option.FP.Get_addEnvArg(_env)))))
            } else { 
                var closeParen bool
                closeParen = convPython_convExp0_9177(Lns_2DDD(self.FP.OutputCallPrefix(_env, node.FP.GetIdTxt(_env), &node.Nodes_Node, node.FP.Get_prefix(_env), symbolInfo)))
                self.FP.WriteRaw(_env, _env.GetVM().String_format("(%s)", Lns_2DDD(convPython_getAddEnvArg_6_(_env, 0, self.option.FP.Get_addEnvArg(_env)))))
                var retType *Ast_TypeInfo
                retType = symbolInfo.FP.Get_typeInfo(_env).FP.Get_retTypeInfoList(_env).GetAt(1)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( retType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                    _env.SetStackVal( Lns_op_not(retType.FP.HasBase(_env))) ).(bool)){
                    self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
                }
                if node.FP.HasNilAccess(_env){
                    self.FP.WriteRaw(_env, "})")
                    if opt.Parent.FP.HasNilAccess(_env){
                    } else { 
                        self.FP.WriteRaw(_env, ")")
                    }
                }
                if closeParen{
                    self.FP.WriteRaw(_env, ")")
                }
            }
        } else {
            Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
        }
    }
}
// 5808: decl @lune.@base.@convPython.convFilter.processReturn
func (self *convPython_convFilter) ProcessReturn(_env *LnsEnv, node *Nodes_ReturnNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "return ")
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            convPython_filter_7_(_env, &expList.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln(_env, "")
}
// 5818: decl @lune.@base.@convPython.convFilter.processTestCase
func (self *convPython_convFilter) ProcessTestCase(_env *LnsEnv, node *Nodes_TestCaseNode,_opt LnsAny) {
    if Lns_op_not(self.enableTest){
        return 
    }
    self.FP.Writeln(_env, "import unittest\nfrom test import support\n\nclass MyTestCase(unittest.TestCase):\n    def setUp(self):\n        pass\n    def tearDown(self):\n        pass\n    def test_case(self):\n")
    self.FP.PushIndent(_env, nil)
    convPython_filter_7_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.PopIndent(_env)
}
// 5843: decl @lune.@base.@convPython.convFilter.processTestBlock
func (self *convPython_convFilter) ProcessTestBlock(_env *LnsEnv, node *Nodes_TestBlockNode,_opt LnsAny) {
    var stmtList *LnsList2_[*Nodes_Node]
    stmtList = node.FP.Get_stmtList(_env)
    for _, _statement := range( stmtList.Items ) {
        statement := _statement
        if Lns_op_not(self.ignoreNodeInInnerBlockSet.Has(LnsInt(statement.FP.Get_kind(_env)))){
            convPython_filter_7_(_env, statement, self, &node.Nodes_Node)
        }
    }
}
// 5855: decl @lune.@base.@convPython.convFilter.processProvide
func (self *convPython_convFilter) ProcessProvide(_env *LnsEnv, node *Nodes_ProvideNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processProvide"
    Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
}
// 5861: decl @lune.@base.@convPython.convFilter.processAlias
func (self *convPython_convFilter) ProcessAlias(_env *LnsEnv, node *Nodes_AliasNode,_opt LnsAny) {
}
// 5866: decl @lune.@base.@convPython.convFilter.processBoxing
func (self *convPython_convFilter) ProcessBoxing(_env *LnsEnv, node *Nodes_BoxingNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processBoxing"
    Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
}
// 5872: decl @lune.@base.@convPython.convFilter.processUnboxing
func (self *convPython_convFilter) ProcessUnboxing(_env *LnsEnv, node *Nodes_UnboxingNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processUnboxing"
    Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(__func__)))
}
// 5878: decl @lune.@base.@convPython.convFilter.processLiteralList
func (self *convPython_convFilter) ProcessLiteralList(_env *LnsEnv, node *Nodes_LiteralListNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "[")
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            convPython_filter_7_(_env, &expList.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.WriteRaw(_env, "]")
}
// 5889: decl @lune.@base.@convPython.convFilter.processLiteralSet
func (self *convPython_convFilter) ProcessLiteralSet(_env *LnsEnv, node *Nodes_LiteralSetNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "set(")
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.expList2Slice(_env, expList, true)
        }
    }
    self.FP.WriteRaw(_env, ")")
}
// 5900: decl @lune.@base.@convPython.convFilter.processLiteralMap
func (self *convPython_convFilter) ProcessLiteralMap(_env *LnsEnv, node *Nodes_LiteralMapNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "{")
    for _, _pair := range( node.FP.Get_pairList(_env).Items ) {
        pair := _pair
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( pair.FP.Get_key(_env).FP.Get_kind(_env) == Nodes_NodeKind_get_LiteralNil(_env)) ||
            _env.SetStackVal( pair.FP.Get_val(_env).FP.Get_kind(_env) == Nodes_NodeKind_get_LiteralNil(_env)) ).(bool){
        } else { 
            convPython_filter_7_(_env, pair.FP.Get_key(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ":")
            convPython_filter_7_(_env, pair.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ",")
        }
    }
    self.FP.WriteRaw(_env, "}")
}
// 5919: decl @lune.@base.@convPython.convFilter.processLiteralArray
func (self *convPython_convFilter) ProcessLiteralArray(_env *LnsEnv, node *Nodes_LiteralArrayNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "NewLnsList(")
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.expList2Slice(_env, expList, true)
        } else {
            self.FP.WriteRaw(_env, "[]LnsAny{}")
        }
    }
    self.FP.WriteRaw(_env, ")")
}
// 5933: decl @lune.@base.@convPython.convFilter.processLiteralChar
func (self *convPython_convFilter) ProcessLiteralChar(_env *LnsEnv, node *Nodes_LiteralCharNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, _env.GetVM().String_format("%d", Lns_2DDD(node.FP.Get_num(_env))))
}
// 5939: decl @lune.@base.@convPython.convFilter.processLiteralInt
func (self *convPython_convFilter) ProcessLiteralInt(_env *LnsEnv, node *Nodes_LiteralIntNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, node.FP.Get_token(_env).Txt)
}
// 5945: decl @lune.@base.@convPython.convFilter.processLiteralReal
func (self *convPython_convFilter) ProcessLiteralReal(_env *LnsEnv, node *Nodes_LiteralRealNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, node.FP.Get_token(_env).Txt)
}
// 5951: decl @lune.@base.@convPython.convFilter.processLiteralString
func (self *convPython_convFilter) ProcessLiteralString(_env *LnsEnv, node *Nodes_LiteralStringNode,_opt LnsAny) {
    var txt string
    txt = node.FP.Get_token(_env).Txt
    {
        _expList := node.FP.Get_dddParam(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.WriteRaw(_env, convPython_str2gostr_14_(_env, txt))
            self.FP.WriteRaw(_env, "%(")
            convPython_filter_7_(_env, &expList.Nodes_Node, self, &node.Nodes_Node)
            self.FP.WriteRaw(_env, ")")
        } else {
            self.FP.WriteRaw(_env, _env.GetVM().String_format("%s", Lns_2DDD(convPython_str2gostr_14_(_env, txt))))
        }
    }
}
// 5965: decl @lune.@base.@convPython.convFilter.processLiteralBool
func (self *convPython_convFilter) ProcessLiteralBool(_env *LnsEnv, node *Nodes_LiteralBoolNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, node.FP.Get_token(_env).Txt)
}
// 5971: decl @lune.@base.@convPython.convFilter.processLiteralNil
func (self *convPython_convFilter) ProcessLiteralNil(_env *LnsEnv, node *Nodes_LiteralNilNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, convPython_literalNil)
}
// 5977: decl @lune.@base.@convPython.convFilter.processBreak
func (self *convPython_convFilter) ProcessBreak(_env *LnsEnv, node *Nodes_BreakNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, "break")
    self.FP.Writeln(_env, "")
}
// 5984: decl @lune.@base.@convPython.convFilter.processLiteralSymbol
func (self *convPython_convFilter) ProcessLiteralSymbol(_env *LnsEnv, node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    self.FP.WriteRaw(_env, node.FP.Get_token(_env).Txt)
}
// 5990: decl @lune.@base.@convPython.convFilter.processLuneControl
func (self *convPython_convFilter) ProcessLuneControl(_env *LnsEnv, node *Nodes_LuneControlNode,_opt LnsAny) {
}
// 6026: decl @lune.@base.@convPython.convFilter.processAbbr
func (self *convPython_convFilter) ProcessAbbr(_env *LnsEnv, node *Nodes_AbbrNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convPython.convFilter.processAbbr"
    Util_err(_env, _env.GetVM().String_format("not support -- %s:%d", Lns_2DDD(__func__, node.FP.Get_pos(_env).LineNo)))
}
// 1821: decl @lune.@base.@convPython.ConvRunner.run
func (self *convPython_ConvRunner) Run(_env *LnsEnv) {
    self.FP.Setup(_env)
    self.FP.pushProcessMode(_env, convPython_ProcessMode__DeclClass)
    LnsLog(_env, _env.GetVM().String_format("class fields div (%d)", Lns_2DDD(self.declMethodItemList.Len())))
    for _, _info := range( self.declMethodItemList.Items ) {
        info := _info
        convPython_filter_7_(_env, &info.FP.Get_fieldNode(_env).Nodes_Node, &self.convPython_convFilter, &info.FP.Get_classNode(_env).Nodes_Node)
    }
    self.FP.popProcessMode(_env)
}
// 1833: decl @lune.@base.@convPython.ConvRunner.getResult
func (self *convPython_ConvRunner) GetResult(_env *LnsEnv) string {
    LnsJoin(_env, self.FP)
    var memStream *Util_memStream
    
    {
        _memStream := Util_memStreamDownCastF(self.orgStream)
        if _memStream == nil{
            Util_err(_env, "convert err ")
        } else {
            memStream = _memStream.(*Util_memStream)
        }
    }
    return memStream.FP.Get_txt(_env)
}
