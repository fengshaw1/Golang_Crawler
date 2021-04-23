package parser

import (
	"go_crawler/crawler/engine"
	"regexp"
	"strings"
)
// 				<a href="//dealer.xcar.com.cn/34897/price_m50767.htm"
const cityRe = `<a[^h]*href="(//dealer.xcar.com.cn/[0-9]*/price_[a-z0-9]*.htm)"[^>]*>([^<]*)</a>`
func ParseCity(contents []byte) engine.ParseResult {
	var url string
	var carname string
	re := regexp.MustCompile(cityRe)		// 挖出所有的超链接和城市名
	matches := re.FindAllSubmatch(contents, -1)	// -1代表所有的匹配
	result := engine.ParseResult{}
	for _, m := range matches {
		url = "https:" + string(m[1])
		carname = string(m[2])
		carname = strings.Replace(carname," ","", -1)
		result.Items = append(result.Items, "Car " + carname)	// 将城市包装进result中
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
