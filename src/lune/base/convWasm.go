// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_convWasm bool
var convWasm__mod__ string
var convWasm_strCmpOpMap *LnsMap2_[string,LnsInt]
var convWasm_i64CmpOpMap *LnsMap2_[string,LnsInt]
var convWasm_f64CmpOpMap *LnsMap2_[string,LnsInt]
var convWasm_mappingMethodSet *LnsSet2_[string]
// decl alge -- VarLoc
type convWasm_VarLoc = LnsAny
type convWasm_VarLoc__Captured struct{
Val1 LnsInt
}
func (self *convWasm_VarLoc__Captured) GetTxt() string {
return "VarLoc.Captured"
}
type convWasm_VarLoc__CellLocal struct{
Val1 LnsInt
}
func (self *convWasm_VarLoc__CellLocal) GetTxt() string {
return "VarLoc.CellLocal"
}
type convWasm_VarLoc__Global struct{
Val1 LnsInt
}
func (self *convWasm_VarLoc__Global) GetTxt() string {
return "VarLoc.Global"
}
type convWasm_VarLoc__Local struct{
Val1 LnsInt
}
func (self *convWasm_VarLoc__Local) GetTxt() string {
return "VarLoc.Local"
}
type convWasm_ArgGetter_19_ func (_env *LnsEnv, arg1 LnsInt,arg2 *WasmBin_ValType) *WasmBin_Ins
// for 379
func convWasm_convExp0_1309(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 711
func convWasm_convExp1_1554(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 734
func convWasm_convExp1_1752(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 769
func convWasm_convExp1_1983(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 787
func convWasm_convExp1_2149(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 835
func convWasm_convExp1_2519(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 1096
func convWasm_convExp1_4268(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 1209
func convWasm_convExp1_4811(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 1244
func convWasm_convExp1_5249(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 1304
func convWasm_convExp1_5825(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 1371
func convWasm_convExp1_6459(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 1453
func convWasm_convExp1_7260(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 1679
func convWasm_convExp1_8578(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*Ast_TypeInfo])
}
// for 1759
func convWasm_convExp1_9143(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 2037
func convWasm_convExp2_2312(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 2057
func convWasm_convExp2_2567(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 2077
func convWasm_convExp2_2805(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 2103
func convWasm_convExp2_2989(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 2168
func convWasm_convExp2_3698(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 2366
func convWasm_convExp2_5156(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 2597
func convWasm_convExp2_6866(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*Ast_TypeInfo])
}
// for 2744
func convWasm_convExp2_8094(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*Ast_TypeInfo])
}
// for 2803
func convWasm_convExp2_8651(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*Ast_TypeInfo])
}
// for 2928
func convWasm_convExp3_403(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*Ast_TypeInfo])
}
// for 3078
func convWasm_convExp3_1604(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 3279
func convWasm_convExp3_3120(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 3307
func convWasm_convExp3_3508(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 3348
func convWasm_convExp3_4120(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 3362
func convWasm_convExp3_4243(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 3685
func convWasm_convExp3_7060(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 4015
func convWasm_convExp4_958(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 4034
func convWasm_convExp4_1114(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*Ast_TypeInfo])
}
// for 4124
func convWasm_convExp4_1688(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 4146
func convWasm_convExp4_1876(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 4166
func convWasm_convExp4_2049(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 4239
func convWasm_convExp4_2416(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 4351
func convWasm_convExp4_3764(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 4382
func convWasm_convExp4_4013(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 4585
func convWasm_convExp4_6840(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 4795
func convWasm_convExp4_8452(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*Ast_TypeInfo])
}
// for 5542
func convWasm_convExp0_9393(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*Ast_TypeInfo])
}
// for 5557
func convWasm_convExp0_9512(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*Ast_TypeInfo])
}
// for 5781
func convWasm_convExp0_11019(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 6219
func convWasm_convExp0_14291(arg1 []LnsAny) (*WasmBin_Ins, *LnsList2_[*WasmBin_ValType], *LnsList2_[*WasmBin_ValType]) {
    return Lns_getFromMulti( arg1, 0 ).(*WasmBin_Ins), Lns_getFromMulti( arg1, 1 ).(*LnsList2_[*WasmBin_ValType]), Lns_getFromMulti( arg1, 2 ).(*LnsList2_[*WasmBin_ValType])
}
// for 6353
func convWasm_convExp0_15179(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 6379
func convWasm_convExp0_15386(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 6623
func convWasm_convExp0_16609(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// 74: decl @lune.@base.@convWasm.lg
func convWasm_lg_2_(_env *LnsEnv, index LnsInt) *WasmBin_Ins {
    return WasmBin_localGet(_env, index)
}

// 77: decl @lune.@base.@convWasm.ls
func convWasm_ls_3_(_env *LnsEnv, index LnsInt,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_localSet(_env, index, val)
}

// 80: decl @lune.@base.@convWasm.p
func convWasm_p_4_(_env *LnsEnv, code LnsInt,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_plain(_env, code, args)
}

// 507: decl @lune.@base.@convWasm.isScalar
func convWasm_isScalar_15_(_env *LnsEnv, rep *WasmBin_ValType) bool {
    return Lns_op_not(rep.FP.IsRef(_env))
}

// 625: decl @lune.@base.@convWasm.isMethodSym
func convWasm_isMethodSym_16_(_env *LnsEnv, symbol *Ast_SymbolInfo) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( symbol.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) &&
        _env.SetStackVal( Lns_op_not(symbol.FP.Get_staticFlag(_env))) &&
        _env.SetStackVal( symbol.FP.Get_name(_env) != "__init") &&
        _env.SetStackVal( symbol.FP.Get_name(_env) != "__free") ).(bool)
}

// 2861: decl @lune.@base.@convWasm.append
func convWasm_append_17_(_env *LnsEnv, list *LnsList2_[*WasmBin_Ins],items *LnsList2_[*WasmBin_Ins]) {
    for _, _item := range( items.Items ) {
        item := _item
        list.Insert(item)
    }
}







// 4390: decl @lune.@base.@convWasm.convFilter.fnFromConv.ok.ifThen
func convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env *LnsEnv, cond *WasmBin_Ins,body *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_ifElse(_env, cond, body, nil, nil)
}




// 4830: decl @lune.@base.@convWasm.makeIfChain
func convWasm_makeIfChain_18_(_env *LnsEnv, conds *LnsList2_[*WasmBin_Ins],bodies *LnsList2_[*LnsList2_[*WasmBin_Ins]],elseBody LnsAny) *WasmBin_Ins {
    var acc LnsAny
    acc = elseBody
    {
        var _forFrom0 LnsInt = conds.Len()
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
            acc = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, conds.GetAt(index), bodies.GetAt(index), acc, nil)))
            _forWork0 += _forDelta0
        }
    }
    if acc != nil{
        acc_403 := acc.(*LnsList2_[*WasmBin_Ins])
        return WasmBin_seq(_env, acc_403)
    }
    return WasmBin_nop(_env)
}





// 6387: decl @lune.@base.@convWasm.getClassDepth
func convWasm_getClassDepth_20_(_env *LnsEnv, info *convWasm_ClassInfo) LnsInt {
    var depth LnsInt
    depth = 0
    var work LnsAny
    work = info.FP.Get_parent(_env)
    for  {
        var parent *convWasm_ClassInfo
        
        {
            _parent := work
            if _parent == nil{
                break
            } else {
                parent = _parent.(*convWasm_ClassInfo)
            }
        }
        depth = depth + 1
        work = parent.FP.Get_parent(_env)
    }
    return depth
}

func convWasm_convFilter_buildToMapDyn___anonymous_0_(_env *LnsEnv, info1 *convWasm_ClassInfo,info2 *convWasm_ClassInfo) bool {
    var depth1 LnsInt
    depth1 = convWasm_getClassDepth_20_(_env, info1)
    var depth2 LnsInt
    depth2 = convWasm_getClassDepth_20_(_env, info2)
    if depth1 != depth2{
        return depth1 > depth2
    }
    return info1.FP.Get_structType(_env) < info2.FP.Get_structType(_env)
}

// 6528: decl @lune.@base.@convWasm.getMacroExp
func convWasm_getMacroExp_21_(_env *LnsEnv, stmt *Nodes_Node) LnsAny {
    {
        _stmtExp := Nodes_StmtExpNodeDownCastF(stmt.FP)
        if !Lns_IsNil( _stmtExp ) {
            stmtExp := _stmtExp.(*Nodes_StmtExpNode)
            return Nodes_ExpMacroExpNodeDownCastF(stmtExp.FP.Get_exp(_env).FP)
        }
    }
    return nil
}


// 6710: decl @lune.@base.@convWasm.createFilter
func ConvWasm_createFilter(_env *LnsEnv, streamName string,stream Lns_oStream,ast *AstInfo_ASTInfo,importAstList *LnsList2_[*AstInfo_ASTInfo],embedLuaList *LnsList2_[*ConvWasm_EmbedLua],bindLnsMap *LnsMap2_[string,string]) *Nodes_Filter {
    return &NewconvWasm_convFilter(_env, streamName, stream, ast, importAstList, embedLuaList, bindLnsMap).Nodes_Filter
}

// declaration Class -- Opt
type ConvWasm_OptMtd interface {
}
type ConvWasm_Opt struct {
    Parent *Nodes_Node
    FP ConvWasm_OptMtd
}
func ConvWasm_Opt2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvWasm_Opt).FP
}
func ConvWasm_Opt_toSlice(slice []LnsAny) []*ConvWasm_Opt {
    ret := make([]*ConvWasm_Opt, len(slice))
    for index, val := range slice {
        ret[index] = val.(ConvWasm_OptDownCast).ToConvWasm_Opt()
    }
    return ret
}
type ConvWasm_OptDownCast interface {
    ToConvWasm_Opt() *ConvWasm_Opt
}
func ConvWasm_OptDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvWasm_OptDownCast)
    if ok { return work.ToConvWasm_Opt() }
    return nil
}
func (obj *ConvWasm_Opt) ToConvWasm_Opt() *ConvWasm_Opt {
    return obj
}
func NewConvWasm_Opt(_env *LnsEnv, arg1 *Nodes_Node) *ConvWasm_Opt {
    obj := &ConvWasm_Opt{}
    obj.FP = obj
    obj.InitConvWasm_Opt(_env, arg1)
    return obj
}
// 54: DeclConstr
func (self *ConvWasm_Opt) InitConvWasm_Opt(_env *LnsEnv, parent *Nodes_Node) {
    self.Parent = parent
}


// declaration Class -- Option
type ConvWasm_OptionMtd interface {
}
type ConvWasm_Option struct {
    FP ConvWasm_OptionMtd
}
func ConvWasm_Option2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvWasm_Option).FP
}
func ConvWasm_Option_toSlice(slice []LnsAny) []*ConvWasm_Option {
    ret := make([]*ConvWasm_Option, len(slice))
    for index, val := range slice {
        ret[index] = val.(ConvWasm_OptionDownCast).ToConvWasm_Option()
    }
    return ret
}
type ConvWasm_OptionDownCast interface {
    ToConvWasm_Option() *ConvWasm_Option
}
func ConvWasm_OptionDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvWasm_OptionDownCast)
    if ok { return work.ToConvWasm_Option() }
    return nil
}
func (obj *ConvWasm_Option) ToConvWasm_Option() *ConvWasm_Option {
    return obj
}
func NewConvWasm_Option(_env *LnsEnv) *ConvWasm_Option {
    obj := &ConvWasm_Option{}
    obj.FP = obj
    obj.InitConvWasm_Option(_env)
    return obj
}
// 60: DeclConstr
func (self *ConvWasm_Option) InitConvWasm_Option(_env *LnsEnv) {
}


// declaration Class -- Val
type convWasm_ValMtd interface {
    Get_ins(_env *LnsEnv) *WasmBin_Ins
    Get_rep(_env *LnsEnv) *WasmBin_ValType
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
}
type convWasm_Val struct {
    ins *WasmBin_Ins
    rep *WasmBin_ValType
    typeInfo *Ast_TypeInfo
    FP convWasm_ValMtd
}
func convWasm_Val2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convWasm_Val).FP
}
func convWasm_Val_toSlice(slice []LnsAny) []*convWasm_Val {
    ret := make([]*convWasm_Val, len(slice))
    for index, val := range slice {
        ret[index] = val.(convWasm_ValDownCast).ToconvWasm_Val()
    }
    return ret
}
type convWasm_ValDownCast interface {
    ToconvWasm_Val() *convWasm_Val
}
func convWasm_ValDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convWasm_ValDownCast)
    if ok { return work.ToconvWasm_Val() }
    return nil
}
func (obj *convWasm_Val) ToconvWasm_Val() *convWasm_Val {
    return obj
}
func NewconvWasm_Val(_env *LnsEnv, arg1 *WasmBin_Ins, arg2 *WasmBin_ValType, arg3 *Ast_TypeInfo) *convWasm_Val {
    obj := &convWasm_Val{}
    obj.FP = obj
    obj.InitconvWasm_Val(_env, arg1, arg2, arg3)
    return obj
}
func (self *convWasm_Val) InitconvWasm_Val(_env *LnsEnv, arg1 *WasmBin_Ins, arg2 *WasmBin_ValType, arg3 *Ast_TypeInfo) {
    self.ins = arg1
    self.rep = arg2
    self.typeInfo = arg3
}
func (self *convWasm_Val) Get_ins(_env *LnsEnv) *WasmBin_Ins{ return self.ins }
func (self *convWasm_Val) Get_rep(_env *LnsEnv) *WasmBin_ValType{ return self.rep }
func (self *convWasm_Val) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }

// declaration Class -- SlotInfo
type convWasm_SlotInfoMtd interface {
    Get_funcType(_env *LnsEnv) LnsInt
    Get_impl(_env *LnsEnv) LnsInt
    Get_methodType(_env *LnsEnv) *Ast_TypeInfo
    Get_name(_env *LnsEnv) string
    Set_impl(_env *LnsEnv, arg1 LnsInt)
}
type convWasm_SlotInfo struct {
    name string
    funcType LnsInt
    methodType *Ast_TypeInfo
    impl LnsInt
    FP convWasm_SlotInfoMtd
}
func convWasm_SlotInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convWasm_SlotInfo).FP
}
func convWasm_SlotInfo_toSlice(slice []LnsAny) []*convWasm_SlotInfo {
    ret := make([]*convWasm_SlotInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(convWasm_SlotInfoDownCast).ToconvWasm_SlotInfo()
    }
    return ret
}
type convWasm_SlotInfoDownCast interface {
    ToconvWasm_SlotInfo() *convWasm_SlotInfo
}
func convWasm_SlotInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convWasm_SlotInfoDownCast)
    if ok { return work.ToconvWasm_SlotInfo() }
    return nil
}
func (obj *convWasm_SlotInfo) ToconvWasm_SlotInfo() *convWasm_SlotInfo {
    return obj
}
func NewconvWasm_SlotInfo(_env *LnsEnv, arg1 string, arg2 LnsInt, arg3 *Ast_TypeInfo, arg4 LnsInt) *convWasm_SlotInfo {
    obj := &convWasm_SlotInfo{}
    obj.FP = obj
    obj.InitconvWasm_SlotInfo(_env, arg1, arg2, arg3, arg4)
    return obj
}
func (self *convWasm_SlotInfo) InitconvWasm_SlotInfo(_env *LnsEnv, arg1 string, arg2 LnsInt, arg3 *Ast_TypeInfo, arg4 LnsInt) {
    self.name = arg1
    self.funcType = arg2
    self.methodType = arg3
    self.impl = arg4
}
func (self *convWasm_SlotInfo) Get_name(_env *LnsEnv) string{ return self.name }
func (self *convWasm_SlotInfo) Get_funcType(_env *LnsEnv) LnsInt{ return self.funcType }
func (self *convWasm_SlotInfo) Get_methodType(_env *LnsEnv) *Ast_TypeInfo{ return self.methodType }
func (self *convWasm_SlotInfo) Get_impl(_env *LnsEnv) LnsInt{ return self.impl }
func (self *convWasm_SlotInfo) Set_impl(_env *LnsEnv, arg1 LnsInt){ self.impl = arg1 }

// declaration Class -- AdvertiseMethod
type convWasm_AdvertiseMethodMtd interface {
    Get_impl(_env *LnsEnv) LnsInt
    Get_memberName(_env *LnsEnv) string
    Get_memberType(_env *LnsEnv) *Ast_TypeInfo
    Get_methodName(_env *LnsEnv) string
    Get_methodType(_env *LnsEnv) *Ast_TypeInfo
}
type convWasm_AdvertiseMethod struct {
    memberName string
    memberType *Ast_TypeInfo
    methodName string
    methodType *Ast_TypeInfo
    impl LnsInt
    FP convWasm_AdvertiseMethodMtd
}
func convWasm_AdvertiseMethod2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convWasm_AdvertiseMethod).FP
}
func convWasm_AdvertiseMethod_toSlice(slice []LnsAny) []*convWasm_AdvertiseMethod {
    ret := make([]*convWasm_AdvertiseMethod, len(slice))
    for index, val := range slice {
        ret[index] = val.(convWasm_AdvertiseMethodDownCast).ToconvWasm_AdvertiseMethod()
    }
    return ret
}
type convWasm_AdvertiseMethodDownCast interface {
    ToconvWasm_AdvertiseMethod() *convWasm_AdvertiseMethod
}
func convWasm_AdvertiseMethodDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convWasm_AdvertiseMethodDownCast)
    if ok { return work.ToconvWasm_AdvertiseMethod() }
    return nil
}
func (obj *convWasm_AdvertiseMethod) ToconvWasm_AdvertiseMethod() *convWasm_AdvertiseMethod {
    return obj
}
func NewconvWasm_AdvertiseMethod(_env *LnsEnv, arg1 string, arg2 *Ast_TypeInfo, arg3 string, arg4 *Ast_TypeInfo, arg5 LnsInt) *convWasm_AdvertiseMethod {
    obj := &convWasm_AdvertiseMethod{}
    obj.FP = obj
    obj.InitconvWasm_AdvertiseMethod(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *convWasm_AdvertiseMethod) InitconvWasm_AdvertiseMethod(_env *LnsEnv, arg1 string, arg2 *Ast_TypeInfo, arg3 string, arg4 *Ast_TypeInfo, arg5 LnsInt) {
    self.memberName = arg1
    self.memberType = arg2
    self.methodName = arg3
    self.methodType = arg4
    self.impl = arg5
}
func (self *convWasm_AdvertiseMethod) Get_memberName(_env *LnsEnv) string{ return self.memberName }
func (self *convWasm_AdvertiseMethod) Get_memberType(_env *LnsEnv) *Ast_TypeInfo{ return self.memberType }
func (self *convWasm_AdvertiseMethod) Get_methodName(_env *LnsEnv) string{ return self.methodName }
func (self *convWasm_AdvertiseMethod) Get_methodType(_env *LnsEnv) *Ast_TypeInfo{ return self.methodType }
func (self *convWasm_AdvertiseMethod) Get_impl(_env *LnsEnv) LnsInt{ return self.impl }

// declaration Class -- ClassInfo
type convWasm_ClassInfoMtd interface {
    Get_advertiseList(_env *LnsEnv) *LnsList2_[*convWasm_AdvertiseMethod]
    Get_ctorFunc(_env *LnsEnv) LnsInt
    Get_ctorFuncType(_env *LnsEnv) LnsInt
    Get_fieldNames(_env *LnsEnv) *LnsList2_[string]
    Get_fieldReps(_env *LnsEnv) *LnsList2_[*WasmBin_ValType]
    Get_fieldTypes(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]
    Get_ifList(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]
    Get_name2field(_env *LnsEnv) *LnsMap2_[string,LnsInt]
    Get_name2slot(_env *LnsEnv) *LnsMap2_[string,LnsInt]
    Get_parent(_env *LnsEnv) LnsAny
    Get_slotList(_env *LnsEnv) *LnsList2_[*convWasm_SlotInfo]
    Get_structType(_env *LnsEnv) LnsInt
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_vtGlobal(_env *LnsEnv) LnsInt
    Get_vtType(_env *LnsEnv) LnsInt
    Set_ctorFunc(_env *LnsEnv, arg1 LnsInt)
    Set_ctorFuncType(_env *LnsEnv, arg1 LnsInt)
}
type convWasm_ClassInfo struct {
    typeInfo *Ast_TypeInfo
    structType LnsInt
    vtType LnsInt
    vtGlobal LnsInt
    parent LnsAny
    fieldNames *LnsList2_[string]
    fieldReps *LnsList2_[*WasmBin_ValType]
    fieldTypes *LnsList2_[*Ast_TypeInfo]
    name2field *LnsMap2_[string,LnsInt]
    slotList *LnsList2_[*convWasm_SlotInfo]
    name2slot *LnsMap2_[string,LnsInt]
    ifList *LnsList2_[*Ast_TypeInfo]
    ctorFunc LnsInt
    advertiseList *LnsList2_[*convWasm_AdvertiseMethod]
    ctorFuncType LnsInt
    FP convWasm_ClassInfoMtd
}
func convWasm_ClassInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convWasm_ClassInfo).FP
}
func convWasm_ClassInfo_toSlice(slice []LnsAny) []*convWasm_ClassInfo {
    ret := make([]*convWasm_ClassInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(convWasm_ClassInfoDownCast).ToconvWasm_ClassInfo()
    }
    return ret
}
type convWasm_ClassInfoDownCast interface {
    ToconvWasm_ClassInfo() *convWasm_ClassInfo
}
func convWasm_ClassInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convWasm_ClassInfoDownCast)
    if ok { return work.ToconvWasm_ClassInfo() }
    return nil
}
func (obj *convWasm_ClassInfo) ToconvWasm_ClassInfo() *convWasm_ClassInfo {
    return obj
}
func NewconvWasm_ClassInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 LnsInt, arg4 LnsInt, arg5 LnsAny) *convWasm_ClassInfo {
    obj := &convWasm_ClassInfo{}
    obj.FP = obj
    obj.InitconvWasm_ClassInfo(_env, arg1, arg2, arg3, arg4, arg5)
    return obj
}
func (self *convWasm_ClassInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *convWasm_ClassInfo) Get_structType(_env *LnsEnv) LnsInt{ return self.structType }
func (self *convWasm_ClassInfo) Get_vtType(_env *LnsEnv) LnsInt{ return self.vtType }
func (self *convWasm_ClassInfo) Get_vtGlobal(_env *LnsEnv) LnsInt{ return self.vtGlobal }
func (self *convWasm_ClassInfo) Get_parent(_env *LnsEnv) LnsAny{ return self.parent }
func (self *convWasm_ClassInfo) Get_fieldNames(_env *LnsEnv) *LnsList2_[string]{ return self.fieldNames }
func (self *convWasm_ClassInfo) Get_fieldReps(_env *LnsEnv) *LnsList2_[*WasmBin_ValType]{ return self.fieldReps }
func (self *convWasm_ClassInfo) Get_fieldTypes(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]{ return self.fieldTypes }
func (self *convWasm_ClassInfo) Get_name2field(_env *LnsEnv) *LnsMap2_[string,LnsInt]{ return self.name2field }
func (self *convWasm_ClassInfo) Get_slotList(_env *LnsEnv) *LnsList2_[*convWasm_SlotInfo]{ return self.slotList }
func (self *convWasm_ClassInfo) Get_name2slot(_env *LnsEnv) *LnsMap2_[string,LnsInt]{ return self.name2slot }
func (self *convWasm_ClassInfo) Get_ifList(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]{ return self.ifList }
func (self *convWasm_ClassInfo) Get_ctorFunc(_env *LnsEnv) LnsInt{ return self.ctorFunc }
func (self *convWasm_ClassInfo) Set_ctorFunc(_env *LnsEnv, arg1 LnsInt){ self.ctorFunc = arg1 }
func (self *convWasm_ClassInfo) Get_advertiseList(_env *LnsEnv) *LnsList2_[*convWasm_AdvertiseMethod]{ return self.advertiseList }
func (self *convWasm_ClassInfo) Get_ctorFuncType(_env *LnsEnv) LnsInt{ return self.ctorFuncType }
func (self *convWasm_ClassInfo) Set_ctorFuncType(_env *LnsEnv, arg1 LnsInt){ self.ctorFuncType = arg1 }
// 150: DeclConstr
func (self *convWasm_ClassInfo) InitconvWasm_ClassInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo,structType LnsInt,vtType LnsInt,vtGlobal LnsInt,parent LnsAny) {
    self.typeInfo = typeInfo
    self.structType = structType
    self.vtType = vtType
    self.vtGlobal = vtGlobal
    self.parent = parent
    self.fieldNames = NewLnsList2_[string]([]string{})
    self.fieldReps = NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{})
    self.fieldTypes = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    self.name2field = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    self.slotList = NewLnsList2_[*convWasm_SlotInfo]([]*convWasm_SlotInfo{})
    self.name2slot = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    self.ifList = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    self.ctorFunc = -1
    self.advertiseList = NewLnsList2_[*convWasm_AdvertiseMethod]([]*convWasm_AdvertiseMethod{})
    self.ctorFuncType = -1
    if parent != nil{
        parent_118 := parent.(*convWasm_ClassInfo)
        for _, _name := range( parent_118.FP.Get_fieldNames(_env).Items ) {
            name := _name
            self.fieldNames.Insert(name)
        }
        for _, _rep := range( parent_118.FP.Get_fieldReps(_env).Items ) {
            rep := _rep
            self.fieldReps.Insert(rep)
        }
        for _, _fieldType := range( parent_118.FP.Get_fieldTypes(_env).Items ) {
            fieldType := _fieldType
            self.fieldTypes.Insert(fieldType)
        }
        for _name, _index := range( parent_118.FP.Get_name2field(_env).Items ) {
            name := _name
            index := _index
            self.name2field.Set(name,index)
        }
        for _, _slot := range( parent_118.FP.Get_slotList(_env).Items ) {
            slot := _slot
            self.slotList.Insert(NewconvWasm_SlotInfo(_env, slot.FP.Get_name(_env), slot.FP.Get_funcType(_env), slot.FP.Get_methodType(_env), slot.FP.Get_impl(_env)))
        }
        for _name, _index := range( parent_118.FP.Get_name2slot(_env).Items ) {
            name := _name
            index := _index
            self.name2slot.Set(name,index)
        }
        for _, _ifType := range( parent_118.FP.Get_ifList(_env).Items ) {
            ifType := _ifType
            self.ifList.Insert(ifType)
        }
    }
}


// declaration Class -- IfInfo
type convWasm_IfInfoMtd interface {
    Get_ifId(_env *LnsEnv) LnsInt
    Get_itType(_env *LnsEnv) LnsInt
    Get_methodFuncTypes(_env *LnsEnv) *LnsList2_[LnsInt]
    Get_methodNames(_env *LnsEnv) *LnsList2_[string]
    Get_name2index(_env *LnsEnv) *LnsMap2_[string,LnsInt]
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
}
type convWasm_IfInfo struct {
    typeInfo *Ast_TypeInfo
    ifId LnsInt
    itType LnsInt
    methodNames *LnsList2_[string]
    methodFuncTypes *LnsList2_[LnsInt]
    name2index *LnsMap2_[string,LnsInt]
    FP convWasm_IfInfoMtd
}
func convWasm_IfInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convWasm_IfInfo).FP
}
func convWasm_IfInfo_toSlice(slice []LnsAny) []*convWasm_IfInfo {
    ret := make([]*convWasm_IfInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(convWasm_IfInfoDownCast).ToconvWasm_IfInfo()
    }
    return ret
}
type convWasm_IfInfoDownCast interface {
    ToconvWasm_IfInfo() *convWasm_IfInfo
}
func convWasm_IfInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convWasm_IfInfoDownCast)
    if ok { return work.ToconvWasm_IfInfo() }
    return nil
}
func (obj *convWasm_IfInfo) ToconvWasm_IfInfo() *convWasm_IfInfo {
    return obj
}
func NewconvWasm_IfInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 LnsInt, arg4 *LnsList2_[string], arg5 *LnsList2_[LnsInt], arg6 *LnsMap2_[string,LnsInt]) *convWasm_IfInfo {
    obj := &convWasm_IfInfo{}
    obj.FP = obj
    obj.InitconvWasm_IfInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *convWasm_IfInfo) InitconvWasm_IfInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 LnsInt, arg4 *LnsList2_[string], arg5 *LnsList2_[LnsInt], arg6 *LnsMap2_[string,LnsInt]) {
    self.typeInfo = arg1
    self.ifId = arg2
    self.itType = arg3
    self.methodNames = arg4
    self.methodFuncTypes = arg5
    self.name2index = arg6
}
func (self *convWasm_IfInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *convWasm_IfInfo) Get_ifId(_env *LnsEnv) LnsInt{ return self.ifId }
func (self *convWasm_IfInfo) Get_itType(_env *LnsEnv) LnsInt{ return self.itType }
func (self *convWasm_IfInfo) Get_methodNames(_env *LnsEnv) *LnsList2_[string]{ return self.methodNames }
func (self *convWasm_IfInfo) Get_methodFuncTypes(_env *LnsEnv) *LnsList2_[LnsInt]{ return self.methodFuncTypes }
func (self *convWasm_IfInfo) Get_name2index(_env *LnsEnv) *LnsMap2_[string,LnsInt]{ return self.name2index }

// declaration Class -- AlgeInfo
type convWasm_AlgeInfoMtd interface {
    Get_baseType(_env *LnsEnv) LnsInt
    Get_name2global(_env *LnsEnv) *LnsMap2_[string,LnsInt]
    Get_name2tag(_env *LnsEnv) *LnsMap2_[string,LnsInt]
    Get_name2type(_env *LnsEnv) *LnsMap2_[string,LnsInt]
    Get_namesGlobal(_env *LnsEnv) LnsInt
    Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo
    Get_valNames(_env *LnsEnv) *LnsList2_[string]
}
type convWasm_AlgeInfo struct {
    typeInfo *Ast_TypeInfo
    baseType LnsInt
    valNames *LnsList2_[string]
    name2tag *LnsMap2_[string,LnsInt]
    name2type *LnsMap2_[string,LnsInt]
    name2global *LnsMap2_[string,LnsInt]
    namesGlobal LnsInt
    FP convWasm_AlgeInfoMtd
}
func convWasm_AlgeInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convWasm_AlgeInfo).FP
}
func convWasm_AlgeInfo_toSlice(slice []LnsAny) []*convWasm_AlgeInfo {
    ret := make([]*convWasm_AlgeInfo, len(slice))
    for index, val := range slice {
        ret[index] = val.(convWasm_AlgeInfoDownCast).ToconvWasm_AlgeInfo()
    }
    return ret
}
type convWasm_AlgeInfoDownCast interface {
    ToconvWasm_AlgeInfo() *convWasm_AlgeInfo
}
func convWasm_AlgeInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convWasm_AlgeInfoDownCast)
    if ok { return work.ToconvWasm_AlgeInfo() }
    return nil
}
func (obj *convWasm_AlgeInfo) ToconvWasm_AlgeInfo() *convWasm_AlgeInfo {
    return obj
}
func NewconvWasm_AlgeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 *LnsList2_[string], arg4 *LnsMap2_[string,LnsInt], arg5 *LnsMap2_[string,LnsInt], arg6 *LnsMap2_[string,LnsInt], arg7 LnsInt) *convWasm_AlgeInfo {
    obj := &convWasm_AlgeInfo{}
    obj.FP = obj
    obj.InitconvWasm_AlgeInfo(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *convWasm_AlgeInfo) InitconvWasm_AlgeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 LnsInt, arg3 *LnsList2_[string], arg4 *LnsMap2_[string,LnsInt], arg5 *LnsMap2_[string,LnsInt], arg6 *LnsMap2_[string,LnsInt], arg7 LnsInt) {
    self.typeInfo = arg1
    self.baseType = arg2
    self.valNames = arg3
    self.name2tag = arg4
    self.name2type = arg5
    self.name2global = arg6
    self.namesGlobal = arg7
}
func (self *convWasm_AlgeInfo) Get_typeInfo(_env *LnsEnv) *Ast_TypeInfo{ return self.typeInfo }
func (self *convWasm_AlgeInfo) Get_baseType(_env *LnsEnv) LnsInt{ return self.baseType }
func (self *convWasm_AlgeInfo) Get_valNames(_env *LnsEnv) *LnsList2_[string]{ return self.valNames }
func (self *convWasm_AlgeInfo) Get_name2tag(_env *LnsEnv) *LnsMap2_[string,LnsInt]{ return self.name2tag }
func (self *convWasm_AlgeInfo) Get_name2type(_env *LnsEnv) *LnsMap2_[string,LnsInt]{ return self.name2type }
func (self *convWasm_AlgeInfo) Get_name2global(_env *LnsEnv) *LnsMap2_[string,LnsInt]{ return self.name2global }
func (self *convWasm_AlgeInfo) Get_namesGlobal(_env *LnsEnv) LnsInt{ return self.namesGlobal }

// declaration Class -- FuncCtx
type convWasm_FuncCtxMtd interface {
    AddLocal(_env *LnsEnv, arg1 *WasmBin_ValType) LnsInt
    Get_breakDepthList(_env *LnsEnv) *LnsList2_[LnsInt]
    Get_depth(_env *LnsEnv) LnsInt
    Get_envLocal(_env *LnsEnv) LnsInt
    Get_envType(_env *LnsEnv) LnsInt
    Get_func(_env *LnsEnv) *WasmBin_Func
    Get_retReps(_env *LnsEnv) *LnsList2_[*WasmBin_ValType]
    Get_retTypes(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]
    Get_selfClass(_env *LnsEnv) LnsAny
    Get_selfLocal(_env *LnsEnv) LnsInt
    Get_sym2loc(_env *LnsEnv) *LnsMap2_[*Ast_SymbolInfo,LnsAny]
    Set_depth(_env *LnsEnv, arg1 LnsInt)
    Set_envLocal(_env *LnsEnv, arg1 LnsInt)
    Set_envType(_env *LnsEnv, arg1 LnsInt)
    Set_selfClass(_env *LnsEnv, arg1 LnsAny)
    Set_selfLocal(_env *LnsEnv, arg1 LnsInt)
}
type convWasm_FuncCtx struct {
    _func *WasmBin_Func
    sym2loc *LnsMap2_[*Ast_SymbolInfo,LnsAny]
    retReps *LnsList2_[*WasmBin_ValType]
    retTypes *LnsList2_[*Ast_TypeInfo]
    depth LnsInt
    breakDepthList *LnsList2_[LnsInt]
    envLocal LnsInt
    envType LnsInt
    selfLocal LnsInt
    selfClass LnsAny
    FP convWasm_FuncCtxMtd
}
func convWasm_FuncCtx2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convWasm_FuncCtx).FP
}
func convWasm_FuncCtx_toSlice(slice []LnsAny) []*convWasm_FuncCtx {
    ret := make([]*convWasm_FuncCtx, len(slice))
    for index, val := range slice {
        ret[index] = val.(convWasm_FuncCtxDownCast).ToconvWasm_FuncCtx()
    }
    return ret
}
type convWasm_FuncCtxDownCast interface {
    ToconvWasm_FuncCtx() *convWasm_FuncCtx
}
func convWasm_FuncCtxDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convWasm_FuncCtxDownCast)
    if ok { return work.ToconvWasm_FuncCtx() }
    return nil
}
func (obj *convWasm_FuncCtx) ToconvWasm_FuncCtx() *convWasm_FuncCtx {
    return obj
}
func NewconvWasm_FuncCtx(_env *LnsEnv, arg1 *WasmBin_Func, arg2 *LnsList2_[*WasmBin_ValType], arg3 *LnsList2_[*Ast_TypeInfo]) *convWasm_FuncCtx {
    obj := &convWasm_FuncCtx{}
    obj.FP = obj
    obj.InitconvWasm_FuncCtx(_env, arg1, arg2, arg3)
    return obj
}
func (self *convWasm_FuncCtx) Get_func(_env *LnsEnv) *WasmBin_Func{ return self._func }
func (self *convWasm_FuncCtx) Get_sym2loc(_env *LnsEnv) *LnsMap2_[*Ast_SymbolInfo,LnsAny]{ return self.sym2loc }
func (self *convWasm_FuncCtx) Get_retReps(_env *LnsEnv) *LnsList2_[*WasmBin_ValType]{ return self.retReps }
func (self *convWasm_FuncCtx) Get_retTypes(_env *LnsEnv) *LnsList2_[*Ast_TypeInfo]{ return self.retTypes }
func (self *convWasm_FuncCtx) Get_depth(_env *LnsEnv) LnsInt{ return self.depth }
func (self *convWasm_FuncCtx) Set_depth(_env *LnsEnv, arg1 LnsInt){ self.depth = arg1 }
func (self *convWasm_FuncCtx) Get_breakDepthList(_env *LnsEnv) *LnsList2_[LnsInt]{ return self.breakDepthList }
func (self *convWasm_FuncCtx) Get_envLocal(_env *LnsEnv) LnsInt{ return self.envLocal }
func (self *convWasm_FuncCtx) Set_envLocal(_env *LnsEnv, arg1 LnsInt){ self.envLocal = arg1 }
func (self *convWasm_FuncCtx) Get_envType(_env *LnsEnv) LnsInt{ return self.envType }
func (self *convWasm_FuncCtx) Set_envType(_env *LnsEnv, arg1 LnsInt){ self.envType = arg1 }
func (self *convWasm_FuncCtx) Get_selfLocal(_env *LnsEnv) LnsInt{ return self.selfLocal }
func (self *convWasm_FuncCtx) Set_selfLocal(_env *LnsEnv, arg1 LnsInt){ self.selfLocal = arg1 }
func (self *convWasm_FuncCtx) Get_selfClass(_env *LnsEnv) LnsAny{ return self.selfClass }
func (self *convWasm_FuncCtx) Set_selfClass(_env *LnsEnv, arg1 LnsAny){ self.selfClass = arg1 }
// 236: DeclConstr
func (self *convWasm_FuncCtx) InitconvWasm_FuncCtx(_env *LnsEnv, _func *WasmBin_Func,retReps *LnsList2_[*WasmBin_ValType],retTypes *LnsList2_[*Ast_TypeInfo]) {
    self._func = _func
    self.sym2loc = NewLnsMap2_[*Ast_SymbolInfo,LnsAny]( map[*Ast_SymbolInfo]LnsAny{})
    self.retReps = retReps
    self.retTypes = retTypes
    self.depth = 0
    self.breakDepthList = NewLnsList2_[LnsInt]([]LnsInt{})
    self.envLocal = -1
    self.envType = -1
    self.selfLocal = -1
    self.selfClass = nil
}


// declaration Class -- EmbedLua
type ConvWasm_EmbedLuaMtd interface {
    Get_chunkname(_env *LnsEnv) string
    Get_code(_env *LnsEnv) string
    Get_mod(_env *LnsEnv) string
}
type ConvWasm_EmbedLua struct {
    mod string
    chunkname string
    code string
    FP ConvWasm_EmbedLuaMtd
}
func ConvWasm_EmbedLua2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*ConvWasm_EmbedLua).FP
}
func ConvWasm_EmbedLua_toSlice(slice []LnsAny) []*ConvWasm_EmbedLua {
    ret := make([]*ConvWasm_EmbedLua, len(slice))
    for index, val := range slice {
        ret[index] = val.(ConvWasm_EmbedLuaDownCast).ToConvWasm_EmbedLua()
    }
    return ret
}
type ConvWasm_EmbedLuaDownCast interface {
    ToConvWasm_EmbedLua() *ConvWasm_EmbedLua
}
func ConvWasm_EmbedLuaDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(ConvWasm_EmbedLuaDownCast)
    if ok { return work.ToConvWasm_EmbedLua() }
    return nil
}
func (obj *ConvWasm_EmbedLua) ToConvWasm_EmbedLua() *ConvWasm_EmbedLua {
    return obj
}
func NewConvWasm_EmbedLua(_env *LnsEnv, arg1 string, arg2 string, arg3 string) *ConvWasm_EmbedLua {
    obj := &ConvWasm_EmbedLua{}
    obj.FP = obj
    obj.InitConvWasm_EmbedLua(_env, arg1, arg2, arg3)
    return obj
}
func (self *ConvWasm_EmbedLua) InitConvWasm_EmbedLua(_env *LnsEnv, arg1 string, arg2 string, arg3 string) {
    self.mod = arg1
    self.chunkname = arg2
    self.code = arg3
}
func (self *ConvWasm_EmbedLua) Get_mod(_env *LnsEnv) string{ return self.mod }
func (self *ConvWasm_EmbedLua) Get_chunkname(_env *LnsEnv) string{ return self.chunkname }
func (self *ConvWasm_EmbedLua) Get_code(_env *LnsEnv) string{ return self.code }

// declaration Class -- convFilter
type convWasm_convFilterMtd interface {
    adaptClosure(_env *LnsEnv, arg1 *WasmBin_Ins, arg2 *Ast_TypeInfo, arg3 *Ast_TypeInfo) *WasmBin_Ins
    addInterface(_env *LnsEnv, arg1 *convWasm_ClassInfo, arg2 *Ast_TypeInfo)
    algeTxt(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_TypeInfo) *WasmBin_Ins
    anyNil(_env *LnsEnv, arg1 *LnsList2_[*convWasm_Val], arg2 LnsInt) *WasmBin_Ins
    assignItem(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_Node, arg3 *convWasm_Val, arg4 *convWasm_Val, arg5 *LnsList2_[*WasmBin_Ins])
    buildList(_env *LnsEnv, arg1 *LnsList2_[*convWasm_Val]) *WasmBin_Ins
    buildToMapDyn(_env *LnsEnv)
    callClosure(_env *LnsEnv, arg1 *WasmBin_Ins, arg2 *Ast_TypeInfo, arg3 LnsAny, arg4 *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins
    callMethodDyn(_env *LnsEnv, arg1 *Nodes_Node, arg2 *WasmBin_Ins, arg3 *WasmBin_ValType, arg4 *Ast_TypeInfo, arg5 string, arg6 convWasm_ArgGetter_19_, arg7 *WasmBin_Func)(*WasmBin_Ins, *LnsList2_[*WasmBin_ValType], *LnsList2_[*WasmBin_ValType])
    cellGet(_env *LnsEnv, arg1 *WasmBin_ValType, arg2 *WasmBin_Ins) *WasmBin_Ins
    coerce(_env *LnsEnv, arg1 *WasmBin_Ins, arg2 *Ast_TypeInfo, arg3 *Ast_TypeInfo) *WasmBin_Ins
    coerceRep(_env *LnsEnv, arg1 *WasmBin_Ins, arg2 *WasmBin_ValType, arg3 *WasmBin_ValType) *WasmBin_Ins
    coerceVal(_env *LnsEnv, arg1 *convWasm_Val, arg2 *Ast_TypeInfo) *WasmBin_Ins
    collectDecl(_env *LnsEnv, arg1 *Nodes_Node)
    collectDynGetters(_env *LnsEnv, arg1 *Nodes_Node)
    collectIfMethods(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *LnsMap2_[string,*Ast_TypeInfo])
    compileAccessors(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 *convWasm_ClassInfo)
    compileAdvertise(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 *convWasm_ClassInfo)
    compileDefaultCtor(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 *convWasm_ClassInfo)
    compileFunc(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo, arg3 *Ast_TypeInfo, arg4 LnsInt) *WasmBin_Ins
    compileMethod(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo, arg3 *Ast_TypeInfo, arg4 LnsInt, arg5 *convWasm_ClassInfo, arg6 *LnsList2_[*WasmBin_ValType])
    compileMethodNode(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_DeclFuncInfo)
    declareVar(_env *LnsEnv, arg1 *Ast_SymbolInfo, arg2 LnsAny) *WasmBin_Ins
    decodeString(_env *LnsEnv, arg1 *Nodes_Node, arg2 string) string
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    enterLoop(_env *LnsEnv)
    enumEq(_env *LnsEnv, arg1 *Ast_EnumTypeInfo, arg2 *WasmBin_Ins, arg3 *Ast_EnumValInfo) *WasmBin_Ins
    envGet(_env *LnsEnv, arg1 LnsInt) *WasmBin_Ins
    error(_env *LnsEnv, arg1 *Nodes_Node, arg2 string)
    exitLoop(_env *LnsEnv, arg1 *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins
    expTypeOf(_env *LnsEnv, arg1 *Nodes_Node) *Ast_TypeInfo
    expandDDD(_env *LnsEnv, arg1 *LnsList2_[*convWasm_Val], arg2 LnsInt, arg3 *LnsList2_[*WasmBin_Ins]) *LnsList2_[*convWasm_Val]
    finalizeClasses(_env *LnsEnv)
    findLoc(_env *LnsEnv, arg1 *Ast_SymbolInfo) LnsAny
    fnAnyToMap(_env *LnsEnv) LnsInt
    fnDivI64(_env *LnsEnv) LnsInt
    fnEnumAllList(_env *LnsEnv, arg1 *Ast_EnumTypeInfo) LnsInt
    fnEnumFrom(_env *LnsEnv, arg1 *Ast_EnumTypeInfo) LnsInt
    fnEnumTxt(_env *LnsEnv, arg1 *Ast_EnumTypeInfo) LnsInt
    fnFromConv(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_TypeInfo) LnsInt
    fnFromMap(_env *LnsEnv, arg1 *Nodes_Node, arg2 *convWasm_ClassInfo) LnsInt
    fnIntOrDefault(_env *LnsEnv, arg1 LnsInt) LnsInt
    fnIntOrMinus1(_env *LnsEnv) LnsInt
    fnIntOrZero(_env *LnsEnv) LnsInt
    fnIsIntVal(_env *LnsEnv) LnsInt
    fnLuaType(_env *LnsEnv) LnsInt
    fnMapToList(_env *LnsEnv) LnsInt
    fnModF64(_env *LnsEnv) LnsInt
    fnModI64(_env *LnsEnv) LnsInt
    fnShift(_env *LnsEnv, arg1 bool) LnsInt
    fnStemIndex(_env *LnsEnv) LnsInt
    fnStrJoin(_env *LnsEnv) LnsInt
    fnToMap(_env *LnsEnv, arg1 *Nodes_Node, arg2 *convWasm_ClassInfo) LnsInt
    fnToMapConv(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_TypeInfo) LnsInt
    fnToMapDyn(_env *LnsEnv) LnsInt
    fnTruthy(_env *LnsEnv) LnsInt
    getAlgeInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny
    getCaptureSyms(_env *LnsEnv, arg1 *Ast_TypeInfo) *LnsList2_[*Ast_SymbolInfo]
    getCellRef(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_SymbolInfo) *WasmBin_Ins
    getClassInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny
    getCondRetInfo(_env *LnsEnv, arg1 *Nodes_Node) LnsAny
    getCtorType(_env *LnsEnv, arg1 *Ast_TypeInfo) *Ast_TypeInfo
    getDefaultVal(_env *LnsEnv, arg1 *WasmBin_ValType) *WasmBin_Ins
    getEnumFuncName(_env *LnsEnv, arg1 *Ast_EnumTypeInfo, arg2 string) string
    getEnumValList(_env *LnsEnv, arg1 *Ast_EnumTypeInfo) *LnsList2_[*Ast_EnumValInfo]
    getExtModuleType(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsAny
    getFieldIndex(_env *LnsEnv, arg1 *Nodes_Node, arg2 *convWasm_ClassInfo, arg3 string) LnsInt
    GetFull(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    getFuncCloInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) *WasmRt_CloInfo
    getGenSrc(_env *LnsEnv, arg1 *Ast_TypeInfo) *Ast_TypeInfo
    getIfInfo(_env *LnsEnv, arg1 *Ast_TypeInfo) *convWasm_IfInfo
    getIfMethodType(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 string) LnsAny
    getLazyInit(_env *LnsEnv, arg1 *Nodes_Node) LnsAny
    getLessFormType(_env *LnsEnv, arg1 *Nodes_ExpCallNode) *Ast_TypeInfo
    getMemberType(_env *LnsEnv, arg1 *convWasm_ClassInfo, arg2 string) *Ast_TypeInfo
    getMethodFuncType(_env *LnsEnv, arg1 *Ast_TypeInfo) LnsInt
    getNilAccPrefix(_env *LnsEnv, arg1 *Nodes_Node) LnsAny
    getObjType(_env *LnsEnv, arg1 *Ast_TypeInfo) *Ast_TypeInfo
    getObjTypeForMethod(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 string) *Ast_TypeInfo
    getRep(_env *LnsEnv, arg1 *Ast_TypeInfo) *WasmBin_ValType
    getRepList(_env *LnsEnv, arg1 *LnsList2_[*Ast_TypeInfo]) *LnsList2_[*WasmBin_ValType]
    getSerrInfo(_env *LnsEnv) *convWasm_ClassInfo
    getSlotRetReps(_env *LnsEnv, arg1 *convWasm_ClassInfo, arg2 *Ast_TypeInfo) *LnsList2_[*WasmBin_ValType]
    getStdStream(_env *LnsEnv, arg1 bool) LnsInt
    getStdinGlobal(_env *LnsEnv) LnsInt
    getStrBuilderBody(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_TypeInfo, arg3 *convWasm_ClassInfo, arg4 LnsInt) LnsAny
    getStringModuleFuncMap(_env *LnsEnv) *LnsMap2_[*Ast_TypeInfo,LnsInt]
    Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl
    isAccessOf(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_Node) bool
    isFuncType(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isGlobalSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) bool
    isLuaval(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isModuleFunc(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 string, arg3 string) bool
    isRunLuaOnLns(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isStreamType(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    isType(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo) bool
    lowerAndOr(_env *LnsEnv, arg1 *Nodes_ExpOp2Node) *convWasm_Val
    lowerAnonymousFunc(_env *LnsEnv, arg1 *Nodes_DeclFuncNode) *convWasm_Val
    lowerApply(_env *LnsEnv, arg1 *Nodes_ApplyNode) *WasmBin_Ins
    lowerApplyGmatch(_env *LnsEnv, arg1 *Nodes_ApplyNode, arg2 *Nodes_ExpCallNode, arg3 *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins
    lowerArgAny(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 LnsInt) *WasmBin_Ins
    lowerArgAs(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 LnsInt, arg3 *Ast_TypeInfo) *WasmBin_Ins
    lowerArgs(_env *LnsEnv, arg1 LnsAny, arg2 *LnsList2_[*Ast_TypeInfo], arg3 *LnsList2_[*WasmBin_Ins]) *LnsList2_[*WasmBin_Ins]
    lowerAssign(_env *LnsEnv, arg1 *Nodes_Node, arg2 *convWasm_Val, arg3 *LnsList2_[*WasmBin_Ins])
    lowerBlock(_env *LnsEnv, arg1 *Nodes_BlockNode) *LnsList2_[*WasmBin_Ins]
    lowerBlockAt(_env *LnsEnv, arg1 *Nodes_BlockNode, arg2 LnsInt) *LnsList2_[*WasmBin_Ins]
    lowerCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerCallField(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 *Nodes_RefFieldNode, arg3 *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerCallRef(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 *Nodes_ExpRefNode, arg3 *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerCallSuper(_env *LnsEnv, arg1 *Nodes_ExpCallSuperNode)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerCallSuperCtor(_env *LnsEnv, arg1 *Nodes_ExpCallSuperCtorNode, arg2 *LnsList2_[*WasmBin_Ins])
    lowerCast(_env *LnsEnv, arg1 *Nodes_ExpCastNode) *convWasm_Val
    lowerCompare(_env *LnsEnv, arg1 *Nodes_Node, arg2 string, arg3 *convWasm_Val, arg4 *convWasm_Val) *WasmBin_Ins
    lowerCond(_env *LnsEnv, arg1 *Nodes_Node) *WasmBin_Ins
    lowerCondRetOf(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsList2_[*WasmBin_Ins])
    lowerDeclClass(_env *LnsEnv, arg1 *Nodes_DeclClassNode, arg2 *LnsList2_[*WasmBin_Ins])
    lowerDeclVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode, arg2 *LnsList2_[*WasmBin_Ins])
    lowerEnumVal(_env *LnsEnv, arg1 *Ast_EnumTypeInfo, arg2 *Ast_EnumValInfo) *convWasm_Val
    lowerEqVals(_env *LnsEnv, arg1 *convWasm_Val, arg2 *convWasm_Val) *WasmBin_Ins
    lowerExp(_env *LnsEnv, arg1 *Nodes_Node) *WasmBin_Ins
    lowerExpAs(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_TypeInfo) *WasmBin_Ins
    lowerExtModuleVal(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_TypeInfo) *WasmBin_Ins
    lowerFor(_env *LnsEnv, arg1 *Nodes_ForNode) *WasmBin_Ins
    lowerForeach(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_SymbolInfo, arg3 LnsAny, arg4 *Nodes_Node, arg5 *Nodes_BlockNode, arg6 bool) *WasmBin_Ins
    lowerGetField(_env *LnsEnv, arg1 *Nodes_GetFieldNode) *convWasm_Val
    lowerIf(_env *LnsEnv, arg1 *Nodes_IfNode) *WasmBin_Ins
    lowerIfMethodCall(_env *LnsEnv, arg1 *Nodes_Node, arg2 string, arg3 *LnsList2_[*WasmBin_Ins], arg4 LnsAny)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerIfUnwrap(_env *LnsEnv, arg1 *Nodes_IfUnwrapNode) *WasmBin_Ins
    lowerImportedModules(_env *LnsEnv)
    lowerListMethod(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 *Nodes_Node, arg3 string)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerLiteralColl(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny, arg3 bool) *convWasm_Val
    lowerLiteralMap(_env *LnsEnv, arg1 *Nodes_LiteralMapNode) *convWasm_Val
    lowerLiteralString(_env *LnsEnv, arg1 *Nodes_LiteralStringNode) *WasmBin_Ins
    lowerLuaCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 LnsInt, arg3 *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerLuaCallSub(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 LnsInt, arg3 LnsAny, arg4 *LnsList2_[*WasmBin_Ins], arg5 *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerMatch(_env *LnsEnv, arg1 *Nodes_MatchNode) *WasmBin_Ins
    lowerMemberRef(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_Node, arg3 string) *convWasm_Val
    lowerMethodCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 *Nodes_Node, arg3 string, arg4 *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerMulti(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsList2_[*WasmBin_Ins]) *LnsList2_[*convWasm_Val]
    lowerNew(_env *LnsEnv, arg1 *Nodes_ExpNewNode) *convWasm_Val
    lowerNewAlgeVal(_env *LnsEnv, arg1 *Nodes_NewAlgeValNode) *convWasm_Val
    lowerNilAcc(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_Node) *convWasm_Val
    lowerNilAccCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 *Nodes_Node)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerObjMethodCall(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Nodes_Node, arg3 string, arg4 *LnsList2_[*WasmBin_Ins], arg5 LnsAny)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerOp1(_env *LnsEnv, arg1 *Nodes_ExpOp1Node) *convWasm_Val
    lowerOp2(_env *LnsEnv, arg1 *Nodes_ExpOp2Node) *convWasm_Val
    lowerPrint(_env *LnsEnv, arg1 *Nodes_ExpCallNode) *WasmBin_Ins
    lowerRawVal(_env *LnsEnv, arg1 *Nodes_Node) *convWasm_Val
    lowerRawValList(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 *LnsList2_[*WasmBin_Ins]) *LnsList2_[*convWasm_Val]
    lowerRefField(_env *LnsEnv, arg1 *Nodes_RefFieldNode) *convWasm_Val
    lowerRefFieldSub(_env *LnsEnv, arg1 *Nodes_RefFieldNode) *convWasm_Val
    lowerRefItem(_env *LnsEnv, arg1 *Nodes_ExpRefItemNode) *convWasm_Val
    lowerReturn(_env *LnsEnv, arg1 *Nodes_ReturnNode, arg2 *LnsList2_[*WasmBin_Ins])
    lowerSetItem(_env *LnsEnv, arg1 *Nodes_ExpSetItemNode, arg2 *LnsList2_[*WasmBin_Ins])
    lowerSetMethod(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 *Nodes_Node, arg3 string)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerSetVal(_env *LnsEnv, arg1 *Nodes_ExpSetValNode, arg2 *LnsList2_[*WasmBin_Ins])
    lowerStaticCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 *Nodes_Node, arg3 string, arg4 *Ast_TypeInfo, arg5 *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerStmt(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsList2_[*WasmBin_Ins])
    lowerStmtExp(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsList2_[*WasmBin_Ins])
    lowerStmtList(_env *LnsEnv, arg1 *LnsList2_[*Nodes_Node]) *LnsList2_[*WasmBin_Ins]
    lowerStrMethod(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 *Nodes_Node, arg3 string)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerStreamMethod(_env *LnsEnv, arg1 *Nodes_ExpCallNode, arg2 *Nodes_Node, arg3 string)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo])
    lowerSwitch(_env *LnsEnv, arg1 *Nodes_SwitchNode) *WasmBin_Ins
    lowerUnwrap(_env *LnsEnv, arg1 *Nodes_ExpUnwrapNode) *convWasm_Val
    lowerUnwrapSet(_env *LnsEnv, arg1 *Nodes_UnwrapSetNode, arg2 *LnsList2_[*WasmBin_Ins])
    lowerVal(_env *LnsEnv, arg1 *Nodes_Node) *convWasm_Val
    lowerValList(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 *LnsList2_[*WasmBin_Ins]) *LnsList2_[*convWasm_Val]
    lowerValListSub(_env *LnsEnv, arg1 *Nodes_ExpListNode, arg2 *LnsList2_[*WasmBin_Ins], arg3 bool) *LnsList2_[*convWasm_Val]
    lowerWhen(_env *LnsEnv, arg1 *Nodes_WhenNode) *WasmBin_Ins
    luaFuncToClosure(_env *LnsEnv, arg1 *WasmBin_Ins, arg2 *Ast_TypeInfo) *WasmBin_Ins
    makeIfAdapter(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 *Ast_TypeInfo, arg3 LnsInt) LnsInt
    mapConvKey(_env *LnsEnv, arg1 *Ast_TypeInfo) string
    needToMapConv(_env *LnsEnv, arg1 *Ast_TypeInfo) bool
    newVal(_env *LnsEnv, arg1 *WasmBin_Ins, arg2 *Ast_TypeInfo) *convWasm_Val
    normType(_env *LnsEnv, arg1 *Ast_TypeInfo) *Ast_TypeInfo
    notSupport(_env *LnsEnv, arg1 *Nodes_Node)
    preDeclare(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsList2_[*WasmBin_Ins])
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
    ProcessContinue(_env *LnsEnv, arg1 *Nodes_ContinueNode, arg2 LnsAny)
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
    ProcessRefTypeFunc(_env *LnsEnv, arg1 *Nodes_RefTypeFuncNode, arg2 LnsAny)
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
    readSym(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_SymbolInfo) *WasmBin_Ins
    readSymSub(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_SymbolInfo) *WasmBin_Ins
    registerClass(_env *LnsEnv, arg1 *Nodes_DeclClassNode) *convWasm_ClassInfo
    registerEmbedLua(_env *LnsEnv) *WasmBin_Ins
    registerFunc(_env *LnsEnv, arg1 *Nodes_DeclFuncInfo, arg2 *Ast_TypeInfo) LnsInt
    registerImportAlias(_env *LnsEnv, arg1 *Nodes_RootNode)
    registerModuleVars(_env *LnsEnv, arg1 *Nodes_RootNode)
    registerModuleVarsIn(_env *LnsEnv, arg1 *LnsList2_[*Nodes_Node])
    resolveSym(_env *LnsEnv, arg1 *Ast_SymbolInfo) *Ast_SymbolInfo
    setupParams(_env *LnsEnv, arg1 *Nodes_DeclFuncInfo, arg2 *Ast_TypeInfo, arg3 *LnsList2_[*WasmBin_Ins])
    spillVal(_env *LnsEnv, arg1 *convWasm_Val, arg2 *LnsList2_[*WasmBin_Ins]) *convWasm_Val
    tempLocal(_env *LnsEnv, arg1 *WasmBin_ValType) LnsInt
    toAny(_env *LnsEnv, arg1 *convWasm_Val) *WasmBin_Ins
    toCond(_env *LnsEnv, arg1 *convWasm_Val) *WasmBin_Ins
    toMapVal(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_TypeInfo, arg3 *WasmBin_Ins) *WasmBin_Ins
    toStr(_env *LnsEnv, arg1 *convWasm_Val) *WasmBin_Ins
    writeSym(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_SymbolInfo, arg3 *WasmBin_Ins) *WasmBin_Ins
}
type convWasm_convFilter struct {
    Nodes_Filter
    streamName string
    embedLuaList *LnsList2_[*ConvWasm_EmbedLua]
    bindLnsMap *LnsMap2_[string,string]
    embedLuaGlobal LnsInt
    stream Lns_oStream
    processInfo *Ast_ProcessInfo
    builtinFuncs *Builtin_BuiltinFuncType
    moduleScope *Ast_Scope
    moduleTypeInfo *Ast_TypeInfo
    importAstList *LnsList2_[*AstInfo_ASTInfo]
    moduleScopeSet *LnsSet2_[*Ast_Scope]
    module2init *LnsMap2_[string,LnsInt]
    module2provide *LnsMap2_[string,*Ast_SymbolInfo]
    symAlias *LnsMap2_[*Ast_SymbolInfo,*Ast_SymbolInfo]
    lazySym2module *LnsMap2_[*Ast_SymbolInfo,string]
    mod *WasmBin_Module
    rt *WasmRt_Runtime
    ctx *convWasm_FuncCtx
    globalSym2loc *LnsMap2_[*Ast_SymbolInfo,LnsAny]
    func2index *LnsMap2_[*Ast_TypeInfo,LnsInt]
    func2decl *LnsMap2_[*Ast_TypeInfo,*Nodes_DeclFuncInfo]
    classType2decl *LnsMap2_[*Ast_TypeInfo,*Nodes_DeclClassNode]
    class2info *LnsMap2_[*Ast_TypeInfo,*convWasm_ClassInfo]
    classInfoList *LnsList2_[*convWasm_ClassInfo]
    serrInfo LnsAny
    if2info *LnsMap2_[*Ast_TypeInfo,*convWasm_IfInfo]
    ifInfoList *LnsList2_[*convWasm_IfInfo]
    algeType2decl *LnsMap2_[*Ast_TypeInfo,*Nodes_DeclAlgeNode]
    alge2info *LnsMap2_[*Ast_TypeInfo,*convWasm_AlgeInfo]
    algeInfoList *LnsList2_[*convWasm_AlgeInfo]
    enumType2decl *LnsMap2_[*Ast_TypeInfo,*Nodes_DeclEnumNode]
    declaredMethodSet *LnsSet2_[*Ast_TypeInfo]
    advertiseMethodSet *LnsSet2_[*Ast_TypeInfo]
    dynGetterNameSet *LnsSet2_[string]
    toMapDynIndex LnsInt
    toMapDynNode LnsAny
    adapterMap *LnsMap2_[string,LnsInt]
    extModule2global *LnsMap2_[*Ast_TypeInfo,LnsInt]
    algeBuildingMap *LnsMap2_[*Ast_TypeInfo,LnsInt]
    node2val *LnsMap2_[*Nodes_Node,*convWasm_Val]
    nilAccDoneSet *LnsSet2_[*Nodes_Node]
    funcSym2name *LnsMap2_[*Ast_SymbolInfo,string]
    prologue *LnsList2_[*WasmBin_Ins]
    stdoutGlobal LnsInt
    stderrGlobal LnsInt
    stdinGlobal LnsInt
    mainIndex LnsInt
    compiledFuncSet *LnsSet2_[LnsInt]
    FP convWasm_convFilterMtd
}
func convWasm_convFilter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*convWasm_convFilter).FP
}
func convWasm_convFilter_toSlice(slice []LnsAny) []*convWasm_convFilter {
    ret := make([]*convWasm_convFilter, len(slice))
    for index, val := range slice {
        ret[index] = val.(convWasm_convFilterDownCast).ToconvWasm_convFilter()
    }
    return ret
}
type convWasm_convFilterDownCast interface {
    ToconvWasm_convFilter() *convWasm_convFilter
}
func convWasm_convFilterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(convWasm_convFilterDownCast)
    if ok { return work.ToconvWasm_convFilter() }
    return nil
}
func (obj *convWasm_convFilter) ToconvWasm_convFilter() *convWasm_convFilter {
    return obj
}
func NewconvWasm_convFilter(_env *LnsEnv, arg1 string, arg2 Lns_oStream, arg3 *AstInfo_ASTInfo, arg4 *LnsList2_[*AstInfo_ASTInfo], arg5 *LnsList2_[*ConvWasm_EmbedLua], arg6 *LnsMap2_[string,string]) *convWasm_convFilter {
    obj := &convWasm_convFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvWasm_convFilter(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
// 352: DeclConstr
func (self *convWasm_convFilter) InitconvWasm_convFilter(_env *LnsEnv, streamName string,stream Lns_oStream,ast *AstInfo_ASTInfo,importAstList *LnsList2_[*AstInfo_ASTInfo],embedLuaList *LnsList2_[*ConvWasm_EmbedLua],bindLnsMap *LnsMap2_[string,string]) {
    self.InitNodes_Filter(_env, true, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env))
    self.streamName = streamName
    self.embedLuaList = embedLuaList
    self.bindLnsMap = bindLnsMap
    self.embedLuaGlobal = -1
    self.stream = stream
    self.processInfo = ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env).FP.Clone(_env)
    self.builtinFuncs = ast.FP.Get_builtinFunc(_env)
    self.moduleTypeInfo = ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env)
    self.moduleScope = Lns_unwrap( ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env)).(*Ast_Scope)
    self.importAstList = importAstList
    self.moduleScopeSet = NewLnsSet2_[*Ast_Scope]([]*Ast_Scope{})
    self.moduleScopeSet.Add(self.moduleScope)
    self.module2init = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    self.module2provide = NewLnsMap2_[string,*Ast_SymbolInfo]( map[string]*Ast_SymbolInfo{})
    self.symAlias = NewLnsMap2_[*Ast_SymbolInfo,*Ast_SymbolInfo]( map[*Ast_SymbolInfo]*Ast_SymbolInfo{})
    self.lazySym2module = NewLnsMap2_[*Ast_SymbolInfo,string]( map[*Ast_SymbolInfo]string{})
    var mod *WasmBin_Module
    mod = NewWasmBin_Module(_env)
    self.rt = NewWasmRt_Runtime(_env, mod)
    self.mod = mod
    var mainIndex LnsInt
    var mainFunc *WasmBin_Func
    mainIndex,mainFunc = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
    self.mainIndex = mainIndex
    self.mod.FP.SetFuncName(_env, mainIndex, "main")
    self.ctx = NewconvWasm_FuncCtx(_env, mainFunc, NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{}))
    self.compiledFuncSet = NewLnsSet2_[LnsInt]([]LnsInt{})
    self.globalSym2loc = NewLnsMap2_[*Ast_SymbolInfo,LnsAny]( map[*Ast_SymbolInfo]LnsAny{})
    self.func2index = NewLnsMap2_[*Ast_TypeInfo,LnsInt]( map[*Ast_TypeInfo]LnsInt{})
    self.func2decl = NewLnsMap2_[*Ast_TypeInfo,*Nodes_DeclFuncInfo]( map[*Ast_TypeInfo]*Nodes_DeclFuncInfo{})
    self.classType2decl = NewLnsMap2_[*Ast_TypeInfo,*Nodes_DeclClassNode]( map[*Ast_TypeInfo]*Nodes_DeclClassNode{})
    self.class2info = NewLnsMap2_[*Ast_TypeInfo,*convWasm_ClassInfo]( map[*Ast_TypeInfo]*convWasm_ClassInfo{})
    self.classInfoList = NewLnsList2_[*convWasm_ClassInfo]([]*convWasm_ClassInfo{})
    self.serrInfo = nil
    self.if2info = NewLnsMap2_[*Ast_TypeInfo,*convWasm_IfInfo]( map[*Ast_TypeInfo]*convWasm_IfInfo{})
    self.ifInfoList = NewLnsList2_[*convWasm_IfInfo]([]*convWasm_IfInfo{})
    self.algeType2decl = NewLnsMap2_[*Ast_TypeInfo,*Nodes_DeclAlgeNode]( map[*Ast_TypeInfo]*Nodes_DeclAlgeNode{})
    self.alge2info = NewLnsMap2_[*Ast_TypeInfo,*convWasm_AlgeInfo]( map[*Ast_TypeInfo]*convWasm_AlgeInfo{})
    self.algeInfoList = NewLnsList2_[*convWasm_AlgeInfo]([]*convWasm_AlgeInfo{})
    self.enumType2decl = NewLnsMap2_[*Ast_TypeInfo,*Nodes_DeclEnumNode]( map[*Ast_TypeInfo]*Nodes_DeclEnumNode{})
    self.declaredMethodSet = NewLnsSet2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    self.advertiseMethodSet = NewLnsSet2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    self.dynGetterNameSet = NewLnsSet2_[string]([]string{})
    self.toMapDynIndex = -1
    self.toMapDynNode = nil
    self.adapterMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    self.extModule2global = NewLnsMap2_[*Ast_TypeInfo,LnsInt]( map[*Ast_TypeInfo]LnsInt{})
    self.algeBuildingMap = NewLnsMap2_[*Ast_TypeInfo,LnsInt]( map[*Ast_TypeInfo]LnsInt{})
    self.node2val = NewLnsMap2_[*Nodes_Node,*convWasm_Val]( map[*Nodes_Node]*convWasm_Val{})
    self.nilAccDoneSet = NewLnsSet2_[*Nodes_Node]([]*Nodes_Node{})
    self.funcSym2name = NewLnsMap2_[*Ast_SymbolInfo,string]( map[*Ast_SymbolInfo]string{})
    self.prologue = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    self.stdoutGlobal = -1
    self.stderrGlobal = -1
    self.stdinGlobal = -1
}


func Lns_convWasm_init(_env *LnsEnv) {
    if init_convWasm { return }
    init_convWasm = true
    convWasm__mod__ = "@lune.@base.@convWasm"
    Lns_InitMod()
    Lns_Ast_init(_env)
    Lns_Nodes_init(_env)
    Lns_Util_init(_env)
    Lns_AstInfo_init(_env)
    Lns_Builtin_init(_env)
    Lns_WasmBin_init(_env)
    Lns_WasmRt_init(_env)
    convWasm_strCmpOpMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{"~=":WasmBin_opI32Ne,"<":WasmBin_opI32LtS,"<=":WasmBin_opI32LeS,">":WasmBin_opI32GtS,">=":WasmBin_opI32GeS,})
    convWasm_i64CmpOpMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{"==":WasmBin_opI64Eq,"~=":WasmBin_opI64Ne,"<":WasmBin_opI64LtS,"<=":WasmBin_opI64LeS,">":WasmBin_opI64GtS,">=":WasmBin_opI64GeS,})
    convWasm_f64CmpOpMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{"==":WasmBin_opF64Eq,"~=":WasmBin_opF64Ne,"<":WasmBin_opF64Lt,"<=":WasmBin_opF64Le,">":WasmBin_opF64Gt,">=":WasmBin_opF64Ge,})
    convWasm_mappingMethodSet = NewLnsSet2_[string](Lns_2DDDGen[string]("_toMap", "_fromMap", "_fromStem"))
}
func init() {
    init_convWasm = false
}
// 248: decl @lune.@base.@convWasm.FuncCtx.addLocal
func (self *convWasm_FuncCtx) AddLocal(_env *LnsEnv, valType *WasmBin_ValType) LnsInt {
    return self._func.FP.AddLocal(_env, valType)
}
// 416: decl @lune.@base.@convWasm.convFilter.error
func (self *convWasm_convFilter) error(_env *LnsEnv, node *Nodes_Node,mess string) {
    Util_err(_env, _env.GetVM().String_format("%s: wasm: %s", Lns_2DDD(node.FP.Get_pos(_env).GetDisplayTxt(_env), mess)))
}
// 421: decl @lune.@base.@convWasm.convFilter.notSupport
func (self *convWasm_convFilter) notSupport(_env *LnsEnv, node *Nodes_Node) {
    self.FP.error(_env, node, _env.GetVM().String_format("not support -- %s", Lns_2DDD(Nodes_getNodeKindName(_env, node.FP.Get_kind(_env)))))
}
// 428: decl @lune.@base.@convWasm.convFilter.normType
func (self *convWasm_convFilter) normType(_env *LnsEnv, typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    var work *Ast_TypeInfo
    work = typeInfo.FP.Get_aliasSrc(_env).FP.Get_srcTypeInfo(_env)
    if work.FP.Get_nilable(_env){
        work = work.FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env).FP.Get_srcTypeInfo(_env)
    }
    if work.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext{
        return Ast_builtinTypeStem
    }
    return work
}
// 442: decl @lune.@base.@convWasm.convFilter.getExtModuleType
func (self *convWasm_convFilter) getExtModuleType(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny {
    var work *Ast_TypeInfo
    work = typeInfo.FP.Get_aliasSrc(_env).FP.Get_srcTypeInfo(_env).FP.Get_nonnilableType(_env)
    if work.FP.Get_kind(_env) == Ast_TypeInfoKind__ExtModule{
        return work
    }
    return nil
}
// 456: decl @lune.@base.@convWasm.convFilter.lowerExtModuleVal
func (self *convWasm_convFilter) lowerExtModuleVal(_env *LnsEnv, node *Nodes_Node,modType *Ast_TypeInfo) *WasmBin_Ins {
    var modGlobal LnsInt
    {
        _work := self.extModule2global.Get(modType)
        if !Lns_IsNil( _work ) {
            work := _work.(LnsInt)
            modGlobal = work
        } else {
            modGlobal = self.mod.FP.AddGlobal(_env, WasmBin_typeAnyRef, WasmBin_refNull(_env, WasmBin_htAny))
            self.extModule2global.Set(modType,modGlobal)
        }
    }
    var normalType *Ast_NormalTypeInfo
    
    {
        _normalType := Ast_NormalTypeInfoDownCastF(modType.FP)
        if _normalType == nil{
            self.FP.error(_env, node, _env.GetVM().String_format("not support module -- %s", Lns_2DDD(modType.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            normalType = _normalType.(*Ast_NormalTypeInfo)
        }
    }
    var args LnsInt
    args = self.FP.tempLocal(_env, self.rt.FP.Get_listType(_env))
    return WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refIsNull(_env, WasmBin_globalGet(_env, modGlobal)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, args, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 1))))), WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, args), self.rt.FP.StrLiteral(_env, normalType.FP.Get_requirePath(_env))))), WasmBin_globalSet(_env, modGlobal, self.rt.FP.ListItem(_env, WasmBin_call(_env, self.rt.FP.FnLuaCall(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, WasmRt_luaRequire), convWasm_lg_2_(_env, args)))), 0)))), nil, nil), WasmBin_globalGet(_env, modGlobal))))
}
// 479: decl @lune.@base.@convWasm.convFilter.isLuaval
func (self *convWasm_convFilter) isLuaval(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    var work *Ast_TypeInfo
    work = typeInfo.FP.Get_aliasSrc(_env).FP.Get_srcTypeInfo(_env)
    if work.FP.Get_nilable(_env){
        work = work.FP.Get_nonnilableType(_env).FP.Get_aliasSrc(_env).FP.Get_srcTypeInfo(_env)
    }
    return work.FP.Get_kind(_env) == Ast_TypeInfoKind__Ext
}
// 488: decl @lune.@base.@convWasm.convFilter.getGenSrc
func (self *convWasm_convFilter) getGenSrc(_env *LnsEnv, typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    var work *Ast_TypeInfo
    work = self.FP.normType(_env, typeInfo)
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = 10
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            var src *Ast_TypeInfo
            src = work.FP.Get_genSrcTypeInfo(_env).FP.Get_aliasSrc(_env).FP.Get_srcTypeInfo(_env)
            if src == work{
                break
            }
            work = src
        }
    }
    return work
}
// 502: decl @lune.@base.@convWasm.convFilter.isType
func (self *convWasm_convFilter) isType(_env *LnsEnv, typeInfo *Ast_TypeInfo,builtin *Ast_TypeInfo) bool {
    return self.FP.normType(_env, typeInfo).FP.Equals(_env, self.processInfo, builtin, nil, nil)
}
// 512: decl @lune.@base.@convWasm.convFilter.getRep
func (self *convWasm_convFilter) getRep(_env *LnsEnv, typeInfo *Ast_TypeInfo) *WasmBin_ValType {
    var orgType *Ast_TypeInfo
    orgType = typeInfo.FP.Get_aliasSrc(_env).FP.Get_srcTypeInfo(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( orgType.FP.Get_nilable(_env)) &&
        _env.SetStackVal( orgType != Ast_builtinTypeNil) &&
        _env.SetStackVal( orgType.FP.Get_nonnilableType(_env) != orgType) ).(bool)){
        var rep *WasmBin_ValType
        rep = self.FP.getRep(_env, orgType.FP.Get_nonnilableType(_env))
        if convWasm_isScalar_15_(_env, rep){
            return WasmBin_typeAnyRef
        }
        return rep
    }
    var work *Ast_TypeInfo
    work = self.FP.normType(_env, orgType)
    if work.FP.Equals(_env, self.processInfo, Ast_builtinTypeString, nil, nil){
        return self.rt.FP.Get_strType(_env)
    }
    if _switch0 := work.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Prim {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( work.FP.Equals(_env, self.processInfo, Ast_builtinTypeInt, nil, nil)) ||
            _env.SetStackVal( work.FP.Equals(_env, self.processInfo, Ast_builtinTypeChar, nil, nil)) ).(bool){
            return WasmBin_typeI64
        }
        if work.FP.Equals(_env, self.processInfo, Ast_builtinTypeReal, nil, nil){
            return WasmBin_typeF64
        }
        if work.FP.Equals(_env, self.processInfo, Ast_builtinTypeBool, nil, nil){
            return WasmBin_typeI32
        }
        return WasmBin_typeAnyRef
    } else if _switch0 == Ast_TypeInfoKind__Enum {
        {
            _enumType := Ast_EnumTypeInfoDownCastF(work.FP)
            if !Lns_IsNil( _enumType ) {
                enumType := _enumType.(*Ast_EnumTypeInfo)
                return self.FP.getRep(_env, enumType.FP.Get_valTypeInfo(_env))
            }
        }
    } else if _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array || _switch0 == Ast_TypeInfoKind__DDD || _switch0 == Ast_TypeInfoKind__Tuple || _switch0 == Ast_TypeInfoKind__Box {
        return self.rt.FP.Get_listType(_env)
    } else if _switch0 == Ast_TypeInfoKind__Map || _switch0 == Ast_TypeInfoKind__Set {
        return self.rt.FP.Get_mapType(_env)
    } else if _switch0 == Ast_TypeInfoKind__Class {
        {
            _info := self.FP.getClassInfo(_env, work)
            if !Lns_IsNil( _info ) {
                info := _info.(*convWasm_ClassInfo)
                return WasmBin_refType(_env, info.FP.Get_structType(_env))
            }
        }
        return self.rt.FP.Get_objType(_env)
    } else if _switch0 == Ast_TypeInfoKind__IF {
        return self.rt.FP.Get_objType(_env)
    } else if _switch0 == Ast_TypeInfoKind__Func || _switch0 == Ast_TypeInfoKind__Form || _switch0 == Ast_TypeInfoKind__FormFunc {
        return WasmBin_refType(_env, self.FP.getFuncCloInfo(_env, work).FP.Get_cloType(_env))
    } else if _switch0 == Ast_TypeInfoKind__Alge {
        {
            _baseType := self.algeBuildingMap.Get(self.FP.getGenSrc(_env, work))
            if !Lns_IsNil( _baseType ) {
                baseType := _baseType.(LnsInt)
                return WasmBin_refType(_env, baseType)
            }
        }
        {
            _info := self.FP.getAlgeInfo(_env, self.FP.getGenSrc(_env, work))
            if !Lns_IsNil( _info ) {
                info := _info.(*convWasm_AlgeInfo)
                return WasmBin_refType(_env, info.FP.Get_baseType(_env))
            }
        }
    }
    return WasmBin_typeAnyRef
}
// 580: decl @lune.@base.@convWasm.convFilter.getRepList
func (self *convWasm_convFilter) getRepList(_env *LnsEnv, typeList *LnsList2_[*Ast_TypeInfo]) *LnsList2_[*WasmBin_ValType] {
    var list *LnsList2_[*WasmBin_ValType]
    list = NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{})
    for _, _typeInfo := range( typeList.Items ) {
        typeInfo := _typeInfo
        list.Insert(self.FP.getRep(_env, typeInfo))
    }
    return list
}
// 589: decl @lune.@base.@convWasm.convFilter.getFuncCloInfo
func (self *convWasm_convFilter) getFuncCloInfo(_env *LnsEnv, funcType *Ast_TypeInfo) *WasmRt_CloInfo {
    var work *Ast_TypeInfo
    work = self.FP.normType(_env, funcType)
    return self.rt.FP.GetCloInfo(_env, self.FP.getRepList(_env, work.FP.Get_argTypeInfoList(_env)), self.FP.getRepList(_env, work.FP.Get_retTypeInfoList(_env)))
}
// 596: decl @lune.@base.@convWasm.convFilter.getMethodFuncType
func (self *convWasm_convFilter) getMethodFuncType(_env *LnsEnv, methodType *Ast_TypeInfo) LnsInt {
    var params *LnsList2_[*WasmBin_ValType]
    params = NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_objType(_env)))
    for _, _param := range( self.FP.getRepList(_env, methodType.FP.Get_argTypeInfoList(_env)).Items ) {
        param := _param
        params.Insert(param)
    }
    return self.mod.FP.GetFuncType(_env, params, self.FP.getRepList(_env, methodType.FP.Get_retTypeInfoList(_env)))
}
// 605: decl @lune.@base.@convWasm.convFilter.getDefaultVal
func (self *convWasm_convFilter) getDefaultVal(_env *LnsEnv, rep *WasmBin_ValType) *WasmBin_Ins {
    if _switch0 := rep.FP.Get_kind(_env); _switch0 == WasmBin_ValKind__I64 {
        return WasmBin_i64Const(_env, 0)
    } else if _switch0 == WasmBin_ValKind__F64 {
        return WasmBin_f64Const(_env, 0.0)
    } else if _switch0 == WasmBin_ValKind__I32 {
        return WasmBin_i32Const(_env, 0)
    } else if _switch0 == WasmBin_ValKind__Ref {
        return WasmBin_refNull(_env, rep.FP.Get_heapType(_env))
    }
// insert a dummy
    return nil
}
// 633: decl @lune.@base.@convWasm.convFilter.getClassInfo
func (self *convWasm_convFilter) getClassInfo(_env *LnsEnv, typeInfo *Ast_TypeInfo) LnsAny {
    var classType *Ast_TypeInfo
    classType = self.FP.getGenSrc(_env, typeInfo)
    {
        _info := self.class2info.Get(classType)
        if !Lns_IsNil( _info ) {
            info := _info.(*convWasm_ClassInfo)
            return info
        }
    }
    var declNode *Nodes_DeclClassNode
    
    {
        _declNode := self.classType2decl.Get(classType)
        if _declNode == nil{
            return nil
        } else {
            declNode = _declNode.(*Nodes_DeclClassNode)
        }
    }
    return self.FP.registerClass(_env, declNode)
}
// 645: decl @lune.@base.@convWasm.convFilter.registerClass
func (self *convWasm_convFilter) registerClass(_env *LnsEnv, declNode *Nodes_DeclClassNode) *convWasm_ClassInfo {
    var classType *Ast_TypeInfo
    classType = self.FP.getGenSrc(_env, declNode.FP.Get_expType(_env))
    {
        _info := self.class2info.Get(classType)
        if !Lns_IsNil( _info ) {
            info := _info.(*convWasm_ClassInfo)
            return info
        }
    }
    var parent LnsAny
    parent = nil
    var superStruct LnsInt
    superStruct = self.rt.FP.Get_tObj(_env)
    var superVt LnsInt
    superVt = self.rt.FP.Get_tVT(_env)
    if classType.FP.HasBase(_env){
        var parentInfo *convWasm_ClassInfo
        
        {
            _parentInfo := self.FP.getClassInfo(_env, classType.FP.Get_baseTypeInfo(_env))
            if _parentInfo == nil{
                self.FP.error(_env, &declNode.Nodes_Node, _env.GetVM().String_format("not found base class -- %s", Lns_2DDD(classType.FP.Get_baseTypeInfo(_env).FP.GetTxt(_env, nil, nil, nil))))
            } else {
                parentInfo = _parentInfo.(*convWasm_ClassInfo)
            }
        }
        parent = parentInfo
        superStruct = parentInfo.FP.Get_structType(_env)
        superVt = parentInfo.FP.Get_vtType(_env)
    }
    var structType LnsInt
    structType = self.mod.FP.ReserveType(_env)
    var vtType LnsInt
    vtType = self.mod.FP.ReserveType(_env)
    var vtGlobal LnsInt
    vtGlobal = self.mod.FP.AddGlobal(_env, WasmBin_refType(_env, vtType), WasmBin_refNull(_env, vtType))
    var info *convWasm_ClassInfo
    info = NewconvWasm_ClassInfo(_env, classType, structType, vtType, vtGlobal, parent)
    self.class2info.Set(classType,info)
    self.classInfoList.Insert(info)
    for _, _memberNode := range( declNode.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode
        if memberNode.FP.Get_staticFlag(_env){
            var symbol *Ast_SymbolInfo
            symbol = memberNode.FP.Get_symbolInfo(_env).FP.GetOrg(_env)
            var rep *WasmBin_ValType
            rep = self.FP.getRep(_env, symbol.FP.Get_typeInfo(_env))
            self.globalSym2loc.Set(symbol,&convWasm_VarLoc__Global{self.mod.FP.AddGlobal(_env, rep, self.FP.getDefaultVal(_env, rep))})
        } else { 
            var name string
            name = memberNode.FP.Get_name(_env).Txt
            info.FP.Get_fieldNames(_env).Insert(name)
            info.FP.Get_fieldReps(_env).Insert(self.FP.getRep(_env, memberNode.FP.Get_symbolInfo(_env).FP.Get_typeInfo(_env)))
            info.FP.Get_fieldTypes(_env).Insert(memberNode.FP.Get_symbolInfo(_env).FP.Get_typeInfo(_env))
            info.FP.Get_name2field(_env).Set(name,info.FP.Get_fieldNames(_env).Len() + 1)
        }
    }
    var autoMethodSet *LnsSet2_[string]
    autoMethodSet = NewLnsSet2_[string]([]string{})
    for _, _memberNode := range( declNode.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode
        if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
            if memberNode.FP.Get_getterMode(_env) != Ast_AccessMode__None{
                autoMethodSet.Add("get_" + memberNode.FP.Get_name(_env).Txt)
            }
            if memberNode.FP.Get_setterMode(_env) != Ast_AccessMode__None{
                autoMethodSet.Add("set_" + memberNode.FP.Get_name(_env).Txt)
            }
        }
    }
    var scope *Ast_Scope
    scope = Lns_unwrap( classType.FP.Get_scope(_env)).(*Ast_Scope)
    {
        __forsortCollection0 := scope.FP.Get_symbol2SymbolInfoMap(_env)
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, _name := range( __forsortSorted0.Items ) {
            symbol := __forsortCollection0.Items[ _name ]
            name := _name
            var methodType *Ast_TypeInfo
            methodType = symbol.FP.Get_typeInfo(_env)
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( Lns_op_not(self.declaredMethodSet.Has(methodType))) &&
                    _env.SetStackVal( Lns_op_not(autoMethodSet.Has(name))) ).(bool))) ||
                _env.SetStackVal( convWasm_mappingMethodSet.Has(name)) ).(bool){
            } else if convWasm_isMethodSym_16_(_env, symbol){
                var impl LnsInt
                impl = -1
                var slotImpl LnsInt
                slotImpl = -1
                if Lns_op_not(methodType.FP.Get_abstractFlag(_env)){
                    var funcType LnsInt
                    funcType = self.FP.getMethodFuncType(_env, methodType)
                    var work LnsInt
                    work,_ = self.mod.FP.AddFuncWithType(_env, funcType, 1 + methodType.FP.Get_argTypeInfoList(_env).Len())
                    impl = work
                    slotImpl = work
                    self.mod.FP.SetFuncName(_env, impl, _env.GetVM().String_format("%s.%s", Lns_2DDD(classType.FP.GetTxt(_env, nil, nil, nil), name)))
                    self.func2index.Set(methodType,impl)
                    {
                        _slotIndex := info.FP.Get_name2slot(_env).Get(name)
                        if !Lns_IsNil( _slotIndex ) {
                            slotIndex := _slotIndex.(LnsInt)
                            var slot *convWasm_SlotInfo
                            slot = info.FP.Get_slotList(_env).GetAt(slotIndex)
                            if slot.FP.Get_funcType(_env) != funcType{
                                slotImpl = self.FP.makeIfAdapter(_env, slot.FP.Get_methodType(_env), methodType, impl)
                            }
                        }
                    }
                }
                {
                    _slotIndex := info.FP.Get_name2slot(_env).Get(name)
                    if !Lns_IsNil( _slotIndex ) {
                        slotIndex := _slotIndex.(LnsInt)
                        info.FP.Get_slotList(_env).GetAt(slotIndex).FP.Set_impl(_env, slotImpl)
                    } else {
                        info.FP.Get_slotList(_env).Insert(NewconvWasm_SlotInfo(_env, name, self.FP.getMethodFuncType(_env, methodType), methodType, impl))
                        info.FP.Get_name2slot(_env).Set(name,info.FP.Get_slotList(_env).Len())
                    }
                }
            } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( symbol.FP.Get_kind(_env) == Ast_SymbolKind__Mtd) &&
                _env.SetStackVal( symbol.FP.Get_staticFlag(_env)) ).(bool)){
                var index LnsInt
                index,_ = self.mod.FP.AddFuncWithType(_env, self.FP.getFuncCloInfo(_env, methodType).FP.Get_funcType(_env), 1 + methodType.FP.Get_argTypeInfoList(_env).Len())
                self.func2index.Set(methodType,index)
                self.mod.FP.SetFuncName(_env, index, _env.GetVM().String_format("%s.%s", Lns_2DDD(classType.FP.GetTxt(_env, nil, nil, nil), name)))
            }
        }
    }
    for _, _advertiseInfo := range( declNode.FP.Get_advertiseList(_env).Items ) {
        advertiseInfo := _advertiseInfo
        var memberType *Ast_TypeInfo
        memberType = advertiseInfo.FP.Get_member(_env).FP.Get_expType(_env)
        var mbrScope *Ast_Scope
        mbrScope = Lns_unwrap( self.FP.getGenSrc(_env, memberType).FP.Get_scope(_env)).(*Ast_Scope)
        for _, _mtdName := range( Ast_getAllMethodName(_env, self.FP.getGenSrc(_env, memberType), Ast_MethodKind__Object).FP.Get_list(_env).Items ) {
            mtdName := _mtdName.(string)
            var child *Ast_TypeInfo
            child = Lns_unwrap( mbrScope.FP.GetTypeInfoField(_env, mtdName, true, mbrScope, Ast_ScopeAccess__Normal)).(*Ast_TypeInfo)
            if child.FP.Get_accessMode(_env) == Ast_AccessMode__Pri{
                continue
            }
            var name string
            name = advertiseInfo.FP.Get_prefix(_env) + mtdName
            var methodType *Ast_TypeInfo
            methodType = child
            {
                _symbol := scope.FP.Get_symbol2SymbolInfoMap(_env).Get(name)
                if !Lns_IsNil( _symbol ) {
                    symbol := _symbol.(*Ast_SymbolInfo)
                    var ownMethod bool
                    ownMethod = self.FP.getGenSrc(_env, symbol.FP.Get_typeInfo(_env).FP.Get_parentInfo(_env)) == classType
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( ownMethod) &&
                        _env.SetStackVal( self.declaredMethodSet.Has(symbol.FP.Get_typeInfo(_env))) ||
                        _env.SetStackVal( autoMethodSet.Has(name)) ).(bool){
                        continue
                    }
                    methodType = symbol.FP.Get_typeInfo(_env)
                }
            }
            var funcType LnsInt
            {
                _slotIndex := info.FP.Get_name2slot(_env).Get(name)
                if !Lns_IsNil( _slotIndex ) {
                    slotIndex := _slotIndex.(LnsInt)
                    funcType = info.FP.Get_slotList(_env).GetAt(slotIndex).FP.Get_funcType(_env)
                } else {
                    funcType = self.FP.getMethodFuncType(_env, methodType)
                }
            }
            var impl LnsInt
            impl,_ = self.mod.FP.AddFuncWithType(_env, funcType, 1 + methodType.FP.Get_argTypeInfoList(_env).Len())
            self.mod.FP.SetFuncName(_env, impl, _env.GetVM().String_format("%s.%s", Lns_2DDD(classType.FP.GetTxt(_env, nil, nil, nil), name)))
            {
                _slotIndex := info.FP.Get_name2slot(_env).Get(name)
                if !Lns_IsNil( _slotIndex ) {
                    slotIndex := _slotIndex.(LnsInt)
                    info.FP.Get_slotList(_env).GetAt(slotIndex).FP.Set_impl(_env, impl)
                } else {
                    info.FP.Get_slotList(_env).Insert(NewconvWasm_SlotInfo(_env, name, funcType, methodType, impl))
                    info.FP.Get_name2slot(_env).Set(name,info.FP.Get_slotList(_env).Len())
                }
            }
            info.FP.Get_advertiseList(_env).Insert(NewconvWasm_AdvertiseMethod(_env, advertiseInfo.FP.Get_member(_env).FP.Get_name(_env).Txt, memberType, mtdName, methodType, impl))
            self.advertiseMethodSet.Add(methodType)
        }
    }
    {
        _initSym := scope.FP.Get_symbol2SymbolInfoMap(_env).Get("__init")
        if !Lns_IsNil( _initSym ) {
            initSym := _initSym.(*Ast_SymbolInfo)
            var ctorType LnsInt
            ctorType = self.FP.getMethodFuncType(_env, initSym.FP.Get_typeInfo(_env))
            var ctor LnsInt
            ctor,_ = self.mod.FP.AddFuncWithType(_env, ctorType, 1 + initSym.FP.Get_typeInfo(_env).FP.Get_argTypeInfoList(_env).Len())
            self.func2index.Set(initSym.FP.Get_typeInfo(_env),ctor)
            self.mod.FP.SetFuncName(_env, ctor, _env.GetVM().String_format("%s.__init", Lns_2DDD(classType.FP.GetTxt(_env, nil, nil, nil))))
            info.FP.Set_ctorFunc(_env, ctor)
            info.FP.Set_ctorFuncType(_env, ctorType)
        }
    }
    for _, _ifType := range( classType.FP.Get_interfaceList(_env).Items ) {
        ifType := _ifType
        self.FP.addInterface(_env, info, ifType)
    }
    var fields *LnsList2_[*WasmBin_FieldType]
    fields = NewLnsList2_[*WasmBin_FieldType](Lns_2DDDGen[*WasmBin_FieldType](WasmBin_field(_env, WasmBin_typeI32, true), WasmBin_field(_env, WasmBin_refType(_env, vtType), false)))
    for _, _rep := range( info.FP.Get_fieldReps(_env).Items ) {
        rep := _rep
        fields.Insert(WasmBin_field(_env, rep, true))
    }
    self.mod.FP.DefineStructType(_env, structType, superStruct, fields)
    var vtFields *LnsList2_[*WasmBin_FieldType]
    vtFields = NewLnsList2_[*WasmBin_FieldType](Lns_2DDDGen[*WasmBin_FieldType](WasmBin_field(_env, self.rt.FP.Get_anyArrType(_env), false)))
    for _, _slot := range( info.FP.Get_slotList(_env).Items ) {
        slot := _slot
        vtFields.Insert(WasmBin_field(_env, WasmBin_refType(_env, slot.FP.Get_funcType(_env)), false))
    }
    self.mod.FP.DefineStructType(_env, vtType, superVt, vtFields)
    return info
}
// 819: decl @lune.@base.@convWasm.convFilter.getSerrInfo
func (self *convWasm_convFilter) getSerrInfo(_env *LnsEnv) *convWasm_ClassInfo {
    {
        _info := self.serrInfo
        if !Lns_IsNil( _info ) {
            info := _info.(*convWasm_ClassInfo)
            return info
        }
    }
    var erType *Ast_TypeInfo
    erType = self.FP.getGenSrc(_env, self.builtinFuncs.Lns___serr.FP.Get_retTypeInfoList(_env).GetAt(1))
    var structType LnsInt
    structType = self.mod.FP.ReserveType(_env)
    var vtType LnsInt
    vtType = self.mod.FP.ReserveType(_env)
    var vtGlobal LnsInt
    vtGlobal = self.mod.FP.AddGlobal(_env, WasmBin_refType(_env, vtType), WasmBin_refNull(_env, vtType))
    var info *convWasm_ClassInfo
    info = NewconvWasm_ClassInfo(_env, erType, structType, vtType, vtGlobal, nil)
    info.FP.Get_fieldNames(_env).Insert("txt")
    info.FP.Get_fieldReps(_env).Insert(self.rt.FP.Get_strType(_env))
    info.FP.Get_fieldTypes(_env).Insert(Ast_builtinTypeString)
    info.FP.Get_name2field(_env).Set("txt",2)
    var methodType *Ast_TypeInfo
    methodType = Lns_unwrap( self.FP.getIfMethodType(_env, erType, "get_txt")).(*Ast_TypeInfo)
    var funcType LnsInt
    funcType = self.FP.getMethodFuncType(_env, methodType)
    var impl LnsInt
    var _func *WasmBin_Func
    impl,_func = self.mod.FP.AddFuncWithType(_env, funcType, 1)
    self.mod.FP.SetFuncName(_env, impl, "LnsErr.get_txt")
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structGet(_env, structType, 2, WasmBin_refCast(_env, structType, convWasm_lg_2_(_env, 0))))))
    info.FP.Get_slotList(_env).Insert(NewconvWasm_SlotInfo(_env, "get_txt", funcType, methodType, impl))
    info.FP.Get_name2slot(_env).Set("get_txt",1)
    self.FP.addInterface(_env, info, erType)
    self.mod.FP.DefineStructType(_env, structType, self.rt.FP.Get_tObj(_env), NewLnsList2_[*WasmBin_FieldType](Lns_2DDDGen[*WasmBin_FieldType](WasmBin_field(_env, WasmBin_typeI32, true), WasmBin_field(_env, WasmBin_refType(_env, vtType), false), WasmBin_field(_env, self.rt.FP.Get_strType(_env), true))))
    self.mod.FP.DefineStructType(_env, vtType, self.rt.FP.Get_tVT(_env), NewLnsList2_[*WasmBin_FieldType](Lns_2DDDGen[*WasmBin_FieldType](WasmBin_field(_env, self.rt.FP.Get_anyArrType(_env), false), WasmBin_field(_env, WasmBin_refType(_env, funcType), false))))
    self.classInfoList.Insert(info)
    self.serrInfo = info
    return info
}
// 854: decl @lune.@base.@convWasm.convFilter.addInterface
func (self *convWasm_convFilter) addInterface(_env *LnsEnv, info *convWasm_ClassInfo,ifType *Ast_TypeInfo) {
    var work *Ast_TypeInfo
    work = self.FP.getGenSrc(_env, ifType)
    for _, _registered := range( info.FP.Get_ifList(_env).Items ) {
        registered := _registered
        if registered == work{
            return 
        }
    }
    info.FP.Get_ifList(_env).Insert(work)
    for _, _parentIf := range( work.FP.Get_interfaceList(_env).Items ) {
        parentIf := _parentIf
        self.FP.addInterface(_env, info, parentIf)
    }
}
// 868: decl @lune.@base.@convWasm.convFilter.collectIfMethods
func (self *convWasm_convFilter) collectIfMethods(_env *LnsEnv, ifType *Ast_TypeInfo,_map *LnsMap2_[string,*Ast_TypeInfo]) {
    for _, _parentIf := range( ifType.FP.Get_interfaceList(_env).Items ) {
        parentIf := _parentIf
        self.FP.collectIfMethods(_env, self.FP.getGenSrc(_env, parentIf), _map)
    }
    {
        _scope := ifType.FP.Get_scope(_env)
        if !Lns_IsNil( _scope ) {
            scope := _scope.(*Ast_Scope)
            for _name, _symbol := range( scope.FP.Get_symbol2SymbolInfoMap(_env).Items ) {
                name := _name
                symbol := _symbol
                if convWasm_isMethodSym_16_(_env, symbol){
                    _map.Set(name,symbol.FP.Get_typeInfo(_env))
                }
            }
        }
    }
}
// 881: decl @lune.@base.@convWasm.convFilter.getIfInfo
func (self *convWasm_convFilter) getIfInfo(_env *LnsEnv, ifType *Ast_TypeInfo) *convWasm_IfInfo {
    var work *Ast_TypeInfo
    work = self.FP.getGenSrc(_env, ifType)
    {
        _info := self.if2info.Get(work)
        if !Lns_IsNil( _info ) {
            info := _info.(*convWasm_IfInfo)
            return info
        }
    }
    var methodMap *LnsMap2_[string,*Ast_TypeInfo]
    methodMap = NewLnsMap2_[string,*Ast_TypeInfo]( map[string]*Ast_TypeInfo{})
    self.FP.collectIfMethods(_env, work, methodMap)
    var names *LnsList2_[string]
    names = NewLnsList2_[string]([]string{})
    var funcTypes *LnsList2_[LnsInt]
    funcTypes = NewLnsList2_[LnsInt]([]LnsInt{})
    var name2index *LnsMap2_[string,LnsInt]
    name2index = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    var fields *LnsList2_[*WasmBin_FieldType]
    fields = NewLnsList2_[*WasmBin_FieldType]([]*WasmBin_FieldType{})
    {
        __forsortCollection0 := methodMap
        __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
        __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
        for _, _name := range( __forsortSorted0.Items ) {
            methodType := __forsortCollection0.Items[ _name ]
            name := _name
            names.Insert(name)
            var funcType LnsInt
            funcType = self.FP.getMethodFuncType(_env, methodType)
            funcTypes.Insert(LnsInt(funcType))
            name2index.Set(name,names.Len() - 1)
            fields.Insert(WasmBin_field(_env, WasmBin_refType(_env, funcType), false))
        }
    }
    var itType LnsInt
    itType = self.mod.FP.AddStructType(_env, nil, fields)
    var info *convWasm_IfInfo
    info = NewconvWasm_IfInfo(_env, work, self.ifInfoList.Len(), itType, names, funcTypes, name2index)
    self.if2info.Set(work,info)
    self.ifInfoList.Insert(info)
    return info
}
// 907: decl @lune.@base.@convWasm.convFilter.getFieldIndex
func (self *convWasm_convFilter) getFieldIndex(_env *LnsEnv, node *Nodes_Node,info *convWasm_ClassInfo,name string) LnsInt {
    var index LnsInt
    
    {
        _index := info.FP.Get_name2field(_env).Get(name)
        if _index == nil{
            self.FP.error(_env, node, _env.GetVM().String_format("not found member -- %s", Lns_2DDD(name)))
        } else {
            index = _index.(LnsInt)
        }
    }
    return index
}
// 916: decl @lune.@base.@convWasm.convFilter.getAlgeInfo
func (self *convWasm_convFilter) getAlgeInfo(_env *LnsEnv, algeType *Ast_TypeInfo) LnsAny {
    {
        _info := self.alge2info.Get(algeType)
        if !Lns_IsNil( _info ) {
            info := _info.(*convWasm_AlgeInfo)
            return info
        }
    }
    var work *Ast_AlgeTypeInfo
    
    {
        _work := Ast_AlgeTypeInfoDownCastF(algeType.FP)
        if _work == nil{
            return nil
        } else {
            work = _work.(*Ast_AlgeTypeInfo)
        }
    }
    var valNames *LnsList2_[string]
    valNames = NewLnsList2_[string]([]string{})
    {
        _declNode := self.algeType2decl.Get(algeType)
        if !Lns_IsNil( _declNode ) {
            declNode := _declNode.(*Nodes_DeclAlgeNode)
            for _, _valInfo := range( declNode.FP.Get_algeValList(_env).Items ) {
                valInfo := _valInfo
                valNames.Insert(valInfo.FP.Get_valSym(_env).FP.Get_name(_env))
            }
        } else {
            {
                __forsortCollection0 := work.FP.Get_valInfoMap(_env)
                __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
                __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
                for _, _name := range( __forsortSorted0.Items ) {
                    name := _name
                    valNames.Insert(name)
                }
            }
        }
    }
    var baseType LnsInt
    baseType = self.mod.FP.AddStructType(_env, self.rt.FP.Get_tBase(_env), NewLnsList2_[*WasmBin_FieldType](Lns_2DDDGen[*WasmBin_FieldType](WasmBin_field(_env, WasmBin_typeI32, true), WasmBin_field(_env, WasmBin_typeI32, false))))
    self.algeBuildingMap.Set(algeType,baseType)
    var name2tag *LnsMap2_[string,LnsInt]
    name2tag = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    var name2type *LnsMap2_[string,LnsInt]
    name2type = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    var name2global *LnsMap2_[string,LnsInt]
    name2global = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    var nameInsList *LnsList2_[*WasmBin_Ins]
    nameInsList = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _index, _name := range( valNames.Items ) {
        index := _index + 1
        name := _name
        name2tag.Set(name,index - 1)
        var valInfo *Ast_AlgeValInfo
        valInfo = Lns_unwrap( work.FP.Get_valInfoMap(_env).Get(name)).(*Ast_AlgeValInfo)
        if valInfo.FP.Get_typeList(_env).Len() == 0{
            name2type.Set(name,baseType)
            var globalIdx LnsInt
            globalIdx = self.mod.FP.AddGlobal(_env, WasmBin_refType(_env, baseType), WasmBin_refNull(_env, baseType))
            name2global.Set(name,globalIdx)
            self.prologue.Insert(WasmBin_globalSet(_env, globalIdx, WasmBin_structNew(_env, baseType, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), WasmBin_i32Const(_env, index - 1))))))
        } else { 
            var fields *LnsList2_[*WasmBin_FieldType]
            fields = NewLnsList2_[*WasmBin_FieldType](Lns_2DDDGen[*WasmBin_FieldType](WasmBin_field(_env, WasmBin_typeI32, true), WasmBin_field(_env, WasmBin_typeI32, false)))
            for _, _paramType := range( valInfo.FP.Get_typeList(_env).Items ) {
                paramType := _paramType
                fields.Insert(WasmBin_field(_env, self.FP.getRep(_env, paramType), false))
            }
            name2type.Set(name,self.mod.FP.AddStructType(_env, baseType, fields))
        }
        nameInsList.Insert(self.rt.FP.StrLiteral(_env, _env.GetVM().String_format("%s.%s", Lns_2DDD(work.FP.Get_rawTxt(_env), name))))
    }
    var namesGlobal LnsInt
    namesGlobal = self.mod.FP.AddGlobal(_env, self.rt.FP.Get_anyArrType(_env), WasmBin_refNull(_env, self.rt.FP.Get_tAnyArr(_env)))
    self.prologue.Insert(WasmBin_globalSet(_env, namesGlobal, WasmBin_arrayNewFixed(_env, self.rt.FP.Get_tAnyArr(_env), nameInsList)))
    var info *convWasm_AlgeInfo
    info = NewconvWasm_AlgeInfo(_env, algeType, baseType, valNames, name2tag, name2type, name2global, namesGlobal)
    self.algeBuildingMap.Set(algeType,nil)
    self.alge2info.Set(algeType,info)
    self.algeInfoList.Insert(info)
    return info
}
// 975: decl @lune.@base.@convWasm.convFilter.isGlobalSym
func (self *convWasm_convFilter) isGlobalSym(_env *LnsEnv, symbol *Ast_SymbolInfo) bool {
    return self.moduleScopeSet.Has(symbol.FP.Get_scope(_env))
}
// 980: decl @lune.@base.@convWasm.convFilter.cellGet
func (self *convWasm_convFilter) cellGet(_env *LnsEnv, rep *WasmBin_ValType,cell *WasmBin_Ins) *WasmBin_Ins {
    var cellType LnsInt
    cellType = self.rt.FP.GetCellType(_env, rep)
    var val *WasmBin_Ins
    val = WasmBin_structGet(_env, cellType, 0, cell)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( rep.FP.IsRef(_env)) &&
        _env.SetStackVal( rep.FP.Get_heapType(_env) != WasmBin_htAny) ).(bool)){
        return WasmBin_refCast(_env, rep.FP.Get_heapType(_env), val)
    }
    return val
}
// 995: decl @lune.@base.@convWasm.convFilter.declareVar
func (self *convWasm_convFilter) declareVar(_env *LnsEnv, symbol *Ast_SymbolInfo,init LnsAny) *WasmBin_Ins {
    var orgSym *Ast_SymbolInfo
    orgSym = symbol.FP.GetOrg(_env)
    var rep *WasmBin_ValType
    rep = self.FP.getRep(_env, orgSym.FP.Get_typeInfo(_env))
    var val *WasmBin_Ins
    if init != nil{
        init_258 := init.(*WasmBin_Ins)
        val = init_258
    } else {
        val = self.FP.getDefaultVal(_env, rep)
    }
    if self.FP.isGlobalSym(_env, orgSym){
        var globalIdx LnsInt
        {
            _loc := self.globalSym2loc.Get(orgSym)
            if !Lns_IsNil( _loc ) {
                loc := _loc
                switch _matchExp0 := loc.(type) {
                case *convWasm_VarLoc__Global:
                    index := _matchExp0.Val1
                    globalIdx = index
                default:
                    globalIdx = self.mod.FP.AddGlobal(_env, rep, self.FP.getDefaultVal(_env, rep))
                }
            } else {
                globalIdx = self.mod.FP.AddGlobal(_env, rep, self.FP.getDefaultVal(_env, rep))
                self.globalSym2loc.Set(orgSym,&convWasm_VarLoc__Global{globalIdx})
            }
        }
        return WasmBin_globalSet(_env, globalIdx, val)
    }
    if orgSym.FP.Get_hasAccessFromClosure(_env){
        var cellType LnsInt
        cellType = self.rt.FP.GetCellType(_env, rep)
        var index LnsInt
        index = self.ctx.FP.AddLocal(_env, WasmBin_refType(_env, cellType))
        self.ctx.FP.Get_sym2loc(_env).Set(orgSym,&convWasm_VarLoc__CellLocal{index})
        return convWasm_ls_3_(_env, index, WasmBin_structNew(_env, cellType, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val))))
    }
    var index LnsInt
    index = self.ctx.FP.AddLocal(_env, rep)
    self.ctx.FP.Get_sym2loc(_env).Set(orgSym,&convWasm_VarLoc__Local{index})
    return convWasm_ls_3_(_env, index, val)
}
// 1032: decl @lune.@base.@convWasm.convFilter.findLoc
func (self *convWasm_convFilter) findLoc(_env *LnsEnv, symbol *Ast_SymbolInfo) LnsAny {
    var orgSym *Ast_SymbolInfo
    orgSym = symbol.FP.GetOrg(_env)
    {
        _loc := self.ctx.FP.Get_sym2loc(_env).Get(orgSym)
        if !Lns_IsNil( _loc ) {
            loc := _loc
            return loc
        }
    }
    return self.globalSym2loc.Get(orgSym)
}
// 1040: decl @lune.@base.@convWasm.convFilter.envGet
func (self *convWasm_convFilter) envGet(_env *LnsEnv, field LnsInt) *WasmBin_Ins {
    return WasmBin_structGet(_env, self.ctx.FP.Get_envType(_env), field, convWasm_lg_2_(_env, self.ctx.FP.Get_envLocal(_env)))
}
// 1045: decl @lune.@base.@convWasm.convFilter.readSym
func (self *convWasm_convFilter) readSym(_env *LnsEnv, node *Nodes_Node,symbol *Ast_SymbolInfo) *WasmBin_Ins {
    {
        _initIns := self.FP.getLazyInit(_env, node)
        if !Lns_IsNil( _initIns ) {
            initIns := _initIns.(*WasmBin_Ins)
            return WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](initIns, self.FP.readSymSub(_env, node, symbol))))
        }
    }
    return self.FP.readSymSub(_env, node, symbol)
}
// 1052: decl @lune.@base.@convWasm.convFilter.readSymSub
func (self *convWasm_convFilter) readSymSub(_env *LnsEnv, node *Nodes_Node,symbol *Ast_SymbolInfo) *WasmBin_Ins {
    var orgSym *Ast_SymbolInfo
    orgSym = self.FP.resolveSym(_env, symbol.FP.GetOrg(_env))
    var rep *WasmBin_ValType
    rep = self.FP.getRep(_env, orgSym.FP.Get_typeInfo(_env))
    {
        _loc := self.FP.findLoc(_env, orgSym)
        if !Lns_IsNil( _loc ) {
            loc := _loc
            switch _matchExp0 := loc.(type) {
            case *convWasm_VarLoc__Local:
                index := _matchExp0.Val1
                return convWasm_lg_2_(_env, index)
            case *convWasm_VarLoc__CellLocal:
                index := _matchExp0.Val1
                return self.FP.cellGet(_env, rep, convWasm_lg_2_(_env, index))
            case *convWasm_VarLoc__Captured:
                field := _matchExp0.Val1
                return self.FP.cellGet(_env, rep, self.FP.envGet(_env, field))
            case *convWasm_VarLoc__Global:
                index := _matchExp0.Val1
                return WasmBin_globalGet(_env, index)
            }
        }
    }
    if orgSym.FP.Get_kind(_env) == Ast_SymbolKind__Fun{
        {
            _funcIndex := self.func2index.Get(orgSym.FP.Get_typeInfo(_env))
            if !Lns_IsNil( _funcIndex ) {
                funcIndex := _funcIndex.(LnsInt)
                var info *WasmRt_CloInfo
                info = self.FP.getFuncCloInfo(_env, orgSym.FP.Get_typeInfo(_env))
                return WasmBin_structNew(_env, info.FP.Get_cloType(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), self.mod.FP.RefFunc(_env, funcIndex))))
            }
        }
    }
    if _switch0 := orgSym.FP.Get_name(_env); _switch0 == "_G" || _switch0 == "_ENV" {
        return WasmBin_call(_env, self.rt.FP.FnLuaToMap(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.ListItem(_env, WasmBin_call(_env, self.rt.FP.FnLuaCall(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, WasmRt_luaGlobals), WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0))))))), 0), WasmBin_i32Const(_env, 0))))
    } else if _switch0 == "__mod__" {
        return self.rt.FP.StrLiteral(_env, self.moduleTypeInfo.FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), false))
    } else if _switch0 == "__func__" {
        var funcType *Ast_TypeInfo
        funcType = orgSym.FP.Get_namespaceTypeInfo(_env)
        var nameSpace string
        nameSpace = convWasm_convExp1_4268(Lns_2DDD(_env.GetVM().String_gsub(funcType.FP.Get_parentInfo(_env).FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), false),"&", "")))
        var funcName string
        funcName = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.funcSym2name.Get(orgSym)) ||
            _env.SetStackVal( "") ).(string)
        if funcName == ""{
            funcName = "<anonymous>"
        }
        return self.rt.FP.StrLiteral(_env, _env.GetVM().String_format("%s.%s", Lns_2DDD(nameSpace, funcName)))
    }
    self.FP.error(_env, node, _env.GetVM().String_format("not support symbol -- %s", Lns_2DDD(orgSym.FP.Get_name(_env))))
// insert a dummy
    return nil
}
// 1109: decl @lune.@base.@convWasm.convFilter.writeSym
func (self *convWasm_convFilter) writeSym(_env *LnsEnv, node *Nodes_Node,symbol *Ast_SymbolInfo,val *WasmBin_Ins) *WasmBin_Ins {
    var orgSym *Ast_SymbolInfo
    orgSym = symbol.FP.GetOrg(_env)
    if orgSym.FP.Get_name(_env) == "_"{
        return WasmBin_drop(_env, val)
    }
    var rep *WasmBin_ValType
    rep = self.FP.getRep(_env, orgSym.FP.Get_typeInfo(_env))
    var loc LnsAny
    
    {
        _loc := self.FP.findLoc(_env, orgSym)
        if _loc == nil{
            self.FP.error(_env, node, _env.GetVM().String_format("not support symbol -- %s", Lns_2DDD(orgSym.FP.Get_name(_env))))
        } else {
            loc = _loc
        }
    }
    switch _matchExp0 := loc.(type) {
    case *convWasm_VarLoc__Local:
        index := _matchExp0.Val1
        return convWasm_ls_3_(_env, index, val)
    case *convWasm_VarLoc__CellLocal:
        index := _matchExp0.Val1
        return WasmBin_structSet(_env, self.rt.FP.GetCellType(_env, rep), 0, convWasm_lg_2_(_env, index), val)
    case *convWasm_VarLoc__Captured:
        field := _matchExp0.Val1
        return WasmBin_structSet(_env, self.rt.FP.GetCellType(_env, rep), 0, self.FP.envGet(_env, field), val)
    case *convWasm_VarLoc__Global:
        index := _matchExp0.Val1
        return WasmBin_globalSet(_env, index, val)
    }
// insert a dummy
    return nil
}
// 1137: decl @lune.@base.@convWasm.convFilter.isFuncType
func (self *convWasm_convFilter) isFuncType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    if _switch0 := self.FP.normType(_env, typeInfo).FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Func || _switch0 == Ast_TypeInfoKind__Form || _switch0 == Ast_TypeInfoKind__FormFunc {
        return true
    }
    return false
}
// 1147: decl @lune.@base.@convWasm.convFilter.coerceRep
func (self *convWasm_convFilter) coerceRep(_env *LnsEnv, ins *WasmBin_Ins,fromRep *WasmBin_ValType,toRep *WasmBin_ValType) *WasmBin_Ins {
    if fromRep.FP.EqualsType(_env, toRep){
        return ins
    }
    if convWasm_isScalar_15_(_env, fromRep){
        if convWasm_isScalar_15_(_env, toRep){
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( fromRep.FP.Get_kind(_env) == WasmBin_ValKind__I64) &&
                _env.SetStackVal( toRep.FP.Get_kind(_env) == WasmBin_ValKind__F64) ).(bool)){
                return convWasm_p_4_(_env, WasmBin_opF64ConvertI64S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](ins)))
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( fromRep.FP.Get_kind(_env) == WasmBin_ValKind__F64) &&
                _env.SetStackVal( toRep.FP.Get_kind(_env) == WasmBin_ValKind__I64) ).(bool)){
                return convWasm_p_4_(_env, WasmBin_opI64TruncF64S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](ins)))
            }
            return ins
        }
        var boxed *WasmBin_Ins
        if _switch0 := fromRep.FP.Get_kind(_env); _switch0 == WasmBin_ValKind__I64 {
            boxed = self.rt.FP.BoxI64(_env, ins)
        } else if _switch0 == WasmBin_ValKind__F64 {
            boxed = self.rt.FP.BoxF64(_env, ins)
        } else if _switch0 == WasmBin_ValKind__I32 || _switch0 == WasmBin_ValKind__Ref {
            boxed = self.rt.FP.BoxBool(_env, ins)
        }
        return boxed
    }
    if convWasm_isScalar_15_(_env, toRep){
        if _switch1 := toRep.FP.Get_kind(_env); _switch1 == WasmBin_ValKind__I64 {
            return self.rt.FP.UnboxI64(_env, ins)
        } else if _switch1 == WasmBin_ValKind__F64 {
            return self.rt.FP.UnboxF64(_env, ins)
        } else if _switch1 == WasmBin_ValKind__I32 || _switch1 == WasmBin_ValKind__Ref {
            return self.rt.FP.UnboxBool(_env, ins)
        }
    }
    if toRep.FP.Get_heapType(_env) == WasmBin_htAny{
        return ins
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( fromRep.FP.Get_heapType(_env) == WasmBin_htAny) &&
        _env.SetStackVal( toRep.FP.EqualsType(_env, self.rt.FP.Get_mapType(_env))) ).(bool)){
        return WasmBin_call(_env, self.FP.fnAnyToMap(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](ins)))
    }
    return WasmBin_refCast(_env, toRep.FP.Get_heapType(_env), ins)
}
// 1205: decl @lune.@base.@convWasm.convFilter.fnAnyToMap
func (self *convWasm_convFilter) fnAnyToMap(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("anyToMap")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_mapType(_env))))
    self.adapterMap.Set("anyToMap",index)
    self.mod.FP.SetFuncName(_env, index, "anyToMap")
    var list LnsInt
    list = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
    var _map LnsInt
    _map = _func.FP.AddLocal(_env, self.rt.FP.Get_mapType(_env))
    var pos LnsInt
    pos = _func.FP.AddLocal(_env, WasmBin_typeI32)
    var item LnsInt
    item = _func.FP.AddLocal(_env, WasmBin_typeAnyRef)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refTest(_env, self.rt.FP.Get_tList(_env), convWasm_lg_2_(_env, 0))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refCast(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0))))))), nil, nil), convWasm_ls_3_(_env, list, WasmBin_refCast(_env, self.rt.FP.Get_tList(_env), convWasm_lg_2_(_env, 0))), convWasm_ls_3_(_env, _map, self.rt.FP.NewMap(_env)), WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, convWasm_lg_2_(_env, list)))))), convWasm_ls_3_(_env, item, WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListItems, convWasm_lg_2_(_env, list)), convWasm_lg_2_(_env, pos))), convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))), WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refIsNull(_env, convWasm_lg_2_(_env, item))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, _map), WasmBin_call(_env, self.rt.FP.FnI64ToStr(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64ExtendI32S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos))))))), convWasm_lg_2_(_env, item)))))), nil, nil), WasmBin_br(_env, 0))), nil))), nil), convWasm_lg_2_(_env, _map))))
    return index
}
// 1240: decl @lune.@base.@convWasm.convFilter.fnMapToList
func (self *convWasm_convFilter) fnMapToList(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("mapToList")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_mapType(_env))), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_listType(_env))))
    self.adapterMap.Set("mapToList",index)
    self.mod.FP.SetFuncName(_env, index, "mapToList")
    var list LnsInt
    list = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
    var pos LnsInt
    pos = _func.FP.AddLocal(_env, WasmBin_typeI64)
    var item LnsInt
    item = _func.FP.AddLocal(_env, WasmBin_typeAnyRef)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, list, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 4))))), WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI64Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i64Const(_env, 1))))), convWasm_ls_3_(_env, item, WasmBin_call(_env, self.rt.FP.FnMapGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), self.rt.FP.BoxI64(_env, convWasm_lg_2_(_env, pos)))))), WasmBin_brIf(_env, 1, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, item))), WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, list), convWasm_lg_2_(_env, item)))), WasmBin_br(_env, 0))), nil))), nil), convWasm_lg_2_(_env, list))))
    return index
}
// 1263: decl @lune.@base.@convWasm.convFilter.coerce
func (self *convWasm_convFilter) coerce(_env *LnsEnv, ins *WasmBin_Ins,fromType *Ast_TypeInfo,toType *Ast_TypeInfo) *WasmBin_Ins {
    var fromRep *WasmBin_ValType
    fromRep = self.FP.getRep(_env, fromType)
    var toRep *WasmBin_ValType
    toRep = self.FP.getRep(_env, toType)
    if fromRep.FP.EqualsType(_env, toRep){
        return ins
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.isFuncType(_env, fromType)) &&
        _env.SetStackVal( self.FP.isFuncType(_env, toType)) ).(bool)){
        return self.FP.adaptClosure(_env, ins, fromType, toType)
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.isLuaval(_env, fromType)) &&
        _env.SetStackVal( Lns_op_not(self.FP.isLuaval(_env, toType))) &&
        _env.SetStackVal( self.FP.isFuncType(_env, toType)) ).(bool)){
        return self.FP.luaFuncToClosure(_env, ins, toType)
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.isType(_env, fromType, Ast_builtinTypeStem)) &&
        _env.SetStackVal( self.FP.isFuncType(_env, toType)) ).(bool)){
        var work LnsInt
        work = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
        return WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, ins), WasmBin_ifElse(_env, WasmBin_refTest(_env, self.rt.FP.Get_tLuaVal(_env), convWasm_lg_2_(_env, work)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.luaFuncToClosure(_env, convWasm_lg_2_(_env, work), toType))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.coerceRep(_env, convWasm_lg_2_(_env, work), fromRep, toRep))), toRep))))
    }
    return self.FP.coerceRep(_env, ins, fromRep, toRep)
}
// 1292: decl @lune.@base.@convWasm.convFilter.luaFuncToClosure
func (self *convWasm_convFilter) luaFuncToClosure(_env *LnsEnv, ins *WasmBin_Ins,toType *Ast_TypeInfo) *WasmBin_Ins {
    var toInfo *WasmRt_CloInfo
    toInfo = self.FP.getFuncCloInfo(_env, toType)
    var key string
    key = _env.GetVM().String_format("luaFunc:%d", Lns_2DDD(toInfo.FP.Get_funcType(_env)))
    var adaptType LnsInt
    {
        _work := self.adapterMap.Get(key)
        if !Lns_IsNil( _work ) {
            work := _work.(LnsInt)
            adaptType = work
        } else {
            adaptType = self.mod.FP.AddStructType(_env, toInfo.FP.Get_cloType(_env), NewLnsList2_[*WasmBin_FieldType](Lns_2DDDGen[*WasmBin_FieldType](WasmBin_field(_env, WasmBin_typeI32, true), WasmBin_field(_env, WasmBin_refType(_env, toInfo.FP.Get_funcType(_env)), false), WasmBin_field(_env, WasmBin_typeAnyRef, false))))
            self.adapterMap.Set(key,adaptType)
            var funcIndex LnsInt
            var _func *WasmBin_Func
            funcIndex,_func = self.mod.FP.AddFuncWithType(_env, toInfo.FP.Get_funcType(_env), 1 + toInfo.FP.Get_params(_env).Len())
            self.adapterMap.Set(key + ":func",funcIndex)
            var args LnsInt
            args = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
            var result LnsInt
            result = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
            var body *LnsList2_[*WasmBin_Ins]
            body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, args, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, toInfo.FP.Get_params(_env).Len())))))))
            var toArgs *LnsList2_[*Ast_TypeInfo]
            toArgs = self.FP.normType(_env, toType).FP.Get_argTypeInfoList(_env)
            for _index, _rep := range( toInfo.FP.Get_params(_env).Items ) {
                index := _index + 1
                rep := _rep
                if self.FP.normType(_env, toArgs.GetAt(index)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    body.Insert(WasmBin_call(_env, self.rt.FP.FnListAppend(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, args), convWasm_lg_2_(_env, index)))))
                } else { 
                    body.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, args), self.FP.coerceRep(_env, convWasm_lg_2_(_env, index), rep, WasmBin_typeAnyRef)))))
                }
            }
            body.Insert(convWasm_ls_3_(_env, result, WasmBin_call(_env, self.rt.FP.FnLuaCallValLazy(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structGet(_env, adaptType, 2, WasmBin_refCast(_env, adaptType, convWasm_lg_2_(_env, 0))), convWasm_lg_2_(_env, args))))))
            var toRetTypes *LnsList2_[*Ast_TypeInfo]
            toRetTypes = self.FP.normType(_env, toType).FP.Get_retTypeInfoList(_env)
            for _index, _rep := range( toInfo.FP.Get_results(_env).Items ) {
                index := _index + 1
                rep := _rep
                if self.FP.normType(_env, toRetTypes.GetAt(index)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    if index != 1{
                        Util_err(_env, _env.GetVM().String_format("wasm: not support the luaval form -- %s", Lns_2DDD(toType.FP.GetTxt(_env, nil, nil, nil))))
                    }
                    body.Insert(convWasm_lg_2_(_env, result))
                } else { 
                    body.Insert(self.FP.coerceRep(_env, WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, result), WasmBin_i64Const(_env, index)))), WasmBin_typeAnyRef, rep))
                }
            }
            _func.FP.SetBody(_env, body)
        }
    }
    var funcIndex LnsInt
    funcIndex = Lns_unwrap( self.adapterMap.Get(key + ":func")).(LnsInt)
    var work LnsInt
    work = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
    return WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, ins), WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, work)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, toInfo.FP.Get_cloType(_env)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structNew(_env, adaptType, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), self.mod.FP.RefFunc(_env, funcIndex), convWasm_lg_2_(_env, work)))))), WasmBin_refType(_env, toInfo.FP.Get_cloType(_env))))))
}
// 1348: decl @lune.@base.@convWasm.convFilter.coerceVal
func (self *convWasm_convFilter) coerceVal(_env *LnsEnv, val *convWasm_Val,toType *Ast_TypeInfo) *WasmBin_Ins {
    return self.FP.coerce(_env, val.FP.Get_ins(_env), val.FP.Get_typeInfo(_env), toType)
}
// 1357: decl @lune.@base.@convWasm.convFilter.adaptClosure
func (self *convWasm_convFilter) adaptClosure(_env *LnsEnv, ins *WasmBin_Ins,fromType *Ast_TypeInfo,toType *Ast_TypeInfo) *WasmBin_Ins {
    var fromInfo *WasmRt_CloInfo
    fromInfo = self.FP.getFuncCloInfo(_env, fromType)
    var toInfo *WasmRt_CloInfo
    toInfo = self.FP.getFuncCloInfo(_env, toType)
    var key string
    key = _env.GetVM().String_format("%d:%d", Lns_2DDD(fromInfo.FP.Get_funcType(_env), toInfo.FP.Get_funcType(_env)))
    var adaptType LnsInt
    {
        _work := self.adapterMap.Get(key)
        if !Lns_IsNil( _work ) {
            work := _work.(LnsInt)
            adaptType = work
        } else {
            adaptType = self.mod.FP.AddStructType(_env, toInfo.FP.Get_cloType(_env), NewLnsList2_[*WasmBin_FieldType](Lns_2DDDGen[*WasmBin_FieldType](WasmBin_field(_env, WasmBin_typeI32, true), WasmBin_field(_env, WasmBin_refType(_env, toInfo.FP.Get_funcType(_env)), false), WasmBin_field(_env, WasmBin_refType(_env, fromInfo.FP.Get_cloType(_env)), false))))
            self.adapterMap.Set(key,adaptType)
            var funcIndex LnsInt
            var _func *WasmBin_Func
            funcIndex,_func = self.mod.FP.AddFuncWithType(_env, toInfo.FP.Get_funcType(_env), 1 + toInfo.FP.Get_params(_env).Len())
            self.adapterMap.Set(key + ":func",funcIndex)
            var inner LnsInt
            inner = _func.FP.AddLocal(_env, WasmBin_refType(_env, fromInfo.FP.Get_cloType(_env)))
            var body *LnsList2_[*WasmBin_Ins]
            body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, inner, WasmBin_structGet(_env, adaptType, 2, WasmBin_refCast(_env, adaptType, convWasm_lg_2_(_env, 0))))))
            var fromArgs *LnsList2_[*Ast_TypeInfo]
            fromArgs = self.FP.normType(_env, fromType).FP.Get_argTypeInfoList(_env)
            var toArgs *LnsList2_[*Ast_TypeInfo]
            toArgs = self.FP.normType(_env, toType).FP.Get_argTypeInfoList(_env)
            var toHasDDD bool
            toHasDDD = _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( toArgs.Len() > 0) &&
                _env.SetStackVal( self.FP.normType(_env, toArgs.GetAt(toArgs.Len())).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)
            var args *LnsList2_[*WasmBin_Ins]
            args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, inner)))
            for _index, _fromArg := range( fromArgs.Items ) {
                index := _index + 1
                fromArg := _fromArg
                var fromRep *WasmBin_ValType
                fromRep = self.FP.getRep(_env, fromArg)
                if self.FP.normType(_env, fromArg).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( toHasDDD) &&
                        _env.SetStackVal( index == toArgs.Len()) ).(bool)){
                        args.Insert(convWasm_lg_2_(_env, toArgs.Len()))
                    } else { 
                        args.Insert(WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0)))))
                    }
                } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( toHasDDD) &&
                    _env.SetStackVal( index >= toArgs.Len()) ).(bool)){
                    args.Insert(self.FP.coerceRep(_env, WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, toArgs.Len()), WasmBin_i64Const(_env, index - toArgs.Len() + 1)))), WasmBin_typeAnyRef, fromRep))
                } else if index <= toArgs.Len(){
                    args.Insert(self.FP.coerceRep(_env, convWasm_lg_2_(_env, index), self.FP.getRep(_env, toArgs.GetAt(index)), fromRep))
                } else { 
                    args.Insert(self.FP.getDefaultVal(_env, fromRep))
                }
            }
            args.Insert(WasmBin_structGet(_env, fromInfo.FP.Get_cloType(_env), WasmRt_fCloCode, convWasm_lg_2_(_env, inner)))
            body.Insert(WasmBin_callRef(_env, fromInfo.FP.Get_funcType(_env), args))
            var fromRets *LnsList2_[*WasmBin_ValType]
            fromRets = fromInfo.FP.Get_results(_env)
            var toRets *LnsList2_[*WasmBin_ValType]
            toRets = toInfo.FP.Get_results(_env)
            var retLocals *LnsList2_[LnsInt]
            retLocals = NewLnsList2_[LnsInt]([]LnsInt{})
            for _, _rep := range( fromRets.Items ) {
                rep := _rep
                retLocals.Insert(LnsInt(_func.FP.AddLocal(_env, rep)))
            }
            {
                var _forFrom0 LnsInt = fromRets.Len()
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
                    body.Insert(convWasm_ls_3_(_env, retLocals.GetAt(index), WasmBin_nop(_env)))
                    _forWork0 += _forDelta0
                }
            }
            var rets *LnsList2_[*WasmBin_Ins]
            rets = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
            var toRetTypes *LnsList2_[*Ast_TypeInfo]
            toRetTypes = self.FP.normType(_env, toType).FP.Get_retTypeInfoList(_env)
            for _index, _toRep := range( toRets.Items ) {
                index := _index + 1
                toRep := _toRep
                var isDDD bool
                isDDD = self.FP.normType(_env, toRetTypes.GetAt(index)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( isDDD) &&
                    _env.SetStackVal( index == toRets.Len()) ).(bool)){
                    var list LnsInt
                    list = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
                    body.Insert(convWasm_ls_3_(_env, list, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0))))))
                    {
                        var _forFrom1 LnsInt = index
                        var _forTo1 LnsInt = fromRets.Len()
                        for _forWork1 := _forFrom1; _forWork1 <= _forTo1; _forWork1++ {
                            work := _forWork1
                            body.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, list), self.FP.coerceRep(_env, convWasm_lg_2_(_env, retLocals.GetAt(work)), fromRets.GetAt(work), WasmBin_typeAnyRef)))))
                        }
                    }
                    rets.Insert(convWasm_lg_2_(_env, list))
                } else if index <= fromRets.Len(){
                    rets.Insert(self.FP.coerceRep(_env, convWasm_lg_2_(_env, retLocals.GetAt(index)), fromRets.GetAt(index), toRep))
                } else { 
                    rets.Insert(self.FP.getDefaultVal(_env, toRep))
                }
            }
            body.Insert(WasmBin_seq(_env, rets))
            _func.FP.SetBody(_env, body)
        }
    }
    var funcIndex LnsInt
    funcIndex = Lns_unwrap( self.adapterMap.Get(key + ":func")).(LnsInt)
    return WasmBin_structNew(_env, adaptType, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), self.mod.FP.RefFunc(_env, funcIndex), WasmBin_refCast(_env, fromInfo.FP.Get_cloType(_env), ins))))
}
// 1444: decl @lune.@base.@convWasm.convFilter.toAny
func (self *convWasm_convFilter) toAny(_env *LnsEnv, val *convWasm_Val) *WasmBin_Ins {
    return self.FP.coerceRep(_env, val.FP.Get_ins(_env), val.FP.Get_rep(_env), WasmBin_typeAnyRef)
}
// 1449: decl @lune.@base.@convWasm.convFilter.fnTruthy
func (self *convWasm_convFilter) fnTruthy(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("truthy")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI32)))
    self.adapterMap.Set("truthy",index)
    self.mod.FP.SetFuncName(_env, index, "truthy")
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0)))))), nil, nil), WasmBin_ifElse(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxBool(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.UnboxBool(_env, convWasm_lg_2_(_env, 0))))))), nil, nil), WasmBin_i32Const(_env, 1))))
    return index
}
// 1465: decl @lune.@base.@convWasm.convFilter.toCond
func (self *convWasm_convFilter) toCond(_env *LnsEnv, val *convWasm_Val) *WasmBin_Ins {
    if _switch0 := val.FP.Get_rep(_env).FP.Get_kind(_env); _switch0 == WasmBin_ValKind__I32 {
        return val.FP.Get_ins(_env)
    } else if _switch0 == WasmBin_ValKind__I64 || _switch0 == WasmBin_ValKind__F64 {
        return WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_drop(_env, val.FP.Get_ins(_env)), WasmBin_i32Const(_env, 1))))
    } else if _switch0 == WasmBin_ValKind__Ref {
        if val.FP.Get_rep(_env).FP.Get_heapType(_env) == WasmBin_htAny{
            return WasmBin_call(_env, self.FP.fnTruthy(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val.FP.Get_ins(_env))))
        }
        return convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refIsNull(_env, val.FP.Get_ins(_env)))))
    }
// insert a dummy
    return nil
}
// 1482: decl @lune.@base.@convWasm.convFilter.lowerCond
func (self *convWasm_convFilter) lowerCond(_env *LnsEnv, node *Nodes_Node) *WasmBin_Ins {
    return self.FP.toCond(_env, self.FP.lowerVal(_env, node))
}
// 1489: decl @lune.@base.@convWasm.convFilter.expTypeOf
func (self *convWasm_convFilter) expTypeOf(_env *LnsEnv, node *Nodes_Node) *Ast_TypeInfo {
    {
        _castNode := Nodes_ExpCastNodeDownCastF(node.FP)
        if !Lns_IsNil( _castNode ) {
            castNode := _castNode.(*Nodes_ExpCastNode)
            return castNode.FP.Get_castType(_env)
        }
    }
    return node.FP.Get_expType(_env)
}
// 1496: decl @lune.@base.@convWasm.convFilter.newVal
func (self *convWasm_convFilter) newVal(_env *LnsEnv, ins *WasmBin_Ins,typeInfo *Ast_TypeInfo) *convWasm_Val {
    return NewconvWasm_Val(_env, ins, self.FP.getRep(_env, typeInfo), typeInfo)
}
// 1501: decl @lune.@base.@convWasm.convFilter.lowerExp
func (self *convWasm_convFilter) lowerExp(_env *LnsEnv, node *Nodes_Node) *WasmBin_Ins {
    return self.FP.coerceVal(_env, self.FP.lowerVal(_env, node), self.FP.expTypeOf(_env, node))
}
// 1506: decl @lune.@base.@convWasm.convFilter.lowerExpAs
func (self *convWasm_convFilter) lowerExpAs(_env *LnsEnv, node *Nodes_Node,typeInfo *Ast_TypeInfo) *WasmBin_Ins {
    return self.FP.coerceVal(_env, self.FP.lowerVal(_env, node), typeInfo)
}
// 1511: decl @lune.@base.@convWasm.convFilter.tempLocal
func (self *convWasm_convFilter) tempLocal(_env *LnsEnv, rep *WasmBin_ValType) LnsInt {
    return self.ctx.FP.AddLocal(_env, rep)
}
// 1516: decl @lune.@base.@convWasm.convFilter.buildList
func (self *convWasm_convFilter) buildList(_env *LnsEnv, vals *LnsList2_[*convWasm_Val]) *WasmBin_Ins {
    var list LnsInt
    list = self.FP.tempLocal(_env, self.rt.FP.Get_listType(_env))
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, list, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, vals.Len())))))))
    for _, _val := range( vals.Items ) {
        val := _val
        if self.FP.normType(_env, val.FP.Get_typeInfo(_env)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            body.Insert(WasmBin_call(_env, self.rt.FP.FnListAppend(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, list), val.FP.Get_ins(_env)))))
        } else { 
            body.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, list), self.FP.toAny(_env, val)))))
        }
    }
    body.Insert(convWasm_lg_2_(_env, list))
    return WasmBin_seq(_env, body)
}
// 1536: decl @lune.@base.@convWasm.convFilter.lowerValList
func (self *convWasm_convFilter) lowerValList(_env *LnsEnv, expList *Nodes_ExpListNode,pre *LnsList2_[*WasmBin_Ins]) *LnsList2_[*convWasm_Val] {
    return self.FP.lowerValListSub(_env, expList, pre, false)
}
// 1542: decl @lune.@base.@convWasm.convFilter.lowerValListSub
func (self *convWasm_convFilter) lowerValListSub(_env *LnsEnv, expList *Nodes_ExpListNode,pre *LnsList2_[*WasmBin_Ins],raw bool) *LnsList2_[*convWasm_Val] {
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    var mRetIndex LnsInt
    {
        _mRetExp := expList.FP.Get_mRetExp(_env)
        if !Lns_IsNil( _mRetExp ) {
            mRetExp := _mRetExp.(*Nodes_MRetExp)
            mRetIndex = mRetExp.FP.Get_index(_env)
        } else {
            mRetIndex = 0
        }
    }
    var castMRet LnsAny
    castMRet = nil
    var castMRetVals *LnsList2_[*convWasm_Val]
    castMRetVals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    var castMRetPos LnsInt
    castMRetPos = 0
    for _index, _exp := range( expList.FP.Get_expList(_env).Items ) {
        index := _index + 1
        exp := _exp
        var mretInCast LnsAny
        mretInCast = nil
        var castType *Ast_TypeInfo
        castType = exp.FP.Get_expType(_env)
        var isRawMRet bool
        isRawMRet = false
        {
            _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
            if !Lns_IsNil( _castNode ) {
                castNode := _castNode.(*Nodes_ExpCastNode)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( castNode.FP.Get_castKind(_env) == Nodes_CastKind__Implicit) &&
                    _env.SetStackVal( castNode.FP.Get_exp(_env).FP.Get_kind(_env) == Nodes_NodeKind_get_ExpMRet(_env)) ).(bool)){
                    mretInCast = castNode.FP.Get_exp(_env)
                    castType = castNode.FP.Get_castType(_env)
                }
            } else {
                if exp.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpMRet(_env){
                    for _, _other := range( expList.FP.Get_expList(_env).Items ) {
                        other := _other
                        {
                            _otherCast := Nodes_ExpCastNodeDownCastF(other.FP)
                            if !Lns_IsNil( _otherCast ) {
                                otherCast := _otherCast.(*Nodes_ExpCastNode)
                                if otherCast.FP.Get_exp(_env) == exp{
                                    mretInCast = exp
                                    isRawMRet = true
                                }
                            }
                        }
                    }
                }
            }
        }
        if Lns_isCondTrue( mretInCast){
            var mret *Nodes_Node
            mret = Lns_unwrap( mretInCast).(*Nodes_Node)
            if castMRet != mret{
                castMRet = mret
                var num LnsInt
                num = 0
                for _, _other := range( expList.FP.Get_expList(_env).Items ) {
                    other := _other
                    {
                        _otherCast := Nodes_ExpCastNodeDownCastF(other.FP)
                        if !Lns_IsNil( _otherCast ) {
                            otherCast := _otherCast.(*Nodes_ExpCastNode)
                            if otherCast.FP.Get_exp(_env) == mret{
                                num = num + 1
                            }
                        } else {
                            if _env.PopVal( _env.IncStack() ||
                                _env.SetStackVal( other == mret) ||
                                _env.SetStackVal( self.FP.isAccessOf(_env, other, mret)) ).(bool){
                                num = num + 1
                            }
                        }
                    }
                }
                castMRetVals = self.FP.expandDDD(_env, self.FP.lowerMulti(_env, mret, pre), num, pre)
                castMRetPos = 0
            }
            castMRetPos = castMRetPos + 1
            var val *convWasm_Val
            if castMRetPos <= castMRetVals.Len(){
                val = castMRetVals.GetAt(castMRetPos)
            } else { 
                val = NewconvWasm_Val(_env, WasmBin_refNull(_env, WasmBin_htAny), WasmBin_typeAnyRef, Ast_builtinTypeNil)
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( isRawMRet) ||
                _env.SetStackVal( raw) ).(bool){
                vals.Insert(val)
            } else { 
                vals.Insert(NewconvWasm_Val(_env, self.FP.coerceVal(_env, val, castType), self.FP.getRep(_env, castType), castType))
            }
        } else if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( castMRet) &&
            _env.SetStackVal( self.FP.isAccessOf(_env, exp, Lns_unwrap( castMRet).(*Nodes_Node))) )){
            castMRetPos = castMRetPos + 1
            if castMRetPos <= castMRetVals.Len(){
                vals.Insert(castMRetVals.GetAt(castMRetPos))
            } else { 
                vals.Insert(NewconvWasm_Val(_env, WasmBin_refNull(_env, WasmBin_htAny), WasmBin_typeAnyRef, Ast_builtinTypeNil))
            }
        } else if exp.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpAccessMRet(_env){
        } else if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( index == mRetIndex) ||
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( index == expList.FP.Get_expList(_env).Len()) &&
                _env.SetStackVal( exp.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpMRet(_env)) ).(bool))) ).(bool){
            for _, _val := range( self.FP.lowerMulti(_env, exp, pre).Items ) {
                val := _val
                vals.Insert(val)
            }
        } else if mRetIndex > 0{
            var val *convWasm_Val
            val = self.FP.lowerVal(_env, exp)
            var work LnsInt
            work = self.FP.tempLocal(_env, val.FP.Get_rep(_env))
            pre.Insert(convWasm_ls_3_(_env, work, val.FP.Get_ins(_env)))
            vals.Insert(NewconvWasm_Val(_env, convWasm_lg_2_(_env, work), val.FP.Get_rep(_env), val.FP.Get_typeInfo(_env)))
        } else { 
            vals.Insert(self.FP.lowerVal(_env, exp))
        }
    }
    return vals
}
// 1639: decl @lune.@base.@convWasm.convFilter.isAccessOf
func (self *convWasm_convFilter) isAccessOf(_env *LnsEnv, node *Nodes_Node,mret *Nodes_Node) bool {
    var accessNode *Nodes_ExpAccessMRetNode
    
    {
        _accessNode := Nodes_ExpAccessMRetNodeDownCastF(node.FP)
        if _accessNode == nil{
            return false
        } else {
            accessNode = _accessNode.(*Nodes_ExpAccessMRetNode)
        }
    }
    if accessNode.FP.Get_mRet(_env) == mret{
        return true
    }
    {
        _mretNode := Nodes_ExpMRetNodeDownCastF(mret.FP)
        if !Lns_IsNil( _mretNode ) {
            mretNode := _mretNode.(*Nodes_ExpMRetNode)
            return accessNode.FP.Get_mRet(_env) == mretNode.FP.Get_mRet(_env)
        }
    }
    return false
}
// 1653: decl @lune.@base.@convWasm.convFilter.lowerMulti
func (self *convWasm_convFilter) lowerMulti(_env *LnsEnv, node *Nodes_Node,pre *LnsList2_[*WasmBin_Ins]) *LnsList2_[*convWasm_Val] {
    var work *Nodes_Node
    work = node
    {
        _mretNode := Nodes_ExpMRetNodeDownCastF(node.FP)
        if !Lns_IsNil( _mretNode ) {
            mretNode := _mretNode.(*Nodes_ExpMRetNode)
            work = mretNode.FP.Get_mRet(_env)
        }
    }
    for  {
        var macroNode *Nodes_ExpMacroExpNode
        
        {
            _macroNode := Nodes_ExpMacroExpNodeDownCastF(work.FP)
            if _macroNode == nil{
                break
            } else {
                macroNode = _macroNode.(*Nodes_ExpMacroExpNode)
            }
        }
        if macroNode.FP.Get_stmtList(_env).Len() != 1{
            break
        }
        var stmt *Nodes_Node
        stmt = macroNode.FP.Get_stmtList(_env).GetAt(1)
        {
            _stmtExp := Nodes_StmtExpNodeDownCastF(stmt.FP)
            if !Lns_IsNil( _stmtExp ) {
                stmtExp := _stmtExp.(*Nodes_StmtExpNode)
                work = stmtExp.FP.Get_exp(_env)
            } else {
                work = stmt
            }
        }
    }
    var callNode *Nodes_ExpCallNode
    
    {
        _callNode := Nodes_ExpCallNodeDownCastF(work.FP)
        if _callNode == nil{
            var val *convWasm_Val
            val = self.FP.lowerVal(_env, work)
            var tmp LnsInt
            tmp = self.FP.tempLocal(_env, val.FP.Get_rep(_env))
            pre.Insert(convWasm_ls_3_(_env, tmp, val.FP.Get_ins(_env)))
            return NewLnsList2_[*convWasm_Val](Lns_2DDDGen[*convWasm_Val](NewconvWasm_Val(_env, convWasm_lg_2_(_env, tmp), val.FP.Get_rep(_env), val.FP.Get_typeInfo(_env))))
        } else {
            callNode = _callNode.(*Nodes_ExpCallNode)
        }
    }
    var ins *WasmBin_Ins
    var retTypes *LnsList2_[*Ast_TypeInfo]
    ins,retTypes = self.FP.lowerCall(_env, callNode)
    pre.Insert(ins)
    var locals *LnsList2_[LnsInt]
    locals = NewLnsList2_[LnsInt]([]LnsInt{})
    for _, _retType := range( retTypes.Items ) {
        retType := _retType
        locals.Insert(LnsInt(self.FP.tempLocal(_env, self.FP.getRep(_env, retType))))
    }
    {
        var _forFrom0 LnsInt = locals.Len()
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
            pre.Insert(convWasm_ls_3_(_env, locals.GetAt(index), WasmBin_nop(_env)))
            _forWork0 += _forDelta0
        }
    }
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    for _index, _retType := range( retTypes.Items ) {
        index := _index + 1
        retType := _retType
        vals.Insert(NewconvWasm_Val(_env, convWasm_lg_2_(_env, locals.GetAt(index)), self.FP.getRep(_env, retType), retType))
    }
    return vals
}
// 1696: decl @lune.@base.@convWasm.convFilter.decodeString
func (self *convWasm_convFilter) decodeString(_env *LnsEnv, node *Nodes_Node,txt string) string {
    var quote string
    quote = _env.GetVM().String_sub(txt,1, 1)
    if quote == "`"{
        var body string
        body = _env.GetVM().String_sub(txt,4, len(txt) - 3)
        if _env.GetVM().String_sub(body,1, 2) == "\r\n"{
            return _env.GetVM().String_sub(body,3, nil)
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( _env.GetVM().String_sub(body,1, 1) == "\n") ||
            _env.SetStackVal( _env.GetVM().String_sub(body,1, 1) == "\r") ).(bool){
            return _env.GetVM().String_sub(body,2, nil)
        }
        return body
    }
    var body string
    body = _env.GetVM().String_sub(txt,2, len(txt) - 1)
    var result string
    result = ""
    var pos LnsInt
    pos = 1
    for pos <= len(body) {
        var ch string
        ch = _env.GetVM().String_sub(body,pos, pos)
        if ch != "\\"{
            result = result + ch
            pos = pos + 1
        } else { 
            var next string
            next = _env.GetVM().String_sub(body,pos + 1, pos + 1)
            pos = pos + 2
            if _switch0 := next; _switch0 == "n" {
                result = result + "\n"
            } else if _switch0 == "t" {
                result = result + "\t"
            } else if _switch0 == "r" {
                result = result + "\r"
            } else if _switch0 == "a" {
                result = result + WasmBin_byte(_env, 7)
            } else if _switch0 == "b" {
                result = result + WasmBin_byte(_env, 8)
            } else if _switch0 == "f" {
                result = result + WasmBin_byte(_env, 12)
            } else if _switch0 == "v" {
                result = result + WasmBin_byte(_env, 11)
            } else if _switch0 == "\\" || _switch0 == "\"" || _switch0 == "'" {
                result = result + next
            } else if _switch0 == "x" {
                var hex string
                hex = _env.GetVM().String_sub(body,pos, pos + 1)
                var val LnsReal
                
                {
                    _val := Lns_tonumber(hex, 16)
                    if _val == nil{
                        self.FP.error(_env, node, _env.GetVM().String_format("illegal escape -- %s", Lns_2DDD(txt)))
                    } else {
                        val = _val.(LnsReal)
                    }
                }
                result = result + WasmBin_byte(_env, (LnsInt)(val))
                pos = pos + 2
            } else {
                var val LnsInt
                val = 0
                var count LnsInt
                count = 0
                pos = pos - 1
                for count < 3 {
                    var code LnsInt
                    
                    {
                        _code := convWasm_convExp1_9143(Lns_2DDD(_env.GetVM().String_byte(body,pos, nil)))
                        if _code == nil{
                            break
                        } else {
                            code = _code.(LnsInt)
                        }
                    }
                    if _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( code < 0x30) ||
                        _env.SetStackVal( code > 0x39) ).(bool){
                        break
                    }
                    val = val * 10 + code - 0x30
                    count = count + 1
                    pos = pos + 1
                }
                if count == 0{
                    self.FP.error(_env, node, _env.GetVM().String_format("illegal escape -- %s", Lns_2DDD(txt)))
                }
                result = result + WasmBin_byte(_env, val)
            }
        }
    }
    return result
}
// 1780: decl @lune.@base.@convWasm.convFilter.lowerLiteralString
func (self *convWasm_convFilter) lowerLiteralString(_env *LnsEnv, node *Nodes_LiteralStringNode) *WasmBin_Ins {
    var data string
    data = self.FP.decodeString(_env, &node.Nodes_Node, node.FP.Get_token(_env).Txt)
    var fmt *WasmBin_Ins
    fmt = self.rt.FP.StrLiteral(_env, data)
    {
        _param := node.FP.Get_orgParam(_env)
        if !Lns_IsNil( _param ) {
            param := _param.(*Nodes_ExpListNode)
            var pre *LnsList2_[*WasmBin_Ins]
            pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
            var vals *LnsList2_[*convWasm_Val]
            vals = self.FP.lowerValList(_env, param, pre)
            pre.Insert(WasmBin_call(_env, self.rt.FP.FnFormat(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](fmt, self.FP.buildList(_env, vals)))))
            return WasmBin_seq(_env, pre)
        }
    }
    return fmt
}
// 1793: decl @lune.@base.@convWasm.convFilter.lowerEqVals
func (self *convWasm_convFilter) lowerEqVals(_env *LnsEnv, val1 *convWasm_Val,val2 *convWasm_Val) *WasmBin_Ins {
    var rep1 *WasmBin_ValType
    rep1 = val1.FP.Get_rep(_env)
    var rep2 *WasmBin_ValType
    rep2 = val2.FP.Get_rep(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( convWasm_isScalar_15_(_env, rep1)) &&
        _env.SetStackVal( convWasm_isScalar_15_(_env, rep2)) ).(bool)){
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( rep1.FP.Get_kind(_env) == WasmBin_ValKind__I32) ||
            _env.SetStackVal( rep2.FP.Get_kind(_env) == WasmBin_ValKind__I32) ).(bool){
            return convWasm_p_4_(_env, WasmBin_opI32Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val1.FP.Get_ins(_env), val2.FP.Get_ins(_env))))
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( rep1.FP.Get_kind(_env) == WasmBin_ValKind__I64) &&
            _env.SetStackVal( rep2.FP.Get_kind(_env) == WasmBin_ValKind__I64) ).(bool)){
            return convWasm_p_4_(_env, WasmBin_opI64Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val1.FP.Get_ins(_env), val2.FP.Get_ins(_env))))
        }
        return convWasm_p_4_(_env, WasmBin_opF64Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.coerceRep(_env, val1.FP.Get_ins(_env), rep1, WasmBin_typeF64), self.FP.coerceRep(_env, val2.FP.Get_ins(_env), rep2, WasmBin_typeF64))))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( val1.FP.Get_typeInfo(_env) == Ast_builtinTypeNil) &&
        _env.SetStackVal( rep2.FP.IsRef(_env)) ).(bool)){
        return WasmBin_refIsNull(_env, val2.FP.Get_ins(_env))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( val2.FP.Get_typeInfo(_env) == Ast_builtinTypeNil) &&
        _env.SetStackVal( rep1.FP.IsRef(_env)) ).(bool)){
        return WasmBin_refIsNull(_env, val1.FP.Get_ins(_env))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( rep1.FP.EqualsType(_env, self.rt.FP.Get_strType(_env))) &&
        _env.SetStackVal( rep2.FP.EqualsType(_env, self.rt.FP.Get_strType(_env))) ).(bool)){
        return WasmBin_call(_env, self.rt.FP.FnStrEq(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val1.FP.Get_ins(_env), val2.FP.Get_ins(_env))))
    }
    return WasmBin_call(_env, self.rt.FP.FnEqAny(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.toAny(_env, val1), self.FP.toAny(_env, val2))))
}
// 1819: decl @lune.@base.@convWasm.convFilter.lowerCompare
func (self *convWasm_convFilter) lowerCompare(_env *LnsEnv, node *Nodes_Node,opTxt string,val1 *convWasm_Val,val2 *convWasm_Val) *WasmBin_Ins {
    var rep1 *WasmBin_ValType
    rep1 = val1.FP.Get_rep(_env)
    var rep2 *WasmBin_ValType
    rep2 = val2.FP.Get_rep(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( rep1.FP.Get_kind(_env) == WasmBin_ValKind__I64) ||
            _env.SetStackVal( rep1.FP.Get_kind(_env) == WasmBin_ValKind__F64) ).(bool))) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( rep2.FP.Get_kind(_env) == WasmBin_ValKind__I64) ||
            _env.SetStackVal( rep2.FP.Get_kind(_env) == WasmBin_ValKind__F64) ).(bool))) ).(bool)){
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( rep1.FP.Get_kind(_env) == WasmBin_ValKind__I64) &&
            _env.SetStackVal( rep2.FP.Get_kind(_env) == WasmBin_ValKind__I64) ).(bool)){
            return convWasm_p_4_(_env, Lns_unwrap( convWasm_i64CmpOpMap.Get(opTxt)).(LnsInt), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val1.FP.Get_ins(_env), val2.FP.Get_ins(_env))))
        }
        return convWasm_p_4_(_env, Lns_unwrap( convWasm_f64CmpOpMap.Get(opTxt)).(LnsInt), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.coerceRep(_env, val1.FP.Get_ins(_env), rep1, WasmBin_typeF64), self.FP.coerceRep(_env, val2.FP.Get_ins(_env), rep2, WasmBin_typeF64))))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( rep1.FP.EqualsType(_env, self.rt.FP.Get_strType(_env))) &&
        _env.SetStackVal( rep2.FP.EqualsType(_env, self.rt.FP.Get_strType(_env))) ).(bool)){
        return convWasm_p_4_(_env, Lns_unwrap( convWasm_strCmpOpMap.Get(opTxt)).(LnsInt), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnStrCmp(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val1.FP.Get_ins(_env), val2.FP.Get_ins(_env)))), WasmBin_i32Const(_env, 0))))
    }
    var any1 *WasmBin_Ins
    any1 = self.FP.toAny(_env, val1)
    var any2 *WasmBin_Ins
    any2 = self.FP.toAny(_env, val2)
    var less LnsInt
    less = self.rt.FP.FnLessAny(_env)
    if _switch0 := opTxt; _switch0 == "<" {
        return WasmBin_call(_env, less, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](any1, any2)))
    } else if _switch0 == ">" {
        return WasmBin_call(_env, less, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](any2, any1)))
    } else if _switch0 == "<=" {
        return convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, less, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](any2, any1))))))
    } else if _switch0 == ">=" {
        return convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, less, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](any1, any2))))))
    }
    self.FP.error(_env, node, _env.GetVM().String_format("illegal op -- %s", Lns_2DDD(opTxt)))
// insert a dummy
    return nil
}
// 1856: decl @lune.@base.@convWasm.convFilter.lowerAndOr
func (self *convWasm_convFilter) lowerAndOr(_env *LnsEnv, node *Nodes_ExpOp2Node) *convWasm_Val {
    var opTxt string
    opTxt = node.FP.Get_op(_env).Txt
    var resultType *Ast_TypeInfo
    resultType = node.FP.Get_expType(_env)
    var resultRep *WasmBin_ValType
    resultRep = self.FP.getRep(_env, resultType)
    if resultRep.FP.Get_kind(_env) == WasmBin_ValKind__I32{
        var cond1 *WasmBin_Ins
        cond1 = self.FP.lowerCond(_env, node.FP.Get_exp1(_env))
        var cond2 *WasmBin_Ins
        cond2 = self.FP.lowerCond(_env, node.FP.Get_exp2(_env))
        if opTxt == "and"{
            return NewconvWasm_Val(_env, WasmBin_ifElse(_env, cond1, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cond2)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0))), WasmBin_typeI32), WasmBin_typeI32, resultType)
        }
        return NewconvWasm_Val(_env, WasmBin_ifElse(_env, cond1, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 1))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cond2)), WasmBin_typeI32), WasmBin_typeI32, resultType)
    }
    var val1 *convWasm_Val
    val1 = self.FP.lowerVal(_env, node.FP.Get_exp1(_env))
    var valType *Ast_TypeInfo
    valType = resultType
    var valRep *WasmBin_ValType
    valRep = resultRep
    for _, _opType := range( NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](val1.FP.Get_typeInfo(_env), node.FP.Get_exp2(_env).FP.Get_expType(_env))).Items ) {
        opType := _opType
        var opRep *WasmBin_ValType
        opRep = self.FP.getRep(_env, opType)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( Lns_op_not(opRep.FP.EqualsType(_env, valRep))) &&
            _env.SetStackVal( valRep.FP.Get_kind(_env) != WasmBin_ValKind__Ref) ||
            _env.SetStackVal( valRep.FP.Get_kind(_env) == WasmBin_ValKind__Ref) &&
            _env.SetStackVal( valRep.FP.Get_heapType(_env) != WasmBin_htAny) &&
            _env.SetStackVal( convWasm_isScalar_15_(_env, opRep)) ).(bool){
            valType = Ast_builtinTypeStem_
            valRep = WasmBin_typeAnyRef
        }
    }
    var work LnsInt
    work = self.FP.tempLocal(_env, val1.FP.Get_rep(_env))
    var cond *WasmBin_Ins
    cond = self.FP.toCond(_env, NewconvWasm_Val(_env, convWasm_lg_2_(_env, work), val1.FP.Get_rep(_env), val1.FP.Get_typeInfo(_env)))
    var first *WasmBin_Ins
    first = self.FP.coerce(_env, convWasm_lg_2_(_env, work), val1.FP.Get_typeInfo(_env), valType)
    var second *WasmBin_Ins
    second = self.FP.lowerExpAs(_env, node.FP.Get_exp2(_env), valType)
    var ins *WasmBin_Ins
    if opTxt == "and"{
        ins = WasmBin_ifElse(_env, cond, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](second)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](first)), valRep)
    } else { 
        ins = WasmBin_ifElse(_env, cond, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](first)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](second)), valRep)
    }
    return NewconvWasm_Val(_env, WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, val1.FP.Get_ins(_env)), ins))), valRep, valType)
}
// 1898: decl @lune.@base.@convWasm.convFilter.toStr
func (self *convWasm_convFilter) toStr(_env *LnsEnv, val *convWasm_Val) *WasmBin_Ins {
    if val.FP.Get_rep(_env).FP.EqualsType(_env, self.rt.FP.Get_strType(_env)){
        return val.FP.Get_ins(_env)
    }
    if _switch0 := val.FP.Get_rep(_env).FP.Get_kind(_env); _switch0 == WasmBin_ValKind__I64 {
        return WasmBin_call(_env, self.rt.FP.FnI64ToStr(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val.FP.Get_ins(_env))))
    } else if _switch0 == WasmBin_ValKind__F64 {
        return WasmBin_call(_env, self.rt.FP.FnRealToStr(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val.FP.Get_ins(_env))))
    } else if _switch0 == WasmBin_ValKind__I32 {
        return WasmBin_ifElse(_env, val.FP.Get_ins(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.StrLiteral(_env, "true"))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.StrLiteral(_env, "false"))), self.rt.FP.Get_strType(_env))
    } else if _switch0 == WasmBin_ValKind__Ref {
        return WasmBin_call(_env, self.rt.FP.FnToStr(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val.FP.Get_ins(_env))))
    }
// insert a dummy
    return nil
}
// 1920: decl @lune.@base.@convWasm.convFilter.lowerOp2
func (self *convWasm_convFilter) lowerOp2(_env *LnsEnv, node *Nodes_ExpOp2Node) *convWasm_Val {
    var opTxt string
    opTxt = node.FP.Get_op(_env).Txt
    var resultType *Ast_TypeInfo
    resultType = node.FP.Get_expType(_env)
    if _switch0 := opTxt; _switch0 == "and" || _switch0 == "or" {
        return self.FP.lowerAndOr(_env, node)
    } else if _switch0 == ".." {
        return self.FP.newVal(_env, WasmBin_call(_env, self.rt.FP.FnStrConcat(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.toStr(_env, self.FP.lowerVal(_env, node.FP.Get_exp1(_env))), self.FP.toStr(_env, self.FP.lowerVal(_env, node.FP.Get_exp2(_env)))))), resultType)
    } else if _switch0 == "==" || _switch0 == "~=" {
        var eq *WasmBin_Ins
        eq = self.FP.lowerEqVals(_env, self.FP.lowerVal(_env, node.FP.Get_exp1(_env)), self.FP.lowerVal(_env, node.FP.Get_exp2(_env)))
        if opTxt == "=="{
            return self.FP.newVal(_env, eq, resultType)
        }
        return self.FP.newVal(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](eq))), resultType)
    } else if _switch0 == "<" || _switch0 == "<=" || _switch0 == ">" || _switch0 == ">=" {
        return self.FP.newVal(_env, self.FP.lowerCompare(_env, &node.Nodes_Node, opTxt, self.FP.lowerVal(_env, node.FP.Get_exp1(_env)), self.FP.lowerVal(_env, node.FP.Get_exp2(_env))), resultType)
    }
    var resultRep *WasmBin_ValType
    resultRep = self.FP.getRep(_env, resultType)
    if resultRep.FP.Get_kind(_env) == WasmBin_ValKind__I64{
        var exp1 *WasmBin_Ins
        exp1 = self.FP.lowerExpAs(_env, node.FP.Get_exp1(_env), Ast_builtinTypeInt)
        var exp2 *WasmBin_Ins
        exp2 = self.FP.lowerExpAs(_env, node.FP.Get_exp2(_env), Ast_builtinTypeInt)
        var ins LnsAny
        ins = nil
        if _switch1 := opTxt; _switch1 == "+" {
            ins = convWasm_p_4_(_env, WasmBin_opI64Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch1 == "-" {
            ins = convWasm_p_4_(_env, WasmBin_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch1 == "*" {
            ins = convWasm_p_4_(_env, WasmBin_opI64Mul, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch1 == "/" {
            ins = WasmBin_call(_env, self.FP.fnDivI64(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch1 == "%" {
            ins = WasmBin_call(_env, self.FP.fnModI64(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch1 == "&" {
            ins = convWasm_p_4_(_env, WasmBin_opI64And, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch1 == "|" {
            ins = convWasm_p_4_(_env, WasmBin_opI64Or, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch1 == "~" {
            ins = convWasm_p_4_(_env, WasmBin_opI64Xor, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch1 == "|<<" {
            ins = WasmBin_call(_env, self.FP.fnShift(_env, true), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch1 == "|>>" {
            ins = WasmBin_call(_env, self.FP.fnShift(_env, false), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        }
        if ins != nil{
            ins_82 := ins.(*WasmBin_Ins)
            return NewconvWasm_Val(_env, ins_82, resultRep, resultType)
        }
    }
    if resultRep.FP.Get_kind(_env) == WasmBin_ValKind__F64{
        var exp1 *WasmBin_Ins
        exp1 = self.FP.lowerExpAs(_env, node.FP.Get_exp1(_env), Ast_builtinTypeReal)
        var exp2 *WasmBin_Ins
        exp2 = self.FP.lowerExpAs(_env, node.FP.Get_exp2(_env), Ast_builtinTypeReal)
        var ins LnsAny
        ins = nil
        if _switch2 := opTxt; _switch2 == "+" {
            ins = convWasm_p_4_(_env, WasmBin_opF64Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch2 == "-" {
            ins = convWasm_p_4_(_env, WasmBin_opF64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch2 == "*" {
            ins = convWasm_p_4_(_env, WasmBin_opF64Mul, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch2 == "/" {
            ins = convWasm_p_4_(_env, WasmBin_opF64Div, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch2 == "%" {
            ins = WasmBin_call(_env, self.FP.fnModF64(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        }
        if ins != nil{
            ins_93 := ins.(*WasmBin_Ins)
            return NewconvWasm_Val(_env, ins_93, resultRep, resultType)
        }
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support op -- %s (%s)", Lns_2DDD(opTxt, resultType.FP.GetTxt(_env, nil, nil, nil))))
// insert a dummy
    return nil
}
// 1988: decl @lune.@base.@convWasm.convFilter.lowerOp1
func (self *convWasm_convFilter) lowerOp1(_env *LnsEnv, node *Nodes_ExpOp1Node) *convWasm_Val {
    var opTxt string
    opTxt = node.FP.Get_op(_env).Txt
    var resultType *Ast_TypeInfo
    resultType = node.FP.Get_expType(_env)
    var val *convWasm_Val
    val = self.FP.lowerVal(_env, node.FP.Get_exp(_env))
    if _switch0 := opTxt; _switch0 == "-" {
        if val.FP.Get_rep(_env).FP.Get_kind(_env) == WasmBin_ValKind__I64{
            return self.FP.newVal(_env, convWasm_p_4_(_env, WasmBin_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i64Const(_env, 0), val.FP.Get_ins(_env)))), resultType)
        }
        if val.FP.Get_rep(_env).FP.Get_kind(_env) == WasmBin_ValKind__F64{
            return self.FP.newVal(_env, convWasm_p_4_(_env, WasmBin_opF64Neg, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val.FP.Get_ins(_env)))), resultType)
        }
    } else if _switch0 == "not" {
        return self.FP.newVal(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.toCond(_env, val)))), resultType)
    } else if _switch0 == "#" {
        if self.FP.isLuaval(_env, node.FP.Get_exp(_env).FP.Get_expType(_env)){
            return self.FP.newVal(_env, WasmBin_call(_env, self.rt.FP.FnLuaLen(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.toAny(_env, val)))), resultType)
        }
        if val.FP.Get_rep(_env).FP.EqualsType(_env, self.rt.FP.Get_strType(_env)){
            return self.FP.newVal(_env, convWasm_p_4_(_env, WasmBin_opI64ExtendI32U, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_arrayLen(_env, val.FP.Get_ins(_env))))), resultType)
        }
        if val.FP.Get_rep(_env).FP.EqualsType(_env, self.rt.FP.Get_listType(_env)){
            return self.FP.newVal(_env, convWasm_p_4_(_env, WasmBin_opI64ExtendI32U, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, val.FP.Get_ins(_env))))), resultType)
        }
        if val.FP.Get_rep(_env).FP.EqualsType(_env, self.rt.FP.Get_mapType(_env)){
            return self.FP.newVal(_env, convWasm_p_4_(_env, WasmBin_opI64ExtendI32U, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structGet(_env, self.rt.FP.Get_tMap(_env), WasmRt_fMapCount, val.FP.Get_ins(_env))))), resultType)
        }
    } else if _switch0 == "~" {
        if val.FP.Get_rep(_env).FP.Get_kind(_env) == WasmBin_ValKind__I64{
            return self.FP.newVal(_env, convWasm_p_4_(_env, WasmBin_opI64Xor, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val.FP.Get_ins(_env), WasmBin_i64Const(_env, -1)))), resultType)
        }
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support op -- %s (%s)", Lns_2DDD(opTxt, val.FP.Get_typeInfo(_env).FP.GetTxt(_env, nil, nil, nil))))
// insert a dummy
    return nil
}
// 2033: decl @lune.@base.@convWasm.convFilter.fnDivI64
func (self *convWasm_convFilter) fnDivI64(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("divI64")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64, WasmBin_typeI64)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64)))
    self.adapterMap.Set("divI64",index)
    var q LnsInt
    q = _func.FP.AddLocal(_env, WasmBin_typeI64)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, q, convWasm_p_4_(_env, WasmBin_opI64DivS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), convWasm_lg_2_(_env, 1))))), WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32And, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64Ne, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64RemS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), convWasm_lg_2_(_env, 1)))), WasmBin_i64Const(_env, 0)))), convWasm_p_4_(_env, WasmBin_opI32Ne, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), WasmBin_i64Const(_env, 0)))), convWasm_p_4_(_env, WasmBin_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 1), WasmBin_i64Const(_env, 0)))))))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, q, convWasm_p_4_(_env, WasmBin_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, q), WasmBin_i64Const(_env, 1))))))), nil, nil), convWasm_lg_2_(_env, q))))
    return index
}
// 2053: decl @lune.@base.@convWasm.convFilter.fnModI64
func (self *convWasm_convFilter) fnModI64(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("modI64")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64, WasmBin_typeI64)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64)))
    self.adapterMap.Set("modI64",index)
    var r LnsInt
    r = _func.FP.AddLocal(_env, WasmBin_typeI64)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, r, convWasm_p_4_(_env, WasmBin_opI64RemS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), convWasm_lg_2_(_env, 1))))), WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32And, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64Ne, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, r), WasmBin_i64Const(_env, 0)))), convWasm_p_4_(_env, WasmBin_opI32Ne, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, r), WasmBin_i64Const(_env, 0)))), convWasm_p_4_(_env, WasmBin_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 1), WasmBin_i64Const(_env, 0)))))))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, r, convWasm_p_4_(_env, WasmBin_opI64Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, r), convWasm_lg_2_(_env, 1))))))), nil, nil), convWasm_lg_2_(_env, r))))
    return index
}
// 2073: decl @lune.@base.@convWasm.convFilter.fnModF64
func (self *convWasm_convFilter) fnModF64(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("modF64")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeF64, WasmBin_typeF64)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeF64)))
    self.adapterMap.Set("modF64",index)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opF64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), convWasm_p_4_(_env, WasmBin_opF64Mul, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opF64Floor, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opF64Div, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), convWasm_lg_2_(_env, 1))))))), convWasm_lg_2_(_env, 1))))))))))
    return index
}
// 2087: decl @lune.@base.@convWasm.convFilter.fnShift
func (self *convWasm_convFilter) fnShift(_env *LnsEnv, left bool) LnsInt {
    var name string
    var code LnsInt
    var revCode LnsInt
    if left{
        name = "shiftL"
        code = WasmBin_opI64Shl
        revCode = WasmBin_opI64ShrU
    } else { 
        name = "shiftR"
        code = WasmBin_opI64ShrU
        revCode = WasmBin_opI64Shl
    }
    {
        _index := self.adapterMap.Get(name)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64, WasmBin_typeI64)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64)))
    self.adapterMap.Set(name,index)
    self.mod.FP.SetFuncName(_env, index, name)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32Or, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64LeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 1), WasmBin_i64Const(_env, -64)))), convWasm_p_4_(_env, WasmBin_opI64GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 1), WasmBin_i64Const(_env, 64))))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i64Const(_env, 0)))))), nil, nil), WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 1), WasmBin_i64Const(_env, 0)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, revCode, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), convWasm_p_4_(_env, WasmBin_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i64Const(_env, 0), convWasm_lg_2_(_env, 1)))))))))))), nil, nil), convWasm_p_4_(_env, code, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), convWasm_lg_2_(_env, 1)))))))
    return index
}
// 2119: decl @lune.@base.@convWasm.convFilter.lowerCast
func (self *convWasm_convFilter) lowerCast(_env *LnsEnv, node *Nodes_ExpCastNode) *convWasm_Val {
    var castType *Ast_TypeInfo
    castType = node.FP.Get_castType(_env)
    var val *convWasm_Val
    val = self.FP.lowerVal(_env, node.FP.Get_exp(_env))
    var castRep *WasmBin_ValType
    castRep = self.FP.getRep(_env, castType)
    if node.FP.Get_castKind(_env) == Nodes_CastKind__Normal{
        var work LnsInt
        work = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
        var anyVal *WasmBin_Ins
        anyVal = self.FP.toAny(_env, val)
        var dstType *Ast_TypeInfo
        dstType = self.FP.normType(_env, castType)
        var test *WasmBin_Ins
        if dstType.FP.Equals(_env, self.processInfo, Ast_builtinTypeInt, nil, nil){
            test = WasmBin_call(_env, self.FP.fnIsIntVal(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, work))))
        } else if dstType.FP.Equals(_env, self.processInfo, Ast_builtinTypeReal, nil, nil){
            test = convWasm_p_4_(_env, WasmBin_opI32Or, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refTest(_env, self.rt.FP.Get_tBoxF64(_env), convWasm_lg_2_(_env, work)), WasmBin_refTest(_env, self.rt.FP.Get_tBoxI64(_env), convWasm_lg_2_(_env, work)))))
        } else if dstType.FP.Equals(_env, self.processInfo, Ast_builtinTypeBool, nil, nil){
            test = WasmBin_refTest(_env, self.rt.FP.Get_tBoxBool(_env), convWasm_lg_2_(_env, work))
        } else { 
            var dstRep *WasmBin_ValType
            dstRep = self.FP.getRep(_env, dstType)
            if dstRep.FP.Get_heapType(_env) == WasmBin_htAny{
                test = convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refIsNull(_env, convWasm_lg_2_(_env, work)))))
            } else { 
                test = WasmBin_refTest(_env, dstRep.FP.Get_heapType(_env), convWasm_lg_2_(_env, work))
            }
        }
        var resultType *Ast_TypeInfo
        resultType = castType.FP.Get_nilableTypeInfo(_env)
        var resultRep *WasmBin_ValType
        resultRep = self.FP.getRep(_env, resultType)
        var okVal *WasmBin_Ins
        okVal = convWasm_lg_2_(_env, work)
        if dstType.FP.Equals(_env, self.processInfo, Ast_builtinTypeReal, nil, nil){
            okVal = self.rt.FP.BoxF64(_env, self.rt.FP.UnboxF64(_env, convWasm_lg_2_(_env, work)))
        } else if dstType.FP.Equals(_env, self.processInfo, Ast_builtinTypeInt, nil, nil){
            okVal = self.rt.FP.BoxI64(_env, convWasm_p_4_(_env, WasmBin_opI64TruncF64S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.UnboxF64(_env, convWasm_lg_2_(_env, work))))))
        }
        return NewconvWasm_Val(_env, WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, anyVal), WasmBin_ifElse(_env, test, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.coerceRep(_env, okVal, WasmBin_typeAnyRef, resultRep))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, resultRep.FP.Get_heapType(_env)))), resultRep)))), resultRep, resultType)
    }
    return NewconvWasm_Val(_env, self.FP.coerce(_env, val.FP.Get_ins(_env), val.FP.Get_typeInfo(_env), castType), castRep, castType)
}
// 2164: decl @lune.@base.@convWasm.convFilter.fnIsIntVal
func (self *convWasm_convFilter) fnIsIntVal(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("isIntVal")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI32)))
    self.adapterMap.Set("isIntVal",index)
    self.mod.FP.SetFuncName(_env, index, "isIntVal")
    var fval LnsInt
    fval = _func.FP.AddLocal(_env, WasmBin_typeF64)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxI64(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 1)))))), nil, nil), WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refTest(_env, self.rt.FP.Get_tBoxF64(_env), convWasm_lg_2_(_env, 0))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0)))))), nil, nil), convWasm_ls_3_(_env, fval, self.rt.FP.UnboxF64(_env, convWasm_lg_2_(_env, 0))), convWasm_p_4_(_env, WasmBin_opI32And, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opF64Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opF64Trunc, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, fval)))), convWasm_lg_2_(_env, fval)))), convWasm_p_4_(_env, WasmBin_opF64Lt, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opF64Abs, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, fval)))), WasmBin_f64Const(_env, 9.2e18))))))))))
    return index
}
// 2183: decl @lune.@base.@convWasm.convFilter.lowerUnwrap
func (self *convWasm_convFilter) lowerUnwrap(_env *LnsEnv, node *Nodes_ExpUnwrapNode) *convWasm_Val {
    var resultType *Ast_TypeInfo
    resultType = node.FP.Get_expType(_env)
    var resultRep *WasmBin_ValType
    resultRep = self.FP.getRep(_env, resultType)
    var val *convWasm_Val
    val = self.FP.lowerVal(_env, node.FP.Get_exp(_env))
    var work LnsInt
    work = self.FP.tempLocal(_env, val.FP.Get_rep(_env))
    var elseBody *WasmBin_Ins
    {
        _defNode := node.FP.Get_default(_env)
        if !Lns_IsNil( _defNode ) {
            defNode := _defNode.(*Nodes_Node)
            elseBody = self.FP.lowerExpAs(_env, defNode, resultType)
        } else {
            elseBody = self.rt.FP.Error(_env, self.rt.FP.StrLiteral(_env, "unwrap nil"))
        }
    }
    return NewconvWasm_Val(_env, WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, val.FP.Get_ins(_env)), WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, work)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](elseBody)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.coerce(_env, convWasm_lg_2_(_env, work), val.FP.Get_typeInfo(_env), resultType))), resultRep)))), resultRep, resultType)
}
// 2204: decl @lune.@base.@convWasm.convFilter.lowerLiteralColl
func (self *convWasm_convFilter) lowerLiteralColl(_env *LnsEnv, node *Nodes_Node,expList LnsAny,isSet bool) *convWasm_Val {
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    if expList != nil{
        expList_176 := expList.(*Nodes_ExpListNode)
        vals = self.FP.lowerValList(_env, expList_176, pre)
    }
    if isSet{
        var set LnsInt
        set = self.FP.tempLocal(_env, self.rt.FP.Get_mapType(_env))
        pre.Insert(convWasm_ls_3_(_env, set, self.rt.FP.NewMap(_env)))
        for _, _val := range( vals.Items ) {
            val := _val
            var work LnsInt
            work = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
            pre.Insert(convWasm_ls_3_(_env, work, self.FP.toAny(_env, val)))
            pre.Insert(WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, set), convWasm_lg_2_(_env, work), convWasm_lg_2_(_env, work)))))
        }
        pre.Insert(convWasm_lg_2_(_env, set))
    } else { 
        pre.Insert(self.FP.buildList(_env, vals))
    }
    return self.FP.newVal(_env, WasmBin_seq(_env, pre), node.FP.Get_expType(_env))
}
// 2227: decl @lune.@base.@convWasm.convFilter.lowerLiteralMap
func (self *convWasm_convFilter) lowerLiteralMap(_env *LnsEnv, node *Nodes_LiteralMapNode) *convWasm_Val {
    var _map LnsInt
    _map = self.FP.tempLocal(_env, self.rt.FP.Get_mapType(_env))
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, _map, self.rt.FP.NewMap(_env))))
    for _, _pair := range( node.FP.Get_pairList(_env).Items ) {
        pair := _pair
        body.Insert(WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, _map), self.FP.toAny(_env, self.FP.lowerVal(_env, pair.FP.Get_key(_env))), self.FP.toAny(_env, self.FP.lowerVal(_env, pair.FP.Get_val(_env)))))))
    }
    body.Insert(convWasm_lg_2_(_env, _map))
    return self.FP.newVal(_env, WasmBin_seq(_env, body), node.FP.Get_expType(_env))
}
// 2240: decl @lune.@base.@convWasm.convFilter.getObjType
func (self *convWasm_convFilter) getObjType(_env *LnsEnv, typeInfo *Ast_TypeInfo) *Ast_TypeInfo {
    var work *Ast_TypeInfo
    work = self.FP.normType(_env, typeInfo)
    if work.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate{
        if work.FP.HasBase(_env){
            return self.FP.getObjType(_env, work.FP.Get_baseTypeInfo(_env))
        }
        if work.FP.Get_interfaceList(_env).Len() > 0{
            return self.FP.getObjType(_env, work.FP.Get_interfaceList(_env).GetAt(1))
        }
    }
    return work
}
// 2259: decl @lune.@base.@convWasm.convFilter.getObjTypeForMethod
func (self *convWasm_convFilter) getObjTypeForMethod(_env *LnsEnv, typeInfo *Ast_TypeInfo,name string) *Ast_TypeInfo {
    var work *Ast_TypeInfo
    work = self.FP.normType(_env, typeInfo)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( work.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) &&
        _env.SetStackVal( Lns_op_not(work.FP.HasBase(_env))) ).(bool)){
        for _, _ifType := range( work.FP.Get_interfaceList(_env).Items ) {
            ifType := _ifType
            if Lns_isCondTrue( self.FP.getIfMethodType(_env, ifType, name)){
                return self.FP.getObjType(_env, ifType)
            }
        }
    }
    return self.FP.getObjType(_env, typeInfo)
}
// 2272: decl @lune.@base.@convWasm.convFilter.lowerMemberRef
func (self *convWasm_convFilter) lowerMemberRef(_env *LnsEnv, node *Nodes_Node,prefix *Nodes_Node,name string) *convWasm_Val {
    var objType *Ast_TypeInfo
    objType = self.FP.getObjType(_env, prefix.FP.Get_expType(_env))
    var info *convWasm_ClassInfo
    
    {
        _info := self.FP.getClassInfo(_env, objType)
        if _info == nil{
            self.FP.error(_env, node, _env.GetVM().String_format("not found class -- %s", Lns_2DDD(objType.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            info = _info.(*convWasm_ClassInfo)
        }
    }
    var fieldIndex LnsInt
    fieldIndex = self.FP.getFieldIndex(_env, node, info, name)
    var rep *WasmBin_ValType
    rep = info.FP.Get_fieldReps(_env).GetAt(fieldIndex - 1)
    var obj *WasmBin_Ins
    obj = self.FP.lowerExpAs(_env, prefix, info.FP.Get_typeInfo(_env))
    var memberType *Ast_TypeInfo
    memberType = self.FP.getMemberType(_env, info, name)
    return NewconvWasm_Val(_env, WasmBin_structGet(_env, info.FP.Get_structType(_env), fieldIndex, WasmBin_refCast(_env, info.FP.Get_structType(_env), obj)), rep, memberType)
}
// 2286: decl @lune.@base.@convWasm.convFilter.getMemberType
func (self *convWasm_convFilter) getMemberType(_env *LnsEnv, info *convWasm_ClassInfo,name string) *Ast_TypeInfo {
    var work LnsAny
    work = info
    for  {
        var classInfo *convWasm_ClassInfo
        
        {
            _classInfo := work
            if _classInfo == nil{
                break
            } else {
                classInfo = _classInfo.(*convWasm_ClassInfo)
            }
        }
        {
            _scope := classInfo.FP.Get_typeInfo(_env).FP.Get_scope(_env)
            if !Lns_IsNil( _scope ) {
                scope := _scope.(*Ast_Scope)
                {
                    _symbol := scope.FP.Get_symbol2SymbolInfoMap(_env).Get(name)
                    if !Lns_IsNil( _symbol ) {
                        symbol := _symbol.(*Ast_SymbolInfo)
                        return symbol.FP.Get_typeInfo(_env)
                    }
                }
            }
        }
        work = classInfo.FP.Get_parent(_env)
    }
    return Ast_builtinTypeStem_
}
// 2303: decl @lune.@base.@convWasm.convFilter.lowerRefField
func (self *convWasm_convFilter) lowerRefField(_env *LnsEnv, node *Nodes_RefFieldNode) *convWasm_Val {
    {
        _initIns := self.FP.getLazyInit(_env, node.FP.Get_prefix(_env))
        if !Lns_IsNil( _initIns ) {
            initIns := _initIns.(*WasmBin_Ins)
            var val *convWasm_Val
            val = self.FP.lowerRefFieldSub(_env, node)
            return NewconvWasm_Val(_env, WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](initIns, val.FP.Get_ins(_env)))), val.FP.Get_rep(_env), val.FP.Get_typeInfo(_env))
        }
    }
    return self.FP.lowerRefFieldSub(_env, node)
}
// 2310: decl @lune.@base.@convWasm.convFilter.lowerRefFieldSub
func (self *convWasm_convFilter) lowerRefFieldSub(_env *LnsEnv, node *Nodes_RefFieldNode) *convWasm_Val {
    var prefix *Nodes_Node
    prefix = node.FP.Get_prefix(_env)
    var fieldName string
    fieldName = node.FP.Get_field(_env).Txt
    var prefixType *Ast_TypeInfo
    prefixType = self.FP.normType(_env, prefix.FP.Get_expType(_env))
    {
        _modType := self.FP.getExtModuleType(_env, prefix.FP.Get_expType(_env))
        if !Lns_IsNil( _modType ) {
            modType := _modType.(*Ast_TypeInfo)
            return NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnLuaGetAt(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExtModuleVal(_env, &node.Nodes_Node, modType), self.rt.FP.StrLiteral(_env, fieldName)))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
        }
    }
    if self.FP.isLuaval(_env, prefix.FP.Get_expType(_env)){
        return NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnLuaGetAt(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExpAs(_env, prefix, Ast_builtinTypeStem_), self.rt.FP.StrLiteral(_env, fieldName)))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
    }
    {
        _symbol := node.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( _symbol ) {
            symbol := _symbol.(*Ast_SymbolInfo)
            var orgSym *Ast_SymbolInfo
            orgSym = symbol.FP.GetOrg(_env)
            var builtin *Builtin_BuiltinFuncType
            builtin = self.builtinFuncs
            var builtinConstMap *LnsMap2_[*Ast_SymbolInfo,LnsInt]
            builtinConstMap = NewLnsMap2_[*Ast_SymbolInfo,LnsInt]( map[*Ast_SymbolInfo]LnsInt{builtin.G__lns_runmode_Sync_sym:0,builtin.G__lns_runmode_Queue_sym:1,builtin.G__lns_runmode_Skip_sym:2,})
            {
                _val := builtinConstMap.Get(orgSym)
                if !Lns_IsNil( _val ) {
                    val := _val.(LnsInt)
                    return self.FP.newVal(_env, WasmBin_i64Const(_env, val), Ast_builtinTypeInt)
                }
            }
            if orgSym == builtin.G__lns_capability_async_sym{
                return self.FP.newVal(_env, WasmBin_i32Const(_env, 0), Ast_builtinTypeBool)
            }
            {
                __ := self.globalSym2loc.Get(orgSym)
                if !Lns_IsNil( __ ) {
                    return NewconvWasm_Val(_env, self.FP.readSym(_env, &node.Nodes_Node, orgSym), self.FP.getRep(_env, orgSym.FP.Get_typeInfo(_env)), orgSym.FP.Get_typeInfo(_env))
                }
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( orgSym.FP.Get_typeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) ||
                _env.SetStackVal( prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Enum) ).(bool){
                {
                    _enumType := Ast_EnumTypeInfoDownCastF(self.FP.normType(_env, orgSym.FP.Get_typeInfo(_env)).FP)
                    if !Lns_IsNil( _enumType ) {
                        enumType := _enumType.(*Ast_EnumTypeInfo)
                        {
                            _valInfo := enumType.FP.Get_name2EnumValInfo(_env).Get(fieldName)
                            if !Lns_IsNil( _valInfo ) {
                                valInfo := _valInfo.(*Ast_EnumValInfo)
                                return self.FP.lowerEnumVal(_env, enumType, valInfo)
                            }
                        }
                    }
                }
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( orgSym.FP.Get_name(_env) == "stdout") ||
                _env.SetStackVal( orgSym.FP.Get_name(_env) == "stderr") ).(bool){
                if prefix.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpRef(_env){
                    return NewconvWasm_Val(_env, WasmBin_globalGet(_env, self.FP.getStdStream(_env, orgSym.FP.Get_name(_env) == "stdout")), self.rt.FP.Get_objType(_env), orgSym.FP.Get_typeInfo(_env))
                }
            }
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( prefix.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpRef(_env)) &&
        _env.SetStackVal( prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Module) ).(bool)){
        var modName string
        modName = convWasm_convExp2_5156(Lns_2DDD(_env.GetVM().String_gsub(prefixType.FP.Get_rawTxt(_env),"^@", "")))
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( modName == "io") &&
            _env.SetStackVal( fieldName == "stdin") ).(bool)){
            return NewconvWasm_Val(_env, WasmBin_globalGet(_env, self.FP.getStdinGlobal(_env)), self.rt.FP.Get_objType(_env), node.FP.Get_expType(_env))
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( modName == "package") &&
            _env.SetStackVal( fieldName == "path") ).(bool)){
            return NewconvWasm_Val(_env, WasmBin_refCast(_env, self.rt.FP.Get_tStr(_env), self.rt.FP.ListItem(_env, WasmBin_call(_env, self.rt.FP.FnLuaCall(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, WasmRt_luaPackagePath), WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0))))))), 0)), self.rt.FP.Get_strType(_env), Ast_builtinTypeString)
        }
    }
    if _switch0 := prefixType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Map {
        return NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnMapGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, prefix), self.rt.FP.StrLiteral(_env, fieldName)))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
    } else if _switch0 == Ast_TypeInfoKind__Class || _switch0 == Ast_TypeInfoKind__Alternate {
        return self.FP.lowerMemberRef(_env, &node.Nodes_Node, prefix, fieldName)
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support field -- %s.%s", Lns_2DDD(prefixType.FP.GetTxt(_env, nil, nil, nil), fieldName)))
// insert a dummy
    return nil
}
// 2397: decl @lune.@base.@convWasm.convFilter.lowerEnumVal
func (self *convWasm_convFilter) lowerEnumVal(_env *LnsEnv, enumType *Ast_EnumTypeInfo,valInfo *Ast_EnumValInfo) *convWasm_Val {
    switch _matchExp0 := valInfo.FP.Get_val(_env).(type) {
    case *Ast_EnumLiteral__Int:
        val := _matchExp0.Val1
        return NewconvWasm_Val(_env, WasmBin_i64Const(_env, val), WasmBin_typeI64, &enumType.Ast_TypeInfo)
    case *Ast_EnumLiteral__Real:
        val := _matchExp0.Val1
        return NewconvWasm_Val(_env, WasmBin_f64Const(_env, val), WasmBin_typeF64, &enumType.Ast_TypeInfo)
    case *Ast_EnumLiteral__Str:
        val := _matchExp0.Val1
        return NewconvWasm_Val(_env, self.rt.FP.StrLiteral(_env, val), self.rt.FP.Get_strType(_env), &enumType.Ast_TypeInfo)
    }
// insert a dummy
    return nil
}
// 2412: decl @lune.@base.@convWasm.convFilter.getStdinGlobal
func (self *convWasm_convFilter) getStdinGlobal(_env *LnsEnv) LnsInt {
    if self.stdinGlobal < 0{
        self.stdinGlobal = self.mod.FP.AddGlobal(_env, self.rt.FP.Get_objType(_env), WasmBin_refNull(_env, self.rt.FP.Get_tObj(_env)))
        self.prologue.Insert(WasmBin_globalSet(_env, self.stdinGlobal, WasmBin_call(_env, self.rt.FP.FnIoStdin(_env), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))))
    }
    return self.stdinGlobal
}
// 2421: decl @lune.@base.@convWasm.convFilter.getStdStream
func (self *convWasm_convFilter) getStdStream(_env *LnsEnv, isStdout bool) LnsInt {
    if self.stdoutGlobal < 0{
        self.stdoutGlobal = self.mod.FP.AddGlobal(_env, self.rt.FP.Get_objType(_env), WasmBin_refNull(_env, self.rt.FP.Get_tObj(_env)))
        self.stderrGlobal = self.mod.FP.AddGlobal(_env, self.rt.FP.Get_objType(_env), WasmBin_refNull(_env, self.rt.FP.Get_tObj(_env)))
        var vt *WasmBin_Ins
        vt = WasmBin_structNew(_env, self.rt.FP.Get_tVT(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, self.rt.FP.Get_tAnyArr(_env)))))
        self.prologue.Insert(WasmBin_globalSet(_env, self.stdoutGlobal, WasmBin_structNew(_env, self.rt.FP.Get_tObj(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), vt)))))
        self.prologue.Insert(WasmBin_globalSet(_env, self.stderrGlobal, WasmBin_structNew(_env, self.rt.FP.Get_tObj(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), vt)))))
    }
    if isStdout{
        return self.stdoutGlobal
    }
    return self.stderrGlobal
}
// 2438: decl @lune.@base.@convWasm.convFilter.lowerRefItem
func (self *convWasm_convFilter) lowerRefItem(_env *LnsEnv, node *Nodes_ExpRefItemNode) *convWasm_Val {
    var valType *Ast_TypeInfo
    valType = self.FP.normType(_env, node.FP.Get_val(_env).FP.Get_expType(_env))
    if self.FP.isLuaval(_env, node.FP.Get_val(_env).FP.Get_expType(_env)){
        var keyIns *WasmBin_Ins
        {
            _index := node.FP.Get_index(_env)
            if !Lns_IsNil( _index ) {
                index := _index.(*Nodes_Node)
                keyIns = self.FP.lowerExpAs(_env, index, Ast_builtinTypeStem_)
            } else {
                keyIns = self.rt.FP.StrLiteral(_env, Lns_unwrap( node.FP.Get_symbol(_env)).(string))
            }
        }
        return NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnLuaGetAt(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExpAs(_env, node.FP.Get_val(_env), Ast_builtinTypeStem_), keyIns))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
    }
    var index *Nodes_Node
    
    {
        _index := node.FP.Get_index(_env)
        if _index == nil{
            {
                _symbol := node.FP.Get_symbol(_env)
                if !Lns_IsNil( _symbol ) {
                    symbol := _symbol.(string)
                    if valType.FP.Get_kind(_env) == Ast_TypeInfoKind__Map{
                        return NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnMapGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, node.FP.Get_val(_env)), self.rt.FP.StrLiteral(_env, symbol)))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
                    }
                    if self.FP.isType(_env, valType, Ast_builtinTypeStem){
                        return NewconvWasm_Val(_env, WasmBin_call(_env, self.FP.fnStemIndex(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, node.FP.Get_val(_env)), self.rt.FP.StrLiteral(_env, symbol)))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
                    }
                }
            }
            self.FP.notSupport(_env, &node.Nodes_Node)
        } else {
            index = _index.(*Nodes_Node)
        }
    }
    if _switch0 := valType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        return NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, node.FP.Get_val(_env)), self.FP.lowerExpAs(_env, index, Ast_builtinTypeInt)))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
    } else if _switch0 == Ast_TypeInfoKind__Map {
        return NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnMapGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, node.FP.Get_val(_env)), self.FP.toAny(_env, self.FP.lowerVal(_env, index))))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
    }
    if self.FP.isType(_env, valType, Ast_builtinTypeStem){
        return NewconvWasm_Val(_env, WasmBin_call(_env, self.FP.fnStemIndex(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, node.FP.Get_val(_env)), self.FP.toAny(_env, self.FP.lowerVal(_env, index))))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
    }
    if self.FP.isType(_env, valType, Ast_builtinTypeString){
        var strVal LnsInt
        strVal = self.FP.tempLocal(_env, self.rt.FP.Get_strType(_env))
        return self.FP.newVal(_env, WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, strVal, self.FP.lowerExp(_env, node.FP.Get_val(_env))), convWasm_p_4_(_env, WasmBin_opI64ExtendI32U, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_arrayGetU(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, strVal), convWasm_p_4_(_env, WasmBin_opI32WrapI64, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExpAs(_env, index, Ast_builtinTypeInt), WasmBin_i64Const(_env, 1)))))))))))))), Ast_builtinTypeInt)
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support index -- %s", Lns_2DDD(valType.FP.GetTxt(_env, nil, nil, nil))))
// insert a dummy
    return nil
}
// 2503: decl @lune.@base.@convWasm.convFilter.lowerNewAlgeVal
func (self *convWasm_convFilter) lowerNewAlgeVal(_env *LnsEnv, node *Nodes_NewAlgeValNode) *convWasm_Val {
    var algeType *Ast_TypeInfo
    algeType = self.FP.getGenSrc(_env, &node.FP.Get_valInfo(_env).FP.Get_algeType(_env).Ast_TypeInfo)
    var info *convWasm_AlgeInfo
    
    {
        _info := self.FP.getAlgeInfo(_env, algeType)
        if _info == nil{
            self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not found alge -- %s", Lns_2DDD(algeType.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            info = _info.(*convWasm_AlgeInfo)
        }
    }
    var name string
    name = node.FP.Get_valInfo(_env).FP.Get_name(_env)
    {
        _globalIdx := info.FP.Get_name2global(_env).Get(name)
        if !Lns_IsNil( _globalIdx ) {
            globalIdx := _globalIdx.(LnsInt)
            return NewconvWasm_Val(_env, WasmBin_globalGet(_env, globalIdx), WasmBin_refType(_env, info.FP.Get_baseType(_env)), algeType)
        }
    }
    var valType LnsInt
    valType = Lns_unwrap( info.FP.Get_name2type(_env).Get(name)).(LnsInt)
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), WasmBin_i32Const(_env, Lns_unwrap( info.FP.Get_name2tag(_env).Get(name)).(LnsInt))))
    for _index, _param := range( node.FP.Get_paramList(_env).Items ) {
        index := _index + 1
        param := _param
        args.Insert(self.FP.lowerExpAs(_env, param, node.FP.Get_valInfo(_env).FP.Get_typeList(_env).GetAt(index)))
    }
    return NewconvWasm_Val(_env, WasmBin_structNew(_env, valType, args), WasmBin_refType(_env, info.FP.Get_baseType(_env)), algeType)
}
// 2521: decl @lune.@base.@convWasm.convFilter.lowerAnonymousFunc
func (self *convWasm_convFilter) lowerAnonymousFunc(_env *LnsEnv, node *Nodes_DeclFuncNode) *convWasm_Val {
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_expType(_env)
    var funcIndex LnsInt
    funcIndex = self.FP.registerFunc(_env, node.FP.Get_declInfo(_env), funcType)
    return self.FP.newVal(_env, self.FP.compileFunc(_env, &node.Nodes_Node, node.FP.Get_declInfo(_env), funcType, funcIndex), funcType)
}
// 2531: decl @lune.@base.@convWasm.convFilter.getNilAccPrefix
func (self *convWasm_convFilter) getNilAccPrefix(_env *LnsEnv, node *Nodes_Node) LnsAny {
    if self.nilAccDoneSet.Has(node){
        return nil
    }
    {
        _refNode := Nodes_RefFieldNodeDownCastF(node.FP)
        if !Lns_IsNil( _refNode ) {
            refNode := _refNode.(*Nodes_RefFieldNode)
            if refNode.FP.Get_nilAccess(_env){
                return refNode.FP.Get_prefix(_env)
            }
        } else {
            {
                _getNode := Nodes_GetFieldNodeDownCastF(node.FP)
                if !Lns_IsNil( _getNode ) {
                    getNode := _getNode.(*Nodes_GetFieldNode)
                    if getNode.FP.Get_nilAccess(_env){
                        return getNode.FP.Get_prefix(_env)
                    }
                } else {
                    {
                        _itemNode := Nodes_ExpRefItemNodeDownCastF(node.FP)
                        if !Lns_IsNil( _itemNode ) {
                            itemNode := _itemNode.(*Nodes_ExpRefItemNode)
                            if itemNode.FP.Get_nilAccess(_env){
                                return itemNode.FP.Get_val(_env)
                            }
                        } else {
                            {
                                _callNode := Nodes_ExpCallNodeDownCastF(node.FP)
                                if !Lns_IsNil( _callNode ) {
                                    callNode := _callNode.(*Nodes_ExpCallNode)
                                    if callNode.FP.Get_nilAccess(_env){
                                        {
                                            _funcNode := Nodes_RefFieldNodeDownCastF(callNode.FP.Get_func(_env).FP)
                                            if !Lns_IsNil( _funcNode ) {
                                                funcNode := _funcNode.(*Nodes_RefFieldNode)
                                                if funcNode.FP.Get_nilAccess(_env){
                                                    return funcNode.FP.Get_prefix(_env)
                                                }
                                            }
                                        }
                                        return callNode.FP.Get_func(_env)
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    return nil
}
// 2565: decl @lune.@base.@convWasm.convFilter.lowerNilAcc
func (self *convWasm_convFilter) lowerNilAcc(_env *LnsEnv, node *Nodes_Node,prefix *Nodes_Node) *convWasm_Val {
    var prefixVal *convWasm_Val
    prefixVal = self.FP.lowerVal(_env, prefix)
    var work LnsInt
    work = self.FP.tempLocal(_env, prefixVal.FP.Get_rep(_env))
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, prefixVal.FP.Get_ins(_env))))
    self.node2val.Set(prefix,NewconvWasm_Val(_env, convWasm_lg_2_(_env, work), prefixVal.FP.Get_rep(_env), prefixVal.FP.Get_typeInfo(_env)))
    self.nilAccDoneSet.Add(node)
    {
        _callNode := Nodes_ExpCallNodeDownCastF(node.FP)
        if !Lns_IsNil( _callNode ) {
            callNode := _callNode.(*Nodes_ExpCallNode)
            self.nilAccDoneSet.Add(callNode.FP.Get_func(_env))
        }
    }
    var resultType *Ast_TypeInfo
    resultType = node.FP.Get_expType(_env)
    var resultRep *WasmBin_ValType
    resultRep = self.FP.getRep(_env, resultType)
    var accessIns *WasmBin_Ins
    accessIns = self.FP.coerceVal(_env, self.FP.lowerVal(_env, node), resultType)
    self.node2val.Set(prefix,nil)
    pre.Insert(WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, work)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.getDefaultVal(_env, resultRep))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](accessIns)), resultRep))
    return self.FP.newVal(_env, WasmBin_seq(_env, pre), resultType)
}
// 2588: decl @lune.@base.@convWasm.convFilter.lowerNilAccCall
func (self *convWasm_convFilter) lowerNilAccCall(_env *LnsEnv, node *Nodes_ExpCallNode,prefix *Nodes_Node)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var prefixVal *convWasm_Val
    prefixVal = self.FP.lowerVal(_env, prefix)
    var work LnsInt
    work = self.FP.tempLocal(_env, prefixVal.FP.Get_rep(_env))
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, prefixVal.FP.Get_ins(_env))))
    self.node2val.Set(prefix,NewconvWasm_Val(_env, convWasm_lg_2_(_env, work), prefixVal.FP.Get_rep(_env), prefixVal.FP.Get_typeInfo(_env)))
    self.nilAccDoneSet.Add(&node.Nodes_Node)
    self.nilAccDoneSet.Add(node.FP.Get_func(_env))
    var callIns *WasmBin_Ins
    var callRetTypes *LnsList2_[*Ast_TypeInfo]
    callIns,callRetTypes = self.FP.lowerCall(_env, node)
    self.node2val.Set(prefix,nil)
    var retTypes *LnsList2_[*Ast_TypeInfo]
    retTypes = NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    var callLocals *LnsList2_[LnsInt]
    callLocals = NewLnsList2_[LnsInt]([]LnsInt{})
    var retLocals *LnsList2_[LnsInt]
    retLocals = NewLnsList2_[LnsInt]([]LnsInt{})
    for _, _retType := range( callRetTypes.Items ) {
        retType := _retType
        var nilableType *Ast_TypeInfo
        nilableType = retType.FP.Get_nilableTypeInfo(_env)
        retTypes.Insert(nilableType)
        callLocals.Insert(LnsInt(self.FP.tempLocal(_env, self.FP.getRep(_env, retType))))
        retLocals.Insert(LnsInt(self.FP.tempLocal(_env, self.FP.getRep(_env, nilableType))))
    }
    var thenBody *LnsList2_[*WasmBin_Ins]
    thenBody = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](callIns))
    var elseBody *LnsList2_[*WasmBin_Ins]
    elseBody = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    {
        var _forFrom0 LnsInt = callLocals.Len()
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
            thenBody.Insert(convWasm_ls_3_(_env, callLocals.GetAt(index), WasmBin_nop(_env)))
            _forWork0 += _forDelta0
        }
    }
    for _index, _retType := range( callRetTypes.Items ) {
        index := _index + 1
        retType := _retType
        thenBody.Insert(convWasm_ls_3_(_env, retLocals.GetAt(index), self.FP.coerce(_env, convWasm_lg_2_(_env, callLocals.GetAt(index)), retType, retTypes.GetAt(index))))
        if self.FP.normType(_env, retType).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            elseBody.Insert(convWasm_ls_3_(_env, retLocals.GetAt(index), WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 1))))))
            elseBody.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, retLocals.GetAt(index)), WasmBin_refNull(_env, WasmBin_htAny)))))
        } else { 
            elseBody.Insert(convWasm_ls_3_(_env, retLocals.GetAt(index), self.FP.getDefaultVal(_env, self.FP.getRep(_env, retTypes.GetAt(index)))))
        }
    }
    pre.Insert(WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, work)), elseBody, thenBody, nil))
    for _, _retLocal := range( retLocals.Items ) {
        retLocal := _retLocal
        pre.Insert(convWasm_lg_2_(_env, retLocal))
    }
    return WasmBin_seq(_env, pre), retTypes
}
// 2635: decl @lune.@base.@convWasm.convFilter.lowerVal
func (self *convWasm_convFilter) lowerVal(_env *LnsEnv, node *Nodes_Node) *convWasm_Val {
    {
        _val := self.node2val.Get(node)
        if !Lns_IsNil( _val ) {
            val := _val.(*convWasm_Val)
            return val
        }
    }
    {
        _prefix := self.FP.getNilAccPrefix(_env, node)
        if !Lns_IsNil( _prefix ) {
            prefix := _prefix.(*Nodes_Node)
            return self.FP.lowerNilAcc(_env, node, prefix)
        }
    }
    var expType *Ast_TypeInfo
    expType = node.FP.Get_expType(_env)
    if _switch0 := node.FP.Get_kind(_env); _switch0 == Nodes_NodeKind_get_TupleConst(_env) {
        var tupleNode *Nodes_TupleConstNode
        tupleNode = Lns_unwrap( Nodes_TupleConstNodeDownCastF(node.FP)).(*Nodes_TupleConstNode)
        var pre *LnsList2_[*WasmBin_Ins]
        pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        var vals *LnsList2_[*convWasm_Val]
        vals = self.FP.lowerValList(_env, tupleNode.FP.Get_expList(_env), pre)
        pre.Insert(self.FP.buildList(_env, vals))
        return self.FP.newVal(_env, WasmBin_seq(_env, pre), expType)
    } else if _switch0 == Nodes_NodeKind_get_LuneKind(_env) {
        var kindNode *Nodes_LuneKindNode
        kindNode = Lns_unwrap( Nodes_LuneKindNodeDownCastF(node.FP)).(*Nodes_LuneKindNode)
        var exp *Nodes_Node
        exp = kindNode.FP.Get_exp(_env)
        {
            _castNode := Nodes_ExpCastNodeDownCastF(exp.FP)
            if !Lns_IsNil( _castNode ) {
                castNode := _castNode.(*Nodes_ExpCastNode)
                if castNode.FP.Get_castKind(_env) == Nodes_CastKind__Implicit{
                    exp = castNode.FP.Get_exp(_env)
                }
            }
        }
        return self.FP.newVal(_env, WasmBin_i64Const(_env, LnsInt(exp.FP.Get_expType(_env).FP.Get_kind(_env))), Ast_builtinTypeInt)
    } else if _switch0 == Nodes_NodeKind_get_Boxing(_env) {
        var boxNode *Nodes_BoxingNode
        boxNode = Lns_unwrap( Nodes_BoxingNodeDownCastF(node.FP)).(*Nodes_BoxingNode)
        var pre *LnsList2_[*WasmBin_Ins]
        pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        pre.Insert(self.FP.buildList(_env, NewLnsList2_[*convWasm_Val](Lns_2DDDGen[*convWasm_Val](self.FP.lowerVal(_env, boxNode.FP.Get_src(_env))))))
        return self.FP.newVal(_env, WasmBin_seq(_env, pre), expType)
    } else if _switch0 == Nodes_NodeKind_get_Unboxing(_env) {
        var unboxNode *Nodes_UnboxingNode
        unboxNode = Lns_unwrap( Nodes_UnboxingNodeDownCastF(node.FP)).(*Nodes_UnboxingNode)
        return NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExpAs(_env, unboxNode.FP.Get_src(_env), unboxNode.FP.Get_src(_env).FP.Get_expType(_env)), WasmBin_i64Const(_env, 1)))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
    } else if _switch0 == Nodes_NodeKind_get_ExpExpandTuple(_env) {
        var expandNode *Nodes_ExpExpandTupleNode
        expandNode = Lns_unwrap( Nodes_ExpExpandTupleNodeDownCastF(node.FP)).(*Nodes_ExpExpandTupleNode)
        return NewconvWasm_Val(_env, self.FP.lowerExpAs(_env, expandNode.FP.Get_exp(_env), expandNode.FP.Get_exp(_env).FP.Get_expType(_env)), self.rt.FP.Get_listType(_env), Ast_builtinTypeDDD)
    } else if _switch0 == Nodes_NodeKind_get_LiteralInt(_env) {
        var literal *Nodes_LiteralIntNode
        literal = Lns_unwrap( Nodes_LiteralIntNodeDownCastF(node.FP)).(*Nodes_LiteralIntNode)
        return self.FP.newVal(_env, WasmBin_i64Const(_env, literal.FP.Get_num(_env)), Ast_builtinTypeInt)
    } else if _switch0 == Nodes_NodeKind_get_LiteralChar(_env) {
        var literal *Nodes_LiteralCharNode
        literal = Lns_unwrap( Nodes_LiteralCharNodeDownCastF(node.FP)).(*Nodes_LiteralCharNode)
        return self.FP.newVal(_env, WasmBin_i64Const(_env, literal.FP.Get_num(_env)), Ast_builtinTypeInt)
    } else if _switch0 == Nodes_NodeKind_get_LiteralReal(_env) {
        var literal *Nodes_LiteralRealNode
        literal = Lns_unwrap( Nodes_LiteralRealNodeDownCastF(node.FP)).(*Nodes_LiteralRealNode)
        return self.FP.newVal(_env, WasmBin_f64Const(_env, literal.FP.Get_num(_env)), Ast_builtinTypeReal)
    } else if _switch0 == Nodes_NodeKind_get_LiteralBool(_env) {
        var literal *Nodes_LiteralBoolNode
        literal = Lns_unwrap( Nodes_LiteralBoolNodeDownCastF(node.FP)).(*Nodes_LiteralBoolNode)
        if literal.FP.Get_token(_env).Txt == "true"{
            return self.FP.newVal(_env, WasmBin_i32Const(_env, 1), Ast_builtinTypeBool)
        }
        return self.FP.newVal(_env, WasmBin_i32Const(_env, 0), Ast_builtinTypeBool)
    } else if _switch0 == Nodes_NodeKind_get_LiteralNil(_env) || _switch0 == Nodes_NodeKind_get_Abbr(_env) {
        return NewconvWasm_Val(_env, WasmBin_refNull(_env, WasmBin_htAny), WasmBin_typeAnyRef, Ast_builtinTypeNil)
    } else if _switch0 == Nodes_NodeKind_get_LiteralString(_env) {
        return self.FP.newVal(_env, self.FP.lowerLiteralString(_env, Lns_unwrap( Nodes_LiteralStringNodeDownCastF(node.FP)).(*Nodes_LiteralStringNode)), Ast_builtinTypeString)
    } else if _switch0 == Nodes_NodeKind_get_LiteralList(_env) {
        return self.FP.lowerLiteralColl(_env, node, (Lns_unwrap( Nodes_LiteralListNodeDownCastF(node.FP)).(*Nodes_LiteralListNode)).FP.Get_expList(_env), false)
    } else if _switch0 == Nodes_NodeKind_get_LiteralArray(_env) {
        return self.FP.lowerLiteralColl(_env, node, (Lns_unwrap( Nodes_LiteralArrayNodeDownCastF(node.FP)).(*Nodes_LiteralArrayNode)).FP.Get_expList(_env), false)
    } else if _switch0 == Nodes_NodeKind_get_LiteralSet(_env) {
        return self.FP.lowerLiteralColl(_env, node, (Lns_unwrap( Nodes_LiteralSetNodeDownCastF(node.FP)).(*Nodes_LiteralSetNode)).FP.Get_expList(_env), true)
    } else if _switch0 == Nodes_NodeKind_get_LiteralMap(_env) {
        return self.FP.lowerLiteralMap(_env, Lns_unwrap( Nodes_LiteralMapNodeDownCastF(node.FP)).(*Nodes_LiteralMapNode))
    } else if _switch0 == Nodes_NodeKind_get_ExpRef(_env) {
        var refNode *Nodes_ExpRefNode
        refNode = Lns_unwrap( Nodes_ExpRefNodeDownCastF(node.FP)).(*Nodes_ExpRefNode)
        var symbol *Ast_SymbolInfo
        symbol = refNode.FP.Get_symbolInfo(_env).FP.GetOrg(_env)
        return NewconvWasm_Val(_env, self.FP.readSym(_env, node, symbol), self.FP.getRep(_env, symbol.FP.Get_typeInfo(_env)), symbol.FP.Get_typeInfo(_env))
    } else if _switch0 == Nodes_NodeKind_get_ExpParen(_env) {
        return self.FP.lowerVal(_env, (Lns_unwrap( Nodes_ExpParenNodeDownCastF(node.FP)).(*Nodes_ExpParenNode)).FP.Get_exp(_env))
    } else if _switch0 == Nodes_NodeKind_get_ExpOp2(_env) {
        return self.FP.lowerOp2(_env, Lns_unwrap( Nodes_ExpOp2NodeDownCastF(node.FP)).(*Nodes_ExpOp2Node))
    } else if _switch0 == Nodes_NodeKind_get_ExpOp1(_env) {
        return self.FP.lowerOp1(_env, Lns_unwrap( Nodes_ExpOp1NodeDownCastF(node.FP)).(*Nodes_ExpOp1Node))
    } else if _switch0 == Nodes_NodeKind_get_ExpCast(_env) {
        return self.FP.lowerCast(_env, Lns_unwrap( Nodes_ExpCastNodeDownCastF(node.FP)).(*Nodes_ExpCastNode))
    } else if _switch0 == Nodes_NodeKind_get_ExpUnwrap(_env) {
        return self.FP.lowerUnwrap(_env, Lns_unwrap( Nodes_ExpUnwrapNodeDownCastF(node.FP)).(*Nodes_ExpUnwrapNode))
    } else if _switch0 == Nodes_NodeKind_get_ExpCall(_env) {
        var ins *WasmBin_Ins
        var retTypes *LnsList2_[*Ast_TypeInfo]
        ins,retTypes = self.FP.lowerCall(_env, Lns_unwrap( Nodes_ExpCallNodeDownCastF(node.FP)).(*Nodes_ExpCallNode))
        if retTypes.Len() == 0{
            return NewconvWasm_Val(_env, ins, WasmBin_typeI32, Ast_builtinTypeNone)
        }
        if retTypes.Len() == 1{
            if self.FP.normType(_env, retTypes.GetAt(1)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
                return NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](ins, WasmBin_i64Const(_env, 1)))), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
            }
            return self.FP.newVal(_env, ins, retTypes.GetAt(1))
        }
        var pre *LnsList2_[*WasmBin_Ins]
        pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        var vals *LnsList2_[*convWasm_Val]
        vals = self.FP.lowerMulti(_env, node, pre)
        pre.Insert(vals.GetAt(1).FP.Get_ins(_env))
        return NewconvWasm_Val(_env, WasmBin_seq(_env, pre), vals.GetAt(1).FP.Get_rep(_env), vals.GetAt(1).FP.Get_typeInfo(_env))
    } else if _switch0 == Nodes_NodeKind_get_ExpMultiTo1(_env) {
        var pre *LnsList2_[*WasmBin_Ins]
        pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        var vals *LnsList2_[*convWasm_Val]
        vals = self.FP.lowerMulti(_env, (Lns_unwrap( Nodes_ExpMultiTo1NodeDownCastF(node.FP)).(*Nodes_ExpMultiTo1Node)).FP.Get_exp(_env), pre)
        if self.FP.normType(_env, vals.GetAt(1).FP.Get_typeInfo(_env)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            pre.Insert(WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](vals.GetAt(1).FP.Get_ins(_env), WasmBin_i64Const(_env, 1)))))
            return NewconvWasm_Val(_env, WasmBin_seq(_env, pre), WasmBin_typeAnyRef, Ast_builtinTypeStem_)
        }
        pre.Insert(vals.GetAt(1).FP.Get_ins(_env))
        return NewconvWasm_Val(_env, WasmBin_seq(_env, pre), vals.GetAt(1).FP.Get_rep(_env), vals.GetAt(1).FP.Get_typeInfo(_env))
    } else if _switch0 == Nodes_NodeKind_get_ExpToDDD(_env) {
        var dddNode *Nodes_ExpToDDDNode
        dddNode = Lns_unwrap( Nodes_ExpToDDDNodeDownCastF(node.FP)).(*Nodes_ExpToDDDNode)
        var pre *LnsList2_[*WasmBin_Ins]
        pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        var vals *LnsList2_[*convWasm_Val]
        vals = self.FP.lowerValList(_env, dddNode.FP.Get_expList(_env), pre)
        pre.Insert(self.FP.buildList(_env, vals))
        return NewconvWasm_Val(_env, WasmBin_seq(_env, pre), self.rt.FP.Get_listType(_env), expType)
    } else if _switch0 == Nodes_NodeKind_get_ExpNew(_env) {
        return self.FP.lowerNew(_env, Lns_unwrap( Nodes_ExpNewNodeDownCastF(node.FP)).(*Nodes_ExpNewNode))
    } else if _switch0 == Nodes_NodeKind_get_RefField(_env) {
        return self.FP.lowerRefField(_env, Lns_unwrap( Nodes_RefFieldNodeDownCastF(node.FP)).(*Nodes_RefFieldNode))
    } else if _switch0 == Nodes_NodeKind_get_GetField(_env) {
        return self.FP.lowerGetField(_env, Lns_unwrap( Nodes_GetFieldNodeDownCastF(node.FP)).(*Nodes_GetFieldNode))
    } else if _switch0 == Nodes_NodeKind_get_ExpRefItem(_env) {
        return self.FP.lowerRefItem(_env, Lns_unwrap( Nodes_ExpRefItemNodeDownCastF(node.FP)).(*Nodes_ExpRefItemNode))
    } else if _switch0 == Nodes_NodeKind_get_ExpOmitEnum(_env) {
        var omitNode *Nodes_ExpOmitEnumNode
        omitNode = Lns_unwrap( Nodes_ExpOmitEnumNodeDownCastF(node.FP)).(*Nodes_ExpOmitEnumNode)
        return self.FP.lowerEnumVal(_env, omitNode.FP.Get_enumTypeInfo(_env), omitNode.FP.Get_valInfo(_env))
    } else if _switch0 == Nodes_NodeKind_get_NewAlgeVal(_env) {
        return self.FP.lowerNewAlgeVal(_env, Lns_unwrap( Nodes_NewAlgeValNodeDownCastF(node.FP)).(*Nodes_NewAlgeValNode))
    } else if _switch0 == Nodes_NodeKind_get_DeclFunc(_env) {
        return self.FP.lowerAnonymousFunc(_env, Lns_unwrap( Nodes_DeclFuncNodeDownCastF(node.FP)).(*Nodes_DeclFuncNode))
    } else if _switch0 == Nodes_NodeKind_get_ExpCallSuper(_env) {
        var ins *WasmBin_Ins
        var retTypes *LnsList2_[*Ast_TypeInfo]
        ins,retTypes = self.FP.lowerCallSuper(_env, Lns_unwrap( Nodes_ExpCallSuperNodeDownCastF(node.FP)).(*Nodes_ExpCallSuperNode))
        if retTypes.Len() == 0{
            return NewconvWasm_Val(_env, ins, WasmBin_typeI32, Ast_builtinTypeNone)
        }
        return self.FP.newVal(_env, ins, retTypes.GetAt(1))
    } else if _switch0 == Nodes_NodeKind_get_ExpMacroExp(_env) {
        var macroNode *Nodes_ExpMacroExpNode
        macroNode = Lns_unwrap( Nodes_ExpMacroExpNodeDownCastF(node.FP)).(*Nodes_ExpMacroExpNode)
        if macroNode.FP.Get_stmtList(_env).Len() == 1{
            var stmt *Nodes_Node
            stmt = macroNode.FP.Get_stmtList(_env).GetAt(1)
            {
                _stmtExp := Nodes_StmtExpNodeDownCastF(stmt.FP)
                if !Lns_IsNil( _stmtExp ) {
                    stmtExp := _stmtExp.(*Nodes_StmtExpNode)
                    return self.FP.lowerVal(_env, stmtExp.FP.Get_exp(_env))
                }
            }
            return self.FP.lowerVal(_env, stmt)
        }
    }
    self.FP.notSupport(_env, node)
// insert a dummy
    return nil
}
// 2830: decl @lune.@base.@convWasm.convFilter.lowerArgs
func (self *convWasm_convFilter) lowerArgs(_env *LnsEnv, argList LnsAny,paramTypes *LnsList2_[*Ast_TypeInfo],pre *LnsList2_[*WasmBin_Ins]) *LnsList2_[*WasmBin_Ins] {
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    if argList != nil{
        argList_435 := argList.(*Nodes_ExpListNode)
        vals = self.FP.lowerValList(_env, argList_435, pre)
    }
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _index, _paramType := range( paramTypes.Items ) {
        index := _index + 1
        paramType := _paramType
        if self.FP.normType(_env, paramType).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD{
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( vals.Len() == index) &&
                _env.SetStackVal( self.FP.normType(_env, vals.GetAt(index).FP.Get_typeInfo(_env)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
                args.Insert(vals.GetAt(index).FP.Get_ins(_env))
            } else { 
                var rest *LnsList2_[*convWasm_Val]
                rest = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
                {
                    var _forFrom0 LnsInt = index
                    var _forTo0 LnsInt = vals.Len()
                    for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                        work := _forWork0
                        rest.Insert(vals.GetAt(work))
                    }
                }
                args.Insert(self.FP.buildList(_env, rest))
            }
            break
        }
        if index <= vals.Len(){
            args.Insert(self.FP.coerceVal(_env, vals.GetAt(index), paramType))
        } else { 
            args.Insert(self.FP.getDefaultVal(_env, self.FP.getRep(_env, paramType)))
        }
    }
    return args
}
// 2868: decl @lune.@base.@convWasm.convFilter.callClosure
func (self *convWasm_convFilter) callClosure(_env *LnsEnv, cloVal *WasmBin_Ins,funcType *Ast_TypeInfo,argList LnsAny,pre *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    var info *WasmRt_CloInfo
    info = self.FP.getFuncCloInfo(_env, funcType)
    var work LnsInt
    work = self.FP.tempLocal(_env, WasmBin_refType(_env, info.FP.Get_cloType(_env)))
    pre.Insert(convWasm_ls_3_(_env, work, WasmBin_refCast(_env, info.FP.Get_cloType(_env), cloVal)))
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, work)))
    convWasm_append_17_(_env, args, self.FP.lowerArgs(_env, argList, self.FP.normType(_env, funcType).FP.Get_argTypeInfoList(_env), pre))
    args.Insert(WasmBin_structGet(_env, info.FP.Get_cloType(_env), WasmRt_fCloCode, convWasm_lg_2_(_env, work)))
    return WasmBin_callRef(_env, info.FP.Get_funcType(_env), args)
}
// 2881: decl @lune.@base.@convWasm.convFilter.lowerPrint
func (self *convWasm_convFilter) lowerPrint(_env *LnsEnv, node *Nodes_ExpCallNode) *WasmBin_Ins {
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    {
        _argList := node.FP.Get_argList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            if argList.FP.Get_expList(_env).Len() == 1{
                {
                    _dddNode := Nodes_ExpToDDDNodeDownCastF(argList.FP.Get_expList(_env).GetAt(1).FP)
                    if !Lns_IsNil( _dddNode ) {
                        dddNode := _dddNode.(*Nodes_ExpToDDDNode)
                        vals = self.FP.lowerValList(_env, dddNode.FP.Get_expList(_env), pre)
                    } else {
                        vals = self.FP.lowerValList(_env, argList, pre)
                    }
                }
            } else { 
                vals = self.FP.lowerValList(_env, argList, pre)
            }
        }
    }
    pre.Insert(WasmBin_call(_env, self.rt.FP.FnPrintList(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.buildList(_env, vals)))))
    return WasmBin_seq(_env, pre)
}
// 2900: decl @lune.@base.@convWasm.convFilter.lowerCallRef
func (self *convWasm_convFilter) lowerCallRef(_env *LnsEnv, node *Nodes_ExpCallNode,refNode *Nodes_ExpRefNode,pre *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var symbol *Ast_SymbolInfo
    symbol = self.FP.resolveSym(_env, refNode.FP.Get_symbolInfo(_env).FP.GetOrg(_env))
    var funcType *Ast_TypeInfo
    funcType = symbol.FP.Get_typeInfo(_env)
    var builtin *Builtin_BuiltinFuncType
    builtin = self.builtinFuncs
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns_print, nil, nil){
        return self.FP.lowerPrint(_env, node), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns_error, nil, nil){
        var args *LnsList2_[*WasmBin_Ins]
        args = self.FP.lowerArgs(_env, node.FP.Get_argList(_env), funcType.FP.Get_argTypeInfoList(_env), pre)
        pre.Insert(self.rt.FP.Error(_env, args.GetAt(1)))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns_tostring, nil, nil){
        var vals *LnsList2_[*convWasm_Val]
        vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                vals = self.FP.lowerValList(_env, argList, pre)
            }
        }
        pre.Insert(self.FP.toStr(_env, vals.GetAt(1)))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString))
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns___run, nil, nil){
        var argList *Nodes_ExpListNode
        argList = Lns_unwrap( node.FP.Get_argList(_env)).(*Nodes_ExpListNode)
        var mode LnsInt
        mode = self.FP.tempLocal(_env, WasmBin_typeI64)
        pre.Insert(convWasm_ls_3_(_env, mode, self.FP.lowerExpAs(_env, argList.FP.Get_expList(_env).GetAt(2), Ast_builtinTypeInt)))
        var callPre *LnsList2_[*WasmBin_Ins]
        callPre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        var callIns *WasmBin_Ins
        var retTypes *LnsList2_[*Ast_TypeInfo]
        callIns,retTypes = self.FP.lowerObjMethodCall(_env, &node.Nodes_Node, argList.FP.Get_expList(_env).GetAt(1), "run", callPre, nil)
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](callIns))
        for range( retTypes.Items ) {
            body.Insert(WasmBin_drop(_env, WasmBin_nop(_env)))
        }
        body.Insert(WasmBin_i32Const(_env, 1))
        pre.Insert(WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI64Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, mode), WasmBin_i64Const(_env, 2)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0))), body, WasmBin_typeI32))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeBool))
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns___join, nil, nil){
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns___serr, nil, nil){
        var info *convWasm_ClassInfo
        info = self.FP.getSerrInfo(_env)
        pre.Insert(WasmBin_structNew(_env, info.FP.Get_structType(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), WasmBin_globalGet(_env, info.FP.Get_vtGlobal(_env)), self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString)))))
        return WasmBin_seq(_env, pre), funcType.FP.Get_retTypeInfoList(_env)
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns_collectgarbage, nil, nil){
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns_tonumber, nil, nil){
        return self.FP.lowerLuaCall(_env, node, WasmRt_luaTonumber, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns_require, nil, nil){
        return self.FP.lowerLuaCall(_env, node, WasmRt_luaRequire, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns__load, nil, nil){
        return self.FP.lowerLuaCall(_env, node, WasmRt_luaLoad, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns_loadfile, nil, nil){
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnLuaLoadfile(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString)))))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeDDD))
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns_expandLuavalMap, nil, nil){
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnLuaExpand(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeStem_)))))
        return WasmBin_seq(_env, pre), funcType.FP.Get_retTypeInfoList(_env)
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.Lns_type, nil, nil){
        pre.Insert(WasmBin_call(_env, self.FP.fnLuaType(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeStem_)))))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString))
    }
    if symbol.FP.Get_kind(_env) == Ast_SymbolKind__Fun{
        {
            _funcIndex := self.func2index.Get(funcType)
            if !Lns_IsNil( _funcIndex ) {
                funcIndex := _funcIndex.(LnsInt)
                var cloArg *WasmBin_Ins
                if Lns_isCondTrue( self.FP.findLoc(_env, symbol)){
                    cloArg = self.FP.readSym(_env, &refNode.Nodes_Node, symbol)
                } else { 
                    cloArg = WasmBin_refNull(_env, self.rt.FP.Get_tClo(_env))
                }
                var args *LnsList2_[*WasmBin_Ins]
                args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cloArg))
                convWasm_append_17_(_env, args, self.FP.lowerArgs(_env, node.FP.Get_argList(_env), funcType.FP.Get_argTypeInfoList(_env), pre))
                pre.Insert(WasmBin_call(_env, funcIndex, args))
                return WasmBin_seq(_env, pre), funcType.FP.Get_retTypeInfoList(_env)
            }
        }
    }
    var cloIns *WasmBin_Ins
    {
        _val := self.node2val.Get(&refNode.Nodes_Node)
        if !Lns_IsNil( _val ) {
            val := _val.(*convWasm_Val)
            cloIns = val.FP.Get_ins(_env)
        } else {
            cloIns = self.FP.readSym(_env, &refNode.Nodes_Node, symbol)
        }
    }
    pre.Insert(self.FP.callClosure(_env, cloIns, funcType, node.FP.Get_argList(_env), pre))
    return WasmBin_seq(_env, pre), self.FP.normType(_env, funcType).FP.Get_retTypeInfoList(_env)
}
// 3001: decl @lune.@base.@convWasm.convFilter.lowerArgAny
func (self *convWasm_convFilter) lowerArgAny(_env *LnsEnv, node *Nodes_ExpCallNode,index LnsInt) *WasmBin_Ins {
    var argList *Nodes_ExpListNode
    argList = Lns_unwrap( node.FP.Get_argList(_env)).(*Nodes_ExpListNode)
    return self.FP.toAny(_env, self.FP.lowerVal(_env, argList.FP.Get_expList(_env).GetAt(index)))
}
// 3005: decl @lune.@base.@convWasm.convFilter.lowerArgAs
func (self *convWasm_convFilter) lowerArgAs(_env *LnsEnv, node *Nodes_ExpCallNode,index LnsInt,typeInfo *Ast_TypeInfo) *WasmBin_Ins {
    var argList *Nodes_ExpListNode
    
    {
        _argList := node.FP.Get_argList(_env)
        if _argList == nil{
            return self.FP.getDefaultVal(_env, self.FP.getRep(_env, typeInfo))
        } else {
            argList = _argList.(*Nodes_ExpListNode)
        }
    }
    if index > argList.FP.Get_expList(_env).Len(){
        return self.FP.getDefaultVal(_env, self.FP.getRep(_env, typeInfo))
    }
    return self.FP.lowerExpAs(_env, argList.FP.Get_expList(_env).GetAt(index), typeInfo)
}
// 3018: decl @lune.@base.@convWasm.convFilter.lowerListMethod
func (self *convWasm_convFilter) lowerListMethod(_env *LnsEnv, node *Nodes_ExpCallNode,prefix *Nodes_Node,name string)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var list *WasmBin_Ins
    list = self.FP.lowerExp(_env, prefix)
    if _switch0 := name; _switch0 == "insert" {
        return WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](list, self.FP.lowerArgAny(_env, node, 1)))), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    } else if _switch0 == "remove" {
        return WasmBin_call(_env, self.rt.FP.FnListRemove(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](list, WasmBin_call(_env, self.FP.fnIntOrZero(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeStem_))))))), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeStem_))
    } else if _switch0 == "sort" {
        var lessInfo *WasmRt_CloInfo
        lessInfo = self.rt.FP.GetLessCloInfo(_env)
        var cmp *WasmBin_Ins
        cmp = WasmBin_refNull(_env, lessInfo.FP.Get_cloType(_env))
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                if argList.FP.Get_expList(_env).Len() > 0{
                    var arg *Nodes_Node
                    arg = argList.FP.Get_expList(_env).GetAt(1)
                    if arg.FP.Get_kind(_env) != Nodes_NodeKind_get_Abbr(_env){
                        var val *convWasm_Val
                        val = self.FP.lowerVal(_env, arg)
                        var lessType *WasmRt_CloInfo
                        lessType = self.rt.FP.GetCloInfo(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef, WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI32)))
                        if val.FP.Get_rep(_env).FP.EqualsType(_env, WasmBin_refType(_env, lessType.FP.Get_cloType(_env))){
                            cmp = val.FP.Get_ins(_env)
                        } else { 
                            cmp = self.FP.adaptClosure(_env, val.FP.Get_ins(_env), val.FP.Get_typeInfo(_env), self.FP.getLessFormType(_env, node))
                        }
                    }
                }
            }
        }
        return WasmBin_call(_env, self.rt.FP.FnListSort(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](list, cmp))), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    } else if _switch0 == "unpack" {
        return WasmBin_call(_env, self.rt.FP.FnListClone(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](list))), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeDDD))
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support list method -- %s", Lns_2DDD(name)))
// insert a dummy
    return nil,nil
}
// 3060: decl @lune.@base.@convWasm.convFilter.getLessFormType
func (self *convWasm_convFilter) getLessFormType(_env *LnsEnv, node *Nodes_ExpCallNode) *Ast_TypeInfo {
    var funcType *Ast_TypeInfo
    funcType = self.FP.normType(_env, node.FP.Get_func(_env).FP.Get_expType(_env))
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(node.FP.Get_func(_env).FP)
        if !Lns_IsNil( _fieldNode ) {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            {
                _symbol := fieldNode.FP.Get_symbolInfo(_env)
                if !Lns_IsNil( _symbol ) {
                    symbol := _symbol.(*Ast_SymbolInfo)
                    var declType *Ast_TypeInfo
                    declType = symbol.FP.Get_typeInfo(_env)
                    if declType.FP.Get_argTypeInfoList(_env).Len() > 0{
                        return declType.FP.Get_argTypeInfoList(_env).GetAt(1)
                    }
                }
            }
        }
    }
    return funcType.FP.Get_argTypeInfoList(_env).GetAt(1)
}
// 3074: decl @lune.@base.@convWasm.convFilter.fnIntOrZero
func (self *convWasm_convFilter) fnIntOrZero(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("intOrZero")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64)))
    self.adapterMap.Set("intOrZero",index)
    self.mod.FP.SetFuncName(_env, index, "intOrZero")
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i64Const(_env, 0)))))), nil, nil), self.rt.FP.UnboxI64(_env, convWasm_lg_2_(_env, 0)))))
    return index
}
// 3088: decl @lune.@base.@convWasm.convFilter.lowerSetMethod
func (self *convWasm_convFilter) lowerSetMethod(_env *LnsEnv, node *Nodes_ExpCallNode,prefix *Nodes_Node,name string)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var set *WasmBin_Ins
    set = self.FP.lowerExp(_env, prefix)
    if _switch0 := name; _switch0 == "add" {
        var work LnsInt
        work = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
        return WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, self.FP.lowerArgAny(_env, node, 1)), WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](set, convWasm_lg_2_(_env, work), convWasm_lg_2_(_env, work))))))), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    } else if _switch0 == "del" {
        return WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](set, self.FP.lowerArgAny(_env, node, 1), WasmBin_refNull(_env, WasmBin_htAny)))), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    } else if _switch0 == "has" {
        return convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refIsNull(_env, WasmBin_call(_env, self.rt.FP.FnMapGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](set, self.FP.lowerArgAny(_env, node, 1)))))))), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeBool))
    } else if _switch0 == "and" || _switch0 == "or" || _switch0 == "sub" {
        var mode LnsInt
        mode = 2
        if name == "and"{
            mode = 0
        } else if name == "or"{
            mode = 1
        }
        return WasmBin_call(_env, self.rt.FP.FnSetOp(_env, mode), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](set, self.FP.lowerArgAs(_env, node, 1, prefix.FP.Get_expType(_env))))), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](prefix.FP.Get_expType(_env)))
    } else if _switch0 == "clone" {
        return WasmBin_call(_env, self.rt.FP.FnMapClone(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](set))), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](prefix.FP.Get_expType(_env)))
    } else if _switch0 == "len" {
        return convWasm_p_4_(_env, WasmBin_opI64ExtendI32U, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structGet(_env, self.rt.FP.Get_tMap(_env), WasmRt_fMapCount, set)))), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeInt))
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support set method -- %s", Lns_2DDD(name)))
// insert a dummy
    return nil,nil
}
// 3130: decl @lune.@base.@convWasm.convFilter.lowerStrMethod
func (self *convWasm_convFilter) lowerStrMethod(_env *LnsEnv, node *Nodes_ExpCallNode,prefix *Nodes_Node,name string)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var strVal *WasmBin_Ins
    strVal = self.FP.lowerExpAs(_env, prefix, Ast_builtinTypeString)
    if _switch0 := name; _switch0 == "sub" {
        var endIns *WasmBin_Ins
        endIns = WasmBin_i64Const(_env, -1)
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( argList.FP.Get_expList(_env).Len() >= 2) &&
                    _env.SetStackVal( argList.FP.Get_expList(_env).GetAt(2).FP.Get_kind(_env) != Nodes_NodeKind_get_Abbr(_env)) ).(bool)){
                    endIns = WasmBin_call(_env, self.FP.fnIntOrMinus1(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 2, Ast_builtinTypeStem_))))
                }
            }
        }
        return WasmBin_call(_env, self.rt.FP.FnStrSub(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](strVal, self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeInt), endIns))), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString))
    } else if _switch0 == "find" {
        var initIns *WasmBin_Ins
        initIns = WasmBin_i64Const(_env, 1)
        var plainIns *WasmBin_Ins
        plainIns = WasmBin_i32Const(_env, 0)
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( argList.FP.Get_expList(_env).Len() >= 2) &&
                    _env.SetStackVal( argList.FP.Get_expList(_env).GetAt(2).FP.Get_kind(_env) != Nodes_NodeKind_get_Abbr(_env)) ).(bool)){
                    initIns = WasmBin_call(_env, self.FP.fnIntOrDefault(_env, 1), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 2, Ast_builtinTypeStem_))))
                }
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( argList.FP.Get_expList(_env).Len() >= 3) &&
                    _env.SetStackVal( argList.FP.Get_expList(_env).GetAt(3).FP.Get_kind(_env) != Nodes_NodeKind_get_Abbr(_env)) ).(bool)){
                    plainIns = self.FP.toCond(_env, self.FP.lowerVal(_env, argList.FP.Get_expList(_env).GetAt(3)))
                }
            }
        }
        return WasmBin_call(_env, self.rt.FP.FnStrFind(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](strVal, self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString), initIns, plainIns))), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeDDD))
    } else if _switch0 == "gsub" {
        return WasmBin_call(_env, self.rt.FP.FnStrGsub(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](strVal, self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString), self.FP.lowerArgAs(_env, node, 2, Ast_builtinTypeString)))), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString, Ast_builtinTypeInt))
    } else if _switch0 == "format" {
        var pre *LnsList2_[*WasmBin_Ins]
        pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        var work LnsInt
        work = self.FP.tempLocal(_env, self.rt.FP.Get_strType(_env))
        pre.Insert(convWasm_ls_3_(_env, work, strVal))
        var vals *LnsList2_[*convWasm_Val]
        vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                if argList.FP.Get_expList(_env).Len() == 1{
                    {
                        _dddNode := Nodes_ExpToDDDNodeDownCastF(argList.FP.Get_expList(_env).GetAt(1).FP)
                        if !Lns_IsNil( _dddNode ) {
                            dddNode := _dddNode.(*Nodes_ExpToDDDNode)
                            vals = self.FP.lowerValList(_env, dddNode.FP.Get_expList(_env), pre)
                        } else {
                            vals = self.FP.lowerValList(_env, argList, pre)
                        }
                    }
                }
            }
        }
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnFormat(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, work), self.FP.buildList(_env, vals)))))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString))
    }
    var luaIdMap *LnsMap2_[string,LnsInt]
    luaIdMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{"byte":WasmRt_luaByte,"rep":WasmRt_luaRep,"lower":WasmRt_luaLower,"upper":WasmRt_luaUpper,"reverse":WasmRt_luaReverse,"len":WasmRt_luaLen,"replace":WasmRt_luaReplace,})
    {
        _luaId := luaIdMap.Get(name)
        if !Lns_IsNil( _luaId ) {
            luaId := _luaId.(LnsInt)
            return self.FP.lowerLuaCall(_env, node, luaId, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](strVal)))
        }
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support str method -- %s", Lns_2DDD(name)))
// insert a dummy
    return nil,nil
}
// 3204: decl @lune.@base.@convWasm.convFilter.lowerLuaCall
func (self *convWasm_convFilter) lowerLuaCall(_env *LnsEnv, node *Nodes_ExpCallNode,luaId LnsInt,headArgs *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    return self.FP.lowerLuaCallSub(_env, node, luaId, nil, headArgs, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}
// 3218: decl @lune.@base.@convWasm.convFilter.lowerLuaCallSub
func (self *convWasm_convFilter) lowerLuaCallSub(_env *LnsEnv, node *Nodes_ExpCallNode,luaId LnsInt,funcVal LnsAny,headArgs *LnsList2_[*WasmBin_Ins],headPre *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    convWasm_append_17_(_env, pre, headPre)
    var args LnsInt
    args = self.FP.tempLocal(_env, self.rt.FP.Get_listType(_env))
    pre.Insert(convWasm_ls_3_(_env, args, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 4))))))
    for _, _ins := range( headArgs.Items ) {
        ins := _ins
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, args), ins))))
    }
    {
        _argList := node.FP.Get_argList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            for _, _exp := range( argList.FP.Get_expList(_env).Items ) {
                exp := _exp
                {
                    _dddNode := Nodes_ExpToDDDNodeDownCastF(exp.FP)
                    if !Lns_IsNil( _dddNode ) {
                        dddNode := _dddNode.(*Nodes_ExpToDDDNode)
                        for _, _val := range( self.FP.lowerValList(_env, dddNode.FP.Get_expList(_env), pre).Items ) {
                            val := _val
                            pre.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, args), self.FP.toAny(_env, val)))))
                        }
                        continue
                    }
                }
                var argIns *WasmBin_Ins
                if exp.FP.Get_kind(_env) == Nodes_NodeKind_get_Abbr(_env){
                    argIns = WasmBin_refNull(_env, WasmBin_htAny)
                } else { 
                    argIns = self.FP.lowerExpAs(_env, exp, Ast_builtinTypeStem_)
                }
                pre.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, args), argIns))))
            }
        }
    }
    var result *WasmBin_Ins
    if funcVal != nil{
        funcVal_147 := funcVal.(*WasmBin_Ins)
        result = WasmBin_call(_env, self.rt.FP.FnLuaCallVal(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](funcVal_147, convWasm_lg_2_(_env, args))))
    } else {
        {
            _nativeFunc := self.rt.FP.GetNativeStrFunc(_env, luaId)
            if !Lns_IsNil( _nativeFunc ) {
                nativeFunc := _nativeFunc.(LnsInt)
                result = WasmBin_call(_env, nativeFunc, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, args))))
            } else {
                result = WasmBin_call(_env, self.rt.FP.FnLuaCall(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, luaId), convWasm_lg_2_(_env, args))))
            }
        }
    }
    var retTypeList *LnsList2_[*Ast_TypeInfo]
    if Lns_isCondTrue( funcVal){
        retTypeList = node.FP.Get_expTypeList(_env)
    } else { 
        retTypeList = node.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_retTypeInfoList(_env)
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( retTypeList.Len() == 1) &&
        _env.SetStackVal( retTypeList.GetAt(1).FP.Get_kind(_env) != Ast_TypeInfoKind__DDD) ).(bool)){
        pre.Insert(self.FP.coerce(_env, self.rt.FP.ListItem(_env, result, 0), Ast_builtinTypeStem_, retTypeList.GetAt(1)))
        return WasmBin_seq(_env, pre), retTypeList
    }
    pre.Insert(result)
    return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeDDD))
}
// 3275: decl @lune.@base.@convWasm.convFilter.fnLuaType
func (self *convWasm_convFilter) fnLuaType(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("luaType")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_strType(_env))))
    self.adapterMap.Set("luaType",index)
    self.mod.FP.SetFuncName(_env, index, "luaType")
    var convWasm_check func(_env *LnsEnv, cond *WasmBin_Ins,name string) *WasmBin_Ins
    convWasm_check = func(_env *LnsEnv, cond *WasmBin_Ins,name string) *WasmBin_Ins {
        return WasmBin_ifElse(_env, cond, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.StrLiteral(_env, name)))))), nil, nil)
    }
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_check(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, 0)), "nil"), convWasm_check(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxI64(_env), convWasm_lg_2_(_env, 0)), "number"), convWasm_check(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxF64(_env), convWasm_lg_2_(_env, 0)), "number"), convWasm_check(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxBool(_env), convWasm_lg_2_(_env, 0)), "boolean"), convWasm_check(_env, WasmBin_refTest(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, 0)), "string"), convWasm_check(_env, WasmBin_refTest(_env, self.rt.FP.Get_tClo(_env), convWasm_lg_2_(_env, 0)), "function"), WasmBin_ifElse(_env, WasmBin_refTest(_env, self.rt.FP.Get_tLuaVal(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_check(_env, convWasm_p_4_(_env, WasmBin_opI32Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.LuaValType(_env, convWasm_lg_2_(_env, 0)), WasmBin_i32Const(_env, WasmRt_luaTFunction)))), "function"), convWasm_check(_env, convWasm_p_4_(_env, WasmBin_opI32Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.LuaValType(_env, convWasm_lg_2_(_env, 0)), WasmBin_i32Const(_env, 7)))), "userdata"), convWasm_check(_env, convWasm_p_4_(_env, WasmBin_opI32Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.LuaValType(_env, convWasm_lg_2_(_env, 0)), WasmBin_i32Const(_env, 8)))), "thread"))), nil, nil), self.rt.FP.StrLiteral(_env, "table"))))
    return index
}
// 3303: decl @lune.@base.@convWasm.convFilter.fnStemIndex
func (self *convWasm_convFilter) fnStemIndex(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("stemIndex")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef, WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)))
    self.adapterMap.Set("stemIndex",index)
    self.mod.FP.SetFuncName(_env, index, "stemIndex")
    var pos LnsInt
    pos = _func.FP.AddLocal(_env, WasmBin_typeI64)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refTest(_env, self.rt.FP.Get_tLuaVal(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnLuaGetAtLazy(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0), convWasm_lg_2_(_env, 1))))))))), nil, nil), WasmBin_ifElse(_env, WasmBin_refTest(_env, self.rt.FP.Get_tList(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxI64(_env), convWasm_lg_2_(_env, 1)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refCast(_env, self.rt.FP.Get_tList(_env), convWasm_lg_2_(_env, 0)), self.rt.FP.UnboxI64(_env, convWasm_lg_2_(_env, 1)))))))))), nil, nil), WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, WasmBin_htAny)))))), nil, nil), WasmBin_ifElse(_env, WasmBin_refTest(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnMapGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refCast(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0)), convWasm_lg_2_(_env, 1))))))))), nil, nil), WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32And, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refTest(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, 0)), WasmBin_refTest(_env, self.rt.FP.Get_tBoxI64(_env), convWasm_lg_2_(_env, 1))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.UnboxI64(_env, convWasm_lg_2_(_env, 1)), WasmBin_i64Const(_env, 1))))), WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32Or, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i64Const(_env, 0)))), convWasm_p_4_(_env, WasmBin_opI64GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), convWasm_p_4_(_env, WasmBin_opI64ExtendI32U, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_arrayLen(_env, WasmBin_refCast(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, 0)))))))))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, WasmBin_htAny)))))), nil, nil), WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.BoxI64(_env, convWasm_p_4_(_env, WasmBin_opI64ExtendI32U, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_arrayGetU(_env, self.rt.FP.Get_tStr(_env), WasmBin_refCast(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, 0)), convWasm_p_4_(_env, WasmBin_opI32WrapI64, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos)))))))))))))), nil, nil), self.rt.FP.Error(_env, self.rt.FP.StrLiteral(_env, "attempt to index a non-table value")))))
    return index
}
// 3343: decl @lune.@base.@convWasm.convFilter.fnIntOrDefault
func (self *convWasm_convFilter) fnIntOrDefault(_env *LnsEnv, defVal LnsInt) LnsInt {
    var name string
    name = _env.GetVM().String_format("intOr:%d", Lns_2DDD(defVal))
    {
        _index := self.adapterMap.Get(name)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64)))
    self.adapterMap.Set(name,index)
    self.mod.FP.SetFuncName(_env, index, name)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i64Const(_env, defVal)))))), nil, nil), self.rt.FP.UnboxI64(_env, convWasm_lg_2_(_env, 0)))))
    return index
}
// 3358: decl @lune.@base.@convWasm.convFilter.fnIntOrMinus1
func (self *convWasm_convFilter) fnIntOrMinus1(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("intOrMinus1")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64)))
    self.adapterMap.Set("intOrMinus1",index)
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i64Const(_env, -1)))))), nil, nil), self.rt.FP.UnboxI64(_env, convWasm_lg_2_(_env, 0)))))
    return index
}
// 3371: decl @lune.@base.@convWasm.convFilter.lowerStreamMethod
func (self *convWasm_convFilter) lowerStreamMethod(_env *LnsEnv, node *Nodes_ExpCallNode,prefix *Nodes_Node,name string)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var obj LnsInt
    obj = self.FP.tempLocal(_env, self.rt.FP.Get_objType(_env))
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, obj, WasmBin_refCast(_env, self.rt.FP.Get_tObj(_env), self.FP.lowerExpAs(_env, prefix, Ast_builtinTypeStem_)))))
    var isFile *WasmBin_Ins
    isFile = self.rt.FP.IsFile(_env, convWasm_lg_2_(_env, obj))
    var ifType *Ast_TypeInfo
    ifType = self.FP.getGenSrc(_env, prefix.FP.Get_expType(_env))
    if self.FP.isLuaval(_env, prefix.FP.Get_expType(_env)){
        ifType = self.FP.getGenSrc(_env, prefix.FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_extedType(_env))
    }
    var isUser *WasmBin_Ins
    isUser = convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refIsNull(_env, WasmBin_structGet(_env, self.rt.FP.Get_tVT(_env), WasmRt_fVtItabs, WasmBin_structGet(_env, self.rt.FP.Get_tObj(_env), WasmRt_fObjVt, convWasm_lg_2_(_env, obj)))))))
    var convWasm_userCall func(_env *LnsEnv, args *LnsList2_[*WasmBin_Ins],retLocals *LnsList2_[LnsInt],retReps *LnsList2_[*WasmBin_ValType]) *WasmBin_Ins
    convWasm_userCall = func(_env *LnsEnv, args *LnsList2_[*WasmBin_Ins],retLocals *LnsList2_[LnsInt],retReps *LnsList2_[*WasmBin_ValType]) *WasmBin_Ins {
        var ifInfo *convWasm_IfInfo
        ifInfo = self.FP.getIfInfo(_env, ifType)
        var methodIndex LnsInt
        
        {
            _methodIndex := ifInfo.FP.Get_name2index(_env).Get(name)
            if _methodIndex == nil{
                self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not found method -- %s", Lns_2DDD(name)))
            } else {
                methodIndex = _methodIndex.(LnsInt)
            }
        }
        var methodType *Ast_TypeInfo
        methodType = Lns_unwrap( self.FP.getIfMethodType(_env, ifType, name)).(*Ast_TypeInfo)
        var callArgs *LnsList2_[*WasmBin_Ins]
        callArgs = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj)))
        for _, _arg := range( args.Items ) {
            arg := _arg
            callArgs.Insert(arg)
        }
        var itab *WasmBin_Ins
        itab = WasmBin_refCast(_env, ifInfo.FP.Get_itType(_env), WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tVT(_env), WasmRt_fVtItabs, WasmBin_structGet(_env, self.rt.FP.Get_tObj(_env), WasmRt_fObjVt, convWasm_lg_2_(_env, obj))), WasmBin_i32Const(_env, ifInfo.FP.Get_ifId(_env))))
        callArgs.Insert(WasmBin_structGet(_env, ifInfo.FP.Get_itType(_env), methodIndex, itab))
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_callRef(_env, ifInfo.FP.Get_methodFuncTypes(_env).GetAt(methodIndex + 1), callArgs)))
        var calleeRets *LnsList2_[*WasmBin_ValType]
        calleeRets = self.FP.getRepList(_env, methodType.FP.Get_retTypeInfoList(_env))
        var works *LnsList2_[LnsInt]
        works = NewLnsList2_[LnsInt]([]LnsInt{})
        for _, _rep := range( calleeRets.Items ) {
            rep := _rep
            works.Insert(LnsInt(self.FP.tempLocal(_env, rep)))
        }
        {
            var _forFrom0 LnsInt = works.Len()
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
                body.Insert(convWasm_ls_3_(_env, works.GetAt(index), WasmBin_nop(_env)))
                _forWork0 += _forDelta0
            }
        }
        for _index, _retLocal := range( retLocals.Items ) {
            index := _index + 1
            retLocal := _retLocal
            if index <= works.Len(){
                body.Insert(convWasm_ls_3_(_env, retLocal, self.FP.coerceRep(_env, convWasm_lg_2_(_env, works.GetAt(index)), calleeRets.GetAt(index), retReps.GetAt(index))))
            } else { 
                body.Insert(convWasm_ls_3_(_env, retLocal, self.FP.getDefaultVal(_env, retReps.GetAt(index))))
            }
        }
        return WasmBin_seq(_env, body)
    }
    if _switch0 := name; _switch0 == "write" {
        var txt LnsInt
        txt = self.FP.tempLocal(_env, self.rt.FP.Get_strType(_env))
        pre.Insert(convWasm_ls_3_(_env, txt, self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString)))
        self.FP.getStdStream(_env, true)
        var ret1 LnsInt
        ret1 = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
        var ret2 LnsInt
        ret2 = self.FP.tempLocal(_env, self.rt.FP.Get_strType(_env))
        pre.Insert(convWasm_ls_3_(_env, ret1, convWasm_lg_2_(_env, obj)))
        pre.Insert(convWasm_ls_3_(_env, ret2, WasmBin_refNull(_env, self.rt.FP.Get_tStr(_env))))
        pre.Insert(WasmBin_ifElse(_env, isFile, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnFileWrite(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj), convWasm_lg_2_(_env, txt)))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refEq(_env, convWasm_lg_2_(_env, obj), WasmBin_globalGet(_env, self.stderrGlobal)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnHostWriteErr(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, txt)))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, isUser, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_userCall(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, txt))), NewLnsList2_[LnsInt](Lns_2DDDGen[LnsInt](ret1, ret2)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef, self.rt.FP.Get_strType(_env)))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnHostWrite(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, txt)))))), nil))), nil))), nil))
        pre.Insert(convWasm_lg_2_(_env, ret1))
        pre.Insert(convWasm_lg_2_(_env, ret2))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeStem_, Ast_builtinTypeString.FP.Get_nilableTypeInfo(_env)))
    } else if _switch0 == "read" {
        var fmt LnsInt
        fmt = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
        pre.Insert(convWasm_ls_3_(_env, fmt, self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeStem_)))
        var ret LnsInt
        ret = self.FP.tempLocal(_env, self.rt.FP.Get_strType(_env))
        pre.Insert(convWasm_ls_3_(_env, ret, WasmBin_refNull(_env, self.rt.FP.Get_tStr(_env))))
        pre.Insert(WasmBin_ifElse(_env, isFile, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, ret, WasmBin_call(_env, self.rt.FP.FnFileRead(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj), convWasm_lg_2_(_env, fmt))))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, isUser, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_userCall(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, fmt))), NewLnsList2_[LnsInt](Lns_2DDDGen[LnsInt](ret)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_strType(_env)))))), nil, nil))), nil))
        pre.Insert(convWasm_lg_2_(_env, ret))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString.FP.Get_nilableTypeInfo(_env)))
    } else if _switch0 == "seek" {
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnFileSeek(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj), self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString), self.FP.lowerArgAs(_env, node, 2, Ast_builtinTypeInt)))))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeDDD))
    } else if _switch0 == "close" || _switch0 == "flush" {
        var fileOp *LnsList2_[*WasmBin_Ins]
        fileOp = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        if name == "close"{
            fileOp.Insert(WasmBin_call(_env, self.rt.FP.FnFileClose(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj)))))
        }
        pre.Insert(WasmBin_ifElse(_env, isFile, fileOp, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, isUser, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_userCall(_env, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), NewLnsList2_[LnsInt]([]LnsInt{}), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{})))), nil, nil))), nil))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support stream method -- %s", Lns_2DDD(name)))
// insert a dummy
    return nil,nil
}
// 3488: decl @lune.@base.@convWasm.convFilter.isStreamType
func (self *convWasm_convFilter) isStreamType(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    var work *Ast_TypeInfo
    work = self.FP.getGenSrc(_env, typeInfo)
    var builtin *Builtin_BuiltinFuncType
    builtin = self.builtinFuncs
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( work == self.FP.getGenSrc(_env, builtin.Ostream_)) ||
        _env.SetStackVal( work == self.FP.getGenSrc(_env, builtin.Istream_)) ||
        _env.SetStackVal( work == self.FP.getGenSrc(_env, builtin.Luastream_)) ).(bool)
}
// 3497: decl @lune.@base.@convWasm.convFilter.lowerIfMethodCall
func (self *convWasm_convFilter) lowerIfMethodCall(_env *LnsEnv, prefix *Nodes_Node,name string,pre *LnsList2_[*WasmBin_Ins],argList LnsAny)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var objType *Ast_TypeInfo
    objType = self.FP.getObjTypeForMethod(_env, prefix.FP.Get_expType(_env), name)
    var ifInfo *convWasm_IfInfo
    ifInfo = self.FP.getIfInfo(_env, objType)
    var methodIndex LnsInt
    
    {
        _methodIndex := ifInfo.FP.Get_name2index(_env).Get(name)
        if _methodIndex == nil{
            self.FP.error(_env, prefix, _env.GetVM().String_format("not found method -- %s", Lns_2DDD(name)))
        } else {
            methodIndex = _methodIndex.(LnsInt)
        }
    }
    var methodType *Ast_TypeInfo
    methodType = Lns_unwrap( self.FP.getIfMethodType(_env, objType, name)).(*Ast_TypeInfo)
    var obj LnsInt
    obj = self.FP.tempLocal(_env, self.rt.FP.Get_objType(_env))
    pre.Insert(convWasm_ls_3_(_env, obj, self.FP.lowerExpAs(_env, prefix, objType)))
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj)))
    convWasm_append_17_(_env, args, self.FP.lowerArgs(_env, argList, methodType.FP.Get_argTypeInfoList(_env), pre))
    var itab *WasmBin_Ins
    itab = WasmBin_refCast(_env, ifInfo.FP.Get_itType(_env), WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tVT(_env), WasmRt_fVtItabs, WasmBin_structGet(_env, self.rt.FP.Get_tObj(_env), WasmRt_fObjVt, convWasm_lg_2_(_env, obj))), WasmBin_i32Const(_env, ifInfo.FP.Get_ifId(_env))))
    args.Insert(WasmBin_structGet(_env, ifInfo.FP.Get_itType(_env), methodIndex, itab))
    pre.Insert(WasmBin_callRef(_env, ifInfo.FP.Get_methodFuncTypes(_env).GetAt(methodIndex + 1), args))
    return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
}
// 3523: decl @lune.@base.@convWasm.convFilter.lowerMethodCall
func (self *convWasm_convFilter) lowerMethodCall(_env *LnsEnv, node *Nodes_ExpCallNode,prefix *Nodes_Node,name string,pre *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var objType *Ast_TypeInfo
    objType = self.FP.getObjType(_env, prefix.FP.Get_expType(_env))
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( objType.FP.Get_kind(_env) != Ast_TypeInfoKind__IF) &&
        _env.SetStackVal( Lns_op_not(self.FP.getClassInfo(_env, objType))) &&
        _env.SetStackVal( Ast_isBuiltin(_env, self.FP.getGenSrc(_env, objType).FP.Get_typeId(_env).Id)) ).(bool)){
        pre.Insert(WasmBin_drop(_env, self.FP.toAny(_env, self.FP.lowerVal(_env, prefix))))
        var retTypes *LnsList2_[*Ast_TypeInfo]
        retTypes = node.FP.Get_func(_env).FP.Get_expType(_env).FP.Get_retTypeInfoList(_env)
        for _, _retType := range( retTypes.Items ) {
            retType := _retType
            pre.Insert(self.FP.getDefaultVal(_env, self.FP.getRep(_env, retType)))
        }
        return WasmBin_seq(_env, pre), retTypes
    }
    return self.FP.lowerObjMethodCall(_env, &node.Nodes_Node, prefix, name, pre, node.FP.Get_argList(_env))
}
// 3543: decl @lune.@base.@convWasm.convFilter.lowerObjMethodCall
func (self *convWasm_convFilter) lowerObjMethodCall(_env *LnsEnv, errNode *Nodes_Node,prefix *Nodes_Node,name string,pre *LnsList2_[*WasmBin_Ins],argList LnsAny)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var objType *Ast_TypeInfo
    objType = self.FP.getObjTypeForMethod(_env, prefix.FP.Get_expType(_env), name)
    if objType.FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
        return self.FP.lowerIfMethodCall(_env, prefix, name, pre, argList)
    }
    var info *convWasm_ClassInfo
    
    {
        _info := self.FP.getClassInfo(_env, objType)
        if _info == nil{
            self.FP.error(_env, errNode, _env.GetVM().String_format("not found class -- %s", Lns_2DDD(objType.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            info = _info.(*convWasm_ClassInfo)
        }
    }
    var slotIndex LnsInt
    
    {
        _slotIndex := info.FP.Get_name2slot(_env).Get(name)
        if _slotIndex == nil{
            self.FP.error(_env, errNode, _env.GetVM().String_format("not found method -- %s", Lns_2DDD(name)))
        } else {
            slotIndex = _slotIndex.(LnsInt)
        }
    }
    var slot *convWasm_SlotInfo
    slot = info.FP.Get_slotList(_env).GetAt(slotIndex)
    var obj LnsInt
    obj = self.FP.tempLocal(_env, WasmBin_refType(_env, info.FP.Get_structType(_env)))
    pre.Insert(convWasm_ls_3_(_env, obj, self.FP.lowerExpAs(_env, prefix, info.FP.Get_typeInfo(_env))))
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj)))
    convWasm_append_17_(_env, args, self.FP.lowerArgs(_env, argList, slot.FP.Get_methodType(_env).FP.Get_argTypeInfoList(_env), pre))
    args.Insert(WasmBin_structGet(_env, info.FP.Get_vtType(_env), slotIndex, WasmBin_structGet(_env, info.FP.Get_structType(_env), WasmRt_fObjVt, convWasm_lg_2_(_env, obj))))
    pre.Insert(WasmBin_callRef(_env, slot.FP.Get_funcType(_env), args))
    return WasmBin_seq(_env, pre), slot.FP.Get_methodType(_env).FP.Get_retTypeInfoList(_env)
}
// 3569: decl @lune.@base.@convWasm.convFilter.getIfMethodType
func (self *convWasm_convFilter) getIfMethodType(_env *LnsEnv, ifType *Ast_TypeInfo,name string) LnsAny {
    var _map *LnsMap2_[string,*Ast_TypeInfo]
    _map = NewLnsMap2_[string,*Ast_TypeInfo]( map[string]*Ast_TypeInfo{})
    self.FP.collectIfMethods(_env, self.FP.getGenSrc(_env, ifType), _map)
    return _map.Get(name)
}
// 3576: decl @lune.@base.@convWasm.convFilter.lowerCallField
func (self *convWasm_convFilter) lowerCallField(_env *LnsEnv, node *Nodes_ExpCallNode,fieldNode *Nodes_RefFieldNode,pre *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var prefix *Nodes_Node
    prefix = fieldNode.FP.Get_prefix(_env)
    var name string
    name = fieldNode.FP.Get_field(_env).Txt
    var prefixType *Ast_TypeInfo
    prefixType = self.FP.normType(_env, prefix.FP.Get_expType(_env))
    if self.FP.isLuaval(_env, prefix.FP.Get_expType(_env)){
        prefixType = self.FP.normType(_env, prefix.FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_extedType(_env))
    }
    var builtin *Builtin_BuiltinFuncType
    builtin = self.builtinFuncs
    var funcType *Ast_TypeInfo
    funcType = node.FP.Get_func(_env).FP.Get_expType(_env)
    {
        _initIns := self.FP.getLazyInit(_env, prefix)
        if !Lns_IsNil( _initIns ) {
            initIns := _initIns.(*WasmBin_Ins)
            pre.Insert(initIns)
        }
    }
    if funcType.FP.Equals(_env, self.processInfo, builtin.String_format, nil, nil){
        var vals *LnsList2_[*convWasm_Val]
        vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                vals = self.FP.lowerValList(_env, argList, pre)
            }
        }
        var rest *LnsList2_[*convWasm_Val]
        rest = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
        {
            var _forFrom0 LnsInt = 2
            var _forTo0 LnsInt = vals.Len()
            for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                index := _forWork0
                rest.Insert(vals.GetAt(index))
            }
        }
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnFormat(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.coerceVal(_env, vals.GetAt(1), Ast_builtinTypeString), self.FP.buildList(_env, rest)))))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString))
    }
    if _switch0 := prefixType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__Box {
        if name == "val"{
            pre.Insert(WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExpAs(_env, prefix, prefix.FP.Get_expType(_env)), WasmBin_i64Const(_env, 1)))))
            return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeStem_))
        }
    } else if _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        return self.FP.lowerListMethod(_env, node, prefix, name)
    } else if _switch0 == Ast_TypeInfoKind__Set {
        return self.FP.lowerSetMethod(_env, node, prefix, name)
    }
    if self.FP.isType(_env, prefixType, Ast_builtinTypeString){
        return self.FP.lowerStrMethod(_env, node, prefix, name)
    }
    if self.FP.isStreamType(_env, prefixType){
        return self.FP.lowerStreamMethod(_env, node, prefix, name)
    }
    {
        _symbol := fieldNode.FP.Get_symbolInfo(_env)
        if !Lns_IsNil( _symbol ) {
            symbol := _symbol.(*Ast_SymbolInfo)
            var methodType *Ast_TypeInfo
            methodType = symbol.FP.GetOrg(_env).FP.Get_typeInfo(_env)
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( symbol.FP.Get_kind(_env) == Ast_SymbolKind__Mbr) ||
                _env.SetStackVal( symbol.FP.Get_kind(_env) == Ast_SymbolKind__Var) ).(bool){
                var cloVal *convWasm_Val
                cloVal = self.FP.lowerVal(_env, &fieldNode.Nodes_Node)
                pre.Insert(self.FP.callClosure(_env, cloVal.FP.Get_ins(_env), methodType, node.FP.Get_argList(_env), pre))
                return WasmBin_seq(_env, pre), self.FP.normType(_env, methodType).FP.Get_retTypeInfoList(_env)
            }
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( methodType.FP.Get_staticFlag(_env)) ||
                _env.SetStackVal( symbol.FP.Get_kind(_env) == Ast_SymbolKind__Fun) ).(bool){
                return self.FP.lowerStaticCall(_env, node, prefix, name, methodType, pre)
            }
        }
    }
    {
        _enumType := Ast_EnumTypeInfoDownCastF(prefixType.FP)
        if !Lns_IsNil( _enumType ) {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            if name == "get__txt"{
                pre.Insert(WasmBin_call(_env, self.FP.fnEnumTxt(_env, enumType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, prefix)))))
                return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString))
            }
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alge) &&
        _env.SetStackVal( name == "get__txt") ).(bool)){
        pre.Insert(self.FP.algeTxt(_env, prefix, prefixType))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString))
    }
    if name == "_toMap"{
        if Lns_op_not(self.toMapDynNode){
            self.toMapDynNode = &node.Nodes_Node
        }
        pre.Insert(WasmBin_call(_env, self.FP.fnToMapDyn(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExpAs(_env, prefix, Ast_builtinTypeStem_)))))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](node.FP.Get_expType(_env)))
    }
    return self.FP.lowerMethodCall(_env, node, prefix, name, pre)
}
// 3664: decl @lune.@base.@convWasm.convFilter.getStringModuleFuncMap
func (self *convWasm_convFilter) getStringModuleFuncMap(_env *LnsEnv) *LnsMap2_[*Ast_TypeInfo,LnsInt] {
    var builtin *Builtin_BuiltinFuncType
    builtin = self.builtinFuncs
    return NewLnsMap2_[*Ast_TypeInfo,LnsInt]( map[*Ast_TypeInfo]LnsInt{builtin.String_find:WasmRt_luaFind,builtin.String_byte:WasmRt_luaByte,builtin.String_rep:WasmRt_luaRep,builtin.String_gsub:WasmRt_luaGsub,builtin.String_lower:WasmRt_luaLower,builtin.String_upper:WasmRt_luaUpper,builtin.String_reverse:WasmRt_luaReverse,builtin.String_sub:WasmRt_luaSub,builtin.Math_random:WasmRt_luaRandom,builtin.Math_randomseed:WasmRt_luaRandomseed,builtin.String_dump:WasmRt_luaDump,})
}
// 3681: decl @lune.@base.@convWasm.convFilter.fnStrJoin
func (self *convWasm_convFilter) fnStrJoin(_env *LnsEnv) LnsInt {
    {
        _index := self.adapterMap.Get("strJoin")
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef, WasmBin_typeAnyRef, WasmBin_typeAnyRef, WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_strType(_env))))
    self.adapterMap.Set("strJoin",index)
    self.mod.FP.SetFuncName(_env, index, "strJoin")
    var list LnsInt
    list = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
    var sep LnsInt
    sep = _func.FP.AddLocal(_env, self.rt.FP.Get_strType(_env))
    var pos LnsInt
    pos = _func.FP.AddLocal(_env, WasmBin_typeI32)
    var last LnsInt
    last = _func.FP.AddLocal(_env, WasmBin_typeI32)
    var result LnsInt
    result = _func.FP.AddLocal(_env, self.rt.FP.Get_strType(_env))
    var first LnsInt
    first = _func.FP.AddLocal(_env, WasmBin_typeI32)
    var total LnsInt
    total = _func.FP.AddLocal(_env, WasmBin_typeI32)
    var work LnsInt
    work = _func.FP.AddLocal(_env, self.rt.FP.Get_strType(_env))
    var convWasm_toI32 func(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins
    convWasm_toI32 = func(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins {
        return convWasm_p_4_(_env, WasmBin_opI32WrapI64, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.UnboxI64(_env, val))))
    }
    var convWasm_item func(_env *LnsEnv) *WasmBin_Ins
    convWasm_item = func(_env *LnsEnv) *WasmBin_Ins {
        return WasmBin_refCast(_env, self.rt.FP.Get_tStr(_env), WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListItems, convWasm_lg_2_(_env, list)), convWasm_p_4_(_env, WasmBin_opI32Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))))
    }
    var convWasm_copyTo func(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins
    convWasm_copyTo = func(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins {
        var strIdx LnsInt
        strIdx = self.rt.FP.Get_tStr(_env)
        return WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, val), WasmBin_arrayCopy(_env, strIdx, strIdx, convWasm_lg_2_(_env, result), convWasm_lg_2_(_env, total), convWasm_lg_2_(_env, work), WasmBin_i32Const(_env, 0), WasmBin_arrayLen(_env, convWasm_lg_2_(_env, work))), convWasm_ls_3_(_env, total, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, total), WasmBin_arrayLen(_env, convWasm_lg_2_(_env, work)))))))))
    }
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, list, WasmBin_refCast(_env, self.rt.FP.Get_tList(_env), convWasm_lg_2_(_env, 0))), convWasm_ls_3_(_env, sep, WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, 1)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.StrLiteral(_env, ""))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refCast(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, 1)))), self.rt.FP.Get_strType(_env))), convWasm_ls_3_(_env, pos, WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, 2)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 1))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_toI32(_env, convWasm_lg_2_(_env, 2)))), WasmBin_typeI32)), convWasm_ls_3_(_env, last, WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, 3)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, convWasm_lg_2_(_env, list)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_toI32(_env, convWasm_lg_2_(_env, 3)))), WasmBin_typeI32)), convWasm_ls_3_(_env, first, convWasm_lg_2_(_env, pos)), WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32GtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), convWasm_lg_2_(_env, last))))), convWasm_ls_3_(_env, total, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, total), WasmBin_arrayLen(_env, convWasm_item(_env)))))), WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), convWasm_lg_2_(_env, last)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, total, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, total), WasmBin_arrayLen(_env, convWasm_lg_2_(_env, sep)))))))), nil, nil), convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))), WasmBin_br(_env, 0))), nil))), nil), convWasm_ls_3_(_env, result, WasmBin_arrayNewDefault(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, total))), convWasm_ls_3_(_env, pos, convWasm_lg_2_(_env, first)), convWasm_ls_3_(_env, total, WasmBin_i32Const(_env, 0)), WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32GtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), convWasm_lg_2_(_env, last))))), convWasm_copyTo(_env, convWasm_item(_env)), WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), convWasm_lg_2_(_env, last)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_copyTo(_env, convWasm_lg_2_(_env, sep)))), nil, nil), convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))), WasmBin_br(_env, 0))), nil))), nil), convWasm_lg_2_(_env, result))))
    return index
}
// 3750: decl @lune.@base.@convWasm.convFilter.lowerStaticCall
func (self *convWasm_convFilter) lowerStaticCall(_env *LnsEnv, node *Nodes_ExpCallNode,prefix *Nodes_Node,name string,methodType *Ast_TypeInfo,pre *LnsList2_[*WasmBin_Ins])(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    if Lns_op_not(self.declaredMethodSet.Has(methodType)){
        if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(name,"^[gs]et_.", nil, nil))){
            var memberName string
            memberName = _env.GetVM().String_sub(name,5, nil)
            {
                _scope := self.FP.getGenSrc(_env, prefix.FP.Get_expType(_env)).FP.Get_scope(_env)
                if !Lns_IsNil( _scope ) {
                    scope := _scope.(*Ast_Scope)
                    {
                        _symbol := scope.FP.Get_symbol2SymbolInfoMap(_env).Get(memberName)
                        if !Lns_IsNil( _symbol ) {
                            symbol := _symbol.(*Ast_SymbolInfo)
                            var orgSym *Ast_SymbolInfo
                            orgSym = symbol.FP.GetOrg(_env)
                            if Lns_isCondTrue( self.globalSym2loc.Get(orgSym)){
                                if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(name,"^get_", nil, nil))){
                                    pre.Insert(self.FP.readSym(_env, &node.Nodes_Node, orgSym))
                                    return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](orgSym.FP.Get_typeInfo(_env)))
                                }
                                pre.Insert(self.FP.writeSym(_env, &node.Nodes_Node, orgSym, self.FP.lowerArgAs(_env, node, 1, orgSym.FP.Get_typeInfo(_env))))
                                return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
                            }
                        }
                    }
                }
            }
        }
    }
    {
        _funcIndex := self.func2index.Get(methodType)
        if !Lns_IsNil( _funcIndex ) {
            funcIndex := _funcIndex.(LnsInt)
            var args *LnsList2_[*WasmBin_Ins]
            args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, self.rt.FP.Get_tClo(_env))))
            convWasm_append_17_(_env, args, self.FP.lowerArgs(_env, node.FP.Get_argList(_env), methodType.FP.Get_argTypeInfoList(_env), pre))
            pre.Insert(WasmBin_call(_env, funcIndex, args))
            return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
        }
    }
    var builtin *Builtin_BuiltinFuncType
    builtin = self.builtinFuncs
    for _strFunc, _luaId := range( self.FP.getStringModuleFuncMap(_env).Items ) {
        strFunc := _strFunc
        luaId := _luaId
        if methodType.FP.Equals(_env, self.processInfo, strFunc, nil, nil){
            return self.FP.lowerLuaCall(_env, node, luaId, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
        }
    }
    for _, _noopFunc := range( NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](builtin.G__lns_runtime_log, builtin.G__lns_runtime_enableLog, builtin.G__lns_runtime_enableDebugLog, builtin.G__lns_runtime_dumpLog, builtin.G__processor_end)).Items ) {
        noopFunc := _noopFunc
        if methodType.FP.Equals(_env, self.processInfo, noopFunc, nil, nil){
            {
                _argList := node.FP.Get_argList(_env)
                if !Lns_IsNil( _argList ) {
                    argList := _argList.(*Nodes_ExpListNode)
                    for _, _exp := range( argList.FP.Get_expList(_env).Items ) {
                        exp := _exp
                        pre.Insert(WasmBin_drop(_env, self.FP.lowerExpAs(_env, exp, Ast_builtinTypeStem_)))
                    }
                }
            }
            return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
        }
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.G__lns_runtime_time, nil, nil){
        pre.Insert(convWasm_p_4_(_env, WasmBin_opI64Mul, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnOsTime(_env), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})), WasmBin_i64Const(_env, 1000)))))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeInt))
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Debug_getinfo, nil, nil){
        pre.Insert(WasmBin_refNull(_env, self.rt.FP.Get_tMap(_env)))
        return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Debug_getlocal, nil, nil){
        pre.Insert(WasmBin_refNull(_env, self.rt.FP.Get_tStr(_env)))
        pre.Insert(WasmBin_refNull(_env, WasmBin_htAny))
        return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Os_exit, nil, nil){
        var codeIns *WasmBin_Ins
        codeIns = WasmBin_i64Const(_env, 0)
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( argList.FP.Get_expList(_env).Len() >= 1) &&
                    _env.SetStackVal( argList.FP.Get_expList(_env).GetAt(1).FP.Get_kind(_env) != Nodes_NodeKind_get_Abbr(_env)) ).(bool)){
                    codeIns = WasmBin_call(_env, self.FP.fnIntOrZero(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeStem_))))
                }
            }
        }
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnOsExit(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](codeIns))))
        pre.Insert(WasmBin_unreachable(_env))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{})
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Os_remove, nil, nil){
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnOsFileOp(_env, "remove"), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString)))))
        return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Os_rename, nil, nil){
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnOsFileOp(_env, "rename"), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString), self.FP.lowerArgAs(_env, node, 2, Ast_builtinTypeString)))))
        return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Os_time, nil, nil){
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( argList.FP.Get_expList(_env).Len() >= 1) &&
                    _env.SetStackVal( argList.FP.Get_expList(_env).GetAt(1).FP.Get_kind(_env) != Nodes_NodeKind_get_Abbr(_env)) ).(bool)){
                    self.FP.error(_env, &node.Nodes_Node, "not support os.time with the argument")
                }
            }
        }
        pre.Insert(self.rt.FP.BoxI64(_env, WasmBin_call(_env, self.rt.FP.FnOsTime(_env), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))))
        return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Package_searchpath, nil, nil){
        return self.FP.lowerLuaCall(_env, node, WasmRt_luaSearchpath, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Os_clock, nil, nil){
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnOsClock(_env), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeReal))
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Io_popen, nil, nil){
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnIoPopen(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString)))))
        return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
    }
    if methodType.FP.Equals(_env, self.processInfo, builtin.Io_open, nil, nil){
        var modeIns *WasmBin_Ins
        modeIns = WasmBin_refNull(_env, self.rt.FP.Get_tStr(_env))
        {
            _argList := node.FP.Get_argList(_env)
            if !Lns_IsNil( _argList ) {
                argList := _argList.(*Nodes_ExpListNode)
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( argList.FP.Get_expList(_env).Len() >= 2) &&
                    _env.SetStackVal( argList.FP.Get_expList(_env).GetAt(2).FP.Get_kind(_env) != Nodes_NodeKind_get_Abbr(_env)) ).(bool)){
                    modeIns = self.FP.lowerArgAs(_env, node, 2, Ast_builtinTypeString.FP.Get_nilableTypeInfo(_env))
                }
            }
        }
        pre.Insert(WasmBin_call(_env, self.rt.FP.FnIoOpen(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeString), modeIns))))
        return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( methodType.FP.Equals(_env, self.processInfo, builtin.G__lns_sync_createFlag, nil, nil)) ||
        _env.SetStackVal( methodType.FP.Equals(_env, self.processInfo, builtin.G__lns_sync__createPipe, nil, nil)) ).(bool){
        pre.Insert(WasmBin_refNull(_env, WasmBin_htAny))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeStem.FP.Get_nilableTypeInfo(_env)))
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( methodType.FP.Equals(_env, self.processInfo, builtin.String__join, nil, nil)) ||
        _env.SetStackVal( methodType.FP.Equals(_env, self.processInfo, builtin.String___join, nil, nil)) ).(bool){
        var args *LnsList2_[*WasmBin_Ins]
        args = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        args.Insert(self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeStem_))
        {
            var _forFrom0 LnsInt = 2
            var _forTo0 LnsInt = 4
            for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                index := _forWork0
                args.Insert(self.FP.lowerArgAs(_env, node, index, Ast_builtinTypeStem_))
            }
        }
        pre.Insert(WasmBin_call(_env, self.FP.fnStrJoin(_env), args))
        return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](Ast_builtinTypeString))
    }
    var classType *Ast_TypeInfo
    classType = self.FP.getGenSrc(_env, prefix.FP.Get_expType(_env))
    {
        _enumType := Ast_EnumTypeInfoDownCastF(classType.FP)
        if !Lns_IsNil( _enumType ) {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            if name == "_from"{
                pre.Insert(self.FP.coerceRep(_env, WasmBin_call(_env, self.FP.fnEnumFrom(_env, enumType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, enumType.FP.Get_valTypeInfo(_env))))), WasmBin_typeAnyRef, self.FP.getRep(_env, enumType.FP.Get_nilableTypeInfo(_env))))
                return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](enumType.FP.Get_nilableTypeInfo(_env)))
            }
        }
    }
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( name == "_fromMap") ||
        _env.SetStackVal( name == "_fromStem") ).(bool){
        {
            _info := self.FP.getClassInfo(_env, classType)
            if !Lns_IsNil( _info ) {
                info := _info.(*convWasm_ClassInfo)
                pre.Insert(WasmBin_call(_env, self.FP.fnFromMap(_env, &node.Nodes_Node, info), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerArgAs(_env, node, 1, Ast_builtinTypeStem_)))))
                return WasmBin_seq(_env, pre), NewLnsList2_[*Ast_TypeInfo](Lns_2DDDGen[*Ast_TypeInfo](info.FP.Get_typeInfo(_env).FP.Get_nilableTypeInfo(_env), Ast_builtinTypeString.FP.Get_nilableTypeInfo(_env)))
            }
        }
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support static call -- %s", Lns_2DDD(name)))
// insert a dummy
    return nil,nil
}
// 3912: decl @lune.@base.@convWasm.convFilter.lowerCall
func (self *convWasm_convFilter) lowerCall(_env *LnsEnv, node *Nodes_ExpCallNode)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    {
        _prefix := self.FP.getNilAccPrefix(_env, &node.Nodes_Node)
        if !Lns_IsNil( _prefix ) {
            prefix := _prefix.(*Nodes_Node)
            return self.FP.lowerNilAccCall(_env, node, prefix)
        }
    }
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var funcNode *Nodes_Node
    funcNode = node.FP.Get_func(_env)
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(funcNode.FP)
        if !Lns_IsNil( _fieldNode ) {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            {
                _modType := self.FP.getExtModuleType(_env, fieldNode.FP.Get_prefix(_env).FP.Get_expType(_env))
                if !Lns_IsNil( _modType ) {
                    modType := _modType.(*Ast_TypeInfo)
                    var modVal LnsInt
                    modVal = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
                    var _func LnsInt
                    _func = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
                    pre.Insert(convWasm_ls_3_(_env, modVal, self.FP.lowerExtModuleVal(_env, &node.Nodes_Node, modType)))
                    pre.Insert(convWasm_ls_3_(_env, _func, WasmBin_call(_env, self.rt.FP.FnLuaGetAt(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, modVal), self.rt.FP.StrLiteral(_env, fieldNode.FP.Get_field(_env).Txt))))))
                    var headArgs *LnsList2_[*WasmBin_Ins]
                    headArgs = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
                    if funcNode.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Method{
                        headArgs.Insert(convWasm_lg_2_(_env, modVal))
                    }
                    return self.FP.lowerLuaCallSub(_env, node, 0, convWasm_lg_2_(_env, _func), headArgs, pre)
                }
            }
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.FP.isLuaval(_env, fieldNode.FP.Get_prefix(_env).FP.Get_expType(_env))) &&
                _env.SetStackVal( Lns_op_not(self.FP.isStreamType(_env, fieldNode.FP.Get_prefix(_env).FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_extedType(_env)))) ).(bool)){
                var obj LnsInt
                obj = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
                var _func LnsInt
                _func = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
                pre.Insert(convWasm_ls_3_(_env, obj, self.FP.lowerExpAs(_env, fieldNode.FP.Get_prefix(_env), Ast_builtinTypeStem_)))
                pre.Insert(convWasm_ls_3_(_env, _func, WasmBin_call(_env, self.rt.FP.FnLuaGetAt(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj), self.rt.FP.StrLiteral(_env, fieldNode.FP.Get_field(_env).Txt))))))
                return self.FP.lowerLuaCallSub(_env, node, 0, convWasm_lg_2_(_env, _func), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj))), pre)
            }
        }
    }
    if self.FP.isLuaval(_env, funcNode.FP.Get_expType(_env)){
        var _func LnsInt
        _func = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
        pre.Insert(convWasm_ls_3_(_env, _func, self.FP.lowerExpAs(_env, funcNode, Ast_builtinTypeStem_)))
        return self.FP.lowerLuaCallSub(_env, node, 0, convWasm_lg_2_(_env, _func), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), pre)
    }
    {
        _refNode := Nodes_ExpRefNodeDownCastF(funcNode.FP)
        if !Lns_IsNil( _refNode ) {
            refNode := _refNode.(*Nodes_ExpRefNode)
            return self.FP.lowerCallRef(_env, node, refNode, pre)
        }
    }
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(funcNode.FP)
        if !Lns_IsNil( _fieldNode ) {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            return self.FP.lowerCallField(_env, node, fieldNode, pre)
        }
    }
    var funcVal *convWasm_Val
    funcVal = self.FP.lowerVal(_env, funcNode)
    pre.Insert(self.FP.callClosure(_env, funcVal.FP.Get_ins(_env), funcVal.FP.Get_typeInfo(_env), node.FP.Get_argList(_env), pre))
    return WasmBin_seq(_env, pre), self.FP.normType(_env, funcVal.FP.Get_typeInfo(_env)).FP.Get_retTypeInfoList(_env)
}
// 3962: decl @lune.@base.@convWasm.convFilter.lowerCallSuper
func (self *convWasm_convFilter) lowerCallSuper(_env *LnsEnv, node *Nodes_ExpCallSuperNode)(*WasmBin_Ins, *LnsList2_[*Ast_TypeInfo]) {
    var methodType *Ast_TypeInfo
    methodType = node.FP.Get_methodType(_env)
    var funcIndex LnsInt
    
    {
        _funcIndex := self.func2index.Get(methodType)
        if _funcIndex == nil{
            self.FP.error(_env, &node.Nodes_Node, "not found super method")
        } else {
            funcIndex = _funcIndex.(LnsInt)
        }
    }
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, self.ctx.FP.Get_selfLocal(_env))))
    convWasm_append_17_(_env, args, self.FP.lowerArgs(_env, node.FP.Get_expList(_env), methodType.FP.Get_argTypeInfoList(_env), pre))
    pre.Insert(WasmBin_call(_env, funcIndex, args))
    return WasmBin_seq(_env, pre), methodType.FP.Get_retTypeInfoList(_env)
}
// 3975: decl @lune.@base.@convWasm.convFilter.lowerNew
func (self *convWasm_convFilter) lowerNew(_env *LnsEnv, node *Nodes_ExpNewNode) *convWasm_Val {
    var info *convWasm_ClassInfo
    
    {
        _info := self.FP.getClassInfo(_env, node.FP.Get_expType(_env))
        if _info == nil{
            self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not found class -- %s", Lns_2DDD(node.FP.Get_expType(_env).FP.GetTxt(_env, nil, nil, nil))))
        } else {
            info = _info.(*convWasm_ClassInfo)
        }
    }
    var objRep *WasmBin_ValType
    objRep = WasmBin_refType(_env, info.FP.Get_structType(_env))
    var obj LnsInt
    obj = self.FP.tempLocal(_env, objRep)
    var fields *LnsList2_[*WasmBin_Ins]
    fields = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), WasmBin_globalGet(_env, info.FP.Get_vtGlobal(_env))))
    for _, _rep := range( info.FP.Get_fieldReps(_env).Items ) {
        rep := _rep
        fields.Insert(self.FP.getDefaultVal(_env, rep))
    }
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, obj, WasmBin_structNew(_env, info.FP.Get_structType(_env), fields))))
    var ctorType *Ast_TypeInfo
    ctorType = self.FP.getCtorType(_env, info.FP.Get_typeInfo(_env))
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj)))
    convWasm_append_17_(_env, args, self.FP.lowerArgs(_env, node.FP.Get_argList(_env), ctorType.FP.Get_argTypeInfoList(_env), pre))
    pre.Insert(WasmBin_call(_env, info.FP.Get_ctorFunc(_env), args))
    pre.Insert(convWasm_lg_2_(_env, obj))
    return NewconvWasm_Val(_env, WasmBin_seq(_env, pre), objRep, node.FP.Get_expType(_env))
}
// 3995: decl @lune.@base.@convWasm.convFilter.algeTxt
func (self *convWasm_convFilter) algeTxt(_env *LnsEnv, prefix *Nodes_Node,prefixType *Ast_TypeInfo) *WasmBin_Ins {
    var algeType *Ast_TypeInfo
    algeType = self.FP.getGenSrc(_env, prefixType)
    var info *convWasm_AlgeInfo
    info = Lns_unwrap( self.FP.getAlgeInfo(_env, algeType)).(*convWasm_AlgeInfo)
    return WasmBin_refCast(_env, self.rt.FP.Get_tStr(_env), WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_globalGet(_env, info.FP.Get_namesGlobal(_env)), WasmBin_structGet(_env, info.FP.Get_baseType(_env), 1, self.FP.lowerExpAs(_env, prefix, algeType))))
}
// 4005: decl @lune.@base.@convWasm.convFilter.getCtorType
func (self *convWasm_convFilter) getCtorType(_env *LnsEnv, classType *Ast_TypeInfo) *Ast_TypeInfo {
    var scope *Ast_Scope
    scope = Lns_unwrap( classType.FP.Get_scope(_env)).(*Ast_Scope)
    return (Lns_unwrap( scope.FP.Get_symbol2SymbolInfoMap(_env).Get("__init")).(*Ast_SymbolInfo)).FP.Get_typeInfo(_env)
}
// 4011: decl @lune.@base.@convWasm.convFilter.lowerGetField
func (self *convWasm_convFilter) lowerGetField(_env *LnsEnv, node *Nodes_GetFieldNode) *convWasm_Val {
    var prefix *Nodes_Node
    prefix = node.FP.Get_prefix(_env)
    var getterType *Ast_TypeInfo
    getterType = node.FP.Get_getterTypeInfo(_env)
    var fieldName string
    fieldName = convWasm_convExp4_958(Lns_2DDD(_env.GetVM().String_gsub(node.FP.Get_field(_env).Txt,"^%$", "")))
    var getterName string
    getterName = "get_" + fieldName
    var prefixType *Ast_TypeInfo
    prefixType = self.FP.normType(_env, prefix.FP.Get_expType(_env))
    {
        _enumType := Ast_EnumTypeInfoDownCastF(prefixType.FP)
        if !Lns_IsNil( _enumType ) {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            if _switch0 := getterName; _switch0 == "get__txt" {
                return self.FP.newVal(_env, WasmBin_call(_env, self.FP.fnEnumTxt(_env, enumType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, prefix)))), Ast_builtinTypeString)
            } else if _switch0 == "get__allList" {
                return self.FP.newVal(_env, WasmBin_call(_env, self.FP.fnEnumAllList(_env, enumType), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})), node.FP.Get_expType(_env))
            }
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Alge) &&
        _env.SetStackVal( getterName == "get__txt") ).(bool)){
        return self.FP.newVal(_env, self.FP.algeTxt(_env, prefix, prefixType), Ast_builtinTypeString)
    }
    if self.FP.getObjType(_env, prefix.FP.Get_expType(_env)).FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
        var ins *WasmBin_Ins
        var retTypes *LnsList2_[*Ast_TypeInfo]
        ins,retTypes = self.FP.lowerIfMethodCall(_env, prefix, getterName, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), nil)
        return self.FP.newVal(_env, ins, retTypes.GetAt(1))
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( getterType.FP.Get_staticFlag(_env)) &&
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( getterType.FP.Get_autoFlag(_env)) ||
            _env.SetStackVal( Lns_op_not(self.declaredMethodSet.Has(getterType))) ).(bool))) ).(bool)){
        {
            _scope := self.FP.getGenSrc(_env, prefix.FP.Get_expType(_env)).FP.Get_scope(_env)
            if !Lns_IsNil( _scope ) {
                scope := _scope.(*Ast_Scope)
                {
                    _symbol := scope.FP.Get_symbol2SymbolInfoMap(_env).Get(fieldName)
                    if !Lns_IsNil( _symbol ) {
                        symbol := _symbol.(*Ast_SymbolInfo)
                        var orgSym *Ast_SymbolInfo
                        orgSym = symbol.FP.GetOrg(_env)
                        if Lns_isCondTrue( self.globalSym2loc.Get(orgSym)){
                            return NewconvWasm_Val(_env, self.FP.readSym(_env, &node.Nodes_Node, orgSym), self.FP.getRep(_env, orgSym.FP.Get_typeInfo(_env)), orgSym.FP.Get_typeInfo(_env))
                        }
                    }
                }
            }
        }
    }
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( getterType.FP.Get_autoFlag(_env)) ||
            _env.SetStackVal( Lns_op_not(self.declaredMethodSet.Has(getterType))) ).(bool))) &&
        _env.SetStackVal( Lns_op_not(self.advertiseMethodSet.Has(getterType))) &&
        _env.SetStackVal( Lns_op_not(self.dynGetterNameSet.Has(getterName))) ).(bool)){
        return self.FP.lowerMemberRef(_env, &node.Nodes_Node, prefix, fieldName)
    }
    if getterType.FP.Get_staticFlag(_env){
        {
            _funcIndex := self.func2index.Get(getterType)
            if !Lns_IsNil( _funcIndex ) {
                funcIndex := _funcIndex.(LnsInt)
                return self.FP.newVal(_env, WasmBin_call(_env, funcIndex, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, self.rt.FP.Get_tClo(_env))))), getterType.FP.Get_retTypeInfoList(_env).GetAt(1))
            }
        }
    }
    var objType *Ast_TypeInfo
    objType = self.FP.getObjType(_env, prefix.FP.Get_expType(_env))
    var info *convWasm_ClassInfo
    
    {
        _info := self.FP.getClassInfo(_env, objType)
        if _info == nil{
            self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not found class -- %s", Lns_2DDD(objType.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            info = _info.(*convWasm_ClassInfo)
        }
    }
    var slotIndex LnsInt
    
    {
        _slotIndex := info.FP.Get_name2slot(_env).Get(getterName)
        if _slotIndex == nil{
            self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not found getter -- %s", Lns_2DDD(getterName)))
        } else {
            slotIndex = _slotIndex.(LnsInt)
        }
    }
    var slot *convWasm_SlotInfo
    slot = info.FP.Get_slotList(_env).GetAt(slotIndex)
    var obj LnsInt
    obj = self.FP.tempLocal(_env, WasmBin_refType(_env, info.FP.Get_structType(_env)))
    var retType *Ast_TypeInfo
    retType = slot.FP.Get_methodType(_env).FP.Get_retTypeInfoList(_env).GetAt(1)
    return self.FP.newVal(_env, WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, obj, self.FP.lowerExpAs(_env, prefix, info.FP.Get_typeInfo(_env))), WasmBin_callRef(_env, slot.FP.Get_funcType(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, obj), WasmBin_structGet(_env, info.FP.Get_vtType(_env), slotIndex, WasmBin_structGet(_env, info.FP.Get_structType(_env), WasmRt_fObjVt, convWasm_lg_2_(_env, obj))))))))), retType)
}
// 4089: decl @lune.@base.@convWasm.convFilter.getEnumValList
func (self *convWasm_convFilter) getEnumValList(_env *LnsEnv, enumType *Ast_EnumTypeInfo) *LnsList2_[*Ast_EnumValInfo] {
    var list *LnsList2_[*Ast_EnumValInfo]
    list = NewLnsList2_[*Ast_EnumValInfo]([]*Ast_EnumValInfo{})
    {
        _declNode := self.enumType2decl.Get(&enumType.Ast_TypeInfo)
        if !Lns_IsNil( _declNode ) {
            declNode := _declNode.(*Nodes_DeclEnumNode)
            for _, _token := range( declNode.FP.Get_valueNameList(_env).Items ) {
                token := _token
                {
                    _valInfo := enumType.FP.Get_name2EnumValInfo(_env).Get(token.Txt)
                    if !Lns_IsNil( _valInfo ) {
                        valInfo := _valInfo.(*Ast_EnumValInfo)
                        list.Insert(valInfo)
                    }
                }
            }
        } else {
            {
                __forsortCollection0 := enumType.FP.Get_name2EnumValInfo(_env)
                __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
                __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
                for _, ___forsortKey0 := range( __forsortSorted0.Items ) {
                    valInfo := __forsortCollection0.Items[ ___forsortKey0 ]
                    list.Insert(valInfo)
                }
            }
        }
    }
    return list
}
// 4105: decl @lune.@base.@convWasm.convFilter.getEnumFuncName
func (self *convWasm_convFilter) getEnumFuncName(_env *LnsEnv, enumType *Ast_EnumTypeInfo,kind string) string {
    return _env.GetVM().String_format("enum%s:%d", Lns_2DDD(kind, enumType.FP.Get_typeId(_env).Id))
}
// 4110: decl @lune.@base.@convWasm.convFilter.enumEq
func (self *convWasm_convFilter) enumEq(_env *LnsEnv, enumType *Ast_EnumTypeInfo,val *WasmBin_Ins,valInfo *Ast_EnumValInfo) *WasmBin_Ins {
    var literal *convWasm_Val
    literal = self.FP.lowerEnumVal(_env, enumType, valInfo)
    return self.FP.lowerEqVals(_env, NewconvWasm_Val(_env, val, literal.FP.Get_rep(_env), &enumType.Ast_TypeInfo), literal)
}
// 4118: decl @lune.@base.@convWasm.convFilter.fnEnumTxt
func (self *convWasm_convFilter) fnEnumTxt(_env *LnsEnv, enumType *Ast_EnumTypeInfo) LnsInt {
    var name string
    name = self.FP.getEnumFuncName(_env, enumType, "txt")
    {
        _index := self.adapterMap.Get(name)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var rep *WasmBin_ValType
    rep = self.FP.getRep(_env, &enumType.Ast_TypeInfo)
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](rep)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_strType(_env))))
    self.adapterMap.Set(name,index)
    self.mod.FP.SetFuncName(_env, index, name)
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _, _valInfo := range( self.FP.getEnumValList(_env, enumType).Items ) {
        valInfo := _valInfo
        body.Insert(WasmBin_ifElse(_env, self.FP.enumEq(_env, enumType, convWasm_lg_2_(_env, 0), valInfo), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.StrLiteral(_env, _env.GetVM().String_format("%s.%s", Lns_2DDD(enumType.FP.Get_rawTxt(_env), valInfo.FP.Get_name(_env))))))))), nil, nil))
    }
    body.Insert(self.rt.FP.StrLiteral(_env, ""))
    _func.FP.SetBody(_env, body)
    return index
}
// 4140: decl @lune.@base.@convWasm.convFilter.fnEnumFrom
func (self *convWasm_convFilter) fnEnumFrom(_env *LnsEnv, enumType *Ast_EnumTypeInfo) LnsInt {
    var name string
    name = self.FP.getEnumFuncName(_env, enumType, "from")
    {
        _index := self.adapterMap.Get(name)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var rep *WasmBin_ValType
    rep = self.FP.getRep(_env, &enumType.Ast_TypeInfo)
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](rep)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)))
    self.adapterMap.Set(name,index)
    self.mod.FP.SetFuncName(_env, index, name)
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _, _valInfo := range( self.FP.getEnumValList(_env, enumType).Items ) {
        valInfo := _valInfo
        body.Insert(WasmBin_ifElse(_env, self.FP.enumEq(_env, enumType, convWasm_lg_2_(_env, 0), valInfo), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.coerceRep(_env, convWasm_lg_2_(_env, 0), rep, WasmBin_typeAnyRef)))))), nil, nil))
    }
    body.Insert(WasmBin_refNull(_env, WasmBin_htAny))
    _func.FP.SetBody(_env, body)
    return index
}
// 4161: decl @lune.@base.@convWasm.convFilter.fnEnumAllList
func (self *convWasm_convFilter) fnEnumAllList(_env *LnsEnv, enumType *Ast_EnumTypeInfo) LnsInt {
    var name string
    name = self.FP.getEnumFuncName(_env, enumType, "all")
    {
        _index := self.adapterMap.Get(name)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_listType(_env))))
    self.adapterMap.Set(name,index)
    self.mod.FP.SetFuncName(_env, index, name)
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    for _, _valInfo := range( self.FP.getEnumValList(_env, enumType).Items ) {
        valInfo := _valInfo
        vals.Insert(self.FP.lowerEnumVal(_env, enumType, valInfo))
    }
    var saved *convWasm_FuncCtx
    saved = self.ctx
    self.ctx = NewconvWasm_FuncCtx(_env, _func, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_listType(_env))), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{}))
    var ins *WasmBin_Ins
    ins = self.FP.buildList(_env, vals)
    self.ctx = saved
    _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](ins)))
    return index
}
// 4189: decl @lune.@base.@convWasm.convFilter.needToMapConv
func (self *convWasm_convFilter) needToMapConv(_env *LnsEnv, typeInfo *Ast_TypeInfo) bool {
    var work *Ast_TypeInfo
    work = self.FP.normType(_env, typeInfo)
    if _switch0 := work.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array || _switch0 == Ast_TypeInfoKind__Map {
        for _, _item := range( work.FP.Get_itemTypeInfoList(_env).Items ) {
            item := _item
            if self.FP.needToMapConv(_env, item){
                return true
            }
        }
        return false
    } else if _switch0 == Ast_TypeInfoKind__Class {
        if self.FP.isType(_env, work, Ast_builtinTypeString){
            return false
        }
        return convWasm_ClassInfo2Stem(self.FP.getClassInfo(_env, work)) != nil
    } else if _switch0 == Ast_TypeInfoKind__Alge {
        return true
    }
    return false
}
// 4214: decl @lune.@base.@convWasm.convFilter.mapConvKey
func (self *convWasm_convFilter) mapConvKey(_env *LnsEnv, typeInfo *Ast_TypeInfo) string {
    var work *Ast_TypeInfo
    work = self.FP.normType(_env, typeInfo)
    var key string
    key = _env.GetVM().String_format("%d", Lns_2DDD(self.FP.getGenSrc(_env, work).FP.Get_typeId(_env).Id))
    if work.FP.Get_itemTypeInfoList(_env).Len() > 0{
        key = key + "<"
        for _index, _item := range( work.FP.Get_itemTypeInfoList(_env).Items ) {
            index := _index + 1
            item := _item
            if index > 1{
                key = key + ","
            }
            key = key + self.FP.mapConvKey(_env, item)
            if item.FP.Get_nilable(_env){
                key = key + "!"
            }
        }
        key = key + ">"
    }
    return key
}
// 4234: decl @lune.@base.@convWasm.convFilter.fnToMapConv
func (self *convWasm_convFilter) fnToMapConv(_env *LnsEnv, node *Nodes_Node,typeInfo *Ast_TypeInfo) LnsInt {
    var name string
    name = "toMapConv:" + self.FP.mapConvKey(_env, typeInfo)
    {
        _index := self.adapterMap.Get(name)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)))
    self.adapterMap.Set(name,index)
    self.mod.FP.SetFuncName(_env, index, name)
    var work *Ast_TypeInfo
    work = self.FP.normType(_env, typeInfo)
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, WasmBin_htAny)))))), nil, nil)))
    if _switch0 := work.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        var elemType *Ast_TypeInfo
        elemType = work.FP.Get_itemTypeInfoList(_env).GetAt(1)
        var src LnsInt
        src = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
        var result LnsInt
        result = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
        var pos LnsInt
        pos = _func.FP.AddLocal(_env, WasmBin_typeI32)
        body.Insert(convWasm_ls_3_(_env, src, WasmBin_refCast(_env, self.rt.FP.Get_tList(_env), convWasm_lg_2_(_env, 0))))
        body.Insert(convWasm_ls_3_(_env, result, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, convWasm_lg_2_(_env, src)))))))
        body.Insert(WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, convWasm_lg_2_(_env, src)))))), WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, result), self.FP.toMapVal(_env, node, elemType, WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListItems, convWasm_lg_2_(_env, src)), convWasm_lg_2_(_env, pos)))))), convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))), WasmBin_br(_env, 0))), nil))), nil))
        body.Insert(convWasm_lg_2_(_env, result))
    } else if _switch0 == Ast_TypeInfoKind__Map {
        var keyType *Ast_TypeInfo
        keyType = work.FP.Get_itemTypeInfoList(_env).GetAt(1)
        var valType *Ast_TypeInfo
        valType = work.FP.Get_itemTypeInfoList(_env).GetAt(2)
        var src LnsInt
        src = _func.FP.AddLocal(_env, self.rt.FP.Get_mapType(_env))
        var result LnsInt
        result = _func.FP.AddLocal(_env, self.rt.FP.Get_mapType(_env))
        var pos LnsInt
        pos = _func.FP.AddLocal(_env, WasmBin_typeI32)
        body.Insert(convWasm_ls_3_(_env, src, WasmBin_refCast(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0))))
        body.Insert(convWasm_ls_3_(_env, result, self.rt.FP.NewMap(_env)))
        body.Insert(WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, pos, WasmBin_call(_env, self.rt.FP.FnMapNext(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, src), convWasm_lg_2_(_env, pos))))), WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 0))))), WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, result), self.FP.toMapVal(_env, node, keyType, self.rt.FP.MapKeyAt(_env, convWasm_lg_2_(_env, src), convWasm_lg_2_(_env, pos))), self.FP.toMapVal(_env, node, valType, self.rt.FP.MapValAt(_env, convWasm_lg_2_(_env, src), convWasm_lg_2_(_env, pos)))))), convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))), WasmBin_br(_env, 0))), nil))), nil))
        body.Insert(convWasm_lg_2_(_env, result))
    } else if _switch0 == Ast_TypeInfoKind__Alge {
        var algeType *Ast_TypeInfo
        algeType = self.FP.getGenSrc(_env, work)
        var algeInfo *convWasm_AlgeInfo
        algeInfo = Lns_unwrap( self.FP.getAlgeInfo(_env, algeType)).(*convWasm_AlgeInfo)
        var algeWork *Ast_AlgeTypeInfo
        algeWork = Lns_unwrap( Ast_AlgeTypeInfoDownCastF(algeType.FP)).(*Ast_AlgeTypeInfo)
        var src LnsInt
        src = _func.FP.AddLocal(_env, WasmBin_refType(_env, algeInfo.FP.Get_baseType(_env)))
        var result LnsInt
        result = _func.FP.AddLocal(_env, self.rt.FP.Get_mapType(_env))
        var params LnsInt
        params = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
        var tag LnsInt
        tag = _func.FP.AddLocal(_env, WasmBin_typeI32)
        body.Insert(convWasm_ls_3_(_env, src, WasmBin_refCast(_env, algeInfo.FP.Get_baseType(_env), convWasm_lg_2_(_env, 0))))
        body.Insert(convWasm_ls_3_(_env, tag, WasmBin_structGet(_env, algeInfo.FP.Get_baseType(_env), 1, convWasm_lg_2_(_env, src))))
        body.Insert(convWasm_ls_3_(_env, result, self.rt.FP.NewMap(_env)))
        for _, _valName := range( algeInfo.FP.Get_valNames(_env).Items ) {
            valName := _valName
            var valInfo *Ast_AlgeValInfo
            valInfo = Lns_unwrap( algeWork.FP.Get_valInfoMap(_env).Get(valName)).(*Ast_AlgeValInfo)
            var valType LnsInt
            valType = Lns_unwrap( algeInfo.FP.Get_name2type(_env).Get(valName)).(LnsInt)
            var thenBody *LnsList2_[*WasmBin_Ins]
            thenBody = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, result), self.rt.FP.BoxI64(_env, WasmBin_i64Const(_env, 1)), self.rt.FP.StrLiteral(_env, valName))))))
            if valInfo.FP.Get_typeList(_env).Len() > 0{
                thenBody.Insert(convWasm_ls_3_(_env, params, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, valInfo.FP.Get_typeList(_env).Len()))))))
                for _paramIndex, _paramType := range( valInfo.FP.Get_typeList(_env).Items ) {
                    paramIndex := _paramIndex + 1
                    paramType := _paramType
                    var paramVal *WasmBin_Ins
                    paramVal = self.FP.coerceRep(_env, WasmBin_structGet(_env, valType, paramIndex + 1, WasmBin_refCast(_env, valType, convWasm_lg_2_(_env, src))), self.FP.getRep(_env, paramType), WasmBin_typeAnyRef)
                    thenBody.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, params), self.FP.toMapVal(_env, node, paramType, paramVal)))))
                }
                thenBody.Insert(WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, result), self.rt.FP.BoxI64(_env, WasmBin_i64Const(_env, 2)), convWasm_lg_2_(_env, params)))))
            }
            body.Insert(WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, tag), WasmBin_i32Const(_env, Lns_unwrap( algeInfo.FP.Get_name2tag(_env).Get(valName)).(LnsInt))))), thenBody, nil, nil))
        }
        body.Insert(convWasm_lg_2_(_env, result))
    } else {
        var info *convWasm_ClassInfo
        
        {
            _info := self.FP.getClassInfo(_env, work)
            if _info == nil{
                self.FP.error(_env, node, _env.GetVM().String_format("not support _toMap -- %s", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil))))
            } else {
                info = _info.(*convWasm_ClassInfo)
            }
        }
        body.Insert(WasmBin_call(_env, self.FP.fnToMap(_env, node, info), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refCast(_env, info.FP.Get_structType(_env), convWasm_lg_2_(_env, 0))))))
    }
    _func.FP.SetBody(_env, body)
    return index
}
// 4338: decl @lune.@base.@convWasm.convFilter.toMapVal
func (self *convWasm_convFilter) toMapVal(_env *LnsEnv, node *Nodes_Node,typeInfo *Ast_TypeInfo,val *WasmBin_Ins) *WasmBin_Ins {
    if Lns_op_not(self.FP.needToMapConv(_env, typeInfo)){
        return val
    }
    return WasmBin_call(_env, self.FP.fnToMapConv(_env, node, typeInfo), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}
// 4346: decl @lune.@base.@convWasm.convFilter.fnToMap
func (self *convWasm_convFilter) fnToMap(_env *LnsEnv, node *Nodes_Node,info *convWasm_ClassInfo) LnsInt {
    var name string
    name = _env.GetVM().String_format("toMap:%d", Lns_2DDD(info.FP.Get_structType(_env)))
    {
        _index := self.adapterMap.Get(name)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_refType(_env, info.FP.Get_structType(_env)))), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_mapType(_env))))
    self.adapterMap.Set(name,index)
    self.mod.FP.SetFuncName(_env, index, name)
    var _map LnsInt
    _map = _func.FP.AddLocal(_env, self.rt.FP.Get_mapType(_env))
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, _map, self.rt.FP.NewMap(_env))))
    for _fieldIndex, _fieldName := range( info.FP.Get_fieldNames(_env).Items ) {
        fieldIndex := _fieldIndex + 1
        fieldName := _fieldName
        var rep *WasmBin_ValType
        rep = info.FP.Get_fieldReps(_env).GetAt(fieldIndex)
        var val *WasmBin_Ins
        val = self.FP.coerceRep(_env, WasmBin_structGet(_env, info.FP.Get_structType(_env), fieldIndex + 1, convWasm_lg_2_(_env, 0)), rep, WasmBin_typeAnyRef)
        body.Insert(WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, _map), self.rt.FP.StrLiteral(_env, fieldName), self.FP.toMapVal(_env, node, info.FP.Get_fieldTypes(_env).GetAt(fieldIndex), val)))))
    }
    body.Insert(convWasm_lg_2_(_env, _map))
    _func.FP.SetBody(_env, body)
    return index
}
// 4375: decl @lune.@base.@convWasm.convFilter.fnFromConv
func (self *convWasm_convFilter) fnFromConv(_env *LnsEnv, node *Nodes_Node,typeInfo *Ast_TypeInfo) LnsInt {
    var name string
    name = "fromConv:" + self.FP.mapConvKey(_env, typeInfo)
    {
        _index := self.adapterMap.Get(name)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var strNull *WasmBin_Ins
    strNull = WasmBin_refNull(_env, self.rt.FP.Get_tStr(_env))
    var anyNull *WasmBin_Ins
    anyNull = WasmBin_refNull(_env, WasmBin_htAny)
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef, self.rt.FP.Get_strType(_env))))
    self.adapterMap.Set(name,index)
    self.mod.FP.SetFuncName(_env, index, name)
    var work *Ast_TypeInfo
    work = self.FP.normType(_env, typeInfo)
    var fail *WasmBin_Ins
    fail = WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](anyNull, strNull)))
    var convWasm_ok func(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins
    convWasm_ok = func(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins {
        return WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val, strNull)))
    }
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var convWasm_posMess func(_env *LnsEnv, posTxt *WasmBin_Ins,mess *WasmBin_Ins,messLocal LnsInt) *WasmBin_Ins
    convWasm_posMess = func(_env *LnsEnv, posTxt *WasmBin_Ins,mess *WasmBin_Ins,messLocal LnsInt) *WasmBin_Ins {
        return WasmBin_ifElse(_env, WasmBin_refIsNull(_env, mess), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](posTxt)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnStrConcat(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](posTxt, WasmBin_call(_env, self.rt.FP.FnStrConcat(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.StrLiteral(_env, "."), convWasm_lg_2_(_env, messLocal))))))))), self.rt.FP.Get_strType(_env))
    }
    {
        _enumType := Ast_EnumTypeInfoDownCastF(work.FP)
        if !Lns_IsNil( _enumType ) {
            enumType := _enumType.(*Ast_EnumTypeInfo)
            var valType *Ast_TypeInfo
            valType = enumType.FP.Get_valTypeInfo(_env)
            var val LnsInt
            val = _func.FP.AddLocal(_env, WasmBin_typeAnyRef)
            body.Insert(WasmBin_call(_env, self.FP.fnFromConv(_env, node, valType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0)))))
            body.Insert(WasmBin_drop(_env, WasmBin_nop(_env)))
            body.Insert(convWasm_ls_3_(_env, val, WasmBin_nop(_env)))
            body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, val)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](fail))))
            body.Insert(convWasm_ok(_env, self.FP.coerceRep(_env, WasmBin_call(_env, self.FP.fnEnumFrom(_env, enumType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.coerceRep(_env, convWasm_lg_2_(_env, val), WasmBin_typeAnyRef, self.FP.getRep(_env, valType))))), self.FP.getRep(_env, enumType.FP.Get_nilableTypeInfo(_env)), WasmBin_typeAnyRef)))
        } else {
            if self.FP.isType(_env, work, Ast_builtinTypeInt){
                body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxI64(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ok(_env, convWasm_lg_2_(_env, 0))))))
                body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxF64(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ok(_env, self.rt.FP.BoxI64(_env, convWasm_p_4_(_env, WasmBin_opI64TruncF64S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opF64Floor, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.UnboxF64(_env, convWasm_lg_2_(_env, 0))))))))))))))
                body.Insert(fail)
            } else if self.FP.isType(_env, work, Ast_builtinTypeReal){
                body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxF64(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ok(_env, convWasm_lg_2_(_env, 0))))))
                body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxI64(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ok(_env, self.rt.FP.BoxF64(_env, convWasm_p_4_(_env, WasmBin_opF64ConvertI64S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.UnboxI64(_env, convWasm_lg_2_(_env, 0)))))))))))
                body.Insert(fail)
            } else if self.FP.isType(_env, work, Ast_builtinTypeBool){
                body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refTest(_env, self.rt.FP.Get_tBoxBool(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ok(_env, convWasm_lg_2_(_env, 0))))))
                body.Insert(fail)
            } else if self.FP.isType(_env, work, Ast_builtinTypeString){
                body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refTest(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ok(_env, convWasm_lg_2_(_env, 0))))))
                body.Insert(fail)
            } else if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( self.FP.isType(_env, work, Ast_builtinTypeStem)) ||
                _env.SetStackVal( work.FP.Get_kind(_env) == Ast_TypeInfoKind__Alternate) ).(bool){
                body.Insert(convWasm_ok(_env, convWasm_lg_2_(_env, 0)))
            } else { 
                if _switch0 := work.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
                    var elemType *Ast_TypeInfo
                    elemType = work.FP.Get_itemTypeInfoList(_env).GetAt(1)
                    var src LnsInt
                    src = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
                    var result LnsInt
                    result = _func.FP.AddLocal(_env, self.rt.FP.Get_listType(_env))
                    var pos LnsInt
                    pos = _func.FP.AddLocal(_env, WasmBin_typeI32)
                    var val LnsInt
                    val = _func.FP.AddLocal(_env, WasmBin_typeAnyRef)
                    var mess LnsInt
                    mess = _func.FP.AddLocal(_env, self.rt.FP.Get_strType(_env))
                    body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refTest(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, 0, WasmBin_call(_env, self.FP.fnMapToList(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refCast(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0))))))))))
                    body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refTest(_env, self.rt.FP.Get_tList(_env), convWasm_lg_2_(_env, 0))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](fail))))
                    body.Insert(convWasm_ls_3_(_env, src, WasmBin_refCast(_env, self.rt.FP.Get_tList(_env), convWasm_lg_2_(_env, 0))))
                    body.Insert(convWasm_ls_3_(_env, result, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, convWasm_lg_2_(_env, src)))))))
                    var loopBody *LnsList2_[*WasmBin_Ins]
                    loopBody = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, convWasm_lg_2_(_env, src)))))), WasmBin_call(_env, self.FP.fnFromConv(_env, node, elemType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListItems, convWasm_lg_2_(_env, src)), convWasm_lg_2_(_env, pos))))), convWasm_ls_3_(_env, mess, WasmBin_nop(_env)), convWasm_ls_3_(_env, val, WasmBin_nop(_env)), convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1)))))))
                    if Lns_op_not(elemType.FP.Get_nilable(_env)){
                        loopBody.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, val)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](anyNull, convWasm_posMess(_env, WasmBin_call(_env, self.rt.FP.FnI64ToStr(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI64ExtendI32S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos))))))), convWasm_lg_2_(_env, mess), mess))))))))
                    }
                    loopBody.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, result), convWasm_lg_2_(_env, val)))))
                    loopBody.Insert(WasmBin_br(_env, 0))
                    body.Insert(WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, loopBody, nil))), nil))
                    body.Insert(convWasm_ok(_env, convWasm_lg_2_(_env, result)))
                } else if _switch0 == Ast_TypeInfoKind__Map || _switch0 == Ast_TypeInfoKind__Set {
                    var isSet bool
                    isSet = work.FP.Get_kind(_env) == Ast_TypeInfoKind__Set
                    var keyType *Ast_TypeInfo
                    keyType = work.FP.Get_itemTypeInfoList(_env).GetAt(1)
                    var src LnsInt
                    src = _func.FP.AddLocal(_env, self.rt.FP.Get_mapType(_env))
                    var result LnsInt
                    result = _func.FP.AddLocal(_env, self.rt.FP.Get_mapType(_env))
                    var pos LnsInt
                    pos = _func.FP.AddLocal(_env, WasmBin_typeI32)
                    var key LnsInt
                    key = _func.FP.AddLocal(_env, WasmBin_typeAnyRef)
                    var val LnsInt
                    val = _func.FP.AddLocal(_env, WasmBin_typeAnyRef)
                    var mess LnsInt
                    mess = _func.FP.AddLocal(_env, self.rt.FP.Get_strType(_env))
                    body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refTest(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](fail))))
                    body.Insert(convWasm_ls_3_(_env, src, WasmBin_refCast(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0))))
                    body.Insert(convWasm_ls_3_(_env, result, self.rt.FP.NewMap(_env)))
                    var loopBody *LnsList2_[*WasmBin_Ins]
                    loopBody = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, pos, WasmBin_call(_env, self.rt.FP.FnMapNext(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, src), convWasm_lg_2_(_env, pos))))), WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 0))))), WasmBin_call(_env, self.FP.fnFromConv(_env, node, keyType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.MapKeyAt(_env, convWasm_lg_2_(_env, src), convWasm_lg_2_(_env, pos))))), WasmBin_drop(_env, WasmBin_nop(_env)), convWasm_ls_3_(_env, key, WasmBin_nop(_env)), convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, key)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](fail)))))
                    if isSet{
                        loopBody.Insert(convWasm_ls_3_(_env, val, convWasm_lg_2_(_env, key)))
                    } else { 
                        loopBody.Insert(WasmBin_call(_env, self.FP.fnFromConv(_env, node, work.FP.Get_itemTypeInfoList(_env).GetAt(2)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.MapValAt(_env, convWasm_lg_2_(_env, src), convWasm_lg_2_(_env, pos))))))
                        loopBody.Insert(convWasm_ls_3_(_env, mess, WasmBin_nop(_env)))
                        loopBody.Insert(convWasm_ls_3_(_env, val, WasmBin_nop(_env)))
                        loopBody.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, val)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](anyNull, WasmBin_call(_env, self.rt.FP.FnToStr(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, key)))))))))))
                    }
                    loopBody.Insert(WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, result), convWasm_lg_2_(_env, key), convWasm_lg_2_(_env, val)))))
                    loopBody.Insert(convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))))
                    loopBody.Insert(WasmBin_br(_env, 0))
                    body.Insert(WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, loopBody, nil))), nil))
                    body.Insert(convWasm_ok(_env, convWasm_lg_2_(_env, result)))
                } else if _switch0 == Ast_TypeInfoKind__Alge {
                    var algeType *Ast_TypeInfo
                    algeType = self.FP.getGenSrc(_env, work)
                    var algeInfo *convWasm_AlgeInfo
                    algeInfo = Lns_unwrap( self.FP.getAlgeInfo(_env, algeType)).(*convWasm_AlgeInfo)
                    var algeWork *Ast_AlgeTypeInfo
                    algeWork = Lns_unwrap( Ast_AlgeTypeInfoDownCastF(algeType.FP)).(*Ast_AlgeTypeInfo)
                    var algeName LnsInt
                    algeName = _func.FP.AddLocal(_env, self.rt.FP.Get_strType(_env))
                    var params LnsInt
                    params = _func.FP.AddLocal(_env, WasmBin_typeAnyRef)
                    var val LnsInt
                    val = _func.FP.AddLocal(_env, WasmBin_typeAnyRef)
                    var convWasm_item func(_env *LnsEnv, obj *WasmBin_Ins,itemIndex LnsInt) *WasmBin_Ins
                    convWasm_item = func(_env *LnsEnv, obj *WasmBin_Ins,itemIndex LnsInt) *WasmBin_Ins {
                        return WasmBin_call(_env, self.FP.fnStemIndex(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](obj, self.rt.FP.BoxI64(_env, WasmBin_i64Const(_env, itemIndex)))))
                    }
                    body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, WasmBin_opI32Or, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refTest(_env, self.rt.FP.Get_tList(_env), convWasm_lg_2_(_env, 0)), WasmBin_refTest(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0)))))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](fail))))
                    body.Insert(convWasm_ls_3_(_env, val, convWasm_item(_env, convWasm_lg_2_(_env, 0), 1)))
                    body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refTest(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, val))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](fail))))
                    body.Insert(convWasm_ls_3_(_env, algeName, WasmBin_refCast(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, val))))
                    body.Insert(convWasm_ls_3_(_env, params, convWasm_item(_env, convWasm_lg_2_(_env, 0), 2)))
                    for _, _valName := range( algeInfo.FP.Get_valNames(_env).Items ) {
                        valName := _valName
                        var valInfo *Ast_AlgeValInfo
                        valInfo = Lns_unwrap( algeWork.FP.Get_valInfoMap(_env).Get(valName)).(*Ast_AlgeValInfo)
                        var valType LnsInt
                        valType = Lns_unwrap( algeInfo.FP.Get_name2type(_env).Get(valName)).(LnsInt)
                        var thenBody *LnsList2_[*WasmBin_Ins]
                        thenBody = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
                        if valInfo.FP.Get_typeList(_env).Len() == 0{
                            thenBody.Insert(convWasm_ok(_env, WasmBin_globalGet(_env, Lns_unwrap( algeInfo.FP.Get_name2global(_env).Get(valName)).(LnsInt))))
                        } else { 
                            thenBody.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, params)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](fail))))
                            var fields *LnsList2_[*WasmBin_Ins]
                            fields = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), WasmBin_i32Const(_env, Lns_unwrap( algeInfo.FP.Get_name2tag(_env).Get(valName)).(LnsInt))))
                            for _paramIndex, _paramType := range( valInfo.FP.Get_typeList(_env).Items ) {
                                paramIndex := _paramIndex + 1
                                paramType := _paramType
                                var paramLocal LnsInt
                                paramLocal = _func.FP.AddLocal(_env, self.FP.getRep(_env, paramType))
                                thenBody.Insert(WasmBin_call(_env, self.FP.fnFromConv(_env, node, paramType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_item(_env, convWasm_lg_2_(_env, params), paramIndex)))))
                                thenBody.Insert(WasmBin_drop(_env, WasmBin_nop(_env)))
                                thenBody.Insert(convWasm_ls_3_(_env, val, WasmBin_nop(_env)))
                                if Lns_op_not(paramType.FP.Get_nilable(_env)){
                                    thenBody.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, val)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](fail))))
                                }
                                thenBody.Insert(convWasm_ls_3_(_env, paramLocal, self.FP.coerceRep(_env, convWasm_lg_2_(_env, val), WasmBin_typeAnyRef, self.FP.getRep(_env, paramType))))
                                fields.Insert(convWasm_lg_2_(_env, paramLocal))
                            }
                            thenBody.Insert(convWasm_ok(_env, WasmBin_structNew(_env, valType, fields)))
                        }
                        body.Insert(convWasm_convFilter_fnFromConv_ok__ifThen_1_(_env, WasmBin_call(_env, self.rt.FP.FnStrEq(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, algeName), self.rt.FP.StrLiteral(_env, valName)))), thenBody))
                    }
                    body.Insert(fail)
                } else {
                    var info *convWasm_ClassInfo
                    
                    {
                        _info := self.FP.getClassInfo(_env, work)
                        if _info == nil{
                            self.FP.error(_env, node, _env.GetVM().String_format("not support _fromMap -- %s", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil))))
                        } else {
                            info = _info.(*convWasm_ClassInfo)
                        }
                    }
                    body.Insert(WasmBin_call(_env, self.FP.fnFromMap(_env, node, info), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0)))))
                }
            }
        }
    }
    _func.FP.SetBody(_env, body)
    return index
}
// 4579: decl @lune.@base.@convWasm.convFilter.fnFromMap
func (self *convWasm_convFilter) fnFromMap(_env *LnsEnv, node *Nodes_Node,info *convWasm_ClassInfo) LnsInt {
    var name string
    name = _env.GetVM().String_format("fromMap:%d", Lns_2DDD(info.FP.Get_structType(_env)))
    {
        _index := self.adapterMap.Get(name)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var objRep *WasmBin_ValType
    objRep = WasmBin_refType(_env, info.FP.Get_structType(_env))
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](objRep, self.rt.FP.Get_strType(_env))))
    self.adapterMap.Set(name,index)
    self.mod.FP.SetFuncName(_env, index, name)
    var _map LnsInt
    _map = _func.FP.AddLocal(_env, self.rt.FP.Get_mapType(_env))
    var val LnsInt
    val = _func.FP.AddLocal(_env, WasmBin_typeAnyRef)
    var mess LnsInt
    mess = _func.FP.AddLocal(_env, self.rt.FP.Get_strType(_env))
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refTest(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, info.FP.Get_structType(_env)), WasmBin_refNull(_env, self.rt.FP.Get_tStr(_env))))))), nil, nil), convWasm_ls_3_(_env, _map, WasmBin_refCast(_env, self.rt.FP.Get_tMap(_env), convWasm_lg_2_(_env, 0)))))
    var fields *LnsList2_[*WasmBin_Ins]
    fields = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), WasmBin_globalGet(_env, info.FP.Get_vtGlobal(_env))))
    for _fieldIndex, _fieldName := range( info.FP.Get_fieldNames(_env).Items ) {
        fieldIndex := _fieldIndex + 1
        fieldName := _fieldName
        var rep *WasmBin_ValType
        rep = info.FP.Get_fieldReps(_env).GetAt(fieldIndex)
        var fieldType *Ast_TypeInfo
        fieldType = info.FP.Get_fieldTypes(_env).GetAt(fieldIndex)
        var fieldLocal LnsInt
        fieldLocal = _func.FP.AddLocal(_env, rep)
        body.Insert(WasmBin_call(_env, self.FP.fnFromConv(_env, node, fieldType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnMapGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, _map), self.rt.FP.StrLiteral(_env, fieldName))))))))
        body.Insert(convWasm_ls_3_(_env, mess, WasmBin_nop(_env)))
        body.Insert(convWasm_ls_3_(_env, val, WasmBin_nop(_env)))
        if Lns_op_not(fieldType.FP.Get_nilable(_env)){
            var nameTxt *WasmBin_Ins
            nameTxt = self.rt.FP.StrLiteral(_env, fieldName)
            body.Insert(WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, val)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, info.FP.Get_structType(_env)), WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, mess)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](nameTxt)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnStrConcat(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.StrLiteral(_env, fieldName + "."), convWasm_lg_2_(_env, mess)))))), self.rt.FP.Get_strType(_env))))))), nil, nil))
        }
        body.Insert(convWasm_ls_3_(_env, fieldLocal, self.FP.coerceRep(_env, convWasm_lg_2_(_env, val), WasmBin_typeAnyRef, rep)))
        fields.Insert(convWasm_lg_2_(_env, fieldLocal))
    }
    body.Insert(WasmBin_structNew(_env, info.FP.Get_structType(_env), fields))
    body.Insert(WasmBin_refNull(_env, self.rt.FP.Get_tStr(_env)))
    _func.FP.SetBody(_env, body)
    return index
}
// 4627: decl @lune.@base.@convWasm.convFilter.lowerStmtList
func (self *convWasm_convFilter) lowerStmtList(_env *LnsEnv, stmtList *LnsList2_[*Nodes_Node]) *LnsList2_[*WasmBin_Ins] {
    var list *LnsList2_[*WasmBin_Ins]
    list = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt
        self.FP.preDeclare(_env, stmt, list)
    }
    for _, _stmt := range( stmtList.Items ) {
        stmt := _stmt
        self.FP.lowerStmt(_env, stmt, list)
    }
    return list
}
// 4639: decl @lune.@base.@convWasm.convFilter.lowerBlock
func (self *convWasm_convFilter) lowerBlock(_env *LnsEnv, node *Nodes_BlockNode) *LnsList2_[*WasmBin_Ins] {
    return self.FP.lowerStmtList(_env, node.FP.Get_stmtList(_env))
}
// 4644: decl @lune.@base.@convWasm.convFilter.lowerBlockAt
func (self *convWasm_convFilter) lowerBlockAt(_env *LnsEnv, node *Nodes_BlockNode,depth LnsInt) *LnsList2_[*WasmBin_Ins] {
    var saved LnsInt
    saved = self.ctx.FP.Get_depth(_env)
    self.ctx.FP.Set_depth(_env, depth)
    var list *LnsList2_[*WasmBin_Ins]
    list = self.FP.lowerBlock(_env, node)
    self.ctx.FP.Set_depth(_env, saved)
    return list
}
// 4653: decl @lune.@base.@convWasm.convFilter.preDeclare
func (self *convWasm_convFilter) preDeclare(_env *LnsEnv, stmt *Nodes_Node,out *LnsList2_[*WasmBin_Ins]) {
    {
        _macroNode := convWasm_getMacroExp_21_(_env, stmt)
        if !Lns_IsNil( _macroNode ) {
            macroNode := _macroNode.(*Nodes_ExpMacroExpNode)
            for _, _child := range( macroNode.FP.Get_stmtList(_env).Items ) {
                child := _child
                self.FP.preDeclare(_env, child, out)
            }
            return 
        }
    }
    {
        _declNode := Nodes_DeclFuncNodeDownCastF(stmt.FP)
        if !Lns_IsNil( _declNode ) {
            declNode := _declNode.(*Nodes_DeclFuncNode)
            {
                _symbol := declNode.FP.Get_declInfo(_env).FP.Get_symbol(_env)
                if !Lns_IsNil( _symbol ) {
                    symbol := _symbol.(*Ast_SymbolInfo)
                    self.FP.registerFunc(_env, declNode.FP.Get_declInfo(_env), declNode.FP.Get_expType(_env))
                    var orgSym *Ast_SymbolInfo
                    orgSym = symbol.FP.GetOrg(_env)
                    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                        _env.SetStackVal( Lns_op_not(self.FP.isGlobalSym(_env, orgSym))) &&
                        _env.SetStackVal( Lns_op_not(self.FP.findLoc(_env, orgSym))) ).(bool)){
                        out.Insert(self.FP.declareVar(_env, orgSym, nil))
                    }
                }
            }
        }
    }
    {
        _classNode := Nodes_DeclClassNodeDownCastF(stmt.FP)
        if !Lns_IsNil( _classNode ) {
            classNode := _classNode.(*Nodes_DeclClassNode)
            if classNode.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Class{
                self.FP.registerClass(_env, classNode)
            }
        }
    }
}
// 4684: decl @lune.@base.@convWasm.convFilter.enterLoop
func (self *convWasm_convFilter) enterLoop(_env *LnsEnv) {
    self.ctx.FP.Set_depth(_env, self.ctx.FP.Get_depth(_env) + 1)
    self.ctx.FP.Get_breakDepthList(_env).Insert(LnsInt(self.ctx.FP.Get_depth(_env)))
    self.ctx.FP.Set_depth(_env, self.ctx.FP.Get_depth(_env) + 1)
}
// 4691: decl @lune.@base.@convWasm.convFilter.exitLoop
func (self *convWasm_convFilter) exitLoop(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    self.ctx.FP.Set_depth(_env, self.ctx.FP.Get_depth(_env) - 2)
    self.ctx.FP.Get_breakDepthList(_env).Remove(nil)
    return WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, body, nil))), nil)
}
// 4704: decl @lune.@base.@convWasm.convFilter.lowerApply
func (self *convWasm_convFilter) lowerApply(_env *LnsEnv, node *Nodes_ApplyNode) *WasmBin_Ins {
    var expList *Nodes_ExpListNode
    expList = node.FP.Get_expList(_env)
    var builtin *Builtin_BuiltinFuncType
    builtin = self.builtinFuncs
    var list *LnsList2_[*WasmBin_Ins]
    list = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    if expList.FP.Get_expList(_env).Len() >= 1{
        var callNode LnsAny
        callNode = nil
        {
            _mretNode := Nodes_ExpMRetNodeDownCastF(expList.FP.Get_expList(_env).GetAt(1).FP)
            if !Lns_IsNil( _mretNode ) {
                mretNode := _mretNode.(*Nodes_ExpMRetNode)
                callNode = Nodes_ExpCallNodeDownCastF(mretNode.FP.Get_mRet(_env).FP)
            }
        }
        if callNode != nil{
            callNode_343 := callNode.(*Nodes_ExpCallNode)
            var funcType *Ast_TypeInfo
            funcType = callNode_343.FP.Get_func(_env).FP.Get_expType(_env)
            var headArgs *LnsList2_[*WasmBin_Ins]
            headArgs = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
            var isGmatch bool
            isGmatch = false
            if funcType.FP.Equals(_env, self.processInfo, builtin.Str_gmatch, nil, nil){
                {
                    _fieldNode := Nodes_RefFieldNodeDownCastF(callNode_343.FP.Get_func(_env).FP)
                    if !Lns_IsNil( _fieldNode ) {
                        fieldNode := _fieldNode.(*Nodes_RefFieldNode)
                        headArgs.Insert(self.FP.lowerExpAs(_env, fieldNode.FP.Get_prefix(_env), Ast_builtinTypeString))
                        isGmatch = true
                    }
                }
            } else if funcType.FP.Equals(_env, self.processInfo, builtin.String_gmatch, nil, nil){
                isGmatch = true
            }
            if isGmatch{
                return self.FP.lowerApplyGmatch(_env, node, callNode_343, headArgs)
            }
        }
    }
    var vals *LnsList2_[*convWasm_Val]
    vals = self.FP.lowerValList(_env, expList, list)
    var funcType *Ast_TypeInfo
    funcType = self.FP.normType(_env, expList.FP.Get_expTypeList(_env).GetAt(1))
    var cloInfo *WasmRt_CloInfo
    cloInfo = self.FP.getFuncCloInfo(_env, funcType)
    var clo LnsInt
    clo = self.FP.tempLocal(_env, WasmBin_refType(_env, cloInfo.FP.Get_cloType(_env)))
    var state LnsInt
    state = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
    var ctrl LnsInt
    ctrl = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
    list.Insert(convWasm_ls_3_(_env, clo, WasmBin_refCast(_env, cloInfo.FP.Get_cloType(_env), vals.GetAt(1).FP.Get_ins(_env))))
    var convWasm_getAny func(_env *LnsEnv, index LnsInt) *WasmBin_Ins
    convWasm_getAny = func(_env *LnsEnv, index LnsInt) *WasmBin_Ins {
        if index <= vals.Len(){
            return self.FP.coerceVal(_env, vals.GetAt(index), Ast_builtinTypeStem_)
        }
        return WasmBin_refNull(_env, WasmBin_htAny)
    }
    list.Insert(convWasm_ls_3_(_env, state, convWasm_getAny(_env, 2)))
    list.Insert(convWasm_ls_3_(_env, ctrl, convWasm_getAny(_env, 3)))
    var argTypes *LnsList2_[*Ast_TypeInfo]
    argTypes = funcType.FP.Get_argTypeInfoList(_env)
    var retTypes *LnsList2_[*Ast_TypeInfo]
    retTypes = funcType.FP.Get_retTypeInfoList(_env)
    self.FP.enterLoop(_env)
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, clo)))
    for _index, _argType := range( argTypes.Items ) {
        index := _index + 1
        argType := _argType
        var argIns *WasmBin_Ins
        if index == 1{
            argIns = convWasm_lg_2_(_env, state)
        } else if index == 2{
            argIns = convWasm_lg_2_(_env, ctrl)
        } else { 
            argIns = WasmBin_refNull(_env, WasmBin_htAny)
        }
        args.Insert(self.FP.coerce(_env, argIns, Ast_builtinTypeStem_, argType))
    }
    args.Insert(WasmBin_structGet(_env, cloInfo.FP.Get_cloType(_env), WasmRt_fCloCode, convWasm_lg_2_(_env, clo)))
    body.Insert(WasmBin_callRef(_env, cloInfo.FP.Get_funcType(_env), args))
    var locals *LnsList2_[LnsInt]
    locals = NewLnsList2_[LnsInt]([]LnsInt{})
    for _, _retType := range( retTypes.Items ) {
        retType := _retType
        locals.Insert(LnsInt(self.FP.tempLocal(_env, self.FP.getRep(_env, retType))))
    }
    {
        var _forFrom0 LnsInt = locals.Len()
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
            body.Insert(convWasm_ls_3_(_env, locals.GetAt(index), WasmBin_nop(_env)))
            _forWork0 += _forDelta0
        }
    }
    body.Insert(convWasm_ls_3_(_env, ctrl, self.FP.coerce(_env, convWasm_lg_2_(_env, locals.GetAt(1)), retTypes.GetAt(1), Ast_builtinTypeStem_)))
    body.Insert(WasmBin_brIf(_env, 1, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, ctrl))))
    for _index, _varSym := range( node.FP.Get_varList(_env).Items ) {
        index := _index + 1
        varSym := _varSym
        var valIns *WasmBin_Ins
        if index <= retTypes.Len(){
            valIns = self.FP.coerce(_env, convWasm_lg_2_(_env, locals.GetAt(index)), retTypes.GetAt(index), varSym.FP.Get_typeInfo(_env))
        } else { 
            valIns = self.FP.getDefaultVal(_env, self.FP.getRep(_env, varSym.FP.Get_typeInfo(_env)))
        }
        body.Insert(self.FP.declareVar(_env, varSym, valIns))
    }
    convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, node.FP.Get_block(_env)))
    body.Insert(WasmBin_br(_env, 0))
    list.Insert(self.FP.exitLoop(_env, body))
    return WasmBin_seq(_env, list)
}
// 4791: decl @lune.@base.@convWasm.convFilter.lowerApplyGmatch
func (self *convWasm_convFilter) lowerApplyGmatch(_env *LnsEnv, node *Nodes_ApplyNode,callNode *Nodes_ExpCallNode,headArgs *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    var list *LnsList2_[*WasmBin_Ins]
    list = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var ins *WasmBin_Ins
    ins,_ = self.FP.lowerLuaCall(_env, callNode, WasmRt_luaGmatchAll, headArgs)
    var result LnsInt
    result = self.FP.tempLocal(_env, self.rt.FP.Get_listType(_env))
    list.Insert(convWasm_ls_3_(_env, result, ins))
    var num LnsInt
    num = self.FP.tempLocal(_env, WasmBin_typeI32)
    list.Insert(convWasm_ls_3_(_env, num, convWasm_p_4_(_env, WasmBin_opI32WrapI64, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.UnboxI64(_env, self.rt.FP.ListItem(_env, convWasm_lg_2_(_env, result), 0)))))))
    var pos LnsInt
    pos = self.FP.tempLocal(_env, WasmBin_typeI32)
    list.Insert(convWasm_ls_3_(_env, pos, WasmBin_i32Const(_env, 1)))
    self.FP.enterLoop(_env)
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    body.Insert(WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, convWasm_lg_2_(_env, result)))))))
    for _index, _varSym := range( node.FP.Get_varList(_env).Items ) {
        index := _index + 1
        varSym := _varSym
        var item *WasmBin_Ins
        item = WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListItems, convWasm_lg_2_(_env, result)), convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, index - 1)))))
        var valIns *WasmBin_Ins
        valIns = WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, index - 1), convWasm_lg_2_(_env, num)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](item)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, WasmBin_htAny))), WasmBin_typeAnyRef)
        body.Insert(self.FP.declareVar(_env, varSym, self.FP.coerce(_env, valIns, Ast_builtinTypeStem_, varSym.FP.Get_typeInfo(_env))))
    }
    body.Insert(convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), convWasm_lg_2_(_env, num))))))
    convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, node.FP.Get_block(_env)))
    body.Insert(WasmBin_br(_env, 0))
    list.Insert(self.FP.exitLoop(_env, body))
    return WasmBin_seq(_env, list)
}
// 4843: decl @lune.@base.@convWasm.convFilter.lowerIf
func (self *convWasm_convFilter) lowerIf(_env *LnsEnv, node *Nodes_IfNode) *WasmBin_Ins {
    var base LnsInt
    base = self.ctx.FP.Get_depth(_env)
    var conds *LnsList2_[*WasmBin_Ins]
    conds = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var bodies *LnsList2_[*LnsList2_[*WasmBin_Ins]]
    bodies = NewLnsList2_[*LnsList2_[*WasmBin_Ins]]([]*LnsList2_[*WasmBin_Ins]{})
    var elseBody LnsAny
    elseBody = nil
    for _, _stmt := range( node.FP.Get_stmtList(_env).Items ) {
        stmt := _stmt
        if stmt.FP.Get_kind(_env) == Nodes_IfKind__Else{
            elseBody = self.FP.lowerBlockAt(_env, stmt.FP.Get_block(_env), base + conds.Len())
        } else { 
            self.ctx.FP.Set_depth(_env, base + conds.Len())
            conds.Insert(self.FP.lowerCond(_env, stmt.FP.Get_exp(_env)))
            bodies.Insert(self.FP.lowerBlockAt(_env, stmt.FP.Get_block(_env), base + conds.Len()))
        }
    }
    self.ctx.FP.Set_depth(_env, base)
    return convWasm_makeIfChain_18_(_env, conds, bodies, elseBody)
}
// 4861: decl @lune.@base.@convWasm.convFilter.lowerSwitch
func (self *convWasm_convFilter) lowerSwitch(_env *LnsEnv, node *Nodes_SwitchNode) *WasmBin_Ins {
    var base LnsInt
    base = self.ctx.FP.Get_depth(_env)
    var val *convWasm_Val
    val = self.FP.lowerVal(_env, node.FP.Get_exp(_env))
    var work LnsInt
    work = self.FP.tempLocal(_env, val.FP.Get_rep(_env))
    var localVal *convWasm_Val
    localVal = NewconvWasm_Val(_env, convWasm_lg_2_(_env, work), val.FP.Get_rep(_env), val.FP.Get_typeInfo(_env))
    var conds *LnsList2_[*WasmBin_Ins]
    conds = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var bodies *LnsList2_[*LnsList2_[*WasmBin_Ins]]
    bodies = NewLnsList2_[*LnsList2_[*WasmBin_Ins]]([]*LnsList2_[*WasmBin_Ins]{})
    for _, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
        caseInfo := _caseInfo
        var cond LnsAny
        cond = nil
        for _, _exp := range( caseInfo.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
            exp := _exp
            var eq *WasmBin_Ins
            eq = self.FP.lowerEqVals(_env, localVal, self.FP.lowerVal(_env, exp))
            {
                _prev := cond
                if !Lns_IsNil( _prev ) {
                    prev := _prev.(*WasmBin_Ins)
                    cond = convWasm_p_4_(_env, WasmBin_opI32Or, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](prev, eq)))
                } else {
                    cond = eq
                }
            }
        }
        conds.Insert(Lns_unwrapDefault( cond, WasmBin_i32Const(_env, 0)).(*WasmBin_Ins))
        bodies.Insert(self.FP.lowerBlockAt(_env, caseInfo.FP.Get_block(_env), base + conds.Len()))
    }
    var elseBody LnsAny
    elseBody = nil
    {
        _defBlock := node.FP.Get_default(_env)
        if !Lns_IsNil( _defBlock ) {
            defBlock := _defBlock.(*Nodes_BlockNode)
            elseBody = self.FP.lowerBlockAt(_env, defBlock, base + conds.Len())
        }
    }
    return WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, work, val.FP.Get_ins(_env)), convWasm_makeIfChain_18_(_env, conds, bodies, elseBody))))
}
// 4888: decl @lune.@base.@convWasm.convFilter.lowerMatch
func (self *convWasm_convFilter) lowerMatch(_env *LnsEnv, node *Nodes_MatchNode) *WasmBin_Ins {
    var base LnsInt
    base = self.ctx.FP.Get_depth(_env)
    var algeType *Ast_TypeInfo
    algeType = self.FP.getGenSrc(_env, node.FP.Get_val(_env).FP.Get_expType(_env))
    var info *convWasm_AlgeInfo
    
    {
        _info := self.FP.getAlgeInfo(_env, algeType)
        if _info == nil{
            self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not found alge -- %s", Lns_2DDD(algeType.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            info = _info.(*convWasm_AlgeInfo)
        }
    }
    var algeVal LnsInt
    algeVal = self.FP.tempLocal(_env, WasmBin_refType(_env, info.FP.Get_baseType(_env)))
    var tag LnsInt
    tag = self.FP.tempLocal(_env, WasmBin_typeI32)
    var conds *LnsList2_[*WasmBin_Ins]
    conds = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var bodies *LnsList2_[*LnsList2_[*WasmBin_Ins]]
    bodies = NewLnsList2_[*LnsList2_[*WasmBin_Ins]]([]*LnsList2_[*WasmBin_Ins]{})
    for _, _caseInfo := range( node.FP.Get_caseList(_env).Items ) {
        caseInfo := _caseInfo
        var name string
        name = caseInfo.FP.Get_valInfo(_env).FP.Get_name(_env)
        conds.Insert(convWasm_p_4_(_env, WasmBin_opI32Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, tag), WasmBin_i32Const(_env, Lns_unwrap( info.FP.Get_name2tag(_env).Get(name)).(LnsInt))))))
        self.ctx.FP.Set_depth(_env, base + conds.Len())
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        var valType LnsInt
        valType = Lns_unwrap( info.FP.Get_name2type(_env).Get(name)).(LnsInt)
        for _index, _symbol := range( caseInfo.FP.Get_valParamNameList(_env).Items ) {
            index := _index + 1
            symbol := _symbol
            var paramType *Ast_TypeInfo
            paramType = caseInfo.FP.Get_valInfo(_env).FP.Get_typeList(_env).GetAt(index)
            var paramVal *WasmBin_Ins
            paramVal = WasmBin_structGet(_env, valType, index + 1, WasmBin_refCast(_env, valType, convWasm_lg_2_(_env, algeVal)))
            body.Insert(self.FP.declareVar(_env, symbol, self.FP.coerce(_env, paramVal, paramType, symbol.FP.Get_typeInfo(_env))))
        }
        convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, caseInfo.FP.Get_block(_env)))
        bodies.Insert(body)
    }
    var elseBody LnsAny
    elseBody = nil
    {
        _defBlock := node.FP.Get_defaultBlock(_env)
        if !Lns_IsNil( _defBlock ) {
            defBlock := _defBlock.(*Nodes_Node)
            self.ctx.FP.Set_depth(_env, base + conds.Len())
            var defBody *LnsList2_[*WasmBin_Ins]
            defBody = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
            self.FP.lowerStmt(_env, defBlock, defBody)
            elseBody = defBody
        }
    }
    self.ctx.FP.Set_depth(_env, base)
    return WasmBin_seq(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, algeVal, self.FP.lowerExpAs(_env, node.FP.Get_val(_env), algeType)), convWasm_ls_3_(_env, tag, WasmBin_structGet(_env, info.FP.Get_baseType(_env), 1, convWasm_lg_2_(_env, algeVal))), convWasm_makeIfChain_18_(_env, conds, bodies, elseBody))))
}
// 4927: decl @lune.@base.@convWasm.convFilter.lowerRawVal
func (self *convWasm_convFilter) lowerRawVal(_env *LnsEnv, node *Nodes_Node) *convWasm_Val {
    {
        _castNode := Nodes_ExpCastNodeDownCastF(node.FP)
        if !Lns_IsNil( _castNode ) {
            castNode := _castNode.(*Nodes_ExpCastNode)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( castNode.FP.Get_castKind(_env) == Nodes_CastKind__Implicit) &&
                _env.SetStackVal( castNode.FP.Get_exp(_env).FP.Get_expType(_env).FP.Get_nilable(_env)) ).(bool)){
                return self.FP.lowerVal(_env, castNode.FP.Get_exp(_env))
            }
        }
    }
    return self.FP.lowerVal(_env, node)
}
// 4937: decl @lune.@base.@convWasm.convFilter.spillVal
func (self *convWasm_convFilter) spillVal(_env *LnsEnv, val *convWasm_Val,pre *LnsList2_[*WasmBin_Ins]) *convWasm_Val {
    var work LnsInt
    work = self.FP.tempLocal(_env, val.FP.Get_rep(_env))
    pre.Insert(convWasm_ls_3_(_env, work, val.FP.Get_ins(_env)))
    return NewconvWasm_Val(_env, convWasm_lg_2_(_env, work), val.FP.Get_rep(_env), val.FP.Get_typeInfo(_env))
}
// 4948: decl @lune.@base.@convWasm.convFilter.expandDDD
func (self *convWasm_convFilter) expandDDD(_env *LnsEnv, vals *LnsList2_[*convWasm_Val],num LnsInt,pre *LnsList2_[*WasmBin_Ins]) *LnsList2_[*convWasm_Val] {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( vals.Len() == 0) ||
        _env.SetStackVal( vals.Len() > num) ||
        _env.SetStackVal( self.FP.normType(_env, vals.GetAt(vals.Len()).FP.Get_typeInfo(_env)).FP.Get_kind(_env) != Ast_TypeInfoKind__DDD) ).(bool){
        return vals
    }
    var work *LnsList2_[*convWasm_Val]
    work = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = vals.Len() - 1
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            index := _forWork0
            work.Insert(vals.GetAt(index))
        }
    }
    var list *convWasm_Val
    list = self.FP.spillVal(_env, vals.GetAt(vals.Len()), pre)
    {
        var _forFrom1 LnsInt = 1
        var _forTo1 LnsInt = num - vals.Len() + 1
        for _forWork1 := _forFrom1; _forWork1 <= _forTo1; _forWork1++ {
            index := _forWork1
            work.Insert(self.FP.spillVal(_env, NewconvWasm_Val(_env, WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](list.FP.Get_ins(_env), WasmBin_i64Const(_env, index)))), WasmBin_typeAnyRef, Ast_builtinTypeStem_), pre))
        }
    }
    return work
}
// 4968: decl @lune.@base.@convWasm.convFilter.lowerRawValList
func (self *convWasm_convFilter) lowerRawValList(_env *LnsEnv, expList *Nodes_ExpListNode,pre *LnsList2_[*WasmBin_Ins]) *LnsList2_[*convWasm_Val] {
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    if Lns_isCondTrue( expList.FP.Get_mRetExp(_env)){
        for _, _val := range( self.FP.lowerValListSub(_env, expList, pre, true).Items ) {
            val := _val
            vals.Insert(self.FP.spillVal(_env, val, pre))
        }
    } else { 
        for _, _exp := range( expList.FP.Get_expList(_env).Items ) {
            exp := _exp
            vals.Insert(self.FP.spillVal(_env, self.FP.lowerRawVal(_env, exp), pre))
        }
    }
    return vals
}
// 4985: decl @lune.@base.@convWasm.convFilter.anyNil
func (self *convWasm_convFilter) anyNil(_env *LnsEnv, vals *LnsList2_[*convWasm_Val],num LnsInt) *WasmBin_Ins {
    var cond *WasmBin_Ins
    cond = WasmBin_i32Const(_env, 0)
    for _index, _val := range( vals.Items ) {
        index := _index + 1
        val := _val
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( val.FP.Get_rep(_env).FP.IsRef(_env)) &&
            _env.SetStackVal( index <= num) ).(bool)){
            cond = convWasm_p_4_(_env, WasmBin_opI32Or, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cond, WasmBin_refIsNull(_env, val.FP.Get_ins(_env)))))
        }
    }
    return cond
}
// 4995: decl @lune.@base.@convWasm.convFilter.lowerDeclVar
func (self *convWasm_convFilter) lowerDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,out *LnsList2_[*WasmBin_Ins]) {
    if node.FP.Get_unwrapFlag(_env){
        var base LnsInt
        base = self.ctx.FP.Get_depth(_env)
        var vals *LnsList2_[*convWasm_Val]
        vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
        {
            _expList := node.FP.Get_expList(_env)
            if !Lns_IsNil( _expList ) {
                expList := _expList.(*Nodes_ExpListNode)
                vals = self.FP.expandDDD(_env, self.FP.lowerRawValList(_env, expList, out), node.FP.Get_symbolInfoList(_env).Len(), out)
            }
        }
        if node.FP.Get_mode(_env) == Nodes_DeclVarMode__Let{
            for _, _symbol := range( node.FP.Get_symbolInfoList(_env).Items ) {
                symbol := _symbol
                out.Insert(self.FP.declareVar(_env, symbol, nil))
            }
        }
        self.ctx.FP.Set_depth(_env, base + 1)
        var nilBody *LnsList2_[*WasmBin_Ins]
        nilBody = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        {
            _block := node.FP.Get_unwrapBlock(_env)
            if !Lns_IsNil( _block ) {
                block := _block.(*Nodes_BlockNode)
                for _index, _symbol := range( node.FP.Get_symbolInfoList(_env).Items ) {
                    index := _index + 1
                    symbol := _symbol
                    {
                        _rawSym := block.FP.Get_scope(_env).FP.Get_symbol2SymbolInfoMap(_env).Get("_" + symbol.FP.Get_name(_env))
                        if !Lns_IsNil( _rawSym ) {
                            rawSym := _rawSym.(*Ast_SymbolInfo)
                            if index <= vals.Len(){
                                nilBody.Insert(self.FP.declareVar(_env, rawSym, self.FP.coerceVal(_env, vals.GetAt(index), rawSym.FP.Get_typeInfo(_env))))
                            }
                        }
                    }
                }
                convWasm_append_17_(_env, nilBody, self.FP.lowerBlock(_env, block))
            }
        }
        var thenBody *LnsList2_[*WasmBin_Ins]
        thenBody = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        for _index, _symbol := range( node.FP.Get_symbolInfoList(_env).Items ) {
            index := _index + 1
            symbol := _symbol
            if index <= vals.Len(){
                thenBody.Insert(self.FP.writeSym(_env, &node.Nodes_Node, symbol, self.FP.coerceVal(_env, vals.GetAt(index), symbol.FP.Get_typeInfo(_env))))
            }
        }
        {
            _block := node.FP.Get_thenBlock(_env)
            if !Lns_IsNil( _block ) {
                block := _block.(*Nodes_BlockNode)
                convWasm_append_17_(_env, thenBody, self.FP.lowerBlock(_env, block))
            }
        }
        self.ctx.FP.Set_depth(_env, base)
        out.Insert(WasmBin_ifElse(_env, self.FP.anyNil(_env, vals, node.FP.Get_symbolInfoList(_env).Len()), nilBody, thenBody, nil))
        return 
    }
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            vals = self.FP.expandDDD(_env, self.FP.lowerValList(_env, expList, out), node.FP.Get_symbolInfoList(_env).Len(), out)
        }
    }
    for _index, _symbol := range( node.FP.Get_symbolInfoList(_env).Items ) {
        index := _index + 1
        symbol := _symbol
        if index <= vals.Len(){
            out.Insert(self.FP.declareVar(_env, symbol, self.FP.coerceVal(_env, vals.GetAt(index), symbol.FP.Get_typeInfo(_env))))
        } else { 
            out.Insert(self.FP.declareVar(_env, symbol, nil))
        }
    }
}
// 5051: decl @lune.@base.@convWasm.convFilter.lowerIfUnwrap
func (self *convWasm_convFilter) lowerIfUnwrap(_env *LnsEnv, node *Nodes_IfUnwrapNode) *WasmBin_Ins {
    var base LnsInt
    base = self.ctx.FP.Get_depth(_env)
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var vals *LnsList2_[*convWasm_Val]
    vals = self.FP.expandDDD(_env, self.FP.lowerRawValList(_env, node.FP.Get_expList(_env), pre), node.FP.Get_varSymList(_env).Len(), pre)
    self.ctx.FP.Set_depth(_env, base + 1)
    var thenBody *LnsList2_[*WasmBin_Ins]
    thenBody = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _index, _symbol := range( node.FP.Get_varSymList(_env).Items ) {
        index := _index + 1
        symbol := _symbol
        if index <= vals.Len(){
            thenBody.Insert(self.FP.declareVar(_env, symbol, self.FP.coerceVal(_env, vals.GetAt(index), symbol.FP.Get_typeInfo(_env))))
        }
    }
    convWasm_append_17_(_env, thenBody, self.FP.lowerBlock(_env, node.FP.Get_block(_env)))
    var elseBody LnsAny
    elseBody = nil
    {
        _block := node.FP.Get_nilBlock(_env)
        if !Lns_IsNil( _block ) {
            block := _block.(*Nodes_BlockNode)
            elseBody = self.FP.lowerBlock(_env, block)
        }
    }
    self.ctx.FP.Set_depth(_env, base)
    var checkNum LnsInt
    checkNum = node.FP.Get_varSymList(_env).Len()
    if checkNum == 0{
        checkNum = vals.Len()
    }
    pre.Insert(WasmBin_ifElse(_env, self.FP.anyNil(_env, vals, checkNum), Lns_unwrapDefault( elseBody, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})).(*LnsList2_[*WasmBin_Ins]), thenBody, nil))
    return WasmBin_seq(_env, pre)
}
// 5079: decl @lune.@base.@convWasm.convFilter.lowerWhen
func (self *convWasm_convFilter) lowerWhen(_env *LnsEnv, node *Nodes_WhenNode) *WasmBin_Ins {
    var base LnsInt
    base = self.ctx.FP.Get_depth(_env)
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    for _, _pair := range( node.FP.Get_symPairList(_env).Items ) {
        pair := _pair
        var src *Ast_SymbolInfo
        src = pair.FP.Get_src(_env).FP.GetOrg(_env)
        vals.Insert(NewconvWasm_Val(_env, self.FP.readSym(_env, &node.Nodes_Node, src), self.FP.getRep(_env, src.FP.Get_typeInfo(_env)), src.FP.Get_typeInfo(_env)))
    }
    self.ctx.FP.Set_depth(_env, base + 1)
    var thenBody *LnsList2_[*WasmBin_Ins]
    thenBody = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _index, _pair := range( node.FP.Get_symPairList(_env).Items ) {
        index := _index + 1
        pair := _pair
        thenBody.Insert(self.FP.declareVar(_env, pair.FP.Get_dst(_env), self.FP.coerceVal(_env, vals.GetAt(index), pair.FP.Get_dst(_env).FP.Get_typeInfo(_env))))
    }
    convWasm_append_17_(_env, thenBody, self.FP.lowerBlock(_env, node.FP.Get_block(_env)))
    var elseBody LnsAny
    elseBody = nil
    {
        _block := node.FP.Get_elseBlock(_env)
        if !Lns_IsNil( _block ) {
            block := _block.(*Nodes_BlockNode)
            elseBody = self.FP.lowerBlock(_env, block)
        }
    }
    self.ctx.FP.Set_depth(_env, base)
    return WasmBin_ifElse(_env, self.FP.anyNil(_env, vals, vals.Len()), Lns_unwrapDefault( elseBody, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})).(*LnsList2_[*WasmBin_Ins]), thenBody, nil)
}
// 5102: decl @lune.@base.@convWasm.convFilter.lowerUnwrapSet
func (self *convWasm_convFilter) lowerUnwrapSet(_env *LnsEnv, node *Nodes_UnwrapSetNode,out *LnsList2_[*WasmBin_Ins]) {
    var base LnsInt
    base = self.ctx.FP.Get_depth(_env)
    var vals *LnsList2_[*convWasm_Val]
    vals = self.FP.expandDDD(_env, self.FP.lowerRawValList(_env, node.FP.Get_srcExpList(_env), out), node.FP.Get_dstExpList(_env).FP.Get_expList(_env).Len(), out)
    self.ctx.FP.Set_depth(_env, base + 1)
    var nilBody *LnsList2_[*WasmBin_Ins]
    nilBody = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    {
        _block := node.FP.Get_unwrapBlock(_env)
        if !Lns_IsNil( _block ) {
            block := _block.(*Nodes_BlockNode)
            convWasm_append_17_(_env, nilBody, self.FP.lowerBlock(_env, block))
        }
    }
    var thenBody *LnsList2_[*WasmBin_Ins]
    thenBody = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _index, _dst := range( node.FP.Get_dstExpList(_env).FP.Get_expList(_env).Items ) {
        index := _index + 1
        dst := _dst
        if index <= vals.Len(){
            self.FP.lowerAssign(_env, dst, vals.GetAt(index), thenBody)
        }
    }
    self.ctx.FP.Set_depth(_env, base)
    out.Insert(WasmBin_ifElse(_env, self.FP.anyNil(_env, vals, node.FP.Get_dstExpList(_env).FP.Get_expList(_env).Len()), nilBody, thenBody, nil))
}
// 5122: decl @lune.@base.@convWasm.convFilter.lowerAssign
func (self *convWasm_convFilter) lowerAssign(_env *LnsEnv, dst *Nodes_Node,val *convWasm_Val,out *LnsList2_[*WasmBin_Ins]) {
    {
        _refNode := Nodes_ExpRefNodeDownCastF(dst.FP)
        if !Lns_IsNil( _refNode ) {
            refNode := _refNode.(*Nodes_ExpRefNode)
            var symbol *Ast_SymbolInfo
            symbol = refNode.FP.Get_symbolInfo(_env).FP.GetOrg(_env)
            out.Insert(self.FP.writeSym(_env, dst, symbol, self.FP.coerceVal(_env, val, symbol.FP.Get_typeInfo(_env))))
            return 
        }
    }
    {
        _fieldNode := Nodes_RefFieldNodeDownCastF(dst.FP)
        if !Lns_IsNil( _fieldNode ) {
            fieldNode := _fieldNode.(*Nodes_RefFieldNode)
            var prefixType *Ast_TypeInfo
            prefixType = self.FP.normType(_env, fieldNode.FP.Get_prefix(_env).FP.Get_expType(_env))
            var name string
            name = fieldNode.FP.Get_field(_env).Txt
            {
                _symbol := fieldNode.FP.Get_symbolInfo(_env)
                if !Lns_IsNil( _symbol ) {
                    symbol := _symbol.(*Ast_SymbolInfo)
                    var orgSym *Ast_SymbolInfo
                    orgSym = symbol.FP.GetOrg(_env)
                    if Lns_isCondTrue( self.globalSym2loc.Get(orgSym)){
                        out.Insert(self.FP.writeSym(_env, dst, orgSym, self.FP.coerceVal(_env, val, orgSym.FP.Get_typeInfo(_env))))
                        return 
                    }
                }
            }
            if prefixType.FP.Get_kind(_env) == Ast_TypeInfoKind__Map{
                out.Insert(WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, fieldNode.FP.Get_prefix(_env)), self.rt.FP.StrLiteral(_env, name), self.FP.toAny(_env, val)))))
                return 
            }
            var objType *Ast_TypeInfo
            objType = self.FP.getObjType(_env, prefixType)
            var info *convWasm_ClassInfo
            
            {
                _info := self.FP.getClassInfo(_env, objType)
                if _info == nil{
                    self.FP.error(_env, dst, _env.GetVM().String_format("not found class -- %s", Lns_2DDD(objType.FP.GetTxt(_env, nil, nil, nil))))
                } else {
                    info = _info.(*convWasm_ClassInfo)
                }
            }
            var fieldIndex LnsInt
            fieldIndex = self.FP.getFieldIndex(_env, dst, info, name)
            out.Insert(WasmBin_structSet(_env, info.FP.Get_structType(_env), fieldIndex, WasmBin_refCast(_env, info.FP.Get_structType(_env), self.FP.lowerExpAs(_env, fieldNode.FP.Get_prefix(_env), info.FP.Get_typeInfo(_env))), self.FP.coerce(_env, val.FP.Get_ins(_env), val.FP.Get_typeInfo(_env), self.FP.getMemberType(_env, info, name))))
            return 
        }
    }
    {
        _itemNode := Nodes_ExpRefItemNodeDownCastF(dst.FP)
        if !Lns_IsNil( _itemNode ) {
            itemNode := _itemNode.(*Nodes_ExpRefItemNode)
            var index *convWasm_Val
            {
                _indexNode := itemNode.FP.Get_index(_env)
                if !Lns_IsNil( _indexNode ) {
                    indexNode := _indexNode.(*Nodes_Node)
                    index = self.FP.lowerVal(_env, indexNode)
                } else {
                    index = NewconvWasm_Val(_env, self.rt.FP.StrLiteral(_env, Lns_unwrap( itemNode.FP.Get_symbol(_env)).(string)), self.rt.FP.Get_strType(_env), Ast_builtinTypeString)
                }
            }
            self.FP.assignItem(_env, dst, itemNode.FP.Get_val(_env), index, val, out)
            return 
        }
    }
    self.FP.error(_env, dst, _env.GetVM().String_format("not support assign -- %s", Lns_2DDD(Nodes_getNodeKindName(_env, dst.FP.Get_kind(_env)))))
}
// 5170: decl @lune.@base.@convWasm.convFilter.assignItem
func (self *convWasm_convFilter) assignItem(_env *LnsEnv, node *Nodes_Node,collNode *Nodes_Node,index *convWasm_Val,val *convWasm_Val,out *LnsList2_[*WasmBin_Ins]) {
    var collType *Ast_TypeInfo
    collType = self.FP.normType(_env, collNode.FP.Get_expType(_env))
    var coll *WasmBin_Ins
    coll = self.FP.lowerExp(_env, collNode)
    if _switch0 := collType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        out.Insert(WasmBin_call(_env, self.rt.FP.FnListSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](coll, self.FP.coerceVal(_env, index, Ast_builtinTypeInt), self.FP.toAny(_env, val)))))
        return 
    } else if _switch0 == Ast_TypeInfoKind__Map {
        out.Insert(WasmBin_call(_env, self.rt.FP.FnMapSet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](coll, self.FP.toAny(_env, index), self.FP.toAny(_env, val)))))
        return 
    }
    self.FP.error(_env, node, _env.GetVM().String_format("not support assign -- %s", Lns_2DDD(collType.FP.GetTxt(_env, nil, nil, nil))))
}
// 5191: decl @lune.@base.@convWasm.convFilter.lowerSetItem
func (self *convWasm_convFilter) lowerSetItem(_env *LnsEnv, node *Nodes_ExpSetItemNode,out *LnsList2_[*WasmBin_Ins]) {
    var index *convWasm_Val
    switch _matchExp0 := node.FP.Get_index(_env).(type) {
    case *Nodes_IndexVal__NodeIdx:
        indexNode := _matchExp0.Val1
        index = self.FP.lowerVal(_env, indexNode)
    case *Nodes_IndexVal__SymIdx:
        symTxt := _matchExp0.Val1
        index = NewconvWasm_Val(_env, self.rt.FP.StrLiteral(_env, symTxt), self.rt.FP.Get_strType(_env), Ast_builtinTypeString)
    }
    var val *convWasm_Val
    {
        _listNode := Nodes_ExpListNodeDownCastF(node.FP.Get_exp2(_env).FP)
        if !Lns_IsNil( _listNode ) {
            listNode := _listNode.(*Nodes_ExpListNode)
            var vals *LnsList2_[*convWasm_Val]
            vals = self.FP.lowerValList(_env, listNode, out)
            val = vals.GetAt(1)
        } else {
            val = self.FP.lowerVal(_env, node.FP.Get_exp2(_env))
        }
    }
    self.FP.assignItem(_env, &node.Nodes_Node, node.FP.Get_val(_env), index, val, out)
}
// 5211: decl @lune.@base.@convWasm.convFilter.lowerSetVal
func (self *convWasm_convFilter) lowerSetVal(_env *LnsEnv, node *Nodes_ExpSetValNode,out *LnsList2_[*WasmBin_Ins]) {
    var dstList *LnsList2_[*Nodes_Node]
    dstList = NewLnsList2_[*Nodes_Node]([]*Nodes_Node{})
    {
        _listNode := Nodes_ExpListNodeDownCastF(node.FP.Get_exp1(_env).FP)
        if !Lns_IsNil( _listNode ) {
            listNode := _listNode.(*Nodes_ExpListNode)
            for _, _exp := range( listNode.FP.Get_expList(_env).Items ) {
                exp := _exp
                dstList.Insert(exp)
            }
        } else {
            dstList.Insert(node.FP.Get_exp1(_env))
        }
    }
    var vals *LnsList2_[*convWasm_Val]
    vals = self.FP.expandDDD(_env, self.FP.lowerValList(_env, node.FP.Get_exp2(_env), out), dstList.Len(), out)
    if dstList.Len() > 1{
        var work *LnsList2_[*convWasm_Val]
        work = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
        for _, _val := range( vals.Items ) {
            val := _val
            work.Insert(self.FP.spillVal(_env, val, out))
        }
        vals = work
    }
    for _index, _dst := range( dstList.Items ) {
        index := _index + 1
        dst := _dst
        if index <= vals.Len(){
            self.FP.lowerAssign(_env, dst, vals.GetAt(index), out)
        } else { 
            self.FP.lowerAssign(_env, dst, NewconvWasm_Val(_env, WasmBin_refNull(_env, WasmBin_htAny), WasmBin_typeAnyRef, Ast_builtinTypeNil), out)
        }
    }
}
// 5239: decl @lune.@base.@convWasm.convFilter.lowerReturn
func (self *convWasm_convFilter) lowerReturn(_env *LnsEnv, node *Nodes_ReturnNode,out *LnsList2_[*WasmBin_Ins]) {
    var vals *LnsList2_[*convWasm_Val]
    vals = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
    {
        _expList := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expList ) {
            expList := _expList.(*Nodes_ExpListNode)
            vals = self.FP.lowerValList(_env, expList, out)
        }
    }
    var retTypes *LnsList2_[*Ast_TypeInfo]
    retTypes = self.ctx.FP.Get_retTypes(_env)
    var rets *LnsList2_[*WasmBin_Ins]
    rets = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _index, _retType := range( retTypes.Items ) {
        index := _index + 1
        retType := _retType
        var retRep *WasmBin_ValType
        retRep = self.ctx.FP.Get_retReps(_env).GetAt(index)
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( self.FP.normType(_env, retType).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) &&
            _env.SetStackVal( index == retTypes.Len()) ).(bool)){
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( vals.Len() == index) &&
                _env.SetStackVal( self.FP.normType(_env, vals.GetAt(index).FP.Get_typeInfo(_env)).FP.Get_kind(_env) == Ast_TypeInfoKind__DDD) ).(bool)){
                rets.Insert(vals.GetAt(index).FP.Get_ins(_env))
            } else { 
                var rest *LnsList2_[*convWasm_Val]
                rest = NewLnsList2_[*convWasm_Val]([]*convWasm_Val{})
                {
                    var _forFrom0 LnsInt = index
                    var _forTo0 LnsInt = vals.Len()
                    for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                        work := _forWork0
                        rest.Insert(vals.GetAt(work))
                    }
                }
                rets.Insert(self.FP.buildList(_env, rest))
            }
        } else if index <= vals.Len(){
            rets.Insert(self.FP.coerceRep(_env, self.FP.coerceVal(_env, vals.GetAt(index), retType), self.FP.getRep(_env, retType), retRep))
        } else { 
            rets.Insert(self.FP.getDefaultVal(_env, retRep))
        }
    }
    out.Insert(WasmBin_ret(_env, rets))
}
// 5268: decl @lune.@base.@convWasm.convFilter.lowerFor
func (self *convWasm_convFilter) lowerFor(_env *LnsEnv, node *Nodes_ForNode) *WasmBin_Ins {
    var isReal bool
    isReal = self.FP.isType(_env, node.FP.Get_val(_env).FP.Get_typeInfo(_env), Ast_builtinTypeReal)
    var rep *WasmBin_ValType
    var numType *Ast_TypeInfo
    var ge LnsInt
    ge = WasmBin_opI64GeS
    var gt LnsInt
    gt = WasmBin_opI64GtS
    var lt LnsInt
    lt = WasmBin_opI64LtS
    var add LnsInt
    add = WasmBin_opI64Add
    var zero *WasmBin_Ins
    if isReal{
        rep = WasmBin_typeF64
        numType = Ast_builtinTypeReal
        ge = WasmBin_opF64Ge
        gt = WasmBin_opF64Gt
        lt = WasmBin_opF64Lt
        add = WasmBin_opF64Add
        zero = WasmBin_f64Const(_env, 0.0)
    } else { 
        rep = WasmBin_typeI64
        numType = Ast_builtinTypeInt
        zero = WasmBin_i64Const(_env, 0)
    }
    var counter LnsInt
    counter = self.FP.tempLocal(_env, rep)
    var toIdx LnsInt
    toIdx = self.FP.tempLocal(_env, rep)
    var stepIdx LnsInt
    stepIdx = self.FP.tempLocal(_env, rep)
    var list *LnsList2_[*WasmBin_Ins]
    list = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    list.Insert(convWasm_ls_3_(_env, counter, self.FP.lowerExpAs(_env, node.FP.Get_init(_env), numType)))
    list.Insert(convWasm_ls_3_(_env, toIdx, self.FP.lowerExpAs(_env, node.FP.Get_to(_env), numType)))
    {
        _delta := node.FP.Get_delta(_env)
        if !Lns_IsNil( _delta ) {
            delta := _delta.(*Nodes_Node)
            list.Insert(convWasm_ls_3_(_env, stepIdx, self.FP.lowerExpAs(_env, delta, numType)))
        } else {
            if isReal{
                list.Insert(convWasm_ls_3_(_env, stepIdx, WasmBin_f64Const(_env, 1.0)))
            } else { 
                list.Insert(convWasm_ls_3_(_env, stepIdx, WasmBin_i64Const(_env, 1)))
            }
        }
    }
    self.FP.enterLoop(_env)
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    body.Insert(WasmBin_brIf(_env, 1, WasmBin_ifElse(_env, convWasm_p_4_(_env, ge, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, stepIdx), zero))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, gt, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, counter), convWasm_lg_2_(_env, toIdx)))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_p_4_(_env, lt, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, counter), convWasm_lg_2_(_env, toIdx)))))), WasmBin_typeI32)))
    body.Insert(self.FP.declareVar(_env, node.FP.Get_val(_env), self.FP.coerceRep(_env, convWasm_lg_2_(_env, counter), rep, self.FP.getRep(_env, node.FP.Get_val(_env).FP.Get_typeInfo(_env)))))
    body.Insert(convWasm_ls_3_(_env, counter, convWasm_p_4_(_env, add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, counter), convWasm_lg_2_(_env, stepIdx))))))
    convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, node.FP.Get_block(_env)))
    body.Insert(WasmBin_br(_env, 0))
    list.Insert(self.FP.exitLoop(_env, body))
    return WasmBin_seq(_env, list)
}
// 5321: decl @lune.@base.@convWasm.convFilter.lowerForeach
func (self *convWasm_convFilter) lowerForeach(_env *LnsEnv, node *Nodes_Node,valSym *Ast_SymbolInfo,keySym LnsAny,exp *Nodes_Node,block *Nodes_BlockNode,sort bool) *WasmBin_Ins {
    var expType *Ast_TypeInfo
    expType = self.FP.normType(_env, exp.FP.Get_expType(_env))
    var list *LnsList2_[*WasmBin_Ins]
    list = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var pos LnsInt
    pos = self.FP.tempLocal(_env, WasmBin_typeI32)
    list.Insert(convWasm_ls_3_(_env, pos, WasmBin_i32Const(_env, 0)))
    var luavalFlag bool
    luavalFlag = self.FP.isLuaval(_env, exp.FP.Get_expType(_env))
    if luavalFlag{
        expType = self.FP.normType(_env, exp.FP.Get_expType(_env).FP.Get_nonnilableType(_env).FP.Get_extedType(_env))
    }
    var convWasm_lowerColl func(_env *LnsEnv) *WasmBin_Ins
    convWasm_lowerColl = func(_env *LnsEnv) *WasmBin_Ins {
        if Lns_op_not(luavalFlag){
            return self.FP.lowerExp(_env, exp)
        }
        var luaVal *WasmBin_Ins
        luaVal = self.FP.lowerExpAs(_env, exp, Ast_builtinTypeStem_)
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( expType.FP.Get_kind(_env) == Ast_TypeInfoKind__List) ||
            _env.SetStackVal( expType.FP.Get_kind(_env) == Ast_TypeInfoKind__Array) ).(bool){
            return WasmBin_call(_env, self.rt.FP.FnLuaToList(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](luaVal)))
        }
        return WasmBin_call(_env, self.rt.FP.FnLuaToMap(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](luaVal, WasmBin_i32Const(_env, 0))))
    }
    if _switch0 := expType.FP.Get_kind(_env); _switch0 == Ast_TypeInfoKind__List || _switch0 == Ast_TypeInfoKind__Array {
        var coll LnsInt
        coll = self.FP.tempLocal(_env, self.rt.FP.Get_listType(_env))
        list.Insert(convWasm_ls_3_(_env, coll, convWasm_lowerColl(_env)))
        self.FP.enterLoop(_env)
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        body.Insert(WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, convWasm_lg_2_(_env, coll)))))))
        var item *WasmBin_Ins
        item = WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListItems, convWasm_lg_2_(_env, coll)), convWasm_lg_2_(_env, pos))
        body.Insert(self.FP.declareVar(_env, valSym, self.FP.coerceRep(_env, item, WasmBin_typeAnyRef, self.FP.getRep(_env, valSym.FP.Get_typeInfo(_env)))))
        body.Insert(convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))))
        if keySym != nil{
            keySym_1443 := keySym.(*Ast_SymbolInfo)
            body.Insert(self.FP.declareVar(_env, keySym_1443, convWasm_p_4_(_env, WasmBin_opI64ExtendI32U, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos))))))
        }
        convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, block))
        body.Insert(WasmBin_br(_env, 0))
        list.Insert(self.FP.exitLoop(_env, body))
        return WasmBin_seq(_env, list)
    } else if _switch0 == Ast_TypeInfoKind__Map || _switch0 == Ast_TypeInfoKind__Set {
        var isSet bool
        isSet = expType.FP.Get_kind(_env) == Ast_TypeInfoKind__Set
        var coll LnsInt
        coll = self.FP.tempLocal(_env, self.rt.FP.Get_mapType(_env))
        list.Insert(convWasm_ls_3_(_env, coll, convWasm_lowerColl(_env)))
        if sort{
            var keys LnsInt
            keys = self.FP.tempLocal(_env, self.rt.FP.Get_listType(_env))
            list.Insert(convWasm_ls_3_(_env, keys, WasmBin_call(_env, self.rt.FP.FnMapKeys(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, coll))))))
            list.Insert(WasmBin_call(_env, self.rt.FP.FnListSort(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, keys), WasmBin_refNull(_env, self.rt.FP.GetLessCloInfo(_env).FP.Get_cloType(_env))))))
            var key LnsInt
            key = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
            self.FP.enterLoop(_env)
            var body *LnsList2_[*WasmBin_Ins]
            body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
            body.Insert(WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListLen, convWasm_lg_2_(_env, keys)))))))
            body.Insert(convWasm_ls_3_(_env, key, WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tList(_env), WasmRt_fListItems, convWasm_lg_2_(_env, keys)), convWasm_lg_2_(_env, pos))))
            body.Insert(convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))))
            var valIns *WasmBin_Ins
            if isSet{
                valIns = convWasm_lg_2_(_env, key)
            } else { 
                valIns = WasmBin_call(_env, self.rt.FP.FnMapGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, coll), convWasm_lg_2_(_env, key))))
            }
            body.Insert(self.FP.declareVar(_env, valSym, self.FP.coerceRep(_env, valIns, WasmBin_typeAnyRef, self.FP.getRep(_env, valSym.FP.Get_typeInfo(_env)))))
            if keySym != nil{
                keySym_1455 := keySym.(*Ast_SymbolInfo)
                body.Insert(self.FP.declareVar(_env, keySym_1455, self.FP.coerceRep(_env, convWasm_lg_2_(_env, key), WasmBin_typeAnyRef, self.FP.getRep(_env, keySym_1455.FP.Get_typeInfo(_env)))))
            }
            convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, block))
            body.Insert(WasmBin_br(_env, 0))
            list.Insert(self.FP.exitLoop(_env, body))
            return WasmBin_seq(_env, list)
        }
        self.FP.enterLoop(_env)
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        body.Insert(convWasm_ls_3_(_env, pos, WasmBin_call(_env, self.rt.FP.FnMapNext(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, coll), convWasm_lg_2_(_env, pos))))))
        body.Insert(WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 0))))))
        var keyIns *WasmBin_Ins
        keyIns = self.rt.FP.MapKeyAt(_env, convWasm_lg_2_(_env, coll), convWasm_lg_2_(_env, pos))
        var valIns *WasmBin_Ins
        if isSet{
            valIns = keyIns
        } else { 
            valIns = self.rt.FP.MapValAt(_env, convWasm_lg_2_(_env, coll), convWasm_lg_2_(_env, pos))
        }
        body.Insert(self.FP.declareVar(_env, valSym, self.FP.coerceRep(_env, valIns, WasmBin_typeAnyRef, self.FP.getRep(_env, valSym.FP.Get_typeInfo(_env)))))
        if keySym != nil{
            keySym_1462 := keySym.(*Ast_SymbolInfo)
            body.Insert(self.FP.declareVar(_env, keySym_1462, self.FP.coerceRep(_env, keyIns, WasmBin_typeAnyRef, self.FP.getRep(_env, keySym_1462.FP.Get_typeInfo(_env)))))
        }
        body.Insert(convWasm_ls_3_(_env, pos, convWasm_p_4_(_env, WasmBin_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, pos), WasmBin_i32Const(_env, 1))))))
        convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, block))
        body.Insert(WasmBin_br(_env, 0))
        list.Insert(self.FP.exitLoop(_env, body))
        return WasmBin_seq(_env, list)
    }
    self.FP.error(_env, node, _env.GetVM().String_format("not support foreach -- %s", Lns_2DDD(expType.FP.GetTxt(_env, nil, nil, nil))))
// insert a dummy
    return nil
}
// 5431: decl @lune.@base.@convWasm.convFilter.getCondRetInfo
func (self *convWasm_convFilter) getCondRetInfo(_env *LnsEnv, node *Nodes_Node) LnsAny {
    {
        _work := Nodes_DeclVarNodeDownCastF(node.FP)
        if !Lns_IsNil( _work ) {
            work := _work.(*Nodes_DeclVarNode)
            return work.FP.Get_condRetInfo(_env)
        } else {
            {
                _work := Nodes_ExpSetValNodeDownCastF(node.FP)
                if !Lns_IsNil( _work ) {
                    work := _work.(*Nodes_ExpSetValNode)
                    return work.FP.Get_condRetInfo(_env)
                } else {
                    {
                        _work := Nodes_ExpSetItemNodeDownCastF(node.FP)
                        if !Lns_IsNil( _work ) {
                            work := _work.(*Nodes_ExpSetItemNode)
                            return work.FP.Get_condRetInfo(_env)
                        } else {
                            {
                                _work := Nodes_IfUnwrapNodeDownCastF(node.FP)
                                if !Lns_IsNil( _work ) {
                                    work := _work.(*Nodes_IfUnwrapNode)
                                    return work.FP.Get_condRetInfo(_env)
                                } else {
                                    {
                                        _work := Nodes_IfNodeDownCastF(node.FP)
                                        if !Lns_IsNil( _work ) {
                                            work := _work.(*Nodes_IfNode)
                                            return work.FP.Get_stmtList(_env).GetAt(1).FP.Get_condRetInfo(_env)
                                        } else {
                                            {
                                                _work := Nodes_StmtExpNodeDownCastF(node.FP)
                                                if !Lns_IsNil( _work ) {
                                                    work := _work.(*Nodes_StmtExpNode)
                                                    {
                                                        _condRetList := Nodes_CondRetListNodeDownCastF(work.FP.Get_exp(_env).FP)
                                                        if !Lns_IsNil( _condRetList ) {
                                                            condRetList := _condRetList.(*Nodes_CondRetListNode)
                                                            return condRetList.FP.Get_info(_env)
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    return nil
}
// 5456: decl @lune.@base.@convWasm.convFilter.lowerCondRetOf
func (self *convWasm_convFilter) lowerCondRetOf(_env *LnsEnv, node *Nodes_Node,out *LnsList2_[*WasmBin_Ins]) {
    var info *Nodes_CondRetInfo
    
    {
        _info := self.FP.getCondRetInfo(_env, node)
        if _info == nil{
            return 
        } else {
            info = _info.(*Nodes_CondRetInfo)
        }
    }
    var retReps *LnsList2_[*WasmBin_ValType]
    retReps = self.ctx.FP.Get_retReps(_env)
    var convWasm_retWith func(_env *LnsEnv, first LnsAny) *WasmBin_Ins
    convWasm_retWith = func(_env *LnsEnv, first LnsAny) *WasmBin_Ins {
        var vals *LnsList2_[*WasmBin_Ins]
        vals = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        for _index, _rep := range( retReps.Items ) {
            index := _index + 1
            rep := _rep
            if first != nil{
                first_1496 := first.(*WasmBin_Ins)
                if index == 1{
                    vals.Insert(first_1496)
                    continue
                }
            }
            vals.Insert(self.FP.getDefaultVal(_env, rep))
        }
        return WasmBin_ret(_env, vals)
    }
    for _, _condRet := range( info.FP.Get_list(_env).Items ) {
        condRet := _condRet
        if _switch0 := condRet.FP.Get_condKind(_env); _switch0 == Nodes_CondRetKind__Nilable {
            var val *convWasm_Val
            val = self.FP.lowerVal(_env, condRet.FP.Get_exp(_env))
            var work LnsInt
            work = self.FP.tempLocal(_env, val.FP.Get_rep(_env))
            out.Insert(convWasm_ls_3_(_env, work, val.FP.Get_ins(_env)))
            out.Insert(WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, work)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_retWith(_env, nil))), nil, nil))
            self.node2val.Set(&condRet.Nodes_Node,NewconvWasm_Val(_env, convWasm_lg_2_(_env, work), val.FP.Get_rep(_env), val.FP.Get_typeInfo(_env)))
        } else if _switch0 == Nodes_CondRetKind__Ret {
            var algeType *Ast_TypeInfo
            algeType = self.FP.getGenSrc(_env, condRet.FP.Get_exp(_env).FP.Get_expType(_env))
            var algeInfo *convWasm_AlgeInfo
            
            {
                _algeInfo := self.FP.getAlgeInfo(_env, algeType)
                if _algeInfo == nil{
                    self.FP.error(_env, &condRet.Nodes_Node, _env.GetVM().String_format("not found alge -- %s", Lns_2DDD(algeType.FP.GetTxt(_env, nil, nil, nil))))
                } else {
                    algeInfo = _algeInfo.(*convWasm_AlgeInfo)
                }
            }
            var algeRep *WasmBin_ValType
            algeRep = WasmBin_refType(_env, algeInfo.FP.Get_baseType(_env))
            var work LnsInt
            work = self.FP.tempLocal(_env, algeRep)
            out.Insert(convWasm_ls_3_(_env, work, self.FP.lowerExpAs(_env, condRet.FP.Get_exp(_env), algeType)))
            out.Insert(WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32Eq, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structGet(_env, algeInfo.FP.Get_baseType(_env), 1, convWasm_lg_2_(_env, work)), WasmBin_i32Const(_env, Lns_unwrap( algeInfo.FP.Get_name2tag(_env).Get("Err")).(LnsInt))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_retWith(_env, self.FP.coerceRep(_env, convWasm_lg_2_(_env, work), algeRep, retReps.GetAt(1))))), nil, nil))
            var okType LnsInt
            okType = Lns_unwrap( algeInfo.FP.Get_name2type(_env).Get("Ok")).(LnsInt)
            self.node2val.Set(&condRet.Nodes_Node,NewconvWasm_Val(_env, WasmBin_structGet(_env, okType, 2, WasmBin_refCast(_env, okType, convWasm_lg_2_(_env, work))), WasmBin_typeAnyRef, Ast_builtinTypeStem_))
        } else if _switch0 == Nodes_CondRetKind__Two {
            var vals *LnsList2_[*convWasm_Val]
            vals = self.FP.lowerMulti(_env, condRet.FP.Get_exp(_env), out)
            var val *convWasm_Val
            val = vals.GetAt(1)
            var errVal *convWasm_Val
            errVal = vals.GetAt(2)
            var retVals *LnsList2_[*WasmBin_Ins]
            retVals = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
            for _index, _rep := range( retReps.Items ) {
                index := _index + 1
                rep := _rep
                if index == 2{
                    retVals.Insert(self.FP.coerceRep(_env, self.FP.coerceVal(_env, errVal, Ast_builtinTypeStem_), WasmBin_typeAnyRef, rep))
                } else { 
                    retVals.Insert(self.FP.getDefaultVal(_env, rep))
                }
            }
            out.Insert(WasmBin_ifElse(_env, WasmBin_refIsNull(_env, val.FP.Get_ins(_env)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, retVals))), nil, nil))
            self.node2val.Set(&condRet.Nodes_Node,val)
        }
    }
}
// 5524: decl @lune.@base.@convWasm.convFilter.lowerStmtExp
func (self *convWasm_convFilter) lowerStmtExp(_env *LnsEnv, exp *Nodes_Node,out *LnsList2_[*WasmBin_Ins]) {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( exp.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpSetVal(_env)) ||
        _env.SetStackVal( exp.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpSetItem(_env)) ).(bool){
        self.FP.lowerStmt(_env, exp, out)
        return 
    }
    {
        _callNode := Nodes_ExpCallNodeDownCastF(exp.FP)
        if !Lns_IsNil( _callNode ) {
            callNode := _callNode.(*Nodes_ExpCallNode)
            var prefix LnsAny
            prefix = self.FP.getNilAccPrefix(_env, &callNode.Nodes_Node)
            var work *LnsList2_[*WasmBin_Ins]
            work = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
            var prefixLocal LnsInt
            prefixLocal = -1
            if prefix != nil{
                prefix_1530 := prefix.(*Nodes_Node)
                var prefixVal *convWasm_Val
                prefixVal = self.FP.lowerVal(_env, prefix_1530)
                prefixLocal = self.FP.tempLocal(_env, prefixVal.FP.Get_rep(_env))
                out.Insert(convWasm_ls_3_(_env, prefixLocal, prefixVal.FP.Get_ins(_env)))
                self.node2val.Set(prefix_1530,NewconvWasm_Val(_env, convWasm_lg_2_(_env, prefixLocal), prefixVal.FP.Get_rep(_env), prefixVal.FP.Get_typeInfo(_env)))
                self.nilAccDoneSet.Add(&callNode.Nodes_Node)
                self.nilAccDoneSet.Add(callNode.FP.Get_func(_env))
            }
            var ins *WasmBin_Ins
            var retTypes *LnsList2_[*Ast_TypeInfo]
            ins,retTypes = self.FP.lowerCall(_env, callNode)
            work.Insert(ins)
            for range( retTypes.Items ) {
                work.Insert(WasmBin_drop(_env, WasmBin_nop(_env)))
            }
            if prefix != nil{
                prefix_1537 := prefix.(*Nodes_Node)
                self.node2val.Set(prefix_1537,nil)
                out.Insert(WasmBin_ifElse(_env, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refIsNull(_env, convWasm_lg_2_(_env, prefixLocal))))), work, nil, nil))
            } else {
                convWasm_append_17_(_env, out, work)
            }
            return 
        }
    }
    {
        _superNode := Nodes_ExpCallSuperNodeDownCastF(exp.FP)
        if !Lns_IsNil( _superNode ) {
            superNode := _superNode.(*Nodes_ExpCallSuperNode)
            var ins *WasmBin_Ins
            var retTypes *LnsList2_[*Ast_TypeInfo]
            ins,retTypes = self.FP.lowerCallSuper(_env, superNode)
            out.Insert(ins)
            for range( retTypes.Items ) {
                out.Insert(WasmBin_drop(_env, WasmBin_nop(_env)))
            }
            return 
        }
    }
    {
        _ctorNode := Nodes_ExpCallSuperCtorNodeDownCastF(exp.FP)
        if !Lns_IsNil( _ctorNode ) {
            ctorNode := _ctorNode.(*Nodes_ExpCallSuperCtorNode)
            self.FP.lowerCallSuperCtor(_env, ctorNode, out)
            return 
        }
    }
    {
        _macroNode := Nodes_ExpMacroExpNodeDownCastF(exp.FP)
        if !Lns_IsNil( _macroNode ) {
            macroNode := _macroNode.(*Nodes_ExpMacroExpNode)
            for _, _ins := range( self.FP.lowerStmtList(_env, macroNode.FP.Get_stmtList(_env)).Items ) {
                ins := _ins
                out.Insert(ins)
            }
            return 
        }
    }
    var val *convWasm_Val
    val = self.FP.lowerVal(_env, exp)
    if val.FP.Get_typeInfo(_env) == Ast_builtinTypeNone{
        out.Insert(val.FP.Get_ins(_env))
    } else { 
        out.Insert(WasmBin_drop(_env, val.FP.Get_ins(_env)))
    }
}
// 5583: decl @lune.@base.@convWasm.convFilter.lowerCallSuperCtor
func (self *convWasm_convFilter) lowerCallSuperCtor(_env *LnsEnv, node *Nodes_ExpCallSuperCtorNode,out *LnsList2_[*WasmBin_Ins]) {
    var info *convWasm_ClassInfo
    
    {
        _info := self.FP.getClassInfo(_env, node.FP.Get_superType(_env))
        if _info == nil{
            self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not found class -- %s", Lns_2DDD(node.FP.Get_superType(_env).FP.GetTxt(_env, nil, nil, nil))))
        } else {
            info = _info.(*convWasm_ClassInfo)
        }
    }
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, self.ctx.FP.Get_selfLocal(_env))))
    convWasm_append_17_(_env, args, self.FP.lowerArgs(_env, node.FP.Get_expList(_env), self.FP.getCtorType(_env, info.FP.Get_typeInfo(_env)).FP.Get_argTypeInfoList(_env), out))
    out.Insert(WasmBin_call(_env, info.FP.Get_ctorFunc(_env), args))
}
// 5594: decl @lune.@base.@convWasm.convFilter.lowerStmt
func (self *convWasm_convFilter) lowerStmt(_env *LnsEnv, node *Nodes_Node,out *LnsList2_[*WasmBin_Ins]) {
    self.FP.lowerCondRetOf(_env, node, out)
    if _switch0 := node.FP.Get_kind(_env); _switch0 == Nodes_NodeKind_get_None(_env) || _switch0 == Nodes_NodeKind_get_BlankLine(_env) || _switch0 == Nodes_NodeKind_get_Shebang(_env) || _switch0 == Nodes_NodeKind_get_LuneControl(_env) || _switch0 == Nodes_NodeKind_get_DeclMacro(_env) || _switch0 == Nodes_NodeKind_get_TestBlock(_env) || _switch0 == Nodes_NodeKind_get_DeclForm(_env) || _switch0 == Nodes_NodeKind_get_ProtoClass(_env) || _switch0 == Nodes_NodeKind_get_DeclEnum(_env) || _switch0 == Nodes_NodeKind_get_DeclAlge(_env) || _switch0 == Nodes_NodeKind_get_ProtoMethod(_env) || _switch0 == Nodes_NodeKind_get_Subfile(_env) || _switch0 == Nodes_NodeKind_get_Alias(_env) || _switch0 == Nodes_NodeKind_get_TestCase(_env) || _switch0 == Nodes_NodeKind_get_Provide(_env) {
        return 
    } else if _switch0 == Nodes_NodeKind_get_DeclFunc(_env) {
        var declNode *Nodes_DeclFuncNode
        declNode = Lns_unwrap( Nodes_DeclFuncNodeDownCastF(node.FP)).(*Nodes_DeclFuncNode)
        var funcIndex LnsInt
        funcIndex = self.FP.registerFunc(_env, declNode.FP.Get_declInfo(_env), declNode.FP.Get_expType(_env))
        var closure *WasmBin_Ins
        closure = self.FP.compileFunc(_env, node, declNode.FP.Get_declInfo(_env), declNode.FP.Get_expType(_env), funcIndex)
        {
            _symbol := declNode.FP.Get_declInfo(_env).FP.Get_symbol(_env)
            if !Lns_IsNil( _symbol ) {
                symbol := _symbol.(*Ast_SymbolInfo)
                if Lns_op_not(self.FP.isGlobalSym(_env, symbol.FP.GetOrg(_env))){
                    out.Insert(self.FP.writeSym(_env, node, symbol, closure))
                }
            }
        }
        return 
    } else if _switch0 == Nodes_NodeKind_get_DeclClass(_env) {
        self.FP.lowerDeclClass(_env, Lns_unwrap( Nodes_DeclClassNodeDownCastF(node.FP)).(*Nodes_DeclClassNode), out)
        return 
    } else if _switch0 == Nodes_NodeKind_get_DeclMethod(_env) {
        var declNode *Nodes_DeclMethodNode
        declNode = Lns_unwrap( Nodes_DeclMethodNodeDownCastF(node.FP)).(*Nodes_DeclMethodNode)
        self.FP.compileMethodNode(_env, node, declNode.FP.Get_declInfo(_env))
        return 
    } else if _switch0 == Nodes_NodeKind_get_StmtExp(_env) {
        var exp *Nodes_Node
        exp = (Lns_unwrap( Nodes_StmtExpNodeDownCastF(node.FP)).(*Nodes_StmtExpNode)).FP.Get_exp(_env)
        {
            _condRetList := Nodes_CondRetListNodeDownCastF(exp.FP)
            if !Lns_IsNil( _condRetList ) {
                condRetList := _condRetList.(*Nodes_CondRetListNode)
                exp = condRetList.FP.Get_exp(_env)
            }
        }
        self.FP.lowerStmtExp(_env, exp, out)
        return 
    } else if _switch0 == Nodes_NodeKind_get_ExpCall(_env) || _switch0 == Nodes_NodeKind_get_ExpCallSuperCtor(_env) || _switch0 == Nodes_NodeKind_get_ExpMacroExp(_env) {
        self.FP.lowerStmtExp(_env, node, out)
        return 
    } else if _switch0 == Nodes_NodeKind_get_DeclVar(_env) {
        self.FP.lowerDeclVar(_env, Lns_unwrap( Nodes_DeclVarNodeDownCastF(node.FP)).(*Nodes_DeclVarNode), out)
        return 
    } else if _switch0 == Nodes_NodeKind_get_ExpSetVal(_env) {
        self.FP.lowerSetVal(_env, Lns_unwrap( Nodes_ExpSetValNodeDownCastF(node.FP)).(*Nodes_ExpSetValNode), out)
        return 
    } else if _switch0 == Nodes_NodeKind_get_ExpSetItem(_env) {
        self.FP.lowerSetItem(_env, Lns_unwrap( Nodes_ExpSetItemNodeDownCastF(node.FP)).(*Nodes_ExpSetItemNode), out)
        return 
    } else if _switch0 == Nodes_NodeKind_get_UnwrapSet(_env) {
        self.FP.lowerUnwrapSet(_env, Lns_unwrap( Nodes_UnwrapSetNodeDownCastF(node.FP)).(*Nodes_UnwrapSetNode), out)
        return 
    } else if _switch0 == Nodes_NodeKind_get_If(_env) {
        out.Insert(self.FP.lowerIf(_env, Lns_unwrap( Nodes_IfNodeDownCastF(node.FP)).(*Nodes_IfNode)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_IfUnwrap(_env) {
        out.Insert(self.FP.lowerIfUnwrap(_env, Lns_unwrap( Nodes_IfUnwrapNodeDownCastF(node.FP)).(*Nodes_IfUnwrapNode)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_When(_env) {
        out.Insert(self.FP.lowerWhen(_env, Lns_unwrap( Nodes_WhenNodeDownCastF(node.FP)).(*Nodes_WhenNode)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Switch(_env) {
        out.Insert(self.FP.lowerSwitch(_env, Lns_unwrap( Nodes_SwitchNodeDownCastF(node.FP)).(*Nodes_SwitchNode)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Match(_env) {
        out.Insert(self.FP.lowerMatch(_env, Lns_unwrap( Nodes_MatchNodeDownCastF(node.FP)).(*Nodes_MatchNode)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_While(_env) {
        var whileNode *Nodes_WhileNode
        whileNode = Lns_unwrap( Nodes_WhileNodeDownCastF(node.FP)).(*Nodes_WhileNode)
        self.FP.enterLoop(_env)
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        if Lns_op_not(whileNode.FP.Get_infinit(_env)){
            body.Insert(WasmBin_brIf(_env, 1, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerCond(_env, whileNode.FP.Get_exp(_env)))))))
        }
        convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, whileNode.FP.Get_block(_env)))
        body.Insert(WasmBin_br(_env, 0))
        out.Insert(self.FP.exitLoop(_env, body))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Repeat(_env) {
        var repeatNode *Nodes_RepeatNode
        repeatNode = Lns_unwrap( Nodes_RepeatNodeDownCastF(node.FP)).(*Nodes_RepeatNode)
        self.FP.enterLoop(_env)
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, repeatNode.FP.Get_block(_env)))
        body.Insert(WasmBin_brIf(_env, 0, convWasm_p_4_(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerCond(_env, repeatNode.FP.Get_exp(_env)))))))
        out.Insert(self.FP.exitLoop(_env, body))
        return 
    } else if _switch0 == Nodes_NodeKind_get_For(_env) {
        out.Insert(self.FP.lowerFor(_env, Lns_unwrap( Nodes_ForNodeDownCastF(node.FP)).(*Nodes_ForNode)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Foreach(_env) {
        var foreachNode *Nodes_ForeachNode
        foreachNode = Lns_unwrap( Nodes_ForeachNodeDownCastF(node.FP)).(*Nodes_ForeachNode)
        out.Insert(self.FP.lowerForeach(_env, node, foreachNode.FP.Get_val(_env), foreachNode.FP.Get_key(_env), foreachNode.FP.Get_exp(_env), foreachNode.FP.Get_block(_env), false))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Forsort(_env) {
        var forsortNode *Nodes_ForsortNode
        forsortNode = Lns_unwrap( Nodes_ForsortNodeDownCastF(node.FP)).(*Nodes_ForsortNode)
        out.Insert(self.FP.lowerForeach(_env, node, forsortNode.FP.Get_val(_env), forsortNode.FP.Get_key(_env), forsortNode.FP.Get_exp(_env), forsortNode.FP.Get_block(_env), true))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Continue(_env) {
        var breakList *LnsList2_[LnsInt]
        breakList = self.ctx.FP.Get_breakDepthList(_env)
        if breakList.Len() == 0{
            self.FP.error(_env, node, "continue outside loop")
        }
        out.Insert(WasmBin_br(_env, self.ctx.FP.Get_depth(_env) - (breakList.GetAt(breakList.Len()) + 1)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Break(_env) {
        var breakList *LnsList2_[LnsInt]
        breakList = self.ctx.FP.Get_breakDepthList(_env)
        if breakList.Len() == 0{
            self.FP.error(_env, node, "break outside loop")
        }
        out.Insert(WasmBin_br(_env, self.ctx.FP.Get_depth(_env) - breakList.GetAt(breakList.Len())))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Return(_env) {
        self.FP.lowerReturn(_env, Lns_unwrap( Nodes_ReturnNodeDownCastF(node.FP)).(*Nodes_ReturnNode), out)
        return 
    } else if _switch0 == Nodes_NodeKind_get_Scope(_env) {
        var scopeNode *Nodes_ScopeNode
        scopeNode = Lns_unwrap( Nodes_ScopeNodeDownCastF(node.FP)).(*Nodes_ScopeNode)
        for _, _symbol := range( scopeNode.FP.Get_symbolList(_env).Items ) {
            symbol := _symbol
            {
                _aliasSym := scopeNode.FP.Get_scope(_env).FP.Get_symbol2SymbolInfoMap(_env).Get(symbol.FP.Get_name(_env))
                if !Lns_IsNil( _aliasSym ) {
                    aliasSym := _aliasSym.(*Ast_SymbolInfo)
                    self.symAlias.Set(aliasSym.FP.GetOrg(_env),self.FP.resolveSym(_env, symbol.FP.GetOrg(_env)))
                }
            }
        }
        convWasm_append_17_(_env, out, self.FP.lowerBlock(_env, scopeNode.FP.Get_block(_env)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Import(_env) {
        var info *Nodes_ImportInfo
        info = (Lns_unwrap( Nodes_ImportNodeDownCastF(node.FP)).(*Nodes_ImportNode)).FP.Get_info(_env)
        var initIndex LnsInt
        
        {
            _initIndex := self.module2init.Get(info.FP.Get_modulePath(_env))
            if _initIndex == nil{
                self.FP.error(_env, node, _env.GetVM().String_format("not found the imported module -- %s", Lns_2DDD(info.FP.Get_modulePath(_env))))
            } else {
                initIndex = _initIndex.(LnsInt)
            }
        }
        if info.FP.Get_lazy(_env) == Nodes_LazyLoad__Off{
            out.Insert(WasmBin_call(_env, initIndex, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})))
        }
        return 
    } else if _switch0 == Nodes_NodeKind_get_Apply(_env) {
        out.Insert(self.FP.lowerApply(_env, Lns_unwrap( Nodes_ApplyNodeDownCastF(node.FP)).(*Nodes_ApplyNode)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_AsyncLock(_env) {
        convWasm_append_17_(_env, out, self.FP.lowerBlock(_env, (Lns_unwrap( Nodes_AsyncLockNodeDownCastF(node.FP)).(*Nodes_AsyncLockNode)).FP.Get_block(_env)))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Block(_env) {
        convWasm_append_17_(_env, out, self.FP.lowerBlock(_env, Lns_unwrap( Nodes_BlockNodeDownCastF(node.FP)).(*Nodes_BlockNode)))
        return 
    }
    self.FP.notSupport(_env, node)
}
// 5776: decl @lune.@base.@convWasm.convFilter.registerFunc
func (self *convWasm_convFilter) registerFunc(_env *LnsEnv, declInfo *Nodes_DeclFuncInfo,funcType *Ast_TypeInfo) LnsInt {
    {
        _index := self.func2index.Get(funcType)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var info *WasmRt_CloInfo
    info = self.FP.getFuncCloInfo(_env, funcType)
    var index LnsInt
    index,_ = self.mod.FP.AddFuncWithType(_env, info.FP.Get_funcType(_env), 1 + funcType.FP.Get_argTypeInfoList(_env).Len())
    self.func2index.Set(funcType,index)
    self.mod.FP.SetFuncName(_env, index, funcType.FP.GetTxt(_env, nil, nil, nil))
    return index
}
// 5788: decl @lune.@base.@convWasm.convFilter.getCaptureSyms
func (self *convWasm_convFilter) getCaptureSyms(_env *LnsEnv, funcType *Ast_TypeInfo) *LnsList2_[*Ast_SymbolInfo] {
    var list *LnsList2_[*Ast_SymbolInfo]
    list = NewLnsList2_[*Ast_SymbolInfo]([]*Ast_SymbolInfo{})
    {
        _scope := funcType.FP.Get_scope(_env)
        if !Lns_IsNil( _scope ) {
            scope := _scope.(*Ast_Scope)
            var addedSet *LnsSet2_[*Ast_SymbolInfo]
            addedSet = NewLnsSet2_[*Ast_SymbolInfo]([]*Ast_SymbolInfo{})
            for _, _symbol := range( scope.FP.Get_closureSymList(_env).Items ) {
                symbol := _symbol
                var orgSym *Ast_SymbolInfo
                orgSym = self.FP.resolveSym(_env, symbol.FP.GetOrg(_env))
                if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                    _env.SetStackVal( Lns_op_not(self.FP.isGlobalSym(_env, orgSym))) &&
                    _env.SetStackVal( Lns_op_not(addedSet.Has(orgSym))) ).(bool)){
                    addedSet.Add(orgSym)
                    list.Insert(orgSym)
                }
            }
        }
    }
    return list
}
// 5805: decl @lune.@base.@convWasm.convFilter.getCellRef
func (self *convWasm_convFilter) getCellRef(_env *LnsEnv, node *Nodes_Node,symbol *Ast_SymbolInfo) *WasmBin_Ins {
    var loc LnsAny
    
    {
        _loc := self.FP.findLoc(_env, symbol)
        if _loc == nil{
            self.FP.error(_env, node, _env.GetVM().String_format("not found closure symbol -- %s", Lns_2DDD(symbol.FP.Get_name(_env))))
        } else {
            loc = _loc
        }
    }
    switch _matchExp0 := loc.(type) {
    case *convWasm_VarLoc__CellLocal:
        index := _matchExp0.Val1
        return convWasm_lg_2_(_env, index)
    case *convWasm_VarLoc__Captured:
        field := _matchExp0.Val1
        return self.FP.envGet(_env, field)
    case *convWasm_VarLoc__Local:
        index := _matchExp0.Val1
        var rep *WasmBin_ValType
        rep = self.FP.getRep(_env, symbol.FP.Get_typeInfo(_env))
        return WasmBin_structNew(_env, self.rt.FP.GetCellType(_env, rep), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, index))))
    case *convWasm_VarLoc__Global:
        self.FP.error(_env, node, _env.GetVM().String_format("global closure symbol -- %s", Lns_2DDD(symbol.FP.Get_name(_env))))
    }
// insert a dummy
    return nil
}
// 5828: decl @lune.@base.@convWasm.convFilter.setupParams
func (self *convWasm_convFilter) setupParams(_env *LnsEnv, declInfo *Nodes_DeclFuncInfo,funcType *Ast_TypeInfo,body *LnsList2_[*WasmBin_Ins]) {
    if declInfo.FP.Get_has__func__Symbol(_env){
        {
            _block := declInfo.FP.Get_body(_env)
            if !Lns_IsNil( _block ) {
                block := _block.(*Nodes_BlockNode)
                var scope *Ast_Scope
                scope = block.FP.Get_scope(_env)
                {
                    var _forFrom0 LnsInt = 1
                    var _forTo0 LnsInt = 3
                    for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                        {
                            _funcSym := scope.FP.Get_symbol2SymbolInfoMap(_env).Get("__func__")
                            if !Lns_IsNil( _funcSym ) {
                                funcSym := _funcSym.(*Ast_SymbolInfo)
                                {
                                    _name := declInfo.FP.Get_name(_env)
                                    if !Lns_IsNil( _name ) {
                                        name := _name.(*Types_Token)
                                        self.funcSym2name.Set(funcSym,name.Txt)
                                    }
                                }
                                break
                            }
                        }
                        scope = scope.FP.Get_parent(_env)
                    }
                }
            }
        }
    }
    for _index, _arg := range( declInfo.FP.Get_argList(_env).Items ) {
        index := _index + 1
        arg := _arg
        if arg.FP.Get_kind(_env) == Nodes_NodeKind_get_DeclArgDDD(_env){
            {
                _scope := funcType.FP.Get_scope(_env)
                if !Lns_IsNil( _scope ) {
                    scope := _scope.(*Ast_Scope)
                    {
                        _symbol := scope.FP.Get_symbol2SymbolInfoMap(_env).Get("...")
                        if !Lns_IsNil( _symbol ) {
                            symbol := _symbol.(*Ast_SymbolInfo)
                            self.ctx.FP.Get_sym2loc(_env).Set(symbol.FP.GetOrg(_env),&convWasm_VarLoc__Local{index})
                        }
                    }
                }
            }
        }
        {
            _argNode := Nodes_DeclArgNodeDownCastF(arg.FP)
            if !Lns_IsNil( _argNode ) {
                argNode := _argNode.(*Nodes_DeclArgNode)
                var symbol *Ast_SymbolInfo
                symbol = argNode.FP.Get_symbolInfo(_env).FP.GetOrg(_env)
                if symbol.FP.Get_hasAccessFromClosure(_env){
                    body.Insert(self.FP.declareVar(_env, symbol, convWasm_lg_2_(_env, index)))
                } else { 
                    self.ctx.FP.Get_sym2loc(_env).Set(symbol,&convWasm_VarLoc__Local{index})
                }
            }
        }
    }
}
// 5870: decl @lune.@base.@convWasm.convFilter.compileFunc
func (self *convWasm_convFilter) compileFunc(_env *LnsEnv, node *Nodes_Node,declInfo *Nodes_DeclFuncInfo,funcType *Ast_TypeInfo,funcIndex LnsInt) *WasmBin_Ins {
    var info *WasmRt_CloInfo
    info = self.FP.getFuncCloInfo(_env, funcType)
    var capSyms *LnsList2_[*Ast_SymbolInfo]
    capSyms = self.FP.getCaptureSyms(_env, funcType)
    var envType LnsInt
    envType = -1
    var cloIns *WasmBin_Ins
    if capSyms.Len() > 0{
        var fields *LnsList2_[*WasmBin_FieldType]
        fields = NewLnsList2_[*WasmBin_FieldType](Lns_2DDDGen[*WasmBin_FieldType](WasmBin_field(_env, WasmBin_typeI32, true), WasmBin_field(_env, WasmBin_refType(_env, info.FP.Get_funcType(_env)), false)))
        var vals *LnsList2_[*WasmBin_Ins]
        vals = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), self.mod.FP.RefFunc(_env, funcIndex)))
        for _, _symbol := range( capSyms.Items ) {
            symbol := _symbol
            var cellType LnsInt
            cellType = self.rt.FP.GetCellType(_env, self.FP.getRep(_env, symbol.FP.Get_typeInfo(_env)))
            fields.Insert(WasmBin_field(_env, WasmBin_refType(_env, cellType), false))
            vals.Insert(self.FP.getCellRef(_env, node, symbol))
        }
        envType = self.mod.FP.AddStructType(_env, info.FP.Get_cloType(_env), fields)
        cloIns = WasmBin_structNew(_env, envType, vals)
    } else { 
        cloIns = WasmBin_structNew(_env, info.FP.Get_cloType(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), self.mod.FP.RefFunc(_env, funcIndex))))
    }
    if self.compiledFuncSet.Has(LnsInt(funcIndex)){
        return cloIns
    }
    self.compiledFuncSet.Add(LnsInt(funcIndex))
    var saved *convWasm_FuncCtx
    saved = self.ctx
    var retTypes *LnsList2_[*Ast_TypeInfo]
    retTypes = self.FP.normType(_env, funcType).FP.Get_retTypeInfoList(_env)
    self.ctx = NewconvWasm_FuncCtx(_env, self.mod.FP.GetFunc(_env, funcIndex), info.FP.Get_results(_env), retTypes)
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    if envType >= 0{
        var envLocal LnsInt
        envLocal = self.ctx.FP.AddLocal(_env, WasmBin_refType(_env, envType))
        self.ctx.FP.Set_envLocal(_env, envLocal)
        self.ctx.FP.Set_envType(_env, envType)
        body.Insert(convWasm_ls_3_(_env, envLocal, WasmBin_refCast(_env, envType, convWasm_lg_2_(_env, 0))))
        for _index, _symbol := range( capSyms.Items ) {
            index := _index + 1
            symbol := _symbol
            self.ctx.FP.Get_sym2loc(_env).Set(symbol,&convWasm_VarLoc__Captured{index + 1})
        }
    }
    self.FP.setupParams(_env, declInfo, funcType, body)
    if self.FP.isRunLuaOnLns(_env, funcType){
        var args LnsInt
        args = self.FP.tempLocal(_env, self.rt.FP.Get_listType(_env))
        var result LnsInt
        result = self.FP.tempLocal(_env, self.rt.FP.Get_listType(_env))
        var mess LnsInt
        mess = self.FP.tempLocal(_env, WasmBin_typeAnyRef)
        if self.embedLuaList.Len() > 0{
            body.Insert(self.FP.registerEmbedLua(_env))
        }
        body.Insert(convWasm_ls_3_(_env, args, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 1))))))
        body.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, args), convWasm_lg_2_(_env, 2)))))
        body.Insert(convWasm_ls_3_(_env, result, WasmBin_call(_env, self.rt.FP.FnLuaCall(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, WasmRt_luaRunLuaOnLns), convWasm_lg_2_(_env, args))))))
        body.Insert(convWasm_ls_3_(_env, mess, WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, result), WasmBin_i64Const(_env, 2))))))
        body.Insert(WasmBin_call(_env, self.rt.FP.FnListGet(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, result), WasmBin_i64Const(_env, 1)))))
        body.Insert(WasmBin_ifElse(_env, WasmBin_refIsNull(_env, convWasm_lg_2_(_env, mess)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.StrLiteral(_env, ""))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refCast(_env, self.rt.FP.Get_tStr(_env), convWasm_lg_2_(_env, mess)))), self.rt.FP.Get_strType(_env)))
    } else if self.FP.isModuleFunc(_env, funcType, "@lune.@base.@Str", "getLineList"){
        body.Insert(WasmBin_call(_env, self.rt.FP.FnStrLineList(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 1)))))
    } else if self.FP.isModuleFunc(_env, funcType, "@lune.@base.@Depend", "getBindLns"){
        var val *WasmBin_Ins
        val = WasmBin_refNull(_env, WasmBin_htAny)
        {
            __forsortCollection0 := self.bindLnsMap
            __forsortSorted0 := __forsortCollection0.CreateKeyListStr()
            __forsortSorted0.Sort( _env, LnsItemKindStr, nil )
            for _, _mod := range( __forsortSorted0.Items ) {
                code := __forsortCollection0.Items[ _mod ]
                mod := _mod
                val = WasmBin_ifElse(_env, WasmBin_call(_env, self.rt.FP.FnStrEq(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 1), self.rt.FP.StrLiteral(_env, mod)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.StrLiteral(_env, code))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)), WasmBin_typeAnyRef)
            }
        }
        body.Insert(self.FP.coerce(_env, val, Ast_builtinTypeStem_, retTypes.GetAt(1)))
    } else if self.FP.isModuleFunc(_env, funcType, "@lune.@base.@Depend", "getFileLastModifiedTime"){
        var mtime LnsInt
        mtime = self.FP.tempLocal(_env, WasmBin_typeF64)
        body.Insert(convWasm_ls_3_(_env, mtime, WasmBin_call(_env, self.rt.FP.FnFileMtime(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 1))))))
        body.Insert(self.FP.coerce(_env, WasmBin_ifElse(_env, WasmBin_plain(_env, WasmBin_opF64Ne, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, mtime), convWasm_lg_2_(_env, mtime)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, WasmBin_htAny))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.BoxF64(_env, convWasm_lg_2_(_env, mtime)))), WasmBin_typeAnyRef), Ast_builtinTypeStem_, retTypes.GetAt(1)))
    } else { 
        {
            _block := declInfo.FP.Get_body(_env)
            if !Lns_IsNil( _block ) {
                block := _block.(*Nodes_BlockNode)
                convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, block))
            }
        }
        if retTypes.Len() > 0{
            body.Insert(WasmBin_unreachable(_env))
        }
    }
    self.ctx.FP.Get_func(_env).FP.SetBody(_env, body)
    self.ctx = saved
    return cloIns
}
// 5960: decl @lune.@base.@convWasm.convFilter.isModuleFunc
func (self *convWasm_convFilter) isModuleFunc(_env *LnsEnv, funcType *Ast_TypeInfo,modName string,name string) bool {
    if funcType.FP.Get_rawTxt(_env) != name{
        return false
    }
    return funcType.FP.GetModule(_env).FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), false) == modName
}
// 5973: decl @lune.@base.@convWasm.convFilter.registerEmbedLua
func (self *convWasm_convFilter) registerEmbedLua(_env *LnsEnv) *WasmBin_Ins {
    if self.embedLuaGlobal < 0{
        self.embedLuaGlobal = self.mod.FP.AddGlobal(_env, WasmBin_typeI32, WasmBin_i32Const(_env, 0))
    }
    var args LnsInt
    args = self.FP.tempLocal(_env, self.rt.FP.Get_listType(_env))
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_globalSet(_env, self.embedLuaGlobal, WasmBin_i32Const(_env, 1))))
    for _, _info := range( self.embedLuaList.Items ) {
        info := _info
        body.Insert(convWasm_ls_3_(_env, args, WasmBin_call(_env, self.rt.FP.FnListNew(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 3))))))
        for _, _txt := range( NewLnsList2_[string](Lns_2DDDGen[string](info.FP.Get_mod(_env), info.FP.Get_code(_env), info.FP.Get_chunkname(_env))).Items ) {
            txt := _txt
            body.Insert(WasmBin_call(_env, self.rt.FP.FnListPush(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, args), self.rt.FP.StrLiteral(_env, txt)))))
        }
        body.Insert(WasmBin_drop(_env, WasmBin_call(_env, self.rt.FP.FnLuaCall(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, WasmRt_luaPreload), convWasm_lg_2_(_env, args))))))
    }
    return WasmBin_ifElse(_env, WasmBin_plain(_env, WasmBin_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_globalGet(_env, self.embedLuaGlobal)))), body, nil, nil)
}
// 5991: decl @lune.@base.@convWasm.convFilter.isRunLuaOnLns
func (self *convWasm_convFilter) isRunLuaOnLns(_env *LnsEnv, funcType *Ast_TypeInfo) bool {
    return self.FP.isModuleFunc(_env, funcType, "@lune.@base.@DependLuaOnLns", "runLuaOnLns")
}
// 5996: decl @lune.@base.@convWasm.convFilter.compileMethod
func (self *convWasm_convFilter) compileMethod(_env *LnsEnv, node *Nodes_Node,declInfo *Nodes_DeclFuncInfo,methodType *Ast_TypeInfo,funcIndex LnsInt,info *convWasm_ClassInfo,retReps *LnsList2_[*WasmBin_ValType]) {
    if self.compiledFuncSet.Has(LnsInt(funcIndex)){
        return 
    }
    self.compiledFuncSet.Add(LnsInt(funcIndex))
    var saved *convWasm_FuncCtx
    saved = self.ctx
    self.ctx = NewconvWasm_FuncCtx(_env, self.mod.FP.GetFunc(_env, funcIndex), retReps, methodType.FP.Get_retTypeInfoList(_env))
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var selfLocal LnsInt
    selfLocal = self.ctx.FP.AddLocal(_env, WasmBin_refType(_env, info.FP.Get_structType(_env)))
    self.ctx.FP.Set_selfLocal(_env, selfLocal)
    self.ctx.FP.Set_selfClass(_env, info)
    body.Insert(convWasm_ls_3_(_env, selfLocal, WasmBin_refCast(_env, info.FP.Get_structType(_env), convWasm_lg_2_(_env, 0))))
    {
        _scope := methodType.FP.Get_scope(_env)
        if !Lns_IsNil( _scope ) {
            scope := _scope.(*Ast_Scope)
            {
                _selfSym := scope.FP.Get_symbol2SymbolInfoMap(_env).Get("self")
                if !Lns_IsNil( _selfSym ) {
                    selfSym := _selfSym.(*Ast_SymbolInfo)
                    var orgSym *Ast_SymbolInfo
                    orgSym = selfSym.FP.GetOrg(_env)
                    if orgSym.FP.Get_hasAccessFromClosure(_env){
                        body.Insert(self.FP.declareVar(_env, orgSym, convWasm_lg_2_(_env, selfLocal)))
                    } else { 
                        self.ctx.FP.Get_sym2loc(_env).Set(orgSym,&convWasm_VarLoc__Local{selfLocal})
                    }
                }
            }
        }
    }
    self.FP.setupParams(_env, declInfo, methodType, body)
    {
        _nativeBody := self.FP.getStrBuilderBody(_env, node, methodType, info, selfLocal)
        if !Lns_IsNil( _nativeBody ) {
            nativeBody := _nativeBody.(*LnsList2_[*WasmBin_Ins])
            convWasm_append_17_(_env, body, nativeBody)
        } else {
            {
                _block := declInfo.FP.Get_body(_env)
                if !Lns_IsNil( _block ) {
                    block := _block.(*Nodes_BlockNode)
                    convWasm_append_17_(_env, body, self.FP.lowerBlock(_env, block))
                }
            }
            if retReps.Len() > 0{
                body.Insert(WasmBin_unreachable(_env))
            }
        }
    }
    self.ctx.FP.Get_func(_env).FP.SetBody(_env, body)
    self.ctx = saved
}
// 6045: decl @lune.@base.@convWasm.convFilter.getStrBuilderBody
func (self *convWasm_convFilter) getStrBuilderBody(_env *LnsEnv, node *Nodes_Node,methodType *Ast_TypeInfo,info *convWasm_ClassInfo,selfLocal LnsInt) LnsAny {
    var classType *Ast_TypeInfo
    classType = methodType.FP.Get_parentInfo(_env)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( classType.FP.Get_rawTxt(_env) != "Builder") ||
        _env.SetStackVal( classType.FP.GetModule(_env).FP.GetFullName(_env, self.FP.Get_typeNameCtrl(_env), self.FP.Get_moduleInfoManager(_env), false) != "@lune.@base.@Str") ).(bool){
        return nil
    }
    var structType LnsInt
    structType = info.FP.Get_structType(_env)
    var txtField LnsInt
    txtField = self.FP.getFieldIndex(_env, node, info, "txt")
    var bufField LnsInt
    bufField = self.FP.getFieldIndex(_env, node, info, "progress")
    var buf *WasmBin_Ins
    buf = WasmBin_refCast(_env, self.rt.FP.Get_tStr(_env), WasmBin_structGet(_env, structType, bufField, convWasm_lg_2_(_env, selfLocal)))
    var convWasm_setBuf func(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins
    convWasm_setBuf = func(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins {
        return WasmBin_structSet(_env, structType, bufField, convWasm_lg_2_(_env, selfLocal), val)
    }
    if _switch0 := methodType.FP.Get_rawTxt(_env); _switch0 == "__init" {
        return NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structSet(_env, structType, txtField, convWasm_lg_2_(_env, selfLocal), self.rt.FP.StrLiteral(_env, "")), convWasm_setBuf(_env, self.rt.FP.NewBuf(_env))))
    } else if _switch0 == "add" {
        return NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_setBuf(_env, WasmBin_call(_env, self.rt.FP.FnBufAdd(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](buf, convWasm_lg_2_(_env, 1)))))))
    } else if _switch0 == "get_txt" {
        return NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.rt.FP.FnBufToStr(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](buf)))))
    } else if _switch0 == "len" {
        return NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.BufLenI64(_env, buf)))
    } else if _switch0 == "clear" {
        return NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.rt.FP.ClearBuf(_env, buf)))
    } else if _switch0 == "flush" {
        return NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    }
    self.FP.error(_env, node, _env.GetVM().String_format("unknown method of Str.Builder -- %s", Lns_2DDD(methodType.FP.Get_rawTxt(_env))))
// insert a dummy
    return nil
}
// 6087: decl @lune.@base.@convWasm.convFilter.getSlotRetReps
func (self *convWasm_convFilter) getSlotRetReps(_env *LnsEnv, info *convWasm_ClassInfo,methodType *Ast_TypeInfo) *LnsList2_[*WasmBin_ValType] {
    {
        _slotIndex := info.FP.Get_name2slot(_env).Get(methodType.FP.Get_rawTxt(_env))
        if !Lns_IsNil( _slotIndex ) {
            slotIndex := _slotIndex.(LnsInt)
            return self.FP.getRepList(_env, info.FP.Get_slotList(_env).GetAt(slotIndex).FP.Get_methodType(_env).FP.Get_retTypeInfoList(_env))
        }
    }
    return self.FP.getRepList(_env, methodType.FP.Get_retTypeInfoList(_env))
}
// 6095: decl @lune.@base.@convWasm.convFilter.compileMethodNode
func (self *convWasm_convFilter) compileMethodNode(_env *LnsEnv, node *Nodes_Node,declInfo *Nodes_DeclFuncInfo) {
    var classType *Ast_TypeInfo
    
    {
        _classType := declInfo.FP.Get_classTypeInfo(_env)
        if _classType == nil{
            self.FP.error(_env, node, "no class")
        } else {
            classType = _classType.(*Ast_TypeInfo)
        }
    }
    var info *convWasm_ClassInfo
    
    {
        _info := self.FP.getClassInfo(_env, classType)
        if _info == nil{
            self.FP.error(_env, node, _env.GetVM().String_format("not found class -- %s", Lns_2DDD(classType.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            info = _info.(*convWasm_ClassInfo)
        }
    }
    var methodType *Ast_TypeInfo
    methodType = node.FP.Get_expType(_env)
    var funcIndex LnsInt
    
    {
        _funcIndex := self.func2index.Get(methodType)
        if _funcIndex == nil{
            if declInfo.FP.Get_kind(_env) == Nodes_FuncKind__Ctor{
                self.FP.compileMethod(_env, node, declInfo, methodType, info.FP.Get_ctorFunc(_env), info, NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
                return 
            }
            self.FP.error(_env, node, _env.GetVM().String_format("not found method -- %s", Lns_2DDD(methodType.FP.GetTxt(_env, nil, nil, nil))))
        } else {
            funcIndex = _funcIndex.(LnsInt)
        }
    }
    if declInfo.FP.Get_staticFlag(_env){
        self.FP.compileFunc(_env, node, declInfo, methodType, funcIndex)
        return 
    }
    if declInfo.FP.Get_kind(_env) == Nodes_FuncKind__Ctor{
        self.FP.compileMethod(_env, node, declInfo, methodType, funcIndex, info, NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
        return 
    }
    self.FP.compileMethod(_env, node, declInfo, methodType, funcIndex, info, self.FP.getRepList(_env, methodType.FP.Get_retTypeInfoList(_env)))
}
// 6124: decl @lune.@base.@convWasm.convFilter.compileDefaultCtor
func (self *convWasm_convFilter) compileDefaultCtor(_env *LnsEnv, node *Nodes_DeclClassNode,info *convWasm_ClassInfo) {
    if self.compiledFuncSet.Has(LnsInt(info.FP.Get_ctorFunc(_env))){
        return 
    }
    self.compiledFuncSet.Add(LnsInt(info.FP.Get_ctorFunc(_env)))
    var ctorType *Ast_TypeInfo
    ctorType = self.FP.getCtorType(_env, info.FP.Get_typeInfo(_env))
    var _func *WasmBin_Func
    _func = self.mod.FP.GetFunc(_env, info.FP.Get_ctorFunc(_env))
    var selfLocal LnsInt
    selfLocal = _func.FP.AddLocal(_env, WasmBin_refType(_env, info.FP.Get_structType(_env)))
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, selfLocal, WasmBin_refCast(_env, info.FP.Get_structType(_env), convWasm_lg_2_(_env, 0)))))
    var superArgNum LnsInt
    superArgNum = 0
    {
        _parent := info.FP.Get_parent(_env)
        if !Lns_IsNil( _parent ) {
            parent := _parent.(*convWasm_ClassInfo)
            var parentCtorType *Ast_TypeInfo
            parentCtorType = self.FP.getCtorType(_env, parent.FP.Get_typeInfo(_env))
            var args *LnsList2_[*WasmBin_Ins]
            args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, selfLocal)))
            for _index, _argType := range( parentCtorType.FP.Get_argTypeInfoList(_env).Items ) {
                index := _index + 1
                argType := _argType
                if node.FP.Get_hasOldCtor(_env){
                    args.Insert(self.FP.getDefaultVal(_env, self.FP.getRep(_env, argType)))
                } else { 
                    args.Insert(convWasm_lg_2_(_env, index))
                }
            }
            body.Insert(WasmBin_call(_env, parent.FP.Get_ctorFunc(_env), args))
            if Lns_op_not(node.FP.Get_hasOldCtor(_env)){
                superArgNum = parentCtorType.FP.Get_argTypeInfoList(_env).Len()
            }
        }
    }
    var argIndex LnsInt
    argIndex = superArgNum + 1
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode
        if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
            if argIndex <= ctorType.FP.Get_argTypeInfoList(_env).Len(){
                var fieldIndex LnsInt
                fieldIndex = self.FP.getFieldIndex(_env, &node.Nodes_Node, info, memberNode.FP.Get_name(_env).Txt)
                body.Insert(WasmBin_structSet(_env, info.FP.Get_structType(_env), fieldIndex, convWasm_lg_2_(_env, selfLocal), self.FP.coerce(_env, convWasm_lg_2_(_env, argIndex), ctorType.FP.Get_argTypeInfoList(_env).GetAt(argIndex), memberNode.FP.Get_symbolInfo(_env).FP.Get_typeInfo(_env))))
            }
            argIndex = argIndex + 1
        }
    }
    _func.FP.SetBody(_env, body)
}
// 6166: decl @lune.@base.@convWasm.convFilter.compileAccessors
func (self *convWasm_convFilter) compileAccessors(_env *LnsEnv, node *Nodes_DeclClassNode,info *convWasm_ClassInfo) {
    var scope *Ast_Scope
    scope = Lns_unwrap( info.FP.Get_typeInfo(_env).FP.Get_scope(_env)).(*Ast_Scope)
    for _, _memberNode := range( node.FP.Get_memberList(_env).Items ) {
        memberNode := _memberNode
        var member string
        member = memberNode.FP.Get_name(_env).Txt
        if Lns_op_not(memberNode.FP.Get_staticFlag(_env)){
            for _, _name := range( NewLnsList2_[string](Lns_2DDDGen[string]("get_" + member, "set_" + member)).Items ) {
                name := _name
                {
                    _symbol := scope.FP.Get_symbol2SymbolInfoMap(_env).Get(name)
                    if !Lns_IsNil( _symbol ) {
                        symbol := _symbol.(*Ast_SymbolInfo)
                        var methodType *Ast_TypeInfo
                        methodType = symbol.FP.Get_typeInfo(_env)
                        if Lns_op_not(self.declaredMethodSet.Has(methodType)){
                            {
                                _funcIndex := self.func2index.Get(methodType)
                                if !Lns_IsNil( _funcIndex ) {
                                    funcIndex := _funcIndex.(LnsInt)
                                    if Lns_op_not(self.compiledFuncSet.Has(LnsInt(funcIndex))){
                                        self.compiledFuncSet.Add(LnsInt(funcIndex))
                                        var _func *WasmBin_Func
                                        _func = self.mod.FP.GetFunc(_env, funcIndex)
                                        var selfIns *WasmBin_Ins
                                        selfIns = WasmBin_refCast(_env, info.FP.Get_structType(_env), convWasm_lg_2_(_env, 0))
                                        var fieldIndex LnsInt
                                        fieldIndex = self.FP.getFieldIndex(_env, &node.Nodes_Node, info, member)
                                        if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(name,"^get_", nil, nil))){
                                            var retReps *LnsList2_[*WasmBin_ValType]
                                            retReps = self.FP.getRepList(_env, methodType.FP.Get_retTypeInfoList(_env))
                                            _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.coerceRep(_env, WasmBin_structGet(_env, info.FP.Get_structType(_env), fieldIndex, selfIns), info.FP.Get_fieldReps(_env).GetAt(fieldIndex - 1), retReps.GetAt(1)))))
                                        } else { 
                                            _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_structSet(_env, info.FP.Get_structType(_env), fieldIndex, selfIns, self.FP.coerceRep(_env, convWasm_lg_2_(_env, 1), self.FP.getRep(_env, methodType.FP.Get_argTypeInfoList(_env).GetAt(1)), info.FP.Get_fieldReps(_env).GetAt(fieldIndex - 1))))))
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
// 6207: decl @lune.@base.@convWasm.convFilter.compileAdvertise
func (self *convWasm_convFilter) compileAdvertise(_env *LnsEnv, node *Nodes_DeclClassNode,info *convWasm_ClassInfo) {
    for _, _adv := range( info.FP.Get_advertiseList(_env).Items ) {
        adv := _adv
        if self.compiledFuncSet.Has(LnsInt(adv.FP.Get_impl(_env))){
            continue
        }
        self.compiledFuncSet.Add(LnsInt(adv.FP.Get_impl(_env)))
        var _func *WasmBin_Func
        _func = self.mod.FP.GetFunc(_env, adv.FP.Get_impl(_env))
        var fieldIndex LnsInt
        fieldIndex = self.FP.getFieldIndex(_env, &node.Nodes_Node, info, adv.FP.Get_memberName(_env))
        var member *WasmBin_Ins
        member = WasmBin_structGet(_env, info.FP.Get_structType(_env), fieldIndex, WasmBin_refCast(_env, info.FP.Get_structType(_env), convWasm_lg_2_(_env, 0)))
        var argReps *LnsList2_[*WasmBin_ValType]
        argReps = self.FP.getRepList(_env, adv.FP.Get_methodType(_env).FP.Get_argTypeInfoList(_env))
        var retReps *LnsList2_[*WasmBin_ValType]
        retReps = self.FP.getSlotRetReps(_env, info, adv.FP.Get_methodType(_env))
        var ins *WasmBin_Ins
        var calleeRetReps *LnsList2_[*WasmBin_ValType]
        ins,_,calleeRetReps = self.FP.callMethodDyn(_env, &node.Nodes_Node, member, info.FP.Get_fieldReps(_env).GetAt(fieldIndex - 1), adv.FP.Get_memberType(_env), adv.FP.Get_methodName(_env), convWasm_ArgGetter_19_(func(_env *LnsEnv, index LnsInt,rep *WasmBin_ValType) *WasmBin_Ins {
            return self.FP.coerceRep(_env, convWasm_lg_2_(_env, index), argReps.GetAt(index), rep)
        }), _func)
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](ins))
        if calleeRetReps.Len() > 0{
            var locals *LnsList2_[LnsInt]
            locals = NewLnsList2_[LnsInt]([]LnsInt{})
            for _, _rep := range( calleeRetReps.Items ) {
                rep := _rep
                locals.Insert(LnsInt(_func.FP.AddLocal(_env, rep)))
            }
            {
                var _forFrom0 LnsInt = locals.Len()
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
                    body.Insert(convWasm_ls_3_(_env, locals.GetAt(index), WasmBin_nop(_env)))
                    _forWork0 += _forDelta0
                }
            }
            for _index, _rep := range( retReps.Items ) {
                index := _index + 1
                rep := _rep
                body.Insert(self.FP.coerceRep(_env, convWasm_lg_2_(_env, locals.GetAt(index)), calleeRetReps.GetAt(index), rep))
            }
        }
        _func.FP.SetBody(_env, body)
    }
}
// 6247: decl @lune.@base.@convWasm.convFilter.callMethodDyn
func (self *convWasm_convFilter) callMethodDyn(_env *LnsEnv, node *Nodes_Node,obj *WasmBin_Ins,objRep *WasmBin_ValType,objType *Ast_TypeInfo,name string,getArg convWasm_ArgGetter_19_,_func *WasmBin_Func)(*WasmBin_Ins, *LnsList2_[*WasmBin_ValType], *LnsList2_[*WasmBin_ValType]) {
    var work *Ast_TypeInfo
    work = self.FP.getObjType(_env, objType)
    var objLocal LnsInt
    objLocal = _func.FP.AddLocal(_env, self.rt.FP.Get_objType(_env))
    var pre *LnsList2_[*WasmBin_Ins]
    pre = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_ls_3_(_env, objLocal, self.FP.coerceRep(_env, obj, objRep, self.rt.FP.Get_objType(_env)))))
    var methodType *Ast_TypeInfo
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var callFuncType LnsInt
    var funcRef *WasmBin_Ins
    if work.FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
        var ifInfo *convWasm_IfInfo
        ifInfo = self.FP.getIfInfo(_env, work)
        var methodIndex LnsInt
        
        {
            _methodIndex := ifInfo.FP.Get_name2index(_env).Get(name)
            if _methodIndex == nil{
                self.FP.error(_env, node, _env.GetVM().String_format("not found method -- %s", Lns_2DDD(name)))
            } else {
                methodIndex = _methodIndex.(LnsInt)
            }
        }
        methodType = Lns_unwrap( self.FP.getIfMethodType(_env, work, name)).(*Ast_TypeInfo)
        args.Insert(convWasm_lg_2_(_env, objLocal))
        callFuncType = ifInfo.FP.Get_methodFuncTypes(_env).GetAt(methodIndex + 1)
        var itab *WasmBin_Ins
        itab = WasmBin_refCast(_env, ifInfo.FP.Get_itType(_env), WasmBin_arrayGet(_env, self.rt.FP.Get_tAnyArr(_env), WasmBin_structGet(_env, self.rt.FP.Get_tVT(_env), WasmRt_fVtItabs, WasmBin_structGet(_env, self.rt.FP.Get_tObj(_env), WasmRt_fObjVt, convWasm_lg_2_(_env, objLocal))), WasmBin_i32Const(_env, ifInfo.FP.Get_ifId(_env))))
        funcRef = WasmBin_structGet(_env, ifInfo.FP.Get_itType(_env), methodIndex, itab)
    } else { 
        var info *convWasm_ClassInfo
        
        {
            _info := self.FP.getClassInfo(_env, work)
            if _info == nil{
                self.FP.error(_env, node, _env.GetVM().String_format("not found class -- %s", Lns_2DDD(work.FP.GetTxt(_env, nil, nil, nil))))
            } else {
                info = _info.(*convWasm_ClassInfo)
            }
        }
        var slotIndex LnsInt
        
        {
            _slotIndex := info.FP.Get_name2slot(_env).Get(name)
            if _slotIndex == nil{
                self.FP.error(_env, node, _env.GetVM().String_format("not found method -- %s", Lns_2DDD(name)))
            } else {
                slotIndex = _slotIndex.(LnsInt)
            }
        }
        var slot *convWasm_SlotInfo
        slot = info.FP.Get_slotList(_env).GetAt(slotIndex)
        methodType = slot.FP.Get_methodType(_env)
        var objWork LnsInt
        objWork = _func.FP.AddLocal(_env, WasmBin_refType(_env, info.FP.Get_structType(_env)))
        pre.Insert(convWasm_ls_3_(_env, objWork, WasmBin_refCast(_env, info.FP.Get_structType(_env), convWasm_lg_2_(_env, objLocal))))
        args.Insert(convWasm_lg_2_(_env, objWork))
        callFuncType = slot.FP.Get_funcType(_env)
        funcRef = WasmBin_structGet(_env, info.FP.Get_vtType(_env), slotIndex, WasmBin_structGet(_env, info.FP.Get_structType(_env), WasmRt_fObjVt, convWasm_lg_2_(_env, objWork)))
    }
    var argReps *LnsList2_[*WasmBin_ValType]
    argReps = self.FP.getRepList(_env, methodType.FP.Get_argTypeInfoList(_env))
    for _index, _rep := range( argReps.Items ) {
        index := _index + 1
        rep := _rep
        args.Insert(getArg(_env, index, rep))
    }
    args.Insert(funcRef)
    pre.Insert(WasmBin_callRef(_env, callFuncType, args))
    return WasmBin_seq(_env, pre), argReps, self.FP.getRepList(_env, methodType.FP.Get_retTypeInfoList(_env))
}
// 6298: decl @lune.@base.@convWasm.convFilter.lowerDeclClass
func (self *convWasm_convFilter) lowerDeclClass(_env *LnsEnv, node *Nodes_DeclClassNode,out *LnsList2_[*WasmBin_Ins]) {
    if node.FP.IsModule(_env){
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( node.FP.Get_lazyLoad(_env) == Nodes_LazyLoad__Off) &&
            _env.SetStackVal( Lns_op_not(node.FP.Get_gluePrefix(_env))) ).(bool)){
            {
                _modType := self.FP.getExtModuleType(_env, node.FP.Get_expType(_env))
                if !Lns_IsNil( _modType ) {
                    modType := _modType.(*Ast_TypeInfo)
                    out.Insert(WasmBin_drop(_env, self.FP.lowerExtModuleVal(_env, &node.Nodes_Node, modType)))
                }
            }
        }
        return 
    }
    if node.FP.Get_expType(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__IF{
        self.FP.getIfInfo(_env, node.FP.Get_expType(_env))
        return 
    }
    var info *convWasm_ClassInfo
    info = self.FP.registerClass(_env, node)
    for _, _field := range( node.FP.Get_fieldList(_env).Items ) {
        field := _field
        if _switch0 := field.FP.Get_kind(_env); _switch0 == Nodes_NodeKind_get_DeclMethod(_env) {
            self.FP.compileMethodNode(_env, field, (Lns_unwrap( Nodes_DeclMethodNodeDownCastF(field.FP)).(*Nodes_DeclMethodNode)).FP.Get_declInfo(_env))
        } else if _switch0 == Nodes_NodeKind_get_DeclConstr(_env) {
            self.FP.compileMethodNode(_env, field, (Lns_unwrap( Nodes_DeclConstrNodeDownCastF(field.FP)).(*Nodes_DeclConstrNode)).FP.Get_declInfo(_env))
        }
    }
    if Lns_op_not(node.FP.HasUserInit(_env)){
        self.FP.compileDefaultCtor(_env, node, info)
    }
    self.FP.compileAccessors(_env, node, info)
    self.FP.compileAdvertise(_env, node, info)
    {
        _initFunc := node.FP.Get_initBlock(_env).FP.Get_func(_env)
        if !Lns_IsNil( _initFunc ) {
            initFunc := _initFunc.(*Nodes_DeclMethodNode)
            {
                _block := initFunc.FP.Get_declInfo(_env).FP.Get_body(_env)
                if !Lns_IsNil( _block ) {
                    block := _block.(*Nodes_BlockNode)
                    convWasm_append_17_(_env, out, self.FP.lowerBlock(_env, block))
                }
            }
        }
    }
}
// 6342: decl @lune.@base.@convWasm.convFilter.makeIfAdapter
func (self *convWasm_convFilter) makeIfAdapter(_env *LnsEnv, ifMethodType *Ast_TypeInfo,implMethodType *Ast_TypeInfo,impl LnsInt) LnsInt {
    var ifArgReps *LnsList2_[*WasmBin_ValType]
    ifArgReps = self.FP.getRepList(_env, ifMethodType.FP.Get_argTypeInfoList(_env))
    var ifRetReps *LnsList2_[*WasmBin_ValType]
    ifRetReps = self.FP.getRepList(_env, ifMethodType.FP.Get_retTypeInfoList(_env))
    var implArgReps *LnsList2_[*WasmBin_ValType]
    implArgReps = self.FP.getRepList(_env, implMethodType.FP.Get_argTypeInfoList(_env))
    var implRetReps *LnsList2_[*WasmBin_ValType]
    implRetReps = self.FP.getRepList(_env, implMethodType.FP.Get_retTypeInfoList(_env))
    var params *LnsList2_[*WasmBin_ValType]
    params = NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_objType(_env)))
    for _, _rep := range( ifArgReps.Items ) {
        rep := _rep
        params.Insert(rep)
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, params, ifRetReps)
    var args *LnsList2_[*WasmBin_Ins]
    args = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_lg_2_(_env, 0)))
    for _argIndex, _rep := range( implArgReps.Items ) {
        argIndex := _argIndex + 1
        rep := _rep
        args.Insert(self.FP.coerceRep(_env, convWasm_lg_2_(_env, argIndex), ifArgReps.GetAt(argIndex), rep))
    }
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, impl, args)))
    if implRetReps.Len() > 0{
        var locals *LnsList2_[LnsInt]
        locals = NewLnsList2_[LnsInt]([]LnsInt{})
        for _, _rep := range( implRetReps.Items ) {
            rep := _rep
            locals.Insert(LnsInt(_func.FP.AddLocal(_env, rep)))
        }
        {
            var _forFrom0 LnsInt = locals.Len()
            var _forTo0 LnsInt = 1
            _forWork0 := _forFrom0
            _forDelta0 := -1
            for {
                if _forDelta0 > 0 {
                   if _forWork0 > _forTo0 { break }
                } else {
                   if _forWork0 < _forTo0 { break }
                }
                retIndex := _forWork0
                body.Insert(convWasm_ls_3_(_env, locals.GetAt(retIndex), WasmBin_nop(_env)))
                _forWork0 += _forDelta0
            }
        }
        for _retIndex, _rep := range( ifRetReps.Items ) {
            retIndex := _retIndex + 1
            rep := _rep
            body.Insert(self.FP.coerceRep(_env, convWasm_lg_2_(_env, locals.GetAt(retIndex)), implRetReps.GetAt(retIndex), rep))
        }
    }
    _func.FP.SetBody(_env, body)
    return index
}
// 6377: decl @lune.@base.@convWasm.convFilter.fnToMapDyn
func (self *convWasm_convFilter) fnToMapDyn(_env *LnsEnv) LnsInt {
    if self.toMapDynIndex < 0{
        var index LnsInt
        index,_ = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeAnyRef)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.rt.FP.Get_mapType(_env))))
        self.mod.FP.SetFuncName(_env, index, "toMapDyn")
        self.toMapDynIndex = index
    }
    return self.toMapDynIndex
}
// 6405: decl @lune.@base.@convWasm.convFilter.buildToMapDyn
func (self *convWasm_convFilter) buildToMapDyn(_env *LnsEnv) {
    var node *Nodes_Node
    
    {
        _node := self.toMapDynNode
        if _node == nil{
            return 
        } else {
            node = _node.(*Nodes_Node)
        }
    }
    var _func *WasmBin_Func
    _func = self.mod.FP.GetFunc(_env, self.toMapDynIndex)
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    var num LnsInt
    num = -1
    for num != self.classInfoList.Len() {
        num = self.classInfoList.Len()
        var infoList *LnsList2_[*convWasm_ClassInfo]
        infoList = NewLnsList2_[*convWasm_ClassInfo]([]*convWasm_ClassInfo{})
        for _, _info := range( self.classInfoList.Items ) {
            info := _info
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( info.FP.Get_typeInfo(_env).FP.Get_kind(_env) == Ast_TypeInfoKind__Class) &&
                _env.SetStackVal( info.FP.Get_typeInfo(_env).FP.IsInheritFrom(_env, self.processInfo, Ast_builtinTypeMapping, nil)) ).(bool)){
                infoList.Insert(info)
            }
        }
        infoList.Sort(_env, LnsItemKindStem, LnsComp(func ( _env *LnsEnv, val1, val2 LnsAny ) bool {return convWasm_convFilter_buildToMapDyn___anonymous_0_( _env, val1.(convWasm_ClassInfoDownCast).ToconvWasm_ClassInfo(), val2.(convWasm_ClassInfoDownCast).ToconvWasm_ClassInfo())}))
        body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        for _, _info := range( infoList.Items ) {
            info := _info
            body.Insert(WasmBin_ifElse(_env, WasmBin_refTest(_env, info.FP.Get_structType(_env), convWasm_lg_2_(_env, 0)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.FP.fnToMap(_env, node, info), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refCast(_env, info.FP.Get_structType(_env), convWasm_lg_2_(_env, 0)))))))))), nil, nil))
        }
    }
    body.Insert(self.rt.FP.Error(_env, self.rt.FP.StrLiteral(_env, "_toMap: not Mapping object")))
    _func.FP.SetBody(_env, body)
}
// 6446: decl @lune.@base.@convWasm.convFilter.finalizeClasses
func (self *convWasm_convFilter) finalizeClasses(_env *LnsEnv) {
    self.FP.buildToMapDyn(_env)
    for _, _info := range( self.classInfoList.Items ) {
        info := _info
        for _, _ifType := range( info.FP.Get_ifList(_env).Items ) {
            ifType := _ifType
            self.FP.getIfInfo(_env, ifType)
        }
    }
    for _, _info := range( self.classInfoList.Items ) {
        info := _info
        var itabs *LnsList2_[*WasmBin_Ins]
        itabs = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        for _, _ifInfo := range( self.ifInfoList.Items ) {
            ifInfo := _ifInfo
            var implemented bool
            implemented = false
            for _, _ifType := range( info.FP.Get_ifList(_env).Items ) {
                ifType := _ifType
                if ifType == ifInfo.FP.Get_typeInfo(_env){
                    implemented = true
                }
            }
            if implemented{
                var funcs *LnsList2_[*WasmBin_Ins]
                funcs = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
                for _index, _name := range( ifInfo.FP.Get_methodNames(_env).Items ) {
                    index := _index + 1
                    name := _name
                    var funcType LnsInt
                    funcType = ifInfo.FP.Get_methodFuncTypes(_env).GetAt(index)
                    var ins *WasmBin_Ins
                    ins = WasmBin_refNull(_env, funcType)
                    {
                        _slotIndex := info.FP.Get_name2slot(_env).Get(name)
                        if !Lns_IsNil( _slotIndex ) {
                            slotIndex := _slotIndex.(LnsInt)
                            var slot *convWasm_SlotInfo
                            slot = info.FP.Get_slotList(_env).GetAt(slotIndex)
                            if slot.FP.Get_impl(_env) >= 0{
                                if slot.FP.Get_funcType(_env) != funcType{
                                    var ifMethodType *Ast_TypeInfo
                                    ifMethodType = Lns_unwrap( self.FP.getIfMethodType(_env, ifInfo.FP.Get_typeInfo(_env), name)).(*Ast_TypeInfo)
                                    ins = self.mod.FP.RefFunc(_env, self.FP.makeIfAdapter(_env, ifMethodType, slot.FP.Get_methodType(_env), slot.FP.Get_impl(_env)))
                                } else { 
                                    ins = self.mod.FP.RefFunc(_env, slot.FP.Get_impl(_env))
                                }
                            }
                        }
                    }
                    funcs.Insert(ins)
                }
                itabs.Insert(WasmBin_structNew(_env, ifInfo.FP.Get_itType(_env), funcs))
            } else { 
                itabs.Insert(WasmBin_refNull(_env, WasmBin_htAny))
            }
        }
        var vals *LnsList2_[*WasmBin_Ins]
        vals = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_arrayNewFixed(_env, self.rt.FP.Get_tAnyArr(_env), itabs)))
        for _, _slot := range( info.FP.Get_slotList(_env).Items ) {
            slot := _slot
            if slot.FP.Get_impl(_env) >= 0{
                vals.Insert(self.mod.FP.RefFunc(_env, slot.FP.Get_impl(_env)))
            } else { 
                vals.Insert(WasmBin_refNull(_env, slot.FP.Get_funcType(_env)))
            }
        }
        self.prologue.Insert(WasmBin_globalSet(_env, info.FP.Get_vtGlobal(_env), WasmBin_structNew(_env, info.FP.Get_vtType(_env), vals)))
    }
}
// 6501: decl @lune.@base.@convWasm.convFilter.collectDecl
func (self *convWasm_convFilter) collectDecl(_env *LnsEnv, node *Nodes_Node) {
    var alreadySet *LnsSet2_[*Nodes_Node]
    alreadySet = NewLnsSet2_[*Nodes_Node]([]*Nodes_Node{})
    node.FP.Visit(_env, Nodes_NodeVisitor(func(_env *LnsEnv, workNode *Nodes_Node,parent *Nodes_Node,relation string,depth LnsInt) LnsInt {
        {
            _classNode := Nodes_DeclClassNodeDownCastF(workNode.FP)
            if !Lns_IsNil( _classNode ) {
                classNode := _classNode.(*Nodes_DeclClassNode)
                self.classType2decl.Set(self.FP.getGenSrc(_env, classNode.FP.Get_expType(_env)),classNode)
            }
        }
        {
            _algeNode := Nodes_DeclAlgeNodeDownCastF(workNode.FP)
            if !Lns_IsNil( _algeNode ) {
                algeNode := _algeNode.(*Nodes_DeclAlgeNode)
                self.algeType2decl.Set(&algeNode.FP.Get_algeType(_env).Ast_TypeInfo,algeNode)
            }
        }
        {
            _enumNode := Nodes_DeclEnumNodeDownCastF(workNode.FP)
            if !Lns_IsNil( _enumNode ) {
                enumNode := _enumNode.(*Nodes_DeclEnumNode)
                self.enumType2decl.Set(&enumNode.FP.Get_enumType(_env).Ast_TypeInfo,enumNode)
            }
        }
        if _switch0 := workNode.FP.Get_kind(_env); _switch0 == Nodes_NodeKind_get_DeclMethod(_env) || _switch0 == Nodes_NodeKind_get_ProtoMethod(_env) || _switch0 == Nodes_NodeKind_get_DeclConstr(_env) {
            self.declaredMethodSet.Add(workNode.FP.Get_expType(_env))
        }
        return Nodes_NodeVisitMode__Child
    }), 0, alreadySet)
}
// 6536: decl @lune.@base.@convWasm.convFilter.registerModuleVars
func (self *convWasm_convFilter) registerModuleVars(_env *LnsEnv, node *Nodes_RootNode) {
    self.FP.registerModuleVarsIn(_env, node.FP.Get_children(_env))
}
// 6539: decl @lune.@base.@convWasm.convFilter.registerModuleVarsIn
func (self *convWasm_convFilter) registerModuleVarsIn(_env *LnsEnv, stmtList *LnsList2_[*Nodes_Node]) {
    for _, _child := range( stmtList.Items ) {
        child := _child
        {
            _macroNode := convWasm_getMacroExp_21_(_env, child)
            if !Lns_IsNil( _macroNode ) {
                macroNode := _macroNode.(*Nodes_ExpMacroExpNode)
                self.FP.registerModuleVarsIn(_env, macroNode.FP.Get_stmtList(_env))
            }
        }
        {
            _declNode := Nodes_DeclVarNodeDownCastF(child.FP)
            if !Lns_IsNil( _declNode ) {
                declNode := _declNode.(*Nodes_DeclVarNode)
                for _, _symbol := range( declNode.FP.Get_symbolInfoList(_env).Items ) {
                    symbol := _symbol
                    var orgSym *Ast_SymbolInfo
                    orgSym = symbol.FP.GetOrg(_env)
                    if self.FP.isGlobalSym(_env, orgSym){
                        var rep *WasmBin_ValType
                        rep = self.FP.getRep(_env, orgSym.FP.Get_typeInfo(_env))
                        self.globalSym2loc.Set(orgSym,&convWasm_VarLoc__Global{self.mod.FP.AddGlobal(_env, rep, self.FP.getDefaultVal(_env, rep))})
                    }
                }
            }
        }
    }
}
// 6561: decl @lune.@base.@convWasm.convFilter.registerImportAlias
func (self *convWasm_convFilter) registerImportAlias(_env *LnsEnv, node *Nodes_RootNode) {
    for _, _child := range( node.FP.Get_children(_env).Items ) {
        child := _child
        {
            _importNode := Nodes_ImportNodeDownCastF(child.FP)
            if !Lns_IsNil( _importNode ) {
                importNode := _importNode.(*Nodes_ImportNode)
                var info *Nodes_ImportInfo
                info = importNode.FP.Get_info(_env)
                {
                    _provideSym := self.module2provide.Get(info.FP.Get_modulePath(_env))
                    if !Lns_IsNil( _provideSym ) {
                        provideSym := _provideSym.(*Ast_SymbolInfo)
                        self.symAlias.Set(info.FP.Get_symbolInfo(_env).FP.GetOrg(_env),provideSym)
                    }
                }
                if info.FP.Get_lazy(_env) != Nodes_LazyLoad__Off{
                    self.lazySym2module.Set(info.FP.Get_symbolInfo(_env).FP.GetOrg(_env),info.FP.Get_modulePath(_env))
                }
            }
        }
    }
}
// 6580: decl @lune.@base.@convWasm.convFilter.getLazyInit
func (self *convWasm_convFilter) getLazyInit(_env *LnsEnv, node *Nodes_Node) LnsAny {
    var refNode *Nodes_ExpRefNode
    
    {
        _refNode := Nodes_ExpRefNodeDownCastF(node.FP)
        if _refNode == nil{
            return nil
        } else {
            refNode = _refNode.(*Nodes_ExpRefNode)
        }
    }
    var modPath string
    
    {
        _modPath := self.lazySym2module.Get(refNode.FP.Get_symbolInfo(_env).FP.GetOrg(_env))
        if _modPath == nil{
            return nil
        } else {
            modPath = _modPath.(string)
        }
    }
    {
        _initIndex := self.module2init.Get(modPath)
        if !Lns_IsNil( _initIndex ) {
            initIndex := _initIndex.(LnsInt)
            return WasmBin_call(_env, initIndex, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
        }
    }
    return nil
}
// 6594: decl @lune.@base.@convWasm.convFilter.resolveSym
func (self *convWasm_convFilter) resolveSym(_env *LnsEnv, symbol *Ast_SymbolInfo) *Ast_SymbolInfo {
    {
        _provideSym := self.symAlias.Get(symbol)
        if !Lns_IsNil( _provideSym ) {
            provideSym := _provideSym.(*Ast_SymbolInfo)
            return provideSym
        }
    }
    return symbol
}
// 6607: decl @lune.@base.@convWasm.convFilter.lowerImportedModules
func (self *convWasm_convFilter) lowerImportedModules(_env *LnsEnv) {
    var mainScope *Ast_Scope
    mainScope = self.moduleScope
    var mainType *Ast_TypeInfo
    mainType = self.moduleTypeInfo
    for _, _ast := range( self.importAstList.Items ) {
        ast := _ast
        var root *Nodes_RootNode
        root = Lns_unwrap( Nodes_RootNodeDownCastF(ast.FP.Get_node(_env).FP)).(*Nodes_RootNode)
        var moduleType *Ast_TypeInfo
        moduleType = ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env)
        self.moduleScope = Lns_unwrap( moduleType.FP.Get_scope(_env)).(*Ast_Scope)
        self.moduleTypeInfo = moduleType
        self.moduleScopeSet.Add(self.moduleScope)
        self.FP.collectDecl(_env, &root.Nodes_Node)
        self.FP.registerModuleVars(_env, root)
        self.FP.registerImportAlias(_env, root)
        {
            _provideNode := Nodes_ProvideNodeDownCastF(root.FP.Get_provideNode(_env))
            if !Lns_IsNil( _provideNode ) {
                provideNode := _provideNode.(*Nodes_ProvideNode)
                self.module2provide.Set(ast.FP.Get_exportInfo(_env).FP.Get_fullName(_env),provideNode.FP.Get_symbol(_env).FP.GetOrg(_env))
            }
        }
        var initIndex LnsInt
        var initFunc *WasmBin_Func
        initIndex,initFunc = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
        self.mod.FP.SetFuncName(_env, initIndex, "init:" + ast.FP.Get_exportInfo(_env).FP.Get_fullName(_env))
        self.module2init.Set(ast.FP.Get_exportInfo(_env).FP.Get_fullName(_env),initIndex)
        var doneGlobal LnsInt
        doneGlobal = self.mod.FP.AddGlobal(_env, WasmBin_typeI32, WasmBin_i32Const(_env, 0))
        var saved *convWasm_FuncCtx
        saved = self.ctx
        self.ctx = NewconvWasm_FuncCtx(_env, initFunc, NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}), NewLnsList2_[*Ast_TypeInfo]([]*Ast_TypeInfo{}))
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, WasmBin_globalGet(_env, doneGlobal), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})))), nil, nil), WasmBin_globalSet(_env, doneGlobal, WasmBin_i32Const(_env, 1))))
        convWasm_append_17_(_env, body, self.FP.lowerStmtList(_env, root.FP.Get_children(_env)))
        initFunc.FP.SetBody(_env, body)
        self.ctx = saved
    }
    self.moduleScope = mainScope
    self.moduleTypeInfo = mainType
}
// 6642: decl @lune.@base.@convWasm.convFilter.collectDynGetters
func (self *convWasm_convFilter) collectDynGetters(_env *LnsEnv, node *Nodes_Node) {
    var alreadySet *LnsSet2_[*Nodes_Node]
    alreadySet = NewLnsSet2_[*Nodes_Node]([]*Nodes_Node{})
    node.FP.Visit(_env, Nodes_NodeVisitor(func(_env *LnsEnv, workNode *Nodes_Node,parent *Nodes_Node,relation string,depth LnsInt) LnsInt {
        {
            _methodNode := Nodes_DeclMethodNodeDownCastF(workNode.FP)
            if !Lns_IsNil( _methodNode ) {
                methodNode := _methodNode.(*Nodes_DeclMethodNode)
                var name string
                name = methodNode.FP.Get_expType(_env).FP.Get_rawTxt(_env)
                if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(name,"^get_", nil, nil))){
                    self.dynGetterNameSet.Add(name)
                }
            }
        }
        {
            _classNode := Nodes_DeclClassNodeDownCastF(workNode.FP)
            if !Lns_IsNil( _classNode ) {
                classNode := _classNode.(*Nodes_DeclClassNode)
                for _, _advertiseInfo := range( classNode.FP.Get_advertiseList(_env).Items ) {
                    advertiseInfo := _advertiseInfo
                    var memberType *Ast_TypeInfo
                    memberType = self.FP.getGenSrc(_env, advertiseInfo.FP.Get_member(_env).FP.Get_expType(_env))
                    for _, _mtdName := range( Ast_getAllMethodName(_env, memberType, Ast_MethodKind__Object).FP.Get_list(_env).Items ) {
                        mtdName := _mtdName.(string)
                        var name string
                        name = advertiseInfo.FP.Get_prefix(_env) + mtdName
                        if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(name,"^get_", nil, nil))){
                            self.dynGetterNameSet.Add(name)
                        }
                    }
                }
            }
        }
        return Nodes_NodeVisitMode__Child
    }), 0, alreadySet)
}
// 6669: decl @lune.@base.@convWasm.convFilter.processRoot
func (self *convWasm_convFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    for _, _ast := range( self.importAstList.Items ) {
        ast := _ast
        self.FP.collectDynGetters(_env, ast.FP.Get_node(_env))
    }
    self.FP.collectDynGetters(_env, &node.Nodes_Node)
    self.FP.lowerImportedModules(_env)
    self.FP.collectDecl(_env, &node.Nodes_Node)
    self.FP.registerModuleVars(_env, node)
    self.FP.registerImportAlias(_env, node)
    var mainBody *LnsList2_[*WasmBin_Ins]
    mainBody = self.FP.lowerStmtList(_env, node.FP.Get_children(_env))
    self.FP.finalizeClasses(_env)
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    convWasm_append_17_(_env, body, self.prologue)
    convWasm_append_17_(_env, body, mainBody)
    {
        _mainSym := self.moduleScope.FP.Get_symbol2SymbolInfoMap(_env).Get("__main")
        if !Lns_IsNil( _mainSym ) {
            mainSym := _mainSym.(*Ast_SymbolInfo)
            if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( mainSym.FP.Get_kind(_env) == Ast_SymbolKind__Fun) &&
                _env.SetStackVal( mainSym.FP.Get_accessMode(_env) == Ast_AccessMode__Pub) ).(bool)){
                {
                    _funcIndex := self.func2index.Get(mainSym.FP.Get_typeInfo(_env))
                    if !Lns_IsNil( _funcIndex ) {
                        funcIndex := _funcIndex.(LnsInt)
                        body.Insert(WasmBin_call(_env, self.rt.FP.FnOsExit(_env), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, funcIndex, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_refNull(_env, self.rt.FP.Get_tClo(_env)), WasmBin_call(_env, self.rt.FP.FnArgList(_env), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})))))))))
                    }
                }
            }
        }
    }
    var mainFunc *WasmBin_Func
    mainFunc = self.mod.FP.GetFunc(_env, self.mainIndex)
    mainFunc.FP.SetBody(_env, body)
    self.mod.FP.ExportFunc(_env, "main", self.mainIndex)
    self.rt.FP.FinalizeLazyLua(_env)
    self.stream.Write(_env, self.mod.FP.Encode(_env))
}
