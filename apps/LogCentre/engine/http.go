package engine

import (
	"github.com/mangenotwork/extras/apps/LogCentre/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartHttp(){
	go func() {
		logger.Info("StartHttp")
		mux := httpser.NewEngine()

		// 获取所有table
		mux.Router("/table",handler.GetLogTable)
		// 获取日志, 时间段参数
		mux.Router("/check/time",handler.CheckLogTime)
		// 获取日志, 前多少个
		mux.Router("/check/count",handler.CheckLogCount)
		// 查看日志文件列表
		mux.Router("/log/dir",handler.LogDir)
		// 查看指定日志文件内容
		mux.Router("/log/file",handler.LogFile)
		// 下载日志文件
		mux.Router("/log/upload",handler.LogUpload)

		mux.Run()

	}()
}

