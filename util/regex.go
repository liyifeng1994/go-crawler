package util

import (
	"regexp"
	"strconv"
)

func ExtractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func ExtractInt(contents []byte, re *regexp.Regexp) int {
	int, err := strconv.Atoi(string(ExtractString(contents, re)))
	if err == nil {
		return int
	}
	return 0
}

func ExtractFloat(contents []byte, re *regexp.Regexp) float32 {
	float, err := strconv.ParseFloat(string(ExtractString(contents, re)), 32)
	if err == nil {
		return float32(float)
	}
	return 0
}
