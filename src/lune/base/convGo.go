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
func convGo_ProcessMode_get__allList_1032_(_env *LnsEnv) *LnsList{
    return convGo_ProcessModeList_
}
var convGo_ProcessModeMap_ = map[LnsInt]string {
  convGo_ProcessMode__DeclClass: "ProcessMode.DeclClass",
  convGo_ProcessMode__DeclTopScopeVar: "ProcessMode.DeclTopScopeVar",
  convGo_ProcessMode__Main: "ProcessMode.Main",
  convGo_ProcessMode__NonClosureFuncDecl: "ProcessMode.NonClosureFuncDecl",
}
func convGo_ProcessMode__from_1026_(_env *LnsEnv, arg1 LnsInt) LnsAny{
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
func convGo_ExpListKind_get__allList_1879_(_env *LnsEnv) *LnsList{
    return convGo_ExpListKindList_
}
var convGo_ExpListKindMap_ = map[LnsInt]string {
  convGo_ExpListKind__Conv: "ExpListKind.Conv",
  convGo_ExpListKind__Direct: "ExpListKind.Direct",
  convGo_ExpListKind__Slice: "ExpListKind.Slice",
}
func convGo_ExpListKind__from_1873_(_env *LnsEnv, arg1 LnsInt) LnsAny{
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
}
var convGo_CallKind__BuiltinCall_Obj = &convGo_CallKind__BuiltinCall{}
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
type convFilter_processRoot__ProcNode_2382_ func (_env *LnsEnv, arg1 *Nodes_Node)
// for 161
func convGo_convExp564(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 368
func convGo_convExp1424(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 372
func convGo_convExp1450(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 373
func convGo_convExp1466(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 374
func convGo_convExp1488(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 375
func convGo_convExp1504(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 378
func convGo_convExp1534(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 383
func convGo_convExp1559(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 386
func convGo_convExp1579(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 389
func convGo_convExp1599(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 711
func convGo_convExp3374(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 715
func convGo_convExp3419(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 4675
func convGo_convExp26515(arg1 []LnsAny) (bool, LnsAny) {
    return Lns_getFromMulti( arg1, 0 ).(bool), Lns_getFromMulti( arg1, 1 )
}
// for 340
func convGo_convExp1352(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 362
func convGo_convExp1370(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 367
func convGo_convExp1401(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 453
func convGo_convExp1915(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 454
func convGo_convExp1944(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 497
func convGo_convExp2134(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1709
func convGo_convExp8354(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1727
func convGo_convExp8460(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1729
func convGo_convExp8499(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1732
func convGo_convExp8537(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1749
func convGo_convExp8641(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1752
func convGo_convExp8682(arg1 []LnsAny) (string, LnsInt) {
    return Lns_getFromMulti( arg1, 0 ).(string), Lns_getFromMulti( arg1, 1 ).(LnsInt)
}
// for 1756
func convGo_convExp8720(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2007
func convGo_convExp11364(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 5435
func convGo_convExp30769(arg1 []LnsAny) bool {
    return Lns_getFromMulti( arg1, 0 ).(bool)
}
// 103: decl @lune.@base.@convGo.isMain
func convGo_isMain_1080_(_env *LnsEnv, funcType *Ast_TypeInfo) bool {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Func) &&
        _env.SetStackVal( funcType.FP.Get_rawTxt(_env) == "__main") &&
        _env.SetStackVal( funcType.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) ).(bool)){
        return true
    }
    return false
}

// 225: decl @lune.@base.@convGo.getAddEnvArg
func convGo_getAddEnvArg_1279_(_env *LnsEnv, argLen LnsInt,addEnvArg bool) string {
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

// 248: decl @lune.@base.@convGo.filter
func convGo_filter_1292_(_env *LnsEnv, node *Nodes_Node,filter *convGo_convFilter,parent *Nodes_Node) {
    node.FP.ProcessFilter(_env, &filter.Nodes_Filter, ConvGo_Opt2Stem(NewConvGo_Opt(_env, parent)))
}

// 254: decl @lune.@base.@convGo.isAnyType
func convGo_isAnyType_1296_(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    var work *Ast_TypeInfo
    work = typeInfo.FP.Get_srcTypeInfo(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_nilable(_env)) ||
        _env.SetStackVal( work == Ast_builtinTypeStem) ).(bool){
        return true
    }
    if _switch1088 := typeInfo.FP.Get_kind(_env); _switch1088 == Ast_TypeInfoKind__Stem || _switch1088 == Ast_TypeInfoKind__Alge {
        return true
    } else if _switch1088 == Ast_TypeInfoKind__Alternate {
        if typeInfo.FP.HasBase(_env){
            return convGo_isAnyType_1296_(_env, typeInfo.FP.Get_baseTypeInfo(_env))
        }
        return true
    } else if _switch1088 == Ast_TypeInfoKind__Ext {
        if typeInfo.FP.Get_extedType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Stem{
            return true
        }
    }
    return false
}

// 278: decl @lune.@base.@convGo.isClosure
func convGo_isClosure_1307_(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
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

// 313: decl @lune.@base.@convGo.concatGLSym
func convGo_concatGLSym_1325_(_env *LnsEnv, name string,external bool) string {
    if external{
        var frontChar LnsInt
        frontChar = LnsInt(name[1-1])
        var front string
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( frontChar >= 97) &&
            _env.SetStackVal( frontChar <= 122) ).(bool)){
            front = _env.LuaVM.String_format("%c", []LnsAny{65 + frontChar - 97})
            
        } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( frontChar >= 65) &&
            _env.SetStackVal( frontChar <= 90) ).(bool)){
            front = _env.LuaVM.String_sub(name,1, 1)
            
        } else { 
            front = _env.LuaVM.String_format("G%c", []LnsAny{frontChar})
            
        }
        return front + _env.LuaVM.String_sub(name,2, nil)
    }
    return name
}

// 330: decl @lune.@base.@convGo.outputGoMain
func ConvGo_outputGoMain(_env *LnsEnv, appName string,mod string,testing bool,path LnsAny,opt *Option_RuntimeOpt) LnsAny {
    var lune_path string
    lune_path = "main.go"
    if path != nil{
        path_233 := path.(string)
        if path_233 != ""{
            lune_path = path_233
            
        }
    }
    var fileObj Lns_luaStream
    
    {
        _fileObj := convGo_convExp1352(Lns_2DDD(Lns_io_open(lune_path, "w")))
        if _fileObj == nil{
            return _env.LuaVM.String_format("failed to open -- %s", []LnsAny{lune_path})
        } else {
            fileObj = _fileObj.(Lns_luaStream)
        }
    }
    var base_mainCode string
    base_mainCode = "package main\n\nimport . \"github.com/ifritJP/LuneScript/src/lune/base/runtime_go\"\n//IMPORT_MAIN:\n\n//IMPORT:\n////TEST:import . \"lns/lune/base\"\n\nfunc main() {\n    Lns_InitModOnce($opt)\n    //TEST:Lns_Testing_init()\n    Lns_init()\n    //TEST:Testing_run( \"\" )\n    //TEST:Testing_outputAllResult(Lns_io_stdout)\n}\n"
    var mainMod string
    mainMod = convGo_convExp1370(Lns_2DDD(_env.LuaVM.String_gsub(mod,".*%.", "")))
    var code string
    code = base_mainCode
    if mod != mainMod{
        var importPath string
        importPath = convGo_convExp1401(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(mod,"%.[^%.]+$", "")).(string),"%.", "/")))
        code = convGo_convExp1424(Lns_2DDD(_env.LuaVM.String_gsub(code,"//IMPORT_MAIN:", _env.LuaVM.String_format("import . \"%s/%s\"", []LnsAny{appName, importPath}))))
        
    }
    if testing{
        code = convGo_convExp1450(Lns_2DDD(_env.LuaVM.String_gsub(code,"Lns_init", _env.LuaVM.String_format("Lns_%s_init", []LnsAny{mainMod}))))
        
        code = convGo_convExp1466(Lns_2DDD(_env.LuaVM.String_gsub(code,"//TEST:", "")))
        
        code = convGo_convExp1488(Lns_2DDD(_env.LuaVM.String_gsub(code,"run%( \"\" %)", _env.LuaVM.String_format("run( \"lune.base.%s\" )", []LnsAny{mainMod}))))
        
        code = convGo_convExp1504(Lns_2DDD(_env.LuaVM.String_gsub(code,"//IMPORT:", "import . \"github.com/ifritJP/LuneScript/src/lune/base\"")))
        
    } else { 
        code = convGo_convExp1534(Lns_2DDD(_env.LuaVM.String_gsub(code,"Lns_init%(%)", _env.LuaVM.String_format("Lns_RunMain( %s___main )", []LnsAny{convGo_concatGLSym_1325_(_env, mainMod, true)}))))
        
    }
    if _switch1602 := opt.FP.Get_int2strMode(_env); _switch1602 == Option_Int2strMode__Int2strModeNeed0 {
        code = convGo_convExp1559(Lns_2DDD(_env.LuaVM.String_gsub(code,"$opt", "LnsRuntimeOpt{ Int2strModeNeed0 }")))
        
    } else if _switch1602 == Option_Int2strMode__Int2strModeUnneed0 {
        code = convGo_convExp1579(Lns_2DDD(_env.LuaVM.String_gsub(code,"$opt", "LnsRuntimeOpt{ Int2strModeUnneed0 }")))
        
    } else if _switch1602 == Option_Int2strMode__Int2strModeDepend {
        code = convGo_convExp1599(Lns_2DDD(_env.LuaVM.String_gsub(code,"$opt", "LnsRuntimeOpt{ Int2strModeDepend }")))
        
    }
    fileObj.Write(_env, code)
    fileObj.Close(_env)
    return nil
}

// 403: decl @lune.@base.@convGo.isInnerDeclType
func convGo_isInnerDeclType_1379_(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
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

// 482: decl @lune.@base.@convGo.normalizeSym
func convGo_normalizeSym_1451_(_env *LnsEnv, name string) string {
    if convGo_golanKeywordSet.Has(name){
        return _env.LuaVM.String_format("_%s", []LnsAny{name})
    }
    return name
}

// 708: decl @lune.@base.@convGo.str2gostr
func convGo_str2gostr_1644_(_env *LnsEnv, txt string) string {
    var work string
    work = txt
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(work, "^```", nil, nil))){
        work = convGo_convExp3374(Lns_2DDD(_env.LuaVM.String_gsub(_env.LuaVM.String_sub(work,4, -4),"^\n", "")))
        
        work = Parser_quoteStr(_env, work)
        
    } else if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(work, "^'", nil, nil))){
        work = convGo_convExp3419(Lns_2DDD(_env.LuaVM.String_gsub(Parser_convFromRawToStr(_env, work),"\"", "\\\"")))
        
        work = _env.LuaVM.String_format("\"%s\"", []LnsAny{work})
        
    }
    return work
}

// 723: decl @lune.@base.@convGo.getOrgTypeInfo
func convGo_getOrgTypeInfo_1660_(_env *LnsEnv, typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    {
        _enumType := Ast_EnumTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env).FP)
        if !Lns_IsNil( _enumType ) {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            return enumType.FP.Get_valTypeInfo(_env)
        }
    }
    return typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)
}


// 902: decl @lune.@base.@convGo.needConvCast
func convGo_needConvCast_1788_(_env *LnsEnv, dstType *Ast_TypeInfo,srcType *Ast_TypeInfo) bool {
    if _switch4412 := dstType.FP.Get_kind(_env); _switch4412 == Ast_TypeInfoKind__Nilable {
        return convGo_needConvCast_1788_(_env, dstType.FP.Get_nonnilableType(_env), srcType)
    } else if _switch4412 == Ast_TypeInfoKind__Class {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( dstType == Ast_builtinTypeString) ||
            _env.SetStackVal( srcType.FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env) == dstType.FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)) ).(bool){
            return false
        } else { 
            return true
        }
    } else if _switch4412 == Ast_TypeInfoKind__IF {
        return false
    } else if _switch4412 == Ast_TypeInfoKind__FormFunc {
        return true
    } else if _switch4412 == Ast_TypeInfoKind__Alternate {
        if Lns_op_not(dstType.FP.HasBase(_env)){
            return false
        } else { 
            return convGo_needConvCast_1788_(_env, dstType.FP.Get_baseTypeInfo(_env), srcType)
        }
    } else if _switch4412 == Ast_TypeInfoKind__Form {
        return true
    } else if _switch4412 == Ast_TypeInfoKind__Prim {
        if Lns_op_not(dstType.FP.Get_nilable(_env)){
            if _switch4376 := dstType; _switch4376 == Ast_builtinTypeInt {
                return true
            } else if _switch4376 == Ast_builtinTypeReal {
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

// 1076: decl @lune.@base.@convGo.getExpListKind
func convGo_getExpListKind_1885_(_env *LnsEnv, dstTypeList *LnsList,node *Nodes_ExpListNode,addEnvArg bool) LnsInt {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( addEnvArg) &&
        _env.SetStackVal( node.FP.Get_mRetExp(_env)) )){
        return convGo_ExpListKind__Conv
    }
    if dstTypeList.Len() < node.FP.Get_expList(_env).Len(){
        if dstTypeList.GetAt(dstTypeList.Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind(_env) != Ast_TypeInfoKind__DDD{
            return convGo_ExpListKind__Conv
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( dstTypeList.Len() > 1) &&
        _env.SetStackVal( node.FP.Get_mRetExp(_env)) )){
        for _, _exp := range( node.FP.Get_expList(_env).Items ) {
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            {
                _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
                if !Lns_IsNil( _castNode ) {
                    castNode := _castNode.(*Nodes_ExpCastNode)
                    if convGo_needConvCast_1788_(_env, castNode.FP.Get_castType(_env), castNode.FP.Get_exp(_env).FP.Get_expType(_env)){
                        return convGo_ExpListKind__Conv
                    }
                }
            }
        }
    }
    var lastExp *Nodes_Node
    lastExp = node.FP.Get_expList(_env).GetAt(node.FP.Get_expList(_env).Len()).(Nodes_NodeDownCast).ToNodes_Node()
    var hasAbbr bool
    if lastExp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
        hasAbbr = true
        
        if node.FP.Get_expList(_env).Len() < 2{
            return convGo_ExpListKind__Direct
        }
        lastExp = node.FP.Get_expList(_env).GetAt(node.FP.Get_expList(_env).Len() - 1).(Nodes_NodeDownCast).ToNodes_Node()
        
    } else { 
        hasAbbr = false
        
    }
    if Lns_isCondTrue( Nodes_ExpToDDDNodeDownCastF(lastExp.FP)){
        var mRetExp *Nodes_MRetExp
        
        {
            _mRetExp := node.FP.Get_mRetExp(_env)
            if _mRetExp == nil{
                return convGo_ExpListKind__Slice
            } else {
                mRetExp = _mRetExp.(*Nodes_MRetExp)
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( mRetExp.FP.Get_index(_env) == 1) &&
            _env.SetStackVal( dstTypeList.GetAt(mRetExp.FP.Get_index(_env)).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
            return convGo_ExpListKind__Slice
        }
        return convGo_ExpListKind__Conv
    }
    if lastExp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        var mRetExp *Nodes_MRetExp
        
        {
            _mRetExp := node.FP.Get_mRetExp(_env)
            if _mRetExp == nil{
                return convGo_ExpListKind__Slice
            } else {
                mRetExp = _mRetExp.(*Nodes_MRetExp)
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( mRetExp.FP.Get_index(_env) == 1) &&
            _env.SetStackVal( dstTypeList.GetAt(mRetExp.FP.Get_index(_env)).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
            return convGo_ExpListKind__Direct
        }
    } else { 
        var mRetExp *Nodes_MRetExp
        
        {
            _mRetExp := node.FP.Get_mRetExp(_env)
            if _mRetExp == nil{
                return convGo_ExpListKind__Direct
            } else {
                mRetExp = _mRetExp.(*Nodes_MRetExp)
            }
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(hasAbbr)) &&
            _env.SetStackVal( mRetExp.FP.Get_index(_env) == 1) ).(bool)){
            return convGo_ExpListKind__Direct
        }
    }
    return convGo_ExpListKind__Conv
}


// 1371: decl @lune.@base.@convGo.isRetGenerics
func convGo_isRetGenerics_2055_(_env *LnsEnv, node *Nodes_ExpCallNode) bool {
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_func(_env).FP.Get_expType(_env)
    for _index, _retType := range( funcType.FP.Get_retTypeInfoList(_env).Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( retType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
            _env.SetStackVal( Lns_op_not(convGo_isAnyType_1296_(_env, node.FP.Get_expTypeList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))) ).(bool)){
            return true
        }
    }
    return false
}

































// 3539: decl @lune.@base.@convGo.getLnsItemKind
func convGo_getLnsItemKind_3923_(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    {
        _kind := convGo_type2LnsItemKindMap.Get(typeInfo)
        if !Lns_IsNil( _kind ) {
            kind := _kind.(string)
            return kind
        }
    }
    return "LnsItemKindStem"
}





// 5060: decl @lune.@base.@convGo.convFilter.processAndOr.isAndOr
func convFilter_processAndOr__isAndOr_4990_(_env *LnsEnv, exp *Nodes_Node) bool {
    {
        _parentNode := Nodes_ExpOp2NodeDownCastF(exp.FP)
        if !Lns_IsNil( _parentNode ) {
            parentNode := _parentNode.(*Nodes_ExpOp2Node)
            if _switch28559 := parentNode.FP.Get_op(_env).Txt; _switch28559 == "and" || _switch28559 == "or" {
                return true
            }
        }
    }
    return false
}

// 5701: decl @lune.@base.@convGo.createFilter
func ConvGo_createFilter(_env *LnsEnv, enableTest bool,streamName string,stream Lns_oStream,ast *TransUnit_ASTInfo,option *ConvGo_Option) *Nodes_Filter {
    return &NewconvGo_convFilter(_env, enableTest, streamName, stream, ast, option).Nodes_Filter
}

// 5711: decl @lune.@base.@convGo.Ast2Code
func convGo_Ast2Code_5451_(_env *LnsEnv, streamName string,ast *TransUnit_ASTInfo) string {
    var stream *Util_memStream
    stream = NewUtil_memStream(_env)
    var conv *Nodes_Filter
    conv = ConvGo_createFilter(_env, false, streamName, stream.FP, ast, NewConvGo_Option(_env, "lns", "lnsgo", "", false, true))
    ast.FP.Get_node(_env).FP.ProcessFilter(_env, conv, ConvGo_Opt2Stem(NewConvGo_Opt(_env, ast.FP.Get_node(_env))))
    return stream.FP.Get_txt(_env)
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
func NewConvGo_Opt(_env *LnsEnv, arg1 *Nodes_Node) *ConvGo_Opt {
    obj := &ConvGo_Opt{}
    obj.FP = obj
    obj.InitConvGo_Opt(_env, arg1)
    return obj
}
// 43: DeclConstr
func (self *ConvGo_Opt) InitConvGo_Opt(_env *LnsEnv, parent *Nodes_Node) {
    self.Parent = parent
    
}


// declaration Class -- Env
type convGo_EnvMtd interface {
    getEnv(_env *LnsEnv, arg1 bool) string
    getLuavm(_env *LnsEnv, arg1 bool) string
    pop(_env *LnsEnv)
    push(_env *LnsEnv, arg1 *Nodes_Node)
}
type convGo_Env struct {
    stack *LnsList
    addEnvArg bool
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
func NewconvGo_Env(_env *LnsEnv, arg1 bool) *convGo_Env {
    obj := &convGo_Env{}
    obj.FP = obj
    obj.InitconvGo_Env(_env, arg1)
    return obj
}
// 63: DeclConstr
func (self *convGo_Env) InitconvGo_Env(_env *LnsEnv, addEnvArg bool) {
    self.stack = NewLnsList([]LnsAny{})
    
    self.addEnvArg = addEnvArg
    
}

// 67: decl @lune.@base.@convGo.Env.push
func (self *convGo_Env) push(_env *LnsEnv, node *Nodes_Node) {
    self.stack.Insert(Nodes_Node2Stem(node))
}

// 70: decl @lune.@base.@convGo.Env.pop
func (self *convGo_Env) pop(_env *LnsEnv) {
    self.stack.Remove(nil)
}

// 73: decl @lune.@base.@convGo.Env.getLuavm
func (self *convGo_Env) getLuavm(_env *LnsEnv, threading bool) string {
    if self.addEnvArg{
        return "_env.LuaVM"
    }
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

// 87: decl @lune.@base.@convGo.Env.getEnv
func (self *convGo_Env) getEnv(_env *LnsEnv, threading bool) string {
    if self.addEnvArg{
        return "_env"
    }
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
    Get_addEnvArg(_env *LnsEnv) bool
    Get_appName(_env *LnsEnv) string
    Get_enableRunner(_env *LnsEnv) bool
    Get_mainModule(_env *LnsEnv) string
    Get_packageName(_env *LnsEnv) string
}
type ConvGo_Option struct {
    packageName string
    appName string
    mainModule string
    addEnvArg bool
    enableRunner bool
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
func NewConvGo_Option(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 bool, arg5 bool) *ConvGo_Option {
    obj := &ConvGo_Option{}
    obj.FP = obj
    obj.InitConvGo_Option(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *ConvGo_Option) InitConvGo_Option(_env *LnsEnv, arg1 string, arg2 string, arg3 string, arg4 bool, arg5 bool) {
    self.packageName = arg1
    self.appName = arg2
    self.mainModule = arg3
    self.addEnvArg = arg4
    self.enableRunner = arg5
}
func (self *ConvGo_Option) Get_packageName(_env *LnsEnv) string{ return self.packageName }
func (self *ConvGo_Option) Get_appName(_env *LnsEnv) string{ return self.appName }
func (self *ConvGo_Option) Get_mainModule(_env *LnsEnv) string{ return self.mainModule }
func (self *ConvGo_Option) Get_addEnvArg(_env *LnsEnv) bool{ return self.addEnvArg }
func (self *ConvGo_Option) Get_enableRunner(_env *LnsEnv) bool{ return self.enableRunner }

// declaration Class -- convFilter
type convGo_convFilterMtd interface {
    concatSymWithType(_env *LnsEnv, arg1 string, arg2 *Ast_TypeInfo) string
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    expList2Slice(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 bool)
    getAccessorSym(_env *LnsEnv, arg1 LnsInt, arg2 string) string
    getAlgeSymbol(_env *LnsEnv, arg1 *Ast_AlgeValInfo) string
    getCanonicalName(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    getConstrSymbol(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    getConv2formName(_env *LnsEnv, arg1 *Nodes_Node) string
    getConvExpName(_env *LnsEnv, arg1 LnsInt, arg2 *Nodes_ExpListNode) string
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
    getVM(_env *LnsEnv, arg1 bool, arg2 *Ast_TypeInfo) LnsAny
    Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl
    IsExtSymbol(_env *LnsEnv, arg1 *Ast_SymbolInfo) bool
    IsExtType(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    IsPubSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) bool
    IsPubType(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isSameModDir(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isSamePackage(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isSamePackageExtModule(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    needConvFormFunc(_env *LnsEnv, arg1 *Nodes_ExpCastNode) bool
    needPrefix(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    OutputAccessor(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputAdvertise(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputAlter2MapFunc(_env *LnsEnv, arg1 *LnsMap)
    outputAny2Type(_env *LnsEnv, arg1 *Ast_TypeInfo)
    outputAsyncItem(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    OutputCallPrefix(_env *LnsEnv, arg1 bool, arg2 LnsInt, arg3 *Nodes_Node, arg4 LnsAny, arg5 *Ast_SymbolInfo)(bool, LnsAny)
    outputCastReceiver(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputClassType(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputConstructor(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputConvExt(_env *LnsEnv, arg1 *Nodes_Node)
    outputConvItemType(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsAny)
    outputConvItemTypeList(_env *LnsEnv, arg1 *LnsList, arg2 LnsAny)
    outputConvToForm(_env *LnsEnv, arg1 *Nodes_ExpCastNode)
    OutputDeclFunc(_env *LnsEnv, arg1 bool, arg2 LnsAny) *convGo_FuncConv
    outputDeclFuncArg(_env *LnsEnv, arg1 *Ast_TypeInfo)
    OutputDeclFuncInfo(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo)
    outputDownCast(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputDummyAbstractMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo)
    outputDummyAbstractMethodOfClass(_env *LnsEnv, arg1 *Ast_TypeInfo)
    outputDummyReturn(_env *LnsEnv, arg1 *LnsList)
    outputForeachLua(_env *LnsEnv, arg1 *Nodes_Node, arg2 bool, arg3 *Nodes_Node, arg4 *Ast_SymbolInfo, arg5 LnsAny, arg6 *Nodes_BlockNode)
    outputIFMethods(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputImplicitCast(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Nodes_Node, arg3 *Nodes_ExpCastNode)
    outputImport(_env *LnsEnv, arg1 *Nodes_ImportNode)
    outputInterfaceType(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputLetVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode)
    outputMapping(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputMethodIF(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputModule(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool)
    outputModuleImport(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputNewSetup(_env *LnsEnv, arg1 string, arg2 *Ast_TypeInfo)
    outputNilAccCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode)
    outputRetType(_env *LnsEnv, arg1 *LnsList)
    outputStaticMember(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputStem2Type(_env *LnsEnv, arg1 *Ast_TypeInfo)
    outputSymbol(_env *LnsEnv, arg1 LnsAny, arg2 string)
    outputToStem(_env *LnsEnv, arg1 *Nodes_DeclClassNode)
    outputTopScopeVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode)
    PopIndent(_env *LnsEnv)
    popProcessMode(_env *LnsEnv)
    ProcessAbbr(_env *LnsEnv, arg1 *Nodes_AbbrNode, arg2 LnsAny)
    ProcessAlias(_env *LnsEnv, arg1 *Nodes_AliasNode, arg2 LnsAny)
    processAndOr(_env *LnsEnv, arg1 *Nodes_ExpOp2Node, arg2 string, arg3 *Nodes_Node)
    ProcessApply(_env *LnsEnv, arg1 *Nodes_ApplyNode, arg2 LnsAny)
    ProcessBlankLine(_env *LnsEnv, arg1 *Nodes_BlankLineNode, arg2 LnsAny)
    ProcessBlock(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBlockSub(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsAny)
    ProcessBoxing(_env *LnsEnv, arg1 *Nodes_BoxingNode, arg2 LnsAny)
    ProcessBreak(_env *LnsEnv, arg1 *Nodes_BreakNode, arg2 LnsAny)
    processCond(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_Node)
    processConvExp(_env *LnsEnv, arg1 LnsInt, arg2 *LnsList, arg3 LnsAny, arg4 bool)
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
    ProcessEnv(_env *LnsEnv, arg1 *Nodes_EnvNode, arg2 LnsAny)
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
    ProcessNewAlgeVal(_env *LnsEnv, arg1 *Nodes_NewAlgeValNode, arg2 LnsAny)
    ProcessNone(_env *LnsEnv, arg1 *Nodes_NoneNode, arg2 LnsAny)
    ProcessProtoClass(_env *LnsEnv, arg1 *Nodes_ProtoClassNode, arg2 LnsAny)
    ProcessProtoMethod(_env *LnsEnv, arg1 *Nodes_ProtoMethodNode, arg2 LnsAny)
    ProcessProvide(_env *LnsEnv, arg1 *Nodes_ProvideNode, arg2 LnsAny)
    ProcessRefField(_env *LnsEnv, arg1 *Nodes_RefFieldNode, arg2 LnsAny)
    ProcessRefType(_env *LnsEnv, arg1 *Nodes_RefTypeNode, arg2 LnsAny)
    ProcessRepeat(_env *LnsEnv, arg1 *Nodes_RepeatNode, arg2 LnsAny)
    ProcessReturn(_env *LnsEnv, arg1 *Nodes_ReturnNode, arg2 LnsAny)
    ProcessRoot(_env *LnsEnv, arg1 *Nodes_RootNode, arg2 LnsAny)
    ProcessScope(_env *LnsEnv, arg1 *Nodes_ScopeNode, arg2 LnsAny)
    processSetFromExpList(_env *LnsEnv, arg1 string, arg2 *LnsList, arg3 *Nodes_ExpListNode, arg4 bool)
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
    PushIndent(_env *LnsEnv, arg1 LnsAny)
    pushProcessMode(_env *LnsEnv, arg1 LnsInt)
    ReturnToSource(_env *LnsEnv)
    SwitchToHeader(_env *LnsEnv)
    type2gotype(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    type2gotypeOrg(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    Write(_env *LnsEnv, arg1 string)
    Writeln(_env *LnsEnv, arg1 string)
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
    moduleType2SymbolMap *LnsMap
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
func NewconvGo_convFilter(_env *LnsEnv, arg1 bool, arg2 string, arg3 Lns_oStream, arg4 *TransUnit_ASTInfo, arg5 *ConvGo_Option) *convGo_convFilter {
    obj := &convGo_convFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvGo_convFilter(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
// advertise -- 120
func (self *convGo_convFilter) PopIndent(_env *LnsEnv) {
self.stream. FP.PopIndent( _env)
}
// advertise -- 120
func (self *convGo_convFilter) PushIndent(_env *LnsEnv, arg1 LnsAny) {
self.stream. FP.PushIndent( _env, arg1)
}
// advertise -- 120
func (self *convGo_convFilter) ReturnToSource(_env *LnsEnv) {
self.stream. FP.ReturnToSource( _env)
}
// advertise -- 120
func (self *convGo_convFilter) SwitchToHeader(_env *LnsEnv) {
self.stream. FP.SwitchToHeader( _env)
}
// advertise -- 120
func (self *convGo_convFilter) Write(_env *LnsEnv, arg1 string) {
self.stream. FP.Write( _env, arg1)
}
// advertise -- 120
func (self *convGo_convFilter) Writeln(_env *LnsEnv, arg1 string) {
self.stream. FP.Writeln( _env, arg1)
}
// 138: DeclConstr
func (self *convGo_convFilter) InitconvGo_convFilter(_env *LnsEnv, enableTest bool,streamName string,stream Lns_oStream,ast *TransUnit_ASTInfo,option *ConvGo_Option) {
    self.InitNodes_Filter(_env, true, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env))
    self.moduleType2SymbolMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.env = NewconvGo_Env(_env, option.FP.Get_addEnvArg(_env))
    
    self.option = option
    
    self.processInfo = ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env).FP.Clone(_env)
    
    self.stream = NewUtil_SimpleSourceOStream(_env, stream, nil, 4)
    
    self.processMode = convGo_ProcessMode__Main
    
    self.processModeStack = NewLnsList([]LnsAny{})
    
    self.moduleTypeInfo = ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env)
    
    self.moduleScope = Lns_unwrap( ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env)).(*Ast_Scope)
    
    self.builtin2runtime = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.type2gotypeMap = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.nodeManager = NewNodes_NodeManager(_env)
    
    self.enableTest = enableTest
    
    self.module2PackSym = NewLnsMap( map[LnsAny]LnsAny{})
    
    var modDir string
    modDir = self.moduleTypeInfo.FP.GetParentFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), false)
    self.modDir = convGo_convExp564(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(modDir,"@", "")).(string),"%.$", "")))
    
    self.noneNode = Nodes_NoneNode_create(_env, self.nodeManager, Parser_noneToken.Pos, false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeNone)}))
    
}

// 168: decl @lune.@base.@convGo.convFilter.getVM
func (self *convGo_convFilter) getVM(_env *LnsEnv, threading bool,typeInfo *Ast_TypeInfo) LnsAny {
    var txt string
    
    {
        _txt := self.builtin2runtime.Get(typeInfo)
        if _txt == nil{
            return nil
        } else {
            txt = _txt.(string)
        }
    }
    return Lns_car(_env.LuaVM.String_gsub(txt,"GETVM", self.env.FP.getLuavm(_env, threading))).(string)
}

// 175: decl @lune.@base.@convGo.convFilter.pushProcessMode
func (self *convGo_convFilter) pushProcessMode(_env *LnsEnv, mode LnsInt) {
    self.processModeStack.Insert(self.processMode)
    self.processMode = mode
    
}

// 179: decl @lune.@base.@convGo.convFilter.popProcessMode
func (self *convGo_convFilter) popProcessMode(_env *LnsEnv) {
    self.processMode = self.processModeStack.GetAt(self.processModeStack.Len()).(LnsInt)
    
    self.processModeStack.Remove(nil)
}

// 184: decl @lune.@base.@convGo.convFilter.isPubType
func (self *convGo_convFilter) IsPubType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if Ast_isPubToExternal(_env, typeInfo.FP.Get_accessMode(_env)){
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Func{
            if _switch735 := typeInfo.FP.Get_parentInfo(_env).FP.Get_kind(_env); _switch735 == Ast_TypeInfoKind__Class || _switch735 == Ast_TypeInfoKind__Enum {
                return self.FP.IsPubType(_env, typeInfo.FP.Get_parentInfo(_env))
            }
        }
        return true
    }
    return Lns_op_not(typeInfo.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil))
}

// 197: decl @lune.@base.@convGo.convFilter.isPubSym
func (self *convGo_convFilter) IsPubSym(_env *LnsEnv, symbol *Ast_SymbolInfo) bool {
    if Ast_isPubToExternal(_env, symbol.FP.Get_accessMode(_env)){
        return true
    }
    return Lns_op_not(symbol.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil))
}

// 204: decl @lune.@base.@convGo.convFilter.isExtType
func (self *convGo_convFilter) IsExtType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    return Lns_op_not(typeInfo.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil))
}

// 207: decl @lune.@base.@convGo.convFilter.isExtSymbol
func (self *convGo_convFilter) IsExtSymbol(_env *LnsEnv, symbol *Ast_SymbolInfo) bool {
    return Lns_op_not(symbol.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil))
}

// 214: decl @lune.@base.@convGo.convFilter.getEnvArgDecl
func (self *convGo_convFilter) getEnvArgDecl(_env *LnsEnv, argLen LnsInt) string {
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

// 416: decl @lune.@base.@convGo.convFilter.getCanonicalName
func (self *convGo_convFilter) getCanonicalName(_env *LnsEnv, typeInfo *Ast_TypeInfo,localFlag bool) string {
    var enumName string
    enumName = typeInfo.FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), localFlag)
    return _env.LuaVM.String_format("%s", []LnsAny{Lns_car(_env.LuaVM.String_gsub(enumName,"&", "")).(string)})
}

// 423: decl @lune.@base.@convGo.convFilter.getModuleName
func (self *convGo_convFilter) getModuleName(_env *LnsEnv, typeInfo *Ast_TypeInfo,assign bool) string {
    if typeInfo == Ast_headTypeInfo{
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
    return Lns_car(_env.LuaVM.String_gsub(moduleType.FP.Get_rawTxt(_env),"@", "")).(string)
}

// 436: decl @lune.@base.@convGo.convFilter.concatSymWithType
func (self *convGo_convFilter) concatSymWithType(_env *LnsEnv, name string,typeInfo *Ast_TypeInfo) string {
    var modName string
    modName = self.FP.getModuleName(_env, typeInfo.FP.GetModule(_env), false)
    var typeName string
    if modName == ""{
        typeName = name
        
    } else { 
        typeName = _env.LuaVM.String_format("%s_%s", []LnsAny{modName, name})
        
    }
    return convGo_concatGLSym_1325_(_env, typeName, self.FP.IsPubType(_env, typeInfo))
}

// 447: decl @lune.@base.@convGo.convFilter.isSamePackageExtModule
func (self *convGo_convFilter) isSamePackageExtModule(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    var extModuleType *Ast_NormalTypeInfo
    
    {
        _extModuleType := Ast_NormalTypeInfoDownCastF(typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env).FP)
        if _extModuleType == nil{
            Util_err(_env, _env.LuaVM.String_format("illegal type -- %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
        } else {
            extModuleType = _extModuleType.(*Ast_NormalTypeInfo)
        }
    }
    var requireParent string
    requireParent = convGo_convExp1915(Lns_2DDD(_env.LuaVM.String_gsub(extModuleType.FP.Get_requirePath(_env),"[^%.]+$", "")))
    var moduleParent string
    moduleParent = convGo_convExp1944(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(self.FP.GetFull(_env, self.moduleTypeInfo, false),"[^@%.]+$", "")).(string),"@", "")))
    return requireParent == moduleParent
}

// 459: decl @lune.@base.@convGo.convFilter.isSameModDir
func (self *convGo_convFilter) isSameModDir(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo) bool {
    if moduleTypeInfo.FP.Get_parentInfo(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo.FP.Get_parentInfo(_env), nil, nil){
        return true
    }
    return false
}

// 468: decl @lune.@base.@convGo.convFilter.isSamePackage
func (self *convGo_convFilter) isSamePackage(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__ExtModule{
        return self.FP.isSamePackageExtModule(_env, typeInfo)
    }
    return self.FP.isSameModDir(_env, typeInfo)
}

// 475: decl @lune.@base.@convGo.convFilter.needPrefix
func (self *convGo_convFilter) needPrefix(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if Ast_isBuiltin(_env, typeInfo.FP.Get_typeId(_env).Id){
        return false
    }
    return Lns_op_not(self.FP.isSamePackage(_env, typeInfo))
}

// 489: decl @lune.@base.@convGo.convFilter.getSymbol
func (self *convGo_convFilter) getSymbol(_env *LnsEnv, kind LnsAny,name string) string {
    __func__ := "@lune.@base.@convGo.convFilter.getSymbol"
    var symbolName string
    symbolName = convGo_normalizeSym_1451_(_env, name)
    switch _exp2829 := kind.(type) {
    case *convGo_SymbolKind__Arg:
        return symbolName
    case *convGo_SymbolKind__Var:
    symbolInfo := _exp2829.Val1
        var modName string
        modName = convGo_convExp2134(Lns_2DDD(_env.LuaVM.String_gsub(self.moduleTypeInfo.FP.Get_rawTxt(_env),"@", "")))
        if Lns_op_not(symbolInfo.FP.GetModule(_env).FP.Equals(_env, self.processInfo, self.moduleTypeInfo, nil, nil)){
            symbolName = _env.LuaVM.String_format("%s_%s", []LnsAny{self.FP.getModuleName(_env, symbolInfo.FP.GetModule(_env), false), symbolInfo.FP.Get_name(_env)})
            
        } else if name == "__mod__"{
            symbolName = _env.LuaVM.String_format("%s__mod__", []LnsAny{modName})
            
        } else if symbolInfo.FP.Get_scope(_env) == self.moduleScope{
            symbolName = convGo_concatGLSym_1325_(_env, _env.LuaVM.String_format("%s_", []LnsAny{modName}) + symbolName, Ast_isPubToExternal(_env, symbolInfo.FP.Get_accessMode(_env)))
            
        } else if Lns_op_not(symbolInfo.FP.Get_posForModToRef(_env)){
            if symbolName != "__func__"{
                symbolName = "_"
                
            }
        }
        if self.FP.needPrefix(_env, symbolInfo.FP.GetModule(_env)){
            symbolName = _env.LuaVM.String_format("%s.%s", []LnsAny{self.FP.getModuleName(_env, symbolInfo.FP.GetModule(_env), true), symbolName})
            
        }
    case *convGo_SymbolKind__Member:
    external := _exp2829.Val1
        symbolName = convGo_concatGLSym_1325_(_env, symbolName, external)
        
    case *convGo_SymbolKind__Func:
    typeInfo := _exp2829.Val1
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
            if _switch2341 := symbolName; _switch2341 == "_toMap" {
                return "ToMap"
            } else {
                symbolName = convGo_concatGLSym_1325_(_env, symbolName, Ast_isPubToExternal(_env, typeInfo.FP.Get_accessMode(_env)))
                
            }
        } else { 
            var prefix LnsAny
            prefix = nil
            if _switch2616 := typeInfo.FP.Get_parentInfo(_env).FP.Get_kind(_env); _switch2616 == Ast_TypeInfoKind__Module || _switch2616 == Ast_TypeInfoKind__Func || _switch2616 == Ast_TypeInfoKind__Method {
                if convGo_isInnerDeclType_1379_(_env, typeInfo){
                    if Lns_op_not(convGo_isClosure_1307_(_env, typeInfo)){
                        var parentName string
                        parentName = typeInfo.FP.GetParentFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), true)
                        symbolName = _env.LuaVM.String_format("%s_%s_%d_", []LnsAny{Lns_car(_env.LuaVM.String_gsub(parentName,"[%.@]", "_")).(string), symbolName, typeInfo.FP.Get_typeId(_env).Id})
                        
                    }
                } else { 
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(Ast_isBuiltin(_env, typeInfo.FP.Get_typeId(_env).Id))) &&
                        _env.SetStackVal( Lns_op_not(self.FP.IsPubType(_env, typeInfo))) ).(bool)){
                        symbolName = _env.LuaVM.String_format("%s_%d_", []LnsAny{symbolName, typeInfo.FP.Get_typeId(_env).Id})
                        
                    }
                    symbolName = self.FP.concatSymWithType(_env, symbolName, typeInfo)
                    
                }
            } else if _switch2616 == Ast_TypeInfoKind__Enum || _switch2616 == Ast_TypeInfoKind__Class {
                var parentInfo *Ast_TypeInfo
                parentInfo = typeInfo.FP.Get_parentInfo(_env)
                symbolName = _env.LuaVM.String_format("%s_%s", []LnsAny{self.FP.getSymbol(_env, &convGo_SymbolKind__Type{parentInfo, false}, parentInfo.FP.Get_rawTxt(_env)), symbolName})
                
                if Lns_op_not(self.FP.IsPubType(_env, typeInfo)){
                    symbolName = _env.LuaVM.String_format("%s_%d_", []LnsAny{symbolName, typeInfo.FP.Get_typeId(_env).Id})
                    
                }
            } else if _switch2616 == Ast_TypeInfoKind__ExtModule {
                symbolName = convGo_concatGLSym_1325_(_env, symbolName, true)
                
                if Lns_op_not(self.FP.isSamePackageExtModule(_env, typeInfo.FP.Get_parentInfo(_env))){
                    prefix = typeInfo.FP.Get_parentInfo(_env).FP.Get_rawTxt(_env)
                    
                }
            } else {
                Util_err(_env, _env.LuaVM.String_format("%s: not support -- %s:%s", []LnsAny{__func__, typeInfo.FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), true), Ast_TypeInfoKind_getTxt( typeInfo.FP.Get_parentInfo(_env).FP.Get_kind(_env))}))
            }
            if Lns_op_not(prefix){
                if self.FP.needPrefix(_env, typeInfo.FP.GetModule(_env)){
                    prefix = self.FP.getModuleName(_env, typeInfo, true)
                    
                }
            }
            if prefix != nil{
                prefix_356 := prefix.(string)
                symbolName = _env.LuaVM.String_format("%s.%s", []LnsAny{prefix_356, symbolName})
                
            }
        }
    case *convGo_SymbolKind__Type:
    typeInfo := _exp2829.Val1
    needPrefix := _exp2829.Val2
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__FormFunc{
            return self.FP.getSymbol(_env, &convGo_SymbolKind__Func{typeInfo}, symbolName)
        }
        var workName string
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( convGo_isInnerDeclType_1379_(_env, typeInfo)) &&
            _env.SetStackVal( Lns_op_not(Ast_isBuiltin(_env, typeInfo.FP.Get_typeId(_env).Id))) ).(bool)){
            workName = _env.LuaVM.String_format("%s%d", []LnsAny{name, typeInfo.FP.Get_typeId(_env).Id})
            
        } else { 
            workName = symbolName
            
        }
        symbolName = self.FP.concatSymWithType(_env, workName, typeInfo)
        
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( needPrefix) &&
            _env.SetStackVal( self.FP.needPrefix(_env, typeInfo.FP.GetModule(_env))) ).(bool)){
            symbolName = _env.LuaVM.String_format("%s.%s", []LnsAny{self.FP.getModuleName(_env, typeInfo, true), symbolName})
            
        }
    case *convGo_SymbolKind__Static:
    typeInfo := _exp2829.Val1
        var workName string
        workName = self.FP.getSymbol(_env, &convGo_SymbolKind__Type{typeInfo, true}, typeInfo.FP.Get_rawTxt(_env))
        symbolName = _env.LuaVM.String_format("%s__%s", []LnsAny{workName, name})
        
    case *convGo_SymbolKind__Normal:
    }
    return symbolName
}

// 614: decl @lune.@base.@convGo.convFilter.getTypeSymbol
func (self *convGo_convFilter) getTypeSymbol(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    var orgType *Ast_TypeInfo
    orgType = typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env)
    return self.FP.getSymbol(_env, &convGo_SymbolKind__Type{orgType, false}, orgType.FP.Get_rawTxt(_env))
}

// 624: decl @lune.@base.@convGo.convFilter.getTypeSymbolWithPrefix
func (self *convGo_convFilter) getTypeSymbolWithPrefix(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    var orgType *Ast_TypeInfo
    orgType = typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env)
    return self.FP.getSymbol(_env, &convGo_SymbolKind__Type{orgType, true}, orgType.FP.Get_rawTxt(_env))
}

// 630: decl @lune.@base.@convGo.convFilter.getConstrSymbol
func (self *convGo_convFilter) getConstrSymbol(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    return _env.LuaVM.String_format("Init%s", []LnsAny{self.FP.getTypeSymbol(_env, typeInfo)})
}

// 634: decl @lune.@base.@convGo.convFilter.getFuncSymbol
func (self *convGo_convFilter) getFuncSymbol(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) &&
        _env.SetStackVal( typeInfo.FP.Get_staticFlag(_env)) ).(bool)){
        return self.FP.getSymbol(_env, &convGo_SymbolKind__Static{typeInfo.FP.Get_parentInfo(_env)}, typeInfo.FP.Get_rawTxt(_env))
    }
    return self.FP.getSymbol(_env, &convGo_SymbolKind__Func{typeInfo}, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( typeInfo.FP.Get_rawTxt(_env) == "") &&
        _env.SetStackVal( "_anonymous") ||
        _env.SetStackVal( typeInfo.FP.Get_rawTxt(_env)) ).(string))
}

