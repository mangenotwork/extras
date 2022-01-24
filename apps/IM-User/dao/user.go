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

func (dao *UserDao) NewUser(u *model.UserBase) error {
	db := conn.GetGorm("imtest")
	u.TableId = utils.RandInt(0, global.MaxUserBaseTable)
	return db.Table(u.TableName()).Create(u).Error
}