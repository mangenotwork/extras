package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartWS(){
	go func() {
		logger.Info("StartWS")
	}()
}