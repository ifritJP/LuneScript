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
})
func WasmBin_InsKind_get__allList(_env *LnsEnv) *LnsList{
    return WasmBin_InsKindList_
}
var WasmBin_InsKindMap_ = map[LnsInt]string {
  WasmBin_InsKind__Block: "InsKind.Block",
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
var WasmBin_typeI32 *WasmBin_ValType
var WasmBin_typeI64 *WasmBin_ValType
var WasmBin_typeF64 *WasmBin_ValType
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

// 282: decl @lune.@base.@WasmBin.op
func WasmBin_op(_env *LnsEnv, code string,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__Plain, code, args, NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), nil, nil)
}

// 285: decl @lune.@base.@WasmBin.seq
func WasmBin_seq(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__Seq, "", NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), body, nil, nil)
}

// 288: decl @lune.@base.@WasmBin.block
func WasmBin_block(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins],result LnsAny) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__Block, "", NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), body, nil, result)
}

// 291: decl @lune.@base.@WasmBin.loop
func WasmBin_loop(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins],result LnsAny) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__Loop, "", NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}), body, nil, result)
}

// 294: decl @lune.@base.@WasmBin.ifElse
func WasmBin_ifElse(_env *LnsEnv, cond *WasmBin_Ins,thenBody *LnsList2_[*WasmBin_Ins],elseBody LnsAny,result LnsAny) *WasmBin_Ins {
    return NewWasmBin_Ins(_env, WasmBin_InsKind__If, "", NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cond)), thenBody, elseBody, result)
}

// 300: decl @lune.@base.@WasmBin.i32Const
func WasmBin_i32Const(_env *LnsEnv, val LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x41) + WasmBin_sleb(_env, val), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 303: decl @lune.@base.@WasmBin.i64Const
func WasmBin_i64Const(_env *LnsEnv, val LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x42) + WasmBin_sleb(_env, val), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 306: decl @lune.@base.@WasmBin.f64Const
func WasmBin_f64Const(_env *LnsEnv, val LnsReal) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x44) + WasmBin_int64LE(_env, WasmBin_f64bits(_env, val)), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 309: decl @lune.@base.@WasmBin.localGet
func WasmBin_localGet(_env *LnsEnv, index LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x20) + WasmBin_uleb(_env, index), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 312: decl @lune.@base.@WasmBin.localSet
func WasmBin_localSet(_env *LnsEnv, index LnsInt,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x21) + WasmBin_uleb(_env, index), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 315: decl @lune.@base.@WasmBin.globalGet
func WasmBin_globalGet(_env *LnsEnv, index LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x23) + WasmBin_uleb(_env, index), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 318: decl @lune.@base.@WasmBin.globalSet
func WasmBin_globalSet(_env *LnsEnv, index LnsInt,val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x24) + WasmBin_uleb(_env, index), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 321: decl @lune.@base.@WasmBin.call
func WasmBin_call(_env *LnsEnv, funcIndex LnsInt,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x10) + WasmBin_uleb(_env, funcIndex), args)
}

// 324: decl @lune.@base.@WasmBin.br
func WasmBin_br(_env *LnsEnv, depth LnsInt) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x0c) + WasmBin_uleb(_env, depth), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 327: decl @lune.@base.@WasmBin.brIf
func WasmBin_brIf(_env *LnsEnv, depth LnsInt,cond *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x0d) + WasmBin_uleb(_env, depth), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](cond)))
}

// 330: decl @lune.@base.@WasmBin.ret
func WasmBin_ret(_env *LnsEnv, vals *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x0f), vals)
}

// 333: decl @lune.@base.@WasmBin.drop
func WasmBin_drop(_env *LnsEnv, val *WasmBin_Ins) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x1a), NewLnsList2_[*WasmBin_Ins](Lns_2DDDGen[*WasmBin_Ins](val)))
}

// 336: decl @lune.@base.@WasmBin.unreachable
func WasmBin_unreachable(_env *LnsEnv) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0x00), NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{}))
}