// 643: decl @lune.@base.@convGo.convFilter.getAlgeSymbol
func (self *convGo_convFilter) getAlgeSymbol(_env *LnsEnv, valInfo *Ast_AlgeValInfo) string {
    return self.FP.getSymbol(_env, &convGo_SymbolKind__Static{&valInfo.FP.Get_algeTpye(_env).Ast_TypeInfo}, valInfo.FP.Get_name(_env))
}

// 647: decl @lune.@base.@convGo.convFilter.getSymbolSym
func (self *convGo_convFilter) getSymbolSym(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) string {
    if _switch3184 := symbolInfo.FP.Get_kind(_env); _switch3184 == Ast_SymbolKind__Fun || _switch3184 == Ast_SymbolKind__Mtd {
        return self.FP.getFuncSymbol(_env, symbolInfo.FP.Get_typeInfo(_env))
    } else if _switch3184 == Ast_SymbolKind__Arg {
        return self.FP.getSymbol(_env, convGo_SymbolKind__Arg_Obj, symbolInfo.FP.Get_name(_env))
    } else if _switch3184 == Ast_SymbolKind__Var {
        return self.FP.getSymbol(_env, &convGo_SymbolKind__Var{symbolInfo}, symbolInfo.FP.Get_name(_env))
    } else if _switch3184 == Ast_SymbolKind__Mbr {
        if symbolInfo.FP.Get_staticFlag(_env){
            return self.FP.getSymbol(_env, &convGo_SymbolKind__Static{symbolInfo.FP.Get_namespaceTypeInfo(_env)}, symbolInfo.FP.Get_name(_env))
        }
        return self.FP.getSymbol(_env, &convGo_SymbolKind__Member{Ast_isPubToExternal(_env, symbolInfo.FP.Get_accessMode(_env))}, symbolInfo.FP.Get_name(_env))
    } else if _switch3184 == Ast_SymbolKind__Typ {
        if Lns_isCondTrue( Ast_AliasTypeInfoDownCastF(symbolInfo.FP.Get_typeInfo(_env).FP)){
            return self.FP.getSymbol(_env, &convGo_SymbolKind__Var{symbolInfo}, symbolInfo.FP.Get_name(_env))
        }
        return self.FP.getTypeSymbol(_env, symbolInfo.FP.Get_typeInfo(_env))
    } else {
        Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{Ast_SymbolKind_getTxt( symbolInfo.FP.Get_kind(_env))}))
    }
// insert a dummy
    return ""
}

// 678: decl @lune.@base.@convGo.convFilter.getAccessorSym
func (self *convGo_convFilter) getAccessorSym(_env *LnsEnv, accessMode LnsInt,name string) string {
    return self.FP.getSymbol(_env, &convGo_SymbolKind__Member{Ast_isPubToExternal(_env, accessMode)}, name)
}

// 682: decl @lune.@base.@convGo.convFilter.outputSymbol
func (self *convGo_convFilter) outputSymbol(_env *LnsEnv, kind LnsAny,name string) {
    self.FP.Write(_env, self.FP.getSymbol(_env, kind, name))
}

// 686: decl @lune.@base.@convGo.convFilter.getConv2formName
func (self *convGo_convFilter) getConv2formName(_env *LnsEnv, node *Nodes_Node) string {
    return _env.LuaVM.String_format("conv2Form%d", []LnsAny{node.FP.Get_id(_env)})
}

// 690: decl @lune.@base.@convGo.convFilter.getConvGenericsName
func (self *convGo_convFilter) getConvGenericsName(_env *LnsEnv, node *Nodes_Node) string {
    return _env.LuaVM.String_format("lns_convGenerics%d", []LnsAny{node.FP.Get_id(_env)})
}

// 695: decl @lune.@base.@convGo.convFilter.getModuleSym
func (self *convGo_convFilter) getModuleSym(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,addDot bool) string {
    {
        _packSym := self.module2PackSym.Get(moduleTypeInfo)
        if !Lns_IsNil( _packSym ) {
            packSym := _packSym.(string)
            if addDot{
                return _env.LuaVM.String_format("%s.", []LnsAny{packSym})
            }
            return packSym
        }
    }
    return ""
}

// 730: decl @lune.@base.@convGo.convFilter.type2gotypeOrg
func (self *convGo_convFilter) type2gotypeOrg(_env *LnsEnv, typeInfo *Ast_TypeInfo,addClassAster bool) string {
    if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        return "[]LnsAny"
    }
    if convGo_isAnyType_1296_(_env, typeInfo){
        return "LnsAny"
    }
    var orgType *Ast_TypeInfo
    orgType = convGo_getOrgTypeInfo_1660_(_env, typeInfo)
    {
        _goType := self.type2gotypeMap.Get(orgType)
        if !Lns_IsNil( _goType ) {
            goType := _goType.(string)
            return goType
        }
    }
    if _switch3676 := orgType.FP.Get_kind(_env); _switch3676 == Ast_TypeInfoKind__Ext {
        if orgType.FP.Get_extedType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Stem{
            return "LnsAny"
        }
        return "*Lns_luaValue"
    } else if _switch3676 == Ast_TypeInfoKind__List || _switch3676 == Ast_TypeInfoKind__Array {
        return "*LnsList"
    } else if _switch3676 == Ast_TypeInfoKind__Set {
        return "*LnsSet"
    } else if _switch3676 == Ast_TypeInfoKind__Map {
        return "*LnsMap"
    } else if _switch3676 == Ast_TypeInfoKind__Form {
        return "LnsForm"
    } else if _switch3676 == Ast_TypeInfoKind__FormFunc {
        return self.FP.getFuncSymbol(_env, typeInfo)
    } else if _switch3676 == Ast_TypeInfoKind__Class {
        if typeInfo.FP.Get_genSrcTypeInfo(_env) == TransUnit_getBuiltinFunc(_env).G__pipe_{
            return "*Lns__pipe"
        }
        var symbol string
        symbol = self.FP.getTypeSymbolWithPrefix(_env, typeInfo)
        if addClassAster{
            return "*" + symbol
        }
        return symbol
    } else if _switch3676 == Ast_TypeInfoKind__IF {
        return self.FP.getTypeSymbolWithPrefix(_env, typeInfo)
    } else if _switch3676 == Ast_TypeInfoKind__Alternate {
        return self.FP.type2gotypeOrg(_env, typeInfo.FP.Get_baseTypeInfo(_env), addClassAster)
    }
    Util_err(_env, _env.LuaVM.String_format("not support yet -- %s", []LnsAny{typeInfo.FP.GetTxt(_env, nil, nil, nil)}))
// insert a dummy
    return ""
}

// 788: decl @lune.@base.@convGo.convFilter.type2gotype
func (self *convGo_convFilter) type2gotype(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    return self.FP.type2gotypeOrg(_env, typeInfo, true)
}

// 802: decl @lune.@base.@convGo.convFilter.outputAny2Type
func (self *convGo_convFilter) outputAny2Type(_env *LnsEnv, dstType *Ast_TypeInfo) {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(convGo_isAnyType_1296_(_env, dstType))) &&
        _env.SetStackVal( dstType.FP.Get_kind(_env) != Ast_TypeInfoKind__Alternate) ).(bool)){
        self.FP.Write(_env, _env.LuaVM.String_format(".(%s)", []LnsAny{self.FP.type2gotype(_env, dstType)}))
    }
}

// 809: decl @lune.@base.@convGo.convFilter.outputStem2Type
func (self *convGo_convFilter) outputStem2Type(_env *LnsEnv, dstType *Ast_TypeInfo) {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( dstType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
        _env.SetStackVal( dstType.FP.HasBase(_env)) ).(bool)){
        self.FP.Write(_env, _env.LuaVM.String_format(".(%s)", []LnsAny{self.FP.type2gotype(_env, dstType)}))
    } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( dstType.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
        _env.SetStackVal( dstType != Ast_builtinTypeString) ).(bool)){
        self.FP.Write(_env, _env.LuaVM.String_format(".(%sDownCast).To%s()", []LnsAny{self.FP.getTypeSymbolWithPrefix(_env, dstType), self.FP.getTypeSymbol(_env, dstType)}))
    } else { 
        self.FP.outputAny2Type(_env, dstType)
    }
}

// 835: decl @lune.@base.@convGo.convFilter.processBlankLine
func (self *convGo_convFilter) ProcessBlankLine(_env *LnsEnv, node *Nodes_BlankLineNode,_opt LnsAny) {
}

// 835: decl @lune.@base.@convGo.convFilter.processNone
func (self *convGo_convFilter) ProcessNone(_env *LnsEnv, node *Nodes_NoneNode,_opt LnsAny) {
}

// 844: decl @lune.@base.@convGo.convFilter.processImport
func (self *convGo_convFilter) ProcessImport(_env *LnsEnv, node *Nodes_ImportNode,_opt LnsAny) {
    var args string
    if self.option.FP.Get_addEnvArg(_env){
        args = "_env"
        
    } else { 
        args = ""
        
    }
    if node.FP.Get_modulePath(_env) == "lune.base.Depend"{
        self.FP.Writeln(_env, _env.LuaVM.String_format("Lns_LuaVer_init( %s )", []LnsAny{args}))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(self.FP.isSameModDir(_env, node.FP.Get_moduleTypeInfo(_env)))) &&
        _env.SetStackVal( Lns_op_not(Ast_isBuiltin(_env, node.FP.Get_moduleTypeInfo(_env).FP.Get_typeId(_env).Id))) ).(bool)){
        self.FP.Write(_env, _env.LuaVM.String_format("%s.", []LnsAny{self.FP.getModuleName(_env, node.FP.Get_moduleTypeInfo(_env), true)}))
    }
    self.FP.Writeln(_env, _env.LuaVM.String_format("Lns_%s_init(%s)", []LnsAny{self.FP.getModuleName(_env, node.FP.Get_moduleTypeInfo(_env), false), args}))
}

// 869: decl @lune.@base.@convGo.convFilter.needConvFormFunc
func (self *convGo_convFilter) needConvFormFunc(_env *LnsEnv, node *Nodes_ExpCastNode) bool {
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
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(argType.FP.Equals(_env, self.processInfo, funcType.FP.Get_argTypeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), nil, nil)){
            return true
        }
    }
    for _index, _retType := range( castType.FP.Get_retTypeInfoList(_env).Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_op_not(retType.FP.Equals(_env, self.processInfo, funcType.FP.Get_retTypeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), nil, nil)){
            return true
        }
    }
    return false
}

// 963: decl @lune.@base.@convGo.convFilter.outputImplicitCast
func (self *convGo_convFilter) outputImplicitCast(_env *LnsEnv, castType *Ast_TypeInfo,node *Nodes_Node,parent *Nodes_ExpCastNode) {
    if _switch5077 := castType.FP.Get_kind(_env); _switch5077 == Ast_TypeInfoKind__Nilable {
        self.FP.outputImplicitCast(_env, castType.FP.Get_nonnilableType(_env), node, parent)
    } else if _switch5077 == Ast_TypeInfoKind__Class {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( castType == Ast_builtinTypeString) ||
            _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) ||
            _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env) == castType.FP.Get_genSrcTypeInfo(_env).FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)) ).(bool){
            convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
        } else { 
            if convGo_isAnyType_1296_(_env, node.FP.Get_expType(_env)){
                self.FP.Write(_env, _env.LuaVM.String_format("%sDownCastF(", []LnsAny{self.FP.getTypeSymbolWithPrefix(_env, castType)}))
                convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
                self.FP.Write(_env, ")")
            } else { 
                self.FP.Write(_env, "&")
                convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
                self.FP.Write(_env, _env.LuaVM.String_format(".%s", []LnsAny{self.FP.getTypeSymbol(_env, castType)}))
            }
        }
    } else if _switch5077 == Ast_TypeInfoKind__IF {
        convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
        if Ast_isClass(_env, node.FP.Get_expType(_env)){
            self.FP.Write(_env, ".FP")
        }
    } else if _switch5077 == Ast_TypeInfoKind__FormFunc {
        _ = node.FP.Get_expType(_env)
        if self.FP.needConvFormFunc(_env, parent){
            self.FP.Write(_env, _env.LuaVM.String_format("%s(", []LnsAny{self.FP.getConv2formName(_env, &parent.Nodes_Node)}))
            convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
            self.FP.Write(_env, ")")
        } else { 
            self.FP.Write(_env, _env.LuaVM.String_format("%s(", []LnsAny{self.FP.getTypeSymbol(_env, castType)}))
            convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
            self.FP.Write(_env, ")")
        }
    } else if _switch5077 == Ast_TypeInfoKind__Alternate {
        if Lns_op_not(castType.FP.HasBase(_env)){
            if Ast_isClass(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env)){
                self.FP.Write(_env, _env.LuaVM.String_format("%s2Stem(", []LnsAny{self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env))}))
                convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
                self.FP.Write(_env, ")")
            } else { 
                convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
            }
        } else { 
            self.FP.outputImplicitCast(_env, castType.FP.Get_baseTypeInfo(_env), node, parent)
        }
    } else if _switch5077 == Ast_TypeInfoKind__Form {
        self.FP.Write(_env, _env.LuaVM.String_format("%s(", []LnsAny{self.FP.getConv2formName(_env, &parent.Nodes_Node)}))
        convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
        self.FP.Write(_env, ")")
    } else if _switch5077 == Ast_TypeInfoKind__Prim {
        if Lns_op_not(node.FP.Get_expType(_env).FP.Get_nilable(_env)){
            if _switch4959 := castType; _switch4959 == Ast_builtinTypeInt {
                self.FP.Write(_env, "LnsInt(")
                convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
                self.FP.Write(_env, ")")
            } else if _switch4959 == Ast_builtinTypeReal {
                self.FP.Write(_env, "LnsReal(")
                convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
                self.FP.Write(_env, ")")
            } else {
                convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
            }
        } else { 
            convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
        }
    } else {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_nilable(_env)) &&
            _env.SetStackVal( Ast_isClass(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env))) ).(bool)){
            self.FP.Write(_env, _env.LuaVM.String_format("%s2Stem(", []LnsAny{self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_expType(_env).FP.Get_nonnilableType(_env))}))
            convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
            self.FP.Write(_env, ")")
        } else { 
            convGo_filter_1292_(_env, node, self, &parent.Nodes_Node)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
                _env.SetStackVal( node.FP.Get_expType(_env) != Ast_builtinTypeString) ).(bool)){
                self.FP.Write(_env, ".FP")
            }
        }
    }
}

// 1149: decl @lune.@base.@convGo.convFilter.getConvExpName
func (self *convGo_convFilter) getConvExpName(_env *LnsEnv, nodeId LnsInt,argListNode *Nodes_ExpListNode) string {
    return _env.LuaVM.String_format("%s_convExp%d", []LnsAny{Lns_car(_env.LuaVM.String_gsub(self.moduleTypeInfo.FP.Get_rawTxt(_env),"@", "")).(string), nodeId})
}

// 1154: decl @lune.@base.@convGo.convFilter.processConvExp
func (self *convGo_convFilter) processConvExp(_env *LnsEnv, nodeId LnsInt,dstTypeList *LnsList,argListNode LnsAny,hasRetEnv bool) {
    var argList *Nodes_ExpListNode
    
    {
        _argList := argListNode
        if _argList == nil{
            return 
        } else {
            argList = _argList.(*Nodes_ExpListNode)
        }
    }
    if convGo_getExpListKind_1885_(_env, dstTypeList, argList, self.option.FP.Get_addEnvArg(_env)) != convGo_ExpListKind__Conv{
        return 
    }
    var expList *LnsList
    expList = argList.FP.Get_expList(_env)
    var mRetIndex LnsAny
    mRetIndex = _env.NilAccFin(_env.NilAccPush(argList.FP.Get_mRetExp(_env)) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
    if Lns_op_not(mRetIndex){
        if expList.GetAt(expList.Len()).(Nodes_NodeDownCast).ToNodes_Node().FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
        } else { 
            return 
        }
    }
    var workList *LnsList
    workList = NewLnsList([]LnsAny{})
    for _, _exp := range( expList.Items ) {
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        var workExp *Nodes_Node
        workExp = Nodes_getUnwraped(_env, exp)
        if workExp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
            break
        }
        workList.Insert(Nodes_Node2Stem(workExp))
    }
    expList = workList
    
    self.FP.Writeln(_env, _env.LuaVM.String_format("// for %d", []LnsAny{argList.FP.Get_pos(_env).LineNo}))
    self.FP.Write(_env, _env.LuaVM.String_format("func %s(", []LnsAny{self.FP.getConvExpName(_env, nodeId, argList)}))
    if hasRetEnv{
        self.FP.Write(_env, self.FP.getEnvArgDecl(_env, expList.Len()))
    }
    for _index, _argExp := range( expList.Items ) {
        index := _index + 1
        argExp := _argExp.(Nodes_NodeDownCast).ToNodes_Node()
        {
            _exp2ddd := Nodes_ExpToDDDNodeDownCastF(argExp.FP)
            if !Lns_IsNil( _exp2ddd ) {
                exp2ddd := _exp2ddd.(*Nodes_ExpToDDDNode)
                for _, _exp := range( exp2ddd.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
                    exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
                    if index != 1{
                        self.FP.Write(_env, ", ")
                    }
                    self.FP.Write(_env, _env.LuaVM.String_format("arg%d ", []LnsAny{index}))
                    self.FP.Write(_env, self.FP.type2gotype(_env, exp.FP.Get_expType(_env)))
                }
            } else {
                if index != 1{
                    self.FP.Write(_env, ", ")
                }
                if mRetIndex == index{
                    self.FP.Write(_env, _env.LuaVM.String_format("arg%d []LnsAny", []LnsAny{index}))
                    break
                } else { 
                    self.FP.Write(_env, _env.LuaVM.String_format("arg%d ", []LnsAny{index}))
                    {
                        _castNode := Nodes_ExpCastNodeDownCastF(argExp.FP)
                        if !Lns_IsNil( _castNode ) {
                            castNode := _castNode.(*Nodes_ExpCastNode)
                            self.FP.Write(_env, self.FP.type2gotype(_env, castNode.FP.Get_castType(_env)))
                        } else {
                            self.FP.Write(_env, self.FP.type2gotype(_env, argExp.FP.Get_expType(_env)))
                        }
                    }
                }
            }
        }
    }
    self.FP.Write(_env, ") ")
    var getRetType func(_env *LnsEnv, retType *Ast_TypeInfo,index LnsInt) *Ast_TypeInfo
    getRetType = func(_env *LnsEnv, retType *Ast_TypeInfo,index LnsInt) *Ast_TypeInfo {
        if retType == Ast_builtinTypeEmpty{
            return argList.FP.GetExpTypeNoDDDAt(_env, index)
        }
        return retType
    }
    var retTypeList *LnsList
    retTypeList = NewLnsList([]LnsAny{})
    for _index, _dstType := range( dstTypeList.Items ) {
        index := _index + 1
        dstType := _dstType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        retTypeList.Insert(Ast_TypeInfo2Stem(getRetType(_env, dstType, index)))
    }
    var needRetEnv bool
    needRetEnv = false
    if retTypeList.Len() >= 2{
        self.FP.Write(_env, "(")
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( hasRetEnv) &&
            _env.SetStackVal( self.option.FP.Get_addEnvArg(_env)) ).(bool)){
            self.FP.Write(_env, "*LnsEnv, ")
            needRetEnv = true
            
        }
        for _index, _retType := range( retTypeList.Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(_env, ", ")
            }
            self.FP.Write(_env, self.FP.type2gotype(_env, getRetType(_env, retType, index)))
        }
        self.FP.Writeln(_env, ") {")
    } else if retTypeList.Len() == 1{
        self.FP.Write(_env, self.FP.type2gotype(_env, getRetType(_env, retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo(), 1)))
        self.FP.Writeln(_env, " {")
    } else { 
        self.FP.Writeln(_env, " {")
    }
    self.FP.Write(_env, "    return ")
    if needRetEnv{
        self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, retTypeList.Len(), self.option.FP.Get_addEnvArg(_env)))
    }
    if mRetIndex != nil{
        mRetIndex_720 := mRetIndex.(LnsInt)
        var restIndex LnsAny
        restIndex = nil
        for _index, _retType := range( retTypeList.Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(_env, ", ")
            }
            if retType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                restIndex = index
                
                break
            }
            if index >= mRetIndex_720{
                var valTxt string
                valTxt = _env.LuaVM.String_format("Lns_getFromMulti( arg%d, %d )", []LnsAny{mRetIndex_720, index - mRetIndex_720})
                var wrote bool
                wrote = false
                if index <= expList.Len(){
                    var exp *Nodes_Node
                    exp = expList.GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                    {
                        _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
                        if !Lns_IsNil( _castNode ) {
                            castNode := _castNode.(*Nodes_ExpCastNode)
                            var srcTxt string
                            if castNode.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_nilable(_env){
                                srcTxt = valTxt
                                
                            } else { 
                                srcTxt = _env.LuaVM.String_format("%s.(%s)", []LnsAny{valTxt, self.FP.type2gotype(_env, castNode.FP.Get_exp(_env).FP.Get_expType(_env))})
                                
                            }
                            var statNode *Nodes_ConvStatNode
                            statNode = Nodes_ConvStatNode_create(_env, self.nodeManager, exp.FP.Get_pos(_env), false, NewLnsList([]LnsAny{Ast_TypeInfo2Stem(exp.FP.Get_expType(_env))}), srcTxt)
                            self.FP.outputImplicitCast(_env, castNode.FP.Get_castType(_env), &statNode.Nodes_Node, castNode)
                            wrote = true
                            
                        }
                    }
                }
                if Lns_op_not(wrote){
                    self.FP.Write(_env, valTxt)
                    self.FP.outputAny2Type(_env, retType)
                }
            } else { 
                self.FP.Write(_env, _env.LuaVM.String_format("arg%d", []LnsAny{index}))
            }
        }
        if restIndex != nil{
            restIndex_741 := restIndex.(LnsInt)
            self.FP.Write(_env, "Lns_2DDD( ")
            for _index, _ := range( expList.Items ) {
                index := _index + 1
                if index >= restIndex_741{
                    if index < expList.Len(){
                        self.FP.Write(_env, _env.LuaVM.String_format("arg%d", []LnsAny{index}))
                    } else { 
                        self.FP.Write(_env, _env.LuaVM.String_format("arg%d[%d:]", []LnsAny{mRetIndex_720, index - mRetIndex_720}))
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
                self.FP.Write(_env, ", ")
            }
            if index <= expList.Len(){
                self.FP.Write(_env, _env.LuaVM.String_format("arg%d", []LnsAny{index}))
            } else { 
                self.FP.Write(_env, "nil")
            }
        }
        self.FP.Writeln(_env, "")
    }
    self.FP.Writeln(_env, "}")
}

