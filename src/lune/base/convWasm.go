// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_convWasm bool
var convWasm__mod__ string
// decl enum -- Kind 
type convWasm_Kind = LnsInt
const convWasm_Kind__Bool = 2
const convWasm_Kind__Int = 0
const convWasm_Kind__Nil = 4
const convWasm_Kind__Real = 1
const convWasm_Kind__Str = 3
const convWasm_Kind__Void = 5
var convWasm_KindList_ = NewLnsList( []LnsAny {
  convWasm_Kind__Int,
  convWasm_Kind__Real,
  convWasm_Kind__Bool,
  convWasm_Kind__Str,
  convWasm_Kind__Nil,
  convWasm_Kind__Void,
})
func convWasm_Kind_get__allList_2_(_env *LnsEnv) *LnsList{
    return convWasm_KindList_
}
var convWasm_KindMap_ = map[LnsInt]string {
  convWasm_Kind__Bool: "Kind.Bool",
  convWasm_Kind__Int: "Kind.Int",
  convWasm_Kind__Nil: "Kind.Nil",
  convWasm_Kind__Real: "Kind.Real",
  convWasm_Kind__Str: "Kind.Str",
  convWasm_Kind__Void: "Kind.Void",
}
func convWasm_Kind__from_1_(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := convWasm_KindMap_[arg1]; ok { return arg1 }
    return nil
}

