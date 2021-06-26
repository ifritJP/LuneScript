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

import "fmt"

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

func (self *Lns_syncFlag) Wait(_env *LnsEnv) {
	if !_env.async {
		// __asyncLock の処理が動くように sync_LnsEnvMutex を unlock する。
		sync_LnsEnvMutex.Unlock()

		// wait 後、sync_LnsEnvMutex を lock するために defer する。
		defer sync_LnsEnvMutex.Lock()
	} else {
		lns_threadMgrInfo.changeRunNum(-1)
		defer lns_threadMgrInfo.changeRunNum(1)
	}
	self.wg.Wait()
}

type LnsRunner interface {
	Run(*LnsEnv)
	GetLnsSyncFlag() *Lns_syncFlag
}

func lnsRunMain(runnerInfo *lnsRunnerInfo, threadMgrInfo *Lns_ThreadMgrInfo) {
	env := createEnv(true, runnerInfo.name, runnerInfo.id)

	lns_threadMgrInfo.log(env, _THREAD_EVENT_START, "")

	LnsExecRunner(env, runnerInfo.runner)

	runnerInfo.runner.GetLnsSyncFlag().wg.Done()

	env.LuaVM.closeVM()

	threadMgrInfo.endToRun(runnerInfo)
}

func LnsRun(_env *LnsEnv, runner LnsRunner, mode int, nameOp LnsAny) bool {
	var name string = ""
	if nameOp != nil {
		name = nameOp.(string)
	}
	return lns_threadMgrInfo.run(runner, mode, _env, name)
}

func LnsExecRunner(_env *LnsEnv, runner LnsRunner) {

	runner.Run(_env)

}

func LnsJoin(_env *LnsEnv, runner LnsRunner) {
	runner.GetLnsSyncFlag().Wait(_env)
}

func LnsLog(_env *LnsEnv, mess string) {
	lns_threadMgrInfo.log(_env, _THREAD_EVENT_LOG, mess)
}

func LnsStartRunnerLog(_env *LnsEnv, valid bool) {
	lns_thread_event_log_on = valid
}

func LnsDumpRunnerLog(_env *LnsEnv, stream Lns_oStream) {
	lns_threadMgrInfo.dumpEventLog(
		func(txt string) {
			stream.Write(_env, txt)
		})
}

type LnsProcessor struct {
	running bool
	// processor 終了管理の flag
	syncFlag *Lns_syncFlag
	// processor の名前
	name string
	// processor へのリクエスト queue
	reqQueue chan func(*LnsEnv)
	// processor へのリクエスト処理終了待ち queue
	respQueue chan bool
}

func LnsCreateProcessor(_env *LnsEnv, name string) *LnsProcessor {
	processor := &LnsProcessor{
		true,
		&Lns_syncFlag{},
		name,
		make(chan func(*LnsEnv), 1),
		make(chan bool, 1),
	}

	lns_threadMgrInfo.run(processor, _RUN_MODE_ON_FULL_IGNORE, _env, name)

	return processor
}

func (self *LnsProcessor) Run(_env *LnsEnv) {
	for true {
		request := <-self.reqQueue
		if Lns_IsNil(request) {
			return
		}
		request(_env)
		self.respQueue <- true
	}
}

func (self *LnsProcessor) GetLnsSyncFlag() *Lns_syncFlag {
	return self.syncFlag
}

func (self *LnsProcessor) Request(_env *LnsEnv, proc func(*LnsEnv)) {
	if !self.running {
		panic(fmt.Sprintf("Thie processer is not running -- %s", self.name))
	}
	self.reqQueue <- proc
	<-self.respQueue
}

func (self *LnsProcessor) End(_env *LnsEnv) {
	if self.running {
		self.running = false
		self.reqQueue <- nil
	}
}
