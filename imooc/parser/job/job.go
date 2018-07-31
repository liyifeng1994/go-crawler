package job

import (
	"fmt"
	"regexp"

	. "lyf/crawler/util"
	"lyf/crawler/engine"
	"lyf/crawler/model"
)

var idUrlRe = regexp.MustCompile(`https://class.imooc.com/sc/([\d]+)`)
var jobRe = regexp.MustCompile(`<h2 class='cat_name'>([^<]+)</h2>`)
var titleRe = regexp.MustCompile(`<h3 class='fixed-course-name'>([^<]+)</h3>`)
var timeRe = regexp.MustCompile(`<dd>([\d]+)<span>小时</span></dd>`)
var studentsNumberRe = regexp.MustCompile(`<dd>([\d]+)<span>人</span></dd>`)
var scoreRe = regexp.MustCompile(`<dd>([^<]+)<span>分</span></dd>`)
var priceRe = regexp.MustCompile(`<span class="pricespan">￥([^<]+)</span>`)

func ParseJob(contents []byte, url string) engine.ParseResult {
	course := model.JobCourse{
		Job:            ExtractString(contents, jobRe),
		Title:          ExtractString(contents, titleRe),
		Time:           fmt.Sprintf("%d小时", ExtractInt(contents, timeRe)),
		StudentsNumber: ExtractInt(contents, studentsNumberRe),
		Score:          ExtractFloat(contents, scoreRe),
		Price:          ExtractFloat(contents, priceRe),
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "imooc",
				Id:      ExtractString([]byte(url), idUrlRe),
				Payload: course,
			},
		},
	}

	return result
}
