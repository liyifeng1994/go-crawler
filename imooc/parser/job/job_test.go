package job

import (
	"testing"
	"io/ioutil"

	"lyf/crawler/model"
	"lyf/crawler/engine"
)

func TestParseJob(t *testing.T) {
	contents, err := ioutil.ReadFile("job_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseJob(contents, "https://class.imooc.com/sc/18")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %d", len(result.Items))
	}

	actual := result.Items[0]

	excepted := engine.Item{
		Url:  "https://class.imooc.com/sc/18",
		Type: "imooc",
		Id:   "18",
		Payload: model.JobCourse{
			Job:            "JAVA攻城狮培养计划",
			Title:          "Java零基础入门",
			Time:           "40小时",
			StudentsNumber: 5556,
			Score:          9.89,
			Price:          358.00,
		},
	}

	if actual != excepted {
		t.Errorf("expected %v; \nbut was %v", excepted, actual)
	}
}
