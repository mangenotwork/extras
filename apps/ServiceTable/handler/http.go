package handler

import (
	"net/http"
	"strings"

	"github.com/mangenotwork/extras/apps/ServiceTable/model"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/utils"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm img helper.\n"+utils.Logo))
}

// Command : SetAdd key value1,value2,
// 集合添加数据
func SetAdd(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	values := httpser.GetUrlArg(r, "values")

	model.SetAddAt(key, utils.SliceDelNullString(strings.Split(values, ",")))
	httpser.OutSucceedBodyJsonP(w, "")
	return
}

// Command : SetAddExpire key value timeUnix
// 集合添加数据并指定过期时间
func SetAddExpire(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	value := httpser.GetUrlArg(r, "value")
	timeUnix := httpser.GetUrlArgInt64(r, "time")
	model.SetAddExpireAt(key, value, timeUnix)
	httpser.OutSucceedBodyJsonP(w, "")
	return
}

// Command : SetValueExpire key value timeUnix
// 指定集合数据过期时间
func SetValueExpire(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	value := httpser.GetUrlArg(r, "value")
	timeUnix := httpser.GetUrlArgInt64(r, "time")
	rse := model.SetValueExpireAt(key, value, timeUnix)
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

// Command : SetGet key
// 获取集合所有数据
func SetGet(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	rse := model.SetGet(key)
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

// Command : SetDel key
// 删除指定集合
func SetDel(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	rse := model.SetDelAt(key)
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

// Command : SetDelValue key value
// 删除指定集合的元素
func SetDelValue(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	value := httpser.GetUrlArg(r, "value")
	rse := model.SetDelValueAt(key, value)
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

// key 是否存在
// Command : KeyHas key
func KeyHas(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	rse := model.Key.Has(key)
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

// key 模糊查询
// Command : KeyLike key
func KeyLike(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	rse := model.Key.Like(key)
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

// key 列表
// Command : KeyAll
func KeyAll(w http.ResponseWriter, r *http.Request) {
	rse := model.Key.GetAll()
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

// KV
// 增,改
// Command : KVAdd key value
func KVAdd(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	value := httpser.GetUrlArg(r, "value")
	model.KVAddAt(key, value)
	httpser.OutSucceedBodyJsonP(w, "")
	return
}

// Command : KVAddExpire key value expire
func KVAddExpire(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	value := httpser.GetUrlArg(r, "value")
	expire := httpser.GetUrlArgInt64(r, "expire")
	model.KVAddExpireAt(key, value, expire)
	httpser.OutSucceedBodyJsonP(w, "")
	return
}

// Command : KVExpire key expire
func KVExpire(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	expire := httpser.GetUrlArgInt64(r, "expire")
	rse := model.KVExpireAt(key, expire)
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

// 删
// Command : KVDel key
func KVDel(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	rse := model.KVDelAt(key)
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

// 查
// Command : KVGet key
func KVGet(w http.ResponseWriter, r *http.Request) {
	key := httpser.GetUrlArg(r, "key")
	is, rse := model.KVGet(key)
	if !is {
		httpser.OutSucceedBodyJsonP(w, "没有这个key的数据")
		return
	}
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}


