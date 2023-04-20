// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Util bool
var Util__mod__ string
var Util_consoleOStream Lns_oStream
var Util_errStream Lns_oStream
var Util_debugFlag bool
var Util_errorCode LnsInt
var Util_s_fs Util_FSIF
type Util_ConsoleWriter func (_env *LnsEnv, arg1 string)
// for 456
func Util_convExp0_2011(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 473
func Util_convExp0_2098(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 432
func Util_convExp0_1878(arg1 []LnsAny) LnsAny {
    return Lns_getFromMulti( arg1, 0 )
}
// for 444
func Util_convExp0_1941(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 475
func Util_convExp0_2114(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 488
func Util_convExp0_2171(arg1 []LnsAny) string {
    return Lns_getFromMulti( arg1, 0 ).(string)
}
// for 66
func Util_convExp0_826(arg1 []LnsAny) []LnsAny {
    return Lns_2DDD( arg1[0:])
}
// 47: decl @lune.@base.@Util.getConsoleOStream
func Util_getConsoleOStream(_env *LnsEnv) Lns_oStream {
    return Util_consoleOStream
}

// 51: decl @lune.@base.@Util.setConsoleOStream
func Util_setConsoleOStream(_env *LnsEnv, stream Lns_oStream,streamErr Lns_oStream) {
    Lns_LockEnvSync( _env, 52, func () {
        Util_consoleOStream = stream
        Util_errStream = streamErr
    })
}

// 58: decl @lune.@base.@Util.setConsoleOStreamWithWriter
func Util_setConsoleOStreamWithWriter(_env *LnsEnv, writer Util_ConsoleWriter,errWriter Util_ConsoleWriter) {
    Util_setConsoleOStream(_env, NewUtil_ConsoleAdapter(_env, writer).FP, NewUtil_ConsoleAdapter(_env, errWriter).FP)
}

// 65: decl @lune.@base.@Util.println
func Util_println(_env *LnsEnv, ddd []LnsAny) {
    var list *LnsList2_[LnsAny]
    list = NewLnsList2_[LnsAny](Lns_2Slice[LnsAny]( Lns_2DDD(ddd)))
    Lns_LockEnvSync( _env, 67, func () {
        for _index, _arg := range( list.Items ) {
            index := _index + 1
            arg := _arg
            Util_consoleOStream.Write(_env, _env.GetVM().String_format("%s", Lns_2DDD(arg)))
            if index != list.Len(){
                Util_consoleOStream.Write(_env, "\t")
            } else { 
                Util_consoleOStream.Write(_env, "\n")
            }
        }
    })
}

// 80: decl @lune.@base.@Util.setDebugFlag
func Util_setDebugFlag(_env *LnsEnv, flag bool) {
    Util_debugFlag = flag
}

// 84: decl @lune.@base.@Util.setErrorCode
func Util_setErrorCode(_env *LnsEnv, code LnsInt) {
    Util_errorCode = code
}

// 88: decl @lune.@base.@Util.debugLog
func Util_debugLog(_env *LnsEnv, message string) {
    Lns_LockEnvSync( _env, 89, func () {
        if Util_debugFlag{
            Util_errStream.Write(_env, message + "\n")
        }
    })
}

// 96: decl @lune.@base.@Util.errorLog
func Util_errorLog(_env *LnsEnv, message string) {
    Lns_LockEnvSync( _env, 97, func () {
        Util_errStream.Write(_env, message + "\n")
    })
}

// 102: decl @lune.@base.@Util.err
func Util_err(_env *LnsEnv, message string) {
    if Util_debugFlag{
        panic(message)
    }
    Util_errorLog(_env, message)
    _env.GetVM().OS_exit(Util_errorCode)
}

// 121: decl @lune.@base.@Util.splitStr
func Util_splitStr(_env *LnsEnv, txt string,pattern string) *LnsList2_[string] {
    var list *LnsList2_[string]
    list = NewLnsList2_[string]([]string{})
        {
            _applyForm0, _applyParam0, _applyPrev0 := _env.GetVM().String_gmatch(txt, pattern)
            for {
                _applyWork0 := _applyForm0.(*Lns_luaValue).Call( Lns_2DDD( _applyParam0, _applyPrev0 ) )
                _applyPrev0 = Lns_getFromMulti(_applyWork0,0)
                if Lns_IsNil( _applyPrev0 ) { break }
                token := _applyPrev0.(string)
                list.Insert(token)
            }
        }
    return list
}

// 131: decl @lune.@base.@Util.splitModule
func Util_splitModule(_env *LnsEnv, modPath string) *LnsList2_[string] {
    return Util_splitStr(_env, modPath, "[^%./:]+")
}

// 400: decl @lune.@base.@Util.log
func Util_log(_env *LnsEnv, message string) {
    if Util_debugFlag{
        Util_errorLog(_env, message)
    }
}

// 406: decl @lune.@base.@Util.printStackTrace
func Util_printStackTrace(_env *LnsEnv) {
    Util_errorLog(_env, Depend_getStackTrace(_env))
}

// 415: decl @lune.@base.@Util.getReadyCode
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
    Log_log(_env, Log_Level__Warn, __func__, 426, Log_CreateMessage(func(_env *LnsEnv) string {
        return _env.GetVM().String_format("not ready %g < %g : %s, %s", Lns_2DDD(tgtTime, depTime, tgtPath, depPath))
    }))
    
    return false
}

// 431: decl @lune.@base.@Util.readFile
func Util_readFile(_env *LnsEnv, path string) LnsAny {
    var fileObj Lns_luaStream
    
    {
        _fileObj := Util_convExp0_1878(Lns_2DDD(Lns_io_open(path, nil)))
        if _fileObj == nil{
            return nil
        } else {
            fileObj = _fileObj.(Lns_luaStream)
        }
    }
    var txt LnsAny
    txt = fileObj.Read(_env, "*a")
    fileObj.Close(_env)
    return txt
}

// 440: decl @lune.@base.@Util.scriptPath2Module
func Util_scriptPath2Module(_env *LnsEnv, path string) string {
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(path,"^/", nil, nil))){
        Util_err(_env, "script must be relative-path -- " + path)
    }
    var mod string
    mod = Util_convExp0_1941(Lns_2DDD(_env.GetVM().String_gsub(Lns_car(_env.GetVM().String_gsub(path,"^./", "")).(string),"/", ".")))
    return Lns_car(_env.GetVM().String_gsub(mod, "%.lns$", "")).(string)
}

// 448: decl @lune.@base.@Util.scriptPath2ModuleFromProjDir
func Util_scriptPath2ModuleFromProjDir(_env *LnsEnv, path string,projDir LnsAny) string {
    if projDir != nil{
        projDir_396 := projDir.(string)
        var workpath string
        if Lns_op_not(Lns_car(_env.GetVM().String_find(projDir_396,"/$", nil, nil))){
            workpath = projDir_396 + "/"
        } else { 
            workpath = projDir_396
        }
        path = Util_convExp0_2011(Lns_2DDD(_env.GetVM().String_gsub(path,"^" + workpath, "")))
    }
    return Util_scriptPath2Module(_env, path)
}

// 461: decl @lune.@base.@Util.pathJoin
func Util_pathJoin(_env *LnsEnv, dir string,path string) string {
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(path,"^/", nil, nil))){
        return path
    }
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(dir,"/$", nil, nil))){
        return _env.GetVM().String_format("%s%s", Lns_2DDD(dir, path))
    }
    return _env.GetVM().String_format("%s/%s", Lns_2DDD(dir, path))
}

