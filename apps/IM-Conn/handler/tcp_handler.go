package handler

import (
	"encoding/json"
	"github.com/mangenotwork/extras/apps/IM-Conn/model"
	"github.com/mangenotwork/extras/common/logger"
)

func TcpHandler(client *model.TcpClient, data *model.CmdData) {

	//device = service.Interactive(cmdData, client)
	// TODO 获取来自客服端的身份信息,并验证
	// client.UserID = ...
	// if 验证失败 { _=conn.Close() }
	// model.TcpClientTable().Insert(client)
	// TODO 心跳

	switch data.Cmd {
		case "Login":
			resByre, resByteErr := json.Marshal(data.Data)
			if resByteErr != nil {
				logger.Debugf("%v",resByteErr)
				return
			}
			var LoginCmdData LoginCmd
			jsonRes := json.Unmarshal(resByre, &LoginCmdData)
			if jsonRes != nil {
				logger.Debugf("%v",jsonRes)
				return
			}
			logger.Debugf("使用 json: %v",LoginCmdData)
			// 登录验证 grpc

	}
}

type LoginCmd struct {
	Account string `json:"account"`
	Password string `json:"password"`
	Device string `json:"device"`
	Source string `json:"source"`
}