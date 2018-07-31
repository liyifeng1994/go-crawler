package coding

import (
	"testing"
	"io/ioutil"

	"lyf/crawler/model"
	"lyf/crawler/engine"
)

func TestParseCoding(t *testing.T) {
	contents, err := ioutil.ReadFile("coding_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCoding(contents, "https://coding.imooc.com/class/237.html")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %d", len(result.Items))
	}

	actual := result.Items[0]

	excepted := engine.Item{
		Url:  "https://coding.imooc.com/class/237.html",
		Type: "imooc",
		Id:   "237",
		Payload: model.CodingCourse{
			Title:          "分布式事务实践 解决数据一致性",
			OriginalTitle:  "分布式事务实践",
			Teacher:        "大漠风",
			Level:          "高级",
			Time:           "14小时",
			StudentsNumber: 195,
			Score:          9.75,
			Price:          348.00,
			Deleted:        false,
		},
	}

	if actual != excepted {
		t.Errorf("expected %v; \nbut was %v", excepted, actual)
	}
}
