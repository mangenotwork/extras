/*
	屏蔽词与词语白名单的存储， 主要由第三方redis进行存储
	数据存储结构使用set
 */

package model

import (
	"github.com/mangenotwork/extras/common/utils"
	"sort"
	"sync"
)

var Words = &wordMaps{
	words : make(map[string]struct{}),
	size : 0,
}

type wordMaps struct {
	words map[string]struct{}
	mtx sync.Mutex
	size int
}

func (wmp *wordMaps) Add(key string) {
	wmp.mtx.Lock()
	defer wmp.mtx.Unlock()
	wmp.words[key] = struct{}{}
	wmp.size++
}

func (wmp *wordMaps) Get() []string {
	wmp.mtx.Lock()
	defer wmp.mtx.Unlock()
	keyList := make([]string, 0, len(wmp.words))
	for k, _ := range wmp.words {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	return keyList
}

func (wmp *wordMaps) Del(key string) {
	wmp.mtx.Lock()
	defer wmp.mtx.Unlock()
	delete(wmp.words, key)
	wmp.size--
}

func (wmp *wordMaps) IsHave(key string) (ok bool) {
	_,ok = wmp.words[key]
	return
}

func (wmp *wordMaps) Save(fileName string) {
	go func() {
		wordStr := ""
		for k,_ := range wmp.words {
			wordStr = wordStr + k +"\n"
		}
		buf := utils.Compressed(wordStr)
		utils.FileWrite(fileName, buf)
	}()
}
