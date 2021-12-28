package main

import (
	"github.com/mangenotwork/extras/apps/ImgHelper/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){
	logger.Info(utils.Logo)
	logger.Info("Starting img helper http server")
	conf.InitConf()
	engine.StartInitialization()
	engine.StartJobServer()

	if conf.Arg.HttpServer.Open {
		engine.StartHttpServer()
	}

	if conf.Arg.GrpcServer.Open {
		engine.StartRpcServer()
	}

	select {}
}