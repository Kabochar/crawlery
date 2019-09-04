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
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)            // scheduler的输入
	out := make(chan ParseResult)    // worker的输出
	e.Scheduler.ConfigMasterWorkerChan(in)    // 把初始请求提交给scheduler

	// 创建 goroutine
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
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
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
