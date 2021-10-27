package service

import (
	"github.com/mangenotwork/extras/apps/BlockWord/model"
	"log"
)

var BlockWorkTrie = model.NewTrie()
var FileName = "./word.action"


// 添加屏蔽词
func AddWord(word string) {
	if BlockWorkTrie.Search(word) {
		return
	}
	BlockWorkTrie.Add(word)
	model.NewWord(model.BlockWordKey).Add(word)
	return
}

// 删除屏蔽词
func DelWord(word string) {
	if !BlockWorkTrie.Search(word) {
		return
	}
	if err := BlockWorkTrie.RemoveWord(word); err!=nil {
		log.Println(err)
		return
	}
	model.NewWord(model.BlockWordKey).Del(word)
	return
}

// 查看当前屏蔽词
func GetWord() []string {
	return model.NewWord(model.BlockWordKey).Get()
}

func WhiteAddWord(word string) {
	if model.WhiteWord.Search([]rune(word)) {
		return
	}
	model.WhiteWord.Insert(word)
	model.NewWord(model.WhiteWordKey).Add(word)
	return
}

// 删除屏蔽词
func WhiteDelWord(word string) {
	if !model.WhiteWord.Search([]rune(word)) {
		return
	}
	model.WhiteWord.Remove(word)
	model.NewWord(model.WhiteWordKey).Del(word)
	return
}

// 查看当前屏蔽词
func WhiteGetWord() []string {
	return model.NewWord(model.WhiteWordKey).Get()
}

// 存储屏蔽词设计
//	增删都会去写入日志操作文件 .action，
//	程序启动会将.action读取然后写入全局词前缀树，词map
func InitWord(){
	for _, v := range model.NewWord(model.BlockWordKey).Get() {
		BlockWorkTrie.Add(v)
	}
	for _, v2 := range model.NewWord(model.WhiteWordKey).Get() {
		model.WhiteWord.Insert(v2)
	}
}

