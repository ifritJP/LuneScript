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
	"time"
)

var lns_thread_event_log_on = false

const (
	_THREAD_EVENT_REQ              = 0
	_THREAD_EVENT_START            = 1
	_THREAD_EVENT_END              = 2
	_THREAD_EVENT_LOG              = 3
	_THREAD_EVENT_QUEUE_START      = 4
	_THREAD_EVENT_START_FROM_SUITE = 5
)

type lnsThreadEvent struct {
	// イベント発生時間
	time time.Time
	// runner ID
	runnerId int
	// runner の名前
	runnerName string
	// イベント
	kind int
	// イベント情報。
	// eventKind が THREAD_EVENT_REQ の時有効。 _RUN_KIND_* の値が入る。
	info1 int
	// eventKind が THREAD_EVENT_REQ の時有効。 起動元 runnerId が入る。
	info2 int
	// イベント発生時の runNum
	runNum int
	// メッセージ
	mess string
}

type lnsRunnerInfo struct {
	runner LnsRunner
	name   string
	id     int
}

type lnsRunSuite struct {
	launched bool
	runnerCh chan *lnsRunnerInfo
}

/**
runnerCh から runnerInfo を取得し、その runner を実行する。
*/
func (self *lnsRunSuite) runnerLoop(threadMgrInfo *Lns_ThreadMgrInfo) {
	self.launched = true
	runnerInfo := <-self.runnerCh
	for runnerInfo != nil {
		lnsRunMain(runnerInfo, threadMgrInfo)
		runnerInfo = threadMgrInfo.endToRun(runnerInfo, self)
		if runnerInfo == nil {
			runnerInfo = <-self.runnerCh
		}
	}
}

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
	// 実行中の数 (aliveNum から Wait 中を除く)
	runNum int
	// limitNum を越えてリクエストされた runner の &lnsRunnerInfo queue
	runQueue *list.List
	// lnsThreadEvent を格納する list
	eventList *list.List

	// 待ち状態の lnsRunSuite リスト
	freeFreeRunSuiteList *list.List
}

func (self *Lns_ThreadMgrInfo) lock() {
	self.mutex.Lock()
}
func (self *Lns_ThreadMgrInfo) unlock() {
	self.mutex.Unlock()
}

func (self *Lns_ThreadMgrInfo) createAsyncEnv(runnerName string) *LnsEnv {
	self.lock()
	defer self.unlock()

	self.totalReqNum++

	return createEnv(true, runnerName, self.totalReqNum)
}

func (self *Lns_ThreadMgrInfo) newThreadEvent(
	runnerId int, runnerName string,
	eventId int, info1 int, info2 int, mess string) *lnsThreadEvent {
	return &lnsThreadEvent{
		time.Now(), runnerId, runnerName, eventId, info1, info2, self.runNum, mess}
}

func (self *Lns_ThreadMgrInfo) newEvent(
	runnerInfo *lnsRunnerInfo, eventId int, info1 int, info2 int) *lnsThreadEvent {
	return self.newThreadEvent(runnerInfo.id, runnerInfo.name, eventId, info1, info2, "")
}

const (
	// 非同期で実行
	_RUN_KIND_ASYNC = 0
	// Queue に登録
	_RUN_KIND_QUEUE = 1
	// 同期で実行
	_RUN_KIND_SYNC = 2
	// 実行しない
	_RUN_KIND_SKIP = 3
)

const (
	// limitNum オーバー時、 スレッドを新規起動せずに、呼び出し元スレッドで動かす
	_RUN_MODE_ON_FULL_SYNC = 0
	// limitNum オーバー時、 キューに入れ、スレッドが空いた時に動かす
	_RUN_MODE_ON_FULL_QUEUE = 1
	// limitNum オーバー時、 処理を動かさない
	_RUN_MODE_ON_FULL_SKIP = 2
	// limitNum オーバーを無視する
	_RUN_MODE_ON_FULL_IGNORE = 3
)

func (self *Lns_ThreadMgrInfo) log(_env *LnsEnv, eventId int, mess string) {
	if lns_thread_event_log_on {
		self.lock()
		defer self.unlock()
		self.eventList.PushBack(self.newThreadEvent(
			_env.runnerId, _env.runnerName, eventId, 0, 0, mess))

	}
}

