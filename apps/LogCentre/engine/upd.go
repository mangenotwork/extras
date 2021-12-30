package engine

import (
	"encoding/json"
	"net"
	"os"
	"strings"
	"time"

	"github.com/mangenotwork/extras/common/boltdb"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

var BoltdbFileName = "data.db"

func StartUdp(){
	go func() {

		logger.Info("StartUdpServer")
		var err error
		var udpListener *net.UDPConn

		// 监听
		udpListener, err = net.ListenUDP("udp", &net.UDPAddr{
			IP: net.ParseIP("0.0.0.0"),
			Port: utils.Str2Int(conf.Arg.UdpServer.Prod),
		})
		if err != nil {
			logger.Error(err)
			return
		}
		logger.Info("Local: <%s> \n", udpListener.LocalAddr().String())

		// 读取数据
		data := make([]byte, 1024*20)
		for {

			n, remoteAddr, err := udpListener.ReadFromUDP(data)
			if err != nil {
				logger.Error("error during read: %s", err)
			}

			if n > 0 && n < 1024*20 {
				logger.Debug("收到日志")
				dataType := data[0:1]
				ip := remoteAddr.IP.String()
				switch string(dataType) {
				case "1":
					// 日志写入文件
					Log2File(data, n, ip)
				case "2":
					// 日志是http日志 记录到db
					HttpLog(data, n)
				case "3":
					// 日志是grpc日志 记录到db
					GrpcLog(data, n)
				}

			}else{
				logger.Error("传入的数据太小或太大, 建议 1~10240个字节")
			}
		}

	}()
}

func Log2File(data []byte, n int, ip string){
	dataStr := string(data[1:n])
	name := ip
	nameList := strings.Split(dataStr, "|")
	if len(nameList) > 0 {
		name = nameList[0]
	}
	writer, err := os.OpenFile( "logs/"+name+time.Now().Format("-20060102")+".log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Error("打开日志文件失败  = ", err)
	}
	d := []byte(ip+"|")
	d = append(d, data[1:n]...)
	writer.Write(d)
}

func HttpLog(data []byte, n int){
	//logger.Info("http 日志 ", string(data[1:n]))
	dataList := strings.Split(string(data[1:n]), "#")
	name := dataList[0]
	timestamp := dataList[1]
	reqIp := dataList[2]
	reqMethod := dataList[3]
	reqUrl := dataList[4]
	reqCode := dataList[5]
	//reqTime := utils.Str2Float64(dataList[5])
	reqTime := dataList[5]
	//logger.Info("reqIp = ", reqIp, " | reqMethod = ", reqMethod, " | reqUrl = ", reqUrl, "| reqCode", reqCode, "| reqTime = ", reqTime)

	// 数据存储到 boltdb
	tableName := name + "-http-req"
	bo, err := boltdb.NewBoltDB(BoltdbFileName)
	if err != nil {
		logger.Error(err)
	}
	defer bo.Close()

	err = bo.CreateTable(tableName)
	if err != nil {
		logger.Error(err)
	}
	d := HttpRepLog{
		AppName: name,
		Timestamp: timestamp,
		Ip: reqIp,
		Method: reqMethod,
		Url: reqUrl,
		Code: reqCode,
		ReqTime: reqTime,
	}
	b, err := json.Marshal(d)
	if err != nil {
		logger.Error(err)
	}
	err = bo.Insert(tableName, timestamp, string(b))
	if err != nil {
		logger.Error(err)
	}


}

func GrpcLog(data []byte, n int){
	//logger.Info("http 日志 ", string(data[1:n]))
	dataList := strings.Split(string(data[1:n]), "#")
	name := dataList[0]
	timestamp := dataList[1]
	state := dataList[2]
	link := dataList[3]
	requestId := dataList[4]
	method := dataList[5]
	reqTime := dataList[6]
	//logger.Info("state = ", state, " | link = ", link, " | requestId = ", requestId, "| method", method, "| reqTime = ", reqTime)

	// 数据存储到 boltdb
	tableName := name + "grpc"
	bo, err := boltdb.NewBoltDB(BoltdbFileName)
	if err != nil {
		logger.Error(err)
	}
	defer bo.Close()

	err = bo.CreateTable(tableName)
	if err != nil {
		logger.Error(err)
	}
	d := GrpcReqLog{
		AppName: name,
		Timestamp: timestamp,
		State: state,
		Link: link,
		RequestId: requestId,
		Method: method,
		ReqTime: reqTime,
	}
	b, err := json.Marshal(d)
	if err != nil {
		logger.Error(err)
	}
	err = bo.Insert(tableName, timestamp, string(b))
	if err != nil {
		logger.Error(err)
	}
}

type HttpRepLog struct {
	AppName string `json:"app_name"`
	Timestamp string `json:"timestamp"`
	Ip string `json:"ip"`
	Method string `json:"method"`
	Url string `json:"url"`
	Code string `json:"code"`
	ReqTime string `json:"req_time"`
}

type GrpcReqLog struct {
	AppName string `json:"app_name"`
	Timestamp string `json:"timestamp"`
	State string `json:"state"`
	Link string `json:"link"`
	RequestId  string `json:"request_id"`
	Method string `json:"method"`
	ReqTime string `json:"req_time"`
}
