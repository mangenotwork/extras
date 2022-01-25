package dao

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/IM-User/global"
	"github.com/mangenotwork/extras/apps/IM-User/model"
	"github.com/mangenotwork/extras/common/conn"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
	"sync"
)

type UserDao struct {

}

// UserBaseHas 验证用户信息是否已经存在
func (dao *UserDao) UserBaseHas(name, account string) bool {
	var (
		wg sync.WaitGroup
		has = false
	)

	for i:=0; i<global.MaxUserBaseTable; i++ {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			var count  int
			table := fmt.Sprintf(global.UserBaseTableName, number)

			db := conn.GetGorm("imtest")
			err := db.Table(table).Where("uname=? or account=?", name, account).Count(&count).Error
			if err != nil {
				logger.Error(err)
			}

			//logger.Debug("count = ", count)
			if count > 0 {
				has = true
			}
		}(i)
	}
	wg.Wait()

	if has {
		logger.Debug("存在 昵称或账号")
	}

	return has
}

// NewUser 新建用户
func (dao *UserDao) NewUser(u *model.UserBase) error {
	db := conn.GetGorm("imtest")
	u.TableId = utils.RandInt(0, global.MaxUserBaseTable)
	err := db.Table(u.TableName()).Create(u).Error
	if err != nil {
		return err
	}
	u.UId = fmt.Sprintf("%04d%d", u.TableId, u.ID)
	return db.Model(&u).Update("uid", u.UId).Error
}

// GetFromAccount 获取用户来自参数 Account
func (dao *UserDao) GetFromAccount(account string) *model.UserBase {
	var (
		wg sync.WaitGroup
		user = &model.UserBase{}
	)

	for i:=0; i<global.MaxUserBaseTable; i++ {
		wg.Add(1)
		go func(number int) {
			defer wg.Done()
			table := fmt.Sprintf(global.UserBaseTableName, number)
			db := conn.GetGorm("imtest")
			err := db.Table(table).Where("account=?", account).First(&user).Error
			if err != nil {
				logger.Error(err)
			}
		}(i)
	}
	wg.Wait()

	logger.Debug(user)
	return user
}

// HasFromUid 用户是否存在
func (dao *UserDao) HasFromUid(tid, id int) bool {
	var count  int
	table := fmt.Sprintf(global.UserBaseTableName, tid)
	db := conn.GetGorm("imtest")
	err := db.Table(table).Where("id=?", id).Count(&count).Error
	if err != nil {
		logger.Error(err)
		return false
	}

	logger.Debug("count = ", count)
	if count > 0 {
		return true
	}
	return false
}