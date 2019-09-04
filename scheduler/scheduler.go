package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// 为每一个Request创建goroutine
	go func() {
		s.workerChan <- request
	}()
}

// 把初始请求发送给 Scheduler
func (s *SimpleScheduler) ConfigMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}