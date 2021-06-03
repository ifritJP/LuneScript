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
	if !_env.async {
		// __noasync で get する際に、
		// __asyncLock の処理が動くように sync_LnsEnvMutex を unlock する。
		sync_LnsEnvMutex.Unlock()

		// get 後、sync_LnsEnvMutex を lock するために defer する。
		defer sync_LnsEnvMutex.Lock()
	}

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
    GetLnsSyncFlag() *Lns_syncFlag
}

func lnsRunMain(self LnsRunner) {
	env := createEnv(true)

	self.Run(env)
    self.GetLnsSyncFlag().wg.Done()

    
	env.LuaVM.closeVM()
}

func LnsRun(self LnsRunner) {
    self.GetLnsSyncFlag().wg.Add(1)
	go lnsRunMain(self)
}

func (self *Lns_syncFlag) Wait( _env *LnsEnv ) {
	if !_env.async {
		// __asyncLock の処理が動くように sync_LnsEnvMutex を unlock する。
		sync_LnsEnvMutex.Unlock()

		// wait 後、sync_LnsEnvMutex を lock するために defer する。
		defer sync_LnsEnvMutex.Lock()
	}
    self.wg.Wait()    
}
