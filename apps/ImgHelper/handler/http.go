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
	"mime/multipart"
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
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	txt := r.FormValue("txt")
	if len(txt) < 1 {
		utils.OutErrBody(w, 2001, fmt.Errorf("txt is null"))
		return
	}
	color := r.FormValue("color")
	fontSize := r.FormValue("font_size")
	fontSizeInt := utils.Str2Int(fontSize)
	dpi := r.FormValue("dpi")
	dpiInt := utils.Str2Int(dpi)
	position := r.FormValue("position")
	positionInt := utils.Str2Int(position)
	out, err := service.WatermarkTxt(file, txt, color, fontSizeInt, dpiInt, positionInt)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

// 添加水印 - 图片水印   WatermarkLogo
func WatermarkLogo(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	logo, _, err := r.FormFile("logo")
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	out, err := service.WatermarkLogo(file, logo)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
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

func Txt2Img(w http.ResponseWriter, r *http.Request) {
	txt := r.FormValue("txt")
	fontSize := r.FormValue("font_size")
	fontSizeInt := utils.Str2Int(fontSize)
	dpi := r.FormValue("dpi")
	dpiInt := utils.Str2Int(dpi)
	spacing := r.FormValue("spacing")
	spacingInt := utils.Str2Int(spacing)
	outType := r.FormValue("out_type")
	out, err := service.Txt2Img(txt, fontSizeInt, dpiInt, spacingInt, outType)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func Img2Gif(w http.ResponseWriter, r *http.Request) {
	_= r.ParseMultipartForm(10 << 20)
	files := r.MultipartForm.File["file"]
	out, err := service.CompositeGif(files)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgRevolve(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	rType := r.FormValue("type")
	defer file.Close()
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	out, err := service.Revolve(file, rType)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgCenter(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	out, err := service.ImgCenter(file)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgStitching(w http.ResponseWriter, r *http.Request) {
	fileCount := r.FormValue("file_count")
	count := utils.Str2Int(fileCount)
	if count < 2 {
		utils.OutErrBody(w, 2001, fmt.Errorf("图片少于两张"))
		return
	}

	fileList := make([]multipart.File, 0, count)
	for i:=0; i<count; i++ {
		file, _, err := r.FormFile(fmt.Sprintf("file_%d", i+1))
		if err != nil {
			utils.OutErrBody(w, 2001, err)
			return
		}
		fileList = append(fileList, file)
		_=file.Close()
	}

	out, err := service.ImgStitching(fileList)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgSudoku(w http.ResponseWriter, r *http.Request) {
	fileCount := r.FormValue("file_count")
	count := utils.Str2Int(fileCount)
	if count < 2 {
		utils.OutErrBody(w, 2001, fmt.Errorf("图片少于两张"))
		return
	}

	fileList := make([]multipart.File, 0, count)
	for i:=0; i<count; i++ {
		file, _, err := r.FormFile(fmt.Sprintf("file_%d", i+1))
		if err != nil {
			utils.OutErrBody(w, 2001, err)
			return
		}
		fileList = append(fileList, file)
		_=file.Close()
	}

	out, err := service.ImgStitchingSudoku(fileList)
	if err != nil {
		utils.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}