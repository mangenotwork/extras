package service

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	qrcodereco "github.com/tuotoo/qrcode"
)

// 生成二维码
func QrCodeBase64(value string) ([]byte, error) {
	return qrcode.Encode(value, qrcode.High, 256)
}

// 生成条形码
func Barcode(value string) ([]byte, error) {
	var err error
	// 创建一个code128编码的 BarcodeIntCS
	cs, _ := code128.Encode(value)
	buf := new(bytes.Buffer)
	// 设置图片像素大小
	qrCode, err := barcode.Scale(cs, 350, 70)
	// 将code128的条形码编码为png图片
	err = png.Encode(buf, qrCode)
	return buf.Bytes(), err
}

// 识别二维码
func QRCodeRecognition() {
	fi, err := os.Open("qrcode.png") // 默认到$GOPATH/src/ 下找
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer fi.Close()

	qrmatrix, err := qrcodereco.Decode(fi)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info(qrmatrix.Content)
}

func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}

// 生成二维码
func createQrCode(content string) (img image.Image, err error) {
	var qrCode *qrcode.QRCode

	qrCode, err = qrcode.New(content, qrcode.Highest)

	if err != nil {
		return nil, fmt.Errorf("创建二维码失败")
	}
	qrCode.DisableBorder = true

	img = qrCode.Image(256)

	return img, nil
}

