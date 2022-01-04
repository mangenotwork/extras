package handler

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/mangenotwork/extras/apps/WordHelper/proto"
	"github.com/mangenotwork/extras/apps/WordHelper/service"
	"github.com/mangenotwork/extras/apps/WordHelper/service/pdf"
)

type GRPCService struct {}

// 结巴分词
func (*GRPCService) FenciJieba(ctx context.Context, req *proto.FenciJiebaReq) (*proto.FenciJiebaResp, error) {
	resp := new(proto.FenciJiebaResp)
	/*
	// req.Type
	// 		1: 全模式
	// 		2: 精确模式
	//		3: 搜索引擎模式
	//		4: 词性标注
	//     == 单独
	//      5: Tokenize 搜索引擎模式
	//      6: Tokenize 默认模式
	//      7: Extract
	 */
	if int(req.Type) > 5 || int(req.Type) < 1 {
		return resp, fmt.Errorf("Type 错误, 应该是 1~4")
	}
	for _, v := range service.JieBa(req.Str, int(req.Type)) {
		resp.Data = append(resp.Data, v.(string))
	}
	return resp, nil
}

// OCR 识别
func (*GRPCService) OCR(ctx context.Context, req *proto.OCRReq) (*proto.OCRResp, error) {
	var err error
	resp := new(proto.OCRResp)
	resp.Data, err = service.OCR(req.File, req.Lang)
	return resp, err
}

// OCR 语言包列表
func (*GRPCService) OCRLanguages(ctx context.Context, req *proto.OCRLangReq) (*proto.OCRLangResp, error) {
	var err error
	resp := new(proto.OCRLangResp)
	resp.Data, err = service.GetOCRLanguages()
	return resp, err
}

// OCR 版本
func (*GRPCService) OCRVersion(ctx context.Context, req *proto.OCRVersionReq) (*proto.OCRVersionResp, error) {
	resp := new(proto.OCRVersionResp)
	resp.Data = service.GetOCRVersion()
	return resp, nil
}

// OCR 识别base64图片
func (*GRPCService) OCRBase64(ctx context.Context, req *proto.OCRBase64Req) (*proto.OCRBase64Resp, error) {
	var err error
	resp := new(proto.OCRBase64Resp)

	lang := req.Lang
	base64img := req.Base64Img

	b, _ := regexp.MatchString(`^data:\s*image\/(\w+);base64,`, base64img)
	if b {
		re, _ := regexp.Compile(`^data:\s*image\/(\w+);base64,`)
		base64img = re.ReplaceAllString(base64img, "")
	}

	sDec, err := base64.StdEncoding.DecodeString(base64img)
	if err != nil {
		return resp, err
	}

	sDec = service.Huihua(sDec)
	//file := "./a.png"
	//err = ioutil.WriteFile(file, sDec, 0666)
	//if err != nil {
	//	logger.Error(err)
	//}
	resp.Data, err = service.OCR(sDec, lang)
	return resp, err
}

// 翻译
func (*GRPCService) Fanyi(ctx context.Context, req *proto.FanyiReq) (*proto.FanyiResp, error) {
	var err error
	resp := new(proto.FanyiResp)
	res, err :=http.Get("http://fanyi.youdao.com/translate?&doctype=json&type=AUTO&i="+req.Word)
	if err != nil {
		return resp, nil
	}
	resp.Body, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return resp, err
}

// PDF 提取文本内容
func (*GRPCService) PDFTxt(ctx context.Context, req *proto.PDFTxtReq) (*proto.PDFTxtResp, error) {
	resp := &proto.PDFTxtResp{
		Data : make([]*proto.PDFTxtBody, 0),
	}

	read := bytes.NewReader(req.File)
	reader, err := pdf.NewReader(read, int64(len(req.File)))
	if err != nil {
		return resp, err
	}

	for i:=1; i< reader.NumPage(); i++ {
		pg := reader.Page(i)
		txt,_ := pg.GetTxt()
		resp.Data = append(resp.Data, &proto.PDFTxtBody{
			Page: int32(i),
			Content: txt,
		})
	}
	return resp, nil
}

