package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartMQ(){
	go func() {
		logger.Info("StartMQ")
	}()
}