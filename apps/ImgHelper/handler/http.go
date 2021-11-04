package handler

import (
	"github.com/mangenotwork/extras/common/utils"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm img helper.\n"+utils.Logo))
}

// 生成二维码  QRCode
func QRCode(w http.ResponseWriter, r *http.Request) {

}

// 生成条形码  Barcode
func Barcode(w http.ResponseWriter, r *http.Request) {

}

// 识别二维码  QRCodeRecognition
func QRCodeRecognition(w http.ResponseWriter, r *http.Request) {

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