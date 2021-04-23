package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// 将输入的路径转化为文件、降低读取速率
	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 307
	expectedUrls := []string {
		"https://bj.xcar.com.cn",
		"https://sh.xcar.com.cn",
		"https://gz.xcar.com.cn",
	}
	expectedCities := []string {
		"City 北京","City 上海","City 广州",
	}

	// 判断查到的数据数量是否等于网页内包含的数据量
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d " +
			"requests; but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but" + "was %s",
				i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d " +
			"requests; but had %d",
			resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city #%d: %s;but " + "was %s",
				i, city, result.Items[i].(string))
		}
	}
}
