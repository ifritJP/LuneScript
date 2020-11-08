// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_Util bool
var Util__mod__ string
var Util_debugFlag bool
var Util_errorCode LnsInt
// for 256
func Util_convExp949(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 30: decl @lune.@base.@Util.setDebugFlag
func Util_setDebugFlag(flag bool) {
    Util_debugFlag = flag
    
}

// 34: decl @lune.@base.@Util.setErrorCode
func Util_setErrorCode(code LnsInt) {
    Util_errorCode = code
    
}

// 38: decl @lune.@base.@Util.errorLog
func Util_errorLog(message string) {
    Lns_io_stderr.Write(message + "\n")
}

// 42: decl @lune.@base.@Util.err
func Util_err(message string) {
    if Util_debugFlag{
        panic(message)
    }
    Util_errorLog(message)
    Lns_getVM().OS_exit(Util_errorCode)
}

// 53: decl @lune.@base.@Util.splitStr
func Util_splitStr(txt string,pattern string) *LnsList {
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    {
        _form138, _param138, _prev138 := Lns_getVM().String_gmatch(txt, pattern)
        for {
            _work138 := _form138.(*Lns_luaValue).Call( Lns_2DDD( _param138, _prev138 ) )
            _prev138 = Lns_getFromMulti(_work138,0)
            if Lns_IsNil( _prev138 ) { break }
            token := _prev138.(string)
            list.Insert(token)
        }
    }
    return list
}

// 221: decl @lune.@base.@Util.log
func Util_log(message string) {
    if Util_debugFlag{
        Util_errorLog(message)
    }
}

// 227: decl @lune.@base.@Util.printStackTrace
func Util_printStackTrace() {
    Util_errorLog(Depend_getStackTrace())
}


// 236: decl @lune.@base.@Util.getReadyCode
func Util_getReadyCode(depPath string,tgtPath string) bool {
    __func__ := "@lune.@base.@Util.getReadyCode"
    var tgtTime LnsReal
    var depTime LnsReal
    
    {
        _tgtTime, _depTime := Depend_getFileLastModifiedTime(tgtPath), Depend_getFileLastModifiedTime(depPath)
        if _tgtTime == nil || _depTime == nil{
            return false
        } else {
            tgtTime = _tgtTime.(LnsReal)
            depTime = _depTime.(LnsReal)
        }
    }
    if tgtTime >= depTime{
        return true
    }
    Log_log(Log_Level__Warn, __func__, 247, Log_CreateMessage(func() string {
        return Lns_getVM().String_format("not ready %g < %g : %s, %s", []LnsAny{tgtTime, depTime, tgtPath, depPath})
    }))
    
    return false
}

// 252: decl @lune.@base.@Util.scriptPath2Module
func Util_scriptPath2Module(path string) string {
    if Lns_isCondTrue( Lns_car(Lns_getVM().String_find(path,"^/", nil, nil))){
        Util_err("script must be relative-path -- " + path)
    }
    var mod string
    mod = Util_convExp949(Lns_2DDD(Lns_getVM().String_gsub(path, "/", ".")))
    return Lns_car(Lns_getVM().String_gsub(mod, "%.lns$", "")).(string)
}

// declaration Class -- OrderedSet
type Util_OrderedSetMtd interface {
    Add(arg1 LnsAny) bool
    Clone() *Util_OrderedSet
    Get_list() *LnsList
    Has(arg1 LnsAny) bool
    RemoveLast()
}
type Util_OrderedSet struct {
    set *LnsSet
    list *LnsList
    FP Util_OrderedSetMtd
}
func Util_OrderedSet2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_OrderedSet).FP
}
type Util_OrderedSetDownCast interface {
    ToUtil_OrderedSet() *Util_OrderedSet
}
func Util_OrderedSetDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_OrderedSetDownCast)
    if ok { return work.ToUtil_OrderedSet() }
    return nil
}
func (obj *Util_OrderedSet) ToUtil_OrderedSet() *Util_OrderedSet {
    return obj
}
func NewUtil_OrderedSet() *Util_OrderedSet {
    obj := &Util_OrderedSet{}
    obj.FP = obj
    obj.InitUtil_OrderedSet()
    return obj
}
func (self *Util_OrderedSet) Get_list() *LnsList{ return self.list }
// 65: DeclConstr
func (self *Util_OrderedSet) InitUtil_OrderedSet() {
    self.set = NewLnsSet([]LnsAny{})
    
    self.list = NewLnsList([]LnsAny{})
    
}

