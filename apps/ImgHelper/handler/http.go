package handler

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/ImgHelper/service"
	"github.com/mangenotwork/extras/common/utils"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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

type ImageInfoBody struct {
	Name string `json:"name"`
	Size int64 `json:"size"`
	Type string `json:"type"`
	Width int `json:"width"`
	Height int `json:"height"`
	Dpi string `json:"dpi"`
	IsEXIF bool `json:"is_exif"`
}

func ImageInfo(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	img, str, err := image.Decode(file)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	imgInfo := ImageInfoBody{
		Name: handler.Filename,
		Size: handler.Size,
		Type: str,
		Width: width,
		Height: height,
		Dpi: fmt.Sprintf("%d*%d dpi", width, height),
		IsEXIF: false,
	}

	if ok := service.NewExifData().ProcessExifStream(file); ok == nil {
		imgInfo.IsEXIF = true
	}

	utils.OutSucceedBody(w, imgInfo)
}

func ImageCompress(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	level := r.FormValue("level")
	defer file.Close()
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	img, str, err := image.Decode(file)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	b := img.Bounds()
	width := b.Max.X
	levelInt := utils.Str2Int(level)
	if levelInt <= 0 {
		levelInt = 1
	}
	out := service.ImgCompress(img, width/levelInt, 0, str)
	_,_=w.Write(out)
}