package engine

import (
	"github.com/mangenotwork/extras/apps/WordHelper/handler"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartHttp(){
	go func() {
		logger.Info("StartHttp")
		mux := httpser.NewEngine()

		// 分词
		mux.Router("/fenci/jieba",  handler.JieBaFenCi)

		// ocr
		mux.Router("/ocr", handler.OCR)
		mux.Router("/ocr/languages", handler.GetOCRLanguages)
		mux.Router("/ocr/version", handler.GetOCRVersion)
		mux.Router("/ocr/base64", handler.OCRBase64Img) // 识别base64图片

		// 翻译
		mux.Router("/fanyi", handler.FanYi)

		// pdf提取
		mux.Router("/pdf/txt", handler.PDFExtractionTxt)
		mux.Router("/pdf/row", handler.PDFExtractionRow)
		mux.Router("/pdf/table", handler.PDFExtractionTable)

		// AES
		mux.Router("/aes/cbc/encrypt", handler.AESCBCEncrypt)
		mux.Router("/aes/cbc/decrypt", handler.AESCBCDecrypt)
		mux.Router("/aes/ecb/encrypt", handler.AESECBEncrypt)
		mux.Router("/aes/ecb/decrypt", handler.AESECBDecrypt)
		mux.Router("/aes/cfb/encrypt", handler.AESCFBEncrypt)
		mux.Router("/aes/cfb/decrypt", handler.AESCFBDecrypt)
		mux.Router("/aes/ctr/encrypt", handler.AESCTREncrypt)
		mux.Router("/aes/ctr/decrypt", handler.AESCTRDecrypt)

		// DES
		mux.Router("/des/cbc/encrypt", handler.DESCBCEncrypt)
		mux.Router("/des/cbc/decrypt", handler.DESCBCDecrypt)
		mux.Router("/des/ecb/encrypt", handler.DESECBEncrypt)
		mux.Router("/des/ecb/decrypt", handler.DESECBDecrypt)
		mux.Router("/des/cfb/encrypt", handler.DESCFBEncrypt)
		mux.Router("/des/cfb/decrypt", handler.DESCFBDecrypt)
		mux.Router("/des/ctr/encrypt", handler.DESCTREncrypt)
		mux.Router("/des/ctr/decrypt", handler.DESCTRDecrypt)

		// md5
		mux.Router("/md5/16", handler.MD516)
		mux.Router("/md5/32", handler.MD532)

		// base64
		mux.Router("/base64/encrypt", handler.Base64Encrypt)
		mux.Router("/base64/decrypt", handler.Base64Decrypt)
		mux.Router("/base64url/encrypt", handler.Base64UrlEncrypt)
		mux.Router("/base64url/decrypt", handler.Base64UrlDecrypt)

		// Hmac
		mux.Router("/hmac/md5", handler.HmacMD5)
		mux.Router("/hmac/sha1", handler.HmacSHA1)
		mux.Router("/hmac/sha256", handler.HmacSHA256)
		mux.Router("/hmac/sha512", handler.HmacSHA512)

		// TODO: PBKDF2

		// md 转 html
		mux.Router("doc/change/md2html", handler.Md2Html)

		mux.Run()

	}()
}
