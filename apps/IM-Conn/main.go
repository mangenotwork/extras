package main

import (
	"github.com/mangenotwork/extras/apps/IM-Conn/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){
	conf.InitConf()
	logger.InitLogger()

	logger.Info(utils.Logo)
	logger.Info("Starting IM-Conn ......")

	engine.StartJob()
	engine.StartMQ()

	if conf.Arg.HttpServer.Open {
		engine.StartWS()
	}

	if conf.Arg.TcpServer.Open {
		engine.StartTCP()
	}

	if conf.Arg.UdpServer.Open {
		engine.StartUDP()
	}

	select {}
}
