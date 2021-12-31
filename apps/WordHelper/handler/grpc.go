package handler

import (
	"context"
	"github.com/mangenotwork/extras/apps/WordHelper/proto"
	"github.com/mangenotwork/extras/apps/WordHelper/service"
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
	//      5: Tokenize 搜索引擎模式
	//      6: Tokenize 默认模式
	//      7: Extract
	 */
	data := service.JieBa(req.Str, req.Type)
	return resp, nil
}

// OCR 识别
func (*GRPCService) OCR(ctx context.Context, req *proto.OCRReq) (*proto.OCRResp, error) {
	resp := new(proto.OCRResp)
	return resp, nil
}

// OCR 语言包列表
func (*GRPCService) OCRLanguages(ctx context.Context, req *proto.OCRLangReq) (*proto.OCRLangResp, error) {
	resp := new(proto.OCRLangResp)
	return resp, nil
}

// OCR 版本
func (*GRPCService) OCRVersion(ctx context.Context, req *proto.OCRVersionReq) (*proto.OCRVersionResp, error) {
	resp := new(proto.OCRVersionResp)
	return resp, nil
}

// OCR 识别base64图片
func (*GRPCService) OCRBase64(ctx context.Context, req *proto.OCRBase64Req) (*proto.OCRBase64Resp, error) {
	resp := new(proto.OCRBase64Resp)
	return resp, nil
}

// 翻译
func (*GRPCService) Fanyi(ctx context.Context, req *proto.FanyiReq) (*proto.FanyiResp, error) {
	resp := new(proto.FanyiResp)
	return resp, nil
}

// PDF 提取文本内容
func (*GRPCService) PDFTxt(ctx context.Context, req *proto.PDFTxtReq) (*proto.PDFTxtResp, error) {
	resp := new(proto.PDFTxtResp)
	return resp, nil
}

// PDF 按行提取文本内容
func (*GRPCService) PDFRow(ctx context.Context, req *proto.PDFRowReq) (*proto.RDFRowResp, error) {
	resp := new(proto.RDFRowResp)
	return resp, nil
}

// PDF 提取表格
func (*GRPCService) PDFTable (ctx context.Context, req *proto.PDFTableReq) (*proto.PDFTableResp, error) {
	resp := new(proto.PDFTableResp)
	return resp, nil
}

// MD 转 html
func (*GRPCService) Md2Html (ctx context.Context, req *proto.Md2HtmlReq) (*proto.Md2HtmlResp, error) {
	resp := new(proto.Md2HtmlResp)
	return resp, nil
}

// AES CBC Encrypt
func (*GRPCService) AESCBCEncrypt (ctx context.Context, req *proto.AESCBCEncryptReq) (*proto.AESCBCEncryptResp, error) {
	resp := new(proto.AESCBCEncryptResp)
	return resp, nil
}

// AES CBC Decrypt
func (*GRPCService) AESCBCDecrypt (ctx context.Context, req *proto.AESCBCDecryptReq) (*proto.AESCBCDecryptResp, error) {
	resp := new(proto.AESCBCDecryptResp)
	return resp, nil
}

// AES ECB Encrypt
func (*GRPCService) AESECBEncrypt (ctx context.Context, req *proto.AESECBEncryptReq) (*proto.AESECBEncryptResp, error) {
	resp := new(proto.AESECBEncryptResp)
	return resp, nil
}

// AES ECB Decrypt
func (*GRPCService) AESECBDecrypt (ctx context.Context, req *proto.AESECBDecryptReq) (*proto.AESECBDecryptResp, error) {
	resp := new(proto.AESECBDecryptResp)
	return resp, nil
}

// AES CFB Encrypt
func (*GRPCService) AESCFBEncrypt (ctx context.Context, req *proto.AESCFBEncryptReq) (*proto.AESCFBEncryptResp, error) {
	resp := new(proto.AESCFBEncryptResp)
	return resp, nil
}

