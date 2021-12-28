package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartTCP(){
	go func() {
		logger.Info("StartTCP")
	}()
}