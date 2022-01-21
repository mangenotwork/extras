package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"time"

	"github.com/mangenotwork/extras/common/command"
	"github.com/mangenotwork/extras/common/utils"
)

func Conversion(file multipart.File, handler *multipart.FileHeader, target string) (string, error){
	fileSuffix := path.Ext(handler.Filename)
	filePrefix := handler.Filename[0:len(handler.Filename) - len(fileSuffix)]
	t := utils.Int642Str(time.Now().Unix())
	pathStr := "./temp/" + filePrefix+"-"+t+fileSuffix
	cur, err := os.Create(pathStr)
	defer cur.Close()
	if err != nil {
		return "", err
	}

	_,err = io.Copy(cur, file)
	if err != nil {
		return "", err
	}

	var ok bool
	var cmdErr error
	var newfileSuffix string

	switch target {
	case "pdf":
		ok, cmdErr = command.Cmd(20*time.Second, "libreoffice", "--invisible", "--convert-to", "pdf", pathStr, "--outdir", "./temp")
		newfileSuffix = ".pdf"
	case "html":
		ok, cmdErr = command.Cmd(20*time.Second, "libreoffice", "--invisible", "--convert-to", "html", pathStr, "--outdir", "./temp")
		newfileSuffix = ".html"
	case "jpg":
		ok, cmdErr = command.Cmd(20*time.Second, "libreoffice", "--invisible", "--convert-to", "jpg", pathStr, "--outdir", "./temp")
		newfileSuffix = ".jpg"
	}
	if cmdErr != nil {
		return "", cmdErr
	}

	err = os.Remove(pathStr)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", fmt.Errorf("转换失败,请重新上传")
	}

	dir,_ := os.Getwd()
	pathStrPdf := dir + "/temp/" + filePrefix+"-" + t + newfileSuffix
	return pathStrPdf, nil
}

