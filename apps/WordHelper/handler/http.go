package handler

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/mangenotwork/extras/apps/WordHelper/service"
	"github.com/mangenotwork/extras/apps/WordHelper/service/pdf"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm img helper.\n"+utils.Logo))
}

func JieBaFenCi(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	jiebaType := httpser.GetUrlArgInt(r, "type")
	data := service.JieBa(str, jiebaType)
	httpser.OutSucceedBody(w, data)
}

func OCR(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	defer file.Close()
	lang := r.FormValue("lang")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
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
		httpser.OutErrBody(w, 2001, err)
		return
	}
	logger.Info(handler.Size)
	httpser.OutSucceedBody(w, jg)
}

func GetOCRLanguages(w http.ResponseWriter, r *http.Request) {
	lang, err := service.GetOCRLanguages()
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, lang)
}

func GetOCRVersion(w http.ResponseWriter, r *http.Request) {
	httpser.OutSucceedBody(w, service.GetOCRVersion())
}

// OCRBase64Img 识别base64 的图片
func OCRBase64Img(w http.ResponseWriter, r *http.Request) {
	lang := r.FormValue("lang")
	base64img := r.FormValue("base64img")

	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64img)
	if b {
		re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
		//allData := re.FindAllSubmatch([]byte(base64img), 2)
		//log.Print(allData)
		base64img = re.ReplaceAllString(base64img, "")
	}
	logger.Info(base64img)
	// Base64 Standard Decoding
	sDec, err := base64.StdEncoding.DecodeString(base64img)
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return
	}

	sDec = service.Huihua(sDec)

	file := "./a.png"
	err = ioutil.WriteFile(file, sDec, 0666)
	if err != nil {
		logger.Error(err)
	}

	jg, err := service.OCR(sDec, lang)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, jg)
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
	word := httpser.GetUrlArg(r, "word")

	res, err :=http.Get("http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i="+word)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	robots, err := ioutil.ReadAll(res.Body)
	_=res.Body.Close()
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(robots)
}

type PDFExtractionBody struct {
	Page int `json:"page"`
	Content interface{} `json:"content"`
}

func PDFExtractionTxt(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	reader, err := pdf.NewReader(file, handler.Size)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
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
	httpser.OutSucceedBody(w, data)
}

func PDFExtractionRow(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	reader, err := pdf.NewReader(file, handler.Size)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
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
	httpser.OutSucceedBody(w, data)
}

func PDFExtractionTable(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	reader, err := pdf.NewReader(file, handler.Size)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
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
	httpser.OutSucceedBody(w, data)
}

type EncryptParam struct {
	Str string `json:"str"`
	Key string `json:"key"`
	Iv string `json:"iv"`
}

type DecryptParam struct {
	Str string `json:"str"`
	Key string `json:"key"`
	Iv string `json:"iv"`
}

// AES
func AESCBCEncrypt(w http.ResponseWriter, r *http.Request) {
	params := &EncryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewAES("cbc", []byte(params.Iv)).Encrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, base64.StdEncoding.EncodeToString(rse))
}

func AESCBCDecrypt(w http.ResponseWriter, r *http.Request) {
	params := &DecryptParam{}
	httpser.GetJsonParam(r, params)
	decoded, err := base64.StdEncoding.DecodeString(params.Str)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	rse, err := service.NewAES("cbc", []byte(params.Iv)).Decrypt(decoded, []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func AESECBEncrypt(w http.ResponseWriter, r *http.Request) {
	params := &EncryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewAES("ecb").Encrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, base64.StdEncoding.EncodeToString(rse))
}

func AESECBDecrypt(w http.ResponseWriter, r *http.Request) {
	params := &DecryptParam{}
	httpser.GetJsonParam(r, params)
	decoded, err := base64.StdEncoding.DecodeString(params.Str)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	rse, err := service.NewAES("ecb").Decrypt(decoded, []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func AESCFBEncrypt(w http.ResponseWriter, r *http.Request) {
	params := &EncryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewAES("cfb").Encrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, base64.StdEncoding.EncodeToString(rse))
}

func AESCFBDecrypt(w http.ResponseWriter, r *http.Request) {
	params := &DecryptParam{}
	httpser.GetJsonParam(r, params)
	decoded, err := base64.StdEncoding.DecodeString(params.Str)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	rse, err := service.NewAES("cfb").Decrypt(decoded, []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func AESCTREncrypt(w http.ResponseWriter, r *http.Request) {
	params := &EncryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewAES("ctr", []byte(params.Iv)).Encrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func AESCTRDecrypt(w http.ResponseWriter, r *http.Request) {
	params := &DecryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewAES("ctr", []byte(params.Iv)).Decrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

// DES
func DESCBCEncrypt(w http.ResponseWriter, r *http.Request) {
	params := &EncryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewDES("cbc", []byte(params.Iv)).Encrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func DESCBCDecrypt(w http.ResponseWriter, r *http.Request) {
	params := &DecryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewDES("cbc", []byte(params.Iv)).Decrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func DESECBEncrypt(w http.ResponseWriter, r *http.Request) {
	params := &EncryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewDES("ecb").Encrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func DESECBDecrypt(w http.ResponseWriter, r *http.Request) {
	params := &DecryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewDES("ecb").Decrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func DESCFBEncrypt(w http.ResponseWriter, r *http.Request) {
	params := &EncryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewDES("cfb").Encrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func DESCFBDecrypt(w http.ResponseWriter, r *http.Request) {
	params := &DecryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewDES("cfb").Decrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func DESCTREncrypt(w http.ResponseWriter, r *http.Request) {
	params := &EncryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewDES("ctr", []byte(params.Iv)).Encrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func DESCTRDecrypt(w http.ResponseWriter, r *http.Request) {
	params := &DecryptParam{}
	httpser.GetJsonParam(r, params)
	rse, err := service.NewDES("ctr", []byte(params.Iv)).Decrypt([]byte(params.Str), []byte(params.Key))
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(rse))
}

func MD516(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	h := md5.New()
	h.Write([]byte(str))
	httpser.OutSucceedBody(w, hex.EncodeToString(h.Sum(nil))[8:24])
}

func MD532(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	h := md5.New()
	h.Write([]byte(str))
	httpser.OutSucceedBody(w, hex.EncodeToString(h.Sum(nil)))
}

func Base64Encrypt(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	httpser.OutSucceedBody(w, base64.StdEncoding.EncodeToString([]byte(str)))
}

func Base64Decrypt(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(decoded))
}

func Base64UrlEncrypt(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	httpser.OutSucceedBody(w, base64.URLEncoding.EncodeToString([]byte(str)))
}

func Base64UrlDecrypt(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	decoded, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	httpser.OutSucceedBody(w, string(decoded))
}

func HmacMD5(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	key := httpser.GetUrlArg(r, "key")
	httpser.OutSucceedBody(w, service.HmacMD5(str,key))
}

func HmacSHA1(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	key := httpser.GetUrlArg(r, "key")
	httpser.OutSucceedBody(w, service.HmacSHA1(str,key))
}

func HmacSHA256(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	key := httpser.GetUrlArg(r, "key")
	httpser.OutSucceedBody(w, service.HmacSHA256(str,key))
}

func HmacSHA512(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	key := httpser.GetUrlArg(r, "key")
	httpser.OutSucceedBody(w, service.HmacSHA512(str,key))
}

func Md2Html(w http.ResponseWriter, r *http.Request) {
	str := httpser.GetUrlArg(r, "str")
	httpser.OutSucceedBody(w, service.MarkdownToHTML(str))
}