// 70: decl @lune.@base.@Util.OrderedSet.add
func (self *Util_OrderedSet) Add(_val LnsAny) bool {
    val := _val
    if Lns_op_not(self.set.Has(val)){
        self.set.Add(val)
        self.list.Insert(val)
        return true
    }
    return false
}

// 79: decl @lune.@base.@Util.OrderedSet.clone
func (self *Util_OrderedSet) Clone() *Util_OrderedSet {
    var obj *Util_OrderedSet
    obj = NewUtil_OrderedSet()
    for _, _val := range( self.list.Items ) {
        val := _val
        obj.set.Add(val)
        obj.list.Insert(val)
    }
    return obj
}

// 88: decl @lune.@base.@Util.OrderedSet.has
func (self *Util_OrderedSet) Has(_val LnsAny) bool {
    val := _val
    return self.set.Has(val)
}

// 92: decl @lune.@base.@Util.OrderedSet.removeLast
func (self *Util_OrderedSet) RemoveLast() {
    if self.list.Len() == 0{
        Util_err("empty")
    }
    self.set.Del(self.list.GetAt(self.list.Len()))
    self.list.Remove(nil)
}


// declaration Class -- memStream
type Util_memStreamMtd interface {
    Close()
    Flush()
    Get_txt() string
    Write(arg1 string)(LnsAny, LnsAny)
}
type Util_memStream struct {
    txt *Str_Builder
    FP Util_memStreamMtd
}
func Util_memStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_memStream).FP
}
type Util_memStreamDownCast interface {
    ToUtil_memStream() *Util_memStream
}
func Util_memStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_memStreamDownCast)
    if ok { return work.ToUtil_memStream() }
    return nil
}
func (obj *Util_memStream) ToUtil_memStream() *Util_memStream {
    return obj
}
func NewUtil_memStream() *Util_memStream {
    obj := &Util_memStream{}
    obj.FP = obj
    obj.InitUtil_memStream()
    return obj
}
// 103: DeclConstr
func (self *Util_memStream) InitUtil_memStream() {
    self.txt = NewStr_Builder()
    
}

// 106: decl @lune.@base.@Util.memStream.get_txt
func (self *Util_memStream) Get_txt() string {
    return self.txt.FP.Get_txt()
}

// 109: decl @lune.@base.@Util.memStream.write
func (self *Util_memStream) Write(val string)(LnsAny, LnsAny) {
    self.txt.FP.Add(val)
    return self.FP, nil
}

// 113: decl @lune.@base.@Util.memStream.close
func (self *Util_memStream) Close() {
}

// 115: decl @lune.@base.@Util.memStream.flush
func (self *Util_memStream) Flush() {
}


type Util_SourceStream interface {
        PopIndent()
        PushIndent(arg1 LnsAny)
        ReturnToSource()
        SwitchToHeader()
        Write(arg1 string)
        Writeln(arg1 string)
}
func Lns_cast2Util_SourceStream( obj LnsAny ) LnsAny {
    if _, ok := obj.(Util_SourceStream); ok { 
        return obj
    }
    return nil
}

