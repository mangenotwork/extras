package engine

import (
	"github.com/mangenotwork/extras/common/logger"
)

func StartUdp(){
	go func() {
		logger.Info("StartUdp")
	}()
}