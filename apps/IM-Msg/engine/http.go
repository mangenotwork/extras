package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartHTTP(){
	go func() {
		logger.Info("StartHTTP")
	}()
}