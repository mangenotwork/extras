package service

import (
	"bytes"
	"fmt"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
)

func ClipperRectangle(file multipart.File, x1, y1, x2, y2 int) (outByte []byte, err error) {
	m, outType, err := image.Decode(file)
	if err != nil {
		return
	}

	out := new(bytes.Buffer)

	switch outType {
	case "jpeg":
		img := m.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(x1, y1, x2, y2)).(*image.YCbCr)
		_=jpeg.Encode(out, subImg, nil)
	case "png":
		var subImg image.Image

		switch m.(type) {
		case *image.NRGBA:
			img := m.(*image.NRGBA)
			subImg = img.SubImage(image.Rect(x1, y1, x2, y2)).(*image.NRGBA)

		case *image.RGBA:
			img := m.(*image.RGBA)
			subImg = img.SubImage(image.Rect(x1, y1, x2, y2)).(*image.RGBA)

		}
		_=png.Encode(out, subImg)

	case "gif":
		img := m.(*image.Paletted)
		subImg := img.SubImage(image.Rect(x1, y1, x2, y2)).(*image.Paletted)
		_=gif.Encode(out, subImg, &gif.Options{})
	case "bmp":
		img := m.(*image.RGBA)
		subImg := img.SubImage(image.Rect(x1, y1, x2, y2)).(*image.RGBA)
		_=bmp.Encode(out, subImg)
	default:
		err = fmt.Errorf("图片类型错误")
		return
	}

	outByte = out.Bytes()
	return
}


type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// TODO BUG
func ClipperRound(file multipart.File, x, y, r int) (outByte []byte, err error) {
	m, outType, err := image.Decode(file)
	if err != nil {
		return
	}

	out := new(bytes.Buffer)

	round := &circle{
		p: image.Point{x, y},
		r : r,
	}

	switch outType {
	case "jpeg":
		img := m.(*image.YCbCr)
		subImg := img.SubImage(round.Bounds()).(*image.YCbCr)
		_=jpeg.Encode(out, subImg, nil)
	case "png":
		var subImg image.Image

		switch m.(type) {
		case *image.NRGBA:
			img := m.(*image.NRGBA)
			subImg = img.SubImage(round.Bounds()).(*image.NRGBA)

		case *image.RGBA:
			img := m.(*image.RGBA)
			subImg = img.SubImage(round.Bounds()).(*image.RGBA)

		}
		_=png.Encode(out, subImg)

	case "gif":
		img := m.(*image.Paletted)
		subImg := img.SubImage(round.Bounds()).(*image.Paletted)
		_=gif.Encode(out, subImg, &gif.Options{})
	case "bmp":
		img := m.(*image.RGBA)
		subImg := img.SubImage(round.Bounds()).(*image.RGBA)
		_=bmp.Encode(out, subImg)
	default:
		err = fmt.Errorf("图片类型错误")
		return
	}

	outByte = out.Bytes()
	return
}
