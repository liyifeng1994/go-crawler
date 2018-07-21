package main

import (
	"lyf/crawler/engine"
	"lyf/crawler/persist"
	"lyf/crawler/scheduler"
	"lyf/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persistr.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}
