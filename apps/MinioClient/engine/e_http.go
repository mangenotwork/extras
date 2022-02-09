package engine

import (
	"github.com/mangenotwork/extras/apps/MinioClient/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartHttp() {
	go func() {
		logger.Info("StartHttp")
		mux := httpser.SimpleEngine()

		mux.Router("/", handler.Hello)
		mux.Router("/err", handler.Error)

		// 查看是否连接Minio
		mux.Router("/hasConn", handler.HasConn)

		// 创建桶
		mux.Router("/bucket/add", handler.BucketAdd)

		// 查看所有桶
		mux.Router("/bucket/list", handler.BucketList)

		// 查看桶文件列表
		mux.Router("/bucket/files", handler.BucketFiles)

		mux.Run()

	}()
}

