package engine

import "log"

func StartJobServer(){
	go func() {
		log.Println("StartJobServer...")

		// 读取 log.data 到内存

		// 没有 log.data 则创建

		// 定时同步更新到 log.data


	}()
}
