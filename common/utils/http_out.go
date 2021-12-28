package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mangenotwork/extras/common/logger"
)

// HTTP  输出 json body 定义
// [Code]
// - 0 成功
// - 1001 参数错误
// - 2001 程序错误
type HttpOutBody struct {
	Code int `json:"code"`
	Timestamp int64 `json:"timestamp"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
const (
	BodyJSON            = "application/json; charset=utf-8"
	BodyAsciiJSON       = "application/json"
	BodyHTML            = "text/html; charset=utf-8"
	BodyJavaScript      = "application/javascript; charset=utf-8"
	BodyXML             = "application/xml; charset=utf-8"
	BodyPlain           = "text/plain; charset=utf-8"
	BodyYAML            = "application/x-yaml; charset=utf-8"
	BodyDownload        = "application/octet-stream; charset=utf-8"
)

func OutSucceedBodyJsonP(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", BodyJavaScript)
	body := &HttpOutBody{
		Code: 0,
		Timestamp: time.Now().Unix(),
		Msg: "succeed",
		Data: data,
	}
	bodyJson, err := body.JsonStr()
	if err != nil {
		OutErrBody(w,2001, err)
	}
	_,_=fmt.Fprintln(w, bodyJson)
	return
}

func OutSucceedBody(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", BodyJSON)
	body := &HttpOutBody{
		Code: 0,
		Timestamp: time.Now().Unix(),
		Msg: "succeed",
		Data: data,
	}
	bodyJson, err := body.JsonStr()
	if err != nil {
		OutErrBody(w,2001, err)
	}
	_,_=fmt.Fprintln(w, bodyJson)
	return
}

func OutErrBody(w http.ResponseWriter, code int,err error) {
	body := &HttpOutBody{
		Code: code,
		Timestamp: time.Now().Unix(),
		Msg: err.Error(),
		Data: nil,
	}
	bodyJson, _ := body.JsonStr()
	_,_=fmt.Fprintln(w, bodyJson)
	return
}

func (m *HttpOutBody) JsonStr() (string,error) {
	b, err := json.Marshal(m)
	if err != nil {
		logger.Error("Umarshal failed:", err)
		return "",err
	}
	return string(b), nil
}

// GetUrlArg 获取URL的GET参数
func GetUrlArg(r *http.Request, name string) string {
	var arg string
	values := r.URL.Query()
	arg=values.Get(name)
	return arg
}

func GetUrlArgInt64(r *http.Request, name string) int64 {
	var arg string
	values := r.URL.Query()
	arg=values.Get(name)
	return Str2Int64(arg)
}

func GetUrlArgInt(r *http.Request, name string) int {
	var arg string
	values := r.URL.Query()
	arg=values.Get(name)
	return Str2Int(arg)
}