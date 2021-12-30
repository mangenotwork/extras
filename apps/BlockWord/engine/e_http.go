package engine

import (
	"github.com/mangenotwork/extras/apps/BlockWord/handler"
	"github.com/mangenotwork/extras/common/middleware"
	"github.com/mangenotwork/extras/common/httpser"
	"net/http"
)


func StartHttpServer(){
	go func() {
		httpser.HttpServer(Router())
	}()
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	m := middleware.Base
	mux.Handle("/hello", m(http.HandlerFunc(handler.Hello)))
	mux.Handle("/", m(http.HandlerFunc(handler.Hello)))

	// swagger:operation POST /v1/do 屏蔽词过滤
	// ---
	// summary: 屏蔽词过滤
	// description: 屏蔽词过滤
	// parameters:
	// - name: str
	//   in: body
	//   description: 词语
	//   type: string
	//   required: true
	// - name: sub
	//   in: body
	//   description: 替换符号
	//   type: string
	//   required: true
	// responses:
	//   200: json: {"code":0,"timestamp":1635232884,"msg":"succeed","data":{"str":"我在路口交通进行???就在这个路口交接","sub":"???"}}
	mux.Handle("/v1/do", m(http.HandlerFunc(handler.Do)))

	// [GET] /v1/add 添加屏蔽词
	mux.Handle("/v1/add", m(http.HandlerFunc(handler.Add)))

	// [GET] /v1/del 删除屏蔽词
	mux.Handle("/v1/del", m(http.HandlerFunc(handler.Del)))

	// [GET] /v1/list 查看所有屏蔽词
	mux.Handle("/v1/list", m(http.HandlerFunc(handler.List)))

	// [GET] /v1/white/add 词语白名单添加
	mux.Handle("/v1/white/add", m(http.HandlerFunc(handler.WhiteAdd)))

	// [GET] /v1/white/del 词语白名单删除
	mux.Handle("/v1/white/del", m(http.HandlerFunc(handler.WhiteDel)))

	// [GET] /v1/white/list 查看所有词语白名单
	mux.Handle("/v1/white/list", m(http.HandlerFunc(handler.WhiteList)))

	// [POST] /v1/ishave 是否存在非法词语
	mux.Handle("/v1/ishave", m(http.HandlerFunc(handler.IsHave)))

	// [POST] /v1/ishave/list 是否存在非法词语并返回非法的词语
	mux.Handle("/v1/ishave/list", m(http.HandlerFunc(handler.IsHaveList)))

	return mux
}