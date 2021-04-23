package engine

import (
	"go_crawler/crawler/fetcher"
	"log"
)

// 引擎、对于源源不断的请求、反复的进行获取和赋值操作、循环往复
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fet111ching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error" + "fetching url %s : %v", r.Url, err)
			continue
		}

		ParseResult := r.ParserFunc(body)
		requests = append(requests, ParseResult.Requests...)

		for _, item := range ParseResult.Items {
			log.Printf("Got Item %v", item)
		}
	}
}