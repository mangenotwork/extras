package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

func OutSucceedBody(w http.ResponseWriter, data interface{}) {
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
		log.Println("Umarshal failed:", err)
		return "",err
	}
	return string(b), nil
}

// GetUrlArg 获取URL的GET参数
func GetUrlArg(r *http.Request,name string)string{
	var arg string
	values := r.URL.Query()
	arg=values.Get(name)
	return arg
}