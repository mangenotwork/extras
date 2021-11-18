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

	// 图片信息获取
	mux.Handle("/image/info", m(http.HandlerFunc(handler.ImageInfo)))

	// 图片压缩
	mux.Handle("/image/compress", m(http.HandlerFunc(handler.ImageCompress)))

	// 图片添加水印
	mux.Handle("/watermark/txt", m(http.HandlerFunc(handler.WatermarkTxt)))   // - 文字水印
	mux.Handle("/watermark/img", m(http.HandlerFunc(handler.WatermarkLogo)))  // - 图片水印
	mux.Handle("/watermark/logo", m(http.HandlerFunc(handler.WatermarkLogo))) // - logo水印

	// 生成文字图片, 应用场景: 文章转图片
	mux.Handle("/txt2img", m(http.HandlerFunc(handler.Txt2Img)))

	// 图片合成gif
	mux.Handle("/img2gif", m(http.HandlerFunc(handler.Img2Gif)))

	// 图片旋转
	mux.Handle("/img/revolve", m(http.HandlerFunc(handler.ImgRevolve)))

	// 图片居中
	mux.Handle("/img/center", m(http.HandlerFunc(handler.ImgCenter)))

	// 图片拼接
	mux.Handle("/img/stitching", m(http.HandlerFunc(handler.ImgStitching)))
	mux.Handle("/img/sudoku", m(http.HandlerFunc(handler.ImgSudoku)))  // 九宫格


	return mux
}