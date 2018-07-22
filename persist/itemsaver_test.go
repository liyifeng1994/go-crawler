package persistr

import (
	"context"
	"encoding/json"
	"testing"

	"lyf/crawler/engine"
	"lyf/crawler/model"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
	excepted := engine.Item{
		Url:  "http://album.zhenai.com/u/1314495053",
		Type: "zhenai",
		Id:   "13144950531",
		Payload: model.Profile{
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
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"

	// Save expected item
	err = Save(client, index, excepted)
	if err != nil {
		panic(err)
	}

	// Fetch save item
	resp, err := client.Get().Index(index).Type(excepted.Type).Id(excepted.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)

	profileActual, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = profileActual

	// Verify result
	if actual != excepted {
		t.Errorf("got %v; expected %v", actual, excepted)
	}
}
