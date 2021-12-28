package raft

import (
	"bytes"
	"io"
	"os"
	"strings"

	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

/* 日志

每个节点存储自己的日志副本(log[])，每条日志记录包含：

索引：该记录在日志中的位置
任期号：该记录首次被创建时的任期号
命令

*/

type LogData struct {
	Index int64
	Term string // 任期号
	Command string
}

func NewLogData(command string) *LogData {
	Index++
	logger.Info("Index = ", Index)
	return &LogData{
		Index : Index,
		Term : MyAddr,
		Command : command,
	}
}

func (data *LogData) ToStr() string {
	var buffer bytes.Buffer
	buffer.WriteString(utils.Any2String(data.Index))
	buffer.WriteString("&")
	buffer.WriteString(data.Term)
	buffer.WriteString("&")
	buffer.WriteString(data.Command)
	buffer.WriteString("\n")
	return buffer.String()
}

func (data *LogData) ToObj(str string){
	strList := strings.Split(str, "&")
	if len(strList) == 3 {
		data.Index = utils.Str2Int64(strList[0])
		data.Term = strList[1]
		data.Command = strList[2]
	}
}

// 追加写入日志
func (data *LogData) Write(){
	fileName := "log.data"

	var f *os.File
	var err error

	if utils.CheckFileExist(fileName) {  //文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0666) //打开文件
		if err != nil{
			logger.Error("file open fail", err)
			return
		}
	}else {  //文件不存在
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			logger.Error("file create fail")
			return
		}
	}

	strTest := data.ToStr()

	//将文件写进去
	n, err1 := io.WriteString(f, strTest)
	if err1 != nil {
		logger.Error("write error", err1)
		return
	}
	logger.Info("写入的字节数是：", n)


	_=f.Close()
}
