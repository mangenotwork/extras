package main

import (
	"github.com/mangenotwork/extras/apps/Push/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/utils"
	"log"
)

func main(){

	log.Println(utils.Logo)
	log.Println("Starting push server")
	conf.InitConf()
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
