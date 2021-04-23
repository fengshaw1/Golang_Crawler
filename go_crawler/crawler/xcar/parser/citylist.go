package parser

import (
	"go_crawler/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(//[a-z]+.xcar.com.cn+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult {
	var url string
	re := regexp.MustCompile(cityListRe)		// 挖出所有的超链接和城市名
	matches := re.FindAllSubmatch(contents, -1)	// -1代表所有的匹配
	result := engine.ParseResult{}
	for _, m := range matches {
		url = "https:" + string(m[1])
		result.Items = append(result.Items, "City " + string(m[2]))	// 将城市包装进result中
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: ParseCity,
		})
	}
	return result
}
