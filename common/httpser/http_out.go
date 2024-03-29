package httpser

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
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
	BodyPDF 			= "application/pdf"
	BodyJPG   			= "image/jpeg"
	BodyPNG	 			= "image/png"
	BodyGif				= "image/gif"
	BodyWord			= "application/msword"
	BodyOctet			= "application/octet-stream"

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

// 输出静态文件
func OutStaticFile(w http.ResponseWriter, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
		_,_=fmt.Fprintln(w, err)
		return
	}
	w.Header().Add("Content-Type", BodyHTML)
	_,_=fmt.Fprintln(w, string(data))
	return
}

func OutPdf(w http.ResponseWriter, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
		_,_=fmt.Fprintln(w, err)
		return
	}
	w.Header().Add("Content-Type", BodyPDF)
	_,_=fmt.Fprintln(w, string(data))
	return
}

func OutJPG(w http.ResponseWriter, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		w.WriteHeader(404)
		_,_=fmt.Fprintln(w, err)
		return
	}
	w.Header().Add("Content-Type", BodyJPG)
	_,_=fmt.Fprintln(w, string(data))
	return
}

// 给客户端下载的静态文件
func OutUploadFile(w http.ResponseWriter, path, fileName string) {
	file, _ := os.Open(path)
	defer file.Close()

	fileHeader := make([]byte, 512)
	_,err := file.Read(fileHeader)
	if err != nil {
		w.WriteHeader(404)
		_,_=fmt.Fprintln(w, err)
		return
	}
	fileStat, _ := file.Stat()

	w.Header().Set("Content-Disposition", "attachment; filename=" + fileName)
	w.Header().Set("Content-Type", http.DetectContentType(fileHeader))
	w.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))
	_,err = file.Seek(0, 0)
	if err != nil {
		w.WriteHeader(404)
		_,_=fmt.Fprintln(w, err)
		return
	}
	_,err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(404)
		_,_=fmt.Fprintln(w, err)
		return
	}

}

func Out404(w http.ResponseWriter) {
	w.WriteHeader(404)
	_,_=fmt.Fprintln(w, "404")
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
	return utils.Str2Int64(arg)
}

func GetUrlArgInt(r *http.Request, name string) int {
	var arg string
	values := r.URL.Query()
	arg=values.Get(name)
	return utils.Str2Int(arg)
}

func GetJsonParam(r *http.Request, param interface{}) {
	decoder:=json.NewDecoder(r.Body)
	_=decoder.Decode(&param)
}

func GetFromArg(r *http.Request, name string) string {
	return r.FormValue(name)
}

func GetFromFile(r *http.Request, name string) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(name)
}

func GetCookie(r *http.Request, name string) (*http.Cookie, error) {
	return r.Cookie(name)
}

func GetCookieVal(r *http.Request, name string) string {
	cookie, err := GetCookie(r, name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func SetCookie(w http.ResponseWriter, name, value string, t int) {
	http.SetCookie(w, &http.Cookie{
		Name:    name,
		Value:   url.QueryEscape(value),
		Expires: time.Now().Add(time.Duration(t) * time.Second),
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
}

func SetCookieMap(w http.ResponseWriter, data map[string]string, t int) {
	for k, v := range data {
		SetCookie(w, k, v, t)
	}
}

func GetClientIp(r *http.Request) string {
	return GetIP(r)
}

func GetHeader(r *http.Request, name string) string {
	return r.Header.Get(name)
}

func GetIp(r *http.Request) (ip string) {
		for _, ip := range strings.Split(r.Header.Get("X-Forward-For"), ",") {
			if net.ParseIP(ip) != nil {
				return ip
			}
		}
		if ip = r.Header.Get("X-Real-IP"); net.ParseIP(ip) != nil {
			return ip
		}
		if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
			if net.ParseIP(ip) != nil {
				return ip
			}
		}
		return "0.0.0.0"
}

