package handler

import (
	"github.com/mangenotwork/extras/apps/IM-Conn/model"
	"github.com/mangenotwork/extras/common/jwt"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func WsHandler(client *model.WsClient, data *model.CmdData) {
	switch data.Cmd {
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
	}
}