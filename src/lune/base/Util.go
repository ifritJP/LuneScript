// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Util bool
var Util__mod__ string
var Util_debugFlag bool
var Util_errorCode LnsInt
// for 309
func Util_convExp1186(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 326
func Util_convExp1292(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 337
func Util_convExp1360(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 297
func Util_convExp1105(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// 32: decl @lune.@base.@Util.setDebugFlag
func Util_setDebugFlag(_env *LnsEnv, flag bool) {
    Util_debugFlag = flag
    
}

// 36: decl @lune.@base.@Util.setErrorCode
func Util_setErrorCode(_env *LnsEnv, code LnsInt) {
    Util_errorCode = code
    
}

// 40: decl @lune.@base.@Util.errorLog
func Util_errorLog(_env *LnsEnv, message string) {
    Lns_io_stderr.Write(_env, message + "\n")
}

// 44: decl @lune.@base.@Util.err
func Util_err(_env *LnsEnv, message string) {
    if Util_debugFlag{
        panic(message)
    }
    Util_errorLog(_env, message)
    _env.LuaVM.OS_exit(Util_errorCode)
}

// 55: decl @lune.@base.@Util.splitStr
func Util_splitStr(_env *LnsEnv, txt string,pattern string) *LnsList {
    var list *LnsList
    list = NewLnsList([]LnsAny{})
    Lns_LockEnvSync( _env, func () {
        {
            _form140, _param140, _prev140 := _env.CommonLuaVM.String_gmatch(txt, pattern)
            for {
                _work140 := _form140.(*Lns_luaValue).Call( Lns_2DDD( _param140, _prev140 ) )
                _prev140 = Lns_getFromMulti(_work140,0)
                if Lns_IsNil( _prev140 ) { break }
                token := _prev140.(string)
                list.Insert(token)
            }
        }
    })
    return list
}

// 262: decl @lune.@base.@Util.log
func Util_log(_env *LnsEnv, message string) {
    if Util_debugFlag{
        Util_errorLog(_env, message)
    }
}

// 268: decl @lune.@base.@Util.printStackTrace
func Util_printStackTrace(_env *LnsEnv) {
    Util_errorLog(_env, Depend_getStackTrace(_env))
}


// 277: decl @lune.@base.@Util.getReadyCode
func Util_getReadyCode(_env *LnsEnv, depPath string,tgtPath string) bool {
    __func__ := "@lune.@base.@Util.getReadyCode"
    var tgtTime LnsReal
    var depTime LnsReal
    
    {
        _tgtTime, _depTime := Depend_getFileLastModifiedTime(_env, tgtPath), Depend_getFileLastModifiedTime(_env, depPath)
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
    Log_log(_env, Log_Level__Warn, __func__, 288, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.LuaVM.String_format("not ready %g < %g : %s, %s", []LnsAny{tgtTime, depTime, tgtPath, depPath})
    }))
    
    return false
}

// 293: decl @lune.@base.@Util.scriptPath2Module
func Util_scriptPath2Module(_env *LnsEnv, path string) string {
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(path,"^/", nil, nil))){
        Util_err(_env, "script must be relative-path -- " + path)
    }
    var mod string
    mod = Util_convExp1105(Lns_2DDD(_env.LuaVM.String_gsub(Lns_car(_env.LuaVM.String_gsub(path,"^./", "")).(string),"/", ".")))
    return Lns_car(_env.LuaVM.String_gsub(mod, "%.lns$", "")).(string)
}

// 301: decl @lune.@base.@Util.scriptPath2ModuleFromProjDir
func Util_scriptPath2ModuleFromProjDir(_env *LnsEnv, path string,projDir LnsAny) string {
    if projDir != nil{
        projDir_235 := projDir.(string)
        var workpath string
        if Lns_op_not(Lns_car(_env.LuaVM.String_find(projDir_235,"/$", nil, nil))){
            workpath = projDir_235 + "/"
            
        } else { 
            workpath = projDir_235
            
        }
        path = Util_convExp1186(Lns_2DDD(_env.LuaVM.String_gsub(path,"^" + workpath, "")))
        
    }
    return Util_scriptPath2Module(_env, path)
}

// 314: decl @lune.@base.@Util.pathJoin
func Util_pathJoin(_env *LnsEnv, dir string,path string) string {
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(path,"^/", nil, nil))){
        return path
    }
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(dir,"/$", nil, nil))){
        return _env.LuaVM.String_format("%s%s", []LnsAny{dir, path})
    }
    return _env.LuaVM.String_format("%s/%s", []LnsAny{dir, path})
}

