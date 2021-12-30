package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mangenotwork/extras/apps/ShortLink/model"
	"github.com/mangenotwork/extras/apps/ShortLink/service"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/middleware"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	logger.Info("Received request %s %s %s\n", r.Method, r.Host, r.RemoteAddr)
	logger.Info(r.URL)
	logger.Info(r.URL.Path,  r.URL.User, r.URL.Query())
	// 获取短链接
	link := new(model.ShortLink)
	err := link.Get(r.URL.Path)
	if err != nil || len(link.Url) < 1 {
		http.Redirect(w, r, "/err", http.StatusMovedPermanently)
		return
	}
	if link.IsPrivacy {
		if httpser.GetUrlArg(r, "password") != link.Password {
			http.Redirect(w, r, "/NotPrivacy", http.StatusMovedPermanently)
			return
		}
	}
	ip := middleware.GetIP(r)
	if !link.IsWhiteList(ip){
		http.Redirect(w, r, "/WhiteNote", http.StatusMovedPermanently)
		return
	}
	if link.IsBlockList(ip) {
		http.Redirect(w, r, "/BlockNote", http.StatusMovedPermanently)
		return
	}
	http.Redirect(w, r, link.Url, http.StatusMovedPermanently)
}

// 隐藏静态页面， 如果是动态页面由于隐藏了host无法实现跨域请求
func Te(w http.ResponseWriter, r *http.Request) {
	logger.Info(r)

	transport :=  http.DefaultTransport

	// step 1
	outReq := new(http.Request)
	*outReq = *r // this only does shallow copies of maps

	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	newUrl,err := url.Parse("https://studygolang.com/articles/6340")
	logger.Info("err = ", err)
	outReq.URL = newUrl
	outReq.Host = newUrl.Host

	logger.Info("outReq.URL = ", outReq)

	// step 2
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}


	// step 3
	for key, value := range res.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}

	w.WriteHeader(res.StatusCode)
	io.Copy(w, res.Body)
	res.Body.Close()
}


func Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_,_=w.Write([]byte("Error: 未知链接!"))
}

func NotPrivacy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	_,_=w.Write([]byte("Error: 访问密码错误! 请设置密码参数如: /url?password="))
}

func WhiteNote(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	_,_=w.Write([]byte("对不起: 只有在白名单内的地址访问."))
}

func BlockNote(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	_,_=w.Write([]byte("对不起: 你再在黑名单内,禁止访问."))
}

type AddParam struct {
	Url string `json:"url"` // 目的地址
	Aging int64 `json:"aging"` // 时效，单位秒
	Deadline int64 `json:"deadline"` // 截止日期， 单位时间戳, 只有当Aging为0时才用
	IsPrivacy bool `json:"is_privacy"` // 是否隐私
	Password string `json:"password"` // 只有当IsPrivacy=true使用
	OpenBlockList bool `json:"open_block_list"` // 是否启用黑名单，启用后黑名单不能访问
	BlockList []string `json:"block_list"` // 访问黑名单， OpenBlockList=true使用
	OpenWhiteList bool `json:"open_white_list"` // 是否启用白名单，启用后只有白名单才能访问
	WhiteList []string `json:"white_list"` // 访问白名单， OpenWhiteList=true使用
}

type AddBody struct {
	Url string `json:"url"` // 短链接地址
	Password string `json:"password"` // 短链接访问密码，空则没有密码
	Expire string `json:"expire"` // 过期时间， 空则永久不过期
}

// 创建短链接
func Add(w http.ResponseWriter, r *http.Request) {
	decoder:=json.NewDecoder(r.Body)
	params := &AddParam{}
	_=decoder.Decode(&params)
	if params.IsPrivacy && len(params.Password) < 1 {
		httpser.OutErrBody(w, 1001, errors.New("设置了隐私但是password为空"))
	}
	// 生成短链接
	exp := params.Aging
	if exp == 0 {
		exp = params.Deadline
	}
	shortLink := &model.ShortLink{
		Short: "/"+service.MustGenerate(),
		Url: params.Url,
		Expiration: exp,
		IsPrivacy : params.IsPrivacy,
		Password : params.Password,
		Creation : time.Now().Unix(),
		View : 0,
		OpenBlockList : params.OpenBlockList,
		OpenWhiteList : params.OpenWhiteList,
		BlockList: params.BlockList,
		WhiteList: params.WhiteList,
	}

	err := shortLink.Save()
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	sLink := &AddBody{
		Url: shortLink.Short,
		Password: params.Password,
		Expire: time.Unix(exp, 0).Format("2006-01-02 15:04:05"),
	}
	httpser.OutSucceedBody(w, sLink)
}

// 查看短链接, 如果是隐私的则需要带密码访问
func Get(w http.ResponseWriter, r *http.Request) {

}

// 修改短链接， 如果是隐私的则需要带密码
func Modify(w http.ResponseWriter, r *http.Request) {

}

// 删除短链接
func Del(w http.ResponseWriter, r *http.Request) {

}