// 1341: decl @lune.@base.@convGo.convFilter.outputNilAccCall
func (self *convGo_convFilter) outputNilAccCall(_env *LnsEnv, node *Nodes_ExpCallNode) {
    if Lns_op_not(node.FP.HasNilAccess(_env)){
        return 
    }
    if node.FP.Get_expTypeList(_env).Len() > convGo_MaxNilAccNum{
        var anys string
        anys = "LnsAny"
        var nils string
        nils = "nil"
        var lists string
        lists = "list[0]"
        {
            var _from6470 LnsInt = 2
            var _to6470 LnsInt = node.FP.Get_expTypeList(_env).Len()
            for _work6470 := _from6470; _work6470 <= _to6470; _work6470++ {
                count := _work6470
                anys = _env.LuaVM.String_format("%s,LnsAny", []LnsAny{anys})
                
                nils = _env.LuaVM.String_format("%s,nil", []LnsAny{nils})
                
                lists = _env.LuaVM.String_format("%s,list[%d]", []LnsAny{lists, count - 1})
                
            }
        }
        var name string
        name = _env.LuaVM.String_format("%s_%d", []LnsAny{Lns_car(_env.LuaVM.String_gsub(self.moduleTypeInfo.FP.Get_rawTxt(_env),"@", "")).(string), node.FP.Get_id(_env)})
        self.FP.Write(_env, _env.LuaVM.String_format("func lns_NilAccCall_%s( env *LnsEnv, call func () (%s) ) bool {\n    return env.NilAccPush( Lns_2DDD( call() ) )\n}\nfunc lns_NilAccFinCall_%s( ret LnsAny ) (%s) {\n    if Lns_IsNil( ret ) {\n        return %s\n    }\n    list := ret.([]LnsAny)\n    return %s\n}\n", []LnsAny{name, anys, name, anys, nils, lists}))
    }
}

// 1382: decl @lune.@base.@convGo.convFilter.processGenericsCall
func (self *convGo_convFilter) processGenericsCall(_env *LnsEnv, node *Nodes_ExpCallNode) {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(convGo_isRetGenerics_2055_(_env, node))) ||
        _env.SetStackVal( node.FP.Get_expTypeList(_env).Len() < 2) ).(bool){
        return 
    }
    var srcTypeList *LnsList
    srcTypeList = node.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_retTypeInfoList(_env)
    var dstTypeList *LnsList
    dstTypeList = node.FP.Get_expTypeList(_env)
    var srcTxt string
    srcTxt = _env.LuaVM.String_format("arg1 %s", []LnsAny{self.FP.type2gotype(_env, srcTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
    var dstTxt string
    dstTxt = _env.LuaVM.String_format("%s", []LnsAny{self.FP.type2gotype(_env, dstTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
    {
        var _from6678 LnsInt = 2
        var _to6678 LnsInt = srcTypeList.Len()
        for _work6678 := _from6678; _work6678 <= _to6678; _work6678++ {
            index := _work6678
            srcTxt = _env.LuaVM.String_format("%s,arg%d %s", []LnsAny{srcTxt, index, self.FP.type2gotype(_env, srcTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
            
        }
    }
    {
        var _from6710 LnsInt = 2
        var _to6710 LnsInt = dstTypeList.Len()
        for _work6710 := _from6710; _work6710 <= _to6710; _work6710++ {
            index := _work6710
            dstTxt = _env.LuaVM.String_format("%s,%s", []LnsAny{dstTxt, self.FP.type2gotype(_env, dstTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo())})
            
        }
    }
    self.FP.Writeln(_env, _env.LuaVM.String_format("func %s(%s) (%s) {", []LnsAny{self.FP.getConvGenericsName(_env, &node.Nodes_Node), srcTxt, dstTxt}))
    self.FP.PushIndent(_env, nil)
    self.FP.Write(_env, "return ")
    for _index, _dstType := range( dstTypeList.Items ) {
        index := _index + 1
        dstType := _dstType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        if index > srcTypeList.Len(){
            self.FP.Write(_env, "nil")
        } else { 
            self.FP.Write(_env, _env.LuaVM.String_format("arg%d", []LnsAny{index}))
            var srcType *Ast_TypeInfo
            srcType = srcTypeList.GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if srcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                self.FP.outputAny2Type(_env, dstType)
            }
        }
    }
    self.FP.Writeln(_env, "")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 1426: decl @lune.@base.@convGo.convFilter.outputRetType
func (self *convGo_convFilter) outputRetType(_env *LnsEnv, retTypeList *LnsList) {
    if _switch6974 := retTypeList.Len(); _switch6974 == 0 {
        self.FP.Write(_env, "")
    } else if _switch6974 == 1 {
        if retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeNeverRet{
            self.FP.Write(_env, " " + self.FP.type2gotype(_env, retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()))
        }
    } else {
        self.FP.Write(_env, "(")
        for _index, _retType := range( retTypeList.Items ) {
            index := _index + 1
            retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(_env, ", ")
            }
            self.FP.Write(_env, self.FP.type2gotype(_env, retType))
        }
        self.FP.Write(_env, ")")
    }
}

// 1459: decl @lune.@base.@convGo.convFilter.outputDeclFunc
func (self *convGo_convFilter) OutputDeclFunc(_env *LnsEnv, addEnvArg bool,funcInfo LnsAny) *convGo_FuncConv {
    var typeInfo *Ast_TypeInfo
    var name LnsAny
    var prefixType *Ast_TypeInfo
    var extFlag bool
    switch _exp7169 := funcInfo.(type) {
    case *convGo_FuncInfo__DeclInfo:
    node := _exp7169.Val1
    workDeclInfo := _exp7169.Val2
        extFlag = false
        
        typeInfo = node.FP.Get_expType(_env)
        
        prefixType = typeInfo.FP.Get_parentInfo(_env)
        
        if Lns_op_not(workDeclInfo.FP.Get_name(_env)){
            if self.processMode == convGo_ProcessMode__NonClosureFuncDecl{
                name = "_anonymous"
                
            } else { 
                name = nil
                
            }
        } else { 
            name = typeInfo.FP.Get_rawTxt(_env)
            
        }
    case *convGo_FuncInfo__Type:
    workTypeInfo := _exp7169.Val1
        extFlag = workTypeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext
        
        typeInfo = workTypeInfo
        
        prefixType = typeInfo.FP.Get_parentInfo(_env)
        
        name = typeInfo.FP.Get_rawTxt(_env)
        
    case *convGo_FuncInfo__WithClass:
    classType := _exp7169.Val1
    methodType := _exp7169.Val2
        extFlag = false
        
        typeInfo = methodType
        
        prefixType = classType
        
        name = typeInfo.FP.Get_rawTxt(_env)
        
    }
    if convGo_isClosure_1307_(_env, typeInfo){
        self.FP.Write(_env, "func")
    } else { 
        if typeInfo.FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
            self.FP.Write(_env, "func ")
            self.FP.Write(_env, "(self *")
            self.FP.Write(_env, self.FP.getTypeSymbol(_env, prefixType))
            self.FP.Write(_env, ") ")
        } else { 
            self.FP.Write(_env, "func ")
        }
        if typeInfo.FP.Get_extedType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__FormFunc{
            if name != nil{
                name_856 := name.(string)
                self.FP.outputSymbol(_env, &convGo_SymbolKind__Func{typeInfo}, name_856)
            }
        }
    }
    self.FP.Write(_env, "(")
    var workType *Ast_TypeInfo
    
    {
        _workType := typeInfo.FP.GetOverridingType(_env)
        if _workType == nil{
            workType = typeInfo
            
        } else {
            workType = _workType.(*Ast_TypeInfo)
        }
    }
    var retTypeList *LnsList
    if extFlag{
        retTypeList = Lns_unwrap( Lns_car(Ast_convToExtTypeList(_env, self.processInfo, workType.FP.Get_retTypeInfoList(_env)))).(*LnsList)
        
    } else { 
        retTypeList = workType.FP.Get_retTypeInfoList(_env)
        
    }
    var funcConv *convGo_FuncConv
    funcConv = NewconvGo_FuncConv(_env, retTypeList)
    if addEnvArg{
        self.FP.Write(_env, self.FP.getEnvArgDecl(_env, workType.FP.Get_argTypeInfoList(_env).Len()))
    }
    switch _exp7577 := funcInfo.(type) {
    case *convGo_FuncInfo__DeclInfo:
    node := _exp7577.Val1
    declInfo := _exp7577.Val2
        for _index, _arg := range( declInfo.FP.Get_argList(_env).Items ) {
            index := _index + 1
            arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
            if index != 1{
                self.FP.Write(_env, ",")
            }
            var argType *Ast_TypeInfo
            argType = workType.FP.Get_argTypeInfoList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if argType.FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                {
                    _argNode := Nodes_DeclArgNodeDownCastF(arg.FP)
                    if !Lns_IsNil( _argNode ) {
                        argNode := _argNode.(*Nodes_DeclArgNode)
                        var argName string
                        argName = self.FP.getSymbolSym(_env, argNode.FP.Get_symbolInfo(_env))
                        self.FP.Write(_env, _env.LuaVM.String_format("_%s ", []LnsAny{argName}))
                        self.FP.Write(_env, self.FP.type2gotype(_env, argType))
                        funcConv.FP.Get_argList(_env).Insert(Ast_SymbolInfo2Stem(argNode.FP.Get_symbolInfo(_env)))
                    } else {
                        convGo_filter_1292_(_env, arg, self, node)
                    }
                }
            } else { 
                convGo_filter_1292_(_env, arg, self, node)
            }
        }
    case *convGo_FuncInfo__Type:
        for _index, _argType := range( workType.FP.Get_argTypeInfoList(_env).Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(_env, ",")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("arg%d %s", []LnsAny{index, self.FP.type2gotype(_env, argType)}))
        }
    case *convGo_FuncInfo__WithClass:
        for _index, _argType := range( workType.FP.Get_argTypeInfoList(_env).Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(_env, ",")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("arg%d %s", []LnsAny{index, self.FP.type2gotype(_env, argType)}))
        }
    }
    self.FP.Write(_env, ")")
    self.FP.outputRetType(_env, retTypeList)
    return funcConv
}

// 1582: decl @lune.@base.@convGo.convFilter.outputConvToForm
func (self *convGo_convFilter) outputConvToForm(_env *LnsEnv, node *Nodes_ExpCastNode) {
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
        self.FP.Writeln(_env, _env.LuaVM.String_format("      \nfunc %s( luaform LnsAny ) LnsForm {\n    return func (argList []LnsAny) []LnsAny {\n        return %s.RunLoadedfunc( luaform.(*Lns_luaValue), argList )\n    }\n}", []LnsAny{self.FP.getConv2formName(_env, &node.Nodes_Node), self.env.FP.getLuavm(_env, node.FP.Get_exp(_env).FP.IsThreading(_env))}))
        return 
    }
    self.FP.Writeln(_env, _env.LuaVM.String_format("// for %d: %s", []LnsAny{node.FP.Get_pos(_env).LineNo, Nodes_getNodeKindName(_env, node.FP.Get_kind(_env))}))
    self.FP.Write(_env, _env.LuaVM.String_format("func %s( src func (%s", []LnsAny{self.FP.getConv2formName(_env, &node.Nodes_Node), self.FP.getEnvArgDecl(_env, funcType.FP.Get_argTypeInfoList(_env).Len())}))
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList(_env).Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        self.FP.Write(_env, _env.LuaVM.String_format("arg%d %s", []LnsAny{index, self.FP.type2gotype(_env, argType)}))
    }
    self.FP.Write(_env, ")")
    self.FP.outputRetType(_env, funcType.FP.Get_retTypeInfoList(_env))
    self.FP.Writeln(_env, ") LnsForm {")
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, _env.LuaVM.String_format("return func (%s argList []LnsAny) []LnsAny {", []LnsAny{self.FP.getEnvArgDecl(_env, 1)}))
    self.FP.PushIndent(_env, nil)
    if funcType.FP.Get_retTypeInfoList(_env).Len() > 0{
        self.FP.Write(_env, "return ")
        if funcType.FP.Get_argTypeInfoList(_env).Len() > 0{
            self.FP.Write(_env, "Lns_2DDD(")
        }
    }
    self.FP.Write(_env, "src(")
    self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, funcType.FP.Get_argTypeInfoList(_env).Len(), self.option.FP.Get_addEnvArg(_env)))
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList(_env).Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        if argType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            self.FP.Write(_env, _env.LuaVM.String_format("argList[ %d: ]", []LnsAny{index - 1}))
            break
        }
        self.FP.Write(_env, _env.LuaVM.String_format("Lns_getFromMulti( argList, %d )", []LnsAny{index - 1}))
    }
    self.FP.Write(_env, ")")
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

// 1659: decl @lune.@base.@convGo.convFilter.processConvStat
func (self *convGo_convFilter) ProcessConvStat(_env *LnsEnv, node *Nodes_ConvStatNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_txt(_env))
}

// 1665: decl @lune.@base.@convGo.convFilter.outputTopScopeVar
func (self *convGo_convFilter) outputTopScopeVar(_env *LnsEnv, node *Nodes_DeclVarNode) {
    for _, _symbolInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( symbolInfo.FP.Get_scope(_env) == self.moduleScope) &&
            _env.SetStackVal( node.FP.Get_mode(_env) == Nodes_DeclVarMode__Let) ).(bool)){
            self.FP.Writeln(_env, _env.LuaVM.String_format("var %s %s", []LnsAny{self.FP.getSymbolSym(_env, symbolInfo), self.FP.type2gotype(_env, symbolInfo.FP.Get_typeInfo(_env))}))
        }
    }
}

// 1674: decl @lune.@base.@convGo.convFilter.outputConvExt
func (self *convGo_convFilter) outputConvExt(_env *LnsEnv, funcNode *Nodes_Node) {
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
    self.FP.Write(_env, _env.LuaVM.String_format("func Lns_callExt%d( args []LnsAny ) (", []LnsAny{funcNode.FP.Get_id(_env)}))
    for _index, _retType := range( funcNode.FP.Get_expType(_env).FP.Get_retTypeInfoList(_env).Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(_env, ",")
        }
        self.FP.Write(_env, self.FP.type2gotype(_env, retType))
    }
    self.FP.Writeln(_env, ") {")
    self.FP.Write(_env, "    return ")
    for _index, _ := range( funcNode.FP.Get_expType(_env).FP.Get_retTypeInfoList(_env).Items ) {
        index := _index + 1
        if index > 1{
            self.FP.Write(_env, ",")
        }
        self.FP.Write(_env, _env.LuaVM.String_format("Lns_getFromMulti( args, %d )", []LnsAny{index - 1}))
    }
    self.FP.Writeln(_env, "")
    self.FP.Writeln(_env, "}")
}

// 1701: decl @lune.@base.@convGo.convFilter.outputModule
func (self *convGo_convFilter) outputModule(_env *LnsEnv, moduleTypeInfo *Ast_TypeInfo,addDot bool) {
    self.FP.Write(_env, self.FP.getModuleSym(_env, moduleTypeInfo, addDot))
}

// 1716: decl @lune.@base.@convGo.convFilter.outputModuleImport
func (self *convGo_convFilter) outputModuleImport(_env *LnsEnv, node *Nodes_DeclClassNode) {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__ExtModule) ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( node.FP.Get_lang(_env) != Types_Lang__Go) &&
            _env.SetStackVal( node.FP.Get_lang(_env) != Types_Lang__Same) ).(bool))) ||
        _env.SetStackVal( self.FP.isSamePackageExtModule(_env, node.FP.Get_expType(_env))) ).(bool){
        return 
    }
    var normalType *Ast_NormalTypeInfo
    normalType = Lns_unwrap( Ast_NormalTypeInfoDownCastF(node.FP.Get_expType(_env).FP)).(*Ast_NormalTypeInfo)
    self.FP.Write(_env, _env.LuaVM.String_format("import %s ", []LnsAny{node.FP.Get_expType(_env).FP.Get_rawTxt(_env)}))
    var mod string
    mod = convGo_convExp8460(Lns_2DDD(_env.LuaVM.String_gsub(normalType.FP.Get_requirePath(_env),"%.[^%.]+$", "")))
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(mod,"^go/", nil, nil))){
        var workMod string
        workMod = convGo_convExp8499(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(mod,"^go/", "")).(string),"%.", "/")).(string),":", ".")))
        self.FP.Writeln(_env, _env.LuaVM.String_format("\"%s\"", []LnsAny{workMod}))
    } else { 
        var path string
        path = convGo_convExp8537(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(normalType.FP.Get_requirePath(_env),"%.", "/")).(string),":", ".")))
        self.FP.Writeln(_env, _env.LuaVM.String_format("\"%s/%s\"", []LnsAny{self.option.FP.Get_appName(_env), path}))
    }
}

// 1738: decl @lune.@base.@convGo.convFilter.outputImport
func (self *convGo_convFilter) outputImport(_env *LnsEnv, node *Nodes_ImportNode) {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.isSameModDir(_env, node.FP.Get_moduleTypeInfo(_env))) ||
        _env.SetStackVal( Ast_isBuiltin(_env, node.FP.Get_moduleTypeInfo(_env).FP.Get_typeId(_env).Id)) ).(bool){
        return 
    }
    self.FP.Write(_env, "import ")
    var packSym string
    packSym = node.FP.Get_assignName(_env)
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(node.FP.Get_modulePath(_env),"^go/", nil, nil))){
        var workMod string
        workMod = convGo_convExp8641(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(node.FP.Get_modulePath(_env),"^go/", "")).(string),"%.", "/")).(string),":", ".")))
        self.FP.Writeln(_env, _env.LuaVM.String_format("%s \"%s\"", []LnsAny{packSym, Lns_car(_env.LuaVM.String_gsub(workMod,"/[^/]+$", "")).(string)}))
    } else { 
        var modulePath string
        var count LnsInt
        modulePath,count = _env.LuaVM.String_gsub(node.FP.Get_modulePath(_env),"([^%.]+)%.[^%.]+$", "/%1")
        if count == 0{
            self.FP.Writeln(_env, _env.LuaVM.String_format("%s \"%s\"", []LnsAny{packSym, self.option.appName}))
        } else { 
            var modDir string
            modDir = convGo_convExp8720(Lns_2DDD(_env.LuaVM.String_gsub(modulePath,"/", "")))
            self.FP.Writeln(_env, _env.LuaVM.String_format("%s \"%s/%s\"", []LnsAny{packSym, self.option.appName, Lns_car(_env.LuaVM.String_gsub(modDir,"%.", "/")).(string)}))
        }
    }
    self.module2PackSym.Set(node.FP.Get_moduleTypeInfo(_env),packSym)
}

// 1765: decl @lune.@base.@convGo.convFilter.processRoot
func (self *convGo_convFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    for _, _importNode := range( node.FP.Get_nodeManager(_env).FP.GetImportNodeList(_env).Items ) {
        importNode := _importNode.(Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        self.moduleType2SymbolMap.Set(importNode.FP.Get_moduleTypeInfo(_env),importNode.FP.Get_symbolInfo(_env))
    }
    for _pragma := range( node.FP.Get_luneHelperInfo(_env).PragmaSet.Items ) {
        pragma := _pragma
        switch _exp8848 := pragma.(type) {
        case *LuneControl_Pragma__limit_conv_code:
        codeSet := _exp8848.Val1
            if Lns_op_not(codeSet.Has(LuneControl_Code__Go)){
                self.FP.Writeln(_env, "// This code is transcompiled by LuneScript.")
                self.FP.Writeln(_env, _env.LuaVM.String_format("package %s", []LnsAny{self.option.packageName}))
                return 
            }
        }
    }
    var isUsingInTest func(_env *LnsEnv, aNode *Nodes_Node) bool
    isUsingInTest = func(_env *LnsEnv, aNode *Nodes_Node) bool {
        for _, _testBlock := range( node.FP.Get_nodeManager(_env).FP.GetTestBlockNodeList(_env).Items ) {
            testBlock := _testBlock.(Nodes_TestBlockNodeDownCast).ToNodes_TestBlockNode()
            if testBlock.FP.IsInnerPos(_env, aNode.FP.Get_pos(_env)){
                return true
            }
        }
        return false
    }
    var builtinFuncs *Builtin_BuiltinFuncType
    builtinFuncs = TransUnit_getBuiltinFunc(_env)
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
    
    self.FP.Writeln(_env, "// This code is transcompiled by LuneScript.")
    self.FP.Writeln(_env, _env.LuaVM.String_format("package %s", []LnsAny{self.option.packageName}))
    self.FP.Writeln(_env, "import . \"github.com/ifritJP/LuneScript/src/lune/base/runtime_go\"")
    for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetImportNodeList(_env).Items ) {
        workNode := _workNode.(Nodes_ImportNodeDownCast).ToNodes_ImportNode()
        self.FP.outputImport(_env, workNode)
    }
    for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclClassNodeList(_env).Items ) {
        workNode := _workNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
        self.FP.outputModuleImport(_env, workNode)
    }
    var initModVar string
    initModVar = _env.LuaVM.String_format("init_%s", []LnsAny{self.FP.getModuleName(_env, node.FP.Get_moduleTypeInfo(_env), false)})
    self.FP.Writeln(_env, _env.LuaVM.String_format("var %s bool", []LnsAny{initModVar}))
    self.FP.pushProcessMode(_env, convGo_ProcessMode__DeclTopScopeVar)
    var modSym *Ast_SymbolInfo
    modSym = Lns_unwrap( self.moduleScope.FP.GetSymbolInfoChild(_env, "__mod__")).(*Ast_SymbolInfo)
    self.FP.Writeln(_env, _env.LuaVM.String_format("var %s string", []LnsAny{self.FP.getSymbolSym(_env, modSym)}))
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_Node)
        procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convGo_filter_1292_(_env, workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclEnumNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclEnumNodeDownCast).ToNodes_DeclEnumNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclVarNodeList(_env).Items ) {
        workNode := _workNode.(Nodes_DeclVarNodeDownCast).ToNodes_DeclVarNode()
        self.FP.outputTopScopeVar(_env, workNode)
    }
    self.FP.popProcessMode(_env)
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_Node)
        procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convGo_filter_1292_(_env, workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclAlgeNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclAlgeNodeDownCast).ToNodes_DeclAlgeNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_Node)
        procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convGo_filter_1292_(_env, workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclFormNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclFormNodeDownCast).ToNodes_DeclFormNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_ExpCallNode)
        procNode = func(_env *LnsEnv, workNode *Nodes_ExpCallNode) {
            self.FP.processGenericsCall(_env, workNode)
            self.FP.outputNilAccCall(_env, workNode)
            self.FP.outputConvExt(_env, workNode.FP.Get_func(_env))
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpCallNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCallNodeDownCast).ToNodes_ExpCallNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_ExpCastNode)
        procNode = func(_env *LnsEnv, workNode *Nodes_ExpCastNode) {
            self.FP.outputConvToForm(_env, workNode)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpCastNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCastNodeDownCast).ToNodes_ExpCastNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_Node)
        procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convGo_filter_1292_(_env, workNode, self, &node.Nodes_Node)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetTestCaseNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_TestCaseNodeDownCast).ToNodes_TestCaseNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_IfUnwrapNode)
        procNode = func(_env *LnsEnv, workNode *Nodes_IfUnwrapNode) {
            var symTypeList *LnsList
            symTypeList = NewLnsList([]LnsAny{})
            {
                var _from10331 LnsInt = 1
                var _to10331 LnsInt = workNode.FP.Get_varSymList(_env).Len()
                for _work10331 := _from10331; _work10331 <= _to10331; _work10331++ {
                    symTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeStem_))
                }
            }
            self.FP.processConvExp(_env, workNode.FP.Get_id(_env), symTypeList, workNode.FP.Get_expList(_env), false)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetIfUnwrapNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_IfUnwrapNodeDownCast).ToNodes_IfUnwrapNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_ExpSetValNode)
        procNode = func(_env *LnsEnv, workNode *Nodes_ExpSetValNode) {
            self.FP.processConvExp(_env, workNode.FP.Get_id(_env), workNode.FP.Get_exp1(_env).FP.Get_expTypeList(_env), workNode.FP.Get_exp2(_env), false)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpSetValNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpSetValNodeDownCast).ToNodes_ExpSetValNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_ExpCallNode)
        procNode = func(_env *LnsEnv, workNode *Nodes_ExpCallNode) {
            var needEnv bool
            needEnv = Lns_op_not(Ast_isBuiltin(_env, workNode.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_typeId(_env).Id))
            self.FP.processConvExp(_env, workNode.FP.Get_id(_env), workNode.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_argTypeInfoList(_env), workNode.FP.Get_argList(_env), needEnv)
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetExpCallNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_ExpCallNodeDownCast).ToNodes_ExpCallNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_DeclVarNode)
        procNode = func(_env *LnsEnv, workNode *Nodes_DeclVarNode) {
            var typeList *LnsList
            typeList = NewLnsList([]LnsAny{})
            {
                _expList := workNode.FP.Get_expList(_env)
                if !Lns_IsNil( _expList ) {
                    expList := _expList.(*Nodes_ExpListNode)
                    for _index, _symbolInfo := range( workNode.FP.Get_symbolInfoList(_env).Items ) {
                        index := _index + 1
                        symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                        if _env.PopVal( _env.IncStack() ||
                            _env.SetStackVal( workNode.FP.Get_mode(_env) == Nodes_DeclVarMode__Let) ||
                            _env.SetStackVal( workNode.FP.Get_mode(_env) == Nodes_DeclVarMode__Unwrap) ).(bool){
                            if workNode.FP.Get_unwrapFlag(_env){
                                typeList.Insert(Ast_TypeInfo2Stem(expList.FP.GetExpTypeNoDDDAt(_env, index)))
                            } else { 
                                typeList.Insert(Ast_TypeInfo2Stem(symbolInfo.FP.Get_typeInfo(_env)))
                            }
                        }
                    }
                    self.FP.processConvExp(_env, workNode.FP.Get_id(_env), typeList, expList, false)
                }
            }
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclVarNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclVarNodeDownCast).ToNodes_DeclVarNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, tmpNode)
            }
        }
    }
    
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_AliasNode)
        procNode = func(_env *LnsEnv, workNode *Nodes_AliasNode) {
            self.FP.Writeln(_env, _env.LuaVM.String_format("type %s = %s", []LnsAny{self.FP.getSymbolSym(_env, workNode.FP.Get_newSymbol(_env)), self.FP.getTypeSymbol(_env, workNode.FP.Get_typeInfo(_env).FP.Get_aliasSrc(_env))}))
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetAliasNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_AliasNodeDownCast).ToNodes_AliasNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, tmpNode)
            }
        }
    }
    
    self.FP.pushProcessMode(_env, convGo_ProcessMode__NonClosureFuncDecl)
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_Node)
        procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convGo_filter_1292_(_env, workNode, self, &node.Nodes_Node)
            self.FP.Writeln(_env, "")
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclFuncNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    self.FP.popProcessMode(_env)
    self.FP.pushProcessMode(_env, convGo_ProcessMode__DeclClass)
    {
        var procNode func(_env *LnsEnv, workNode *Nodes_Node)
        procNode = func(_env *LnsEnv, workNode *Nodes_Node) {
            convGo_filter_1292_(_env, workNode, self, &node.Nodes_Node)
            self.FP.Writeln(_env, "")
        }
        for _, _tmpNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclClassNodeList(_env).Items ) {
            tmpNode := _tmpNode.(Nodes_DeclClassNodeDownCast).ToNodes_DeclClassNode()
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.enableTest) ||
                _env.SetStackVal( Lns_op_not(isUsingInTest(_env, &tmpNode.Nodes_Node))) ).(bool){
                procNode(_env, &tmpNode.Nodes_Node)
            }
        }
    }
    
    self.FP.popProcessMode(_env)
    self.FP.Writeln(_env, _env.LuaVM.String_format("func Lns_%s_init(%s) {", []LnsAny{self.FP.getModuleName(_env, node.FP.Get_moduleTypeInfo(_env), false), self.FP.getEnvArgDecl(_env, 0)}))
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, _env.LuaVM.String_format("if %s { return }", []LnsAny{initModVar}))
    self.FP.Writeln(_env, _env.LuaVM.String_format("%s = true", []LnsAny{initModVar}))
    self.FP.Writeln(_env, _env.LuaVM.String_format("%s = \"%s\"", []LnsAny{self.FP.getSymbolSym(_env, modSym), node.FP.Get_moduleTypeInfo(_env).FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), nil)}))
    self.FP.Writeln(_env, "Lns_InitMod()")
    var modulePath string
    modulePath = convGo_convExp11364(Lns_2DDD(_env.LuaVM.String_gsub(node.FP.Get_moduleTypeInfo(_env).FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), nil),"@", "")))
    var moduleName string
    moduleName = self.FP.getModuleName(_env, node.FP.Get_moduleTypeInfo(_env), false)
    if self.enableTest{
        for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetTestCaseNodeList(_env).Items ) {
            workNode := _workNode.(Nodes_TestCaseNodeDownCast).ToNodes_TestCaseNode()
            self.FP.Writeln(_env, _env.LuaVM.String_format("Testing_registerTestcase( \"%s\", \"%s\", lns_test_%s_%s )", []LnsAny{modulePath, workNode.FP.Get_name(_env).Txt, moduleName, workNode.FP.Get_name(_env).Txt}))
        }
    }
    for _, _child := range( node.FP.Get_children(_env).Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(convGo_ignoreNodeInInnerBlockSet.Has(child.FP.Get_kind(_env))){
            convGo_filter_1292_(_env, child, self, &node.Nodes_Node)
        }
    }
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
    if self.option.mainModule == Lns_car(_env.LuaVM.String_gsub(self.FP.getCanonicalName(_env, self.moduleTypeInfo, false),"@", "")).(string){
        var hasMain bool
        hasMain = false
        for _, _workNode := range( node.FP.Get_nodeManager(_env).FP.GetDeclFuncNodeList(_env).Items ) {
            workNode := _workNode.(Nodes_DeclFuncNodeDownCast).ToNodes_DeclFuncNode()
            if convGo_isMain_1080_(_env, workNode.FP.Get_expType(_env)){
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
            self.FP.Writeln(_env, _env.LuaVM.String_format("func %s___main( %sargs *LnsList ) LnsInt {", []LnsAny{convGo_concatGLSym_1325_(_env, moduleSym, true), self.FP.getEnvArgDecl(_env, 1)}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("Lns_%s_init( %s )", []LnsAny{moduleSym, callArgs}))
            self.FP.Writeln(_env, "return 0")
            self.FP.Writeln(_env, "}")
        }
    }
    self.FP.Writeln(_env, "func init() {")
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, _env.LuaVM.String_format("%s = false", []LnsAny{initModVar}))
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 2066: decl @lune.@base.@convGo.convFilter.processSubfile
func (self *convGo_convFilter) ProcessSubfile(_env *LnsEnv, node *Nodes_SubfileNode,_opt LnsAny) {
}

// 2072: decl @lune.@base.@convGo.convFilter.processEnv
func (self *convGo_convFilter) ProcessEnv(_env *LnsEnv, node *Nodes_EnvNode,_opt LnsAny) {
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, "Lns_LockEnvSync()")
    self.FP.PopIndent(_env)
    self.env.FP.push(_env, &node.Nodes_Node)
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.env.FP.pop(_env)
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, "Lns_UnlockEnvSync()")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 2093: decl @lune.@base.@convGo.convFilter.processBlockSub
func (self *convGo_convFilter) ProcessBlockSub(_env *LnsEnv, node *Nodes_BlockNode,_opt LnsAny) {
    if _switch11804 := node.FP.Get_blockKind(_env); _switch11804 == Nodes_BlockKind__Block {
        self.FP.Writeln(_env, "{")
    }
    self.FP.pushProcessMode(_env, convGo_ProcessMode__Main)
    self.FP.PushIndent(_env, nil)
    for _, _child := range( node.FP.Get_stmtList(_env).Items ) {
        child := _child.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(convGo_ignoreNodeInInnerBlockSet.Has(child.FP.Get_kind(_env))){
            convGo_filter_1292_(_env, child, self, &node.Nodes_Node)
        }
    }
    self.FP.PopIndent(_env)
    self.FP.popProcessMode(_env)
    if _switch11879 := node.FP.Get_blockKind(_env); _switch11879 == Nodes_BlockKind__Block {
        self.FP.Writeln(_env, "}")
    }
}

