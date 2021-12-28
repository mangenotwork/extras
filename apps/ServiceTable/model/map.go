package model

import (
	"sync"
	"time"

	"github.com/mangenotwork/extras/apps/ServiceTable/raft"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

/*
	定义数据结构
	key : value  存储值
 */

var (
	KVData map[string]*KVDataValue
	mapDataOnce sync.Once
)

type KVDataValue struct {
	Value string
	Expire int64
}

func InitMapData(){
	mapDataOnce.Do(
		func() {
			KVData = make(map[string]*KVDataValue)
		})
}

// 增,改
// Command : KVAdd key value
func KVAdd(key, value string) {
	KVData[key] = &KVDataValue{
		Value: value,
		Expire: -1, // 永久不过期
	}
	Key.Insert(key, "kv")
}
func KVAddAt(key, value string) {
	KVAdd(key, value)
	raft.NewLogData("KVAdd "+key+" "+value).Write()
}

// Command : KVAddExpire key value expire
func KVAddExpire(key, value string, expire int64) {
	KVData[key] = &KVDataValue{
		Value: value,
		Expire: expire,
	}
	Key.Insert(key, "kv")
}
func KVAddExpireAt(key, value string, expire int64) {
	KVAddExpire(key, value, expire)
	raft.NewLogData("KVAddExpire "+key+" "+value+" "+utils.Any2String(expire)).Write()
}

// Command : KVExpire key expire
func KVExpire(key string, expire int64) int {
	if v,ok := KVData[key]; ok {
		v.Expire = expire
		return 1
	}
	return 0
}
func KVExpireAt(key string, expire int64) int {
	rse := KVExpire(key, expire)
	if rse == 1 {
		raft.NewLogData("KVExpire "+key+" "+utils.Any2String(expire)).Write()
	}
	return rse
}

// 删
// Command : KVDel key
func KVDel(key string) int {
	if _,ok := KVData[key]; ok {
		delete(KVData, key)
		Key.Remove(key)
		return 1
	}
	return 0
}
func KVDelAt(key string) int {
	rse := KVDel(key)
	if rse == 1 {
		raft.NewLogData("KVDel "+key).Write()
	}
	return rse
}

// 查
// Command : KVGet key
func KVGet(key string) (bool, string) {
	logger.Info("KVGet", KVData[key])
	if v,ok := KVData[key]; ok {
		logger.Info(v)
		if v.Expire < 0 || v.Expire > time.Now().Unix() {
			return true,v.Value
		} else {
			// 过期删除
			KVDelAt(key)
		}
	}
	return false,""
}


