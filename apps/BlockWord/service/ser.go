package service

import (
	"bufio"
	"github.com/mangenotwork/extras/apps/BlockWord/model"
	"github.com/mangenotwork/extras/common/utils"
	"log"
	"os"
)

var BlockWorkTrie = model.NewTrie()
var FileName = "./word.action"


// 添加屏蔽词
func AddWord(word string) error {
	if model.Words.IsHave(word) {
		return nil
	}
	if err := BlockWorkTrie.AddWord(word); err!=nil {
		return err
	}
	model.Words.Add(word)
	buf := utils.Compressed("add|"+word)
	buf = append(buf, '\n')
	utils.FileWrite(FileName, buf)
	return nil
}

// 删除屏蔽词
func DelWord(word string) error {
	if !model.Words.IsHave(word) {
		return nil
	}
	if err := BlockWorkTrie.Remove(word); err!=nil {
		return err
	}
	model.Words.Del(word)
	buf := utils.Compressed("del|"+word)
	buf = append(buf, '\n')
	utils.FileWrite(FileName, buf)
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
		if len(str) < 5 {
			continue
		}
		doType := str[:3]
		word := str[4:]
		switch doType {
		case "add":
			if err := BlockWorkTrie.AddWord(word); err==nil {
				model.Words.Add(word)
			}
		case "del":
			if err := BlockWorkTrie.Remove(word); err==nil {
				model.Words.Del(word)
			}
		}
	}
}