// 471: decl @lune.@base.@Util.parentPath
func Util_parentPath(_env *LnsEnv, path string) string {
    if Lns_isCondTrue( Lns_car(_env.GetVM().String_find(path,"/$", nil, nil))){
        path = Util_convExp0_2098(Lns_2DDD(_env.GetVM().String_gsub(path,"/$", "")))
    }
    var parent string
    parent = Util_convExp0_2114(Lns_2DDD(_env.GetVM().String_gsub(path,"/[^/]+$", "")))
    if parent == path{
        return "./"
    }
    return parent
}

// 482: decl @lune.@base.@Util.searchProjDir
func Util_searchProjDir(_env *LnsEnv, dir string) LnsAny {
    var work string
    work = dir
    for work != "" {
        if Depend_existFile(_env, Util_pathJoin(_env, work, "lune.js")){
            return work
        }
        var parent string
        parent = Util_convExp0_2171(Lns_2DDD(_env.GetVM().String_gsub(work,"/[^/]+$", "")))
        if parent == work{
            return nil
        }
        work = parent
    }
    return nil
}

// 525: decl @lune.@base.@Util.setFS
func Util_setFS(_env *LnsEnv, fs Util_FSIF) {
    Util_s_fs = fs
}

// 528: decl @lune.@base.@Util.openRd
func Util_openRd(_env *LnsEnv, path string) LnsAny {
    return Util_s_fs.OpenRd(_env, path)
}


