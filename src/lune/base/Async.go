// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Async bool
var Async__mod__ string
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
    Loop(_env *LnsEnv)
    Start(_env *LnsEnv)
}
type Async_Pipe struct {
    LnsThread
    pipe LnsAny
    started bool
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
// 40: DeclConstr
func (self *Async_Pipe) InitAsync_Pipe(_env *LnsEnv, pipe LnsAny) {
    self.InitLnsThread(_env)
    self.pipe = pipe
    
    self.started = false
    
}


// 48: decl @lune.@base.@Async.Pipe.getNext
func (self *Async_Pipe) GetNext(_env *LnsEnv) LnsAny {
    if self.started{
        {
            _pipe := self.pipe
            if !Lns_IsNil( _pipe ) {
                pipe := _pipe.(*Lns__pipe)
                var val LnsAny
                
                {
                    _val := pipe.FP.Get(_env)
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

// 60: decl @lune.@base.@Async.Pipe.loop
func (self *Async_Pipe) Loop(_env *LnsEnv) {
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
                pipe.FP.Put(_env, nil)
                break
            } else {
                val = _val.(*Async_PipeItem)
            }
        }
        pipe.FP.Put(_env, val.FP.Get_item(_env))
    }
}

// 72: decl @lune.@base.@Async.Pipe.start
func (self *Async_Pipe) Start(_env *LnsEnv) {
    self.started = true
    
    go self.LoopMain()
}


// declaration Class -- WritePipe
type Async_WritePipeMtd interface {
    Loop(_env *LnsEnv)
    Process(_env *LnsEnv, arg1 LnsAny)
    Put(_env *LnsEnv, arg1 LnsAny)
    Start(_env *LnsEnv)
}
type Async_WritePipe struct {
    LnsThread
    pipe LnsAny
    started bool
    FP Async_WritePipeMtd
}
func Async_WritePipe2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Async_WritePipe).FP
}
type Async_WritePipeDownCast interface {
    ToAsync_WritePipe() *Async_WritePipe
}
func Async_WritePipeDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Async_WritePipeDownCast)
    if ok { return work.ToAsync_WritePipe() }
    return nil
}
func (obj *Async_WritePipe) ToAsync_WritePipe() *Async_WritePipe {
    return obj
}
// 85: DeclConstr
func (self *Async_WritePipe) InitAsync_WritePipe(_env *LnsEnv, pipe LnsAny) {
    self.InitLnsThread(_env)
    self.pipe = pipe
    
    self.started = false
    
}


// 97: decl @lune.@base.@Async.WritePipe.put
func (self *Async_WritePipe) Put(_env *LnsEnv, _item LnsAny) {
    item := _item
    if self.started{
        {
            _pipe := self.pipe
            if !Lns_IsNil( _pipe ) {
                pipe := _pipe.(*Lns__pipe)
                pipe.FP.Put(_env, item)
            }
        }
    } else { 
        panic("not started")
    }
    self.FP.Process(_env, item)
}

// 108: decl @lune.@base.@Async.WritePipe.loop
func (self *Async_WritePipe) Loop(_env *LnsEnv) {
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
        var item LnsAny
        item = pipe.FP.Get(_env)
        self.FP.Process(_env, item)
        if Lns_op_not(item){
            break
        }
    }
}

// 120: decl @lune.@base.@Async.WritePipe.start
func (self *Async_WritePipe) Start(_env *LnsEnv) {
    self.started = true
    
    go self.LoopMain()
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
