// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_WasmBin bool
var WasmBin__mod__ string
// decl enum -- ValKind 
type WasmBin_ValKind = LnsInt
const WasmBin_ValKind__F64 = 2
const WasmBin_ValKind__I32 = 0
const WasmBin_ValKind__I64 = 1
const WasmBin_ValKind__Ref = 3
var WasmBin_ValKindList_ = NewLnsList( []LnsAny {
  WasmBin_ValKind__I32,
  WasmBin_ValKind__I64,
  WasmBin_ValKind__F64,
  WasmBin_ValKind__Ref,
})
func WasmBin_ValKind_get__allList(_env *LnsEnv) *LnsList{
    return WasmBin_ValKindList_
}
var WasmBin_ValKindMap_ = map[LnsInt]string {
  WasmBin_ValKind__F64: "ValKind.F64",
  WasmBin_ValKind__I32: "ValKind.I32",
  WasmBin_ValKind__I64: "ValKind.I64",
  WasmBin_ValKind__Ref: "ValKind.Ref",
}
func WasmBin_ValKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := WasmBin_ValKindMap_[arg1]; ok { return arg1 }
    return nil
}

func WasmBin_ValKind_getTxt(arg1 LnsInt) string {
    return WasmBin_ValKindMap_[arg1];
}
// decl enum -- InsKind 
type WasmBin_InsKind = LnsInt
const WasmBin_InsKind__Block = 2
const WasmBin_InsKind__FuncRef = 5
const WasmBin_InsKind__If = 4
const WasmBin_InsKind__Loop = 3
const WasmBin_InsKind__Plain = 0
const WasmBin_InsKind__Seq = 1
var WasmBin_InsKindList_ = NewLnsList( []LnsAny {
  WasmBin_InsKind__Plain,
  WasmBin_InsKind__Seq,
  WasmBin_InsKind__Block,
  WasmBin_InsKind__Loop,
  WasmBin_InsKind__If,
  WasmBin_InsKind__FuncRef,
})
func WasmBin_InsKind_get__allList(_env *LnsEnv) *LnsList{
    return WasmBin_InsKindList_
}
var WasmBin_InsKindMap_ = map[LnsInt]string {
  WasmBin_InsKind__Block: "InsKind.Block",
  WasmBin_InsKind__FuncRef: "InsKind.FuncRef",
  WasmBin_InsKind__If: "InsKind.If",
  WasmBin_InsKind__Loop: "InsKind.Loop",
  WasmBin_InsKind__Plain: "InsKind.Plain",
  WasmBin_InsKind__Seq: "InsKind.Seq",
}
func WasmBin_InsKind__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := WasmBin_InsKindMap_[arg1]; ok { return arg1 }
    return nil
}

func WasmBin_InsKind_getTxt(arg1 LnsInt) string {
    return WasmBin_InsKindMap_[arg1];
}
var WasmBin_htAny LnsInt
var WasmBin_htEq LnsInt
var WasmBin_htI31 LnsInt
var WasmBin_htStruct LnsInt
var WasmBin_htArray LnsInt
var WasmBin_htFunc LnsInt
var WasmBin_htExtern LnsInt
var WasmBin_htNone LnsInt
var WasmBin_typeI32 *WasmBin_ValType
var WasmBin_typeI64 *WasmBin_ValType
var WasmBin_typeF64 *WasmBin_ValType
var WasmBin_typeAnyRef *WasmBin_ValType
var WasmBin_typeEqRef *WasmBin_ValType
var WasmBin_opUnreachable LnsInt
var WasmBin_opNop LnsInt
var WasmBin_opDrop LnsInt
var WasmBin_opSelect LnsInt
var WasmBin_opI32Eqz LnsInt
var WasmBin_opI32Eq LnsInt
var WasmBin_opI32Ne LnsInt
var WasmBin_opI32LtS LnsInt
var WasmBin_opI32LtU LnsInt
var WasmBin_opI32GtS LnsInt
var WasmBin_opI32GtU LnsInt
var WasmBin_opI32LeS LnsInt
var WasmBin_opI32LeU LnsInt
var WasmBin_opI32GeS LnsInt
var WasmBin_opI32GeU LnsInt
var WasmBin_opI64Eqz LnsInt
var WasmBin_opI64Eq LnsInt
var WasmBin_opI64Ne LnsInt
var WasmBin_opI64LtS LnsInt
var WasmBin_opI64LtU LnsInt
var WasmBin_opI64GtS LnsInt
var WasmBin_opI64GtU LnsInt
var WasmBin_opI64LeS LnsInt
var WasmBin_opI64LeU LnsInt
var WasmBin_opI64GeS LnsInt
var WasmBin_opI64GeU LnsInt
var WasmBin_opF64Eq LnsInt
var WasmBin_opF64Ne LnsInt
var WasmBin_opF64Lt LnsInt
var WasmBin_opF64Gt LnsInt
var WasmBin_opF64Le LnsInt
var WasmBin_opF64Ge LnsInt
var WasmBin_opI32Add LnsInt
var WasmBin_opI32Sub LnsInt
var WasmBin_opI32Mul LnsInt
var WasmBin_opI32DivS LnsInt
var WasmBin_opI32DivU LnsInt
var WasmBin_opI32RemS LnsInt
var WasmBin_opI32RemU LnsInt
var WasmBin_opI32And LnsInt
var WasmBin_opI32Or LnsInt
var WasmBin_opI32Xor LnsInt
var WasmBin_opI32Shl LnsInt
var WasmBin_opI32ShrS LnsInt
var WasmBin_opI32ShrU LnsInt
var WasmBin_opI64Add LnsInt
var WasmBin_opI64Sub LnsInt
var WasmBin_opI64Mul LnsInt
var WasmBin_opI64DivS LnsInt
var WasmBin_opI64DivU LnsInt
var WasmBin_opI64RemS LnsInt
var WasmBin_opI64RemU LnsInt
var WasmBin_opI64And LnsInt
var WasmBin_opI64Or LnsInt
var WasmBin_opI64Xor LnsInt
var WasmBin_opI64Shl LnsInt
var WasmBin_opI64ShrS LnsInt
var WasmBin_opI64ShrU LnsInt
var WasmBin_opF64Abs LnsInt
var WasmBin_opF64Neg LnsInt
var WasmBin_opF64Ceil LnsInt
var WasmBin_opF64Floor LnsInt
var WasmBin_opF64Trunc LnsInt
var WasmBin_opF64Sqrt LnsInt
var WasmBin_opF64Add LnsInt
var WasmBin_opF64Sub LnsInt
var WasmBin_opF64Mul LnsInt
var WasmBin_opF64Div LnsInt
var WasmBin_opI32WrapI64 LnsInt
var WasmBin_opI64ExtendI32S LnsInt
var WasmBin_opI64ExtendI32U LnsInt
var WasmBin_opI64TruncF64S LnsInt
var WasmBin_opF64ConvertI64S LnsInt
var WasmBin_opI64ReinterpretF64 LnsInt
var WasmBin_opF64ReinterpretI64 LnsInt
var WasmBin_opRefIsNull LnsInt
var WasmBin_opRefEq LnsInt
// 38: decl @lune.@base.@WasmBin.byte
func WasmBin_byte(_env *LnsEnv, val LnsInt) string {
    return _env.GetVM().String_format("%c", Lns_2DDD(val & 0xff))
}

// 43: decl @lune.@base.@WasmBin.uleb
func WasmBin_uleb(_env *LnsEnv, val LnsInt) string {
    var work LnsInt
    work = val
    var txt string
    txt = ""
    for  {
        var bt LnsInt
        bt = work & 0x7f
        work = Lns_shiftR(work, 7)
        if work != 0{
            bt = bt | 0x80
        }
        txt = txt + WasmBin_byte(_env, bt)
        if work == 0{
            break
        }
    }
    return txt
}

