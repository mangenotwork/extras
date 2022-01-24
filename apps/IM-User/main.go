package main

import (
	"github.com/mangenotwork/extras/apps/IM-User/dao"
	"github.com/mangenotwork/extras/apps/IM-User/engine"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){
	conf.InitConf()
	logger.InitLogger()

	conn.InitMysqlDriver() // 初始化 mysql driver
	conn.InitMysqlGorm() // 初始化 mysql gorm

	dao.InitUserBaseTable() // 初始化表


	logger.Info(utils.Logo)
	logger.Info("Starting IM-User ......")

	engine.StartJob()

	if conf.Arg.HttpServer.Open {
		engine.StartHTTP()
	}

	if conf.Arg.GrpcServer.Open {
		engine.StartRPC()
	}

	select {}
}