// 340: decl @lune.@base.@WasmBin.gcOp
func WasmBin_gcOp(_env *LnsEnv, sub LnsInt,imm string,args *LnsList2_[*WasmBin_Ins]) *WasmBin_Ins {
    return WasmBin_op(_env, WasmBin_byte(_env, 0xfb) + WasmBin_uleb(_env, sub) + imm, args)
}

// 175: decl @lune.@base.@WasmBin.ValType.encode
func (self *WasmBin_ValType) Encode(_env *LnsEnv) string {
    if _switch0 := self.kind; _switch0 == WasmBin_ValKind__I32 {
        return WasmBin_byte(_env, 0x7f)
    } else if _switch0 == WasmBin_ValKind__I64 {
        return WasmBin_byte(_env, 0x7e)
    } else if _switch0 == WasmBin_ValKind__F64 {
        return WasmBin_byte(_env, 0x7c)
    } else if _switch0 == WasmBin_ValKind__Ref {
        if self.nullable{
            return WasmBin_byte(_env, 0x63) + WasmBin_sleb(_env, self.typeIdx)
        }
        return WasmBin_byte(_env, 0x64) + WasmBin_sleb(_env, self.typeIdx)
    }
// insert a dummy
    return ""
}
// 231: decl @lune.@base.@WasmBin.Ins.encode
func (self *WasmBin_Ins) Encode(_env *LnsEnv) string {
    var txt string
    txt = ""
    for _, _arg := range( self.args.Items ) {
        arg := _arg
        txt = txt + arg.FP.Encode(_env)
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
        txt = txt + self.code
    } else if _switch0 == WasmBin_InsKind__Seq {
        for _, _ins := range( self.body.Items ) {
            ins := _ins
            txt = txt + ins.FP.Encode(_env)
        }
    } else if _switch0 == WasmBin_InsKind__Block || _switch0 == WasmBin_InsKind__Loop {
        if self.kind == WasmBin_InsKind__Block{
            txt = txt + WasmBin_byte(_env, 0x02)
        } else { 
            txt = txt + WasmBin_byte(_env, 0x03)
        }
        txt = txt + blockType
        for _, _ins := range( self.body.Items ) {
            ins := _ins
            txt = txt + ins.FP.Encode(_env)
        }
        txt = txt + WasmBin_byte(_env, 0x0b)
    } else if _switch0 == WasmBin_InsKind__If {
        txt = txt + WasmBin_byte(_env, 0x04) + blockType
        for _, _ins := range( self.body.Items ) {
            ins := _ins
            txt = txt + ins.FP.Encode(_env)
        }
        {
            _elseBody := self.elseBody
            if !Lns_IsNil( _elseBody ) {
                elseBody := _elseBody.(*LnsList2_[*WasmBin_Ins])
                txt = txt + WasmBin_byte(_env, 0x05)
                for _, _ins := range( elseBody.Items ) {
                    ins := _ins
                    txt = txt + ins.FP.Encode(_env)
                }
            }
        }
        txt = txt + WasmBin_byte(_env, 0x0b)
    }
    return txt
}
// 360: decl @lune.@base.@WasmBin.Func.addLocal
func (self *WasmBin_Func) AddLocal(_env *LnsEnv, valType *WasmBin_ValType) LnsInt {
    self.localTypes.Insert(valType)
    return self.paramNum + self.localTypes.Len() - 1
}
// 365: decl @lune.@base.@WasmBin.Func.setBody
func (self *WasmBin_Func) SetBody(_env *LnsEnv, body *LnsList2_[*WasmBin_Ins]) {
    self.body = body
}
// 369: decl @lune.@base.@WasmBin.Func.encode
func (self *WasmBin_Func) Encode(_env *LnsEnv) string {
    var localList *LnsList2_[string]
    localList = NewLnsList2_[string]([]string{})
    for _, _localType := range( self.localTypes.Items ) {
        localType := _localType
        localList.Insert(WasmBin_uleb(_env, 1) + localType.FP.Encode(_env))
    }
    var txt string
    txt = WasmBin_vec(_env, localList)
    for _, _ins := range( self.body.Items ) {
        ins := _ins
        txt = txt + ins.FP.Encode(_env)
    }
    txt = txt + WasmBin_byte(_env, 0x0b)
    return WasmBin_uleb(_env, len(txt)) + txt
}
// 408: decl @lune.@base.@WasmBin.Module.addType
func (self *WasmBin_Module) AddType(_env *LnsEnv, encoded string) LnsInt {
    self.typeList.Insert(encoded)
    return self.typeList.Len() - 1
}
// 413: decl @lune.@base.@WasmBin.Module.addByteArrayType
func (self *WasmBin_Module) AddByteArrayType(_env *LnsEnv) LnsInt {
    return self.FP.AddType(_env, WasmBin_byte(_env, 0x5e) + WasmBin_byte(_env, 0x78) + WasmBin_byte(_env, 0x01))
}
// 418: decl @lune.@base.@WasmBin.Module.getFuncType
func (self *WasmBin_Module) GetFuncType(_env *LnsEnv, params *LnsList2_[*WasmBin_ValType],results *LnsList2_[*WasmBin_ValType]) LnsInt {
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
    var encoded string
    encoded = WasmBin_byte(_env, 0x60) + WasmBin_vec(_env, paramList) + WasmBin_vec(_env, resultList)
    {
        _index := self.funcTypeMap.Get(encoded)
        if !Lns_IsNil( _index ) {
            index := _index.(LnsInt)
            return index
        }
    }
    var index LnsInt
    index = self.FP.AddType(_env, encoded)
    self.funcTypeMap.Set(encoded,index)
    return index
}
// 441: decl @lune.@base.@WasmBin.Module.importFunc
func (self *WasmBin_Module) ImportFunc(_env *LnsEnv, modName string,funcName string,params *LnsList2_[*WasmBin_ValType],results *LnsList2_[*WasmBin_ValType]) LnsInt {
    if self.funcList.Len() != 0{
        panic("import after func")
    }
    var typeIdx LnsInt
    typeIdx = self.FP.GetFuncType(_env, params, results)
    self.importList.Insert(WasmBin_name(_env, modName) + WasmBin_name(_env, funcName) + WasmBin_byte(_env, 0x00) + WasmBin_uleb(_env, typeIdx))
    self.importFuncNum = self.importFuncNum + 1
    return self.importFuncNum - 1
}
// 455: decl @lune.@base.@WasmBin.Module.addFunc
func (self *WasmBin_Module) AddFunc(_env *LnsEnv, params *LnsList2_[*WasmBin_ValType],results *LnsList2_[*WasmBin_ValType])(LnsInt, *WasmBin_Func) {
    var typeIdx LnsInt
    typeIdx = self.FP.GetFuncType(_env, params, results)
    var _func *WasmBin_Func
    _func = NewWasmBin_Func(_env, typeIdx, params.Len())
    self.funcList.Insert(_func)
    return self.importFuncNum + self.funcList.Len() - 1, _func
}
// 463: decl @lune.@base.@WasmBin.Module.addGlobal
func (self *WasmBin_Module) AddGlobal(_env *LnsEnv, valType *WasmBin_ValType,init *WasmBin_Ins) LnsInt {
    self.globalList.Insert(valType.FP.Encode(_env) + WasmBin_byte(_env, 0x01) + init.FP.Encode(_env) + WasmBin_byte(_env, 0x0b))
    return self.globalList.Len() - 1
}
// 468: decl @lune.@base.@WasmBin.Module.exportFunc
func (self *WasmBin_Module) ExportFunc(_env *LnsEnv, exportName string,funcIndex LnsInt) {
    self.exportList.Insert(WasmBin_name(_env, exportName) + WasmBin_byte(_env, 0x00) + WasmBin_uleb(_env, funcIndex))
}
// 473: decl @lune.@base.@WasmBin.Module.addData
func (self *WasmBin_Module) AddData(_env *LnsEnv, data string) LnsInt {
    self.dataList.Insert(WasmBin_byte(_env, 0x01) + WasmBin_uleb(_env, len(data)) + data)
    return self.dataList.Len() - 1
}
// 478: decl @lune.@base.@WasmBin.Module.encode
func (self *WasmBin_Module) Encode(_env *LnsEnv) string {
    var txt string
    txt = _env.GetVM().String_format("%sasm%s", Lns_2DDD(WasmBin_byte(_env, 0), WasmBin_byte(_env, 1) + WasmBin_byte(_env, 0) + WasmBin_byte(_env, 0) + WasmBin_byte(_env, 0)))
    txt = txt + WasmBin_section(_env, 1, WasmBin_vec(_env, self.typeList))
    if self.importList.Len() > 0{
        txt = txt + WasmBin_section(_env, 2, WasmBin_vec(_env, self.importList))
    }
    var funcTypeList *LnsList2_[string]
    funcTypeList = NewLnsList2_[string]([]string{})
    var codeList *LnsList2_[string]
    codeList = NewLnsList2_[string]([]string{})
    for _, __func := range( self.funcList.Items ) {
        _func := __func
        funcTypeList.Insert(WasmBin_uleb(_env, _func.FP.Get_typeIdx(_env)))
        codeList.Insert(_func.FP.Encode(_env))
    }
    txt = txt + WasmBin_section(_env, 3, WasmBin_vec(_env, funcTypeList))
    if self.globalList.Len() > 0{
        txt = txt + WasmBin_section(_env, 6, WasmBin_vec(_env, self.globalList))
    }
    if self.exportList.Len() > 0{
        txt = txt + WasmBin_section(_env, 7, WasmBin_vec(_env, self.exportList))
    }
    if self.dataList.Len() > 0{
        txt = txt + WasmBin_section(_env, 12, WasmBin_uleb(_env, self.dataList.Len()))
    }
    txt = txt + WasmBin_section(_env, 10, WasmBin_vec(_env, codeList))
    if self.dataList.Len() > 0{
        txt = txt + WasmBin_section(_env, 11, WasmBin_vec(_env, self.dataList))
    }
    return txt
}
// declaration Class -- ValType
type WasmBin_ValTypeMtd interface {
    Encode(_env *LnsEnv) string
    Get_kind(_env *LnsEnv) LnsInt
    Get_nullable(_env *LnsEnv) bool
    Get_typeIdx(_env *LnsEnv) LnsInt
}
type WasmBin_ValType struct {
    kind LnsInt
    typeIdx LnsInt
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
    self.typeIdx = arg2
    self.nullable = arg3
}
func (self *WasmBin_ValType) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *WasmBin_ValType) Get_typeIdx(_env *LnsEnv) LnsInt{ return self.typeIdx }
func (self *WasmBin_ValType) Get_nullable(_env *LnsEnv) bool{ return self.nullable }

