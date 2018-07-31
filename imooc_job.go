package main

import (
	"fmt"

	"lyf/crawler/engine"
	"lyf/crawler/persist"
	"lyf/crawler/scheduler"
	parser "lyf/crawler/imooc/parser/job"
)

func main() {
	itemChan, err := persistr.ItemSaver("dating_course_job")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	var seeds []engine.Request
	for i := 0; i < 50; i++ {
		seeds = append(seeds, engine.Request{
			Url:        fmt.Sprintf("https://class.imooc.com/sc/%d", i+1),
			ParserFunc: parser.ParseJob,
		})
	}
	e.Run(seeds...)
}