// 2119: decl @lune.@base.@convGo.convFilter.expList2Slice
func (self *convGo_convFilter) expList2Slice(_env *LnsEnv, subList *Nodes_ExpListNode,toStem bool) {
    var processExp func(_env *LnsEnv, exp *Nodes_Node)
    processExp = func(_env *LnsEnv, exp *Nodes_Node) {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( toStem) &&
            _env.SetStackVal( Ast_isClass(_env, exp.FP.Get_expType(_env))) ).(bool)){
            self.FP.Write(_env, _env.LuaVM.String_format("%s2Stem(", []LnsAny{self.FP.getTypeSymbolWithPrefix(_env, exp.FP.Get_expType(_env).FP.Get_nonnilableType(_env))}))
            convGo_filter_1292_(_env, Nodes_getCastUnwraped(_env, exp), self, &subList.Nodes_Node)
            self.FP.Write(_env, ")")
        } else { 
            convGo_filter_1292_(_env, exp, self, &subList.Nodes_Node)
        }
    }
    var mRetIndex LnsAny
    mRetIndex = _env.NilAccFin(_env.NilAccPush(subList.FP.Get_mRetExp(_env)) && 
    Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( mRetIndex) &&
        _env.SetStackVal( mRetIndex == 1) )){
        var subExp *Nodes_Node
        subExp = subList.FP.Get_expList(_env).GetAt(1).(Nodes_NodeDownCast).ToNodes_Node()
        if subExp.FP.Get_expType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__DDD{
            self.FP.Write(_env, "Lns_2DDD(")
            processExp(_env, subExp)
            self.FP.Write(_env, ")")
        } else { 
            processExp(_env, subExp)
        }
    } else { 
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( mRetIndex) &&
            _env.SetStackVal( mRetIndex != 1) )){
            self.FP.Write(_env, "append( ")
        }
        self.FP.Write(_env, "[]LnsAny{")
        for _subIndex, _subExp := range( subList.FP.Get_expList(_env).Items ) {
            subIndex := _subIndex + 1
            subExp := _subExp.(Nodes_NodeDownCast).ToNodes_Node()
            if mRetIndex == subIndex{
                if mRetIndex != 1{
                    self.FP.Write(_env, "}, ")
                }
                if subExp.FP.Get_expType(_env).FP.Get_kind(_env) != Ast_TypeInfoKind__DDD{
                    self.FP.Write(_env, "Lns_2DDD(")
                    processExp(_env, subExp)
                    self.FP.Write(_env, ")")
                } else { 
                    processExp(_env, subExp)
                }
                self.FP.Write(_env, "...")
                break
            }
            if subIndex != 1{
                self.FP.Write(_env, ", ")
            }
            processExp(_env, subExp)
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( mRetIndex) &&
            _env.SetStackVal( mRetIndex != 1) )){
            self.FP.Write(_env, ")")
        } else { 
            self.FP.Write(_env, "}")
        }
    }
}

// 2182: decl @lune.@base.@convGo.convFilter.processSetFromExpList
func (self *convGo_convFilter) processSetFromExpList(_env *LnsEnv, convArgFuncName string,dstTypeList *LnsList,expListNode *Nodes_ExpListNode,addEnvArg bool) {
    if _switch12744 := convGo_getExpListKind_1885_(_env, dstTypeList, expListNode, addEnvArg); _switch12744 == convGo_ExpListKind__Conv {
        self.FP.Write(_env, _env.LuaVM.String_format("%s(", []LnsAny{convArgFuncName}))
        var mRetIndex LnsAny
        mRetIndex = _env.NilAccFin(_env.NilAccPush(expListNode.FP.Get_mRetExp(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
        self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, expListNode.FP.Get_expList(_env).Len(), addEnvArg))
        for _index, _exp := range( expListNode.FP.Get_expList(_env).Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            if exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
                break
            }
            if index != 1{
                self.FP.Write(_env, ", ")
            }
            if mRetIndex == index{
                self.FP.Write(_env, "Lns_2DDD(")
                convGo_filter_1292_(_env, Nodes_getCastUnwraped(_env, exp), self, &expListNode.Nodes_Node)
                self.FP.Write(_env, ")")
                break
            }
            convGo_filter_1292_(_env, exp, self, &expListNode.Nodes_Node)
            if _env.NilAccFin(_env.NilAccPush(expListNode.FP.Get_mRetExp(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)})) == index{
                break
            }
        }
        self.FP.Write(_env, ")")
    } else if _switch12744 == convGo_ExpListKind__Slice {
        self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, dstTypeList.Len(), addEnvArg))
        for _index, _argType := range( dstTypeList.Items ) {
            index := _index + 1
            argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index != 1{
                self.FP.Write(_env, ", ")
            }
            if expListNode.FP.Get_expList(_env).Len() >= index{
                var argExp *Nodes_Node
                argExp = expListNode.FP.Get_expList(_env).GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                {
                    _exp2ddd := Nodes_ExpToDDDNodeDownCastF(argExp.FP)
                    if !Lns_IsNil( _exp2ddd ) {
                        exp2ddd := _exp2ddd.(*Nodes_ExpToDDDNode)
                        self.FP.expList2Slice(_env, exp2ddd.FP.Get_expList(_env), false)
                    } else {
                        if argExp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr{
                            if argType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                                self.FP.Write(_env, "[]LnsAny{}")
                            } else { 
                                self.FP.Write(_env, "nil")
                            }
                        } else { 
                            convGo_filter_1292_(_env, argExp, self, &expListNode.Nodes_Node)
                        }
                    }
                }
            } else { 
                self.FP.Write(_env, "[]LnsAny{}")
            }
        }
    } else if _switch12744 == convGo_ExpListKind__Direct {
        self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, dstTypeList.Len(), addEnvArg))
        var mRetIndex LnsAny
        mRetIndex = _env.NilAccFin(_env.NilAccPush(expListNode.FP.Get_mRetExp(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Nodes_MRetExp).FP.Get_index(_env)}))
        for _index, _dstType := range( dstTypeList.Items ) {
            index := _index + 1
            dstType := _dstType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if mRetIndex == index - 1{
                break
            }
            if index != 1{
                self.FP.Write(_env, ", ")
            }
            var exp *Nodes_Node
            if expListNode.FP.Get_expList(_env).Len() < index{
                exp = &self.noneNode.Nodes_Node
                
            } else { 
                exp = expListNode.FP.Get_expList(_env).GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( index == dstTypeList.Len()) &&
                _env.SetStackVal( dstType.FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( expListNode.FP.Get_expList(_env).Len() < index) ||
                    _env.SetStackVal( exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr) ).(bool){
                    self.FP.Write(_env, "[]LnsAny{}")
                } else { 
                    convGo_filter_1292_(_env, exp, self, &expListNode.Nodes_Node)
                }
            } else { 
                if _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( expListNode.FP.Get_expList(_env).Len() < index) ||
                    _env.SetStackVal( exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Abbr) ).(bool){
                    self.FP.Write(_env, "nil")
                } else if expListNode.FP.Get_expTypeList(_env).GetAt(index).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    self.FP.Write(_env, "Lns_car(")
                    convGo_filter_1292_(_env, exp, self, &expListNode.Nodes_Node)
                    self.FP.Write(_env, ")")
                } else { 
                    convGo_filter_1292_(_env, exp, self, &expListNode.Nodes_Node)
                }
            }
        }
    }
}

// 2289: decl @lune.@base.@convGo.convFilter.processStmtExp
func (self *convGo_convFilter) ProcessStmtExp(_env *LnsEnv, node *Nodes_StmtExpNode,_opt LnsAny) {
    convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Writeln(_env, "")
}

// 2296: decl @lune.@base.@convGo.convFilter.processDeclAlge
func (self *convGo_convFilter) ProcessDeclAlge(_env *LnsEnv, node *Nodes_DeclAlgeNode,_opt LnsAny) {
    self.FP.Writeln(_env, _env.LuaVM.String_format("// decl alge -- %s", []LnsAny{node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
    self.FP.Writeln(_env, _env.LuaVM.String_format("type %s = LnsAny", []LnsAny{self.FP.getTypeSymbol(_env, &node.FP.Get_algeType(_env).Ast_TypeInfo)}))
    {
        __collection12959 := node.FP.Get_algeType(_env).FP.Get_valInfoMap(_env)
        __sorted12959 := __collection12959.CreateKeyListStr()
        __sorted12959.Sort( LnsItemKindStr, nil )
        for _, ___key12959 := range( __sorted12959.Items ) {
            valInfo := __collection12959.Items[ ___key12959 ].(Ast_AlgeValInfoDownCast).ToAst_AlgeValInfo()
            var algeSym string
            algeSym = self.FP.getAlgeSymbol(_env, valInfo)
            self.FP.Writeln(_env, _env.LuaVM.String_format("type %s struct{", []LnsAny{algeSym}))
            for _index, _paramType := range( valInfo.FP.Get_typeList(_env).Items ) {
                index := _index + 1
                paramType := _paramType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                self.FP.Writeln(_env, _env.LuaVM.String_format("Val%d %s", []LnsAny{index, self.FP.type2gotype(_env, paramType)}))
            }
            self.FP.Writeln(_env, "}")
            if valInfo.FP.Get_typeList(_env).Len() == 0{
                self.FP.Writeln(_env, _env.LuaVM.String_format("var %s_Obj = &%s{}", []LnsAny{algeSym, algeSym}))
            }
            self.FP.Writeln(_env, _env.LuaVM.String_format("func (self *%s) GetTxt() string {", []LnsAny{algeSym}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("return \"%s.%s\"", []LnsAny{node.FP.Get_algeType(_env).FP.Get_rawTxt(_env), valInfo.FP.Get_name(_env)}))
            self.FP.Writeln(_env, "}")
        }
    }
}

// 2318: decl @lune.@base.@convGo.convFilter.processNewAlgeVal
func (self *convGo_convFilter) ProcessNewAlgeVal(_env *LnsEnv, node *Nodes_NewAlgeValNode,_opt LnsAny) {
    var algeSym string
    algeSym = self.FP.getAlgeSymbol(_env, node.FP.Get_valInfo(_env))
    if node.FP.Get_valInfo(_env).FP.Get_typeList(_env).Len() == 0{
        self.FP.Write(_env, _env.LuaVM.String_format("%s_Obj", []LnsAny{algeSym}))
    } else { 
        self.FP.Write(_env, _env.LuaVM.String_format("&%s{", []LnsAny{algeSym}))
        for _index, _param := range( node.FP.Get_paramList(_env).Items ) {
            index := _index + 1
            param := _param.(Nodes_NodeDownCast).ToNodes_Node()
            if index > 1{
                self.FP.Write(_env, ", ")
            }
            convGo_filter_1292_(_env, param, self, &node.Nodes_Node)
        }
        self.FP.Write(_env, "}")
    }
}

// 2337: decl @lune.@base.@convGo.convFilter.processDeclMember
func (self *convGo_convFilter) ProcessDeclMember(_env *LnsEnv, node *Nodes_DeclMemberNode,_opt LnsAny) {
    self.FP.outputSymbol(_env, &convGo_SymbolKind__Member{Ast_isPubToExternal(_env, node.FP.Get_accessMode(_env))}, node.FP.Get_name(_env).Txt)
    self.FP.Write(_env, _env.LuaVM.String_format(" %s", []LnsAny{self.FP.type2gotype(_env, node.FP.Get_refType(_env).FP.Get_expType(_env))}))
    self.FP.Writeln(_env, "")
}

// 2346: decl @lune.@base.@convGo.convFilter.processExpMacroExp
func (self *convGo_convFilter) ProcessExpMacroExp(_env *LnsEnv, node *Nodes_ExpMacroExpNode,_opt LnsAny) {
    for _, _stmt := range( node.FP.Get_stmtList(_env).Items ) {
        stmt := _stmt.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(convGo_ignoreNodeInInnerBlockSet.Has(stmt.FP.Get_kind(_env))){
            convGo_filter_1292_(_env, stmt, self, &node.Nodes_Node)
        }
    }
}

// 2356: decl @lune.@base.@convGo.convFilter.processDeclMacro
func (self *convGo_convFilter) ProcessDeclMacro(_env *LnsEnv, node *Nodes_DeclMacroNode,_opt LnsAny) {
}

// 2362: decl @lune.@base.@convGo.convFilter.processExpMacroStat
func (self *convGo_convFilter) ProcessExpMacroStat(_env *LnsEnv, node *Nodes_ExpMacroStatNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpMacroStat"
    Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
}

// 2368: decl @lune.@base.@convGo.convFilter.outputDeclFuncArg
func (self *convGo_convFilter) outputDeclFuncArg(_env *LnsEnv, funcType *Ast_TypeInfo) {
    self.FP.Write(_env, self.FP.getEnvArgDecl(_env, funcType.FP.Get_argTypeInfoList(_env).Len()))
    for _index, _argType := range( funcType.FP.Get_argTypeInfoList(_env).Items ) {
        index := _index + 1
        argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index != 1{
            self.FP.Write(_env, ", ")
        }
        self.FP.Write(_env, _env.LuaVM.String_format("arg%d ", []LnsAny{index}))
        self.FP.Write(_env, self.FP.type2gotype(_env, argType))
    }
}

// 2380: decl @lune.@base.@convGo.convFilter.processDeclConstr
func (self *convGo_convFilter) ProcessDeclConstr(_env *LnsEnv, node *Nodes_DeclConstrNode,_opt LnsAny) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType(_env).FP.Get_parentInfo(_env)
    var className string
    className = self.FP.getTypeSymbol(_env, classType)
    self.FP.Writeln(_env, _env.LuaVM.String_format("// %d: %s", []LnsAny{node.FP.Get_pos(_env).LineNo, Nodes_getNodeKindName(_env, node.FP.Get_kind(_env))}))
    self.FP.Write(_env, _env.LuaVM.String_format("func (self *%s) %s(", []LnsAny{className, self.FP.getConstrSymbol(_env, classType)}))
    self.FP.Write(_env, self.FP.getEnvArgDecl(_env, node.FP.Get_declInfo(_env).FP.Get_argList(_env).Len()))
    for _index, _arg := range( node.FP.Get_declInfo(_env).FP.Get_argList(_env).Items ) {
        index := _index + 1
        arg := _arg.(Nodes_NodeDownCast).ToNodes_Node()
        if index != 1{
            self.FP.Write(_env, ",")
        }
        convGo_filter_1292_(_env, arg, self, &node.Nodes_Node)
    }
    self.FP.Writeln(_env, ") {")
    convGo_filter_1292_(_env, &Lns_unwrap( node.FP.Get_declInfo(_env).FP.Get_body(_env)).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "}")
}

// 2403: decl @lune.@base.@convGo.convFilter.processDeclDestr
func (self *convGo_convFilter) ProcessDeclDestr(_env *LnsEnv, node *Nodes_DeclDestrNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processDeclDestr"
    Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
}

// 2409: decl @lune.@base.@convGo.convFilter.processExpCallSuperCtor
func (self *convGo_convFilter) ProcessExpCallSuperCtor(_env *LnsEnv, node *Nodes_ExpCallSuperCtorNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("self.%s(", []LnsAny{self.FP.getConstrSymbol(_env, node.FP.Get_superType(_env))}))
    {
        _argList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), argList), node.FP.Get_methodType(_env).FP.Get_argTypeInfoList(_env), argList, self.option.FP.Get_addEnvArg(_env))
        } else {
            self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, 0, self.option.FP.Get_addEnvArg(_env)))
        }
    }
    self.FP.Writeln(_env, ")")
}

// 2423: decl @lune.@base.@convGo.convFilter.processExpCallSuper
func (self *convGo_convFilter) ProcessExpCallSuper(_env *LnsEnv, node *Nodes_ExpCallSuperNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("self.%s.%s(", []LnsAny{self.FP.getTypeSymbol(_env, node.FP.Get_methodType(_env).FP.Get_parentInfo(_env)), self.FP.getFuncSymbol(_env, node.FP.Get_methodType(_env))}))
    {
        _argList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), argList), node.FP.Get_methodType(_env).FP.Get_argTypeInfoList(_env), argList, self.option.FP.Get_addEnvArg(_env))
        } else {
            self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, 0, self.option.FP.Get_addEnvArg(_env)))
        }
    }
    self.FP.Write(_env, ")")
}

// 2438: decl @lune.@base.@convGo.convFilter.outputDummyReturn
func (self *convGo_convFilter) outputDummyReturn(_env *LnsEnv, retTypeInfoList *LnsList) {
    self.FP.Writeln(_env, "// insert a dummy")
    self.FP.Write(_env, "    return ")
    for _index, _retType := range( retTypeInfoList.Items ) {
        index := _index + 1
        retType := _retType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if index > 1{
            self.FP.Write(_env, ",")
        }
        if convGo_isAnyType_1296_(_env, retType){
            self.FP.Write(_env, "nil")
        } else { 
            if _switch13772 := convGo_getOrgTypeInfo_1660_(_env, retType); _switch13772 == Ast_builtinTypeInt {
                self.FP.Write(_env, "0")
            } else if _switch13772 == Ast_builtinTypeReal {
                self.FP.Write(_env, "0.0")
            } else if _switch13772 == Ast_builtinTypeBool {
                self.FP.Write(_env, "false")
            } else if _switch13772 == Ast_builtinTypeString {
                self.FP.Write(_env, "\"\"")
            } else {
                self.FP.Write(_env, "nil")
            }
        }
    }
    self.FP.Writeln(_env, "")
}

// 2470: decl @lune.@base.@convGo.convFilter.outputDeclFuncInfo
func (self *convGo_convFilter) OutputDeclFuncInfo(_env *LnsEnv, node *Nodes_Node,declInfo *Nodes_DeclFuncInfo) {
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_expType(_env)
    if funcType.FP.Get_abstractFlag(_env){
        return 
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( declInfo.FP.Get_name(_env)) &&
        _env.SetStackVal( Lns_op_not(convGo_isClosure_1307_(_env, funcType))) )){
        self.FP.Writeln(_env, _env.LuaVM.String_format("// %d: decl %s", []LnsAny{node.FP.Get_pos(_env).LineNo, self.FP.getCanonicalName(_env, funcType, false)}))
    }
    var convFunc *convGo_FuncConv
    convFunc = self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__DeclInfo{node, declInfo})
    self.FP.Writeln(_env, " {")
    self.FP.PushIndent(_env, nil)
    if convGo_isMain_1080_(_env, funcType){
        self.FP.Writeln(_env, _env.LuaVM.String_format("Lns_%s_init( %s )", []LnsAny{self.FP.getModuleName(_env, self.moduleTypeInfo, false), convGo_getAddEnvArg_1279_(_env, 0, self.option.FP.Get_addEnvArg(_env))}))
    }
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
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(_env, "__func__")})/* 2503:29 */)).(*Ast_SymbolInfo)
        self.FP.Writeln(_env, _env.LuaVM.String_format("%s := \"%s.%s\"", []LnsAny{self.FP.getSymbolSym(_env, funcSym_), nameSpace, funcName}))
    }
    for _, _convArg := range( convFunc.FP.Get_argList(_env).Items ) {
        convArg := _convArg.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if Lns_isCondTrue( convArg.FP.Get_posForModToRef(_env)){
            self.FP.Write(_env, _env.LuaVM.String_format("%s := _%s", []LnsAny{self.FP.getSymbolSym(_env, convArg), self.FP.getSymbolSym(_env, convArg)}))
            self.FP.outputAny2Type(_env, convArg.FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
    }
    self.FP.PopIndent(_env)
    {
        _body := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( _body ) {
            body := _body.(*Nodes_BlockNode)
            convGo_filter_1292_(_env, &body.Nodes_Node, self, node)
            var retTypeInfoList *LnsList
            retTypeInfoList = funcType.FP.Get_retTypeInfoList(_env)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( retTypeInfoList.Len() > 0) &&
                _env.SetStackVal( retTypeInfoList.GetAt(retTypeInfoList.Len()).(Ast_TypeInfoDownCast).ToAst_TypeInfo() != Ast_builtinTypeNeverRet) ).(bool)){
                var needReturn bool
                needReturn = false
                {
                    var _from14150 LnsInt = body.FP.Get_stmtList(_env).Len()
                    var _to14150 LnsInt = 1
                    _work14150 := _from14150
                    _delta14150 := -1
                    for {
                        if _delta14150 > 0 {
                           if _work14150 > _to14150 { break }
                        } else {
                           if _work14150 < _to14150 { break }
                        }
                        index := _work14150
                        var statment *Nodes_Node
                        statment = body.FP.Get_stmtList(_env).GetAt(index).(Nodes_NodeDownCast).ToNodes_Node()
                        if _switch14148 := statment.FP.Get_kind(_env); _switch14148 == Nodes_NodeKind_get_BlankLine(_env) {
                        } else if _switch14148 == Nodes_NodeKind_get_Return(_env) {
                            break
                        } else {
                            needReturn = true
                            
                            break
                        }
                        _work14150 += _delta14150
                    }
                }
                if needReturn{
                    self.FP.outputDummyReturn(_env, funcType.FP.Get_retTypeInfoList(_env))
                }
            }
        }
    }
    self.FP.Write(_env, "}")
    if Lns_isCondTrue( declInfo.FP.Get_name(_env)){
        self.FP.Writeln(_env, "")
    }
}

// 2556: decl @lune.@base.@convGo.convFilter.processDeclMethod
func (self *convGo_convFilter) ProcessDeclMethod(_env *LnsEnv, node *Nodes_DeclMethodNode,_opt LnsAny) {
    self.FP.OutputDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env))
}

// 2562: decl @lune.@base.@convGo.convFilter.processProtoMethod
func (self *convGo_convFilter) ProcessProtoMethod(_env *LnsEnv, node *Nodes_ProtoMethodNode,_opt LnsAny) {
}

// 2568: decl @lune.@base.@convGo.convFilter.getEnumGetTxtSym
func (self *convGo_convFilter) getEnumGetTxtSym(_env *LnsEnv, enumType *Ast_EnumTypeInfo) string {
    return _env.LuaVM.String_format("%s_getTxt", []LnsAny{self.FP.getTypeSymbol(_env, &enumType.Ast_TypeInfo)})
}

// 2572: decl @lune.@base.@convGo.convFilter.processDeclEnum
func (self *convGo_convFilter) ProcessDeclEnum(_env *LnsEnv, node *Nodes_DeclEnumNode,_opt LnsAny) {
    if _switch14807 := self.processMode; _switch14807 == convGo_ProcessMode__DeclTopScopeVar {
        self.FP.Writeln(_env, _env.LuaVM.String_format("// decl enum -- %s ", []LnsAny{node.FP.Get_enumType(_env).FP.GetTxt(_env, nil, nil, nil)}))
        self.FP.Writeln(_env, _env.LuaVM.String_format("type %s = %s", []LnsAny{self.FP.getTypeSymbol(_env, &node.FP.Get_enumType(_env).Ast_TypeInfo), self.FP.type2gotype(_env, &node.FP.Get_enumType(_env).Ast_TypeInfo)}))
        {
            __collection14419 := node.FP.Get_enumType(_env).FP.Get_name2EnumValInfo(_env)
            __sorted14419 := __collection14419.CreateKeyListStr()
            __sorted14419.Sort( LnsItemKindStr, nil )
            for _, ___key14419 := range( __sorted14419.Items ) {
                valInfo := __collection14419.Items[ ___key14419 ].(Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
                self.FP.Write(_env, _env.LuaVM.String_format("const %s = ", []LnsAny{self.FP.getSymbol(_env, &convGo_SymbolKind__Static{&node.FP.Get_enumType(_env).Ast_TypeInfo}, valInfo.FP.Get_name(_env))}))
                switch _exp14409 := valInfo.FP.Get_val(_env).(type) {
                case *Ast_EnumLiteral__Int:
                val := _exp14409.Val1
                    self.FP.Write(_env, _env.LuaVM.String_format("%d", []LnsAny{val}))
                case *Ast_EnumLiteral__Real:
                val := _exp14409.Val1
                    self.FP.Write(_env, _env.LuaVM.String_format("%g", []LnsAny{val}))
                case *Ast_EnumLiteral__Str:
                val := _exp14409.Val1
                    self.FP.Write(_env, convGo_str2gostr_1644_(_env, _env.LuaVM.String_format("\"%s\"", []LnsAny{val})))
                }
                self.FP.Writeln(_env, "")
            }
        }
        var listName string
        listName = _env.LuaVM.String_format("%sList_", []LnsAny{self.FP.getTypeSymbol(_env, &node.FP.Get_enumType(_env).Ast_TypeInfo)})
        self.FP.Writeln(_env, _env.LuaVM.String_format("var %s = NewLnsList( []LnsAny {", []LnsAny{listName}))
        for _, _valName := range( node.FP.Get_valueNameList(_env).Items ) {
            valName := _valName.(Types_TokenDownCast).ToTypes_Token()
            self.FP.Writeln(_env, _env.LuaVM.String_format("  %s,", []LnsAny{self.FP.getSymbol(_env, &convGo_SymbolKind__Static{&node.FP.Get_enumType(_env).Ast_TypeInfo}, valName.Txt)}))
        }
        self.FP.Writeln(_env, "})")
        var scope *Ast_Scope
        scope = Lns_unwrap( node.FP.Get_enumType(_env).FP.Get_scope(_env)).(*Ast_Scope)
        self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__Type{Lns_unwrap( scope.FP.GetTypeInfoChild(_env, "get__allList")).(*Ast_TypeInfo)})
        self.FP.Writeln(_env, "{")
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, _env.LuaVM.String_format("return %s", []LnsAny{listName}))
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
        var mapName string
        mapName = _env.LuaVM.String_format("%sMap_", []LnsAny{self.FP.getTypeSymbol(_env, &node.FP.Get_enumType(_env).Ast_TypeInfo)})
        self.FP.Writeln(_env, _env.LuaVM.String_format("var %s = map[%s]string {", []LnsAny{mapName, self.FP.type2gotype(_env, node.FP.Get_enumType(_env).FP.Get_valTypeInfo(_env))}))
        {
            __collection14649 := node.FP.Get_enumType(_env).FP.Get_name2EnumValInfo(_env)
            __sorted14649 := __collection14649.CreateKeyListStr()
            __sorted14649.Sort( LnsItemKindStr, nil )
            for _, ___key14649 := range( __sorted14649.Items ) {
                valInfo := __collection14649.Items[ ___key14649 ].(Ast_EnumValInfoDownCast).ToAst_EnumValInfo()
                self.FP.Writeln(_env, _env.LuaVM.String_format("  %s: \"%s.%s\",", []LnsAny{self.FP.getSymbol(_env, &convGo_SymbolKind__Static{&node.FP.Get_enumType(_env).Ast_TypeInfo}, valInfo.FP.Get_name(_env)), node.FP.Get_expType(_env).FP.Get_rawTxt(_env), valInfo.FP.Get_name(_env)}))
            }
        }
        self.FP.Writeln(_env, "}")
        self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__Type{Lns_unwrap( scope.FP.GetTypeInfoChild(_env, "_from")).(*Ast_TypeInfo)})
        self.FP.Writeln(_env, "{")
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, _env.LuaVM.String_format("if _, ok := %s[arg1]; ok { return arg1 }", []LnsAny{mapName}))
        self.FP.Writeln(_env, "return nil")
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
        self.FP.Writeln(_env, "")
        self.FP.Writeln(_env, _env.LuaVM.String_format("func %s(arg1 %s) string {", []LnsAny{self.FP.getEnumGetTxtSym(_env, node.FP.Get_enumType(_env)), self.FP.type2gotype(_env, node.FP.Get_enumType(_env).FP.Get_valTypeInfo(_env))}))
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, _env.LuaVM.String_format("return %s[arg1];", []LnsAny{mapName}))
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    } else {
    }
}

// 2654: decl @lune.@base.@convGo.convFilter.processUnwrapSet
func (self *convGo_convFilter) ProcessUnwrapSet(_env *LnsEnv, node *Nodes_UnwrapSetNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processUnwrapSet"
    Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
}

// 2660: decl @lune.@base.@convGo.convFilter.processIfUnwrap
func (self *convGo_convFilter) ProcessIfUnwrap(_env *LnsEnv, node *Nodes_IfUnwrapNode,_opt LnsAny) {
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    var tempTypeList *LnsList
    tempTypeList = NewLnsList([]LnsAny{})
    for _index, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        index := _index + 1
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if index > 1{
            self.FP.Write(_env, ", ")
        }
        if varSym.FP.Get_name(_env) == "_"{
            self.FP.Write(_env, "_")
        } else { 
            self.FP.Write(_env, "_" + self.FP.getSymbolSym(_env, varSym))
        }
        tempTypeList.Insert(Ast_TypeInfo2Stem(Ast_builtinTypeStem_))
    }
    if convGo_getExpListKind_1885_(_env, tempTypeList, node.FP.Get_expList(_env), self.option.FP.Get_addEnvArg(_env)) == convGo_ExpListKind__Direct{
        {
            var _from14976 LnsInt = node.FP.Get_varSymList(_env).Len() + 1
            var _to14976 LnsInt = node.FP.Get_expList(_env).FP.Get_expTypeList(_env).Len()
            for _work14976 := _from14976; _work14976 <= _to14976; _work14976++ {
                self.FP.Write(_env, ", _")
            }
        }
    }
    self.FP.Write(_env, " := ")
    self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), node.FP.Get_expList(_env)), tempTypeList, node.FP.Get_expList(_env), false)
    self.FP.Writeln(_env, "")
    self.FP.Write(_env, "if ")
    var hasSym bool
    hasSym = false
    for _, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if varSym.FP.Get_name(_env) != "_"{
            if hasSym{
                self.FP.Write(_env, " && ")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("!Lns_IsNil( _%s )", []LnsAny{self.FP.getSymbolSym(_env, varSym)}))
            hasSym = true
            
        }
    }
    self.FP.Writeln(_env, " {")
    self.FP.PushIndent(_env, nil)
    for _index, _varSym := range( node.FP.Get_varSymList(_env).Items ) {
        index := _index + 1
        varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
        if varSym.FP.Get_name(_env) != "_"{
            if varSym.FP.HasAccess(_env){
                self.FP.Write(_env, _env.LuaVM.String_format("%s := _%s", []LnsAny{self.FP.getSymbolSym(_env, varSym), self.FP.getSymbolSym(_env, varSym)}))
                if node.FP.Get_expList(_env).FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                    self.FP.outputAny2Type(_env, varSym.FP.Get_typeInfo(_env))
                }
                self.FP.Writeln(_env, "")
            }
        }
    }
    self.FP.PopIndent(_env)
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    {
        _nilBlock := node.FP.Get_nilBlock(_env)
        if !Lns_IsNil( _nilBlock ) {
            nilBlock := _nilBlock.(*Nodes_BlockNode)
            self.FP.Writeln(_env, "} else {")
            convGo_filter_1292_(_env, &nilBlock.Nodes_Node, self, &node.Nodes_Node)
            self.FP.Writeln(_env, "}")
        } else {
            self.FP.Writeln(_env, "}")
        }
    }
    self.FP.PopIndent(_env)
    self.FP.Write(_env, "}")
    self.FP.Writeln(_env, "")
}

// 2738: decl @lune.@base.@convGo.convFilter.outputLetVar
func (self *convGo_convFilter) outputLetVar(_env *LnsEnv, node *Nodes_DeclVarNode) {
    var declVar func(_env *LnsEnv)
    declVar = func(_env *LnsEnv) {
        if node.FP.Get_symbolInfoList(_env).GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_scope(_env) == self.moduleScope{
            return 
        }
        for _, _symbolInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
            symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if Lns_isCondTrue( symbolInfo.FP.Get_posForModToRef(_env)){
                self.FP.Writeln(_env, _env.LuaVM.String_format("var %s %s", []LnsAny{self.FP.getSymbolSym(_env, symbolInfo), self.FP.type2gotype(_env, symbolInfo.FP.Get_typeInfo(_env))}))
            }
        }
    }
    if node.FP.Get_unwrapFlag(_env){
        {
            _expList, _unwrapBlock := node.FP.Get_expList(_env), node.FP.Get_unwrapBlock(_env)
            if !Lns_IsNil( _expList ) && !Lns_IsNil( _unwrapBlock ) {
                expList := _expList.(*Nodes_ExpListNode)
                unwrapBlock := _unwrapBlock.(*Nodes_BlockNode)
                if node.FP.Get_symbolInfoList(_env).GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo().FP.Get_scope(_env) != self.moduleScope{
                    declVar(_env)
                }
                self.FP.Writeln(_env, "")
                self.FP.Writeln(_env, "{")
                self.FP.PushIndent(_env, nil)
                for _index, _varInfo := range( node.FP.Get_varList(_env).Items ) {
                    index := _index + 1
                    varInfo := _varInfo.(Nodes_VarInfoDownCast).ToNodes_VarInfo()
                    if index != 1{
                        self.FP.Write(_env, ", ")
                    }
                    self.FP.Write(_env, _env.LuaVM.String_format("_%s", []LnsAny{convGo_normalizeSym_1451_(_env, varInfo.FP.Get_name(_env).Txt)}))
                }
                var tmpVarTypeList *LnsList
                tmpVarTypeList = NewLnsList([]LnsAny{})
                for _index, _ := range( node.FP.Get_symbolInfoList(_env).Items ) {
                    index := _index + 1
                    tmpVarTypeList.Insert(Ast_TypeInfo2Stem(expList.FP.GetExpTypeNoDDDAt(_env, index)))
                }
                if convGo_getExpListKind_1885_(_env, tmpVarTypeList, expList, self.option.FP.Get_addEnvArg(_env)) == convGo_ExpListKind__Direct{
                    {
                        var _from15483 LnsInt = tmpVarTypeList.Len() + 1
                        var _to15483 LnsInt = expList.FP.Get_expTypeList(_env).Len()
                        for _work15483 := _from15483; _work15483 <= _to15483; _work15483++ {
                            self.FP.Write(_env, ", _")
                        }
                    }
                }
                self.FP.Write(_env, " := ")
                self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), expList), tmpVarTypeList, expList, false)
                self.FP.Writeln(_env, "")
                self.FP.Write(_env, "if ")
                var hasCond bool
                hasCond = false
                for _index, _varInfo := range( node.FP.Get_varList(_env).Items ) {
                    index := _index + 1
                    varInfo := _varInfo.(Nodes_VarInfoDownCast).ToNodes_VarInfo()
                    if expList.FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                        if hasCond{
                            self.FP.Write(_env, " || ")
                        }
                        self.FP.Write(_env, _env.LuaVM.String_format("_%s == nil", []LnsAny{convGo_normalizeSym_1451_(_env, varInfo.FP.Get_name(_env).Txt)}))
                        hasCond = true
                        
                    }
                }
                self.FP.Writeln(_env, "{")
                convGo_filter_1292_(_env, &unwrapBlock.Nodes_Node, self, &node.Nodes_Node)
                {
                    _thenBlock := node.FP.Get_thenBlock(_env)
                    if !Lns_IsNil( _thenBlock ) {
                        thenBlock := _thenBlock.(*Nodes_BlockNode)
                        self.FP.Writeln(_env, "} else {")
                        self.FP.PushIndent(_env, nil)
                        for _index, _varInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
                            index := _index + 1
                            varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            self.FP.Write(_env, _env.LuaVM.String_format("%s = _%s", []LnsAny{self.FP.getSymbolSym(_env, varInfo), convGo_normalizeSym_1451_(_env, varInfo.FP.Get_name(_env))}))
                            if expList.FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                                self.FP.outputAny2Type(_env, varInfo.FP.Get_typeInfo(_env))
                            }
                            self.FP.Writeln(_env, "")
                        }
                        self.FP.PopIndent(_env)
                        convGo_filter_1292_(_env, &thenBlock.Nodes_Node, self, &node.Nodes_Node)
                        self.FP.Writeln(_env, "}")
                    } else {
                        self.FP.Writeln(_env, "} else {")
                        self.FP.PushIndent(_env, nil)
                        for _index, _varInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
                            index := _index + 1
                            varInfo := _varInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                            self.FP.Write(_env, _env.LuaVM.String_format("%s = _%s", []LnsAny{self.FP.getSymbolSym(_env, varInfo), convGo_normalizeSym_1451_(_env, varInfo.FP.Get_name(_env))}))
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
        declVar(_env)
        {
            _expList := node.FP.Get_expList(_env)
            if !Lns_IsNil( _expList ) {
                expList := _expList.(*Nodes_ExpListNode)
                var varTypeList *LnsList
                varTypeList = NewLnsList([]LnsAny{})
                for _index, _symbolInfo := range( node.FP.Get_symbolInfoList(_env).Items ) {
                    index := _index + 1
                    symbolInfo := _symbolInfo.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                    varTypeList.Insert(Ast_TypeInfo2Stem(symbolInfo.FP.Get_typeInfo(_env)))
                    if index > 1{
                        self.FP.Write(_env, ",")
                    }
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( symbolInfo.FP.Get_posForModToRef(_env)) ||
                        _env.SetStackVal( Ast_isPubToExternal(_env, symbolInfo.FP.Get_accessMode(_env))) )){
                        self.FP.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{self.FP.getSymbolSym(_env, symbolInfo)}))
                    } else { 
                        self.FP.Write(_env, "_")
                    }
                }
                if convGo_getExpListKind_1885_(_env, varTypeList, expList, self.option.FP.Get_addEnvArg(_env)) == convGo_ExpListKind__Direct{
                    {
                        var _from15957 LnsInt = varTypeList.Len() + 1
                        var _to15957 LnsInt = expList.FP.Get_expTypeList(_env).Len()
                        for _work15957 := _from15957; _work15957 <= _to15957; _work15957++ {
                            self.FP.Write(_env, ", _")
                        }
                    }
                }
                self.FP.Write(_env, " = ")
                self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), expList), varTypeList, expList, false)
                self.FP.Writeln(_env, "")
            }
        }
    }
}

