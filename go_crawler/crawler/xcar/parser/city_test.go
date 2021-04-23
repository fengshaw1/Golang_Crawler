package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// 测试ParseCity
func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_data.html")
	if err !=nil {
		panic(err)
	}

	result := ParseCity(contents)

	fmt.Println(result)
}