// 37: decl @lune.@base.@Util.ConsoleAdapter.write
func (self *Util_ConsoleAdapter) Write(_env *LnsEnv, txt string)(LnsAny, LnsAny) {
    self.writer(_env, txt)
    return self.FP, nil
}
// 41: decl @lune.@base.@Util.ConsoleAdapter.flush
func (self *Util_ConsoleAdapter) Flush(_env *LnsEnv) {
}
// 43: decl @lune.@base.@Util.ConsoleAdapter.close
func (self *Util_ConsoleAdapter) Close(_env *LnsEnv) {
}
// 144: decl @lune.@base.@Util.OrderedSet.add
func (self *Util_OrderedSet) Add(_env *LnsEnv, _val LnsAny) bool {
    val := _val
    if Lns_op_not(self.set.Has(val)){
        self.set.Add(val)
        self.list.Insert(val)
        return true
    }
    return false
}
// 153: decl @lune.@base.@Util.OrderedSet.clone
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
// 162: decl @lune.@base.@Util.OrderedSet.has
func (self *Util_OrderedSet) Has(_env *LnsEnv, _val LnsAny) bool {
    val := _val
    return self.set.Has(val)
}
// 166: decl @lune.@base.@Util.OrderedSet.removeLast
func (self *Util_OrderedSet) RemoveLast(_env *LnsEnv) {
    if self.list.Len() == 0{
        Util_err(_env, "empty")
    }
    self.set.Del(self.list.GetAt(self.list.Len()))
    self.list.Remove(nil)
}
// 184: decl @lune.@base.@Util.OrderdMap.clear
func (self *Util_OrderdMap) Clear(_env *LnsEnv) {
    self._map = NewLnsMap( map[LnsAny]LnsAny{})
    self.keyList = NewLnsList([]LnsAny{})
}
// 189: decl @lune.@base.@Util.OrderdMap.add
func (self *Util_OrderdMap) Add(_env *LnsEnv, _key LnsAny,_val LnsAny,overwrite bool) {
    key := _key
    val := _val
    if Lns_isCondTrue( self._map.Get(key)){
        if overwrite{
            self._map.Set(key,val)
        }
        return 
    }
    self._map.Set(key,val)
    self.keyList.Insert(key)
}
// 206: decl @lune.@base.@Util.memStream.get_txt
func (self *Util_memStream) Get_txt(_env *LnsEnv) string {
    self.txt.FP.Flush(_env)
    return self.txt.FP.Get_txt(_env)
}
// 210: decl @lune.@base.@Util.memStream.write
func (self *Util_memStream) Write(_env *LnsEnv, val string)(LnsAny, LnsAny) {
    self.txt.FP.Add(_env, val)
    return self.FP, nil
}
// 214: decl @lune.@base.@Util.memStream.close
func (self *Util_memStream) Close(_env *LnsEnv) {
    self.txt.FP.Flush(_env)
}
// 217: decl @lune.@base.@Util.memStream.flush
func (self *Util_memStream) Flush(_env *LnsEnv) {
    self.txt.FP.Flush(_env)
}
// 247: decl @lune.@base.@Util.TxtStream.getSubstring
func (self *Util_TxtStream) GetSubstring(_env *LnsEnv, fromLineNo LnsInt,toLineNo LnsAny) string {
    var txt string
    txt = ""
    var to LnsInt
    to = Lns_unwrapDefault( toLineNo, self.lineList.Len() + 1).(LnsInt)
    {
        var _forFrom0 LnsInt = fromLineNo
        var _forTo0 LnsInt = to - 1
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            index := _forWork0
            if _env.PopVal( _env.IncStack() ||
                _env.SetStackVal( index < 1) ||
                _env.SetStackVal( index > self.lineList.Len()) ).(bool){
                break
            }
            txt = _env.GetVM().String_format("%s%s", Lns_2DDD(txt, self.lineList.GetAt(index).(string)))
        }
    }
    return txt
}
// 259: decl @lune.@base.@Util.TxtStream.read
func (self *Util_TxtStream) Read(_env *LnsEnv, mode LnsAny) LnsAny {
    if mode != "*l"{
        Util_err(_env, _env.GetVM().String_format("not support -- %s", Lns_2DDD(mode)))
    }
    if self.lineNo > self.lineList.Len(){
        return nil
    }
    self.lineNo = self.lineNo + 1
    var line string
    line = self.lineList.GetAt(self.lineNo - 1).(string)
    if Str_endsWith(_env, line, "\n"){
        return _env.GetVM().String_sub(line,1, len(line) - 1)
    }
    return line
}
// 273: decl @lune.@base.@Util.TxtStream.close
func (self *Util_TxtStream) Close(_env *LnsEnv) {
}
// 280: decl @lune.@base.@Util.NullOStream.write
func (self *Util_NullOStream) Write(_env *LnsEnv, val string)(LnsAny, LnsAny) {
    return self.FP, nil
}
// 283: decl @lune.@base.@Util.NullOStream.close
func (self *Util_NullOStream) Close(_env *LnsEnv) {
}
// 285: decl @lune.@base.@Util.NullOStream.flush
func (self *Util_NullOStream) Flush(_env *LnsEnv) {
}
// 333: decl @lune.@base.@Util.SimpleSourceOStream.get_indent
func (self *Util_SimpleSourceOStream) get_indent(_env *LnsEnv) LnsInt {
    if self.indentQueue.Len() > 0{
        return self.indentQueue.GetAt(self.indentQueue.Len())
    }
    return 0
}
// 340: decl @lune.@base.@Util.SimpleSourceOStream.writeRaw
func (self *Util_SimpleSourceOStream) WriteRaw(_env *LnsEnv, txt string) {
    if self.needIndent{
        self.nowStream.Write(_env, self.indentSpace)
        self.needIndent = false
    }
    self.nowStream.Write(_env, txt)
}
// 349: decl @lune.@base.@Util.SimpleSourceOStream.write
func (self *Util_SimpleSourceOStream) Write(_env *LnsEnv, txt string) {
    if Lns_op_not(Lns_car(_env.GetVM().String_find(txt,"\n", 1, true))){
        self.FP.WriteRaw(_env, txt)
        return 
    }
    var stream Lns_oStream
    stream = self.nowStream
    for _, _line := range( Str_getLineList(_env, txt).Items ) {
        line := _line.(string)
        if self.needIndent{
            stream.Write(_env, self.indentSpace)
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
// 369: decl @lune.@base.@Util.SimpleSourceOStream.writeln
func (self *Util_SimpleSourceOStream) Writeln(_env *LnsEnv, txt string) {
    self.FP.Write(_env, txt)
    self.FP.Write(_env, "\n")
    self.needIndent = true
}
// 375: decl @lune.@base.@Util.SimpleSourceOStream.pushIndent
func (self *Util_SimpleSourceOStream) PushIndent(_env *LnsEnv, newIndent LnsAny) {
    var indent LnsInt
    indent = Lns_unwrapDefault( newIndent, self.FP.get_indent(_env) + self.stepIndent).(LnsInt)
    self.indentQueue.Insert(LnsInt(indent))
    if indent > Util_SimpleSourceOStream__indentSpaceList.Len(){
        Util_err(_env, _env.GetVM().String_format("overflow indent -- %d", Lns_2DDD(indent)))
    }
    self.indentSpace = Util_SimpleSourceOStream__indentSpaceList.GetAt(indent + 1)
}
// 384: decl @lune.@base.@Util.SimpleSourceOStream.popIndent
func (self *Util_SimpleSourceOStream) PopIndent(_env *LnsEnv) {
    if self.indentQueue.Len() == 0{
        Util_err(_env, "self.indentQueue == 0")
    }
    self.indentQueue.Remove(nil)
    self.indentSpace = Util_SimpleSourceOStream__indentSpaceList.GetAt(self.FP.get_indent(_env) + 1)
}
// 392: decl @lune.@base.@Util.SimpleSourceOStream.switchToHeader
func (self *Util_SimpleSourceOStream) SwitchToHeader(_env *LnsEnv) {
    self.nowStream = self.headStream
}
// 395: decl @lune.@base.@Util.SimpleSourceOStream.returnToSource
func (self *Util_SimpleSourceOStream) ReturnToSource(_env *LnsEnv) {
    self.nowStream = self.srcStream
}
// 502: decl @lune.@base.@Util.FS.openRd
func (self *Util_FS) OpenRd(_env *LnsEnv, path string) LnsAny {
    return Lns_car(Lns_io_open(path, nil))
}
// 515: decl @lune.@base.@Util.MappedFS.openRd
func (self *Util_MappedFS) OpenRd(_env *LnsEnv, path string) LnsAny {
    var bin string
    
    {
        _bin := self.path2bin.Get(path)
        if _bin == nil{
            return nil
        } else {
            bin = _bin.(string)
        }
    }
    return NewUtil_TxtStream(_env, bin).FP
}
// declaration Class -- ConsoleAdapter
type Util_ConsoleAdapterMtd interface {
    Close(_env *LnsEnv)
    Flush(_env *LnsEnv)
    Write(_env *LnsEnv, arg1 string)(LnsAny, LnsAny)
}
type Util_ConsoleAdapter struct {
    writer Util_ConsoleWriter
    FP Util_ConsoleAdapterMtd
}
func Util_ConsoleAdapter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_ConsoleAdapter).FP
}
func Util_ConsoleAdapter_toSlice(slice []LnsAny) []*Util_ConsoleAdapter {
    ret := make([]*Util_ConsoleAdapter, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_ConsoleAdapterDownCast).ToUtil_ConsoleAdapter()
    }
    return ret
}
type Util_ConsoleAdapterDownCast interface {
    ToUtil_ConsoleAdapter() *Util_ConsoleAdapter
}
func Util_ConsoleAdapterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_ConsoleAdapterDownCast)
    if ok { return work.ToUtil_ConsoleAdapter() }
    return nil
}
func (obj *Util_ConsoleAdapter) ToUtil_ConsoleAdapter() *Util_ConsoleAdapter {
    return obj
}
func NewUtil_ConsoleAdapter(_env *LnsEnv, arg1 Util_ConsoleWriter) *Util_ConsoleAdapter {
    obj := &Util_ConsoleAdapter{}
    obj.FP = obj
    obj.InitUtil_ConsoleAdapter(_env, arg1)
    return obj
}
func (self *Util_ConsoleAdapter) InitUtil_ConsoleAdapter(_env *LnsEnv, arg1 Util_ConsoleWriter) {
    self.writer = arg1
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
func Util_OrderedSet_toSlice(slice []LnsAny) []*Util_OrderedSet {
    ret := make([]*Util_OrderedSet, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_OrderedSetDownCast).ToUtil_OrderedSet()
    }
    return ret
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
// 139: DeclConstr
func (self *Util_OrderedSet) InitUtil_OrderedSet(_env *LnsEnv) {
    self.set = NewLnsSet([]LnsAny{})
    self.list = NewLnsList([]LnsAny{})
}


// declaration Class -- OrderdMap
type Util_OrderdMapMtd interface {
    Add(_env *LnsEnv, arg1 LnsAny, arg2 LnsAny, arg3 bool)
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
func Util_OrderdMap_toSlice(slice []LnsAny) []*Util_OrderdMap {
    ret := make([]*Util_OrderdMap, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_OrderdMapDownCast).ToUtil_OrderdMap()
    }
    return ret
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
// 179: DeclConstr
func (self *Util_OrderdMap) InitUtil_OrderdMap(_env *LnsEnv) {
    self._map = NewLnsMap( map[LnsAny]LnsAny{})
    self.keyList = NewLnsList([]LnsAny{})
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
func Util_memStream_toSlice(slice []LnsAny) []*Util_memStream {
    ret := make([]*Util_memStream, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_memStreamDownCast).ToUtil_memStream()
    }
    return ret
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
// 203: DeclConstr
func (self *Util_memStream) InitUtil_memStream(_env *LnsEnv) {
    self.txt = NewStr_Builder(_env)
}


// declaration Class -- TxtStream
type Util_TxtStreamMtd interface {
    Close(_env *LnsEnv)
    GetSubstring(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny) string
    Get_lineNo(_env *LnsEnv) LnsInt
    Get_txt(_env *LnsEnv) string
    Read(_env *LnsEnv, arg1 LnsAny) LnsAny
}
type Util_TxtStream struct {
    txt string
    lineList *LnsList
    start LnsInt
    eof bool
    lineNo LnsInt
    FP Util_TxtStreamMtd
}
func Util_TxtStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_TxtStream).FP
}
func Util_TxtStream_toSlice(slice []LnsAny) []*Util_TxtStream {
    ret := make([]*Util_TxtStream, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_TxtStreamDownCast).ToUtil_TxtStream()
    }
    return ret
}
type Util_TxtStreamDownCast interface {
    ToUtil_TxtStream() *Util_TxtStream
}
func Util_TxtStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_TxtStreamDownCast)
    if ok { return work.ToUtil_TxtStream() }
    return nil
}
func (obj *Util_TxtStream) ToUtil_TxtStream() *Util_TxtStream {
    return obj
}
func NewUtil_TxtStream(_env *LnsEnv, arg1 string) *Util_TxtStream {
    obj := &Util_TxtStream{}
    obj.FP = obj
    obj.InitUtil_TxtStream(_env, arg1)
    return obj
}
func (self *Util_TxtStream) Get_txt(_env *LnsEnv) string{ return self.txt }
func (self *Util_TxtStream) Get_lineNo(_env *LnsEnv) LnsInt{ return self.lineNo }
// 228: DeclConstr
func (self *Util_TxtStream) InitUtil_TxtStream(_env *LnsEnv, txt string) {
    self.txt = txt
    self.start = 1
    self.eof = false
    self.lineList = Str_getLineList(_env, self.txt)
    self.lineNo = 1
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
func Util_NullOStream_toSlice(slice []LnsAny) []*Util_NullOStream {
    ret := make([]*Util_NullOStream, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_NullOStreamDownCast).ToUtil_NullOStream()
    }
    return ret
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
var Util_SimpleSourceOStream__indentSpaceList *LnsList2_[string]
// 299: decl @lune.@base.@Util.SimpleSourceOStream.___init
func Util_SimpleSourceOStream____init_1_(_env *LnsEnv) {
    var list *LnsList2_[string]
    list = NewLnsList2_[string]([]string{})
    var txt string
    txt = ""
    {
        var _forFrom0 LnsInt = 1
        var _forTo0 LnsInt = 100
        for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
            list.Insert(txt)
            txt = txt + " "
        }
    }
    Util_SimpleSourceOStream__indentSpaceList = list
}

