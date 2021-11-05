package main

import (
	"github.com/mangenotwork/extras/apps/WordHelper/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/utils"
	"log"
)

func main(){
	log.Println(utils.Logo)
	log.Println("Starting img helper http server")
	conf.InitConf()
	engine.StartJobServer()

	if conf.Arg.HttpServer.Open {
		engine.StartHttpServer()
	}

	if conf.Arg.GrpcServer.Open {
		engine.StartRpcServer()
	}

	select {}
}