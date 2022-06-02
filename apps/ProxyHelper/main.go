package main

import (
	"github.com/mangenotwork/extras/apps/ProxyHelper/service"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func main(){

	conf.InitConf()
	logger.InitLogger()

	logger.Info(utils.Logo)
	logger.Info("Starting ProxyHelper server.")

	service.SockerRun()
}