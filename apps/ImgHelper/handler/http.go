package handler

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/ImgHelper/service"
	"github.com/mangenotwork/extras/common/utils"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm img helper.\n"+utils.Logo))
}

// 生成二维码  QRCode
func QRCode(w http.ResponseWriter, r *http.Request) {
	value := utils.GetUrlArg(r, "value")
	if len(value) < 1 {
		utils.OutErrBody(w, 1001, fmt.Errorf("没有输入内容"))
		return
	}
	img, err := service.QrCodeBase64(value)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	w.Header().Set("Content-Typee", "image/png")
	_,_=w.Write(img)
}

// 生成条形码  Barcode
func Barcode(w http.ResponseWriter, r *http.Request) {
	value := utils.GetUrlArg(r, "value")
	if len(value) < 1 {
		utils.OutErrBody(w, 1001, fmt.Errorf("没有输入内容"))
		return
	}
	img, err := service.Barcode(value)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	w.Header().Set("Content-Typee", "image/png")
	_,_=w.Write(img)
}

// 识别二维码  QRCodeRecognition
func QRCodeRecognition(w http.ResponseWriter, r *http.Request) {
	log.Println(r)

}

// 识别条形码  BarcodeRecognition
func BarcodeRecognition(w http.ResponseWriter, r *http.Request) {

}

// 添加水印 - 文字水印   WatermarkTxt
func WatermarkTxt(w http.ResponseWriter, r *http.Request) {

}

// 添加水印 - 图片水印   WatermarkImg
func WatermarkImg(w http.ResponseWriter, r *http.Request) {

}