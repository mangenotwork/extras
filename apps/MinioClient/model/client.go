package model

import (
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/minio/minio-go/v6"
)

var MinioClient *minio.Client

func InitMinioClient(){
	var err error
	MinioClient, err = minio.New(conf.Arg.Minio.Host, conf.Arg.Minio.Access, conf.Arg.Minio.Secret, false)
	if err != nil {
		logger.Error("创建 MinIO 客户端失败", err)
		return
	}
	logger.Info("创建 MinIO 客户端成功")
}