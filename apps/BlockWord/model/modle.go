/*
	屏蔽词与词语白名单的存储， 主要由第三方redis进行存储
	数据存储结构使用set
 */

package model

import (
	"sort"
	"sync"

	"github.com/garyburd/redigo/redis"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
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

const (
	BlockWordKey = "bw"
	WhiteWordKey = "ww"
)

func addWord(key, word string) (err error) {
	c := conn.RedisConn().Get()
	defer c.Close()
	_, err = c.Do("SADD", key, word)
	if err != nil {
		logger.Error(err)
	}
	return
}

func getWord(key string) (res []string, err error) {
	c := conn.RedisConn().Get()
	defer c.Close()
	res, err = redis.Strings(c.Do("SMEMBERS", key))
	return
}

func delWord(key, word string) (err error){
	c := conn.RedisConn().Get()
	defer c.Close()
	_, err = c.Do("SREM", key, word)
	return
}

type Word interface {
	Add(word string)
	Get() []string
	Del(word string)
}

func NewWord(key string) Word {
	if key == WhiteWordKey {
		return new(whiteWord)
	}
	return new(blockWord)
}

type blockWord struct {}

func (*blockWord) Add(word string) {
	if err := addWord(BlockWordKey, word); err != nil {
		logger.Error(err)
	}
}

func (*blockWord) Get() []string {
	res, err := getWord(BlockWordKey)
	if err != nil {
		return nil
	}
	return res
}

func (*blockWord) Del(word string) {
	if err := delWord(BlockWordKey, word); err != nil {
		logger.Error(err)
	}
}

type whiteWord struct {}

func (*whiteWord) Add(word string) {
	if err := addWord(WhiteWordKey, word); err != nil {
		logger.Error(err)
	}
}

func (*whiteWord) Get() []string {
	res, err := getWord(WhiteWordKey)
	if err != nil {
		return nil
	}
	return res
}

func (*whiteWord) Del(word string) {
	if err := delWord(WhiteWordKey, word); err != nil {
		logger.Error(err)
	}
}


