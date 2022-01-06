package engine

import (
	"github.com/mangenotwork/extras/apps/BlockWord/service"
	"github.com/mangenotwork/extras/common/logger"
)

func StartJobServer(){
	go func() {
		logger.Info("加载屏蔽词...")
		service.InitWord()
		logger.Info("加载完成!")
	}()
}
