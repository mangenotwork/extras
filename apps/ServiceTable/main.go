package main

import (
	"github.com/mangenotwork/extras/apps/ConfigCenter/engine"
	"github.com/mangenotwork/extras/apps/ConfigCenter/raft"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/utils"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	go func() {
		raft.StartCluster()
	}()


	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	select {
	case s := <-ch:
		// TODO 通知退出
		log.Println("通知退出....")
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}
}



