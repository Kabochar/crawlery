# Crawler

### V0.1：单任务版

流程说明：

1，配置种子请求，项目的初始入口。

2，把初始入口信息发送给爬虫引擎，引擎把其作为任务信息放入任务队列，只要任务队列不空就一直从任务队列中取任务。

3，取出任务后，engine 把要请求的任务 交给 Fetcher 模块，Fetcher 模块负责抓取 URL 页面数据，然后把数据返回给 Engine。

4，Engine收到网页数后，把数据交给解析  Parser  模块，Parser  解析出需要的数据后返回给  Engine，Engine  收到解析出的信息在控制台打印出来。

请求种子：Seed 模块

中间分发：Engine 模块

获取 URL 数据：Fetcher 模块

数据提取：Parse 模块

任务队列：requests []Request

