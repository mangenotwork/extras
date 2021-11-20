package service

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

func ImgCenter(file multipart.File) (outByte []byte, err error) {
	m, outType, err := image.Decode(file)
	if err != nil {
		return
	}

	max := m.Bounds().Dx()
	// 居中后距离最底部的高度为(x-y)/2
	temp := (max - m.Bounds().Dy()) / 2
	centerImage := image.NewRGBA(image.Rect(0, 0, max, max))
	for x := m.Bounds().Min.X; x < m.Bounds().Max.X; x++ {
		for y := m.Bounds().Min.Y; y < m.Bounds().Max.Y; y++ {
			centerImage.Set(x, temp+y, m.At(x, y))
		}
	}

	out := new(bytes.Buffer)
	switch outType {
	case "png","PNG":
		_=png.Encode(out, centerImage)
	case "jpg", "jpeg", "JPG", "JPEG":
		_=jpeg.Encode(out, centerImage, nil)
	case "gif", "GIF":
		_=gif.Encode(out, centerImage, nil)
	}
	outByte = out.Bytes()
	return

}

func ImgInvert(file multipart.File) (outByte []byte, err error) {
	m, outType, err := image.Decode(file)
	if err != nil {
		return
	}

	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			rUint8 := uint8(r >> 8)
			gUint8 := uint8(g >> 8)
			bUint8 := uint8(b >> 8)
			aUint8 := uint8(a >> 8)
			rUint8 = 255 - rUint8
			gUint8 = 255 - gUint8
			bUint8 = 255 - bUint8
			newRgba.SetRGBA(i, j, color.RGBA{rUint8, gUint8, bUint8, aUint8})
		}
	}

	out := new(bytes.Buffer)
	switch outType {
	case "png","PNG":
		_=png.Encode(out, newRgba)
	case "jpg", "jpeg", "JPG", "JPEG":
		_=jpeg.Encode(out, newRgba, nil)
	case "gif", "GIF":
		_=gif.Encode(out, newRgba, nil)
	}
	outByte = out.Bytes()
	return
}