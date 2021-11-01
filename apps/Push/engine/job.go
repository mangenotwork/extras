package engine

import (
	"github.com/mangenotwork/extras/apps/Push/model"
	"log"
)

func StartJobServer(){
	go func() {
		log.Println("StartJobServer")

		model.TopicMap = make(map[string]*model.Topic)

	}()
}