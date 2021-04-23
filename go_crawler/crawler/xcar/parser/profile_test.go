package parser

import (
	"go_crawler/crawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err !=nil {
		panic(err)
	}

	result := ParseProfile(contents)

	if len(result.Items) != 1 {
		t.Errorf("Result should contain 1" + "element;but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Price: "38.06万",
		Factory: "奇瑞·捷豹路虎",
		Level:"中型车",
		Motor:"184kW(2.0L涡轮增压)",
		PowerType:"汽油机",
		Year: "2019",
		Speed: "240",
		Acceleration: "7.3",
		Oil: "6.7",
	}

	if profile!= expected {
		 t.Errorf("expected %v ; but was %v",
		 	expected, profile)
	}
}
