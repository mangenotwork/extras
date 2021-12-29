package main

import (
	"github.com/mangenotwork/extras/apps/LogCentre/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){
	conf.InitConf()
	logger.Info(utils.Logo)
	logger.Info("Starting log center http server")
	engine.StartJob()

	if conf.Arg.HttpServer.Open {
		engine.StartHttp()
	}

	if conf.Arg.UdpServer.Open {
		engine.StartUdp()
	}

	select {}
}