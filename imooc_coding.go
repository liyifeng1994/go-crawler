package main

import (
	"fmt"

	"lyf/crawler/engine"
	"lyf/crawler/persist"
	"lyf/crawler/scheduler"
	parser "lyf/crawler/imooc/parser/coding"
)

func main() {
	itemChan, err := persistr.ItemSaver("dating_course_coding")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	var seeds []engine.Request
	for i := 0; i < 300; i++ {
		seeds = append(seeds, engine.Request{
			Url:        fmt.Sprintf("https://coding.imooc.com/class/%d.html", i+1),
			ParserFunc: parser.ParseCoding,
		})
	}
	e.Run(seeds...)
}
