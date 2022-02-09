package handler

import (
	"github.com/mangenotwork/extras/apps/MinioClient/model"
	"github.com/mangenotwork/extras/common/httpser"
	"net/http"
)

func HasConn(w http.ResponseWriter, r *http.Request) {
	if model.MinioClient != nil {
		httpser.OutSucceedBodyJsonP(w, "连接成功!")
		return
	}
	httpser.OutSucceedBodyJsonP(w, "连接失败!")
	return
}