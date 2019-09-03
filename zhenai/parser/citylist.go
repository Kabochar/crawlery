package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// 解析城市列表
func ParseCityList(bytes []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	// submatch 是 [][][]byte 类型数据
	// 第一个[]表示匹配到多少条数据，第二个[]表示匹配的数据中要提取的任容
	submatch := re.FindAllSubmatch(bytes, -1)
	result := engine.ParseResult{}
	limit := 3
	for _, item := range submatch {
		result.Items = append(result.Items, "City:"+string(item[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(item[1]), // 每一个城市对应的URL
			ParseFunc: ParseCity,       // 使用城市解析器
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
