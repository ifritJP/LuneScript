// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Writer bool
var Writer__mod__ string
// for 58
func Writer_convExp0_346(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 59
func Writer_convExp0_364(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 60
func Writer_convExp0_382(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 61
func Writer_convExp0_400(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 62
func Writer_convExp0_418(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 265
func Writer_convExp0_1237(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 266
func Writer_convExp0_1255(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 267
func Writer_convExp0_1273(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 49: decl @lune.@base.@Writer.XML.convertXmlTxt
func Writer_XML_convertXmlTxt_1_(_env *LnsEnv, val LnsAny) string {
    if val == ""{
        return ""
    }
    if Lns_type(val) == "number"{
        return _env.GetVM().String_format("%g", []LnsAny{Lns_forceCastReal(val)})
    }
    var txt string
    txt = _env.GetVM().String_format("%s", []LnsAny{val})
    txt = Writer_convExp0_346(Lns_2DDD(_env.GetVM().String_gsub(txt, "&", "&amp;")))
    txt = Writer_convExp0_364(Lns_2DDD(_env.GetVM().String_gsub(txt, ">", "&gt;")))
    txt = Writer_convExp0_382(Lns_2DDD(_env.GetVM().String_gsub(txt, "<", "&lt;")))
    txt = Writer_convExp0_400(Lns_2DDD(_env.GetVM().String_gsub(txt, "\"", "&quot;")))
    txt = Writer_convExp0_418(Lns_2DDD(_env.GetVM().String_gsub(txt, "'", "&apos;")))
    return txt
}
// 66: decl @lune.@base.@Writer.XML.startElement
func (self *Writer_XML) StartElement(_env *LnsEnv, name string) {
    self.elementList.Insert(name)
    self.stream.Write(_env, _env.GetVM().String_format("<%s>", []LnsAny{name}))
    self.depth = self.depth + 1
}
// 72: decl @lune.@base.@Writer.XML.startParent
func (self *Writer_XML) StartParent(_env *LnsEnv, name string,arrayFlag bool) {
    self.FP.StartElement(_env, name)
}
// 75: decl @lune.@base.@Writer.XML.endElement
func (self *Writer_XML) EndElement(_env *LnsEnv) {
    var name LnsAny
    name = self.elementList.Remove(nil)
    self.stream.Write(_env, _env.GetVM().String_format("</%s>", []LnsAny{name}))
    self.depth = self.depth - 1
    if self.depth == 0{
        self.stream.Write(_env, "\n")
    } else if self.depth < 0{
        Util_err(_env, "illegal depth")
    }
}
// 86: decl @lune.@base.@Writer.XML.writeValue
func (self *Writer_XML) WriteValue(_env *LnsEnv, val LnsAny) {
    self.stream.Write(_env, Writer_XML_convertXmlTxt_1_(_env, val))
}
// 89: decl @lune.@base.@Writer.XML.write
func (self *Writer_XML) Write(_env *LnsEnv, name string,val LnsAny) {
    self.FP.StartElement(_env, name)
    self.FP.WriteValue(_env, val)
    self.FP.EndElement(_env)
}
// 94: decl @lune.@base.@Writer.XML.fin
func (self *Writer_XML) Fin(_env *LnsEnv) {
}
// 114: decl @lune.@base.@Writer.JSON.startLayer
func (self *Writer_JSON) startLayer(_env *LnsEnv, arrayFlag bool,madeByArrayFlag bool) {
    var info *Writer_JsonLayer
    info = NewWriter_JsonLayer(_env, "none", arrayFlag, self.prevName, madeByArrayFlag, NewLnsSet([]LnsAny{}), true, false)
    self.layerQueue.Insert(Writer_JsonLayer2Stem(info))
    self.stream.Write(_env, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( arrayFlag) &&
        _env.SetStackVal( "[") ||
        _env.SetStackVal( "{") ).(string))
}
// 135: decl @lune.@base.@Writer.JSON.getLayerInfo
func (self *Writer_JSON) getLayerInfo(_env *LnsEnv) LnsAny {
    if self.layerQueue.Len() == 0{
        return nil
    }
    return self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer()
}
// 142: decl @lune.@base.@Writer.JSON.endLayer
func (self *Writer_JSON) EndLayer(_env *LnsEnv) {
    if self.layerQueue.Len() == 0{
        Util_err(_env, "illegal depth")
    }
    for self.layerQueue.Len() > 0 {
        var info *Writer_JsonLayer
        info = Lns_unwrap( self.FP.getLayerInfo(_env)).(*Writer_JsonLayer)
        if info.ArrayFlag{
            self.stream.Write(_env, "]")
        } else { 
            self.stream.Write(_env, "}")
        }
        self.layerQueue.Remove(nil)
        var parentInfo LnsAny
        parentInfo = self.FP.getLayerInfo(_env)
        if Lns_op_not(_env.NilAccFin(_env.NilAccPush(parentInfo) && 
        _env.NilAccPush(_env.NilAccPop().(*Writer_JsonLayer).MadeByArrayFlag))){
            break
        }
    }
    if self.layerQueue.Len() == 0{
        self.stream.Write(_env, "\n")
    }
}
// 166: decl @lune.@base.@Writer.JSON.equalLayerState
func (self *Writer_JSON) EqualLayerState(_env *LnsEnv, state string) bool {
    return self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer().State == state
}
// 170: decl @lune.@base.@Writer.JSON.isArrayLayer
func (self *Writer_JSON) IsArrayLayer(_env *LnsEnv) bool {
    return self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer().ArrayFlag
}
// 174: decl @lune.@base.@Writer.JSON.setLayerState
func (self *Writer_JSON) SetLayerState(_env *LnsEnv, state string) {
    self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer().State = state
}
// 178: decl @lune.@base.@Writer.JSON.getLayerName
func (self *Writer_JSON) GetLayerName(_env *LnsEnv) string {
    return self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer().Name
}
// 182: decl @lune.@base.@Writer.JSON.addElementName
func (self *Writer_JSON) AddElementName(_env *LnsEnv, name string) {
    var info *Writer_JsonLayer
    info = Lns_unwrap( self.FP.getLayerInfo(_env)).(*Writer_JsonLayer)
    var nameSet *LnsSet
    nameSet = info.ElementNameSet
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(info.ArrayFlag)) &&
        _env.SetStackVal( nameSet.Has(name)) ).(bool)){
        Util_err(_env, "exist same name: " + name)
    }
    nameSet.Add(name)
}
// 191: decl @lune.@base.@Writer.JSON.startParent
func (self *Writer_JSON) StartParent(_env *LnsEnv, name string,arrayFlag bool) {
    self.FP.AddElementName(_env, name)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.EqualLayerState(_env, "termed")) ||
        _env.SetStackVal( self.FP.EqualLayerState(_env, "named")) ||
        _env.SetStackVal( self.FP.EqualLayerState(_env, "valued")) ).(bool){
        self.stream.Write(_env, ",")
    } else if self.FP.EqualLayerState(_env, "none"){
    }
    var parentInfo LnsAny
    parentInfo = self.FP.getLayerInfo(_env)
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(arrayFlag)) &&
        _env.SetStackVal( _env.NilAccFin(_env.NilAccPush(parentInfo) && 
        _env.NilAccPush(_env.NilAccPop().(*Writer_JsonLayer).ArrayFlag))) )){
        self.FP.startLayer(_env, false, true)
    }
    self.stream.Write(_env, _env.GetVM().String_format("\"%s\": ", []LnsAny{name}))
    self.FP.startLayer(_env, arrayFlag, false)
    self.prevName = name
}
// 214: decl @lune.@base.@Writer.JSON.startElement
func (self *Writer_JSON) StartElement(_env *LnsEnv, name string) {
    self.FP.AddElementName(_env, name)
    if self.FP.EqualLayerState(_env, "termed"){
        self.stream.Write(_env, ",")
    } else if self.FP.EqualLayerState(_env, "named"){
        Util_err(_env, "illegal layer state")
    } else if self.FP.EqualLayerState(_env, "none"){
    }
    if self.FP.IsArrayLayer(_env){
        self.FP.startLayer(_env, false, true)
    }
    var info *Writer_JsonLayer
    info = Lns_unwrap( self.FP.getLayerInfo(_env)).(*Writer_JsonLayer)
    if info.OpenElement{
        Util_err(_env, "illegal openElement")
    }
    info.OpenElement = true
    self.stream.Write(_env, _env.GetVM().String_format("\"%s\": ", []LnsAny{name}))
    self.FP.SetLayerState(_env, "named")
    self.prevName = name
}
// 242: decl @lune.@base.@Writer.JSON.endElement
func (self *Writer_JSON) EndElement(_env *LnsEnv) {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.EqualLayerState(_env, "none")) ||
        _env.SetStackVal( self.FP.EqualLayerState(_env, "termed")) ).(bool){
        self.FP.EndLayer(_env)
    } else if self.FP.EqualLayerState(_env, "valued"){
        var info *Writer_JsonLayer
        info = Lns_unwrap( self.FP.getLayerInfo(_env)).(*Writer_JsonLayer)
        if info.OpenElement{
            info.OpenElement = false
        }
        if Lns_isCondTrue( _env.NilAccFin(_env.NilAccPush(self.FP.getLayerInfo(_env)) && 
        _env.NilAccPush(_env.NilAccPop().(*Writer_JsonLayer).MadeByArrayFlag))){
            self.FP.EndLayer(_env)
        }
    } else { 
        Util_err(_env, _env.GetVM().String_format("illegal layer state %s", []LnsAny{self.FP.GetLayerName(_env)}))
    }
    self.FP.SetLayerState(_env, "termed")
}
// 261: decl @lune.@base.@Writer.JSON.convertJsonTxt
func Writer_JSON_convertJsonTxt_12_(_env *LnsEnv, txt string) string {
    if txt == ""{
        return ""
    }
    txt = Writer_convExp0_1237(Lns_2DDD(_env.GetVM().String_gsub(txt, "\"", "\\\"")))
    txt = Writer_convExp0_1255(Lns_2DDD(_env.GetVM().String_gsub(txt, "\\", "\\\\")))
    txt = Writer_convExp0_1273(Lns_2DDD(_env.GetVM().String_gsub(txt, "\n", "\\n")))
    return txt
}
// 272: decl @lune.@base.@Writer.JSON.writeValue
func (self *Writer_JSON) WriteValue(_env *LnsEnv, val LnsAny) {
    var txt string
    txt = ""
    var typeId string
    typeId = Lns_type(val)
    if typeId == "number"{
        txt = _env.GetVM().String_format("%g", []LnsAny{Lns_forceCastReal(val)})
    } else if typeId == "boolean"{
        txt = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( val) &&
            _env.SetStackVal( "true") ||
            _env.SetStackVal( "false") ).(string)
    } else { 
        txt = _env.GetVM().String_format("\"%s\"", []LnsAny{Writer_JSON_convertJsonTxt_12_(_env, _env.GetVM().String_format("%s", []LnsAny{val}))})
    }
    self.stream.Write(_env, txt)
    self.FP.SetLayerState(_env, "valued")
}
// 289: decl @lune.@base.@Writer.JSON.write
func (self *Writer_JSON) Write(_env *LnsEnv, name string,val LnsAny) {
    self.FP.StartElement(_env, name)
    self.FP.WriteValue(_env, val)
    self.FP.EndElement(_env)
}
// 295: decl @lune.@base.@Writer.JSON.fin
func (self *Writer_JSON) Fin(_env *LnsEnv) {
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( self.FP.EqualLayerState(_env, "none")) ||
        _env.SetStackVal( self.FP.EqualLayerState(_env, "termed")) ).(bool){
        self.FP.EndLayer(_env)
    } else { 
        Util_err(_env, "illegal")
    }
}
type Writer_Writer interface {
        EndElement(_env *LnsEnv)
        Fin(_env *LnsEnv)
        StartElement(_env *LnsEnv, arg1 string)
        StartParent(_env *LnsEnv, arg1 string, arg2 bool)
        Write(_env *LnsEnv, arg1 string, arg2 LnsAny)
        WriteValue(_env *LnsEnv, arg1 LnsAny)
}
func Lns_cast2Writer_Writer( obj LnsAny ) LnsAny {
    if _, ok := obj.(Writer_Writer); ok { 
        return obj
    }
    return nil
}

