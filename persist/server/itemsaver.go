package main

import (
	"log"
	"fmt"
	"flag"

	"github.com/olivere/elastic"

	"lyf/crawler/rpcsupport"
	"lyf/crawler/persist"
	"lyf/crawler/config"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return rpcsupport.ServeRpc(host, &persistr.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
