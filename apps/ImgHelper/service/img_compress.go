package service

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
)

/*
// Nearest-neighbor interpolation
	NearestNeighbor InterpolationFunction = iota
	// Bilinear interpolation
	Bilinear
	// Bicubic interpolation (with cubic hermite spline)
	Bicubic
	// Mitchell-Netravali interpolation
	MitchellNetravali
	// Lanczos interpolation (a=2)
	Lanczos2
	// Lanczos interpolation (a=3)
	Lanczos3
 */
func 	ImgCompress(img image.Image, width, height int, outType string) []byte {
	// resize.Resize 使用插值函数interp创建具有新尺寸（宽度，高度）的缩放图像。 如果宽度或高度设置为0，则将其设置为保留宽高比值。
	m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	out := new(bytes.Buffer)

	switch outType {
	case "png","PNG":
		_=png.Encode(out, m)
	case "jpg", "jpeg", "JPG", "JPEG":
		_=jpeg.Encode(out, m, nil)
	case "gif", "GIF":
		_=gif.Encode(out, m, nil)
	}
	return out.Bytes()
}




