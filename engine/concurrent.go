package engine

import "log"

// 并发引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler    // 任务调度器
	WorkerCount int            // 任务并发数量
}

// 任务调度器
type Scheduler interface {
	Submit(request Request) // 提交任务
	ConfigMasterWorkerChan(chan Request)    // 配置初始请求任务
	WorkerReady(w chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 创建 goruntine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	// engine把请求任务提交给 Scheduler
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	itemCount := 0
	for {
		// 接受 Worker 的解析结果
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: #%d: %v\n", itemCount, item)
			itemCount++
		}

		// 然后把 Worker 解析出的 Request 送给 Scheduler
		// Scheduler.Submit，一个 Request 一个 worker goroutine 处理
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// 创建任务，调用worker，分发goroutine
func createWorker(out chan ParseResult, s Scheduler) {
	// 为每一个Worker创建一个channel
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in) // 告诉调度器任务空闲
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
