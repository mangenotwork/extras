package service

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/IM-User/dao"
	"github.com/mangenotwork/extras/apps/IM-User/model"
	"time"
)

type UserAPI interface {
	Register()
}

type UserParam struct {
	Name string `json:"name"` // 昵称
	Account string `json:"account"` // 账号
	Password string `json:"password"`
}

func (param *UserParam) verifyName() error {
	if len(param.Name) < 1 || len(param.Name) > 15 {
		return fmt.Errorf("昵称太长")
	}
	return nil
}

func (param *UserParam) verifyAccount() error {
	if len(param.Account) < 1 || len(param.Account) > 15 {
		return fmt.Errorf("账号太长")
	}
	return nil
}

func (param *UserParam) verifyPassword() error {
	if len(param.Password) < 6 {
		return fmt.Errorf("密码太短")
	}
	return nil
}

func (param *UserParam) Register() string {
	err := param.verifyName()
	if err != nil {
		return err.Error()
	}
	err = param.verifyAccount()
	if err != nil {
		return err.Error()
	}
	err = param.verifyPassword()
	if err != nil {
		return err.Error()
	}

	// 昵称, 账号是否重复
	has := new(dao.UserDao).UserBaseHas(param.Name, param.Account)
	if has {
		return "创建失败,账号或昵称已存在"
	}

	// 创建账号
	newUser := &model.UserBase{
		UName : param.Name,
		Account : param.Account,
		Password : param.Password,
		Created : time.Now().Unix(),
	}
	err = new(dao.UserDao).NewUser(newUser)
	if err != nil {
		return err.Error()
	}
	return "创建成功"
}