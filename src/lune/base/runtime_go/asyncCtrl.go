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

import (
	"container/list"
	"fmt"
	"io"
	"sync"
)

type Lns_ThreadMgrInfo struct {
	// この構造体への排他
	mutex sync.Mutex
	// 起動リクエストされた回数
	totalReqNum int
	// runQueue に入っているリクエストを含む実行中の runner の数
	threadNum int
	// threadNum の最大値
	peakNum int
	// runQueue を除く実行中の runner の数
	aliveNum int
	// aliveNum の上限
	limitNum int
	// limitNum を越えてリクエストされた runner
	runQueue *list.List
}

func (self *Lns_ThreadMgrInfo) lock() {
	self.mutex.Lock()
}
func (self *Lns_ThreadMgrInfo) unlock() {
	self.mutex.Unlock()
}

const (
	RUN_KIND_ASYNC = 0
	RUN_KIND_QUEUE = 1
	RUN_KIND_SYNC  = 2
	RUN_KIND_SKIP  = 3
)

const (
	// limitNum オーバー時、 スレッドを新規起動せずに、呼び出し元スレッドで動かす
	RUN_MODE_ON_FULL_SYNC = 0
	// limitNum オーバー時、 キューに入れ、スレッドが空いた時に動かす
	RUN_MODE_ON_FULL_QUEUE = 1
	// limitNum オーバー時、 処理を動かさない
	RUN_MODE_ON_FULL_SKIP = 2
)

/**
  runner を実行する。

  @param runner 実行する処理
  @param mode モード
  @return 実行した場合 true
*/
func (self *Lns_ThreadMgrInfo) run(runner LnsRunner, mode int, _env *LnsEnv) bool {

	process := func() int {
		self.lock()
		defer self.unlock()

		self.totalReqNum++
		self.threadNum++
		if self.peakNum < self.threadNum {
			self.peakNum = self.threadNum
		}
		kind := RUN_KIND_ASYNC
		if self.aliveNum >= self.limitNum {
			if mode == 0 {
				kind = RUN_KIND_SYNC
			} else if mode == 1 {
				kind = RUN_KIND_QUEUE
			} else if mode == 2 {
				kind = RUN_KIND_SKIP
			} else {
				panic(fmt.Sprintf("illegal async run mode -- %d", mode))
			}
		} else {
			self.aliveNum++
		}
		return kind
	}

	result := true
	switch kind := process(); kind {
	case RUN_KIND_ASYNC:
		runner.GetLnsSyncFlag().wg.Add(1)
		go lnsRunMain(runner, self)
	case RUN_KIND_QUEUE:
		runner.GetLnsSyncFlag().wg.Add(1)
		self.runQueue.PushBack(runner)
	case RUN_KIND_SYNC:
		self.threadNum--
		runner.GetLnsSyncFlag().wg.Add(1)
		LnsExecRunner(_env, runner)
		runner.GetLnsSyncFlag().wg.Done()
	case RUN_KIND_SKIP:
		self.threadNum--
		result = false
	default:
		panic(fmt.Sprintf("illegal run_kind -- %d", kind))
	}
	return result
}

/**
非同期起動した runner の終了処理。

runner が runQueue にある場合、先頭の runner を取り出して起動する。
*/
func (self *Lns_ThreadMgrInfo) endToRun(runner LnsRunner) {
	process := func() LnsRunner {
		self.lock()
		defer self.unlock()

		self.threadNum--
		self.aliveNum--

		if self.runQueue.Len() > 0 {
			self.aliveNum++
			front := self.runQueue.Front()
			return self.runQueue.Remove(front).(LnsRunner)
		}
		return nil
	}

	nextRunner := process()
	if nextRunner != nil {
		go lnsRunMain(nextRunner, self)
	}
}
func (self *Lns_ThreadMgrInfo) SetLimit(limit int) {
	self.limitNum = limit
}
func Lns_setThreadLimit(limit int) {
	lns_threadMgrInfo.SetLimit(limit)
}

func (self *Lns_ThreadMgrInfo) dump(writer io.Writer) {
	writer.Write([]byte(fmt.Sprintf("totalReqNum = %d\n", self.totalReqNum)))
	writer.Write([]byte(fmt.Sprintf("threadNum = %d\n", self.threadNum)))
	writer.Write([]byte(fmt.Sprintf("peakNum = %d\n", self.peakNum)))
}

var lns_threadMgrInfo *Lns_ThreadMgrInfo

func lns_newThreadMgrInfo() *Lns_ThreadMgrInfo {
	threadInfo := &Lns_ThreadMgrInfo{}
	threadInfo.runQueue = list.New()
	threadInfo.limitNum = 200
	return threadInfo
}

func init() {
	lns_threadMgrInfo = lns_newThreadMgrInfo()
}

type Lns__pipe struct {
	ch  chan LnsAny
	end bool
	FP  Lns_pipeMtd
}

func NewLnspipe(count int) *Lns__pipe {
	pipe := &Lns__pipe{}
	pipe.ch = make(chan LnsAny, count)
	pipe.FP = pipe
	pipe.end = false
	return pipe
}

func (self *Lns__pipe) put(val LnsAny) {
	if Lns_IsNil(val) {
		self.end = true
	}
	self.ch <- val
}
func (self *Lns__pipe) get() LnsAny {
	if len(self.ch) == 0 && self.end {
		return nil
	}
	return <-self.ch
}

func Lns_LockEnvSync(_env *LnsEnv, callback func()) {
	if _env.async {
		// __noasync が待ちになるまで待つために lock する
		sync_LnsEnvMutex.Lock()
		_env.async = false

		// 処理終了後に lock を開放するために defer する。
		defer func() {
			_env.async = true
			sync_LnsEnvMutex.Unlock()
		}()

		callback()

	} else {
		callback()
	}
}

type Lns_syncFlag struct {
	wg sync.WaitGroup
}
