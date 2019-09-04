package engine

import "log"

// 并发引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler    // 任务调度器
	WorkerCount int            // 任务并发数量
}

// 任务调度器
type Scheduler interface {
	ReadyNotifier
	Submit(request Request) // 提交任务
	WorkerChan() chan Request // 自动分配，每个 worker 一个 channel OR 所有 worker 共用一个 channel。
	Run()
}

// 这里避免直接传入 Scheduler 过于繁重
type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 创建 goroutine
	for i := 0; i < e.WorkerCount; i++ {
		// 任务是每个 worker 一个 channel 还是
		// 所有 worker 共用一个 channel
		// 由 WorkerChan 来决定
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
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
			if isDuplicated(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

// 创建任务，调用worker，分发goroutine
func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	// 为每一个Worker创建一个channel
	go func() {
		for {
			ready.WorkerReady(in) // 告诉调度器任务空闲

			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

// 去重判断。仅限当次运行
var visitedUrls = make(map[string]bool)

func isDuplicated(url string) bool {
	if visitedUrls[url] {
		return true
	}

	visitedUrls[url] = true

	return false
}