package service

import (
	"bytes"
	"github.com/nfnt/resize"
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

func ImgGray(file multipart.File) (outByte []byte, err error) {
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
			_, g, _, a := colorRgb.RGBA()
			gUint8 := uint8(g >> 8)
			aUint8 := uint8(a >> 8)
			newRgba.SetRGBA(i, j, color.RGBA{gUint8, gUint8, gUint8, aUint8})
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

func Img2Txt(file multipart.File) (outByte []byte, err error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return
	}
	// 改一下图片的尺寸
	imgBounds := img.Bounds()
	height := 50 //150
	width := 80 //150
	if imgBounds.Dy() < 150 {
		height = imgBounds.Dy()
	}
	if imgBounds.Dx() < 150 {
		width = imgBounds.Dx()
	}
	m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	arr := []string{"M", "N", "H", "Q", "$", "O", "C", "?", "7", ">", "!", ":", "–", ";", "."}

	out := new(bytes.Buffer)

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			colorRgb := m.At(j, i)
			_, g, _, _ := colorRgb.RGBA()
			avg := uint8(g >> 8)
			num := avg / 18
			out.WriteString(arr[num])
			if j == dx-1 {
				out.WriteString("\n")
			}
		}
	}

	outByte = out.Bytes()
	return
}

func ImgAlpha(file multipart.File, percentage float64) (outByte []byte, err error) {
	m, outType, err := image.Decode(file)
	if err != nil {
		return
	}

	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA64(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := m.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			opacity := uint16(float64(a)*percentage)
			//颜色模型转换，至关重要！
			v := newRgba.ColorModel().Convert(color.NRGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: opacity})
			//Alpha = 0: Full transparent
			rr, gg, bb, aa := v.RGBA()
			newRgba.SetRGBA64(i, j, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})
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