// 61: decl @lune.@base.@WasmBin.sleb
func WasmBin_sleb(_env *LnsEnv, val LnsInt) string {
    var work LnsInt
    work = val
    var txt string
    txt = ""
    for  {
        var bt LnsInt
        bt = work & 0x7f
        work = Lns_shiftR(work, 7)
        if val < 0{
            work = work | (Lns_shiftL(-1, 57))
        }
        var done bool
        if _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( work == 0) &&
                _env.SetStackVal( bt & 0x40 == 0) ).(bool))) ||
            _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
                _env.SetStackVal( work == -1) &&
                _env.SetStackVal( bt & 0x40 != 0) ).(bool))) ).(bool){
            done = true
        } else { 
            done = false
        }
        if done{
            txt = txt + WasmBin_byte(_env, bt)
            break
        }
        txt = txt + WasmBin_byte(_env, bt | 0x80)
    }
    return txt
}

// 87: decl @lune.@base.@WasmBin.int64LE
func WasmBin_int64LE(_env *LnsEnv, val LnsInt) string {
    var work LnsInt
    work = val
    var txt string
    txt = ""
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = 8
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            txt = txt + WasmBin_byte(_env, work & 0xff)
            work = Lns_shiftR(work, 8)
        }
    }
    return txt
}

// 102: decl @lune.@base.@WasmBin.f64bits
func WasmBin_f64bits(_env *LnsEnv, val LnsReal) LnsInt {
    if val != val{
        return 0x7ff8000000000000
    }
    var work LnsReal
    work = val
    var sign LnsInt
    sign = 0
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( work < 0.0) ||
        _env.SetStackVal( (_env.PopVal( _env.IncStack() ||
            _env.SetStackVal( work == 0.0) &&
            _env.SetStackVal( 1.0 / work < 0.0) ).(bool))) ).(bool){
        sign = 1
        work = -work
    }
    if work == 0.0{
        return Lns_shiftL(sign, 63)
    }
    if work > 1.7976931348623157e308{
        return (Lns_shiftL(sign, 63)) | (Lns_shiftL(0x7ff, 52))
    }
    var exp LnsInt
    exp = 0
    for work >= 2.0 {
        work = work / 2.0
        exp = exp + 1
    }
    for Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( work < 1.0) &&
        _env.SetStackVal( exp > -1022) ).(bool)) {
        work = work * 2.0
        exp = exp - 1
    }
    var two52 LnsReal
    two52 = 4503599627370496.0
    if work < 1.0{
        var mant LnsInt
        mant = (LnsInt)((work * two52))
        return (Lns_shiftL(sign, 63)) | mant
    }
    var mant LnsInt
    mant = (LnsInt)(((work - 1.0) * two52))
    return (Lns_shiftL(sign, 63)) | (Lns_shiftL((exp + 1023), 52)) | mant
}

// 140: decl @lune.@base.@WasmBin.name
func WasmBin_name(_env *LnsEnv, txt string) string {
    return WasmBin_uleb(_env, len(txt)) + txt
}

// 145: decl @lune.@base.@WasmBin.vec
func WasmBin_vec(_env *LnsEnv, list *LnsList2_[string]) string {
    var txt string
    txt = WasmBin_uleb(_env, list.Len())
    for _, _item := range( list.Items ) {
        item := _item
        txt = txt + item
    }
    return txt
}

// 154: decl @lune.@base.@WasmBin.section
func WasmBin_section(_env *LnsEnv, id LnsInt,body string) string {
    return WasmBin_byte(_env, id) + WasmBin_uleb(_env, len(body)) + body
}

// 220: decl @lune.@base.@WasmBin.refType
func WasmBin_refType(_env *LnsEnv, heapType LnsInt) *WasmBin_ValType {
    return NewWasmBin_ValType(_env, WasmBin_ValKind__Ref, heapType, true)
}

// 247: decl @lune.@base.@WasmBin.field
func WasmBin_field(_env *LnsEnv, valType *WasmBin_ValType,mutable bool) *WasmBin_FieldType {
    return NewWasmBin_FieldType(_env, valType, false, mutable)
}

// 250: decl @lune.@base.@WasmBin.fieldI8
func WasmBin_fieldI8(_env *LnsEnv, mutable bool) *WasmBin_FieldType {
    return NewWasmBin_FieldType(_env, nil, true, mutable)
}

// 429: decl @lune.@base.@WasmBin.op
func WasmBin_op(_env *LnsEnv, code string,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__Plain, code, args, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), nil, nil, 0)
}

// 433: decl @lune.@base.@WasmBin.plain
func WasmBin_plain(_env *LnsEnv, code LnsInt,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__Plain, WasmBin_byte(_env, code), args, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), nil, nil, 0)
}

// 436: decl @lune.@base.@WasmBin.seq
func WasmBin_seq(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__Seq, "", NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), body, nil, nil, 0)
}

// 439: decl @lune.@base.@WasmBin.block
func WasmBin_block(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins],result LnsAny) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__Block, "", NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), body, nil, result, 0)
}

// 442: decl @lune.@base.@WasmBin.loop
func WasmBin_loop(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins],result LnsAny) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__Loop, "", NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), body, nil, result, 0)
}

// 445: decl @lune.@base.@WasmBin.ifElse
func WasmBin_ifElse(_env *LnsEnv, cond *WasmBin_Ins,thenBody *LnsList2_[*WasmBin_Ins],elseBody LnsAny,result LnsAny) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__If, "", NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cond)), thenBody, elseBody, result, 0)
}

// 451: decl @lune.@base.@WasmBin.i32Const
func WasmBin_i32Const(_env *LnsEnv, val LnsInt) *WasmBin_Ins {
    var work LnsInt
    work = val & 0xffffffff
    if work >= 0x80000000{
        work = work - 0x100000000
    }
    return WasmBin_op(_env, WasmBin_byte(_env, 0x41) + WasmBin_sleb(_env, work), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 459: decl @lune.@base.@WasmBin.i64Const
func WasmBin_i64Const(_env *LnsEnv, val LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x42) + WasmBin_sleb(_env, val), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 462: decl @lune.@base.@WasmBin.f64Const
func WasmBin_f64Const(_env *LnsEnv, val LnsReal) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x44) + WasmBin_int64LE(_env, WasmBin_f64bits(_env, val)), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 465: decl @lune.@base.@WasmBin.localGet
func WasmBin_localGet(_env *LnsEnv, index LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x20) + WasmBin_uleb(_env, index), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 468: decl @lune.@base.@WasmBin.localSet
func WasmBin_localSet(_env *LnsEnv, index LnsInt,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x21) + WasmBin_uleb(_env, index), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 471: decl @lune.@base.@WasmBin.localTee
func WasmBin_localTee(_env *LnsEnv, index LnsInt,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x22) + WasmBin_uleb(_env, index), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 474: decl @lune.@base.@WasmBin.globalGet
func WasmBin_globalGet(_env *LnsEnv, index LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x23) + WasmBin_uleb(_env, index), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 477: decl @lune.@base.@WasmBin.globalSet
func WasmBin_globalSet(_env *LnsEnv, index LnsInt,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x24) + WasmBin_uleb(_env, index), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 481: decl @lune.@base.@WasmBin.call
func WasmBin_call(_env *LnsEnv, funcId LnsInt,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__FuncRef, WasmBin_byte(_env, 0x10), args, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), nil, nil, funcId)
}

// 485: decl @lune.@base.@WasmBin.callRef
func WasmBin_callRef(_env *LnsEnv, funcTypeIdx LnsInt,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x14) + WasmBin_uleb(_env, funcTypeIdx), args)
}

// 488: decl @lune.@base.@WasmBin.br
func WasmBin_br(_env *LnsEnv, depth LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x0c) + WasmBin_uleb(_env, depth), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 491: decl @lune.@base.@WasmBin.brIf
func WasmBin_brIf(_env *LnsEnv, depth LnsInt,cond *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x0d) + WasmBin_uleb(_env, depth), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cond)))
}

// 494: decl @lune.@base.@WasmBin.ret
func WasmBin_ret(_env *LnsEnv, vals *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x0f), vals)
}

// 497: decl @lune.@base.@WasmBin.drop
func WasmBin_drop(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x1a), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 500: decl @lune.@base.@WasmBin.unreachable
func WasmBin_unreachable(_env *LnsEnv) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x00), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 503: decl @lune.@base.@WasmBin.nop
func WasmBin_nop(_env *LnsEnv) *WasmBin_Ins {
    return WasmBin_op(_env, "", NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 507: decl @lune.@base.@WasmBin.select
func WasmBin__select(_env *LnsEnv, val1 *WasmBin_Ins,val2 *WasmBin_Ins,cond *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x1b), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val1, val2, cond)))
}

