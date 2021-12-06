package raft

import (
	"bufio"
	"io"
	"log"
	"os"
)

// 读取 log.data 到内存
// 没有 log.data 则创建
func LogDataToMemory(){
	fileName := "log.data"

	var f *os.File
	var err error

	if checkFileExist(fileName) {  //文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
		if err != nil{
			log.Println("file open fail", err)
			return
		}
		// 读取文件
		defer f.Close()

		br := bufio.NewReader(f)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			log.Println(string(a))
		}

	}else {  //文件不存在
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			log.Println("file create fail")
			return
		}
	}

}

func SetTestData(){

	data := &LogData{
		Index : 1,
		Term : "aaaaa",
		Command : "add aaaaa",
	}
	data.Wait()

	data.Index = 2
	data.Term = "bbbbb"
	data.Command = "add bbbbb"
	data.Wait()

	data.Index = 3
	data.Term = "ccccc"
	data.Command = "add ccccc"
	data.Wait()
}