type Util_SimpleSourceOStreamMtd interface {
    get_indent(_env *LnsEnv) LnsInt
    Get_stepIndent(_env *LnsEnv) LnsInt
    PopIndent(_env *LnsEnv)
    PushIndent(_env *LnsEnv, arg1 LnsAny)
    ReturnToSource(_env *LnsEnv)
    SwitchToHeader(_env *LnsEnv)
    Write(_env *LnsEnv, arg1 string)
    WriteRaw(_env *LnsEnv, arg1 string)
    Writeln(_env *LnsEnv, arg1 string)
}
type Util_SimpleSourceOStream struct {
    nowStream Lns_oStream
    srcStream Lns_oStream
    headStream Lns_oStream
    needIndent bool
    stepIndent LnsInt
    curLineNo LnsInt
    indentQueue *LnsList2_[LnsInt]
    indentSpace string
    indentSpaceList *LnsList2_[string]
    FP Util_SimpleSourceOStreamMtd
}
func Util_SimpleSourceOStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_SimpleSourceOStream).FP
}
func Util_SimpleSourceOStream_toSlice(slice []LnsAny) []*Util_SimpleSourceOStream {
    ret := make([]*Util_SimpleSourceOStream, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_SimpleSourceOStreamDownCast).ToUtil_SimpleSourceOStream()
    }
    return ret
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
func (self *Util_SimpleSourceOStream) Get_stepIndent(_env *LnsEnv) LnsInt{ return self.stepIndent }
// 322: DeclConstr
func (self *Util_SimpleSourceOStream) InitUtil_SimpleSourceOStream(_env *LnsEnv, stream Lns_oStream,headStream LnsAny,stepIndent LnsInt) {
    self.srcStream = stream
    self.nowStream = stream
    self.headStream = Lns_unwrapDefault( headStream, stream).(Lns_oStream)
    self.needIndent = true
    self.curLineNo = 0
    self.stepIndent = stepIndent
    self.indentQueue = NewLnsList2_[LnsInt](Lns_2DDDGen[LnsInt](0))
    self.indentSpace = ""
}


