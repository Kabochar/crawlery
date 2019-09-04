package scheduler

import "crawler/engine"

// 单任务版 调度

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	// 此时所有 worker 共用同一个 channel，直接返回即可
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(w chan engine.Request) {
	// todo
}

func (s *SimpleScheduler) Run() {
	// 创建出 worker channel
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// 发送请求到 worker chan
	// 为什么需要 goroutine 发送？ 避免createWorker持续等待。
	// 每个 request 进一个 goroutine
	go func() {
		s.workerChan <- request
	}()
}