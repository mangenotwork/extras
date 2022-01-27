package engine

import (
	"encoding/json"
	"fmt"
	"github.com/mangenotwork/extras/apps/IM-Conn/handler"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mangenotwork/extras/apps/IM-Conn/model"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
)

func StartWS(){
	go func() {
		logger.Info("StartWS")

		mux := httpser.NewEngine()

		mux.RouterFunc("/ws", Ws)

		mux.Run()
	}()
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  65535,
	WriteBufferSize: 65535,
	HandshakeTimeout: 5*time.Second,
	Error: Err,
	CheckOrigin: func(r *http.Request) bool {
		if r.Method != "GET" {
			logger.Error("method is not GET")
			return false
		}
		if  strings.Index(r.URL.Path, "/ws") == -1 {
			logger.Error("path error")
			return false
		}
		return true
	},
}


func Ws(w http.ResponseWriter, r *http.Request) {
	if !websocket.IsWebSocketUpgrade(r) {
		Err(w, r, 2001, fmt.Errorf("非websocket请求"))
	}

	st := time.Now()
	token := httpser.GetUrlArg(r, "token")
	deviceId := httpser.GetUrlArg(r, "device")
	source := httpser.GetUrlArg(r, "source")
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("websocket upgrade error:%v", err)
		return
	}

	uid, err := handler.AuthToken(token)
	if err != nil {
		_= conn.WriteMessage(websocket.BinaryMessage, []byte("[未知身份连接]device为空, 客户端可以send数据来确认身份"))
		_= conn.Close()
		Err(w, r, 2001, fmt.Errorf("[未知身份连接]device为空, 客户端可以send数据来确认身份"))
		return
	}

	if len(deviceId) < 1 {
		_= conn.WriteMessage(websocket.BinaryMessage, []byte("[未知身份连接]device为空, 客户端可以send数据来确认身份"))
		_= conn.Close()
		Err(w, r, 2001, fmt.Errorf("[未知身份连接]device为空, 客户端可以send数据来确认身份"))
		return
	}
	logger.Info("[连接日志] 连接成功. 用时 = ", time.Now().Sub(st))
	logger.Info("RemoteAddr = ", conn.RemoteAddr(), " | LocalAddr = ", conn.LocalAddr())

	client := &model.WsClient{
		Conn : conn,
		UserID : uid,
		IP : conn.RemoteAddr().String(),
		DeviceID : deviceId,
		DeviceType : source,
		HeartBeat : make(chan []byte),
	}

	model.WsClientTable().Insert(client) // 加入ws哈希表

	//接收数据
	go func() {

		for {
			_, data, err := client.Conn.ReadMessage()
			if err != nil {
				_=client.Conn.Close()
				// 释放客户端连接
				model.WsClientTable().Del(client)
				return
			}
			logger.Info(data)
			if len(data) < 1 {
				continue
			}

			logger.Info(string(data))
			cmdData := &model.CmdData{}
			jsonErr := json.Unmarshal(data, &cmdData)
			if jsonErr != nil {
				client.Send(cmdData.SendMsg("非法数据格式", 1001))
				continue
			}
			handler.WsHandler(client, cmdData)

		}

	}()


	// 处理心跳和僵尸连接
	go func(){
		for {
			timer := time.NewTimer(10 * time.Second)

			select {
			case <-timer.C:
				// 10秒内收不到来自客户端的心跳连接断开
				_=client.Conn.Close()
				model.WsClientTable().Del(client)

				// 接收心跳
			case <-client.HeartBeat:
				// 重置
				log.Println("重置 = ", 10)
				timer.Reset(10 * time.Second)
			}
		}

	}()

}

func Err(w http.ResponseWriter, r *http.Request, status int, reason error) {
	httpser.OutErrBody(w, status, reason)
}

