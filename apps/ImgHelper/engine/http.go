package engine

import (
	"github.com/mangenotwork/extras/apps/ImgHelper/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartHttp(){
	go func() {
		logger.Info("StartHttp")
		mux := httpser.NewEngine()

		// 生成二维码  QRCode
		mux.Router("/qrcode", handler.QRCode)

		// 生成条形码  Barcode
		mux.Router("/barcode", handler.Barcode)

		// 识别二维码  QRCodeRecognition
		mux.Router("/qrcode/recognition", handler.QRCodeRecognition)

		// 识别条形码  BarcodeRecognition
		mux.Router("/barcode/recognition", handler.BarcodeRecognition)

		// 图片信息获取
		mux.Router("/image/info", handler.ImageInfo)

		// 图片压缩
		mux.Router("/image/compress", handler.ImageCompress)

		// 图片添加水印
		mux.Router("/watermark/txt", handler.WatermarkTxt)   // - 文字水印
		mux.Router("/watermark/img", handler.WatermarkLogo)  // - 图片水印
		mux.Router("/watermark/logo", handler.WatermarkLogo) // - logo水印

		// 生成文字图片, 应用场景: 文章转图片
		mux.Router("/txt2img", handler.Txt2Img)

		// 图片合成gif
		mux.Router("/img2gif", handler.Img2Gif)

		// 图片旋转
		mux.Router("/img/revolve", handler.ImgRevolve)

		// 图片居中
		mux.Router("/img/center", handler.ImgCenter)

		// 图片拼接
		mux.Router("/img/stitching", handler.ImgStitching) // 默认垂直拼接
		mux.Router("/img/sudoku", handler.ImgSudoku)  // 九宫格

		// 图片剪裁, 平均等份裁剪, 矩形裁剪, 圆形裁剪
		//mux.Handle("/img/clipper", m(http.HandlerFunc(handler.ImgClipper)))
		mux.Router("/img/clipper/rect", handler.ImgClipperRectangle)
		mux.Router("/img/clipper/round", handler.ImgClipperRound)

		// 图片色彩反转
		mux.Router("/img/invert", handler.ImgInvert)

		// 图片灰化
		mux.Router("/img/gray", handler.ImgGray)

		// 图片转为字符画
		mux.Router("/img2txt", handler.Img2Txt)

		// 图片透明
		mux.Router("/img/alpha", handler.ImgAlpha)

		mux.Run()

	}()
}