/**
  runner を実行する。

  @param runner 実行する処理
  @param mode モード _RUN_MODE
  @return 実行した場合 true
*/
func (self *Lns_ThreadMgrInfo) run(
	runner LnsRunner, mode int, _env *LnsEnv, name string) bool {

	// mode から、条件にあう _RUN_KIND を判断して返す
	//
	// @return int _RUN_KIND
	// @return lnsRunnerInfo Runner 管理情報
	// @return lnsRunSuite runner を動かす lnsRunSuite。
	//   runner を直ぐに非同期で動かさない場合は nil。
	process := func() (int, *lnsRunnerInfo, *lnsRunSuite) {
		self.lock()
		defer self.unlock()

		self.totalReqNum++
		self.threadNum++
		if self.peakNum < self.threadNum {
			self.peakNum = self.threadNum
		}
		var runnerName string
		if name == "" {
			runnerName = fmt.Sprintf("%d", self.totalReqNum)
		} else {
			runnerName = name
		}
		runnerInfo := &lnsRunnerInfo{runner, runnerName, self.totalReqNum}

		kind := _RUN_KIND_ASYNC
		if self.aliveNum >= self.limitNum && mode != _RUN_MODE_ON_FULL_IGNORE {
			if mode == _RUN_MODE_ON_FULL_SYNC {
				kind = _RUN_KIND_SYNC
			} else if mode == _RUN_MODE_ON_FULL_QUEUE {
				kind = _RUN_KIND_QUEUE
			} else if mode == _RUN_MODE_ON_FULL_SKIP {
				kind = _RUN_KIND_SKIP
			} else {
				panic(fmt.Sprintf("illegal async run mode -- %d", mode))
			}
		} else {
			self.aliveNum++
			self.runNum++
		}

		if lns_thread_event_log_on {
			self.eventList.PushBack(
				self.newEvent(runnerInfo, _THREAD_EVENT_REQ, kind, _env.runnerId))
			if kind == _RUN_KIND_SYNC {
				self.eventList.PushBack(
					self.newEvent(runnerInfo, _THREAD_EVENT_START, 0, 0))
			}
		}

		var runSuite *lnsRunSuite = nil
		if kind == _RUN_KIND_ASYNC {
			if self.freeFreeRunSuiteList.Len() == 0 {
				// 空がなければ新規に作成
				runSuite = &lnsRunSuite{}
				runSuite.runnerCh = make(chan *lnsRunnerInfo, 1)
				runSuite.launched = false
			} else {
				// 空きを利用する
				element := self.freeFreeRunSuiteList.Back()
				self.freeFreeRunSuiteList.Remove(element)
				runSuite = element.Value.(*lnsRunSuite)

				if lns_thread_event_log_on {
					self.eventList.PushBack(
						self.newEvent(runnerInfo, _THREAD_EVENT_START_FROM_SUITE, 0, 0))
				}
			}
		}

		return kind, runnerInfo, runSuite
	}

	result := true
	switch kind, runnerInfo, runSuite := process(); kind {
	case _RUN_KIND_ASYNC:
		// 非同期で直ぐに動かす
		runner.GetLnsSyncFlag().wg.Add(1)
		runSuite.runnerCh <- runnerInfo
		if !runSuite.launched {
			go runSuite.runnerLoop(self)
		}
	case _RUN_KIND_QUEUE:
		// Queue に詰む
		runner.GetLnsSyncFlag().wg.Add(1)
		self.runQueue.PushBack(runnerInfo)
	case _RUN_KIND_SYNC:
		// 同期で動かす
		self.threadNum--
		runner.GetLnsSyncFlag().wg.Add(1)

		LnsExecRunner(_env, runner)

		if lns_thread_event_log_on {
			self.lock()
			self.eventList.PushBack(
				self.newEvent(runnerInfo, _THREAD_EVENT_END, 0, 0))
			self.unlock()
		}

		runner.GetLnsSyncFlag().wg.Done()
	case _RUN_KIND_SKIP:
		// 動かさない
		self.threadNum--
		result = false
	default:
		// エラー
		panic(fmt.Sprintf("illegal run_kind -- %d", kind))
	}
	return result
}

