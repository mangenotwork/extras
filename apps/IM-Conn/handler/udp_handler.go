package handler

import (
	"github.com/mangenotwork/extras/apps/IM-Conn/model"
	"github.com/mangenotwork/extras/apps/IM-Conn/service"
	"github.com/mangenotwork/extras/common/logger"
)

func UdpHandler(client model.UdpClient, data *model.CmdData) {

	switch data.Cmd {
		// 登录
		case "Login":
			loginCmdData := service.NewLoginCmd()
			err := loginCmdData.Serialize(data.Data)
			if err != nil {
				logger.Errorf("%v",err)
				return
			}
			logger.Debugf("使用 json: %v",loginCmdData)
			// 登录验证 grpc
			uid, token, err := Login(loginCmdData.Account, loginCmdData.Password)
			if err == nil {
				logger.Debug("写入uid = ", uid)
				client.UserID = uid
				logger.Debug("Login client.UserID = ", client.UserID)
				client.DeviceID = loginCmdData.Device
				client.Source = loginCmdData.Source
				model.UdpClientTable().Insert(&client)
				// 下发 token
				sendToken := model.NewCmdData()
				client.Send(sendToken.SendCmd("Token", token))
			}

		case "Hello":
			uid := data.GetUid()
			logger.Debug("Hello uid = ", uid)
			if len(uid) == 0  {
				sendErr := model.NewCmdData()
				client.Send(sendErr.SendCmd("Error", "token is error!"))
			}
			c, err := model.UdpClientTable().Get(uid)
			logger.Debug(c, err)
	}
}