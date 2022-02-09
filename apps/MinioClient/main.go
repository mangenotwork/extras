package main

import (
	"github.com/mangenotwork/extras/apps/MinioClient/engine"
	"github.com/mangenotwork/extras/apps/MinioClient/model"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){
	conf.InitConf()
	logger.InitLogger()

	logger.Info(utils.Logo)
	logger.Info("Starting block word http server")

	model.InitMinioClient()

	engine.StartJob()

	if conf.Arg.HttpServer.Open {
		engine.StartHttp()
	}

	select {}
}
