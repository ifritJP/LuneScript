// +build !noEnvArg

/*
MIT License

Copyright (c) 2018,2020 ifritJP

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package runtimelns

type Lns_pipeMtd interface {
	Put(_env *LnsEnv, val LnsAny)
	Get(_env *LnsEnv) LnsAny
}

func (self *Lns__pipe) Put(_env *LnsEnv, val LnsAny) {
	self.put(val)
}
func (self *Lns__pipe) Get(_env *LnsEnv) LnsAny {
	return self.get()
}

type LnsThreadMtd interface {
	Loop(*LnsEnv)
}

func (self *LnsThread) InitLnsThread(_env *LnsEnv) {
	self.initLnsThread()
}

func (self *LnsThread) runLoop() {
	self.FP.Loop(self.LnsEnv)
}

type LnsRunner interface {
	Run(*LnsEnv)
}

func lnsRunMain(self LnsRunner) {
    env := createEnv()
    
	self.Run( env )
    
    env.LuaVM.closeVM()
}

func LnsRun(self LnsRunner) {
	go lnsRunMain( self )
}
