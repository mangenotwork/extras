/*

宿主机 需要安装  tesseract

#显示安装的语言包
tesseract --list-langs

#显示帮助
tesseract --help
tesseract --help-extra
tesseract --version


https://tesseract-ocr.github.io/tessdoc/Data-Files.html  下载词典



## 字体库

字体库 (tessdata_best) : 基于LSTM引擎的训练数据，最佳最准确的
https://github.com/tesseract-ocr/tessdata_best

字体库 (tessdata) :  支持双引擎（LSTM和传统引擎），但LSTM训练数据不是最新的版本
https://github.com/tesseract-ocr/tessdata

字体库(tessdata_fast) :  基于LSTM引擎的训练数据，快速（精简）版本
https://github.com/tesseract-ocr/tessdata_fast

总结 :  推荐使用tessdata_best，虽然识别速度相对于tessdata_fast稍慢，但是准确率可以保证



*/
package service

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"

	gs "github.com/otiai10/gosseract/v2"
)

func GetOCRVersion() string {
	return gs.Version()
}

func GetOCRLanguages() ([]string, error) {
	return gs.GetAvailableLanguages()
}

func OCR(imgData []byte, lang string) (string, error){
	if lang == "" {
		lang = "chi_sim"
	}
	client := gs.NewClient()
	client.SetLanguage(lang)
	defer client.Close()
	client.SetImageFromBytes(imgData)
	return client.Text()
}

func Huihua(imgData []byte) []byte {

	buffer := bytes.NewReader(imgData)
	m, outType, err := image.Decode(buffer)
	if err != nil {
		return nil
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
	return out.Bytes()
}

/*
 TODO 提高识别

1. 做好图片的二值化
	二值化就是将图像中灰度值大于某个临界灰度值的像素点设置为灰度最大值，灰度值小于某个临界灰度值的像素点设置为灰度最小值。

2. 合理的降噪

3. 图片resize; 调整尺寸
	Tesseract对于dpi >= 300的图片有更好的识别效果

4. 图片旋转到合适的角度

5. 图片切割

6. 合理的训练自己的识别库。

 */