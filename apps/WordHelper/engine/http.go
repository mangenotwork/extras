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

	return mux
}