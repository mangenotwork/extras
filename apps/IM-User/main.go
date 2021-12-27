package main

import (
	"github.com/mangenotwork/extras/apps/IM-User/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/utils"
	"log"
)

func main(){

	log.Println(utils.Logo)
	log.Println("Starting IM-User ......")
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
