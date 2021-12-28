package main

import (
	"github.com/mangenotwork/extras/apps/IM-User/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){

	logger.Info(utils.Logo)
	logger.Info("Starting IM-User ......")
	conf.InitConf()
	engine.StartJob()

	if conf.Arg.HttpServer.Open {
		engine.StartHTTP()
	}

	if conf.Arg.GrpcServer.Open {
		engine.StartRPC()
	}

	select {}
}
