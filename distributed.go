package main

import (
	"fmt"

	"lyf/crawler/engine"
	"lyf/crawler/scheduler"
	"lyf/crawler/zhenai/parser"
	"lyf/crawler/config"
	itemsaver "lyf/crawler/persist/client"
	worker "lyf/crawler/worker/client"
)

func main() {
	itemChan, err := itemsaver.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	processor, err := worker.CreateProcessor(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
