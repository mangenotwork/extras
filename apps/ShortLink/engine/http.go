package engine

import (
	"github.com/mangenotwork/extras/apps/ShortLink/handler"
	"github.com/mangenotwork/extras/common/middleware"
	"github.com/mangenotwork/extras/common/utils"
	"net/http"
)

func StartHttpServer(){
	go func(){
		utils.HttpServer(Router())
	}()
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	m := middleware.Base
	mux.Handle("/", m(http.HandlerFunc(handler.Hello)))

	mux.Handle("/err", m(http.HandlerFunc(handler.Error)))
	mux.Handle("/NotPrivacy", m(http.HandlerFunc(handler.NotPrivacy)))
	mux.Handle("/WhiteNote", m(http.HandlerFunc(handler.WhiteNote)))
	mux.Handle("/BlockNote", m(http.HandlerFunc(handler.BlockNote)))

	mux.Handle("/ttttt", m(http.HandlerFunc(handler.Te)))

	// [post] /v1/add  创建短链接
	mux.Handle("/v1/add", m(http.HandlerFunc(handler.Add)))

	// [post] /v1/modify  修改短链接
	mux.Handle("/v1/modify", m(http.HandlerFunc(handler.Modify)))

	// [post] /v1/get   获取短链接信息
	mux.Handle("/v1/get", m(http.HandlerFunc(handler.Get)))

	// [post] /v1/del   删除短链接
	mux.Handle("/v1/del", m(http.HandlerFunc(handler.Del)))

	return mux
}