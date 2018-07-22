package persistr

import (
	"log"

	"github.com/olivere/elastic"

	"lyf/crawler/engine"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := Save(s.Client, s.Index, item)
	if err == nil {
		log.Printf("Item %v saved.", item)
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
