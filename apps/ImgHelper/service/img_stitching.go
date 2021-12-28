package service

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"

	"github.com/mangenotwork/extras/common/logger"
)

// 选用最大图片的宽度, 垂直依次拼接
func ImgStitching(fileList []multipart.File) (outByte []byte, err error) {
	imgList := make([]image.Image, 0)
	height := 0
	width := 0
	outType := ""
	for _, m := range fileList {
		img, imgType, imgErr := image.Decode(m)
		if imgErr != nil {
			err = imgErr
			return
		}
		outType = imgType
		height+=img.Bounds().Dy()
		if width <= img.Bounds().Dx() {
			width = img.Bounds().Dx()
		}
		imgList = append(imgList, img)
	}

	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	lastY := 0
	for _, m := range imgList {
		src := resize.Resize(uint(width), uint(m.Bounds().Dy()), m, resize.Lanczos3)
		draw.Draw(rgba, src.Bounds().Add(image.Pt(0, lastY)), src, image.ZP, draw.Src)
		lastY += m.Bounds().Dy()
	}

	out := new(bytes.Buffer)
	switch outType {
	case "png","PNG":
		_=png.Encode(out, rgba)
	case "jpg", "jpeg", "JPG", "JPEG":
		_=jpeg.Encode(out, rgba, nil)
	case "gif", "GIF":
		_=gif.Encode(out, rgba, nil)
	}
	outByte = out.Bytes()
	return

}

// 选用最大图片宽高, 九宫格拼接
func ImgStitchingSudoku(fileList []multipart.File) (outByte []byte, err error) {
	imgList := make([]image.Image, 0)
	height := 0
	width := 0
	outType := ""
	for _, m := range fileList {
		img, imgType, imgErr := image.Decode(m)
		if imgErr != nil {
			err = imgErr
			return
		}
		outType = imgType

		if width <= img.Bounds().Dx() {
			width = img.Bounds().Dx()
		}
		if height <= img.Bounds().Dy() {
			height = img.Bounds().Dy()
		}
		imgList = append(imgList, img)
	}
	//logger.Info(height, width, outType)

	td := 2
	if len(fileList) >= 9 {
		td = 3
	}
	if len(fileList) >= 16 {
		td = 4
	}
	tr := len(fileList)/td
	logger.Info("td, tr = ", td, tr)
	rgba := image.NewRGBA(image.Rect(0, 0, td*width, tr*height))

	x := 0
	y := 0
	n := 0
	for _, m := range imgList {
		src := resize.Resize(uint(width), uint(height), m, resize.Lanczos3)
		if n >= td {
			x = 0
			n = 0
			y += height
		}
		draw.Draw(rgba, src.Bounds().Add(image.Pt(x, y)), src, image.ZP, draw.Src)
		x+=m.Bounds().Dx()
		n++
	}

	outImg := resize.Resize(uint(width), uint(height), rgba, resize.Lanczos3)
	out := new(bytes.Buffer)
	switch outType {
	case "png","PNG":
		_=png.Encode(out, outImg)
	case "jpg", "jpeg", "JPG", "JPEG":
		_=jpeg.Encode(out, outImg, nil)
	case "gif", "GIF":
		_=gif.Encode(out, outImg, nil)
	}
	outByte = out.Bytes()
	return
}

// TODO 选用最小图片高度, 水平依次拼接
func ImgStitchingParallel() {

}