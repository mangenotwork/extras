package engine

import (
	"github.com/mangenotwork/extras/apps/WordHelper/handler"
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

	// 分词
	mux.Handle("/fenci/jieba",  m(http.HandlerFunc(handler.JieBaFenCi)))

	// ocr
	mux.Handle("/ocr", m(http.HandlerFunc(handler.OCR)))
	mux.Handle("/ocr/languages", m(http.HandlerFunc(handler.GetOCRLanguages)))
	mux.Handle("/ocr/version", m(http.HandlerFunc(handler.GetOCRVersion)))

	// 翻译
	mux.Handle("/fanyi", m(http.HandlerFunc(handler.FanYi)))

	// pdf提取
	mux.Handle("/pdf/txt", m(http.HandlerFunc(handler.PDFExtractionTxt)))
	mux.Handle("/pdf/row", m(http.HandlerFunc(handler.PDFExtractionRow)))
	mux.Handle("/pdf/table", m(http.HandlerFunc(handler.PDFExtractionTable)))

	// AES
	mux.Handle("/aes/cbc/encrypt", m(http.HandlerFunc(handler.AESCBCEncrypt)))
	mux.Handle("/aes/cbc/decrypt", m(http.HandlerFunc(handler.AESCBCDecrypt)))
	mux.Handle("/aes/ecb/encrypt", m(http.HandlerFunc(handler.AESECBEncrypt)))
	mux.Handle("/aes/ecb/decrypt", m(http.HandlerFunc(handler.AESECBDecrypt)))
	mux.Handle("/aes/cfb/encrypt", m(http.HandlerFunc(handler.AESCFBEncrypt)))
	mux.Handle("/aes/cfb/decrypt", m(http.HandlerFunc(handler.AESCFBDecrypt)))
	mux.Handle("/aes/ctr/encrypt", m(http.HandlerFunc(handler.AESCTREncrypt)))
	mux.Handle("/aes/ctr/decrypt", m(http.HandlerFunc(handler.AESCTRDecrypt)))

	// DES
	mux.Handle("/des/cbc/encrypt", m(http.HandlerFunc(handler.DESCBCEncrypt)))
	mux.Handle("/des/cbc/decrypt", m(http.HandlerFunc(handler.DESCBCDecrypt)))
	mux.Handle("/des/ecb/encrypt", m(http.HandlerFunc(handler.DESECBEncrypt)))
	mux.Handle("/des/ecb/decrypt", m(http.HandlerFunc(handler.DESECBDecrypt)))
	mux.Handle("/des/cfb/encrypt", m(http.HandlerFunc(handler.DESCFBEncrypt)))
	mux.Handle("/des/cfb/decrypt", m(http.HandlerFunc(handler.DESCFBDecrypt)))
	mux.Handle("/des/ctr/encrypt", m(http.HandlerFunc(handler.DESCTREncrypt)))
	mux.Handle("/des/ctr/decrypt", m(http.HandlerFunc(handler.DESCTRDecrypt)))

	// md5
	mux.Handle("/md5/16", m(http.HandlerFunc(handler.MD516)))
	mux.Handle("/md5/32", m(http.HandlerFunc(handler.MD532)))

	// base64
	mux.Handle("/base64/encrypt", m(http.HandlerFunc(handler.Base64Encrypt)))
	mux.Handle("/base64/decrypt", m(http.HandlerFunc(handler.Base64Decrypt)))
	mux.Handle("/base64url/encrypt", m(http.HandlerFunc(handler.Base64UrlEncrypt)))
	mux.Handle("/base64url/decrypt", m(http.HandlerFunc(handler.Base64UrlDecrypt)))

	// Hmac
	mux.Handle("/hmac/md5", m(http.HandlerFunc(handler.HmacMD5)))
	mux.Handle("/hmac/sha1", m(http.HandlerFunc(handler.HmacSHA1)))
	mux.Handle("/hmac/sha256", m(http.HandlerFunc(handler.HmacSHA256)))
	mux.Handle("/hmac/sha512", m(http.HandlerFunc(handler.HmacSHA512)))

	// TODO: PBKDF2


	return mux
}