// 510: decl @lune.@base.@WasmBin.refNull
func WasmBin_refNull(_env *LnsEnv, heapType LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0xd0) + WasmBin_sleb(_env, heapType), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 513: decl @lune.@base.@WasmBin.refIsNull
func WasmBin_refIsNull(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0xd1), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 516: decl @lune.@base.@WasmBin.refEq
func WasmBin_refEq(_env *LnsEnv, val1 *WasmBin_Ins,val2 *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0xd3), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val1, val2)))
}

// 520: decl @lune.@base.@WasmBin.gcOp
func WasmBin_gcOp(_env *LnsEnv, sub LnsInt,imm string,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0xfb) + WasmBin_uleb(_env, sub) + imm, args)
}

// 523: decl @lune.@base.@WasmBin.structNew
func WasmBin_structNew(_env *LnsEnv, typeIdx LnsInt,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x00, WasmBin_uleb(_env, typeIdx), args)
}

// 526: decl @lune.@base.@WasmBin.structNewDefault
func WasmBin_structNewDefault(_env *LnsEnv, typeIdx LnsInt) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x01, WasmBin_uleb(_env, typeIdx), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 529: decl @lune.@base.@WasmBin.structGet
func WasmBin_structGet(_env *LnsEnv, typeIdx LnsInt,fieldIdx LnsInt,obj *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x02, WasmBin_uleb(_env, typeIdx) + WasmBin_uleb(_env, fieldIdx), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](obj)))
}

// 532: decl @lune.@base.@WasmBin.structSet
func WasmBin_structSet(_env *LnsEnv, typeIdx LnsInt,fieldIdx LnsInt,obj *WasmBin_Ins,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x05, WasmBin_uleb(_env, typeIdx) + WasmBin_uleb(_env, fieldIdx), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](obj, val)))
}

// 535: decl @lune.@base.@WasmBin.arrayNew
func WasmBin_arrayNew(_env *LnsEnv, typeIdx LnsInt,init *WasmBin_Ins,len *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x06, WasmBin_uleb(_env, typeIdx), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](init, len)))
}

// 538: decl @lune.@base.@WasmBin.arrayNewDefault
func WasmBin_arrayNewDefault(_env *LnsEnv, typeIdx LnsInt,len *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x07, WasmBin_uleb(_env, typeIdx), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](len)))
}

// 541: decl @lune.@base.@WasmBin.arrayNewFixed
func WasmBin_arrayNewFixed(_env *LnsEnv, typeIdx LnsInt,vals *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x08, WasmBin_uleb(_env, typeIdx) + WasmBin_uleb(_env, vals.Len()), vals)
}

// 544: decl @lune.@base.@WasmBin.arrayNewData
func WasmBin_arrayNewData(_env *LnsEnv, typeIdx LnsInt,dataIdx LnsInt,offset *WasmBin_Ins,len *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x09, WasmBin_uleb(_env, typeIdx) + WasmBin_uleb(_env, dataIdx), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](offset, len)))
}

// 547: decl @lune.@base.@WasmBin.arrayGet
func WasmBin_arrayGet(_env *LnsEnv, typeIdx LnsInt,arr *WasmBin_Ins,index *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x0b, WasmBin_uleb(_env, typeIdx), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](arr, index)))
}

// 550: decl @lune.@base.@WasmBin.arrayGetU
func WasmBin_arrayGetU(_env *LnsEnv, typeIdx LnsInt,arr *WasmBin_Ins,index *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x0d, WasmBin_uleb(_env, typeIdx), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](arr, index)))
}

// 553: decl @lune.@base.@WasmBin.arraySet
func WasmBin_arraySet(_env *LnsEnv, typeIdx LnsInt,arr *WasmBin_Ins,index *WasmBin_Ins,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x0e, WasmBin_uleb(_env, typeIdx), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](arr, index, val)))
}

// 556: decl @lune.@base.@WasmBin.arrayLen
func WasmBin_arrayLen(_env *LnsEnv, arr *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x0f, "", NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](arr)))
}

// 559: decl @lune.@base.@WasmBin.arrayCopy
func WasmBin_arrayCopy(_env *LnsEnv, dstType LnsInt,srcType LnsInt,dst *WasmBin_Ins,dstIndex *WasmBin_Ins,src *WasmBin_Ins,srcIndex *WasmBin_Ins,len *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x11, WasmBin_uleb(_env, dstType) + WasmBin_uleb(_env, srcType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](dst, dstIndex, src, srcIndex, len)))
}

// 566: decl @lune.@base.@WasmBin.refTest
func WasmBin_refTest(_env *LnsEnv, heapType LnsInt,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x14, WasmBin_sleb(_env, heapType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 570: decl @lune.@base.@WasmBin.refCast
func WasmBin_refCast(_env *LnsEnv, heapType LnsInt,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_gcOp(_env, 0x17, WasmBin_sleb(_env, heapType), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 575: decl @lune.@base.@WasmBin.i32Load8U
func WasmBin_i32Load8U(_env *LnsEnv, addr *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x2d) + WasmBin_uleb(_env, 0) + WasmBin_uleb(_env, 0), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](addr)))
}

// 579: decl @lune.@base.@WasmBin.i32Store8
func WasmBin_i32Store8(_env *LnsEnv, addr *WasmBin_Ins,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x3a) + WasmBin_uleb(_env, 0) + WasmBin_uleb(_env, 0), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](addr, val)))
}

// 668: decl @lune.@base.@WasmBin.encodeFuncType
func WasmBin_encodeFuncType_62_(_env *LnsEnv, params *LnsList2_[*WasmBin_ValType],results *LnsList2_[*WasmBin_ValType]) string {
    var paramList *LnsList2_[string]
    paramList = NewLnsList2_[string]([]string{})
    for _, _param := range( params.Items ) {
        param := _param
        paramList.Insert(param.FP.Encode(_env))
    }
    var resultList *LnsList2_[string]
    resultList = NewLnsList2_[string]([]string{})
    for _, _result := range( results.Items ) {
        result := _result
        resultList.Insert(result.FP.Encode(_env))
    }
    return WasmBin_byte(_env, 0x60) + WasmBin_vec(_env, paramList) + WasmBin_vec(_env, resultList)
}