func convWasm_Kind_getTxt(arg1 LnsInt) string {
    return convWasm_KindMap_[arg1];
}
var convWasm_opI32Eqz LnsInt
var convWasm_opI32Eq LnsInt
var convWasm_opI32Ne LnsInt
var convWasm_opI32LtS LnsInt
var convWasm_opI32GtS LnsInt
var convWasm_opI32LeS LnsInt
var convWasm_opI32GeS LnsInt
var convWasm_opI32Add LnsInt
var convWasm_opI32Sub LnsInt
var convWasm_opI32And LnsInt
var convWasm_opI32Or LnsInt
var convWasm_opI64Eqz LnsInt
var convWasm_opI64Eq LnsInt
var convWasm_opI64Ne LnsInt
var convWasm_opI64LtS LnsInt
var convWasm_opI64GtS LnsInt
var convWasm_opI64LeS LnsInt
var convWasm_opI64GeS LnsInt
var convWasm_opI64Add LnsInt
var convWasm_opI64Sub LnsInt
var convWasm_opI64Mul LnsInt
var convWasm_opI64DivS LnsInt
var convWasm_opI64RemS LnsInt
var convWasm_opI64And LnsInt
var convWasm_opI64Or LnsInt
var convWasm_opI64Xor LnsInt
var convWasm_opI64Shl LnsInt
var convWasm_opI64ShrU LnsInt
var convWasm_opF64Eq LnsInt
var convWasm_opF64Ne LnsInt
var convWasm_opF64Lt LnsInt
var convWasm_opF64Gt LnsInt
var convWasm_opF64Le LnsInt
var convWasm_opF64Ge LnsInt
var convWasm_opF64Neg LnsInt
var convWasm_opF64Floor LnsInt
var convWasm_opF64Add LnsInt
var convWasm_opF64Sub LnsInt
var convWasm_opF64Mul LnsInt
var convWasm_opF64Div LnsInt
var convWasm_opI64ExtendI32U LnsInt
var convWasm_opF64ConvertI64S LnsInt
var convWasm_opI64TruncF64S LnsInt
var convWasm_opRefNull LnsInt
var convWasm_strCmpOpMap *LnsMap2_[string,LnsInt]
var convWasm_i64CmpOpMap *LnsMap2_[string,LnsInt]
var convWasm_f64CmpOpMap *LnsMap2_[string,LnsInt]
var convWasm_gcArrayNewData LnsInt
var convWasm_gcArrayNewDefault LnsInt
var convWasm_gcArrayGetU LnsInt
var convWasm_gcArrayLen LnsInt
var convWasm_gcArrayCopy LnsInt
// for 303
func convWasm_convExp1_377(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 320
func convWasm_convExp1_609(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 335
func convWasm_convExp1_825(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 347
func convWasm_convExp1_930(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 348
func convWasm_convExp1_952(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 370
func convWasm_convExp1_1226(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 391
func convWasm_convExp1_1495(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 417
func convWasm_convExp1_1866(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 424
func convWasm_convExp1_1939(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 429
func convWasm_convExp1_2002(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 492
func convWasm_convExp2_379(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 1158
func convWasm_convExp0_1303(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// for 1207
func convWasm_convExp0_1564(arg1 []LnsAny) (LnsInt, *WasmBin_Func) {
    return Lns_getFromMulti( arg1, 0 ).(LnsInt), Lns_getFromMulti( arg1, 1 ).(*WasmBin_Func)
}
// 134: decl @lune.@base.@convWasm.plain
func convWasm_plain_3_(_env *LnsEnv, code LnsInt,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, code), args)
}



// 1226: decl @lune.@base.@convWasm.createFilter
func ConvWasm_createFilter(_env *LnsEnv, streamName string,stream Lns_oStream,ast *AstInfo_ASTInfo) *Nodes_Filter {
    return &NewconvWasm_convFilter(_env, streamName, stream, ast).Nodes_Filter
}

// 222: decl @lune.@base.@convWasm.convFilter.error
func (self *convWasm_convFilter) error(_env *LnsEnv, node *Nodes_Node,mess string) {
    Util_err(_env, _env.GetVM().String_format("%s:%d:%d: wasm: %s", Lns_2DDD(self.streamName, node.FP.Get_pos(_env).LineNo, node.FP.Get_pos(_env).Column, mess)))
}
// 227: decl @lune.@base.@convWasm.convFilter.notSupport
func (self *convWasm_convFilter) notSupport(_env *LnsEnv, node *Nodes_Node) {
    self.FP.error(_env, node, _env.GetVM().String_format("not support -- %s", Lns_2DDD(Nodes_getNodeKindName(_env, node.FP.Get_kind(_env)))))
}
// 231: decl @lune.@base.@convWasm.convFilter.getKind
func (self *convWasm_convFilter) getKind(_env *LnsEnv, node *Nodes_Node,typeInfo *Ast_TypeInfo) LnsInt {
    if typeInfo.FP.Equals(_env, self.processInfo, Ast_builtinTypeInt, nil, nil){
        return convWasm_Kind__Int
    }
    if typeInfo.FP.Equals(_env, self.processInfo, Ast_builtinTypeReal, nil, nil){
        return convWasm_Kind__Real
    }
    if typeInfo.FP.Equals(_env, self.processInfo, Ast_builtinTypeBool, nil, nil){
        return convWasm_Kind__Bool
    }
    if typeInfo.FP.Equals(_env, self.processInfo, Ast_builtinTypeString, nil, nil){
        return convWasm_Kind__Str
    }
    if typeInfo.FP.Equals(_env, self.processInfo, Ast_builtinTypeNil, nil, nil){
        return convWasm_Kind__Nil
    }
    self.FP.error(_env, node, _env.GetVM().String_format("not support type -- %s", Lns_2DDD(typeInfo.FP.GetTxt(_env, nil, nil, nil))))
// insert a dummy
    return 0
}
// 250: decl @lune.@base.@convWasm.convFilter.getValType
func (self *convWasm_convFilter) getValType(_env *LnsEnv, node *Nodes_Node,kind LnsInt) *WasmBin_ValType {
    if _switch0 := kind; _switch0 == convWasm_Kind__Int {
        return WasmBin_typeI64
    } else if _switch0 == convWasm_Kind__Real {
        return WasmBin_typeF64
    } else if _switch0 == convWasm_Kind__Bool {
        return WasmBin_typeI32
    } else if _switch0 == convWasm_Kind__Str {
        return self.strType
    }
    self.FP.error(_env, node, _env.GetVM().String_format("not support type -- %s", Lns_2DDD(convWasm_Kind_getTxt( kind))))
// insert a dummy
    return nil
}
// 269: decl @lune.@base.@convWasm.convFilter.getDefaultVal
func (self *convWasm_convFilter) getDefaultVal(_env *LnsEnv, kind LnsInt) *WasmBin_Ins {
    if _switch0 := kind; _switch0 == convWasm_Kind__Real {
        return WasmBin_f64Const(_env, 0.0)
    } else if _switch0 == convWasm_Kind__Bool {
        return WasmBin_i32Const(_env, 0)
    } else if _switch0 == convWasm_Kind__Str {
        return WasmBin_op(_env, WasmBin_byte(_env, convWasm_opRefNull) + WasmBin_sleb(_env, self.strTypeIdx), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
    }
    return WasmBin_i64Const(_env, 0)
}
// 286: decl @lune.@base.@convWasm.convFilter.getCurFunc
func (self *convWasm_convFilter) getCurFunc(_env *LnsEnv, node *Nodes_Node) *WasmBin_Func {
    var _func *WasmBin_Func
    
    {
        __func := self.curFunc
        if __func == nil{
            self.FP.error(_env, node, "no function")
        } else {
            _func = __func.(*WasmBin_Func)
        }
    }
    return _func
}
// 294: decl @lune.@base.@convWasm.convFilter.setupRuntime
func (self *convWasm_convFilter) setupRuntime(_env *LnsEnv) {
    var i64 *WasmBin_ValType
    i64 = WasmBin_typeI64
    var f64 *WasmBin_ValType
    f64 = WasmBin_typeF64
    var i32 *WasmBin_ValType
    i32 = WasmBin_typeI32
    var strType *WasmBin_ValType
    strType = self.strType
    var strIdx string
    strIdx = WasmBin_uleb(_env, self.strTypeIdx)
    {
        var index LnsInt
        var _func *WasmBin_Func
        index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i64, i64)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i64)))
        var q LnsInt
        q = _func.FP.AddLocal(_env, i64)
        _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localSet(_env, q, convWasm_plain_3_(_env, convWasm_opI64DivS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), WasmBin_localGet(_env, 1))))), WasmBin_ifElse(_env, convWasm_plain_3_(_env, convWasm_opI32And, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI64Ne, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI64RemS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), WasmBin_localGet(_env, 1)))), WasmBin_i64Const(_env, 0)))), convWasm_plain_3_(_env, convWasm_opI32Ne, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), WasmBin_i64Const(_env, 0)))), convWasm_plain_3_(_env, convWasm_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 1), WasmBin_i64Const(_env, 0)))))))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localSet(_env, q, convWasm_plain_3_(_env, convWasm_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, q), WasmBin_i64Const(_env, 1))))))), nil, nil), WasmBin_localGet(_env, q))))
        self.fnDivI64 = index
    }
    {
        var index LnsInt
        var _func *WasmBin_Func
        index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i64, i64)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i64)))
        var r LnsInt
        r = _func.FP.AddLocal(_env, i64)
        _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localSet(_env, r, convWasm_plain_3_(_env, convWasm_opI64RemS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), WasmBin_localGet(_env, 1))))), WasmBin_ifElse(_env, convWasm_plain_3_(_env, convWasm_opI32And, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI64Ne, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, r), WasmBin_i64Const(_env, 0)))), convWasm_plain_3_(_env, convWasm_opI32Ne, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, r), WasmBin_i64Const(_env, 0)))), convWasm_plain_3_(_env, convWasm_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 1), WasmBin_i64Const(_env, 0)))))))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localSet(_env, r, convWasm_plain_3_(_env, convWasm_opI64Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, r), WasmBin_localGet(_env, 1))))))), nil, nil), WasmBin_localGet(_env, r))))
        self.fnModI64 = index
    }
    {
        var index LnsInt
        var _func *WasmBin_Func
        index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](f64, f64)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](f64)))
        _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opF64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), convWasm_plain_3_(_env, convWasm_opF64Mul, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opF64Floor, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opF64Div, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), WasmBin_localGet(_env, 1))))))), WasmBin_localGet(_env, 1))))))))))
        self.fnModF64 = index
    }
    {
        var indexL LnsInt
        var funcL *WasmBin_Func
        indexL,funcL = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i64, i64)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i64)))
        var indexR LnsInt
        var funcR *WasmBin_Func
        indexR,funcR = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i64, i64)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i64)))
        var convWasm_makeBody func(_env *LnsEnv, code LnsInt,revCode LnsInt) *LnsList2_[*WasmBin_Ins]
        convWasm_makeBody = func(_env *LnsEnv, code LnsInt,revCode LnsInt) *LnsList2_[*WasmBin_Ins] {
            return NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, convWasm_plain_3_(_env, convWasm_opI32Or, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI64LeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 1), WasmBin_i64Const(_env, -64)))), convWasm_plain_3_(_env, convWasm_opI64GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 1), WasmBin_i64Const(_env, 64))))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i64Const(_env, 0)))))), nil, nil), WasmBin_ifElse(_env, convWasm_plain_3_(_env, convWasm_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 1), WasmBin_i64Const(_env, 0)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, revCode, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), convWasm_plain_3_(_env, convWasm_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i64Const(_env, 0), WasmBin_localGet(_env, 1)))))))))))), nil, nil), convWasm_plain_3_(_env, code, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), WasmBin_localGet(_env, 1))))))
        }
        funcL.FP.SetBody(_env, convWasm_makeBody(_env, convWasm_opI64Shl, convWasm_opI64ShrU))
        funcR.FP.SetBody(_env, convWasm_makeBody(_env, convWasm_opI64ShrU, convWasm_opI64Shl))
        self.fnShiftL = indexL
        self.fnShiftR = indexR
    }
    {
        var index LnsInt
        var _func *WasmBin_Func
        index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](strType, strType)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](strType)))
        var lenA LnsInt
        lenA = _func.FP.AddLocal(_env, i32)
        var lenB LnsInt
        lenB = _func.FP.AddLocal(_env, i32)
        var work LnsInt
        work = _func.FP.AddLocal(_env, strType)
        _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localSet(_env, lenA, WasmBin_gcOp(_env, convWasm_gcArrayLen, "", NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0))))), WasmBin_localSet(_env, lenB, WasmBin_gcOp(_env, convWasm_gcArrayLen, "", NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 1))))), WasmBin_localSet(_env, work, WasmBin_gcOp(_env, convWasm_gcArrayNewDefault, strIdx, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, lenA), WasmBin_localGet(_env, lenB)))))))), WasmBin_gcOp(_env, convWasm_gcArrayCopy, strIdx + strIdx, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, work), WasmBin_i32Const(_env, 0), WasmBin_localGet(_env, 0), WasmBin_i32Const(_env, 0), WasmBin_localGet(_env, lenA)))), WasmBin_gcOp(_env, convWasm_gcArrayCopy, strIdx + strIdx, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, work), WasmBin_localGet(_env, lenA), WasmBin_localGet(_env, 1), WasmBin_i32Const(_env, 0), WasmBin_localGet(_env, lenB)))), WasmBin_localGet(_env, work))))
        self.fnStrConcat = index
    }
    {
        var index LnsInt
        var _func *WasmBin_Func
        index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](strType, strType)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i32)))
        var lenA LnsInt
        lenA = _func.FP.AddLocal(_env, i32)
        var lenB LnsInt
        lenB = _func.FP.AddLocal(_env, i32)
        var pos LnsInt
        pos = _func.FP.AddLocal(_env, i32)
        var diff LnsInt
        diff = _func.FP.AddLocal(_env, i32)
        _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localSet(_env, lenA, WasmBin_gcOp(_env, convWasm_gcArrayLen, "", NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0))))), WasmBin_localSet(_env, lenB, WasmBin_gcOp(_env, convWasm_gcArrayLen, "", NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 1))))), WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_brIf(_env, 1, convWasm_plain_3_(_env, convWasm_opI32GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, pos), WasmBin_localGet(_env, lenA))))), WasmBin_brIf(_env, 1, convWasm_plain_3_(_env, convWasm_opI32GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, pos), WasmBin_localGet(_env, lenB))))), WasmBin_localSet(_env, diff, convWasm_plain_3_(_env, convWasm_opI32Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_gcOp(_env, convWasm_gcArrayGetU, strIdx, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), WasmBin_localGet(_env, pos)))), WasmBin_gcOp(_env, convWasm_gcArrayGetU, strIdx, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 1), WasmBin_localGet(_env, pos)))))))), WasmBin_ifElse(_env, WasmBin_localGet(_env, diff), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ret(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, diff)))))), nil, nil), WasmBin_localSet(_env, pos, convWasm_plain_3_(_env, convWasm_opI32Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, pos), WasmBin_i32Const(_env, 1))))), WasmBin_br(_env, 0))), nil))), nil), convWasm_plain_3_(_env, convWasm_opI32Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, lenA), WasmBin_localGet(_env, lenB)))))))
        self.fnStrCmp = index
    }
    {
        var index LnsInt
        var _func *WasmBin_Func
        index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](strType, strType)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i32)))
        _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_call(_env, self.fnStrCmp, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), WasmBin_localGet(_env, 1))))))))))
        self.fnStrEq = index
    }
    {
        var index LnsInt
        var _func *WasmBin_Func
        index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](strType)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i32)))
        _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_gcOp(_env, convWasm_gcArrayLen, "", NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0)))))))
        self.mod.FP.ExportFunc(_env, "lns_str_len", index)
    }
    {
        var index LnsInt
        var _func *WasmBin_Func
        index,_func = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](strType, i32)), NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](i32)))
        _func.FP.SetBody(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_gcOp(_env, convWasm_gcArrayGetU, strIdx, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, 0), WasmBin_localGet(_env, 1)))))))
        self.mod.FP.ExportFunc(_env, "lns_str_get", index)
    }
}
// 436: decl @lune.@base.@convWasm.convFilter.decodeString
func (self *convWasm_convFilter) decodeString(_env *LnsEnv, node *Nodes_Node,txt string) string {
    var quote string
    quote = _env.GetVM().String_sub(txt,1, 1)
    if quote == "`"{
        return _env.GetVM().String_sub(txt,4, len(txt) - 3)
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
                        _code := convWasm_convExp2_379(Lns_2DDD(_env.GetVM().String_byte(body,pos, nil)))
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
// 514: decl @lune.@base.@convWasm.convFilter.getExpKind
func (self *convWasm_convFilter) getExpKind(_env *LnsEnv, node *Nodes_Node) LnsInt {
    {
        _castNode := Nodes_ExpCastNodeDownCastF(node.FP)
        if !Lns_IsNil( _castNode ) {
            castNode := _castNode.(*Nodes_ExpCastNode)
            return self.FP.getKind(_env, node, castNode.FP.Get_castType(_env))
        }
    }
    if node.FP.Get_kind(_env) == Nodes_NodeKind_get_ExpCall(_env){
        var callNode *Nodes_ExpCallNode
        
        {
            _callNode := Nodes_ExpCallNodeDownCastF(node.FP)
            if _callNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                callNode = _callNode.(*Nodes_ExpCallNode)
            }
        }
        var funcType *Ast_TypeInfo
        funcType = callNode.FP.Get_func(_env).FP.Get_expType(_env)
        var retList *LnsList2_[*Ast_TypeInfo]
        retList = funcType.FP.Get_retTypeInfoList(_env)
        if retList.Len() == 0{
            return convWasm_Kind__Void
        }
        if retList.Len() > 1{
            self.FP.error(_env, node, "not support multiple return values")
        }
        return self.FP.getKind(_env, node, retList.GetAt(1))
    }
    return self.FP.getKind(_env, node, node.FP.Get_expType(_env))
}
// 537: decl @lune.@base.@convWasm.convFilter.toReal
func (self *convWasm_convFilter) toReal(_env *LnsEnv, node *Nodes_Node) *WasmBin_Ins {
    var ins *WasmBin_Ins
    ins = self.FP.lowerExp(_env, node)
    if self.FP.getExpKind(_env, node) == convWasm_Kind__Int{
        return convWasm_plain_3_(_env, convWasm_opF64ConvertI64S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](ins)))
    }
    return ins
}
// 546: decl @lune.@base.@convWasm.convFilter.lowerOp2
func (self *convWasm_convFilter) lowerOp2(_env *LnsEnv, node *Nodes_ExpOp2Node) *WasmBin_Ins {
    var opTxt string
    opTxt = node.FP.Get_op(_env).Txt
    var kind1 LnsInt
    kind1 = self.FP.getExpKind(_env, node.FP.Get_exp1(_env))
    var kind2 LnsInt
    kind2 = self.FP.getExpKind(_env, node.FP.Get_exp2(_env))
    if _switch0 := opTxt; _switch0 == "and" || _switch0 == "or" {
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( kind1 != convWasm_Kind__Bool) ||
            _env.SetStackVal( kind2 != convWasm_Kind__Bool) ).(bool){
            self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support non bool %s", Lns_2DDD(opTxt)))
        }
        var exp1 *WasmBin_Ins
        exp1 = self.FP.lowerExp(_env, node.FP.Get_exp1(_env))
        self.depth = self.depth + 1
        var exp2 *WasmBin_Ins
        exp2 = self.FP.lowerExp(_env, node.FP.Get_exp2(_env))
        self.depth = self.depth - 1
        if opTxt == "and"{
            return WasmBin_ifElse(_env, exp1, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp2)), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0))), WasmBin_typeI32)
        }
        return WasmBin_ifElse(_env, exp1, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 1))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp2)), WasmBin_typeI32)
    } else if _switch0 == ".." {
        return WasmBin_call(_env, self.fnStrConcat, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, node.FP.Get_exp1(_env)), self.FP.lowerExp(_env, node.FP.Get_exp2(_env)))))
    }
    if _switch2 := opTxt; _switch2 == "==" || _switch2 == "~=" || _switch2 == "<" || _switch2 == "<=" || _switch2 == ">" || _switch2 == ">=" {
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( kind1 == convWasm_Kind__Str) &&
            _env.SetStackVal( kind2 == convWasm_Kind__Str) ).(bool)){
            var cmp *WasmBin_Ins
            cmp = WasmBin_call(_env, self.fnStrCmp, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, node.FP.Get_exp1(_env)), self.FP.lowerExp(_env, node.FP.Get_exp2(_env)))))
            if opTxt == "=="{
                return convWasm_plain_3_(_env, convWasm_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cmp)))
            }
            var code LnsInt
            
            {
                _code := convWasm_strCmpOpMap.Get(opTxt)
                if _code == nil{
                    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("illegal op -- %s", Lns_2DDD(opTxt)))
                } else {
                    code = _code.(LnsInt)
                }
            }
            return convWasm_plain_3_(_env, code, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cmp, WasmBin_i32Const(_env, 0))))
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( kind1 == convWasm_Kind__Bool) &&
            _env.SetStackVal( kind2 == convWasm_Kind__Bool) ).(bool)){
            var code LnsInt
            if _switch1 := opTxt; _switch1 == "==" {
                code = convWasm_opI32Eq
            } else if _switch1 == "~=" {
                code = convWasm_opI32Ne
            } else {
                self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support bool %s", Lns_2DDD(opTxt)))
            }
            return convWasm_plain_3_(_env, code, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, node.FP.Get_exp1(_env)), self.FP.lowerExp(_env, node.FP.Get_exp2(_env)))))
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( kind1 == convWasm_Kind__Int) &&
            _env.SetStackVal( kind2 == convWasm_Kind__Int) ).(bool)){
            var code LnsInt
            
            {
                _code := convWasm_i64CmpOpMap.Get(opTxt)
                if _code == nil{
                    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("illegal op -- %s", Lns_2DDD(opTxt)))
                } else {
                    code = _code.(LnsInt)
                }
            }
            return convWasm_plain_3_(_env, code, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, node.FP.Get_exp1(_env)), self.FP.lowerExp(_env, node.FP.Get_exp2(_env)))))
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( kind1 == convWasm_Kind__Int) ||
                _env.SetStackVal( kind1 == convWasm_Kind__Real) ).(bool))) &&
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( kind2 == convWasm_Kind__Int) ||
                _env.SetStackVal( kind2 == convWasm_Kind__Real) ).(bool))) ).(bool)){
            var code LnsInt
            
            {
                _code := convWasm_f64CmpOpMap.Get(opTxt)
                if _code == nil{
                    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("illegal op -- %s", Lns_2DDD(opTxt)))
                } else {
                    code = _code.(LnsInt)
                }
            }
            return convWasm_plain_3_(_env, code, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.toReal(_env, node.FP.Get_exp1(_env)), self.FP.toReal(_env, node.FP.Get_exp2(_env)))))
        }
        self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support compare -- %s %s %s", Lns_2DDD(convWasm_Kind_getTxt( kind1), opTxt, convWasm_Kind_getTxt( kind2))))
    }
    var resultKind LnsInt
    resultKind = self.FP.getKind(_env, &node.Nodes_Node, node.FP.Get_expType(_env))
    if resultKind == convWasm_Kind__Int{
        var exp1 *WasmBin_Ins
        exp1 = self.FP.lowerExp(_env, node.FP.Get_exp1(_env))
        var exp2 *WasmBin_Ins
        exp2 = self.FP.lowerExp(_env, node.FP.Get_exp2(_env))
        if _switch3 := opTxt; _switch3 == "+" {
            return convWasm_plain_3_(_env, convWasm_opI64Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch3 == "-" {
            return convWasm_plain_3_(_env, convWasm_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch3 == "*" {
            return convWasm_plain_3_(_env, convWasm_opI64Mul, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch3 == "/" {
            return WasmBin_call(_env, self.fnDivI64, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch3 == "%" {
            return WasmBin_call(_env, self.fnModI64, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch3 == "&" {
            return convWasm_plain_3_(_env, convWasm_opI64And, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch3 == "|" {
            return convWasm_plain_3_(_env, convWasm_opI64Or, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch3 == "~" {
            return convWasm_plain_3_(_env, convWasm_opI64Xor, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch3 == "|<<" {
            return WasmBin_call(_env, self.fnShiftL, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch3 == "|>>" {
            return WasmBin_call(_env, self.fnShiftR, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        }
    }
    if resultKind == convWasm_Kind__Real{
        var exp1 *WasmBin_Ins
        exp1 = self.FP.toReal(_env, node.FP.Get_exp1(_env))
        var exp2 *WasmBin_Ins
        exp2 = self.FP.toReal(_env, node.FP.Get_exp2(_env))
        if _switch4 := opTxt; _switch4 == "+" {
            return convWasm_plain_3_(_env, convWasm_opF64Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch4 == "-" {
            return convWasm_plain_3_(_env, convWasm_opF64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch4 == "*" {
            return convWasm_plain_3_(_env, convWasm_opF64Mul, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch4 == "/" {
            return convWasm_plain_3_(_env, convWasm_opF64Div, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        } else if _switch4 == "%" {
            return WasmBin_call(_env, self.fnModF64, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp1, exp2)))
        }
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support op -- %s (%s)", Lns_2DDD(opTxt, convWasm_Kind_getTxt( resultKind))))
// insert a dummy
    return nil
}
// 647: decl @lune.@base.@convWasm.convFilter.lowerOp1
func (self *convWasm_convFilter) lowerOp1(_env *LnsEnv, node *Nodes_ExpOp1Node) *WasmBin_Ins {
    var opTxt string
    opTxt = node.FP.Get_op(_env).Txt
    var kind LnsInt
    kind = self.FP.getExpKind(_env, node.FP.Get_exp(_env))
    var exp *WasmBin_Ins
    exp = self.FP.lowerExp(_env, node.FP.Get_exp(_env))
    if _switch0 := opTxt; _switch0 == "-" {
        if kind == convWasm_Kind__Int{
            return convWasm_plain_3_(_env, convWasm_opI64Sub, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i64Const(_env, 0), exp)))
        }
        if kind == convWasm_Kind__Real{
            return convWasm_plain_3_(_env, convWasm_opF64Neg, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp)))
        }
    } else if _switch0 == "not" {
        if kind == convWasm_Kind__Bool{
            return convWasm_plain_3_(_env, convWasm_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp)))
        }
    } else if _switch0 == "#" {
        if kind == convWasm_Kind__Str{
            return convWasm_plain_3_(_env, convWasm_opI64ExtendI32U, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_gcOp(_env, convWasm_gcArrayLen, "", NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp))))))
        }
    } else if _switch0 == "~" {
        if kind == convWasm_Kind__Int{
            return convWasm_plain_3_(_env, convWasm_opI64Xor, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp, WasmBin_i64Const(_env, -1))))
        }
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support op -- %s (%s)", Lns_2DDD(opTxt, convWasm_Kind_getTxt( kind))))
// insert a dummy
    return nil
}
// 680: decl @lune.@base.@convWasm.convFilter.lowerArgs
func (self *convWasm_convFilter) lowerArgs(_env *LnsEnv, node *Nodes_ExpCallNode) *LnsList2_[*WasmBin_Ins] {
    var list *LnsList2_[*WasmBin_Ins]
    list = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    {
        _argList := node.FP.Get_argList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            if Lns_isCondTrue( argList.FP.Get_mRetExp(_env)){
                self.FP.error(_env, &node.Nodes_Node, "not support multiple return values in args")
            }
            for _, _arg := range( argList.FP.Get_expList(_env).Items ) {
                arg := _arg
                list.Insert(self.FP.lowerExp(_env, arg))
            }
        }
    }
    return list
}
// 698: decl @lune.@base.@convWasm.convFilter.getDDDArgs
func (self *convWasm_convFilter) getDDDArgs(_env *LnsEnv, node *Nodes_ExpCallNode) *LnsList2_[*Nodes_Node] {
    var list *LnsList2_[*Nodes_Node]
    list = NewLnsList2_[*Nodes_Node]([]*Nodes_Node{})
    {
        _argList := node.FP.Get_argList(_env)
        if !Lns_IsNil( _argList ) {
            argList := _argList.(*Nodes_ExpListNode)
            if Lns_isCondTrue( argList.FP.Get_mRetExp(_env)){
                self.FP.error(_env, &node.Nodes_Node, "not support multiple return values in args")
            }
            for _, _arg := range( argList.FP.Get_expList(_env).Items ) {
                arg := _arg
                {
                    _toDDD := Nodes_ExpToDDDNodeDownCastF(arg.FP)
                    if !Lns_IsNil( _toDDD ) {
                        toDDD := _toDDD.(*Nodes_ExpToDDDNode)
                        for _, _exp := range( toDDD.FP.Get_expList(_env).FP.Get_expList(_env).Items ) {
                            exp := _exp
                            list.Insert(exp)
                        }
                    } else {
                        list.Insert(arg)
                    }
                }
            }
        }
    }
    return list
}
// 718: decl @lune.@base.@convWasm.convFilter.lowerPrint
func (self *convWasm_convFilter) lowerPrint(_env *LnsEnv, node *Nodes_ExpCallNode) *WasmBin_Ins {
    var list *LnsList2_[*WasmBin_Ins]
    list = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    {
        for _index, _arg := range( self.FP.getDDDArgs(_env, node).Items ) {
            index := _index + 1
            arg := _arg
            if index > 1{
                list.Insert(WasmBin_call(_env, self.fnPrintSep, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})))
            }
            var kind LnsInt
            kind = self.FP.getExpKind(_env, arg)
            if _switch0 := kind; _switch0 == convWasm_Kind__Int {
                list.Insert(WasmBin_call(_env, self.fnPrintI64, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, arg)))))
            } else if _switch0 == convWasm_Kind__Real {
                list.Insert(WasmBin_call(_env, self.fnPrintF64, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, arg)))))
            } else if _switch0 == convWasm_Kind__Bool {
                list.Insert(WasmBin_call(_env, self.fnPrintBool, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, arg)))))
            } else if _switch0 == convWasm_Kind__Str {
                list.Insert(WasmBin_call(_env, self.fnPrintStr, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, arg)))))
            } else if _switch0 == convWasm_Kind__Nil {
                list.Insert(WasmBin_call(_env, self.fnPrintNil, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})))
            } else if _switch0 == convWasm_Kind__Void {
                self.FP.error(_env, arg, "void value")
            }
        }
    }
    list.Insert(WasmBin_call(_env, self.fnPrintEnd, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})))
    return WasmBin_seq(_env, list)
}
// 752: decl @lune.@base.@convWasm.convFilter.lowerCall
func (self *convWasm_convFilter) lowerCall(_env *LnsEnv, node *Nodes_ExpCallNode) *WasmBin_Ins {
    if node.FP.Get_func(_env).FP.Get_expType(_env).FP.Equals(_env, self.processInfo, self.builtinFuncs.Lns_print, nil, nil){
        return self.FP.lowerPrint(_env, node)
    }
    var refNode *Nodes_ExpRefNode
    
    {
        _refNode := Nodes_ExpRefNodeDownCastF(node.FP.Get_func(_env).FP)
        if _refNode == nil{
            self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support func -- %s", Lns_2DDD(Nodes_getNodeKindName(_env, node.FP.Get_func(_env).FP.Get_kind(_env)))))
        } else {
            refNode = _refNode.(*Nodes_ExpRefNode)
        }
    }
    var funcIndex LnsInt
    
    {
        _funcIndex := self.sym2func.Get(refNode.FP.Get_symbolInfo(_env).FP.GetOrg(_env))
        if _funcIndex == nil{
            self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support func -- %s", Lns_2DDD(refNode.FP.Get_symbolInfo(_env).FP.Get_name(_env))))
        } else {
            funcIndex = _funcIndex.(LnsInt)
        }
    }
    return WasmBin_call(_env, funcIndex, self.FP.lowerArgs(_env, node))
}
// 766: decl @lune.@base.@convWasm.convFilter.lowerRef
func (self *convWasm_convFilter) lowerRef(_env *LnsEnv, node *Nodes_ExpRefNode) *WasmBin_Ins {
    var symbol *Ast_SymbolInfo
    symbol = node.FP.Get_symbolInfo(_env).FP.GetOrg(_env)
    {
        _index := self.sym2local.Get(symbol)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return WasmBin_localGet(_env, index)
        }
    }
    {
        _index := self.sym2global.Get(symbol)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return WasmBin_globalGet(_env, index)
        }
    }
    self.FP.error(_env, &node.Nodes_Node, _env.GetVM().String_format("not support symbol -- %s", Lns_2DDD(symbol.FP.Get_name(_env))))
