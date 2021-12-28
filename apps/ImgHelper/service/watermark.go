package service

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
)

var fontTTF *truetype.Font

func FontTTFInit() {
	logger.Info("加载字体")
	workPath, _ := os.Getwd()
	logger.Info(workPath + conf.Arg.TTF)
	fontBytes, err := ioutil.ReadFile(workPath + conf.Arg.TTF)
	if err != nil {
		logger.Error(err)
		return
	}
	fontTTF, err = freetype.ParseFont(fontBytes)
	if err != nil {
		logger.Error(err)
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

func WatermarkLogo(imgFile, logoFile multipart.File) (outByte []byte, err error) {
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

// outType  png, jpg; 默认png
func Txt2Img(txt string, fontSize, dpi, spacing int, outType string) (outByte []byte, err error) {
	if fontSize <= 0 {
		fontSize = 16
	}
	if dpi <= 0 {
		dpi = 72
	}
	if spacing <= 0 {
		spacing = 2
	}
	fg, bg := image.Black, image.White

	txtList := strings.Split(txt,"\n")
	max := 0
	for _,v := range txtList {
		if utf8.RuneCountInString(v) > max {
			max = utf8.RuneCountInString(v)
		}
	}

	rgba := image.NewRGBA(image.Rect(0, 0, max*fontSize+20, len(txtList)*fontSize*2+20))
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(float64(dpi))
	c.SetFont(fontTTF)
	c.SetFontSize(float64(fontSize))
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	pt := freetype.Pt(10, 10+int(c.PointToFixed(float64(fontSize))>>6))
	for _, s := range txtList {
		_, err = c.DrawString(s, pt)
		if err != nil {
			logger.Error(err)
			return
		}
		pt.Y += c.PointToFixed(float64(fontSize) * float64(spacing))
	}

	out := new(bytes.Buffer)
	switch outType {
	case "png","PNG":
		_=png.Encode(out, rgba)
	case "jpg", "jpeg", "JPG", "JPEG":
		_=jpeg.Encode(out, rgba, nil)
	case "gif", "GIF":
		_=gif.Encode(out, rgba, nil)
	default:
		_=png.Encode(out, rgba)
	}
	outByte = out.Bytes()
	return
}

