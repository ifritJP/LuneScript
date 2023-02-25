// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Writer bool
var Writer__mod__ string
// for 59
func Writer_convExp0_343(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 60
func Writer_convExp0_361(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 61
func Writer_convExp0_379(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 62
func Writer_convExp0_397(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 63
func Writer_convExp0_415(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 266
func Writer_convExp0_1232(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 267
func Writer_convExp0_1250(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 268
func Writer_convExp0_1268(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 50: decl @lune.@base.@Writer.XML.convertXmlTxt
func Writer_XML_convertXmlTxt_1_(_env *LnsEnv, val LnsAny) string {
    if val == ""{
        return ""
    }
    if Lns_type(val) == "number"{
        return _env.GetVM().String_format("%g", Lns_2DDD(Lns_forceCastReal(val)))
    }
    var txt string
    txt = _env.GetVM().String_format("%s", Lns_2DDD(val))
    txt = Writer_convExp0_343(Lns_2DDD(_env.GetVM().String_gsub(txt, "&", "&amp;")))
    txt = Writer_convExp0_361(Lns_2DDD(_env.GetVM().String_gsub(txt, ">", "&gt;")))
    txt = Writer_convExp0_379(Lns_2DDD(_env.GetVM().String_gsub(txt, "<", "&lt;")))
    txt = Writer_convExp0_397(Lns_2DDD(_env.GetVM().String_gsub(txt, "\"", "&quot;")))
    txt = Writer_convExp0_415(Lns_2DDD(_env.GetVM().String_gsub(txt, "'", "&apos;")))
    return txt
}
// 67: decl @lune.@base.@Writer.XML.startElement
func (self *Writer_XML) StartElement(_env *LnsEnv, name string) {
    self.elementList.Insert(name)
    self.stream.Write(_env, _env.GetVM().String_format("<%s>", Lns_2DDD(name)))
    self.depth = self.depth + 1
}
// 73: decl @lune.@base.@Writer.XML.startParent
func (self *Writer_XML) StartParent(_env *LnsEnv, name string,arrayFlag bool) {
    self.FP.StartElement(_env, name)
}
// 76: decl @lune.@base.@Writer.XML.endElement
func (self *Writer_XML) EndElement(_env *LnsEnv) {
    var name LnsAny
    name = self.elementList.Remove(nil)
    self.stream.Write(_env, _env.GetVM().String_format("</%s>", Lns_2DDD(name)))
    self.depth = self.depth - 1
    if self.depth == 0{
        self.stream.Write(_env, "\n")
    } else if self.depth < 0{
        Util_err(_env, "illegal depth")
    }
}
// 87: decl @lune.@base.@Writer.XML.writeValue
func (self *Writer_XML) WriteValue(_env *LnsEnv, val LnsAny) {
    self.stream.Write(_env, Writer_XML_convertXmlTxt_1_(_env, val))
}
// 90: decl @lune.@base.@Writer.XML.write
func (self *Writer_XML) Write(_env *LnsEnv, name string,val LnsAny) {
    self.FP.StartElement(_env, name)
    self.FP.WriteValue(_env, val)
    self.FP.EndElement(_env)
}
// 95: decl @lune.@base.@Writer.XML.fin
func (self *Writer_XML) Fin(_env *LnsEnv) {
}
// 115: decl @lune.@base.@Writer.JSON.startLayer
func (self *Writer_JSON) startLayer(_env *LnsEnv, arrayFlag bool,madeByArrayFlag bool) {
    var info *Writer_JsonLayer
    info = NewWriter_JsonLayer(_env, "none", arrayFlag, self.prevName, madeByArrayFlag, NewLnsSet2_[string]([]string{}), true, false)
    self.layerQueue.Insert(info)
    self.stream.Write(_env, _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( arrayFlag) &&
        _env.SetStackVal( "[") ||
        _env.SetStackVal( "{") ).(string))
}
// 136: decl @lune.@base.@Writer.JSON.getLayerInfo
func (self *Writer_JSON) getLayerInfo(_env *LnsEnv) LnsAny {
    if self.layerQueue.Len() == 0{
        return nil
    }
    return self.layerQueue.GetAt(self.layerQueue.Len())
}
// 143: decl @lune.@base.@Writer.JSON.endLayer
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
// 167: decl @lune.@base.@Writer.JSON.equalLayerState
func (self *Writer_JSON) EqualLayerState(_env *LnsEnv, state string) bool {
    return self.layerQueue.GetAt(self.layerQueue.Len()).State == state
}
// 171: decl @lune.@base.@Writer.JSON.isArrayLayer
func (self *Writer_JSON) IsArrayLayer(_env *LnsEnv) bool {
    return self.layerQueue.GetAt(self.layerQueue.Len()).ArrayFlag
}
// 175: decl @lune.@base.@Writer.JSON.setLayerState
func (self *Writer_JSON) SetLayerState(_env *LnsEnv, state string) {
    self.layerQueue.GetAt(self.layerQueue.Len()).State = state
}
// 179: decl @lune.@base.@Writer.JSON.getLayerName
func (self *Writer_JSON) GetLayerName(_env *LnsEnv) string {
    return self.layerQueue.GetAt(self.layerQueue.Len()).Name
}
// 183: decl @lune.@base.@Writer.JSON.addElementName
func (self *Writer_JSON) AddElementName(_env *LnsEnv, name string) {
    var info *Writer_JsonLayer
    info = Lns_unwrap( self.FP.getLayerInfo(_env)).(*Writer_JsonLayer)
    var nameSet *LnsSet2_[string]
    nameSet = info.ElementNameSet
    if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(info.ArrayFlag)) &&
        _env.SetStackVal( nameSet.Has(name)) ).(bool)){
        Util_err(_env, "exist same name: " + name)
    }
    nameSet.Add(name)
}
// 192: decl @lune.@base.@Writer.JSON.startParent
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
    self.stream.Write(_env, _env.GetVM().String_format("\"%s\": ", Lns_2DDD(name)))
    self.FP.startLayer(_env, arrayFlag, false)
    self.prevName = name
}
// 215: decl @lune.@base.@Writer.JSON.startElement
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
    self.stream.Write(_env, _env.GetVM().String_format("\"%s\": ", Lns_2DDD(name)))
    self.FP.SetLayerState(_env, "named")
    self.prevName = name
}
// 243: decl @lune.@base.@Writer.JSON.endElement
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
        Util_err(_env, _env.GetVM().String_format("illegal layer state %s", Lns_2DDD(self.FP.GetLayerName(_env))))
    }
    self.FP.SetLayerState(_env, "termed")
}
// 262: decl @lune.@base.@Writer.JSON.convertJsonTxt
func Writer_JSON_convertJsonTxt_12_(_env *LnsEnv, txt string) string {
    if txt == ""{
        return ""
    }
    txt = Writer_convExp0_1232(Lns_2DDD(_env.GetVM().String_gsub(txt, "\"", "\\\"")))
    txt = Writer_convExp0_1250(Lns_2DDD(_env.GetVM().String_gsub(txt, "\\", "\\\\")))
    txt = Writer_convExp0_1268(Lns_2DDD(_env.GetVM().String_gsub(txt, "\n", "\\n")))
    return txt
}
// 273: decl @lune.@base.@Writer.JSON.writeValue
func (self *Writer_JSON) WriteValue(_env *LnsEnv, val LnsAny) {
    var txt string
    txt = ""
    var typeId string
    typeId = Lns_type(val)
    if typeId == "number"{
        txt = _env.GetVM().String_format("%g", Lns_2DDD(Lns_forceCastReal(val)))
    } else if typeId == "boolean"{
        txt = _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( val) &&
            _env.SetStackVal( "true") ||
            _env.SetStackVal( "false") ).(string)
    } else { 
        txt = _env.GetVM().String_format("\"%s\"", Lns_2DDD(Writer_JSON_convertJsonTxt_12_(_env, _env.GetVM().String_format("%s", Lns_2DDD(val)))))
    }
    self.stream.Write(_env, txt)
    self.FP.SetLayerState(_env, "valued")
}
// 290: decl @lune.@base.@Writer.JSON.write
func (self *Writer_JSON) Write(_env *LnsEnv, name string,val LnsAny) {
    self.FP.StartElement(_env, name)
    self.FP.WriteValue(_env, val)
    self.FP.EndElement(_env)
}
// 296: decl @lune.@base.@Writer.JSON.fin
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
    elementList *LnsList2_[string]
    depth LnsInt
    FP Writer_XMLMtd
}
func Writer_XML2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Writer_XML).FP
}
func Writer_XML_toSlice(slice []LnsAny) []*Writer_XML {
    ret := make([]*Writer_XML, len(slice))
    for index, val := range slice {
        ret[index] = val.(Writer_XMLDownCast).ToWriter_XML()
    }
    return ret
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
// 44: DeclConstr
func (self *Writer_XML) InitWriter_XML(_env *LnsEnv, stream Lns_oStream) {
    self.stream = stream
    self.elementList = NewLnsList2_[string]([]string{})
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
    ElementNameSet *LnsSet2_[string]
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
func Writer_JsonLayer_toSlice(slice []LnsAny) []*Writer_JsonLayer {
    ret := make([]*Writer_JsonLayer, len(slice))
    for index, val := range slice {
        ret[index] = val.(Writer_JsonLayerDownCast).ToWriter_JsonLayer()
    }
    return ret
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
func NewWriter_JsonLayer(_env *LnsEnv, arg1 string, arg2 bool, arg3 string, arg4 bool, arg5 *LnsSet2_[string], arg6 bool, arg7 bool) *Writer_JsonLayer {
    obj := &Writer_JsonLayer{}
    obj.FP = obj
    obj.InitWriter_JsonLayer(_env, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Writer_JsonLayer) InitWriter_JsonLayer(_env *LnsEnv, arg1 string, arg2 bool, arg3 string, arg4 bool, arg5 *LnsSet2_[string], arg6 bool, arg7 bool) {
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
    layerQueue *LnsList2_[*Writer_JsonLayer]
    prevName string
    FP Writer_JSONMtd
}
func Writer_JSON2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Writer_JSON).FP
}
func Writer_JSON_toSlice(slice []LnsAny) []*Writer_JSON {
    ret := make([]*Writer_JSON, len(slice))
    for index, val := range slice {
        ret[index] = val.(Writer_JSONDownCast).ToWriter_JSON()
    }
    return ret
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
// 129: DeclConstr
func (self *Writer_JSON) InitWriter_JSON(_env *LnsEnv, stream Lns_oStream) {
    self.stream = stream
    self.layerQueue = NewLnsList2_[*Writer_JsonLayer]([]*Writer_JsonLayer{})
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
