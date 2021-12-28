package service

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/mangenotwork/extras/apps/ServiceTable/model"
	"github.com/mangenotwork/extras/apps/ServiceTable/raft"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func InitRaft(){
	raft.MyAddr = conf.Arg.Cluster.MyAddr
	raft.Cluster = strings.Split(conf.Arg.Cluster.InitCluster, ";")
	raft.IsCluster = conf.Arg.Cluster.Open
	model.InitSetData()
	model.InitMapData()
}


// 读取 log.data 到内存
// 没有 log.data 则创建
func LogDataToMemory(){
	fileName := "log.data"

	var f *os.File
	var err error

	if utils.CheckFileExist(fileName) {  //文件存在
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
		if err != nil{
			logger.Error("file open fail", err)
			return
		}
		// 读取文件
		defer f.Close()
		data := raft.LogData{}
		br := bufio.NewReader(f)
		for {
			a, _, c := br.ReadLine()
			if c == io.EOF {
				break
			}
			logger.Info(string(a))
			data.ToObj(string(a))
			CommandDo(data.Command)
		}
		raft.Index = data.Index

	}else {  //文件不存在
		f, err = os.Create(fileName) //创建文件
		if err != nil {
			logger.Error("file create fail")
			return
		}
	}

}



func CommandDo(cmdStr string) {
	cmdArg := strings.Split(cmdStr, " ")
	l := len(cmdArg)
	if l < 1 {
		return
	}
	cmd := cmdArg[0]
	switch cmd {
	case "SetAdd":
		// Command : SetAdd key value1,value2,
		if l < 3 {
			return
		}
		key := cmdArg[1]
		values := cmdArg[2]
		model.SetAdd(key, utils.SliceDelNullString(strings.Split(values, ",")))

	case "SetAddExpire":
		// Command : SetAddExpire key value timeUnix
		if l < 4 {
			return
		}
		key := cmdArg[1]
		value := cmdArg[2]
		timeUnix := cmdArg[3]
		model.SetValueExpire(key, value, utils.Str2Int64(timeUnix))

	case "SetValueExpire":
		// Command : SetValueExpire key value timeUnix
		if l < 4 {
			return
		}
		key := cmdArg[1]
		value := cmdArg[2]
		timeUnix := cmdArg[3]
		model.SetValueExpire(key, value, utils.Str2Int64(timeUnix))

	case "SetGet":
		// Command : SetGet key
		// 读操作跳过

	case "SetDel":
		// Command : SetDel key
		if l < 2 {
			return
		}
		key := cmdArg[1]
		model.SetDel(key)

	case "SetDelValue":
		// Command : SetDelValue key value
		if l < 3 {
			return
		}
		key := cmdArg[1]
		value := cmdArg[2]
		model.SetDelValue(key, value)

	case "KVAdd":
		// Command : KVAdd key value
		if l < 3 {
			return
		}
		key := cmdArg[1]
		value := cmdArg[2]
		logger.Info("Command : KVAdd key value")
		model.KVAdd(key, value)

	case "KVAddExpire":
		// Command : KVAddExpire key value expire
		if l < 4 {
			return
		}
		key := cmdArg[1]
		value := cmdArg[2]
		expire := cmdArg[3]
		model.KVAddExpire(key, value, utils.Str2Int64(expire))

	case "KVExpire":
		// Command : KVExpire key expire
		if l < 3 {
			return
		}
		key := cmdArg[1]
		expire := cmdArg[2]
		model.KVExpire(key, utils.Str2Int64(expire))

	case "KVDel":
		// Command : KVDel key
		if l < 2 {
			return
		}
		key := cmdArg[1]
		model.KVDel(key)


	}
}