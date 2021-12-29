package main

import (
	"github.com/mangenotwork/extras/apps/Push/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){
	conf.InitConf()

	logger.Info(utils.Logo)
	logger.Info("Starting push server")

	engine.StartJobServer()
	engine.StartMqServer()

	if conf.Arg.HttpServer.Open {
		engine.StartHttpServer()
	}

	if conf.Arg.TcpServer.Open {
		engine.StartTcpServer()
	}
	if conf.Arg.UdpServer.Open {
		engine.StartUdpServer()
	}

	select {}
}
