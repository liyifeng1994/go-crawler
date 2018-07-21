package main

import (
	"lyf/crawler/engine"
	"lyf/crawler/persist"
	"lyf/crawler/scheduler"
	"lyf/crawler/zhenai/parser"
	"lyf/crawler/config"
)

func main() {
	itemChan, err := persistr.ItemSaver(config.ElasticIndex)
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
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
