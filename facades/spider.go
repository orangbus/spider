package facades

import (
	"log"

	"github.com/orangbus/spider"
	"github.com/orangbus/spider/contracts"
)

func Spider() contracts.Spider {
	instance, err := spider.App.Make(spider.Binding)
	if err != nil {
		log.Println(err)
		return nil
	}

	return instance.(contracts.Spider)
}
