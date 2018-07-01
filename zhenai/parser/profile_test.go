package parser

import (
	"testing"
	"io/ioutil"
	"lyf/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "风中的蒲公英")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %d", len(result.Items))
	}

	profile := result.Items[0].(model.Profile)

	excepted := model.Profile{
		Name:       "风中的蒲公英",
		Gender:     "女",
		Age:        41,
		Height:     158,
		Weight:     48,
		Income:     "3001-5000元",
		Marriage:   "离异",
		Education:  "中专",
		Occupation: "公务员",
		Hukou:      "四川阿坝",
		Xinzuo:     "处女座",
		House:      "已购房",
		Car:        "未购车",
	}

	if profile != excepted {
		t.Errorf("expected %v; but was %v", excepted, profile)
	}
}
