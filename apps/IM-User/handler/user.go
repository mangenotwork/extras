package handler

import (
	"github.com/mangenotwork/extras/apps/IM-User/service"
	"github.com/mangenotwork/extras/common/httpser"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	params := &service.UserParam{}
	httpser.GetJsonParam(r, params)
	// 执行注册业务
	params.Register()

	httpser.OutSucceedBodyJsonP(w, params)
	return
}