// 185: decl @lune.@base.@WasmBin.ValType.encode
func (self *WasmBin_ValType) Encode(_env *LnsEnv) string {
    if _switch0 := self.kind; _switch0 == WasmBin_ValKind__I32 {
        return WasmBin_byte(_env, 0x7f)
    } else if _switch0 == WasmBin_ValKind__I64 {
        return WasmBin_byte(_env, 0x7e)
    } else if _switch0 == WasmBin_ValKind__F64 {
        return WasmBin_byte(_env, 0x7c)
    } else if _switch0 == WasmBin_ValKind__Ref {
        if self.nullable{
            return WasmBin_byte(_env, 0x63) + WasmBin_sleb(_env, self.heapType)
        }
        return WasmBin_byte(_env, 0x64) + WasmBin_sleb(_env, self.heapType)
    }
// insert a dummy
    return ""
}
// 205: decl @lune.@base.@WasmBin.ValType.isRef
func (self *WasmBin_ValType) IsRef(_env *LnsEnv) bool {
    return self.kind == WasmBin_ValKind__Ref
}
// 209: decl @lune.@base.@WasmBin.ValType.equalsType
func (self *WasmBin_ValType) EqualsType(_env *LnsEnv, other *WasmBin_ValType) bool {
    return _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.kind == other.kind) &&
        _env.SetStackVal( self.heapType == other.heapType) &&
        _env.SetStackVal( self.nullable == other.nullable) ).(bool)
}
// 233: decl @lune.@base.@WasmBin.FieldType.encode
func (self *WasmBin_FieldType) Encode(_env *LnsEnv) string {
    var txt string
    {
        _valType := self.valType
        if !Lns_IsNil( _valType ) {
            valType := _valType.(*WasmBin_ValType)
            txt = valType.FP.Encode(_env)
        } else {
            txt = WasmBin_byte(_env, 0x78)
        }
    }
    if self.mutable{
        return txt + WasmBin_byte(_env, 0x01)
    }
    return txt + WasmBin_byte(_env, 0x00)
}
// 377: decl @lune.@base.@WasmBin.Ins.encodeTo
func (self *WasmBin_Ins) EncodeTo(_env *LnsEnv, stream Lns_oStream,funcIndexList *LnsList2_[LnsInt]) {
    for _, _arg := range( self.args.Items ) {
        arg := _arg
        arg.FP.EncodeTo(_env, stream, funcIndexList)
    }
    var blockType string
    {
        _result := self.result
        if !Lns_IsNil( _result ) {
            result := _result.(*WasmBin_ValType)
            blockType = result.FP.Encode(_env)
        } else {
            blockType = WasmBin_byte(_env, 0x40)
        }
    }
    if _switch0 := self.kind; _switch0 == WasmBin_InsKind__Plain {
        stream.Write(_env, self.code)
    } else if _switch0 == WasmBin_InsKind__FuncRef {
        stream.Write(_env, self.code + WasmBin_uleb(_env, funcIndexList.GetAt(self.funcId + 1)))
    } else if _switch0 == WasmBin_InsKind__Seq {
        for _, _ins := range( self.body.Items ) {
            ins := _ins
            ins.FP.EncodeTo(_env, stream, funcIndexList)
        }
    } else if _switch0 == WasmBin_InsKind__Block || _switch0 == WasmBin_InsKind__Loop {
        if self.kind == WasmBin_InsKind__Block{
            stream.Write(_env, WasmBin_byte(_env, 0x02))
        } else { 
            stream.Write(_env, WasmBin_byte(_env, 0x03))
        }
        stream.Write(_env, blockType)
        for _, _ins := range( self.body.Items ) {
            ins := _ins
            ins.FP.EncodeTo(_env, stream, funcIndexList)
        }
        stream.Write(_env, WasmBin_byte(_env, 0x0b))
    } else if _switch0 == WasmBin_InsKind__If {
        stream.Write(_env, WasmBin_byte(_env, 0x04) + blockType)
        for _, _ins := range( self.body.Items ) {
            ins := _ins
            ins.FP.EncodeTo(_env, stream, funcIndexList)
        }
        {
            _elseBody := self.elseBody
            if !Lns_IsNil( _elseBody ) {
                elseBody := _elseBody.(*LnsList2_[*WasmBin_Ins])
                stream.Write(_env, WasmBin_byte(_env, 0x05))
                for _, _ins := range( elseBody.Items ) {
                    ins := _ins
                    ins.FP.EncodeTo(_env, stream, funcIndexList)
                }
            }
        }
        stream.Write(_env, WasmBin_byte(_env, 0x0b))
    }
}
// 591: decl @lune.@base.@WasmBin.memStream.get_txt
func (self *WasmBin_memStream) Get_txt(_env *LnsEnv) string {
    self.FP.Flush(_env)
    var txt string
    txt = ""
    for _, _item := range( self.list.Items ) {
        item := _item
        txt = txt + item
    }
    self.list = NewLnsList2_[string](Lns_2DDDGen[string](txt))
    return txt
}
// 600: decl @lune.@base.@WasmBin.memStream.write
func (self *WasmBin_memStream) Write(_env *LnsEnv, val string)(LnsAny, LnsAny) {
    self.work = self.work + val
    if len(self.work) > 4000{
        self.FP.Flush(_env)
    }
    return self.FP, nil
}
// 607: decl @lune.@base.@WasmBin.memStream.close
func (self *WasmBin_memStream) Close(_env *LnsEnv) {
    self.FP.Flush(_env)
}
// 610: decl @lune.@base.@WasmBin.memStream.flush
func (self *WasmBin_memStream) Flush(_env *LnsEnv) {
    if len(self.work) > 0{
        self.list.Insert(self.work)
        self.work = ""
    }
}
// 637: decl @lune.@base.@WasmBin.Func.addLocal
func (self *WasmBin_Func) AddLocal(_env *LnsEnv, valType *WasmBin_ValType) LnsInt {
    self.localTypes.Insert(valType)
    return self.paramNum + self.localTypes.Len() - 1
}
// 642: decl @lune.@base.@WasmBin.Func.setBody
func (self *WasmBin_Func) SetBody(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins]) {
    self.body = body
    self.hasBody = true
}
// 647: decl @lune.@base.@WasmBin.Func.encode
func (self *WasmBin_Func) Encode(_env *LnsEnv, funcIndexList *LnsList2_[LnsInt]) string {
    var localList *LnsList2_[string]
    localList = NewLnsList2_[string]([]string{})
    for _, _localType := range( self.localTypes.Items ) {
        localType := _localType
        localList.Insert(WasmBin_uleb(_env, 1) + localType.FP.Encode(_env))
    }
    var stream *WasmBin_memStream
    stream = NewWasmBin_memStream(_env)
    stream.FP.Write(_env, WasmBin_vec(_env, localList))
    for _, _ins := range( self.body.Items ) {
        ins := _ins
        ins.FP.EncodeTo(_env, stream.FP, funcIndexList)
    }
    if Lns_op_not(self.hasBody){
        stream.FP.Write(_env, WasmBin_byte(_env, 0x00))
    }
    stream.FP.Write(_env, WasmBin_byte(_env, 0x0b))
    var txt string
    txt = stream.FP.Get_txt(_env)
    return WasmBin_uleb(_env, len(txt)) + txt
}
// 735: decl @lune.@base.@WasmBin.Module.setFuncName
func (self *WasmBin_Module) SetFuncName(_env *LnsEnv, funcId LnsInt,funcName string) {
    self.funcNameMap.Set(funcId,funcName)
}
// 740: decl @lune.@base.@WasmBin.Module.reserveType
func (self *WasmBin_Module) ReserveType(_env *LnsEnv) LnsInt {
    self.typeList.Insert("")
    return self.typeList.Len() - 1
}
// 744: decl @lune.@base.@WasmBin.Module.defineType
func (self *WasmBin_Module) defineType(_env *LnsEnv, index LnsInt,superIdx LnsAny,comp string) {
    var superList *LnsList2_[string]
    superList = NewLnsList2_[string]([]string{})
    if superIdx != nil{
        superIdx_23 := superIdx.(LnsInt)
        superList.Insert(WasmBin_uleb(_env, superIdx_23))
    }
    self.typeList.Set(index + 1,WasmBin_byte(_env, 0x50) + WasmBin_vec(_env, superList) + comp)
}
// 752: decl @lune.@base.@WasmBin.Module.defineStructType
func (self *WasmBin_Module) DefineStructType(_env *LnsEnv, index LnsInt,superIdx LnsAny,fields *LnsList2_[*WasmBin_FieldType]) {
    var list *LnsList2_[string]
    list = NewLnsList2_[string]([]string{})
    for _, _fieldType := range( fields.Items ) {
        fieldType := _fieldType
        list.Insert(fieldType.FP.Encode(_env))
    }
    self.FP.defineType(_env, index, superIdx, WasmBin_byte(_env, 0x5f) + WasmBin_vec(_env, list))
}
// 759: decl @lune.@base.@WasmBin.Module.addStructType
func (self *WasmBin_Module) AddStructType(_env *LnsEnv, superIdx LnsAny,fields *LnsList2_[*WasmBin_FieldType]) LnsInt {
    var index LnsInt
    index = self.FP.ReserveType(_env)
    self.FP.DefineStructType(_env, index, superIdx, fields)
    return index
}
// 764: decl @lune.@base.@WasmBin.Module.addArrayType
func (self *WasmBin_Module) AddArrayType(_env *LnsEnv, fieldType *WasmBin_FieldType) LnsInt {
    var index LnsInt
    index = self.FP.ReserveType(_env)
    self.FP.defineType(_env, index, nil, WasmBin_byte(_env, 0x5e) + fieldType.FP.Encode(_env))
    return index
}
// 771: decl @lune.@base.@WasmBin.Module.getFuncType
func (self *WasmBin_Module) GetFuncType(_env *LnsEnv, params *LnsList2_[*WasmBin_ValType],results *LnsList2_[*WasmBin_ValType]) LnsInt {
    var encoded string
    encoded = WasmBin_encodeFuncType_62_(_env, params, results)
    {
        _index := self.funcTypeMap.Get(encoded)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    index = self.FP.ReserveType(_env)
    self.FP.defineType(_env, index, nil, encoded)
    self.funcTypeMap.Set(encoded,index)
    return index
}
// 789: decl @lune.@base.@WasmBin.Module.importFunc
func (self *WasmBin_Module) ImportFunc(_env *LnsEnv, modName string,funcName string,params *LnsList2_[*WasmBin_ValType],results *LnsList2_[*WasmBin_ValType]) LnsInt {
    self.importList.Insert(WasmBin_name(_env, modName) + WasmBin_name(_env, funcName))
    self.importTypeList.Insert(WasmBin_encodeFuncType_62_(_env, params, results))
    var funcId LnsInt
    funcId = self.nextFuncId
    self.nextFuncId = self.nextFuncId + 1
    self.importIdList.Insert(LnsInt(funcId))
    return funcId
}
// 803: decl @lune.@base.@WasmBin.Module.importMemory
func (self *WasmBin_Module) ImportMemory(_env *LnsEnv, modName string,memName string) {
    self.importMemList.Insert(WasmBin_name(_env, modName) + WasmBin_name(_env, memName) + WasmBin_byte(_env, 0x02) + WasmBin_byte(_env, 0x00) + WasmBin_uleb(_env, 0))
}
// 810: decl @lune.@base.@WasmBin.Module.addFunc
func (self *WasmBin_Module) AddFunc(_env *LnsEnv, params *LnsList2_[*WasmBin_ValType],results *LnsList2_[*WasmBin_ValType])(LnsInt, *WasmBin_Func) {
    var typeIdx LnsInt
    typeIdx = self.FP.GetFuncType(_env, params, results)
    return self.FP.AddFuncWithType(_env, typeIdx, params.Len())
}
// 815: decl @lune.@base.@WasmBin.Module.addFuncWithType
func (self *WasmBin_Module) AddFuncWithType(_env *LnsEnv, typeIdx LnsInt,paramNum LnsInt)(LnsInt, *WasmBin_Func) {
    var _func *WasmBin_Func
    _func = NewWasmBin_Func(_env, typeIdx, paramNum)
    var funcId LnsInt
    funcId = self.nextFuncId
    self.nextFuncId = self.nextFuncId + 1
    self.funcList.Insert(_func)
    self.funcIdList.Insert(LnsInt(funcId))
    self.id2func.Set(funcId,_func)
    return funcId, _func
}
// 826: decl @lune.@base.@WasmBin.Module.getFunc
func (self *WasmBin_Module) GetFunc(_env *LnsEnv, funcId LnsInt) *WasmBin_Func {
    return Lns_unwrap( self.id2func.Get(funcId)).(*WasmBin_Func)
}
// 831: decl @lune.@base.@WasmBin.Module.refFunc
func (self *WasmBin_Module) RefFunc(_env *LnsEnv, funcId LnsInt) *WasmBin_Ins {
    if Lns_op_not(self.refFuncSet.Has(LnsInt(funcId))){
        self.refFuncSet.Add(LnsInt(funcId))
        self.refFuncList.Insert(LnsInt(funcId))
    }
    return NewWasmBin_Ins(_env, WasmBin_InsKind__FuncRef, WasmBin_byte(_env, 0xd2), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), nil, nil, funcId)
}
// 840: decl @lune.@base.@WasmBin.Module.addGlobal
func (self *WasmBin_Module) AddGlobal(_env *LnsEnv, valType *WasmBin_ValType,init *WasmBin_Ins) LnsInt {
    self.globalList.Insert(valType)
    self.globalInitList.Insert(init)
    return self.globalList.Len() - 1
}
// 846: decl @lune.@base.@WasmBin.Module.exportFunc
func (self *WasmBin_Module) ExportFunc(_env *LnsEnv, exportName string,funcId LnsInt) {
    self.exportNameList.Insert(exportName)
    self.exportIdList.Insert(LnsInt(funcId))
}
// 852: decl @lune.@base.@WasmBin.Module.addData
func (self *WasmBin_Module) AddData(_env *LnsEnv, data string) LnsInt {
    self.dataList.Insert(WasmBin_byte(_env, 0x01) + WasmBin_uleb(_env, len(data)) + data)
    return self.dataList.Len() - 1
}
// 857: decl @lune.@base.@WasmBin.Module.encode
func (self *WasmBin_Module) Encode(_env *LnsEnv) string {
    var funcIndexList *LnsList2_[LnsInt]
    funcIndexList = NewLnsList2_[LnsInt]([]LnsInt{})
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = self.nextFuncId
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            funcIndexList.Insert(LnsInt(0))
        }
    }
    for _index, _funcId := range( self.importIdList.Items ) {
        index := _index + 1
        funcId := _funcId
        funcIndexList.Set(funcId + 1,index - 1)
    }
    for _index, _funcId := range( self.funcIdList.Items ) {
        index := _index + 1
        funcId := _funcId
        funcIndexList.Set(funcId + 1,self.importIdList.Len() + index - 1)
    }
    var stream *WasmBin_memStream
    stream = NewWasmBin_memStream(_env)
    stream.FP.Write(_env, _env.GetVM().String_format("%sasm%s", Lns_2DDD(WasmBin_byte(_env, 0), WasmBin_byte(_env, 1) + WasmBin_byte(_env, 0) + WasmBin_byte(_env, 0) + WasmBin_byte(_env, 0))))
    var typeEntryList *LnsList2_[string]
    typeEntryList = NewLnsList2_[string](Lns_2DDDGen[string](WasmBin_byte(_env, 0x4e) + WasmBin_vec(_env, self.typeList)))
    var importTypeMap *LnsMap2_[string,LnsInt]
    importTypeMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    var importList *LnsList2_[string]
    importList = NewLnsList2_[string]([]string{})
    for _index, _entry := range( self.importList.Items ) {
        index := _index + 1
        entry := _entry
        var funcType string
        funcType = self.importTypeList.GetAt(index)
        var typeIdx LnsInt
        {
            _work := importTypeMap.Get(funcType)
            if !Lns_IsNil( _work ) {
                work := _work.(LnsInt)
                typeIdx = work
            } else {
                typeIdx = self.typeList.Len() + typeEntryList.Len() - 1
                importTypeMap.Set(funcType,typeIdx)
                typeEntryList.Insert(funcType)
            }
        }
        importList.Insert(entry + WasmBin_byte(_env, 0x00) + WasmBin_uleb(_env, typeIdx))
    }
    stream.FP.Write(_env, WasmBin_section(_env, 1, WasmBin_vec(_env, typeEntryList)))
    for _, _entry := range( self.importMemList.Items ) {
        entry := _entry
        importList.Insert(entry)
    }
    if importList.Len() > 0{
        stream.FP.Write(_env, WasmBin_section(_env, 2, WasmBin_vec(_env, importList)))
    }
    var funcTypeList *LnsList2_[string]
    funcTypeList = NewLnsList2_[string]([]string{})
    var codeList *LnsList2_[string]
    codeList = NewLnsList2_[string]([]string{})
    for _, __func := range( self.funcList.Items ) {
        _func := __func
        funcTypeList.Insert(WasmBin_uleb(_env, _func.FP.Get_typeIdx(_env)))
        codeList.Insert(_func.FP.Encode(_env, funcIndexList))
    }
    stream.FP.Write(_env, WasmBin_section(_env, 3, WasmBin_vec(_env, funcTypeList)))
    if self.globalList.Len() > 0{
        var globalList *LnsList2_[string]
        globalList = NewLnsList2_[string]([]string{})
        for _index, _valType := range( self.globalList.Items ) {
            index := _index + 1
            valType := _valType
            var initStream *WasmBin_memStream
            initStream = NewWasmBin_memStream(_env)
            self.globalInitList.GetAt(index).FP.EncodeTo(_env, initStream.FP, funcIndexList)
            globalList.Insert(valType.FP.Encode(_env) + WasmBin_byte(_env, 0x01) + initStream.FP.Get_txt(_env) + WasmBin_byte(_env, 0x0b))
        }
        stream.FP.Write(_env, WasmBin_section(_env, 6, WasmBin_vec(_env, globalList)))
    }
    if self.exportNameList.Len() > 0{
        var exportList *LnsList2_[string]
        exportList = NewLnsList2_[string]([]string{})
        for _index, _exportName := range( self.exportNameList.Items ) {
            index := _index + 1
            exportName := _exportName
            exportList.Insert(WasmBin_name(_env, exportName) + WasmBin_byte(_env, 0x00) + WasmBin_uleb(_env, funcIndexList.GetAt(self.exportIdList.GetAt(index) + 1)))
        }
        stream.FP.Write(_env, WasmBin_section(_env, 7, WasmBin_vec(_env, exportList)))
    }
    if self.refFuncList.Len() > 0{
        var idxList *LnsList2_[string]
        idxList = NewLnsList2_[string]([]string{})
        for _, _funcId := range( self.refFuncList.Items ) {
            funcId := _funcId
            idxList.Insert(WasmBin_uleb(_env, funcIndexList.GetAt(funcId + 1)))
        }
        stream.FP.Write(_env, WasmBin_section(_env, 9, WasmBin_uleb(_env, 1) + WasmBin_byte(_env, 0x03) + WasmBin_byte(_env, 0x00) + WasmBin_vec(_env, idxList)))
    }
    if self.dataList.Len() > 0{
        stream.FP.Write(_env, WasmBin_section(_env, 12, WasmBin_uleb(_env, self.dataList.Len())))
    }
    stream.FP.Write(_env, WasmBin_section(_env, 10, WasmBin_vec(_env, codeList)))
    if self.dataList.Len() > 0{
        stream.FP.Write(_env, WasmBin_section(_env, 11, WasmBin_vec(_env, self.dataList)))
    }
    var index2name *LnsMap2_[LnsInt,string]
    index2name = NewLnsMap2_[LnsInt,string]( map[LnsInt]string{})
    for _funcId, _funcName := range( self.funcNameMap.Items ) {
        funcId := _funcId
        funcName := _funcName
        index2name.Set(funcIndexList.GetAt(funcId + 1),funcName)
    }
    var nameList *LnsList2_[string]
    nameList = NewLnsList2_[string]([]string{})
    {
        __forsortCollection0 := index2name
        __forsortSorted0 := __forsortCollection0.CreateKeyListInt()
        __forsortSorted0.Sort( _env, LnsItemKindInt, nil )
        for _, _funcIndex := range( __forsortSorted0.Items ) {
            funcName := __forsortCollection0.Items[ _funcIndex ]
            funcIndex := _funcIndex
            nameList.Insert(WasmBin_uleb(_env, funcIndex) + WasmBin_name(_env, funcName))
        }
    }
    if nameList.Len() > 0{
        var funcNames string
        funcNames = WasmBin_vec(_env, nameList)
        stream.FP.Write(_env, WasmBin_section(_env, 0, WasmBin_name(_env, "name") + WasmBin_byte(_env, 1) + WasmBin_uleb(_env, len(funcNames)) + funcNames))
    }
    return stream.FP.Get_txt(_env)
}
// declaration Class -- ValType
type WasmBin_ValTypeMtd interface {
    Encode(_env *LnsEnv) string
    EqualsType(_env *LnsEnv, arg1 *WasmBin_ValType) bool
    Get_heapType(_env *LnsEnv) LnsInt
    Get_kind(_env *LnsEnv) LnsInt
    Get_nullable(_env *LnsEnv) bool
    IsRef(_env *LnsEnv) bool
}
type WasmBin_ValType struct {
    kind LnsInt
    heapType LnsInt
    nullable bool
    FP WasmBin_ValTypeMtd
}
func WasmBin_ValType2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*WasmBin_ValType).FP
}
func WasmBin_ValType_toSlice(slice []LnsAny) []*WasmBin_ValType {
    ret := make([]*WasmBin_ValType, len(slice))
    for index, val := range slice {
        ret[index] = val.(WasmBin_ValTypeDownCast).ToWasmBin_ValType()
    }
    return ret
}
type WasmBin_ValTypeDownCast interface {
    ToWasmBin_ValType() *WasmBin_ValType
}
func WasmBin_ValTypeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(WasmBin_ValTypeDownCast)
    if ok { return work.ToWasmBin_ValType() }
    return nil
}
func (obj *WasmBin_ValType) ToWasmBin_ValType() *WasmBin_ValType {
    return obj
}
func NewWasmBin_ValType(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 bool) *WasmBin_ValType {
    obj := &WasmBin_ValType{}
    obj.FP = obj
    obj.InitWasmBin_ValType(_env, arg1, arg2, arg3)
    return obj
}
func (self *WasmBin_ValType) InitWasmBin_ValType(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt, arg3 bool) {
    self.kind = arg1
    self.heapType = arg2
    self.nullable = arg3
}
func (self *WasmBin_ValType) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *WasmBin_ValType) Get_heapType(_env *LnsEnv) LnsInt{ return self.heapType }
func (self *WasmBin_ValType) Get_nullable(_env *LnsEnv) bool{ return self.nullable }

