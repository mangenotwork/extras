package handler

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024*100,
	WriteBufferSize: 65535,
	HandshakeTimeout: 5*time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_,_=w.Write([]byte("Hello ManGe"))
}

func Ws(w http.ResponseWriter, r *http.Request) {
	st := time.Now()
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("websocket upgrade error:%v", err)
		return
	}
	log.Println("[连接日志] 连接成功. 用时 = ", time.Now().Sub(st))
	log.Println(conn)
}