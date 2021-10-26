package main

import (
	"github.com/mangenotwork/extras/apps/BlockWord/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/utils"
	"log"
)

func main(){

	log.Println(utils.Logo)
	log.Println("Starting block word http server")
	conf.InitConf()
	engine.StartJobServer()

	if conf.Arg.HttpServer.Open {
		engine.StartHttpServer()
	}

	if conf.Arg.HttpServer.Open {
		engine.StartRpcServer()
	}

	select {}
}
