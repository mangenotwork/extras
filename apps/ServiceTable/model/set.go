package model

import (
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
	}
}

// Command : SetAddExpire key value timeUnix
// 集合添加数据并指定过期时间
func SetAddExpire(key, value string, timeUnix int64) {
	// 不存在就创建
	if _,ok := SetData[key]; !ok {
		SetData[key] = newSet()
	}
	SetData[key].AddExpire(value, timeUnix)
}

// Command : SetValueExpire key value timeUnix
// 指定集合数据过期时间
func SetValueExpire(key, value string, timeUnix int64) int {
	// 不存在就创建
	if _,ok := SetData[key]; !ok {
		return 0
	}
	return SetData[key].Expire(value, timeUnix)
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
		return 1
	}
	return 0
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
	rse := make([]string, len(s.data))
	for k, _ := range s.data {
		exp := s.expire[k]
		if exp == -1 || exp - t > 0 {
			rse = append(rse, k)
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