package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartUDP(){
	go func() {
		logger.Info("StartUDP")
	}()
}