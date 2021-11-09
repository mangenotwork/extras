package handler

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/WordHelper/service"
	"github.com/mangenotwork/extras/common/utils"
	"log"
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

func OCR(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	lang := r.FormValue("lang")
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	var buff = make([]byte, handler.Size)
	n, err := file.Read(buff)
	if err != nil {
		fmt.Println(err)
		return
	}
	jg, err := service.OCR(buff[:n], lang)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	log.Println(handler.Size)
	utils.OutSucceedBody(w, jg)
}

func GetOCRLanguages(w http.ResponseWriter, r *http.Request) {
	lang, err := service.GetOCRLanguages()
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	utils.OutSucceedBody(w, lang)
}

func GetOCRVersion(w http.ResponseWriter, r *http.Request) {
	utils.OutSucceedBody(w, service.GetOCRVersion())
}