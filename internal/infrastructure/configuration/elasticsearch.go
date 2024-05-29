package configuration

import (
	"github.com/olivere/elastic/v7"
	"log"
)

func ProvideElasticClient() *elastic.Client {
	client, err := elastic.NewClient()
	if err != nil {
		log.Fatal("failed to create elasticsearch client")
	}
	return client
}
