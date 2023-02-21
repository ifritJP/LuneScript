// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Log bool
var Log__mod__ string
// decl enum -- Level 
type Log_Level = LnsInt
const Log_Level__Debug = 5
const Log_Level__Err = 1
const Log_Level__Fatal = 0
const Log_Level__Info = 4
const Log_Level__Log = 3
const Log_Level__Trace = 6
const Log_Level__Warn = 2
var Log_LevelList_ = NewLnsList( []LnsAny {
  Log_Level__Fatal,
  Log_Level__Err,
  Log_Level__Warn,
  Log_Level__Log,
  Log_Level__Info,
  Log_Level__Debug,
  Log_Level__Trace,
})
func Log_Level_get__allList(_env *LnsEnv) *LnsList{
    return Log_LevelList_
}
var Log_LevelMap_ = map[LnsInt]string {
  Log_Level__Debug: "Level.Debug",
  Log_Level__Err: "Level.Err",
  Log_Level__Fatal: "Level.Fatal",
  Log_Level__Info: "Level.Info",
  Log_Level__Log: "Level.Log",
  Log_Level__Trace: "Level.Trace",
  Log_Level__Warn: "Level.Warn",
}
func Log_Level__from(_env *LnsEnv, arg1 LnsInt) LnsAny{
    if _, ok := Log_LevelMap_[arg1]; ok { return arg1 }
    return nil
}

func Log_Level_getTxt(arg1 LnsInt) string {
    return Log_LevelMap_[arg1];
}
var Log_name2levelMap *LnsMap
var Log_control *Log_Control
type Log_CreateMessage func (_env *LnsEnv) string
// 54: decl @lune.@base.@Log.str2level
func Log_str2level(_env *LnsEnv, txt string) LnsAny {
    return Log_name2levelMap.Get(txt)
}


// 83: decl @lune.@base.@Log.setLevel
func Log_setLevel(_env *LnsEnv, level LnsInt) {
    Log_control = NewLog_Control(_env, level)
    if level >= Log_Level__Log{
        Depend_setRuntimeLog(_env, true)
    }
}

// 90: decl @lune.@base.@Log.getLevel
func Log_getLevel(_env *LnsEnv) LnsInt {
    return Log_control.FP.get_level(_env)
}

// 94: decl @lune.@base.@Log.log
func Log_log(_env *LnsEnv, level LnsInt,funcName string,lineNo LnsInt,callback Log_CreateMessage) {
    Log_control.FP.log(_env, level, funcName, lineNo, callback)
}

// 98: decl @lune.@base.@Log.direct
func Log_direct(_env *LnsEnv, level LnsInt,funcName string,lineNo LnsInt,mess string) {
    Log_control.FP.direct(_env, level, funcName, lineNo, mess)
}

// 61: decl @lune.@base.@Log.Control.log
func (self *Log_Control) log(_env *LnsEnv, level LnsInt,funcName string,lineNo LnsInt,callback Log_CreateMessage) {
    var logStream Lns_oStream
    logStream = Lns_io_stderr
    if level <= self.level{
        var nowClock LnsReal
        nowClock = _env.GetVM().OS_clock()
        logStream.Write(_env, _env.GetVM().String_format("%6d:%s:%s:%d:", Lns_2DDD((LnsInt)((nowClock * LnsReal(1000))), Log_Level_getTxt( level), funcName, lineNo)))
        logStream.Write(_env, callback(_env))
        logStream.Write(_env, "\n")
    }
}
// 73: decl @lune.@base.@Log.Control.direct
func (self *Log_Control) direct(_env *LnsEnv, level LnsInt,funcName string,lineNo LnsInt,mess string) {
    self.FP.log(_env, level, funcName, lineNo, Log_CreateMessage(func(_env *LnsEnv) string {
        return mess
    }))
}
// declaration Class -- Control
type Log_ControlMtd interface {
    direct(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 string)
    get_level(_env *LnsEnv) LnsInt
    log(_env *LnsEnv, arg1 LnsInt, arg2 string, arg3 LnsInt, arg4 Log_CreateMessage)
}
type Log_Control struct {
    level LnsInt
    FP Log_ControlMtd
}
func Log_Control2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Log_Control).FP
}
func Log_Control_toSlice(slice []LnsAny) []*Log_Control {
    ret := make([]*Log_Control, len(slice))
    for index, val := range slice {
        ret[index] = val.(Log_ControlDownCast).ToLog_Control()
    }
    return ret
}
type Log_ControlDownCast interface {
    ToLog_Control() *Log_Control
}
func Log_ControlDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Log_ControlDownCast)
    if ok { return work.ToLog_Control() }
    return nil
}
func (obj *Log_Control) ToLog_Control() *Log_Control {
    return obj
}
func NewLog_Control(_env *LnsEnv, arg1 LnsInt) *Log_Control {
    obj := &Log_Control{}
    obj.FP = obj
    obj.InitLog_Control(_env, arg1)
    return obj
}
func (self *Log_Control) InitLog_Control(_env *LnsEnv, arg1 LnsInt) {
    self.level = arg1
}
func (self *Log_Control) get_level(_env *LnsEnv) LnsInt{ return self.level }

func Lns_Log_init(_env *LnsEnv) {
    if init_Log { return }
    init_Log = true
    Log__mod__ = "@lune.@base.@Log"
    Lns_InitMod()
    Lns_LuaVer_init( _env )
    Lns_Depend_init(_env)
    {
        var work *LnsMap
        work = NewLnsMap( map[LnsAny]LnsAny{})
        work.Set("fatal",Log_Level__Fatal)
        work.Set("error",Log_Level__Err)
        work.Set("warn",Log_Level__Warn)
        work.Set("log",Log_Level__Log)
        work.Set("info",Log_Level__Info)
        work.Set("debug",Log_Level__Debug)
        work.Set("trace",Log_Level__Trace)
        Log_name2levelMap = work
    }
    Log_control = NewLog_Control(_env, Log_Level__Err)
}
func init() {
    init_Log = false
}