// AES CFB Decrypt
func (*GRPCService) AESCFBDecrypt (ctx context.Context, req *proto.AESCFBDecryptReq) (*proto.AESCFBDecryptResp, error) {
	resp := new(proto.AESCFBDecryptResp)
	return resp, nil
}

// AES CTR Encrypt
func (*GRPCService) AESCTREncrypt (ctx context.Context, req *proto.AESCTREncryptReq) (*proto.AESCTREncryptResp, error) {
	resp := new(proto.AESCTREncryptResp)
	return resp, nil
}

// AES CTR Decrypt
func (*GRPCService) AESCTRDecrypt (ctx context.Context, req *proto.AESCTRDecryptReq) (*proto.AESCTRDecryptResp, error) {
	resp := new(proto.AESCTRDecryptResp)
	return resp, nil
}

// DES CBC Encrypt
func (*GRPCService) DESCBCEncrypt (ctx context.Context, req *proto.DESCBCEncryptReq) (*proto.DESCBCEncryptResp, error) {
	resp := new(proto.DESCBCEncryptResp)
	return resp, nil
}

// DES CBC Decrypt
func (*GRPCService) DESCBCDecrypt (ctx context.Context, req *proto.DESCBCDecryptReq) (*proto.DESCBCDecryptResp, error) {
	resp := new(proto.DESCBCDecryptResp)
	return resp, nil
}

// DES ECB Encrypt
func (*GRPCService) DESECBEncrypt (ctx context.Context, req *proto.DESECBEncryptReq) (*proto.DESECBEncryptResp, error) {
	resp := new(proto.DESECBEncryptResp)
	return resp, nil
}

// DES ECB Decrypt
func (*GRPCService) DESECBDecrypt (ctx context.Context, req *proto.DESECBDecryptReq) (*proto.DESECBDecryptResp, error) {
	resp := new(proto.DESECBDecryptResp)
	return resp, nil
}

// DES CFB Encrypt
func (*GRPCService) DESCFBEncrypt (ctx context.Context, req *proto.DESCFBEncryptReq) (*proto.DESCFBEncryptResp, error) {
	resp := new(proto.DESCFBEncryptResp)
	return resp, nil
}

// DES CFB Decrypt
func (*GRPCService) DESCFBDecrypt (ctx context.Context, req *proto.DESCFBDecryptReq) (*proto.DESCFBDecryptResp, error) {
	resp := new(proto.DESCFBDecryptResp)
	return resp, nil
}

// DES CTR Encrypt
func (*GRPCService) DESCTREncrypt (ctx context.Context, req *proto.DESCTREncryptReq) (*proto.DESCTREncryptResp, error) {
	resp := new(proto.DESCTREncryptResp)
	return resp, nil
}

// DES CTR Decrypt
func (*GRPCService) DESCTRDecrypt (ctx context.Context, req *proto.DESCTRDecryptReq) (*proto.DESCTRDecryptResp, error) {
	resp := new(proto.DESCTRDecryptResp)
	return resp, nil
}

// Hmac Md5
func (*GRPCService) HmacMd5 (ctx context.Context, req *proto.HmacMd5Req) (*proto.HmacMd5Resp, error) {
	resp := new(proto.HmacMd5Resp)
	return resp, nil
}

// Hmac Sha1
func (*GRPCService) HmacSha1 (ctx context.Context, req *proto.HmacSha1Req) (*proto.HmacSha1Resp, error) {
	resp := new(proto.HmacSha1Resp)
	return resp, nil
}

// Hmac Sha256
func (*GRPCService) HmacSha256 (ctx context.Context, req *proto.HmacSha256Req) (*proto.HmacSha256Resp, error) {
	resp := new(proto.HmacSha256Resp)
	return resp, nil
}

// Hmac Sha512
func (*GRPCService) HmacSha512 (ctx context.Context, req *proto.HmacSha512Req) (*proto.HmacSha512Resp, error) {
	resp := new(proto.HmacSha512Resp)
	return resp, nil
}