// declaration Class -- FieldType
type WasmBin_FieldTypeMtd interface {
    Encode(_env *LnsEnv) string
    Get_mutable(_env *LnsEnv) bool
    Get_packedI8(_env *LnsEnv) bool
    Get_valType(_env *LnsEnv) LnsAny
}
type WasmBin_FieldType struct {
    valType LnsAny
    packedI8 bool
    mutable bool
    FP WasmBin_FieldTypeMtd
}
func WasmBin_FieldType2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*WasmBin_FieldType).FP
}
func WasmBin_FieldType_toSlice(slice []LnsAny) []*WasmBin_FieldType {
    ret := make([]*WasmBin_FieldType, len(slice))
    for index, val := range slice {
        ret[index] = val.(WasmBin_FieldTypeDownCast).ToWasmBin_FieldType()
    }
    return ret
}
type WasmBin_FieldTypeDownCast interface {
    ToWasmBin_FieldType() *WasmBin_FieldType
}
func WasmBin_FieldTypeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(WasmBin_FieldTypeDownCast)
    if ok { return work.ToWasmBin_FieldType() }
    return nil
}
func (obj *WasmBin_FieldType) ToWasmBin_FieldType() *WasmBin_FieldType {
    return obj
}
func NewWasmBin_FieldType(_env *LnsEnv, arg1 LnsAny, arg2 bool, arg3 bool) *WasmBin_FieldType {
    obj := &WasmBin_FieldType{}
    obj.FP = obj
    obj.InitWasmBin_FieldType(_env, arg1, arg2, arg3)
    return obj
}
func (self *WasmBin_FieldType) InitWasmBin_FieldType(_env *LnsEnv, arg1 LnsAny, arg2 bool, arg3 bool) {
    self.valType = arg1
    self.packedI8 = arg2
    self.mutable = arg3
}
func (self *WasmBin_FieldType) Get_valType(_env *LnsEnv) LnsAny{ return self.valType }
func (self *WasmBin_FieldType) Get_packedI8(_env *LnsEnv) bool{ return self.packedI8 }
func (self *WasmBin_FieldType) Get_mutable(_env *LnsEnv) bool{ return self.mutable }

