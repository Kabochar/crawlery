package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{ // 配置爬虫引擎
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{        // 配置爬虫目标信息
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
