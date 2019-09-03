package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	// 配置请求信息
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