// PDF 按行提取文本内容
func (*GRPCService) PDFRow(ctx context.Context, req *proto.PDFRowReq) (*proto.RDFRowResp, error) {
	resp := &proto.RDFRowResp{
		Data : make([]*proto.PDFRowBody, 0),
	}

	read := bytes.NewReader(req.File)
	reader, err := pdf.NewReader(read, int64(len(req.File)))
	if err != nil {
		return resp, err
	}

	for i:=1; i< reader.NumPage(); i++ {
		pg := reader.Page(i)
		row, _ := pg.GetRow()
		resp.Data = append(resp.Data, &proto.PDFRowBody{
			Page: int32(i),
			Content: row,
		})
	}
	return resp, nil
}

// PDF 提取表格
func (*GRPCService) PDFTable (ctx context.Context, req *proto.PDFTableReq) (*proto.PDFTableResp, error) {
	resp := &proto.PDFTableResp{
		Data : make([]*proto.PDFTableBody, 0),
	}

	read := bytes.NewReader(req.File)
	reader, err := pdf.NewReader(read, int64(len(req.File)))
	if err != nil {
		return resp, err
	}

	for i:=1; i< reader.NumPage(); i++ {
		pg := reader.Page(i)
		data := &proto.PDFTableBody{
			Page: int32(i),
			Content: make([]*proto.PDFTableBodyMap, 0),
		}
		for _, v := range pg.GetTable() {
			data.Content = append(data.Content, &proto.PDFTableBodyMap{
				Data: v,
			})
		}
		resp.Data = append(resp.Data, data)
	}
	return resp, nil
}

// MD 转 html
func (*GRPCService) Md2Html (ctx context.Context, req *proto.Md2HtmlReq) (*proto.Md2HtmlResp, error) {
	resp := new(proto.Md2HtmlResp)
	resp.Data = service.MarkdownToHTML(req.Str)
	return resp, nil
}

