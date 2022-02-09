package handler

import (
	"github.com/mangenotwork/extras/apps/MinioClient/model"
	"github.com/mangenotwork/extras/apps/MinioClient/service"
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

type BucketPostParam struct {
	Name string `json:"name"`
}

func BucketAdd(w http.ResponseWriter, r *http.Request) {
	params := &BucketPostParam{}
	httpser.GetJsonParam(r, params)
	err := service.MakeBucket(params.Name)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBodyJsonP(w, "创建成功!")
	return
}

func BucketList(w http.ResponseWriter, r *http.Request) {
	rse, err := service.BucketList()
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

func BucketFiles(w http.ResponseWriter, r *http.Request) {
	bucket := httpser.GetUrlArg(r, "bucket")
	rse := service.BucketFiles(bucket)
	httpser.OutSucceedBodyJsonP(w, rse)
	return
}




