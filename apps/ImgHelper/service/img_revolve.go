package service

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

// 旋转90度
func  Revolve(file multipart.File, revolveType string) (outByte []byte, err error)  {
	m, outType, err := image.Decode(file)
	if err != nil {
		return
	}
	if len(revolveType) <= 0 {
		revolveType = "90"
	}

	var outImg *image.RGBA

	switch revolveType {
	case "90":
		outImg = image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
		// 矩阵旋转
		for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
			for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
				//  设置像素点
				outImg.Set(m.Bounds().Max.Y-x, y, m.At(y, x))
			}
		}
	case "180":
		outImg = image.NewRGBA(image.Rect(0, 0, m.Bounds().Dx(), m.Bounds().Dy() ))
		// 矩阵旋转
		for x := m.Bounds().Min.X; x < m.Bounds().Max.X; x++ {
			for y := m.Bounds().Min.Y; y < m.Bounds().Max.Y; y++ {
				//  设置像素点
				outImg.Set(m.Bounds().Max.X-x, m.Bounds().Max.Y-y, m.At(x, y))
			}
		}
	case "270":
		outImg = image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
		// 矩阵旋转
		for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
			for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
				// 设置像素点
				outImg.Set(x, m.Bounds().Max.X-y, m.At(y, x))
			}
		}
	default:
		err = fmt.Errorf("type must is : 90, 180, 270.")
		return
	}

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
