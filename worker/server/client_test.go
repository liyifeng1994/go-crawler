package main

import (
	"fmt"
	"time"
	"testing"

	"lyf/crawler/config"
	"lyf/crawler/rpcsupport"
	"lyf/crawler/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// TODO: Use a fake fetcher to handle the url.
	// So we don't get data from zhenai.com
	req := worker.Request{
		Url: "http://album.zhenai.com/u/1314495053",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "风中的蒲公英",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}

	// TODO: Verify results
}