// declaration Class -- SimpleSourceOStream
type Util_SimpleSourceOStreamMtd interface {
    get_indent() LnsInt
    PopIndent()
    PushIndent(arg1 LnsAny)
    ReturnToSource()
    SwitchToHeader()
    Write(arg1 string)
    Writeln(arg1 string)
}
type Util_SimpleSourceOStream struct {
    nowStream Lns_oStream
    srcStream Lns_oStream
    headStream Lns_oStream
    needIndent bool
    stepIndent LnsInt
    curLineNo LnsInt
    indentQueue *LnsList
    FP Util_SimpleSourceOStreamMtd
}
func Util_SimpleSourceOStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_SimpleSourceOStream).FP
}
type Util_SimpleSourceOStreamDownCast interface {
    ToUtil_SimpleSourceOStream() *Util_SimpleSourceOStream
}
func Util_SimpleSourceOStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_SimpleSourceOStreamDownCast)
    if ok { return work.ToUtil_SimpleSourceOStream() }
    return nil
}
func (obj *Util_SimpleSourceOStream) ToUtil_SimpleSourceOStream() *Util_SimpleSourceOStream {
    return obj
}
func NewUtil_SimpleSourceOStream(arg1 Lns_oStream, arg2 LnsAny, arg3 LnsInt) *Util_SimpleSourceOStream {
    obj := &Util_SimpleSourceOStream{}
    obj.FP = obj
    obj.InitUtil_SimpleSourceOStream(arg1, arg2, arg3)
    return obj
}
// 137: DeclConstr
func (self *Util_SimpleSourceOStream) InitUtil_SimpleSourceOStream(stream Lns_oStream,headStream LnsAny,stepIndent LnsInt) {
    self.srcStream = stream
    
    self.nowStream = stream
    
    self.headStream = Lns_unwrapDefault( headStream, stream).(Lns_oStream)
    
    self.needIndent = true
    
    self.curLineNo = 0
    
    self.stepIndent = stepIndent
    
    self.indentQueue = NewLnsList([]LnsAny{0})
    
}

// 147: decl @lune.@base.@Util.SimpleSourceOStream.get_indent
func (self *Util_SimpleSourceOStream) get_indent() LnsInt {
    if self.indentQueue.Len() > 0{
        return self.indentQueue.GetAt(self.indentQueue.Len()).(LnsInt)
    }
    return 0
}

// 168: decl @lune.@base.@Util.SimpleSourceOStream.write
func (self *Util_SimpleSourceOStream) Write(txt string) {
    var stream Lns_oStream
    stream = self.nowStream
    for _, _line := range( Str_getLineList(txt).Items ) {
        line := _line.(string)
        if self.needIndent{
            stream.Write(Lns_getVM().String_rep(" ", self.FP.get_indent()))
            self.needIndent = false
            
        }
        if Lns_isCondTrue( Lns_GetEnv().PopVal( Lns_GetEnv().IncStack() ||
            Lns_GetEnv().SetStackVal( len(line) > 0) &&
            Lns_GetEnv().SetStackVal( LnsInt(line[len(line)-1]) == 10) ).(bool)){
            self.curLineNo = self.curLineNo + 1
            
        }
        stream.Write(line)
    }
}

// 195: decl @lune.@base.@Util.SimpleSourceOStream.writeln
func (self *Util_SimpleSourceOStream) Writeln(txt string) {
    self.FP.Write(txt)
    self.FP.Write("\n")
    self.needIndent = true
    
}

// 201: decl @lune.@base.@Util.SimpleSourceOStream.pushIndent
func (self *Util_SimpleSourceOStream) PushIndent(newIndent LnsAny) {
    var indent LnsInt
    indent = Lns_unwrapDefault( newIndent, self.FP.get_indent() + self.stepIndent).(LnsInt)
    self.indentQueue.Insert(indent)
}

// 206: decl @lune.@base.@Util.SimpleSourceOStream.popIndent
func (self *Util_SimpleSourceOStream) PopIndent() {
    if self.indentQueue.Len() == 0{
        Util_err("self.indentQueue == 0")
    }
    self.indentQueue.Remove(nil)
}

// 213: decl @lune.@base.@Util.SimpleSourceOStream.switchToHeader
func (self *Util_SimpleSourceOStream) SwitchToHeader() {
    self.nowStream = self.headStream
    
}

// 216: decl @lune.@base.@Util.SimpleSourceOStream.returnToSource
func (self *Util_SimpleSourceOStream) ReturnToSource() {
    self.nowStream = self.srcStream
    
}


func Lns_Util_init() {
    if init_Util { return }
    init_Util = true
    Util__mod__ = "@lune.@base.@Util"
    Lns_InitMod()
    Lns_LuaVer_init()
    Lns_Depend_init()
    Lns_Log_init()
    Lns_Str_init()
    Util_debugFlag = true
    Util_errorCode = 1
}
func init() {
    init_Util = false
}