// declaration Class -- XML
type Writer_XMLMtd interface {
    EndElement(_env *LnsEnv)
    Fin(_env *LnsEnv)
    StartElement(_env *LnsEnv, arg1 string)
    StartParent(_env *LnsEnv, arg1 string, arg2 bool)
    Write(_env *LnsEnv, arg1 string, arg2 LnsAny)
    WriteValue(_env *LnsEnv, arg1 LnsAny)
}
type Writer_XML struct {
    stream Lns_oStream
    elementList *LnsList
    depth LnsInt
    FP Writer_XMLMtd
}
func Writer_XML2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Writer_XML).FP
}
type Writer_XMLDownCast interface {
    ToWriter_XML() *Writer_XML
}
func Writer_XMLDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Writer_XMLDownCast)
    if ok { return work.ToWriter_XML() }
    return nil
}
func (obj *Writer_XML) ToWriter_XML() *Writer_XML {
    return obj
}
func NewWriter_XML(_env *LnsEnv, arg1 Lns_oStream) *Writer_XML {
    obj := &Writer_XML{}
    obj.FP = obj
    obj.InitWriter_XML(_env, arg1)
    return obj
}
// 43: DeclConstr
func (self *Writer_XML) InitWriter_XML(_env *LnsEnv, stream Lns_oStream) {
    self.stream = stream
    self.elementList = NewLnsList([]LnsAny{})
    self.depth = 0
}


