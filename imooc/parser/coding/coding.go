package coding

import (
	"regexp"

	. "lyf/crawler/util"
	"lyf/crawler/engine"
	"lyf/crawler/model"
)

var idUrlRe = regexp.MustCompile(`https://coding.imooc.com/class/([\d]+).html`)
var titleRe = regexp.MustCompile(`<h3 class='fixed-course-name' title='([^']+)'>`)
var originalTitleRe = regexp.MustCompile(`<title>([^-]+)-慕课网[实战]*[课程]*</title>`)
var teacherRe = regexp.MustCompile(`<div class="nickname">([^<]+)</div>`)
var levelRe = regexp.MustCompile(`<span class="meta-value"><strong>([初|中|高]级)</strong></span>`)
var timeRe = regexp.MustCompile(`<span class="meta-value"><strong>[ ]*([\d]+小时[^<]*)</strong></span>`)
var studentsNumberRe = regexp.MustCompile(`<span class="meta-value"><strong>([\d]+)</strong></span>`)
var scoreRe = regexp.MustCompile(`<span class="meta-value"><strong>([^分]+)分</strong></span>`)
var priceRe = regexp.MustCompile(`<span class="r fixed-nav-prices">￥([^<]+)</span>`)
var deletedRe = regexp.MustCompile(`<span>报名人数已满</span>`)

func ParseCoding(contents []byte, url string) engine.ParseResult {
	course := model.CodingCourse{
		Title:          ExtractString(contents, titleRe),
		OriginalTitle:  ExtractString(contents, originalTitleRe),
		Teacher:        ExtractString(contents, teacherRe),
		Level:          ExtractString(contents, levelRe),
		Time:           ExtractString(contents, timeRe),
		StudentsNumber: ExtractInt(contents, studentsNumberRe),
		Score:          ExtractFloat(contents, scoreRe),
		Price:          ExtractFloat(contents, priceRe),
		Deleted:        deletedRe.Match(contents),
	}

	if course.Title == "" {
		return engine.ParseResult{}
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