// 324: decl @lune.@base.@Util.parentPath
func Util_parentPath(_env *LnsEnv, path string) string {
    if Lns_isCondTrue( Lns_car(_env.LuaVM.String_find(path,"/$", nil, nil))){
        path = Util_convExp1292(Lns_2DDD(_env.LuaVM.String_gsub(path,"/$", "")))
        
    }
    return Lns_car(_env.LuaVM.String_gsub(path,"/[^/]+$", "")).(string)
}

// 331: decl @lune.@base.@Util.searchProjDir
func Util_searchProjDir(_env *LnsEnv, dir string) LnsAny {
    var work string
    work = dir
    for work != "" {
        if Depend_existFile(_env, Util_pathJoin(_env, work, "lune.js")){
            return work
        }
        work = Util_convExp1360(Lns_2DDD(_env.LuaVM.String_gsub(work,"/[^/]+$", "")))
        
    }
    return nil
}

// declaration Class -- OrderedSet
type Util_OrderedSetMtd interface {
    Add(_env *LnsEnv, arg1 LnsAny) bool
    Clone(_env *LnsEnv) *Util_OrderedSet
    Get_list(_env *LnsEnv) *LnsList
    Has(_env *LnsEnv, arg1 LnsAny) bool
    RemoveLast(_env *LnsEnv)
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
func NewUtil_OrderedSet(_env *LnsEnv) *Util_OrderedSet {
    obj := &Util_OrderedSet{}
    obj.FP = obj
    obj.InitUtil_OrderedSet(_env)
    return obj
}
func (self *Util_OrderedSet) Get_list(_env *LnsEnv) *LnsList{ return self.list }
// 69: DeclConstr
func (self *Util_OrderedSet) InitUtil_OrderedSet(_env *LnsEnv) {
    self.set = NewLnsSet([]LnsAny{})
    
    self.list = NewLnsList([]LnsAny{})
    
}

// 74: decl @lune.@base.@Util.OrderedSet.add
func (self *Util_OrderedSet) Add(_env *LnsEnv, _val LnsAny) bool {
    val := _val
    if Lns_op_not(self.set.Has(val)){
        self.set.Add(val)
        self.list.Insert(val)
        return true
    }
    return false
}

// 83: decl @lune.@base.@Util.OrderedSet.clone
func (self *Util_OrderedSet) Clone(_env *LnsEnv) *Util_OrderedSet {
    var obj *Util_OrderedSet
    obj = NewUtil_OrderedSet(_env)
    for _, _val := range( self.list.Items ) {
        val := _val
        obj.set.Add(val)
        obj.list.Insert(val)
    }
    return obj
}

// 92: decl @lune.@base.@Util.OrderedSet.has
func (self *Util_OrderedSet) Has(_env *LnsEnv, _val LnsAny) bool {
    val := _val
    return self.set.Has(val)
}

// 96: decl @lune.@base.@Util.OrderedSet.removeLast
func (self *Util_OrderedSet) RemoveLast(_env *LnsEnv) {
    if self.list.Len() == 0{
        Util_err(_env, "empty")
    }
    self.set.Del(self.list.GetAt(self.list.Len()))
    self.list.Remove(nil)
}


// declaration Class -- OrderdMap
type Util_OrderdMapMtd interface {
    Add(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny)
    Clear(_env *LnsEnv)
    Get_keyList(_env *LnsEnv) *LnsList
    Get_map(_env *LnsEnv) *LnsMap
}
type Util_OrderdMap struct {
    _map *LnsMap
    keyList *LnsList
    FP Util_OrderdMapMtd
}
func Util_OrderdMap2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_OrderdMap).FP
}
type Util_OrderdMapDownCast interface {
    ToUtil_OrderdMap() *Util_OrderdMap
}
func Util_OrderdMapDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_OrderdMapDownCast)
    if ok { return work.ToUtil_OrderdMap() }
    return nil
}
func (obj *Util_OrderdMap) ToUtil_OrderdMap() *Util_OrderdMap {
    return obj
}
func NewUtil_OrderdMap(_env *LnsEnv) *Util_OrderdMap {
    obj := &Util_OrderdMap{}
    obj.FP = obj
    obj.InitUtil_OrderdMap(_env)
    return obj
}
func (self *Util_OrderdMap) Get_map(_env *LnsEnv) *LnsMap{ return self._map }
func (self *Util_OrderdMap) Get_keyList(_env *LnsEnv) *LnsList{ return self.keyList }
// 109: DeclConstr
func (self *Util_OrderdMap) InitUtil_OrderdMap(_env *LnsEnv) {
    self._map = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.keyList = NewLnsList([]LnsAny{})
    
}

