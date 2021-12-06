package engine

import (
	"github.com/mangenotwork/extras/apps/ServiceTable/service"
	"log"
)

func StartJobServer(){
	go func() {
		log.Println("StartJobServer...")

		// 初始化将日志数据写入内存
		service.LogDataToMemory()

		// 定时同步更新到 log.data


	}()
}
