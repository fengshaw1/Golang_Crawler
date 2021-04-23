package main

import (
	"go_crawler/crawler/engine"
	"go_crawler/crawler/xcar/parser"
)

func main()  {
	engine.Run(engine.Request{
		Url: "https://www.xcar.com.cn/",
		ParserFunc: parser.ParseCityList,
	})
}