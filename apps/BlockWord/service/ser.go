package service

import (
	"bufio"
	"github.com/mangenotwork/extras/apps/BlockWord/model"
	"github.com/mangenotwork/extras/common/utils"
	"log"
	"os"
	"strings"
)

var BlockWorkTrie = model.NewTrie()
var FileName = "./word.action"


// 添加屏蔽词
func AddWord(word string) error {
	if model.Words.IsHave(word) {
		return nil
	}
	BlockWorkTrie.Add(word)
	model.Words.Add(word)
	model.Words.Save(FileName)
	return nil
}

// 删除屏蔽词
func DelWord(word string) error {
	if !model.Words.IsHave(word) {
		return nil
	}
	if err := BlockWorkTrie.RemoveWord(word); err!=nil {
		return err
	}
	model.Words.Del(word)
	model.Words.Save(FileName)
	return nil
}

// 查看当前屏蔽词
func GetWord() []string {
	return model.Words.Get()
}


// 存储屏蔽词设计
//	增删都会去写入日志操作文件 .action，
//	程序启动会将.action读取然后写入全局词前缀树，词map
func InitWord(){
	file, err := os.Open(FileName)
	if err != nil {
		log.Println("[Error] 存储屏蔽词的文件不存在!")
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Bytes()
		str := utils.Decompress(data)
		for _,v := range strings.Split(str, "\n") {
			if err := BlockWorkTrie.Add(v); err==nil {
				model.Words.Add(v)
			}
		}
	}
}

