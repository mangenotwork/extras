package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartJobServer(){
	go func() {
		logger.Info("StartJobServer...")
	}()
}

