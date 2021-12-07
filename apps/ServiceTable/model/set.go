package model

import (
	"github.com/mangenotwork/extras/apps/ServiceTable/raft"
	"github.com/mangenotwork/extras/common/utils"
	"strings"
	"sync"
	"time"
)

/*
定义数据结构
key : set list  集合存储表
 */

var (
	SetData map[string]*set
	setDataOnce sync.Once
)

func InitSetData(){
	setDataOnce.Do(
		func() {
			SetData = make(map[string]*set)
	})
}

// Command : SetAdd key value1,value2,
// 集合添加数据
func SetAdd(key string, value []string) {
	// 不存在就创建
	if _,ok := SetData[key]; !ok {
		SetData[key] = newSet()
	}
	for _, v := range value {
		SetData[key].Add(v)
		Key.Insert(key, "set")
	}
}
func SetAddAt(key string, value []string) {
	SetAdd(key, value)
	// 存日志
	raft.NewLogData("SetAdd "+key+" "+strings.Join(value, ",")).Write()
}

// Command : SetAddExpire key value timeUnix
// 集合添加数据并指定过期时间
func SetAddExpire(key, value string, timeUnix int64) {
	// 不存在就创建
	if _,ok := SetData[key]; !ok {
		SetData[key] = newSet()
	}
	SetData[key].AddExpire(value, timeUnix)
	if timeUnix > time.Now().Unix() {
		Key.Insert(key, "set")
	}
}
func SetAddExpireAt(key, value string, timeUnix int64) {
	SetAddExpire(key, value, timeUnix)
	raft.NewLogData("SetAddExpire "+key+" "+value+" "+utils.Any2String(timeUnix)).Write()
}

// Command : SetValueExpire key value timeUnix
// 指定集合数据过期时间
func SetValueExpire(key, value string, timeUnix int64) int {
	// 不存在就创建
	if _,ok := SetData[key]; !ok {
		return 0
	}
	rse := SetData[key].Expire(value, timeUnix)
	return rse
}
func SetValueExpireAt(key, value string, timeUnix int64) int {
	rse := SetValueExpire(key, value, timeUnix)
	if rse == 1 {
		raft.NewLogData("SetValueExpire "+key+" "+value+" "+utils.Any2String(timeUnix)).Write()
	}
	return rse
}

// Command : SetGet key
// 获取集合所有数据
func SetGet(key string) []string {
	if v, ok := SetData[key]; ok {
		return v.All()
	}
	return []string{}
}

// Command : SetDel key
// 删除指定集合
func SetDel(key string) int {
	if _, ok := SetData[key]; ok {
		delete(SetData, key)
		Key.Remove(key)
		return 1
	}
	return 0
}
func SetDelAt(key string) int {
	rse := SetDel(key)
	if rse == 1 {
		raft.NewLogData("SetDel "+key).Write()
	}
	return SetDel(key)
}

// Command : SetDelValue key value
// 删除指定集合的元素
func SetDelValue(key, value string) int {
	if s, ok := SetData[key]; ok {
		s.Delete(value)
		return 1
	}
	return 0
}
func SetDelValueAt(key, value string) int {
	rse := SetDelValue(key, value)
	if rse == 1 {
		raft.NewLogData("SetDelValue "+key+" "+value).Write()
	}
	return rse
}

type set struct {
	data map[string]struct{}
	expire map[string]int64
}


func newSet() *set {
	return &set{
		data : make(map[string]struct{}),
		expire: make(map[string]int64),
	}
}

func (s set) Has(key string) bool {
	_, ok := s.data[key]
	return ok
}

func (s set) Add(key string) {
	s.data[key] = struct{}{}
	s.expire[key] = -1 // -1代表永久不过期
}

func (s set) AddExpire(key string, t int64) {
	s.data[key] = struct{}{}
	s.expire[key] = t
}

func (s set) Delete(key string) {
	delete(s.data, key)
	delete(s.expire, key)
}

// 过期的不取
func (s set) All() []string {
	t := time.Now().Unix()
	rse := make([]string, 0)
	for k, _ := range s.data {
		exp := s.expire[k]
		if exp == -1 || exp - t > 0 {
			rse = append(rse, k)
		}else{
			// 删除过期
			s.Delete(k)
		}
	}
	return rse
}

func (s set) Expire(key string, t int64) int {
	if _, ok := s.expire[key]; ok {
		s.expire[key] = t
		return 1
	}
	return 0
}