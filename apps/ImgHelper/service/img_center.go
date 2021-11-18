package service

import (
	"bytes"
	"image"
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
