package parser

import (
	"go_crawler/crawler/engine"
	"go_crawler/crawler/model"
	"regexp"
)

// 预先编译、节省时间、提高程序效率
var priceRe   = regexp.MustCompile(`<tr><td class="back_hui">厂商指导价</td><td>([^<]+)</td></tr>`)
var factoryRe = regexp.MustCompile(`<tr><td class="back_hui">厂商</td><td>([^<]+)</td></tr>`)
var levelRe   = regexp.MustCompile(`<tr><td class="back_hui">级别</td><td>([^<]+)</td></tr>`)
var motorRe   = regexp.MustCompile(`<tr><td class="back_hui">发动机</td><td>([^<]+)</td></tr>`)
var powerTypeRe = regexp.MustCompile(`<tr><td class="back_hui">动力类型</td><td>([^<]+)</td></tr>`)
var yearRe    = regexp.MustCompile(`<tr><td class="back_hui">上市年份</td><td>([^<]+)</td></tr>`)
var speedRe   = regexp.MustCompile(`<tr><td class="back_hui">最高车速(km/h)</td><td>([^<]+)</td>`)
var speedUpRe = regexp.MustCompile(`<tr><td class="back_hui">0-100加速时间(s)</td><td>([^<]+)</td></tr>`)
var oilRe     = regexp.MustCompile(`<tr><td class="back_hui">工信部油耗(L/100km)</td><td>([^<]+)</td></tr>`)
func ParseProfile(contents []byte) engine.ParseResult  {
	profile :=model.Profile{}
	//price, err := strconv.Atoi(
	//	extractString(contents, priceRe))
	// 整数部分进行校验转换、其他无需操作
	//if err != nil {
	//	profile.Price = price
	//}
	profile.Price = extractString(
		contents, priceRe)
	profile.Factory = extractString(
		contents, factoryRe)
	profile.Level = extractString(
		contents, levelRe)
	profile.Motor = extractString(
		contents, motorRe)
	profile.PowerType = extractString(
		contents, powerTypeRe)
	profile.Year = extractString(
		contents, yearRe)
	profile.Speed = extractString(
		contents, speedRe)
	profile.Acceleration = extractString(
		contents, speedUpRe)
	profile.Oil = extractString(
		contents, oilRe)
	// 将得到的result返回、值为profile
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
	// 修改前代码
	//match := priceRe.FindSubmatch(contents)
	//if match != nil {
	//	// 第一个括号即为价格
	//	price, err := strconv.Atoi(string(match[1]))
	//	if err!=nil {
	//		// 说明获得了对应的价格
	//		profile.Price = price
	//	}
	//}
	//
	//match = FactoryRe.FindSubmatch(contents)
	//
	//if match != nil {
	//	profile.Factory = string(match[1])
	//}
}

// 提取公共部分
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	// 如果得到的是字段名 + 值的形式
	if len(match) >= 2 {
		// 返回第一个（第0个是整个字符串）
		return string(match[1])
	} else {
		return ""
	}

}