// AES CBC Encrypt
func (*GRPCService) AESCBCEncrypt (ctx context.Context, req *proto.AESCBCEncryptReq) (*proto.AESCBCEncryptResp, error) {
	resp := new(proto.AESCBCEncryptResp)
	b, err := service.NewAES("cbc", []byte(req.Param.Iv)).Encrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// AES CBC Decrypt
func (*GRPCService) AESCBCDecrypt (ctx context.Context, req *proto.AESCBCDecryptReq) (*proto.AESCBCDecryptResp, error) {
	resp := new(proto.AESCBCDecryptResp)
	decoded, err := base64.StdEncoding.DecodeString(req.Param.Str)
	if err != nil {
		return resp, err
	}
	b, err := service.NewAES("cbc", []byte(req.Param.Iv)).Decrypt(decoded, []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// AES ECB Encrypt
func (*GRPCService) AESECBEncrypt (ctx context.Context, req *proto.AESECBEncryptReq) (*proto.AESECBEncryptResp, error) {
	resp := new(proto.AESECBEncryptResp)
	b, err := service.NewAES("ecb", []byte(req.Param.Iv)).Encrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// AES ECB Decrypt
func (*GRPCService) AESECBDecrypt (ctx context.Context, req *proto.AESECBDecryptReq) (*proto.AESECBDecryptResp, error) {
	resp := new(proto.AESECBDecryptResp)
	decoded, err := base64.StdEncoding.DecodeString(req.Param.Str)
	if err != nil {
		return resp, err
	}
	b, err := service.NewAES("ecb", []byte(req.Param.Iv)).Decrypt(decoded, []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// AES CFB Encrypt
func (*GRPCService) AESCFBEncrypt (ctx context.Context, req *proto.AESCFBEncryptReq) (*proto.AESCFBEncryptResp, error) {
	resp := new(proto.AESCFBEncryptResp)
	b, err := service.NewAES("cfb", []byte(req.Param.Iv)).Encrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// AES CFB Decrypt
func (*GRPCService) AESCFBDecrypt (ctx context.Context, req *proto.AESCFBDecryptReq) (*proto.AESCFBDecryptResp, error) {
	resp := new(proto.AESCFBDecryptResp)
	decoded, err := base64.StdEncoding.DecodeString(req.Param.Str)
	if err != nil {
		return resp, err
	}
	b, err := service.NewAES("cfb", []byte(req.Param.Iv)).Decrypt(decoded, []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// AES CTR Encrypt
func (*GRPCService) AESCTREncrypt (ctx context.Context, req *proto.AESCTREncryptReq) (*proto.AESCTREncryptResp, error) {
	resp := new(proto.AESCTREncryptResp)
	b, err := service.NewAES("ctr", []byte(req.Param.Iv)).Encrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// AES CTR Decrypt
func (*GRPCService) AESCTRDecrypt (ctx context.Context, req *proto.AESCTRDecryptReq) (*proto.AESCTRDecryptResp, error) {
	resp := new(proto.AESCTRDecryptResp)
	decoded, err := base64.StdEncoding.DecodeString(req.Param.Str)
	if err != nil {
		return resp, err
	}
	b, err := service.NewAES("ctr", []byte(req.Param.Iv)).Decrypt(decoded, []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// DES CBC Encrypt
func (*GRPCService) DESCBCEncrypt (ctx context.Context, req *proto.DESCBCEncryptReq) (*proto.DESCBCEncryptResp, error) {
	resp := new(proto.DESCBCEncryptResp)
	b, err := service.NewDES("cbc", []byte(req.Param.Iv)).Encrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// DES CBC Decrypt
func (*GRPCService) DESCBCDecrypt (ctx context.Context, req *proto.DESCBCDecryptReq) (*proto.DESCBCDecryptResp, error) {
	resp := new(proto.DESCBCDecryptResp)
	b, err := service.NewDES("cbc", []byte(req.Param.Iv)).Decrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// DES ECB Encrypt
func (*GRPCService) DESECBEncrypt (ctx context.Context, req *proto.DESECBEncryptReq) (*proto.DESECBEncryptResp, error) {
	resp := new(proto.DESECBEncryptResp)
	b, err := service.NewDES("ecb", []byte(req.Param.Iv)).Encrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// DES ECB Decrypt
func (*GRPCService) DESECBDecrypt (ctx context.Context, req *proto.DESECBDecryptReq) (*proto.DESECBDecryptResp, error) {
	resp := new(proto.DESECBDecryptResp)
	b, err := service.NewDES("ecb", []byte(req.Param.Iv)).Decrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// DES CFB Encrypt
func (*GRPCService) DESCFBEncrypt (ctx context.Context, req *proto.DESCFBEncryptReq) (*proto.DESCFBEncryptResp, error) {
	resp := new(proto.DESCFBEncryptResp)
	b, err := service.NewDES("cfb", []byte(req.Param.Iv)).Encrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// DES CFB Decrypt
func (*GRPCService) DESCFBDecrypt (ctx context.Context, req *proto.DESCFBDecryptReq) (*proto.DESCFBDecryptResp, error) {
	resp := new(proto.DESCFBDecryptResp)
	b, err := service.NewDES("cfb", []byte(req.Param.Iv)).Decrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// DES CTR Encrypt
func (*GRPCService) DESCTREncrypt (ctx context.Context, req *proto.DESCTREncryptReq) (*proto.DESCTREncryptResp, error) {
	resp := new(proto.DESCTREncryptResp)
	b, err := service.NewDES("ctr", []byte(req.Param.Iv)).Encrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// DES CTR Decrypt
func (*GRPCService) DESCTRDecrypt (ctx context.Context, req *proto.DESCTRDecryptReq) (*proto.DESCTRDecryptResp, error) {
	resp := new(proto.DESCTRDecryptResp)
	b, err := service.NewDES("ctr", []byte(req.Param.Iv)).Decrypt([]byte(req.Param.Str), []byte(req.Param.Key))
	resp.Data = string(b)
	return resp, err
}

// Hmac Md5
func (*GRPCService) HmacMd5 (ctx context.Context, req *proto.HmacMd5Req) (*proto.HmacMd5Resp, error) {
	resp := new(proto.HmacMd5Resp)
	resp.Data = service.HmacMD5(req.Str, req.Key)
	return resp, nil
}

// Hmac Sha1
func (*GRPCService) HmacSha1 (ctx context.Context, req *proto.HmacSha1Req) (*proto.HmacSha1Resp, error) {
	resp := new(proto.HmacSha1Resp)
	resp.Data = service.HmacSHA1(req.Str, req.Key)
	return resp, nil
}

// Hmac Sha256
func (*GRPCService) HmacSha256 (ctx context.Context, req *proto.HmacSha256Req) (*proto.HmacSha256Resp, error) {
	resp := new(proto.HmacSha256Resp)
	resp.Data = service.HmacSHA256(req.Str, req.Key)
	return resp, nil
}

// Hmac Sha512
func (*GRPCService) HmacSha512 (ctx context.Context, req *proto.HmacSha512Req) (*proto.HmacSha512Resp, error) {
	resp := new(proto.HmacSha512Resp)
	resp.Data = service.HmacSHA512(req.Str, req.Key)
	return resp, nil
}