// declaration Class -- Ins
type WasmBin_InsMtd interface {
    EncodeTo(_env *LnsEnv, arg1 Lns_oStream, arg2 *LnsList2_[LnsInt])
    Get_args(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]
    Get_body(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]
    Get_code(_env *LnsEnv) string
    Get_elseBody(_env *LnsEnv) LnsAny
    Get_funcId(_env *LnsEnv) LnsInt
    Get_kind(_env *LnsEnv) LnsInt
    Get_result(_env *LnsEnv) LnsAny
}
type WasmBin_Ins struct {
    kind LnsInt
    code string
    args *LnsList2_[*WasmBin_Ins]
    body *LnsList2_[*WasmBin_Ins]
    elseBody LnsAny
    result LnsAny
    funcId LnsInt
    FP WasmBin_InsMtd
}
func WasmBin_Ins2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*WasmBin_Ins).FP
}
func WasmBin_Ins_toSlice(slice []LnsAny) []*WasmBin_Ins {
    ret := make([]*WasmBin_Ins, len(slice))
    for index, val := range slice {
        ret[index] = val.(WasmBin_InsDownCast).ToWasmBin_Ins()
    }
    return ret
}
type WasmBin_InsDownCast interface {
    ToWasmBin_Ins() *WasmBin_Ins
}
func WasmBin_InsDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(WasmBin_InsDownCast)
    if ok { return work.ToWasmBin_Ins() }
    return nil
}
func (obj *WasmBin_Ins) ToWasmBin_Ins() *WasmBin_Ins {
    return obj
}
func NewWasmBin_Ins(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 *LnsList2_[*WasmBin_Ins], arg4 *LnsList2_[*WasmBin_Ins], arg5 LnsAny, arg6 LnsAny, arg7 LnsInt) *WasmBin_Ins {
    obj := &WasmBin_Ins{}
    obj.FP = obj
    obj.InitWasmBin_Ins(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *WasmBin_Ins) InitWasmBin_Ins(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 *LnsList2_[*WasmBin_Ins], arg4 *LnsList2_[*WasmBin_Ins], arg5 LnsAny, arg6 LnsAny, arg7 LnsInt) {
    self.kind = arg1
    self.code = arg2
    self.args = arg3
    self.body = arg4
    self.elseBody = arg5
    self.result = arg6
    self.funcId = arg7
}
func (self *WasmBin_Ins) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *WasmBin_Ins) Get_code(_env *LnsEnv) string{ return self.code }
func (self *WasmBin_Ins) Get_args(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]{ return self.args }
func (self *WasmBin_Ins) Get_body(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]{ return self.body }
func (self *WasmBin_Ins) Get_elseBody(_env *LnsEnv) LnsAny{ return self.elseBody }
func (self *WasmBin_Ins) Get_result(_env *LnsEnv) LnsAny{ return self.result }
func (self *WasmBin_Ins) Get_funcId(_env *LnsEnv) LnsInt{ return self.funcId }

