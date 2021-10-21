package engine

import (
	"github.com/mangenotwork/extras/apps/BlockWord/service"
	"log"
)

func StartJobSrc(){
	go func() {
		// 初始化树
		// 是否有存储文件，没有创建，存在读取数据
		log.Println("加载屏蔽词...")
		service.InitWord()
		log.Println("加载完成!")
	}()
}
