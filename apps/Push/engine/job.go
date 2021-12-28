package engine

import (
	"github.com/mangenotwork/extras/apps/Push/model"
	"github.com/mangenotwork/extras/common/logger"
)

func StartJobServer(){
	go func() {
		logger.Info("StartJobServer")

		model.TopicMap = make(map[string]*model.Topic)

	}()
}