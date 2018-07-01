package parser

import (
	"lyf/crawler/engine"
	"lyf/crawler/model"
	"regexp"
	"strconv"
)

var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{
		Name:       name,
		Gender:     extractString(contents, genderRe),
		Age:        extractInt(contents, ageRe),
		Height:     extractInt(contents, heightRe),
		Weight:     extractInt(contents, weightRe),
		Income:     extractString(contents, incomeRe),
		Marriage:   extractString(contents, marriageRe),
		Education:  extractString(contents, educationRe),
		Occupation: extractString(contents, occupationRe),
		Hukou:      extractString(contents, hukouRe),
		Xinzuo:     extractString(contents, xinzuoRe),
		House:      extractString(contents, houseRe),
		Car:        extractString(contents, carRe),
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	int, err := strconv.Atoi(string(extractString(contents, re)))
	if err == nil {
		return int
	}
	return 0
}
