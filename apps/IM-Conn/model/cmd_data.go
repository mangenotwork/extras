package model

import (
	"encoding/json"
	"github.com/mangenotwork/extras/common/jwt"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

type CmdData struct {
	Cmd string `json:"cmd"`
	Data interface{} `json:"data"`  // obj
	Msg string `json:"msg"`
	Code int `json:"code"`
	Token string `json:"token"`
}

func NewCmdData() *CmdData {
	return new(CmdData)
}

func (c *CmdData) Byte() []byte {
	if data, err := json.Marshal(c); err == nil {
		return data
	}else{
		return []byte(err.Error())
	}
}

func (c *CmdData) SendMsg(str string, code int) []byte {
	c.Msg = str
	return c.Byte()
}

func (c *CmdData) SendCmd(cmd string, data interface{}) []byte {
	c.Cmd = cmd
	c.Data = data
	return c.Byte()
}

func (c *CmdData) VerifyToken(clientUid int64) bool {
	logger.Debug("c.Token = ", c.Token, " | uid = ", clientUid)
	j := jwt.NewJWT()
	err := j.ParseToken(c.Token)
	if err != nil {
		logger.Debug("Token err = ", err)
		return false
	}

	uid := utils.Str2Int64(j.GetString("uid"))
	if uid == clientUid {
		return true
	}
	return false
}

func (c *CmdData) GetUid() string {
	j := jwt.NewJWT()
	err := j.ParseToken(c.Token)
	if err != nil {
		logger.Debug("Token err = ", err)
		return ""
	}

	return j.GetString("uid")
}