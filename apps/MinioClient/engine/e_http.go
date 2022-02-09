package engine

import (
	"github.com/mangenotwork/extras/apps/MinioClient/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartHttp() {
	go func() {
		logger.Info("StartHttp")
		mux := httpser.NewEngine()

		// 查看是否连接Minio
		mux.Router("/hasConn", handler.HasConn)


		mux.Run()

	}()
}