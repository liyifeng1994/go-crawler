package main

import (
	"log"
	"fmt"

	"github.com/olivere/elastic"

	"lyf/crawler/rpcsupport"
	"lyf/crawler/persist"
	"lyf/crawler/config"
)

func main() {
	log.Fatal(serveRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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
