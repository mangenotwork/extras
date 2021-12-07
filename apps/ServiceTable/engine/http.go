package engine

import (
	"github.com/mangenotwork/extras/apps/ServiceTable/handler"
	"github.com/mangenotwork/extras/common/middleware"
	"github.com/mangenotwork/extras/common/utils"
	"net/http"
)

func StartHttpServer(){
	go func() {
		utils.HttpServer(Router())
	}()
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	m := middleware.Base
	mux.Handle("/hello", m(http.HandlerFunc(handler.Hello)))
	mux.Handle("/", m(http.HandlerFunc(handler.Hello)))


	// Set
	// Command : SetAdd key value1,value2,
	// 集合添加数据
	mux.Handle("/set/add", m(http.HandlerFunc(handler.SetAdd)))
	// Command : SetAddExpire key value timeUnix
	// 集合添加数据并指定过期时间
	mux.Handle("/set/addExpire", m(http.HandlerFunc(handler.SetAddExpire)))
	// Command : SetValueExpire key value timeUnix
	// 指定集合数据过期时间
	mux.Handle("/set/valueExpire", m(http.HandlerFunc(handler.SetValueExpire)))
	// Command : SetGet key
	// 获取集合所有数据
	mux.Handle("/set/get", m(http.HandlerFunc(handler.SetGet)))
	// Command : SetDel key
	// 删除指定集合
	mux.Handle("/set/del", m(http.HandlerFunc(handler.SetDel)))
	// Command : SetDelValue key value
	// 删除指定集合的元素
	mux.Handle("/set/value", m(http.HandlerFunc(handler.SetDelValue)))


	// key 是否存在
	// Command : KeyHas key
	mux.Handle("/key/has", m(http.HandlerFunc(handler.KeyHas)))
	// key 模糊查询
	// Command : KeyLike key
	mux.Handle("/key/like", m(http.HandlerFunc(handler.KeyLike)))
	// key 列表
	// Command : KeyAll
	mux.Handle("/key/all", m(http.HandlerFunc(handler.KeyAll)))


	// KV
	// 增,改
	// Command : KVAdd key value
	mux.Handle("/kv/add", m(http.HandlerFunc(handler.KVAdd)))
	// Command : KVAddExpire key value expire
	mux.Handle("/kv/addExpire", m(http.HandlerFunc(handler.KVAddExpire)))
	// Command : KVExpire key expire
	mux.Handle("/kv/expire", m(http.HandlerFunc(handler.KVExpire)))
	// 删
	// Command : KVDel key
	mux.Handle("/kv/del", m(http.HandlerFunc(handler.KVDel)))
	// 查
	// Command : KVGet key
	mux.Handle("/kv/get", m(http.HandlerFunc(handler.KVGet)))


	return mux
}