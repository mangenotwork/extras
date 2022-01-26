package handler

import (
	"github.com/mangenotwork/extras/apps/IM-Conn/model"
	"github.com/mangenotwork/extras/apps/IM-Conn/service"
	"github.com/mangenotwork/extras/common/jwt"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func TcpHandler(client *model.TcpClient, data *model.CmdData) {

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
				client.UserID = uid
				client.DeviceID = loginCmdData.Device
				client.Source = loginCmdData.Source
				model.TcpClientTable().Insert(client)
				// 下发 token
				sendToken := model.NewCmdData()
				client.Send(sendToken.SendCmd("Token", token))
			}

		// 心跳
		case "HeartBeat":
			logger.Debug("HeartBeat doing ...")
			token := utils.Any2String(data.Data)
			logger.Debug("token = ", token)
			j := jwt.NewJWT()
			err := j.ParseToken(token)
			if err != nil {
				logger.Debug("Token err = ", err)
				sendToken := model.NewCmdData()
				client.Send(sendToken.SendCmd("Error", "Token err"))
				return
			}

			uid := j.GetString("uid")
			if uid == client.UserID {
				client.HeartBeat <- []byte("1")
			}

		//


	}
}

