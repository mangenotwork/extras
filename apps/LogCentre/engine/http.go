package engine

import (
	"net/http"

	"github.com/mangenotwork/extras/apps/LogCentre/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/middleware"
)

func StartHttp(){
	go func() {
		logger.Info("StartHttp")
		httpser.HttpServer(Router())
	}()
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	m := middleware.Base
	mux.Handle("/hello", m(http.HandlerFunc(handler.Hello)))
	mux.Handle("/", m(http.HandlerFunc(handler.Hello)))

	// 获取所有table
	mux.Handle("/table", m(http.HandlerFunc(handler.GetLogTable)))
	// 获取日志, 时间段参数
	mux.Handle("/check/time", m(http.HandlerFunc(handler.CheckLogTime)))
	// 获取日志, 前多少个
	mux.Handle("/check/count", m(http.HandlerFunc(handler.CheckLogCount)))
	// 查看日志文件列表
	mux.Handle("/log/dir", m(http.HandlerFunc(handler.LogDir)))
	// 查看指定日志文件内容
	// 下载日志文件

	return mux
}

