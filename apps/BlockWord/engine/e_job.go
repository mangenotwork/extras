package engine

import (
	"github.com/mangenotwork/extras/apps/BlockWord/service"
	"github.com/mangenotwork/extras/common/logger"
)

func StartJobServer(){
	go func() {
		// 初始化树
		// 是否有存储文件，没有创建，存在读取数据
		logger.Info("加载屏蔽词...")
		service.InitWord()
		logger.Info("加载完成!")
	}()
}
