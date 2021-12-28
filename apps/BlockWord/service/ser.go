package service

import (
	"github.com/mangenotwork/extras/apps/BlockWord/model"
	"github.com/mangenotwork/extras/common/logger"
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
		logger.Error(err)
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

func WhiteDelWord(word string) {
	if !model.WhiteWord.Search([]rune(word)) {
		return
	}
	model.WhiteWord.Remove(word)
	model.NewWord(model.WhiteWordKey).Del(word)
	return
}

func WhiteGetWord() []string {
	return model.NewWord(model.WhiteWordKey).Get()
}

func InitWord(){
	for _, v := range model.NewWord(model.BlockWordKey).Get() {
		BlockWorkTrie.Add(v)
	}
	for _, v2 := range model.NewWord(model.WhiteWordKey).Get() {
		model.WhiteWord.Insert(v2)
	}
}

