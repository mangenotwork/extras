package main

import (
	"github.com/mangenotwork/extras/apps/WordHelper/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){
	conf.InitConf()

	logger.Info(utils.Logo)
	logger.Info("Starting img helper http server")

	engine.StartJobServer()

	if conf.Arg.HttpServer.Open {
		engine.StartHttpServer()
	}

	if conf.Arg.GrpcServer.Open {
		engine.StartRpcServer()
	}

	select {}
}