// insert a dummy
    return nil
}
// 778: decl @lune.@base.@convWasm.convFilter.lowerSetSym
func (self *convWasm_convFilter) lowerSetSym(_env *LnsEnv, node *Nodes_Node,symbol *Ast_SymbolInfo,val *WasmBin_Ins) *WasmBin_Ins {
    var orgSym *Ast_SymbolInfo
    orgSym = symbol.FP.GetOrg(_env)
    {
        _index := self.sym2local.Get(orgSym)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return WasmBin_localSet(_env, index, val)
        }
    }
    {
        _index := self.sym2global.Get(orgSym)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return WasmBin_globalSet(_env, index, val)
        }
    }
    self.FP.error(_env, node, _env.GetVM().String_format("not support symbol -- %s", Lns_2DDD(symbol.FP.Get_name(_env))))
// insert a dummy
    return nil
}
// 790: decl @lune.@base.@convWasm.convFilter.lowerExp
func (self *convWasm_convFilter) lowerExp(_env *LnsEnv, node *Nodes_Node) *WasmBin_Ins {
    if _switch0 := node.FP.Get_kind(_env); _switch0 == Nodes_NodeKind_get_LiteralInt(_env) {
        var literal *Nodes_LiteralIntNode
        
        {
            _literal := Nodes_LiteralIntNodeDownCastF(node.FP)
            if _literal == nil{
                self.FP.error(_env, node, "cast")
            } else {
                literal = _literal.(*Nodes_LiteralIntNode)
            }
        }
        return WasmBin_i64Const(_env, literal.FP.Get_num(_env))
    } else if _switch0 == Nodes_NodeKind_get_LiteralReal(_env) {
        var literal *Nodes_LiteralRealNode
        
        {
            _literal := Nodes_LiteralRealNodeDownCastF(node.FP)
            if _literal == nil{
                self.FP.error(_env, node, "cast")
            } else {
                literal = _literal.(*Nodes_LiteralRealNode)
            }
        }
        return WasmBin_f64Const(_env, literal.FP.Get_num(_env))
    } else if _switch0 == Nodes_NodeKind_get_LiteralBool(_env) {
        var literal *Nodes_LiteralBoolNode
        
        {
            _literal := Nodes_LiteralBoolNodeDownCastF(node.FP)
            if _literal == nil{
                self.FP.error(_env, node, "cast")
            } else {
                literal = _literal.(*Nodes_LiteralBoolNode)
            }
        }
        if literal.FP.Get_token(_env).Txt == "true"{
            return WasmBin_i32Const(_env, 1)
        }
        return WasmBin_i32Const(_env, 0)
    } else if _switch0 == Nodes_NodeKind_get_LiteralString(_env) {
        var literal *Nodes_LiteralStringNode
        
        {
            _literal := Nodes_LiteralStringNodeDownCastF(node.FP)
            if _literal == nil{
                self.FP.error(_env, node, "cast")
            } else {
                literal = _literal.(*Nodes_LiteralStringNode)
            }
        }
        if Lns_isCondTrue( literal.FP.Get_orgParam(_env)){
            self.FP.error(_env, node, "not support format string")
        }
        var data string
        data = self.FP.decodeString(_env, node, literal.FP.Get_token(_env).Txt)
        var dataIdx LnsInt
        dataIdx = self.mod.FP.AddData(_env, data)
        return WasmBin_gcOp(_env, convWasm_gcArrayNewData, WasmBin_uleb(_env, self.strTypeIdx) + WasmBin_uleb(_env, dataIdx), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_i32Const(_env, 0), WasmBin_i32Const(_env, len(data)))))
    } else if _switch0 == Nodes_NodeKind_get_ExpRef(_env) {
        var refNode *Nodes_ExpRefNode
        
        {
            _refNode := Nodes_ExpRefNodeDownCastF(node.FP)
            if _refNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                refNode = _refNode.(*Nodes_ExpRefNode)
            }
        }
        return self.FP.lowerRef(_env, refNode)
    } else if _switch0 == Nodes_NodeKind_get_ExpParen(_env) {
        var parenNode *Nodes_ExpParenNode
        
        {
            _parenNode := Nodes_ExpParenNodeDownCastF(node.FP)
            if _parenNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                parenNode = _parenNode.(*Nodes_ExpParenNode)
            }
        }
        return self.FP.lowerExp(_env, parenNode.FP.Get_exp(_env))
    } else if _switch0 == Nodes_NodeKind_get_ExpOp2(_env) {
        var op2Node *Nodes_ExpOp2Node
        
        {
            _op2Node := Nodes_ExpOp2NodeDownCastF(node.FP)
            if _op2Node == nil{
                self.FP.error(_env, node, "cast")
            } else {
                op2Node = _op2Node.(*Nodes_ExpOp2Node)
            }
        }
        return self.FP.lowerOp2(_env, op2Node)
    } else if _switch0 == Nodes_NodeKind_get_ExpOp1(_env) {
        var op1Node *Nodes_ExpOp1Node
        
        {
            _op1Node := Nodes_ExpOp1NodeDownCastF(node.FP)
            if _op1Node == nil{
                self.FP.error(_env, node, "cast")
            } else {
                op1Node = _op1Node.(*Nodes_ExpOp1Node)
            }
        }
        return self.FP.lowerOp1(_env, op1Node)
    } else if _switch0 == Nodes_NodeKind_get_ExpCall(_env) {
        var callNode *Nodes_ExpCallNode
        
        {
            _callNode := Nodes_ExpCallNodeDownCastF(node.FP)
            if _callNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                callNode = _callNode.(*Nodes_ExpCallNode)
            }
        }
        return self.FP.lowerCall(_env, callNode)
    } else if _switch0 == Nodes_NodeKind_get_ExpCast(_env) {
        var castNode *Nodes_ExpCastNode
        
        {
            _castNode := Nodes_ExpCastNodeDownCastF(node.FP)
            if _castNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                castNode = _castNode.(*Nodes_ExpCastNode)
            }
        }
        var srcKind LnsInt
        srcKind = self.FP.getExpKind(_env, castNode.FP.Get_exp(_env))
        var dstKind LnsInt
        dstKind = self.FP.getKind(_env, node, castNode.FP.Get_castType(_env))
        var exp *WasmBin_Ins
        exp = self.FP.lowerExp(_env, castNode.FP.Get_exp(_env))
        if srcKind == dstKind{
            return exp
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( srcKind == convWasm_Kind__Int) &&
            _env.SetStackVal( dstKind == convWasm_Kind__Real) ).(bool)){
            return convWasm_plain_3_(_env, convWasm_opF64ConvertI64S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp)))
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( srcKind == convWasm_Kind__Real) &&
            _env.SetStackVal( dstKind == convWasm_Kind__Int) ).(bool)){
            return convWasm_plain_3_(_env, convWasm_opI64TruncF64S, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](exp)))
        }
        self.FP.error(_env, node, _env.GetVM().String_format("not support cast -- %s -> %s", Lns_2DDD(convWasm_Kind_getTxt( srcKind), convWasm_Kind_getTxt( dstKind))))
    }
    self.FP.notSupport(_env, node)
