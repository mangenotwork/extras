package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mangenotwork/extras/apps/ServiceTable/engine"
	"github.com/mangenotwork/extras/apps/ServiceTable/raft"
	"github.com/mangenotwork/extras/apps/ServiceTable/service"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){
	conf.InitConf()
	logger.InitLogger()

	logger.Info(utils.Logo)
	logger.Info("Starting img helper http server")

	engine.StartJobServer()

	if conf.Arg.HttpServer.Open {
		engine.StartHttpServer()
	}

	if conf.Arg.GrpcServer.Open {
		engine.StartRpcServer()
	}

	service.InitRaft()
	go func() {
		raft.StartCluster()
	}()


	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	select {
	case s := <-ch:
		// TODO 通知退出
		logger.Info("通知退出....")
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}
}



