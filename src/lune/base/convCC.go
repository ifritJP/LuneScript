// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_convCC bool
var convCC__mod__ string
// decl enum -- ConvMode 
const ConvCC_ConvMode__ConvMeta = 2
const ConvCC_ConvMode__Convert = 1
const ConvCC_ConvMode__Exec = 0
var ConvCC_ConvModeList_ = NewLnsList( []LnsAny {
  ConvCC_ConvMode__Exec,
  ConvCC_ConvMode__Convert,
  ConvCC_ConvMode__ConvMeta,
})
func ConvCC_ConvMode_get__allList() *LnsList{
    return ConvCC_ConvModeList_
}
var ConvCC_ConvModeMap_ = map[LnsInt]string {
  ConvCC_ConvMode__ConvMeta: "ConvMode.ConvMeta",
  ConvCC_ConvMode__Convert: "ConvMode.Convert",
  ConvCC_ConvMode__Exec: "ConvMode.Exec",
}
func ConvCC_ConvMode__from(arg1 LnsInt) LnsAny{
    if _, ok := ConvCC_ConvModeMap_[arg1]; ok { return arg1 }
    return nil
}

func ConvCC_ConvMode_getTxt(arg1 LnsInt) string {
    return ConvCC_ConvModeMap_[arg1];
}
// decl enum -- ValKind 
const convCC_ValKind__Any = 1
const convCC_ValKind__Other = 5
const convCC_ValKind__Prim = 0
const convCC_ValKind__Stem = 3
const convCC_ValKind__StemWork = 2
const convCC_ValKind__Var = 4
var convCC_ValKindList_ = NewLnsList( []LnsAny {
  convCC_ValKind__Prim,
  convCC_ValKind__Any,
  convCC_ValKind__StemWork,
  convCC_ValKind__Stem,
  convCC_ValKind__Var,
  convCC_ValKind__Other,
})
func convCC_ValKind_get__allList_1164_() *LnsList{
    return convCC_ValKindList_
}
var convCC_ValKindMap_ = map[LnsInt]string {
  convCC_ValKind__Any: "ValKind.Any",
  convCC_ValKind__Other: "ValKind.Other",
  convCC_ValKind__Prim: "ValKind.Prim",
  convCC_ValKind__Stem: "ValKind.Stem",
  convCC_ValKind__StemWork: "ValKind.StemWork",
  convCC_ValKind__Var: "ValKind.Var",
}
func convCC_ValKind__from_1157_(arg1 LnsInt) LnsAny{
    if _, ok := convCC_ValKindMap_[arg1]; ok { return arg1 }
    return nil
}

func convCC_ValKind_getTxt(arg1 LnsInt) string {
    return convCC_ValKindMap_[arg1];
}
// decl enum -- ProcessMode 
const convCC_ProcessMode__DefClass = 6
const convCC_ProcessMode__DefWrap = 10
const convCC_ProcessMode__Form = 7
const convCC_ProcessMode__Immediate = 8
const convCC_ProcessMode__Include = 0
const convCC_ProcessMode__InitFuncSym = 9
const convCC_ProcessMode__InitModule = 3
const convCC_ProcessMode__Intermediate = 4
const convCC_ProcessMode__Prototype = 1
const convCC_ProcessMode__StringFormat = 5
const convCC_ProcessMode__WideScopeVer = 2
var convCC_ProcessModeList_ = NewLnsList( []LnsAny {
  convCC_ProcessMode__Include,
  convCC_ProcessMode__Prototype,
  convCC_ProcessMode__WideScopeVer,
  convCC_ProcessMode__InitModule,
  convCC_ProcessMode__Intermediate,
  convCC_ProcessMode__StringFormat,
  convCC_ProcessMode__DefClass,
  convCC_ProcessMode__Form,
  convCC_ProcessMode__Immediate,
  convCC_ProcessMode__InitFuncSym,
  convCC_ProcessMode__DefWrap,
})
func convCC_ProcessMode_get__allList_1213_() *LnsList{
    return convCC_ProcessModeList_
}
var convCC_ProcessModeMap_ = map[LnsInt]string {
  convCC_ProcessMode__DefClass: "ProcessMode.DefClass",
  convCC_ProcessMode__DefWrap: "ProcessMode.DefWrap",
  convCC_ProcessMode__Form: "ProcessMode.Form",
  convCC_ProcessMode__Immediate: "ProcessMode.Immediate",
  convCC_ProcessMode__Include: "ProcessMode.Include",
  convCC_ProcessMode__InitFuncSym: "ProcessMode.InitFuncSym",
  convCC_ProcessMode__InitModule: "ProcessMode.InitModule",
  convCC_ProcessMode__Intermediate: "ProcessMode.Intermediate",
  convCC_ProcessMode__Prototype: "ProcessMode.Prototype",
  convCC_ProcessMode__StringFormat: "ProcessMode.StringFormat",
  convCC_ProcessMode__WideScopeVer: "ProcessMode.WideScopeVer",
}
func convCC_ProcessMode__from_1206_(arg1 LnsInt) LnsAny{
    if _, ok := convCC_ProcessModeMap_[arg1]; ok { return arg1 }
    return nil
}

func convCC_ProcessMode_getTxt(arg1 LnsInt) string {
    return convCC_ProcessModeMap_[arg1];
}
// decl enum -- Out2HMode 
const convCC_Out2HMode__HeaderPri = 1
const convCC_Out2HMode__HeaderPub = 0
const convCC_Out2HMode__SourcePri = 3
const convCC_Out2HMode__SourcePub = 2
var convCC_Out2HModeList_ = NewLnsList( []LnsAny {
  convCC_Out2HMode__HeaderPub,
  convCC_Out2HMode__HeaderPri,
  convCC_Out2HMode__SourcePub,
  convCC_Out2HMode__SourcePri,
})
func convCC_Out2HMode_get__allList_1515_() *LnsList{
    return convCC_Out2HModeList_
}
var convCC_Out2HModeMap_ = map[LnsInt]string {
  convCC_Out2HMode__HeaderPri: "Out2HMode.HeaderPri",
  convCC_Out2HMode__HeaderPub: "Out2HMode.HeaderPub",
  convCC_Out2HMode__SourcePri: "Out2HMode.SourcePri",
  convCC_Out2HMode__SourcePub: "Out2HMode.SourcePub",
}
func convCC_Out2HMode__from_1508_(arg1 LnsInt) LnsAny{
    if _, ok := convCC_Out2HModeMap_[arg1]; ok { return arg1 }
    return nil
}

func convCC_Out2HMode_getTxt(arg1 LnsInt) string {
    return convCC_Out2HModeMap_[arg1];
}
// decl enum -- FuncWrap 
const convCC_FuncWrap__CallWrap = 1
const convCC_FuncWrap__NilWrap = 2
const convCC_FuncWrap__Normal = 0
var convCC_FuncWrapList_ = NewLnsList( []LnsAny {
  convCC_FuncWrap__Normal,
  convCC_FuncWrap__CallWrap,
  convCC_FuncWrap__NilWrap,
})
func convCC_FuncWrap_get__allList_2008_() *LnsList{
    return convCC_FuncWrapList_
}
var convCC_FuncWrapMap_ = map[LnsInt]string {
  convCC_FuncWrap__CallWrap: "FuncWrap.CallWrap",
  convCC_FuncWrap__NilWrap: "FuncWrap.NilWrap",
  convCC_FuncWrap__Normal: "FuncWrap.Normal",
}
func convCC_FuncWrap__from_2001_(arg1 LnsInt) LnsAny{
    if _, ok := convCC_FuncWrapMap_[arg1]; ok { return arg1 }
    return nil
}

func convCC_FuncWrap_getTxt(arg1 LnsInt) string {
    return convCC_FuncWrapMap_[arg1];
}
var convCC_cTypeInt string
var convCC_cTypeReal string
var convCC_cTypeBool string
var convCC_cTypeStem string
var convCC_cTypeAny string
var convCC_cTypeAnyP string
var convCC_cTypeAnyPP string
var convCC_cTypeEnvP string
var convCC_cTypeVarP string
var convCC_cTypeMod string
var convCC_cTypeModP string
var convCC_cTypeBlockP string
var convCC_cValNil string
var convCC_cValNone string
var convCC_cValDDD0 string
var convCC_accessAny string
var convCC_stepIndent LnsInt
var convCC_scopeAccess LnsInt
var convCC_invalidSymbolId LnsInt
type convCC_DstInfo__Node struct{
Val1 *Nodes_Node
Val2 bool
}
func (self *convCC_DstInfo__Node) GetTxt() string {
return "DstInfo.Node"
}
type convCC_DstInfo__Symbol struct{
Val1 Ast_LowSymbol
Val2 LnsAny
Val3 bool
}
func (self *convCC_DstInfo__Symbol) GetTxt() string {
return "DstInfo.Symbol"
}
type convCC_CollectionKind__Array struct{
}
var convCC_CollectionKind__Array_Obj = &convCC_CollectionKind__Array{}
func (self *convCC_CollectionKind__Array) GetTxt() string {
return "CollectionKind.Array"
}
type convCC_CollectionKind__ExtArray struct{
Val1 *Ast_TypeInfo
}
func (self *convCC_CollectionKind__ExtArray) GetTxt() string {
return "CollectionKind.ExtArray"
}
type convCC_CollectionKind__ExtList struct{
Val1 *Ast_TypeInfo
}
func (self *convCC_CollectionKind__ExtList) GetTxt() string {
return "CollectionKind.ExtList"
}
type convCC_CollectionKind__ExtMap struct{
Val1 *Ast_TypeInfo
}
func (self *convCC_CollectionKind__ExtMap) GetTxt() string {
return "CollectionKind.ExtMap"
}
type convCC_CollectionKind__ExtSet struct{
Val1 *Ast_TypeInfo
}
func (self *convCC_CollectionKind__ExtSet) GetTxt() string {
return "CollectionKind.ExtSet"
}
type convCC_CollectionKind__List struct{
}
var convCC_CollectionKind__List_Obj = &convCC_CollectionKind__List{}
func (self *convCC_CollectionKind__List) GetTxt() string {
return "CollectionKind.List"
}
type convCC_CollectionKind__Map struct{
}
var convCC_CollectionKind__Map_Obj = &convCC_CollectionKind__Map{}
func (self *convCC_CollectionKind__Map) GetTxt() string {
return "CollectionKind.Map"
}
type convCC_CollectionKind__Set struct{
}
var convCC_CollectionKind__Set_Obj = &convCC_CollectionKind__Set{}
func (self *convCC_CollectionKind__Set) GetTxt() string {
return "CollectionKind.Set"
}
type convCC_MRetInfo__DDD struct{
Val1 *Nodes_ExpListNode
}
func (self *convCC_MRetInfo__DDD) GetTxt() string {
return "MRetInfo.DDD"
}
type convCC_MRetInfo__Form struct{
}
var convCC_MRetInfo__Form_Obj = &convCC_MRetInfo__Form{}
func (self *convCC_MRetInfo__Form) GetTxt() string {
return "MRetInfo.Form"
}
type convCC_MRetInfo__FormFunc struct{
Val1 *Ast_TypeInfo
}
func (self *convCC_MRetInfo__FormFunc) GetTxt() string {
return "MRetInfo.FormFunc"
}
type convCC_MRetInfo__Format struct{
Val1 string
Val2 *Nodes_ExpListNode
}
func (self *convCC_MRetInfo__Format) GetTxt() string {
return "MRetInfo.Format"
}
type convCC_MRetInfo__Func struct{
Val1 *Nodes_Node
}
func (self *convCC_MRetInfo__Func) GetTxt() string {
return "MRetInfo.Func"
}
type convCC_MRetInfo__Method struct{
Val1 *Ast_TypeInfo
}
func (self *convCC_MRetInfo__Method) GetTxt() string {
return "MRetInfo.Method"
}
type convCC_processExp_1881_ func ()
type convCC_process2stemCallback_2057_ func ()
type convCC_outputMacroStmtBlock_2219_ func ()
type convCC_processRValue_2298_ func ()
type convCC_ProcessToValForm_2538_ func (arg1 LnsInt)
// for 5418: ExpCast
func conv2Form24644( src func ()) LnsForm {
    return func (argList []LnsAny) []LnsAny {
        src()
        return []LnsAny{}
    }
}
// for 362
func convCC_convExp854(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 367
func convCC_convExp921(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 10418
func convCC_convExp44564(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1648
func convCC_convExp7157(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1986
func convCC_convExp9781(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 4768
func convCC_convExp22103(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 6230
func convCC_convExp27616(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 6408
func convCC_convExp28556(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 6409
func convCC_convExp28570(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 10779
func convCC_convExp46440(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 10780
func convCC_convExp46454(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}

// 81: decl @lune.@base.@convCC.isClosure
func convCC_isClosure_1011_(funcType *Ast_TypeInfo) bool {
    {
        _scope := funcType.FP.Get_scope()
        if _scope != nil {
            scope := _scope.(*Ast_Scope)
            return scope.FP.Get_closureSymList().Len() > 0
        }
    }
    return false
}

// 210: decl @lune.@base.@convCC.getValKind
func convCC_getValKind_1167_(valType *Ast_TypeInfo) LnsInt {
    var expType *Ast_TypeInfo
    expType = valType.FP.Get_srcTypeInfo()
    if expType.FP.Get_nilable(){
        return convCC_ValKind__Stem
    }
    if _switch466 := expType.FP.Get_kind(); _switch466 == Ast_TypeInfoKind__Alternate || _switch466 == Ast_TypeInfoKind__Stem || _switch466 == Ast_TypeInfoKind__DDD || _switch466 == Ast_TypeInfoKind__Alge {
        return convCC_ValKind__Stem
    }
    if _switch522 := expType; _switch522 == Ast_builtinTypeInt || _switch522 == Ast_builtinTypeChar {
        return convCC_ValKind__Prim
    } else if _switch522 == Ast_builtinTypeReal {
        return convCC_ValKind__Prim
    } else if _switch522 == Ast_builtinTypeBool {
        return convCC_ValKind__Prim
    } else {
        {
            _enumType := Ast_EnumTypeInfoDownCastF(expType.FP)
            if _enumType != nil {
                enumType := _enumType.(*Ast_EnumTypeInfo)
                return convCC_getValKind_1167_(enumType.FP.Get_valTypeInfo())
            }
        }
        return convCC_ValKind__Any
    }
// insert a dummy
    return 0
}

// 247: decl @lune.@base.@convCC.isStemType
func convCC_isStemType_1171_(valType *Ast_TypeInfo) bool {
    return convCC_getValKind_1167_(valType) == convCC_ValKind__Stem
}

// 252: decl @lune.@base.@convCC.getRetKind
func convCC_getRetKind_1178_(retTypeList *LnsList) LnsInt {
    if _switch582 := retTypeList.Len(); _switch582 == 0 {
        return convCC_ValKind__Other
    } else if _switch582 == 1 {
        return convCC_getValKind_1167_(retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
    }
    return convCC_ValKind__Stem
}


// 275: decl @lune.@base.@convCC.getCType
func convCC_getCType_1188_(valType *Ast_TypeInfo) string {
    var expType *Ast_TypeInfo
    expType = valType.FP.Get_srcTypeInfo()
    if _switch697 := expType; _switch697 == Ast_builtinTypeInt || _switch697 == Ast_builtinTypeChar {
        return convCC_cTypeInt
    } else if _switch697 == Ast_builtinTypeReal {
        return convCC_cTypeReal
    } else if _switch697 == Ast_builtinTypeBool {
        return convCC_cTypeBool
    } else {
        {
            _enumType := Ast_EnumTypeInfoDownCastF(expType.FP)
            if _enumType != nil {
                enumType := _enumType.(*Ast_EnumTypeInfo)
                return convCC_getCType_1188_(enumType.FP.Get_valTypeInfo())
            }
        }
        if convCC_getValKind_1167_(valType) == convCC_ValKind__Any{
            return convCC_cTypeAnyP
        }
        return convCC_cTypeStem
    }
// insert a dummy
    return ""
}

// 304: decl @lune.@base.@convCC.getCRetType
func convCC_getCRetType_1195_(retTypeList *LnsList) string {
    if _switch736 := retTypeList.Len(); _switch736 == 0 {
        return "void"
    } else if _switch736 == 1 {
        return convCC_getCType_1188_(retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
    }
    return convCC_cTypeStem
}

// 316: decl @lune.@base.@convCC.getBlockName
func convCC_getBlockName_1198_(scope *Ast_Scope) string {
    return Lns_getVM().String_format("pBlock_%X", []LnsAny{scope.FP.Get_scopeId()})
}

// 346: decl @lune.@base.@convCC.isClassMember
func convCC_isClassMember_1216_(symbol Ast_LowSymbol) bool {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(symbol.Get_scope().FP.Get_ownerTypeInfo()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.Get_kind()})) == Ast_TypeInfoKind__Class) &&
        Lns_GetEnv().SetStackVal( symbol.Get_kind() == Ast_SymbolKind__Mbr) &&
        Lns_GetEnv().SetStackVal( symbol.Get_staticFlag()) ).(bool)){
        return true
    }
    return false
}

// 359: decl @lune.@base.@convCC.str2cstr
func convCC_str2cstr_1219_(txt string) string {
    var work string
    work = txt
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(work, "^```", nil, nil))){
        work = convCC_convExp854(Lns_2DDD(Lns_getVM().String_gsub((Lns_getVM().String_format("%q", []LnsAny{Lns_getVM().String_sub(work,4, -4)})),"\\\n", "\\n")))
        
    } else if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(work, "^'", nil, nil))){
        work = Lns_getVM().String_format("\"%s\"", []LnsAny{Lns_car(Lns_getVM().String_gsub((Lns_getVM().String_format("%s", []LnsAny{Lns_getVM().String_sub(work,2, -2)})),"\"", "\\\"")).(string)})
        
    }
    work = convCC_convExp921(Lns_2DDD(Lns_getVM().String_gsub(work,"\\9", "\\t")))
    
    return work
}

// 718: decl @lune.@base.@convCC.getOrgTypeInfo
func convCC_getOrgTypeInfo_1430_(typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    {
        _enumType := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo().FP.Get_nonnilableType().FP)
        if _enumType != nil {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            return enumType.FP.Get_valTypeInfo()
        }
    }
    return typeInfo.FP.Get_srcTypeInfo().FP.Get_nonnilableType()
}

// 725: decl @lune.@base.@convCC.getAccessPrimValFromSymbolDirect
func convCC_getAccessPrimValFromSymbolDirect_1433_(symName string,valKind LnsInt,symType *Ast_TypeInfo) string {
    var txt string
    txt = symName
    if _switch3235 := valKind; _switch3235 == convCC_ValKind__Var {
        txt = txt + "->stem"
        
    } else if _switch3235 == convCC_ValKind__Stem {
    } else if _switch3235 == convCC_ValKind__Prim {
        return txt
    }
    if _switch3290 := convCC_getOrgTypeInfo_1430_(symType); _switch3290 == Ast_builtinTypeInt || _switch3290 == Ast_builtinTypeChar {
        txt = txt + ".val.intVal"
        
    } else if _switch3290 == Ast_builtinTypeReal {
        txt = txt + ".val.realVal"
        
    } else if _switch3290 == Ast_builtinTypeBool {
        txt = txt + ".val.boolVal"
        
    }
    return txt
}

// 755: decl @lune.@base.@convCC.createSymbolParam
func convCC_createSymbolParam_1436_(name string,valKind LnsInt,cTypeTxt string) *convCC_SymbolParam {
    if _switch3399 := valKind; _switch3399 == convCC_ValKind__Stem {
        return NewconvCC_SymbolParam(convCC_ValKind__Stem, 0, convCC_cTypeStem)
    } else if _switch3399 == convCC_ValKind__Any {
        return NewconvCC_SymbolParam(convCC_ValKind__Any, 0, convCC_cTypeAnyP)
    } else if _switch3399 == convCC_ValKind__Prim {
        return NewconvCC_SymbolParam(convCC_ValKind__Prim, 0, cTypeTxt)
    } else if _switch3399 == convCC_ValKind__Other {
        return NewconvCC_SymbolParam(convCC_ValKind__Other, 0, "void")
    } else {
        Util_err(Lns_getVM().String_format("not support %s:%s", []LnsAny{name, convCC_ValKind_getTxt( valKind)}))
    }
// insert a dummy
    return nil
}

// 1003: decl @lune.@base.@convCC.getLiteralStrAny
func convCC_getLiteralStrAny_1497_(txt string) string {
    return Lns_getVM().String_format("lns_litStr2any( _pEnv, %s )", []LnsAny{txt})
}

// 1007: decl @lune.@base.@convCC.getLiteralStrStem
func convCC_getLiteralStrStem_1500_(txt string) string {
    return Lns_getVM().String_format("LNS_STEM_ANY( %s )", []LnsAny{convCC_getLiteralStrAny_1497_(txt)})
}

// 1019: decl @lune.@base.@convCC.getOut2HeaderPrefix
func convCC_getOut2HeaderPrefix_1518_(mode LnsInt) string {
    if _switch4498 := mode; _switch4498 == convCC_Out2HMode__HeaderPub {
        return "extern "
    } else if _switch4498 == convCC_Out2HMode__SourcePri {
        return "static "
    }
    return ""
}

// 1213: decl @lune.@base.@convCC.processAddModuleGlobal
func convCC_processAddModuleGlobal_1639_(stream Util_SourceStream,valName string) {
    stream.Writeln(Lns_getVM().String_format("lns_mtd_List_insert( _pEnv, *lns_module_globalStemList, %s );", []LnsAny{valName}))
}

// 1221: decl @lune.@base.@convCC.filter
func convCC_filter_1642_(node *Nodes_Node,filter *convCC_convFilter,parent *Nodes_Node) {
    node.FP.ProcessFilter(&filter.Nodes_Filter, ConvCC_Opt2Stem(NewConvCC_Opt(parent)))
}



// 1289: decl @lune.@base.@convCC.getSymbolIndex
func convCC_getSymbolIndex_1657_(symbol *Ast_SymbolInfo) LnsInt {
    var param LnsAny
    
    {
        _param := symbol.FP.Get_convModuleParam()
        if _param == nil{
            return 0
        } else {
            param = _param
        }
    }
    return (param.(*convCC_SymbolParam)).Index
}



// 1496: decl @lune.@base.@convCC.registerBuiltin
func convCC_registerBuiltin_1761_() {
    var builtin *TransUnit_BuiltinFuncType
    builtin = TransUnit_getBuiltinFunc()
    for _, _symbol := range( builtin.FP.Get_allSymbol().Items ) {
        symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        var param *convCC_SymbolParam
        if _switch6377 := symbol.FP.Get_kind(); _switch6377 == Ast_SymbolKind__Mtd || _switch6377 == Ast_SymbolKind__Fun {
            var retTypeList *LnsList
            retTypeList = symbol.FP.Get_typeInfo().FP.Get_retTypeInfoList()
            param = convCC_createSymbolParam_1436_(symbol.FP.Get_name(), convCC_getRetKind_1178_(retTypeList), convCC_getCRetType_1195_(retTypeList))
            
        } else if _switch6377 == Ast_SymbolKind__Mbr || _switch6377 == Ast_SymbolKind__Var {
            param = convCC_createSymbolParam_1436_(symbol.FP.Get_name(), convCC_getValKind_1167_(symbol.FP.Get_typeInfo()), convCC_getCType_1188_(symbol.FP.Get_typeInfo()))
            
        } else {
            Util_err(Lns_getVM().String_format("illeal symbol -- %s %d", []LnsAny{symbol.FP.Get_name(), 1512}))
        }
        symbol.FP.Set_convModuleParam(param.FP)
    }
}




// 1867: decl @lune.@base.@convCC.getAccessPrimValFromStem
func convCC_getAccessPrimValFromStem_1851_(dddFlag bool,typeInfo *Ast_TypeInfo,index LnsInt) string {
    var txt string
    txt = ""
    if dddFlag{
        txt = Lns_getVM().String_format(".val.pAny->val.ddd.stemList[ %d ]", []LnsAny{index})
        
    }
    var expType *Ast_TypeInfo
    {
        _enumType := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo().FP)
        if _enumType != nil {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            expType = enumType.FP.Get_valTypeInfo()
            
        } else {
            expType = typeInfo.FP.Get_srcTypeInfo()
            
        }
    }
    if _switch9428 := expType; _switch9428 == Ast_builtinTypeInt || _switch9428 == Ast_builtinTypeChar {
        txt = txt + ".val.intVal"
        
    } else if _switch9428 == Ast_builtinTypeReal {
        txt = txt + ".val.realVal"
        
    } else if _switch9428 == Ast_builtinTypeBool {
        txt = txt + ".val.boolVal"
        
    } else {
        if convCC_getValKind_1167_(expType) == convCC_ValKind__Any{
            txt = txt + ".val.pAny"
            
        }
    }
    return txt
}

// 1901: decl @lune.@base.@convCC.getAccessValFromStem
func convCC_getAccessValFromStem_1854_(typeInfo *Ast_TypeInfo) string {
    var txt string
    var expType *Ast_TypeInfo
    {
        _enumType := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo().FP)
        if _enumType != nil {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            expType = enumType.FP.Get_valTypeInfo()
            
        } else {
            expType = typeInfo.FP.Get_srcTypeInfo()
            
        }
    }
    if _switch9543 := expType; _switch9543 == Ast_builtinTypeInt || _switch9543 == Ast_builtinTypeChar {
        txt = ".val.intVal"
        
    } else if _switch9543 == Ast_builtinTypeReal {
        txt = ".val.realVal"
        
    } else if _switch9543 == Ast_builtinTypeBool {
        txt = ".val.boolVal"
        
    } else {
        if convCC_getValKind_1167_(typeInfo) == convCC_ValKind__Any{
            txt = convCC_accessAny
            
        } else { 
            txt = ""
            
        }
    }
    return txt
}

// 2095: decl @lune.@base.@convCC.getLiteral2Stem
func convCC_getLiteral2Stem_1884_(valTxt string,typeInfo *Ast_TypeInfo) string {
    if _switch10288 := typeInfo.FP.Get_srcTypeInfo(); _switch10288 == Ast_builtinTypeInt || _switch10288 == Ast_builtinTypeChar {
        return Lns_getVM().String_format("LNS_STEM_INT( %s )", []LnsAny{valTxt})
    } else if _switch10288 == Ast_builtinTypeReal {
        return Lns_getVM().String_format("LNS_STEM_REAL( %s )", []LnsAny{valTxt})
    } else if _switch10288 == Ast_builtinTypeBool {
        return Lns_getVM().String_format("LNS_STEM_BOOL( %s )", []LnsAny{valTxt})
    } else {
        return "NULL"
    }
// insert a dummy
    return ""
}

// 2113: decl @lune.@base.@convCC.getStemTypeId
func convCC_getStemTypeId_1887_(typeInfo *Ast_TypeInfo) string {
    if typeInfo.FP.Get_nilable(){
        return "lns_stem_type_none"
    }
    if _switch10348 := typeInfo; _switch10348 == Ast_builtinTypeInt || _switch10348 == Ast_builtinTypeChar {
        return "lns_stem_type_int"
    } else if _switch10348 == Ast_builtinTypeInt || _switch10348 == Ast_builtinTypeReal {
        return "lns_stem_type_real"
    } else if _switch10348 == Ast_builtinTypeBool {
        return "lns_stem_type_bool"
    } else {
        return "lns_stem_type_any"
    }
// insert a dummy
    return ""
}

// 2133: decl @lune.@base.@convCC.getPrepareClosure
func convCC_getPrepareClosure_1894_(scopeMgr *convCC_ScopeMgr,funcName string,argNum LnsInt,hasDDD bool,symList *LnsList) string {
    var txt string
    txt = Lns_getVM().String_format("lns_func2any( _pEnv, (lns_closure_t *)%s, %d, %s, %d", []LnsAny{funcName, argNum, hasDDD, symList.Len()})
    
    for _, _symbolInfo := range( symList.Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        txt = txt + ", "
        
        txt = txt + scopeMgr.FP.Symbol2Any(symbolInfo.FP)
        
    }
    txt = txt + ")"
    
    return txt
}

// 2150: decl @lune.@base.@convCC.getFunc2any
func convCC_getFunc2any_1897_(moduleCtrl *convCC_ModuleCtrl,scopeMgr *convCC_ScopeMgr,funcType *Ast_TypeInfo) string {
    var argList *LnsList
    argList = funcType.FP.Get_argTypeInfoList()
    var hasDDD bool
    hasDDD = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( argList.Len() > 0) &&
        Lns_GetEnv().SetStackVal( argList.GetAt(argList.Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() == Ast_TypeInfoKind__DDD) ||
        Lns_GetEnv().SetStackVal( false) ).(bool)
    return convCC_getPrepareClosure_1894_(scopeMgr, moduleCtrl.FP.GetFuncName(funcType), funcType.FP.Get_argTypeInfoList().Len(), hasDDD, (Lns_unwrap( funcType.FP.Get_scope()).(*Ast_Scope)).FP.Get_closureSymList())
}




// 2632: decl @lune.@base.@convCC.processAlgeNewProto
func convCC_processAlgeNewProto_1925_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,typeInfo *Ast_TypeInfo,valInfo *Ast_AlgeValInfo) {
    stream.Write(Lns_getVM().String_format("%s %s( %s _pEnv", []LnsAny{convCC_cTypeStem, moduleCtrl.FP.GetNewAlgeCName(typeInfo, valInfo.FP.Get_name()), convCC_cTypeEnvP}))
    for _index, _valTypeInfo := range( valInfo.FP.Get_typeList().Items ) {
        index := _index + 1
        valTypeInfo := _valTypeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        stream.Write(Lns_getVM().String_format(", %s _val%d", []LnsAny{convCC_getCType_1188_(valTypeInfo), index}))
    }
    stream.Write(")")
}



// 2646: decl @lune.@base.@convCC.processAlgePrototype
func convCC_processAlgePrototype_1928_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,node *Nodes_DeclAlgeNode) {
    var algeType *Ast_AlgeTypeInfo
    algeType = node.FP.Get_algeType()
    var valList *LnsList
    valList = NewLnsList([]LnsAny{})
    {
        __collection12408 := algeType.FP.Get_valInfoMap()
        __sorted12408 := __collection12408.CreateKeyListStr()
        __sorted12408.Sort( LnsItemKindStr, nil )
        for _, ___key12408 := range( __sorted12408.Items ) {
            valInfo := __collection12408.Items[ ___key12408 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            valList.Insert(Ast_AlgeValInfo2Stem(valInfo))
        }
    }
    var process func()
    process = func() {
        stream.Writeln("typedef enum {")
        stream.PushIndent(nil)
        var enumName string
        enumName = moduleCtrl.FP.GetAlgeEnumCName(node.FP.Get_expType())
        for _index, _valInfo := range( valList.Items ) {
            index := _index + 1
            valInfo := _valInfo.(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            if index > 1{
                stream.Writeln(",")
            }
            stream.Write(Lns_getVM().String_format("%s_%s", []LnsAny{enumName, valInfo.FP.Get_name()}))
        }
        stream.Writeln("")
        stream.PopIndent()
        stream.Writeln(Lns_getVM().String_format("} %s;", []LnsAny{enumName}))
        for _, _valInfo := range( valList.Items ) {
            valInfo := _valInfo.(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            if valInfo.FP.Get_typeList().Len() > 0{
                stream.Writeln("typedef struct {")
                stream.PushIndent(nil)
                for _index, _typeInfo := range( valInfo.FP.Get_typeList().Items ) {
                    index := _index + 1
                    typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    stream.Writeln(Lns_getVM().String_format("%s _val%d;", []LnsAny{convCC_getCType_1188_(typeInfo), index}))
                }
                stream.PopIndent()
                stream.Writeln(Lns_getVM().String_format("} %s;", []LnsAny{moduleCtrl.FP.GetAlgeValStrCName(node.FP.Get_expType(), valInfo.FP.Get_name())}))
            }
        }
        for _, _valInfo := range( valList.Items ) {
            valInfo := _valInfo.(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            if valInfo.FP.Get_typeList().Len() > 0{
                convCC_processAlgeNewProto_1925_(stream, moduleCtrl, node.FP.Get_expType(), valInfo)
                stream.Writeln(";")
            }
        }
    }
    {
        var processwork func(out2HMode LnsInt)
        processwork = func(out2HMode LnsInt) {
            if _switch12674 := out2HMode; _switch12674 == convCC_Out2HMode__HeaderPub || _switch12674 == convCC_Out2HMode__SourcePri {
                process()
            }
        }
        if Ast_isPubToExternal(node.FP.Get_expType().FP.Get_accessMode()){
            stream.SwitchToHeader()
            processwork(convCC_Out2HMode__HeaderPub)
            stream.ReturnToSource()
            processwork(convCC_Out2HMode__SourcePub)
        } else { 
            processwork(convCC_Out2HMode__SourcePri)
        }
    }
    
    stream.Writeln(Lns_getVM().String_format("static void %s( %s _pEnv );", []LnsAny{moduleCtrl.FP.GetAlgeInitCName(node.FP.Get_expType()), convCC_cTypeEnvP}))
}



// 2716: decl @lune.@base.@convCC.processAlgeWideScope
func convCC_processAlgeWideScope_1944_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,node *Nodes_DeclAlgeNode) {
    var algeType *Ast_AlgeTypeInfo
    algeType = node.FP.Get_algeType()
    var valList *LnsList
    valList = NewLnsList([]LnsAny{})
    {
        __collection12797 := algeType.FP.Get_valInfoMap()
        __sorted12797 := __collection12797.CreateKeyListStr()
        __sorted12797.Sort( LnsItemKindStr, nil )
        for _, ___key12797 := range( __sorted12797.Items ) {
            valInfo := __collection12797.Items[ ___key12797 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            valList.Insert(Ast_AlgeValInfo2Stem(valInfo))
        }
    }
    var process func(out2HMode LnsInt)
    process = func(out2HMode LnsInt) {
        var prefix string
        prefix = convCC_getOut2HeaderPrefix_1518_(out2HMode)
        for _, _valInfo := range( valList.Items ) {
            valInfo := _valInfo.(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            if valInfo.FP.Get_typeList().Len() == 0{
                var varName string
                varName = moduleCtrl.FP.GetAlgeValCName(node.FP.Get_expType(), valInfo.FP.Get_name())
                stream.Writeln(Lns_getVM().String_format("%s%s %s;", []LnsAny{prefix, convCC_cTypeStem, varName}))
                stream.Writeln(Lns_getVM().String_format("%s%s %s_any;", []LnsAny{prefix, convCC_cTypeAny, varName}))
            }
        }
    }
    {
        var processwork func(out2HMode LnsInt)
        processwork = func(out2HMode LnsInt) {
            if _switch12938 := out2HMode; _switch12938 == convCC_Out2HMode__HeaderPub || _switch12938 == convCC_Out2HMode__SourcePri || _switch12938 == convCC_Out2HMode__SourcePub {
                process(out2HMode)
            }
        }
        if Ast_isPubToExternal(node.FP.Get_expType().FP.Get_accessMode()){
            stream.SwitchToHeader()
            processwork(convCC_Out2HMode__HeaderPub)
            stream.ReturnToSource()
            processwork(convCC_Out2HMode__SourcePub)
        } else { 
            processwork(convCC_Out2HMode__SourcePri)
        }
    }
    
    var algeTypeName string
    algeTypeName = moduleCtrl.FP.GetAlgeCName(node.FP.Get_expType())
    stream.Writeln(Lns_getVM().String_format("static %s %s_type2NameMap;", []LnsAny{convCC_cTypeAnyP, algeTypeName}))
}

// 2755: decl @lune.@base.@convCC.processAlgeForm
func convCC_processAlgeForm_1960_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,node *Nodes_DeclAlgeNode) {
    var algeType *Ast_AlgeTypeInfo
    algeType = node.FP.Get_algeType()
    var valList *LnsList
    valList = NewLnsList([]LnsAny{})
    {
        __collection13065 := algeType.FP.Get_valInfoMap()
        __sorted13065 := __collection13065.CreateKeyListStr()
        __sorted13065.Sort( LnsItemKindStr, nil )
        for _, ___key13065 := range( __sorted13065.Items ) {
            valInfo := __collection13065.Items[ ___key13065 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            valList.Insert(Ast_AlgeValInfo2Stem(valInfo))
        }
    }
    var algeName string
    algeName = moduleCtrl.FP.GetAlgeCName(&algeType.Ast_TypeInfo)
    var type2NameMapName string
    type2NameMapName = Lns_getVM().String_format("%s_type2NameMap", []LnsAny{algeName})
    if Lns_op_not(Ast_isPubToExternal(algeType.FP.Get_accessMode())){
        stream.Write("static ")
    }
    stream.Writeln(Lns_getVM().String_format("%s %s_get__txt( %s _pEnv, %s pAny ) {", []LnsAny{convCC_cTypeAnyP, algeName, convCC_cTypeEnvP, convCC_cTypeAnyP}))
    stream.PushIndent(nil)
    stream.Writeln(Lns_getVM().String_format("return lns_mtd_Map_get( _pEnv, %s, LNS_STEM_INT( pAny->val.alge.type ) )%s;", []LnsAny{type2NameMapName, convCC_accessAny}))
    stream.PopIndent()
    stream.Writeln("}")
    stream.Writeln(Lns_getVM().String_format("static void %s( %s _pEnv ) {", []LnsAny{moduleCtrl.FP.GetAlgeInitCName(&algeType.Ast_TypeInfo), convCC_cTypeEnvP}))
    stream.PushIndent(nil)
    stream.Writeln(type2NameMapName + " = lns_class_Map_new( _pEnv );")
    convCC_processAddModuleGlobal_1639_(stream, Lns_getVM().String_format("LNS_STEM_ANY( %s )", []LnsAny{type2NameMapName}))
    var fullName string
    fullName = moduleCtrl.FP.GetFullName(&algeType.Ast_TypeInfo)
    var enumName string
    enumName = moduleCtrl.FP.GetAlgeEnumCName(&algeType.Ast_TypeInfo)
    for _, _valInfo := range( valList.Items ) {
        valInfo := _valInfo.(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
        stream.Write(Lns_getVM().String_format("lns_mtd_Map_add( _pEnv, %s, LNS_STEM_INT( %s_%s ), ", []LnsAny{type2NameMapName, enumName, valInfo.FP.Get_name()}))
        stream.Writeln(Lns_getVM().String_format("LNS_STEM_ANY( lns_litStr2any( _pEnv, \"%s.%s\" ) ) );", []LnsAny{fullName, valInfo.FP.Get_name()}))
    }
    for _, _valInfo := range( valList.Items ) {
        valInfo := _valInfo.(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
        if valInfo.FP.Get_typeList().Len() == 0{
            var varName string
            varName = moduleCtrl.FP.GetAlgeValCName(&algeType.Ast_TypeInfo, valInfo.FP.Get_name())
            stream.Writeln(Lns_getVM().String_format("lns_init_alge( &%s, &%s_any, %s_%s );", []LnsAny{varName, varName, enumName, valInfo.FP.Get_name()}))
        }
    }
    stream.PopIndent()
    stream.Writeln("}")
    for _, _valInfo := range( valList.Items ) {
        valInfo := _valInfo.(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
        if valInfo.FP.Get_typeList().Len() > 0{
            var hasAnyFlag bool
            hasAnyFlag = false
            for _, _valType := range( valInfo.FP.Get_typeList().Items ) {
                valType := _valType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if convCC_isStemType_1171_(valType){
                    hasAnyFlag = true
                    
                    break
                }
            }
            var valStruct string
            valStruct = moduleCtrl.FP.GetAlgeValStrCName(&algeType.Ast_TypeInfo, valInfo.FP.Get_name())
            var gcTxt string
            if hasAnyFlag{
                gcTxt = Lns_getVM().String_format("lns_gc_alge_%s_%s", []LnsAny{algeName, valInfo.FP.Get_name()})
                
                stream.Writeln(Lns_getVM().String_format("static void %s( %s _pEnv, void * pVal ) {", []LnsAny{gcTxt, convCC_cTypeEnvP}))
                stream.PushIndent(nil)
                stream.Writeln(Lns_getVM().String_format("%s *pWorkVal = (%s *)pVal;", []LnsAny{valStruct, valStruct}))
                for _paramIndex, _valType := range( valInfo.FP.Get_typeList().Items ) {
                    paramIndex := _paramIndex + 1
                    valType := _valType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if convCC_isStemType_1171_(valType){
                        stream.Writeln(Lns_getVM().String_format("lns_decre_ref( _pEnv, pWorkVal->_val%d.val.pAny );", []LnsAny{paramIndex}))
                    }
                }
                stream.PopIndent()
                stream.Writeln("}")
            } else { 
                gcTxt = "NULL"
                
            }
            convCC_processAlgeNewProto_1925_(stream, moduleCtrl, &algeType.Ast_TypeInfo, valInfo)
            stream.Writeln("{")
            stream.PushIndent(nil)
            stream.Writeln(Lns_getVM().String_format("%s pAny = lns_alge_new( _pEnv, %s_%s, sizeof( %s ), %s );", []LnsAny{convCC_cTypeAnyP, enumName, valInfo.FP.Get_name(), valStruct, gcTxt}))
            stream.Writeln(Lns_getVM().String_format("%s *pVal = pAny->val.alge.pVal;", []LnsAny{valStruct}))
            for _paramIndex, _valType := range( valInfo.FP.Get_typeList().Items ) {
                paramIndex := _paramIndex + 1
                valType := _valType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if convCC_isStemType_1171_(valType){
                    stream.Writeln(Lns_getVM().String_format("lns_setQ( pVal->_val%d, _val%d );", []LnsAny{paramIndex, paramIndex}))
                } else { 
                    stream.Writeln(Lns_getVM().String_format("pVal->_val%d = _val%d;", []LnsAny{paramIndex, paramIndex}))
                }
            }
            stream.Writeln("return LNS_STEM_ANY( pAny );")
            stream.PopIndent()
            stream.Writeln("}")
        }
    }
}

// 3035: decl @lune.@base.@convCC.getMethodTypeTxt
func convCC_getMethodTypeTxt_1987_(retTypeList *LnsList) string {
    if retTypeList.Len() == 1{
        var retType *Ast_TypeInfo
        retType = retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_srcTypeInfo()
        {
            _enumType := Ast_EnumTypeInfoDownCastF(retType.FP)
            if _enumType != nil {
                enumType := _enumType.(*Ast_EnumTypeInfo)
                retType = enumType.FP.Get_valTypeInfo()
                
            }
        }
        if _switch13978 := retType; _switch13978 == Ast_builtinTypeInt || _switch13978 == Ast_builtinTypeChar {
            return "lns_method_int_t"
        } else if _switch13978 == Ast_builtinTypeReal {
            return "lns_method_real_t"
        } else if _switch13978 == Ast_builtinTypeBool {
            return "lns_method_bool_t"
        }
        if convCC_getValKind_1167_(retType) == convCC_ValKind__Any{
            return "lns_method_any_t"
        }
    }
    return "lns_method_t"
}

// 3059: decl @lune.@base.@convCC.processNewConstrProto
func convCC_processNewConstrProto_1990_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,node *Nodes_DeclClassNode,out2HMode LnsInt,outputBuiltinFlag bool) {
    stream.Write(convCC_getOut2HeaderPrefix_1518_(out2HMode))
    stream.Write(Lns_getVM().String_format("%s %s( %s _pEnv", []LnsAny{convCC_cTypeAnyP, moduleCtrl.FP.GetNewName(node.FP.Get_expType()), convCC_cTypeEnvP}))
    if Lns_op_not(outputBuiltinFlag){
        var scope *Ast_Scope
        scope = Lns_unwrap( node.FP.Get_expType().FP.Get_scope()).(*Ast_Scope)
        var initFuncType *Ast_TypeInfo
        initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField("__init", true, scope, convCC_scopeAccess)).(*Ast_TypeInfo)
        for _index, _argType := range( initFuncType.FP.Get_argTypeInfoList().Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            stream.Write(Lns_getVM().String_format(", %s arg%d", []LnsAny{convCC_getCType_1188_(argType), index}))
        }
    }
    stream.Write(")")
}

// 3081: decl @lune.@base.@convCC.processDeclAlgeSub
func convCC_processDeclAlgeSub_1993_(stream Util_SourceStream,node *Nodes_DeclArgNode) {
    stream.Write(convCC_getCType_1188_(node.FP.Get_expType()))
    if node.FP.Get_symbolInfo().FP.Get_hasAccessFromClosure(){
        stream.Write(" _")
    } else if node.FP.Get_symbolInfo().FP.Get_mutable(){
        stream.Write(" _")
    } else { 
        stream.Write(" ")
    }
    stream.Write(node.FP.Get_name().Txt)
}

// 3110: decl @lune.@base.@convCC.processMethodDeclTxt
func convCC_processMethodDeclTxt_2015_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,wrapKind LnsInt,methodTypeInfo *Ast_TypeInfo,argList LnsAny) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( methodTypeInfo.FP.Get_rawTxt() != "__init") &&
        Lns_GetEnv().SetStackVal( wrapKind == convCC_FuncWrap__Normal) ).(bool)){
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(methodTypeInfo.FP.Get_staticFlag())) ||
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(methodTypeInfo.FP.Get_accessMode()))) ||
            Lns_GetEnv().SetStackVal( Lns_op_not(Ast_isPubToExternal(methodTypeInfo.FP.Get_parentInfo().FP.Get_accessMode()))) ).(bool){
            stream.Write("static ")
        }
    }
    var name string
    var objDecl string
    var retType string
    if _switch14416 := wrapKind; _switch14416 == convCC_FuncWrap__Normal {
        name = moduleCtrl.FP.GetMethodCName(methodTypeInfo)
        
        objDecl = Lns_getVM().String_format(", %s pObj", []LnsAny{convCC_cTypeAnyP})
        
        retType = convCC_getCRetType_1195_(methodTypeInfo.FP.Get_retTypeInfoList())
        
    } else if _switch14416 == convCC_FuncWrap__CallWrap {
        name = moduleCtrl.FP.GetCallMethodCName(methodTypeInfo)
        
        objDecl = Lns_getVM().String_format(", %s pObj", []LnsAny{convCC_cTypeAnyP})
        
        retType = convCC_getCRetType_1195_(methodTypeInfo.FP.Get_retTypeInfoList())
        
    } else if _switch14416 == convCC_FuncWrap__NilWrap {
        name = moduleCtrl.FP.GetNilMethodCName(methodTypeInfo)
        
        objDecl = Lns_getVM().String_format(", %s obj", []LnsAny{convCC_cTypeStem})
        
        if methodTypeInfo.FP.Get_retTypeInfoList().Len() == 0{
            retType = "void"
            
        } else { 
            retType = convCC_cTypeStem
            
        }
    }
    stream.Write(Lns_getVM().String_format("%s %s( %s _pEnv", []LnsAny{retType, name, convCC_cTypeEnvP}))
    if methodTypeInfo.FP.Get_staticFlag(){
        if convCC_isClosure_1011_(methodTypeInfo){
            stream.Write(Lns_getVM().String_format(", %s _pForm", []LnsAny{convCC_cTypeAnyP}))
        }
    } else { 
        stream.Write(objDecl)
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( methodTypeInfo.FP.Get_rawTxt() == "___init") &&
        Lns_GetEnv().SetStackVal( methodTypeInfo.FP.Get_staticFlag()) ).(bool)){
        stream.Write(Lns_getVM().String_format(", %s %s", []LnsAny{convCC_cTypeBlockP, convCC_getBlockName_1198_(Lns_unwrap( methodTypeInfo.FP.GetModule().FP.Get_scope()).(*Ast_Scope))}))
    } else { 
        if argList != nil{
            argList_7068 := argList.(*LnsList)
            for _, _argNode := range( argList_7068.Items ) {
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                stream.Write(", ")
                {
                    _declArgNode := Nodes_DeclArgNodeDownCastF(argNode.FP)
                    if _declArgNode != nil {
                        declArgNode := _declArgNode.(*Nodes_DeclArgNode)
                        convCC_processDeclAlgeSub_1993_(stream, declArgNode)
                    } else {
                        stream.Write(Lns_getVM().String_format("%s _pDDD", []LnsAny{convCC_cTypeStem}))
                    }
                }
            }
        } else {
            for _index, _arg := range( methodTypeInfo.FP.Get_argTypeInfoList().Items ) {
                index := _index + 1
                arg := _arg.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                stream.Write(Lns_getVM().String_format(", %s arg%d", []LnsAny{convCC_getCType_1188_(arg), index}))
            }
        }
    }
    stream.Write(")")
}


// 3188: decl @lune.@base.@convCC.processDeclMethodTable
func convCC_processDeclMethodTable_2020_(stream Util_SourceStream,classTypeInfo *Ast_TypeInfo) {
    var outputField func(name string,retTypeList *LnsList)
    outputField = func(name string,retTypeList *LnsList) {
        var methodType string
        methodType = convCC_getMethodTypeTxt_1987_(retTypeList)
        stream.Writeln(Lns_getVM().String_format("%s * %s;", []LnsAny{methodType, name}))
    }
    var nameSet *Util_OrderedSet
    nameSet = Ast_getAllMethodName(classTypeInfo, Ast_MethodKind__Object)
    for _, _name := range( nameSet.FP.Get_list().Items ) {
        name := _name.(string)
        {
            _symbolInfo := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(classTypeInfo.FP.Get_scope()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.GetSymbolInfoField(name, true, Lns_unwrap( classTypeInfo.FP.Get_scope()).(*Ast_Scope), Ast_ScopeAccess__Normal)})/* 3217:28 */)
            if _symbolInfo != nil {
                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                outputField(name, symbolInfo.FP.Get_typeInfo().FP.Get_retTypeInfoList())
            }
        }
    }
}


// 3227: decl @lune.@base.@convCC.processDeclMemberTable
func convCC_processDeclMemberTable_2034_(normal bool,stream Util_SourceStream,classTypeInfo *Ast_TypeInfo) {
    var outputVal func(scope *Ast_Scope)
    outputVal = func(scope *Ast_Scope) {
        {
            _inherit := scope.FP.Get_inherit()
            if _inherit != nil {
                inherit := _inherit.(*Ast_Scope)
                outputVal(inherit)
            }
        }
        {
            __collection14766 := scope.FP.Get_symbol2SymbolInfoMap()
            __sorted14766 := __collection14766.CreateKeyListStr()
            __sorted14766.Sort( LnsItemKindStr, nil )
            for _, ___key14766 := range( __sorted14766.Items ) {
                symbolInfo := __collection14766.Items[ ___key14766 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if _switch14764 := symbolInfo.FP.Get_kind(); _switch14764 == Ast_SymbolKind__Mbr {
                    if Lns_op_not(symbolInfo.FP.Get_staticFlag()){
                        stream.Writeln(Lns_getVM().String_format("%s %s;", []LnsAny{convCC_getCType_1188_(symbolInfo.FP.Get_typeInfo()), symbolInfo.FP.Get_name()}))
                    }
                }
            }
        }
    }
    stream.Writeln("// member")
    if normal{
        outputVal(Lns_unwrap( classTypeInfo.FP.Get_scope()).(*Ast_Scope))
    } else { 
        stream.Writeln("void * pExt;")
    }
}

// 3254: decl @lune.@base.@convCC.hasGC
func convCC_hasGC_2040_(classTypeInfo *Ast_TypeInfo) bool {
    {
        _scope := classTypeInfo.FP.Get_scope()
        if _scope != nil {
            scope := _scope.(*Ast_Scope)
            if Lns_isCondTrue( scope.FP.GetSymbolInfoField("_gc", true, scope, convCC_scopeAccess)){
                return true
            }
        }
    }
    var workInfo *Ast_TypeInfo
    workInfo = classTypeInfo
    for workInfo.FP.HasBase() {
        workInfo = workInfo.FP.Get_baseTypeInfo()
        
        {
            _scope := classTypeInfo.FP.Get_scope()
            if _scope != nil {
                scope := _scope.(*Ast_Scope)
                if Lns_isCondTrue( scope.FP.GetSymbolInfoField("_gc", true, scope, convCC_scopeAccess)){
                    return true
                }
            }
        }
    }
    return false
}



// 3273: decl @lune.@base.@convCC.processPrototypeMethod
func convCC_processPrototypeMethod_2047_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,declArgNodeList LnsAny,funcTypeInfo *Ast_TypeInfo) {
    var processHeader func(out2HMode LnsInt)
    processHeader = func(out2HMode LnsInt) {
        if out2HMode == convCC_Out2HMode__HeaderPub{
            stream.Write("extern ")
        }
        if Lns_op_not(funcTypeInfo.FP.Get_staticFlag()){
            convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, convCC_FuncWrap__CallWrap, funcTypeInfo, declArgNodeList)
            stream.Writeln(";")
            convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, convCC_FuncWrap__NilWrap, funcTypeInfo, declArgNodeList)
            stream.Writeln(";")
        } else { 
            convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, convCC_FuncWrap__Normal, funcTypeInfo, declArgNodeList)
            stream.Writeln(";")
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_parentInfo().FP.Get_kind() == Ast_TypeInfoKind__Class) &&
        Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_rawTxt() == "__init") ).(bool)){
        convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, convCC_FuncWrap__Normal, funcTypeInfo, declArgNodeList)
        stream.Writeln(";")
    } else { 
        {
            var processwork func(out2HMode LnsInt)
            processwork = func(out2HMode LnsInt) {
                if _switch15273 := out2HMode; _switch15273 == convCC_Out2HMode__SourcePri {
                    if funcTypeInfo.FP.Get_parentInfo().FP.Get_kind() == Ast_TypeInfoKind__Class{
                        processHeader(out2HMode)
                        convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, convCC_FuncWrap__Normal, funcTypeInfo, declArgNodeList)
                        stream.Writeln(";")
                    }
                } else if _switch15273 == convCC_Out2HMode__SourcePub {
                    if funcTypeInfo.FP.Get_parentInfo().FP.Get_kind() == Ast_TypeInfoKind__Class{
                        if Lns_op_not(funcTypeInfo.FP.Get_staticFlag()){
                            convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, convCC_FuncWrap__Normal, funcTypeInfo, declArgNodeList)
                            stream.Writeln(";")
                        }
                    }
                } else if _switch15273 == convCC_Out2HMode__HeaderPub {
                    processHeader(out2HMode)
                }
            }
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Ast_isPubToExternal(funcTypeInfo.FP.Get_parentInfo().FP.Get_accessMode())) &&
                Lns_GetEnv().SetStackVal( Ast_isPubToExternal(funcTypeInfo.FP.Get_accessMode())) ).(bool)){
                stream.SwitchToHeader()
                processwork(convCC_Out2HMode__HeaderPub)
                stream.ReturnToSource()
                processwork(convCC_Out2HMode__SourcePub)
            } else { 
                processwork(convCC_Out2HMode__SourcePri)
            }
        }
        
    }
}

// 3347: decl @lune.@base.@convCC.process2stem
func convCC_process2stem_2060_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,scopeMgr *convCC_ScopeMgr,valKind LnsInt,typeInfo *Ast_TypeInfo,parent *Nodes_Node,callback convCC_process2stemCallback_2057_) {
    if _switch15596 := valKind; _switch15596 == convCC_ValKind__Stem {
        callback()
    } else {
        var expType *Ast_TypeInfo
        expType = typeInfo.FP.Get_srcTypeInfo()
        {
            _enumType := Ast_EnumTypeInfoDownCastF(expType.FP)
            if _enumType != nil {
                enumType := _enumType.(*Ast_EnumTypeInfo)
                expType = enumType.FP.Get_valTypeInfo()
                
            }
        }
        if _switch15593 := expType; _switch15593 == Ast_builtinTypeInt || _switch15593 == Ast_builtinTypeChar {
            stream.Write("LNS_STEM_INT( ")
            callback()
            stream.Write(")")
        } else if _switch15593 == Ast_builtinTypeReal {
            stream.Write("LNS_STEM_REAL( ")
            callback()
            stream.Write(")")
        } else if _switch15593 == Ast_builtinTypeBool {
            stream.Write("LNS_STEM_BOOL( ")
            callback()
            stream.Write(")")
        } else {
            if _switch15591 := expType.FP.Get_kind(); _switch15591 == Ast_TypeInfoKind__DDD {
                stream.Write("_pDDD")
            } else if _switch15591 == Ast_TypeInfoKind__Func {
                if Lns_isCondTrue( expType.FP.Get_scope()){
                    stream.Write("LNS_STEM_ANY(")
                    stream.Write(convCC_getFunc2any_1897_(moduleCtrl, scopeMgr, expType))
                    stream.Write(")")
                } else { 
                    Util_err("illegal func")
                }
            } else {
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( valKind == convCC_ValKind__Var) &&
                    Lns_GetEnv().SetStackVal( convCC_getValKind_1167_(expType) == convCC_ValKind__Stem) ).(bool)){
                    callback()
                } else { 
                    stream.Write("LNS_STEM_ANY( ")
                    callback()
                    stream.Write(")")
                }
            }
        }
    }
}


// 3417: decl @lune.@base.@convCC.processDeclCallMethodWrapper
func convCC_processDeclCallMethodWrapper_2063_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,scopeMgr *convCC_ScopeMgr,parent *Nodes_Node,funcTypeInfo *Ast_TypeInfo,callFlag bool) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_rawTxt() != "__init") &&
        Lns_GetEnv().SetStackVal( Lns_op_not(funcTypeInfo.FP.Get_staticFlag())) ).(bool)){
        convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( callFlag) &&
            Lns_GetEnv().SetStackVal( convCC_FuncWrap__CallWrap) ||
            Lns_GetEnv().SetStackVal( convCC_FuncWrap__NilWrap) ).(LnsInt), funcTypeInfo, nil)
        stream.Writeln("{")
        if Lns_op_not(callFlag){
            var retVal string
            if _switch15705 := funcTypeInfo.FP.Get_retTypeInfoList().Len(); _switch15705 == 0 {
                retVal = ""
                
            } else if _switch15705 == 1 {
                retVal = convCC_cValNil
                
            } else {
                retVal = convCC_cValDDD0
                
            }
            stream.Writeln(Lns_getVM().String_format("if ( obj.type == lns_stem_type_nil ) { return %s; }", []LnsAny{retVal}))
            stream.Writeln(Lns_getVM().String_format("%s pObj = obj.val.pAny;", []LnsAny{convCC_cTypeAnyP}))
        }
        if funcTypeInfo.FP.Get_retTypeInfoList().Len() != 0{
            stream.Write("return ")
        }
        var retList *LnsList
        retList = funcTypeInfo.FP.Get_retTypeInfoList()
        var process func()
        process = func() {
            var classTypeInfo *Ast_TypeInfo
            classTypeInfo = funcTypeInfo.FP.Get_parentInfo()
            var className string
            className = moduleCtrl.FP.GetClassCName(classTypeInfo)
            stream.Write(Lns_getVM().String_format("lns_mtd_%s( pObj )->%s( _pEnv, ", []LnsAny{className, funcTypeInfo.FP.Get_rawTxt()}))
            if classTypeInfo.FP.Get_kind() == Ast_TypeInfoKind__IF{
                stream.Write("lns_getImpObj( pObj ) ")
            } else { 
                stream.Write("pObj ")
            }
            for _index, _ := range( funcTypeInfo.FP.Get_argTypeInfoList().Items ) {
                index := _index + 1
                stream.Write(Lns_getVM().String_format(", arg%d", []LnsAny{index}))
            }
            stream.Write(")")
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(callFlag)) &&
            Lns_GetEnv().SetStackVal( retList.Len() > 0) ).(bool)){
            convCC_process2stem_2060_(stream, moduleCtrl, scopeMgr, convCC_getRetKind_1178_(retList), retList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), parent, convCC_process2stemCallback_2057_(process))
        } else { 
            process()
        }
        stream.Writeln(";")
        stream.Writeln("}")
    }
}

// 3485: decl @lune.@base.@convCC.getAccessMember
func convCC_getAccessMember_2069_(className string,obj string,member string) string {
    return Lns_getVM().String_format("lns_obj_%s( %s )->%s", []LnsAny{className, obj, member})
}

// 3489: decl @lune.@base.@convCC.getAccessMethod
func convCC_getAccessMethod_2072_(className string,obj string,method string) string {
    return Lns_getVM().String_format("lns_mtd_%s( %s )->%s", []LnsAny{className, obj, method})
}

// 3497: decl @lune.@base.@convCC.processAdvertise
func convCC_processAdvertise_2075_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,scopeMgr *convCC_ScopeMgr,processMode LnsInt,node *Nodes_DeclClassNode) {
    var declMethodNameSet *LnsSet
    declMethodNameSet = NewLnsSet([]LnsAny{})
    for _, _field := range( node.FP.Get_fieldList().Items ) {
        field := _field.(Nodes_NodeDownCast).ToNodes_Node()
        {
            _declMethodNode := Nodes_DeclMethodNodeDownCastF(field.FP)
            if _declMethodNode != nil {
                declMethodNode := _declMethodNode.(*Nodes_DeclMethodNode)
                {
                    _name := declMethodNode.FP.Get_declInfo().FP.Get_name()
                    if _name != nil {
                        name := _name.(*Types_Token)
                        declMethodNameSet.Add(name.Txt)
                    }
                }
            }
        }
    }
    for _, _member := range( node.FP.Get_memberList().Items ) {
        member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if member.FP.Get_getterMode() != Ast_AccessMode__None{
            declMethodNameSet.Add("get_" + member.FP.Get_name().Txt)
        }
        if member.FP.Get_setterMode() != Ast_AccessMode__None{
            declMethodNameSet.Add("set_" + member.FP.Get_name().Txt)
        }
    }
    for _, _advInfo := range( node.FP.Get_advertiseList().Items ) {
        advInfo := _advInfo.(Nodes_AdvertiseInfoDownCast).ToNodes_AdvertiseInfo()
        var member *Nodes_DeclMemberNode
        member = advInfo.FP.Get_member()
        stream.Writeln(Lns_getVM().String_format("// for advertise %s.%s --->", []LnsAny{node.FP.Get_name().Txt, member.FP.Get_name().Txt}))
        for _, _name := range( Ast_getAllMethodName(member.FP.Get_expType(), Ast_MethodKind__Object).FP.Get_list().Items ) {
            name := _name.(string)
            if Lns_op_not(declMethodNameSet.Has(name)){
                var methodSym *Ast_SymbolInfo
                methodSym = Lns_unwrap( node.FP.Get_scope().FP.GetSymbolInfoField(name, true, node.FP.Get_scope(), Ast_ScopeAccess__Normal)).(*Ast_SymbolInfo)
                var methodType *Ast_TypeInfo
                methodType = methodSym.FP.Get_typeInfo()
                if methodType.FP.Get_accessMode() != Ast_AccessMode__Pri{
                    if _switch16390 := processMode; _switch16390 == convCC_ProcessMode__Prototype {
                        convCC_processPrototypeMethod_2047_(stream, moduleCtrl, nil, methodType)
                    } else if _switch16390 == convCC_ProcessMode__DefClass {
                        convCC_processDeclCallMethodWrapper_2063_(stream, moduleCtrl, scopeMgr, &node.Nodes_Node, methodType, true)
                        convCC_processDeclCallMethodWrapper_2063_(stream, moduleCtrl, scopeMgr, &node.Nodes_Node, methodType, false)
                        convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, convCC_FuncWrap__Normal, methodType, nil)
                        stream.Writeln("{")
                        var className string
                        className = moduleCtrl.FP.GetClassCName(node.FP.Get_expType())
                        var memberClassName string
                        memberClassName = moduleCtrl.FP.GetClassCName(member.FP.Get_expType())
                        stream.PushIndent(nil)
                        stream.Writeln(Lns_getVM().String_format("%s pVal = %s;", []LnsAny{convCC_cTypeAnyP, convCC_getAccessMember_2069_(className, "pObj", member.FP.Get_name().Txt)}))
                        if methodType.FP.Get_retTypeInfoList().Len() != 0{
                            stream.Write("return ")
                        }
                        stream.Write(Lns_getVM().String_format("%s( _pEnv, pVal", []LnsAny{convCC_getAccessMethod_2072_(memberClassName, "pVal", name)}))
                        for _index, _ := range( methodType.FP.Get_argTypeInfoList().Items ) {
                            index := _index + 1
                            stream.Write(Lns_getVM().String_format(", arg%d", []LnsAny{index}))
                        }
                        stream.Writeln(");")
                        stream.PopIndent()
                        stream.Writeln("}")
                    }
                }
            }
        }
        stream.Writeln(Lns_getVM().String_format("// <-- for advertise %s.%s", []LnsAny{node.FP.Get_name().Txt, member.FP.Get_name().Txt}))
    }
}

// 3592: decl @lune.@base.@convCC.processDeclClassPrototype
func convCC_processDeclClassPrototype_2088_(normal bool,stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,node *Nodes_DeclClassNode) {
    var className string
    className = moduleCtrl.FP.GetClassCName(node.FP.Get_expType())
    stream.Writeln(Lns_getVM().String_format("static void mtd_%s__del( lns_env_t * _pEnv, %s pObj );", []LnsAny{className, convCC_cTypeAnyP}))
    if Lns_op_not(normal){
        stream.Writeln(Lns_getVM().String_format("static void mtd_%s__delExt( lns_env_t * _pEnv, %s pObj );", []LnsAny{className, convCC_cTypeAnyP}))
    }
    if convCC_hasGC_2040_(node.FP.Get_expType()){
        stream.Writeln(Lns_getVM().String_format("static void mtd_%s__gc( lns_env_t * _pEnv, %s pObj );", []LnsAny{className, convCC_cTypeAnyP}))
    }
    for _, _member := range( node.FP.Get_memberList().Items ) {
        member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var memberName string
        memberName = member.FP.Get_name().Txt
        if member.FP.Get_getterMode() != Ast_AccessMode__None{
            var getterType *Ast_TypeInfo
            getterType = Lns_unwrap( node.FP.Get_scope().FP.GetTypeInfoField(Lns_getVM().String_format("get_%s", []LnsAny{memberName}), true, node.FP.Get_scope(), convCC_scopeAccess)).(*Ast_TypeInfo)
            convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, convCC_FuncWrap__Normal, getterType, nil)
            stream.Writeln(";")
        }
        if member.FP.Get_setterMode() != Ast_AccessMode__None{
            var setterType *Ast_TypeInfo
            setterType = Lns_unwrap( node.FP.Get_scope().FP.GetTypeInfoField(Lns_getVM().String_format("set_%s", []LnsAny{memberName}), true, node.FP.Get_scope(), convCC_scopeAccess)).(*Ast_TypeInfo)
            convCC_processMethodDeclTxt_2015_(stream, moduleCtrl, convCC_FuncWrap__Normal, setterType, nil)
            stream.Writeln(";")
        }
    }
}

// 3630: decl @lune.@base.@convCC.processDefaultCtor
func convCC_processDefaultCtor_2091_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,scopeMgr *convCC_ScopeMgr,node *Nodes_DeclClassNode) {
    var className string
    className = moduleCtrl.FP.GetClassCName(node.FP.Get_expType())
    if Lns_op_not(node.FP.HasUserInit()){
        stream.Write(Lns_getVM().String_format("static void %s( lns_env_t * _pEnv, %s pAny", []LnsAny{moduleCtrl.FP.GetCtorName(node.FP.Get_expType()), convCC_cTypeAnyP}))
        var ctorType *Ast_TypeInfo
        ctorType = Lns_unwrap( node.FP.Get_scope().FP.GetTypeInfoField("__init", true, node.FP.Get_scope(), convCC_scopeAccess)).(*Ast_TypeInfo)
        for _index, _argType := range( ctorType.FP.Get_argTypeInfoList().Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            stream.Write(Lns_getVM().String_format(", %s _arg%d", []LnsAny{convCC_getCType_1188_(argType), index}))
        }
        stream.Writeln(") {")
        stream.PushIndent(nil)
        var memberNum LnsInt
        memberNum = 0
        for _, _member := range( node.FP.Get_memberList().Items ) {
            member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if Lns_op_not(member.FP.Get_staticFlag()){
                memberNum = memberNum + 1
                
            }
        }
        var superArgNum LnsInt
        {
            _baseScope := node.FP.Get_scope().FP.Get_inherit()
            if _baseScope != nil {
                baseScope := _baseScope.(*Ast_Scope)
                var superInitType *Ast_TypeInfo
                superInitType = Lns_unwrap( baseScope.FP.GetTypeInfoField("__init", true, baseScope, convCC_scopeAccess)).(*Ast_TypeInfo)
                stream.Write(Lns_getVM().String_format("%s( _pEnv, pAny", []LnsAny{moduleCtrl.FP.GetCtorName(node.FP.Get_expType().FP.Get_baseTypeInfo())}))
                if ctorType.FP.Get_argTypeInfoList().Len() >= superInitType.FP.Get_argTypeInfoList().Len() + memberNum{
                    superArgNum = superInitType.FP.Get_argTypeInfoList().Len()
                    
                    for _index, _ := range( superInitType.FP.Get_argTypeInfoList().Items ) {
                        index := _index + 1
                        stream.Write(Lns_getVM().String_format(", _arg%d", []LnsAny{index}))
                    }
                } else { 
                    superArgNum = 0
                    
                    for range( superInitType.FP.Get_argTypeInfoList().Items ) {
                        stream.Write(Lns_getVM().String_format(", %s", []LnsAny{convCC_cValNil}))
                    }
                }
                stream.Writeln(");")
            } else {
                superArgNum = 0
                
            }
        }
        stream.Writeln(Lns_getVM().String_format("%s * pObj = lns_obj_%s( pAny );", []LnsAny{className, className}))
        var argIndex LnsInt
        argIndex = superArgNum
        for _, _member := range( node.FP.Get_memberList().Items ) {
            member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if Lns_op_not(member.FP.Get_staticFlag()){
                argIndex = argIndex + 1
                
                var valKind LnsInt
                valKind = scopeMgr.FP.GetSymbolValKind(member.FP.Get_symbolInfo().FP)
                if _switch17055 := valKind; _switch17055 == convCC_ValKind__Stem {
                    stream.Writeln(Lns_getVM().String_format("lns_setQ( pObj->%s, _arg%d );", []LnsAny{member.FP.Get_name().Txt, argIndex}))
                } else if _switch17055 == convCC_ValKind__Any {
                    stream.Writeln(Lns_getVM().String_format("lns_setQ_any( &pObj->%s, _arg%d );", []LnsAny{member.FP.Get_name().Txt, argIndex}))
                } else if _switch17055 == convCC_ValKind__Prim {
                    stream.Writeln(Lns_getVM().String_format("pObj->%s = _arg%d;", []LnsAny{member.FP.Get_name().Txt, argIndex}))
                } else {
                    Util_err(Lns_getVM().String_format("no support -- %s:%s:%d", []LnsAny{member.FP.Get_name().Txt, convCC_ValKind_getTxt( valKind), 3713}))
                }
            }
        }
        stream.PopIndent()
        stream.Writeln("}")
    }
}

// 3723: decl @lune.@base.@convCC.processIFObjDecl
func convCC_processIFObjDecl_2094_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,classType *Ast_TypeInfo) {
    if classType.FP.HasBase(){
        convCC_processIFObjDecl_2094_(stream, moduleCtrl, classType.FP.Get_baseTypeInfo())
    }
    for _, _ifType := range( classType.FP.Get_interfaceList().Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        stream.Writeln(Lns_getVM().String_format("%s %s;", []LnsAny{convCC_cTypeAny, moduleCtrl.FP.GetClassCName(ifType)}))
    }
}

// 3735: decl @lune.@base.@convCC.processIFObjInit
func convCC_processIFObjInit_2097_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,classType *Ast_TypeInfo,impClassType *Ast_TypeInfo) {
    if classType.FP.HasBase(){
        convCC_processIFObjInit_2097_(stream, moduleCtrl, classType.FP.Get_baseTypeInfo(), impClassType)
    }
    var className string
    className = moduleCtrl.FP.GetClassCName(impClassType)
    for _, _ifType := range( classType.FP.Get_interfaceList().Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var ifName string
        ifName = moduleCtrl.FP.GetClassCName(ifType)
        stream.Writeln(Lns_getVM().String_format("lns_init_if( &pObj->imp.%s, _pEnv, pAny, &lns_if_%s_imp_%s, %s );", []LnsAny{ifName, className, ifName, ifName}))
    }
}














// 4488: decl @lune.@base.@convCC.processInitMethodTable
func convCC_processInitMethodTable_2167_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,classTypeInfo *Ast_TypeInfo) {
    var outputField func(name string,retTypeList *LnsList)
    outputField = func(name string,retTypeList *LnsList) {
        var methodType string
        methodType = convCC_getMethodTypeTxt_1987_(retTypeList)
        stream.Writeln(Lns_getVM().String_format("(%s *)%s,", []LnsAny{methodType, name}))
    }
    var scope *Ast_Scope
    scope = Lns_unwrap( classTypeInfo.FP.Get_scope()).(*Ast_Scope)
    var nameSet *Util_OrderedSet
    nameSet = Ast_getAllMethodName(classTypeInfo, Ast_MethodKind__Object)
    for _, _name := range( nameSet.FP.Get_list().Items ) {
        name := _name.(string)
        var symbolInfo *Ast_SymbolInfo
        symbolInfo = Lns_unwrap( scope.FP.GetSymbolInfoField(name, true, scope, Ast_ScopeAccess__Normal)).(*Ast_SymbolInfo)
        if Lns_op_not(symbolInfo.FP.Get_typeInfo().FP.Get_abstractFlag()){
            outputField(moduleCtrl.FP.GetMethodCName(symbolInfo.FP.Get_typeInfo()), symbolInfo.FP.Get_typeInfo().FP.Get_retTypeInfoList())
        } else { 
            stream.Writeln("NULL,")
        }
    }
}



// 4528: decl @lune.@base.@convCC.processInitIFMethodTable
func convCC_processInitIFMethodTable_2181_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,ifType *Ast_TypeInfo,classTypeInfo *Ast_TypeInfo) {
    var outputField func(name string,retTypeList *LnsList)
    outputField = func(name string,retTypeList *LnsList) {
        var methodType string
        methodType = convCC_getMethodTypeTxt_1987_(retTypeList)
        stream.Writeln(Lns_getVM().String_format("(%s *)%s,", []LnsAny{methodType, name}))
    }
    var outputVal func(scope *Ast_Scope,impScope *Ast_Scope)
    outputVal = func(scope *Ast_Scope,impScope *Ast_Scope) {
        {
            _inherit := scope.FP.Get_inherit()
            if _inherit != nil {
                inherit := _inherit.(*Ast_Scope)
                outputVal(inherit, impScope)
            }
        }
        {
            __collection21373 := scope.FP.Get_symbol2SymbolInfoMap()
            __sorted21373 := __collection21373.CreateKeyListStr()
            __sorted21373.Sort( LnsItemKindStr, nil )
            for _, ___key21373 := range( __sorted21373.Items ) {
                symbolInfo := __collection21373.Items[ ___key21373 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if _switch21371 := symbolInfo.FP.Get_kind(); _switch21371 == Ast_SymbolKind__Mtd {
                    if symbolInfo.FP.Get_name() != "__init"{
                        var impMethodSym *Ast_SymbolInfo
                        impMethodSym = Lns_unwrap( impScope.FP.GetSymbolInfoField(symbolInfo.FP.Get_name(), true, impScope, convCC_scopeAccess)).(*Ast_SymbolInfo)
                        var impMethodType *Ast_TypeInfo
                        impMethodType = impMethodSym.FP.Get_typeInfo()
                        outputField(moduleCtrl.FP.GetMethodCName(impMethodType), impMethodType.FP.Get_retTypeInfoList())
                    }
                }
            }
        }
    }
    outputVal(Lns_unwrap( ifType.FP.Get_scope()).(*Ast_Scope), Lns_unwrap( classTypeInfo.FP.Get_scope()).(*Ast_Scope))
}

// 4560: decl @lune.@base.@convCC.processIFMethodDataInit
func convCC_processIFMethodDataInit_2194_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,classType *Ast_TypeInfo,orgClassType *Ast_TypeInfo) {
    if classType.FP.HasBase(){
        convCC_processIFMethodDataInit_2194_(stream, moduleCtrl, classType.FP.Get_baseTypeInfo(), orgClassType)
    }
    if Lns_op_not(orgClassType.FP.Get_abstractFlag()){
        var className string
        className = moduleCtrl.FP.GetClassCName(orgClassType)
        for _, _ifType := range( classType.FP.Get_interfaceList().Items ) {
            ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            var ifName string
            ifName = moduleCtrl.FP.GetClassCName(ifType)
            stream.Writeln(Lns_getVM().String_format("static lns_mtd_%s_t lns_if_%s_imp_%s = {", []LnsAny{ifName, className, ifName}))
            stream.PushIndent(nil)
            convCC_processInitIFMethodTable_2181_(stream, moduleCtrl, ifType, orgClassType)
            stream.PopIndent()
            stream.Writeln("};")
        }
    }
}

// 4584: decl @lune.@base.@convCC.processClassMeta
func convCC_processClassMeta_2197_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,classTypeInfo *Ast_TypeInfo) {
    var className string
    className = moduleCtrl.FP.GetClassCName(classTypeInfo)
    stream.Write(Lns_getVM().String_format("lns_type_meta_t %s = { \"%s\", &%s, {", []LnsAny{moduleCtrl.FP.GetClassMetaName(classTypeInfo), className, moduleCtrl.FP.GetClassMetaName(classTypeInfo.FP.Get_baseTypeInfo())}))
    for _, _ifType := range( classTypeInfo.FP.Get_interfaceList().Items ) {
        ifType := _ifType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        stream.Write(Lns_getVM().String_format("&%s, ", []LnsAny{moduleCtrl.FP.GetClassMetaName(ifType)}))
    }
    stream.Writeln("NULL } };")
}



// 4598: decl @lune.@base.@convCC.processClassDataInit
func convCC_processClassDataInit_2204_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,scopeMgr *convCC_ScopeMgr,classTypeInfo *Ast_TypeInfo,fieldList *LnsList) {
    var className string
    className = moduleCtrl.FP.GetClassCName(classTypeInfo)
    if Lns_op_not(Ast_isPubToExternal(classTypeInfo.FP.Get_accessMode())){
        stream.Write("static ")
    }
    convCC_processClassMeta_2197_(stream, moduleCtrl, classTypeInfo)
    if Lns_op_not(classTypeInfo.FP.Get_abstractFlag()){
        if Lns_op_not(Ast_isPubToExternal(classTypeInfo.FP.Get_accessMode())){
            stream.Write("static ")
        }
        stream.Writeln(Lns_getVM().String_format("lns_mtd_%s_t lns_mtd_%s = {", []LnsAny{className, className}))
        stream.PushIndent(nil)
        stream.Writeln(Lns_getVM().String_format("mtd_%s__del,", []LnsAny{className}))
        if convCC_hasGC_2040_(classTypeInfo){
            stream.Writeln(Lns_getVM().String_format("mtd_%s__gc,", []LnsAny{className}))
        } else { 
            stream.Writeln("NULL,")
        }
        convCC_processInitMethodTable_2167_(stream, moduleCtrl, classTypeInfo)
        stream.PopIndent()
        stream.Writeln("};")
    }
    var process func(out2HMode LnsInt,symbolInfo *Ast_SymbolInfo)
    process = func(out2HMode LnsInt,symbolInfo *Ast_SymbolInfo) {
        if _switch21844 := out2HMode; _switch21844 == convCC_Out2HMode__HeaderPub || _switch21844 == convCC_Out2HMode__SourcePub || _switch21844 == convCC_Out2HMode__SourcePri {
            stream.Writeln(Lns_getVM().String_format("%s%s %s;", []LnsAny{convCC_getOut2HeaderPrefix_1518_(out2HMode), Lns_car(scopeMgr.FP.GetCTypeForSym(symbolInfo.FP)).(string), moduleCtrl.FP.GetClassMemberName(symbolInfo.FP)}))
        }
    }
    for _, _symbolInfo := range( (Lns_unwrap( classTypeInfo.FP.Get_scope()).(*Ast_Scope)).FP.Get_symbol2SymbolInfoMap().Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if convCC_isClassMember_1216_(symbolInfo.FP){
            {
                var processwork func(out2HMode LnsInt)
                processwork = func(out2HMode LnsInt) {
                    process(out2HMode, symbolInfo)
                }
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Ast_isPubToExternal(classTypeInfo.FP.Get_accessMode())) &&
                    Lns_GetEnv().SetStackVal( Ast_isPubToExternal(symbolInfo.FP.Get_accessMode())) ).(bool)){
                    stream.SwitchToHeader()
                    processwork(convCC_Out2HMode__HeaderPub)
                    stream.ReturnToSource()
                    processwork(convCC_Out2HMode__SourcePub)
                } else { 
                    processwork(convCC_Out2HMode__SourcePri)
                }
            }
            
        }
    }
}




// 5621: decl @lune.@base.@convCC.processToIF
func convCC_processToIF_2310_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,expType *Ast_TypeInfo,process convCC_processExp_1881_) {
    if expType.FP.Get_kind() == Ast_TypeInfoKind__IF{
        stream.Write("lns_toIF( _pEnv, ")
        process()
        stream.Write(Lns_getVM().String_format(", &lns_type_meta_%s )", []LnsAny{moduleCtrl.FP.GetClassCName(expType)}))
        stream.Write(convCC_accessAny)
    } else { 
        process()
    }
}


// 5637: decl @lune.@base.@convCC.processGetMRet
func convCC_processGetMRet_2313_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,typeInfo *Ast_TypeInfo,index LnsInt) {
    var process func()
    process = func() {
        stream.Write(Lns_getVM().String_format("lns_getMRet( _pEnv, %d )", []LnsAny{index}))
        stream.Write(convCC_getAccessValFromStem_1854_(typeInfo))
    }
    if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__IF{
        convCC_processToIF_2310_(stream, moduleCtrl, typeInfo, convCC_processExp_1881_(process))
    } else { 
        process()
    }
}




// 5770: decl @lune.@base.@convCC.processAlterAccessVal
func convCC_processAlterAccessVal_2356_(stream Util_SourceStream,srcTypeList *LnsList,dstTypeList *LnsList) {
    if dstTypeList.Len() == 1{
        if srcTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Alternate{
            stream.Write(convCC_getAccessValFromStem_1854_(dstTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))
        }
    }
}

// 5791: decl @lune.@base.@convCC.processAlterToActualType
func convCC_processAlterToActualType_2359_(stream Util_SourceStream,moduleCtrl *convCC_ModuleCtrl,fromType *Ast_TypeInfo,toType *Ast_TypeInfo,process convCC_processExp_1881_) {
    if fromType.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
        if toType.FP.Get_kind() == Ast_TypeInfoKind__IF{
            convCC_processToIF_2310_(stream, moduleCtrl, toType, process)
        } else { 
            process()
        }
    } else { 
        process()
    }
}














// 7406: decl @lune.@base.@convCC.getCollectionKind
func convCC_getCollectionKind_2616_(typeInfo *Ast_TypeInfo) LnsAny {
    if _switch33190 := typeInfo.FP.Get_kind(); _switch33190 == Ast_TypeInfoKind__List {
        return convCC_CollectionKind__List_Obj
    } else if _switch33190 == Ast_TypeInfoKind__Array {
        return convCC_CollectionKind__Array_Obj
    } else if _switch33190 == Ast_TypeInfoKind__Set {
        return convCC_CollectionKind__Set_Obj
    } else if _switch33190 == Ast_TypeInfoKind__Map {
        return convCC_CollectionKind__Map_Obj
    } else if _switch33190 == Ast_TypeInfoKind__Ext {
        var extType *Ast_ExtTypeInfo
        extType = Lns_unwrap( Ast_ExtTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo().FP)).(*Ast_ExtTypeInfo)
        var extedType *Ast_TypeInfo
        extedType = extType.FP.Get_extedType()
        if _switch33188 := extedType.FP.Get_kind(); _switch33188 == Ast_TypeInfoKind__List {
            return &convCC_CollectionKind__ExtList{extedType}
        } else if _switch33188 == Ast_TypeInfoKind__Array {
            return &convCC_CollectionKind__ExtArray{extedType}
        } else if _switch33188 == Ast_TypeInfoKind__Set {
            return &convCC_CollectionKind__ExtSet{extedType}
        } else if _switch33188 == Ast_TypeInfoKind__Map {
            return &convCC_CollectionKind__ExtMap{extedType}
        }
    }
    Util_err(Lns_getVM().String_format("unknown collection type -- %s", []LnsAny{typeInfo.FP.GetTxt(nil, nil, nil)}))
// insert a dummy
    return nil
}



// 7766: decl @lune.@base.@convCC.needMRetWrap
func convCC_needMRetWrap_2651_(funcArgTypeList *LnsList,argNodeList *Nodes_ExpListNode) bool {
    var mRetExp *Nodes_MRetExp
    
    {
        _mRetExp := argNodeList.FP.Get_mRetExp()
        if _mRetExp == nil{
            return false
        } else {
            mRetExp = _mRetExp.(*Nodes_MRetExp)
        }
    }
    var argTypeList *LnsList
    argTypeList = argNodeList.FP.Get_expTypeList()
    for _index, _funcArgType := range( funcArgTypeList.Items ) {
        index := _index + 1
        funcArgType := _funcArgType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var argType *Ast_TypeInfo
        argType = argTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if funcArgType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
            return argType.FP.Get_kind() != Ast_TypeInfoKind__DDD
        }
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( argType.FP.Get_kind() == Ast_TypeInfoKind__DDD) ||
            Lns_GetEnv().SetStackVal( mRetExp.FP.Get_index() == index) ).(bool){
            return true
        }
        if mRetExp.FP.Get_index() == index{
            return true
        }
    }
    return true
}






// 8112: decl @lune.@base.@convCC.getMRetFuncName
func convCC_getMRetFuncName_2692_(node *Nodes_Node) string {
    return Lns_getVM().String_format("l_call_mret_%d", []LnsAny{node.FP.Get_id()})
}

// 8583: decl @lune.@base.@convCC.getFormNilWrapper
func convCC_getFormNilWrapper_2713_(node *Nodes_Node) string {
    return Lns_getVM().String_format("l_nil_form_%d", []LnsAny{node.FP.Get_id()})
}

// 8739: decl @lune.@base.@convCC.convFilter.processExpCall.process.fieldCall.processEnumAlge
func convFilter_processExpCall_process_fieldCall__processEnumAlge_2734_() {
}







// 9526: decl @lune.@base.@convCC.convFilter.processAndOr.isAndOr
func convFilter_processAndOr__isAndOr_2783_(exp *Nodes_Node) bool {
    {
        _parentNode := Nodes_ExpOp2NodeDownCastF(exp.FP)
        if _parentNode != nil {
            parentNode := _parentNode.(*Nodes_ExpOp2Node)
            if _switch40777 := parentNode.FP.Get_op().Txt; _switch40777 == "and" || _switch40777 == "or" {
                return true
            }
        }
    }
    return false
}







// 10362: decl @lune.@base.@convCC.getLiteralListFuncName
func convCC_getLiteralListFuncName_2851_(node *Nodes_LiteralListNode) string {
    return Lns_getVM().String_format("lns_list_%X", []LnsAny{node.FP.Get_id()})
}

// 10466: decl @lune.@base.@convCC.getLiteralSetFuncName
func convCC_getLiteralSetFuncName_2879_(node *Nodes_LiteralSetNode) string {
    return Lns_getVM().String_format("lns_set_%X", []LnsAny{node.FP.Get_id()})
}

// 10495: decl @lune.@base.@convCC.getLiteralMapFuncName
func convCC_getLiteralMapFuncName_2889_(node *Nodes_LiteralMapNode) string {
    return Lns_getVM().String_format("lns_map_%X", []LnsAny{node.FP.Get_id()})
}

// 10574: decl @lune.@base.@convCC.getLiteralArrayFuncName
func convCC_getLiteralArrayFuncName_2913_(node *Nodes_LiteralArrayNode) string {
    return Lns_getVM().String_format("lns_array_%X", []LnsAny{node.FP.Get_id()})
}

// 10767: decl @lune.@base.@convCC.createFilter
func ConvCC_createFilter(enableTest bool,outputBuiltin bool,streamName string,stream Lns_oStream,headerStream Lns_oStream,ast *TransUnit_ASTInfo) *Nodes_Filter {
    return &NewconvCC_convFilter(enableTest, outputBuiltin, streamName, stream, headerStream, ast).Nodes_Filter
}

// 10775: decl @lune.@base.@convCC.outputBootcode
func ConvCC_outputBootcode(stream Lns_oStream,launchModuleName string) {
    var srcStream *Util_SimpleSourceOStream
    srcStream = NewUtil_SimpleSourceOStream(stream, nil, convCC_stepIndent)
    var launchModulePath string
    launchModulePath = convCC_convExp46440(Lns_2DDD(Lns_getVM().String_gsub(launchModuleName,"%.", "/")))
    var moduleName string
    moduleName = convCC_convExp46454(Lns_2DDD(Lns_getVM().String_gsub(launchModuleName,"%.", "_")))
    srcStream.FP.Writeln(Lns_getVM().String_format("#include <lunescript.h>\n#include <%s.h>\n    \nvoid lns_run_module( %s _pEnv ) {\n   %s pInfo = lns_init_%s( _pEnv );\n   lns_test( _pEnv, pInfo );      \n}\n", []LnsAny{launchModulePath, convCC_cTypeEnvP, convCC_cTypeModP, moduleName}))
}

// declaration Class -- PubVarInfo
type convCC_PubVarInfoMtd interface {
}
type convCC_PubVarInfo struct {
    StaticFlag bool
    AccessMode LnsInt
    Mutable bool
    TypeInfo *Ast_TypeInfo
    FP convCC_PubVarInfoMtd
}
func convCC_PubVarInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_PubVarInfo).FP
}
type convCC_PubVarInfoDownCast interface {
    ToconvCC_PubVarInfo() *convCC_PubVarInfo
}
func convCC_PubVarInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_PubVarInfoDownCast)
    if ok { return work.ToconvCC_PubVarInfo() }
    return nil
}
func (obj *convCC_PubVarInfo) ToconvCC_PubVarInfo() *convCC_PubVarInfo {
    return obj
}
func NewconvCC_PubVarInfo(arg1 bool, arg2 LnsInt, arg3 bool, arg4 *Ast_TypeInfo) *convCC_PubVarInfo {
    obj := &convCC_PubVarInfo{}
    obj.FP = obj
    obj.InitconvCC_PubVarInfo(arg1, arg2, arg3, arg4)
    return obj
}
func (self *convCC_PubVarInfo) InitconvCC_PubVarInfo(arg1 bool, arg2 LnsInt, arg3 bool, arg4 *Ast_TypeInfo) {
    self.StaticFlag = arg1
    self.AccessMode = arg2
    self.Mutable = arg3
    self.TypeInfo = arg4
}

// declaration Class -- PubFuncInfo
type convCC_PubFuncInfoMtd interface {
}
type convCC_PubFuncInfo struct {
    AccessMode LnsInt
    TypeInfo *Ast_TypeInfo
    FP convCC_PubFuncInfoMtd
}
func convCC_PubFuncInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_PubFuncInfo).FP
}
type convCC_PubFuncInfoDownCast interface {
    ToconvCC_PubFuncInfo() *convCC_PubFuncInfo
}
func convCC_PubFuncInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_PubFuncInfoDownCast)
    if ok { return work.ToconvCC_PubFuncInfo() }
    return nil
}
func (obj *convCC_PubFuncInfo) ToconvCC_PubFuncInfo() *convCC_PubFuncInfo {
    return obj
}
func NewconvCC_PubFuncInfo(arg1 LnsInt, arg2 *Ast_TypeInfo) *convCC_PubFuncInfo {
    obj := &convCC_PubFuncInfo{}
    obj.FP = obj
    obj.InitconvCC_PubFuncInfo(arg1, arg2)
    return obj
}
func (self *convCC_PubFuncInfo) InitconvCC_PubFuncInfo(arg1 LnsInt, arg2 *Ast_TypeInfo) {
    self.AccessMode = arg1
    self.TypeInfo = arg2
}

// declaration Class -- ModuleInfo
type convCC_ModuleInfoMtd interface {
    Get_assignName() string
    Get_isLazyLoad() bool
    Get_modulePath() string
}
type convCC_ModuleInfo struct {
    assignName string
    modulePath string
    isLazyLoad bool
    FP convCC_ModuleInfoMtd
}
func convCC_ModuleInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_ModuleInfo).FP
}
type convCC_ModuleInfoDownCast interface {
    ToconvCC_ModuleInfo() *convCC_ModuleInfo
}
func convCC_ModuleInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_ModuleInfoDownCast)
    if ok { return work.ToconvCC_ModuleInfo() }
    return nil
}
func (obj *convCC_ModuleInfo) ToconvCC_ModuleInfo() *convCC_ModuleInfo {
    return obj
}
func NewconvCC_ModuleInfo(arg1 string, arg2 string, arg3 bool) *convCC_ModuleInfo {
    obj := &convCC_ModuleInfo{}
    obj.FP = obj
    obj.InitconvCC_ModuleInfo(arg1, arg2, arg3)
    return obj
}
func (self *convCC_ModuleInfo) InitconvCC_ModuleInfo(arg1 string, arg2 string, arg3 bool) {
    self.assignName = arg1
    self.modulePath = arg2
    self.isLazyLoad = arg3
}
func (self *convCC_ModuleInfo) Get_assignName() string{ return self.assignName }
func (self *convCC_ModuleInfo) Get_modulePath() string{ return self.modulePath }
func (self *convCC_ModuleInfo) Get_isLazyLoad() bool{ return self.isLazyLoad }

// declaration Class -- Opt
type ConvCC_OptMtd interface {
}
type ConvCC_Opt struct {
    Node *Nodes_Node
    FP ConvCC_OptMtd
}
func ConvCC_Opt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvCC_Opt).FP
}
type ConvCC_OptDownCast interface {
    ToConvCC_Opt() *ConvCC_Opt
}
func ConvCC_OptDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvCC_OptDownCast)
    if ok { return work.ToConvCC_Opt() }
    return nil
}
func (obj *ConvCC_Opt) ToConvCC_Opt() *ConvCC_Opt {
    return obj
}
func NewConvCC_Opt(arg1 *Nodes_Node) *ConvCC_Opt {
    obj := &ConvCC_Opt{}
    obj.FP = obj
    obj.InitConvCC_Opt(arg1)
    return obj
}
func (self *ConvCC_Opt) InitConvCC_Opt(arg1 *Nodes_Node) {
    self.Node = arg1
}

// declaration Class -- DepthInfo
type convCC_DepthInfoMtd interface {
    Get_blockDepth() LnsInt
    PopDepth()
    PushDepth()
}
type convCC_DepthInfo struct {
    blockDepth LnsInt
    FP convCC_DepthInfoMtd
}
func convCC_DepthInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_DepthInfo).FP
}
type convCC_DepthInfoDownCast interface {
    ToconvCC_DepthInfo() *convCC_DepthInfo
}
func convCC_DepthInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_DepthInfoDownCast)
    if ok { return work.ToconvCC_DepthInfo() }
    return nil
}
func (obj *convCC_DepthInfo) ToconvCC_DepthInfo() *convCC_DepthInfo {
    return obj
}
func NewconvCC_DepthInfo() *convCC_DepthInfo {
    obj := &convCC_DepthInfo{}
    obj.FP = obj
    obj.InitconvCC_DepthInfo()
    return obj
}
func (self *convCC_DepthInfo) Get_blockDepth() LnsInt{ return self.blockDepth }
// 130: DeclConstr
func (self *convCC_DepthInfo) InitconvCC_DepthInfo() {
    self.blockDepth = 1
    
}

// 133: decl @lune.@base.@convCC.DepthInfo.pushDepth
func (self *convCC_DepthInfo) PushDepth() {
    self.blockDepth = self.blockDepth + 1
    
}

// 136: decl @lune.@base.@convCC.DepthInfo.popDepth
func (self *convCC_DepthInfo) PopDepth() {
    self.blockDepth = self.blockDepth - 1
    
}


// declaration Class -- DepthStack
type convCC_DepthStackMtd interface {
    Current() *convCC_DepthInfo
    CurrentR() *convCC_DepthInfo
    DelInfo()
    Get_blockDepth() LnsInt
    NewInfo(arg1 *convCC_DepthInfo)
    PopDepth()
    PushDepth()
}
type convCC_DepthStack struct {
    stack *LnsList
    FP convCC_DepthStackMtd
}
func convCC_DepthStack2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_DepthStack).FP
}
type convCC_DepthStackDownCast interface {
    ToconvCC_DepthStack() *convCC_DepthStack
}
func convCC_DepthStackDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_DepthStackDownCast)
    if ok { return work.ToconvCC_DepthStack() }
    return nil
}
func (obj *convCC_DepthStack) ToconvCC_DepthStack() *convCC_DepthStack {
    return obj
}
func NewconvCC_DepthStack() *convCC_DepthStack {
    obj := &convCC_DepthStack{}
    obj.FP = obj
    obj.InitconvCC_DepthStack()
    return obj
}
// 144: DeclConstr
func (self *convCC_DepthStack) InitconvCC_DepthStack() {
    self.stack = NewLnsList([]LnsAny{})
    
}

// 148: decl @lune.@base.@convCC.DepthStack.newInfo
func (self *convCC_DepthStack) NewInfo(_info *convCC_DepthInfo) {
    info := _info
    self.stack.Insert(info)
}

// 151: decl @lune.@base.@convCC.DepthStack.delInfo
func (self *convCC_DepthStack) DelInfo() {
    self.stack.Remove(nil)
}

// 154: decl @lune.@base.@convCC.DepthStack.current
func (self *convCC_DepthStack) Current() *convCC_DepthInfo {
    if self.stack.Len() == 0{
        Util_err("stack empty")
    }
    return self.stack.GetAt(self.stack.Len()).(*convCC_DepthInfo)
}

// 160: decl @lune.@base.@convCC.DepthStack.currentR
func (self *convCC_DepthStack) CurrentR() *convCC_DepthInfo {
    if self.stack.Len() == 0{
        Util_err("stack empty")
    }
    return self.stack.GetAt(self.stack.Len()).(*convCC_DepthInfo)
}

// 166: decl @lune.@base.@convCC.DepthStack.pushDepth
func (self *convCC_DepthStack) PushDepth() {
    self.FP.Current().FP.PushDepth()
}

// 169: decl @lune.@base.@convCC.DepthStack.popDepth
func (self *convCC_DepthStack) PopDepth() {
    self.FP.Current().FP.PopDepth()
}

// 172: decl @lune.@base.@convCC.DepthStack.get_blockDepth
func (self *convCC_DepthStack) Get_blockDepth() LnsInt {
    return self.FP.CurrentR().FP.Get_blockDepth()
}


// declaration Class -- RoutineInfo
type convCC_RoutineInfoMtd interface {
    Get_blockDepth() LnsInt
    Get_funcInfo() *Ast_TypeInfo
    PopDepth()
    PushDepth()
}
type convCC_RoutineInfo struct {
    convCC_DepthInfo
    funcInfo *Ast_TypeInfo
    FP convCC_RoutineInfoMtd
}
func convCC_RoutineInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_RoutineInfo).FP
}
type convCC_RoutineInfoDownCast interface {
    ToconvCC_RoutineInfo() *convCC_RoutineInfo
}
func convCC_RoutineInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_RoutineInfoDownCast)
    if ok { return work.ToconvCC_RoutineInfo() }
    return nil
}
func (obj *convCC_RoutineInfo) ToconvCC_RoutineInfo() *convCC_RoutineInfo {
    return obj
}
func NewconvCC_RoutineInfo(arg1 *Ast_TypeInfo) *convCC_RoutineInfo {
    obj := &convCC_RoutineInfo{}
    obj.FP = obj
    obj.convCC_DepthInfo.FP = obj
    obj.InitconvCC_RoutineInfo(arg1)
    return obj
}
func (self *convCC_RoutineInfo) Get_funcInfo() *Ast_TypeInfo{ return self.funcInfo }
// 182: DeclConstr
func (self *convCC_RoutineInfo) InitconvCC_RoutineInfo(funcInfo *Ast_TypeInfo) {
    self.InitconvCC_DepthInfo()
    self.funcInfo = funcInfo
    
}


// declaration Class -- ModuleCtrl
type convCC_ModuleCtrlMtd interface {
    GetAlgeCName(arg1 *Ast_TypeInfo) string
    GetAlgeEnumCName(arg1 *Ast_TypeInfo) string
    GetAlgeInitCName(arg1 *Ast_TypeInfo) string
    GetAlgeValCName(arg1 *Ast_TypeInfo, arg2 string) string
    GetAlgeValStrCName(arg1 *Ast_TypeInfo, arg2 string) string
    GetBuiltinFuncName(arg1 *Ast_SymbolInfo) LnsAny
    GetBuiltinFuncNameFromType(arg1 *Ast_TypeInfo) LnsAny
    GetCallFormName(arg1 *Ast_TypeInfo) string
    GetCallMethodCName(arg1 *Ast_TypeInfo) string
    GetCanonicalName(arg1 *Ast_TypeInfo) string
    GetClassCName(arg1 *Ast_TypeInfo) string
    GetClassMemberName(arg1 Ast_LowSymbol) string
    GetClassMetaName(arg1 *Ast_TypeInfo) string
    GetCtorName(arg1 *Ast_TypeInfo) string
    GetEnumFuncName(arg1 *Ast_EnumTypeInfo, arg2 string) string
    GetEnumTypeName(arg1 *Ast_TypeInfo) string
    GetEnumVal2NameMapName(arg1 *Ast_TypeInfo) string
    GetEnumValCName(arg1 *Ast_TypeInfo, arg2 string) string
    GetFilePath(arg1 *Ast_TypeInfo) string
    GetFormName(arg1 *Ast_TypeInfo) string
    GetFullName(arg1 *Ast_TypeInfo) string
    GetFuncCastWrapName(arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo) string
    GetFuncName(arg1 *Ast_TypeInfo) string
    GetMethodCName(arg1 *Ast_TypeInfo) string
    GetNewAlgeCName(arg1 *Ast_TypeInfo, arg2 string) string
    GetNewName(arg1 *Ast_TypeInfo) string
    GetNilMethodCName(arg1 *Ast_TypeInfo) string
    GetSymbolName(arg1 Ast_LowSymbol) string
}
type convCC_ModuleCtrl struct {
    moduleInfoManager Ast_ModuleInfoManager
    typeNameCtrl *Ast_TypeNameCtrl
    builtinSym2CFuncMap *LnsMap
    builtinType2CFuncMap *LnsMap
    FP convCC_ModuleCtrlMtd
}
func convCC_ModuleCtrl2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_ModuleCtrl).FP
}
type convCC_ModuleCtrlDownCast interface {
    ToconvCC_ModuleCtrl() *convCC_ModuleCtrl
}
func convCC_ModuleCtrlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_ModuleCtrlDownCast)
    if ok { return work.ToconvCC_ModuleCtrl() }
    return nil
}
func (obj *convCC_ModuleCtrl) ToconvCC_ModuleCtrl() *convCC_ModuleCtrl {
    return obj
}
func NewconvCC_ModuleCtrl(arg1 *Ast_TypeNameCtrl, arg2 Ast_ModuleInfoManager) *convCC_ModuleCtrl {
    obj := &convCC_ModuleCtrl{}
    obj.FP = obj
    obj.InitconvCC_ModuleCtrl(arg1, arg2)
    return obj
}
// 383: decl @lune.@base.@convCC.ModuleCtrl.setupBuiltinSym
func convCC_ModuleCtrl_setupBuiltinSym_1246_()(*LnsMap, *LnsMap) {
    var builtinFunc *TransUnit_BuiltinFuncType
    builtinFunc = TransUnit_getBuiltinFunc()
    var symMap *LnsMap
    symMap = NewLnsMap( map[LnsAny]LnsAny{})
    var typeMap *LnsMap
    typeMap = NewLnsMap( map[LnsAny]LnsAny{})
    typeMap.Set(builtinFunc.Lns_type,"lns_f_" + "type")
    
    typeMap.Set(builtinFunc.Lns_error,"lns_f_" + "error")
    
    typeMap.Set(builtinFunc.Lns_print,"lns_f_" + "print")
    
    typeMap.Set(builtinFunc.Lns_tonumber,"lns_f_" + "tonumber")
    
    typeMap.Set(builtinFunc.Lns_tostring,"lns_f_" + "tostring")
    
    typeMap.Set(builtinFunc.Lns_loadfile,"lns_f_" + "loadfile")
    
    typeMap.Set(builtinFunc.Lns_require,"lns_f_" + "require")
    
    typeMap.Set(builtinFunc.Lns_collectgarbage,"lns_f_" + "collectgarbage")
    
    typeMap.Set(builtinFunc.Lns__fcall,"lns_f_" + "_fcall")
    
    typeMap.Set(builtinFunc.Lns__load,"lns_f_" + "_load")
    
    typeMap.Set(builtinFunc.Lns__kind,"lns_f_" + "_kind")
    
    typeMap.Set(builtinFunc.Io_open,"mtd_lns_" + "io_open")
    
    typeMap.Set(builtinFunc.Io_popen,"mtd_lns_" + "io_popen")
    
    typeMap.Set(builtinFunc.Package_searchpath,"mtd_lns_" + "package_searchpath")
    
    typeMap.Set(builtinFunc.Os_clock,"mtd_lns_" + "os_clock")
    
    typeMap.Set(builtinFunc.Os_date,"mtd_lns_" + "os_date")
    
    typeMap.Set(builtinFunc.Os_difftime,"mtd_lns_" + "os_difftime")
    
    typeMap.Set(builtinFunc.Os_exit,"mtd_lns_" + "os_exit")
    
    typeMap.Set(builtinFunc.Os_remove,"mtd_lns_" + "os_remove")
    
    typeMap.Set(builtinFunc.Os_rename,"mtd_lns_" + "os_rename")
    
    typeMap.Set(builtinFunc.Os_time,"mtd_lns_" + "os_time")
    
    typeMap.Set(builtinFunc.String_byte,"mtd_lns_" + "string_byte")
    
    typeMap.Set(builtinFunc.String_dump,"mtd_lns_" + "string_dump")
    
    typeMap.Set(builtinFunc.String_find,"mtd_lns_" + "string_find")
    
    typeMap.Set(builtinFunc.String_format,"mtd_lns_" + "string_format")
    
    typeMap.Set(builtinFunc.String_gmatch,"mtd_lns_" + "string_gmatch")
    
    typeMap.Set(builtinFunc.String_gsub,"mtd_lns_" + "string_gsub")
    
    typeMap.Set(builtinFunc.String_lower,"mtd_lns_" + "string_lower")
    
    typeMap.Set(builtinFunc.String_rep,"mtd_lns_" + "string_rep")
    
    typeMap.Set(builtinFunc.String_reverse,"mtd_lns_" + "string_reverse")
    
    typeMap.Set(builtinFunc.String_sub,"mtd_lns_" + "string_sub")
    
    typeMap.Set(builtinFunc.String_upper,"mtd_lns_" + "string_upper")
    
    typeMap.Set(builtinFunc.Math_random,"mtd_lns_" + "math_random")
    
    typeMap.Set(builtinFunc.Math_randomseed,"mtd_lns_" + "math_randomseed")
    
    typeMap.Set(builtinFunc.Debug_getinfo,"mtd_lns_" + "debug_getinfo")
    
    typeMap.Set(builtinFunc.Debug_getlocal,"mtd_lns_" + "debug_getlocal")
    
    return symMap, typeMap
}

// 441: DeclConstr
func (self *convCC_ModuleCtrl) InitconvCC_ModuleCtrl(typeNameCtrl *Ast_TypeNameCtrl,moduleInfoManager Ast_ModuleInfoManager) {
    self.builtinSym2CFuncMap, self.builtinType2CFuncMap = convCC_ModuleCtrl_setupBuiltinSym_1246_()
    
    self.typeNameCtrl = typeNameCtrl
    
    self.moduleInfoManager = moduleInfoManager
    
}

// 449: decl @lune.@base.@convCC.ModuleCtrl.getBuiltinFuncName
func (self *convCC_ModuleCtrl) GetBuiltinFuncName(symbol *Ast_SymbolInfo) LnsAny {
    return self.builtinSym2CFuncMap.Items[symbol]
}

// 453: decl @lune.@base.@convCC.ModuleCtrl.getBuiltinFuncNameFromType
func (self *convCC_ModuleCtrl) GetBuiltinFuncNameFromType(typeInfo *Ast_TypeInfo) LnsAny {
    return self.builtinType2CFuncMap.Items[typeInfo]
}

// 462: decl @lune.@base.@convCC.ModuleCtrl.getFilePath
func (self *convCC_ModuleCtrl) GetFilePath(moduleTypeInfo *Ast_TypeInfo) string {
    var workName string
    workName = moduleTypeInfo.FP.GetFullName(self.typeNameCtrl, self.moduleInfoManager, false)
    var fullName string
    fullName = Lns_getVM().String_format("%s", []LnsAny{Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(workName,"[&@]", "")).(string),"%.", "/")).(string)})
    return fullName
}

// 470: decl @lune.@base.@convCC.ModuleCtrl.getCanonicalName
func (self *convCC_ModuleCtrl) GetCanonicalName(typeInfo *Ast_TypeInfo) string {
    return typeInfo.FP.GetFullName(self.typeNameCtrl, self.moduleInfoManager, false)
}

// 477: decl @lune.@base.@convCC.ModuleCtrl.getFullName
func (self *convCC_ModuleCtrl) GetFullName(typeInfo *Ast_TypeInfo) string {
    {
        _alterType := Ast_AlternateTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo().FP)
        if _alterType != nil {
            alterType := _alterType.(*Ast_AlternateTypeInfo)
            if alterType.FP.HasBase(){
                typeInfo = alterType.FP.Get_baseTypeInfo()
                
            }
        }
    }
    typeInfo = typeInfo.FP.Get_srcTypeInfo().FP.Get_genSrcTypeInfo()
    
    var workName string
    workName = typeInfo.FP.GetFullName(self.typeNameCtrl, self.moduleInfoManager, false)
    var fullName string
    fullName = Lns_getVM().String_format("%s", []LnsAny{Lns_car(Lns_getVM().String_gsub(Lns_car(Lns_getVM().String_gsub(workName,"[&@]", "")).(string),"%.", "_")).(string)})
    if Ast_isPubToExternal(typeInfo.FP.Get_accessMode()){
        return fullName
    }
    return Lns_getVM().String_format("_%d_%s", []LnsAny{typeInfo.FP.Get_typeId(), fullName})
}

// 495: decl @lune.@base.@convCC.ModuleCtrl.getAlgeCName
func (self *convCC_ModuleCtrl) GetAlgeCName(algeType *Ast_TypeInfo) string {
    return self.FP.GetFullName(algeType)
}

// 498: decl @lune.@base.@convCC.ModuleCtrl.getAlgeEnumCName
func (self *convCC_ModuleCtrl) GetAlgeEnumCName(algeType *Ast_TypeInfo) string {
    return Lns_getVM().String_format("lns_algeType_%s", []LnsAny{self.FP.GetAlgeCName(algeType)})
}

// 501: decl @lune.@base.@convCC.ModuleCtrl.getAlgeValCName
func (self *convCC_ModuleCtrl) GetAlgeValCName(algeType *Ast_TypeInfo,valName string) string {
    return Lns_getVM().String_format("lns__alge_%s_%s", []LnsAny{self.FP.GetFullName(algeType), valName})
}

// 504: decl @lune.@base.@convCC.ModuleCtrl.getAlgeValStrCName
func (self *convCC_ModuleCtrl) GetAlgeValStrCName(algeType *Ast_TypeInfo,valName string) string {
    return Lns_getVM().String_format("lns__alge_%s_%s_t", []LnsAny{self.FP.GetFullName(algeType), valName})
}

// 507: decl @lune.@base.@convCC.ModuleCtrl.getNewAlgeCName
func (self *convCC_ModuleCtrl) GetNewAlgeCName(algeType *Ast_TypeInfo,valName string) string {
    return Lns_getVM().String_format("lns__new_alge_%s_%s", []LnsAny{self.FP.GetFullName(algeType), valName})
}

// 511: decl @lune.@base.@convCC.ModuleCtrl.getAlgeInitCName
func (self *convCC_ModuleCtrl) GetAlgeInitCName(algeType *Ast_TypeInfo) string {
    return Lns_getVM().String_format("lns__init_alge_%s", []LnsAny{self.FP.GetAlgeCName(algeType)})
}

// 516: decl @lune.@base.@convCC.ModuleCtrl.getEnumTypeName
func (self *convCC_ModuleCtrl) GetEnumTypeName(typeInfo *Ast_TypeInfo) string {
    var srcType *Ast_TypeInfo
    srcType = typeInfo.FP.Get_srcTypeInfo()
    var fullName string
    fullName = self.FP.GetFullName(srcType)
    if Ast_isPubToExternal(typeInfo.FP.Get_accessMode()){
        return fullName
    }
    return Lns_getVM().String_format("e_%s", []LnsAny{fullName})
}

// 526: decl @lune.@base.@convCC.ModuleCtrl.getEnumValCName
func (self *convCC_ModuleCtrl) GetEnumValCName(typeInfo *Ast_TypeInfo,valName string) string {
    return Lns_getVM().String_format("%s__%s", []LnsAny{self.FP.GetEnumTypeName(typeInfo), valName})
}

// 530: decl @lune.@base.@convCC.ModuleCtrl.getEnumVal2NameMapName
func (self *convCC_ModuleCtrl) GetEnumVal2NameMapName(enumType *Ast_TypeInfo) string {
    return Lns_getVM().String_format("%s_val2NameMap", []LnsAny{self.FP.GetEnumTypeName(enumType)})
}

// 534: decl @lune.@base.@convCC.ModuleCtrl.getClassCName
func (self *convCC_ModuleCtrl) GetClassCName(classType *Ast_TypeInfo) string {
    return "lns_" + self.FP.GetFullName(classType)
}

// 539: decl @lune.@base.@convCC.ModuleCtrl.getNewName
func (self *convCC_ModuleCtrl) GetNewName(classType *Ast_TypeInfo) string {
    return Lns_getVM().String_format("lns_class_%s_new", []LnsAny{self.FP.GetClassCName(classType)})
}

// 543: decl @lune.@base.@convCC.ModuleCtrl.getCtorName
func (self *convCC_ModuleCtrl) GetCtorName(classType *Ast_TypeInfo) string {
    return Lns_getVM().String_format("mtd_%s___init", []LnsAny{self.FP.GetClassCName(classType)})
}

// 547: decl @lune.@base.@convCC.ModuleCtrl.getClassMetaName
func (self *convCC_ModuleCtrl) GetClassMetaName(classType *Ast_TypeInfo) string {
    if classType.FP.Get_srcTypeInfo() == Ast_headTypeInfo{
        return "lns_type_meta_lns__root"
    }
    return Lns_getVM().String_format("lns_type_meta_%s", []LnsAny{self.FP.GetClassCName(classType)})
}

// 555: decl @lune.@base.@convCC.ModuleCtrl.getMethodCName
func (self *convCC_ModuleCtrl) GetMethodCName(methodTypeInfo *Ast_TypeInfo) string {
    return Lns_getVM().String_format("mtd_%s_%s", []LnsAny{self.FP.GetClassCName(methodTypeInfo.FP.Get_parentInfo()), methodTypeInfo.FP.Get_rawTxt()})
}

// 561: decl @lune.@base.@convCC.ModuleCtrl.getFuncName
func (self *convCC_ModuleCtrl) GetFuncName(typeInfo *Ast_TypeInfo) string {
    if typeInfo.FP.Get_rawTxt() == ""{
        return Lns_getVM().String_format("lns_anonymous_%d", []LnsAny{typeInfo.FP.Get_typeId()})
    }
    if _switch2610 := typeInfo.FP.Get_accessMode(); _switch2610 == Ast_AccessMode__Pub || _switch2610 == Ast_AccessMode__Global {
        {
            _cFuncName := self.FP.GetBuiltinFuncNameFromType(typeInfo)
            if _cFuncName != nil {
                cFuncName := _cFuncName.(string)
                return cFuncName
            }
        }
        if _switch2599 := typeInfo.FP.Get_parentInfo().FP.Get_kind(); _switch2599 == Ast_TypeInfoKind__Class || _switch2599 == Ast_TypeInfoKind__Enum {
            return self.FP.GetMethodCName(typeInfo)
        }
        return self.FP.GetFullName(typeInfo)
    }
    if typeInfo.FP.Get_parentInfo().FP.Get_kind() == Ast_TypeInfoKind__Class{
        return self.FP.GetMethodCName(typeInfo)
    }
    return Lns_getVM().String_format("lns_f_%d_%s", []LnsAny{typeInfo.FP.Get_typeId(), typeInfo.FP.Get_rawTxt()})
}

// 586: decl @lune.@base.@convCC.ModuleCtrl.getNilMethodCName
func (self *convCC_ModuleCtrl) GetNilMethodCName(methodTypeInfo *Ast_TypeInfo) string {
    return Lns_getVM().String_format("l_nil_mtd_%s_%s", []LnsAny{self.FP.GetClassCName(methodTypeInfo.FP.Get_parentInfo()), methodTypeInfo.FP.Get_rawTxt()})
}

// 592: decl @lune.@base.@convCC.ModuleCtrl.getCallMethodCName
func (self *convCC_ModuleCtrl) GetCallMethodCName(methodTypeInfo *Ast_TypeInfo) string {
    if _switch2746 := methodTypeInfo.FP.Get_parentInfo().FP.Get_kind(); _switch2746 == Ast_TypeInfoKind__List {
        return Lns_getVM().String_format("lns_mtd_List_%s", []LnsAny{methodTypeInfo.FP.Get_rawTxt()})
    } else if _switch2746 == Ast_TypeInfoKind__Array {
        return Lns_getVM().String_format("lns_mtd_Array_%s", []LnsAny{methodTypeInfo.FP.Get_rawTxt()})
    } else if _switch2746 == Ast_TypeInfoKind__Set {
        return Lns_getVM().String_format("lns_mtd_Set_%s", []LnsAny{methodTypeInfo.FP.Get_rawTxt()})
    } else if _switch2746 == Ast_TypeInfoKind__Map {
        return Lns_getVM().String_format("lns_mtd_Map_%s", []LnsAny{methodTypeInfo.FP.Get_rawTxt()})
    }
    return Lns_getVM().String_format("l_call_mtd_%s_%s", []LnsAny{self.FP.GetClassCName(methodTypeInfo.FP.Get_parentInfo()), methodTypeInfo.FP.Get_rawTxt()})
}

// 612: decl @lune.@base.@convCC.ModuleCtrl.getClassMemberName
func (self *convCC_ModuleCtrl) GetClassMemberName(symbolInfo Ast_LowSymbol) string {
    var classTypeInfo *Ast_TypeInfo
    classTypeInfo = symbolInfo.Get_scope().FP.GetClassTypeInfo()
    return Lns_getVM().String_format("l_var_%s_%s", []LnsAny{self.FP.GetClassCName(classTypeInfo), symbolInfo.Get_name()})
}

// 621: decl @lune.@base.@convCC.ModuleCtrl.getSymbolName
func (self *convCC_ModuleCtrl) GetSymbolName(symbolInfo Ast_LowSymbol) string {
    if symbolInfo.Get_typeInfo().FP.Get_kind() == Ast_TypeInfoKind__DDD{
        return "_pDDD"
    }
    if symbolInfo.Get_kind() == Ast_SymbolKind__Mbr{
        if convCC_isClassMember_1216_(symbolInfo){
            return self.FP.GetClassMemberName(symbolInfo)
        }
        return symbolInfo.Get_name()
    }
    if Ast_isPubToExternal(symbolInfo.Get_accessMode()){
        if symbolInfo.Get_accessMode() == Ast_AccessMode__Global{
            return "lns_" + symbolInfo.Get_name()
        }
        var moduleType *Ast_TypeInfo
        moduleType = symbolInfo.Get_scope().FP.GetNamespaceTypeInfo().FP.GetModule()
        return Lns_getVM().String_format("lns_%s_%s", []LnsAny{self.FP.GetFullName(moduleType), symbolInfo.Get_name()})
    }
    if _switch2939 := symbolInfo.Get_kind(); _switch2939 == Ast_SymbolKind__Var {
        if symbolInfo.Get_symbolId() == convCC_invalidSymbolId{
            return symbolInfo.Get_name()
        }
        return Lns_getVM().String_format("lns_%s_%d", []LnsAny{symbolInfo.Get_name(), symbolInfo.Get_symbolId()})
    }
    return symbolInfo.Get_name()
}

// 649: decl @lune.@base.@convCC.ModuleCtrl.getFormName
func (self *convCC_ModuleCtrl) GetFormName(typeInfo *Ast_TypeInfo) string {
    return Lns_getVM().String_format("l_form_%s", []LnsAny{self.FP.GetFullName(typeInfo)})
}

// 652: decl @lune.@base.@convCC.ModuleCtrl.getCallFormName
func (self *convCC_ModuleCtrl) GetCallFormName(typeInfo *Ast_TypeInfo) string {
    return Lns_getVM().String_format("lns_call_formFunc_%s", []LnsAny{self.FP.GetFullName(typeInfo)})
}

// 656: decl @lune.@base.@convCC.ModuleCtrl.getFuncCastWrapName
func (self *convCC_ModuleCtrl) GetFuncCastWrapName(orgFunc *Ast_TypeInfo,castType *Ast_TypeInfo) string {
    return Lns_getVM().String_format("wrap_%s_2_%s", []LnsAny{self.FP.GetFuncName(orgFunc), self.FP.GetFuncName(castType)})
}

// 660: decl @lune.@base.@convCC.ModuleCtrl.getEnumFuncName
func (self *convCC_ModuleCtrl) GetEnumFuncName(enumType *Ast_EnumTypeInfo,name string) string {
    var scope *Ast_Scope
    scope = Lns_unwrap( enumType.FP.Get_scope()).(*Ast_Scope)
    return self.FP.GetFuncName((Lns_unwrap( scope.FP.GetSymbolInfoChild(name)).(*Ast_SymbolInfo)).FP.Get_typeInfo())
}


// declaration Class -- SymbolParam
type convCC_SymbolParamMtd interface {
}
type convCC_SymbolParam struct {
    Kind LnsInt
    Index LnsInt
    TypeTxt string
    FP convCC_SymbolParamMtd
}
func convCC_SymbolParam2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_SymbolParam).FP
}
type convCC_SymbolParamDownCast interface {
    ToconvCC_SymbolParam() *convCC_SymbolParam
}
func convCC_SymbolParamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_SymbolParamDownCast)
    if ok { return work.ToconvCC_SymbolParam() }
    return nil
}
func (obj *convCC_SymbolParam) ToconvCC_SymbolParam() *convCC_SymbolParam {
    return obj
}
func NewconvCC_SymbolParam(arg1 LnsInt, arg2 LnsInt, arg3 string) *convCC_SymbolParam {
    obj := &convCC_SymbolParam{}
    obj.FP = obj
    obj.InitconvCC_SymbolParam(arg1, arg2, arg3)
    return obj
}
func (self *convCC_SymbolParam) InitconvCC_SymbolParam(arg1 LnsInt, arg2 LnsInt, arg3 string) {
    self.Kind = arg1
    self.Index = arg2
    self.TypeTxt = arg3
}

// declaration Class -- WorkSymbol
type convCC_WorkSymbolMtd interface {
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
type convCC_WorkSymbol struct {
    scope *Ast_Scope
    accessMode LnsInt
    name string
    typeInfo *Ast_TypeInfo
    kind LnsInt
    staticFlag bool
    convModuleParam *convCC_SymbolParam
    FP convCC_WorkSymbolMtd
}
func convCC_WorkSymbol2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_WorkSymbol).FP
}
type convCC_WorkSymbolDownCast interface {
    ToconvCC_WorkSymbol() *convCC_WorkSymbol
}
func convCC_WorkSymbolDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_WorkSymbolDownCast)
    if ok { return work.ToconvCC_WorkSymbol() }
    return nil
}
func (obj *convCC_WorkSymbol) ToconvCC_WorkSymbol() *convCC_WorkSymbol {
    return obj
}
func NewconvCC_WorkSymbol(arg1 *Ast_Scope, arg2 LnsInt, arg3 string, arg4 *Ast_TypeInfo, arg5 LnsInt, arg6 bool, arg7 *convCC_SymbolParam) *convCC_WorkSymbol {
    obj := &convCC_WorkSymbol{}
    obj.FP = obj
    obj.InitconvCC_WorkSymbol(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *convCC_WorkSymbol) InitconvCC_WorkSymbol(arg1 *Ast_Scope, arg2 LnsInt, arg3 string, arg4 *Ast_TypeInfo, arg5 LnsInt, arg6 bool, arg7 *convCC_SymbolParam) {
    self.scope = arg1
    self.accessMode = arg2
    self.name = arg3
    self.typeInfo = arg4
    self.kind = arg5
    self.staticFlag = arg6
    self.convModuleParam = arg7
}
func (self *convCC_WorkSymbol) Get_scope() *Ast_Scope{ return self.scope }
func (self *convCC_WorkSymbol) Get_accessMode() LnsInt{ return self.accessMode }
func (self *convCC_WorkSymbol) Get_name() string{ return self.name }
func (self *convCC_WorkSymbol) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }
func (self *convCC_WorkSymbol) Get_kind() LnsInt{ return self.kind }
func (self *convCC_WorkSymbol) Get_staticFlag() bool{ return self.staticFlag }
func (self *convCC_WorkSymbol) Get_convModuleParam() LnsAny{ return self.convModuleParam.FP }
// 693: decl @lune.@base.@convCC.WorkSymbol.get_mutable
func (self *convCC_WorkSymbol) Get_mutable() bool {
    return false
}

// 696: decl @lune.@base.@convCC.WorkSymbol.get_isLazyLoad
func (self *convCC_WorkSymbol) Get_isLazyLoad() bool {
    return false
}

// 700: decl @lune.@base.@convCC.WorkSymbol.get_symbolId
func (self *convCC_WorkSymbol) Get_symbolId() LnsInt {
    return convCC_invalidSymbolId
}

// 703: decl @lune.@base.@convCC.WorkSymbol.get_hasAccessFromClosure
func (self *convCC_WorkSymbol) Get_hasAccessFromClosure() bool {
    return false
}


// declaration Class -- ScopeInfo
type convCC_ScopeInfoMtd interface {
}
type convCC_ScopeInfo struct {
    AnyNum LnsInt
    StemNum LnsInt
    VarNum LnsInt
    FP convCC_ScopeInfoMtd
}
func convCC_ScopeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_ScopeInfo).FP
}
type convCC_ScopeInfoDownCast interface {
    ToconvCC_ScopeInfo() *convCC_ScopeInfo
}
func convCC_ScopeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_ScopeInfoDownCast)
    if ok { return work.ToconvCC_ScopeInfo() }
    return nil
}
func (obj *convCC_ScopeInfo) ToconvCC_ScopeInfo() *convCC_ScopeInfo {
    return obj
}
func NewconvCC_ScopeInfo(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt) *convCC_ScopeInfo {
    obj := &convCC_ScopeInfo{}
    obj.FP = obj
    obj.InitconvCC_ScopeInfo(arg1, arg2, arg3)
    return obj
}
func (self *convCC_ScopeInfo) InitconvCC_ScopeInfo(arg1 LnsInt, arg2 LnsInt, arg3 LnsInt) {
    self.AnyNum = arg1
    self.StemNum = arg2
    self.VarNum = arg3
}

// declaration Class -- ScopeMgr
type convCC_ScopeMgrMtd interface {
    GetAccessPrimValFromSymbol(arg1 Ast_LowSymbol) string
    GetAccessPrimValFromSymbolOnly(arg1 Ast_LowSymbol) string
    GetCTypeForSym(arg1 Ast_LowSymbol)(string, LnsInt)
    GetSymbolParam(arg1 Ast_LowSymbol) *convCC_SymbolParam
    GetSymbolValKind(arg1 Ast_LowSymbol) LnsInt
    Get_numOf__func__() LnsInt
    Setup(arg1 *Ast_Scope, arg2 *LnsList)
    SetupScopeParam(arg1 *Ast_Scope)(LnsInt, LnsInt, LnsInt)
    setupScopeParamSub(arg1 *Ast_Scope) *convCC_ScopeInfo
    Symbol2Any(arg1 Ast_LowSymbol) string
}
type convCC_ScopeMgr struct {
    scope2InfoMap *LnsMap
    moduleCtrl *convCC_ModuleCtrl
    numOf__func__ LnsInt
    moduleBlockAnyNum LnsInt
    FP convCC_ScopeMgrMtd
}
func convCC_ScopeMgr2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_ScopeMgr).FP
}
type convCC_ScopeMgrDownCast interface {
    ToconvCC_ScopeMgr() *convCC_ScopeMgr
}
func convCC_ScopeMgrDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_ScopeMgrDownCast)
    if ok { return work.ToconvCC_ScopeMgr() }
    return nil
}
func (obj *convCC_ScopeMgr) ToconvCC_ScopeMgr() *convCC_ScopeMgr {
    return obj
}
func NewconvCC_ScopeMgr(arg1 *convCC_ModuleCtrl) *convCC_ScopeMgr {
    obj := &convCC_ScopeMgr{}
    obj.FP = obj
    obj.InitconvCC_ScopeMgr(arg1)
    return obj
}
func (self *convCC_ScopeMgr) Get_numOf__func__() LnsInt{ return self.numOf__func__ }
// 785: DeclConstr
func (self *convCC_ScopeMgr) InitconvCC_ScopeMgr(moduleCtrl *convCC_ModuleCtrl) {
    self.scope2InfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.moduleCtrl = moduleCtrl
    
    self.numOf__func__ = 0
    
    self.moduleBlockAnyNum = 0
    
}

// 792: decl @lune.@base.@convCC.ScopeMgr.setSymbolParam
func convCC_ScopeMgr_setSymbolParam_1456_(scopeInfo *convCC_ScopeInfo,symbol *Ast_SymbolInfo) *convCC_SymbolParam {
    var param *convCC_SymbolParam
    if _switch3686 := convCC_getValKind_1167_(symbol.FP.Get_typeInfo()); _switch3686 == convCC_ValKind__Stem {
        param = NewconvCC_SymbolParam(convCC_ValKind__Stem, scopeInfo.StemNum, convCC_cTypeStem)
        
        scopeInfo.StemNum = scopeInfo.StemNum + 1
        
    } else if _switch3686 == convCC_ValKind__Any {
        if symbol.FP.Get_name() == "self"{
            param = NewconvCC_SymbolParam(convCC_ValKind__Any, 0, convCC_cTypeAnyP)
            
        } else { 
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( symbol.FP.Get_kind() == Ast_SymbolKind__Var) ||
                Lns_GetEnv().SetStackVal( symbol.FP.Get_kind() == Ast_SymbolKind__Mbr) ).(bool){
                param = NewconvCC_SymbolParam(convCC_ValKind__Any, scopeInfo.AnyNum, convCC_cTypeAnyPP)
                
                scopeInfo.AnyNum = scopeInfo.AnyNum + 1
                
            } else { 
                if symbol.FP.Get_mutable(){
                    param = NewconvCC_SymbolParam(convCC_ValKind__Any, scopeInfo.AnyNum, convCC_cTypeAnyPP)
                    
                    scopeInfo.AnyNum = scopeInfo.AnyNum + 1
                    
                } else { 
                    param = NewconvCC_SymbolParam(convCC_ValKind__Any, 0, convCC_cTypeAnyP)
                    
                }
            }
        }
    } else if _switch3686 == convCC_ValKind__Prim {
        param = NewconvCC_SymbolParam(convCC_ValKind__Prim, 0, convCC_getCType_1188_(symbol.FP.Get_typeInfo()))
        
    } else {
        Util_err(Lns_getVM().String_format("not support %s", []LnsAny{symbol.FP.Get_typeInfo().FP.GetTxt(nil, nil, nil)}))
    }
    return param
}

// 832: decl @lune.@base.@convCC.ScopeMgr.setupScopeParamSub
func (self *convCC_ScopeMgr) setupScopeParamSub(scope *Ast_Scope) *convCC_ScopeInfo {
    {
        _scopeInfo := self.scope2InfoMap.Items[scope]
        if _scopeInfo != nil {
            scopeInfo := _scopeInfo.(*convCC_ScopeInfo)
            return scopeInfo
        }
    }
    var scopeInfo *convCC_ScopeInfo
    scopeInfo = NewconvCC_ScopeInfo(0, 0, 0)
    
    if Lns_isCondTrue( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(scope.FP.Get_ownerTypeInfo()) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_TypeInfo).FP.IsModule()})/* 843:13 */)){
        scopeInfo.AnyNum = 2
        
    }
    {
        __collection4049 := scope.FP.Get_symbol2SymbolInfoMap()
        __sorted4049 := __collection4049.CreateKeyListStr()
        __sorted4049.Sort( LnsItemKindStr, nil )
        for _, ___key4049 := range( __sorted4049.Items ) {
            symbol := __collection4049.Items[ ___key4049 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if Lns_op_not(symbol.FP.Get_convModuleParam()){
                var param *convCC_SymbolParam
                if symbol.FP.Get_name() != "__func__"{
                    if _switch3998 := symbol.FP.Get_kind(); _switch3998 == Ast_SymbolKind__Var || _switch3998 == Ast_SymbolKind__Arg {
                        if symbol.FP.Get_hasAccessFromClosure(){
                            param = NewconvCC_SymbolParam(convCC_ValKind__Var, scopeInfo.VarNum, convCC_cTypeVarP)
                            
                            scopeInfo.VarNum = scopeInfo.VarNum + 1
                            
                        } else { 
                            param = convCC_ScopeMgr_setSymbolParam_1456_(scopeInfo, symbol)
                            
                        }
                    } else if _switch3998 == Ast_SymbolKind__Fun {
                        if symbol.FP.Get_hasAccessFromClosure(){
                            param = NewconvCC_SymbolParam(convCC_ValKind__Var, scopeInfo.VarNum, convCC_cTypeVarP)
                            
                            scopeInfo.VarNum = scopeInfo.VarNum + 1
                            
                        } else { 
                            param = convCC_createSymbolParam_1436_(symbol.FP.Get_name(), convCC_getValKind_1167_(symbol.FP.Get_typeInfo()), convCC_getCType_1188_(symbol.FP.Get_typeInfo()))
                            
                        }
                    } else if _switch3998 == Ast_SymbolKind__Mtd {
                        var retTypeList *LnsList
                        retTypeList = symbol.FP.Get_typeInfo().FP.Get_retTypeInfoList()
                        param = convCC_createSymbolParam_1436_(symbol.FP.Get_name(), convCC_getRetKind_1178_(retTypeList), convCC_getCRetType_1195_(retTypeList))
                        
                    } else if _switch3998 == Ast_SymbolKind__Mbr {
                        if convCC_isClassMember_1216_(symbol.FP){
                            param = (Lns_unwrap( symbol.FP.Get_convModuleParam())).(*convCC_SymbolParam)
                            
                        } else { 
                            param = convCC_createSymbolParam_1436_(symbol.FP.Get_name(), convCC_getValKind_1167_(symbol.FP.Get_typeInfo()), convCC_getCType_1188_(symbol.FP.Get_typeInfo()))
                            
                        }
                    } else {
                        param = NewconvCC_SymbolParam(convCC_ValKind__Other, 0, convCC_cTypeStem)
                        
                    }
                } else { 
                    param = NewconvCC_SymbolParam(convCC_ValKind__Any, self.numOf__func__ + self.moduleBlockAnyNum, convCC_cTypeAnyP)
                    
                    self.numOf__func__ = self.numOf__func__ + 1
                    
                }
                symbol.FP.Set_convModuleParam(param.FP)
            }
        }
    }
    self.scope2InfoMap.Set(scope,scopeInfo)
    return scopeInfo
}

// 913: decl @lune.@base.@convCC.ScopeMgr.setup
func (self *convCC_ScopeMgr) Setup(scope *Ast_Scope,declMemberList *LnsList) {
    self.FP.setupScopeParamSub(scope)
    for _, _declMember := range( declMemberList.Items ) {
        declMember := _declMember.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var scopeInfo *convCC_ScopeInfo
        scopeInfo = Lns_unwrap( self.scope2InfoMap.Items[scope]).(*convCC_ScopeInfo)
        var symbol *Ast_SymbolInfo
        symbol = declMember.FP.Get_symbolInfo()
        if convCC_isClassMember_1216_(symbol.FP){
            symbol.FP.Set_convModuleParam(convCC_ScopeMgr_setSymbolParam_1456_(scopeInfo, symbol).FP)
        }
    }
}

// 926: decl @lune.@base.@convCC.ScopeMgr.create
func convCC_ScopeMgr_create_1469_(moduleCtrl *convCC_ModuleCtrl,initBlockScope *Ast_Scope) *convCC_ScopeMgr {
    var scopeMgr *convCC_ScopeMgr
    scopeMgr = NewconvCC_ScopeMgr(moduleCtrl)
    var param *convCC_ScopeInfo
    param = scopeMgr.FP.setupScopeParamSub(initBlockScope)
    scopeMgr.moduleBlockAnyNum = param.AnyNum
    
    return scopeMgr
}

// 934: decl @lune.@base.@convCC.ScopeMgr.setupScopeParam
func (self *convCC_ScopeMgr) SetupScopeParam(scope *Ast_Scope)(LnsInt, LnsInt, LnsInt) {
    var scopeInfo *convCC_ScopeInfo
    scopeInfo = self.FP.setupScopeParamSub(scope)
    return scopeInfo.AnyNum, scopeInfo.StemNum, scopeInfo.VarNum
}

// 940: decl @lune.@base.@convCC.ScopeMgr.getSymbolParam
func (self *convCC_ScopeMgr) GetSymbolParam(symbol Ast_LowSymbol) *convCC_SymbolParam {
    {
        _param := symbol.Get_convModuleParam()
        if _param != nil {
            param := _param
            return param.(*convCC_SymbolParam)
        }
    }
    var scope *Ast_Scope
    scope = symbol.Get_scope()
    if Lns_op_not(self.scope2InfoMap.Items[scope]){
        self.FP.SetupScopeParam(scope)
        {
            _param := symbol.Get_convModuleParam()
            if _param != nil {
                param := _param
                return param.(*convCC_SymbolParam)
            }
        }
    }
    Util_err(Lns_getVM().String_format("illegal symbol -- %s %s %s %d", []LnsAny{symbol.Get_name(), Ast_SymbolKind_getTxt( symbol.Get_kind()), self.moduleCtrl.FP.GetCanonicalName(symbol.Get_scope().FP.GetNamespaceTypeInfo()), 956}))
// insert a dummy
    return nil
}

// 959: decl @lune.@base.@convCC.ScopeMgr.getSymbolValKind
func (self *convCC_ScopeMgr) GetSymbolValKind(symbol Ast_LowSymbol) LnsInt {
    var symbolParam *convCC_SymbolParam
    symbolParam = self.FP.GetSymbolParam(symbol)
    return symbolParam.Kind
}

// 965: decl @lune.@base.@convCC.ScopeMgr.getCTypeForSym
func (self *convCC_ScopeMgr) GetCTypeForSym(symbol Ast_LowSymbol)(string, LnsInt) {
    var param *convCC_SymbolParam
    param = self.FP.GetSymbolParam(symbol)
    return param.TypeTxt, param.Kind
}


// 987: decl @lune.@base.@convCC.ScopeMgr.getAccessPrimValFromSymbol
func (self *convCC_ScopeMgr) GetAccessPrimValFromSymbol(symbolInfo Ast_LowSymbol) string {
    return convCC_getAccessPrimValFromSymbolDirect_1433_(self.moduleCtrl.FP.GetSymbolName(symbolInfo), self.FP.GetSymbolValKind(symbolInfo), symbolInfo.Get_typeInfo())
}

// 994: decl @lune.@base.@convCC.ScopeMgr.getAccessPrimValFromSymbolOnly
func (self *convCC_ScopeMgr) GetAccessPrimValFromSymbolOnly(symbolInfo Ast_LowSymbol) string {
    return convCC_getAccessPrimValFromSymbolDirect_1433_("", self.FP.GetSymbolValKind(symbolInfo), symbolInfo.Get_typeInfo())
}

// 2162: decl @lune.@base.@convCC.ScopeMgr.symbol2Any
func (self *convCC_ScopeMgr) Symbol2Any(symbol Ast_LowSymbol) string {
    if symbol.Get_kind() == Ast_SymbolKind__Fun{
        return convCC_getFunc2any_1897_(self.moduleCtrl, self, symbol.Get_typeInfo())
    }
    var name string
    name = self.moduleCtrl.FP.GetSymbolName(symbol)
    if _switch10595 := self.FP.GetSymbolValKind(symbol); _switch10595 == convCC_ValKind__Var {
        return name
    } else {
        Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{symbol.Get_typeInfo().FP.GetTxt(nil, nil, nil)}))
    }
// insert a dummy
    return ""
}


// declaration Class -- convFilter
type convCC_convFilterMtd interface {
    AccessPrimVal(arg1 *Nodes_Node, arg2 *Nodes_Node)
    accessPrimValFromAny(arg1 bool, arg2 *Ast_TypeInfo, arg3 LnsInt)
    createRefNodeFromSym(arg1 *Ast_SymbolInfo) *Nodes_ExpRefNode
    DefaultProcess(arg1 *Nodes_Node, arg2 LnsAny)
    getFullName(arg1 *Ast_TypeInfo) string
    getValKindOfNode(arg1 *Nodes_Node) LnsInt
    Get_moduleInfoManager() Ast_ModuleInfoManager
    Get_optStack() *LnsList
    Get_typeNameCtrl() *Ast_TypeNameCtrl
    isManagedAnySymbol(arg1 Ast_LowSymbol) bool
    isStemSym(arg1 Ast_LowSymbol) bool
    isStemVal(arg1 *Nodes_Node) bool
    needsWrapper(arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo) bool
    outputAlter2MapFunc(arg1 Util_SourceStream, arg2 *LnsMap)
    OutputDeclMacro(arg1 string, arg2 *LnsList, arg3 convCC_outputMacroStmtBlock_2219_)
    PopIndent()
    popRoutine()
    popStream()
    ProcessAbbr(arg1 *Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(arg1 *Nodes_AliasNode, arg2 LnsAny)
    processAndOr(arg1 *Nodes_ExpOp2Node, arg2 string, arg3 *Nodes_Node)
    ProcessApply(arg1 *Nodes_ApplyNode, arg2 LnsAny)
    processArgClosure(arg1 *Nodes_DeclFuncInfo)
    ProcessBlankLine(arg1 *Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(arg1 *Nodes_BlockNode, arg2 LnsAny)
    processBlockPostProcess()
    processBlockPreProcess(arg1 *Ast_Scope)
    ProcessBlockSub(arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(arg1 *Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(arg1 *Nodes_BreakNode, arg2 LnsAny)
    processBuiltin() *Nodes_NodeManager
    processCall(arg1 LnsAny, arg2 *Ast_TypeInfo, arg3 bool, arg4 LnsAny)
    ProcessCallArgList(arg1 *Ast_TypeInfo, arg2 LnsAny)
    processCallUserForm(arg1 string, arg2 *Ast_TypeInfo, arg3 *LnsList)
    ProcessCallWithMRet(arg1 *Nodes_Node, arg2 string, arg3 string, arg4 LnsAny, arg5 *Nodes_ExpListNode)
    processClosureFunc(arg1 *Nodes_DeclFuncInfo)
    processConcat(arg1 *Nodes_ExpOp2Node, arg2 *Nodes_Node)
    ProcessConvStat(arg1 *Nodes_ConvStatNode, arg2 LnsAny)
    processCreateDDD(arg1 *Nodes_Node, arg2 *LnsList)
    processCreateMRet(arg1 *LnsList, arg2 *LnsList, arg3 *Nodes_Node)
    ProcessDeclAdvertise(arg1 *Nodes_DeclAdvertiseNode, arg2 LnsAny)
    ProcessDeclAlge(arg1 *Nodes_DeclAlgeNode, arg2 LnsAny)
    ProcessDeclArg(arg1 *Nodes_DeclArgNode, arg2 LnsAny)
    ProcessDeclArgDDD(arg1 *Nodes_DeclArgDDDNode, arg2 LnsAny)
    ProcessDeclClass(arg1 *Nodes_DeclClassNode, arg2 LnsAny)
    processDeclClassDef(arg1 *Nodes_DeclClassNode)
    processDeclClassNodePrototype(arg1 *Nodes_DeclClassNode)
    ProcessDeclConstr(arg1 *Nodes_DeclConstrNode, arg2 LnsAny)
    ProcessDeclDestr(arg1 *Nodes_DeclDestrNode, arg2 LnsAny)
    ProcessDeclEnum(arg1 *Nodes_DeclEnumNode, arg2 LnsAny)
    ProcessDeclForm(arg1 *Nodes_DeclFormNode, arg2 LnsAny)
    ProcessDeclFunc(arg1 *Nodes_DeclFuncNode, arg2 LnsAny)
    ProcessDeclMacro(arg1 *Nodes_DeclMacroNode, arg2 LnsAny)
    ProcessDeclMember(arg1 *Nodes_DeclMemberNode, arg2 LnsAny)
    ProcessDeclMethod(arg1 *Nodes_DeclMethodNode, arg2 LnsAny)
    processDeclMethodInfo(arg1 *Nodes_DeclFuncInfo, arg2 *Ast_TypeInfo, arg3 *Nodes_Node)
    ProcessDeclVar(arg1 *Nodes_DeclVarNode, arg2 LnsAny)
    processDeclVarAndSet(arg1 *LnsList, arg2 LnsAny)
    processDeclVarC(arg1 bool, arg2 Ast_LowSymbol, arg3 bool, arg4 LnsAny)
    ProcessEnv(arg1 *Nodes_EnvNode, arg2 LnsAny)
    processEquals(arg1 bool, arg2 *Ast_TypeInfo, arg3 *Ast_TypeInfo, arg4 convCC_ProcessToValForm_2538_, arg5 convCC_ProcessToValForm_2538_)
    ProcessExpAccessMRet(arg1 *Nodes_ExpAccessMRetNode, arg2 LnsAny)
    ProcessExpCall(arg1 *Nodes_ExpCallNode, arg2 LnsAny)
    processExpCallDefWrap(arg1 *Nodes_ExpCallNode, arg2 *ConvCC_Opt)
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
    processForeachSetupVal(arg1 *Nodes_Node, arg2 *Ast_Scope, arg3 string, arg4 string, arg5 *Ast_TypeInfo)
    ProcessForsort(arg1 *Nodes_ForsortNode, arg2 LnsAny)
    processFuncCast(arg1 *Nodes_ExpCastNode)
    processFuncCast2Form(arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo)
    processFuncPrototype(arg1 *Nodes_Node, arg2 LnsInt, arg3 bool, arg4 string, arg5 string, arg6 *LnsList, arg7 bool)
    ProcessGetField(arg1 *Nodes_GetFieldNode, arg2 LnsAny)
    ProcessIf(arg1 *Nodes_IfNode, arg2 LnsAny)
    ProcessIfUnwrap(arg1 *Nodes_IfUnwrapNode, arg2 LnsAny)
    ProcessImport(arg1 *Nodes_ImportNode, arg2 LnsAny)
    processInitModule(arg1 *Nodes_RootNode)
    ProcessLiteralArray(arg1 *Nodes_LiteralArrayNode, arg2 LnsAny)
    ProcessLiteralBool(arg1 *Nodes_LiteralBoolNode, arg2 LnsAny)
    ProcessLiteralChar(arg1 *Nodes_LiteralCharNode, arg2 LnsAny)
    ProcessLiteralInt(arg1 *Nodes_LiteralIntNode, arg2 LnsAny)
    ProcessLiteralList(arg1 *Nodes_LiteralListNode, arg2 LnsAny)
    processLiteralListSub(arg1 string, arg2 *Nodes_Node, arg3 LnsAny, arg4 string)
    ProcessLiteralMap(arg1 *Nodes_LiteralMapNode, arg2 LnsAny)
    processLiteralMapSub(arg1 *Nodes_LiteralMapNode)
    ProcessLiteralNil(arg1 *Nodes_LiteralNilNode, arg2 LnsAny)
    processLiteralNode(arg1 *Nodes_Node, arg2 *Nodes_Node)
    ProcessLiteralReal(arg1 *Nodes_LiteralRealNode, arg2 LnsAny)
    ProcessLiteralSet(arg1 *Nodes_LiteralSetNode, arg2 LnsAny)
    ProcessLiteralString(arg1 *Nodes_LiteralStringNode, arg2 LnsAny)
    ProcessLiteralSymbol(arg1 *Nodes_LiteralSymbolNode, arg2 LnsAny)
    processLiteralVal(arg1 *Nodes_Node, arg2 *Nodes_Node)
    processLoopPostProcess()
    processLoopPreProcess(arg1 *Nodes_BlockNode)
    ProcessLuneControl(arg1 *Nodes_LuneControlNode, arg2 LnsAny)
    ProcessLuneKind(arg1 *Nodes_LuneKindNode, arg2 LnsAny)
    processMapForeachSetupVal(arg1 *Nodes_Node, arg2 *Ast_TypeInfo, arg3 *Ast_Scope, arg4 LnsAny, arg5 LnsAny, arg6 string, arg7 string)
    processMapping(arg1 *Nodes_DeclClassNode, arg2 *Ast_TypeInfo, arg3 LnsInt)
    ProcessMatch(arg1 *Nodes_MatchNode, arg2 LnsAny)
    ProcessNewAlgeVal(arg1 *Nodes_NewAlgeValNode, arg2 LnsAny)
    processNewInsance(arg1 *Ast_TypeInfo, arg2 bool)
    ProcessNone(arg1 *Nodes_NoneNode, arg2 LnsAny)
    processPoolForeachSetupVal(arg1 *Nodes_Node, arg2 *Ast_TypeInfo, arg3 *Ast_Scope, arg4 LnsAny, arg5 LnsAny)
    ProcessProtoClass(arg1 *Nodes_ProtoClassNode, arg2 LnsAny)
    ProcessProtoMethod(arg1 *Nodes_ProtoMethodNode, arg2 LnsAny)
    ProcessProvide(arg1 *Nodes_ProvideNode, arg2 LnsAny)
    ProcessRefField(arg1 *Nodes_RefFieldNode, arg2 LnsAny)
    ProcessRefType(arg1 *Nodes_RefTypeNode, arg2 LnsAny)
    ProcessRepeat(arg1 *Nodes_RepeatNode, arg2 LnsAny)
    ProcessReturn(arg1 *Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(arg1 *Nodes_RootNode, arg2 LnsAny)
    ProcessScope(arg1 *Nodes_ScopeNode, arg2 LnsAny)
    processSetSymSingle(arg1 *Nodes_Node, arg2 LnsAny, arg3 Ast_LowSymbol, arg4 bool, arg5 Ast_LowSymbol, arg6 bool)
    processSetValSingle(arg1 *Nodes_Node, arg2 LnsAny, arg3 Ast_LowSymbol, arg4 bool, arg5 *Nodes_Node, arg6 LnsInt, arg7 bool)
    processSetValSingleDirect(arg1 *Nodes_Node, arg2 LnsAny, arg3 Ast_LowSymbol, arg4 bool, arg5 LnsInt, arg6 *Ast_TypeInfo, arg7 LnsInt, arg8 bool, arg9 convCC_processRValue_2298_)
    processSetValSingleNode(arg1 *Nodes_Node, arg2 *Nodes_Node, arg3 bool, arg4 *Nodes_Node, arg5 LnsInt, arg6 bool)
    processSetValToDst(arg1 *Nodes_Node, arg2 *LnsList, arg3 *LnsList, arg4 LnsAny)
    processSetValToNode(arg1 *Nodes_Node, arg2 *Nodes_Node, arg3 *LnsSet, arg4 *LnsList, arg5 LnsAny)
    processSetValToSym(arg1 *Nodes_Node, arg2 *LnsList, arg3 bool, arg4 *LnsList, arg5 LnsAny, arg6 LnsAny)
    processStme2Val(arg1 *Ast_TypeInfo, arg2 string)
    ProcessStmtExp(arg1 *Nodes_StmtExpNode, arg2 LnsAny)
    ProcessSubfile(arg1 *Nodes_SubfileNode, arg2 LnsAny)
    ProcessSwitch(arg1 *Nodes_SwitchNode, arg2 LnsAny)
    ProcessSym2Any(arg1 Ast_LowSymbol)
    processSym2stem(arg1 Ast_LowSymbol)
    processSymForSetOp(arg1 *Nodes_Node, arg2 LnsInt, arg3 *Ast_TypeInfo, arg4 Ast_LowSymbol)
    ProcessTestBlock(arg1 *Nodes_TestBlockNode, arg2 LnsAny)
    ProcessTestCase(arg1 *Nodes_TestCaseNode, arg2 LnsAny)
    ProcessUnboxing(arg1 *Nodes_UnboxingNode, arg2 LnsAny)
    ProcessUnwrapSet(arg1 *Nodes_UnwrapSetNode, arg2 LnsAny)
    ProcessVal2any(arg1 *Nodes_Node, arg2 *Nodes_Node)
    ProcessVal2stem(arg1 *Nodes_Node, arg2 *Nodes_Node)
    processValForSetOp(arg1 *Nodes_Node, arg2 LnsInt, arg3 *Ast_TypeInfo, arg4 *Nodes_Node, arg5 LnsInt, arg6 bool)
    ProcessWhen(arg1 *Nodes_WhenNode, arg2 LnsAny)
    ProcessWhile(arg1 *Nodes_WhileNode, arg2 LnsAny)
    processWrapForm2Func(arg1 *Ast_TypeInfo)
    process__func__symbol(arg1 *Ast_TypeInfo, arg2 bool, arg3 string)
    PushIndent(arg1 LnsAny)
    pushRoutine(arg1 *Ast_TypeInfo, arg2 *Nodes_BlockNode)
    pushStream() *Util_memStream
    ReturnToSource()
    SwitchToHeader()
    Write(arg1 string)
    Writeln(arg1 string)
}
type convCC_convFilter struct {
    Nodes_Filter
    streamName string
    stream *Util_SimpleSourceOStream
    streamQueue *LnsList
    moduleCtrl *convCC_ModuleCtrl
    pubVarName2InfoMap *LnsMap
    pubFuncName2InfoMap *LnsMap
    moduleTypeInfo *Ast_TypeInfo
    ast *Nodes_RootNode
    routineInfoStack *convCC_DepthStack
    loopInfoStack *convCC_DepthStack
    processMode LnsInt
    duringDeclFunc bool
    processingNode LnsAny
    processedNodeSet *LnsSet
    accessSymbolSet *Util_OrderedSet
    literalNode2AccessSymbolSet *LnsMap
    scopeMgr *convCC_ScopeMgr
    outputBuiltinFlag bool
    enableTest bool
    canConv bool
    dummyNodeManager *Nodes_NodeManager
    processInfo *Ast_ProcessInfo
    FP convCC_convFilterMtd
}
func convCC_convFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_convFilter).FP
}
type convCC_convFilterDownCast interface {
    ToconvCC_convFilter() *convCC_convFilter
}
func convCC_convFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_convFilterDownCast)
    if ok { return work.ToconvCC_convFilter() }
    return nil
}
func (obj *convCC_convFilter) ToconvCC_convFilter() *convCC_convFilter {
    return obj
}
func NewconvCC_convFilter(arg1 bool, arg2 bool, arg3 string, arg4 Lns_oStream, arg5 Lns_oStream, arg6 *TransUnit_ASTInfo) *convCC_convFilter {
    obj := &convCC_convFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvCC_convFilter(arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *convCC_convFilter) PopIndent() {
self.stream. FP.PopIndent( )
}
func (self *convCC_convFilter) PushIndent(arg1 LnsAny) {
self.stream. FP.PushIndent( arg1)
}
func (self *convCC_convFilter) ReturnToSource() {
self.stream. FP.ReturnToSource( )
}
func (self *convCC_convFilter) SwitchToHeader() {
self.stream. FP.SwitchToHeader( )
}
func (self *convCC_convFilter) Write(arg1 string) {
self.stream. FP.Write( arg1)
}
func (self *convCC_convFilter) Writeln(arg1 string) {
self.stream. FP.Writeln( arg1)
}
// 1112: decl @lune.@base.@convCC.convFilter.createRefNodeFromSym
func (self *convCC_convFilter) createRefNodeFromSym(symbol *Ast_SymbolInfo) *Nodes_ExpRefNode {
    return Nodes_ExpRefNode_create(self.dummyNodeManager, Lns_unwrap( symbol.FP.Get_pos()).(*Types_Position), false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(symbol.FP.Get_typeInfo())}), symbol)
}

// 1117: DeclConstr
func (self *convCC_convFilter) InitconvCC_convFilter(enableTest bool,outputBuiltin bool,streamName string,stream Lns_oStream,headerStream Lns_oStream,ast *TransUnit_ASTInfo) {
    self.InitNodes_Filter(true, ast.FP.Get_moduleTypeInfo(), ast.FP.Get_moduleTypeInfo().FP.Get_scope())
    self.processInfo = ast.FP.Get_processInfo()
    
    self.dummyNodeManager = NewNodes_NodeManager()
    
    self.canConv = true
    
    self.enableTest = enableTest
    
    self.outputBuiltinFlag = outputBuiltin
    
    self.processingNode = nil
    
    self.processedNodeSet = NewLnsSet([]LnsAny{})
    
    self.accessSymbolSet = NewUtil_OrderedSet()
    
    self.literalNode2AccessSymbolSet = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.duringDeclFunc = false
    
    self.processMode = convCC_ProcessMode__Prototype
    
    self.moduleTypeInfo = ast.FP.Get_moduleTypeInfo()
    
    self.routineInfoStack = NewconvCC_DepthStack()
    
    self.routineInfoStack.FP.NewInfo(&NewconvCC_RoutineInfo(Ast_builtinTypeNone).convCC_DepthInfo)
    self.routineInfoStack.FP.NewInfo(&NewconvCC_RoutineInfo(ast.FP.Get_moduleTypeInfo()).convCC_DepthInfo)
    self.loopInfoStack = NewconvCC_DepthStack()
    
    self.loopInfoStack.FP.NewInfo(NewconvCC_DepthInfo())
    self.loopInfoStack.FP.NewInfo(NewconvCC_DepthInfo())
    self.ast = Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node().FP)).(*Nodes_RootNode)
    
    self.streamName = streamName
    
    self.streamQueue = NewLnsList([]LnsAny{})
    
    self.pubVarName2InfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.pubFuncName2InfoMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.moduleCtrl = NewconvCC_ModuleCtrl(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager())
    
    self.scopeMgr = convCC_ScopeMgr_create_1469_(self.moduleCtrl, Lns_unwrap( ast.FP.Get_moduleTypeInfo().FP.Get_scope()).(*Ast_Scope))
    
    self.stream = NewUtil_SimpleSourceOStream(stream, headerStream, convCC_stepIndent)
    
}

// 1176: decl @lune.@base.@convCC.convFilter.pushStream
func (self *convCC_convFilter) pushStream() *Util_memStream {
    self.streamQueue.Insert(Util_SimpleSourceOStream2Stem(self.stream))
    var stream *Util_memStream
    stream = NewUtil_memStream()
    self.stream = NewUtil_SimpleSourceOStream(stream.FP, nil, convCC_stepIndent)
    
    return stream
}

// 1186: decl @lune.@base.@convCC.convFilter.popStream
func (self *convCC_convFilter) popStream() {
    if self.streamQueue.Len() == 0{
        Util_err("streamQueue is empty.")
    }
    self.stream = self.streamQueue.GetAt(self.streamQueue.Len()).(Util_SimpleSourceOStreamDownCast).ToUtil_SimpleSourceOStream()
    
    self.streamQueue.Remove(nil)
}

// 1197: decl @lune.@base.@convCC.convFilter.getFullName
func (self *convCC_convFilter) getFullName(typeInfo *Ast_TypeInfo) string {
    return self.moduleCtrl.FP.GetFullName(typeInfo)
}



// 1254: decl @lune.@base.@convCC.convFilter.processNone
func (self *convCC_convFilter) ProcessNone(node *Nodes_NoneNode,_opt LnsAny) {
}

// 1260: decl @lune.@base.@convCC.convFilter.processImport
func (self *convCC_convFilter) ProcessImport(node *Nodes_ImportNode,_opt LnsAny) {
    if self.processMode == convCC_ProcessMode__Include{
        var process func(out2HMode LnsInt)
        process = func(out2HMode LnsInt) {
            if _switch5266 := out2HMode; _switch5266 == convCC_Out2HMode__HeaderPub {
                if node.FP.Get_symbolInfo().FP.Get_scope() != node.FP.Get_moduleTypeInfo().FP.Get_scope(){
                    return 
                }
            } else if _switch5266 == convCC_Out2HMode__SourcePub || _switch5266 == convCC_Out2HMode__SourcePri {
            } else {
                return 
            }
            self.FP.Writeln(Lns_getVM().String_format("#include<%s.h>", []LnsAny{Lns_car(Lns_getVM().String_gsub(node.FP.Get_modulePath(),"%.", "/")).(string)}))
        }
        {
            var processwork func(out2HMode LnsInt)
            processwork = func(out2HMode LnsInt) {
                process(out2HMode)
            }
            if true{
                self.stream.FP.SwitchToHeader()
                processwork(convCC_Out2HMode__HeaderPub)
                self.stream.FP.ReturnToSource()
                processwork(convCC_Out2HMode__SourcePub)
            } else { 
                processwork(convCC_Out2HMode__SourcePri)
            }
        }
        
    } else { 
        self.FP.Writeln(Lns_getVM().String_format("lns_init_%s( _pEnv );", []LnsAny{self.moduleCtrl.FP.GetFullName(node.FP.Get_moduleTypeInfo())}))
    }
}

// 1296: decl @lune.@base.@convCC.convFilter.processInitModule
func (self *convCC_convFilter) processInitModule(node *Nodes_RootNode) {
    var anyNum LnsInt
    var stemNum LnsInt
    var varNum LnsInt
    anyNum,stemNum,varNum = self.scopeMgr.FP.SetupScopeParam(self.ast.FP.Get_moduleScope())
    self.processMode = convCC_ProcessMode__InitModule
    
    var moduleFullName string
    moduleFullName = self.moduleCtrl.FP.GetFullName(node.FP.Get_moduleTypeInfo())
    var moduleInfoName string
    moduleInfoName = Lns_getVM().String_format("lns_moduleInfo_%s", []LnsAny{moduleFullName})
    if self.outputBuiltinFlag{
        self.FP.Writeln("static void lns_init_lns_builtin_Sub( lns_env_t * _pEnv );")
    } else if Lns_op_not(self.canConv){
        self.FP.Writeln(Lns_getVM().String_format("extern void lns_init_%s_Sub( lns_env_t * _pEnv );", []LnsAny{moduleFullName}))
    }
    var process func(out2HMode LnsInt)
    process = func(out2HMode LnsInt) {
        self.FP.Write(Lns_getVM().String_format("%s%s lns_init_%s( %s _pEnv )", []LnsAny{convCC_getOut2HeaderPrefix_1518_(out2HMode), convCC_cTypeModP, moduleFullName, convCC_cTypeEnvP}))
        if _switch5568 := out2HMode; _switch5568 == convCC_Out2HMode__HeaderPub {
            self.FP.Writeln(";")
        } else if _switch5568 == convCC_Out2HMode__SourcePub {
            self.FP.Writeln("{")
        }
    }
    {
        var processwork func(out2HMode LnsInt)
        processwork = func(out2HMode LnsInt) {
            process(out2HMode)
        }
        if true{
            self.stream.FP.SwitchToHeader()
            processwork(convCC_Out2HMode__HeaderPub)
            self.stream.FP.ReturnToSource()
            processwork(convCC_Out2HMode__SourcePub)
        } else { 
            processwork(convCC_Out2HMode__SourcePri)
        }
    }
    
    self.FP.PushIndent(nil)
    self.FP.Writeln(Lns_getVM().String_format("if ( %s.readyFlag ) {", []LnsAny{moduleInfoName}))
    self.FP.PushIndent(nil)
    self.FP.Writeln(Lns_getVM().String_format("return &%s;", []LnsAny{moduleInfoName}))
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.FP.Writeln(Lns_getVM().String_format("%s.readyFlag = true;", []LnsAny{moduleInfoName}))
    self.FP.Writeln(Lns_getVM().String_format("lns_add2list( &_pEnv->loadModuleTop, &%s);", []LnsAny{moduleInfoName}))
    self.FP.Writeln("")
    var moduleBlockName string
    moduleBlockName = convCC_getBlockName_1198_(self.ast.FP.Get_moduleScope())
    self.FP.Writeln(Lns_getVM().String_format("lns_block_t * %s = lns_enter_module( _pEnv, %d, %d, %d );", []LnsAny{moduleBlockName, anyNum + self.scopeMgr.FP.Get_numOf__func__(), stemNum, varNum}))
    self.FP.Writeln(Lns_getVM().String_format("%s.pBlock = %s;", []LnsAny{moduleInfoName, moduleBlockName}))
    self.FP.Writeln(Lns_getVM().String_format("lns_set_block_any( %s, 0, lns_module_globalStemList);", []LnsAny{moduleBlockName}))
    self.FP.Writeln("lns_setQ_any( lns_module_globalStemList, lns_class_List_new( _pEnv ));")
    self.FP.Writeln(Lns_getVM().String_format("lns_set_block_any( %s, 1, lns_module_path);", []LnsAny{moduleBlockName}))
    self.FP.Writeln(Lns_getVM().String_format("lns_setQ_any( lns_module_path, lns_litStr2any( _pEnv, \"%s\"));", []LnsAny{node.FP.Get_moduleTypeInfo().FP.GetFullName(self.FP.Get_typeNameCtrl(), self.FP.Get_moduleInfoManager(), nil)}))
    for _, _blockNode := range( node.FP.Get_nodeManager().FP.GetBlockNodeList().Items ) {
        blockNode := _blockNode.(Nodes_BlockNodeDownCast).ToNodes_BlockNode()
        self.scopeMgr.FP.SetupScopeParam(blockNode.FP.Get_scope())
    }
    self.FP.Writeln("lns_enter_block( _pEnv, 0, 0, 0 );")
    if self.canConv{
        self.FP.Writeln(Lns_getVM().String_format("initFuncSym( _pEnv, %s );", []LnsAny{moduleBlockName}))
    }
    self.FP.Writeln("")
    if self.outputBuiltinFlag{
        self.FP.Writeln("lns_init_lns_builtin_Sub( _pEnv );")
    }
    if Lns_op_not(self.canConv){
        self.FP.Writeln(Lns_getVM().String_format("lns_init_%s_Sub( _pEnv );", []LnsAny{moduleFullName}))
    } else { 
        for _, _declAlgeNode := range( node.FP.Get_nodeManager().FP.GetDeclAlgeNodeList().Items ) {
            declAlgeNode := _declAlgeNode.(Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
            convCC_filter_1642_(&declAlgeNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _child := range( node.FP.Get_children().Items ) {
            child := _child.(Nodes_NodeDownCast).ToNodes_Node()
            if _switch6017 := child.FP.Get_kind(); _switch6017 == Nodes_NodeKind_get_DeclAlge() || _switch6017 == Nodes_NodeKind_get_DeclFunc() || _switch6017 == Nodes_NodeKind_get_DeclMacro() || _switch6017 == Nodes_NodeKind_get_TestCase() {
            } else {
                convCC_filter_1642_(child, self, &node.Nodes_Node)
                self.FP.Writeln("")
            }
        }
    }
    self.FP.Writeln("lns_leave_block( _pEnv );")
    self.FP.Writeln(Lns_getVM().String_format("return &%s;", []LnsAny{moduleInfoName}))
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 1519: decl @lune.@base.@convCC.convFilter.processBuiltin
func (self *convCC_convFilter) processBuiltin() *Nodes_NodeManager {
    var nodeManager *Nodes_NodeManager
    nodeManager = NewNodes_NodeManager()
    var dummyPos *Types_Position
    dummyPos = NewTypes_Position(0, 0, "builtin")
    var createNodeFromSymbol func(classInfo LnsAny,symbol *Ast_SymbolInfo) LnsAny
    createNodeFromSymbol = func(classInfo LnsAny,symbol *Ast_SymbolInfo) LnsAny {
        var token *Types_Token
        token = NewTypes_Token(Types_TokenKind__Symb, symbol.FP.Get_name(), dummyPos, false, nil)
        if _switch6812 := symbol.FP.Get_kind(); _switch6812 == Ast_SymbolKind__Mtd || _switch6812 == Ast_SymbolKind__Fun {
            var argList *LnsList
            argList = NewLnsList([]LnsAny{})
            for _index, _argType := range( symbol.FP.Get_typeInfo().FP.Get_argTypeInfoList().Items ) {
                index := _index + 1
                argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var argToken *Types_Token
                argToken = NewTypes_Token(Types_TokenKind__Symb, Lns_getVM().String_format("arg%d", []LnsAny{index}), dummyPos, false, nil)
                var dummyScope *Ast_Scope
                dummyScope = NewAst_Scope(self.processInfo, nil, false, nil, nil)
                var argSym *convCC_BuiltinArgSymbolInfo
                argSym = NewconvCC_BuiltinArgSymbolInfo(dummyScope, argToken.Txt, argType, nil, symbol.FP.Get_typeInfo())
                argList.Insert(Nodes_DeclArgNode2Stem(Nodes_DeclArgNode_create(nodeManager, dummyPos, false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(argType)}), argToken, &argSym.Ast_SymbolInfo, nil)))
            }
            if classInfo != nil{
                classInfo_6407 := classInfo.(*Ast_TypeInfo)
                var declFuncInfo *Nodes_DeclFuncInfo
                declFuncInfo = NewNodes_DeclFuncInfo(Nodes_FuncKind__Mtd, classInfo_6407, nil, token, symbol, argList, false, Ast_AccessMode__Pub, nil, symbol.FP.Get_typeInfo().FP.Get_retTypeInfoList(), false, false)
                return &Nodes_DeclMethodNode_create(nodeManager, dummyPos, false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(symbol.FP.Get_typeInfo())}), declFuncInfo).Nodes_Node
            } else {
                var declFuncInfo *Nodes_DeclFuncInfo
                declFuncInfo = NewNodes_DeclFuncInfo(Nodes_FuncKind__Func, nil, nil, token, symbol, argList, false, Ast_AccessMode__Pub, nil, symbol.FP.Get_typeInfo().FP.Get_retTypeInfoList(), false, false)
                return &Nodes_DeclFuncNode_create(nodeManager, dummyPos, false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(symbol.FP.Get_typeInfo())}), declFuncInfo).Nodes_Node
            }
        } else if _switch6812 == Ast_SymbolKind__Var {
            var varToken *Types_Token
            varToken = NewTypes_Token(Types_TokenKind__Symb, symbol.FP.Get_name(), dummyPos, false, nil)
            return &Nodes_DeclVarNode_create(nodeManager, dummyPos, false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(symbol.FP.Get_typeInfo())}), Nodes_DeclVarMode__Let, Ast_AccessMode__Pub, true, NewLnsList([]LnsAny{Nodes_VarInfo2Stem(NewNodes_VarInfo(varToken, nil, symbol.FP.Get_typeInfo()))}), nil, NewLnsList([]LnsAny{Ast_SymbolInfo2Stem(symbol)}), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(symbol.FP.Get_typeInfo())}), false, nil, nil, NewLnsList([]LnsAny{}), nil).Nodes_Node
        } else if _switch6812 == Ast_SymbolKind__Mbr {
            return nil
        } else {
            Util_err(Lns_getVM().String_format("illegal kind -- %s", []LnsAny{Ast_SymbolKind_getTxt( symbol.FP.Get_kind())}))
        }
    // insert a dummy
        return nil
    }
    var builtin *TransUnit_BuiltinFuncType
    builtin = TransUnit_getBuiltinFunc()
    for _, _classInfo := range( builtin.FP.Get_allClass().Items ) {
        classInfo := _classInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if _switch6995 := classInfo.FP.Get_kind(); _switch6995 == Ast_TypeInfoKind__List || _switch6995 == Ast_TypeInfoKind__Array || _switch6995 == Ast_TypeInfoKind__Set || _switch6995 == Ast_TypeInfoKind__Map || _switch6995 == Ast_TypeInfoKind__Box {
        } else {
            if classInfo != Ast_builtinTypeString{
                Lns_print([]LnsAny{classInfo.FP.GetTxt(nil, nil, nil)})
                var classScope *Ast_Scope
                classScope = Lns_unwrap( classInfo.FP.Get_scope()).(*Ast_Scope)
                var fieldList *LnsList
                fieldList = NewLnsList([]LnsAny{})
                Nodes_DeclClassNode_create(nodeManager, dummyPos, false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(classInfo)}), Ast_AccessMode__Pub, NewTypes_Token(Types_TokenKind__Symb, classInfo.FP.Get_rawTxt(), dummyPos, false, nil), false, classInfo.FP.Get_rawTxt(), nil, Nodes_LazyLoad__Off, false, fieldList, NewLnsList([]LnsAny{}), fieldList, NewLnsList([]LnsAny{}), classScope, NewNodes_ClassInitBlockInfo(nil), NewLnsList([]LnsAny{}), NewLnsList([]LnsAny{}), NewLnsList([]LnsAny{}), NewLnsSet([]LnsAny{}))
                {
                    __collection6990 := classScope.FP.Get_symbol2SymbolInfoMap()
                    __sorted6990 := __collection6990.CreateKeyListStr()
                    __sorted6990.Sort( LnsItemKindStr, nil )
                    for _, ___key6990 := range( __sorted6990.Items ) {
                        field := __collection6990.Items[ ___key6990 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                        {
                            _node := createNodeFromSymbol(classInfo, field)
                            if _node != nil {
                                node := _node.(*Nodes_Node)
                                fieldList.Insert(Nodes_Node2Stem(node))
                            }
                        }
                    }
                }
            }
        }
    }
    for _, _symbol := range( builtin.FP.Get_allSymbol().Items ) {
        symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symbol.FP.Get_kind() == Ast_SymbolKind__Var) ||
            Lns_GetEnv().SetStackVal( symbol.FP.Get_kind() == Ast_SymbolKind__Fun) ||
            Lns_GetEnv().SetStackVal( symbol.FP.Get_namespaceTypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Root) ).(bool){
            createNodeFromSymbol(nil, symbol)
        }
    }
    return nodeManager
}

// 1617: decl @lune.@base.@convCC.convFilter.processRoot
func (self *convCC_convFilter) ProcessRoot(node *Nodes_RootNode,_opt LnsAny) {
    var nodeManager *Nodes_NodeManager
    convCC_registerBuiltin_1761_()
    if self.outputBuiltinFlag{
        nodeManager = self.FP.processBuiltin()
        
    } else { 
        nodeManager = node.FP.Get_nodeManager()
        
    }
    self.scopeMgr.FP.Setup(self.ast.FP.Get_moduleScope(), node.FP.Get_nodeManager().FP.GetDeclMemberNodeList())
    for _pragma := range( node.FP.Get_luneHelperInfo().PragmaSet.Items ) {
        pragma := _pragma
        switch _exp7127 := pragma.(type) {
        case *LuneControl_Pragma__limit_conv_code:
        codeSet := _exp7127.Val1
            if Lns_op_not(codeSet.Has(LuneControl_Code__C)){
                self.canConv = false
                
                break
            }
        }
    }
    self.stream.FP.SwitchToHeader()
    var ifdefname string
    ifdefname = convCC_convExp7157(Lns_2DDD(Lns_getVM().String_gsub(self.moduleCtrl.FP.GetFilePath(self.moduleTypeInfo),"/", "_")))
    self.FP.Writeln(Lns_getVM().String_format("#ifndef __%s__\n#define __%s__\n       ", []LnsAny{ifdefname, ifdefname}))
    self.stream.FP.ReturnToSource()
    self.FP.Writeln(Lns_getVM().String_format("// %s", []LnsAny{self.streamName}))
    self.FP.Writeln("#include <lunescript.h>")
    self.FP.Writeln(Lns_getVM().String_format("#include <%s.h>", []LnsAny{self.moduleCtrl.FP.GetFilePath(node.FP.Get_moduleTypeInfo())}))
    self.processMode = convCC_ProcessMode__Include
    
    for _, _importNode := range( nodeManager.FP.GetImportNodeList().Items ) {
        importNode := _importNode.(Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        convCC_filter_1642_(&importNode.Nodes_Node, self, &node.Nodes_Node)
    }
    self.processMode = convCC_ProcessMode__Prototype
    
    for _, _workNode := range( nodeManager.FP.GetTestCaseNodeList().Items ) {
        workNode := _workNode.(Nodes_TestCaseNodeDownCast).ToNodes_TestCaseNode()
        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
    }
    var moduleName string
    moduleName = self.moduleCtrl.FP.GetFullName(node.FP.Get_moduleTypeInfo())
    self.FP.Write(Lns_getVM().String_format("static %s lns_moduleInfo_%s = {NULL,NULL,false, NULL, \"%s\", {", []LnsAny{convCC_cTypeMod, moduleName, moduleName}))
    for _, _workNode := range( nodeManager.FP.GetTestCaseNodeList().Items ) {
        workNode := _workNode.(Nodes_TestCaseNodeDownCast).ToNodes_TestCaseNode()
        self.FP.Write(Lns_getVM().String_format("%s__test_%s, ", []LnsAny{moduleName, workNode.FP.Get_name().Txt}))
    }
    self.FP.Writeln("NULL } };")
    self.FP.Writeln(Lns_getVM().String_format("static %s lns_module_globalStemList;", []LnsAny{convCC_cTypeAnyPP}))
    self.FP.Writeln(Lns_getVM().String_format("static %s lns_module_path = NULL;", []LnsAny{convCC_cTypeAnyPP}))
    var process func(onlyPub bool)
    process = func(onlyPub bool) {
        self.processMode = convCC_ProcessMode__Prototype
        
        for _, _workNode := range( nodeManager.FP.GetDeclEnumNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclEnumNodeDownCast).ToNodes_DeclEnumNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclFormNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclFormNodeDownCast).ToNodes_DeclFormNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclFuncNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclAlgeNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclClassNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclConstrNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclConstrNodeDownCast).ToNodes_DeclConstrNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclMethodNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclMethodNodeDownCast).ToNodes_DeclMethodNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetProtoMethodNodeList().Items ) {
            workNode := _workNode.(Nodes_ProtoMethodNodeDownCast).ToNodes_ProtoMethodNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        if self.canConv{
            for _, _workNode := range( nodeManager.FP.GetExpToDDDNodeList().Items ) {
                workNode := _workNode.(Nodes_ExpToDDDNodeDownCast).ToNodes_ExpToDDDNode()
                if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                    if onlyPub{
                        if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                            convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                        }
                    } else { 
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                }
            }
            
            for _, _workNode := range( nodeManager.FP.GetLiteralStringNodeList().Items ) {
                workNode := _workNode.(Nodes_LiteralStringNodeDownCast).ToNodes_LiteralStringNode()
                if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                    if onlyPub{
                        if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                            convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                        }
                    } else { 
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                }
            }
            
            for _, _workNode := range( nodeManager.FP.GetExpCastNodeList().Items ) {
                workNode := _workNode.(Nodes_ExpCastNodeDownCast).ToNodes_ExpCastNode()
                if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                    if onlyPub{
                        if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                            convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                        }
                    } else { 
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                }
            }
            
        }
        self.processMode = convCC_ProcessMode__WideScopeVer
        
        for _, _decl := range( nodeManager.FP.GetDeclVarNodeList().Items ) {
            decl := _decl.(Nodes_DeclVarNodeDownCast).ToNodes_DeclVarNode()
            convCC_filter_1642_(&decl.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _workNode := range( nodeManager.FP.GetDeclAlgeNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclClassNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclConstrNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclConstrNodeDownCast).ToNodes_DeclConstrNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclMethodNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclMethodNodeDownCast).ToNodes_DeclMethodNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        for _, _workNode := range( nodeManager.FP.GetDeclFuncNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
        self.processMode = convCC_ProcessMode__DefClass
        
        for _, _workNode := range( nodeManager.FP.GetDeclClassNodeList().Items ) {
            workNode := _workNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            if Lns_op_not(workNode.FP.Get_macroArgFlag()){
                if onlyPub{
                    if Ast_isPubToExternal(workNode.FP.Get_expType().FP.Get_accessMode()){
                        convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                    }
                } else { 
                    convCC_filter_1642_(&workNode.Nodes_Node, self, &node.Nodes_Node)
                }
            }
        }
        
    }
    process(Lns_op_not(self.canConv))
    if self.canConv{
        self.processMode = convCC_ProcessMode__StringFormat
        
        for _, _litStr := range( nodeManager.FP.GetLiteralStringNodeList().Items ) {
            litStr := _litStr.(Nodes_LiteralStringNodeDownCast).ToNodes_LiteralStringNode()
            if Lns_op_not(litStr.FP.Get_macroArgFlag()){
                convCC_filter_1642_(&litStr.Nodes_Node, self, &node.Nodes_Node)
            }
        }
        self.processMode = convCC_ProcessMode__Immediate
        
        self.processedNodeSet = NewLnsSet([]LnsAny{})
        
        var procssLiteralCtor func(literalNodeList *LnsList)
        procssLiteralCtor = func(literalNodeList *LnsList) {
            for _, _literalNode := range( literalNodeList.Items ) {
                literalNode := _literalNode.(Nodes_NodeDownCast).ToNodes_Node()
                self.processingNode = literalNode
                
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( Lns_op_not(self.processedNodeSet.Has(Nodes_Node2Stem(literalNode)))) &&
                    Lns_GetEnv().SetStackVal( Lns_op_not(literalNode.FP.Get_macroArgFlag())) ).(bool)){
                    self.accessSymbolSet = NewUtil_OrderedSet()
                    
                    convCC_filter_1642_(literalNode, self, &node.Nodes_Node)
                    self.processedNodeSet.Add(Nodes_RootNode2Stem(node))
                }
            }
        }
        procssLiteralCtor(nodeManager.FP.GetLiteralListNodeList())
        procssLiteralCtor(nodeManager.FP.GetLiteralArrayNodeList())
        procssLiteralCtor(nodeManager.FP.GetLiteralSetNodeList())
        procssLiteralCtor(nodeManager.FP.GetLiteralMapNodeList())
        self.processingNode = nil
        
        self.processMode = convCC_ProcessMode__Intermediate
        
        for _, _callNode := range( nodeManager.FP.GetExpCallNodeList().Items ) {
            callNode := _callNode.(Nodes_ExpCallNodeDownCast).ToNodes_ExpCallNode()
            convCC_filter_1642_(&callNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _dddNode := range( nodeManager.FP.GetExpToDDDNodeList().Items ) {
            dddNode := _dddNode.(Nodes_ExpToDDDNodeDownCast).ToNodes_ExpToDDDNode()
            convCC_filter_1642_(&dddNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _castNode := range( nodeManager.FP.GetExpCastNodeList().Items ) {
            castNode := _castNode.(Nodes_ExpCastNodeDownCast).ToNodes_ExpCastNode()
            convCC_filter_1642_(&castNode.Nodes_Node, self, &node.Nodes_Node)
        }
        self.processMode = convCC_ProcessMode__DefWrap
        
        for _, _callNode := range( nodeManager.FP.GetExpCallNodeList().Items ) {
            callNode := _callNode.(Nodes_ExpCallNodeDownCast).ToNodes_ExpCallNode()
            convCC_filter_1642_(&callNode.Nodes_Node, self, &node.Nodes_Node)
        }
        self.processMode = convCC_ProcessMode__InitFuncSym
        
        self.FP.Writeln(Lns_getVM().String_format("static void initFuncSym( %s _pEnv, %s pBlock )\n{", []LnsAny{convCC_cTypeEnvP, convCC_cTypeBlockP}))
        self.FP.PushIndent(nil)
        for _, _declConstrNode := range( nodeManager.FP.GetDeclConstrNodeList().Items ) {
            declConstrNode := _declConstrNode.(Nodes_DeclConstrNodeDownCast).ToNodes_DeclConstrNode()
            convCC_filter_1642_(&declConstrNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _declMethodNode := range( nodeManager.FP.GetDeclMethodNodeList().Items ) {
            declMethodNode := _declMethodNode.(Nodes_DeclMethodNodeDownCast).ToNodes_DeclMethodNode()
            convCC_filter_1642_(&declMethodNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _declFuncNode := range( nodeManager.FP.GetDeclFuncNodeList().Items ) {
            declFuncNode := _declFuncNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            convCC_filter_1642_(&declFuncNode.Nodes_Node, self, &node.Nodes_Node)
        }
        self.FP.PopIndent()
        self.FP.Writeln("}")
        self.processMode = convCC_ProcessMode__Form
        
        for _, _declEnumNode := range( nodeManager.FP.GetDeclEnumNodeList().Items ) {
            declEnumNode := _declEnumNode.(Nodes_DeclEnumNodeDownCast).ToNodes_DeclEnumNode()
            convCC_filter_1642_(&declEnumNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _declAlgeNode := range( nodeManager.FP.GetDeclAlgeNodeList().Items ) {
            declAlgeNode := _declAlgeNode.(Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
            convCC_filter_1642_(&declAlgeNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _declConstrNode := range( nodeManager.FP.GetDeclConstrNodeList().Items ) {
            declConstrNode := _declConstrNode.(Nodes_DeclConstrNodeDownCast).ToNodes_DeclConstrNode()
            convCC_filter_1642_(&declConstrNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _declMethodNode := range( nodeManager.FP.GetDeclMethodNodeList().Items ) {
            declMethodNode := _declMethodNode.(Nodes_DeclMethodNodeDownCast).ToNodes_DeclMethodNode()
            convCC_filter_1642_(&declMethodNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _declMethodNode := range( nodeManager.FP.GetProtoMethodNodeList().Items ) {
            declMethodNode := _declMethodNode.(Nodes_ProtoMethodNodeDownCast).ToNodes_ProtoMethodNode()
            convCC_filter_1642_(&declMethodNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _declFormNode := range( nodeManager.FP.GetDeclFormNodeList().Items ) {
            declFormNode := _declFormNode.(Nodes_DeclFormNodeDownCast).ToNodes_DeclFormNode()
            convCC_filter_1642_(&declFormNode.Nodes_Node, self, &node.Nodes_Node)
        }
        for _, _declFuncNode := range( nodeManager.FP.GetDeclFuncNodeList().Items ) {
            declFuncNode := _declFuncNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            self.duringDeclFunc = false
            
            convCC_filter_1642_(&declFuncNode.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.processInitModule(node)
    for _, _testBlock := range( nodeManager.FP.GetTestCaseNodeList().Items ) {
        testBlock := _testBlock.(Nodes_TestCaseNodeDownCast).ToNodes_TestCaseNode()
        convCC_filter_1642_(&testBlock.Nodes_Node, self, &node.Nodes_Node)
    }
    if self.outputBuiltinFlag{
        self.FP.Writeln("#include \"lns_builtinInc.c\"")
    }
    self.stream.FP.SwitchToHeader()
    self.FP.Writeln("#endif")
    self.stream.FP.ReturnToSource()
}

// 1862: decl @lune.@base.@convCC.convFilter.processSubfile
func (self *convCC_convFilter) ProcessSubfile(node *Nodes_SubfileNode,_opt LnsAny) {
}

// 1935: decl @lune.@base.@convCC.convFilter.processBlockPreProcess
func (self *convCC_convFilter) processBlockPreProcess(scope *Ast_Scope) {
    self.FP.PushIndent(nil)
    var anyNum LnsInt
    var stemNum LnsInt
    var varNum LnsInt
    anyNum,stemNum,varNum = self.scopeMgr.FP.SetupScopeParam(scope)
    self.FP.Writeln(Lns_getVM().String_format("lns_block_t * %s = lns_enter_block( _pEnv, %d, %d, %d );", []LnsAny{convCC_getBlockName_1198_(scope), anyNum, stemNum, varNum}))
    self.routineInfoStack.FP.PushDepth()
    self.loopInfoStack.FP.PushDepth()
}

// 1944: decl @lune.@base.@convCC.convFilter.processBlockPostProcess
func (self *convCC_convFilter) processBlockPostProcess() {
    self.loopInfoStack.FP.PopDepth()
    self.routineInfoStack.FP.PopDepth()
    self.FP.Writeln("lns_leave_block( _pEnv );")
    self.FP.PopIndent()
}

// 1951: decl @lune.@base.@convCC.convFilter.pushRoutine
func (self *convCC_convFilter) pushRoutine(funcType *Ast_TypeInfo,blockNode *Nodes_BlockNode) {
    self.FP.processBlockPreProcess(blockNode.FP.Get_scope())
    self.routineInfoStack.FP.NewInfo(&NewconvCC_RoutineInfo(funcType).convCC_DepthInfo)
}

// 1956: decl @lune.@base.@convCC.convFilter.popRoutine
func (self *convCC_convFilter) popRoutine() {
    self.routineInfoStack.FP.DelInfo()
    self.FP.processBlockPostProcess()
}

// 1961: decl @lune.@base.@convCC.convFilter.processLoopPreProcess
func (self *convCC_convFilter) processLoopPreProcess(blockNode *Nodes_BlockNode) {
    self.FP.processBlockPreProcess(blockNode.FP.Get_scope())
    self.loopInfoStack.FP.NewInfo(NewconvCC_DepthInfo())
}

// 1966: decl @lune.@base.@convCC.convFilter.processLoopPostProcess
func (self *convCC_convFilter) processLoopPostProcess() {
    self.loopInfoStack.FP.DelInfo()
    self.FP.processBlockPostProcess()
}

// 1971: decl @lune.@base.@convCC.convFilter.processBlockSub
func (self *convCC_convFilter) ProcessBlockSub(node *Nodes_BlockNode,_opt LnsAny) {
    self.scopeMgr.FP.SetupScopeParam(node.FP.Get_scope())
    var scope *Ast_Scope
    scope = node.FP.Get_scope()
    {
        __collection9824 := scope.FP.Get_closureSymMap()
        __sorted9824 := __collection9824.CreateKeyListInt()
        __sorted9824.Sort( LnsItemKindInt, nil )
        for _, ___key9824 := range( __sorted9824.Items ) {
            symbol := __collection9824.Items[ ___key9824 ].(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            var typeTxt string
            typeTxt = convCC_convExp9781(Lns_2DDD(self.scopeMgr.FP.GetCTypeForSym(symbol.FP)))
            self.FP.Write(Lns_getVM().String_format("%s %s = l_form_closure_var( _pForm, %d )", []LnsAny{typeTxt, self.moduleCtrl.FP.GetSymbolName(symbol.FP), Lns_unwrap( scope.FP.Get_closureSym2NumMap().Items[symbol]).(LnsInt)}))
            self.FP.Writeln(";")
        }
    }
    var loopFlag bool
    loopFlag = false
    var readyBlock bool
    readyBlock = false
    var word string
    word = ""
    if _switch10066 := node.FP.Get_blockKind(); _switch10066 == Nodes_BlockKind__If || _switch10066 == Nodes_BlockKind__Elseif {
        word = "{"
        
    } else if _switch10066 == Nodes_BlockKind__Else {
        word = ""
        
    } else if _switch10066 == Nodes_BlockKind__While {
        loopFlag = true
        
    } else if _switch10066 == Nodes_BlockKind__Repeat {
        word = ""
        
        loopFlag = true
        
    } else if _switch10066 == Nodes_BlockKind__For {
        word = ""
        
        loopFlag = true
        
    } else if _switch10066 == Nodes_BlockKind__Apply {
        word = ""
        
        loopFlag = true
        
    } else if _switch10066 == Nodes_BlockKind__Foreach {
        word = ""
        
        loopFlag = true
        
    } else if _switch10066 == Nodes_BlockKind__Macro {
        word = ""
        
    } else if _switch10066 == Nodes_BlockKind__Func {
        readyBlock = true
        
        word = ""
        
    } else if _switch10066 == Nodes_BlockKind__Default {
        word = ""
        
    } else if _switch10066 == Nodes_BlockKind__Block {
        word = "{"
        
    } else if _switch10066 == Nodes_BlockKind__LetUnwrap {
        readyBlock = true
        
        word = ""
        
    } else if _switch10066 == Nodes_BlockKind__LetUnwrapThenDo {
        word = ""
        
    } else if _switch10066 == Nodes_BlockKind__IfUnwrap {
        readyBlock = true
        
        word = ""
        
    } else if _switch10066 == Nodes_BlockKind__When {
        readyBlock = true
        
        word = ""
        
    }
    if loopFlag{
        readyBlock = true
        
    }
    self.FP.Writeln(Lns_getVM().String_format("%s // %d", []LnsAny{word, node.FP.Get_pos().LineNo}))
    if Lns_op_not(readyBlock){
        self.FP.processBlockPreProcess(node.FP.Get_scope())
    }
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList()
    for _, _statement := range( stmtList.Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        convCC_filter_1642_(statement, self, &node.Nodes_Node)
        self.FP.Writeln("")
    }
    if Lns_op_not(readyBlock){
        self.FP.processBlockPostProcess()
    }
    if node.FP.Get_blockKind() == Nodes_BlockKind__Block{
        self.FP.Writeln("}")
    }
}

// 2085: decl @lune.@base.@convCC.convFilter.processStmtExp
func (self *convCC_convFilter) ProcessStmtExp(node *Nodes_StmtExpNode,_opt LnsAny) {
    convCC_filter_1642_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Write(Lns_getVM().String_format("; // %d", []LnsAny{node.FP.Get_pos().LineNo}))
}

// 2180: decl @lune.@base.@convCC.convFilter.processSym2stem
func (self *convCC_convFilter) processSym2stem(symbolInfo Ast_LowSymbol) {
    var valKind LnsInt
    valKind = self.scopeMgr.FP.GetSymbolValKind(symbolInfo)
    if _switch10713 := valKind; _switch10713 == convCC_ValKind__Any {
        self.FP.Write("LNS_STEM_ANY( ")
        if symbolInfo.Get_kind() == Ast_SymbolKind__Var{
            self.FP.Write("*")
        }
        self.FP.Write(self.moduleCtrl.FP.GetSymbolName(symbolInfo))
        self.FP.Write(")")
        return 
    } else if _switch10713 == convCC_ValKind__Var {
        self.FP.Write(self.moduleCtrl.FP.GetSymbolName(symbolInfo))
        self.FP.Write("->stem")
        return 
    } else if _switch10713 == convCC_ValKind__Stem {
        self.FP.Write(self.moduleCtrl.FP.GetSymbolName(symbolInfo))
        return 
    }
    var expType *Ast_TypeInfo
    expType = symbolInfo.Get_typeInfo().FP.Get_srcTypeInfo()
    {
        _enumType := Ast_EnumTypeInfoDownCastF(expType.FP)
        if _enumType != nil {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            expType = enumType.FP.Get_valTypeInfo()
            
        }
    }
    if _switch10963 := expType; _switch10963 == Ast_builtinTypeInt || _switch10963 == Ast_builtinTypeChar {
        self.FP.Write("LNS_STEM_INT( ")
        self.FP.Write("")
        self.FP.Write(self.scopeMgr.FP.GetAccessPrimValFromSymbol(symbolInfo))
        self.FP.Write(")")
    } else if _switch10963 == Ast_builtinTypeReal {
        self.FP.Write("LNS_STEM_REAL( ")
        self.FP.Write(self.scopeMgr.FP.GetAccessPrimValFromSymbol(symbolInfo))
        self.FP.Write(")")
    } else if _switch10963 == Ast_builtinTypeBool {
        self.FP.Write("LNS_STEM_BOOL( ")
        self.FP.Write(self.scopeMgr.FP.GetAccessPrimValFromSymbol(symbolInfo))
        self.FP.Write(")")
    } else if _switch10963 == Ast_builtinTypeStem || _switch10963 == Ast_builtinTypeStem_ {
        self.FP.Write(self.moduleCtrl.FP.GetSymbolName(symbolInfo))
    } else {
        if _switch10961 := expType.FP.Get_kind(); _switch10961 == Ast_TypeInfoKind__DDD {
            self.FP.Write("_pDDD")
        } else if _switch10961 == Ast_TypeInfoKind__Func {
            if Lns_isCondTrue( expType.FP.Get_scope()){
                self.FP.Write("LNS_STEM_ANY(")
                self.FP.Write(convCC_getFunc2any_1897_(self.moduleCtrl, self.scopeMgr, expType))
                self.FP.Write(")")
            } else { 
                Util_err("illegal func")
            }
        } else {
            self.FP.Write(self.moduleCtrl.FP.GetSymbolName(symbolInfo))
        }
    }
}

// 2252: decl @lune.@base.@convCC.convFilter.processDeclEnum
func (self *convCC_convFilter) ProcessDeclEnum(node *Nodes_DeclEnumNode,_opt LnsAny) {
    var enumType *Ast_EnumTypeInfo
    enumType = Lns_unwrap( Ast_EnumTypeInfoDownCastF(node.FP.Get_expType().FP)).(*Ast_EnumTypeInfo)
    var enumFullName string
    enumFullName = self.moduleCtrl.FP.GetEnumTypeName(&enumType.Ast_TypeInfo)
    var fullName string
    fullName = self.FP.getFullName(&enumType.Ast_TypeInfo)
    var isStrEnum bool
    isStrEnum = enumType.FP.Get_valTypeInfo().FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil)
    if _switch12229 := self.processMode; _switch12229 == convCC_ProcessMode__Prototype {
        var process func(out2HMode LnsInt)
        process = func(out2HMode LnsInt) {
            var prefix string
            prefix = convCC_getOut2HeaderPrefix_1518_(out2HMode)
            for _, _valName := range( node.FP.Get_valueNameList().Items ) {
                valName := _valName.(Types_TokenDownCast).ToTypes_Token()
                var valInfo *Ast_EnumValInfo
                valInfo = Lns_unwrap( enumType.FP.GetEnumValInfo(valName.Txt)).(*Ast_EnumValInfo)
                var enumValName string
                enumValName = self.moduleCtrl.FP.GetEnumValCName(&enumType.Ast_TypeInfo, valName.Txt)
                if isStrEnum{
                    self.FP.Writeln(Lns_getVM().String_format("%s%s %s;", []LnsAny{prefix, convCC_cTypeAnyP, enumValName}))
                } else { 
                    if _switch11137 := out2HMode; _switch11137 == convCC_Out2HMode__HeaderPub || _switch11137 == convCC_Out2HMode__SourcePri {
                        var valTxt string
                        valTxt = Lns_getVM().String_format("%s", []LnsAny{Ast_getEnumLiteralVal(valInfo.FP.Get_val())})
                        self.FP.Writeln(Lns_getVM().String_format("#define %s %s", []LnsAny{enumValName, valTxt}))
                    }
                }
            }
            if _switch11231 := out2HMode; _switch11231 == convCC_Out2HMode__HeaderPub || _switch11231 == convCC_Out2HMode__SourcePri {
                self.FP.Writeln(Lns_getVM().String_format("%s%s %s_get__allList( lns_env_t * _pEnv );", []LnsAny{prefix, convCC_cTypeAnyP, enumFullName}))
                self.FP.Writeln(Lns_getVM().String_format("%s%s %s_get__txt( %s _pEnv, %s val );", []LnsAny{prefix, convCC_cTypeAnyP, enumFullName, convCC_cTypeEnvP, convCC_getCType_1188_(enumType.FP.Get_valTypeInfo())}))
                self.FP.Writeln(Lns_getVM().String_format("%s%s %s( %s _pEnv, %s val );", []LnsAny{prefix, convCC_cTypeStem, self.moduleCtrl.FP.GetEnumFuncName(enumType, "_from"), convCC_cTypeEnvP, convCC_getCType_1188_(enumType.FP.Get_valTypeInfo())}))
            }
        }
        {
            var processwork func(out2HMode LnsInt)
            processwork = func(out2HMode LnsInt) {
                process(out2HMode)
            }
            if Ast_isPubToExternal(enumType.FP.Get_accessMode()){
                self.stream.FP.SwitchToHeader()
                processwork(convCC_Out2HMode__HeaderPub)
                self.stream.FP.ReturnToSource()
                processwork(convCC_Out2HMode__SourcePub)
            } else { 
                processwork(convCC_Out2HMode__SourcePri)
            }
        }
        
        self.FP.Writeln(Lns_getVM().String_format("static %s %s_val2NameMap;", []LnsAny{convCC_cTypeAnyP, enumFullName}))
        self.FP.Writeln(Lns_getVM().String_format("static %s %s_allList;", []LnsAny{convCC_cTypeAnyP, enumFullName}))
    } else if _switch12229 == convCC_ProcessMode__Form {
        if Lns_op_not(Ast_isPubToExternal(enumType.FP.Get_accessMode())){
            self.FP.Write("static ")
        }
        self.FP.Writeln(Lns_getVM().String_format("%s %s_get__allList( lns_env_t * _pEnv )", []LnsAny{convCC_cTypeAnyP, enumFullName}))
        self.FP.Writeln("{")
        self.FP.Writeln(Lns_getVM().String_format("    return %s_allList;", []LnsAny{enumFullName}))
        self.FP.Writeln("}")
        {
            if Lns_op_not(Ast_isPubToExternal(enumType.FP.Get_accessMode())){
                self.FP.Write("static ")
            }
            self.FP.Writeln(Lns_getVM().String_format("%s %s_get__txt( %s _pEnv, %s val ) {", []LnsAny{convCC_cTypeAnyP, enumFullName, convCC_cTypeEnvP, convCC_getCType_1188_(enumType.FP.Get_valTypeInfo())}))
            self.FP.PushIndent(nil)
            self.FP.Write(Lns_getVM().String_format("%s _work =  lns_mtd_Map_get( _pEnv, %s, ", []LnsAny{convCC_cTypeStem, self.moduleCtrl.FP.GetEnumVal2NameMapName(&enumType.Ast_TypeInfo)}))
            var workSym *convCC_WorkSymbol
            workSym = NewconvCC_WorkSymbol(Lns_unwrap( self.moduleTypeInfo.FP.Get_scope()).(*Ast_Scope), Ast_AccessMode__Local, "val", enumType.FP.Get_valTypeInfo(), Ast_SymbolKind__Arg, false, NewconvCC_SymbolParam(convCC_getValKind_1167_(enumType.FP.Get_valTypeInfo()), 0, convCC_getCType_1188_(enumType.FP.Get_valTypeInfo())))
            self.FP.processSym2stem(workSym.FP)
            self.FP.Writeln(");")
            self.FP.Writeln(Lns_getVM().String_format("return _work%s;", []LnsAny{convCC_accessAny}))
            self.FP.PopIndent()
            self.FP.Writeln("}")
        }
        {
            if Lns_op_not(Ast_isPubToExternal(enumType.FP.Get_accessMode())){
                self.FP.Write("static ")
            }
            self.FP.Writeln(Lns_getVM().String_format("%s %s( %s _pEnv, %s val ) {", []LnsAny{convCC_cTypeStem, self.moduleCtrl.FP.GetEnumFuncName(enumType, "_from"), convCC_cTypeEnvP, convCC_getCType_1188_(enumType.FP.Get_valTypeInfo())}))
            self.FP.PushIndent(nil)
            self.FP.Write(Lns_getVM().String_format("%s key = ", []LnsAny{convCC_cTypeStem}))
            var workSym *convCC_WorkSymbol
            workSym = NewconvCC_WorkSymbol(Lns_unwrap( self.moduleTypeInfo.FP.Get_scope()).(*Ast_Scope), Ast_AccessMode__Local, "val", enumType.FP.Get_valTypeInfo(), Ast_SymbolKind__Arg, false, NewconvCC_SymbolParam(convCC_getValKind_1167_(enumType.FP.Get_valTypeInfo()), 0, convCC_getCType_1188_(enumType.FP.Get_valTypeInfo())))
            self.FP.processSym2stem(workSym.FP)
            self.FP.Writeln(";")
            self.FP.Writeln(Lns_getVM().String_format("%s _work = lns_mtd_Map_get( _pEnv, %s, key );", []LnsAny{convCC_cTypeStem, self.moduleCtrl.FP.GetEnumVal2NameMapName(&enumType.Ast_TypeInfo)}))
            self.FP.Writeln("if ( _work.type == lns_stem_type_nil ) {")
            self.FP.PushIndent(nil)
            self.FP.Writeln("return lns_global.nilStem;")
            self.FP.PopIndent()
            self.FP.Writeln("}")
            self.FP.Writeln("return key;")
            self.FP.PopIndent()
            self.FP.Writeln("}")
        }
        self.FP.Writeln(Lns_getVM().String_format("static void init_%s( lns_env_t * _pEnv )", []LnsAny{enumFullName}))
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        var anyVarList *LnsList
        anyVarList = NewLnsList([]LnsAny{})
        if isStrEnum{
            for _, _valName := range( node.FP.Get_valueNameList().Items ) {
                valName := _valName.(Types_TokenDownCast).ToTypes_Token()
                var valInfo *Ast_EnumValInfo
                valInfo = Lns_unwrap( enumType.FP.GetEnumValInfo(valName.Txt)).(*Ast_EnumValInfo)
                var valTxt string
                valTxt = Lns_getVM().String_format("\"%s\"", []LnsAny{Ast_getEnumLiteralVal(valInfo.FP.Get_val())})
                var anyVar string
                anyVar = self.moduleCtrl.FP.GetEnumValCName(&enumType.Ast_TypeInfo, valName.Txt)
                anyVarList.Insert(Lns_getVM().String_format("LNS_STEM_ANY( %s )", []LnsAny{anyVar}))
                self.FP.Writeln(Lns_getVM().String_format("%s = lns_litStr2any( _pEnv, %s );", []LnsAny{anyVar, valTxt}))
            }
        } else { 
            for _, _valName := range( node.FP.Get_valueNameList().Items ) {
                valName := _valName.(Types_TokenDownCast).ToTypes_Token()
                var valInfo *Ast_EnumValInfo
                valInfo = Lns_unwrap( enumType.FP.GetEnumValInfo(valName.Txt)).(*Ast_EnumValInfo)
                var valTxt string
                valTxt = Lns_getVM().String_format("%s", []LnsAny{Ast_getEnumLiteralVal(valInfo.FP.Get_val())})
                var anyVar string
                anyVar = Lns_getVM().String_format("_%s", []LnsAny{valName.Txt})
                anyVarList.Insert(anyVar)
                self.FP.Write(Lns_getVM().String_format("%s %s = ", []LnsAny{convCC_cTypeStem, anyVar}))
                self.FP.Write(convCC_getLiteral2Stem_1884_(valTxt, enumType.FP.Get_valTypeInfo()))
                self.FP.Writeln(";")
            }
        }
        var allListName string
        allListName = Lns_getVM().String_format("%s_allList", []LnsAny{enumFullName})
        self.FP.Write(allListName)
        self.FP.Writeln(" = lns_class_List_new( _pEnv );")
        convCC_processAddModuleGlobal_1639_(self.stream.FP, Lns_getVM().String_format("LNS_STEM_ANY( %s )", []LnsAny{allListName}))
        for _, _anyVar := range( anyVarList.Items ) {
            anyVar := _anyVar.(string)
            self.FP.Writeln(Lns_getVM().String_format("lns_mtd_List_insert( _pEnv, %s_allList, %s );", []LnsAny{enumFullName, anyVar}))
        }
        var val2NameMapName string
        val2NameMapName = Lns_getVM().String_format("%s_val2NameMap", []LnsAny{enumFullName})
        self.FP.Write(val2NameMapName)
        self.FP.Writeln(" = lns_class_Map_new( _pEnv );")
        convCC_processAddModuleGlobal_1639_(self.stream.FP, Lns_getVM().String_format("LNS_STEM_ANY( %s )", []LnsAny{val2NameMapName}))
        for _index, _anyVar := range( anyVarList.Items ) {
            index := _index + 1
            anyVar := _anyVar.(string)
            self.FP.Writeln(Lns_getVM().String_format("lns_mtd_Map_add( _pEnv, %s_val2NameMap, %s, ", []LnsAny{enumFullName, anyVar}))
            self.FP.Writeln(Lns_getVM().String_format("  LNS_STEM_ANY( lns_litStr2any( _pEnv, \"%s.%s\" ) ) );", []LnsAny{fullName, node.FP.Get_valueNameList().GetAt(index).(Types_TokenDownCast).ToTypes_Token().Txt}))
        }
        self.FP.PopIndent()
        self.FP.Writeln("}")
    } else if _switch12229 == convCC_ProcessMode__InitModule {
        self.FP.Writeln(Lns_getVM().String_format("init_%s( _pEnv );", []LnsAny{enumFullName}))
    }
}

// 2884: decl @lune.@base.@convCC.convFilter.processDeclAlge
func (self *convCC_convFilter) ProcessDeclAlge(node *Nodes_DeclAlgeNode,_opt LnsAny) {
    if _switch13759 := self.processMode; _switch13759 == convCC_ProcessMode__Prototype {
        convCC_processAlgePrototype_1928_(self.stream.FP, self.moduleCtrl, node)
    } else if _switch13759 == convCC_ProcessMode__WideScopeVer {
        convCC_processAlgeWideScope_1944_(self.stream.FP, self.moduleCtrl, node)
    } else if _switch13759 == convCC_ProcessMode__Form {
        convCC_processAlgeForm_1960_(self.stream.FP, self.moduleCtrl, node)
    } else if _switch13759 == convCC_ProcessMode__InitModule {
        self.FP.Writeln(Lns_getVM().String_format("%s( _pEnv );", []LnsAny{self.moduleCtrl.FP.GetAlgeInitCName(node.FP.Get_expType())}))
    }
}

// 2960: decl @lune.@base.@convCC.convFilter.processNewAlgeVal
func (self *convCC_convFilter) ProcessNewAlgeVal(node *Nodes_NewAlgeValNode,_opt LnsAny) {
    var valInfo *Ast_AlgeValInfo
    valInfo = node.FP.Get_valInfo()
    if valInfo.FP.Get_typeList().Len() == 0{
        var valName string
        valName = self.moduleCtrl.FP.GetAlgeValCName(&node.FP.Get_algeTypeInfo().Ast_TypeInfo, valInfo.FP.Get_name())
        self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{valName}))
    } else { 
        self.FP.Write(self.moduleCtrl.FP.GetNewAlgeCName(&node.FP.Get_algeTypeInfo().Ast_TypeInfo, valInfo.FP.Get_name()))
        self.FP.Write("( _pEnv")
        for _, _arg := range( node.FP.Get_paramList().Items ) {
            arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
            self.FP.Write(",")
            convCC_filter_1642_(arg, self, &node.Nodes_Node)
        }
        self.FP.Write(")")
    }
}

// 3011: decl @lune.@base.@convCC.convFilter.outputAlter2MapFunc
func (self *convCC_convFilter) outputAlter2MapFunc(stream Util_SourceStream,alt2Map *LnsMap) {
}

// 3752: decl @lune.@base.@convCC.convFilter.processNewInsance
func (self *convCC_convFilter) processNewInsance(classType *Ast_TypeInfo,callInit bool) {
    var className string
    className = self.moduleCtrl.FP.GetClassCName(classType)
    self.FP.Writeln(Lns_getVM().String_format("lns_class_new_( _pEnv, %s, pAny, pObj );", []LnsAny{className}))
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( callInit) &&
        Lns_GetEnv().SetStackVal( Lns_op_not(self.outputBuiltinFlag)) ).(bool)){
        self.FP.Write(Lns_getVM().String_format("%s( _pEnv, pAny", []LnsAny{self.moduleCtrl.FP.GetCtorName(classType)}))
        var scope *Ast_Scope
        scope = Lns_unwrap( classType.FP.Get_scope()).(*Ast_Scope)
        if Lns_op_not(self.outputBuiltinFlag){
            var initFuncType *Ast_TypeInfo
            initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField("__init", true, scope, convCC_scopeAccess)).(*Ast_TypeInfo)
            for _index, _ := range( initFuncType.FP.Get_argTypeInfoList().Items ) {
                index := _index + 1
                self.FP.Write(Lns_getVM().String_format(", arg%d", []LnsAny{index}))
            }
        }
        self.FP.Writeln(");")
    }
    if (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.outputBuiltinFlag) ||
        Lns_GetEnv().SetStackVal( Lns_op_not(self.canConv)) ).(bool)){
        self.FP.Writeln("pObj->pExt = NULL;")
    }
    self.FP.Writeln("pObj->pImp = &pObj->imp;")
    self.FP.Writeln("pObj->imp.sentinel.type = lns_value_type_none;")
    convCC_processIFObjInit_2097_(self.stream.FP, self.moduleCtrl, classType, classType)
}

// 3781: decl @lune.@base.@convCC.convFilter.processMapping
func (self *convCC_convFilter) processMapping(node *Nodes_DeclClassNode,classType *Ast_TypeInfo,out2HMode LnsInt) {
    if Lns_op_not(classType.FP.IsInheritFrom(self.processInfo, Ast_builtinTypeMapping, nil)){
        return 
    }
    var classScope *Ast_Scope
    classScope = Lns_unwrap( classType.FP.Get_scope()).(*Ast_Scope)
    var toMapMtdSym *Ast_SymbolInfo
    toMapMtdSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild("_toMap")).(*Ast_SymbolInfo)
    var fromMapMtdSym *Ast_SymbolInfo
    fromMapMtdSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild("_fromMap")).(*Ast_SymbolInfo)
    var className string
    className = self.moduleCtrl.FP.GetClassCName(classType)
    var processDeclToMap func(callFlag bool)
    processDeclToMap = func(callFlag bool) {
        self.FP.Write(Lns_getVM().String_format("%s%s ", []LnsAny{convCC_getOut2HeaderPrefix_1518_(out2HMode), convCC_cTypeAnyP}))
        if callFlag{
            self.FP.Write(self.moduleCtrl.FP.GetCallMethodCName(toMapMtdSym.FP.Get_typeInfo()))
        } else { 
            self.FP.Write(self.moduleCtrl.FP.GetMethodCName(toMapMtdSym.FP.Get_typeInfo()))
        }
        self.FP.Write(Lns_getVM().String_format("( %s _pEnv, %s pObj)", []LnsAny{convCC_cTypeEnvP, convCC_getCRetType_1195_(toMapMtdSym.FP.Get_typeInfo().FP.Get_retTypeInfoList())}))
    }
    var processDeclFromMap func(sub bool)
    processDeclFromMap = func(sub bool) {
        self.FP.Write(Lns_getVM().String_format("%s%s ", []LnsAny{convCC_getOut2HeaderPrefix_1518_(out2HMode), convCC_cTypeStem}))
        self.FP.Write(self.moduleCtrl.FP.GetMethodCName(fromMapMtdSym.FP.Get_typeInfo()))
        if sub{
            self.FP.Write("Sub")
        }
        self.FP.Write(Lns_getVM().String_format("( %s _pEnv", []LnsAny{convCC_cTypeEnvP}))
        if sub{
            self.FP.Write(", const lns_fromVal_info_t * pInfoArray")
        }
        self.FP.Write(Lns_getVM().String_format(", %s mapStem)", []LnsAny{convCC_getCRetType_1195_(fromMapMtdSym.FP.Get_typeInfo().FP.Get_retTypeInfoList())}))
    }
    var processToMapBody func()
    processToMapBody = func() {
        processDeclToMap(false)
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        self.FP.Writeln(Lns_getVM().String_format("%s pMap = lns_class_Map_new( _pEnv );", []LnsAny{convCC_cTypeAnyP}))
        for _, _varName := range( Ast_getAllNameForKind(classType, Ast_MethodKind__Object, Ast_SymbolKind__Mbr).FP.Get_list().Items ) {
            varName := _varName.(string)
            self.FP.Write("lns_mtd_Map_add( _pEnv, pMap, ")
            self.FP.Write(Lns_getVM().String_format("LNS_STEM_ANY( lns_litStr2any( _pEnv, \"%s\" ) ), ", []LnsAny{varName}))
            var memberSym *Ast_SymbolInfo
            memberSym = Lns_unwrap( classScope.FP.GetSymbolInfoField(varName, true, classScope, Ast_ScopeAccess__Full)).(*Ast_SymbolInfo)
            var nonNilMemberType *Ast_TypeInfo
            nonNilMemberType = memberSym.FP.Get_typeInfo().FP.Get_nonnilableType().FP.Get_srcTypeInfo()
            var valKind LnsInt
            valKind = convCC_getValKind_1167_(memberSym.FP.Get_typeInfo())
            var valTxt LnsAny
            valTxt = convCC_getAccessMember_2069_(className, "pObj", varName)
            if _switch17864 := valKind; _switch17864 == convCC_ValKind__Prim {
            } else if _switch17864 == convCC_ValKind__Stem {
                self.FP.Writeln(Lns_getVM().String_format("lns_toMapFromStem( _pEnv, %s ) );", []LnsAny{valTxt}))
                valTxt = nil
                
            } else if _switch17864 == convCC_ValKind__Any {
                if nonNilMemberType == Ast_builtinTypeString{
                } else { 
                    self.FP.Writeln(Lns_getVM().String_format("lns_toMapFromStem( _pEnv, LNS_STEM_ANY( %s ) ) );", []LnsAny{valTxt}))
                    valTxt = nil
                    
                }
            } else {
                Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{convCC_ValKind_getTxt( valKind)}))
            }
            if valTxt != nil{
                valTxt_7389 := valTxt.(string)
                convCC_process2stem_2060_(self.stream.FP, self.moduleCtrl, self.scopeMgr, convCC_getValKind_1167_(memberSym.FP.Get_typeInfo()), memberSym.FP.Get_typeInfo(), &node.Nodes_Node, convCC_process2stemCallback_2057_(func() {
                    self.FP.Write(valTxt_7389)
                }))
                self.FP.Writeln(");")
            }
        }
        self.FP.Writeln("return pMap;")
        self.FP.PopIndent()
        self.FP.Writeln("}")
        processDeclToMap(true)
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        self.FP.Write("return ")
        self.FP.Write(convCC_getAccessMethod_2072_(className, "pObj", "_toMap"))
        self.FP.Writeln("( _pEnv, pObj );")
        self.FP.PopIndent()
        self.FP.Writeln("}")
    }
    var processFromMapBody func()
    processFromMapBody = func() {
        processDeclFromMap(false)
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        self.FP.Writeln(Lns_getVM().String_format("return %sSub( _pEnv, NULL, mapStem );", []LnsAny{self.moduleCtrl.FP.GetMethodCName(fromMapMtdSym.FP.Get_typeInfo())}))
        self.FP.PopIndent()
        self.FP.Writeln("}")
        processDeclFromMap(true)
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        self.FP.Writeln(Lns_getVM().String_format("if ( mapStem.type == lns_stem_type_nil ) { return %s; }", []LnsAny{convCC_cValNil}))
        self.FP.Writeln(Lns_getVM().String_format("lns_any_t * pMap = mapStem%s;", []LnsAny{convCC_accessAny}))
        self.FP.Writeln("lns_any_t * pErr = NULL;")
        for _, _varName := range( Ast_getAllNameForKind(classType, Ast_MethodKind__Object, Ast_SymbolKind__Mbr).FP.Get_list().Items ) {
            varName := _varName.(string)
            var memberSym *Ast_SymbolInfo
            memberSym = Lns_unwrap( classScope.FP.GetSymbolInfoField(varName, true, classScope, Ast_ScopeAccess__Full)).(*Ast_SymbolInfo)
            var nonNilMemberType *Ast_TypeInfo
            nonNilMemberType = memberSym.FP.Get_typeInfo().FP.Get_nonnilableType().FP.Get_srcTypeInfo()
            var fromMapSym LnsAny
            if _switch18280 := nonNilMemberType.FP.Get_kind(); _switch18280 == Ast_TypeInfoKind__List {
                fromMapSym = "lns_fromMapToList"
                
            } else if _switch18280 == Ast_TypeInfoKind__Array {
                fromMapSym = "lns_fromMapToArray"
                
            } else if _switch18280 == Ast_TypeInfoKind__Set {
                fromMapSym = "lns_fromMapToSet"
                
            } else if _switch18280 == Ast_TypeInfoKind__Map {
                fromMapSym = "lns_fromMapToMap"
                
            } else {
                {
                    _memberClassScope := nonNilMemberType.FP.Get_scope()
                    if _memberClassScope != nil {
                        memberClassScope := _memberClassScope.(*Ast_Scope)
                        {
                            _symbol := memberClassScope.FP.GetSymbolInfoField("_fromMap", true, memberClassScope, Ast_ScopeAccess__Normal)
                            if _symbol != nil {
                                symbol := _symbol.(*Ast_SymbolInfo)
                                fromMapSym = self.moduleCtrl.FP.GetMethodCName(symbol.FP.Get_typeInfo())
                                
                            } else {
                                fromMapSym = nil
                                
                            }
                        }
                    } else {
                        fromMapSym = nil
                        
                    }
                }
            }
            var process func(nilable bool)
            process = func(nilable bool) {
                var kind string
                if _switch18348 := nonNilMemberType; _switch18348 == Ast_builtinTypeInt || _switch18348 == Ast_builtinTypeChar {
                    kind = "lns_stem_type_int"
                    
                } else if _switch18348 == Ast_builtinTypeReal {
                    kind = "lns_stem_type_real"
                    
                } else if _switch18348 == Ast_builtinTypeBool {
                    kind = "lns_stem_type_bool"
                    
                } else {
                    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{memberSym.FP.Get_typeInfo().FP.GetTxt(nil, nil, nil)}))
                }
                self.FP.Writeln(Lns_getVM().String_format("lns_check_err_from_map( pErr, _pEnv, pMap, %s, %s, %s, %s );", []LnsAny{Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( nilable) &&
                    Lns_GetEnv().SetStackVal( "true") ||
                    Lns_GetEnv().SetStackVal( "false") ).(string), memberSym.FP.Get_name(), kind, convCC_getAccessPrimValFromStem_1851_(false, memberSym.FP.Get_typeInfo(), 0)}))
            }
            self.FP.Writeln(Lns_getVM().String_format("%s %s;", []LnsAny{convCC_getCType_1188_(memberSym.FP.Get_typeInfo()), memberSym.FP.Get_name()}))
            var nilable bool
            nilable = memberSym.FP.Get_typeInfo().FP.Get_nilable()
            if nilable{
                self.FP.Writeln(Lns_getVM().String_format("%s = lns_global.nilStem;", []LnsAny{memberSym.FP.Get_name()}))
            }
            if fromMapSym != nil{
                fromMapSym_7423 := fromMapSym.(string)
                self.FP.Write("lns_check_err_from_map_class")
                var infoValName string
                if memberSym.FP.Get_typeInfo().FP.Get_itemTypeInfoList().Len() > 0{
                    infoValName = Lns_getVM().String_format("&info_0_1_%s_%s", []LnsAny{className, memberSym.FP.Get_name()})
                    
                } else { 
                    infoValName = "NULL"
                    
                }
                self.FP.Writeln(Lns_getVM().String_format("( pErr, _pEnv, pMap, %s, %s, %sSub, %s, %s );", []LnsAny{nilable, memberSym.FP.Get_name(), fromMapSym_7423, infoValName, convCC_getAccessPrimValFromStem_1851_(false, memberSym.FP.Get_typeInfo(), 0)}))
            } else {
                if nonNilMemberType.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil){
                    self.FP.Writeln(Lns_getVM().String_format("lns_check_err_from_map_str( pErr, _pEnv, pMap, %s, %s, %s );", []LnsAny{nilable, memberSym.FP.Get_name(), convCC_getAccessPrimValFromStem_1851_(false, memberSym.FP.Get_typeInfo(), 0)}))
                } else if nonNilMemberType.FP.Get_kind() == Ast_TypeInfoKind__Alternate{
                    self.FP.Writeln(Lns_getVM().String_format("lns_check_err_from_map_stem( pErr, _pEnv, pMap, %s, %s );", []LnsAny{nilable, memberSym.FP.Get_name()}))
                } else { 
                    process(nilable)
                }
            }
        }
        self.FP.Writeln("if ( pErr != NULL ) {")
        self.FP.PushIndent(nil)
        self.FP.Write(Lns_getVM().String_format("return lns_createMRet( _pEnv, false, 2, %s, ", []LnsAny{convCC_cValNil}))
        self.FP.Writeln("LNS_STEM_ANY( pErr ) );")
        self.FP.PopIndent()
        self.FP.Writeln("}")
        self.FP.processNewInsance(classType, false)
        for _, _varName := range( Ast_getAllNameForKind(classType, Ast_MethodKind__Object, Ast_SymbolKind__Mbr).FP.Get_list().Items ) {
            varName := _varName.(string)
            var memberSym *Ast_SymbolInfo
            memberSym = Lns_unwrap( classScope.FP.GetSymbolInfoField(varName, true, classScope, Ast_ScopeAccess__Full)).(*Ast_SymbolInfo)
            if _switch18760 := convCC_getValKind_1167_(memberSym.FP.Get_typeInfo()); _switch18760 == convCC_ValKind__Stem {
                self.FP.Writeln(Lns_getVM().String_format("lns_setQ( pObj->%s, %s );", []LnsAny{varName, varName}))
            } else if _switch18760 == convCC_ValKind__Any {
                self.FP.Writeln(Lns_getVM().String_format("lns_setQ_any( &pObj->%s, %s );", []LnsAny{varName, varName}))
            } else if _switch18760 == convCC_ValKind__Prim {
                self.FP.Writeln(Lns_getVM().String_format("pObj->%s = %s;", []LnsAny{varName, varName}))
            }
        }
        self.FP.Writeln("return lns_createMRet( _pEnv, false, 1, LNS_STEM_ANY( pAny ) );")
        self.FP.PopIndent()
        self.FP.Writeln("}")
    }
    var processFromMapInfo func(memberSym *Ast_SymbolInfo)
    processFromMapInfo = func(memberSym *Ast_SymbolInfo) {
        if memberSym.FP.Get_typeInfo().FP.Get_nonnilableType().FP.Get_itemTypeInfoList().Len() == 0{
            return 
        }
        var processGenType func(genType *Ast_TypeInfo,name string,depth LnsInt,index LnsInt)
        processGenType = func(genType *Ast_TypeInfo,name string,depth LnsInt,index LnsInt) {
            self.FP.Write(Lns_getVM().String_format("const lns_fromVal_info_t info_%d_%d_%s = { ", []LnsAny{depth, index, name}))
            self.FP.Write(Lns_getVM().String_format("%s, ", []LnsAny{genType.FP.Get_nilable()}))
            if _switch19016 := genType.FP.Get_nonnilableType().FP.Get_srcTypeInfo(); _switch19016 == Ast_builtinTypeStem {
                self.FP.Write("lns_fromMapToStemSub")
            } else if _switch19016 == Ast_builtinTypeInt || _switch19016 == Ast_builtinTypeChar {
                self.FP.Write("lns_fromMapToIntSub")
            } else if _switch19016 == Ast_builtinTypeReal {
                self.FP.Write("lns_fromMapToRealSub")
            } else if _switch19016 == Ast_builtinTypeBool {
                self.FP.Write("lns_fromMapToBoolSub")
            } else if _switch19016 == Ast_builtinTypeString {
                self.FP.Write("lns_fromMapToStrSub")
            } else {
                if _switch19014 := genType.FP.Get_nonnilableType().FP.Get_kind(); _switch19014 == Ast_TypeInfoKind__List || _switch19014 == Ast_TypeInfoKind__Array {
                    self.FP.Write("lns_fromMapToListSub")
                } else if _switch19014 == Ast_TypeInfoKind__Set {
                    self.FP.Write("lns_fromMapToSetSub")
                } else if _switch19014 == Ast_TypeInfoKind__Map {
                    self.FP.Write("lns_fromMapToMapSub")
                } else {
                    {
                        _memberClassScope := genType.FP.Get_scope()
                        if _memberClassScope != nil {
                            memberClassScope := _memberClassScope.(*Ast_Scope)
                            {
                                _symbol := memberClassScope.FP.GetSymbolInfoField("_fromMap", true, memberClassScope, Ast_ScopeAccess__Normal)
                                if _symbol != nil {
                                    symbol := _symbol.(*Ast_SymbolInfo)
                                    self.FP.Write(Lns_getVM().String_format("%sSub", []LnsAny{self.moduleCtrl.FP.GetMethodCName(symbol.FP.Get_typeInfo())}))
                                }
                            }
                        }
                    }
                }
            }
            if genType.FP.Get_itemTypeInfoList().Len() > 0{
                self.FP.Write(", { ")
                for _subIndex, _ := range( genType.FP.Get_itemTypeInfoList().Items ) {
                    subIndex := _subIndex + 1
                    if subIndex > 1{
                        self.FP.Write(", ")
                    }
                    self.FP.Write(Lns_getVM().String_format("&info_%d_%d_%s", []LnsAny{depth + 1, subIndex, name}))
                }
                self.FP.Write("}")
            }
            self.FP.Writeln("};")
        }
        var process func(typeInfo *Ast_TypeInfo,name string,depth LnsInt,index LnsInt)
        process = func(typeInfo *Ast_TypeInfo,name string,depth LnsInt,index LnsInt) {
            for _genIndex, _genType := range( typeInfo.FP.Get_itemTypeInfoList().Items ) {
                genIndex := _genIndex + 1
                genType := _genType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if genType.FP.Get_itemTypeInfoList().Len() > 0{
                    for _, _subGenType := range( genType.FP.Get_itemTypeInfoList().Items ) {
                        subGenType := _subGenType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        process(subGenType, name, depth + 2, genIndex)
                    }
                }
                processGenType(genType, name, depth + 1, genIndex)
            }
            processGenType(typeInfo, name, depth, index)
        }
        process(memberSym.FP.Get_typeInfo(), Lns_getVM().String_format("%s_%s", []LnsAny{className, memberSym.FP.Get_name()}), 0, 1)
    }
    if _switch19323 := self.processMode; _switch19323 == convCC_ProcessMode__Prototype {
        processDeclToMap(true)
        self.FP.Writeln(";")
        processDeclToMap(false)
        self.FP.Writeln(";")
        for _, _varName := range( Ast_getAllNameForKind(classType, Ast_MethodKind__Object, Ast_SymbolKind__Mbr).FP.Get_list().Items ) {
            varName := _varName.(string)
            var memberSym *Ast_SymbolInfo
            memberSym = Lns_unwrap( classScope.FP.GetSymbolInfoField(varName, true, classScope, Ast_ScopeAccess__Full)).(*Ast_SymbolInfo)
            processFromMapInfo(memberSym)
        }
        processDeclFromMap(false)
        self.FP.Writeln(";")
        processDeclFromMap(true)
        self.FP.Writeln(";")
    } else if _switch19323 == convCC_ProcessMode__DefClass {
        processToMapBody()
        processFromMapBody()
    }
}

// 4157: decl @lune.@base.@convCC.convFilter.processDeclClassNodePrototype
func (self *convCC_convFilter) processDeclClassNodePrototype(node *Nodes_DeclClassNode) {
    var className string
    className = self.moduleCtrl.FP.GetClassCName(node.FP.Get_expType())
    var kind LnsInt
    kind = node.FP.Get_expType().FP.Get_kind()
    var process func(out2HMode LnsInt)
    process = func(out2HMode LnsInt) {
        self.FP.Writeln(Lns_getVM().String_format("typedef struct lns_mtd_%s_t {", []LnsAny{className}))
        self.FP.PushIndent(nil)
        if kind == Ast_TypeInfoKind__Class{
            self.FP.Writeln("lns_del_t * _del;")
            self.FP.Writeln("lns_gc_t * _gc;")
        }
        convCC_processDeclMethodTable_2020_(self.stream.FP, node.FP.Get_expType())
        self.FP.PopIndent()
        self.FP.Writeln(Lns_getVM().String_format("} lns_mtd_%s_t;", []LnsAny{className}))
        if kind == Ast_TypeInfoKind__Class{
            self.FP.Writeln(Lns_getVM().String_format("typedef struct u_if_imp_%s_t {", []LnsAny{className}))
            self.FP.PushIndent(nil)
            convCC_processIFObjDecl_2094_(self.stream.FP, self.moduleCtrl, node.FP.Get_expType())
            self.FP.Writeln(Lns_getVM().String_format("%s sentinel;", []LnsAny{convCC_cTypeAny}))
            self.FP.PopIndent()
            self.FP.Writeln(Lns_getVM().String_format("} u_if_imp_%s_t;", []LnsAny{className}))
        }
        self.FP.Writeln(Lns_getVM().String_format("typedef struct %s {", []LnsAny{className}))
        self.FP.PushIndent(nil)
        self.FP.Writeln("lns_type_meta_t * pMeta;")
        if _switch19657 := kind; _switch19657 == Ast_TypeInfoKind__Class {
            self.FP.Writeln(Lns_getVM().String_format("u_if_imp_%s_t * pImp;", []LnsAny{className}))
            self.FP.Writeln(Lns_getVM().String_format("lns_mtd_%s_t * pMtd;", []LnsAny{className}))
            convCC_processDeclMemberTable_2034_(Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Lns_op_not(self.outputBuiltinFlag)) &&
                Lns_GetEnv().SetStackVal( self.canConv) ).(bool), self.stream.FP, node.FP.Get_expType())
            self.FP.Writeln("// interface implements")
            self.FP.Writeln(Lns_getVM().String_format("u_if_imp_%s_t imp;", []LnsAny{className}))
        } else if _switch19657 == Ast_TypeInfoKind__IF {
            self.FP.Writeln(Lns_getVM().String_format("%s pObj;", []LnsAny{convCC_cTypeAnyP}))
            self.FP.Writeln(Lns_getVM().String_format("lns_mtd_%s_t * pMtd;", []LnsAny{className}))
        }
        self.FP.PopIndent()
        self.FP.Writeln(Lns_getVM().String_format("} %s;", []LnsAny{className}))
        if _switch19832 := kind; _switch19832 == Ast_TypeInfoKind__Class {
            self.FP.Writeln(Lns_getVM().String_format("#define lns_mtd_%s( OBJ )                     \\\n                (((%s*)OBJ->val.classVal)->pMtd )", []LnsAny{className, className}))
            self.FP.Writeln(Lns_getVM().String_format("#define lns_obj_%s( OBJ ) ((%s*)OBJ->val.classVal)", []LnsAny{className, className}))
            self.FP.Writeln(Lns_getVM().String_format("#define lns_if_%s( OBJ ) ((%s*)OBJ->val.classVal)->pImp", []LnsAny{className, className}))
            if Lns_op_not(node.FP.Get_expType().FP.Get_abstractFlag()){
                convCC_processNewConstrProto_1990_(self.stream.FP, self.moduleCtrl, node, out2HMode, self.outputBuiltinFlag)
                self.stream.FP.Writeln(";")
            }
            self.FP.processMapping(node, node.FP.Get_expType(), out2HMode)
        } else if _switch19832 == Ast_TypeInfoKind__IF {
            self.FP.Writeln(Lns_getVM().String_format("#define lns_mtd_%s( OBJ )                     \\\n             ((%s*)&OBJ->val.ifVal)->pMtd", []LnsAny{className, className}))
            if out2HMode == convCC_Out2HMode__HeaderPub{
                self.FP.Writeln(Lns_getVM().String_format("extern lns_type_meta_t %s;", []LnsAny{self.moduleCtrl.FP.GetClassMetaName(node.FP.Get_expType())}))
            }
        }
    }
    {
        var processwork func(out2HMode LnsInt)
        processwork = func(out2HMode LnsInt) {
            if _switch19893 := out2HMode; _switch19893 == convCC_Out2HMode__HeaderPub || _switch19893 == convCC_Out2HMode__SourcePri {
                process(out2HMode)
            }
        }
        if Ast_isPubToExternal(node.FP.Get_expType().FP.Get_accessMode()){
            self.stream.FP.SwitchToHeader()
            processwork(convCC_Out2HMode__HeaderPub)
            self.stream.FP.ReturnToSource()
            processwork(convCC_Out2HMode__SourcePub)
        } else { 
            processwork(convCC_Out2HMode__SourcePri)
        }
    }
    
    if kind == Ast_TypeInfoKind__Class{
        convCC_processDeclClassPrototype_2088_(Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(self.outputBuiltinFlag)) &&
            Lns_GetEnv().SetStackVal( self.canConv) ).(bool), self.stream.FP, self.moduleCtrl, node)
        convCC_processAdvertise_2075_(self.stream.FP, self.moduleCtrl, self.scopeMgr, self.processMode, node)
        if Lns_op_not(self.outputBuiltinFlag){
            convCC_processDefaultCtor_2091_(self.stream.FP, self.moduleCtrl, self.scopeMgr, node)
        }
    }
}

// 4269: decl @lune.@base.@convCC.convFilter.isManagedAnySymbol
func (self *convCC_convFilter) isManagedAnySymbol(symbol Ast_LowSymbol) bool {
    var scope *Ast_Scope
    scope = symbol.Get_scope()
    if scope.FP.GetNamespaceTypeInfo().FP.GetModule() != self.moduleTypeInfo{
        return false
    }
    var valKind LnsInt
    valKind = self.scopeMgr.FP.GetSymbolValKind(symbol)
    var varName string
    varName = self.moduleCtrl.FP.GetSymbolName(symbol)
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( valKind == convCC_ValKind__Any) &&
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( symbol.Get_kind() == Ast_SymbolKind__Var) ||
            Lns_GetEnv().SetStackVal( symbol.Get_kind() == Ast_SymbolKind__Arg) &&
            Lns_GetEnv().SetStackVal( symbol.Get_mutable()) ||
            Lns_GetEnv().SetStackVal( convCC_isClassMember_1216_(symbol)) ).(bool))) ).(bool)){
        if varName == "self"{
            return false
        }
        return true
    }
    return false
}

// 4294: decl @lune.@base.@convCC.convFilter.processDeclClassDef
func (self *convCC_convFilter) processDeclClassDef(node *Nodes_DeclClassNode) {
    var className string
    className = self.moduleCtrl.FP.GetClassCName(node.FP.Get_expType())
    self.FP.Writeln(Lns_getVM().String_format("static void mtd_%s__del( lns_env_t * _pEnv, %s pObj ) {", []LnsAny{className, convCC_cTypeAnyP}))
    self.FP.PushIndent(nil)
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.outputBuiltinFlag) ||
        Lns_GetEnv().SetStackVal( Lns_op_not(self.canConv)) ).(bool){
        self.FP.Writeln(Lns_getVM().String_format("mtd_%s__delExt( _pEnv, pObj );", []LnsAny{className}))
    }
    if node.FP.Get_expType().FP.HasBase(){
        self.FP.Writeln(Lns_getVM().String_format("mtd_%s__del( _pEnv, pObj );", []LnsAny{self.moduleCtrl.FP.GetClassCName(node.FP.Get_expType().FP.Get_baseTypeInfo())}))
    }
    for _, _member := range( node.FP.Get_memberList().Items ) {
        member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if Lns_op_not(member.FP.Get_staticFlag()){
            var valKind LnsInt
            valKind = self.scopeMgr.FP.GetSymbolValKind(member.FP.Get_symbolInfo().FP)
            if _switch20337 := valKind; _switch20337 == convCC_ValKind__Stem {
                var typeInfo *Ast_TypeInfo
                typeInfo = member.FP.Get_symbolInfo().FP.Get_typeInfo()
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( typeInfo.FP.Get_nilable()) &&
                    Lns_GetEnv().SetStackVal( convCC_getValKind_1167_(typeInfo.FP.Get_nonnilableType()) == convCC_ValKind__Prim) ).(bool)){
                } else { 
                    self.FP.Writeln(Lns_getVM().String_format("lns_decre_ref_stem( _pEnv, %s );", []LnsAny{convCC_getAccessMember_2069_(className, "pObj", member.FP.Get_name().Txt)}))
                }
            } else if _switch20337 == convCC_ValKind__Any {
                self.FP.Writeln(Lns_getVM().String_format("lns_decre_ref( _pEnv, %s );", []LnsAny{convCC_getAccessMember_2069_(className, "pObj", member.FP.Get_name().Txt)}))
            }
        }
    }
    self.FP.PopIndent()
    self.FP.Writeln("}")
    if Lns_op_not(node.FP.Get_expType().FP.Get_abstractFlag()){
        convCC_processNewConstrProto_1990_(self.stream.FP, self.moduleCtrl, node, convCC_Out2HMode__SourcePub, self.outputBuiltinFlag)
        self.FP.Writeln(Lns_getVM().String_format("{ // %d", []LnsAny{node.FP.Get_pos().LineNo}))
        self.FP.PushIndent(nil)
        self.FP.processNewInsance(node.FP.Get_expType(), true)
        self.FP.Writeln("return pAny;")
        self.FP.PopIndent()
        self.FP.Writeln("}")
    }
    for _, _member := range( node.FP.Get_memberList().Items ) {
        member := _member.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var memberName string
        memberName = member.FP.Get_name().Txt
        if member.FP.Get_getterMode() != Ast_AccessMode__None{
            var getterType *Ast_TypeInfo
            getterType = Lns_unwrap( node.FP.Get_scope().FP.GetTypeInfoField(Lns_getVM().String_format("get_%s", []LnsAny{memberName}), true, node.FP.Get_scope(), convCC_scopeAccess)).(*Ast_TypeInfo)
            if getterType.FP.Get_autoFlag(){
                if getterType.FP.Get_staticFlag(){
                    convCC_processMethodDeclTxt_2015_(self.stream.FP, self.moduleCtrl, convCC_FuncWrap__CallWrap, getterType, nil)
                    self.FP.Writeln("{")
                    self.FP.PushIndent(nil)
                    if self.FP.isManagedAnySymbol(member.FP.Get_symbolInfo().FP){
                        self.FP.Writeln(Lns_getVM().String_format("return *%s;", []LnsAny{self.moduleCtrl.FP.GetClassMemberName(member.FP.Get_symbolInfo().FP)}))
                    } else { 
                        self.FP.Writeln(Lns_getVM().String_format("return %s;", []LnsAny{self.moduleCtrl.FP.GetClassMemberName(member.FP.Get_symbolInfo().FP)}))
                    }
                    self.FP.PopIndent()
                    self.FP.Writeln("}")
                } else { 
                    convCC_processMethodDeclTxt_2015_(self.stream.FP, self.moduleCtrl, convCC_FuncWrap__Normal, getterType, nil)
                    self.FP.Writeln("{")
                    self.FP.PushIndent(nil)
                    self.FP.Writeln(Lns_getVM().String_format("return %s;", []LnsAny{convCC_getAccessMember_2069_(className, "pObj", memberName)}))
                    self.FP.PopIndent()
                    self.FP.Writeln("}")
                    convCC_processMethodDeclTxt_2015_(self.stream.FP, self.moduleCtrl, convCC_FuncWrap__CallWrap, getterType, nil)
                    self.FP.Writeln("{")
                    self.FP.PushIndent(nil)
                    self.FP.Writeln(Lns_getVM().String_format("return lns_mtd_%s( pObj )->get_%s( _pEnv, pObj );", []LnsAny{className, memberName}))
                    self.FP.PopIndent()
                    self.FP.Writeln("}")
                }
            }
        }
        if member.FP.Get_setterMode() != Ast_AccessMode__None{
            var setterType *Ast_TypeInfo
            setterType = Lns_unwrap( node.FP.Get_scope().FP.GetTypeInfoField(Lns_getVM().String_format("set_%s", []LnsAny{memberName}), true, node.FP.Get_scope(), convCC_scopeAccess)).(*Ast_TypeInfo)
            var process func(accessMemberTxt string)
            process = func(accessMemberTxt string) {
                self.FP.Writeln("{")
                self.FP.PushIndent(nil)
                var valKind LnsInt
                valKind = self.scopeMgr.FP.GetSymbolValKind(member.FP.Get_symbolInfo().FP)
                if _switch20878 := valKind; _switch20878 == convCC_ValKind__Stem {
                    self.FP.Writeln(Lns_getVM().String_format("lns_setq( _pEnv, %s, arg1 );", []LnsAny{accessMemberTxt}))
                } else if _switch20878 == convCC_ValKind__Any {
                    self.FP.Writeln(Lns_getVM().String_format("lns_setq_any( _pEnv, &%s, arg1 );", []LnsAny{accessMemberTxt}))
                } else if _switch20878 == convCC_ValKind__Prim {
                    self.FP.Writeln(Lns_getVM().String_format("%s = arg1;", []LnsAny{accessMemberTxt}))
                } else {
                    Util_err(Lns_getVM().String_format("no support -- %s:%s:%d", []LnsAny{member.FP.Get_symbolInfo().FP.Get_name(), convCC_ValKind_getTxt( valKind), 4420}))
                }
                self.FP.PopIndent()
                self.FP.Writeln("}")
            }
            if setterType.FP.Get_autoFlag(){
                if setterType.FP.Get_staticFlag(){
                    convCC_processMethodDeclTxt_2015_(self.stream.FP, self.moduleCtrl, convCC_FuncWrap__Normal, setterType, nil)
                    var txt string
                    if self.FP.isManagedAnySymbol(member.FP.Get_symbolInfo().FP){
                        txt = Lns_getVM().String_format("(*%s)", []LnsAny{self.moduleCtrl.FP.GetClassMemberName(member.FP.Get_symbolInfo().FP)})
                        
                    } else { 
                        txt = self.moduleCtrl.FP.GetClassMemberName(member.FP.Get_symbolInfo().FP)
                        
                    }
                    process(txt)
                } else { 
                    convCC_processMethodDeclTxt_2015_(self.stream.FP, self.moduleCtrl, convCC_FuncWrap__Normal, setterType, nil)
                    process(convCC_getAccessMember_2069_(className, "pObj", memberName))
                    convCC_processMethodDeclTxt_2015_(self.stream.FP, self.moduleCtrl, convCC_FuncWrap__CallWrap, setterType, nil)
                    self.FP.Writeln("{")
                    self.FP.PushIndent(nil)
                    self.FP.Writeln(Lns_getVM().String_format("lns_mtd_%s( pObj )->set_%s( _pEnv, pObj, arg1 );", []LnsAny{className, memberName}))
                    self.FP.PopIndent()
                    self.FP.Writeln("}")
                }
            }
        }
    }
    convCC_processAdvertise_2075_(self.stream.FP, self.moduleCtrl, self.scopeMgr, self.processMode, node)
}

// 4658: decl @lune.@base.@convCC.convFilter.processDeclMember
func (self *convCC_convFilter) ProcessDeclMember(node *Nodes_DeclMemberNode,_opt LnsAny) {
}

// 4665: decl @lune.@base.@convCC.convFilter.processExpMacroExp
func (self *convCC_convFilter) ProcessExpMacroExp(node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    for _, _stmt := range( node.FP.Get_stmtList().Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        convCC_filter_1642_(stmt, self, &node.Nodes_Node)
        self.FP.Writeln("")
    }
}

// 4677: decl @lune.@base.@convCC.convFilter.outputDeclMacro
func (self *convCC_convFilter) OutputDeclMacro(name string,argNameList *LnsList,callback convCC_outputMacroStmtBlock_2219_) {
}

// 4721: decl @lune.@base.@convCC.convFilter.processDeclMacro
func (self *convCC_convFilter) ProcessDeclMacro(node *Nodes_DeclMacroNode,_opt LnsAny) {
}

// 4740: decl @lune.@base.@convCC.convFilter.processExpMacroStat
func (self *convCC_convFilter) ProcessExpMacroStat(node *Nodes_ExpMacroStatNode,_opt LnsAny) {
}

// 4764: decl @lune.@base.@convCC.convFilter.processDeclVarC
func (self *convCC_convFilter) processDeclVarC(declFlag bool,_var Ast_LowSymbol,init0 bool,manageScope LnsAny) {
    if declFlag{
        var typeTxt string
        typeTxt = convCC_convExp22103(Lns_2DDD(self.scopeMgr.FP.GetCTypeForSym(_var)))
        self.FP.Writeln(Lns_getVM().String_format("%s %s;", []LnsAny{typeTxt, self.moduleCtrl.FP.GetSymbolName(_var)}))
    }
    var valKind LnsInt
    valKind = self.scopeMgr.FP.GetSymbolValKind(_var)
    if valKind == convCC_ValKind__Prim{
        if init0{
            self.FP.Writeln(Lns_getVM().String_format("%s = 0;", []LnsAny{self.moduleCtrl.FP.GetSymbolName(_var)}))
        }
        return 
    }
    var scope *Ast_Scope
    
    {
        _scope := manageScope
        if _scope == nil{
            scope = _var.Get_scope()
            
        } else {
            scope = _scope.(*Ast_Scope)
        }
    }
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( Lns_op_not(init0)) ||
        Lns_GetEnv().SetStackVal( convCC_isStemType_1171_(_var.Get_typeInfo())) ).(bool){
        {
            _symbolInfo := Ast_SymbolInfoDownCastF(_var)
            if _symbolInfo != nil {
                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                if _switch22344 := valKind; _switch22344 == convCC_ValKind__Any {
                    self.FP.Write("lns_set_block_any")
                    self.FP.Writeln(Lns_getVM().String_format("( %s, %d, %s );", []LnsAny{convCC_getBlockName_1198_(scope), convCC_getSymbolIndex_1657_(symbolInfo), self.moduleCtrl.FP.GetSymbolName(_var)}))
                } else if _switch22344 == convCC_ValKind__Stem {
                    self.FP.Write("lns_set_block_stem")
                    self.FP.Writeln(Lns_getVM().String_format("( %s, %d, %s );", []LnsAny{convCC_getBlockName_1198_(scope), convCC_getSymbolIndex_1657_(symbolInfo), self.moduleCtrl.FP.GetSymbolName(_var)}))
                } else if _switch22344 == convCC_ValKind__Var {
                    var typeTxt string
                    typeTxt = convCC_getStemTypeId_1887_(convCC_getOrgTypeInfo_1430_(_var.Get_typeInfo()))
                    self.FP.Writeln(Lns_getVM().String_format("lns_set_block_var( %s, %d, %s, %s );", []LnsAny{convCC_getBlockName_1198_(scope), convCC_getSymbolIndex_1657_(symbolInfo), typeTxt, self.moduleCtrl.FP.GetSymbolName(_var)}))
                }
            }
        }
    } else { 
        var initVal string
        initVal = convCC_getLiteral2Stem_1884_("0", _var.Get_typeInfo())
        {
            _symbolInfo := Ast_SymbolInfoDownCastF(_var)
            if _symbolInfo != nil {
                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                if _switch22421 := valKind; _switch22421 == convCC_ValKind__Stem {
                    self.FP.Write("lns_initVal_stem")
                } else if _switch22421 == convCC_ValKind__Any {
                    self.FP.Write("lns_initVal_any")
                } else if _switch22421 == convCC_ValKind__Var {
                    self.FP.Write("lns_initVal_var")
                } else {
                    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{convCC_ValKind_getTxt( valKind)}))
                }
                self.FP.Writeln(Lns_getVM().String_format("( %s, %s, %d, %s );", []LnsAny{self.moduleCtrl.FP.GetSymbolName(_var), convCC_getBlockName_1198_(scope), convCC_getSymbolIndex_1657_(symbolInfo), initVal}))
            } else {
                self.FP.Writeln(Lns_getVM().String_format("%s = %s;", []LnsAny{self.moduleCtrl.FP.GetSymbolName(_var), initVal}))
            }
        }
    }
}

// 4839: decl @lune.@base.@convCC.convFilter.process__func__symbol
func (self *convCC_convFilter) process__func__symbol(funcTypeInfo *Ast_TypeInfo,has__func__Symbol bool,funcName string) {
    if Lns_op_not(has__func__Symbol){
        return 
    }
    if _switch22633 := self.processMode; _switch22633 == convCC_ProcessMode__WideScopeVer {
        var scope *Ast_Scope
        scope = Lns_unwrap( funcTypeInfo.FP.Get_scope()).(*Ast_Scope)
        var symbol *Ast_SymbolInfo
        symbol = Lns_unwrap( scope.FP.GetSymbolInfoChild("__func__")).(*Ast_SymbolInfo)
        self.FP.Writeln(Lns_getVM().String_format("static %s %s = NULL;", []LnsAny{convCC_cTypeAnyPP, self.moduleCtrl.FP.GetSymbolName(symbol.FP)}))
    } else if _switch22633 == convCC_ProcessMode__InitFuncSym {
        var scope *Ast_Scope
        scope = Lns_unwrap( funcTypeInfo.FP.Get_scope()).(*Ast_Scope)
        var symbol *Ast_SymbolInfo
        symbol = Lns_unwrap( scope.FP.GetSymbolInfoChild("__func__")).(*Ast_SymbolInfo)
        var symbolParam *convCC_SymbolParam
        symbolParam = self.scopeMgr.FP.GetSymbolParam(symbol.FP)
        var symbolName string
        symbolName = self.moduleCtrl.FP.GetSymbolName(symbol.FP)
        self.FP.Writeln(Lns_getVM().String_format("lns_set_block_any( pBlock, %d, %s );", []LnsAny{symbolParam.Index, symbolName}))
        self.FP.Writeln(Lns_getVM().String_format("lns_setQ_any( %s, lns_litStr2any( _pEnv, \"%s\") );", []LnsAny{symbolName, funcName}))
    }
}

// 4869: decl @lune.@base.@convCC.convFilter.processArgClosure
func (self *convCC_convFilter) processArgClosure(declInfo *Nodes_DeclFuncInfo) {
    for _, _argNode := range( declInfo.FP.Get_argList().Items ) {
        argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
        {
            _declArg := Nodes_DeclArgNodeDownCastF(argNode.FP)
            if _declArg != nil {
                declArg := _declArg.(*Nodes_DeclArgNode)
                var symbolInfo *Ast_SymbolInfo
                symbolInfo = declArg.FP.Get_symbolInfo()
                if symbolInfo.FP.Get_hasAccessFromClosure(){
                    var symbolParam *convCC_SymbolParam
                    symbolParam = self.scopeMgr.FP.GetSymbolParam(symbolInfo.FP)
                    self.FP.Writeln(Lns_getVM().String_format("%s %s;", []LnsAny{symbolParam.TypeTxt, self.moduleCtrl.FP.GetSymbolName(symbolInfo.FP)}))
                    self.FP.Write("lns_initVal_var(")
                    self.FP.Write(Lns_getVM().String_format(" %s, %s, %d, ", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbolInfo.FP), convCC_getBlockName_1198_(symbolInfo.FP.Get_scope()), symbolParam.Index}))
                    var valKind LnsInt
                    valKind = convCC_getValKind_1167_(symbolInfo.FP.Get_typeInfo())
                    var workSymbol *convCC_WorkSymbol
                    workSymbol = NewconvCC_WorkSymbol(symbolInfo.FP.Get_scope(), symbolInfo.FP.Get_accessMode(), Lns_getVM().String_format("_%s", []LnsAny{symbolInfo.FP.Get_name()}), symbolInfo.FP.Get_typeInfo(), symbolInfo.FP.Get_kind(), symbolInfo.FP.Get_staticFlag(), NewconvCC_SymbolParam(valKind, 0, convCC_getCType_1188_(symbolInfo.FP.Get_typeInfo())))
                    self.FP.processSym2stem(workSymbol.FP)
                    self.FP.Writeln(");")
                } else if symbolInfo.FP.Get_mutable(){
                    self.FP.processDeclVarC(true, symbolInfo.FP, false, nil)
                    if _switch22906 := convCC_getValKind_1167_(symbolInfo.FP.Get_typeInfo()); _switch22906 == convCC_ValKind__Stem {
                        self.FP.Writeln(Lns_getVM().String_format("lns_setQ( %s, _%s );", []LnsAny{symbolInfo.FP.Get_name(), symbolInfo.FP.Get_name()}))
                    } else if _switch22906 == convCC_ValKind__Any {
                        self.FP.Writeln(Lns_getVM().String_format("lns_setQ_any( %s, _%s );", []LnsAny{symbolInfo.FP.Get_name(), symbolInfo.FP.Get_name()}))
                    } else if _switch22906 == convCC_ValKind__Prim {
                        self.FP.Writeln(Lns_getVM().String_format("%s = _%s;", []LnsAny{symbolInfo.FP.Get_name(), symbolInfo.FP.Get_name()}))
                    }
                }
            }
        }
    }
}

// 4919: decl @lune.@base.@convCC.convFilter.processDeclMethodInfo
func (self *convCC_convFilter) processDeclMethodInfo(declInfo *Nodes_DeclFuncInfo,funcTypeInfo *Ast_TypeInfo,parent *Nodes_Node) {
    if _switch23385 := self.processMode; _switch23385 == convCC_ProcessMode__Prototype {
        convCC_processPrototypeMethod_2047_(self.stream.FP, self.moduleCtrl, declInfo.FP.Get_argList(), funcTypeInfo)
    } else if _switch23385 == convCC_ProcessMode__Form {
        var classType *Ast_TypeInfo
        classType = funcTypeInfo.FP.Get_parentInfo()
        var className string
        className = self.moduleCtrl.FP.GetClassCName(classType)
        {
            _body := declInfo.FP.Get_body()
            if _body != nil {
                body := _body.(*Nodes_BlockNode)
                convCC_processMethodDeclTxt_2015_(self.stream.FP, self.moduleCtrl, convCC_FuncWrap__Normal, funcTypeInfo, declInfo.FP.Get_argList())
                self.FP.Writeln("{")
                self.FP.PushIndent(nil)
                self.FP.pushRoutine(funcTypeInfo, body)
                {
                    _declClassNode := declInfo.FP.Get_declClassNode()
                    if _declClassNode != nil {
                        declClassNode := _declClassNode.(*Nodes_DeclClassNode)
                        if Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(declInfo.FP.Get_name()) && 
                        Lns_GetEnv().NilAccPush(Lns_GetEnv().NilAccPop().(*Types_Token).Txt)) == "__init"{
                            for _, _memberSym := range( declClassNode.FP.Get_uninitMemberList().Items ) {
                                memberSym := _memberSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                                if declInfo.FP.Get_staticFlag() == memberSym.FP.Get_staticFlag(){
                                    var memberAccess string
                                    if declInfo.FP.Get_staticFlag(){
                                        memberAccess = self.moduleCtrl.FP.GetClassMemberName(memberSym.FP)
                                        
                                    } else { 
                                        memberAccess = convCC_getAccessMember_2069_(className, "pObj", memberSym.FP.Get_name())
                                        
                                    }
                                    self.FP.Writeln(Lns_getVM().String_format("%s = %s;", []LnsAny{memberAccess, convCC_cValNil}))
                                }
                            }
                        }
                    }
                }
                if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                    Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_staticFlag()) &&
                    Lns_GetEnv().SetStackVal( funcTypeInfo.FP.Get_rawTxt() == "___init") ).(bool)){
                    for _, _symbol := range( (Lns_unwrap( classType.FP.Get_scope()).(*Ast_Scope)).FP.Get_symbol2SymbolInfoMap().Items ) {
                        symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                        if convCC_isClassMember_1216_(symbol.FP){
                            self.FP.processDeclVarC(false, symbol.FP, false, self.moduleTypeInfo.FP.Get_scope())
                        }
                    }
                }
                var scope *Ast_Scope
                scope = body.FP.Get_scope()
                {
                    _selfSymbol := scope.FP.GetSymbolInfoChild("self")
                    if _selfSymbol != nil {
                        selfSymbol := _selfSymbol.(*Ast_SymbolInfo)
                        var symbolParam *convCC_SymbolParam
                        symbolParam = self.scopeMgr.FP.GetSymbolParam(selfSymbol.FP)
                        if symbolParam.Kind == convCC_ValKind__Var{
                            self.FP.Writeln(Lns_getVM().String_format("%s self;", []LnsAny{convCC_cTypeVarP}))
                            self.FP.Writeln(Lns_getVM().String_format("lns_initVal_var( self, %s, %d, LNS_STEM_ANY( pObj ) );", []LnsAny{convCC_getBlockName_1198_(scope), symbolParam.Index}))
                        } else { 
                            if Lns_op_not(funcTypeInfo.FP.Get_staticFlag()){
                                self.FP.Writeln(Lns_getVM().String_format("%s self = pObj;", []LnsAny{convCC_cTypeAnyP}))
                            }
                        }
                    }
                }
                self.FP.processArgClosure(declInfo)
                self.duringDeclFunc = true
                
                convCC_filter_1642_(&body.Nodes_Node, self, parent)
                self.duringDeclFunc = false
                
                self.FP.popRoutine()
                self.FP.PopIndent()
                self.FP.Writeln("}")
            }
        }
        convCC_processDeclCallMethodWrapper_2063_(self.stream.FP, self.moduleCtrl, self.scopeMgr, parent, funcTypeInfo, true)
        convCC_processDeclCallMethodWrapper_2063_(self.stream.FP, self.moduleCtrl, self.scopeMgr, parent, funcTypeInfo, false)
    } else if _switch23385 == convCC_ProcessMode__InitFuncSym || _switch23385 == convCC_ProcessMode__WideScopeVer {
        self.FP.process__func__symbol(funcTypeInfo, declInfo.FP.Get_has__func__Symbol(), self.moduleCtrl.FP.GetMethodCName(funcTypeInfo))
    }
}

// 5013: decl @lune.@base.@convCC.convFilter.processDeclConstr
func (self *convCC_convFilter) ProcessDeclConstr(node *Nodes_DeclConstrNode,_opt LnsAny) {
    self.FP.processDeclMethodInfo(node.FP.Get_declInfo(), node.FP.Get_expType(), &node.Nodes_Node)
}

// 5019: decl @lune.@base.@convCC.convFilter.processDeclDestr
func (self *convCC_convFilter) ProcessDeclDestr(node *Nodes_DeclDestrNode,_opt LnsAny) {
}

// 5038: decl @lune.@base.@convCC.convFilter.getValKindOfNode
func (self *convCC_convFilter) getValKindOfNode(node *Nodes_Node) LnsInt {
    if node.FP.Get_expTypeList().Len() > 1{
        return convCC_ValKind__Stem
    }
    var symbolList *LnsList
    symbolList = node.FP.GetSymbolInfo()
    if symbolList.Len() > 0{
        return self.scopeMgr.FP.GetSymbolValKind(symbolList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP)
    }
    return convCC_getValKind_1167_(node.FP.Get_expType())
}

// 5056: decl @lune.@base.@convCC.convFilter.processVal2stem
func (self *convCC_convFilter) ProcessVal2stem(node *Nodes_Node,parent *Nodes_Node) {
    convCC_process2stem_2060_(self.stream.FP, self.moduleCtrl, self.scopeMgr, self.FP.getValKindOfNode(node), node.FP.Get_expType(), parent, convCC_process2stemCallback_2057_(func() {
        convCC_filter_1642_(node, self, parent)
    }))
}

// 5078: decl @lune.@base.@convCC.convFilter.processCallArgList
func (self *convCC_convFilter) ProcessCallArgList(funcType *Ast_TypeInfo,expListNode LnsAny) {
    var funcArgTypeList *LnsList
    funcArgTypeList = funcType.FP.Get_argTypeInfoList()
    var abbrValTxt string
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.moduleCtrl.FP.GetBuiltinFuncNameFromType(funcType)) ||
        Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( funcType.FP.Get_kind() == Ast_TypeInfoKind__Method) &&
            Lns_GetEnv().SetStackVal( funcType.FP.Get_parentInfo().FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil)) ).(bool))) )){
        abbrValTxt = convCC_cValNone
        
    } else { 
        abbrValTxt = convCC_cValNil
        
    }
    var processAbbr func(funcArgType *Ast_TypeInfo)
    processAbbr = func(funcArgType *Ast_TypeInfo) {
        if funcArgType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
            self.FP.Write(convCC_cValDDD0)
        } else { 
            self.FP.Write(abbrValTxt)
        }
    }
    if expListNode != nil{
        expListNode_7853 := expListNode.(*Nodes_ExpListNode)
        var expList *LnsList
        expList = expListNode_7853.FP.Get_expList()
        for _index, _funcArgType := range( funcArgTypeList.Items ) {
            index := _index + 1
            funcArgType := _funcArgType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            self.FP.Write(", ")
            if expList.Len() >= index{
                var expNode *Nodes_Node
                expNode = expList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                if expNode.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr{
                    processAbbr(funcArgType)
                } else { 
                    if funcArgType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                        if expNode.FP.Get_kind() == Nodes_NodeKind_get_Abbr(){
                            self.FP.Write(convCC_cValDDD0)
                        } else { 
                            convCC_filter_1642_(expNode, self, &expListNode_7853.Nodes_Node)
                        }
                        return 
                    } else { 
                        if convCC_isStemType_1171_(funcArgType){
                            self.FP.ProcessVal2stem(expNode, &expListNode_7853.Nodes_Node)
                        } else { 
                            convCC_filter_1642_(expNode, self, &expListNode_7853.Nodes_Node)
                        }
                    }
                }
            } else { 
                processAbbr(funcArgType)
            }
        }
    } else {
        for _, _funcArgType := range( funcArgTypeList.Items ) {
            funcArgType := _funcArgType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            self.FP.Write(", ")
            processAbbr(funcArgType)
        }
    }
}

// 5169: decl @lune.@base.@convCC.convFilter.processExpCallSuper
func (self *convCC_convFilter) ProcessExpCallSuper(node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    var funcType *Ast_TypeInfo
    if node.FP.Get_methodType().FP.Get_rawTxt() == "__init"{
        self.FP.Write(Lns_getVM().String_format("%s( _pEnv, pObj", []LnsAny{self.moduleCtrl.FP.GetCtorName(node.FP.Get_superType())}))
        var superScope *Ast_Scope
        superScope = Lns_unwrap( node.FP.Get_superType().FP.Get_scope()).(*Ast_Scope)
        funcType = Lns_unwrap( superScope.FP.GetTypeInfoChild("__init")).(*Ast_TypeInfo)
        
    } else { 
        self.FP.Write(Lns_getVM().String_format("%s( _pEnv, pObj", []LnsAny{self.moduleCtrl.FP.GetMethodCName(node.FP.Get_methodType())}))
        funcType = node.FP.Get_methodType()
        
    }
    self.FP.ProcessCallArgList(funcType, node.FP.Get_expList())
    self.FP.Writeln(");")
}

// 5188: decl @lune.@base.@convCC.convFilter.processDeclMethod
func (self *convCC_convFilter) ProcessDeclMethod(node *Nodes_DeclMethodNode,_opt LnsAny) {
    self.FP.processDeclMethodInfo(node.FP.Get_declInfo(), node.FP.Get_expType(), &node.Nodes_Node)
}

// 5194: decl @lune.@base.@convCC.convFilter.processProtoMethod
func (self *convCC_convFilter) ProcessProtoMethod(node *Nodes_ProtoMethodNode,_opt LnsAny) {
    if node.FP.Get_expType().FP.Get_abstractFlag(){
        self.FP.processDeclMethodInfo(node.FP.Get_declInfo(), node.FP.Get_expType(), &node.Nodes_Node)
    }
}

// 5203: decl @lune.@base.@convCC.convFilter.processUnwrapSet
func (self *convCC_convFilter) ProcessUnwrapSet(node *Nodes_UnwrapSetNode,_opt LnsAny) {
}

// 5237: decl @lune.@base.@convCC.convFilter.accessPrimValFromAny
func (self *convCC_convFilter) accessPrimValFromAny(dddFlag bool,typeInfo *Ast_TypeInfo,index LnsInt) {
    self.FP.Write(convCC_getAccessPrimValFromStem_1851_(dddFlag, typeInfo, index))
}

// 5243: decl @lune.@base.@convCC.convFilter.isStemSym
func (self *convCC_convFilter) isStemSym(symbolInfo Ast_LowSymbol) bool {
    return self.scopeMgr.FP.GetSymbolValKind(symbolInfo) != convCC_ValKind__Prim
}

// 5248: decl @lune.@base.@convCC.convFilter.isStemVal
func (self *convCC_convFilter) isStemVal(node *Nodes_Node) bool {
    if node.FP.Get_expTypeList().Len() > 1{
        return false
    }
    var symbolList *LnsList
    symbolList = node.FP.GetSymbolInfo()
    if symbolList.Len() > 0{
        return self.scopeMgr.FP.GetSymbolValKind(symbolList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP) != convCC_ValKind__Prim
    }
    return convCC_isStemType_1171_(node.FP.Get_expType())
}

// 5264: decl @lune.@base.@convCC.convFilter.accessPrimVal
func (self *convCC_convFilter) AccessPrimVal(exp *Nodes_Node,parent *Nodes_Node) {
    if _switch24237 := self.FP.getValKindOfNode(exp); _switch24237 == convCC_ValKind__Var {
        convCC_filter_1642_(exp, self, parent)
    } else if _switch24237 == convCC_ValKind__Prim {
        convCC_filter_1642_(exp, self, parent)
    } else if _switch24237 == convCC_ValKind__Stem {
        convCC_filter_1642_(exp, self, parent)
        self.FP.accessPrimValFromAny(exp.FP.Get_expTypeList().Len() > 1, exp.FP.Get_expType(), 0)
    } else if _switch24237 == convCC_ValKind__Any {
        convCC_filter_1642_(exp, self, parent)
    } else {
        Util_err(Lns_getVM().String_format("not support -- %d", []LnsAny{5297}))
    }
}

// 5331: decl @lune.@base.@convCC.convFilter.processSym2Any
func (self *convCC_convFilter) ProcessSym2Any(symbol Ast_LowSymbol) {
    var valKind LnsInt
    valKind = self.scopeMgr.FP.GetSymbolValKind(symbol)
    var symName string
    if self.FP.isManagedAnySymbol(symbol){
        symName = Lns_getVM().String_format("(*%s)", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbol)})
        
    } else { 
        symName = self.moduleCtrl.FP.GetSymbolName(symbol)
        
    }
    if _switch24379 := valKind; _switch24379 == convCC_ValKind__Stem {
        self.FP.Write(symName)
        self.FP.Write(convCC_accessAny)
    } else if _switch24379 == convCC_ValKind__Any {
        self.FP.Write(symName)
    } else if _switch24379 == convCC_ValKind__Var {
        self.FP.Write(symName)
        self.FP.Write(Lns_getVM().String_format("->stem%s", []LnsAny{convCC_accessAny}))
    } else {
        Util_err(Lns_getVM().String_format("not suppport -- %s, %d", []LnsAny{convCC_ValKind_getTxt( valKind), 5354}))
    }
}

// 5361: decl @lune.@base.@convCC.convFilter.processVal2any
func (self *convCC_convFilter) ProcessVal2any(node *Nodes_Node,parent *Nodes_Node) {
    var valKind LnsInt
    valKind = self.FP.getValKindOfNode(node)
    if _switch24487 := valKind; _switch24487 == convCC_ValKind__Stem {
        convCC_filter_1642_(node, self, parent)
        self.FP.Write(convCC_accessAny)
    } else if _switch24487 == convCC_ValKind__Any {
        convCC_filter_1642_(node, self, parent)
    } else if _switch24487 == convCC_ValKind__Var {
        convCC_filter_1642_(node, self, parent)
    } else {
        Util_err(Lns_getVM().String_format("not suppport -- %d, %s, %s, %d", []LnsAny{node.FP.Get_pos().LineNo, convCC_ValKind_getTxt( valKind), Nodes_getNodeKindName(node.FP.Get_kind()), 5379}))
    }
}

// 5401: decl @lune.@base.@convCC.convFilter.processSetValSingleDirect
func (self *convCC_convFilter) processSetValSingleDirect(parent *Nodes_Node,node LnsAny,_var Ast_LowSymbol,initFlag bool,expValKind LnsInt,expValType *Ast_TypeInfo,index LnsInt,firstMRet bool,processVal convCC_processRValue_2298_) {
    var valKind LnsInt
    valKind = self.scopeMgr.FP.GetSymbolValKind(_var)
    var varName string
    varName = self.moduleCtrl.FP.GetSymbolName(_var)
    if self.FP.isManagedAnySymbol(_var){
        varName = Lns_getVM().String_format("(*%s)", []LnsAny{varName})
        
    }
    var processPrefix LnsAny
    processPrefix = nil
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(node)
        if _fieldNode != nil {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            if Lns_isCondTrue( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(fieldNode.FP.Get_symbolInfo()) && 
            Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_SymbolInfo).FP.Get_staticFlag()}))){
            } else { 
                processPrefix = conv2Form24644(func() {
                    var prefixNode *Nodes_Node
                    prefixNode = fieldNode.FP.Get_prefix()
                    var className string
                    className = self.moduleCtrl.FP.GetClassCName(prefixNode.FP.Get_expType())
                    self.FP.Write(Lns_getVM().String_format("lns_obj_%s( ", []LnsAny{className}))
                    self.FP.ProcessVal2any(prefixNode, &fieldNode.Nodes_Node)
                    self.FP.Write(")->")
                })
                
            }
        }
    }
    if _var.Get_symbolId() == convCC_invalidSymbolId{
        if valKind == expValKind{
            self.FP.Write(Lns_getVM().String_format("%s = ", []LnsAny{varName}))
            processVal()
            self.FP.Write(";")
            return 
        }
        Util_err(Lns_getVM().String_format("illegal %s %s %s -- %d", []LnsAny{_var.Get_name(), convCC_ValKind_getTxt( valKind), convCC_ValKind_getTxt( expValKind), 5438}))
    }
    if _switch24970 := valKind; _switch24970 == convCC_ValKind__Var {
        if initFlag{
            self.FP.Write(Lns_getVM().String_format("lns_setQ( %s->stem, ", []LnsAny{varName}))
        } else { 
            self.FP.Write(Lns_getVM().String_format("lns_setq( _pEnv, %s->stem, ", []LnsAny{varName}))
        }
        convCC_process2stem_2060_(self.stream.FP, self.moduleCtrl, self.scopeMgr, expValKind, expValType, parent, convCC_process2stemCallback_2057_(processVal))
        self.FP.Write(" );")
    } else if _switch24970 == convCC_ValKind__Stem {
        if _switch24865 := expValKind; _switch24865 == convCC_ValKind__Stem || _switch24865 == convCC_ValKind__Any || _switch24865 == convCC_ValKind__Prim {
            if initFlag{
                self.FP.Write("lns_setQ( ")
            } else { 
                self.FP.Write("lns_setq( _pEnv, ")
            }
            if processPrefix != nil{
                processPrefix_7996 := processPrefix.(LnsForm)
                processPrefix_7996(Lns_2DDD([]LnsAny{}))
            }
            self.FP.Write(Lns_getVM().String_format("%s, ", []LnsAny{varName}))
            processVal()
            self.FP.Write(" );")
        }
    } else if _switch24970 == convCC_ValKind__Any {
        if initFlag{
            self.FP.Write("lns_setQ_any( &")
        } else { 
            self.FP.Write("lns_setq_any( _pEnv, &")
        }
        if processPrefix != nil{
            processPrefix_8001 := processPrefix.(LnsForm)
            processPrefix_8001(Lns_2DDD([]LnsAny{}))
        }
        self.FP.Write(Lns_getVM().String_format("%s, ", []LnsAny{varName}))
        processVal()
        self.FP.Write(" );")
    } else {
        if processPrefix != nil{
            processPrefix_8004 := processPrefix.(LnsForm)
            processPrefix_8004(Lns_2DDD([]LnsAny{}))
        }
        self.FP.Write(Lns_getVM().String_format("%s = ", []LnsAny{varName}))
        processVal()
        self.FP.Write(";")
    }
}

// 5566: decl @lune.@base.@convCC.convFilter.processSymForSetOp
func (self *convCC_convFilter) processSymForSetOp(parent *Nodes_Node,dstKind LnsInt,dstTypeInfo *Ast_TypeInfo,symbol Ast_LowSymbol) {
    var srcKind LnsInt
    srcKind = self.scopeMgr.FP.GetSymbolValKind(symbol)
    if dstKind != srcKind{
        if _switch25081 := dstKind; _switch25081 == convCC_ValKind__Prim {
            self.FP.Write(self.scopeMgr.FP.GetAccessPrimValFromSymbol(symbol))
            return 
        } else if _switch25081 == convCC_ValKind__Stem {
            self.FP.processSym2stem(symbol)
            return 
        } else if _switch25081 == convCC_ValKind__Var {
            self.FP.processSym2stem(symbol)
            return 
        } else if _switch25081 == convCC_ValKind__Any {
            self.FP.ProcessSym2Any(symbol)
            return 
        } else {
            Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{convCC_ValKind_getTxt( dstKind)}))
        }
    }
    var symName string
    if self.FP.isManagedAnySymbol(symbol){
        symName = Lns_getVM().String_format("(*%s)", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbol)})
        
    } else { 
        symName = self.moduleCtrl.FP.GetSymbolName(symbol)
        
    }
    self.FP.Write(symName)
}

// 5652: decl @lune.@base.@convCC.convFilter.needsWrapper
func (self *convCC_convFilter) needsWrapper(orgFunc *Ast_TypeInfo,castType *Ast_TypeInfo) bool {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( orgFunc.FP.Get_argTypeInfoList().Len() == castType.FP.Get_argTypeInfoList().Len()) &&
        Lns_GetEnv().SetStackVal( orgFunc.FP.Get_retTypeInfoList().Len() == castType.FP.Get_retTypeInfoList().Len()) ).(bool)){
        var check func(typeList1 *LnsList,typeList2 *LnsList) bool
        check = func(typeList1 *LnsList,typeList2 *LnsList) bool {
            for _index, _type1 := range( typeList1.Items ) {
                index := _index + 1
                type1 := _type1.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if Lns_op_not(type1.FP.Equals(self.processInfo, typeList2.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), nil, nil)){
                    return false
                }
            }
            return true
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( check(orgFunc.FP.Get_argTypeInfoList(), castType.FP.Get_argTypeInfoList())) &&
            Lns_GetEnv().SetStackVal( check(orgFunc.FP.Get_retTypeInfoList(), castType.FP.Get_retTypeInfoList())) ).(bool)){
            return false
        }
    }
    return true
}

// 5675: decl @lune.@base.@convCC.convFilter.processFuncCast2Form
func (self *convCC_convFilter) processFuncCast2Form(castType *Ast_TypeInfo,orgFunc *Ast_TypeInfo) {
    if Lns_op_not(self.FP.needsWrapper(orgFunc, castType)){
        self.FP.Write(convCC_getFunc2any_1897_(self.moduleCtrl, self.scopeMgr, orgFunc))
        return 
    }
    var closureSymList *LnsList
    closureSymList = Lns_unwrapDefault( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(orgFunc.FP.Get_scope()) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_closureSymList()})), NewLnsList([]LnsAny{})).(*LnsList)
    var argList *LnsList
    argList = castType.FP.Get_argTypeInfoList()
    var hasDDD bool
    hasDDD = Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( argList.Len() > 0) &&
        Lns_GetEnv().SetStackVal( argList.GetAt(argList.Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() == Ast_TypeInfoKind__DDD) ||
        Lns_GetEnv().SetStackVal( false) ).(bool)
    self.FP.Write(convCC_getPrepareClosure_1894_(self.scopeMgr, self.moduleCtrl.FP.GetFuncCastWrapName(orgFunc, castType), argList.Len(), hasDDD, closureSymList))
}

// 5706: decl @lune.@base.@convCC.convFilter.processValForSetOp
func (self *convCC_convFilter) processValForSetOp(parent *Nodes_Node,dstKind LnsInt,dstTypeInfo *Ast_TypeInfo,exp *Nodes_Node,index LnsInt,firstMRet bool) {
    var valKind LnsInt
    valKind = self.FP.getValKindOfNode(exp)
    var accessVal func()
    accessVal = func() {
        if firstMRet{
            convCC_processGetMRet_2313_(self.stream.FP, self.moduleCtrl, exp.FP.Get_expType(), 0)
        } else { 
            if exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Func{
                self.FP.processFuncCast2Form(dstTypeInfo, exp.FP.Get_expType())
            } else { 
                convCC_filter_1642_(exp, self, parent)
            }
        }
    }
    var processVal func()
    processVal = func() {
        if firstMRet{
            accessVal()
        } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( Lns_op_not(firstMRet)) &&
            Lns_GetEnv().SetStackVal( Nodes_hasMultiValNode(exp)) ).(bool)){
            self.FP.Write("lns_fromDDD( ")
            accessVal()
            self.FP.Write(convCC_accessAny)
            self.FP.Write(Lns_getVM().String_format(", %d )", []LnsAny{index}))
            self.FP.Write(convCC_getAccessValFromStem_1854_(exp.FP.Get_expType()))
        } else { 
            if dstKind == convCC_ValKind__Stem{
                self.FP.ProcessVal2stem(exp, parent)
            } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( dstKind == convCC_ValKind__Var) &&
                Lns_GetEnv().SetStackVal( valKind != convCC_ValKind__Stem) ).(bool)){
                accessVal()
            } else { 
                accessVal()
            }
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( dstKind == convCC_ValKind__Prim) &&
        Lns_GetEnv().SetStackVal( valKind == convCC_ValKind__Stem) ).(bool)){
        var expSymList *LnsList
        expSymList = exp.FP.GetSymbolInfo()
        if expSymList.Len() > 0{
            self.FP.Write(self.scopeMgr.FP.GetAccessPrimValFromSymbol(expSymList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP))
        } else { 
            processVal()
        }
    } else { 
        processVal()
    }
}

// 5819: decl @lune.@base.@convCC.convFilter.processSetValSingle
func (self *convCC_convFilter) processSetValSingle(parent *Nodes_Node,node LnsAny,_var Ast_LowSymbol,initFlag bool,exp *Nodes_Node,index LnsInt,firstMRet bool) {
    var expValKind LnsInt
    if firstMRet{
        expValKind = convCC_getValKind_1167_(exp.FP.Get_expType())
        
    } else { 
        expValKind = self.FP.getValKindOfNode(exp)
        
    }
    self.FP.processSetValSingleDirect(parent, node, _var, initFlag, expValKind, exp.FP.Get_expType(), index, firstMRet, convCC_processRValue_2298_(func() {
        self.FP.processValForSetOp(parent, self.scopeMgr.FP.GetSymbolValKind(_var), _var.Get_typeInfo(), exp, index, firstMRet)
    }))
}

// 5849: decl @lune.@base.@convCC.convFilter.processSetSymSingle
func (self *convCC_convFilter) processSetSymSingle(parent *Nodes_Node,node LnsAny,_var Ast_LowSymbol,initFlag bool,symbol Ast_LowSymbol,toIF bool) {
    var process func()
    process = func() {
        self.FP.processSymForSetOp(parent, self.scopeMgr.FP.GetSymbolValKind(_var), _var.Get_typeInfo(), symbol)
    }
    self.FP.processSetValSingleDirect(parent, node, _var, initFlag, self.scopeMgr.FP.GetSymbolValKind(symbol), symbol.Get_typeInfo(), 1, false, convCC_processRValue_2298_(func() {
        if toIF{
            convCC_processToIF_2310_(self.stream.FP, self.moduleCtrl, symbol.Get_typeInfo(), convCC_processExp_1881_(process))
        } else { 
            process()
        }
    }))
}

// 5882: decl @lune.@base.@convCC.convFilter.processSetValSingleNode
func (self *convCC_convFilter) processSetValSingleNode(parent *Nodes_Node,_var *Nodes_Node,initFlag bool,exp *Nodes_Node,index LnsInt,firstMRet bool) {
    var symbolList *LnsList
    symbolList = _var.FP.GetSymbolInfo()
    if symbolList.Len() > 0{
        self.FP.processSetValSingle(parent, _var, symbolList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP, initFlag, exp, index, firstMRet)
        return 
    }
    if _switch26477 := _var.FP.Get_kind(); _switch26477 == Nodes_NodeKind_get_ExpRefItem() {
        {
            _refItemNode := Nodes_ExpRefItemNodeDownCastF(_var.FP)
            if _refItemNode != nil {
                refItemNode := _refItemNode.(*Nodes_ExpRefItemNode)
                var dstType *Ast_TypeInfo
                dstType = refItemNode.FP.Get_val().FP.Get_expType()
                if _switch26452 := dstType.FP.Get_kind(); _switch26452 == Ast_TypeInfoKind__Map {
                    self.FP.Write("lns_mtd_Map_add( _pEnv, ")
                    self.FP.ProcessVal2any(refItemNode.FP.Get_val(), _var)
                    self.FP.Write(", ")
                    {
                        _indexNode := refItemNode.FP.Get_index()
                        if _indexNode != nil {
                            indexNode := _indexNode.(*Nodes_Node)
                            self.FP.ProcessVal2stem(indexNode, _var)
                        } else {
                            self.FP.Write(convCC_getLiteralStrStem_1500_(Lns_getVM().String_format("\"%s\"", []LnsAny{Lns_unwrap( refItemNode.FP.Get_symbol()).(string)})))
                        }
                    }
                    self.FP.Write(", ")
                    self.FP.processValForSetOp(parent, convCC_ValKind__Stem, dstType.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), exp, index, firstMRet)
                    self.FP.Write(")")
                } else if _switch26452 == Ast_TypeInfoKind__List {
                    self.FP.Write("void lns_mtd_List_setAt( _pEnv, ")
                    self.FP.ProcessVal2any(refItemNode.FP.Get_val(), _var)
                    self.FP.Write(", ")
                    {
                        _indexNode := refItemNode.FP.Get_index()
                        if _indexNode != nil {
                            indexNode := _indexNode.(*Nodes_Node)
                            self.FP.ProcessVal2stem(indexNode, _var)
                        } else {
                            self.FP.Write(convCC_getLiteralStrStem_1500_(Lns_getVM().String_format("\"%s\"", []LnsAny{Lns_unwrap( refItemNode.FP.Get_symbol()).(string)})))
                        }
                    }
                    self.FP.Write(", ")
                    self.FP.processValForSetOp(parent, convCC_ValKind__Stem, dstType.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), exp, index, firstMRet)
                    self.FP.Write(")")
                } else {
                    Util_err(Lns_getVM().String_format("not support type -- %s", []LnsAny{Ast_TypeInfoKind_getTxt( dstType.FP.Get_kind())}))
                }
            }
        }
    } else {
        Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{Nodes_getNodeKindName(_var.FP.Get_kind())}))
    }
}

// 5964: decl @lune.@base.@convCC.convFilter.processSetValToDst
func (self *convCC_convFilter) processSetValToDst(parent *Nodes_Node,dstList *LnsList,expList *LnsList,mRetExp LnsAny) {
    var mRetIndex LnsAny
    mRetIndex = Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(mRetExp) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_MRetExp).FP.Get_index()}))
    for _index, _exp := range( expList.Items ) {
        index := _index + 1
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        var is1stMRet bool
        is1stMRet = index == mRetIndex
        if is1stMRet{
            if mRetExp != nil{
                mRetExp_8201 := mRetExp.(*Nodes_MRetExp)
                self.FP.Write("lns_setMRet( _pEnv, ")
                convCC_filter_1642_(mRetExp_8201.FP.Get_exp(), self, parent)
                self.FP.Write(convCC_accessAny)
                self.FP.Writeln(");")
            }
        }
        if index > dstList.Len(){
            return 
        }
        if index == expList.Len(){
            {
                var _from26692 LnsInt = index
                var _to26692 LnsInt = dstList.Len()
                for _work26692 := _from26692; _work26692 <= _to26692; _work26692++ {
                    dstIndex := _work26692
                    var accessIndex LnsInt
                    if mRetIndex != nil{
                        mRetIndex_8208 := mRetIndex.(LnsInt)
                        accessIndex = index - mRetIndex_8208
                        
                    } else {
                        accessIndex = 0
                        
                    }
                    switch _exp26682 := dstList.GetAt(dstIndex).(type) {
                    case *convCC_DstInfo__Symbol:
                    symbolInfo := _exp26682.Val1
                    dstNode := _exp26682.Val2
                    initFlag := _exp26682.Val3
                        self.FP.processSetValSingle(parent, dstNode, symbolInfo, initFlag, exp, accessIndex, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( is1stMRet) &&
                            Lns_GetEnv().SetStackVal( dstIndex == index) ).(bool))
                    case *convCC_DstInfo__Node:
                    dstNode := _exp26682.Val1
                    initFlag := _exp26682.Val2
                        self.FP.processSetValSingleNode(parent, dstNode, initFlag, exp, accessIndex, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                            Lns_GetEnv().SetStackVal( is1stMRet) &&
                            Lns_GetEnv().SetStackVal( dstIndex == index) ).(bool))
                    }
                    self.FP.Writeln("")
                }
            }
        } else { 
            var accessIndex LnsInt
            if mRetIndex != nil{
                mRetIndex_8220 := mRetIndex.(LnsInt)
                accessIndex = index - mRetIndex_8220
                
            } else {
                accessIndex = 0
                
            }
            switch _exp26764 := dstList.GetAt(index).(type) {
            case *convCC_DstInfo__Symbol:
            symbolInfo := _exp26764.Val1
            dstNode := _exp26764.Val2
            initFlag := _exp26764.Val3
                self.FP.processSetValSingle(parent, dstNode, symbolInfo, initFlag, exp, accessIndex, is1stMRet)
            case *convCC_DstInfo__Node:
            dstNode := _exp26764.Val1
            initFlag := _exp26764.Val2
                self.FP.processSetValSingleNode(parent, dstNode, initFlag, exp, accessIndex, is1stMRet)
            }
            self.FP.Writeln("")
        }
    }
}

// 6049: decl @lune.@base.@convCC.convFilter.processSetValToSym
func (self *convCC_convFilter) processSetValToSym(parent *Nodes_Node,varSymList *LnsList,initFlag bool,expList *LnsList,varNode LnsAny,mRetExp LnsAny) {
    var varNodeList *LnsList
    {
        _expListNode := Nodes_ExpListNodeDownCastF(varNode)
        if _expListNode != nil {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            varNodeList = expListNode.FP.Get_expList()
            
        } else {
            if varNode != nil{
                varNode_8244 := varNode.(*Nodes_Node)
                varNodeList = NewLnsList([]LnsAny{Nodes_Node2Stem(varNode_8244)})
                
            } else {
                varNodeList = NewLnsList([]LnsAny{})
                
            }
        }
    }
    var dstList *LnsList
    dstList = NewLnsList([]LnsAny{})
    for _index, _symbol := range( varSymList.Items ) {
        index := _index + 1
        symbol := _symbol.(Ast_LowSymbol)
        var node LnsAny
        if index <= varNodeList.Len(){
            node = varNodeList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
            
        } else { 
            node = nil
            
        }
        dstList.Insert(&convCC_DstInfo__Symbol{symbol, node, initFlag})
    }
    self.FP.processSetValToDst(parent, dstList, expList, mRetExp)
}

// 6091: decl @lune.@base.@convCC.convFilter.processSetValToNode
func (self *convCC_convFilter) processSetValToNode(parent *Nodes_Node,dstNode *Nodes_Node,initSymSet *LnsSet,expList *LnsList,mRetExp LnsAny) {
    var isInitSym func(node *Nodes_Node) bool
    isInitSym = func(node *Nodes_Node) bool {
        var symbolList *LnsList
        symbolList = node.FP.GetSymbolInfo()
        if symbolList.Len() > 0{
            return initSymSet.Has(Ast_SymbolInfo2Stem(symbolList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()))
        }
        return false
    }
    var dstList *LnsList
    dstList = NewLnsList([]LnsAny{})
    {
        _expListNode := Nodes_ExpListNodeDownCastF(dstNode.FP)
        if _expListNode != nil {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            for _, _node := range( expListNode.FP.Get_expList().Items ) {
                node := _node.(Nodes_NodeDownCast).ToNodes_Node()
                dstList.Insert(&convCC_DstInfo__Node{node, isInitSym(node)})
            }
        } else {
            dstList.Insert(&convCC_DstInfo__Node{dstNode, isInitSym(dstNode)})
        }
    }
    self.FP.processSetValToDst(parent, dstList, expList, mRetExp)
}

// 6121: decl @lune.@base.@convCC.convFilter.processDeclVarAndSet
func (self *convCC_convFilter) processDeclVarAndSet(varSymList *LnsList,expListNode LnsAny) {
    for _, __var := range( varSymList.Items ) {
        _var := __var.(Ast_LowSymbol)
        var symbolParam *convCC_SymbolParam
        symbolParam = self.scopeMgr.FP.GetSymbolParam(_var)
        var typeTxt string
        var valKind LnsInt
        typeTxt,valKind = symbolParam.TypeTxt, symbolParam.Kind
        var declVarFlag bool
        if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( varSymList.GetAt(1).(Ast_LowSymbol).Get_scope() != self.ast.FP.Get_moduleScope()) ||
            Lns_GetEnv().SetStackVal( symbolParam.Index == convCC_invalidSymbolId) ).(bool){
            declVarFlag = true
            
        } else { 
            declVarFlag = false
            
        }
        if valKind != convCC_ValKind__Prim{
            self.FP.processDeclVarC(declVarFlag, _var, valKind != convCC_ValKind__Var, nil)
        } else { 
            if declVarFlag{
                self.FP.Writeln(Lns_getVM().String_format("%s %s;", []LnsAny{typeTxt, self.moduleCtrl.FP.GetSymbolName(_var)}))
            }
        }
    }
    if expListNode != nil{
        expListNode_8292 := expListNode.(*Nodes_ExpListNode)
        self.FP.processSetValToSym(&expListNode_8292.Nodes_Node, varSymList, true, expListNode_8292.FP.Get_expList(), nil, expListNode_8292.FP.Get_mRetExp())
    }
}

// 6154: decl @lune.@base.@convCC.convFilter.processIfUnwrap
func (self *convCC_convFilter) ProcessIfUnwrap(node *Nodes_IfUnwrapNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    var expListNode *Nodes_ExpListNode
    expListNode = node.FP.Get_expList()
    var workSymList *LnsList
    workSymList = NewLnsList([]LnsAny{})
    for _, _varSym := range( node.FP.Get_varSymList().Items ) {
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        var workSymbol *convCC_WorkSymbol
        workSymbol = NewconvCC_WorkSymbol(varSym.FP.Get_scope(), varSym.FP.Get_accessMode(), Lns_getVM().String_format("_%s", []LnsAny{varSym.FP.Get_name()}), varSym.FP.Get_typeInfo().FP.Get_nilableTypeInfo(), varSym.FP.Get_kind(), varSym.FP.Get_staticFlag(), NewconvCC_SymbolParam(convCC_ValKind__Stem, 0, convCC_cTypeStem))
        workSymList.Insert(convCC_WorkSymbol2Stem(workSymbol))
    }
    self.FP.processDeclVarAndSet(workSymList, expListNode)
    self.FP.Write("if ( ")
    for _index, _workSym := range( workSymList.Items ) {
        index := _index + 1
        workSym := _workSym.(convCC_WorkSymbolDownCast).ToconvCC_WorkSymbol()
        self.FP.Write(Lns_getVM().String_format("%s.type != lns_stem_type_nil", []LnsAny{self.moduleCtrl.FP.GetSymbolName(workSym.FP)}))
        if index != workSymList.Len(){
            self.FP.Write(" && ")
        }
    }
    self.FP.Writeln(") {")
    self.FP.processBlockPreProcess(node.FP.Get_block().FP.Get_scope())
    for _index, _varSym := range( node.FP.Get_varSymList().Items ) {
        index := _index + 1
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        self.FP.processDeclVarC(true, varSym.FP, false, nil)
        self.FP.processSetSymSingle(&node.Nodes_Node, nil, varSym.FP, true, workSymList.GetAt(index).(convCC_WorkSymbolDownCast).ToconvCC_WorkSymbol().FP, false)
        self.FP.Writeln("")
    }
    convCC_filter_1642_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.processBlockPostProcess()
    self.FP.Writeln("}")
    {
        __exp := node.FP.Get_nilBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Writeln("else {")
            convCC_filter_1642_(&_exp.Nodes_Node, self, &node.Nodes_Node)
            self.FP.Writeln("}")
        }
    }
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 6219: decl @lune.@base.@convCC.convFilter.processDeclVar
func (self *convCC_convFilter) ProcessDeclVar(node *Nodes_DeclVarNode,_opt LnsAny) {
    if _switch27766 := self.processMode; _switch27766 == convCC_ProcessMode__WideScopeVer {
        if node.FP.Get_mode() == Nodes_DeclVarMode__Let{
            var varSymList *LnsList
            varSymList = node.FP.Get_symbolInfoList()
            if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( varSymList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_scope() == self.moduleTypeInfo.FP.Get_scope()) ||
                Lns_GetEnv().SetStackVal( self.outputBuiltinFlag) ).(bool){
                var process func(out2HMode LnsInt,_var *Ast_SymbolInfo)
                process = func(out2HMode LnsInt,_var *Ast_SymbolInfo) {
                    var typeTxt string
                    typeTxt = convCC_convExp27616(Lns_2DDD(self.scopeMgr.FP.GetCTypeForSym(_var.FP)))
                    self.FP.Write(convCC_getOut2HeaderPrefix_1518_(out2HMode))
                    self.FP.Writeln(Lns_getVM().String_format("%s %s;", []LnsAny{typeTxt, self.moduleCtrl.FP.GetSymbolName(_var.FP)}))
                }
                for _, __var := range( varSymList.Items ) {
                    _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    {
                        var processwork func(out2HMode LnsInt)
                        processwork = func(out2HMode LnsInt) {
                            process(out2HMode, _var)
                        }
                        if Ast_isPubToExternal(_var.FP.Get_accessMode()){
                            self.stream.FP.SwitchToHeader()
                            processwork(convCC_Out2HMode__HeaderPub)
                            self.stream.FP.ReturnToSource()
                            processwork(convCC_Out2HMode__SourcePub)
                        } else { 
                            processwork(convCC_Out2HMode__SourcePri)
                        }
                    }
                    
                }
            }
        }
        return 
    } else if _switch27766 == convCC_ProcessMode__InitModule || _switch27766 == convCC_ProcessMode__Form {
    } else {
        return 
    }
    if Lns_isCondTrue( node.FP.Get_syncBlock()){
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        for _, _varInfo := range( node.FP.Get_syncVarList().Items ) {
            varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            self.FP.Writeln(Lns_getVM().String_format("_sync_%s", []LnsAny{varInfo.FP.Get_name()}))
        }
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
    }
    var varSymList *LnsList
    varSymList = node.FP.Get_symbolInfoList()
    if node.FP.Get_unwrapFlag(){
        var unwrapScope *Ast_Scope
        unwrapScope = Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(node.FP.Get_unwrapBlock()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_BlockNode).FP.Get_scope()}))).(*Ast_Scope)
        var workSymList *LnsList
        workSymList = NewLnsList([]LnsAny{})
        for _, _varSym := range( varSymList.Items ) {
            varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            var tmpVarSym *Ast_SymbolInfo
            tmpVarSym = Lns_unwrap( unwrapScope.FP.GetSymbolInfoChild(Lns_getVM().String_format("_%s", []LnsAny{varSym.FP.Get_name()}))).(*Ast_SymbolInfo)
            workSymList.Insert(Ast_SymbolInfo2Stem(tmpVarSym))
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( node.FP.Get_mode() == Nodes_DeclVarMode__Let) &&
                Lns_GetEnv().SetStackVal( varSym.FP.Get_scope() != self.moduleTypeInfo.FP.Get_scope()) ).(bool)){
                self.FP.processDeclVarC(true, varSym.FP, false, nil)
            }
        }
        self.FP.Writeln("{")
        self.FP.processBlockPreProcess(unwrapScope)
        self.FP.processDeclVarAndSet(workSymList, node.FP.Get_expList())
        {
            _unwrapBlock := node.FP.Get_unwrapBlock()
            if _unwrapBlock != nil {
                unwrapBlock := _unwrapBlock.(*Nodes_BlockNode)
                self.FP.Writeln("")
                self.FP.Write("if ( ")
                var firstFlag bool
                firstFlag = true
                for _, __var := range( workSymList.Items ) {
                    _var := __var.(Ast_LowSymbol)
                    if self.scopeMgr.FP.GetSymbolValKind(_var) == convCC_ValKind__Stem{
                        if _var.Get_typeInfo().FP.Get_nilable(){
                            if firstFlag{
                                firstFlag = false
                                
                            } else { 
                                self.FP.Write(" || ")
                            }
                            self.FP.Write(Lns_getVM().String_format("lns_stem_type_nil == %s.type", []LnsAny{self.moduleCtrl.FP.GetSymbolName(_var)}))
                        }
                    }
                }
                self.FP.Writeln(" ) {")
                self.FP.PushIndent(nil)
                convCC_filter_1642_(&unwrapBlock.Nodes_Node, self, &node.Nodes_Node)
                self.FP.PopIndent()
                {
                    _thenBlock := node.FP.Get_thenBlock()
                    if _thenBlock != nil {
                        thenBlock := _thenBlock.(*Nodes_BlockNode)
                        self.FP.Writeln("}")
                        self.FP.Writeln("else {")
                        self.FP.PushIndent(nil)
                        for _index, __var := range( varSymList.Items ) {
                            index := _index + 1
                            _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            self.FP.processSetSymSingle(&node.Nodes_Node, nil, _var.FP, true, workSymList.GetAt(index).(Ast_LowSymbol), false)
                            self.FP.Writeln("")
                        }
                        self.FP.PopIndent()
                        convCC_filter_1642_(&thenBlock.Nodes_Node, self, &node.Nodes_Node)
                        self.FP.Writeln("}")
                    } else {
                        self.FP.Writeln("}")
                        self.FP.Writeln("else {")
                        self.FP.PushIndent(nil)
                        for _index, __var := range( varSymList.Items ) {
                            index := _index + 1
                            _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            self.FP.processSetSymSingle(&node.Nodes_Node, nil, _var.FP, true, workSymList.GetAt(index).(Ast_LowSymbol), false)
                            self.FP.Writeln("")
                        }
                        self.FP.PopIndent()
                        self.FP.Writeln("}")
                    }
                }
                self.FP.processBlockPostProcess()
                self.FP.Writeln("}")
            }
        }
    } else { 
        self.FP.processDeclVarAndSet(varSymList, node.FP.Get_expList())
    }
    {
        __exp := node.FP.Get_syncBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            convCC_filter_1642_(&_exp.Nodes_Node, self, &node.Nodes_Node)
            for _, _varInfo := range( node.FP.Get_syncVarList().Items ) {
                varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.Writeln(Lns_getVM().String_format("_sync_%s = %s", []LnsAny{varInfo.FP.Get_name(), varInfo.FP.Get_name()}))
            }
            self.FP.PopIndent()
            self.FP.Writeln("}")
            for _, _varInfo := range( node.FP.Get_syncVarList().Items ) {
                varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.Writeln(Lns_getVM().String_format("%s = _sync_%s", []LnsAny{varInfo.FP.Get_name(), varInfo.FP.Get_name()}))
            }
            self.FP.PopIndent()
            self.FP.Writeln("}")
        }
    }
    if node.FP.Get_accessMode() == Ast_AccessMode__Pub{
        self.FP.Writeln("")
        for _index, __var := range( varSymList.Items ) {
            index := _index + 1
            _var := __var.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            var name string
            name = self.moduleCtrl.FP.GetSymbolName(_var.FP)
            self.pubVarName2InfoMap.Set(name,NewconvCC_PubVarInfo(node.FP.Get_staticFlag(), node.FP.Get_accessMode(), node.FP.Get_symbolInfoList().GetAt(index).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_mutable(), node.FP.Get_typeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))
        }
    }
}

// 6389: decl @lune.@base.@convCC.convFilter.processWhen
func (self *convCC_convFilter) ProcessWhen(node *Nodes_WhenNode,_opt LnsAny) {
    self.FP.Write("if ( ")
    for _index, _symPair := range( node.FP.Get_symPairList().Items ) {
        index := _index + 1
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        self.FP.processSym2stem(symPair.FP.Get_src().FP)
        self.FP.Write(".type != lns_stem_type_nil")
        if index != node.FP.Get_symPairList().Len(){
            self.FP.Write(" && ")
        }
    }
    self.FP.Writeln(" ) ")
    self.FP.Writeln("{")
    self.FP.processBlockPreProcess(node.FP.Get_block().FP.Get_scope())
    for _, _symPair := range( node.FP.Get_symPairList().Items ) {
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        var srcSymbol *Ast_SymbolInfo
        srcSymbol = symPair.FP.Get_src()
        var dstSymbol *Ast_SymbolInfo
        dstSymbol = symPair.FP.Get_dst()
        var srcTypeTxt string
        srcTypeTxt = convCC_convExp28556(Lns_2DDD(self.scopeMgr.FP.GetCTypeForSym(srcSymbol.FP)))
        var dstTypeTxt string
        dstTypeTxt = convCC_convExp28570(Lns_2DDD(self.scopeMgr.FP.GetCTypeForSym(dstSymbol.FP)))
        if srcTypeTxt != dstTypeTxt{
            self.FP.processDeclVarC(true, dstSymbol.FP, false, nil)
            self.FP.processSetSymSingle(&node.Nodes_Node, nil, dstSymbol.FP, true, srcSymbol.FP, false)
            self.FP.Writeln("")
        } else { 
            self.FP.Writeln(Lns_getVM().String_format("%s %s = %s;", []LnsAny{dstTypeTxt, self.moduleCtrl.FP.GetSymbolName(dstSymbol.FP), self.moduleCtrl.FP.GetSymbolName(srcSymbol.FP)}))
        }
    }
    convCC_filter_1642_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.processBlockPostProcess()
    {
        __exp := node.FP.Get_elseBlock()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Write("} else {")
            convCC_filter_1642_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln("}")
}

// 6449: decl @lune.@base.@convCC.convFilter.processDeclArg
func (self *convCC_convFilter) ProcessDeclArg(node *Nodes_DeclArgNode,_opt LnsAny) {
    convCC_processDeclAlgeSub_1993_(self.stream.FP, node)
}

// 6475: decl @lune.@base.@convCC.convFilter.processDeclArgDDD
func (self *convCC_convFilter) ProcessDeclArgDDD(node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("%s _pDDD", []LnsAny{convCC_cTypeStem}))
}

// 6481: decl @lune.@base.@convCC.convFilter.processExpSubDDD
func (self *convCC_convFilter) ProcessExpSubDDD(node *Nodes_ExpSubDDDNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("lns_createSubDDD( _pEnv, %d, LNS_STEM_ANY( _pEnv->pMRet ) )", []LnsAny{node.FP.Get_remainIndex()}))
}

// 6488: decl @lune.@base.@convCC.convFilter.processFuncPrototype
func (self *convCC_convFilter) processFuncPrototype(parent *Nodes_Node,accessMode LnsInt,needFormVal bool,name string,retType string,argList *LnsList,termFlag bool) {
    var process func(out2HMode LnsInt)
    process = func(out2HMode LnsInt) {
        self.FP.Write(Lns_getVM().String_format("%s%s %s( %s _pEnv", []LnsAny{convCC_getOut2HeaderPrefix_1518_(out2HMode), retType, name, convCC_cTypeEnvP}))
        if needFormVal{
            self.FP.Write(Lns_getVM().String_format(", %s _pForm", []LnsAny{convCC_cTypeAnyP}))
        }
        for _, _arg := range( argList.Items ) {
            arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
            self.FP.Write(", ")
            convCC_filter_1642_(arg, self, parent)
        }
        self.FP.Write(" )")
        if termFlag{
            self.FP.Writeln(";")
        }
    }
    if termFlag{
        {
            var processwork func(out2HMode LnsInt)
            processwork = func(out2HMode LnsInt) {
                if out2HMode != convCC_Out2HMode__SourcePub{
                    process(out2HMode)
                }
            }
            if Ast_isPubToExternal(accessMode){
                self.stream.FP.SwitchToHeader()
                processwork(convCC_Out2HMode__HeaderPub)
                self.stream.FP.ReturnToSource()
                processwork(convCC_Out2HMode__SourcePub)
            } else { 
                processwork(convCC_Out2HMode__SourcePri)
            }
        }
        
    } else { 
        process(convCC_Out2HMode__SourcePub)
    }
}

// 6524: decl @lune.@base.@convCC.convFilter.processCallUserForm
func (self *convCC_convFilter) processCallUserForm(formName string,formType *Ast_TypeInfo,argNameList *LnsList) {
    var process func(prefix string)
    process = func(prefix string) {
        self.FP.PushIndent(nil)
        self.FP.Write(prefix)
        for _, _argName := range( argNameList.Items ) {
            argName := _argName.(string)
            self.FP.Write(", ")
            self.FP.Write(argName)
        }
        self.FP.Writeln(");")
        self.FP.PopIndent()
    }
    self.FP.Writeln(Lns_getVM().String_format("if lns_isClosure( %s ) {", []LnsAny{formName}))
    if _switch29248 := convCC_getCRetType_1195_(formType.FP.Get_retTypeInfoList()); _switch29248 == "void" {
        process(Lns_getVM().String_format("lns_closure( %s )( _pEnv, %s", []LnsAny{formName, formName}))
    } else if _switch29248 == convCC_cTypeAny {
        process(Lns_getVM().String_format("return lns_closure_any( %s )( _pEnv, %s", []LnsAny{formName, formName}))
    } else if _switch29248 == convCC_cTypeInt {
        process(Lns_getVM().String_format("return lns_closure_int( %s )( _pEnv, %s", []LnsAny{formName, formName}))
    } else if _switch29248 == convCC_cTypeReal {
        process(Lns_getVM().String_format("return lns_closure_real( %s )( _pEnv, %s", []LnsAny{formName, formName}))
    } else if _switch29248 == convCC_cTypeBool {
        process(Lns_getVM().String_format("return lns_closure_bool( %s )( _pEnv, %s", []LnsAny{formName, formName}))
    } else {
        process(Lns_getVM().String_format("return lns_closure( %s )( _pEnv, %s", []LnsAny{formName, formName}))
    }
    self.FP.Writeln("}")
    self.FP.Writeln("else {")
    if _switch29371 := convCC_getCRetType_1195_(formType.FP.Get_retTypeInfoList()); _switch29371 == "void" {
        process(Lns_getVM().String_format("lns_func( %s )( _pEnv", []LnsAny{formName}))
    } else if _switch29371 == convCC_cTypeAnyP {
        process(Lns_getVM().String_format("return lns_func_any( %s )( _pEnv", []LnsAny{formName}))
    } else if _switch29371 == convCC_cTypeInt {
        process(Lns_getVM().String_format("return lns_func_int( %s )( _pEnv", []LnsAny{formName}))
    } else if _switch29371 == convCC_cTypeReal {
        process(Lns_getVM().String_format("return lns_func_real( %s )( _pEnv", []LnsAny{formName}))
    } else if _switch29371 == convCC_cTypeBool {
        process(Lns_getVM().String_format("return lns_func_bool( %s )( _pEnv", []LnsAny{formName}))
    } else {
        process(Lns_getVM().String_format("return lns_func( %s )( _pEnv", []LnsAny{formName}))
    }
    self.FP.Writeln("}")
}

// 6584: decl @lune.@base.@convCC.convFilter.processDeclForm
func (self *convCC_convFilter) ProcessDeclForm(node *Nodes_DeclFormNode,_opt LnsAny) {
    var formType *Ast_TypeInfo
    formType = node.FP.Get_expType()
    if _switch29575 := self.processMode; _switch29575 == convCC_ProcessMode__Prototype {
        self.FP.processFuncPrototype(&node.Nodes_Node, formType.FP.Get_accessMode(), true, self.moduleCtrl.FP.GetCallFormName(formType), convCC_getCRetType_1195_(formType.FP.Get_retTypeInfoList()), node.FP.Get_argList(), true)
    } else if _switch29575 == convCC_ProcessMode__Form {
        self.FP.processFuncPrototype(&node.Nodes_Node, formType.FP.Get_accessMode(), true, self.moduleCtrl.FP.GetCallFormName(formType), convCC_getCRetType_1195_(formType.FP.Get_retTypeInfoList()), node.FP.Get_argList(), false)
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        var argNameList *LnsList
        argNameList = NewLnsList([]LnsAny{})
        for _, _arg := range( node.FP.Get_argList().Items ) {
            arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
            {
                _workArg := Nodes_DeclArgNodeDownCastF(arg.FP)
                if _workArg != nil {
                    workArg := _workArg.(*Nodes_DeclArgNode)
                    argNameList.Insert(workArg.FP.Get_name().Txt)
                } else {
                    if Lns_isCondTrue( Nodes_DeclArgDDDNodeDownCastF(arg.FP)){
                        argNameList.Insert("_pDDD")
                    }
                }
            }
        }
        self.FP.processCallUserForm("_pForm", formType, argNameList)
        self.FP.PopIndent()
        self.FP.Writeln("}")
    }
}

// 6694: decl @lune.@base.@convCC.convFilter.processClosureFunc
func (self *convCC_convFilter) processClosureFunc(declInfo *Nodes_DeclFuncInfo) {
    var simpleName *Types_Token
    
    {
        _simpleName := declInfo.FP.Get_name()
        if _simpleName == nil{
            return 
        } else {
            simpleName = _simpleName.(*Types_Token)
        }
    }
    var scope *Ast_Scope
    
    {
        _scope := Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(declInfo.FP.Get_body()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Nodes_BlockNode).FP.Get_scope()}))
        if _scope == nil{
            return 
        } else {
            scope = _scope.(*Ast_Scope)
        }
    }
    var funcSym *Ast_SymbolInfo
    
    {
        _funcSym := scope.FP.GetSymbolInfo(simpleName.Txt, scope, false, Ast_ScopeAccess__Normal)
        if _funcSym == nil{
            return 
        } else {
            funcSym = _funcSym.(*Ast_SymbolInfo)
        }
    }
    if Lns_op_not(funcSym.FP.Get_hasAccessFromClosure()){
        return 
    }
    var symbolParam *convCC_SymbolParam
    symbolParam = self.scopeMgr.FP.GetSymbolParam(funcSym.FP)
    var symbolName string
    symbolName = self.moduleCtrl.FP.GetSymbolName(funcSym.FP)
    self.FP.Writeln(Lns_getVM().String_format("%s %s;", []LnsAny{convCC_cTypeVarP, symbolName}))
    self.FP.Write(Lns_getVM().String_format("lns_initVal_var( %s, %s, %d, ", []LnsAny{symbolName, convCC_getBlockName_1198_(scope.FP.Get_parent()), symbolParam.Index}))
    self.FP.Writeln(Lns_getVM().String_format("LNS_STEM_ANY( %s ) );", []LnsAny{convCC_getFunc2any_1897_(self.moduleCtrl, self.scopeMgr, funcSym.FP.Get_typeInfo())}))
}

// 6717: decl @lune.@base.@convCC.convFilter.processDeclFunc
func (self *convCC_convFilter) ProcessDeclFunc(node *Nodes_DeclFuncNode,_opt LnsAny) {
    opt := _opt.(*ConvCC_Opt)
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo()
    var name string
    name = self.moduleCtrl.FP.GetFuncName(node.FP.Get_expType())
    var processFuncPrototype func(termFlag bool)
    processFuncPrototype = func(termFlag bool) {
        self.FP.processFuncPrototype(&node.Nodes_Node, declInfo.FP.Get_accessMode(), convCC_isClosure_1011_(node.FP.Get_expType()), name, convCC_getCRetType_1195_(node.FP.Get_expType().FP.Get_retTypeInfoList()), declInfo.FP.Get_argList(), termFlag)
    }
    if _switch29868 := self.processMode; _switch29868 == convCC_ProcessMode__Form {
    } else if _switch29868 == convCC_ProcessMode__Prototype {
        processFuncPrototype(true)
        return 
    } else if _switch29868 == convCC_ProcessMode__InitFuncSym || _switch29868 == convCC_ProcessMode__WideScopeVer {
        self.FP.process__func__symbol(node.FP.Get_expType(), declInfo.FP.Get_has__func__Symbol(), self.moduleCtrl.FP.GetFuncName(node.FP.Get_expType()))
        return 
    } else {
        if _switch29865 := opt.Node.FP.Get_kind(); _switch29865 == Nodes_NodeKind_get_Block() || _switch29865 == Nodes_NodeKind_get_ExpMacroExp() {
        } else {
            self.FP.Write(convCC_getFunc2any_1897_(self.moduleCtrl, self.scopeMgr, node.FP.Get_expType()))
        }
        return 
    }
    if self.duringDeclFunc{
        if opt.Node.FP.Get_kind() == Nodes_NodeKind_get_Block(){
            self.FP.processClosureFunc(declInfo)
            return 
        }
        self.FP.Write(self.moduleCtrl.FP.GetFuncName(node.FP.Get_expType()))
        return 
    }
    var body *Nodes_BlockNode
    
    {
        _body := declInfo.FP.Get_body()
        if _body == nil{
            return 
        } else {
            body = _body.(*Nodes_BlockNode)
        }
    }
    self.duringDeclFunc = true
    
    processFuncPrototype(false)
    self.FP.Writeln("")
    self.FP.Writeln("{")
    self.FP.pushRoutine(node.FP.Get_expType(), body)
    self.FP.processArgClosure(declInfo)
    var breakKind LnsInt
    breakKind = Nodes_BreakKind__None
    convCC_filter_1642_(&body.Nodes_Node, self, &node.Nodes_Node)
    self.FP.popRoutine()
    breakKind = body.FP.GetBreakKind(Nodes_CheckBreakMode__Normal)
    
    if _switch30026 := breakKind; _switch30026 == Nodes_BreakKind__Return || _switch30026 == Nodes_BreakKind__NeverRet {
    } else {
    }
    self.FP.Writeln("}")
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType()
    if expType.FP.Get_accessMode() == Ast_AccessMode__Pub{
        self.pubFuncName2InfoMap.Set(name,NewconvCC_PubFuncInfo(declInfo.FP.Get_accessMode(), node.FP.Get_expType()))
    }
}

// 6814: decl @lune.@base.@convCC.convFilter.processRefType
func (self *convCC_convFilter) ProcessRefType(node *Nodes_RefTypeNode,_opt LnsAny) {
}

// 6828: decl @lune.@base.@convCC.convFilter.processIf
func (self *convCC_convFilter) ProcessIf(node *Nodes_IfNode,_opt LnsAny) {
    var valList *LnsList
    valList = node.FP.Get_stmtList()
    for _index, _val := range( valList.Items ) {
        index := _index + 1
        val := _val.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if index == 1{
            self.FP.Write("if ( lns_isCondTrue( ")
            self.FP.ProcessVal2stem(val.FP.Get_exp(), &node.Nodes_Node)
            self.FP.Write(") )")
        } else if val.FP.Get_kind() == Nodes_IfKind__ElseIf{
            self.FP.Write("else if ( lns_isCondTrue( ")
            self.FP.ProcessVal2stem(val.FP.Get_exp(), &node.Nodes_Node)
            self.FP.Write(") )")
        } else { 
            self.FP.Writeln("else {")
        }
        self.FP.Write(" ")
        convCC_filter_1642_(&val.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
        self.FP.Write("}")
    }
}

// 6855: decl @lune.@base.@convCC.convFilter.processEquals
func (self *convCC_convFilter) processEquals(eqFlag bool,type1 *Ast_TypeInfo,type2 *Ast_TypeInfo,process1 convCC_ProcessToValForm_2538_,process2 convCC_ProcessToValForm_2538_) {
    var valKind1 LnsInt
    valKind1 = convCC_getValKind_1167_(type1)
    var valKind2 LnsInt
    valKind2 = convCC_getValKind_1167_(type2)
    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( valKind1 == convCC_ValKind__Stem) ||
        Lns_GetEnv().SetStackVal( valKind2 == convCC_ValKind__Stem) ).(bool){
        if Lns_op_not(eqFlag){
            self.FP.Write("!")
        }
        self.FP.Write("lns_equals( ")
        process1(convCC_ValKind__Stem)
        self.FP.Write(",")
        process2(convCC_ValKind__Stem)
        self.FP.Write(")")
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( valKind1 == convCC_ValKind__Any) &&
        Lns_GetEnv().SetStackVal( valKind2 == convCC_ValKind__Any) ).(bool)){
        if Lns_op_not(eqFlag){
            self.FP.Write("!")
        }
        self.FP.Write("lns_equals_any( ")
        process1(convCC_ValKind__Any)
        self.FP.Write(",")
        process2(convCC_ValKind__Any)
        self.FP.Write(")")
    } else if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( valKind1 == convCC_ValKind__Prim) &&
        Lns_GetEnv().SetStackVal( valKind2 == convCC_ValKind__Prim) ).(bool)){
        process1(convCC_ValKind__Prim)
        if eqFlag{
            self.FP.Write(" == ")
        } else { 
            self.FP.Write(" != ")
        }
        process2(convCC_ValKind__Prim)
    } else { 
        Util_err(Lns_getVM().String_format("illegal kind %s -- %s", []LnsAny{convCC_ValKind_getTxt( valKind1), convCC_ValKind_getTxt( valKind2)}))
    }
}

// 6898: decl @lune.@base.@convCC.convFilter.processSwitch
func (self *convCC_convFilter) ProcessSwitch(node *Nodes_SwitchNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    var expType *Ast_TypeInfo
    expType = node.FP.Get_exp().FP.Get_expType()
    var expSymName string
    expSymName = Lns_getVM().String_format("_switchExp%d", []LnsAny{node.FP.Get_id()})
    self.FP.Write(Lns_getVM().String_format("%s %s = ", []LnsAny{convCC_getCType_1188_(expType), expSymName}))
    convCC_filter_1642_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Writeln(";")
    var expValKind LnsInt
    expValKind = self.FP.getValKindOfNode(node.FP.Get_exp())
    var expStemName string
    if expValKind != convCC_ValKind__Stem{
        expStemName = expSymName + "Stem"
        
        self.FP.Write(Lns_getVM().String_format("%s %s = ", []LnsAny{convCC_cTypeStem, expStemName}))
        convCC_process2stem_2060_(self.stream.FP, self.moduleCtrl, self.scopeMgr, expValKind, expType, &node.Nodes_Node, convCC_process2stemCallback_2057_(func() {
            self.FP.Write(expSymName)
        }))
        self.FP.Writeln(";")
    } else { 
        expStemName = expSymName
        
    }
    for _index, _caseInfo := range( node.FP.Get_caseList().Items ) {
        index := _index + 1
        caseInfo := _caseInfo.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        if index == 1{
            self.FP.Write("if ( ")
        } else { 
            self.FP.Writeln("}")
            self.FP.Write("else if ( ")
        }
        var expList *Nodes_ExpListNode
        expList = caseInfo.FP.Get_expList()
        for _listIndex, _expNode := range( expList.FP.Get_expList().Items ) {
            listIndex := _listIndex + 1
            expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
            if listIndex != 1{
                self.FP.Write(" || ")
            }
            self.FP.processEquals(true, expType, expNode.FP.Get_expType(), convCC_ProcessToValForm_2538_(func(valKind LnsInt) {
                if valKind == convCC_ValKind__Stem{
                    self.FP.Write(expStemName)
                } else { 
                    self.FP.Write(expSymName)
                }
            }), convCC_ProcessToValForm_2538_(func(valKind LnsInt) {
                if _switch30811 := valKind; _switch30811 == convCC_ValKind__Stem {
                    self.FP.ProcessVal2stem(expNode, &node.Nodes_Node)
                } else if _switch30811 == convCC_ValKind__Any {
                    self.FP.ProcessVal2any(expNode, &node.Nodes_Node)
                } else if _switch30811 == convCC_ValKind__Prim {
                    self.FP.AccessPrimVal(expNode, &node.Nodes_Node)
                }
            }))
        }
        self.FP.Writeln(" ) {")
        convCC_filter_1642_(&caseInfo.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    }
    {
        __exp := node.FP.Get_default()
        if __exp != nil {
            _exp := __exp.(*Nodes_BlockNode)
            self.FP.Writeln("}")
            self.FP.Writeln("else {")
            convCC_filter_1642_(&_exp.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln("}")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 6985: decl @lune.@base.@convCC.convFilter.processMatch
func (self *convCC_convFilter) ProcessMatch(node *Nodes_MatchNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Write(Lns_getVM().String_format("%s _matchExp = ", []LnsAny{convCC_cTypeAnyP}))
    convCC_filter_1642_(node.FP.Get_val(), self, &node.Nodes_Node)
    self.FP.Write(convCC_accessAny)
    self.FP.Writeln(";")
    self.FP.Writeln("switch( _matchExp->val.alge.type ) {")
    var algeType *Ast_AlgeTypeInfo
    algeType = node.FP.Get_algeTypeInfo()
    var enumName string
    enumName = self.moduleCtrl.FP.GetAlgeEnumCName(&algeType.Ast_TypeInfo)
    for _, _caseInfo := range( node.FP.Get_caseList().Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        var valInfo *Ast_AlgeValInfo
        valInfo = caseInfo.FP.Get_valInfo()
        self.FP.Writeln(Lns_getVM().String_format("case %s_%s:", []LnsAny{enumName, valInfo.FP.Get_name()}))
        self.FP.PushIndent(nil)
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        if valInfo.FP.Get_typeList().Len() > 0{
            var structTxt string
            structTxt = self.moduleCtrl.FP.GetAlgeValStrCName(&algeType.Ast_TypeInfo, valInfo.FP.Get_name())
            self.FP.Writeln(Lns_getVM().String_format("%s * _pVal = (%s *)_matchExp->val.alge.pVal;", []LnsAny{structTxt, structTxt}))
            for _paramIndex, _paramType := range( valInfo.FP.Get_typeList().Items ) {
                paramIndex := _paramIndex + 1
                paramType := _paramType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                var paramName *Ast_SymbolInfo
                paramName = caseInfo.FP.Get_valParamNameList().GetAt(paramIndex).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.Writeln(Lns_getVM().String_format("%s %s = _pVal->_val%d;", []LnsAny{convCC_getCType_1188_(paramType), paramName.FP.Get_name(), paramIndex}))
            }
        }
        self.FP.PopIndent()
        convCC_filter_1642_(&caseInfo.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
        self.FP.Writeln("}")
        self.FP.Writeln("break;")
        self.FP.PopIndent()
    }
    self.FP.Writeln("}")
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 7061: decl @lune.@base.@convCC.convFilter.processWhile
func (self *convCC_convFilter) ProcessWhile(node *Nodes_WhileNode,_opt LnsAny) {
    self.FP.processLoopPreProcess(node.FP.Get_block())
    self.FP.Write("while ( ")
    if node.FP.Get_exp().FP.Get_expType().FP.Get_srcTypeInfo() == Ast_builtinTypeBool{
        convCC_filter_1642_(node.FP.Get_exp(), self, &node.Nodes_Node)
    } else { 
        self.FP.Write("lns_isCondTrue( ")
        self.FP.ProcessVal2stem(node.FP.Get_exp(), &node.Nodes_Node)
        self.FP.Write(")")
    }
    self.FP.Writeln(" )")
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Writeln("lns_reset_block( _pEnv );")
    convCC_filter_1642_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.FP.processLoopPostProcess()
}

// 7089: decl @lune.@base.@convCC.convFilter.processRepeat
func (self *convCC_convFilter) ProcessRepeat(node *Nodes_RepeatNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.processLoopPreProcess(node.FP.Get_block())
    self.FP.Writeln("while ( true ) {")
    self.FP.PushIndent(nil)
    self.FP.Writeln("lns_reset_block( _pEnv );")
    convCC_filter_1642_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.Write("if ( ")
    if node.FP.Get_exp().FP.Get_expType().FP.Get_srcTypeInfo() == Ast_builtinTypeBool{
        convCC_filter_1642_(node.FP.Get_exp(), self, &node.Nodes_Node)
    } else { 
        self.FP.Write("lns_isCondTrue(")
        self.FP.ProcessVal2stem(node.FP.Get_exp(), &node.Nodes_Node)
        self.FP.Write(")")
    }
    self.FP.Writeln(") { break; }")
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.FP.processLoopPostProcess()
    self.FP.Writeln("}")
}

// 7120: decl @lune.@base.@convCC.convFilter.processFor
func (self *convCC_convFilter) ProcessFor(node *Nodes_ForNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Writeln(Lns_getVM().String_format("%s _to;", []LnsAny{convCC_cTypeInt}))
    self.FP.Writeln(Lns_getVM().String_format("%s _inc;", []LnsAny{convCC_cTypeInt}))
    self.FP.Writeln(Lns_getVM().String_format("%s %s;", []LnsAny{convCC_cTypeInt, self.moduleCtrl.FP.GetSymbolName(node.FP.Get_val().FP)}))
    self.FP.processSetValSingle(&node.Nodes_Node, nil, node.FP.Get_val().FP, true, node.FP.Get_to(), 0, false)
    self.FP.Writeln("")
    self.FP.Writeln(Lns_getVM().String_format("_to = %s;", []LnsAny{self.moduleCtrl.FP.GetSymbolName(node.FP.Get_val().FP)}))
    {
        __exp := node.FP.Get_delta()
        if __exp != nil {
            _exp := __exp.(*Nodes_Node)
            self.FP.processSetValToSym(&node.Nodes_Node, NewLnsList([]LnsAny{Ast_SymbolInfo2Stem(node.FP.Get_val())}), true, NewLnsList([]LnsAny{Nodes_Node2Stem(_exp)}), nil, nil)
            self.FP.Writeln(Lns_getVM().String_format("_inc = %s;", []LnsAny{self.moduleCtrl.FP.GetSymbolName(node.FP.Get_val().FP)}))
        } else {
            self.FP.Writeln("_inc = 1;")
        }
    }
    self.FP.processSetValToSym(&node.Nodes_Node, NewLnsList([]LnsAny{Ast_SymbolInfo2Stem(node.FP.Get_val())}), true, NewLnsList([]LnsAny{Nodes_Node2Stem(node.FP.Get_init())}), nil, nil)
    self.FP.Writeln("")
    self.FP.processLoopPreProcess(node.FP.Get_block())
    var indexSym string
    indexSym = self.moduleCtrl.FP.GetSymbolName(node.FP.Get_val().FP)
    self.FP.Writeln(Lns_getVM().String_format("for (; ; %s += _inc ) {", []LnsAny{indexSym}))
    self.FP.PushIndent(nil)
    self.FP.Writeln(Lns_getVM().String_format("if ( ( _inc >= 0 && %s > _to ) || ( _inc < 0 && %s < _to ) ) { break; }", []LnsAny{indexSym, indexSym}))
    self.FP.Writeln("lns_reset_block( _pEnv );")
    convCC_filter_1642_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.FP.processLoopPostProcess()
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 7164: decl @lune.@base.@convCC.convFilter.processCreateDDD
func (self *convCC_convFilter) processCreateDDD(parent *Nodes_Node,expList *LnsList) {
    self.FP.Write("lns_createDDD")
    var lastExp *Nodes_Node
    lastExp = expList.GetAt(expList.Len()).(Nodes_NodeDownCast).ToNodes_Node()
    self.FP.Write(Lns_getVM().String_format("( _pEnv, %s, %d", []LnsAny{Nodes_hasMultiValNode(lastExp), expList.Len()}))
    for _, _exp := range( expList.Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.Write(", ")
        self.FP.ProcessVal2stem(exp, parent)
    }
    self.FP.Write(")")
}

// 7178: decl @lune.@base.@convCC.convFilter.processApply
func (self *convCC_convFilter) ProcessApply(node *Nodes_ApplyNode,_opt LnsAny) {
    self.FP.Writeln("{")
    var varList *LnsList
    varList = node.FP.Get_varList()
    var iteExpTypeList *LnsList
    iteExpTypeList = node.FP.Get_expList().FP.Get_expTypeList()
    var iteFuncType *Ast_TypeInfo
    iteFuncType = iteExpTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
    var dummyId LnsInt
    dummyId = varList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_symbolId()
    var dummyScope *Ast_Scope
    dummyScope = NewAst_Scope(self.processInfo, self.moduleTypeInfo.FP.Get_scope(), false, nil, nil)
    var formSym *Ast_SymbolInfo
    formSym = Lns_unwrap( Lns_car(dummyScope.FP.AddLocalVar(self.processInfo, false, false, Lns_getVM().String_format("_form%d", []LnsAny{dummyId}), node.FP.Get_pos(), iteFuncType, Ast_MutMode__IMut))).(*Ast_SymbolInfo)
    var paramSym *Ast_SymbolInfo
    paramSym = Lns_unwrap( Lns_car(dummyScope.FP.AddLocalVar(self.processInfo, false, false, Lns_getVM().String_format("_param%d", []LnsAny{dummyId}), node.FP.Get_pos(), iteExpTypeList.GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut))).(*Ast_SymbolInfo)
    var stateSym *Ast_SymbolInfo
    stateSym = Lns_unwrap( Lns_car(dummyScope.FP.AddLocalVar(self.processInfo, false, false, Lns_getVM().String_format("_state%d", []LnsAny{dummyId}), node.FP.Get_pos(), iteExpTypeList.GetAt(3).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), Ast_MutMode__IMut))).(*Ast_SymbolInfo)
    self.scopeMgr.FP.SetupScopeParam(dummyScope)
    self.FP.processBlockPreProcess(dummyScope)
    var symList *LnsList
    symList = NewLnsList([]LnsAny{})
    symList.Insert(Ast_SymbolInfo2Stem(formSym))
    symList.Insert(Ast_SymbolInfo2Stem(paramSym))
    symList.Insert(Ast_SymbolInfo2Stem(stateSym))
    self.FP.processDeclVarAndSet(symList, node.FP.Get_expList())
    self.FP.Writeln("{")
    self.FP.processLoopPreProcess(node.FP.Get_block())
    self.FP.Writeln("while ( true ) {")
    self.FP.PushIndent(nil)
    self.FP.Writeln("lns_reset_block( _pEnv );")
    for _, _varSym := range( node.FP.Get_varList().Items ) {
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        self.FP.processDeclVarC(true, varSym.FP, false, nil)
    }
    var workSymName string
    workSymName = Lns_getVM().String_format("_workMret%d", []LnsAny{dummyId})
    self.FP.Write(Lns_getVM().String_format("%s %s = ", []LnsAny{convCC_cTypeAnyP, workSymName}))
    if formSym.FP.Get_typeInfo().FP.Get_kind() == Ast_TypeInfoKind__Ext{
        self.FP.Write(Lns_getVM().String_format("lns_lua_callForm( _pEnv, *%s, ", []LnsAny{self.moduleCtrl.FP.GetSymbolName(formSym.FP)}))
        self.FP.Write("")
        var expList *LnsList
        expList = NewLnsList([]LnsAny{})
        expList.Insert(Nodes_ExpRefNode2Stem(self.FP.createRefNodeFromSym(paramSym)))
        expList.Insert(Nodes_ExpRefNode2Stem(self.FP.createRefNodeFromSym(stateSym)))
        self.FP.processCreateDDD(&node.Nodes_Node, expList)
        self.FP.Writeln(Lns_getVM().String_format(")%s;", []LnsAny{convCC_accessAny}))
    } else { 
        self.FP.Write(Lns_getVM().String_format("lns_func( *%s )( _pEnv, ", []LnsAny{self.moduleCtrl.FP.GetSymbolName(formSym.FP)}))
        if self.scopeMgr.FP.GetSymbolValKind(paramSym.FP) == convCC_ValKind__Any{
            self.FP.Write("*")
        }
        self.FP.Writeln(Lns_getVM().String_format("%s, %s)%s;", []LnsAny{self.moduleCtrl.FP.GetSymbolName(paramSym.FP), self.moduleCtrl.FP.GetSymbolName(stateSym.FP), convCC_accessAny}))
    }
    self.FP.Writeln(Lns_getVM().String_format("if ( lns_fromDDD( %s, 0 ).type == lns_stem_type_nil ) {", []LnsAny{workSymName}))
    self.FP.Writeln("   break;")
    self.FP.Writeln("}")
    for _index, _varSym := range( node.FP.Get_varList().Items ) {
        index := _index + 1
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        var valKind LnsInt
        valKind = self.scopeMgr.FP.GetSymbolValKind(varSym.FP)
        var varName string
        varName = self.moduleCtrl.FP.GetSymbolName(varSym.FP)
        if _switch32648 := valKind; _switch32648 == convCC_ValKind__Stem {
            self.FP.Writeln(Lns_getVM().String_format("lns_setq( _pEnv, %s, lns_fromDDD( %s, %d ) );", []LnsAny{varName, workSymName, index - 1}))
        } else if _switch32648 == convCC_ValKind__Any {
            self.FP.Writeln(Lns_getVM().String_format("lns_setq_any( _pEnv, %s, lns_fromDDD( %s, %d )%s );", []LnsAny{varName, workSymName, index - 1, convCC_accessAny}))
        } else if _switch32648 == convCC_ValKind__Prim {
            self.FP.Writeln(Lns_getVM().String_format("%s = lns_fromDDD( %s, %d )%s", []LnsAny{varName, workSymName, index - 1, convCC_getAccessValFromStem_1854_(varSym.FP.Get_typeInfo())}))
        } else {
            Util_err(Lns_getVM().String_format("no support -- %s:%s:%d", []LnsAny{varSym.FP.Get_name(), convCC_ValKind_getTxt( valKind), 7276}))
        }
    }
    convCC_filter_1642_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.processSetSymSingle(&node.Nodes_Node, nil, stateSym.FP, false, node.FP.Get_varList().GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP, false)
    self.FP.Writeln("")
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.FP.processLoopPostProcess()
    self.FP.Writeln("}")
    self.FP.processBlockPostProcess()
    self.FP.Writeln("}")
}

// 7321: decl @lune.@base.@convCC.convFilter.processForeachSetupVal
func (self *convCC_convFilter) processForeachSetupVal(parent *Nodes_Node,scope *Ast_Scope,workTxt string,symTxt string,symType *Ast_TypeInfo) {
    var symbolInfo *Ast_SymbolInfo
    
    {
        _symbolInfo := scope.FP.GetSymbolInfoChild(symTxt)
        if _symbolInfo == nil{
            Util_err(Lns_getVM().String_format("not found symTxt -- %s", []LnsAny{symTxt}))
        } else {
            symbolInfo = _symbolInfo.(*Ast_SymbolInfo)
        }
    }
    self.FP.processDeclVarC(true, symbolInfo.FP, false, nil)
    var srcSymbol *convCC_WorkSymbol
    srcSymbol = NewconvCC_WorkSymbol(symbolInfo.FP.Get_scope(), symbolInfo.FP.Get_accessMode(), workTxt, symbolInfo.FP.Get_typeInfo(), symbolInfo.FP.Get_kind(), symbolInfo.FP.Get_staticFlag(), NewconvCC_SymbolParam(convCC_ValKind__Stem, 0, convCC_cTypeStem))
    self.FP.processSetSymSingle(parent, nil, symbolInfo.FP, true, srcSymbol.FP, true)
}

// 7351: decl @lune.@base.@convCC.convFilter.processPoolForeachSetupVal
func (self *convCC_convFilter) processPoolForeachSetupVal(parent *Nodes_Node,loopType *Ast_TypeInfo,scope *Ast_Scope,keyToken LnsAny,valToken LnsAny) {
    var valType *Ast_TypeInfo
    valType = loopType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
    var valSymTxt string
    if loopType.FP.Get_kind() == Ast_TypeInfoKind__Set{
        {
            __exp := keyToken
            if __exp != nil {
                _exp := __exp.(*Ast_SymbolInfo)
                valSymTxt = _exp.FP.Get_name()
                
            } else {
                Util_err("keyToken is nil")
            }
        }
    } else { 
        {
            __exp := valToken
            if __exp != nil {
                _exp := __exp.(*Ast_SymbolInfo)
                valSymTxt = _exp.FP.Get_name()
                
            } else {
                Util_err("valToken is nil")
            }
        }
    }
    self.FP.processForeachSetupVal(parent, scope, "_val", valSymTxt, valType)
}

// 7380: decl @lune.@base.@convCC.convFilter.processMapForeachSetupVal
func (self *convCC_convFilter) processMapForeachSetupVal(parent *Nodes_Node,loopType *Ast_TypeInfo,scope *Ast_Scope,keyToken LnsAny,valToken LnsAny,keyTxt string,valTxt string) {
    if keyToken != nil{
        keyToken_8754 := keyToken.(*Ast_SymbolInfo)
        self.FP.processForeachSetupVal(parent, scope, keyTxt, keyToken_8754.FP.Get_name(), loopType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
    }
    self.FP.Writeln("")
    if valToken != nil{
        valToken_8756 := valToken.(*Ast_SymbolInfo)
        self.FP.processForeachSetupVal(parent, scope, valTxt, valToken_8756.FP.Get_name(), loopType.FP.Get_itemTypeInfoList().GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
    }
}

// 7442: decl @lune.@base.@convCC.convFilter.processForeach
func (self *convCC_convFilter) ProcessForeach(node *Nodes_ForeachNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Write(Lns_getVM().String_format("%s _obj = ", []LnsAny{convCC_cTypeAnyP}))
    self.FP.ProcessVal2any(node.FP.Get_exp(), &node.Nodes_Node)
    self.FP.Writeln(";")
    var indexSymbol LnsAny
    var loopType *Ast_TypeInfo
    loopType = node.FP.Get_exp().FP.Get_expType()
    var collectionKind LnsAny
    collectionKind = convCC_getCollectionKind_2616_(loopType)
    if _switch33552 := loopType.FP.Get_kind(); _switch33552 == Ast_TypeInfoKind__List || _switch33552 == Ast_TypeInfoKind__Array {
        self.FP.Writeln(Lns_getVM().String_format("%s _itAny = lns_itList_new( _pEnv, _obj );", []LnsAny{convCC_cTypeAnyP}))
        {
            _keyToken := node.FP.Get_key()
            if _keyToken != nil {
                keyToken := _keyToken.(*Ast_SymbolInfo)
                var workSymbol *Ast_SymbolInfo
                
                {
                    _workSymbol := node.FP.Get_block().FP.Get_scope().FP.GetSymbolInfoChild(keyToken.FP.Get_name())
                    if _workSymbol == nil{
                        Util_err(Lns_getVM().String_format("not found symbol -- %s", []LnsAny{keyToken.FP.Get_name()}))
                    } else {
                        workSymbol = _workSymbol.(*Ast_SymbolInfo)
                    }
                }
                indexSymbol = workSymbol
                
                if self.scopeMgr.FP.GetSymbolValKind(workSymbol.FP) != convCC_ValKind__Prim{
                    self.FP.Writeln(Lns_getVM().String_format("int _%s = 0;", []LnsAny{keyToken.FP.Get_name()}))
                } else { 
                    self.FP.processDeclVarC(true, workSymbol.FP, true, nil)
                }
            } else {
                indexSymbol = nil
                
            }
        }
        self.FP.Writeln(Lns_getVM().String_format("%s _val;", []LnsAny{convCC_cTypeStem}))
    } else if _switch33552 == Ast_TypeInfoKind__Set {
        self.FP.Writeln(Lns_getVM().String_format("%s _itAny = lns_itSet_new( _pEnv, _obj );", []LnsAny{convCC_cTypeAnyP}))
        indexSymbol = nil
        
        self.FP.Writeln(Lns_getVM().String_format("%s _val;", []LnsAny{convCC_cTypeStem}))
    } else if _switch33552 == Ast_TypeInfoKind__Map {
        self.FP.Writeln(Lns_getVM().String_format("%s _itAny = lns_itMap_new( _pEnv, _obj );", []LnsAny{convCC_cTypeAnyP}))
        indexSymbol = nil
        
        self.FP.Writeln("lns_Map_entry_t _entry;")
    } else if _switch33552 == Ast_TypeInfoKind__Ext {
        indexSymbol = nil
        
        switch collectionKind.(type) {
        case *convCC_CollectionKind__ExtMap:
            self.FP.Writeln(Lns_getVM().String_format("%s _itAny = lns_lua_itMap_new( _pEnv, _obj );", []LnsAny{convCC_cTypeAnyP}))
            self.FP.Writeln("lns_Map_entry_t _entry;")
        }
    } else {
        Util_err(Lns_getVM().String_format("illegal kind -- %s", []LnsAny{Ast_TypeInfoKind_getTxt( loopType.FP.Get_kind())}))
    }
    self.FP.processLoopPreProcess(node.FP.Get_block())
    if _switch33645 := loopType.FP.Get_kind(); _switch33645 == Ast_TypeInfoKind__List || _switch33645 == Ast_TypeInfoKind__Array {
        self.FP.Writeln("for ( ; lns_itList_hasNext( _pEnv, _itAny, &_val );")
        self.FP.Writeln("      lns_itList_inc( _pEnv, _itAny ) )")
    } else if _switch33645 == Ast_TypeInfoKind__Set {
        self.FP.Writeln("for ( ; lns_itSet_hasNext( _pEnv, _itAny, &_val );")
        self.FP.Writeln("      lns_itSet_inc( _pEnv, _itAny ) )")
    } else if _switch33645 == Ast_TypeInfoKind__Map {
        self.FP.Writeln("for ( ; lns_itMap_hasNext( _pEnv, _itAny, &_entry );")
        self.FP.Writeln("      lns_itMap_inc( _pEnv, _itAny ) )")
    } else if _switch33645 == Ast_TypeInfoKind__Ext {
        switch collectionKind.(type) {
        case *convCC_CollectionKind__ExtMap:
            self.FP.Writeln("while ( lns_lua_itMap_hasNext( _pEnv, _itAny ) )")
        }
    }
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Writeln("lns_reset_block( _pEnv );")
    if indexSymbol != nil{
        indexSymbol_8814 := indexSymbol.(*Ast_SymbolInfo)
        if self.scopeMgr.FP.GetSymbolValKind(indexSymbol_8814.FP) != convCC_ValKind__Prim{
            self.FP.Writeln(Lns_getVM().String_format("_%s++;", []LnsAny{self.moduleCtrl.FP.GetSymbolName(indexSymbol_8814.FP)}))
            self.FP.processDeclVarC(true, indexSymbol_8814.FP, true, nil)
            self.FP.processSetValSingleDirect(&node.Nodes_Node, nil, indexSymbol_8814.FP, true, convCC_ValKind__Prim, Ast_builtinTypeInt, 0, false, convCC_processRValue_2298_(func() {
                self.FP.Write(Lns_getVM().String_format("_%s", []LnsAny{self.moduleCtrl.FP.GetSymbolName(indexSymbol_8814.FP)}))
            }))
            self.FP.Writeln("")
        } else { 
            self.FP.Writeln(Lns_getVM().String_format("%s++;", []LnsAny{self.moduleCtrl.FP.GetSymbolName(indexSymbol_8814.FP)}))
        }
    }
    if _switch33916 := loopType.FP.Get_kind(); _switch33916 == Ast_TypeInfoKind__List || _switch33916 == Ast_TypeInfoKind__Set || _switch33916 == Ast_TypeInfoKind__Array {
        self.FP.processPoolForeachSetupVal(&node.Nodes_Node, loopType, node.FP.Get_block().FP.Get_scope(), node.FP.Get_key(), node.FP.Get_val())
    } else if _switch33916 == Ast_TypeInfoKind__Map {
        self.FP.processMapForeachSetupVal(&node.Nodes_Node, loopType, node.FP.Get_block().FP.Get_scope(), node.FP.Get_key(), node.FP.Get_val(), "_entry.key", "_entry.val")
    } else {
        switch _exp33914 := collectionKind.(type) {
        case *convCC_CollectionKind__ExtMap:
        extedType := _exp33914.Val1
            self.FP.Writeln("lns_lua_itMap_getEntry( _pEnv, _itAny, &_entry );")
            self.FP.processMapForeachSetupVal(&node.Nodes_Node, extedType, node.FP.Get_block().FP.Get_scope(), node.FP.Get_key(), node.FP.Get_val(), "_entry.key", "_entry.val")
        }
    }
    convCC_filter_1642_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.FP.processLoopPostProcess()
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 7584: decl @lune.@base.@convCC.convFilter.processForsort
func (self *convCC_convFilter) ProcessForsort(node *Nodes_ForsortNode,_opt LnsAny) {
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Write(Lns_getVM().String_format("%s _obj = ", []LnsAny{convCC_cTypeAnyP}))
    self.FP.ProcessVal2any(node.FP.Get_exp(), &node.Nodes_Node)
    self.FP.Writeln(";")
    var loopType *Ast_TypeInfo
    loopType = node.FP.Get_exp().FP.Get_expType()
    if _switch34178 := loopType.FP.Get_kind(); _switch34178 == Ast_TypeInfoKind__Set {
        self.FP.Writeln(Lns_getVM().String_format("%s _pList = lns_mtd_Map_createKeyList( _pEnv, _obj );", []LnsAny{convCC_cTypeAnyP}))
        self.FP.Writeln(Lns_getVM().String_format("lns_mtd_List( _pList )->sort( _pEnv, _pList, %s );", []LnsAny{convCC_cValNil}))
        self.FP.Writeln(Lns_getVM().String_format("%s _itAny = lns_itList_new( _pEnv, _pList );", []LnsAny{convCC_cTypeAnyP}))
        self.FP.Writeln(Lns_getVM().String_format("%s _val;", []LnsAny{convCC_cTypeStem}))
    } else if _switch34178 == Ast_TypeInfoKind__Map {
        self.FP.Writeln(Lns_getVM().String_format("%s _pKeyList = lns_mtd_Map_createKeyList( _pEnv, _obj );", []LnsAny{convCC_cTypeAnyP}))
        self.FP.Writeln(Lns_getVM().String_format("lns_mtd_List( _pKeyList )->sort( _pEnv, _pKeyList, %s );", []LnsAny{convCC_cValNil}))
        self.FP.Writeln(Lns_getVM().String_format("%s _itAny = lns_itList_new( _pEnv, _pKeyList );", []LnsAny{convCC_cTypeAnyP}))
        self.FP.Writeln(Lns_getVM().String_format("%s _key;", []LnsAny{convCC_cTypeStem}))
    } else {
        Util_err(Lns_getVM().String_format("illegal kind -- %s", []LnsAny{Ast_TypeInfoKind_getTxt( loopType.FP.Get_kind())}))
    }
    self.FP.processLoopPreProcess(node.FP.Get_block())
    if _switch34232 := loopType.FP.Get_kind(); _switch34232 == Ast_TypeInfoKind__Set {
        self.FP.Writeln("for ( ; lns_itList_hasNext( _pEnv, _itAny, &_val );")
        self.FP.Writeln("      lns_itList_inc( _pEnv, _itAny ) )")
    } else if _switch34232 == Ast_TypeInfoKind__Map {
        self.FP.Writeln("for ( ; lns_itList_hasNext( _pEnv, _itAny, &_key );")
        self.FP.Writeln("      lns_itList_inc( _pEnv, _itAny ) )")
    }
    self.FP.Writeln("{")
    self.FP.Writeln("lns_reset_block( _pEnv );")
    if _switch34311 := loopType.FP.Get_kind(); _switch34311 == Ast_TypeInfoKind__Set {
        self.FP.processPoolForeachSetupVal(&node.Nodes_Node, loopType, node.FP.Get_block().FP.Get_scope(), node.FP.Get_key(), node.FP.Get_val())
    } else if _switch34311 == Ast_TypeInfoKind__Map {
        self.FP.processMapForeachSetupVal(&node.Nodes_Node, loopType, node.FP.Get_block().FP.Get_scope(), node.FP.Get_key(), node.FP.Get_val(), "_key", "lns_mtd_Map_get( _pEnv, _obj, _key )")
    } else {
    }
    convCC_filter_1642_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln("}")
    self.FP.processLoopPostProcess()
    self.FP.Writeln("}")
    self.FP.PopIndent()
}

// 7657: decl @lune.@base.@convCC.convFilter.processExpUnwrap
func (self *convCC_convFilter) ProcessExpUnwrap(node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    var processUnwrap func(typeTxt string)
    processUnwrap = func(typeTxt string) {
        {
            _defVal := node.FP.Get_default()
            if _defVal != nil {
                defVal := _defVal.(*Nodes_Node)
                self.FP.Write(Lns_getVM().String_format("lns_unwrap_%sDefault( ", []LnsAny{typeTxt}))
                self.FP.ProcessVal2stem(node.FP.Get_exp(), &node.Nodes_Node)
                self.FP.Write(",")
                self.FP.AccessPrimVal(defVal, &node.Nodes_Node)
                self.FP.Write(")")
            } else {
                self.FP.Write(Lns_getVM().String_format("lns_unwrap_%s( ", []LnsAny{typeTxt}))
                self.FP.ProcessVal2stem(node.FP.Get_exp(), &node.Nodes_Node)
                self.FP.Write(")")
            }
        }
    }
    if _switch34640 := convCC_getOrgTypeInfo_1430_(node.FP.Get_expType()); _switch34640 == Ast_builtinTypeInt || _switch34640 == Ast_builtinTypeChar {
        processUnwrap("int")
    } else if _switch34640 == Ast_builtinTypeReal {
        processUnwrap("real")
    } else if _switch34640 == Ast_builtinTypeBool {
        processUnwrap("bool")
    } else {
        if _switch34575 := self.FP.getValKindOfNode(&node.Nodes_Node); _switch34575 == convCC_ValKind__Stem {
            self.FP.Write("lns_unwrap_stem( ")
        } else if _switch34575 == convCC_ValKind__Any {
            self.FP.Write("lns_unwrap_any( ")
        } else {
            Util_err(Lns_getVM().String_format("no support -- %s: %d", []LnsAny{convCC_ValKind_getTxt( self.FP.getValKindOfNode(&node.Nodes_Node)), 7697}))
        }
        self.FP.ProcessVal2stem(node.FP.Get_exp(), &node.Nodes_Node)
        {
            _defVal := node.FP.Get_default()
            if _defVal != nil {
                defVal := _defVal.(*Nodes_Node)
                self.FP.Write(",")
                self.FP.ProcessVal2stem(defVal, &node.Nodes_Node)
                self.FP.Write(")")
            } else {
                self.FP.Write(Lns_getVM().String_format(", %s )", []LnsAny{convCC_cValNone}))
            }
        }
    }
}

// 7728: decl @lune.@base.@convCC.convFilter.processCreateMRet
func (self *convCC_convFilter) processCreateMRet(retTypeList *LnsList,expList *LnsList,parent *Nodes_Node) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( expList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__DDD) &&
        Lns_GetEnv().SetStackVal( expList.Len() == 1) ).(bool)){
        self.FP.Write("_pDDD")
        return 
    }
    self.FP.Write("lns_createMRet")
    var lastExp *Nodes_Node
    lastExp = expList.GetAt(expList.Len()).(Nodes_NodeDownCast).ToNodes_Node()
    self.FP.Write(Lns_getVM().String_format("( _pEnv, %s, %d", []LnsAny{Nodes_hasMultiValNode(lastExp), expList.Len()}))
    for _, _exp := range( expList.Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.Write(", ")
        self.FP.ProcessVal2stem(exp, parent)
    }
    self.FP.Write(")")
}

// 7812: decl @lune.@base.@convCC.convFilter.processCallWithMRet
func (self *convCC_convFilter) ProcessCallWithMRet(parent *Nodes_Node,mRetFuncName string,retTypeName string,mRetInfo LnsAny,argList *Nodes_ExpListNode) {
    var mRetExp *Nodes_MRetExp
    
    {
        _mRetExp := argList.FP.Get_mRetExp()
        if _mRetExp == nil{
            return 
        } else {
            mRetExp = _mRetExp.(*Nodes_MRetExp)
        }
    }
    switch _exp34973 := mRetInfo.(type) {
    case *convCC_MRetInfo__Method:
    funcType := _exp34973.Val1
        if Lns_op_not(convCC_needMRetWrap_2651_(funcType.FP.Get_argTypeInfoList(), argList)){
            return 
        }
    case *convCC_MRetInfo__Form:
        if Lns_op_not(convCC_needMRetWrap_2651_(NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeDDD)}), argList)){
            return 
        }
    case *convCC_MRetInfo__FormFunc:
    funcType := _exp34973.Val1
        if Lns_op_not(convCC_needMRetWrap_2651_(funcType.FP.Get_argTypeInfoList(), argList)){
            return 
        }
    case *convCC_MRetInfo__Func:
    funcNode := _exp34973.Val1
        if Lns_op_not(convCC_needMRetWrap_2651_(funcNode.FP.Get_expType().FP.Get_argTypeInfoList(), argList)){
            return 
        }
    case *convCC_MRetInfo__DDD:
    case *convCC_MRetInfo__Format:
    }
    var processDeclMRetProto func()
    processDeclMRetProto = func() {
        self.FP.Write(Lns_getVM().String_format("static %s %s( %s _pEnv", []LnsAny{retTypeName, mRetFuncName, convCC_cTypeEnvP}))
        var processArgs func()
        processArgs = func() {
            for _index, _argNode := range( argList.FP.Get_expList().Items ) {
                index := _index + 1
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                if index >= mRetExp.FP.Get_index(){
                    break
                }
                var argType *Ast_TypeInfo
                argType = argNode.FP.Get_expType()
                self.FP.Write(Lns_getVM().String_format(", %s arg%d", []LnsAny{convCC_getCType_1188_(argType), index}))
            }
            self.FP.Write(Lns_getVM().String_format(", %s pMRet )", []LnsAny{convCC_cTypeStem}))
        }
        switch mRetInfo.(type) {
        case *convCC_MRetInfo__Method:
            self.FP.Write(Lns_getVM().String_format(", %s _pObj", []LnsAny{convCC_cTypeAnyP}))
            processArgs()
        case *convCC_MRetInfo__Form:
            self.FP.Write(Lns_getVM().String_format(", %s _pForm", []LnsAny{convCC_cTypeAnyP}))
            processArgs()
        case *convCC_MRetInfo__FormFunc:
            self.FP.Write(Lns_getVM().String_format(", %s _pForm", []LnsAny{convCC_cTypeAnyP}))
            processArgs()
        case *convCC_MRetInfo__DDD:
            processArgs()
        case *convCC_MRetInfo__Func:
            processArgs()
        case *convCC_MRetInfo__Format:
            processArgs()
        }
    }
    if _switch36301 := self.processMode; _switch36301 == convCC_ProcessMode__Intermediate || _switch36301 == convCC_ProcessMode__StringFormat {
        processDeclMRetProto()
        self.FP.Writeln(Lns_getVM().String_format("// %d", []LnsAny{parent.FP.Get_pos().LineNo}))
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        var argTypeList *LnsList
        argTypeList = NewLnsList([]LnsAny{})
        var processSetArg func(primFlag bool)
        processSetArg = func(primFlag bool) {
            for _index, _argNode := range( argList.FP.Get_expList().Items ) {
                index := _index + 1
                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                var argType *Ast_TypeInfo
                argType = argNode.FP.Get_expType()
                if index == mRetExp.FP.Get_index(){
                    self.FP.Writeln(Lns_getVM().String_format("lns_setMRet( _pEnv, pMRet%s );", []LnsAny{convCC_accessAny}))
                }
                if index >= mRetExp.FP.Get_index(){
                    if _switch35556 := argNode.FP.Get_kind(); _switch35556 == Nodes_NodeKind_get_ExpToDDD() {
                        var toDDDNode *Nodes_ExpToDDDNode
                        toDDDNode = Lns_unwrap( Nodes_ExpToDDDNodeDownCastF(argNode.FP)).(*Nodes_ExpToDDDNode)
                        self.FP.Write(Lns_getVM().String_format("%s arg%d = ", []LnsAny{convCC_cTypeStem, index}))
                        self.FP.Write("lns_createDDD")
                        var expList *LnsList
                        expList = toDDDNode.FP.Get_expList().FP.Get_expList()
                        var lastExp *Nodes_Node
                        lastExp = expList.GetAt(expList.Len()).(Nodes_NodeDownCast).ToNodes_Node()
                        self.FP.Write(Lns_getVM().String_format("( _pEnv, %s, %d", []LnsAny{Nodes_hasMultiValNode(lastExp), expList.Len()}))
                        for _workIndex, _ := range( expList.Items ) {
                            workIndex := _workIndex + 1
                            self.FP.Write(Lns_getVM().String_format(", lns_getMRet( _pEnv, %d )", []LnsAny{workIndex + index - 2}))
                        }
                        self.FP.Write(")")
                        argTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeDDD))
                    } else if _switch35556 == Nodes_NodeKind_get_ExpSubDDD() {
                        self.FP.Write(Lns_getVM().String_format("%s arg%d = ", []LnsAny{convCC_cTypeStem, index}))
                        convCC_filter_1642_(argNode, self, parent)
                    } else {
                        {
                            _castNode := Nodes_ExpCastNodeDownCastF(argNode.FP)
                            if _castNode != nil {
                                castNode := _castNode.(*Nodes_ExpCastNode)
                                argType = castNode.FP.Get_castType()
                                
                            }
                        }
                        var typeTxt string
                        if primFlag{
                            typeTxt = convCC_getCType_1188_(argType)
                            
                            argTypeList.Insert(Ast_TypeInfo2Stem(argType.FP.Get_srcTypeInfo()))
                        } else { 
                            typeTxt = convCC_cTypeStem
                            
                            argTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeStem))
                        }
                        self.FP.Write(Lns_getVM().String_format("%s arg%d = ", []LnsAny{typeTxt, index}))
                        if argType.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                            if index == mRetExp.FP.Get_index(){
                                self.FP.Write("pMRet")
                            } else { 
                                self.FP.Write(Lns_getVM().String_format("lns_createSubDDD( _pEnv, %d, pMRet )", []LnsAny{index - mRetExp.FP.Get_index()}))
                            }
                        } else { 
                            self.FP.Write(Lns_getVM().String_format("lns_getMRet( _pEnv, %d )", []LnsAny{index - mRetExp.FP.Get_index()}))
                        }
                        if primFlag{
                            self.FP.Write(convCC_getAccessValFromStem_1854_(argType))
                        } else { 
                            self.FP.Write("->val.pAny")
                        }
                    }
                    self.FP.Writeln(Lns_getVM().String_format("; // %s", []LnsAny{argType.FP.GetTxt(self.FP.Get_typeNameCtrl(), nil, nil)}))
                } else { 
                    argTypeList.Insert(Ast_TypeInfo2Stem(argNode.FP.Get_expType()))
                }
            }
            if retTypeName != "void"{
                self.FP.Write("return ")
            }
        }
        var processArg2Stem func(index LnsInt,typeInfo *Ast_TypeInfo)
        processArg2Stem = func(index LnsInt,typeInfo *Ast_TypeInfo) {
            if argTypeList.Len() >= index{
                if _switch35770 := argTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(); _switch35770 == Ast_builtinTypeInt || _switch35770 == Ast_builtinTypeChar {
                    self.FP.Write("LNS_STEM_INT(")
                    self.FP.Write(Lns_getVM().String_format("arg%d )", []LnsAny{index}))
                } else if _switch35770 == Ast_builtinTypeReal {
                    self.FP.Write("LNS_STEM_REAL(")
                    self.FP.Write(Lns_getVM().String_format("arg%d )", []LnsAny{index}))
                } else if _switch35770 == Ast_builtinTypeBool {
                    self.FP.Write("LNS_STEM_BOOL(")
                    self.FP.Write(Lns_getVM().String_format("arg%d )", []LnsAny{index}))
                } else {
                    if convCC_getValKind_1167_(argTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()) == convCC_ValKind__Any{
                        self.FP.Write("LNS_STEM_ANY(")
                        self.FP.Write(Lns_getVM().String_format("arg%d )", []LnsAny{index}))
                    } else { 
                        self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
                    }
                }
            } else { 
                if typeInfo.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                    var offset LnsInt
                    offset = index - mRetExp.FP.Get_index()
                    if offset > 0{
                        self.FP.Write(Lns_getVM().String_format("lns_createSubDDD( _pEnv, %d, pMRet )", []LnsAny{index - mRetExp.FP.Get_index()}))
                    } else if offset == 0{
                        self.FP.Write("pMRet")
                    } else { 
                        self.FP.Write(convCC_cValDDD0)
                    }
                } else { 
                    self.FP.Write(convCC_cValNone)
                }
            }
        }
        var processCreateDDD func(expList *LnsList)
        processCreateDDD = func(expList *LnsList) {
            self.FP.Write("lns_createDDD")
            var lastExp *Nodes_Node
            lastExp = expList.GetAt(expList.Len()).(Nodes_NodeDownCast).ToNodes_Node()
            self.FP.Write(Lns_getVM().String_format("( _pEnv, %s, %d", []LnsAny{Nodes_hasMultiValNode(lastExp), expList.Len()}))
            {
                var _from35925 LnsInt = 1
                var _to35925 LnsInt = expList.Len()
                for _work35925 := _from35925; _work35925 <= _to35925; _work35925++ {
                    index := _work35925
                    self.FP.Write(", ")
                    processArg2Stem(index, Ast_builtinTypeNone)
                }
            }
        }
        var funcTypeInfo LnsAny
        funcTypeInfo = nil
        switch _exp36197 := mRetInfo.(type) {
        case *convCC_MRetInfo__Method:
        funcType := _exp36197.Val1
            funcTypeInfo = funcType
            
            processSetArg(true)
            self.FP.Write(Lns_getVM().String_format("%s( _pEnv, _pObj", []LnsAny{self.moduleCtrl.FP.GetCallMethodCName(funcType)}))
        case *convCC_MRetInfo__Form:
            processSetArg(true)
            self.FP.Write("lns_closure( _pForm )( _pEnv, pForm")
            processCreateDDD(argList.FP.Get_expList())
        case *convCC_MRetInfo__FormFunc:
        funcType := _exp36197.Val1
            funcTypeInfo = funcType
            
            processSetArg(true)
            self.FP.Write(Lns_getVM().String_format("%s( _pEnv, _pForm", []LnsAny{self.moduleCtrl.FP.GetCallFormName(funcType)}))
        case *convCC_MRetInfo__Func:
        funcNode := _exp36197.Val1
            funcTypeInfo = funcNode.FP.Get_expType()
            
            processSetArg(true)
            var wroteFuncFlag bool
            wroteFuncFlag = false
            _ = TransUnit_getBuiltinFunc()
            {
                _cFuncName := self.moduleCtrl.FP.GetBuiltinFuncNameFromType(funcNode.FP.Get_expType())
                if _cFuncName != nil {
                    cFuncName := _cFuncName.(string)
                    wroteFuncFlag = true
                    
                    self.FP.Write(cFuncName + "(")
                }
            }
            if Lns_op_not(wroteFuncFlag){
                convCC_filter_1642_(funcNode, self, parent)
                self.FP.Write("(")
            }
            self.FP.Write(" _pEnv")
        case *convCC_MRetInfo__DDD:
        expListNode := _exp36197.Val1
            processSetArg(true)
            processCreateDDD(expListNode.FP.Get_expList())
        case *convCC_MRetInfo__Format:
        format := _exp36197.Val1
        expListNode := _exp36197.Val2
            processSetArg(true)
            self.FP.Write("mtd_lns_string_format( _pEnv, ")
            self.FP.Write(convCC_getLiteralStrAny_1497_(format))
            self.FP.Write(", ")
            processCreateDDD(expListNode.FP.Get_expList())
            self.FP.Write(")")
        }
        if funcTypeInfo != nil{
            funcTypeInfo_9029 := funcTypeInfo.(*Ast_TypeInfo)
            for _index, _argType := range( funcTypeInfo_9029.FP.Get_argTypeInfoList().Items ) {
                index := _index + 1
                argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if convCC_getValKind_1167_(argType) == convCC_ValKind__Stem{
                    self.FP.Write(", ")
                    processArg2Stem(index, argType)
                } else { 
                    self.FP.Write(Lns_getVM().String_format(", arg%d", []LnsAny{index}))
                }
            }
        }
        self.FP.PopIndent()
        self.FP.Writeln(");")
        self.FP.Writeln("}")
    } else if _switch36301 == convCC_ProcessMode__Prototype {
        processDeclMRetProto()
        self.FP.Writeln(Lns_getVM().String_format("; // %d", []LnsAny{argList.FP.Get_pos().LineNo}))
    }
}

// 8117: decl @lune.@base.@convCC.convFilter.processExpToDDD
func (self *convCC_convFilter) ProcessExpToDDD(node *Nodes_ExpToDDDNode,_opt LnsAny) {
    if _switch36388 := self.processMode; _switch36388 == convCC_ProcessMode__Intermediate || _switch36388 == convCC_ProcessMode__Prototype {
        {
            _mRetExp := node.FP.Get_expList().FP.Get_mRetExp()
            if _mRetExp != nil {
                mRetExp := _mRetExp.(*Nodes_MRetExp)
                if mRetExp.FP.Get_index() > 0{
                    self.FP.ProcessCallWithMRet(&node.Nodes_Node, convCC_getMRetFuncName_2692_(&node.Nodes_Node), convCC_cTypeStem, &convCC_MRetInfo__DDD{node.FP.Get_expList()}, node.FP.Get_expList())
                }
            }
        }
        return 
    }
    var expList *LnsList
    expList = node.FP.Get_expList().FP.Get_expList()
    {
        _mRetExp := node.FP.Get_expList().FP.Get_mRetExp()
        if _mRetExp != nil {
            mRetExp := _mRetExp.(*Nodes_MRetExp)
            self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{convCC_getMRetFuncName_2692_(&node.Nodes_Node)}))
            for _index, _exp := range( expList.Items ) {
                index := _index + 1
                exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
                if index > mRetExp.FP.Get_index(){
                    break
                }
                self.FP.Write(", ")
                convCC_filter_1642_(exp, self, &node.Nodes_Node)
            }
            self.FP.Write(")")
        } else {
            self.FP.processCreateDDD(&node.Nodes_Node, expList)
        }
    }
}

// 8154: decl @lune.@base.@convCC.convFilter.processExpNew
func (self *convCC_convFilter) ProcessExpNew(node *Nodes_ExpNewNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{self.moduleCtrl.FP.GetNewName(node.FP.Get_symbol().FP.Get_expType())}))
    {
        __exp := node.FP.Get_argList()
        if __exp != nil {
            _exp := __exp.(*Nodes_ExpListNode)
            self.FP.ProcessCallArgList(node.FP.Get_ctorTypeInfo(), _exp)
        }
    }
    self.FP.Write(")")
}

// 8168: decl @lune.@base.@convCC.convFilter.processCall
func (self *convCC_convFilter) processCall(funcSym LnsAny,funcType *Ast_TypeInfo,setArgFlag bool,argList LnsAny) {
    if Lns_op_not(setArgFlag){
        self.FP.Write("_pEnv")
        {
            _scope := funcType.FP.Get_scope()
            if _scope != nil {
                scope := _scope.(*Ast_Scope)
                if scope.FP.Get_closureSymList().Len() > 0{
                    self.FP.Write(", ")
                    var setFlag bool
                    setFlag = false
                    if funcSym != nil{
                        funcSym_9082 := funcSym.(*Ast_SymbolInfo)
                        if funcSym_9082.FP.Get_hasAccessFromClosure(){
                            self.FP.ProcessSym2Any(funcSym_9082.FP)
                            setFlag = true
                            
                        }
                    }
                    if Lns_op_not(setFlag){
                        self.FP.Write(convCC_getPrepareClosure_1894_(self.scopeMgr, "NULL", 0, false, scope.FP.Get_closureSymList()))
                    }
                }
            }
        }
    }
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( funcType.FP.Get_kind() == Ast_TypeInfoKind__Func) &&
        Lns_GetEnv().SetStackVal( funcType.FP.Get_rawTxt() == "___init") &&
        Lns_GetEnv().SetStackVal( funcType.FP.Get_parentInfo().FP.Get_kind() == Ast_TypeInfoKind__Class) ).(bool)){
        self.FP.Write(Lns_getVM().String_format(", %s", []LnsAny{convCC_getBlockName_1198_(self.ast.FP.Get_moduleScope())}))
    } else { 
        if argList != nil{
            argList_9088 := argList.(*Nodes_ExpListNode)
            var expList *LnsList
            expList = NewLnsList([]LnsAny{})
            for _, _expNode := range( argList_9088.FP.Get_expList().Items ) {
                expNode := _expNode.(Nodes_NodeDownCast).ToNodes_Node()
                if expNode.FP.Get_expType().FP.Get_kind() != Ast_TypeInfoKind__Abbr{
                    expList.Insert(Nodes_Node2Stem(expNode))
                }
            }
            self.FP.ProcessCallArgList(funcType, argList_9088)
        } else {
            self.FP.ProcessCallArgList(funcType, nil)
        }
    }
    self.FP.Write(" )")
}

// 8219: decl @lune.@base.@convCC.convFilter.processDeclClass
func (self *convCC_convFilter) ProcessDeclClass(node *Nodes_DeclClassNode,_opt LnsAny) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType()
    var classCanonicalName string
    classCanonicalName = self.moduleCtrl.FP.GetCanonicalName(classType)
    self.FP.Writeln(Lns_getVM().String_format("// decl class %s (%s)-->", []LnsAny{classCanonicalName, convCC_ProcessMode_getTxt( self.processMode)}))
    if _switch37006 := self.processMode; _switch37006 == convCC_ProcessMode__Prototype {
        self.FP.processDeclClassNodePrototype(node)
    } else if _switch37006 == convCC_ProcessMode__WideScopeVer {
        if _switch36900 := classType.FP.Get_kind(); _switch36900 == Ast_TypeInfoKind__Class {
            convCC_processIFMethodDataInit_2194_(self.stream.FP, self.moduleCtrl, classType, classType)
            convCC_processClassDataInit_2204_(self.stream.FP, self.moduleCtrl, self.scopeMgr, classType, node.FP.Get_fieldList())
        } else if _switch36900 == Ast_TypeInfoKind__IF {
            convCC_processClassMeta_2197_(self.stream.FP, self.moduleCtrl, classType)
        }
    } else if _switch37006 == convCC_ProcessMode__DefClass {
        if classType.FP.Get_kind() == Ast_TypeInfoKind__Class{
            self.FP.processDeclClassDef(node)
            self.FP.processMapping(node, classType, Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( Ast_isPubToExternal(classType.FP.Get_accessMode())) &&
                Lns_GetEnv().SetStackVal( convCC_Out2HMode__SourcePub) ||
                Lns_GetEnv().SetStackVal( convCC_Out2HMode__SourcePri) ).(LnsInt))
        }
    } else if _switch37006 == convCC_ProcessMode__Form || _switch37006 == convCC_ProcessMode__InitModule {
        {
            _initBlockNode := node.FP.Get_initBlock().FP.Get_func()
            if _initBlockNode != nil {
                initBlockNode := _initBlockNode.(*Nodes_DeclMethodNode)
                self.FP.Write(Lns_getVM().String_format("%s( ", []LnsAny{self.moduleCtrl.FP.GetMethodCName(initBlockNode.FP.Get_expType())}))
                self.FP.processCall(nil, initBlockNode.FP.Get_expType(), false, nil)
                self.FP.Writeln(";")
            }
        }
    }
    self.FP.Writeln(Lns_getVM().String_format("// <--- decl class %s (%s)", []LnsAny{classCanonicalName, convCC_ProcessMode_getTxt( self.processMode)}))
}

// 8588: decl @lune.@base.@convCC.convFilter.processExpCallDefWrap
func (self *convCC_convFilter) processExpCallDefWrap(node *Nodes_ExpCallNode,opt *ConvCC_Opt) {
    var funcType *Ast_TypeInfo
    if node.FP.Get_nilAccess(){
        funcType = node.FP.Get_func().FP.Get_expType().FP.Get_nonnilableType()
        
    } else { 
        return 
    }
    if funcType.FP.Get_kind() != Ast_TypeInfoKind__FormFunc{
        return 
    }
    var retCode string
    var retType string
    if convCC_getCRetType_1195_(funcType.FP.Get_retTypeInfoList()) == "void"{
        retType = "void"
        
        retCode = ""
        
    } else { 
        retType = convCC_cTypeStem
        
        retCode = convCC_cValNil
        
    }
    var argNameList *LnsList
    argNameList = NewLnsList([]LnsAny{})
    self.FP.Write(Lns_getVM().String_format("static %s %s( %s * _pEnv, %s form", []LnsAny{retType, convCC_getFormNilWrapper_2713_(&node.Nodes_Node), convCC_cTypeEnvP, convCC_cTypeStem}))
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList().Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var name string
        name = Lns_getVM().String_format("arg%d", []LnsAny{index})
        self.FP.Write(Lns_getVM().String_format(", %s %s", []LnsAny{convCC_getCType_1188_(argType), name}))
        argNameList.Insert(name)
    }
    self.FP.Writeln(")")
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Write(Lns_getVM().String_format("if ( form.type == lns_stem_type_nil ) {\n   return %s;\n}\n", []LnsAny{retCode}))
    self.FP.processCallUserForm(Lns_getVM().String_format("form%s", []LnsAny{convCC_accessAny}), funcType, argNameList)
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 8638: decl @lune.@base.@convCC.convFilter.processExpCall
func (self *convCC_convFilter) ProcessExpCall(node *Nodes_ExpCallNode,_opt LnsAny) {
    opt := _opt.(*ConvCC_Opt)
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_func().FP.Get_expType().FP.Get_nonnilableType()
    if _switch37435 := self.processMode; _switch37435 == convCC_ProcessMode__Intermediate || _switch37435 == convCC_ProcessMode__Prototype {
        {
            _argList := node.FP.Get_argList()
            if _argList != nil {
                argList := _argList.(*Nodes_ExpListNode)
                var funcNode *Nodes_Node
                funcNode = node.FP.Get_func()
                var mRetInfo LnsAny
                if _switch37385 := funcNode.FP.Get_expType().FP.Get_kind(); _switch37385 == Ast_TypeInfoKind__Method {
                    mRetInfo = &convCC_MRetInfo__Method{funcType}
                    
                } else if _switch37385 == Ast_TypeInfoKind__Form {
                    mRetInfo = convCC_MRetInfo__Form_Obj
                    
                } else if _switch37385 == Ast_TypeInfoKind__FormFunc {
                    mRetInfo = &convCC_MRetInfo__FormFunc{node.FP.Get_func().FP.Get_expType()}
                    
                } else {
                    mRetInfo = &convCC_MRetInfo__Func{node.FP.Get_func()}
                    
                }
                self.FP.ProcessCallWithMRet(&node.Nodes_Node, convCC_getMRetFuncName_2692_(&node.Nodes_Node), convCC_getCRetType_1195_(node.FP.Get_expTypeList()), mRetInfo, argList)
            }
        }
        return 
    } else if _switch37435 == convCC_ProcessMode__DefWrap {
        self.FP.processExpCallDefWrap(node, opt)
        return 
    }
    var process func()
    process = func() {
        if funcType.FP.Get_kind() == Ast_TypeInfoKind__Form{
            self.FP.Write("lns_call_form( _pEnv, ")
            self.FP.ProcessVal2any(node.FP.Get_func(), &node.Nodes_Node)
            {
                _argList := node.FP.Get_argList()
                if _argList != nil {
                    argList := _argList.(*Nodes_ExpListNode)
                    if argList.FP.Get_expList().Len() > 0{
                        self.FP.ProcessCallArgList(funcType, argList)
                    }
                } else {
                    self.FP.Write(", lns_global.ddd0")
                }
            }
            self.FP.Write(" )")
            return 
        }
        var wroteFuncFlag bool
        wroteFuncFlag = false
        var setArgFlag bool
        setArgFlag = false
        var fieldCall func() bool
        fieldCall = func() bool {
            var fieldNode *Nodes_RefFieldNode
            
            {
                _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func().FP)
                if _fieldNode == nil{
                    return true
                } else {
                    fieldNode = _fieldNode.(*Nodes_RefFieldNode)
                }
            }
            var prefixNode *Nodes_Node
            prefixNode = fieldNode.FP.Get_prefix()
            var prefixType *Ast_TypeInfo
            prefixType = prefixNode.FP.Get_expType()
            if node.FP.Get_nilAccess(){
            } else { 
                if _switch37634 := prefixType.FP.Get_kind(); _switch37634 == Ast_TypeInfoKind__Enum || _switch37634 == Ast_TypeInfoKind__Alge {
                    convFilter_processExpCall_process_fieldCall__processEnumAlge_2734_()
                } else if _switch37634 == Ast_TypeInfoKind__Class {
                    if prefixType.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil){
                        wroteFuncFlag = true
                        
                        setArgFlag = true
                        
                        self.FP.Write(Lns_getVM().String_format("mtd_lns_string_%s( _pEnv, ", []LnsAny{fieldNode.FP.Get_field().Txt}))
                        convCC_filter_1642_(prefixNode, self, &fieldNode.Nodes_Node)
                    }
                }
            }
            return true
        }
        var funcSym LnsAny
        funcSym = nil
        if Lns_op_not(fieldCall()){
            return 
        }
        {
            _refNode := Nodes_ExpRefNodeDownCastF(node.FP.Get_func().FP)
            if _refNode != nil {
                refNode := _refNode.(*Nodes_ExpRefNode)
                {
                    _cFuncTxt := self.moduleCtrl.FP.GetBuiltinFuncNameFromType(refNode.FP.Get_expType())
                    if _cFuncTxt != nil {
                        cFuncTxt := _cFuncTxt.(string)
                        wroteFuncFlag = true
                        
                        self.FP.Write(cFuncTxt + "(")
                    }
                }
            }
        }
        if Lns_op_not(wroteFuncFlag){
            var funcSymList *LnsList
            funcSymList = node.FP.Get_func().FP.GetSymbolInfo()
            if funcSymList.Len() > 0{
                var workFuncSym *Ast_SymbolInfo
                workFuncSym = funcSymList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.GetOrg()
                funcSym = workFuncSym
                
                {
                    _cFuncName := self.moduleCtrl.FP.GetBuiltinFuncNameFromType(workFuncSym.FP.Get_typeInfo())
                    if _cFuncName != nil {
                        cFuncName := _cFuncName.(string)
                        wroteFuncFlag = true
                        
                        self.FP.Write(cFuncName)
                        self.FP.Write("(")
                    }
                }
            }
        }
        if Lns_op_not(wroteFuncFlag){
            if _switch37998 := funcType.FP.Get_kind(); _switch37998 == Ast_TypeInfoKind__Method {
                {
                    _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func().FP)
                    if _fieldNode != nil {
                        fieldNode := _fieldNode.(*Nodes_RefFieldNode)
                        if node.FP.Get_nilAccess(){
                            self.FP.Write(self.moduleCtrl.FP.GetNilMethodCName(funcType))
                            self.FP.Write("( _pEnv, ")
                            self.FP.ProcessVal2stem(fieldNode.FP.Get_prefix(), &fieldNode.Nodes_Node)
                        } else { 
                            self.FP.Write(self.moduleCtrl.FP.GetCallMethodCName(funcType))
                            self.FP.Write("( _pEnv, ")
                            self.FP.ProcessVal2any(fieldNode.FP.Get_prefix(), &fieldNode.Nodes_Node)
                        }
                    }
                }
                wroteFuncFlag = true
                
                setArgFlag = true
                
            } else if _switch37998 == Ast_TypeInfoKind__Func {
                self.FP.Write(Lns_getVM().String_format("%s( ", []LnsAny{self.moduleCtrl.FP.GetFuncName(funcType)}))
                wroteFuncFlag = true
                
            } else if _switch37998 == Ast_TypeInfoKind__FormFunc {
                if node.FP.Get_nilAccess(){
                    self.FP.Write(Lns_getVM().String_format("%s( _pEnv, ", []LnsAny{convCC_getFormNilWrapper_2713_(&node.Nodes_Node)}))
                } else { 
                    self.FP.Write(Lns_getVM().String_format("%s( _pEnv, ", []LnsAny{self.moduleCtrl.FP.GetCallFormName(funcType)}))
                }
                self.FP.ProcessVal2any(node.FP.Get_func(), &node.Nodes_Node)
                wroteFuncFlag = true
                
                setArgFlag = true
                
            }
        }
        if Lns_op_not(wroteFuncFlag){
            convCC_filter_1642_(node.FP.Get_func(), self, &node.Nodes_Node)
            self.FP.Write("( ")
        }
        self.FP.processCall(funcSym, funcType, setArgFlag, node.FP.Get_argList())
    }
    var call func()
    call = func() {
        var isMret bool
        isMret = false
        {
            _argList := node.FP.Get_argList()
            if _argList != nil {
                argList := _argList.(*Nodes_ExpListNode)
                {
                    _mRetExp := argList.FP.Get_mRetExp()
                    if _mRetExp != nil {
                        mRetExp := _mRetExp.(*Nodes_MRetExp)
                        if convCC_needMRetWrap_2651_(node.FP.Get_func().FP.Get_expType().FP.Get_argTypeInfoList(), argList){
                            isMret = true
                            
                            self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{convCC_getMRetFuncName_2692_(&node.Nodes_Node)}))
                            var funcNode *Nodes_Node
                            funcNode = node.FP.Get_func()
                            if _switch38174 := funcNode.FP.Get_expType().FP.Get_kind(); _switch38174 == Ast_TypeInfoKind__Method {
                                {
                                    _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func().FP)
                                    if _fieldNode != nil {
                                        fieldNode := _fieldNode.(*Nodes_RefFieldNode)
                                        self.FP.Write(", ")
                                        self.FP.ProcessVal2any(fieldNode.FP.Get_prefix(), &fieldNode.Nodes_Node)
                                    }
                                }
                            } else if _switch38174 == Ast_TypeInfoKind__Form || _switch38174 == Ast_TypeInfoKind__FormFunc {
                                self.FP.Write(", ")
                                self.FP.ProcessVal2any(node.FP.Get_func(), &node.Nodes_Node)
                            }
                            for _index, _argNode := range( argList.FP.Get_expList().Items ) {
                                index := _index + 1
                                argNode := _argNode.(Nodes_NodeDownCast).ToNodes_Node()
                                if index <= mRetExp.FP.Get_index(){
                                    self.FP.Write(", ")
                                    convCC_filter_1642_(argNode, self, &argList.Nodes_Node)
                                }
                            }
                            self.FP.Write(")")
                        }
                    }
                }
            }
        }
        if Lns_op_not(isMret){
            process()
        }
        convCC_processAlterAccessVal_2356_(self.stream.FP, funcType.FP.Get_retTypeInfoList(), node.FP.Get_expTypeList())
    }
    var retTypeInfoList *LnsList
    retTypeInfoList = funcType.FP.Get_retTypeInfoList()
    if retTypeInfoList.Len() == 1{
        convCC_processAlterToActualType_2359_(self.stream.FP, self.moduleCtrl, retTypeInfoList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), node.FP.Get_expType(), convCC_processExp_1881_(call))
    } else { 
        call()
    }
}

// 9029: decl @lune.@base.@convCC.convFilter.processExpAccessMRet
func (self *convCC_convFilter) ProcessExpAccessMRet(node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
    convCC_processGetMRet_2313_(self.stream.FP, self.moduleCtrl, node.FP.Get_expType(), node.FP.Get_index() - 1)
}

// 9037: decl @lune.@base.@convCC.convFilter.processExpList
func (self *convCC_convFilter) ProcessExpList(node *Nodes_ExpListNode,_opt LnsAny) {
    var expList *LnsList
    expList = node.FP.Get_expList()
    for _index, _exp := range( expList.Items ) {
        index := _index + 1
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        if exp.FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Abbr{
            break
        }
        if index > 1{
            self.FP.Write(", ")
        }
        convCC_filter_1642_(exp, self, &node.Nodes_Node)
    }
}

// 9053: decl @lune.@base.@convCC.convFilter.processExpOp1
func (self *convCC_convFilter) ProcessExpOp1(node *Nodes_ExpOp1Node,_opt LnsAny) {
    var op string
    op = node.FP.Get_op().Txt
    if _switch38587 := op; _switch38587 == "~" || _switch38587 == "+" || _switch38587 == "-" {
        self.FP.Write(op)
        self.FP.AccessPrimVal(node.FP.Get_exp(), &node.Nodes_Node)
    } else if _switch38587 == "not" {
        self.FP.Write("lns_op_not( _pEnv, ")
        self.FP.ProcessVal2stem(node.FP.Get_exp(), &node.Nodes_Node)
        self.FP.Write(")")
    } else if _switch38587 == "#" {
        var expType *Ast_TypeInfo
        expType = node.FP.Get_exp().FP.Get_expType().FP.Get_srcTypeInfo()
        if expType.FP.Get_kind() == Ast_TypeInfoKind__List{
            self.FP.Write("lns_mtd_List_len( _pEnv, ")
            self.FP.ProcessVal2any(node.FP.Get_exp(), &node.Nodes_Node)
            self.FP.Write(")")
        } else if expType == Ast_builtinTypeString{
            self.FP.ProcessVal2any(node.FP.Get_exp(), &node.Nodes_Node)
            self.FP.Write("->val.str.len")
        } else { 
            Util_err(Lns_getVM().String_format("not support type -- %s", []LnsAny{expType.FP.GetTxt(nil, nil, nil)}))
        }
    } else {
        Util_err(Lns_getVM().String_format("not support op -- %s", []LnsAny{op}))
    }
}

// 9091: decl @lune.@base.@convCC.convFilter.processExpMultiTo1
func (self *convCC_convFilter) ProcessExpMultiTo1(node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    self.FP.Write("lns_fromDDD( ")
    convCC_filter_1642_(node.FP.Get_exp(), self, &node.Nodes_Node)
    self.FP.Write(convCC_accessAny)
    self.FP.Write(", 0 )")
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( node.FP.Get_exp().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__DDD) &&
        Lns_GetEnv().SetStackVal( Ast_isNumberType(node.FP.Get_expType().FP.Get_srcTypeInfo().FP.Get_nonnilableType())) ).(bool)){
    } else { 
        self.FP.Write(convCC_getAccessValFromStem_1854_(node.FP.Get_exp().FP.Get_expType()))
    }
}

// 9110: decl @lune.@base.@convCC.convFilter.processStme2Val
func (self *convCC_convFilter) processStme2Val(dstType *Ast_TypeInfo,srcStemTxt string) {
    if _switch38807 := dstType; _switch38807 == Ast_builtinTypeInt || _switch38807 == Ast_builtinTypeChar {
        self.FP.Write("lns_stem2int( ")
        self.FP.Write(srcStemTxt)
        self.FP.Write(")")
    } else if _switch38807 == Ast_builtinTypeReal {
        self.FP.Write("lns_stem2real( ")
        self.FP.Write(srcStemTxt)
        self.FP.Write(")")
    } else if _switch38807 == Ast_builtinTypeBool {
        self.FP.Write("lns_stem2bool( ")
        self.FP.Write(srcStemTxt)
        self.FP.Write(")")
    } else {
        self.FP.Write(srcStemTxt)
        if convCC_getValKind_1167_(dstType) == convCC_ValKind__Any{
            self.FP.Write(convCC_accessAny)
        }
    }
}

// 9137: decl @lune.@base.@convCC.convFilter.processFuncCast
func (self *convCC_convFilter) processFuncCast(node *Nodes_ExpCastNode) {
    var castType *Ast_TypeInfo
    castType = node.FP.Get_castType()
    if _switch38835 := castType.FP.Get_kind(); _switch38835 == Ast_TypeInfoKind__Func || _switch38835 == Ast_TypeInfoKind__Form || _switch38835 == Ast_TypeInfoKind__FormFunc {
    } else {
        return 
    }
    var orgFunc *Ast_TypeInfo
    orgFunc = node.FP.Get_exp().FP.Get_expType()
    var closureSymList *LnsList
    closureSymList = Lns_unwrapDefault( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(orgFunc.FP.Get_scope()) && 
    Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.Get_closureSymList()})), NewLnsList([]LnsAny{})).(*LnsList)
    if _switch38890 := orgFunc.FP.Get_nonnilableType().FP.Get_kind(); _switch38890 == Ast_TypeInfoKind__Func || _switch38890 == Ast_TypeInfoKind__Form || _switch38890 == Ast_TypeInfoKind__FormFunc {
    } else if _switch38890 == Ast_TypeInfoKind__Stem {
        return 
    } else {
        Util_err(Lns_getVM().String_format("illegal kind -- %s, %d", []LnsAny{Ast_TypeInfoKind_getTxt( orgFunc.FP.Get_nonnilableType().FP.Get_kind()), 9161}))
    }
    if Lns_op_not(self.FP.needsWrapper(orgFunc, castType)){
        return 
    }
    self.FP.Write(Lns_getVM().String_format("static %s %s( %s _pEnv", []LnsAny{convCC_getCRetType_1195_(castType.FP.Get_retTypeInfoList()), self.moduleCtrl.FP.GetFuncCastWrapName(orgFunc, castType), convCC_cTypeEnvP}))
    if closureSymList.Len() > 0{
        self.FP.Write(Lns_getVM().String_format(", %s _pForm", []LnsAny{convCC_cTypeAnyP}))
    }
    for _index, _argType := range( castType.FP.Get_argTypeInfoList().Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        self.FP.Write(", ")
        self.FP.Write(Lns_getVM().String_format("%s arg%d", []LnsAny{convCC_getCType_1188_(argType), index}))
    }
    self.FP.Write(")")
    if self.processMode == convCC_ProcessMode__Prototype{
        self.FP.Writeln(";")
        return 
    }
    self.FP.Writeln(Lns_getVM().String_format("// %d", []LnsAny{node.FP.Get_pos().LineNo}))
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    for _index, _typeInfo := range( orgFunc.FP.Get_argTypeInfoList().Items ) {
        index := _index + 1
        typeInfo := _typeInfo.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        self.FP.Write(Lns_getVM().String_format("%s var%d = ", []LnsAny{convCC_getCType_1188_(typeInfo), index}))
        var dstType *Ast_TypeInfo
        dstType = orgFunc.FP.Get_argTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var srcType *Ast_TypeInfo
        srcType = castType.FP.Get_argTypeInfoList().GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( index == castType.FP.Get_argTypeInfoList().Len()) &&
            Lns_GetEnv().SetStackVal( (Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( orgFunc.FP.Get_argTypeInfoList().Len() != castType.FP.Get_argTypeInfoList().Len()) ||
                Lns_GetEnv().SetStackVal( Lns_op_not(dstType.FP.Equals(self.processInfo, srcType, nil, nil))) ).(bool))) ).(bool)){
            if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                Lns_GetEnv().SetStackVal( srcType.FP.Get_kind() == Ast_TypeInfoKind__DDD) &&
                Lns_GetEnv().SetStackVal( dstType.FP.Get_kind() != Ast_TypeInfoKind__DDD) ).(bool)){
                var dddSym string
                dddSym = Lns_getVM().String_format("arg%d%s", []LnsAny{index, convCC_accessAny})
                self.FP.processStme2Val(dstType, Lns_getVM().String_format("lns_fromDDD( %s, 0 )", []LnsAny{dddSym}))
                self.FP.Writeln(";")
                {
                    var _from39289 LnsInt = index + 1
                    var _to39289 LnsInt = orgFunc.FP.Get_argTypeInfoList().Len()
                    for _work39289 := _from39289; _work39289 <= _to39289; _work39289++ {
                        subIndex := _work39289
                        var dstTypeSub *Ast_TypeInfo
                        dstTypeSub = orgFunc.FP.Get_argTypeInfoList().GetAt(subIndex).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                        self.FP.Write(Lns_getVM().String_format("%s var%d = ", []LnsAny{convCC_getCType_1188_(typeInfo), subIndex}))
                        if dstTypeSub.FP.Get_kind() == Ast_TypeInfoKind__DDD{
                            self.FP.Write(Lns_getVM().String_format("lns_createSubDDD( _pEnv, %d, arg%d )", []LnsAny{subIndex - index, index}))
                        } else { 
                            var dddSymSub string
                            dddSymSub = Lns_getVM().String_format("arg%d%s", []LnsAny{index, convCC_accessAny})
                            self.FP.processStme2Val(dstTypeSub, Lns_getVM().String_format("lns_fromDDD( %s, %d )", []LnsAny{dddSymSub, subIndex - index}))
                        }
                        self.FP.Writeln(";")
                    }
                }
                break
            } else { 
                self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
            }
        } else { 
            self.FP.Write(Lns_getVM().String_format("arg%d", []LnsAny{index}))
        }
        self.FP.Writeln(";")
    }
    if orgFunc.FP.Get_retTypeInfoList().Len() > 0{
        self.FP.Write(Lns_getVM().String_format("%s ret = ", []LnsAny{convCC_getCRetType_1195_(orgFunc.FP.Get_retTypeInfoList())}))
    }
    self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{self.moduleCtrl.FP.GetFuncName(orgFunc)}))
    if closureSymList.Len() > 0{
        self.FP.Write(", _pForm")
    }
    for _index, _ := range( orgFunc.FP.Get_argTypeInfoList().Items ) {
        index := _index + 1
        self.FP.Write(Lns_getVM().String_format(", var%s", []LnsAny{index}))
    }
    self.FP.Writeln(");")
    if castType.FP.Get_retTypeInfoList().Len() > 0{
        self.FP.Write("return ")
        if orgFunc.FP.Get_retTypeInfoList().Len() == 0{
            if castType.FP.Get_retTypeInfoList().Len() > 1{
                self.FP.Write(convCC_cValDDD0)
            } else { 
                self.FP.Write(convCC_cValNil)
            }
        } else if orgFunc.FP.Get_retTypeInfoList().Len() == 1{
            if castType.FP.Get_retTypeInfoList().Len() > 1{
                if orgFunc.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind() == Ast_TypeInfoKind__DDD{
                    self.FP.Write("ret")
                } else { 
                    self.FP.Write("lns_createMRet( _pEnv, false, 1, ")
                    convCC_process2stem_2060_(self.stream.FP, self.moduleCtrl, self.scopeMgr, convCC_getValKind_1167_(orgFunc.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()), orgFunc.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), &node.Nodes_Node, convCC_process2stemCallback_2057_(func() {
                        self.FP.Write("ret")
                    }))
                    self.FP.Write(")")
                }
            } else { 
                if convCC_getValKind_1167_(castType.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()) == convCC_ValKind__Stem{
                    convCC_process2stem_2060_(self.stream.FP, self.moduleCtrl, self.scopeMgr, convCC_getValKind_1167_(orgFunc.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()), orgFunc.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), &node.Nodes_Node, convCC_process2stemCallback_2057_(func() {
                        self.FP.Write("ret")
                    }))
                } else { 
                    self.FP.Write("ret")
                }
            }
        } else { 
            self.FP.Write("ret")
        }
        self.FP.Writeln(";")
    }
    self.FP.PopIndent()
    self.FP.Writeln("}")
}

// 9304: decl @lune.@base.@convCC.convFilter.processExpCast
func (self *convCC_convFilter) ProcessExpCast(node *Nodes_ExpCastNode,_opt LnsAny) {
    if _switch39733 := self.processMode; _switch39733 == convCC_ProcessMode__Prototype || _switch39733 == convCC_ProcessMode__Intermediate {
        self.FP.processFuncCast(node)
        return 
    } else {
    }
    var exp *Nodes_Node
    exp = node.FP.Get_exp()
    var expType *Ast_TypeInfo
    expType = exp.FP.Get_expType()
    var castType *Ast_TypeInfo
    castType = node.FP.Get_castType()
    if _switch40555 := node.FP.Get_castKind(); _switch40555 == Nodes_CastKind__Implicit {
        if _switch39868 := castType.FP.Get_kind(); _switch39868 == Ast_TypeInfoKind__IF {
            if expType.FP.Get_kind() == Ast_TypeInfoKind__Class{
                self.FP.Write(Lns_getVM().String_format("lns_getIF( _pEnv, &lns_if_%s( ", []LnsAny{self.moduleCtrl.FP.GetClassCName(expType)}))
                self.FP.ProcessVal2any(node.FP.Get_exp(), &node.Nodes_Node)
                self.FP.Write(Lns_getVM().String_format(")->%s )", []LnsAny{self.moduleCtrl.FP.GetClassCName(castType)}))
            }
        } else if _switch39868 == Ast_TypeInfoKind__FormFunc {
            self.FP.processFuncCast2Form(castType, expType)
        } else if _switch39868 == Ast_TypeInfoKind__Form {
            self.FP.processFuncCast2Form(castType, expType)
        } else {
            convCC_filter_1642_(exp, self, &node.Nodes_Node)
        }
    } else if _switch40555 == Nodes_CastKind__Force {
        if _switch40067 := convCC_getValKind_1167_(castType); _switch40067 == convCC_ValKind__Stem {
            self.FP.ProcessVal2stem(exp, &node.Nodes_Node)
        } else if _switch40067 == convCC_ValKind__Any {
            self.FP.ProcessVal2any(exp, &node.Nodes_Node)
        } else if _switch40067 == convCC_ValKind__Prim {
            if convCC_isStemType_1171_(expType){
                if _switch40048 := castType.FP.Get_srcTypeInfo(); _switch40048 == Ast_builtinTypeInt || _switch40048 == Ast_builtinTypeChar {
                    self.FP.Write("lns_stem2int( ")
                    convCC_filter_1642_(exp, self, &node.Nodes_Node)
                    self.FP.Write(")")
                } else if _switch40048 == Ast_builtinTypeReal {
                    self.FP.Write("lns_stem2real( ")
                    convCC_filter_1642_(exp, self, &node.Nodes_Node)
                    self.FP.Write(")")
                } else if _switch40048 == Ast_builtinTypeBool {
                    self.FP.Write("lns_stem2bool( ")
                    convCC_filter_1642_(exp, self, &node.Nodes_Node)
                    self.FP.Write(")")
                } else {
                    Util_err(Lns_getVM().String_format("illegal cast -- %s", []LnsAny{castType.FP.GetTxt(nil, nil, nil)}))
                }
            } else { 
                convCC_filter_1642_(exp, self, &node.Nodes_Node)
            }
        }
    } else if _switch40555 == Nodes_CastKind__Normal {
        var nonNilCastType *Ast_TypeInfo
        nonNilCastType = castType.FP.Get_nonnilableType()
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( nonNilCastType.FP.Get_kind() == Ast_TypeInfoKind__Class) &&
            Lns_GetEnv().SetStackVal( Lns_op_not(nonNilCastType.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil))) ).(bool)){
            self.FP.Write("lns_castClass( ")
            self.FP.ProcessVal2stem(exp, &node.Nodes_Node)
            self.FP.Write(Lns_getVM().String_format(", &%s )", []LnsAny{self.moduleCtrl.FP.GetClassMetaName(nonNilCastType)}))
        } else if nonNilCastType.FP.Get_kind() == Ast_TypeInfoKind__IF{
            self.FP.Write("lns_castIf( _pEnv, ")
            self.FP.ProcessVal2stem(exp, &node.Nodes_Node)
            self.FP.Write(Lns_getVM().String_format(", &%s )", []LnsAny{self.moduleCtrl.FP.GetClassMetaName(nonNilCastType)}))
        } else { 
            if convCC_getValKind_1167_(nonNilCastType) == convCC_ValKind__Any{
                var kindTxt string
                var workType *Ast_TypeInfo
                {
                    _enumType := Ast_EnumTypeInfoDownCastF(nonNilCastType.FP)
                    if _enumType != nil {
                        enumType := _enumType.(*Ast_EnumTypeInfo)
                        workType = enumType.FP.Get_valTypeInfo()
                        
                    } else {
                        workType = nonNilCastType
                        
                    }
                }
                if _switch40392 := workType.FP.Get_kind(); _switch40392 == Ast_TypeInfoKind__List {
                    kindTxt = "lns_value_type_List"
                    
                } else if _switch40392 == Ast_TypeInfoKind__Array {
                    kindTxt = "lns_value_type_Array"
                    
                } else if _switch40392 == Ast_TypeInfoKind__Map {
                    kindTxt = "lns_value_type_Map"
                    
                } else if _switch40392 == Ast_TypeInfoKind__Class {
                    if workType.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil){
                        kindTxt = "lns_value_type_str"
                        
                    } else { 
                        Util_err("not support")
                    }
                } else if _switch40392 == Ast_TypeInfoKind__IF {
                    Util_err("not support")
                } else if _switch40392 == Ast_TypeInfoKind__Func {
                    kindTxt = "lns_value_type_form"
                    
                } else if _switch40392 == Ast_TypeInfoKind__Alge {
                    kindTxt = "lns_value_type_alge"
                    
                } else if _switch40392 == Ast_TypeInfoKind__DDD {
                    kindTxt = "lns_value_type_ddd"
                    
                } else if _switch40392 == Ast_TypeInfoKind__Set {
                    kindTxt = "lns_value_type_Set"
                    
                } else if _switch40392 == Ast_TypeInfoKind__Form {
                    kindTxt = "lns_value_type_form"
                    
                } else if _switch40392 == Ast_TypeInfoKind__FormFunc {
                    kindTxt = "lns_value_type_form"
                    
                } else {
                    Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{castType.FP.GetTxt(nil, nil, nil)}))
                }
                self.FP.Write("lns_castAny( ")
                self.FP.ProcessVal2stem(exp, &node.Nodes_Node)
                self.FP.Write(Lns_getVM().String_format(", %s )", []LnsAny{kindTxt}))
            } else { 
                if nonNilCastType.FP.Get_kind() != Ast_TypeInfoKind__Stem{
                    var kindTxt string
                    if _switch40498 := nonNilCastType.FP.Get_srcTypeInfo(); _switch40498 == Ast_builtinTypeInt || _switch40498 == Ast_builtinTypeChar {
                        kindTxt = "lns_stem_type_int"
                        
                    } else if _switch40498 == Ast_builtinTypeReal {
                        kindTxt = "lns_stem_type_real"
                        
                    } else if _switch40498 == Ast_builtinTypeBool {
                        kindTxt = "lns_stem_type_bool"
                        
                    } else {
                        Util_err(Lns_getVM().String_format("not support -- %s", []LnsAny{castType.FP.GetTxt(nil, nil, nil)}))
                    }
                    self.FP.Write("lns_castStem( ")
                    self.FP.ProcessVal2stem(exp, &node.Nodes_Node)
                    self.FP.Write(Lns_getVM().String_format(", %s )", []LnsAny{kindTxt}))
                } else { 
                    convCC_filter_1642_(exp, self, &node.Nodes_Node)
                }
            }
        }
    }
}

// 9486: decl @lune.@base.@convCC.convFilter.processExpParen
func (self *convCC_convFilter) ProcessExpParen(node *Nodes_ExpParenNode,_opt LnsAny) {
    if node.FP.Get_exp().FP.Get_expTypeList().Len() == 1{
        self.FP.Write("(")
        self.FP.AccessPrimVal(node.FP.Get_exp(), &node.Nodes_Node)
        self.FP.Write(" )")
    } else { 
        convCC_processToIF_2310_(self.stream.FP, self.moduleCtrl, node.FP.Get_expType(), convCC_processExp_1881_(func() {
            self.FP.AccessPrimVal(node.FP.Get_exp(), &node.Nodes_Node)
        }))
    }
}

// 9506: decl @lune.@base.@convCC.convFilter.processWrapForm2Func
func (self *convCC_convFilter) processWrapForm2Func(funcType *Ast_TypeInfo) {
    self.FP.Write(Lns_getVM().String_format("static %s _wrap_%s_%d( %s _pEnv, %s _pForm, ", []LnsAny{convCC_cTypeStem, funcType.FP.Get_rawTxt(), funcType.FP.Get_typeId(), convCC_cTypeEnvP, convCC_cTypeAnyP}))
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList().Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        self.FP.Write(Lns_getVM().String_format(", %s arg%d", []LnsAny{convCC_getCType_1188_(argType), index}))
    }
    self.FP.Writeln(")")
    self.FP.Writeln("{")
    self.FP.Writeln("return %s( _pEnv, _pForm")
    self.FP.Writeln("}")
}

// 9523: decl @lune.@base.@convCC.convFilter.processAndOr
func (self *convCC_convFilter) processAndOr(node *Nodes_ExpOp2Node,opTxt string,parent *Nodes_Node) {
    var firstFlag bool
    firstFlag = Lns_op_not(convFilter_processAndOr__isAndOr_2783_(parent))
    if firstFlag{
        self.FP.Writeln("lns_popVal( _pEnv, lns_incStack( _pEnv ) ||")
        self.FP.PushIndent(nil)
    }
    var opCC string
    if opTxt == "and"{
        opCC = "&&"
        
    } else { 
        opCC = "||"
        
    }
    if convFilter_processAndOr__isAndOr_2783_(node.FP.Get_exp1()){
        convCC_filter_1642_(node.FP.Get_exp1(), self, &node.Nodes_Node)
    } else { 
        self.FP.Write("lns_setStackVal( _pEnv, ")
        self.FP.ProcessVal2stem(node.FP.Get_exp1(), &node.Nodes_Node)
        self.FP.Write(") ")
    }
    self.FP.Writeln(opCC)
    if convFilter_processAndOr__isAndOr_2783_(node.FP.Get_exp2()){
        convCC_filter_1642_(node.FP.Get_exp2(), self, &node.Nodes_Node)
    } else { 
        self.FP.Write("lns_setStackVal( _pEnv, ")
        self.FP.ProcessVal2stem(node.FP.Get_exp2(), &node.Nodes_Node)
        self.FP.Write(") ")
    }
    if firstFlag{
        self.FP.Write(")")
        if Lns_op_not(convCC_isStemType_1171_(node.FP.Get_expType())){
            self.FP.Write(convCC_getAccessPrimValFromStem_1851_(false, node.FP.Get_expType(), 0))
        }
        self.FP.PopIndent()
    }
}

// 9580: decl @lune.@base.@convCC.convFilter.processConcat
func (self *convCC_convFilter) processConcat(node *Nodes_ExpOp2Node,parent *Nodes_Node) {
    self.FP.Write("lns_strconcat( _pEnv, ")
    self.FP.ProcessVal2any(node.FP.Get_exp1(), &node.Nodes_Node)
    self.FP.Write(", ")
    self.FP.ProcessVal2any(node.FP.Get_exp2(), &node.Nodes_Node)
    self.FP.Write(")")
}

// 9590: decl @lune.@base.@convCC.convFilter.processExpSetVal
func (self *convCC_convFilter) ProcessExpSetVal(node *Nodes_ExpSetValNode,_opt LnsAny) {
    var expList *Nodes_ExpListNode
    expList = node.FP.Get_exp2()
    var mRetExp LnsAny
    mRetExp = node.FP.Get_exp2().FP.Get_mRetExp()
    self.FP.processSetValToNode(&node.Nodes_Node, node.FP.Get_exp1(), node.FP.Get_initSymSet(), expList.FP.Get_expList(), mRetExp)
}

// 9600: decl @lune.@base.@convCC.convFilter.processExpOp2
func (self *convCC_convFilter) ProcessExpOp2(node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt.(*ConvCC_Opt)
    var opTxt string
    opTxt = node.FP.Get_op().Txt
    if _switch41498 := opTxt; _switch41498 == "and" || _switch41498 == "or" {
        self.FP.processAndOr(node, opTxt, opt.Node)
    } else if _switch41498 == ".." {
        self.FP.processConcat(node, opt.Node)
    } else {
        {
            __exp := Ast_bitBinOpMap.Items[opTxt]
            if __exp != nil {
                _exp := __exp.(LnsInt)
                if _switch41198 := _exp; _switch41198 == Ast_BitOpKind__LShift {
                    opTxt = "<<"
                    
                } else if _switch41198 == Ast_BitOpKind__RShift {
                    opTxt = ">>"
                    
                }
                convCC_filter_1642_(node.FP.Get_exp1(), self, &node.Nodes_Node)
                self.FP.Write(" " + opTxt + " ")
                convCC_filter_1642_(node.FP.Get_exp2(), self, &node.Nodes_Node)
            } else {
                if Ast_compOpSet.Has(opTxt){
                    self.FP.processEquals(opTxt == "==", node.FP.Get_exp1().FP.Get_expType(), node.FP.Get_exp2().FP.Get_expType(), convCC_ProcessToValForm_2538_(func(valKind LnsInt) {
                        if _switch41326 := valKind; _switch41326 == convCC_ValKind__Stem {
                            self.FP.ProcessVal2stem(node.FP.Get_exp1(), &node.Nodes_Node)
                        } else if _switch41326 == convCC_ValKind__Any {
                            self.FP.ProcessVal2any(node.FP.Get_exp1(), &node.Nodes_Node)
                        } else if _switch41326 == convCC_ValKind__Prim {
                            self.FP.AccessPrimVal(node.FP.Get_exp1(), &node.Nodes_Node)
                        }
                    }), convCC_ProcessToValForm_2538_(func(valKind LnsInt) {
                        if _switch41387 := valKind; _switch41387 == convCC_ValKind__Stem {
                            self.FP.ProcessVal2stem(node.FP.Get_exp2(), &node.Nodes_Node)
                        } else if _switch41387 == convCC_ValKind__Any {
                            self.FP.ProcessVal2any(node.FP.Get_exp2(), &node.Nodes_Node)
                        } else if _switch41387 == convCC_ValKind__Prim {
                            self.FP.AccessPrimVal(node.FP.Get_exp2(), &node.Nodes_Node)
                        }
                    }))
                } else if Ast_mathCompOpSet.Has(opTxt){
                    self.FP.AccessPrimVal(node.FP.Get_exp1(), &node.Nodes_Node)
                    self.FP.Write(" " + opTxt + " ")
                    self.FP.AccessPrimVal(node.FP.Get_exp2(), &node.Nodes_Node)
                } else { 
                    convCC_filter_1642_(node.FP.Get_exp1(), self, &node.Nodes_Node)
                    self.FP.Write(" " + opTxt + " ")
                    convCC_filter_1642_(node.FP.Get_exp2(), self, &node.Nodes_Node)
                }
            }
        }
    }
}

// 9722: decl @lune.@base.@convCC.convFilter.processExpRef
func (self *convCC_convFilter) ProcessExpRef(node *Nodes_ExpRefNode,_opt LnsAny) {
    if self.processMode == convCC_ProcessMode__Immediate{
        self.accessSymbolSet.FP.Add(Ast_SymbolInfo2Stem(node.FP.Get_symbolInfo()))
    }
    if node.FP.Get_symbolInfo().FP.Get_name() == "super"{
        var funcType *Ast_TypeInfo
        funcType = node.FP.Get_expType()
        self.FP.Write(Lns_getVM().String_format("%s.%s", []LnsAny{self.FP.getFullName(funcType.FP.Get_parentInfo()), funcType.FP.Get_rawTxt()}))
    } else if node.FP.Get_symbolInfo().FP.Get_name() == "__mod__"{
        self.FP.Write("*lns_module_path")
    } else if node.FP.Get_symbolInfo().FP.Get_name() == "..."{
        self.FP.Write("_pDDD")
    } else { 
        {
            _cFuncName := self.moduleCtrl.FP.GetBuiltinFuncNameFromType(node.FP.Get_expType())
            if _cFuncName != nil {
                cFuncName := _cFuncName.(string)
                self.FP.Write(cFuncName)
            } else {
                var symbolInfo *Ast_SymbolInfo
                symbolInfo = node.FP.Get_symbolInfo()
                var valKind LnsInt
                valKind = self.scopeMgr.FP.GetSymbolValKind(symbolInfo.FP)
                if valKind == convCC_ValKind__Var{
                    self.FP.Write(Lns_getVM().String_format("%s->stem", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbolInfo.FP)}))
                    self.FP.Write(convCC_getAccessValFromStem_1854_(symbolInfo.FP.Get_typeInfo()))
                } else { 
                    if Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
                        Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_kind() == Ast_SymbolKind__Fun) ||
                        Lns_GetEnv().SetStackVal( symbolInfo.FP.Get_typeInfo().FP.Get_kind() == Ast_TypeInfoKind__Func) ).(bool){
                        self.FP.Write(self.moduleCtrl.FP.GetFuncName(symbolInfo.FP.Get_typeInfo()))
                    } else { 
                        if self.FP.isManagedAnySymbol(symbolInfo.FP){
                            self.FP.Write(Lns_getVM().String_format("(*%s)", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbolInfo.FP)}))
                        } else { 
                            self.FP.Write(self.moduleCtrl.FP.GetSymbolName(symbolInfo.FP))
                        }
                    }
                }
            }
        }
    }
}

// 9776: decl @lune.@base.@convCC.convFilter.processExpRefItem
func (self *convCC_convFilter) ProcessExpRefItem(node *Nodes_ExpRefItemNode,_opt LnsAny) {
    var process func()
    process = func() {
        self.FP.Write("lns_stem_refAt( _pEnv, ")
        self.FP.ProcessVal2stem(node.FP.Get_val(), &node.Nodes_Node)
        self.FP.Write(", ")
        {
            _index := node.FP.Get_index()
            if _index != nil {
                index := _index.(*Nodes_Node)
                self.FP.ProcessVal2stem(index, &node.Nodes_Node)
            } else {
                self.FP.Write(convCC_getLiteralStrStem_1500_(Lns_getVM().String_format("\"%s\"", []LnsAny{Lns_unwrap( node.FP.Get_symbol()).(string)})))
            }
        }
        self.FP.Write(")")
    }
    if node.FP.Get_nilAccess(){
        convCC_processToIF_2310_(self.stream.FP, self.moduleCtrl, node.FP.Get_expType(), convCC_processExp_1881_(process))
        return 
    }
    var val *Nodes_Node
    val = node.FP.Get_val()
    var valType *Ast_TypeInfo
    valType = val.FP.Get_expType()
    if valType.FP.Equals(self.processInfo, Ast_builtinTypeString, nil, nil){
        self.FP.AccessPrimVal(val, &node.Nodes_Node)
        self.FP.Write("->val.str.pStr[")
        {
            _indexNode := node.FP.Get_index()
            if _indexNode != nil {
                indexNode := _indexNode.(*Nodes_Node)
                convCC_filter_1642_(indexNode, self, &node.Nodes_Node)
            } else {
                panic("index is nil")
            }
        }
        self.FP.Write("- 1 ]")
    } else if node.FP.Get_isLValue(){
        Util_err("not support -- L-Value")
    } else { 
        convCC_processToIF_2310_(self.stream.FP, self.moduleCtrl, node.FP.Get_expType(), convCC_processExp_1881_(func() {
            if _switch42178 := valType.FP.Get_kind(); _switch42178 == Ast_TypeInfoKind__List {
                self.FP.Write("lns_mtd_List_refAt( _pEnv, ")
                self.FP.ProcessVal2any(val, &node.Nodes_Node)
                self.FP.Write(", ")
                self.FP.AccessPrimVal(Lns_unwrap( node.FP.Get_index()).(*Nodes_Node), &node.Nodes_Node)
                self.FP.Write(")")
                self.FP.Write(convCC_getAccessValFromStem_1854_(valType.FP.Get_itemTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))
            } else if _switch42178 == Ast_TypeInfoKind__Map {
                self.FP.Write("lns_mtd_Map_get( _pEnv, ")
                self.FP.ProcessVal2any(val, &node.Nodes_Node)
                self.FP.Write(", ")
                {
                    _index := node.FP.Get_index()
                    if _index != nil {
                        index := _index.(*Nodes_Node)
                        self.FP.ProcessVal2stem(index, &node.Nodes_Node)
                    } else {
                        self.FP.Write(convCC_getLiteralStrStem_1500_(Lns_getVM().String_format("\"%s\"", []LnsAny{Lns_unwrap( node.FP.Get_symbol()).(string)})))
                    }
                }
                self.FP.Write(")")
            } else if _switch42178 == Ast_TypeInfoKind__Stem {
                process()
            } else {
                Util_err(Lns_getVM().String_format("not support:%s -- %d:%d", []LnsAny{Ast_TypeInfoKind_getTxt( valType.FP.Get_kind()), 9850, node.FP.Get_pos().LineNo}))
            }
        }))
    }
}

// 9857: decl @lune.@base.@convCC.convFilter.processRefField
func (self *convCC_convFilter) ProcessRefField(node *Nodes_RefFieldNode,_opt LnsAny) {
    if node.FP.Get_nilAccess(){
        {
            _symbolInfo := node.FP.Get_symbolInfo()
            if _symbolInfo != nil {
                symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
                if _switch42317 := symbolInfo.FP.Get_kind(); _switch42317 == Ast_SymbolKind__Mbr {
                    var prefixType *Ast_TypeInfo
                    prefixType = convCC_getOrgTypeInfo_1430_(node.FP.Get_prefix().FP.Get_expType())
                    if prefixType.FP.Get_kind() == Ast_TypeInfoKind__Class{
                        self.FP.Write("lns_refFieldNil( _pEnv, ")
                        self.FP.ProcessVal2stem(node.FP.Get_prefix(), &node.Nodes_Node)
                        self.FP.Write(Lns_getVM().String_format(", offsetof( %s, %s ), %s )", []LnsAny{self.moduleCtrl.FP.GetClassCName(prefixType), symbolInfo.FP.Get_name(), convCC_getStemTypeId_1887_(symbolInfo.FP.Get_typeInfo().FP.Get_srcTypeInfo())}))
                    } else { 
                        Util_err("not support -- " + prefixType.FP.GetTxt(nil, nil, nil))
                    }
                } else {
                    Util_err("not support -- " + Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind()))
                }
            } else {
                Util_err("not support")
            }
        }
        return 
    }
    {
        _symbolInfo := node.FP.Get_symbolInfo()
        if _symbolInfo != nil {
            symbolInfo := _symbolInfo.(*Ast_SymbolInfo)
            if symbolInfo.FP.Get_typeInfo().FP.Get_kind() == Ast_TypeInfoKind__Enum{
                if symbolInfo.FP.Get_kind() == Ast_SymbolKind__Mbr{
                    if symbolInfo.FP.Get_namespaceTypeInfo().FP.Get_kind() == Ast_TypeInfoKind__Enum{
                        self.FP.Write(self.moduleCtrl.FP.GetEnumTypeName(symbolInfo.FP.Get_typeInfo()))
                        self.FP.Write(Lns_getVM().String_format("__%s", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbolInfo.FP)}))
                        return 
                    }
                } else { 
                    Util_err("illegal access")
                }
            }
            if _switch42706 := symbolInfo.FP.Get_kind(); _switch42706 == Ast_SymbolKind__Mbr {
                if Ast_isClass(node.FP.Get_prefix().FP.Get_expType()){
                    if symbolInfo.FP.Get_staticFlag(){
                        var symbolName string
                        symbolName = self.moduleCtrl.FP.GetClassMemberName(symbolInfo.FP)
                        if self.FP.isManagedAnySymbol(symbolInfo.FP){
                            self.FP.Write(Lns_getVM().String_format("(*%s)", []LnsAny{symbolName}))
                        } else { 
                            self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{symbolName}))
                        }
                    } else { 
                        var className string
                        className = self.moduleCtrl.FP.GetClassCName(node.FP.Get_prefix().FP.Get_expType())
                        self.FP.Write(Lns_getVM().String_format("lns_obj_%s( ", []LnsAny{className}))
                        self.FP.ProcessVal2any(node.FP.Get_prefix(), &node.Nodes_Node)
                        self.FP.Write(Lns_getVM().String_format(")->%s", []LnsAny{node.FP.Get_field().Txt}))
                    }
                }
            } else if _switch42706 == Ast_SymbolKind__Var {
                if node.FP.Get_prefix().FP.Get_expType().FP.Get_kind() == Ast_TypeInfoKind__Module{
                    if symbolInfo.FP.Get_staticFlag(){
                        var symbolName string
                        symbolName = self.moduleCtrl.FP.GetSymbolName(symbolInfo.FP)
                        if self.FP.isManagedAnySymbol(symbolInfo.FP){
                            self.FP.Write(Lns_getVM().String_format("(*%s)", []LnsAny{symbolName}))
                        } else { 
                            self.FP.Write(Lns_getVM().String_format("%s", []LnsAny{symbolName}))
                        }
                    } else { 
                        var className string
                        className = self.moduleCtrl.FP.GetClassCName(node.FP.Get_prefix().FP.Get_expType())
                        self.FP.Write(Lns_getVM().String_format("lns_obj_%s( ", []LnsAny{className}))
                        self.FP.ProcessVal2any(node.FP.Get_prefix(), &node.Nodes_Node)
                        self.FP.Write(Lns_getVM().String_format(")->%s", []LnsAny{node.FP.Get_field().Txt}))
                    }
                }
            } else if _switch42706 == Ast_SymbolKind__Mtd {
                if Lns_op_not(symbolInfo.FP.Get_staticFlag()){
                    Util_err("not support yet. instanse method.")
                }
                self.FP.Write(self.moduleCtrl.FP.GetMethodCName(symbolInfo.FP.Get_typeInfo()))
            }
        }
    }
}

// 9976: decl @lune.@base.@convCC.convFilter.processExpOmitEnum
func (self *convCC_convFilter) ProcessExpOmitEnum(node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    self.FP.Write(self.moduleCtrl.FP.GetEnumValCName(node.FP.Get_expType(), node.FP.Get_valInfo().FP.Get_name()))
}

// 9995: decl @lune.@base.@convCC.convFilter.processGetField
func (self *convCC_convFilter) ProcessGetField(node *Nodes_GetFieldNode,_opt LnsAny) {
    var prefixNode *Nodes_Node
    prefixNode = node.FP.Get_prefix()
    var prefixType *Ast_TypeInfo
    prefixType = prefixNode.FP.Get_expType().FP.Get_nonnilableType()
    var fieldTxt string
    fieldTxt = node.FP.Get_field().Txt
    if _switch43317 := prefixType.FP.Get_kind(); _switch43317 == Ast_TypeInfoKind__Enum {
        if node.FP.Get_nilAccess(){
            Util_err(Lns_getVM().String_format("not support -- %d:%d:%s", []LnsAny{10006, node.FP.Get_pos().LineNo, fieldTxt}))
        }
        var enumFullName string
        enumFullName = self.moduleCtrl.FP.GetEnumTypeName(prefixType)
        if _switch42895 := fieldTxt; _switch42895 == "_allList" {
            self.FP.Write(Lns_getVM().String_format("%s_get__allList( _pEnv )", []LnsAny{enumFullName}))
        } else if _switch42895 == "_txt" {
            self.FP.Write(Lns_getVM().String_format("%s_get__txt( _pEnv, ", []LnsAny{enumFullName}))
            convCC_filter_1642_(prefixNode, self, &node.Nodes_Node)
            self.FP.Write(")")
        } else {
            Util_err(Lns_getVM().String_format("not support -- %d:%d:%s", []LnsAny{10020, node.FP.Get_pos().LineNo, fieldTxt}))
        }
    } else if _switch43317 == Ast_TypeInfoKind__Alge {
        if node.FP.Get_nilAccess(){
            Util_err(Lns_getVM().String_format("not support -- %d:%d:%s", []LnsAny{10027, node.FP.Get_pos().LineNo, fieldTxt}))
        }
        var algeName string
        algeName = self.moduleCtrl.FP.GetAlgeCName(prefixType)
        if _switch42996 := fieldTxt; _switch42996 == "_txt" {
            self.FP.Write(Lns_getVM().String_format("%s_get__txt( _pEnv, ", []LnsAny{algeName}))
            self.FP.ProcessVal2any(prefixNode, &node.Nodes_Node)
            self.FP.Write(")")
        } else {
            Util_err(Lns_getVM().String_format("not support -- %d:%d:%s", []LnsAny{10038, node.FP.Get_pos().LineNo, fieldTxt}))
        }
    } else if _switch43317 == Ast_TypeInfoKind__Class || _switch43317 == Ast_TypeInfoKind__IF {
        var getterType *Ast_TypeInfo
        getterType = Lns_unwrap( Lns_GetEnv().NilAccFin(Lns_GetEnv().NilAccPush(prefixType.FP.Get_scope()) && 
        Lns_NilAccCall1( Lns_GetEnv(), func () LnsAny { return Lns_GetEnv().NilAccPop().(*Ast_Scope).FP.GetTypeInfoField(Lns_getVM().String_format("get_%s", []LnsAny{fieldTxt}), true, Lns_unwrap( prefixType.FP.Get_scope()).(*Ast_Scope), convCC_scopeAccess)})/* 10043:34 */)).(*Ast_TypeInfo)
        var process func()
        process = func() {
            if node.FP.Get_nilAccess(){
                var typeInfo string
                typeInfo = convCC_getCType_1188_(getterType.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                if _switch43129 := typeInfo; _switch43129 == convCC_cTypeInt {
                    self.FP.Write("l_nil_mtd_getter_int( _pEnv, ")
                } else if _switch43129 == convCC_cTypeReal {
                    self.FP.Write("l_nil_mtd_getter_real( _pEnv, ")
                } else if _switch43129 == convCC_cTypeBool {
                    self.FP.Write("l_nil_mtd_getter_bool( _pEnv, ")
                } else if _switch43129 == convCC_cTypeAnyP {
                    self.FP.Write("l_nil_mtd_getter_any( _pEnv, ")
                } else if _switch43129 == convCC_cTypeStem {
                    self.FP.Write("l_nil_mtd_getter( _pEnv, ")
                } else {
                    Util_err(Lns_getVM().String_format("not support -- %d:%d:%s", []LnsAny{10067, node.FP.Get_pos().LineNo, fieldTxt}))
                }
                self.FP.ProcessVal2stem(prefixNode, &node.Nodes_Node)
                self.FP.Write(", ")
                self.FP.Write(self.moduleCtrl.FP.GetMethodCName(getterType))
                self.FP.Write(")")
            } else { 
                self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{self.moduleCtrl.FP.GetCallMethodCName(getterType)}))
                if Lns_op_not(getterType.FP.Get_staticFlag()){
                    self.FP.Write(", ")
                    self.FP.ProcessVal2any(prefixNode, &node.Nodes_Node)
                }
                self.FP.Write(")")
                convCC_processAlterAccessVal_2356_(self.stream.FP, getterType.FP.Get_retTypeInfoList(), node.FP.Get_expTypeList())
            }
        }
        if node.FP.Get_expTypeList().Len() == 1{
            convCC_processAlterToActualType_2359_(self.stream.FP, self.moduleCtrl, getterType.FP.Get_retTypeInfoList().GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), node.FP.Get_expType(), convCC_processExp_1881_(process))
        } else { 
            process()
        }
    } else {
        Util_err(Lns_getVM().String_format("not support -- %d:%d:%s", []LnsAny{10100, node.FP.Get_pos().LineNo, Ast_TypeInfoKind_getTxt( prefixType.FP.Get_kind())}))
    }
}

// 10129: decl @lune.@base.@convCC.convFilter.processReturn
func (self *convCC_convFilter) ProcessReturn(node *Nodes_ReturnNode,_opt LnsAny) {
    var retTypeInfoList *LnsList
    retTypeInfoList = self.routineInfoStack.FP.Current().FP.(convCC_RoutineInfoDownCast).ToconvCC_RoutineInfo().FP.Get_funcInfo().FP.Get_retTypeInfoList()
    var blockStart bool
    {
        _expListNode := node.FP.Get_expList()
        if _expListNode != nil {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            var expList *LnsList
            expList = expListNode.FP.Get_expList()
            var retKind LnsInt
            retKind = convCC_getRetKind_1178_(retTypeInfoList)
            var needSetRet bool
            needSetRet = true
            self.FP.Writeln("{")
            blockStart = true
            
            self.FP.PushIndent(nil)
            self.FP.Write(Lns_getVM().String_format("%s _ret = ", []LnsAny{convCC_getCRetType_1195_(retTypeInfoList)}))
            if retTypeInfoList.Len() >= 2{
                self.FP.processCreateMRet(retTypeInfoList, expList, &node.Nodes_Node)
            } else if retTypeInfoList.Len() == 1{
                if _switch43510 := retKind; _switch43510 == convCC_ValKind__Stem {
                    self.FP.ProcessVal2stem(expList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node(), &node.Nodes_Node)
                } else if _switch43510 == convCC_ValKind__Any {
                    self.FP.ProcessVal2any(expList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node(), &node.Nodes_Node)
                } else if _switch43510 == convCC_ValKind__Prim {
                    convCC_filter_1642_(expList.GetAt(1).(Nodes_NodeDownCast).ToNodes_Node(), self, &node.Nodes_Node)
                } else {
                    Util_err(Lns_getVM().String_format("no support -- %d", []LnsAny{10158}))
                }
            } else { 
            }
            self.FP.Writeln(";")
            if needSetRet{
                if _switch43625 := retKind; _switch43625 == convCC_ValKind__Stem {
                    if self.routineInfoStack.FP.Get_blockDepth() == 1{
                        self.FP.Writeln("lns_setRet( _pEnv, _ret );")
                    } else { 
                        self.FP.Writeln(Lns_getVM().String_format("lns_setRetAtBlock( LNS_BLOCK_AT( _pEnv, %d ), _ret );", []LnsAny{self.routineInfoStack.FP.Get_blockDepth()}))
                    }
                } else if _switch43625 == convCC_ValKind__Any {
                    if self.routineInfoStack.FP.Get_blockDepth() == 1{
                        self.FP.Writeln("lns_setRet( _pEnv, LNS_STEM_ANY( _ret ) );")
                    } else { 
                        self.FP.Writeln(Lns_getVM().String_format("lns_setRetAtBlock( LNS_BLOCK_AT( _pEnv, %d ), LNS_STEM_ANY( _ret ) );", []LnsAny{self.routineInfoStack.FP.Get_blockDepth()}))
                    }
                } else if _switch43625 == convCC_ValKind__Prim {
                } else {
                    Util_err(Lns_getVM().String_format("no support -- %d", []LnsAny{10191}))
                }
            }
        } else {
            blockStart = false
            
        }
    }
    if self.routineInfoStack.FP.Get_blockDepth() == 1{
        self.FP.Writeln("lns_leave_block( _pEnv );")
    } else { 
        self.FP.Writeln(Lns_getVM().String_format("lns_leave_blockMulti( _pEnv, %d );", []LnsAny{self.routineInfoStack.FP.Get_blockDepth()}))
    }
    if retTypeInfoList.Len() != 0{
        self.FP.Writeln("return _ret;")
    } else { 
        self.FP.Writeln("return;")
    }
    if blockStart{
        self.FP.PopIndent()
        self.FP.Writeln("}")
    }
}

// 10222: decl @lune.@base.@convCC.convFilter.processTestCase
func (self *convCC_convFilter) ProcessTestCase(node *Nodes_TestCaseNode,_opt LnsAny) {
    if Lns_op_not(self.enableTest){
        return 
    }
    var moduleName string
    moduleName = self.moduleCtrl.FP.GetFullName(self.moduleTypeInfo)
    var processDecl func()
    processDecl = func() {
        self.FP.Write(Lns_getVM().String_format("void %s__test_%s( %s _pEnv )", []LnsAny{moduleName, node.FP.Get_name().Txt, convCC_cTypeEnvP}))
    }
    if _switch43877 := self.processMode; _switch43877 == convCC_ProcessMode__Prototype {
        processDecl()
        self.FP.Writeln(";")
    } else {
        processDecl()
        self.FP.Writeln("{")
        self.FP.PushIndent(nil)
        self.FP.Writeln(Lns_getVM().String_format("printf( \"%s:\\n\" );", []LnsAny{node.FP.Get_name().Txt}))
        convCC_filter_1642_(&node.FP.Get_block().Nodes_Node, self, &node.Nodes_Node)
        self.FP.Writeln("lns_init_lune_base_Testing( _pEnv );")
        self.FP.Writeln("lune_base_Testing_outputAllResult( _pEnv, lns_io_stdout );")
        self.FP.PopIndent()
        self.FP.Writeln("}")
    }
}

// 10258: decl @lune.@base.@convCC.convFilter.processProvide
func (self *convCC_convFilter) ProcessProvide(node *Nodes_ProvideNode,_opt LnsAny) {
}

// 10263: decl @lune.@base.@convCC.convFilter.processAlias
func (self *convCC_convFilter) ProcessAlias(node *Nodes_AliasNode,_opt LnsAny) {
}

// 10273: decl @lune.@base.@convCC.convFilter.processBoxing
func (self *convCC_convFilter) ProcessBoxing(node *Nodes_BoxingNode,_opt LnsAny) {
}

// 10283: decl @lune.@base.@convCC.convFilter.processUnboxing
func (self *convCC_convFilter) ProcessUnboxing(node *Nodes_UnboxingNode,_opt LnsAny) {
}

// 10291: decl @lune.@base.@convCC.convFilter.processLiteralVal
func (self *convCC_convFilter) processLiteralVal(exp *Nodes_Node,parent *Nodes_Node) {
    if self.processMode != convCC_ProcessMode__Immediate{
        var symbolList *LnsList
        symbolList = exp.FP.GetSymbolInfo()
        if symbolList.Len() > 0{
            var valKind LnsInt
            _,valKind = self.scopeMgr.FP.GetCTypeForSym(symbolList.GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP)
            if valKind != convCC_ValKind__Prim{
                self.FP.ProcessVal2stem(exp, parent)
                return 
            }
        }
    }
    var valType *Ast_TypeInfo
    valType = exp.FP.Get_expType().FP.Get_srcTypeInfo()
    {
        _enumType := Ast_EnumTypeInfoDownCastF(valType.FP)
        if _enumType != nil {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            valType = enumType.FP.Get_valTypeInfo()
            
        }
    }
    if _switch44318 := valType; _switch44318 == Ast_builtinTypeInt || _switch44318 == Ast_builtinTypeChar {
        self.FP.Write("lns_imdInt( ")
        convCC_filter_1642_(exp, self, parent)
        self.FP.Write(")")
    } else if _switch44318 == Ast_builtinTypeReal {
        self.FP.Write("lns_imdReal( ")
        convCC_filter_1642_(exp, self, parent)
        self.FP.Write(")")
    } else if _switch44318 == Ast_builtinTypeBool {
        self.FP.Write("lns_imdBool( ")
        convCC_filter_1642_(exp, self, parent)
        self.FP.Write(")")
    } else if _switch44318 == Ast_builtinTypeString {
        {
            _strNode := Nodes_LiteralStringNodeDownCastF(exp.FP)
            if _strNode != nil {
                strNode := _strNode.(*Nodes_LiteralStringNode)
                if Lns_op_not(strNode.FP.Get_orgParam()){
                    self.FP.Write(Lns_getVM().String_format("lns_imdStr( %s )", []LnsAny{convCC_str2cstr_1219_(strNode.FP.Get_token().Txt)}))
                    return 
                }
            }
        }
        self.FP.Write("lns_imdAny( ")
        convCC_filter_1642_(exp, self, parent)
        self.FP.Write(")")
    } else {
        if _switch44316 := valType.FP.Get_kind(); _switch44316 == Ast_TypeInfoKind__List || _switch44316 == Ast_TypeInfoKind__Set || _switch44316 == Ast_TypeInfoKind__Map || _switch44316 == Ast_TypeInfoKind__Array || _switch44316 == Ast_TypeInfoKind__Class {
            self.FP.Write("lns_imdAny( ")
            convCC_filter_1642_(exp, self, parent)
            self.FP.Write(")")
        } else if _switch44316 == Ast_TypeInfoKind__DDD {
            self.FP.Write("lns_imdAny( ")
            self.FP.ProcessVal2any(exp, parent)
            self.FP.Write(")")
        } else if _switch44316 == Ast_TypeInfoKind__Alternate || _switch44316 == Ast_TypeInfoKind__Stem || _switch44316 == Ast_TypeInfoKind__Alge {
            self.FP.Write("lns_imdStem( ")
            convCC_filter_1642_(exp, self, parent)
            self.FP.Write(")")
        } else {
            Util_err(Lns_getVM().String_format("illegal type -- %s", []LnsAny{valType.FP.GetTxt(nil, nil, nil)}))
        }
    }
}

// 10366: decl @lune.@base.@convCC.convFilter.processLiteralNode
func (self *convCC_convFilter) processLiteralNode(exp *Nodes_Node,parent *Nodes_Node) {
    if _switch44415 := exp.FP.Get_kind(); _switch44415 == Nodes_NodeKind_get_LiteralList() || _switch44415 == Nodes_NodeKind_get_LiteralMap() || _switch44415 == Nodes_NodeKind_get_LiteralArray() || _switch44415 == Nodes_NodeKind_get_LiteralSet() {
        self.processingNode = exp
        
        convCC_filter_1642_(exp, self, parent)
    } else {
        self.FP.pushStream()
        convCC_filter_1642_(exp, self, parent)
        self.FP.popStream()
    }
}

// 10384: decl @lune.@base.@convCC.convFilter.processLiteralListSub
func (self *convCC_convFilter) processLiteralListSub(collectionType string,node *Nodes_Node,expListNodeOrg LnsAny,literalFuncName string) {
    if self.processedNodeSet.Has(Nodes_Node2Stem(node)){
        {
            _set := self.literalNode2AccessSymbolSet.Items[node]
            if _set != nil {
                set := _set.(*Util_OrderedSet)
                for _, _symbol := range( set.FP.Get_list().Items ) {
                    symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    self.accessSymbolSet.FP.Add(Ast_SymbolInfo2Stem(symbol))
                }
            }
        }
        return 
    }
    self.processedNodeSet.Add(Nodes_Node2Stem(node))
    var expListNode *Nodes_ExpListNode
    
    {
        _expListNode := expListNodeOrg
        if _expListNode == nil{
            return 
        } else {
            expListNode = _expListNode.(*Nodes_ExpListNode)
        }
    }
    if expListNode.FP.Get_expList().Len() == 0{
        return 
    }
    for _, _exp := range( expListNode.FP.Get_expList().Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.processLiteralNode(exp, node)
    }
    self.processingNode = node
    
    self.FP.Write(Lns_getVM().String_format("static %s %s( %s _pEnv", []LnsAny{convCC_cTypeAnyP, literalFuncName, convCC_cTypeEnvP}))
    for _, _symbol := range( self.accessSymbolSet.FP.Get_list().Items ) {
        symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        var ctype string
        ctype = convCC_convExp44564(Lns_2DDD(self.scopeMgr.FP.GetCTypeForSym(symbol.FP)))
        
        self.FP.Write(Lns_getVM().String_format(", %s %s", []LnsAny{ctype, self.moduleCtrl.FP.GetSymbolName(symbol.FP)}))
    }
    self.FP.Writeln(")")
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Write(Lns_getVM().String_format("lns_imd%s( list", []LnsAny{collectionType}))
    self.FP.PushIndent(nil)
    for _, _exp := range( expListNode.FP.Get_expList().Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        self.FP.Write(", ")
        self.FP.processLiteralVal(exp, node)
    }
    self.FP.PopIndent()
    self.FP.Writeln(");")
    self.FP.Writeln(Lns_getVM().String_format("return lns_create%s( _pEnv, list );", []LnsAny{collectionType}))
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.literalNode2AccessSymbolSet.Set(node,self.accessSymbolSet.FP.Clone())
}

// 10441: decl @lune.@base.@convCC.convFilter.processLiteralList
func (self *convCC_convFilter) ProcessLiteralList(node *Nodes_LiteralListNode,_opt LnsAny) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.processMode == convCC_ProcessMode__Immediate) &&
        Lns_GetEnv().SetStackVal( self.processingNode == &node.Nodes_Node) ).(bool)){
        self.FP.processLiteralListSub("List", &node.Nodes_Node, node.FP.Get_expList(), convCC_getLiteralListFuncName_2851_(node))
    } else { 
        if Lns_isCondTrue( node.FP.Get_expList()){
            self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{convCC_getLiteralListFuncName_2851_(node)}))
            var symbolSet *Util_OrderedSet
            
            {
                _symbolSet := self.literalNode2AccessSymbolSet.Items[&node.Nodes_Node]
                if _symbolSet == nil{
                    return 
                } else {
                    symbolSet = _symbolSet.(*Util_OrderedSet)
                }
            }
            for _, _symbol := range( symbolSet.FP.Get_list().Items ) {
                symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.Write(Lns_getVM().String_format(", %s", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbol.FP)}))
            }
            self.FP.Write(")")
        } else { 
            self.FP.Write("lns_class_List_new( _pEnv )")
        }
    }
}

// 10470: decl @lune.@base.@convCC.convFilter.processLiteralSet
func (self *convCC_convFilter) ProcessLiteralSet(node *Nodes_LiteralSetNode,_opt LnsAny) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.processMode == convCC_ProcessMode__Immediate) &&
        Lns_GetEnv().SetStackVal( self.processingNode == &node.Nodes_Node) ).(bool)){
        self.FP.processLiteralListSub("Set", &node.Nodes_Node, node.FP.Get_expList(), convCC_getLiteralSetFuncName_2879_(node))
    } else { 
        if Lns_isCondTrue( node.FP.Get_expList()){
            self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{convCC_getLiteralSetFuncName_2879_(node)}))
            var symbolSet *Util_OrderedSet
            
            {
                _symbolSet := self.literalNode2AccessSymbolSet.Items[&node.Nodes_Node]
                if _symbolSet == nil{
                    return 
                } else {
                    symbolSet = _symbolSet.(*Util_OrderedSet)
                }
            }
            for _, _symbol := range( symbolSet.FP.Get_list().Items ) {
                symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.Write(Lns_getVM().String_format(", %s", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbol.FP)}))
            }
            self.FP.Write(")")
        } else { 
            self.FP.Write("lns_class_Set_new( _pEnv )")
        }
    }
}

// 10499: decl @lune.@base.@convCC.convFilter.processLiteralMapSub
func (self *convCC_convFilter) processLiteralMapSub(node *Nodes_LiteralMapNode) {
    if self.processedNodeSet.Has(Nodes_LiteralMapNode2Stem(node)){
        {
            _set := self.literalNode2AccessSymbolSet.Items[&node.Nodes_Node]
            if _set != nil {
                set := _set.(*Util_OrderedSet)
                for _, _symbol := range( set.FP.Get_list().Items ) {
                    symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    self.accessSymbolSet.FP.Add(Ast_SymbolInfo2Stem(symbol))
                }
            }
        }
        return 
    }
    self.processedNodeSet.Add(Nodes_LiteralMapNode2Stem(node))
    var pairList *LnsList
    pairList = node.FP.Get_pairList()
    if pairList.Len() == 0{
        return 
    }
    for _, _pair := range( pairList.Items ) {
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        self.FP.processLiteralNode(pair.FP.Get_key(), &node.Nodes_Node)
        self.FP.processLiteralNode(pair.FP.Get_val(), &node.Nodes_Node)
    }
    self.processingNode = &node.Nodes_Node
    
    self.FP.Write(Lns_getVM().String_format("static %s %s( %s _pEnv", []LnsAny{convCC_cTypeAnyP, convCC_getLiteralMapFuncName_2889_(node), convCC_cTypeEnvP}))
    for _, _symbol := range( self.accessSymbolSet.FP.Get_list().Items ) {
        symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        self.FP.Write(Lns_getVM().String_format(", %s %s", []LnsAny{Lns_car(self.scopeMgr.FP.GetCTypeForSym(symbol.FP)).(string), self.moduleCtrl.FP.GetSymbolName(symbol.FP)}))
    }
    self.FP.Writeln(")")
    self.FP.Writeln("{")
    self.FP.PushIndent(nil)
    self.FP.Write("lns_imdMap( list")
    self.FP.PushIndent(nil)
    for _, _pair := range( pairList.Items ) {
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        self.FP.Writeln(", ")
        self.FP.Write("{ ")
        self.FP.processLiteralVal(pair.FP.Get_key(), &node.Nodes_Node)
        self.FP.Write(", ")
        self.FP.processLiteralVal(pair.FP.Get_val(), &node.Nodes_Node)
        self.FP.Write("} ")
    }
    self.FP.PopIndent()
    self.FP.Writeln(");")
    self.FP.Writeln("return lns_createMap( _pEnv, list );")
    self.FP.PopIndent()
    self.FP.Writeln("}")
    self.literalNode2AccessSymbolSet.Set(&node.Nodes_Node,self.accessSymbolSet.FP.Clone())
}

// 10551: decl @lune.@base.@convCC.convFilter.processLiteralMap
func (self *convCC_convFilter) ProcessLiteralMap(node *Nodes_LiteralMapNode,_opt LnsAny) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.processMode == convCC_ProcessMode__Immediate) &&
        Lns_GetEnv().SetStackVal( self.processingNode == &node.Nodes_Node) ).(bool)){
        self.FP.processLiteralMapSub(node)
    } else { 
        if node.FP.Get_pairList().Len() > 0{
            self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{convCC_getLiteralMapFuncName_2889_(node)}))
            var symbolSet *Util_OrderedSet
            
            {
                _symbolSet := self.literalNode2AccessSymbolSet.Items[&node.Nodes_Node]
                if _symbolSet == nil{
                    return 
                } else {
                    symbolSet = _symbolSet.(*Util_OrderedSet)
                }
            }
            for _, _symbol := range( symbolSet.FP.Get_list().Items ) {
                symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.Write(Lns_getVM().String_format(", %s", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbol.FP)}))
            }
            self.FP.Write(")")
        } else { 
            self.FP.Write("lns_class_Map_new( _pEnv )")
        }
    }
}

// 10578: decl @lune.@base.@convCC.convFilter.processLiteralArray
func (self *convCC_convFilter) ProcessLiteralArray(node *Nodes_LiteralArrayNode,_opt LnsAny) {
    if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
        Lns_GetEnv().SetStackVal( self.processMode == convCC_ProcessMode__Immediate) &&
        Lns_GetEnv().SetStackVal( self.processingNode == &node.Nodes_Node) ).(bool)){
        self.FP.processLiteralListSub("List", &node.Nodes_Node, node.FP.Get_expList(), convCC_getLiteralArrayFuncName_2913_(node))
    } else { 
        if Lns_isCondTrue( node.FP.Get_expList()){
            self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{convCC_getLiteralArrayFuncName_2913_(node)}))
            var symbolSet *Util_OrderedSet
            
            {
                _symbolSet := self.literalNode2AccessSymbolSet.Items[&node.Nodes_Node]
                if _symbolSet == nil{
                    return 
                } else {
                    symbolSet = _symbolSet.(*Util_OrderedSet)
                }
            }
            for _, _symbol := range( symbolSet.FP.Get_list().Items ) {
                symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.Write(Lns_getVM().String_format(", %s", []LnsAny{self.moduleCtrl.FP.GetSymbolName(symbol.FP)}))
            }
            self.FP.Write(")")
        } else { 
            self.FP.Write("lns_class_List_new( _pEnv )")
        }
    }
}

// 10603: decl @lune.@base.@convCC.convFilter.processLiteralChar
func (self *convCC_convFilter) ProcessLiteralChar(node *Nodes_LiteralCharNode,_opt LnsAny) {
    self.FP.Write(Lns_getVM().String_format("%d", []LnsAny{node.FP.Get_num()}))
}

// 10609: decl @lune.@base.@convCC.convFilter.processLiteralInt
func (self *convCC_convFilter) ProcessLiteralInt(node *Nodes_LiteralIntNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 10615: decl @lune.@base.@convCC.convFilter.processLiteralReal
func (self *convCC_convFilter) ProcessLiteralReal(node *Nodes_LiteralRealNode,_opt LnsAny) {
    self.FP.Write(node.FP.Get_token().Txt)
}

// 10621: decl @lune.@base.@convCC.convFilter.processLiteralString
func (self *convCC_convFilter) ProcessLiteralString(node *Nodes_LiteralStringNode,_opt LnsAny) {
    var txt string
    txt = convCC_str2cstr_1219_(node.FP.Get_token().Txt)
    if _switch46055 := self.processMode; _switch46055 == convCC_ProcessMode__Prototype {
        {
            _expListNode := node.FP.Get_orgParam()
            if _expListNode != nil {
                expListNode := _expListNode.(*Nodes_ExpListNode)
                if Lns_isCondTrue( expListNode.FP.Get_mRetExp()){
                    self.FP.ProcessCallWithMRet(&node.Nodes_Node, convCC_getMRetFuncName_2692_(&node.Nodes_Node), convCC_cTypeAnyP, &convCC_MRetInfo__Format{txt, expListNode}, expListNode)
                } else { 
                    self.FP.Write(Lns_getVM().String_format("static %s lns_litstr_%d( %s _pEnv", []LnsAny{convCC_cTypeAnyP, node.FP.Get_id(), convCC_cTypeEnvP}))
                    for _index, _ := range( expListNode.FP.Get_expList().Items ) {
                        index := _index + 1
                        self.FP.Write(Lns_getVM().String_format(", %s arg%d", []LnsAny{convCC_cTypeStem, index}))
                    }
                    self.FP.Writeln(");")
                }
            }
        }
        return 
    } else if _switch46055 == convCC_ProcessMode__StringFormat {
        {
            _expListNode := node.FP.Get_orgParam()
            if _expListNode != nil {
                expListNode := _expListNode.(*Nodes_ExpListNode)
                if Lns_isCondTrue( expListNode.FP.Get_mRetExp()){
                    self.FP.ProcessCallWithMRet(&node.Nodes_Node, convCC_getMRetFuncName_2692_(&node.Nodes_Node), convCC_cTypeAnyP, &convCC_MRetInfo__Format{txt, expListNode}, expListNode)
                } else { 
                    self.FP.Write(Lns_getVM().String_format("static %s lns_litstr_%d( %s _pEnv", []LnsAny{convCC_cTypeAnyP, node.FP.Get_id(), convCC_cTypeEnvP}))
                    for _index, _ := range( expListNode.FP.Get_expList().Items ) {
                        index := _index + 1
                        self.FP.Write(Lns_getVM().String_format(", %s arg%d", []LnsAny{convCC_cTypeStem, index}))
                    }
                    self.FP.Writeln(Lns_getVM().String_format(") // %d", []LnsAny{node.FP.Get_pos().LineNo}))
                    self.FP.Writeln("{")
                    self.FP.PushIndent(nil)
                    self.FP.Write("return mtd_lns_string_format( _pEnv, ")
                    self.FP.Write(convCC_getLiteralStrAny_1497_(txt))
                    self.FP.Write(", ")
                    var expList *LnsList
                    expList = expListNode.FP.Get_expList()
                    self.FP.Write("lns_createDDD")
                    var lastExp *Nodes_Node
                    lastExp = expList.GetAt(expList.Len()).(Nodes_NodeDownCast).ToNodes_Node()
                    self.FP.Write(Lns_getVM().String_format("( _pEnv, %s, %d", []LnsAny{Nodes_hasMultiValNode(lastExp), expList.Len()}))
                    {
                        var _from46025 LnsInt = 1
                        var _to46025 LnsInt = expList.Len()
                        for _work46025 := _from46025; _work46025 <= _to46025; _work46025++ {
                            index := _work46025
                            self.FP.Write(Lns_getVM().String_format(", arg%d", []LnsAny{index}))
                        }
                    }
                    self.FP.Writeln(") );")
                    self.FP.PopIndent()
                    self.FP.Writeln("}")
                }
            }
        }
        return 
    }
    {
        _expListNode := node.FP.Get_orgParam()
        if _expListNode != nil {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            {
                _mRetExp := expListNode.FP.Get_mRetExp()
                if _mRetExp != nil {
                    mRetExp := _mRetExp.(*Nodes_MRetExp)
                    self.FP.Write(Lns_getVM().String_format("%s( _pEnv", []LnsAny{convCC_getMRetFuncName_2692_(&node.Nodes_Node)}))
                    for _index, _exp := range( expListNode.FP.Get_expList().Items ) {
                        index := _index + 1
                        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
                        if index > mRetExp.FP.Get_index(){
                            break
                        }
                        self.FP.Write(", ")
                        convCC_filter_1642_(exp, self, &node.Nodes_Node)
                    }
                    self.FP.Write(")")
                } else {
                    self.FP.Write(Lns_getVM().String_format("lns_litstr_%d( _pEnv ", []LnsAny{node.FP.Get_id()}))
                    for _, _exp := range( expListNode.FP.Get_expList().Items ) {
                        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
                        self.FP.Write(", ")
                        self.FP.ProcessVal2stem(exp, &node.Nodes_Node)
                    }
                    self.FP.Write(")")
                }
            }
        } else {
            self.FP.Write(convCC_getLiteralStrAny_1497_(txt))
        }
    }
}

// 10721: decl @lune.@base.@convCC.convFilter.processLiteralBool
func (self *convCC_convFilter) ProcessLiteralBool(node *Nodes_LiteralBoolNode,_opt LnsAny) {
    if node.FP.Get_token().Txt == "true"{
        self.FP.Write("true")
    } else { 
        self.FP.Write("false")
    }
}

// 10732: decl @lune.@base.@convCC.convFilter.processLiteralNil
func (self *convCC_convFilter) ProcessLiteralNil(node *Nodes_LiteralNilNode,_opt LnsAny) {
    self.FP.Write(convCC_cValNil)
}

// 10738: decl @lune.@base.@convCC.convFilter.processBreak
func (self *convCC_convFilter) ProcessBreak(node *Nodes_BreakNode,_opt LnsAny) {
    if self.loopInfoStack.FP.Get_blockDepth() > 1{
        if self.loopInfoStack.FP.Get_blockDepth() == 2{
            self.FP.Writeln("lns_leave_block( _pEnv );")
        } else { 
            self.FP.Writeln(Lns_getVM().String_format("lns_leave_blockMulti( _pEnv, %d );", []LnsAny{self.loopInfoStack.FP.Get_blockDepth() - 1}))
        }
    }
    self.FP.Write("break;")
}

// 10753: decl @lune.@base.@convCC.convFilter.processLiteralSymbol
func (self *convCC_convFilter) ProcessLiteralSymbol(node *Nodes_LiteralSymbolNode,_opt LnsAny) {
}

// 10759: decl @lune.@base.@convCC.convFilter.processAbbr
func (self *convCC_convFilter) ProcessAbbr(node *Nodes_AbbrNode,_opt LnsAny) {
    Util_err("illegal")
}


// declaration Class -- BuiltinArgSymbolInfo
type convCC_BuiltinArgSymbolInfoMtd interface {
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
type convCC_BuiltinArgSymbolInfo struct {
    Ast_SymbolInfo
    scope *Ast_Scope
    name string
    typeInfo *Ast_TypeInfo
    convModuleParam LnsAny
    namespaceTypeInfo *Ast_TypeInfo
    FP convCC_BuiltinArgSymbolInfoMtd
}
func convCC_BuiltinArgSymbolInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convCC_BuiltinArgSymbolInfo).FP
}
type convCC_BuiltinArgSymbolInfoDownCast interface {
    ToconvCC_BuiltinArgSymbolInfo() *convCC_BuiltinArgSymbolInfo
}
func convCC_BuiltinArgSymbolInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convCC_BuiltinArgSymbolInfoDownCast)
    if ok { return work.ToconvCC_BuiltinArgSymbolInfo() }
    return nil
}
func (obj *convCC_BuiltinArgSymbolInfo) ToconvCC_BuiltinArgSymbolInfo() *convCC_BuiltinArgSymbolInfo {
    return obj
}
func NewconvCC_BuiltinArgSymbolInfo(arg1 *Ast_Scope, arg2 string, arg3 *Ast_TypeInfo, arg4 LnsAny, arg5 *Ast_TypeInfo) *convCC_BuiltinArgSymbolInfo {
    obj := &convCC_BuiltinArgSymbolInfo{}
    obj.FP = obj
    obj.Ast_SymbolInfo.FP = obj
    obj.InitconvCC_BuiltinArgSymbolInfo(arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *convCC_BuiltinArgSymbolInfo) InitconvCC_BuiltinArgSymbolInfo(arg1 *Ast_Scope, arg2 string, arg3 *Ast_TypeInfo, arg4 LnsAny, arg5 *Ast_TypeInfo) {
    self.Ast_SymbolInfo.InitAst_SymbolInfo( )
    self.scope = arg1
    self.name = arg2
    self.typeInfo = arg3
    self.convModuleParam = arg4
    self.namespaceTypeInfo = arg5
}
func (self *convCC_BuiltinArgSymbolInfo) Get_scope() *Ast_Scope{ return self.scope }
func (self *convCC_BuiltinArgSymbolInfo) Get_name() string{ return self.name }
func (self *convCC_BuiltinArgSymbolInfo) Get_typeInfo() *Ast_TypeInfo{ return self.typeInfo }
func (self *convCC_BuiltinArgSymbolInfo) Set_typeInfo(arg1 *Ast_TypeInfo){ self.typeInfo = arg1 }
func (self *convCC_BuiltinArgSymbolInfo) Get_convModuleParam() LnsAny{ return self.convModuleParam }
func (self *convCC_BuiltinArgSymbolInfo) Set_convModuleParam(arg1 LnsAny){ self.convModuleParam = arg1 }
func (self *convCC_BuiltinArgSymbolInfo) Get_namespaceTypeInfo() *Ast_TypeInfo{ return self.namespaceTypeInfo }
// 1429: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_canBeLeft
func (self *convCC_BuiltinArgSymbolInfo) Get_canBeLeft() bool {
    return false
}

// 1433: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_canBeRight
func (self *convCC_BuiltinArgSymbolInfo) Get_canBeRight() bool {
    return true
}

// 1437: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_symbolId
func (self *convCC_BuiltinArgSymbolInfo) Get_symbolId() LnsInt {
    return 0
}

// 1441: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_accessMode
func (self *convCC_BuiltinArgSymbolInfo) Get_accessMode() LnsInt {
    return Ast_AccessMode__Pub
}

// 1444: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_staticFlag
func (self *convCC_BuiltinArgSymbolInfo) Get_staticFlag() bool {
    return false
}

// 1447: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_isLazyLoad
func (self *convCC_BuiltinArgSymbolInfo) Get_isLazyLoad() bool {
    return false
}

// 1450: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_kind
func (self *convCC_BuiltinArgSymbolInfo) Get_kind() LnsInt {
    return Ast_SymbolKind__Arg
}

// 1453: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_pos
func (self *convCC_BuiltinArgSymbolInfo) Get_pos() LnsAny {
    return nil
}

// 1457: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_mutable
func (self *convCC_BuiltinArgSymbolInfo) Get_mutable() bool {
    return false
}

// 1460: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_mutMode
func (self *convCC_BuiltinArgSymbolInfo) Get_mutMode() LnsInt {
    return Ast_MutMode__IMut
}

// 1464: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_hasValueFlag
func (self *convCC_BuiltinArgSymbolInfo) Get_hasValueFlag() bool {
    return true
}

// 1467: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.set_hasValueFlag
func (self *convCC_BuiltinArgSymbolInfo) Set_hasValueFlag(arg bool) {
}

// 1469: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_hasAccessFromClosure
func (self *convCC_BuiltinArgSymbolInfo) Get_hasAccessFromClosure() bool {
    return false
}

// 1472: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.set_hasAccessFromClosure
func (self *convCC_BuiltinArgSymbolInfo) Set_hasAccessFromClosure(flag bool) {
}

// 1474: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_posForLatestMod
func (self *convCC_BuiltinArgSymbolInfo) Get_posForLatestMod() LnsAny {
    return nil
}

// 1477: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.set_posForLatestMod
func (self *convCC_BuiltinArgSymbolInfo) Set_posForLatestMod(pos LnsAny) {
}

// 1479: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.get_posForModToRef
func (self *convCC_BuiltinArgSymbolInfo) Get_posForModToRef() LnsAny {
    return nil
}

// 1482: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.set_posForModToRef
func (self *convCC_BuiltinArgSymbolInfo) Set_posForModToRef(pos LnsAny) {
}

// 1485: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.canAccess
func (self *convCC_BuiltinArgSymbolInfo) CanAccess(fromScope *Ast_Scope,access LnsInt) LnsAny {
    return &self.Ast_SymbolInfo
}

// 1491: decl @lune.@base.@convCC.BuiltinArgSymbolInfo.getOrg
func (self *convCC_BuiltinArgSymbolInfo) GetOrg() *Ast_SymbolInfo {
    return &self.Ast_SymbolInfo
}


func Lns_convCC_init() {
    if init_convCC { return }
    init_convCC = true
    convCC__mod__ = "@lune.@base.@convCC"
    Lns_InitMod()
    Lns_Ver_init()
    Lns_Ast_init()
    Lns_Nodes_init()
    Lns_Util_init()
    Lns_TransUnit_init()
    Lns_Parser_init()
    Lns_LuneControl_init()
    convCC_cTypeInt = "lns_int_t"
    convCC_cTypeReal = "lns_real_t"
    convCC_cTypeBool = "lns_bool_t"
    convCC_cTypeStem = "lns_stem_t"
    convCC_cTypeAny = "lns_any_t"
    convCC_cTypeAnyP = "lns_any_t *"
    convCC_cTypeAnyPP = "lns_any_t **"
    convCC_cTypeEnvP = "lns_env_t *"
    convCC_cTypeVarP = "lns_closureVar_t *"
    convCC_cTypeMod = "lns_module_t"
    convCC_cTypeModP = "lns_module_t *"
    convCC_cTypeBlockP = "lns_block_t *"
    convCC_cValNil = "lns_global.nilStem"
    convCC_cValNone = "lns_global.noneStem"
    convCC_cValDDD0 = "lns_global.ddd0"
    convCC_accessAny = ".val.pAny"
    convCC_stepIndent = 4
    convCC_scopeAccess = Ast_ScopeAccess__Full
    convCC_invalidSymbolId = -1
}
func init() {
    init_convCC = false
}