type Util_FSIF interface {
        OpenRd(_env *LnsEnv, arg1 string) LnsAny
}
func Lns_cast2Util_FSIF( obj LnsAny ) LnsAny {
    if _, ok := obj.(Util_FSIF); ok { 
        return obj
    }
    return nil
}

// declaration Class -- FS
type Util_FSMtd interface {
    OpenRd(_env *LnsEnv, arg1 string) LnsAny
}
type Util_FS struct {
    FP Util_FSMtd
}
func Util_FS2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_FS).FP
}
func Util_FS_toSlice(slice []LnsAny) []*Util_FS {
    ret := make([]*Util_FS, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_FSDownCast).ToUtil_FS()
    }
    return ret
}
type Util_FSDownCast interface {
    ToUtil_FS() *Util_FS
}
func Util_FSDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_FSDownCast)
    if ok { return work.ToUtil_FS() }
    return nil
}
func (obj *Util_FS) ToUtil_FS() *Util_FS {
    return obj
}
func NewUtil_FS(_env *LnsEnv) *Util_FS {
    obj := &Util_FS{}
    obj.FP = obj
    obj.InitUtil_FS(_env)
    return obj
}
func (self *Util_FS) InitUtil_FS(_env *LnsEnv) {
}

// declaration Class -- MappedFS
type Util_MappedFSMtd interface {
    OpenRd(_env *LnsEnv, arg1 string) LnsAny
}
type Util_MappedFS struct {
    path2bin *LnsMap2_[string,string]
    FP Util_MappedFSMtd
}
func Util_MappedFS2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Util_MappedFS).FP
}
func Util_MappedFS_toSlice(slice []LnsAny) []*Util_MappedFS {
    ret := make([]*Util_MappedFS, len(slice))
    for index, val := range slice {
        ret[index] = val.(Util_MappedFSDownCast).ToUtil_MappedFS()
    }
    return ret
}
type Util_MappedFSDownCast interface {
    ToUtil_MappedFS() *Util_MappedFS
}
func Util_MappedFSDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Util_MappedFSDownCast)
    if ok { return work.ToUtil_MappedFS() }
    return nil
}
func (obj *Util_MappedFS) ToUtil_MappedFS() *Util_MappedFS {
    return obj
}
func NewUtil_MappedFS(_env *LnsEnv, arg1 *LnsMap2_[string,string]) *Util_MappedFS {
    obj := &Util_MappedFS{}
    obj.FP = obj
    obj.InitUtil_MappedFS(_env, arg1)
    return obj
}
// 511: DeclConstr
func (self *Util_MappedFS) InitUtil_MappedFS(_env *LnsEnv, path2bin *LnsMap2_[string,string]) {
    self.path2bin = path2bin
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
    Util_consoleOStream = Lns_io_stdout
    Util_errStream = Lns_io_stderr
    Util_debugFlag = true
    Util_errorCode = 1
    Util_SimpleSourceOStream____init_1_(_env)
    Util_s_fs = NewUtil_FS(_env).FP
}
func init() {
    init_Util = false
}
