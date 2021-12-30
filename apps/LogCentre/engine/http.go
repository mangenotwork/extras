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

	return mux
}