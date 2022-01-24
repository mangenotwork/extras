package handler

import (
	"github.com/mangenotwork/extras/apps/IM-User/service"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/jwt"
	"github.com/mangenotwork/extras/common/logger"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	params := &service.UserParam{}
	httpser.GetJsonParam(r, params)
	// 执行注册业务
	rse := params.Register()

	httpser.OutSucceedBodyJsonP(w, rse)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	params := &service.UserParam{}
	httpser.GetJsonParam(r, params)
	token, err := params.Token()
	if err != nil {
		httpser.OutErrBody(w, 2000, err)
		return
	}
	httpser.OutSucceedBodyJsonP(w, token)
}


func Jwt(w http.ResponseWriter, r *http.Request) {
	j := jwt.NewJWT()
	j.AddClaims("uid", 100)
	j.AddClaims("name", "aaaa")
	j.AddClaims("isok", true)
	token, err := j.Token()
	if err != nil {
		httpser.OutErrBody(w, 404, err)
		return
	}
	httpser.OutSucceedBodyJsonP(w, token)
}

func JwtGet(w http.ResponseWriter, r *http.Request) {
	token := httpser.GetUrlArg(r, "token")
	j := jwt.NewJWT()
	err := j.ParseToken(token)
	if err != nil {
		httpser.OutErrBody(w, 404, err)
		return
	}

	logger.Debug("uid = ", j.Get("uid"))
	logger.Debug("uid = ", j.GetInt("uid"))
	logger.Debug("name = ", j.GetString("name"))
	logger.Debug("isok = ", j.GetBool("isok"))

	logger.Debug("is expire = ", j.IsExpire())

	logger.Debug(" Expire = ", j.GetInt64("Expire"))
	logger.Debug("Expire = ", j.Get("Expire"))

	httpser.OutSucceedBodyJsonP(w, token)
}