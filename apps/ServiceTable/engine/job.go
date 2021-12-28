package engine

import (
	"github.com/mangenotwork/extras/apps/ServiceTable/service"
	"github.com/mangenotwork/extras/common/logger"
)

func StartJobServer(){
	go func() {
		logger.Info("StartJobServer...")

		// 初始化将日志数据写入内存
		service.LogDataToMemory()

		// 定时同步更新到 log.data


	}()
}