// 2869: decl @lune.@base.@convGo.convFilter.processDeclVar
func (self *convGo_convFilter) ProcessDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processDeclVar"
    if _switch16460 := node.FP.Get_mode(_env); _switch16460 == Nodes_DeclVarMode__Let {
        self.FP.outputLetVar(_env, node)
    } else if _switch16460 == Nodes_DeclVarMode__Unwrap {
        self.FP.Writeln(_env, "{")
        self.FP.PushIndent(_env, nil)
        for _, _varSym := range( node.FP.Get_symbolInfoList(_env).Items ) {
            varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            self.FP.Writeln(_env, _env.LuaVM.String_format("var _%s LnsAny", []LnsAny{varSym.FP.Get_name(_env)}))
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
        var setVals func(_env *LnsEnv)
        setVals = func(_env *LnsEnv) {
            for _index, _varSym := range( node.FP.Get_symbolInfoList(_env).Items ) {
                index := _index + 1
                varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                self.FP.Write(_env, _env.LuaVM.String_format("%s = _%s", []LnsAny{self.FP.getSymbolSym(_env, varSym), varSym.FP.Get_name(_env)}))
                if expList.FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                    self.FP.outputAny2Type(_env, varSym.FP.Get_typeInfo(_env))
                }
                self.FP.Writeln(_env, "")
            }
        }
        var typeList *LnsList
        typeList = NewLnsList([]LnsAny{})
        for _index, _varSym := range( node.FP.Get_symbolInfoList(_env).Items ) {
            index := _index + 1
            varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            typeList.Insert(Ast_TypeInfo2Stem(varSym.FP.Get_typeInfo(_env)))
            if index > 1{
                self.FP.Write(_env, ",")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("_%s", []LnsAny{varSym.FP.Get_name(_env)}))
        }
        if convGo_getExpListKind_1885_(_env, typeList, expList, self.option.FP.Get_addEnvArg(_env)) == convGo_ExpListKind__Direct{
            {
                var _from16237 LnsInt = node.FP.Get_symbolInfoList(_env).Len() + 1
                var _to16237 LnsInt = expList.FP.Get_expTypeList(_env).Len()
                for _work16237 := _from16237; _work16237 <= _to16237; _work16237++ {
                    self.FP.Write(_env, ",_")
                }
            }
        }
        self.FP.Write(_env, " = ")
        self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), expList), typeList, expList, false)
        self.FP.Writeln(_env, "")
        self.FP.Write(_env, "if ")
        var hasCond bool
        hasCond = false
        for _index, _varSym := range( node.FP.Get_symbolInfoList(_env).Items ) {
            index := _index + 1
            varSym := _varSym.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if expList.FP.GetExpTypeNoDDDAt(_env, index).FP.Get_nilable(_env){
                if hasCond{
                    self.FP.Write(_env, " || ")
                }
                self.FP.Write(_env, _env.LuaVM.String_format("Lns_IsNil( _%s )", []LnsAny{varSym.FP.Get_name(_env)}))
                hasCond = true
                
            }
        }
        self.FP.Writeln(_env, " {")
        convGo_filter_1292_(_env, &Lns_unwrap( node.FP.Get_unwrapBlock(_env)).(*Nodes_BlockNode).Nodes_Node, self, &node.Nodes_Node)
        {
            _thenBlock := node.FP.Get_thenBlock(_env)
            if !Lns_IsNil( _thenBlock ) {
                thenBlock := _thenBlock.(*Nodes_BlockNode)
                self.FP.Writeln(_env, "} else {")
                self.FP.PushIndent(_env, nil)
                setVals(_env)
                self.FP.PopIndent(_env)
                convGo_filter_1292_(_env, &thenBlock.Nodes_Node, self, &node.Nodes_Node)
                self.FP.Writeln(_env, "}")
            } else {
                self.FP.Writeln(_env, "}")
                setVals(_env)
            }
        }
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    } else {
        Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
    }
}

// 2949: decl @lune.@base.@convGo.convFilter.processWhen
func (self *convGo_convFilter) ProcessWhen(_env *LnsEnv, node *Nodes_WhenNode,_opt LnsAny) {
    self.FP.Write(_env, "if ")
    for _index, _symPair := range( node.FP.Get_symPairList(_env).Items ) {
        index := _index + 1
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        if index > 1{
            self.FP.Write(_env, " && ")
        }
        self.FP.Write(_env, _env.LuaVM.String_format("%s != nil", []LnsAny{self.FP.getSymbolSym(_env, symPair.FP.Get_src(_env))}))
        symPair.FP.Get_dst(_env).FP.Set_convModuleParam(_env, true)
    }
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    for _, _symPair := range( node.FP.Get_symPairList(_env).Items ) {
        symPair := _symPair.(Nodes_UnwrapSymbolPairDownCast).ToNodes_UnwrapSymbolPair()
        if Lns_isCondTrue( symPair.FP.Get_dst(_env).FP.Get_posForModToRef(_env)){
            self.FP.Write(_env, _env.LuaVM.String_format("%s_%d := %s", []LnsAny{symPair.FP.Get_dst(_env).FP.Get_name(_env), symPair.FP.Get_dst(_env).FP.Get_symbolId(_env), self.FP.getSymbolSym(_env, symPair.FP.Get_src(_env))}))
            self.FP.outputAny2Type(_env, symPair.FP.Get_dst(_env).FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
    }
    self.FP.PopIndent(_env)
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    {
        _elseBlock := node.FP.Get_elseBlock(_env)
        if !Lns_IsNil( _elseBlock ) {
            elseBlock := _elseBlock.(*Nodes_BlockNode)
            self.FP.Writeln(_env, "} else {")
            convGo_filter_1292_(_env, &elseBlock.Nodes_Node, self, &node.Nodes_Node)
            self.FP.Write(_env, "}")
        } else {
            self.FP.Write(_env, "}")
        }
    }
    self.FP.Writeln(_env, "")
}

// 2986: decl @lune.@base.@convGo.convFilter.processDeclArg
func (self *convGo_convFilter) ProcessDeclArg(_env *LnsEnv, node *Nodes_DeclArgNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("%s ", []LnsAny{self.FP.getSymbolSym(_env, node.FP.Get_symbolInfo(_env))}))
    convGo_filter_1292_(_env, &Lns_unwrap( node.FP.Get_argType(_env)).(*Nodes_RefTypeNode).Nodes_Node, self, &node.Nodes_Node)
}

// 2993: decl @lune.@base.@convGo.convFilter.processDeclArgDDD
func (self *convGo_convFilter) ProcessDeclArgDDD(_env *LnsEnv, node *Nodes_DeclArgDDDNode,_opt LnsAny) {
    self.FP.Write(_env, "ddd []LnsAny")
}

// 2999: decl @lune.@base.@convGo.convFilter.processExpSubDDD
func (self *convGo_convFilter) ProcessExpSubDDD(_env *LnsEnv, node *Nodes_ExpSubDDDNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpSubDDD"
    Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
}

// 3006: decl @lune.@base.@convGo.convFilter.processDeclForm
func (self *convGo_convFilter) ProcessDeclForm(_env *LnsEnv, node *Nodes_DeclFormNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("type %s ", []LnsAny{self.FP.getFuncSymbol(_env, node.FP.Get_expType(_env))}))
    self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__Type{node.FP.Get_expType(_env)})
    self.FP.Writeln(_env, "")
}

// 3015: decl @lune.@base.@convGo.convFilter.processDeclFunc
func (self *convGo_convFilter) ProcessDeclFunc(_env *LnsEnv, node *Nodes_DeclFuncNode,_opt LnsAny) {
    {
        _funcSym := node.FP.Get_declInfo(_env).FP.Get_symbol(_env)
        if !Lns_IsNil( _funcSym ) {
            funcSym := _funcSym.(*Ast_SymbolInfo)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(funcSym.FP.Get_posForModToRef(_env))) &&
                _env.SetStackVal( Lns_op_not(Ast_isPubToExternal(_env, funcSym.FP.Get_accessMode(_env)))) ).(bool)){
                return 
            }
        }
    }
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_expType(_env)
    if (self.processMode == convGo_ProcessMode__NonClosureFuncDecl) == convGo_isClosure_1307_(_env, node.FP.Get_expType(_env)){
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.processMode != convGo_ProcessMode__NonClosureFuncDecl) &&
            _env.SetStackVal( Lns_op_not(node.FP.Get_declInfo(_env).FP.Get_symbol(_env))) ).(bool)){
            self.FP.Write(_env, self.FP.getFuncSymbol(_env, funcType))
        }
        return 
    }
    if convGo_isClosure_1307_(_env, funcType){
        {
            _funcSym := node.FP.Get_declInfo(_env).FP.Get_symbol(_env)
            if !Lns_IsNil( _funcSym ) {
                funcSym := _funcSym.(*Ast_SymbolInfo)
                self.FP.Write(_env, "var ")
                self.FP.outputSymbol(_env, &convGo_SymbolKind__Func{funcType}, funcSym.FP.Get_name(_env))
                self.FP.Write(_env, " ")
                self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__DeclInfo{&node.Nodes_Node, node.FP.Get_declInfo(_env)})
                self.FP.Writeln(_env, "")
                self.FP.outputSymbol(_env, &convGo_SymbolKind__Func{funcType}, funcSym.FP.Get_name(_env))
                self.FP.Write(_env, " = ")
            }
        }
    }
    self.FP.OutputDeclFuncInfo(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env))
}

// 3048: decl @lune.@base.@convGo.convFilter.processRefType
func (self *convGo_convFilter) ProcessRefType(_env *LnsEnv, node *Nodes_RefTypeNode,_opt LnsAny) {
    self.FP.Write(_env, self.FP.type2gotype(_env, node.FP.Get_expType(_env)))
}

// 3070: decl @lune.@base.@convGo.convFilter.processCond
func (self *convGo_convFilter) processCond(_env *LnsEnv, node *Nodes_Node,parent *Nodes_Node) {
    if node.FP.Get_expType(_env) != Ast_builtinTypeBool{
        self.FP.Write(_env, "Lns_isCondTrue( ")
        convGo_filter_1292_(_env, node, self, parent)
        self.FP.Write(_env, ")")
    } else { 
        convGo_filter_1292_(_env, node, self, parent)
    }
}

// 3081: decl @lune.@base.@convGo.convFilter.processIf
func (self *convGo_convFilter) ProcessIf(_env *LnsEnv, node *Nodes_IfNode,_opt LnsAny) {
    for _, _stmt := range( node.FP.Get_stmtList(_env).Items ) {
        stmt := _stmt.(Nodes_IfStmtInfoDownCast).ToNodes_IfStmtInfo()
        if _switch17251 := stmt.FP.Get_kind(_env); _switch17251 == Nodes_IfKind__If {
            self.FP.Write(_env, "if ")
            self.FP.processCond(_env, stmt.FP.Get_exp(_env), &node.Nodes_Node)
            self.FP.Writeln(_env, "{")
            convGo_filter_1292_(_env, &stmt.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        } else if _switch17251 == Nodes_IfKind__ElseIf {
            self.FP.Write(_env, "} else if ")
            self.FP.processCond(_env, stmt.FP.Get_exp(_env), &node.Nodes_Node)
            self.FP.Writeln(_env, "{")
            convGo_filter_1292_(_env, &stmt.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        } else if _switch17251 == Nodes_IfKind__Else {
            self.FP.Writeln(_env, "} else { ")
            convGo_filter_1292_(_env, &stmt.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln(_env, "}")
}

// 3107: decl @lune.@base.@convGo.convFilter.processSwitch
func (self *convGo_convFilter) ProcessSwitch(_env *LnsEnv, node *Nodes_SwitchNode,_opt LnsAny) {
    var valName string
    valName = _env.LuaVM.String_format("_switch%d", []LnsAny{node.FP.Get_id(_env)})
    self.FP.Write(_env, _env.LuaVM.String_format("if %s := ", []LnsAny{valName}))
    convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, "; ")
    for _caseIndex, _caseNode := range( node.FP.Get_caseList(_env).Items ) {
        caseIndex := _caseIndex + 1
        caseNode := _caseNode.(Nodes_CaseInfoDownCast).ToNodes_CaseInfo()
        if caseIndex != 1{
            self.FP.Write(_env, "} else if ")
        }
        for _index, _exp := range( caseNode.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
            index := _index + 1
            exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
            if index != 1{
                self.FP.Write(_env, " || ")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("%s == ", []LnsAny{valName}))
            convGo_filter_1292_(_env, exp, self, &caseNode.FP.Get_expList(_env).Nodes_Node)
        }
        self.FP.Writeln(_env, " {")
        convGo_filter_1292_(_env, &caseNode.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    }
    {
        _defaultBlock := node.FP.Get_default(_env)
        if !Lns_IsNil( _defaultBlock ) {
            defaultBlock := _defaultBlock.(*Nodes_BlockNode)
            self.FP.Writeln(_env, "} else {")
            convGo_filter_1292_(_env, &defaultBlock.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln(_env, "}")
}

// 3140: decl @lune.@base.@convGo.convFilter.processMatch
func (self *convGo_convFilter) ProcessMatch(_env *LnsEnv, node *Nodes_MatchNode,_opt LnsAny) {
    var hasAccessing func(_env *LnsEnv) bool
    hasAccessing = func(_env *LnsEnv) bool {
        for _, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
            caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
            for _, _symbol := range( caseInfo.FP.Get_valParamNameList(_env).Items ) {
                symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
                if Lns_isCondTrue( symbol.FP.Get_posForModToRef(_env)){
                    return true
                }
            }
        }
        return false
    }
    var val string
    if hasAccessing(_env){
        val = _env.LuaVM.String_format("_exp%d", []LnsAny{node.FP.Get_id(_env)})
        
        self.FP.Write(_env, _env.LuaVM.String_format("switch %s := ", []LnsAny{val}))
    } else { 
        val = ""
        
        self.FP.Write(_env, "switch ")
    }
    convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
    self.FP.Writeln(_env, ".(type) {")
    for _, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
        caseInfo := _caseInfo.(Nodes_MatchCaseDownCast).ToNodes_MatchCase()
        self.FP.Writeln(_env, _env.LuaVM.String_format("case *%s:", []LnsAny{self.FP.getAlgeSymbol(_env, caseInfo.FP.Get_valInfo(_env))}))
        for _index, _symbol := range( caseInfo.FP.Get_valParamNameList(_env).Items ) {
            index := _index + 1
            symbol := _symbol.(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            if Lns_isCondTrue( symbol.FP.Get_posForModToRef(_env)){
                self.FP.Writeln(_env, _env.LuaVM.String_format("%s := %s.Val%d", []LnsAny{self.FP.getSymbolSym(_env, symbol), val, index}))
            }
        }
        convGo_filter_1292_(_env, &caseInfo.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    }
    {
        _defaultBlock := node.FP.Get_defaultBlock(_env)
        if !Lns_IsNil( _defaultBlock ) {
            defaultBlock := _defaultBlock.(*Nodes_Node)
            self.FP.Writeln(_env, "default:")
            convGo_filter_1292_(_env, defaultBlock, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln(_env, "}")
}

// 3180: decl @lune.@base.@convGo.convFilter.processWhile
func (self *convGo_convFilter) ProcessWhile(_env *LnsEnv, node *Nodes_WhileNode,_opt LnsAny) {
    self.FP.Write(_env, "for ")
    if Lns_op_not(node.FP.Get_infinit(_env)){
        self.FP.processCond(_env, node.FP.Get_exp(_env), &node.Nodes_Node)
    }
    self.FP.Writeln(_env, " {")
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "}")
}

// 3194: decl @lune.@base.@convGo.convFilter.processRepeat
func (self *convGo_convFilter) ProcessRepeat(_env *LnsEnv, node *Nodes_RepeatNode,_opt LnsAny) {
    self.FP.Writeln(_env, "for {")
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.PushIndent(_env, nil)
    self.FP.Write(_env, "if ")
    self.FP.processCond(_env, node.FP.Get_exp(_env), &node.Nodes_Node)
    self.FP.Writeln(_env, "{ break }")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 3212: decl @lune.@base.@convGo.convFilter.processFor
func (self *convGo_convFilter) ProcessFor(_env *LnsEnv, node *Nodes_ForNode,_opt LnsAny) {
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    var fromSym string
    fromSym = _env.LuaVM.String_format("_from%d", []LnsAny{node.FP.Get_id(_env)})
    var toSym string
    toSym = _env.LuaVM.String_format("_to%d", []LnsAny{node.FP.Get_id(_env)})
    var deltaSym string
    deltaSym = _env.LuaVM.String_format("_delta%d", []LnsAny{node.FP.Get_id(_env)})
    var workSym string
    workSym = _env.LuaVM.String_format("_work%d", []LnsAny{node.FP.Get_id(_env)})
    var valType string
    valType = _env.LuaVM.String_format("%s", []LnsAny{self.FP.type2gotype(_env, node.FP.Get_init(_env).FP.Get_expType(_env))})
    self.FP.Write(_env, _env.LuaVM.String_format("var %s %s = ", []LnsAny{fromSym, valType}))
    convGo_filter_1292_(_env, node.FP.Get_init(_env), self, &node.Nodes_Node)
    self.FP.Writeln(_env, "")
    self.FP.Write(_env, _env.LuaVM.String_format("var %s %s = ", []LnsAny{toSym, valType}))
    convGo_filter_1292_(_env, node.FP.Get_to(_env), self, &node.Nodes_Node)
    self.FP.Writeln(_env, "")
    {
        _delta := node.FP.Get_delta(_env)
        if !Lns_IsNil( _delta ) {
            delta := _delta.(*Nodes_Node)
            self.FP.Writeln(_env, _env.LuaVM.String_format("%s := %s", []LnsAny{workSym, fromSym}))
            self.FP.Write(_env, _env.LuaVM.String_format("%s := ", []LnsAny{deltaSym}))
            convGo_filter_1292_(_env, delta, self, &node.Nodes_Node)
            self.FP.Writeln(_env, "")
            self.FP.Writeln(_env, "for {")
            self.FP.PushIndent(_env, nil)
            self.FP.Writeln(_env, _env.LuaVM.String_format("if %s > 0 {", []LnsAny{deltaSym}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("   if %s > %s { break }", []LnsAny{workSym, toSym}))
            self.FP.Writeln(_env, "} else {")
            self.FP.Writeln(_env, _env.LuaVM.String_format("   if %s < %s { break }", []LnsAny{workSym, toSym}))
            self.FP.Writeln(_env, "}")
            self.FP.PopIndent(_env)
        } else {
            self.FP.Writeln(_env, _env.LuaVM.String_format("for %s := %s; %s <= %s; %s++ {", []LnsAny{workSym, fromSym, workSym, toSym, workSym}))
        }
    }
    self.FP.PushIndent(_env, nil)
    if Lns_isCondTrue( node.FP.Get_val(_env).FP.Get_posForModToRef(_env)){
        self.FP.Writeln(_env, _env.LuaVM.String_format("%s := %s", []LnsAny{node.FP.Get_val(_env).FP.Get_name(_env), workSym}))
    }
    self.FP.PopIndent(_env)
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    if Lns_isCondTrue( node.FP.Get_delta(_env)){
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, _env.LuaVM.String_format("%s += %s", []LnsAny{workSym, deltaSym}))
        self.FP.PopIndent(_env)
    }
    self.FP.Writeln(_env, "}")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 3276: decl @lune.@base.@convGo.convFilter.processApply
func (self *convGo_convFilter) ProcessApply(_env *LnsEnv, node *Nodes_ApplyNode,_opt LnsAny) {
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    var formSym string
    formSym = _env.LuaVM.String_format("_form%d", []LnsAny{node.FP.Get_id(_env)})
    var paramSym string
    paramSym = _env.LuaVM.String_format("_param%d", []LnsAny{node.FP.Get_id(_env)})
    var prevSym string
    prevSym = _env.LuaVM.String_format("_prev%d", []LnsAny{node.FP.Get_id(_env)})
    if Lns_isCondTrue( node.FP.Get_expList(_env).FP.Get_mRetExp(_env)){
        self.FP.Write(_env, _env.LuaVM.String_format("%s, %s, %s := ", []LnsAny{formSym, paramSym, prevSym}))
        convGo_filter_1292_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, &node.Nodes_Node)
        self.FP.Writeln(_env, "")
    } else { 
        self.FP.Write(_env, _env.LuaVM.String_format("%s, %s, %s := ", []LnsAny{formSym, paramSym, prevSym}))
        convGo_filter_1292_(_env, node.FP.Get_expList(_env).FP.Get_expList(_env).GetAt(1).(Nodes_NodeDownCast).ToNodes_Node(), self, &node.FP.Get_expList(_env).Nodes_Node)
        self.FP.Write(_env, ",")
        convGo_filter_1292_(_env, node.FP.Get_expList(_env).FP.Get_expList(_env).GetAt(2).(Nodes_NodeDownCast).ToNodes_Node(), self, &node.FP.Get_expList(_env).Nodes_Node)
        self.FP.Write(_env, ", LnsAny(")
        convGo_filter_1292_(_env, node.FP.Get_expList(_env).FP.Get_expList(_env).GetAt(3).(Nodes_NodeDownCast).ToNodes_Node(), self, &node.FP.Get_expList(_env).Nodes_Node)
        self.FP.Writeln(_env, ")")
    }
    self.FP.Writeln(_env, "for {")
    self.FP.PushIndent(_env, nil)
    var setTxt string
    setTxt = prevSym
    {
        var _from18573 LnsInt = 2
        var _to18573 LnsInt = node.FP.Get_varList(_env).Len()
        for _work18573 := _from18573; _work18573 <= _to18573; _work18573++ {
            index := _work18573
            var symInfo *Ast_SymbolInfo
            symInfo = node.FP.Get_varList(_env).GetAt(index).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
            self.FP.Writeln(_env, _env.LuaVM.String_format("var %s %s", []LnsAny{self.FP.getSymbolSym(_env, symInfo), self.FP.type2gotype(_env, symInfo.FP.Get_typeInfo(_env))}))
            setTxt = _env.LuaVM.String_format("%s, %s", []LnsAny{setTxt, self.FP.getSymbolSym(_env, symInfo)})
            
        }
    }
    if node.FP.Get_expList(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        var workSym string
        workSym = _env.LuaVM.String_format("_work%d", []LnsAny{node.FP.Get_id(_env)})
        self.FP.Writeln(_env, _env.LuaVM.String_format("%s := %s.(*Lns_luaValue).Call( Lns_2DDD( %s, %s ) )", []LnsAny{workSym, formSym, paramSym, prevSym}))
        self.FP.Write(_env, _env.LuaVM.String_format("%s = ", []LnsAny{setTxt}))
        for _index, _ := range( node.FP.Get_varList(_env).Items ) {
            index := _index + 1
            if index > 1{
                self.FP.Write(_env, ",")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("Lns_getFromMulti(%s,%d)", []LnsAny{workSym, index - 1}))
        }
        self.FP.Writeln(_env, "")
    } else { 
        self.FP.Writeln(_env, _env.LuaVM.String_format("%s = %s( %s %s, %s )", []LnsAny{setTxt, formSym, convGo_getAddEnvArg_1279_(_env, 2, self.option.FP.Get_addEnvArg(_env)), paramSym, prevSym}))
    }
    self.FP.Writeln(_env, _env.LuaVM.String_format("if Lns_IsNil( %s ) { break }", []LnsAny{prevSym}))
    var topSymInfo *Ast_SymbolInfo
    topSymInfo = node.FP.Get_varList(_env).GetAt(1).(Ast_SymbolInfoDownCast).ToAst_SymbolInfo()
    if topSymInfo.FP.Get_name(_env) != "_"{
        self.FP.Writeln(_env, _env.LuaVM.String_format("%s := %s.(%s)", []LnsAny{self.FP.getSymbolSym(_env, topSymInfo), prevSym, self.FP.type2gotype(_env, topSymInfo.FP.Get_typeInfo(_env))}))
    }
    self.FP.PopIndent(_env)
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "}")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 3349: decl @lune.@base.@convGo.convFilter.outputForeachLua
func (self *convGo_convFilter) outputForeachLua(_env *LnsEnv, node *Nodes_Node,sortFlag bool,exp *Nodes_Node,val *Ast_SymbolInfo,key LnsAny,block *Nodes_BlockNode) {
    __func__ := "@lune.@base.@convGo.convFilter.outputForeachLua"
    if _switch19354 := exp.FP.Get_expType(_env).FP.Get_extedType(_env).FP.Get_kind(_env); _switch19354 == Ast_TypeInfoKind__List || _switch19354 == Ast_TypeInfoKind__Map {
        self.FP.Writeln(_env, "{")
        self.FP.PushIndent(_env, nil)
        var tmpExp string
        tmpExp = _env.LuaVM.String_format("_exp%d", []LnsAny{node.FP.Get_id(_env)})
        var tmpKey string
        tmpKey = _env.LuaVM.String_format("_key%d", []LnsAny{node.FP.Get_id(_env)})
        var tmpVal string
        tmpVal = _env.LuaVM.String_format("_val%d", []LnsAny{node.FP.Get_id(_env)})
        var tmpIndex string
        tmpIndex = _env.LuaVM.String_format("_index%d", []LnsAny{node.FP.Get_id(_env)})
        var sorted string
        sorted = _env.LuaVM.String_format("_sorted%d", []LnsAny{node.FP.Get_id(_env)})
        self.FP.Write(_env, _env.LuaVM.String_format("%s := ", []LnsAny{tmpExp}))
        convGo_filter_1292_(_env, exp, self, node)
        self.FP.Writeln(_env, "")
        var tmpValName string
        if val.FP.HasAccess(_env){
            tmpValName = tmpVal
            
        } else { 
            tmpValName = "_"
            
        }
        if sortFlag{
            self.FP.Writeln(_env, _env.LuaVM.String_format("%s := %s.SortMapKeyList( %s )", []LnsAny{sorted, self.env.FP.getLuavm(_env, node.FP.IsThreading(_env)), tmpExp}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("%s, %s := %s.Get1stFromMap()", []LnsAny{tmpIndex, tmpKey, sorted}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("for %s != nil {", []LnsAny{tmpIndex}))
            self.FP.PushIndent(_env, nil)
        } else { 
            self.FP.Writeln(_env, _env.LuaVM.String_format("%s, %s := %s.Get1stFromMap()", []LnsAny{tmpKey, tmpValName, tmpExp}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("for %s != nil {", []LnsAny{tmpKey}))
            self.FP.PushIndent(_env, nil)
        }
        {
            _keySym := key
            if !Lns_IsNil( _keySym ) {
                keySym := _keySym.(*Ast_SymbolInfo)
                if keySym.FP.HasAccess(_env){
                    self.FP.Write(_env, _env.LuaVM.String_format("%s := %s", []LnsAny{self.FP.getSymbolSym(_env, keySym), tmpKey}))
                    self.FP.outputAny2Type(_env, keySym.FP.Get_typeInfo(_env))
                    self.FP.Writeln(_env, "")
                }
            }
        }
        if val.FP.HasAccess(_env){
            if sortFlag{
                self.FP.Write(_env, _env.LuaVM.String_format("%s := %s.GetAt( %s )", []LnsAny{self.FP.getSymbolSym(_env, val), tmpExp, tmpKey}))
            } else { 
                self.FP.Write(_env, _env.LuaVM.String_format("%s := %s", []LnsAny{self.FP.getSymbolSym(_env, val), tmpVal}))
            }
            self.FP.outputAny2Type(_env, val.FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
        self.FP.PopIndent(_env)
        convGo_filter_1292_(_env, &block.Nodes_Node, self, node)
        self.FP.PushIndent(_env, nil)
        if Lns_op_not(sortFlag){
            self.FP.Write(_env, _env.LuaVM.String_format("%s, %s = ", []LnsAny{tmpKey, tmpValName}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("%s.NextFromMap( %s )", []LnsAny{tmpExp, tmpKey}))
        } else { 
            self.FP.Write(_env, _env.LuaVM.String_format("%s, %s = ", []LnsAny{tmpIndex, tmpKey}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("%s.NextFromMap( %s )", []LnsAny{sorted, tmpIndex}))
        }
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    } else {
        Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
    }
}

// 3429: decl @lune.@base.@convGo.convFilter.processForeach
func (self *convGo_convFilter) ProcessForeach(_env *LnsEnv, node *Nodes_ForeachNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processForeach"
    if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_srcTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        self.FP.outputForeachLua(_env, &node.Nodes_Node, false, node.FP.Get_exp(_env), node.FP.Get_val(_env), node.FP.Get_key(_env), node.FP.Get_block(_env))
        return 
    }
    var hasAccessLoopSym LnsAny
    hasAccessLoopSym = _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(node.FP.Get_key(_env)) && 
        Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_SymbolInfo).FP.Get_posForModToRef(_env)}))) ||
        _env.SetStackVal( node.FP.Get_val(_env).FP.Get_posForModToRef(_env)) )
    self.FP.Write(_env, "for ")
    var loopExpType *Ast_TypeInfo
    loopExpType = node.FP.Get_exp(_env).FP.Get_expType(_env)
    if _switch20238 := loopExpType.FP.Get_kind(_env); _switch20238 == Ast_TypeInfoKind__List || _switch20238 == Ast_TypeInfoKind__Array {
        var valName string
        valName = self.FP.getSymbolSym(_env, node.FP.Get_val(_env))
        var itemType *Ast_TypeInfo
        itemType = loopExpType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_isCondTrue( hasAccessLoopSym){
            {
                _key := node.FP.Get_key(_env)
                if !Lns_IsNil( _key ) {
                    key := _key.(*Ast_SymbolInfo)
                    if key.FP.Get_name(_env) != "_"{
                        self.FP.Write(_env, _env.LuaVM.String_format("_%s", []LnsAny{key.FP.Get_name(_env)}))
                    } else { 
                        self.FP.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{key.FP.Get_name(_env)}))
                    }
                    
                    self.FP.Write(_env, ", ")
                } else {
                    self.FP.Write(_env, "_, ")
                }
            }
            if valName != "_"{
                self.FP.Write(_env, _env.LuaVM.String_format("_%s", []LnsAny{valName}))
            } else { 
                self.FP.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{valName}))
            }
            
            self.FP.Write(_env, " := ")
        }
        self.FP.Write(_env, "range( ")
        convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.Writeln(_env, ".Items ) {")
        self.FP.PushIndent(_env, nil)
        {
            _key := node.FP.Get_key(_env)
            if !Lns_IsNil( _key ) {
                key := _key.(*Ast_SymbolInfo)
                if Lns_isCondTrue( key.FP.Get_posForModToRef(_env)){
                    self.FP.Writeln(_env, _env.LuaVM.String_format("%s := _%s + 1", []LnsAny{self.FP.getSymbolSym(_env, key), key.FP.Get_name(_env)}))
                }
            }
        }
        if valName != "_"{
            self.FP.Write(_env, _env.LuaVM.String_format("%s := _%s", []LnsAny{valName, valName}))
            self.FP.outputStem2Type(_env, itemType)
            self.FP.Writeln(_env, "")
        }
        
        self.FP.PopIndent(_env)
    } else if _switch20238 == Ast_TypeInfoKind__Map {
        var valName string
        valName = self.FP.getSymbolSym(_env, node.FP.Get_val(_env))
        var itemType *Ast_TypeInfo
        itemType = loopExpType.FP.Get_itemTypeInfoList(_env).GetAt(2).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var keyType *Ast_TypeInfo
        keyType = loopExpType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if Lns_isCondTrue( hasAccessLoopSym){
            {
                _key := node.FP.Get_key(_env)
                if !Lns_IsNil( _key ) {
                    key := _key.(*Ast_SymbolInfo)
                    if key.FP.Get_name(_env) != "_"{
                        self.FP.Write(_env, _env.LuaVM.String_format("_%s", []LnsAny{key.FP.Get_name(_env)}))
                    } else { 
                        self.FP.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{key.FP.Get_name(_env)}))
                    }
                    
                    self.FP.Write(_env, ", ")
                } else {
                    self.FP.Write(_env, "_, ")
                }
            }
            if valName != "_"{
                self.FP.Write(_env, _env.LuaVM.String_format("_%s", []LnsAny{valName}))
            } else { 
                self.FP.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{valName}))
            }
            
            self.FP.Write(_env, " := ")
        }
        self.FP.Write(_env, "range( ")
        convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.Writeln(_env, ".Items ) {")
        self.FP.PushIndent(_env, nil)
        {
            _key := node.FP.Get_key(_env)
            if !Lns_IsNil( _key ) {
                key := _key.(*Ast_SymbolInfo)
                if key.FP.Get_name(_env) != "_"{
                    self.FP.Write(_env, _env.LuaVM.String_format("%s := _%s", []LnsAny{key.FP.Get_name(_env), key.FP.Get_name(_env)}))
                    self.FP.outputStem2Type(_env, keyType)
                    self.FP.Writeln(_env, "")
                }
                
            }
        }
        if valName != "_"{
            self.FP.Write(_env, _env.LuaVM.String_format("%s := _%s", []LnsAny{valName, valName}))
            self.FP.outputStem2Type(_env, itemType)
            self.FP.Writeln(_env, "")
        }
        
        self.FP.PopIndent(_env)
    } else if _switch20238 == Ast_TypeInfoKind__Set {
        var valType *Ast_TypeInfo
        valType = loopExpType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        var valName string
        valName = self.FP.getSymbolSym(_env, node.FP.Get_val(_env))
        if Lns_isCondTrue( hasAccessLoopSym){
            if valName != "_"{
                self.FP.Write(_env, _env.LuaVM.String_format("_%s", []LnsAny{valName}))
            } else { 
                self.FP.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{valName}))
            }
            
            self.FP.Write(_env, " := ")
        }
        self.FP.Write(_env, "range( ")
        convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.Writeln(_env, ".Items ) {")
        self.FP.PushIndent(_env, nil)
        if valName != "_"{
            self.FP.Write(_env, _env.LuaVM.String_format("%s := _%s", []LnsAny{valName, valName}))
            self.FP.outputStem2Type(_env, valType)
            self.FP.Writeln(_env, "")
        }
        
        self.FP.PopIndent(_env)
    } else {
        Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
    }
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Write(_env, "}")
    self.FP.Writeln(_env, "")
}

// 3547: decl @lune.@base.@convGo.convFilter.processForsort
func (self *convGo_convFilter) ProcessForsort(_env *LnsEnv, node *Nodes_ForsortNode,_opt LnsAny) {
    if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_srcTypeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        self.FP.outputForeachLua(_env, &node.Nodes_Node, true, node.FP.Get_exp(_env), node.FP.Get_val(_env), node.FP.Get_key(_env), node.FP.Get_block(_env))
        return 
    }
    var keySym LnsAny
    var valSym LnsAny
    var keyTypeInfo *Ast_TypeInfo
    keyTypeInfo = node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
    if node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Set{
        keySym = node.FP.Get_val(_env)
        
        valSym = node.FP.Get_key(_env)
        
    } else { 
        keySym = node.FP.Get_key(_env)
        
        valSym = node.FP.Get_val(_env)
        
    }
    self.FP.Writeln(_env, "{")
    self.FP.PushIndent(_env, nil)
    var collSym string
    collSym = _env.LuaVM.String_format("__collection%d", []LnsAny{node.FP.Get_id(_env)})
    self.FP.Write(_env, _env.LuaVM.String_format("%s := ", []LnsAny{collSym}))
    convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Writeln(_env, "")
    var sortSym string
    sortSym = _env.LuaVM.String_format("__sorted%d", []LnsAny{node.FP.Get_id(_env)})
    self.FP.Write(_env, _env.LuaVM.String_format("%s := %s.", []LnsAny{sortSym, collSym}))
    if _switch20568 := keyTypeInfo; _switch20568 == Ast_builtinTypeInt {
        self.FP.Writeln(_env, "CreateKeyListInt()")
    } else if _switch20568 == Ast_builtinTypeReal {
        self.FP.Writeln(_env, "CreateKeyListReal()")
    } else if _switch20568 == Ast_builtinTypeString {
        self.FP.Writeln(_env, "CreateKeyListStr()")
    } else {
        self.FP.Writeln(_env, "CreateKeyListStem()")
    }
    self.FP.Writeln(_env, _env.LuaVM.String_format("%s.Sort( %s, nil )", []LnsAny{sortSym, convGo_getLnsItemKind_3923_(_env, keyTypeInfo)}))
    self.FP.Write(_env, "for _, ")
    var key string
    key = _env.LuaVM.String_format("__key%d", []LnsAny{node.FP.Get_id(_env)})
    if keySym != nil{
        keySym_1983 := keySym.(*Ast_SymbolInfo)
        key = _env.LuaVM.String_format("%s", []LnsAny{self.FP.getSymbolSym(_env, keySym_1983)})
        
    }
    self.FP.Write(_env, _env.LuaVM.String_format("_%s", []LnsAny{key}))
    self.FP.Writeln(_env, _env.LuaVM.String_format(" := range( %s.Items ) {", []LnsAny{sortSym}))
    self.FP.PushIndent(_env, nil)
    if valSym != nil{
        valSym_1985 := valSym.(*Ast_SymbolInfo)
        if Lns_isCondTrue( valSym_1985.FP.Get_posForModToRef(_env)){
            self.FP.Write(_env, _env.LuaVM.String_format("%s := %s.Items[ _%s ]", []LnsAny{self.FP.getSymbolSym(_env, valSym_1985), collSym, key}))
            self.FP.outputStem2Type(_env, valSym_1985.FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
    }
    if keySym != nil{
        keySym_1988 := keySym.(*Ast_SymbolInfo)
        if Lns_isCondTrue( keySym_1988.FP.Get_posForModToRef(_env)){
            self.FP.Write(_env, _env.LuaVM.String_format("%s := _%s", []LnsAny{key, key}))
            self.FP.outputStem2Type(_env, keySym_1988.FP.Get_typeInfo(_env))
            self.FP.Writeln(_env, "")
        }
    }
    self.FP.PopIndent(_env)
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "}")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 3625: decl @lune.@base.@convGo.convFilter.processExpUnwrap
func (self *convGo_convFilter) ProcessExpUnwrap(_env *LnsEnv, node *Nodes_ExpUnwrapNode,_opt LnsAny) {
    {
        _def := node.FP.Get_default(_env)
        if !Lns_IsNil( _def ) {
            def := _def.(*Nodes_Node)
            self.FP.Write(_env, "Lns_unwrapDefault( ")
            convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ", ")
            convGo_filter_1292_(_env, def, self, &node.Nodes_Node)
        } else {
            self.FP.Write(_env, "Lns_unwrap( ")
            convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        }
    }
    self.FP.Write(_env, ")")
    self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
}

// 3641: decl @lune.@base.@convGo.convFilter.processExpToDDD
func (self *convGo_convFilter) ProcessExpToDDD(_env *LnsEnv, node *Nodes_ExpToDDDNode,_opt LnsAny) {
    if Lns_isCondTrue( node.FP.Get_expList(_env).FP.Get_mRetExp(_env)){
        convGo_filter_1292_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, &node.Nodes_Node)
    } else { 
        self.FP.Write(_env, "[]LnsAny{ ")
        convGo_filter_1292_(_env, &node.FP.Get_expList(_env).Nodes_Node, self, &node.Nodes_Node)
        self.FP.Write(_env, "}")
    }
}

