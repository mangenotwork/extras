package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartRpcServer(){
	go func() {
		logger.Info("StartRpcServer...")
	}()
}