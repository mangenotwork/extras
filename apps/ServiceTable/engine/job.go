package engine

import (
	"github.com/mangenotwork/extras/apps/ServiceTable/raft"
	"log"
)

func StartJobServer(){
	go func() {
		log.Println("StartJobServer...")

		raft.SetTestData()

		raft.LogDataToMemory()

		// 定时同步更新到 log.data


	}()
}