// 3653: decl @lune.@base.@convGo.convFilter.processExpNew
func (self *convGo_convFilter) ProcessExpNew(_env *LnsEnv, node *Nodes_ExpNewNode,_opt LnsAny) {
    var className string
    className = self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env))
    if self.FP.isSamePackage(_env, node.FP.Get_expType(_env).FP.GetModule(_env)){
        self.FP.Write(_env, _env.LuaVM.String_format("New%s(", []LnsAny{className}))
    } else { 
        {
            _refTypeNode := Nodes_RefTypeNodeDownCastF(node.FP.Get_symbol(_env).FP)
            if !Lns_IsNil( _refTypeNode ) {
                refTypeNode := _refTypeNode.(*Nodes_RefTypeNode)
                {
                    _refNode := Nodes_RefFieldNodeDownCastF(refTypeNode.FP.Get_name(_env).FP)
                    if !Lns_IsNil( _refNode ) {
                        refNode := _refNode.(*Nodes_RefFieldNode)
                        convGo_filter_1292_(_env, refNode.FP.Get_prefix(_env), self, &node.Nodes_Node)
                        self.FP.Write(_env, ".")
                    }
                }
            }
        }
        self.FP.Write(_env, _env.LuaVM.String_format("New%s(", []LnsAny{className}))
    }
    {
        _argList := node.FP.Get_argList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            var scope *Ast_Scope
            scope = Lns_unwrap( node.FP.Get_expType(_env).FP.Get_scope(_env)).(*Ast_Scope)
            var initFuncType *Ast_TypeInfo
            initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField(_env, "__init", true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
            self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), argList), initFuncType.FP.Get_argTypeInfoList(_env), argList, self.option.FP.Get_addEnvArg(_env))
        } else {
            self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, 0, self.option.FP.Get_addEnvArg(_env)))
        }
    }
    self.FP.Write(_env, ")")
}

// 3681: decl @lune.@base.@convGo.convFilter.outputIFMethods
func (self *convGo_convFilter) outputIFMethods(_env *LnsEnv, node *Nodes_DeclClassNode) {
    self.FP.PushIndent(_env, nil)
    var name2MtdType *LnsMap
    name2MtdType = NewLnsMap( map[LnsAny]LnsAny{})
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
        __collection21371 := name2MtdType
        __sorted21371 := __collection21371.CreateKeyListStr()
        __sorted21371.Sort( LnsItemKindStr, nil )
        for _, _name := range( __sorted21371.Items ) {
            typeInfo := __collection21371.Items[ _name ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            name := _name.(string)
            self.FP.Write(_env, _env.LuaVM.String_format("%s(", []LnsAny{self.FP.getSymbol(_env, &convGo_SymbolKind__Func{typeInfo}, name)}))
            if name != "_toMap"{
                self.FP.Write(_env, self.FP.getEnvArgDecl(_env, typeInfo.FP.Get_argTypeInfoList(_env).Len()))
            }
            for _index, _argType := range( typeInfo.FP.Get_argTypeInfoList(_env).Items ) {
                index := _index + 1
                argType := _argType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if index != 1{
                    self.FP.Write(_env, ", ")
                }
                self.FP.Write(_env, _env.LuaVM.String_format("arg%d %s", []LnsAny{index, self.FP.type2gotype(_env, argType)}))
            }
            self.FP.Write(_env, ")")
            self.FP.outputRetType(_env, typeInfo.FP.Get_retTypeInfoList(_env))
            self.FP.Writeln(_env, "")
        }
    }
    self.FP.PopIndent(_env)
}

// 3715: decl @lune.@base.@convGo.convFilter.outputMethodIF
func (self *convGo_convFilter) outputMethodIF(_env *LnsEnv, node *Nodes_DeclClassNode) {
    self.FP.Write(_env, "type ")
    self.FP.Write(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
    self.FP.Writeln(_env, "Mtd interface {")
    self.FP.outputIFMethods(_env, node)
    self.FP.Writeln(_env, "}")
}

// 3725: decl @lune.@base.@convGo.convFilter.outputInterfaceType
func (self *convGo_convFilter) outputInterfaceType(_env *LnsEnv, node *Nodes_DeclClassNode) {
    self.FP.Writeln(_env, _env.LuaVM.String_format("type %s interface {", []LnsAny{self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env))}))
    self.FP.PushIndent(_env, nil)
    self.FP.outputIFMethods(_env, node)
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
    var typeName string
    typeName = self.FP.type2gotype(_env, node.FP.Get_expType(_env))
    self.FP.Writeln(_env, _env.LuaVM.String_format("func Lns_cast2%s( obj LnsAny ) LnsAny {", []LnsAny{typeName}))
    self.FP.Writeln(_env, _env.LuaVM.String_format("    if _, ok := obj.(%s); ok { ", []LnsAny{typeName}))
    self.FP.Writeln(_env, "        return obj")
    self.FP.Writeln(_env, "    }")
    self.FP.Writeln(_env, "    return nil")
    self.FP.Writeln(_env, "}")
}

// 3750: decl @lune.@base.@convGo.convFilter.outputClassType
func (self *convGo_convFilter) outputClassType(_env *LnsEnv, node *Nodes_DeclClassNode) {
    self.FP.Write(_env, "type ")
    self.FP.Write(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
    self.FP.Writeln(_env, " struct {")
    self.FP.PushIndent(_env, nil)
    if node.FP.Get_expType(_env).FP.HasBase(_env){
        var superClassType *Ast_TypeInfo
        superClassType = node.FP.Get_expType(_env).FP.Get_baseTypeInfo(_env)
        self.FP.Writeln(_env, self.FP.getTypeSymbolWithPrefix(_env, superClassType))
    }
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        convGo_filter_1292_(_env, &memberNode.Nodes_Node, self, &node.Nodes_Node)
    }
    self.FP.Write(_env, "FP ")
    self.FP.Write(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
    self.FP.Writeln(_env, "Mtd")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 3779: decl @lune.@base.@convGo.convFilter.outputToStem
func (self *convGo_convFilter) outputToStem(_env *LnsEnv, node *Nodes_DeclClassNode) {
    self.FP.Writeln(_env, _env.LuaVM.String_format("func %s2Stem( obj LnsAny ) LnsAny {", []LnsAny{self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_expType(_env))}))
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, "if obj == nil {")
    self.FP.Writeln(_env, "    return nil")
    self.FP.Writeln(_env, "}")
    self.FP.Writeln(_env, _env.LuaVM.String_format("return obj.(%s).FP", []LnsAny{self.FP.type2gotype(_env, node.FP.Get_expType(_env))}))
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 3792: decl @lune.@base.@convGo.convFilter.outputDownCast
func (self *convGo_convFilter) outputDownCast(_env *LnsEnv, node *Nodes_DeclClassNode) {
    var symbol string
    symbol = self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env))
    self.FP.Write(_env, "type ")
    self.FP.Writeln(_env, _env.LuaVM.String_format("%sDownCast interface {", []LnsAny{symbol}))
    self.FP.PushIndent(_env, nil)
    self.FP.Write(_env, "To")
    self.FP.Write(_env, symbol)
    self.FP.Write(_env, "() *")
    self.FP.Write(_env, symbol)
    self.FP.Writeln(_env, "")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
    self.FP.Writeln(_env, _env.LuaVM.String_format("func %sDownCastF( multi ...LnsAny ) LnsAny {", []LnsAny{symbol}))
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, "if len( multi ) == 0 { return nil }")
    self.FP.Writeln(_env, "obj := multi[ 0 ]")
    self.FP.Writeln(_env, "if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }")
    self.FP.Writeln(_env, _env.LuaVM.String_format("work, ok := obj.(%sDownCast)", []LnsAny{self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_expType(_env))}))
    self.FP.Writeln(_env, _env.LuaVM.String_format("if ok { return work.To%s() }", []LnsAny{symbol}))
    self.FP.Writeln(_env, "return nil")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 3818: decl @lune.@base.@convGo.convFilter.outputCastReceiver
func (self *convGo_convFilter) outputCastReceiver(_env *LnsEnv, node *Nodes_DeclClassNode) {
    self.FP.Write(_env, "func (obj *")
    self.FP.Write(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
    self.FP.Write(_env, ") To")
    self.FP.Write(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
    self.FP.Write(_env, "() *")
    self.FP.Write(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
    self.FP.Writeln(_env, " {")
    self.FP.PushIndent(_env, nil)
    self.FP.Writeln(_env, "return obj")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 3832: decl @lune.@base.@convGo.convFilter.outputNewSetup
func (self *convGo_convFilter) outputNewSetup(_env *LnsEnv, objName string,classType *Ast_TypeInfo) {
    var className string
    className = self.FP.getTypeSymbol(_env, classType)
    self.FP.Writeln(_env, _env.LuaVM.String_format("%s := &%s{}", []LnsAny{objName, className}))
    self.FP.Writeln(_env, _env.LuaVM.String_format("%s.FP = %s", []LnsAny{objName, objName}))
    {
        var workType *Ast_TypeInfo
        workType = classType
        for workType.FP.HasBase(_env) {
            workType = workType.FP.Get_baseTypeInfo(_env)
            
            var superName string
            superName = self.FP.getTypeSymbol(_env, workType)
            self.FP.Writeln(_env, _env.LuaVM.String_format("%s.%s.FP = %s", []LnsAny{objName, superName, objName}))
        }
    }
}

// 3849: decl @lune.@base.@convGo.convFilter.outputConstructor
func (self *convGo_convFilter) outputConstructor(_env *LnsEnv, node *Nodes_DeclClassNode) {
    var scope *Ast_Scope
    scope = Lns_unwrap( node.FP.Get_expType(_env).FP.Get_scope(_env)).(*Ast_Scope)
    var initFuncType *Ast_TypeInfo
    initFuncType = Lns_unwrap( scope.FP.GetTypeInfoField(_env, "__init", true, scope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
    var className string
    className = self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env))
    var ctorName string
    ctorName = self.FP.getConstrSymbol(_env, node.FP.Get_expType(_env))
    if Lns_op_not(node.FP.Get_expType(_env).FP.Get_abstractFlag(_env)){
        self.FP.Write(_env, _env.LuaVM.String_format("func New%s(", []LnsAny{className}))
        self.FP.outputDeclFuncArg(_env, initFuncType)
        self.FP.Writeln(_env, _env.LuaVM.String_format(") *%s {", []LnsAny{className}))
        self.FP.PushIndent(_env, nil)
        self.FP.outputNewSetup(_env, "obj", node.FP.Get_expType(_env))
        self.FP.Write(_env, _env.LuaVM.String_format("obj.%s(", []LnsAny{ctorName}))
        self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, initFuncType.FP.Get_argTypeInfoList(_env).Len(), self.option.FP.Get_addEnvArg(_env)))
        for _index, _ := range( initFuncType.FP.Get_argTypeInfoList(_env).Items ) {
            index := _index + 1
            if index != 1{
                self.FP.Write(_env, ", ")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("arg%d", []LnsAny{index}))
        }
        self.FP.Writeln(_env, ")")
        self.FP.Writeln(_env, "return obj")
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    }
    if Lns_op_not(node.FP.HasUserInit(_env)){
        self.FP.Write(_env, _env.LuaVM.String_format("func (self *%s) %s(", []LnsAny{className, ctorName}))
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
            self.FP.Write(_env, _env.LuaVM.String_format("self.%s.%s( ", []LnsAny{superName, self.FP.getConstrSymbol(_env, superType)}))
            self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, baseInitFuncType.FP.Get_argTypeInfoList(_env).Len(), self.option.FP.Get_addEnvArg(_env)))
            {
                var _from22612 LnsInt = 1
                var _to22612 LnsInt = baseInitFuncType.FP.Get_argTypeInfoList(_env).Len()
                for _work22612 := _from22612; _work22612 <= _to22612; _work22612++ {
                    index := _work22612
                    if index != 1{
                        self.FP.Write(_env, ",")
                    }
                    if node.FP.Get_hasOldCtor(_env){
                        self.FP.Write(_env, "nil")
                    } else { 
                        self.FP.Write(_env, _env.LuaVM.String_format("arg%d", []LnsAny{index}))
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
        var argIndex LnsInt
        argIndex = superArgNum + 1
        for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
                self.FP.Writeln(_env, _env.LuaVM.String_format("self.%s = arg%d", []LnsAny{self.FP.getSymbol(_env, &convGo_SymbolKind__Member{Ast_isPubToExternal(_env, memberNode.FP.Get_accessMode(_env))}, memberNode.FP.Get_name(_env).Txt), argIndex}))
                argIndex = argIndex + 1
                
            }
        }
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    }
}

// 3940: decl @lune.@base.@convGo.convFilter.outputAccessor
func (self *convGo_convFilter) OutputAccessor(_env *LnsEnv, node *Nodes_DeclClassNode) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType(_env)
    if classType.FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
        return 
    }
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        var memberNameToken *Types_Token
        memberNameToken = memberNode.FP.Get_name(_env)
        var memberName string
        memberName = memberNameToken.Txt
        var memberSym *Ast_SymbolInfo
        memberSym = memberNode.FP.Get_symbolInfo(_env)
        if memberNode.FP.Get_getterMode(_env) != Ast_AccessMode__None{
            var getterName string
            getterName = _env.LuaVM.String_format("get_%s", []LnsAny{memberName})
            var getterSym *Ast_SymbolInfo
            getterSym = Lns_unwrap( _env.NilAccFin(_env.NilAccPush(classType.FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(_env, getterName)})/* 3954:33 */)).(*Ast_SymbolInfo)
            self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__Type{getterSym.FP.Get_typeInfo(_env)})
            var retType *Ast_TypeInfo
            retType = getterSym.FP.Get_typeInfo(_env).FP.Get_retTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo().FP.Get_srcTypeInfo(_env)
            self.FP.Write(_env, "{ return ")
            var prefix string
            if memberSym.FP.Get_staticFlag(_env){
                prefix = ""
                
            } else { 
                prefix = "self."
                
            }
            if retType.FP.Get_srcTypeInfo(_env) != memberSym.FP.Get_typeInfo(_env).FP.Get_srcTypeInfo(_env){
                if retType.FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
                    if Ast_isClass(_env, memberSym.FP.Get_typeInfo(_env).FP.Get_srcTypeInfo(_env)){
                        self.FP.Write(_env, _env.LuaVM.String_format("%s%s.FP", []LnsAny{prefix, self.FP.getSymbolSym(_env, memberSym)}))
                    }
                } else if Ast_isClass(_env, retType){
                    self.FP.Write(_env, _env.LuaVM.String_format("&%s%s.%s", []LnsAny{prefix, self.FP.getSymbolSym(_env, memberSym), self.FP.getTypeSymbol(_env, retType)}))
                } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( retType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                    _env.SetStackVal( retType.FP.HasBase(_env)) ).(bool)){
                    self.FP.Write(_env, _env.LuaVM.String_format("%s%s.FP", []LnsAny{prefix, self.FP.getSymbolSym(_env, memberSym)}))
                    self.FP.outputStem2Type(_env, retType)
                } else { 
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( convGo_isAnyType_1296_(_env, retType)) &&
                        _env.SetStackVal( Ast_isClass(_env, memberSym.FP.Get_typeInfo(_env))) ).(bool)){
                        self.FP.Write(_env, _env.LuaVM.String_format("%s%s.FP", []LnsAny{prefix, self.FP.getSymbolSym(_env, memberSym)}))
                    } else { 
                        Util_err(_env, "not support")
                    }
                }
            } else { 
                self.FP.Write(_env, _env.LuaVM.String_format("%s%s", []LnsAny{prefix, self.FP.getSymbolSym(_env, memberSym)}))
            }
            self.FP.Writeln(_env, " }")
        }
        if memberNode.FP.Get_setterMode(_env) != Ast_AccessMode__None{
            var setterName string
            setterName = _env.LuaVM.String_format("set_%s", []LnsAny{memberName})
            var setterSym *Ast_SymbolInfo
            setterSym = Lns_unwrap( _env.NilAccFin(_env.NilAccPush(classType.FP.Get_scope(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_Scope).FP.GetSymbolInfoChild(_env, setterName)})/* 3995:33 */)).(*Ast_SymbolInfo)
            self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__Type{setterSym.FP.Get_typeInfo(_env)})
            self.FP.Write(_env, _env.LuaVM.String_format("{ self.%s = arg1 ", []LnsAny{self.FP.getSymbolSym(_env, memberSym)}))
            self.FP.Writeln(_env, "}")
        }
    }
}

// 4004: decl @lune.@base.@convGo.convFilter.outputStaticMember
func (self *convGo_convFilter) outputStaticMember(_env *LnsEnv, node *Nodes_DeclClassNode) {
    if self.processMode == convGo_ProcessMode__DeclClass{
        for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
            memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
            if memberNode.FP.Get_staticFlag(_env){
                self.FP.Writeln(_env, _env.LuaVM.String_format("var %s %s", []LnsAny{self.FP.getSymbol(_env, &convGo_SymbolKind__Static{node.FP.Get_expType(_env)}, memberNode.FP.Get_name(_env).Txt), self.FP.type2gotype(_env, memberNode.FP.Get_expType(_env))}))
            }
        }
        {
            _initBlock := node.FP.Get_initBlock(_env).FP.Get_func(_env)
            if !Lns_IsNil( _initBlock ) {
                initBlock := _initBlock.(*Nodes_DeclMethodNode)
                convGo_filter_1292_(_env, &initBlock.Nodes_Node, self, &node.Nodes_Node)
                self.FP.Writeln(_env, "")
            }
        }
    } else { 
        {
            _initBlock := node.FP.Get_initBlock(_env).FP.Get_func(_env)
            if !Lns_IsNil( _initBlock ) {
                initBlock := _initBlock.(*Nodes_DeclMethodNode)
                self.FP.Writeln(_env, _env.LuaVM.String_format("%s(%s)", []LnsAny{self.FP.getFuncSymbol(_env, initBlock.FP.Get_expType(_env)), convGo_getAddEnvArg_1279_(_env, 0, self.option.FP.Get_addEnvArg(_env))}))
            }
        }
    }
}

// 4036: decl @lune.@base.@convGo.convFilter.getFromStemName
func (self *convGo_convFilter) getFromStemName(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    __func__ := "@lune.@base.@convGo.convFilter.getFromStemName"
    var workTypeInfo *Ast_TypeInfo
    workTypeInfo = convGo_getOrgTypeInfo_1660_(_env, typeInfo)
    {
        _name := convGo_type2FromStemNameMap.Get(workTypeInfo)
        if !Lns_IsNil( _name ) {
            name := _name.(string)
            return name
        }
    }
    if _switch23410 := workTypeInfo.FP.Get_kind(_env); _switch23410 == Ast_TypeInfoKind__List || _switch23410 == Ast_TypeInfoKind__Array {
        return "Lns_ToList"
    } else if _switch23410 == Ast_TypeInfoKind__Set {
        return "Lns_ToSet"
    } else if _switch23410 == Ast_TypeInfoKind__Map {
        return "Lns_ToLnsMap"
    } else if _switch23410 == Ast_TypeInfoKind__Class {
        return _env.LuaVM.String_format("%s_FromMap", []LnsAny{self.FP.getTypeSymbol(_env, workTypeInfo)})
    } else {
        Util_err(_env, _env.LuaVM.String_format("%s: not support -- %s", []LnsAny{__func__, Ast_TypeInfoKind_getTxt( workTypeInfo.FP.Get_kind(_env))}))
    }
// insert a dummy
    return ""
}


// 4063: decl @lune.@base.@convGo.convFilter.outputConvItemType
func (self *convGo_convFilter) outputConvItemType(_env *LnsEnv, typeInfo *Ast_TypeInfo,alt2type LnsAny) {
    var workTypeInfo *Ast_TypeInfo
    workTypeInfo = typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)
    if typeInfo.FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
        if alt2type != nil{
            alt2type_2199 := alt2type.(*LnsMap)
            {
                _alt := alt2type_2199.Get(workTypeInfo)
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
            self.FP.Write(_env, _env.LuaVM.String_format("paramList[%d]", []LnsAny{altType.FP.Get_altIndex(_env) - 1}))
        } else {
            self.FP.Writeln(_env, "Lns_ToObjParam{")
            self.FP.PushIndent(_env, nil)
            self.FP.Write(_env, _env.LuaVM.String_format("%sSub, %s,", []LnsAny{self.FP.getFromStemName(_env, workTypeInfo), typeInfo.FP.Get_nilable(_env)}))
            self.FP.outputConvItemTypeList(_env, workTypeInfo.FP.Get_itemTypeInfoList(_env), alt2type)
            self.FP.PopIndent(_env)
            self.FP.Write(_env, "}")
        }
    }
}

// 4088: decl @lune.@base.@convGo.convFilter.outputConvItemTypeList
func (self *convGo_convFilter) outputConvItemTypeList(_env *LnsEnv, itemTypeInfoList *LnsList,alt2type LnsAny) {
    if itemTypeInfoList.Len() > 0{
        self.FP.Write(_env, "[]Lns_ToObjParam{")
        self.FP.PushIndent(_env, nil)
        for _index, _itemType := range( itemTypeInfoList.Items ) {
            index := _index + 1
            itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if index > 1{
                self.FP.Write(_env, ",")
            }
            self.FP.outputConvItemType(_env, itemType, alt2type)
        }
        self.FP.PopIndent(_env)
        self.FP.Write(_env, "}")
    } else { 
        self.FP.Write(_env, "nil")
    }
}

// 4108: decl @lune.@base.@convGo.convFilter.outputAlter2MapFunc
func (self *convGo_convFilter) outputAlter2MapFunc(_env *LnsEnv, alt2Map *LnsMap) {
    __func__ := "@lune.@base.@convGo.convFilter.outputAlter2MapFunc"
    self.FP.Write(_env, "{")
    for _altType, _assinType := range( alt2Map.Items ) {
        altType := _altType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        assinType := _assinType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
        if altType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
            if assinType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                Util_err(_env, _env.LuaVM.String_format("not support: %s", []LnsAny{__func__}))
            } else { 
                self.FP.outputConvItemType(_env, assinType, alt2Map)
            }
        }
    }
    self.FP.Write(_env, "}")
}

// 4127: decl @lune.@base.@convGo.convFilter.outputAsyncItem
func (self *convGo_convFilter) outputAsyncItem(_env *LnsEnv, node *Nodes_DeclClassNode) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType(_env)
    var classScope *Ast_Scope
    classScope = Lns_unwrap( classType.FP.Get_scope(_env)).(*Ast_Scope)
    var createSym *Ast_SymbolInfo
    createSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild(_env, "_createPipe")).(*Ast_SymbolInfo)
    self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__WithClass{classType, createSym.FP.Get_typeInfo(_env)})
    self.FP.Writeln(_env, "{")
    self.FP.Writeln(_env, "   return NewLnspipe( arg1 )")
    self.FP.Writeln(_env, "}")
}

