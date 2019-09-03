package parser

import (
	"crawler/engine"
	"regexp"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

// 用户性别正则，因为在用户详情页没有性别信息，所以在用户性别在用户列表页面获取
var sexRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)

// 城市页面用户解析器
func ParseCity(bytes []byte) engine.ParseResult {
	submatch := cityRe.FindAllSubmatch(bytes, -1)
	gendermatch := sexRe.FindAllSubmatch(bytes, -1)

	result := engine.ParseResult{}

	for k, item := range submatch {
		name := string(item[2])
		gender := string(gendermatch[k][1])

		result.Items = append(result.Items, "User:"+name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(item[1]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name, gender)
			},
		})
	}
	return result
}