// declaration Class -- memStream
type WasmBin_memStreamMtd interface {
    Close(_env *LnsEnv)
    Flush(_env *LnsEnv)
    Get_txt(_env *LnsEnv) string
    Write(_env *LnsEnv, arg1 string)(LnsAny, LnsAny)
}
type WasmBin_memStream struct {
    list *LnsList2_[string]
    work string
    FP WasmBin_memStreamMtd
}
func WasmBin_memStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*WasmBin_memStream).FP
}
func WasmBin_memStream_toSlice(slice []LnsAny) []*WasmBin_memStream {
    ret := make([]*WasmBin_memStream, len(slice))
    for index, val := range slice {
        ret[index] = val.(WasmBin_memStreamDownCast).ToWasmBin_memStream()
    }
    return ret
}
type WasmBin_memStreamDownCast interface {
    ToWasmBin_memStream() *WasmBin_memStream
}
func WasmBin_memStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(WasmBin_memStreamDownCast)
    if ok { return work.ToWasmBin_memStream() }
    return nil
}
func (obj *WasmBin_memStream) ToWasmBin_memStream() *WasmBin_memStream {
    return obj
}
func NewWasmBin_memStream(_env *LnsEnv) *WasmBin_memStream {
    obj := &WasmBin_memStream{}
    obj.FP = obj
    obj.InitWasmBin_memStream(_env)
    return obj
}
// 587: DeclConstr
func (self *WasmBin_memStream) InitWasmBin_memStream(_env *LnsEnv) {
    self.list = NewLnsList2_[string]([]string{})
    self.work = ""
}


// declaration Class -- Func
type WasmBin_FuncMtd interface {
    AddLocal(_env *LnsEnv, arg1 *WasmBin_ValType) LnsInt
    Encode(_env *LnsEnv, arg1 *LnsList2_[LnsInt]) string
    Get_body(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]
    Get_hasBody(_env *LnsEnv) bool
    Get_localTypes(_env *LnsEnv) *LnsList2_[*WasmBin_ValType]
    Get_paramNum(_env *LnsEnv) LnsInt
    Get_typeIdx(_env *LnsEnv) LnsInt
    SetBody(_env *LnsEnv, arg1 *LnsList2_[*WasmBin_Ins])
}
type WasmBin_Func struct {
    typeIdx LnsInt
    localTypes *LnsList2_[*WasmBin_ValType]
    body *LnsList2_[*WasmBin_Ins]
    paramNum LnsInt
    hasBody bool
    FP WasmBin_FuncMtd
}
func WasmBin_Func2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*WasmBin_Func).FP
}
func WasmBin_Func_toSlice(slice []LnsAny) []*WasmBin_Func {
    ret := make([]*WasmBin_Func, len(slice))
    for index, val := range slice {
        ret[index] = val.(WasmBin_FuncDownCast).ToWasmBin_Func()
    }
    return ret
}
type WasmBin_FuncDownCast interface {
    ToWasmBin_Func() *WasmBin_Func
}
func WasmBin_FuncDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(WasmBin_FuncDownCast)
    if ok { return work.ToWasmBin_Func() }
    return nil
}
func (obj *WasmBin_Func) ToWasmBin_Func() *WasmBin_Func {
    return obj
}
func NewWasmBin_Func(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt) *WasmBin_Func {
    obj := &WasmBin_Func{}
    obj.FP = obj
    obj.InitWasmBin_Func(_env, arg1, arg2)
    return obj
}
func (self *WasmBin_Func) Get_typeIdx(_env *LnsEnv) LnsInt{ return self.typeIdx }
func (self *WasmBin_Func) Get_localTypes(_env *LnsEnv) *LnsList2_[*WasmBin_ValType]{ return self.localTypes }
func (self *WasmBin_Func) Get_body(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]{ return self.body }
func (self *WasmBin_Func) Get_paramNum(_env *LnsEnv) LnsInt{ return self.paramNum }
func (self *WasmBin_Func) Get_hasBody(_env *LnsEnv) bool{ return self.hasBody }
// 628: DeclConstr
func (self *WasmBin_Func) InitWasmBin_Func(_env *LnsEnv, typeIdx LnsInt,paramNum LnsInt) {
    self.typeIdx = typeIdx
    self.paramNum = paramNum
    self.localTypes = NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{})
    self.body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    self.hasBody = false
}