// 4138: decl @lune.@base.@convGo.convFilter.outputMapping
func (self *convGo_convFilter) outputMapping(_env *LnsEnv, node *Nodes_DeclClassNode) {
    var classType *Ast_TypeInfo
    classType = node.FP.Get_expType(_env)
    var className string
    className = self.FP.getTypeSymbol(_env, classType)
    self.FP.Writeln(_env, _env.LuaVM.String_format("func (self *%s) ToMapSetup( obj *LnsMap ) *LnsMap {", []LnsAny{className}))
    self.FP.PushIndent(_env, nil)
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
            self.FP.Writeln(_env, _env.LuaVM.String_format("obj.Items[\"%s\"] = Lns_ToCollection( self.%s )", []LnsAny{memberNode.FP.Get_symbolInfo(_env).FP.Get_name(_env), self.FP.getSymbolSym(_env, memberNode.FP.Get_symbolInfo(_env))}))
        }
    }
    if classType.FP.HasBase(_env){
        self.FP.Writeln(_env, _env.LuaVM.String_format("return self.%s.ToMapSetup( obj )", []LnsAny{self.FP.getTypeSymbol(_env, classType.FP.Get_baseTypeInfo(_env))}))
    } else { 
        self.FP.Writeln(_env, "return obj")
    }
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
    self.FP.Writeln(_env, _env.LuaVM.String_format("func (self *%s) ToMap() *LnsMap {", []LnsAny{className}))
    self.FP.Writeln(_env, "    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )")
    self.FP.Writeln(_env, "}")
    var fromStemName string
    fromStemName = self.FP.getFromStemName(_env, classType)
    var classScope *Ast_Scope
    classScope = Lns_unwrap( classType.FP.Get_scope(_env)).(*Ast_Scope)
    if Lns_op_not(classType.FP.Get_abstractFlag(_env)){
        var envStr string
        envStr = convGo_getAddEnvArg_1279_(_env, 1, self.option.FP.Get_addEnvArg(_env))
        {
            var fromMapSym *Ast_SymbolInfo
            fromMapSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild(_env, "_fromMap")).(*Ast_SymbolInfo)
            self.FP.Writeln(_env, _env.LuaVM.String_format("func %s(%s arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){", []LnsAny{self.FP.getSymbolSym(_env, fromMapSym), envStr}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("   return %s( arg1, paramList )", []LnsAny{fromStemName}))
            self.FP.Writeln(_env, "}")
        }
        {
            var fromStemSym *Ast_SymbolInfo
            fromStemSym = Lns_unwrap( classScope.FP.GetSymbolInfoChild(_env, "_fromStem")).(*Ast_SymbolInfo)
            self.FP.Writeln(_env, _env.LuaVM.String_format("func %s(%s arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){", []LnsAny{self.FP.getSymbolSym(_env, fromStemSym), envStr}))
            self.FP.Writeln(_env, _env.LuaVM.String_format("   return %s( arg1, paramList )", []LnsAny{fromStemName}))
            self.FP.Writeln(_env, "}")
        }
        self.FP.Writeln(_env, _env.LuaVM.String_format("func %s( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {", []LnsAny{fromStemName}))
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, _env.LuaVM.String_format("_,conv,mess := %sSub(obj,false, paramList);", []LnsAny{fromStemName}))
        self.FP.Writeln(_env, "return conv,mess")
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
        self.FP.Writeln(_env, _env.LuaVM.String_format("func %sSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {", []LnsAny{fromStemName}))
        self.FP.PushIndent(_env, nil)
        self.FP.Writeln(_env, "var objMap *LnsMap")
        self.FP.Writeln(_env, "if work, ok := obj.(*LnsMap); !ok {")
        self.FP.Writeln(_env, "   return false, nil, \"no map -- \" + Lns_ToString(obj)")
        self.FP.Writeln(_env, "} else {")
        self.FP.Writeln(_env, "   objMap = work")
        self.FP.Writeln(_env, "}")
        self.FP.outputNewSetup(_env, "newObj", classType)
        self.FP.Writeln(_env, _env.LuaVM.String_format("return %sMain( newObj, objMap, paramList )", []LnsAny{fromStemName}))
        self.FP.PopIndent(_env)
        self.FP.Writeln(_env, "}")
    }
    self.FP.Writeln(_env, _env.LuaVM.String_format("func %sMain( newObj %s, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {", []LnsAny{fromStemName, self.FP.type2gotype(_env, classType)}))
    self.FP.PushIndent(_env, nil)
    if classType.FP.HasBase(_env){
        self.FP.Writeln(_env, _env.LuaVM.String_format("if ok,_,mess := %sMain( &newObj.%s, objMap, paramList ); !ok {", []LnsAny{self.FP.getFromStemName(_env, classType.FP.Get_baseTypeInfo(_env)), self.FP.getTypeSymbol(_env, classType.FP.Get_baseTypeInfo(_env))}))
        self.FP.Writeln(_env, "    return false,nil,mess")
        self.FP.Writeln(_env, "}")
    }
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode.(Nodes_DeclMemberNodeDownCast).ToNodes_DeclMemberNode()
        if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
            var memberName string
            memberName = self.FP.getSymbolSym(_env, memberNode.FP.Get_symbolInfo(_env))
            self.FP.Write(_env, "if ok,conv,mess := ")
            if memberNode.FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                for _index, _itemType := range( classType.FP.Get_itemTypeInfoList(_env).Items ) {
                    index := _index + 1
                    itemType := _itemType.(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                    if itemType == memberNode.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env){
                        self.FP.Write(_env, _env.LuaVM.String_format("paramList[%d].Func( objMap.Items[\"%s\"], %s, paramList[%d].Child", []LnsAny{index - 1, memberNode.FP.Get_symbolInfo(_env).FP.Get_name(_env), memberNode.FP.Get_expType(_env).FP.Get_nilable(_env), index - 1}))
                    }
                }
            } else { 
                self.FP.Write(_env, _env.LuaVM.String_format("%sSub( objMap.Items[\"%s\"], %s, ", []LnsAny{self.FP.getFromStemName(_env, memberNode.FP.Get_expType(_env)), memberNode.FP.Get_symbolInfo(_env).FP.Get_name(_env), memberNode.FP.Get_expType(_env).FP.Get_nilable(_env)}))
                self.FP.outputConvItemTypeList(_env, memberNode.FP.Get_expType(_env).FP.Get_itemTypeInfoList(_env), nil)
            }
            self.FP.Writeln(_env, "); !ok {")
            self.FP.Writeln(_env, _env.LuaVM.String_format("   return false,nil,\"%s:\" + mess.(string)", []LnsAny{memberNode.FP.Get_symbolInfo(_env).FP.Get_name(_env)}))
            self.FP.Writeln(_env, "} else {")
            self.FP.Write(_env, _env.LuaVM.String_format("   newObj.%s = conv", []LnsAny{memberName}))
            self.FP.outputAny2Type(_env, memberNode.FP.Get_expType(_env))
            self.FP.Writeln(_env, "")
            self.FP.Writeln(_env, "}")
        }
    }
    self.FP.Writeln(_env, "return true, newObj, nil")
    self.FP.PopIndent(_env)
    self.FP.Writeln(_env, "}")
}

// 4266: decl @lune.@base.@convGo.convFilter.outputDummyAbstractMethod
func (self *convGo_convFilter) outputDummyAbstractMethod(_env *LnsEnv, classType *Ast_TypeInfo,methodType *Ast_TypeInfo) {
    self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__WithClass{classType, methodType})
    self.FP.Writeln(_env, "{")
    self.FP.outputDummyReturn(_env, methodType.FP.Get_retTypeInfoList(_env))
    self.FP.Writeln(_env, "}")
}

// 4278: decl @lune.@base.@convGo.convFilter.outputDummyAbstractMethodOfClass
func (self *convGo_convFilter) outputDummyAbstractMethodOfClass(_env *LnsEnv, classType *Ast_TypeInfo) {
    var name2typeMap *LnsMap
    name2typeMap = NewLnsMap( map[LnsAny]LnsAny{})
    var scope *Ast_Scope
    scope = Lns_unwrap( classType.FP.Get_scope(_env)).(*Ast_Scope)
    scope.FP.FilterSymbolInfoIfField(_env, scope, Ast_ScopeAccess__Normal, Ast_filterForm(func(_env *LnsEnv, symbolInfo *Ast_SymbolInfo) bool {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( symbolInfo.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) &&
            _env.SetStackVal( symbolInfo.FP.Get_typeInfo(_env).FP.Get_abstractFlag(_env)) ).(bool)){
            {
                _methodType := scope.FP.GetTypeInfoChild(_env, symbolInfo.FP.Get_name(_env))
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
        __collection24770 := name2typeMap
        __sorted24770 := __collection24770.CreateKeyListStr()
        __sorted24770.Sort( LnsItemKindStr, nil )
        for _, ___key24770 := range( __sorted24770.Items ) {
            methodType := __collection24770.Items[ ___key24770 ].(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            self.FP.outputDummyAbstractMethod(_env, classType, methodType)
        }
    }
}

// 4308: decl @lune.@base.@convGo.convFilter.outputAdvertise
func (self *convGo_convFilter) outputAdvertise(_env *LnsEnv, node *Nodes_DeclClassNode) {
    __func__ := "@lune.@base.@convGo.convFilter.outputAdvertise"
    var methodNameSet *LnsSet
    methodNameSet = node.FP.CreateMethodNameSetWithoutAdv(_env)
    for _, _adv := range( node.FP.Get_advertiseList(_env).Items ) {
        adv := _adv.(Nodes_AdvertiseInfoDownCast).ToNodes_AdvertiseInfo()
        if adv.FP.Get_prefix(_env) != ""{
            Util_err(_env, _env.LuaVM.String_format("%s: not support advertise with prefix", []LnsAny{__func__}))
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
                        self.FP.Writeln(_env, _env.LuaVM.String_format("// advertise -- %d", []LnsAny{node.FP.Get_pos(_env).LineNo}))
                        self.FP.OutputDeclFunc(_env, self.option.FP.Get_addEnvArg(_env), &convGo_FuncInfo__WithClass{node.FP.Get_expType(_env), funcType})
                        self.FP.Writeln(_env, " {")
                        if funcType.FP.Get_retTypeInfoList(_env).Len() > 0{
                            self.FP.Write(_env, "    return ")
                        }
                        self.FP.Write(_env, _env.LuaVM.String_format("self.%s. ", []LnsAny{self.FP.getSymbolSym(_env, adv.FP.Get_member(_env).FP.Get_symbolInfo(_env))}))
                        if adv.FP.Get_member(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Class{
                            self.FP.Write(_env, "FP.")
                        }
                        self.FP.Write(_env, _env.LuaVM.String_format("%s( ", []LnsAny{self.FP.getSymbolSym(_env, symbol)}))
                        self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, funcType.FP.Get_argTypeInfoList(_env).Len(), self.option.FP.Get_addEnvArg(_env)))
                        for _index, _ := range( funcType.FP.Get_argTypeInfoList(_env).Items ) {
                            index := _index + 1
                            if index > 1{
                                self.FP.Write(_env, ",")
                            }
                            self.FP.Write(_env, _env.LuaVM.String_format("arg%d", []LnsAny{index}))
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

// 4353: decl @lune.@base.@convGo.convFilter.processProtoClass
func (self *convGo_convFilter) ProcessProtoClass(_env *LnsEnv, node *Nodes_ProtoClassNode,_opt LnsAny) {
}

// 4358: decl @lune.@base.@convGo.convFilter.processDeclClass
func (self *convGo_convFilter) ProcessDeclClass(_env *LnsEnv, node *Nodes_DeclClassNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processDeclClass"
    if node.FP.IsModule(_env){
        return 
    }
    if self.processMode == convGo_ProcessMode__DeclClass{
        if _switch25339 := node.FP.Get_expType(_env).FP.Get_kind(_env); _switch25339 == Ast_TypeInfoKind__Class {
            self.FP.Writeln(_env, _env.LuaVM.String_format("// declaration Class -- %s", []LnsAny{node.FP.Get_expType(_env).FP.Get_rawTxt(_env)}))
            self.FP.outputStaticMember(_env, node)
            self.FP.outputMethodIF(_env, node)
            self.FP.outputClassType(_env, node)
            self.FP.outputToStem(_env, node)
            self.FP.outputDownCast(_env, node)
            self.FP.outputCastReceiver(_env, node)
            self.FP.outputConstructor(_env, node)
            self.FP.OutputAccessor(_env, node)
            self.FP.outputDummyAbstractMethodOfClass(_env, node.FP.Get_expType(_env))
            self.FP.outputAdvertise(_env, node)
            if node.FP.Get_expType(_env).FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeMapping, nil){
                self.FP.outputMapping(_env, node)
            }
            if node.FP.Get_expType(_env).FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeAsyncItem, nil){
                self.FP.outputAsyncItem(_env, node)
            }
            for _, _fieldNode := range( node.FP.Get_fieldList(_env).Items ) {
                fieldNode := _fieldNode.(Nodes_NodeDownCast).ToNodes_Node()
                if Lns_isCondTrue( Nodes_DeclMemberNodeDownCastF(fieldNode.FP)){
                } else { 
                    convGo_filter_1292_(_env, fieldNode, self, &node.Nodes_Node)
                    self.FP.Writeln(_env, "")
                }
            }
        } else if _switch25339 == Ast_TypeInfoKind__IF {
            self.FP.outputInterfaceType(_env, node)
        } else {
            Util_err(_env, _env.LuaVM.String_format("%s: not support -- %s:%d", []LnsAny{__func__, Ast_TypeInfoKind_getTxt( node.FP.Get_expType(_env).FP.Get_kind(_env)), node.FP.Get_pos(_env).LineNo}))
        }
    } else { 
        self.FP.outputStaticMember(_env, node)
    }
}

// 4445: decl @lune.@base.@convGo.convFilter.outputCallPrefix
func (self *convGo_convFilter) OutputCallPrefix(_env *LnsEnv, threading bool,callId LnsInt,node *Nodes_Node,prefixNode LnsAny,funcSymbol *Ast_SymbolInfo)(bool, LnsAny) {
    var funcType *Ast_TypeInfo
    funcType = funcSymbol.FP.Get_typeInfo(_env)
    var nilAccName string
    nilAccName = _env.LuaVM.String_format("%s_%d", []LnsAny{Lns_car(_env.LuaVM.String_gsub(self.moduleTypeInfo.FP.Get_rawTxt(_env),"@", "")).(string), callId})
    var callKind LnsAny
    callKind = convGo_CallKind__Normal_Obj
    var extCallFlag bool
    if prefixNode != nil{
        prefixNode_2361 := prefixNode.(*Nodes_Node)
        extCallFlag = prefixNode_2361.FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext
        
    } else {
        extCallFlag = false
        
    }
    if extCallFlag{
        callKind = convGo_CallKind__LuaCall_Obj
        
    }
    var getEnvTxt string
    getEnvTxt = self.env.FP.getEnv(_env, threading)
    var processNilAcc func(_env *LnsEnv, workPrefixNode *Nodes_Node)
    processNilAcc = func(_env *LnsEnv, workPrefixNode *Nodes_Node) {
        if Lns_op_not(node.FP.HasNilAccess(_env)){
            return 
        }
        var retNum LnsInt
        retNum = funcType.FP.Get_retTypeInfoList(_env).Len()
        if _switch25622 := retNum; _switch25622 == 0 {
            self.FP.Write(_env, _env.LuaVM.String_format("Lns_NilAccCall0( %s, func () {", []LnsAny{getEnvTxt}))
        } else if _switch25622 == 1 {
            self.FP.Write(_env, _env.LuaVM.String_format("Lns_NilAccCall1( %s, func () LnsAny { return ", []LnsAny{getEnvTxt}))
        } else {
            if retNum <= convGo_MaxNilAccNum{
                var anys string
                anys = "LnsAny"
                {
                    var _from25557 LnsInt = 2
                    var _to25557 LnsInt = retNum
                    for _work25557 := _from25557; _work25557 <= _to25557; _work25557++ {
                        anys = _env.LuaVM.String_format("%s,LnsAny", []LnsAny{anys})
                        
                    }
                }
                self.FP.Write(_env, _env.LuaVM.String_format("Lns_NilAccCall%d( %s, func () (%s) { return ", []LnsAny{retNum, getEnvTxt, anys}))
            } else { 
                var args string
                args = "LnsAny"
                {
                    var _from25600 LnsInt = 2
                    var _to25600 LnsInt = retNum
                    for _work25600 := _from25600; _work25600 <= _to25600; _work25600++ {
                        args = _env.LuaVM.String_format("%s,LnsAny", []LnsAny{args})
                        
                    }
                }
                self.FP.Write(_env, _env.LuaVM.String_format("lns_NilAccCall_%s( %s, func () (%s) { return ", []LnsAny{nilAccName, getEnvTxt, args}))
            }
        }
        if extCallFlag{
            if funcType.FP.Get_retTypeInfoList(_env).Len() > 1{
                self.FP.Write(_env, _env.LuaVM.String_format("Lns_callExt%d( ", []LnsAny{node.FP.Get_id(_env)}))
            }
        }
        self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccPop().(%s)", []LnsAny{getEnvTxt, self.FP.type2gotype(_env, workPrefixNode.FP.Get_expType(_env).FP.Get_nonnilableType(_env))}))
    }
    var closeParen bool
    closeParen = false
    if prefixNode != nil{
        prefixNode_2386 := prefixNode.(*Nodes_Node)
        if node.FP.HasNilAccess(_env){
            if funcType.FP.Get_retTypeInfoList(_env).Len() >= 2{
                if funcType.FP.Get_retTypeInfoList(_env).Len() <= convGo_MaxNilAccNum{
                    self.FP.Write(_env, _env.LuaVM.String_format("Lns_NilAccFinCall%d( ", []LnsAny{funcType.FP.Get_retTypeInfoList(_env).Len()}))
                } else { 
                    self.FP.Write(_env, _env.LuaVM.String_format("lns_NilAccFinCall_%s(", []LnsAny{nilAccName}))
                }
                closeParen = true
                
            }
        }
        var prefixType *Ast_TypeInfo
        prefixType = prefixNode_2386.FP.Get_expType(_env).FP.Get_nonnilableType(_env)
        if prefixType == Ast_builtinTypeString{
            if node.FP.HasNilAccess(_env){
                Util_err(_env, "not support nilAccName")
            }
            {
                _runtime := self.FP.getVM(_env, threading, funcType)
                if !Lns_IsNil( _runtime ) {
                    runtime := _runtime.(string)
                    callKind = &convGo_CallKind__RuntimeCall{prefixNode_2386}
                    
                    self.FP.Write(_env, runtime)
                }
            }
        } else { 
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( Lns_op_not(funcType.FP.Get_staticFlag(_env))) ||
                _env.SetStackVal( funcSymbol.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) ).(bool){
                if node.FP.HasNilAccess(_env){
                    if Lns_op_not(prefixNode_2386.FP.HasNilAccess(_env)){
                        self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccFin(", []LnsAny{getEnvTxt}))
                        self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccPush(", []LnsAny{getEnvTxt}))
                        convGo_filter_1292_(_env, prefixNode_2386, self, node)
                        self.FP.Writeln(_env, ") && ")
                    } else { 
                        convGo_filter_1292_(_env, prefixNode_2386, self, node)
                        self.FP.Writeln(_env, "&&")
                    }
                } else { 
                    if extCallFlag{
                        if funcType.FP.Get_retTypeInfoList(_env).Len() > 1{
                            self.FP.Write(_env, _env.LuaVM.String_format("Lns_callExt%d( ", []LnsAny{node.FP.Get_id(_env)}))
                        }
                    }
                    convGo_filter_1292_(_env, prefixNode_2386, self, node)
                }
            } else { 
            }
            processNilAcc(_env, prefixNode_2386)
            if prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                self.FP.Write(_env, _env.LuaVM.String_format(".CallMethod( \"%s\", Lns_2DDD", []LnsAny{funcSymbol.FP.Get_name(_env)}))
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
                    if _switch26129 := prefixKind; _switch26129 == Ast_TypeInfoKind__Class {
                        self.FP.Write(_env, _env.LuaVM.String_format(".FP.%s", []LnsAny{self.FP.getSymbolSym(_env, funcSymbol)}))
                    } else {
                        var builtinFuncs *Builtin_BuiltinFuncType
                        builtinFuncs = TransUnit_getBuiltinFunc(_env)
                        {
                            _runtime := self.FP.getVM(_env, threading, funcType)
                            if !Lns_IsNil( _runtime ) {
                                runtime := _runtime.(string)
                                self.FP.Write(_env, runtime)
                                callKind = convGo_CallKind__BuiltinCall_Obj
                                
                            } else {
                                if _switch26106 := funcType; _switch26106 == builtinFuncs.List_sort || _switch26106 == builtinFuncs.Array_sort {
                                    callKind = &convGo_CallKind__SortCall{prefixType.FP.Get_itemTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()}
                                    
                                }
                                self.FP.Write(_env, _env.LuaVM.String_format(".%s", []LnsAny{self.FP.getSymbolSym(_env, funcSymbol)}))
                            }
                        }
                    }
                } else { 
                    if _switch26230 := funcType.FP.Get_kind(_env); _switch26230 == Ast_TypeInfoKind__Method {
                        if _switch26198 := prefixKind; _switch26198 == Ast_TypeInfoKind__Class {
                            self.FP.Write(_env, _env.LuaVM.String_format(".FP.%s", []LnsAny{self.FP.getSymbolSym(_env, funcSymbol)}))
                            if funcSymbol.FP.Get_name(_env) == "_toMap"{
                                callKind = convGo_CallKind__BuiltinCall_Obj
                                
                            }
                        } else {
                            self.FP.Write(_env, _env.LuaVM.String_format(".%s", []LnsAny{self.FP.getSymbolSym(_env, funcSymbol)}))
                        }
                    } else {
                        if funcSymbol.FP.Get_kind(_env) == Ast_SymbolKind__Mbr{
                            self.FP.Write(_env, ".")
                        }
                        self.FP.Write(_env, self.FP.getSymbolSym(_env, funcSymbol))
                    }
                }
            }
        }
    }
    return closeParen, callKind
}

// 4626: decl @lune.@base.@convGo.convFilter.processExpCall
func (self *convGo_convFilter) ProcessExpCall(_env *LnsEnv, node *Nodes_ExpCallNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpCall"
    opt := _opt.(*ConvGo_Opt)
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_nonnilableType(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Method) &&
        _env.SetStackVal( funcType.FP.Get_parentInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) &&
        _env.SetStackVal( funcType.FP.Get_rawTxt(_env) == "get__txt") ).(bool)){
        self.FP.Write(_env, _env.LuaVM.String_format("%s(", []LnsAny{self.FP.getEnumGetTxtSym(_env, Lns_unwrap( Ast_EnumTypeInfoDownCastF(funcType.FP.Get_parentInfo(_env).FP.Get_aliasSrc(_env).FP)).(*Ast_EnumTypeInfo))}))
        var fieldNode *Nodes_RefFieldNode
        
        {
            _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func(_env).FP)
            if _fieldNode == nil{
                Util_err(_env, "not support -- __func__")
            } else {
                fieldNode = _fieldNode.(*Nodes_RefFieldNode)
            }
        }
        convGo_filter_1292_(_env, fieldNode.FP.Get_prefix(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, ")")
        return 
    }
    var retGenerics bool
    if opt.Parent.FP.Get_kind(_env) == Nodes_NodeKind_get_StmtExp(_env){
        retGenerics = false
        
    } else { 
        retGenerics = convGo_isRetGenerics_2055_(_env, node)
        
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( retGenerics) &&
            _env.SetStackVal( funcType.FP.Get_retTypeInfoList(_env).Len() != 1) ).(bool)){
            self.FP.Write(_env, _env.LuaVM.String_format("%s(", []LnsAny{self.FP.getConvGenericsName(_env, &node.Nodes_Node)}))
        }
    }
    var addEnvArg bool
    addEnvArg = self.option.FP.Get_addEnvArg(_env)
    var closeParen bool
    closeParen = false
    var callKind LnsAny
    callKind = convGo_CallKind__Normal_Obj
    var withPrefix bool
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func(_env).FP)
        if !Lns_IsNil( _fieldNode ) {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            if _switch26474 := fieldNode.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_kind(_env); _switch26474 == Ast_TypeInfoKind__List || _switch26474 == Ast_TypeInfoKind__Map || _switch26474 == Ast_TypeInfoKind__Set || _switch26474 == Ast_TypeInfoKind__Array {
                addEnvArg = false
                
            }
            withPrefix = true
            
            closeParen, callKind = self.FP.OutputCallPrefix(_env, node.FP.IsThreading(_env), node.FP.Get_id(_env), &fieldNode.Nodes_Node, fieldNode.FP.Get_prefix(_env), Lns_unwrap( fieldNode.FP.Get_symbolInfo(_env)).(*Ast_SymbolInfo))
            
        } else {
            withPrefix = false
            
            if Ast_isBuiltin(_env, funcType.FP.Get_typeId(_env).Id){
                {
                    _runtime := self.FP.getVM(_env, node.FP.IsThreading(_env), funcType)
                    if !Lns_IsNil( _runtime ) {
                        runtime := _runtime.(string)
                        self.FP.Write(_env, runtime)
                        addEnvArg = false
                        
                    } else {
                        if _switch26617 := funcType.FP.Get_srcTypeInfo(_env); _switch26617 == Ast_builtinTypeForm {
                            convGo_filter_1292_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
                            callKind = convGo_CallKind__FormCall_Obj
                            
                        } else {
                            Util_err(_env, _env.LuaVM.String_format("%s: not support -- %s:%d", []LnsAny{__func__, funcType.FP.GetTxt(_env, nil, nil, nil), node.FP.Get_pos(_env).LineNo}))
                        }
                    }
                }
            } else { 
                if funcType.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                    self.FP.Write(_env, _env.LuaVM.String_format("%s.RunLoadedfunc", []LnsAny{self.env.FP.getLuavm(_env, node.FP.IsThreading(_env))}))
                    callKind = convGo_CallKind__RunLoaded_Obj
                    
                } else { 
                    convGo_filter_1292_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
                }
            }
        }
    }
    self.FP.Write(_env, "(")
    var skipArg bool
    skipArg = false
    var closeTxt LnsAny
    closeTxt = nil
    switch _exp26969 := callKind.(type) {
    case *convGo_CallKind__RuntimeCall:
    prefixNode := _exp26969.Val1
        convGo_filter_1292_(_env, prefixNode, self, node.FP.Get_func(_env))
        if Lns_isCondTrue( node.FP.Get_argList(_env)){
            self.FP.Write(_env, ",")
        }
        addEnvArg = false
        
    case *convGo_CallKind__FormCall:
        self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, 1, addEnvArg))
        self.FP.Write(_env, "Lns_2DDD(")
        addEnvArg = false
        
    case *convGo_CallKind__RunLoaded:
        convGo_filter_1292_(_env, node.FP.Get_func(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, ",")
        if Lns_op_not(node.FP.Get_argList(_env)){
            self.FP.Write(_env, "[]LnsAny{}")
        } else { 
            self.FP.Write(_env, "Lns_2DDD(")
            closeTxt = ")"
            
        }
        addEnvArg = false
        
    case *convGo_CallKind__SortCall:
    typeInfo := _exp26969.Val1
        self.FP.Write(_env, _env.LuaVM.String_format("%s, ", []LnsAny{convGo_getLnsItemKind_3923_(_env, typeInfo)}))
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                if argList.FP.Get_expType(_env).FP.Get_argTypeInfoList(_env).Len() == 2{
                    skipArg = true
                    
                    self.FP.Write(_env, "LnsComp(func ( val1, val2 LnsAny ) bool {")
                    self.FP.Write(_env, "return ")
                    self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), argList), funcType.FP.Get_argTypeInfoList(_env), argList, false)
                    var argType string
                    argType = self.FP.type2gotype(_env, argList.FP.Get_expType(_env).FP.Get_argTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
                    self.FP.Write(_env, _env.LuaVM.String_format("( val1.(%s), val2.(%s) )", []LnsAny{argType, argType}))
                    self.FP.Write(_env, "})")
                }
            }
        }
    case *convGo_CallKind__BuiltinCall:
        addEnvArg = false
        
    case *convGo_CallKind__LuaCall:
        closeTxt = ")"
        
    }
    if Lns_op_not(skipArg){
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), argList), funcType.FP.Get_argTypeInfoList(_env), argList, addEnvArg)
            } else {
                self.FP.Write(_env, convGo_getAddEnvArg_1279_(_env, 0, addEnvArg))
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
        self.FP.Write(_env, ",")
        self.FP.outputConvItemTypeList(_env, funcType.FP.Get_parentInfo(_env).FP.Get_itemTypeInfoList(_env), fieldNode.FP.Get_prefix(_env).FP.Get_expType(_env).FP.CreateAlt2typeMap(_env, false))
    }
    if closeTxt != nil{
        closeTxt_2484 := closeTxt.(string)
        self.FP.Write(_env, closeTxt_2484)
    }
    self.FP.Write(_env, ")")
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( callKind == convGo_CallKind__LuaCall_Obj) ||
        _env.SetStackVal( callKind == convGo_CallKind__RunLoaded_Obj) ).(bool){
        if funcType.FP.Get_retTypeInfoList(_env).Len() == 1{
            if opt.Parent.FP.Get_kind(_env) != Nodes_NodeKind_get_StmtExp(_env){
                self.FP.Write(_env, "[0]")
                var retTypeList *LnsList
                retTypeList = Lns_unwrap( Lns_car(Ast_convToExtTypeList(_env, self.processInfo, funcType.FP.Get_retTypeInfoList(_env)))).(*LnsList)
                self.FP.outputAny2Type(_env, retTypeList.GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo())
            }
        } else if funcType.FP.Get_retTypeInfoList(_env).Len() > 1{
            self.FP.Write(_env, ")")
        }
    }
    if retGenerics{
        if funcType.FP.Get_retTypeInfoList(_env).Len() == 1{
            var retType *Ast_TypeInfo
            retType = funcType.FP.Get_retTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
            if convGo_isAnyType_1296_(_env, retType){
                if Lns_op_not(convGo_isAnyType_1296_(_env, node.FP.Get_expType(_env))){
                    self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
                }
            } else if retType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
                if retType.FP.Get_srcTypeInfo(_env) != node.FP.Get_expType(_env).FP.Get_srcTypeInfo(_env){
                    self.FP.Write(_env, ".FP")
                    self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
                }
            }
        } else { 
            self.FP.Write(_env, ")")
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( withPrefix) &&
        _env.SetStackVal( node.FP.HasNilAccess(_env)) ).(bool)){
        self.FP.Write(_env, "})")
        self.FP.Write(_env, _env.LuaVM.String_format("/* %d:%d */", []LnsAny{node.FP.Get_pos(_env).LineNo, node.FP.Get_pos(_env).Column}))
        if opt.Parent.FP.HasNilAccess(_env){
        } else { 
            self.FP.Write(_env, ")")
        }
        if closeParen{
            self.FP.Write(_env, ")")
        }
    }
    if callKind == convGo_CallKind__FormCall_Obj{
        self.FP.Write(_env, ")")
    }
}

// 4845: decl @lune.@base.@convGo.convFilter.processExpMRet
func (self *convGo_convFilter) ProcessExpMRet(_env *LnsEnv, node *Nodes_ExpMRetNode,_opt LnsAny) {
    convGo_filter_1292_(_env, node.FP.Get_mRet(_env), self, &node.Nodes_Node)
}

// 4852: decl @lune.@base.@convGo.convFilter.processExpAccessMRet
func (self *convGo_convFilter) ProcessExpAccessMRet(_env *LnsEnv, node *Nodes_ExpAccessMRetNode,_opt LnsAny) {
}

// 4860: decl @lune.@base.@convGo.convFilter.processExpList
func (self *convGo_convFilter) ProcessExpList(_env *LnsEnv, node *Nodes_ExpListNode,_opt LnsAny) {
    for _index, _exp := range( node.FP.Get_expList(_env).Items ) {
        index := _index + 1
        exp := _exp.(Nodes_NodeDownCast).ToNodes_Node()
        if index != 1{
            self.FP.Write(_env, ", ")
        }
        {
            _mRetExp := node.FP.Get_mRetExp(_env)
            if !Lns_IsNil( _mRetExp ) {
                mRetExp := _mRetExp.(*Nodes_MRetExp)
                if mRetExp.FP.Get_index(_env) == index{
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( index == 1) ||
                        _env.SetStackVal( exp.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool){
                        convGo_filter_1292_(_env, exp, self, &node.Nodes_Node)
                    } else { 
                        self.FP.Write(_env, "Lns_2DDD(")
                        convGo_filter_1292_(_env, exp, self, &node.Nodes_Node)
                        self.FP.Write(_env, ")")
                    }
                    break
                }
            }
        }
        convGo_filter_1292_(_env, exp, self, &node.Nodes_Node)
    }
}

// 4884: decl @lune.@base.@convGo.convFilter.processExpOp1
func (self *convGo_convFilter) ProcessExpOp1(_env *LnsEnv, node *Nodes_ExpOp1Node,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpOp1"
    if _switch27805 := node.FP.Get_op(_env).Txt; _switch27805 == "~" {
        self.FP.Write(_env, "^")
        convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    } else if _switch27805 == "+" || _switch27805 == "-" {
        self.FP.Write(_env, node.FP.Get_op(_env).Txt)
        convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    } else if _switch27805 == "not" {
        self.FP.Write(_env, "Lns_op_not(")
        convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, ")")
    } else if _switch27805 == "#" {
        if _switch27784 := node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env); _switch27784 == Ast_TypeInfoKind__List {
            convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ".Len()")
        } else if _switch27784 == Ast_TypeInfoKind__Ext {
            if _switch27751 := node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_extedType(_env).FP.Get_kind(_env); _switch27751 == Ast_TypeInfoKind__List {
                convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
                self.FP.Write(_env, ".Len()")
            } else {
                Util_err(_env, _env.LuaVM.String_format("%s: not support -- %s", []LnsAny{__func__, node.FP.Get_exp(_env).FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil)}))
            }
        } else {
            self.FP.Write(_env, "len(")
            convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ")")
        }
    } else {
        Util_err(_env, _env.LuaVM.String_format("%s: not support -- %s", []LnsAny{__func__, node.FP.Get_op(_env).Txt}))
    }
}

// 4932: decl @lune.@base.@convGo.convFilter.processExpMultiTo1
func (self *convGo_convFilter) ProcessExpMultiTo1(_env *LnsEnv, node *Nodes_ExpMultiTo1Node,_opt LnsAny) {
    self.FP.Write(_env, "Lns_car(")
    convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, ")")
    self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
}

// 4944: decl @lune.@base.@convGo.convFilter.processExpCast
func (self *convGo_convFilter) ProcessExpCast(_env *LnsEnv, node *Nodes_ExpCastNode,_opt LnsAny) {
    if _switch28212 := node.FP.Get_castKind(_env); _switch28212 == Nodes_CastKind__Force {
        if convGo_isAnyType_1296_(_env, node.FP.Get_exp(_env).FP.Get_expType(_env)){
            if _switch27978 := node.FP.Get_castType(_env); _switch27978 == Ast_builtinTypeInt {
                self.FP.Write(_env, "Lns_forceCastInt(")
                convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
                self.FP.Write(_env, ")")
            } else if _switch27978 == Ast_builtinTypeReal {
                self.FP.Write(_env, "Lns_forceCastReal(")
                convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
                self.FP.Write(_env, ")")
            } else {
                convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
                self.FP.outputAny2Type(_env, node.FP.Get_castType(_env))
            }
        } else { 
            self.FP.Write(_env, _env.LuaVM.String_format("(%s)(", []LnsAny{self.FP.type2gotype(_env, node.FP.Get_castType(_env))}))
            convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ")")
        }
    } else if _switch28212 == Nodes_CastKind__Implicit {
        if node.FP.Get_exp(_env).FP.Get_expTypeList(_env).Len() > 1{
            convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        } else { 
            self.FP.outputImplicitCast(_env, node.FP.Get_castType(_env), node.FP.Get_exp(_env), node)
        }
    } else if _switch28212 == Nodes_CastKind__Normal {
        var typeName string
        typeName = self.FP.getTypeSymbolWithPrefix(_env, node.FP.Get_castType(_env))
        var castType *Ast_TypeInfo
        castType = node.FP.Get_castType(_env).FP.Get_nonnilableType(_env)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( castType.FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
            _env.SetStackVal( castType != Ast_builtinTypeString) ).(bool)){
            self.FP.Write(_env, _env.LuaVM.String_format("%sDownCastF(", []LnsAny{typeName}))
            convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( node.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
                _env.SetStackVal( node.FP.Get_exp(_env).FP.Get_expType(_env) != Ast_builtinTypeString) ).(bool)){
                self.FP.Write(_env, ".FP")
            }
            self.FP.Write(_env, ")")
        } else { 
            self.FP.Write(_env, _env.LuaVM.String_format("Lns_cast2%s( ", []LnsAny{self.FP.type2gotype(_env, castType)}))
            convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ")")
        }
    }
}

// 5004: decl @lune.@base.@convGo.convFilter.processExpParen
func (self *convGo_convFilter) ProcessExpParen(_env *LnsEnv, node *Nodes_ExpParenNode,_opt LnsAny) {
    if node.FP.Get_exp(_env).FP.Get_expTypeList(_env).Len() >= 2{
        self.FP.Write(_env, "Lns_car(")
        convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, ")")
        self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
    } else { 
        self.FP.Write(_env, "(")
        convGo_filter_1292_(_env, node.FP.Get_exp(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, ")")
    }
}

