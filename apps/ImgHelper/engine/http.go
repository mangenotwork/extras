package engine

import (
	"github.com/mangenotwork/extras/apps/ImgHelper/handler"
	"github.com/mangenotwork/extras/common/middleware"
	"github.com/mangenotwork/extras/common/utils"
	"net/http"
)

func StartHttpServer(){
	go func() {
		utils.HttpServer(Router())
	}()
}

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	m := middleware.Base
	mux.Handle("/hello", m(http.HandlerFunc(handler.Hello)))
	mux.Handle("/", m(http.HandlerFunc(handler.Hello)))

	// 生成二维码  QRCode
	mux.Handle("/qrcode", m(http.HandlerFunc(handler.QRCode)))

	// 生成条形码  Barcode
	mux.Handle("/barcode", m(http.HandlerFunc(handler.Barcode)))

	// 识别二维码  QRCodeRecognition
	mux.Handle("/qrcode/recognition", m(http.HandlerFunc(handler.QRCodeRecognition)))

	// 识别条形码  BarcodeRecognition
	mux.Handle("/barcode/recognition", m(http.HandlerFunc(handler.BarcodeRecognition)))

	// 添加水印 - 文字水印   WatermarkTxt
	mux.Handle("/watermark/txt", m(http.HandlerFunc(handler.WatermarkTxt)))

	// 添加水印 - 图片水印   WatermarkImg
	mux.Handle("/watermark/img", m(http.HandlerFunc(handler.WatermarkImg)))

	// 图片信息获取
	mux.Handle("/image/info", m(http.HandlerFunc(handler.ImageInfo)))

	// 图片压缩
	mux.Handle("/image/compress", m(http.HandlerFunc(handler.ImageCompress)))

	return mux
}