// insert a dummy
    return nil
}
// 878: decl @lune.@base.@convWasm.convFilter.lowerBlock
func (self *convWasm_convFilter) lowerBlock(_env *LnsEnv, node *Nodes_BlockNode) *LnsList2_[*WasmBin_Ins] {
    var list *LnsList2_[*WasmBin_Ins]
    list = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _, _stmt := range( node.FP.Get_stmtList(_env).Items ) {
        stmt := _stmt
        self.FP.lowerStmt(_env, stmt, list)
    }
    return list
}
// 892: decl @lune.@base.@convWasm.convFilter.enterLoop
func (self *convWasm_convFilter) enterLoop(_env *LnsEnv) {
    self.depth = self.depth + 1
    self.breakDepthList.Insert(LnsInt(self.depth))
    self.depth = self.depth + 1
}
// 899: decl @lune.@base.@convWasm.convFilter.exitLoop
func (self *convWasm_convFilter) exitLoop(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    self.depth = self.depth - 2
    self.breakDepthList.Remove(nil)
    return WasmBin_block(_env, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_loop(_env, body, nil))), nil)
}
// 905: decl @lune.@base.@convWasm.convFilter.lowerDeclVar
func (self *convWasm_convFilter) lowerDeclVar(_env *LnsEnv, node *Nodes_DeclVarNode,out *LnsList2_[*WasmBin_Ins]) {
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( node.FP.Get_unwrapFlag(_env)) ||
        _env.SetStackVal( node.FP.Get_condRetInfo(_env)) )){
        self.FP.error(_env, &node.Nodes_Node, "not support let!")
    }
    var expList *LnsList2_[*Nodes_Node]
    expList = NewLnsList2_[*Nodes_Node]([]*Nodes_Node{})
    {
        _expListNode := node.FP.Get_expList(_env)
        if !Lns_IsNil( _expListNode ) {
            expListNode := _expListNode.(*Nodes_ExpListNode)
            if Lns_isCondTrue( expListNode.FP.Get_mRetExp(_env)){
                self.FP.error(_env, &node.Nodes_Node, "not support multiple return values")
            }
            for _, _exp := range( expListNode.FP.Get_expList(_env).Items ) {
                exp := _exp
                expList.Insert(exp)
            }
        }
    }
    for _index, _symbol := range( node.FP.Get_symbolInfoList(_env).Items ) {
        index := _index + 1
        symbol := _symbol
        var orgSym *Ast_SymbolInfo
        orgSym = symbol.FP.GetOrg(_env)
        var kind LnsInt
        kind = self.FP.getKind(_env, &node.Nodes_Node, symbol.FP.Get_typeInfo(_env))
        var valType *WasmBin_ValType
        valType = self.FP.getValType(_env, &node.Nodes_Node, kind)
        if WasmBin_Func2Stem(self.curFunc) == nil{
            self.FP.error(_env, &node.Nodes_Node, "no function")
        }
        if symbol.FP.Get_scope(_env) == self.moduleScope{
            self.sym2global.Set(orgSym,self.mod.FP.AddGlobal(_env, valType, self.FP.getDefaultVal(_env, kind)))
        } else { 
            var _func *WasmBin_Func
            _func = self.FP.getCurFunc(_env, &node.Nodes_Node)
            self.sym2local.Set(orgSym,_func.FP.AddLocal(_env, valType))
        }
        if index <= expList.Len(){
            out.Insert(self.FP.lowerSetSym(_env, &node.Nodes_Node, orgSym, self.FP.lowerExp(_env, expList.GetAt(index))))
        }
    }
}
// 938: decl @lune.@base.@convWasm.convFilter.lowerIf
func (self *convWasm_convFilter) lowerIf(_env *LnsEnv, node *Nodes_IfNode) *WasmBin_Ins {
    var convWasm_process func(_env *LnsEnv, index LnsInt) *LnsList2_[*WasmBin_Ins]
    convWasm_process = func(_env *LnsEnv, index LnsInt) *LnsList2_[*WasmBin_Ins] {
        var stmt *Nodes_IfStmtInfo
        stmt = node.FP.Get_stmtList(_env).GetAt(index)
        if stmt.FP.Get_kind(_env) == Nodes_IfKind__Else{
            return self.FP.lowerBlock(_env, stmt.FP.Get_block(_env))
        }
        var cond *WasmBin_Ins
        cond = self.FP.lowerExp(_env, stmt.FP.Get_exp(_env))
        self.depth = self.depth + 1
        var thenBody *LnsList2_[*WasmBin_Ins]
        thenBody = self.FP.lowerBlock(_env, stmt.FP.Get_block(_env))
        var elseBody LnsAny
        elseBody = nil
        if index < node.FP.Get_stmtList(_env).Len(){
            elseBody = convWasm_process(_env, index + 1)
        }
        self.depth = self.depth - 1
        return NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_ifElse(_env, cond, thenBody, elseBody, nil)))
    }
    return WasmBin_seq(_env, convWasm_process(_env, 1))
}
// 958: decl @lune.@base.@convWasm.convFilter.lowerFor
func (self *convWasm_convFilter) lowerFor(_env *LnsEnv, node *Nodes_ForNode) *WasmBin_Ins {
    if self.FP.getKind(_env, &node.Nodes_Node, node.FP.Get_val(_env).FP.Get_typeInfo(_env)) != convWasm_Kind__Int{
        self.FP.error(_env, &node.Nodes_Node, "not support non int for")
    }
    var _func *WasmBin_Func
    _func = self.FP.getCurFunc(_env, &node.Nodes_Node)
    var valIdx LnsInt
    valIdx = _func.FP.AddLocal(_env, WasmBin_typeI64)
    self.sym2local.Set(node.FP.Get_val(_env).FP.GetOrg(_env),valIdx)
    var toIdx LnsInt
    toIdx = _func.FP.AddLocal(_env, WasmBin_typeI64)
    var stepIdx LnsInt
    stepIdx = _func.FP.AddLocal(_env, WasmBin_typeI64)
    var list *LnsList2_[*WasmBin_Ins]
    list = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    list.Insert(WasmBin_localSet(_env, valIdx, self.FP.lowerExp(_env, node.FP.Get_init(_env))))
    list.Insert(WasmBin_localSet(_env, toIdx, self.FP.lowerExp(_env, node.FP.Get_to(_env))))
    {
        _delta := node.FP.Get_delta(_env)
        if !Lns_IsNil( _delta ) {
            delta := _delta.(*Nodes_Node)
            list.Insert(WasmBin_localSet(_env, stepIdx, self.FP.lowerExp(_env, delta)))
        } else {
            list.Insert(WasmBin_localSet(_env, stepIdx, WasmBin_i64Const(_env, 1)))
        }
    }
    self.FP.enterLoop(_env)
    {
        var body *LnsList2_[*WasmBin_Ins]
        body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        self.depth = self.depth + 1
        var exitCond *WasmBin_Ins
        exitCond = WasmBin_ifElse(_env, convWasm_plain_3_(_env, convWasm_opI64GeS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, stepIdx), WasmBin_i64Const(_env, 0)))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI64GtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, valIdx), WasmBin_localGet(_env, toIdx)))))), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](convWasm_plain_3_(_env, convWasm_opI64LtS, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, valIdx), WasmBin_localGet(_env, toIdx)))))), WasmBin_typeI32)
        self.depth = self.depth - 1
        body.Insert(WasmBin_brIf(_env, 1, exitCond))
        for _, _ins := range( self.FP.lowerBlock(_env, node.FP.Get_block(_env)).Items ) {
            ins := _ins
            body.Insert(ins)
        }
        body.Insert(WasmBin_localSet(_env, valIdx, convWasm_plain_3_(_env, convWasm_opI64Add, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](WasmBin_localGet(_env, valIdx), WasmBin_localGet(_env, stepIdx))))))
        body.Insert(WasmBin_br(_env, 0))
        list.Insert(self.FP.exitLoop(_env, body))
    }
    return WasmBin_seq(_env, list)
}
// 999: decl @lune.@base.@convWasm.convFilter.lowerStmt
func (self *convWasm_convFilter) lowerStmt(_env *LnsEnv, node *Nodes_Node,out *LnsList2_[*WasmBin_Ins]) {
    if _switch0 := node.FP.Get_kind(_env); _switch0 == Nodes_NodeKind_get_None(_env) || _switch0 == Nodes_NodeKind_get_BlankLine(_env) || _switch0 == Nodes_NodeKind_get_Shebang(_env) || _switch0 == Nodes_NodeKind_get_LuneControl(_env) {
        return 
    } else if _switch0 == Nodes_NodeKind_get_DeclFunc(_env) {
        if self.depth == 0{
            return 
        }
    } else if _switch0 == Nodes_NodeKind_get_StmtExp(_env) {
        var stmtNode *Nodes_StmtExpNode
        
        {
            _stmtNode := Nodes_StmtExpNodeDownCastF(node.FP)
            if _stmtNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                stmtNode = _stmtNode.(*Nodes_StmtExpNode)
            }
        }
        if stmtNode.FP.Get_exp(_env).FP.Get_kind(_env) == Nodes_NodeKind_get_ExpSetVal(_env){
            self.FP.lowerStmt(_env, stmtNode.FP.Get_exp(_env), out)
            return 
        }
        var ins *WasmBin_Ins
        ins = self.FP.lowerExp(_env, stmtNode.FP.Get_exp(_env))
        if self.FP.getExpKind(_env, stmtNode.FP.Get_exp(_env)) == convWasm_Kind__Void{
            out.Insert(ins)
        } else { 
            out.Insert(WasmBin_drop(_env, ins))
        }
        return 
    } else if _switch0 == Nodes_NodeKind_get_DeclVar(_env) {
        var declNode *Nodes_DeclVarNode
        
        {
            _declNode := Nodes_DeclVarNodeDownCastF(node.FP)
            if _declNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                declNode = _declNode.(*Nodes_DeclVarNode)
            }
        }
        self.FP.lowerDeclVar(_env, declNode, out)
        return 
    } else if _switch0 == Nodes_NodeKind_get_ExpSetVal(_env) {
        var setNode *Nodes_ExpSetValNode
        
        {
            _setNode := Nodes_ExpSetValNodeDownCastF(node.FP)
            if _setNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                setNode = _setNode.(*Nodes_ExpSetValNode)
            }
        }
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( setNode.FP.Get_LeftSymList(_env).Len() != 1) ||
            _env.SetStackVal( setNode.FP.Get_exp2(_env).FP.Get_expList(_env).Len() != 1) ).(bool){
            self.FP.error(_env, node, "not support multiple assign")
        }
        out.Insert(self.FP.lowerSetSym(_env, node, setNode.FP.Get_LeftSymList(_env).GetAt(1), self.FP.lowerExp(_env, setNode.FP.Get_exp2(_env).FP.Get_expList(_env).GetAt(1))))
        return 
    } else if _switch0 == Nodes_NodeKind_get_If(_env) {
        var ifNode *Nodes_IfNode
        
        {
            _ifNode := Nodes_IfNodeDownCastF(node.FP)
            if _ifNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                ifNode = _ifNode.(*Nodes_IfNode)
            }
        }
        out.Insert(self.FP.lowerIf(_env, ifNode))
        return 
    } else if _switch0 == Nodes_NodeKind_get_While(_env) {
        var whileNode *Nodes_WhileNode
        
        {
            _whileNode := Nodes_WhileNodeDownCastF(node.FP)
            if _whileNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                whileNode = _whileNode.(*Nodes_WhileNode)
            }
        }
        self.FP.enterLoop(_env)
        {
            var body *LnsList2_[*WasmBin_Ins]
            body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
            if Lns_op_not(whileNode.FP.Get_infinit(_env)){
                body.Insert(WasmBin_brIf(_env, 1, convWasm_plain_3_(_env, convWasm_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, whileNode.FP.Get_exp(_env)))))))
            }
            for _, _ins := range( self.FP.lowerBlock(_env, whileNode.FP.Get_block(_env)).Items ) {
                ins := _ins
                body.Insert(ins)
            }
            body.Insert(WasmBin_br(_env, 0))
            out.Insert(self.FP.exitLoop(_env, body))
        }
        return 
    } else if _switch0 == Nodes_NodeKind_get_Repeat(_env) {
        var repeatNode *Nodes_RepeatNode
        
        {
            _repeatNode := Nodes_RepeatNodeDownCastF(node.FP)
            if _repeatNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                repeatNode = _repeatNode.(*Nodes_RepeatNode)
            }
        }
        self.FP.enterLoop(_env)
        {
            var body *LnsList2_[*WasmBin_Ins]
            body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
            for _, _ins := range( self.FP.lowerBlock(_env, repeatNode.FP.Get_block(_env)).Items ) {
                ins := _ins
                body.Insert(ins)
            }
            body.Insert(WasmBin_brIf(_env, 0, convWasm_plain_3_(_env, convWasm_opI32Eqz, NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](self.FP.lowerExp(_env, repeatNode.FP.Get_exp(_env)))))))
            out.Insert(self.FP.exitLoop(_env, body))
        }
        return 
    } else if _switch0 == Nodes_NodeKind_get_For(_env) {
        var forNode *Nodes_ForNode
        
        {
            _forNode := Nodes_ForNodeDownCastF(node.FP)
            if _forNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                forNode = _forNode.(*Nodes_ForNode)
            }
        }
        out.Insert(self.FP.lowerFor(_env, forNode))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Break(_env) {
        if self.breakDepthList.Len() == 0{
            self.FP.error(_env, node, "break outside loop")
        }
        out.Insert(WasmBin_br(_env, self.depth - self.breakDepthList.GetAt(self.breakDepthList.Len())))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Return(_env) {
        var retNode *Nodes_ReturnNode
        
        {
            _retNode := Nodes_ReturnNodeDownCastF(node.FP)
            if _retNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                retNode = _retNode.(*Nodes_ReturnNode)
            }
        }
        var vals *LnsList2_[*WasmBin_Ins]
        vals = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
        {
            _expList := retNode.FP.Get_expList(_env)
            if !Lns_IsNil( _expList ) {
                expList := _expList.(*Nodes_ExpListNode)
                for _, _exp := range( expList.FP.Get_expList(_env).Items ) {
                    exp := _exp
                    vals.Insert(self.FP.lowerExp(_env, exp))
                }
            }
        }
        out.Insert(WasmBin_ret(_env, vals))
        return 
    } else if _switch0 == Nodes_NodeKind_get_Scope(_env) {
        var scopeNode *Nodes_ScopeNode
        
        {
            _scopeNode := Nodes_ScopeNodeDownCastF(node.FP)
            if _scopeNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                scopeNode = _scopeNode.(*Nodes_ScopeNode)
            }
        }
        for _, _ins := range( self.FP.lowerBlock(_env, scopeNode.FP.Get_block(_env)).Items ) {
            ins := _ins
            out.Insert(ins)
        }
        return 
    } else if _switch0 == Nodes_NodeKind_get_Block(_env) {
        var blockNode *Nodes_BlockNode
        
        {
            _blockNode := Nodes_BlockNodeDownCastF(node.FP)
            if _blockNode == nil{
                self.FP.error(_env, node, "cast")
            } else {
                blockNode = _blockNode.(*Nodes_BlockNode)
            }
        }
        for _, _ins := range( self.FP.lowerBlock(_env, blockNode).Items ) {
            ins := _ins
            out.Insert(ins)
        }
        return 
    }
    self.FP.notSupport(_env, node)
}
// 1139: decl @lune.@base.@convWasm.convFilter.registerFunc
func (self *convWasm_convFilter) registerFunc(_env *LnsEnv, node *Nodes_DeclFuncNode) *WasmBin_Func {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo(_env)
    if Lns_isCondTrue( declInfo.FP.Get_classTypeInfo(_env)){
        self.FP.error(_env, &node.Nodes_Node, "not support method")
    }
    var symbol *Ast_SymbolInfo
    
    {
        _symbol := declInfo.FP.Get_symbol(_env)
        if _symbol == nil{
            self.FP.error(_env, &node.Nodes_Node, "not support anonymous function")
        } else {
            symbol = _symbol.(*Ast_SymbolInfo)
        }
    }
    var params *LnsList2_[*WasmBin_ValType]
    params = NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{})
    for _, _arg := range( declInfo.FP.Get_argList(_env).Items ) {
        arg := _arg
        var argNode *Nodes_DeclArgNode
        
        {
            _argNode := Nodes_DeclArgNodeDownCastF(arg.FP)
            if _argNode == nil{
                self.FP.error(_env, arg, _env.GetVM().String_format("not support arg -- %s", Lns_2DDD(Nodes_getNodeKindName(_env, arg.FP.Get_kind(_env)))))
            } else {
                argNode = _argNode.(*Nodes_DeclArgNode)
            }
        }
        params.Insert(self.FP.getValType(_env, arg, self.FP.getKind(_env, arg, argNode.FP.Get_symbolInfo(_env).FP.Get_typeInfo(_env))))
    }
    var results *LnsList2_[*WasmBin_ValType]
    results = NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{})
    for _, _retType := range( declInfo.FP.Get_retTypeInfoList(_env).Items ) {
        retType := _retType
        results.Insert(self.FP.getValType(_env, &node.Nodes_Node, self.FP.getKind(_env, &node.Nodes_Node, retType)))
    }
    var index LnsInt
    var _func *WasmBin_Func
    index,_func = self.mod.FP.AddFunc(_env, params, results)
    self.sym2func.Set(symbol.FP.GetOrg(_env),index)
    return _func
}
// 1164: decl @lune.@base.@convWasm.convFilter.lowerFunc
func (self *convWasm_convFilter) lowerFunc(_env *LnsEnv, node *Nodes_DeclFuncNode,_func *WasmBin_Func) {
    var declInfo *Nodes_DeclFuncInfo
    declInfo = node.FP.Get_declInfo(_env)
    for _index, _arg := range( declInfo.FP.Get_argList(_env).Items ) {
        index := _index + 1
        arg := _arg
        var argNode *Nodes_DeclArgNode
        
        {
            _argNode := Nodes_DeclArgNodeDownCastF(arg.FP)
            if _argNode == nil{
                self.FP.error(_env, arg, "cast")
            } else {
                argNode = _argNode.(*Nodes_DeclArgNode)
            }
        }
        self.sym2local.Set(argNode.FP.Get_symbolInfo(_env).FP.GetOrg(_env),index - 1)
    }
    self.curFunc = _func
    self.depth = 1
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    {
        _block := declInfo.FP.Get_body(_env)
        if !Lns_IsNil( _block ) {
            block := _block.(*Nodes_BlockNode)
            for _, _ins := range( self.FP.lowerBlock(_env, block).Items ) {
                ins := _ins
                body.Insert(ins)
            }
        }
    }
    if declInfo.FP.Get_retTypeInfoList(_env).Len() > 0{
        body.Insert(WasmBin_unreachable(_env))
    }
    _func.FP.SetBody(_env, body)
    self.curFunc = nil
    self.depth = 0
}
// 1189: decl @lune.@base.@convWasm.convFilter.processRoot
func (self *convWasm_convFilter) ProcessRoot(_env *LnsEnv, node *Nodes_RootNode,_opt LnsAny) {
    self.FP.setupRuntime(_env)
    var funcList *LnsList2_[*Nodes_DeclFuncNode]
    funcList = NewLnsList2_[*Nodes_DeclFuncNode]([]*Nodes_DeclFuncNode{})
    var wasmFuncList *LnsList2_[*WasmBin_Func]
    wasmFuncList = NewLnsList2_[*WasmBin_Func]([]*WasmBin_Func{})
    for _, _child := range( node.FP.Get_children(_env).Items ) {
        child := _child
        if child.FP.Get_kind(_env) == Nodes_NodeKind_get_DeclFunc(_env){
            var declNode *Nodes_DeclFuncNode
            
            {
                _declNode := Nodes_DeclFuncNodeDownCastF(child.FP)
                if _declNode == nil{
                    self.FP.error(_env, child, "cast")
                } else {
                    declNode = _declNode.(*Nodes_DeclFuncNode)
                }
            }
            funcList.Insert(declNode)
            wasmFuncList.Insert(self.FP.registerFunc(_env, declNode))
        }
    }
    var mainIndex LnsInt
    var mainFunc *WasmBin_Func
    mainIndex,mainFunc = self.mod.FP.AddFunc(_env, NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
    self.curFunc = mainFunc
    self.depth = 0
    var body *LnsList2_[*WasmBin_Ins]
    body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    for _, _child := range( node.FP.Get_children(_env).Items ) {
        child := _child
        self.FP.lowerStmt(_env, child, body)
    }
    mainFunc.FP.SetBody(_env, body)
    self.curFunc = nil
    self.mod.FP.ExportFunc(_env, "main", mainIndex)
    for _index, _declNode := range( funcList.Items ) {
        index := _index + 1
        declNode := _declNode
        self.FP.lowerFunc(_env, declNode, wasmFuncList.GetAt(index))
    }
    self.stream.Write(_env, self.mod.FP.Encode(_env))
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
// 52: DeclConstr
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
// 58: DeclConstr
func (self *ConvWasm_Option) InitConvWasm_Option(_env *LnsEnv) {
}


// declaration Class -- convFilter
type convWasm_convFilterMtd interface {
    decodeString(_env *LnsEnv, arg1 *Nodes_Node, arg2 string) string
    DefaultProcess(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsAny)
    enterLoop(_env *LnsEnv)
    error(_env *LnsEnv, arg1 *Nodes_Node, arg2 string)
    exitLoop(_env *LnsEnv, arg1 *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins
    getCurFunc(_env *LnsEnv, arg1 *Nodes_Node) *WasmBin_Func
    getDDDArgs(_env *LnsEnv, arg1 *Nodes_ExpCallNode) *LnsList2_[*Nodes_Node]
    getDefaultVal(_env *LnsEnv, arg1 LnsInt) *WasmBin_Ins
    getExpKind(_env *LnsEnv, arg1 *Nodes_Node) LnsInt
    GetFull(_env *LnsEnv, arg1 *Ast_TypeInfo, arg2 bool) string
    getKind(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_TypeInfo) LnsInt
    getValType(_env *LnsEnv, arg1 *Nodes_Node, arg2 LnsInt) *WasmBin_ValType
    Get_moduleInfoManager(_env *LnsEnv) Ast_ModuleInfoManager
    Get_optStack(_env *LnsEnv) *LnsList
    Get_typeNameCtrl(_env *LnsEnv) *Ast_TypeNameCtrl
    lowerArgs(_env *LnsEnv, arg1 *Nodes_ExpCallNode) *LnsList2_[*WasmBin_Ins]
    lowerBlock(_env *LnsEnv, arg1 *Nodes_BlockNode) *LnsList2_[*WasmBin_Ins]
    lowerCall(_env *LnsEnv, arg1 *Nodes_ExpCallNode) *WasmBin_Ins
    lowerDeclVar(_env *LnsEnv, arg1 *Nodes_DeclVarNode, arg2 *LnsList2_[*WasmBin_Ins])
    lowerExp(_env *LnsEnv, arg1 *Nodes_Node) *WasmBin_Ins
    lowerFor(_env *LnsEnv, arg1 *Nodes_ForNode) *WasmBin_Ins
    lowerFunc(_env *LnsEnv, arg1 *Nodes_DeclFuncNode, arg2 *WasmBin_Func)
    lowerIf(_env *LnsEnv, arg1 *Nodes_IfNode) *WasmBin_Ins
    lowerOp1(_env *LnsEnv, arg1 *Nodes_ExpOp1Node) *WasmBin_Ins
    lowerOp2(_env *LnsEnv, arg1 *Nodes_ExpOp2Node) *WasmBin_Ins
    lowerPrint(_env *LnsEnv, arg1 *Nodes_ExpCallNode) *WasmBin_Ins
    lowerRef(_env *LnsEnv, arg1 *Nodes_ExpRefNode) *WasmBin_Ins
    lowerSetSym(_env *LnsEnv, arg1 *Nodes_Node, arg2 *Ast_SymbolInfo, arg3 *WasmBin_Ins) *WasmBin_Ins
    lowerStmt(_env *LnsEnv, arg1 *Nodes_Node, arg2 *LnsList2_[*WasmBin_Ins])
    notSupport(_env *LnsEnv, arg1 *Nodes_Node)
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
    registerFunc(_env *LnsEnv, arg1 *Nodes_DeclFuncNode) *WasmBin_Func
    setupRuntime(_env *LnsEnv)
    toReal(_env *LnsEnv, arg1 *Nodes_Node) *WasmBin_Ins
}
type convWasm_convFilter struct {
    Nodes_Filter
    streamName string
    stream Lns_oStream
    processInfo *Ast_ProcessInfo
    builtinFuncs *Builtin_BuiltinFuncType
    moduleScope *Ast_Scope
    mod *WasmBin_Module
    strTypeIdx LnsInt
    strType *WasmBin_ValType
    fnPrintI64 LnsInt
    fnPrintF64 LnsInt
    fnPrintBool LnsInt
    fnPrintStr LnsInt
    fnPrintNil LnsInt
    fnPrintSep LnsInt
    fnPrintEnd LnsInt
    fnStrConcat LnsInt
    fnStrEq LnsInt
    fnStrCmp LnsInt
    fnDivI64 LnsInt
    fnModI64 LnsInt
    fnModF64 LnsInt
    fnShiftL LnsInt
    fnShiftR LnsInt
    sym2func *LnsMap2_[*Ast_SymbolInfo,LnsInt]
    sym2global *LnsMap2_[*Ast_SymbolInfo,LnsInt]
    sym2local *LnsMap2_[*Ast_SymbolInfo,LnsInt]
    curFunc LnsAny
    depth LnsInt
    breakDepthList *LnsList2_[LnsInt]
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
func NewconvWasm_convFilter(_env *LnsEnv, arg1 string, arg2 Lns_oStream, arg3 *AstInfo_ASTInfo) *convWasm_convFilter {
    obj := &convWasm_convFilter{}
    obj.FP = obj
    obj.Nodes_Filter.FP = obj
    obj.InitconvWasm_convFilter(_env, arg1, arg2, arg3)
    return obj
}
// 179: DeclConstr
func (self *convWasm_convFilter) InitconvWasm_convFilter(_env *LnsEnv, streamName string,stream Lns_oStream,ast *AstInfo_ASTInfo) {
    self.InitNodes_Filter(_env, true, ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env), ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env))
    self.streamName = streamName
    self.stream = stream
    self.processInfo = ast.FP.Get_exportInfo(_env).FP.Get_processInfo(_env).FP.Clone(_env)
    self.builtinFuncs = ast.FP.Get_builtinFunc(_env)
    self.moduleScope = Lns_unwrap( ast.FP.Get_exportInfo(_env).FP.Get_moduleTypeInfo(_env).FP.Get_scope(_env)).(*Ast_Scope)
    var mod *WasmBin_Module
    mod = NewWasmBin_Module(_env)
    self.strTypeIdx = mod.FP.AddByteArrayType(_env)
    self.strType = NewWasmBin_ValType(_env, WasmBin_ValKind__Ref, self.strTypeIdx, true)
    self.fnPrintI64 = mod.FP.ImportFunc(_env, "lns", "print_i64", NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI64)), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
    self.fnPrintF64 = mod.FP.ImportFunc(_env, "lns", "print_f64", NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeF64)), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
    self.fnPrintBool = mod.FP.ImportFunc(_env, "lns", "print_bool", NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](WasmBin_typeI32)), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
    self.fnPrintStr = mod.FP.ImportFunc(_env, "lns", "print_str", NewLnsList2_[*WasmBin_ValType](Lns_2DDDGen[*WasmBin_ValType](self.strType)), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
    self.fnPrintNil = mod.FP.ImportFunc(_env, "lns", "print_nil", NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
    self.fnPrintSep = mod.FP.ImportFunc(_env, "lns", "print_sep", NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
    self.fnPrintEnd = mod.FP.ImportFunc(_env, "lns", "print_end", NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}), NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{}))
    self.mod = mod
    self.fnStrConcat = 0
    self.fnStrEq = 0
    self.fnStrCmp = 0
    self.fnDivI64 = 0
    self.fnModI64 = 0
    self.fnModF64 = 0
    self.fnShiftL = 0
    self.fnShiftR = 0
    self.sym2func = NewLnsMap2_[*Ast_SymbolInfo,LnsInt]( map[*Ast_SymbolInfo]LnsInt{})
    self.sym2global = NewLnsMap2_[*Ast_SymbolInfo,LnsInt]( map[*Ast_SymbolInfo]LnsInt{})
    self.sym2local = NewLnsMap2_[*Ast_SymbolInfo,LnsInt]( map[*Ast_SymbolInfo]LnsInt{})
    self.curFunc = nil
    self.depth = 0
    self.breakDepthList = NewLnsList2_[LnsInt]([]LnsInt{})
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
    convWasm_opI32Eqz = 0x45
    convWasm_opI32Eq = 0x46
    convWasm_opI32Ne = 0x47
    convWasm_opI32LtS = 0x48
    convWasm_opI32GtS = 0x4a
    convWasm_opI32LeS = 0x4c
    convWasm_opI32GeS = 0x4e
    convWasm_opI32Add = 0x6a
    convWasm_opI32Sub = 0x6b
    convWasm_opI32And = 0x71
    convWasm_opI32Or = 0x72
    convWasm_opI64Eqz = 0x50
    convWasm_opI64Eq = 0x51
    convWasm_opI64Ne = 0x52
    convWasm_opI64LtS = 0x53
    convWasm_opI64GtS = 0x55
    convWasm_opI64LeS = 0x57
    convWasm_opI64GeS = 0x59
    convWasm_opI64Add = 0x7c
    convWasm_opI64Sub = 0x7d
    convWasm_opI64Mul = 0x7e
    convWasm_opI64DivS = 0x7f
    convWasm_opI64RemS = 0x81
    convWasm_opI64And = 0x83
    convWasm_opI64Or = 0x84
    convWasm_opI64Xor = 0x85
    convWasm_opI64Shl = 0x86
    convWasm_opI64ShrU = 0x88
    convWasm_opF64Eq = 0x61
    convWasm_opF64Ne = 0x62
    convWasm_opF64Lt = 0x63
    convWasm_opF64Gt = 0x64
    convWasm_opF64Le = 0x65
    convWasm_opF64Ge = 0x66
    convWasm_opF64Neg = 0x9a
    convWasm_opF64Floor = 0x9c
    convWasm_opF64Add = 0xa0
    convWasm_opF64Sub = 0xa1
    convWasm_opF64Mul = 0xa2
    convWasm_opF64Div = 0xa3
    convWasm_opI64ExtendI32U = 0xad
    convWasm_opF64ConvertI64S = 0xb9
    convWasm_opI64TruncF64S = 0xb0
    convWasm_opRefNull = 0xd0
    convWasm_strCmpOpMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{"~=":convWasm_opI32Ne,"<":convWasm_opI32LtS,"<=":convWasm_opI32LeS,">":convWasm_opI32GtS,">=":convWasm_opI32GeS,})
    convWasm_i64CmpOpMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{"==":convWasm_opI64Eq,"~=":convWasm_opI64Ne,"<":convWasm_opI64LtS,"<=":convWasm_opI64LeS,">":convWasm_opI64GtS,">=":convWasm_opI64GeS,})
    convWasm_f64CmpOpMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{"==":convWasm_opF64Eq,"~=":convWasm_opF64Ne,"<":convWasm_opF64Lt,"<=":convWasm_opF64Le,">":convWasm_opF64Gt,">=":convWasm_opF64Ge,})
    convWasm_gcArrayNewData = 0x09
    convWasm_gcArrayNewDefault = 0x07
    convWasm_gcArrayGetU = 0x0d
    convWasm_gcArrayLen = 0x0f
    convWasm_gcArrayCopy = 0x11
}
func init() {
    init_convWasm = false
}
