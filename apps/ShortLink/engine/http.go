package engine

import (
	"github.com/mangenotwork/extras/apps/ShortLink/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartHttp(){
	go func() {
		logger.Info("StartHttp")
		mux := httpser.NewEngine()

		mux.Router("/err", handler.Error)
		mux.Router("/NotPrivacy", handler.NotPrivacy)
		mux.Router("/WhiteNote", handler.WhiteNote)
		mux.Router("/BlockNote", handler.BlockNote)

		mux.Router("/ttttt", handler.Te)

		// [post] /v1/add  创建短链接
		mux.Router("/v1/add", handler.Add)

		// [post] /v1/modify  修改短链接
		mux.Router("/v1/modify", handler.Modify)

		// [post] /v1/get   获取短链接信息
		mux.Router("/v1/get", handler.Get)

		// [post] /v1/del   删除短链接
		mux.Router("/v1/del", handler.Del)

		mux.Run()

	}()
}

