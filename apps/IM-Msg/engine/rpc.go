package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartRPC(){
	go func() {
		logger.Info("StartRPC")
	}()
}