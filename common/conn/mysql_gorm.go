package conn

import (
	"fmt"
	"time"


	"github.com/jinzhu/gorm"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/logger"
)

/*
	gorm
	应用场景, 所有业务
 */

var mysqlGorm map[string]*gorm.DB

func InitMysqlGorm() {
	mysqlGorm = make(map[string]*gorm.DB, len(conf.Arg.Mysql))
	for _, v := range conf.Arg.Mysql {
		m, err := newORM(v.DBName, v.User, v.Password, v.Host, false)
		if err != nil {
			logger.Panic(err)
		}
		mysqlGorm[v.DBName] = m
	}
}

// newORM newORM
func newORM(database, user, password, host string, disablePrepared bool) (*gorm.DB, error) {
	var (
		orm *gorm.DB
		err error
	)
	if database == "" || user == "" || password == "" || host == "" {
		panic("数据库配置信息获取失败")
	}

	str := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database) + "?charset=utf8mb4&parseTime=true&loc=Local"
	if disablePrepared {
		str = str + "&interpolateParams=true"
	}
	for orm, err = gorm.Open("mysql", str); err != nil; {
		logger.Error(fmt.Sprintf("[DB]-[%v] 连接异常:%v，正在重试: %v", database, err, str))
		time.Sleep(5 * time.Second)
		orm, err = gorm.Open("mysql", str)
	}
	orm.LogMode(true)
	orm.CommonDB()
	return orm, err
}

func GetGorm(name string) *gorm.DB {
	m, ok := mysqlGorm[name]
	if !ok {
		logger.Panic("[DB] 未init")
	}
	return m
}
