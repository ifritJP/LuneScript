// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Async bool
var Async__mod__ string
type Async_EndProcessFunc_3_ func (_env *LnsEnv, arg1 *Async_RunnerBase)
// 51: decl @lune.@base.@Async.Pipe.getNext
func (self *Async_Pipe) GetNext(_env *LnsEnv) LnsAny {
    if self.started{
        {
            _pipe := self.pipe
            if !Lns_IsNil( _pipe ) {
                pipe := _pipe.(*Lns__pipe)
                var val LnsAny
                
                {
                    _val := (pipe.Get(_env))
                    if _val == nil{
                        return nil
                    } else {
                        val = _val
                    }
                }
                return NewAsync_PipeItem(_env, val)
            }
        }
    }
    return self.FP.Access(_env)
}
// 64: decl @lune.@base.@Async.Pipe.run
func (self *Async_Pipe) Run(_env *LnsEnv) {
    self.FP.Setup(_env)
    self.setuped = true
    var pipe *Lns__pipe
    
    {
        _pipe := self.pipe
        if _pipe == nil{
            return 
        } else {
            pipe = _pipe.(*Lns__pipe)
        }
    }
    for  {
        var val *Async_PipeItem
        
        {
            _val := self.FP.Access(_env)
            if _val == nil{
                pipe.Put(_env, nil)
                break
            } else {
                val = _val.(*Async_PipeItem)
            }
        }
        pipe.Put(_env, val.FP.Get_item(_env))
    }
}
// 79: decl @lune.@base.@Async.Pipe.start
func (self *Async_Pipe) Start(_env *LnsEnv) {
    self.started = true
}
// 82: decl @lune.@base.@Async.Pipe.stop
func (self *Async_Pipe) Stop(_env *LnsEnv) {
    self.started = false
    if Lns_op_not(self.setuped){
        self.FP.Setup(_env)
    }
}
// 128: decl @lune.@base.@Async.RunnerBase.run
func (self *Async_RunnerBase) Run(_env *LnsEnv) {
    self.artifact = self.FP.RunMain(_env)
    self.ranFlag = true
    {
        _pipe := self.pipe
        if !Lns_IsNil( _pipe ) {
            pipe := _pipe.(*Lns__pipe)
            pipe.Put(_env, Async_RunnerBase2Stem(self))
        }
    }
}
// 167: decl @lune.@base.@Async.Waiter.startRunner
func (self *Async_Waiter) startRunner(_env *LnsEnv, runner *Async_RunnerBase,mode LnsInt,name string) bool {
    self.asyncNum = self.asyncNum + 1
    var result bool
    result = LnsRun(_env, runner.FP, mode, name)
    if _env.PopVal( _env.IncStack() ||
        _env.SetStackVal( Lns_op_not(self.pipe)) ||
        _env.SetStackVal( Lns_op_not(result)) ).(bool){
        self.asyncNum = self.asyncNum - 1
        self.finRunnerList.Insert(Async_RunnerBase2Stem(runner))
    }
    return result
}
// 183: decl @lune.@base.@Async.Waiter.wait
func (self *Async_Waiter) Wait(_env *LnsEnv, _func Async_EndProcessFunc_3_) {
    for _, _runner := range( self.finRunnerList.Items ) {
        runner := _runner.(Async_RunnerBaseDownCast).ToAsync_RunnerBase()
        _func(_env, runner)
    }
    {
        _pipe := self.pipe
        if !Lns_IsNil( _pipe ) {
            pipe := _pipe.(*Lns__pipe)
            {
                var _forFrom0 LnsInt = 1
                var _forTo0 LnsInt = self.asyncNum
                for _forWork0 := _forFrom0; _forWork0 <= _forTo0; _forWork0++ {
                    var runner *Async_RunnerBase
                    
                    {
                        _runner := Async_RunnerBaseDownCastF(pipe.Get(_env))
                        if _runner == nil{
                            break
                        } else {
                            runner = _runner.(*Async_RunnerBase)
                        }
                    }
                    _func(_env, runner)
                }
            }
        }
    }
}
// declaration Class -- PipeItem
type Async_PipeItemMtd interface {
    Get_item(_env *LnsEnv) LnsAny
}
type Async_PipeItem struct {
    item LnsAny
    FP Async_PipeItemMtd
}
func Async_PipeItem2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Async_PipeItem).FP
}
type Async_PipeItemDownCast interface {
    ToAsync_PipeItem() *Async_PipeItem
}
func Async_PipeItemDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Async_PipeItemDownCast)
    if ok { return work.ToAsync_PipeItem() }
    return nil
}
func (obj *Async_PipeItem) ToAsync_PipeItem() *Async_PipeItem {
    return obj
}
func NewAsync_PipeItem(_env *LnsEnv, arg1 LnsAny) *Async_PipeItem {
    obj := &Async_PipeItem{}
    obj.FP = obj
    obj.InitAsync_PipeItem(_env, arg1)
    return obj
}
func (self *Async_PipeItem) InitAsync_PipeItem(_env *LnsEnv, arg1 LnsAny) {
    self.item = arg1
}
func (self *Async_PipeItem) Get_item(_env *LnsEnv) LnsAny{ return self.item }

