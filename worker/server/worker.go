package main

import (
	"fmt"
	"log"

	"lyf/crawler/rpcsupport"
	"lyf/crawler/worker"
	"lyf/crawler/config"
)

func main() {
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))
}
