package handler

import (
	"github.com/mangenotwork/extras/apps/MinioClient/model"
	"github.com/mangenotwork/extras/apps/MinioClient/service"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
	"net/http"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	logger.Infof("Received request %s %s %s\n", r.Method, r.Host, r.RemoteAddr)
	logger.Info(r.URL)
	logger.Info(r.URL.Path,  r.URL.User, r.URL.Query())

	objUrlList := strings.Split(r.URL.Path, "/")
	if len(objUrlList) < 1 {
		http.Redirect(w, r, "/err", http.StatusMovedPermanently)
		return
	}

	obj := objUrlList[len(objUrlList)-1]
	bucket := strings.Replace(r.URL.Path, "/"+obj, "", -1)
	compact := r.URL.Query().Get("compact")

	service.GetFile(w, bucket, obj, compact)

}

func Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_,_=w.Write([]byte("Error: 未知链接!"))
}

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




