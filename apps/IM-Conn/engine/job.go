package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartJob(){
	go func() {
		logger.Info("StartJob")
	}()
}