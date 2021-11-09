package handler

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/WordHelper/service"
	"github.com/mangenotwork/extras/apps/WordHelper/service/pdf"
	"github.com/mangenotwork/extras/common/utils"
	"io/ioutil"
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
	defer file.Close()
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

/*
 http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i=计算
	ZH_CN2EN 中文　»　英语
	ZH_CN2JA 中文　»　日语
	ZH_CN2KR 中文　»　韩语
	ZH_CN2FR 中文　»　法语
	ZH_CN2RU 中文　»　俄语
	ZH_CN2SP 中文　»　西语
	EN2ZH_CN 英语　»　中文
	JA2ZH_CN 日语　»　中文
	KR2ZH_CN 韩语　»　中文
	FR2ZH_CN 法语　»　中文
	RU2ZH_CN 俄语　»　中文
	SP2ZH_CN 西语　»　中文
*/
func FanYi(w http.ResponseWriter, r *http.Request) {
	word := utils.GetUrlArg(r, "word")

	res, err :=http.Get("http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i="+word)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	robots, err := ioutil.ReadAll(res.Body)
	_=res.Body.Close()
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	w.Write(robots)
}

type PDFExtractionBody struct {
	Page int `json:"page"`
	Content interface{} `json:"content"`
}

func PDFExtractionTxt(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	reader, err := pdf.NewReader(file, handler.Size)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	data := make([]*PDFExtractionBody, 0, reader.NumPage())
	for i:=1; i< reader.NumPage(); i++ {
		pg := reader.Page(i)
		txt,_ := pg.GetTxt()
		data = append(data, &PDFExtractionBody{
			Page: i,
			Content: txt,
		})
	}
	utils.OutSucceedBody(w, data)
}

func PDFExtractionRow(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	reader, err := pdf.NewReader(file, handler.Size)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	data := make([]*PDFExtractionBody, 0, reader.NumPage())
	for i:=1; i< reader.NumPage(); i++ {
		pg := reader.Page(i)
		row, _ := pg.GetRow()
		data = append(data, &PDFExtractionBody{
			Page: i,
			Content: row,
		})
	}
	utils.OutSucceedBody(w, data)
}

func PDFExtractionTable(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	reader, err := pdf.NewReader(file, handler.Size)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	data := make([]*PDFExtractionBody, 0, reader.NumPage())
	for i:=1; i< reader.NumPage(); i++ {
		pg := reader.Page(i)
		data = append(data, &PDFExtractionBody{
			Page: i,
			Content: pg.GetTable(),
		})
	}
	utils.OutSucceedBody(w, data)
}

