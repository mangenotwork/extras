package handler

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"mime/multipart"
	"net/http"
	"path"
	"strings"

	"github.com/mangenotwork/extras/apps/ImgHelper/service"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	_,_= w.Write([]byte("Hello l'm img helper.\n"+utils.Logo))
}

// 生成二维码  QRCode
func QRCode(w http.ResponseWriter, r *http.Request) {
	value := httpser.GetUrlArg(r, "value")
	if len(value) < 1 {
		httpser.OutErrBody(w, 1001, fmt.Errorf("没有输入内容"))
		return
	}
	img, err := service.QrCodeBase64(value)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	w.Header().Set("Content-Typee", "image/png")
	_,_=w.Write(img)
}

// 生成条形码  Barcode
func Barcode(w http.ResponseWriter, r *http.Request) {
	value := httpser.GetUrlArg(r, "value")
	if len(value) < 1 {
		httpser.OutErrBody(w, 1001, fmt.Errorf("没有输入内容"))
		return
	}
	img, err := service.Barcode(value)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	w.Header().Set("Content-Typee", "image/png")
	_,_=w.Write(img)
}

// 识别二维码  QRCodeRecognition
func QRCodeRecognition(w http.ResponseWriter, r *http.Request) {
	logger.Info(r)

}

// 识别条形码  BarcodeRecognition
func BarcodeRecognition(w http.ResponseWriter, r *http.Request) {

}

// 添加水印 - 文字水印   WatermarkTxt
func WatermarkTxt(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	txt := r.FormValue("txt")
	if len(txt) < 1 {
		httpser.OutErrBody(w, 2001, fmt.Errorf("txt is null"))
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
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

// 添加水印 - 图片水印   WatermarkLogo
func WatermarkLogo(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	logo, _, err := r.FormFile("logo")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	out, err := service.WatermarkLogo(file, logo)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
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
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	img, str, err := image.Decode(file)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
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

	httpser.OutSucceedBody(w, imgInfo)
}

func ImageCompress(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	level := r.FormValue("level")
	img, str, err := image.Decode(file)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	b := img.Bounds()
	width := b.Max.X
	levelInt := utils.Str2Int(level)
	if levelInt <= 0 {
		levelInt = 1
	}
	out := service.ImgCompress(img, width/levelInt, 0, str)


	//ext := path.Ext(head.Filename)
	//ext = strings.Replace(ext, ".", "", -1)
	//w.Header().Add("Content-Type", "image/"+ext)
	_,_=w.Write(out)
}

func ImageCompressBase64(w http.ResponseWriter, r *http.Request) {
	file, head, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	ext := path.Ext(head.Filename)
	ext = strings.Replace(ext, ".", "", -1)

	defer file.Close()
	level := r.FormValue("level")
	logger.Error(level)
	img, str, err := image.Decode(file)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	b := img.Bounds()
	width := b.Max.X
	levelInt := utils.Str2Int(level)
	if levelInt <= 0 {
		levelInt = 1
	}
	out := service.ImgCompress(img, width/levelInt, 0, str)

	encodeString := base64.StdEncoding.EncodeToString(out)

	logger.Error(encodeString)

	w.WriteHeader(200)

	_,_=w.Write([]byte("data:image/"+ext+";base64,"+encodeString))

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
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func Img2Gif(w http.ResponseWriter, r *http.Request) {
	_= r.ParseMultipartForm(10 << 20)
	files := r.MultipartForm.File["file"]
	out, err := service.CompositeGif(files)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgRevolve(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	rType := r.FormValue("type")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer func(){_=file.Close()}()
	out, err := service.Revolve(file, rType)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgCenter(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	out, err := service.ImgCenter(file)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgStitching(w http.ResponseWriter, r *http.Request) {
	fileCount := r.FormValue("file_count")
	count := utils.Str2Int(fileCount)
	if count < 2 {
		httpser.OutErrBody(w, 2001, fmt.Errorf("图片少于两张"))
		return
	}

	fileList := make([]multipart.File, 0, count)
	for i:=0; i<count; i++ {
		file, _, err := r.FormFile(fmt.Sprintf("file_%d", i+1))
		if err != nil {
			httpser.OutErrBody(w, 2001, err)
			return
		}
		fileList = append(fileList, file)
		_=file.Close()
	}

	out, err := service.ImgStitching(fileList)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgSudoku(w http.ResponseWriter, r *http.Request) {
	fileCount := r.FormValue("file_count")
	count := utils.Str2Int(fileCount)
	if count < 2 {
		httpser.OutErrBody(w, 2001, fmt.Errorf("图片少于两张"))
		return
	}

	fileList := make([]multipart.File, 0, count)
	for i:=0; i<count; i++ {
		file, _, err := r.FormFile(fmt.Sprintf("file_%d", i+1))
		if err != nil {
			httpser.OutErrBody(w, 2001, err)
			return
		}
		fileList = append(fileList, file)
		_=file.Close()
	}

	out, err := service.ImgStitchingSudoku(fileList)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

// 平均裁剪成多份 偶数份
func ImgClipper(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()

	number := utils.Str2Int(r.FormValue("number"))
	if number%2 != 0 {
		httpser.OutErrBody(w, 2001, fmt.Errorf("number 应该为偶数"))
		return
	}


}

// 按坐标矩形裁剪  x1,y1,x2,y2
func ImgClipperRectangle(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	x1 := utils.Str2Int(r.FormValue("x1"))
	y1 := utils.Str2Int(r.FormValue("y1"))
	x2 := utils.Str2Int(r.FormValue("x2"))
	y2 := utils.Str2Int(r.FormValue("y2"))

	out, err := service.ClipperRectangle(file, x1, y1, x2, y2)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)

}

// 按坐标,半径圆形裁剪   x,y,r
func ImgClipperRound(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	x := utils.Str2Int(r.FormValue("x"))
	y := utils.Str2Int(r.FormValue("y"))
	radius := utils.Str2Int(r.FormValue("r"))
	out, err := service.ClipperRound(file, x, y, radius)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgInvert(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	out, err := service.ImgInvert(file)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func ImgGray(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	out, err := service.ImgGray(file)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}

func Img2Txt(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	defer file.Close()
	out, err := service.Img2Txt(file)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_,_=w.Write(out)
}

func ImgAlpha(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	percentage := utils.Str2Float64(r.FormValue("percentage"))
	defer file.Close()
	out, err := service.ImgAlpha(file, percentage)
	if err != nil {
		httpser.OutErrBody(w, 2001, err)
		return
	}
	_,_=w.Write(out)
}