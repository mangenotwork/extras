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
	// 获取http 请求日志, 时间段参数
	mux.Handle("/http/req", m(http.HandlerFunc(handler.HttpReqLog)))


	return mux
}

