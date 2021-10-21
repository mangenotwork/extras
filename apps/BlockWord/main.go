package main

import (
	"github.com/mangenotwork/extras/apps/BlockWord/engine"
	"github.com/mangenotwork/extras/common/utils"
	"log"
)

func main(){

	// 打印logo
	log.Println(utils.Logo)
	log.Println("Starting block word http server")

	engine.StartJobSrc()
	engine.StartRpcSrc()
	engine.StartHttpSrc()
}