// 5021: decl @lune.@base.@convGo.convFilter.processExpSetVal
func (self *convGo_convFilter) ProcessExpSetVal(_env *LnsEnv, node *Nodes_ExpSetValNode,_opt LnsAny) {
    convGo_filter_1292_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
    if convGo_getExpListKind_1885_(_env, node.FP.Get_exp1(_env).FP.Get_expTypeList(_env), node.FP.Get_exp2(_env), self.option.FP.Get_addEnvArg(_env)) == convGo_ExpListKind__Direct{
        {
            var _from28373 LnsInt = node.FP.Get_exp1(_env).FP.Get_expTypeList(_env).Len() + 1
            var _to28373 LnsInt = node.FP.Get_exp2(_env).FP.Get_expTypeList(_env).Len()
            for _work28373 := _from28373; _work28373 <= _to28373; _work28373++ {
                self.FP.Write(_env, ",_")
            }
        }
    }
    self.FP.Write(_env, " = ")
    self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), node.FP.Get_exp2(_env)), node.FP.Get_exp1(_env).FP.Get_expTypeList(_env), node.FP.Get_exp2(_env), false)
    self.FP.Writeln(_env, "")
}

// 5039: decl @lune.@base.@convGo.convFilter.processExpSetItem
func (self *convGo_convFilter) ProcessExpSetItem(_env *LnsEnv, node *Nodes_ExpSetItemNode,_opt LnsAny) {
    convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, ".Set(")
    switch _exp28485 := node.FP.Get_index(_env).(type) {
    case *Nodes_IndexVal__NodeIdx:
    index := _exp28485.Val1
        convGo_filter_1292_(_env, index, self, &node.Nodes_Node)
    case *Nodes_IndexVal__SymIdx:
    index := _exp28485.Val1
        self.FP.Write(_env, _env.LuaVM.String_format("\"%s\"", []LnsAny{index}))
    }
    self.FP.Write(_env, ",")
    convGo_filter_1292_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
    self.FP.Write(_env, ")")
}

// 5057: decl @lune.@base.@convGo.convFilter.processAndOr
func (self *convGo_convFilter) processAndOr(_env *LnsEnv, node *Nodes_ExpOp2Node,opTxt string,parent *Nodes_Node) {
    var getEnvTxt string
    getEnvTxt = self.env.FP.getEnv(_env, node.FP.IsThreading(_env))
    var firstFlag bool
    firstFlag = Lns_op_not(convFilter_processAndOr__isAndOr_4990_(_env, parent))
    if firstFlag{
        self.FP.Writeln(_env, _env.LuaVM.String_format("%s.PopVal( %s.IncStack() ||", []LnsAny{getEnvTxt, getEnvTxt}))
        self.FP.PushIndent(_env, nil)
    }
    var opCC string
    if opTxt == "and"{
        opCC = "&&"
        
    } else { 
        opCC = "||"
        
    }
    if convFilter_processAndOr__isAndOr_4990_(_env, node.FP.Get_exp1(_env)){
        convGo_filter_1292_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
    } else { 
        self.FP.Write(_env, _env.LuaVM.String_format("%s.SetStackVal( ", []LnsAny{getEnvTxt}))
        convGo_filter_1292_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, ") ")
    }
    self.FP.Writeln(_env, opCC)
    if convFilter_processAndOr__isAndOr_4990_(_env, node.FP.Get_exp2(_env)){
        convGo_filter_1292_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
    } else { 
        self.FP.Write(_env, _env.LuaVM.String_format("%s.SetStackVal( ", []LnsAny{getEnvTxt}))
        convGo_filter_1292_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, ") ")
    }
    if firstFlag{
        self.FP.Write(_env, ")")
        self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
        self.FP.PopIndent(_env)
    }
}

// 5118: decl @lune.@base.@convGo.convFilter.processExpOp2
func (self *convGo_convFilter) ProcessExpOp2(_env *LnsEnv, node *Nodes_ExpOp2Node,_opt LnsAny) {
    opt := _opt.(*ConvGo_Opt)
    var opTxt string
    opTxt = node.FP.Get_op(_env).Txt
    if _switch29039 := opTxt; _switch29039 == "and" || _switch29039 == "or" {
        self.FP.processAndOr(_env, node, opTxt, opt.Parent)
    } else if _switch29039 == ".." {
        convGo_filter_1292_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, " + ")
        convGo_filter_1292_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
    } else {
        {
            __exp := Ast_bitBinOpMap.Get(opTxt)
            if !Lns_IsNil( __exp ) {
                _exp := __exp.(LnsInt)
                if _switch28926 := _exp; _switch28926 == Ast_BitOpKind__LShift {
                    opTxt = "<<"
                    
                } else if _switch28926 == Ast_BitOpKind__RShift {
                    opTxt = ">>"
                    
                }
                convGo_filter_1292_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
                self.FP.Write(_env, " " + opTxt + " ")
                convGo_filter_1292_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
            } else {
                convGo_filter_1292_(_env, node.FP.Get_exp1(_env), self, &node.Nodes_Node)
                {
                    _op := convGo_op2map.Get(opTxt)
                    if !Lns_IsNil( _op ) {
                        op := _op.(string)
                        self.FP.Write(_env, _env.LuaVM.String_format(" %s ", []LnsAny{op}))
                    } else {
                        self.FP.Write(_env, _env.LuaVM.String_format(" %s ", []LnsAny{opTxt}))
                    }
                }
                convGo_filter_1292_(_env, node.FP.Get_exp2(_env), self, &node.Nodes_Node)
            }
        }
    }
}

// 5163: decl @lune.@base.@convGo.convFilter.processExpRef
func (self *convGo_convFilter) ProcessExpRef(_env *LnsEnv, node *Nodes_ExpRefNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpRef"
    if node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
        self.FP.Write(_env, "ddd")
    } else { 
        if Lns_isCondTrue( node.FP.Get_symbolInfo(_env).FP.Get_convModuleParam(_env)){
            self.FP.Write(_env, _env.LuaVM.String_format("%s_%d", []LnsAny{node.FP.Get_symbolInfo(_env).FP.Get_name(_env), node.FP.Get_symbolInfo(_env).FP.Get_symbolId(_env)}))
        } else { 
            if _switch29251 := node.FP.Get_symbolInfo(_env).FP.Get_kind(_env); _switch29251 == Ast_SymbolKind__Var || _switch29251 == Ast_SymbolKind__Arg {
                self.FP.Write(_env, self.FP.getSymbolSym(_env, node.FP.Get_symbolInfo(_env)))
            } else if _switch29251 == Ast_SymbolKind__Fun {
                if Ast_isBuiltin(_env, node.FP.Get_expType(_env).FP.Get_typeId(_env).Id){
                    var builtinFunc *Builtin_BuiltinFuncType
                    builtinFunc = TransUnit_getBuiltinFunc(_env)
                    if _switch29173 := node.FP.Get_expType(_env); _switch29173 == builtinFunc.Lns_print {
                        self.FP.Write(_env, "Lns_print")
                    } else {
                        Util_err(_env, _env.LuaVM.String_format("%s: not support -- %s", []LnsAny{__func__, node.FP.Get_symbolInfo(_env).FP.Get_name(_env)}))
                    }
                } else { 
                    self.FP.Write(_env, self.FP.getSymbol(_env, &convGo_SymbolKind__Func{node.FP.Get_expType(_env)}, node.FP.Get_symbolInfo(_env).FP.Get_name(_env)))
                }
            } else if _switch29251 == Ast_SymbolKind__Typ {
                if node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Module{
                    self.FP.Write(_env, node.FP.Get_symbolInfo(_env).FP.Get_name(_env))
                } else { 
                    self.FP.Write(_env, self.FP.getTypeSymbol(_env, node.FP.Get_expType(_env)))
                }
            } else {
                self.FP.Write(_env, node.FP.Get_symbolInfo(_env).FP.Get_name(_env))
            }
        }
    }
}

// 5211: decl @lune.@base.@convGo.convFilter.processExpRefItem
func (self *convGo_convFilter) ProcessExpRefItem(_env *LnsEnv, node *Nodes_ExpRefItemNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processExpRefItem"
    var getEnvTxt string
    getEnvTxt = self.env.FP.getEnv(_env, node.FP.IsThreading(_env))
    var prefixType *Ast_TypeInfo
    prefixType = node.FP.Get_val(_env).FP.Get_expType(_env).FP.Get_nonnilableType(_env)
    if _switch30048 := prefixType.FP.Get_kind(_env); _switch30048 == Ast_TypeInfoKind__Ext {
        if node.FP.Get_nilAccess(_env){
            self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccFin( %s.NilAccPush( ", []LnsAny{getEnvTxt, getEnvTxt}))
            convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ") && ")
            self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccPush( %s.NilAccPop().(*Lns_luaValue)", []LnsAny{getEnvTxt, getEnvTxt}))
        } else { 
            convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            if prefixType.FP.Get_extedType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Stem{
                self.FP.Write(_env, ".(*Lns_luaValue)")
            }
        }
        self.FP.Write(_env, ".GetAt(")
        {
            _index := node.FP.Get_index(_env)
            if !Lns_IsNil( _index ) {
                index := _index.(*Nodes_Node)
                convGo_filter_1292_(_env, index, self, &node.Nodes_Node)
            } else {
                self.FP.Write(_env, _env.LuaVM.String_format("\"%s\"", []LnsAny{convGo_str2gostr_1644_(_env, Lns_unwrap( node.FP.Get_symbol(_env)).(string))}))
            }
        }
        self.FP.Write(_env, ")")
        self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
    } else if _switch30048 == Ast_TypeInfoKind__List || _switch30048 == Ast_TypeInfoKind__Array {
        if node.FP.Get_nilAccess(_env){
            if Lns_op_not(node.FP.Get_val(_env).FP.HasNilAccess(_env)){
                self.FP.Writeln(_env, _env.LuaVM.String_format("%s.NilAccFin( %s.NilAccPush(", []LnsAny{getEnvTxt, getEnvTxt}))
                convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
                self.FP.Writeln(_env, ") && ")
            } else { 
                convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
                self.FP.Writeln(_env, "&&")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccPush( %s.NilAccPop().(*LnsList)", []LnsAny{getEnvTxt, getEnvTxt}))
            self.FP.Write(_env, ".GetAt(")
            convGo_filter_1292_(_env, Lns_unwrap( node.FP.Get_index(_env)).(*Nodes_Node), self, &node.Nodes_Node)
            self.FP.Write(_env, ")")
            self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
            self.FP.Write(_env, "))")
        } else { 
            convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ".GetAt(")
            convGo_filter_1292_(_env, Lns_unwrap( node.FP.Get_index(_env)).(*Nodes_Node), self, &node.Nodes_Node)
            self.FP.Write(_env, ")")
            self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
        }
    } else if _switch30048 == Ast_TypeInfoKind__Map {
        if node.FP.Get_nilAccess(_env){
            if Lns_op_not(node.FP.Get_val(_env).FP.HasNilAccess(_env)){
                self.FP.Writeln(_env, _env.LuaVM.String_format("%s.NilAccFin( %s.NilAccPush(", []LnsAny{getEnvTxt, getEnvTxt}))
                convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
                self.FP.Writeln(_env, ") && ")
            } else { 
                convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
                self.FP.Writeln(_env, "&&")
            }
            self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccPush( %s.NilAccPop().(*LnsMap)", []LnsAny{getEnvTxt, getEnvTxt}))
        } else { 
            convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            if prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Stem{
                self.FP.Write(_env, ".(*LnsMap)")
            }
        }
        self.FP.Write(_env, ".Get(")
        {
            _index := node.FP.Get_index(_env)
            if !Lns_IsNil( _index ) {
                index := _index.(*Nodes_Node)
                convGo_filter_1292_(_env, index, self, &node.Nodes_Node)
            } else {
                self.FP.Write(_env, _env.LuaVM.String_format("\"%s\"", []LnsAny{convGo_str2gostr_1644_(_env, Lns_unwrap( node.FP.Get_symbol(_env)).(string))}))
            }
        }
        self.FP.Write(_env, ")")
        self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
        if node.FP.Get_nilAccess(_env){
            self.FP.Write(_env, "))")
        }
    } else if _switch30048 == Ast_TypeInfoKind__Stem {
        self.FP.Write(_env, "Lns_FromStemGetAt(")
        convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
        self.FP.Write(_env, ",")
        {
            _index := node.FP.Get_index(_env)
            if !Lns_IsNil( _index ) {
                index := _index.(*Nodes_Node)
                convGo_filter_1292_(_env, index, self, &node.Nodes_Node)
            } else {
                self.FP.Write(_env, _env.LuaVM.String_format("\"%s\"", []LnsAny{convGo_str2gostr_1644_(_env, Lns_unwrap( node.FP.Get_symbol(_env)).(string))}))
            }
        }
        self.FP.Write(_env, _env.LuaVM.String_format(", %s )", []LnsAny{node.FP.Get_nilAccess(_env)}))
        self.FP.outputStem2Type(_env, node.FP.Get_expType(_env))
    } else {
        if prefixType == Ast_builtinTypeString{
            self.FP.Write(_env, "LnsInt(")
            convGo_filter_1292_(_env, node.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, "[")
            convGo_filter_1292_(_env, Lns_unwrap( node.FP.Get_index(_env)).(*Nodes_Node), self, &node.Nodes_Node)
            self.FP.Write(_env, "-1])")
        } else { 
            Util_err(_env, _env.LuaVM.String_format("%s: not support -- %d, %s", []LnsAny{__func__, node.FP.Get_pos(_env).LineNo, Ast_TypeInfoKind_getTxt( prefixType.FP.Get_kind(_env))}))
        }
    }
}

// 5322: decl @lune.@base.@convGo.convFilter.processRefField
func (self *convGo_convFilter) ProcessRefField(_env *LnsEnv, node *Nodes_RefFieldNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processRefField"
    opt := _opt.(*ConvGo_Opt)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) &&
        _env.SetStackVal( node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) ).(bool)){
        self.FP.Write(_env, self.FP.getSymbol(_env, &convGo_SymbolKind__Static{node.FP.Get_expType(_env)}, node.FP.Get_field(_env).Txt))
        return 
    }
    {
        _symbol := node.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( _symbol ) {
            symbol := _symbol.(*Ast_SymbolInfo)
            var orgSym *Ast_SymbolInfo
            orgSym = symbol.FP.GetOrg(_env)
            var builtinFuncs *Builtin_BuiltinFuncType
            builtinFuncs = TransUnit_getBuiltinFunc(_env)
            if builtinFuncs.FP.Get_allSymbolSet(_env).Has(Ast_SymbolInfo2Stem(orgSym)){
                self.FP.Write(_env, _env.LuaVM.String_format("Lns_%s_%s", []LnsAny{Lns_car(_env.LuaVM.String_gsub(node.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_rawTxt(_env),"@", "")).(string), orgSym.FP.Get_name(_env)}))
                return 
            }
            if symbol.FP.Get_staticFlag(_env){
                self.FP.Write(_env, self.FP.getSymbolSym(_env, symbol))
                return 
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(symbol.FP.Get_scope(_env).FP.Get_ownerTypeInfo(_env)) && 
                Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Ast_TypeInfo).FP.Get_kind(_env)})) == Ast_TypeInfoKind__Module) &&
                _env.SetStackVal( symbol.FP.Get_kind(_env) == Ast_SymbolKind__Var) ).(bool)){
                self.FP.Write(_env, self.FP.getSymbolSym(_env, symbol))
                return 
            }
        }
    }
    var getEnvTxt string
    getEnvTxt = self.env.FP.getEnv(_env, node.FP.IsThreading(_env))
    var openParenNum LnsInt
    if Lns_op_not(node.FP.HasNilAccess(_env)){
        openParenNum = 0
        
        if Lns_op_not(node.FP.Get_prefix(_env).FP.HasNilAccess(_env)){
            convGo_filter_1292_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
        } else { 
            Util_err(_env, _env.LuaVM.String_format("%s: not support", []LnsAny{__func__}))
        }
    } else { 
        if Lns_op_not(node.FP.Get_prefix(_env).FP.HasNilAccess(_env)){
            self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccFin(", []LnsAny{getEnvTxt}))
            self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccPush(", []LnsAny{getEnvTxt}))
            convGo_filter_1292_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
            self.FP.Writeln(_env, ") && ")
        } else { 
            convGo_filter_1292_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
            self.FP.Writeln(_env, "&&")
        }
        self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccPush(", []LnsAny{getEnvTxt}))
        if opt.Parent.FP.HasNilAccess(_env){
            openParenNum = 1
            
        } else { 
            openParenNum = 2
            
        }
        self.FP.Write(_env, _env.LuaVM.String_format("%s.NilAccPop().(%s)", []LnsAny{getEnvTxt, self.FP.type2gotype(_env, node.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_nonnilableType(_env))}))
    }
    self.FP.Write(_env, ".")
    {
        _symbol := node.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( _symbol ) {
            symbol := _symbol.(*Ast_SymbolInfo)
            if node.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
                self.FP.Write(_env, _env.LuaVM.String_format("GetAt( \"%s\" )", []LnsAny{symbol.FP.Get_name(_env)}))
                self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
            } else { 
                self.FP.Write(_env, self.FP.getSymbolSym(_env, symbol))
                var orgSym *Ast_SymbolInfo
                orgSym = symbol.FP.GetOrg(_env)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( orgSym.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) &&
                    _env.SetStackVal( orgSym.FP.Get_typeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                    _env.SetStackVal( convGo_isAnyType_1296_(_env, orgSym.FP.Get_typeInfo(_env))) ).(bool)){
                    self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
                }
            }
        } else {
            Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
        }
    }
    {
        var _from30556 LnsInt = 1
        var _to30556 LnsInt = openParenNum
        for _work30556 := _from30556; _work30556 <= _to30556; _work30556++ {
            self.FP.Write(_env, ")")
        }
    }
}

// 5407: decl @lune.@base.@convGo.convFilter.processExpOmitEnum
func (self *convGo_convFilter) ProcessExpOmitEnum(_env *LnsEnv, node *Nodes_ExpOmitEnumNode,_opt LnsAny) {
    self.FP.Write(_env, self.FP.getSymbol(_env, &convGo_SymbolKind__Static{node.FP.Get_expType(_env).FP.Get_aliasSrc(_env)}, node.FP.Get_valInfo(_env).FP.Get_name(_env)))
}

// 5413: decl @lune.@base.@convGo.convFilter.processGetField
func (self *convGo_convFilter) ProcessGetField(_env *LnsEnv, node *Nodes_GetFieldNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processGetField"
    opt := _opt.(*ConvGo_Opt)
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
                        self.FP.Write(_env, _env.LuaVM.String_format("%s( ", []LnsAny{self.FP.getEnumGetTxtSym(_env, enumType)}))
                        convGo_filter_1292_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
                        self.FP.Write(_env, ")")
                        return 
                    }
                }
                if Lns_isCondTrue( Ast_AlgeTypeInfoDownCastF(symbolInfo.FP.Get_namespaceTypeInfo(_env).FP)){
                    convGo_filter_1292_(_env, node.FP.Get_prefix(_env), self, &node.Nodes_Node)
                    self.FP.Write(_env, ".(LnsAlgeVal).GetTxt()")
                    return 
                }
            }
            if symbolInfo.FP.Get_staticFlag(_env){
                self.FP.Write(_env, self.FP.getSymbolSym(_env, symbolInfo))
                self.FP.Write(_env, _env.LuaVM.String_format("(%s)", []LnsAny{convGo_getAddEnvArg_1279_(_env, 0, self.option.FP.Get_addEnvArg(_env))}))
            } else { 
                var closeParen bool
                closeParen = convGo_convExp30769(Lns_2DDD(self.FP.OutputCallPrefix(_env, node.FP.IsThreading(_env), node.FP.Get_id(_env), &node.Nodes_Node, node.FP.Get_prefix(_env), symbolInfo)))
                self.FP.Write(_env, _env.LuaVM.String_format("(%s)", []LnsAny{convGo_getAddEnvArg_1279_(_env, 0, self.option.FP.Get_addEnvArg(_env))}))
                var retType *Ast_TypeInfo
                retType = symbolInfo.FP.Get_typeInfo(_env).FP.Get_retTypeInfoList(_env).GetAt(1).(Ast_TypeInfoDownCast).ToAst_TypeInfo()
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( retType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
                    _env.SetStackVal( Lns_op_not(retType.FP.HasBase(_env))) ).(bool)){
                    self.FP.outputAny2Type(_env, node.FP.Get_expType(_env))
                }
                if node.FP.HasNilAccess(_env){
                    self.FP.Write(_env, "})")
                    if opt.Parent.FP.HasNilAccess(_env){
                    } else { 
                        self.FP.Write(_env, ")")
                    }
                }
                if closeParen{
                    self.FP.Write(_env, ")")
                }
            }
        } else {
            Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
        }
    }
}

// 5462: decl @lune.@base.@convGo.convFilter.processReturn
func (self *convGo_convFilter) ProcessReturn(_env *LnsEnv, node *Nodes_ReturnNode,_opt LnsAny) {
    self.FP.Write(_env, "return ")
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            convGo_filter_1292_(_env, &expList.Nodes_Node, self, &node.Nodes_Node)
        }
    }
    self.FP.Writeln(_env, "")
}

// 5472: decl @lune.@base.@convGo.convFilter.processTestCase
func (self *convGo_convFilter) ProcessTestCase(_env *LnsEnv, node *Nodes_TestCaseNode,_opt LnsAny) {
    if Lns_op_not(self.enableTest){
        return 
    }
    self.FP.Writeln(_env, _env.LuaVM.String_format("func lns_test_%s_%s( %s *Testing_Ctrl ) {", []LnsAny{self.FP.getModuleName(_env, self.moduleTypeInfo, false), node.FP.Get_name(_env).Txt, node.FP.Get_ctrlName(_env)}))
    convGo_filter_1292_(_env, &node.FP.Get_block(_env).Nodes_Node, self, &node.Nodes_Node)
    self.FP.Writeln(_env, "}")
}

// 5487: decl @lune.@base.@convGo.convFilter.processTestBlock
func (self *convGo_convFilter) ProcessTestBlock(_env *LnsEnv, node *Nodes_TestBlockNode,_opt LnsAny) {
    var stmtList *LnsList
    stmtList = node.FP.Get_stmtList(_env)
    for _, _statement := range( stmtList.Items ) {
        statement := _statement.(Nodes_NodeDownCast).ToNodes_Node()
        if Lns_op_not(convGo_ignoreNodeInInnerBlockSet.Has(statement.FP.Get_kind(_env))){
            convGo_filter_1292_(_env, statement, self, &node.Nodes_Node)
        }
    }
}

// 5499: decl @lune.@base.@convGo.convFilter.processProvide
func (self *convGo_convFilter) ProcessProvide(_env *LnsEnv, node *Nodes_ProvideNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processProvide"
    Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
}

// 5505: decl @lune.@base.@convGo.convFilter.processAlias
func (self *convGo_convFilter) ProcessAlias(_env *LnsEnv, node *Nodes_AliasNode,_opt LnsAny) {
}

// 5510: decl @lune.@base.@convGo.convFilter.processBoxing
func (self *convGo_convFilter) ProcessBoxing(_env *LnsEnv, node *Nodes_BoxingNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processBoxing"
    Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
}

// 5516: decl @lune.@base.@convGo.convFilter.processUnboxing
func (self *convGo_convFilter) ProcessUnboxing(_env *LnsEnv, node *Nodes_UnboxingNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processUnboxing"
    Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
}

// 5522: decl @lune.@base.@convGo.convFilter.processLiteralList
func (self *convGo_convFilter) ProcessLiteralList(_env *LnsEnv, node *Nodes_LiteralListNode,_opt LnsAny) {
    self.FP.Write(_env, "NewLnsList(")
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.expList2Slice(_env, expList, true)
        } else {
            self.FP.Write(_env, "[]LnsAny{}")
        }
    }
    self.FP.Write(_env, ")")
}

// 5535: decl @lune.@base.@convGo.convFilter.processLiteralSet
func (self *convGo_convFilter) ProcessLiteralSet(_env *LnsEnv, node *Nodes_LiteralSetNode,_opt LnsAny) {
    self.FP.Write(_env, "NewLnsSet(")
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.expList2Slice(_env, expList, true)
        } else {
            self.FP.Write(_env, "[]LnsAny{}")
        }
    }
    self.FP.Write(_env, ")")
}

// 5549: decl @lune.@base.@convGo.convFilter.processLiteralMap
func (self *convGo_convFilter) ProcessLiteralMap(_env *LnsEnv, node *Nodes_LiteralMapNode,_opt LnsAny) {
    var hasNilable bool
    hasNilable = false
    self.FP.Write(_env, "NewLnsMap( map[LnsAny]LnsAny{")
    for _, _pair := range( node.FP.Get_pairList(_env).Items ) {
        pair := _pair.(Nodes_PairItemDownCast).ToNodes_PairItem()
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( pair.FP.Get_key(_env).FP.Get_kind(_env) == Nodes_NodeKind_get_LiteralNil(_env)) ||
            _env.SetStackVal( pair.FP.Get_val(_env).FP.Get_kind(_env) == Nodes_NodeKind_get_LiteralNil(_env)) ).(bool){
        } else { 
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( pair.FP.Get_key(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Nilable) ||
                _env.SetStackVal( pair.FP.Get_val(_env).FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Nilable) ).(bool){
                hasNilable = true
                
            }
            convGo_filter_1292_(_env, pair.FP.Get_key(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ":")
            convGo_filter_1292_(_env, pair.FP.Get_val(_env), self, &node.Nodes_Node)
            self.FP.Write(_env, ",")
        }
    }
    self.FP.Write(_env, "})")
    if hasNilable{
        self.FP.Write(_env, ".Correct()")
    }
}

// 5577: decl @lune.@base.@convGo.convFilter.processLiteralArray
func (self *convGo_convFilter) ProcessLiteralArray(_env *LnsEnv, node *Nodes_LiteralArrayNode,_opt LnsAny) {
    self.FP.Write(_env, "NewLnsList(")
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.expList2Slice(_env, expList, true)
        } else {
            self.FP.Write(_env, "[]LnsAny{}")
        }
    }
    self.FP.Write(_env, ")")
}

// 5591: decl @lune.@base.@convGo.convFilter.processLiteralChar
func (self *convGo_convFilter) ProcessLiteralChar(_env *LnsEnv, node *Nodes_LiteralCharNode,_opt LnsAny) {
    self.FP.Write(_env, _env.LuaVM.String_format("%d", []LnsAny{node.FP.Get_num(_env)}))
}

// 5597: decl @lune.@base.@convGo.convFilter.processLiteralInt
func (self *convGo_convFilter) ProcessLiteralInt(_env *LnsEnv, node *Nodes_LiteralIntNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
}

// 5603: decl @lune.@base.@convGo.convFilter.processLiteralReal
func (self *convGo_convFilter) ProcessLiteralReal(_env *LnsEnv, node *Nodes_LiteralRealNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
}

// 5609: decl @lune.@base.@convGo.convFilter.processLiteralString
func (self *convGo_convFilter) ProcessLiteralString(_env *LnsEnv, node *Nodes_LiteralStringNode,_opt LnsAny) {
    var txt string
    txt = node.FP.Get_token(_env).Txt
    {
        _expList := node.FP.Get_dddParam(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            self.FP.Write(_env, _env.LuaVM.String_format("%s.String_format(%s, ", []LnsAny{self.env.FP.getLuavm(_env, node.FP.IsThreading(_env)), convGo_str2gostr_1644_(_env, txt)}))
            self.FP.processSetFromExpList(_env, self.FP.getConvExpName(_env, node.FP.Get_id(_env), expList), NewLnsList([]LnsAny{Ast_TypeInfo2Stem(Ast_builtinTypeDDD)}), expList, false)
            self.FP.Write(_env, ")")
        } else {
            self.FP.Write(_env, _env.LuaVM.String_format("%s", []LnsAny{convGo_str2gostr_1644_(_env, txt)}))
        }
    }
}

// 5625: decl @lune.@base.@convGo.convFilter.processLiteralBool
func (self *convGo_convFilter) ProcessLiteralBool(_env *LnsEnv, node *Nodes_LiteralBoolNode,_opt LnsAny) {
    self.FP.Write(_env, node.FP.Get_token(_env).Txt)
}

// 5631: decl @lune.@base.@convGo.convFilter.processLiteralNil
func (self *convGo_convFilter) ProcessLiteralNil(_env *LnsEnv, node *Nodes_LiteralNilNode,_opt LnsAny) {
    self.FP.Write(_env, "nil")
}

// 5637: decl @lune.@base.@convGo.convFilter.processBreak
func (self *convGo_convFilter) ProcessBreak(_env *LnsEnv, node *Nodes_BreakNode,_opt LnsAny) {
    self.FP.Write(_env, "break")
    self.FP.Writeln(_env, "")
}

// 5644: decl @lune.@base.@convGo.convFilter.processLiteralSymbol
func (self *convGo_convFilter) ProcessLiteralSymbol(_env *LnsEnv, node *Nodes_LiteralSymbolNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processLiteralSymbol"
    Util_err(_env, _env.LuaVM.String_format("not support -- %s", []LnsAny{__func__}))
}

// 5650: decl @lune.@base.@convGo.convFilter.processLuneControl
func (self *convGo_convFilter) ProcessLuneControl(_env *LnsEnv, node *Nodes_LuneControlNode,_opt LnsAny) {
    switch node.FP.Get_pragma(_env).(type) {
    case *LuneControl_Pragma__default__init:
    case *LuneControl_Pragma__default__init_old:
    case *LuneControl_Pragma__disable_mut_control:
    case *LuneControl_Pragma__ignore_symbol_:
    case *LuneControl_Pragma__load__lune_module:
    case *LuneControl_Pragma__limit_conv_code:
    case *LuneControl_Pragma__use_async:
    case *LuneControl_Pragma__run_async_pipe:
        self.FP.Writeln(_env, "go self.LoopMain()")
    case *LuneControl_Pragma__run_async_runner:
        if self.option.FP.Get_enableRunner(_env){
            self.FP.Writeln(_env, "LnsRun( self )")
        } else { 
            if self.option.FP.Get_addEnvArg(_env){
                self.FP.Writeln(_env, "self.Run( _env )")
            } else { 
                self.FP.Writeln(_env, "self.Run()")
            }
        }
    case *LuneControl_Pragma__default_async_func:
    case *LuneControl_Pragma__default_async_all:
    case *LuneControl_Pragma__default_async_this_class:
    case *LuneControl_Pragma__default_noasync_this_class:
    }
}

// 5693: decl @lune.@base.@convGo.convFilter.processAbbr
func (self *convGo_convFilter) ProcessAbbr(_env *LnsEnv, node *Nodes_AbbrNode,_opt LnsAny) {
    __func__ := "@lune.@base.@convGo.convFilter.processAbbr"
    Util_err(_env, _env.LuaVM.String_format("not support -- %s:%d", []LnsAny{__func__, node.FP.Get_pos(_env).LineNo}))
}


// declaration Class -- FuncConv
type convGo_FuncConvMtd interface {
    Get_argList(_env *LnsEnv) *LnsList
    Get_retList(_env *LnsEnv) *LnsList
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
func NewconvGo_FuncConv(_env *LnsEnv, arg1 *LnsList) *convGo_FuncConv {
    obj := &convGo_FuncConv{}
    obj.FP = obj
    obj.InitconvGo_FuncConv(_env, arg1)
    return obj
}
func (self *convGo_FuncConv) Get_argList(_env *LnsEnv) *LnsList{ return self.argList }
func (self *convGo_FuncConv) Get_retList(_env *LnsEnv) *LnsList{ return self.retList }
// 1453: DeclConstr
func (self *convGo_FuncConv) InitconvGo_FuncConv(_env *LnsEnv, retList *LnsList) {
    self.argList = NewLnsList([]LnsAny{})
    
    self.retList = retList
    
}


func Lns_convGo_init(_env *LnsEnv) {
    if init_convGo { return }
    init_convGo = true
    convGo__mod__ = "@lune.@base.@convGo"
    Lns_InitMod()
    Lns_Ver_init(_env)
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Util_init(_env)
    Lns_TransUnit_init(_env)
    Lns_LuaVer_init(_env)
    Lns_Parser_init(_env)
    Lns_LuneControl_init(_env)
    Lns_Types_init(_env)
    Lns_Option_init(_env)
    convGo_MaxNilAccNum = 3
    convGo_ignoreNodeInInnerBlockSet = NewLnsSet([]LnsAny{Nodes_NodeKind_get_DeclAlge(_env), Nodes_NodeKind_get_DeclEnum(_env), Nodes_NodeKind_get_DeclMethod(_env), Nodes_NodeKind_get_DeclForm(_env), Nodes_NodeKind_get_DeclMacro(_env), Nodes_NodeKind_get_TestCase(_env)})
    convGo_golanKeywordSet = NewLnsSet([]LnsAny{"func", "select", "defer", "go", "chan", "package", "const", "fallthrough", "range", "continue", "var", "map", "default"})
    
    
    convGo_type2LnsItemKindMap = NewLnsMap( map[LnsAny]LnsAny{Ast_builtinTypeInt:"LnsItemKindInt",Ast_builtinTypeReal:"LnsItemKindReal",Ast_builtinTypeString:"LnsItemKindStr",})
    convGo_type2FromStemNameMap = NewLnsMap( map[LnsAny]LnsAny{Ast_builtinTypeInt:"Lns_ToInt",Ast_builtinTypeReal:"Lns_ToReal",Ast_builtinTypeBool:"Lns_ToBool",Ast_builtinTypeString:"Lns_ToStr",Ast_builtinTypeStem:"Lns_ToStem",})
    convGo_op2map = NewLnsMap( map[LnsAny]LnsAny{"..":"+","~=":"!=",})
    Lns_convLua_init(_env)
}
func init() {
    init_convGo = false
}
