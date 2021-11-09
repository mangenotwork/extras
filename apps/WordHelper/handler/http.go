package handler

import (
	"github.com/mangenotwork/extras/apps/WordHelper/service"
	"github.com/mangenotwork/extras/common/utils"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm img helper.\n"+utils.Logo))
}

func JieBaFenCi(w http.ResponseWriter, r *http.Request) {
	str := utils.GetUrlArg(r, "str")
	jiebaType := utils.GetUrlArgInt(r, "type")
	data := service.JieBa(str, jiebaType)
	utils.OutSucceedBody(w, data)
}