// declaration Class -- Pipe
type Async_PipeMtd interface {
    Access(_env *LnsEnv) LnsAny
    GetNext(_env *LnsEnv) LnsAny
    Run(_env *LnsEnv)
    Setup(_env *LnsEnv)
    Start(_env *LnsEnv)
    Stop(_env *LnsEnv)
}
type Async_Pipe struct {
    pipe LnsAny
    started bool
    setuped bool
    FP Async_PipeMtd
}
func Async_Pipe2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Async_Pipe).FP
}
type Async_PipeDownCast interface {
    ToAsync_Pipe() *Async_Pipe
}
func Async_PipeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Async_PipeDownCast)
    if ok { return work.ToAsync_Pipe() }
    return nil
}
func (obj *Async_Pipe) ToAsync_Pipe() *Async_Pipe {
    return obj
}
// 41: DeclConstr
func (self *Async_Pipe) InitAsync_Pipe(_env *LnsEnv, pipe LnsAny) {
    self.pipe = pipe
    self.started = false
    self.setuped = false
}




// declaration Class -- RunnerBase
type Async_RunnerBaseMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    Get_artifact(_env *LnsEnv) LnsAny
    Get_ranFlag(_env *LnsEnv) bool
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv) LnsAny
}
type Async_RunnerBase struct {
    _syncFlag *Lns_syncFlag
    pipe LnsAny
    artifact LnsAny
    ranFlag bool
    FP Async_RunnerBaseMtd
}
func (self *Async_RunnerBase) GetLnsSyncFlag() *Lns_syncFlag { return self._syncFlag }
func Async_RunnerBase2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Async_RunnerBase).FP
}
type Async_RunnerBaseDownCast interface {
    ToAsync_RunnerBase() *Async_RunnerBase
}
func Async_RunnerBaseDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Async_RunnerBaseDownCast)
    if ok { return work.ToAsync_RunnerBase() }
    return nil
}
func (obj *Async_RunnerBase) ToAsync_RunnerBase() *Async_RunnerBase {
    return obj
}
func (self *Async_RunnerBase) Get_artifact(_env *LnsEnv) LnsAny{ return self.artifact }
func (self *Async_RunnerBase) Get_ranFlag(_env *LnsEnv) bool{ return self.ranFlag }
// 111: DeclConstr
func (self *Async_RunnerBase) InitAsync_RunnerBase(_env *LnsEnv, pipe LnsAny) {
    self._syncFlag = &Lns_syncFlag{}
    self.pipe = pipe
    self.artifact = nil
    self.ranFlag = false
}



// declaration Class -- Waiter
type Async_WaiterMtd interface {
    Get_pipe(_env *LnsEnv) LnsAny
    startRunner(_env *LnsEnv, arg1 *Async_RunnerBase, arg2 LnsInt, arg3 string) bool
    Wait(_env *LnsEnv, arg1 Async_EndProcessFunc_3_)
}
type Async_Waiter struct {
    pipe LnsAny
    asyncNum LnsInt
    finRunnerList *LnsList
    FP Async_WaiterMtd
}
func Async_Waiter2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Async_Waiter).FP
}
type Async_WaiterDownCast interface {
    ToAsync_Waiter() *Async_Waiter
}
func Async_WaiterDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Async_WaiterDownCast)
    if ok { return work.ToAsync_Waiter() }
    return nil
}
func (obj *Async_Waiter) ToAsync_Waiter() *Async_Waiter {
    return obj
}
func NewAsync_Waiter(_env *LnsEnv, arg1 LnsInt) *Async_Waiter {
    obj := &Async_Waiter{}
    obj.FP = obj
    obj.InitAsync_Waiter(_env, arg1)
    return obj
}
func (self *Async_Waiter) Get_pipe(_env *LnsEnv) LnsAny{ return self.pipe }
// 153: DeclConstr
func (self *Async_Waiter) InitAsync_Waiter(_env *LnsEnv, pipeItemCount LnsInt) {
    self.pipe = LnsAny(NewLnspipe( pipeItemCount))
    self.asyncNum = 0
    self.finRunnerList = NewLnsList([]LnsAny{})
}


func Lns_Async_init(_env *LnsEnv) {
    if init_Async { return }
    init_Async = true
    Async__mod__ = "@lune.@base.@Async"
    Lns_InitMod()
}
func init() {
    init_Async = false
}
