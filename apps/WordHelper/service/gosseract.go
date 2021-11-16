/*

宿主机 需要安装  tesseract

#显示安装的语言包
tesseract --list-langs

#显示帮助
tesseract --help
tesseract --help-extra
tesseract --version


https://tesseract-ocr.github.io/tessdoc/Data-Files  下载词典

*/
package service


import (
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
	text, err := client.Text()
	return text, err
}
