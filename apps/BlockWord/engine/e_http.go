package engine

import (
	"github.com/mangenotwork/extras/apps/BlockWord/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)


func StartHttp(){
	go func() {
		logger.Info("StartHttp")
		mux := httpser.NewEngine()

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
		mux.Router("/v1/do", handler.Do)

		// [GET] /v1/add 添加屏蔽词
		mux.Router("/v1/add", handler.Add)

		// [GET] /v1/del 删除屏蔽词
		mux.Router("/v1/del", handler.Del)

		// [GET] /v1/list 查看所有屏蔽词
		mux.Router("/v1/list", handler.List)

		// [GET] /v1/white/add 词语白名单添加
		mux.Router("/v1/white/add", handler.WhiteAdd)

		// [GET] /v1/white/del 词语白名单删除
		mux.Router("/v1/white/del", handler.WhiteDel)

		// [GET] /v1/white/list 查看所有词语白名单
		mux.Router("/v1/white/list", handler.WhiteList)

		// [POST] /v1/ishave 是否存在非法词语
		mux.Router("/v1/ishave", handler.IsHave)

		// [POST] /v1/ishave/list 是否存在非法词语并返回非法的词语
		mux.Router("/v1/ishave/list", handler.IsHaveList)

		mux.Run()

	}()
}