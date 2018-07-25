package main

import (
	"fmt"
	"log"
	"flag"

	"lyf/crawler/rpcsupport"
	"lyf/crawler/worker"
	"lyf/crawler/fetcher"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
