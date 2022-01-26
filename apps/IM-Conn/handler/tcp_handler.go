package handler

import (
	"encoding/json"
	"github.com/mangenotwork/extras/apps/IM-Conn/model"
	"github.com/mangenotwork/extras/common/logger"
)

func TcpHandler(client *model.TcpClient, data *model.CmdData) {

	switch data.Cmd {
		case "Login":
			loginCmdData := NewLoginCmd()
			err := loginCmdData.Serialize(data.Data)
			if err != nil {
				logger.Errorf("%v",err)
				return
			}
			logger.Debugf("使用 json: %v",loginCmdData)
			// 登录验证 grpc
			uid, token, err := Login(loginCmdData.Account, loginCmdData.Password)
			if err == nil {
				client.UserID = uid
				client.DeviceID = loginCmdData.Device
				client.Source = loginCmdData.Source
				model.TcpClientTable().Insert(client)
				// 下发 token
				client.Send([]byte(token))
			}

		// TODO 心跳

	}
}

// LoginCmd 交互命令 Login
type LoginCmd struct {
	Account string `json:"account"`
	Password string `json:"password"`
	Device string `json:"device"`
	Source string `json:"source"`
}

func NewLoginCmd() *LoginCmd {
	return &LoginCmd{}
}

func (l *LoginCmd) Serialize(data interface{}) error {
	resByre, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(resByre, &l)
}