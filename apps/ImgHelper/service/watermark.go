package service

import (
	"bytes"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var fontTTF *truetype.Font

func init() {
	log.Println("加载字体")
	workPath, _ := os.Getwd()
	fontBytes, err := ioutil.ReadFile(workPath+"/ttf/Alibaba-PuHuiTi-Heavy.ttf")
	if err != nil {
		log.Println(err)
	}
	fontTTF, err = freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
	}
}

// position 1:左下角, 2:居中, 3:左上角, 4:右上角, 5:右下角
func WatermarkTxt(file multipart.File, txt, color string, fontSize, dpi, position int) (outByte []byte, err error) {
	strings.Replace(color, "#", "", -1)
	if len(color) < 6 {
		color = "FF0000"
	}
	if fontSize <= 0 {
		fontSize = 18
	}
	if dpi <= 0 {
		dpi = 72
	}

	fontWeight := utf8.RuneCountInString(txt)*fontSize
	fontHeight := fontSize

	imgObj, outType, err := image.Decode(file)
	if err != nil {
		return
	}
	img := image.NewNRGBA(imgObj.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, imgObj.At(x, y))
		}
	}

	f := freetype.NewContext()
	f.SetDPI(float64(dpi))
	f.SetFont(fontTTF)
	f.SetFontSize(float64(fontSize))
	f.SetClip(imgObj.Bounds())
	f.SetDst(img)
	f.SetSrc(image.NewUniform(hex2rgb(color)))
	f.SetHinting(font.HintingFull)

	var pt fixed.Point26_6
	switch position {
	case 2:
		// 居中
		pt = freetype.Pt( (img.Bounds().Dx()-fontWeight)/2 , (img.Bounds().Dy()-fontHeight)/2)
	case 3:
		//左上角
		pt = freetype.Pt(img.Bounds().Dx()-fontWeight, fontHeight)
	case 4:
		//右上角
		pt = freetype.Pt(img.Bounds().Dx()*(1/100), fontHeight)
	case 5:
		//右下角
		pt = freetype.Pt(img.Bounds().Dx()*(1/100), img.Bounds().Dy()-fontHeight)
	default:
		//左下角
		pt = freetype.Pt(img.Bounds().Dx()-fontWeight, img.Bounds().Dy()-fontHeight)
	}
	_, err = f.DrawString(txt, pt)
	if err != nil {
		return
	}

	out := new(bytes.Buffer)
	switch outType {
	case "png","PNG":
		_=png.Encode(out, img)
	case "jpg", "jpeg", "JPG", "JPEG":
		_=jpeg.Encode(out, img, nil)
	case "gif", "GIF":
		_=gif.Encode(out, img, nil)
	}
	outByte = out.Bytes()
	return
}

func hex2rgb(str string) color.Color {
	r, _ := strconv.ParseInt(str[:2], 16, 10)
	g, _ := strconv.ParseInt(str[2:4], 16, 18)
	b, _ := strconv.ParseInt(str[4:], 16, 10)
	return color.RGBA{A: 255, R: uint8(r), G: uint8(g), B: uint8(b)}
}

func WatermarkLogo(imgFile, logoFile multipart.File) (outByte []byte, err error){
	logoImg, _, err := image.Decode(logoFile)
	if err != nil {
		return
	}
	img, outType, err := image.Decode(imgFile)
	if err != nil {
		return
	}

	offset := image.Pt(img.Bounds().Dx()-logoImg.Bounds().Dx()-10, img.Bounds().Dy()-logoImg.Bounds().Dy()-10)
	b := img.Bounds()
	m := image.NewNRGBA(b)

	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, logoImg.Bounds().Add(offset), logoImg, image.ZP, draw.Over)

	out := new(bytes.Buffer)
	switch outType {
	case "png","PNG":
		_=png.Encode(out, m)
	case "jpg", "jpeg", "JPG", "JPEG":
		_=jpeg.Encode(out, m, nil)
	case "gif", "GIF":
		_=gif.Encode(out, m, nil)
	}
	outByte = out.Bytes()
	return
}