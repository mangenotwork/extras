package model

import "encoding/json"

type CmdData struct {
	Cmd string `json:"cmd"`
	Data interface{} `json:"data"`
}

func CmdDataMsg(str string) []byte {
	msg := &CmdData{
		Cmd: "Message",
		Data: str,
	}
	return msg.Byte()
}

func (c *CmdData) Byte() []byte {
	if data, err := json.Marshal(c); err == nil {
		return data
	}else{
		return []byte(err.Error())
	}
}

type AuthData struct {
	Device string `json:"device"`
}
