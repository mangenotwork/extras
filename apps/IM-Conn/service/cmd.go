package service

import "encoding/json"

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