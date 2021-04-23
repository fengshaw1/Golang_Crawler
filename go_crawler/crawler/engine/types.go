package engine

// 请求
type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

// 转化后的结果
type ParseResult struct {
	Requests [] Request
	Items []interface{}
}

// 空转换、测试
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