// declaration Class -- Module
type WasmBin_ModuleMtd interface {
    AddArrayType(_env *LnsEnv, arg1 *WasmBin_FieldType) LnsInt
    AddData(_env *LnsEnv, arg1 string) LnsInt
    AddFunc(_env *LnsEnv, arg1 *LnsList2_[*WasmBin_ValType], arg2 *LnsList2_[*WasmBin_ValType])(LnsInt, *WasmBin_Func)
    AddFuncWithType(_env *LnsEnv, arg1 LnsInt, arg2 LnsInt)(LnsInt, *WasmBin_Func)
    AddGlobal(_env *LnsEnv, arg1 *WasmBin_ValType, arg2 *WasmBin_Ins) LnsInt
    AddStructType(_env *LnsEnv, arg1 LnsAny, arg2 *LnsList2_[*WasmBin_FieldType]) LnsInt
    DefineStructType(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 *LnsList2_[*WasmBin_FieldType])
    defineType(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny, arg3 string)
    Encode(_env *LnsEnv) string
    ExportFunc(_env *LnsEnv, arg1 string, arg2 LnsInt)
    GetFunc(_env *LnsEnv, arg1 LnsInt) *WasmBin_Func
    GetFuncType(_env *LnsEnv, arg1 *LnsList2_[*WasmBin_ValType], arg2 *LnsList2_[*WasmBin_ValType]) LnsInt
    ImportFunc(_env *LnsEnv, arg1 string, arg2 string, arg3 *LnsList2_[*WasmBin_ValType], arg4 *LnsList2_[*WasmBin_ValType]) LnsInt
    ImportMemory(_env *LnsEnv, arg1 string, arg2 string)
    RefFunc(_env *LnsEnv, arg1 LnsInt) *WasmBin_Ins
    ReserveType(_env *LnsEnv) LnsInt
    SetFuncName(_env *LnsEnv, arg1 LnsInt, arg2 string)
}
type WasmBin_Module struct {
    typeList *LnsList2_[string]
    funcTypeMap *LnsMap2_[string,LnsInt]
    importList *LnsList2_[string]
    importTypeList *LnsList2_[string]
    importIdList *LnsList2_[LnsInt]
    importMemList *LnsList2_[string]
    funcList *LnsList2_[*WasmBin_Func]
    funcIdList *LnsList2_[LnsInt]
    id2func *LnsMap2_[LnsInt,*WasmBin_Func]
    nextFuncId LnsInt
    globalList *LnsList2_[*WasmBin_ValType]
    globalInitList *LnsList2_[*WasmBin_Ins]
    exportNameList *LnsList2_[string]
    exportIdList *LnsList2_[LnsInt]
    dataList *LnsList2_[string]
    funcNameMap *LnsMap2_[LnsInt,string]
    refFuncSet *LnsSet2_[LnsInt]
    refFuncList *LnsList2_[LnsInt]
    FP WasmBin_ModuleMtd
}
func WasmBin_Module2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*WasmBin_Module).FP
}
func WasmBin_Module_toSlice(slice []LnsAny) []*WasmBin_Module {
    ret := make([]*WasmBin_Module, len(slice))
    for index, val := range slice {
        ret[index] = val.(WasmBin_ModuleDownCast).ToWasmBin_Module()
    }
    return ret
}
type WasmBin_ModuleDownCast interface {
    ToWasmBin_Module() *WasmBin_Module
}
func WasmBin_ModuleDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(WasmBin_ModuleDownCast)
    if ok { return work.ToWasmBin_Module() }
    return nil
}
func (obj *WasmBin_Module) ToWasmBin_Module() *WasmBin_Module {
    return obj
}
func NewWasmBin_Module(_env *LnsEnv) *WasmBin_Module {
    obj := &WasmBin_Module{}
    obj.FP = obj
    obj.InitWasmBin_Module(_env)
    return obj
}
// 713: DeclConstr
func (self *WasmBin_Module) InitWasmBin_Module(_env *LnsEnv) {
    self.typeList = NewLnsList2_[string]([]string{})
    self.funcTypeMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    self.importList = NewLnsList2_[string]([]string{})
    self.importTypeList = NewLnsList2_[string]([]string{})
    self.importIdList = NewLnsList2_[LnsInt]([]LnsInt{})
    self.importMemList = NewLnsList2_[string]([]string{})
    self.funcList = NewLnsList2_[*WasmBin_Func]([]*WasmBin_Func{})
    self.funcIdList = NewLnsList2_[LnsInt]([]LnsInt{})
    self.id2func = NewLnsMap2_[LnsInt,*WasmBin_Func]( map[LnsInt]*WasmBin_Func{})
    self.nextFuncId = 0
    self.globalList = NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{})
    self.globalInitList = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
    self.exportNameList = NewLnsList2_[string]([]string{})
    self.exportIdList = NewLnsList2_[LnsInt]([]LnsInt{})
    self.dataList = NewLnsList2_[string]([]string{})
    self.refFuncSet = NewLnsSet2_[LnsInt]([]LnsInt{})
    self.refFuncList = NewLnsList2_[LnsInt]([]LnsInt{})
    self.funcNameMap = NewLnsMap2_[LnsInt,string]( map[LnsInt]string{})
}


func Lns_WasmBin_init(_env *LnsEnv) {
    if init_WasmBin { return }
    init_WasmBin = true
    WasmBin__mod__ = "@lune.@base.@WasmBin"
    Lns_InitMod()
    WasmBin_htAny = -18
    WasmBin_htEq = -19
    WasmBin_htI31 = -20
    WasmBin_htStruct = -21
    WasmBin_htArray = -22
    WasmBin_htFunc = -16
    WasmBin_htExtern = -17
    WasmBin_htNone = -15
    WasmBin_typeI32 = NewWasmBin_ValType(_env, WasmBin_ValKind__I32, 0, false)
    WasmBin_typeI64 = NewWasmBin_ValType(_env, WasmBin_ValKind__I64, 0, false)
    WasmBin_typeF64 = NewWasmBin_ValType(_env, WasmBin_ValKind__F64, 0, false)
    WasmBin_typeAnyRef = WasmBin_refType(_env, WasmBin_htAny)
    WasmBin_typeEqRef = WasmBin_refType(_env, WasmBin_htEq)
    WasmBin_opUnreachable = 0x00
    WasmBin_opNop = 0x01
    WasmBin_opDrop = 0x1a
    WasmBin_opSelect = 0x1b
    WasmBin_opI32Eqz = 0x45
    WasmBin_opI32Eq = 0x46
    WasmBin_opI32Ne = 0x47
    WasmBin_opI32LtS = 0x48
    WasmBin_opI32LtU = 0x49
    WasmBin_opI32GtS = 0x4a
    WasmBin_opI32GtU = 0x4b
    WasmBin_opI32LeS = 0x4c
    WasmBin_opI32LeU = 0x4d
    WasmBin_opI32GeS = 0x4e
    WasmBin_opI32GeU = 0x4f
    WasmBin_opI64Eqz = 0x50
    WasmBin_opI64Eq = 0x51
    WasmBin_opI64Ne = 0x52
    WasmBin_opI64LtS = 0x53
    WasmBin_opI64LtU = 0x54
    WasmBin_opI64GtS = 0x55
    WasmBin_opI64GtU = 0x56
    WasmBin_opI64LeS = 0x57
    WasmBin_opI64LeU = 0x58
    WasmBin_opI64GeS = 0x59
    WasmBin_opI64GeU = 0x5a
    WasmBin_opF64Eq = 0x61
    WasmBin_opF64Ne = 0x62
    WasmBin_opF64Lt = 0x63
    WasmBin_opF64Gt = 0x64
    WasmBin_opF64Le = 0x65
    WasmBin_opF64Ge = 0x66
    WasmBin_opI32Add = 0x6a
    WasmBin_opI32Sub = 0x6b
    WasmBin_opI32Mul = 0x6c
    WasmBin_opI32DivS = 0x6d
    WasmBin_opI32DivU = 0x6e
    WasmBin_opI32RemS = 0x6f
    WasmBin_opI32RemU = 0x70
    WasmBin_opI32And = 0x71
    WasmBin_opI32Or = 0x72
    WasmBin_opI32Xor = 0x73
    WasmBin_opI32Shl = 0x74
    WasmBin_opI32ShrS = 0x75
    WasmBin_opI32ShrU = 0x76
    WasmBin_opI64Add = 0x7c
    WasmBin_opI64Sub = 0x7d
    WasmBin_opI64Mul = 0x7e
    WasmBin_opI64DivS = 0x7f
    WasmBin_opI64DivU = 0x80
    WasmBin_opI64RemS = 0x81
    WasmBin_opI64RemU = 0x82
    WasmBin_opI64And = 0x83
    WasmBin_opI64Or = 0x84
    WasmBin_opI64Xor = 0x85
    WasmBin_opI64Shl = 0x86
    WasmBin_opI64ShrS = 0x87
    WasmBin_opI64ShrU = 0x88
    WasmBin_opF64Abs = 0x99
    WasmBin_opF64Neg = 0x9a
    WasmBin_opF64Ceil = 0x9b
    WasmBin_opF64Floor = 0x9c
    WasmBin_opF64Trunc = 0x9d
    WasmBin_opF64Sqrt = 0x9f
    WasmBin_opF64Add = 0xa0
    WasmBin_opF64Sub = 0xa1
    WasmBin_opF64Mul = 0xa2
    WasmBin_opF64Div = 0xa3
    WasmBin_opI32WrapI64 = 0xa7
    WasmBin_opI64ExtendI32S = 0xac
    WasmBin_opI64ExtendI32U = 0xad
    WasmBin_opI64TruncF64S = 0xb0
    WasmBin_opF64ConvertI64S = 0xb9
    WasmBin_opI64ReinterpretF64 = 0xbd
    WasmBin_opF64ReinterpretI64 = 0xbf
    WasmBin_opRefIsNull = 0xd1
    WasmBin_opRefEq = 0xd3
}
func init() {
    init_WasmBin = false
}
