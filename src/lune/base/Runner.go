// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Runner bool
var Runner__mod__ string
// declaration Class -- RunItem
type Runner_RunItemMtd interface {
    ToMap() *LnsMap
    Get_result(_env *LnsEnv) bool
}
type Runner_RunItem struct {
    result bool
    FP Runner_RunItemMtd
}
func Runner_RunItem2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Runner_RunItem).FP
}
type Runner_RunItemDownCast interface {
    ToRunner_RunItem() *Runner_RunItem
}
func Runner_RunItemDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Runner_RunItemDownCast)
    if ok { return work.ToRunner_RunItem() }
    return nil
}
func (obj *Runner_RunItem) ToRunner_RunItem() *Runner_RunItem {
    return obj
}
func NewRunner_RunItem(_env *LnsEnv, arg1 bool) *Runner_RunItem {
    obj := &Runner_RunItem{}
    obj.FP = obj
    obj.InitRunner_RunItem(_env, arg1)
    return obj
}
func (self *Runner_RunItem) InitRunner_RunItem(_env *LnsEnv, arg1 bool) {
    self.result = arg1
}
func (self *Runner_RunItem) Get_result(_env *LnsEnv) bool{ return self.result }
func (self *Runner_RunItem) ToMapSetup( obj *LnsMap ) *LnsMap {
    obj.Items["result"] = Lns_ToCollection( self.result )
    return obj
}
func (self *Runner_RunItem) ToMap() *LnsMap {
    return self.ToMapSetup( NewLnsMap( map[LnsAny]LnsAny{} ) )
}
func Runner_RunItem__fromMap_1022_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Runner_RunItem_FromMap( arg1, paramList )
}
func Runner_RunItem__fromStem_1026_(_env,  arg1 LnsAny, paramList []Lns_ToObjParam)(LnsAny, LnsAny){
   return Runner_RunItem_FromMap( arg1, paramList )
}
func Runner_RunItem_FromMap( obj LnsAny, paramList []Lns_ToObjParam ) (LnsAny, LnsAny) {
    _,conv,mess := Runner_RunItem_FromMapSub(obj,false, paramList);
    return conv,mess
}
func Runner_RunItem_FromMapSub( obj LnsAny, nilable bool, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    var objMap *LnsMap
    if work, ok := obj.(*LnsMap); !ok {
       return false, nil, "no map -- " + Lns_ToString(obj)
    } else {
       objMap = work
    }
    newObj := &Runner_RunItem{}
    newObj.FP = newObj
    return Runner_RunItem_FromMapMain( newObj, objMap, paramList )
}
func Runner_RunItem_FromMapMain( newObj *Runner_RunItem, objMap *LnsMap, paramList []Lns_ToObjParam ) (bool, LnsAny, LnsAny) {
    if ok,conv,mess := Lns_ToBoolSub( objMap.Items["result"], false, nil); !ok {
       return false,nil,"result:" + mess.(string)
    } else {
       newObj.result = conv.(bool)
    }
    return true, newObj, nil
}
func Runner_RunItem__createPipe_1034_(_env *LnsEnv, arg1 LnsInt) LnsAny{
   return NewLnspipe( arg1 )
}

// declaration Class -- Runner
type Runner_RunnerMtd interface {
    GetResult(_env *LnsEnv) bool
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    Start(_env *LnsEnv)
}
type Runner_Runner struct {
    pipe LnsAny
    FP Runner_RunnerMtd
}
func Runner_Runner2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Runner_Runner).FP
}
type Runner_RunnerDownCast interface {
    ToRunner_Runner() *Runner_Runner
}
func Runner_RunnerDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Runner_RunnerDownCast)
    if ok { return work.ToRunner_Runner() }
    return nil
}
func (obj *Runner_Runner) ToRunner_Runner() *Runner_Runner {
    return obj
}
// 35: DeclConstr
func (self *Runner_Runner) InitRunner_Runner(_env *LnsEnv) {
    self.pipe = Runner_RunItem__createPipe_1034_(_env, 1)
    
}


// 41: decl @lune.@base.@Runner.Runner.run
func (self *Runner_Runner) Run(_env *LnsEnv) {
    self.FP.RunMain(_env)
    {
        __exp := self.pipe
        if !Lns_IsNil( __exp ) {
            _exp := __exp.(*Lns__pipe)
            _exp.FP.Put(_env, Runner_RunItem2Stem(NewRunner_RunItem(_env, true)))
        }
    }
}

// 50: decl @lune.@base.@Runner.Runner.start
func (self *Runner_Runner) Start(_env *LnsEnv) {
    LnsRun( self )
}

// 53: decl @lune.@base.@Runner.Runner.getResult
func (self *Runner_Runner) GetResult(_env *LnsEnv) bool {
    {
        _pipe := self.pipe
        if !Lns_IsNil( _pipe ) {
            pipe := _pipe.(*Lns__pipe)
            return Lns_unwrapDefault( _env.NilAccFin(_env.NilAccPush(pipe.FP.Get(_env)) && 
            Lns_NilAccCall1( _env, func () LnsAny { return _env.NilAccPop().(*Runner_RunItem).FP.Get_result(_env)})), false).(bool)
        }
    }
    return false
}


func Lns_Runner_init(_env *LnsEnv) {
    if init_Runner { return }
    init_Runner = true
    Runner__mod__ = "@lune.@base.@Runner"
    Lns_InitMod()
}
func init() {
    init_Runner = false
}
