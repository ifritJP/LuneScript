// This code is transcompiled by LuneScript.
package lnsc
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_Runner bool
var Runner__mod__ string
// declaration Class -- Runner
type Runner_RunnerMtd interface {
    GetLnsSyncFlag() *Lns_syncFlag
    Run(_env *LnsEnv)
    RunMain(_env *LnsEnv)
    Start(_env *LnsEnv, arg1 LnsInt, arg2 LnsAny)
}
type Runner_Runner struct {
    _syncFlag *Lns_syncFlag
    FP Runner_RunnerMtd
}
func (self *Runner_Runner) GetLnsSyncFlag() *Lns_syncFlag { return self._syncFlag }
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
// 29: DeclConstr
func (self *Runner_Runner) InitRunner_Runner(_env *LnsEnv) {
    self._syncFlag = &Lns_syncFlag{}
}


// 34: decl @lune.@base.@Runner.Runner.run
func (self *Runner_Runner) Run(_env *LnsEnv) {
    self.FP.RunMain(_env)
}

// 39: decl @lune.@base.@Runner.Runner.start
func (self *Runner_Runner) Start(_env *LnsEnv, mode LnsInt,name LnsAny) {
    LnsRun(_env, self.FP, mode, name)
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
