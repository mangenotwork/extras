package engine

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mangenotwork/extras/apps/IM-Conn/model"
	"github.com/mangenotwork/extras/common/httpser"
	"github.com/mangenotwork/extras/common/logger"
	"net/http"
	"time"
)

func StartWS(){
	go func() {
		logger.Info("StartWS")

		mux := httpser.NewEngine()

		mux.Router("/ws", Ws)

	}()
}

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024*100,
	WriteBufferSize: 65535,
	HandshakeTimeout: 5*time.Second,
	Error: Err,
	CheckOrigin: func(r *http.Request) bool {
		if r.Method != "GET" {
			fmt.Println("method is not GET")
			return false
		}
		if r.URL.Path != "/ws" {
			fmt.Println("path error")
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

	// TODO Token JWT解码 得出 user_id, device_id, source

	userId := httpser.GetUrlArgInt64(r, "user_id")
	deviceId := httpser.GetUrlArg(r, "device")
	source := httpser.GetUrlArg(r, "source")
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("websocket upgrade error:%v", err)
		return
	}

	// TODO 身份验证

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
		UserID : userId,
		IP : conn.RemoteAddr().String(),
		DeviceID : deviceId,
		DeviceType : source,
	}

	model.WsClientTable().Insert(client) // 加入ws哈希表

	// TODO 接收数据

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

			//cmdData := &model.CmdData{}
			//jsonErr := json.Unmarshal(data, &cmdData)
			//if jsonErr != nil {
			//	_=conn.WriteMessage(websocket.BinaryMessage, model.CmdDataMsg("非法数据格式"))
			//	continue
			//}
			//logger.Info(cmdData)
			//device = service.Interactive(cmdData, client)
		}

	}()

}

func Err(w http.ResponseWriter, r *http.Request, status int, reason error) {
	httpser.OutErrBody(w, status, reason)
}

