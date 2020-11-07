// This code is transcompiled by LuneScript.
package lnsc
import . "lnsc/lune/base/runtime_go"
var init_Async bool
var Async__mod__ string
// declaration Class -- PipeItem
type Async_PipeItemMtd interface {
    Get_item() LnsAny
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
func NewAsync_PipeItem(arg1 LnsAny) *Async_PipeItem {
    obj := &Async_PipeItem{}
    obj.FP = obj
    obj.InitAsync_PipeItem(arg1)
    return obj
}
func (self *Async_PipeItem) InitAsync_PipeItem(arg1 LnsAny) {
    self.item = arg1
}
func (self *Async_PipeItem) Get_item() LnsAny{ return self.item }

// declaration Class -- Pipe
type Async_PipeMtd interface {
    Access() LnsAny
    GetNext() LnsAny
    Loop()
    Start()
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
// 35: DeclConstr
func (self *Async_Pipe) InitAsync_Pipe(pipe LnsAny) {
    self.InitLnsThread()
    self.pipe = pipe
    
    self.started = false
    
}


// 43: decl @lune.@base.@Async.Pipe.getNext
func (self *Async_Pipe) GetNext() LnsAny {
    if self.started{
        {
            _pipe := self.pipe
            if _pipe != nil {
                pipe := _pipe.(*Lns__pipe)
                var val LnsAny
                
                {
                    _val := pipe.FP.Get()
                    if _val == nil{
                        return nil
                    } else {
                        val = _val
                    }
                }
                return NewAsync_PipeItem(val)
            }
        }
    }
    return self.FP.Access()
}

// 55: decl @lune.@base.@Async.Pipe.loop
func (self *Async_Pipe) Loop() {
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
            _val := self.FP.Access()
            if _val == nil{
                break
            } else {
                val = _val.(*Async_PipeItem)
            }
        }
        pipe.FP.Put(Async_PipeItem2Stem(val))
    }
}

// 66: decl @lune.@base.@Async.Pipe.start
func (self *Async_Pipe) Start() {
    self.started = true
    
    go self.FP.Loop()
}


func Lns_Async_init() {
    if init_Async { return }
    init_Async = true
    Async__mod__ = "@lune.@base.@Async"
    Lns_InitMod()
}
func init() {
    init_Async = false
}
