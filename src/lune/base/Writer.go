// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_Writer bool
var Writer__mod__ string
// for 56
func Writer_convExp147(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 57
func Writer_convExp165(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 58
func Writer_convExp183(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 59
func Writer_convExp201(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 60
func Writer_convExp219(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 263
func Writer_convExp1211(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 264
func Writer_convExp1229(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 265
func Writer_convExp1247(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
type Writer_Writer interface {
        EndElement()
        Fin()
        StartElement(arg1 string)
        StartParent(arg1 string, arg2 bool)
        Write(arg1 string, arg2 LnsAny)
        WriteValue(arg1 LnsAny)
}
func Lns_cast2Writer_Writer( obj LnsAny ) LnsAny {
    if _, ok := obj.(Writer_Writer); ok { 
        return obj
    }
    return nil
}

// declaration Class -- XML
type Writer_XMLMtd interface {
    EndElement()
    Fin()
    StartElement(arg1 string)
    StartParent(arg1 string, arg2 bool)
    Write(arg1 string, arg2 LnsAny)
    WriteValue(arg1 LnsAny)
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
func NewWriter_XML(arg1 Lns_oStream) *Writer_XML {
    obj := &Writer_XML{}
    obj.FP = obj
    obj.InitWriter_XML(arg1)
    return obj
}
// 41: DeclConstr
func (self *Writer_XML) InitWriter_XML(stream Lns_oStream) {
    self.stream = stream
    
    self.elementList = NewLnsList([]LnsAny{})
    
    self.depth = 0
    
}

// 47: decl @lune.@base.@Writer.XML.convertXmlTxt
func Writer_XML_convertXmlTxt_1052_(val LnsAny) string {
    if val == ""{
        return ""
    }
    if Lns_type(val) == "number"{
        return Lns_getVM().String_format("%g", []LnsAny{Lns_forceCastReal(val)})
    }
    var txt string
    txt = Lns_getVM().String_format("%s", []LnsAny{val})
    txt = Writer_convExp147(Lns_2DDD(Lns_getVM().String_gsub(txt, "&", "&amp;")))
    
    txt = Writer_convExp165(Lns_2DDD(Lns_getVM().String_gsub(txt, ">", "&gt;")))
    
    txt = Writer_convExp183(Lns_2DDD(Lns_getVM().String_gsub(txt, "<", "&lt;")))
    
    txt = Writer_convExp201(Lns_2DDD(Lns_getVM().String_gsub(txt, "\"", "&quot;")))
    
    txt = Writer_convExp219(Lns_2DDD(Lns_getVM().String_gsub(txt, "'", "&apos;")))
    
    return txt
}

// 64: decl @lune.@base.@Writer.XML.startElement
func (self *Writer_XML) StartElement(name string) {
    self.elementList.Insert(name)
    self.stream.Write(Lns_getVM().String_format("<%s>", []LnsAny{name}))
    self.depth = self.depth + 1
    
}

// 70: decl @lune.@base.@Writer.XML.startParent
func (self *Writer_XML) StartParent(name string,arrayFlag bool) {
    self.FP.StartElement(name)
}

// 73: decl @lune.@base.@Writer.XML.endElement
func (self *Writer_XML) EndElement() {
    var name LnsAny
    name = self.elementList.Remove(nil)
    self.stream.Write(Lns_getVM().String_format("</%s>", []LnsAny{name}))
    self.depth = self.depth - 1
    
    if self.depth == 0{
        self.stream.Write("\n")
    } else if self.depth < 0{
        Util_err("illegal depth")
    }
}

// 84: decl @lune.@base.@Writer.XML.writeValue
func (self *Writer_XML) WriteValue(val LnsAny) {
    self.stream.Write(Writer_XML_convertXmlTxt_1052_(val))
}

// 87: decl @lune.@base.@Writer.XML.write
func (self *Writer_XML) Write(name string,val LnsAny) {
    self.FP.StartElement(name)
    self.FP.WriteValue(val)
    self.FP.EndElement()
}

// 92: decl @lune.@base.@Writer.XML.fin
func (self *Writer_XML) Fin() {
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
func NewWriter_JsonLayer(arg1 string, arg2 bool, arg3 string, arg4 bool, arg5 *LnsSet, arg6 bool, arg7 bool) *Writer_JsonLayer {
    obj := &Writer_JsonLayer{}
    obj.FP = obj
    obj.InitWriter_JsonLayer(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    return obj
}
func (self *Writer_JsonLayer) InitWriter_JsonLayer(arg1 string, arg2 bool, arg3 string, arg4 bool, arg5 *LnsSet, arg6 bool, arg7 bool) {
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
    AddElementName(arg1 string)
    EndElement()
    EndLayer()
    EqualLayerState(arg1 string) bool
    Fin()
    getLayerInfo() LnsAny
    GetLayerName() string
    IsArrayLayer() bool
    SetLayerState(arg1 string)
    StartElement(arg1 string)
    startLayer(arg1 bool, arg2 bool)
    StartParent(arg1 string, arg2 bool)
    Write(arg1 string, arg2 LnsAny)
    WriteValue(arg1 LnsAny)
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
func NewWriter_JSON(arg1 Lns_oStream) *Writer_JSON {
    obj := &Writer_JSON{}
    obj.FP = obj
    obj.InitWriter_JSON(arg1)
    return obj
}
// 112: decl @lune.@base.@Writer.JSON.startLayer
func (self *Writer_JSON) startLayer(arrayFlag bool,madeByArrayFlag bool) {
    var info *Writer_JsonLayer
    info = NewWriter_JsonLayer("none", arrayFlag, self.prevName, madeByArrayFlag, NewLnsSet([]LnsAny{}), true, false)
    self.layerQueue.Insert(Writer_JsonLayer2Stem(info))
    self.stream.Write(Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( arrayFlag) &&
        Lns_setStackVal( "[") ||
        Lns_setStackVal( "{") ).(string))
}

// 126: DeclConstr
func (self *Writer_JSON) InitWriter_JSON(stream Lns_oStream) {
    self.stream = stream
    
    self.layerQueue = NewLnsList([]LnsAny{})
    
    self.prevName = ""
    
    self.FP.startLayer(false, false)
}

// 133: decl @lune.@base.@Writer.JSON.getLayerInfo
func (self *Writer_JSON) getLayerInfo() LnsAny {
    if self.layerQueue.Len() == 0{
        return nil
    }
    return self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer()
}

// 140: decl @lune.@base.@Writer.JSON.endLayer
func (self *Writer_JSON) EndLayer() {
    if self.layerQueue.Len() == 0{
        Util_err("illegal depth")
    }
    for self.layerQueue.Len() > 0 {
        var info *Writer_JsonLayer
        info = Lns_unwrap( self.FP.getLayerInfo()).(*Writer_JsonLayer)
        if info.ArrayFlag{
            self.stream.Write("]")
        } else { 
            self.stream.Write("}")
        }
        self.layerQueue.Remove(nil)
        var parentInfo LnsAny
        parentInfo = self.FP.getLayerInfo()
        if Lns_op_not(Lns_NilAccFin(Lns_NilAccPush(parentInfo) && 
        Lns_NilAccPush(Lns_NilAccPop().(*Writer_JsonLayer).MadeByArrayFlag))){
            break
        }
    }
    if self.layerQueue.Len() == 0{
        self.stream.Write("\n")
    }
}

// 164: decl @lune.@base.@Writer.JSON.equalLayerState
func (self *Writer_JSON) EqualLayerState(state string) bool {
    return self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer().State == state
}

// 168: decl @lune.@base.@Writer.JSON.isArrayLayer
func (self *Writer_JSON) IsArrayLayer() bool {
    return self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer().ArrayFlag
}

// 172: decl @lune.@base.@Writer.JSON.setLayerState
func (self *Writer_JSON) SetLayerState(state string) {
    self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer().State = state
    
}

// 176: decl @lune.@base.@Writer.JSON.getLayerName
func (self *Writer_JSON) GetLayerName() string {
    return self.layerQueue.GetAt(self.layerQueue.Len()).(Writer_JsonLayerDownCast).ToWriter_JsonLayer().Name
}

// 180: decl @lune.@base.@Writer.JSON.addElementName
func (self *Writer_JSON) AddElementName(name string) {
    var info *Writer_JsonLayer
    info = Lns_unwrap( self.FP.getLayerInfo()).(*Writer_JsonLayer)
    var nameSet *LnsSet
    nameSet = info.ElementNameSet
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( Lns_op_not(info.ArrayFlag)) &&
        Lns_setStackVal( nameSet.Has(name)) ).(bool)){
        Util_err("exist same name: " + name)
    }
    nameSet.Add(name)
}

// 189: decl @lune.@base.@Writer.JSON.startParent
func (self *Writer_JSON) StartParent(name string,arrayFlag bool) {
    self.FP.AddElementName(name)
    if Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( self.FP.EqualLayerState("termed")) ||
        Lns_setStackVal( self.FP.EqualLayerState("named")) ||
        Lns_setStackVal( self.FP.EqualLayerState("valued")) ).(bool){
        self.stream.Write(",")
    } else if self.FP.EqualLayerState("none"){
    }
    var parentInfo LnsAny
    parentInfo = self.FP.getLayerInfo()
    if Lns_isCondTrue( Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( Lns_op_not(arrayFlag)) &&
        Lns_setStackVal( Lns_NilAccFin(Lns_NilAccPush(parentInfo) && 
        Lns_NilAccPush(Lns_NilAccPop().(*Writer_JsonLayer).ArrayFlag))) )){
        self.FP.startLayer(false, true)
    }
    self.stream.Write(Lns_getVM().String_format("\"%s\": ", []LnsAny{name}))
    self.FP.startLayer(arrayFlag, false)
    self.prevName = name
    
}

// 212: decl @lune.@base.@Writer.JSON.startElement
func (self *Writer_JSON) StartElement(name string) {
    self.FP.AddElementName(name)
    if self.FP.EqualLayerState("termed"){
        self.stream.Write(",")
    } else if self.FP.EqualLayerState("named"){
        Util_err("illegal layer state")
    } else if self.FP.EqualLayerState("none"){
    }
    if self.FP.IsArrayLayer(){
        self.FP.startLayer(false, true)
    }
    var info *Writer_JsonLayer
    info = Lns_unwrap( self.FP.getLayerInfo()).(*Writer_JsonLayer)
    if info.OpenElement{
        Util_err("illegal openElement")
    }
    info.OpenElement = true
    
    self.stream.Write(Lns_getVM().String_format("\"%s\": ", []LnsAny{name}))
    self.FP.SetLayerState("named")
    self.prevName = name
    
}

// 240: decl @lune.@base.@Writer.JSON.endElement
func (self *Writer_JSON) EndElement() {
    if Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( self.FP.EqualLayerState("none")) ||
        Lns_setStackVal( self.FP.EqualLayerState("termed")) ).(bool){
        self.FP.EndLayer()
    } else if self.FP.EqualLayerState("valued"){
        var info *Writer_JsonLayer
        info = Lns_unwrap( self.FP.getLayerInfo()).(*Writer_JsonLayer)
        if info.OpenElement{
            info.OpenElement = false
            
        }
        if Lns_isCondTrue( Lns_NilAccFin(Lns_NilAccPush(self.FP.getLayerInfo()) && 
        Lns_NilAccPush(Lns_NilAccPop().(*Writer_JsonLayer).MadeByArrayFlag))){
            self.FP.EndLayer()
        }
    } else { 
        Util_err(Lns_getVM().String_format("illegal layer state %s", []LnsAny{self.FP.GetLayerName()}))
    }
    self.FP.SetLayerState("termed")
}

// 259: decl @lune.@base.@Writer.JSON.convertJsonTxt
func Writer_JSON_convertJsonTxt_1142_(txt string) string {
    if txt == ""{
        return ""
    }
    txt = Writer_convExp1211(Lns_2DDD(Lns_getVM().String_gsub(txt, "\"", "\\\"")))
    
    txt = Writer_convExp1229(Lns_2DDD(Lns_getVM().String_gsub(txt, "\\", "\\\\")))
    
    txt = Writer_convExp1247(Lns_2DDD(Lns_getVM().String_gsub(txt, "\n", "\\n")))
    
    return txt
}

// 270: decl @lune.@base.@Writer.JSON.writeValue
func (self *Writer_JSON) WriteValue(val LnsAny) {
    var txt string
    txt = ""
    var typeId string
    typeId = Lns_type(val)
    if typeId == "number"{
        txt = Lns_getVM().String_format("%g", []LnsAny{Lns_forceCastReal(val)})
        
    } else if typeId == "boolean"{
        txt = Lns_popVal( Lns_incStack() ||
            Lns_setStackVal( val) &&
            Lns_setStackVal( "true") ||
            Lns_setStackVal( "false") ).(string)
        
    } else { 
        txt = Lns_getVM().String_format("\"%s\"", []LnsAny{Writer_JSON_convertJsonTxt_1142_(Lns_getVM().String_format("%s", []LnsAny{val}))})
        
    }
    self.stream.Write(txt)
    self.FP.SetLayerState("valued")
}

// 287: decl @lune.@base.@Writer.JSON.write
func (self *Writer_JSON) Write(name string,val LnsAny) {
    self.FP.StartElement(name)
    self.FP.WriteValue(val)
    self.FP.EndElement()
}

// 293: decl @lune.@base.@Writer.JSON.fin
func (self *Writer_JSON) Fin() {
    if Lns_popVal( Lns_incStack() ||
        Lns_setStackVal( self.FP.EqualLayerState("none")) ||
        Lns_setStackVal( self.FP.EqualLayerState("termed")) ).(bool){
        self.FP.EndLayer()
    } else { 
        Util_err("illegal")
    }
}


func Lns_Writer_init() {
    if init_Writer { return }
    init_Writer = true
    Writer__mod__ = "@lune.@base.@Writer"
    Lns_InitMod()
    Lns_Util_init()
}
func init() {
    init_Writer = false
}
