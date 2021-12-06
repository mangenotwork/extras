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

	return mux
}