/**
非同期起動した runner の終了処理。

runner が runQueue にある場合、先頭の runner を取り出して起動する。
*/
func (self *Lns_ThreadMgrInfo) endToRun(info *lnsRunnerInfo, runSuite *lnsRunSuite) *lnsRunnerInfo {
	self.lock()
	defer self.unlock()

	if lns_thread_event_log_on {
		self.eventList.PushBack(
			self.newEvent(info, _THREAD_EVENT_END, 0, 0))
	}

	self.threadNum--
	self.aliveNum--
	self.runNum--

	if self.runQueue.Len() > 0 {
		self.aliveNum++
		self.runNum++
		front := self.runQueue.Front()

		nextRunnerInfo := self.runQueue.Remove(front).(*lnsRunnerInfo)

		if lns_thread_event_log_on {
			self.eventList.PushBack(
				self.newEvent(nextRunnerInfo, _THREAD_EVENT_QUEUE_START, info.id, 0))
		}

		return nextRunnerInfo
	}

	// 空きに追加
	self.freeFreeRunSuiteList.PushBack(runSuite)
	return nil
}

func (self *Lns_ThreadMgrInfo) dumpEventLog(write func(txt string)) {
	self.lock()
	defer self.unlock()

	element := self.eventList.Front()
	for element != nil {
		event := element.Value.(*lnsThreadEvent)
		element = element.Next()

		delta := event.time.Sub(lns_profStartTime)
		write(fmt.Sprintf(
			"%5d:%3d:%d:%d:%3d:%3d:%s:%s", delta.Milliseconds(),
			event.runnerId, event.kind, event.info1, event.info2,
			event.runNum, event.runnerName, event.mess))
		write("\n")

	}
}

func (self *Lns_ThreadMgrInfo) SetLimit(limit int) {
	self.limitNum = limit
}
func Lns_setThreadLimit(limit int) {
	lns_threadMgrInfo.SetLimit(limit)
}

func (self *Lns_ThreadMgrInfo) dump(writer io.Writer) {
	writer.Write([]byte(fmt.Sprintf("--------- thread ---------\n")))
	writer.Write([]byte(fmt.Sprintf("totalReqNum = %d\n", self.totalReqNum)))
	writer.Write([]byte(fmt.Sprintf("threadNum = %d\n", self.threadNum)))
	writer.Write([]byte(fmt.Sprintf("peakNum = %d\n", self.peakNum)))
	writer.Write([]byte(fmt.Sprintf(
		"freeFreeRunSuiteList = %d\n", self.freeFreeRunSuiteList.Len())))
}

var lns_threadMgrInfo *Lns_ThreadMgrInfo

func lns_newThreadMgrInfo() *Lns_ThreadMgrInfo {
	threadInfo := &Lns_ThreadMgrInfo{}
	threadInfo.runQueue = list.New()
	threadInfo.eventList = list.New()
	threadInfo.freeFreeRunSuiteList = list.New()
	threadInfo.limitNum = 200
	return threadInfo
}

var lns_profStartTime time.Time

func init() {
	lns_threadMgrInfo = lns_newThreadMgrInfo()
	lns_profStartTime = time.Now()
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

func (self *Lns_ThreadMgrInfo) changeRunNum(delta int) {
	if lns_thread_event_log_on {
		self.lock()
		defer self.unlock()

		self.runNum += delta
	}
}

func Lns_LockEnvSync(_env *LnsEnv, lineNo int, callback func()) {
	if _env.async {
		// __noasync が待ちになるまで待つために lock する
		var prev time.Time
		if lns_thread_event_log_on {
			prev = time.Now()
		}
		sync_LnsEnvMutex.Lock()
		_env.async = false

		if lns_thread_event_log_on {
			delta := time.Now().Sub(prev).Milliseconds()

			if delta > 1 {
				lns_threadMgrInfo.log(_env, _THREAD_EVENT_LOG,
					fmt.Sprintf("wait -- %d from %d", delta, lineNo))
			}
		}

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

type G__lns_Flag interface {
	Set(_env *LnsEnv)
	Wait(_env *LnsEnv)
}

type Lns_syncFlag struct {
	wg sync.WaitGroup
}

func LnsCreateSyncFlag(_env *LnsEnv) G__lns_Flag {
	flag := &Lns_syncFlag{}
	flag.wg.Add(1)
	return flag
}
func (self *Lns_syncFlag) Set(_env *LnsEnv) {
	self.wg.Done()
}