// declaration Class -- JsonLayer
type Writer_JsonLayerMtd interface {
}
type Writer_JsonLayer struct {
    State string
    ArrayFlag bool
    Name string
    MadeByArrayFlag bool
    ElementNameSet *LnsSet
    ParentFlag bool
    OpenElement bool
    FP Writer_JsonLayerMtd
}
func Writer_JsonLayer2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Writer_JsonLayer).FP
}
type Writer_JsonLayerDownCast interface {
    ToWriter_JsonLayer() *Writer_JsonLayer
}
func Writer_JsonLayerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Writer_JsonLayerDownCast)
    if ok { return work.ToWriter_JsonLayer() }
    return nil
}
func (obj *Writer_JsonLayer) ToWriter_JsonLayer() *Writer_JsonLayer {
    return obj
}
func NewWriter_JsonLayer(_env *LnsEnv, arg1 string, arg2 bool, arg3 string, arg4 bool, arg5 *LnsSet, arg6 bool, arg7 bool) *Writer_JsonLayer {
    obj := &Writer_JsonLayer{}
    obj.FP = obj
    obj.InitWriter_JsonLayer(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Writer_JsonLayer) InitWriter_JsonLayer(_env *LnsEnv, arg1 string, arg2 bool, arg3 string, arg4 bool, arg5 *LnsSet, arg6 bool, arg7 bool) {
    self.State = arg1
    self.ArrayFlag = arg2
    self.Name = arg3
    self.MadeByArrayFlag = arg4
    self.ElementNameSet = arg5
    self.ParentFlag = arg6
    self.OpenElement = arg7
}

// declaration Class -- JSON
type Writer_JSONMtd interface {
    AddElementName(_env *LnsEnv, arg1 string)
    EndElement(_env *LnsEnv)
    EndLayer(_env *LnsEnv)
    EqualLayerState(_env *LnsEnv, arg1 string) bool
    Fin(_env *LnsEnv)
    getLayerInfo(_env *LnsEnv) LnsAny
    GetLayerName(_env *LnsEnv) string
    IsArrayLayer(_env *LnsEnv) bool
    SetLayerState(_env *LnsEnv, arg1 string)
    StartElement(_env *LnsEnv, arg1 string)
    startLayer(_env *LnsEnv, arg1 bool, arg2 bool)
    StartParent(_env *LnsEnv, arg1 string, arg2 bool)
    Write(_env *LnsEnv, arg1 string, arg2 LnsAny)
    WriteValue(_env *LnsEnv, arg1 LnsAny)
}
type Writer_JSON struct {
    stream Lns_oStream
    layerQueue *LnsList
    prevName string
    FP Writer_JSONMtd
}
func Writer_JSON2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Writer_JSON).FP
}
type Writer_JSONDownCast interface {
    ToWriter_JSON() *Writer_JSON
}
func Writer_JSONDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Writer_JSONDownCast)
    if ok { return work.ToWriter_JSON() }
    return nil
}
func (obj *Writer_JSON) ToWriter_JSON() *Writer_JSON {
    return obj
}
func NewWriter_JSON(_env *LnsEnv, arg1 Lns_oStream) *Writer_JSON {
    obj := &Writer_JSON{}
    obj.FP = obj
    obj.InitWriter_JSON(_env, arg1)
    return obj
}
// 128: DeclConstr
func (self *Writer_JSON) InitWriter_JSON(_env *LnsEnv, stream Lns_oStream) {
    self.stream = stream
    self.layerQueue = NewLnsList([]LnsAny{})
    self.prevName = ""
    self.FP.startLayer(_env, false, false)
}


func Lns_Writer_init(_env *LnsEnv) {
    if init_Writer { return }
    init_Writer = true
    Writer__mod__ = "@lune.@base.@Writer"
    Lns_InitMod()
    Lns_Util_init(_env)
}
func init() {
    init_Writer = false
}
