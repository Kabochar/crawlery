package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.QueuedScheduler{},    // 队列实现调度器
		Scheduler:   &scheduler.SimpleScheduler{},    // 简单并发调度
		WorkerCount: 50,
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