// 114: decl @lune.@base.@Util.OrderdMap.clear
func (self *Util_OrderdMap) Clear(_env *LnsEnv) {
    self._map = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.keyList = NewLnsList([]LnsAny{})
    
}

// 119: decl @lune.@base.@Util.OrderdMap.add
func (self *Util_OrderdMap) Add(_env *LnsEnv, _key LnsAny,_val LnsAny) {
    key := _key
    val := _val
    if Lns_isCondTrue( self._map.Get(key)){
        return 
    }
    self._map.Set(key,val)
    self.keyList.Insert(key)
}


// declaration Class -- memStream
type Util_memStreamMtd interface {
    Close(_env *LnsEnv)
    Flush(_env *LnsEnv)
    Get_txt(_env *LnsEnv) string
    Write(_env *LnsEnv, arg1 string)(LnsAny, LnsAny)
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
func NewUtil_memStream(_env *LnsEnv) *Util_memStream {
    obj := &Util_memStream{}
    obj.FP = obj
    obj.InitUtil_memStream(_env)
    return obj
}
// 130: DeclConstr
func (self *Util_memStream) InitUtil_memStream(_env *LnsEnv) {
    self.txt = NewStr_Builder(_env)
    
}

// 133: decl @lune.@base.@Util.memStream.get_txt
func (self *Util_memStream) Get_txt(_env *LnsEnv) string {
    self.txt.FP.Flush(_env)
    return self.txt.FP.Get_txt(_env)
}

// 137: decl @lune.@base.@Util.memStream.write
func (self *Util_memStream) Write(_env *LnsEnv, val string)(LnsAny, LnsAny) {
    self.txt.FP.Add(_env, val)
    return self.FP, nil
}

// 141: decl @lune.@base.@Util.memStream.close
func (self *Util_memStream) Close(_env *LnsEnv) {
    self.txt.FP.Flush(_env)
}

// 144: decl @lune.@base.@Util.memStream.flush
func (self *Util_memStream) Flush(_env *LnsEnv) {
    self.txt.FP.Flush(_env)
}


// declaration Class -- NullOStream
type Util_NullOStreamMtd interface {
    Close(_env *LnsEnv)
    Flush(_env *LnsEnv)
    Write(_env *LnsEnv, arg1 string)(LnsAny, LnsAny)
}
type Util_NullOStream struct {
    FP Util_NullOStreamMtd
}
func Util_NullOStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_NullOStream).FP
}
type Util_NullOStreamDownCast interface {
    ToUtil_NullOStream() *Util_NullOStream
}
func Util_NullOStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_NullOStreamDownCast)
    if ok { return work.ToUtil_NullOStream() }
    return nil
}
func (obj *Util_NullOStream) ToUtil_NullOStream() *Util_NullOStream {
    return obj
}
func NewUtil_NullOStream(_env *LnsEnv) *Util_NullOStream {
    obj := &Util_NullOStream{}
    obj.FP = obj
    obj.InitUtil_NullOStream(_env)
    return obj
}
func (self *Util_NullOStream) InitUtil_NullOStream(_env *LnsEnv) {
}
// 150: decl @lune.@base.@Util.NullOStream.write
func (self *Util_NullOStream) Write(_env *LnsEnv, val string)(LnsAny, LnsAny) {
    return self.FP, nil
}

// 153: decl @lune.@base.@Util.NullOStream.close
func (self *Util_NullOStream) Close(_env *LnsEnv) {
}

// 155: decl @lune.@base.@Util.NullOStream.flush
func (self *Util_NullOStream) Flush(_env *LnsEnv) {
}


type Util_SourceStream interface {
        PopIndent(_env *LnsEnv)
        PushIndent(_env *LnsEnv, arg1 LnsAny)
        ReturnToSource(_env *LnsEnv)
        SwitchToHeader(_env *LnsEnv)
        Write(_env *LnsEnv, arg1 string)
        Writeln(_env *LnsEnv, arg1 string)
}
func Lns_cast2Util_SourceStream( obj LnsAny ) LnsAny {
    if _, ok := obj.(Util_SourceStream); ok { 
        return obj
    }
    return nil
}