// declaration Class -- Ins
type WasmBin_InsMtd interface {
    Encode(_env *LnsEnv) string
    Get_args(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]
    Get_body(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]
    Get_code(_env *LnsEnv) string
    Get_elseBody(_env *LnsEnv) LnsAny
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
func NewWasmBin_Ins(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 *LnsList2_[*WasmBin_Ins], arg4 *LnsList2_[*WasmBin_Ins], arg5 LnsAny, arg6 LnsAny) *WasmBin_Ins {
    obj := &WasmBin_Ins{}
    obj.FP = obj
    obj.InitWasmBin_Ins(_env, arg1, arg2, arg3, arg4, arg5, arg6)
    return obj
}
func (self *WasmBin_Ins) InitWasmBin_Ins(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 *LnsList2_[*WasmBin_Ins], arg4 *LnsList2_[*WasmBin_Ins], arg5 LnsAny, arg6 LnsAny) {
    self.kind = arg1
    self.code = arg2
    self.args = arg3
    self.body = arg4
    self.elseBody = arg5
    self.result = arg6
}
func (self *WasmBin_Ins) Get_kind(_env *LnsEnv) LnsInt{ return self.kind }
func (self *WasmBin_Ins) Get_code(_env *LnsEnv) string{ return self.code }
func (self *WasmBin_Ins) Get_args(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]{ return self.args }
func (self *WasmBin_Ins) Get_body(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]{ return self.body }
func (self *WasmBin_Ins) Get_elseBody(_env *LnsEnv) LnsAny{ return self.elseBody }
func (self *WasmBin_Ins) Get_result(_env *LnsEnv) LnsAny{ return self.result }

// declaration Class -- Func
type WasmBin_FuncMtd interface {
    AddLocal(_env *LnsEnv, arg1 *WasmBin_ValType) LnsInt
    Encode(_env *LnsEnv) string
    Get_body(_env *LnsEnv) *LnsList2_[*WasmBin_Ins]
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
// 352: DeclConstr
func (self *WasmBin_Func) InitWasmBin_Func(_env *LnsEnv, typeIdx LnsInt,paramNum LnsInt) {
    self.typeIdx = typeIdx
    self.paramNum = paramNum
    self.localTypes = NewLnsList2_[*WasmBin_ValType]([]*WasmBin_ValType{})
    self.body = NewLnsList2_[*WasmBin_Ins]([]*WasmBin_Ins{})
}


// declaration Class -- Module
type WasmBin_ModuleMtd interface {
    AddByteArrayType(_env *LnsEnv) LnsInt
    AddData(_env *LnsEnv, arg1 string) LnsInt
    AddFunc(_env *LnsEnv, arg1 *LnsList2_[*WasmBin_ValType], arg2 *LnsList2_[*WasmBin_ValType])(LnsInt, *WasmBin_Func)
    AddGlobal(_env *LnsEnv, arg1 *WasmBin_ValType, arg2 *WasmBin_Ins) LnsInt
    AddType(_env *LnsEnv, arg1 string) LnsInt
    Encode(_env *LnsEnv) string
    ExportFunc(_env *LnsEnv, arg1 string, arg2 LnsInt)
    GetFuncType(_env *LnsEnv, arg1 *LnsList2_[*WasmBin_ValType], arg2 *LnsList2_[*WasmBin_ValType]) LnsInt
    Get_importFuncNum(_env *LnsEnv) LnsInt
    ImportFunc(_env *LnsEnv, arg1 string, arg2 string, arg3 *LnsList2_[*WasmBin_ValType], arg4 *LnsList2_[*WasmBin_ValType]) LnsInt
}
type WasmBin_Module struct {
    typeList *LnsList2_[string]
    funcTypeMap *LnsMap2_[string,LnsInt]
    importList *LnsList2_[string]
    importFuncNum LnsInt
    funcList *LnsList2_[*WasmBin_Func]
    globalList *LnsList2_[string]
    exportList *LnsList2_[string]
    dataList *LnsList2_[string]
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
func (self *WasmBin_Module) Get_importFuncNum(_env *LnsEnv) LnsInt{ return self.importFuncNum }
// 396: DeclConstr
func (self *WasmBin_Module) InitWasmBin_Module(_env *LnsEnv) {
    self.typeList = NewLnsList2_[string]([]string{})
    self.funcTypeMap = NewLnsMap2_[string,LnsInt]( map[string]LnsInt{})
    self.importList = NewLnsList2_[string]([]string{})
    self.importFuncNum = 0
    self.funcList = NewLnsList2_[*WasmBin_Func]([]*WasmBin_Func{})
    self.globalList = NewLnsList2_[string]([]string{})
    self.exportList = NewLnsList2_[string]([]string{})
    self.dataList = NewLnsList2_[string]([]string{})
}


func Lns_WasmBin_init(_env *LnsEnv) {
    if init_WasmBin { return }
    init_WasmBin = true
    WasmBin__mod__ = "@lune.@base.@WasmBin"
    Lns_InitMod()
    WasmBin_typeI32 = NewWasmBin_ValType(_env, WasmBin_ValKind__I32, 0, false)
    WasmBin_typeI64 = NewWasmBin_ValType(_env, WasmBin_ValKind__I64, 0, false)
    WasmBin_typeF64 = NewWasmBin_ValType(_env, WasmBin_ValKind__F64, 0, false)
}
func init() {
    init_WasmBin = false
}
