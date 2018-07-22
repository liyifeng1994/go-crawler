package client

import (
	"lyf/crawler/engine"
	"lyf/crawler/config"
	"lyf/crawler/worker"
	"lyf/crawler/rpcsupport"
)

func CreateProcessor(host string) (engine.Processor, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)

		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