// declaration Class -- SimpleSourceOStream
type Util_SimpleSourceOStreamMtd interface {
    get_indent(_env *LnsEnv) LnsInt
    PopIndent(_env *LnsEnv)
    PushIndent(_env *LnsEnv, arg1 LnsAny)
    ReturnToSource(_env *LnsEnv)
    SwitchToHeader(_env *LnsEnv)
    Write(_env *LnsEnv, arg1 string)
    Writeln(_env *LnsEnv, arg1 string)
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
func NewUtil_SimpleSourceOStream(_env *LnsEnv, arg1 Lns_oStream, arg2 LnsAny, arg3 LnsInt) *Util_SimpleSourceOStream {
    obj := &Util_SimpleSourceOStream{}
    obj.FP = obj
    obj.InitUtil_SimpleSourceOStream(_env, arg1, arg2, arg3)
    return obj
}
// 178: DeclConstr
func (self *Util_SimpleSourceOStream) InitUtil_SimpleSourceOStream(_env *LnsEnv, stream Lns_oStream,headStream LnsAny,stepIndent LnsInt) {
    self.srcStream = stream
    
    self.nowStream = stream
    
    self.headStream = Lns_unwrapDefault( headStream, stream).(Lns_oStream)
    
    self.needIndent = true
    
    self.curLineNo = 0
    
    self.stepIndent = stepIndent
    
    self.indentQueue = NewLnsList([]LnsAny{0})
    
}

// 188: decl @lune.@base.@Util.SimpleSourceOStream.get_indent
func (self *Util_SimpleSourceOStream) get_indent(_env *LnsEnv) LnsInt {
    if self.indentQueue.Len() > 0{
        return self.indentQueue.GetAt(self.indentQueue.Len()).(LnsInt)
    }
    return 0
}

// 209: decl @lune.@base.@Util.SimpleSourceOStream.write
func (self *Util_SimpleSourceOStream) Write(_env *LnsEnv, txt string) {
    var stream Lns_oStream
    stream = self.nowStream
    for _, _line := range( Str_getLineList(_env, txt).Items ) {
        line := _line.(string)
        if self.needIndent{
            stream.Write(_env, _env.LuaVM.String_rep(" ", self.FP.get_indent(_env)))
            self.needIndent = false
            
        }
        if Lns_isCondTrue( _env.PopVal( _env.IncStack() ||
            _env.SetStackVal( len(line) > 0) &&
            _env.SetStackVal( LnsInt(line[len(line)-1]) == 10) ).(bool)){
            self.curLineNo = self.curLineNo + 1
            
        }
        stream.Write(_env, line)
    }
}

// 236: decl @lune.@base.@Util.SimpleSourceOStream.writeln
func (self *Util_SimpleSourceOStream) Writeln(_env *LnsEnv, txt string) {
    self.FP.Write(_env, txt)
    self.FP.Write(_env, "\n")
    self.needIndent = true
    
}

// 242: decl @lune.@base.@Util.SimpleSourceOStream.pushIndent
func (self *Util_SimpleSourceOStream) PushIndent(_env *LnsEnv, newIndent LnsAny) {
    var indent LnsInt
    indent = Lns_unwrapDefault( newIndent, self.FP.get_indent(_env) + self.stepIndent).(LnsInt)
    self.indentQueue.Insert(indent)
}

// 247: decl @lune.@base.@Util.SimpleSourceOStream.popIndent
func (self *Util_SimpleSourceOStream) PopIndent(_env *LnsEnv) {
    if self.indentQueue.Len() == 0{
        Util_err(_env, "self.indentQueue == 0")
    }
    self.indentQueue.Remove(nil)
}

// 254: decl @lune.@base.@Util.SimpleSourceOStream.switchToHeader
func (self *Util_SimpleSourceOStream) SwitchToHeader(_env *LnsEnv) {
    self.nowStream = self.headStream
    
}

// 257: decl @lune.@base.@Util.SimpleSourceOStream.returnToSource
func (self *Util_SimpleSourceOStream) ReturnToSource(_env *LnsEnv) {
    self.nowStream = self.srcStream
    
}


func Lns_Util_init(_env *LnsEnv) {
    if init_Util { return }
    init_Util = true
    Util__mod__ = "@lune.@base.@Util"
    Lns_InitMod()
    Lns_LuaVer_init( _env )
    Lns_Depend_init(_env)
    Lns_Log_init(_env)
    Lns_Str_init(_env)
    Util_debugFlag = true
    Util_errorCode = 1
}
func init() {
    init_Util = false
}
