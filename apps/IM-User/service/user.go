package service

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/IM-User/dao"
	"github.com/mangenotwork/extras/apps/IM-User/model"
	"github.com/mangenotwork/extras/common/jwt"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
	"time"
)

type UserAPI interface {
	Register() string // 用户注册业务
	Token()	(string, error) // 获取用户token
	AuthToken() (int64, error) // 用户token 的验证
}

// UserParam 用户业务相关参数
type UserParam struct {
	Name string `json:"name"` // 昵称
	Account string `json:"account"` // 账号
	Password string `json:"password"`
	TokenStr string `json:"token_str"` // 用户token
}

// verifyName 验证用户昵称
func (param *UserParam) verifyName() error {
	if len(param.Name) < 1 || len(param.Name) > 15 {
		return fmt.Errorf("昵称太长")
	}
	return nil
}

// verifyAccount 验证用户账号
func (param *UserParam) verifyAccount() error {
	if len(param.Account) < 1 || len(param.Account) > 15 {
		return fmt.Errorf("账号太长")
	}
	return nil
}

// verifyPassword 验证用户密码
func (param *UserParam) verifyPassword() error {
	if len(param.Password) < 6 {
		return fmt.Errorf("密码太短")
	}
	return nil
}

// Register 用户注册
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

	// 会写id


	return "创建成功"
}

// Token 获取用户token
func (param *UserParam) Token() (string, string,  error) {
	err := param.verifyPassword()
	if err != nil {
		return "", "", err
	}
	u := new(dao.UserDao).GetFromAccount(param.Account)
	if param.Password != u.Password {
		return "", "", fmt.Errorf("密码错误")
	}

	j := jwt.NewJWT()
	j.AddClaims("uid", u.UId)
	j.AddClaims("name", u.UName)
	j.AddClaims("isok", u.Account)
	token, err := j.Token()
	return token, u.UId, err
}

// AuthToken 用户token 的验证
// return  ->  1: 验证通过  2: 验证失败
func (param *UserParam) AuthToken() (int64, string, error) {

	j := jwt.NewJWT()
	err := j.ParseToken(param.TokenStr)
	if err != nil {
		return 2, "", err
	}

	// 是否过期
	if j.IsExpire() {
		return 2, "", fmt.Errorf("token 过期")
	}

	// 是否存在
	uid := utils.Any2String(j.Get("uid"))
	logger.Debug("uid = ", uid)
	tid, id := model.SplitUId(uid)
	logger.Debug("tid, id = ", tid, id)
	if !new(dao.UserDao).HasFromUid(tid, id) {
		return 2, "", fmt.Errorf("用户不存在")
	}

	return 1, uid, nil
}