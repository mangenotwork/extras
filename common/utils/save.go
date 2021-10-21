package utils

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"encoding/gob"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// GobEncodeStr 将数据gob序列化存储到文件
func GobEncodeStr(data string) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	return buf.Bytes()
}

func GobDecoder(data []byte) (res string) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&res)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	return
}

// Compressed 压缩字符串
func Compressed(cache string) []byte {
	var content bytes.Buffer
	b := []byte(cache)
	w := zlib.NewWriter(&content)
	_,_=w.Write(b)
	_=w.Close()
	return content.Bytes()
}

// Decompress 解压字符串
func Decompress(data []byte) string {
	if len(data) < 1 {
		return ""
	}
	var out bytes.Buffer
	r, err := zlib.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return ""
	}
	_, err = io.Copy(&out, r)
	if err != nil {
		return ""
	}
	return out.String()
}

// 检查文件是否存在
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 文件写入数据
func FileWrite(fileName string, data []byte)  {
	var (
		f *os.File
		err error
	)
	if CheckFileExist(fileName) {
		f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil{
			log.Println("file open fail : ", err)
			return
		}
	}else {
		f, err = os.Create(fileName)
		if err != nil {
			log.Println("file create fail : ", err)
			return
		}
	}
	defer f.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(f)
	_,_=write.Write(data)
	//Flush将缓存的文件真正写入到文件中
	_=write.Flush()
	return
}

// 读取文件
func FileRead(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}